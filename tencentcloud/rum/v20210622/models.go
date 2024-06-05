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

package v20210622

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type CreateProjectRequestParams struct {
	// 应用名称(不为空且最长为 200)
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 业务系统 ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 项目抽样率(大于等于 0)
	Rate *string `json:"Rate,omitnil,omitempty" name:"Rate"`

	// 是否开启聚类
	EnableURLGroup *uint64 `json:"EnableURLGroup,omitnil,omitempty" name:"EnableURLGroup"`

	// 项目类型("web", "mp", "android", "ios", "node", "hippy", "weex", "viola", "rn")
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 项目对应仓库地址(可选，最长为 256)
	Repo *string `json:"Repo,omitnil,omitempty" name:"Repo"`

	// 项目对应网页地址(可选，最长为 256)
	URL *string `json:"URL,omitnil,omitempty" name:"URL"`

	// 应用描述(可选，最长为 1000)
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`
}

type CreateProjectRequest struct {
	*tchttp.BaseRequest
	
	// 应用名称(不为空且最长为 200)
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 业务系统 ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 项目抽样率(大于等于 0)
	Rate *string `json:"Rate,omitnil,omitempty" name:"Rate"`

	// 是否开启聚类
	EnableURLGroup *uint64 `json:"EnableURLGroup,omitnil,omitempty" name:"EnableURLGroup"`

	// 项目类型("web", "mp", "android", "ios", "node", "hippy", "weex", "viola", "rn")
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 项目对应仓库地址(可选，最长为 256)
	Repo *string `json:"Repo,omitnil,omitempty" name:"Repo"`

	// 项目对应网页地址(可选，最长为 256)
	URL *string `json:"URL,omitnil,omitempty" name:"URL"`

	// 应用描述(可选，最长为 1000)
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`
}

func (r *CreateProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "InstanceID")
	delete(f, "Rate")
	delete(f, "EnableURLGroup")
	delete(f, "Type")
	delete(f, "Repo")
	delete(f, "URL")
	delete(f, "Desc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProjectResponseParams struct {
	// 项目 id
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 项目唯一key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateProjectResponse struct {
	*tchttp.BaseResponse
	Response *CreateProjectResponseParams `json:"Response"`
}

func (r *CreateProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReleaseFileRequestParams struct {
	// 项目 id
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`

	// 文件信息列表
	Files []*ReleaseFile `json:"Files,omitnil,omitempty" name:"Files"`
}

type CreateReleaseFileRequest struct {
	*tchttp.BaseRequest
	
	// 项目 id
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`

	// 文件信息列表
	Files []*ReleaseFile `json:"Files,omitnil,omitempty" name:"Files"`
}

func (r *CreateReleaseFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReleaseFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectID")
	delete(f, "Files")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateReleaseFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReleaseFileResponseParams struct {
	// 调用结果
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateReleaseFileResponse struct {
	*tchttp.BaseResponse
	Response *CreateReleaseFileResponseParams `json:"Response"`
}

func (r *CreateReleaseFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReleaseFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStarProjectRequestParams struct {
	// 实例ID：taw-123
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

type CreateStarProjectRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID：taw-123
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

func (r *CreateStarProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStarProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStarProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStarProjectResponseParams struct {
	// 接口返回信息
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateStarProjectResponse struct {
	*tchttp.BaseResponse
	Response *CreateStarProjectResponseParams `json:"Response"`
}

func (r *CreateStarProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStarProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTawInstanceRequestParams struct {
	// 片区Id，(至少大于0)
	AreaId *int64 `json:"AreaId,omitnil,omitempty" name:"AreaId"`

	// 计费类型, (1=后付费)
	ChargeType *int64 `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 数据保存时间，(至少大于0)
	DataRetentionDays *int64 `json:"DataRetentionDays,omitnil,omitempty" name:"DataRetentionDays"`

	// 实例名称，(最大长度不超过255字节)
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 实例描述，(最大长度不超过1024字节)
	InstanceDesc *string `json:"InstanceDesc,omitnil,omitempty" name:"InstanceDesc"`

	// 每天数据上报量，（不作量级限制）
	CountNum *string `json:"CountNum,omitnil,omitempty" name:"CountNum"`

	// 数据存储时长计费
	PeriodRetain *string `json:"PeriodRetain,omitnil,omitempty" name:"PeriodRetain"`

	// 实例购买渠道("cdn" 等)
	BuyingChannel *string `json:"BuyingChannel,omitnil,omitempty" name:"BuyingChannel"`

	// 预付费资源包类型(仅预付费需要)
	ResourcePackageType *uint64 `json:"ResourcePackageType,omitnil,omitempty" name:"ResourcePackageType"`

	// 预付费资源包数量(仅预付费需要)
	ResourcePackageNum *uint64 `json:"ResourcePackageNum,omitnil,omitempty" name:"ResourcePackageNum"`

	// 实例类型 1:原web相关类型 2:app端类型
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`
}

type CreateTawInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 片区Id，(至少大于0)
	AreaId *int64 `json:"AreaId,omitnil,omitempty" name:"AreaId"`

	// 计费类型, (1=后付费)
	ChargeType *int64 `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 数据保存时间，(至少大于0)
	DataRetentionDays *int64 `json:"DataRetentionDays,omitnil,omitempty" name:"DataRetentionDays"`

	// 实例名称，(最大长度不超过255字节)
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 实例描述，(最大长度不超过1024字节)
	InstanceDesc *string `json:"InstanceDesc,omitnil,omitempty" name:"InstanceDesc"`

	// 每天数据上报量，（不作量级限制）
	CountNum *string `json:"CountNum,omitnil,omitempty" name:"CountNum"`

	// 数据存储时长计费
	PeriodRetain *string `json:"PeriodRetain,omitnil,omitempty" name:"PeriodRetain"`

	// 实例购买渠道("cdn" 等)
	BuyingChannel *string `json:"BuyingChannel,omitnil,omitempty" name:"BuyingChannel"`

	// 预付费资源包类型(仅预付费需要)
	ResourcePackageType *uint64 `json:"ResourcePackageType,omitnil,omitempty" name:"ResourcePackageType"`

	// 预付费资源包数量(仅预付费需要)
	ResourcePackageNum *uint64 `json:"ResourcePackageNum,omitnil,omitempty" name:"ResourcePackageNum"`

	// 实例类型 1:原web相关类型 2:app端类型
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`
}

func (r *CreateTawInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTawInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AreaId")
	delete(f, "ChargeType")
	delete(f, "DataRetentionDays")
	delete(f, "InstanceName")
	delete(f, "Tags")
	delete(f, "InstanceDesc")
	delete(f, "CountNum")
	delete(f, "PeriodRetain")
	delete(f, "BuyingChannel")
	delete(f, "ResourcePackageType")
	delete(f, "ResourcePackageNum")
	delete(f, "InstanceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTawInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTawInstanceResponseParams struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 预付费订单 id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTawInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateTawInstanceResponseParams `json:"Response"`
}

func (r *CreateTawInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTawInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWhitelistRequestParams struct {
	// 实例ID：taw-123
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 备注（暂未作字节数限制）
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// uin：业务方标识
	WhitelistUin *string `json:"WhitelistUin,omitnil,omitempty" name:"WhitelistUin"`

	// 业务方标识
	Aid *string `json:"Aid,omitnil,omitempty" name:"Aid"`
}

type CreateWhitelistRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID：taw-123
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 备注（暂未作字节数限制）
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// uin：业务方标识
	WhitelistUin *string `json:"WhitelistUin,omitnil,omitempty" name:"WhitelistUin"`

	// 业务方标识
	Aid *string `json:"Aid,omitnil,omitempty" name:"Aid"`
}

func (r *CreateWhitelistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWhitelistRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	delete(f, "Remark")
	delete(f, "WhitelistUin")
	delete(f, "Aid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWhitelistRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWhitelistResponseParams struct {
	// 消息
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 白名单ID
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateWhitelistResponse struct {
	*tchttp.BaseResponse
	Response *CreateWhitelistResponseParams `json:"Response"`
}

func (r *CreateWhitelistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWhitelistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstanceRequestParams struct {
	// 需要删除的实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DeleteInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 需要删除的实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DeleteInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteInstanceResponseParams `json:"Response"`
}

func (r *DeleteInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProjectRequestParams struct {
	// 需要删除的项目 ID
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`
}

type DeleteProjectRequest struct {
	*tchttp.BaseRequest
	
	// 需要删除的项目 ID
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`
}

func (r *DeleteProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProjectResponseParams struct {
	// 操作信息
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteProjectResponse struct {
	*tchttp.BaseResponse
	Response *DeleteProjectResponseParams `json:"Response"`
}

func (r *DeleteProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReleaseFileRequestParams struct {
	// 文件 id
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

type DeleteReleaseFileRequest struct {
	*tchttp.BaseRequest
	
	// 文件 id
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

func (r *DeleteReleaseFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReleaseFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteReleaseFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReleaseFileResponseParams struct {
	// 接口请求返回字符串
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteReleaseFileResponse struct {
	*tchttp.BaseResponse
	Response *DeleteReleaseFileResponseParams `json:"Response"`
}

func (r *DeleteReleaseFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReleaseFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStarProjectRequestParams struct {
	// 实例ID：taw-123
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

type DeleteStarProjectRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID：taw-123
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

func (r *DeleteStarProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStarProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteStarProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStarProjectResponseParams struct {
	// 返回消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteStarProjectResponse struct {
	*tchttp.BaseResponse
	Response *DeleteStarProjectResponseParams `json:"Response"`
}

func (r *DeleteStarProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStarProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWhitelistRequestParams struct {
	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 名单ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`
}

type DeleteWhitelistRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 名单ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`
}

func (r *DeleteWhitelistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWhitelistRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWhitelistRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWhitelistResponseParams struct {
	// 消息success
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteWhitelistResponse struct {
	*tchttp.BaseResponse
	Response *DeleteWhitelistResponseParams `json:"Response"`
}

func (r *DeleteWhitelistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWhitelistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppDimensionMetricsRequestParams struct {
	// app 项目ID
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`

	// 查询的表名
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 查询指标 fields
	Fields *string `json:"Fields,omitnil,omitempty" name:"Fields"`

	// 查询的过滤条件
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 查询简单过滤条件
	FilterSimple *string `json:"FilterSimple,omitnil,omitempty" name:"FilterSimple"`

	// group by 条件
	GroupBy []*string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// order by 条件
	OrderBy []*string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// limit 参数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// offset 参数
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 业务上下文参数
	BusinessContext *string `json:"BusinessContext,omitnil,omitempty" name:"BusinessContext"`
}

type DescribeAppDimensionMetricsRequest struct {
	*tchttp.BaseRequest
	
	// app 项目ID
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`

	// 查询的表名
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 查询指标 fields
	Fields *string `json:"Fields,omitnil,omitempty" name:"Fields"`

	// 查询的过滤条件
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 查询简单过滤条件
	FilterSimple *string `json:"FilterSimple,omitnil,omitempty" name:"FilterSimple"`

	// group by 条件
	GroupBy []*string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// order by 条件
	OrderBy []*string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// limit 参数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// offset 参数
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 业务上下文参数
	BusinessContext *string `json:"BusinessContext,omitnil,omitempty" name:"BusinessContext"`
}

func (r *DescribeAppDimensionMetricsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppDimensionMetricsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectID")
	delete(f, "From")
	delete(f, "Fields")
	delete(f, "Filter")
	delete(f, "FilterSimple")
	delete(f, "GroupBy")
	delete(f, "OrderBy")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "BusinessContext")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAppDimensionMetricsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppDimensionMetricsResponseParams struct {
	// 查询数据返回
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAppDimensionMetricsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAppDimensionMetricsResponseParams `json:"Response"`
}

func (r *DescribeAppDimensionMetricsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppDimensionMetricsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppMetricsDataRequestParams struct {
	// app 项目ID
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`

	// 查询的表名
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 查询指标 field
	Fields *string `json:"Fields,omitnil,omitempty" name:"Fields"`

	// 查询的过滤条件
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 查询简单过滤条件
	FilterSimple *string `json:"FilterSimple,omitnil,omitempty" name:"FilterSimple"`

	// group by 条件
	GroupBy []*string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// order by 条件
	OrderBy []*string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// limit 参数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// offset 参数
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// group by 参数
	GroupByModifier *string `json:"GroupByModifier,omitnil,omitempty" name:"GroupByModifier"`
}

type DescribeAppMetricsDataRequest struct {
	*tchttp.BaseRequest
	
	// app 项目ID
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`

	// 查询的表名
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 查询指标 field
	Fields *string `json:"Fields,omitnil,omitempty" name:"Fields"`

	// 查询的过滤条件
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 查询简单过滤条件
	FilterSimple *string `json:"FilterSimple,omitnil,omitempty" name:"FilterSimple"`

	// group by 条件
	GroupBy []*string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// order by 条件
	OrderBy []*string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// limit 参数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// offset 参数
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// group by 参数
	GroupByModifier *string `json:"GroupByModifier,omitnil,omitempty" name:"GroupByModifier"`
}

func (r *DescribeAppMetricsDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppMetricsDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectID")
	delete(f, "From")
	delete(f, "Fields")
	delete(f, "Filter")
	delete(f, "FilterSimple")
	delete(f, "GroupBy")
	delete(f, "OrderBy")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "GroupByModifier")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAppMetricsDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppMetricsDataResponseParams struct {
	// 查询数据返回
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAppMetricsDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAppMetricsDataResponseParams `json:"Response"`
}

func (r *DescribeAppMetricsDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppMetricsDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppSingleCaseDetailListRequestParams struct {
	// app 项目ID
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`

	// 查询的表名
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 查询指标 field
	Fields *string `json:"Fields,omitnil,omitempty" name:"Fields"`

	// 查询的过滤条件
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 查询简单过滤条件
	FilterSimple *string `json:"FilterSimple,omitnil,omitempty" name:"FilterSimple"`

	// group by 条件
	GroupBy []*string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// order by 条件
	OrderBy []*string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// limit 参数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// offset 参数
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeAppSingleCaseDetailListRequest struct {
	*tchttp.BaseRequest
	
	// app 项目ID
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`

	// 查询的表名
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 查询指标 field
	Fields *string `json:"Fields,omitnil,omitempty" name:"Fields"`

	// 查询的过滤条件
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 查询简单过滤条件
	FilterSimple *string `json:"FilterSimple,omitnil,omitempty" name:"FilterSimple"`

	// group by 条件
	GroupBy []*string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// order by 条件
	OrderBy []*string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// limit 参数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// offset 参数
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeAppSingleCaseDetailListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppSingleCaseDetailListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectID")
	delete(f, "From")
	delete(f, "Fields")
	delete(f, "Filter")
	delete(f, "FilterSimple")
	delete(f, "GroupBy")
	delete(f, "OrderBy")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAppSingleCaseDetailListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppSingleCaseDetailListResponseParams struct {
	// 查询数据返回
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAppSingleCaseDetailListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAppSingleCaseDetailListResponseParams `json:"Response"`
}

func (r *DescribeAppSingleCaseDetailListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppSingleCaseDetailListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppSingleCaseListRequestParams struct {
	// app 项目 ID
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`

	// 查询的表名
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 查询指标 field
	Fields *string `json:"Fields,omitnil,omitempty" name:"Fields"`

	// 查询的过滤条件
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 查询简单过滤条件
	FilterSimple *string `json:"FilterSimple,omitnil,omitempty" name:"FilterSimple"`

	// group by 条件
	GroupBy []*string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// order by 条件
	OrderBy []*string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// limit 参数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// offset 参数
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeAppSingleCaseListRequest struct {
	*tchttp.BaseRequest
	
	// app 项目 ID
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`

	// 查询的表名
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 查询指标 field
	Fields *string `json:"Fields,omitnil,omitempty" name:"Fields"`

	// 查询的过滤条件
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 查询简单过滤条件
	FilterSimple *string `json:"FilterSimple,omitnil,omitempty" name:"FilterSimple"`

	// group by 条件
	GroupBy []*string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// order by 条件
	OrderBy []*string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// limit 参数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// offset 参数
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeAppSingleCaseListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppSingleCaseListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectID")
	delete(f, "From")
	delete(f, "Fields")
	delete(f, "Filter")
	delete(f, "FilterSimple")
	delete(f, "GroupBy")
	delete(f, "OrderBy")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAppSingleCaseListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppSingleCaseListResponseParams struct {
	// 查询数据返回
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAppSingleCaseListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAppSingleCaseListResponseParams `json:"Response"`
}

func (r *DescribeAppSingleCaseListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppSingleCaseListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataCustomUrlRequestParams struct {
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// top：资源top视图，allcount：性能视图，day：14天数据，condition：条件列表，pagepv：性能视图，area：请求速度分布，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图/ext1视图等等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型（1,2,3,4,5,100），1表示WIFI, 2表示2G, 3表示3G, 4表示4G, 5表示5G, 6表示6G, 100表示未知。
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时计算方式
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 自定义测速的key的值
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

type DescribeDataCustomUrlRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// top：资源top视图，allcount：性能视图，day：14天数据，condition：条件列表，pagepv：性能视图，area：请求速度分布，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图/ext1视图等等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型（1,2,3,4,5,100），1表示WIFI, 2表示2G, 3表示3G, 4表示4G, 5表示5G, 6表示6G, 100表示未知。
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时计算方式
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 自定义测速的key的值
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

func (r *DescribeDataCustomUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataCustomUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "CostType")
	delete(f, "Url")
	delete(f, "Env")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataCustomUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataCustomUrlResponseParams struct {
	// 返回值
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataCustomUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataCustomUrlResponseParams `json:"Response"`
}

func (r *DescribeDataCustomUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataCustomUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataEventUrlRequestParams struct {
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// allcount：性能视图，day：14天数据，condition：条件列表，ckuv：获取uv趋势，ckpv：获取pv趋势，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图/ext1视图等等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 筛选条件
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

type DescribeDataEventUrlRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// allcount：性能视图，day：14天数据，condition：条件列表，ckuv：获取uv趋势，ckpv：获取pv趋势，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图/ext1视图等等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 筛选条件
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

func (r *DescribeDataEventUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataEventUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "Name")
	delete(f, "Env")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataEventUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataEventUrlResponseParams struct {
	// 返回值
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataEventUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataEventUrlResponseParams `json:"Response"`
}

func (r *DescribeDataEventUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataEventUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataFetchProjectRequestParams struct {
	// 开始时间，示例值：1625454840
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// allcount：性能视图，day：14天数据，condition：条件列表，area：请求速度分布，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图/ext1视图等等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间，示例值：1625454840
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级（1表示白名单日志，2表示一般日志，4表示错误日志，8表示Promise 错误，16表示Ajax 请求异常，32表示JS 加载异常，64表示图片加载异常，128表示css 加载异常，256表示console.error，512表示音视频资源异常，1024表示retcode 异常，2048表示aegis report，4096表示PV日志，8192表示自定义事件，16384表示小程序 页面不存在，32768表示websocket错误，65536表示js bridge错误）
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型（1,2,3,4,5,100），1表示WIFI, 2表示2G, 3表示3G, 4表示4G, 5表示5G, 6表示6G, 100表示未知。
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时计算方式
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 来源
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// httpcode响应码
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// retcode
	Ret *string `json:"Ret,omitnil,omitempty" name:"Ret"`
}

type DescribeDataFetchProjectRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间，示例值：1625454840
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// allcount：性能视图，day：14天数据，condition：条件列表，area：请求速度分布，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图/ext1视图等等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间，示例值：1625454840
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级（1表示白名单日志，2表示一般日志，4表示错误日志，8表示Promise 错误，16表示Ajax 请求异常，32表示JS 加载异常，64表示图片加载异常，128表示css 加载异常，256表示console.error，512表示音视频资源异常，1024表示retcode 异常，2048表示aegis report，4096表示PV日志，8192表示自定义事件，16384表示小程序 页面不存在，32768表示websocket错误，65536表示js bridge错误）
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型（1,2,3,4,5,100），1表示WIFI, 2表示2G, 3表示3G, 4表示4G, 5表示5G, 6表示6G, 100表示未知。
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时计算方式
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 来源
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// httpcode响应码
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// retcode
	Ret *string `json:"Ret,omitnil,omitempty" name:"Ret"`
}

func (r *DescribeDataFetchProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataFetchProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "CostType")
	delete(f, "Url")
	delete(f, "Env")
	delete(f, "Status")
	delete(f, "Ret")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataFetchProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataFetchProjectResponseParams struct {
	// 返回值
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataFetchProjectResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataFetchProjectResponseParams `json:"Response"`
}

func (r *DescribeDataFetchProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataFetchProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataFetchUrlInfoRequestParams struct {
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时计算方式
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 来源
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

type DescribeDataFetchUrlInfoRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时计算方式
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 来源
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

func (r *DescribeDataFetchUrlInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataFetchUrlInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "CostType")
	delete(f, "Url")
	delete(f, "Env")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataFetchUrlInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataFetchUrlInfoResponseParams struct {
	// 返回值
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataFetchUrlInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataFetchUrlInfoResponseParams `json:"Response"`
}

func (r *DescribeDataFetchUrlInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataFetchUrlInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataFetchUrlRequestParams struct {
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// allcount：性能视图，pagepv：pv视图，day：14天数据，count40x：40X视图，count50x：50X视图，count5xand4x：40∑50视图，top：资源top视图，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图/ext1视图等等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时计算方式
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 来源
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// httpcode响应码
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// retcode
	Ret *string `json:"Ret,omitnil,omitempty" name:"Ret"`

	// 网络状态
	NetStatus *string `json:"NetStatus,omitnil,omitempty" name:"NetStatus"`
}

type DescribeDataFetchUrlRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// allcount：性能视图，pagepv：pv视图，day：14天数据，count40x：40X视图，count50x：50X视图，count5xand4x：40∑50视图，top：资源top视图，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图/ext1视图等等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时计算方式
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 来源
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// httpcode响应码
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// retcode
	Ret *string `json:"Ret,omitnil,omitempty" name:"Ret"`

	// 网络状态
	NetStatus *string `json:"NetStatus,omitnil,omitempty" name:"NetStatus"`
}

func (r *DescribeDataFetchUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataFetchUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "CostType")
	delete(f, "Url")
	delete(f, "Env")
	delete(f, "Status")
	delete(f, "Ret")
	delete(f, "NetStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataFetchUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataFetchUrlResponseParams struct {
	// 返回值
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataFetchUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataFetchUrlResponseParams `json:"Response"`
}

func (r *DescribeDataFetchUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataFetchUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataLogUrlInfoRequestParams struct {
	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 时间戳
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 时间戳
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeDataLogUrlInfoRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 时间戳
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 时间戳
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *DescribeDataLogUrlInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataLogUrlInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ID")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataLogUrlInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataLogUrlInfoResponseParams struct {
	// 返回字符串
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataLogUrlInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataLogUrlInfoResponseParams `json:"Response"`
}

func (r *DescribeDataLogUrlInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataLogUrlInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataLogUrlStatisticsRequestParams struct {
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// analysis：异常分析，compare：异常列表对比，allcount：性能视图，condition：条件列表，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图/ext1视图等等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 环境区分
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

type DescribeDataLogUrlStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// analysis：异常分析，compare：异常列表对比，allcount：性能视图，condition：条件列表，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图/ext1视图等等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 环境区分
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

func (r *DescribeDataLogUrlStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataLogUrlStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "Env")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataLogUrlStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataLogUrlStatisticsResponseParams struct {
	// 返回值
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataLogUrlStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataLogUrlStatisticsResponseParams `json:"Response"`
}

func (r *DescribeDataLogUrlStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataLogUrlStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataPerformancePageRequestParams struct {
	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// pagepv：pv视图，allcount：性能视图，falls：页面加载瀑布图，samp：首屏时间，day：14天数据，nettype：网络/平台视图，performance：页面性能TOP视图，version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：ISP视图/地区视图/浏览器视图等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 网络类型
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 耗时计算方式
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 环境变量
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// 网络状态
	NetStatus *string `json:"NetStatus,omitnil,omitempty" name:"NetStatus"`
}

type DescribeDataPerformancePageRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// pagepv：pv视图，allcount：性能视图，falls：页面加载瀑布图，samp：首屏时间，day：14天数据，nettype：网络/平台视图，performance：页面性能TOP视图，version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：ISP视图/地区视图/浏览器视图等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 网络类型
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 耗时计算方式
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 环境变量
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// 网络状态
	NetStatus *string `json:"NetStatus,omitnil,omitempty" name:"NetStatus"`
}

func (r *DescribeDataPerformancePageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataPerformancePageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ID")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Type")
	delete(f, "Level")
	delete(f, "Isp")
	delete(f, "Area")
	delete(f, "NetType")
	delete(f, "Platform")
	delete(f, "Device")
	delete(f, "VersionNum")
	delete(f, "ExtFirst")
	delete(f, "ExtSecond")
	delete(f, "ExtThird")
	delete(f, "IsAbroad")
	delete(f, "Browser")
	delete(f, "Os")
	delete(f, "Engine")
	delete(f, "Brand")
	delete(f, "From")
	delete(f, "CostType")
	delete(f, "Env")
	delete(f, "NetStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataPerformancePageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataPerformancePageResponseParams struct {
	// 返回值
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataPerformancePageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataPerformancePageResponseParams `json:"Response"`
}

func (r *DescribeDataPerformancePageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataPerformancePageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataPvUrlInfoRequestParams struct {
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

type DescribeDataPvUrlInfoRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

func (r *DescribeDataPvUrlInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataPvUrlInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "Env")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataPvUrlInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataPvUrlInfoResponseParams struct {
	// 返回值
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataPvUrlInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataPvUrlInfoResponseParams `json:"Response"`
}

func (r *DescribeDataPvUrlInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataPvUrlInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataPvUrlStatisticsRequestParams struct {
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// allcount：性能视图，day：14天数据，vp：性能，ckuv：uv，ckpv：pv，condition：条件列表，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图/ext1视图等等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// group by 参数值枚举1:1m  2:5m  3:30m  4:1h 
	//  5:1d
	GroupByType *int64 `json:"GroupByType,omitnil,omitempty" name:"GroupByType"`

	// 1: 查询智研
	// 0: 走旧逻辑，已下线，勿使用
	IsNewData *int64 `json:"IsNewData,omitnil,omitempty" name:"IsNewData"`
}

type DescribeDataPvUrlStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// allcount：性能视图，day：14天数据，vp：性能，ckuv：uv，ckpv：pv，condition：条件列表，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图/ext1视图等等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// group by 参数值枚举1:1m  2:5m  3:30m  4:1h 
	//  5:1d
	GroupByType *int64 `json:"GroupByType,omitnil,omitempty" name:"GroupByType"`

	// 1: 查询智研
	// 0: 走旧逻辑，已下线，勿使用
	IsNewData *int64 `json:"IsNewData,omitnil,omitempty" name:"IsNewData"`
}

func (r *DescribeDataPvUrlStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataPvUrlStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "Env")
	delete(f, "GroupByType")
	delete(f, "IsNewData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataPvUrlStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataPvUrlStatisticsResponseParams struct {
	// 返回值
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataPvUrlStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataPvUrlStatisticsResponseParams `json:"Response"`
}

func (r *DescribeDataPvUrlStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataPvUrlStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataReportCountRequestParams struct {
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 上报类型（custom，event，log，miniProgramData，performance，pv，speed，webvitals）
	ReportType *string `json:"ReportType,omitnil,omitempty" name:"ReportType"`

	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
}

type DescribeDataReportCountRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 上报类型（custom，event，log，miniProgramData，performance，pv，speed，webvitals）
	ReportType *string `json:"ReportType,omitnil,omitempty" name:"ReportType"`

	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
}

func (r *DescribeDataReportCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataReportCountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ReportType")
	delete(f, "InstanceID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataReportCountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataReportCountResponseParams struct {
	// 返回值
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataReportCountResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataReportCountResponseParams `json:"Response"`
}

func (r *DescribeDataReportCountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataReportCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataRequestParams struct {
	// 查询字符串
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

type DescribeDataRequest struct {
	*tchttp.BaseRequest
	
	// 查询字符串
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

func (r *DescribeDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Query")
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataResponseParams struct {
	// 返回字符串
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataResponseParams `json:"Response"`
}

func (r *DescribeDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataSetUrlStatisticsRequestParams struct {
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// allcount：性能视图，data：小程序，component：小程序相关，day：14天数据，nettype：网络/平台视图，performance：页面性能TOP视图，version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：ISP视图/地区视图/浏览器视图等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时计算
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// 获取package
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`
}

type DescribeDataSetUrlStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// allcount：性能视图，data：小程序，component：小程序相关，day：14天数据，nettype：网络/平台视图，performance：页面性能TOP视图，version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：ISP视图/地区视图/浏览器视图等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时计算
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// 获取package
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`
}

func (r *DescribeDataSetUrlStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataSetUrlStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "CostType")
	delete(f, "Env")
	delete(f, "PackageType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataSetUrlStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataSetUrlStatisticsResponseParams struct {
	// 返回值
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataSetUrlStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataSetUrlStatisticsResponseParams `json:"Response"`
}

func (r *DescribeDataSetUrlStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataSetUrlStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataStaticProjectRequestParams struct {
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// allcount：性能视图，day：14天数据，condition：条件列表，area：请求速度分布，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图/ext1视图等等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级（1表示白名单日志，2表示一般日志，4表示错误日志，8表示Promise 错误，16表示Ajax 请求异常，32表示JS 加载异常，64表示图片加载异常，128表示css 加载异常，256表示console.error，512表示音视频资源异常，1024表示retcode 异常，2048表示aegis report，4096表示PV日志，8192表示自定义事件，16384表示小程序 页面不存在，32768表示websocket错误，65536表示js bridge错误）
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型（1,2,3,4,5,100），1表示WIFI, 2表示2G, 3表示3G, 4表示4G, 5表示5G, 6表示6G, 100表示未知。
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时计算
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 来源
	Url []*string `json:"Url,omitnil,omitempty" name:"Url"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

type DescribeDataStaticProjectRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// allcount：性能视图，day：14天数据，condition：条件列表，area：请求速度分布，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图/ext1视图等等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级（1表示白名单日志，2表示一般日志，4表示错误日志，8表示Promise 错误，16表示Ajax 请求异常，32表示JS 加载异常，64表示图片加载异常，128表示css 加载异常，256表示console.error，512表示音视频资源异常，1024表示retcode 异常，2048表示aegis report，4096表示PV日志，8192表示自定义事件，16384表示小程序 页面不存在，32768表示websocket错误，65536表示js bridge错误）
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型（1,2,3,4,5,100），1表示WIFI, 2表示2G, 3表示3G, 4表示4G, 5表示5G, 6表示6G, 100表示未知。
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时计算
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 来源
	Url []*string `json:"Url,omitnil,omitempty" name:"Url"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

func (r *DescribeDataStaticProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataStaticProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "CostType")
	delete(f, "Url")
	delete(f, "Env")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataStaticProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataStaticProjectResponseParams struct {
	// 返回值
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataStaticProjectResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataStaticProjectResponseParams `json:"Response"`
}

func (r *DescribeDataStaticProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataStaticProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataStaticResourceRequestParams struct {
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// top：资源top视图，count40x：40X视图，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图//ext1视图等等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时计算方式
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 来源
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

type DescribeDataStaticResourceRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// top：资源top视图，count40x：40X视图，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图//ext1视图等等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时计算方式
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 来源
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

func (r *DescribeDataStaticResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataStaticResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "CostType")
	delete(f, "Url")
	delete(f, "Env")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataStaticResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataStaticResourceResponseParams struct {
	// 返回值
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataStaticResourceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataStaticResourceResponseParams `json:"Response"`
}

func (r *DescribeDataStaticResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataStaticResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataStaticUrlRequestParams struct {
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// pagepv：性能视图，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图//ext1视图等等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时计算方式
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 来源
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

type DescribeDataStaticUrlRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// pagepv：性能视图，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图//ext1视图等等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时计算方式
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 来源
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

func (r *DescribeDataStaticUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataStaticUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "CostType")
	delete(f, "Url")
	delete(f, "Env")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataStaticUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataStaticUrlResponseParams struct {
	// 返回值
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataStaticUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataStaticUrlResponseParams `json:"Response"`
}

func (r *DescribeDataStaticUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataStaticUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataWebVitalsPageRequestParams struct {
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 类型暂无
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时计算
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

type DescribeDataWebVitalsPageRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 类型暂无
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时计算
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

func (r *DescribeDataWebVitalsPageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataWebVitalsPageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Type")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "CostType")
	delete(f, "Env")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataWebVitalsPageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataWebVitalsPageResponseParams struct {
	// 返回值
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataWebVitalsPageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataWebVitalsPageResponseParams `json:"Response"`
}

func (r *DescribeDataWebVitalsPageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataWebVitalsPageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeErrorRequestParams struct {
	// 日期
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

type DescribeErrorRequest struct {
	*tchttp.BaseRequest
	
	// 日期
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

func (r *DescribeErrorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeErrorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Date")
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeErrorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeErrorResponseParams struct {
	// 内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeErrorResponse struct {
	*tchttp.BaseResponse
	Response *DescribeErrorResponseParams `json:"Response"`
}

func (r *DescribeErrorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeErrorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectLimitsRequestParams struct {
	// 项目ID
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`
}

type DescribeProjectLimitsRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`
}

func (r *DescribeProjectLimitsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectLimitsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectLimitsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectLimitsResponseParams struct {
	// 上报率数组列表
	ProjectLimitSet []*ProjectLimit `json:"ProjectLimitSet,omitnil,omitempty" name:"ProjectLimitSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProjectLimitsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProjectLimitsResponseParams `json:"Response"`
}

func (r *DescribeProjectLimitsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectLimitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectsRequestParams struct {
	// 分页每页数目，整型
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页页码，整型
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤参数；demo模式传{"Name": "IsDemo", "Values":["1"]}
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 该参数已废弃，demo模式请在Filters内注明
	IsDemo *int64 `json:"IsDemo,omitnil,omitempty" name:"IsDemo"`
}

type DescribeProjectsRequest struct {
	*tchttp.BaseRequest
	
	// 分页每页数目，整型
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页页码，整型
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤参数；demo模式传{"Name": "IsDemo", "Values":["1"]}
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 该参数已废弃，demo模式请在Filters内注明
	IsDemo *int64 `json:"IsDemo,omitnil,omitempty" name:"IsDemo"`
}

func (r *DescribeProjectsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "IsDemo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectsResponseParams struct {
	// 列表总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 项目列表
	ProjectSet []*RumProject `json:"ProjectSet,omitnil,omitempty" name:"ProjectSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProjectsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProjectsResponseParams `json:"Response"`
}

func (r *DescribeProjectsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePvListRequestParams struct {
	// ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 对PV指标的查询维度。获取day：d，   获取min则不填。
	Dimension *string `json:"Dimension,omitnil,omitempty" name:"Dimension"`
}

type DescribePvListRequest struct {
	*tchttp.BaseRequest
	
	// ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 对PV指标的查询维度。获取day：d，   获取min则不填。
	Dimension *string `json:"Dimension,omitnil,omitempty" name:"Dimension"`
}

func (r *DescribePvListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePvListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "EndTime")
	delete(f, "StartTime")
	delete(f, "Dimension")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePvListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePvListResponseParams struct {
	// pv列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectPvSet []*RumPvInfo `json:"ProjectPvSet,omitnil,omitempty" name:"ProjectPvSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePvListResponse struct {
	*tchttp.BaseResponse
	Response *DescribePvListResponseParams `json:"Response"`
}

func (r *DescribePvListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePvListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReleaseFileSignRequestParams struct {
	// 超时时间，不填默认是 5 分钟
	Timeout *int64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// bucket类型，不填默认1:web，2:app
	FileType *int64 `json:"FileType,omitnil,omitempty" name:"FileType"`
}

type DescribeReleaseFileSignRequest struct {
	*tchttp.BaseRequest
	
	// 超时时间，不填默认是 5 分钟
	Timeout *int64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// bucket类型，不填默认1:web，2:app
	FileType *int64 `json:"FileType,omitnil,omitempty" name:"FileType"`
}

func (r *DescribeReleaseFileSignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReleaseFileSignRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Timeout")
	delete(f, "FileType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReleaseFileSignRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReleaseFileSignResponseParams struct {
	// 临时密钥key
	SecretKey *string `json:"SecretKey,omitnil,omitempty" name:"SecretKey"`

	// 临时密钥 id
	SecretID *string `json:"SecretID,omitnil,omitempty" name:"SecretID"`

	// 临时密钥临时 token
	SessionToken *string `json:"SessionToken,omitnil,omitempty" name:"SessionToken"`

	// 开始时间戳
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 过期时间戳
	ExpiredTime *int64 `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeReleaseFileSignResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReleaseFileSignResponseParams `json:"Response"`
}

func (r *DescribeReleaseFileSignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReleaseFileSignResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReleaseFilesRequestParams struct {
	// 项目 id
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`

	// 文件版本
	FileVersion *string `json:"FileVersion,omitnil,omitempty" name:"FileVersion"`
}

type DescribeReleaseFilesRequest struct {
	*tchttp.BaseRequest
	
	// 项目 id
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`

	// 文件版本
	FileVersion *string `json:"FileVersion,omitnil,omitempty" name:"FileVersion"`
}

func (r *DescribeReleaseFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReleaseFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectID")
	delete(f, "FileVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReleaseFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReleaseFilesResponseParams struct {
	// 文件信息列表
	Files []*ReleaseFile `json:"Files,omitnil,omitempty" name:"Files"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeReleaseFilesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReleaseFilesResponseParams `json:"Response"`
}

func (r *DescribeReleaseFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReleaseFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRumGroupLogRequestParams struct {
	// 排序方式  desc  asc（必填）
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 开始时间（必填）
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 单次查询返回的原始日志条数，最大值为100（必填）
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 页数，第几页
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 查询语句，参考控制台请求参数，语句长度最大为4096（必填）
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 结束时间（必填）
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID（必填）
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 聚合字段
	GroupField *string `json:"GroupField,omitnil,omitempty" name:"GroupField"`
}

type DescribeRumGroupLogRequest struct {
	*tchttp.BaseRequest
	
	// 排序方式  desc  asc（必填）
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 开始时间（必填）
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 单次查询返回的原始日志条数，最大值为100（必填）
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 页数，第几页
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 查询语句，参考控制台请求参数，语句长度最大为4096（必填）
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 结束时间（必填）
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID（必填）
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 聚合字段
	GroupField *string `json:"GroupField,omitnil,omitempty" name:"GroupField"`
}

func (r *DescribeRumGroupLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRumGroupLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrderBy")
	delete(f, "StartTime")
	delete(f, "Limit")
	delete(f, "Page")
	delete(f, "Query")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "GroupField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRumGroupLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRumGroupLogResponseParams struct {
	// 返回字符串
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRumGroupLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRumGroupLogResponseParams `json:"Response"`
}

func (r *DescribeRumGroupLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRumGroupLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRumLogExportRequestParams struct {
	// 导出标识name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 开始时间（必填）
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询语句，参考控制台请求参数，语句长度最大为4096（必填）
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 结束时间（必填）
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID（必填）
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// field条件
	Fields []*string `json:"Fields,omitnil,omitempty" name:"Fields"`
}

type DescribeRumLogExportRequest struct {
	*tchttp.BaseRequest
	
	// 导出标识name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 开始时间（必填）
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询语句，参考控制台请求参数，语句长度最大为4096（必填）
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 结束时间（必填）
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID（必填）
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// field条件
	Fields []*string `json:"Fields,omitnil,omitempty" name:"Fields"`
}

func (r *DescribeRumLogExportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRumLogExportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "StartTime")
	delete(f, "Query")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "Fields")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRumLogExportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRumLogExportResponseParams struct {
	// 返回字符串
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRumLogExportResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRumLogExportResponseParams `json:"Response"`
}

func (r *DescribeRumLogExportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRumLogExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRumLogExportsRequestParams struct {
	// 页面大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页数，第几页
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 项目ID（必填）
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

type DescribeRumLogExportsRequest struct {
	*tchttp.BaseRequest
	
	// 页面大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页数，第几页
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 项目ID（必填）
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

func (r *DescribeRumLogExportsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRumLogExportsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageSize")
	delete(f, "PageNum")
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRumLogExportsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRumLogExportsResponseParams struct {
	// 返回字符串
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRumLogExportsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRumLogExportsResponseParams `json:"Response"`
}

func (r *DescribeRumLogExportsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRumLogExportsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRumLogListRequestParams struct {
	// 排序方式  desc  asc（必填）
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 开始时间（必填）格式为时间戳 毫秒
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 单次查询返回的原始日志条数，最大值为100（必填）
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 页数，第几页
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 查询语句，参考控制台请求参数，语句长度最大为4096（必填）
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 结束时间（必填）格式为时间戳 毫秒
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID（必填）
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

type DescribeRumLogListRequest struct {
	*tchttp.BaseRequest
	
	// 排序方式  desc  asc（必填）
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 开始时间（必填）格式为时间戳 毫秒
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 单次查询返回的原始日志条数，最大值为100（必填）
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 页数，第几页
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 查询语句，参考控制台请求参数，语句长度最大为4096（必填）
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 结束时间（必填）格式为时间戳 毫秒
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID（必填）
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

func (r *DescribeRumLogListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRumLogListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrderBy")
	delete(f, "StartTime")
	delete(f, "Limit")
	delete(f, "Page")
	delete(f, "Query")
	delete(f, "EndTime")
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRumLogListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRumLogListResponseParams struct {
	// 返回字符串
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRumLogListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRumLogListResponseParams `json:"Response"`
}

func (r *DescribeRumLogListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRumLogListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRumStatsLogListRequestParams struct {
	// 开始时间（必填）
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 单次查询返回的原始日志条数，最大值为100（必填）
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询语句，参考控制台请求参数，语句长度最大为4096（必填）
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 结束时间（必填）
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID（必填）
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

type DescribeRumStatsLogListRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间（必填）
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 单次查询返回的原始日志条数，最大值为100（必填）
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询语句，参考控制台请求参数，语句长度最大为4096（必填）
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 结束时间（必填）
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID（必填）
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

func (r *DescribeRumStatsLogListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRumStatsLogListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Limit")
	delete(f, "Query")
	delete(f, "EndTime")
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRumStatsLogListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRumStatsLogListResponseParams struct {
	// 返回字符串
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRumStatsLogListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRumStatsLogListResponseParams `json:"Response"`
}

func (r *DescribeRumStatsLogListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRumStatsLogListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScoresRequestParams struct {
	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 该参数已废弃
	IsDemo *int64 `json:"IsDemo,omitnil,omitempty" name:"IsDemo"`
}

type DescribeScoresRequest struct {
	*tchttp.BaseRequest
	
	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 该参数已废弃
	IsDemo *int64 `json:"IsDemo,omitnil,omitempty" name:"IsDemo"`
}

func (r *DescribeScoresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScoresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EndTime")
	delete(f, "StartTime")
	delete(f, "ID")
	delete(f, "IsDemo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScoresRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScoresResponseParams struct {
	// 数组
	ScoreSet []*ScoreInfo `json:"ScoreSet,omitnil,omitempty" name:"ScoreSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeScoresResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScoresResponseParams `json:"Response"`
}

func (r *DescribeScoresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScoresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTawAreasRequestParams struct {
	// 片区Id
	AreaIds []*int64 `json:"AreaIds,omitnil,omitempty" name:"AreaIds"`

	// 片区Key
	AreaKeys []*string `json:"AreaKeys,omitnil,omitempty" name:"AreaKeys"`

	// 分页Limit，默认根据AreaKeys和AreaStatuses参数查询所有。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 片区状态(1=有效，2=无效)
	AreaStatuses []*int64 `json:"AreaStatuses,omitnil,omitempty" name:"AreaStatuses"`

	// 分页Offset，默认根据AreaKeys和AreaStatuses参数查询所有。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeTawAreasRequest struct {
	*tchttp.BaseRequest
	
	// 片区Id
	AreaIds []*int64 `json:"AreaIds,omitnil,omitempty" name:"AreaIds"`

	// 片区Key
	AreaKeys []*string `json:"AreaKeys,omitnil,omitempty" name:"AreaKeys"`

	// 分页Limit，默认根据AreaKeys和AreaStatuses参数查询所有。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 片区状态(1=有效，2=无效)
	AreaStatuses []*int64 `json:"AreaStatuses,omitnil,omitempty" name:"AreaStatuses"`

	// 分页Offset，默认根据AreaKeys和AreaStatuses参数查询所有。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeTawAreasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTawAreasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AreaIds")
	delete(f, "AreaKeys")
	delete(f, "Limit")
	delete(f, "AreaStatuses")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTawAreasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTawAreasResponseParams struct {
	// 片区总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 片区列表
	AreaSet []*RumAreaInfo `json:"AreaSet,omitnil,omitempty" name:"AreaSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTawAreasResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTawAreasResponseParams `json:"Response"`
}

func (r *DescribeTawAreasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTawAreasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTawInstancesRequestParams struct {
	// 计费状态
	ChargeStatuses []*int64 `json:"ChargeStatuses,omitnil,omitempty" name:"ChargeStatuses"`

	// 计费类型
	ChargeTypes []*int64 `json:"ChargeTypes,omitnil,omitempty" name:"ChargeTypes"`

	// 分页Limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 片区Id
	AreaIds []*int64 `json:"AreaIds,omitnil,omitempty" name:"AreaIds"`

	// 实例状态(1=创建中，2=运行中，3=异常，4=重启中，5=停止中，6=已停止，7=销毁中，8=已销毁), 该参数已废弃，请在Filters内注明
	InstanceStatuses []*int64 `json:"InstanceStatuses,omitnil,omitempty" name:"InstanceStatuses"`

	// 实例Id, 该参数已废弃，请在Filters内注明
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 过滤参数；demo模式传{"Name": "IsDemo", "Values":["1"]}
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 该参数已废弃，demo模式请在Filters内注明
	IsDemo *int64 `json:"IsDemo,omitnil,omitempty" name:"IsDemo"`
}

type DescribeTawInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 计费状态
	ChargeStatuses []*int64 `json:"ChargeStatuses,omitnil,omitempty" name:"ChargeStatuses"`

	// 计费类型
	ChargeTypes []*int64 `json:"ChargeTypes,omitnil,omitempty" name:"ChargeTypes"`

	// 分页Limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 片区Id
	AreaIds []*int64 `json:"AreaIds,omitnil,omitempty" name:"AreaIds"`

	// 实例状态(1=创建中，2=运行中，3=异常，4=重启中，5=停止中，6=已停止，7=销毁中，8=已销毁), 该参数已废弃，请在Filters内注明
	InstanceStatuses []*int64 `json:"InstanceStatuses,omitnil,omitempty" name:"InstanceStatuses"`

	// 实例Id, 该参数已废弃，请在Filters内注明
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 过滤参数；demo模式传{"Name": "IsDemo", "Values":["1"]}
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 该参数已废弃，demo模式请在Filters内注明
	IsDemo *int64 `json:"IsDemo,omitnil,omitempty" name:"IsDemo"`
}

func (r *DescribeTawInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTawInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChargeStatuses")
	delete(f, "ChargeTypes")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "AreaIds")
	delete(f, "InstanceStatuses")
	delete(f, "InstanceIds")
	delete(f, "Filters")
	delete(f, "IsDemo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTawInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTawInstancesResponseParams struct {
	// 实例列表
	InstanceSet []*RumInstanceInfo `json:"InstanceSet,omitnil,omitempty" name:"InstanceSet"`

	// 实例总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTawInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTawInstancesResponseParams `json:"Response"`
}

func (r *DescribeTawInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTawInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUvListRequestParams struct {
	// ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 获取day：d，   min:m
	Dimension *string `json:"Dimension,omitnil,omitempty" name:"Dimension"`
}

type DescribeUvListRequest struct {
	*tchttp.BaseRequest
	
	// ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 获取day：d，   min:m
	Dimension *string `json:"Dimension,omitnil,omitempty" name:"Dimension"`
}

func (r *DescribeUvListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUvListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "EndTime")
	delete(f, "StartTime")
	delete(f, "Dimension")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUvListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUvListResponseParams struct {
	// uv列表
	ProjectUvSet []*RumUvInfo `json:"ProjectUvSet,omitnil,omitempty" name:"ProjectUvSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUvListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUvListResponseParams `json:"Response"`
}

func (r *DescribeUvListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUvListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWhitelistsRequestParams struct {
	// 实例instance-ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
}

type DescribeWhitelistsRequest struct {
	*tchttp.BaseRequest
	
	// 实例instance-ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
}

func (r *DescribeWhitelistsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWhitelistsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWhitelistsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWhitelistsResponseParams struct {
	// 白名单列表
	WhitelistSet []*Whitelist `json:"WhitelistSet,omitnil,omitempty" name:"WhitelistSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWhitelistsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWhitelistsResponseParams `json:"Response"`
}

func (r *DescribeWhitelistsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWhitelistsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 一个或者多个过滤值。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// 过滤键的名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

// Predefined struct for user
type ModifyInstanceRequestParams struct {
	// 要修改的实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 新的实例名称(长度最大不超过255)
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 新的实例描述(长度最大不超过1024)
	InstanceDesc *string `json:"InstanceDesc,omitnil,omitempty" name:"InstanceDesc"`
}

type ModifyInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 要修改的实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 新的实例名称(长度最大不超过255)
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 新的实例描述(长度最大不超过1024)
	InstanceDesc *string `json:"InstanceDesc,omitnil,omitempty" name:"InstanceDesc"`
}

func (r *ModifyInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InstanceName")
	delete(f, "InstanceDesc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceResponseParams `json:"Response"`
}

func (r *ModifyInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProjectLimitRequestParams struct {
	// 项目ID
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`

	// 取值为[log speed performance webvitals pv event custom miniProgramData]其中之一
	ProjectInterface *string `json:"ProjectInterface,omitnil,omitempty" name:"ProjectInterface"`

	// 上报比例   10代表10%
	ReportRate *int64 `json:"ReportRate,omitnil,omitempty" name:"ReportRate"`

	// 上报类型 1：比例  2：上报量
	ReportType *int64 `json:"ReportType,omitnil,omitempty" name:"ReportType"`

	// 主键ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

type ModifyProjectLimitRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`

	// 取值为[log speed performance webvitals pv event custom miniProgramData]其中之一
	ProjectInterface *string `json:"ProjectInterface,omitnil,omitempty" name:"ProjectInterface"`

	// 上报比例   10代表10%
	ReportRate *int64 `json:"ReportRate,omitnil,omitempty" name:"ReportRate"`

	// 上报类型 1：比例  2：上报量
	ReportType *int64 `json:"ReportType,omitnil,omitempty" name:"ReportType"`

	// 主键ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

func (r *ModifyProjectLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProjectLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectID")
	delete(f, "ProjectInterface")
	delete(f, "ReportRate")
	delete(f, "ReportType")
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyProjectLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProjectLimitResponseParams struct {
	// 返回信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyProjectLimitResponse struct {
	*tchttp.BaseResponse
	Response *ModifyProjectLimitResponseParams `json:"Response"`
}

func (r *ModifyProjectLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProjectLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProjectRequestParams struct {
	// 项目 id
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 应用名称(可选，不为空且最长为 200字符)
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 项目网页地址(可选，最长为 256)
	URL *string `json:"URL,omitnil,omitempty" name:"URL"`

	// 项目仓库地址(可选，最长为 256)
	Repo *string `json:"Repo,omitnil,omitempty" name:"Repo"`

	// 项目需要转移到的实例 id(可选)
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 项目采样率(可选)
	Rate *string `json:"Rate,omitnil,omitempty" name:"Rate"`

	// 是否开启聚类(可选)
	EnableURLGroup *uint64 `json:"EnableURLGroup,omitnil,omitempty" name:"EnableURLGroup"`

	// 项目类型(可接受值为 "web", "mp", "android", "ios", "node", "hippy", "weex", "viola", "rn")
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 应用描述(可选，最长为 1000字符)
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`
}

type ModifyProjectRequest struct {
	*tchttp.BaseRequest
	
	// 项目 id
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 应用名称(可选，不为空且最长为 200字符)
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 项目网页地址(可选，最长为 256)
	URL *string `json:"URL,omitnil,omitempty" name:"URL"`

	// 项目仓库地址(可选，最长为 256)
	Repo *string `json:"Repo,omitnil,omitempty" name:"Repo"`

	// 项目需要转移到的实例 id(可选)
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 项目采样率(可选)
	Rate *string `json:"Rate,omitnil,omitempty" name:"Rate"`

	// 是否开启聚类(可选)
	EnableURLGroup *uint64 `json:"EnableURLGroup,omitnil,omitempty" name:"EnableURLGroup"`

	// 项目类型(可接受值为 "web", "mp", "android", "ios", "node", "hippy", "weex", "viola", "rn")
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 应用描述(可选，最长为 1000字符)
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`
}

func (r *ModifyProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ID")
	delete(f, "Name")
	delete(f, "URL")
	delete(f, "Repo")
	delete(f, "InstanceID")
	delete(f, "Rate")
	delete(f, "EnableURLGroup")
	delete(f, "Type")
	delete(f, "Desc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProjectResponseParams struct {
	// 操作信息
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 项目id
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyProjectResponse struct {
	*tchttp.BaseResponse
	Response *ModifyProjectResponseParams `json:"Response"`
}

func (r *ModifyProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProjectLimit struct {
	// 接口
	ProjectInterface *string `json:"ProjectInterface,omitnil,omitempty" name:"ProjectInterface"`

	// 上报率
	ReportRate *int64 `json:"ReportRate,omitnil,omitempty" name:"ReportRate"`

	// 上报类型 1：上报率  2：上报量限制
	ReportType *int64 `json:"ReportType,omitnil,omitempty" name:"ReportType"`

	// 主键ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 项目ID
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`
}

type ReleaseFile struct {
	// 文件版本
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 文件唯一 key
	FileKey *string `json:"FileKey,omitnil,omitempty" name:"FileKey"`

	// 文件名
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文件哈希值
	FileHash *string `json:"FileHash,omitnil,omitempty" name:"FileHash"`

	// 文件 id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

// Predefined struct for user
type ResumeInstanceRequestParams struct {
	// 需要恢复的实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type ResumeInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 需要恢复的实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *ResumeInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResumeInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResumeInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ResumeInstanceResponseParams `json:"Response"`
}

func (r *ResumeInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeProjectRequestParams struct {
	// 项目 id
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type ResumeProjectRequest struct {
	*tchttp.BaseRequest
	
	// 项目 id
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *ResumeProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResumeProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeProjectResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResumeProjectResponse struct {
	*tchttp.BaseResponse
	Response *ResumeProjectResponseParams `json:"Response"`
}

func (r *ResumeProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RumAreaInfo struct {
	// 片区Id
	AreaId *int64 `json:"AreaId,omitnil,omitempty" name:"AreaId"`

	// 片区状态(1=有效，2=无效)
	AreaStatus *int64 `json:"AreaStatus,omitnil,omitempty" name:"AreaStatus"`

	// 片区名称
	AreaName *string `json:"AreaName,omitnil,omitempty" name:"AreaName"`

	// 片区Key
	AreaKey *string `json:"AreaKey,omitnil,omitempty" name:"AreaKey"`

	// 地域码表 id
	AreaRegionID *string `json:"AreaRegionID,omitnil,omitempty" name:"AreaRegionID"`

	// 地域码表 code 如 ap-xxx（xxx 为地域词）
	AreaRegionCode *string `json:"AreaRegionCode,omitnil,omitempty" name:"AreaRegionCode"`

	// 地域缩写
	AreaAbbr *string `json:"AreaAbbr,omitnil,omitempty" name:"AreaAbbr"`
}

type RumInstanceInfo struct {
	// 实例状态(1=创建中，2=运行中，3=异常，4=重启中，5=停止中，6=已停止，7=已删除)
	InstanceStatus *int64 `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// 片区Id
	AreaId *int64 `json:"AreaId,omitnil,omitempty" name:"AreaId"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 集群Id
	ClusterId *int64 `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 实例描述
	InstanceDesc *string `json:"InstanceDesc,omitnil,omitempty" name:"InstanceDesc"`

	// 计费状态(1=使用中，2=已过期，3=已销毁，4=分配中，5=分配失败)
	ChargeStatus *int64 `json:"ChargeStatus,omitnil,omitempty" name:"ChargeStatus"`

	// 计费类型(1=免费版，2=预付费，3=后付费)
	ChargeType *int64 `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 更新时间
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// 数据保留时间(天)
	DataRetentionDays *int64 `json:"DataRetentionDays,omitnil,omitempty" name:"DataRetentionDays"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 实例类型 1:原web相关类型 2:app端类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`
}

type RumProject struct {
	// 项目名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 创建者 id
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 实例 id
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 项目类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 项目仓库地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Repo *string `json:"Repo,omitnil,omitempty" name:"Repo"`

	// 项目网址地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	URL *string `json:"URL,omitnil,omitempty" name:"URL"`

	// 项目采样频率
	Rate *string `json:"Rate,omitnil,omitempty" name:"Rate"`

	// 项目唯一key（长度 12 位）
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 是否开启url聚类
	EnableURLGroup *int64 `json:"EnableURLGroup,omitnil,omitempty" name:"EnableURLGroup"`

	// 实例名
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 项目 ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 实例 key
	InstanceKey *string `json:"InstanceKey,omitnil,omitempty" name:"InstanceKey"`

	// 项目描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 是否星标  1:是 0:否
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsStar *int64 `json:"IsStar,omitnil,omitempty" name:"IsStar"`

	// 项目状态(1 创建中，2 运行中，3 异常，4 重启中，5 停止中，6 已停止， 7 销毁中，8 已销毁)
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectStatus *int64 `json:"ProjectStatus,omitnil,omitempty" name:"ProjectStatus"`

	// 日志接入点，用户忽略。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessPoint *string `json:"AccessPoint,omitnil,omitempty" name:"AccessPoint"`
}

type RumPvInfo struct {
	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// pv访问量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pv *string `json:"Pv,omitnil,omitempty" name:"Pv"`

	// 时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type RumUvInfo struct {
	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// uv访问量
	Uv *string `json:"Uv,omitnil,omitempty" name:"Uv"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type ScoreInfo struct {
	// duration
	StaticDuration *string `json:"StaticDuration,omitnil,omitempty" name:"StaticDuration"`

	// pv
	PagePv *string `json:"PagePv,omitnil,omitempty" name:"PagePv"`

	// 失败
	ApiFail *string `json:"ApiFail,omitnil,omitempty" name:"ApiFail"`

	// 请求
	ApiNum *string `json:"ApiNum,omitnil,omitempty" name:"ApiNum"`

	// fail
	StaticFail *string `json:"StaticFail,omitnil,omitempty" name:"StaticFail"`

	// 项目id
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`

	// uv
	PageUv *string `json:"PageUv,omitnil,omitempty" name:"PageUv"`

	// 请求次数
	ApiDuration *string `json:"ApiDuration,omitnil,omitempty" name:"ApiDuration"`

	// 分数
	Score *string `json:"Score,omitnil,omitempty" name:"Score"`

	// error
	PageError *string `json:"PageError,omitnil,omitempty" name:"PageError"`

	// num
	StaticNum *string `json:"StaticNum,omitnil,omitempty" name:"StaticNum"`

	// num
	RecordNum *int64 `json:"RecordNum,omitnil,omitempty" name:"RecordNum"`

	// Duration
	PageDuration *string `json:"PageDuration,omitnil,omitempty" name:"PageDuration"`

	// 时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

// Predefined struct for user
type StopInstanceRequestParams struct {
	// 需要停止的实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type StopInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 需要停止的实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *StopInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopInstanceResponse struct {
	*tchttp.BaseResponse
	Response *StopInstanceResponseParams `json:"Response"`
}

func (r *StopInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopProjectRequestParams struct {
	// 项目 id
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type StopProjectRequest struct {
	*tchttp.BaseRequest
	
	// 项目 id
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *StopProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopProjectResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopProjectResponse struct {
	*tchttp.BaseResponse
	Response *StopProjectResponseParams `json:"Response"`
}

func (r *StopProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {
	// 标签key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 标签value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type Whitelist struct {
	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 截止时间
	Ttl *string `json:"Ttl,omitnil,omitempty" name:"Ttl"`

	// 白名单自增ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 业务唯一标识
	WhitelistUin *string `json:"WhitelistUin,omitnil,omitempty" name:"WhitelistUin"`

	// 创建者ID
	CreateUser *string `json:"CreateUser,omitnil,omitempty" name:"CreateUser"`

	// aid标识
	Aid *string `json:"Aid,omitnil,omitempty" name:"Aid"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}