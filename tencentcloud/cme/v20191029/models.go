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

package v20191029

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AddMemberInfo struct {

	// 团队成员 ID。
	MemberId *string `json:"MemberId,omitempty" name:"MemberId"`

	// 团队成员备注。
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type AddTeamMemberRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 团队 ID。
	TeamId *string `json:"TeamId,omitempty" name:"TeamId"`

	// 要添加的成员列表，一次最多添加30个成员。
	TeamMembers []*AddMemberInfo `json:"TeamMembers,omitempty" name:"TeamMembers" list`

	// 操作者。填写用户的 Id，用于标识调用者及校验操作权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *AddTeamMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddTeamMemberRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddTeamMemberResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddTeamMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddTeamMemberResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AudioMaterial struct {

	// 素材元信息。
	MetaData *MediaMetaData `json:"MetaData,omitempty" name:"MetaData"`

	// 素材媒体文件的 URL 地址。
	MaterialUrl *string `json:"MaterialUrl,omitempty" name:"MaterialUrl"`

	// 素材媒体文件的封面图片地址。
	CoverUrl *string `json:"CoverUrl,omitempty" name:"CoverUrl"`

	// 素材状态。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaterialStatus *MaterialStatus `json:"MaterialStatus,omitempty" name:"MaterialStatus"`
}

type AudioStreamInfo struct {

	// 码率，单位：bps。
	Bitrate *uint64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 采样率，单位：hz。
	SamplingRate *uint64 `json:"SamplingRate,omitempty" name:"SamplingRate"`

	// 编码格式。
	Codec *string `json:"Codec,omitempty" name:"Codec"`
}

type AuthorizationInfo struct {

	// 被授权者实体。
	Authorizee *Entity `json:"Authorizee,omitempty" name:"Authorizee"`

	// 详细授权值。 取值有：
	// <li>R：可读，可以浏览素材，但不能使用该素材（将其添加到 Project），或复制到自己的媒资库中</li>
	// <li>X：可用，可以使用该素材（将其添加到 Project），但不能将其复制到自己的媒资库中，意味着被授权者无法将该资源进一步扩散给其他个人或团队。</li>
	// <li>C：可复制，既可以使用该素材（将其添加到 Project），也可以将其复制到自己的媒资库中。</li>
	// <li>W：可修改、删除媒资。</li>
	PermissionSet []*string `json:"PermissionSet,omitempty" name:"PermissionSet" list`
}

type Authorizer struct {

	// 授权者类型，取值有：
	// <li>PERSON：个人。</li>
	// <li>TEAM：团队。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Id，当 Type=PERSON，取值为用户 Id。当Type=TEAM，取值为团队 ID。
	Id *string `json:"Id,omitempty" name:"Id"`
}

type CMEExportInfo struct {

	// 导出的归属者。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 导出的素材名称，不得超过30个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 导出的素材信息，不得超过50个字符。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 导出的素材分类路径，长度不能超过15字符。
	ClassPath *string `json:"ClassPath,omitempty" name:"ClassPath"`

	// 导出的素材标签，单个标签不得超过10个字符。
	TagSet []*string `json:"TagSet,omitempty" name:"TagSet" list`
}

type ClassInfo struct {

	// 归属者。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 分类路径。
	ClassPath *string `json:"ClassPath,omitempty" name:"ClassPath"`
}

type CreateClassRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 归属者。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 分类路径。
	ClassPath *string `json:"ClassPath,omitempty" name:"ClassPath"`

	// 操作者。填写用户的 Id，用于标识调用者及校验操作权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *CreateClassRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateClassRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateClassResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClassResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateClassResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLinkRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 链接类型，取值有:
	// <li>CLASS: 分类链接；</li>
	// <li> MATERIAL：素材链接。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 链接名称，不能超过30个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 链接归属实体。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 目标资源Id。取值：
	// <li>当 Type 为 MATERIAL 时填素材 ID；</li>
	// <li>当 Type 为 CLASS 时填写分类路径。</li>
	DestinationId *string `json:"DestinationId,omitempty" name:"DestinationId"`

	// 目标资源归属者。
	DestinationOwner *Entity `json:"DestinationOwner,omitempty" name:"DestinationOwner"`

	// 链接的分类路径，如填"/a/b"则代表链接属于该分类路径，不填则默认为根路径。
	ClassPath *string `json:"ClassPath,omitempty" name:"ClassPath"`

	// 链接标签，单个标签长度不能超过10，数组长度不能超过10。
	Tags []*string `json:"Tags,omitempty" name:"Tags" list`

	// 操作者。填写用户的 Id，用于标识调用者及校验操作权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *CreateLinkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLinkRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLinkResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 新建链接的素材 Id。
		MaterialId *string `json:"MaterialId,omitempty" name:"MaterialId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLinkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLinkResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateProjectRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 项目类别，取值有：
	// <li>VIDEO_EDIT：视频编辑。</li>
	Category *string `json:"Category,omitempty" name:"Category"`

	// 项目名称，不可超过30个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 画布宽高比，取值有：
	// <li>16:9；</li>
	// <li>9:16。</li>
	AspectRatio *string `json:"AspectRatio,omitempty" name:"AspectRatio"`

	// 归属者。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`
}

func (r *CreateProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateProjectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 项目 Id。
		ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateProjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTeamRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 团队名称，限30个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 团队所有者，指定用户 ID。
	OwnerId *string `json:"OwnerId,omitempty" name:"OwnerId"`

	// 团队所有者的备注，限30个字符。
	OwnerRemark *string `json:"OwnerRemark,omitempty" name:"OwnerRemark"`

	// 自定义团队 ID。创建后不可修改，限20个英文字符及"-"。同时不能以 cmetid_开头。不填会生成默认团队 ID。
	TeamId *string `json:"TeamId,omitempty" name:"TeamId"`
}

func (r *CreateTeamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTeamRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTeamResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 创建的团队 ID。
		TeamId *string `json:"TeamId,omitempty" name:"TeamId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTeamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTeamResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteClassRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 归属者。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 分类路径。
	ClassPath *string `json:"ClassPath,omitempty" name:"ClassPath"`

	// 操作者。填写用户的 Id，用于标识调用者及校验操作权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *DeleteClassRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClassRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteClassResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClassResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClassResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLoginStatusRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 用户 Id 列表，N 从 0 开始取值，最大 19。
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds" list`
}

func (r *DeleteLoginStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLoginStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLoginStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLoginStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLoginStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMaterialRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 素材 Id。
	MaterialId *string `json:"MaterialId,omitempty" name:"MaterialId"`

	// 操作者。填写用户的 Id，用于标识调用者及校验操作权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *DeleteMaterialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMaterialRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMaterialResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMaterialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMaterialResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteProjectRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 项目 Id。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DeleteProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteProjectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteProjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTeamMembersRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 团队 ID。
	TeamId *string `json:"TeamId,omitempty" name:"TeamId"`

	// 要删除的成员列表。
	MemberIds []*string `json:"MemberIds,omitempty" name:"MemberIds" list`

	// 操作者。填写用户的 Id，用于标识调用者及校验操作权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *DeleteTeamMembersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTeamMembersRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTeamMembersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTeamMembersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTeamMembersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTeamRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 要删除的团队  ID。
	TeamId *string `json:"TeamId,omitempty" name:"TeamId"`

	// 操作者。填写用户的 Id，用于标识调用者及校验操作权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *DeleteTeamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTeamRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTeamResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTeamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTeamResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClassRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 归属者。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 操作者。填写用户的 Id，用于标识调用者及校验操作权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *DescribeClassRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClassRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClassResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分类信息列表。
		ClassInfoSet []*ClassInfo `json:"ClassInfoSet,omitempty" name:"ClassInfoSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClassResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClassResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeJoinTeamsRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 团队成员　ID。
	MemberId *string `json:"MemberId,omitempty" name:"MemberId"`

	// 分页偏移量，默认值：0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回记录条数，默认值：30，最大值：30。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeJoinTeamsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeJoinTeamsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeJoinTeamsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 团队列表
		TeamSet []*JoinTeamInfo `json:"TeamSet,omitempty" name:"TeamSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeJoinTeamsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeJoinTeamsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLoginStatusRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 用户 Id 列表，N 从 0 开始取值，最大 19。
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds" list`
}

func (r *DescribeLoginStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLoginStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLoginStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 用户登录状态列表。
		LoginStatusInfoSet []*LoginStatusInfo `json:"LoginStatusInfoSet,omitempty" name:"LoginStatusInfoSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLoginStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLoginStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMaterialsRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 素材 ID 列表，N 从 0 开始取值，最大 19。
	MaterialIds []*string `json:"MaterialIds,omitempty" name:"MaterialIds" list`

	// 列表排序，支持下列排序字段：
	// <li>CreateTime：创建时间；</li>
	// <li>UpdateTime：更新时间。</li>
	Sort *SortBy `json:"Sort,omitempty" name:"Sort"`

	// 操作者。填写用户的 Id，用于标识调用者及校验操作权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *DescribeMaterialsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMaterialsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMaterialsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 素材列表信息。
		MaterialInfoSet []*MaterialInfo `json:"MaterialInfoSet,omitempty" name:"MaterialInfoSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMaterialsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMaterialsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectsRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 项目 Id 列表，N 从 0 开始取值，最大 19。
	ProjectIds []*string `json:"ProjectIds,omitempty" name:"ProjectIds" list`

	// 画布宽高比集合。
	AspectRatioSet []*string `json:"AspectRatioSet,omitempty" name:"AspectRatioSet" list`

	// 项目类别集合。
	CategorySet []*string `json:"CategorySet,omitempty" name:"CategorySet" list`

	// 列表排序，支持下列排序字段：
	// <li>CreateTime：创建时间；</li>
	// <li>UpdateTime：更新时间。</li>
	Sort *SortBy `json:"Sort,omitempty" name:"Sort"`

	// 项目归属者。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 分页返回的起始偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页返回的记录条数，默认值：10。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeProjectsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProjectsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 项目信息列表。
		ProjectInfoSet []*ProjectInfo `json:"ProjectInfoSet,omitempty" name:"ProjectInfoSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProjectsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProjectsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceAuthorizationRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 归属者。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 资源。
	Resource *Resource `json:"Resource,omitempty" name:"Resource"`

	// 操作者。填写用户的 Id，用于标识调用者及校验操作权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *DescribeResourceAuthorizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeResourceAuthorizationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceAuthorizationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的资源授权记录总数。
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 授权信息列表。
		AuthorizationInfoSet []*AuthorizationInfo `json:"AuthorizationInfoSet,omitempty" name:"AuthorizationInfoSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourceAuthorizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeResourceAuthorizationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSharedSpaceRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 被授权目标实体。
	Authorizee *Entity `json:"Authorizee,omitempty" name:"Authorizee"`

	// 操作者。填写用户的 Id，用于标识调用者及校验操作权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *DescribeSharedSpaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSharedSpaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSharedSpaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询到的共享空间总数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 各个共享空间对应的授权者信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
		AuthorizerSet []*Authorizer `json:"AuthorizerSet,omitempty" name:"AuthorizerSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSharedSpaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSharedSpaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskDetailRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 任务 Id。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeTaskDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务状态，取值有：
	// <li>PROCESSING：处理中：</li>
	// <li>SUCCESS：成功；</li>
	// <li>FAIL：失败。</li>
		Status *string `json:"Status,omitempty" name:"Status"`

		// 任务进度，取值为：0~100。
		Progress *uint64 `json:"Progress,omitempty" name:"Progress"`

		// 错误码。
	// <li>0：成功；</li>
	// <li>其他值：失败。</li>
		ErrCode *uint64 `json:"ErrCode,omitempty" name:"ErrCode"`

		// 错误信息。
		ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

		// 任务类型，取值有：
	// <li>VIDEO_EDIT_PROJECT_EXPORT：视频编辑项目导出。</li>
		TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

		// 导出项目输出信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
		VideoEditProjectOutput *VideoEditProjectOutput `json:"VideoEditProjectOutput,omitempty" name:"VideoEditProjectOutput"`

		// 创建时间，格式按照 ISO 8601 标准表示。
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTasksRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 项目 Id。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务类型集合，取值有：
	// <li>VIDEO_EDIT_PROJECT_EXPORT：视频编辑项目导出。</li>
	TaskTypeSet []*string `json:"TaskTypeSet,omitempty" name:"TaskTypeSet" list`

	// 任务状态集合，取值有：
	// <li>PROCESSING：处理中；</li>
	// <li>SUCCESS：成功；</li>
	// <li>FAIL：失败。</li>
	StatusSet []*string `json:"StatusSet,omitempty" name:"StatusSet" list`

	// 分页返回的起始偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页返回的记录条数，默认值：10。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTasksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合搜索条件的记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 任务基础信息列表。
		TaskBaseInfoSet []*TaskBaseInfo `json:"TaskBaseInfoSet,omitempty" name:"TaskBaseInfoSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTasksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTeamMembersRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 团队 ID。
	TeamId *string `json:"TeamId,omitempty" name:"TeamId"`

	// 成员 ID 列表，限指定30个指定成员。
	MemberIds []*string `json:"MemberIds,omitempty" name:"MemberIds" list`

	// 分页偏移量，默认值：0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回记录条数，默认值：30，最大值：30。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 操作者。填写用户的 Id，用于标识调用者及校验操作权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *DescribeTeamMembersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTeamMembersRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTeamMembersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 团队成员列表。
		MemberSet []*TeamMemberInfo `json:"MemberSet,omitempty" name:"MemberSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTeamMembersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTeamMembersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTeamsRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 团队 ID 列表，限30个。
	TeamIds []*string `json:"TeamIds,omitempty" name:"TeamIds" list`
}

func (r *DescribeTeamsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTeamsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTeamsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 团队列表。
		TeamSet []*TeamInfo `json:"TeamSet,omitempty" name:"TeamSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTeamsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTeamsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Entity struct {

	// 类型，取值有：
	// <li>PERSON：个人。</li>
	// <li>TEAM：团队。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Id，当 Type=PERSON，取值为用户 Id，当 Type=TEAM，取值为团队 Id。
	Id *string `json:"Id,omitempty" name:"Id"`
}

type ExportVideoByEditorTrackDataRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 导出模板 Id，目前不支持自定义创建，只支持下面的预置模板 Id。
	// <li>10：分辨率为 480P，输出视频格式为 MP4；</li>
	// <li>11：分辨率为 720P，输出视频格式为 MP4；</li>
	// <li>12：分辨率为 1080P，输出视频格式为 MP4。</li>
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 导出目标。
	// <li>CME：云剪，即导出为云剪素材；</li>
	// <li>VOD：云点播，即导出为云点播媒资。</li>
	ExportDestination *string `json:"ExportDestination,omitempty" name:"ExportDestination"`

	// 在线编辑轨道数据。
	TrackData *string `json:"TrackData,omitempty" name:"TrackData"`

	// 导出的云剪素材信息。指定 ExportDestination = CME 时有效。
	CMEExportInfo *CMEExportInfo `json:"CMEExportInfo,omitempty" name:"CMEExportInfo"`

	// 导出的云点播媒资信息。指定 ExportDestination = VOD 时有效。
	VODExportInfo *VODExportInfo `json:"VODExportInfo,omitempty" name:"VODExportInfo"`

	// 操作者。填写用户的 Id，用于标识调用者及校验操作权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *ExportVideoByEditorTrackDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExportVideoByEditorTrackDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ExportVideoByEditorTrackDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务 Id。
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportVideoByEditorTrackDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExportVideoByEditorTrackDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ExportVideoEditProjectRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 项目 Id。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 导出模板 Id，目前不支持自定义创建，只支持下面的预置模板 Id。
	// <li>10：分辨率为 480P，输出视频格式为 MP4；</li>
	// <li>11：分辨率为 720P，输出视频格式为 MP4；</li>
	// <li>12：分辨率为 1080P，输出视频格式为 MP4。</li>
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 导出目标。
	// <li>CME：云剪，即导出为云剪素材；</li>
	// <li>VOD：云点播，即导出为云点播媒资。</li>
	ExportDestination *string `json:"ExportDestination,omitempty" name:"ExportDestination"`

	// 导出的云剪素材信息。指定 ExportDestination = CME 时有效。
	CMEExportInfo *CMEExportInfo `json:"CMEExportInfo,omitempty" name:"CMEExportInfo"`

	// 导出的云点播媒资信息。指定 ExportDestination = VOD 时有效。
	VODExportInfo *VODExportInfo `json:"VODExportInfo,omitempty" name:"VODExportInfo"`
}

func (r *ExportVideoEditProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExportVideoEditProjectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ExportVideoEditProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务 Id。
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportVideoEditProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExportVideoEditProjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type FlattenListMediaRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 素材分类路径，例如填写"/a/b"，则代表平铺该分类路径下及其子分类路径下的素材信息。
	ClassPath *string `json:"ClassPath,omitempty" name:"ClassPath"`

	// 素材路径的归属者。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 分页偏移量，默认值：0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：50。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 操作者。填写用户的 Id，用于标识调用者及校验操作权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *FlattenListMediaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *FlattenListMediaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type FlattenListMediaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的记录总数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 该分类路径下及其子分类下的所有素材。
		MaterialInfoSet []*MaterialInfo `json:"MaterialInfoSet,omitempty" name:"MaterialInfoSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *FlattenListMediaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *FlattenListMediaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GrantResourceAuthorizationRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 资源所属实体。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 被授权资源。
	Resources []*Resource `json:"Resources,omitempty" name:"Resources" list`

	// 被授权目标实体。
	Authorizees []*Entity `json:"Authorizees,omitempty" name:"Authorizees" list`

	// 详细授权值。 取值有：
	// <li>R：可读，可以浏览素材，但不能使用该素材（将其添加到 Project），或复制到自己的媒资库中</li>
	// <li>X：可用，可以使用该素材（将其添加到 Project），但不能将其复制到自己的媒资库中，意味着被授权者无法将该资源进一步扩散给其他个人或团队。</li>
	// <li>C：可复制，既可以使用该素材（将其添加到 Project），也可以将其复制到自己的媒资库中。</li>
	// <li>W：可修改、删除媒资。</li>
	Permissions []*string `json:"Permissions,omitempty" name:"Permissions" list`

	// 操作者。填写用户的 Id，用于标识调用者及校验操作权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *GrantResourceAuthorizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GrantResourceAuthorizationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GrantResourceAuthorizationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GrantResourceAuthorizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GrantResourceAuthorizationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImageMaterial struct {

	// 图片高度，单位：px。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 图片宽度，单位：px。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 素材媒体文件的 URL 地址。
	MaterialUrl *string `json:"MaterialUrl,omitempty" name:"MaterialUrl"`

	// 图片大小，单位：字节。
	Size *int64 `json:"Size,omitempty" name:"Size"`
}

type ImportMaterialRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 云点播媒资 FileId。
	VodFileId *string `json:"VodFileId,omitempty" name:"VodFileId"`

	// 素材归属者。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 素材名称，不能超过30个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 素材分类路径，形如："/a/b"，层级数不能超过10，每个层级长度不能超过15字符。若不填则默认为根路径。
	ClassPath *string `json:"ClassPath,omitempty" name:"ClassPath"`

	// 素材标签，单个标签长度不能超过10，数组长度不能超过10。
	Tags []*string `json:"Tags,omitempty" name:"Tags" list`

	// 素材预处理任务模板 ID。取值：
	// <li>10：进行编辑预处理。</li>
	PreProcessDefinition *int64 `json:"PreProcessDefinition,omitempty" name:"PreProcessDefinition"`

	// 操作者。填写用户的 Id，用于标识调用者及校验操作权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *ImportMaterialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImportMaterialRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImportMaterialResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 素材 Id。
		MaterialId *string `json:"MaterialId,omitempty" name:"MaterialId"`

		// 素材预处理任务 ID，如果未指定发起预处理任务则为空。
		PreProcessTaskId *string `json:"PreProcessTaskId,omitempty" name:"PreProcessTaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportMaterialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImportMaterialResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImportMediaToProjectRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 项目 Id。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 云点播媒资 FileId。
	VodFileId *string `json:"VodFileId,omitempty" name:"VodFileId"`

	// 素材名称，不能超过30个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 素材预处理任务模板 ID，取值：
	// <li>10：进行编辑预处理。</li>
	// 注意：如果填0则不进行处理。
	PreProcessDefinition *int64 `json:"PreProcessDefinition,omitempty" name:"PreProcessDefinition"`
}

func (r *ImportMediaToProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImportMediaToProjectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImportMediaToProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 素材 Id。
		MaterialId *string `json:"MaterialId,omitempty" name:"MaterialId"`

		// 素材预处理任务 ID，如果未指定发起预处理任务则为空。
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportMediaToProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImportMediaToProjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type IntegerRange struct {

	// 按整形代表值的下限检索。
	LowerBound *int64 `json:"LowerBound,omitempty" name:"LowerBound"`

	// 按整形代表值的上限检索。
	UpperBound *int64 `json:"UpperBound,omitempty" name:"UpperBound"`
}

type JoinTeamInfo struct {

	// 团队 ID。
	TeamId *string `json:"TeamId,omitempty" name:"TeamId"`

	// 团队名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 团队成员个数
	MemberCount *uint64 `json:"MemberCount,omitempty" name:"MemberCount"`

	// 成员在团队中的角色，取值有：
	// <li>Owner：团队所有者，添加团队成员及修改团队成员解决时不能填此角色；</li>
	// <li>Admin：团队管理员；</li>
	// <li>Member：普通成员。</li>
	Role *string `json:"Role,omitempty" name:"Role"`
}

type LinkMaterial struct {

	// 链接类型取值:
	// <li>CLASS: 分类链接;</li>
	// <li> MATERIAL：素材链接。</li>
	LinkType *string `json:"LinkType,omitempty" name:"LinkType"`

	// 链接状态取值：
	// <li> Normal：正常 ；</li>
	// <li>NotFound：链接目标不存在；</li> <li>Forbidden：无权限。</li>
	LinkStatus *string `json:"LinkStatus,omitempty" name:"LinkStatus"`

	// 素材链接详细信息，当LinkType="MATERIAL"时有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LinkMaterialInfo *LinkMaterialInfo `json:"LinkMaterialInfo,omitempty" name:"LinkMaterialInfo"`

	// 分类链接目标信息，当LinkType=“CLASS”时有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LinkClassInfo *ClassInfo `json:"LinkClassInfo,omitempty" name:"LinkClassInfo"`
}

type LinkMaterialInfo struct {

	// 素材基本信息。
	BasicInfo *MaterialBasicInfo `json:"BasicInfo,omitempty" name:"BasicInfo"`

	// 视频素材信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoMaterial *VideoMaterial `json:"VideoMaterial,omitempty" name:"VideoMaterial"`

	// 音频素材信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioMaterial *AudioMaterial `json:"AudioMaterial,omitempty" name:"AudioMaterial"`

	// 图片素材信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageMaterial *ImageMaterial `json:"ImageMaterial,omitempty" name:"ImageMaterial"`
}

type ListMediaRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 素材分类路径，例如填写"/a/b"，则代表浏览该分类路径下的素材和子分类信息。
	ClassPath *string `json:"ClassPath,omitempty" name:"ClassPath"`

	// 素材和分类的归属者。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 分页偏移量，默认值：0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：50。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 操作者。填写用户的 Id，用于标识调用者及校验操作权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *ListMediaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListMediaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListMediaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的素材记录总数。
		MaterialTotalCount *int64 `json:"MaterialTotalCount,omitempty" name:"MaterialTotalCount"`

		// 浏览分类路径下的素材列表信息。
		MaterialInfoSet []*MaterialInfo `json:"MaterialInfoSet,omitempty" name:"MaterialInfoSet" list`

		// 浏览分类路径下的一级子类。
		ClassInfoSet []*ClassInfo `json:"ClassInfoSet,omitempty" name:"ClassInfoSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListMediaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListMediaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LoginStatusInfo struct {

	// 用户 Id。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户登录状态。
	// <li>Online：在线；</li>
	// <li>Offline：离线。</li>
	Status *string `json:"Status,omitempty" name:"Status"`
}

type MaterialBasicInfo struct {

	// 素材 Id。
	MaterialId *string `json:"MaterialId,omitempty" name:"MaterialId"`

	// 素材类型，取值为：音频（AUDIO）、视频（VIDEO）、图片（IMAGE）、链接（LINK）、字幕 （SUBTITLE）、转场（TRANSITION）、滤镜（FILTER）、文本文字（TEXT）、图文动效（TEXT_IMAGE）。
	MaterialType *string `json:"MaterialType,omitempty" name:"MaterialType"`

	// 素材归属实体。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 素材名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 素材文件的创建时间，使用 ISO 日期格式。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 素材文件的最近更新时间（如修改视频属性、发起视频处理等会触发更新媒体文件信息的操作），使用 ISO 日期格式。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 素材的分类目录路径。
	ClassPath *string `json:"ClassPath,omitempty" name:"ClassPath"`

	// 素材标签信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagSet []*string `json:"TagSet,omitempty" name:"TagSet" list`

	// 素材媒体文件的预览图。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PreviewUrl *string `json:"PreviewUrl,omitempty" name:"PreviewUrl"`
}

type MaterialInfo struct {

	// 素材基本信息。
	BasicInfo *MaterialBasicInfo `json:"BasicInfo,omitempty" name:"BasicInfo"`

	// 视频素材信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoMaterial *VideoMaterial `json:"VideoMaterial,omitempty" name:"VideoMaterial"`

	// 音频素材信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioMaterial *AudioMaterial `json:"AudioMaterial,omitempty" name:"AudioMaterial"`

	// 图片素材信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageMaterial *ImageMaterial `json:"ImageMaterial,omitempty" name:"ImageMaterial"`

	// 链接素材信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LinkMaterial *LinkMaterial `json:"LinkMaterial,omitempty" name:"LinkMaterial"`
}

type MaterialStatus struct {

	// 素材编辑可用状态，取值有：
	// <li>NORMAL：正常，可直接用于编辑；</li>
	// <li>ABNORMAL : 异常，不可用于编辑；</li>
	// <li>PROCESSING：处理中，暂不可用于编辑。</li>
	EditorUsableStatus *string `json:"EditorUsableStatus,omitempty" name:"EditorUsableStatus"`
}

type MediaImageSpriteInfo struct {

	// 雪碧图小图的高度。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 雪碧图小图的宽度。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 雪碧图小图的总数量。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 截取雪碧图输出的地址。
	ImageUrlSet []*string `json:"ImageUrlSet,omitempty" name:"ImageUrlSet" list`

	// 雪碧图子图位置与时间关系的 WebVtt 文件地址。WebVtt 文件表明了各个雪碧图小图对应的时间点，以及在雪碧大图里的坐标位置，一般被播放器用于实现预览。
	WebVttUrl *string `json:"WebVttUrl,omitempty" name:"WebVttUrl"`
}

type MediaMetaData struct {

	// 大小。
	Size *uint64 `json:"Size,omitempty" name:"Size"`

	// 容器类型。
	Container *string `json:"Container,omitempty" name:"Container"`

	// 视频流码率平均值与音频流码率平均值之和，单位：bps。
	Bitrate *uint64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 视频流高度的最大值，单位：px。
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// 视频流宽度的最大值，单位：px。
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// 时长，单位：秒。
	Duration *float64 `json:"Duration,omitempty" name:"Duration"`

	// 视频拍摄时的选择角度，单位：度
	Rotate *int64 `json:"Rotate,omitempty" name:"Rotate"`

	// 视频流信息。
	VideoStreamInfoSet []*VideoStreamInfo `json:"VideoStreamInfoSet,omitempty" name:"VideoStreamInfoSet" list`

	// 音频流信息。
	AudioStreamInfoSet []*AudioStreamInfo `json:"AudioStreamInfoSet,omitempty" name:"AudioStreamInfoSet" list`
}

type ModifyMaterialRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 素材 Id。
	MaterialId *string `json:"MaterialId,omitempty" name:"MaterialId"`

	// 素材归属。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 素材名称，不能超过30个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 素材标签，单个标签长度不能超过10个字符，数组长度不能超过10。
	Tags []*string `json:"Tags,omitempty" name:"Tags" list`

	// 素材分类路径，例如填写"/a/b"，则代表该素材存储的路径为"/a/b"。
	ClassPath *string `json:"ClassPath,omitempty" name:"ClassPath"`

	// 操作者。填写用户的 Id，用于标识调用者及校验操作权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *ModifyMaterialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMaterialRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyMaterialResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMaterialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMaterialResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyProjectRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 项目 Id。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 项目名称，不可超过30个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 画布宽高比，取值有：
	// <li>16:9；</li>
	// <li>9:16。</li>
	AspectRatio *string `json:"AspectRatio,omitempty" name:"AspectRatio"`

	// 归属者。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`
}

func (r *ModifyProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyProjectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyProjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTeamMemberRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 团队 ID。
	TeamId *string `json:"TeamId,omitempty" name:"TeamId"`

	// 团队成员 ID。
	MemberId *string `json:"MemberId,omitempty" name:"MemberId"`

	// 成员备注，允许设置备注为空，不为空时长度不能超过15个字符。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 成员角色，取值：
	// <li>Admin：团队管理员；</li>
	// <li>Member：普通成员。</li>
	Role *string `json:"Role,omitempty" name:"Role"`

	// 操作者。填写用户的 Id，用于标识调用者及校验操作权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *ModifyTeamMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTeamMemberRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTeamMemberResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTeamMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTeamMemberResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTeamRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 团队 ID。
	TeamId *string `json:"TeamId,omitempty" name:"TeamId"`

	// 团队名称，不能超过 30 个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 操作者。填写用户的 Id，用于标识调用者及校验操作权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *ModifyTeamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTeamRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTeamResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTeamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTeamResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MoveClassRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 归属者。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 源分类路径。
	SourceClassPath *string `json:"SourceClassPath,omitempty" name:"SourceClassPath"`

	// 目标分类路径。
	DestinationClassPath *string `json:"DestinationClassPath,omitempty" name:"DestinationClassPath"`

	// 操作者。填写用户的 Id，用于标识调用者及校验操作权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *MoveClassRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MoveClassRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MoveClassResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MoveClassResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MoveClassResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ProjectInfo struct {

	// 项目 Id。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 项目名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 画布宽高比。
	AspectRatio *string `json:"AspectRatio,omitempty" name:"AspectRatio"`

	// 项目类别。
	Category *string `json:"Category,omitempty" name:"Category"`

	// 归属者。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 项目封面图片地址。
	CoverUrl *string `json:"CoverUrl,omitempty" name:"CoverUrl"`

	// 项目创建时间，格式按照 ISO 8601 标准表示。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 项目更新时间，格式按照 ISO 8601 标准表示。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type Resource struct {

	// 类型，取值有：
	// <li>MATERIAL：素材。</li>
	// <li>CLASS：分类。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 资源 Id，当 Type 为 MATERIAL 时，取值为素材 Id；当 Type 为 CLASS 时，取值为分类路径 ClassPath。
	Id *string `json:"Id,omitempty" name:"Id"`
}

type RevokeResourceAuthorizationRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 资源所属实体。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 被授权资源。
	Resources []*Resource `json:"Resources,omitempty" name:"Resources" list`

	// 被授权目标实体。
	Authorizees []*Entity `json:"Authorizees,omitempty" name:"Authorizees" list`

	// 详细授权值。 取值有：
	// <li>R：可读，可以浏览素材，但不能使用该素材（将其添加到 Project），或复制到自己的媒资库中</li>
	// <li>X：可用，可以使用该素材（将其添加到 Project），但不能将其复制到自己的媒资库中，意味着被授权者无法将该资源进一步扩散给其他个人或团队。</li>
	// <li>C：可复制，既可以使用该素材（将其添加到 Project），也可以将其复制到自己的媒资库中。</li>
	// <li>W：可修改、删除媒资。</li>
	Permissions []*string `json:"Permissions,omitempty" name:"Permissions" list`

	// 操作者。填写用户的 Id，用于标识调用者及校验操作权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *RevokeResourceAuthorizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RevokeResourceAuthorizationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RevokeResourceAuthorizationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RevokeResourceAuthorizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RevokeResourceAuthorizationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchMaterialRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 指定搜索空间，数组长度不得超过5。
	SearchScopes []*SearchScope `json:"SearchScopes,omitempty" name:"SearchScopes" list`

	// 素材类型，取值：
	// <li>AUDIO：音频；</li>
	// <li>VIDEO：视频 ；</li>
	// <li>IMAGE：图片。</li>
	MaterialTypes []*string `json:"MaterialTypes,omitempty" name:"MaterialTypes" list`

	// 搜索文本，模糊匹配素材名称或描述信息，匹配项越多，匹配度越高，排序越优先。长度限制：64 个字符。
	Text *string `json:"Text,omitempty" name:"Text"`

	// 按画质检索，取值为：LD/SD/HD/FHD/2K/4K。
	Resolution *string `json:"Resolution,omitempty" name:"Resolution"`

	// 按素材时长检索，单位s。
	DurationRange *IntegerRange `json:"DurationRange,omitempty" name:"DurationRange"`

	// 按照素材创建时间检索。
	CreateTimeRange *TimeRange `json:"CreateTimeRange,omitempty" name:"CreateTimeRange"`

	// 标签集合，匹配集合中任意元素。单个标签长度限制：10 个字符。数组长度限制：10。
	Tags []*string `json:"Tags,omitempty" name:"Tags" list`

	// 排序方式。Sort.Field 可选值：CreateTime。指定 Text 搜索时，将根据匹配度排序，该字段无效。
	Sort *SortBy `json:"Sort,omitempty" name:"Sort"`

	// 偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回记录条数，默认值：50。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 操作者。填写用户的 Id，用于标识调用者及校验操作权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *SearchMaterialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchMaterialRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchMaterialResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合记录总条数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 素材信息，仅返回基础信息。
		MaterialInfoSet []*MaterialInfo `json:"MaterialInfoSet,omitempty" name:"MaterialInfoSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchMaterialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchMaterialResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchScope struct {

	// 分类路径归属。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 按分类路径检索。 不填则默认按根分类路径检索。
	ClassPath *string `json:"ClassPath,omitempty" name:"ClassPath"`
}

type SortBy struct {

	// 排序字段。
	Field *string `json:"Field,omitempty" name:"Field"`

	// 排序方式，可选值：Asc（升序）、Desc（降序），默认降序。
	Order *string `json:"Order,omitempty" name:"Order"`
}

type TaskBaseInfo struct {

	// 任务 Id。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务类型，取值有：
	// <li>VIDEO_EDIT_PROJECT_EXPORT：项目导出。</li>
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 任务状态，取值有：
	// <li>PROCESSING：处理中：</li>
	// <li>SUCCESS：成功；</li>
	// <li>FAIL：失败。</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 任务进度，取值为：0~100。
	Progress *uint64 `json:"Progress,omitempty" name:"Progress"`

	// 错误码。
	// <li>0：成功；</li>
	// <li>其他值：失败。</li>
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 创建时间，格式按照 ISO 8601 标准表示。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type TeamInfo struct {

	// 团队 ID。
	TeamId *string `json:"TeamId,omitempty" name:"TeamId"`

	// 团队名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 团队成员个数
	MemberCount *uint64 `json:"MemberCount,omitempty" name:"MemberCount"`

	// 团队创建时间，格式按照 ISO 8601 标准表示。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 团队最后更新时间，格式按照 ISO 8601 标准表示。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type TeamMemberInfo struct {

	// 团队成员 ID。
	MemberId *string `json:"MemberId,omitempty" name:"MemberId"`

	// 团队成员备注。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 团队成员角色，取值：
	// <li>Owner：团队所有者，添加团队成员及修改团队成员解决时不能填此角色；</li>
	// <li>Admin：团队管理员；</li>
	// <li>Member：普通成员。</li>
	Role *string `json:"Role,omitempty" name:"Role"`
}

type TimeRange struct {

	// 开始时间，使用 ISO 日期格式。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，使用 ISO 日期格式。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type VODExportInfo struct {

	// 导出的媒资名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 导出的媒资分类 Id。
	ClassId *uint64 `json:"ClassId,omitempty" name:"ClassId"`
}

type VideoEditProjectOutput struct {

	// 云点播媒资 FileId。
	VodFileId *string `json:"VodFileId,omitempty" name:"VodFileId"`

	// 导出的媒资 URL。
	URL *string `json:"URL,omitempty" name:"URL"`

	// 元信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetaData *MediaMetaData `json:"MetaData,omitempty" name:"MetaData"`
}

type VideoMaterial struct {

	// 素材元信息。
	MetaData *MediaMetaData `json:"MetaData,omitempty" name:"MetaData"`

	// 雪碧图信息。
	ImageSpriteInfo *MediaImageSpriteInfo `json:"ImageSpriteInfo,omitempty" name:"ImageSpriteInfo"`

	// 素材媒体文件的 URL 地址
	MaterialUrl *string `json:"MaterialUrl,omitempty" name:"MaterialUrl"`

	// 素材媒体文件的封面图片地址。
	CoverUrl *string `json:"CoverUrl,omitempty" name:"CoverUrl"`

	// 媒体文件分辨率。取值为：LD/SD/HD/FHD/2K/4K。
	Resolution *string `json:"Resolution,omitempty" name:"Resolution"`

	// 素材状态。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaterialStatus *MaterialStatus `json:"MaterialStatus,omitempty" name:"MaterialStatus"`
}

type VideoStreamInfo struct {

	// 码率，单位：bps。
	Bitrate *uint64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 高度，单位：px。
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// 宽度，单位：px。
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// 编码格式。
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// 帧率，单位：hz。
	Fps *uint64 `json:"Fps,omitempty" name:"Fps"`
}
