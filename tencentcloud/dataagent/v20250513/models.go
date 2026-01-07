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

package v20250513

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type AddChunkRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 文件ID
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 新增chunk的后面一个ChunkID。如果是空就是插到队尾。插入位置的下一个 chunkId。如果插到最前面，传入原切片的第一个。
	BeforeChunkId *string `json:"BeforeChunkId,omitnil,omitempty" name:"BeforeChunkId"`

	// 显式指定的位置,实际的位置。从 0 开始计算。0 代表插到最前面，chunk total 代表插到最后面。
	InsertPos *int64 `json:"InsertPos,omitnil,omitempty" name:"InsertPos"`

	// chunk内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 新 Chunk 插入到目标 Chunk ​之后的位置。插入位置的上一个 chunkId
	AfterChunkId *string `json:"AfterChunkId,omitnil,omitempty" name:"AfterChunkId"`

	// 知识库id
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`
}

type AddChunkRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 文件ID
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 新增chunk的后面一个ChunkID。如果是空就是插到队尾。插入位置的下一个 chunkId。如果插到最前面，传入原切片的第一个。
	BeforeChunkId *string `json:"BeforeChunkId,omitnil,omitempty" name:"BeforeChunkId"`

	// 显式指定的位置,实际的位置。从 0 开始计算。0 代表插到最前面，chunk total 代表插到最后面。
	InsertPos *int64 `json:"InsertPos,omitnil,omitempty" name:"InsertPos"`

	// chunk内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 新 Chunk 插入到目标 Chunk ​之后的位置。插入位置的上一个 chunkId
	AfterChunkId *string `json:"AfterChunkId,omitnil,omitempty" name:"AfterChunkId"`

	// 知识库id
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`
}

func (r *AddChunkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddChunkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "FileId")
	delete(f, "BeforeChunkId")
	delete(f, "InsertPos")
	delete(f, "Content")
	delete(f, "AfterChunkId")
	delete(f, "KnowledgeBaseId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddChunkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddChunkResponseParams struct {
	// 新增的ChunkId
	ChunkId *string `json:"ChunkId,omitnil,omitempty" name:"ChunkId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddChunkResponse struct {
	*tchttp.BaseResponse
	Response *AddChunkResponseParams `json:"Response"`
}

func (r *AddChunkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddChunkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChatAIRequestParams struct {
	// 会话ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 问题内容
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 上下文
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 模型
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 是否深度思考
	DeepThinking *bool `json:"DeepThinking,omitnil,omitempty" name:"DeepThinking"`

	// 数据源id
	DataSourceIds []*string `json:"DataSourceIds,omitnil,omitempty" name:"DataSourceIds"`

	// agent类型
	AgentType *string `json:"AgentType,omitnil,omitempty" name:"AgentType"`

	// 需要重新生成答案的记录ID
	OldRecordId *string `json:"OldRecordId,omitnil,omitempty" name:"OldRecordId"`

	// 知识库id列表
	KnowledgeBaseIds []*string `json:"KnowledgeBaseIds,omitnil,omitempty" name:"KnowledgeBaseIds"`
}

type ChatAIRequest struct {
	*tchttp.BaseRequest
	
	// 会话ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 问题内容
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 上下文
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 模型
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 是否深度思考
	DeepThinking *bool `json:"DeepThinking,omitnil,omitempty" name:"DeepThinking"`

	// 数据源id
	DataSourceIds []*string `json:"DataSourceIds,omitnil,omitempty" name:"DataSourceIds"`

	// agent类型
	AgentType *string `json:"AgentType,omitnil,omitempty" name:"AgentType"`

	// 需要重新生成答案的记录ID
	OldRecordId *string `json:"OldRecordId,omitnil,omitempty" name:"OldRecordId"`

	// 知识库id列表
	KnowledgeBaseIds []*string `json:"KnowledgeBaseIds,omitnil,omitempty" name:"KnowledgeBaseIds"`
}

func (r *ChatAIRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChatAIRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionId")
	delete(f, "InstanceId")
	delete(f, "Question")
	delete(f, "Context")
	delete(f, "Model")
	delete(f, "DeepThinking")
	delete(f, "DataSourceIds")
	delete(f, "AgentType")
	delete(f, "OldRecordId")
	delete(f, "KnowledgeBaseIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChatAIRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChatAIResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。本接口为流式响应接口，当请求成功时，RequestId 会被放在 HTTP 响应的 Header "X-TC-RequestId" 中。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChatAIResponse struct {
	tchttp.BaseSSEResponse `json:"-"`
	Response *ChatAIResponseParams `json:"Response"`
}

func (r *ChatAIResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChatAIResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Chunk struct {
	// 切片ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 切片内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 切片的字数
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 切片概要
	Summary *string `json:"Summary,omitnil,omitempty" name:"Summary"`
}

type ColumnInfo struct {
	// 列名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 列类型：int, bigint, double, date, datetime, string，decimal
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 列名称描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 列索引
	Index *int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 原始字段名称
	OriginalName *string `json:"OriginalName,omitnil,omitempty" name:"OriginalName"`
}

type CosFileInfo struct {
	// 文件名称，包含后缀
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文件类型，"PDF", "DOC", "DOCX", "XLS", "XLSX", "PPT", "PPTX", "MD", "TXT", "PNG", "JPG", "JPEG", "CSV"
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 用户文件的cosurl
	UserCosUrl *string `json:"UserCosUrl,omitnil,omitempty" name:"UserCosUrl"`
}

// Predefined struct for user
type CreateDataAgentSessionRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type CreateDataAgentSessionRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *CreateDataAgentSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDataAgentSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDataAgentSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDataAgentSessionResponseParams struct {
	// 会话
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDataAgentSessionResponse struct {
	*tchttp.BaseResponse
	Response *CreateDataAgentSessionResponseParams `json:"Response"`
}

func (r *CreateDataAgentSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDataAgentSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteChunkRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 文件ID
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 切片ID
	ChunkIds []*string `json:"ChunkIds,omitnil,omitempty" name:"ChunkIds"`

	// 知识库id
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`
}

type DeleteChunkRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 文件ID
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 切片ID
	ChunkIds []*string `json:"ChunkIds,omitnil,omitempty" name:"ChunkIds"`

	// 知识库id
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`
}

func (r *DeleteChunkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteChunkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "FileId")
	delete(f, "ChunkIds")
	delete(f, "KnowledgeBaseId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteChunkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteChunkResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteChunkResponse struct {
	*tchttp.BaseResponse
	Response *DeleteChunkResponseParams `json:"Response"`
}

func (r *DeleteChunkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteChunkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDataAgentSessionRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 会话ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

type DeleteDataAgentSessionRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 会话ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

func (r *DeleteDataAgentSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDataAgentSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SessionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDataAgentSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDataAgentSessionResponseParams struct {
	// 删除的会话ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDataAgentSessionResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDataAgentSessionResponseParams `json:"Response"`
}

func (r *DeleteDataAgentSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDataAgentSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FileInfo struct {
	// 文件名称
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文件大小，字节
	FileSize *float64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// 文件类型,0=文本,1=表格，默认0
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 文件ID
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 状态，0：数据处理中  1：可用 -1：错误
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 操作者
	CreateUser *string `json:"CreateUser,omitnil,omitempty" name:"CreateUser"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 分片策略
	ChunkConfig *KnowledgeTaskConfig `json:"ChunkConfig,omitnil,omitempty" name:"ChunkConfig"`

	// 文件来源0=unknow,1=user_cos,2=local
	Source *int64 `json:"Source,omitnil,omitempty" name:"Source"`

	// 文件url
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// 是否官方示例，0=否，1=是
	IsShowCase *int64 `json:"IsShowCase,omitnil,omitempty" name:"IsShowCase"`

	// 文档摘要
	DocumentSummary *string `json:"DocumentSummary,omitnil,omitempty" name:"DocumentSummary"`
}

// Predefined struct for user
type GetJobsByKnowledgeBaseIdRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 知识库id
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`
}

type GetJobsByKnowledgeBaseIdRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 知识库id
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`
}

func (r *GetJobsByKnowledgeBaseIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetJobsByKnowledgeBaseIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "KnowledgeBaseId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetJobsByKnowledgeBaseIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetJobsByKnowledgeBaseIdResponseParams struct {
	// 任务列表详情
	Jobs []*UploadJob `json:"Jobs,omitnil,omitempty" name:"Jobs"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetJobsByKnowledgeBaseIdResponse struct {
	*tchttp.BaseResponse
	Response *GetJobsByKnowledgeBaseIdResponseParams `json:"Response"`
}

func (r *GetJobsByKnowledgeBaseIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetJobsByKnowledgeBaseIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetKnowledgeBaseFileListRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 默认 1 表示第一页，可以不填
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 默认 10 一页展示 10 条，可以不填
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 知识库id
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`
}

type GetKnowledgeBaseFileListRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 默认 1 表示第一页，可以不填
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 默认 10 一页展示 10 条，可以不填
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 知识库id
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`
}

func (r *GetKnowledgeBaseFileListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetKnowledgeBaseFileListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Page")
	delete(f, "PageSize")
	delete(f, "KnowledgeBaseId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetKnowledgeBaseFileListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetKnowledgeBaseFileListResponseParams struct {
	// 文件信息列表
	FileList []*FileInfo `json:"FileList,omitnil,omitempty" name:"FileList"`

	// 文件信息总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetKnowledgeBaseFileListResponse struct {
	*tchttp.BaseResponse
	Response *GetKnowledgeBaseFileListResponseParams `json:"Response"`
}

func (r *GetKnowledgeBaseFileListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetKnowledgeBaseFileListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetKnowledgeBaseListRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type GetKnowledgeBaseListRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *GetKnowledgeBaseListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetKnowledgeBaseListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetKnowledgeBaseListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetKnowledgeBaseListResponseParams struct {
	// 用户实例所有知识库列表
	KnowledgeBaseList []*KnowledgeBase `json:"KnowledgeBaseList,omitnil,omitempty" name:"KnowledgeBaseList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetKnowledgeBaseListResponse struct {
	*tchttp.BaseResponse
	Response *GetKnowledgeBaseListResponseParams `json:"Response"`
}

func (r *GetKnowledgeBaseListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetKnowledgeBaseListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetSessionDetailsRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 会话ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

type GetSessionDetailsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 会话ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

func (r *GetSessionDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSessionDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SessionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetSessionDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetSessionDetailsResponseParams struct {
	// 会话记录详情
	RecordList []*Record `json:"RecordList,omitnil,omitempty" name:"RecordList"`

	// 记录总数
	RecordCount *int64 `json:"RecordCount,omitnil,omitempty" name:"RecordCount"`

	// 当前在运行的record信息
	RunRecord *string `json:"RunRecord,omitnil,omitempty" name:"RunRecord"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetSessionDetailsResponse struct {
	*tchttp.BaseResponse
	Response *GetSessionDetailsResponseParams `json:"Response"`
}

func (r *GetSessionDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSessionDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetUploadJobDetailsRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 任务id
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type GetUploadJobDetailsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 任务id
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *GetUploadJobDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetUploadJobDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetUploadJobDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetUploadJobDetailsResponseParams struct {
	// 任务详情
	Job *UploadJob `json:"Job,omitnil,omitempty" name:"Job"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetUploadJobDetailsResponse struct {
	*tchttp.BaseResponse
	Response *GetUploadJobDetailsResponseParams `json:"Response"`
}

func (r *GetUploadJobDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetUploadJobDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KnowledgeBase struct {
	// 知识库id
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`

	// 知识库名称
	KnowledgeBaseName *string `json:"KnowledgeBaseName,omitnil,omitempty" name:"KnowledgeBaseName"`

	// 知识库描述
	KnowledgeBaseDesc *string `json:"KnowledgeBaseDesc,omitnil,omitempty" name:"KnowledgeBaseDesc"`

	// 创建者subuin
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 文件数量
	FileNum *int64 `json:"FileNum,omitnil,omitempty" name:"FileNum"`

	// 知识库关联的数据库列表，目前是只绑定一个数据源，数组预留拓展
	DatasourceIds []*string `json:"DatasourceIds,omitnil,omitempty" name:"DatasourceIds"`
}

type KnowledgeTaskConfig struct {
	// 切片类型  0:自定义切片，1：智能切片
	ChunkType *int64 `json:"ChunkType,omitnil,omitempty" name:"ChunkType"`

	// /智能切片：最小值 1000，默认 4800 自定义切片：正整数即可,默认值 1000
	MaxChunkSize *int64 `json:"MaxChunkSize,omitnil,omitempty" name:"MaxChunkSize"`

	//  切片分隔符,自定义切片使用：默认值为：["\n\n", "\n", "。", "！", "？", "，", ""]
	// 注意：此字段可能返回 null，表示取不到有效值。
	Delimiters []*string `json:"Delimiters,omitnil,omitempty" name:"Delimiters"`

	// 自定义切片使用:默认0 可重叠字符长度
	ChunkOverlap *int64 `json:"ChunkOverlap,omitnil,omitempty" name:"ChunkOverlap"`

	// 表格类文档解析
	Columns []*ColumnInfo `json:"Columns,omitnil,omitempty" name:"Columns"`

	// 带检索的索引列表
	Indexes []*int64 `json:"Indexes,omitnil,omitempty" name:"Indexes"`

	// 0：不生成文档摘要，1：生成文档概要。默认0，当取1时，GenParaSummary必须也为1
	GenDocSummary *int64 `json:"GenDocSummary,omitnil,omitempty" name:"GenDocSummary"`

	// 0：不生成段落摘要，1：生成段落概要。默认0
	GenParaSummary *int64 `json:"GenParaSummary,omitnil,omitempty" name:"GenParaSummary"`
}

// Predefined struct for user
type ModifyChunkRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 文件ID
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 切片ID
	ChunkId *string `json:"ChunkId,omitnil,omitempty" name:"ChunkId"`

	// 编辑后的文本
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 知识库id
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`
}

type ModifyChunkRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 文件ID
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 切片ID
	ChunkId *string `json:"ChunkId,omitnil,omitempty" name:"ChunkId"`

	// 编辑后的文本
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 知识库id
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`
}

func (r *ModifyChunkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyChunkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "FileId")
	delete(f, "ChunkId")
	delete(f, "Content")
	delete(f, "KnowledgeBaseId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyChunkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyChunkResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyChunkResponse struct {
	*tchttp.BaseResponse
	Response *ModifyChunkResponseParams `json:"Response"`
}

func (r *ModifyChunkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyChunkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyKnowledgeBaseRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 操作类型：Create，Update，Delete
	OperateType *string `json:"OperateType,omitnil,omitempty" name:"OperateType"`

	// 知识库id，update和delete时必填
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`

	// 知识库名称，create和update时必填。只允许字母、数字、汉字、下划线
	KnowledgeBaseName *string `json:"KnowledgeBaseName,omitnil,omitempty" name:"KnowledgeBaseName"`

	// 知识库描述，create和update时必填
	KnowledgeBaseDesc *string `json:"KnowledgeBaseDesc,omitnil,omitempty" name:"KnowledgeBaseDesc"`

	// 1仅自己使用，2指定用户，0全员
	UseScope *int64 `json:"UseScope,omitnil,omitempty" name:"UseScope"`

	// 可使用用户列表
	AuthorityUins []*string `json:"AuthorityUins,omitnil,omitempty" name:"AuthorityUins"`
}

type ModifyKnowledgeBaseRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 操作类型：Create，Update，Delete
	OperateType *string `json:"OperateType,omitnil,omitempty" name:"OperateType"`

	// 知识库id，update和delete时必填
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`

	// 知识库名称，create和update时必填。只允许字母、数字、汉字、下划线
	KnowledgeBaseName *string `json:"KnowledgeBaseName,omitnil,omitempty" name:"KnowledgeBaseName"`

	// 知识库描述，create和update时必填
	KnowledgeBaseDesc *string `json:"KnowledgeBaseDesc,omitnil,omitempty" name:"KnowledgeBaseDesc"`

	// 1仅自己使用，2指定用户，0全员
	UseScope *int64 `json:"UseScope,omitnil,omitempty" name:"UseScope"`

	// 可使用用户列表
	AuthorityUins []*string `json:"AuthorityUins,omitnil,omitempty" name:"AuthorityUins"`
}

func (r *ModifyKnowledgeBaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyKnowledgeBaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "OperateType")
	delete(f, "KnowledgeBaseId")
	delete(f, "KnowledgeBaseName")
	delete(f, "KnowledgeBaseDesc")
	delete(f, "UseScope")
	delete(f, "AuthorityUins")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyKnowledgeBaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyKnowledgeBaseResponseParams struct {
	// 知识库id
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyKnowledgeBaseResponse struct {
	*tchttp.BaseResponse
	Response *ModifyKnowledgeBaseResponseParams `json:"Response"`
}

func (r *ModifyKnowledgeBaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyKnowledgeBaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryChunkListRequestParams struct {
	// 表示第一页
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 默认一页展示 10 条
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 知识库id
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`
}

type QueryChunkListRequest struct {
	*tchttp.BaseRequest
	
	// 表示第一页
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 默认一页展示 10 条
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 知识库id
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`
}

func (r *QueryChunkListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryChunkListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Page")
	delete(f, "PageSize")
	delete(f, "KnowledgeBaseId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryChunkListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryChunkListResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 分片信息
	Chunks []*Chunk `json:"Chunks,omitnil,omitempty" name:"Chunks"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryChunkListResponse struct {
	*tchttp.BaseResponse
	Response *QueryChunkListResponseParams `json:"Response"`
}

func (r *QueryChunkListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryChunkListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Record struct {
	// 问题内容
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 回答内容
	Answer *string `json:"Answer,omitnil,omitempty" name:"Answer"`

	// 思考内容
	Think *string `json:"Think,omitnil,omitempty" name:"Think"`

	// 任务列表
	TaskList []*Task `json:"TaskList,omitnil,omitempty" name:"TaskList"`

	// 记录创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 记录更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 记录id
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 总结内容
	FinalSummary *string `json:"FinalSummary,omitnil,omitempty" name:"FinalSummary"`

	// 会话ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 1=赞，2=踩，0=无反馈
	Feedback *int64 `json:"Feedback,omitnil,omitempty" name:"Feedback"`

	// 数据库信息
	DbInfo *string `json:"DbInfo,omitnil,omitempty" name:"DbInfo"`

	// 错误信息
	ErrorContext *string `json:"ErrorContext,omitnil,omitempty" name:"ErrorContext"`

	// TaskList的string字符串
	TaskListStr *string `json:"TaskListStr,omitnil,omitempty" name:"TaskListStr"`

	// 知识库id列表
	KnowledgeBaseIds []*string `json:"KnowledgeBaseIds,omitnil,omitempty" name:"KnowledgeBaseIds"`

	// 上下文
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`
}

type StepExpand struct {
	// 标题
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// cellid数组
	CellIds []*string `json:"CellIds,omitnil,omitempty" name:"CellIds"`
}

type StepInfo struct {
	// 步骤id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 步骤名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 步骤状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 类型(text/expand)
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 总结
	Summary *string `json:"Summary,omitnil,omitempty" name:"Summary"`

	// 步骤扩展结构
	Expand *StepExpand `json:"Expand,omitnil,omitempty" name:"Expand"`

	// 描述
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`
}

// Predefined struct for user
type StopChatAIRequestParams struct {
	// 会话ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type StopChatAIRequest struct {
	*tchttp.BaseRequest
	
	// 会话ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *StopChatAIRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopChatAIRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionId")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopChatAIRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopChatAIResponseParams struct {
	// 会话
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopChatAIResponse struct {
	*tchttp.BaseResponse
	Response *StopChatAIResponseParams `json:"Response"`
}

func (r *StopChatAIResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopChatAIResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Task struct {
	// 任务ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 任务名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 任务状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务步骤列表
	StepInfoList []*StepInfo `json:"StepInfoList,omitnil,omitempty" name:"StepInfoList"`
}

// Predefined struct for user
type UploadAndCommitFileRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 上传文件列表
	CosFiles []*CosFileInfo `json:"CosFiles,omitnil,omitempty" name:"CosFiles"`

	// 知识库id
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`
}

type UploadAndCommitFileRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 上传文件列表
	CosFiles []*CosFileInfo `json:"CosFiles,omitnil,omitempty" name:"CosFiles"`

	// 知识库id
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`
}

func (r *UploadAndCommitFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadAndCommitFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "CosFiles")
	delete(f, "KnowledgeBaseId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadAndCommitFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadAndCommitFileResponseParams struct {
	// 上传任务
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UploadAndCommitFileResponse struct {
	*tchttp.BaseResponse
	Response *UploadAndCommitFileResponseParams `json:"Response"`
}

func (r *UploadAndCommitFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadAndCommitFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UploadJob struct {
	// id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 任务id
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 知识库id
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`

	// uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// subuin
	SubUin *string `json:"SubUin,omitnil,omitempty" name:"SubUin"`

	// Pending、FileUploading、
	// FileParsing、
	// Success、
	// Failed 
	// 	
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 任务更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 错误信息
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}