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

package v20220105

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type ApiDatasourceConfig struct {
	// API数据源解析结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	FieldsJsonData *string `json:"FieldsJsonData,omitnil,omitempty" name:"FieldsJsonData"`

	// 连接类型1:直连 2:抽取
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConnectionType *uint64 `json:"ConnectionType,omitnil,omitempty" name:"ConnectionType"`

	// 抽取频率配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrequencyConfig *FrequencyConfig `json:"FrequencyConfig,omitnil,omitempty" name:"FrequencyConfig"`

	// 请求URL
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 1:GET 2:POST
	// 注意：此字段可能返回 null，表示取不到有效值。
	RequestMethod *uint64 `json:"RequestMethod,omitnil,omitempty" name:"RequestMethod"`

	// 请求头
	// 注意：此字段可能返回 null，表示取不到有效值。
	RequestHeader *string `json:"RequestHeader,omitnil,omitempty" name:"RequestHeader"`

	// 请求参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RequestParams *string `json:"RequestParams,omitnil,omitempty" name:"RequestParams"`

	// 请求体
	// 注意：此字段可能返回 null，表示取不到有效值。
	RequestBody *string `json:"RequestBody,omitnil,omitempty" name:"RequestBody"`

	// 用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 1: 无鉴权 2:BASIC_AUTH
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthorizationType *uint64 `json:"AuthorizationType,omitnil,omitempty" name:"AuthorizationType"`

	// 表id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableId *uint64 `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 路径DbName映射
	// 注意：此字段可能返回 null，表示取不到有效值。
	JsonPathDbNameMap *string `json:"JsonPathDbNameMap,omitnil,omitempty" name:"JsonPathDbNameMap"`

	// 鉴权API
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthApi *string `json:"AuthApi,omitnil,omitempty" name:"AuthApi"`

	// 应用Key
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// 应用密钥
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppSecret *string `json:"AppSecret,omitnil,omitempty" name:"AppSecret"`

	// 数据密钥Key
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretKey *string `json:"SecretKey,omitnil,omitempty" name:"SecretKey"`

	// 数据密钥初始化向量
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretIv *string `json:"SecretIv,omitnil,omitempty" name:"SecretIv"`
}

type ApiKeyAuthApplyVO struct {
	// <p>id</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>企业id</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CorpId *string `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// <p>apiKey</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiKey *string `json:"ApiKey,omitnil,omitempty" name:"ApiKey"`

	// <p>默认用户</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultUser *string `json:"DefaultUser,omitnil,omitempty" name:"DefaultUser"`

	// <p>创建人</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedUser *string `json:"CreatedUser,omitnil,omitempty" name:"CreatedUser"`

	// <p>创建时间</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// <p>更新人</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedUser *string `json:"UpdatedUser,omitnil,omitempty" name:"UpdatedUser"`

	// <p>更新时间</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`
}

type ApiKeyAuthApplyVOList struct {
	// <p>总数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// <p>页数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPages *int64 `json:"TotalPages,omitnil,omitempty" name:"TotalPages"`

	// <p>列表数据</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*ApiKeyAuthApplyVO `json:"List,omitnil,omitempty" name:"List"`
}

// Predefined struct for user
type ApplyEmbedIntervalRequestParams struct {
	// 分享项目id
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 分享页面id，嵌出看板时此为空值0，ChatBI嵌出时不传
	PageId *uint64 `json:"PageId,omitnil,omitempty" name:"PageId"`

	// 需要申请延期的Token
	BIToken *string `json:"BIToken,omitnil,omitempty" name:"BIToken"`

	// 备用字段
	ExtraParam *string `json:"ExtraParam,omitnil,omitempty" name:"ExtraParam"`

	// embed：页面/看板嵌出
	// chatBIEmbed：ChatBI嵌出
	Intention *string `json:"Intention,omitnil,omitempty" name:"Intention"`

	// panel, 看板；page，页面
	// project，ChatBI嵌出时
	Scope *string `json:"Scope,omitnil,omitempty" name:"Scope"`
}

type ApplyEmbedIntervalRequest struct {
	*tchttp.BaseRequest
	
	// 分享项目id
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 分享页面id，嵌出看板时此为空值0，ChatBI嵌出时不传
	PageId *uint64 `json:"PageId,omitnil,omitempty" name:"PageId"`

	// 需要申请延期的Token
	BIToken *string `json:"BIToken,omitnil,omitempty" name:"BIToken"`

	// 备用字段
	ExtraParam *string `json:"ExtraParam,omitnil,omitempty" name:"ExtraParam"`

	// embed：页面/看板嵌出
	// chatBIEmbed：ChatBI嵌出
	Intention *string `json:"Intention,omitnil,omitempty" name:"Intention"`

	// panel, 看板；page，页面
	// project，ChatBI嵌出时
	Scope *string `json:"Scope,omitnil,omitempty" name:"Scope"`
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
	delete(f, "Intention")
	delete(f, "Scope")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyEmbedIntervalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyEmbedIntervalResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// 额外参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 结果数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *ApplyEmbedTokenInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 结果描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`
}

type BaseStateAction struct {
	// 编辑是否可见
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShowEdit *bool `json:"ShowEdit,omitnil,omitempty" name:"ShowEdit"`

	// 编辑是否可点击
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsEdit *bool `json:"IsEdit,omitnil,omitempty" name:"IsEdit"`

	// 编辑按钮的文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	EditText *string `json:"EditText,omitnil,omitempty" name:"EditText"`

	// 编辑不可用时的提示文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	EditTips *string `json:"EditTips,omitnil,omitempty" name:"EditTips"`

	// 删除是否可见
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShowDel *bool `json:"ShowDel,omitnil,omitempty" name:"ShowDel"`

	// 删除是否可点击
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDel *bool `json:"IsDel,omitnil,omitempty" name:"IsDel"`

	// 删除按钮的文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	DelText *string `json:"DelText,omitnil,omitempty" name:"DelText"`

	// 删除不可用时的提示文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	DelTips *string `json:"DelTips,omitnil,omitempty" name:"DelTips"`

	// 复制是否可见
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShowCopy *bool `json:"ShowCopy,omitnil,omitempty" name:"ShowCopy"`

	// 是否可见
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShowView *bool `json:"ShowView,omitnil,omitempty" name:"ShowView"`

	// 是否可重命名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShowRename *bool `json:"ShowRename,omitnil,omitempty" name:"ShowRename"`
}

// Predefined struct for user
type ClearEmbedTokenRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 主账号id
	UserCorpId *string `json:"UserCorpId,omitnil,omitempty" name:"UserCorpId"`

	// panel , page
	Scope *string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// page id
	PageId *string `json:"PageId,omitnil,omitempty" name:"PageId"`
}

type ClearEmbedTokenRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 主账号id
	UserCorpId *string `json:"UserCorpId,omitnil,omitempty" name:"UserCorpId"`

	// panel , page
	Scope *string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// page id
	PageId *string `json:"PageId,omitnil,omitempty" name:"PageId"`
}

func (r *ClearEmbedTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ClearEmbedTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "UserCorpId")
	delete(f, "Scope")
	delete(f, "PageId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ClearEmbedTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ClearEmbedTokenResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// 额外消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 提示消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *bool `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ClearEmbedTokenResponse struct {
	*tchttp.BaseResponse
	Response *ClearEmbedTokenResponseParams `json:"Response"`
}

func (r *ClearEmbedTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ClearEmbedTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CorpUserListData struct {
	// 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*UserIdAndUserName `json:"List,omitnil,omitempty" name:"List"`

	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 页数
	TotalPages *int64 `json:"TotalPages,omitnil,omitempty" name:"TotalPages"`
}

// Predefined struct for user
type CreateAuthApiKeyRequestParams struct {
	// <p>默认用户</p>
	DefaultUser *string `json:"DefaultUser,omitnil,omitempty" name:"DefaultUser"`
}

type CreateAuthApiKeyRequest struct {
	*tchttp.BaseRequest
	
	// <p>默认用户</p>
	DefaultUser *string `json:"DefaultUser,omitnil,omitempty" name:"DefaultUser"`
}

func (r *CreateAuthApiKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAuthApiKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DefaultUser")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAuthApiKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAuthApiKeyResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// <p>&quot;&quot;</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// <p>&quot;success&quot;</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// <p>数据</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *ApiKeyAuthApplyVO `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAuthApiKeyResponse struct {
	*tchttp.BaseResponse
	Response *CreateAuthApiKeyResponseParams `json:"Response"`
}

func (r *CreateAuthApiKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAuthApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCorpTagRequestParams struct {
	// 标签名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type CreateCorpTagRequest struct {
	*tchttp.BaseRequest
	
	// 标签名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *CreateCorpTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCorpTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCorpTagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCorpTagResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// 扩展
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *DataId `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCorpTagResponse struct {
	*tchttp.BaseResponse
	Response *CreateCorpTagResponseParams `json:"Response"`
}

func (r *CreateCorpTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCorpTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDataTableRequestParams struct {
	// 后端提供字典：数据表类型，1、数据库建表，2、SQL建表，3、Excel上传，4、API接入，5、腾讯文档，6、云数据库，7、手工输入，8、关联查询
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 数据表名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 无
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 所属文件夹
	FoldId *uint64 `json:"FoldId,omitnil,omitempty" name:"FoldId"`

	// 数据源Id
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 物理表名
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// sql语句
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// excel地址
	ExcelUrl *string `json:"ExcelUrl,omitnil,omitempty" name:"ExcelUrl"`

	// 字段配置
	Fields []*TableField `json:"Fields,omitnil,omitempty" name:"Fields"`

	// 多表关联使用: 1:数据源原表,2:本地表,3:Excel表,4:API表
	TableNodeType *uint64 `json:"TableNodeType,omitnil,omitempty" name:"TableNodeType"`

	// 多表关联的原始表信息
	Tables []*JoinSourceTable `json:"Tables,omitnil,omitempty" name:"Tables"`

	// 多表关联的关联信息
	Joins []*JoinRelation `json:"Joins,omitnil,omitempty" name:"Joins"`

	// 补充信息，如api数据源信息，腾讯文档接入信息等
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`

	// 是否是异步
	AsyncRequest *bool `json:"AsyncRequest,omitnil,omitempty" name:"AsyncRequest"`

	// 依赖的异步事务id
	ParentTranId *string `json:"ParentTranId,omitnil,omitempty" name:"ParentTranId"`

	// API数据源配置
	ApiDatasourceConfig *ApiDatasourceConfig `json:"ApiDatasourceConfig,omitnil,omitempty" name:"ApiDatasourceConfig"`

	// 1
	ParamList []*ParamCreateDTO `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// dlc高级参数
	DlcExtInfo *string `json:"DlcExtInfo,omitnil,omitempty" name:"DlcExtInfo"`

	// 是否查询数据库
	QueryDbData *string `json:"QueryDbData,omitnil,omitempty" name:"QueryDbData"`

	// 数据表备注
	TableComment *string `json:"TableComment,omitnil,omitempty" name:"TableComment"`

	// 是否查询字段备注
	QueryFieldRemark *int64 `json:"QueryFieldRemark,omitnil,omitempty" name:"QueryFieldRemark"`

	// 字段备注列表
	FieldRemarkList []*FieldRemarkDTO `json:"FieldRemarkList,omitnil,omitempty" name:"FieldRemarkList"`
}

type CreateDataTableRequest struct {
	*tchttp.BaseRequest
	
	// 后端提供字典：数据表类型，1、数据库建表，2、SQL建表，3、Excel上传，4、API接入，5、腾讯文档，6、云数据库，7、手工输入，8、关联查询
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 数据表名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 无
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 所属文件夹
	FoldId *uint64 `json:"FoldId,omitnil,omitempty" name:"FoldId"`

	// 数据源Id
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 物理表名
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// sql语句
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// excel地址
	ExcelUrl *string `json:"ExcelUrl,omitnil,omitempty" name:"ExcelUrl"`

	// 字段配置
	Fields []*TableField `json:"Fields,omitnil,omitempty" name:"Fields"`

	// 多表关联使用: 1:数据源原表,2:本地表,3:Excel表,4:API表
	TableNodeType *uint64 `json:"TableNodeType,omitnil,omitempty" name:"TableNodeType"`

	// 多表关联的原始表信息
	Tables []*JoinSourceTable `json:"Tables,omitnil,omitempty" name:"Tables"`

	// 多表关联的关联信息
	Joins []*JoinRelation `json:"Joins,omitnil,omitempty" name:"Joins"`

	// 补充信息，如api数据源信息，腾讯文档接入信息等
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`

	// 是否是异步
	AsyncRequest *bool `json:"AsyncRequest,omitnil,omitempty" name:"AsyncRequest"`

	// 依赖的异步事务id
	ParentTranId *string `json:"ParentTranId,omitnil,omitempty" name:"ParentTranId"`

	// API数据源配置
	ApiDatasourceConfig *ApiDatasourceConfig `json:"ApiDatasourceConfig,omitnil,omitempty" name:"ApiDatasourceConfig"`

	// 1
	ParamList []*ParamCreateDTO `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// dlc高级参数
	DlcExtInfo *string `json:"DlcExtInfo,omitnil,omitempty" name:"DlcExtInfo"`

	// 是否查询数据库
	QueryDbData *string `json:"QueryDbData,omitnil,omitempty" name:"QueryDbData"`

	// 数据表备注
	TableComment *string `json:"TableComment,omitnil,omitempty" name:"TableComment"`

	// 是否查询字段备注
	QueryFieldRemark *int64 `json:"QueryFieldRemark,omitnil,omitempty" name:"QueryFieldRemark"`

	// 字段备注列表
	FieldRemarkList []*FieldRemarkDTO `json:"FieldRemarkList,omitnil,omitempty" name:"FieldRemarkList"`
}

func (r *CreateDataTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDataTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "Name")
	delete(f, "ProjectId")
	delete(f, "FoldId")
	delete(f, "DatasourceId")
	delete(f, "TableName")
	delete(f, "Sql")
	delete(f, "ExcelUrl")
	delete(f, "Fields")
	delete(f, "TableNodeType")
	delete(f, "Tables")
	delete(f, "Joins")
	delete(f, "ExtInfo")
	delete(f, "AsyncRequest")
	delete(f, "ParentTranId")
	delete(f, "ApiDatasourceConfig")
	delete(f, "ParamList")
	delete(f, "DlcExtInfo")
	delete(f, "QueryDbData")
	delete(f, "TableComment")
	delete(f, "QueryFieldRemark")
	delete(f, "FieldRemarkList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDataTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDataTableResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// 成功返回数据表的id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *IdDTO `json:"Data,omitnil,omitempty" name:"Data"`

	// 额外信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 错误提示
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDataTableResponse struct {
	*tchttp.BaseResponse
	Response *CreateDataTableResponseParams `json:"Response"`
}

func (r *CreateDataTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDataTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDatasourceCloudRequestParams struct {
	// <p>后端提供字典：域类型，1、腾讯云，2、本地</p>
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// <p>驱动</p><p>枚举值：</p><ul><li>MYSQL： MySQL数据库</li><li>PRESTO： PRESTO数据库</li><li>POSTGRE： PostgreSQL数据库</li><li>DLC： DLC数据库</li><li>MSSQL： 微软SQL Server数据库</li></ul>
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// <p>数据库编码</p>
	Charset *string `json:"Charset,omitnil,omitempty" name:"Charset"`

	// <p>用户名</p>
	DbUser *string `json:"DbUser,omitnil,omitempty" name:"DbUser"`

	// <p>密码</p>
	DbPwd *string `json:"DbPwd,omitnil,omitempty" name:"DbPwd"`

	// <p>数据库名称</p>
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// <p>数据库别名</p>
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// <p>项目ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>公有云内网ip</p>
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// <p>公有云内网端口</p>
	Vport *string `json:"Vport,omitnil,omitempty" name:"Vport"`

	// <p>vpc标识</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>统一vpc标识</p>
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// <p>区域标识（gz,bj)</p>
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// <p>扩展参数</p>
	ExtraParam *string `json:"ExtraParam,omitnil,omitempty" name:"ExtraParam"`

	// <p>实例Id</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>数据源产品名</p>
	ProdDbName *string `json:"ProdDbName,omitnil,omitempty" name:"ProdDbName"`

	// <p>第三方数据源标识</p>
	DataOrigin *string `json:"DataOrigin,omitnil,omitempty" name:"DataOrigin"`

	// <p>第三方项目id</p>
	DataOriginProjectId *string `json:"DataOriginProjectId,omitnil,omitempty" name:"DataOriginProjectId"`

	// <p>第三方数据源id</p>
	DataOriginDatasourceId *string `json:"DataOriginDatasourceId,omitnil,omitempty" name:"DataOriginDatasourceId"`

	// <p>集群id</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>数据库schema</p>
	Schema *string `json:"Schema,omitnil,omitempty" name:"Schema"`

	// <p>数据库版本</p>
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`
}

type CreateDatasourceCloudRequest struct {
	*tchttp.BaseRequest
	
	// <p>后端提供字典：域类型，1、腾讯云，2、本地</p>
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// <p>驱动</p><p>枚举值：</p><ul><li>MYSQL： MySQL数据库</li><li>PRESTO： PRESTO数据库</li><li>POSTGRE： PostgreSQL数据库</li><li>DLC： DLC数据库</li><li>MSSQL： 微软SQL Server数据库</li></ul>
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// <p>数据库编码</p>
	Charset *string `json:"Charset,omitnil,omitempty" name:"Charset"`

	// <p>用户名</p>
	DbUser *string `json:"DbUser,omitnil,omitempty" name:"DbUser"`

	// <p>密码</p>
	DbPwd *string `json:"DbPwd,omitnil,omitempty" name:"DbPwd"`

	// <p>数据库名称</p>
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// <p>数据库别名</p>
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// <p>项目ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>公有云内网ip</p>
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// <p>公有云内网端口</p>
	Vport *string `json:"Vport,omitnil,omitempty" name:"Vport"`

	// <p>vpc标识</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>统一vpc标识</p>
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// <p>区域标识（gz,bj)</p>
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// <p>扩展参数</p>
	ExtraParam *string `json:"ExtraParam,omitnil,omitempty" name:"ExtraParam"`

	// <p>实例Id</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>数据源产品名</p>
	ProdDbName *string `json:"ProdDbName,omitnil,omitempty" name:"ProdDbName"`

	// <p>第三方数据源标识</p>
	DataOrigin *string `json:"DataOrigin,omitnil,omitempty" name:"DataOrigin"`

	// <p>第三方项目id</p>
	DataOriginProjectId *string `json:"DataOriginProjectId,omitnil,omitempty" name:"DataOriginProjectId"`

	// <p>第三方数据源id</p>
	DataOriginDatasourceId *string `json:"DataOriginDatasourceId,omitnil,omitempty" name:"DataOriginDatasourceId"`

	// <p>集群id</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>数据库schema</p>
	Schema *string `json:"Schema,omitnil,omitempty" name:"Schema"`

	// <p>数据库版本</p>
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`
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
	delete(f, "ClusterId")
	delete(f, "Schema")
	delete(f, "DbVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDatasourceCloudRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDatasourceCloudResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// <p>成功无</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *IdDTO `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>额外信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// <p>提示</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// <p>HOST</p>
	DbHost *string `json:"DbHost,omitnil,omitempty" name:"DbHost"`

	// <p>端口</p>
	DbPort *uint64 `json:"DbPort,omitnil,omitempty" name:"DbPort"`

	// <p>后端提供字典：域类型，1、腾讯云，2、本地</p>
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// <p>驱动</p><p>枚举值：</p><ul><li>MYSQL： MySQL数据库</li><li>PRESTO： PRESTO数据库</li><li>POSTGRE： PostgreSQL数据库</li><li>DLC： DLC数据库</li><li>MSSQL： 微软SQL Server数据库</li></ul>
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// <p>数据库编码</p>
	Charset *string `json:"Charset,omitnil,omitempty" name:"Charset"`

	// <p>用户名</p>
	DbUser *string `json:"DbUser,omitnil,omitempty" name:"DbUser"`

	// <p>密码</p>
	DbPwd *string `json:"DbPwd,omitnil,omitempty" name:"DbPwd"`

	// <p>数据库名称</p>
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// <p>数据库别名</p>
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// <p>项目id</p>
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>catalog值</p>
	Catalog *string `json:"Catalog,omitnil,omitempty" name:"Catalog"`

	// <p>第三方数据源标识</p>
	DataOrigin *string `json:"DataOrigin,omitnil,omitempty" name:"DataOrigin"`

	// <p>第三方项目id</p>
	DataOriginProjectId *string `json:"DataOriginProjectId,omitnil,omitempty" name:"DataOriginProjectId"`

	// <p>第三方数据源id</p>
	DataOriginDatasourceId *string `json:"DataOriginDatasourceId,omitnil,omitempty" name:"DataOriginDatasourceId"`

	// <p>扩展参数</p>
	ExtraParam *string `json:"ExtraParam,omitnil,omitempty" name:"ExtraParam"`

	// <p>腾讯云私有网络统一标识</p>
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// <p>私有网络ip</p>
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// <p>私有网络端口</p>
	Vport *string `json:"Vport,omitnil,omitempty" name:"Vport"`

	// <p>腾讯云私有网络标识</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>操作权限限制</p>
	OperationAuthLimit []*string `json:"OperationAuthLimit,omitnil,omitempty" name:"OperationAuthLimit"`

	// <p>开启vpc</p>
	UseVPC *bool `json:"UseVPC,omitnil,omitempty" name:"UseVPC"`

	// <p>地域</p>
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// <p>数据库schema</p>
	Schema *string `json:"Schema,omitnil,omitempty" name:"Schema"`

	// <p>数据库版本</p>
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`
}

type CreateDatasourceRequest struct {
	*tchttp.BaseRequest
	
	// <p>HOST</p>
	DbHost *string `json:"DbHost,omitnil,omitempty" name:"DbHost"`

	// <p>端口</p>
	DbPort *uint64 `json:"DbPort,omitnil,omitempty" name:"DbPort"`

	// <p>后端提供字典：域类型，1、腾讯云，2、本地</p>
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// <p>驱动</p><p>枚举值：</p><ul><li>MYSQL： MySQL数据库</li><li>PRESTO： PRESTO数据库</li><li>POSTGRE： PostgreSQL数据库</li><li>DLC： DLC数据库</li><li>MSSQL： 微软SQL Server数据库</li></ul>
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// <p>数据库编码</p>
	Charset *string `json:"Charset,omitnil,omitempty" name:"Charset"`

	// <p>用户名</p>
	DbUser *string `json:"DbUser,omitnil,omitempty" name:"DbUser"`

	// <p>密码</p>
	DbPwd *string `json:"DbPwd,omitnil,omitempty" name:"DbPwd"`

	// <p>数据库名称</p>
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// <p>数据库别名</p>
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// <p>项目id</p>
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>catalog值</p>
	Catalog *string `json:"Catalog,omitnil,omitempty" name:"Catalog"`

	// <p>第三方数据源标识</p>
	DataOrigin *string `json:"DataOrigin,omitnil,omitempty" name:"DataOrigin"`

	// <p>第三方项目id</p>
	DataOriginProjectId *string `json:"DataOriginProjectId,omitnil,omitempty" name:"DataOriginProjectId"`

	// <p>第三方数据源id</p>
	DataOriginDatasourceId *string `json:"DataOriginDatasourceId,omitnil,omitempty" name:"DataOriginDatasourceId"`

	// <p>扩展参数</p>
	ExtraParam *string `json:"ExtraParam,omitnil,omitempty" name:"ExtraParam"`

	// <p>腾讯云私有网络统一标识</p>
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// <p>私有网络ip</p>
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// <p>私有网络端口</p>
	Vport *string `json:"Vport,omitnil,omitempty" name:"Vport"`

	// <p>腾讯云私有网络标识</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>操作权限限制</p>
	OperationAuthLimit []*string `json:"OperationAuthLimit,omitnil,omitempty" name:"OperationAuthLimit"`

	// <p>开启vpc</p>
	UseVPC *bool `json:"UseVPC,omitnil,omitempty" name:"UseVPC"`

	// <p>地域</p>
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// <p>数据库schema</p>
	Schema *string `json:"Schema,omitnil,omitempty" name:"Schema"`

	// <p>数据库版本</p>
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`
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
	delete(f, "OperationAuthLimit")
	delete(f, "UseVPC")
	delete(f, "RegionId")
	delete(f, "Schema")
	delete(f, "DbVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDatasourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDatasourceResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// <p>数据源id</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *IdDTO `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>额外信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// <p>提示</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 分享页面id，嵌出看板时此为空值0，ChatBI嵌出时不传
	PageId *uint64 `json:"PageId,omitnil,omitempty" name:"PageId"`

	// embed表示页面看板嵌出，chatBIEmbed表示ChatBI嵌出
	Intention *string `json:"Intention,omitnil,omitempty" name:"Intention"`

	// page表示嵌出页面，panel表示嵌出整个看板，ChatBI嵌出时使用project
	Scope *string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// 过期时间。 单位：分钟 最大值：240。即，4小时 默认值：240
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 备用字段
	ExtraParam *string `json:"ExtraParam,omitnil,omitempty" name:"ExtraParam"`

	// 使用者企业Id(仅用于多用户)
	UserCorpId *string `json:"UserCorpId,omitnil,omitempty" name:"UserCorpId"`

	// 使用者Id(仅用于多用户)
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 访问次数限制，限制范围1-99999，为空则不设置访问次数限制
	TicketNum *int64 `json:"TicketNum,omitnil,omitempty" name:"TicketNum"`

	// 全局筛选参数 报表过滤条件的全局参数。 格式为JSON格式的字符串
	// **目前仅支持字符类型页面参数绑定到全局参数
	// **
	// [
	//     {
	//         "ParamKey": "name",  //页面参数名称
	//         "JoinType": "AND",     // 连接方式,目前仅支持AND
	//         "WhereList": [
	//             {
	//                 "Operator": "-neq",   // 操作符，参考以下说明
	//                 "Value": [                   //操作值，单值数组只传一个值
	//                     "zZWJMD",
	//                     "ZzVGHX",
	//                     "湖南省",
	//                     "河北省"
	//                 ]
	//             }
	//         ]
	//     },
	//     {
	//         "ParamKey": "genderParam",
	//         "JoinType": "AND",
	//         "WhereList": [
	//             {
	//                 "Operator": "-neq",
	//                 "Value": [
	//                     "男"
	//                 ]
	//             }
	//         ]
	//     }
	// ]
	// 
	// 
	// 
	// Operator 目前支持
	// -neq  不等于!=操作符
	// -eq  等于=操作符
	// -is     in操作符
	GlobalParam *string `json:"GlobalParam,omitnil,omitempty" name:"GlobalParam"`

	// 100 不绑定用户, 一次创建一个token，UserCorpId和UserId 非必填，不支持 ChatBI 嵌出
	// 200 单用户单token , 一次创建一个token， UserCorpId和UserId 必填
	// 300 单用户多token, 一次创建多个token，UserCorpId和UserId 必填
	TokenType *int64 `json:"TokenType,omitnil,omitempty" name:"TokenType"`

	// 一次创建的token数
	TokenNum *int64 `json:"TokenNum,omitnil,omitempty" name:"TokenNum"`

	// 嵌出显示配置，目前为ChatBI嵌出场景用，TableFilter表示数据表列表过滤，SqlView表示sql查看功能
	ConfigParam *string `json:"ConfigParam,omitnil,omitempty" name:"ConfigParam"`
}

type CreateEmbedTokenRequest struct {
	*tchttp.BaseRequest
	
	// 分享项目id
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 分享页面id，嵌出看板时此为空值0，ChatBI嵌出时不传
	PageId *uint64 `json:"PageId,omitnil,omitempty" name:"PageId"`

	// embed表示页面看板嵌出，chatBIEmbed表示ChatBI嵌出
	Intention *string `json:"Intention,omitnil,omitempty" name:"Intention"`

	// page表示嵌出页面，panel表示嵌出整个看板，ChatBI嵌出时使用project
	Scope *string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// 过期时间。 单位：分钟 最大值：240。即，4小时 默认值：240
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 备用字段
	ExtraParam *string `json:"ExtraParam,omitnil,omitempty" name:"ExtraParam"`

	// 使用者企业Id(仅用于多用户)
	UserCorpId *string `json:"UserCorpId,omitnil,omitempty" name:"UserCorpId"`

	// 使用者Id(仅用于多用户)
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 访问次数限制，限制范围1-99999，为空则不设置访问次数限制
	TicketNum *int64 `json:"TicketNum,omitnil,omitempty" name:"TicketNum"`

	// 全局筛选参数 报表过滤条件的全局参数。 格式为JSON格式的字符串
	// **目前仅支持字符类型页面参数绑定到全局参数
	// **
	// [
	//     {
	//         "ParamKey": "name",  //页面参数名称
	//         "JoinType": "AND",     // 连接方式,目前仅支持AND
	//         "WhereList": [
	//             {
	//                 "Operator": "-neq",   // 操作符，参考以下说明
	//                 "Value": [                   //操作值，单值数组只传一个值
	//                     "zZWJMD",
	//                     "ZzVGHX",
	//                     "湖南省",
	//                     "河北省"
	//                 ]
	//             }
	//         ]
	//     },
	//     {
	//         "ParamKey": "genderParam",
	//         "JoinType": "AND",
	//         "WhereList": [
	//             {
	//                 "Operator": "-neq",
	//                 "Value": [
	//                     "男"
	//                 ]
	//             }
	//         ]
	//     }
	// ]
	// 
	// 
	// 
	// Operator 目前支持
	// -neq  不等于!=操作符
	// -eq  等于=操作符
	// -is     in操作符
	GlobalParam *string `json:"GlobalParam,omitnil,omitempty" name:"GlobalParam"`

	// 100 不绑定用户, 一次创建一个token，UserCorpId和UserId 非必填，不支持 ChatBI 嵌出
	// 200 单用户单token , 一次创建一个token， UserCorpId和UserId 必填
	// 300 单用户多token, 一次创建多个token，UserCorpId和UserId 必填
	TokenType *int64 `json:"TokenType,omitnil,omitempty" name:"TokenType"`

	// 一次创建的token数
	TokenNum *int64 `json:"TokenNum,omitnil,omitempty" name:"TokenNum"`

	// 嵌出显示配置，目前为ChatBI嵌出场景用，TableFilter表示数据表列表过滤，SqlView表示sql查看功能
	ConfigParam *string `json:"ConfigParam,omitnil,omitempty" name:"ConfigParam"`
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
	delete(f, "Intention")
	delete(f, "Scope")
	delete(f, "ExpireTime")
	delete(f, "ExtraParam")
	delete(f, "UserCorpId")
	delete(f, "UserId")
	delete(f, "TicketNum")
	delete(f, "GlobalParam")
	delete(f, "TokenType")
	delete(f, "TokenNum")
	delete(f, "ConfigParam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEmbedTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEmbedTokenResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// 额外信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *EmbedTokenInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 结果描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreatePermissionRanksRequestParams struct {
	// <p>页数</p>
	TableId *int64 `json:"TableId,omitnil,omitempty" name:"TableId"`

	// <p>模式</p><p>枚举值：</p><ul><li>ALL： 全部</li><li>Specify： 指定</li><li>TAG： 标签</li></ul><p>默认值：ALL</p>
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// <p>角色类型</p><p>枚举值：</p><ul><li>ROLES： 按角色</li><li>Others： 其它</li></ul><p>默认值：Others</p>
	RoleType *string `json:"RoleType,omitnil,omitempty" name:"RoleType"`

	// <p>所有页码</p>
	RoleId *int64 `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// <p>规则信息</p>
	RulerInfo *string `json:"RulerInfo,omitnil,omitempty" name:"RulerInfo"`

	// <p>类型</p><p>枚举值：</p><ul><li>ROW： 行权限</li><li>COLUMN： 列权限</li></ul><p>默认值：ROW</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>开启状态</p><p>枚举值：</p><ul><li>Open： 开启</li><li>Close： 关闭</li></ul><p>默认值：Close</p>
	OpenStatus *string `json:"OpenStatus,omitnil,omitempty" name:"OpenStatus"`

	// <p>项目id</p>
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>行列权限配置</p>
	RowColumnConfigList []*RowColumnConfig `json:"RowColumnConfigList,omitnil,omitempty" name:"RowColumnConfigList"`
}

type CreatePermissionRanksRequest struct {
	*tchttp.BaseRequest
	
	// <p>页数</p>
	TableId *int64 `json:"TableId,omitnil,omitempty" name:"TableId"`

	// <p>模式</p><p>枚举值：</p><ul><li>ALL： 全部</li><li>Specify： 指定</li><li>TAG： 标签</li></ul><p>默认值：ALL</p>
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// <p>角色类型</p><p>枚举值：</p><ul><li>ROLES： 按角色</li><li>Others： 其它</li></ul><p>默认值：Others</p>
	RoleType *string `json:"RoleType,omitnil,omitempty" name:"RoleType"`

	// <p>所有页码</p>
	RoleId *int64 `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// <p>规则信息</p>
	RulerInfo *string `json:"RulerInfo,omitnil,omitempty" name:"RulerInfo"`

	// <p>类型</p><p>枚举值：</p><ul><li>ROW： 行权限</li><li>COLUMN： 列权限</li></ul><p>默认值：ROW</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>开启状态</p><p>枚举值：</p><ul><li>Open： 开启</li><li>Close： 关闭</li></ul><p>默认值：Close</p>
	OpenStatus *string `json:"OpenStatus,omitnil,omitempty" name:"OpenStatus"`

	// <p>项目id</p>
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>行列权限配置</p>
	RowColumnConfigList []*RowColumnConfig `json:"RowColumnConfigList,omitnil,omitempty" name:"RowColumnConfigList"`
}

func (r *CreatePermissionRanksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePermissionRanksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TableId")
	delete(f, "Mode")
	delete(f, "RoleType")
	delete(f, "RoleId")
	delete(f, "RulerInfo")
	delete(f, "Type")
	delete(f, "OpenStatus")
	delete(f, "ProjectId")
	delete(f, "RowColumnConfigList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePermissionRanksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePermissionRanksResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// <p>消息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// <p>112</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// <p>1</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePermissionRanksResponse struct {
	*tchttp.BaseResponse
	Response *CreatePermissionRanksResponseParams `json:"Response"`
}

func (r *CreatePermissionRanksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePermissionRanksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProjectRequestParams struct {
	// <p>项目名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>logo底色</p>
	ColorCode *string `json:"ColorCode,omitnil,omitempty" name:"ColorCode"`

	// <p>项目Logo</p>
	Logo *string `json:"Logo,omitnil,omitempty" name:"Logo"`

	// <p>备注</p>
	Mark *string `json:"Mark,omitnil,omitempty" name:"Mark"`

	// <p>是否允许用户申请</p>
	IsApply *bool `json:"IsApply,omitnil,omitempty" name:"IsApply"`

	// <p>默认看板</p><p>枚举值：</p><ul><li>1： 项目看板</li><li>2： 我的看板</li></ul>
	DefaultPanelType *int64 `json:"DefaultPanelType,omitnil,omitempty" name:"DefaultPanelType"`

	// <p>管理平台</p>
	ManagePlatform *string `json:"ManagePlatform,omitnil,omitempty" name:"ManagePlatform"`
}

type CreateProjectRequest struct {
	*tchttp.BaseRequest
	
	// <p>项目名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>logo底色</p>
	ColorCode *string `json:"ColorCode,omitnil,omitempty" name:"ColorCode"`

	// <p>项目Logo</p>
	Logo *string `json:"Logo,omitnil,omitempty" name:"Logo"`

	// <p>备注</p>
	Mark *string `json:"Mark,omitnil,omitempty" name:"Mark"`

	// <p>是否允许用户申请</p>
	IsApply *bool `json:"IsApply,omitnil,omitempty" name:"IsApply"`

	// <p>默认看板</p><p>枚举值：</p><ul><li>1： 项目看板</li><li>2： 我的看板</li></ul>
	DefaultPanelType *int64 `json:"DefaultPanelType,omitnil,omitempty" name:"DefaultPanelType"`

	// <p>管理平台</p>
	ManagePlatform *string `json:"ManagePlatform,omitnil,omitempty" name:"ManagePlatform"`
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
	delete(f, "ManagePlatform")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProjectResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// <p>额外数据</p>
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// <p>数据</p>
	Data *Data `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>返回信息</p>
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

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
type CreateTagTableRequestParams struct {
	// 标签表名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 标签表关联的项目id
	AutoImportProjectId *int64 `json:"AutoImportProjectId,omitnil,omitempty" name:"AutoImportProjectId"`

	// 标签表关联的数据表id
	AutoImportTableId *int64 `json:"AutoImportTableId,omitnil,omitempty" name:"AutoImportTableId"`

	// uin对应字段
	AutoImportUinField *string `json:"AutoImportUinField,omitnil,omitempty" name:"AutoImportUinField"`
}

type CreateTagTableRequest struct {
	*tchttp.BaseRequest
	
	// 标签表名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 标签表关联的项目id
	AutoImportProjectId *int64 `json:"AutoImportProjectId,omitnil,omitempty" name:"AutoImportProjectId"`

	// 标签表关联的数据表id
	AutoImportTableId *int64 `json:"AutoImportTableId,omitnil,omitempty" name:"AutoImportTableId"`

	// uin对应字段
	AutoImportUinField *string `json:"AutoImportUinField,omitnil,omitempty" name:"AutoImportUinField"`
}

func (r *CreateTagTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTagTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "AutoImportProjectId")
	delete(f, "AutoImportTableId")
	delete(f, "AutoImportUinField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTagTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTagTableResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// 数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *CreateTagTableVO `json:"Data,omitnil,omitempty" name:"Data"`

	// 扩展
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTagTableResponse struct {
	*tchttp.BaseResponse
	Response *CreateTagTableResponseParams `json:"Response"`
}

func (r *CreateTagTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTagTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTagTableVO struct {
	// 标签表id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`
}

// Predefined struct for user
type CreateUserGroupMemberRequestParams struct {
	// <p>用户组id</p>
	GroupId *uint64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// <p>用户id集合</p>
	UserIdList []*string `json:"UserIdList,omitnil,omitempty" name:"UserIdList"`
}

type CreateUserGroupMemberRequest struct {
	*tchttp.BaseRequest
	
	// <p>用户组id</p>
	GroupId *uint64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// <p>用户id集合</p>
	UserIdList []*string `json:"UserIdList,omitnil,omitempty" name:"UserIdList"`
}

func (r *CreateUserGroupMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserGroupMemberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "UserIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserGroupMemberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserGroupMemberResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// <p>额外信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// <p>结果信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// <p>数据</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateUserGroupMemberResponse struct {
	*tchttp.BaseResponse
	Response *CreateUserGroupMemberResponseParams `json:"Response"`
}

func (r *CreateUserGroupMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserGroupMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserGroupRequestParams struct {
	// <p>组管理员</p>
	AdminUserId *string `json:"AdminUserId,omitnil,omitempty" name:"AdminUserId"`

	// <p>描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>用户组名称</p>
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// <p>位置</p>
	Location *int64 `json:"Location,omitnil,omitempty" name:"Location"`

	// <p>父用户组id</p>
	ParentId *int64 `json:"ParentId,omitnil,omitempty" name:"ParentId"`
}

type CreateUserGroupRequest struct {
	*tchttp.BaseRequest
	
	// <p>组管理员</p>
	AdminUserId *string `json:"AdminUserId,omitnil,omitempty" name:"AdminUserId"`

	// <p>描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>用户组名称</p>
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// <p>位置</p>
	Location *int64 `json:"Location,omitnil,omitempty" name:"Location"`

	// <p>父用户组id</p>
	ParentId *int64 `json:"ParentId,omitnil,omitempty" name:"ParentId"`
}

func (r *CreateUserGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AdminUserId")
	delete(f, "Description")
	delete(f, "GroupName")
	delete(f, "Location")
	delete(f, "ParentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserGroupResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// <p>额外信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// <p>结果信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// <p>数据</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *UserGroupVO `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateUserGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateUserGroupResponseParams `json:"Response"`
}

func (r *CreateUserGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserRoleProjectRequestParams struct {
	// <p>项目ID</p>
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>角色ID列表</p>
	RoleIdList []*int64 `json:"RoleIdList,omitnil,omitempty" name:"RoleIdList"`

	// <p>用户列表（废弃）</p>
	//
	// Deprecated: UserList is deprecated.
	UserList []*UserIdAndUserName `json:"UserList,omitnil,omitempty" name:"UserList"`

	// <p>用户列表（新）</p>
	UserInfoList []*UserInfo `json:"UserInfoList,omitnil,omitempty" name:"UserInfoList"`
}

type CreateUserRoleProjectRequest struct {
	*tchttp.BaseRequest
	
	// <p>项目ID</p>
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>角色ID列表</p>
	RoleIdList []*int64 `json:"RoleIdList,omitnil,omitempty" name:"RoleIdList"`

	// <p>用户列表（废弃）</p>
	UserList []*UserIdAndUserName `json:"UserList,omitnil,omitempty" name:"UserList"`

	// <p>用户列表（新）</p>
	UserInfoList []*UserInfo `json:"UserInfoList,omitnil,omitempty" name:"UserInfoList"`
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
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// <p>扩展</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// <p>数据</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *DataId `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>消息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	RoleIdList []*int64 `json:"RoleIdList,omitnil,omitempty" name:"RoleIdList"`

	// 用户列表（废弃）
	//
	// Deprecated: UserList is deprecated.
	UserList []*UserIdAndUserName `json:"UserList,omitnil,omitempty" name:"UserList"`

	// 用户列表（新）
	UserInfoList []*UserInfo `json:"UserInfoList,omitnil,omitempty" name:"UserInfoList"`

	// 用户组id列表
	UserGroups []*uint64 `json:"UserGroups,omitnil,omitempty" name:"UserGroups"`
}

type CreateUserRoleRequest struct {
	*tchttp.BaseRequest
	
	// 角色ID列表
	RoleIdList []*int64 `json:"RoleIdList,omitnil,omitempty" name:"RoleIdList"`

	// 用户列表（废弃）
	UserList []*UserIdAndUserName `json:"UserList,omitnil,omitempty" name:"UserList"`

	// 用户列表（新）
	UserInfoList []*UserInfo `json:"UserInfoList,omitnil,omitempty" name:"UserInfoList"`

	// 用户组id列表
	UserGroups []*uint64 `json:"UserGroups,omitnil,omitempty" name:"UserGroups"`
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
	delete(f, "UserGroups")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserRoleResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// 扩展
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *DataId `json:"Data,omitnil,omitempty" name:"Data"`

	// 消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// url
	// 注意：此字段可能返回 null，表示取不到有效值。
	EditUrl *string `json:"EditUrl,omitnil,omitempty" name:"EditUrl"`
}

type DataId struct {
	// 数据id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`
}

type DatasourceInfo struct {
	// 数据库ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 数据库名称
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 域类型，1、腾讯云，2、本地
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// 数据库别名
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// 数据库驱动
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// HOST
	DbHost *string `json:"DbHost,omitnil,omitempty" name:"DbHost"`

	// 端口
	DbPort *uint64 `json:"DbPort,omitnil,omitempty" name:"DbPort"`

	// 用户名
	DbUser *string `json:"DbUser,omitnil,omitempty" name:"DbUser"`

	// 数据库编码
	Charset *string `json:"Charset,omitnil,omitempty" name:"Charset"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 修改时间
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// 创建人
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedUser *string `json:"CreatedUser,omitnil,omitempty" name:"CreatedUser"`

	// catalog值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Catalog *string `json:"Catalog,omitnil,omitempty" name:"Catalog"`

	// 连接类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConnectType *string `json:"ConnectType,omitnil,omitempty" name:"ConnectType"`

	// 项目id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 数据源描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 数据源状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 来源平台
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourcePlat *string `json:"SourcePlat,omitnil,omitempty" name:"SourcePlat"`

	// 扩展参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraParam *string `json:"ExtraParam,omitnil,omitempty" name:"ExtraParam"`

	// 额外信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddInfo *string `json:"AddInfo,omitnil,omitempty" name:"AddInfo"`

	// 项目名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 引擎类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// 数据源负责人
	// 注意：此字段可能返回 null，表示取不到有效值。
	Manager *string `json:"Manager,omitnil,omitempty" name:"Manager"`

	// 特定操作人员白名单
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperatorWhitelist *string `json:"OperatorWhitelist,omitnil,omitempty" name:"OperatorWhitelist"`

	// 数据源的vpc信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 数据源的uniqVpc信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 数据源的地域信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 操作属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	StateAction *BaseStateAction `json:"StateAction,omitnil,omitempty" name:"StateAction"`

	// 更新人
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedUser *string `json:"UpdatedUser,omitnil,omitempty" name:"UpdatedUser"`

	// 权限列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PermissionList []*PermissionGroup `json:"PermissionList,omitnil,omitempty" name:"PermissionList"`

	// 权限值列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthList []*string `json:"AuthList,omitnil,omitempty" name:"AuthList"`

	// 第三方数据源标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataOrigin *string `json:"DataOrigin,omitnil,omitempty" name:"DataOrigin"`

	// 第三方项目id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataOriginProjectId *string `json:"DataOriginProjectId,omitnil,omitempty" name:"DataOriginProjectId"`

	// 第三方数据源id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataOriginDatasourceId *string `json:"DataOriginDatasourceId,omitnil,omitempty" name:"DataOriginDatasourceId"`

	// 集群id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 数据源名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbTypeName *string `json:"DbTypeName,omitnil,omitempty" name:"DbTypeName"`

	// 开启vpc
	// 注意：此字段可能返回 null，表示取不到有效值。
	UseVPC *bool `json:"UseVPC,omitnil,omitempty" name:"UseVPC"`

	// 所属人ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Owner *string `json:"Owner,omitnil,omitempty" name:"Owner"`

	// 所属人名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerName *string `json:"OwnerName,omitnil,omitempty" name:"OwnerName"`

	// 数据库schema
	// 注意：此字段可能返回 null，表示取不到有效值。
	Schema *string `json:"Schema,omitnil,omitempty" name:"Schema"`

	// 数据库版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`
}

type DatasourceInfoData struct {
	// 数据源详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*DatasourceInfo `json:"List,omitnil,omitempty" name:"List"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 总页数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPages *int64 `json:"TotalPages,omitnil,omitempty" name:"TotalPages"`
}

// Predefined struct for user
type DeleteAuthApiKeyRequestParams struct {
	// <p>ApiKey</p>
	ApiKey *string `json:"ApiKey,omitnil,omitempty" name:"ApiKey"`
}

type DeleteAuthApiKeyRequest struct {
	*tchttp.BaseRequest
	
	// <p>ApiKey</p>
	ApiKey *string `json:"ApiKey,omitnil,omitempty" name:"ApiKey"`
}

func (r *DeleteAuthApiKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAuthApiKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApiKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAuthApiKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAuthApiKeyResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// <p>&quot;&quot;</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// <p>&quot;success&quot;</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// <p>数据</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAuthApiKeyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAuthApiKeyResponseParams `json:"Response"`
}

func (r *DeleteAuthApiKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAuthApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDatasourceRequestParams struct {
	// 数据源id
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 项目id
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DeleteDatasourceRequest struct {
	*tchttp.BaseRequest
	
	// 数据源id
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 项目id
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
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
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// 数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 扩展
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 信息
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// <p>项目ID</p>
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>随机数</p>
	Seed *string `json:"Seed,omitnil,omitempty" name:"Seed"`

	// <p>默认看板</p><p>枚举值：</p><ul><li>1： 项目看板</li><li>2： 我的看板</li></ul>
	DefaultPanelType *int64 `json:"DefaultPanelType,omitnil,omitempty" name:"DefaultPanelType"`
}

type DeleteProjectRequest struct {
	*tchttp.BaseRequest
	
	// <p>项目ID</p>
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>随机数</p>
	Seed *string `json:"Seed,omitnil,omitempty" name:"Seed"`

	// <p>默认看板</p><p>枚举值：</p><ul><li>1： 项目看板</li><li>2： 我的看板</li></ul>
	DefaultPanelType *int64 `json:"DefaultPanelType,omitnil,omitempty" name:"DefaultPanelType"`
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
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// <p>”“</p>
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// <p>&quot;&quot;</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>&quot;&quot;</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

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

// Predefined struct for user
type DeleteUserGroupMemberRequestParams struct {
	// <p>用户组id</p>
	GroupId *uint64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// <p>用户id集合</p>
	UserIdList []*string `json:"UserIdList,omitnil,omitempty" name:"UserIdList"`
}

type DeleteUserGroupMemberRequest struct {
	*tchttp.BaseRequest
	
	// <p>用户组id</p>
	GroupId *uint64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// <p>用户id集合</p>
	UserIdList []*string `json:"UserIdList,omitnil,omitempty" name:"UserIdList"`
}

func (r *DeleteUserGroupMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserGroupMemberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "UserIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUserGroupMemberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserGroupMemberResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// <p>额外信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// <p>结果信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// <p>数据</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteUserGroupMemberResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUserGroupMemberResponseParams `json:"Response"`
}

func (r *DeleteUserGroupMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserGroupMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserGroupRequestParams struct {
	// <p>用户组id</p>
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`
}

type DeleteUserGroupRequest struct {
	*tchttp.BaseRequest
	
	// <p>用户组id</p>
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`
}

func (r *DeleteUserGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUserGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserGroupResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// <p>额外信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// <p>结果信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// <p>数据</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteUserGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUserGroupResponseParams `json:"Response"`
}

func (r *DeleteUserGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserRoleProjectRequestParams struct {
	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type DeleteUserRoleProjectRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
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
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// 扩展
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type DeleteUserRoleRequest struct {
	*tchttp.BaseRequest
	
	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
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
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// 扩展
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeAuthApiKeyInfoRequestParams struct {
	// <p>ApiKey</p>
	ApiKey *string `json:"ApiKey,omitnil,omitempty" name:"ApiKey"`
}

type DescribeAuthApiKeyInfoRequest struct {
	*tchttp.BaseRequest
	
	// <p>ApiKey</p>
	ApiKey *string `json:"ApiKey,omitnil,omitempty" name:"ApiKey"`
}

func (r *DescribeAuthApiKeyInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuthApiKeyInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApiKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuthApiKeyInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuthApiKeyInfoResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// <p>&quot;&quot;</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// <p>&quot;success&quot;</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// <p>数据</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *ApiKeyAuthApplyVO `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAuthApiKeyInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuthApiKeyInfoResponseParams `json:"Response"`
}

func (r *DescribeAuthApiKeyInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuthApiKeyInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuthApiKeyListRequestParams struct {
	// <p>全部</p><p>默认值：false</p>
	AllPage *bool `json:"AllPage,omitnil,omitempty" name:"AllPage"`

	// <p>页码</p><p>默认值：0</p>
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// <p>分页大小</p><p>默认值：10</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>关键字过滤</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

type DescribeAuthApiKeyListRequest struct {
	*tchttp.BaseRequest
	
	// <p>全部</p><p>默认值：false</p>
	AllPage *bool `json:"AllPage,omitnil,omitempty" name:"AllPage"`

	// <p>页码</p><p>默认值：0</p>
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// <p>分页大小</p><p>默认值：10</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>关键字过滤</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

func (r *DescribeAuthApiKeyListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuthApiKeyListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AllPage")
	delete(f, "PageNo")
	delete(f, "PageSize")
	delete(f, "Keyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuthApiKeyListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuthApiKeyListResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// <p>{}</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// <p>信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// <p>数据</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *ApiKeyAuthApplyVOList `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAuthApiKeyListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuthApiKeyListResponseParams `json:"Response"`
}

func (r *DescribeAuthApiKeyListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuthApiKeyListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatasourceListRequestParams struct {
	// 无
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 返回所有页面，默认false
	AllPage *bool `json:"AllPage,omitnil,omitempty" name:"AllPage"`

	// 数据库名称检索
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 无
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// 无
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 搜索关键词
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 过滤无权限列表的参数（0 全量，1 使用权限，2 编辑权限）
	PermissionType *int64 `json:"PermissionType,omitnil,omitempty" name:"PermissionType"`
}

type DescribeDatasourceListRequest struct {
	*tchttp.BaseRequest
	
	// 无
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 返回所有页面，默认false
	AllPage *bool `json:"AllPage,omitnil,omitempty" name:"AllPage"`

	// 数据库名称检索
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 无
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// 无
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 搜索关键词
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 过滤无权限列表的参数（0 全量，1 使用权限，2 编辑权限）
	PermissionType *int64 `json:"PermissionType,omitnil,omitempty" name:"PermissionType"`
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
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// 列表详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *DatasourceInfoData `json:"Data,omitnil,omitempty" name:"Data"`

	// 信息
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 信息
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribePageWidgetListRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 页面id
	PageId *string `json:"PageId,omitnil,omitempty" name:"PageId"`
}

type DescribePageWidgetListRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 页面id
	PageId *string `json:"PageId,omitnil,omitempty" name:"PageId"`
}

func (r *DescribePageWidgetListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePageWidgetListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "PageId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePageWidgetListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePageWidgetListResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// 扩展参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 返回数据结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *WidgetListVO `json:"Data,omitnil,omitempty" name:"Data"`

	// 返回消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePageWidgetListResponse struct {
	*tchttp.BaseResponse
	Response *DescribePageWidgetListResponseParams `json:"Response"`
}

func (r *DescribePageWidgetListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePageWidgetListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePermissionRanksInfoRequestParams struct {
	// <p>页数</p>
	TableId *int64 `json:"TableId,omitnil,omitempty" name:"TableId"`

	// <p>模式</p><p>枚举值：</p><ul><li>ALL： 全部</li><li>Specify： 指定</li><li>TAG： 标签</li></ul><p>默认值：ALL</p>
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// <p>角色类型</p><p>枚举值：</p><ul><li>ROLES： 按角色</li><li>Others： 其它</li></ul><p>默认值：Others</p>
	RoleType *string `json:"RoleType,omitnil,omitempty" name:"RoleType"`

	// <p>所有页码</p>
	RoleId *int64 `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// <p>类型</p><p>枚举值：</p><ul><li>ROW： 行权限</li><li>COLUMN： 列权限</li></ul><p>默认值：ROW</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>项目id</p>
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DescribePermissionRanksInfoRequest struct {
	*tchttp.BaseRequest
	
	// <p>页数</p>
	TableId *int64 `json:"TableId,omitnil,omitempty" name:"TableId"`

	// <p>模式</p><p>枚举值：</p><ul><li>ALL： 全部</li><li>Specify： 指定</li><li>TAG： 标签</li></ul><p>默认值：ALL</p>
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// <p>角色类型</p><p>枚举值：</p><ul><li>ROLES： 按角色</li><li>Others： 其它</li></ul><p>默认值：Others</p>
	RoleType *string `json:"RoleType,omitnil,omitempty" name:"RoleType"`

	// <p>所有页码</p>
	RoleId *int64 `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// <p>类型</p><p>枚举值：</p><ul><li>ROW： 行权限</li><li>COLUMN： 列权限</li></ul><p>默认值：ROW</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>项目id</p>
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *DescribePermissionRanksInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePermissionRanksInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TableId")
	delete(f, "Mode")
	delete(f, "RoleType")
	delete(f, "RoleId")
	delete(f, "Type")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePermissionRanksInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePermissionRanksInfoResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// <p>消息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// <p>112</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// <p>无</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *RankInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePermissionRanksInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribePermissionRanksInfoResponseParams `json:"Response"`
}

func (r *DescribePermissionRanksInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePermissionRanksInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePermissionRoleInfoRequestParams struct {
	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 页数 
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// 条数
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 所有页码
	AllPage *bool `json:"AllPage,omitnil,omitempty" name:"AllPage"`
}

type DescribePermissionRoleInfoRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 页数 
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// 条数
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 所有页码
	AllPage *bool `json:"AllPage,omitnil,omitempty" name:"AllPage"`
}

func (r *DescribePermissionRoleInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePermissionRoleInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "PageNo")
	delete(f, "PageSize")
	delete(f, "AllPage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePermissionRoleInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePermissionRoleInfoResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// 数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*Role `json:"Data,omitnil,omitempty" name:"Data"`

	// 消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 112 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePermissionRoleInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribePermissionRoleInfoResponseParams `json:"Response"`
}

func (r *DescribePermissionRoleInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePermissionRoleInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePermissionStatusInfoRequestParams struct {
	// <p>页数</p>
	TableId *int64 `json:"TableId,omitnil,omitempty" name:"TableId"`

	// <p>类型</p><p>枚举值：</p><ul><li>ROW： 行权限</li><li>COLUMN： 列权限</li></ul><p>默认值：ROW</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>1</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DescribePermissionStatusInfoRequest struct {
	*tchttp.BaseRequest
	
	// <p>页数</p>
	TableId *int64 `json:"TableId,omitnil,omitempty" name:"TableId"`

	// <p>类型</p><p>枚举值：</p><ul><li>ROW： 行权限</li><li>COLUMN： 列权限</li></ul><p>默认值：ROW</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>1</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *DescribePermissionStatusInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePermissionStatusInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TableId")
	delete(f, "Type")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePermissionStatusInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePermissionStatusInfoResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// <p>消息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// <p>112</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// <p>1</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *RowColumnStatus `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePermissionStatusInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribePermissionStatusInfoResponseParams `json:"Response"`
}

func (r *DescribePermissionStatusInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePermissionStatusInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectInfoRequestParams struct {
	// <p>项目Id</p>
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>默认看板</p><p>枚举值：</p><ul><li>1： 项目看板</li><li>2： 我的看板</li></ul>
	DefaultPanelType *int64 `json:"DefaultPanelType,omitnil,omitempty" name:"DefaultPanelType"`
}

type DescribeProjectInfoRequest struct {
	*tchttp.BaseRequest
	
	// <p>项目Id</p>
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>默认看板</p><p>枚举值：</p><ul><li>1： 项目看板</li><li>2： 我的看板</li></ul>
	DefaultPanelType *int64 `json:"DefaultPanelType,omitnil,omitempty" name:"DefaultPanelType"`
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
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// <p>&quot;&quot;</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// <p>&quot;&quot;</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// <p>项目详情</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *Project `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页标
	PageNo *uint64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// 检索模糊字段
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 是否全部展示，如果是ture，则忽略分页
	AllPage *bool `json:"AllPage,omitnil,omitempty" name:"AllPage"`

	// 角色信息
	ModuleCollection *string `json:"ModuleCollection,omitnil,omitempty" name:"ModuleCollection"`

	// moduleId集合
	ModuleIdList []*string `json:"ModuleIdList,omitnil,omitempty" name:"ModuleIdList"`
}

type DescribeProjectListRequest struct {
	*tchttp.BaseRequest
	
	// 页容，初版默认20，将来可能根据屏幕宽度动态变化
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页标
	PageNo *uint64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// 检索模糊字段
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 是否全部展示，如果是ture，则忽略分页
	AllPage *bool `json:"AllPage,omitnil,omitempty" name:"AllPage"`

	// 角色信息
	ModuleCollection *string `json:"ModuleCollection,omitnil,omitempty" name:"ModuleCollection"`

	// moduleId集合
	ModuleIdList []*string `json:"ModuleIdList,omitnil,omitempty" name:"ModuleIdList"`
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
	delete(f, "ModuleIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectListResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// 额外信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 接口信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *ProjectListData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeResourceUserGroupPageListRequestParams struct {
	// 项目Id
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 页面Id
	PageId *int64 `json:"PageId,omitnil,omitempty" name:"PageId"`

	// 资源类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 是否分页
	AllPage *int64 `json:"AllPage,omitnil,omitempty" name:"AllPage"`

	// 模糊搜索关键字
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 页码
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// 每页条数
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 标签过滤条件
	TagValueList []*ResourceTagValue `json:"TagValueList,omitnil,omitempty" name:"TagValueList"`

	// 角色
	ModuleCollection *string `json:"ModuleCollection,omitnil,omitempty" name:"ModuleCollection"`

	// 是否授权
	ResourceValue *string `json:"ResourceValue,omitnil,omitempty" name:"ResourceValue"`

	// 权限类型
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 父id
	ParentId *int64 `json:"ParentId,omitnil,omitempty" name:"ParentId"`

	// 资源id
	EntityId *int64 `json:"EntityId,omitnil,omitempty" name:"EntityId"`
}

type DescribeResourceUserGroupPageListRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 页面Id
	PageId *int64 `json:"PageId,omitnil,omitempty" name:"PageId"`

	// 资源类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 是否分页
	AllPage *int64 `json:"AllPage,omitnil,omitempty" name:"AllPage"`

	// 模糊搜索关键字
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 页码
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// 每页条数
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 标签过滤条件
	TagValueList []*ResourceTagValue `json:"TagValueList,omitnil,omitempty" name:"TagValueList"`

	// 角色
	ModuleCollection *string `json:"ModuleCollection,omitnil,omitempty" name:"ModuleCollection"`

	// 是否授权
	ResourceValue *string `json:"ResourceValue,omitnil,omitempty" name:"ResourceValue"`

	// 权限类型
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 父id
	ParentId *int64 `json:"ParentId,omitnil,omitempty" name:"ParentId"`

	// 资源id
	EntityId *int64 `json:"EntityId,omitnil,omitempty" name:"EntityId"`
}

func (r *DescribeResourceUserGroupPageListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceUserGroupPageListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "PageId")
	delete(f, "ResourceType")
	delete(f, "AllPage")
	delete(f, "Keyword")
	delete(f, "PageNo")
	delete(f, "PageSize")
	delete(f, "TagValueList")
	delete(f, "ModuleCollection")
	delete(f, "ResourceValue")
	delete(f, "ResourceName")
	delete(f, "ParentId")
	delete(f, "EntityId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourceUserGroupPageListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceUserGroupPageListResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// 扩展
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *UserGroupPageTreeVO `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeResourceUserGroupPageListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourceUserGroupPageListResponseParams `json:"Response"`
}

func (r *DescribeResourceUserGroupPageListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceUserGroupPageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSourceFieldListRequestParams struct {
	// 数据源Id
	DataSourceId *int64 `json:"DataSourceId,omitnil,omitempty" name:"DataSourceId"`

	// 表名
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// sql内容
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// 项目id
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 是否是异步
	AsyncRequest *bool `json:"AsyncRequest,omitnil,omitempty" name:"AsyncRequest"`

	// 异步事务id
	TranId *string `json:"TranId,omitnil,omitempty" name:"TranId"`

	// 11
	ParamList []*ParamCreateDTO `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// DLC扩展参数
	DlcExtInfo *string `json:"DlcExtInfo,omitnil,omitempty" name:"DlcExtInfo"`

	// 是否查询数据库
	QueryDbData *string `json:"QueryDbData,omitnil,omitempty" name:"QueryDbData"`

	// 数据表 Id
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 后端提供字典：数据表类型，1、数据库建表，2、SQL建表，3、Excel上传，4、API接入，5、腾讯文档，6、云数据库，7、手工输入，8、关联查询
	TableType *uint64 `json:"TableType,omitnil,omitempty" name:"TableType"`
}

type DescribeSourceFieldListRequest struct {
	*tchttp.BaseRequest
	
	// 数据源Id
	DataSourceId *int64 `json:"DataSourceId,omitnil,omitempty" name:"DataSourceId"`

	// 表名
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// sql内容
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// 项目id
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 是否是异步
	AsyncRequest *bool `json:"AsyncRequest,omitnil,omitempty" name:"AsyncRequest"`

	// 异步事务id
	TranId *string `json:"TranId,omitnil,omitempty" name:"TranId"`

	// 11
	ParamList []*ParamCreateDTO `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// DLC扩展参数
	DlcExtInfo *string `json:"DlcExtInfo,omitnil,omitempty" name:"DlcExtInfo"`

	// 是否查询数据库
	QueryDbData *string `json:"QueryDbData,omitnil,omitempty" name:"QueryDbData"`

	// 数据表 Id
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 后端提供字典：数据表类型，1、数据库建表，2、SQL建表，3、Excel上传，4、API接入，5、腾讯文档，6、云数据库，7、手工输入，8、关联查询
	TableType *uint64 `json:"TableType,omitnil,omitempty" name:"TableType"`
}

func (r *DescribeSourceFieldListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSourceFieldListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataSourceId")
	delete(f, "TableName")
	delete(f, "Sql")
	delete(f, "ProjectId")
	delete(f, "AsyncRequest")
	delete(f, "TranId")
	delete(f, "ParamList")
	delete(f, "DlcExtInfo")
	delete(f, "QueryDbData")
	delete(f, "TableId")
	delete(f, "TableType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSourceFieldListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSourceFieldListResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// 额外信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 表中字段的列表
	Data *TableColumnListData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSourceFieldListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSourceFieldListResponseParams `json:"Response"`
}

func (r *DescribeSourceFieldListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSourceFieldListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserGroupInfoRequestParams struct {
	// 用户组id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`
}

type DescribeUserGroupInfoRequest struct {
	*tchttp.BaseRequest
	
	// 用户组id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`
}

func (r *DescribeUserGroupInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserGroupInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserGroupInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserGroupInfoResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserGroupInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserGroupInfoResponseParams `json:"Response"`
}

func (r *DescribeUserGroupInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserGroupInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserGroupMemberListRequestParams struct {
	// <p>用户组id集合</p>
	GroupIds []*int64 `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// <p>asc正序,desc倒序</p>
	CreatedOnOrder *string `json:"CreatedOnOrder,omitnil,omitempty" name:"CreatedOnOrder"`

	// <p>搜索关键字</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>分页大小</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>分页页码</p>
	PageNo *uint64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// <p>是否需要分页</p>
	AllPage *bool `json:"AllPage,omitnil,omitempty" name:"AllPage"`
}

type DescribeUserGroupMemberListRequest struct {
	*tchttp.BaseRequest
	
	// <p>用户组id集合</p>
	GroupIds []*int64 `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// <p>asc正序,desc倒序</p>
	CreatedOnOrder *string `json:"CreatedOnOrder,omitnil,omitempty" name:"CreatedOnOrder"`

	// <p>搜索关键字</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>分页大小</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>分页页码</p>
	PageNo *uint64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// <p>是否需要分页</p>
	AllPage *bool `json:"AllPage,omitnil,omitempty" name:"AllPage"`
}

func (r *DescribeUserGroupMemberListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserGroupMemberListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupIds")
	delete(f, "CreatedOnOrder")
	delete(f, "Keyword")
	delete(f, "PageSize")
	delete(f, "PageNo")
	delete(f, "AllPage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserGroupMemberListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserGroupMemberListResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// <p>额外信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// <p>结果信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// <p>数据</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *DescribeUserGroupMemberPageListContainer `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserGroupMemberListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserGroupMemberListResponseParams `json:"Response"`
}

func (r *DescribeUserGroupMemberListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserGroupMemberListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserGroupMemberPageListContainer struct {
	// 列表数据集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*UserGroupMemberVO `json:"List,omitnil,omitempty" name:"List"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 总页数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPages *int64 `json:"TotalPages,omitnil,omitempty" name:"TotalPages"`
}

// Predefined struct for user
type DescribeUserGroupTreeListRequestParams struct {
	// 项目Id
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 是否分页
	AllPage *int64 `json:"AllPage,omitnil,omitempty" name:"AllPage"`

	// 页码
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// 每页条数
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 标签过滤条件
	TagValueList []*ResourceTagValue `json:"TagValueList,omitnil,omitempty" name:"TagValueList"`

	// 节点集合
	Nodes []*UserGroupTreeNodeDTO `json:"Nodes,omitnil,omitempty" name:"Nodes"`

	// 用户组id集合
	GroupIds []*int64 `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// 查询下一个code
	QueryNextNode *bool `json:"QueryNextNode,omitnil,omitempty" name:"QueryNextNode"`

	// 父id集合
	ParentIds []*int64 `json:"ParentIds,omitnil,omitempty" name:"ParentIds"`

	// 是否查询所有节点
	QueryAllNode *bool `json:"QueryAllNode,omitnil,omitempty" name:"QueryAllNode"`

	// 过滤组id集合
	FilterGroupIds []*int64 `json:"FilterGroupIds,omitnil,omitempty" name:"FilterGroupIds"`

	// 模糊搜索关键字
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

type DescribeUserGroupTreeListRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 是否分页
	AllPage *int64 `json:"AllPage,omitnil,omitempty" name:"AllPage"`

	// 页码
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// 每页条数
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 标签过滤条件
	TagValueList []*ResourceTagValue `json:"TagValueList,omitnil,omitempty" name:"TagValueList"`

	// 节点集合
	Nodes []*UserGroupTreeNodeDTO `json:"Nodes,omitnil,omitempty" name:"Nodes"`

	// 用户组id集合
	GroupIds []*int64 `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// 查询下一个code
	QueryNextNode *bool `json:"QueryNextNode,omitnil,omitempty" name:"QueryNextNode"`

	// 父id集合
	ParentIds []*int64 `json:"ParentIds,omitnil,omitempty" name:"ParentIds"`

	// 是否查询所有节点
	QueryAllNode *bool `json:"QueryAllNode,omitnil,omitempty" name:"QueryAllNode"`

	// 过滤组id集合
	FilterGroupIds []*int64 `json:"FilterGroupIds,omitnil,omitempty" name:"FilterGroupIds"`

	// 模糊搜索关键字
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

func (r *DescribeUserGroupTreeListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserGroupTreeListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "AllPage")
	delete(f, "PageNo")
	delete(f, "PageSize")
	delete(f, "TagValueList")
	delete(f, "Nodes")
	delete(f, "GroupIds")
	delete(f, "QueryNextNode")
	delete(f, "ParentIds")
	delete(f, "QueryAllNode")
	delete(f, "FilterGroupIds")
	delete(f, "Keyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserGroupTreeListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserGroupTreeListResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// 扩展
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*UserGroupTreeVO `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserGroupTreeListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserGroupTreeListResponseParams `json:"Response"`
}

func (r *DescribeUserGroupTreeListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserGroupTreeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserProjectListRequestParams struct {
	// <p>项目ID</p>
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>无</p>
	AllPage *bool `json:"AllPage,omitnil,omitempty" name:"AllPage"`

	// <p>无</p>
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// <p>无</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>是否过滤掉企业管理员</p>
	IsFilterPerAuthUser *bool `json:"IsFilterPerAuthUser,omitnil,omitempty" name:"IsFilterPerAuthUser"`

	// <p>是否过滤掉当前用户</p>
	IsFilterCurrentUser *bool `json:"IsFilterCurrentUser,omitnil,omitempty" name:"IsFilterCurrentUser"`

	// <p>关键字</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

type DescribeUserProjectListRequest struct {
	*tchttp.BaseRequest
	
	// <p>项目ID</p>
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>无</p>
	AllPage *bool `json:"AllPage,omitnil,omitempty" name:"AllPage"`

	// <p>无</p>
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// <p>无</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>是否过滤掉企业管理员</p>
	IsFilterPerAuthUser *bool `json:"IsFilterPerAuthUser,omitnil,omitempty" name:"IsFilterPerAuthUser"`

	// <p>是否过滤掉当前用户</p>
	IsFilterCurrentUser *bool `json:"IsFilterCurrentUser,omitnil,omitempty" name:"IsFilterCurrentUser"`

	// <p>关键字</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
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
	delete(f, "IsFilterPerAuthUser")
	delete(f, "IsFilterCurrentUser")
	delete(f, "Keyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserProjectListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserProjectListResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// <p>数据</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *CorpUserListData `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>扩展</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// <p>消息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// 页数
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 全部页码
	AllPage *bool `json:"AllPage,omitnil,omitempty" name:"AllPage"`

	// 0 企业用户 1 访客 不填表示所有用户
	UserType *string `json:"UserType,omitnil,omitempty" name:"UserType"`

	// 模糊搜索的关键字
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 是否只获取绑定企微应用的
	IsOnlyBindAppUser *bool `json:"IsOnlyBindAppUser,omitnil,omitempty" name:"IsOnlyBindAppUser"`
}

type DescribeUserRoleListRequest struct {
	*tchttp.BaseRequest
	
	// 页码
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// 页数
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 全部页码
	AllPage *bool `json:"AllPage,omitnil,omitempty" name:"AllPage"`

	// 0 企业用户 1 访客 不填表示所有用户
	UserType *string `json:"UserType,omitnil,omitempty" name:"UserType"`

	// 模糊搜索的关键字
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 是否只获取绑定企微应用的
	IsOnlyBindAppUser *bool `json:"IsOnlyBindAppUser,omitnil,omitempty" name:"IsOnlyBindAppUser"`
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
	delete(f, "IsOnlyBindAppUser")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserRoleListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserRoleListResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// 扩展描述信息(提供更多异常信息,用于辅助判断)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *UserRoleListData `json:"Data,omitnil,omitempty" name:"Data"`

	// 消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// 页数
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 是否只获取绑定企微应用的
	IsOnlyBindAppUser *bool `json:"IsOnlyBindAppUser,omitnil,omitempty" name:"IsOnlyBindAppUser"`

	// 是否获取全部数据
	AllPage *bool `json:"AllPage,omitnil,omitempty" name:"AllPage"`

	// 角色编码
	RoleCode *string `json:"RoleCode,omitnil,omitempty" name:"RoleCode"`

	// 用户id列表
	UserIdList []*string `json:"UserIdList,omitnil,omitempty" name:"UserIdList"`

	// 搜索关键词
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

type DescribeUserRoleProjectListRequest struct {
	*tchttp.BaseRequest
	
	// 页码
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// 页数
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 是否只获取绑定企微应用的
	IsOnlyBindAppUser *bool `json:"IsOnlyBindAppUser,omitnil,omitempty" name:"IsOnlyBindAppUser"`

	// 是否获取全部数据
	AllPage *bool `json:"AllPage,omitnil,omitempty" name:"AllPage"`

	// 角色编码
	RoleCode *string `json:"RoleCode,omitnil,omitempty" name:"RoleCode"`

	// 用户id列表
	UserIdList []*string `json:"UserIdList,omitnil,omitempty" name:"UserIdList"`

	// 搜索关键词
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
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
	delete(f, "IsOnlyBindAppUser")
	delete(f, "AllPage")
	delete(f, "RoleCode")
	delete(f, "UserIdList")
	delete(f, "Keyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserRoleProjectListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserRoleProjectListResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// 扩展
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *UserRoleListData `json:"Data,omitnil,omitempty" name:"Data"`

	// 消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type EditCorpTagRequestParams struct {
	// 标签ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 标签名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 导入标签的模式(manual/auto)
	ImportType *string `json:"ImportType,omitnil,omitempty" name:"ImportType"`

	// 自动导入标签表的id
	AutoImportTagTableId *int64 `json:"AutoImportTagTableId,omitnil,omitempty" name:"AutoImportTagTableId"`

	// 自动导入标签的关联字段
	AutoImportField *string `json:"AutoImportField,omitnil,omitempty" name:"AutoImportField"`

	// 是否异步请求
	AsyncRequest *bool `json:"AsyncRequest,omitnil,omitempty" name:"AsyncRequest"`

	// 自动导入标签表的表名
	AutoImportTagTableName *string `json:"AutoImportTagTableName,omitnil,omitempty" name:"AutoImportTagTableName"`

	// 事务id
	TranId *string `json:"TranId,omitnil,omitempty" name:"TranId"`
}

type EditCorpTagRequest struct {
	*tchttp.BaseRequest
	
	// 标签ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 标签名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 导入标签的模式(manual/auto)
	ImportType *string `json:"ImportType,omitnil,omitempty" name:"ImportType"`

	// 自动导入标签表的id
	AutoImportTagTableId *int64 `json:"AutoImportTagTableId,omitnil,omitempty" name:"AutoImportTagTableId"`

	// 自动导入标签的关联字段
	AutoImportField *string `json:"AutoImportField,omitnil,omitempty" name:"AutoImportField"`

	// 是否异步请求
	AsyncRequest *bool `json:"AsyncRequest,omitnil,omitempty" name:"AsyncRequest"`

	// 自动导入标签表的表名
	AutoImportTagTableName *string `json:"AutoImportTagTableName,omitnil,omitempty" name:"AutoImportTagTableName"`

	// 事务id
	TranId *string `json:"TranId,omitnil,omitempty" name:"TranId"`
}

func (r *EditCorpTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EditCorpTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Name")
	delete(f, "ImportType")
	delete(f, "AutoImportTagTableId")
	delete(f, "AutoImportField")
	delete(f, "AsyncRequest")
	delete(f, "AutoImportTagTableName")
	delete(f, "TranId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EditCorpTagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EditCorpTagResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// 扩展
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *EditTagVO `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EditCorpTagResponse struct {
	*tchttp.BaseResponse
	Response *EditCorpTagResponseParams `json:"Response"`
}

func (r *EditCorpTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EditCorpTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditTagVO struct {
	// 事务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranId *string `json:"TranId,omitnil,omitempty" name:"TranId"`

	// 事务状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranStatus *int64 `json:"TranStatus,omitnil,omitempty" name:"TranStatus"`

	// 标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`
}

type EmbedTokenInfo struct {
	// 信息标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 令牌
	// 注意：此字段可能返回 null，表示取不到有效值。
	BIToken *string `json:"BIToken,omitnil,omitempty" name:"BIToken"`

	// 项目Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 创建人
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedUser *string `json:"CreatedUser,omitnil,omitempty" name:"CreatedUser"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 更新人
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedUser *string `json:"UpdatedUser,omitnil,omitempty" name:"UpdatedUser"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// 页面Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageId *string `json:"PageId,omitnil,omitempty" name:"PageId"`

	// 备用
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraParam *string `json:"ExtraParam,omitnil,omitempty" name:"ExtraParam"`

	// 嵌出类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Scope *string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// 过期时间，分钟为单位，最大240
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *uint64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 使用者企业Id(仅用于多用户)
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserCorpId *string `json:"UserCorpId,omitnil,omitempty" name:"UserCorpId"`

	// 使用者Id(仅用于多用户)
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 访问次数限制，限制范围1-99999，为空则不设置访问次数限制
	// 注意：此字段可能返回 null，表示取不到有效值。
	TicketNum *int64 `json:"TicketNum,omitnil,omitempty" name:"TicketNum"`

	// 全局参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	GlobalParam *string `json:"GlobalParam,omitnil,omitempty" name:"GlobalParam"`

	// embed表示页面看板嵌出，chatBIEmbed表示ChatBI嵌出
	// 注意：此字段可能返回 null，表示取不到有效值。
	Intention *string `json:"Intention,omitnil,omitempty" name:"Intention"`

	// 100 无绑定用户
	// 200 单用户单token
	// 300 单用户 多token
	// 注意：此字段可能返回 null，表示取不到有效值。
	TokenType *int64 `json:"TokenType,omitnil,omitempty" name:"TokenType"`

	// token 数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TokenNum *int64 `json:"TokenNum,omitnil,omitempty" name:"TokenNum"`

	// 是否单用户多token
	// 注意：此字段可能返回 null，表示取不到有效值。
	SingleUserMultiToken *bool `json:"SingleUserMultiToken,omitnil,omitempty" name:"SingleUserMultiToken"`

	// 嵌出显示配置，目前为ChatBI嵌出场景用，TableFilter表示数据表列表过滤，SqlView表示sql查看功能
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigParam *string `json:"ConfigParam,omitnil,omitempty" name:"ConfigParam"`
}

type EmptyValue struct {
	// 空值展示样式类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 空值展示样式类型对应具体的展示字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	Custom *string `json:"Custom,omitnil,omitempty" name:"Custom"`
}

type EmptyValueConfig struct {
	// 数值类字段空值样式配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Number *EmptyValue `json:"Number,omitnil,omitempty" name:"Number"`

	// 字符串字段空置样式配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	String *EmptyValue `json:"String,omitnil,omitempty" name:"String"`
}

type ErrorInfo struct {
	// 错误说明字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorTip *string `json:"ErrorTip,omitnil,omitempty" name:"ErrorTip"`

	// 原始异常信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// 错误等级字段
	// ERROR
	// WARN
	// INFO
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorLevel *string `json:"ErrorLevel,omitnil,omitempty" name:"ErrorLevel"`

	// 指引文档链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	DocLink *string `json:"DocLink,omitnil,omitempty" name:"DocLink"`

	// 快速指引说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	FAQ *string `json:"FAQ,omitnil,omitempty" name:"FAQ"`

	// 预留字段1
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReservedField *string `json:"ReservedField,omitnil,omitempty" name:"ReservedField"`
}

// Predefined struct for user
type ExportScreenPageRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 页面id
	PageId *string `json:"PageId,omitnil,omitempty" name:"PageId"`

	// 画布类型。栅格画布：GRID；自由画布：FREE
	CanvasType *string `json:"CanvasType,omitnil,omitempty" name:"CanvasType"`

	// 图片导出类型。base64；url（有效期：1天）
	PicType *string `json:"PicType,omitnil,omitempty" name:"PicType"`

	// 组件Ids。为空时，导出整个页面
	WidgetIds []*string `json:"WidgetIds,omitnil,omitempty" name:"WidgetIds"`

	// 是否是异步请求
	AsyncRequest *bool `json:"AsyncRequest,omitnil,omitempty" name:"AsyncRequest"`

	// 事务id
	TranId *string `json:"TranId,omitnil,omitempty" name:"TranId"`
}

type ExportScreenPageRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 页面id
	PageId *string `json:"PageId,omitnil,omitempty" name:"PageId"`

	// 画布类型。栅格画布：GRID；自由画布：FREE
	CanvasType *string `json:"CanvasType,omitnil,omitempty" name:"CanvasType"`

	// 图片导出类型。base64；url（有效期：1天）
	PicType *string `json:"PicType,omitnil,omitempty" name:"PicType"`

	// 组件Ids。为空时，导出整个页面
	WidgetIds []*string `json:"WidgetIds,omitnil,omitempty" name:"WidgetIds"`

	// 是否是异步请求
	AsyncRequest *bool `json:"AsyncRequest,omitnil,omitempty" name:"AsyncRequest"`

	// 事务id
	TranId *string `json:"TranId,omitnil,omitempty" name:"TranId"`
}

func (r *ExportScreenPageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportScreenPageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "PageId")
	delete(f, "CanvasType")
	delete(f, "PicType")
	delete(f, "WidgetIds")
	delete(f, "AsyncRequest")
	delete(f, "TranId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportScreenPageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportScreenPageResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// 扩展参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 返回数据结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *PageScreenListVO `json:"Data,omitnil,omitempty" name:"Data"`

	// 返回消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ExportScreenPageResponse struct {
	*tchttp.BaseResponse
	Response *ExportScreenPageResponseParams `json:"Response"`
}

func (r *ExportScreenPageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportScreenPageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FieldRemarkDTO struct {
	// 字段名
	// 注意：此字段可能返回 null，表示取不到有效值。
	FieldName *string `json:"FieldName,omitnil,omitempty" name:"FieldName"`

	// 字段备注列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comment []*string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

type FrequencyConfig struct {
	// 周期
	// 注意：此字段可能返回 null，表示取不到有效值。
	Frequency *string `json:"Frequency,omitnil,omitempty" name:"Frequency"`

	// 日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	Dates []*int64 `json:"Dates,omitnil,omitempty" name:"Dates"`

	// 时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 间隔时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntervalTime *uint64 `json:"IntervalTime,omitnil,omitempty" name:"IntervalTime"`

	// 1:SECOND,2:MINUTE,3:HOUR,4:DAY
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntervalTimeUnit *uint64 `json:"IntervalTimeUnit,omitnil,omitempty" name:"IntervalTimeUnit"`

	// 小时列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Hours []*int64 `json:"Hours,omitnil,omitempty" name:"Hours"`

	// 分钟列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Minute []*int64 `json:"Minute,omitnil,omitempty" name:"Minute"`
}

type IdDTO struct {
	// <p>请求id</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>key</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessKey *string `json:"AccessKey,omitnil,omitempty" name:"AccessKey"`

	// <p>id</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>事务id</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranId *string `json:"TranId,omitnil,omitempty" name:"TranId"`

	// <p>事务状态</p><p>枚举值：</p><ul><li>1： 处理中</li><li>2： 处理成功</li><li>3： 处理失败</li></ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranStatus *int64 `json:"TranStatus,omitnil,omitempty" name:"TranStatus"`
}

type JoinRelation struct {
	// 表关联关系id,前端使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	JoinId *string `json:"JoinId,omitnil,omitempty" name:"JoinId"`

	// 原表节点id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceTableNodeId *string `json:"SourceTableNodeId,omitnil,omitempty" name:"SourceTableNodeId"`

	// 目标表节点id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetTableNodeId *string `json:"TargetTableNodeId,omitnil,omitempty" name:"TargetTableNodeId"`

	// 多表关联的关联类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	JoinType *string `json:"JoinType,omitnil,omitempty" name:"JoinType"`

	// 多表关联的字段列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Fields []*JoinRelationField `json:"Fields,omitnil,omitempty" name:"Fields"`
}

type JoinRelationField struct {
	// 字段关联关系id,前端使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	FieldJoinId *string `json:"FieldJoinId,omitnil,omitempty" name:"FieldJoinId"`

	// 原表字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceField *TableField `json:"SourceField,omitnil,omitempty" name:"SourceField"`

	// 目标表字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetField *TableField `json:"TargetField,omitnil,omitempty" name:"TargetField"`
}

type JoinSourceTable struct {
	// 1:数据源原表,2:本地表,3:Excel表,4:API表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableNodeType *uint64 `json:"TableNodeType,omitnil,omitempty" name:"TableNodeType"`

	// 原始表节点Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableNodeId *string `json:"TableNodeId,omitnil,omitempty" name:"TableNodeId"`

	// 父节点id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentId *string `json:"ParentId,omitnil,omitempty" name:"ParentId"`

	// 非必填, 数据源原表没有ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 必填,数据源原表用原始表名, 其他类型用BI的逻辑表名
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 原始表需要展示的字段列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Fields []*TableField `json:"Fields,omitnil,omitempty" name:"Fields"`

	// 数据源id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceId *uint64 `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 非必填,前端展示的数据源别名,excel建表需要
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableAlias *string `json:"TableAlias,omitnil,omitempty" name:"TableAlias"`
}

// Predefined struct for user
type ModifyAuthApiKeyRequestParams struct {
	// <p>ApiKey</p>
	ApiKey *string `json:"ApiKey,omitnil,omitempty" name:"ApiKey"`

	// <p>默认用户</p>
	DefaultUser *string `json:"DefaultUser,omitnil,omitempty" name:"DefaultUser"`
}

type ModifyAuthApiKeyRequest struct {
	*tchttp.BaseRequest
	
	// <p>ApiKey</p>
	ApiKey *string `json:"ApiKey,omitnil,omitempty" name:"ApiKey"`

	// <p>默认用户</p>
	DefaultUser *string `json:"DefaultUser,omitnil,omitempty" name:"DefaultUser"`
}

func (r *ModifyAuthApiKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAuthApiKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApiKey")
	delete(f, "DefaultUser")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAuthApiKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAuthApiKeyResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// <p>&quot;&quot;</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// <p>&quot;success&quot;</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// <p>数据</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *ApiKeyAuthApplyVO `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAuthApiKeyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAuthApiKeyResponseParams `json:"Response"`
}

func (r *ModifyAuthApiKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAuthApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDatasourceCloudRequestParams struct {
	// <p>后端提供字典：域类型，1、腾讯云，2、本地</p>
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// <p>驱动</p><p>枚举值：</p><ul><li>MYSQL： MySQL数据库</li><li>PRESTO： PRESTO数据库</li><li>POSTGRE： PostgreSQL数据库</li><li>DLC： DLC数据库</li><li>MSSQL： 微软SQL Server数据库</li></ul>
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// <p>数据库编码</p>
	Charset *string `json:"Charset,omitnil,omitempty" name:"Charset"`

	// <p>用户名</p>
	DbUser *string `json:"DbUser,omitnil,omitempty" name:"DbUser"`

	// <p>密码</p>
	DbPwd *string `json:"DbPwd,omitnil,omitempty" name:"DbPwd"`

	// <p>数据库名称</p>
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// <p>数据库别名</p>
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// <p>项目ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>住键</p>
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>公有云内网ip</p>
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// <p>公有云内网端口</p>
	Vport *string `json:"Vport,omitnil,omitempty" name:"Vport"`

	// <p>vpc标识</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>统一vpc标识</p>
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// <p>区域标识（gz,bj)</p>
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// <p>扩展参数</p>
	ExtraParam *string `json:"ExtraParam,omitnil,omitempty" name:"ExtraParam"`

	// <p>实例id</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>数据源产品名</p>
	ProdDbName *string `json:"ProdDbName,omitnil,omitempty" name:"ProdDbName"`

	// <p>第三方数据源标识</p>
	DataOrigin *string `json:"DataOrigin,omitnil,omitempty" name:"DataOrigin"`

	// <p>第三方项目id</p>
	DataOriginProjectId *string `json:"DataOriginProjectId,omitnil,omitempty" name:"DataOriginProjectId"`

	// <p>第三方数据源id</p>
	DataOriginDatasourceId *string `json:"DataOriginDatasourceId,omitnil,omitempty" name:"DataOriginDatasourceId"`

	// <p>集群id</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>数据库schema</p>
	Schema *string `json:"Schema,omitnil,omitempty" name:"Schema"`

	// <p>数据库版本</p>
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`
}

type ModifyDatasourceCloudRequest struct {
	*tchttp.BaseRequest
	
	// <p>后端提供字典：域类型，1、腾讯云，2、本地</p>
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// <p>驱动</p><p>枚举值：</p><ul><li>MYSQL： MySQL数据库</li><li>PRESTO： PRESTO数据库</li><li>POSTGRE： PostgreSQL数据库</li><li>DLC： DLC数据库</li><li>MSSQL： 微软SQL Server数据库</li></ul>
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// <p>数据库编码</p>
	Charset *string `json:"Charset,omitnil,omitempty" name:"Charset"`

	// <p>用户名</p>
	DbUser *string `json:"DbUser,omitnil,omitempty" name:"DbUser"`

	// <p>密码</p>
	DbPwd *string `json:"DbPwd,omitnil,omitempty" name:"DbPwd"`

	// <p>数据库名称</p>
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// <p>数据库别名</p>
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// <p>项目ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>住键</p>
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>公有云内网ip</p>
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// <p>公有云内网端口</p>
	Vport *string `json:"Vport,omitnil,omitempty" name:"Vport"`

	// <p>vpc标识</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>统一vpc标识</p>
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// <p>区域标识（gz,bj)</p>
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// <p>扩展参数</p>
	ExtraParam *string `json:"ExtraParam,omitnil,omitempty" name:"ExtraParam"`

	// <p>实例id</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>数据源产品名</p>
	ProdDbName *string `json:"ProdDbName,omitnil,omitempty" name:"ProdDbName"`

	// <p>第三方数据源标识</p>
	DataOrigin *string `json:"DataOrigin,omitnil,omitempty" name:"DataOrigin"`

	// <p>第三方项目id</p>
	DataOriginProjectId *string `json:"DataOriginProjectId,omitnil,omitempty" name:"DataOriginProjectId"`

	// <p>第三方数据源id</p>
	DataOriginDatasourceId *string `json:"DataOriginDatasourceId,omitnil,omitempty" name:"DataOriginDatasourceId"`

	// <p>集群id</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>数据库schema</p>
	Schema *string `json:"Schema,omitnil,omitempty" name:"Schema"`

	// <p>数据库版本</p>
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`
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
	delete(f, "ClusterId")
	delete(f, "Schema")
	delete(f, "DbVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDatasourceCloudRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDatasourceCloudResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// <p>无</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>额外信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// <p>提示</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// <p>HOST</p>
	DbHost *string `json:"DbHost,omitnil,omitempty" name:"DbHost"`

	// <p>端口</p>
	DbPort *uint64 `json:"DbPort,omitnil,omitempty" name:"DbPort"`

	// <p>后端提供字典：域类型，1、腾讯云，2、本地</p>
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// <p>驱动</p><p>枚举值：</p><ul><li>MYSQL： MySQL数据库</li><li>PRESTO： PRESTO数据库</li><li>POSTGRE： PostgreSQL数据库</li><li>DLC： DLC数据库</li><li>MSSQL： 微软SQL Server数据库</li></ul>
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// <p>数据库编码</p>
	Charset *string `json:"Charset,omitnil,omitempty" name:"Charset"`

	// <p>用户名</p>
	DbUser *string `json:"DbUser,omitnil,omitempty" name:"DbUser"`

	// <p>密码</p>
	DbPwd *string `json:"DbPwd,omitnil,omitempty" name:"DbPwd"`

	// <p>数据库名称</p>
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// <p>数据库别名</p>
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// <p>数据源id</p>
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>项目ID</p>
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>catalog值</p>
	Catalog *string `json:"Catalog,omitnil,omitempty" name:"Catalog"`

	// <p>第三方数据源标识</p>
	DataOrigin *string `json:"DataOrigin,omitnil,omitempty" name:"DataOrigin"`

	// <p>第三方项目id</p>
	DataOriginProjectId *string `json:"DataOriginProjectId,omitnil,omitempty" name:"DataOriginProjectId"`

	// <p>第三方数据源id</p>
	DataOriginDatasourceId *string `json:"DataOriginDatasourceId,omitnil,omitempty" name:"DataOriginDatasourceId"`

	// <p>扩展参数</p>
	ExtraParam *string `json:"ExtraParam,omitnil,omitempty" name:"ExtraParam"`

	// <p>腾讯云私有网络统一标识</p>
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// <p>私有网络ip</p>
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// <p>私有网络端口</p>
	Vport *string `json:"Vport,omitnil,omitempty" name:"Vport"`

	// <p>腾讯云私有网络标识</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>开启vpc</p>
	UseVPC *bool `json:"UseVPC,omitnil,omitempty" name:"UseVPC"`

	// <p>地域</p>
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// <p>数据库schema</p>
	Schema *string `json:"Schema,omitnil,omitempty" name:"Schema"`

	// <p>数据库版本</p>
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`
}

type ModifyDatasourceRequest struct {
	*tchttp.BaseRequest
	
	// <p>HOST</p>
	DbHost *string `json:"DbHost,omitnil,omitempty" name:"DbHost"`

	// <p>端口</p>
	DbPort *uint64 `json:"DbPort,omitnil,omitempty" name:"DbPort"`

	// <p>后端提供字典：域类型，1、腾讯云，2、本地</p>
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// <p>驱动</p><p>枚举值：</p><ul><li>MYSQL： MySQL数据库</li><li>PRESTO： PRESTO数据库</li><li>POSTGRE： PostgreSQL数据库</li><li>DLC： DLC数据库</li><li>MSSQL： 微软SQL Server数据库</li></ul>
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// <p>数据库编码</p>
	Charset *string `json:"Charset,omitnil,omitempty" name:"Charset"`

	// <p>用户名</p>
	DbUser *string `json:"DbUser,omitnil,omitempty" name:"DbUser"`

	// <p>密码</p>
	DbPwd *string `json:"DbPwd,omitnil,omitempty" name:"DbPwd"`

	// <p>数据库名称</p>
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// <p>数据库别名</p>
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// <p>数据源id</p>
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>项目ID</p>
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>catalog值</p>
	Catalog *string `json:"Catalog,omitnil,omitempty" name:"Catalog"`

	// <p>第三方数据源标识</p>
	DataOrigin *string `json:"DataOrigin,omitnil,omitempty" name:"DataOrigin"`

	// <p>第三方项目id</p>
	DataOriginProjectId *string `json:"DataOriginProjectId,omitnil,omitempty" name:"DataOriginProjectId"`

	// <p>第三方数据源id</p>
	DataOriginDatasourceId *string `json:"DataOriginDatasourceId,omitnil,omitempty" name:"DataOriginDatasourceId"`

	// <p>扩展参数</p>
	ExtraParam *string `json:"ExtraParam,omitnil,omitempty" name:"ExtraParam"`

	// <p>腾讯云私有网络统一标识</p>
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// <p>私有网络ip</p>
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// <p>私有网络端口</p>
	Vport *string `json:"Vport,omitnil,omitempty" name:"Vport"`

	// <p>腾讯云私有网络标识</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>开启vpc</p>
	UseVPC *bool `json:"UseVPC,omitnil,omitempty" name:"UseVPC"`

	// <p>地域</p>
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// <p>数据库schema</p>
	Schema *string `json:"Schema,omitnil,omitempty" name:"Schema"`

	// <p>数据库版本</p>
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`
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
	delete(f, "UseVPC")
	delete(f, "RegionId")
	delete(f, "Schema")
	delete(f, "DbVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDatasourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDatasourceResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// <p>无</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>额外信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// <p>提示</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// <p>项目Id</p>
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>名字</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>颜色值</p>
	ColorCode *string `json:"ColorCode,omitnil,omitempty" name:"ColorCode"`

	// <p>图标</p>
	Logo *string `json:"Logo,omitnil,omitempty" name:"Logo"`

	// <p>备注</p>
	Mark *string `json:"Mark,omitnil,omitempty" name:"Mark"`

	// <p>可申请</p>
	IsApply *bool `json:"IsApply,omitnil,omitempty" name:"IsApply"`

	// <p>种子</p>
	Seed *string `json:"Seed,omitnil,omitempty" name:"Seed"`

	// <p>默认看板</p><p>枚举值：</p><ul><li>1： 项目看板</li><li>2： 我的看板</li></ul>
	DefaultPanelType *int64 `json:"DefaultPanelType,omitnil,omitempty" name:"DefaultPanelType"`

	// <p>2</p>
	PanelScope *string `json:"PanelScope,omitnil,omitempty" name:"PanelScope"`

	// <p>项目管理平台</p>
	ManagePlatform *string `json:"ManagePlatform,omitnil,omitempty" name:"ManagePlatform"`
}

type ModifyProjectRequest struct {
	*tchttp.BaseRequest
	
	// <p>项目Id</p>
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>名字</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>颜色值</p>
	ColorCode *string `json:"ColorCode,omitnil,omitempty" name:"ColorCode"`

	// <p>图标</p>
	Logo *string `json:"Logo,omitnil,omitempty" name:"Logo"`

	// <p>备注</p>
	Mark *string `json:"Mark,omitnil,omitempty" name:"Mark"`

	// <p>可申请</p>
	IsApply *bool `json:"IsApply,omitnil,omitempty" name:"IsApply"`

	// <p>种子</p>
	Seed *string `json:"Seed,omitnil,omitempty" name:"Seed"`

	// <p>默认看板</p><p>枚举值：</p><ul><li>1： 项目看板</li><li>2： 我的看板</li></ul>
	DefaultPanelType *int64 `json:"DefaultPanelType,omitnil,omitempty" name:"DefaultPanelType"`

	// <p>2</p>
	PanelScope *string `json:"PanelScope,omitnil,omitempty" name:"PanelScope"`

	// <p>项目管理平台</p>
	ManagePlatform *string `json:"ManagePlatform,omitnil,omitempty" name:"ManagePlatform"`
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
	delete(f, "ManagePlatform")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProjectResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// <p>额外信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// <p>返回数据</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>结果</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

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
type ModifyResourceUserGroupRequestParams struct {
	// 项目Id
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 用户组id集合
	UserGroupIds []*int64 `json:"UserGroupIds,omitnil,omitempty" name:"UserGroupIds"`

	// 操作的资源权限
	Resource *ResourceListDTO `json:"Resource,omitnil,omitempty" name:"Resource"`

	// 资源类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 是否查询所有孩子节点
	QueryChildren *bool `json:"QueryChildren,omitnil,omitempty" name:"QueryChildren"`
}

type ModifyResourceUserGroupRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 用户组id集合
	UserGroupIds []*int64 `json:"UserGroupIds,omitnil,omitempty" name:"UserGroupIds"`

	// 操作的资源权限
	Resource *ResourceListDTO `json:"Resource,omitnil,omitempty" name:"Resource"`

	// 资源类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 是否查询所有孩子节点
	QueryChildren *bool `json:"QueryChildren,omitnil,omitempty" name:"QueryChildren"`
}

func (r *ModifyResourceUserGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourceUserGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "UserGroupIds")
	delete(f, "Resource")
	delete(f, "ResourceType")
	delete(f, "QueryChildren")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyResourceUserGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourceUserGroupResourceRequestParams struct {
	// 项目Id
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 用户id
	UserGroupId *int64 `json:"UserGroupId,omitnil,omitempty" name:"UserGroupId"`

	// 资源
	Resource *UserResourceDTO `json:"Resource,omitnil,omitempty" name:"Resource"`

	// 实体类
	EntityIds []*int64 `json:"EntityIds,omitnil,omitempty" name:"EntityIds"`

	// 资源类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`
}

type ModifyResourceUserGroupResourceRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 用户id
	UserGroupId *int64 `json:"UserGroupId,omitnil,omitempty" name:"UserGroupId"`

	// 资源
	Resource *UserResourceDTO `json:"Resource,omitnil,omitempty" name:"Resource"`

	// 实体类
	EntityIds []*int64 `json:"EntityIds,omitnil,omitempty" name:"EntityIds"`

	// 资源类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`
}

func (r *ModifyResourceUserGroupResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourceUserGroupResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "UserGroupId")
	delete(f, "Resource")
	delete(f, "EntityIds")
	delete(f, "ResourceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyResourceUserGroupResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourceUserGroupResourceResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// 扩展
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *int64 `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyResourceUserGroupResourceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyResourceUserGroupResourceResponseParams `json:"Response"`
}

func (r *ModifyResourceUserGroupResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourceUserGroupResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourceUserGroupResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// 扩展
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *int64 `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyResourceUserGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyResourceUserGroupResponseParams `json:"Response"`
}

func (r *ModifyResourceUserGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourceUserGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourceUserRequestParams struct {
	// 项目Id
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 用户id
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 资源
	Resource *UserResourceDTO `json:"Resource,omitnil,omitempty" name:"Resource"`

	// 实体类
	EntityIds []*int64 `json:"EntityIds,omitnil,omitempty" name:"EntityIds"`

	// 资源类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`
}

type ModifyResourceUserRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 用户id
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 资源
	Resource *UserResourceDTO `json:"Resource,omitnil,omitempty" name:"Resource"`

	// 实体类
	EntityIds []*int64 `json:"EntityIds,omitnil,omitempty" name:"EntityIds"`

	// 资源类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`
}

func (r *ModifyResourceUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourceUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "UserId")
	delete(f, "Resource")
	delete(f, "EntityIds")
	delete(f, "ResourceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyResourceUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourceUserResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// 扩展
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *int64 `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyResourceUserResponse struct {
	*tchttp.BaseResponse
	Response *ModifyResourceUserResponseParams `json:"Response"`
}

func (r *ModifyResourceUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourceUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTagTableRequestParams struct {
	// 标签表名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 标签表关联的项目id
	AutoImportProjectId *int64 `json:"AutoImportProjectId,omitnil,omitempty" name:"AutoImportProjectId"`

	// 标签表关联的数据表id
	AutoImportTableId *int64 `json:"AutoImportTableId,omitnil,omitempty" name:"AutoImportTableId"`

	// uin对应字段
	AutoImportUinField *string `json:"AutoImportUinField,omitnil,omitempty" name:"AutoImportUinField"`

	// 标签表id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`
}

type ModifyTagTableRequest struct {
	*tchttp.BaseRequest
	
	// 标签表名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 标签表关联的项目id
	AutoImportProjectId *int64 `json:"AutoImportProjectId,omitnil,omitempty" name:"AutoImportProjectId"`

	// 标签表关联的数据表id
	AutoImportTableId *int64 `json:"AutoImportTableId,omitnil,omitempty" name:"AutoImportTableId"`

	// uin对应字段
	AutoImportUinField *string `json:"AutoImportUinField,omitnil,omitempty" name:"AutoImportUinField"`

	// 标签表id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`
}

func (r *ModifyTagTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTagTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "AutoImportProjectId")
	delete(f, "AutoImportTableId")
	delete(f, "AutoImportUinField")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTagTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTagTableResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// 数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *ModifyTagTableVO `json:"Data,omitnil,omitempty" name:"Data"`

	// 扩展
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTagTableResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTagTableResponseParams `json:"Response"`
}

func (r *ModifyTagTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTagTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTagTableVO struct {
	// 标签表id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`
}

// Predefined struct for user
type ModifyUserDetailInfoRequestParams struct {
	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 角色ID 列表
	RoleIdList []*int64 `json:"RoleIdList,omitnil,omitempty" name:"RoleIdList"`

	// 邮箱
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 手机号
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// 手机区号
	AreaCode *string `json:"AreaCode,omitnil,omitempty" name:"AreaCode"`

	// 企业微信应用用户id
	AppUserId *string `json:"AppUserId,omitnil,omitempty" name:"AppUserId"`

	// 是否开启手机验证码登录（0 关闭，1 开启）
	LoginSecurityStatus *int64 `json:"LoginSecurityStatus,omitnil,omitempty" name:"LoginSecurityStatus"`

	// 是否开启密码过期提醒（0 关闭，1 开启
	ResetPassWordTip *int64 `json:"ResetPassWordTip,omitnil,omitempty" name:"ResetPassWordTip"`

	// 强制修改密码（0 关闭，1 开启）
	ForceResetPassWord *int64 `json:"ForceResetPassWord,omitnil,omitempty" name:"ForceResetPassWord"`

	// 密码过期提醒时间，30、60、90（默认）、180天
	PasswordExpired *int64 `json:"PasswordExpired,omitnil,omitempty" name:"PasswordExpired"`

	// 用户组id
	UserGroupIdList []*int64 `json:"UserGroupIdList,omitnil,omitempty" name:"UserGroupIdList"`
}

type ModifyUserDetailInfoRequest struct {
	*tchttp.BaseRequest
	
	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 角色ID 列表
	RoleIdList []*int64 `json:"RoleIdList,omitnil,omitempty" name:"RoleIdList"`

	// 邮箱
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 手机号
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// 手机区号
	AreaCode *string `json:"AreaCode,omitnil,omitempty" name:"AreaCode"`

	// 企业微信应用用户id
	AppUserId *string `json:"AppUserId,omitnil,omitempty" name:"AppUserId"`

	// 是否开启手机验证码登录（0 关闭，1 开启）
	LoginSecurityStatus *int64 `json:"LoginSecurityStatus,omitnil,omitempty" name:"LoginSecurityStatus"`

	// 是否开启密码过期提醒（0 关闭，1 开启
	ResetPassWordTip *int64 `json:"ResetPassWordTip,omitnil,omitempty" name:"ResetPassWordTip"`

	// 强制修改密码（0 关闭，1 开启）
	ForceResetPassWord *int64 `json:"ForceResetPassWord,omitnil,omitempty" name:"ForceResetPassWord"`

	// 密码过期提醒时间，30、60、90（默认）、180天
	PasswordExpired *int64 `json:"PasswordExpired,omitnil,omitempty" name:"PasswordExpired"`

	// 用户组id
	UserGroupIdList []*int64 `json:"UserGroupIdList,omitnil,omitempty" name:"UserGroupIdList"`
}

func (r *ModifyUserDetailInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserDetailInfoRequest) FromJsonString(s string) error {
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
	delete(f, "AppUserId")
	delete(f, "LoginSecurityStatus")
	delete(f, "ResetPassWordTip")
	delete(f, "ForceResetPassWord")
	delete(f, "PasswordExpired")
	delete(f, "UserGroupIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserDetailInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserDetailInfoResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// 扩展
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyUserDetailInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUserDetailInfoResponseParams `json:"Response"`
}

func (r *ModifyUserDetailInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserDetailInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserGroupRequestParams struct {
	// <p>用户组更新list</p>
	UpdateList []*UserGroupUpdateDTO `json:"UpdateList,omitnil,omitempty" name:"UpdateList"`
}

type ModifyUserGroupRequest struct {
	*tchttp.BaseRequest
	
	// <p>用户组更新list</p>
	UpdateList []*UserGroupUpdateDTO `json:"UpdateList,omitnil,omitempty" name:"UpdateList"`
}

func (r *ModifyUserGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UpdateList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserGroupResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// <p>额外信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// <p>结果信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// <p>数据</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*UserGroupVO `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyUserGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUserGroupResponseParams `json:"Response"`
}

func (r *ModifyUserGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserRoleProjectRequestParams struct {
	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 角色ID 列表
	RoleIdList []*int64 `json:"RoleIdList,omitnil,omitempty" name:"RoleIdList"`

	// 邮箱
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 企业微信应用用户id
	AppUserId *string `json:"AppUserId,omitnil,omitempty" name:"AppUserId"`
}

type ModifyUserRoleProjectRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 角色ID 列表
	RoleIdList []*int64 `json:"RoleIdList,omitnil,omitempty" name:"RoleIdList"`

	// 邮箱
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 企业微信应用用户id
	AppUserId *string `json:"AppUserId,omitnil,omitempty" name:"AppUserId"`
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
	delete(f, "AppUserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserRoleProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserRoleProjectResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// 扩展
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 角色ID 列表
	RoleIdList []*int64 `json:"RoleIdList,omitnil,omitempty" name:"RoleIdList"`

	// 邮箱
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 手机号
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// 手机区号
	AreaCode *string `json:"AreaCode,omitnil,omitempty" name:"AreaCode"`

	// 企业微信应用用户id
	AppUserId *string `json:"AppUserId,omitnil,omitempty" name:"AppUserId"`

	// 是否开启手机验证码登录（0 关闭，1 开启）
	LoginSecurityStatus *int64 `json:"LoginSecurityStatus,omitnil,omitempty" name:"LoginSecurityStatus"`

	// 是否开启密码过期提醒（0 关闭，1 开启
	ResetPassWordTip *int64 `json:"ResetPassWordTip,omitnil,omitempty" name:"ResetPassWordTip"`

	// 强制修改密码（0 关闭，1 开启）
	ForceResetPassWord *int64 `json:"ForceResetPassWord,omitnil,omitempty" name:"ForceResetPassWord"`

	// 密码过期提醒时间，30、60、90（默认）、180天
	PasswordExpired *int64 `json:"PasswordExpired,omitnil,omitempty" name:"PasswordExpired"`
}

type ModifyUserRoleRequest struct {
	*tchttp.BaseRequest
	
	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 角色ID 列表
	RoleIdList []*int64 `json:"RoleIdList,omitnil,omitempty" name:"RoleIdList"`

	// 邮箱
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 手机号
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// 手机区号
	AreaCode *string `json:"AreaCode,omitnil,omitempty" name:"AreaCode"`

	// 企业微信应用用户id
	AppUserId *string `json:"AppUserId,omitnil,omitempty" name:"AppUserId"`

	// 是否开启手机验证码登录（0 关闭，1 开启）
	LoginSecurityStatus *int64 `json:"LoginSecurityStatus,omitnil,omitempty" name:"LoginSecurityStatus"`

	// 是否开启密码过期提醒（0 关闭，1 开启
	ResetPassWordTip *int64 `json:"ResetPassWordTip,omitnil,omitempty" name:"ResetPassWordTip"`

	// 强制修改密码（0 关闭，1 开启）
	ForceResetPassWord *int64 `json:"ForceResetPassWord,omitnil,omitempty" name:"ForceResetPassWord"`

	// 密码过期提醒时间，30、60、90（默认）、180天
	PasswordExpired *int64 `json:"PasswordExpired,omitnil,omitempty" name:"PasswordExpired"`
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
	delete(f, "AppUserId")
	delete(f, "LoginSecurityStatus")
	delete(f, "ResetPassWordTip")
	delete(f, "ForceResetPassWord")
	delete(f, "PasswordExpired")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserRoleResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// 扩展
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type ModifyUserTagRequestParams struct {
	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 标签信息
	TagList []*UserTagInfo `json:"TagList,omitnil,omitempty" name:"TagList"`
}

type ModifyUserTagRequest struct {
	*tchttp.BaseRequest
	
	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 标签信息
	TagList []*UserTagInfo `json:"TagList,omitnil,omitempty" name:"TagList"`
}

func (r *ModifyUserTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "TagList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserTagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserTagResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// 消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 额外信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyUserTagResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUserTagResponseParams `json:"Response"`
}

func (r *ModifyUserTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PageScreenListVO struct {
	// 图片导出类型。base64；url
	// 注意：此字段可能返回 null，表示取不到有效值。
	PicType *string `json:"PicType,omitnil,omitempty" name:"PicType"`

	// 图片列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*PageScreenVO `json:"List,omitnil,omitempty" name:"List"`

	// 异步事务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranId *string `json:"TranId,omitnil,omitempty" name:"TranId"`

	// 事务状态
	// 1: 处理中; 2: 处理成功; 3 处理失败(错误内容见外层Msg)
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranStatus *int64 `json:"TranStatus,omitnil,omitempty" name:"TranStatus"`
}

type PageScreenVO struct {
	// 截图base64或 url
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 组件Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	WidgetId *string `json:"WidgetId,omitnil,omitempty" name:"WidgetId"`
}

type ParamCreateDTO struct {
	// 参数名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// 默认值
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultValue *string `json:"DefaultValue,omitnil,omitempty" name:"DefaultValue"`

	// 参数类型，string/datetime/number
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamType *string `json:"ParamType,omitnil,omitempty" name:"ParamType"`

	// 格式化类型，yyyy-MM-dd HH:mm:ss.SSS（只有时间必填）
	// 注意：此字段可能返回 null，表示取不到有效值。
	FormatRule *string `json:"FormatRule,omitnil,omitempty" name:"FormatRule"`

	// 复杂类型，格式化的另一种表达，例如YYYY-MM
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComplexType *string `json:"ComplexType,omitnil,omitempty" name:"ComplexType"`

	// 作用域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Scope *string `json:"Scope,omitnil,omitempty" name:"Scope"`
}

type PermissionComponent struct {
	// <p>权限值</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModuleId *string `json:"ModuleId,omitnil,omitempty" name:"ModuleId"`

	// <p>可用性</p><p>枚举值：</p><ul><li>usable： 可用</li><li>visible： 可见</li><li>disabled： 不可用</li><li>hidden： 隐藏</li></ul><p>默认值：disabled</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncludeType *string `json:"IncludeType,omitnil,omitempty" name:"IncludeType"`

	// <p>目标升级版本</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpgradeVersionType *string `json:"UpgradeVersionType,omitnil,omitempty" name:"UpgradeVersionType"`

	// <p>补充信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tips *string `json:"Tips,omitnil,omitempty" name:"Tips"`

	// <p>补充信息的key值</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TipsKey *string `json:"TipsKey,omitnil,omitempty" name:"TipsKey"`
}

type PermissionGroup struct {
	// 分组名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModuleGroup *string `json:"ModuleGroup,omitnil,omitempty" name:"ModuleGroup"`

	// 权限列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Components []*PermissionComponent `json:"Components,omitnil,omitempty" name:"Components"`
}

type Project struct {
	// 项目ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 项目Logo
	// 注意：此字段可能返回 null，表示取不到有效值。
	Logo *string `json:"Logo,omitnil,omitempty" name:"Logo"`

	// 项目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// logo底色
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColorCode *string `json:"ColorCode,omitnil,omitempty" name:"ColorCode"`

	// 创建人
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedUser *string `json:"CreatedUser,omitnil,omitempty" name:"CreatedUser"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 成员个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemberCount *int64 `json:"MemberCount,omitnil,omitempty" name:"MemberCount"`

	// 页面个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageCount *int64 `json:"PageCount,omitnil,omitempty" name:"PageCount"`

	// 最后修改报表、简报名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastModifyName *string `json:"LastModifyName,omitnil,omitempty" name:"LastModifyName"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	Apply *bool `json:"Apply,omitnil,omitempty" name:"Apply"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedUser *string `json:"UpdatedUser,omitnil,omitempty" name:"UpdatedUser"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	CorpId *string `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mark *string `json:"Mark,omitnil,omitempty" name:"Mark"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	Seed *string `json:"Seed,omitnil,omitempty" name:"Seed"`

	// 项目内权限列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthList []*string `json:"AuthList,omitnil,omitempty" name:"AuthList"`

	// 默认看板
	// 注意：此字段可能返回 null，表示取不到有效值。
	PanelScope *string `json:"PanelScope,omitnil,omitempty" name:"PanelScope"`

	// 是否被托管
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsExternalManage *bool `json:"IsExternalManage,omitnil,omitempty" name:"IsExternalManage"`

	// 托管平台名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ManagePlatform *string `json:"ManagePlatform,omitnil,omitempty" name:"ManagePlatform"`

	// 定制化参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigList []*ProjectConfigList `json:"ConfigList,omitnil,omitempty" name:"ConfigList"`

	// 创建人
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedUserName *string `json:"CreatedUserName,omitnil,omitempty" name:"CreatedUserName"`

	// 所属人id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Owner *string `json:"Owner,omitnil,omitempty" name:"Owner"`

	// 所属人
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerName *string `json:"OwnerName,omitnil,omitempty" name:"OwnerName"`

	// 仪表盘页面数
	// 注意：此字段可能返回 null，表示取不到有效值。
	NormalCount *int64 `json:"NormalCount,omitnil,omitempty" name:"NormalCount"`

	// 自由画布页面数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FreeCount *int64 `json:"FreeCount,omitnil,omitempty" name:"FreeCount"`

	// 即席分析页面数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdhocCount *int64 `json:"AdhocCount,omitnil,omitempty" name:"AdhocCount"`

	// 简报页面数
	// 注意：此字段可能返回 null，表示取不到有效值。
	BriefingCount *int64 `json:"BriefingCount,omitnil,omitempty" name:"BriefingCount"`
}

type ProjectConfigList struct {
	// 模块组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModuleGroup *string `json:"ModuleGroup,omitnil,omitempty" name:"ModuleGroup"`

	// 内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Components []*ProjectConfigResult `json:"Components,omitnil,omitempty" name:"Components"`
}

type ProjectConfigResult struct {
	// <p>配置名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModuleId *string `json:"ModuleId,omitnil,omitempty" name:"ModuleId"`

	// <p>配置方式</p><p>枚举值：</p><ul><li>usable： 可用</li><li>visible： 可见</li><li>disabled： 不可用</li><li>hidden： 隐藏</li></ul><p>默认值：disabled</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncludeType *string `json:"IncludeType,omitnil,omitempty" name:"IncludeType"`

	// <p>额外参数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Params *string `json:"Params,omitnil,omitempty" name:"Params"`
}

type ProjectListData struct {
	// 数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*Project `json:"List,omitnil,omitempty" name:"List"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPages *uint64 `json:"TotalPages,omitnil,omitempty" name:"TotalPages"`
}

// Predefined struct for user
type QueryUserGroupMemberRequestParams struct {
	// <p>用户组id集合</p>
	GroupIds []*int64 `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// <p>搜索关键字</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>分页大小</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>分页页码</p>
	PageNo *uint64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// <p>是否需要分页</p>
	AllPage *bool `json:"AllPage,omitnil,omitempty" name:"AllPage"`

	// <p>用户组节点信息</p>
	Nodes []*UserGroupTreeNodeDTO `json:"Nodes,omitnil,omitempty" name:"Nodes"`

	// <p>标签值</p>
	TagValueList []*ResourceTagValue `json:"TagValueList,omitnil,omitempty" name:"TagValueList"`

	// <p>需要过滤的用户组</p>
	FilterGroupIds []*int64 `json:"FilterGroupIds,omitnil,omitempty" name:"FilterGroupIds"`
}

type QueryUserGroupMemberRequest struct {
	*tchttp.BaseRequest
	
	// <p>用户组id集合</p>
	GroupIds []*int64 `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// <p>搜索关键字</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>分页大小</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>分页页码</p>
	PageNo *uint64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// <p>是否需要分页</p>
	AllPage *bool `json:"AllPage,omitnil,omitempty" name:"AllPage"`

	// <p>用户组节点信息</p>
	Nodes []*UserGroupTreeNodeDTO `json:"Nodes,omitnil,omitempty" name:"Nodes"`

	// <p>标签值</p>
	TagValueList []*ResourceTagValue `json:"TagValueList,omitnil,omitempty" name:"TagValueList"`

	// <p>需要过滤的用户组</p>
	FilterGroupIds []*int64 `json:"FilterGroupIds,omitnil,omitempty" name:"FilterGroupIds"`
}

func (r *QueryUserGroupMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryUserGroupMemberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupIds")
	delete(f, "Keyword")
	delete(f, "PageSize")
	delete(f, "PageNo")
	delete(f, "AllPage")
	delete(f, "Nodes")
	delete(f, "TagValueList")
	delete(f, "FilterGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryUserGroupMemberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryUserGroupMemberResponseParams struct {
	// 自定义错误信息对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// <p>额外信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// <p>结果信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// <p>数据</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *DescribeUserGroupMemberPageListContainer `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryUserGroupMemberResponse struct {
	*tchttp.BaseResponse
	Response *QueryUserGroupMemberResponseParams `json:"Response"`
}

func (r *QueryUserGroupMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryUserGroupMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RankInfo struct {
	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// `
	// 注意：此字段可能返回 null，表示取不到有效值。
	RulerInfo *string `json:"RulerInfo,omitnil,omitempty" name:"RulerInfo"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleId *int64 `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleType *string `json:"RoleType,omitnil,omitempty" name:"RoleType"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableId *int64 `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 行列权限配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	RowColumnConfigList []*RowColumnConfig `json:"RowColumnConfigList,omitnil,omitempty" name:"RowColumnConfigList"`
}

type ResourceDTO struct {
	// 资源名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 是否启用
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceValue *bool `json:"ResourceValue,omitnil,omitempty" name:"ResourceValue"`

	// 能否变更
	// 注意：此字段可能返回 null，表示取不到有效值。
	CanChange *bool `json:"CanChange,omitnil,omitempty" name:"CanChange"`

	// 提示
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tips *string `json:"Tips,omitnil,omitempty" name:"Tips"`
}

type ResourceItem struct {
	// 资源名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 资源值
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceValue *bool `json:"ResourceValue,omitnil,omitempty" name:"ResourceValue"`

	// 是否可变更
	// 注意：此字段可能返回 null，表示取不到有效值。
	CanChange *bool `json:"CanChange,omitnil,omitempty" name:"CanChange"`

	// 提示信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tips *string `json:"Tips,omitnil,omitempty" name:"Tips"`
}

type ResourceListDTO struct {
	// 资源id
	// 注意：此字段可能返回 null，表示取不到有效值。
	EntityId *int64 `json:"EntityId,omitnil,omitempty" name:"EntityId"`

	// 资源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeType *int64 `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// 资源权限
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceList []*ResourceItem `json:"ResourceList,omitnil,omitempty" name:"ResourceList"`
}

type ResourceTagValue struct {
	// 标签id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 标签名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 标签值列表
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type Role struct {
	// 角色ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 角色名称
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 企业ID
	CorpId *string `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 角色类型
	RoleType *string `json:"RoleType,omitnil,omitempty" name:"RoleType"`

	// 范围
	Scope *string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 创建人
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedUser *string `json:"CreatedUser,omitnil,omitempty" name:"CreatedUser"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// 更新人
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedUser *string `json:"UpdatedUser,omitnil,omitempty" name:"UpdatedUser"`

	// 是否为全局角色（0 不是， 1 是）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScopeType *int64 `json:"ScopeType,omitnil,omitempty" name:"ScopeType"`

	// 是否可被选
	// 注意：此字段可能返回 null，表示取不到有效值。
	CanChoose *bool `json:"CanChoose,omitnil,omitempty" name:"CanChoose"`

	// 角色key
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModuleCollection *string `json:"ModuleCollection,omitnil,omitempty" name:"ModuleCollection"`
}

type RowColumnConfig struct {
	// 行列权限规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	RulerInfo *string `json:"RulerInfo,omitnil,omitempty" name:"RulerInfo"`

	// 标签值列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagValueList []*RowColumnTagValue `json:"TagValueList,omitnil,omitempty" name:"TagValueList"`
}

type RowColumnStatus struct {
	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableId *int64 `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	OpenStatus *string `json:"OpenStatus,omitnil,omitempty" name:"OpenStatus"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleType *string `json:"RoleType,omitnil,omitempty" name:"RoleType"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleId *int64 `json:"RoleId,omitnil,omitempty" name:"RoleId"`
}

type RowColumnTagValue struct {
	// 标签id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 标签名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 标签值列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type TableColumn struct {
	// 列名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 列的别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	AliasName *string `json:"AliasName,omitnil,omitempty" name:"AliasName"`

	// 列的类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// 段类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	FieldType *string `json:"FieldType,omitnil,omitempty" name:"FieldType"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mark *string `json:"Mark,omitnil,omitempty" name:"Mark"`

	// excel名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExcelName *string `json:"ExcelName,omitnil,omitempty" name:"ExcelName"`

	// 关联的字典表Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DictId *int64 `json:"DictId,omitnil,omitempty" name:"DictId"`

	// 关联的字典表表名
	// 注意：此字段可能返回 null，表示取不到有效值。
	DictName *string `json:"DictName,omitnil,omitempty" name:"DictName"`

	// 多表关联新增字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableNodeId *string `json:"TableNodeId,omitnil,omitempty" name:"TableNodeId"`

	// 字段所属的表名
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 用户设置的带格式的目标复杂格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	FieldComplexType *string `json:"FieldComplexType,omitnil,omitempty" name:"FieldComplexType"`

	// 格式规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	FormatRule *string `json:"FormatRule,omitnil,omitempty" name:"FormatRule"`

	// 数据字段是否过滤空值
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsFilter *bool `json:"IsFilter,omitnil,omitempty" name:"IsFilter"`

	// 计算字段类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	CalcType *string `json:"CalcType,omitnil,omitempty" name:"CalcType"`

	// 计算字段的公式内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	CalcFormula *string `json:"CalcFormula,omitnil,omitempty" name:"CalcFormula"`

	// 计算字段的中文公式内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	CalcDesc *string `json:"CalcDesc,omitnil,omitempty" name:"CalcDesc"`

	// Api数据源json路径名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	JsonPathName *string `json:"JsonPathName,omitnil,omitempty" name:"JsonPathName"`

	// 地理类型标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	Granularity *string `json:"Granularity,omitnil,omitempty" name:"Granularity"`

	// 自定义地图Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	GeoJsonId *uint64 `json:"GeoJsonId,omitnil,omitempty" name:"GeoJsonId"`

	// 空值展示样式配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	EmptyValueConfig *EmptyValueConfig `json:"EmptyValueConfig,omitnil,omitempty" name:"EmptyValueConfig"`

	// 原列名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbFieldName *string `json:"DbFieldName,omitnil,omitempty" name:"DbFieldName"`

	// 是否是复制字段操作
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsCopyOperation *bool `json:"IsCopyOperation,omitnil,omitempty" name:"IsCopyOperation"`

	// 是否从普通字段复制
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsCopyFromNormal *bool `json:"IsCopyFromNormal,omitnil,omitempty" name:"IsCopyFromNormal"`
}

type TableColumnListData struct {
	// 表中的列的列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*TableColumn `json:"List,omitnil,omitempty" name:"List"`

	// 异步事务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranId *string `json:"TranId,omitnil,omitempty" name:"TranId"`

	// 异步事务状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranStatus *int64 `json:"TranStatus,omitnil,omitempty" name:"TranStatus"`
}

type TableField struct {
	// db里的字段column名
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// bi展示名
	AliasName *string `json:"AliasName,omitnil,omitempty" name:"AliasName"`

	// db里的字段类型
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// BI归类后的抽象字段类型,比如字符串,数字,时间
	FieldType *string `json:"FieldType,omitnil,omitempty" name:"FieldType"`

	// 字段组合计算公式后生成的复杂明细类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	FieldComplexType *string `json:"FieldComplexType,omitnil,omitempty" name:"FieldComplexType"`

	// 字段描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mark *string `json:"Mark,omitnil,omitempty" name:"Mark"`

	// 字段计算公式
	// 注意：此字段可能返回 null，表示取不到有效值。
	FormatRule *string `json:"FormatRule,omitnil,omitempty" name:"FormatRule"`

	// 数据字段是否过滤空值
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsFilter *bool `json:"IsFilter,omitnil,omitempty" name:"IsFilter"`

	// 计算字段类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	CalcType *string `json:"CalcType,omitnil,omitempty" name:"CalcType"`

	// 计算字段的公式内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	CalcFormula *string `json:"CalcFormula,omitnil,omitempty" name:"CalcFormula"`

	// 计算字段的中文公式内容, 前端展示使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	CalcDesc *string `json:"CalcDesc,omitnil,omitempty" name:"CalcDesc"`

	// 关联字典表id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DictId *uint64 `json:"DictId,omitnil,omitempty" name:"DictId"`

	// 关联字典表名
	// 注意：此字段可能返回 null，表示取不到有效值。
	DictName *string `json:"DictName,omitnil,omitempty" name:"DictName"`

	// 非必填, 多表关联新增字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableNodeId *string `json:"TableNodeId,omitnil,omitempty" name:"TableNodeId"`

	// excel
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExcelName *string `json:"ExcelName,omitnil,omitempty" name:"ExcelName"`

	// 非必填,多表关联新增字段,字段所属的表名
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// api数据源路径名
	// 注意：此字段可能返回 null，表示取不到有效值。
	JsonPathName *string `json:"JsonPathName,omitnil,omitempty" name:"JsonPathName"`

	// 地理字段标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	Granularity *string `json:"Granularity,omitnil,omitempty" name:"Granularity"`

	// 地图id
	// 注意：此字段可能返回 null，表示取不到有效值。
	GeoJsonId *uint64 `json:"GeoJsonId,omitnil,omitempty" name:"GeoJsonId"`

	// 空值展示样式配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	EmptyValueConfig *EmptyValueConfig `json:"EmptyValueConfig,omitnil,omitempty" name:"EmptyValueConfig"`

	// 原列名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbFieldName *string `json:"DbFieldName,omitnil,omitempty" name:"DbFieldName"`

	// 是否是复制字段操作
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsCopyOperation *bool `json:"IsCopyOperation,omitnil,omitempty" name:"IsCopyOperation"`

	// 是否从普通字段复制
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsCopyFromNormal *bool `json:"IsCopyFromNormal,omitnil,omitempty" name:"IsCopyFromNormal"`
}

type UserGroupDTO struct {
	// id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 用户组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 父节点id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentId *int64 `json:"ParentId,omitnil,omitempty" name:"ParentId"`

	// 是否为默认
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDefault *int64 `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`

	// 管理员用户id
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdminUserId *string `json:"AdminUserId,omitnil,omitempty" name:"AdminUserId"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 定位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Location *int64 `json:"Location,omitnil,omitempty" name:"Location"`
}

type UserGroupMemberVO struct {
	// 用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 用户id
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 创建人
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedBy *string `json:"CreatedBy,omitnil,omitempty" name:"CreatedBy"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedOn *string `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`
}

type UserGroupPageTreeVO struct {
	// 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*UserGroupTreeVO `json:"List,omitnil,omitempty" name:"List"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 页数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPages *int64 `json:"TotalPages,omitnil,omitempty" name:"TotalPages"`
}

type UserGroupTreeNodeDTO struct {
	// 用户组id
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 是否查询子节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueryChildNodes *bool `json:"QueryChildNodes,omitnil,omitempty" name:"QueryChildNodes"`
}

type UserGroupTreeVO struct {
	// 当前实体ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 用户组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 父id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentId *int64 `json:"ParentId,omitnil,omitempty" name:"ParentId"`

	// 父节点名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentName *string `json:"ParentName,omitnil,omitempty" name:"ParentName"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDefault *int64 `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`

	// 管理员账号id
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdminUserId *string `json:"AdminUserId,omitnil,omitempty" name:"AdminUserId"`

	// 用户集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserList []*UserVO `json:"UserList,omitnil,omitempty" name:"UserList"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 排序
	// 注意：此字段可能返回 null，表示取不到有效值。
	Location *int64 `json:"Location,omitnil,omitempty" name:"Location"`

	// 孩子节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	Children []*UserGroupTreeVO `json:"Children,omitnil,omitempty" name:"Children"`

	// 是否有孩子节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	HasChildren *bool `json:"HasChildren,omitnil,omitempty" name:"HasChildren"`

	// 资源集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceList []*ResourceDTO `json:"ResourceList,omitnil,omitempty" name:"ResourceList"`
}

type UserGroupUpdateDTO struct {
	// 组管理员
	AdminUserId *string `json:"AdminUserId,omitnil,omitempty" name:"AdminUserId"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 用户组名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 排序位置
	Location *int64 `json:"Location,omitnil,omitempty" name:"Location"`

	// 父节点id
	ParentId *int64 `json:"ParentId,omitnil,omitempty" name:"ParentId"`

	// 用户组id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 父节点名称
	ParentName *string `json:"ParentName,omitnil,omitempty" name:"ParentName"`
}

type UserGroupUserInfoVO struct {
	// 用户id
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`
}

type UserGroupVO struct {
	// 用户组id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 用户组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 所属用户组id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentId *int64 `json:"ParentId,omitnil,omitempty" name:"ParentId"`

	// 所属用户组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentName *string `json:"ParentName,omitnil,omitempty" name:"ParentName"`

	// 是否默认用户组
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDefault *int64 `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`

	// 用户组管理员
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdminUserId *string `json:"AdminUserId,omitnil,omitempty" name:"AdminUserId"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 排序位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Location *int64 `json:"Location,omitnil,omitempty" name:"Location"`

	// 用户信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserList []*UserGroupUserInfoVO `json:"UserList,omitnil,omitempty" name:"UserList"`
}

type UserIdAndUserName struct {
	// <p>用户ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>用户名</p>
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// <p>企业ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CorpId *string `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// <p>电子邮箱</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// <p>最后一次登录时间</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastLogin *string `json:"LastLogin,omitnil,omitempty" name:"LastLogin"`

	// <p>用户状态</p><p>枚举值：</p><ul><li>1： 启用</li><li>0： 停用</li></ul><p>默认值：1</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>首次登录是否修改密码</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstModify *int64 `json:"FirstModify,omitnil,omitempty" name:"FirstModify"`

	// <p>手机号码</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// <p>手机区号</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AreaCode *string `json:"AreaCode,omitnil,omitempty" name:"AreaCode"`

	// <p>创建人</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedUser *string `json:"CreatedUser,omitnil,omitempty" name:"CreatedUser"`

	// <p>创建时间</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// <p>修改人</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedUser *string `json:"UpdatedUser,omitnil,omitempty" name:"UpdatedUser"`

	// <p>更改时间</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// <p>系统全局角色</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	GlobalUserName *string `json:"GlobalUserName,omitnil,omitempty" name:"GlobalUserName"`

	// <p>系统全局角色编码</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	GlobalUserCode *string `json:"GlobalUserCode,omitnil,omitempty" name:"GlobalUserCode"`

	// <p>手机号</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// <p>1</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>1</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppUserId *string `json:"AppUserId,omitnil,omitempty" name:"AppUserId"`

	// <p>1</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppUserAliasName *string `json:"AppUserAliasName,omitnil,omitempty" name:"AppUserAliasName"`

	// <p>1</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppUserName *string `json:"AppUserName,omitnil,omitempty" name:"AppUserName"`

	// <p>1</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	InValidateAppRange *bool `json:"InValidateAppRange,omitnil,omitempty" name:"InValidateAppRange"`

	// <p>-1 免激活  0 未激活  1 已激活 空代表待绑定</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	EmailActivationStatus *int64 `json:"EmailActivationStatus,omitnil,omitempty" name:"EmailActivationStatus"`

	// <p>1</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`
}

type UserInfo struct {
	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 邮箱
	// 注意：此字段可能返回 null，表示取不到有效值。
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 手机号
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// 手机号区号
	// 注意：此字段可能返回 null，表示取不到有效值。
	AreaCode *string `json:"AreaCode,omitnil,omitempty" name:"AreaCode"`

	// 企微账号id
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppUserId *string `json:"AppUserId,omitnil,omitempty" name:"AppUserId"`

	// 企微账号名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppUserName *string `json:"AppUserName,omitnil,omitempty" name:"AppUserName"`
}

type UserResourceDTO struct {
	// 企业id
	// 注意：此字段可能返回 null，表示取不到有效值。
	CorpId *string `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 用户id
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 资源列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceList []*ResourceItem `json:"ResourceList,omitnil,omitempty" name:"ResourceList"`
}

type UserRoleListData struct {
	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 总页数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPages *int64 `json:"TotalPages,omitnil,omitempty" name:"TotalPages"`

	// 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*UserRoleListDataUserRoleInfo `json:"List,omitnil,omitempty" name:"List"`
}

type UserRoleListDataRoleInfo struct {
	// 角色名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 角色ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleId *int64 `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 项目ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 是否为全局角色（0 不是 1 是）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScopeType *int64 `json:"ScopeType,omitnil,omitempty" name:"ScopeType"`

	// 角色key
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModuleCollection *string `json:"ModuleCollection,omitnil,omitempty" name:"ModuleCollection"`
}

type UserRoleListDataUserRoleInfo struct {
	// 业务ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 角色列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleList []*UserRoleListDataRoleInfo `json:"RoleList,omitnil,omitempty" name:"RoleList"`

	// 角色ID列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleIdList []*uint64 `json:"RoleIdList,omitnil,omitempty" name:"RoleIdList"`

	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 企业ID
	CorpId *string `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 邮箱
	// 注意：此字段可能返回 null，表示取不到有效值。
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 创建人
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedUser *string `json:"CreatedUser,omitnil,omitempty" name:"CreatedUser"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 更新人
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedUser *string `json:"UpdatedUser,omitnil,omitempty" name:"UpdatedUser"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// 最后一次登录时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastLogin *string `json:"LastLogin,omitnil,omitempty" name:"LastLogin"`

	// 账号状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 手机号码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// 手机号区号
	// 注意：此字段可能返回 null，表示取不到有效值。
	AreaCode *string `json:"AreaCode,omitnil,omitempty" name:"AreaCode"`

	// 是否为主账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	RootAccount *bool `json:"RootAccount,omitnil,omitempty" name:"RootAccount"`

	// 是否为企业管理员
	// 注意：此字段可能返回 null，表示取不到有效值。
	CorpAdmin *bool `json:"CorpAdmin,omitnil,omitempty" name:"CorpAdmin"`

	// 企微用户id
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppUserId *string `json:"AppUserId,omitnil,omitempty" name:"AppUserId"`

	// 昵称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppUserAliasName *string `json:"AppUserAliasName,omitnil,omitempty" name:"AppUserAliasName"`

	// 应用用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppUserName *string `json:"AppUserName,omitnil,omitempty" name:"AppUserName"`

	// 是否在可见范围内
	// 注意：此字段可能返回 null，表示取不到有效值。
	InValidateAppRange *bool `json:"InValidateAppRange,omitnil,omitempty" name:"InValidateAppRange"`

	// 用户openid
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppOpenUserId *string `json:"AppOpenUserId,omitnil,omitempty" name:"AppOpenUserId"`

	// 邮箱激活状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	EmailActivationStatus *int64 `json:"EmailActivationStatus,omitnil,omitempty" name:"EmailActivationStatus"`

	// 用户组信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserGroupList []*UserGroupDTO `json:"UserGroupList,omitnil,omitempty" name:"UserGroupList"`
}

type UserTagInfo struct {
	// 标签ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 标签名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 是否被托管
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsExternalManage *bool `json:"IsExternalManage,omitnil,omitempty" name:"IsExternalManage"`

	// 标签托管平台
	// 注意：此字段可能返回 null，表示取不到有效值。
	ManagePlatform *string `json:"ManagePlatform,omitnil,omitempty" name:"ManagePlatform"`

	// 导入类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImportType *string `json:"ImportType,omitnil,omitempty" name:"ImportType"`
}

type UserVO struct {
	// u1
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// zhang
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`
}

type WidgetListVO struct {
	// uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	CorpId *string `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 项目id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 页面id
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageId *string `json:"PageId,omitnil,omitempty" name:"PageId"`

	// 组件数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	WidgetList []*WidgetVO `json:"WidgetList,omitnil,omitempty" name:"WidgetList"`
}

type WidgetVO struct {
	// 组件Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	WidgetId *string `json:"WidgetId,omitnil,omitempty" name:"WidgetId"`

	// 组件name
	// 注意：此字段可能返回 null，表示取不到有效值。
	WidgetName *string `json:"WidgetName,omitnil,omitempty" name:"WidgetName"`
}