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

package v20231130

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AppConfig struct {
	// 知识问答管理应用配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	KnowledgeQa *KnowledgeQaConfig `json:"KnowledgeQa,omitnil,omitempty" name:"KnowledgeQa"`

	// 知识摘要应用配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Summary *SummaryConfig `json:"Summary,omitnil,omitempty" name:"Summary"`

	// 标签提取应用配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Classify *ClassifyConfig `json:"Classify,omitnil,omitempty" name:"Classify"`
}

type AppInfo struct {
	// 应用类型；knowledge_qa-知识问答管理；summary-知识摘要；classifys-知识标签提取
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`

	// 应用类型描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppTypeDesc *string `json:"AppTypeDesc,omitnil,omitempty" name:"AppTypeDesc"`

	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 应用头像
	// 注意：此字段可能返回 null，表示取不到有效值。
	Avatar *string `json:"Avatar,omitnil,omitempty" name:"Avatar"`

	// 应用描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 应用状态，1：未上线，2：运行中，3：停用
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppStatus *uint64 `json:"AppStatus,omitnil,omitempty" name:"AppStatus"`

	// 状态说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppStatusDesc *string `json:"AppStatusDesc,omitnil,omitempty" name:"AppStatusDesc"`

	// 修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 最后修改人
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 模型名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 模型别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelAliasName *string `json:"ModelAliasName,omitnil,omitempty" name:"ModelAliasName"`
}

type AppModel struct {
	// 模型名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模型描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 上下文指代轮次
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContextLimit *uint64 `json:"ContextLimit,omitnil,omitempty" name:"ContextLimit"`

	// 模型别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	AliasName *string `json:"AliasName,omitnil,omitempty" name:"AliasName"`
}

type AttrLabel struct {
	// 属性标签来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *uint64 `json:"Source,omitnil,omitempty" name:"Source"`

	// 属性ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttrBizId *string `json:"AttrBizId,omitnil,omitempty" name:"AttrBizId"`

	// 属性标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttrKey *string `json:"AttrKey,omitnil,omitempty" name:"AttrKey"`

	// 属性名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttrName *string `json:"AttrName,omitnil,omitempty" name:"AttrName"`

	// 标签ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Labels []*Label `json:"Labels,omitnil,omitempty" name:"Labels"`
}

type AttrLabelDetail struct {
	// 属性ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttrBizId *string `json:"AttrBizId,omitnil,omitempty" name:"AttrBizId"`

	// 属性标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttrKey *string `json:"AttrKey,omitnil,omitempty" name:"AttrKey"`

	// 属性名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttrName *string `json:"AttrName,omitnil,omitempty" name:"AttrName"`

	// 标签名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelNames []*string `json:"LabelNames,omitnil,omitempty" name:"LabelNames"`

	// 属性标签是否在更新中
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsUpdating *bool `json:"IsUpdating,omitnil,omitempty" name:"IsUpdating"`
}

type AttrLabelRefer struct {
	// 属性标签来源，1：属性标签
	Source *uint64 `json:"Source,omitnil,omitempty" name:"Source"`

	// 属性ID
	AttributeBizId *string `json:"AttributeBizId,omitnil,omitempty" name:"AttributeBizId"`

	// 标签ID
	LabelBizIds []*string `json:"LabelBizIds,omitnil,omitempty" name:"LabelBizIds"`
}

type AttributeFilters struct {
	// 检索，属性或标签名称
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`
}

type AttributeLabel struct {
	// 标签ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelBizId *string `json:"LabelBizId,omitnil,omitempty" name:"LabelBizId"`

	// 标签名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelName *string `json:"LabelName,omitnil,omitempty" name:"LabelName"`

	// 相似标签名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SimilarLabels []*string `json:"SimilarLabels,omitnil,omitempty" name:"SimilarLabels"`
}

type BaseConfig struct {
	// 应用名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 机器人头像
	Avatar *string `json:"Avatar,omitnil,omitempty" name:"Avatar"`

	// 机器人描述
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`
}

// Predefined struct for user
type CheckAttributeLabelExistRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 属性名称
	LabelName *string `json:"LabelName,omitnil,omitempty" name:"LabelName"`

	// 属性ID
	AttributeBizId *string `json:"AttributeBizId,omitnil,omitempty" name:"AttributeBizId"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 滚动加载，最后一个属性标签ID
	LastLabelBizId *string `json:"LastLabelBizId,omitnil,omitempty" name:"LastLabelBizId"`
}

type CheckAttributeLabelExistRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 属性名称
	LabelName *string `json:"LabelName,omitnil,omitempty" name:"LabelName"`

	// 属性ID
	AttributeBizId *string `json:"AttributeBizId,omitnil,omitempty" name:"AttributeBizId"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 滚动加载，最后一个属性标签ID
	LastLabelBizId *string `json:"LastLabelBizId,omitnil,omitempty" name:"LastLabelBizId"`
}

func (r *CheckAttributeLabelExistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckAttributeLabelExistRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "LabelName")
	delete(f, "AttributeBizId")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	delete(f, "LastLabelBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckAttributeLabelExistRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckAttributeLabelExistResponseParams struct {
	// 是否存在
	IsExist *bool `json:"IsExist,omitnil,omitempty" name:"IsExist"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckAttributeLabelExistResponse struct {
	*tchttp.BaseResponse
	Response *CheckAttributeLabelExistResponseParams `json:"Response"`
}

func (r *CheckAttributeLabelExistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckAttributeLabelExistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckAttributeLabelReferRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 属性标签
	LabelBizId *string `json:"LabelBizId,omitnil,omitempty" name:"LabelBizId"`

	// 属性ID
	AttributeBizId []*string `json:"AttributeBizId,omitnil,omitempty" name:"AttributeBizId"`
}

type CheckAttributeLabelReferRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 属性标签
	LabelBizId *string `json:"LabelBizId,omitnil,omitempty" name:"LabelBizId"`

	// 属性ID
	AttributeBizId []*string `json:"AttributeBizId,omitnil,omitempty" name:"AttributeBizId"`
}

func (r *CheckAttributeLabelReferRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckAttributeLabelReferRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	delete(f, "LabelBizId")
	delete(f, "AttributeBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckAttributeLabelReferRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckAttributeLabelReferResponseParams struct {
	// 是否引用
	IsRefer *bool `json:"IsRefer,omitnil,omitempty" name:"IsRefer"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckAttributeLabelReferResponse struct {
	*tchttp.BaseResponse
	Response *CheckAttributeLabelReferResponseParams `json:"Response"`
}

func (r *CheckAttributeLabelReferResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckAttributeLabelReferResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClassifyConfig struct {
	// 模型配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Model *AppModel `json:"Model,omitnil,omitempty" name:"Model"`

	// 标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Labels []*ClassifyLabel `json:"Labels,omitnil,omitempty" name:"Labels"`
}

type ClassifyLabel struct {
	// 标签名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 标签描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 标签取值范围
	// 注意：此字段可能返回 null，表示取不到有效值。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type Context struct {
	// 消息记录ID信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordBizId *string `json:"RecordBizId,omitnil,omitempty" name:"RecordBizId"`

	// 是否为用户
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsVisitor *bool `json:"IsVisitor,omitnil,omitempty" name:"IsVisitor"`

	// 昵称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 头像
	// 注意：此字段可能返回 null，表示取不到有效值。
	Avatar *string `json:"Avatar,omitnil,omitempty" name:"Avatar"`

	// 消息内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

// Predefined struct for user
type CreateAppRequestParams struct {
	// 应用类型；knowledge_qa-知识问答管理；summary-知识摘要；classifys-知识标签提取
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`

	// 应用基础配置
	BaseConfig *BaseConfig `json:"BaseConfig,omitnil,omitempty" name:"BaseConfig"`
}

type CreateAppRequest struct {
	*tchttp.BaseRequest
	
	// 应用类型；knowledge_qa-知识问答管理；summary-知识摘要；classifys-知识标签提取
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`

	// 应用基础配置
	BaseConfig *BaseConfig `json:"BaseConfig,omitnil,omitempty" name:"BaseConfig"`
}

func (r *CreateAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppType")
	delete(f, "BaseConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAppResponseParams struct {
	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAppResponse struct {
	*tchttp.BaseResponse
	Response *CreateAppResponseParams `json:"Response"`
}

func (r *CreateAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAttributeLabelRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 属性标识
	AttrKey *string `json:"AttrKey,omitnil,omitempty" name:"AttrKey"`

	// 属性名称
	AttrName *string `json:"AttrName,omitnil,omitempty" name:"AttrName"`

	// 属性标签
	Labels []*AttributeLabel `json:"Labels,omitnil,omitempty" name:"Labels"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

type CreateAttributeLabelRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 属性标识
	AttrKey *string `json:"AttrKey,omitnil,omitempty" name:"AttrKey"`

	// 属性名称
	AttrName *string `json:"AttrName,omitnil,omitempty" name:"AttrName"`

	// 属性标签
	Labels []*AttributeLabel `json:"Labels,omitnil,omitempty" name:"Labels"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

func (r *CreateAttributeLabelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAttributeLabelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "AttrKey")
	delete(f, "AttrName")
	delete(f, "Labels")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAttributeLabelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAttributeLabelResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAttributeLabelResponse struct {
	*tchttp.BaseResponse
	Response *CreateAttributeLabelResponseParams `json:"Response"`
}

func (r *CreateAttributeLabelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAttributeLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCorpRequestParams struct {
	// 企业全称
	FullName *string `json:"FullName,omitnil,omitempty" name:"FullName"`

	// 联系人名称
	ContactName *string `json:"ContactName,omitnil,omitempty" name:"ContactName"`

	// 联系人邮箱
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 联系人手机号
	Telephone *string `json:"Telephone,omitnil,omitempty" name:"Telephone"`
}

type CreateCorpRequest struct {
	*tchttp.BaseRequest
	
	// 企业全称
	FullName *string `json:"FullName,omitnil,omitempty" name:"FullName"`

	// 联系人名称
	ContactName *string `json:"ContactName,omitnil,omitempty" name:"ContactName"`

	// 联系人邮箱
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 联系人手机号
	Telephone *string `json:"Telephone,omitnil,omitempty" name:"Telephone"`
}

func (r *CreateCorpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCorpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FullName")
	delete(f, "ContactName")
	delete(f, "Email")
	delete(f, "Telephone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCorpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCorpResponseParams struct {
	// 企业ID
	CorpBizId *string `json:"CorpBizId,omitnil,omitempty" name:"CorpBizId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCorpResponse struct {
	*tchttp.BaseResponse
	Response *CreateCorpResponseParams `json:"Response"`
}

func (r *CreateCorpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCorpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateQACateRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 父级业务ID
	ParentBizId *string `json:"ParentBizId,omitnil,omitempty" name:"ParentBizId"`

	// 分类名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type CreateQACateRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 父级业务ID
	ParentBizId *string `json:"ParentBizId,omitnil,omitempty" name:"ParentBizId"`

	// 分类名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *CreateQACateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateQACateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "ParentBizId")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateQACateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateQACateResponseParams struct {
	// 是否可新增
	CanAdd *bool `json:"CanAdd,omitnil,omitempty" name:"CanAdd"`

	// 是否可编辑
	CanEdit *bool `json:"CanEdit,omitnil,omitempty" name:"CanEdit"`

	// 是否可删除
	CanDelete *bool `json:"CanDelete,omitnil,omitempty" name:"CanDelete"`

	// 分类业务ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateQACateResponse struct {
	*tchttp.BaseResponse
	Response *CreateQACateResponseParams `json:"Response"`
}

func (r *CreateQACateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateQACateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateQARequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 问题
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 答案
	Answer *string `json:"Answer,omitnil,omitempty" name:"Answer"`

	// 属性标签适用范围 1：全部，2：按条件
	AttrRange *uint64 `json:"AttrRange,omitnil,omitempty" name:"AttrRange"`

	// 自定义参数
	CustomParam *string `json:"CustomParam,omitnil,omitempty" name:"CustomParam"`

	// 属性标签引用
	AttrLabels []*AttrLabelRefer `json:"AttrLabels,omitnil,omitempty" name:"AttrLabels"`

	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// 分类ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`

	// 有效开始时间，unix时间戳
	ExpireStart *string `json:"ExpireStart,omitnil,omitempty" name:"ExpireStart"`

	// 有效结束时间，unix时间戳，0代表永久有效
	ExpireEnd *string `json:"ExpireEnd,omitnil,omitempty" name:"ExpireEnd"`
}

type CreateQARequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 问题
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 答案
	Answer *string `json:"Answer,omitnil,omitempty" name:"Answer"`

	// 属性标签适用范围 1：全部，2：按条件
	AttrRange *uint64 `json:"AttrRange,omitnil,omitempty" name:"AttrRange"`

	// 自定义参数
	CustomParam *string `json:"CustomParam,omitnil,omitempty" name:"CustomParam"`

	// 属性标签引用
	AttrLabels []*AttrLabelRefer `json:"AttrLabels,omitnil,omitempty" name:"AttrLabels"`

	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// 分类ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`

	// 有效开始时间，unix时间戳
	ExpireStart *string `json:"ExpireStart,omitnil,omitempty" name:"ExpireStart"`

	// 有效结束时间，unix时间戳，0代表永久有效
	ExpireEnd *string `json:"ExpireEnd,omitnil,omitempty" name:"ExpireEnd"`
}

func (r *CreateQARequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateQARequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "Question")
	delete(f, "Answer")
	delete(f, "AttrRange")
	delete(f, "CustomParam")
	delete(f, "AttrLabels")
	delete(f, "DocBizId")
	delete(f, "CateBizId")
	delete(f, "ExpireStart")
	delete(f, "ExpireEnd")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateQARequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateQAResponseParams struct {
	// 问答ID
	QaBizId *string `json:"QaBizId,omitnil,omitempty" name:"QaBizId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateQAResponse struct {
	*tchttp.BaseResponse
	Response *CreateQAResponseParams `json:"Response"`
}

func (r *CreateQAResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateQAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRejectedQuestionRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 拒答问题
	// 
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 拒答问题来源的数据源唯一id，取值1，2
	// 
	BusinessSource *uint64 `json:"BusinessSource,omitnil,omitempty" name:"BusinessSource"`

	// 拒答问题来源的数据源唯一id
	// 
	BusinessId *string `json:"BusinessId,omitnil,omitempty" name:"BusinessId"`
}

type CreateRejectedQuestionRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 拒答问题
	// 
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 拒答问题来源的数据源唯一id，取值1，2
	// 
	BusinessSource *uint64 `json:"BusinessSource,omitnil,omitempty" name:"BusinessSource"`

	// 拒答问题来源的数据源唯一id
	// 
	BusinessId *string `json:"BusinessId,omitnil,omitempty" name:"BusinessId"`
}

func (r *CreateRejectedQuestionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRejectedQuestionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "Question")
	delete(f, "BusinessSource")
	delete(f, "BusinessId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRejectedQuestionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRejectedQuestionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRejectedQuestionResponse struct {
	*tchttp.BaseResponse
	Response *CreateRejectedQuestionResponseParams `json:"Response"`
}

func (r *CreateRejectedQuestionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRejectedQuestionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReleaseRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 发布描述
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`
}

type CreateReleaseRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 发布描述
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`
}

func (r *CreateReleaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReleaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "Desc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateReleaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReleaseResponseParams struct {
	// 发布ID
	ReleaseBizId *string `json:"ReleaseBizId,omitnil,omitempty" name:"ReleaseBizId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateReleaseResponse struct {
	*tchttp.BaseResponse
	Response *CreateReleaseResponseParams `json:"Response"`
}

func (r *CreateReleaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReleaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Credentials struct {
	// token
	// 注意：此字段可能返回 null，表示取不到有效值。
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// 临时证书密钥ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TmpSecretId *string `json:"TmpSecretId,omitnil,omitempty" name:"TmpSecretId"`

	// 临时证书密钥Key
	// 注意：此字段可能返回 null，表示取不到有效值。
	TmpSecretKey *string `json:"TmpSecretKey,omitnil,omitempty" name:"TmpSecretKey"`
}

// Predefined struct for user
type DeleteAppRequestParams struct {
	// 应用ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 应用类型；knowledge_qa-知识问答管理；summary-知识摘要；classifys-知识标签提取
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`
}

type DeleteAppRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 应用类型；knowledge_qa-知识问答管理；summary-知识摘要；classifys-知识标签提取
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`
}

func (r *DeleteAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppBizId")
	delete(f, "AppType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAppResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAppResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAppResponseParams `json:"Response"`
}

func (r *DeleteAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAttributeLabelRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 属性ID
	AttributeBizIds []*string `json:"AttributeBizIds,omitnil,omitempty" name:"AttributeBizIds"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

type DeleteAttributeLabelRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 属性ID
	AttributeBizIds []*string `json:"AttributeBizIds,omitnil,omitempty" name:"AttributeBizIds"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

func (r *DeleteAttributeLabelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAttributeLabelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "AttributeBizIds")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAttributeLabelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAttributeLabelResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAttributeLabelResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAttributeLabelResponseParams `json:"Response"`
}

func (r *DeleteAttributeLabelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAttributeLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDocRequestParams struct {
	// 文档业务ID列表
	DocBizIds []*string `json:"DocBizIds,omitnil,omitempty" name:"DocBizIds"`

	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`
}

type DeleteDocRequest struct {
	*tchttp.BaseRequest
	
	// 文档业务ID列表
	DocBizIds []*string `json:"DocBizIds,omitnil,omitempty" name:"DocBizIds"`

	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`
}

func (r *DeleteDocRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDocRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DocBizIds")
	delete(f, "BotBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDocRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDocResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDocResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDocResponseParams `json:"Response"`
}

func (r *DeleteDocResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDocResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteQACateRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 分类业务ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`
}

type DeleteQACateRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 分类业务ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`
}

func (r *DeleteQACateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteQACateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "CateBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteQACateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteQACateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteQACateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteQACateResponseParams `json:"Response"`
}

func (r *DeleteQACateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteQACateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteQARequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 问答ID
	QaBizIds []*string `json:"QaBizIds,omitnil,omitempty" name:"QaBizIds"`
}

type DeleteQARequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 问答ID
	QaBizIds []*string `json:"QaBizIds,omitnil,omitempty" name:"QaBizIds"`
}

func (r *DeleteQARequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteQARequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "QaBizIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteQARequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteQAResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteQAResponse struct {
	*tchttp.BaseResponse
	Response *DeleteQAResponseParams `json:"Response"`
}

func (r *DeleteQAResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteQAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRejectedQuestionRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 拒答问题来源的数据源唯一id
	// 
	// 
	RejectedBizIds []*string `json:"RejectedBizIds,omitnil,omitempty" name:"RejectedBizIds"`
}

type DeleteRejectedQuestionRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 拒答问题来源的数据源唯一id
	// 
	// 
	RejectedBizIds []*string `json:"RejectedBizIds,omitnil,omitempty" name:"RejectedBizIds"`
}

func (r *DeleteRejectedQuestionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRejectedQuestionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "RejectedBizIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRejectedQuestionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRejectedQuestionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRejectedQuestionResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRejectedQuestionResponseParams `json:"Response"`
}

func (r *DeleteRejectedQuestionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRejectedQuestionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppRequestParams struct {
	// 应用ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 应用类型；knowledge_qa-知识问答管理；summary-知识摘要；classifys-知识标签提取
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`

	// 是否发布后的配置
	IsRelease *bool `json:"IsRelease,omitnil,omitempty" name:"IsRelease"`
}

type DescribeAppRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 应用类型；knowledge_qa-知识问答管理；summary-知识摘要；classifys-知识标签提取
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`

	// 是否发布后的配置
	IsRelease *bool `json:"IsRelease,omitnil,omitempty" name:"IsRelease"`
}

func (r *DescribeAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppBizId")
	delete(f, "AppType")
	delete(f, "IsRelease")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppResponseParams struct {
	// 应用 ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 应用类型；knowledge_qa-知识问答管理；summary-知识摘要；classifys-知识标签提取
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`

	// 应用类型说明
	AppTypeDesc *string `json:"AppTypeDesc,omitnil,omitempty" name:"AppTypeDesc"`

	// 应用类型说明
	BaseConfig *BaseConfig `json:"BaseConfig,omitnil,omitempty" name:"BaseConfig"`

	// 应用配置
	AppConfig *AppConfig `json:"AppConfig,omitnil,omitempty" name:"AppConfig"`

	// 头像是否在申诉中
	AvatarInAppeal *bool `json:"AvatarInAppeal,omitnil,omitempty" name:"AvatarInAppeal"`

	// 角色描述是否在申诉中
	RoleInAppeal *bool `json:"RoleInAppeal,omitnil,omitempty" name:"RoleInAppeal"`

	// 名称是否在申诉中
	NameInAppeal *bool `json:"NameInAppeal,omitnil,omitempty" name:"NameInAppeal"`

	// 欢迎语是否在申诉中
	GreetingInAppeal *bool `json:"GreetingInAppeal,omitnil,omitempty" name:"GreetingInAppeal"`

	// 未知问题回复语是否在申诉中
	BareAnswerInAppeal *bool `json:"BareAnswerInAppeal,omitnil,omitempty" name:"BareAnswerInAppeal"`

	// 应用appKey
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAppResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAppResponseParams `json:"Response"`
}

func (r *DescribeAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAttributeLabelRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 属性ID
	AttributeBizId *string `json:"AttributeBizId,omitnil,omitempty" name:"AttributeBizId"`

	// 每次加载的数量 
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 查询标签或相似标签
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 滚动加载游标的标签ID
	LastLabelBizId *string `json:"LastLabelBizId,omitnil,omitempty" name:"LastLabelBizId"`
}

type DescribeAttributeLabelRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 属性ID
	AttributeBizId *string `json:"AttributeBizId,omitnil,omitempty" name:"AttributeBizId"`

	// 每次加载的数量 
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 查询标签或相似标签
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 滚动加载游标的标签ID
	LastLabelBizId *string `json:"LastLabelBizId,omitnil,omitempty" name:"LastLabelBizId"`
}

func (r *DescribeAttributeLabelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAttributeLabelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "AttributeBizId")
	delete(f, "Limit")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	delete(f, "Query")
	delete(f, "LastLabelBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAttributeLabelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAttributeLabelResponseParams struct {
	// 属性ID
	AttributeBizId *string `json:"AttributeBizId,omitnil,omitempty" name:"AttributeBizId"`

	// 属性标识
	AttrKey *string `json:"AttrKey,omitnil,omitempty" name:"AttrKey"`

	// 属性名称
	AttrName *string `json:"AttrName,omitnil,omitempty" name:"AttrName"`

	// 标签数量
	LabelNumber *string `json:"LabelNumber,omitnil,omitempty" name:"LabelNumber"`

	// 标签名称
	Labels []*AttributeLabel `json:"Labels,omitnil,omitempty" name:"Labels"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAttributeLabelResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAttributeLabelResponseParams `json:"Response"`
}

func (r *DescribeAttributeLabelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAttributeLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCorpRequestParams struct {

}

type DescribeCorpRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeCorpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCorpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCorpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCorpResponseParams struct {
	// 企业ID
	CorpBizId *string `json:"CorpBizId,omitnil,omitempty" name:"CorpBizId"`

	// 机器人配额
	RobotQuota *uint64 `json:"RobotQuota,omitnil,omitempty" name:"RobotQuota"`

	// 企业全称
	FullName *string `json:"FullName,omitnil,omitempty" name:"FullName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCorpResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCorpResponseParams `json:"Response"`
}

func (r *DescribeCorpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCorpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDocRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`
}

type DescribeDocRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`
}

func (r *DescribeDocRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDocRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "DocBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDocRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDocResponseParams struct {
	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// 文件名称
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文件类型
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// cos路径
	CosUrl *string `json:"CosUrl,omitnil,omitempty" name:"CosUrl"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 文档状态(5审核通过 7审核中 8审核不通过 9审核通过 10待发布 11发布中 12发布成功 13学习中 14学习失败)
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 文档状态描述
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// 生成失败原因
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 答案中是否引用
	IsRefer *bool `json:"IsRefer,omitnil,omitempty" name:"IsRefer"`

	// 问答对数量
	QaNum *int64 `json:"QaNum,omitnil,omitempty" name:"QaNum"`

	// 是否删除
	IsDeleted *bool `json:"IsDeleted,omitnil,omitempty" name:"IsDeleted"`

	// 文档来源
	Source *int64 `json:"Source,omitnil,omitempty" name:"Source"`

	// 文档来源描述
	SourceDesc *string `json:"SourceDesc,omitnil,omitempty" name:"SourceDesc"`

	// 是否允许重新生成
	IsAllowRestart *bool `json:"IsAllowRestart,omitnil,omitempty" name:"IsAllowRestart"`

	// qa是否已删除
	IsDeletedQa *bool `json:"IsDeletedQa,omitnil,omitempty" name:"IsDeletedQa"`

	// 问答是否生成中
	IsCreatingQa *bool `json:"IsCreatingQa,omitnil,omitempty" name:"IsCreatingQa"`

	// 是否允许删除
	IsAllowDelete *bool `json:"IsAllowDelete,omitnil,omitempty" name:"IsAllowDelete"`

	// 是否允许操作引用开关
	IsAllowRefer *bool `json:"IsAllowRefer,omitnil,omitempty" name:"IsAllowRefer"`

	// 是否生成过问答
	IsCreatedQa *bool `json:"IsCreatedQa,omitnil,omitempty" name:"IsCreatedQa"`

	// 文档字符量
	DocCharSize *string `json:"DocCharSize,omitnil,omitempty" name:"DocCharSize"`

	// 是否允许编辑
	IsAllowEdit *bool `json:"IsAllowEdit,omitnil,omitempty" name:"IsAllowEdit"`

	// 属性标签适用范围 1：全部，2：按条件范围
	AttrRange *int64 `json:"AttrRange,omitnil,omitempty" name:"AttrRange"`

	// 属性标签
	AttrLabels []*AttrLabel `json:"AttrLabels,omitnil,omitempty" name:"AttrLabels"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDocResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDocResponseParams `json:"Response"`
}

func (r *DescribeDocResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDocResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQARequestParams struct {
	// QA业务ID
	QaBizId *string `json:"QaBizId,omitnil,omitempty" name:"QaBizId"`

	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`
}

type DescribeQARequest struct {
	*tchttp.BaseRequest
	
	// QA业务ID
	QaBizId *string `json:"QaBizId,omitnil,omitempty" name:"QaBizId"`

	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`
}

func (r *DescribeQARequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQARequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "QaBizId")
	delete(f, "BotBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeQARequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQAResponseParams struct {
	// QA业务ID
	QaBizId *string `json:"QaBizId,omitnil,omitempty" name:"QaBizId"`

	// 问题
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 答案
	Answer *string `json:"Answer,omitnil,omitempty" name:"Answer"`

	// 自定义参数
	CustomParam *string `json:"CustomParam,omitnil,omitempty" name:"CustomParam"`

	// 来源
	Source *uint64 `json:"Source,omitnil,omitempty" name:"Source"`

	// 来源描述
	SourceDesc *string `json:"SourceDesc,omitnil,omitempty" name:"SourceDesc"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 状态描述
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// 分类ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`

	// 是否允许校验
	IsAllowAccept *bool `json:"IsAllowAccept,omitnil,omitempty" name:"IsAllowAccept"`

	// 是否允许删除
	IsAllowDelete *bool `json:"IsAllowDelete,omitnil,omitempty" name:"IsAllowDelete"`

	// 是否允许编辑
	IsAllowEdit *bool `json:"IsAllowEdit,omitnil,omitempty" name:"IsAllowEdit"`

	// 文档id
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// 文档名称
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文档类型
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 分片ID
	SegmentBizId *string `json:"SegmentBizId,omitnil,omitempty" name:"SegmentBizId"`

	// 分片内容
	PageContent *string `json:"PageContent,omitnil,omitempty" name:"PageContent"`

	// 分片高亮内容
	Highlights []*Highlight `json:"Highlights,omitnil,omitempty" name:"Highlights"`

	// 分片内容
	OrgData *string `json:"OrgData,omitnil,omitempty" name:"OrgData"`

	// 属性标签适用范围
	AttrRange *int64 `json:"AttrRange,omitnil,omitempty" name:"AttrRange"`

	// 属性标签
	AttrLabels []*AttrLabel `json:"AttrLabels,omitnil,omitempty" name:"AttrLabels"`

	// 有效开始时间，unix时间戳
	ExpireStart *string `json:"ExpireStart,omitnil,omitempty" name:"ExpireStart"`

	// 有效结束时间，unix时间戳，0代表永久有效
	ExpireEnd *string `json:"ExpireEnd,omitnil,omitempty" name:"ExpireEnd"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeQAResponse struct {
	*tchttp.BaseResponse
	Response *DescribeQAResponseParams `json:"Response"`
}

func (r *DescribeQAResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReferRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 引用ID
	ReferBizIds []*string `json:"ReferBizIds,omitnil,omitempty" name:"ReferBizIds"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

type DescribeReferRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 引用ID
	ReferBizIds []*string `json:"ReferBizIds,omitnil,omitempty" name:"ReferBizIds"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

func (r *DescribeReferRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReferRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "ReferBizIds")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReferRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReferResponseParams struct {
	// 引用列表
	List []*ReferDetail `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeReferResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReferResponseParams `json:"Response"`
}

func (r *DescribeReferResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReferResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReleaseInfoRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`
}

type DescribeReleaseInfoRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`
}

func (r *DescribeReleaseInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReleaseInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReleaseInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReleaseInfoResponseParams struct {
	// 最后发布时间
	LastTime *string `json:"LastTime,omitnil,omitempty" name:"LastTime"`

	// 发布状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 是否编辑过
	IsUpdated *bool `json:"IsUpdated,omitnil,omitempty" name:"IsUpdated"`

	// 失败原因
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeReleaseInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReleaseInfoResponseParams `json:"Response"`
}

func (r *DescribeReleaseInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReleaseInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReleaseRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 发布详情
	ReleaseBizId *string `json:"ReleaseBizId,omitnil,omitempty" name:"ReleaseBizId"`
}

type DescribeReleaseRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 发布详情
	ReleaseBizId *string `json:"ReleaseBizId,omitnil,omitempty" name:"ReleaseBizId"`
}

func (r *DescribeReleaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReleaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "ReleaseBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReleaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReleaseResponseParams struct {
	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 发布描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 发布状态(1待发布 2发布中 3发布成功 4发布失败 5发布中 6发布中 7发布失败 9发布暂停)
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 发布状态描述
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeReleaseResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReleaseResponseParams `json:"Response"`
}

func (r *DescribeReleaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReleaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRobotBizIDByAppKeyRequestParams struct {
	// 机器人appkey
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`
}

type DescribeRobotBizIDByAppKeyRequest struct {
	*tchttp.BaseRequest
	
	// 机器人appkey
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`
}

func (r *DescribeRobotBizIDByAppKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRobotBizIDByAppKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRobotBizIDByAppKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRobotBizIDByAppKeyResponseParams struct {
	// 机器人业务ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRobotBizIDByAppKeyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRobotBizIDByAppKeyResponseParams `json:"Response"`
}

func (r *DescribeRobotBizIDByAppKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRobotBizIDByAppKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStorageCredentialRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`
}

type DescribeStorageCredentialRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`
}

func (r *DescribeStorageCredentialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStorageCredentialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStorageCredentialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStorageCredentialResponseParams struct {
	// 密钥信息
	Credentials *Credentials `json:"Credentials,omitnil,omitempty" name:"Credentials"`

	// 失效时间
	ExpiredTime *int64 `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 对象存储桶
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// 对象存储可用区
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 目录
	FilePath *string `json:"FilePath,omitnil,omitempty" name:"FilePath"`

	// 存储类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 主号
	CorpUin *string `json:"CorpUin,omitnil,omitempty" name:"CorpUin"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeStorageCredentialResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStorageCredentialResponseParams `json:"Response"`
}

func (r *DescribeStorageCredentialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStorageCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUnsatisfiedReplyContextRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 回复ID
	ReplyBizId *string `json:"ReplyBizId,omitnil,omitempty" name:"ReplyBizId"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

type DescribeUnsatisfiedReplyContextRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 回复ID
	ReplyBizId *string `json:"ReplyBizId,omitnil,omitempty" name:"ReplyBizId"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

func (r *DescribeUnsatisfiedReplyContextRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUnsatisfiedReplyContextRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "ReplyBizId")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUnsatisfiedReplyContextRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUnsatisfiedReplyContextResponseParams struct {
	// 不满意回复上下文
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*Context `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUnsatisfiedReplyContextResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUnsatisfiedReplyContextResponseParams `json:"Response"`
}

func (r *DescribeUnsatisfiedReplyContextResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUnsatisfiedReplyContextResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EmbeddingObject struct {
	// 向量
	Embedding []*float64 `json:"Embedding,omitnil,omitempty" name:"Embedding"`
}

// Predefined struct for user
type ExportAttributeLabelRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 属性ID
	AttributeBizIds []*string `json:"AttributeBizIds,omitnil,omitempty" name:"AttributeBizIds"`

	// 根据筛选数据导出
	Filters *AttributeFilters `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type ExportAttributeLabelRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 属性ID
	AttributeBizIds []*string `json:"AttributeBizIds,omitnil,omitempty" name:"AttributeBizIds"`

	// 根据筛选数据导出
	Filters *AttributeFilters `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *ExportAttributeLabelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportAttributeLabelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	delete(f, "AttributeBizIds")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportAttributeLabelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportAttributeLabelResponseParams struct {
	// 导出任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ExportAttributeLabelResponse struct {
	*tchttp.BaseResponse
	Response *ExportAttributeLabelResponseParams `json:"Response"`
}

func (r *ExportAttributeLabelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportAttributeLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportQAListRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// QA业务ID
	QaBizIds []*string `json:"QaBizIds,omitnil,omitempty" name:"QaBizIds"`

	// 查询参数
	Filters *QAQuery `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type ExportQAListRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// QA业务ID
	QaBizIds []*string `json:"QaBizIds,omitnil,omitempty" name:"QaBizIds"`

	// 查询参数
	Filters *QAQuery `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *ExportQAListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportQAListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "QaBizIds")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportQAListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportQAListResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ExportQAListResponse struct {
	*tchttp.BaseResponse
	Response *ExportQAListResponseParams `json:"Response"`
}

func (r *ExportQAListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportQAListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportUnsatisfiedReplyRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 勾选导出ID列表
	ReplyBizIds []*string `json:"ReplyBizIds,omitnil,omitempty" name:"ReplyBizIds"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 检索过滤器
	Filters *Filters `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type ExportUnsatisfiedReplyRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 勾选导出ID列表
	ReplyBizIds []*string `json:"ReplyBizIds,omitnil,omitempty" name:"ReplyBizIds"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 检索过滤器
	Filters *Filters `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *ExportUnsatisfiedReplyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportUnsatisfiedReplyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "ReplyBizIds")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportUnsatisfiedReplyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportUnsatisfiedReplyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ExportUnsatisfiedReplyResponse struct {
	*tchttp.BaseResponse
	Response *ExportUnsatisfiedReplyResponseParams `json:"Response"`
}

func (r *ExportUnsatisfiedReplyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportUnsatisfiedReplyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filters struct {
	// 检索，用户问题或答案
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 错误类型检索
	Reasons []*string `json:"Reasons,omitnil,omitempty" name:"Reasons"`
}

// Predefined struct for user
type GenerateQARequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文档ID
	DocBizIds []*string `json:"DocBizIds,omitnil,omitempty" name:"DocBizIds"`
}

type GenerateQARequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文档ID
	DocBizIds []*string `json:"DocBizIds,omitnil,omitempty" name:"DocBizIds"`
}

func (r *GenerateQARequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateQARequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "DocBizIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GenerateQARequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GenerateQAResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GenerateQAResponse struct {
	*tchttp.BaseResponse
	Response *GenerateQAResponseParams `json:"Response"`
}

func (r *GenerateQAResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateQAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAppKnowledgeCountRequestParams struct {
	// 类型：doc-文档；qa-问答对
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 应用ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 登录用户主账号(集成商模式必填)	
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)	
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

type GetAppKnowledgeCountRequest struct {
	*tchttp.BaseRequest
	
	// 类型：doc-文档；qa-问答对
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 应用ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 登录用户主账号(集成商模式必填)	
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)	
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

func (r *GetAppKnowledgeCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAppKnowledgeCountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "AppBizId")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetAppKnowledgeCountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAppKnowledgeCountResponseParams struct {
	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetAppKnowledgeCountResponse struct {
	*tchttp.BaseResponse
	Response *GetAppKnowledgeCountResponseParams `json:"Response"`
}

func (r *GetAppKnowledgeCountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAppKnowledgeCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAppSecretRequestParams struct {
	// 应用ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`
}

type GetAppSecretRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`
}

func (r *GetAppSecretRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAppSecretRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetAppSecretRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAppSecretResponseParams struct {
	// 应用密钥
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 是否发布
	IsRelease *bool `json:"IsRelease,omitnil,omitempty" name:"IsRelease"`

	// 是否有查看权限
	HasPermission *bool `json:"HasPermission,omitnil,omitempty" name:"HasPermission"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetAppSecretResponse struct {
	*tchttp.BaseResponse
	Response *GetAppSecretResponseParams `json:"Response"`
}

func (r *GetAppSecretResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAppSecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDocPreviewRequestParams struct {
	// 文档业务ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`
}

type GetDocPreviewRequest struct {
	*tchttp.BaseRequest
	
	// 文档业务ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`
}

func (r *GetDocPreviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDocPreviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DocBizId")
	delete(f, "BotBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDocPreviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDocPreviewResponseParams struct {
	// 文件名
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文件类型
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// cos路径
	CosUrl *string `json:"CosUrl,omitnil,omitempty" name:"CosUrl"`

	// cos临时地址
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// cos桶
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetDocPreviewResponse struct {
	*tchttp.BaseResponse
	Response *GetDocPreviewResponseParams `json:"Response"`
}

func (r *GetDocPreviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDocPreviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetEmbeddingRequestParams struct {
	// 模型名称
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 需要 embedding 的文本, 单条文本最大长度500个字符, 总条数最大7条
	Inputs []*string `json:"Inputs,omitnil,omitempty" name:"Inputs"`

	// 是否在线, 后台异步任务使用离线, 实时任务使用在线, 默认值: false
	Online *bool `json:"Online,omitnil,omitempty" name:"Online"`
}

type GetEmbeddingRequest struct {
	*tchttp.BaseRequest
	
	// 模型名称
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 需要 embedding 的文本, 单条文本最大长度500个字符, 总条数最大7条
	Inputs []*string `json:"Inputs,omitnil,omitempty" name:"Inputs"`

	// 是否在线, 后台异步任务使用离线, 实时任务使用在线, 默认值: false
	Online *bool `json:"Online,omitnil,omitempty" name:"Online"`
}

func (r *GetEmbeddingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetEmbeddingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Model")
	delete(f, "Inputs")
	delete(f, "Online")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetEmbeddingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetEmbeddingResponseParams struct {
	// 特征
	Data []*EmbeddingObject `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetEmbeddingResponse struct {
	*tchttp.BaseResponse
	Response *GetEmbeddingResponseParams `json:"Response"`
}

func (r *GetEmbeddingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetEmbeddingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetMsgRecordRequestParams struct {
	// 类型
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 数量
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 会话sessionid
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 最后一条记录ID
	LastRecordId *string `json:"LastRecordId,omitnil,omitempty" name:"LastRecordId"`

	// 机器人AppKey
	BotAppKey *string `json:"BotAppKey,omitnil,omitempty" name:"BotAppKey"`
}

type GetMsgRecordRequest struct {
	*tchttp.BaseRequest
	
	// 类型
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 数量
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 会话sessionid
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 最后一条记录ID
	LastRecordId *string `json:"LastRecordId,omitnil,omitempty" name:"LastRecordId"`

	// 机器人AppKey
	BotAppKey *string `json:"BotAppKey,omitnil,omitempty" name:"BotAppKey"`
}

func (r *GetMsgRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetMsgRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "Count")
	delete(f, "SessionId")
	delete(f, "LastRecordId")
	delete(f, "BotAppKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetMsgRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetMsgRecordResponseParams struct {
	// 会话记录
	Records []*MsgRecord `json:"Records,omitnil,omitempty" name:"Records"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetMsgRecordResponse struct {
	*tchttp.BaseResponse
	Response *GetMsgRecordResponseParams `json:"Response"`
}

func (r *GetMsgRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetMsgRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTaskStatusRequestParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务类型
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`
}

type GetTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务类型
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`
}

func (r *GetTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTaskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "TaskType")
	delete(f, "BotBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTaskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTaskStatusResponseParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务类型
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 任务状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 任务参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Params *TaskParams `json:"Params,omitnil,omitempty" name:"Params"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *GetTaskStatusResponseParams `json:"Response"`
}

func (r *GetTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetWsTokenReq_Label struct {
	// 标签名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 标签值
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

// Predefined struct for user
type GetWsTokenRequestParams struct {
	// 接入类型
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 机器人AppKey
	BotAppKey *string `json:"BotAppKey,omitnil,omitempty" name:"BotAppKey"`

	// 坐席ID
	VisitorBizId *string `json:"VisitorBizId,omitnil,omitempty" name:"VisitorBizId"`

	// 坐席标签
	VisitorLabels []*GetWsTokenReq_Label `json:"VisitorLabels,omitnil,omitempty" name:"VisitorLabels"`
}

type GetWsTokenRequest struct {
	*tchttp.BaseRequest
	
	// 接入类型
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 机器人AppKey
	BotAppKey *string `json:"BotAppKey,omitnil,omitempty" name:"BotAppKey"`

	// 坐席ID
	VisitorBizId *string `json:"VisitorBizId,omitnil,omitempty" name:"VisitorBizId"`

	// 坐席标签
	VisitorLabels []*GetWsTokenReq_Label `json:"VisitorLabels,omitnil,omitempty" name:"VisitorLabels"`
}

func (r *GetWsTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetWsTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "BotAppKey")
	delete(f, "VisitorBizId")
	delete(f, "VisitorLabels")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetWsTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetWsTokenResponseParams struct {
	// token值
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetWsTokenResponse struct {
	*tchttp.BaseResponse
	Response *GetWsTokenResponseParams `json:"Response"`
}

func (r *GetWsTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetWsTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GroupQARequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// QA业务ID列表
	QaBizIds []*string `json:"QaBizIds,omitnil,omitempty" name:"QaBizIds"`

	// 分组 ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`
}

type GroupQARequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// QA业务ID列表
	QaBizIds []*string `json:"QaBizIds,omitnil,omitempty" name:"QaBizIds"`

	// 分组 ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`
}

func (r *GroupQARequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GroupQARequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "QaBizIds")
	delete(f, "CateBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GroupQARequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GroupQAResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GroupQAResponse struct {
	*tchttp.BaseResponse
	Response *GroupQAResponseParams `json:"Response"`
}

func (r *GroupQAResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GroupQAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Highlight struct {
	// 高亮启始位置
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartPos *string `json:"StartPos,omitnil,omitempty" name:"StartPos"`

	// 高亮结束位置
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndPos *string `json:"EndPos,omitnil,omitempty" name:"EndPos"`

	// 高亮子文本
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

// Predefined struct for user
type IgnoreUnsatisfiedReplyRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 不满意回复ID
	ReplyBizIds []*string `json:"ReplyBizIds,omitnil,omitempty" name:"ReplyBizIds"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

type IgnoreUnsatisfiedReplyRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 不满意回复ID
	ReplyBizIds []*string `json:"ReplyBizIds,omitnil,omitempty" name:"ReplyBizIds"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

func (r *IgnoreUnsatisfiedReplyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IgnoreUnsatisfiedReplyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "ReplyBizIds")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IgnoreUnsatisfiedReplyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IgnoreUnsatisfiedReplyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type IgnoreUnsatisfiedReplyResponse struct {
	*tchttp.BaseResponse
	Response *IgnoreUnsatisfiedReplyResponseParams `json:"Response"`
}

func (r *IgnoreUnsatisfiedReplyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IgnoreUnsatisfiedReplyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsTransferIntentRequestParams struct {
	// 内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 机器人appKey
	BotAppKey *string `json:"BotAppKey,omitnil,omitempty" name:"BotAppKey"`
}

type IsTransferIntentRequest struct {
	*tchttp.BaseRequest
	
	// 内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 机器人appKey
	BotAppKey *string `json:"BotAppKey,omitnil,omitempty" name:"BotAppKey"`
}

func (r *IsTransferIntentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsTransferIntentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Content")
	delete(f, "BotAppKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IsTransferIntentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsTransferIntentResponseParams struct {
	// 是否意图转人工
	Hit *bool `json:"Hit,omitnil,omitempty" name:"Hit"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type IsTransferIntentResponse struct {
	*tchttp.BaseResponse
	Response *IsTransferIntentResponseParams `json:"Response"`
}

func (r *IsTransferIntentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsTransferIntentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KnowledgeQaConfig struct {
	// 欢迎语，200字符以内
	// 注意：此字段可能返回 null，表示取不到有效值。
	Greeting *string `json:"Greeting,omitnil,omitempty" name:"Greeting"`

	// 角色描述，300字符以内
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleDescription *string `json:"RoleDescription,omitnil,omitempty" name:"RoleDescription"`

	// 模型配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Model *AppModel `json:"Model,omitnil,omitempty" name:"Model"`

	// 知识搜索配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Search []*KnowledgeQaSearch `json:"Search,omitnil,omitempty" name:"Search"`

	// 知识管理输出配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *KnowledgeQaOutput `json:"Output,omitnil,omitempty" name:"Output"`
}

type KnowledgeQaOutput struct {
	// 输出方式 1：流式 2：非流式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Method *uint64 `json:"Method,omitnil,omitempty" name:"Method"`

	// 通用模型回复
	// 注意：此字段可能返回 null，表示取不到有效值。
	UseGeneralKnowledge *bool `json:"UseGeneralKnowledge,omitnil,omitempty" name:"UseGeneralKnowledge"`

	// 未知回复语，300字符以内
	// 注意：此字段可能返回 null，表示取不到有效值。
	BareAnswer *string `json:"BareAnswer,omitnil,omitempty" name:"BareAnswer"`
}

type KnowledgeQaSearch struct {
	// 知识来源 doc：文档，qa：问答  taskflow：业务流程，search：搜索增强
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 问答-回复灵活度 1：已采纳答案直接回复 2：已采纳润色后回复
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReplyFlexibility *uint64 `json:"ReplyFlexibility,omitnil,omitempty" name:"ReplyFlexibility"`

	// 搜索增强-搜索引擎状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	UseSearchEngine *bool `json:"UseSearchEngine,omitnil,omitempty" name:"UseSearchEngine"`

	// 是否显示搜索引擎检索状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShowSearchEngine *bool `json:"ShowSearchEngine,omitnil,omitempty" name:"ShowSearchEngine"`

	// 知识来源，是否选择
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsEnabled *bool `json:"IsEnabled,omitnil,omitempty" name:"IsEnabled"`

	// 问答最大召回数量, 默认2，限制5
	// 注意：此字段可能返回 null，表示取不到有效值。
	QaTopN *uint64 `json:"QaTopN,omitnil,omitempty" name:"QaTopN"`

	// 文档最大召回数量, 默认3，限制5
	// 注意：此字段可能返回 null，表示取不到有效值。
	DocTopN *uint64 `json:"DocTopN,omitnil,omitempty" name:"DocTopN"`
}

type Label struct {
	// 标签ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelBizId *string `json:"LabelBizId,omitnil,omitempty" name:"LabelBizId"`

	// 标签名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelName *string `json:"LabelName,omitnil,omitempty" name:"LabelName"`
}

// Predefined struct for user
type ListAppCategoryRequestParams struct {

}

type ListAppCategoryRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ListAppCategoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAppCategoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAppCategoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAppCategoryResponseParams struct {
	// 应用类型列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*ListAppCategoryRspOption `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAppCategoryResponse struct {
	*tchttp.BaseResponse
	Response *ListAppCategoryResponseParams `json:"Response"`
}

func (r *ListAppCategoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAppCategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAppCategoryRspOption struct {
	// 类型名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 类型值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 类型log
	// 注意：此字段可能返回 null，表示取不到有效值。
	Logo *string `json:"Logo,omitnil,omitempty" name:"Logo"`
}

// Predefined struct for user
type ListAppRequestParams struct {
	// 应用类型；knowledge_qa-知识问答管理；summary-知识摘要；classifys-知识标签提取
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`

	// 每页数目，整型
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页码，整型
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 关键词：应用/修改人
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 登录用户子账号(集成商模式必填)	
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

type ListAppRequest struct {
	*tchttp.BaseRequest
	
	// 应用类型；knowledge_qa-知识问答管理；summary-知识摘要；classifys-知识标签提取
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`

	// 每页数目，整型
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页码，整型
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 关键词：应用/修改人
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 登录用户子账号(集成商模式必填)	
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

func (r *ListAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppType")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "Keyword")
	delete(f, "LoginSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAppResponseParams struct {
	// 数量
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// 标签列表
	List []*AppInfo `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAppResponse struct {
	*tchttp.BaseResponse
	Response *ListAppResponseParams `json:"Response"`
}

func (r *ListAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAttributeLabelRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 查询内容
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`
}

type ListAttributeLabelRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 查询内容
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`
}

func (r *ListAttributeLabelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAttributeLabelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	delete(f, "Query")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAttributeLabelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAttributeLabelResponseParams struct {
	// 总数
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// 列表
	List []*AttrLabelDetail `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAttributeLabelResponse struct {
	*tchttp.BaseResponse
	Response *ListAttributeLabelResponseParams `json:"Response"`
}

func (r *ListAttributeLabelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAttributeLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListDocItem struct {
	// 文档ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// 文件名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文件类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// cos路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosUrl *string `json:"CosUrl,omitnil,omitempty" name:"CosUrl"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 文档状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 文档状态描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// 原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 答案中是否引用
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsRefer *bool `json:"IsRefer,omitnil,omitempty" name:"IsRefer"`

	// 问答对数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	QaNum *int64 `json:"QaNum,omitnil,omitempty" name:"QaNum"`

	// 是否已删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDeleted *bool `json:"IsDeleted,omitnil,omitempty" name:"IsDeleted"`

	// 文档来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *int64 `json:"Source,omitnil,omitempty" name:"Source"`

	// 文档来源描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceDesc *string `json:"SourceDesc,omitnil,omitempty" name:"SourceDesc"`

	// 是否允许重新生成
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAllowRestart *bool `json:"IsAllowRestart,omitnil,omitempty" name:"IsAllowRestart"`

	// qa是否已删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDeletedQa *bool `json:"IsDeletedQa,omitnil,omitempty" name:"IsDeletedQa"`

	// 问答是否生成中
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsCreatingQa *bool `json:"IsCreatingQa,omitnil,omitempty" name:"IsCreatingQa"`

	// 是否允许删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAllowDelete *bool `json:"IsAllowDelete,omitnil,omitempty" name:"IsAllowDelete"`

	// 是否允许操作引用开关
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAllowRefer *bool `json:"IsAllowRefer,omitnil,omitempty" name:"IsAllowRefer"`

	// 问答是否生成过
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsCreatedQa *bool `json:"IsCreatedQa,omitnil,omitempty" name:"IsCreatedQa"`

	// 文档字符量
	// 注意：此字段可能返回 null，表示取不到有效值。
	DocCharSize *string `json:"DocCharSize,omitnil,omitempty" name:"DocCharSize"`

	// 属性标签适用范围
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttrRange *uint64 `json:"AttrRange,omitnil,omitempty" name:"AttrRange"`

	// 属性标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttrLabels []*AttrLabel `json:"AttrLabels,omitnil,omitempty" name:"AttrLabels"`

	// 是否允许编辑
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAllowEdit *bool `json:"IsAllowEdit,omitnil,omitempty" name:"IsAllowEdit"`

	// 外部引用链接类型 0：系统链接 1：自定义链接
	// 值为1时，WebUrl 字段不能为空，否则不生效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReferUrlType *uint64 `json:"ReferUrlType,omitnil,omitempty" name:"ReferUrlType"`

	// 网页(或自定义链接)地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebUrl *string `json:"WebUrl,omitnil,omitempty" name:"WebUrl"`

	// 有效开始时间，unix时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireStart *string `json:"ExpireStart,omitnil,omitempty" name:"ExpireStart"`

	// 有效结束时间，unix时间戳，0代表永久有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireEnd *string `json:"ExpireEnd,omitnil,omitempty" name:"ExpireEnd"`
}

// Predefined struct for user
type ListDocRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询内容
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 文档状态： 7 审核中、8 审核失败、10 待发布、11 发布中、12 已发布、13 学习中、14 学习失败 20 已过期
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type ListDocRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询内容
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 文档状态： 7 审核中、8 审核失败、10 待发布、11 发布中、12 已发布、13 学习中、14 学习失败 20 已过期
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *ListDocRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDocRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Query")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListDocRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDocResponseParams struct {
	// 文档数量
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// 文档列表
	List []*ListDocItem `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListDocResponse struct {
	*tchttp.BaseResponse
	Response *ListDocResponseParams `json:"Response"`
}

func (r *ListDocResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDocResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListModelRequestParams struct {
	// 应用类型；knowledge_qa-知识问答管理；summary-知识摘要；classifys-知识标签提取
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`

	// 登录用户主账号(集成商模式必填)	
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)	
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

type ListModelRequest struct {
	*tchttp.BaseRequest
	
	// 应用类型；knowledge_qa-知识问答管理；summary-知识摘要；classifys-知识标签提取
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`

	// 登录用户主账号(集成商模式必填)	
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)	
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

func (r *ListModelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListModelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppType")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListModelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListModelResponseParams struct {
	// 模型列表
	List []*ModelInfo `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListModelResponse struct {
	*tchttp.BaseResponse
	Response *ListModelResponseParams `json:"Response"`
}

func (r *ListModelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListModelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListQACateRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`
}

type ListQACateRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`
}

func (r *ListQACateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListQACateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListQACateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListQACateResponseParams struct {
	// 列表
	List []*QACate `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListQACateResponse struct {
	*tchttp.BaseResponse
	Response *ListQACateResponseParams `json:"Response"`
}

func (r *ListQACateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListQACateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListQARequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询问题
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 校验状态(1未校验2采纳3不采纳)
	AcceptStatus []*int64 `json:"AcceptStatus,omitnil,omitempty" name:"AcceptStatus"`

	// 发布状态(2待发布 3发布中 4已发布 7审核中 8审核失败 9人工申述中 11人工申述失败)
	ReleaseStatus []*int64 `json:"ReleaseStatus,omitnil,omitempty" name:"ReleaseStatus"`

	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// 来源(1 文档生成 2 批量导入 3 手动添加)
	Source *int64 `json:"Source,omitnil,omitempty" name:"Source"`

	// 查询答案
	QueryAnswer *string `json:"QueryAnswer,omitnil,omitempty" name:"QueryAnswer"`

	// QA业务ID列表
	QaBizIds []*string `json:"QaBizIds,omitnil,omitempty" name:"QaBizIds"`
}

type ListQARequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询问题
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 校验状态(1未校验2采纳3不采纳)
	AcceptStatus []*int64 `json:"AcceptStatus,omitnil,omitempty" name:"AcceptStatus"`

	// 发布状态(2待发布 3发布中 4已发布 7审核中 8审核失败 9人工申述中 11人工申述失败)
	ReleaseStatus []*int64 `json:"ReleaseStatus,omitnil,omitempty" name:"ReleaseStatus"`

	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// 来源(1 文档生成 2 批量导入 3 手动添加)
	Source *int64 `json:"Source,omitnil,omitempty" name:"Source"`

	// 查询答案
	QueryAnswer *string `json:"QueryAnswer,omitnil,omitempty" name:"QueryAnswer"`

	// QA业务ID列表
	QaBizIds []*string `json:"QaBizIds,omitnil,omitempty" name:"QaBizIds"`
}

func (r *ListQARequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListQARequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Query")
	delete(f, "AcceptStatus")
	delete(f, "ReleaseStatus")
	delete(f, "DocBizId")
	delete(f, "Source")
	delete(f, "QueryAnswer")
	delete(f, "QaBizIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListQARequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListQAResponseParams struct {
	// 问答数量
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// 待校验问答数量
	WaitVerifyTotal *string `json:"WaitVerifyTotal,omitnil,omitempty" name:"WaitVerifyTotal"`

	// 未采纳问答数量
	NotAcceptedTotal *string `json:"NotAcceptedTotal,omitnil,omitempty" name:"NotAcceptedTotal"`

	// 已采纳问答数量
	AcceptedTotal *string `json:"AcceptedTotal,omitnil,omitempty" name:"AcceptedTotal"`

	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 问答详情
	List []*ListQaItem `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListQAResponse struct {
	*tchttp.BaseResponse
	Response *ListQAResponseParams `json:"Response"`
}

func (r *ListQAResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListQAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListQaItem struct {
	// 问答ID
	QaBizId *string `json:"QaBizId,omitnil,omitempty" name:"QaBizId"`

	// 问题
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 答案
	Answer *string `json:"Answer,omitnil,omitempty" name:"Answer"`

	// 来源
	Source *int64 `json:"Source,omitnil,omitempty" name:"Source"`

	// 来源描述
	SourceDesc *string `json:"SourceDesc,omitnil,omitempty" name:"SourceDesc"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 状态描述
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 是否允许编辑
	IsAllowEdit *bool `json:"IsAllowEdit,omitnil,omitempty" name:"IsAllowEdit"`

	// 是否允许删除
	IsAllowDelete *bool `json:"IsAllowDelete,omitnil,omitempty" name:"IsAllowDelete"`

	// 是否允许校验
	IsAllowAccept *bool `json:"IsAllowAccept,omitnil,omitempty" name:"IsAllowAccept"`

	// 文档名称
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文档类型
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 问答字符数
	QaCharSize *string `json:"QaCharSize,omitnil,omitempty" name:"QaCharSize"`
}

// Predefined struct for user
type ListRejectedQuestionPreviewRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询内容
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 发布单ID
	ReleaseBizId *string `json:"ReleaseBizId,omitnil,omitempty" name:"ReleaseBizId"`

	// 状态(1新增2更新3删除)
	Actions []*uint64 `json:"Actions,omitnil,omitempty" name:"Actions"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type ListRejectedQuestionPreviewRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询内容
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 发布单ID
	ReleaseBizId *string `json:"ReleaseBizId,omitnil,omitempty" name:"ReleaseBizId"`

	// 状态(1新增2更新3删除)
	Actions []*uint64 `json:"Actions,omitnil,omitempty" name:"Actions"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *ListRejectedQuestionPreviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRejectedQuestionPreviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Query")
	delete(f, "ReleaseBizId")
	delete(f, "Actions")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListRejectedQuestionPreviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRejectedQuestionPreviewResponseParams struct {
	// 文档数量
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// 文档列表
	List []*ReleaseRejectedQuestion `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListRejectedQuestionPreviewResponse struct {
	*tchttp.BaseResponse
	Response *ListRejectedQuestionPreviewResponseParams `json:"Response"`
}

func (r *ListRejectedQuestionPreviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRejectedQuestionPreviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRejectedQuestionRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询内容
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`
}

type ListRejectedQuestionRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询内容
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`
}

func (r *ListRejectedQuestionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRejectedQuestionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Query")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListRejectedQuestionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRejectedQuestionResponseParams struct {
	// 总数
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// 拒答问题列表
	List []*RejectedQuestion `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListRejectedQuestionResponse struct {
	*tchttp.BaseResponse
	Response *ListRejectedQuestionResponseParams `json:"Response"`
}

func (r *ListRejectedQuestionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRejectedQuestionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListReleaseConfigPreviewRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询内容
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 发布单ID
	ReleaseBizId *string `json:"ReleaseBizId,omitnil,omitempty" name:"ReleaseBizId"`

	// 状态(1新增2更新3删除)
	Actions []*uint64 `json:"Actions,omitnil,omitempty" name:"Actions"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 发布状态
	ReleaseStatus []*uint64 `json:"ReleaseStatus,omitnil,omitempty" name:"ReleaseStatus"`
}

type ListReleaseConfigPreviewRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询内容
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 发布单ID
	ReleaseBizId *string `json:"ReleaseBizId,omitnil,omitempty" name:"ReleaseBizId"`

	// 状态(1新增2更新3删除)
	Actions []*uint64 `json:"Actions,omitnil,omitempty" name:"Actions"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 发布状态
	ReleaseStatus []*uint64 `json:"ReleaseStatus,omitnil,omitempty" name:"ReleaseStatus"`
}

func (r *ListReleaseConfigPreviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListReleaseConfigPreviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Query")
	delete(f, "ReleaseBizId")
	delete(f, "Actions")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ReleaseStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListReleaseConfigPreviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListReleaseConfigPreviewResponseParams struct {
	// 数量
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// 配置项列表
	List []*ReleaseConfigs `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListReleaseConfigPreviewResponse struct {
	*tchttp.BaseResponse
	Response *ListReleaseConfigPreviewResponseParams `json:"Response"`
}

func (r *ListReleaseConfigPreviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListReleaseConfigPreviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListReleaseDocPreviewRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询内容
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 发布业务ID
	ReleaseBizId *string `json:"ReleaseBizId,omitnil,omitempty" name:"ReleaseBizId"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 状态(1新增2修改3删除)
	Actions []*uint64 `json:"Actions,omitnil,omitempty" name:"Actions"`
}

type ListReleaseDocPreviewRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询内容
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 发布业务ID
	ReleaseBizId *string `json:"ReleaseBizId,omitnil,omitempty" name:"ReleaseBizId"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 状态(1新增2修改3删除)
	Actions []*uint64 `json:"Actions,omitnil,omitempty" name:"Actions"`
}

func (r *ListReleaseDocPreviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListReleaseDocPreviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Query")
	delete(f, "ReleaseBizId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Actions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListReleaseDocPreviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListReleaseDocPreviewResponseParams struct {
	// 文档数量
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// 文档列表
	List []*ReleaseDoc `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListReleaseDocPreviewResponse struct {
	*tchttp.BaseResponse
	Response *ListReleaseDocPreviewResponseParams `json:"Response"`
}

func (r *ListReleaseDocPreviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListReleaseDocPreviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListReleaseItem struct {
	// 版本ID
	ReleaseBizId *string `json:"ReleaseBizId,omitnil,omitempty" name:"ReleaseBizId"`

	// 发布人
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 发布描述
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 发布状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 发布状态描述
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// 失败原因
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 发布成功数
	SuccessCount *int64 `json:"SuccessCount,omitnil,omitempty" name:"SuccessCount"`

	// 发布失败数
	FailCount *int64 `json:"FailCount,omitnil,omitempty" name:"FailCount"`
}

// Predefined struct for user
type ListReleaseQAPreviewRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询内容
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 发布单ID
	ReleaseBizId *string `json:"ReleaseBizId,omitnil,omitempty" name:"ReleaseBizId"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 状态(1新增2修改3删除)
	Actions []*uint64 `json:"Actions,omitnil,omitempty" name:"Actions"`

	// 发布状态(4发布成功5发布失败)
	ReleaseStatus []*uint64 `json:"ReleaseStatus,omitnil,omitempty" name:"ReleaseStatus"`
}

type ListReleaseQAPreviewRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询内容
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 发布单ID
	ReleaseBizId *string `json:"ReleaseBizId,omitnil,omitempty" name:"ReleaseBizId"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 状态(1新增2修改3删除)
	Actions []*uint64 `json:"Actions,omitnil,omitempty" name:"Actions"`

	// 发布状态(4发布成功5发布失败)
	ReleaseStatus []*uint64 `json:"ReleaseStatus,omitnil,omitempty" name:"ReleaseStatus"`
}

func (r *ListReleaseQAPreviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListReleaseQAPreviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Query")
	delete(f, "ReleaseBizId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Actions")
	delete(f, "ReleaseStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListReleaseQAPreviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListReleaseQAPreviewResponseParams struct {
	// 文档数量
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// 文档列表
	List []*ReleaseQA `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListReleaseQAPreviewResponse struct {
	*tchttp.BaseResponse
	Response *ListReleaseQAPreviewResponseParams `json:"Response"`
}

func (r *ListReleaseQAPreviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListReleaseQAPreviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListReleaseRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type ListReleaseRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *ListReleaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListReleaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListReleaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListReleaseResponseParams struct {
	// 发布列表数量
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// 发布列表
	List []*ListReleaseItem `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListReleaseResponse struct {
	*tchttp.BaseResponse
	Response *ListReleaseResponseParams `json:"Response"`
}

func (r *ListReleaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListReleaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListSelectDocRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文档名称
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文档状态： 7 审核中、8 审核失败、10 待发布、11 发布中、12 已发布、13 学习中、14 学习失败 20 已过期
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type ListSelectDocRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文档名称
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文档状态： 7 审核中、8 审核失败、10 待发布、11 发布中、12 已发布、13 学习中、14 学习失败 20 已过期
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *ListSelectDocRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSelectDocRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "FileName")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListSelectDocRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListSelectDocResponseParams struct {
	// 下拉框内容
	List []*Option `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListSelectDocResponse struct {
	*tchttp.BaseResponse
	Response *ListSelectDocResponseParams `json:"Response"`
}

func (r *ListSelectDocResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSelectDocResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUnsatisfiedReplyRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 用户请求(问题或答案)
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 错误类型检索
	Reasons []*string `json:"Reasons,omitnil,omitempty" name:"Reasons"`
}

type ListUnsatisfiedReplyRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 用户请求(问题或答案)
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 错误类型检索
	Reasons []*string `json:"Reasons,omitnil,omitempty" name:"Reasons"`
}

func (r *ListUnsatisfiedReplyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUnsatisfiedReplyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	delete(f, "Query")
	delete(f, "Reasons")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListUnsatisfiedReplyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUnsatisfiedReplyResponseParams struct {
	// 总数
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// 不满意回复列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*UnsatisfiedReply `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListUnsatisfiedReplyResponse struct {
	*tchttp.BaseResponse
	Response *ListUnsatisfiedReplyResponseParams `json:"Response"`
}

func (r *ListUnsatisfiedReplyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUnsatisfiedReplyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Message struct {
	// role表示角色  user标识用户提问，assistant标识返回的答案
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 对话内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type ModelInfo struct {
	// 模型名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 模型描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelDesc *string `json:"ModelDesc,omitnil,omitempty" name:"ModelDesc"`

	// 模型名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AliasName *string `json:"AliasName,omitnil,omitempty" name:"AliasName"`
}

// Predefined struct for user
type ModifyAppRequestParams struct {
	// 应用 ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 应用类型；knowledge_qa-知识问答管理；summary-知识摘要；classifys-知识标签提取
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`

	// 应用基础配置
	BaseConfig *BaseConfig `json:"BaseConfig,omitnil,omitempty" name:"BaseConfig"`

	// 登录用户子账号(集成商模式必填)	
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 应用配置
	AppConfig *AppConfig `json:"AppConfig,omitnil,omitempty" name:"AppConfig"`
}

type ModifyAppRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 应用类型；knowledge_qa-知识问答管理；summary-知识摘要；classifys-知识标签提取
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`

	// 应用基础配置
	BaseConfig *BaseConfig `json:"BaseConfig,omitnil,omitempty" name:"BaseConfig"`

	// 登录用户子账号(集成商模式必填)	
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 应用配置
	AppConfig *AppConfig `json:"AppConfig,omitnil,omitempty" name:"AppConfig"`
}

func (r *ModifyAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppBizId")
	delete(f, "AppType")
	delete(f, "BaseConfig")
	delete(f, "LoginSubAccountUin")
	delete(f, "AppConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAppResponseParams struct {
	// 应用App
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAppResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAppResponseParams `json:"Response"`
}

func (r *ModifyAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAttributeLabelRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 属性ID
	AttributeBizId *string `json:"AttributeBizId,omitnil,omitempty" name:"AttributeBizId"`

	// 属性标识
	AttrKey *string `json:"AttrKey,omitnil,omitempty" name:"AttrKey"`

	// 属性名称
	AttrName *string `json:"AttrName,omitnil,omitempty" name:"AttrName"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 删除的属性标签
	DeleteLabelBizIds []*string `json:"DeleteLabelBizIds,omitnil,omitempty" name:"DeleteLabelBizIds"`

	// 新增或编辑的属性标签
	Labels []*AttributeLabel `json:"Labels,omitnil,omitempty" name:"Labels"`
}

type ModifyAttributeLabelRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 属性ID
	AttributeBizId *string `json:"AttributeBizId,omitnil,omitempty" name:"AttributeBizId"`

	// 属性标识
	AttrKey *string `json:"AttrKey,omitnil,omitempty" name:"AttrKey"`

	// 属性名称
	AttrName *string `json:"AttrName,omitnil,omitempty" name:"AttrName"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 删除的属性标签
	DeleteLabelBizIds []*string `json:"DeleteLabelBizIds,omitnil,omitempty" name:"DeleteLabelBizIds"`

	// 新增或编辑的属性标签
	Labels []*AttributeLabel `json:"Labels,omitnil,omitempty" name:"Labels"`
}

func (r *ModifyAttributeLabelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAttributeLabelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "AttributeBizId")
	delete(f, "AttrKey")
	delete(f, "AttrName")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	delete(f, "DeleteLabelBizIds")
	delete(f, "Labels")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAttributeLabelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAttributeLabelResponseParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAttributeLabelResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAttributeLabelResponseParams `json:"Response"`
}

func (r *ModifyAttributeLabelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAttributeLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDocAttrRangeRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文档ID
	DocBizIds []*string `json:"DocBizIds,omitnil,omitempty" name:"DocBizIds"`

	// 属性标签适用范围 1：全部，2：按条件
	AttrRange *uint64 `json:"AttrRange,omitnil,omitempty" name:"AttrRange"`

	// 属性标签引用
	AttrLabels []*AttrLabelRefer `json:"AttrLabels,omitnil,omitempty" name:"AttrLabels"`
}

type ModifyDocAttrRangeRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文档ID
	DocBizIds []*string `json:"DocBizIds,omitnil,omitempty" name:"DocBizIds"`

	// 属性标签适用范围 1：全部，2：按条件
	AttrRange *uint64 `json:"AttrRange,omitnil,omitempty" name:"AttrRange"`

	// 属性标签引用
	AttrLabels []*AttrLabelRefer `json:"AttrLabels,omitnil,omitempty" name:"AttrLabels"`
}

func (r *ModifyDocAttrRangeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDocAttrRangeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "DocBizIds")
	delete(f, "AttrRange")
	delete(f, "AttrLabels")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDocAttrRangeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDocAttrRangeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDocAttrRangeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDocAttrRangeResponseParams `json:"Response"`
}

func (r *ModifyDocAttrRangeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDocAttrRangeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDocRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// 是否引用链接
	IsRefer *bool `json:"IsRefer,omitnil,omitempty" name:"IsRefer"`

	// 属性标签适用范围 1：全部，2：按条件
	AttrRange *uint64 `json:"AttrRange,omitnil,omitempty" name:"AttrRange"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 适用范围，关联的属性标签
	AttrLabels []*AttrLabelRefer `json:"AttrLabels,omitnil,omitempty" name:"AttrLabels"`

	// 网页(或自定义链接)地址
	WebUrl *string `json:"WebUrl,omitnil,omitempty" name:"WebUrl"`

	// 外部引用链接类型 0：系统链接 1：自定义链接
	// 值为1时，WebUrl 字段不能为空，否则不生效。
	ReferUrlType *uint64 `json:"ReferUrlType,omitnil,omitempty" name:"ReferUrlType"`

	// 有效开始时间，unix时间戳
	ExpireStart *string `json:"ExpireStart,omitnil,omitempty" name:"ExpireStart"`

	// 有效结束时间，unix时间戳，0代表永久有效
	ExpireEnd *string `json:"ExpireEnd,omitnil,omitempty" name:"ExpireEnd"`
}

type ModifyDocRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// 是否引用链接
	IsRefer *bool `json:"IsRefer,omitnil,omitempty" name:"IsRefer"`

	// 属性标签适用范围 1：全部，2：按条件
	AttrRange *uint64 `json:"AttrRange,omitnil,omitempty" name:"AttrRange"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 适用范围，关联的属性标签
	AttrLabels []*AttrLabelRefer `json:"AttrLabels,omitnil,omitempty" name:"AttrLabels"`

	// 网页(或自定义链接)地址
	WebUrl *string `json:"WebUrl,omitnil,omitempty" name:"WebUrl"`

	// 外部引用链接类型 0：系统链接 1：自定义链接
	// 值为1时，WebUrl 字段不能为空，否则不生效。
	ReferUrlType *uint64 `json:"ReferUrlType,omitnil,omitempty" name:"ReferUrlType"`

	// 有效开始时间，unix时间戳
	ExpireStart *string `json:"ExpireStart,omitnil,omitempty" name:"ExpireStart"`

	// 有效结束时间，unix时间戳，0代表永久有效
	ExpireEnd *string `json:"ExpireEnd,omitnil,omitempty" name:"ExpireEnd"`
}

func (r *ModifyDocRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDocRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "DocBizId")
	delete(f, "IsRefer")
	delete(f, "AttrRange")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	delete(f, "AttrLabels")
	delete(f, "WebUrl")
	delete(f, "ReferUrlType")
	delete(f, "ExpireStart")
	delete(f, "ExpireEnd")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDocRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDocResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDocResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDocResponseParams `json:"Response"`
}

func (r *ModifyDocResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDocResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyQAAttrRangeRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 问答ID
	QaBizIds []*string `json:"QaBizIds,omitnil,omitempty" name:"QaBizIds"`

	// 属性标签适用范围 1：全部，2：按条件
	AttrRange *uint64 `json:"AttrRange,omitnil,omitempty" name:"AttrRange"`

	// 属性标签引用
	AttrLabels []*AttrLabelRefer `json:"AttrLabels,omitnil,omitempty" name:"AttrLabels"`
}

type ModifyQAAttrRangeRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 问答ID
	QaBizIds []*string `json:"QaBizIds,omitnil,omitempty" name:"QaBizIds"`

	// 属性标签适用范围 1：全部，2：按条件
	AttrRange *uint64 `json:"AttrRange,omitnil,omitempty" name:"AttrRange"`

	// 属性标签引用
	AttrLabels []*AttrLabelRefer `json:"AttrLabels,omitnil,omitempty" name:"AttrLabels"`
}

func (r *ModifyQAAttrRangeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyQAAttrRangeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "QaBizIds")
	delete(f, "AttrRange")
	delete(f, "AttrLabels")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyQAAttrRangeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyQAAttrRangeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyQAAttrRangeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyQAAttrRangeResponseParams `json:"Response"`
}

func (r *ModifyQAAttrRangeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyQAAttrRangeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyQACateRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 分类名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分类业务ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`
}

type ModifyQACateRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 分类名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分类业务ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`
}

func (r *ModifyQACateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyQACateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "Name")
	delete(f, "CateBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyQACateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyQACateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyQACateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyQACateResponseParams `json:"Response"`
}

func (r *ModifyQACateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyQACateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyQARequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 问答ID
	QaBizId *string `json:"QaBizId,omitnil,omitempty" name:"QaBizId"`

	// 问题
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 答案
	Answer *string `json:"Answer,omitnil,omitempty" name:"Answer"`

	// 自定义参数
	CustomParam *string `json:"CustomParam,omitnil,omitempty" name:"CustomParam"`

	// 属性标签适用范围 1：全部，2：按条件
	AttrRange *uint64 `json:"AttrRange,omitnil,omitempty" name:"AttrRange"`

	// 属性标签引用
	AttrLabels []*AttrLabelRefer `json:"AttrLabels,omitnil,omitempty" name:"AttrLabels"`

	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// 分类ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`

	// 有效开始时间，unix时间戳
	ExpireStart *string `json:"ExpireStart,omitnil,omitempty" name:"ExpireStart"`

	// 有效结束时间，unix时间戳，0代表永久有效
	ExpireEnd *string `json:"ExpireEnd,omitnil,omitempty" name:"ExpireEnd"`
}

type ModifyQARequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 问答ID
	QaBizId *string `json:"QaBizId,omitnil,omitempty" name:"QaBizId"`

	// 问题
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 答案
	Answer *string `json:"Answer,omitnil,omitempty" name:"Answer"`

	// 自定义参数
	CustomParam *string `json:"CustomParam,omitnil,omitempty" name:"CustomParam"`

	// 属性标签适用范围 1：全部，2：按条件
	AttrRange *uint64 `json:"AttrRange,omitnil,omitempty" name:"AttrRange"`

	// 属性标签引用
	AttrLabels []*AttrLabelRefer `json:"AttrLabels,omitnil,omitempty" name:"AttrLabels"`

	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// 分类ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`

	// 有效开始时间，unix时间戳
	ExpireStart *string `json:"ExpireStart,omitnil,omitempty" name:"ExpireStart"`

	// 有效结束时间，unix时间戳，0代表永久有效
	ExpireEnd *string `json:"ExpireEnd,omitnil,omitempty" name:"ExpireEnd"`
}

func (r *ModifyQARequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyQARequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "QaBizId")
	delete(f, "Question")
	delete(f, "Answer")
	delete(f, "CustomParam")
	delete(f, "AttrRange")
	delete(f, "AttrLabels")
	delete(f, "DocBizId")
	delete(f, "CateBizId")
	delete(f, "ExpireStart")
	delete(f, "ExpireEnd")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyQARequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyQAResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyQAResponse struct {
	*tchttp.BaseResponse
	Response *ModifyQAResponseParams `json:"Response"`
}

func (r *ModifyQAResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyQAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRejectedQuestionRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 拒答问题
	// 
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 拒答问题来源的数据源唯一id
	// 
	// 
	RejectedBizId *string `json:"RejectedBizId,omitnil,omitempty" name:"RejectedBizId"`
}

type ModifyRejectedQuestionRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 拒答问题
	// 
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 拒答问题来源的数据源唯一id
	// 
	// 
	RejectedBizId *string `json:"RejectedBizId,omitnil,omitempty" name:"RejectedBizId"`
}

func (r *ModifyRejectedQuestionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRejectedQuestionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "Question")
	delete(f, "RejectedBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRejectedQuestionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRejectedQuestionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRejectedQuestionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRejectedQuestionResponseParams `json:"Response"`
}

func (r *ModifyRejectedQuestionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRejectedQuestionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MsgRecord struct {
	// 内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 记录ID
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 关联记录ID
	RelatedRecordId *string `json:"RelatedRecordId,omitnil,omitempty" name:"RelatedRecordId"`

	// 是否来自自己
	IsFromSelf *bool `json:"IsFromSelf,omitnil,omitempty" name:"IsFromSelf"`

	// 发送者名称
	FromName *string `json:"FromName,omitnil,omitempty" name:"FromName"`

	// 发送者头像
	FromAvatar *string `json:"FromAvatar,omitnil,omitempty" name:"FromAvatar"`

	// 时间戳
	Timestamp *string `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 是否已读
	HasRead *bool `json:"HasRead,omitnil,omitempty" name:"HasRead"`

	// 评价
	Score *uint64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 是否评分
	CanRating *bool `json:"CanRating,omitnil,omitempty" name:"CanRating"`

	// 记录类型
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 引用来源
	References []*MsgRecordReference `json:"References,omitnil,omitempty" name:"References"`

	// 评价原因
	Reasons []*string `json:"Reasons,omitnil,omitempty" name:"Reasons"`

	// 是否大模型
	IsLlmGenerated *bool `json:"IsLlmGenerated,omitnil,omitempty" name:"IsLlmGenerated"`
}

type MsgRecordReference struct {
	// id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 链接
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 类型
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 来源文档ID
	DocId *string `json:"DocId,omitnil,omitempty" name:"DocId"`
}

type Option struct {
	// 文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 文件字符数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CharSize *string `json:"CharSize,omitnil,omitempty" name:"CharSize"`

	// 文件类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`
}

// Predefined struct for user
type ParseDocRequestParams struct {
	// 文件名称(需要包括文件后缀, 最大长度1024字节)
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 文件下载链接 (支持的文件类型: docx, txt, markdown, pdf)
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 任务ID, 用于幂等去重, 业务自行定义(最大长度64字节)
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 切分策略
	Policy *string `json:"Policy,omitnil,omitempty" name:"Policy"`

	// 默认值: split
	Operate *string `json:"Operate,omitnil,omitempty" name:"Operate"`
}

type ParseDocRequest struct {
	*tchttp.BaseRequest
	
	// 文件名称(需要包括文件后缀, 最大长度1024字节)
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 文件下载链接 (支持的文件类型: docx, txt, markdown, pdf)
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 任务ID, 用于幂等去重, 业务自行定义(最大长度64字节)
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 切分策略
	Policy *string `json:"Policy,omitnil,omitempty" name:"Policy"`

	// 默认值: split
	Operate *string `json:"Operate,omitnil,omitempty" name:"Operate"`
}

func (r *ParseDocRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ParseDocRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Url")
	delete(f, "TaskId")
	delete(f, "Policy")
	delete(f, "Operate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ParseDocRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ParseDocResponseParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ParseDocResponse struct {
	*tchttp.BaseResponse
	Response *ParseDocResponseParams `json:"Response"`
}

func (r *ParseDocResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ParseDocResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QACate struct {
	// QA分类的业务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`

	// 分类名称
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分类下QA数量
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 是否可新增
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	CanAdd *bool `json:"CanAdd,omitnil,omitempty" name:"CanAdd"`

	// 是否可编辑
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	CanEdit *bool `json:"CanEdit,omitnil,omitempty" name:"CanEdit"`

	// 是否可删除
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	CanDelete *bool `json:"CanDelete,omitnil,omitempty" name:"CanDelete"`

	// 子分类
	// 注意：此字段可能返回 null，表示取不到有效值。
	Children []*QACate `json:"Children,omitnil,omitempty" name:"Children"`
}

type QAList struct {
	// 问答ID
	QaBizId *string `json:"QaBizId,omitnil,omitempty" name:"QaBizId"`

	// 是否采纳
	IsAccepted *bool `json:"IsAccepted,omitnil,omitempty" name:"IsAccepted"`

	// 分类ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`

	// 问题
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 答案
	Answer *string `json:"Answer,omitnil,omitempty" name:"Answer"`
}

type QAQuery struct {
	// 页码
	// 
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 查询内容
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 分类ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`

	// 校验状态
	AcceptStatus []*uint64 `json:"AcceptStatus,omitnil,omitempty" name:"AcceptStatus"`

	// 发布状态
	ReleaseStatus []*uint64 `json:"ReleaseStatus,omitnil,omitempty" name:"ReleaseStatus"`

	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// QAID
	QaBizId *string `json:"QaBizId,omitnil,omitempty" name:"QaBizId"`

	// 来源
	Source *uint64 `json:"Source,omitnil,omitempty" name:"Source"`

	// 查询答案
	QueryAnswer *string `json:"QueryAnswer,omitnil,omitempty" name:"QueryAnswer"`
}

// Predefined struct for user
type QueryParseDocResultRequestParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type QueryParseDocResultRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *QueryParseDocResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryParseDocResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryParseDocResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryParseDocResultResponseParams struct {
	// 等待 / 执行中 / 成功 / 失败
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 解析后的文件内容
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 文件下载地址
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 解析失败原因
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryParseDocResultResponse struct {
	*tchttp.BaseResponse
	Response *QueryParseDocResultResponseParams `json:"Response"`
}

func (r *QueryParseDocResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryParseDocResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryRewriteRequestParams struct {
	// 需要改写的问题
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 需要改写的多轮历史会话
	Messages []*Message `json:"Messages,omitnil,omitempty" name:"Messages"`

	// 模型名称
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`
}

type QueryRewriteRequest struct {
	*tchttp.BaseRequest
	
	// 需要改写的问题
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 需要改写的多轮历史会话
	Messages []*Message `json:"Messages,omitnil,omitempty" name:"Messages"`

	// 模型名称
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`
}

func (r *QueryRewriteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryRewriteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Question")
	delete(f, "Messages")
	delete(f, "Model")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryRewriteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryRewriteResponseParams struct {
	// 改写结果
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryRewriteResponse struct {
	*tchttp.BaseResponse
	Response *QueryRewriteResponseParams `json:"Response"`
}

func (r *QueryRewriteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryRewriteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RateMsgRecordRequestParams struct {
	// 机器人appKey
	BotAppKey *string `json:"BotAppKey,omitnil,omitempty" name:"BotAppKey"`

	// 消息ID
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 1点赞2点踩
	Score *uint64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 原因
	Reasons []*string `json:"Reasons,omitnil,omitempty" name:"Reasons"`
}

type RateMsgRecordRequest struct {
	*tchttp.BaseRequest
	
	// 机器人appKey
	BotAppKey *string `json:"BotAppKey,omitnil,omitempty" name:"BotAppKey"`

	// 消息ID
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 1点赞2点踩
	Score *uint64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 原因
	Reasons []*string `json:"Reasons,omitnil,omitempty" name:"Reasons"`
}

func (r *RateMsgRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RateMsgRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotAppKey")
	delete(f, "RecordId")
	delete(f, "Score")
	delete(f, "Reasons")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RateMsgRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RateMsgRecordResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RateMsgRecordResponse struct {
	*tchttp.BaseResponse
	Response *RateMsgRecordResponseParams `json:"Response"`
}

func (r *RateMsgRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RateMsgRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReferDetail struct {
	// 引用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReferBizId *string `json:"ReferBizId,omitnil,omitempty" name:"ReferBizId"`

	// 文档类型 (1 QA, 2 文档段)
	// 注意：此字段可能返回 null，表示取不到有效值。
	DocType *uint64 `json:"DocType,omitnil,omitempty" name:"DocType"`

	// 文档名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DocName *string `json:"DocName,omitnil,omitempty" name:"DocName"`

	// 分片内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageContent *string `json:"PageContent,omitnil,omitempty" name:"PageContent"`

	// 问题
	// 注意：此字段可能返回 null，表示取不到有效值。
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 答案
	// 注意：此字段可能返回 null，表示取不到有效值。
	Answer *string `json:"Answer,omitnil,omitempty" name:"Answer"`

	// 置信度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// 标记
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mark *uint64 `json:"Mark,omitnil,omitempty" name:"Mark"`

	// 分片高亮内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Highlights []*Highlight `json:"Highlights,omitnil,omitempty" name:"Highlights"`

	// 原始内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrgData *string `json:"OrgData,omitnil,omitempty" name:"OrgData"`
}

type RejectedQuestion struct {
	// 拒答问题ID
	// 
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	RejectedBizId *string `json:"RejectedBizId,omitnil,omitempty" name:"RejectedBizId"`

	// 被拒答的问题
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 状态描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// 更新时间
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 是否允许编辑
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAllowEdit *bool `json:"IsAllowEdit,omitnil,omitempty" name:"IsAllowEdit"`

	// 是否允许删除
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAllowDelete *bool `json:"IsAllowDelete,omitnil,omitempty" name:"IsAllowDelete"`
}

type ReleaseConfigs struct {
	// 配置项描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigItem *string `json:"ConfigItem,omitnil,omitempty" name:"ConfigItem"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Action *uint64 `json:"Action,omitnil,omitempty" name:"Action"`

	// 变更后的内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 变更前的内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastValue *string `json:"LastValue,omitnil,omitempty" name:"LastValue"`

	// 变更内容(优先级展示content内容,content为空取value内容)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

type ReleaseDoc struct {
	// 文件名
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文件类型
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 状态
	Action *uint64 `json:"Action,omitnil,omitempty" name:"Action"`

	// 状态描述
	ActionDesc *string `json:"ActionDesc,omitnil,omitempty" name:"ActionDesc"`

	// 失败原因
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 文档业务ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`
}

type ReleaseQA struct {
	// 问题
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 状态
	Action *uint64 `json:"Action,omitnil,omitempty" name:"Action"`

	// 状态描述
	ActionDesc *string `json:"ActionDesc,omitnil,omitempty" name:"ActionDesc"`

	// 来源
	Source *uint64 `json:"Source,omitnil,omitempty" name:"Source"`

	// 来源描述
	SourceDesc *string `json:"SourceDesc,omitnil,omitempty" name:"SourceDesc"`

	// 文件名字
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文档类型
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 失败原因
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 发布状态
	ReleaseStatus *uint64 `json:"ReleaseStatus,omitnil,omitempty" name:"ReleaseStatus"`

	// QAID
	QaBizId *string `json:"QaBizId,omitnil,omitempty" name:"QaBizId"`

	// 文档业务ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`
}

type ReleaseRejectedQuestion struct {
	// 问题
	// 注意：此字段可能返回 null，表示取不到有效值。
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Action *uint64 `json:"Action,omitnil,omitempty" name:"Action"`

	// 状态描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActionDesc *string `json:"ActionDesc,omitnil,omitempty" name:"ActionDesc"`

	// 失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

// Predefined struct for user
type ResetSessionRequestParams struct {
	// 会话ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

type ResetSessionRequest struct {
	*tchttp.BaseRequest
	
	// 会话ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

func (r *ResetSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetSessionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResetSessionResponse struct {
	*tchttp.BaseResponse
	Response *ResetSessionResponseParams `json:"Response"`
}

func (r *ResetSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RetryDocAuditRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`
}

type RetryDocAuditRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`
}

func (r *RetryDocAuditRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RetryDocAuditRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "DocBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RetryDocAuditRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RetryDocAuditResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RetryDocAuditResponse struct {
	*tchttp.BaseResponse
	Response *RetryDocAuditResponseParams `json:"Response"`
}

func (r *RetryDocAuditResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RetryDocAuditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RetryDocParseRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`
}

type RetryDocParseRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`
}

func (r *RetryDocParseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RetryDocParseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "DocBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RetryDocParseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RetryDocParseResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RetryDocParseResponse struct {
	*tchttp.BaseResponse
	Response *RetryDocParseResponseParams `json:"Response"`
}

func (r *RetryDocParseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RetryDocParseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RetryReleaseRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 发布业务ID
	ReleaseBizId *string `json:"ReleaseBizId,omitnil,omitempty" name:"ReleaseBizId"`
}

type RetryReleaseRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 发布业务ID
	ReleaseBizId *string `json:"ReleaseBizId,omitnil,omitempty" name:"ReleaseBizId"`
}

func (r *RetryReleaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RetryReleaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "ReleaseBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RetryReleaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RetryReleaseResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RetryReleaseResponse struct {
	*tchttp.BaseResponse
	Response *RetryReleaseResponseParams `json:"Response"`
}

func (r *RetryReleaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RetryReleaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SaveDocRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文件名
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文件类型(md|txt|docx|pdf|xlsx)
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// cos路径
	CosUrl *string `json:"CosUrl,omitnil,omitempty" name:"CosUrl"`

	// ETag 全称为 Entity Tag，是对象被创建时标识对象内容的信息标签，可用于检查对象的内容是否发生变化
	ETag *string `json:"ETag,omitnil,omitempty" name:"ETag"`

	// cos_hash x-cos-hash-crc64ecma 头部中的 CRC64编码进行校验上传到云端的文件和本地文件的一致性
	CosHash *string `json:"CosHash,omitnil,omitempty" name:"CosHash"`

	// 文件大小
	Size *string `json:"Size,omitnil,omitempty" name:"Size"`

	// 属性标签适用范围 1：全部，2：按条件范围
	AttrRange *uint64 `json:"AttrRange,omitnil,omitempty" name:"AttrRange"`

	// 来源(0 源文件导入 1 网页导入)
	Source *int64 `json:"Source,omitnil,omitempty" name:"Source"`

	// 网页(或自定义链接)地址
	WebUrl *string `json:"WebUrl,omitnil,omitempty" name:"WebUrl"`

	// 属性标签引用
	AttrLabels []*AttrLabelRefer `json:"AttrLabels,omitnil,omitempty" name:"AttrLabels"`

	// 外部引用链接类型 0：系统链接 1：自定义链接
	// 值为1时，WebUrl 字段不能为空，否则不生效。
	ReferUrlType *uint64 `json:"ReferUrlType,omitnil,omitempty" name:"ReferUrlType"`

	// 有效开始时间，unix时间戳
	ExpireStart *string `json:"ExpireStart,omitnil,omitempty" name:"ExpireStart"`

	// 有效结束时间，unix时间戳，0代表永久有效
	ExpireEnd *string `json:"ExpireEnd,omitnil,omitempty" name:"ExpireEnd"`

	// 是否引用链接
	IsRefer *bool `json:"IsRefer,omitnil,omitempty" name:"IsRefer"`

	// 文档操作类型：1：批量导入；2:文档导入
	Opt *uint64 `json:"Opt,omitnil,omitempty" name:"Opt"`
}

type SaveDocRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文件名
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文件类型(md|txt|docx|pdf|xlsx)
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// cos路径
	CosUrl *string `json:"CosUrl,omitnil,omitempty" name:"CosUrl"`

	// ETag 全称为 Entity Tag，是对象被创建时标识对象内容的信息标签，可用于检查对象的内容是否发生变化
	ETag *string `json:"ETag,omitnil,omitempty" name:"ETag"`

	// cos_hash x-cos-hash-crc64ecma 头部中的 CRC64编码进行校验上传到云端的文件和本地文件的一致性
	CosHash *string `json:"CosHash,omitnil,omitempty" name:"CosHash"`

	// 文件大小
	Size *string `json:"Size,omitnil,omitempty" name:"Size"`

	// 属性标签适用范围 1：全部，2：按条件范围
	AttrRange *uint64 `json:"AttrRange,omitnil,omitempty" name:"AttrRange"`

	// 来源(0 源文件导入 1 网页导入)
	Source *int64 `json:"Source,omitnil,omitempty" name:"Source"`

	// 网页(或自定义链接)地址
	WebUrl *string `json:"WebUrl,omitnil,omitempty" name:"WebUrl"`

	// 属性标签引用
	AttrLabels []*AttrLabelRefer `json:"AttrLabels,omitnil,omitempty" name:"AttrLabels"`

	// 外部引用链接类型 0：系统链接 1：自定义链接
	// 值为1时，WebUrl 字段不能为空，否则不生效。
	ReferUrlType *uint64 `json:"ReferUrlType,omitnil,omitempty" name:"ReferUrlType"`

	// 有效开始时间，unix时间戳
	ExpireStart *string `json:"ExpireStart,omitnil,omitempty" name:"ExpireStart"`

	// 有效结束时间，unix时间戳，0代表永久有效
	ExpireEnd *string `json:"ExpireEnd,omitnil,omitempty" name:"ExpireEnd"`

	// 是否引用链接
	IsRefer *bool `json:"IsRefer,omitnil,omitempty" name:"IsRefer"`

	// 文档操作类型：1：批量导入；2:文档导入
	Opt *uint64 `json:"Opt,omitnil,omitempty" name:"Opt"`
}

func (r *SaveDocRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SaveDocRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "FileName")
	delete(f, "FileType")
	delete(f, "CosUrl")
	delete(f, "ETag")
	delete(f, "CosHash")
	delete(f, "Size")
	delete(f, "AttrRange")
	delete(f, "Source")
	delete(f, "WebUrl")
	delete(f, "AttrLabels")
	delete(f, "ReferUrlType")
	delete(f, "ExpireStart")
	delete(f, "ExpireEnd")
	delete(f, "IsRefer")
	delete(f, "Opt")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SaveDocRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SaveDocResponseParams struct {
	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// 导入错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 错误链接
	ErrorLink *string `json:"ErrorLink,omitnil,omitempty" name:"ErrorLink"`

	// 错误链接文本
	ErrorLinkText *string `json:"ErrorLinkText,omitnil,omitempty" name:"ErrorLinkText"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SaveDocResponse struct {
	*tchttp.BaseResponse
	Response *SaveDocResponseParams `json:"Response"`
}

func (r *SaveDocResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SaveDocResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopDocParseRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`
}

type StopDocParseRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`
}

func (r *StopDocParseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopDocParseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "DocBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopDocParseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopDocParseResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopDocParseResponse struct {
	*tchttp.BaseResponse
	Response *StopDocParseResponseParams `json:"Response"`
}

func (r *StopDocParseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopDocParseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SummaryConfig struct {
	// 模型配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Model *AppModel `json:"Model,omitnil,omitempty" name:"Model"`

	// 知识摘要输出配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *SummaryOutput `json:"Output,omitnil,omitempty" name:"Output"`
}

type SummaryOutput struct {
	// 输出方式 1：流式 2：非流式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Method *uint64 `json:"Method,omitnil,omitempty" name:"Method"`

	// 输出要求 1：文本总结 2：自定义要求
	// 注意：此字段可能返回 null，表示取不到有效值。
	Requirement *uint64 `json:"Requirement,omitnil,omitempty" name:"Requirement"`

	// 自定义要求指令
	// 注意：此字段可能返回 null，表示取不到有效值。
	RequireCommand *string `json:"RequireCommand,omitnil,omitempty" name:"RequireCommand"`
}

type TaskParams struct {
	// 下载地址,需要通过cos桶临时秘钥去下载
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosPath *string `json:"CosPath,omitnil,omitempty" name:"CosPath"`
}

type UnsatisfiedReply struct {
	// 不满意回复ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReplyBizId *string `json:"ReplyBizId,omitnil,omitempty" name:"ReplyBizId"`

	// 消息记录ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordBizId *string `json:"RecordBizId,omitnil,omitempty" name:"RecordBizId"`

	// 用户问题
	// 注意：此字段可能返回 null，表示取不到有效值。
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 机器人回复
	// 注意：此字段可能返回 null，表示取不到有效值。
	Answer *string `json:"Answer,omitnil,omitempty" name:"Answer"`

	// 错误类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reasons []*string `json:"Reasons,omitnil,omitempty" name:"Reasons"`
}

// Predefined struct for user
type UploadAttributeLabelRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文件名
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// cos路径
	CosUrl *string `json:"CosUrl,omitnil,omitempty" name:"CosUrl"`

	// x-cos-hash-crc64ecma 头部中的 CRC64编码进行校验上传到云端的文件和本地文件的一致性
	CosHash *string `json:"CosHash,omitnil,omitempty" name:"CosHash"`

	// 文件大小
	Size *string `json:"Size,omitnil,omitempty" name:"Size"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

type UploadAttributeLabelRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文件名
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// cos路径
	CosUrl *string `json:"CosUrl,omitnil,omitempty" name:"CosUrl"`

	// x-cos-hash-crc64ecma 头部中的 CRC64编码进行校验上传到云端的文件和本地文件的一致性
	CosHash *string `json:"CosHash,omitnil,omitempty" name:"CosHash"`

	// 文件大小
	Size *string `json:"Size,omitnil,omitempty" name:"Size"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

func (r *UploadAttributeLabelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadAttributeLabelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "FileName")
	delete(f, "CosUrl")
	delete(f, "CosHash")
	delete(f, "Size")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadAttributeLabelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadAttributeLabelResponseParams struct {
	// 导入错误
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 错误链接
	ErrorLink *string `json:"ErrorLink,omitnil,omitempty" name:"ErrorLink"`

	// 错误链接文本
	ErrorLinkText *string `json:"ErrorLinkText,omitnil,omitempty" name:"ErrorLinkText"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UploadAttributeLabelResponse struct {
	*tchttp.BaseResponse
	Response *UploadAttributeLabelResponseParams `json:"Response"`
}

func (r *UploadAttributeLabelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadAttributeLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyQARequestParams struct {
	// 问答列表
	List []*QAList `json:"List,omitnil,omitempty" name:"List"`

	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

type VerifyQARequest struct {
	*tchttp.BaseRequest
	
	// 问答列表
	List []*QAList `json:"List,omitnil,omitempty" name:"List"`

	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

func (r *VerifyQARequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyQARequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "List")
	delete(f, "BotBizId")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifyQARequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyQAResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type VerifyQAResponse struct {
	*tchttp.BaseResponse
	Response *VerifyQAResponseParams `json:"Response"`
}

func (r *VerifyQAResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyQAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}