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

package v20210108

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type CreateKnowledgeSetRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 知识库标识
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 知识库名称
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 描述
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 知识库的meta信息
	Meta *string `json:"Meta,omitnil,omitempty" name:"Meta"`
}

type CreateKnowledgeSetRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 知识库标识
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 知识库名称
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 描述
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 知识库的meta信息
	Meta *string `json:"Meta,omitnil,omitempty" name:"Meta"`
}

func (r *CreateKnowledgeSetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateKnowledgeSetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Name")
	delete(f, "Title")
	delete(f, "Desc")
	delete(f, "Meta")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateKnowledgeSetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateKnowledgeSetResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateKnowledgeSetResponse struct {
	*tchttp.BaseResponse
	Response *CreateKnowledgeSetResponseParams `json:"Response"`
}

func (r *CreateKnowledgeSetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateKnowledgeSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataSourceDetail struct {
	// 数据源 ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 数据源名称
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 数据源标识
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数据源类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 数据源描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 数据源配置
	Schema *string `json:"Schema,omitnil,omitempty" name:"Schema"`

	// cms 项目状态, 0: 重新获取详情信息，1： 不需要重新获取详情信息
	CmsProject *string `json:"CmsProject,omitnil,omitempty" name:"CmsProject"`

	// 当前为环境 id
	PkgId *string `json:"PkgId,omitnil,omitempty" name:"PkgId"`

	// schema 版本信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchemaVersion *string `json:"SchemaVersion,omitnil,omitempty" name:"SchemaVersion"`

	// 创建者用户 ID
	CreatorId *string `json:"CreatorId,omitnil,omitempty" name:"CreatorId"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// 环境 id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSourceVersion *string `json:"DataSourceVersion,omitnil,omitempty" name:"DataSourceVersion"`

	// 所属应用数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppUsageList []*DataSourceLinkApp `json:"AppUsageList,omitnil,omitempty" name:"AppUsageList"`

	// 发布时间
	PublishedAt *string `json:"PublishedAt,omitnil,omitempty" name:"PublishedAt"`

	// 子数据源ids
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChildDataSourceIds []*string `json:"ChildDataSourceIds,omitnil,omitempty" name:"ChildDataSourceIds"`

	// 数据源发布信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Fun *string `json:"Fun,omitnil,omitempty" name:"Fun"`

	// 云函数状态 1 Active 2 Creating 3 Updating 4 Deleting  9 Deleted 11 CreatFailed  12 UpdateFailed 13 DeleteFailed 21 UpdateTimeOut
	ScfStatus *uint64 `json:"ScfStatus,omitnil,omitempty" name:"ScfStatus"`

	// 自定义方法
	Methods *string `json:"Methods,omitnil,omitempty" name:"Methods"`

	// 子数据源名数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChildDataSourceNames []*string `json:"ChildDataSourceNames,omitnil,omitempty" name:"ChildDataSourceNames"`

	// 是否旧数据源 1 新 0 旧
	IsNewDataSource *int64 `json:"IsNewDataSource,omitnil,omitempty" name:"IsNewDataSource"`

	// 数据源视图id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ViewId *string `json:"ViewId,omitnil,omitempty" name:"ViewId"`

	// 数据源属性配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Configuration *string `json:"Configuration,omitnil,omitempty" name:"Configuration"`

	// 外部数据源模板code
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateCode *string `json:"TemplateCode,omitnil,omitempty" name:"TemplateCode"`

	// 外部数据源模板来源 0 空模板 1 腾讯文档 2 腾讯会议 3 企业微信 4 微信电商
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *int64 `json:"Source,omitnil,omitempty" name:"Source"`

	// 发布版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: PublishVersion is deprecated.
	PublishVersion *string `json:"PublishVersion,omitnil,omitempty" name:"PublishVersion"`

	// 发布视图id
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublishViewId *string `json:"PublishViewId,omitnil,omitempty" name:"PublishViewId"`

	// 数据源子类型   "database" 标准模型 "custom-database" 自定义模型 "system" 系统模型 "connector" 连接器 "custom-connector" 自定义连接器 "hidden" 隐藏数据源
	SubType *string `json:"SubType,omitnil,omitempty" name:"SubType"`

	// 授权状态  0 授权无效 1 授权有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthStatus *int64 `json:"AuthStatus,omitnil,omitempty" name:"AuthStatus"`

	// 数据源授权信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthInfo *TicketAuthInfo `json:"AuthInfo,omitnil,omitempty" name:"AuthInfo"`

	// 1发布0未发布
	PublishStatus *int64 `json:"PublishStatus,omitnil,omitempty" name:"PublishStatus"`

	// 更新版本
	UpdateVersion *int64 `json:"UpdateVersion,omitnil,omitempty" name:"UpdateVersion"`

	// 模型关联关系字段列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelationFieldList []*RelationField `json:"RelationFieldList,omitnil,omitempty" name:"RelationFieldList"`

	// db实例类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbInstanceType *string `json:"DbInstanceType,omitnil,omitempty" name:"DbInstanceType"`

	// 体验环境db表名
	// 注意：此字段可能返回 null，表示取不到有效值。
	PreviewTableName *string `json:"PreviewTableName,omitnil,omitempty" name:"PreviewTableName"`

	// 正式环境db表名
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublishedTableName *string `json:"PublishedTableName,omitnil,omitempty" name:"PublishedTableName"`

	// DB来源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbSourceType *string `json:"DbSourceType,omitnil,omitempty" name:"DbSourceType"`
}

type DataSourceDetailItems struct {
	// 数据详情列表
	Rows []*DataSourceDetail `json:"Rows,omitnil,omitempty" name:"Rows"`

	// 数据源列表总个数
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type DataSourceLinkApp struct {
	// 应用Id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 应用名称
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 是否编辑状态使用
	EditStatusUse *int64 `json:"EditStatusUse,omitnil,omitempty" name:"EditStatusUse"`

	// 是否预览状态使用
	PreviewStatusUse *int64 `json:"PreviewStatusUse,omitnil,omitempty" name:"PreviewStatusUse"`

	// 是否正式状态使用
	OnlineStatusUse *int64 `json:"OnlineStatusUse,omitnil,omitempty" name:"OnlineStatusUse"`

	// 数据源ID
	DataSourceId *string `json:"DataSourceId,omitnil,omitempty" name:"DataSourceId"`
}

type DataSourceQueryOption struct {
	// 数据源标识模糊匹配
	LikeName *string `json:"LikeName,omitnil,omitempty" name:"LikeName"`

	// 数据源名称模糊匹配
	LikeTitle *string `json:"LikeTitle,omitnil,omitempty" name:"LikeTitle"`
}

// Predefined struct for user
type DeleteKnowledgeDocumentSetRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 知识库标识
	CollectionView *string `json:"CollectionView,omitnil,omitempty" name:"CollectionView"`

	// 删除时制定的条件
	Query *DocumentQuery `json:"Query,omitnil,omitempty" name:"Query"`
}

type DeleteKnowledgeDocumentSetRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 知识库标识
	CollectionView *string `json:"CollectionView,omitnil,omitempty" name:"CollectionView"`

	// 删除时制定的条件
	Query *DocumentQuery `json:"Query,omitnil,omitempty" name:"Query"`
}

func (r *DeleteKnowledgeDocumentSetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteKnowledgeDocumentSetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "CollectionView")
	delete(f, "Query")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteKnowledgeDocumentSetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteKnowledgeDocumentSetResponseParams struct {
	// 新增文件返回信息
	Data *DeleteKnowledgeDocumentSetRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteKnowledgeDocumentSetResponse struct {
	*tchttp.BaseResponse
	Response *DeleteKnowledgeDocumentSetResponseParams `json:"Response"`
}

func (r *DeleteKnowledgeDocumentSetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteKnowledgeDocumentSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteKnowledgeDocumentSetRsp struct {
	// 删除文档数量。
	AffectedCount *int64 `json:"AffectedCount,omitnil,omitempty" name:"AffectedCount"`
}

// Predefined struct for user
type DeleteKnowledgeSetRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 知识库标识
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DeleteKnowledgeSetRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 知识库标识
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *DeleteKnowledgeSetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteKnowledgeSetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteKnowledgeSetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteKnowledgeSetResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteKnowledgeSetResponse struct {
	*tchttp.BaseResponse
	Response *DeleteKnowledgeSetResponseParams `json:"Response"`
}

func (r *DeleteKnowledgeSetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteKnowledgeSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataSourceListRequestParams struct {
	// 每页条数
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页码
	PageIndex *int64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`

	// 环境 id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 应用id数组
	Appids []*string `json:"Appids,omitnil,omitempty" name:"Appids"`

	// 数据源id数组
	DataSourceIds []*string `json:"DataSourceIds,omitnil,omitempty" name:"DataSourceIds"`

	// 数据源名称数组
	DataSourceNames []*string `json:"DataSourceNames,omitnil,omitempty" name:"DataSourceNames"`

	// 数据源类型 database-自建数据源；cloud-integration-自定义数据源
	DataSourceType *string `json:"DataSourceType,omitnil,omitempty" name:"DataSourceType"`

	// 数据源模糊查询参数
	QueryOption *DataSourceQueryOption `json:"QueryOption,omitnil,omitempty" name:"QueryOption"`

	// 数据源视图Id数组
	ViewIds []*string `json:"ViewIds,omitnil,omitempty" name:"ViewIds"`

	// 查询未关联应用的数据源，0:未关联，该参数配合 AppIds 参数一块使用
	AppLinkStatus *int64 `json:"AppLinkStatus,omitnil,omitempty" name:"AppLinkStatus"`

	// 查询应用绑定数据源: 0: 否,1: 是
	QueryBindToApp *int64 `json:"QueryBindToApp,omitnil,omitempty" name:"QueryBindToApp"`

	// 查询连接器 0 数据模型 1 连接器 2 自定义连接器
	QueryConnector *int64 `json:"QueryConnector,omitnil,omitempty" name:"QueryConnector"`

	// 废弃中
	NotQuerySubTypeList []*string `json:"NotQuerySubTypeList,omitnil,omitempty" name:"NotQuerySubTypeList"`

	// 查询channelList
	ChannelList []*string `json:"ChannelList,omitnil,omitempty" name:"ChannelList"`

	// 是否查询数据源关联关系
	QueryDataSourceRelationList *bool `json:"QueryDataSourceRelationList,omitnil,omitempty" name:"QueryDataSourceRelationList"`

	// db实例类型
	DbInstanceType *string `json:"DbInstanceType,omitnil,omitempty" name:"DbInstanceType"`

	// 数据库表名列表
	DatabaseTableNames []*string `json:"DatabaseTableNames,omitnil,omitempty" name:"DatabaseTableNames"`

	// 是否查询系统模型，默认为true，需要显示设置为False才能过滤系统模型
	QuerySystemModel *bool `json:"QuerySystemModel,omitnil,omitempty" name:"QuerySystemModel"`
}

type DescribeDataSourceListRequest struct {
	*tchttp.BaseRequest
	
	// 每页条数
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页码
	PageIndex *int64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`

	// 环境 id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 应用id数组
	Appids []*string `json:"Appids,omitnil,omitempty" name:"Appids"`

	// 数据源id数组
	DataSourceIds []*string `json:"DataSourceIds,omitnil,omitempty" name:"DataSourceIds"`

	// 数据源名称数组
	DataSourceNames []*string `json:"DataSourceNames,omitnil,omitempty" name:"DataSourceNames"`

	// 数据源类型 database-自建数据源；cloud-integration-自定义数据源
	DataSourceType *string `json:"DataSourceType,omitnil,omitempty" name:"DataSourceType"`

	// 数据源模糊查询参数
	QueryOption *DataSourceQueryOption `json:"QueryOption,omitnil,omitempty" name:"QueryOption"`

	// 数据源视图Id数组
	ViewIds []*string `json:"ViewIds,omitnil,omitempty" name:"ViewIds"`

	// 查询未关联应用的数据源，0:未关联，该参数配合 AppIds 参数一块使用
	AppLinkStatus *int64 `json:"AppLinkStatus,omitnil,omitempty" name:"AppLinkStatus"`

	// 查询应用绑定数据源: 0: 否,1: 是
	QueryBindToApp *int64 `json:"QueryBindToApp,omitnil,omitempty" name:"QueryBindToApp"`

	// 查询连接器 0 数据模型 1 连接器 2 自定义连接器
	QueryConnector *int64 `json:"QueryConnector,omitnil,omitempty" name:"QueryConnector"`

	// 废弃中
	NotQuerySubTypeList []*string `json:"NotQuerySubTypeList,omitnil,omitempty" name:"NotQuerySubTypeList"`

	// 查询channelList
	ChannelList []*string `json:"ChannelList,omitnil,omitempty" name:"ChannelList"`

	// 是否查询数据源关联关系
	QueryDataSourceRelationList *bool `json:"QueryDataSourceRelationList,omitnil,omitempty" name:"QueryDataSourceRelationList"`

	// db实例类型
	DbInstanceType *string `json:"DbInstanceType,omitnil,omitempty" name:"DbInstanceType"`

	// 数据库表名列表
	DatabaseTableNames []*string `json:"DatabaseTableNames,omitnil,omitempty" name:"DatabaseTableNames"`

	// 是否查询系统模型，默认为true，需要显示设置为False才能过滤系统模型
	QuerySystemModel *bool `json:"QuerySystemModel,omitnil,omitempty" name:"QuerySystemModel"`
}

func (r *DescribeDataSourceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataSourceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageSize")
	delete(f, "PageIndex")
	delete(f, "EnvId")
	delete(f, "Appids")
	delete(f, "DataSourceIds")
	delete(f, "DataSourceNames")
	delete(f, "DataSourceType")
	delete(f, "QueryOption")
	delete(f, "ViewIds")
	delete(f, "AppLinkStatus")
	delete(f, "QueryBindToApp")
	delete(f, "QueryConnector")
	delete(f, "NotQuerySubTypeList")
	delete(f, "ChannelList")
	delete(f, "QueryDataSourceRelationList")
	delete(f, "DbInstanceType")
	delete(f, "DatabaseTableNames")
	delete(f, "QuerySystemModel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataSourceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataSourceListResponseParams struct {
	// data 数据
	Data *DataSourceDetailItems `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataSourceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataSourceListResponseParams `json:"Response"`
}

func (r *DescribeDataSourceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataSourceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKnowledgeDocumentSetDetailRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 知识库标识
	CollectionView *string `json:"CollectionView,omitnil,omitempty" name:"CollectionView"`

	// 文件名
	DocumentSetName *string `json:"DocumentSetName,omitnil,omitempty" name:"DocumentSetName"`

	// 文件id
	DocumentSetId *string `json:"DocumentSetId,omitnil,omitempty" name:"DocumentSetId"`
}

type DescribeKnowledgeDocumentSetDetailRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 知识库标识
	CollectionView *string `json:"CollectionView,omitnil,omitempty" name:"CollectionView"`

	// 文件名
	DocumentSetName *string `json:"DocumentSetName,omitnil,omitempty" name:"DocumentSetName"`

	// 文件id
	DocumentSetId *string `json:"DocumentSetId,omitnil,omitempty" name:"DocumentSetId"`
}

func (r *DescribeKnowledgeDocumentSetDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKnowledgeDocumentSetDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "CollectionView")
	delete(f, "DocumentSetName")
	delete(f, "DocumentSetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKnowledgeDocumentSetDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKnowledgeDocumentSetDetailResponseParams struct {
	// 新增文件返回信息
	Data *DescribeKnowledgeDocumentSetDetailRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeKnowledgeDocumentSetDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKnowledgeDocumentSetDetailResponseParams `json:"Response"`
}

func (r *DescribeKnowledgeDocumentSetDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKnowledgeDocumentSetDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKnowledgeDocumentSetDetailRsp struct {
	// 获取的数量。
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 文档信息
	DocumentSet *KnowledgeDocumentSet `json:"DocumentSet,omitnil,omitempty" name:"DocumentSet"`
}

// Predefined struct for user
type DescribeKnowledgeDocumentSetListRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 知识库标识
	CollectionView *string `json:"CollectionView,omitnil,omitempty" name:"CollectionView"`

	// 查询条件
	Query *PageQuery `json:"Query,omitnil,omitempty" name:"Query"`
}

type DescribeKnowledgeDocumentSetListRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 知识库标识
	CollectionView *string `json:"CollectionView,omitnil,omitempty" name:"CollectionView"`

	// 查询条件
	Query *PageQuery `json:"Query,omitnil,omitempty" name:"Query"`
}

func (r *DescribeKnowledgeDocumentSetListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKnowledgeDocumentSetListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "CollectionView")
	delete(f, "Query")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKnowledgeDocumentSetListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKnowledgeDocumentSetListResponseParams struct {
	// 新增文件返回信息
	Data *DescribeKnowledgeDocumentSetListRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeKnowledgeDocumentSetListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKnowledgeDocumentSetListResponseParams `json:"Response"`
}

func (r *DescribeKnowledgeDocumentSetListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKnowledgeDocumentSetListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKnowledgeDocumentSetListRsp struct {
	// 文件集
	DocumentSets []*QureyKnowledgeDocumentSet `json:"DocumentSets,omitnil,omitempty" name:"DocumentSets"`

	// 条数
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`
}

// Predefined struct for user
type DescribeKnowledgeSetListRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 知识库标识，精准查询
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 知识库名称，精准查询
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 分页起始位
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询条数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// NoPage标识不分页
	QueryMode *string `json:"QueryMode,omitnil,omitempty" name:"QueryMode"`
}

type DescribeKnowledgeSetListRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 知识库标识，精准查询
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 知识库名称，精准查询
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 分页起始位
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询条数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// NoPage标识不分页
	QueryMode *string `json:"QueryMode,omitnil,omitempty" name:"QueryMode"`
}

func (r *DescribeKnowledgeSetListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKnowledgeSetListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Name")
	delete(f, "Title")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "QueryMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKnowledgeSetListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKnowledgeSetListResponseParams struct {
	// 知识库列表
	Data *KnowledgeSetRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeKnowledgeSetListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKnowledgeSetListResponseParams `json:"Response"`
}

func (r *DescribeKnowledgeSetListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKnowledgeSetListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DocumentQuery struct {
	// 文件ids
	DocumentSetId []*string `json:"DocumentSetId,omitnil,omitempty" name:"DocumentSetId"`

	// 文件名集合
	DocumentSetName []*string `json:"DocumentSetName,omitnil,omitempty" name:"DocumentSetName"`

	// 使用创建 CollectionView 指定的 Filter 索引的字段设置查询过滤表达式
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type KnowledgeDocumentSet struct {
	// 文档id
	DocumentSetId *string `json:"DocumentSetId,omitnil,omitempty" name:"DocumentSetId"`

	// 文档名
	DocumentSetName *string `json:"DocumentSetName,omitnil,omitempty" name:"DocumentSetName"`

	// 文件完整内容。
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 文件内容前 200个字符。
	TextPrefix *string `json:"TextPrefix,omitnil,omitempty" name:"TextPrefix"`

	// 文件详情
	DocumentSetInfo *KnowledgeDocumentSetInfo `json:"DocumentSetInfo,omitnil,omitempty" name:"DocumentSetInfo"`

	// 文件拆分信息
	SplitterPreprocess *KnowledgeSplitterPreprocess `json:"SplitterPreprocess,omitnil,omitempty" name:"SplitterPreprocess"`

	// 未使用
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 文档标题
	FileTitle *string `json:"FileTitle,omitnil,omitempty" name:"FileTitle"`

	// 文档元信息，必须为jsonstring
	FileMetaData *string `json:"FileMetaData,omitnil,omitempty" name:"FileMetaData"`

	// 作者
	Author *string `json:"Author,omitnil,omitempty" name:"Author"`

	// 上传文件状态
	DocStatus *string `json:"DocStatus,omitnil,omitempty" name:"DocStatus"`

	// 文件上传失败的具体原因
	ErrMsg *string `json:"ErrMsg,omitnil,omitempty" name:"ErrMsg"`

	// Cos存储文件ID
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`
}

type KnowledgeDocumentSetInfo struct {
	// 文件的字符数。
	TextLength *uint64 `json:"TextLength,omitnil,omitempty" name:"TextLength"`

	// 文件的字节数。
	ByteLength *uint64 `json:"ByteLength,omitnil,omitempty" name:"ByteLength"`

	// 文件被预处理、Embedding 向量化的进度。
	IndexedProgress *uint64 `json:"IndexedProgress,omitnil,omitempty" name:"IndexedProgress"`

	// 文件预处理、Embedding 向量化的状态。
	// New：等待解析。
	// Loading：文件解析中。
	// Failure：文件解析、写入出错。
	// Ready：文件解析、写入完成。
	IndexedStatus *string `json:"IndexedStatus,omitnil,omitempty" name:"IndexedStatus"`

	// 文件创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 文件最后更新时间。
	LastUpdateTime *string `json:"LastUpdateTime,omitnil,omitempty" name:"LastUpdateTime"`

	// 文件关键字。
	Keywords *string `json:"Keywords,omitnil,omitempty" name:"Keywords"`
}

type KnowledgeSet struct {
	// 知识库标识
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 知识库名称
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 描述
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 状态，
	// NOT_ENABLED未启用
	// ENABLED 已启用
	Active *string `json:"Active,omitnil,omitempty" name:"Active"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 知识库的meta信息
	Meta *string `json:"Meta,omitnil,omitempty" name:"Meta"`

	// 知识库容量,单位字节
	TotalSize *string `json:"TotalSize,omitnil,omitempty" name:"TotalSize"`
}

type KnowledgeSetRsp struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 知识库列表
	KnowledgeSets []*KnowledgeSet `json:"KnowledgeSets,omitnil,omitempty" name:"KnowledgeSets"`
}

type KnowledgeSplitterPreprocess struct {
	// 在对文件拆分时，配置是否将 Title 追加到切分后的段落后面一并 Embedding。取值如下所示：
	// false：不追加。
	// true：将段落 Title 追加到切分后的段落。
	AppendTitleToChunk *bool `json:"AppendTitleToChunk,omitnil,omitempty" name:"AppendTitleToChunk"`

	// 在对文件拆分时，配置是否将关键字 keywords 追加到切分后的段落一并 Embedding。取值如下所示：
	// false：不追加。
	// true：将全文的 keywords 追加到切分后的段落。
	AppendKeywordsToChunk *bool `json:"AppendKeywordsToChunk,omitnil,omitempty" name:"AppendKeywordsToChunk"`
}

type PageQuery struct {
	// 文件id数组，表示要查询的文件的所有 ID，支持批量查询，数组元素范围[1,20]。
	DocumentSetId []*string `json:"DocumentSetId,omitnil,omitempty" name:"DocumentSetId"`

	// 表示要查询的文档名称，支持批量查询，数组元素范围[1,20]。
	DocumentSetName []*string `json:"DocumentSetName,omitnil,omitempty" name:"DocumentSetName"`

	// 取值范围：[1,16384]
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 设置分页偏移量，用于控制分页查询返回结果的起始位置，方便用户对数据进行分页展示和浏览。
	// 取值：为 limit 整数倍。
	// 计算公式：offset=limit*(page-1)。
	// 例如：当 limit = 10，page = 2 时，分页偏移量 offset = 10 * (2 - 1) = 10，表示从查询结果的第 11 条记录开始返回数据。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 设置后，其他字段为空值
	OutputFields []*string `json:"OutputFields,omitnil,omitempty" name:"OutputFields"`

	// 使用创建 CollectionView 指定的 Filter 索引的字段设置查询过滤表达式。
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type QureyKnowledgeDocumentSet struct {
	// 文件id
	DocumentSetId *string `json:"DocumentSetId,omitnil,omitempty" name:"DocumentSetId"`

	// 文件名
	DocumentSetName *string `json:"DocumentSetName,omitnil,omitempty" name:"DocumentSetName"`

	// 文件内容前 200个字符。
	TextPrefix *string `json:"TextPrefix,omitnil,omitempty" name:"TextPrefix"`

	// 文件拆分信息
	SplitterPreprocess *KnowledgeSplitterPreprocess `json:"SplitterPreprocess,omitnil,omitempty" name:"SplitterPreprocess"`

	// 文件详情
	DocumentSetInfo *QureyKnowledgeDocumentSetInfo `json:"DocumentSetInfo,omitnil,omitempty" name:"DocumentSetInfo"`

	// 文件标题
	FileTitle *string `json:"FileTitle,omitnil,omitempty" name:"FileTitle"`

	// 文件元信息，必须为jsonstring
	FileMetaData *string `json:"FileMetaData,omitnil,omitempty" name:"FileMetaData"`

	// name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 作者
	Author *string `json:"Author,omitnil,omitempty" name:"Author"`

	// 文档上传状态
	DocStatus *string `json:"DocStatus,omitnil,omitempty" name:"DocStatus"`

	// 上传文件失败时具体的错误消息
	ErrMsg *string `json:"ErrMsg,omitnil,omitempty" name:"ErrMsg"`

	// Cos存储文件ID
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`
}

type QureyKnowledgeDocumentSetInfo struct {
	// 文件的字符数。
	TextLength *uint64 `json:"TextLength,omitnil,omitempty" name:"TextLength"`

	// 文件的字节数。
	ByteLength *uint64 `json:"ByteLength,omitnil,omitempty" name:"ByteLength"`

	// 文件被预处理、Embedding 向量化的进度。
	IndexedProgress *uint64 `json:"IndexedProgress,omitnil,omitempty" name:"IndexedProgress"`

	// 文件预处理、Embedding 向量化的状态。
	// New：等待解析。
	// Loading：文件解析中。
	// Failure：文件解析、写入出错。
	// Ready：文件解析、写入完成。
	IndexedStatus *string `json:"IndexedStatus,omitnil,omitempty" name:"IndexedStatus"`

	// 错误信息
	IndexedErrorMsg *string `json:"IndexedErrorMsg,omitnil,omitempty" name:"IndexedErrorMsg"`

	// 文件创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 文件最后更新时间。
	LastUpdateTime *string `json:"LastUpdateTime,omitnil,omitempty" name:"LastUpdateTime"`

	// 文件关键字。
	Keywords *string `json:"Keywords,omitnil,omitempty" name:"Keywords"`
}

type RelationField struct {
	// 关联关系字段
	Field *string `json:"Field,omitnil,omitempty" name:"Field"`

	// 关联关系格式
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 关联数据源名称
	RelateDataSourceName *string `json:"RelateDataSourceName,omitnil,omitempty" name:"RelateDataSourceName"`
}

type SearchDocInfo struct {
	// 知识库名称
	CollectionViewName *string `json:"CollectionViewName,omitnil,omitempty" name:"CollectionViewName"`

	// 文档Id
	DocSetId *string `json:"DocSetId,omitnil,omitempty" name:"DocSetId"`

	// 文档Name
	DocSetName *string `json:"DocSetName,omitnil,omitempty" name:"DocSetName"`

	// 文档类型
	DocType *string `json:"DocType,omitnil,omitempty" name:"DocType"`

	// 文档标题
	FileTitle *string `json:"FileTitle,omitnil,omitempty" name:"FileTitle"`

	// 文档元信息
	FileMetaData *string `json:"FileMetaData,omitnil,omitempty" name:"FileMetaData"`

	// 文档描述
	DocDesc *string `json:"DocDesc,omitnil,omitempty" name:"DocDesc"`

	// 文档大小
	FileSize *uint64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// Cos存储文件ID
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`
}

// Predefined struct for user
type SearchDocListRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 知识库名称
	CollectionView *string `json:"CollectionView,omitnil,omitempty" name:"CollectionView"`

	// 搜索模式
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// 搜索值
	SearchValue *string `json:"SearchValue,omitnil,omitempty" name:"SearchValue"`

	// 页码
	PageNo *uint64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// 页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type SearchDocListRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 知识库名称
	CollectionView *string `json:"CollectionView,omitnil,omitempty" name:"CollectionView"`

	// 搜索模式
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// 搜索值
	SearchValue *string `json:"SearchValue,omitnil,omitempty" name:"SearchValue"`

	// 页码
	PageNo *uint64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// 页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *SearchDocListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchDocListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "CollectionView")
	delete(f, "SearchKey")
	delete(f, "SearchValue")
	delete(f, "PageNo")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchDocListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchDocListResponseParams struct {
	// 知识库文档搜索数据
	Data *SearchDocRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SearchDocListResponse struct {
	*tchttp.BaseResponse
	Response *SearchDocListResponseParams `json:"Response"`
}

func (r *SearchDocListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchDocListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchDocRsp struct {
	// 文档基本信息
	DocInfos []*SearchDocInfo `json:"DocInfos,omitnil,omitempty" name:"DocInfos"`

	// 文档总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`
}

type TicketAuthInfo struct {
	// 授权用户
	AuthUser *string `json:"AuthUser,omitnil,omitempty" name:"AuthUser"`
}

// Predefined struct for user
type UpdateKnowledgeSetRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 知识库标识
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 知识库名称
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 描述
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 状态;ENABLED启用；NOT_ENABLED不启用
	Active *string `json:"Active,omitnil,omitempty" name:"Active"`

	// 知识库的meta信息
	Meta *string `json:"Meta,omitnil,omitempty" name:"Meta"`
}

type UpdateKnowledgeSetRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 知识库标识
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 知识库名称
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 描述
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 状态;ENABLED启用；NOT_ENABLED不启用
	Active *string `json:"Active,omitnil,omitempty" name:"Active"`

	// 知识库的meta信息
	Meta *string `json:"Meta,omitnil,omitempty" name:"Meta"`
}

func (r *UpdateKnowledgeSetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateKnowledgeSetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Name")
	delete(f, "Title")
	delete(f, "Desc")
	delete(f, "Active")
	delete(f, "Meta")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateKnowledgeSetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateKnowledgeSetResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateKnowledgeSetResponse struct {
	*tchttp.BaseResponse
	Response *UpdateKnowledgeSetResponseParams `json:"Response"`
}

func (r *UpdateKnowledgeSetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateKnowledgeSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadKnowledgeDocumentSetRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 知识库标识
	CollectionView *string `json:"CollectionView,omitnil,omitempty" name:"CollectionView"`

	// 状态;ENABLED启用；NOT_ENABLED不启用
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 腾讯云文件存储位置的可读地址
	CosUrl *string `json:"CosUrl,omitnil,omitempty" name:"CosUrl"`

	// 文件类型，例如: .docx, .md
	DocumentType *string `json:"DocumentType,omitnil,omitempty" name:"DocumentType"`

	// 对文件的描述
	DocumentDesc *string `json:"DocumentDesc,omitnil,omitempty" name:"DocumentDesc"`

	// 文件标题
	FileTitle *string `json:"FileTitle,omitnil,omitempty" name:"FileTitle"`

	// 文件元信息，为jsonstring
	FileMetaData *string `json:"FileMetaData,omitnil,omitempty" name:"FileMetaData"`

	// 文件id
	DocumentSetId *string `json:"DocumentSetId,omitnil,omitempty" name:"DocumentSetId"`

	// 使用 regex 分割文档
	Delimiter *string `json:"Delimiter,omitnil,omitempty" name:"Delimiter"`

	// Cos存储文件ID
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`
}

type UploadKnowledgeDocumentSetRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 知识库标识
	CollectionView *string `json:"CollectionView,omitnil,omitempty" name:"CollectionView"`

	// 状态;ENABLED启用；NOT_ENABLED不启用
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 腾讯云文件存储位置的可读地址
	CosUrl *string `json:"CosUrl,omitnil,omitempty" name:"CosUrl"`

	// 文件类型，例如: .docx, .md
	DocumentType *string `json:"DocumentType,omitnil,omitempty" name:"DocumentType"`

	// 对文件的描述
	DocumentDesc *string `json:"DocumentDesc,omitnil,omitempty" name:"DocumentDesc"`

	// 文件标题
	FileTitle *string `json:"FileTitle,omitnil,omitempty" name:"FileTitle"`

	// 文件元信息，为jsonstring
	FileMetaData *string `json:"FileMetaData,omitnil,omitempty" name:"FileMetaData"`

	// 文件id
	DocumentSetId *string `json:"DocumentSetId,omitnil,omitempty" name:"DocumentSetId"`

	// 使用 regex 分割文档
	Delimiter *string `json:"Delimiter,omitnil,omitempty" name:"Delimiter"`

	// Cos存储文件ID
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`
}

func (r *UploadKnowledgeDocumentSetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadKnowledgeDocumentSetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "CollectionView")
	delete(f, "FileName")
	delete(f, "CosUrl")
	delete(f, "DocumentType")
	delete(f, "DocumentDesc")
	delete(f, "FileTitle")
	delete(f, "FileMetaData")
	delete(f, "DocumentSetId")
	delete(f, "Delimiter")
	delete(f, "FileId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadKnowledgeDocumentSetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadKnowledgeDocumentSetResponseParams struct {
	// 新增文件返回信息
	Data *UploadKnowledgeDocumentSetRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UploadKnowledgeDocumentSetResponse struct {
	*tchttp.BaseResponse
	Response *UploadKnowledgeDocumentSetResponseParams `json:"Response"`
}

func (r *UploadKnowledgeDocumentSetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadKnowledgeDocumentSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UploadKnowledgeDocumentSetRsp struct {
	// 给文件分配的 ID 信息。
	//
	// Deprecated: DocumentSetId is deprecated.
	DocumentSetId *string `json:"DocumentSetId,omitnil,omitempty" name:"DocumentSetId"`

	// 文件名
	DocumentSetName *string `json:"DocumentSetName,omitnil,omitempty" name:"DocumentSetName"`

	// 文件标题
	FileTitle *string `json:"FileTitle,omitnil,omitempty" name:"FileTitle"`

	// 文件元信息，为jsonstring
	FileMetaData *string `json:"FileMetaData,omitnil,omitempty" name:"FileMetaData"`

	// Cos存储文件ID
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`
}