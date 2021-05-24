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
    "errors"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AccountInfo struct {

	// 用户 Id。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户手机号码。
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 用户昵称。
	Nick *string `json:"Nick,omitempty" name:"Nick"`

	// 账号状态，取值：
	// <li>Normal：有效；</li>
	// <li>Stopped：无效。</li>
	Status *string `json:"Status,omitempty" name:"Status"`
}

type AddMemberInfo struct {

	// 团队成员 ID。
	MemberId *string `json:"MemberId,omitempty" name:"MemberId"`

	// 团队成员备注。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 团队成员角色，不填则默认添加普通成员。可选值：
	// <li>Admin：团队管理员；</li>
	// <li>Member：普通成员。</li>
	Role *string `json:"Role,omitempty" name:"Role"`
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddTeamMemberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "TeamId")
	delete(f, "TeamMembers")
	delete(f, "Operator")
	if len(f) > 0 {
		return errors.New("AddTeamMemberRequest has unknown keys!")
	}
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddTeamMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AudioMaterial struct {

	// 素材元信息。
	MetaData *MediaMetaData `json:"MetaData,omitempty" name:"MetaData"`

	// 素材媒体文件的播放 URL 地址。
	MaterialUrl *string `json:"MaterialUrl,omitempty" name:"MaterialUrl"`

	// 素材媒体文件的封面图片地址。
	CoverUrl *string `json:"CoverUrl,omitempty" name:"CoverUrl"`

	// 素材状态。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaterialStatus *MaterialStatus `json:"MaterialStatus,omitempty" name:"MaterialStatus"`

	// 素材媒体文件的原始 URL 地址。
	OriginalUrl *string `json:"OriginalUrl,omitempty" name:"OriginalUrl"`

	// 云点播媒资 FileId。
	VodFileId *string `json:"VodFileId,omitempty" name:"VodFileId"`
}

type AudioStreamInfo struct {

	// 码率，单位：bps。
	Bitrate *uint64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 采样率，单位：hz。
	SamplingRate *uint64 `json:"SamplingRate,omitempty" name:"SamplingRate"`

	// 编码格式。
	Codec *string `json:"Codec,omitempty" name:"Codec"`
}

type AudioTrackItem struct {

	// 音频媒体来源类型，取值有：
	// <ul>
	// <li>VOD ：素材来源于云点播文件 ；</li>
	// <li>CME ：视频来源于制作云媒体文件 ；</li>
	// <li>EXTERNAL ：视频来源于媒资绑定。</li>
	// </ul>
	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`

	// 音频片段的媒体来源，可以是：
	// <ul>
	// <li>当 SourceType 为 VOD 时，为云点播的媒体文件 FileId ，会默认将该 FileId 导入到项目中 ；</li>
	// <li>当 SourceType 为 CME 时，为制作云的媒体 ID，项目归属者必须对该云媒资有访问权限；</li>
	// <li>当 SourceType 为 EXTERNAL 时，为媒资绑定的 Definition 与 MediaKey 中间用冒号分隔合并后的字符串，格式为 Definition:MediaKey 。</li>
	// </ul>
	SourceMedia *string `json:"SourceMedia,omitempty" name:"SourceMedia"`

	// 音频片段取自媒体文件的起始时间，单位为秒。0 表示从媒体开始位置截取。默认为0。
	SourceMediaStartTime *float64 `json:"SourceMediaStartTime,omitempty" name:"SourceMediaStartTime"`

	// 音频片段的时长，单位为秒。默认和媒体本身长度一致，表示截取全部媒体。
	Duration *float64 `json:"Duration,omitempty" name:"Duration"`
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

	// 导出媒体归属，个人或团队。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 导出的媒体名称，不得超过30个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 导出的媒体信息，不得超过50个字符。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 导出的媒体分类路径，长度不能超过15字符。
	ClassPath *string `json:"ClassPath,omitempty" name:"ClassPath"`

	// 导出的媒体标签，单个标签不得超过10个字符。
	TagSet []*string `json:"TagSet,omitempty" name:"TagSet" list`

	// 第三方平台发布信息列表。暂未正式对外，请勿使用。
	ThirdPartyPublishInfos []*ThirdPartyPublishInfo `json:"ThirdPartyPublishInfos,omitempty" name:"ThirdPartyPublishInfos" list`
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

	// 操作者。填写用户的 Id，用于标识调用者及校验分类创建权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *CreateClassRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClassRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "Owner")
	delete(f, "ClassPath")
	delete(f, "Operator")
	if len(f) > 0 {
		return errors.New("CreateClassRequest has unknown keys!")
	}
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClassResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLinkRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 链接类型，取值有:
	// <li>CLASS: 分类链接；</li>
	// <li> MATERIAL：媒体文件链接。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 链接名称，不能超过30个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 链接归属者。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 目标资源Id。取值：
	// <li>当 Type 为 MATERIAL 时填媒体 ID；</li>
	// <li>当 Type 为 CLASS 时填写分类路径。</li>
	DestinationId *string `json:"DestinationId,omitempty" name:"DestinationId"`

	// 目标资源归属者。
	DestinationOwner *Entity `json:"DestinationOwner,omitempty" name:"DestinationOwner"`

	// 链接的分类路径，如填"/a/b"则代表链接属于该分类路径，不填则默认为根路径。
	ClassPath *string `json:"ClassPath,omitempty" name:"ClassPath"`

	// 操作者。填写用户的 Id，用于标识调用者及校验操作权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *CreateLinkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLinkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "Type")
	delete(f, "Name")
	delete(f, "Owner")
	delete(f, "DestinationId")
	delete(f, "DestinationOwner")
	delete(f, "ClassPath")
	delete(f, "Operator")
	if len(f) > 0 {
		return errors.New("CreateLinkRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateLinkResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 新建链接的媒体 Id。
		MaterialId *string `json:"MaterialId,omitempty" name:"MaterialId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLinkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLinkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProjectRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 项目名称，不可超过30个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 项目归属者。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 项目类别，取值有：
	// <li>VIDEO_EDIT：视频编辑。</li>
	// <li>SWITCHER：导播台。</li>
	// <li>VIDEO_SEGMENTATION：视频拆条。</li>
	// <li>STREAM_CONNECT：云转推。</li>
	// <li>RECORD_REPLAY：录制回放。</li>
	Category *string `json:"Category,omitempty" name:"Category"`

	// 项目模式，一个项目可以有多种模式并相互切换。
	// 当 Category 为 VIDEO_EDIT 时，可选模式有：
	// <li>Default：默认模式。</li>
	// <li>VideoEditTemplate：视频编辑模板制作模式。</li>
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// 画布宽高比。
	// 该字段已经废弃，请使用具体项目输入中的 AspectRatio 字段。
	AspectRatio *string `json:"AspectRatio,omitempty" name:"AspectRatio"`

	// 项目描述信息。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 导播台信息，仅当项目类型为 SWITCHER 时必填。
	SwitcherProjectInput *SwitcherProjectInput `json:"SwitcherProjectInput,omitempty" name:"SwitcherProjectInput"`

	// 直播剪辑信息，暂未开放，请勿使用。
	LiveStreamClipProjectInput *LiveStreamClipProjectInput `json:"LiveStreamClipProjectInput,omitempty" name:"LiveStreamClipProjectInput"`

	// 视频编辑信息，仅当项目类型为 VIDEO_EDIT 时必填。
	VideoEditProjectInput *VideoEditProjectInput `json:"VideoEditProjectInput,omitempty" name:"VideoEditProjectInput"`

	// 视频拆条信息，仅当项目类型为 VIDEO_SEGMENTATION  时必填。
	VideoSegmentationProjectInput *VideoSegmentationProjectInput `json:"VideoSegmentationProjectInput,omitempty" name:"VideoSegmentationProjectInput"`

	// 云转推项目信息，仅当项目类型为 STREAM_CONNECT 时必填。
	StreamConnectProjectInput *StreamConnectProjectInput `json:"StreamConnectProjectInput,omitempty" name:"StreamConnectProjectInput"`

	// 录制回放项目信息，仅当项目类型为 RECORD_REPLAY 时必填。
	RecordReplayProjectInput *RecordReplayProjectInput `json:"RecordReplayProjectInput,omitempty" name:"RecordReplayProjectInput"`
}

func (r *CreateProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "Name")
	delete(f, "Owner")
	delete(f, "Category")
	delete(f, "Mode")
	delete(f, "AspectRatio")
	delete(f, "Description")
	delete(f, "SwitcherProjectInput")
	delete(f, "LiveStreamClipProjectInput")
	delete(f, "VideoEditProjectInput")
	delete(f, "VideoSegmentationProjectInput")
	delete(f, "StreamConnectProjectInput")
	delete(f, "RecordReplayProjectInput")
	if len(f) > 0 {
		return errors.New("CreateProjectRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 项目 Id。
		ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

		// 输入源推流信息。
	//  <li> 当 Catagory 为 STREAM_CONNECT 时，数组返回长度为 2 ，第 0 个代表主输入源，第 1 个代表备输入源。只有当各自输入源类型为推流时才有有效内容。</li>
		RtmpPushInputInfoSet []*RtmpPushInputInfo `json:"RtmpPushInputInfoSet,omitempty" name:"RtmpPushInputInfoSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTeamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "Name")
	delete(f, "OwnerId")
	delete(f, "OwnerRemark")
	delete(f, "TeamId")
	if len(f) > 0 {
		return errors.New("CreateTeamRequest has unknown keys!")
	}
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClassRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "Owner")
	delete(f, "ClassPath")
	delete(f, "Operator")
	if len(f) > 0 {
		return errors.New("DeleteClassRequest has unknown keys!")
	}
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoginStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "UserIds")
	if len(f) > 0 {
		return errors.New("DeleteLoginStatusRequest has unknown keys!")
	}
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoginStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMaterialRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 媒体 Id。
	MaterialId *string `json:"MaterialId,omitempty" name:"MaterialId"`

	// 操作者。填写用户的 Id，用于标识调用者及校验媒体删除权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *DeleteMaterialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMaterialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "MaterialId")
	delete(f, "Operator")
	if len(f) > 0 {
		return errors.New("DeleteMaterialRequest has unknown keys!")
	}
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMaterialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProjectRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 项目 Id。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 操作者。填写用户的 Id，用于标识调用者及校验对项目删除操作权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *DeleteProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "ProjectId")
	delete(f, "Operator")
	if len(f) > 0 {
		return errors.New("DeleteProjectRequest has unknown keys!")
	}
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTeamMembersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "TeamId")
	delete(f, "MemberIds")
	delete(f, "Operator")
	if len(f) > 0 {
		return errors.New("DeleteTeamMembersRequest has unknown keys!")
	}
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTeamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "TeamId")
	delete(f, "Operator")
	if len(f) > 0 {
		return errors.New("DeleteTeamRequest has unknown keys!")
	}
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTeamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountsRequest struct {
	*tchttp.BaseRequest

	// 平台唯一标识。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 手机号码。
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 分页返回的起始偏移量，默认值：0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页返回的记录条数，默认值：10，最大值：20。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "Phone")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return errors.New("DescribeAccountsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合搜索条件的记录总数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 账号信息列表。
		AccountInfoSet []*AccountInfo `json:"AccountInfoSet,omitempty" name:"AccountInfoSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountsResponse) FromJsonString(s string) error {
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClassRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "Owner")
	delete(f, "Operator")
	if len(f) > 0 {
		return errors.New("DescribeClassRequest has unknown keys!")
	}
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJoinTeamsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "MemberId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return errors.New("DescribeJoinTeamsRequest has unknown keys!")
	}
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoginStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "UserIds")
	if len(f) > 0 {
		return errors.New("DescribeLoginStatusRequest has unknown keys!")
	}
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoginStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMaterialsRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 媒体 ID 列表，N 从 0 开始取值，最大 19。
	MaterialIds []*string `json:"MaterialIds,omitempty" name:"MaterialIds" list`

	// 列表排序，支持下列排序字段：
	// <li>CreateTime：创建时间；</li>
	// <li>UpdateTime：更新时间。</li>
	Sort *SortBy `json:"Sort,omitempty" name:"Sort"`

	// 操作者。填写用户的 Id，用于标识调用者及校验媒体的访问权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *DescribeMaterialsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMaterialsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "MaterialIds")
	delete(f, "Sort")
	delete(f, "Operator")
	if len(f) > 0 {
		return errors.New("DescribeMaterialsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMaterialsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 媒体列表信息。
		MaterialInfoSet []*MaterialInfo `json:"MaterialInfoSet,omitempty" name:"MaterialInfoSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMaterialsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMaterialsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePlatformsRequest struct {
	*tchttp.BaseRequest

	// 平台集合。
	Platforms []*string `json:"Platforms,omitempty" name:"Platforms" list`

	// 平台绑定的 license Id 集合。
	LicenseIds []*string `json:"LicenseIds,omitempty" name:"LicenseIds" list`

	// 分页返回的起始偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页返回的记录条数，默认值：10。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribePlatformsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePlatformsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platforms")
	delete(f, "LicenseIds")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return errors.New("DescribePlatformsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePlatformsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合搜索条件的记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 平台信息列表。
		PlatformInfoSet []*PlatformInfo `json:"PlatformInfoSet,omitempty" name:"PlatformInfoSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePlatformsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePlatformsResponse) FromJsonString(s string) error {
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

	// 项目类别，取值有：
	// <li>VIDEO_EDIT：视频编辑。</li>
	// <li>SWITCHER：导播台。</li>
	// <li>VIDEO_SEGMENTATION：视频拆条。</li>
	// <li>STREAM_CONNECT：云转推。</li>
	// <li>RECORD_REPLAY：录制回放。</li>
	CategorySet []*string `json:"CategorySet,omitempty" name:"CategorySet" list`

	// 项目模式，一个项目可以有多种模式并相互切换。
	// 当 Category 为 VIDEO_EDIT 时，可选模式有：
	// <li>Default：默认模式。</li>
	// <li>VideoEditTemplate：视频编辑模板制作模式。</li>
	Modes []*string `json:"Modes,omitempty" name:"Modes" list`

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

	// 操作者。填写用户的 Id，用于标识调用者及校验项目访问权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *DescribeProjectsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "ProjectIds")
	delete(f, "AspectRatioSet")
	delete(f, "CategorySet")
	delete(f, "Modes")
	delete(f, "Sort")
	delete(f, "Owner")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Operator")
	if len(f) > 0 {
		return errors.New("DescribeProjectsRequest has unknown keys!")
	}
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceAuthorizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "Owner")
	delete(f, "Resource")
	delete(f, "Operator")
	if len(f) > 0 {
		return errors.New("DescribeResourceAuthorizationRequest has unknown keys!")
	}
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceAuthorizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSharedSpaceRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 被授权目标,，个人或团队。
	Authorizee *Entity `json:"Authorizee,omitempty" name:"Authorizee"`

	// 操作者。填写用户的 Id，用于标识调用者及校验操作权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *DescribeSharedSpaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSharedSpaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "Authorizee")
	delete(f, "Operator")
	if len(f) > 0 {
		return errors.New("DescribeSharedSpaceRequest has unknown keys!")
	}
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSharedSpaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskDetailRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 任务 Id。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 操作者。填写用户的 Id，用于标识调用者及校验对任务的访问权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *DescribeTaskDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "TaskId")
	delete(f, "Operator")
	if len(f) > 0 {
		return errors.New("DescribeTaskDetailRequest has unknown keys!")
	}
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

	// 操作者。填写用户的 Id，用于标识调用者及校验对任务的访问权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *DescribeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "ProjectId")
	delete(f, "TaskTypeSet")
	delete(f, "StatusSet")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Operator")
	if len(f) > 0 {
		return errors.New("DescribeTasksRequest has unknown keys!")
	}
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTeamMembersRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 团队 ID。
	TeamId *string `json:"TeamId,omitempty" name:"TeamId"`

	// 成员 ID 列表，限指定30个指定成员。如不填，则返回指定团队下的所有成员。
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTeamMembersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "TeamId")
	delete(f, "MemberIds")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Operator")
	if len(f) > 0 {
		return errors.New("DescribeTeamMembersRequest has unknown keys!")
	}
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTeamMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTeamsRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 团队 ID 列表，限30个。若不填，则默认获取平台下所有团队。
	TeamIds []*string `json:"TeamIds,omitempty" name:"TeamIds" list`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回记录条数，默认值：20，最大值：30。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTeamsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTeamsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "TeamIds")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return errors.New("DescribeTeamsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTeamsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTeamsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EmptyTrackItem struct {

	// 持续时间，单位为秒。
	Duration *float64 `json:"Duration,omitempty" name:"Duration"`
}

type Entity struct {

	// 类型，取值有：
	// <li>PERSON：个人。</li>
	// <li>TEAM：团队。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Id，当 Type=PERSON，取值为用户 Id，当 Type=TEAM，取值为团队 Id。
	Id *string `json:"Id,omitempty" name:"Id"`
}

type EventContent struct {

	// 事件类型，可取值为：
	// <li>Storage.NewFileCreated：新文件产生。</li>
	EventType *string `json:"EventType,omitempty" name:"EventType"`

	// 新文件产生事件信息。仅当 EventType 为 Storage.NewFileCreated 时有效。
	StorageNewFileCreatedEvent *StorageNewFileCreatedEvent `json:"StorageNewFileCreatedEvent,omitempty" name:"StorageNewFileCreatedEvent"`
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

	// 在线编辑轨道数据。轨道数据相关介绍，请查看 [视频合成协议](https://cloud.tencent.com/document/product/1156/51225)。
	TrackData *string `json:"TrackData,omitempty" name:"TrackData"`

	// 导出的云剪素材信息。指定 ExportDestination = CME 时有效。
	CMEExportInfo *CMEExportInfo `json:"CMEExportInfo,omitempty" name:"CMEExportInfo"`

	// 导出的云点播媒资信息。指定 ExportDestination = VOD 时有效。
	VODExportInfo *VODExportInfo `json:"VODExportInfo,omitempty" name:"VODExportInfo"`

	// 操作者。填写用户的 Id，用于标识调用者及校验导出操作权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *ExportVideoByEditorTrackDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportVideoByEditorTrackDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "Definition")
	delete(f, "ExportDestination")
	delete(f, "TrackData")
	delete(f, "CMEExportInfo")
	delete(f, "VODExportInfo")
	delete(f, "Operator")
	if len(f) > 0 {
		return errors.New("ExportVideoByEditorTrackDataRequest has unknown keys!")
	}
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportVideoByEditorTrackDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportVideoByTemplateRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 视频编辑模板  Id。
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 导出模板 Id，目前不支持自定义创建，只支持下面的预置模板 Id。
	// <li>10：分辨率为 480P，输出视频格式为 MP4；</li>
	// <li>11：分辨率为 720P，输出视频格式为 MP4；</li>
	// <li>12：分辨率为 1080P，输出视频格式为 MP4。</li>
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 导出目标，可取值为：
	// <li>CME：云剪，即导出为云剪媒体；</li>
	// <li>VOD：云点播，即导出为云点播媒资。</li>
	ExportDestination *string `json:"ExportDestination,omitempty" name:"ExportDestination"`

	// 需要替换的素材信息。
	SlotReplacements []*SlotReplacementInfo `json:"SlotReplacements,omitempty" name:"SlotReplacements" list`

	// 导出的云剪媒体信息。指定 ExportDestination = CME 时有效。
	CMEExportInfo *CMEExportInfo `json:"CMEExportInfo,omitempty" name:"CMEExportInfo"`

	// 导出的云点播媒资信息。指定 ExportDestination = VOD 时有效。
	VODExportInfo *VODExportInfo `json:"VODExportInfo,omitempty" name:"VODExportInfo"`

	// 操作者。填写用户的 Id，用于标识调用者及校验项目导出权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *ExportVideoByTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportVideoByTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "TemplateId")
	delete(f, "Definition")
	delete(f, "ExportDestination")
	delete(f, "SlotReplacements")
	delete(f, "CMEExportInfo")
	delete(f, "VODExportInfo")
	delete(f, "Operator")
	if len(f) > 0 {
		return errors.New("ExportVideoByTemplateRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ExportVideoByTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 导出任务 Id。
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportVideoByTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportVideoByTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportVideoByVideoSegmentationDataRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 视频拆条项目 Id 。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 指定需要导出的智能拆条片段的组 Id 。
	SegmentGroupId *string `json:"SegmentGroupId,omitempty" name:"SegmentGroupId"`

	// 指定需要导出的智能拆条片段 Id  集合。
	SegmentIds []*string `json:"SegmentIds,omitempty" name:"SegmentIds" list`

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

	// 操作者。填写用户的 Id，用于标识调用者及校验操作权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *ExportVideoByVideoSegmentationDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportVideoByVideoSegmentationDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "ProjectId")
	delete(f, "SegmentGroupId")
	delete(f, "SegmentIds")
	delete(f, "Definition")
	delete(f, "ExportDestination")
	delete(f, "CMEExportInfo")
	delete(f, "VODExportInfo")
	delete(f, "Operator")
	if len(f) > 0 {
		return errors.New("ExportVideoByVideoSegmentationDataRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ExportVideoByVideoSegmentationDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务 Id。
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportVideoByVideoSegmentationDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportVideoByVideoSegmentationDataResponse) FromJsonString(s string) error {
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
	// <li>CME：云剪，即导出为云剪媒体；</li>
	// <li>VOD：云点播，即导出为云点播媒资。</li>
	ExportDestination *string `json:"ExportDestination,omitempty" name:"ExportDestination"`

	// 视频封面图片文件（如 jpeg, png 等）进行 Base64 编码后的字符串，仅支持 gif、jpeg、png 三种图片格式，原图片文件不能超过2 M大 小。
	CoverData *string `json:"CoverData,omitempty" name:"CoverData"`

	// 导出的云剪媒体信息。指定 ExportDestination = CME 时有效。
	CMEExportInfo *CMEExportInfo `json:"CMEExportInfo,omitempty" name:"CMEExportInfo"`

	// 导出的云点播媒资信息。指定 ExportDestination = VOD 时有效。
	VODExportInfo *VODExportInfo `json:"VODExportInfo,omitempty" name:"VODExportInfo"`

	// 操作者。填写用户的 Id，用于标识调用者及校验项目导出权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *ExportVideoEditProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportVideoEditProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "ProjectId")
	delete(f, "Definition")
	delete(f, "ExportDestination")
	delete(f, "CoverData")
	delete(f, "CMEExportInfo")
	delete(f, "VODExportInfo")
	delete(f, "Operator")
	if len(f) > 0 {
		return errors.New("ExportVideoEditProjectRequest has unknown keys!")
	}
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportVideoEditProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExternalMediaInfo struct {

	// 媒资绑定模板 ID。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 媒资绑定媒体路径或文件 ID。
	MediaKey *string `json:"MediaKey,omitempty" name:"MediaKey"`
}

type FlattenListMediaRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 媒体分类路径，例如填写"/a/b"，则代表平铺该分类路径下及其子分类路径下的媒体信息。
	ClassPath *string `json:"ClassPath,omitempty" name:"ClassPath"`

	// 媒体分类的归属者。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 分页偏移量，默认值：0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：50。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 操作者。填写用户的 Id，用于标识调用者及校验媒体访问权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *FlattenListMediaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FlattenListMediaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "ClassPath")
	delete(f, "Owner")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Operator")
	if len(f) > 0 {
		return errors.New("FlattenListMediaRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type FlattenListMediaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的记录总数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 该分类路径下及其子分类下的所有媒体基础信息列表。
		MaterialInfoSet []*MaterialInfo `json:"MaterialInfoSet,omitempty" name:"MaterialInfoSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *FlattenListMediaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FlattenListMediaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GenerateVideoSegmentationSchemeByAiRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 视频拆条项目 Id 。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 操作者。填写用户的 Id，用于标识调用者及校验操作权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *GenerateVideoSegmentationSchemeByAiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateVideoSegmentationSchemeByAiRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "ProjectId")
	delete(f, "Operator")
	if len(f) > 0 {
		return errors.New("GenerateVideoSegmentationSchemeByAiRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GenerateVideoSegmentationSchemeByAiResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 视频智能拆条任务 Id 。
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GenerateVideoSegmentationSchemeByAiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateVideoSegmentationSchemeByAiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GrantResourceAuthorizationRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 资源归属者，个人或者团队。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 被授权资源。
	Resources []*Resource `json:"Resources,omitempty" name:"Resources" list`

	// 被授权目标，个人或者团队。
	Authorizees []*Entity `json:"Authorizees,omitempty" name:"Authorizees" list`

	// 详细授权值。 取值有：
	// <li>R：可读，可以浏览媒体，但不能使用该媒体文件（将其添加到 Project），或复制到自己的媒资库中</li>
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GrantResourceAuthorizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "Owner")
	delete(f, "Resources")
	delete(f, "Authorizees")
	delete(f, "Permissions")
	delete(f, "Operator")
	if len(f) > 0 {
		return errors.New("GrantResourceAuthorizationRequest has unknown keys!")
	}
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GrantResourceAuthorizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HandleStreamConnectProjectRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 云转推项目Id 。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 请参考 [操作类型](#Operation)
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 转推输入源操作参数。具体操作方式详见 [操作类型](#Operation) 及下文示例。
	InputInfo *StreamInputInfo `json:"InputInfo,omitempty" name:"InputInfo"`

	// 主备输入源标识，取值有：
	// <li> Main ：主源；</li>
	// <li> Backup ：备源。</li>
	InputEndpoint *string `json:"InputEndpoint,omitempty" name:"InputEndpoint"`

	// 转推输出源操作参数。具体操作方式详见 [操作类型](#Operation) 及下文示例。
	OutputInfo *StreamConnectOutput `json:"OutputInfo,omitempty" name:"OutputInfo"`

	// 云转推当前预计结束时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。具体操作方式详见 [操作类型](#Operation) 及下文示例。
	CurrentStopTime *string `json:"CurrentStopTime,omitempty" name:"CurrentStopTime"`
}

func (r *HandleStreamConnectProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *HandleStreamConnectProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "ProjectId")
	delete(f, "Operation")
	delete(f, "InputInfo")
	delete(f, "InputEndpoint")
	delete(f, "OutputInfo")
	delete(f, "CurrentStopTime")
	if len(f) > 0 {
		return errors.New("HandleStreamConnectProjectRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type HandleStreamConnectProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 输入源推流地址，当 Operation 取值 AddInput 且 InputType 为 RtmpPush 类型时有效。
		StreamInputRtmpPushUrl *string `json:"StreamInputRtmpPushUrl,omitempty" name:"StreamInputRtmpPushUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *HandleStreamConnectProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *HandleStreamConnectProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageMaterial struct {

	// 图片高度，单位：px。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 图片宽度，单位：px。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 素材媒体文件的展示 URL 地址。
	MaterialUrl *string `json:"MaterialUrl,omitempty" name:"MaterialUrl"`

	// 图片大小，单位：字节。
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// 素材媒体文件的原始 URL 地址。
	OriginalUrl *string `json:"OriginalUrl,omitempty" name:"OriginalUrl"`

	// 云点播媒资 FileId。
	VodFileId *string `json:"VodFileId,omitempty" name:"VodFileId"`
}

type ImportMaterialRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 媒体归属者，团队或个人。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 媒体名称，不能超过30个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 导入媒资类型，取值：
	// <li>VOD：云点播文件；</li>
	// <li>EXTERNAL：媒资绑定。</li>
	// 注意：如果不填默认为云点播文件。
	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`

	// 云点播媒资 FileId，仅当 SourceType 为 VOD 时有效。
	VodFileId *string `json:"VodFileId,omitempty" name:"VodFileId"`

	// 原始媒资文件信息，当 SourceType 取值 EXTERNAL 的时候必填。
	ExternalMediaInfo *ExternalMediaInfo `json:"ExternalMediaInfo,omitempty" name:"ExternalMediaInfo"`

	// 媒体分类路径，形如："/a/b"，层级数不能超过10，每个层级长度不能超过15字符。若不填则默认为根路径。
	ClassPath *string `json:"ClassPath,omitempty" name:"ClassPath"`

	// 媒体预处理任务模板 ID。取值：
	// <li>10：进行编辑预处理。</li>
	PreProcessDefinition *int64 `json:"PreProcessDefinition,omitempty" name:"PreProcessDefinition"`

	// 操作者。填写用户的 Id，用于标识调用者及校验操作权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *ImportMaterialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportMaterialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "Owner")
	delete(f, "Name")
	delete(f, "SourceType")
	delete(f, "VodFileId")
	delete(f, "ExternalMediaInfo")
	delete(f, "ClassPath")
	delete(f, "PreProcessDefinition")
	delete(f, "Operator")
	if len(f) > 0 {
		return errors.New("ImportMaterialRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ImportMaterialResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 媒体 Id。
		MaterialId *string `json:"MaterialId,omitempty" name:"MaterialId"`

		// 媒体文预处理任务 ID，如果未指定发起预处理任务则为空。
		PreProcessTaskId *string `json:"PreProcessTaskId,omitempty" name:"PreProcessTaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportMaterialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportMaterialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportMediaToProjectRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 项目 Id。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 导入媒资类型，取值：
	// <li>VOD：云点播文件；</li>
	// <li>EXTERNAL：媒资绑定。</li>
	// 注意：如果不填默认为云点播文件。
	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`

	// 云点播媒资文件 Id，当 SourceType 取值 VOD 或者缺省的时候必填。
	VodFileId *string `json:"VodFileId,omitempty" name:"VodFileId"`

	// 原始媒资文件信息，当 SourceType 取值 EXTERNAL 的时候必填。
	ExternalMediaInfo *ExternalMediaInfo `json:"ExternalMediaInfo,omitempty" name:"ExternalMediaInfo"`

	// 媒体名称，不能超过30个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 媒体预处理任务模板 ID，取值：
	// <li>10：进行编辑预处理。</li>
	// 注意：如果填0则不进行处理。
	PreProcessDefinition *int64 `json:"PreProcessDefinition,omitempty" name:"PreProcessDefinition"`

	// 操作者。填写用户的 Id，用于标识调用者及校验项目和媒体文件访问权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *ImportMediaToProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportMediaToProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "ProjectId")
	delete(f, "SourceType")
	delete(f, "VodFileId")
	delete(f, "ExternalMediaInfo")
	delete(f, "Name")
	delete(f, "PreProcessDefinition")
	delete(f, "Operator")
	if len(f) > 0 {
		return errors.New("ImportMediaToProjectRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ImportMediaToProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 媒体 Id。
		MaterialId *string `json:"MaterialId,omitempty" name:"MaterialId"`

		// 媒体预处理任务 ID，如果未指定发起预处理任务则为空。
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportMediaToProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

type KuaishouPublishInfo struct {

	// 视频发布标题，限30个字符。
	Title *string `json:"Title,omitempty" name:"Title"`
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

	// 媒体分类路径，例如填写"/a/b"，则代表浏览该分类路径下的媒体和子分类信息。
	ClassPath *string `json:"ClassPath,omitempty" name:"ClassPath"`

	// 媒体和分类的归属者。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 分页偏移量，默认值：0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：50。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 操作者。填写用户的 Id，用于标识调用者及校验对媒体的访问权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *ListMediaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListMediaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "ClassPath")
	delete(f, "Owner")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Operator")
	if len(f) > 0 {
		return errors.New("ListMediaRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListMediaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的媒体记录总数。
		MaterialTotalCount *int64 `json:"MaterialTotalCount,omitempty" name:"MaterialTotalCount"`

		// 浏览分类路径下的媒体列表信息。
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListMediaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LivePullInputInfo struct {

	// 直播拉流地址。
	InputUrl *string `json:"InputUrl,omitempty" name:"InputUrl"`
}

type LiveStreamClipProjectInput struct {

	// 直播流播放地址，目前仅支持 HLS 和 FLV 格式。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 直播流录制时长，单位为秒，最大值为 7200。
	StreamRecordDuration *uint64 `json:"StreamRecordDuration,omitempty" name:"StreamRecordDuration"`
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

	// 媒体 Id。
	MaterialId *string `json:"MaterialId,omitempty" name:"MaterialId"`

	// 媒体类型，取值为：
	// <li> AUDIO :音频;</li>
	// <li> VIDEO :视频;</li>
	// <li> IMAGE :图片;</li>
	// <li> LINK  :链接.</li>
	// <li> OTHER : 其他.</li>
	MaterialType *string `json:"MaterialType,omitempty" name:"MaterialType"`

	// 媒体归属实体。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 媒体名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 媒体文件的创建时间，使用 ISO 日期格式。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 媒体文件的最近更新时间（如修改视频属性、发起视频处理等会触发更新媒体文件信息的操作），使用 ISO 日期格式。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 媒体的分类路径。
	ClassPath *string `json:"ClassPath,omitempty" name:"ClassPath"`

	// 预置标签列表。
	PresetTagSet []*PresetTagInfo `json:"PresetTagSet,omitempty" name:"PresetTagSet" list`

	// 人工标签列表。
	TagSet []*string `json:"TagSet,omitempty" name:"TagSet" list`

	// 媒体文件的预览图。
	PreviewUrl *string `json:"PreviewUrl,omitempty" name:"PreviewUrl"`

	// 媒体绑定的标签信息列表 。
	// 该字段已废弃。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagInfoSet []*MaterialTagInfo `json:"TagInfoSet,omitempty" name:"TagInfoSet" list`
}

type MaterialInfo struct {

	// 媒体基本信息。
	BasicInfo *MaterialBasicInfo `json:"BasicInfo,omitempty" name:"BasicInfo"`

	// 视频媒体信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoMaterial *VideoMaterial `json:"VideoMaterial,omitempty" name:"VideoMaterial"`

	// 音频媒体信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioMaterial *AudioMaterial `json:"AudioMaterial,omitempty" name:"AudioMaterial"`

	// 图片媒体信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageMaterial *ImageMaterial `json:"ImageMaterial,omitempty" name:"ImageMaterial"`

	// 链接媒体信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LinkMaterial *LinkMaterial `json:"LinkMaterial,omitempty" name:"LinkMaterial"`

	// 模板媒体信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoEditTemplateMaterial *VideoEditTemplateMaterial `json:"VideoEditTemplateMaterial,omitempty" name:"VideoEditTemplateMaterial"`

	// 其他类型媒体信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OtherMaterial *OtherMaterial `json:"OtherMaterial,omitempty" name:"OtherMaterial"`
}

type MaterialStatus struct {

	// 素材编辑可用状态，取值有：
	// <li>NORMAL：正常，可直接用于编辑；</li>
	// <li>ABNORMAL : 异常，不可用于编辑；</li>
	// <li>PROCESSING：处理中，暂不可用于编辑。</li>
	EditorUsableStatus *string `json:"EditorUsableStatus,omitempty" name:"EditorUsableStatus"`
}

type MaterialTagInfo struct {

	// 标签类型，取值为：
	// <li>PRESET：预置标签；</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 标签 Id 。当标签类型为 PRESET 时，标签 Id 为预置标签 Id 。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 标签名称。
	Name *string `json:"Name,omitempty" name:"Name"`
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

type MediaReplacementInfo struct {

	// 素材 ID。
	MaterialId *string `json:"MaterialId,omitempty" name:"MaterialId"`

	// 替换媒体选取的开始时间，单位为秒，默认为 0。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`
}

type MediaTrack struct {

	// 轨道类型，取值有：
	// <ul>
	// <li>Video ：视频轨道。视频轨道由以下 Item 组成：<ul><li>VideoTrackItem</li><li>EmptyTrackItem</li><li>MediaTransitionItem</li></ul> </li>
	// <li>Audio ：音频轨道。音频轨道由以下 Item 组成：<ul><li>AudioTrackItem</li><li>EmptyTrackItem</li></ul> </li>
	// </ul>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 轨道上的媒体片段列表。
	TrackItems []*MediaTrackItem `json:"TrackItems,omitempty" name:"TrackItems" list`
}

type MediaTrackItem struct {

	// 片段类型。取值有：
	// <li>Video：视频片段；</li>
	// <li>Audio：音频片段；</li>
	// <li>Empty：空白片段；</li>
	// <li>Transition：转场。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 视频片段，当 Type = Video 时有效。
	VideoItem *VideoTrackItem `json:"VideoItem,omitempty" name:"VideoItem"`

	// 音频片段，当 Type = Audio 时有效。
	AudioItem *AudioTrackItem `json:"AudioItem,omitempty" name:"AudioItem"`

	// 空白片段，当 Type = Empty 时有效。空片段用于时间轴的占位。<li>如需要两个音频片段之间有一段时间的静音，可以用 EmptyTrackItem 来进行占位。</li>
	// <li>使用 EmptyTrackItem 进行占位，来定位某个Item。</li>
	EmptyItem *EmptyTrackItem `json:"EmptyItem,omitempty" name:"EmptyItem"`

	// 转场，当 Type = Transition 时有效。
	TransitionItem *MediaTransitionItem `json:"TransitionItem,omitempty" name:"TransitionItem"`
}

type MediaTransitionItem struct {

	// 转场 Id 。暂只支持一个转场。
	TransitionId *string `json:"TransitionId,omitempty" name:"TransitionId"`

	// 转场持续时间，单位为秒，默认为2秒。进行转场处理的两个媒体片段，第二个片段在轨道上的起始时间会自动进行调整，设置为前面一个片段的结束时间减去转场的持续时间。
	Duration *float64 `json:"Duration,omitempty" name:"Duration"`
}

type ModifyMaterialRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 媒体 Id。
	MaterialId *string `json:"MaterialId,omitempty" name:"MaterialId"`

	// 媒体或分类路径归属。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 媒体名称，不能超过30个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 媒体分类路径，例如填写"/a/b"，则代表该媒体存储的路径为"/a/b"。若修改分类路径，则 Owner 字段必填。
	ClassPath *string `json:"ClassPath,omitempty" name:"ClassPath"`

	// 操作者。填写用户的 Id，用于标识调用者及校验操作权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *ModifyMaterialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMaterialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "MaterialId")
	delete(f, "Owner")
	delete(f, "Name")
	delete(f, "ClassPath")
	delete(f, "Operator")
	if len(f) > 0 {
		return errors.New("ModifyMaterialRequest has unknown keys!")
	}
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

	// 项目归属者。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 项目模式，一个项目可以有多种模式并相互切换。
	// 当 Category 为 VIDEO_EDIT 时，可选模式有：
	// <li>Defualt：默认模式。</li>
	// <li>VideoEditTemplate：视频编辑模板制作模式。</li>
	Mode *string `json:"Mode,omitempty" name:"Mode"`
}

func (r *ModifyProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "ProjectId")
	delete(f, "Name")
	delete(f, "AspectRatio")
	delete(f, "Owner")
	delete(f, "Mode")
	if len(f) > 0 {
		return errors.New("ModifyProjectRequest has unknown keys!")
	}
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTeamMemberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "TeamId")
	delete(f, "MemberId")
	delete(f, "Remark")
	delete(f, "Role")
	delete(f, "Operator")
	if len(f) > 0 {
		return errors.New("ModifyTeamMemberRequest has unknown keys!")
	}
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTeamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "TeamId")
	delete(f, "Name")
	delete(f, "Operator")
	if len(f) > 0 {
		return errors.New("ModifyTeamRequest has unknown keys!")
	}
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MoveClassRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "Owner")
	delete(f, "SourceClassPath")
	delete(f, "DestinationClassPath")
	delete(f, "Operator")
	if len(f) > 0 {
		return errors.New("MoveClassRequest has unknown keys!")
	}
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MoveClassResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MoveResourceRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 待移动的原始资源信息，包含原始媒体或分类资源，以及资源归属。
	SourceResource *ResourceInfo `json:"SourceResource,omitempty" name:"SourceResource"`

	// 目标信息，包含分类及归属，仅支持移动资源到分类。
	DestinationResource *ResourceInfo `json:"DestinationResource,omitempty" name:"DestinationResource"`

	// 操作者。填写用户的 Id，用于标识调用者及校验资源访问以及写权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *MoveResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MoveResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "SourceResource")
	delete(f, "DestinationResource")
	delete(f, "Operator")
	if len(f) > 0 {
		return errors.New("MoveResourceRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type MoveResourceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MoveResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MoveResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OtherMaterial struct {

	// 素材媒体文件的播放 URL 地址。
	MaterialUrl *string `json:"MaterialUrl,omitempty" name:"MaterialUrl"`

	// 云点播媒资 FileId。
	VodFileId *string `json:"VodFileId,omitempty" name:"VodFileId"`
}

type ParseEventRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 回调事件内容。
	EventContent *string `json:"EventContent,omitempty" name:"EventContent"`
}

func (r *ParseEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ParseEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "EventContent")
	if len(f) > 0 {
		return errors.New("ParseEventRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ParseEventResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 事件内容。
		EventContent *EventContent `json:"EventContent,omitempty" name:"EventContent"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ParseEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ParseEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PenguinMediaPlatformPublishInfo struct {

	// 视频发布标题。
	Title *string `json:"Title,omitempty" name:"Title"`

	// 视频发布描述信息。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 视频标签。
	Tags []*string `json:"Tags,omitempty" name:"Tags" list`

	// 视频分类，详见[企鹅号官网](https://open.om.qq.com/resources/resourcesCenter)视频分类。
	Category *int64 `json:"Category,omitempty" name:"Category"`
}

type PlatformInfo struct {

	// 平台名称。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 平台描述。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 云点播子应用 Id。
	VodSubAppId *uint64 `json:"VodSubAppId,omitempty" name:"VodSubAppId"`

	// 平台绑定的 license Id。
	LicenseId *string `json:"LicenseId,omitempty" name:"LicenseId"`

	// 创建时间，格式按照 ISO 8601 标准表示。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间，格式按照 ISO 8601 标准表示。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type PresetTagInfo struct {

	// 标签 Id 。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 标签名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 父级预设 Id。
	ParentTagId *string `json:"ParentTagId,omitempty" name:"ParentTagId"`
}

type ProjectInfo struct {

	// 项目 Id。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 项目名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 画布宽高比。
	AspectRatio *string `json:"AspectRatio,omitempty" name:"AspectRatio"`

	// 项目类别，取值有：
	// <li>VIDEO_EDIT：视频编辑。</li>
	// <li>SWITCHER：导播台。</li>
	// <li>VIDEO_SEGMENTATION：视频拆条。</li>
	// <li>STREAM_CONNECT：云转推。</li>
	// <li>RECORD_REPLAY：录制回放。</li>
	Category *string `json:"Category,omitempty" name:"Category"`

	// 归属者。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 项目封面图片地址。
	CoverUrl *string `json:"CoverUrl,omitempty" name:"CoverUrl"`

	// 云转推项目信息，仅当项目类别取值 STREAM_CONNECT 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StreamConnectProjectInfo *StreamConnectProjectInfo `json:"StreamConnectProjectInfo,omitempty" name:"StreamConnectProjectInfo"`

	// 项目创建时间，格式按照 ISO 8601 标准表示。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 项目更新时间，格式按照 ISO 8601 标准表示。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type RecordReplayProjectInput struct {

	// 录制拉流地址。
	PullStreamUrl *string `json:"PullStreamUrl,omitempty" name:"PullStreamUrl"`

	// 录制文件归属者。
	MaterialOwner *Entity `json:"MaterialOwner,omitempty" name:"MaterialOwner"`

	// 录制文件存储分类路径。
	MaterialClassPath *string `json:"MaterialClassPath,omitempty" name:"MaterialClassPath"`

	// 回放推流地址。
	PushStreamUrl *string `json:"PushStreamUrl,omitempty" name:"PushStreamUrl"`
}

type Resource struct {

	// 类型，取值有：
	// <li>MATERIAL：素材。</li>
	// <li>CLASS：分类。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 资源 Id，当 Type 为 MATERIAL 时，取值为素材 Id；当 Type 为 CLASS 时，取值为分类路径 ClassPath。
	Id *string `json:"Id,omitempty" name:"Id"`
}

type ResourceInfo struct {

	// 媒资和分类资源。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resource *Resource `json:"Resource,omitempty" name:"Resource"`

	// 资源归属，个人或团队。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RevokeResourceAuthorizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "Owner")
	delete(f, "Resources")
	delete(f, "Authorizees")
	delete(f, "Permissions")
	delete(f, "Operator")
	if len(f) > 0 {
		return errors.New("RevokeResourceAuthorizationRequest has unknown keys!")
	}
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RevokeResourceAuthorizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RtmpPushInputInfo struct {

	// 直播推流地址有效期，单位：秒 。
	ExpiredSecond *uint64 `json:"ExpiredSecond,omitempty" name:"ExpiredSecond"`

	// 直播推流地址，入参不填默认由云剪生成。
	PushUrl *string `json:"PushUrl,omitempty" name:"PushUrl"`
}

type SearchMaterialRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 指定搜索空间，数组长度不得超过5。
	SearchScopes []*SearchScope `json:"SearchScopes,omitempty" name:"SearchScopes" list`

	// 媒体类型，取值：
	// <li>AUDIO：音频；</li>
	// <li>VIDEO：视频 ；</li>
	// <li>IMAGE：图片。</li>
	MaterialTypes []*string `json:"MaterialTypes,omitempty" name:"MaterialTypes" list`

	// 搜索文本，模糊匹配媒体名称或描述信息，匹配项越多，匹配度越高，排序越优先。长度限制：15个字符。
	Text *string `json:"Text,omitempty" name:"Text"`

	// 按画质检索，取值为：LD/SD/HD/FHD/2K/4K。
	Resolution *string `json:"Resolution,omitempty" name:"Resolution"`

	// 按媒体时长检索，单位s。
	DurationRange *IntegerRange `json:"DurationRange,omitempty" name:"DurationRange"`

	// 按照媒体创建时间检索。
	CreateTimeRange *TimeRange `json:"CreateTimeRange,omitempty" name:"CreateTimeRange"`

	// 按标签检索，填入检索的标签名。
	Tags []*string `json:"Tags,omitempty" name:"Tags" list`

	// 排序方式。Sort.Field 可选值：CreateTime。指定 Text 搜索时，将根据匹配度排序，该字段无效。
	Sort *SortBy `json:"Sort,omitempty" name:"Sort"`

	// 偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回记录条数，默认值：50。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 操作者。填写用户的 Id，用于标识调用者及校验媒体访问权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *SearchMaterialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchMaterialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "SearchScopes")
	delete(f, "MaterialTypes")
	delete(f, "Text")
	delete(f, "Resolution")
	delete(f, "DurationRange")
	delete(f, "CreateTimeRange")
	delete(f, "Tags")
	delete(f, "Sort")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Operator")
	if len(f) > 0 {
		return errors.New("SearchMaterialRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SearchMaterialResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合记录总条数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 媒体信息，仅返回基础信息。
		MaterialInfoSet []*MaterialInfo `json:"MaterialInfoSet,omitempty" name:"MaterialInfoSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchMaterialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchMaterialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchScope struct {

	// 分类路径归属。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 按分类路径检索。 不填则默认按根分类路径检索。
	ClassPath *string `json:"ClassPath,omitempty" name:"ClassPath"`
}

type SlotInfo struct {

	// 卡槽 Id。
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 素材类型，同素材素材，可取值有：
	// <li> AUDIO :音频;</li>
	// <li> VIDEO :视频;</li>
	// <li> IMAGE :图片。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 默认素材 Id。
	DefaultMaterialId *string `json:"DefaultMaterialId,omitempty" name:"DefaultMaterialId"`

	// 素材时长，单位秒。
	Duration *float64 `json:"Duration,omitempty" name:"Duration"`
}

type SlotReplacementInfo struct {

	// 卡槽 Id。
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 替换类型，可取值有：
	// <li> AUDIO :音频;</li>
	// <li> VIDEO :视频;</li>
	// <li> IMAGE :图片。</li>
	// 注意：这里必须保证替换的素材类型与模板轨道数据的素材类型一致。
	ReplacementType *string `json:"ReplacementType,omitempty" name:"ReplacementType"`

	// 媒体替换信息，仅当要替换的媒体类型为音频、视频、图片时有效。
	MediaReplacementInfo *MediaReplacementInfo `json:"MediaReplacementInfo,omitempty" name:"MediaReplacementInfo"`
}

type SortBy struct {

	// 排序字段。
	Field *string `json:"Field,omitempty" name:"Field"`

	// 排序方式，可选值：Asc（升序）、Desc（降序），默认降序。
	Order *string `json:"Order,omitempty" name:"Order"`
}

type StorageNewFileCreatedEvent struct {

	// 云点播文件  Id。
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 媒体 Id。
	MaterialId *string `json:"MaterialId,omitempty" name:"MaterialId"`

	// 操作者 Id。
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 操作类型，可取值为：
	// <li>Upload：上传；</li>
	// <li>PullUpload：拉取上传；</li>
	// <li>Record：直播录制。</li>
	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`

	// 媒体归属。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 媒体分类路径。
	ClassPath *string `json:"ClassPath,omitempty" name:"ClassPath"`
}

type StreamConnectOutput struct {

	// 云转推输出源标识，转推项目级别唯一。若不填则由后端生成。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 云转推输出源名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 云转推输出源类型，取值：
	// <li>URL ：URL类型</li>
	// 不填默认为URL类型。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 云转推推流地址。
	PushUrl *string `json:"PushUrl,omitempty" name:"PushUrl"`
}

type StreamConnectOutputInfo struct {

	// 输出源。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StreamConnectOutput *StreamConnectOutput `json:"StreamConnectOutput,omitempty" name:"StreamConnectOutput"`

	// 输出流状态：
	// <li>On ：开；</li>
	// <li>Off ：关 。</li>
	PushSwitch *string `json:"PushSwitch,omitempty" name:"PushSwitch"`
}

type StreamConnectProjectInfo struct {

	// 转推项目状态，取值有：
	// <li>Working ：转推中；</li>
	// <li>Idle ：空闲中。</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 当前转推输入源，取值有：
	// <li>Main ：主输入源；</li>
	// <li>Backup ：备输入源。</li>
	CurrentInputEndpoint *string `json:"CurrentInputEndpoint,omitempty" name:"CurrentInputEndpoint"`

	// 当前转推开始时间， 采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。仅 Status 取值 Working 时有效。
	CurrentStartTime *string `json:"CurrentStartTime,omitempty" name:"CurrentStartTime"`

	// 当前转推计划结束时间， 采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。仅 Status 取值 Working 时有效。
	CurrentStopTime *string `json:"CurrentStopTime,omitempty" name:"CurrentStopTime"`

	// 上一次转推结束时间， 采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。仅 Status 取值 Idle 时有效。
	LastStopTime *string `json:"LastStopTime,omitempty" name:"LastStopTime"`

	// 云转推主输入源。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MainInput *StreamInputInfo `json:"MainInput,omitempty" name:"MainInput"`

	// 云转推备输入源。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupInput *StreamInputInfo `json:"BackupInput,omitempty" name:"BackupInput"`

	// 云转推输出源。
	OutputSet []*StreamConnectOutputInfo `json:"OutputSet,omitempty" name:"OutputSet" list`
}

type StreamConnectProjectInput struct {

	// 云转推主输入源信息。
	MainInput *StreamInputInfo `json:"MainInput,omitempty" name:"MainInput"`

	// 云转推备输入源信息。
	BackupInput *StreamInputInfo `json:"BackupInput,omitempty" name:"BackupInput"`

	// 云转推输出源信息。
	Outputs []*StreamConnectOutput `json:"Outputs,omitempty" name:"Outputs" list`
}

type StreamInputInfo struct {

	// 流输入类型，取值：
	// <li>VodPull ： 点播拉流；</li>
	// <li>LivePull ：直播拉流；</li>
	// <li>RtmpPush ： 直播推流。</li>
	InputType *string `json:"InputType,omitempty" name:"InputType"`

	// 点播拉流信息，当 InputType = VodPull 时必填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VodPullInputInfo *VodPullInputInfo `json:"VodPullInputInfo,omitempty" name:"VodPullInputInfo"`

	// 直播拉流信息，当 InputType = LivePull  时必填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LivePullInputInfo *LivePullInputInfo `json:"LivePullInputInfo,omitempty" name:"LivePullInputInfo"`

	// 直播推流信息，当 InputType = RtmpPush 时必填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RtmpPushInputInfo *RtmpPushInputInfo `json:"RtmpPushInputInfo,omitempty" name:"RtmpPushInputInfo"`
}

type SwitcherPgmOutputConfig struct {

	// 导播台输出模板 ID，可取值：
	// <li>10001：分辨率为1080 P；</li>
	// <li>10002：分辨率为720 P；</li>
	// <li>10003：分辨率为480 P。</li>
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 导播台输出宽，单位：像素。
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// 导播台输出高，单位：像素。
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// 导播台输出帧率，单位：帧/秒
	Fps *uint64 `json:"Fps,omitempty" name:"Fps"`

	// 导播台输出码率， 单位：bit/s。
	BitRate *uint64 `json:"BitRate,omitempty" name:"BitRate"`
}

type SwitcherProjectInput struct {

	// 导播台停止时间，格式按照 ISO 8601 标准表示。若不填，该值默认为当前时间加七天。
	StopTime *string `json:"StopTime,omitempty" name:"StopTime"`

	// 导播台主监输出配置信息。若不填，默认输出 720P。
	PgmOutputConfig *SwitcherPgmOutputConfig `json:"PgmOutputConfig,omitempty" name:"PgmOutputConfig"`
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

type ThirdPartyPublishInfo struct {

	// 发布通道  ID。
	ChannelMaterialId *string `json:"ChannelMaterialId,omitempty" name:"ChannelMaterialId"`

	// 企鹅号发布信息，如果使用的发布通道为企鹅号时必填。
	PenguinMediaPlatformPublishInfo *PenguinMediaPlatformPublishInfo `json:"PenguinMediaPlatformPublishInfo,omitempty" name:"PenguinMediaPlatformPublishInfo"`

	// 新浪微博发布信息，如果使用的发布通道为新浪微博时必填。
	WeiboPublishInfo *WeiboPublishInfo `json:"WeiboPublishInfo,omitempty" name:"WeiboPublishInfo"`

	// 快手发布信息，如果使用的发布通道为快手时必填。
	KuaishouPublishInfo *KuaishouPublishInfo `json:"KuaishouPublishInfo,omitempty" name:"KuaishouPublishInfo"`
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

	// 第三方平台发布信息列表。暂未正式对外，请勿使用。
	ThirdPartyPublishInfos []*ThirdPartyPublishInfo `json:"ThirdPartyPublishInfos,omitempty" name:"ThirdPartyPublishInfos" list`
}

type VideoEditProjectInput struct {

	// 画布宽高比，取值有：
	// <li>16:9；</li>
	// <li>9:16；</li>
	// <li>2:1。</li>
	// 默认值 16:9 。
	AspectRatio *string `json:"AspectRatio,omitempty" name:"AspectRatio"`

	// 视频编辑模板媒体 ID ，通过模板媒体导入项目轨道数据时填写。
	VideoEditTemplateId *string `json:"VideoEditTemplateId,omitempty" name:"VideoEditTemplateId"`

	// 输入的媒体轨道列表，包括视频、音频，等媒体组成的多个轨道信息。其中：<li>输入的多个轨道在时间轴上和输出媒体文件的时间轴对齐；</li><li>时间轴上相同时间点的各个轨道的素材进行重叠，视频或者图片按轨道顺序进行图像的叠加，轨道顺序高的素材叠加在上面，音频素材进行混音；</li><li>视频、音频，每一种类型的轨道最多支持10个。</li>
	// 注：当从模板导入项目时（即 VideoEditTemplateId 不为空时），该参数无效。
	InitTracks []*MediaTrack `json:"InitTracks,omitempty" name:"InitTracks" list`
}

type VideoEditProjectOutput struct {

	// 导出的云剪素材 MaterialId，仅当导出为云剪素材时有效。
	MaterialId *string `json:"MaterialId,omitempty" name:"MaterialId"`

	// 云点播媒资 FileId。
	VodFileId *string `json:"VodFileId,omitempty" name:"VodFileId"`

	// 导出的媒资 URL。
	URL *string `json:"URL,omitempty" name:"URL"`

	// 元信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetaData *MediaMetaData `json:"MetaData,omitempty" name:"MetaData"`
}

type VideoEditTemplateMaterial struct {

	// 视频编辑模板宽高比。
	AspectRatio *string `json:"AspectRatio,omitempty" name:"AspectRatio"`

	// 卡槽信息。
	SlotSet []*SlotInfo `json:"SlotSet,omitempty" name:"SlotSet" list`
}

type VideoMaterial struct {

	// 素材元信息。
	MetaData *MediaMetaData `json:"MetaData,omitempty" name:"MetaData"`

	// 雪碧图信息。
	ImageSpriteInfo *MediaImageSpriteInfo `json:"ImageSpriteInfo,omitempty" name:"ImageSpriteInfo"`

	// 素材媒体文件的播放 URL 地址。
	MaterialUrl *string `json:"MaterialUrl,omitempty" name:"MaterialUrl"`

	// 素材媒体文件的封面图片地址。
	CoverUrl *string `json:"CoverUrl,omitempty" name:"CoverUrl"`

	// 媒体文件分辨率。取值为：LD/SD/HD/FHD/2K/4K。
	Resolution *string `json:"Resolution,omitempty" name:"Resolution"`

	// 素材状态。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaterialStatus *MaterialStatus `json:"MaterialStatus,omitempty" name:"MaterialStatus"`

	// 素材媒体文件的原始 URL 地址。
	OriginalUrl *string `json:"OriginalUrl,omitempty" name:"OriginalUrl"`

	// 云点播媒资 FileId。
	VodFileId *string `json:"VodFileId,omitempty" name:"VodFileId"`
}

type VideoSegmentationProjectInput struct {

	// 画布宽高比，取值有：
	// <li>16:9；</li>
	// <li>9:16；</li>
	// <li>2:1。</li>
	// 默认值 16:9 。
	AspectRatio *string `json:"AspectRatio,omitempty" name:"AspectRatio"`

	// 视频拆条处理模型，不填则默认为手工分割视频。取值 ：
	// <li>AI.GameHighlights.PUBG：和平精英集锦 ;</li>
	// <li>AI.GameHighlights.Honor OfKings：王者荣耀集锦 ;</li>
	// <li>AI.SportHighlights.Football：足球集锦 </li>
	// <li>AI.SportHighlights.Basketball：篮球集锦 ；</li>
	// <li>AI.PersonSegmentation：人物集锦  ;</li>
	// <li>AI.NewsSegmentation：新闻拆条。</li>
	ProcessModel *string `json:"ProcessModel,omitempty" name:"ProcessModel"`
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

type VideoTrackItem struct {

	// 视频媒体来源类型，取值有：
	// <ul>
	// <li>VOD ：媒体来源于云点播文件 。</li>
	// <li>CME ：视频来源制作云媒体文件。</li>
	// <li>EXTERNAL ：视频来源于媒资绑定。</li>
	// </ul>
	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`

	// 视频片段的媒体文件来源，取值为：
	// <ul>
	// <li>当 SourceType 为 VOD 时，为云点播的媒体文件 FileId ，会默认将该 FileId 导入到项目中；</li>
	// <li>当 SourceType 为 CME 时，为制作云的媒体 ID，项目归属者必须对该云媒资有访问权限；</li>
	// <li>当 SourceType 为 EXTERNAL 时，为媒资绑定的 Definition 与 MediaKey 中间用冒号分隔合并后的字符串，格式为 Definition:MediaKey 。</li>
	// </ul>
	SourceMedia *string `json:"SourceMedia,omitempty" name:"SourceMedia"`

	// 视频片段取自媒体文件的起始时间，单位为秒。默认为0。
	SourceMediaStartTime *float64 `json:"SourceMediaStartTime,omitempty" name:"SourceMediaStartTime"`

	// 视频片段时长，单位为秒。默认取视频媒体文件本身长度，表示截取全部媒体文件。如果源文件是图片，Duration需要大于0。
	Duration *float64 `json:"Duration,omitempty" name:"Duration"`

	// 视频片段原点距离画布原点的水平位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示视频片段 XPos 为画布宽度指定百分比的位置，如 10% 表示 XPos 为画布口宽度的 10%。</li>
	// <li>当字符串以 px 结尾，表示视频片段 XPos 单位为像素，如 100px 表示 XPos 为100像素。</li>
	// 默认值：0px。
	XPos *string `json:"XPos,omitempty" name:"XPos"`

	// 视频片段原点距离画布原点的垂直位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示视频片段 YPos 为画布高度指定百分比的位置，如 10% 表示 YPos 为画布高度的 10%。</li>
	// <li>当字符串以 px 结尾，表示视频片段 YPos 单位为像素，如 100px 表示 YPos 为100像素。</li>
	// 默认值：0px。
	YPos *string `json:"YPos,omitempty" name:"YPos"`

	// 视频原点位置，取值有：
	// <li>Center：坐标原点为中心位置，如画布中心。</li>
	// 默认值 ：Center。
	CoordinateOrigin *string `json:"CoordinateOrigin,omitempty" name:"CoordinateOrigin"`

	// 视频片段的高度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示视频片段 Height 为画布高度的百分比大小，如 10% 表示 Height 为画布高度的 10%；</li>
	// <li>当字符串以 px 结尾，表示视频片段 Height 单位为像素，如 100px 表示 Height 为100像素；</li>
	// <li>当 Width、Height 均为空，则 Width 和 Height 取视频媒体文件本身的 Width、Height；</li>
	// <li>当 Width 为空，Height 非空，则 Width 按比例缩放；</li>
	// <li>当 Width 非空，Height 为空，则 Height 按比例缩放。</li>
	Height *string `json:"Height,omitempty" name:"Height"`

	// 视频片段的宽度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示视频片段 Width 为画布宽度的百分比大小，如 10% 表示 Width 为画布宽度的 10%；</li>
	// <li>当字符串以 px 结尾，表示视频片段 Width 单位为像素，如 100px 表示 Width 为100像素；</li>
	// <li>当 Width、Height 均为空，则 Width 和 Height 取视频媒体文件本身的 Width、Height；</li>
	// <li>当 Width 为空，Height 非空，则 Width 按比例缩放；</li>
	// <li>当 Width 非空，Height 为空，则 Height 按比例缩放。</li>
	Width *string `json:"Width,omitempty" name:"Width"`
}

type VodPullInputInfo struct {

	// 点播输入拉流 URL 。
	InputUrls []*string `json:"InputUrls,omitempty" name:"InputUrls" list`

	// 播放次数，取值有：
	// <li>-1 : 循环播放，直到转推结束；</li>
	// <li>0 : 不循环；</li>
	// <li>大于0 : 具体循环次数，次数和时间以先结束的为准。</li>
	// 默认不循环。
	LoopTimes *int64 `json:"LoopTimes,omitempty" name:"LoopTimes"`
}

type WeiboPublishInfo struct {

	// 视频发布标题。
	Title *string `json:"Title,omitempty" name:"Title"`

	// 视频发布描述信息。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 微博可见性，可取值为：
	// <li>Public：公开，所有人可见；</li>
	// <li>Private：私有，仅自己可见。</li>
	// 
	// 默认为 Public，所有人可见。
	Visible *string `json:"Visible,omitempty" name:"Visible"`
}
