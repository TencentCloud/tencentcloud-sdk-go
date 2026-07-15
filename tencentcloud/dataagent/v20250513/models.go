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
type AddSceneRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 场景
	Scene *Scene `json:"Scene,omitnil,omitempty" name:"Scene"`

	// 1仅自己使用，2指定用户，0全员
	UseScope *int64 `json:"UseScope,omitnil,omitempty" name:"UseScope"`

	// 可使用用户列表
	AuthorityUins []*string `json:"AuthorityUins,omitnil,omitempty" name:"AuthorityUins"`
}

type AddSceneRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 场景
	Scene *Scene `json:"Scene,omitnil,omitempty" name:"Scene"`

	// 1仅自己使用，2指定用户，0全员
	UseScope *int64 `json:"UseScope,omitnil,omitempty" name:"UseScope"`

	// 可使用用户列表
	AuthorityUins []*string `json:"AuthorityUins,omitnil,omitempty" name:"AuthorityUins"`
}

func (r *AddSceneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddSceneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Scene")
	delete(f, "UseScope")
	delete(f, "AuthorityUins")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddSceneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddSceneResponseParams struct {
	// 场景id
	SceneId *string `json:"SceneId,omitnil,omitempty" name:"SceneId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddSceneResponse struct {
	*tchttp.BaseResponse
	Response *AddSceneResponseParams `json:"Response"`
}

func (r *AddSceneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddSceneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AppendDocument struct {
	// <p>文件名称</p>
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// <p>文件id</p>
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <p>文件url</p>
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// <p>文件大小</p>
	FileSize *float64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`
}

// Predefined struct for user
type AppendKnowledgeTaskRequestParams struct {
	// <p>实例id</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>知识库id</p>
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`

	// <p>文件id</p>
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <p>追加的文档列表</p>
	Documents []*AppendDocument `json:"Documents,omitnil,omitempty" name:"Documents"`
}

type AppendKnowledgeTaskRequest struct {
	*tchttp.BaseRequest
	
	// <p>实例id</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>知识库id</p>
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`

	// <p>文件id</p>
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <p>追加的文档列表</p>
	Documents []*AppendDocument `json:"Documents,omitnil,omitempty" name:"Documents"`
}

func (r *AppendKnowledgeTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AppendKnowledgeTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "KnowledgeBaseId")
	delete(f, "FileId")
	delete(f, "Documents")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AppendKnowledgeTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AppendKnowledgeTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AppendKnowledgeTaskResponse struct {
	*tchttp.BaseResponse
	Response *AppendKnowledgeTaskResponseParams `json:"Response"`
}

func (r *AppendKnowledgeTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AppendKnowledgeTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChatAIRequestParams struct {
	// <p>会话ID</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>问题内容</p>
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// <p>上下文</p>
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// <p>模型</p>
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// <p>是否深度思考</p>
	DeepThinking *bool `json:"DeepThinking,omitnil,omitempty" name:"DeepThinking"`

	// <p>数据源id</p>
	DataSourceIds []*string `json:"DataSourceIds,omitnil,omitempty" name:"DataSourceIds"`

	// <p>agent类型</p>
	AgentType *string `json:"AgentType,omitnil,omitempty" name:"AgentType"`

	// <p>需要重新生成答案的记录ID</p>
	OldRecordId *string `json:"OldRecordId,omitnil,omitempty" name:"OldRecordId"`

	// <p>知识库id列表</p>
	KnowledgeBaseIds []*string `json:"KnowledgeBaseIds,omitnil,omitempty" name:"KnowledgeBaseIds"`

	// <p>版本信息</p>
	ArchVersion *string `json:"ArchVersion,omitnil,omitempty" name:"ArchVersion"`
}

type ChatAIRequest struct {
	*tchttp.BaseRequest
	
	// <p>会话ID</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>问题内容</p>
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// <p>上下文</p>
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// <p>模型</p>
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// <p>是否深度思考</p>
	DeepThinking *bool `json:"DeepThinking,omitnil,omitempty" name:"DeepThinking"`

	// <p>数据源id</p>
	DataSourceIds []*string `json:"DataSourceIds,omitnil,omitempty" name:"DataSourceIds"`

	// <p>agent类型</p>
	AgentType *string `json:"AgentType,omitnil,omitempty" name:"AgentType"`

	// <p>需要重新生成答案的记录ID</p>
	OldRecordId *string `json:"OldRecordId,omitnil,omitempty" name:"OldRecordId"`

	// <p>知识库id列表</p>
	KnowledgeBaseIds []*string `json:"KnowledgeBaseIds,omitnil,omitempty" name:"KnowledgeBaseIds"`

	// <p>版本信息</p>
	ArchVersion *string `json:"ArchVersion,omitnil,omitempty" name:"ArchVersion"`
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
	delete(f, "ArchVersion")
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
	// <p>切片ID</p>
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>切片内容</p>
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// <p>切片的字数</p>
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// <p>切片概要</p>
	Summary *string `json:"Summary,omitnil,omitempty" name:"Summary"`

	// <p>分段类型</p><p>枚举值：</p><ul><li>0： 自动分段</li><li>1： 新建分段</li></ul>
	ChunkSource *int64 `json:"ChunkSource,omitnil,omitempty" name:"ChunkSource"`
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

	// 批量删除 会话id 列表
	SessionIds []*string `json:"SessionIds,omitnil,omitempty" name:"SessionIds"`
}

type DeleteDataAgentSessionRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 会话ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 批量删除 会话id 列表
	SessionIds []*string `json:"SessionIds,omitnil,omitempty" name:"SessionIds"`
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
	delete(f, "SessionIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDataAgentSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDataAgentSessionResponseParams struct {
	// 删除的会话ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 删除的会话ID列表
	SessionIds []*string `json:"SessionIds,omitnil,omitempty" name:"SessionIds"`

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

// Predefined struct for user
type DeleteSceneRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 场景id
	SceneId *string `json:"SceneId,omitnil,omitempty" name:"SceneId"`
}

type DeleteSceneRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 场景id
	SceneId *string `json:"SceneId,omitnil,omitempty" name:"SceneId"`
}

func (r *DeleteSceneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSceneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SceneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSceneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSceneResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSceneResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSceneResponseParams `json:"Response"`
}

func (r *DeleteSceneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSceneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExampleQA struct {
	// 示例记录的唯一业务 ID
	ExampleId *string `json:"ExampleId,omitnil,omitempty" name:"ExampleId"`

	// 问题列表
	Questions []*string `json:"Questions,omitnil,omitempty" name:"Questions"`

	// 对应的标准答案或回复
	Answer *string `json:"Answer,omitnil,omitempty" name:"Answer"`

	// 内容类型，类型包含 'text', 'sql', 'code' 
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 记录的创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 记录的最后更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

// Predefined struct for user
type ExecuteAgentApiRequestParams struct {
	// <p>参数路径</p>
	RequestPath *string `json:"RequestPath,omitnil,omitempty" name:"RequestPath"`

	// <p>参数值</p>
	RequestData *string `json:"RequestData,omitnil,omitempty" name:"RequestData"`

	// <p>post还是get</p><p>枚举值：</p><ul><li>post： post请求</li><li>get： get请求</li></ul>
	RequestType *string `json:"RequestType,omitnil,omitempty" name:"RequestType"`
}

type ExecuteAgentApiRequest struct {
	*tchttp.BaseRequest
	
	// <p>参数路径</p>
	RequestPath *string `json:"RequestPath,omitnil,omitempty" name:"RequestPath"`

	// <p>参数值</p>
	RequestData *string `json:"RequestData,omitnil,omitempty" name:"RequestData"`

	// <p>post还是get</p><p>枚举值：</p><ul><li>post： post请求</li><li>get： get请求</li></ul>
	RequestType *string `json:"RequestType,omitnil,omitempty" name:"RequestType"`
}

func (r *ExecuteAgentApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExecuteAgentApiRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RequestPath")
	delete(f, "RequestData")
	delete(f, "RequestType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExecuteAgentApiRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExecuteAgentApiResponseParams struct {
	// <p>请求路径</p>
	RequestPath *string `json:"RequestPath,omitnil,omitempty" name:"RequestPath"`

	// <p>返回的具体指</p>
	AgentData *string `json:"AgentData,omitnil,omitempty" name:"AgentData"`

	// <p>错误码信息</p>
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ExecuteAgentApiResponse struct {
	*tchttp.BaseResponse
	Response *ExecuteAgentApiResponseParams `json:"Response"`
}

func (r *ExecuteAgentApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExecuteAgentApiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExecuteAgentApiV1RequestParams struct {
	// <p>参数路径</p>
	RequestPath *string `json:"RequestPath,omitnil,omitempty" name:"RequestPath"`

	// <p>post还是get</p><p>枚举值：</p><ul><li>post： post请求</li><li>get： get请求</li></ul>
	RequestType *string `json:"RequestType,omitnil,omitempty" name:"RequestType"`
}

type ExecuteAgentApiV1Request struct {
	*tchttp.BaseRequest
	
	// <p>参数路径</p>
	RequestPath *string `json:"RequestPath,omitnil,omitempty" name:"RequestPath"`

	// <p>post还是get</p><p>枚举值：</p><ul><li>post： post请求</li><li>get： get请求</li></ul>
	RequestType *string `json:"RequestType,omitnil,omitempty" name:"RequestType"`
}

func (r *ExecuteAgentApiV1Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExecuteAgentApiV1Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RequestPath")
	delete(f, "RequestType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExecuteAgentApiV1Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExecuteAgentApiV1ResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。本接口为流式响应接口，当请求成功时，RequestId 会被放在 HTTP 响应的 Header "X-TC-RequestId" 中。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ExecuteAgentApiV1Response struct {
	tchttp.BaseSSEResponse `json:"-"`
	Response *ExecuteAgentApiV1ResponseParams `json:"Response"`
}

func (r *ExecuteAgentApiV1Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExecuteAgentApiV1Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FileInfo struct {
	// <p>文件名称</p>
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// <p>文件大小，字节</p>
	FileSize *float64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// <p>文件类型,0=文本,1=表格，默认0</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>文件ID</p>
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <p>状态，0：数据处理中  1：可用 -1：错误</p>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>操作者</p>
	CreateUser *string `json:"CreateUser,omitnil,omitempty" name:"CreateUser"`

	// <p>创建时间</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>更新时间</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>分片策略</p>
	ChunkConfig *KnowledgeTaskConfig `json:"ChunkConfig,omitnil,omitempty" name:"ChunkConfig"`

	// <p>文件来源0=unknow,1=user_cos,2=local</p>
	Source *int64 `json:"Source,omitnil,omitempty" name:"Source"`

	// <p>文件url</p>
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// <p>是否官方示例，0=否，1=是</p>
	IsShowCase *int64 `json:"IsShowCase,omitnil,omitempty" name:"IsShowCase"`

	// <p>文档摘要</p>
	DocumentSummary *string `json:"DocumentSummary,omitnil,omitempty" name:"DocumentSummary"`

	// <p>网页地址</p>
	WebUrl *string `json:"WebUrl,omitnil,omitempty" name:"WebUrl"`

	// <p>文件能力标识列表</p>
	Capabilities []*string `json:"Capabilities,omitnil,omitempty" name:"Capabilities"`
}

type FileTaskStatus struct {
	// <p>文件id</p>
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <p>状态</p><p>枚举值：</p><ul><li>0： 处理中</li><li>1： 可用</li><li>-1： 错误</li></ul>
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>是否已拉取过状态</p><p>枚举值：</p><ul><li>0： 未被拉取过状态</li><li>1： 已被拉取过状态</li></ul>
	IsTerminated *uint64 `json:"IsTerminated,omitnil,omitempty" name:"IsTerminated"`

	// <p>错误信息，状态-1时不为空</p>
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`
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

// Predefined struct for user
type GetUserInstanceListRequestParams struct {

}

type GetUserInstanceListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *GetUserInstanceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetUserInstanceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetUserInstanceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetUserInstanceListResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetUserInstanceListResponse struct {
	*tchttp.BaseResponse
	Response *GetUserInstanceListResponseParams `json:"Response"`
}

func (r *GetUserInstanceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetUserInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KnowledgeBase struct {
	// <p>知识库id</p>
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`

	// <p>知识库名称</p>
	KnowledgeBaseName *string `json:"KnowledgeBaseName,omitnil,omitempty" name:"KnowledgeBaseName"`

	// <p>知识库描述</p>
	KnowledgeBaseDesc *string `json:"KnowledgeBaseDesc,omitnil,omitempty" name:"KnowledgeBaseDesc"`

	// <p>创建者subuin</p>
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// <p>创建时间</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>文件数量</p>
	FileNum *int64 `json:"FileNum,omitnil,omitempty" name:"FileNum"`

	// <p>知识库关联的数据库列表，目前是只绑定一个数据源，数组预留拓展</p>
	DatasourceIds []*string `json:"DatasourceIds,omitnil,omitempty" name:"DatasourceIds"`

	// <p>知识库任务配置</p>
	Config *KnowledgeTaskConfig `json:"Config,omitnil,omitempty" name:"Config"`
}

type KnowledgeTaskConfig struct {
	// <p>切片类型  0:自定义切片，1：智能切片</p>
	ChunkType *int64 `json:"ChunkType,omitnil,omitempty" name:"ChunkType"`

	// <p>/智能切片：最小值 1000，默认 4800 自定义切片：正整数即可,默认值 1000</p>
	MaxChunkSize *int64 `json:"MaxChunkSize,omitnil,omitempty" name:"MaxChunkSize"`

	// <p>切片分隔符,自定义切片使用：默认值为：[&quot;\n\n&quot;, &quot;\n&quot;, &quot;。&quot;, &quot;！&quot;, &quot;？&quot;, &quot;，&quot;, &quot;&quot;]</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Delimiters []*string `json:"Delimiters,omitnil,omitempty" name:"Delimiters"`

	// <p>自定义切片使用:默认0 可重叠字符长度</p>
	ChunkOverlap *int64 `json:"ChunkOverlap,omitnil,omitempty" name:"ChunkOverlap"`

	// <p>表格类文档解析</p>
	Columns []*ColumnInfo `json:"Columns,omitnil,omitempty" name:"Columns"`

	// <p>带检索的索引列表</p>
	Indexes []*int64 `json:"Indexes,omitnil,omitempty" name:"Indexes"`

	// <p>0：不生成文档摘要，1：生成文档概要。默认0，当取1时，GenParaSummary必须也为1</p>
	GenDocSummary *int64 `json:"GenDocSummary,omitnil,omitempty" name:"GenDocSummary"`

	// <p>0：不生成段落摘要，1：生成段落概要。默认0</p>
	GenParaSummary *int64 `json:"GenParaSummary,omitnil,omitempty" name:"GenParaSummary"`

	// <p>0：不开启图片理解，1：开启图片理解。默认1</p><p>取值范围：[1, 10000]</p><p>默认值：1</p>
	EnableImageUnderstanding *int64 `json:"EnableImageUnderstanding,omitnil,omitempty" name:"EnableImageUnderstanding"`

	// <p>是否开启表格结构化提取</p><p>枚举值：</p><ul><li>0： 不开启表格提取</li><li>1： 开启表格提取</li></ul><p>默认值：1</p>
	EnableExtractDb *int64 `json:"EnableExtractDb,omitnil,omitempty" name:"EnableExtractDb"`
}

type ModelUserAuthority struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 模块，分为知识库knowledge、数据源datasource、自定义场景scene
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// 对象创建者
	CreatorUin *string `json:"CreatorUin,omitnil,omitempty" name:"CreatorUin"`

	// 对象id,分为知识库id、数据源id、场景id
	ObjectId *string `json:"ObjectId,omitnil,omitempty" name:"ObjectId"`

	// 作用范围：1仅自己使用，2指定用户，0全员
	UseScope *int64 `json:"UseScope,omitnil,omitempty" name:"UseScope"`

	// 可使用的用户列表
	AuthorityUins []*string `json:"AuthorityUins,omitnil,omitempty" name:"AuthorityUins"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

// Predefined struct for user
type ModifyChunkRequestParams struct {
	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>文件ID</p>
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <p>切片ID</p>
	ChunkId *string `json:"ChunkId,omitnil,omitempty" name:"ChunkId"`

	// <p>编辑后的文本</p>
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// <p>分段概要</p>
	Summary *string `json:"Summary,omitnil,omitempty" name:"Summary"`

	// <p>知识库id</p>
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`
}

type ModifyChunkRequest struct {
	*tchttp.BaseRequest
	
	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>文件ID</p>
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <p>切片ID</p>
	ChunkId *string `json:"ChunkId,omitnil,omitempty" name:"ChunkId"`

	// <p>编辑后的文本</p>
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// <p>分段概要</p>
	Summary *string `json:"Summary,omitnil,omitempty" name:"Summary"`

	// <p>知识库id</p>
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
	delete(f, "Summary")
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
	// <p>实例id</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>操作类型：Create，Update，Delete</p>
	OperateType *string `json:"OperateType,omitnil,omitempty" name:"OperateType"`

	// <p>知识库id，update和delete时必填</p>
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`

	// <p>知识库名称，create和update时必填。只允许字母、数字、汉字、下划线</p>
	KnowledgeBaseName *string `json:"KnowledgeBaseName,omitnil,omitempty" name:"KnowledgeBaseName"`

	// <p>知识库描述，create和update时必填</p>
	KnowledgeBaseDesc *string `json:"KnowledgeBaseDesc,omitnil,omitempty" name:"KnowledgeBaseDesc"`

	// <p>1仅自己使用，2指定用户，0全员</p>
	UseScope *int64 `json:"UseScope,omitnil,omitempty" name:"UseScope"`

	// <p>可使用用户列表</p>
	AuthorityUins []*string `json:"AuthorityUins,omitnil,omitempty" name:"AuthorityUins"`

	// <p>知识库任务配置</p>
	Config *KnowledgeTaskConfig `json:"Config,omitnil,omitempty" name:"Config"`
}

type ModifyKnowledgeBaseRequest struct {
	*tchttp.BaseRequest
	
	// <p>实例id</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>操作类型：Create，Update，Delete</p>
	OperateType *string `json:"OperateType,omitnil,omitempty" name:"OperateType"`

	// <p>知识库id，update和delete时必填</p>
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`

	// <p>知识库名称，create和update时必填。只允许字母、数字、汉字、下划线</p>
	KnowledgeBaseName *string `json:"KnowledgeBaseName,omitnil,omitempty" name:"KnowledgeBaseName"`

	// <p>知识库描述，create和update时必填</p>
	KnowledgeBaseDesc *string `json:"KnowledgeBaseDesc,omitnil,omitempty" name:"KnowledgeBaseDesc"`

	// <p>1仅自己使用，2指定用户，0全员</p>
	UseScope *int64 `json:"UseScope,omitnil,omitempty" name:"UseScope"`

	// <p>可使用用户列表</p>
	AuthorityUins []*string `json:"AuthorityUins,omitnil,omitempty" name:"AuthorityUins"`

	// <p>知识库任务配置</p>
	Config *KnowledgeTaskConfig `json:"Config,omitnil,omitempty" name:"Config"`
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
	delete(f, "Config")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyKnowledgeBaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyKnowledgeBaseResponseParams struct {
	// <p>知识库id</p>
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
type ModifyUserAuthorityRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分为知识库knowledge、数据源datasource、自定义场景scene
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// 对象id,分为知识库id、数据源id、场景id
	ObjectId *string `json:"ObjectId,omitnil,omitempty" name:"ObjectId"`

	// 作用范围：1仅自己使用，2指定用户，0全员
	UseScope *int64 `json:"UseScope,omitnil,omitempty" name:"UseScope"`

	// 可使用的用户列表，UseScope=0/1,取值为[]
	AuthorityUins []*string `json:"AuthorityUins,omitnil,omitempty" name:"AuthorityUins"`
}

type ModifyUserAuthorityRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分为知识库knowledge、数据源datasource、自定义场景scene
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// 对象id,分为知识库id、数据源id、场景id
	ObjectId *string `json:"ObjectId,omitnil,omitempty" name:"ObjectId"`

	// 作用范围：1仅自己使用，2指定用户，0全员
	UseScope *int64 `json:"UseScope,omitnil,omitempty" name:"UseScope"`

	// 可使用的用户列表，UseScope=0/1,取值为[]
	AuthorityUins []*string `json:"AuthorityUins,omitnil,omitempty" name:"AuthorityUins"`
}

func (r *ModifyUserAuthorityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserAuthorityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Module")
	delete(f, "ObjectId")
	delete(f, "UseScope")
	delete(f, "AuthorityUins")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserAuthorityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserAuthorityResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyUserAuthorityResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUserAuthorityResponseParams `json:"Response"`
}

func (r *ModifyUserAuthorityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserAuthorityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryChunkListRequestParams struct {
	// <p>表示第一页</p>
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// <p>默认一页展示 10 条</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>知识库id</p>
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`
}

type QueryChunkListRequest struct {
	*tchttp.BaseRequest
	
	// <p>表示第一页</p>
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// <p>默认一页展示 10 条</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>知识库id</p>
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
	// <p>总数</p>
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// <p>文档的自动分段数</p>
	AutoTotal *int64 `json:"AutoTotal,omitnil,omitempty" name:"AutoTotal"`

	// <p>文档的手动新建分段数</p>
	ManualTotal *int64 `json:"ManualTotal,omitnil,omitempty" name:"ManualTotal"`

	// <p>分片信息</p>
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

// Predefined struct for user
type QueryKnowledgeTaskRequestParams struct {
	// <p>实例id</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>知识库id</p>
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`

	// <p>文件id列表</p>
	FileIds []*string `json:"FileIds,omitnil,omitempty" name:"FileIds"`
}

type QueryKnowledgeTaskRequest struct {
	*tchttp.BaseRequest
	
	// <p>实例id</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>知识库id</p>
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitnil,omitempty" name:"KnowledgeBaseId"`

	// <p>文件id列表</p>
	FileIds []*string `json:"FileIds,omitnil,omitempty" name:"FileIds"`
}

func (r *QueryKnowledgeTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryKnowledgeTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "KnowledgeBaseId")
	delete(f, "FileIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryKnowledgeTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryKnowledgeTaskResponseParams struct {
	// <p>文档任务详情对象</p>
	FileList []*FileTaskStatus `json:"FileList,omitnil,omitempty" name:"FileList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryKnowledgeTaskResponse struct {
	*tchttp.BaseResponse
	Response *QueryKnowledgeTaskResponseParams `json:"Response"`
}

func (r *QueryKnowledgeTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryKnowledgeTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QuerySceneListRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 场景id
	SceneId *string `json:"SceneId,omitnil,omitempty" name:"SceneId"`

	// 场景名称
	SceneName *string `json:"SceneName,omitnil,omitempty" name:"SceneName"`

	// 页数
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 页的大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type QuerySceneListRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 场景id
	SceneId *string `json:"SceneId,omitnil,omitempty" name:"SceneId"`

	// 场景名称
	SceneName *string `json:"SceneName,omitnil,omitempty" name:"SceneName"`

	// 页数
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 页的大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *QuerySceneListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QuerySceneListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SceneId")
	delete(f, "SceneName")
	delete(f, "Page")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QuerySceneListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QuerySceneListResponseParams struct {
	// 场景列表
	Datas []*Scene `json:"Datas,omitnil,omitempty" name:"Datas"`

	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QuerySceneListResponse struct {
	*tchttp.BaseResponse
	Response *QuerySceneListResponseParams `json:"Response"`
}

func (r *QuerySceneListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QuerySceneListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryUserAuthorityRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分为知识库knowledge、数据源datasource、自定义场景scene
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// 对象id,分为知识库id、数据源id、场景id
	ObjectId *string `json:"ObjectId,omitnil,omitempty" name:"ObjectId"`
}

type QueryUserAuthorityRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分为知识库knowledge、数据源datasource、自定义场景scene
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// 对象id,分为知识库id、数据源id、场景id
	ObjectId *string `json:"ObjectId,omitnil,omitempty" name:"ObjectId"`
}

func (r *QueryUserAuthorityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryUserAuthorityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Module")
	delete(f, "ObjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryUserAuthorityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryUserAuthorityResponseParams struct {
	// 对象权限信息
	ModelUserAuthority *ModelUserAuthority `json:"ModelUserAuthority,omitnil,omitempty" name:"ModelUserAuthority"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryUserAuthorityResponse struct {
	*tchttp.BaseResponse
	Response *QueryUserAuthorityResponseParams `json:"Response"`
}

func (r *QueryUserAuthorityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryUserAuthorityResponse) FromJsonString(s string) error {
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

type Scene struct {
	// <p>场景ID</p>
	SceneId *string `json:"SceneId,omitnil,omitempty" name:"SceneId"`

	// <p>场景名称</p>
	SceneName *string `json:"SceneName,omitnil,omitempty" name:"SceneName"`

	// <p>技能列表，包含：rag（知识检索）、data_analytics（数据分析）、data_prediction（数据预测）</p>
	Skills []*string `json:"Skills,omitnil,omitempty" name:"Skills"`

	// <p>提示词文本</p>
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// <p>描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>检索配置</p>
	SearchConfig *SearchConfig `json:"SearchConfig,omitnil,omitempty" name:"SearchConfig"`

	// <p>示例问答列表</p>
	ExampleQAList []*ExampleQA `json:"ExampleQAList,omitnil,omitempty" name:"ExampleQAList"`

	// <p>记录的创建时间</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>记录的最后更新时间</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>创建者Uin</p>
	CreatorUin *string `json:"CreatorUin,omitnil,omitempty" name:"CreatorUin"`

	// <p>知识</p>
	Knowledge *string `json:"Knowledge,omitnil,omitempty" name:"Knowledge"`
}

type SearchConfig struct {
	// 检索类型：0:混合搜索 1：向量搜索 2：全文搜索
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 召回数量最大值
	Num *int64 `json:"Num,omitnil,omitempty" name:"Num"`

	// 权重配置
	EmbeddingWeight *float64 `json:"EmbeddingWeight,omitnil,omitempty" name:"EmbeddingWeight"`

	// 0:关闭 1:开启，默认1
	Rerank *int64 `json:"Rerank,omitnil,omitempty" name:"Rerank"`

	// 0:关闭 1:开启，默认0
	AutoRag *int64 `json:"AutoRag,omitnil,omitempty" name:"AutoRag"`

	// AutoRag关联的知识库ID列表
	KnowledgeBaseIds []*string `json:"KnowledgeBaseIds,omitnil,omitempty" name:"KnowledgeBaseIds"`

	// AutoRag搜索状态：0-未完成，1-已完成。仅当AutoRag=1时，该字段有效
	SearchStatus *int64 `json:"SearchStatus,omitnil,omitempty" name:"SearchStatus"`
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
type UpdateSceneRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 场景
	Scene *Scene `json:"Scene,omitnil,omitempty" name:"Scene"`
}

type UpdateSceneRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 场景
	Scene *Scene `json:"Scene,omitnil,omitempty" name:"Scene"`
}

func (r *UpdateSceneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateSceneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Scene")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateSceneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateSceneResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateSceneResponse struct {
	*tchttp.BaseResponse
	Response *UpdateSceneResponseParams `json:"Response"`
}

func (r *UpdateSceneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateSceneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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