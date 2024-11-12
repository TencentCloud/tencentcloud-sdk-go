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

	// token余量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TokenBalance *float64 `json:"TokenBalance,omitnil,omitempty" name:"TokenBalance"`

	// 是否使用上下文指代轮次
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsUseContext *bool `json:"IsUseContext,omitnil,omitempty" name:"IsUseContext"`

	// 上下文记忆轮数
	// 注意：此字段可能返回 null，表示取不到有效值。
	HistoryLimit *uint64 `json:"HistoryLimit,omitnil,omitempty" name:"HistoryLimit"`

	// 使用类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsageType *string `json:"UsageType,omitnil,omitempty" name:"UsageType"`
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

	// 应用头像url，在CreateApp和ModifyApp中作为入参必填。
	// 作为入参传入说明：
	// 1. 传入的url图片限制为jpeg和png，大小限制为500KB，url链接需允许head请求。
	// 2. 如果用户没有对象存储，可使用“获取文件上传临时密钥”(DescribeStorageCredential)接口，获取cos临时密钥和上传路径，自行上传头像至cos中并获取访问链接。
	Avatar *string `json:"Avatar,omitnil,omitempty" name:"Avatar"`

	// 应用描述
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`
}

// Predefined struct for user
type CheckAttributeLabelExistRequestParams struct {
	// 应用ID
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
	
	// 应用ID
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
	// 应用ID
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
	
	// 应用ID
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

	// 欢迎语，200字符以内
	// 注意：此字段可能返回 null，表示取不到有效值。
	Greeting *string `json:"Greeting,omitnil,omitempty" name:"Greeting"`
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

	// 文档信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileInfos []*MsgFileInfo `json:"FileInfos,omitnil,omitempty" name:"FileInfos"`

	// 回复方式，15：澄清确认回复
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReplyMethod *uint64 `json:"ReplyMethod,omitnil,omitempty" name:"ReplyMethod"`
}

// Predefined struct for user
type ConvertDocumentRequestParams struct {
	// 图片的 Url 地址。 支持的图片格式：PNG、JPG、JPEG、PDF，暂不支持 GIF 格式。 支持的图片大小：所下载图片经 Base64 编码后不超过 8M。图片下载时间不超过 3 秒。 支持的图片像素：单边介于20-10000px之间。 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// 图片的 Base64 值。 支持的图片格式：PNG、JPG、JPEG、PDF，暂不支持 GIF 格式。 支持的图片大小：所下载图片经Base64编码后不超过 8M。图片下载时间不超过 3 秒。 支持的图片像素：单边介于20-10000px之间。 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	FileBase64 *string `json:"FileBase64,omitnil,omitempty" name:"FileBase64"`

	// 当传入文件是PDF类型（FileType=PDF）时，用来指定pdf识别的起始页码，识别的页码包含当前值。
	FileStartPageNumber *int64 `json:"FileStartPageNumber,omitnil,omitempty" name:"FileStartPageNumber"`

	// 当传入文件是PDF类型（FileType=PDF）时，用来指定pdf识别的结束页码，识别的页码包含当前值。
	// 建议一次请求的页面不超过3页。
	FileEndPageNumber *int64 `json:"FileEndPageNumber,omitnil,omitempty" name:"FileEndPageNumber"`
}

type ConvertDocumentRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Url 地址。 支持的图片格式：PNG、JPG、JPEG、PDF，暂不支持 GIF 格式。 支持的图片大小：所下载图片经 Base64 编码后不超过 8M。图片下载时间不超过 3 秒。 支持的图片像素：单边介于20-10000px之间。 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// 图片的 Base64 值。 支持的图片格式：PNG、JPG、JPEG、PDF，暂不支持 GIF 格式。 支持的图片大小：所下载图片经Base64编码后不超过 8M。图片下载时间不超过 3 秒。 支持的图片像素：单边介于20-10000px之间。 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	FileBase64 *string `json:"FileBase64,omitnil,omitempty" name:"FileBase64"`

	// 当传入文件是PDF类型（FileType=PDF）时，用来指定pdf识别的起始页码，识别的页码包含当前值。
	FileStartPageNumber *int64 `json:"FileStartPageNumber,omitnil,omitempty" name:"FileStartPageNumber"`

	// 当传入文件是PDF类型（FileType=PDF）时，用来指定pdf识别的结束页码，识别的页码包含当前值。
	// 建议一次请求的页面不超过3页。
	FileEndPageNumber *int64 `json:"FileEndPageNumber,omitnil,omitempty" name:"FileEndPageNumber"`
}

func (r *ConvertDocumentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ConvertDocumentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileUrl")
	delete(f, "FileBase64")
	delete(f, "FileStartPageNumber")
	delete(f, "FileEndPageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ConvertDocumentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ConvertDocumentResponseParams struct {
	// 识别生成的word文件base64编码的字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	WordRecognizeInfo []*WordRecognizeInfo `json:"WordRecognizeInfo,omitnil,omitempty" name:"WordRecognizeInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ConvertDocumentResponse struct {
	*tchttp.BaseResponse
	Response *ConvertDocumentResponseParams `json:"Response"`
}

func (r *ConvertDocumentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ConvertDocumentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Coord struct {
	// 横坐标
	X *int64 `json:"X,omitnil,omitempty" name:"X"`

	// 纵坐标
	Y *int64 `json:"Y,omitnil,omitempty" name:"Y"`
}

// Predefined struct for user
type CreateAppRequestParams struct {
	// 应用类型；knowledge_qa-知识问答管理
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`

	// 应用基础配置
	BaseConfig *BaseConfig `json:"BaseConfig,omitnil,omitempty" name:"BaseConfig"`
}

type CreateAppRequest struct {
	*tchttp.BaseRequest
	
	// 应用类型；knowledge_qa-知识问答管理
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
	// 应用ID
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
	
	// 应用ID
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
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 父级业务ID
	ParentBizId *string `json:"ParentBizId,omitnil,omitempty" name:"ParentBizId"`

	// 分类名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type CreateQACateRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
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
	// 应用ID
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

	// 相似问内容
	SimilarQuestions []*string `json:"SimilarQuestions,omitnil,omitempty" name:"SimilarQuestions"`
}

type CreateQARequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
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

	// 相似问内容
	SimilarQuestions []*string `json:"SimilarQuestions,omitnil,omitempty" name:"SimilarQuestions"`
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
	delete(f, "SimilarQuestions")
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

type CreateReconstructDocumentFlowConfig struct {
	// Markdown文件中表格返回的形式
	// 0，表格以MD形式返回
	// 1，表格以HTML形式返回
	// 默认为1
	TableResultType *string `json:"TableResultType,omitnil,omitempty" name:"TableResultType"`

	// 智能文档解析返回结果的格式
	// 0：只返回全文MD；
	// 1：只返回每一页的OCR原始Json；
	// 2：只返回每一页的MD，
	// 3：返回全文MD + 每一页的OCR原始Json；
	// 4：返回全文MD + 每一页的MD，
	// 默认值为3（返回全文MD + 每一页的OCR原始Json）
	// 
	ResultType *string `json:"ResultType,omitnil,omitempty" name:"ResultType"`
}

// Predefined struct for user
type CreateReconstructDocumentFlowRequestParams struct {
	// 文件类型。支持的文件类型：PDF、DOC、DOCX、PPT、PPTX、MD、TXT、XLS、XLSX、CSV、PNG、JPG、JPEG、BMP、GIF、WEBP、HEIC、EPS、ICNS、IM、PCX、PPM、TIFF、XBM、HEIF、JP2
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 文件的 Base64 值。支持的文件大小：所下载文件经Base64编码后不超过 8M。文件下载时间不超过 3 秒。支持的图片像素：单边介于20-10000px之间。文件的 FileUrl、FileBase64 必须提供一个，如果都提供，只使用 FileUrl。
	FileBase64 *string `json:"FileBase64,omitnil,omitempty" name:"FileBase64"`

	// 文件的 Url 地址。支持的文件大小：所下载文件经 Base64 编码后不超过 100M。文件下载时间不超过 15 秒。支持的图片像素：单边介于20-10000px之间。 文件存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议文件存储于腾讯云。非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// 当传入文件是PDF类型时，用来指定pdf识别的起始页码，识别的页码包含当前值。默认为1，表示从pdf文件的第1页开始识别。
	FileStartPageNumber *int64 `json:"FileStartPageNumber,omitnil,omitempty" name:"FileStartPageNumber"`

	// 当传入文件是PDF类型时，用来指定pdf识别的结束页码，识别的页码包含当前值。默认为100，表示识别到pdf文件的第100页。单次调用最多支持识别100页内容，即FileEndPageNumber-FileStartPageNumber需要不大于100。
	FileEndPageNumber *int64 `json:"FileEndPageNumber,omitnil,omitempty" name:"FileEndPageNumber"`

	// 创建文档解析任务配置信息
	Config *CreateReconstructDocumentFlowConfig `json:"Config,omitnil,omitempty" name:"Config"`
}

type CreateReconstructDocumentFlowRequest struct {
	*tchttp.BaseRequest
	
	// 文件类型。支持的文件类型：PDF、DOC、DOCX、PPT、PPTX、MD、TXT、XLS、XLSX、CSV、PNG、JPG、JPEG、BMP、GIF、WEBP、HEIC、EPS、ICNS、IM、PCX、PPM、TIFF、XBM、HEIF、JP2
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 文件的 Base64 值。支持的文件大小：所下载文件经Base64编码后不超过 8M。文件下载时间不超过 3 秒。支持的图片像素：单边介于20-10000px之间。文件的 FileUrl、FileBase64 必须提供一个，如果都提供，只使用 FileUrl。
	FileBase64 *string `json:"FileBase64,omitnil,omitempty" name:"FileBase64"`

	// 文件的 Url 地址。支持的文件大小：所下载文件经 Base64 编码后不超过 100M。文件下载时间不超过 15 秒。支持的图片像素：单边介于20-10000px之间。 文件存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议文件存储于腾讯云。非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// 当传入文件是PDF类型时，用来指定pdf识别的起始页码，识别的页码包含当前值。默认为1，表示从pdf文件的第1页开始识别。
	FileStartPageNumber *int64 `json:"FileStartPageNumber,omitnil,omitempty" name:"FileStartPageNumber"`

	// 当传入文件是PDF类型时，用来指定pdf识别的结束页码，识别的页码包含当前值。默认为100，表示识别到pdf文件的第100页。单次调用最多支持识别100页内容，即FileEndPageNumber-FileStartPageNumber需要不大于100。
	FileEndPageNumber *int64 `json:"FileEndPageNumber,omitnil,omitempty" name:"FileEndPageNumber"`

	// 创建文档解析任务配置信息
	Config *CreateReconstructDocumentFlowConfig `json:"Config,omitnil,omitempty" name:"Config"`
}

func (r *CreateReconstructDocumentFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReconstructDocumentFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileType")
	delete(f, "FileBase64")
	delete(f, "FileUrl")
	delete(f, "FileStartPageNumber")
	delete(f, "FileEndPageNumber")
	delete(f, "Config")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateReconstructDocumentFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReconstructDocumentFlowResponseParams struct {
	// 任务唯一id。30天内可以通过GetReconstructDocumentResult接口查询TaskId对应的处理结果。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateReconstructDocumentFlowResponse struct {
	*tchttp.BaseResponse
	Response *CreateReconstructDocumentFlowResponseParams `json:"Response"`
}

func (r *CreateReconstructDocumentFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReconstructDocumentFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRejectedQuestionRequestParams struct {
	// 应用ID
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
	
	// 应用ID
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
	// 应用ID
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
	
	// 应用ID
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

	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`
}

type DeleteDocRequest struct {
	*tchttp.BaseRequest
	
	// 文档业务ID列表
	DocBizIds []*string `json:"DocBizIds,omitnil,omitempty" name:"DocBizIds"`

	// 应用ID
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
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 分类业务ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`
}

type DeleteQACateRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
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
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 问答ID
	QaBizIds []*string `json:"QaBizIds,omitnil,omitempty" name:"QaBizIds"`
}

type DeleteQARequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
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
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 拒答问题来源的数据源唯一id
	// 
	// 
	RejectedBizIds []*string `json:"RejectedBizIds,omitnil,omitempty" name:"RejectedBizIds"`
}

type DeleteRejectedQuestionRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
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

	// 应用状态，1：未上线，2：运行中，3：停用
	AppStatus *uint64 `json:"AppStatus,omitnil,omitempty" name:"AppStatus"`

	// 状态说明
	AppStatusDesc *string `json:"AppStatusDesc,omitnil,omitempty" name:"AppStatusDesc"`

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
	// 应用ID
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
	
	// 应用ID
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
type DescribeCallStatsGraphRequestParams struct {
	// uin
	UinAccount []*string `json:"UinAccount,omitnil,omitempty" name:"UinAccount"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 子业务类型
	SubBizType *string `json:"SubBizType,omitnil,omitempty" name:"SubBizType"`

	// 模型标识
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 应用id列表
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`
}

type DescribeCallStatsGraphRequest struct {
	*tchttp.BaseRequest
	
	// uin
	UinAccount []*string `json:"UinAccount,omitnil,omitempty" name:"UinAccount"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 子业务类型
	SubBizType *string `json:"SubBizType,omitnil,omitempty" name:"SubBizType"`

	// 模型标识
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 应用id列表
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`
}

func (r *DescribeCallStatsGraphRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCallStatsGraphRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UinAccount")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	delete(f, "SubBizType")
	delete(f, "ModelName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "AppBizIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCallStatsGraphRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCallStatsGraphResponseParams struct {
	// 统计信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*Stat `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCallStatsGraphResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCallStatsGraphResponseParams `json:"Response"`
}

func (r *DescribeCallStatsGraphResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCallStatsGraphResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConcurrencyUsageGraphRequestParams struct {
	// 模型标识
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// uin
	UinAccount []*string `json:"UinAccount,omitnil,omitempty" name:"UinAccount"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 子业务类型
	SubBizType *string `json:"SubBizType,omitnil,omitempty" name:"SubBizType"`

	// 应用id列表
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`
}

type DescribeConcurrencyUsageGraphRequest struct {
	*tchttp.BaseRequest
	
	// 模型标识
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// uin
	UinAccount []*string `json:"UinAccount,omitnil,omitempty" name:"UinAccount"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 子业务类型
	SubBizType *string `json:"SubBizType,omitnil,omitempty" name:"SubBizType"`

	// 应用id列表
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`
}

func (r *DescribeConcurrencyUsageGraphRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConcurrencyUsageGraphRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "UinAccount")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	delete(f, "SubBizType")
	delete(f, "AppBizIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConcurrencyUsageGraphRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConcurrencyUsageGraphResponseParams struct {
	// 统计信息
	X []*string `json:"X,omitnil,omitempty" name:"X"`

	// 可用并发y轴坐标
	AvailableY []*int64 `json:"AvailableY,omitnil,omitempty" name:"AvailableY"`

	// 成功调用并发y轴坐标
	SuccessCallY []*int64 `json:"SuccessCallY,omitnil,omitempty" name:"SuccessCallY"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConcurrencyUsageGraphResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConcurrencyUsageGraphResponseParams `json:"Response"`
}

func (r *DescribeConcurrencyUsageGraphResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConcurrencyUsageGraphResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConcurrencyUsageRequestParams struct {
	// 模型标识
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 应用id列表
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`
}

type DescribeConcurrencyUsageRequest struct {
	*tchttp.BaseRequest
	
	// 模型标识
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 应用id列表
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`
}

func (r *DescribeConcurrencyUsageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConcurrencyUsageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "AppBizIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConcurrencyUsageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConcurrencyUsageResponseParams struct {
	// 可用并发数
	AvailableConcurrency *uint64 `json:"AvailableConcurrency,omitnil,omitempty" name:"AvailableConcurrency"`

	// 并发峰值
	ConcurrencyPeak *uint64 `json:"ConcurrencyPeak,omitnil,omitempty" name:"ConcurrencyPeak"`

	// 调用超可用次数
	ExceedUsageTime *uint64 `json:"ExceedUsageTime,omitnil,omitempty" name:"ExceedUsageTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConcurrencyUsageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConcurrencyUsageResponseParams `json:"Response"`
}

func (r *DescribeConcurrencyUsageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConcurrencyUsageResponse) FromJsonString(s string) error {
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
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`
}

type DescribeDocRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
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

	// 文档状态： 1-未生成 2-生成中 3-生成成功 4-生成失败 5-删除中 6-删除成功 7-审核中 8-审核失败 9-审核成功 10-待发布 11-发布中 12-已发布 13-学习中 14-学习失败 15-更新中 16-更新失败 17-解析中 18-解析失败 19-导入失败 20-已过期 21-超量失效 22-超量失效恢复
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
type DescribeKnowledgeUsagePieGraphRequestParams struct {
	// 应用ID数组
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`
}

type DescribeKnowledgeUsagePieGraphRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID数组
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`
}

func (r *DescribeKnowledgeUsagePieGraphRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKnowledgeUsagePieGraphRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppBizIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKnowledgeUsagePieGraphRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKnowledgeUsagePieGraphResponseParams struct {
	// 所有应用已用的字符总数
	AvailableCharSize *string `json:"AvailableCharSize,omitnil,omitempty" name:"AvailableCharSize"`

	// 应用饼图详情列表
	List []*KnowledgeCapacityPieGraphDetail `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeKnowledgeUsagePieGraphResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKnowledgeUsagePieGraphResponseParams `json:"Response"`
}

func (r *DescribeKnowledgeUsagePieGraphResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKnowledgeUsagePieGraphResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKnowledgeUsageRequestParams struct {

}

type DescribeKnowledgeUsageRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeKnowledgeUsageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKnowledgeUsageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKnowledgeUsageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKnowledgeUsageResponseParams struct {
	// 可用字符数
	AvailableCharSize *string `json:"AvailableCharSize,omitnil,omitempty" name:"AvailableCharSize"`

	// 超量字符数
	ExceedCharSize *string `json:"ExceedCharSize,omitnil,omitempty" name:"ExceedCharSize"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeKnowledgeUsageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKnowledgeUsageResponseParams `json:"Response"`
}

func (r *DescribeKnowledgeUsageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKnowledgeUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQARequestParams struct {
	// QA业务ID
	QaBizId *string `json:"QaBizId,omitnil,omitempty" name:"QaBizId"`

	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`
}

type DescribeQARequest struct {
	*tchttp.BaseRequest
	
	// QA业务ID
	QaBizId *string `json:"QaBizId,omitnil,omitempty" name:"QaBizId"`

	// 应用ID
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

	// 相似问列表信息
	SimilarQuestions []*SimilarQuestion `json:"SimilarQuestions,omitnil,omitempty" name:"SimilarQuestions"`

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
	// 应用ID
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
	
	// 应用ID
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
	// 应用appkey
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`
}

type DescribeRobotBizIDByAppKeyRequest struct {
	*tchttp.BaseRequest
	
	// 应用appkey
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
	// 应用业务ID
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
type DescribeSearchStatsGraphRequestParams struct {
	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// uin列表
	UinAccount []*string `json:"UinAccount,omitnil,omitempty" name:"UinAccount"`

	// 子业务类型
	SubBizType *string `json:"SubBizType,omitnil,omitempty" name:"SubBizType"`

	// 模型标识
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 应用id列表
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`
}

type DescribeSearchStatsGraphRequest struct {
	*tchttp.BaseRequest
	
	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// uin列表
	UinAccount []*string `json:"UinAccount,omitnil,omitempty" name:"UinAccount"`

	// 子业务类型
	SubBizType *string `json:"SubBizType,omitnil,omitempty" name:"SubBizType"`

	// 模型标识
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 应用id列表
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`
}

func (r *DescribeSearchStatsGraphRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSearchStatsGraphRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	delete(f, "UinAccount")
	delete(f, "SubBizType")
	delete(f, "ModelName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "AppBizIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSearchStatsGraphRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSearchStatsGraphResponseParams struct {
	// 统计结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*Stat `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSearchStatsGraphResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSearchStatsGraphResponseParams `json:"Response"`
}

func (r *DescribeSearchStatsGraphResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSearchStatsGraphResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSegmentsRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文档ID
	SegBizId []*string `json:"SegBizId,omitnil,omitempty" name:"SegBizId"`
}

type DescribeSegmentsRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文档ID
	SegBizId []*string `json:"SegBizId,omitnil,omitempty" name:"SegBizId"`
}

func (r *DescribeSegmentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSegmentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "SegBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSegmentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSegmentsResponseParams struct {
	// 片段列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*DocSegment `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSegmentsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSegmentsResponseParams `json:"Response"`
}

func (r *DescribeSegmentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSegmentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStorageCredentialRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文件类型,正常的文件名类型后缀，例如 xlsx、pdf、 docx、png 等
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// IsPublic用于上传文件时选择场景，当上传为对话端文件时IsPublic为true，上传文档库文件时场景IsPublic为false
	IsPublic *bool `json:"IsPublic,omitnil,omitempty" name:"IsPublic"`

	// 存储类型: offline:离线文件，realtime:实时文件；为空默认为offline
	TypeKey *string `json:"TypeKey,omitnil,omitempty" name:"TypeKey"`
}

type DescribeStorageCredentialRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文件类型,正常的文件名类型后缀，例如 xlsx、pdf、 docx、png 等
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// IsPublic用于上传文件时选择场景，当上传为对话端文件时IsPublic为true，上传文档库文件时场景IsPublic为false
	IsPublic *bool `json:"IsPublic,omitnil,omitempty" name:"IsPublic"`

	// 存储类型: offline:离线文件，realtime:实时文件；为空默认为offline
	TypeKey *string `json:"TypeKey,omitnil,omitempty" name:"TypeKey"`
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
	delete(f, "FileType")
	delete(f, "IsPublic")
	delete(f, "TypeKey")
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

	// 文件存储目录
	FilePath *string `json:"FilePath,omitnil,omitempty" name:"FilePath"`

	// 存储类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 主号
	CorpUin *string `json:"CorpUin,omitnil,omitempty" name:"CorpUin"`

	// 图片存储目录
	ImagePath *string `json:"ImagePath,omitnil,omitempty" name:"ImagePath"`

	// 上传存储路径，到具体文件
	UploadPath *string `json:"UploadPath,omitnil,omitempty" name:"UploadPath"`

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
type DescribeTokenUsageGraphRequestParams struct {
	// 腾讯云主账号
	UinAccount []*string `json:"UinAccount,omitnil,omitempty" name:"UinAccount"`

	// 知识引擎子业务类型:  FileParse(文档解析)、Embedding、Rewrite(多轮改写)、 Concurrency(并发)、KnowledgeSummary(知识总结)   KnowledgeQA(知识问答)、KnowledgeCapacity(知识库容量)、SearchEngine(搜索引擎)
	SubBizType *string `json:"SubBizType,omitnil,omitempty" name:"SubBizType"`

	// 模型标识
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 应用id列表
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`
}

type DescribeTokenUsageGraphRequest struct {
	*tchttp.BaseRequest
	
	// 腾讯云主账号
	UinAccount []*string `json:"UinAccount,omitnil,omitempty" name:"UinAccount"`

	// 知识引擎子业务类型:  FileParse(文档解析)、Embedding、Rewrite(多轮改写)、 Concurrency(并发)、KnowledgeSummary(知识总结)   KnowledgeQA(知识问答)、KnowledgeCapacity(知识库容量)、SearchEngine(搜索引擎)
	SubBizType *string `json:"SubBizType,omitnil,omitempty" name:"SubBizType"`

	// 模型标识
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 应用id列表
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`
}

func (r *DescribeTokenUsageGraphRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTokenUsageGraphRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UinAccount")
	delete(f, "SubBizType")
	delete(f, "ModelName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "AppBizIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTokenUsageGraphRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTokenUsageGraphResponseParams struct {
	// 总消耗
	Total []*Stat `json:"Total,omitnil,omitempty" name:"Total"`

	// 输入消耗
	Input []*Stat `json:"Input,omitnil,omitempty" name:"Input"`

	// 输出消耗
	Output []*Stat `json:"Output,omitnil,omitempty" name:"Output"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTokenUsageGraphResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTokenUsageGraphResponseParams `json:"Response"`
}

func (r *DescribeTokenUsageGraphResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTokenUsageGraphResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTokenUsageRequestParams struct {
	// 腾讯云主账号
	UinAccount []*string `json:"UinAccount,omitnil,omitempty" name:"UinAccount"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 知识引擎子业务类型:  FileParse(文档解析)、Embedding、Rewrite(多轮改写)、 Concurrency(并发)、KnowledgeSummary(知识总结)   KnowledgeQA(知识问答)、KnowledgeCapacity(知识库容量)、SearchEngine(搜索引擎)
	SubBizType *string `json:"SubBizType,omitnil,omitempty" name:"SubBizType"`

	// 模型标识
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 应用id列表
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`
}

type DescribeTokenUsageRequest struct {
	*tchttp.BaseRequest
	
	// 腾讯云主账号
	UinAccount []*string `json:"UinAccount,omitnil,omitempty" name:"UinAccount"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 知识引擎子业务类型:  FileParse(文档解析)、Embedding、Rewrite(多轮改写)、 Concurrency(并发)、KnowledgeSummary(知识总结)   KnowledgeQA(知识问答)、KnowledgeCapacity(知识库容量)、SearchEngine(搜索引擎)
	SubBizType *string `json:"SubBizType,omitnil,omitempty" name:"SubBizType"`

	// 模型标识
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 应用id列表
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`
}

func (r *DescribeTokenUsageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTokenUsageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UinAccount")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	delete(f, "SubBizType")
	delete(f, "ModelName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "AppBizIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTokenUsageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTokenUsageResponseParams struct {
	// 总token消耗量
	TotalTokenUsage *float64 `json:"TotalTokenUsage,omitnil,omitempty" name:"TotalTokenUsage"`

	// 输入token消耗
	InputTokenUsage *float64 `json:"InputTokenUsage,omitnil,omitempty" name:"InputTokenUsage"`

	// 输出token消耗
	OutputTokenUsage *float64 `json:"OutputTokenUsage,omitnil,omitempty" name:"OutputTokenUsage"`

	// 接口调用次数
	ApiCallStats *uint64 `json:"ApiCallStats,omitnil,omitempty" name:"ApiCallStats"`

	// 搜索服务调用次数
	SearchUsage *float64 `json:"SearchUsage,omitnil,omitempty" name:"SearchUsage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTokenUsageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTokenUsageResponseParams `json:"Response"`
}

func (r *DescribeTokenUsageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTokenUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUnsatisfiedReplyContextRequestParams struct {
	// 应用ID
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
	
	// 应用ID
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

type DocSegment struct {
	// 片段ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 业务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessId *string `json:"BusinessId,omitnil,omitempty" name:"BusinessId"`

	// 文件类型(markdown,word,txt)
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 文档切片类型(segment-文档切片 table-表格)
	// 注意：此字段可能返回 null，表示取不到有效值。
	SegmentType *string `json:"SegmentType,omitnil,omitempty" name:"SegmentType"`

	// 标题
	// 注意：此字段可能返回 null，表示取不到有效值。
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 段落内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageContent *string `json:"PageContent,omitnil,omitempty" name:"PageContent"`

	// 段落原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrgData *string `json:"OrgData,omitnil,omitempty" name:"OrgData"`

	// 文章ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DocId *string `json:"DocId,omitnil,omitempty" name:"DocId"`

	// 文档业务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// 文档链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	DocUrl *string `json:"DocUrl,omitnil,omitempty" name:"DocUrl"`
}

type DocumentElement struct {
	// 文档元素索引
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index *int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 元素类型，包括paragraph、table、formula、figure、title、header、footer、figure_text
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 元素内容，当type为figure或formula(公式识别关闭)时该字段内容为图片的位置
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 元素坐标，左上角(x1, y1)，右上角(x2, y2)，右下角(x3, y3)，左下角(x4, y4)
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Polygon *Polygon `json:"Polygon,omitnil,omitempty" name:"Polygon"`

	// 元素层级
	// 注意：此字段可能返回 null，表示取不到有效值。
	Level *int64 `json:"Level,omitnil,omitempty" name:"Level"`

	// 入参开启EnableInsetImage后返回，表示在InsetImagePackage中的内嵌图片名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InsetImageName *string `json:"InsetImageName,omitnil,omitempty" name:"InsetImageName"`

	// 嵌套的文档元素信息，一般包含的是文档内嵌入图片的文字识别结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Elements []*DocumentElement `json:"Elements,omitnil,omitempty" name:"Elements"`
}

type DocumentRecognizeInfo struct {
	// 输入PDF文件的页码，从1开始。输入图片的话值始终为1
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 旋转角度
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Angle *int64 `json:"Angle,omitnil,omitempty" name:"Angle"`

	// AI算法识别处理后的图片高度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// AI算法识别处理后的图片宽度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 图片的原始高度，输入PDF文件则表示单页PDF转图片之后的图片高度
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginHeight *int64 `json:"OriginHeight,omitnil,omitempty" name:"OriginHeight"`

	// 图片的原始宽度，输入PDF文件则表示单页PDF转图片之后的图片宽度
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginWidth *int64 `json:"OriginWidth,omitnil,omitempty" name:"OriginWidth"`

	// 文档元素信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Elements []*DocumentElement `json:"Elements,omitnil,omitempty" name:"Elements"`

	// 旋转角度
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	RotatedAngle *float64 `json:"RotatedAngle,omitnil,omitempty" name:"RotatedAngle"`
}

type EmbeddingObject struct {
	// 向量
	Embedding []*float64 `json:"Embedding,omitnil,omitempty" name:"Embedding"`
}

// Predefined struct for user
type ExportAttributeLabelRequestParams struct {
	// 应用ID
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
	
	// 应用ID
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
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// QA业务ID
	QaBizIds []*string `json:"QaBizIds,omitnil,omitempty" name:"QaBizIds"`

	// 查询参数
	Filters *QAQuery `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type ExportQAListRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
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
	// 应用ID
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
	
	// 应用ID
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

type FileInfo struct {
	// 文件名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文件大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileSize *string `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// 文件的URL地址，COS地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// 文件类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 解析后返回的DocID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DocId *string `json:"DocId,omitnil,omitempty" name:"DocId"`
}

type Filters struct {
	// 检索，用户问题或答案
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 错误类型检索
	Reasons []*string `json:"Reasons,omitnil,omitempty" name:"Reasons"`
}

// Predefined struct for user
type GenerateQARequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文档ID
	DocBizIds []*string `json:"DocBizIds,omitnil,omitempty" name:"DocBizIds"`
}

type GenerateQARequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
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
type GetAnswerTypeDataCountRequestParams struct {
	// 开始日期
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束日期
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 应用id
	AppBizId []*string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 消息来源(1、分享用户端  2、对话API  3、对话测试  4、应用评测)
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)	
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

type GetAnswerTypeDataCountRequest struct {
	*tchttp.BaseRequest
	
	// 开始日期
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束日期
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 应用id
	AppBizId []*string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 消息来源(1、分享用户端  2、对话API  3、对话测试  4、应用评测)
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)	
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

func (r *GetAnswerTypeDataCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAnswerTypeDataCountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "AppBizId")
	delete(f, "Type")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetAnswerTypeDataCountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAnswerTypeDataCountResponseParams struct {
	// 总消息数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 大模型直接回复总数
	ModelReplyCount *uint64 `json:"ModelReplyCount,omitnil,omitempty" name:"ModelReplyCount"`

	// 知识型回复总数
	KnowledgeCount *uint64 `json:"KnowledgeCount,omitnil,omitempty" name:"KnowledgeCount"`

	// 任务流回复总数
	TaskFlowCount *uint64 `json:"TaskFlowCount,omitnil,omitempty" name:"TaskFlowCount"`

	// 搜索引擎回复总数
	SearchEngineCount *uint64 `json:"SearchEngineCount,omitnil,omitempty" name:"SearchEngineCount"`

	// 图片理解回复总数
	ImageUnderstandingCount *uint64 `json:"ImageUnderstandingCount,omitnil,omitempty" name:"ImageUnderstandingCount"`

	// 拒答回复总数
	RejectCount *uint64 `json:"RejectCount,omitnil,omitempty" name:"RejectCount"`

	// 敏感回复总数
	SensitiveCount *uint64 `json:"SensitiveCount,omitnil,omitempty" name:"SensitiveCount"`

	// 并发超限回复总数
	ConcurrentLimitCount *uint64 `json:"ConcurrentLimitCount,omitnil,omitempty" name:"ConcurrentLimitCount"`

	// 未知问题回复总数
	UnknownIssuesCount *uint64 `json:"UnknownIssuesCount,omitnil,omitempty" name:"UnknownIssuesCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetAnswerTypeDataCountResponse struct {
	*tchttp.BaseResponse
	Response *GetAnswerTypeDataCountResponseParams `json:"Response"`
}

func (r *GetAnswerTypeDataCountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAnswerTypeDataCountResponse) FromJsonString(s string) error {
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
	// 文档BizID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 存储类型: offline:离线文件，realtime:实时文件；为空默认为offline
	TypeKey *string `json:"TypeKey,omitnil,omitempty" name:"TypeKey"`
}

type GetDocPreviewRequest struct {
	*tchttp.BaseRequest
	
	// 文档BizID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 存储类型: offline:离线文件，realtime:实时文件；为空默认为offline
	TypeKey *string `json:"TypeKey,omitnil,omitempty" name:"TypeKey"`
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
	delete(f, "TypeKey")
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

	// 消耗量，返回TotalToken
	Usage *Usage `json:"Usage,omitnil,omitempty" name:"Usage"`

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
type GetLikeDataCountRequestParams struct {
	// 开始日期
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束日期
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 应用id
	AppBizId []*string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 消息来源(1、分享用户端  2、对话API)
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)	
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

type GetLikeDataCountRequest struct {
	*tchttp.BaseRequest
	
	// 开始日期
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束日期
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 应用id
	AppBizId []*string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 消息来源(1、分享用户端  2、对话API)
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)	
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

func (r *GetLikeDataCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetLikeDataCountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "AppBizId")
	delete(f, "Type")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetLikeDataCountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetLikeDataCountResponseParams struct {
	// 可评价消息数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 评价数
	AppraisalTotal *uint64 `json:"AppraisalTotal,omitnil,omitempty" name:"AppraisalTotal"`

	// 参评率
	ParticipationRate *float64 `json:"ParticipationRate,omitnil,omitempty" name:"ParticipationRate"`

	// 点赞数
	LikeTotal *uint64 `json:"LikeTotal,omitnil,omitempty" name:"LikeTotal"`

	// 点赞率
	LikeRate *float64 `json:"LikeRate,omitnil,omitempty" name:"LikeRate"`

	// 点踩数
	DislikeTotal *uint64 `json:"DislikeTotal,omitnil,omitempty" name:"DislikeTotal"`

	// 点踩率
	DislikeRate *float64 `json:"DislikeRate,omitnil,omitempty" name:"DislikeRate"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetLikeDataCountResponse struct {
	*tchttp.BaseResponse
	Response *GetLikeDataCountResponseParams `json:"Response"`
}

func (r *GetLikeDataCountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetLikeDataCountResponse) FromJsonString(s string) error {
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

	// 应用AppKey, 当Type=5[API访客]时, 该字段必填
	BotAppKey *string `json:"BotAppKey,omitnil,omitempty" name:"BotAppKey"`

	// 场景, 体验: 1; 正式: 2
	Scene *uint64 `json:"Scene,omitnil,omitempty" name:"Scene"`
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

	// 应用AppKey, 当Type=5[API访客]时, 该字段必填
	BotAppKey *string `json:"BotAppKey,omitnil,omitempty" name:"BotAppKey"`

	// 场景, 体验: 1; 正式: 2
	Scene *uint64 `json:"Scene,omitnil,omitempty" name:"Scene"`
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
	delete(f, "Scene")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetMsgRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetMsgRecordResponseParams struct {
	// 会话记录
	Records []*MsgRecord `json:"Records,omitnil,omitempty" name:"Records"`

	// session 清除关联上下文时间, 单位 ms
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionDisassociatedTimestamp *string `json:"SessionDisassociatedTimestamp,omitnil,omitempty" name:"SessionDisassociatedTimestamp"`

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
type GetReconstructDocumentResultRequestParams struct {
	// 任务唯一Id。[CreateReconstructDocumentFlow](https://cloud.tencent.com/document/product/1759/107506) 返回的TaskId。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type GetReconstructDocumentResultRequest struct {
	*tchttp.BaseRequest
	
	// 任务唯一Id。[CreateReconstructDocumentFlow](https://cloud.tencent.com/document/product/1759/107506) 返回的TaskId。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *GetReconstructDocumentResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetReconstructDocumentResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetReconstructDocumentResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetReconstructDocumentResultResponseParams struct {
	// 任务状态: Success->执行完成；Processing->执行中；Failed->执行失败；WaitExecute->等待执行。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 本次文档解析的结果文件，存储在腾讯云COS的下载URL，下载URL的有效期为10分钟。
	DocumentRecognizeResultUrl *string `json:"DocumentRecognizeResultUrl,omitnil,omitempty" name:"DocumentRecognizeResultUrl"`

	// 本次文档解析失败的页码信息。
	FailedPages []*ReconstructDocumentFailedPage `json:"FailedPages,omitnil,omitempty" name:"FailedPages"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetReconstructDocumentResultResponse struct {
	*tchttp.BaseResponse
	Response *GetReconstructDocumentResultResponseParams `json:"Response"`
}

func (r *GetReconstructDocumentResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetReconstructDocumentResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTaskStatusRequestParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务类型
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`
}

type GetTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务类型
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 应用ID
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

	// 应用AppKey（应用发布后在应用页面[发布管理]-[调用信息]-[API管理]处获取）
	BotAppKey *string `json:"BotAppKey,omitnil,omitempty" name:"BotAppKey"`

	// 访客ID（外部输入，建议唯一，标识当前接入会话的用户）
	VisitorBizId *string `json:"VisitorBizId,omitnil,omitempty" name:"VisitorBizId"`

	// 知识标签（用于知识库中知识的检索过滤）
	VisitorLabels []*GetWsTokenReq_Label `json:"VisitorLabels,omitnil,omitempty" name:"VisitorLabels"`
}

type GetWsTokenRequest struct {
	*tchttp.BaseRequest
	
	// 接入类型
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 应用AppKey（应用发布后在应用页面[发布管理]-[调用信息]-[API管理]处获取）
	BotAppKey *string `json:"BotAppKey,omitnil,omitempty" name:"BotAppKey"`

	// 访客ID（外部输入，建议唯一，标识当前接入会话的用户）
	VisitorBizId *string `json:"VisitorBizId,omitnil,omitempty" name:"VisitorBizId"`

	// 知识标签（用于知识库中知识的检索过滤）
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
	// token值（有效期60s）
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// 余额; 余额大于 0 时表示有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	Balance *float64 `json:"Balance,omitnil,omitempty" name:"Balance"`

	// 对话窗输入字符限制
	InputLenLimit *int64 `json:"InputLenLimit,omitnil,omitempty" name:"InputLenLimit"`

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
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// QaBizID列表
	QaBizIds []*string `json:"QaBizIds,omitnil,omitempty" name:"QaBizIds"`

	// 分组 ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`
}

type GroupQARequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// QaBizID列表
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
	// 高亮起始位置
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

type HistorySummary struct {
	// 助手
	// 注意：此字段可能返回 null，表示取不到有效值。
	Assistant *string `json:"Assistant,omitnil,omitempty" name:"Assistant"`

	// 用户
	// 注意：此字段可能返回 null，表示取不到有效值。
	User *string `json:"User,omitnil,omitempty" name:"User"`
}

// Predefined struct for user
type IgnoreUnsatisfiedReplyRequestParams struct {
	// 应用ID
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
	
	// 应用ID
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

type InvokeAPI struct {
	// 请求方法，如GET/POST等
	// 注意：此字段可能返回 null，表示取不到有效值。
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// 请求地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// header参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeaderValues []*StrValue `json:"HeaderValues,omitnil,omitempty" name:"HeaderValues"`

	// 入参Query
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueryValues []*StrValue `json:"QueryValues,omitnil,omitempty" name:"QueryValues"`

	// Post请求的原始数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	RequestPostBody *string `json:"RequestPostBody,omitnil,omitempty" name:"RequestPostBody"`

	// 返回的原始数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResponseBody *string `json:"ResponseBody,omitnil,omitempty" name:"ResponseBody"`

	// 出参
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResponseValues []*ValueInfo `json:"ResponseValues,omitnil,omitempty" name:"ResponseValues"`

	// 异常信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailMessage *string `json:"FailMessage,omitnil,omitempty" name:"FailMessage"`
}

// Predefined struct for user
type IsTransferIntentRequestParams struct {
	// 内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 应用appKey
	BotAppKey *string `json:"BotAppKey,omitnil,omitempty" name:"BotAppKey"`
}

type IsTransferIntentRequest struct {
	*tchttp.BaseRequest
	
	// 内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 应用appKey
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

type KnowledgeCapacityPieGraphDetail struct {
	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 应用使用的字符数
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsedCharSize *string `json:"UsedCharSize,omitnil,omitempty" name:"UsedCharSize"`

	// 应用占比
	// 注意：此字段可能返回 null，表示取不到有效值。
	Proportion *float64 `json:"Proportion,omitnil,omitempty" name:"Proportion"`
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

	// 是否展示问题澄清开关
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShowQuestionClarify *bool `json:"ShowQuestionClarify,omitnil,omitempty" name:"ShowQuestionClarify"`

	// 是否打开问题澄清
	// 注意：此字段可能返回 null，表示取不到有效值。
	UseQuestionClarify *bool `json:"UseQuestionClarify,omitnil,omitempty" name:"UseQuestionClarify"`

	// 问题澄清关键词列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	QuestionClarifyKeywords []*string `json:"QuestionClarifyKeywords,omitnil,omitempty" name:"QuestionClarifyKeywords"`

	// 是否打开推荐问题开关
	// 注意：此字段可能返回 null，表示取不到有效值。
	UseRecommended *bool `json:"UseRecommended,omitnil,omitempty" name:"UseRecommended"`
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

	// 检索置信度，针对文档和问答有效，最小0.01，最大0.99
	// 注意：此字段可能返回 null，表示取不到有效值。
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// 资源状态 1：资源可用；2：资源已用尽
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceStatus *uint64 `json:"ResourceStatus,omitnil,omitempty" name:"ResourceStatus"`
}

type KnowledgeSummary struct {
	// 1是问答 2是文档片段
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 知识内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
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
	// 应用ID
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
	
	// 应用ID
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

	// 是否允许重试，0：否，1：是
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAllowRetry *bool `json:"IsAllowRetry,omitnil,omitempty" name:"IsAllowRetry"`
}

// Predefined struct for user
type ListDocRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询内容
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 文档状态： 1-未生成 2-生成中 3-生成成功 4-生成失败 5-删除中 6-删除成功  7-审核中  8-审核失败 9-审核成功  10-待发布  11-发布中  12-已发布  13-学习中  14-学习失败  15-更新中  16-更新失败  17-解析中  18-解析失败  19-导入失败   20-已过期 21-超量失效 22-超量失效恢复
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type ListDocRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询内容
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 文档状态： 1-未生成 2-生成中 3-生成成功 4-生成失败 5-删除中 6-删除成功  7-审核中  8-审核失败 9-审核成功  10-待发布  11-发布中  12-已发布  13-学习中  14-学习失败  15-更新中  16-更新失败  17-解析中  18-解析失败  19-导入失败   20-已过期 21-超量失效 22-超量失效恢复
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
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`
}

type ListQACateRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
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
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询问题
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 校验状态(1未校验2采纳3不采纳)
	AcceptStatus []*int64 `json:"AcceptStatus,omitnil,omitempty" name:"AcceptStatus"`

	// 发布状态(2待发布 3发布中 4已发布 7审核中 8审核失败 9人工申述中 11人工申述失败 12已过期 13超量失效 14超量失效恢复)
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
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询问题
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 校验状态(1未校验2采纳3不采纳)
	AcceptStatus []*int64 `json:"AcceptStatus,omitnil,omitempty" name:"AcceptStatus"`

	// 发布状态(2待发布 3发布中 4已发布 7审核中 8审核失败 9人工申述中 11人工申述失败 12已过期 13超量失效 14超量失效恢复)
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

	// 有效开始时间，unix时间戳
	ExpireStart *string `json:"ExpireStart,omitnil,omitempty" name:"ExpireStart"`

	// 有效结束时间，unix时间戳，0代表永久有效
	ExpireEnd *string `json:"ExpireEnd,omitnil,omitempty" name:"ExpireEnd"`

	// 属性标签适用范围 1：全部，2：按条件
	AttrRange *int64 `json:"AttrRange,omitnil,omitempty" name:"AttrRange"`

	// 属性标签
	AttrLabels []*AttrLabel `json:"AttrLabels,omitnil,omitempty" name:"AttrLabels"`

	// 相似问个数
	SimilarQuestionNum *uint64 `json:"SimilarQuestionNum,omitnil,omitempty" name:"SimilarQuestionNum"`
}

// Predefined struct for user
type ListRejectedQuestionPreviewRequestParams struct {
	// 应用ID
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
	
	// 应用ID
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
	// 应用ID
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
	
	// 应用ID
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
	// 应用ID
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
}

type ListReleaseDocPreviewRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
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
	// 应用ID
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
	
	// 应用ID
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
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文档名称
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文档状态： 7 审核中、8 审核失败、10 待发布、11 发布中、12 已发布、13 学习中、14 学习失败 20 已过期
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type ListSelectDocRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
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
	// 应用ID
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
	
	// 应用ID
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

	// 资源状态 1：资源可用；2：资源已用尽
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceStatus *uint64 `json:"ResourceStatus,omitnil,omitempty" name:"ResourceStatus"`

	// 提示词内容字符限制
	// 注意：此字段可能返回 null，表示取不到有效值。
	PromptWordsLimit *string `json:"PromptWordsLimit,omitnil,omitempty" name:"PromptWordsLimit"`
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
	// 应用ID
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
	
	// 应用ID
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
	// 应用ID
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
	
	// 应用ID
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
	// 应用ID
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
	
	// 应用ID
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
	// 应用ID
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
	
	// 应用ID
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
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 分类名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分类业务ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`
}

type ModifyQACateRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
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
	// 应用ID
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

	// 相似问修改信息(相似问没有修改则不传)
	SimilarQuestionModify *SimilarQuestionModify `json:"SimilarQuestionModify,omitnil,omitempty" name:"SimilarQuestionModify"`
}

type ModifyQARequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
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

	// 相似问修改信息(相似问没有修改则不传)
	SimilarQuestionModify *SimilarQuestionModify `json:"SimilarQuestionModify,omitnil,omitempty" name:"SimilarQuestionModify"`
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
	delete(f, "SimilarQuestionModify")
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
	// 应用ID
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
	
	// 应用ID
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

type MsgFileInfo struct {
	// 文档名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文档大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileSize *string `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// 文档URL
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// 文档类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 文档ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DocId *string `json:"DocId,omitnil,omitempty" name:"DocId"`
}

type MsgRecord struct {
	// 内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 当前记录所对应的 Session ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

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

	// 是否展示反馈按钮
	// 注意：此字段可能返回 null，表示取不到有效值。
	CanFeedback *bool `json:"CanFeedback,omitnil,omitempty" name:"CanFeedback"`

	// 记录类型
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 引用来源
	References []*MsgRecordReference `json:"References,omitnil,omitempty" name:"References"`

	// 评价原因
	Reasons []*string `json:"Reasons,omitnil,omitempty" name:"Reasons"`

	// 是否大模型
	IsLlmGenerated *bool `json:"IsLlmGenerated,omitnil,omitempty" name:"IsLlmGenerated"`

	// 图片链接，可公有读
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageUrls []*string `json:"ImageUrls,omitnil,omitempty" name:"ImageUrls"`

	// 当次 token 统计信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TokenStat *TokenStat `json:"TokenStat,omitnil,omitempty" name:"TokenStat"`

	// 回复方式
	// 1:大模型直接回复;
	// 2:保守回复, 未知问题回复;
	// 3:拒答问题回复;
	// 4:敏感回复;
	// 5:问答对直接回复, 已采纳问答对优先回复;
	// 6:欢迎语回复;
	// 7:并发超限回复;
	// 8:全局干预知识;
	// 9:任务流程过程回复, 当历史记录中 task_flow.type = 0 时, 为大模型回复;
	// 10:任务流程答案回复;
	// 11:搜索引擎回复;
	// 12:知识润色后回复;
	// 13:图片理解回复;
	// 14:实时文档回复;
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReplyMethod *uint64 `json:"ReplyMethod,omitnil,omitempty" name:"ReplyMethod"`

	// 选项卡, 用于多轮对话
	// 注意：此字段可能返回 null，表示取不到有效值。
	OptionCards []*string `json:"OptionCards,omitnil,omitempty" name:"OptionCards"`

	// 任务信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskFlow *TaskFlowInfo `json:"TaskFlow,omitnil,omitempty" name:"TaskFlow"`

	// 用户传入的文件信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileInfos []*FileInfo `json:"FileInfos,omitnil,omitempty" name:"FileInfos"`

	// 参考来源引用位置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	QuoteInfos []*QuoteInfo `json:"QuoteInfos,omitnil,omitempty" name:"QuoteInfos"`
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

	// 文件下载链接 (支持的文件类型: docx, txt, markdown, pdf), 该地址需要外网可以直接无状态访问
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 任务ID, 用于幂等去重, 业务自行定义(最大长度64字节)
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 切分策略
	Policy *string `json:"Policy,omitnil,omitempty" name:"Policy"`

	// 默认值: parse
	//
	// Deprecated: Operate is deprecated.
	Operate *string `json:"Operate,omitnil,omitempty" name:"Operate"`
}

type ParseDocRequest struct {
	*tchttp.BaseRequest
	
	// 文件名称(需要包括文件后缀, 最大长度1024字节)
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 文件下载链接 (支持的文件类型: docx, txt, markdown, pdf), 该地址需要外网可以直接无状态访问
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 任务ID, 用于幂等去重, 业务自行定义(最大长度64字节)
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 切分策略
	Policy *string `json:"Policy,omitnil,omitempty" name:"Policy"`

	// 默认值: parse
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

type Polygon struct {
	// 左上顶点坐标
	LeftTop *Coord `json:"LeftTop,omitnil,omitempty" name:"LeftTop"`

	// 右上顶点坐标
	RightTop *Coord `json:"RightTop,omitnil,omitempty" name:"RightTop"`

	// 右下顶点坐标
	RightBottom *Coord `json:"RightBottom,omitnil,omitempty" name:"RightBottom"`

	// 左下顶点坐标
	LeftBottom *Coord `json:"LeftBottom,omitnil,omitempty" name:"LeftBottom"`
}

type Procedure struct {
	// 执行过程英语名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 中文名, 用于展示
	// 注意：此字段可能返回 null，表示取不到有效值。
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 状态常量: 使用中: processing, 成功: success, 失败: failed
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 消耗 token 数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 调试信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Debugging *ProcedureDebugging `json:"Debugging,omitnil,omitempty" name:"Debugging"`

	// 计费资源状态，1：可用，2：不可用
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceStatus *uint64 `json:"ResourceStatus,omitnil,omitempty" name:"ResourceStatus"`
}

type ProcedureDebugging struct {
	// 检索query
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 系统prompt
	// 注意：此字段可能返回 null，表示取不到有效值。
	System *string `json:"System,omitnil,omitempty" name:"System"`

	// 多轮历史信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Histories []*HistorySummary `json:"Histories,omitnil,omitempty" name:"Histories"`

	// 检索知识
	// 注意：此字段可能返回 null，表示取不到有效值。
	Knowledge []*KnowledgeSummary `json:"Knowledge,omitnil,omitempty" name:"Knowledge"`

	// 任务流程
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskFlow *TaskFlowSummary `json:"TaskFlow,omitnil,omitempty" name:"TaskFlow"`
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

	// 应用ID
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

	// 消耗量，输出页数
	Usage *Usage `json:"Usage,omitnil,omitempty" name:"Usage"`

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

	// 消耗量，返回输入token数，输出token数以及总token数
	Usage *Usage `json:"Usage,omitnil,omitempty" name:"Usage"`

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

type QuoteInfo struct {
	// 参考来源位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Position *uint64 `json:"Position,omitnil,omitempty" name:"Position"`

	// 参考来源索引顺序
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`
}

// Predefined struct for user
type RateMsgRecordRequestParams struct {
	// 应用appKey
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
	
	// 应用appKey
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

type ReconstructDocumentConfig struct {
	// 生成的Markdown中是否嵌入图片
	EnableInsetImage *bool `json:"EnableInsetImage,omitnil,omitempty" name:"EnableInsetImage"`
}

type ReconstructDocumentFailedPage struct {
	// 失败页码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

// Predefined struct for user
type ReconstructDocumentRequestParams struct {
	// 文件的 Base64 值。 支持的文件格式：PNG、JPG、JPEG、PDF。 支持的文件大小：所下载文件经Base64编码后不超过 8M。文件下载时间不超过 3 秒。 支持的图片像素：单边介于20-10000px之间。 文件的 FileUrl、FileBase64 必须提供一个，如果都提供，只使用 FileUrl。
	FileBase64 *string `json:"FileBase64,omitnil,omitempty" name:"FileBase64"`

	// 文件的 Url 地址。 支持的文件格式：PNG、JPG、JPEG、PDF。 支持的文件大小：所下载文件经 Base64 编码后不超过 8M。文件下载时间不超过 3 秒。 支持的图片像素：单边介于20-10000px之间。 文件存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议文件存储于腾讯云。 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// 当传入文件是PDF类型时，用来指定pdf识别的起始页码，识别的页码包含当前值。默认为1，表示从pdf文件的第1页开始识别。
	FileStartPageNumber *int64 `json:"FileStartPageNumber,omitnil,omitempty" name:"FileStartPageNumber"`

	// 当传入文件是PDF类型时，用来指定pdf识别的结束页码，识别的页码包含当前值。默认为10，表示识别到pdf文件的第10页。单次调用最多支持识别10页内容，即FileEndPageNumber-FileStartPageNumber需要不大于10。
	FileEndPageNumber *int64 `json:"FileEndPageNumber,omitnil,omitempty" name:"FileEndPageNumber"`

	// 配置选项，支持配置是否在生成的Markdown中是否嵌入图片
	Config *ReconstructDocumentConfig `json:"Config,omitnil,omitempty" name:"Config"`
}

type ReconstructDocumentRequest struct {
	*tchttp.BaseRequest
	
	// 文件的 Base64 值。 支持的文件格式：PNG、JPG、JPEG、PDF。 支持的文件大小：所下载文件经Base64编码后不超过 8M。文件下载时间不超过 3 秒。 支持的图片像素：单边介于20-10000px之间。 文件的 FileUrl、FileBase64 必须提供一个，如果都提供，只使用 FileUrl。
	FileBase64 *string `json:"FileBase64,omitnil,omitempty" name:"FileBase64"`

	// 文件的 Url 地址。 支持的文件格式：PNG、JPG、JPEG、PDF。 支持的文件大小：所下载文件经 Base64 编码后不超过 8M。文件下载时间不超过 3 秒。 支持的图片像素：单边介于20-10000px之间。 文件存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议文件存储于腾讯云。 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// 当传入文件是PDF类型时，用来指定pdf识别的起始页码，识别的页码包含当前值。默认为1，表示从pdf文件的第1页开始识别。
	FileStartPageNumber *int64 `json:"FileStartPageNumber,omitnil,omitempty" name:"FileStartPageNumber"`

	// 当传入文件是PDF类型时，用来指定pdf识别的结束页码，识别的页码包含当前值。默认为10，表示识别到pdf文件的第10页。单次调用最多支持识别10页内容，即FileEndPageNumber-FileStartPageNumber需要不大于10。
	FileEndPageNumber *int64 `json:"FileEndPageNumber,omitnil,omitempty" name:"FileEndPageNumber"`

	// 配置选项，支持配置是否在生成的Markdown中是否嵌入图片
	Config *ReconstructDocumentConfig `json:"Config,omitnil,omitempty" name:"Config"`
}

func (r *ReconstructDocumentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReconstructDocumentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileBase64")
	delete(f, "FileUrl")
	delete(f, "FileStartPageNumber")
	delete(f, "FileEndPageNumber")
	delete(f, "Config")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReconstructDocumentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReconstructDocumentResponseParams struct {
	// 识别生成的Markdown文件base64编码的字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	MarkdownBase64 *string `json:"MarkdownBase64,omitnil,omitempty" name:"MarkdownBase64"`

	// 输入文件中嵌入的图片放在一个文件夹中打包为.zip压缩文件，识别生成的Markdown文件通过路径关联插入本文件夹中的图片。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InsetImagePackage *string `json:"InsetImagePackage,omitnil,omitempty" name:"InsetImagePackage"`

	// 输入文件中嵌入的图片中文字内容的识别结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	DocumentRecognizeInfo []*DocumentRecognizeInfo `json:"DocumentRecognizeInfo,omitnil,omitempty" name:"DocumentRecognizeInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReconstructDocumentResponse struct {
	*tchttp.BaseResponse
	Response *ReconstructDocumentResponseParams `json:"Response"`
}

func (r *ReconstructDocumentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReconstructDocumentResponse) FromJsonString(s string) error {
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

	// 页码信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageInfos []*uint64 `json:"PageInfos,omitnil,omitempty" name:"PageInfos"`

	// sheet信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SheetInfos []*string `json:"SheetInfos,omitnil,omitempty" name:"SheetInfos"`

	// 文档ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`
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

	// 是否仅清空会话关联
	IsOnlyEmptyTheDialog *bool `json:"IsOnlyEmptyTheDialog,omitnil,omitempty" name:"IsOnlyEmptyTheDialog"`
}

type ResetSessionRequest struct {
	*tchttp.BaseRequest
	
	// 会话ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 是否仅清空会话关联
	IsOnlyEmptyTheDialog *bool `json:"IsOnlyEmptyTheDialog,omitnil,omitempty" name:"IsOnlyEmptyTheDialog"`
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
	delete(f, "IsOnlyEmptyTheDialog")
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
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`
}

type RetryDocAuditRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
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
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`
}

type RetryDocParseRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
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

type RunNodeInfo struct {
	// 节点类型，0:未指定，1:开始节点，2:API节点，3:询问节点，4:答案节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeType *int64 `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// 节点ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 节点名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// 请求的API
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvokeApi *InvokeAPI `json:"InvokeApi,omitnil,omitempty" name:"InvokeApi"`

	// 当前节点的所有槽位的值，key：SlotID。没有值的时候也要返回空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlotValues []*ValueInfo `json:"SlotValues,omitnil,omitempty" name:"SlotValues"`
}

// Predefined struct for user
type SaveDocRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文件名
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文件类型(md|txt|docx|pdf|xlsx)
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 平台cos路径，与DescribeStorageCredential接口查询UploadPath参数保持一致
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

	// 文档操作类型：1：批量导入（批量导入问答对）；2:文档导入（正常导入单个文档）
	Opt *uint64 `json:"Opt,omitnil,omitempty" name:"Opt"`
}

type SaveDocRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文件名
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文件类型(md|txt|docx|pdf|xlsx)
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 平台cos路径，与DescribeStorageCredential接口查询UploadPath参数保持一致
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

	// 文档操作类型：1：批量导入（批量导入问答对）；2:文档导入（正常导入单个文档）
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

type SimilarQuestion struct {
	// 相似问ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SimBizId *string `json:"SimBizId,omitnil,omitempty" name:"SimBizId"`

	// 相似问内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`
}

type SimilarQuestionModify struct {
	// 需要添加的相似问(内容)列表
	AddQuestions []*string `json:"AddQuestions,omitnil,omitempty" name:"AddQuestions"`

	// 需要更新的相似问列表
	UpdateQuestions []*SimilarQuestion `json:"UpdateQuestions,omitnil,omitempty" name:"UpdateQuestions"`

	// 需要删除的相似问列表
	DeleteQuestions []*SimilarQuestion `json:"DeleteQuestions,omitnil,omitempty" name:"DeleteQuestions"`
}

type Stat struct {
	// x轴时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	X *string `json:"X,omitnil,omitempty" name:"X"`

	// y轴统计值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Y *float64 `json:"Y,omitnil,omitempty" name:"Y"`
}

// Predefined struct for user
type StopDocParseRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`
}

type StopDocParseRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
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

type StrValue struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type SummaryConfig struct {
	// 模型配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Model *AppModel `json:"Model,omitnil,omitempty" name:"Model"`

	// 知识摘要输出配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *SummaryOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// 欢迎语，200字符以内
	// 注意：此字段可能返回 null，表示取不到有效值。
	Greeting *string `json:"Greeting,omitnil,omitempty" name:"Greeting"`
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

type TaskFlowInfo struct {
	// 任务流程ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskFlowId *string `json:"TaskFlowId,omitnil,omitempty" name:"TaskFlowId"`

	// 任务流程名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskFlowName *string `json:"TaskFlowName,omitnil,omitempty" name:"TaskFlowName"`

	// Query 重写结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueryRewrite *string `json:"QueryRewrite,omitnil,omitempty" name:"QueryRewrite"`

	// 命中意图
	// 注意：此字段可能返回 null，表示取不到有效值。
	HitIntent *string `json:"HitIntent,omitnil,omitempty" name:"HitIntent"`

	// 任务流程回复类型
	// 0: 任务流回复
	// 1: 任务流静默
	// 2: 任务流拉回话术
	// 3: 任务流自定义回复
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`
}

type TaskFlowSummary struct {
	// 任务流程名
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntentName *string `json:"IntentName,omitnil,omitempty" name:"IntentName"`

	// 实体列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedSlotValues []*ValueInfo `json:"UpdatedSlotValues,omitnil,omitempty" name:"UpdatedSlotValues"`

	// 节点列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunNodes []*RunNodeInfo `json:"RunNodes,omitnil,omitempty" name:"RunNodes"`

	// 意图判断
	// 注意：此字段可能返回 null，表示取不到有效值。
	Purposes []*string `json:"Purposes,omitnil,omitempty" name:"Purposes"`
}

type TaskParams struct {
	// 下载地址,需要通过cos桶临时密钥去下载
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosPath *string `json:"CosPath,omitnil,omitempty" name:"CosPath"`
}

type TokenStat struct {
	// 会话 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 请求 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`

	// 对应哪条会话, 会话 ID, 用于回答的消息存储使用, 可提前生成, 保存消息时使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// token 已使用数
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsedCount *uint64 `json:"UsedCount,omitnil,omitempty" name:"UsedCount"`

	// 免费 token 数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FreeCount *uint64 `json:"FreeCount,omitnil,omitempty" name:"FreeCount"`

	// 订单总 token 数
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderCount *uint64 `json:"OrderCount,omitnil,omitempty" name:"OrderCount"`

	// 当前执行状态汇总, 常量: 使用中: processing, 成功: success, 失败: failed
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusSummary *string `json:"StatusSummary,omitnil,omitempty" name:"StatusSummary"`

	// 当前执行状态汇总后中文展示
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusSummaryTitle *string `json:"StatusSummaryTitle,omitnil,omitempty" name:"StatusSummaryTitle"`

	// 当前请求执行时间, 单位 ms
	// 注意：此字段可能返回 null，表示取不到有效值。
	Elapsed *uint64 `json:"Elapsed,omitnil,omitempty" name:"Elapsed"`

	// 当前请求消耗 token 数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TokenCount *uint64 `json:"TokenCount,omitnil,omitempty" name:"TokenCount"`

	// 执行过程信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Procedures []*Procedure `json:"Procedures,omitnil,omitempty" name:"Procedures"`

	// 执行过程信息TraceId
	// 注意：此字段可能返回 null，表示取不到有效值。
	TraceId *string `json:"TraceId,omitnil,omitempty" name:"TraceId"`
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

	// 应用回复
	// 注意：此字段可能返回 null，表示取不到有效值。
	Answer *string `json:"Answer,omitnil,omitempty" name:"Answer"`

	// 错误类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reasons []*string `json:"Reasons,omitnil,omitempty" name:"Reasons"`
}

// Predefined struct for user
type UploadAttributeLabelRequestParams struct {
	// 应用ID
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
	
	// 应用ID
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

type Usage struct {
	// 文档页数
	TotalPages *int64 `json:"TotalPages,omitnil,omitempty" name:"TotalPages"`

	// 输入token数
	InputTokens *int64 `json:"InputTokens,omitnil,omitempty" name:"InputTokens"`

	// 输出token数
	OutputTokens *int64 `json:"OutputTokens,omitnil,omitempty" name:"OutputTokens"`

	// 总token数
	TotalTokens *int64 `json:"TotalTokens,omitnil,omitempty" name:"TotalTokens"`
}

type ValueInfo struct {
	// 值ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 值类型：0:未知或者空, 1:string, 2:int, 3:float, 4:bool, 5:array(字符串数组), 6: object_array(结构体数组), 7: object(结构体)
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValueType *int64 `json:"ValueType,omitnil,omitempty" name:"ValueType"`

	// string
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValueStr *string `json:"ValueStr,omitnil,omitempty" name:"ValueStr"`

	// int（避免精度丢失使用字符串返回）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValueInt *string `json:"ValueInt,omitnil,omitempty" name:"ValueInt"`

	// float
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValueFloat *float64 `json:"ValueFloat,omitnil,omitempty" name:"ValueFloat"`

	// bool
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValueBool *bool `json:"ValueBool,omitnil,omitempty" name:"ValueBool"`

	// array
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValueStrArray []*string `json:"ValueStrArray,omitnil,omitempty" name:"ValueStrArray"`
}

// Predefined struct for user
type VerifyQARequestParams struct {
	// 问答列表
	List []*QAList `json:"List,omitnil,omitempty" name:"List"`

	// 应用ID
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

	// 应用ID
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

type WordRecognizeInfo struct {
	// 输入文件的页码数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// word的base64
	// 注意：此字段可能返回 null，表示取不到有效值。
	WordBase64 *string `json:"WordBase64,omitnil,omitempty" name:"WordBase64"`
}