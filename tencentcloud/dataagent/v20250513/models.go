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

	// 	 新 Chunk 插入到目标 Chunk ​之后的位置。插入位置的上一个 chunkId
	AfterChunkId *string `json:"AfterChunkId,omitnil,omitempty" name:"AfterChunkId"`
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

	// 	 新 Chunk 插入到目标 Chunk ​之后的位置。插入位置的上一个 chunkId
	AfterChunkId *string `json:"AfterChunkId,omitnil,omitempty" name:"AfterChunkId"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddChunkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddChunkResponseParams struct {
	// 新增的chunkid
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

// Predefined struct for user
type CreateDataAgentSessionRequestParams struct {

}

type CreateDataAgentSessionRequest struct {
	*tchttp.BaseRequest
	
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
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDataAgentSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDataAgentSessionResponseParams struct {
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
}

type DeleteChunkRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 文件ID
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 切片ID
	ChunkIds []*string `json:"ChunkIds,omitnil,omitempty" name:"ChunkIds"`
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

}

type DeleteDataAgentSessionRequest struct {
	*tchttp.BaseRequest
	
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
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDataAgentSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDataAgentSessionResponseParams struct {
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

}

type GetSessionDetailsRequest struct {
	*tchttp.BaseRequest
	
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
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetSessionDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetSessionDetailsResponseParams struct {
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
	// 默认 1 表示第一页
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 	 默认 10 一页展示 10 条
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type QueryChunkListRequest struct {
	*tchttp.BaseRequest
	
	// 默认 1 表示第一页
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 	 默认 10 一页展示 10 条
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryChunkListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryChunkListResponseParams struct {
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
type StopChatAIRequestParams struct {

}

type StopChatAIRequest struct {
	*tchttp.BaseRequest
	
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
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopChatAIRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopChatAIResponseParams struct {
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