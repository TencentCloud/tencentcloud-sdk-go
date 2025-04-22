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

package v20240522

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AttributeItem struct {
	// 属性id
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttributeId *string `json:"AttributeId,omitnil,omitempty" name:"AttributeId"`

	// 属性标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttributeKey *string `json:"AttributeKey,omitnil,omitempty" name:"AttributeKey"`

	// 属性名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttributeName *string `json:"AttributeName,omitnil,omitempty" name:"AttributeName"`

	// 标签名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Labels []*AttributeLabelItem `json:"Labels,omitnil,omitempty" name:"Labels"`
}

type AttributeLabelItem struct {
	// 标签id
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelId *string `json:"LabelId,omitnil,omitempty" name:"LabelId"`

	// 标签名称，最大80个英文字符
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelName *string `json:"LabelName,omitnil,omitempty" name:"LabelName"`
}

type AttributeLabelReferItem struct {
	// 属性id
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttributeId *string `json:"AttributeId,omitnil,omitempty" name:"AttributeId"`

	// 标签id
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelIds []*string `json:"LabelIds,omitnil,omitempty" name:"LabelIds"`
}

// Predefined struct for user
type ChatCompletionsRequestParams struct {
	// 模型名称
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 聊天上下文信息。
	// 说明：
	// 1. 长度最多为 40，按对话时间从旧到新在数组中排列。
	// 2. Message.Role 可选值：system、user、assistant。
	// 其中，system 角色可选，如存在则必须位于列表的最开始。user 和 assistant 需交替出现，以 user 提问开始，user 提问结束，Content 不能为空。Role 的顺序示例：[system（可选） user assistant user assistant user ...]。
	Messages []*Message `json:"Messages,omitnil,omitempty" name:"Messages"`

	// 是否流式输出
	Stream *bool `json:"Stream,omitnil,omitempty" name:"Stream"`

	// 控制生成的随机性，较高的值会产生更多样化的输出。
	Temperature *float64 `json:"Temperature,omitnil,omitempty" name:"Temperature"`

	// 最大生成的token数量，默认为4096，最大可设置为16384
	MaxTokens *int64 `json:"MaxTokens,omitnil,omitempty" name:"MaxTokens"`
}

type ChatCompletionsRequest struct {
	*tchttp.BaseRequest
	
	// 模型名称
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 聊天上下文信息。
	// 说明：
	// 1. 长度最多为 40，按对话时间从旧到新在数组中排列。
	// 2. Message.Role 可选值：system、user、assistant。
	// 其中，system 角色可选，如存在则必须位于列表的最开始。user 和 assistant 需交替出现，以 user 提问开始，user 提问结束，Content 不能为空。Role 的顺序示例：[system（可选） user assistant user assistant user ...]。
	Messages []*Message `json:"Messages,omitnil,omitempty" name:"Messages"`

	// 是否流式输出
	Stream *bool `json:"Stream,omitnil,omitempty" name:"Stream"`

	// 控制生成的随机性，较高的值会产生更多样化的输出。
	Temperature *float64 `json:"Temperature,omitnil,omitempty" name:"Temperature"`

	// 最大生成的token数量，默认为4096，最大可设置为16384
	MaxTokens *int64 `json:"MaxTokens,omitnil,omitempty" name:"MaxTokens"`
}

func (r *ChatCompletionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChatCompletionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Model")
	delete(f, "Messages")
	delete(f, "Stream")
	delete(f, "Temperature")
	delete(f, "MaxTokens")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChatCompletionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChatCompletionsResponseParams struct {
	// Unix 时间戳，单位为秒。
	Created *int64 `json:"Created,omitnil,omitempty" name:"Created"`

	// Token 统计信息。
	// 按照总 Token 数量计费。
	Usage *ChatUsage `json:"Usage,omitnil,omitempty" name:"Usage"`

	// 本次请求的 RequestId。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 回复内容。
	Choices []*Choice `json:"Choices,omitnil,omitempty" name:"Choices"`

	// 模型名称。
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。本接口为流式响应接口，当请求成功时，RequestId 会被放在 HTTP 响应的 Header "X-TC-RequestId" 中。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChatCompletionsResponse struct {
	tchttp.BaseSSEResponse `json:"-"`
	Response *ChatCompletionsResponseParams `json:"Response"`
}

func (r *ChatCompletionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChatCompletionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChatUsage struct {
	// 输入token数
	PromptTokens *int64 `json:"PromptTokens,omitnil,omitempty" name:"PromptTokens"`

	// 输出token数
	CompletionTokens *int64 `json:"CompletionTokens,omitnil,omitempty" name:"CompletionTokens"`

	// 总token数
	TotalTokens *int64 `json:"TotalTokens,omitnil,omitempty" name:"TotalTokens"`
}

type Choice struct {
	// 结束标志位，可能为 stop、 content_filter。
	// stop 表示输出正常结束。
	// content_filter 只在开启流式输出审核时会出现，表示安全审核未通过。
	FinishReason *string `json:"FinishReason,omitnil,omitempty" name:"FinishReason"`

	// 增量返回值，流式调用时使用该字段。
	Delta *Delta `json:"Delta,omitnil,omitempty" name:"Delta"`

	// 返回值，非流式调用时使用该字段。
	Message *Message `json:"Message,omitnil,omitempty" name:"Message"`

	// 索引值，流式调用时使用该字段。
	Index *int64 `json:"Index,omitnil,omitempty" name:"Index"`
}

// Predefined struct for user
type CreateAttributeLabelRequestParams struct {
	// 知识库ID
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`

	// 属性标识，最大40个英文字符，如style
	AttributeKey *string `json:"AttributeKey,omitnil,omitempty" name:"AttributeKey"`

	// 属性名称，最大80个英文字符，如风格
	AttributeName *string `json:"AttributeName,omitnil,omitempty" name:"AttributeName"`

	// 属性标签信息
	Labels []*AttributeLabelItem `json:"Labels,omitnil,omitempty" name:"Labels"`
}

type CreateAttributeLabelRequest struct {
	*tchttp.BaseRequest
	
	// 知识库ID
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`

	// 属性标识，最大40个英文字符，如style
	AttributeKey *string `json:"AttributeKey,omitnil,omitempty" name:"AttributeKey"`

	// 属性名称，最大80个英文字符，如风格
	AttributeName *string `json:"AttributeName,omitnil,omitempty" name:"AttributeName"`

	// 属性标签信息
	Labels []*AttributeLabelItem `json:"Labels,omitnil,omitempty" name:"Labels"`
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
	delete(f, "KnowledgeBaseId")
	delete(f, "AttributeKey")
	delete(f, "AttributeName")
	delete(f, "Labels")
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
type CreateKnowledgeBaseRequestParams struct {

}

type CreateKnowledgeBaseRequest struct {
	*tchttp.BaseRequest
	
}

func (r *CreateKnowledgeBaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateKnowledgeBaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateKnowledgeBaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateKnowledgeBaseResponseParams struct {
	// 知识库ID
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateKnowledgeBaseResponse struct {
	*tchttp.BaseResponse
	Response *CreateKnowledgeBaseResponseParams `json:"Response"`
}

func (r *CreateKnowledgeBaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateKnowledgeBaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateQARequestParams struct {
	// 知识库ID
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`

	// 问题，最大1000个英文字符
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 答案，最大4000个英文字符
	Answer *string `json:"Answer,omitnil,omitempty" name:"Answer"`

	// 属性标签
	AttributeLabels []*AttributeLabelReferItem `json:"AttributeLabels,omitnil,omitempty" name:"AttributeLabels"`
}

type CreateQARequest struct {
	*tchttp.BaseRequest
	
	// 知识库ID
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`

	// 问题，最大1000个英文字符
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 答案，最大4000个英文字符
	Answer *string `json:"Answer,omitnil,omitempty" name:"Answer"`

	// 属性标签
	AttributeLabels []*AttributeLabelReferItem `json:"AttributeLabels,omitnil,omitempty" name:"AttributeLabels"`
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
	delete(f, "KnowledgeBaseId")
	delete(f, "Question")
	delete(f, "Answer")
	delete(f, "AttributeLabels")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateQARequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateQAResponseParams struct {
	// 问答对ID
	QaId *string `json:"QaId,omitnil,omitempty" name:"QaId"`

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
	// 默认为0
	TableResultType *string `json:"TableResultType,omitnil,omitempty" name:"TableResultType"`

	// 智能文档解析返回结果的格式
	// 0：只返回全文MD；
	// 1：只返回每一页的OCR原始Json；
	// 2：只返回每一页的MD，
	// 3：返回全文MD + 每一页的OCR原始Json；
	// 4：返回全文MD + 每一页的MD，
	// 默认值为0
	ResultType *string `json:"ResultType,omitnil,omitempty" name:"ResultType"`
}

// Predefined struct for user
type CreateReconstructDocumentFlowRequestParams struct {
	// 文件类型。
	// 
	// **支持的文件类型：**
	// - `PDF`、`DOC`、`DOCX`、`XLS`、`XLSX`、`PPT`、`PPTX`、`MD`、`TXT`、`PNG`、`JPG`、`JPEG`、`CSV`、`HTML`、`EPUB`、`BMP`、`GIF`、`WEBP`、`HEIC`、`EPS`、`ICNS`、`IM`、`PCX`、`PPM`、`TIFF`、`XBM`、`HEIF`、`JP2`
	// 
	// **支持的文件大小：**
	//  - `PDF` 最大300M
	//  - `DOCX`、`DOC`、`PPT`、`PPTX` 最大 200M
	//  - `TXT`、`MD` 最大10M
	//  - 其他 最大20M
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 文件的 URL 地址。
	// 文件存储于腾讯云的 URL 可保障更高的下载速度和稳定性，建议文件存储于腾讯云。 非腾讯云存储的 URL 速度和稳定性可能受一定影响。
	// 参考：[腾讯云COS文档](https://cloud.tencent.com/document/product/436/7749)
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// 文件的 Base64 值。
	// 支持的文件类型： PNG、JPG、JPEG、PDF、GIF、BMP、TIFF
	// 支持的文件大小：所下载文件经Base64编码后不超过 8M。文件下载时间不超过 3 秒。
	// 支持的图片像素：单边介于20-10000px之间。
	// 文件的 FileUrl、FileBase64 必须提供一个，如果都提供，只使用 FileUrl。
	FileBase64 *string `json:"FileBase64,omitnil,omitempty" name:"FileBase64"`

	// 文档的起始页码。
	// 当传入文件是PDF、PDF、PPT、PPTX、DOC类型时，用来指定识别的起始页码，识别的页码包含当前值。
	FileStartPageNumber *int64 `json:"FileStartPageNumber,omitnil,omitempty" name:"FileStartPageNumber"`

	// 文档的结束页码。
	// 当传入文件是PDF、PDF、PPT、PPTX、DOC类型时，用来指定识别的结束页码，识别的页码包含当前值。
	FileEndPageNumber *int64 `json:"FileEndPageNumber,omitnil,omitempty" name:"FileEndPageNumber"`

	// 创建文档解析任务配置信息。
	Config *CreateReconstructDocumentFlowConfig `json:"Config,omitnil,omitempty" name:"Config"`
}

type CreateReconstructDocumentFlowRequest struct {
	*tchttp.BaseRequest
	
	// 文件类型。
	// 
	// **支持的文件类型：**
	// - `PDF`、`DOC`、`DOCX`、`XLS`、`XLSX`、`PPT`、`PPTX`、`MD`、`TXT`、`PNG`、`JPG`、`JPEG`、`CSV`、`HTML`、`EPUB`、`BMP`、`GIF`、`WEBP`、`HEIC`、`EPS`、`ICNS`、`IM`、`PCX`、`PPM`、`TIFF`、`XBM`、`HEIF`、`JP2`
	// 
	// **支持的文件大小：**
	//  - `PDF` 最大300M
	//  - `DOCX`、`DOC`、`PPT`、`PPTX` 最大 200M
	//  - `TXT`、`MD` 最大10M
	//  - 其他 最大20M
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 文件的 URL 地址。
	// 文件存储于腾讯云的 URL 可保障更高的下载速度和稳定性，建议文件存储于腾讯云。 非腾讯云存储的 URL 速度和稳定性可能受一定影响。
	// 参考：[腾讯云COS文档](https://cloud.tencent.com/document/product/436/7749)
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// 文件的 Base64 值。
	// 支持的文件类型： PNG、JPG、JPEG、PDF、GIF、BMP、TIFF
	// 支持的文件大小：所下载文件经Base64编码后不超过 8M。文件下载时间不超过 3 秒。
	// 支持的图片像素：单边介于20-10000px之间。
	// 文件的 FileUrl、FileBase64 必须提供一个，如果都提供，只使用 FileUrl。
	FileBase64 *string `json:"FileBase64,omitnil,omitempty" name:"FileBase64"`

	// 文档的起始页码。
	// 当传入文件是PDF、PDF、PPT、PPTX、DOC类型时，用来指定识别的起始页码，识别的页码包含当前值。
	FileStartPageNumber *int64 `json:"FileStartPageNumber,omitnil,omitempty" name:"FileStartPageNumber"`

	// 文档的结束页码。
	// 当传入文件是PDF、PDF、PPT、PPTX、DOC类型时，用来指定识别的结束页码，识别的页码包含当前值。
	FileEndPageNumber *int64 `json:"FileEndPageNumber,omitnil,omitempty" name:"FileEndPageNumber"`

	// 创建文档解析任务配置信息。
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
	delete(f, "FileUrl")
	delete(f, "FileBase64")
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

type CreateSplitDocumentFlowConfig struct {
	// Markdown文件中表格返回的形式
	// 0，表格以MD形式返回
	// 1，表格以HTML形式返回
	// 默认为
	//
	// Deprecated: TableResultType is deprecated.
	TableResultType *string `json:"TableResultType,omitnil,omitempty" name:"TableResultType"`

	// 智能文档解析返回结果的格式
	// 0：只返回全文MD；
	// 1：只返回每一页的OCR原始Json；
	// 2：只返回每一页的MD，
	// 3：返回全文MD + 每一页的OCR原始Json；
	// 4：返回全文MD + 每一页的MD，
	// 默认值为3（返回全文MD + 每一页的OCR原始Json）
	// 
	//
	// Deprecated: ResultType is deprecated.
	ResultType *string `json:"ResultType,omitnil,omitempty" name:"ResultType"`

	// 是否开启mllm
	EnableMllm *bool `json:"EnableMllm,omitnil,omitempty" name:"EnableMllm"`

	// 最大分片长度
	MaxChunkSize *int64 `json:"MaxChunkSize,omitnil,omitempty" name:"MaxChunkSize"`
}

// Predefined struct for user
type CreateSplitDocumentFlowRequestParams struct {
	// 文件类型。
	// 
	// **支持的文件类型：**
	// - `PDF`、`DOC`、`DOCX`、`XLS`、`XLSX`、`PPT`、`PPTX`、`MD`、`TXT`、`PNG`、`JPG`、`JPEG`、`CSV`、`HTML`、`EPUB`
	// 
	// **支持的文件大小：**
	//  - `PDF` 最大300M
	//  - `DOCX`、`DOC`、`PPT`、`PPTX` 最大 200M
	//  - `TXT`、`MD` 最大10M
	//  - 其他 最大20M
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 文件的 URL 地址。
	// 文件存储于腾讯云的 URL 可保障更高的下载速度和稳定性，建议文件存储于腾讯云。 非腾讯云存储的 URL 速度和稳定性可能受一定影响。
	// 参考：[腾讯云COS文档](https://cloud.tencent.com/document/product/436/7749)
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// 文件名，可选。
	// **需带文件类型后缀**，当文件名无法从传入的`FileUrl`获取时需要通过该字段来明确。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文件的 Base64 值。
	// 支持的文件大小：所下载文件经Base64编码后不超过 8M。文件下载时间不超过 3 秒。
	// 支持的图片像素：单边介于20-10000px之间。
	// 文件的 FileUrl、FileBase64 必须提供一个，如果都提供，只使用 FileUrl。
	//
	// Deprecated: FileBase64 is deprecated.
	FileBase64 *string `json:"FileBase64,omitnil,omitempty" name:"FileBase64"`

	// 文档的起始页码。
	// 当传入文件是PDF、PDF、PPT、PPTX、DOC类型时，用来指定识别的起始页码，识别的页码包含当前值。
	FileStartPageNumber *int64 `json:"FileStartPageNumber,omitnil,omitempty" name:"FileStartPageNumber"`

	// 文档的结束页码。
	// 当传入文件是PDF、PDF、PPT、PPTX、DOC类型时，用来指定识别的结束页码，识别的页码包含当前值。
	FileEndPageNumber *int64 `json:"FileEndPageNumber,omitnil,omitempty" name:"FileEndPageNumber"`

	// 文档拆分任务的配置信息。
	Config *CreateSplitDocumentFlowConfig `json:"Config,omitnil,omitempty" name:"Config"`
}

type CreateSplitDocumentFlowRequest struct {
	*tchttp.BaseRequest
	
	// 文件类型。
	// 
	// **支持的文件类型：**
	// - `PDF`、`DOC`、`DOCX`、`XLS`、`XLSX`、`PPT`、`PPTX`、`MD`、`TXT`、`PNG`、`JPG`、`JPEG`、`CSV`、`HTML`、`EPUB`
	// 
	// **支持的文件大小：**
	//  - `PDF` 最大300M
	//  - `DOCX`、`DOC`、`PPT`、`PPTX` 最大 200M
	//  - `TXT`、`MD` 最大10M
	//  - 其他 最大20M
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 文件的 URL 地址。
	// 文件存储于腾讯云的 URL 可保障更高的下载速度和稳定性，建议文件存储于腾讯云。 非腾讯云存储的 URL 速度和稳定性可能受一定影响。
	// 参考：[腾讯云COS文档](https://cloud.tencent.com/document/product/436/7749)
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// 文件名，可选。
	// **需带文件类型后缀**，当文件名无法从传入的`FileUrl`获取时需要通过该字段来明确。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文件的 Base64 值。
	// 支持的文件大小：所下载文件经Base64编码后不超过 8M。文件下载时间不超过 3 秒。
	// 支持的图片像素：单边介于20-10000px之间。
	// 文件的 FileUrl、FileBase64 必须提供一个，如果都提供，只使用 FileUrl。
	FileBase64 *string `json:"FileBase64,omitnil,omitempty" name:"FileBase64"`

	// 文档的起始页码。
	// 当传入文件是PDF、PDF、PPT、PPTX、DOC类型时，用来指定识别的起始页码，识别的页码包含当前值。
	FileStartPageNumber *int64 `json:"FileStartPageNumber,omitnil,omitempty" name:"FileStartPageNumber"`

	// 文档的结束页码。
	// 当传入文件是PDF、PDF、PPT、PPTX、DOC类型时，用来指定识别的结束页码，识别的页码包含当前值。
	FileEndPageNumber *int64 `json:"FileEndPageNumber,omitnil,omitempty" name:"FileEndPageNumber"`

	// 文档拆分任务的配置信息。
	Config *CreateSplitDocumentFlowConfig `json:"Config,omitnil,omitempty" name:"Config"`
}

func (r *CreateSplitDocumentFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSplitDocumentFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileType")
	delete(f, "FileUrl")
	delete(f, "FileName")
	delete(f, "FileBase64")
	delete(f, "FileStartPageNumber")
	delete(f, "FileEndPageNumber")
	delete(f, "Config")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSplitDocumentFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSplitDocumentFlowResponseParams struct {
	// 拆分任务唯一ID。
	// 30天内可以通过`GetSplitDocumentResult`接口查询TaskId对应的拆分结果。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSplitDocumentFlowResponse struct {
	*tchttp.BaseResponse
	Response *CreateSplitDocumentFlowResponseParams `json:"Response"`
}

func (r *CreateSplitDocumentFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSplitDocumentFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAttributeLabelsRequestParams struct {
	// 知识库ID
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`

	// 属性ID
	AttributeIds []*string `json:"AttributeIds,omitnil,omitempty" name:"AttributeIds"`
}

type DeleteAttributeLabelsRequest struct {
	*tchttp.BaseRequest
	
	// 知识库ID
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`

	// 属性ID
	AttributeIds []*string `json:"AttributeIds,omitnil,omitempty" name:"AttributeIds"`
}

func (r *DeleteAttributeLabelsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAttributeLabelsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KnowledgeBaseId")
	delete(f, "AttributeIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAttributeLabelsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAttributeLabelsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAttributeLabelsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAttributeLabelsResponseParams `json:"Response"`
}

func (r *DeleteAttributeLabelsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAttributeLabelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDocsRequestParams struct {
	// 知识库ID
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`

	// 文档ID列表。支持批量删除，数量不超过100
	DocIds []*string `json:"DocIds,omitnil,omitempty" name:"DocIds"`
}

type DeleteDocsRequest struct {
	*tchttp.BaseRequest
	
	// 知识库ID
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`

	// 文档ID列表。支持批量删除，数量不超过100
	DocIds []*string `json:"DocIds,omitnil,omitempty" name:"DocIds"`
}

func (r *DeleteDocsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDocsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KnowledgeBaseId")
	delete(f, "DocIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDocsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDocsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDocsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDocsResponseParams `json:"Response"`
}

func (r *DeleteDocsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDocsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteKnowledgeBaseRequestParams struct {
	// 知识库ID
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`
}

type DeleteKnowledgeBaseRequest struct {
	*tchttp.BaseRequest
	
	// 知识库ID
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`
}

func (r *DeleteKnowledgeBaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteKnowledgeBaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KnowledgeBaseId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteKnowledgeBaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteKnowledgeBaseResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteKnowledgeBaseResponse struct {
	*tchttp.BaseResponse
	Response *DeleteKnowledgeBaseResponseParams `json:"Response"`
}

func (r *DeleteKnowledgeBaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteKnowledgeBaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteQAsRequestParams struct {
	// 知识库ID
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`

	// 问答对ID列表。支持批量删除，数量不超过100
	QaIds []*string `json:"QaIds,omitnil,omitempty" name:"QaIds"`
}

type DeleteQAsRequest struct {
	*tchttp.BaseRequest
	
	// 知识库ID
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`

	// 问答对ID列表。支持批量删除，数量不超过100
	QaIds []*string `json:"QaIds,omitnil,omitempty" name:"QaIds"`
}

func (r *DeleteQAsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteQAsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KnowledgeBaseId")
	delete(f, "QaIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteQAsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteQAsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteQAsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteQAsResponseParams `json:"Response"`
}

func (r *DeleteQAsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteQAsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Delta struct {
	// 角色名称。
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 内容详情。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 思维链内容。 ReasoningConent参数仅支持出参，且只有deepseek-r1模型会返回。
	ReasoningContent *string `json:"ReasoningContent,omitnil,omitempty" name:"ReasoningContent"`
}

// Predefined struct for user
type DescribeDocRequestParams struct {
	// 知识库ID
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`

	// 文档ID
	DocId *string `json:"DocId,omitnil,omitempty" name:"DocId"`
}

type DescribeDocRequest struct {
	*tchttp.BaseRequest
	
	// 知识库ID
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`

	// 文档ID
	DocId *string `json:"DocId,omitnil,omitempty" name:"DocId"`
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
	delete(f, "KnowledgeBaseId")
	delete(f, "DocId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDocRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDocResponseParams struct {
	// 文档ID
	DocId *string `json:"DocId,omitnil,omitempty" name:"DocId"`

	// 状态，
	// 
	// - Uploading  上传中  
	// - Auditing 审核中
	// - Parsing 解析中  
	// - ParseFailed 解析失败
	// - Indexing 创建索引中  
	// - IndexFailed 创建索引失败
	// - Success  发布成功
	// - Failed  失败
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 文件名
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 属性标签
	AttributeLabels []*AttributeLabelReferItem `json:"AttributeLabels,omitnil,omitempty" name:"AttributeLabels"`

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

type DocItem struct {
	// 文档id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DocId *string `json:"DocId,omitnil,omitempty" name:"DocId"`

	//  状态，
	// - Uploading  上传中  
	// - Auditing 审核中
	// - Parsing 解析中  
	// - ParseFailed 解析失败
	// - Indexing 创建索引中  
	// - IndexFailed 创建索引失败
	// - Success  发布成功
	// - Failed  失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 文件名
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 属性标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttributeLabels []*AttributeLabelReferItem `json:"AttributeLabels,omitnil,omitempty" name:"AttributeLabels"`
}

type DocumentUsage struct {
	// 文档拆分任务的页数
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 文档拆分任务消耗的总token数
	//
	// Deprecated: TotalToken is deprecated.
	TotalToken *int64 `json:"TotalToken,omitnil,omitempty" name:"TotalToken"`

	// 文档拆分任务消耗的总token数
	TotalTokens *int64 `json:"TotalTokens,omitnil,omitempty" name:"TotalTokens"`
}

type EmbeddingObject struct {
	// 向量
	Embedding []*float64 `json:"Embedding,omitnil,omitempty" name:"Embedding"`
}

// Predefined struct for user
type GetCharacterUsageRequestParams struct {

}

type GetCharacterUsageRequest struct {
	*tchttp.BaseRequest
	
}

func (r *GetCharacterUsageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCharacterUsageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetCharacterUsageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCharacterUsageResponseParams struct {
	// 已用字符数
	Used *uint64 `json:"Used,omitnil,omitempty" name:"Used"`

	// 可用字符数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetCharacterUsageResponse struct {
	*tchttp.BaseResponse
	Response *GetCharacterUsageResponseParams `json:"Response"`
}

func (r *GetCharacterUsageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCharacterUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetEmbeddingRequestParams struct {
	// 模型名称
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 需要 embedding 的文本, 单条文本最大长度500个字符, 总条数最大7条
	Inputs []*string `json:"Inputs,omitnil,omitempty" name:"Inputs"`
}

type GetEmbeddingRequest struct {
	*tchttp.BaseRequest
	
	// 模型名称
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 需要 embedding 的文本, 单条文本最大长度500个字符, 总条数最大7条
	Inputs []*string `json:"Inputs,omitnil,omitempty" name:"Inputs"`
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
type GetReconstructDocumentResultRequestParams struct {
	// 解析任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type GetReconstructDocumentResultRequest struct {
	*tchttp.BaseRequest
	
	// 解析任务ID
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
	// 任务状态。
	// - `Success`：执行完成
	// - `Processing`：执行中
	// -  `Pause`: 暂停
	// -  `Failed`：执行失败
	// -  `WaitExecute`：等待执行
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 解析结果的临时下载地址。文件类型为zip压缩包，下载链接有效期30分钟
	DocumentRecognizeResultUrl *string `json:"DocumentRecognizeResultUrl,omitnil,omitempty" name:"DocumentRecognizeResultUrl"`

	// 文档解析失败的页码
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
type GetSplitDocumentResultRequestParams struct {
	// 拆分任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type GetSplitDocumentResultRequest struct {
	*tchttp.BaseRequest
	
	// 拆分任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *GetSplitDocumentResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSplitDocumentResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetSplitDocumentResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetSplitDocumentResultResponseParams struct {
	// 任务状态。
	// 
	// - `Success`：执行完成
	// - `Processing`：执行中
	// - `Pause`: 暂停
	// - `Failed`：执行失败
	// - `WaitExecute`：等待执行
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 拆分结果的临时下载地址。
	// 文件类型为zip压缩包，下载链接有效期30分钟。
	// 压缩包内包含\*.md、\*.jsonl、\*mllm.json以及images文件夹。
	// 
	// > **jsonl** 结构说明：
	// - `page_content`:
	//  用于生成嵌入（embedding）库，以便于检索使用。该字段中的图片将使用占位符替换。
	// - `org_data`:
	//  表示与 page_content 对应的最小语义完整性块，用于问答模型的处理。
	// - `big_data`:
	//  表示与 page_content 对应的最大语义完整性块，也用于问答模型的处理。
	DocumentRecognizeResultUrl *string `json:"DocumentRecognizeResultUrl,omitnil,omitempty" name:"DocumentRecognizeResultUrl"`

	// 文档拆分失败的页码
	//
	// Deprecated: FailedPages is deprecated.
	FailedPages []*SplitDocumentFailedPage `json:"FailedPages,omitnil,omitempty" name:"FailedPages"`

	// 文档拆分任务的用量
	Usage *DocumentUsage `json:"Usage,omitnil,omitempty" name:"Usage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetSplitDocumentResultResponse struct {
	*tchttp.BaseResponse
	Response *GetSplitDocumentResultResponseParams `json:"Response"`
}

func (r *GetSplitDocumentResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSplitDocumentResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportQAsRequestParams struct {
	// 知识库ID
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`

	// 文件名
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文件的 Url 地址。文件存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议文件存储于腾讯云。 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	// 导入模板：https://cdn.xiaowei.qq.com/lke/assets//static/批量导入问答模板v6.xlsx
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// 文件类型，仅支持XLSX格式，请使用模板
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`
}

type ImportQAsRequest struct {
	*tchttp.BaseRequest
	
	// 知识库ID
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`

	// 文件名
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文件的 Url 地址。文件存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议文件存储于腾讯云。 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	// 导入模板：https://cdn.xiaowei.qq.com/lke/assets//static/批量导入问答模板v6.xlsx
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// 文件类型，仅支持XLSX格式，请使用模板
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`
}

func (r *ImportQAsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportQAsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KnowledgeBaseId")
	delete(f, "FileName")
	delete(f, "FileUrl")
	delete(f, "FileType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImportQAsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportQAsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ImportQAsResponse struct {
	*tchttp.BaseResponse
	Response *ImportQAsResponseParams `json:"Response"`
}

func (r *ImportQAsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportQAsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LabelItem struct {
	// 属性key
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 标签值
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

// Predefined struct for user
type ListAttributeLabelsRequestParams struct {
	// 知识库ID
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`

	// 页码，默认1
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数目，最大50，默认20
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type ListAttributeLabelsRequest struct {
	*tchttp.BaseRequest
	
	// 知识库ID
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`

	// 页码，默认1
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数目，最大50，默认20
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *ListAttributeLabelsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAttributeLabelsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KnowledgeBaseId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAttributeLabelsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAttributeLabelsResponseParams struct {
	// 属性总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 属性标签列表
	List []*AttributeItem `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAttributeLabelsResponse struct {
	*tchttp.BaseResponse
	Response *ListAttributeLabelsResponseParams `json:"Response"`
}

func (r *ListAttributeLabelsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAttributeLabelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDocsRequestParams struct {
	// 知识库ID
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`

	// 页码，默认1
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数目，最大50，默认20
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type ListDocsRequest struct {
	*tchttp.BaseRequest
	
	// 知识库ID
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`

	// 页码，默认1
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数目，最大50，默认20
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *ListDocsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDocsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KnowledgeBaseId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListDocsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDocsResponseParams struct {
	// 文档总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 文档信息
	List []*DocItem `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListDocsResponse struct {
	*tchttp.BaseResponse
	Response *ListDocsResponseParams `json:"Response"`
}

func (r *ListDocsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDocsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListQAsRequestParams struct {
	// 知识库ID
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`

	// 页码，默认1
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数目，最大50，默认20
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type ListQAsRequest struct {
	*tchttp.BaseRequest
	
	// 知识库ID
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`

	// 页码，默认1
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数目，最大50，默认20
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *ListQAsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListQAsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KnowledgeBaseId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListQAsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListQAsResponseParams struct {
	// 问答对总数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 问答对信息
	List []*QaItem `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListQAsResponse struct {
	*tchttp.BaseResponse
	Response *ListQAsResponseParams `json:"Response"`
}

func (r *ListQAsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListQAsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Message struct {
	// 角色
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 思维链内容。
	// ReasoningConent参数仅支持出参，且只有deepseek-r1模型会返回。
	ReasoningContent *string `json:"ReasoningContent,omitnil,omitempty" name:"ReasoningContent"`
}

// Predefined struct for user
type ModifyAttributeLabelRequestParams struct {
	// 知识库ID
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`

	// 属性ID
	AttributeId *string `json:"AttributeId,omitnil,omitempty" name:"AttributeId"`

	// 属性标识，最大40个英文字符，如style
	AttributeKey *string `json:"AttributeKey,omitnil,omitempty" name:"AttributeKey"`

	// 属性名称，最大80个英文字符，如风格
	AttributeName *string `json:"AttributeName,omitnil,omitempty" name:"AttributeName"`

	// 属性标签
	Labels []*AttributeLabelItem `json:"Labels,omitnil,omitempty" name:"Labels"`
}

type ModifyAttributeLabelRequest struct {
	*tchttp.BaseRequest
	
	// 知识库ID
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`

	// 属性ID
	AttributeId *string `json:"AttributeId,omitnil,omitempty" name:"AttributeId"`

	// 属性标识，最大40个英文字符，如style
	AttributeKey *string `json:"AttributeKey,omitnil,omitempty" name:"AttributeKey"`

	// 属性名称，最大80个英文字符，如风格
	AttributeName *string `json:"AttributeName,omitnil,omitempty" name:"AttributeName"`

	// 属性标签
	Labels []*AttributeLabelItem `json:"Labels,omitnil,omitempty" name:"Labels"`
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
	delete(f, "KnowledgeBaseId")
	delete(f, "AttributeId")
	delete(f, "AttributeKey")
	delete(f, "AttributeName")
	delete(f, "Labels")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAttributeLabelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAttributeLabelResponseParams struct {
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
type ModifyQARequestParams struct {
	// 知识库ID
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`

	// 问答对ID
	QaId *string `json:"QaId,omitnil,omitempty" name:"QaId"`

	// 问题，最大1000个英文字符
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 答案，最大4000个英文字符
	Answer *string `json:"Answer,omitnil,omitempty" name:"Answer"`

	// 属性标签
	AttributeLabels []*AttributeLabelReferItem `json:"AttributeLabels,omitnil,omitempty" name:"AttributeLabels"`
}

type ModifyQARequest struct {
	*tchttp.BaseRequest
	
	// 知识库ID
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`

	// 问答对ID
	QaId *string `json:"QaId,omitnil,omitempty" name:"QaId"`

	// 问题，最大1000个英文字符
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 答案，最大4000个英文字符
	Answer *string `json:"Answer,omitnil,omitempty" name:"Answer"`

	// 属性标签
	AttributeLabels []*AttributeLabelReferItem `json:"AttributeLabels,omitnil,omitempty" name:"AttributeLabels"`
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
	delete(f, "KnowledgeBaseId")
	delete(f, "QaId")
	delete(f, "Question")
	delete(f, "Answer")
	delete(f, "AttributeLabels")
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

type QaItem struct {
	// 问答id
	// 注意：此字段可能返回 null，表示取不到有效值。
	QaId *string `json:"QaId,omitnil,omitempty" name:"QaId"`

	// 问题
	// 注意：此字段可能返回 null，表示取不到有效值。
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 答案
	// 注意：此字段可能返回 null，表示取不到有效值。
	Answer *string `json:"Answer,omitnil,omitempty" name:"Answer"`

	// 属性标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttributeLabels []*AttributeLabelReferItem `json:"AttributeLabels,omitnil,omitempty" name:"AttributeLabels"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

// Predefined struct for user
type QueryRewriteRequestParams struct {
	// 需要改写的多轮历史会话，每轮历史对话需要包含user（问）和assistant（答）成对输入，由于模型字符限制，最多提供4轮对话。针对最后一轮对话进行改写
	Messages []*Message `json:"Messages,omitnil,omitempty" name:"Messages"`

	// 模型名称
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`
}

type QueryRewriteRequest struct {
	*tchttp.BaseRequest
	
	// 需要改写的多轮历史会话，每轮历史对话需要包含user（问）和assistant（答）成对输入，由于模型字符限制，最多提供4轮对话。针对最后一轮对话进行改写
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

type ReconstructDocumentFailedPage struct {
	// 失败页码	
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

type ReconstructDocumentSSEConfig struct {
	// Markdown文件中表格返回的形式
	// 0，表格以MD形式返回
	// 1，表格以HTML形式返回
	// 默认为0
	TableResultType *string `json:"TableResultType,omitnil,omitempty" name:"TableResultType"`

	// Markdown文件中图片返回的形式
	// 0:markdown中图片以链接形式返回
	// 1:markdown中图片只返回图片中提取的文本内容
	// 默认是0
	MarkdownImageResponseType *string `json:"MarkdownImageResponseType,omitnil,omitempty" name:"MarkdownImageResponseType"`

	// Markdown文件中是否包含页码信息
	ReturnPageFormat *bool `json:"ReturnPageFormat,omitnil,omitempty" name:"ReturnPageFormat"`

	// 自定义输出页码样式,{{p}}为页码占位符，开启ReturnPageFormat生效。未填默认样式:<page_num>page {{p}}</page_num>
	PageFormat *string `json:"PageFormat,omitnil,omitempty" name:"PageFormat"`
}

// Predefined struct for user
type ReconstructDocumentSSERequestParams struct {
	// 文件类型。
	// **支持的文件类型**：PDF、DOC、DOCX、PPT、PPTX、MD、TXT、XLS、XLSX、CSV、PNG、JPG、JPEG、BMP、GIF、WEBP、HEIC、EPS、ICNS、IM、PCX、PPM、TIFF、XBM、HEIF、JP2
	// **支持的文件大小**：
	// - PDF、DOC、DOCX、PPT、PPTX 支持100M
	// - MD、TXT、XLS、XLSX、CSV 支持10M
	// - 其他支持20M
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 文件的 URL 地址。
	// 文件存储于腾讯云的 URL 可保障更高的下载速度和稳定性，建议文件存储于腾讯云。 非腾讯云存储的 URL 速度和稳定性可能受一定影响。
	// 参考：[腾讯云COS文档](https://cloud.tencent.com/document/product/436/7749)
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// 文件的 Base64 值。
	// 支持的文件大小：所下载文件经Base64编码后不超过 8M。文件下载时间不超过 3 秒。
	// 支持的图片像素：单边介于20-10000px之间。
	// 文件的 FileUrl、FileBase64 必须提供一个，如果都提供，只使用 FileUrl。
	FileBase64 *string `json:"FileBase64,omitnil,omitempty" name:"FileBase64"`

	// 文档的起始页码。
	// 当传入文件是PDF、PDF、PPT、PPTX、DOC类型时，用来指定识别的起始页码，识别的页码包含当前值。
	FileStartPageNumber *int64 `json:"FileStartPageNumber,omitnil,omitempty" name:"FileStartPageNumber"`

	// 文档的结束页码。
	// 当传入文件是PDF、PDF、PPT、PPTX、DOC类型时，用来指定识别的结束页码，识别的页码包含当前值。
	FileEndPageNumber *int64 `json:"FileEndPageNumber,omitnil,omitempty" name:"FileEndPageNumber"`

	// 文档解析配置信息	
	Config *ReconstructDocumentSSEConfig `json:"Config,omitnil,omitempty" name:"Config"`
}

type ReconstructDocumentSSERequest struct {
	*tchttp.BaseRequest
	
	// 文件类型。
	// **支持的文件类型**：PDF、DOC、DOCX、PPT、PPTX、MD、TXT、XLS、XLSX、CSV、PNG、JPG、JPEG、BMP、GIF、WEBP、HEIC、EPS、ICNS、IM、PCX、PPM、TIFF、XBM、HEIF、JP2
	// **支持的文件大小**：
	// - PDF、DOC、DOCX、PPT、PPTX 支持100M
	// - MD、TXT、XLS、XLSX、CSV 支持10M
	// - 其他支持20M
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 文件的 URL 地址。
	// 文件存储于腾讯云的 URL 可保障更高的下载速度和稳定性，建议文件存储于腾讯云。 非腾讯云存储的 URL 速度和稳定性可能受一定影响。
	// 参考：[腾讯云COS文档](https://cloud.tencent.com/document/product/436/7749)
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// 文件的 Base64 值。
	// 支持的文件大小：所下载文件经Base64编码后不超过 8M。文件下载时间不超过 3 秒。
	// 支持的图片像素：单边介于20-10000px之间。
	// 文件的 FileUrl、FileBase64 必须提供一个，如果都提供，只使用 FileUrl。
	FileBase64 *string `json:"FileBase64,omitnil,omitempty" name:"FileBase64"`

	// 文档的起始页码。
	// 当传入文件是PDF、PDF、PPT、PPTX、DOC类型时，用来指定识别的起始页码，识别的页码包含当前值。
	FileStartPageNumber *int64 `json:"FileStartPageNumber,omitnil,omitempty" name:"FileStartPageNumber"`

	// 文档的结束页码。
	// 当传入文件是PDF、PDF、PPT、PPTX、DOC类型时，用来指定识别的结束页码，识别的页码包含当前值。
	FileEndPageNumber *int64 `json:"FileEndPageNumber,omitnil,omitempty" name:"FileEndPageNumber"`

	// 文档解析配置信息	
	Config *ReconstructDocumentSSEConfig `json:"Config,omitnil,omitempty" name:"Config"`
}

func (r *ReconstructDocumentSSERequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReconstructDocumentSSERequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileType")
	delete(f, "FileUrl")
	delete(f, "FileBase64")
	delete(f, "FileStartPageNumber")
	delete(f, "FileEndPageNumber")
	delete(f, "Config")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReconstructDocumentSSERequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReconstructDocumentSSEResponseParams struct {
	// 任务ID。本次请求的唯一标识
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 响应类型。1：返回进度信息，2：返回解析结果
	ResponseType *string `json:"ResponseType,omitnil,omitempty" name:"ResponseType"`

	// 进度。0~100
	Progress *string `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 进度信息。
	ProgressMessage *string `json:"ProgressMessage,omitnil,omitempty" name:"ProgressMessage"`

	// 文档解析结果的临时下载地址。
	// 文件类型为zip压缩包，下载链接有效期30分钟。
	// 压缩包内包含*.md、*.json以及images文件夹。
	DocumentRecognizeResultUrl *string `json:"DocumentRecognizeResultUrl,omitnil,omitempty" name:"DocumentRecognizeResultUrl"`

	// 文档解析失败的页码。
	FailedPages []*ReconstructDocumentFailedPage `json:"FailedPages,omitnil,omitempty" name:"FailedPages"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。本接口为流式响应接口，当请求成功时，RequestId 会被放在 HTTP 响应的 Header "X-TC-RequestId" 中。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReconstructDocumentSSEResponse struct {
	tchttp.BaseSSEResponse `json:"-"`
	Response *ReconstructDocumentSSEResponseParams `json:"Response"`
}

func (r *ReconstructDocumentSSEResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReconstructDocumentSSEResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetrievalRecord struct {
	// 检索结果的元数据
	Metadata *RetrievalRecordMetadata `json:"Metadata,omitnil,omitempty" name:"Metadata"`

	// 检索到的标题
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 检索到的内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type RetrievalRecordMetadata struct {
	// 结果的类型。
	// - `DOC`：文档
	// - `QA`：问答对
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 检索结果的来源。
	// - `SEMANTIC`：从语义检索中得到的结果
	// - `FULL_TEXT`：从全文检索中得到的结果
	ResultSource *string `json:"ResultSource,omitnil,omitempty" name:"ResultSource"`

	// 切片在文档中的页码，仅部分文档支持
	ChunkPageNumbers []*int64 `json:"ChunkPageNumbers,omitnil,omitempty" name:"ChunkPageNumbers"`
}

type RetrievalSetting struct {
	// 检索的类型，不填该参数则检索全部。
	// - `DOC`：文档
	// - `QA`：QA
	// 
	// 仅RetrieveKnowledge接口支持该参数
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 返回个数
	TopK *int64 `json:"TopK,omitnil,omitempty" name:"TopK"`

	// 分数过滤
	ScoreThreshold *float64 `json:"ScoreThreshold,omitnil,omitempty" name:"ScoreThreshold"`
}

// Predefined struct for user
type RetrieveKnowledgeRequestParams struct {
	// 知识库ID。
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`

	// 用于检索的文本。
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 检索方法，默认使用`HYBRID`混合检索。
	// - `SEMANTIC`：语义检索
	// - `FULL_TEXT`：全文检索
	// - `HYBRID`：混合检索
	RetrievalMethod *string `json:"RetrievalMethod,omitnil,omitempty" name:"RetrievalMethod"`

	// 检索设置。
	RetrievalSetting *RetrievalSetting `json:"RetrievalSetting,omitnil,omitempty" name:"RetrievalSetting"`

	// 标签过滤。
	AttributeLabels []*LabelItem `json:"AttributeLabels,omitnil,omitempty" name:"AttributeLabels"`
}

type RetrieveKnowledgeRequest struct {
	*tchttp.BaseRequest
	
	// 知识库ID。
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`

	// 用于检索的文本。
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 检索方法，默认使用`HYBRID`混合检索。
	// - `SEMANTIC`：语义检索
	// - `FULL_TEXT`：全文检索
	// - `HYBRID`：混合检索
	RetrievalMethod *string `json:"RetrievalMethod,omitnil,omitempty" name:"RetrievalMethod"`

	// 检索设置。
	RetrievalSetting *RetrievalSetting `json:"RetrievalSetting,omitnil,omitempty" name:"RetrievalSetting"`

	// 标签过滤。
	AttributeLabels []*LabelItem `json:"AttributeLabels,omitnil,omitempty" name:"AttributeLabels"`
}

func (r *RetrieveKnowledgeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RetrieveKnowledgeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KnowledgeBaseId")
	delete(f, "Query")
	delete(f, "RetrievalMethod")
	delete(f, "RetrievalSetting")
	delete(f, "AttributeLabels")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RetrieveKnowledgeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RetrieveKnowledgeResponseParams struct {
	// 检索结果
	Records []*RetrievalRecord `json:"Records,omitnil,omitempty" name:"Records"`

	// 检索结果数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RetrieveKnowledgeResponse struct {
	*tchttp.BaseResponse
	Response *RetrieveKnowledgeResponseParams `json:"Response"`
}

func (r *RetrieveKnowledgeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RetrieveKnowledgeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunRerankRequestParams struct {
	// 查询内容
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 文档列表，最多20个
	Docs []*string `json:"Docs,omitnil,omitempty" name:"Docs"`

	// 模型名称, 默认: lke-reranker-base
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`
}

type RunRerankRequest struct {
	*tchttp.BaseRequest
	
	// 查询内容
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 文档列表，最多20个
	Docs []*string `json:"Docs,omitnil,omitempty" name:"Docs"`

	// 模型名称, 默认: lke-reranker-base
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`
}

func (r *RunRerankRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunRerankRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Query")
	delete(f, "Docs")
	delete(f, "Model")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunRerankRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunRerankResponseParams struct {
	// 相关性, 数值越大越相关
	ScoreList []*float64 `json:"ScoreList,omitnil,omitempty" name:"ScoreList"`

	// 消耗量，仅返回TotalToken
	Usage *Usage `json:"Usage,omitnil,omitempty" name:"Usage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RunRerankResponse struct {
	*tchttp.BaseResponse
	Response *RunRerankResponseParams `json:"Response"`
}

func (r *RunRerankResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunRerankResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SegmentationConfig struct {
	// 最大分片长度
	MaxChunkSize *int64 `json:"MaxChunkSize,omitnil,omitempty" name:"MaxChunkSize"`
}

type SplitDocumentFailedPage struct {
	// 失败页码	
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

// Predefined struct for user
type UploadDocRealtimeRequestParams struct {
	// 知识库ID
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`

	// 文件名，可选。
	// **需带文件类型后缀**，当文件名无法从传入的`FileUrl`获取时需要通过该字段来明确。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文件类型。
	// **支持的文件类型：**
	// - `PDF`、`DOC`、`DOCX`、`XLS`、`XLSX`、`PPT`、`PPTX`、`MD`、`TXT`、`PNG`、`JPG`、`JPEG`、`CSV`、`HTML`、`EPUB`
	// 
	// **支持的文件大小：**
	//  - `PDF`、`DOCX`、`DOC`、`PPT`、`PPTX` 最大 200M 
	//  - `TXT`、`MD` 最大10M 
	//  - 其他 最大20M
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 文件的 URL 地址。
	// 文件存储于腾讯云的 URL 可保障更高的下载速度和稳定性，建议文件存储于腾讯云。 非腾讯云存储的 URL 速度和稳定性可能受一定影响。
	// 参考：[腾讯云COS文档](https://cloud.tencent.com/document/product/436/7749)
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// 过期时间的秒数，最长24小时，默认24小时
	ExpireTime *int64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`
}

type UploadDocRealtimeRequest struct {
	*tchttp.BaseRequest
	
	// 知识库ID
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`

	// 文件名，可选。
	// **需带文件类型后缀**，当文件名无法从传入的`FileUrl`获取时需要通过该字段来明确。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文件类型。
	// **支持的文件类型：**
	// - `PDF`、`DOC`、`DOCX`、`XLS`、`XLSX`、`PPT`、`PPTX`、`MD`、`TXT`、`PNG`、`JPG`、`JPEG`、`CSV`、`HTML`、`EPUB`
	// 
	// **支持的文件大小：**
	//  - `PDF`、`DOCX`、`DOC`、`PPT`、`PPTX` 最大 200M 
	//  - `TXT`、`MD` 最大10M 
	//  - 其他 最大20M
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 文件的 URL 地址。
	// 文件存储于腾讯云的 URL 可保障更高的下载速度和稳定性，建议文件存储于腾讯云。 非腾讯云存储的 URL 速度和稳定性可能受一定影响。
	// 参考：[腾讯云COS文档](https://cloud.tencent.com/document/product/436/7749)
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// 过期时间的秒数，最长24小时，默认24小时
	ExpireTime *int64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`
}

func (r *UploadDocRealtimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadDocRealtimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KnowledgeBaseId")
	delete(f, "FileName")
	delete(f, "FileType")
	delete(f, "FileUrl")
	delete(f, "ExpireTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadDocRealtimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadDocRealtimeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。本接口为流式响应接口，当请求成功时，RequestId 会被放在 HTTP 响应的 Header "X-TC-RequestId" 中。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UploadDocRealtimeResponse struct {
	tchttp.BaseSSEResponse `json:"-"`
	Response *UploadDocRealtimeResponseParams `json:"Response"`
}

func (r *UploadDocRealtimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadDocRealtimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadDocRequestParams struct {
	// 知识库ID
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`

	// 文件名。
	// **需带文件类型后缀**
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文件类型。
	// 
	// **支持的文件类型：**
	// - `PDF`、`DOC`、`DOCX`、`XLS`、`XLSX`、`PPT`、`PPTX`、`MD`、`TXT`、`PNG`、`JPG`、`JPEG`、`CSV`
	// 
	// **支持的文件大小：**
	//  - `PDF`、`DOCX`、`DOC`、`PPT`、`PPTX` 最大 200M
	//  - `TXT`、`MD` 最大10M
	//  - 其他 最大20M
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 文件的 URL 地址。
	// 文件存储于腾讯云的 URL 可保障更高的下载速度和稳定性，建议文件存储于腾讯云。 非腾讯云存储的 URL 速度和稳定性可能受一定影响。
	// 参考：[腾讯云COS文档](https://cloud.tencent.com/document/product/436/7749)
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// 属性标签引用
	//
	// Deprecated: AttributeLabel is deprecated.
	AttributeLabel []*AttributeLabelReferItem `json:"AttributeLabel,omitnil,omitempty" name:"AttributeLabel"`

	// 属性标签引用
	AttributeLabels []*AttributeLabelReferItem `json:"AttributeLabels,omitnil,omitempty" name:"AttributeLabels"`

	// 分段信息
	Config *SegmentationConfig `json:"Config,omitnil,omitempty" name:"Config"`
}

type UploadDocRequest struct {
	*tchttp.BaseRequest
	
	// 知识库ID
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`

	// 文件名。
	// **需带文件类型后缀**
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文件类型。
	// 
	// **支持的文件类型：**
	// - `PDF`、`DOC`、`DOCX`、`XLS`、`XLSX`、`PPT`、`PPTX`、`MD`、`TXT`、`PNG`、`JPG`、`JPEG`、`CSV`
	// 
	// **支持的文件大小：**
	//  - `PDF`、`DOCX`、`DOC`、`PPT`、`PPTX` 最大 200M
	//  - `TXT`、`MD` 最大10M
	//  - 其他 最大20M
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 文件的 URL 地址。
	// 文件存储于腾讯云的 URL 可保障更高的下载速度和稳定性，建议文件存储于腾讯云。 非腾讯云存储的 URL 速度和稳定性可能受一定影响。
	// 参考：[腾讯云COS文档](https://cloud.tencent.com/document/product/436/7749)
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// 属性标签引用
	AttributeLabel []*AttributeLabelReferItem `json:"AttributeLabel,omitnil,omitempty" name:"AttributeLabel"`

	// 属性标签引用
	AttributeLabels []*AttributeLabelReferItem `json:"AttributeLabels,omitnil,omitempty" name:"AttributeLabels"`

	// 分段信息
	Config *SegmentationConfig `json:"Config,omitnil,omitempty" name:"Config"`
}

func (r *UploadDocRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadDocRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KnowledgeBaseId")
	delete(f, "FileName")
	delete(f, "FileType")
	delete(f, "FileUrl")
	delete(f, "AttributeLabel")
	delete(f, "AttributeLabels")
	delete(f, "Config")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadDocRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadDocResponseParams struct {
	// 文档ID
	DocId *string `json:"DocId,omitnil,omitempty" name:"DocId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UploadDocResponse struct {
	*tchttp.BaseResponse
	Response *UploadDocResponseParams `json:"Response"`
}

func (r *UploadDocResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadDocResponse) FromJsonString(s string) error {
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