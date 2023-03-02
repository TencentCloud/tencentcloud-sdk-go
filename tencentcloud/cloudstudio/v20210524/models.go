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

package v20210524

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AgentSpaceDTO struct {
	// 工作空间名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 镜像id
	ImageId *int64 `json:"ImageId,omitempty" name:"ImageId"`

	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 云服务器登录名称
	RemoteUser *string `json:"RemoteUser,omitempty" name:"RemoteUser"`

	// 云服务器登录地址
	RemoteHost *string `json:"RemoteHost,omitempty" name:"RemoteHost"`

	// 云服务器登录端口
	RemotePort *string `json:"RemotePort,omitempty" name:"RemotePort"`

	// 工作空间类型
	WorkspaceType *string `json:"WorkspaceType,omitempty" name:"WorkspaceType"`

	// 工作空间版本
	WorkspaceVersion *int64 `json:"WorkspaceVersion,omitempty" name:"WorkspaceVersion"`

	// 工作空间资源结构
	WorkspaceResourceDTO *WorkspaceResourceDTO `json:"WorkspaceResourceDTO,omitempty" name:"WorkspaceResourceDTO"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

// Predefined struct for user
type CreateCustomizeTemplatesRequestParams struct {
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitempty" name:"CloudStudioSessionTeam"`

	// 无
	UserDefinedTemplateParams *UserDefinedTemplateParams `json:"UserDefinedTemplateParams,omitempty" name:"UserDefinedTemplateParams"`
}

type CreateCustomizeTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitempty" name:"CloudStudioSessionTeam"`

	// 无
	UserDefinedTemplateParams *UserDefinedTemplateParams `json:"UserDefinedTemplateParams,omitempty" name:"UserDefinedTemplateParams"`
}

func (r *CreateCustomizeTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomizeTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CloudStudioSessionTeam")
	delete(f, "UserDefinedTemplateParams")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCustomizeTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCustomizeTemplatesResponseParams struct {
	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *WorkspaceTemplateInfo `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCustomizeTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *CreateCustomizeTemplatesResponseParams `json:"Response"`
}

func (r *CreateCustomizeTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomizeTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkspaceByAgentRequestParams struct {
	// 无
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitempty" name:"CloudStudioSessionTeam"`

	// 无
	AgentSpaceDTO *AgentSpaceDTO `json:"AgentSpaceDTO,omitempty" name:"AgentSpaceDTO"`
}

type CreateWorkspaceByAgentRequest struct {
	*tchttp.BaseRequest
	
	// 无
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitempty" name:"CloudStudioSessionTeam"`

	// 无
	AgentSpaceDTO *AgentSpaceDTO `json:"AgentSpaceDTO,omitempty" name:"AgentSpaceDTO"`
}

func (r *CreateWorkspaceByAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkspaceByAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CloudStudioSessionTeam")
	delete(f, "AgentSpaceDTO")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWorkspaceByAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkspaceByAgentResponseParams struct {
	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *WorkspaceInfoDTO `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateWorkspaceByAgentResponse struct {
	*tchttp.BaseResponse
	Response *CreateWorkspaceByAgentResponseParams `json:"Response"`
}

func (r *CreateWorkspaceByAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkspaceByAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkspaceByTemplateRequestParams struct {
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitempty" name:"CloudStudioSessionTeam"`

	// 模板ID
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

type CreateWorkspaceByTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitempty" name:"CloudStudioSessionTeam"`

	// 模板ID
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *CreateWorkspaceByTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkspaceByTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CloudStudioSessionTeam")
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWorkspaceByTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkspaceByTemplateResponseParams struct {
	// 创建工作空间返回的信息
	Data *WorkspaceInfo `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateWorkspaceByTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateWorkspaceByTemplateResponseParams `json:"Response"`
}

func (r *CreateWorkspaceByTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkspaceByTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkspaceByVersionControlRequestParams struct {
	// 工作空间结构
	WorkspaceDTO *WorkspaceDTO `json:"WorkspaceDTO,omitempty" name:"WorkspaceDTO"`

	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitempty" name:"CloudStudioSessionTeam"`
}

type CreateWorkspaceByVersionControlRequest struct {
	*tchttp.BaseRequest
	
	// 工作空间结构
	WorkspaceDTO *WorkspaceDTO `json:"WorkspaceDTO,omitempty" name:"WorkspaceDTO"`

	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitempty" name:"CloudStudioSessionTeam"`
}

func (r *CreateWorkspaceByVersionControlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkspaceByVersionControlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceDTO")
	delete(f, "CloudStudioSessionTeam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWorkspaceByVersionControlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkspaceByVersionControlResponseParams struct {
	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *WorkspaceInfoDTO `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateWorkspaceByVersionControlResponse struct {
	*tchttp.BaseResponse
	Response *CreateWorkspaceByVersionControlResponseParams `json:"Response"`
}

func (r *CreateWorkspaceByVersionControlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkspaceByVersionControlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkspaceTemporaryTokenRequestParams struct {
	// 创建工作空间凭证 DTO
	WorkspaceTokenDTO *WorkspaceTokenDTO `json:"WorkspaceTokenDTO,omitempty" name:"WorkspaceTokenDTO"`
}

type CreateWorkspaceTemporaryTokenRequest struct {
	*tchttp.BaseRequest
	
	// 创建工作空间凭证 DTO
	WorkspaceTokenDTO *WorkspaceTokenDTO `json:"WorkspaceTokenDTO,omitempty" name:"WorkspaceTokenDTO"`
}

func (r *CreateWorkspaceTemporaryTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkspaceTemporaryTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceTokenDTO")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWorkspaceTemporaryTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkspaceTemporaryTokenResponseParams struct {
	// 工作空间临时访问 token 信息
	Data *WorkspaceTokenInfoV0 `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateWorkspaceTemporaryTokenResponse struct {
	*tchttp.BaseResponse
	Response *CreateWorkspaceTemporaryTokenResponseParams `json:"Response"`
}

func (r *CreateWorkspaceTemporaryTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkspaceTemporaryTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomizeTemplatesPresetsInfo struct {
	// 模板tag列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*string `json:"Tags,omitempty" name:"Tags"`

	// 模板图标列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Icons []*string `json:"Icons,omitempty" name:"Icons"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Templates *UserDefinedTemplateParams `json:"Templates,omitempty" name:"Templates"`
}

// Predefined struct for user
type DeleteCustomizeTemplatesByIdRequestParams struct {
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitempty" name:"CloudStudioSessionTeam"`

	// 模板ID
	Id *int64 `json:"Id,omitempty" name:"Id"`
}

type DeleteCustomizeTemplatesByIdRequest struct {
	*tchttp.BaseRequest
	
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitempty" name:"CloudStudioSessionTeam"`

	// 模板ID
	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteCustomizeTemplatesByIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCustomizeTemplatesByIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CloudStudioSessionTeam")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCustomizeTemplatesByIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCustomizeTemplatesByIdResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteCustomizeTemplatesByIdResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCustomizeTemplatesByIdResponseParams `json:"Response"`
}

func (r *DeleteCustomizeTemplatesByIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCustomizeTemplatesByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomizeTemplatesByIdRequestParams struct {
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitempty" name:"CloudStudioSessionTeam"`

	// 模板ID
	Id *int64 `json:"Id,omitempty" name:"Id"`
}

type DescribeCustomizeTemplatesByIdRequest struct {
	*tchttp.BaseRequest
	
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitempty" name:"CloudStudioSessionTeam"`

	// 模板ID
	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeCustomizeTemplatesByIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomizeTemplatesByIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CloudStudioSessionTeam")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomizeTemplatesByIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomizeTemplatesByIdResponseParams struct {
	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *WorkspaceTemplateInfo `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCustomizeTemplatesByIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomizeTemplatesByIdResponseParams `json:"Response"`
}

func (r *DescribeCustomizeTemplatesByIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomizeTemplatesByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomizeTemplatesPresetsRequestParams struct {
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitempty" name:"CloudStudioSessionTeam"`

	// 空间标识
	SpaceKey *string `json:"SpaceKey,omitempty" name:"SpaceKey"`
}

type DescribeCustomizeTemplatesPresetsRequest struct {
	*tchttp.BaseRequest
	
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitempty" name:"CloudStudioSessionTeam"`

	// 空间标识
	SpaceKey *string `json:"SpaceKey,omitempty" name:"SpaceKey"`
}

func (r *DescribeCustomizeTemplatesPresetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomizeTemplatesPresetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CloudStudioSessionTeam")
	delete(f, "SpaceKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomizeTemplatesPresetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomizeTemplatesPresetsResponseParams struct {
	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *CustomizeTemplatesPresetsInfo `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCustomizeTemplatesPresetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomizeTemplatesPresetsResponseParams `json:"Response"`
}

func (r *DescribeCustomizeTemplatesPresetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomizeTemplatesPresetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomizeTemplatesRequestParams struct {
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitempty" name:"CloudStudioSessionTeam"`
}

type DescribeCustomizeTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitempty" name:"CloudStudioSessionTeam"`
}

func (r *DescribeCustomizeTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomizeTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CloudStudioSessionTeam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomizeTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomizeTemplatesResponseParams struct {
	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*WorkspaceTemplateInfo `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCustomizeTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomizeTemplatesResponseParams `json:"Response"`
}

func (r *DescribeCustomizeTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomizeTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkspaceEnvListRequestParams struct {
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitempty" name:"CloudStudioSessionTeam"`
}

type DescribeWorkspaceEnvListRequest struct {
	*tchttp.BaseRequest
	
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitempty" name:"CloudStudioSessionTeam"`
}

func (r *DescribeWorkspaceEnvListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkspaceEnvListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CloudStudioSessionTeam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWorkspaceEnvListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkspaceEnvListResponseParams struct {
	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*ImageUserDTO `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWorkspaceEnvListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWorkspaceEnvListResponseParams `json:"Response"`
}

func (r *DescribeWorkspaceEnvListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkspaceEnvListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkspaceNameExistRequestParams struct {
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitempty" name:"CloudStudioSessionTeam"`

	// 工作空间名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 工作空间ID
	WorkspaceId *string `json:"WorkspaceId,omitempty" name:"WorkspaceId"`
}

type DescribeWorkspaceNameExistRequest struct {
	*tchttp.BaseRequest
	
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitempty" name:"CloudStudioSessionTeam"`

	// 工作空间名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 工作空间ID
	WorkspaceId *string `json:"WorkspaceId,omitempty" name:"WorkspaceId"`
}

func (r *DescribeWorkspaceNameExistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkspaceNameExistRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CloudStudioSessionTeam")
	delete(f, "Name")
	delete(f, "WorkspaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWorkspaceNameExistRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkspaceNameExistResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWorkspaceNameExistResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWorkspaceNameExistResponseParams `json:"Response"`
}

func (r *DescribeWorkspaceNameExistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkspaceNameExistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkspaceStatusListRequestParams struct {
	// xxx
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitempty" name:"CloudStudioSessionTeam"`
}

type DescribeWorkspaceStatusListRequest struct {
	*tchttp.BaseRequest
	
	// xxx
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitempty" name:"CloudStudioSessionTeam"`
}

func (r *DescribeWorkspaceStatusListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkspaceStatusListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CloudStudioSessionTeam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWorkspaceStatusListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkspaceStatusListResponseParams struct {
	// xxx
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*WorkspaceStatusInfo `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWorkspaceStatusListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWorkspaceStatusListResponseParams `json:"Response"`
}

func (r *DescribeWorkspaceStatusListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkspaceStatusListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkspaceStatusRequestParams struct {
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitempty" name:"CloudStudioSessionTeam"`

	// 空间标识
	SpaceKey *string `json:"SpaceKey,omitempty" name:"SpaceKey"`
}

type DescribeWorkspaceStatusRequest struct {
	*tchttp.BaseRequest
	
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitempty" name:"CloudStudioSessionTeam"`

	// 空间标识
	SpaceKey *string `json:"SpaceKey,omitempty" name:"SpaceKey"`
}

func (r *DescribeWorkspaceStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkspaceStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CloudStudioSessionTeam")
	delete(f, "SpaceKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWorkspaceStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkspaceStatusResponseParams struct {
	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *WorkspaceStatusInfo `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWorkspaceStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWorkspaceStatusResponseParams `json:"Response"`
}

func (r *DescribeWorkspaceStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkspaceStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageUserDTO struct {
	// 镜像模板ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 镜像模板名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// Tag时间
	Tag *string `json:"Tag,omitempty" name:"Tag"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 中文描述
	DescriptionCN *string `json:"DescriptionCN,omitempty" name:"DescriptionCN"`

	// 图标地址
	IconUrl *string `json:"IconUrl,omitempty" name:"IconUrl"`

	// 创建人
	Author *string `json:"Author,omitempty" name:"Author"`

	// 访问状态
	Visible *string `json:"Visible,omitempty" name:"Visible"`

	// 版本
	WorkspaceVersion *int64 `json:"WorkspaceVersion,omitempty" name:"WorkspaceVersion"`

	// 分类
	Sort *int64 `json:"Sort,omitempty" name:"Sort"`
}

// Predefined struct for user
type ModifyCustomizeTemplateVersionControlRequestParams struct {
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitempty" name:"CloudStudioSessionTeam"`

	// 模板ID
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 仓库地址
	Url *string `json:"Url,omitempty" name:"Url"`

	// 代码仓库分支/标签
	Ref *string `json:"Ref,omitempty" name:"Ref"`

	// 代码仓库 ref 类型
	RefType *string `json:"RefType,omitempty" name:"RefType"`
}

type ModifyCustomizeTemplateVersionControlRequest struct {
	*tchttp.BaseRequest
	
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitempty" name:"CloudStudioSessionTeam"`

	// 模板ID
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 仓库地址
	Url *string `json:"Url,omitempty" name:"Url"`

	// 代码仓库分支/标签
	Ref *string `json:"Ref,omitempty" name:"Ref"`

	// 代码仓库 ref 类型
	RefType *string `json:"RefType,omitempty" name:"RefType"`
}

func (r *ModifyCustomizeTemplateVersionControlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomizeTemplateVersionControlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CloudStudioSessionTeam")
	delete(f, "TemplateId")
	delete(f, "Url")
	delete(f, "Ref")
	delete(f, "RefType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCustomizeTemplateVersionControlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomizeTemplateVersionControlResponseParams struct {
	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *WorkspaceTemplateInfo `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyCustomizeTemplateVersionControlResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCustomizeTemplateVersionControlResponseParams `json:"Response"`
}

func (r *ModifyCustomizeTemplateVersionControlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomizeTemplateVersionControlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomizeTemplatesFullByIdRequestParams struct {
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitempty" name:"CloudStudioSessionTeam"`

	// 模板ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 自定义模板参数
	UserDefinedTemplateParams *UserDefinedTemplateParams `json:"UserDefinedTemplateParams,omitempty" name:"UserDefinedTemplateParams"`
}

type ModifyCustomizeTemplatesFullByIdRequest struct {
	*tchttp.BaseRequest
	
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitempty" name:"CloudStudioSessionTeam"`

	// 模板ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 自定义模板参数
	UserDefinedTemplateParams *UserDefinedTemplateParams `json:"UserDefinedTemplateParams,omitempty" name:"UserDefinedTemplateParams"`
}

func (r *ModifyCustomizeTemplatesFullByIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomizeTemplatesFullByIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CloudStudioSessionTeam")
	delete(f, "Id")
	delete(f, "UserDefinedTemplateParams")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCustomizeTemplatesFullByIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomizeTemplatesFullByIdResponseParams struct {
	// 自定义模板返回值
	Data *WorkspaceTemplateInfo `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyCustomizeTemplatesFullByIdResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCustomizeTemplatesFullByIdResponseParams `json:"Response"`
}

func (r *ModifyCustomizeTemplatesFullByIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomizeTemplatesFullByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomizeTemplatesPartByIdRequestParams struct {
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitempty" name:"CloudStudioSessionTeam"`

	// 模板ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 自定义模板Patched参数
	UserDefinedTemplatePatchedParams *UserDefinedTemplatePatchedParams `json:"UserDefinedTemplatePatchedParams,omitempty" name:"UserDefinedTemplatePatchedParams"`
}

type ModifyCustomizeTemplatesPartByIdRequest struct {
	*tchttp.BaseRequest
	
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitempty" name:"CloudStudioSessionTeam"`

	// 模板ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 自定义模板Patched参数
	UserDefinedTemplatePatchedParams *UserDefinedTemplatePatchedParams `json:"UserDefinedTemplatePatchedParams,omitempty" name:"UserDefinedTemplatePatchedParams"`
}

func (r *ModifyCustomizeTemplatesPartByIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomizeTemplatesPartByIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CloudStudioSessionTeam")
	delete(f, "Id")
	delete(f, "UserDefinedTemplatePatchedParams")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCustomizeTemplatesPartByIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomizeTemplatesPartByIdResponseParams struct {
	// 自定义模板返回值
	Data *WorkspaceTemplateInfo `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyCustomizeTemplatesPartByIdResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCustomizeTemplatesPartByIdResponseParams `json:"Response"`
}

func (r *ModifyCustomizeTemplatesPartByIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomizeTemplatesPartByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWorkspaceAttributesRequestParams struct {
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitempty" name:"CloudStudioSessionTeam"`

	// 工作空间ID
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" name:"WorkspaceId"`

	// 工作空间名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 工作空间描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// xxx
	PriceId *int64 `json:"PriceId,omitempty" name:"PriceId"`
}

type ModifyWorkspaceAttributesRequest struct {
	*tchttp.BaseRequest
	
	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitempty" name:"CloudStudioSessionTeam"`

	// 工作空间ID
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" name:"WorkspaceId"`

	// 工作空间名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 工作空间描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// xxx
	PriceId *int64 `json:"PriceId,omitempty" name:"PriceId"`
}

func (r *ModifyWorkspaceAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWorkspaceAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CloudStudioSessionTeam")
	delete(f, "WorkspaceId")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "PriceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyWorkspaceAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWorkspaceAttributesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyWorkspaceAttributesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyWorkspaceAttributesResponseParams `json:"Response"`
}

func (r *ModifyWorkspaceAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWorkspaceAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecoverWorkspaceRequestParams struct {
	// 无
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitempty" name:"CloudStudioSessionTeam"`

	// 无
	SpaceKey *string `json:"SpaceKey,omitempty" name:"SpaceKey"`
}

type RecoverWorkspaceRequest struct {
	*tchttp.BaseRequest
	
	// 无
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitempty" name:"CloudStudioSessionTeam"`

	// 无
	SpaceKey *string `json:"SpaceKey,omitempty" name:"SpaceKey"`
}

func (r *RecoverWorkspaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecoverWorkspaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CloudStudioSessionTeam")
	delete(f, "SpaceKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecoverWorkspaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecoverWorkspaceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RecoverWorkspaceResponse struct {
	*tchttp.BaseResponse
	Response *RecoverWorkspaceResponseParams `json:"Response"`
}

func (r *RecoverWorkspaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecoverWorkspaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveWorkspaceRequestParams struct {
	// 无
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitempty" name:"CloudStudioSessionTeam"`

	// 无
	SpaceKey *string `json:"SpaceKey,omitempty" name:"SpaceKey"`

	// 是否强制，true或者false
	Force *bool `json:"Force,omitempty" name:"Force"`
}

type RemoveWorkspaceRequest struct {
	*tchttp.BaseRequest
	
	// 无
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitempty" name:"CloudStudioSessionTeam"`

	// 无
	SpaceKey *string `json:"SpaceKey,omitempty" name:"SpaceKey"`

	// 是否强制，true或者false
	Force *bool `json:"Force,omitempty" name:"Force"`
}

func (r *RemoveWorkspaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveWorkspaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CloudStudioSessionTeam")
	delete(f, "SpaceKey")
	delete(f, "Force")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveWorkspaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveWorkspaceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RemoveWorkspaceResponse struct {
	*tchttp.BaseResponse
	Response *RemoveWorkspaceResponseParams `json:"Response"`
}

func (r *RemoveWorkspaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveWorkspaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunWorkspaceRequestParams struct {
	// 空间标识
	SpaceKey *string `json:"SpaceKey,omitempty" name:"SpaceKey"`

	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitempty" name:"CloudStudioSessionTeam"`
}

type RunWorkspaceRequest struct {
	*tchttp.BaseRequest
	
	// 空间标识
	SpaceKey *string `json:"SpaceKey,omitempty" name:"SpaceKey"`

	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitempty" name:"CloudStudioSessionTeam"`
}

func (r *RunWorkspaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunWorkspaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceKey")
	delete(f, "CloudStudioSessionTeam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunWorkspaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunWorkspaceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RunWorkspaceResponse struct {
	*tchttp.BaseResponse
	Response *RunWorkspaceResponseParams `json:"Response"`
}

func (r *RunWorkspaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunWorkspaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopWorkspaceRequestParams struct {
	// 空间标识
	SpaceKey *string `json:"SpaceKey,omitempty" name:"SpaceKey"`

	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitempty" name:"CloudStudioSessionTeam"`

	// 是否强制终止，true或者false
	Force *string `json:"Force,omitempty" name:"Force"`
}

type StopWorkspaceRequest struct {
	*tchttp.BaseRequest
	
	// 空间标识
	SpaceKey *string `json:"SpaceKey,omitempty" name:"SpaceKey"`

	// 用户所属组
	CloudStudioSessionTeam *string `json:"CloudStudioSessionTeam,omitempty" name:"CloudStudioSessionTeam"`

	// 是否强制终止，true或者false
	Force *string `json:"Force,omitempty" name:"Force"`
}

func (r *StopWorkspaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopWorkspaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceKey")
	delete(f, "CloudStudioSessionTeam")
	delete(f, "Force")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopWorkspaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopWorkspaceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StopWorkspaceResponse struct {
	*tchttp.BaseResponse
	Response *StopWorkspaceResponseParams `json:"Response"`
}

func (r *StopWorkspaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopWorkspaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserDefinedTemplateParams struct {
	// 模板名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 模板图标地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Icon *string `json:"Icon,omitempty" name:"Icon"`

	// 模板标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*string `json:"Tags,omitempty" name:"Tags"`

	// 模板来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *string `json:"Source,omitempty" name:"Source"`

	// 模板描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 模板仓库类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionControlType *string `json:"VersionControlType,omitempty" name:"VersionControlType"`

	// 模板地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionControlUrl *string `json:"VersionControlUrl,omitempty" name:"VersionControlUrl"`
}

type UserDefinedTemplatePatchedParams struct {
	// 模板来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *string `json:"Source,omitempty" name:"Source"`

	// 模板名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 模板图标地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Icon *string `json:"Icon,omitempty" name:"Icon"`

	// 模板描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 模板标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*string `json:"Tags,omitempty" name:"Tags"`
}

type UserInfoRsp struct {
	// 用户ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 用户验证信息
	AuthenticationUserInfo *UserSubInfo `json:"AuthenticationUserInfo,omitempty" name:"AuthenticationUserInfo"`

	// 头像地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Avatar *string `json:"Avatar,omitempty" name:"Avatar"`

	// 介绍
	// 注意：此字段可能返回 null，表示取不到有效值。
	Features *string `json:"Features,omitempty" name:"Features"`

	// 状况
	PreviewStatus *int64 `json:"PreviewStatus,omitempty" name:"PreviewStatus"`
}

type UserSubInfo struct {
	// 团队名称
	Team *string `json:"Team,omitempty" name:"Team"`

	// 用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 昵称
	NickName *string `json:"NickName,omitempty" name:"NickName"`

	// 是否为管理员
	IsAdmin *bool `json:"IsAdmin,omitempty" name:"IsAdmin"`

	// xxx
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsTrial *bool `json:"IsTrial,omitempty" name:"IsTrial"`
}

type WorkspaceDTO struct {
	// 工作空间名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 代码来源类型
	VersionControlType *string `json:"VersionControlType,omitempty" name:"VersionControlType"`

	// 镜像id
	ImageId *int64 `json:"ImageId,omitempty" name:"ImageId"`

	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 工作空间版本
	WorkspaceVersion *int64 `json:"WorkspaceVersion,omitempty" name:"WorkspaceVersion"`

	// 工作空间资源结构
	WorkspaceResourceDTO *WorkspaceResourceDTO `json:"WorkspaceResourceDTO,omitempty" name:"WorkspaceResourceDTO"`

	// 代码仓库地址
	VersionControlUrl *string `json:"VersionControlUrl,omitempty" name:"VersionControlUrl"`

	// 代码Ref是分支还是标签
	VersionControlRef *string `json:"VersionControlRef,omitempty" name:"VersionControlRef"`

	// 代码Ref地址
	VersionControlRefType *string `json:"VersionControlRefType,omitempty" name:"VersionControlRefType"`

	// 快照Uid
	SnapshotUid *string `json:"SnapshotUid,omitempty" name:"SnapshotUid"`

	// 模板id
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 价格id
	PriceId *int64 `json:"PriceId,omitempty" name:"PriceId"`

	// 初始化状态
	InitializeStatus *int64 `json:"InitializeStatus,omitempty" name:"InitializeStatus"`

	// 描述
	VersionControlDesc *string `json:"VersionControlDesc,omitempty" name:"VersionControlDesc"`
}

type WorkspaceInfo struct {
	// 工作空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" name:"WorkspaceId"`

	// 工作空间标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpaceKey *string `json:"SpaceKey,omitempty" name:"SpaceKey"`
}

type WorkspaceInfoDTO struct {
	// 工作空间创建时间
	CreateDate *string `json:"CreateDate,omitempty" name:"CreateDate"`

	// 空间key
	SpaceKey *string `json:"SpaceKey,omitempty" name:"SpaceKey"`

	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" name:"WorkspaceId"`
}

type WorkspaceResourceDTO struct {
	// CPU核心数
	CpuCoreNumber *uint64 `json:"CpuCoreNumber,omitempty" name:"CpuCoreNumber"`

	// 一般内存
	NormalMemory *uint64 `json:"NormalMemory,omitempty" name:"NormalMemory"`

	// 系统存储
	SystemStorage *uint64 `json:"SystemStorage,omitempty" name:"SystemStorage"`

	// 用户存储
	UserStorage *uint64 `json:"UserStorage,omitempty" name:"UserStorage"`

	// GPU数
	GpuNumber *uint64 `json:"GpuNumber,omitempty" name:"GpuNumber"`

	// GPU内存
	GpuMemory *uint64 `json:"GpuMemory,omitempty" name:"GpuMemory"`
}

type WorkspaceShareInfo struct {
	// 共享或不共享状态
	Status *bool `json:"Status,omitempty" name:"Status"`

	// 是否与我共享
	// 注意：此字段可能返回 null，表示取不到有效值。
	WithMe *bool `json:"WithMe,omitempty" name:"WithMe"`

	// 开始共享的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`

	// 停止共享的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 停止共享的时间
	Users []*UserInfoRsp `json:"Users,omitempty" name:"Users"`
}

type WorkspaceStatusInfo struct {
	// 空间ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 空间名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 所属人
	Owner *UserInfoRsp `json:"Owner,omitempty" name:"Owner"`

	// 空间标识
	SpaceKey *string `json:"SpaceKey,omitempty" name:"SpaceKey"`

	// 状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 最后操作时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastOpsDate *string `json:"LastOpsDate,omitempty" name:"LastOpsDate"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 共享状态
	Share *WorkspaceShareInfo `json:"Share,omitempty" name:"Share"`

	// 空间类型
	WorkspaceType *string `json:"WorkspaceType,omitempty" name:"WorkspaceType"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitempty" name:"Label"`

	// 空间版本
	WorkspaceVersion *int64 `json:"WorkspaceVersion,omitempty" name:"WorkspaceVersion"`

	// 图标地址
	ImageIcon *string `json:"ImageIcon,omitempty" name:"ImageIcon"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateDate *string `json:"CreateDate,omitempty" name:"CreateDate"`

	// 版本控制地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionControlUrl *string `json:"VersionControlUrl,omitempty" name:"VersionControlUrl"`

	// 版本控制描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionControlDesc *string `json:"VersionControlDesc,omitempty" name:"VersionControlDesc"`

	// 版本控制引用
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionControlRef *string `json:"VersionControlRef,omitempty" name:"VersionControlRef"`

	// 版本控制引用类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionControlRefType *string `json:"VersionControlRefType,omitempty" name:"VersionControlRefType"`

	// 版本控制类型
	VersionControlType *string `json:"VersionControlType,omitempty" name:"VersionControlType"`

	// 模板ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 快照ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SnapshotUid *string `json:"SnapshotUid,omitempty" name:"SnapshotUid"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecDesc *string `json:"SpecDesc,omitempty" name:"SpecDesc"`

	// CPU数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 内存
	// 注意：此字段可能返回 null，表示取不到有效值。
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
}

type WorkspaceTemplateInfo struct {
	// 模板ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 模板分类
	Category *string `json:"Category,omitempty" name:"Category"`

	// 模板名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 模板描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 中文描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	DescriptionEN *string `json:"DescriptionEN,omitempty" name:"DescriptionEN"`

	// 模板标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags *string `json:"Tags,omitempty" name:"Tags"`

	// 模板图标地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Icon *string `json:"Icon,omitempty" name:"Icon"`

	// 默认仓库类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionControlType *string `json:"VersionControlType,omitempty" name:"VersionControlType"`

	// 默认仓库地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionControlUrl *string `json:"VersionControlUrl,omitempty" name:"VersionControlUrl"`

	// 默认仓库描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionControlDesc *string `json:"VersionControlDesc,omitempty" name:"VersionControlDesc"`

	// 默认仓库所属人
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionControlOwner *string `json:"VersionControlOwner,omitempty" name:"VersionControlOwner"`

	// 默认仓库引用地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionControlRef *string `json:"VersionControlRef,omitempty" name:"VersionControlRef"`

	// 默认仓库引用类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionControlRefType *string `json:"VersionControlRefType,omitempty" name:"VersionControlRefType"`

	// 用户自定义仓库地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserVersionControlUrl *string `json:"UserVersionControlUrl,omitempty" name:"UserVersionControlUrl"`

	// 用户自定义仓库类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserVersionControlType *string `json:"UserVersionControlType,omitempty" name:"UserVersionControlType"`

	// 用户自定义仓库引用
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserVersionControlRef *string `json:"UserVersionControlRef,omitempty" name:"UserVersionControlRef"`

	// 用户自定义仓库引用类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserVersionControlRefType *string `json:"UserVersionControlRefType,omitempty" name:"UserVersionControlRefType"`

	// xxx
	// 注意：此字段可能返回 null，表示取不到有效值。
	DevFile *string `json:"DevFile,omitempty" name:"DevFile"`

	// xxx
	// 注意：此字段可能返回 null，表示取不到有效值。
	PluginFile *string `json:"PluginFile,omitempty" name:"PluginFile"`

	// xxx
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrebuildFile *string `json:"PrebuildFile,omitempty" name:"PrebuildFile"`

	// 是否标记
	Marked *bool `json:"Marked,omitempty" name:"Marked"`

	// 标记状态
	MarkAt *int64 `json:"MarkAt,omitempty" name:"MarkAt"`

	// 创建时间
	CreateDate *string `json:"CreateDate,omitempty" name:"CreateDate"`

	// 最后修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastModified *string `json:"LastModified,omitempty" name:"LastModified"`

	// 编号
	Sort *int64 `json:"Sort,omitempty" name:"Sort"`

	// xxx
	// 注意：此字段可能返回 null，表示取不到有效值。
	SnapshotUid *string `json:"SnapshotUid,omitempty" name:"SnapshotUid"`

	// 用户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *int64 `json:"UserId,omitempty" name:"UserId"`

	// 用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Author *string `json:"Author,omitempty" name:"Author"`

	// 是否属于当前用户
	Me *bool `json:"Me,omitempty" name:"Me"`

	// xxx
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthorAvatar *string `json:"AuthorAvatar,omitempty" name:"AuthorAvatar"`
}

type WorkspaceTokenDTO struct {
	// 工作空间 SpaceKey
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpaceKey *string `json:"SpaceKey,omitempty" name:"SpaceKey"`

	// token过期时间，单位是秒，默认 3600
	// 注意：此字段可能返回 null，表示取不到有效值。
	TokenExpiredLimitSec *uint64 `json:"TokenExpiredLimitSec,omitempty" name:"TokenExpiredLimitSec"`
}

type WorkspaceTokenInfoV0 struct {
	// 访问工作空间临时凭证
	Token *string `json:"Token,omitempty" name:"Token"`

	// token 过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpiredTime *string `json:"ExpiredTime,omitempty" name:"ExpiredTime"`
}