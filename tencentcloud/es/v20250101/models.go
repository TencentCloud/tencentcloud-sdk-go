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

package v20250101

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type ChatCompletionsRequestParams struct {
	// 会话内容，按对话时间从旧到新在数组中排列，长度受模型窗口大小限制
	Messages []*Message `json:"Messages,omitnil,omitempty" name:"Messages"`

	// 模型名称
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 是否以流式接口的形式返回数据，默认true
	Stream *bool `json:"Stream,omitnil,omitempty" name:"Stream"`

	// 取值区间为[0.0, 1.0], 非必要不建议使用, 不合理的取值会影响效果 
	TopP *float64 `json:"TopP,omitnil,omitempty" name:"TopP"`

	// 取值区间为[0.0, 2.0], 非必要不建议使用, 不合理的取值会影响效果 
	Temperature *float64 `json:"Temperature,omitnil,omitempty" name:"Temperature"`

	// 是否开启联网搜索。默认为 false。
	OnlineSearch *bool `json:"OnlineSearch,omitnil,omitempty" name:"OnlineSearch"`

	// 当 OnlineSearch 为 true 时，指定的搜索引擎，默认为 bing。
	OnlineSearchOptions *OnlineSearchOptions `json:"OnlineSearchOptions,omitnil,omitempty" name:"OnlineSearchOptions"`
}

type ChatCompletionsRequest struct {
	*tchttp.BaseRequest
	
	// 会话内容，按对话时间从旧到新在数组中排列，长度受模型窗口大小限制
	Messages []*Message `json:"Messages,omitnil,omitempty" name:"Messages"`

	// 模型名称
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 是否以流式接口的形式返回数据，默认true
	Stream *bool `json:"Stream,omitnil,omitempty" name:"Stream"`

	// 取值区间为[0.0, 1.0], 非必要不建议使用, 不合理的取值会影响效果 
	TopP *float64 `json:"TopP,omitnil,omitempty" name:"TopP"`

	// 取值区间为[0.0, 2.0], 非必要不建议使用, 不合理的取值会影响效果 
	Temperature *float64 `json:"Temperature,omitnil,omitempty" name:"Temperature"`

	// 是否开启联网搜索。默认为 false。
	OnlineSearch *bool `json:"OnlineSearch,omitnil,omitempty" name:"OnlineSearch"`

	// 当 OnlineSearch 为 true 时，指定的搜索引擎，默认为 bing。
	OnlineSearchOptions *OnlineSearchOptions `json:"OnlineSearchOptions,omitnil,omitempty" name:"OnlineSearchOptions"`
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
	delete(f, "Messages")
	delete(f, "ModelName")
	delete(f, "Stream")
	delete(f, "TopP")
	delete(f, "Temperature")
	delete(f, "OnlineSearch")
	delete(f, "OnlineSearchOptions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChatCompletionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChatCompletionsResponseParams struct {
	// 此次请求的id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 回复内容
	Choices []*Choice `json:"Choices,omitnil,omitempty" name:"Choices"`

	// token使用量
	Usage *TokenUsage `json:"Usage,omitnil,omitempty" name:"Usage"`

	// 联网搜索结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OnlineSearchContent []*WebContent `json:"OnlineSearchContent,omitnil,omitempty" name:"OnlineSearchContent"`

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

type Choice struct {
	// 返回的回复。
	Message *OutputMessage `json:"Message,omitnil,omitempty" name:"Message"`
}

type Chunk struct {
	// chunk索引。切片顺序 id。
	Index *uint64 `json:"Index,omitnil,omitempty" name:"Index"`

	// chunk内容。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type ChunkConfig struct {
	// 最大分片长度
	MaxChunkSize *uint64 `json:"MaxChunkSize,omitnil,omitempty" name:"MaxChunkSize"`

	// 分隔符列表
	Delimiters []*string `json:"Delimiters,omitnil,omitempty" name:"Delimiters"`
}

type ChunkConfigAsync struct {
	// 最大分片长度
	MaxChunkSize *int64 `json:"MaxChunkSize,omitnil,omitempty" name:"MaxChunkSize"`
}

type ChunkDocument struct {
	// 文件类型
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 文件的 base64值
	FileContent *string `json:"FileContent,omitnil,omitempty" name:"FileContent"`
}

// Predefined struct for user
type ChunkDocumentAsyncRequestParams struct {
	// 文件信息
	Document *Document `json:"Document,omitnil,omitempty" name:"Document"`

	// 模型名称
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 文件切片配置
	Config *ChunkConfigAsync `json:"Config,omitnil,omitempty" name:"Config"`
}

type ChunkDocumentAsyncRequest struct {
	*tchttp.BaseRequest
	
	// 文件信息
	Document *Document `json:"Document,omitnil,omitempty" name:"Document"`

	// 模型名称
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 文件切片配置
	Config *ChunkConfigAsync `json:"Config,omitnil,omitempty" name:"Config"`
}

func (r *ChunkDocumentAsyncRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChunkDocumentAsyncRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Document")
	delete(f, "ModelName")
	delete(f, "Config")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChunkDocumentAsyncRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChunkDocumentAsyncResponseParams struct {
	// 任务 ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChunkDocumentAsyncResponse struct {
	*tchttp.BaseResponse
	Response *ChunkDocumentAsyncResponseParams `json:"Response"`
}

func (r *ChunkDocumentAsyncResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChunkDocumentAsyncResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChunkDocumentRequestParams struct {
	// 文件切片文件信息
	Document *ChunkDocument `json:"Document,omitnil,omitempty" name:"Document"`

	// 模型名称
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 文件切片配置
	Config *ChunkConfig `json:"Config,omitnil,omitempty" name:"Config"`
}

type ChunkDocumentRequest struct {
	*tchttp.BaseRequest
	
	// 文件切片文件信息
	Document *ChunkDocument `json:"Document,omitnil,omitempty" name:"Document"`

	// 模型名称
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 文件切片配置
	Config *ChunkConfig `json:"Config,omitnil,omitempty" name:"Config"`
}

func (r *ChunkDocumentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChunkDocumentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Document")
	delete(f, "ModelName")
	delete(f, "Config")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChunkDocumentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChunkDocumentResponseParams struct {
	// 无
	Chunks []*Chunk `json:"Chunks,omitnil,omitempty" name:"Chunks"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChunkDocumentResponse struct {
	*tchttp.BaseResponse
	Response *ChunkDocumentResponseParams `json:"Response"`
}

func (r *ChunkDocumentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChunkDocumentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Document struct {
	// 文件类型。
	// 支持的文件类型：PDF、DOC、DOCX、PPT、PPTX、MD、TXT、XLS、XLSX、CSV、PNG、JPG、JPEG、BMP、GIF、WEBP、HEIC、EPS、ICNS、IM、PCX、PPM、TIFF、XBM、HEIF、JP2
	// 支持的文件大小：
	// - PDF、DOC、DOCX、PPT、PPTX 支持100M
	// - MD、TXT、XLS、XLSX、CSV 支持10M
	// - 其他支持20M
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 文件存储于腾讯云的 URL 可保障更高的下载速度和稳定性，使用腾讯云COS 文件地址。
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// 文件的 base64 值，携带 MineType前缀信息。编码后的后的文件不超过 10M。
	// 支持的文件大小：所下载文件经Base64编码后不超过 8M。文件下载时间不超过3秒。
	// 支持的图片像素：单边介于20-10000px之间。
	FileContent *string `json:"FileContent,omitnil,omitempty" name:"FileContent"`

	// 文件名称，当使用 base64上传的时候使用。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`
}

type DocumentChunkUsage struct {
	//  解析页面数量
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 消耗 token数量
	TotalTokens *int64 `json:"TotalTokens,omitnil,omitempty" name:"TotalTokens"`
}

type EmbeddingData struct {
	// embedding 内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Embedding []*float64 `json:"Embedding,omitnil,omitempty" name:"Embedding"`

	// 索引序号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index *int64 `json:"Index,omitnil,omitempty" name:"Index"`
}

// Predefined struct for user
type GetDocumentChunkResultRequestParams struct {
	//  任务 ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type GetDocumentChunkResultRequest struct {
	*tchttp.BaseRequest
	
	//  任务 ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *GetDocumentChunkResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDocumentChunkResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDocumentChunkResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDocumentChunkResultResponseParams struct {
	// 任务状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 切片结果
	DocumentChunkResultUrl *string `json:"DocumentChunkResultUrl,omitnil,omitempty" name:"DocumentChunkResultUrl"`

	// 用量
	Usage *DocumentChunkUsage `json:"Usage,omitnil,omitempty" name:"Usage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetDocumentChunkResultResponse struct {
	*tchttp.BaseResponse
	Response *GetDocumentChunkResultResponseParams `json:"Response"`
}

func (r *GetDocumentChunkResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDocumentChunkResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDocumentParseResultRequestParams struct {
	// 任务 Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type GetDocumentParseResultRequest struct {
	*tchttp.BaseRequest
	
	// 任务 Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *GetDocumentParseResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDocumentParseResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDocumentParseResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDocumentParseResultResponseParams struct {
	// 任务状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 结果文件
	DocumentParseResultUrl *string `json:"DocumentParseResultUrl,omitnil,omitempty" name:"DocumentParseResultUrl"`

	// 失败的页码
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedPages []*int64 `json:"FailedPages,omitnil,omitempty" name:"FailedPages"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetDocumentParseResultResponse struct {
	*tchttp.BaseResponse
	Response *GetDocumentParseResultResponseParams `json:"Response"`
}

func (r *GetDocumentParseResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDocumentParseResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTextEmbeddingRequestParams struct {
	// 模型名称
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 需进行向量化的文本集
	Texts []*string `json:"Texts,omitnil,omitempty" name:"Texts"`
}

type GetTextEmbeddingRequest struct {
	*tchttp.BaseRequest
	
	// 模型名称
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 需进行向量化的文本集
	Texts []*string `json:"Texts,omitnil,omitempty" name:"Texts"`
}

func (r *GetTextEmbeddingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTextEmbeddingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelName")
	delete(f, "Texts")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTextEmbeddingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTextEmbeddingResponseParams struct {
	// 结果集
	Data []*EmbeddingData `json:"Data,omitnil,omitempty" name:"Data"`

	// 消耗token数量
	Usage *Usage `json:"Usage,omitnil,omitempty" name:"Usage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetTextEmbeddingResponse struct {
	*tchttp.BaseResponse
	Response *GetTextEmbeddingResponseParams `json:"Response"`
}

func (r *GetTextEmbeddingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTextEmbeddingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Message struct {
	// 角色, ‘system', ‘user'，'assistant'或者'tool', 在message中, 除了system，其他必须是user与assistant交替(一问一答) 
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 具体文本内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type OnlineSearchOptions struct {
	// 搜索引擎。支持 bing 和 sogou。
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`
}

type OutputMessage struct {
	// 角色
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 文本内容	
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 推理内容	
	ReasoningContent *string `json:"ReasoningContent,omitnil,omitempty" name:"ReasoningContent"`
}

type ParseDocument struct {
	// 文件类型。
	// 支持的文件类型：PDF、DOC、DOCX、PPT、PPTX、MD、TXT、XLS、XLSX、CSV、PNG、JPG、JPEG、BMP、GIF、WEBP、HEIC、EPS、ICNS、IM、PCX、PPM、TIFF、XBM、HEIF、JP2
	// 支持的文件大小：
	// - PDF、DOC、DOCX、PPT、PPTX 支持100M
	// - MD、TXT、XLS、XLSX、CSV 支持10M
	// - 其他支持20M
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 文件存储于腾讯云的 URL 可保障更高的下载速度和稳定性，使用腾讯云COS 文件地址。
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// 文件的 base64 值，携带 MineType前缀信息。编码后的后的文件不超过 10M。
	// 支持的文件大小：所下载文件经Base64编码后不超过 8M。文件下载时间不超过3秒。
	// 支持的图片像素：单边介于20-10000px之间。
	FileContent *string `json:"FileContent,omitnil,omitempty" name:"FileContent"`
}

// Predefined struct for user
type ParseDocumentAsyncRequestParams struct {
	// 文件信息
	Document *Document `json:"Document,omitnil,omitempty" name:"Document"`

	// 模型名称
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`
}

type ParseDocumentAsyncRequest struct {
	*tchttp.BaseRequest
	
	// 文件信息
	Document *Document `json:"Document,omitnil,omitempty" name:"Document"`

	// 模型名称
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`
}

func (r *ParseDocumentAsyncRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ParseDocumentAsyncRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Document")
	delete(f, "ModelName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ParseDocumentAsyncRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ParseDocumentAsyncResponseParams struct {
	// 任务 id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ParseDocumentAsyncResponse struct {
	*tchttp.BaseResponse
	Response *ParseDocumentAsyncResponseParams `json:"Response"`
}

func (r *ParseDocumentAsyncResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ParseDocumentAsyncResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ParseDocumentRequestParams struct {
	// 文件信息
	Document *ParseDocument `json:"Document,omitnil,omitempty" name:"Document"`

	// 模型名称
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`
}

type ParseDocumentRequest struct {
	*tchttp.BaseRequest
	
	// 文件信息
	Document *ParseDocument `json:"Document,omitnil,omitempty" name:"Document"`

	// 模型名称
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`
}

func (r *ParseDocumentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ParseDocumentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Document")
	delete(f, "ModelName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ParseDocumentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ParseDocumentResponseParams struct {
	// 进度
	Progress *string `json:"Progress,omitnil,omitempty" name:"Progress"`

	//  解析文件结果
	DocumentParseResultUrl *string `json:"DocumentParseResultUrl,omitnil,omitempty" name:"DocumentParseResultUrl"`

	// 失败页码
	FailedPages []*int64 `json:"FailedPages,omitnil,omitempty" name:"FailedPages"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。本接口为流式响应接口，当请求成功时，RequestId 会被放在 HTTP 响应的 Header "X-TC-RequestId" 中。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ParseDocumentResponse struct {
	tchttp.BaseSSEResponse `json:"-"`
	Response *ParseDocumentResponseParams `json:"Response"`
}

func (r *ParseDocumentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ParseDocumentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RerankResult struct {
	// 对应的doc在输入候选doc数组中的位置索引值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index *int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 相似度分数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelevanceScore *float64 `json:"RelevanceScore,omitnil,omitempty" name:"RelevanceScore"`

	// doc原文内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Document *string `json:"Document,omitnil,omitempty" name:"Document"`
}

// Predefined struct for user
type RunRerankRequestParams struct {
	// 模型名称
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 查询文本
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 待排序的候选doc列表
	Documents []*string `json:"Documents,omitnil,omitempty" name:"Documents"`

	// 排序返回的top文档数量, 如果没有指定则返回全部候选doc，如果指定的top_n值大于输入的候选doc数量，返回全部doc
	TopN *int64 `json:"TopN,omitnil,omitempty" name:"TopN"`

	// 返回的排序结果列表里面是否返回每一条document原文，默认值False
	ReturnDocuments *bool `json:"ReturnDocuments,omitnil,omitempty" name:"ReturnDocuments"`
}

type RunRerankRequest struct {
	*tchttp.BaseRequest
	
	// 模型名称
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 查询文本
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 待排序的候选doc列表
	Documents []*string `json:"Documents,omitnil,omitempty" name:"Documents"`

	// 排序返回的top文档数量, 如果没有指定则返回全部候选doc，如果指定的top_n值大于输入的候选doc数量，返回全部doc
	TopN *int64 `json:"TopN,omitnil,omitempty" name:"TopN"`

	// 返回的排序结果列表里面是否返回每一条document原文，默认值False
	ReturnDocuments *bool `json:"ReturnDocuments,omitnil,omitempty" name:"ReturnDocuments"`
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
	delete(f, "ModelName")
	delete(f, "Query")
	delete(f, "Documents")
	delete(f, "TopN")
	delete(f, "ReturnDocuments")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunRerankRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunRerankResponseParams struct {
	// 输出结果集
	Data []*RerankResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 消耗token数量
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

type TokenUsage struct {
	// 表示prompt的tokens数，多次返回中保持不变
	PromptTokens *uint64 `json:"PromptTokens,omitnil,omitempty" name:"PromptTokens"`

	// 回答的token总数，在流式返回中，表示到目前为止所有completion的tokens总数，多次返回中持续累加        
	CompletionTokens *uint64 `json:"CompletionTokens,omitnil,omitempty" name:"CompletionTokens"`

	// 表示prompt_tokens和completion_tokens之和 
	TotalTokens *uint64 `json:"TotalTokens,omitnil,omitempty" name:"TotalTokens"`
}

type Usage struct {
	// tokens总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalTokens *int64 `json:"TotalTokens,omitnil,omitempty" name:"TotalTokens"`
}

type WebContent struct {
	// 搜素问题	
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 标题
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 链接
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 时间
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 网页内容	
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 切片索引
	ChunkIndex *string `json:"ChunkIndex,omitnil,omitempty" name:"ChunkIndex"`

	// 分数
	Score *string `json:"Score,omitnil,omitempty" name:"Score"`
}

type WebPage struct {
	// 标题
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// url
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 网页摘要
	// 注意：此字段可能返回 null，表示取不到有效值。
	Summary *string `json:"Summary,omitnil,omitempty" name:"Summary"`

	// 网页收录时间。可能为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// Markdown 格式的网页正文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

// Predefined struct for user
type WebSearchRequestParams struct {
	// 查询
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 搜索的网页数量，默认20
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 指定域名，gov.cn 可匹配 *.gov.cn的域名。
	Site *string `json:"Site,omitnil,omitempty" name:"Site"`

	// 是否获取返回网页全文，默认 false。
	FetchContent *bool `json:"FetchContent,omitnil,omitempty" name:"FetchContent"`

	// 域名白名单，在不指定 Site 时，只保存匹配白名单域名的网页。
	WhiteSites []*string `json:"WhiteSites,omitnil,omitempty" name:"WhiteSites"`

	// 域名黑名单，在不指定 Site 和白名单时，过滤黑名单中的域名。
	BlackSites []*string `json:"BlackSites,omitnil,omitempty" name:"BlackSites"`

	// 秒级时间戳，搜索网页的开始时间，默认不限制开始时间。
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 秒级时间戳，搜索网页的结束时间，默认为现在。
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 指定搜索引擎，可选混合搜索 mixed，或 bing, baidu, sogou, 默认为 sogou
	SearchEngine *string `json:"SearchEngine,omitnil,omitempty" name:"SearchEngine"`
}

type WebSearchRequest struct {
	*tchttp.BaseRequest
	
	// 查询
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 搜索的网页数量，默认20
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 指定域名，gov.cn 可匹配 *.gov.cn的域名。
	Site *string `json:"Site,omitnil,omitempty" name:"Site"`

	// 是否获取返回网页全文，默认 false。
	FetchContent *bool `json:"FetchContent,omitnil,omitempty" name:"FetchContent"`

	// 域名白名单，在不指定 Site 时，只保存匹配白名单域名的网页。
	WhiteSites []*string `json:"WhiteSites,omitnil,omitempty" name:"WhiteSites"`

	// 域名黑名单，在不指定 Site 和白名单时，过滤黑名单中的域名。
	BlackSites []*string `json:"BlackSites,omitnil,omitempty" name:"BlackSites"`

	// 秒级时间戳，搜索网页的开始时间，默认不限制开始时间。
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 秒级时间戳，搜索网页的结束时间，默认为现在。
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 指定搜索引擎，可选混合搜索 mixed，或 bing, baidu, sogou, 默认为 sogou
	SearchEngine *string `json:"SearchEngine,omitnil,omitempty" name:"SearchEngine"`
}

func (r *WebSearchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *WebSearchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Query")
	delete(f, "Count")
	delete(f, "Site")
	delete(f, "FetchContent")
	delete(f, "WhiteSites")
	delete(f, "BlackSites")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SearchEngine")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "WebSearchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type WebSearchResponseParams struct {
	// 查询
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 响应状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 执行搜索的引擎
	SearchEngine *string `json:"SearchEngine,omitnil,omitempty" name:"SearchEngine"`

	// 搜索结果
	Results []*WebPage `json:"Results,omitnil,omitempty" name:"Results"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type WebSearchResponse struct {
	*tchttp.BaseResponse
	Response *WebSearchResponseParams `json:"Response"`
}

func (r *WebSearchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *WebSearchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}