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

package v20260715

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type AddConsoleUsersRequestParams struct {
	// <p>用户 UIN 列表，单次最多100个</p>
	UserUins []*string `json:"UserUins,omitnil,omitempty" name:"UserUins"`

	// <p>角色 ID 列表</p><p>枚举值：</p><ul><li>2001： 控制台管理员</li><li>2002： 控制台成员</li></ul><p>可通过 ListConsoleRoles 接口获取</p>
	RoleIds []*string `json:"RoleIds,omitnil,omitempty" name:"RoleIds"`
}

type AddConsoleUsersRequest struct {
	*tchttp.BaseRequest
	
	// <p>用户 UIN 列表，单次最多100个</p>
	UserUins []*string `json:"UserUins,omitnil,omitempty" name:"UserUins"`

	// <p>角色 ID 列表</p><p>枚举值：</p><ul><li>2001： 控制台管理员</li><li>2002： 控制台成员</li></ul><p>可通过 ListConsoleRoles 接口获取</p>
	RoleIds []*string `json:"RoleIds,omitnil,omitempty" name:"RoleIds"`
}

func (r *AddConsoleUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddConsoleUsersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserUins")
	delete(f, "RoleIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddConsoleUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddConsoleUsersResponseParams struct {
	// <p>返回结果</p>
	Data *AddConsoleUsersRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddConsoleUsersResponse struct {
	*tchttp.BaseResponse
	Response *AddConsoleUsersResponseParams `json:"Response"`
}

func (r *AddConsoleUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddConsoleUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddConsoleUsersRsp struct {
	// 操作是否成功
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type AddWorkspaceGroupsRequestParams struct {
	// <p>用户组ID，可通过 ListConsoleGroups 接口获取</p>
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// <p>工作空间ID</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>角色ID，可通过 ListWorkspaceRoles 接口获取</p>
	RoleIds []*string `json:"RoleIds,omitnil,omitempty" name:"RoleIds"`
}

type AddWorkspaceGroupsRequest struct {
	*tchttp.BaseRequest
	
	// <p>用户组ID，可通过 ListConsoleGroups 接口获取</p>
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// <p>工作空间ID</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>角色ID，可通过 ListWorkspaceRoles 接口获取</p>
	RoleIds []*string `json:"RoleIds,omitnil,omitempty" name:"RoleIds"`
}

func (r *AddWorkspaceGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddWorkspaceGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupIds")
	delete(f, "WorkspaceId")
	delete(f, "RoleIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddWorkspaceGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddWorkspaceGroupsResponseParams struct {
	// <p>添加结果</p>
	Data *PlatformStatusRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddWorkspaceGroupsResponse struct {
	*tchttp.BaseResponse
	Response *AddWorkspaceGroupsResponseParams `json:"Response"`
}

func (r *AddWorkspaceGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddWorkspaceGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddWorkspaceUsersRequestParams struct {
	// <p>用户ID</p>
	UserUins []*string `json:"UserUins,omitnil,omitempty" name:"UserUins"`

	// <p>工作空间ID</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>角色ID，可通过 ListWorkspaceRoles 接口获取</p>
	RoleIds []*string `json:"RoleIds,omitnil,omitempty" name:"RoleIds"`
}

type AddWorkspaceUsersRequest struct {
	*tchttp.BaseRequest
	
	// <p>用户ID</p>
	UserUins []*string `json:"UserUins,omitnil,omitempty" name:"UserUins"`

	// <p>工作空间ID</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>角色ID，可通过 ListWorkspaceRoles 接口获取</p>
	RoleIds []*string `json:"RoleIds,omitnil,omitempty" name:"RoleIds"`
}

func (r *AddWorkspaceUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddWorkspaceUsersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserUins")
	delete(f, "WorkspaceId")
	delete(f, "RoleIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddWorkspaceUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddWorkspaceUsersResponseParams struct {
	// <p>结果</p>
	Data *PlatformStatusRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddWorkspaceUsersResponse struct {
	*tchttp.BaseResponse
	Response *AddWorkspaceUsersResponseParams `json:"Response"`
}

func (r *AddWorkspaceUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddWorkspaceUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AdvancedDependencyConfig struct {
	// 逻辑运算符号OR / AND
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 任务运行条件规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Conditions []*TaskRunConditionRule `json:"Conditions,omitnil,omitempty" name:"Conditions"`
}

type AdvancedParameter struct {
	// 参数名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamKey *string `json:"ParamKey,omitnil,omitempty" name:"ParamKey"`

	// 参数值
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamValue *string `json:"ParamValue,omitnil,omitempty" name:"ParamValue"`
}

type AlarmBrief struct {
	// 告警 ID，创建时无需传入，由服务端生成
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmId *string `json:"AlarmId,omitnil,omitempty" name:"AlarmId"`

	// 告警的监控对象类型，如工作流、任务等，当前支持 1. WORKFLOW 2. TASK
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmMonitorType *string `json:"AlarmMonitorType,omitnil,omitempty" name:"AlarmMonitorType"`

	// 告警组，最多 50 个
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmGroups []*AlarmGroup `json:"AlarmGroups,omitnil,omitempty" name:"AlarmGroups"`

	// 被跳过时免打扰，默认值 false
	// 注意：此字段可能返回 null，表示取不到有效值。
	DoNotDisturbWhenSkipped *bool `json:"DoNotDisturbWhenSkipped,omitnil,omitempty" name:"DoNotDisturbWhenSkipped"`

	// 被手动终止时免打扰，默认值 false
	// 注意：此字段可能返回 null，表示取不到有效值。
	DoNotDisturbWhenManuallyTerminated *bool `json:"DoNotDisturbWhenManuallyTerminated,omitnil,omitempty" name:"DoNotDisturbWhenManuallyTerminated"`

	// 最后一次重试前免打扰，默认值 false
	// 注意：此字段可能返回 null，表示取不到有效值。
	DoNotDisturbUntilTheLastRetry *bool `json:"DoNotDisturbUntilTheLastRetry,omitnil,omitempty" name:"DoNotDisturbUntilTheLastRetry"`
}

type AlarmGroup struct {
	// <p>通知渠道ID，可通过基础平台通知渠道相关接口获取</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// <p>通知渠道名称，可以是用户组名称或邮箱地址</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`

	// <p>是否启用邮件渠道，默认值：false</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsEmailChannel *bool `json:"IsEmailChannel,omitnil,omitempty" name:"IsEmailChannel"`

	// <p>告警条件列表。取值：<br>START：启动<br>SUCCESS：成功<br>FAILURE：失败<br>MONITOR_INDICATOR_ALARM：监控指标告警</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmConditions []*string `json:"AlarmConditions,omitnil,omitempty" name:"AlarmConditions"`

	// <p>通知渠道类型。取值：0 未指定，1 Email，2 Webhook，3 Teams，4 Slack</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelType *int64 `json:"ChannelType,omitnil,omitempty" name:"ChannelType"`
}

type AsyncActionRsp struct {
	// 多个操作项的结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActionResults []*RunActionBrief `json:"ActionResults,omitnil,omitempty" name:"ActionResults"`
}

type AsyncOperation struct {
	// 是否异步执行；ZIP 解压创建时为 true
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAsync *bool `json:"IsAsync,omitnil,omitempty" name:"IsAsync"`

	// Workspace 持久化的异步作业 ID，用于查询作业进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 调用方生成的提交幂等与链路追踪标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperationId *string `json:"OperationId,omitnil,omitempty" name:"OperationId"`

	// 异步作业状态：0-未指定，1-已受理，2-解压中，3-回调处理中，4-成功，5-部分失败，6-失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type Audit struct {
	// 创建者。注意：此字段可能返回null，表示取不到有效值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 创建时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 最后修改者。注意：此字段可能返回null，表示取不到有效值
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastModifier *string `json:"LastModifier,omitnil,omitempty" name:"LastModifier"`

	// 最后修改时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastModifiedAt *string `json:"LastModifiedAt,omitnil,omitempty" name:"LastModifiedAt"`

	// 创建者名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatorName *string `json:"CreatorName,omitnil,omitempty" name:"CreatorName"`

	// 最后修改者名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastModifierName *string `json:"LastModifierName,omitnil,omitempty" name:"LastModifierName"`
}

type BucketPartitioning struct {
	// 分桶字段
	FiledNames []*string `json:"FiledNames,omitnil,omitempty" name:"FiledNames"`

	// 分桶数
	NumBuckets *string `json:"NumBuckets,omitnil,omitempty" name:"NumBuckets"`
}

type Catalog struct {
	// <p>数据目录唯一id</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>数据目录名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>数据目录类型</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>数据目录描述信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// <p>状态信息：注册中0，待测试1，连接成功2，连接失败3，删除中4，已删除5</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>数据目录属性</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Properties []*KVPair `json:"Properties,omitnil,omitempty" name:"Properties"`

	// <p>操作人</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// <p>责任人信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetaOwner *MetaOwner `json:"MetaOwner,omitnil,omitempty" name:"MetaOwner"`

	// <p>审计信息</p>
	Audit *Audit `json:"Audit,omitnil,omitempty" name:"Audit"`

	// <p>资产全局唯一ID</p>
	AssetGuid *string `json:"AssetGuid,omitnil,omitempty" name:"AssetGuid"`

	// <p>当前workspace权限级别，取值对应PermissionLevel，0-无权限，1-只读，2-读写</p>
	CurrentWorkspacePermissionLevel *int64 `json:"CurrentWorkspacePermissionLevel,omitnil,omitempty" name:"CurrentWorkspacePermissionLevel"`

	// <p>tccatalog属性中的tccatalog.identifier</p>
	Identifier *string `json:"Identifier,omitnil,omitempty" name:"Identifier"`

	// <p>当前用户对该实体的权限信息</p>
	PermissionDetail *PermissionDetail `json:"PermissionDetail,omitnil,omitempty" name:"PermissionDetail"`

	// <p>是否为外部数据目录（基于数据源连接创建的Catalog，如MYSQL、EMR_HIVE、TCHOUSE_D、POSTGRESQL、DLC等）。true:外部目录, false:内置目录（TABLE/MODEL/VOLUME）</p>
	IsExternal *bool `json:"IsExternal,omitnil,omitempty" name:"IsExternal"`

	// <p>标签信息列表。注意：仅当请求中 FetchTags=true 时返回，默认不返回</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*CommonTagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>数据目录来源</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CatalogSource *string `json:"CatalogSource,omitnil,omitempty" name:"CatalogSource"`

	// <p>连接侧数据目录名或数据库名（仅分析版三段式数据源有值，二段式数据源为空字符串）。Doris/StarRocks/TChouse-D 等 MySQL 内核三段式数据源：external catalog 名称；PostgreSQL/Tencent PG/GaussDB/Hologres 等 PG 内核三段式数据源：databaseName；MySQL 等二段式数据源：空字符串。注意：此字段仅分析版（CatalogSource=CONNECTION）有效，专业版（METALAKE）不填充。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConnectionCatalogName *string `json:"ConnectionCatalogName,omitnil,omitempty" name:"ConnectionCatalogName"`
}

type ColumnInfo struct {
	// 字段名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 字段类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 字段描述
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 字段设置
	FieldSetting *string `json:"FieldSetting,omitnil,omitempty" name:"FieldSetting"`

	// 是否为主键
	IsPrimaryKey *bool `json:"IsPrimaryKey,omitnil,omitempty" name:"IsPrimaryKey"`

	// 字段策略
	DataMaskStrategyInfo *DataMaskStrategyInfo `json:"DataMaskStrategyInfo,omitnil,omitempty" name:"DataMaskStrategyInfo"`

	// 字段级标签列表，仅当请求中FetchOption.FetchTags=true时返回，默认不返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*CommonTagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 字段关联的字典维度信息，仅当请求中FetchOption.FetchDictDimensions=true时返回，默认不返回；若该字段未关联任何字典维度则为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	DictDimension *DictDimensionInfo `json:"DictDimension,omitnil,omitempty" name:"DictDimension"`
}

type CommonFailItem struct {
	// <p>uin或者groupId</p>
	Item *string `json:"Item,omitnil,omitempty" name:"Item"`

	// <p>错误信息</p>
	FailReason *string `json:"FailReason,omitnil,omitempty" name:"FailReason"`
}

type CommonTagInfo struct {
	// 标签ID
	LabelId *string `json:"LabelId,omitnil,omitempty" name:"LabelId"`

	// 标签名称
	LabelName *string `json:"LabelName,omitnil,omitempty" name:"LabelName"`

	// 标签值ID，属性标签（LabelType=3）可为0
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelValueId *string `json:"LabelValueId,omitnil,omitempty" name:"LabelValueId"`

	// 标签值，脱敏标签（LabelType=4）时可为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelValue *string `json:"LabelValue,omitnil,omitempty" name:"LabelValue"`

	// 标签类型，取值参考LabelType枚举定义：1-治理标签，2-自定义标签，3-属性标签，4-脱敏标签
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 标签是否已删除。true表示该LabelId在meta_biz_label中查不到记录，标签已被物理删除；false（默认）表示标签仍存在
	Deleted *bool `json:"Deleted,omitnil,omitempty" name:"Deleted"`
}

type ConsoleGroupInfo struct {
	// 用户组 ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 用户组名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 角色列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Roles []*RoleBasicInfo `json:"Roles,omitnil,omitempty" name:"Roles"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 用户组下用户数量
	UserCount *int64 `json:"UserCount,omitnil,omitempty" name:"UserCount"`

	// 用户组类型。取值为枚举数值的字符串形式："0"=控制台系统类型（包含全部user）、"1"=控制台自定义类型、"2"=工作空间系统类型、"3"=工作空间自定义类型
	GroupType *string `json:"GroupType,omitnil,omitempty" name:"GroupType"`
}

type ConsoleGroupUserInfo struct {
	// 用户 UIN
	UserUin *string `json:"UserUin,omitnil,omitempty" name:"UserUin"`

	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 昵称
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type ConsoleRoleInfo struct {
	// 角色基本信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	BasicInfo *RoleBasicInfo `json:"BasicInfo,omitnil,omitempty" name:"BasicInfo"`

	// 角色元信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetaData *RoleMetaData `json:"MetaData,omitnil,omitempty" name:"MetaData"`

	// 角色权限
	// 注意：此字段可能返回 null，表示取不到有效值。
	Permissions []*RolePermission `json:"Permissions,omitnil,omitempty" name:"Permissions"`
}

type ConsoleUserInfo struct {
	// 用户 UIN
	UserUin *string `json:"UserUin,omitnil,omitempty" name:"UserUin"`

	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 昵称
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// 角色列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Roles []*RoleBasicInfo `json:"Roles,omitnil,omitempty" name:"Roles"`

	// 用户来源，group：用户组、user:用户
	UserSource *string `json:"UserSource,omitnil,omitempty" name:"UserSource"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 是否主账号
	IsOwner *bool `json:"IsOwner,omitnil,omitempty" name:"IsOwner"`

	// 0: 普通用户 1: entraId用户
	UserTag *int64 `json:"UserTag,omitnil,omitempty" name:"UserTag"`

	// 是否具有 admin 权限的子账号
	IsAdmin *bool `json:"IsAdmin,omitnil,omitempty" name:"IsAdmin"`
}

// Predefined struct for user
type CreateCatalogRequestParams struct {
	// catalog名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// catalog类型, 可选值TABLE、MODEL、VOLUME
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 工作空间唯一id
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 描述
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// connection 的 ID
	ConnectionId *string `json:"ConnectionId,omitnil,omitempty" name:"ConnectionId"`

	// 数据目录来源，可选（融合版新增字段），取值参考 CatalogSourceEnum：METALAKE（专业版）/ CONNECTION（分析版），不传时默认按 METALAKE 处理
	CatalogSource *string `json:"CatalogSource,omitnil,omitempty" name:"CatalogSource"`

	// 连接侧数据目录名或数据库名，三段式数据源必填，二段式数据源（如 MySQL）不填。Doris/StarRocks/TChouse-D 等 MySQL 内核三段式数据源填写 external catalog 名称；PostgreSQL/Tencent PG/GaussDB/Hologres 等 PG 内核三段式数据源填写 databaseName；MySQL 等二段式数据源不填。
	ConnectionCatalogName *string `json:"ConnectionCatalogName,omitnil,omitempty" name:"ConnectionCatalogName"`
}

type CreateCatalogRequest struct {
	*tchttp.BaseRequest
	
	// catalog名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// catalog类型, 可选值TABLE、MODEL、VOLUME
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 工作空间唯一id
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 描述
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// connection 的 ID
	ConnectionId *string `json:"ConnectionId,omitnil,omitempty" name:"ConnectionId"`

	// 数据目录来源，可选（融合版新增字段），取值参考 CatalogSourceEnum：METALAKE（专业版）/ CONNECTION（分析版），不传时默认按 METALAKE 处理
	CatalogSource *string `json:"CatalogSource,omitnil,omitempty" name:"CatalogSource"`

	// 连接侧数据目录名或数据库名，三段式数据源必填，二段式数据源（如 MySQL）不填。Doris/StarRocks/TChouse-D 等 MySQL 内核三段式数据源填写 external catalog 名称；PostgreSQL/Tencent PG/GaussDB/Hologres 等 PG 内核三段式数据源填写 databaseName；MySQL 等二段式数据源不填。
	ConnectionCatalogName *string `json:"ConnectionCatalogName,omitnil,omitempty" name:"ConnectionCatalogName"`
}

func (r *CreateCatalogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCatalogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "WorkspaceId")
	delete(f, "Comment")
	delete(f, "ConnectionId")
	delete(f, "CatalogSource")
	delete(f, "ConnectionCatalogName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCatalogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCatalogResponseParams struct {
	// 创建catalog响应
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *CreateCatalogRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCatalogResponse struct {
	*tchttp.BaseResponse
	Response *CreateCatalogResponseParams `json:"Response"`
}

func (r *CreateCatalogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCatalogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCatalogRsp struct {
	// 新创建的数据目录的id
	// 注意：此字段可能返回 null，表示取不到有效值。
	CatalogId *string `json:"CatalogId,omitnil,omitempty" name:"CatalogId"`
}

// Predefined struct for user
type CreateConsoleGroupRequestParams struct {
	// <p>用户组名称</p>
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// <p>用户组别名</p>
	GroupNickname *string `json:"GroupNickname,omitnil,omitempty" name:"GroupNickname"`

	// <p>用户组描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateConsoleGroupRequest struct {
	*tchttp.BaseRequest
	
	// <p>用户组名称</p>
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// <p>用户组别名</p>
	GroupNickname *string `json:"GroupNickname,omitnil,omitempty" name:"GroupNickname"`

	// <p>用户组描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateConsoleGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConsoleGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupName")
	delete(f, "GroupNickname")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateConsoleGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConsoleGroupResponseParams struct {
	// <p>返回结果</p>
	Data *CreateConsoleGroupRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateConsoleGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateConsoleGroupResponseParams `json:"Response"`
}

func (r *CreateConsoleGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConsoleGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateConsoleGroupRsp struct {
	// 创建成功的用户组 ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

// Predefined struct for user
type CreateFileRequestParams struct {
	// <p>工作空间 ID。来源：ListWorkspaces 接口返回的 WorkspaceId</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>文件名，含后缀，最长 255 字节。不能以 . 或 .. 开头/结尾，不能含空格与控制字符</p>
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// <p>父文件夹路径，以 / 开头、末尾不带 /，根目录传 /。来源：ListFiles 接口返回的 Path</p>
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// <p>文件类型。取值：FILE（普通文件/脚本）、NOTEBOOK_FILE（Notebook）、SQL_FILE（SQL文件）。对应 common/domain/entity.proto EntityType</p>
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// <p>文件运行配置</p>
	FileConfig *FileConfig `json:"FileConfig,omitnil,omitempty" name:"FileConfig"`

	// <p>绑定的 BundleId。来源：ListBundles 接口返回的 BundleId</p>
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// <p>绑定的 BundleInfo，JSON 字符串</p>
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`

	// <p>文件初始内容。不传则按FileType 生成默认内容</p>
	Storage *FileStorage `json:"Storage,omitnil,omitempty" name:"Storage"`

	// 是否将 Storage 中的 ZIP 文件异步解压创建，默认 false。true 时异步作业由 Workspace 负责全生命周期，响应仅通过 AsyncOperation 返回作业信息（FileId 为空）；作业进度查询由基础平台 WS 接口实现，不在本协议中定义。
	ExtractArchive *bool `json:"ExtractArchive,omitnil,omitempty" name:"ExtractArchive"`
}

type CreateFileRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间 ID。来源：ListWorkspaces 接口返回的 WorkspaceId</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>文件名，含后缀，最长 255 字节。不能以 . 或 .. 开头/结尾，不能含空格与控制字符</p>
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// <p>父文件夹路径，以 / 开头、末尾不带 /，根目录传 /。来源：ListFiles 接口返回的 Path</p>
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// <p>文件类型。取值：FILE（普通文件/脚本）、NOTEBOOK_FILE（Notebook）、SQL_FILE（SQL文件）。对应 common/domain/entity.proto EntityType</p>
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// <p>文件运行配置</p>
	FileConfig *FileConfig `json:"FileConfig,omitnil,omitempty" name:"FileConfig"`

	// <p>绑定的 BundleId。来源：ListBundles 接口返回的 BundleId</p>
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// <p>绑定的 BundleInfo，JSON 字符串</p>
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`

	// <p>文件初始内容。不传则按FileType 生成默认内容</p>
	Storage *FileStorage `json:"Storage,omitnil,omitempty" name:"Storage"`

	// 是否将 Storage 中的 ZIP 文件异步解压创建，默认 false。true 时异步作业由 Workspace 负责全生命周期，响应仅通过 AsyncOperation 返回作业信息（FileId 为空）；作业进度查询由基础平台 WS 接口实现，不在本协议中定义。
	ExtractArchive *bool `json:"ExtractArchive,omitnil,omitempty" name:"ExtractArchive"`
}

func (r *CreateFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "FileName")
	delete(f, "ParentFolderPath")
	delete(f, "FileType")
	delete(f, "FileConfig")
	delete(f, "BundleId")
	delete(f, "BundleInfo")
	delete(f, "Storage")
	delete(f, "ExtractArchive")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFileResponseParams struct {
	// <p>返回结果</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *FileInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateFileResponse struct {
	*tchttp.BaseResponse
	Response *CreateFileResponseParams `json:"Response"`
}

func (r *CreateFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFolderRequestParams struct {
	// <p>工作空间名称</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>文件夹名称</p>
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`

	// <p>文件夹类型</p><p>枚举值：</p><ul><li>FOLDER： 文件夹</li><li>GIT_FOLDER： git文件夹</li></ul>
	FolderType *string `json:"FolderType,omitnil,omitempty" name:"FolderType"`

	// <p>父节点</p>
	ParentFolder *FolderLocator `json:"ParentFolder,omitnil,omitempty" name:"ParentFolder"`

	// <p>git配置，FolderType=GIT_FOLDER 时必填</p>
	GitConfig *GitRepoConfig `json:"GitConfig,omitnil,omitempty" name:"GitConfig"`
}

type CreateFolderRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间名称</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>文件夹名称</p>
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`

	// <p>文件夹类型</p><p>枚举值：</p><ul><li>FOLDER： 文件夹</li><li>GIT_FOLDER： git文件夹</li></ul>
	FolderType *string `json:"FolderType,omitnil,omitempty" name:"FolderType"`

	// <p>父节点</p>
	ParentFolder *FolderLocator `json:"ParentFolder,omitnil,omitempty" name:"ParentFolder"`

	// <p>git配置，FolderType=GIT_FOLDER 时必填</p>
	GitConfig *GitRepoConfig `json:"GitConfig,omitnil,omitempty" name:"GitConfig"`
}

func (r *CreateFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "FolderName")
	delete(f, "FolderType")
	delete(f, "ParentFolder")
	delete(f, "GitConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFolderResponseParams struct {
	// <p>创建文件夹结果</p>
	Data *CreateFolderRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateFolderResponse struct {
	*tchttp.BaseResponse
	Response *CreateFolderResponseParams `json:"Response"`
}

func (r *CreateFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateFolderRsp struct {
	// <p>文件夹 ID</p>
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`
}

// Predefined struct for user
type CreateSchemaRequestParams struct {
	// catalog名称
	CatalogName *string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`

	// schema名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 调用时所在workspace唯一id
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 描述
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 数据源连接ID，可选（融合版新增字段）。分析版catalog不支持创建schema，传入非空时服务端返回ANA_CATALOG_NOT_SUPPORTED错误
	ConnectionId *string `json:"ConnectionId,omitnil,omitempty" name:"ConnectionId"`

	// 数据目录来源，可选（Setats融合新增字段）。取值：METALAKE（专业版tclake）/ CONNECTION（分析版Linked Catalog）/ SETATS（Setats元数据）。路由优先级：显式非空CatalogSource > 非空ConnectionId（视为CONNECTION）> 默认METALAKE；CatalogSource非空时ConnectionId将被忽略。
	CatalogSource *string `json:"CatalogSource,omitnil,omitempty" name:"CatalogSource"`
}

type CreateSchemaRequest struct {
	*tchttp.BaseRequest
	
	// catalog名称
	CatalogName *string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`

	// schema名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 调用时所在workspace唯一id
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 描述
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 数据源连接ID，可选（融合版新增字段）。分析版catalog不支持创建schema，传入非空时服务端返回ANA_CATALOG_NOT_SUPPORTED错误
	ConnectionId *string `json:"ConnectionId,omitnil,omitempty" name:"ConnectionId"`

	// 数据目录来源，可选（Setats融合新增字段）。取值：METALAKE（专业版tclake）/ CONNECTION（分析版Linked Catalog）/ SETATS（Setats元数据）。路由优先级：显式非空CatalogSource > 非空ConnectionId（视为CONNECTION）> 默认METALAKE；CatalogSource非空时ConnectionId将被忽略。
	CatalogSource *string `json:"CatalogSource,omitnil,omitempty" name:"CatalogSource"`
}

func (r *CreateSchemaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSchemaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CatalogName")
	delete(f, "Name")
	delete(f, "WorkspaceId")
	delete(f, "Comment")
	delete(f, "ConnectionId")
	delete(f, "CatalogSource")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSchemaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSchemaResponseParams struct {
	// schema信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *CreateSchemaRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSchemaResponse struct {
	*tchttp.BaseResponse
	Response *CreateSchemaResponseParams `json:"Response"`
}

func (r *CreateSchemaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSchemaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSchemaRsp struct {
	// schema信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Schema *Schema `json:"Schema,omitnil,omitempty" name:"Schema"`
}

// Predefined struct for user
type CreateWorkflowRequestParams struct {
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>工作流基本信息。必填，其中 WorkflowName 必填且工作空间内唯一</p>
	BaseInfo *WorkflowBaseInfo `json:"BaseInfo,omitnil,omitempty" name:"BaseInfo"`

	// <p>工作流调度配置</p>
	Trigger []*WorkflowTriggerConfiguration `json:"Trigger,omitnil,omitempty" name:"Trigger"`

	// <p>工作流参数列表</p>
	ParamList []*ParamInfo `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// <p>标签列表</p>
	LabelList []*LabelBrief `json:"LabelList,omitnil,omitempty" name:"LabelList"`

	// <p>工作流告警配置</p>
	Alarm *AlarmBrief `json:"Alarm,omitnil,omitempty" name:"Alarm"`

	// <p>监控指标配置。若告警条件中选择了监控告警，则本字段必填</p>
	MonitorMetric *MonitorMetricBrief `json:"MonitorMetric,omitnil,omitempty" name:"MonitorMetric"`

	// <p>工作流高级设置</p>
	AdvanceConfig *WorkflowAdvanceConfig `json:"AdvanceConfig,omitnil,omitempty" name:"AdvanceConfig"`

	// <p>工作流任务列表</p>
	TaskList []*WorkflowTask `json:"TaskList,omitnil,omitempty" name:"TaskList"`

	// <p>BundleId，可通过 Bundle 相关接口获取</p>
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// <p>Bundle信息</p>
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`

	// <p>Git配置ID，可通过 Git 配置相关接口获取</p>
	GitConfigId *string `json:"GitConfigId,omitnil,omitempty" name:"GitConfigId"`

	// <p>Git分支信息</p>
	GitBranch *string `json:"GitBranch,omitnil,omitempty" name:"GitBranch"`
}

type CreateWorkflowRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>工作流基本信息。必填，其中 WorkflowName 必填且工作空间内唯一</p>
	BaseInfo *WorkflowBaseInfo `json:"BaseInfo,omitnil,omitempty" name:"BaseInfo"`

	// <p>工作流调度配置</p>
	Trigger []*WorkflowTriggerConfiguration `json:"Trigger,omitnil,omitempty" name:"Trigger"`

	// <p>工作流参数列表</p>
	ParamList []*ParamInfo `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// <p>标签列表</p>
	LabelList []*LabelBrief `json:"LabelList,omitnil,omitempty" name:"LabelList"`

	// <p>工作流告警配置</p>
	Alarm *AlarmBrief `json:"Alarm,omitnil,omitempty" name:"Alarm"`

	// <p>监控指标配置。若告警条件中选择了监控告警，则本字段必填</p>
	MonitorMetric *MonitorMetricBrief `json:"MonitorMetric,omitnil,omitempty" name:"MonitorMetric"`

	// <p>工作流高级设置</p>
	AdvanceConfig *WorkflowAdvanceConfig `json:"AdvanceConfig,omitnil,omitempty" name:"AdvanceConfig"`

	// <p>工作流任务列表</p>
	TaskList []*WorkflowTask `json:"TaskList,omitnil,omitempty" name:"TaskList"`

	// <p>BundleId，可通过 Bundle 相关接口获取</p>
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// <p>Bundle信息</p>
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`

	// <p>Git配置ID，可通过 Git 配置相关接口获取</p>
	GitConfigId *string `json:"GitConfigId,omitnil,omitempty" name:"GitConfigId"`

	// <p>Git分支信息</p>
	GitBranch *string `json:"GitBranch,omitnil,omitempty" name:"GitBranch"`
}

func (r *CreateWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkflowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "BaseInfo")
	delete(f, "Trigger")
	delete(f, "ParamList")
	delete(f, "LabelList")
	delete(f, "Alarm")
	delete(f, "MonitorMetric")
	delete(f, "AdvanceConfig")
	delete(f, "TaskList")
	delete(f, "BundleId")
	delete(f, "BundleInfo")
	delete(f, "GitConfigId")
	delete(f, "GitBranch")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWorkflowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkflowResponseParams struct {
	// <p>创建工作流响应内容</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *CreateWorkflowRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *CreateWorkflowResponseParams `json:"Response"`
}

func (r *CreateWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkflowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateWorkflowRsp struct {
	// 工作流ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

// Predefined struct for user
type CreateWorkspaceRequestParams struct {
	// <p>工作空间名称，max_len=128</p>
	WorkspaceName *string `json:"WorkspaceName,omitnil,omitempty" name:"WorkspaceName"`

	// <p>工作空间地域（如 ap-guangzhou），max_len=64</p>
	WorkspaceRegion *string `json:"WorkspaceRegion,omitnil,omitempty" name:"WorkspaceRegion"`

	// <p>工作空间描述，max_len=300</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateWorkspaceRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间名称，max_len=128</p>
	WorkspaceName *string `json:"WorkspaceName,omitnil,omitempty" name:"WorkspaceName"`

	// <p>工作空间地域（如 ap-guangzhou），max_len=64</p>
	WorkspaceRegion *string `json:"WorkspaceRegion,omitnil,omitempty" name:"WorkspaceRegion"`

	// <p>工作空间描述，max_len=300</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
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
	delete(f, "WorkspaceName")
	delete(f, "WorkspaceRegion")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWorkspaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkspaceResponseParams struct {
	// <p>创建成功的工作空间ID</p>
	Data *CreateWorkspaceRsp `json:"Data,omitnil,omitempty" name:"Data"`

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
type CreateWorkspaceRoleRequestParams struct {
	// <p>工作空间id</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>角色基础信息</p>
	BasicInfo *RoleBasicInfo `json:"BasicInfo,omitnil,omitempty" name:"BasicInfo"`

	// <p>角色权限</p>
	Permissions []*RolePermission `json:"Permissions,omitnil,omitempty" name:"Permissions"`
}

type CreateWorkspaceRoleRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间id</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>角色基础信息</p>
	BasicInfo *RoleBasicInfo `json:"BasicInfo,omitnil,omitempty" name:"BasicInfo"`

	// <p>角色权限</p>
	Permissions []*RolePermission `json:"Permissions,omitnil,omitempty" name:"Permissions"`
}

func (r *CreateWorkspaceRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkspaceRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "BasicInfo")
	delete(f, "Permissions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWorkspaceRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkspaceRoleResponseParams struct {
	// <p>创建工作空间角色返回</p>
	Data *CreateWorkspaceRoleResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateWorkspaceRoleResponse struct {
	*tchttp.BaseResponse
	Response *CreateWorkspaceRoleResponseParams `json:"Response"`
}

func (r *CreateWorkspaceRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkspaceRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateWorkspaceRoleResult struct {
	// 角色id
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 创建工作空间角色是否成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

type CreateWorkspaceRsp struct {
	// 创建成功的工作空间ID
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`
}

type DataMaskGroup struct {
	// 工作组 ID，对应 WeData 用户组 ID，必填，字符串类型。示例值："17738042827013749"
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 脱敏方法（策略类型），必填，直接透传给 DLC。合法值：MASK_SHOW_FIRST_4（保留前四个字符）、MASK_SHOW_LAST_4（保留后四个字符）、MASK_HASH（哈希脱敏）、MASK_DATE_SHOW_YEAR（日期脱敏保留年份）、MASK_NULL（置为 NULL）、MASK_DEFAULT（默认值脱敏）
	StrategyType *string `json:"StrategyType,omitnil,omitempty" name:"StrategyType"`

	// 用户组名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`
}

type DataMaskLabelInfo struct {
	// 敏感标签 ID，对应 catalog-manager meta_biz_label.id
	LabelId *string `json:"LabelId,omitnil,omitempty" name:"LabelId"`

	// 敏感标签名称，对应 catalog-manager meta_biz_label.name
	LabelName *string `json:"LabelName,omitnil,omitempty" name:"LabelName"`

	// 标签类型，固定值 4（DESENSITIZATION_LABEL 脱敏标签）。参考 enum LabelType：1=业务标签、2=技术标签、3=分类标签、4=脱敏标签
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`
}

type DataMaskStrategyInfo struct {
	// 策略名称，必填，长度 1~128 字符，同一项目下唯一
	// 注意：此字段可能返回 null，表示取不到有效值。
	StrategyName *string `json:"StrategyName,omitnil,omitempty" name:"StrategyName"`

	// 用户组脱敏规则列表，必填，至少包含一条规则。每条规则描述某个用户组对应的脱敏方法
	// 注意：此字段可能返回 null，表示取不到有效值。
	Groups []*DataMaskGroup `json:"Groups,omitnil,omitempty" name:"Groups"`

	// 策略类型，必填，直接透传给 DLC。合法值：MASK_SHOW_FIRST_4（保留前四个字符）、MASK_SHOW_LAST_4（保留后四个字符）、MASK_HASH（哈希脱敏）、MASK_DATE_SHOW_YEAR（日期脱敏保留年份）、MASK_NULL（置为 NULL）、MASK_DEFAULT（默认值脱敏）
	// 注意：此字段可能返回 null，表示取不到有效值。
	StrategyType *string `json:"StrategyType,omitnil,omitempty" name:"StrategyType"`

	// 策略描述，可选，最大 512 字符
	// 注意：此字段可能返回 null，表示取不到有效值。
	StrategyDesc *string `json:"StrategyDesc,omitnil,omitempty" name:"StrategyDesc"`

	// 策略ID
	StrategyId *string `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`

	// 该脱敏策略当前绑定的敏感标签列表。使用场景：表详情弹窗回显“此策略关联了哪些敏感标签”，前端可 hover 展示。数据源：catalog-manager 域 meta_label_mask_policy_rel JOIN meta_biz_label，按 policy_id 反查。未关联任何敏感标签时为空数组。前置条件：仅当所在 ColumnInfo.DataMaskStrategyInfo 非空时才有查询意义。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Labels []*DataMaskLabelInfo `json:"Labels,omitnil,omitempty" name:"Labels"`
}

// Predefined struct for user
type DeleteCatalogRequestParams struct {
	// 数据目录名
	CatalogName *string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`

	// 调用时所在workspace唯一id
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 数据源连接ID，可选（融合版新增字段）。非空→走分析版路径，空/缺省→走专业版TcLake路径
	ConnectionId *string `json:"ConnectionId,omitnil,omitempty" name:"ConnectionId"`

	// 数据目录来源，可选（Setats融合新增字段）。取值：METALAKE（专业版tclake）/ CONNECTION（分析版Linked Catalog）/ SETATS（Setats元数据）。路由优先级：显式非空CatalogSource > 非空ConnectionId（视为CONNECTION）> 默认METALAKE；CatalogSource非空时ConnectionId将被忽略。
	CatalogSource *string `json:"CatalogSource,omitnil,omitempty" name:"CatalogSource"`
}

type DeleteCatalogRequest struct {
	*tchttp.BaseRequest
	
	// 数据目录名
	CatalogName *string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`

	// 调用时所在workspace唯一id
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 数据源连接ID，可选（融合版新增字段）。非空→走分析版路径，空/缺省→走专业版TcLake路径
	ConnectionId *string `json:"ConnectionId,omitnil,omitempty" name:"ConnectionId"`

	// 数据目录来源，可选（Setats融合新增字段）。取值：METALAKE（专业版tclake）/ CONNECTION（分析版Linked Catalog）/ SETATS（Setats元数据）。路由优先级：显式非空CatalogSource > 非空ConnectionId（视为CONNECTION）> 默认METALAKE；CatalogSource非空时ConnectionId将被忽略。
	CatalogSource *string `json:"CatalogSource,omitnil,omitempty" name:"CatalogSource"`
}

func (r *DeleteCatalogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCatalogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CatalogName")
	delete(f, "WorkspaceId")
	delete(f, "ConnectionId")
	delete(f, "CatalogSource")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCatalogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCatalogResponseParams struct {
	// 操作结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *DeleteCatalogRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCatalogResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCatalogResponseParams `json:"Response"`
}

func (r *DeleteCatalogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCatalogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCatalogRsp struct {
	// 操作结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`
}

// Predefined struct for user
type DeleteConsoleGroupsRequestParams struct {
	// 要删除的用户组 ID 列表，单次最多100个；可通过 ListConsoleGroups 接口获取
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`
}

type DeleteConsoleGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 要删除的用户组 ID 列表，单次最多100个；可通过 ListConsoleGroups 接口获取
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`
}

func (r *DeleteConsoleGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConsoleGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteConsoleGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteConsoleGroupsResponseParams struct {
	// <p>返回结果</p>
	Data *DeleteConsoleGroupsRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteConsoleGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteConsoleGroupsResponseParams `json:"Response"`
}

func (r *DeleteConsoleGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConsoleGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteConsoleGroupsRsp struct {
	// 操作是否成功
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type DeleteFileRequestParams struct {
	// <p>工作空间 ID。来源：ListWorkspaces 接口返回的 WorkspaceId</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>文件 ID。来源：CreateFile / ListFiles / GetFile 接口返回的 FileId</p>
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <p>文件类型。取值：FILE（普通文件/脚本）、NOTEBOOK_FILE（Notebook）、SQL_FILE（SQL文件）。对应 common/domain/entity.proto EntityType</p>
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`
}

type DeleteFileRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间 ID。来源：ListWorkspaces 接口返回的 WorkspaceId</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>文件 ID。来源：CreateFile / ListFiles / GetFile 接口返回的 FileId</p>
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <p>文件类型。取值：FILE（普通文件/脚本）、NOTEBOOK_FILE（Notebook）、SQL_FILE（SQL文件）。对应 common/domain/entity.proto EntityType</p>
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`
}

func (r *DeleteFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "FileId")
	delete(f, "FileType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFileResponseParams struct {
	// <p>返回结果</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *DeleteFileResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteFileResponse struct {
	*tchttp.BaseResponse
	Response *DeleteFileResponseParams `json:"Response"`
}

func (r *DeleteFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteFileResult struct {
	// <p>被删除的文件 ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <p>删除是否成功</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type DeleteFolderRequestParams struct {
	// <p>工作空间id</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>待删除的文件夹</p>
	Folder *FolderLocator `json:"Folder,omitnil,omitempty" name:"Folder"`

	// <p>软删除还是从回收站硬删除</p><p>枚举值：</p><ul><li>false： 软删除到回收站</li><li>true： 从回收站硬删除</li></ul>
	ForceDelete *bool `json:"ForceDelete,omitnil,omitempty" name:"ForceDelete"`
}

type DeleteFolderRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间id</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>待删除的文件夹</p>
	Folder *FolderLocator `json:"Folder,omitnil,omitempty" name:"Folder"`

	// <p>软删除还是从回收站硬删除</p><p>枚举值：</p><ul><li>false： 软删除到回收站</li><li>true： 从回收站硬删除</li></ul>
	ForceDelete *bool `json:"ForceDelete,omitnil,omitempty" name:"ForceDelete"`
}

func (r *DeleteFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "Folder")
	delete(f, "ForceDelete")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFolderResponseParams struct {
	// <p>删除文件夹结果</p>
	Data *DeleteFolderRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteFolderResponse struct {
	*tchttp.BaseResponse
	Response *DeleteFolderResponseParams `json:"Response"`
}

func (r *DeleteFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteFolderRsp struct {
	// <p>删除文件夹状态</p>
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type DeleteSchemaRequestParams struct {
	// 数据目录名
	CatalogName *string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`

	// 数据库名
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// 调用时所在workspace唯一id
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 数据源连接ID，可选（融合版新增字段）。分析版catalog不支持删除schema，传入非空时服务端返回ANA_CATALOG_NOT_SUPPORTED错误
	ConnectionId *string `json:"ConnectionId,omitnil,omitempty" name:"ConnectionId"`

	// 数据目录来源，可选（Setats融合新增字段）。取值：METALAKE（专业版tclake）/ CONNECTION（分析版Linked Catalog）/ SETATS（Setats元数据）。路由优先级：显式非空CatalogSource > 非空ConnectionId（视为CONNECTION）> 默认METALAKE；CatalogSource非空时ConnectionId将被忽略。
	CatalogSource *string `json:"CatalogSource,omitnil,omitempty" name:"CatalogSource"`
}

type DeleteSchemaRequest struct {
	*tchttp.BaseRequest
	
	// 数据目录名
	CatalogName *string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`

	// 数据库名
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// 调用时所在workspace唯一id
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 数据源连接ID，可选（融合版新增字段）。分析版catalog不支持删除schema，传入非空时服务端返回ANA_CATALOG_NOT_SUPPORTED错误
	ConnectionId *string `json:"ConnectionId,omitnil,omitempty" name:"ConnectionId"`

	// 数据目录来源，可选（Setats融合新增字段）。取值：METALAKE（专业版tclake）/ CONNECTION（分析版Linked Catalog）/ SETATS（Setats元数据）。路由优先级：显式非空CatalogSource > 非空ConnectionId（视为CONNECTION）> 默认METALAKE；CatalogSource非空时ConnectionId将被忽略。
	CatalogSource *string `json:"CatalogSource,omitnil,omitempty" name:"CatalogSource"`
}

func (r *DeleteSchemaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSchemaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CatalogName")
	delete(f, "SchemaName")
	delete(f, "WorkspaceId")
	delete(f, "ConnectionId")
	delete(f, "CatalogSource")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSchemaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSchemaResponseParams struct {
	// 操作结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *DeleteSchemaRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSchemaResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSchemaResponseParams `json:"Response"`
}

func (r *DeleteSchemaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSchemaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSchemaRsp struct {
	// 操作结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`
}

// Predefined struct for user
type DeleteTableRequestParams struct {
	// 数据目录名
	CatalogName *string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`

	// 数据库名
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// 表名
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 调用时所在workspace唯一id
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 数据源连接ID，可选（融合版新增字段）。分析版catalog不支持删除表，传入非空时服务端返回ANA_CATALOG_NOT_SUPPORTED错误
	ConnectionId *string `json:"ConnectionId,omitnil,omitempty" name:"ConnectionId"`

	// 数据目录来源，可选（Setats融合新增字段）。取值：METALAKE（专业版tclake）/ CONNECTION（分析版Linked Catalog）/ SETATS（Setats元数据）。路由优先级：显式非空CatalogSource > 非空ConnectionId（视为CONNECTION）> 默认METALAKE；CatalogSource非空时ConnectionId将被忽略。
	CatalogSource *string `json:"CatalogSource,omitnil,omitempty" name:"CatalogSource"`
}

type DeleteTableRequest struct {
	*tchttp.BaseRequest
	
	// 数据目录名
	CatalogName *string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`

	// 数据库名
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// 表名
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 调用时所在workspace唯一id
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 数据源连接ID，可选（融合版新增字段）。分析版catalog不支持删除表，传入非空时服务端返回ANA_CATALOG_NOT_SUPPORTED错误
	ConnectionId *string `json:"ConnectionId,omitnil,omitempty" name:"ConnectionId"`

	// 数据目录来源，可选（Setats融合新增字段）。取值：METALAKE（专业版tclake）/ CONNECTION（分析版Linked Catalog）/ SETATS（Setats元数据）。路由优先级：显式非空CatalogSource > 非空ConnectionId（视为CONNECTION）> 默认METALAKE；CatalogSource非空时ConnectionId将被忽略。
	CatalogSource *string `json:"CatalogSource,omitnil,omitempty" name:"CatalogSource"`
}

func (r *DeleteTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CatalogName")
	delete(f, "SchemaName")
	delete(f, "TableName")
	delete(f, "WorkspaceId")
	delete(f, "ConnectionId")
	delete(f, "CatalogSource")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTableResponseParams struct {
	// 操作结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *DeleteTableRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTableResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTableResponseParams `json:"Response"`
}

func (r *DeleteTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTableRsp struct {
	// 操作结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`
}

// Predefined struct for user
type DeleteWorkflowRequestParams struct {
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>待删除的工作流ID，可通过 ListWorkflows 获取。必填</p>
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

type DeleteWorkflowRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>待删除的工作流ID，可通过 ListWorkflows 获取。必填</p>
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

func (r *DeleteWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWorkflowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "WorkflowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWorkflowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWorkflowResponseParams struct {
	// <p>删除工作流响应内容</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *DeleteWorkflowRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *DeleteWorkflowResponseParams `json:"Response"`
}

func (r *DeleteWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWorkflowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteWorkflowRsp struct {
	// 删除状态，true 表示成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type DeleteWorkspaceRequestParams struct {
	// <p>工作空间ID</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`
}

type DeleteWorkspaceRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间ID</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`
}

func (r *DeleteWorkspaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWorkspaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWorkspaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWorkspaceResponseParams struct {
	// <p>操作结果</p>
	Data *DeleteWorkspaceRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteWorkspaceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteWorkspaceResponseParams `json:"Response"`
}

func (r *DeleteWorkspaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWorkspaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWorkspaceRoleRequestParams struct {
	// 工作空间ID
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 角色ID，可通过 ListWorkspaceRoles 接口获取
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`
}

type DeleteWorkspaceRoleRequest struct {
	*tchttp.BaseRequest
	
	// 工作空间ID
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 角色ID，可通过 ListWorkspaceRoles 接口获取
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`
}

func (r *DeleteWorkspaceRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWorkspaceRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "RoleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWorkspaceRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWorkspaceRoleResponseParams struct {
	// 删除工作空间角色结果
	Data *DeleteWorkspaceRoleResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteWorkspaceRoleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteWorkspaceRoleResponseParams `json:"Response"`
}

func (r *DeleteWorkspaceRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWorkspaceRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteWorkspaceRoleResult struct {
	// 删除工作空间角色是否成功
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

type DeleteWorkspaceRsp struct {
	// 操作是否成功
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

type DependOnBrief struct {
	// 任务ID，可通过 ListWorkflowTasks 获取
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`
}

type DictDimensionInfo struct {
	// 关联的标准字典代码ID
	DictCodeId *string `json:"DictCodeId,omitnil,omitempty" name:"DictCodeId"`

	// 字典项列表（key-value映射）
	// 注意：此字段可能返回 null，表示取不到有效值。
	DictItems []*DictItem `json:"DictItems,omitnil,omitempty" name:"DictItems"`
}

type DictItem struct {
	// 字典键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 字典值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type FeatureTableDetail struct {
	// 特征表主键字段列表
	PrimaryKeys []*string `json:"PrimaryKeys,omitnil,omitempty" name:"PrimaryKeys"`

	// 特征表时间戳字段名称
	TimestampField *string `json:"TimestampField,omitnil,omitempty" name:"TimestampField"`

	// 在线特征表对应的来源表名称（离线表）
	SourceTableName *string `json:"SourceTableName,omitnil,omitempty" name:"SourceTableName"`

	// 在线特征表关联的工作流ID
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 离线特征表对应生成的在线特征表名称
	OnlineTableName *string `json:"OnlineTableName,omitnil,omitempty" name:"OnlineTableName"`
}

type FetchOption struct {
	// <p>是否在响应中返回权限列表，默认false</p>
	FetchPermissions *bool `json:"FetchPermissions,omitnil,omitempty" name:"FetchPermissions"`

	// <p>是否获取特征表详情，当AssetType为TABLE时有效</p>
	FetchFeatureTableDetail *bool `json:"FetchFeatureTableDetail,omitnil,omitempty" name:"FetchFeatureTableDetail"`

	// <p>按权限过滤，传入权限列表，仅返回当前用户拥有指定权限的实体。例如传入[&quot;SELECT_TABLE&quot;]则仅返回当前用户有SELECT_TABLE权限的实体。只对list接口生效，为空时不进行权限过滤</p>
	FilterPermissions []*string `json:"FilterPermissions,omitnil,omitempty" name:"FilterPermissions"`

	// <p>是否在响应中返回负责人信息。不传或为true时返回负责人信息（默认返回），显式传false时不返回</p>
	FetchOwners *bool `json:"FetchOwners,omitnil,omitempty" name:"FetchOwners"`

	// <p>是否将用户Uin转换为用户名(userName)。影响范围：Audit中的CreatorName/LastModifierName、MetaOwner中的OwnerName。不传或为true时执行转换（默认转换），显式传false时不转换</p>
	FetchUserInfo *bool `json:"FetchUserInfo,omitnil,omitempty" name:"FetchUserInfo"`

	// <p>是否返回字段脱敏策略信息，默认不返回，传true则会查询表字段对应的字段脱敏策略信息</p>
	FetchMask *bool `json:"FetchMask,omitnil,omitempty" name:"FetchMask"`

	// <p>是否返回标签信息，默认不返回。传true时，GetTable/ListTables/GetCatalog/ListCatalogs/GetSchema/ListSchemas/GetView/ListViews/GetFunction/ListFunctions/GetVolume/ListVolumes/GetModel/ListModels等接口会在对应实体中返回标签（Tags）字段</p>
	FetchTags *bool `json:"FetchTags,omitnil,omitempty" name:"FetchTags"`

	// <p>是否返回字段关联的字典维度信息，默认不传，不返回</p>
	FetchDimensions *bool `json:"FetchDimensions,omitnil,omitempty" name:"FetchDimensions"`
}

type FileConfig struct {
	// <p>高级运行参数，变量替换用，map-json String,String</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Params *string `json:"Params,omitnil,omitempty" name:"Params"`

	// <p>执行资源 ID。来源：ListComputeResources 接口返回的 ResourceId</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// <p>默认 catalog</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultCatalog *string `json:"DefaultCatalog,omitnil,omitempty" name:"DefaultCatalog"`

	// <p>默认 schema</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultSchema *string `json:"DefaultSchema,omitnil,omitempty" name:"DefaultSchema"`

	// <p>高级配置，JSON 字符串</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdvanceConfig *string `json:"AdvanceConfig,omitnil,omitempty" name:"AdvanceConfig"`

	// <p>扩展参数，JSON 字符串</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraParams *string `json:"ExtraParams,omitnil,omitempty" name:"ExtraParams"`

	// <p>Notebook 交互控件定义，JSON 字符串</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Widgets *string `json:"Widgets,omitnil,omitempty" name:"Widgets"`

	// <p>各单元格输出配置。仅 Get 出参返回，入参忽略</p>
	OutputConf []*FileOutputConf `json:"OutputConf,omitnil,omitempty" name:"OutputConf"`

	// <p>SQL脚本语法标记</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SqlSyntax *string `json:"SqlSyntax,omitnil,omitempty" name:"SqlSyntax"`

	// <p>平台集群id</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>默认 Catalog 的来源类型，用于区分默认 Catalog 属于哪一套 Catalog 体系，可选。长度上限 32。取值：METALAKE - TC Catalog（元数据湖 Catalog）；SETATS - Setats Catalog。空值/非法值时后端读取兜底为 METALAKE。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultCatalogSource *string `json:"DefaultCatalogSource,omitnil,omitempty" name:"DefaultCatalogSource"`
}

type FileInfo struct {
	// <p>主账号 AppId</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>工作空间 ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>文件 ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <p>文件名，含后缀</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// <p>文件类型。取值：FILE（普通文件/脚本）、NOTEBOOK_FILE（Notebook）、SQL_FILE（SQL文件）。对应 common/domain/entity.proto EntityType</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// <p>文件在工作空间中的完整路径，以 / 开头，如 /etl/daily/demo.ipynb</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// <p>文件运行配置</p>
	FileConfig *FileConfig `json:"FileConfig,omitnil,omitempty" name:"FileConfig"`

	// <p>绑定的 BundleId</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// <p>绑定的 BundleInfo，JSON 字符串</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`

	// <p>文件状态。active=正常，deleted=已删除</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>文件负责人用户名</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUserName *string `json:"OwnerUserName,omitnil,omitempty" name:"OwnerUserName"`

	// <p>创建人子账号 Uin</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// <p>最近更新人子账号 Uin</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateUserUin *string `json:"UpdateUserUin,omitnil,omitempty" name:"UpdateUserUin"`

	// <p>创建时间，毫秒级时间戳</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>最近更新时间，毫秒级时间戳</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>文件存储信息。仅当请求 IncludeContent=true 时返回内容</p>
	Storage *FileStorage `json:"Storage,omitnil,omitempty" name:"Storage"`

	// <p>当前调用方对该文件的权限点列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Permissions *string `json:"Permissions,omitnil,omitempty" name:"Permissions"`

	// <p>是否已发布</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReleaseStatus *bool `json:"ReleaseStatus,omitnil,omitempty" name:"ReleaseStatus"`

	// <p>资源模式。1=分布式，2=单节点</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceMode *int64 `json:"ResourceMode,omitnil,omitempty" name:"ResourceMode"`

	// <p>ZIP 异步创建时透传 Workspace 作业信息；普通同步创建或其他复用该返回结构的接口不设置该字段</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsyncOperation *AsyncOperation `json:"AsyncOperation,omitnil,omitempty" name:"AsyncOperation"`

	// <p>脚本版本ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`
}

type FileMeta struct {
	// <p>文件id</p>
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <p>文件/文件夹名称</p>
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// <p>文件类型</p>
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// <p>创建时间，毫秒秒级时间戳</p><p>参数格式：时间戳</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>更新时间</p><p>参数格式：时间戳字符串</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>acl权限类型</p>
	AllowActions []*string `json:"AllowActions,omitnil,omitempty" name:"AllowActions"`

	// <p>是否收藏</p>
	IsFavorite *bool `json:"IsFavorite,omitnil,omitempty" name:"IsFavorite"`

	// <p>文件path</p>
	PathName *string `json:"PathName,omitnil,omitempty" name:"PathName"`

	// <p>是否系统创建</p>
	IsSystemGenerated *bool `json:"IsSystemGenerated,omitnil,omitempty" name:"IsSystemGenerated"`
}

type FileNode struct {
	// <p>当前节点</p>
	Node *FileMeta `json:"Node,omitnil,omitempty" name:"Node"`

	// <p>父节点</p>
	Parent *FileMeta `json:"Parent,omitnil,omitempty" name:"Parent"`

	// <p>创建人</p>
	Creator *UserInfo `json:"Creator,omitnil,omitempty" name:"Creator"`

	// <p>拥有者</p>
	Owner *UserInfo `json:"Owner,omitnil,omitempty" name:"Owner"`

	// <p>节点类型</p>
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// <p>原始路径</p>
	OriginPath *string `json:"OriginPath,omitnil,omitempty" name:"OriginPath"`

	// <p>回收时间</p>
	DeleteTime *string `json:"DeleteTime,omitnil,omitempty" name:"DeleteTime"`

	// <p>文件git配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	GitConfig *GitRepoConfig `json:"GitConfig,omitnil,omitempty" name:"GitConfig"`
}

type FileOutputConf struct {
	// 单元格 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CellId *string `json:"CellId,omitnil,omitempty" name:"CellId"`

	// Dashboard 图表配置，JSON 字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	DashboardConf *string `json:"DashboardConf,omitnil,omitempty" name:"DashboardConf"`

	// 执行结果文件的预签名下载链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputPath *string `json:"OutputPath,omitnil,omitempty" name:"OutputPath"`
}

type FileStorage struct {
	// 存储类型
	StorageType *int64 `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 存储路径
	StoragePath *string `json:"StoragePath,omitnil,omitempty" name:"StoragePath"`

	// 文件内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type FolderLocator struct {
	// <p>节点id</p>
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// <p>节点path</p>
	PathName *string `json:"PathName,omitnil,omitempty" name:"PathName"`
}

// Predefined struct for user
type GetCatalogRequestParams struct {
	// <p>数据目录名</p>
	CatalogName *string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`

	// <p>调用时所在workspace唯一id</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>数据目录id</p>
	CatalogId *string `json:"CatalogId,omitnil,omitempty" name:"CatalogId"`

	// <p>数据获取选项，可选，控制是否返回权限信息及按权限过滤</p>
	FetchOption *FetchOption `json:"FetchOption,omitnil,omitempty" name:"FetchOption"`

	// <p>数据源连接ID，可选（融合版新增字段）。非空→走分析版路径，空/缺省→走专业版TcLake路径</p>
	ConnectionId *string `json:"ConnectionId,omitnil,omitempty" name:"ConnectionId"`

	// <p>数据目录来源，可选（Setats融合新增字段）。取值：METALAKE（专业版tclake）/ CONNECTION（分析版Linked Catalog）/ SETATS（Setats元数据）。路由优先级：显式非空CatalogSource &gt; 非空ConnectionId（视为CONNECTION）&gt; 默认METALAKE；CatalogSource非空时ConnectionId将被忽略。</p>
	CatalogSource *string `json:"CatalogSource,omitnil,omitempty" name:"CatalogSource"`
}

type GetCatalogRequest struct {
	*tchttp.BaseRequest
	
	// <p>数据目录名</p>
	CatalogName *string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`

	// <p>调用时所在workspace唯一id</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>数据目录id</p>
	CatalogId *string `json:"CatalogId,omitnil,omitempty" name:"CatalogId"`

	// <p>数据获取选项，可选，控制是否返回权限信息及按权限过滤</p>
	FetchOption *FetchOption `json:"FetchOption,omitnil,omitempty" name:"FetchOption"`

	// <p>数据源连接ID，可选（融合版新增字段）。非空→走分析版路径，空/缺省→走专业版TcLake路径</p>
	ConnectionId *string `json:"ConnectionId,omitnil,omitempty" name:"ConnectionId"`

	// <p>数据目录来源，可选（Setats融合新增字段）。取值：METALAKE（专业版tclake）/ CONNECTION（分析版Linked Catalog）/ SETATS（Setats元数据）。路由优先级：显式非空CatalogSource &gt; 非空ConnectionId（视为CONNECTION）&gt; 默认METALAKE；CatalogSource非空时ConnectionId将被忽略。</p>
	CatalogSource *string `json:"CatalogSource,omitnil,omitempty" name:"CatalogSource"`
}

func (r *GetCatalogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCatalogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CatalogName")
	delete(f, "WorkspaceId")
	delete(f, "CatalogId")
	delete(f, "FetchOption")
	delete(f, "ConnectionId")
	delete(f, "CatalogSource")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetCatalogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCatalogResponseParams struct {
	// <p>catalog详情</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *GetCatalogRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetCatalogResponse struct {
	*tchttp.BaseResponse
	Response *GetCatalogResponseParams `json:"Response"`
}

func (r *GetCatalogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCatalogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCatalogRsp struct {
	// catalog详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Catalog *Catalog `json:"Catalog,omitnil,omitempty" name:"Catalog"`
}

// Predefined struct for user
type GetFileRequestParams struct {
	// <p>工作空间 ID。来源：ListWorkspaces 接口返回的 WorkspaceId</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>文件 ID。来源：CreateFile / ListFiles 接口返回的 FileId。与 FilePath 二选一</p>
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <p>文件类型。取值：FILE（普通文件/脚本）、NOTEBOOK_FILE（Notebook）、SQL_FILE（SQL文件）。对应 common/domain/entity.proto EntityType</p>
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// <p>是否返回文件内容。true 时 Storage.Content 返回 base64 内容，默认 false</p>
	IncludeContent *bool `json:"IncludeContent,omitnil,omitempty" name:"IncludeContent"`

	// <p>文件版本 ID。来源：ListFileVersions 接口返回的 VersionId。不传则读取最新版本</p>
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// <p>文件完整路径，以 / 开头，如 /etl/daily/demo.ipynb。与 FileId 二选一</p>
	FilePath *string `json:"FilePath,omitnil,omitempty" name:"FilePath"`
}

type GetFileRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间 ID。来源：ListWorkspaces 接口返回的 WorkspaceId</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>文件 ID。来源：CreateFile / ListFiles 接口返回的 FileId。与 FilePath 二选一</p>
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <p>文件类型。取值：FILE（普通文件/脚本）、NOTEBOOK_FILE（Notebook）、SQL_FILE（SQL文件）。对应 common/domain/entity.proto EntityType</p>
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// <p>是否返回文件内容。true 时 Storage.Content 返回 base64 内容，默认 false</p>
	IncludeContent *bool `json:"IncludeContent,omitnil,omitempty" name:"IncludeContent"`

	// <p>文件版本 ID。来源：ListFileVersions 接口返回的 VersionId。不传则读取最新版本</p>
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// <p>文件完整路径，以 / 开头，如 /etl/daily/demo.ipynb。与 FileId 二选一</p>
	FilePath *string `json:"FilePath,omitnil,omitempty" name:"FilePath"`
}

func (r *GetFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "FileId")
	delete(f, "FileType")
	delete(f, "IncludeContent")
	delete(f, "VersionId")
	delete(f, "FilePath")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetFileResponseParams struct {
	// <p>返回结果</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *FileInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetFileResponse struct {
	*tchttp.BaseResponse
	Response *GetFileResponseParams `json:"Response"`
}

func (r *GetFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetFolderRequestParams struct {
	// <p>工作空间id</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>文件夹定位</p>
	Folder *FolderLocator `json:"Folder,omitnil,omitempty" name:"Folder"`
}

type GetFolderRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间id</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>文件夹定位</p>
	Folder *FolderLocator `json:"Folder,omitnil,omitempty" name:"Folder"`
}

func (r *GetFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "Folder")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetFolderResponseParams struct {
	// <p>文件夹详情结果</p>
	Data *GetFolderRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetFolderResponse struct {
	*tchttp.BaseResponse
	Response *GetFolderResponseParams `json:"Response"`
}

func (r *GetFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetFolderRsp struct {
	// <p>文件夹信息</p>
	Folder *FileNode `json:"Folder,omitnil,omitempty" name:"Folder"`
}

// Predefined struct for user
type GetSchemaRequestParams struct {
	// <p>数据目录名</p>
	CatalogName *string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`

	// <p>数据库名</p>
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// <p>调用时所在workspace唯一id</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>数据获取选项，可选，控制是否返回权限信息及按权限过滤</p>
	FetchOption *FetchOption `json:"FetchOption,omitnil,omitempty" name:"FetchOption"`

	// 数据源连接ID，可选（融合版新增字段）。非空→走分析版路径，空/缺省→走专业版TcLake路径
	ConnectionId *string `json:"ConnectionId,omitnil,omitempty" name:"ConnectionId"`

	// 数据目录来源，可选（Setats融合新增字段）。取值：METALAKE（专业版tclake）/ CONNECTION（分析版Linked Catalog）/ SETATS（Setats元数据）。路由优先级：显式非空CatalogSource > 非空ConnectionId（视为CONNECTION）> 默认METALAKE；CatalogSource非空时ConnectionId将被忽略。
	CatalogSource *string `json:"CatalogSource,omitnil,omitempty" name:"CatalogSource"`
}

type GetSchemaRequest struct {
	*tchttp.BaseRequest
	
	// <p>数据目录名</p>
	CatalogName *string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`

	// <p>数据库名</p>
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// <p>调用时所在workspace唯一id</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>数据获取选项，可选，控制是否返回权限信息及按权限过滤</p>
	FetchOption *FetchOption `json:"FetchOption,omitnil,omitempty" name:"FetchOption"`

	// 数据源连接ID，可选（融合版新增字段）。非空→走分析版路径，空/缺省→走专业版TcLake路径
	ConnectionId *string `json:"ConnectionId,omitnil,omitempty" name:"ConnectionId"`

	// 数据目录来源，可选（Setats融合新增字段）。取值：METALAKE（专业版tclake）/ CONNECTION（分析版Linked Catalog）/ SETATS（Setats元数据）。路由优先级：显式非空CatalogSource > 非空ConnectionId（视为CONNECTION）> 默认METALAKE；CatalogSource非空时ConnectionId将被忽略。
	CatalogSource *string `json:"CatalogSource,omitnil,omitempty" name:"CatalogSource"`
}

func (r *GetSchemaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSchemaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CatalogName")
	delete(f, "SchemaName")
	delete(f, "WorkspaceId")
	delete(f, "FetchOption")
	delete(f, "ConnectionId")
	delete(f, "CatalogSource")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetSchemaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetSchemaResponseParams struct {
	// <p>schema详情</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *GetSchemaRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetSchemaResponse struct {
	*tchttp.BaseResponse
	Response *GetSchemaResponseParams `json:"Response"`
}

func (r *GetSchemaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSchemaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSchemaRsp struct {
	// schema详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Schema *Schema `json:"Schema,omitnil,omitempty" name:"Schema"`
}

// Predefined struct for user
type GetTableRequestParams struct {
	// <p>数据目录名</p>
	CatalogName *string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`

	// <p>数据库名</p>
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// <p>表名</p>
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// <p>调用时所在workspace唯一id</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>数据获取选项，可选，控制是否返回权限信息及按权限过滤</p>
	FetchOption *FetchOption `json:"FetchOption,omitnil,omitempty" name:"FetchOption"`

	// 数据源连接ID，可选（融合版新增字段）。非空→走分析版路径，空/缺省→走专业版TcLake路径
	ConnectionId *string `json:"ConnectionId,omitnil,omitempty" name:"ConnectionId"`

	// 数据目录来源，可选（Setats融合新增字段）。取值：METALAKE（专业版tclake）/ CONNECTION（分析版Linked Catalog）/ SETATS（Setats元数据）。路由优先级：显式非空CatalogSource > 非空ConnectionId（视为CONNECTION）> 默认METALAKE；CatalogSource非空时ConnectionId将被忽略。
	CatalogSource *string `json:"CatalogSource,omitnil,omitempty" name:"CatalogSource"`
}

type GetTableRequest struct {
	*tchttp.BaseRequest
	
	// <p>数据目录名</p>
	CatalogName *string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`

	// <p>数据库名</p>
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// <p>表名</p>
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// <p>调用时所在workspace唯一id</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>数据获取选项，可选，控制是否返回权限信息及按权限过滤</p>
	FetchOption *FetchOption `json:"FetchOption,omitnil,omitempty" name:"FetchOption"`

	// 数据源连接ID，可选（融合版新增字段）。非空→走分析版路径，空/缺省→走专业版TcLake路径
	ConnectionId *string `json:"ConnectionId,omitnil,omitempty" name:"ConnectionId"`

	// 数据目录来源，可选（Setats融合新增字段）。取值：METALAKE（专业版tclake）/ CONNECTION（分析版Linked Catalog）/ SETATS（Setats元数据）。路由优先级：显式非空CatalogSource > 非空ConnectionId（视为CONNECTION）> 默认METALAKE；CatalogSource非空时ConnectionId将被忽略。
	CatalogSource *string `json:"CatalogSource,omitnil,omitempty" name:"CatalogSource"`
}

func (r *GetTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CatalogName")
	delete(f, "SchemaName")
	delete(f, "TableName")
	delete(f, "WorkspaceId")
	delete(f, "FetchOption")
	delete(f, "ConnectionId")
	delete(f, "CatalogSource")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTableResponseParams struct {
	// <p>表信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *GetTableRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetTableResponse struct {
	*tchttp.BaseResponse
	Response *GetTableResponseParams `json:"Response"`
}

func (r *GetTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTableRsp struct {
	// table详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Table *TableInfo `json:"Table,omitnil,omitempty" name:"Table"`
}

// Predefined struct for user
type GetTablesRequestParams struct {
	// <p>调用时所在workspace唯一id，必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>表的全名列表，格式为 catalogName.schemaName.tableName，必填，至少1条，最大20条</p>
	FullNames []*string `json:"FullNames,omitnil,omitempty" name:"FullNames"`

	// <p>数据获取选项，可选，控制是否返回权限信息及按权限过滤</p>
	FetchOption *FetchOption `json:"FetchOption,omitnil,omitempty" name:"FetchOption"`

	// 数据源连接ID，可选（融合版新增字段）。非空→走分析版路径，要求所有FullNames属于同一catalog（分析版额外约束）；空/缺省→走专业版TcLake路径
	ConnectionId *string `json:"ConnectionId,omitnil,omitempty" name:"ConnectionId"`

	// 数据目录来源，可选（Setats融合新增字段）。取值：METALAKE（专业版tclake）/ CONNECTION（分析版Linked Catalog）/ SETATS（Setats元数据）。路由优先级：显式非空CatalogSource > 非空ConnectionId（视为CONNECTION）> 默认METALAKE；CatalogSource非空时ConnectionId将被忽略。
	CatalogSource *string `json:"CatalogSource,omitnil,omitempty" name:"CatalogSource"`
}

type GetTablesRequest struct {
	*tchttp.BaseRequest
	
	// <p>调用时所在workspace唯一id，必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>表的全名列表，格式为 catalogName.schemaName.tableName，必填，至少1条，最大20条</p>
	FullNames []*string `json:"FullNames,omitnil,omitempty" name:"FullNames"`

	// <p>数据获取选项，可选，控制是否返回权限信息及按权限过滤</p>
	FetchOption *FetchOption `json:"FetchOption,omitnil,omitempty" name:"FetchOption"`

	// 数据源连接ID，可选（融合版新增字段）。非空→走分析版路径，要求所有FullNames属于同一catalog（分析版额外约束）；空/缺省→走专业版TcLake路径
	ConnectionId *string `json:"ConnectionId,omitnil,omitempty" name:"ConnectionId"`

	// 数据目录来源，可选（Setats融合新增字段）。取值：METALAKE（专业版tclake）/ CONNECTION（分析版Linked Catalog）/ SETATS（Setats元数据）。路由优先级：显式非空CatalogSource > 非空ConnectionId（视为CONNECTION）> 默认METALAKE；CatalogSource非空时ConnectionId将被忽略。
	CatalogSource *string `json:"CatalogSource,omitnil,omitempty" name:"CatalogSource"`
}

func (r *GetTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "FullNames")
	delete(f, "FetchOption")
	delete(f, "ConnectionId")
	delete(f, "CatalogSource")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTablesResponseParams struct {
	// <p>返回结果</p>
	Data *GetTablesRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetTablesResponse struct {
	*tchttp.BaseResponse
	Response *GetTablesResponseParams `json:"Response"`
}

func (r *GetTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTablesRsp struct {
	// 表信息列表。注意：仅返回当前workspace有权限的catalog下的表，无权限的表会被静默跳过
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*TableInfo `json:"Items,omitnil,omitempty" name:"Items"`
}

// Predefined struct for user
type GetWorkflowRequestParams struct {
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>工作流ID，可通过 ListWorkflows 获取。必填</p>
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

type GetWorkflowRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>工作流ID，可通过 ListWorkflows 获取。必填</p>
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

func (r *GetWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetWorkflowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "WorkflowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetWorkflowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetWorkflowResponseParams struct {
	// <p>获取工作流详细信息响应内容</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *GetWorkflowRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *GetWorkflowResponseParams `json:"Response"`
}

func (r *GetWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetWorkflowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetWorkflowRsp struct {
	// <p>工作空间ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>工作流基本信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BaseInfo *WorkflowBaseInfoDetail `json:"BaseInfo,omitnil,omitempty" name:"BaseInfo"`

	// <p>工作流调度配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Trigger []*WorkflowTriggerConfiguration `json:"Trigger,omitnil,omitempty" name:"Trigger"`

	// <p>工作流参数列表 参数名必填且只能包含数字、大小写字母、空格、.$@#!%^&amp;*()-_+=&gt;</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamList []*ParamInfo `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// <p>标签 标签名必填且只能包含数字、大小写字母、空格、.$@#!%^&amp;*()-_+=&gt;</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelList []*LabelBrief `json:"LabelList,omitnil,omitempty" name:"LabelList"`

	// <p>工作流告警配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alarm *AlarmBrief `json:"Alarm,omitnil,omitempty" name:"Alarm"`

	// <p>监控指标配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	MonitorMetric *MonitorMetricBrief `json:"MonitorMetric,omitnil,omitempty" name:"MonitorMetric"`

	// <p>工作流高级设置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdvanceConfig *WorkflowAdvanceConfig `json:"AdvanceConfig,omitnil,omitempty" name:"AdvanceConfig"`

	// <p>工作流任务列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskList []*WorkflowTask `json:"TaskList,omitnil,omitempty" name:"TaskList"`

	// <p>工作流绑定的 Bundle唯一标识，未绑定时为空</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// <p>Bundle信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`

	// <p>Git配置ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	GitConfigId *string `json:"GitConfigId,omitnil,omitempty" name:"GitConfigId"`

	// <p>Git分支信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	GitBranch *string `json:"GitBranch,omitnil,omitempty" name:"GitBranch"`
}

// Predefined struct for user
type GetWorkflowRunRequestParams struct {
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>工作流运行ID，可通过 ListWorkflowRuns 获取。必填</p>
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`
}

type GetWorkflowRunRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>工作流运行ID，可通过 ListWorkflowRuns 获取。必填</p>
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`
}

func (r *GetWorkflowRunRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetWorkflowRunRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "WorkflowRunId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetWorkflowRunRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetWorkflowRunResponseParams struct {
	// <p>查询工作流运行详情响应内容</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *GetWorkflowRunRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetWorkflowRunResponse struct {
	*tchttp.BaseResponse
	Response *GetWorkflowRunResponseParams `json:"Response"`
}

func (r *GetWorkflowRunResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetWorkflowRunResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetWorkflowRunRsp struct {
	// 工作流运行信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowRun *WorkflowRun `json:"WorkflowRun,omitnil,omitempty" name:"WorkflowRun"`
}

// Predefined struct for user
type GetWorkflowTaskRunRequestParams struct {
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>任务运行ID，可通过 ListWorkflowTaskRuns 获取。必填</p>
	WorkflowTaskRunId *string `json:"WorkflowTaskRunId,omitnil,omitempty" name:"WorkflowTaskRunId"`

	// <p>内嵌工作流任务运行列表选项（仅限 FOR_EACH 任务使用）。非必填</p>
	InnerWorkflowTaskRunListOption *InnerWorkflowTaskRunListOption `json:"InnerWorkflowTaskRunListOption,omitnil,omitempty" name:"InnerWorkflowTaskRunListOption"`
}

type GetWorkflowTaskRunRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>任务运行ID，可通过 ListWorkflowTaskRuns 获取。必填</p>
	WorkflowTaskRunId *string `json:"WorkflowTaskRunId,omitnil,omitempty" name:"WorkflowTaskRunId"`

	// <p>内嵌工作流任务运行列表选项（仅限 FOR_EACH 任务使用）。非必填</p>
	InnerWorkflowTaskRunListOption *InnerWorkflowTaskRunListOption `json:"InnerWorkflowTaskRunListOption,omitnil,omitempty" name:"InnerWorkflowTaskRunListOption"`
}

func (r *GetWorkflowTaskRunRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetWorkflowTaskRunRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "WorkflowTaskRunId")
	delete(f, "InnerWorkflowTaskRunListOption")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetWorkflowTaskRunRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetWorkflowTaskRunResponseParams struct {
	// <p>查询任务运行详情响应内容</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *GetWorkflowTaskRunRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetWorkflowTaskRunResponse struct {
	*tchttp.BaseResponse
	Response *GetWorkflowTaskRunResponseParams `json:"Response"`
}

func (r *GetWorkflowTaskRunResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetWorkflowTaskRunResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetWorkflowTaskRunRsp struct {
	// <p>任务名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// <p>任务运行ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowTaskRunId *string `json:"WorkflowTaskRunId,omitnil,omitempty" name:"WorkflowTaskRunId"`

	// <p>运行状态。取值参考工作流任务运行状态枚举，如 Pending / Running / Succeeded / Failed / Killed</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunState *string `json:"RunState,omitnil,omitempty" name:"RunState"`

	// <p>工作空间ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>工作流ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// <p>工作流运行ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`

	// <p>任务ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <p>任务类型名称，请参考数据结构TaskType中TaskTypeName字段描述</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTypeName *string `json:"TaskTypeName,omitnil,omitempty" name:"TaskTypeName"`

	// <p>任务版本ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskVersionId *string `json:"TaskVersionId,omitnil,omitempty" name:"TaskVersionId"`

	// <p>触发类型</p><p>枚举值：</p><ul><li>Scheduler： 调度触发</li><li>Manual： 手动触发</li></ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TriggerType *string `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// <p>所属资源组ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// <p>错误码</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorCodeString *string `json:"ErrorCodeString,omitnil,omitempty" name:"ErrorCodeString"`

	// <p>运行用户UIN</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunUserUin *string `json:"RunUserUin,omitnil,omitempty" name:"RunUserUin"`

	// <p>运行用户名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunUserName *string `json:"RunUserName,omitnil,omitempty" name:"RunUserName"`

	// <p>创建人UIN</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// <p>执行平台执行ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// <p>创建时间，单位：毫秒时间戳</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>更新时间，单位：毫秒时间戳</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>依赖任务完成时间，单位：毫秒时间戳</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DependenceFinishedTime *string `json:"DependenceFinishedTime,omitnil,omitempty" name:"DependenceFinishedTime"`

	// <p>运行开始时间，单位：毫秒时间戳</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunStartTime *string `json:"RunStartTime,omitnil,omitempty" name:"RunStartTime"`

	// <p>运行结束时间，单位：毫秒时间戳</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunEndTime *string `json:"RunEndTime,omitnil,omitempty" name:"RunEndTime"`

	// <p>运行时长，单位：秒</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunCostTime *string `json:"RunCostTime,omitnil,omitempty" name:"RunCostTime"`

	// <p>等待时长（依赖就绪到开始运行的等待耗时），单位：秒</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WaitTime *string `json:"WaitTime,omitnil,omitempty" name:"WaitTime"`

	// <p>下发执行平台时间，单位：毫秒时间戳</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IssueTime *string `json:"IssueTime,omitnil,omitempty" name:"IssueTime"`

	// <p>时区</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`

	// <p>依赖上游任务ID列表。保留字段，暂时返回为[]</p><p>保留字段，暂时返回为[]</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DependOnList []*string `json:"DependOnList,omitnil,omitempty" name:"DependOnList"`

	// <p>运行参数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunParams *string `json:"RunParams,omitnil,omitempty" name:"RunParams"`

	// <p>任务扩展信息，包含脚本路径</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTypeExtensions *string `json:"TaskTypeExtensions,omitnil,omitempty" name:"TaskTypeExtensions"`

	// <p>任务X坐标</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	LeftCoordinate *float64 `json:"LeftCoordinate,omitnil,omitempty" name:"LeftCoordinate"`

	// <p>任务Y坐标</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopCoordinate *float64 `json:"TopCoordinate,omitnil,omitempty" name:"TopCoordinate"`

	// <p>重试次数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RetryTimes *int64 `json:"RetryTimes,omitnil,omitempty" name:"RetryTimes"`

	// <p>工作流名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// <p>重跑次数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RerunTimes *int64 `json:"RerunTimes,omitnil,omitempty" name:"RerunTimes"`

	// <p>是否最新一次运行</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsLatestRun *bool `json:"IsLatestRun,omitnil,omitempty" name:"IsLatestRun"`

	// <p>资源组信息列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupInfoList []*ResourceGroupInfo `json:"ResourceGroupInfoList,omitnil,omitempty" name:"ResourceGroupInfoList"`

	// <p>错误消息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// <p>运行结果</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunResult *string `json:"RunResult,omitnil,omitempty" name:"RunResult"`

	// <p>内嵌工作流任务运行详情（仅限 FOR_EACH 任务，其他任务类型不返回该字段）</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerWorkflowTaskRun *InnerWorkflowTaskRun `json:"InnerWorkflowTaskRun,omitnil,omitempty" name:"InnerWorkflowTaskRun"`

	// <p>计划调度时间</p><p>参数格式：毫秒时间戳（UTC）</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScheduledTime *string `json:"ScheduledTime,omitnil,omitempty" name:"ScheduledTime"`
}

// Predefined struct for user
type GetWorkspaceRequestParams struct {
	// <p>工作空间ID</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`
}

type GetWorkspaceRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间ID</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`
}

func (r *GetWorkspaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetWorkspaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetWorkspaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetWorkspaceResponseParams struct {
	// <p>工作空间详情</p>
	Data *GetWorkspaceRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetWorkspaceResponse struct {
	*tchttp.BaseResponse
	Response *GetWorkspaceResponseParams `json:"Response"`
}

func (r *GetWorkspaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetWorkspaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetWorkspaceRsp struct {
	// 工作空间详情
	WorkspaceInfo *WorkspaceInfo `json:"WorkspaceInfo,omitnil,omitempty" name:"WorkspaceInfo"`
}

type GitRepoConfig struct {
	// <p>检出规则</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SparseCheckout *SparseCheckoutConfig `json:"SparseCheckout,omitnil,omitempty" name:"SparseCheckout"`

	// <p>Git 仓库地址</p>
	RepoUrl *string `json:"RepoUrl,omitnil,omitempty" name:"RepoUrl"`

	// <p>分支名</p>
	Branch *string `json:"Branch,omitnil,omitempty" name:"Branch"`

	// <p>关联的 gitAuth 配置名称</p>
	AuthConfigName *string `json:"AuthConfigName,omitnil,omitempty" name:"AuthConfigName"`
}

type GroupUserInfo struct {
	// <p>用户UIN</p>
	UserUin *string `json:"UserUin,omitnil,omitempty" name:"UserUin"`

	// <p>用户名</p>
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// <p>昵称</p>
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// <p>创建时间</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>更新时间</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type Index struct {
	// 字段键值类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 字段名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FieldNames []*string `json:"FieldNames,omitnil,omitempty" name:"FieldNames"`

	// 字段键值类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexType *string `json:"IndexType,omitnil,omitempty" name:"IndexType"`
}

type InnerWorkflowTaskBrief struct {
	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 任务类型名称，请参考数据结构TaskType中TaskTypeName字段描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTypeName *string `json:"TaskTypeName,omitnil,omitempty" name:"TaskTypeName"`
}

type InnerWorkflowTaskRun struct {
	// 当前页码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 总记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 总页数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPageNumber *int64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`

	// 迭代运行列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*InnerWorkflowTaskRunIteration `json:"Items,omitnil,omitempty" name:"Items"`

	// 迭代次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	IterationCount *int64 `json:"IterationCount,omitnil,omitempty" name:"IterationCount"`

	// 失败次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailureCount *int64 `json:"FailureCount,omitnil,omitempty" name:"FailureCount"`

	// 成功次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessCount *int64 `json:"SuccessCount,omitnil,omitempty" name:"SuccessCount"`

	// 内嵌工作流ID，可通过 ListWorkflows 获取
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerWorkflowId *string `json:"InnerWorkflowId,omitnil,omitempty" name:"InnerWorkflowId"`

	// 内嵌任务ID，可通过 ListWorkflowTasks 获取
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerTaskId *string `json:"InnerTaskId,omitnil,omitempty" name:"InnerTaskId"`

	// 内嵌任务运行状态数量统计（实例业务枚举键值对列表）
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerTaskRunBizEnumInfos []*ScheduleBizEnumBrief `json:"InnerTaskRunBizEnumInfos,omitnil,omitempty" name:"InnerTaskRunBizEnumInfos"`
}

type InnerWorkflowTaskRunIteration struct {
	// <p>内嵌工作流运行ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`

	// <p>迭代序号</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IterationIndex *string `json:"IterationIndex,omitnil,omitempty" name:"IterationIndex"`

	// <p>运行开始时间，单位：毫秒时间戳</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunStartTime *string `json:"RunStartTime,omitnil,omitempty" name:"RunStartTime"`

	// <p>运行结束时间，单位：毫秒时间戳</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunEndTime *string `json:"RunEndTime,omitnil,omitempty" name:"RunEndTime"`

	// <p>运行状态（参考工作流运行状态枚举）</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunState *string `json:"RunState,omitnil,omitempty" name:"RunState"`

	// <p>运行时长，单位：秒</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunCostTime *string `json:"RunCostTime,omitnil,omitempty" name:"RunCostTime"`

	// <p>运行参数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowParams *string `json:"WorkflowParams,omitnil,omitempty" name:"WorkflowParams"`

	// <p>错误码</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorCodeString *string `json:"ErrorCodeString,omitnil,omitempty" name:"ErrorCodeString"`

	// <p>内嵌工作流内部的任务运行</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerTaskRun *InnerWorkflowTaskRunIterationBrief `json:"InnerTaskRun,omitnil,omitempty" name:"InnerTaskRun"`
}

type InnerWorkflowTaskRunIterationBrief struct {
	// <p>任务运行ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowTaskRunId *string `json:"WorkflowTaskRunId,omitnil,omitempty" name:"WorkflowTaskRunId"`

	// <p>迭代序号</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IterationIndex *string `json:"IterationIndex,omitnil,omitempty" name:"IterationIndex"`

	// <p>运行开始时间，单位：毫秒时间戳</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunStartTime *string `json:"RunStartTime,omitnil,omitempty" name:"RunStartTime"`

	// <p>运行结束时间，单位：毫秒时间戳</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunEndTime *string `json:"RunEndTime,omitnil,omitempty" name:"RunEndTime"`

	// <p>运行状态</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunState *string `json:"RunState,omitnil,omitempty" name:"RunState"`

	// <p>运行时长，单位：秒</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunCostTime *string `json:"RunCostTime,omitnil,omitempty" name:"RunCostTime"`

	// <p>运行参数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskParams *string `json:"TaskParams,omitnil,omitempty" name:"TaskParams"`

	// <p>错误码</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorCodeString *string `json:"ErrorCodeString,omitnil,omitempty" name:"ErrorCodeString"`
}

type InnerWorkflowTaskRunListOption struct {
	// <p>分页页码，从 1 开始。非必填，默认 1</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页大小。非必填，默认 10，取值范围 [10, 200]</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>迭代运行状态，精确匹配。非必填，多选（多个值之间为 OR 关系）。</p><p>可填 SUCCESS / FAILED 等，具体参考本接口出参 InnerWorkflowTaskRunIteration.RunState 字段返回值。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunStates []*string `json:"RunStates,omitnil,omitempty" name:"RunStates"`
}

type KVPair struct {
	// 键
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type KillWorkflowRunRequestParams struct {
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>工作流ID，可通过 ListWorkflows 获取。必填</p>
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// <p>待终止的工作流运行ID列表，可通过 ListWorkflowRuns 获取</p>
	WorkflowRunIds []*string `json:"WorkflowRunIds,omitnil,omitempty" name:"WorkflowRunIds"`

	// <p>是否终止该工作流下所有未进入终态的运行。非必填，默认 false</p>
	KillAllRuns *bool `json:"KillAllRuns,omitnil,omitempty" name:"KillAllRuns"`

	// <p>是否只终止处于等待中（Pending）状态的运行。非必填，默认 false</p>
	OnlyKillPendingRuns *bool `json:"OnlyKillPendingRuns,omitnil,omitempty" name:"OnlyKillPendingRuns"`
}

type KillWorkflowRunRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>工作流ID，可通过 ListWorkflows 获取。必填</p>
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// <p>待终止的工作流运行ID列表，可通过 ListWorkflowRuns 获取</p>
	WorkflowRunIds []*string `json:"WorkflowRunIds,omitnil,omitempty" name:"WorkflowRunIds"`

	// <p>是否终止该工作流下所有未进入终态的运行。非必填，默认 false</p>
	KillAllRuns *bool `json:"KillAllRuns,omitnil,omitempty" name:"KillAllRuns"`

	// <p>是否只终止处于等待中（Pending）状态的运行。非必填，默认 false</p>
	OnlyKillPendingRuns *bool `json:"OnlyKillPendingRuns,omitnil,omitempty" name:"OnlyKillPendingRuns"`
}

func (r *KillWorkflowRunRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *KillWorkflowRunRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "WorkflowId")
	delete(f, "WorkflowRunIds")
	delete(f, "KillAllRuns")
	delete(f, "OnlyKillPendingRuns")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "KillWorkflowRunRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type KillWorkflowRunResponseParams struct {
	// <p>终止工作流的运行响应内容</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *AsyncActionRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type KillWorkflowRunResponse struct {
	*tchttp.BaseResponse
	Response *KillWorkflowRunResponseParams `json:"Response"`
}

func (r *KillWorkflowRunResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *KillWorkflowRunResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LabelBrief struct {
	// 标签名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelKey *string `json:"LabelKey,omitnil,omitempty" name:"LabelKey"`

	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelValue *string `json:"LabelValue,omitnil,omitempty" name:"LabelValue"`

	// 标签名称ID，可通过标签相关接口获取
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelKeyId *string `json:"LabelKeyId,omitnil,omitempty" name:"LabelKeyId"`

	// 标签值ID，可通过标签相关接口获取
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelValueId *string `json:"LabelValueId,omitnil,omitempty" name:"LabelValueId"`
}

// Predefined struct for user
type ListCatalogsRequestParams struct {
	// <p>工作空间唯一id</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>最大结果条数</p>
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// <p>分页token</p>
	PageToken *string `json:"PageToken,omitnil,omitempty" name:"PageToken"`

	// <p>catalog类型，可填：TABLE、MODEL、VOLUME</p>
	Types []*string `json:"Types,omitnil,omitempty" name:"Types"`

	// <p>是否只返回当前工作空间有权限的，false时会额外返回当前用户有权限管理的catalog</p>
	CurrentWorkspaceOnly *bool `json:"CurrentWorkspaceOnly,omitnil,omitempty" name:"CurrentWorkspaceOnly"`

	// <p>数据获取选项，可选，控制是否返回权限信息及按权限过滤</p>
	FetchOption *FetchOption `json:"FetchOption,omitnil,omitempty" name:"FetchOption"`

	// 数据目录来源筛选（融合版新增字段），取值参考 CatalogSourceEnum。空数组或不传→只查tclake；["METALAKE"]→只查tclake；["CONNECTION"]→只查分析版ana_catalog；["METALAKE","CONNECTION"]→合并返回
	CatalogSources []*string `json:"CatalogSources,omitnil,omitempty" name:"CatalogSources"`
}

type ListCatalogsRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间唯一id</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>最大结果条数</p>
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// <p>分页token</p>
	PageToken *string `json:"PageToken,omitnil,omitempty" name:"PageToken"`

	// <p>catalog类型，可填：TABLE、MODEL、VOLUME</p>
	Types []*string `json:"Types,omitnil,omitempty" name:"Types"`

	// <p>是否只返回当前工作空间有权限的，false时会额外返回当前用户有权限管理的catalog</p>
	CurrentWorkspaceOnly *bool `json:"CurrentWorkspaceOnly,omitnil,omitempty" name:"CurrentWorkspaceOnly"`

	// <p>数据获取选项，可选，控制是否返回权限信息及按权限过滤</p>
	FetchOption *FetchOption `json:"FetchOption,omitnil,omitempty" name:"FetchOption"`

	// 数据目录来源筛选（融合版新增字段），取值参考 CatalogSourceEnum。空数组或不传→只查tclake；["METALAKE"]→只查tclake；["CONNECTION"]→只查分析版ana_catalog；["METALAKE","CONNECTION"]→合并返回
	CatalogSources []*string `json:"CatalogSources,omitnil,omitempty" name:"CatalogSources"`
}

func (r *ListCatalogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListCatalogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "MaxResults")
	delete(f, "PageToken")
	delete(f, "Types")
	delete(f, "CurrentWorkspaceOnly")
	delete(f, "FetchOption")
	delete(f, "CatalogSources")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListCatalogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListCatalogsResponseParams struct {
	// <p>操作结果</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *ListCatalogsRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListCatalogsResponse struct {
	*tchttp.BaseResponse
	Response *ListCatalogsResponseParams `json:"Response"`
}

func (r *ListCatalogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListCatalogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCatalogsRsp struct {
	// 数据目录列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*Catalog `json:"Items,omitnil,omitempty" name:"Items"`

	// 下页分页token
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextPageToken *string `json:"NextPageToken,omitnil,omitempty" name:"NextPageToken"`
}

// Predefined struct for user
type ListConsoleGroupUsersRequestParams struct {
	// <p>用户组 ID，可通过 ListConsoleGroups 接口获取</p>
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// <p>用户名称或 UIN 模糊匹配</p>
	UserKeyword *string `json:"UserKeyword,omitnil,omitempty" name:"UserKeyword"`

	// <p>通过 UIN 批量查询用户信息</p>
	UserUins []*string `json:"UserUins,omitnil,omitempty" name:"UserUins"`

	// <p>多字段排序，如 [{Name: &#39;CreateTime&#39;, Direction: &#39;Desc&#39;}, {Name: &#39;UserName&#39;, Direction: &#39;Asc&#39;}]，默认按创建时间降序</p>
	OrderBys []*OrderBy `json:"OrderBys,omitnil,omitempty" name:"OrderBys"`

	// <p>页码，从1开始，默认1</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页大小，默认10，最小10，最大100</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type ListConsoleGroupUsersRequest struct {
	*tchttp.BaseRequest
	
	// <p>用户组 ID，可通过 ListConsoleGroups 接口获取</p>
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// <p>用户名称或 UIN 模糊匹配</p>
	UserKeyword *string `json:"UserKeyword,omitnil,omitempty" name:"UserKeyword"`

	// <p>通过 UIN 批量查询用户信息</p>
	UserUins []*string `json:"UserUins,omitnil,omitempty" name:"UserUins"`

	// <p>多字段排序，如 [{Name: &#39;CreateTime&#39;, Direction: &#39;Desc&#39;}, {Name: &#39;UserName&#39;, Direction: &#39;Asc&#39;}]，默认按创建时间降序</p>
	OrderBys []*OrderBy `json:"OrderBys,omitnil,omitempty" name:"OrderBys"`

	// <p>页码，从1开始，默认1</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页大小，默认10，最小10，最大100</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *ListConsoleGroupUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListConsoleGroupUsersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "UserKeyword")
	delete(f, "UserUins")
	delete(f, "OrderBys")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListConsoleGroupUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListConsoleGroupUsersResponseParams struct {
	// <p>返回结果</p>
	Data *ListConsoleGroupUsersRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListConsoleGroupUsersResponse struct {
	*tchttp.BaseResponse
	Response *ListConsoleGroupUsersResponseParams `json:"Response"`
}

func (r *ListConsoleGroupUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListConsoleGroupUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListConsoleGroupUsersRsp struct {
	// 用户组成员列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*ConsoleGroupUserInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// 当前页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 总记录数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 总页数
	TotalPageNumber *int64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`
}

// Predefined struct for user
type ListConsoleGroupsRequestParams struct {
	// <p>页码，从1开始，默认1</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页大小，默认10，最小10，最大100</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>通过用户组ID批量查询；用户组ID可通过 ListConsoleGroups 接口获取</p>
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// <p>用户组名称模糊匹配</p>
	GroupKeyword *string `json:"GroupKeyword,omitnil,omitempty" name:"GroupKeyword"`

	// <p>多字段排序，如 [{Name: &#39;CreateTime&#39;, Direction: &#39;Desc&#39;}, {Name: &#39;UserName&#39;, Direction: &#39;Asc&#39;}]，默认按创建时间降序</p>
	OrderBys []*OrderBy `json:"OrderBys,omitnil,omitempty" name:"OrderBys"`
}

type ListConsoleGroupsRequest struct {
	*tchttp.BaseRequest
	
	// <p>页码，从1开始，默认1</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页大小，默认10，最小10，最大100</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>通过用户组ID批量查询；用户组ID可通过 ListConsoleGroups 接口获取</p>
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// <p>用户组名称模糊匹配</p>
	GroupKeyword *string `json:"GroupKeyword,omitnil,omitempty" name:"GroupKeyword"`

	// <p>多字段排序，如 [{Name: &#39;CreateTime&#39;, Direction: &#39;Desc&#39;}, {Name: &#39;UserName&#39;, Direction: &#39;Asc&#39;}]，默认按创建时间降序</p>
	OrderBys []*OrderBy `json:"OrderBys,omitnil,omitempty" name:"OrderBys"`
}

func (r *ListConsoleGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListConsoleGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "GroupIds")
	delete(f, "GroupKeyword")
	delete(f, "OrderBys")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListConsoleGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListConsoleGroupsResponseParams struct {
	// <p>返回结果</p>
	Data *ListConsoleGroupsRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListConsoleGroupsResponse struct {
	*tchttp.BaseResponse
	Response *ListConsoleGroupsResponseParams `json:"Response"`
}

func (r *ListConsoleGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListConsoleGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListConsoleGroupsRsp struct {
	// 用户组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*ConsoleGroupInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// 当前页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 总记录数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 总页数
	TotalPageNumber *int64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`
}

// Predefined struct for user
type ListConsoleRolesRequestParams struct {
	// <p>页码，从1开始，默认1</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页大小，默认10，最小10，最大100</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>角色名称或描述模糊匹配</p>
	RoleKeyword *string `json:"RoleKeyword,omitnil,omitempty" name:"RoleKeyword"`
}

type ListConsoleRolesRequest struct {
	*tchttp.BaseRequest
	
	// <p>页码，从1开始，默认1</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页大小，默认10，最小10，最大100</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>角色名称或描述模糊匹配</p>
	RoleKeyword *string `json:"RoleKeyword,omitnil,omitempty" name:"RoleKeyword"`
}

func (r *ListConsoleRolesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListConsoleRolesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "RoleKeyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListConsoleRolesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListConsoleRolesResponseParams struct {
	// <p>返回结果</p>
	Data *ListConsoleRolesRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListConsoleRolesResponse struct {
	*tchttp.BaseResponse
	Response *ListConsoleRolesResponseParams `json:"Response"`
}

func (r *ListConsoleRolesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListConsoleRolesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListConsoleRolesRsp struct {
	// 角色列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*ConsoleRoleInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// 当前页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 总记录数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 总页数
	TotalPageNumber *int64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`
}

// Predefined struct for user
type ListConsoleUsersRequestParams struct {
	// <p>页码，从1开始，默认1</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页大小，默认10，最小10，最大100</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>用户名称与 UIN 模糊匹配</p>
	UserKeyword *string `json:"UserKeyword,omitnil,omitempty" name:"UserKeyword"`

	// <p>用于过滤角色关联的用户</p><p>枚举值：</p><ul><li>2001： 控制台管理员</li><li>2002： 控制台成员</li></ul><p>可通过 ListConsoleRoles 接口获取</p>
	RoleIds []*string `json:"RoleIds,omitnil,omitempty" name:"RoleIds"`

	// <p>多字段排序，如 [{Name: &#39;CreateTime&#39;, Direction: &#39;Desc&#39;}, {Name: &#39;UserName&#39;, Direction: &#39;Asc&#39;}]，默认按创建时间降序</p>
	OrderBys []*OrderBy `json:"OrderBys,omitnil,omitempty" name:"OrderBys"`
}

type ListConsoleUsersRequest struct {
	*tchttp.BaseRequest
	
	// <p>页码，从1开始，默认1</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页大小，默认10，最小10，最大100</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>用户名称与 UIN 模糊匹配</p>
	UserKeyword *string `json:"UserKeyword,omitnil,omitempty" name:"UserKeyword"`

	// <p>用于过滤角色关联的用户</p><p>枚举值：</p><ul><li>2001： 控制台管理员</li><li>2002： 控制台成员</li></ul><p>可通过 ListConsoleRoles 接口获取</p>
	RoleIds []*string `json:"RoleIds,omitnil,omitempty" name:"RoleIds"`

	// <p>多字段排序，如 [{Name: &#39;CreateTime&#39;, Direction: &#39;Desc&#39;}, {Name: &#39;UserName&#39;, Direction: &#39;Asc&#39;}]，默认按创建时间降序</p>
	OrderBys []*OrderBy `json:"OrderBys,omitnil,omitempty" name:"OrderBys"`
}

func (r *ListConsoleUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListConsoleUsersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "UserKeyword")
	delete(f, "RoleIds")
	delete(f, "OrderBys")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListConsoleUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListConsoleUsersResponseParams struct {
	// <p>控制台用户列表</p>
	Data *ListConsoleUsersRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListConsoleUsersResponse struct {
	*tchttp.BaseResponse
	Response *ListConsoleUsersResponseParams `json:"Response"`
}

func (r *ListConsoleUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListConsoleUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListConsoleUsersRsp struct {
	// 用户列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*ConsoleUserInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// 当前页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 总记录数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 总页数
	TotalPageNumber *int64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`
}

// Predefined struct for user
type ListFilesRequestParams struct {
	// <p>工作空间id</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>父目录，不填默认查询根节点</p>
	Parent *FolderLocator `json:"Parent,omitnil,omitempty" name:"Parent"`

	// <p>按文件类型过滤</p>
	FileTypes []*string `json:"FileTypes,omitnil,omitempty" name:"FileTypes"`

	// <p>文件名模糊匹配</p>
	NameKeyword *string `json:"NameKeyword,omitnil,omitempty" name:"NameKeyword"`

	// <p>按所有者UIN过滤，多值为或关系</p>
	OwnerUserUins []*string `json:"OwnerUserUins,omitnil,omitempty" name:"OwnerUserUins"`

	// <p>是否只列出文件夹，默认 false</p>
	OnlyFolder *bool `json:"OnlyFolder,omitnil,omitempty" name:"OnlyFolder"`

	// <p>排序字段列表，如创建时间 [{Name: &#39;CreateTime&#39;, Direction: &#39;DESC&#39;}]，文件名称 [{Name: &#39;Name&#39;, Direction: &#39;ASC&#39;}]</p>
	OrderBys []*OrderBy `json:"OrderBys,omitnil,omitempty" name:"OrderBys"`

	// <p>页码，默认1，最小值1</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页条数，默认10，最小值10，最大值100</p><p>取值范围：[10, 100]</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type ListFilesRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间id</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>父目录，不填默认查询根节点</p>
	Parent *FolderLocator `json:"Parent,omitnil,omitempty" name:"Parent"`

	// <p>按文件类型过滤</p>
	FileTypes []*string `json:"FileTypes,omitnil,omitempty" name:"FileTypes"`

	// <p>文件名模糊匹配</p>
	NameKeyword *string `json:"NameKeyword,omitnil,omitempty" name:"NameKeyword"`

	// <p>按所有者UIN过滤，多值为或关系</p>
	OwnerUserUins []*string `json:"OwnerUserUins,omitnil,omitempty" name:"OwnerUserUins"`

	// <p>是否只列出文件夹，默认 false</p>
	OnlyFolder *bool `json:"OnlyFolder,omitnil,omitempty" name:"OnlyFolder"`

	// <p>排序字段列表，如创建时间 [{Name: &#39;CreateTime&#39;, Direction: &#39;DESC&#39;}]，文件名称 [{Name: &#39;Name&#39;, Direction: &#39;ASC&#39;}]</p>
	OrderBys []*OrderBy `json:"OrderBys,omitnil,omitempty" name:"OrderBys"`

	// <p>页码，默认1，最小值1</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页条数，默认10，最小值10，最大值100</p><p>取值范围：[10, 100]</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *ListFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "Parent")
	delete(f, "FileTypes")
	delete(f, "NameKeyword")
	delete(f, "OwnerUserUins")
	delete(f, "OnlyFolder")
	delete(f, "OrderBys")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListFilesResponseParams struct {
	// <p>文件列表结果</p>
	Data *ListFilesRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListFilesResponse struct {
	*tchttp.BaseResponse
	Response *ListFilesResponseParams `json:"Response"`
}

func (r *ListFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListFilesRsp struct {
	// <p>文件/文件夹节点列表</p>
	Items []*FileNode `json:"Items,omitnil,omitempty" name:"Items"`

	// <p>当前页码</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页条数</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>总条数</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>总页数</p>
	TotalPageNumber *int64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`
}

type ListPartition struct {
	// 分区名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分区列表
	List []*Literal `json:"List,omitnil,omitempty" name:"List"`

	// 属性
	Properties []*KVPair `json:"Properties,omitnil,omitempty" name:"Properties"`
}

type ListPartitioning struct {
	// 分区字段
	FiledNames []*string `json:"FiledNames,omitnil,omitempty" name:"FiledNames"`

	// 分区
	Assignments []*ListPartition `json:"Assignments,omitnil,omitempty" name:"Assignments"`
}

// Predefined struct for user
type ListSchemasRequestParams struct {
	// <p>数据目录名</p>
	CatalogName *string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`

	// <p>调用时所在workspace唯一id</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>最大结果条数</p>
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// <p>分页token</p>
	PageToken *string `json:"PageToken,omitnil,omitempty" name:"PageToken"`

	// <p>数据获取选项，可选，控制是否返回权限信息及按权限过滤</p>
	FetchOption *FetchOption `json:"FetchOption,omitnil,omitempty" name:"FetchOption"`

	// 数据源连接ID，可选（融合版新增字段）。非空→走分析版路径，空/缺省→走专业版TcLake路径
	ConnectionId *string `json:"ConnectionId,omitnil,omitempty" name:"ConnectionId"`

	// 数据目录来源，可选（Setats融合新增字段）。取值：METALAKE（专业版tclake）/ CONNECTION（分析版Linked Catalog）/ SETATS（Setats元数据）。路由优先级：显式非空CatalogSource > 非空ConnectionId（视为CONNECTION）> 默认METALAKE；CatalogSource非空时ConnectionId将被忽略。
	CatalogSource *string `json:"CatalogSource,omitnil,omitempty" name:"CatalogSource"`
}

type ListSchemasRequest struct {
	*tchttp.BaseRequest
	
	// <p>数据目录名</p>
	CatalogName *string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`

	// <p>调用时所在workspace唯一id</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>最大结果条数</p>
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// <p>分页token</p>
	PageToken *string `json:"PageToken,omitnil,omitempty" name:"PageToken"`

	// <p>数据获取选项，可选，控制是否返回权限信息及按权限过滤</p>
	FetchOption *FetchOption `json:"FetchOption,omitnil,omitempty" name:"FetchOption"`

	// 数据源连接ID，可选（融合版新增字段）。非空→走分析版路径，空/缺省→走专业版TcLake路径
	ConnectionId *string `json:"ConnectionId,omitnil,omitempty" name:"ConnectionId"`

	// 数据目录来源，可选（Setats融合新增字段）。取值：METALAKE（专业版tclake）/ CONNECTION（分析版Linked Catalog）/ SETATS（Setats元数据）。路由优先级：显式非空CatalogSource > 非空ConnectionId（视为CONNECTION）> 默认METALAKE；CatalogSource非空时ConnectionId将被忽略。
	CatalogSource *string `json:"CatalogSource,omitnil,omitempty" name:"CatalogSource"`
}

func (r *ListSchemasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSchemasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CatalogName")
	delete(f, "WorkspaceId")
	delete(f, "MaxResults")
	delete(f, "PageToken")
	delete(f, "FetchOption")
	delete(f, "ConnectionId")
	delete(f, "CatalogSource")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListSchemasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListSchemasResponseParams struct {
	// <p>操作结果</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *ListSchemasRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListSchemasResponse struct {
	*tchttp.BaseResponse
	Response *ListSchemasResponseParams `json:"Response"`
}

func (r *ListSchemasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSchemasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSchemasRsp struct {
	// schema列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*Schema `json:"Items,omitnil,omitempty" name:"Items"`

	// 下页分页token
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextPageToken *string `json:"NextPageToken,omitnil,omitempty" name:"NextPageToken"`
}

// Predefined struct for user
type ListTablesRequestParams struct {
	// <p>数据目录名</p>
	CatalogName *string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`

	// <p>数据库名</p>
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// <p>调用时所在workspace唯一id</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>最大结果条数</p>
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// <p>分页token</p>
	PageToken *string `json:"PageToken,omitnil,omitempty" name:"PageToken"`

	// <p>数据获取选项，可选，控制是否返回权限信息及按权限过滤</p>
	FetchOption *FetchOption `json:"FetchOption,omitnil,omitempty" name:"FetchOption"`

	// 数据源连接ID，可选（融合版新增字段）。非空→走分析版路径，空/缺省→走专业版TcLake路径
	ConnectionId *string `json:"ConnectionId,omitnil,omitempty" name:"ConnectionId"`

	// 数据目录来源，可选（Setats融合新增字段）。取值：METALAKE（专业版tclake）/ CONNECTION（分析版Linked Catalog）/ SETATS（Setats元数据）。路由优先级：显式非空CatalogSource > 非空ConnectionId（视为CONNECTION）> 默认METALAKE；CatalogSource非空时ConnectionId将被忽略。
	CatalogSource *string `json:"CatalogSource,omitnil,omitempty" name:"CatalogSource"`
}

type ListTablesRequest struct {
	*tchttp.BaseRequest
	
	// <p>数据目录名</p>
	CatalogName *string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`

	// <p>数据库名</p>
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// <p>调用时所在workspace唯一id</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>最大结果条数</p>
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// <p>分页token</p>
	PageToken *string `json:"PageToken,omitnil,omitempty" name:"PageToken"`

	// <p>数据获取选项，可选，控制是否返回权限信息及按权限过滤</p>
	FetchOption *FetchOption `json:"FetchOption,omitnil,omitempty" name:"FetchOption"`

	// 数据源连接ID，可选（融合版新增字段）。非空→走分析版路径，空/缺省→走专业版TcLake路径
	ConnectionId *string `json:"ConnectionId,omitnil,omitempty" name:"ConnectionId"`

	// 数据目录来源，可选（Setats融合新增字段）。取值：METALAKE（专业版tclake）/ CONNECTION（分析版Linked Catalog）/ SETATS（Setats元数据）。路由优先级：显式非空CatalogSource > 非空ConnectionId（视为CONNECTION）> 默认METALAKE；CatalogSource非空时ConnectionId将被忽略。
	CatalogSource *string `json:"CatalogSource,omitnil,omitempty" name:"CatalogSource"`
}

func (r *ListTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CatalogName")
	delete(f, "SchemaName")
	delete(f, "WorkspaceId")
	delete(f, "MaxResults")
	delete(f, "PageToken")
	delete(f, "FetchOption")
	delete(f, "ConnectionId")
	delete(f, "CatalogSource")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTablesResponseParams struct {
	// <p>操作结果</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *ListTablesRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListTablesResponse struct {
	*tchttp.BaseResponse
	Response *ListTablesResponseParams `json:"Response"`
}

func (r *ListTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListTablesRsp struct {
	// tabl列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*TableInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// 下页分页token
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextPageToken *string `json:"NextPageToken,omitnil,omitempty" name:"NextPageToken"`
}

// Predefined struct for user
type ListWorkflowRunsRequestParams struct {
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>分页页码，从 1 开始。非必填，默认 1</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页大小。非必填，默认 10，取值范围 [10, 200]</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>工作流ID，精确匹配。非必填，单值</p>
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// <p>工作流名称关键字，对 WorkflowName 做模糊匹配。非必填，单值</p>
	WorkflowNameKeyword *string `json:"WorkflowNameKeyword,omitnil,omitempty" name:"WorkflowNameKeyword"`

	// <p>运行创建时间下界，范围匹配（CreateTime &gt;= 本值），单位：毫秒时间戳。<br>非必填，单值，对应出参 WorkflowRun.CreateTime</p>
	CreateStartTime *string `json:"CreateStartTime,omitnil,omitempty" name:"CreateStartTime"`

	// <p>运行创建时间上界，范围匹配（CreateTime &lt;= 本值），单位：毫秒时间戳。<br>非必填，单值，对应出参 WorkflowRun.CreateTime</p>
	CreateEndTime *string `json:"CreateEndTime,omitnil,omitempty" name:"CreateEndTime"`

	// <p>运行状态，精确匹配。非必填，多选（多个值之间为 OR 关系）。运行状态。CREATE(&quot;初始化&quot;), QUEUED(&quot;等待中&quot;), PENDING(&quot;准备中&quot;), RUNNING(&quot;运行中&quot;), SKIPPED(&quot;跳过运行&quot;), SUCCESS(&quot;成功&quot;), FAILED(&quot;失败&quot;), TERMINATING(&quot;终止中&quot;), TERMINATED(&quot;终止&quot;), CANCELLED(&quot;被手动终止&quot;)等</p>
	RunStates []*string `json:"RunStates,omitnil,omitempty" name:"RunStates"`

	// <p>错误码，精确匹配。非必填，多选（多个值之间为 OR 关系）</p>
	ErrorCodeStrings []*string `json:"ErrorCodeStrings,omitnil,omitempty" name:"ErrorCodeStrings"`

	// <p>运行人UIN，精确匹配。非必填，多选（多个值之间为 OR 关系）</p>
	RunUserUins []*string `json:"RunUserUins,omitnil,omitempty" name:"RunUserUins"`

	// <p>标签名称ID，精确匹配，可通过标签相关接口获取。非必填，多选（多个值之间为 OR 关系）</p>
	LabelKeyIds []*string `json:"LabelKeyIds,omitnil,omitempty" name:"LabelKeyIds"`

	// <p>标签值ID，精确匹配，可通过标签相关接口获取。非必填，多选（多个值之间为 OR 关系）</p>
	LabelValueIds []*string `json:"LabelValueIds,omitnil,omitempty" name:"LabelValueIds"`

	// <p>排序条件，多个之间按数组顺序表示优先级。非必填，默认按 CreateTime Desc。<br>可排序字段白名单：CreateTime、EndTime、RunCostTime</p>
	OrderBys []*OrderBy `json:"OrderBys,omitnil,omitempty" name:"OrderBys"`
}

type ListWorkflowRunsRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>分页页码，从 1 开始。非必填，默认 1</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页大小。非必填，默认 10，取值范围 [10, 200]</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>工作流ID，精确匹配。非必填，单值</p>
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// <p>工作流名称关键字，对 WorkflowName 做模糊匹配。非必填，单值</p>
	WorkflowNameKeyword *string `json:"WorkflowNameKeyword,omitnil,omitempty" name:"WorkflowNameKeyword"`

	// <p>运行创建时间下界，范围匹配（CreateTime &gt;= 本值），单位：毫秒时间戳。<br>非必填，单值，对应出参 WorkflowRun.CreateTime</p>
	CreateStartTime *string `json:"CreateStartTime,omitnil,omitempty" name:"CreateStartTime"`

	// <p>运行创建时间上界，范围匹配（CreateTime &lt;= 本值），单位：毫秒时间戳。<br>非必填，单值，对应出参 WorkflowRun.CreateTime</p>
	CreateEndTime *string `json:"CreateEndTime,omitnil,omitempty" name:"CreateEndTime"`

	// <p>运行状态，精确匹配。非必填，多选（多个值之间为 OR 关系）。运行状态。CREATE(&quot;初始化&quot;), QUEUED(&quot;等待中&quot;), PENDING(&quot;准备中&quot;), RUNNING(&quot;运行中&quot;), SKIPPED(&quot;跳过运行&quot;), SUCCESS(&quot;成功&quot;), FAILED(&quot;失败&quot;), TERMINATING(&quot;终止中&quot;), TERMINATED(&quot;终止&quot;), CANCELLED(&quot;被手动终止&quot;)等</p>
	RunStates []*string `json:"RunStates,omitnil,omitempty" name:"RunStates"`

	// <p>错误码，精确匹配。非必填，多选（多个值之间为 OR 关系）</p>
	ErrorCodeStrings []*string `json:"ErrorCodeStrings,omitnil,omitempty" name:"ErrorCodeStrings"`

	// <p>运行人UIN，精确匹配。非必填，多选（多个值之间为 OR 关系）</p>
	RunUserUins []*string `json:"RunUserUins,omitnil,omitempty" name:"RunUserUins"`

	// <p>标签名称ID，精确匹配，可通过标签相关接口获取。非必填，多选（多个值之间为 OR 关系）</p>
	LabelKeyIds []*string `json:"LabelKeyIds,omitnil,omitempty" name:"LabelKeyIds"`

	// <p>标签值ID，精确匹配，可通过标签相关接口获取。非必填，多选（多个值之间为 OR 关系）</p>
	LabelValueIds []*string `json:"LabelValueIds,omitnil,omitempty" name:"LabelValueIds"`

	// <p>排序条件，多个之间按数组顺序表示优先级。非必填，默认按 CreateTime Desc。<br>可排序字段白名单：CreateTime、EndTime、RunCostTime</p>
	OrderBys []*OrderBy `json:"OrderBys,omitnil,omitempty" name:"OrderBys"`
}

func (r *ListWorkflowRunsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListWorkflowRunsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "WorkflowId")
	delete(f, "WorkflowNameKeyword")
	delete(f, "CreateStartTime")
	delete(f, "CreateEndTime")
	delete(f, "RunStates")
	delete(f, "ErrorCodeStrings")
	delete(f, "RunUserUins")
	delete(f, "LabelKeyIds")
	delete(f, "LabelValueIds")
	delete(f, "OrderBys")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListWorkflowRunsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListWorkflowRunsResponseParams struct {
	// <p>工作流运行列表响应内容</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *ListWorkflowRunsRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListWorkflowRunsResponse struct {
	*tchttp.BaseResponse
	Response *ListWorkflowRunsResponseParams `json:"Response"`
}

func (r *ListWorkflowRunsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListWorkflowRunsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListWorkflowRunsRsp struct {
	// 当前页码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 总记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 总页数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPageNumber *int64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`

	// 工作流运行列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*WorkflowRun `json:"Items,omitnil,omitempty" name:"Items"`

	// 工作流运行状态数量统计。
	// 统计口径为当前筛选条件下的全量数据，不受 PageNumber / PageSize 影响
	// 注意：此字段可能返回 null，表示取不到有效值。
	BizStateEnumInfos []*ScheduleBizEnumBrief `json:"BizStateEnumInfos,omitnil,omitempty" name:"BizStateEnumInfos"`

	// 工作流运行错误码数量统计。
	// 统计口径为当前筛选条件下的全量数据，不受 PageNumber / PageSize 影响
	// 注意：此字段可能返回 null，表示取不到有效值。
	BizErrorCodeEnumInfos []*ScheduleBizEnumBrief `json:"BizErrorCodeEnumInfos,omitnil,omitempty" name:"BizErrorCodeEnumInfos"`
}

// Predefined struct for user
type ListWorkflowTaskRunsRequestParams struct {
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>任务ID，可通过 ListWorkflowTasks 获取。非必填，精确匹配。与 WorkflowRunId 至少传一个：仅传 TaskId 时查询该任务的全部运行历史。</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <p>工作流运行ID，可通过 ListWorkflowRuns 获取。非必填，精确匹配。与 TaskId 至少传一个：仅传 WorkflowRunId 时查询该次工作流运行下的全部任务运行。</p>
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`

	// <p>分页页码，从 1 开始。非必填，默认 1</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页大小。非必填，默认 10，取值范围 [10, 200]</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type ListWorkflowTaskRunsRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>任务ID，可通过 ListWorkflowTasks 获取。非必填，精确匹配。与 WorkflowRunId 至少传一个：仅传 TaskId 时查询该任务的全部运行历史。</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <p>工作流运行ID，可通过 ListWorkflowRuns 获取。非必填，精确匹配。与 TaskId 至少传一个：仅传 WorkflowRunId 时查询该次工作流运行下的全部任务运行。</p>
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`

	// <p>分页页码，从 1 开始。非必填，默认 1</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页大小。非必填，默认 10，取值范围 [10, 200]</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *ListWorkflowTaskRunsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListWorkflowTaskRunsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "TaskId")
	delete(f, "WorkflowRunId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListWorkflowTaskRunsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListWorkflowTaskRunsResponseParams struct {
	// <p>查询工作流任务历史运行列表响应内容</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *ListWorkflowTaskRunsRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListWorkflowTaskRunsResponse struct {
	*tchttp.BaseResponse
	Response *ListWorkflowTaskRunsResponseParams `json:"Response"`
}

func (r *ListWorkflowTaskRunsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListWorkflowTaskRunsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListWorkflowTaskRunsRsp struct {
	// 当前页码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 总记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 总页数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPageNumber *int64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`

	// 任务运行历史列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*WorkflowTaskRun `json:"Items,omitnil,omitempty" name:"Items"`
}

// Predefined struct for user
type ListWorkflowsRequestParams struct {
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>分页页码，从 1 开始。非必填，默认 1</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页大小。非必填，默认 10，取值范围 [10, 200]</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>工作流名称关键字，对 WorkflowName 做模糊匹配。非必填，单值</p>
	WorkflowNameKeyword *string `json:"WorkflowNameKeyword,omitnil,omitempty" name:"WorkflowNameKeyword"`

	// <p>工作流名称，精确匹配。非必填，多选（多个值之间为 OR 关系）</p>
	WorkflowNames []*string `json:"WorkflowNames,omitnil,omitempty" name:"WorkflowNames"`

	// <p>工作流ID，精确匹配。非必填，多选（多个值之间为 OR 关系）</p>
	WorkflowIds []*string `json:"WorkflowIds,omitnil,omitempty" name:"WorkflowIds"`

	// <p>工作流运行人UIN，精确匹配。非必填，多选（多个值之间为 OR 关系）</p>
	RunUserUins []*string `json:"RunUserUins,omitnil,omitempty" name:"RunUserUins"`

	// <p>标签名称ID，精确匹配，可通过标签相关接口获取。非必填，多选（多个值之间为 OR 关系）</p>
	LabelKeyIds []*string `json:"LabelKeyIds,omitnil,omitempty" name:"LabelKeyIds"`

	// <p>标签值ID，精确匹配，可通过标签相关接口获取。非必填，多选（多个值之间为 OR 关系）</p>
	LabelValueIds []*string `json:"LabelValueIds,omitnil,omitempty" name:"LabelValueIds"`

	// <p>快速筛选类型。非必填，单值</p><p>对齐老云 API（wedata/2025-10-10）文档示例值：</p><ul><li>MY_FAVORITE：我收藏的</li><li>MY_OWNER：我负责的</li><li>MY_AUTHORITY：我有权限</li><li>WorkflowId：支持多个工作流ID筛选</li></ul><p>后端实现现状：当前仅 MY_FAVORITE 生效（设置 favoriteUserUin 过滤当前用户收藏），MY_OWNER / MY_AUTHORITY 暂未在 Service 层实现，传入会被忽略（按全量返回）。</p>
	QuickSelectionType *string `json:"QuickSelectionType,omitnil,omitempty" name:"QuickSelectionType"`

	// <p>排序条件，多个之间按数组顺序表示优先级。非必填。<br>可排序字段白名单：CreateTime</p>
	OrderBys []*OrderBy `json:"OrderBys,omitnil,omitempty" name:"OrderBys"`
}

type ListWorkflowsRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>分页页码，从 1 开始。非必填，默认 1</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页大小。非必填，默认 10，取值范围 [10, 200]</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>工作流名称关键字，对 WorkflowName 做模糊匹配。非必填，单值</p>
	WorkflowNameKeyword *string `json:"WorkflowNameKeyword,omitnil,omitempty" name:"WorkflowNameKeyword"`

	// <p>工作流名称，精确匹配。非必填，多选（多个值之间为 OR 关系）</p>
	WorkflowNames []*string `json:"WorkflowNames,omitnil,omitempty" name:"WorkflowNames"`

	// <p>工作流ID，精确匹配。非必填，多选（多个值之间为 OR 关系）</p>
	WorkflowIds []*string `json:"WorkflowIds,omitnil,omitempty" name:"WorkflowIds"`

	// <p>工作流运行人UIN，精确匹配。非必填，多选（多个值之间为 OR 关系）</p>
	RunUserUins []*string `json:"RunUserUins,omitnil,omitempty" name:"RunUserUins"`

	// <p>标签名称ID，精确匹配，可通过标签相关接口获取。非必填，多选（多个值之间为 OR 关系）</p>
	LabelKeyIds []*string `json:"LabelKeyIds,omitnil,omitempty" name:"LabelKeyIds"`

	// <p>标签值ID，精确匹配，可通过标签相关接口获取。非必填，多选（多个值之间为 OR 关系）</p>
	LabelValueIds []*string `json:"LabelValueIds,omitnil,omitempty" name:"LabelValueIds"`

	// <p>快速筛选类型。非必填，单值</p><p>对齐老云 API（wedata/2025-10-10）文档示例值：</p><ul><li>MY_FAVORITE：我收藏的</li><li>MY_OWNER：我负责的</li><li>MY_AUTHORITY：我有权限</li><li>WorkflowId：支持多个工作流ID筛选</li></ul><p>后端实现现状：当前仅 MY_FAVORITE 生效（设置 favoriteUserUin 过滤当前用户收藏），MY_OWNER / MY_AUTHORITY 暂未在 Service 层实现，传入会被忽略（按全量返回）。</p>
	QuickSelectionType *string `json:"QuickSelectionType,omitnil,omitempty" name:"QuickSelectionType"`

	// <p>排序条件，多个之间按数组顺序表示优先级。非必填。<br>可排序字段白名单：CreateTime</p>
	OrderBys []*OrderBy `json:"OrderBys,omitnil,omitempty" name:"OrderBys"`
}

func (r *ListWorkflowsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListWorkflowsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "WorkflowNameKeyword")
	delete(f, "WorkflowNames")
	delete(f, "WorkflowIds")
	delete(f, "RunUserUins")
	delete(f, "LabelKeyIds")
	delete(f, "LabelValueIds")
	delete(f, "QuickSelectionType")
	delete(f, "OrderBys")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListWorkflowsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListWorkflowsResponseParams struct {
	// <p>查询工作流列表响应内容</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *ListWorkflowsRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListWorkflowsResponse struct {
	*tchttp.BaseResponse
	Response *ListWorkflowsResponseParams `json:"Response"`
}

func (r *ListWorkflowsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListWorkflowsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListWorkflowsRsp struct {
	// 当前页码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 总记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 总页数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPageNumber *int64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`

	// 工作流列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*WorkflowBrief `json:"Items,omitnil,omitempty" name:"Items"`
}

// Predefined struct for user
type ListWorkspaceGroupUsersRequestParams struct {
	// <p>工作空间id</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>用户组ID，可通过 ListWorkspaceGroups 接口获取</p>
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// <p>关键字，支持按用户名模糊搜索</p>
	NameKeywords *string `json:"NameKeywords,omitnil,omitempty" name:"NameKeywords"`

	// <p>用户UIN列表，按UIN精确筛选</p>
	UserUins []*string `json:"UserUins,omitnil,omitempty" name:"UserUins"`

	// <p>&quot;OrderBys&quot;: [{&quot;Name&quot;: &quot;CreateTime&quot;, &quot;Direction&quot;: &quot;DESC&quot; }]</p>
	OrderBys []*OrderBy `json:"OrderBys,omitnil,omitempty" name:"OrderBys"`

	// <p>页码，从1开始，默认1</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页大小</p><p>取值范围：[1, 100]</p><p>默认值：10</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type ListWorkspaceGroupUsersRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间id</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>用户组ID，可通过 ListWorkspaceGroups 接口获取</p>
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// <p>关键字，支持按用户名模糊搜索</p>
	NameKeywords *string `json:"NameKeywords,omitnil,omitempty" name:"NameKeywords"`

	// <p>用户UIN列表，按UIN精确筛选</p>
	UserUins []*string `json:"UserUins,omitnil,omitempty" name:"UserUins"`

	// <p>&quot;OrderBys&quot;: [{&quot;Name&quot;: &quot;CreateTime&quot;, &quot;Direction&quot;: &quot;DESC&quot; }]</p>
	OrderBys []*OrderBy `json:"OrderBys,omitnil,omitempty" name:"OrderBys"`

	// <p>页码，从1开始，默认1</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页大小</p><p>取值范围：[1, 100]</p><p>默认值：10</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *ListWorkspaceGroupUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListWorkspaceGroupUsersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "GroupId")
	delete(f, "NameKeywords")
	delete(f, "UserUins")
	delete(f, "OrderBys")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListWorkspaceGroupUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListWorkspaceGroupUsersResponseParams struct {
	// <p>返回结果</p>
	Data *ListWorkspaceGroupUsersRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListWorkspaceGroupUsersResponse struct {
	*tchttp.BaseResponse
	Response *ListWorkspaceGroupUsersResponseParams `json:"Response"`
}

func (r *ListWorkspaceGroupUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListWorkspaceGroupUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListWorkspaceGroupUsersRsp struct {
	// <p>用户列表</p>
	Items []*GroupUserInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// <p>当前页码</p><p>默认值：1</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页大小</p><p>默认值：10</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>总记录数</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>总页数</p>
	TotalPageNumber *int64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`
}

// Predefined struct for user
type ListWorkspaceGroupsRequestParams struct {
	// <p>工作空间ID</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>通过用户组ID批量查询；用户组ID可通过 ListWorkspaceGroups 接口获取</p>
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// <p>用户组名称模糊匹配</p>
	GroupKeyword *string `json:"GroupKeyword,omitnil,omitempty" name:"GroupKeyword"`

	// <p>用于过滤角色关联的用户组；可通过 ListWorkspaceRoles 接口获取</p>
	RoleIds []*string `json:"RoleIds,omitnil,omitempty" name:"RoleIds"`

	// <p>多字段排序，如 [{Name: 'CreateTime', Direction: 'Desc'}]；传入单个即单字段排序，默认按创建时间降序</p>
	OrderBys []*OrderBy `json:"OrderBys,omitnil,omitempty" name:"OrderBys"`

	// <p>页码，从1开始，默认1；小于1时按1处理</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页大小，默认10，取值范围为 [1,100]，越界时自动收敛到边界值</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>用户组类型筛选。取值为枚举数值的字符串形式："0"=控制台系统类型（包含全部user）、"1"=控制台自定义类型、"2"=工作空间系统类型（包含工作空间全部user）、"3"=工作空间自定义类型</p>
	GroupType *string `json:"GroupType,omitnil,omitempty" name:"GroupType"`
}

type ListWorkspaceGroupsRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间ID</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>通过用户组ID批量查询；用户组ID可通过 ListWorkspaceGroups 接口获取</p>
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// <p>用户组名称模糊匹配</p>
	GroupKeyword *string `json:"GroupKeyword,omitnil,omitempty" name:"GroupKeyword"`

	// <p>用于过滤角色关联的用户组；可通过 ListWorkspaceRoles 接口获取</p>
	RoleIds []*string `json:"RoleIds,omitnil,omitempty" name:"RoleIds"`

	// <p>多字段排序，如 [{Name: 'CreateTime', Direction: 'Desc'}]；传入单个即单字段排序，默认按创建时间降序</p>
	OrderBys []*OrderBy `json:"OrderBys,omitnil,omitempty" name:"OrderBys"`

	// <p>页码，从1开始，默认1；小于1时按1处理</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页大小，默认10，取值范围为 [1,100]，越界时自动收敛到边界值</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>用户组类型筛选。取值为枚举数值的字符串形式："0"=控制台系统类型（包含全部user）、"1"=控制台自定义类型、"2"=工作空间系统类型（包含工作空间全部user）、"3"=工作空间自定义类型</p>
	GroupType *string `json:"GroupType,omitnil,omitempty" name:"GroupType"`
}

func (r *ListWorkspaceGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListWorkspaceGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "GroupIds")
	delete(f, "GroupKeyword")
	delete(f, "RoleIds")
	delete(f, "OrderBys")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "GroupType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListWorkspaceGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListWorkspaceGroupsResponseParams struct {
	// <p>工作空间用户组列表</p>
	Data *ListWorkspaceGroupsRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListWorkspaceGroupsResponse struct {
	*tchttp.BaseResponse
	Response *ListWorkspaceGroupsResponseParams `json:"Response"`
}

func (r *ListWorkspaceGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListWorkspaceGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListWorkspaceGroupsRsp struct {
	// 用户组列表
	Items []*WorkspaceGroupInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// 当前页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 总记录数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 总页数
	TotalPageNumber *int64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`
}

// Predefined struct for user
type ListWorkspaceUsersRequestParams struct {
	// <p>工作空间ID</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>用户名称或 UIN 模糊匹配</p>
	UserKeyword *string `json:"UserKeyword,omitnil,omitempty" name:"UserKeyword"`

	// <p>角色ID列表，过滤角色关联的用户；可通过 ListWorkspaceRoles 接口获取</p>
	RoleIds []*string `json:"RoleIds,omitnil,omitempty" name:"RoleIds"`

	// <p>用户UIN列表，按UIN精确筛选</p>
	UserUins []*string `json:"UserUins,omitnil,omitempty" name:"UserUins"`

	// <p>排序字段列表，如 [{Name: CreateTime, Direction: DESC}]；传入单个即单字段排序，默认按创建时间降序</p>
	OrderBys []*OrderBy `json:"OrderBys,omitnil,omitempty" name:"OrderBys"`

	// <p>页码，默认1，从1开始；小于1时按1处理</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页条数，默认10，取值范围为 [1,100]，越界时自动收敛到边界值</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>扩展信息（首次请求为空，翻页请求带响应中的ExtendInfo。连续翻页且item涉及权限过滤等场景，用于定位下次offset）</p>
	ExtendInfo *string `json:"ExtendInfo,omitnil,omitempty" name:"ExtendInfo"`
}

type ListWorkspaceUsersRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间ID</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>用户名称或 UIN 模糊匹配</p>
	UserKeyword *string `json:"UserKeyword,omitnil,omitempty" name:"UserKeyword"`

	// <p>角色ID列表，过滤角色关联的用户；可通过 ListWorkspaceRoles 接口获取</p>
	RoleIds []*string `json:"RoleIds,omitnil,omitempty" name:"RoleIds"`

	// <p>用户UIN列表，按UIN精确筛选</p>
	UserUins []*string `json:"UserUins,omitnil,omitempty" name:"UserUins"`

	// <p>排序字段列表，如 [{Name: CreateTime, Direction: DESC}]；传入单个即单字段排序，默认按创建时间降序</p>
	OrderBys []*OrderBy `json:"OrderBys,omitnil,omitempty" name:"OrderBys"`

	// <p>页码，默认1，从1开始；小于1时按1处理</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页条数，默认10，取值范围为 [1,100]，越界时自动收敛到边界值</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>扩展信息（首次请求为空，翻页请求带响应中的ExtendInfo。连续翻页且item涉及权限过滤等场景，用于定位下次offset）</p>
	ExtendInfo *string `json:"ExtendInfo,omitnil,omitempty" name:"ExtendInfo"`
}

func (r *ListWorkspaceUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListWorkspaceUsersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "UserKeyword")
	delete(f, "RoleIds")
	delete(f, "UserUins")
	delete(f, "OrderBys")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "ExtendInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListWorkspaceUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListWorkspaceUsersResponseParams struct {
	// <p>查询工作空间用户列表响应</p>
	Data *ListWorkspaceUsersRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListWorkspaceUsersResponse struct {
	*tchttp.BaseResponse
	Response *ListWorkspaceUsersResponseParams `json:"Response"`
}

func (r *ListWorkspaceUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListWorkspaceUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListWorkspaceUsersRsp struct {
	// 用户列表
	Items []*WorkspaceUserInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// 当前页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 总记录数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 总页数
	TotalPageNumber *int64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`

	// 扩展信息（连续翻页场景由服务端返回，下次翻页请求原样带上）
	ExtendInfo *string `json:"ExtendInfo,omitnil,omitempty" name:"ExtendInfo"`
}

// Predefined struct for user
type ListWorkspacesRequestParams struct {
	// <p>工作空间ID精确匹配</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>工作空间名称模糊匹配</p>
	WorkspaceKeyword *string `json:"WorkspaceKeyword,omitnil,omitempty" name:"WorkspaceKeyword"`

	// <p>工作空间状态过滤（多选）：0=未指定 1=创建中 2=创建失败 3=正常运行中 4=已删除</p>
	StatusList []*int64 `json:"StatusList,omitnil,omitempty" name:"StatusList"`

	// <p>多字段排序，如 [{Name: 'CreateTime', Direction: 'Desc'}]；传入单个即单字段排序，默认按创建时间降序</p>
	OrderBys []*OrderBy `json:"OrderBys,omitnil,omitempty" name:"OrderBys"`

	// <p>页码，从1开始，默认1</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页大小，默认10，最小10，最大100</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>工作空间地域过滤（多选），如 ap-guangzhou</p>
	WorkspaceRegion []*string `json:"WorkspaceRegion,omitnil,omitempty" name:"WorkspaceRegion"`

	// <p>创建者UIN过滤（多选）</p>
	Creator []*string `json:"Creator,omitnil,omitempty" name:"Creator"`
}

type ListWorkspacesRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间ID精确匹配</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>工作空间名称模糊匹配</p>
	WorkspaceKeyword *string `json:"WorkspaceKeyword,omitnil,omitempty" name:"WorkspaceKeyword"`

	// <p>工作空间状态过滤（多选）：0=未指定 1=创建中 2=创建失败 3=正常运行中 4=已删除</p>
	StatusList []*int64 `json:"StatusList,omitnil,omitempty" name:"StatusList"`

	// <p>多字段排序，如 [{Name: 'CreateTime', Direction: 'Desc'}]；传入单个即单字段排序，默认按创建时间降序</p>
	OrderBys []*OrderBy `json:"OrderBys,omitnil,omitempty" name:"OrderBys"`

	// <p>页码，从1开始，默认1</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页大小，默认10，最小10，最大100</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>工作空间地域过滤（多选），如 ap-guangzhou</p>
	WorkspaceRegion []*string `json:"WorkspaceRegion,omitnil,omitempty" name:"WorkspaceRegion"`

	// <p>创建者UIN过滤（多选）</p>
	Creator []*string `json:"Creator,omitnil,omitempty" name:"Creator"`
}

func (r *ListWorkspacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListWorkspacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "WorkspaceKeyword")
	delete(f, "StatusList")
	delete(f, "OrderBys")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "WorkspaceRegion")
	delete(f, "Creator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListWorkspacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListWorkspacesResponseParams struct {
	// <p>工作空间列表</p>
	Data *ListWorkspacesRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListWorkspacesResponse struct {
	*tchttp.BaseResponse
	Response *ListWorkspacesResponseParams `json:"Response"`
}

func (r *ListWorkspacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListWorkspacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListWorkspacesRsp struct {
	// <p>工作空间列表</p>
	Items []*WorkspaceInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// <p>当前页码</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页大小</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>总记录数</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>总页数</p>
	TotalPageNumber *int64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`

	// <p>是否控制台管理员</p>
	IsConsoleAdmin *bool `json:"IsConsoleAdmin,omitnil,omitempty" name:"IsConsoleAdmin"`
}

type Literal struct {
	// 数值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 类型
	DataType *string `json:"DataType,omitnil,omitempty" name:"DataType"`
}

type MetaOwner struct {
	// 元数据名称（全名）:catalog.schema.table
	// 注意：此字段可能返回 null，表示取不到有效值。
	FullName *string `json:"FullName,omitnil,omitempty" name:"FullName"`

	// 所有者类型:User
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerType *string `json:"OwnerType,omitnil,omitempty" name:"OwnerType"`

	// 所有者:唯一标识(uin)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Owner *string `json:"Owner,omitnil,omitempty" name:"Owner"`

	// 所有者名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerName *string `json:"OwnerName,omitnil,omitempty" name:"OwnerName"`
}

type MonitorMetricBrief struct {
	// 监控指标 ID，创建时无需传入，由服务端生成
	// 注意：此字段可能返回 null，表示取不到有效值。
	MonitorMetricId *string `json:"MonitorMetricId,omitnil,omitempty" name:"MonitorMetricId"`

	// 告警的监控对象类型，如工作流、任务等，当前支持 1. WORKFLOW 2. TASK
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmMonitorType *string `json:"AlarmMonitorType,omitnil,omitempty" name:"AlarmMonitorType"`

	// 监控指标列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Metrics []*MonitorMetricItem `json:"Metrics,omitnil,omitempty" name:"Metrics"`
}

type MonitorMetricItem struct {
	// 监控指标类型,有三种类型：1. RUN_DURATION（运行时长）2. WAIT_DURATION（等待时长）3. COMPLETION_TIME（完成时间）
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricType *string `json:"MetricType,omitnil,omitempty" name:"MetricType"`

	// 警告阈值，单位为毫秒级别，对于COMPLETION_TIME:从当日时间点00:00起算
	// 注意：此字段可能返回 null，表示取不到有效值。
	WarningThreshold *string `json:"WarningThreshold,omitnil,omitempty" name:"WarningThreshold"`

	// 超时阈值，单位为毫秒级别，对于COMPLETION_TIME:从当日时间点00:00起算
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeoutThreshold *string `json:"TimeoutThreshold,omitnil,omitempty" name:"TimeoutThreshold"`
}

type OrderBy struct {
	// 排序方向，Asc（升序）或 Desc（降序），大小写不敏感
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 排序字段名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type ParamInfo struct {
	// 参数ID，创建时无需传入，由服务端生成
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamId *string `json:"ParamId,omitnil,omitempty" name:"ParamId"`

	// 参数名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamKey *string `json:"ParamKey,omitnil,omitempty" name:"ParamKey"`

	// 参数值
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamValue *string `json:"ParamValue,omitnil,omitempty" name:"ParamValue"`
}

type Partitioning struct {
	// 转换策略
	Strategy *string `json:"Strategy,omitnil,omitempty" name:"Strategy"`

	// 分桶策略参数
	NumBuckets *int64 `json:"NumBuckets,omitnil,omitempty" name:"NumBuckets"`

	// 字段名称
	FieldNames []*string `json:"FieldNames,omitnil,omitempty" name:"FieldNames"`

	// 策略参数
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 字段名称
	FieldName []*string `json:"FieldName,omitnil,omitempty" name:"FieldName"`

	// 按年分区策略
	YearPartitioning *SingleFieldPartitioning `json:"YearPartitioning,omitnil,omitempty" name:"YearPartitioning"`

	// 按月分区策略
	MonthPartitioning *SingleFieldPartitioning `json:"MonthPartitioning,omitnil,omitempty" name:"MonthPartitioning"`

	// 按天分区策略
	DayPartitioning *SingleFieldPartitioning `json:"DayPartitioning,omitnil,omitempty" name:"DayPartitioning"`

	// 按小时分区策略
	HourPartitioning *SingleFieldPartitioning `json:"HourPartitioning,omitnil,omitempty" name:"HourPartitioning"`

	// 按字段分区策略
	IdentityPartitioning *SingleFieldPartitioning `json:"IdentityPartitioning,omitnil,omitempty" name:"IdentityPartitioning"`

	// 列表分区策略
	ListPartitioning *ListPartitioning `json:"ListPartitioning,omitnil,omitempty" name:"ListPartitioning"`

	// 范围分区策略
	RangePartitioning *RangePartitioning `json:"RangePartitioning,omitnil,omitempty" name:"RangePartitioning"`

	// 分桶分区策略
	BucketPartitioning *BucketPartitioning `json:"BucketPartitioning,omitnil,omitempty" name:"BucketPartitioning"`

	// 截断分区策略
	TruncatePartitioning *TruncatePartitioning `json:"TruncatePartitioning,omitnil,omitempty" name:"TruncatePartitioning"`
}

type PermissionDetail struct {
	// 当前用户对该实体拥有的权限列表
	Permissions []*string `json:"Permissions,omitnil,omitempty" name:"Permissions"`

	// catalog在工作空间上的权限信息（可选）。取值：WORKSPACE_READONLY（只读）或WORKSPACE_READWRITE（读写）
	CatalogWorkspacePrivilege *string `json:"CatalogWorkspacePrivilege,omitnil,omitempty" name:"CatalogWorkspacePrivilege"`

	// deny权限总列表（用户deny ∪ 角色deny ∪ 继承deny，已去重，已从Permissions中排除）
	DenyPrivilegeList []*string `json:"DenyPrivilegeList,omitnil,omitempty" name:"DenyPrivilegeList"`
}

type PlatformStatusRsp struct {
	// 通用返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

type RangePartition struct {
	// 分区名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 下界
	Lower *Literal `json:"Lower,omitnil,omitempty" name:"Lower"`

	// 上界
	Upper *Literal `json:"Upper,omitnil,omitempty" name:"Upper"`

	// 属性
	Properties []*KVPair `json:"Properties,omitnil,omitempty" name:"Properties"`
}

type RangePartitioning struct {
	// 分区字段
	FieldName *string `json:"FieldName,omitnil,omitempty" name:"FieldName"`

	// 分区
	Assignments []*RangePartition `json:"Assignments,omitnil,omitempty" name:"Assignments"`
}

// Predefined struct for user
type RemoveConsoleUsersRequestParams struct {
	// <p>必填，待移除的用户 UIN 列表，单次最多10个</p>
	UserUins []*string `json:"UserUins,omitnil,omitempty" name:"UserUins"`
}

type RemoveConsoleUsersRequest struct {
	*tchttp.BaseRequest
	
	// <p>必填，待移除的用户 UIN 列表，单次最多10个</p>
	UserUins []*string `json:"UserUins,omitnil,omitempty" name:"UserUins"`
}

func (r *RemoveConsoleUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveConsoleUsersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserUins")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveConsoleUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveConsoleUsersResponseParams struct {
	// <p>批量移除控制台用户结果</p>
	Data *RemoveConsoleUsersRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemoveConsoleUsersResponse struct {
	*tchttp.BaseResponse
	Response *RemoveConsoleUsersResponseParams `json:"Response"`
}

func (r *RemoveConsoleUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveConsoleUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoveConsoleUsersRsp struct {
	// <p>请求已完成处理；即使部分失败也为 true，逐个结果以 SuccessUins/FailItems 为准</p>
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>删除成功的用户 UIN 列表</p>
	SuccessUins []*string `json:"SuccessUins,omitnil,omitempty" name:"SuccessUins"`

	// <p>失败项列表（Item 为用户 UIN，FailReason 为失败原因）</p>
	FailItems []*CommonFailItem `json:"FailItems,omitnil,omitempty" name:"FailItems"`
}

// Predefined struct for user
type RemoveWorkspaceGroupsRequestParams struct {
	// <p>用户组ID列表，可通过 ListWorkspaceGroups 接口获取</p>
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// <p>工作空间ID</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`
}

type RemoveWorkspaceGroupsRequest struct {
	*tchttp.BaseRequest
	
	// <p>用户组ID列表，可通过 ListWorkspaceGroups 接口获取</p>
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// <p>工作空间ID</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`
}

func (r *RemoveWorkspaceGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveWorkspaceGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupIds")
	delete(f, "WorkspaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveWorkspaceGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveWorkspaceGroupsResponseParams struct {
	// <p>移除工作空间用户组是否成功</p>
	Data *PlatformStatusRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemoveWorkspaceGroupsResponse struct {
	*tchttp.BaseResponse
	Response *RemoveWorkspaceGroupsResponseParams `json:"Response"`
}

func (r *RemoveWorkspaceGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveWorkspaceGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RerunWorkflowRunRequestParams struct {
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>工作流ID，可通过 ListWorkflows 获取。必填</p>
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// <p>工作流运行ID，可通过 ListWorkflowRuns 获取。必填</p>
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`

	// <p>运行类型。必填。取值：1 普通运行，2 高级运行</p>
	RunType *int64 `json:"RunType,omitnil,omitempty" name:"RunType"`

	// <p>运行类型为高级运行时填写的自定义运行参数</p>
	AdvancedParams []*TaskSchedulingParameterBrief `json:"AdvancedParams,omitnil,omitempty" name:"AdvancedParams"`

	// <p>本次需要重跑指定的任务ID集合，可通过 ListWorkflowTasks 获取，不传默认重跑该工作流下所有任务</p>
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// <p>计划调度时间列表配置</p>
	ScheduledTimeConfig *ScheduledTimeConfig `json:"ScheduledTimeConfig,omitnil,omitempty" name:"ScheduledTimeConfig"`
}

type RerunWorkflowRunRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>工作流ID，可通过 ListWorkflows 获取。必填</p>
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// <p>工作流运行ID，可通过 ListWorkflowRuns 获取。必填</p>
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`

	// <p>运行类型。必填。取值：1 普通运行，2 高级运行</p>
	RunType *int64 `json:"RunType,omitnil,omitempty" name:"RunType"`

	// <p>运行类型为高级运行时填写的自定义运行参数</p>
	AdvancedParams []*TaskSchedulingParameterBrief `json:"AdvancedParams,omitnil,omitempty" name:"AdvancedParams"`

	// <p>本次需要重跑指定的任务ID集合，可通过 ListWorkflowTasks 获取，不传默认重跑该工作流下所有任务</p>
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// <p>计划调度时间列表配置</p>
	ScheduledTimeConfig *ScheduledTimeConfig `json:"ScheduledTimeConfig,omitnil,omitempty" name:"ScheduledTimeConfig"`
}

func (r *RerunWorkflowRunRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RerunWorkflowRunRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "WorkflowId")
	delete(f, "WorkflowRunId")
	delete(f, "RunType")
	delete(f, "AdvancedParams")
	delete(f, "TaskIds")
	delete(f, "ScheduledTimeConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RerunWorkflowRunRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RerunWorkflowRunResponseParams struct {
	// <p>重跑工作流响应内容</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *AsyncActionRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RerunWorkflowRunResponse struct {
	*tchttp.BaseResponse
	Response *RerunWorkflowRunResponseParams `json:"Response"`
}

func (r *RerunWorkflowRunResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RerunWorkflowRunResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceGroupInfo struct {
	// <p>资源组ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// <p>资源组名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupName *string `json:"ResourceGroupName,omitnil,omitempty" name:"ResourceGroupName"`

	// <p>资源组状态</p><p>参数格式：0 // 未指定 1 // 待创建 2 // 创建中 3 // 运行中 4 // 已停止 5 // 停止中 6 // 启动中 7 // 更新中 8 // 删除中 9 // 已删除 10 // 用户主动启动 / 自动启动（有任务提交且自动启停开启） 11 // 可用: 仅存在于数据计算型 12 // 不可用: 仅存在于数据计算型 13 // 失败</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupStatus *string `json:"ResourceGroupStatus,omitnil,omitempty" name:"ResourceGroupStatus"`
}

type RoleBasicInfo struct {
	// <p>角色ID，可通过 ListWorkspaceRoles / ListConsoleRoles 接口获取；创建/更新角色时无需填写（由系统自动生成，创建成功后从响应 Data.RoleId 获取）</p>
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>角色名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>角色描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>显示名称</p>
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// <p>角色类型：workspace=工作空间系统角色，workspace_custom=工作空间自定义角色，console=控制台角色；查询时返回，创建时无需填写（系统固定创建自定义角色）</p>
	RoleType *string `json:"RoleType,omitnil,omitempty" name:"RoleType"`

	// <p>角色来源（查询侧字段，创建时无需填写）：0=未指定，1=用户直绑，2=用户组继承，3=两者都有</p>
	Source *int64 `json:"Source,omitnil,omitempty" name:"Source"`

	// <p>继承来源的用户组名称列表（查询侧字段，创建时无需填写），Source=1 时为空</p>
	GroupNames []*string `json:"GroupNames,omitnil,omitempty" name:"GroupNames"`
}

type RoleMetaData struct {
	// 创建者
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新者
	Updater *string `json:"Updater,omitnil,omitempty" name:"Updater"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type RolePermission struct {
	// <p>模块ID，须为当前租户已开通的功能模块（叶子节点）的模块ID（层级编码字符串，如 101=快速开始、109=工作流、116101103=工作空间管理_角色权限），非法值返回 InvalidParameterValue。模块清单获取方式：调用 ListWorkspaceModules 接口查询工作空间模块列表，取返回 Items[].ModuleId 字段（IsLeaf 为 true 的叶子节点）；也可通过控制台「工作空间设置-角色权限」页面查看</p>
	ModuleId *string `json:"ModuleId,omitnil,omitempty" name:"ModuleId"`

	// <p>模块访问权限，单值：R=只读，RW=读写，RWD=读写删除，N=无权限</p>
	Permissions *string `json:"Permissions,omitnil,omitempty" name:"Permissions"`
}

type RunActionBrief struct {
	// 工作流ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 工作流名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// 操作动作ID，用于追踪具体的执行动作
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunActionId *string `json:"RunActionId,omitnil,omitempty" name:"RunActionId"`

	// 失败错误信息，操作失败时返回具体的错误描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// 操作状态，true 表示成功，false 表示失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	OpStatus *bool `json:"OpStatus,omitnil,omitempty" name:"OpStatus"`

	// 工作流运行ID。重跑 / 终止场景返回被操作的运行ID；运行工作流场景为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`
}

// Predefined struct for user
type RunWorkflowRequestParams struct {
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>工作流ID，可通过 ListWorkflows 获取。必填</p>
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// <p>运行类型。必填。取值：1 普通运行，2 高级运行</p>
	RunType *int64 `json:"RunType,omitnil,omitempty" name:"RunType"`

	// <p>运行类型为高级运行时填写的自定义运行参数</p>
	AdvancedParams []*TaskSchedulingParameterBrief `json:"AdvancedParams,omitnil,omitempty" name:"AdvancedParams"`

	// <p>本次需要运行指定的任务ID集合，可通过 ListWorkflowTasks 获取，不传默认运行该工作流下所有任务</p>
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// <p>幂等令牌。非必填，相同令牌的重复请求只会触发一次运行</p>
	IdempotencyToken *string `json:"IdempotencyToken,omitnil,omitempty" name:"IdempotencyToken"`

	// <p>计划调度时间列表配置</p>
	ScheduledTimeConfig *ScheduledTimeConfig `json:"ScheduledTimeConfig,omitnil,omitempty" name:"ScheduledTimeConfig"`
}

type RunWorkflowRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>工作流ID，可通过 ListWorkflows 获取。必填</p>
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// <p>运行类型。必填。取值：1 普通运行，2 高级运行</p>
	RunType *int64 `json:"RunType,omitnil,omitempty" name:"RunType"`

	// <p>运行类型为高级运行时填写的自定义运行参数</p>
	AdvancedParams []*TaskSchedulingParameterBrief `json:"AdvancedParams,omitnil,omitempty" name:"AdvancedParams"`

	// <p>本次需要运行指定的任务ID集合，可通过 ListWorkflowTasks 获取，不传默认运行该工作流下所有任务</p>
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// <p>幂等令牌。非必填，相同令牌的重复请求只会触发一次运行</p>
	IdempotencyToken *string `json:"IdempotencyToken,omitnil,omitempty" name:"IdempotencyToken"`

	// <p>计划调度时间列表配置</p>
	ScheduledTimeConfig *ScheduledTimeConfig `json:"ScheduledTimeConfig,omitnil,omitempty" name:"ScheduledTimeConfig"`
}

func (r *RunWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunWorkflowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "WorkflowId")
	delete(f, "RunType")
	delete(f, "AdvancedParams")
	delete(f, "TaskIds")
	delete(f, "IdempotencyToken")
	delete(f, "ScheduledTimeConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunWorkflowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunWorkflowResponseParams struct {
	// <p>运行工作流响应内容</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *AsyncActionRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RunWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *RunWorkflowResponseParams `json:"Response"`
}

func (r *RunWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunWorkflowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScheduleBizEnumBrief struct {
	// 枚举标签键
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelKey *string `json:"LabelKey,omitnil,omitempty" name:"LabelKey"`

	// 枚举标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelValue *string `json:"LabelValue,omitnil,omitempty" name:"LabelValue"`

	// 枚举项统计数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type ScheduledTimeConfig struct {
	// <p>调度时区，IANA 时区 ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScheduledTimeZone *string `json:"ScheduledTimeZone,omitnil,omitempty" name:"ScheduledTimeZone"`

	// <p>调度生效开始时间</p><p>参数格式：毫秒时间戳（UTC）</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>调度生效结束时间</p><p>参数格式：毫秒时间戳（UTC）</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>周期类型</p><p>枚举值：</p><ul><li>DAY_CYCLE： 天</li><li>HOUR_CYCLE： 小时</li><li>MINUTE_CYCLE： 分钟</li><li>WEEK_CYCLE： 周</li></ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleType *string `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// <p>周期步长</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleNum *int64 `json:"CycleNum,omitnil,omitempty" name:"CycleNum"`
}

type Schema struct {
	// <p>schema名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>描述。注意：此字段可能返回null，表示取不到有效值</p>
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// <p>属性。注意：此字段可能返回null，表示取不到有效值</p>
	Properties []*KVPair `json:"Properties,omitnil,omitempty" name:"Properties"`

	// <p>审计信息。注意：此字段可能返回null，表示取不到有效值</p>
	Audit *Audit `json:"Audit,omitnil,omitempty" name:"Audit"`

	// <p>owner信息</p>
	MetaOwner *MetaOwner `json:"MetaOwner,omitnil,omitempty" name:"MetaOwner"`

	// <p>资产全局唯一ID，通过WedataAssetUIDUtils.generateUID生成</p>
	AssetGuid *string `json:"AssetGuid,omitnil,omitempty" name:"AssetGuid"`

	// <p>当前用户对该schema的权限信息。注意：此字段可能返回null，请求中未开启FetchPermissions时不返回</p>
	PermissionDetail *PermissionDetail `json:"PermissionDetail,omitnil,omitempty" name:"PermissionDetail"`

	// <p>标签信息列表。注意：此字段可能返回null，请求中未开启FetchTags时不返回</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*CommonTagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type SingleFieldPartitioning struct {
	// 分区字段
	FieldName *string `json:"FieldName,omitnil,omitempty" name:"FieldName"`
}

type SparseCheckoutConfig struct {
	// <p>是否启用稀疏检出</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// <p>是否使用 cone 模式（推荐 true，按目录匹配更高效）</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConeMode *bool `json:"ConeMode,omitnil,omitempty" name:"ConeMode"`

	// <p>稀疏检出路径列表（如 [&quot;src/module-a/&quot;, &quot;docs/&quot;]）</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Patterns []*string `json:"Patterns,omitnil,omitempty" name:"Patterns"`
}

type StandardUserInfo struct {
	// 用户UIN
	UserUin *string `json:"UserUin,omitnil,omitempty" name:"UserUin"`

	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 昵称
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// 0: 普通用户 1: entraId用户
	UserTag *string `json:"UserTag,omitnil,omitempty" name:"UserTag"`
}

// Predefined struct for user
type StartComputeRequestParams struct {
	// 计算资源 ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 工作空间 ID
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`
}

type StartComputeRequest struct {
	*tchttp.BaseRequest
	
	// 计算资源 ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 工作空间 ID
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`
}

func (r *StartComputeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartComputeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceId")
	delete(f, "WorkspaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartComputeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartComputeResponseParams struct {
	// 启动计算资源返回结果
	Data *StartComputeRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartComputeResponse struct {
	*tchttp.BaseResponse
	Response *StartComputeResponseParams `json:"Response"`
}

func (r *StartComputeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartComputeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartComputeRsp struct {
	// 操作是否成功
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type StopComputeRequestParams struct {
	// 计算资源 ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 工作空间 ID
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`
}

type StopComputeRequest struct {
	*tchttp.BaseRequest
	
	// 计算资源 ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 工作空间 ID
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`
}

func (r *StopComputeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopComputeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceId")
	delete(f, "WorkspaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopComputeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopComputeResponseParams struct {
	// 停止计算资源返回结果
	Data *StopComputeRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopComputeResponse struct {
	*tchttp.BaseResponse
	Response *StopComputeResponseParams `json:"Response"`
}

func (r *StopComputeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopComputeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopComputeRsp struct {
	// 操作是否成功
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

type TableInfo struct {
	// 表名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 字段信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Columns []*ColumnInfo `json:"Columns,omitnil,omitempty" name:"Columns"`

	// 属性值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Properties []*KVPair `json:"Properties,omitnil,omitempty" name:"Properties"`

	// 分区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Partitioning []*Partitioning `json:"Partitioning,omitnil,omitempty" name:"Partitioning"`

	// 索引
	// 注意：此字段可能返回 null，表示取不到有效值。
	Indexes []*Index `json:"Indexes,omitnil,omitempty" name:"Indexes"`

	// 编辑者
	// 注意：此字段可能返回 null，表示取不到有效值。
	Audit *Audit `json:"Audit,omitnil,omitempty" name:"Audit"`

	// 表版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableFormat *string `json:"TableFormat,omitnil,omitempty" name:"TableFormat"`

	// owner信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetaOwner *MetaOwner `json:"MetaOwner,omitnil,omitempty" name:"MetaOwner"`

	// 资产全局唯一ID
	AssetGuid *string `json:"AssetGuid,omitnil,omitempty" name:"AssetGuid"`

	// 总文件数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilesCount *string `json:"FilesCount,omitnil,omitempty" name:"FilesCount"`

	// 总数据量大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilesSize *string `json:"FilesSize,omitnil,omitempty" name:"FilesSize"`

	// 表类型
	// 
	// - Managed 内部
	// - External 外部
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableType *string `json:"TableType,omitnil,omitempty" name:"TableType"`

	// 特征表类型：Online-在线表 Offline-离线表 空字符串-非特征表
	FeatureTableType *string `json:"FeatureTableType,omitnil,omitempty" name:"FeatureTableType"`

	// 特征表详细信息，仅当FeatureTableType不为空时有值
	FeatureTableDetail *FeatureTableDetail `json:"FeatureTableDetail,omitnil,omitempty" name:"FeatureTableDetail"`

	// 当前用户对该实体的权限信息
	PermissionDetail *PermissionDetail `json:"PermissionDetail,omitnil,omitempty" name:"PermissionDetail"`

	// 表级标签列表，仅当请求中FetchOption.FetchTags=true时返回，默认不返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*CommonTagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type TagValueInfo struct {
	// 标签值ID
	TagValueId *string `json:"TagValueId,omitnil,omitempty" name:"TagValueId"`

	// 标签值内容，查询时填充，写入时不传
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`

	// 标签来源，参考 web_enum_standard.proto -> TagSource：0=未指定 1=用户直绑 2=用户组继承 3=两者都有，查询时填充，写入时不传
	Source *int64 `json:"Source,omitnil,omitempty" name:"Source"`

	// 继承来源的用户组名称列表，Source=1 时为空，查询时填充，写入时不传
	GroupNames []*string `json:"GroupNames,omitnil,omitempty" name:"GroupNames"`
}

type TaskRetryStrategy struct {
	// 最多重试次数，默认3
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxRetryTimes *int64 `json:"MaxRetryTimes,omitnil,omitempty" name:"MaxRetryTimes"`

	// 重试之间等待时间，默认5
	// 注意：此字段可能返回 null，表示取不到有效值。
	RetryBetweenWaitTime *int64 `json:"RetryBetweenWaitTime,omitnil,omitempty" name:"RetryBetweenWaitTime"`

	// 重试之间等待时间单位
	// 毫秒：MILLISECOND秒：SECOND分钟（默认）：MINUTE小时：HOUR
	// 注意：此字段可能返回 null，表示取不到有效值。
	RetryBetweenWaitTimeUnit *string `json:"RetryBetweenWaitTimeUnit,omitnil,omitempty" name:"RetryBetweenWaitTimeUnit"`

	// 任务运行失败时重试开关，默认为true
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskRunFailureRetrySwitch *bool `json:"TaskRunFailureRetrySwitch,omitnil,omitempty" name:"TaskRunFailureRetrySwitch"`

	// 任务运行超时时重试开关，默认为false
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskRunTimeoutRetrySwitch *bool `json:"TaskRunTimeoutRetrySwitch,omitnil,omitempty" name:"TaskRunTimeoutRetrySwitch"`
}

type TaskRunConditionRule struct {
	// <p>上游任务ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpstreamTaskId *string `json:"UpstreamTaskId,omitnil,omitempty" name:"UpstreamTaskId"`

	// <p>上游任务名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpstreamTaskName *string `json:"UpstreamTaskName,omitnil,omitempty" name:"UpstreamTaskName"`

	// <p>任务可运行条件<br>支持的状态值： - SUCCESS: 成功 - FAILED: 失败 - UPSTREAM_FAILED: 上游失败 - EXCLUDED: 排除运行</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllowedStates []*string `json:"AllowedStates,omitnil,omitempty" name:"AllowedStates"`
}

type TaskSchedulingParameterBrief struct {
	// 参数名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamKey *string `json:"ParamKey,omitnil,omitempty" name:"ParamKey"`

	// 参数值
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamValue *string `json:"ParamValue,omitnil,omitempty" name:"ParamValue"`
}

type TaskType struct {
	// <p>任务类型：SQL：用于执行SQL查询和数据处理操作；DATA_INTEGRATION：用于离线数据接入操作；NOTEBOOK：用于运行Notebook脚本；RUN_WORKFLOW：用于执行嵌套工作流；PYTHON：用于运行Python脚本；RAY_JOB：用于运行Ray作业；DATA_QUALITY：用于数据质量监控；IF_ELSE：用于条件分支判断；FOR_EACH：用于循环遍历执行；</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTypeName *string `json:"TaskTypeName,omitnil,omitempty" name:"TaskTypeName"`

	// <p>Notebook 类型扩展信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: Notebook is deprecated.
	Notebook *TaskTypeNotebookExt `json:"Notebook,omitnil,omitempty" name:"Notebook"`

	// <p>任务扩展属性列表，具体填写参考TaskType.TaskPropertyList描述</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTypePropertyList []*TaskTypeProperty `json:"TaskTypePropertyList,omitnil,omitempty" name:"TaskTypePropertyList"`

	// <p>运行时属性列表，具体填写参考TaskType.RuntimePropertyList描述</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuntimePropertyList []*TaskTypeProperty `json:"RuntimePropertyList,omitnil,omitempty" name:"RuntimePropertyList"`
}

type TaskTypeNotebookExt struct {
	// <p>脚本来源。取值：SCRIPT_SOURCE_LOCAL（本地）/ SCRIPT_SOURCE_GIT（Git 仓库）/<br>SCRIPT_SOURCE_CFS（CFS 文件系统）/ SCRIPT_SOURCE_COS（COS 对象存储）/<br>SCRIPT_SOURCE_WORKSPACE（工作空间）</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: Source is deprecated.
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// <p>前端显示使用，对执行平台无意义</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: DisplayPath is deprecated.
	DisplayPath *string `json:"DisplayPath,omitnil,omitempty" name:"DisplayPath"`

	// <p>Notebook 相对路径</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: NotebookPath is deprecated.
	NotebookPath *string `json:"NotebookPath,omitnil,omitempty" name:"NotebookPath"`

	// <p>Notebook 绝对路径</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: NotebookAbsolutePath is deprecated.
	NotebookAbsolutePath *string `json:"NotebookAbsolutePath,omitnil,omitempty" name:"NotebookAbsolutePath"`
}

type TaskTypeProperty struct {
	// <p>属性键名。用于TaskType内的入参时，若不传或为空，该属性项被服务端静默忽略（不报错），等同于该属性未配置。合法 PropertyKey 清单请参考TaskType中描述。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	PropertyKey *string `json:"PropertyKey,omitnil,omitempty" name:"PropertyKey"`

	// <p>属性值</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	PropertyValue *string `json:"PropertyValue,omitnil,omitempty" name:"PropertyValue"`
}

type TruncatePartitioning struct {
	// 字段名
	FiledName *string `json:"FiledName,omitnil,omitempty" name:"FiledName"`

	// 截取长度
	Width *string `json:"Width,omitnil,omitempty" name:"Width"`
}

// Predefined struct for user
type UnbindWorkflowBundleRequestParams struct {
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>工作流ID，可通过 ListWorkflows 获取。必填</p>
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

type UnbindWorkflowBundleRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>工作流ID，可通过 ListWorkflows 获取。必填</p>
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

func (r *UnbindWorkflowBundleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindWorkflowBundleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "WorkflowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindWorkflowBundleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindWorkflowBundleResponseParams struct {
	// <p>解绑工作流Bundle信息响应内容</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *UnbindWorkflowBundleRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UnbindWorkflowBundleResponse struct {
	*tchttp.BaseResponse
	Response *UnbindWorkflowBundleResponseParams `json:"Response"`
}

func (r *UnbindWorkflowBundleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindWorkflowBundleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnbindWorkflowBundleRsp struct {
	// 操作状态，true 表示成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type UpdateConsoleGroupRequestParams struct {
	// <p>用户组 ID，可通过 ListConsoleGroups 接口获取</p>
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// <p>修改标识：USER_GROUP_OPER_TYPE_ADD_USER(1)=添加成员、USER_GROUP_OPER_TYPE_DELETE_USER(2)=删除成员、USER_GROUP_OPER_TYPE_BASIC_INFO(3)=基础信息（别名和描述）</p>
	OperType *int64 `json:"OperType,omitnil,omitempty" name:"OperType"`

	// <p>用户组名称</p>
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// <p>用户组别名</p>
	GroupNickname *string `json:"GroupNickname,omitnil,omitempty" name:"GroupNickname"`

	// <p>用户组描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>成员 UIN 列表（OperType 为添加/删除成员时使用），单次最多100个</p>
	UserUins []*string `json:"UserUins,omitnil,omitempty" name:"UserUins"`
}

type UpdateConsoleGroupRequest struct {
	*tchttp.BaseRequest
	
	// <p>用户组 ID，可通过 ListConsoleGroups 接口获取</p>
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// <p>修改标识：USER_GROUP_OPER_TYPE_ADD_USER(1)=添加成员、USER_GROUP_OPER_TYPE_DELETE_USER(2)=删除成员、USER_GROUP_OPER_TYPE_BASIC_INFO(3)=基础信息（别名和描述）</p>
	OperType *int64 `json:"OperType,omitnil,omitempty" name:"OperType"`

	// <p>用户组名称</p>
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// <p>用户组别名</p>
	GroupNickname *string `json:"GroupNickname,omitnil,omitempty" name:"GroupNickname"`

	// <p>用户组描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>成员 UIN 列表（OperType 为添加/删除成员时使用），单次最多100个</p>
	UserUins []*string `json:"UserUins,omitnil,omitempty" name:"UserUins"`
}

func (r *UpdateConsoleGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateConsoleGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "OperType")
	delete(f, "GroupName")
	delete(f, "GroupNickname")
	delete(f, "Description")
	delete(f, "UserUins")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateConsoleGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateConsoleGroupResponseParams struct {
	// <p>返回结果</p>
	Data *UpdateConsoleGroupRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateConsoleGroupResponse struct {
	*tchttp.BaseResponse
	Response *UpdateConsoleGroupResponseParams `json:"Response"`
}

func (r *UpdateConsoleGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateConsoleGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateConsoleGroupRsp struct {
	// 操作是否成功
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type UpdateConsoleUsersRequestParams struct {
	// <p>用户 UIN 列表，单次最多100个</p>
	UserUins []*string `json:"UserUins,omitnil,omitempty" name:"UserUins"`

	// <p>角色 ID 列表</p><p>枚举值：</p><ul><li>2001： 控制台管理员</li><li>2002： 控制台成员</li></ul><p>可通过 ListConsoleRoles 接口获取</p>
	RoleIds []*string `json:"RoleIds,omitnil,omitempty" name:"RoleIds"`
}

type UpdateConsoleUsersRequest struct {
	*tchttp.BaseRequest
	
	// <p>用户 UIN 列表，单次最多100个</p>
	UserUins []*string `json:"UserUins,omitnil,omitempty" name:"UserUins"`

	// <p>角色 ID 列表</p><p>枚举值：</p><ul><li>2001： 控制台管理员</li><li>2002： 控制台成员</li></ul><p>可通过 ListConsoleRoles 接口获取</p>
	RoleIds []*string `json:"RoleIds,omitnil,omitempty" name:"RoleIds"`
}

func (r *UpdateConsoleUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateConsoleUsersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserUins")
	delete(f, "RoleIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateConsoleUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateConsoleUsersResponseParams struct {
	// <p>返回结果</p>
	Data *UpdateConsoleUsersRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateConsoleUsersResponse struct {
	*tchttp.BaseResponse
	Response *UpdateConsoleUsersResponseParams `json:"Response"`
}

func (r *UpdateConsoleUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateConsoleUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateConsoleUsersRsp struct {
	// 操作是否成功
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type UpdateFileRequestParams struct {
	// <p>工作空间 ID。来源：ListWorkspaces 接口返回的 WorkspaceId</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>文件 ID。来源：CreateFile / ListFiles / GetFile 接口返回的 FileId</p>
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <p>文件运行配置。不传则不更新配置</p>
	FileConfig *FileConfig `json:"FileConfig,omitnil,omitempty" name:"FileConfig"`

	// <p>文件类型。取值：FILE（普通文件/脚本）、NOTEBOOK_FILE（Notebook）、SQL_FILE（SQL文件）。对应 common/domain/entity.proto EntityType</p>
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// <p>绑定的 BundleId。来源：ListBundles 接口返回的 BundleId</p>
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// <p>绑定的 BundleInfo，JSON 字符串</p>
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`

	// <p>文件内容。不传则不更新内容</p>
	Storage *FileStorage `json:"Storage,omitnil,omitempty" name:"Storage"`

	// <p>目标文件名，非空且与当前文件名不同时执行 rename 动作。长度不超过 SCRIPT_NAME_MAX_LENGTH，禁止以 . 或 .. 开头/结尾，禁止空格、双点、控制字符及 Linux 保留名（参考 docs/linux_filename_rules.md）。与 ExtensionType 一起校验后缀合法性</p>
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// <p>目标父目录路径，非空时执行 move 动作。根目录传 /；与 FileName 可同时出现，语义为「移动+重命名」。与 CreateFile 的 ParentFolderPath 保持一致</p>
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// <p>目标父目录的 FileType。取值：FOLDER、GIT_FOLDER。仅当 ParentFolderPath 非空时使用；缺省时按解析出的父目录实际类型处理</p>
	TargetFileType *string `json:"TargetFileType,omitnil,omitempty" name:"TargetFileType"`

	// <p>动作类型（选填，默认取值1，未来版本会强制校验）。取值：1 = UPDATE_CONTENT（仅更新 FileConfig / Storage / Bundle*，禁止传 FileName / ParentFolderPath / TargetFileType）；2 = RENAME（仅重命名，必须传 FileName，禁止传 ParentFolderPath / FileConfig / Storage / Bundle*）；3 = MOVE（仅移动，必须传 ParentFolderPath，禁止传 FileName / FileConfig / Storage / Bundle*）。参数互斥校验失败会返回 ParamIllegal 错误</p>
	UpdateAction *int64 `json:"UpdateAction,omitnil,omitempty" name:"UpdateAction"`
}

type UpdateFileRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间 ID。来源：ListWorkspaces 接口返回的 WorkspaceId</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>文件 ID。来源：CreateFile / ListFiles / GetFile 接口返回的 FileId</p>
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <p>文件运行配置。不传则不更新配置</p>
	FileConfig *FileConfig `json:"FileConfig,omitnil,omitempty" name:"FileConfig"`

	// <p>文件类型。取值：FILE（普通文件/脚本）、NOTEBOOK_FILE（Notebook）、SQL_FILE（SQL文件）。对应 common/domain/entity.proto EntityType</p>
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// <p>绑定的 BundleId。来源：ListBundles 接口返回的 BundleId</p>
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// <p>绑定的 BundleInfo，JSON 字符串</p>
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`

	// <p>文件内容。不传则不更新内容</p>
	Storage *FileStorage `json:"Storage,omitnil,omitempty" name:"Storage"`

	// <p>目标文件名，非空且与当前文件名不同时执行 rename 动作。长度不超过 SCRIPT_NAME_MAX_LENGTH，禁止以 . 或 .. 开头/结尾，禁止空格、双点、控制字符及 Linux 保留名（参考 docs/linux_filename_rules.md）。与 ExtensionType 一起校验后缀合法性</p>
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// <p>目标父目录路径，非空时执行 move 动作。根目录传 /；与 FileName 可同时出现，语义为「移动+重命名」。与 CreateFile 的 ParentFolderPath 保持一致</p>
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// <p>目标父目录的 FileType。取值：FOLDER、GIT_FOLDER。仅当 ParentFolderPath 非空时使用；缺省时按解析出的父目录实际类型处理</p>
	TargetFileType *string `json:"TargetFileType,omitnil,omitempty" name:"TargetFileType"`

	// <p>动作类型（选填，默认取值1，未来版本会强制校验）。取值：1 = UPDATE_CONTENT（仅更新 FileConfig / Storage / Bundle*，禁止传 FileName / ParentFolderPath / TargetFileType）；2 = RENAME（仅重命名，必须传 FileName，禁止传 ParentFolderPath / FileConfig / Storage / Bundle*）；3 = MOVE（仅移动，必须传 ParentFolderPath，禁止传 FileName / FileConfig / Storage / Bundle*）。参数互斥校验失败会返回 ParamIllegal 错误</p>
	UpdateAction *int64 `json:"UpdateAction,omitnil,omitempty" name:"UpdateAction"`
}

func (r *UpdateFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "FileId")
	delete(f, "FileConfig")
	delete(f, "FileType")
	delete(f, "BundleId")
	delete(f, "BundleInfo")
	delete(f, "Storage")
	delete(f, "FileName")
	delete(f, "ParentFolderPath")
	delete(f, "TargetFileType")
	delete(f, "UpdateAction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateFileResponseParams struct {
	// <p>返回结果</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *FileInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateFileResponse struct {
	*tchttp.BaseResponse
	Response *UpdateFileResponseParams `json:"Response"`
}

func (r *UpdateFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateFolderRequestParams struct {
	// <p>工作空间ID</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>待更新文件夹</p>
	Folder *FolderLocator `json:"Folder,omitnil,omitempty" name:"Folder"`

	// <p>操作类型</p><p>枚举值：</p><ul><li>1： 重命名</li><li>2： 移动</li></ul>
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`

	// <p>重命名后的文件名，OperationType = 1时生效</p>
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`

	// <p>移动的目的文件夹，OperationType = 2时生效</p>
	TargetParent *FolderLocator `json:"TargetParent,omitnil,omitempty" name:"TargetParent"`
}

type UpdateFolderRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间ID</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>待更新文件夹</p>
	Folder *FolderLocator `json:"Folder,omitnil,omitempty" name:"Folder"`

	// <p>操作类型</p><p>枚举值：</p><ul><li>1： 重命名</li><li>2： 移动</li></ul>
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`

	// <p>重命名后的文件名，OperationType = 1时生效</p>
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`

	// <p>移动的目的文件夹，OperationType = 2时生效</p>
	TargetParent *FolderLocator `json:"TargetParent,omitnil,omitempty" name:"TargetParent"`
}

func (r *UpdateFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "Folder")
	delete(f, "OperationType")
	delete(f, "FolderName")
	delete(f, "TargetParent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateFolderResponseParams struct {
	// <p>更新文件夹结果</p>
	Data *UpdateFolderRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateFolderResponse struct {
	*tchttp.BaseResponse
	Response *UpdateFolderResponseParams `json:"Response"`
}

func (r *UpdateFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateFolderRsp struct {
	// <p>更新文件夹结果，true为成功</p>
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type UpdateWorkflowRequestParams struct {
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>待更新的工作流ID，可通过 ListWorkflows 获取。必填</p>
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// <p>需要清空的字段名列表，用于将指定字段重置为空</p>
	FieldToRemoveList []*string `json:"FieldToRemoveList,omitnil,omitempty" name:"FieldToRemoveList"`

	// <p>更新后的工作流配置，仅传入需要变更的部分即可</p>
	NewSetting *Workflow `json:"NewSetting,omitnil,omitempty" name:"NewSetting"`
}

type UpdateWorkflowRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>待更新的工作流ID，可通过 ListWorkflows 获取。必填</p>
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// <p>需要清空的字段名列表，用于将指定字段重置为空</p>
	FieldToRemoveList []*string `json:"FieldToRemoveList,omitnil,omitempty" name:"FieldToRemoveList"`

	// <p>更新后的工作流配置，仅传入需要变更的部分即可</p>
	NewSetting *Workflow `json:"NewSetting,omitnil,omitempty" name:"NewSetting"`
}

func (r *UpdateWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateWorkflowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "WorkflowId")
	delete(f, "FieldToRemoveList")
	delete(f, "NewSetting")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateWorkflowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateWorkflowResponseParams struct {
	// <p>更新工作流响应内容</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *UpdateWorkflowRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *UpdateWorkflowResponseParams `json:"Response"`
}

func (r *UpdateWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateWorkflowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateWorkflowRsp struct {
	// 更新状态，true 表示成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type UpdateWorkspaceGroupsRolesRequestParams struct {
	// 用户组ID，可通过 ListWorkspaceGroups 接口获取
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// 工作空间ID
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 角色ID，可通过 ListWorkspaceRoles 接口获取
	RoleIds []*string `json:"RoleIds,omitnil,omitempty" name:"RoleIds"`
}

type UpdateWorkspaceGroupsRolesRequest struct {
	*tchttp.BaseRequest
	
	// 用户组ID，可通过 ListWorkspaceGroups 接口获取
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// 工作空间ID
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 角色ID，可通过 ListWorkspaceRoles 接口获取
	RoleIds []*string `json:"RoleIds,omitnil,omitempty" name:"RoleIds"`
}

func (r *UpdateWorkspaceGroupsRolesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateWorkspaceGroupsRolesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupIds")
	delete(f, "WorkspaceId")
	delete(f, "RoleIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateWorkspaceGroupsRolesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateWorkspaceGroupsRolesResponseParams struct {
	// 修改工作空间用户组角色是否成功
	Data *PlatformStatusRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateWorkspaceGroupsRolesResponse struct {
	*tchttp.BaseResponse
	Response *UpdateWorkspaceGroupsRolesResponseParams `json:"Response"`
}

func (r *UpdateWorkspaceGroupsRolesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateWorkspaceGroupsRolesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateWorkspaceRequestParams struct {
	// <p>工作空间ID</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>工作空间名称，max_len=128</p>
	WorkspaceName *string `json:"WorkspaceName,omitnil,omitempty" name:"WorkspaceName"`

	// <p>工作空间描述，max_len=300</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type UpdateWorkspaceRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间ID</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>工作空间名称，max_len=128</p>
	WorkspaceName *string `json:"WorkspaceName,omitnil,omitempty" name:"WorkspaceName"`

	// <p>工作空间描述，max_len=300</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *UpdateWorkspaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateWorkspaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "WorkspaceName")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateWorkspaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateWorkspaceResponseParams struct {
	// <p>操作结果</p>
	Data *UpdateWorkspaceRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateWorkspaceResponse struct {
	*tchttp.BaseResponse
	Response *UpdateWorkspaceResponseParams `json:"Response"`
}

func (r *UpdateWorkspaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateWorkspaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateWorkspaceRoleRequestParams struct {
	// 工作空间ID
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 角色信息
	BasicInfo *RoleBasicInfo `json:"BasicInfo,omitnil,omitempty" name:"BasicInfo"`

	// 功能点权限
	Permissions []*RolePermission `json:"Permissions,omitnil,omitempty" name:"Permissions"`
}

type UpdateWorkspaceRoleRequest struct {
	*tchttp.BaseRequest
	
	// 工作空间ID
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 角色信息
	BasicInfo *RoleBasicInfo `json:"BasicInfo,omitnil,omitempty" name:"BasicInfo"`

	// 功能点权限
	Permissions []*RolePermission `json:"Permissions,omitnil,omitempty" name:"Permissions"`
}

func (r *UpdateWorkspaceRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateWorkspaceRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "BasicInfo")
	delete(f, "Permissions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateWorkspaceRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateWorkspaceRoleResponseParams struct {
	// 更新工作空间角色结果
	Data *UpdateWorkspaceRoleResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateWorkspaceRoleResponse struct {
	*tchttp.BaseResponse
	Response *UpdateWorkspaceRoleResponseParams `json:"Response"`
}

func (r *UpdateWorkspaceRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateWorkspaceRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateWorkspaceRoleResult struct {
	// 更新工作空间角色是否成功
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

type UpdateWorkspaceRsp struct {
	// 操作是否成功
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type UpdateWorkspaceUsersRolesRequestParams struct {
	// <p>用户ID</p>
	UserUins []*string `json:"UserUins,omitnil,omitempty" name:"UserUins"`

	// <p>工作空间ID</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>角色ID，可通过 ListWorkspaceRoles 接口获取</p>
	RoleIds []*string `json:"RoleIds,omitnil,omitempty" name:"RoleIds"`
}

type UpdateWorkspaceUsersRolesRequest struct {
	*tchttp.BaseRequest
	
	// <p>用户ID</p>
	UserUins []*string `json:"UserUins,omitnil,omitempty" name:"UserUins"`

	// <p>工作空间ID</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>角色ID，可通过 ListWorkspaceRoles 接口获取</p>
	RoleIds []*string `json:"RoleIds,omitnil,omitempty" name:"RoleIds"`
}

func (r *UpdateWorkspaceUsersRolesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateWorkspaceUsersRolesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserUins")
	delete(f, "WorkspaceId")
	delete(f, "RoleIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateWorkspaceUsersRolesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateWorkspaceUsersRolesResponseParams struct {
	// <p>结果</p>
	Data *PlatformStatusRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateWorkspaceUsersRolesResponse struct {
	*tchttp.BaseResponse
	Response *UpdateWorkspaceUsersRolesResponseParams `json:"Response"`
}

func (r *UpdateWorkspaceUsersRolesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateWorkspaceUsersRolesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserInfo struct {
	// <p>uin</p>
	UserUin *string `json:"UserUin,omitnil,omitempty" name:"UserUin"`

	// <p>子用户名称</p>
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// <p>子用户昵称</p>
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// <p>0: 普通用户 1: entraId用户</p>
	UserTag *string `json:"UserTag,omitnil,omitempty" name:"UserTag"`
}

type UserTagInfo struct {
	// <p>标签维度ID</p>
	TagId *string `json:"TagId,omitnil,omitempty" name:"TagId"`

	// <p>标签名称，查询时填充，写入时不传</p>
	TagName *string `json:"TagName,omitnil,omitempty" name:"TagName"`

	// <p>该维度下绑定的所有标签值</p>
	Values []*TagValueInfo `json:"Values,omitnil,omitempty" name:"Values"`
}

type Workflow struct {
	// <p>工作空间ID，可通过 ListWorkspaces 获取</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>工作流基本信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BaseInfo *WorkflowBaseInfo `json:"BaseInfo,omitnil,omitempty" name:"BaseInfo"`

	// <p>工作流调度配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Trigger []*WorkflowTriggerConfiguration `json:"Trigger,omitnil,omitempty" name:"Trigger"`

	// <p>工作流参数列表 参数名必填且只能包含数字、大小写字母、空格、.$@#!%^&amp;*()-_+=&gt; 最长128个字符</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamList []*ParamInfo `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// <p>标签 标签名必填且只能包含数字、大小写字母、空格、.$@#!%^&amp;*()-_+=&gt;'，最长128个字符</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelList []*LabelBrief `json:"LabelList,omitnil,omitempty" name:"LabelList"`

	// <p>工作流告警配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alarm *AlarmBrief `json:"Alarm,omitnil,omitempty" name:"Alarm"`

	// <p>监控指标配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	MonitorMetric *MonitorMetricBrief `json:"MonitorMetric,omitnil,omitempty" name:"MonitorMetric"`

	// <p>工作流高级设置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdvanceConfig *WorkflowAdvanceConfig `json:"AdvanceConfig,omitnil,omitempty" name:"AdvanceConfig"`

	// <p>工作流任务列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskList []*WorkflowTask `json:"TaskList,omitnil,omitempty" name:"TaskList"`

	// <p>BundleId，可通过 Bundle 相关接口获取</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// <p>Bundle信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`

	// <p>GIT配置ID，对应GetWorkspaceConfig接口中的ConfigKey</p>    -
	// 注意：此字段可能返回 null，表示取不到有效值。
	GitConfigId *string `json:"GitConfigId,omitnil,omitempty" name:"GitConfigId"`

	// <p>Git分支信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	GitBranch *string `json:"GitBranch,omitnil,omitempty" name:"GitBranch"`
}

type WorkflowAdvanceConfig struct {
	// 排队模式，ON（默认）, OFF
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueuingMode *string `json:"QueuingMode,omitnil,omitempty" name:"QueuingMode"`

	// 	
	// 默认值为1
	// 
	// QueuingMode为ON时，MaxConcurrentNum 设置才生效；只能输入大于0的整数，输入非法值自动转换为1
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxConcurrentNum *int64 `json:"MaxConcurrentNum,omitnil,omitempty" name:"MaxConcurrentNum"`
}

type WorkflowBaseInfo struct {
	// 工作流名称，长度不超过 1024
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// 工作流ID，创建时无需传入，由服务端生成
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 工作流运行人UIN
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunUserUin *string `json:"RunUserUin,omitnil,omitempty" name:"RunUserUin"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 工作流负责人用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUserName *string `json:"OwnerUserName,omitnil,omitempty" name:"OwnerUserName"`

	// <p>创建人UIN。系统生成字段，仅出参返回，入参无需传入（传入不生效，服务端忽略且不报错）</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: CreateUserUin is deprecated.
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`
}

type WorkflowBaseInfoDetail struct {
	// 工作流名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// 工作流ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 创建人UIN
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// 工作流运行人UIN
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunUserUin *string `json:"RunUserUin,omitnil,omitempty" name:"RunUserUin"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 工作流负责人用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUserName *string `json:"OwnerUserName,omitnil,omitempty" name:"OwnerUserName"`

	// 工作流负责人UIN
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUserUin *string `json:"OwnerUserUin,omitnil,omitempty" name:"OwnerUserUin"`

	// 工作流负责人展示名
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerDisplayName *string `json:"OwnerDisplayName,omitnil,omitempty" name:"OwnerDisplayName"`

	// 创建时间，单位：毫秒时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间，单位：毫秒时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type WorkflowBrief struct {
	// <p>工作流名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// <p>工作流ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// <p>描述</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>创建人UIN</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// <p>工作流负责人用户名</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUserName *string `json:"OwnerUserName,omitnil,omitempty" name:"OwnerUserName"`

	// <p>工作流负责人UIN</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUserUin *string `json:"OwnerUserUin,omitnil,omitempty" name:"OwnerUserUin"`

	// <p>工作流负责人展示名</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerDisplayName *string `json:"OwnerDisplayName,omitnil,omitempty" name:"OwnerDisplayName"`

	// <p>创建时间，单位：毫秒时间戳</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>更新时间，单位：毫秒时间戳</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>标签列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelList []*LabelBrief `json:"LabelList,omitnil,omitempty" name:"LabelList"`

	// <p>工作流调度配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Trigger []*WorkflowTriggerConfiguration `json:"Trigger,omitnil,omitempty" name:"Trigger"`

	// <p>工作流运行人UIN</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunUserUin *string `json:"RunUserUin,omitnil,omitempty" name:"RunUserUin"`

	// <p>工作流运行人用户名</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunUserName *string `json:"RunUserName,omitnil,omitempty" name:"RunUserName"`

	// <p>工作流任务节点列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskList []*WorkflowTaskNodeBrief `json:"TaskList,omitnil,omitempty" name:"TaskList"`

	// <p>工作流运行情况列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowRunList []*WorkflowRunBrief `json:"WorkflowRunList,omitnil,omitempty" name:"WorkflowRunList"`

	// <p>资源组信息列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupInfoList []*ResourceGroupInfo `json:"ResourceGroupInfoList,omitnil,omitempty" name:"ResourceGroupInfoList"`

	// <p>授权权限类型<br>PERMISSION_TYPE_UNSPECIFIED：未指定权限<br>MANAGE : 管理权限：包含所有操作权限<br>RUN : 运行权限：可执行实体<br>VIEW : 查看权限：可查看实体内容</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Permission *string `json:"Permission,omitnil,omitempty" name:"Permission"`

	// <p>工作流绑定的 Bundle 唯一标识，未绑定时为空</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// <p>Bundle信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`

	// <p>Git配置ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	GitConfigId *string `json:"GitConfigId,omitnil,omitempty" name:"GitConfigId"`

	// <p>Git分支信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	GitBranch *string `json:"GitBranch,omitnil,omitempty" name:"GitBranch"`
}

type WorkflowRun struct {
	// <p>主账号ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>工作流名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// <p>工作流ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// <p>工作流运行ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`

	// <p>工作空间ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>触发方式，Scheduler、ManualTrigger、Event (参考SchedulerTriggerType)</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TriggerType *string `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// <p>运行开始时间，单位：毫秒时间戳</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunStartTime *string `json:"RunStartTime,omitnil,omitempty" name:"RunStartTime"`

	// <p>pending 状态开始时间，单位：毫秒时间戳</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	PendingStartTime *string `json:"PendingStartTime,omitnil,omitempty" name:"PendingStartTime"`

	// <p>queue 状态开始时间，单位：毫秒时间戳</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueueStartTime *string `json:"QueueStartTime,omitnil,omitempty" name:"QueueStartTime"`

	// <p>运行结束时间，单位：毫秒时间戳</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunEndTime *string `json:"RunEndTime,omitnil,omitempty" name:"RunEndTime"`

	// <p>终态时间，运行进入终态时都有值，单位：毫秒时间戳</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>运行时长，单位：秒</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunCostTime *string `json:"RunCostTime,omitnil,omitempty" name:"RunCostTime"`

	// <p>并发排队花费时间，单位：秒</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueueCostTime *string `json:"QueueCostTime,omitnil,omitempty" name:"QueueCostTime"`

	// <p>等待资源花费时间，单位：秒</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	PendingCostTime *string `json:"PendingCostTime,omitnil,omitempty" name:"PendingCostTime"`

	// <p>运行状态。CREATE(&quot;初始化&quot;),     QUEUED(&quot;等待中&quot;),     PENDING(&quot;准备中&quot;),     RUNNING(&quot;运行中&quot;),     SKIPPED(&quot;跳过运行&quot;),     SUCCESS(&quot;成功&quot;),     FAILED(&quot;失败&quot;),     TERMINATING(&quot;终止中&quot;),     TERMINATED(&quot;终止&quot;),     CANCELLED(&quot;被手动终止&quot;)等</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunState *string `json:"RunState,omitnil,omitempty" name:"RunState"`

	// <p>计算资源（任务的资源组ID集合）</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupIds []*string `json:"ResourceGroupIds,omitnil,omitempty" name:"ResourceGroupIds"`

	// <p>运行用户UIN</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunUserUin *string `json:"RunUserUin,omitnil,omitempty" name:"RunUserUin"`

	// <p>运行用户名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunUserName *string `json:"RunUserName,omitnil,omitempty" name:"RunUserName"`

	// <p>错误码</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorCodeString *string `json:"ErrorCodeString,omitnil,omitempty" name:"ErrorCodeString"`

	// <p>运行参数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowParams *string `json:"WorkflowParams,omitnil,omitempty" name:"WorkflowParams"`

	// <p>工作流版本ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowVersionId *string `json:"WorkflowVersionId,omitnil,omitempty" name:"WorkflowVersionId"`

	// <p>当前工作流是否支持重跑</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SupportRerun *bool `json:"SupportRerun,omitnil,omitempty" name:"SupportRerun"`

	// <p>工作流运行创建时间，单位：毫秒时间戳</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>重跑次数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RerunTimes *int64 `json:"RerunTimes,omitnil,omitempty" name:"RerunTimes"`

	// <p>运行的任务范围，任务ID列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SelectedTaskIds []*string `json:"SelectedTaskIds,omitnil,omitempty" name:"SelectedTaskIds"`

	// <p>资源组信息列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupInfoList []*ResourceGroupInfo `json:"ResourceGroupInfoList,omitnil,omitempty" name:"ResourceGroupInfoList"`

	// <p>标签列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelList []*LabelBrief `json:"LabelList,omitnil,omitempty" name:"LabelList"`

	// <p>父工作流运行ID 【由嵌套工作流触发独有】</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentWorkflowRunId *string `json:"ParentWorkflowRunId,omitnil,omitempty" name:"ParentWorkflowRunId"`

	// <p>父工作流任务运行ID 【由嵌套工作流触发独有】</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentWorkflowTaskRunId *string `json:"ParentWorkflowTaskRunId,omitnil,omitempty" name:"ParentWorkflowTaskRunId"`

	// <p>父工作流任务运行名称 【由嵌套工作流触发独有】</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentWorkflowTaskRunName *string `json:"ParentWorkflowTaskRunName,omitnil,omitempty" name:"ParentWorkflowTaskRunName"`

	// <p>授权权限类型<br>PERMISSION_TYPE_UNSPECIFIED：未指定权限<br>MANAGE : 管理权限：包含所有操作权限<br>RUN : 运行权限：可执行实体<br>VIEW : 查看权限：可查看实体内容</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Permission *string `json:"Permission,omitnil,omitempty" name:"Permission"`

	// <p>工作流高级运行时用户填入的参数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdvancedParameters []*AdvancedParameter `json:"AdvancedParameters,omitnil,omitempty" name:"AdvancedParameters"`

	// <p>计划调度时间</p><p>参数格式：毫秒时间戳（UTC）</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScheduledTime *string `json:"ScheduledTime,omitnil,omitempty" name:"ScheduledTime"`
}

type WorkflowRunBrief struct {
	// 工作流运行ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`

	// 运行开始时间，单位：毫秒时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunStartTime *string `json:"RunStartTime,omitnil,omitempty" name:"RunStartTime"`

	// <p>运行状态。CREATE(&quot;初始化&quot;),     QUEUED(&quot;等待中&quot;),     PENDING(&quot;准备中&quot;),     RUNNING(&quot;运行中&quot;),     SKIPPED(&quot;跳过运行&quot;),     SUCCESS(&quot;成功&quot;),     FAILED(&quot;失败&quot;),     TERMINATING(&quot;终止中&quot;),     TERMINATED(&quot;终止&quot;),     CANCELLED(&quot;被手动终止&quot;)等</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunState *string `json:"RunState,omitnil,omitempty" name:"RunState"`

	// 运行错误码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorCodeString *string `json:"ErrorCodeString,omitnil,omitempty" name:"ErrorCodeString"`
}

type WorkflowTask struct {
	// 任务参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamList []*ParamInfo `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// 任务依赖
	// 注意：此字段可能返回 null，表示取不到有效值。
	DependOnList []*DependOnBrief `json:"DependOnList,omitnil,omitempty" name:"DependOnList"`

	// 任务ID，创建时无需传入，由服务端生成
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 任务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskType *TaskType `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 资源组ID，可通过资源组相关接口获取
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// 任务描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 任务告警
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alarm *AlarmBrief `json:"Alarm,omitnil,omitempty" name:"Alarm"`

	// 监控指标
	// 注意：此字段可能返回 null，表示取不到有效值。
	MonitorMetric *MonitorMetricBrief `json:"MonitorMetric,omitnil,omitempty" name:"MonitorMetric"`

	// 任务重试策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskRetryStrategy *TaskRetryStrategy `json:"TaskRetryStrategy,omitnil,omitempty" name:"TaskRetryStrategy"`

	// <p>任依赖运行条件</p><ul><li>ALL_SUCCESS: 全部成功：所有上游依赖任务均已执行并成功</li><li>ONE_SUCCESS: 至少一个成功：至少有一个上游依赖任务成功</li><li>NONE_FAILED: 目前没有失败：没有依赖任务失败，并且至少有一个依赖任务在运行中</li><li>ALL_DONE: 全部完成：所有上游依赖任务均已执行并完成（无论成功或失败</li><li>ONE_FAILED: 至少一个失败：至少有一个上游依赖任务失败</li><li>ALL_FAILED: 全部失败：所有上游依赖任务都失败</li><li>ALL_DONE_AT_LEAST_ONE_SUCCESS：上游全部完成至少一个成功: 所有上游依赖任务都达到终态时，进行依赖判断，至少有一个成功，则依赖判断成功，否则就是跳过运行</li><li>ALL_SKIPPED：上游全部完成，没有跳过运行: 所有上游依赖任务都达到终态时，进行依赖判断, 如果上游状态全部都是成功、失败、上游失败状态，则依赖判断成功，否则为跳过运行</li><li>ONE_DONE：至少一个完成：上游只要有一个完成了，就进行依赖判断，且依赖判断成功，否则还是等待上游</li><li>ALL_DONE_NONE_FAILED_AT_LEAST_ONE_SUCCESS：上游全部完成，没有失败，至少有一个成功: 所有上游依赖任务都达到终态时，进行依赖判断，上游没有一个失败且至少有一个成功的情况下，依赖判断成功，否则就是跳过运行</li><li>NONE_SKIPPED：上游全部完成，没有跳过运行: 所有上游依赖任务都达到终态时，进行依赖判断, 如果上游状态全部都是成功、失败、上游失败状态，则依赖判断成功，否则为跳过运行</li><li>ALL_DONE_AT_LEAST_ONE_FAILED：上游全部完成至少一个失败: 所有上游依赖任务都达到终态时，进行依赖判断，至少有一个失败，则依赖判断成功，否则就是跳过运行</li><li>ADVANCED:运行条件为高级模式时配置</li></ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DependOnRunCondition *string `json:"DependOnRunCondition,omitnil,omitempty" name:"DependOnRunCondition"`

	// 任务X坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	LeftCoordinate *float64 `json:"LeftCoordinate,omitnil,omitempty" name:"LeftCoordinate"`

	// 任务Y坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopCoordinate *float64 `json:"TopCoordinate,omitnil,omitempty" name:"TopCoordinate"`

	// <p>任务高级运行参数，当DependOnRunCondition为ADVANCED时配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdvancedDependencyConfig *AdvancedDependencyConfig `json:"AdvancedDependencyConfig,omitnil,omitempty" name:"AdvancedDependencyConfig"`

	// <p>内嵌任务（FOR_EACH任务的子任务）</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerTask *WorkflowTask `json:"InnerTask,omitnil,omitempty" name:"InnerTask"`

	// <p>创建时间，单位：毫秒时间戳。系统生成字段，仅出参返回，入参无需传入（传入不生效，服务端忽略且不报错）</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>更新时间，单位：毫秒时间戳。系统生成字段，仅出参返回，入参无需传入（传入不生效，服务端忽略且不报错）</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>创建人UIN。系统生成字段，仅出参返回，入参无需传入（传入不生效，服务端忽略且不报错）</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`
}

type WorkflowTaskNodeBrief struct {
	// <p>工作流ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// <p>任务ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <p>任务名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// <p>任务类型名称，请参考数据结构TaskType中TaskTypeName字段描述</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTypeName *string `json:"TaskTypeName,omitnil,omitempty" name:"TaskTypeName"`

	// <p>任务依赖列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DependOnList []*DependOnBrief `json:"DependOnList,omitnil,omitempty" name:"DependOnList"`

	// <p>任务资源组ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// <p>任务资源组名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupName *string `json:"ResourceGroupName,omitnil,omitempty" name:"ResourceGroupName"`

	// <p>任务X坐标</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	LeftCoordinate *float64 `json:"LeftCoordinate,omitnil,omitempty" name:"LeftCoordinate"`

	// <p>任务Y坐标</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopCoordinate *float64 `json:"TopCoordinate,omitnil,omitempty" name:"TopCoordinate"`

	// <p>任务重试策略</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskRetryStrategy *TaskRetryStrategy `json:"TaskRetryStrategy,omitnil,omitempty" name:"TaskRetryStrategy"`

	// <p>任务依赖运行条件</p><p>ALL_SUCCESS: 全部成功：所有上游依赖任务均已执行并成功<br>ONE_SUCCESS: 至少一个成功：至少有一个上游依赖任务成功<br>NONE_FAILED: 目前没有失败：没有依赖任务失败，并且至少有一个依赖任务在运行中<br>ALL_DONE: 全部完成：所有上游依赖任务均已执行并完成（无论成功或失败<br>ONE_FAILED: 至少一个失败：至少有一个上游依赖任务失败<br>ALL_FAILED: 全部失败：所有上游依赖任务都失败<br>ALL_DONE_AT_LEAST_ONE_SUCCESS：上游全部完成至少一个成功: 所有上游依赖任务都达到终态时，进行依赖判断，至少有一个成功，则依赖判断成功，否则就是跳过运行<br>ALL_SKIPPED：上游全部完成，没有跳过运行: 所有上游依赖任务都达到终态时，进行依赖判断, 如果上游状态全部都是成功、失败、上游失败状态，则依赖判断成功，否则为跳过运行<br>ONE_DONE：至少一个完成：上游只要有一个完成了，就进行依赖判断，且依赖判断成功，否则还是等待上游<br>ALL_DONE_NONE_FAILED_AT_LEAST_ONE_SUCCESS：上游全部完成，没有失败，至少有一个成功: 所有上游依赖任务都达到终态时，进行依赖判断，上游没有一个失败且至少有一个成功的情况下，依赖判断成功，否则就是跳过运行<br>NONE_SKIPPED：上游全部完成，没有跳过运行: 所有上游依赖任务都达到终态时，进行依赖判断, 如果上游状态全部都是成功、失败、上游失败状态，则依赖判断成功，否则为跳过运行<br>ALL_DONE_AT_LEAST_ONE_FAILED：上游全部完成至少一个失败: 所有上游依赖任务都达到终态时，进行依赖判断，至少有一个失败，则依赖判断成功，否则就是跳过运行<br>ADVANCED:运行条件为高级模式时配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DependOnRunCondition *string `json:"DependOnRunCondition,omitnil,omitempty" name:"DependOnRunCondition"`

	// <p>高级依赖配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdvancedDependencyConfig *AdvancedDependencyConfig `json:"AdvancedDependencyConfig,omitnil,omitempty" name:"AdvancedDependencyConfig"`

	// <p>内嵌工作流任务节点</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerTask *WorkflowTaskNodeBrief `json:"InnerTask,omitnil,omitempty" name:"InnerTask"`
}

type WorkflowTaskRun struct {
	// <p>任务名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// <p>任务运行ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowTaskRunId *string `json:"WorkflowTaskRunId,omitnil,omitempty" name:"WorkflowTaskRunId"`

	// <p>运行状态。如CREATE("初始化"), QUEUED("等待中"), PENDING("准备中"), RUNNING("运行中"), SKIPPED("跳过运行"), SUCCESS("成功"), FAILED("失败"), TERMINATING("终止中"), TERMINATED("终止"), CANCELLED("被手动终止")等</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunState *string `json:"RunState,omitnil,omitempty" name:"RunState"`

	// <p>工作空间ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>工作流ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// <p>工作流运行ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`

	// <p>任务ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <p>任务类型名称，请参考数据结构TaskType中TaskTypeName字段描述</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTypeName *string `json:"TaskTypeName,omitnil,omitempty" name:"TaskTypeName"`

	// <p>任务版本ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskVersionId *string `json:"TaskVersionId,omitnil,omitempty" name:"TaskVersionId"`

	// <p>触发类型</p><p>枚举值：</p><ul><li>Manual： 手动触发</li><li>Scheduler： 调度触发</li></ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TriggerType *string `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// <p>所属资源组ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// <p>错误码</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorCodeString *string `json:"ErrorCodeString,omitnil,omitempty" name:"ErrorCodeString"`

	// <p>运行用户UIN</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunUserUin *string `json:"RunUserUin,omitnil,omitempty" name:"RunUserUin"`

	// <p>运行用户名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunUserName *string `json:"RunUserName,omitnil,omitempty" name:"RunUserName"`

	// <p>创建人UIN</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// <p>执行平台执行ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// <p>创建时间，单位：毫秒时间戳</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>更新时间，单位：毫秒时间戳</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>依赖任务完成时间，单位：毫秒时间戳</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DependenceFinishedTime *string `json:"DependenceFinishedTime,omitnil,omitempty" name:"DependenceFinishedTime"`

	// <p>运行开始时间，单位：毫秒时间戳</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunStartTime *string `json:"RunStartTime,omitnil,omitempty" name:"RunStartTime"`

	// <p>运行结束时间，单位：毫秒时间戳</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunEndTime *string `json:"RunEndTime,omitnil,omitempty" name:"RunEndTime"`

	// <p>运行时长，单位：秒</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunCostTime *string `json:"RunCostTime,omitnil,omitempty" name:"RunCostTime"`

	// <p>等待时长（依赖就绪到开始运行的等待耗时），单位：秒</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WaitTime *string `json:"WaitTime,omitnil,omitempty" name:"WaitTime"`

	// <p>下发执行平台时间，单位：毫秒时间戳</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IssueTime *string `json:"IssueTime,omitnil,omitempty" name:"IssueTime"`

	// <p>时区</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`

	// <p>依赖上游任务ID列表。保留字段，暂时返回为[]</p><p>保留字段，暂时返回为[]</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DependOnList []*string `json:"DependOnList,omitnil,omitempty" name:"DependOnList"`

	// <p>运行参数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunParams *string `json:"RunParams,omitnil,omitempty" name:"RunParams"`

	// <p>任务扩展信息，包含脚本路径</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTypeExtensions *string `json:"TaskTypeExtensions,omitnil,omitempty" name:"TaskTypeExtensions"`

	// <p>任务X坐标</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	LeftCoordinate *float64 `json:"LeftCoordinate,omitnil,omitempty" name:"LeftCoordinate"`

	// <p>任务Y坐标</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopCoordinate *float64 `json:"TopCoordinate,omitnil,omitempty" name:"TopCoordinate"`

	// <p>重试次数，为 0 则表示首次运行</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RetryTimes *int64 `json:"RetryTimes,omitnil,omitempty" name:"RetryTimes"`

	// <p>工作流名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// <p>重跑次数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RerunTimes *int64 `json:"RerunTimes,omitnil,omitempty" name:"RerunTimes"`

	// <p>是否最新一次运行</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsLatestRun *bool `json:"IsLatestRun,omitnil,omitempty" name:"IsLatestRun"`

	// <p>资源组信息列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupInfoList []*ResourceGroupInfo `json:"ResourceGroupInfoList,omitnil,omitempty" name:"ResourceGroupInfoList"`

	// <p>运行结果</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunResult *string `json:"RunResult,omitnil,omitempty" name:"RunResult"`

	// <p>任务依赖运行条件</p><p>ALL_SUCCESS: 全部成功：所有上游依赖任务均已执行并成功<br>ONE_SUCCESS: 至少一个成功：至少有一个上游依赖任务成功<br>NONE_FAILED: 目前没有失败：没有依赖任务失败，并且至少有一个依赖任务在运行中<br>ALL_DONE: 全部完成：所有上游依赖任务均已执行并完成（无论成功或失败<br>ONE_FAILED: 至少一个失败：至少有一个上游依赖任务失败<br>ALL_FAILED: 全部失败：所有上游依赖任务都失败<br>ALL_DONE_AT_LEAST_ONE_SUCCESS：上游全部完成至少一个成功: 所有上游依赖任务都达到终态时，进行依赖判断，至少有一个成功，则依赖判断成功，否则就是跳过运行<br>ALL_SKIPPED：上游全部完成，没有跳过运行: 所有上游依赖任务都达到终态时，进行依赖判断, 如果上游状态全部都是成功、失败、上游失败状态，则依赖判断成功，否则为跳过运行<br>ONE_DONE：至少一个完成：上游只要有一个完成了，就进行依赖判断，且依赖判断成功，否则还是等待上游<br>ALL_DONE_NONE_FAILED_AT_LEAST_ONE_SUCCESS：上游全部完成，没有失败，至少有一个成功: 所有上游依赖任务都达到终态时，进行依赖判断，上游没有一个失败且至少有一个成功的情况下，依赖判断成功，否则就是跳过运行<br>NONE_SKIPPED：上游全部完成，没有跳过运行: 所有上游依赖任务都达到终态时，进行依赖判断, 如果上游状态全部都是成功、失败、上游失败状态，则依赖判断成功，否则为跳过运行<br>ALL_DONE_AT_LEAST_ONE_FAILED：上游全部完成至少一个失败: 所有上游依赖任务都达到终态时，进行依赖判断，至少有一个失败，则依赖判断成功，否则就是跳过运行<br>ADVANCED:运行条件为高级模式时配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DependOnRunCondition *string `json:"DependOnRunCondition,omitnil,omitempty" name:"DependOnRunCondition"`

	// <p>高级依赖配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdvancedDependencyConfig *AdvancedDependencyConfig `json:"AdvancedDependencyConfig,omitnil,omitempty" name:"AdvancedDependencyConfig"`

	// <p>内嵌工作流任务信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerTask *InnerWorkflowTaskBrief `json:"InnerTask,omitnil,omitempty" name:"InnerTask"`

	// <p>计划调度时间</p><p>参数格式：毫秒时间戳，UTC</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScheduledTime *string `json:"ScheduledTime,omitnil,omitempty" name:"ScheduledTime"`
}

type WorkflowTriggerAdvancedConfiguration struct {
	// <p>&lt;p&gt;该工作流下的所有任务重试模式，仅当TriggerMode为CONTINUE_RUN时有效。</p><p>枚举值：</p><ul><li>onFailure： 失败时自动重试</li><li>never： 从不重试</li></ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskRetryMode *string `json:"TaskRetryMode,omitnil,omitempty" name:"TaskRetryMode"`
}

type WorkflowTriggerConfiguration struct {
	// <p>调度配置ID，创建时无需传入，由服务端生成</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TriggerId *string `json:"TriggerId,omitnil,omitempty" name:"TriggerId"`

	// <p>调度状态 启动：START，暂停：PAUSE</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchedulerStatus *string `json:"SchedulerStatus,omitnil,omitempty" name:"SchedulerStatus"`

	// <p>触发方式，</p><ul><li>定时触发：TIME_TRIGGER</li><li>持续运行：CONTINUE_RUN</li></ul><p>注意：</p><ul><li>TIME_TRIGGER 模式下，SchedulerStatus、SchedulerTimeZone、StartTime、EndTime、ConfigMode、CycleType、CrontabExpression 必填；</li><li>CONTINUE_RUN 模式下，AdvancedConfig必填；</li></ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TriggerMode *string `json:"TriggerMode,omitnil,omitempty" name:"TriggerMode"`

	// <p>调度时区</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchedulerTimeZone *string `json:"SchedulerTimeZone,omitnil,omitempty" name:"SchedulerTimeZone"`

	// <p>调度生效时间，单位：毫秒时间戳。必须小于 EndTime</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>调度结束时间，单位：毫秒时间戳。必须大于 StartTime</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>配置方式，常规：COMMON，CRON表达式：CRON_EXPRESSION</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigMode *string `json:"ConfigMode,omitnil,omitempty" name:"ConfigMode"`

	// <p>周期类型：支持的类型为 ONEOFF_CYCLE: 一次性 YEAR_CYCLE: 年 MONTH_CYCLE: 月 WEEK_CYCLE: 周 DAY_CYCLE: 天<br>HOUR_CYCLE: 小时 MINUTE_CYCLE: 分钟 CRONTAB_CYCLE: crontab表达式类型</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleType *string `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// <p>cron表达式</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CrontabExpression *string `json:"CrontabExpression,omitnil,omitempty" name:"CrontabExpression"`

	// <p>Json格式，对账使用</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraInfo *string `json:"ExtraInfo,omitnil,omitempty" name:"ExtraInfo"`

	// <p>高级配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdvancedConfig *WorkflowTriggerAdvancedConfiguration `json:"AdvancedConfig,omitnil,omitempty" name:"AdvancedConfig"`
}

type WorkspaceGroupInfo struct {
	// <p>用户组ID</p>
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// <p>用户组名称</p>
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// <p>角色列表</p>
	Roles []*RoleBasicInfo `json:"Roles,omitnil,omitempty" name:"Roles"`

	// <p>创建时间，毫秒时间戳</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>更新时间，毫秒时间戳</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>用户组下用户数量</p>
	UserCount *int64 `json:"UserCount,omitnil,omitempty" name:"UserCount"`

	// <p>用户组类型。取值为枚举数值的字符串形式：&quot;0&quot;=控制台系统类型（包含全部user）、&quot;1&quot;=控制台自定义类型、&quot;2&quot;=工作空间系统类型（包含工作空间全部user）、&quot;3&quot;=工作空间自定义类型</p>
	GroupType *string `json:"GroupType,omitnil,omitempty" name:"GroupType"`

	// <p>该组直接绑定的标签（不含继承）</p>
	Tags []*UserTagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type WorkspaceInfo struct {
	// 工作空间ID
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 工作空间名称
	WorkspaceName *string `json:"WorkspaceName,omitnil,omitempty" name:"WorkspaceName"`

	// 工作空间描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 工作空间地域（如 ap-guangzhou）
	WorkspaceRegion *string `json:"WorkspaceRegion,omitnil,omitempty" name:"WorkspaceRegion"`

	// 工作空间状态：0=未指定 1=创建中 2=创建失败 3=正常运行中 4=已删除
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 失败原因（Status=2 创建失败时有值）
	ErrorReason *string `json:"ErrorReason,omitnil,omitempty" name:"ErrorReason"`

	// 创建者信息
	Creator *StandardUserInfo `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 创建时间，毫秒时间戳
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间，毫秒时间戳
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 当前用户是否拥有该工作空间的访问权限
	HasAccess *bool `json:"HasAccess,omitnil,omitempty" name:"HasAccess"`
}

type WorkspaceUserInfo struct {
	// <p>用户UIN</p>
	UserUin *string `json:"UserUin,omitnil,omitempty" name:"UserUin"`

	// <p>用户名</p>
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// <p>昵称</p>
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// <p>角色列表</p>
	Roles []*RoleBasicInfo `json:"Roles,omitnil,omitempty" name:"Roles"`

	// <p>用户来源，group：用户组、user：用户</p>
	UserSource *string `json:"UserSource,omitnil,omitempty" name:"UserSource"`

	// <p>创建时间，毫秒时间戳</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>更新时间，毫秒时间戳</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>是否主账号</p>
	IsOwner *bool `json:"IsOwner,omitnil,omitempty" name:"IsOwner"`

	// <p>0: 普通用户 1: entraId用户</p>
	UserTag *int64 `json:"UserTag,omitnil,omitempty" name:"UserTag"`

	// <p>是否具有 admin 权限的子账号</p>
	IsAdmin *bool `json:"IsAdmin,omitnil,omitempty" name:"IsAdmin"`

	// <p>合并后标签（用户直绑 + 用户组继承）</p>
	Tags []*UserTagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`
}