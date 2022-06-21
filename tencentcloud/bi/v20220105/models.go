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

package v20220105

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type ApplyEmbedIntervalRequestParams struct {
	// 分享项目id，必选
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 分享页面id，嵌出看板时此为空值0
	PageId *uint64 `json:"PageId,omitempty" name:"PageId"`

	// 需要申请延期的Token
	BIToken *string `json:"BIToken,omitempty" name:"BIToken"`

	// 备用字段
	ExtraParam *string `json:"ExtraParam,omitempty" name:"ExtraParam"`

	// panel,看板；page，页面
	Scope *string `json:"Scope,omitempty" name:"Scope"`
}

type ApplyEmbedIntervalRequest struct {
	*tchttp.BaseRequest
	
	// 分享项目id，必选
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 分享页面id，嵌出看板时此为空值0
	PageId *uint64 `json:"PageId,omitempty" name:"PageId"`

	// 需要申请延期的Token
	BIToken *string `json:"BIToken,omitempty" name:"BIToken"`

	// 备用字段
	ExtraParam *string `json:"ExtraParam,omitempty" name:"ExtraParam"`

	// panel,看板；page，页面
	Scope *string `json:"Scope,omitempty" name:"Scope"`
}

func (r *ApplyEmbedIntervalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyEmbedIntervalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "PageId")
	delete(f, "BIToken")
	delete(f, "ExtraParam")
	delete(f, "Scope")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyEmbedIntervalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyEmbedIntervalResponseParams struct {
	// 额外参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitempty" name:"Extra"`

	// 结果数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *ApplyEmbedTokenInfo `json:"Data,omitempty" name:"Data"`

	// 结果描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ApplyEmbedIntervalResponse struct {
	*tchttp.BaseResponse
	Response *ApplyEmbedIntervalResponseParams `json:"Response"`
}

func (r *ApplyEmbedIntervalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyEmbedIntervalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplyEmbedTokenInfo struct {
	// 申请结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`
}

// Predefined struct for user
type CreateEmbedTokenRequestParams struct {
	// 分享项目id，必选
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 分享页面id，嵌出看板时此为空值0
	PageId *uint64 `json:"PageId,omitempty" name:"PageId"`

	// page表示嵌出页面，panel表嵌出整个看板
	Scope *string `json:"Scope,omitempty" name:"Scope"`

	// 过期时间。 单位：分钟 最大值：240。即，4小时 默认值：240
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 备用字段
	ExtraParam *string `json:"ExtraParam,omitempty" name:"ExtraParam"`
}

type CreateEmbedTokenRequest struct {
	*tchttp.BaseRequest
	
	// 分享项目id，必选
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 分享页面id，嵌出看板时此为空值0
	PageId *uint64 `json:"PageId,omitempty" name:"PageId"`

	// page表示嵌出页面，panel表嵌出整个看板
	Scope *string `json:"Scope,omitempty" name:"Scope"`

	// 过期时间。 单位：分钟 最大值：240。即，4小时 默认值：240
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 备用字段
	ExtraParam *string `json:"ExtraParam,omitempty" name:"ExtraParam"`
}

func (r *CreateEmbedTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEmbedTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "PageId")
	delete(f, "Scope")
	delete(f, "ExpireTime")
	delete(f, "ExtraParam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEmbedTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEmbedTokenResponseParams struct {
	// 额外信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitempty" name:"Extra"`

	// 数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *EmbedTokenInfo `json:"Data,omitempty" name:"Data"`

	// 结果描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateEmbedTokenResponse struct {
	*tchttp.BaseResponse
	Response *CreateEmbedTokenResponseParams `json:"Response"`
}

func (r *CreateEmbedTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEmbedTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EmbedTokenInfo struct {
	// 信息标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 令牌
	// 注意：此字段可能返回 null，表示取不到有效值。
	BIToken *string `json:"BIToken,omitempty" name:"BIToken"`

	// 项目Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 创建人
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedUser *string `json:"CreatedUser,omitempty" name:"CreatedUser"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 更新人
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedUser *string `json:"UpdatedUser,omitempty" name:"UpdatedUser"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`

	// 页面Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageId *string `json:"PageId,omitempty" name:"PageId"`

	// 备用
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraParam *string `json:"ExtraParam,omitempty" name:"ExtraParam"`

	// 嵌出类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Scope *string `json:"Scope,omitempty" name:"Scope"`

	// 过期时间，分钟为单位，最大240
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *uint64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
}