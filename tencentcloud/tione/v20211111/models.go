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

package v20211111

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AggregatePublicAlgoVersion struct {
	// 用于聚合的系列名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 算法公共版本列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicAlgoVersions []*PublicAlgoVersion `json:"PublicAlgoVersions,omitnil,omitempty" name:"PublicAlgoVersions"`
}

type AnnotationTaskInfo struct {
	// <p>标注任务id</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <p>标注任务名称</p>
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// <p>数据集id</p>
	DatasetId *string `json:"DatasetId,omitnil,omitempty" name:"DatasetId"`

	// <p>数据集名称</p>
	DatasetName *string `json:"DatasetName,omitnil,omitempty" name:"DatasetName"`

	// <p>标注场景名称</p>
	SceneName *string `json:"SceneName,omitnil,omitempty" name:"SceneName"`

	// <p>标注任务的label信息数组</p>
	LabelValueList []*LabelValue `json:"LabelValueList,omitnil,omitempty" name:"LabelValueList"`

	// <p>tag详情数组</p>
	CamTagList []*CamTag `json:"CamTagList,omitnil,omitempty" name:"CamTagList"`

	// <p>任务状态</p>
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>创建任务失败原因说明</p>
	AbnormalMsg *string `json:"AbnormalMsg,omitnil,omitempty" name:"AbnormalMsg"`

	// <p>标注任务是否正在提交</p>
	IsSubmitting *bool `json:"IsSubmitting,omitnil,omitempty" name:"IsSubmitting"`

	// <p>任务详情描述</p>
	TaskNote *string `json:"TaskNote,omitnil,omitempty" name:"TaskNote"`

	// <p>数据集版本</p>
	DataSetVersion *string `json:"DataSetVersion,omitnil,omitempty" name:"DataSetVersion"`

	// <p>已经标注的图片数量</p>
	NumAnnotated *uint64 `json:"NumAnnotated,omitnil,omitempty" name:"NumAnnotated"`

	// <p>标注的总图片数量</p>
	NumTotal *uint64 `json:"NumTotal,omitnil,omitempty" name:"NumTotal"`

	// <p>创建任务的时间戳</p>
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>Ocr Tool Type</p>
	OcrToolType *uint64 `json:"OcrToolType,omitnil,omitempty" name:"OcrToolType"`

	// <p>Ocr Text Attribute Annotate Enable</p>
	OcrTextAttributeAnnotateEnable *bool `json:"OcrTextAttributeAnnotateEnable,omitnil,omitempty" name:"OcrTextAttributeAnnotateEnable"`

	// <p>导出格式</p>
	ExportFormat *string `json:"ExportFormat,omitnil,omitempty" name:"ExportFormat"`

	// <p>提交错误说明</p>
	SubmittingErrorMsg *string `json:"SubmittingErrorMsg,omitnil,omitempty" name:"SubmittingErrorMsg"`

	// <p>ocr任务类型：1-识别。2-智能结构化</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrAnnotationContentType *uint64 `json:"OcrAnnotationContentType,omitnil,omitempty" name:"OcrAnnotationContentType"`

	// <p>OCR任务：是否启用辅助标注</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableAuxiliaryAnnotation *bool `json:"EnableAuxiliaryAnnotation,omitnil,omitempty" name:"EnableAuxiliaryAnnotation"`

	// <p>数据集创建者UIN</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetCreator *string `json:"DatasetCreator,omitnil,omitempty" name:"DatasetCreator"`

	// <p>任务创建者UIN</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// <p>数据集创建者名称</p>
	DatasetCreatorName *string `json:"DatasetCreatorName,omitnil,omitempty" name:"DatasetCreatorName"`

	// <p>任务创建者名称</p>
	CreatorName *string `json:"CreatorName,omitnil,omitempty" name:"CreatorName"`

	// <p>标注任务状态枚举</p><p>枚举值：</p><ul><li>CREATING： 创建中</li><li>CREATE_SUCCESS： 创建成功，可标注</li><li>CREATE_FAILED： 创建失败</li><li>SUBMITTING： 提交中</li><li>SUBMIT_SUCCESS： 提交成功，需重启才可标注</li><li>SUBMIT_FAILED： 提交失败</li><li>ABNORMAL： 数据版本异常，需删除重建（大模型场景）</li></ul>
	TaskStatus *string `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`
}

type Attribute struct {
	// 为‘List’时属性值取Values 否则取Value
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 属性key
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 属性值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 属性值列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type AuthToken struct {
	// AuthToken 基础信息
	Base *AuthTokenBase `json:"Base,omitnil,omitempty" name:"Base"`

	// AuthToken 限流数组
	Limits []*AuthTokenLimit `json:"Limits,omitnil,omitempty" name:"Limits"`
}

type AuthTokenBase struct {
	// token 值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// token 别名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// token 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// token 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// token状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// token的唯一id，与value一一对应，重置后id也会一并变化
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type AuthTokenLimit struct {
	// 限频策略：PerMinute 每分钟限频；PerDay 每日限频
	Strategy *string `json:"Strategy,omitnil,omitempty" name:"Strategy"`

	// 上限值
	Max *uint64 `json:"Max,omitnil,omitempty" name:"Max"`
}

type CBSConfig struct {
	// 存储大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	VolumeSizeInGB *int64 `json:"VolumeSizeInGB,omitnil,omitempty" name:"VolumeSizeInGB"`
}

type CFSConfig struct {
	// cfs的实例的ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 存储的路径
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// cfs的挂载类型，可选值为：STORAGE、SOURCE 分别表示存储拓展模式和数据源模式，默认为 STORAGE
	// 注意：此字段可能返回 null，表示取不到有效值。
	MountType *string `json:"MountType,omitnil,omitempty" name:"MountType"`

	// 协议 1: NFS, 2: TURBO
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`
}

type CFSTurbo struct {
	// CFSTurbo实例id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// CFSTurbo路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`
}

type CamTag struct {
	// tag键值
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// tag值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type ChatCompletionRequestParams struct {
	// 对话的目标模型ID。
	// 自行部署的开源大模型聊天：部署的模型服务组ID，形如ms-q7pfr29p。
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 输入对话历史。旧的对话在前，数组中最后一项应该为这次的问题。
	Messages []*Message `json:"Messages,omitnil,omitempty" name:"Messages"`

	// 仅当模型为自行部署的开源大模型时生效。采样随机值，默认值为0.7，取值范围[0,2]。较高的值(如0.8)将使输出更加随机，而较低的值(如0.2)将使输出更加确定。建议仅修改此参数或TopP，但不建议两者都修改。
	Temperature *float64 `json:"Temperature,omitnil,omitempty" name:"Temperature"`

	// 仅当模型为自行部署的开源大模型时生效。核采样，默认值为1，取值范围[0,1]。指的是预先设置一个概率界限 p，然后将所有可能生成的token，根据概率大小从高到低排列，依次选取。当这些选取的token的累积概率大于或等于 p 值时停止，然后从已经选取的token中进行采样，生成下一个token。例如top_p为0.1时意味着模型只考虑累积概率为10%的token。建议仅修改此参数或Temperature，不建议两者都修改。
	TopP *float64 `json:"TopP,omitnil,omitempty" name:"TopP"`

	// 仅当模型为自行部署的开源大模型时生效。默认 512，模型可生成内容的最长 token 数量，最大不能超过模型支持的上下文长度。
	MaxTokens *int64 `json:"MaxTokens,omitnil,omitempty" name:"MaxTokens"`
}

type ChatCompletionRequest struct {
	*tchttp.BaseRequest
	
	// 对话的目标模型ID。
	// 自行部署的开源大模型聊天：部署的模型服务组ID，形如ms-q7pfr29p。
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 输入对话历史。旧的对话在前，数组中最后一项应该为这次的问题。
	Messages []*Message `json:"Messages,omitnil,omitempty" name:"Messages"`

	// 仅当模型为自行部署的开源大模型时生效。采样随机值，默认值为0.7，取值范围[0,2]。较高的值(如0.8)将使输出更加随机，而较低的值(如0.2)将使输出更加确定。建议仅修改此参数或TopP，但不建议两者都修改。
	Temperature *float64 `json:"Temperature,omitnil,omitempty" name:"Temperature"`

	// 仅当模型为自行部署的开源大模型时生效。核采样，默认值为1，取值范围[0,1]。指的是预先设置一个概率界限 p，然后将所有可能生成的token，根据概率大小从高到低排列，依次选取。当这些选取的token的累积概率大于或等于 p 值时停止，然后从已经选取的token中进行采样，生成下一个token。例如top_p为0.1时意味着模型只考虑累积概率为10%的token。建议仅修改此参数或Temperature，不建议两者都修改。
	TopP *float64 `json:"TopP,omitnil,omitempty" name:"TopP"`

	// 仅当模型为自行部署的开源大模型时生效。默认 512，模型可生成内容的最长 token 数量，最大不能超过模型支持的上下文长度。
	MaxTokens *int64 `json:"MaxTokens,omitnil,omitempty" name:"MaxTokens"`
}

func (r *ChatCompletionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChatCompletionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Model")
	delete(f, "Messages")
	delete(f, "Temperature")
	delete(f, "TopP")
	delete(f, "MaxTokens")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChatCompletionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChatCompletionResponseParams struct {
	// 对话的模型服务组ID
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 本次问答的答案。
	Choices []*Choice `json:"Choices,omitnil,omitempty" name:"Choices"`

	// 会话Id。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// token统计
	Usage *Usage `json:"Usage,omitnil,omitempty" name:"Usage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChatCompletionResponse struct {
	*tchttp.BaseResponse
	Response *ChatCompletionResponseParams `json:"Response"`
}

func (r *ChatCompletionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChatCompletionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Choice struct {
	// 对话结果
	Message *Message `json:"Message,omitnil,omitempty" name:"Message"`

	// 结束理由: stop, length, content_filter, null
	FinishReason *string `json:"FinishReason,omitnil,omitempty" name:"FinishReason"`

	// 序号
	Index *int64 `json:"Index,omitnil,omitempty" name:"Index"`
}

type CodeRepoConfig struct {
	// 代码仓库Id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 代码仓下载目标地址
	TargetPath *string `json:"TargetPath,omitnil,omitempty" name:"TargetPath"`
}

type Container struct {
	// 名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerId *string `json:"ContainerId,omitnil,omitempty" name:"ContainerId"`

	// 镜像地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// 容器状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *ContainerStatus `json:"Status,omitnil,omitempty" name:"Status"`
}

type ContainerStatus struct {
	// 重启次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RestartCount *int64 `json:"RestartCount,omitnil,omitempty" name:"RestartCount"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 是否就绪
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ready *bool `json:"Ready,omitnil,omitempty" name:"Ready"`

	// 状态原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 容器的错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

type CosPathInfo struct {
	// 存储桶
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// 所在地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 路径列表，目前只支持单个
	// 注意：此字段可能返回 null，表示取不到有效值。
	Paths []*string `json:"Paths,omitnil,omitempty" name:"Paths"`
}

// Predefined struct for user
type CreateDataSourceRequestParams struct {
	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`

	// 数据源名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数据源类型英文名
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 数据源权限，取值有RW RO
	Permission *string `json:"Permission,omitnil,omitempty" name:"Permission"`

	// 存储实例ID
	StorageId *string `json:"StorageId,omitnil,omitempty" name:"StorageId"`

	// 数据源挂载配置
	MountConfigure *MountConfigureInfo `json:"MountConfigure,omitnil,omitempty" name:"MountConfigure"`

	// 标签配置
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateDataSourceRequest struct {
	*tchttp.BaseRequest
	
	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`

	// 数据源名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数据源类型英文名
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 数据源权限，取值有RW RO
	Permission *string `json:"Permission,omitnil,omitempty" name:"Permission"`

	// 存储实例ID
	StorageId *string `json:"StorageId,omitnil,omitempty" name:"StorageId"`

	// 数据源挂载配置
	MountConfigure *MountConfigureInfo `json:"MountConfigure,omitnil,omitempty" name:"MountConfigure"`

	// 标签配置
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *CreateDataSourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDataSourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TiProjectId")
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "Permission")
	delete(f, "StorageId")
	delete(f, "MountConfigure")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDataSourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDataSourceResponseParams struct {
	// 数据源ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDataSourceResponse struct {
	*tchttp.BaseResponse
	Response *CreateDataSourceResponseParams `json:"Response"`
}

func (r *CreateDataSourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDataSourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDatasetRequestParams struct {
	// 数据集名称，不超过60个字符，仅支持中英文、数字、下划线"_"、短横"-"，只能以中英文、数字开头
	DatasetName *string `json:"DatasetName,omitnil,omitempty" name:"DatasetName"`

	// 数据集类型:
	// TYPE_DATASET_TEXT，文本
	// TYPE_DATASET_IMAGE，图片
	// TYPE_DATASET_TABLE，表格
	// TYPE_DATASET_OTHER，其他
	DatasetType *string `json:"DatasetType,omitnil,omitempty" name:"DatasetType"`

	// 数据源cos路径
	StorageDataPath *CosPathInfo `json:"StorageDataPath,omitnil,omitempty" name:"StorageDataPath"`

	// 数据集标签cos存储路径
	StorageLabelPath *CosPathInfo `json:"StorageLabelPath,omitnil,omitempty" name:"StorageLabelPath"`

	// 数据集标签
	DatasetTags []*Tag `json:"DatasetTags,omitnil,omitempty" name:"DatasetTags"`

	// 数据集标注状态:
	// STATUS_NON_ANNOTATED，未标注
	// STATUS_ANNOTATED，已标注
	AnnotationStatus *string `json:"AnnotationStatus,omitnil,omitempty" name:"AnnotationStatus"`

	// 标注类型:
	// ANNOTATION_TYPE_CLASSIFICATION，图片分类
	// ANNOTATION_TYPE_DETECTION，目标检测
	// ANNOTATION_TYPE_SEGMENTATION，图片分割
	// ANNOTATION_TYPE_TRACKING，目标跟踪
	// ANNOTATION_TYPE_OCR，OCR
	// ANNOTATION_TYPE_TEXT_CLASSIFICATION，文本分类
	AnnotationType *string `json:"AnnotationType,omitnil,omitempty" name:"AnnotationType"`

	// 标注格式:
	// ANNOTATION_FORMAT_TI，TI平台格式
	// ANNOTATION_FORMAT_PASCAL，Pascal Voc
	// ANNOTATION_FORMAT_COCO，COCO
	// ANNOTATION_FORMAT_FILE，文件目录结构
	// ANNOTATION_FORMAT_TEXT_TI，文本类型TI平台格式
	// ANNOTATION_FORMAT_TXT，文本类型TXT格式
	// ANNOTATION_FORMAT_CSV，文本类型CSV格式
	// ANNOTATION_FORMAT_JSON，文本类型JSON格式
	AnnotationFormat *string `json:"AnnotationFormat,omitnil,omitempty" name:"AnnotationFormat"`

	// 表头信息
	SchemaInfos []*SchemaInfo `json:"SchemaInfos,omitnil,omitempty" name:"SchemaInfos"`

	// 数据是否存在表头
	IsSchemaExisted *bool `json:"IsSchemaExisted,omitnil,omitempty" name:"IsSchemaExisted"`

	// 导入文件粒度
	// TYPE_TEXT_LINE，按行
	// TYPE_TEXT_FILE，按文件
	ContentType *string `json:"ContentType,omitnil,omitempty" name:"ContentType"`

	// 数据集建模一级类别。LLM,CV,STRUCTURE,OTHER
	DatasetScene *string `json:"DatasetScene,omitnil,omitempty" name:"DatasetScene"`

	// 数据集标签。
	SceneTags []*string `json:"SceneTags,omitnil,omitempty" name:"SceneTags"`

	// 数据集CFS配置。仅支持LLM场景
	CFSConfig *CFSConfig `json:"CFSConfig,omitnil,omitempty" name:"CFSConfig"`
}

type CreateDatasetRequest struct {
	*tchttp.BaseRequest
	
	// 数据集名称，不超过60个字符，仅支持中英文、数字、下划线"_"、短横"-"，只能以中英文、数字开头
	DatasetName *string `json:"DatasetName,omitnil,omitempty" name:"DatasetName"`

	// 数据集类型:
	// TYPE_DATASET_TEXT，文本
	// TYPE_DATASET_IMAGE，图片
	// TYPE_DATASET_TABLE，表格
	// TYPE_DATASET_OTHER，其他
	DatasetType *string `json:"DatasetType,omitnil,omitempty" name:"DatasetType"`

	// 数据源cos路径
	StorageDataPath *CosPathInfo `json:"StorageDataPath,omitnil,omitempty" name:"StorageDataPath"`

	// 数据集标签cos存储路径
	StorageLabelPath *CosPathInfo `json:"StorageLabelPath,omitnil,omitempty" name:"StorageLabelPath"`

	// 数据集标签
	DatasetTags []*Tag `json:"DatasetTags,omitnil,omitempty" name:"DatasetTags"`

	// 数据集标注状态:
	// STATUS_NON_ANNOTATED，未标注
	// STATUS_ANNOTATED，已标注
	AnnotationStatus *string `json:"AnnotationStatus,omitnil,omitempty" name:"AnnotationStatus"`

	// 标注类型:
	// ANNOTATION_TYPE_CLASSIFICATION，图片分类
	// ANNOTATION_TYPE_DETECTION，目标检测
	// ANNOTATION_TYPE_SEGMENTATION，图片分割
	// ANNOTATION_TYPE_TRACKING，目标跟踪
	// ANNOTATION_TYPE_OCR，OCR
	// ANNOTATION_TYPE_TEXT_CLASSIFICATION，文本分类
	AnnotationType *string `json:"AnnotationType,omitnil,omitempty" name:"AnnotationType"`

	// 标注格式:
	// ANNOTATION_FORMAT_TI，TI平台格式
	// ANNOTATION_FORMAT_PASCAL，Pascal Voc
	// ANNOTATION_FORMAT_COCO，COCO
	// ANNOTATION_FORMAT_FILE，文件目录结构
	// ANNOTATION_FORMAT_TEXT_TI，文本类型TI平台格式
	// ANNOTATION_FORMAT_TXT，文本类型TXT格式
	// ANNOTATION_FORMAT_CSV，文本类型CSV格式
	// ANNOTATION_FORMAT_JSON，文本类型JSON格式
	AnnotationFormat *string `json:"AnnotationFormat,omitnil,omitempty" name:"AnnotationFormat"`

	// 表头信息
	SchemaInfos []*SchemaInfo `json:"SchemaInfos,omitnil,omitempty" name:"SchemaInfos"`

	// 数据是否存在表头
	IsSchemaExisted *bool `json:"IsSchemaExisted,omitnil,omitempty" name:"IsSchemaExisted"`

	// 导入文件粒度
	// TYPE_TEXT_LINE，按行
	// TYPE_TEXT_FILE，按文件
	ContentType *string `json:"ContentType,omitnil,omitempty" name:"ContentType"`

	// 数据集建模一级类别。LLM,CV,STRUCTURE,OTHER
	DatasetScene *string `json:"DatasetScene,omitnil,omitempty" name:"DatasetScene"`

	// 数据集标签。
	SceneTags []*string `json:"SceneTags,omitnil,omitempty" name:"SceneTags"`

	// 数据集CFS配置。仅支持LLM场景
	CFSConfig *CFSConfig `json:"CFSConfig,omitnil,omitempty" name:"CFSConfig"`
}

func (r *CreateDatasetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDatasetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatasetName")
	delete(f, "DatasetType")
	delete(f, "StorageDataPath")
	delete(f, "StorageLabelPath")
	delete(f, "DatasetTags")
	delete(f, "AnnotationStatus")
	delete(f, "AnnotationType")
	delete(f, "AnnotationFormat")
	delete(f, "SchemaInfos")
	delete(f, "IsSchemaExisted")
	delete(f, "ContentType")
	delete(f, "DatasetScene")
	delete(f, "SceneTags")
	delete(f, "CFSConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDatasetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDatasetResponseParams struct {
	// 数据集ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetId *string `json:"DatasetId,omitnil,omitempty" name:"DatasetId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDatasetResponse struct {
	*tchttp.BaseResponse
	Response *CreateDatasetResponseParams `json:"Response"`
}

func (r *CreateDatasetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDatasetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateExportRequestParams struct {
	// 服务类型，TRAIN为任务式建模, NOTEBOOK为Notebook, INFER为在线服务, BATCH为批量预测枚举值：- TRAIN- NOTEBOOK- INFER- BATCH
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// 服务ID，和Service参数对应，不同Service的服务ID获取方式不同，具体如下：- Service类型为TRAIN：  调用[DescribeTrainingTask接口](/document/product/851/75089)查询训练任务详情，ServiceId为接口返回值中Response.TrainingTaskDetail.LatestInstanceId- Service类型为NOTEBOOK：  调用[DescribeNotebook接口](/document/product/851/95662)查询Notebook详情，ServiceId为接口返回值中Response.NotebookDetail.PodName- Service类型为INFER：  调用[DescribeModelServiceGroup接口](/document/product/851/82285)查询服务组详情，ServiceId为接口返回值中Response.ServiceGroup.Services.ServiceId- Service类型为BATCH：  调用[DescribeBatchTask接口](/document/product/851/80180)查询跑批任务详情，ServiceId为接口返回值中Response.BatchTaskDetail.LatestInstanceId
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 日志查询开始时间（RFC3339格式的时间字符串），默认值为当前时间的前一个小时
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 日志查询结束时间（RFC3339格式的时间字符串），开始时间和结束时间必须同时填或同时不填，默认值为当前时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 日志导出数据格式。json，csv，默认为csv
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// Pod的名称，即需要查询服务对应的Pod，和Service参数对应，不同Service的PodName获取方式不同，具体如下：- Service类型为TRAIN：  调用[DescribeTrainingTaskPods接口](/document/product/851/75088)查询训练任务pod列表，PodName为接口返回值中Response.PodNames- Service类型为NOTEBOOK：  调用[DescribeNotebook接口](/document/product/851/95662)查询Notebook详情，PodName为接口返回值中Response.NotebookDetail.PodName- Service类型为INFER：  调用[DescribeModelService接口](/document/product/851/82287)查询单个服务详情，PodName为接口返回值中Response.Service.ServiceInfo.PodInfos- Service类型为BATCH：  调用[DescribeBatchTask接口](/document/product/851/80180)查询跑批任务详情，PodName为接口返回值中Response.BatchTaskDetail. PodList注：支持结尾通配符*
	PodName *string `json:"PodName,omitnil,omitempty" name:"PodName"`

	// 描述任务的类型
	JobCategory *string `json:"JobCategory,omitnil,omitempty" name:"JobCategory"`

	// 实例的类型
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 查实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 日志类型： PLATFORM_INIT, PLATFORM_SANITY_CHECK, USER
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type CreateExportRequest struct {
	*tchttp.BaseRequest
	
	// 服务类型，TRAIN为任务式建模, NOTEBOOK为Notebook, INFER为在线服务, BATCH为批量预测枚举值：- TRAIN- NOTEBOOK- INFER- BATCH
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// 服务ID，和Service参数对应，不同Service的服务ID获取方式不同，具体如下：- Service类型为TRAIN：  调用[DescribeTrainingTask接口](/document/product/851/75089)查询训练任务详情，ServiceId为接口返回值中Response.TrainingTaskDetail.LatestInstanceId- Service类型为NOTEBOOK：  调用[DescribeNotebook接口](/document/product/851/95662)查询Notebook详情，ServiceId为接口返回值中Response.NotebookDetail.PodName- Service类型为INFER：  调用[DescribeModelServiceGroup接口](/document/product/851/82285)查询服务组详情，ServiceId为接口返回值中Response.ServiceGroup.Services.ServiceId- Service类型为BATCH：  调用[DescribeBatchTask接口](/document/product/851/80180)查询跑批任务详情，ServiceId为接口返回值中Response.BatchTaskDetail.LatestInstanceId
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 日志查询开始时间（RFC3339格式的时间字符串），默认值为当前时间的前一个小时
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 日志查询结束时间（RFC3339格式的时间字符串），开始时间和结束时间必须同时填或同时不填，默认值为当前时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 日志导出数据格式。json，csv，默认为csv
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// Pod的名称，即需要查询服务对应的Pod，和Service参数对应，不同Service的PodName获取方式不同，具体如下：- Service类型为TRAIN：  调用[DescribeTrainingTaskPods接口](/document/product/851/75088)查询训练任务pod列表，PodName为接口返回值中Response.PodNames- Service类型为NOTEBOOK：  调用[DescribeNotebook接口](/document/product/851/95662)查询Notebook详情，PodName为接口返回值中Response.NotebookDetail.PodName- Service类型为INFER：  调用[DescribeModelService接口](/document/product/851/82287)查询单个服务详情，PodName为接口返回值中Response.Service.ServiceInfo.PodInfos- Service类型为BATCH：  调用[DescribeBatchTask接口](/document/product/851/80180)查询跑批任务详情，PodName为接口返回值中Response.BatchTaskDetail. PodList注：支持结尾通配符*
	PodName *string `json:"PodName,omitnil,omitempty" name:"PodName"`

	// 描述任务的类型
	JobCategory *string `json:"JobCategory,omitnil,omitempty" name:"JobCategory"`

	// 实例的类型
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 查实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 日志类型： PLATFORM_INIT, PLATFORM_SANITY_CHECK, USER
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *CreateExportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateExportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Service")
	delete(f, "ServiceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Format")
	delete(f, "PodName")
	delete(f, "JobCategory")
	delete(f, "InstanceType")
	delete(f, "InstanceId")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateExportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateExportResponseParams struct {
	// 日志下载任务的ID
	ExportId *string `json:"ExportId,omitnil,omitempty" name:"ExportId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateExportResponse struct {
	*tchttp.BaseResponse
	Response *CreateExportResponseParams `json:"Response"`
}

func (r *CreateExportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateModelServiceAuthTokenRequestParams struct {
	// 服务组 id
	ServiceGroupId *string `json:"ServiceGroupId,omitnil,omitempty" name:"ServiceGroupId"`

	// token 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Description 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateModelServiceAuthTokenRequest struct {
	*tchttp.BaseRequest
	
	// 服务组 id
	ServiceGroupId *string `json:"ServiceGroupId,omitnil,omitempty" name:"ServiceGroupId"`

	// token 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Description 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateModelServiceAuthTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateModelServiceAuthTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceGroupId")
	delete(f, "Name")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateModelServiceAuthTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateModelServiceAuthTokenResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateModelServiceAuthTokenResponse struct {
	*tchttp.BaseResponse
	Response *CreateModelServiceAuthTokenResponseParams `json:"Response"`
}

func (r *CreateModelServiceAuthTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateModelServiceAuthTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateModelServiceRequestParams struct {
	// <p>新增版本时需要填写</p>
	ServiceGroupId *string `json:"ServiceGroupId,omitnil,omitempty" name:"ServiceGroupId"`

	// <p>不超过60个字，仅支持英文、数字、下划线&quot;_&quot;、短横&quot;-&quot;，只能以英文、数字开头</p>
	ServiceGroupName *string `json:"ServiceGroupName,omitnil,omitempty" name:"ServiceGroupName"`

	// <p>模型服务的描述</p>
	ServiceDescription *string `json:"ServiceDescription,omitnil,omitempty" name:"ServiceDescription"`

	// <p>付费模式,有 PREPAID （包年包月）和 POSTPAID_BY_HOUR（按量付费）</p>
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// <p>预付费模式下所属的资源组id，同服务组下唯一</p>
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// <p>模型信息，需要挂载模型时填写</p>
	ModelInfo *ModelInfo `json:"ModelInfo,omitnil,omitempty" name:"ModelInfo"`

	// <p>镜像信息，配置服务运行所需的镜像地址等信息</p>
	ImageInfo *ImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`

	// <p>环境变量，可选参数，用于配置容器中的环境变量</p>
	Env []*EnvVar `json:"Env,omitnil,omitempty" name:"Env"`

	// <p>资源描述，指定包年包月模式下的cpu,mem,gpu等信息，后付费无需填写</p>
	Resources *ResourceInfo `json:"Resources,omitnil,omitempty" name:"Resources"`

	// <p>使用DescribeBillingSpecs接口返回的规格列表中的值，或者参考实例列表:<br>TI.S.MEDIUM.POST    2C4G<br>TI.S.LARGE.POST    4C8G<br>TI.S.2XLARGE16.POST    8C16G<br>TI.S.2XLARGE32.POST    8C32G<br>TI.S.4XLARGE32.POST    16C32G<br>TI.S.4XLARGE64.POST    16C64G<br>TI.S.6XLARGE48.POST    24C48G<br>TI.S.6XLARGE96.POST    24C96G<br>TI.S.8XLARGE64.POST    32C64G<br>TI.S.8XLARGE128.POST 32C128G<br>TI.GN7.LARGE20.POST    4C20G T4<em>1/4<br>TI.GN7.2XLARGE40.POST    10C40G T4</em>1/2<br>TI.GN7.2XLARGE32.POST    8C32G T4<em>1<br>TI.GN7.5XLARGE80.POST    20C80G T4</em>1<br>TI.GN7.8XLARGE128.POST    32C128G T4<em>1<br>TI.GN7.10XLARGE160.POST    40C160G T4</em>2<br>TI.GN7.20XLARGE320.POST    80C320G T4*4</p>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// <p>扩缩容类型 支持：自动 - &quot;AUTO&quot;, 手动 - &quot;MANUAL&quot;,默认为MANUAL</p>
	ScaleMode *string `json:"ScaleMode,omitnil,omitempty" name:"ScaleMode"`

	// <p>实例数量, 不同计费模式和调节模式下对应关系如下<br>PREPAID 和 POSTPAID_BY_HOUR:<br>手动调节模式下对应 实例数量<br>自动调节模式下对应 基于时间的默认策略的实例数量<br>HYBRID_PAID:<br>后付费实例手动调节模式下对应 实例数量<br>后付费实例自动调节模式下对应 时间策略的默认策略的实例数量</p>
	Replicas *int64 `json:"Replicas,omitnil,omitempty" name:"Replicas"`

	// <p>自动伸缩信息</p>
	HorizontalPodAutoscaler *HorizontalPodAutoscaler `json:"HorizontalPodAutoscaler,omitnil,omitempty" name:"HorizontalPodAutoscaler"`

	// <p>是否开启日志投递，开启后需填写配置投递到指定cls</p>
	LogEnable *bool `json:"LogEnable,omitnil,omitempty" name:"LogEnable"`

	// <p>日志配置，需要投递服务日志到指定cls时填写</p>
	LogConfig *LogConfig `json:"LogConfig,omitnil,omitempty" name:"LogConfig"`

	// <p>是否开启接口鉴权，开启后自动生成token信息，访问需要token鉴权</p>
	AuthorizationEnable *bool `json:"AuthorizationEnable,omitnil,omitempty" name:"AuthorizationEnable"`

	// <p>腾讯云标签</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>是否新增版本</p>
	NewVersion *bool `json:"NewVersion,omitnil,omitempty" name:"NewVersion"`

	// <p>定时任务配置，使用定时策略时填写</p>
	CronScaleJobs []*CronScaleJob `json:"CronScaleJobs,omitnil,omitempty" name:"CronScaleJobs"`

	// <p>自动伸缩策略配置 HPA : 通过HPA进行弹性伸缩 CRON 通过定时任务进行伸缩</p>
	ScaleStrategy *string `json:"ScaleStrategy,omitnil,omitempty" name:"ScaleStrategy"`

	// <p>计费模式[HYBRID_PAID]时生效, 用于标识混合计费模式下的预付费实例数</p>
	HybridBillingPrepaidReplicas *int64 `json:"HybridBillingPrepaidReplicas,omitnil,omitempty" name:"HybridBillingPrepaidReplicas"`

	// <p>[AUTO_ML 自动学习，自动学习正式发布 AUTO_ML_FORMAL, DEFAULT 默认]</p>
	CreateSource *string `json:"CreateSource,omitnil,omitempty" name:"CreateSource"`

	// <p>是否开启模型的热更新。默认不开启</p>
	ModelHotUpdateEnable *bool `json:"ModelHotUpdateEnable,omitnil,omitempty" name:"ModelHotUpdateEnable"`

	// <p>定时停止配置</p>
	ScheduledAction *ScheduledAction `json:"ScheduledAction,omitnil,omitempty" name:"ScheduledAction"`

	// <p>挂载配置，目前只支持CFS</p>
	VolumeMount *VolumeMount `json:"VolumeMount,omitnil,omitempty" name:"VolumeMount"`

	// <p>服务限速限流相关配置</p>
	ServiceLimit *ServiceLimit `json:"ServiceLimit,omitnil,omitempty" name:"ServiceLimit"`

	// <p>回调地址，用于回调创建服务状态信息，回调格式&amp;内容详情见：<a href="https://cloud.tencent.com/document/product/851/84292">TI-ONE 接口回调说明</a></p>
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// <p>是否开启模型的加速, 仅对StableDiffusion(动态加速)格式的模型有效。</p>
	ModelTurboEnable *bool `json:"ModelTurboEnable,omitnil,omitempty" name:"ModelTurboEnable"`

	// <p>服务分类</p>
	ServiceCategory *string `json:"ServiceCategory,omitnil,omitempty" name:"ServiceCategory"`

	// <p>服务的启动命令，如遇特殊字符导致配置失败，可使用CommandBase64参数</p>
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// <p>是否开启TIONE内网访问外部，此功能仅支持后付费机型与从TIONE平台购买的预付费机型；使用从CVM选择资源组时此配置不生效。</p>
	ServiceEIP *ServiceEIP `json:"ServiceEIP,omitnil,omitempty" name:"ServiceEIP"`

	// <p>服务的启动命令，以base64格式进行输入，与Command同时配置时，仅当前参数生效</p>
	CommandBase64 *string `json:"CommandBase64,omitnil,omitempty" name:"CommandBase64"`

	// <p>服务端口，仅在非内置镜像时生效，默认8501。不支持输入8501-8510,6006,9092</p>
	ServicePort *int64 `json:"ServicePort,omitnil,omitempty" name:"ServicePort"`

	// <p>服务的部署类型 [标准部署，分布式多机部署，] 默认STANDARD</p><p>枚举值：</p><ul><li>STANDARD： 标准部署</li><li>DIST： 多机分布式部署</li><li>ROLE_SET： 多角色部署</li></ul>
	DeployType *string `json:"DeployType,omitnil,omitempty" name:"DeployType"`

	// <p>单副本下的实例数，仅在部署类型为DIST时生效，默认1</p>
	InstancePerReplicas *int64 `json:"InstancePerReplicas,omitnil,omitempty" name:"InstancePerReplicas"`

	// <p>服务的优雅退出时限。单位为秒，默认值为30，最小为1</p>
	TerminationGracePeriodSeconds *int64 `json:"TerminationGracePeriodSeconds,omitnil,omitempty" name:"TerminationGracePeriodSeconds"`

	// <p>服务实例停止前执行的命令，执行完毕或执行时间超过优雅退出时限后实例结束</p>
	PreStopCommand []*string `json:"PreStopCommand,omitnil,omitempty" name:"PreStopCommand"`

	// <p>是否启用 grpc 端口</p>
	GrpcEnable *bool `json:"GrpcEnable,omitnil,omitempty" name:"GrpcEnable"`

	// <p>健康探针</p>
	HealthProbe *HealthProbe `json:"HealthProbe,omitnil,omitempty" name:"HealthProbe"`

	// <p>滚动更新策略</p>
	RollingUpdate *RollingUpdate `json:"RollingUpdate,omitnil,omitempty" name:"RollingUpdate"`

	// <p>sidecar配置</p>
	Sidecar *SidecarSpec `json:"Sidecar,omitnil,omitempty" name:"Sidecar"`

	// <p>数据盘批量挂载配置，当前仅支持CFS，仅针对“模型来源-腾讯云存储、模型来源-腾讯云容器镜像、模型来源-资源组、模型来源-数据源”。</p>
	VolumeMounts []*VolumeMount `json:"VolumeMounts,omitnil,omitempty" name:"VolumeMounts"`

	// <p>调度策略 [binpack] 优先占满整机，尽量避免碎卡（默认值）[spread] 优先分散在各个节点，确保服务高可用</p>
	SchedulingStrategy *string `json:"SchedulingStrategy,omitnil,omitempty" name:"SchedulingStrategy"`

	// <p>网关日志投递相关配置</p>
	GatewayLogConfig *LogConfig `json:"GatewayLogConfig,omitnil,omitempty" name:"GatewayLogConfig"`

	// <p>网关相关配置</p>
	GatewayConfig *GatewayConfig `json:"GatewayConfig,omitnil,omitempty" name:"GatewayConfig"`
}

type CreateModelServiceRequest struct {
	*tchttp.BaseRequest
	
	// <p>新增版本时需要填写</p>
	ServiceGroupId *string `json:"ServiceGroupId,omitnil,omitempty" name:"ServiceGroupId"`

	// <p>不超过60个字，仅支持英文、数字、下划线&quot;_&quot;、短横&quot;-&quot;，只能以英文、数字开头</p>
	ServiceGroupName *string `json:"ServiceGroupName,omitnil,omitempty" name:"ServiceGroupName"`

	// <p>模型服务的描述</p>
	ServiceDescription *string `json:"ServiceDescription,omitnil,omitempty" name:"ServiceDescription"`

	// <p>付费模式,有 PREPAID （包年包月）和 POSTPAID_BY_HOUR（按量付费）</p>
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// <p>预付费模式下所属的资源组id，同服务组下唯一</p>
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// <p>模型信息，需要挂载模型时填写</p>
	ModelInfo *ModelInfo `json:"ModelInfo,omitnil,omitempty" name:"ModelInfo"`

	// <p>镜像信息，配置服务运行所需的镜像地址等信息</p>
	ImageInfo *ImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`

	// <p>环境变量，可选参数，用于配置容器中的环境变量</p>
	Env []*EnvVar `json:"Env,omitnil,omitempty" name:"Env"`

	// <p>资源描述，指定包年包月模式下的cpu,mem,gpu等信息，后付费无需填写</p>
	Resources *ResourceInfo `json:"Resources,omitnil,omitempty" name:"Resources"`

	// <p>使用DescribeBillingSpecs接口返回的规格列表中的值，或者参考实例列表:<br>TI.S.MEDIUM.POST    2C4G<br>TI.S.LARGE.POST    4C8G<br>TI.S.2XLARGE16.POST    8C16G<br>TI.S.2XLARGE32.POST    8C32G<br>TI.S.4XLARGE32.POST    16C32G<br>TI.S.4XLARGE64.POST    16C64G<br>TI.S.6XLARGE48.POST    24C48G<br>TI.S.6XLARGE96.POST    24C96G<br>TI.S.8XLARGE64.POST    32C64G<br>TI.S.8XLARGE128.POST 32C128G<br>TI.GN7.LARGE20.POST    4C20G T4<em>1/4<br>TI.GN7.2XLARGE40.POST    10C40G T4</em>1/2<br>TI.GN7.2XLARGE32.POST    8C32G T4<em>1<br>TI.GN7.5XLARGE80.POST    20C80G T4</em>1<br>TI.GN7.8XLARGE128.POST    32C128G T4<em>1<br>TI.GN7.10XLARGE160.POST    40C160G T4</em>2<br>TI.GN7.20XLARGE320.POST    80C320G T4*4</p>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// <p>扩缩容类型 支持：自动 - &quot;AUTO&quot;, 手动 - &quot;MANUAL&quot;,默认为MANUAL</p>
	ScaleMode *string `json:"ScaleMode,omitnil,omitempty" name:"ScaleMode"`

	// <p>实例数量, 不同计费模式和调节模式下对应关系如下<br>PREPAID 和 POSTPAID_BY_HOUR:<br>手动调节模式下对应 实例数量<br>自动调节模式下对应 基于时间的默认策略的实例数量<br>HYBRID_PAID:<br>后付费实例手动调节模式下对应 实例数量<br>后付费实例自动调节模式下对应 时间策略的默认策略的实例数量</p>
	Replicas *int64 `json:"Replicas,omitnil,omitempty" name:"Replicas"`

	// <p>自动伸缩信息</p>
	HorizontalPodAutoscaler *HorizontalPodAutoscaler `json:"HorizontalPodAutoscaler,omitnil,omitempty" name:"HorizontalPodAutoscaler"`

	// <p>是否开启日志投递，开启后需填写配置投递到指定cls</p>
	LogEnable *bool `json:"LogEnable,omitnil,omitempty" name:"LogEnable"`

	// <p>日志配置，需要投递服务日志到指定cls时填写</p>
	LogConfig *LogConfig `json:"LogConfig,omitnil,omitempty" name:"LogConfig"`

	// <p>是否开启接口鉴权，开启后自动生成token信息，访问需要token鉴权</p>
	AuthorizationEnable *bool `json:"AuthorizationEnable,omitnil,omitempty" name:"AuthorizationEnable"`

	// <p>腾讯云标签</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>是否新增版本</p>
	NewVersion *bool `json:"NewVersion,omitnil,omitempty" name:"NewVersion"`

	// <p>定时任务配置，使用定时策略时填写</p>
	CronScaleJobs []*CronScaleJob `json:"CronScaleJobs,omitnil,omitempty" name:"CronScaleJobs"`

	// <p>自动伸缩策略配置 HPA : 通过HPA进行弹性伸缩 CRON 通过定时任务进行伸缩</p>
	ScaleStrategy *string `json:"ScaleStrategy,omitnil,omitempty" name:"ScaleStrategy"`

	// <p>计费模式[HYBRID_PAID]时生效, 用于标识混合计费模式下的预付费实例数</p>
	HybridBillingPrepaidReplicas *int64 `json:"HybridBillingPrepaidReplicas,omitnil,omitempty" name:"HybridBillingPrepaidReplicas"`

	// <p>[AUTO_ML 自动学习，自动学习正式发布 AUTO_ML_FORMAL, DEFAULT 默认]</p>
	CreateSource *string `json:"CreateSource,omitnil,omitempty" name:"CreateSource"`

	// <p>是否开启模型的热更新。默认不开启</p>
	ModelHotUpdateEnable *bool `json:"ModelHotUpdateEnable,omitnil,omitempty" name:"ModelHotUpdateEnable"`

	// <p>定时停止配置</p>
	ScheduledAction *ScheduledAction `json:"ScheduledAction,omitnil,omitempty" name:"ScheduledAction"`

	// <p>挂载配置，目前只支持CFS</p>
	VolumeMount *VolumeMount `json:"VolumeMount,omitnil,omitempty" name:"VolumeMount"`

	// <p>服务限速限流相关配置</p>
	ServiceLimit *ServiceLimit `json:"ServiceLimit,omitnil,omitempty" name:"ServiceLimit"`

	// <p>回调地址，用于回调创建服务状态信息，回调格式&amp;内容详情见：<a href="https://cloud.tencent.com/document/product/851/84292">TI-ONE 接口回调说明</a></p>
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// <p>是否开启模型的加速, 仅对StableDiffusion(动态加速)格式的模型有效。</p>
	ModelTurboEnable *bool `json:"ModelTurboEnable,omitnil,omitempty" name:"ModelTurboEnable"`

	// <p>服务分类</p>
	ServiceCategory *string `json:"ServiceCategory,omitnil,omitempty" name:"ServiceCategory"`

	// <p>服务的启动命令，如遇特殊字符导致配置失败，可使用CommandBase64参数</p>
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// <p>是否开启TIONE内网访问外部，此功能仅支持后付费机型与从TIONE平台购买的预付费机型；使用从CVM选择资源组时此配置不生效。</p>
	ServiceEIP *ServiceEIP `json:"ServiceEIP,omitnil,omitempty" name:"ServiceEIP"`

	// <p>服务的启动命令，以base64格式进行输入，与Command同时配置时，仅当前参数生效</p>
	CommandBase64 *string `json:"CommandBase64,omitnil,omitempty" name:"CommandBase64"`

	// <p>服务端口，仅在非内置镜像时生效，默认8501。不支持输入8501-8510,6006,9092</p>
	ServicePort *int64 `json:"ServicePort,omitnil,omitempty" name:"ServicePort"`

	// <p>服务的部署类型 [标准部署，分布式多机部署，] 默认STANDARD</p><p>枚举值：</p><ul><li>STANDARD： 标准部署</li><li>DIST： 多机分布式部署</li><li>ROLE_SET： 多角色部署</li></ul>
	DeployType *string `json:"DeployType,omitnil,omitempty" name:"DeployType"`

	// <p>单副本下的实例数，仅在部署类型为DIST时生效，默认1</p>
	InstancePerReplicas *int64 `json:"InstancePerReplicas,omitnil,omitempty" name:"InstancePerReplicas"`

	// <p>服务的优雅退出时限。单位为秒，默认值为30，最小为1</p>
	TerminationGracePeriodSeconds *int64 `json:"TerminationGracePeriodSeconds,omitnil,omitempty" name:"TerminationGracePeriodSeconds"`

	// <p>服务实例停止前执行的命令，执行完毕或执行时间超过优雅退出时限后实例结束</p>
	PreStopCommand []*string `json:"PreStopCommand,omitnil,omitempty" name:"PreStopCommand"`

	// <p>是否启用 grpc 端口</p>
	GrpcEnable *bool `json:"GrpcEnable,omitnil,omitempty" name:"GrpcEnable"`

	// <p>健康探针</p>
	HealthProbe *HealthProbe `json:"HealthProbe,omitnil,omitempty" name:"HealthProbe"`

	// <p>滚动更新策略</p>
	RollingUpdate *RollingUpdate `json:"RollingUpdate,omitnil,omitempty" name:"RollingUpdate"`

	// <p>sidecar配置</p>
	Sidecar *SidecarSpec `json:"Sidecar,omitnil,omitempty" name:"Sidecar"`

	// <p>数据盘批量挂载配置，当前仅支持CFS，仅针对“模型来源-腾讯云存储、模型来源-腾讯云容器镜像、模型来源-资源组、模型来源-数据源”。</p>
	VolumeMounts []*VolumeMount `json:"VolumeMounts,omitnil,omitempty" name:"VolumeMounts"`

	// <p>调度策略 [binpack] 优先占满整机，尽量避免碎卡（默认值）[spread] 优先分散在各个节点，确保服务高可用</p>
	SchedulingStrategy *string `json:"SchedulingStrategy,omitnil,omitempty" name:"SchedulingStrategy"`

	// <p>网关日志投递相关配置</p>
	GatewayLogConfig *LogConfig `json:"GatewayLogConfig,omitnil,omitempty" name:"GatewayLogConfig"`

	// <p>网关相关配置</p>
	GatewayConfig *GatewayConfig `json:"GatewayConfig,omitnil,omitempty" name:"GatewayConfig"`
}

func (r *CreateModelServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateModelServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceGroupId")
	delete(f, "ServiceGroupName")
	delete(f, "ServiceDescription")
	delete(f, "ChargeType")
	delete(f, "ResourceGroupId")
	delete(f, "ModelInfo")
	delete(f, "ImageInfo")
	delete(f, "Env")
	delete(f, "Resources")
	delete(f, "InstanceType")
	delete(f, "ScaleMode")
	delete(f, "Replicas")
	delete(f, "HorizontalPodAutoscaler")
	delete(f, "LogEnable")
	delete(f, "LogConfig")
	delete(f, "AuthorizationEnable")
	delete(f, "Tags")
	delete(f, "NewVersion")
	delete(f, "CronScaleJobs")
	delete(f, "ScaleStrategy")
	delete(f, "HybridBillingPrepaidReplicas")
	delete(f, "CreateSource")
	delete(f, "ModelHotUpdateEnable")
	delete(f, "ScheduledAction")
	delete(f, "VolumeMount")
	delete(f, "ServiceLimit")
	delete(f, "CallbackUrl")
	delete(f, "ModelTurboEnable")
	delete(f, "ServiceCategory")
	delete(f, "Command")
	delete(f, "ServiceEIP")
	delete(f, "CommandBase64")
	delete(f, "ServicePort")
	delete(f, "DeployType")
	delete(f, "InstancePerReplicas")
	delete(f, "TerminationGracePeriodSeconds")
	delete(f, "PreStopCommand")
	delete(f, "GrpcEnable")
	delete(f, "HealthProbe")
	delete(f, "RollingUpdate")
	delete(f, "Sidecar")
	delete(f, "VolumeMounts")
	delete(f, "SchedulingStrategy")
	delete(f, "GatewayLogConfig")
	delete(f, "GatewayConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateModelServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateModelServiceResponseParams struct {
	// <p>生成的模型服务</p>
	Service *Service `json:"Service,omitnil,omitempty" name:"Service"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateModelServiceResponse struct {
	*tchttp.BaseResponse
	Response *CreateModelServiceResponseParams `json:"Response"`
}

func (r *CreateModelServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateModelServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMountLimitRequestParams struct {
	// 数据源类型英文名
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 存储实例ID
	StorageId *string `json:"StorageId,omitnil,omitempty" name:"StorageId"`

	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`

	// 限制开关是否开启，只有开启时才有限制，默认关闭
	LimitMount *bool `json:"LimitMount,omitnil,omitempty" name:"LimitMount"`
}

type CreateMountLimitRequest struct {
	*tchttp.BaseRequest
	
	// 数据源类型英文名
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 存储实例ID
	StorageId *string `json:"StorageId,omitnil,omitempty" name:"StorageId"`

	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`

	// 限制开关是否开启，只有开启时才有限制，默认关闭
	LimitMount *bool `json:"LimitMount,omitnil,omitempty" name:"LimitMount"`
}

func (r *CreateMountLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMountLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "StorageId")
	delete(f, "TiProjectId")
	delete(f, "LimitMount")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMountLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMountLimitResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateMountLimitResponse struct {
	*tchttp.BaseResponse
	Response *CreateMountLimitResponseParams `json:"Response"`
}

func (r *CreateMountLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMountLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNotebookRequestParams struct {
	// <p>名称。不超过60个字符，仅支持中英文、数字、下划线&quot;_&quot;、短横&quot;-&quot;，只能以中英文、数字开头</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>计算资源付费模式 ，可选值为：<br>PREPAID：预付费，即包年包月<br>POSTPAID_BY_HOUR：按小时后付费</p>
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// <p>计算资源配置</p>
	ResourceConf *ResourceConf `json:"ResourceConf,omitnil,omitempty" name:"ResourceConf"`

	// <p>是否上报日志</p>
	LogEnable *bool `json:"LogEnable,omitnil,omitempty" name:"LogEnable"`

	// <p>是否ROOT权限</p>
	RootAccess *bool `json:"RootAccess,omitnil,omitempty" name:"RootAccess"`

	// <p>是否自动停止</p>
	AutoStopping *bool `json:"AutoStopping,omitnil,omitempty" name:"AutoStopping"`

	// <p>是否访问公网</p>
	DirectInternetAccess *bool `json:"DirectInternetAccess,omitnil,omitempty" name:"DirectInternetAccess"`

	// <p>资源组ID(for预付费)</p>
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// <p>Vpc-Id</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>子网Id</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>存储的类型。取值包含：<br>FREE：预付费的免费存储<br>CLOUD_PREMIUM：高性能云硬盘<br>CLOUD_SSD：SSD云硬盘<br>CFS：CFS存储<br>CFS_TURBO：CFS Turbo存储<br>GooseFSx：GooseFSx存储</p>
	VolumeSourceType *string `json:"VolumeSourceType,omitnil,omitempty" name:"VolumeSourceType"`

	// <p>云硬盘存储卷大小，单位GB</p>
	VolumeSizeInGB *uint64 `json:"VolumeSizeInGB,omitnil,omitempty" name:"VolumeSizeInGB"`

	// <p>CFS存储的配置</p>
	VolumeSourceCFS *CFSConfig `json:"VolumeSourceCFS,omitnil,omitempty" name:"VolumeSourceCFS"`

	// <p>日志配置</p>
	LogConfig *LogConfig `json:"LogConfig,omitnil,omitempty" name:"LogConfig"`

	// <p>生命周期脚本的ID</p>
	LifecycleScriptId *string `json:"LifecycleScriptId,omitnil,omitempty" name:"LifecycleScriptId"`

	// <p>默认GIT存储库的ID</p>
	DefaultCodeRepoId *string `json:"DefaultCodeRepoId,omitnil,omitempty" name:"DefaultCodeRepoId"`

	// <p>其他GIT存储库的ID，最多3个</p>
	AdditionalCodeRepoIds []*string `json:"AdditionalCodeRepoIds,omitnil,omitempty" name:"AdditionalCodeRepoIds"`

	// <p>自动停止时间，单位小时</p>
	AutomaticStopTime *uint64 `json:"AutomaticStopTime,omitnil,omitempty" name:"AutomaticStopTime"`

	// <p>标签配置</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>数据存储挂载配置</p>
	DataConfigs []*DataConfig `json:"DataConfigs,omitnil,omitempty" name:"DataConfigs"`

	// <p>镜像信息</p>
	ImageInfo *ImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`

	// <p>镜像类型，包括SYSTEM、TCR、CCR</p>
	ImageType *string `json:"ImageType,omitnil,omitempty" name:"ImageType"`

	// <p>SSH配置信息</p>
	SSHConfig *SSHConfig `json:"SSHConfig,omitnil,omitempty" name:"SSHConfig"`

	// <p>GooseFS存储配置</p>
	VolumeSourceGooseFS *GooseFS `json:"VolumeSourceGooseFS,omitnil,omitempty" name:"VolumeSourceGooseFS"`

	// <p>描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateNotebookRequest struct {
	*tchttp.BaseRequest
	
	// <p>名称。不超过60个字符，仅支持中英文、数字、下划线&quot;_&quot;、短横&quot;-&quot;，只能以中英文、数字开头</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>计算资源付费模式 ，可选值为：<br>PREPAID：预付费，即包年包月<br>POSTPAID_BY_HOUR：按小时后付费</p>
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// <p>计算资源配置</p>
	ResourceConf *ResourceConf `json:"ResourceConf,omitnil,omitempty" name:"ResourceConf"`

	// <p>是否上报日志</p>
	LogEnable *bool `json:"LogEnable,omitnil,omitempty" name:"LogEnable"`

	// <p>是否ROOT权限</p>
	RootAccess *bool `json:"RootAccess,omitnil,omitempty" name:"RootAccess"`

	// <p>是否自动停止</p>
	AutoStopping *bool `json:"AutoStopping,omitnil,omitempty" name:"AutoStopping"`

	// <p>是否访问公网</p>
	DirectInternetAccess *bool `json:"DirectInternetAccess,omitnil,omitempty" name:"DirectInternetAccess"`

	// <p>资源组ID(for预付费)</p>
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// <p>Vpc-Id</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>子网Id</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>存储的类型。取值包含：<br>FREE：预付费的免费存储<br>CLOUD_PREMIUM：高性能云硬盘<br>CLOUD_SSD：SSD云硬盘<br>CFS：CFS存储<br>CFS_TURBO：CFS Turbo存储<br>GooseFSx：GooseFSx存储</p>
	VolumeSourceType *string `json:"VolumeSourceType,omitnil,omitempty" name:"VolumeSourceType"`

	// <p>云硬盘存储卷大小，单位GB</p>
	VolumeSizeInGB *uint64 `json:"VolumeSizeInGB,omitnil,omitempty" name:"VolumeSizeInGB"`

	// <p>CFS存储的配置</p>
	VolumeSourceCFS *CFSConfig `json:"VolumeSourceCFS,omitnil,omitempty" name:"VolumeSourceCFS"`

	// <p>日志配置</p>
	LogConfig *LogConfig `json:"LogConfig,omitnil,omitempty" name:"LogConfig"`

	// <p>生命周期脚本的ID</p>
	LifecycleScriptId *string `json:"LifecycleScriptId,omitnil,omitempty" name:"LifecycleScriptId"`

	// <p>默认GIT存储库的ID</p>
	DefaultCodeRepoId *string `json:"DefaultCodeRepoId,omitnil,omitempty" name:"DefaultCodeRepoId"`

	// <p>其他GIT存储库的ID，最多3个</p>
	AdditionalCodeRepoIds []*string `json:"AdditionalCodeRepoIds,omitnil,omitempty" name:"AdditionalCodeRepoIds"`

	// <p>自动停止时间，单位小时</p>
	AutomaticStopTime *uint64 `json:"AutomaticStopTime,omitnil,omitempty" name:"AutomaticStopTime"`

	// <p>标签配置</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>数据存储挂载配置</p>
	DataConfigs []*DataConfig `json:"DataConfigs,omitnil,omitempty" name:"DataConfigs"`

	// <p>镜像信息</p>
	ImageInfo *ImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`

	// <p>镜像类型，包括SYSTEM、TCR、CCR</p>
	ImageType *string `json:"ImageType,omitnil,omitempty" name:"ImageType"`

	// <p>SSH配置信息</p>
	SSHConfig *SSHConfig `json:"SSHConfig,omitnil,omitempty" name:"SSHConfig"`

	// <p>GooseFS存储配置</p>
	VolumeSourceGooseFS *GooseFS `json:"VolumeSourceGooseFS,omitnil,omitempty" name:"VolumeSourceGooseFS"`

	// <p>描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateNotebookRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNotebookRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "ChargeType")
	delete(f, "ResourceConf")
	delete(f, "LogEnable")
	delete(f, "RootAccess")
	delete(f, "AutoStopping")
	delete(f, "DirectInternetAccess")
	delete(f, "ResourceGroupId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "VolumeSourceType")
	delete(f, "VolumeSizeInGB")
	delete(f, "VolumeSourceCFS")
	delete(f, "LogConfig")
	delete(f, "LifecycleScriptId")
	delete(f, "DefaultCodeRepoId")
	delete(f, "AdditionalCodeRepoIds")
	delete(f, "AutomaticStopTime")
	delete(f, "Tags")
	delete(f, "DataConfigs")
	delete(f, "ImageInfo")
	delete(f, "ImageType")
	delete(f, "SSHConfig")
	delete(f, "VolumeSourceGooseFS")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNotebookRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNotebookResponseParams struct {
	// <p>notebook标志</p>
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateNotebookResponse struct {
	*tchttp.BaseResponse
	Response *CreateNotebookResponseParams `json:"Response"`
}

func (r *CreateNotebookResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNotebookResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePresignedNotebookUrlRequestParams struct {
	// Notebook ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type CreatePresignedNotebookUrlRequest struct {
	*tchttp.BaseRequest
	
	// Notebook ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

func (r *CreatePresignedNotebookUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePresignedNotebookUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePresignedNotebookUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePresignedNotebookUrlResponseParams struct {
	// 携带认证TOKEN的URL
	AuthorizedUrl *string `json:"AuthorizedUrl,omitnil,omitempty" name:"AuthorizedUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePresignedNotebookUrlResponse struct {
	*tchttp.BaseResponse
	Response *CreatePresignedNotebookUrlResponseParams `json:"Response"`
}

func (r *CreatePresignedNotebookUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePresignedNotebookUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTrainingModelRequestParams struct {
	// 导入方式
	// MODEL：导入新模型
	// VERSION：导入新版本
	// EXIST：导入现有版本
	ImportMethod *string `json:"ImportMethod,omitnil,omitempty" name:"ImportMethod"`

	// 推理环境来源（SYSTEM/CUSTOM）
	ReasoningEnvironmentSource *string `json:"ReasoningEnvironmentSource,omitnil,omitempty" name:"ReasoningEnvironmentSource"`

	// 模型名称，不超过60个字符，仅支持中英文、数字、下划线"_"、短横"-"，只能以中英文、数字开头
	TrainingModelName *string `json:"TrainingModelName,omitnil,omitempty" name:"TrainingModelName"`

	// 标签配置
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 训练任务名称
	TrainingJobName *string `json:"TrainingJobName,omitnil,omitempty" name:"TrainingJobName"`

	// 模型来源cos目录，以/结尾
	TrainingModelCosPath *CosPathInfo `json:"TrainingModelCosPath,omitnil,omitempty" name:"TrainingModelCosPath"`

	// 算法框架 （PYTORCH/TENSORFLOW/DETECTRON2/PMML/MMDETECTION/ONNX)
	AlgorithmFramework *string `json:"AlgorithmFramework,omitnil,omitempty" name:"AlgorithmFramework"`

	// 推理环境
	ReasoningEnvironment *string `json:"ReasoningEnvironment,omitnil,omitempty" name:"ReasoningEnvironment"`

	// 训练指标，最多支持1000字符
	TrainingModelIndex *string `json:"TrainingModelIndex,omitnil,omitempty" name:"TrainingModelIndex"`

	// 模型版本
	TrainingModelVersion *string `json:"TrainingModelVersion,omitnil,omitempty" name:"TrainingModelVersion"`

	// 自定义推理环境
	ReasoningImageInfo *ImageInfo `json:"ReasoningImageInfo,omitnil,omitempty" name:"ReasoningImageInfo"`

	// 模型移动方式（CUT/COPY）
	ModelMoveMode *string `json:"ModelMoveMode,omitnil,omitempty" name:"ModelMoveMode"`

	// 训练任务ID
	TrainingJobId *string `json:"TrainingJobId,omitnil,omitempty" name:"TrainingJobId"`

	// 模型ID（导入新模型不需要，导入新版本需要）
	TrainingModelId *string `json:"TrainingModelId,omitnil,omitempty" name:"TrainingModelId"`

	// 模型存储cos目录
	ModelOutputPath *CosPathInfo `json:"ModelOutputPath,omitnil,omitempty" name:"ModelOutputPath"`

	// 模型来源 （JOB/COS）
	TrainingModelSource *string `json:"TrainingModelSource,omitnil,omitempty" name:"TrainingModelSource"`

	// 模型偏好
	TrainingPreference *string `json:"TrainingPreference,omitnil,omitempty" name:"TrainingPreference"`

	// 自动学习任务ID（已废弃）
	AutoMLTaskId *string `json:"AutoMLTaskId,omitnil,omitempty" name:"AutoMLTaskId"`

	// 任务版本
	TrainingJobVersion *string `json:"TrainingJobVersion,omitnil,omitempty" name:"TrainingJobVersion"`

	// 模型版本类型；
	// 枚举值：NORMAL(通用)  ACCELERATE(加速)
	// 注意:  默认为NORMAL
	ModelVersionType *string `json:"ModelVersionType,omitnil,omitempty" name:"ModelVersionType"`

	// 模型格式 （PYTORCH/TORCH_SCRIPT/DETECTRON2/SAVED_MODEL/FROZEN_GRAPH/PMML/MMDETECTION/ONNX/HUGGING_FACE_BERT/HUGGING_FACE_STABLE_DIFFUSION/HUGGING_FACE_STABLE_DIFFUSION_LORA/WEB_UI_STABLE_DIFFUSION）
	ModelFormat *string `json:"ModelFormat,omitnil,omitempty" name:"ModelFormat"`

	// 推理镜像ID
	ReasoningEnvironmentId *string `json:"ReasoningEnvironmentId,omitnil,omitempty" name:"ReasoningEnvironmentId"`

	// 模型自动清理开关(true/false)，当前版本仅支持SAVED_MODEL格式模型
	AutoClean *string `json:"AutoClean,omitnil,omitempty" name:"AutoClean"`

	// 模型数量保留上限(默认值为24个，上限为24，下限为1，步长为1)
	MaxReservedModels *uint64 `json:"MaxReservedModels,omitnil,omitempty" name:"MaxReservedModels"`

	// 模型清理周期(默认值为1分钟，上限为1440，下限为1分钟，步长为1)
	ModelCleanPeriod *uint64 `json:"ModelCleanPeriod,omitnil,omitempty" name:"ModelCleanPeriod"`

	// 是否QAT模型
	IsQAT *bool `json:"IsQAT,omitnil,omitempty" name:"IsQAT"`
}

type CreateTrainingModelRequest struct {
	*tchttp.BaseRequest
	
	// 导入方式
	// MODEL：导入新模型
	// VERSION：导入新版本
	// EXIST：导入现有版本
	ImportMethod *string `json:"ImportMethod,omitnil,omitempty" name:"ImportMethod"`

	// 推理环境来源（SYSTEM/CUSTOM）
	ReasoningEnvironmentSource *string `json:"ReasoningEnvironmentSource,omitnil,omitempty" name:"ReasoningEnvironmentSource"`

	// 模型名称，不超过60个字符，仅支持中英文、数字、下划线"_"、短横"-"，只能以中英文、数字开头
	TrainingModelName *string `json:"TrainingModelName,omitnil,omitempty" name:"TrainingModelName"`

	// 标签配置
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 训练任务名称
	TrainingJobName *string `json:"TrainingJobName,omitnil,omitempty" name:"TrainingJobName"`

	// 模型来源cos目录，以/结尾
	TrainingModelCosPath *CosPathInfo `json:"TrainingModelCosPath,omitnil,omitempty" name:"TrainingModelCosPath"`

	// 算法框架 （PYTORCH/TENSORFLOW/DETECTRON2/PMML/MMDETECTION/ONNX)
	AlgorithmFramework *string `json:"AlgorithmFramework,omitnil,omitempty" name:"AlgorithmFramework"`

	// 推理环境
	ReasoningEnvironment *string `json:"ReasoningEnvironment,omitnil,omitempty" name:"ReasoningEnvironment"`

	// 训练指标，最多支持1000字符
	TrainingModelIndex *string `json:"TrainingModelIndex,omitnil,omitempty" name:"TrainingModelIndex"`

	// 模型版本
	TrainingModelVersion *string `json:"TrainingModelVersion,omitnil,omitempty" name:"TrainingModelVersion"`

	// 自定义推理环境
	ReasoningImageInfo *ImageInfo `json:"ReasoningImageInfo,omitnil,omitempty" name:"ReasoningImageInfo"`

	// 模型移动方式（CUT/COPY）
	ModelMoveMode *string `json:"ModelMoveMode,omitnil,omitempty" name:"ModelMoveMode"`

	// 训练任务ID
	TrainingJobId *string `json:"TrainingJobId,omitnil,omitempty" name:"TrainingJobId"`

	// 模型ID（导入新模型不需要，导入新版本需要）
	TrainingModelId *string `json:"TrainingModelId,omitnil,omitempty" name:"TrainingModelId"`

	// 模型存储cos目录
	ModelOutputPath *CosPathInfo `json:"ModelOutputPath,omitnil,omitempty" name:"ModelOutputPath"`

	// 模型来源 （JOB/COS）
	TrainingModelSource *string `json:"TrainingModelSource,omitnil,omitempty" name:"TrainingModelSource"`

	// 模型偏好
	TrainingPreference *string `json:"TrainingPreference,omitnil,omitempty" name:"TrainingPreference"`

	// 自动学习任务ID（已废弃）
	AutoMLTaskId *string `json:"AutoMLTaskId,omitnil,omitempty" name:"AutoMLTaskId"`

	// 任务版本
	TrainingJobVersion *string `json:"TrainingJobVersion,omitnil,omitempty" name:"TrainingJobVersion"`

	// 模型版本类型；
	// 枚举值：NORMAL(通用)  ACCELERATE(加速)
	// 注意:  默认为NORMAL
	ModelVersionType *string `json:"ModelVersionType,omitnil,omitempty" name:"ModelVersionType"`

	// 模型格式 （PYTORCH/TORCH_SCRIPT/DETECTRON2/SAVED_MODEL/FROZEN_GRAPH/PMML/MMDETECTION/ONNX/HUGGING_FACE_BERT/HUGGING_FACE_STABLE_DIFFUSION/HUGGING_FACE_STABLE_DIFFUSION_LORA/WEB_UI_STABLE_DIFFUSION）
	ModelFormat *string `json:"ModelFormat,omitnil,omitempty" name:"ModelFormat"`

	// 推理镜像ID
	ReasoningEnvironmentId *string `json:"ReasoningEnvironmentId,omitnil,omitempty" name:"ReasoningEnvironmentId"`

	// 模型自动清理开关(true/false)，当前版本仅支持SAVED_MODEL格式模型
	AutoClean *string `json:"AutoClean,omitnil,omitempty" name:"AutoClean"`

	// 模型数量保留上限(默认值为24个，上限为24，下限为1，步长为1)
	MaxReservedModels *uint64 `json:"MaxReservedModels,omitnil,omitempty" name:"MaxReservedModels"`

	// 模型清理周期(默认值为1分钟，上限为1440，下限为1分钟，步长为1)
	ModelCleanPeriod *uint64 `json:"ModelCleanPeriod,omitnil,omitempty" name:"ModelCleanPeriod"`

	// 是否QAT模型
	IsQAT *bool `json:"IsQAT,omitnil,omitempty" name:"IsQAT"`
}

func (r *CreateTrainingModelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTrainingModelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImportMethod")
	delete(f, "ReasoningEnvironmentSource")
	delete(f, "TrainingModelName")
	delete(f, "Tags")
	delete(f, "TrainingJobName")
	delete(f, "TrainingModelCosPath")
	delete(f, "AlgorithmFramework")
	delete(f, "ReasoningEnvironment")
	delete(f, "TrainingModelIndex")
	delete(f, "TrainingModelVersion")
	delete(f, "ReasoningImageInfo")
	delete(f, "ModelMoveMode")
	delete(f, "TrainingJobId")
	delete(f, "TrainingModelId")
	delete(f, "ModelOutputPath")
	delete(f, "TrainingModelSource")
	delete(f, "TrainingPreference")
	delete(f, "AutoMLTaskId")
	delete(f, "TrainingJobVersion")
	delete(f, "ModelVersionType")
	delete(f, "ModelFormat")
	delete(f, "ReasoningEnvironmentId")
	delete(f, "AutoClean")
	delete(f, "MaxReservedModels")
	delete(f, "ModelCleanPeriod")
	delete(f, "IsQAT")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTrainingModelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTrainingModelResponseParams struct {
	// 模型ID，TrainingModel ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 模型版本ID
	TrainingModelVersionId *string `json:"TrainingModelVersionId,omitnil,omitempty" name:"TrainingModelVersionId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTrainingModelResponse struct {
	*tchttp.BaseResponse
	Response *CreateTrainingModelResponseParams `json:"Response"`
}

func (r *CreateTrainingModelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTrainingModelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTrainingTaskRequestParams struct {
	// <p>训练任务名称，不超过60个字符，仅支持中英文、数字、下划线&quot;_&quot;、短横&quot;-&quot;，只能以中英文、数字开头</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>计费模式，eg：PREPAID 包年包月（资源组）;<br>POSTPAID_BY_HOUR 按量计费</p>
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// <p>资源配置，需填写对应算力规格ID和节点数量，算力规格ID查询接口为DescribeBillingSpecsPrice，eg：[{&quot;Role&quot;:&quot;WORKER&quot;, &quot;InstanceType&quot;: &quot;TI.S.MEDIUM.POST&quot;, &quot;InstanceNum&quot;: 1}]</p>
	ResourceConfigInfos []*ResourceConfigInfo `json:"ResourceConfigInfos,omitnil,omitempty" name:"ResourceConfigInfos"`

	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`

	// <p>训练框架名称，通过DescribeTrainingFrameworks接口查询，eg：SPARK、PYSPARK、TENSORFLOW、PYTORCH</p>
	FrameworkName *string `json:"FrameworkName,omitnil,omitempty" name:"FrameworkName"`

	// <p>训练框架版本，通过DescribeTrainingFrameworks接口查询，eg：1.15、1.9</p>
	FrameworkVersion *string `json:"FrameworkVersion,omitnil,omitempty" name:"FrameworkVersion"`

	// <p>训练框架环境，通过DescribeTrainingFrameworks接口查询，eg：tf1.15-py3.7-cpu、torch1.9-py3.8-cuda11.1-gpu</p>
	FrameworkEnvironment *string `json:"FrameworkEnvironment,omitnil,omitempty" name:"FrameworkEnvironment"`

	// <p>预付费专用资源组ID，通过DescribeBillingResourceGroups接口查询</p>
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// <p>标签配置</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>自定义镜像信息</p>
	ImageInfo *ImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`

	// <p>COS代码包路径</p>
	CodePackagePath *CosPathInfo `json:"CodePackagePath,omitnil,omitempty" name:"CodePackagePath"`

	// <p>任务的启动命令，按任务训练模式输入，如遇特殊字符导致配置失败，可使用EncodedStartCmdInfo参数</p>
	StartCmdInfo *StartCmdInfo `json:"StartCmdInfo,omitnil,omitempty" name:"StartCmdInfo"`

	// <p>训练模式，通过DescribeTrainingFrameworks接口查询，eg：PS_WORKER、DDP、MPI、HOROVOD</p>
	TrainingMode *string `json:"TrainingMode,omitnil,omitempty" name:"TrainingMode"`

	// <p>数据配置，依赖DataSource字段，数量不超过10个</p>
	DataConfigs []*DataConfig `json:"DataConfigs,omitnil,omitempty" name:"DataConfigs"`

	// <p>VPC Id</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>子网Id</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>COS训练输出路径</p>
	Output *CosPathInfo `json:"Output,omitnil,omitempty" name:"Output"`

	// <p>CLS日志配置</p>
	LogConfig *LogConfig `json:"LogConfig,omitnil,omitempty" name:"LogConfig"`

	// <p>调优参数，不超过2048个字符</p>
	TuningParameters *string `json:"TuningParameters,omitnil,omitempty" name:"TuningParameters"`

	// <p>是否上报日志</p>
	LogEnable *bool `json:"LogEnable,omitnil,omitempty" name:"LogEnable"`

	// <p>备注，不超过1024个字符</p>
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// <p>数据来源，eg：DATASET、COS、CFS、CFSTurbo、HDFS、GooseFSx</p>
	DataSource *string `json:"DataSource,omitnil,omitempty" name:"DataSource"`

	// <p>回调地址，用于创建/启动/停止训练任务的异步回调。回调格式&amp;内容详见：<a href="https://cloud.tencent.com/document/product/851/84292">[TI-ONE接口回调说明]</a></p>
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// <p>编码后的任务启动命令，与StartCmdInfo同时配置时，仅当前参数生效</p>
	EncodedStartCmdInfo *EncodedStartCmdInfo `json:"EncodedStartCmdInfo,omitnil,omitempty" name:"EncodedStartCmdInfo"`

	// <p>代码仓库配置</p>
	CodeRepos []*CodeRepoConfig `json:"CodeRepos,omitnil,omitempty" name:"CodeRepos"`

	// <p>网络暴露配置</p>
	ExposeNetworkConfig *ExposeNetworkConfig `json:"ExposeNetworkConfig,omitnil,omitempty" name:"ExposeNetworkConfig"`

	// <p>环境变量</p>
	Envs []*EnvVar `json:"Envs,omitnil,omitempty" name:"Envs"`
}

type CreateTrainingTaskRequest struct {
	*tchttp.BaseRequest
	
	// <p>训练任务名称，不超过60个字符，仅支持中英文、数字、下划线&quot;_&quot;、短横&quot;-&quot;，只能以中英文、数字开头</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>计费模式，eg：PREPAID 包年包月（资源组）;<br>POSTPAID_BY_HOUR 按量计费</p>
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// <p>资源配置，需填写对应算力规格ID和节点数量，算力规格ID查询接口为DescribeBillingSpecsPrice，eg：[{&quot;Role&quot;:&quot;WORKER&quot;, &quot;InstanceType&quot;: &quot;TI.S.MEDIUM.POST&quot;, &quot;InstanceNum&quot;: 1}]</p>
	ResourceConfigInfos []*ResourceConfigInfo `json:"ResourceConfigInfos,omitnil,omitempty" name:"ResourceConfigInfos"`

	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`

	// <p>训练框架名称，通过DescribeTrainingFrameworks接口查询，eg：SPARK、PYSPARK、TENSORFLOW、PYTORCH</p>
	FrameworkName *string `json:"FrameworkName,omitnil,omitempty" name:"FrameworkName"`

	// <p>训练框架版本，通过DescribeTrainingFrameworks接口查询，eg：1.15、1.9</p>
	FrameworkVersion *string `json:"FrameworkVersion,omitnil,omitempty" name:"FrameworkVersion"`

	// <p>训练框架环境，通过DescribeTrainingFrameworks接口查询，eg：tf1.15-py3.7-cpu、torch1.9-py3.8-cuda11.1-gpu</p>
	FrameworkEnvironment *string `json:"FrameworkEnvironment,omitnil,omitempty" name:"FrameworkEnvironment"`

	// <p>预付费专用资源组ID，通过DescribeBillingResourceGroups接口查询</p>
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// <p>标签配置</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>自定义镜像信息</p>
	ImageInfo *ImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`

	// <p>COS代码包路径</p>
	CodePackagePath *CosPathInfo `json:"CodePackagePath,omitnil,omitempty" name:"CodePackagePath"`

	// <p>任务的启动命令，按任务训练模式输入，如遇特殊字符导致配置失败，可使用EncodedStartCmdInfo参数</p>
	StartCmdInfo *StartCmdInfo `json:"StartCmdInfo,omitnil,omitempty" name:"StartCmdInfo"`

	// <p>训练模式，通过DescribeTrainingFrameworks接口查询，eg：PS_WORKER、DDP、MPI、HOROVOD</p>
	TrainingMode *string `json:"TrainingMode,omitnil,omitempty" name:"TrainingMode"`

	// <p>数据配置，依赖DataSource字段，数量不超过10个</p>
	DataConfigs []*DataConfig `json:"DataConfigs,omitnil,omitempty" name:"DataConfigs"`

	// <p>VPC Id</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>子网Id</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>COS训练输出路径</p>
	Output *CosPathInfo `json:"Output,omitnil,omitempty" name:"Output"`

	// <p>CLS日志配置</p>
	LogConfig *LogConfig `json:"LogConfig,omitnil,omitempty" name:"LogConfig"`

	// <p>调优参数，不超过2048个字符</p>
	TuningParameters *string `json:"TuningParameters,omitnil,omitempty" name:"TuningParameters"`

	// <p>是否上报日志</p>
	LogEnable *bool `json:"LogEnable,omitnil,omitempty" name:"LogEnable"`

	// <p>备注，不超过1024个字符</p>
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// <p>数据来源，eg：DATASET、COS、CFS、CFSTurbo、HDFS、GooseFSx</p>
	DataSource *string `json:"DataSource,omitnil,omitempty" name:"DataSource"`

	// <p>回调地址，用于创建/启动/停止训练任务的异步回调。回调格式&amp;内容详见：<a href="https://cloud.tencent.com/document/product/851/84292">[TI-ONE接口回调说明]</a></p>
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// <p>编码后的任务启动命令，与StartCmdInfo同时配置时，仅当前参数生效</p>
	EncodedStartCmdInfo *EncodedStartCmdInfo `json:"EncodedStartCmdInfo,omitnil,omitempty" name:"EncodedStartCmdInfo"`

	// <p>代码仓库配置</p>
	CodeRepos []*CodeRepoConfig `json:"CodeRepos,omitnil,omitempty" name:"CodeRepos"`

	// <p>网络暴露配置</p>
	ExposeNetworkConfig *ExposeNetworkConfig `json:"ExposeNetworkConfig,omitnil,omitempty" name:"ExposeNetworkConfig"`

	// <p>环境变量</p>
	Envs []*EnvVar `json:"Envs,omitnil,omitempty" name:"Envs"`
}

func (r *CreateTrainingTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTrainingTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "ChargeType")
	delete(f, "ResourceConfigInfos")
	delete(f, "TiProjectId")
	delete(f, "FrameworkName")
	delete(f, "FrameworkVersion")
	delete(f, "FrameworkEnvironment")
	delete(f, "ResourceGroupId")
	delete(f, "Tags")
	delete(f, "ImageInfo")
	delete(f, "CodePackagePath")
	delete(f, "StartCmdInfo")
	delete(f, "TrainingMode")
	delete(f, "DataConfigs")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "Output")
	delete(f, "LogConfig")
	delete(f, "TuningParameters")
	delete(f, "LogEnable")
	delete(f, "Remark")
	delete(f, "DataSource")
	delete(f, "CallbackUrl")
	delete(f, "EncodedStartCmdInfo")
	delete(f, "CodeRepos")
	delete(f, "ExposeNetworkConfig")
	delete(f, "Envs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTrainingTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTrainingTaskResponseParams struct {
	// <p>训练任务ID</p>
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTrainingTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateTrainingTaskResponseParams `json:"Response"`
}

func (r *CreateTrainingTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTrainingTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CronScaleJob struct {
	// Cron表达式，标识任务的执行时间，精确到分钟级
	Schedule *string `json:"Schedule,omitnil,omitempty" name:"Schedule"`

	// 定时任务名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 目标实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetReplicas *int64 `json:"TargetReplicas,omitnil,omitempty" name:"TargetReplicas"`

	// 目标min
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinReplicas *int64 `json:"MinReplicas,omitnil,omitempty" name:"MinReplicas"`

	// 目标max
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxReplicas *int64 `json:"MaxReplicas,omitnil,omitempty" name:"MaxReplicas"`

	// 例外时间，Cron表达式，在对应时间内不执行任务。最多支持3条。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExcludeDates []*string `json:"ExcludeDates,omitnil,omitempty" name:"ExcludeDates"`
}

type CrossTenantENIInfo struct {
	// Pod IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrimaryIP *string `json:"PrimaryIP,omitnil,omitempty" name:"PrimaryIP"`

	// Pod Port
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`
}

type DataConfig struct {
	// <p>映射路径</p>
	MappingPath *string `json:"MappingPath,omitnil,omitempty" name:"MappingPath"`

	// <p>存储用途<br>可选值为 BUILTIN_CODE, BUILTIN_DATA, BUILTIN_MODEL, USER_DATA, USER_CODE, USER_MODEL, OUTPUT, OTHER</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSourceUsage *string `json:"DataSourceUsage,omitnil,omitempty" name:"DataSourceUsage"`

	// <p>DATASET、COS、CFS、CFSTurbo、GooseFSx、HDFS、WEDATA_HDFS</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSourceType *string `json:"DataSourceType,omitnil,omitempty" name:"DataSourceType"`

	// <p>来自数据集的数据</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSetSource *DataSetConfig `json:"DataSetSource,omitnil,omitempty" name:"DataSetSource"`

	// <p>来自cos的数据</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	COSSource *CosPathInfo `json:"COSSource,omitnil,omitempty" name:"COSSource"`

	// <p>来自CFS的数据</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CFSSource *CFSConfig `json:"CFSSource,omitnil,omitempty" name:"CFSSource"`

	// <p>来自HDFS的数据</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	HDFSSource *HDFSConfig `json:"HDFSSource,omitnil,omitempty" name:"HDFSSource"`

	// <p>配置GooseFS的数据</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	GooseFSSource *GooseFS `json:"GooseFSSource,omitnil,omitempty" name:"GooseFSSource"`

	// <p>配置TurboFS的数据</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CFSTurboSource *CFSTurbo `json:"CFSTurboSource,omitnil,omitempty" name:"CFSTurboSource"`

	// <p>来自本地磁盘的信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocalDiskSource *LocalDisk `json:"LocalDiskSource,omitnil,omitempty" name:"LocalDiskSource"`

	// <p>CBS配置信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CBSSource *CBSConfig `json:"CBSSource,omitnil,omitempty" name:"CBSSource"`

	// <p>主机路径信息</p>
	HostPathSource *HostPath `json:"HostPathSource,omitnil,omitempty" name:"HostPathSource"`

	// <p>公有云数据源</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicDataSource *PublicDataSourceFS `json:"PublicDataSource,omitnil,omitempty" name:"PublicDataSource"`
}

type DataPoint struct {
	// 指标名字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 值
	Value *float64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type DataSetConfig struct {
	// 数据集ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type DataSourceInfo struct {
	// <p>数据源ID</p>
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>数据源名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>创建者uin</p>
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// <p>创建者名称</p>
	CreatorName *string `json:"CreatorName,omitnil,omitempty" name:"CreatorName"`

	// <p>数据源类型英文名</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>数据源权限，取值有RW RO</p>
	Permission *string `json:"Permission,omitnil,omitempty" name:"Permission"`

	// <p>数据源所属存储实例ID</p>
	StorageId *string `json:"StorageId,omitnil,omitempty" name:"StorageId"`

	// <p>数据源所属存储实例名称</p>
	StorageName *string `json:"StorageName,omitnil,omitempty" name:"StorageName"`

	// <p>数据源挂载配置</p>
	MountConfigure *MountConfigureInfo `json:"MountConfigure,omitnil,omitempty" name:"MountConfigure"`

	// <p>创建时间, 格式为yyyy-mm-ddThh:mm:ssZ</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>更新时间, 格式为yyyy-mm-ddThh:mm:ssZ</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>限制开关是否开启，只有开启时才有限制</p>
	LimitMount *bool `json:"LimitMount,omitnil,omitempty" name:"LimitMount"`

	// <p>标签配置</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>额外配置,对应存储实例的额外配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraConf *StorageExtraConf `json:"ExtraConf,omitnil,omitempty" name:"ExtraConf"`
}

type DatasetGroup struct {
	// 数据集ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetId *string `json:"DatasetId,omitnil,omitempty" name:"DatasetId"`

	// 数据集名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetName *string `json:"DatasetName,omitnil,omitempty" name:"DatasetName"`

	// 创建者
	// 注意：此字段可能返回 null，表示取不到有效值。
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 数据集版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetVersion *string `json:"DatasetVersion,omitnil,omitempty" name:"DatasetVersion"`

	// 数据集类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetType *string `json:"DatasetType,omitnil,omitempty" name:"DatasetType"`

	// 数据集标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetTags []*Tag `json:"DatasetTags,omitnil,omitempty" name:"DatasetTags"`

	// 数据集标注任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetAnnotationTaskName *string `json:"DatasetAnnotationTaskName,omitnil,omitempty" name:"DatasetAnnotationTaskName"`

	// 数据集标注任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetAnnotationTaskId *string `json:"DatasetAnnotationTaskId,omitnil,omitempty" name:"DatasetAnnotationTaskId"`

	// 处理进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Process *uint64 `json:"Process,omitnil,omitempty" name:"Process"`

	// 数据集状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetStatus *string `json:"DatasetStatus,omitnil,omitempty" name:"DatasetStatus"`

	// 错误详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 外部关联TASKType
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalTaskType *string `json:"ExternalTaskType,omitnil,omitempty" name:"ExternalTaskType"`

	// 数据集大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetSize *string `json:"DatasetSize,omitnil,omitempty" name:"DatasetSize"`

	// 数据集数据量
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileNum *uint64 `json:"FileNum,omitnil,omitempty" name:"FileNum"`

	// 数据集源COS路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageDataPath *CosPathInfo `json:"StorageDataPath,omitnil,omitempty" name:"StorageDataPath"`

	// 数据集标签存储路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageLabelPath *CosPathInfo `json:"StorageLabelPath,omitnil,omitempty" name:"StorageLabelPath"`

	// 数据集版本聚合详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetVersions []*DatasetInfo `json:"DatasetVersions,omitnil,omitempty" name:"DatasetVersions"`

	// 数据集标注状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnnotationStatus *string `json:"AnnotationStatus,omitnil,omitempty" name:"AnnotationStatus"`

	// 数据集类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnnotationType *string `json:"AnnotationType,omitnil,omitempty" name:"AnnotationType"`

	// 数据集标注格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnnotationFormat *string `json:"AnnotationFormat,omitnil,omitempty" name:"AnnotationFormat"`

	// 数据集范围
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetScope *string `json:"DatasetScope,omitnil,omitempty" name:"DatasetScope"`

	// 数据集OCR子场景
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrScene *string `json:"OcrScene,omitnil,omitempty" name:"OcrScene"`

	// 数据集字典修改状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnnotationKeyStatus *string `json:"AnnotationKeyStatus,omitnil,omitempty" name:"AnnotationKeyStatus"`

	// 文本数据集导入方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContentType *string `json:"ContentType,omitnil,omitempty" name:"ContentType"`

	// 数据集建模类别。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetScene *string `json:"DatasetScene,omitnil,omitempty" name:"DatasetScene"`

	// CFS配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	CFSConfig *CFSConfig `json:"CFSConfig,omitnil,omitempty" name:"CFSConfig"`

	// 数据集标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	SceneTags []*string `json:"SceneTags,omitnil,omitempty" name:"SceneTags"`

	// 已标注数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	NumAnnotated *uint64 `json:"NumAnnotated,omitnil,omitempty" name:"NumAnnotated"`

	// 标注规范
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnnotationSpecification *string `json:"AnnotationSpecification,omitnil,omitempty" name:"AnnotationSpecification"`

	// 标注Schema是否配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnnotationSchemaConfigured *bool `json:"AnnotationSchemaConfigured,omitnil,omitempty" name:"AnnotationSchemaConfigured"`

	// 创建者名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatorNickname *string `json:"CreatorNickname,omitnil,omitempty" name:"CreatorNickname"`

	// cfs路径是否有修改
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsCfsUpdated *bool `json:"IsCfsUpdated,omitnil,omitempty" name:"IsCfsUpdated"`
}

type DatasetInfo struct {
	// 数据集id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetId *string `json:"DatasetId,omitnil,omitempty" name:"DatasetId"`

	// 数据集名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetName *string `json:"DatasetName,omitnil,omitempty" name:"DatasetName"`

	// 数据集创建者
	// 注意：此字段可能返回 null，表示取不到有效值。
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 数据集版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetVersion *string `json:"DatasetVersion,omitnil,omitempty" name:"DatasetVersion"`

	// 数据集类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetType *string `json:"DatasetType,omitnil,omitempty" name:"DatasetType"`

	// 数据集标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetTags []*Tag `json:"DatasetTags,omitnil,omitempty" name:"DatasetTags"`

	// 数据集对应标注任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetAnnotationTaskName *string `json:"DatasetAnnotationTaskName,omitnil,omitempty" name:"DatasetAnnotationTaskName"`

	// 数据集对应标注任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetAnnotationTaskId *string `json:"DatasetAnnotationTaskId,omitnil,omitempty" name:"DatasetAnnotationTaskId"`

	// 处理进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Process *int64 `json:"Process,omitnil,omitempty" name:"Process"`

	// 数据集状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetStatus *string `json:"DatasetStatus,omitnil,omitempty" name:"DatasetStatus"`

	// 错误详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 数据集创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 数据集更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 外部任务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalTaskType *string `json:"ExternalTaskType,omitnil,omitempty" name:"ExternalTaskType"`

	// 数据集存储大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetSize *string `json:"DatasetSize,omitnil,omitempty" name:"DatasetSize"`

	// 数据集数据数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileNum *uint64 `json:"FileNum,omitnil,omitempty" name:"FileNum"`

	// 数据集源cos 路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageDataPath *CosPathInfo `json:"StorageDataPath,omitnil,omitempty" name:"StorageDataPath"`

	// 数据集输出cos路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageLabelPath *CosPathInfo `json:"StorageLabelPath,omitnil,omitempty" name:"StorageLabelPath"`

	// 数据集标注状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnnotationStatus *string `json:"AnnotationStatus,omitnil,omitempty" name:"AnnotationStatus"`

	// 数据集类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnnotationType *string `json:"AnnotationType,omitnil,omitempty" name:"AnnotationType"`

	// 数据集标注格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnnotationFormat *string `json:"AnnotationFormat,omitnil,omitempty" name:"AnnotationFormat"`

	// 数据集范围
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetScope *string `json:"DatasetScope,omitnil,omitempty" name:"DatasetScope"`

	// 数据集OCR子场景
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrScene *string `json:"OcrScene,omitnil,omitempty" name:"OcrScene"`

	// 数据集字典修改状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnnotationKeyStatus *string `json:"AnnotationKeyStatus,omitnil,omitempty" name:"AnnotationKeyStatus"`

	// 内容类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContentType *string `json:"ContentType,omitnil,omitempty" name:"ContentType"`

	// 数据集建模类别。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetScene *string `json:"DatasetScene,omitnil,omitempty" name:"DatasetScene"`

	// CFS配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	CFSConfig *CFSConfig `json:"CFSConfig,omitnil,omitempty" name:"CFSConfig"`

	// 数据集标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	SceneTags []*string `json:"SceneTags,omitnil,omitempty" name:"SceneTags"`

	// 已标注数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	NumAnnotated *uint64 `json:"NumAnnotated,omitnil,omitempty" name:"NumAnnotated"`

	// 标注规范
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnnotationSpecification *string `json:"AnnotationSpecification,omitnil,omitempty" name:"AnnotationSpecification"`

	// 标注Schema是否配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnnotationSchemaConfigured *bool `json:"AnnotationSchemaConfigured,omitnil,omitempty" name:"AnnotationSchemaConfigured"`

	// 创建者名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatorNickname *string `json:"CreatorNickname,omitnil,omitempty" name:"CreatorNickname"`

	// cfs路径是否有修改
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsCfsUpdated *bool `json:"IsCfsUpdated,omitnil,omitempty" name:"IsCfsUpdated"`
}

type DefaultInnerCallInfo struct {
	// 可以进行调用的VPC-ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcIds []*string `json:"VpcIds,omitnil,omitempty" name:"VpcIds"`

	// 默认内网调用地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerHttpAddr *string `json:"InnerHttpAddr,omitnil,omitempty" name:"InnerHttpAddr"`
}

type DefaultNginxGatewayCallInfo struct {
	// host
	// 注意：此字段可能返回 null，表示取不到有效值。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
}

// Predefined struct for user
type DeleteDataSourceRequestParams struct {
	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`

	// 数据源ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type DeleteDataSourceRequest struct {
	*tchttp.BaseRequest
	
	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`

	// 数据源ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

func (r *DeleteDataSourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDataSourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TiProjectId")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDataSourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDataSourceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDataSourceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDataSourceResponseParams `json:"Response"`
}

func (r *DeleteDataSourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDataSourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDatasetRequestParams struct {
	// 数据集id
	DatasetId *string `json:"DatasetId,omitnil,omitempty" name:"DatasetId"`

	// 是否删除cos标签文件
	DeleteLabelEnable *bool `json:"DeleteLabelEnable,omitnil,omitempty" name:"DeleteLabelEnable"`
}

type DeleteDatasetRequest struct {
	*tchttp.BaseRequest
	
	// 数据集id
	DatasetId *string `json:"DatasetId,omitnil,omitempty" name:"DatasetId"`

	// 是否删除cos标签文件
	DeleteLabelEnable *bool `json:"DeleteLabelEnable,omitnil,omitempty" name:"DeleteLabelEnable"`
}

func (r *DeleteDatasetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDatasetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatasetId")
	delete(f, "DeleteLabelEnable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDatasetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDatasetResponseParams struct {
	// 删除的datasetId
	DatasetId *string `json:"DatasetId,omitnil,omitempty" name:"DatasetId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDatasetResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDatasetResponseParams `json:"Response"`
}

func (r *DeleteDatasetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDatasetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteExportRequestParams struct {
	// 日志下载任务的ID
	ExportId *string `json:"ExportId,omitnil,omitempty" name:"ExportId"`
}

type DeleteExportRequest struct {
	*tchttp.BaseRequest
	
	// 日志下载任务的ID
	ExportId *string `json:"ExportId,omitnil,omitempty" name:"ExportId"`
}

func (r *DeleteExportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteExportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ExportId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteExportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteExportResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteExportResponse struct {
	*tchttp.BaseResponse
	Response *DeleteExportResponseParams `json:"Response"`
}

func (r *DeleteExportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteModelServiceAuthTokenRequestParams struct {
	// 服务组 id
	ServiceGroupId *string `json:"ServiceGroupId,omitnil,omitempty" name:"ServiceGroupId"`

	// token 值
	AuthTokenValue *string `json:"AuthTokenValue,omitnil,omitempty" name:"AuthTokenValue"`
}

type DeleteModelServiceAuthTokenRequest struct {
	*tchttp.BaseRequest
	
	// 服务组 id
	ServiceGroupId *string `json:"ServiceGroupId,omitnil,omitempty" name:"ServiceGroupId"`

	// token 值
	AuthTokenValue *string `json:"AuthTokenValue,omitnil,omitempty" name:"AuthTokenValue"`
}

func (r *DeleteModelServiceAuthTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteModelServiceAuthTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceGroupId")
	delete(f, "AuthTokenValue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteModelServiceAuthTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteModelServiceAuthTokenResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteModelServiceAuthTokenResponse struct {
	*tchttp.BaseResponse
	Response *DeleteModelServiceAuthTokenResponseParams `json:"Response"`
}

func (r *DeleteModelServiceAuthTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteModelServiceAuthTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteModelServiceGroupRequestParams struct {
	// 服务id
	ServiceGroupId *string `json:"ServiceGroupId,omitnil,omitempty" name:"ServiceGroupId"`
}

type DeleteModelServiceGroupRequest struct {
	*tchttp.BaseRequest
	
	// 服务id
	ServiceGroupId *string `json:"ServiceGroupId,omitnil,omitempty" name:"ServiceGroupId"`
}

func (r *DeleteModelServiceGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteModelServiceGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteModelServiceGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteModelServiceGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteModelServiceGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteModelServiceGroupResponseParams `json:"Response"`
}

func (r *DeleteModelServiceGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteModelServiceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteModelServiceRequestParams struct {
	// 服务id
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 服务分类
	//
	// Deprecated: ServiceCategory is deprecated.
	ServiceCategory *string `json:"ServiceCategory,omitnil,omitempty" name:"ServiceCategory"`
}

type DeleteModelServiceRequest struct {
	*tchttp.BaseRequest
	
	// 服务id
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 服务分类
	ServiceCategory *string `json:"ServiceCategory,omitnil,omitempty" name:"ServiceCategory"`
}

func (r *DeleteModelServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteModelServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "ServiceCategory")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteModelServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteModelServiceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteModelServiceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteModelServiceResponseParams `json:"Response"`
}

func (r *DeleteModelServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteModelServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMountLimitRequestParams struct {
	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`

	// 数据源类型英文名
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 存储实例ID
	StorageId *string `json:"StorageId,omitnil,omitempty" name:"StorageId"`
}

type DeleteMountLimitRequest struct {
	*tchttp.BaseRequest
	
	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`

	// 数据源类型英文名
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 存储实例ID
	StorageId *string `json:"StorageId,omitnil,omitempty" name:"StorageId"`
}

func (r *DeleteMountLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMountLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TiProjectId")
	delete(f, "Type")
	delete(f, "StorageId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMountLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMountLimitResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteMountLimitResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMountLimitResponseParams `json:"Response"`
}

func (r *DeleteMountLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMountLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNotebookRequestParams struct {
	// notebook id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type DeleteNotebookRequest struct {
	*tchttp.BaseRequest
	
	// notebook id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

func (r *DeleteNotebookRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNotebookRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNotebookRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNotebookResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteNotebookResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNotebookResponseParams `json:"Response"`
}

func (r *DeleteNotebookResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNotebookResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTrainingModelRequestParams struct {
	// 模型ID
	TrainingModelId *string `json:"TrainingModelId,omitnil,omitempty" name:"TrainingModelId"`

	// 是否同步清理cos
	EnableDeleteCos *bool `json:"EnableDeleteCos,omitnil,omitempty" name:"EnableDeleteCos"`

	// 删除模型类型，枚举值：NORMAL 普通，ACCELERATE 加速，不传则删除所有
	ModelVersionType *string `json:"ModelVersionType,omitnil,omitempty" name:"ModelVersionType"`
}

type DeleteTrainingModelRequest struct {
	*tchttp.BaseRequest
	
	// 模型ID
	TrainingModelId *string `json:"TrainingModelId,omitnil,omitempty" name:"TrainingModelId"`

	// 是否同步清理cos
	EnableDeleteCos *bool `json:"EnableDeleteCos,omitnil,omitempty" name:"EnableDeleteCos"`

	// 删除模型类型，枚举值：NORMAL 普通，ACCELERATE 加速，不传则删除所有
	ModelVersionType *string `json:"ModelVersionType,omitnil,omitempty" name:"ModelVersionType"`
}

func (r *DeleteTrainingModelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTrainingModelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TrainingModelId")
	delete(f, "EnableDeleteCos")
	delete(f, "ModelVersionType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTrainingModelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTrainingModelResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTrainingModelResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTrainingModelResponseParams `json:"Response"`
}

func (r *DeleteTrainingModelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTrainingModelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTrainingModelVersionRequestParams struct {
	// 模型版本ID
	TrainingModelVersionId *string `json:"TrainingModelVersionId,omitnil,omitempty" name:"TrainingModelVersionId"`

	// 是否同步清理cos
	EnableDeleteCos *bool `json:"EnableDeleteCos,omitnil,omitempty" name:"EnableDeleteCos"`
}

type DeleteTrainingModelVersionRequest struct {
	*tchttp.BaseRequest
	
	// 模型版本ID
	TrainingModelVersionId *string `json:"TrainingModelVersionId,omitnil,omitempty" name:"TrainingModelVersionId"`

	// 是否同步清理cos
	EnableDeleteCos *bool `json:"EnableDeleteCos,omitnil,omitempty" name:"EnableDeleteCos"`
}

func (r *DeleteTrainingModelVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTrainingModelVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TrainingModelVersionId")
	delete(f, "EnableDeleteCos")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTrainingModelVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTrainingModelVersionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTrainingModelVersionResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTrainingModelVersionResponseParams `json:"Response"`
}

func (r *DeleteTrainingModelVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTrainingModelVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTrainingTaskRequestParams struct {
	// 训练任务ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`
}

type DeleteTrainingTaskRequest struct {
	*tchttp.BaseRequest
	
	// 训练任务ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`
}

func (r *DeleteTrainingTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTrainingTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "TiProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTrainingTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTrainingTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTrainingTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTrainingTaskResponseParams `json:"Response"`
}

func (r *DeleteTrainingTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTrainingTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAnnotatedTaskListRequestParams struct {
	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 页面大小，默认为10
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件数组，支持数据集ID，标注场景、任务状态、数据集名称、人物名称的过滤，后面两个支持模糊查询
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 标签过滤条件
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`

	// 排序方向：Asc Desc
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

type DescribeAnnotatedTaskListRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 页面大小，默认为10
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件数组，支持数据集ID，标注场景、任务状态、数据集名称、人物名称的过滤，后面两个支持模糊查询
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 标签过滤条件
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`

	// 排序方向：Asc Desc
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

func (r *DescribeAnnotatedTaskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAnnotatedTaskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "TagFilters")
	delete(f, "Order")
	delete(f, "OrderField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAnnotatedTaskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAnnotatedTaskListResponseParams struct {
	// 任务列表总数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 标注任务详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskList []*AnnotationTaskInfo `json:"TaskList,omitnil,omitempty" name:"TaskList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAnnotatedTaskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAnnotatedTaskListResponseParams `json:"Response"`
}

func (r *DescribeAnnotatedTaskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAnnotatedTaskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillingResourceGroupAttachedWorkspacesRequestParams struct {
	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`
}

type DescribeBillingResourceGroupAttachedWorkspacesRequest struct {
	*tchttp.BaseRequest
	
	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`
}

func (r *DescribeBillingResourceGroupAttachedWorkspacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillingResourceGroupAttachedWorkspacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TiProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillingResourceGroupAttachedWorkspacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillingResourceGroupAttachedWorkspacesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBillingResourceGroupAttachedWorkspacesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillingResourceGroupAttachedWorkspacesResponseParams `json:"Response"`
}

func (r *DescribeBillingResourceGroupAttachedWorkspacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillingResourceGroupAttachedWorkspacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillingResourceGroupRequestParams struct {
	// 资源组id, 取值为创建资源组接口(CreateBillingResourceGroup)响应中的ResourceGroupId
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`

	// 过滤条件
	// 注意: 
	// 1. Filter.Name 只支持以下枚举值:
	//     InstanceId (资源组节点id)
	//     InstanceStatus (资源组节点状态)
	// 2. Filter.Values: 长度为1且Filter.Fuzzy=true时，支持模糊查询; 不为1时，精确查询
	// 3. Filter.Negative: 是否取反，默认为false
	// 4. Filter.Fuzzy: 是否模糊查询，默认为false
	// 5. 每次请求的Filters的上限为10，Filter.Values的上限为100
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页查询起始位置，如：Limit为10，第一页Offset为0，第二页Offset为10...即每页左边为闭区间; 默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询每页大小，默认20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序方向; 枚举值: ASC | DESC；默认DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段; 枚举值: CreateTime (创建时间) ｜ ExpireTime (到期时间)；默认CreateTime
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

type DescribeBillingResourceGroupRequest struct {
	*tchttp.BaseRequest
	
	// 资源组id, 取值为创建资源组接口(CreateBillingResourceGroup)响应中的ResourceGroupId
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`

	// 过滤条件
	// 注意: 
	// 1. Filter.Name 只支持以下枚举值:
	//     InstanceId (资源组节点id)
	//     InstanceStatus (资源组节点状态)
	// 2. Filter.Values: 长度为1且Filter.Fuzzy=true时，支持模糊查询; 不为1时，精确查询
	// 3. Filter.Negative: 是否取反，默认为false
	// 4. Filter.Fuzzy: 是否模糊查询，默认为false
	// 5. 每次请求的Filters的上限为10，Filter.Values的上限为100
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页查询起始位置，如：Limit为10，第一页Offset为0，第二页Offset为10...即每页左边为闭区间; 默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询每页大小，默认20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序方向; 枚举值: ASC | DESC；默认DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段; 枚举值: CreateTime (创建时间) ｜ ExpireTime (到期时间)；默认CreateTime
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

func (r *DescribeBillingResourceGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillingResourceGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceGroupId")
	delete(f, "TiProjectId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "OrderField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillingResourceGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillingResourceGroupResponseParams struct {
	// 资源组节点总数； 注意接口是分页拉取的，total是指资源组节点总数，不是本次返回中InstanceSet数组的大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 资源组节点信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceSet []*Instance `json:"InstanceSet,omitnil,omitempty" name:"InstanceSet"`

	// 资源组纳管类型
	ResourceGroupSWType *string `json:"ResourceGroupSWType,omitnil,omitempty" name:"ResourceGroupSWType"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBillingResourceGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillingResourceGroupResponseParams `json:"Response"`
}

func (r *DescribeBillingResourceGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillingResourceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillingResourceGroupsRequestParams struct {
	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`

	// Filter.Name: 枚举值: ResourceGroupId (资源组id列表)                    ResourceGroupName (资源组名称列表)                    AvailableNodeCount（资源组中可用节点数量） Filter.Values: 长度为1且Filter.Fuzzy=true时，支持模糊查询; 不为1时，精确查询每次请求的Filters的上限为5，Filter.Values的上限为100
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 标签过滤
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`

	// 偏移量，默认为0；分页查询起始位置，如：Limit为100，第一页Offset为0，第二页OffSet为100....即每页左边为闭区间
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询每页大小，默认20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 支持模糊查找资源组id和资源组名
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// 是否不展示节点列表; 
	// true: 不展示，false 展示；
	// 默认为false
	DontShowInstanceSet *bool `json:"DontShowInstanceSet,omitnil,omitempty" name:"DontShowInstanceSet"`
}

type DescribeBillingResourceGroupsRequest struct {
	*tchttp.BaseRequest
	
	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`

	// Filter.Name: 枚举值: ResourceGroupId (资源组id列表)                    ResourceGroupName (资源组名称列表)                    AvailableNodeCount（资源组中可用节点数量） Filter.Values: 长度为1且Filter.Fuzzy=true时，支持模糊查询; 不为1时，精确查询每次请求的Filters的上限为5，Filter.Values的上限为100
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 标签过滤
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`

	// 偏移量，默认为0；分页查询起始位置，如：Limit为100，第一页Offset为0，第二页OffSet为100....即每页左边为闭区间
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询每页大小，默认20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 支持模糊查找资源组id和资源组名
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// 是否不展示节点列表; 
	// true: 不展示，false 展示；
	// 默认为false
	DontShowInstanceSet *bool `json:"DontShowInstanceSet,omitnil,omitempty" name:"DontShowInstanceSet"`
}

func (r *DescribeBillingResourceGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillingResourceGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TiProjectId")
	delete(f, "Filters")
	delete(f, "TagFilters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchWord")
	delete(f, "DontShowInstanceSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillingResourceGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillingResourceGroupsResponseParams struct {
	// 资源组总数； 注意接口是分页拉取的，total是指资源组总数，不是本次返回中ResourceGroupSet数组的大小
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 资源组详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupSet []*ResourceGroup `json:"ResourceGroupSet,omitnil,omitempty" name:"ResourceGroupSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBillingResourceGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillingResourceGroupsResponseParams `json:"Response"`
}

func (r *DescribeBillingResourceGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillingResourceGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillingResourceInstanceRunningJobsRequestParams struct {
	// 资源组id
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// 资源组节点id
	ResourceInstanceId *string `json:"ResourceInstanceId,omitnil,omitempty" name:"ResourceInstanceId"`

	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`
}

type DescribeBillingResourceInstanceRunningJobsRequest struct {
	*tchttp.BaseRequest
	
	// 资源组id
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// 资源组节点id
	ResourceInstanceId *string `json:"ResourceInstanceId,omitnil,omitempty" name:"ResourceInstanceId"`

	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`
}

func (r *DescribeBillingResourceInstanceRunningJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillingResourceInstanceRunningJobsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceGroupId")
	delete(f, "ResourceInstanceId")
	delete(f, "TiProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillingResourceInstanceRunningJobsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillingResourceInstanceRunningJobsResponseParams struct {
	// 资源组节点运行中的任务信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceInstanceRunningJobInfos []*ResourceInstanceRunningJobInfo `json:"ResourceInstanceRunningJobInfos,omitnil,omitempty" name:"ResourceInstanceRunningJobInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBillingResourceInstanceRunningJobsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillingResourceInstanceRunningJobsResponseParams `json:"Response"`
}

func (r *DescribeBillingResourceInstanceRunningJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillingResourceInstanceRunningJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillingSpecsPriceRequestParams struct {
	// 询价参数，支持批量询价
	SpecsParam []*SpecUnit `json:"SpecsParam,omitnil,omitempty" name:"SpecsParam"`
}

type DescribeBillingSpecsPriceRequest struct {
	*tchttp.BaseRequest
	
	// 询价参数，支持批量询价
	SpecsParam []*SpecUnit `json:"SpecsParam,omitnil,omitempty" name:"SpecsParam"`
}

func (r *DescribeBillingSpecsPriceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillingSpecsPriceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpecsParam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillingSpecsPriceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillingSpecsPriceResponseParams struct {
	// 计费项价格，支持批量返回
	SpecsPrice []*SpecPrice `json:"SpecsPrice,omitnil,omitempty" name:"SpecsPrice"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBillingSpecsPriceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillingSpecsPriceResponseParams `json:"Response"`
}

func (r *DescribeBillingSpecsPriceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillingSpecsPriceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillingSpecsRequestParams struct {
	// 付费模式：POSTPAID_BY_HOUR按量计费、PREPAID包年包月
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 枚举值：空、TRAIN、NOTEBOOK、INFERENCE或EMS
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 资源类型：["", "CALC", "CPU", "GPU", "GPU-SW"]
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`
}

type DescribeBillingSpecsRequest struct {
	*tchttp.BaseRequest
	
	// 付费模式：POSTPAID_BY_HOUR按量计费、PREPAID包年包月
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 枚举值：空、TRAIN、NOTEBOOK、INFERENCE或EMS
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 资源类型：["", "CALC", "CPU", "GPU", "GPU-SW"]
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`
}

func (r *DescribeBillingSpecsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillingSpecsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChargeType")
	delete(f, "TaskType")
	delete(f, "ResourceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillingSpecsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillingSpecsResponseParams struct {
	// 计费项列表
	Specs []*Spec `json:"Specs,omitnil,omitempty" name:"Specs"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBillingSpecsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillingSpecsResponseParams `json:"Response"`
}

func (r *DescribeBillingSpecsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillingSpecsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBuildInImagesRequestParams struct {
	// 镜像过滤器
	ImageFilters []*ImageFIlter `json:"ImageFilters,omitnil,omitempty" name:"ImageFilters"`
}

type DescribeBuildInImagesRequest struct {
	*tchttp.BaseRequest
	
	// 镜像过滤器
	ImageFilters []*ImageFIlter `json:"ImageFilters,omitnil,omitempty" name:"ImageFilters"`
}

func (r *DescribeBuildInImagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBuildInImagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageFilters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBuildInImagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBuildInImagesResponseParams struct {
	// 内置镜像详情列表
	BuildInImageInfos []*ImageInfo `json:"BuildInImageInfos,omitnil,omitempty" name:"BuildInImageInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBuildInImagesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBuildInImagesResponseParams `json:"Response"`
}

func (r *DescribeBuildInImagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBuildInImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataSourceRequestParams struct {
	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`

	// 数据源id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type DescribeDataSourceRequest struct {
	*tchttp.BaseRequest
	
	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`

	// 数据源id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

func (r *DescribeDataSourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataSourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TiProjectId")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataSourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataSourceResponseParams struct {
	// 数据源信息
	DataSourceInfo *DataSourceInfo `json:"DataSourceInfo,omitnil,omitempty" name:"DataSourceInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataSourceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataSourceResponseParams `json:"Response"`
}

func (r *DescribeDataSourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataSourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataSourcesRequestParams struct {
	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 标签过滤条件
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序的依据字段
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 输出列表的排列顺序。取值范围：ASC：升序排列 DESC：降序排列
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

type DescribeDataSourcesRequest struct {
	*tchttp.BaseRequest
	
	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 标签过滤条件
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序的依据字段
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 输出列表的排列顺序。取值范围：ASC：升序排列 DESC：降序排列
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

func (r *DescribeDataSourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataSourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TiProjectId")
	delete(f, "Filters")
	delete(f, "TagFilters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderField")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataSourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataSourcesResponseParams struct {
	// 数据源列表
	DataSourceInfos []*DataSourceInfo `json:"DataSourceInfos,omitnil,omitempty" name:"DataSourceInfos"`

	// 总条数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataSourcesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataSourcesResponseParams `json:"Response"`
}

func (r *DescribeDataSourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataSourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatasetsRequestParams struct {
	// 数据集id列表
	DatasetIds []*string `json:"DatasetIds,omitnil,omitempty" name:"DatasetIds"`

	// 数据集查询过滤条件，多个Filter之间的关系为逻辑与（AND）关系，过滤字段Filter.Name，类型为String
	// DatasetName，数据集名称
	// DatasetScope，数据集范围，SCOPE_DATASET_PRIVATE或SCOPE_DATASET_PUBLIC
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 标签过滤条件
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`

	// 排序值，支持Asc或Desc，默认Desc
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段，支持CreateTime或UpdateTime，默认CreateTime
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 偏移值
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数据个数，默认20，最大支持200
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 是否检查CFS。若开启，则在CFS挂载好之前，不会返回数据集列表。
	CFSChecking *bool `json:"CFSChecking,omitnil,omitempty" name:"CFSChecking"`

	// 是否返回CFS详情。
	CFSDetail *bool `json:"CFSDetail,omitnil,omitempty" name:"CFSDetail"`
}

type DescribeDatasetsRequest struct {
	*tchttp.BaseRequest
	
	// 数据集id列表
	DatasetIds []*string `json:"DatasetIds,omitnil,omitempty" name:"DatasetIds"`

	// 数据集查询过滤条件，多个Filter之间的关系为逻辑与（AND）关系，过滤字段Filter.Name，类型为String
	// DatasetName，数据集名称
	// DatasetScope，数据集范围，SCOPE_DATASET_PRIVATE或SCOPE_DATASET_PUBLIC
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 标签过滤条件
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`

	// 排序值，支持Asc或Desc，默认Desc
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段，支持CreateTime或UpdateTime，默认CreateTime
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 偏移值
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数据个数，默认20，最大支持200
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 是否检查CFS。若开启，则在CFS挂载好之前，不会返回数据集列表。
	CFSChecking *bool `json:"CFSChecking,omitnil,omitempty" name:"CFSChecking"`

	// 是否返回CFS详情。
	CFSDetail *bool `json:"CFSDetail,omitnil,omitempty" name:"CFSDetail"`
}

func (r *DescribeDatasetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatasetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatasetIds")
	delete(f, "Filters")
	delete(f, "TagFilters")
	delete(f, "Order")
	delete(f, "OrderField")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "CFSChecking")
	delete(f, "CFSDetail")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatasetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatasetsResponseParams struct {
	// 数据集总量（名称维度）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 数据集按照数据集名称聚合的分组
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetGroups []*DatasetGroup `json:"DatasetGroups,omitnil,omitempty" name:"DatasetGroups"`

	// 数据集ID总量
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetIdNums *uint64 `json:"DatasetIdNums,omitnil,omitempty" name:"DatasetIdNums"`

	// 若开启了CFSChecking，则检查CFS是否准备完毕。若CFS未准备完毕，则返回true，并且TotalCount为0，DatasetGroups为空。
	CFSNotReady *bool `json:"CFSNotReady,omitnil,omitempty" name:"CFSNotReady"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDatasetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatasetsResponseParams `json:"Response"`
}

func (r *DescribeDatasetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatasetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEventsRequestParams struct {
	// 服务类型，TRAIN为任务式建模, NOTEBOOK为Notebook, INFER为在线服务, BATCH为批量预测
	// 枚举值：
	// - TRAIN
	// - NOTEBOOK
	// - INFER
	// - BATCH
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// 服务ID，和Service参数对应，不同Service的服务ID获取方式不同，具体如下：
	// - Service类型为TRAIN：
	//   调用[DescribeTrainingTask接口](/document/product/851/75089)查询训练任务详情，ServiceId为接口返回值中Response.TrainingTaskDetail.LatestInstanceId
	// - Service类型为NOTEBOOK：
	//   调用[DescribeNotebook接口](/document/product/851/95662)查询Notebook详情，ServiceId为接口返回值中Response.NotebookDetail.PodName
	// - Service类型为INFER：
	//   调用[DescribeModelServiceGroup接口](/document/product/851/82285)查询服务组详情，ServiceId为接口返回值中Response.ServiceGroup.Services.ServiceId
	// - Service类型为BATCH：
	//   调用[DescribeBatchTask接口](/document/product/851/80180)查询跑批任务详情，ServiceId为接口返回值中Response.BatchTaskDetail.LatestInstanceId
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 查询事件最早发生的时间（RFC3339格式的时间字符串），默认值为当前时间的前一天
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询事件最晚发生的时间（RFC3339格式的时间字符串），默认值为当前时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页Limit，默认值为100，最大值为100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页Offset，默认值为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排列顺序（可选值为ASC, DESC ），默认为DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序的依据字段（可选值为FirstTimestamp, LastTimestamp），默认值为LastTimestamp
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 过滤条件
	// 注意: 
	// 1. Filter.Name：目前支持ResourceKind（按事件关联的资源类型过滤）；Type（按事件类型过滤）
	// 2. Filter.Values：
	// 对于Name为ResourceKind，Values的可选取值为Deployment, Replicaset, Pod等K8S资源类型；
	// 对于Name为Type，Values的可选取值仅为Normal或者Warning；
	// Values为多个的时候表示同时满足
	// 3. Filter. Negative和Filter. Fuzzy没有使用
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeEventsRequest struct {
	*tchttp.BaseRequest
	
	// 服务类型，TRAIN为任务式建模, NOTEBOOK为Notebook, INFER为在线服务, BATCH为批量预测
	// 枚举值：
	// - TRAIN
	// - NOTEBOOK
	// - INFER
	// - BATCH
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// 服务ID，和Service参数对应，不同Service的服务ID获取方式不同，具体如下：
	// - Service类型为TRAIN：
	//   调用[DescribeTrainingTask接口](/document/product/851/75089)查询训练任务详情，ServiceId为接口返回值中Response.TrainingTaskDetail.LatestInstanceId
	// - Service类型为NOTEBOOK：
	//   调用[DescribeNotebook接口](/document/product/851/95662)查询Notebook详情，ServiceId为接口返回值中Response.NotebookDetail.PodName
	// - Service类型为INFER：
	//   调用[DescribeModelServiceGroup接口](/document/product/851/82285)查询服务组详情，ServiceId为接口返回值中Response.ServiceGroup.Services.ServiceId
	// - Service类型为BATCH：
	//   调用[DescribeBatchTask接口](/document/product/851/80180)查询跑批任务详情，ServiceId为接口返回值中Response.BatchTaskDetail.LatestInstanceId
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 查询事件最早发生的时间（RFC3339格式的时间字符串），默认值为当前时间的前一天
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询事件最晚发生的时间（RFC3339格式的时间字符串），默认值为当前时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页Limit，默认值为100，最大值为100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页Offset，默认值为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排列顺序（可选值为ASC, DESC ），默认为DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序的依据字段（可选值为FirstTimestamp, LastTimestamp），默认值为LastTimestamp
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 过滤条件
	// 注意: 
	// 1. Filter.Name：目前支持ResourceKind（按事件关联的资源类型过滤）；Type（按事件类型过滤）
	// 2. Filter.Values：
	// 对于Name为ResourceKind，Values的可选取值为Deployment, Replicaset, Pod等K8S资源类型；
	// 对于Name为Type，Values的可选取值仅为Normal或者Warning；
	// Values为多个的时候表示同时满足
	// 3. Filter. Negative和Filter. Fuzzy没有使用
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Service")
	delete(f, "ServiceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "OrderField")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEventsResponseParams struct {
	// 事件的列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Events []*Event `json:"Events,omitnil,omitempty" name:"Events"`

	// 此次查询的事件的个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEventsResponseParams `json:"Response"`
}

func (r *DescribeEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExportRequestParams struct {
	// 日志下载任务的ID
	ExportId *string `json:"ExportId,omitnil,omitempty" name:"ExportId"`
}

type DescribeExportRequest struct {
	*tchttp.BaseRequest
	
	// 日志下载任务的ID
	ExportId *string `json:"ExportId,omitnil,omitempty" name:"ExportId"`
}

func (r *DescribeExportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ExportId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExportResponseParams struct {
	// 日志下载任务的ID
	ExportId *string `json:"ExportId,omitnil,omitempty" name:"ExportId"`

	// 日志下载文件名
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 日志导出路径,有效期一个小时，请尽快使用该路径下载。
	CosPath *string `json:"CosPath,omitnil,omitempty" name:"CosPath"`

	// 下载任务创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 日志文件大小
	FileSize *string `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// 日志下载状态。Processing:导出正在进行中，Completed:导出完成，Failed:导出失败，Expired:日志导出已过期(三天有效期), Queuing 排队中
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeExportResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExportResponseParams `json:"Response"`
}

func (r *DescribeExportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInferTemplatesRequestParams struct {

}

type DescribeInferTemplatesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeInferTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInferTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInferTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInferTemplatesResponseParams struct {
	// 模板列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrameworkTemplates []*InferTemplateGroup `json:"FrameworkTemplates,omitnil,omitempty" name:"FrameworkTemplates"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInferTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInferTemplatesResponseParams `json:"Response"`
}

func (r *DescribeInferTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInferTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogsRequestParams struct {
	// <p>服务类型，TRAIN为任务式建模, NOTEBOOK为Notebook, INFER为在线服务, BATCH为批量预测<br>枚举值：</p><ul><li>TRAIN</li><li>NOTEBOOK</li><li>INFER</li><li>BATCH</li></ul>
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// <p>日志查询开始时间（RFC3339格式的时间字符串），默认值为当前时间的前一个小时</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>日志查询结束时间（RFC3339格式的时间字符串），默认值为当前时间</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>日志查询条数，默认值100，最大值1000</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>服务ID，和Service参数对应，不同Service的服务ID获取方式不同，具体如下：</p><ul><li>Service类型为TRAIN：<br>调用<a href="/document/product/851/75089">DescribeTrainingTask接口</a>查询训练任务详情，ServiceId为接口返回值中Response.TrainingTaskDetail.LatestInstanceId</li><li>Service类型为NOTEBOOK：<br>调用<a href="/document/product/851/95662">DescribeNotebook接口</a>查询Notebook详情，ServiceId为接口返回值中Response.NotebookDetail.PodName</li><li>Service类型为INFER：<br>调用<a href="/document/product/851/82285">DescribeModelServiceGroup接口</a>查询服务组详情，ServiceId为接口返回值中Response.ServiceGroup.Services.ServiceId</li><li>Service类型为BATCH：<br>调用<a href="/document/product/851/80180">DescribeBatchTask接口</a>查询跑批任务详情，ServiceId为接口返回值中Response.BatchTaskDetail.LatestInstanceId</li></ul>
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// <p>Pod的名称，即需要查询服务对应的Pod，和Service参数对应，不同Service的PodName获取方式不同，具体如下：</p><ul><li>Service类型为TRAIN：<br>调用<a href="/document/product/851/75088">DescribeTrainingTaskPods接口</a>查询训练任务pod列表，PodName为接口返回值中Response.PodNames</li><li>Service类型为NOTEBOOK：<br>调用<a href="/document/product/851/95662">DescribeNotebook接口</a>查询Notebook详情，PodName为接口返回值中Response.NotebookDetail.PodName</li><li>Service类型为INFER：<br>调用<a href="/document/product/851/82287">DescribeModelService接口</a>查询单个服务详情，PodName为接口返回值中Response.Service.ServiceInfo.PodInfos</li><li>Service类型为BATCH：<br>调用<a href="/document/product/851/80180">DescribeBatchTask接口</a>查询跑批任务详情，PodName为接口返回值中Response.BatchTaskDetail. PodList<br>注：支持结尾通配符*</li></ul>
	PodName *string `json:"PodName,omitnil,omitempty" name:"PodName"`

	// <p>排序方向（可选值为ASC, DESC ），默认为DESC</p>
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// <p>按哪个字段排序（可选值为Timestamp），默认值为Timestamp</p>
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// <p>日志查询上下文，查询下一页的时候需要回传这个字段，该字段来自本接口的返回</p>
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// <p>过滤条件<br>注意: </p><ol><li>Filter.Name：目前只支持Key（也就是按关键字过滤日志）</li><li>Filter.Values：表示过滤日志的关键字；Values为多个的时候表示同时满足</li><li>Filter. Negative和Filter. Fuzzy没有使用</li></ol>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>使用OFFSET分页查询时，指定返回的数据偏移量，默认为0</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeLogsRequest struct {
	*tchttp.BaseRequest
	
	// <p>服务类型，TRAIN为任务式建模, NOTEBOOK为Notebook, INFER为在线服务, BATCH为批量预测<br>枚举值：</p><ul><li>TRAIN</li><li>NOTEBOOK</li><li>INFER</li><li>BATCH</li></ul>
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// <p>日志查询开始时间（RFC3339格式的时间字符串），默认值为当前时间的前一个小时</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>日志查询结束时间（RFC3339格式的时间字符串），默认值为当前时间</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>日志查询条数，默认值100，最大值1000</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>服务ID，和Service参数对应，不同Service的服务ID获取方式不同，具体如下：</p><ul><li>Service类型为TRAIN：<br>调用<a href="/document/product/851/75089">DescribeTrainingTask接口</a>查询训练任务详情，ServiceId为接口返回值中Response.TrainingTaskDetail.LatestInstanceId</li><li>Service类型为NOTEBOOK：<br>调用<a href="/document/product/851/95662">DescribeNotebook接口</a>查询Notebook详情，ServiceId为接口返回值中Response.NotebookDetail.PodName</li><li>Service类型为INFER：<br>调用<a href="/document/product/851/82285">DescribeModelServiceGroup接口</a>查询服务组详情，ServiceId为接口返回值中Response.ServiceGroup.Services.ServiceId</li><li>Service类型为BATCH：<br>调用<a href="/document/product/851/80180">DescribeBatchTask接口</a>查询跑批任务详情，ServiceId为接口返回值中Response.BatchTaskDetail.LatestInstanceId</li></ul>
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// <p>Pod的名称，即需要查询服务对应的Pod，和Service参数对应，不同Service的PodName获取方式不同，具体如下：</p><ul><li>Service类型为TRAIN：<br>调用<a href="/document/product/851/75088">DescribeTrainingTaskPods接口</a>查询训练任务pod列表，PodName为接口返回值中Response.PodNames</li><li>Service类型为NOTEBOOK：<br>调用<a href="/document/product/851/95662">DescribeNotebook接口</a>查询Notebook详情，PodName为接口返回值中Response.NotebookDetail.PodName</li><li>Service类型为INFER：<br>调用<a href="/document/product/851/82287">DescribeModelService接口</a>查询单个服务详情，PodName为接口返回值中Response.Service.ServiceInfo.PodInfos</li><li>Service类型为BATCH：<br>调用<a href="/document/product/851/80180">DescribeBatchTask接口</a>查询跑批任务详情，PodName为接口返回值中Response.BatchTaskDetail. PodList<br>注：支持结尾通配符*</li></ul>
	PodName *string `json:"PodName,omitnil,omitempty" name:"PodName"`

	// <p>排序方向（可选值为ASC, DESC ），默认为DESC</p>
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// <p>按哪个字段排序（可选值为Timestamp），默认值为Timestamp</p>
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// <p>日志查询上下文，查询下一页的时候需要回传这个字段，该字段来自本接口的返回</p>
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// <p>过滤条件<br>注意: </p><ol><li>Filter.Name：目前只支持Key（也就是按关键字过滤日志）</li><li>Filter.Values：表示过滤日志的关键字；Values为多个的时候表示同时满足</li><li>Filter. Negative和Filter. Fuzzy没有使用</li></ol>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>使用OFFSET分页查询时，指定返回的数据偏移量，默认为0</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Service")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "ServiceId")
	delete(f, "PodName")
	delete(f, "Order")
	delete(f, "OrderField")
	delete(f, "Context")
	delete(f, "Filters")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogsResponseParams struct {
	// <p>分页的游标</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// <p>日志数组</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*LogIdentity `json:"Content,omitnil,omitempty" name:"Content"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLogsResponseParams `json:"Response"`
}

func (r *DescribeLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelAccelerateTaskRequestParams struct {
	// 模型加速任务ID
	ModelAccTaskId *string `json:"ModelAccTaskId,omitnil,omitempty" name:"ModelAccTaskId"`
}

type DescribeModelAccelerateTaskRequest struct {
	*tchttp.BaseRequest
	
	// 模型加速任务ID
	ModelAccTaskId *string `json:"ModelAccTaskId,omitnil,omitempty" name:"ModelAccTaskId"`
}

func (r *DescribeModelAccelerateTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelAccelerateTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelAccTaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelAccelerateTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelAccelerateTaskResponseParams struct {
	// 模型加速任务详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelAccelerateTask *ModelAccelerateTask `json:"ModelAccelerateTask,omitnil,omitempty" name:"ModelAccelerateTask"`

	// 模型加速时长，单位s
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelAccRuntimeInSecond *uint64 `json:"ModelAccRuntimeInSecond,omitnil,omitempty" name:"ModelAccRuntimeInSecond"`

	// 模型加速任务开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelAccStartTime *string `json:"ModelAccStartTime,omitnil,omitempty" name:"ModelAccStartTime"`

	// 模型加速任务结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelAccEndTime *string `json:"ModelAccEndTime,omitnil,omitempty" name:"ModelAccEndTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModelAccelerateTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelAccelerateTaskResponseParams `json:"Response"`
}

func (r *DescribeModelAccelerateTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelAccelerateTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelAccelerateVersionsRequestParams struct {
	// 过滤条件
	//     Filter.Name: 枚举值: ModelJobName (任务名称)|TrainingModelVersionId (模型版本id)
	//     Filter.Values: 当长度为1时，支持模糊查询; 不为1时，精确查询
	// 每次请求的Filters的上限为10，Filter.Values的上限为100
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段; 枚举值: CreateTime (创建时间) ；默认CreateTime
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 排序方向; 枚举值: ASC | DESC；默认DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 分页查询起始位置，如：Limit为100，第一页Offset为0，第二页Offset为100...即每页左边为闭区间; 默认0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询每页大小，最大20000; 默认10
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 模型ID
	TrainingModelId *string `json:"TrainingModelId,omitnil,omitempty" name:"TrainingModelId"`
}

type DescribeModelAccelerateVersionsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件
	//     Filter.Name: 枚举值: ModelJobName (任务名称)|TrainingModelVersionId (模型版本id)
	//     Filter.Values: 当长度为1时，支持模糊查询; 不为1时，精确查询
	// 每次请求的Filters的上限为10，Filter.Values的上限为100
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段; 枚举值: CreateTime (创建时间) ；默认CreateTime
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 排序方向; 枚举值: ASC | DESC；默认DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 分页查询起始位置，如：Limit为100，第一页Offset为0，第二页Offset为100...即每页左边为闭区间; 默认0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询每页大小，最大20000; 默认10
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 模型ID
	TrainingModelId *string `json:"TrainingModelId,omitnil,omitempty" name:"TrainingModelId"`
}

func (r *DescribeModelAccelerateVersionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelAccelerateVersionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "OrderField")
	delete(f, "Order")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "TrainingModelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelAccelerateVersionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelAccelerateVersionsResponseParams struct {
	// 优化模型总数； 注意接口是分页拉取的，total是指优化模型节点总数，不是本次返回中ModelAccelerateVersions数组的大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 优化模型列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelAccelerateVersions []*ModelAccelerateVersion `json:"ModelAccelerateVersions,omitnil,omitempty" name:"ModelAccelerateVersions"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModelAccelerateVersionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelAccelerateVersionsResponseParams `json:"Response"`
}

func (r *DescribeModelAccelerateVersionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelAccelerateVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelServiceCallInfoRequestParams struct {
	// 服务组id
	ServiceGroupId *string `json:"ServiceGroupId,omitnil,omitempty" name:"ServiceGroupId"`

	// 服务分类
	//
	// Deprecated: ServiceCategory is deprecated.
	ServiceCategory *string `json:"ServiceCategory,omitnil,omitempty" name:"ServiceCategory"`
}

type DescribeModelServiceCallInfoRequest struct {
	*tchttp.BaseRequest
	
	// 服务组id
	ServiceGroupId *string `json:"ServiceGroupId,omitnil,omitempty" name:"ServiceGroupId"`

	// 服务分类
	ServiceCategory *string `json:"ServiceCategory,omitnil,omitempty" name:"ServiceCategory"`
}

func (r *DescribeModelServiceCallInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelServiceCallInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceGroupId")
	delete(f, "ServiceCategory")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelServiceCallInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelServiceCallInfoResponseParams struct {
	// 服务调用信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceCallInfo *ServiceCallInfo `json:"ServiceCallInfo,omitnil,omitempty" name:"ServiceCallInfo"`

	// 升级网关调用信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InferGatewayCallInfo *InferGatewayCallInfo `json:"InferGatewayCallInfo,omitnil,omitempty" name:"InferGatewayCallInfo"`

	// 默认nginx网关的调用信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultNginxGatewayCallInfo *DefaultNginxGatewayCallInfo `json:"DefaultNginxGatewayCallInfo,omitnil,omitempty" name:"DefaultNginxGatewayCallInfo"`

	// 太极服务的调用信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TJCallInfo *TJCallInfo `json:"TJCallInfo,omitnil,omitempty" name:"TJCallInfo"`

	// 内网调用信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntranetCallInfo *IntranetCallInfo `json:"IntranetCallInfo,omitnil,omitempty" name:"IntranetCallInfo"`

	// 基于新网关的服务调用信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceCallInfoV2 *ServiceCallInfoV2 `json:"ServiceCallInfoV2,omitnil,omitempty" name:"ServiceCallInfoV2"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModelServiceCallInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelServiceCallInfoResponseParams `json:"Response"`
}

func (r *DescribeModelServiceCallInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelServiceCallInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelServiceGroupRequestParams struct {
	// 服务组ID
	ServiceGroupId *string `json:"ServiceGroupId,omitnil,omitempty" name:"ServiceGroupId"`

	// 服务分类
	//
	// Deprecated: ServiceCategory is deprecated.
	ServiceCategory *string `json:"ServiceCategory,omitnil,omitempty" name:"ServiceCategory"`
}

type DescribeModelServiceGroupRequest struct {
	*tchttp.BaseRequest
	
	// 服务组ID
	ServiceGroupId *string `json:"ServiceGroupId,omitnil,omitempty" name:"ServiceGroupId"`

	// 服务分类
	ServiceCategory *string `json:"ServiceCategory,omitnil,omitempty" name:"ServiceCategory"`
}

func (r *DescribeModelServiceGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelServiceGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceGroupId")
	delete(f, "ServiceCategory")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelServiceGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelServiceGroupResponseParams struct {
	// 服务组信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceGroup *ServiceGroup `json:"ServiceGroup,omitnil,omitempty" name:"ServiceGroup"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModelServiceGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelServiceGroupResponseParams `json:"Response"`
}

func (r *DescribeModelServiceGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelServiceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelServiceGroupsRequestParams struct {
	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`

	// <p>偏移量，默认为0</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>返回数量，默认为20，最大值为100</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>输出列表的排列顺序。取值范围：ASC：升序排列 DESC：降序排列</p>
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// <p>排序的依据字段， 取值范围 &quot;CreateTime&quot; &quot;UpdateTime&quot;</p>
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// <p>分页参数，支持的分页过滤Name包括：<br>[&quot;ClusterId&quot;, &quot;ServiceId&quot;, &quot;ServiceGroupName&quot;, &quot;ServiceGroupId&quot;,&quot;Status&quot;,&quot;CreatedBy&quot;,&quot;ModelVersionId&quot;]</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>标签过滤参数</p>
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`

	// <p>服务分类</p>
	//
	// Deprecated: ServiceCategory is deprecated.
	ServiceCategory *string `json:"ServiceCategory,omitnil,omitempty" name:"ServiceCategory"`
}

type DescribeModelServiceGroupsRequest struct {
	*tchttp.BaseRequest
	
	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`

	// <p>偏移量，默认为0</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>返回数量，默认为20，最大值为100</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>输出列表的排列顺序。取值范围：ASC：升序排列 DESC：降序排列</p>
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// <p>排序的依据字段， 取值范围 &quot;CreateTime&quot; &quot;UpdateTime&quot;</p>
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// <p>分页参数，支持的分页过滤Name包括：<br>[&quot;ClusterId&quot;, &quot;ServiceId&quot;, &quot;ServiceGroupName&quot;, &quot;ServiceGroupId&quot;,&quot;Status&quot;,&quot;CreatedBy&quot;,&quot;ModelVersionId&quot;]</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>标签过滤参数</p>
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`

	// <p>服务分类</p>
	ServiceCategory *string `json:"ServiceCategory,omitnil,omitempty" name:"ServiceCategory"`
}

func (r *DescribeModelServiceGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelServiceGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TiProjectId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "OrderField")
	delete(f, "Filters")
	delete(f, "TagFilters")
	delete(f, "ServiceCategory")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelServiceGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelServiceGroupsResponseParams struct {
	// <p>推理服务组数量。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>服务组信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceGroups []*ServiceGroup `json:"ServiceGroups,omitnil,omitempty" name:"ServiceGroups"`

	// <p>当前uin和region下全量服务组数量</p>
	GlobalTotalCount *int64 `json:"GlobalTotalCount,omitnil,omitempty" name:"GlobalTotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModelServiceGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelServiceGroupsResponseParams `json:"Response"`
}

func (r *DescribeModelServiceGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelServiceGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelServiceHotUpdatedRequestParams struct {
	// 镜像信息，配置服务运行所需的镜像地址等信息
	ImageInfo *ImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`

	// 模型信息，需要挂载模型时填写
	ModelInfo *ModelInfo `json:"ModelInfo,omitnil,omitempty" name:"ModelInfo"`

	// 挂载信息
	VolumeMount *VolumeMount `json:"VolumeMount,omitnil,omitempty" name:"VolumeMount"`
}

type DescribeModelServiceHotUpdatedRequest struct {
	*tchttp.BaseRequest
	
	// 镜像信息，配置服务运行所需的镜像地址等信息
	ImageInfo *ImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`

	// 模型信息，需要挂载模型时填写
	ModelInfo *ModelInfo `json:"ModelInfo,omitnil,omitempty" name:"ModelInfo"`

	// 挂载信息
	VolumeMount *VolumeMount `json:"VolumeMount,omitnil,omitempty" name:"VolumeMount"`
}

func (r *DescribeModelServiceHotUpdatedRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelServiceHotUpdatedRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageInfo")
	delete(f, "ModelInfo")
	delete(f, "VolumeMount")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelServiceHotUpdatedRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelServiceHotUpdatedResponseParams struct {
	// 模型加速标志位.Allowed 允许模型加速. Forbidden 禁止模型加速
	ModelTurboFlag *string `json:"ModelTurboFlag,omitnil,omitempty" name:"ModelTurboFlag"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModelServiceHotUpdatedResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelServiceHotUpdatedResponseParams `json:"Response"`
}

func (r *DescribeModelServiceHotUpdatedResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelServiceHotUpdatedResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelServiceRequestParams struct {
	// 服务id
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 服务分类
	//
	// Deprecated: ServiceCategory is deprecated.
	ServiceCategory *string `json:"ServiceCategory,omitnil,omitempty" name:"ServiceCategory"`
}

type DescribeModelServiceRequest struct {
	*tchttp.BaseRequest
	
	// 服务id
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 服务分类
	ServiceCategory *string `json:"ServiceCategory,omitnil,omitempty" name:"ServiceCategory"`
}

func (r *DescribeModelServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "ServiceCategory")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelServiceResponseParams struct {
	// 服务信息
	Service *Service `json:"Service,omitnil,omitempty" name:"Service"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModelServiceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelServiceResponseParams `json:"Response"`
}

func (r *DescribeModelServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMountInstanceRequestParams struct {
	// 数据源类型英文名
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`

	// 存储实例ID
	StorageId *string `json:"StorageId,omitnil,omitempty" name:"StorageId"`
}

type DescribeMountInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 数据源类型英文名
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`

	// 存储实例ID
	StorageId *string `json:"StorageId,omitnil,omitempty" name:"StorageId"`
}

func (r *DescribeMountInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMountInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "TiProjectId")
	delete(f, "StorageId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMountInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMountInstanceResponseParams struct {
	// 挂载的实例详情
	MountInstance *MountInstanceInfo `json:"MountInstance,omitnil,omitempty" name:"MountInstance"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMountInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMountInstanceResponseParams `json:"Response"`
}

func (r *DescribeMountInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMountInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMountInstancesRequestParams struct {
	// 数据源类型英文名
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeMountInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 数据源类型英文名
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeMountInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMountInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "TiProjectId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMountInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMountInstancesResponseParams struct {
	// 挂载的实例列表
	MountInstances []*MountInstanceInfo `json:"MountInstances,omitnil,omitempty" name:"MountInstances"`

	// 总条数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMountInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMountInstancesResponseParams `json:"Response"`
}

func (r *DescribeMountInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMountInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMountLimitsRequestParams struct {
	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 输出列表的排列顺序。取值范围：ASC：升序排列 DESC：降序排列
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序的依据字段
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

type DescribeMountLimitsRequest struct {
	*tchttp.BaseRequest
	
	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 输出列表的排列顺序。取值范围：ASC：升序排列 DESC：降序排列
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序的依据字段
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

func (r *DescribeMountLimitsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMountLimitsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TiProjectId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "OrderField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMountLimitsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMountLimitsResponseParams struct {
	// 挂载限制列表
	MountLimits []*MountLimitInfo `json:"MountLimits,omitnil,omitempty" name:"MountLimits"`

	// 总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMountLimitsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMountLimitsResponseParams `json:"Response"`
}

func (r *DescribeMountLimitsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMountLimitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebookRequestParams struct {
	// notebook id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`
}

type DescribeNotebookRequest struct {
	*tchttp.BaseRequest
	
	// notebook id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`
}

func (r *DescribeNotebookRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebookRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "TiProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNotebookRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebookResponseParams struct {
	// 详情
	NotebookDetail *NotebookDetail `json:"NotebookDetail,omitnil,omitempty" name:"NotebookDetail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNotebookResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNotebookResponseParams `json:"Response"`
}

func (r *DescribeNotebookResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebookResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebooksRequestParams struct {
	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页返回的实例数，默认为10
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 输出列表的排列顺序。取值范围：ASC：升序排列 DESC：降序排列。默认为DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 根据哪个字段排序，如：CreateTime、UpdateTime，默认为UpdateTime
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 过滤器，eg：[{ "Name": "Id", "Values": ["nb-123456789"] }]
	// 
	// 取值范围
	// Name（名称）：notebook1
	// Id（notebook ID）：nb-123456789
	// Status（状态）：Starting / Running / Stopped / Stopping / Failed / SubmitFailed
	// Creator（创建者 uin）：100014761913
	// ChargeType（计费类型）：PREPAID（预付费）/ POSTPAID_BY_HOUR（后付费）
	// ChargeStatus（计费状态）：NOT_BILLING（未开始计费）/ BILLING（计费中）/ BILLING_STORAGE（存储计费中）/ARREARS_STOP（欠费停止）
	// DefaultCodeRepoId（默认代码仓库ID）：cr-123456789
	// AdditionalCodeRepoId（关联代码仓库ID）：cr-123456789
	// LifecycleScriptId（生命周期ID）：ls-12312312311312
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 标签过滤器，eg：[{ "TagKey": "TagKeyA", "TagValue": ["TagValueA"] }]
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`
}

type DescribeNotebooksRequest struct {
	*tchttp.BaseRequest
	
	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页返回的实例数，默认为10
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 输出列表的排列顺序。取值范围：ASC：升序排列 DESC：降序排列。默认为DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 根据哪个字段排序，如：CreateTime、UpdateTime，默认为UpdateTime
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 过滤器，eg：[{ "Name": "Id", "Values": ["nb-123456789"] }]
	// 
	// 取值范围
	// Name（名称）：notebook1
	// Id（notebook ID）：nb-123456789
	// Status（状态）：Starting / Running / Stopped / Stopping / Failed / SubmitFailed
	// Creator（创建者 uin）：100014761913
	// ChargeType（计费类型）：PREPAID（预付费）/ POSTPAID_BY_HOUR（后付费）
	// ChargeStatus（计费状态）：NOT_BILLING（未开始计费）/ BILLING（计费中）/ BILLING_STORAGE（存储计费中）/ARREARS_STOP（欠费停止）
	// DefaultCodeRepoId（默认代码仓库ID）：cr-123456789
	// AdditionalCodeRepoId（关联代码仓库ID）：cr-123456789
	// LifecycleScriptId（生命周期ID）：ls-12312312311312
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 标签过滤器，eg：[{ "TagKey": "TagKeyA", "TagValue": ["TagValueA"] }]
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`
}

func (r *DescribeNotebooksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebooksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TiProjectId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "OrderField")
	delete(f, "Filters")
	delete(f, "TagFilters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNotebooksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebooksResponseParams struct {
	// notebook详情
	NotebookSet []*NotebookSetItem `json:"NotebookSet,omitnil,omitempty" name:"NotebookSet"`

	// 总条数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNotebooksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNotebooksResponseParams `json:"Response"`
}

func (r *DescribeNotebooksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebooksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePlatformImagesRequestParams struct {
	// 过滤器,  Name支持ImageId/ImageName/SupportDataPipeline/AllowSaveAllContent/ImageRange，其中ImageRange支持枚举值Train,Inference,Notebook
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移信息
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量, 默认100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribePlatformImagesRequest struct {
	*tchttp.BaseRequest
	
	// 过滤器,  Name支持ImageId/ImageName/SupportDataPipeline/AllowSaveAllContent/ImageRange，其中ImageRange支持枚举值Train,Inference,Notebook
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移信息
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量, 默认100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribePlatformImagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePlatformImagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePlatformImagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePlatformImagesResponseParams struct {
	// 数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 镜像列表
	PlatformImageInfos []*PlatformImageInfo `json:"PlatformImageInfos,omitnil,omitempty" name:"PlatformImageInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePlatformImagesResponse struct {
	*tchttp.BaseResponse
	Response *DescribePlatformImagesResponseParams `json:"Response"`
}

func (r *DescribePlatformImagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePlatformImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePublicAlgoVersionListRequestParams struct {
	// 过滤器
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认10
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 是否需要聚合
	NeedsAggregate *bool `json:"NeedsAggregate,omitnil,omitempty" name:"NeedsAggregate"`
}

type DescribePublicAlgoVersionListRequest struct {
	*tchttp.BaseRequest
	
	// 过滤器
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认10
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 是否需要聚合
	NeedsAggregate *bool `json:"NeedsAggregate,omitnil,omitempty" name:"NeedsAggregate"`
}

func (r *DescribePublicAlgoVersionListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePublicAlgoVersionListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "NeedsAggregate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePublicAlgoVersionListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePublicAlgoVersionListResponseParams struct {
	// 算法版本数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 公共算法版本列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicAlgoVersions []*PublicAlgoVersion `json:"PublicAlgoVersions,omitnil,omitempty" name:"PublicAlgoVersions"`

	// 聚合后的公共算法版本列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AggregatePublicAlgoVersions []*AggregatePublicAlgoVersion `json:"AggregatePublicAlgoVersions,omitnil,omitempty" name:"AggregatePublicAlgoVersions"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePublicAlgoVersionListResponse struct {
	*tchttp.BaseResponse
	Response *DescribePublicAlgoVersionListResponseParams `json:"Response"`
}

func (r *DescribePublicAlgoVersionListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePublicAlgoVersionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubAccountLinuxUserInfosRequestParams struct {
	// 分页偏移量（0 表示全量）
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量（0 表示全量）
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeSubAccountLinuxUserInfosRequest struct {
	*tchttp.BaseRequest
	
	// 分页偏移量（0 表示全量）
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量（0 表示全量）
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeSubAccountLinuxUserInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubAccountLinuxUserInfosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSubAccountLinuxUserInfosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubAccountLinuxUserInfosResponseParams struct {
	// 子账号信息列表
	SubAccountList []*SubAccountInfo `json:"SubAccountList,omitnil,omitempty" name:"SubAccountList"`

	// 总数（配合分页使用）
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSubAccountLinuxUserInfosResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSubAccountLinuxUserInfosResponseParams `json:"Response"`
}

func (r *DescribeSubAccountLinuxUserInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubAccountLinuxUserInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingModelVersionRequestParams struct {
	// 模型版本ID
	TrainingModelVersionId *string `json:"TrainingModelVersionId,omitnil,omitempty" name:"TrainingModelVersionId"`
}

type DescribeTrainingModelVersionRequest struct {
	*tchttp.BaseRequest
	
	// 模型版本ID
	TrainingModelVersionId *string `json:"TrainingModelVersionId,omitnil,omitempty" name:"TrainingModelVersionId"`
}

func (r *DescribeTrainingModelVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingModelVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TrainingModelVersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTrainingModelVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingModelVersionResponseParams struct {
	// 模型版本
	TrainingModelVersion *TrainingModelVersionDTO `json:"TrainingModelVersion,omitnil,omitempty" name:"TrainingModelVersion"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTrainingModelVersionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTrainingModelVersionResponseParams `json:"Response"`
}

func (r *DescribeTrainingModelVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingModelVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingModelVersionsRequestParams struct {
	// 模型ID
	TrainingModelId *string `json:"TrainingModelId,omitnil,omitempty" name:"TrainingModelId"`

	// 过滤条件
	// Filter.Name: 枚举值:
	//     TrainingModelVersionId (模型版本ID)
	//     ModelVersionType (模型版本类型) 其值支持: NORMAL(通用) ACCELERATE (加速)
	//     ModelFormat（模型格式）其值Filter.Values支持：
	// TORCH_SCRIPT/PYTORCH/DETECTRON2/SAVED_MODEL/FROZEN_GRAPH/PMML
	//     AlgorithmFramework (算法框架) 其值Filter.Values支持：TENSORFLOW/PYTORCH/DETECTRON2
	// Filter.Values: 当长度为1时，支持模糊查询; 不为1时，精确查询
	// 每次请求的Filters的上限为10，Filter.Values的上限为100
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeTrainingModelVersionsRequest struct {
	*tchttp.BaseRequest
	
	// 模型ID
	TrainingModelId *string `json:"TrainingModelId,omitnil,omitempty" name:"TrainingModelId"`

	// 过滤条件
	// Filter.Name: 枚举值:
	//     TrainingModelVersionId (模型版本ID)
	//     ModelVersionType (模型版本类型) 其值支持: NORMAL(通用) ACCELERATE (加速)
	//     ModelFormat（模型格式）其值Filter.Values支持：
	// TORCH_SCRIPT/PYTORCH/DETECTRON2/SAVED_MODEL/FROZEN_GRAPH/PMML
	//     AlgorithmFramework (算法框架) 其值Filter.Values支持：TENSORFLOW/PYTORCH/DETECTRON2
	// Filter.Values: 当长度为1时，支持模糊查询; 不为1时，精确查询
	// 每次请求的Filters的上限为10，Filter.Values的上限为100
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeTrainingModelVersionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingModelVersionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TrainingModelId")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTrainingModelVersionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingModelVersionsResponseParams struct {
	// 模型版本列表
	TrainingModelVersions []*TrainingModelVersionDTO `json:"TrainingModelVersions,omitnil,omitempty" name:"TrainingModelVersions"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTrainingModelVersionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTrainingModelVersionsResponseParams `json:"Response"`
}

func (r *DescribeTrainingModelVersionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingModelVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingTaskPodsRequestParams struct {
	// 训练任务ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`
}

type DescribeTrainingTaskPodsRequest struct {
	*tchttp.BaseRequest
	
	// 训练任务ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`
}

func (r *DescribeTrainingTaskPodsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingTaskPodsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "TiProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTrainingTaskPodsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingTaskPodsResponseParams struct {
	// pod名称列表
	PodNames []*string `json:"PodNames,omitnil,omitempty" name:"PodNames"`

	// 数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// pod详细信息
	PodInfoList []*PodInfo `json:"PodInfoList,omitnil,omitempty" name:"PodInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTrainingTaskPodsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTrainingTaskPodsResponseParams `json:"Response"`
}

func (r *DescribeTrainingTaskPodsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingTaskPodsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingTaskRequestParams struct {
	// 训练任务ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`

	// 训练任务实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeTrainingTaskRequest struct {
	*tchttp.BaseRequest
	
	// 训练任务ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`

	// 训练任务实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeTrainingTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "TiProjectId")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTrainingTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingTaskResponseParams struct {
	// 训练任务详情
	TrainingTaskDetail *TrainingTaskDetail `json:"TrainingTaskDetail,omitnil,omitempty" name:"TrainingTaskDetail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTrainingTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTrainingTaskResponseParams `json:"Response"`
}

func (r *DescribeTrainingTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingTasksRequestParams struct {
	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`

	// 过滤器，eg：[{ "Name": "Id", "Values": ["train-23091792777383936"] }]
	// 
	// 取值范围：
	// Name（名称）：task1
	// Id（task ID）：train-23091792777383936
	// Status（状态）：SUBMITTING/PENDING/STARTING / RUNNING / STOPPING / STOPPED / FAILED / SUCCEED / SUBMIT_FAILED
	// ResourceGroupId（资源组 Id）：trsg-kvvfrwl7
	// Creator（创建者 uin）：100014761913
	// ChargeType（计费类型）：PREPAID（预付费）/ POSTPAID_BY_HOUR（后付费）
	// CHARGE_STATUS（计费状态）：NOT_BILLING（未开始计费）/ BILLING（计费中）/ ARREARS_STOP（欠费停止）
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 标签过滤器，eg：[{ "TagKey": "TagKeyA", "TagValue": ["TagValueA"] }]
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为10，最大为50
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 输出列表的排列顺序。取值范围：ASC（升序排列）/ DESC（降序排列），默认为DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序的依据字段， 取值范围 "CreateTime" 、"UpdateTime"、"StartTime"，默认为UpdateTime
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

type DescribeTrainingTasksRequest struct {
	*tchttp.BaseRequest
	
	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`

	// 过滤器，eg：[{ "Name": "Id", "Values": ["train-23091792777383936"] }]
	// 
	// 取值范围：
	// Name（名称）：task1
	// Id（task ID）：train-23091792777383936
	// Status（状态）：SUBMITTING/PENDING/STARTING / RUNNING / STOPPING / STOPPED / FAILED / SUCCEED / SUBMIT_FAILED
	// ResourceGroupId（资源组 Id）：trsg-kvvfrwl7
	// Creator（创建者 uin）：100014761913
	// ChargeType（计费类型）：PREPAID（预付费）/ POSTPAID_BY_HOUR（后付费）
	// CHARGE_STATUS（计费状态）：NOT_BILLING（未开始计费）/ BILLING（计费中）/ ARREARS_STOP（欠费停止）
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 标签过滤器，eg：[{ "TagKey": "TagKeyA", "TagValue": ["TagValueA"] }]
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为10，最大为50
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 输出列表的排列顺序。取值范围：ASC（升序排列）/ DESC（降序排列），默认为DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序的依据字段， 取值范围 "CreateTime" 、"UpdateTime"、"StartTime"，默认为UpdateTime
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

func (r *DescribeTrainingTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TiProjectId")
	delete(f, "Filters")
	delete(f, "TagFilters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "OrderField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTrainingTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingTasksResponseParams struct {
	// 训练任务集
	TrainingTaskSet []*TrainingTaskSetItem `json:"TrainingTaskSet,omitnil,omitempty" name:"TrainingTaskSet"`

	// 数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTrainingTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTrainingTasksResponseParams `json:"Response"`
}

func (r *DescribeTrainingTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkspacesRequestParams struct {
	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`

	// <p>过滤条件</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>偏移量</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>数量</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>排序字段</p>
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// <p>排序方式</p>
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

type DescribeWorkspacesRequest struct {
	*tchttp.BaseRequest
	
	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`

	// <p>过滤条件</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>偏移量</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>数量</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>排序字段</p>
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// <p>排序方式</p>
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
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
	delete(f, "TiProjectId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderField")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWorkspacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkspacesResponseParams struct {
	// <p>总数</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>工作空间列表</p>
	Workspaces []*Workspace `json:"Workspaces,omitnil,omitempty" name:"Workspaces"`

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

type DeviceImageInfo struct {
	// 设备类型, 支持GPU等
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 镜像信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageInfo *ImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`
}

type DeviceMaterialInfo struct {
	// 设备信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 物料信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaterialInfo *MaterialInfo `json:"MaterialInfo,omitnil,omitempty" name:"MaterialInfo"`
}

type EncodedStartCmdInfo struct {
	// 任务的启动命令，以base64格式输入，注意转换时需要完整输入{"StartCmd":"","PsStartCmd":"","WorkerStartCmd":""}
	StartCmdInfo *string `json:"StartCmdInfo,omitnil,omitempty" name:"StartCmdInfo"`
}

type EnvVar struct {
	// 环境变量key
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 环境变量value
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type Event struct {
	// 事件的id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 事件的具体信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 事件第一次发生的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstTimestamp *string `json:"FirstTimestamp,omitnil,omitempty" name:"FirstTimestamp"`

	// 事件最后一次发生的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastTimestamp *string `json:"LastTimestamp,omitnil,omitempty" name:"LastTimestamp"`

	// 事件发生的次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 事件的类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 事件关联的资源的类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceKind *string `json:"ResourceKind,omitnil,omitempty" name:"ResourceKind"`

	// 事件关联的资源的名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`
}

type ExecAction struct {
	// 执行命令列表
	Command []*string `json:"Command,omitnil,omitempty" name:"Command"`
}

type ExposeNetworkConfig struct {
	// ssh配置
	SSHConfig *SSHConfig `json:"SSHConfig,omitnil,omitempty" name:"SSHConfig"`

	// 容器端口暴露到公网配置
	ExposePortConfig *ExposePortConfig `json:"ExposePortConfig,omitnil,omitempty" name:"ExposePortConfig"`
}

type ExposePortConfig struct {
	// 是否开启暴露容器服务端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// vpc id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// clb id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClbId *string `json:"ClbId,omitnil,omitempty" name:"ClbId"`

	// clb domain
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClbHost *string `json:"ClbHost,omitnil,omitempty" name:"ClbHost"`
}

type Filter struct {
	// 过滤字段名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 过滤字段取值
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// 是否开启反向查询
	Negative *bool `json:"Negative,omitnil,omitempty" name:"Negative"`

	// 是否开启模糊匹配
	Fuzzy *bool `json:"Fuzzy,omitnil,omitempty" name:"Fuzzy"`
}

type GatewayConfig struct {
	// 网关类型
	GatewayType *string `json:"GatewayType,omitnil,omitempty" name:"GatewayType"`

	// 网关调度算法：轮询、一致性哈希等
	SchedulingAlgorithm *string `json:"SchedulingAlgorithm,omitnil,omitempty" name:"SchedulingAlgorithm"`

	// 一致性哈希使用的Header字段名
	HashHeaderKey *string `json:"HashHeaderKey,omitnil,omitempty" name:"HashHeaderKey"`
}

type GooseFS struct {
	// goosefs实例id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// GooseFS类型，包括GooseFS和GooseFSx
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// GooseFSx实例需要挂载的路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// GooseFS命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	NameSpace *string `json:"NameSpace,omitnil,omitempty" name:"NameSpace"`
}

type GooseFSx struct {
	// goosefsx实例id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// GooseFSx实例需要挂载的路径
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`
}

type GpuDetail struct {
	// GPU 显卡类型；枚举值: V100 A100 T4
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// GPU 显卡数；单位为1/100卡，比如100代表1卡
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *uint64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type GroupResource struct {
	// CPU核数; 单位为1/1000核，比如100表示0.1核
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存；单位为MB
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 总卡数；GPUDetail 显卡数之和；单位为1/100卡，比如100代表1卡
	// 注意：此字段可能返回 null，表示取不到有效值。
	Gpu *uint64 `json:"Gpu,omitnil,omitempty" name:"Gpu"`

	// Gpu详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	GpuDetailSet []*GpuDetail `json:"GpuDetailSet,omitnil,omitempty" name:"GpuDetailSet"`
}

type HDFSConfig struct {
	// 集群实例ID,实例ID形如: emr-xxxxxxxx
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 路径
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`
}

type HTTPGetAction struct {
	// http 路径
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 调用端口
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`
}

type HealthProbe struct {
	// 存活探针
	LivenessProbe *Probe `json:"LivenessProbe,omitnil,omitempty" name:"LivenessProbe"`

	// 就绪探针
	ReadinessProbe *Probe `json:"ReadinessProbe,omitnil,omitempty" name:"ReadinessProbe"`

	// 启动探针
	StartupProbe *Probe `json:"StartupProbe,omitnil,omitempty" name:"StartupProbe"`
}

type HorizontalPodAutoscaler struct {
	// 最小实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinReplicas *int64 `json:"MinReplicas,omitnil,omitempty" name:"MinReplicas"`

	// 最大实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxReplicas *int64 `json:"MaxReplicas,omitnil,omitempty" name:"MaxReplicas"`

	// 支持：
	// "gpu-util": GPU利用率。范围{10, 100}      "cpu-util": CPU利用率。范围{10, 100}	      "memory-util": 内存利用率。范围{10, 100}      "service-qps": 单个实例QPS值。范围{1, 5000}
	// "concurrency-util":单个实例请求数量值。范围{1,100000}
	// 注意：此字段可能返回 null，表示取不到有效值。
	HpaMetrics []*Option `json:"HpaMetrics,omitnil,omitempty" name:"HpaMetrics"`

	// 扩容观察期，单位秒
	ScaleUpStabilizationWindowSeconds *int64 `json:"ScaleUpStabilizationWindowSeconds,omitnil,omitempty" name:"ScaleUpStabilizationWindowSeconds"`

	// 缩容观察期，单位秒
	ScaleDownStabilizationWindowSeconds *int64 `json:"ScaleDownStabilizationWindowSeconds,omitnil,omitempty" name:"ScaleDownStabilizationWindowSeconds"`
}

type HostPath struct {
	// 需要挂载的主机路径
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`
}

type HyperParameter struct {
	// 最大nnz数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxNNZ *string `json:"MaxNNZ,omitnil,omitempty" name:"MaxNNZ"`

	// slot数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlotNum *string `json:"SlotNum,omitnil,omitempty" name:"SlotNum"`

	// gpu cache 使用率
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuCachePercentage *string `json:"CpuCachePercentage,omitnil,omitempty" name:"CpuCachePercentage"`

	// cpu cache 使用率
	// 注意：此字段可能返回 null，表示取不到有效值。
	GpuCachePercentage *string `json:"GpuCachePercentage,omitnil,omitempty" name:"GpuCachePercentage"`

	// 是否开启分布式模式(true/false)
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableDistributed *string `json:"EnableDistributed,omitnil,omitempty" name:"EnableDistributed"`

	// TORCH_SCRIPT、MMDETECTION、DETECTRON2、HUGGINGFACE格式在进行优化时切分子图的最小算子数目，一般无需进行改动，默认为3
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinBlockSizePt *string `json:"MinBlockSizePt,omitnil,omitempty" name:"MinBlockSizePt"`

	// FROZEN_GRAPH、SAVED_MODEL格式在进行优化时切分子图的最小算子数目，一般无需进行改动，默认为10
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinBlockSizeTf *string `json:"MinBlockSizeTf,omitnil,omitempty" name:"MinBlockSizeTf"`

	// Stable Diffusion 模型优化参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PipelineArgs *string `json:"PipelineArgs,omitnil,omitempty" name:"PipelineArgs"`

	// Stable Diffusion 模型优化参数，控制Lora模型的影响效果
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoraScale *string `json:"LoraScale,omitnil,omitempty" name:"LoraScale"`
}

type ImageFIlter struct {
	// 过滤字段名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 过滤值
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// 是否反选
	Negative *bool `json:"Negative,omitnil,omitempty" name:"Negative"`
}

type ImageInfo struct {
	// 镜像类型：TCR为腾讯云TCR镜像; CCR为腾讯云TCR个人版镜像，PreSet为平台预置镜像，CUSTOM为第三方自定义镜像
	ImageType *string `json:"ImageType,omitnil,omitempty" name:"ImageType"`

	// 镜像地址
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// TCR镜像对应的地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegistryRegion *string `json:"RegistryRegion,omitnil,omitempty" name:"RegistryRegion"`

	// TCR镜像对应的实例id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 是否允许导出全部内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllowSaveAllContent *bool `json:"AllowSaveAllContent,omitnil,omitempty" name:"AllowSaveAllContent"`

	// 镜像名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageName *string `json:"ImageName,omitnil,omitempty" name:"ImageName"`

	// 是否支持数据构建
	// 注意：此字段可能返回 null，表示取不到有效值。
	SupportDataPipeline *bool `json:"SupportDataPipeline,omitnil,omitempty" name:"SupportDataPipeline"`
}

type ImageUrl struct {
	// 图片url
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type InferCodeInfo struct {
	// 推理代码所在的cos详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosPathInfo *CosPathInfo `json:"CosPathInfo,omitnil,omitempty" name:"CosPathInfo"`
}

type InferGatewayCallInfo struct {
	// 内网http调用地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcHttpAddr *string `json:"VpcHttpAddr,omitnil,omitempty" name:"VpcHttpAddr"`

	// 内网https调用地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcHttpsAddr *string `json:"VpcHttpsAddr,omitnil,omitempty" name:"VpcHttpsAddr"`

	// 内网grpc调用地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcGrpcTlsAddr *string `json:"VpcGrpcTlsAddr,omitnil,omitempty" name:"VpcGrpcTlsAddr"`

	// 可访问的vpcid
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 后端ip对应的子网
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`
}

type InferTemplate struct {
	// 模板ID
	InferTemplateId *string `json:"InferTemplateId,omitnil,omitempty" name:"InferTemplateId"`

	// 模板镜像
	InferTemplateImage *string `json:"InferTemplateImage,omitnil,omitempty" name:"InferTemplateImage"`
}

type InferTemplateGroup struct {
	// 算法框架
	// 注意：此字段可能返回 null，表示取不到有效值。
	Framework *string `json:"Framework,omitnil,omitempty" name:"Framework"`

	// 版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrameworkVersion *string `json:"FrameworkVersion,omitnil,omitempty" name:"FrameworkVersion"`

	// 支持的训练框架集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	Groups []*string `json:"Groups,omitnil,omitempty" name:"Groups"`

	// 镜像模板参数列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	InferTemplates []*InferTemplate `json:"InferTemplates,omitnil,omitempty" name:"InferTemplates"`
}

type IngressPrivateLinkInfo struct {
	// 用户VpcId
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 用户子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 内网http调用地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerHttpAddr []*string `json:"InnerHttpAddr,omitnil,omitempty" name:"InnerHttpAddr"`

	// 内网https调用地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerHttpsAddr []*string `json:"InnerHttpsAddr,omitnil,omitempty" name:"InnerHttpsAddr"`

	// 私有连接状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *string `json:"State,omitnil,omitempty" name:"State"`
}

type Instance struct {
	// 资源组节点id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 节点已用资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsedResource *ResourceInfo `json:"UsedResource,omitnil,omitempty" name:"UsedResource"`

	// 节点总资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalResource *ResourceInfo `json:"TotalResource,omitnil,omitempty" name:"TotalResource"`

	// 节点状态 
	// 注意：此字段为枚举值
	// 说明: 
	// DEPLOYING: 部署中
	// RUNNING: 运行中 
	// DEPLOY_FAILED: 部署失败
	// RELEASING 释放中 
	// RELEASED：已释放 
	// EXCEPTION：异常
	// DEBT_OR_EXPIRED: 欠费过期
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceStatus *string `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// 创建人
	SubUin *string `json:"SubUin,omitnil,omitempty" name:"SubUin"`

	// 创建时间: 
	// 注意：北京时间，比如: 2021-12-01 12:00:00
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 到期时间
	// 注意：北京时间，比如：2021-12-11 12:00:00
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 自动续费标识
	// 注意：此字段为枚举值
	// 说明：
	// NOTIFY_AND_MANUAL_RENEW：手动续费(取消自动续费)且到期通知
	// NOTIFY_AND_AUTO_RENEW：自动续费且到期通知
	// DISABLE_NOTIFY_AND_MANUAL_RENEW：手动续费(取消自动续费)且到期不通知
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoRenewFlag *string `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 计费项ID
	SpecId *string `json:"SpecId,omitnil,omitempty" name:"SpecId"`

	// 计费项别名
	SpecAlias *string `json:"SpecAlias,omitnil,omitempty" name:"SpecAlias"`

	// 计费项特性列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecFeatures []*string `json:"SpecFeatures,omitnil,omitempty" name:"SpecFeatures"`

	// 纳管cvmid
	CvmInstanceId *string `json:"CvmInstanceId,omitnil,omitempty" name:"CvmInstanceId"`

	// 部署失败错误码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrCode *string `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 部署失败错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMsg *string `json:"ErrMsg,omitnil,omitempty" name:"ErrMsg"`
}

type IntranetCallInfo struct {
	// 私有连接通道信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	IngressPrivateLinkInfo *IngressPrivateLinkInfo `json:"IngressPrivateLinkInfo,omitnil,omitempty" name:"IngressPrivateLinkInfo"`

	// 共享弹性网卡信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceEIPInfo []*ServiceEIPInfo `json:"ServiceEIPInfo,omitnil,omitempty" name:"ServiceEIPInfo"`

	// 默认内网调用信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultInnerCallInfos []*DefaultInnerCallInfo `json:"DefaultInnerCallInfos,omitnil,omitempty" name:"DefaultInnerCallInfos"`

	// 私有连接信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateLinkInfos []*PrivateLinkInfo `json:"PrivateLinkInfos,omitnil,omitempty" name:"PrivateLinkInfos"`

	// 基于新网关的私有连接信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateLinkInfosV2 []*PrivateLinkInfo `json:"PrivateLinkInfosV2,omitnil,omitempty" name:"PrivateLinkInfosV2"`
}

type LabelValue struct {
	// 标签名称
	LabelName *string `json:"LabelName,omitnil,omitempty" name:"LabelName"`

	// 标签的颜色
	LabelColor *string `json:"LabelColor,omitnil,omitempty" name:"LabelColor"`
}

type LocalDisk struct {
	// 节点ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 本地路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocalPath *string `json:"LocalPath,omitnil,omitempty" name:"LocalPath"`
}

type LogConfig struct {
	// 日志需要投递到cls的日志集
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// 日志需要投递到cls的主题
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

type LogIdentity struct {
	// 单条日志的ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 单条日志的内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 这条日志对应的Pod名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodName *string `json:"PodName,omitnil,omitempty" name:"PodName"`

	// 日志的时间戳（RFC3339格式的时间字符串）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *string `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`
}

type MaterialInfo struct {
	// 存储类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// Cos存储信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosPathInfo *CosPathInfo `json:"CosPathInfo,omitnil,omitempty" name:"CosPathInfo"`

	// 物料名，支持Code、Model
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaterialName *string `json:"MaterialName,omitnil,omitempty" name:"MaterialName"`

	// 物料类型，支持PreSet(预置)、 Custom(自定义)
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaterialType *string `json:"MaterialType,omitnil,omitempty" name:"MaterialType"`

	// 训练任务挂载路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	MountPath *string `json:"MountPath,omitnil,omitempty" name:"MountPath"`
}

type Message struct {
	// 角色名。支持三个角色：system、user、assistant，其中system仅开头可出现一次，也可忽略。
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 对话输入内容。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 多模态对话输入内容，Content与MultiModalContents两者只需要填写其中一个即可，当对话中包含多模态对话信息时，则填写本参数
	MultiModalContents []*MultiModalContent `json:"MultiModalContents,omitnil,omitempty" name:"MultiModalContents"`
}

type MetricData struct {
	// 训练任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 时间戳.unix timestamp,单位为秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 用户uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 本次上报数据所处的训练周期数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Epoch *int64 `json:"Epoch,omitnil,omitempty" name:"Epoch"`

	// 本次上报数据所处的训练迭代次数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Step *int64 `json:"Step,omitnil,omitempty" name:"Step"`

	// 训练停止所需的迭代总数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalSteps *int64 `json:"TotalSteps,omitnil,omitempty" name:"TotalSteps"`

	// 数据点。数组元素为不同指标的数据。数组长度不超过10。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Points []*DataPoint `json:"Points,omitnil,omitempty" name:"Points"`
}

type ModelAccelerateTask struct {
	// 模型加速任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelAccTaskId *string `json:"ModelAccTaskId,omitnil,omitempty" name:"ModelAccTaskId"`

	// 模型加速任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelAccTaskName *string `json:"ModelAccTaskName,omitnil,omitempty" name:"ModelAccTaskName"`

	// 模型ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 模型名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 模型版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelVersion *string `json:"ModelVersion,omitnil,omitempty" name:"ModelVersion"`

	// 模型来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelSource *string `json:"ModelSource,omitnil,omitempty" name:"ModelSource"`

	// 优化级别
	// 注意：此字段可能返回 null，表示取不到有效值。
	OptimizationLevel *string `json:"OptimizationLevel,omitnil,omitempty" name:"OptimizationLevel"`

	// 任务状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskStatus *string `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// input节点个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelInputNum *uint64 `json:"ModelInputNum,omitnil,omitempty" name:"ModelInputNum"`

	// input节点信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelInputInfos []*ModelInputInfo `json:"ModelInputInfos,omitnil,omitempty" name:"ModelInputInfos"`

	// GPU型号
	// 注意：此字段可能返回 null，表示取不到有效值。
	GPUType *string `json:"GPUType,omitnil,omitempty" name:"GPUType"`

	// 计费模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 加速比
	// 注意：此字段可能返回 null，表示取不到有效值。
	Speedup *string `json:"Speedup,omitnil,omitempty" name:"Speedup"`

	// 模型输入cos路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelInputPath *CosPathInfo `json:"ModelInputPath,omitnil,omitempty" name:"ModelInputPath"`

	// 模型输出cos路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelOutputPath *CosPathInfo `json:"ModelOutputPath,omitnil,omitempty" name:"ModelOutputPath"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 算法框架
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlgorithmFramework *string `json:"AlgorithmFramework,omitnil,omitempty" name:"AlgorithmFramework"`

	// 排队个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	WaitNumber *uint64 `json:"WaitNumber,omitnil,omitempty" name:"WaitNumber"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 任务进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskProgress *uint64 `json:"TaskProgress,omitnil,omitempty" name:"TaskProgress"`

	// 模型格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelFormat *string `json:"ModelFormat,omitnil,omitempty" name:"ModelFormat"`

	// 模型Tensor信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TensorInfos []*string `json:"TensorInfos,omitnil,omitempty" name:"TensorInfos"`

	// 模型专业参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	HyperParameter *HyperParameter `json:"HyperParameter,omitnil,omitempty" name:"HyperParameter"`

	// 加速引擎版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccEngineVersion *string `json:"AccEngineVersion,omitnil,omitempty" name:"AccEngineVersion"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 优化模型是否已保存到模型仓库
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsSaved *bool `json:"IsSaved,omitnil,omitempty" name:"IsSaved"`

	// SAVED_MODEL保存时配置的签名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelSignature *string `json:"ModelSignature,omitnil,omitempty" name:"ModelSignature"`

	// 是否是QAT模型
	// 注意：此字段可能返回 null，表示取不到有效值。
	QATModel *bool `json:"QATModel,omitnil,omitempty" name:"QATModel"`

	// 加速引擎对应的框架版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrameworkVersion *string `json:"FrameworkVersion,omitnil,omitempty" name:"FrameworkVersion"`

	// 模型版本ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelVersionId *string `json:"ModelVersionId,omitnil,omitempty" name:"ModelVersionId"`

	// 资源组id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// 资源组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupName *string `json:"ResourceGroupName,omitnil,omitempty" name:"ResourceGroupName"`
}

type ModelAccelerateVersion struct {
	// 模型id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 优化模型版本id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelVersionId *string `json:"ModelVersionId,omitnil,omitempty" name:"ModelVersionId"`

	// 优化任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelJobId *string `json:"ModelJobId,omitnil,omitempty" name:"ModelJobId"`

	// 优化任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelJobName *string `json:"ModelJobName,omitnil,omitempty" name:"ModelJobName"`

	// 优化后模型版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelVersion *string `json:"ModelVersion,omitnil,omitempty" name:"ModelVersion"`

	// 加速比
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpeedUp *string `json:"SpeedUp,omitnil,omitempty" name:"SpeedUp"`

	// 模型来源/任务名称/任务版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelSource *ModelSource `json:"ModelSource,omitnil,omitempty" name:"ModelSource"`

	// 模型cos路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosPathInfo *CosPathInfo `json:"CosPathInfo,omitnil,omitempty" name:"CosPathInfo"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 模型规范
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelFormat *string `json:"ModelFormat,omitnil,omitempty" name:"ModelFormat"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Progress *uint64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// GPU类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	GPUType *string `json:"GPUType,omitnil,omitempty" name:"GPUType"`

	// 模型cos路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelCosPath *CosPathInfo `json:"ModelCosPath,omitnil,omitempty" name:"ModelCosPath"`
}

type ModelInfo struct {
	// <p>模型版本id, DescribeTrainingModelVersion查询模型接口时的id<br>自动学习类型的模型填写自动学习的任务id</p>
	ModelVersionId *string `json:"ModelVersionId,omitnil,omitempty" name:"ModelVersionId"`

	// <p>模型id</p>
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// <p>模型名</p>
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// <p>模型版本</p>
	ModelVersion *string `json:"ModelVersion,omitnil,omitempty" name:"ModelVersion"`

	// <p>模型来源</p>
	ModelSource *string `json:"ModelSource,omitnil,omitempty" name:"ModelSource"`

	// <p>cos路径信息</p>
	CosPathInfo *CosPathInfo `json:"CosPathInfo,omitnil,omitempty" name:"CosPathInfo"`

	// <p>GooseFSx的配置，ModelSource为GooseFSx时有效</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	GooseFSx *GooseFSx `json:"GooseFSx,omitnil,omitempty" name:"GooseFSx"`

	// <p>模型对应的算法框架，预留</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlgorithmFramework *string `json:"AlgorithmFramework,omitnil,omitempty" name:"AlgorithmFramework"`

	// <p>默认为 NORMAL, 已加速模型: ACCELERATE, 自动学习模型 AUTO_ML</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelType *string `json:"ModelType,omitnil,omitempty" name:"ModelType"`

	// <p>模型格式</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelFormat *string `json:"ModelFormat,omitnil,omitempty" name:"ModelFormat"`

	// <p>是否为私有化大模型</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsPrivateModel *bool `json:"IsPrivateModel,omitnil,omitempty" name:"IsPrivateModel"`

	// <p>模型的类别 多模态MultiModal, 文本大模型 LLM</p>
	ModelCategory *string `json:"ModelCategory,omitnil,omitempty" name:"ModelCategory"`

	// <p>数据源的配置</p>
	PublicDataSource *PublicDataSourceFS `json:"PublicDataSource,omitnil,omitempty" name:"PublicDataSource"`
}

type ModelInputInfo struct {
	// input数据类型
	// FIXED：固定
	// RANGE：浮动
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelInputType *string `json:"ModelInputType,omitnil,omitempty" name:"ModelInputType"`

	// input数据尺寸
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelInputDimension []*string `json:"ModelInputDimension,omitnil,omitempty" name:"ModelInputDimension"`
}

type ModelSource struct {
	// 来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 来源任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobName *string `json:"JobName,omitnil,omitempty" name:"JobName"`

	// 来源任务版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobVersion *string `json:"JobVersion,omitnil,omitempty" name:"JobVersion"`

	// 来源任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 模型名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 算法框架
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlgorithmFramework *string `json:"AlgorithmFramework,omitnil,omitempty" name:"AlgorithmFramework"`

	// 训练偏好
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrainingPreference *string `json:"TrainingPreference,omitnil,omitempty" name:"TrainingPreference"`

	// 推理环境来源，SYSTEM/CUSTOM
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReasoningEnvironmentSource *string `json:"ReasoningEnvironmentSource,omitnil,omitempty" name:"ReasoningEnvironmentSource"`

	// 推理环境
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReasoningEnvironment *string `json:"ReasoningEnvironment,omitnil,omitempty" name:"ReasoningEnvironment"`

	// 推理环境id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReasoningEnvironmentId *string `json:"ReasoningEnvironmentId,omitnil,omitempty" name:"ReasoningEnvironmentId"`

	// 自定义推理环境
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReasoningImageInfo *ImageInfo `json:"ReasoningImageInfo,omitnil,omitempty" name:"ReasoningImageInfo"`
}

// Predefined struct for user
type ModifyModelServiceAuthTokenRequestParams struct {
	// 服务组 id
	ServiceGroupId *string `json:"ServiceGroupId,omitnil,omitempty" name:"ServiceGroupId"`

	// 是否需要重置，如果为 true，重置 token 值
	NeedReset *bool `json:"NeedReset,omitnil,omitempty" name:"NeedReset"`

	// AuthToken 数据
	AuthToken *AuthToken `json:"AuthToken,omitnil,omitempty" name:"AuthToken"`
}

type ModifyModelServiceAuthTokenRequest struct {
	*tchttp.BaseRequest
	
	// 服务组 id
	ServiceGroupId *string `json:"ServiceGroupId,omitnil,omitempty" name:"ServiceGroupId"`

	// 是否需要重置，如果为 true，重置 token 值
	NeedReset *bool `json:"NeedReset,omitnil,omitempty" name:"NeedReset"`

	// AuthToken 数据
	AuthToken *AuthToken `json:"AuthToken,omitnil,omitempty" name:"AuthToken"`
}

func (r *ModifyModelServiceAuthTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModelServiceAuthTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceGroupId")
	delete(f, "NeedReset")
	delete(f, "AuthToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyModelServiceAuthTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModelServiceAuthTokenResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyModelServiceAuthTokenResponse struct {
	*tchttp.BaseResponse
	Response *ModifyModelServiceAuthTokenResponseParams `json:"Response"`
}

func (r *ModifyModelServiceAuthTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModelServiceAuthTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModelServiceAuthorizationRequestParams struct {
	// 服务组Id
	ServiceGroupId *string `json:"ServiceGroupId,omitnil,omitempty" name:"ServiceGroupId"`

	// 是否开启鉴权,true表示开启
	AuthorizationEnable *bool `json:"AuthorizationEnable,omitnil,omitempty" name:"AuthorizationEnable"`
}

type ModifyModelServiceAuthorizationRequest struct {
	*tchttp.BaseRequest
	
	// 服务组Id
	ServiceGroupId *string `json:"ServiceGroupId,omitnil,omitempty" name:"ServiceGroupId"`

	// 是否开启鉴权,true表示开启
	AuthorizationEnable *bool `json:"AuthorizationEnable,omitnil,omitempty" name:"AuthorizationEnable"`
}

func (r *ModifyModelServiceAuthorizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModelServiceAuthorizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceGroupId")
	delete(f, "AuthorizationEnable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyModelServiceAuthorizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModelServiceAuthorizationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyModelServiceAuthorizationResponse struct {
	*tchttp.BaseResponse
	Response *ModifyModelServiceAuthorizationResponseParams `json:"Response"`
}

func (r *ModifyModelServiceAuthorizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModelServiceAuthorizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModelServiceRequestParams struct {
	// 服务id
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 模型信息，需要挂载模型时填写
	ModelInfo *ModelInfo `json:"ModelInfo,omitnil,omitempty" name:"ModelInfo"`

	// 镜像信息，配置服务运行所需的镜像地址等信息
	ImageInfo *ImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`

	// 环境变量，可选参数，用于配置容器中的环境变量
	Env []*EnvVar `json:"Env,omitnil,omitempty" name:"Env"`

	// 资源描述，指定预付费模式下的cpu,mem,gpu等信息，后付费无需填写
	Resources *ResourceInfo `json:"Resources,omitnil,omitempty" name:"Resources"`

	// 使用DescribeBillingSpecs接口返回的规格列表中的值，或者参考实例列表:
	// TI.S.MEDIUM.POST	2C4G
	// TI.S.LARGE.POST	4C8G
	// TI.S.2XLARGE16.POST	8C16G
	// TI.S.2XLARGE32.POST	8C32G
	// TI.S.4XLARGE32.POST	16C32G
	// TI.S.4XLARGE64.POST	16C64G
	// TI.S.6XLARGE48.POST	24C48G
	// TI.S.6XLARGE96.POST	24C96G
	// TI.S.8XLARGE64.POST	32C64G
	// TI.S.8XLARGE128.POST 32C128G
	// TI.GN7.LARGE20.POST	4C20G T4*1/4
	// TI.GN7.2XLARGE40.POST	10C40G T4*1/2
	// TI.GN7.2XLARGE32.POST	8C32G T4*1
	// TI.GN7.5XLARGE80.POST	20C80G T4*1
	// TI.GN7.8XLARGE128.POST	32C128G T4*1
	// TI.GN7.10XLARGE160.POST	40C160G T4*2
	// TI.GN7.20XLARGE320.POST	80C320G T4*4
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 扩缩容类型 支持：自动 - "AUTO", 手动 - "MANUAL"
	ScaleMode *string `json:"ScaleMode,omitnil,omitempty" name:"ScaleMode"`

	// 实例数量, 不同计费模式和调节模式下对应关系如下
	// PREPAID 和 POSTPAID_BY_HOUR:
	// 手动调节模式下对应 实例数量
	// 自动调节模式下对应 基于时间的默认策略的实例数量
	// HYBRID_PAID:
	// 后付费实例手动调节模式下对应 实例数量
	// 后付费实例自动调节模式下对应 时间策略的默认策略的实例数量
	Replicas *int64 `json:"Replicas,omitnil,omitempty" name:"Replicas"`

	// 自动伸缩信息
	HorizontalPodAutoscaler *HorizontalPodAutoscaler `json:"HorizontalPodAutoscaler,omitnil,omitempty" name:"HorizontalPodAutoscaler"`

	// 是否开启日志投递，开启后需填写配置投递到指定cls
	LogEnable *bool `json:"LogEnable,omitnil,omitempty" name:"LogEnable"`

	// 日志配置，需要投递服务日志到指定cls时填写
	LogConfig *LogConfig `json:"LogConfig,omitnil,omitempty" name:"LogConfig"`

	// 特殊更新行为： "STOP": 停止, "RESUME": 重启, "SCALE": 扩缩容, 存在这些特殊更新行为时，会忽略其他更新字段
	ServiceAction *string `json:"ServiceAction,omitnil,omitempty" name:"ServiceAction"`

	// 服务的描述
	ServiceDescription *string `json:"ServiceDescription,omitnil,omitempty" name:"ServiceDescription"`

	// 自动伸缩策略
	ScaleStrategy *string `json:"ScaleStrategy,omitnil,omitempty" name:"ScaleStrategy"`

	// 自动伸缩策略配置 HPA : 通过HPA进行弹性伸缩 CRON 通过定时任务进行伸缩
	CronScaleJobs []*CronScaleJob `json:"CronScaleJobs,omitnil,omitempty" name:"CronScaleJobs"`

	// 计费模式[HYBRID_PAID]时生效, 用于标识混合计费模式下的预付费实例数, 若不填则默认为1
	HybridBillingPrepaidReplicas *int64 `json:"HybridBillingPrepaidReplicas,omitnil,omitempty" name:"HybridBillingPrepaidReplicas"`

	// 是否开启模型的热更新。默认不开启
	ModelHotUpdateEnable *bool `json:"ModelHotUpdateEnable,omitnil,omitempty" name:"ModelHotUpdateEnable"`

	// 定时停止配置
	ScheduledAction *ScheduledAction `json:"ScheduledAction,omitnil,omitempty" name:"ScheduledAction"`

	// 服务限速限流相关配置
	ServiceLimit *ServiceLimit `json:"ServiceLimit,omitnil,omitempty" name:"ServiceLimit"`

	// 挂载配置，目前只支持CFS
	VolumeMount *VolumeMount `json:"VolumeMount,omitnil,omitempty" name:"VolumeMount"`

	// 是否开启模型的加速, 仅对StableDiffusion(动态加速)格式的模型有效。默认不开启
	ModelTurboEnable *bool `json:"ModelTurboEnable,omitnil,omitempty" name:"ModelTurboEnable"`

	// 服务的启动命令，如遇特殊字符导致配置失败，可使用CommandBase64参数
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// 是否开启TIONE内网访问外部，此功能仅支持后付费机型与从TIONE平台购买的预付费机型；使用从CVM选择资源组时此配置不生效。
	ServiceEIP *ServiceEIP `json:"ServiceEIP,omitnil,omitempty" name:"ServiceEIP"`

	// 服务的启动命令，以base64格式进行输入，与Command同时配置时，仅当前参数生效
	CommandBase64 *string `json:"CommandBase64,omitnil,omitempty" name:"CommandBase64"`

	// 服务端口，仅在非内置镜像时生效，默认8501。不支持输入8501-8510,6006,9092
	ServicePort *int64 `json:"ServicePort,omitnil,omitempty" name:"ServicePort"`

	// 单副本下的实例数，仅在部署类型为DIST时生效，默认1
	InstancePerReplicas *int64 `json:"InstancePerReplicas,omitnil,omitempty" name:"InstancePerReplicas"`

	// 服务的优雅退出时限。单位为秒，默认值为30，最小为1
	TerminationGracePeriodSeconds *int64 `json:"TerminationGracePeriodSeconds,omitnil,omitempty" name:"TerminationGracePeriodSeconds"`

	// 服务实例停止前执行的命令，执行完毕或执行时间超过优雅退出时限后实例结束
	PreStopCommand []*string `json:"PreStopCommand,omitnil,omitempty" name:"PreStopCommand"`

	// 是否启动grpc端口
	GrpcEnable *bool `json:"GrpcEnable,omitnil,omitempty" name:"GrpcEnable"`

	// 健康探针
	HealthProbe *HealthProbe `json:"HealthProbe,omitnil,omitempty" name:"HealthProbe"`

	// 滚动更新策略
	RollingUpdate *RollingUpdate `json:"RollingUpdate,omitnil,omitempty" name:"RollingUpdate"`

	// sidecar配置
	Sidecar *SidecarSpec `json:"Sidecar,omitnil,omitempty" name:"Sidecar"`

	// 资源组 id
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// 数据盘批量挂载配置，当前仅支持CFS，仅针对“模型来源-腾讯云存储、模型来源-腾讯云容器镜像、模型来源-资源组、模型来源-数据源”。
	VolumeMounts []*VolumeMount `json:"VolumeMounts,omitnil,omitempty" name:"VolumeMounts"`

	// 调度策略 [binpack] 优先占满整机，尽量避免碎卡（默认值）[spread] 优先分散在各个节点，确保服务高可用
	SchedulingStrategy *string `json:"SchedulingStrategy,omitnil,omitempty" name:"SchedulingStrategy"`

	// 目标工作空间，不为0则进行迁移，源服务只允许在默认空间
	TargetProjectId *int64 `json:"TargetProjectId,omitnil,omitempty" name:"TargetProjectId"`
}

type ModifyModelServiceRequest struct {
	*tchttp.BaseRequest
	
	// 服务id
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 模型信息，需要挂载模型时填写
	ModelInfo *ModelInfo `json:"ModelInfo,omitnil,omitempty" name:"ModelInfo"`

	// 镜像信息，配置服务运行所需的镜像地址等信息
	ImageInfo *ImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`

	// 环境变量，可选参数，用于配置容器中的环境变量
	Env []*EnvVar `json:"Env,omitnil,omitempty" name:"Env"`

	// 资源描述，指定预付费模式下的cpu,mem,gpu等信息，后付费无需填写
	Resources *ResourceInfo `json:"Resources,omitnil,omitempty" name:"Resources"`

	// 使用DescribeBillingSpecs接口返回的规格列表中的值，或者参考实例列表:
	// TI.S.MEDIUM.POST	2C4G
	// TI.S.LARGE.POST	4C8G
	// TI.S.2XLARGE16.POST	8C16G
	// TI.S.2XLARGE32.POST	8C32G
	// TI.S.4XLARGE32.POST	16C32G
	// TI.S.4XLARGE64.POST	16C64G
	// TI.S.6XLARGE48.POST	24C48G
	// TI.S.6XLARGE96.POST	24C96G
	// TI.S.8XLARGE64.POST	32C64G
	// TI.S.8XLARGE128.POST 32C128G
	// TI.GN7.LARGE20.POST	4C20G T4*1/4
	// TI.GN7.2XLARGE40.POST	10C40G T4*1/2
	// TI.GN7.2XLARGE32.POST	8C32G T4*1
	// TI.GN7.5XLARGE80.POST	20C80G T4*1
	// TI.GN7.8XLARGE128.POST	32C128G T4*1
	// TI.GN7.10XLARGE160.POST	40C160G T4*2
	// TI.GN7.20XLARGE320.POST	80C320G T4*4
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 扩缩容类型 支持：自动 - "AUTO", 手动 - "MANUAL"
	ScaleMode *string `json:"ScaleMode,omitnil,omitempty" name:"ScaleMode"`

	// 实例数量, 不同计费模式和调节模式下对应关系如下
	// PREPAID 和 POSTPAID_BY_HOUR:
	// 手动调节模式下对应 实例数量
	// 自动调节模式下对应 基于时间的默认策略的实例数量
	// HYBRID_PAID:
	// 后付费实例手动调节模式下对应 实例数量
	// 后付费实例自动调节模式下对应 时间策略的默认策略的实例数量
	Replicas *int64 `json:"Replicas,omitnil,omitempty" name:"Replicas"`

	// 自动伸缩信息
	HorizontalPodAutoscaler *HorizontalPodAutoscaler `json:"HorizontalPodAutoscaler,omitnil,omitempty" name:"HorizontalPodAutoscaler"`

	// 是否开启日志投递，开启后需填写配置投递到指定cls
	LogEnable *bool `json:"LogEnable,omitnil,omitempty" name:"LogEnable"`

	// 日志配置，需要投递服务日志到指定cls时填写
	LogConfig *LogConfig `json:"LogConfig,omitnil,omitempty" name:"LogConfig"`

	// 特殊更新行为： "STOP": 停止, "RESUME": 重启, "SCALE": 扩缩容, 存在这些特殊更新行为时，会忽略其他更新字段
	ServiceAction *string `json:"ServiceAction,omitnil,omitempty" name:"ServiceAction"`

	// 服务的描述
	ServiceDescription *string `json:"ServiceDescription,omitnil,omitempty" name:"ServiceDescription"`

	// 自动伸缩策略
	ScaleStrategy *string `json:"ScaleStrategy,omitnil,omitempty" name:"ScaleStrategy"`

	// 自动伸缩策略配置 HPA : 通过HPA进行弹性伸缩 CRON 通过定时任务进行伸缩
	CronScaleJobs []*CronScaleJob `json:"CronScaleJobs,omitnil,omitempty" name:"CronScaleJobs"`

	// 计费模式[HYBRID_PAID]时生效, 用于标识混合计费模式下的预付费实例数, 若不填则默认为1
	HybridBillingPrepaidReplicas *int64 `json:"HybridBillingPrepaidReplicas,omitnil,omitempty" name:"HybridBillingPrepaidReplicas"`

	// 是否开启模型的热更新。默认不开启
	ModelHotUpdateEnable *bool `json:"ModelHotUpdateEnable,omitnil,omitempty" name:"ModelHotUpdateEnable"`

	// 定时停止配置
	ScheduledAction *ScheduledAction `json:"ScheduledAction,omitnil,omitempty" name:"ScheduledAction"`

	// 服务限速限流相关配置
	ServiceLimit *ServiceLimit `json:"ServiceLimit,omitnil,omitempty" name:"ServiceLimit"`

	// 挂载配置，目前只支持CFS
	VolumeMount *VolumeMount `json:"VolumeMount,omitnil,omitempty" name:"VolumeMount"`

	// 是否开启模型的加速, 仅对StableDiffusion(动态加速)格式的模型有效。默认不开启
	ModelTurboEnable *bool `json:"ModelTurboEnable,omitnil,omitempty" name:"ModelTurboEnable"`

	// 服务的启动命令，如遇特殊字符导致配置失败，可使用CommandBase64参数
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// 是否开启TIONE内网访问外部，此功能仅支持后付费机型与从TIONE平台购买的预付费机型；使用从CVM选择资源组时此配置不生效。
	ServiceEIP *ServiceEIP `json:"ServiceEIP,omitnil,omitempty" name:"ServiceEIP"`

	// 服务的启动命令，以base64格式进行输入，与Command同时配置时，仅当前参数生效
	CommandBase64 *string `json:"CommandBase64,omitnil,omitempty" name:"CommandBase64"`

	// 服务端口，仅在非内置镜像时生效，默认8501。不支持输入8501-8510,6006,9092
	ServicePort *int64 `json:"ServicePort,omitnil,omitempty" name:"ServicePort"`

	// 单副本下的实例数，仅在部署类型为DIST时生效，默认1
	InstancePerReplicas *int64 `json:"InstancePerReplicas,omitnil,omitempty" name:"InstancePerReplicas"`

	// 服务的优雅退出时限。单位为秒，默认值为30，最小为1
	TerminationGracePeriodSeconds *int64 `json:"TerminationGracePeriodSeconds,omitnil,omitempty" name:"TerminationGracePeriodSeconds"`

	// 服务实例停止前执行的命令，执行完毕或执行时间超过优雅退出时限后实例结束
	PreStopCommand []*string `json:"PreStopCommand,omitnil,omitempty" name:"PreStopCommand"`

	// 是否启动grpc端口
	GrpcEnable *bool `json:"GrpcEnable,omitnil,omitempty" name:"GrpcEnable"`

	// 健康探针
	HealthProbe *HealthProbe `json:"HealthProbe,omitnil,omitempty" name:"HealthProbe"`

	// 滚动更新策略
	RollingUpdate *RollingUpdate `json:"RollingUpdate,omitnil,omitempty" name:"RollingUpdate"`

	// sidecar配置
	Sidecar *SidecarSpec `json:"Sidecar,omitnil,omitempty" name:"Sidecar"`

	// 资源组 id
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// 数据盘批量挂载配置，当前仅支持CFS，仅针对“模型来源-腾讯云存储、模型来源-腾讯云容器镜像、模型来源-资源组、模型来源-数据源”。
	VolumeMounts []*VolumeMount `json:"VolumeMounts,omitnil,omitempty" name:"VolumeMounts"`

	// 调度策略 [binpack] 优先占满整机，尽量避免碎卡（默认值）[spread] 优先分散在各个节点，确保服务高可用
	SchedulingStrategy *string `json:"SchedulingStrategy,omitnil,omitempty" name:"SchedulingStrategy"`

	// 目标工作空间，不为0则进行迁移，源服务只允许在默认空间
	TargetProjectId *int64 `json:"TargetProjectId,omitnil,omitempty" name:"TargetProjectId"`
}

func (r *ModifyModelServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModelServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "ModelInfo")
	delete(f, "ImageInfo")
	delete(f, "Env")
	delete(f, "Resources")
	delete(f, "InstanceType")
	delete(f, "ScaleMode")
	delete(f, "Replicas")
	delete(f, "HorizontalPodAutoscaler")
	delete(f, "LogEnable")
	delete(f, "LogConfig")
	delete(f, "ServiceAction")
	delete(f, "ServiceDescription")
	delete(f, "ScaleStrategy")
	delete(f, "CronScaleJobs")
	delete(f, "HybridBillingPrepaidReplicas")
	delete(f, "ModelHotUpdateEnable")
	delete(f, "ScheduledAction")
	delete(f, "ServiceLimit")
	delete(f, "VolumeMount")
	delete(f, "ModelTurboEnable")
	delete(f, "Command")
	delete(f, "ServiceEIP")
	delete(f, "CommandBase64")
	delete(f, "ServicePort")
	delete(f, "InstancePerReplicas")
	delete(f, "TerminationGracePeriodSeconds")
	delete(f, "PreStopCommand")
	delete(f, "GrpcEnable")
	delete(f, "HealthProbe")
	delete(f, "RollingUpdate")
	delete(f, "Sidecar")
	delete(f, "ResourceGroupId")
	delete(f, "VolumeMounts")
	delete(f, "SchedulingStrategy")
	delete(f, "TargetProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyModelServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModelServiceResponseParams struct {
	// 生成的模型服务
	// 注意：此字段可能返回 null，表示取不到有效值。
	Service *Service `json:"Service,omitnil,omitempty" name:"Service"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyModelServiceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyModelServiceResponseParams `json:"Response"`
}

func (r *ModifyModelServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModelServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNotebookRequestParams struct {
	// <p>notebook id</p>
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>名称。不超过60个字符，仅支持中英文、数字、下划线&quot;_&quot;、短横&quot;-&quot;，只能以中英文、数字开头</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>（不允许修改）计算资源付费模式 ，可选值为：<br>PREPAID：预付费，即包年包月<br>POSTPAID_BY_HOUR：按小时后付费</p>
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// <p>计算资源配置</p>
	ResourceConf *ResourceConf `json:"ResourceConf,omitnil,omitempty" name:"ResourceConf"`

	// <p>是否自动停止</p>
	AutoStopping *bool `json:"AutoStopping,omitnil,omitempty" name:"AutoStopping"`

	// <p>是否访问公网</p>
	DirectInternetAccess *bool `json:"DirectInternetAccess,omitnil,omitempty" name:"DirectInternetAccess"`

	// <p>是否ROOT权限</p>
	RootAccess *bool `json:"RootAccess,omitnil,omitempty" name:"RootAccess"`

	// <p>是否上报日志</p>
	LogEnable *bool `json:"LogEnable,omitnil,omitempty" name:"LogEnable"`

	// <p>资源组ID(for预付费)</p>
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// <p>（不允许修改）Vpc-Id</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>（不允许修改）子网Id</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>存储卷大小，单位GB</p>
	VolumeSizeInGB *uint64 `json:"VolumeSizeInGB,omitnil,omitempty" name:"VolumeSizeInGB"`

	// <p>（不允许修改）存储的类型。取值包含：<br>    FREE:    预付费的免费存储<br>    CLOUD_PREMIUM： 高性能云硬盘<br>    CLOUD_SSD： SSD云硬盘<br>    CFS:     CFS存储，包含NFS和turbo</p>
	VolumeSourceType *string `json:"VolumeSourceType,omitnil,omitempty" name:"VolumeSourceType"`

	// <p>（不允许修改）CFS存储的配置</p>
	VolumeSourceCFS *CFSConfig `json:"VolumeSourceCFS,omitnil,omitempty" name:"VolumeSourceCFS"`

	// <p>日志配置</p>
	LogConfig *LogConfig `json:"LogConfig,omitnil,omitempty" name:"LogConfig"`

	// <p>生命周期脚本的ID</p>
	LifecycleScriptId *string `json:"LifecycleScriptId,omitnil,omitempty" name:"LifecycleScriptId"`

	// <p>默认GIT存储库的ID</p>
	DefaultCodeRepoId *string `json:"DefaultCodeRepoId,omitnil,omitempty" name:"DefaultCodeRepoId"`

	// <p>其他GIT存储库的ID，最多3个</p>
	AdditionalCodeRepoIds []*string `json:"AdditionalCodeRepoIds,omitnil,omitempty" name:"AdditionalCodeRepoIds"`

	// <p>自动停止时间，单位小时</p>
	AutomaticStopTime *int64 `json:"AutomaticStopTime,omitnil,omitempty" name:"AutomaticStopTime"`

	// <p>标签配置</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>数据配置，只支持WEDATA_HDFS</p>
	DataConfigs []*DataConfig `json:"DataConfigs,omitnil,omitempty" name:"DataConfigs"`

	// <p>镜像信息</p>
	ImageInfo *ImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`

	// <p>镜像类型，包括SYSTEM、TCR、CCR</p>
	ImageType *string `json:"ImageType,omitnil,omitempty" name:"ImageType"`

	// <p>SSH配置</p>
	SSHConfig *SSHConfig `json:"SSHConfig,omitnil,omitempty" name:"SSHConfig"`

	// <p>自定义环境变量</p>
	Envs []*EnvVar `json:"Envs,omitnil,omitempty" name:"Envs"`

	// <p>描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ModifyNotebookRequest struct {
	*tchttp.BaseRequest
	
	// <p>notebook id</p>
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>名称。不超过60个字符，仅支持中英文、数字、下划线&quot;_&quot;、短横&quot;-&quot;，只能以中英文、数字开头</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>（不允许修改）计算资源付费模式 ，可选值为：<br>PREPAID：预付费，即包年包月<br>POSTPAID_BY_HOUR：按小时后付费</p>
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// <p>计算资源配置</p>
	ResourceConf *ResourceConf `json:"ResourceConf,omitnil,omitempty" name:"ResourceConf"`

	// <p>是否自动停止</p>
	AutoStopping *bool `json:"AutoStopping,omitnil,omitempty" name:"AutoStopping"`

	// <p>是否访问公网</p>
	DirectInternetAccess *bool `json:"DirectInternetAccess,omitnil,omitempty" name:"DirectInternetAccess"`

	// <p>是否ROOT权限</p>
	RootAccess *bool `json:"RootAccess,omitnil,omitempty" name:"RootAccess"`

	// <p>是否上报日志</p>
	LogEnable *bool `json:"LogEnable,omitnil,omitempty" name:"LogEnable"`

	// <p>资源组ID(for预付费)</p>
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// <p>（不允许修改）Vpc-Id</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>（不允许修改）子网Id</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>存储卷大小，单位GB</p>
	VolumeSizeInGB *uint64 `json:"VolumeSizeInGB,omitnil,omitempty" name:"VolumeSizeInGB"`

	// <p>（不允许修改）存储的类型。取值包含：<br>    FREE:    预付费的免费存储<br>    CLOUD_PREMIUM： 高性能云硬盘<br>    CLOUD_SSD： SSD云硬盘<br>    CFS:     CFS存储，包含NFS和turbo</p>
	VolumeSourceType *string `json:"VolumeSourceType,omitnil,omitempty" name:"VolumeSourceType"`

	// <p>（不允许修改）CFS存储的配置</p>
	VolumeSourceCFS *CFSConfig `json:"VolumeSourceCFS,omitnil,omitempty" name:"VolumeSourceCFS"`

	// <p>日志配置</p>
	LogConfig *LogConfig `json:"LogConfig,omitnil,omitempty" name:"LogConfig"`

	// <p>生命周期脚本的ID</p>
	LifecycleScriptId *string `json:"LifecycleScriptId,omitnil,omitempty" name:"LifecycleScriptId"`

	// <p>默认GIT存储库的ID</p>
	DefaultCodeRepoId *string `json:"DefaultCodeRepoId,omitnil,omitempty" name:"DefaultCodeRepoId"`

	// <p>其他GIT存储库的ID，最多3个</p>
	AdditionalCodeRepoIds []*string `json:"AdditionalCodeRepoIds,omitnil,omitempty" name:"AdditionalCodeRepoIds"`

	// <p>自动停止时间，单位小时</p>
	AutomaticStopTime *int64 `json:"AutomaticStopTime,omitnil,omitempty" name:"AutomaticStopTime"`

	// <p>标签配置</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>数据配置，只支持WEDATA_HDFS</p>
	DataConfigs []*DataConfig `json:"DataConfigs,omitnil,omitempty" name:"DataConfigs"`

	// <p>镜像信息</p>
	ImageInfo *ImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`

	// <p>镜像类型，包括SYSTEM、TCR、CCR</p>
	ImageType *string `json:"ImageType,omitnil,omitempty" name:"ImageType"`

	// <p>SSH配置</p>
	SSHConfig *SSHConfig `json:"SSHConfig,omitnil,omitempty" name:"SSHConfig"`

	// <p>自定义环境变量</p>
	Envs []*EnvVar `json:"Envs,omitnil,omitempty" name:"Envs"`

	// <p>描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *ModifyNotebookRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNotebookRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Name")
	delete(f, "ChargeType")
	delete(f, "ResourceConf")
	delete(f, "AutoStopping")
	delete(f, "DirectInternetAccess")
	delete(f, "RootAccess")
	delete(f, "LogEnable")
	delete(f, "ResourceGroupId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "VolumeSizeInGB")
	delete(f, "VolumeSourceType")
	delete(f, "VolumeSourceCFS")
	delete(f, "LogConfig")
	delete(f, "LifecycleScriptId")
	delete(f, "DefaultCodeRepoId")
	delete(f, "AdditionalCodeRepoIds")
	delete(f, "AutomaticStopTime")
	delete(f, "Tags")
	delete(f, "DataConfigs")
	delete(f, "ImageInfo")
	delete(f, "ImageType")
	delete(f, "SSHConfig")
	delete(f, "Envs")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNotebookRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNotebookResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyNotebookResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNotebookResponseParams `json:"Response"`
}

func (r *ModifyNotebookResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNotebookResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNotebookTagsRequestParams struct {
	// Notebook Id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Notebook修改标签集合
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type ModifyNotebookTagsRequest struct {
	*tchttp.BaseRequest
	
	// Notebook Id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Notebook修改标签集合
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *ModifyNotebookTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNotebookTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNotebookTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNotebookTagsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyNotebookTagsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNotebookTagsResponseParams `json:"Response"`
}

func (r *ModifyNotebookTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNotebookTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyServiceGroupWeightsRequestParams struct {
	// 服务组id
	ServiceGroupId *string `json:"ServiceGroupId,omitnil,omitempty" name:"ServiceGroupId"`

	// 权重设置
	Weights []*WeightEntry `json:"Weights,omitnil,omitempty" name:"Weights"`
}

type ModifyServiceGroupWeightsRequest struct {
	*tchttp.BaseRequest
	
	// 服务组id
	ServiceGroupId *string `json:"ServiceGroupId,omitnil,omitempty" name:"ServiceGroupId"`

	// 权重设置
	Weights []*WeightEntry `json:"Weights,omitnil,omitempty" name:"Weights"`
}

func (r *ModifyServiceGroupWeightsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyServiceGroupWeightsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceGroupId")
	delete(f, "Weights")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyServiceGroupWeightsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyServiceGroupWeightsResponseParams struct {
	// 更新权重后的服务组信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceGroup *ServiceGroup `json:"ServiceGroup,omitnil,omitempty" name:"ServiceGroup"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyServiceGroupWeightsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyServiceGroupWeightsResponseParams `json:"Response"`
}

func (r *ModifyServiceGroupWeightsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyServiceGroupWeightsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MountConfigureInfo struct {
	// 数据源的相对路径，支持<@subaccount>这样的占位符
	WorkDir *string `json:"WorkDir,omitnil,omitempty" name:"WorkDir"`
}

type MountInstanceInfo struct {
	// <p>类型英文名</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>存储实例ID</p>
	StorageId *string `json:"StorageId,omitnil,omitempty" name:"StorageId"`

	// <p>存储实例名称</p>
	StorageName *string `json:"StorageName,omitnil,omitempty" name:"StorageName"`

	// <p>状态</p><p>枚举值：</p><ul><li>0： 可挂载（正常）</li><li>1： 不可挂载（挂载限制）</li><li>2： 不可挂载（存储配置关闭）</li></ul>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>额外配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraConf *StorageExtraConf `json:"ExtraConf,omitnil,omitempty" name:"ExtraConf"`
}

type MountLimitInfo struct {
	// <p>数据源类型英文名</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>数据源所属存储实例ID</p>
	StorageId *string `json:"StorageId,omitnil,omitempty" name:"StorageId"`

	// <p>数据源所属存储实例名称</p>
	StorageName *string `json:"StorageName,omitnil,omitempty" name:"StorageName"`

	// <p>限制开关是否开启，只有开启时才有限制</p>
	LimitMount *bool `json:"LimitMount,omitnil,omitempty" name:"LimitMount"`

	// <p>创建者uin</p>
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// <p>创建者名称</p>
	CreatorName *string `json:"CreatorName,omitnil,omitempty" name:"CreatorName"`

	// <p>创建时间, 格式为yyyy-mm-ddThh:mm:ssZ</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>更新时间, 格式为yyyy-mm-ddThh:mm:ssZ</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>额外配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraConf *StorageExtraConf `json:"ExtraConf,omitnil,omitempty" name:"ExtraConf"`
}

type MultiModalContent struct {
	// 对话类型，text表示文本对话内容，image_url表示图片对话内容
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 文本对话内容，当Type为text时需要填写该值
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 图片对话内容，当Type为image_url时需要填写该值
	ImageUrl *ImageUrl `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`
}

type NotebookDetail struct {
	// <p>notebook  ID</p>
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>notebook 名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>生命周期脚本</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	LifecycleScriptId *string `json:"LifecycleScriptId,omitnil,omitempty" name:"LifecycleScriptId"`

	// <p>Pod-Name</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodName *string `json:"PodName,omitnil,omitempty" name:"PodName"`

	// <p>Update-Time</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>是否访问公网</p>
	DirectInternetAccess *bool `json:"DirectInternetAccess,omitnil,omitempty" name:"DirectInternetAccess"`

	// <p>预付费专用资源组</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// <p>标签配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>是否自动停止</p>
	AutoStopping *bool `json:"AutoStopping,omitnil,omitempty" name:"AutoStopping"`

	// <p>其他GIT存储库，最多3个，单个<br>长度不超过512字符</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdditionalCodeRepoIds []*string `json:"AdditionalCodeRepoIds,omitnil,omitempty" name:"AdditionalCodeRepoIds"`

	// <p>自动停止时间，单位小时</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutomaticStopTime *int64 `json:"AutomaticStopTime,omitnil,omitempty" name:"AutomaticStopTime"`

	// <p>资源配置</p>
	ResourceConf *ResourceConf `json:"ResourceConf,omitnil,omitempty" name:"ResourceConf"`

	// <p>默认GIT存储库，长度不超过512字符</p>
	DefaultCodeRepoId *string `json:"DefaultCodeRepoId,omitnil,omitempty" name:"DefaultCodeRepoId"`

	// <p>训练输出</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>是否上报日志</p>
	LogEnable *bool `json:"LogEnable,omitnil,omitempty" name:"LogEnable"`

	// <p>日志配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogConfig *LogConfig `json:"LogConfig,omitnil,omitempty" name:"LogConfig"`

	// <p>VPC ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>子网ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>任务状态</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>运行时长</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuntimeInSeconds *uint64 `json:"RuntimeInSeconds,omitnil,omitempty" name:"RuntimeInSeconds"`

	// <p>创建时间</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>训练开始时间</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>计费状态，eg：BILLING计费中，ARREARS_STOP欠费停止，NOT_BILLING不在计费中</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeStatus *string `json:"ChargeStatus,omitnil,omitempty" name:"ChargeStatus"`

	// <p>是否ROOT权限</p>
	RootAccess *bool `json:"RootAccess,omitnil,omitempty" name:"RootAccess"`

	// <p>计贺金额信息，eg:2.00元/小时</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillingInfos []*string `json:"BillingInfos,omitnil,omitempty" name:"BillingInfos"`

	// <p>存储卷大小 （单位时GB，最小10GB，必须是10G的倍数）</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	VolumeSizeInGB *uint64 `json:"VolumeSizeInGB,omitnil,omitempty" name:"VolumeSizeInGB"`

	// <p>失败原因</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailureReason *string `json:"FailureReason,omitnil,omitempty" name:"FailureReason"`

	// <p>计算资源付费模式 (- PREPAID：预付费，即包年包月 - POSTPAID_BY_HOUR：按小时后付费)</p>
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// <p>后付费资源规格说明</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceTypeAlias *string `json:"InstanceTypeAlias,omitnil,omitempty" name:"InstanceTypeAlias"`

	// <p>预付费资源组名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupName *string `json:"ResourceGroupName,omitnil,omitempty" name:"ResourceGroupName"`

	// <p>存储的类型。取值包含：<br>    FREE:        预付费的免费存储<br>    CLOUD_PREMIUM： 高性能云硬盘<br>    CLOUD_SSD： SSD云硬盘<br>    CFS:     CFS存储，包含NFS和turbo</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	VolumeSourceType *string `json:"VolumeSourceType,omitnil,omitempty" name:"VolumeSourceType"`

	// <p>CFS存储的配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	VolumeSourceCFS *CFSConfig `json:"VolumeSourceCFS,omitnil,omitempty" name:"VolumeSourceCFS"`

	// <p>数据配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataConfigs []*DataConfig `json:"DataConfigs,omitnil,omitempty" name:"DataConfigs"`

	// <p>notebook 信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// <p>数据源来源，eg：WeData_HDFS</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSource *string `json:"DataSource,omitnil,omitempty" name:"DataSource"`

	// <p>镜像信息</p>
	ImageInfo *ImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`

	// <p>镜像类型</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageType *string `json:"ImageType,omitnil,omitempty" name:"ImageType"`

	// <p>SSH配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SSHConfig *SSHConfig `json:"SSHConfig,omitnil,omitempty" name:"SSHConfig"`

	// <p>GooseFS存储配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	VolumeSourceGooseFS *GooseFS `json:"VolumeSourceGooseFS,omitnil,omitempty" name:"VolumeSourceGooseFS"`

	// <p>子用户ID</p>
	SubUin *string `json:"SubUin,omitnil,omitempty" name:"SubUin"`

	// <p>调度节点ID</p>
	ResourceGroupInstanceId *string `json:"ResourceGroupInstanceId,omitnil,omitempty" name:"ResourceGroupInstanceId"`

	// <p>子用户名称</p>
	SubUinName *string `json:"SubUinName,omitnil,omitempty" name:"SubUinName"`

	// <p>任务实例创建时间</p>
	JobCreateTime *string `json:"JobCreateTime,omitnil,omitempty" name:"JobCreateTime"`

	// <p>Appid</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>描述</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>最后操作者信息</p>
	LatestOperatorInfo *OperatorInfo `json:"LatestOperatorInfo,omitnil,omitempty" name:"LatestOperatorInfo"`
}

type NotebookSetItem struct {
	// <p>notebook ID</p>
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>notebook 名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>计费模式</p>
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// <p>资源配置</p>
	ResourceConf *ResourceConf `json:"ResourceConf,omitnil,omitempty" name:"ResourceConf"`

	// <p>预付费资源组</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// <p>存储卷大小</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	VolumeSizeInGB *uint64 `json:"VolumeSizeInGB,omitnil,omitempty" name:"VolumeSizeInGB"`

	// <p>计费金额信息，eg：2.00元/小时 (for后付费)</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillingInfos []*string `json:"BillingInfos,omitnil,omitempty" name:"BillingInfos"`

	// <p>标签配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>创建时间</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>启动时间</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>更新时间</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>运行时间</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuntimeInSeconds *uint64 `json:"RuntimeInSeconds,omitnil,omitempty" name:"RuntimeInSeconds"`

	// <p>计费状态</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeStatus *string `json:"ChargeStatus,omitnil,omitempty" name:"ChargeStatus"`

	// <p>状态</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>错误原因</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailureReason *string `json:"FailureReason,omitnil,omitempty" name:"FailureReason"`

	// <p>结束时间</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>Pod名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodName *string `json:"PodName,omitnil,omitempty" name:"PodName"`

	// <p>后付费资源规格名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceTypeAlias *string `json:"InstanceTypeAlias,omitnil,omitempty" name:"InstanceTypeAlias"`

	// <p>预付费资源组名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupName *string `json:"ResourceGroupName,omitnil,omitempty" name:"ResourceGroupName"`

	// <p>是否自动终止</p>
	AutoStopping *bool `json:"AutoStopping,omitnil,omitempty" name:"AutoStopping"`

	// <p>自动停止时间</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutomaticStopTime *uint64 `json:"AutomaticStopTime,omitnil,omitempty" name:"AutomaticStopTime"`

	// <p>存储的类型。取值包含：<br>    FREE:        预付费的免费存储<br>    CLOUD_PREMIUM： 高性能云硬盘<br>    CLOUD_SSD： SSD云硬盘<br>    CFS:     CFS存储，包含NFS和turbo</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	VolumeSourceType *string `json:"VolumeSourceType,omitnil,omitempty" name:"VolumeSourceType"`

	// <p>CFS存储的配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	VolumeSourceCFS *CFSConfig `json:"VolumeSourceCFS,omitnil,omitempty" name:"VolumeSourceCFS"`

	// <p>notebook 信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// <p>notebook用户类型</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserTypes []*string `json:"UserTypes,omitnil,omitempty" name:"UserTypes"`

	// <p>SSH配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SSHConfig *SSHConfig `json:"SSHConfig,omitnil,omitempty" name:"SSHConfig"`

	// <p>GooseFS存储配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	VolumeSourceGooseFS *GooseFS `json:"VolumeSourceGooseFS,omitnil,omitempty" name:"VolumeSourceGooseFS"`

	// <p>子用户ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubUin *string `json:"SubUin,omitnil,omitempty" name:"SubUin"`

	// <p>子用户名称</p>
	SubUinName *string `json:"SubUinName,omitnil,omitempty" name:"SubUinName"`

	// <p>AppId</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>容器服务暴露端口配置</p>
	ExposePortConfig *ExposePortConfig `json:"ExposePortConfig,omitnil,omitempty" name:"ExposePortConfig"`

	// <p>描述</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>操作者信息</p>
	LatestOperatorInfo *OperatorInfo `json:"LatestOperatorInfo,omitnil,omitempty" name:"LatestOperatorInfo"`
}

type NumOrPercent struct {
	// Num,Percent ,分别表示数量和百分比，默认为 Num
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 数值
	Value *int64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type OperatorInfo struct {
	// <p>操作者子账号 UIN</p>
	SubUin *string `json:"SubUin,omitnil,omitempty" name:"SubUin"`

	// <p>操作者子账号名称</p>
	SubUinName *string `json:"SubUinName,omitnil,omitempty" name:"SubUinName"`

	// <p>是否为平台操作</p>
	IsPlatformOperator *bool `json:"IsPlatformOperator,omitnil,omitempty" name:"IsPlatformOperator"`

	// <p>操作类型</p>
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`
}

type Option struct {
	// 指标名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 指标值
	Value *int64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type PlatformImageInfo struct {
	// 框架名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Framework *string `json:"Framework,omitnil,omitempty" name:"Framework"`

	// 镜像类型: ccr or tcr
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageType *string `json:"ImageType,omitnil,omitempty" name:"ImageType"`

	// 镜像地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// TCR镜像示例所属地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegistryRegion *string `json:"RegistryRegion,omitnil,omitempty" name:"RegistryRegion"`

	// TCR镜像所属实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 镜像名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageName *string `json:"ImageName,omitnil,omitempty" name:"ImageName"`

	// 镜像Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 框架版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrameworkVersion *string `json:"FrameworkVersion,omitnil,omitempty" name:"FrameworkVersion"`

	// 支持的gpu列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SupportGpuList []*string `json:"SupportGpuList,omitnil,omitempty" name:"SupportGpuList"`

	// 描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 业务属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraAttributes []*Attribute `json:"ExtraAttributes,omitnil,omitempty" name:"ExtraAttributes"`

	// 镜像适用场景Train/Inference/Notebook
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageRange []*string `json:"ImageRange,omitnil,omitempty" name:"ImageRange"`

	// 是否支持分布式部署
	// 注意：此字段可能返回 null，表示取不到有效值。
	SupportDistributedDeploy *bool `json:"SupportDistributedDeploy,omitnil,omitempty" name:"SupportDistributedDeploy"`

	// 支持的地域 all(所有地域)/autonomous(自动驾驶地域)/general(通用地域)
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionScope *string `json:"RegionScope,omitnil,omitempty" name:"RegionScope"`
}

type Pod struct {
	// pod名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// pod的唯一id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uid *string `json:"Uid,omitnil,omitempty" name:"Uid"`

	// 服务付费模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// pod的状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Phase *string `json:"Phase,omitnil,omitempty" name:"Phase"`

	// pod的IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	IP *string `json:"IP,omitnil,omitempty" name:"IP"`

	// pod的创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 容器列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: Containers is deprecated.
	Containers *Container `json:"Containers,omitnil,omitempty" name:"Containers"`

	// 容器列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerInfos []*Container `json:"ContainerInfos,omitnil,omitempty" name:"ContainerInfos"`

	// 容器调用信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CrossTenantENIInfo *CrossTenantENIInfo `json:"CrossTenantENIInfo,omitnil,omitempty" name:"CrossTenantENIInfo"`

	// 实例的状态信息
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 实例的开始调度时间
	StartScheduleTime *string `json:"StartScheduleTime,omitnil,omitempty" name:"StartScheduleTime"`

	// 实例状态的补充信息
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 当前实例所在的节点 IP
	NodeIP *string `json:"NodeIP,omitnil,omitempty" name:"NodeIP"`

	// 当前实例所在节点id
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 当时实例所属资源组 id
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// 资源组名称
	ResourceGroupName *string `json:"ResourceGroupName,omitnil,omitempty" name:"ResourceGroupName"`

	// 实例的资源占用信息
	ResourceInfo *ResourceInfo `json:"ResourceInfo,omitnil,omitempty" name:"ResourceInfo"`
}

type PodInfo struct {
	// pod名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// pod的IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	IP *string `json:"IP,omitnil,omitempty" name:"IP"`

	// pod状态。eg：SUBMITTING提交中、PENDING排队中、RUNNING运行中、SUCCEEDED已完成、FAILED异常、TERMINATING停止中、TERMINATED已停止
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// pod启动时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// pod结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// pod资源配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceConfigInfo *ResourceConfigInfo `json:"ResourceConfigInfo,omitnil,omitempty" name:"ResourceConfigInfo"`

	// Pod所属任务的SubUin信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubUin *string `json:"SubUin,omitnil,omitempty" name:"SubUin"`
}

type PodSSHInfo struct {
	// pod访问ip
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// pod ssh访问端口
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// ssh访问命令
	LoginCommand *string `json:"LoginCommand,omitnil,omitempty" name:"LoginCommand"`
}

type PrivateLinkInfo struct {
	// 私有连接所在的VPCID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 私有连接所在的子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// HTTP内网调用地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerHttpAddr []*string `json:"InnerHttpAddr,omitnil,omitempty" name:"InnerHttpAddr"`

	// HTTPS内网调用地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerHttpsAddr []*string `json:"InnerHttpsAddr,omitnil,omitempty" name:"InnerHttpsAddr"`

	// 私有连接状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// grpc内网调用地址
	InnerGrpcAddr []*string `json:"InnerGrpcAddr,omitnil,omitempty" name:"InnerGrpcAddr"`
}

type Probe struct {
	// 探针行为
	ProbeAction *ProbeAction `json:"ProbeAction,omitnil,omitempty" name:"ProbeAction"`

	// 等待服务启动的延迟
	InitialDelaySeconds *int64 `json:"InitialDelaySeconds,omitnil,omitempty" name:"InitialDelaySeconds"`

	// 轮询检查时间间隔
	PeriodSeconds *int64 `json:"PeriodSeconds,omitnil,omitempty" name:"PeriodSeconds"`

	// 检查超时时长
	TimeoutSeconds *int64 `json:"TimeoutSeconds,omitnil,omitempty" name:"TimeoutSeconds"`

	// 检测失败认定次数
	FailureThreshold *int64 `json:"FailureThreshold,omitnil,omitempty" name:"FailureThreshold"`

	// 检测成功认定次数，就绪默认 3，存活/启动默认 1
	SuccessThreshold *int64 `json:"SuccessThreshold,omitnil,omitempty" name:"SuccessThreshold"`
}

type ProbeAction struct {
	// http get 行为
	HTTPGet *HTTPGetAction `json:"HTTPGet,omitnil,omitempty" name:"HTTPGet"`

	// 执行命令检查 行为
	Exec *ExecAction `json:"Exec,omitnil,omitempty" name:"Exec"`

	// tcp socket 检查行为
	TCPSocket *TCPSocketAction `json:"TCPSocket,omitnil,omitempty" name:"TCPSocket"`

	// 探针类型，默认 HTTPGet，可选值：HTTPGet、Exec、TCPSocket
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`
}

type PublicAlgoVersion struct {
	// 公共算法版本Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicAlgoVersionId *string `json:"PublicAlgoVersionId,omitnil,omitempty" name:"PublicAlgoVersionId"`

	// 对应的公共算法组Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicAlgoGroupId *string `json:"PublicAlgoGroupId,omitnil,omitempty" name:"PublicAlgoGroupId"`

	// 版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 模型简介
	// 注意：此字段可能返回 null，表示取不到有效值。
	Introduction *string `json:"Introduction,omitnil,omitempty" name:"Introduction"`

	// 预览信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PreviewInfo *string `json:"PreviewInfo,omitnil,omitempty" name:"PreviewInfo"`

	// 预置训练镜像信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PresetTrainImageInfo *ImageInfo `json:"PresetTrainImageInfo,omitnil,omitempty" name:"PresetTrainImageInfo"`

	// 预置训练代码信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PresetTrainCodeInfo *MaterialInfo `json:"PresetTrainCodeInfo,omitnil,omitempty" name:"PresetTrainCodeInfo"`

	// 预置模型信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PresetModelInfo *MaterialInfo `json:"PresetModelInfo,omitnil,omitempty" name:"PresetModelInfo"`

	// 是否已经被导入到我的算法
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsImported *bool `json:"IsImported,omitnil,omitempty" name:"IsImported"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 默认训练资源规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultResourceSpec *ResourceSpec `json:"DefaultResourceSpec,omitnil,omitempty" name:"DefaultResourceSpec"`

	// 默认推理资源规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultInferenceResourceSpec *ResourceSpec `json:"DefaultInferenceResourceSpec,omitnil,omitempty" name:"DefaultInferenceResourceSpec"`

	// 是否支持直接部署推理服务
	// 注意：此字段可能返回 null，表示取不到有效值。
	SupportDeploy *bool `json:"SupportDeploy,omitnil,omitempty" name:"SupportDeploy"`

	// 内置训练数据集
	// 注意：此字段可能返回 null，表示取不到有效值。
	PresetTrainDataset *MaterialInfo `json:"PresetTrainDataset,omitnil,omitempty" name:"PresetTrainDataset"`

	// 训练代码包下载路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrainCodeDownloadUrl *string `json:"TrainCodeDownloadUrl,omitnil,omitempty" name:"TrainCodeDownloadUrl"`

	// 内置数据下载路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrainDataDownloadUrl *string `json:"TrainDataDownloadUrl,omitnil,omitempty" name:"TrainDataDownloadUrl"`

	// 训练参数列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrainParams []*TrainParam `json:"TrainParams,omitnil,omitempty" name:"TrainParams"`

	// 训练启动命令
	// 注意：此字段可能返回 null，表示取不到有效值。
	PresetTrainCodeStartCmd *string `json:"PresetTrainCodeStartCmd,omitnil,omitempty" name:"PresetTrainCodeStartCmd"`

	// 是否非公开模型
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsPrivateModel *bool `json:"IsPrivateModel,omitnil,omitempty" name:"IsPrivateModel"`

	// 各种设备下的训练镜像
	// 注意：此字段可能返回 null，表示取不到有效值。
	PresetTrainImageInfoList []*DeviceImageInfo `json:"PresetTrainImageInfoList,omitnil,omitempty" name:"PresetTrainImageInfoList"`

	// 各种设备下的推理镜像
	// 注意：此字段可能返回 null，表示取不到有效值。
	PresetInferenceImageInfoList []*DeviceImageInfo `json:"PresetInferenceImageInfoList,omitnil,omitempty" name:"PresetInferenceImageInfoList"`

	// 各种设备下的训练代码信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PresetTrainCodeInfoList []*DeviceMaterialInfo `json:"PresetTrainCodeInfoList,omitnil,omitempty" name:"PresetTrainCodeInfoList"`

	// 各种设备下的内置模型信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PresetModelInfoList []*DeviceMaterialInfo `json:"PresetModelInfoList,omitnil,omitempty" name:"PresetModelInfoList"`

	// 模型类别，比如LLM/MultiModal
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelCategory *string `json:"ModelCategory,omitnil,omitempty" name:"ModelCategory"`

	// 公共算法Id
	PublicAlgoSeriesId *string `json:"PublicAlgoSeriesId,omitnil,omitempty" name:"PublicAlgoSeriesId"`
}

type PublicDataSourceFS struct {
	// 数据源id
	DataSourceId *string `json:"DataSourceId,omitnil,omitempty" name:"DataSourceId"`

	// 相对数据源子路径
	SubPath *string `json:"SubPath,omitnil,omitempty" name:"SubPath"`
}

// Predefined struct for user
type PushTrainingMetricsRequestParams struct {
	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`

	// 指标数据
	Data []*MetricData `json:"Data,omitnil,omitempty" name:"Data"`
}

type PushTrainingMetricsRequest struct {
	*tchttp.BaseRequest
	
	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`

	// 指标数据
	Data []*MetricData `json:"Data,omitnil,omitempty" name:"Data"`
}

func (r *PushTrainingMetricsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PushTrainingMetricsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TiProjectId")
	delete(f, "Data")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PushTrainingMetricsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PushTrainingMetricsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PushTrainingMetricsResponse struct {
	*tchttp.BaseResponse
	Response *PushTrainingMetricsResponseParams `json:"Response"`
}

func (r *PushTrainingMetricsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PushTrainingMetricsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RDMAConfig struct {
	// 是否开启RDMA
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`
}

type ResourceConf struct {
	// cpu 处理器资源, 单位为1/1000核 (for预付费)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// memory 内存资源, 单位为1M (for预付费)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// gpu Gpu卡资源，单位为1/100卡，例如GpuType=T4时，1 Gpu = 1/100 T4卡 (for预付费)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Gpu *uint64 `json:"Gpu,omitnil,omitempty" name:"Gpu"`

	// GpuType 卡类型，参考资源组上对应的卡类型。eg: H800,A800,A100,T4,P4,V100等 (for预付费)
	// 注意：此字段可能返回 null，表示取不到有效值。
	GpuType *string `json:"GpuType,omitnil,omitempty" name:"GpuType"`

	// 计算规格 (for后付费)，可选值如下：
	// TI.S.LARGE.POST: 4C8G 
	// TI.S.2XLARGE16.POST:  8C16G 
	// TI.S.2XLARGE32.POST:  8C32G 
	// TI.S.4XLARGE32.POST:  16C32G
	// TI.S.4XLARGE64.POST:  16C64G
	// TI.S.6XLARGE48.POST:  24C48G
	// TI.S.6XLARGE96.POST:  24C96G
	// TI.S.8XLARGE64.POST:  32C64G
	// TI.S.8XLARGE128.POST : 32C128G
	// TI.GN10.2XLARGE40.POST: 8C40G V100*1 
	// TI.GN10.5XLARGE80.POST:  18C80G V100*2 
	// TI.GN10.10XLARGE160.POST :  32C160G V100*4
	// TI.GN10.20XLARGE320.POST :  72C320G V100*8
	// TI.GN7.8XLARGE128.POST: 32C128G T4*1 
	// TI.GN7.10XLARGE160.POST: 40C160G T4*2 
	// TI.GN7.20XLARGE320.POST: 80C320G T4*4
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`
}

type ResourceConfigInfo struct {
	// 角色，eg：PS、WORKER、DRIVER、EXECUTOR
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// cpu核数，使用资源组时需配置。单位：1/1000，即1000=1核
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存，使用资源组时需配置。单位为MB
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// gpu卡类型，使用资源组时需配置
	GpuType *string `json:"GpuType,omitnil,omitempty" name:"GpuType"`

	// gpu卡数，使用资源组时需配置。单位：1/100，即100=1卡
	Gpu *uint64 `json:"Gpu,omitnil,omitempty" name:"Gpu"`

	// 算力规格ID
	// 计算规格 (for后付费)，可选值如下：
	// TI.S.LARGE.POST: 4C8G 
	// TI.S.2XLARGE16.POST:  8C16G 
	// TI.S.2XLARGE32.POST:  8C32G 
	// TI.S.4XLARGE32.POST:  16C32G
	// TI.S.4XLARGE64.POST:  16C64G
	// TI.S.6XLARGE48.POST:  24C48G
	// TI.S.6XLARGE96.POST:  24C96G
	// TI.S.8XLARGE64.POST:  32C64G
	// TI.S.8XLARGE128.POST : 32C128G
	// TI.GN10.2XLARGE40.POST: 8C40G V100*1 
	// TI.GN10.5XLARGE80.POST:  18C80G V100*2 
	// TI.GN10.10XLARGE160.POST :  32C160G V100*4
	// TI.GN10.20XLARGE320.POST :  72C320G V100*8
	// TI.GN7.8XLARGE128.POST: 32C128G T4*1 
	// TI.GN7.10XLARGE160.POST: 40C160G T4*2 
	// TI.GN7.20XLARGE320.POST: 80C32
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 计算节点数
	InstanceNum *uint64 `json:"InstanceNum,omitnil,omitempty" name:"InstanceNum"`

	// 算力规格名称
	// 计算规格 (for后付费)，可选值如下：
	// 4C8G 
	// 8C16G 
	// 8C32G 
	// 16C32G
	// 6C64G
	// 24C48G
	// 24C96G
	// 32C64G
	// 32C128G
	// 8C40G V100*1 
	// 8C80G V100*2 
	// 32C160G V100*4
	// 72C320G V100*8
	// 32C128G T4*1 
	// 40C160G T4*2 
	// 80C32
	InstanceTypeAlias *string `json:"InstanceTypeAlias,omitnil,omitempty" name:"InstanceTypeAlias"`

	// RDMA配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	RDMAConfig *RDMAConfig `json:"RDMAConfig,omitnil,omitempty" name:"RDMAConfig"`
}

type ResourceGroup struct {
	// 资源组id
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// 资源组名称
	ResourceGroupName *string `json:"ResourceGroupName,omitnil,omitempty" name:"ResourceGroupName"`

	// 可用节点个数(运行中的节点)
	FreeInstance *uint64 `json:"FreeInstance,omitnil,omitempty" name:"FreeInstance"`

	// 总节点个数(所有节点)
	TotalInstance *uint64 `json:"TotalInstance,omitnil,omitempty" name:"TotalInstance"`

	// 资源组已用的资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsedResource *GroupResource `json:"UsedResource,omitnil,omitempty" name:"UsedResource"`

	// 资源组总资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalResource *GroupResource `json:"TotalResource,omitnil,omitempty" name:"TotalResource"`

	// 节点信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceSet []*Instance `json:"InstanceSet,omitnil,omitempty" name:"InstanceSet"`

	// 标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagSet []*Tag `json:"TagSet,omitnil,omitempty" name:"TagSet"`
}

type ResourceGroupInWorkspace struct {
	// <p>资源组ID</p>
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// <p>资源组名称</p>
	ResourceGroupName *string `json:"ResourceGroupName,omitnil,omitempty" name:"ResourceGroupName"`

	// <p>地域</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>是否有运行中的任务/服务占用</p>
	Occupied *bool `json:"Occupied,omitnil,omitempty" name:"Occupied"`
}

type ResourceGroupInfo struct {
	// 资源组 id
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// 资源组名称
	ResourceGroupName *string `json:"ResourceGroupName,omitnil,omitempty" name:"ResourceGroupName"`
}

type ResourceInfo struct {
	// 处理器资源, 单位为1/1000核
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存资源, 单位为1M
	// 注意：此字段可能返回 null，表示取不到有效值。
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Gpu卡个数资源, 单位为0.01单位的GpuType.
	// Gpu=100表示使用了“一张”gpu卡, 但此处的“一张”卡有可能是虚拟化后的1/4卡, 也有可能是整张卡. 取决于实例的机型
	// 例1 实例的机型带有1张虚拟gpu卡, 每张虚拟gpu卡对应1/4张实际T4卡, 则此时 GpuType=T4, Gpu=100, RealGpu=25.
	// 例2 实例的机型带有4张gpu整卡, 每张卡对应1张实际T4卡, 则 此时 GpuType=T4, Gpu=400, RealGpu=400.
	// 注意：此字段可能返回 null，表示取不到有效值。
	Gpu *uint64 `json:"Gpu,omitnil,omitempty" name:"Gpu"`

	// Gpu卡型号 T4或者V100。仅展示当前 GPU 卡型号，若存在多类型同时使用，则参考 RealGpuDetailSet 的值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	GpuType *string `json:"GpuType,omitnil,omitempty" name:"GpuType"`

	// 创建或更新时无需填写，仅展示需要关注
	// 后付费非整卡实例对应的实际的Gpu卡资源, 表示gpu资源对应实际的gpu卡个数.
	// RealGpu=100表示实际使用了一张gpu卡, 对应实际的实例机型, 有可能代表带有1/4卡的实例4个, 或者带有1/2卡的实例2个, 或者带有1卡的实力1个.
	RealGpu *uint64 `json:"RealGpu,omitnil,omitempty" name:"RealGpu"`

	// 创建或更新时无需填写，仅展示需要关注。详细的GPU使用信息。
	RealGpuDetailSet []*GpuDetail `json:"RealGpuDetailSet,omitnil,omitempty" name:"RealGpuDetailSet"`

	// 是否开启rdma
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableRDMA *bool `json:"EnableRDMA,omitnil,omitempty" name:"EnableRDMA"`

	// root disk size(GB)
	RootDisk *uint64 `json:"RootDisk,omitnil,omitempty" name:"RootDisk"`

	// data disk size(GB)
	DataDisk *uint64 `json:"DataDisk,omitnil,omitempty" name:"DataDisk"`
}

type ResourceInstanceRunningJobInfo struct {
	// <p>pod名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodName *string `json:"PodName,omitnil,omitempty" name:"PodName"`

	// <p>任务类型</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// <p>任务id</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <p>任务自定义名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`
}

type ResourceSpec struct {
	// 规格简称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecAlias *string `json:"SpecAlias,omitnil,omitempty" name:"SpecAlias"`

	// 规格Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecId *string `json:"SpecId,omitnil,omitempty" name:"SpecId"`

	// 规则名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`
}

type RollingUpdate struct {
	// 滚动更新的最大不可用
	MaxUnavailable *NumOrPercent `json:"MaxUnavailable,omitnil,omitempty" name:"MaxUnavailable"`

	// 滚动更新的最大新增实例
	MaxSurge *NumOrPercent `json:"MaxSurge,omitnil,omitempty" name:"MaxSurge"`
}

type SSHConfig struct {
	// 是否开启ssh
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 公钥信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicKey *string `json:"PublicKey,omitnil,omitempty" name:"PublicKey"`

	// 端口号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 登录命令
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoginCommand *string `json:"LoginCommand,omitnil,omitempty" name:"LoginCommand"`

	// 登录地址是否改变
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAddressChanged *bool `json:"IsAddressChanged,omitnil,omitempty" name:"IsAddressChanged"`

	// POD访问信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodSSHInfo *PodSSHInfo `json:"PodSSHInfo,omitnil,omitempty" name:"PodSSHInfo"`
}

type ScheduledAction struct {
	// 是否要定时停止服务，true or false。true 则 ScheduleStopTime 必填， false 则 ScheduleStopTime 不生效
	ScheduleStop *bool `json:"ScheduleStop,omitnil,omitempty" name:"ScheduleStop"`

	// 要执行定时停止的时间，格式：“2022-01-26T19:46:22+08:00”
	ScheduleStopTime *string `json:"ScheduleStopTime,omitnil,omitempty" name:"ScheduleStopTime"`
}

type SchedulingPolicy struct {
	// 是否启用了跨资源组调度开关
	CrossResourceGroupScheduling *bool `json:"CrossResourceGroupScheduling,omitnil,omitempty" name:"CrossResourceGroupScheduling"`
}

type SchemaInfo struct {
	// 长度30字符内
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数据类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type Service struct {
	// 服务组id
	ServiceGroupId *string `json:"ServiceGroupId,omitnil,omitempty" name:"ServiceGroupId"`

	// 服务id
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 服务组名
	ServiceGroupName *string `json:"ServiceGroupName,omitnil,omitempty" name:"ServiceGroupName"`

	// 服务描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceDescription *string `json:"ServiceDescription,omitnil,omitempty" name:"ServiceDescription"`

	// 服务的详细信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceInfo *ServiceInfo `json:"ServiceInfo,omitnil,omitempty" name:"ServiceInfo"`

	// 集群id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 付费类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 包年包月服务的资源组id，按量计费的服务为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// 包年包月服务对应的资源组名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupName *string `json:"ResourceGroupName,omitnil,omitempty" name:"ResourceGroupName"`

	// 服务的标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 服务所在的 ingress 的 name
	// 注意：此字段可能返回 null，表示取不到有效值。
	IngressName *string `json:"IngressName,omitnil,omitempty" name:"IngressName"`

	// 创建者
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedBy *string `json:"CreatedBy,omitnil,omitempty" name:"CreatedBy"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 主账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 子账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubUin *string `json:"SubUin,omitnil,omitempty" name:"SubUin"`

	// app_id
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 服务的业务状态
	// CREATING 创建中
	// CREATE_FAILED 创建失败
	// CREATE_SUCCEED 创建成功
	// ARREARS_STOP 因欠费停止
	// WHITELIST_STOP 白名单额度不足
	// RELEASE_FAILED 资源释放失败
	// WHITELIST_RELEASE_FAILED 白名单资源释放失败
	// TIMEOUT_EXCEPTION 创建超时异常
	// BILLING 计费中
	// WHITELIST_USING 白名单试用中
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessStatus *string `json:"BusinessStatus,omitnil,omitempty" name:"BusinessStatus"`

	// 已废弃,以ServiceInfo中的对应为准
	//
	// Deprecated: ServiceLimit is deprecated.
	ServiceLimit *ServiceLimit `json:"ServiceLimit,omitnil,omitempty" name:"ServiceLimit"`

	// 已废弃,以ServiceInfo中的对应为准
	//
	// Deprecated: ScheduledAction is deprecated.
	ScheduledAction *ScheduledAction `json:"ScheduledAction,omitnil,omitempty" name:"ScheduledAction"`

	// 服务创建失败的原因，创建成功后该字段为默认值 CREATE_SUCCEED
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateFailedReason *string `json:"CreateFailedReason,omitnil,omitempty" name:"CreateFailedReason"`

	// 服务状态
	// CREATING 创建中
	// CREATE_FAILED 创建失败
	// TIMEOUT_EXCEPTION 创建超时异常
	// Normal 正常运行中
	// Stopped 已停止
	// Stopping 停止中
	// Abnormal 异常
	// Pending 启动中
	// Waiting 就绪中
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 费用信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillingInfo *string `json:"BillingInfo,omitnil,omitempty" name:"BillingInfo"`

	// 模型权重
	// 注意：此字段可能返回 null，表示取不到有效值。
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// 服务的创建来源
	// AUTO_ML: 来自自动学习的一键发布
	// DEFAULT: 其他来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateSource *string `json:"CreateSource,omitnil,omitempty" name:"CreateSource"`

	// 版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 服务组下服务的最高版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestVersion *string `json:"LatestVersion,omitnil,omitempty" name:"LatestVersion"`

	// 资源组类别 托管 NORMAL，纳管 SW
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupSWType *string `json:"ResourceGroupSWType,omitnil,omitempty" name:"ResourceGroupSWType"`

	// 服务的归档状态  Waiting 等待归档中，Archived 已归档
	// 注意：此字段可能返回 null，表示取不到有效值。
	ArchiveStatus *string `json:"ArchiveStatus,omitnil,omitempty" name:"ArchiveStatus"`

	// 服务的部署类型 [STANDARD 标准部署，DIST 分布式多机部署] 默认STANDARD
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployType *string `json:"DeployType,omitnil,omitempty" name:"DeployType"`

	// 单副本下的实例数，仅在部署类型为DIST时生效，默认1
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstancePerReplicas *string `json:"InstancePerReplicas,omitnil,omitempty" name:"InstancePerReplicas"`

	// 用于监控查询的Source
	// 枚举值，部分情况下与CreateSource不同，通过该字段兼容
	MonitorSource *string `json:"MonitorSource,omitnil,omitempty" name:"MonitorSource"`

	// 服务创建者的子账号名称
	SubUinName *string `json:"SubUinName,omitnil,omitempty" name:"SubUinName"`

	// 服务的调度策略
	SchedulingPolicy *SchedulingPolicy `json:"SchedulingPolicy,omitnil,omitempty" name:"SchedulingPolicy"`

	// 外部的资源组信息，表示借调了哪些资源组的资源
	ExternalResourceGroups []*ResourceGroupInfo `json:"ExternalResourceGroups,omitnil,omitempty" name:"ExternalResourceGroups"`
}

type ServiceCallInfo struct {
	// 服务组id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceGroupId *string `json:"ServiceGroupId,omitnil,omitempty" name:"ServiceGroupId"`

	// 内网http调用地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerHttpAddr *string `json:"InnerHttpAddr,omitnil,omitempty" name:"InnerHttpAddr"`

	// 内网https调用地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerHttpsAddr *string `json:"InnerHttpsAddr,omitnil,omitempty" name:"InnerHttpsAddr"`

	// 内网http调用地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	OuterHttpAddr *string `json:"OuterHttpAddr,omitnil,omitempty" name:"OuterHttpAddr"`

	// 内网https调用地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	OuterHttpsAddr *string `json:"OuterHttpsAddr,omitnil,omitempty" name:"OuterHttpsAddr"`

	// 调用key
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// 调用secret
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppSecret *string `json:"AppSecret,omitnil,omitempty" name:"AppSecret"`

	// 鉴权是否开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthorizationEnable *bool `json:"AuthorizationEnable,omitnil,omitempty" name:"AuthorizationEnable"`
}

type ServiceCallInfoV2 struct {
	// 服务组id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceGroupId *string `json:"ServiceGroupId,omitnil,omitempty" name:"ServiceGroupId"`

	// 服务的公网调用地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	InternetEndpoint *string `json:"InternetEndpoint,omitnil,omitempty" name:"InternetEndpoint"`

	// 鉴权是否开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthorizationEnable *bool `json:"AuthorizationEnable,omitnil,omitempty" name:"AuthorizationEnable"`

	// 鉴权token，仅当AuthorizationEnable为true时有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthToken *string `json:"AuthToken,omitnil,omitempty" name:"AuthToken"`

	// LLM token 列表
	AuthTokens []*AuthToken `json:"AuthTokens,omitnil,omitempty" name:"AuthTokens"`

	// 是否开启限流
	EnableLimit *bool `json:"EnableLimit,omitnil,omitempty" name:"EnableLimit"`

	// 访问grpc时需携带的虚拟Host
	GrpcHost *string `json:"GrpcHost,omitnil,omitempty" name:"GrpcHost"`

	// 网关相关配置
	GatewayConfig *GatewayConfig `json:"GatewayConfig,omitnil,omitempty" name:"GatewayConfig"`
}

type ServiceEIP struct {
	// 是否开启TIONE内网到外部的访问
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableEIP *bool `json:"EnableEIP,omitnil,omitempty" name:"EnableEIP"`

	// 用户VpcId
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 用户subnetId
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`
}

type ServiceEIPInfo struct {
	// 服务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 用户VpcId
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 用户子网Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`
}

type ServiceGroup struct {
	// <p>服务组id</p>
	ServiceGroupId *string `json:"ServiceGroupId,omitnil,omitempty" name:"ServiceGroupId"`

	// <p>服务组名</p>
	ServiceGroupName *string `json:"ServiceGroupName,omitnil,omitempty" name:"ServiceGroupName"`

	// <p>创建者</p>
	CreatedBy *string `json:"CreatedBy,omitnil,omitempty" name:"CreatedBy"`

	// <p>创建时间</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>更新时间</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>主账号</p>
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// <p>服务组下服务总数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceCount *uint64 `json:"ServiceCount,omitnil,omitempty" name:"ServiceCount"`

	// <p>服务组下在运行的服务数量</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunningServiceCount *uint64 `json:"RunningServiceCount,omitnil,omitempty" name:"RunningServiceCount"`

	// <p>服务描述</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Services []*Service `json:"Services,omitnil,omitempty" name:"Services"`

	// <p>服务组状态，与服务一致<br> CREATING 创建中<br>     CREATE_FAILED 创建失败<br>     Normal    正常运行中<br>     Stopped  已停止<br>     Stopping 停止中<br>     Abnormal 异常<br>     Pending 启动中<br>     Waiting 就绪中</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>服务组标签</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>服务组下最高版本</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestVersion *string `json:"LatestVersion,omitnil,omitempty" name:"LatestVersion"`

	// <p>服务的业务状态<br>CREATING 创建中<br>     CREATE_FAILED 创建失败<br>     ARREARS_STOP 因欠费被强制停止<br>     BILLING 计费中<br>     WHITELIST_USING 白名单试用中<br>     WHITELIST_STOP 白名单额度不足</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessStatus *string `json:"BusinessStatus,omitnil,omitempty" name:"BusinessStatus"`

	// <p>服务的计费信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillingInfo *string `json:"BillingInfo,omitnil,omitempty" name:"BillingInfo"`

	// <p>服务的创建来源</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateSource *string `json:"CreateSource,omitnil,omitempty" name:"CreateSource"`

	// <p>服务组的权重更新状态<br>UPDATING 更新中<br>     UPDATED 更新成功<br>     UPDATE_FAILED 更新失败</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WeightUpdateStatus *string `json:"WeightUpdateStatus,omitnil,omitempty" name:"WeightUpdateStatus"`

	// <p>服务组下运行的pod数量</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReplicasCount *uint64 `json:"ReplicasCount,omitnil,omitempty" name:"ReplicasCount"`

	// <p>服务组下期望的pod数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvailableReplicasCount *uint64 `json:"AvailableReplicasCount,omitnil,omitempty" name:"AvailableReplicasCount"`

	// <p>服务组的subuin</p>
	SubUin *string `json:"SubUin,omitnil,omitempty" name:"SubUin"`

	// <p>服务组的app_id</p>
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>是否开启鉴权</p>
	AuthorizationEnable *bool `json:"AuthorizationEnable,omitnil,omitempty" name:"AuthorizationEnable"`

	// <p>限流鉴权 token 列表</p>
	AuthTokens []*AuthToken `json:"AuthTokens,omitnil,omitempty" name:"AuthTokens"`

	// <p>用于监控的创建来源字段</p>
	MonitorSource *string `json:"MonitorSource,omitnil,omitempty" name:"MonitorSource"`

	// <p>子用户的 nickname</p>
	SubUinName *string `json:"SubUinName,omitnil,omitempty" name:"SubUinName"`

	// <p>网关日志投递相关配置</p>
	GatewayLogConfig *LogConfig `json:"GatewayLogConfig,omitnil,omitempty" name:"GatewayLogConfig"`

	// <p>网关路由相关配置</p>
	GatewayConfig *GatewayConfig `json:"GatewayConfig,omitnil,omitempty" name:"GatewayConfig"`
}

type ServiceInfo struct {
	// <p>期望运行的Pod数量，停止状态是0<br>不同计费模式和调节模式下对应关系如下<br>PREPAID 和 POSTPAID_BY_HOUR:<br>手动调节模式下对应 实例数量<br>自动调节模式下对应 基于时间的默认策略的实例数量<br>HYBRID_PAID:<br>后付费实例手动调节模式下对应 实例数量<br>后付费实例自动调节模式下对应 时间策略的默认策略的实例数量</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Replicas *int64 `json:"Replicas,omitnil,omitempty" name:"Replicas"`

	// <p>镜像信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageInfo *ImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`

	// <p>环境变量</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Env []*EnvVar `json:"Env,omitnil,omitempty" name:"Env"`

	// <p>资源信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resources *ResourceInfo `json:"Resources,omitnil,omitempty" name:"Resources"`

	// <p>后付费实例对应的机型规格</p>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// <p>模型信息</p>
	ModelInfo *ModelInfo `json:"ModelInfo,omitnil,omitempty" name:"ModelInfo"`

	// <p>是否启用日志</p>
	LogEnable *bool `json:"LogEnable,omitnil,omitempty" name:"LogEnable"`

	// <p>日志配置</p>
	LogConfig *LogConfig `json:"LogConfig,omitnil,omitempty" name:"LogConfig"`

	// <p>是否开启鉴权</p>
	AuthorizationEnable *bool `json:"AuthorizationEnable,omitnil,omitempty" name:"AuthorizationEnable"`

	// <p>hpa配置</p>
	HorizontalPodAutoscaler *HorizontalPodAutoscaler `json:"HorizontalPodAutoscaler,omitnil,omitempty" name:"HorizontalPodAutoscaler"`

	// <p>服务的状态描述</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *WorkloadStatus `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>权重</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// <p>资源总量</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceTotal *ResourceInfo `json:"ResourceTotal,omitnil,omitempty" name:"ResourceTotal"`

	// <p>历史实例数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	OldReplicas *int64 `json:"OldReplicas,omitnil,omitempty" name:"OldReplicas"`

	// <p>计费模式[HYBRID_PAID]时生效, 用于标识混合计费模式下的预付费实例数, 若不填则默认为1</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	HybridBillingPrepaidReplicas *int64 `json:"HybridBillingPrepaidReplicas,omitnil,omitempty" name:"HybridBillingPrepaidReplicas"`

	// <p>历史 HYBRID_PAID 时的实例数，用户恢复服务</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	OldHybridBillingPrepaidReplicas *int64 `json:"OldHybridBillingPrepaidReplicas,omitnil,omitempty" name:"OldHybridBillingPrepaidReplicas"`

	// <p>是否开启模型的热更新。默认不开启</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelHotUpdateEnable *bool `json:"ModelHotUpdateEnable,omitnil,omitempty" name:"ModelHotUpdateEnable"`

	// <p>服务的规格别名</p>
	InstanceAlias *string `json:"InstanceAlias,omitnil,omitempty" name:"InstanceAlias"`

	// <p>实例数量调节方式,默认为手动<br>支持：自动 - &quot;AUTO&quot;, 手动 - &quot;MANUAL&quot;</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScaleMode *string `json:"ScaleMode,omitnil,omitempty" name:"ScaleMode"`

	// <p>定时伸缩任务</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CronScaleJobs []*CronScaleJob `json:"CronScaleJobs,omitnil,omitempty" name:"CronScaleJobs"`

	// <p>定时伸缩策略</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScaleStrategy *string `json:"ScaleStrategy,omitnil,omitempty" name:"ScaleStrategy"`

	// <p>定时停止的配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScheduledAction *ScheduledAction `json:"ScheduledAction,omitnil,omitempty" name:"ScheduledAction"`

	// <p>实例列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: PodList is deprecated.
	PodList []*string `json:"PodList,omitnil,omitempty" name:"PodList"`

	// <p>Pod列表信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: Pods is deprecated.
	Pods *Pod `json:"Pods,omitnil,omitempty" name:"Pods"`

	// <p>Pod列表信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodInfos []*Pod `json:"PodInfos,omitnil,omitempty" name:"PodInfos"`

	// <p>服务限速限流相关配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceLimit *ServiceLimit `json:"ServiceLimit,omitnil,omitempty" name:"ServiceLimit"`

	// <p>是否开启模型的加速, 仅对StableDiffusion(动态加速)格式的模型有效。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelTurboEnable *bool `json:"ModelTurboEnable,omitnil,omitempty" name:"ModelTurboEnable"`

	// <p>挂载</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	VolumeMount *VolumeMount `json:"VolumeMount,omitnil,omitempty" name:"VolumeMount"`

	// <p>推理代码信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	InferCodeInfo *InferCodeInfo `json:"InferCodeInfo,omitnil,omitempty" name:"InferCodeInfo"`

	// <p>服务的启动命令</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// <p>开启TIONE内网访问外部设置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceEIP *ServiceEIP `json:"ServiceEIP,omitnil,omitempty" name:"ServiceEIP"`

	// <p>服务端口，默认为8501</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServicePort *int64 `json:"ServicePort,omitnil,omitempty" name:"ServicePort"`

	// <p>服务的优雅退出时限。单位为秒，默认值为30，最小为1</p>
	TerminationGracePeriodSeconds *int64 `json:"TerminationGracePeriodSeconds,omitnil,omitempty" name:"TerminationGracePeriodSeconds"`

	// <p>服务实例停止前执行的命令，执行完毕或执行时间超过优雅退出时限后实例结束</p>
	PreStopCommand []*string `json:"PreStopCommand,omitnil,omitempty" name:"PreStopCommand"`

	// <p>是否启用grpc端口</p>
	GrpcEnable *bool `json:"GrpcEnable,omitnil,omitempty" name:"GrpcEnable"`

	// <p>健康探针</p>
	HealthProbe *HealthProbe `json:"HealthProbe,omitnil,omitempty" name:"HealthProbe"`

	// <p>滚动更新配置</p>
	RollingUpdate *RollingUpdate `json:"RollingUpdate,omitnil,omitempty" name:"RollingUpdate"`

	// <p>单副本下的实例数，仅在部署类型为DIST、ROLE时生效，默认1</p>
	InstancePerReplicas *int64 `json:"InstancePerReplicas,omitnil,omitempty" name:"InstancePerReplicas"`

	// <p>批量数据盘挂载配置</p>
	VolumeMounts []*VolumeMount `json:"VolumeMounts,omitnil,omitempty" name:"VolumeMounts"`

	// <p>调度策略 [binpack] 优先占满整机，尽量避免碎卡（默认值）[spread] 优先分散在各个节点，确保服务高可用</p>
	SchedulingStrategy *string `json:"SchedulingStrategy,omitnil,omitempty" name:"SchedulingStrategy"`

	// <p>服务实际运行的节点数</p>
	NodeCount *int64 `json:"NodeCount,omitnil,omitempty" name:"NodeCount"`
}

type ServiceLimit struct {
	// 是否开启实例层面限流限速，true or false。true 则 InstanceRpsLimit 必填， false 则 InstanceRpsLimit 不生效
	EnableInstanceRpsLimit *bool `json:"EnableInstanceRpsLimit,omitnil,omitempty" name:"EnableInstanceRpsLimit"`

	// 每个服务实例的 request per second 限速, 0 为不限流
	InstanceRpsLimit *int64 `json:"InstanceRpsLimit,omitnil,omitempty" name:"InstanceRpsLimit"`

	// 是否开启单实例最大并发数限制，true or false。true 则 InstanceReqLimit 必填， false 则 InstanceReqLimit 不生效
	EnableInstanceReqLimit *bool `json:"EnableInstanceReqLimit,omitnil,omitempty" name:"EnableInstanceReqLimit"`

	// 每个服务实例的最大并发
	InstanceReqLimit *int64 `json:"InstanceReqLimit,omitnil,omitempty" name:"InstanceReqLimit"`
}

type SidecarSpec struct {
	// 镜像配置
	ImageInfo *ImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`
}

type Spec struct {
	// 计费项标签
	SpecId *string `json:"SpecId,omitnil,omitempty" name:"SpecId"`

	// 计费项名称
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// 计费项显示名称
	SpecAlias *string `json:"SpecAlias,omitnil,omitempty" name:"SpecAlias"`

	// 是否售罄
	Available *bool `json:"Available,omitnil,omitempty" name:"Available"`

	// 当前资源售罄时，可用的区域有哪些
	AvailableRegion []*string `json:"AvailableRegion,omitnil,omitempty" name:"AvailableRegion"`

	// 当前计费项支持的特性
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecFeatures []*string `json:"SpecFeatures,omitnil,omitempty" name:"SpecFeatures"`

	// 计费项类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecType *string `json:"SpecType,omitnil,omitempty" name:"SpecType"`

	// GPU类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	GpuType *string `json:"GpuType,omitnil,omitempty" name:"GpuType"`

	// 计费项CategoryId
	// 注意：此字段可能返回 null，表示取不到有效值。
	CategoryId *string `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`
}

type SpecPrice struct {
	// 计费项名称
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// 原价，单位：分。最大值42亿，超过则返回0
	TotalCost *uint64 `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`

	// 优惠后的价格，单位：分
	RealTotalCost *uint64 `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// 计费项数量
	SpecCount *uint64 `json:"SpecCount,omitnil,omitempty" name:"SpecCount"`
}

type SpecUnit struct {
	// 计费项名称
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// 计费项数量,建议不超过100万
	SpecCount *uint64 `json:"SpecCount,omitnil,omitempty" name:"SpecCount"`
}

type StartCmdInfo struct {
	// 启动命令
	StartCmd *string `json:"StartCmd,omitnil,omitempty" name:"StartCmd"`

	// ps启动命令
	PsStartCmd *string `json:"PsStartCmd,omitnil,omitempty" name:"PsStartCmd"`

	// worker启动命令
	WorkerStartCmd *string `json:"WorkerStartCmd,omitnil,omitempty" name:"WorkerStartCmd"`
}

// Predefined struct for user
type StartNotebookRequestParams struct {
	// notebook id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type StartNotebookRequest struct {
	*tchttp.BaseRequest
	
	// notebook id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

func (r *StartNotebookRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartNotebookRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartNotebookRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartNotebookResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartNotebookResponse struct {
	*tchttp.BaseResponse
	Response *StartNotebookResponseParams `json:"Response"`
}

func (r *StartNotebookResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartNotebookResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartTrainingTaskRequestParams struct {
	// 训练任务ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`
}

type StartTrainingTaskRequest struct {
	*tchttp.BaseRequest
	
	// 训练任务ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`
}

func (r *StartTrainingTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartTrainingTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "TiProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartTrainingTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartTrainingTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartTrainingTaskResponse struct {
	*tchttp.BaseResponse
	Response *StartTrainingTaskResponseParams `json:"Response"`
}

func (r *StartTrainingTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartTrainingTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StatefulSetCondition struct {
	// 信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// Status of the condition, one of True, False, Unknown.
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 上次更新的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastTransitionTime *string `json:"LastTransitionTime,omitnil,omitempty" name:"LastTransitionTime"`

	// 上次更新的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastUpdateTime *string `json:"LastUpdateTime,omitnil,omitempty" name:"LastUpdateTime"`
}

// Predefined struct for user
type StopModelAccelerateTaskRequestParams struct {
	// 模型加速任务ID
	ModelAccTaskId *string `json:"ModelAccTaskId,omitnil,omitempty" name:"ModelAccTaskId"`
}

type StopModelAccelerateTaskRequest struct {
	*tchttp.BaseRequest
	
	// 模型加速任务ID
	ModelAccTaskId *string `json:"ModelAccTaskId,omitnil,omitempty" name:"ModelAccTaskId"`
}

func (r *StopModelAccelerateTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopModelAccelerateTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelAccTaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopModelAccelerateTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopModelAccelerateTaskResponseParams struct {
	// 模型加速任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelAccTaskId *string `json:"ModelAccTaskId,omitnil,omitempty" name:"ModelAccTaskId"`

	// 异步任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsyncTaskId *string `json:"AsyncTaskId,omitnil,omitempty" name:"AsyncTaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopModelAccelerateTaskResponse struct {
	*tchttp.BaseResponse
	Response *StopModelAccelerateTaskResponseParams `json:"Response"`
}

func (r *StopModelAccelerateTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopModelAccelerateTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopNotebookRequestParams struct {
	// notebook id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type StopNotebookRequest struct {
	*tchttp.BaseRequest
	
	// notebook id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

func (r *StopNotebookRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopNotebookRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopNotebookRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopNotebookResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopNotebookResponse struct {
	*tchttp.BaseResponse
	Response *StopNotebookResponseParams `json:"Response"`
}

func (r *StopNotebookResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopNotebookResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopTrainingTaskRequestParams struct {
	// 训练任务ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`
}

type StopTrainingTaskRequest struct {
	*tchttp.BaseRequest
	
	// 训练任务ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`
}

func (r *StopTrainingTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopTrainingTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "TiProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopTrainingTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopTrainingTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopTrainingTaskResponse struct {
	*tchttp.BaseResponse
	Response *StopTrainingTaskResponseParams `json:"Response"`
}

func (r *StopTrainingTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopTrainingTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StorageExtraConf struct {
	// cfs的存储类型
	//   // HP:通用性能型
	//   // SD:通用标准型
	//   // TP:turbo性能型
	//   // TB:turbo标准型
	//   // THP:吞吐型
	CFSStorageType *string `json:"CFSStorageType,omitnil,omitempty" name:"CFSStorageType"`

	// cfs的协议
	CFSProtocol *string `json:"CFSProtocol,omitnil,omitempty" name:"CFSProtocol"`
}

type SubAccountInfo struct {
	// 腾讯云主账号UIN
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 腾讯云子账号UIN
	SubUin *string `json:"SubUin,omitnil,omitempty" name:"SubUin"`

	// 子账号名称
	SubUinName *string `json:"SubUinName,omitnil,omitempty" name:"SubUinName"`

	// 子账号在Linux下的UID
	LinuxUid *int64 `json:"LinuxUid,omitnil,omitempty" name:"LinuxUid"`

	// 子账号在Linux下的GID
	LinuxGid *int64 `json:"LinuxGid,omitnil,omitempty" name:"LinuxGid"`

	// 子账号在Linux下的用户名
	LinuxUserName *string `json:"LinuxUserName,omitnil,omitempty" name:"LinuxUserName"`

	// 是否开启 root 登录
	EnableRootLogin *bool `json:"EnableRootLogin,omitnil,omitempty" name:"EnableRootLogin"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type TCPSocketAction struct {
	// 调用端口
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`
}

type TJCallInfo struct {
	// 调用地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpAddr *string `json:"HttpAddr,omitnil,omitempty" name:"HttpAddr"`

	// token
	// 注意：此字段可能返回 null，表示取不到有效值。
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// 调用示例
	// 注意：此字段可能返回 null，表示取不到有效值。
	CallExample *string `json:"CallExample,omitnil,omitempty" name:"CallExample"`
}

type Tag struct {
	// 标签键
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TagFilter struct {
	// 标签键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 多个标签值
	TagValues []*string `json:"TagValues,omitnil,omitempty" name:"TagValues"`
}

type TrainParam struct {
	// 参数名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 默认参数值
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultValue *string `json:"DefaultValue,omitnil,omitempty" name:"DefaultValue"`

	// 参数注释
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 参数类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 是否必选
	Required *bool `json:"Required,omitnil,omitempty" name:"Required"`

	// 参数值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 参数范围
	Range []*string `json:"Range,omitnil,omitempty" name:"Range"`

	// 参数选项
	Enum []*string `json:"Enum,omitnil,omitempty" name:"Enum"`
}

type TrainingModelVersionDTO struct {
	// 模型id
	TrainingModelId *string `json:"TrainingModelId,omitnil,omitempty" name:"TrainingModelId"`

	// 模型版本id
	TrainingModelVersionId *string `json:"TrainingModelVersionId,omitnil,omitempty" name:"TrainingModelVersionId"`

	// 模型版本
	TrainingModelVersion *string `json:"TrainingModelVersion,omitnil,omitempty" name:"TrainingModelVersion"`

	// 模型来源
	TrainingModelSource *string `json:"TrainingModelSource,omitnil,omitempty" name:"TrainingModelSource"`

	// 创建时间
	TrainingModelCreateTime *string `json:"TrainingModelCreateTime,omitnil,omitempty" name:"TrainingModelCreateTime"`

	// 创建人uin
	TrainingModelCreator *string `json:"TrainingModelCreator,omitnil,omitempty" name:"TrainingModelCreator"`

	// 算法框架
	AlgorithmFramework *string `json:"AlgorithmFramework,omitnil,omitempty" name:"AlgorithmFramework"`

	// 推理环境
	ReasoningEnvironment *string `json:"ReasoningEnvironment,omitnil,omitempty" name:"ReasoningEnvironment"`

	// 推理环境来源
	ReasoningEnvironmentSource *string `json:"ReasoningEnvironmentSource,omitnil,omitempty" name:"ReasoningEnvironmentSource"`

	// 模型指标
	TrainingModelIndex *string `json:"TrainingModelIndex,omitnil,omitempty" name:"TrainingModelIndex"`

	// 训练任务名称
	TrainingJobName *string `json:"TrainingJobName,omitnil,omitempty" name:"TrainingJobName"`

	// 模型cos路径
	TrainingModelCosPath *CosPathInfo `json:"TrainingModelCosPath,omitnil,omitempty" name:"TrainingModelCosPath"`

	// 模型名称
	TrainingModelName *string `json:"TrainingModelName,omitnil,omitempty" name:"TrainingModelName"`

	// 训练任务id
	TrainingJobId *string `json:"TrainingJobId,omitnil,omitempty" name:"TrainingJobId"`

	// 自定义推理环境
	ReasoningImageInfo *ImageInfo `json:"ReasoningImageInfo,omitnil,omitempty" name:"ReasoningImageInfo"`

	// 模型版本创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 模型处理状态
	// STATUS_SUCCESS：导入成功，STATUS_FAILED：导入失败 ，STATUS_RUNNING：导入中
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrainingModelStatus *string `json:"TrainingModelStatus,omitnil,omitempty" name:"TrainingModelStatus"`

	// 模型处理进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrainingModelProgress *uint64 `json:"TrainingModelProgress,omitnil,omitempty" name:"TrainingModelProgress"`

	// 模型错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrainingModelErrorMsg *string `json:"TrainingModelErrorMsg,omitnil,omitempty" name:"TrainingModelErrorMsg"`

	// 模型格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrainingModelFormat *string `json:"TrainingModelFormat,omitnil,omitempty" name:"TrainingModelFormat"`

	// 模型版本类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionType *string `json:"VersionType,omitnil,omitempty" name:"VersionType"`

	// GPU类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	GPUType *string `json:"GPUType,omitnil,omitempty" name:"GPUType"`

	// 模型自动清理开关
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoClean *string `json:"AutoClean,omitnil,omitempty" name:"AutoClean"`

	// 模型清理周期
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelCleanPeriod *uint64 `json:"ModelCleanPeriod,omitnil,omitempty" name:"ModelCleanPeriod"`

	// 模型数量保留上限
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxReservedModels *uint64 `json:"MaxReservedModels,omitnil,omitempty" name:"MaxReservedModels"`

	// 模型热更新目录
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelHotUpdatePath *CosPathInfo `json:"ModelHotUpdatePath,omitnil,omitempty" name:"ModelHotUpdatePath"`

	// 推理环境id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReasoningEnvironmentId *string `json:"ReasoningEnvironmentId,omitnil,omitempty" name:"ReasoningEnvironmentId"`

	// 训练任务版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrainingJobVersion *string `json:"TrainingJobVersion,omitnil,omitempty" name:"TrainingJobVersion"`

	// 训练偏好
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrainingPreference *string `json:"TrainingPreference,omitnil,omitempty" name:"TrainingPreference"`

	// 自动学习任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoMLTaskId *string `json:"AutoMLTaskId,omitnil,omitempty" name:"AutoMLTaskId"`

	// 是否QAT模型
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsQAT *bool `json:"IsQAT,omitnil,omitempty" name:"IsQAT"`
}

type TrainingTaskDetail struct {
	// <p>训练任务ID</p>
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>训练任务名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>主账号uin</p>
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// <p>子账号uin</p>
	SubUin *string `json:"SubUin,omitnil,omitempty" name:"SubUin"`

	// <p>创建者名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubUinName *string `json:"SubUinName,omitnil,omitempty" name:"SubUinName"`

	// <p>地域</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>训练框架名称，eg：SPARK、PYSARK、TENSORFLOW、PYTORCH</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrameworkName *string `json:"FrameworkName,omitnil,omitempty" name:"FrameworkName"`

	// <p>训练框架版本</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrameworkVersion *string `json:"FrameworkVersion,omitnil,omitempty" name:"FrameworkVersion"`

	// <p>框架运行环境</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrameworkEnvironment *string `json:"FrameworkEnvironment,omitnil,omitempty" name:"FrameworkEnvironment"`

	// <p>计费模式</p>
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// <p>预付费专用资源组</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// <p>资源配置</p>
	ResourceConfigInfos []*ResourceConfigInfo `json:"ResourceConfigInfos,omitnil,omitempty" name:"ResourceConfigInfos"`

	// <p>标签</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>训练模式，eg：PS_WORKER、DDP、MPI、HOROVOD</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrainingMode *string `json:"TrainingMode,omitnil,omitempty" name:"TrainingMode"`

	// <p>代码包</p>
	CodePackagePath *CosPathInfo `json:"CodePackagePath,omitnil,omitempty" name:"CodePackagePath"`

	// <p>启动命令信息</p>
	StartCmdInfo *StartCmdInfo `json:"StartCmdInfo,omitnil,omitempty" name:"StartCmdInfo"`

	// <p>数据来源，eg：DATASET、COS</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSource *string `json:"DataSource,omitnil,omitempty" name:"DataSource"`

	// <p>数据配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataConfigs []*DataConfig `json:"DataConfigs,omitnil,omitempty" name:"DataConfigs"`

	// <p>调优参数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TuningParameters *string `json:"TuningParameters,omitnil,omitempty" name:"TuningParameters"`

	// <p>训练输出</p>
	Output *CosPathInfo `json:"Output,omitnil,omitempty" name:"Output"`

	// <p>是否上报日志</p>
	LogEnable *bool `json:"LogEnable,omitnil,omitempty" name:"LogEnable"`

	// <p>日志配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogConfig *LogConfig `json:"LogConfig,omitnil,omitempty" name:"LogConfig"`

	// <p>VPC ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>子网ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>自定义镜像信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageInfo *ImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`

	// <p>运行时长</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuntimeInSeconds *uint64 `json:"RuntimeInSeconds,omitnil,omitempty" name:"RuntimeInSeconds"`

	// <p>创建时间</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>训练开始时间</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>计费状态，eg：BILLING计费中，ARREARS_STOP欠费停止，NOT_BILLING不在计费中</p>
	ChargeStatus *string `json:"ChargeStatus,omitnil,omitempty" name:"ChargeStatus"`

	// <p>最近一次实例ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestInstanceId *string `json:"LatestInstanceId,omitnil,omitempty" name:"LatestInstanceId"`

	// <p>TensorBoard ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TensorBoardId *string `json:"TensorBoardId,omitnil,omitempty" name:"TensorBoardId"`

	// <p>备注</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// <p>失败原因</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailureReason *string `json:"FailureReason,omitnil,omitempty" name:"FailureReason"`

	// <p>更新时间</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>训练结束时间</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>计费金额信息，eg：2.00元/小时 (按量计费)</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillingInfo *string `json:"BillingInfo,omitnil,omitempty" name:"BillingInfo"`

	// <p>预付费专用资源组名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupName *string `json:"ResourceGroupName,omitnil,omitempty" name:"ResourceGroupName"`

	// <p>任务信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// <p>任务状态，eg：STARTING启动中、RUNNING运行中、STOPPING停止中、STOPPED已停止、FAILED异常、SUCCEED已完成</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>回调地址</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// <p>任务关联的代码仓库配置</p>
	CodeRepos []*CodeRepoConfig `json:"CodeRepos,omitnil,omitempty" name:"CodeRepos"`

	// <p>暴露网络配置</p>
	ExposeNetworkConfig *ExposeNetworkConfig `json:"ExposeNetworkConfig,omitnil,omitempty" name:"ExposeNetworkConfig"`

	// <p>操作者信息</p>
	OperatorInfo *OperatorInfo `json:"OperatorInfo,omitnil,omitempty" name:"OperatorInfo"`
}

type TrainingTaskSetItem struct {
	// <p>训练任务ID</p>
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>训练任务名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>框架名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrameworkName *string `json:"FrameworkName,omitnil,omitempty" name:"FrameworkName"`

	// <p>训练框架版本</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrameworkVersion *string `json:"FrameworkVersion,omitnil,omitempty" name:"FrameworkVersion"`

	// <p>框架运行环境</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrameworkEnvironment *string `json:"FrameworkEnvironment,omitnil,omitempty" name:"FrameworkEnvironment"`

	// <p>计费模式</p>
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// <p>计费状态，eg：BILLING计费中，ARREARS_STOP欠费停止，NOT_BILLING不在计费中</p>
	ChargeStatus *string `json:"ChargeStatus,omitnil,omitempty" name:"ChargeStatus"`

	// <p>预付费专用资源组</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// <p>资源配置</p>
	ResourceConfigInfos []*ResourceConfigInfo `json:"ResourceConfigInfos,omitnil,omitempty" name:"ResourceConfigInfos"`

	// <p>训练模式eg：PS_WORKER、DDP、MPI、HOROVOD</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrainingMode *string `json:"TrainingMode,omitnil,omitempty" name:"TrainingMode"`

	// <p>任务状态，eg：SUBMITTING提交中、PENDING排队中、<br>STARTING启动中、RUNNING运行中、STOPPING停止中、STOPPED已停止、FAILED异常、SUCCEED已完成</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>运行时长</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuntimeInSeconds *uint64 `json:"RuntimeInSeconds,omitnil,omitempty" name:"RuntimeInSeconds"`

	// <p>创建时间</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>训练开始时间</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>训练结束时间</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>训练输出</p>
	Output *CosPathInfo `json:"Output,omitnil,omitempty" name:"Output"`

	// <p>失败原因</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailureReason *string `json:"FailureReason,omitnil,omitempty" name:"FailureReason"`

	// <p>更新时间</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>计费金额信息，eg：2.00元/小时 (按量计费)</p>
	BillingInfo *string `json:"BillingInfo,omitnil,omitempty" name:"BillingInfo"`

	// <p>预付费专用资源组名称</p>
	ResourceGroupName *string `json:"ResourceGroupName,omitnil,omitempty" name:"ResourceGroupName"`

	// <p>自定义镜像信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageInfo *ImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`

	// <p>任务信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// <p>标签配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>回调地址</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// <p>任务subUin信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubUin *string `json:"SubUin,omitnil,omitempty" name:"SubUin"`

	// <p>任务创建者名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubUinName *string `json:"SubUinName,omitnil,omitempty" name:"SubUinName"`

	// <p>任务AppId</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>环境变量</p>
	Envs []*EnvVar `json:"Envs,omitnil,omitempty" name:"Envs"`

	// <p>操作者信息</p>
	LatestOperatorInfo *OperatorInfo `json:"LatestOperatorInfo,omitnil,omitempty" name:"LatestOperatorInfo"`
}

// Predefined struct for user
type UpdateDataSourceRequestParams struct {
	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`

	// 数据源ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 数据源名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数据源权限，取值有RW RO
	Permission *string `json:"Permission,omitnil,omitempty" name:"Permission"`

	// 数据源挂载配置
	MountConfigure *MountConfigureInfo `json:"MountConfigure,omitnil,omitempty" name:"MountConfigure"`
}

type UpdateDataSourceRequest struct {
	*tchttp.BaseRequest
	
	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`

	// 数据源ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 数据源名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数据源权限，取值有RW RO
	Permission *string `json:"Permission,omitnil,omitempty" name:"Permission"`

	// 数据源挂载配置
	MountConfigure *MountConfigureInfo `json:"MountConfigure,omitnil,omitempty" name:"MountConfigure"`
}

func (r *UpdateDataSourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDataSourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TiProjectId")
	delete(f, "Id")
	delete(f, "Name")
	delete(f, "Permission")
	delete(f, "MountConfigure")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateDataSourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDataSourceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateDataSourceResponse struct {
	*tchttp.BaseResponse
	Response *UpdateDataSourceResponseParams `json:"Response"`
}

func (r *UpdateDataSourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDataSourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateMountLimitRequestParams struct {
	// 数据源类型英文名
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 存储实例ID
	StorageId *string `json:"StorageId,omitnil,omitempty" name:"StorageId"`

	// 限制开关是否开启，只有开启时才有限制，默认关闭
	LimitMount *bool `json:"LimitMount,omitnil,omitempty" name:"LimitMount"`

	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`
}

type UpdateMountLimitRequest struct {
	*tchttp.BaseRequest
	
	// 数据源类型英文名
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 存储实例ID
	StorageId *string `json:"StorageId,omitnil,omitempty" name:"StorageId"`

	// 限制开关是否开启，只有开启时才有限制，默认关闭
	LimitMount *bool `json:"LimitMount,omitnil,omitempty" name:"LimitMount"`

	// <p>TI工作空间ID</p><p>仅用于“工作空间”白名单功能。如需使用，请联系TI管理员开通白名单。</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`
}

func (r *UpdateMountLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateMountLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "StorageId")
	delete(f, "LimitMount")
	delete(f, "TiProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateMountLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateMountLimitResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateMountLimitResponse struct {
	*tchttp.BaseResponse
	Response *UpdateMountLimitResponseParams `json:"Response"`
}

func (r *UpdateMountLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateMountLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateSubAccountLinuxUserInfoRequestParams struct {
	// 子账号信息列表
	SubAccountList []*SubAccountInfo `json:"SubAccountList,omitnil,omitempty" name:"SubAccountList"`
}

type UpdateSubAccountLinuxUserInfoRequest struct {
	*tchttp.BaseRequest
	
	// 子账号信息列表
	SubAccountList []*SubAccountInfo `json:"SubAccountList,omitnil,omitempty" name:"SubAccountList"`
}

func (r *UpdateSubAccountLinuxUserInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateSubAccountLinuxUserInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubAccountList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateSubAccountLinuxUserInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateSubAccountLinuxUserInfoResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateSubAccountLinuxUserInfoResponse struct {
	*tchttp.BaseResponse
	Response *UpdateSubAccountLinuxUserInfoResponseParams `json:"Response"`
}

func (r *UpdateSubAccountLinuxUserInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateSubAccountLinuxUserInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Usage struct {
	// 生成的token数目
	CompletionTokens *int64 `json:"CompletionTokens,omitnil,omitempty" name:"CompletionTokens"`

	// 输入的token数目
	PromptTokens *int64 `json:"PromptTokens,omitnil,omitempty" name:"PromptTokens"`

	// 总共token数目
	TotalTokens *int64 `json:"TotalTokens,omitnil,omitempty" name:"TotalTokens"`
}

type VolumeMount struct {
	// <p>cfs的配置信息</p>
	CFSConfig *CFSConfig `json:"CFSConfig,omitnil,omitempty" name:"CFSConfig"`

	// <p>挂载源类型，CFS、COS、PUBLIC_DATA_SOURCE，默认为CFS</p>
	VolumeSourceType *string `json:"VolumeSourceType,omitnil,omitempty" name:"VolumeSourceType"`

	// <p>自定义容器内挂载路径</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	MountPath *string `json:"MountPath,omitnil,omitempty" name:"MountPath"`

	// <p>挂载数据源时的配置信息</p>
	PublicDataSource *PublicDataSourceFS `json:"PublicDataSource,omitnil,omitempty" name:"PublicDataSource"`
}

type WeightEntry struct {
	// 服务id
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 流量权重值，ServiceGroup 下，不同服务版本根据权重比例分配流量
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`
}

type WorkloadStatus struct {
	// 当前实例数
	Replicas *int64 `json:"Replicas,omitnil,omitempty" name:"Replicas"`

	// 更新的实例数
	UpdatedReplicas *int64 `json:"UpdatedReplicas,omitnil,omitempty" name:"UpdatedReplicas"`

	// 就绪的实例数
	ReadyReplicas *int64 `json:"ReadyReplicas,omitnil,omitempty" name:"ReadyReplicas"`

	// 可用的实例数
	AvailableReplicas *int64 `json:"AvailableReplicas,omitnil,omitempty" name:"AvailableReplicas"`

	// 不可用的实例数
	UnavailableReplicas *int64 `json:"UnavailableReplicas,omitnil,omitempty" name:"UnavailableReplicas"`

	// Normal	正常运行中
	// Abnormal	服务异常，例如容器启动失败等
	// Waiting	服务等待中，例如容器下载镜像过程等
	// Stopped   已停止 
	// Pending 启动中
	// Stopping 停止中
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 工作负载的状况信息
	//
	// Deprecated: StatefulSetCondition is deprecated.
	StatefulSetCondition []*StatefulSetCondition `json:"StatefulSetCondition,omitnil,omitempty" name:"StatefulSetCondition"`

	// 工作负载历史的状况信息
	Conditions []*StatefulSetCondition `json:"Conditions,omitnil,omitempty" name:"Conditions"`

	// 状态异常时，展示原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`
}

type Workspace struct {
	// <p>项目ID</p>
	TiProjectId *string `json:"TiProjectId,omitnil,omitempty" name:"TiProjectId"`

	// <p>名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>创建时间</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>绑定的资源组信息</p>
	ResourceGroups []*ResourceGroupInWorkspace `json:"ResourceGroups,omitnil,omitempty" name:"ResourceGroups"`

	// <p>当前用户对此空间拥有的权限</p>
	ActionType []*string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// <p>工作空间状态</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}