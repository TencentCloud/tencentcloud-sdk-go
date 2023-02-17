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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
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

// Predefined struct for user
type AddTeamMemberRequestParams struct {
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 团队 ID。
	TeamId *string `json:"TeamId,omitempty" name:"TeamId"`

	// 要添加的成员列表，一次最多添加30个成员。
	TeamMembers []*AddMemberInfo `json:"TeamMembers,omitempty" name:"TeamMembers"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以向任意团队中添加成员。如果指定操作者，则操作者必须为管理员或者团队所有者。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type AddTeamMemberRequest struct {
	*tchttp.BaseRequest
	
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 团队 ID。
	TeamId *string `json:"TeamId,omitempty" name:"TeamId"`

	// 要添加的成员列表，一次最多添加30个成员。
	TeamMembers []*AddMemberInfo `json:"TeamMembers,omitempty" name:"TeamMembers"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以向任意团队中添加成员。如果指定操作者，则操作者必须为管理员或者团队所有者。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *AddTeamMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddTeamMemberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddTeamMemberResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddTeamMemberResponse struct {
	*tchttp.BaseResponse
	Response *AddTeamMemberResponseParams `json:"Response"`
}

func (r *AddTeamMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
	// <li>EXTERNAL ：视频来源于媒资绑定，如果媒体不是存储在腾讯云点播中或者云创中，都需要使用媒资绑定。</li>
	// </ul>
	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`

	// 音频媒体，可取值为：
	// <ul>
	// <li>当 SourceType 为 VOD 时，参数填云点播 FileId ；</li>
	// <li>当 SourceType 为 CME 时，参数填多媒体创作引擎媒体 Id；</li>
	// <li>当 SourceType 为 EXTERNAL 时，目前仅支持外部媒体 URL(如`https://www.example.com/a.mp3`)，参数填写规则请参见注意事项。</li>
	// </ul>
	// 
	// 注意：
	// <li>当 SourceType 为 EXTERNAL 并且媒体 URL Scheme 为 `https` 时(如：`https://www.example.com/a.mp3`)，参数为：`1000000:www.example.com/a.mp3`。</li>
	// <li>当 SourceType 为 EXTERNAL 并且媒体 URL Scheme 为 `http` 时(如：`http://www.example.com/b.mp3`)，参数为：`1000001:www.example.com/b.mp3`。</li>
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
	PermissionSet []*string `json:"PermissionSet,omitempty" name:"PermissionSet"`
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

	// 导出的媒体分类路径，长度不能超过15字符。不存在默认创建。
	ClassPath *string `json:"ClassPath,omitempty" name:"ClassPath"`

	// 导出的媒体标签，单个标签不得超过10个字符。
	TagSet []*string `json:"TagSet,omitempty" name:"TagSet"`

	// 第三方平台发布信息列表。暂未正式对外，请勿使用。
	ThirdPartyPublishInfos []*ThirdPartyPublishInfo `json:"ThirdPartyPublishInfos,omitempty" name:"ThirdPartyPublishInfos"`
}

type ClassCreatedEvent struct {
	// 分类归属。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 分类路径。
	ClassPath *string `json:"ClassPath,omitempty" name:"ClassPath"`
}

type ClassDeletedEvent struct {
	// 删除的分类归属。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 删除的分类路径列表。
	ClassPathSet []*string `json:"ClassPathSet,omitempty" name:"ClassPathSet"`
}

type ClassInfo struct {
	// 归属者。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 分类路径。
	ClassPath *string `json:"ClassPath,omitempty" name:"ClassPath"`
}

type ClassMovedEvent struct {
	// 源分类归属。
	SourceOwner *Entity `json:"SourceOwner,omitempty" name:"SourceOwner"`

	// 源分类路径列表。
	SourceClassPathSet []*string `json:"SourceClassPathSet,omitempty" name:"SourceClassPathSet"`

	// 目标分类归属。
	DestinationOwner *Entity `json:"DestinationOwner,omitempty" name:"DestinationOwner"`

	// 目标分类归属。
	DestinationClassPath *string `json:"DestinationClassPath,omitempty" name:"DestinationClassPath"`
}

// Predefined struct for user
type CopyProjectRequestParams struct {
	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 被复制的项目 ID。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 复制后的项目名称，不填为原项目名称+"(副本)"。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 复制后的项目归属者，不填为原项目归属者。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 操作者。填写用户的 Id，用于标识调用者及校验操作权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type CopyProjectRequest struct {
	*tchttp.BaseRequest
	
	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 被复制的项目 ID。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 复制后的项目名称，不填为原项目名称+"(副本)"。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 复制后的项目归属者，不填为原项目归属者。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 操作者。填写用户的 Id，用于标识调用者及校验操作权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *CopyProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "ProjectId")
	delete(f, "Name")
	delete(f, "Owner")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CopyProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CopyProjectResponseParams struct {
	// 复制后的项目 ID。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CopyProjectResponse struct {
	*tchttp.BaseResponse
	Response *CopyProjectResponseParams `json:"Response"`
}

func (r *CopyProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CosPublishInputInfo struct {
	// 发布生成的对象存储文件所在的 COS Bucket 名，如 TopRankVideo-125xxx88。
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// 发布生成的对象存储文件所在的 COS Bucket 所属园区，如 ap-chongqing。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 发布生成的视频在 COS 存储的对象键。对象键（ObjectKey）是对象（Object）在存储桶（Bucket）中的唯一标识。
	VideoKey *string `json:"VideoKey,omitempty" name:"VideoKey"`

	// 发布生成的封面在 COS 存储的对象键。
	CoverKey *string `json:"CoverKey,omitempty" name:"CoverKey"`
}

// Predefined struct for user
type CreateClassRequestParams struct {
	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 归属者。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 分类路径。
	ClassPath *string `json:"ClassPath,omitempty" name:"ClassPath"`

	// 操作者。填写用户的 Id，用于标识调用者及校验分类创建权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
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

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClassRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClassResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateClassResponse struct {
	*tchttp.BaseResponse
	Response *CreateClassResponseParams `json:"Response"`
}

func (r *CreateClassResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClassResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLinkRequestParams struct {
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 链接类型，可取值有:
	// <li>CLASS: 分类链接；</li>
	// <li> MATERIAL：媒体文件链接。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 链接名称，不能超过30个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 链接归属者。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 目标资源Id。可取值有：
	// <li>当 Type 为 MATERIAL 时填媒体 ID；</li>
	// <li>当 Type 为 CLASS 时填写分类路径。</li>
	DestinationId *string `json:"DestinationId,omitempty" name:"DestinationId"`

	// 目标资源归属者。
	DestinationOwner *Entity `json:"DestinationOwner,omitempty" name:"DestinationOwner"`

	// 链接的分类路径，如填"/a/b"则代表链接属于该分类路径，不填则默认为根路径。
	ClassPath *string `json:"ClassPath,omitempty" name:"ClassPath"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以创建任意源及目标资源的链接。如果指定操作者，则操作者必须对源资源有读权限，对目标媒体有写权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type CreateLinkRequest struct {
	*tchttp.BaseRequest
	
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 链接类型，可取值有:
	// <li>CLASS: 分类链接；</li>
	// <li> MATERIAL：媒体文件链接。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 链接名称，不能超过30个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 链接归属者。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 目标资源Id。可取值有：
	// <li>当 Type 为 MATERIAL 时填媒体 ID；</li>
	// <li>当 Type 为 CLASS 时填写分类路径。</li>
	DestinationId *string `json:"DestinationId,omitempty" name:"DestinationId"`

	// 目标资源归属者。
	DestinationOwner *Entity `json:"DestinationOwner,omitempty" name:"DestinationOwner"`

	// 链接的分类路径，如填"/a/b"则代表链接属于该分类路径，不填则默认为根路径。
	ClassPath *string `json:"ClassPath,omitempty" name:"ClassPath"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以创建任意源及目标资源的链接。如果指定操作者，则操作者必须对源资源有读权限，对目标媒体有写权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *CreateLinkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLinkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLinkResponseParams struct {
	// 新建链接的媒体 Id。
	MaterialId *string `json:"MaterialId,omitempty" name:"MaterialId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateLinkResponse struct {
	*tchttp.BaseResponse
	Response *CreateLinkResponseParams `json:"Response"`
}

func (r *CreateLinkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLinkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProjectRequestParams struct {
	// 平台 Id，指定访问的平台。平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 项目名称，不可超过30个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 项目归属者，即项目的所有者，后续操作只有该所有者有权限操作。
	// 
	// 注：目前所有项目只能设置归属个人，暂不支持团队项目。
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
	// <li>Default：默认模式，即普通视频编辑项目。</li>
	// <li>VideoEditTemplate：剪辑模板制作模式，用于制作剪辑模板。</li>
	// 
	// 注：不填则为默认模式。
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// 画布宽高比。
	// 该字段已经废弃，请使用具体项目输入中的 AspectRatio 字段。
	AspectRatio *string `json:"AspectRatio,omitempty" name:"AspectRatio"`

	// 项目描述信息。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 导播台项目输入信息，仅当项目类型为 SWITCHER 时必填。
	SwitcherProjectInput *SwitcherProjectInput `json:"SwitcherProjectInput,omitempty" name:"SwitcherProjectInput"`

	// 直播剪辑项目输入信息，暂未开放，请勿使用。
	LiveStreamClipProjectInput *LiveStreamClipProjectInput `json:"LiveStreamClipProjectInput,omitempty" name:"LiveStreamClipProjectInput"`

	// 视频编辑项目输入信息，仅当项目类型为 VIDEO_EDIT 时必填。
	VideoEditProjectInput *VideoEditProjectInput `json:"VideoEditProjectInput,omitempty" name:"VideoEditProjectInput"`

	// 视频拆条项目输入信息，仅当项目类型为 VIDEO_SEGMENTATION  时必填。
	VideoSegmentationProjectInput *VideoSegmentationProjectInput `json:"VideoSegmentationProjectInput,omitempty" name:"VideoSegmentationProjectInput"`

	// 云转推项目输入信息，仅当项目类型为 STREAM_CONNECT 时必填。
	StreamConnectProjectInput *StreamConnectProjectInput `json:"StreamConnectProjectInput,omitempty" name:"StreamConnectProjectInput"`

	// 录制回放项目输入信息，仅当项目类型为 RECORD_REPLAY 时必填。
	RecordReplayProjectInput *RecordReplayProjectInput `json:"RecordReplayProjectInput,omitempty" name:"RecordReplayProjectInput"`
}

type CreateProjectRequest struct {
	*tchttp.BaseRequest
	
	// 平台 Id，指定访问的平台。平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 项目名称，不可超过30个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 项目归属者，即项目的所有者，后续操作只有该所有者有权限操作。
	// 
	// 注：目前所有项目只能设置归属个人，暂不支持团队项目。
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
	// <li>Default：默认模式，即普通视频编辑项目。</li>
	// <li>VideoEditTemplate：剪辑模板制作模式，用于制作剪辑模板。</li>
	// 
	// 注：不填则为默认模式。
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// 画布宽高比。
	// 该字段已经废弃，请使用具体项目输入中的 AspectRatio 字段。
	AspectRatio *string `json:"AspectRatio,omitempty" name:"AspectRatio"`

	// 项目描述信息。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 导播台项目输入信息，仅当项目类型为 SWITCHER 时必填。
	SwitcherProjectInput *SwitcherProjectInput `json:"SwitcherProjectInput,omitempty" name:"SwitcherProjectInput"`

	// 直播剪辑项目输入信息，暂未开放，请勿使用。
	LiveStreamClipProjectInput *LiveStreamClipProjectInput `json:"LiveStreamClipProjectInput,omitempty" name:"LiveStreamClipProjectInput"`

	// 视频编辑项目输入信息，仅当项目类型为 VIDEO_EDIT 时必填。
	VideoEditProjectInput *VideoEditProjectInput `json:"VideoEditProjectInput,omitempty" name:"VideoEditProjectInput"`

	// 视频拆条项目输入信息，仅当项目类型为 VIDEO_SEGMENTATION  时必填。
	VideoSegmentationProjectInput *VideoSegmentationProjectInput `json:"VideoSegmentationProjectInput,omitempty" name:"VideoSegmentationProjectInput"`

	// 云转推项目输入信息，仅当项目类型为 STREAM_CONNECT 时必填。
	StreamConnectProjectInput *StreamConnectProjectInput `json:"StreamConnectProjectInput,omitempty" name:"StreamConnectProjectInput"`

	// 录制回放项目输入信息，仅当项目类型为 RECORD_REPLAY 时必填。
	RecordReplayProjectInput *RecordReplayProjectInput `json:"RecordReplayProjectInput,omitempty" name:"RecordReplayProjectInput"`
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProjectResponseParams struct {
	// 项目 Id。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// <li> 当 Catagory 为 STREAM_CONNECT 时，数组返回长度为2 ，第0个代表主输入源推流信息，第1个代表备输入源推流信息。只有当各自输入源类型为推流时才有有效内容。</li>
	RtmpPushInputInfoSet []*RtmpPushInputInfo `json:"RtmpPushInputInfoSet,omitempty" name:"RtmpPushInputInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type CreateTeamRequestParams struct {
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

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTeamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTeamResponseParams struct {
	// 创建的团队 ID。
	TeamId *string `json:"TeamId,omitempty" name:"TeamId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateTeamResponse struct {
	*tchttp.BaseResponse
	Response *CreateTeamResponseParams `json:"Response"`
}

func (r *CreateTeamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTeamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVideoEncodingPresetRequestParams struct {
	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 配置名，可用来简单描述该配置的作用。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 封装格式，可选值：
	// <li>mp4 ；</li>
	// <li>mov 。</li>
	// 默认值：mp4。
	Container *string `json:"Container,omitempty" name:"Container"`

	// 是否去除视频数据，可选值：
	// <li>0：保留；</li>
	// <li>1：去除。</li>
	// 默认值：0。
	RemoveVideo *int64 `json:"RemoveVideo,omitempty" name:"RemoveVideo"`

	// 是否去除音频数据，可选值：
	// <li>0：保留；</li>
	// <li>1：去除。</li>
	// 默认值：0。
	RemoveAudio *int64 `json:"RemoveAudio,omitempty" name:"RemoveAudio"`

	// 编码配置的视频设置。默认值参考VideoEncodingPresetVideoSetting 定义。
	VideoSetting *VideoEncodingPresetVideoSetting `json:"VideoSetting,omitempty" name:"VideoSetting"`

	// 编码配置的音频设置。默认值参考VideoEncodingPresetAudioSetting 定义。
	AudioSetting *VideoEncodingPresetAudioSetting `json:"AudioSetting,omitempty" name:"AudioSetting"`
}

type CreateVideoEncodingPresetRequest struct {
	*tchttp.BaseRequest
	
	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 配置名，可用来简单描述该配置的作用。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 封装格式，可选值：
	// <li>mp4 ；</li>
	// <li>mov 。</li>
	// 默认值：mp4。
	Container *string `json:"Container,omitempty" name:"Container"`

	// 是否去除视频数据，可选值：
	// <li>0：保留；</li>
	// <li>1：去除。</li>
	// 默认值：0。
	RemoveVideo *int64 `json:"RemoveVideo,omitempty" name:"RemoveVideo"`

	// 是否去除音频数据，可选值：
	// <li>0：保留；</li>
	// <li>1：去除。</li>
	// 默认值：0。
	RemoveAudio *int64 `json:"RemoveAudio,omitempty" name:"RemoveAudio"`

	// 编码配置的视频设置。默认值参考VideoEncodingPresetVideoSetting 定义。
	VideoSetting *VideoEncodingPresetVideoSetting `json:"VideoSetting,omitempty" name:"VideoSetting"`

	// 编码配置的音频设置。默认值参考VideoEncodingPresetAudioSetting 定义。
	AudioSetting *VideoEncodingPresetAudioSetting `json:"AudioSetting,omitempty" name:"AudioSetting"`
}

func (r *CreateVideoEncodingPresetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVideoEncodingPresetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "Name")
	delete(f, "Container")
	delete(f, "RemoveVideo")
	delete(f, "RemoveAudio")
	delete(f, "VideoSetting")
	delete(f, "AudioSetting")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVideoEncodingPresetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVideoEncodingPresetResponseParams struct {
	// 模板 ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateVideoEncodingPresetResponse struct {
	*tchttp.BaseResponse
	Response *CreateVideoEncodingPresetResponseParams `json:"Response"`
}

func (r *CreateVideoEncodingPresetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVideoEncodingPresetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClassRequestParams struct {
	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 归属者。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 分类路径。
	ClassPath *string `json:"ClassPath,omitempty" name:"ClassPath"`

	// 操作者。填写用户的 Id，用于标识调用者及校验操作权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
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

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteClassRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClassResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteClassResponse struct {
	*tchttp.BaseResponse
	Response *DeleteClassResponseParams `json:"Response"`
}

func (r *DeleteClassResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClassResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLoginStatusRequestParams struct {
	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 用户 Id 列表，N 从 0 开始取值，最大 19。
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`
}

type DeleteLoginStatusRequest struct {
	*tchttp.BaseRequest
	
	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 用户 Id 列表，N 从 0 开始取值，最大 19。
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`
}

func (r *DeleteLoginStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoginStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "UserIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLoginStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLoginStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteLoginStatusResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLoginStatusResponseParams `json:"Response"`
}

func (r *DeleteLoginStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoginStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMaterialRequestParams struct {
	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 媒体 Id。
	MaterialId *string `json:"MaterialId,omitempty" name:"MaterialId"`

	// 操作者。填写用户的 Id，用于标识调用者及校验媒体删除权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
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

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMaterialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMaterialResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteMaterialResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMaterialResponseParams `json:"Response"`
}

func (r *DeleteMaterialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMaterialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProjectRequestParams struct {
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 要删除的项目 Id。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以删除一切项目。如果指定操作者，则操作者必须为项目所有者。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type DeleteProjectRequest struct {
	*tchttp.BaseRequest
	
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 要删除的项目 Id。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以删除一切项目。如果指定操作者，则操作者必须为项目所有者。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
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
	delete(f, "Platform")
	delete(f, "ProjectId")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProjectResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DeleteTeamMembersRequestParams struct {
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 团队 ID。
	TeamId *string `json:"TeamId,omitempty" name:"TeamId"`

	// 要删除的成员列表。
	MemberIds []*string `json:"MemberIds,omitempty" name:"MemberIds"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以删除所有团队的成员。如果指定操作者，则操作者必须为团队管理员或者所有者。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type DeleteTeamMembersRequest struct {
	*tchttp.BaseRequest
	
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 团队 ID。
	TeamId *string `json:"TeamId,omitempty" name:"TeamId"`

	// 要删除的成员列表。
	MemberIds []*string `json:"MemberIds,omitempty" name:"MemberIds"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以删除所有团队的成员。如果指定操作者，则操作者必须为团队管理员或者所有者。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *DeleteTeamMembersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTeamMembersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTeamMembersResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteTeamMembersResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTeamMembersResponseParams `json:"Response"`
}

func (r *DeleteTeamMembersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTeamMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTeamRequestParams struct {
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 要删除的团队  ID。
	TeamId *string `json:"TeamId,omitempty" name:"TeamId"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以删除所有团队。如果指定操作者，则操作者必须为团队所有者。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type DeleteTeamRequest struct {
	*tchttp.BaseRequest
	
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 要删除的团队  ID。
	TeamId *string `json:"TeamId,omitempty" name:"TeamId"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以删除所有团队。如果指定操作者，则操作者必须为团队所有者。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *DeleteTeamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTeamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTeamResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteTeamResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTeamResponseParams `json:"Response"`
}

func (r *DeleteTeamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTeamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVideoEncodingPresetRequestParams struct {
	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 要删除的视频编码配置 ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

type DeleteVideoEncodingPresetRequest struct {
	*tchttp.BaseRequest
	
	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 要删除的视频编码配置 ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteVideoEncodingPresetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVideoEncodingPresetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVideoEncodingPresetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVideoEncodingPresetResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteVideoEncodingPresetResponse struct {
	*tchttp.BaseResponse
	Response *DeleteVideoEncodingPresetResponseParams `json:"Response"`
}

func (r *DeleteVideoEncodingPresetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVideoEncodingPresetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountsRequestParams struct {
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 手机号码。指定手机号获取账号信息，目前仅支持国内手机号，且号码不加地区码 `+86` 等。
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 分页返回的起始偏移量，默认值：0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页返回的记录条数，默认值：10，最大值：20。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeAccountsRequest struct {
	*tchttp.BaseRequest
	
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 手机号码。指定手机号获取账号信息，目前仅支持国内手机号，且号码不加地区码 `+86` 等。
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

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountsResponseParams struct {
	// 符合搜索条件的记录总数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 账号信息列表。
	AccountInfoSet []*AccountInfo `json:"AccountInfoSet,omitempty" name:"AccountInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAccountsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccountsResponseParams `json:"Response"`
}

func (r *DescribeAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClassRequestParams struct {
	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 归属者。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 操作者。填写用户的 Id，用于标识调用者及校验操作权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
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

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClassRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClassResponseParams struct {
	// 分类信息列表。
	ClassInfoSet []*ClassInfo `json:"ClassInfoSet,omitempty" name:"ClassInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClassResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClassResponseParams `json:"Response"`
}

func (r *DescribeClassResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClassResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJoinTeamsRequestParams struct {
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 团队成员　ID。
	MemberId *string `json:"MemberId,omitempty" name:"MemberId"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回记录条数，默认值：30，最大值：30。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeJoinTeamsRequest struct {
	*tchttp.BaseRequest
	
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 团队成员　ID。
	MemberId *string `json:"MemberId,omitempty" name:"MemberId"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回记录条数，默认值：30，最大值：30。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeJoinTeamsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeJoinTeamsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJoinTeamsResponseParams struct {
	// 符合条件的记录总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 团队列表。
	TeamSet []*JoinTeamInfo `json:"TeamSet,omitempty" name:"TeamSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeJoinTeamsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeJoinTeamsResponseParams `json:"Response"`
}

func (r *DescribeJoinTeamsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJoinTeamsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoginStatusRequestParams struct {
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 用户 Id 列表，N 从0开始取值，最大19。
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`
}

type DescribeLoginStatusRequest struct {
	*tchttp.BaseRequest
	
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 用户 Id 列表，N 从0开始取值，最大19。
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`
}

func (r *DescribeLoginStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoginStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "UserIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLoginStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoginStatusResponseParams struct {
	// 用户登录状态列表。
	LoginStatusInfoSet []*LoginStatusInfo `json:"LoginStatusInfoSet,omitempty" name:"LoginStatusInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLoginStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLoginStatusResponseParams `json:"Response"`
}

func (r *DescribeLoginStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoginStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMaterialsRequestParams struct {
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 媒体 ID 列表，一次最多可拉取20个媒体的信息。
	MaterialIds []*string `json:"MaterialIds,omitempty" name:"MaterialIds"`

	// 列表排序，支持下列排序字段：
	// <li>CreateTime：创建时间；</li>
	// <li>UpdateTime：更新时间。</li>
	Sort *SortBy `json:"Sort,omitempty" name:"Sort"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以获取任意媒体的信息。如果指定操作者，则操作者必须对媒体有读权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type DescribeMaterialsRequest struct {
	*tchttp.BaseRequest
	
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 媒体 ID 列表，一次最多可拉取20个媒体的信息。
	MaterialIds []*string `json:"MaterialIds,omitempty" name:"MaterialIds"`

	// 列表排序，支持下列排序字段：
	// <li>CreateTime：创建时间；</li>
	// <li>UpdateTime：更新时间。</li>
	Sort *SortBy `json:"Sort,omitempty" name:"Sort"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以获取任意媒体的信息。如果指定操作者，则操作者必须对媒体有读权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *DescribeMaterialsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMaterialsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMaterialsResponseParams struct {
	// 媒体列表信息。
	MaterialInfoSet []*MaterialInfo `json:"MaterialInfoSet,omitempty" name:"MaterialInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeMaterialsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMaterialsResponseParams `json:"Response"`
}

func (r *DescribeMaterialsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMaterialsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePlatformsRequestParams struct {
	// 平台 Id 列表。如果不填，则不按平台 Id 进行过滤。
	Platforms []*string `json:"Platforms,omitempty" name:"Platforms"`

	// 平台绑定的 License Id 列表。如果不填，则不按平台绑定的 License Id 进行过滤。
	LicenseIds []*string `json:"LicenseIds,omitempty" name:"LicenseIds"`

	// 分页返回的起始偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页返回的记录条数，默认值：10，最大值：20。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribePlatformsRequest struct {
	*tchttp.BaseRequest
	
	// 平台 Id 列表。如果不填，则不按平台 Id 进行过滤。
	Platforms []*string `json:"Platforms,omitempty" name:"Platforms"`

	// 平台绑定的 License Id 列表。如果不填，则不按平台绑定的 License Id 进行过滤。
	LicenseIds []*string `json:"LicenseIds,omitempty" name:"LicenseIds"`

	// 分页返回的起始偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页返回的记录条数，默认值：10，最大值：20。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribePlatformsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePlatformsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePlatformsResponseParams struct {
	// 符合查询条件的记录总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 平台信息列表。
	PlatformInfoSet []*PlatformInfo `json:"PlatformInfoSet,omitempty" name:"PlatformInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePlatformsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePlatformsResponseParams `json:"Response"`
}

func (r *DescribePlatformsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePlatformsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectsRequestParams struct {
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 项目 Id 过滤参数列表，最大支持20个项目 Id 过滤。如果不填不需要项目 Id 进行过滤。
	ProjectIds []*string `json:"ProjectIds,omitempty" name:"ProjectIds"`

	// 画布宽高比过滤参数列表。如果不填则不用画布宽高比进行过滤。
	AspectRatioSet []*string `json:"AspectRatioSet,omitempty" name:"AspectRatioSet"`

	// 项目类型过滤参数列表，取值有：
	// <li>VIDEO_EDIT：视频编辑。</li>
	// <li>SWITCHER：导播台。</li>
	// <li>VIDEO_SEGMENTATION：视频拆条。</li>
	// <li>STREAM_CONNECT：云转推。</li>
	// <li>RECORD_REPLAY：录制回放。</li>
	// 
	// 注：如果不填则不使用项目类型进行过滤。
	CategorySet []*string `json:"CategorySet,omitempty" name:"CategorySet"`

	// 项目模式过滤参数列表，一个项目可以有多种模式并相互切换。
	// 当 Category 为 VIDEO_EDIT 时，可选模式有：
	// <li>Default：默认模式。</li>
	// <li>VideoEditTemplate：视频编辑模板制作模式。</li>
	// 
	// 注：不填不使用项目模式进行过滤。
	Modes []*string `json:"Modes,omitempty" name:"Modes"`

	// 结果排序方式，支持下列排序字段：
	// <li>CreateTime：创建时间；</li>
	// <li>UpdateTime：更新时间。</li>
	// 
	// 注：如不填，则使用项目创建时间倒序排列。
	Sort *SortBy `json:"Sort,omitempty" name:"Sort"`

	// 项目所有者，目前仅支持个人项目过滤。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 分页返回的起始偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页返回的记录条数，默认值：10。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以查询一切用户项目信息。如果指定操作者，则操作者必须为项目所有者。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type DescribeProjectsRequest struct {
	*tchttp.BaseRequest
	
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 项目 Id 过滤参数列表，最大支持20个项目 Id 过滤。如果不填不需要项目 Id 进行过滤。
	ProjectIds []*string `json:"ProjectIds,omitempty" name:"ProjectIds"`

	// 画布宽高比过滤参数列表。如果不填则不用画布宽高比进行过滤。
	AspectRatioSet []*string `json:"AspectRatioSet,omitempty" name:"AspectRatioSet"`

	// 项目类型过滤参数列表，取值有：
	// <li>VIDEO_EDIT：视频编辑。</li>
	// <li>SWITCHER：导播台。</li>
	// <li>VIDEO_SEGMENTATION：视频拆条。</li>
	// <li>STREAM_CONNECT：云转推。</li>
	// <li>RECORD_REPLAY：录制回放。</li>
	// 
	// 注：如果不填则不使用项目类型进行过滤。
	CategorySet []*string `json:"CategorySet,omitempty" name:"CategorySet"`

	// 项目模式过滤参数列表，一个项目可以有多种模式并相互切换。
	// 当 Category 为 VIDEO_EDIT 时，可选模式有：
	// <li>Default：默认模式。</li>
	// <li>VideoEditTemplate：视频编辑模板制作模式。</li>
	// 
	// 注：不填不使用项目模式进行过滤。
	Modes []*string `json:"Modes,omitempty" name:"Modes"`

	// 结果排序方式，支持下列排序字段：
	// <li>CreateTime：创建时间；</li>
	// <li>UpdateTime：更新时间。</li>
	// 
	// 注：如不填，则使用项目创建时间倒序排列。
	Sort *SortBy `json:"Sort,omitempty" name:"Sort"`

	// 项目所有者，目前仅支持个人项目过滤。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 分页返回的起始偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页返回的记录条数，默认值：10。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以查询一切用户项目信息。如果指定操作者，则操作者必须为项目所有者。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectsResponseParams struct {
	// 符合条件的记录总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 项目信息列表。
	ProjectInfoSet []*ProjectInfo `json:"ProjectInfoSet,omitempty" name:"ProjectInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeResourceAuthorizationRequestParams struct {
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 归属者。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 资源。
	Resource *Resource `json:"Resource,omitempty" name:"Resource"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以查询任意资源的被授权情况。如果指定操作者，则操作者必须对被授权资源有读权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type DescribeResourceAuthorizationRequest struct {
	*tchttp.BaseRequest
	
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 归属者。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 资源。
	Resource *Resource `json:"Resource,omitempty" name:"Resource"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以查询任意资源的被授权情况。如果指定操作者，则操作者必须对被授权资源有读权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *DescribeResourceAuthorizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourceAuthorizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceAuthorizationResponseParams struct {
	// 符合条件的资源授权记录总数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 授权信息列表。
	AuthorizationInfoSet []*AuthorizationInfo `json:"AuthorizationInfoSet,omitempty" name:"AuthorizationInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeResourceAuthorizationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourceAuthorizationResponseParams `json:"Response"`
}

func (r *DescribeResourceAuthorizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceAuthorizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSharedSpaceRequestParams struct {
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 被授权目标，个人或团队。
	Authorizee *Entity `json:"Authorizee,omitempty" name:"Authorizee"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以查询任意个人或者团队的共享空间。如果指定操作者，则操作者必须本人或者团队成员。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type DescribeSharedSpaceRequest struct {
	*tchttp.BaseRequest
	
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 被授权目标，个人或团队。
	Authorizee *Entity `json:"Authorizee,omitempty" name:"Authorizee"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以查询任意个人或者团队的共享空间。如果指定操作者，则操作者必须本人或者团队成员。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *DescribeSharedSpaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSharedSpaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSharedSpaceResponseParams struct {
	// 查询到的共享空间总数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 各个共享空间对应的授权者信息。
	AuthorizerSet []*Authorizer `json:"AuthorizerSet,omitempty" name:"AuthorizerSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSharedSpaceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSharedSpaceResponseParams `json:"Response"`
}

func (r *DescribeSharedSpaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSharedSpaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskDetailRequestParams struct {
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 任务 Id。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以获取任意任务信息。如果指定操作者，则操作者需要是任务发起者。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type DescribeTaskDetailRequest struct {
	*tchttp.BaseRequest
	
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 任务 Id。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以获取任意任务信息。如果指定操作者，则操作者需要是任务发起者。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *DescribeTaskDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskDetailResponseParams struct {
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

	// 导出项目输出信息。仅当 TaskType 为 VIDEO_EDIT_PROJECT_EXPORT 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoEditProjectOutput *VideoEditProjectOutput `json:"VideoEditProjectOutput,omitempty" name:"VideoEditProjectOutput"`

	// 创建时间，格式按照 ISO 8601 标准表示。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTaskDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskDetailResponseParams `json:"Response"`
}

func (r *DescribeTaskDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTasksRequestParams struct {
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 项目 Id，使用项目 Id 进行过滤。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务类型集合，取值有：
	// <li>VIDEO_EDIT_PROJECT_EXPORT：视频编辑项目导出。</li>
	// 
	// 注：不填不使用任务类型进行过滤。
	TaskTypeSet []*string `json:"TaskTypeSet,omitempty" name:"TaskTypeSet"`

	// 任务状态集合，取值有：
	// <li>PROCESSING：处理中；</li>
	// <li>SUCCESS：成功；</li>
	// <li>FAIL：失败。</li>
	// 
	// 注：不填则不使用任务状态进行过滤。
	StatusSet []*string `json:"StatusSet,omitempty" name:"StatusSet"`

	// 分页返回的起始偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页返回的记录条数，默认值：10。最大值：20。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以获取所有任务信息。如果指定操作者，则操作者需要是任务发起者。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type DescribeTasksRequest struct {
	*tchttp.BaseRequest
	
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 项目 Id，使用项目 Id 进行过滤。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务类型集合，取值有：
	// <li>VIDEO_EDIT_PROJECT_EXPORT：视频编辑项目导出。</li>
	// 
	// 注：不填不使用任务类型进行过滤。
	TaskTypeSet []*string `json:"TaskTypeSet,omitempty" name:"TaskTypeSet"`

	// 任务状态集合，取值有：
	// <li>PROCESSING：处理中；</li>
	// <li>SUCCESS：成功；</li>
	// <li>FAIL：失败。</li>
	// 
	// 注：不填则不使用任务状态进行过滤。
	StatusSet []*string `json:"StatusSet,omitempty" name:"StatusSet"`

	// 分页返回的起始偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页返回的记录条数，默认值：10。最大值：20。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以获取所有任务信息。如果指定操作者，则操作者需要是任务发起者。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *DescribeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTasksResponseParams struct {
	// 符合搜索条件的记录总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 任务基础信息列表。
	TaskBaseInfoSet []*TaskBaseInfo `json:"TaskBaseInfoSet,omitempty" name:"TaskBaseInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTasksResponseParams `json:"Response"`
}

func (r *DescribeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTeamMembersRequestParams struct {
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 团队 ID。
	TeamId *string `json:"TeamId,omitempty" name:"TeamId"`

	// 成员 ID 列表，限指定30个指定成员。如不填，则返回指定团队下的所有成员。
	MemberIds []*string `json:"MemberIds,omitempty" name:"MemberIds"`

	// 分页偏移量，默认值：0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回记录条数，默认值：30，最大值：30。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以拉取任意团队成员的信息。如果指定操作者，则操作者必须为团队成员。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type DescribeTeamMembersRequest struct {
	*tchttp.BaseRequest
	
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 团队 ID。
	TeamId *string `json:"TeamId,omitempty" name:"TeamId"`

	// 成员 ID 列表，限指定30个指定成员。如不填，则返回指定团队下的所有成员。
	MemberIds []*string `json:"MemberIds,omitempty" name:"MemberIds"`

	// 分页偏移量，默认值：0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回记录条数，默认值：30，最大值：30。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以拉取任意团队成员的信息。如果指定操作者，则操作者必须为团队成员。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *DescribeTeamMembersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTeamMembersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTeamMembersResponseParams struct {
	// 符合条件的记录总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 团队成员列表。
	MemberSet []*TeamMemberInfo `json:"MemberSet,omitempty" name:"MemberSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTeamMembersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTeamMembersResponseParams `json:"Response"`
}

func (r *DescribeTeamMembersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTeamMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTeamsRequestParams struct {
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 团队 ID 列表，限30个。若不填，则默认获取平台下所有团队。
	TeamIds []*string `json:"TeamIds,omitempty" name:"TeamIds"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回记录条数，默认值：20，最大值：30。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeTeamsRequest struct {
	*tchttp.BaseRequest
	
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 团队 ID 列表，限30个。若不填，则默认获取平台下所有团队。
	TeamIds []*string `json:"TeamIds,omitempty" name:"TeamIds"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回记录条数，默认值：20，最大值：30。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTeamsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTeamsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTeamsResponseParams struct {
	// 符合条件的记录总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 团队列表。
	TeamSet []*TeamInfo `json:"TeamSet,omitempty" name:"TeamSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTeamsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTeamsResponseParams `json:"Response"`
}

func (r *DescribeTeamsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTeamsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoEncodingPresetsRequestParams struct {
	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 要查询的配置 ID 列表。填写该参数则按照配置 ID 进行查询。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`

	// 分页大小，默认20。最大值50。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页起始，默认0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeVideoEncodingPresetsRequest struct {
	*tchttp.BaseRequest
	
	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 要查询的配置 ID 列表。填写该参数则按照配置 ID 进行查询。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`

	// 分页大小，默认20。最大值50。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页起始，默认0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeVideoEncodingPresetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoEncodingPresetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "Ids")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVideoEncodingPresetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoEncodingPresetsResponseParams struct {
	// 符合条件的编码配置总个数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 视频编码配置信息。
	VideoEncodingPresetSet []*VideoEncodingPreset `json:"VideoEncodingPresetSet,omitempty" name:"VideoEncodingPresetSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVideoEncodingPresetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVideoEncodingPresetsResponseParams `json:"Response"`
}

func (r *DescribeVideoEncodingPresetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoEncodingPresetsResponse) FromJsonString(s string) error {
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
	// 事件类型，可取值有：
	// <li>Storage.NewFileCreated：新文件产生事件；</li>
	// <li>Project.StreamConnect.StatusChanged：云转推项目状态变更事件；</li>
	// <li>Project.Switcher.StatusChanged：导播台项目状态变更事件；</li>
	// <li>Material.Imported：媒体导入事件；</li>
	// <li>Material.Added：媒体添加事件；</li>
	// <li>Material.Moved：媒体移动事件；</li>
	// <li>Material.Modified：媒体变更事件；</li>
	// <li>Material.Deleted：媒体删除事件；</li>
	// <li>Class.Created：分类新增事件；</li>
	// <li>Class.Moved：分类移动事件；</li>
	// <li>Class.Deleted：分类删除事件；</li>
	// <li>Task.VideoExportCompleted：视频导出完成事件； </li>
	// <li>Project.MediaCast.StatusChanged：点播转直播项目状态变更事件。 </li>
	EventType *string `json:"EventType,omitempty" name:"EventType"`

	// 操作者，表示触发事件的操作者。如果是 `cmeid_system` 表示平台管理员操作。
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 新文件产生事件。仅当 EventType 为 Storage.NewFileCreated 时有效。
	StorageNewFileCreatedEvent *StorageNewFileCreatedEvent `json:"StorageNewFileCreatedEvent,omitempty" name:"StorageNewFileCreatedEvent"`

	// 云转推项目状态变更事件。仅当 EventType 为 Project.StreamConnect.StatusChanged 时有效。
	ProjectStreamConnectStatusChangedEvent *ProjectStreamConnectStatusChangedEvent `json:"ProjectStreamConnectStatusChangedEvent,omitempty" name:"ProjectStreamConnectStatusChangedEvent"`

	// 导播台项目状态变更事件。仅当 EventType 为 Project.Switcher.StatusChanged 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectSwitcherStatusChangedEvent *ProjectSwitcherStatusChangedEvent `json:"ProjectSwitcherStatusChangedEvent,omitempty" name:"ProjectSwitcherStatusChangedEvent"`

	// 媒体导入事件。仅当 EventType 为 Material.Imported 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaterialImportedEvent *MaterialImportedEvent `json:"MaterialImportedEvent,omitempty" name:"MaterialImportedEvent"`

	// 媒体添加事件。仅当 EventType 为 Material.Added 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaterialAddedEvent *MaterialAddedEvent `json:"MaterialAddedEvent,omitempty" name:"MaterialAddedEvent"`

	// 媒体移动事件。仅当 EventType 为 Material.Moved 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaterialMovedEvent *MaterialMovedEvent `json:"MaterialMovedEvent,omitempty" name:"MaterialMovedEvent"`

	// 媒体更新事件。仅当 EventType 为 Material.Modified 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaterialModifiedEvent *MaterialModifiedEvent `json:"MaterialModifiedEvent,omitempty" name:"MaterialModifiedEvent"`

	// 媒体删除事件。仅当 EventType 为 Material.Deleted 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaterialDeletedEvent *MaterialDeletedEvent `json:"MaterialDeletedEvent,omitempty" name:"MaterialDeletedEvent"`

	// 分类创建事件。仅当 EventType 为 Class.Created 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClassCreatedEvent *ClassCreatedEvent `json:"ClassCreatedEvent,omitempty" name:"ClassCreatedEvent"`

	// 分类移动事件。仅当 EventType 为 Class.Moved 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClassMovedEvent *ClassMovedEvent `json:"ClassMovedEvent,omitempty" name:"ClassMovedEvent"`

	// 分类删除事件。仅当 EventType 为 Class.Deleted 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClassDeletedEvent *ClassDeletedEvent `json:"ClassDeletedEvent,omitempty" name:"ClassDeletedEvent"`

	// 视频导出完成事件。仅当 EventType 为 Task.VideoExportCompleted 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoExportCompletedEvent *VideoExportCompletedEvent `json:"VideoExportCompletedEvent,omitempty" name:"VideoExportCompletedEvent"`

	// 点播转直播项目状态变更事件。仅当 EventType 为 Project.MediaCast.StatusChanged 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectMediaCastStatusChangedEvent *ProjectMediaCastStatusChangedEvent `json:"ProjectMediaCastStatusChangedEvent,omitempty" name:"ProjectMediaCastStatusChangedEvent"`
}

// Predefined struct for user
type ExportVideoByEditorTrackDataRequestParams struct {
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 导出视频预设配置 Id，推荐优先使用下面的默认预设配置 Id，有其他需求可通过接口定制预设配置。
	// <li>10：分辨率为 480P，输出视频格式为 MP4；</li>
	// <li>11：分辨率为 720P，输出视频格式为 MP4；</li>
	// <li>12：分辨率为 1080P，输出视频格式为 MP4。</li>
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 导出目标，指定导出视频的目标媒资库，可取值有：
	// <li>CME：多媒体创建引擎，即导出到多媒体创作引擎媒资库，此导出目标在云点播媒资库依然可见；</li>
	// <li>VOD：云点播，即导出为云点播媒资库，此导出目标在多媒体创作引擎媒资库将不可见。</li>
	ExportDestination *string `json:"ExportDestination,omitempty" name:"ExportDestination"`

	// 轨道数据，用于描述待导出视频的内容。关于轨道数据的格式请查看 [视频合成协议](https://cloud.tencent.com/document/product/1156/51225)。文档中也描述了如何在页面上查看一个剪辑项目的轨道数据，该能力可以帮助开发者更方便地构造自己的轨道数据。
	TrackData *string `json:"TrackData,omitempty" name:"TrackData"`

	// 轨道数据对应的画布宽高比，配合预设配置中的视频短边尺寸，可决定导出画面的尺寸。例：
	// <li>如果 AspectRatio 取值 16:9，预设配置选为12（短边1080），则导出尺寸为 1920 * 1080；</li>
	// <li>如果 AspectRatio 取值 9:16，预设配置选为11（短边720），则导出尺寸为 720 *1280。</li>
	AspectRatio *string `json:"AspectRatio,omitempty" name:"AspectRatio"`

	// 视频封面图片文件（如 jpeg, png 等）进行 Base64 编码后的字符串，仅支持 gif、jpeg、png 三种图片格式，原图片文件不能超过2 M大 小。
	CoverData *string `json:"CoverData,omitempty" name:"CoverData"`

	// 导出的多媒体创作引擎媒体信息。当导出目标为 CME 时必填。
	CMEExportInfo *CMEExportInfo `json:"CMEExportInfo,omitempty" name:"CMEExportInfo"`

	// 导出的云点播媒资信息。当导出目标为 VOD 时必填。
	VODExportInfo *VODExportInfo `json:"VODExportInfo,omitempty" name:"VODExportInfo"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，无权限限制。如果指定操作者，轨道数据中使的媒资该操作者需要拥有使用权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type ExportVideoByEditorTrackDataRequest struct {
	*tchttp.BaseRequest
	
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 导出视频预设配置 Id，推荐优先使用下面的默认预设配置 Id，有其他需求可通过接口定制预设配置。
	// <li>10：分辨率为 480P，输出视频格式为 MP4；</li>
	// <li>11：分辨率为 720P，输出视频格式为 MP4；</li>
	// <li>12：分辨率为 1080P，输出视频格式为 MP4。</li>
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 导出目标，指定导出视频的目标媒资库，可取值有：
	// <li>CME：多媒体创建引擎，即导出到多媒体创作引擎媒资库，此导出目标在云点播媒资库依然可见；</li>
	// <li>VOD：云点播，即导出为云点播媒资库，此导出目标在多媒体创作引擎媒资库将不可见。</li>
	ExportDestination *string `json:"ExportDestination,omitempty" name:"ExportDestination"`

	// 轨道数据，用于描述待导出视频的内容。关于轨道数据的格式请查看 [视频合成协议](https://cloud.tencent.com/document/product/1156/51225)。文档中也描述了如何在页面上查看一个剪辑项目的轨道数据，该能力可以帮助开发者更方便地构造自己的轨道数据。
	TrackData *string `json:"TrackData,omitempty" name:"TrackData"`

	// 轨道数据对应的画布宽高比，配合预设配置中的视频短边尺寸，可决定导出画面的尺寸。例：
	// <li>如果 AspectRatio 取值 16:9，预设配置选为12（短边1080），则导出尺寸为 1920 * 1080；</li>
	// <li>如果 AspectRatio 取值 9:16，预设配置选为11（短边720），则导出尺寸为 720 *1280。</li>
	AspectRatio *string `json:"AspectRatio,omitempty" name:"AspectRatio"`

	// 视频封面图片文件（如 jpeg, png 等）进行 Base64 编码后的字符串，仅支持 gif、jpeg、png 三种图片格式，原图片文件不能超过2 M大 小。
	CoverData *string `json:"CoverData,omitempty" name:"CoverData"`

	// 导出的多媒体创作引擎媒体信息。当导出目标为 CME 时必填。
	CMEExportInfo *CMEExportInfo `json:"CMEExportInfo,omitempty" name:"CMEExportInfo"`

	// 导出的云点播媒资信息。当导出目标为 VOD 时必填。
	VODExportInfo *VODExportInfo `json:"VODExportInfo,omitempty" name:"VODExportInfo"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，无权限限制。如果指定操作者，轨道数据中使的媒资该操作者需要拥有使用权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *ExportVideoByEditorTrackDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
	delete(f, "AspectRatio")
	delete(f, "CoverData")
	delete(f, "CMEExportInfo")
	delete(f, "VODExportInfo")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportVideoByEditorTrackDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportVideoByEditorTrackDataResponseParams struct {
	// 任务 Id。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ExportVideoByEditorTrackDataResponse struct {
	*tchttp.BaseResponse
	Response *ExportVideoByEditorTrackDataResponseParams `json:"Response"`
}

func (r *ExportVideoByEditorTrackDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportVideoByEditorTrackDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportVideoByTemplateRequestParams struct {
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 视频编辑模板  Id。
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 导出视频预设配置 Id，推荐优先使用下面的默认预设配置 Id，有其他需求可通过接口定制预设配置。
	// <li>10：分辨率为 480P，输出视频格式为 MP4；</li>
	// <li>11：分辨率为 720P，输出视频格式为 MP4；</li>
	// <li>12：分辨率为 1080P，输出视频格式为 MP4。</li>
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 导出目标，指定导出视频的目标媒资库，可取值有：
	// <li>CME：多媒体创作引擎，即导出为多媒体创作引擎媒资库，此导出目标在云点播媒资库依然可见；</li>
	// <li>VOD：云点播，即导出为云点播媒资库，此导出目标在多媒体创作引擎媒资库将不可见。</li>
	ExportDestination *string `json:"ExportDestination,omitempty" name:"ExportDestination"`

	// 需要替换的素材信息。
	SlotReplacements []*SlotReplacementInfo `json:"SlotReplacements,omitempty" name:"SlotReplacements"`

	// 导出的多媒体创作引擎媒资信息。当导出目标为 CME 时必填。
	CMEExportInfo *CMEExportInfo `json:"CMEExportInfo,omitempty" name:"CMEExportInfo"`

	// 导出的云点播媒资信息。当导出目标为 VOD 时必填。
	VODExportInfo *VODExportInfo `json:"VODExportInfo,omitempty" name:"VODExportInfo"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，无权限限制。如果指定操作者，则操作者需要有替换媒体及剪辑模板的权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type ExportVideoByTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 视频编辑模板  Id。
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 导出视频预设配置 Id，推荐优先使用下面的默认预设配置 Id，有其他需求可通过接口定制预设配置。
	// <li>10：分辨率为 480P，输出视频格式为 MP4；</li>
	// <li>11：分辨率为 720P，输出视频格式为 MP4；</li>
	// <li>12：分辨率为 1080P，输出视频格式为 MP4。</li>
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 导出目标，指定导出视频的目标媒资库，可取值有：
	// <li>CME：多媒体创作引擎，即导出为多媒体创作引擎媒资库，此导出目标在云点播媒资库依然可见；</li>
	// <li>VOD：云点播，即导出为云点播媒资库，此导出目标在多媒体创作引擎媒资库将不可见。</li>
	ExportDestination *string `json:"ExportDestination,omitempty" name:"ExportDestination"`

	// 需要替换的素材信息。
	SlotReplacements []*SlotReplacementInfo `json:"SlotReplacements,omitempty" name:"SlotReplacements"`

	// 导出的多媒体创作引擎媒资信息。当导出目标为 CME 时必填。
	CMEExportInfo *CMEExportInfo `json:"CMEExportInfo,omitempty" name:"CMEExportInfo"`

	// 导出的云点播媒资信息。当导出目标为 VOD 时必填。
	VODExportInfo *VODExportInfo `json:"VODExportInfo,omitempty" name:"VODExportInfo"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，无权限限制。如果指定操作者，则操作者需要有替换媒体及剪辑模板的权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *ExportVideoByTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportVideoByTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportVideoByTemplateResponseParams struct {
	// 导出任务 Id。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ExportVideoByTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ExportVideoByTemplateResponseParams `json:"Response"`
}

func (r *ExportVideoByTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportVideoByTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportVideoByVideoSegmentationDataRequestParams struct {
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 视频拆条项目 Id 。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 指定需要导出的智能拆条片段的组 Id 。
	SegmentGroupId *string `json:"SegmentGroupId,omitempty" name:"SegmentGroupId"`

	// 指定需要导出的智能拆条片段 Id  集合。
	SegmentIds []*string `json:"SegmentIds,omitempty" name:"SegmentIds"`

	// 导出模板 Id，目前不支持自定义创建，只支持下面的预置模板 Id。
	// <li>10：分辨率为 480P，输出视频格式为 MP4；</li>
	// <li>11：分辨率为 720P，输出视频格式为 MP4；</li>
	// <li>12：分辨率为 1080P，输出视频格式为 MP4。</li>
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 导出目标，指定导出视频的目标媒资库，可取值有：
	// <li>CME：多媒体创作引擎，即导出为多媒体创作引擎媒资库，此导出目标在云点播媒资库依然可见；</li>
	// <li>VOD：云点播，即导出为云点播媒资库，此导出目标在多媒体创作引擎媒资库将不可见。</li>
	ExportDestination *string `json:"ExportDestination,omitempty" name:"ExportDestination"`

	// 导出的多媒体创作引擎媒体信息。当导出目标为 CME 时必填。
	CMEExportInfo *CMEExportInfo `json:"CMEExportInfo,omitempty" name:"CMEExportInfo"`

	// 导出的云点播媒资信息。当导出目标为 VOD 时必填。
	VODExportInfo *VODExportInfo `json:"VODExportInfo,omitempty" name:"VODExportInfo"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以操作任意智能拆条项目。如果指定操作者，则操作者必须为项目所有。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type ExportVideoByVideoSegmentationDataRequest struct {
	*tchttp.BaseRequest
	
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 视频拆条项目 Id 。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 指定需要导出的智能拆条片段的组 Id 。
	SegmentGroupId *string `json:"SegmentGroupId,omitempty" name:"SegmentGroupId"`

	// 指定需要导出的智能拆条片段 Id  集合。
	SegmentIds []*string `json:"SegmentIds,omitempty" name:"SegmentIds"`

	// 导出模板 Id，目前不支持自定义创建，只支持下面的预置模板 Id。
	// <li>10：分辨率为 480P，输出视频格式为 MP4；</li>
	// <li>11：分辨率为 720P，输出视频格式为 MP4；</li>
	// <li>12：分辨率为 1080P，输出视频格式为 MP4。</li>
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 导出目标，指定导出视频的目标媒资库，可取值有：
	// <li>CME：多媒体创作引擎，即导出为多媒体创作引擎媒资库，此导出目标在云点播媒资库依然可见；</li>
	// <li>VOD：云点播，即导出为云点播媒资库，此导出目标在多媒体创作引擎媒资库将不可见。</li>
	ExportDestination *string `json:"ExportDestination,omitempty" name:"ExportDestination"`

	// 导出的多媒体创作引擎媒体信息。当导出目标为 CME 时必填。
	CMEExportInfo *CMEExportInfo `json:"CMEExportInfo,omitempty" name:"CMEExportInfo"`

	// 导出的云点播媒资信息。当导出目标为 VOD 时必填。
	VODExportInfo *VODExportInfo `json:"VODExportInfo,omitempty" name:"VODExportInfo"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以操作任意智能拆条项目。如果指定操作者，则操作者必须为项目所有。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *ExportVideoByVideoSegmentationDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportVideoByVideoSegmentationDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportVideoByVideoSegmentationDataResponseParams struct {
	// 任务 Id。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ExportVideoByVideoSegmentationDataResponse struct {
	*tchttp.BaseResponse
	Response *ExportVideoByVideoSegmentationDataResponseParams `json:"Response"`
}

func (r *ExportVideoByVideoSegmentationDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportVideoByVideoSegmentationDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportVideoEditProjectRequestParams struct {
	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 项目 Id。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 视频编码配置 ID，支持自定义创建，推荐优先使用系统预置的导出配置。
	// <li>10：分辨率为 480P，输出视频格式为 MP4；</li>
	// <li>11：分辨率为 720P，输出视频格式为 MP4；</li>
	// <li>12：分辨率为 1080P，输出视频格式为 MP4。</li>
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 导出目标，指定导出视频的目标媒资库，可取值有：
	// <li>CME：多媒体创作引擎，即导出为多媒体创作引擎媒资库，此导出目标在云点播媒资库依然可见；</li>
	// <li>VOD：云点播，即导出为云点播媒资库，此导出目标在多媒体创作引擎媒资库将不可见。</li>
	ExportDestination *string `json:"ExportDestination,omitempty" name:"ExportDestination"`

	// 视频封面图片文件（如 jpeg, png 等）进行 Base64 编码后的字符串，仅支持 gif、jpeg、png 三种图片格式，原图片文件不能超过2 M大 小。
	CoverData *string `json:"CoverData,omitempty" name:"CoverData"`

	// 导出的多媒体创作引擎媒体信息。当导出目标为 CME 时必填。
	CMEExportInfo *CMEExportInfo `json:"CMEExportInfo,omitempty" name:"CMEExportInfo"`

	// 导出的云点播媒资信息。当导出目标为 VOD 时必填。
	VODExportInfo *VODExportInfo `json:"VODExportInfo,omitempty" name:"VODExportInfo"`

	// 视频导出扩展参数。可以覆盖导出模板中的参数，灵活的指定导出规格及参数。
	ExportExtensionArgs *VideoExportExtensionArgs `json:"ExportExtensionArgs,omitempty" name:"ExportExtensionArgs"`

	// 操作者。填写用户的 Id，用于标识调用者及校验项目导出权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type ExportVideoEditProjectRequest struct {
	*tchttp.BaseRequest
	
	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 项目 Id。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 视频编码配置 ID，支持自定义创建，推荐优先使用系统预置的导出配置。
	// <li>10：分辨率为 480P，输出视频格式为 MP4；</li>
	// <li>11：分辨率为 720P，输出视频格式为 MP4；</li>
	// <li>12：分辨率为 1080P，输出视频格式为 MP4。</li>
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 导出目标，指定导出视频的目标媒资库，可取值有：
	// <li>CME：多媒体创作引擎，即导出为多媒体创作引擎媒资库，此导出目标在云点播媒资库依然可见；</li>
	// <li>VOD：云点播，即导出为云点播媒资库，此导出目标在多媒体创作引擎媒资库将不可见。</li>
	ExportDestination *string `json:"ExportDestination,omitempty" name:"ExportDestination"`

	// 视频封面图片文件（如 jpeg, png 等）进行 Base64 编码后的字符串，仅支持 gif、jpeg、png 三种图片格式，原图片文件不能超过2 M大 小。
	CoverData *string `json:"CoverData,omitempty" name:"CoverData"`

	// 导出的多媒体创作引擎媒体信息。当导出目标为 CME 时必填。
	CMEExportInfo *CMEExportInfo `json:"CMEExportInfo,omitempty" name:"CMEExportInfo"`

	// 导出的云点播媒资信息。当导出目标为 VOD 时必填。
	VODExportInfo *VODExportInfo `json:"VODExportInfo,omitempty" name:"VODExportInfo"`

	// 视频导出扩展参数。可以覆盖导出模板中的参数，灵活的指定导出规格及参数。
	ExportExtensionArgs *VideoExportExtensionArgs `json:"ExportExtensionArgs,omitempty" name:"ExportExtensionArgs"`

	// 操作者。填写用户的 Id，用于标识调用者及校验项目导出权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *ExportVideoEditProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
	delete(f, "ExportExtensionArgs")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportVideoEditProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportVideoEditProjectResponseParams struct {
	// 任务 Id。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ExportVideoEditProjectResponse struct {
	*tchttp.BaseResponse
	Response *ExportVideoEditProjectResponseParams `json:"Response"`
}

func (r *ExportVideoEditProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportVideoEditProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExternalMediaInfo struct {
	// 目前仅支持绑定 COS 桶的媒体，请填写存储对象 Key 值，例如：`example-folder/example.mp4`。
	MediaKey *string `json:"MediaKey,omitempty" name:"MediaKey"`

	// 该字段废弃，请勿使用。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 媒资挂载的存储 Id。
	StorageId *string `json:"StorageId,omitempty" name:"StorageId"`
}

// Predefined struct for user
type FlattenListMediaRequestParams struct {
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 媒体分类路径，例如填写"/a/b"，则代表平铺该分类路径下及其子分类路径下的媒体信息。
	ClassPath *string `json:"ClassPath,omitempty" name:"ClassPath"`

	// 媒体分类的归属者。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 分页偏移量，默认值：0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：50。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以平铺查询任意分类下的媒体信息。如果指定操作者，则操作者必须对当前分类有读权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type FlattenListMediaRequest struct {
	*tchttp.BaseRequest
	
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 媒体分类路径，例如填写"/a/b"，则代表平铺该分类路径下及其子分类路径下的媒体信息。
	ClassPath *string `json:"ClassPath,omitempty" name:"ClassPath"`

	// 媒体分类的归属者。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 分页偏移量，默认值：0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：50。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以平铺查询任意分类下的媒体信息。如果指定操作者，则操作者必须对当前分类有读权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *FlattenListMediaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FlattenListMediaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FlattenListMediaResponseParams struct {
	// 符合条件的记录总数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 该分类路径下及其子分类下的所有媒体基础信息列表。
	MaterialInfoSet []*MaterialInfo `json:"MaterialInfoSet,omitempty" name:"MaterialInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type FlattenListMediaResponse struct {
	*tchttp.BaseResponse
	Response *FlattenListMediaResponseParams `json:"Response"`
}

func (r *FlattenListMediaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FlattenListMediaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GenerateVideoSegmentationSchemeByAiRequestParams struct {
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 视频拆条项目 Id 。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以对任务视频拆条项目发起拆条任务。如果指定操作者，则操作者必须为项目所有者。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type GenerateVideoSegmentationSchemeByAiRequest struct {
	*tchttp.BaseRequest
	
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 视频拆条项目 Id 。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以对任务视频拆条项目发起拆条任务。如果指定操作者，则操作者必须为项目所有者。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *GenerateVideoSegmentationSchemeByAiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GenerateVideoSegmentationSchemeByAiRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GenerateVideoSegmentationSchemeByAiResponseParams struct {
	// 视频智能拆条任务 Id 。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GenerateVideoSegmentationSchemeByAiResponse struct {
	*tchttp.BaseResponse
	Response *GenerateVideoSegmentationSchemeByAiResponseParams `json:"Response"`
}

func (r *GenerateVideoSegmentationSchemeByAiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateVideoSegmentationSchemeByAiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GrantResourceAuthorizationRequestParams struct {
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 资源归属者，个人或者团队。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 被授权资源。
	Resources []*Resource `json:"Resources,omitempty" name:"Resources"`

	// 被授权目标，个人或者团队。
	Authorizees []*Entity `json:"Authorizees,omitempty" name:"Authorizees"`

	// 详细授权值。 取值有：
	// <li>R：可读，可以浏览媒体，但不能使用该媒体文件（将其添加到 Project），或复制到自己的媒资库中</li>
	// <li>X：可用，可以使用该素材（将其添加到 Project），但不能将其复制到自己的媒资库中，意味着被授权者无法将该资源进一步扩散给其他个人或团队。</li>
	// <li>C：可复制，既可以使用该素材（将其添加到 Project），也可以将其复制到自己的媒资库中。</li>
	// <li>W：可修改、删除媒资。</li>
	Permissions []*string `json:"Permissions,omitempty" name:"Permissions"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以授权任意归属者的资源。如果指定操作者，则操作者必须对资源拥有写权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type GrantResourceAuthorizationRequest struct {
	*tchttp.BaseRequest
	
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 资源归属者，个人或者团队。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 被授权资源。
	Resources []*Resource `json:"Resources,omitempty" name:"Resources"`

	// 被授权目标，个人或者团队。
	Authorizees []*Entity `json:"Authorizees,omitempty" name:"Authorizees"`

	// 详细授权值。 取值有：
	// <li>R：可读，可以浏览媒体，但不能使用该媒体文件（将其添加到 Project），或复制到自己的媒资库中</li>
	// <li>X：可用，可以使用该素材（将其添加到 Project），但不能将其复制到自己的媒资库中，意味着被授权者无法将该资源进一步扩散给其他个人或团队。</li>
	// <li>C：可复制，既可以使用该素材（将其添加到 Project），也可以将其复制到自己的媒资库中。</li>
	// <li>W：可修改、删除媒资。</li>
	Permissions []*string `json:"Permissions,omitempty" name:"Permissions"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以授权任意归属者的资源。如果指定操作者，则操作者必须对资源拥有写权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *GrantResourceAuthorizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GrantResourceAuthorizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GrantResourceAuthorizationResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GrantResourceAuthorizationResponse struct {
	*tchttp.BaseResponse
	Response *GrantResourceAuthorizationResponseParams `json:"Response"`
}

func (r *GrantResourceAuthorizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GrantResourceAuthorizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type HandleStreamConnectProjectRequestParams struct {
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 云转推项目 Id 。
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

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以操作所有云转推项目。如果指定操作者，则操作者必须为项目所有者。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type HandleStreamConnectProjectRequest struct {
	*tchttp.BaseRequest
	
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 云转推项目 Id 。
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

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以操作所有云转推项目。如果指定操作者，则操作者必须为项目所有者。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *HandleStreamConnectProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "HandleStreamConnectProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type HandleStreamConnectProjectResponseParams struct {
	// 输入源推流地址，当 Operation 取值 AddInput 且 InputType 为 RtmpPush 类型时有效。
	StreamInputRtmpPushUrl *string `json:"StreamInputRtmpPushUrl,omitempty" name:"StreamInputRtmpPushUrl"`

	// 点播输入源播放进度信息，当 Operation 取值 DescribeInputPlayInfo 且 InputType 为 VodPull 类型时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VodPullInputPlayInfo *VodPullInputPlayInfo `json:"VodPullInputPlayInfo,omitempty" name:"VodPullInputPlayInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type HandleStreamConnectProjectResponse struct {
	*tchttp.BaseResponse
	Response *HandleStreamConnectProjectResponseParams `json:"Response"`
}

func (r *HandleStreamConnectProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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

// Predefined struct for user
type ImportMaterialRequestParams struct {
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 媒体归属者，可支持归属团队或个人。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 媒体名称，不能超过30个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 导入媒资类型，取值：
	// <li>VOD：云点播文件；</li>
	// <li>EXTERNAL：媒资绑定。</li>
	// 
	// 注意：如果不填默认为云点播文件，如果媒体存储在非腾讯云点播中，都需要使用媒资绑定。另外，导入云点播的文件，使用云点播的子应用 Id 必须与创建多媒体创作引擎平台时使用的云点播子应用一致。
	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`

	// 云点播媒资 FileId，仅当 SourceType 为 VOD 时有效。
	VodFileId *string `json:"VodFileId,omitempty" name:"VodFileId"`

	// 原始媒资文件信息，当 SourceType 取值 EXTERNAL 的时候必填。
	ExternalMediaInfo *ExternalMediaInfo `json:"ExternalMediaInfo,omitempty" name:"ExternalMediaInfo"`

	// 媒体分类路径，形如："/a/b"，层级数不能超过10，每个层级长度不能超过15字符。若不填则默认为根路径。
	ClassPath *string `json:"ClassPath,omitempty" name:"ClassPath"`

	// 媒体预处理任务参数 ID。可取值有：
	// <li>10：进行编辑预处理。</li>
	PreProcessDefinition *int64 `json:"PreProcessDefinition,omitempty" name:"PreProcessDefinition"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以向任意团队或者个人导入媒体。如果指定操作者，如果媒体归属为个人，则操作者必须与归属者一致；如果媒体归属为团队，则必须为团队可导入媒体的团队成员(如果没有特殊设置，所有团队成员可导入媒体)。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type ImportMaterialRequest struct {
	*tchttp.BaseRequest
	
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 媒体归属者，可支持归属团队或个人。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 媒体名称，不能超过30个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 导入媒资类型，取值：
	// <li>VOD：云点播文件；</li>
	// <li>EXTERNAL：媒资绑定。</li>
	// 
	// 注意：如果不填默认为云点播文件，如果媒体存储在非腾讯云点播中，都需要使用媒资绑定。另外，导入云点播的文件，使用云点播的子应用 Id 必须与创建多媒体创作引擎平台时使用的云点播子应用一致。
	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`

	// 云点播媒资 FileId，仅当 SourceType 为 VOD 时有效。
	VodFileId *string `json:"VodFileId,omitempty" name:"VodFileId"`

	// 原始媒资文件信息，当 SourceType 取值 EXTERNAL 的时候必填。
	ExternalMediaInfo *ExternalMediaInfo `json:"ExternalMediaInfo,omitempty" name:"ExternalMediaInfo"`

	// 媒体分类路径，形如："/a/b"，层级数不能超过10，每个层级长度不能超过15字符。若不填则默认为根路径。
	ClassPath *string `json:"ClassPath,omitempty" name:"ClassPath"`

	// 媒体预处理任务参数 ID。可取值有：
	// <li>10：进行编辑预处理。</li>
	PreProcessDefinition *int64 `json:"PreProcessDefinition,omitempty" name:"PreProcessDefinition"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以向任意团队或者个人导入媒体。如果指定操作者，如果媒体归属为个人，则操作者必须与归属者一致；如果媒体归属为团队，则必须为团队可导入媒体的团队成员(如果没有特殊设置，所有团队成员可导入媒体)。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *ImportMaterialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImportMaterialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportMaterialResponseParams struct {
	// 媒体 Id。
	MaterialId *string `json:"MaterialId,omitempty" name:"MaterialId"`

	// 媒体文预处理任务 ID，如果未指定发起预处理任务则为空。
	PreProcessTaskId *string `json:"PreProcessTaskId,omitempty" name:"PreProcessTaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ImportMaterialResponse struct {
	*tchttp.BaseResponse
	Response *ImportMaterialResponseParams `json:"Response"`
}

func (r *ImportMaterialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportMaterialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportMediaInfo struct {
	// 云点播文件 FileId。
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 媒体 Id。
	MaterialId *string `json:"MaterialId,omitempty" name:"MaterialId"`
}

// Predefined struct for user
type ImportMediaToProjectRequestParams struct {
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 项目 Id。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 导入媒资类型，取值：
	// <li>VOD：云点播文件；</li>
	// <li>EXTERNAL：媒资绑定。</li>
	// 
	// 注意：如果不填默认为云点播文件，如果媒体存储在非腾讯云点播中，都需要使用媒资绑定。
	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`

	// 云点播媒资文件 Id，当 SourceType 取值 VOD 或者缺省的时候必填。
	VodFileId *string `json:"VodFileId,omitempty" name:"VodFileId"`

	// 原始媒资文件信息，当 SourceType 取值 EXTERNAL 的时候必填。
	ExternalMediaInfo *ExternalMediaInfo `json:"ExternalMediaInfo,omitempty" name:"ExternalMediaInfo"`

	// 媒体名称，不能超过30个字符。如果不填，则媒体名称为点播媒资文件名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 媒体预处理配置 ID，取值：
	// <li>10：进行视频编辑预处理。</li>
	// 
	// 注意：如果填0或者不填则不进行处理，如果原始视频不可在浏览器直接播放将无法在编辑页面编辑。
	PreProcessDefinition *int64 `json:"PreProcessDefinition,omitempty" name:"PreProcessDefinition"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以向所有视频编辑项目导入媒体；如果指定操作者，则操作者必须为项目所有者。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type ImportMediaToProjectRequest struct {
	*tchttp.BaseRequest
	
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 项目 Id。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 导入媒资类型，取值：
	// <li>VOD：云点播文件；</li>
	// <li>EXTERNAL：媒资绑定。</li>
	// 
	// 注意：如果不填默认为云点播文件，如果媒体存储在非腾讯云点播中，都需要使用媒资绑定。
	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`

	// 云点播媒资文件 Id，当 SourceType 取值 VOD 或者缺省的时候必填。
	VodFileId *string `json:"VodFileId,omitempty" name:"VodFileId"`

	// 原始媒资文件信息，当 SourceType 取值 EXTERNAL 的时候必填。
	ExternalMediaInfo *ExternalMediaInfo `json:"ExternalMediaInfo,omitempty" name:"ExternalMediaInfo"`

	// 媒体名称，不能超过30个字符。如果不填，则媒体名称为点播媒资文件名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 媒体预处理配置 ID，取值：
	// <li>10：进行视频编辑预处理。</li>
	// 
	// 注意：如果填0或者不填则不进行处理，如果原始视频不可在浏览器直接播放将无法在编辑页面编辑。
	PreProcessDefinition *int64 `json:"PreProcessDefinition,omitempty" name:"PreProcessDefinition"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以向所有视频编辑项目导入媒体；如果指定操作者，则操作者必须为项目所有者。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *ImportMediaToProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImportMediaToProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportMediaToProjectResponseParams struct {
	// 媒体 Id。
	MaterialId *string `json:"MaterialId,omitempty" name:"MaterialId"`

	// 媒体预处理任务 ID，如果未指定发起预处理任务则为空。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ImportMediaToProjectResponse struct {
	*tchttp.BaseResponse
	Response *ImportMediaToProjectResponseParams `json:"Response"`
}

func (r *ImportMediaToProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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

	// 团队成员个数。
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

// Predefined struct for user
type ListMediaRequestParams struct {
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 媒体分类路径，例如填写"/a/b"，则代表浏览该分类路径下的媒体和子分类信息。
	ClassPath *string `json:"ClassPath,omitempty" name:"ClassPath"`

	// 媒体和分类的归属者。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 分页偏移量，默认值：0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：50。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以浏览任意分类的信息。如果指定操作者，则操作者必须对分类有读权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type ListMediaRequest struct {
	*tchttp.BaseRequest
	
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 媒体分类路径，例如填写"/a/b"，则代表浏览该分类路径下的媒体和子分类信息。
	ClassPath *string `json:"ClassPath,omitempty" name:"ClassPath"`

	// 媒体和分类的归属者。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 分页偏移量，默认值：0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：50。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以浏览任意分类的信息。如果指定操作者，则操作者必须对分类有读权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *ListMediaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListMediaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListMediaResponseParams struct {
	// 符合条件的媒体记录总数。
	MaterialTotalCount *int64 `json:"MaterialTotalCount,omitempty" name:"MaterialTotalCount"`

	// 浏览分类路径下的媒体列表信息。
	MaterialInfoSet []*MaterialInfo `json:"MaterialInfoSet,omitempty" name:"MaterialInfoSet"`

	// 浏览分类路径下的一级子类。
	ClassInfoSet []*ClassInfo `json:"ClassInfoSet,omitempty" name:"ClassInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListMediaResponse struct {
	*tchttp.BaseResponse
	Response *ListMediaResponseParams `json:"Response"`
}

func (r *ListMediaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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

type MaterialAddedEvent struct {
	// 添加的媒体 Id 列表。
	MaterialIdSet []*string `json:"MaterialIdSet,omitempty" name:"MaterialIdSet"`

	// 添加的媒体归属。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 添加的媒体分类路径。
	ClassPath *string `json:"ClassPath,omitempty" name:"ClassPath"`
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
	PresetTagSet []*PresetTagInfo `json:"PresetTagSet,omitempty" name:"PresetTagSet"`

	// 人工标签列表。
	TagSet []*string `json:"TagSet,omitempty" name:"TagSet"`

	// 媒体文件的预览图。
	PreviewUrl *string `json:"PreviewUrl,omitempty" name:"PreviewUrl"`

	// 媒体绑定的标签信息列表 。
	// 该字段已废弃。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagInfoSet []*MaterialTagInfo `json:"TagInfoSet,omitempty" name:"TagInfoSet"`
}

type MaterialDeletedEvent struct {
	// 删除的媒体 Id 列表。
	MaterialIdSet []*string `json:"MaterialIdSet,omitempty" name:"MaterialIdSet"`
}

type MaterialImportedEvent struct {
	// 导入的媒体信息列表。
	MediaInfoSet []*ImportMediaInfo `json:"MediaInfoSet,omitempty" name:"MediaInfoSet"`

	// 媒体归属。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 媒体分类路径。
	ClassPath *string `json:"ClassPath,omitempty" name:"ClassPath"`
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

type MaterialModifiedEvent struct {
	// 媒体 Id。
	MaterialId *string `json:"MaterialId,omitempty" name:"MaterialId"`

	// 更新后的媒体名称。如未更新则为空。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 更新后的媒体预置标签列表。如未更新媒体预置标签，则该字段为空数组。
	PresetTagIdSet []*string `json:"PresetTagIdSet,omitempty" name:"PresetTagIdSet"`

	// 更新后的媒体自定义标签列表。如未更新媒体自定义标签，则该字段为空数组。
	TagSet []*string `json:"TagSet,omitempty" name:"TagSet"`
}

type MaterialMovedEvent struct {
	// 要移动的媒体 Id 列表。
	MaterialIdSet []*string `json:"MaterialIdSet,omitempty" name:"MaterialIdSet"`

	// 源媒体归属。
	SourceOwner *Entity `json:"SourceOwner,omitempty" name:"SourceOwner"`

	// 源媒体分类路径。
	SourceClassPath *string `json:"SourceClassPath,omitempty" name:"SourceClassPath"`

	// 目标媒体分类归属。
	DestinationOwner *Entity `json:"DestinationOwner,omitempty" name:"DestinationOwner"`

	// 目标媒体分类路径。
	DestinationClassPath *string `json:"DestinationClassPath,omitempty" name:"DestinationClassPath"`
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

type MediaCastDestinationInfo struct {
	// 输出源序号。由系统进行分配。
	Index *int64 `json:"Index,omitempty" name:"Index"`

	// 输出源的名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 输出直播流地址。支持的直播流类型为 RTMP 和 SRT。
	PushUrl *string `json:"PushUrl,omitempty" name:"PushUrl"`
}

type MediaCastDestinationInterruptInfo struct {
	// 发生断流的输出源信息。
	DestinationInfo *MediaCastDestinationInfo `json:"DestinationInfo,omitempty" name:"DestinationInfo"`

	// 输出源断流原因，取值有：
	// <li>SystemError：系统错误；</li>
	// <li>Unknown：未知错误。</li>
	Reason *string `json:"Reason,omitempty" name:"Reason"`
}

type MediaCastOutputMediaSetting struct {
	// 视频配置。
	VideoSetting *MediaCastVideoSetting `json:"VideoSetting,omitempty" name:"VideoSetting"`
}

type MediaCastPlaySetting struct {
	// 循环播放次数。LoopCount 和 EndTime 同时只能有一个生效。默认循环播放次数为一次。如果同时设置了 LoopCount 和 EndTime 参数，优先使用 LoopCount 参数。
	LoopCount *int64 `json:"LoopCount,omitempty" name:"LoopCount"`

	// 结束时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type MediaCastProjectInfo struct {
	// 点播转直播项目状态，取值有：
	// <li>Working ：运行中；</li>
	// <li>Idle ：空闲。</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 输入源列表。
	SourceInfos []*MediaCastSourceInfo `json:"SourceInfos,omitempty" name:"SourceInfos"`

	// 输出源列表。
	DestinationInfos []*MediaCastDestinationInfo `json:"DestinationInfos,omitempty" name:"DestinationInfos"`

	// 输出媒体配置。
	OutputMediaSetting *MediaCastOutputMediaSetting `json:"OutputMediaSetting,omitempty" name:"OutputMediaSetting"`

	// 播放参数。
	PlaySetting *MediaCastPlaySetting `json:"PlaySetting,omitempty" name:"PlaySetting"`

	// 项目启动时间。采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 项目结束时间。采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。如果项目还在运行中，改字段为空。
	StopTime *string `json:"StopTime,omitempty" name:"StopTime"`
}

type MediaCastSourceInfo struct {
	// 输入源的媒体类型，取值有：
	// <li>CME：多媒体创作引擎的媒体文件；</li>
	// <li>VOD：云点播的媒资文件。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 多媒体创作引擎的媒体 ID。当 Type = CME  时必填。
	MaterialId *string `json:"MaterialId,omitempty" name:"MaterialId"`

	// 云点播媒体文件 ID。当 Type = VOD 时必填。
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 序号，位于输入源列表中的序号，由系统分配。
	Index *int64 `json:"Index,omitempty" name:"Index"`
}

type MediaCastSourceInterruptInfo struct {
	// 发生断流的输入源信息。
	SourceInfo *MediaCastSourceInfo `json:"SourceInfo,omitempty" name:"SourceInfo"`

	// 输入源断开原因。取值有：
	// <li>SystemError：系统错误；</li>
	// <li>Unknown：未知错误。</li>
	Reason *string `json:"Reason,omitempty" name:"Reason"`
}

type MediaCastVideoSetting struct {
	// 视频宽度，单位：px，默认值为1280。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 视频高度，单位：px，默认值为720。支持的视频分辨率最大为1920*1080。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 视频码率，单位：kbps，默认值为2500。最大值为10000 kbps。
	Bitrate *int64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 视频帧率，单位：Hz，默认值为25。最大值为60。
	FrameRate *float64 `json:"FrameRate,omitempty" name:"FrameRate"`
}

type MediaImageSpriteInfo struct {
	// 雪碧图小图的高度。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 雪碧图小图的宽度。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 雪碧图小图的总数量。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 截取雪碧图输出的地址。
	ImageUrlSet []*string `json:"ImageUrlSet,omitempty" name:"ImageUrlSet"`

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
	VideoStreamInfoSet []*VideoStreamInfo `json:"VideoStreamInfoSet,omitempty" name:"VideoStreamInfoSet"`

	// 音频流信息。
	AudioStreamInfoSet []*AudioStreamInfo `json:"AudioStreamInfoSet,omitempty" name:"AudioStreamInfoSet"`
}

type MediaPreprocessOperation struct {
	// 预处理操作的类型，取值范围：
	// <li>ImageTextMask：图片文字遮罩。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 预处理操作参数。
	// 当 Type 取值 ImageTextMask 时，参数为要保留的文字。
	Args []*string `json:"Args,omitempty" name:"Args"`
}

type MediaReplacementInfo struct {
	// 替换的媒体类型，取值有：
	// <li>CMEMaterialId：替换的媒体类型为媒体 ID；</li>
	// <li>ImageUrl：替换的媒体类型为图片 URL；</li>
	// 
	// 注：默认为 CMEMaterialId 。
	MediaType *string `json:"MediaType,omitempty" name:"MediaType"`

	// 媒体 ID。
	// 当媒体类型取值为 CMEMaterialId 时有效。
	MaterialId *string `json:"MaterialId,omitempty" name:"MaterialId"`

	// 媒体 URL。
	// 当媒体类型取值为 ImageUrl 时有效，
	// 图片仅支持 jpg、png 格式，且大小不超过 2M 。
	MediaUrl *string `json:"MediaUrl,omitempty" name:"MediaUrl"`

	// 替换媒体选取的开始时间，单位为秒，默认为 0。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 预处理操作。
	// 注：目前该功能暂不支持，请勿使用。
	PreprocessOperation *MediaPreprocessOperation `json:"PreprocessOperation,omitempty" name:"PreprocessOperation"`
}

type MediaTrack struct {
	// 轨道类型，取值有：
	// <ul>
	// <li>Video ：视频轨道。视频轨道由以下 Item 组成：<ul><li>VideoTrackItem</li><li>EmptyTrackItem</li><li>MediaTransitionItem</li></ul> </li>
	// <li>Audio ：音频轨道。音频轨道由以下 Item 组成：<ul><li>AudioTrackItem</li><li>EmptyTrackItem</li></ul> </li>
	// </ul>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 轨道上的媒体片段列表。
	TrackItems []*MediaTrackItem `json:"TrackItems,omitempty" name:"TrackItems"`
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

// Predefined struct for user
type ModifyMaterialRequestParams struct {
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 要修改的媒体 Id。
	MaterialId *string `json:"MaterialId,omitempty" name:"MaterialId"`

	// 媒体归属者。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 媒体名称，不能超过30个字符，不填则不修改。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 媒体分类路径，例如填写"/a/b"，则代表该媒体存储的路径为"/a/b"。若修改分类路径，则 Owner 字段必填。
	ClassPath *string `json:"ClassPath,omitempty" name:"ClassPath"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以修改任意媒体的信息。如果指定操作者，则操作者必须对媒体有写权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type ModifyMaterialRequest struct {
	*tchttp.BaseRequest
	
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 要修改的媒体 Id。
	MaterialId *string `json:"MaterialId,omitempty" name:"MaterialId"`

	// 媒体归属者。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 媒体名称，不能超过30个字符，不填则不修改。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 媒体分类路径，例如填写"/a/b"，则代表该媒体存储的路径为"/a/b"。若修改分类路径，则 Owner 字段必填。
	ClassPath *string `json:"ClassPath,omitempty" name:"ClassPath"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以修改任意媒体的信息。如果指定操作者，则操作者必须对媒体有写权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *ModifyMaterialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMaterialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMaterialResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyMaterialResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMaterialResponseParams `json:"Response"`
}

func (r *ModifyMaterialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMaterialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProjectRequestParams struct {
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 项目 Id。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 项目名称，不可超过30个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 画布宽高比，值为视频编辑项目画布宽与高的像素值的比值，如 16:9、9:16 等。
	AspectRatio *string `json:"AspectRatio,omitempty" name:"AspectRatio"`

	// 项目所有者。目前仅支持个人项目，不支持团队项目。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 项目模式，一个项目可以有多种模式并相互切换。
	// 当 Category 为 VIDEO_EDIT 时，可选模式有：
	// <li>Default：默认模式，即普通视频编辑项目。</li>
	// <li>VideoEditTemplate：剪辑模板制作模式，用于制作剪辑模板。</li>
	Mode *string `json:"Mode,omitempty" name:"Mode"`
}

type ModifyProjectRequest struct {
	*tchttp.BaseRequest
	
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 项目 Id。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 项目名称，不可超过30个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 画布宽高比，值为视频编辑项目画布宽与高的像素值的比值，如 16:9、9:16 等。
	AspectRatio *string `json:"AspectRatio,omitempty" name:"AspectRatio"`

	// 项目所有者。目前仅支持个人项目，不支持团队项目。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 项目模式，一个项目可以有多种模式并相互切换。
	// 当 Category 为 VIDEO_EDIT 时，可选模式有：
	// <li>Default：默认模式，即普通视频编辑项目。</li>
	// <li>VideoEditTemplate：剪辑模板制作模式，用于制作剪辑模板。</li>
	Mode *string `json:"Mode,omitempty" name:"Mode"`
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
	delete(f, "Platform")
	delete(f, "ProjectId")
	delete(f, "Name")
	delete(f, "AspectRatio")
	delete(f, "Owner")
	delete(f, "Mode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProjectResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

// Predefined struct for user
type ModifyTeamMemberRequestParams struct {
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 团队 ID。
	TeamId *string `json:"TeamId,omitempty" name:"TeamId"`

	// 团队成员 ID。
	MemberId *string `json:"MemberId,omitempty" name:"MemberId"`

	// 成员备注，长度不能超过15个字符。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 成员角色，可取值有：
	// <li>Admin：团队管理员；</li>
	// <li>Member：普通成员。</li>
	Role *string `json:"Role,omitempty" name:"Role"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以个改任意团队成员的信息。如果指定操作者，则操作者必须为团队的管理员或者所有者。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type ModifyTeamMemberRequest struct {
	*tchttp.BaseRequest
	
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 团队 ID。
	TeamId *string `json:"TeamId,omitempty" name:"TeamId"`

	// 团队成员 ID。
	MemberId *string `json:"MemberId,omitempty" name:"MemberId"`

	// 成员备注，长度不能超过15个字符。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 成员角色，可取值有：
	// <li>Admin：团队管理员；</li>
	// <li>Member：普通成员。</li>
	Role *string `json:"Role,omitempty" name:"Role"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以个改任意团队成员的信息。如果指定操作者，则操作者必须为团队的管理员或者所有者。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *ModifyTeamMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTeamMemberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTeamMemberResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyTeamMemberResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTeamMemberResponseParams `json:"Response"`
}

func (r *ModifyTeamMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTeamMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTeamRequestParams struct {
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 团队 ID。
	TeamId *string `json:"TeamId,omitempty" name:"TeamId"`

	// 团队名称。团队名称不能置空，并且不能超过30个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以修改所有团队的信息。如果指定操作者，则操作者必须为团队管理员或者所有者。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type ModifyTeamRequest struct {
	*tchttp.BaseRequest
	
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 团队 ID。
	TeamId *string `json:"TeamId,omitempty" name:"TeamId"`

	// 团队名称。团队名称不能置空，并且不能超过30个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以修改所有团队的信息。如果指定操作者，则操作者必须为团队管理员或者所有者。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *ModifyTeamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTeamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTeamResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyTeamResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTeamResponseParams `json:"Response"`
}

func (r *ModifyTeamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTeamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVideoEncodingPresetRequestParams struct {
	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 配置 ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 更改后的视频编码配置名，不填则不修改。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否去除视频数据，可选值：
	// <li>0：保留；</li>
	// <li>1：去除。</li>
	// 默认值：0。
	RemoveVideo *int64 `json:"RemoveVideo,omitempty" name:"RemoveVideo"`

	// 是否去除音频数据，可选值：
	// <li>0：保留；</li>
	// <li>1：去除。</li>
	// 默认值：0。
	RemoveAudio *int64 `json:"RemoveAudio,omitempty" name:"RemoveAudio"`

	// 更改后的编码配置的视频设置。
	VideoSetting *VideoEncodingPresetVideoSettingForUpdate `json:"VideoSetting,omitempty" name:"VideoSetting"`

	// 更改后的编码配置的音频设置。
	AudioSetting *VideoEncodingPresetAudioSettingForUpdate `json:"AudioSetting,omitempty" name:"AudioSetting"`
}

type ModifyVideoEncodingPresetRequest struct {
	*tchttp.BaseRequest
	
	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 配置 ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 更改后的视频编码配置名，不填则不修改。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否去除视频数据，可选值：
	// <li>0：保留；</li>
	// <li>1：去除。</li>
	// 默认值：0。
	RemoveVideo *int64 `json:"RemoveVideo,omitempty" name:"RemoveVideo"`

	// 是否去除音频数据，可选值：
	// <li>0：保留；</li>
	// <li>1：去除。</li>
	// 默认值：0。
	RemoveAudio *int64 `json:"RemoveAudio,omitempty" name:"RemoveAudio"`

	// 更改后的编码配置的视频设置。
	VideoSetting *VideoEncodingPresetVideoSettingForUpdate `json:"VideoSetting,omitempty" name:"VideoSetting"`

	// 更改后的编码配置的音频设置。
	AudioSetting *VideoEncodingPresetAudioSettingForUpdate `json:"AudioSetting,omitempty" name:"AudioSetting"`
}

func (r *ModifyVideoEncodingPresetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVideoEncodingPresetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "Id")
	delete(f, "Name")
	delete(f, "RemoveVideo")
	delete(f, "RemoveAudio")
	delete(f, "VideoSetting")
	delete(f, "AudioSetting")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVideoEncodingPresetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVideoEncodingPresetResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyVideoEncodingPresetResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVideoEncodingPresetResponseParams `json:"Response"`
}

func (r *ModifyVideoEncodingPresetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVideoEncodingPresetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MoveClassRequestParams struct {
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

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MoveClassRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MoveClassResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type MoveClassResponse struct {
	*tchttp.BaseResponse
	Response *MoveClassResponseParams `json:"Response"`
}

func (r *MoveClassResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MoveClassResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MoveResourceRequestParams struct {
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 待移动的原始资源信息，包含原始媒体或分类资源，以及资源归属。
	SourceResource *ResourceInfo `json:"SourceResource,omitempty" name:"SourceResource"`

	// 目标信息，包含分类及归属，仅支持移动资源到分类。
	DestinationResource *ResourceInfo `json:"DestinationResource,omitempty" name:"DestinationResource"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以移动任务资源。如果指定操作者，则操作者必须对源及目标资源有写权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type MoveResourceRequest struct {
	*tchttp.BaseRequest
	
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 待移动的原始资源信息，包含原始媒体或分类资源，以及资源归属。
	SourceResource *ResourceInfo `json:"SourceResource,omitempty" name:"SourceResource"`

	// 目标信息，包含分类及归属，仅支持移动资源到分类。
	DestinationResource *ResourceInfo `json:"DestinationResource,omitempty" name:"DestinationResource"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以移动任务资源。如果指定操作者，则操作者必须对源及目标资源有写权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *MoveResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MoveResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MoveResourceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type MoveResourceResponse struct {
	*tchttp.BaseResponse
	Response *MoveResourceResponseParams `json:"Response"`
}

func (r *MoveResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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

// Predefined struct for user
type ParseEventRequestParams struct {
	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 回调事件内容。
	EventContent *string `json:"EventContent,omitempty" name:"EventContent"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ParseEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "EventContent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ParseEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ParseEventResponseParams struct {
	// 事件内容。
	EventContent *EventContent `json:"EventContent,omitempty" name:"EventContent"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ParseEventResponse struct {
	*tchttp.BaseResponse
	Response *ParseEventResponseParams `json:"Response"`
}

func (r *ParseEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
	Tags []*string `json:"Tags,omitempty" name:"Tags"`

	// 视频分类，详见[企鹅号官网](https://open.om.qq.com/resources/resourcesCenter)视频分类。
	Category *int64 `json:"Category,omitempty" name:"Category"`
}

type PlatformInfo struct {
	// 平台标识。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 平台描述。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 云点播子应用 Id。
	VodSubAppId *uint64 `json:"VodSubAppId,omitempty" name:"VodSubAppId"`

	// 平台绑定的 license Id。
	LicenseId *string `json:"LicenseId,omitempty" name:"LicenseId"`

	// 平台状态，可取值为：
	// <li>Normal：正常，可使用。；</li>
	// <li>Stopped：已停用，暂无法使用；</li>
	// <li>Expired：已过期，需要重新购买会员包。</li>
	Status *string `json:"Status,omitempty" name:"Status"`

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

	// 点播转直播项目信息，仅当项目类别取值为 MEDIA_CAST 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MediaCastProjectInfo *MediaCastProjectInfo `json:"MediaCastProjectInfo,omitempty" name:"MediaCastProjectInfo"`

	// 项目更新时间，格式按照 ISO 8601 标准表示。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 项目创建时间，格式按照 ISO 8601 标准表示。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type ProjectMediaCastStatusChangedEvent struct {
	// 项目 Id。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 项目状态，取值有：
	// <li>Started：点播转直播开始；</li>
	// <li>Stopped：点播转直播结束；</li>
	// <li>SourceInterrupted：点播转直播输入断流；</li>
	// <li>DestinationInterrupted：点播转直播输出断流。</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 点播转直播输入断流信息，仅当 Status 取值 SourceInterrupted 时有效。
	SourceInterruptInfo *MediaCastSourceInterruptInfo `json:"SourceInterruptInfo,omitempty" name:"SourceInterruptInfo"`

	// 点播转直播输出断流信息，仅当 Status 取值 DestinationInterrupted 时有效。
	DestinationInterruptInfo *MediaCastDestinationInterruptInfo `json:"DestinationInterruptInfo,omitempty" name:"DestinationInterruptInfo"`
}

type ProjectStreamConnectStatusChangedEvent struct {
	// 项目 Id。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 项目状态，取值有：
	// <li>Working：云转推推流开始；</li>
	// <li>Stopped：云转推推流结束；</li>
	// <li>InputInterrupted：云转推输入断流；</li>
	// <li>OutputInterrupted：云转推输出断流。</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 云转推输入断流信息，仅当 Status 取值 InputInterrupted 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputInterruptInfo *StreamConnectInputInterruptInfo `json:"InputInterruptInfo,omitempty" name:"InputInterruptInfo"`

	// 云转推输出断流信息，仅当 Status 取值 OutputInterrupted 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputInterruptInfo *StreamConnectOutputInterruptInfo `json:"OutputInterruptInfo,omitempty" name:"OutputInterruptInfo"`
}

type ProjectSwitcherStatusChangedEvent struct {
	// 导播台项目 Id。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 导播台项目状态，可取值有：
	// <li>Started：导播台启动；</li>
	// <li>Stopped：导播台停止；</li>
	// <li>PvwStarted：导播台 PVW 开启；</li>
	// <li>PgmStarted：导播台 PGM 开启，输出推流开始；</li>
	// <li>PvwStopped：导播台 PVW 停止；</li>
	// <li>PgmStopped：导播台 PGM 停止，输出推流结束。</li>
	Status *string `json:"Status,omitempty" name:"Status"`
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

// Predefined struct for user
type RevokeResourceAuthorizationRequestParams struct {
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 资源所属实体。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 被授权资源。
	Resources []*Resource `json:"Resources,omitempty" name:"Resources"`

	// 被授权目标实体。
	Authorizees []*Entity `json:"Authorizees,omitempty" name:"Authorizees"`

	// 详细授权值。 取值有：
	// <li>R：可读，可以浏览素材，但不能使用该素材（将其添加到 Project），或复制到自己的媒资库中</li>
	// <li>X：可用，可以使用该素材（将其添加到 Project），但不能将其复制到自己的媒资库中，意味着被授权者无法将该资源进一步扩散给其他个人或团队。</li>
	// <li>C：可复制，既可以使用该素材（将其添加到 Project），也可以将其复制到自己的媒资库中。</li>
	// <li>W：可修改、删除媒资。</li>
	Permissions []*string `json:"Permissions,omitempty" name:"Permissions"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，撤销任意资源的授权权限。如果指定操作者，则操作者必须对被授权资源有写权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type RevokeResourceAuthorizationRequest struct {
	*tchttp.BaseRequest
	
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 资源所属实体。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 被授权资源。
	Resources []*Resource `json:"Resources,omitempty" name:"Resources"`

	// 被授权目标实体。
	Authorizees []*Entity `json:"Authorizees,omitempty" name:"Authorizees"`

	// 详细授权值。 取值有：
	// <li>R：可读，可以浏览素材，但不能使用该素材（将其添加到 Project），或复制到自己的媒资库中</li>
	// <li>X：可用，可以使用该素材（将其添加到 Project），但不能将其复制到自己的媒资库中，意味着被授权者无法将该资源进一步扩散给其他个人或团队。</li>
	// <li>C：可复制，既可以使用该素材（将其添加到 Project），也可以将其复制到自己的媒资库中。</li>
	// <li>W：可修改、删除媒资。</li>
	Permissions []*string `json:"Permissions,omitempty" name:"Permissions"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，撤销任意资源的授权权限。如果指定操作者，则操作者必须对被授权资源有写权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *RevokeResourceAuthorizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RevokeResourceAuthorizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RevokeResourceAuthorizationResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RevokeResourceAuthorizationResponse struct {
	*tchttp.BaseResponse
	Response *RevokeResourceAuthorizationResponseParams `json:"Response"`
}

func (r *RevokeResourceAuthorizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RevokeResourceAuthorizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RtmpPushInputInfo struct {
	// 直播推流地址有效期，单位：秒 。
	ExpiredSecond *uint64 `json:"ExpiredSecond,omitempty" name:"ExpiredSecond"`

	// 直播推流地址，入参不填默认由多媒体创作引擎生成。
	PushUrl *string `json:"PushUrl,omitempty" name:"PushUrl"`
}

// Predefined struct for user
type SearchMaterialRequestParams struct {
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 指定搜索空间，数组长度不得超过5。
	SearchScopes []*SearchScope `json:"SearchScopes,omitempty" name:"SearchScopes"`

	// 媒体类型，可取值有：
	// <li>AUDIO：音频；</li>
	// <li>VIDEO：视频 ；</li>
	// <li>IMAGE：图片；</li>
	// <li>VIDEO_EDIT_TEMPLATE：剪辑模板。</li>
	MaterialTypes []*string `json:"MaterialTypes,omitempty" name:"MaterialTypes"`

	// 搜索文本，模糊匹配媒体名称或描述信息，匹配项越多，匹配度越高，排序越优先。长度限制：15个字符。
	Text *string `json:"Text,omitempty" name:"Text"`

	// 按画质检索，取值为：LD/SD/HD/FHD/2K/4K。
	Resolution *string `json:"Resolution,omitempty" name:"Resolution"`

	// 按媒体时长检索，单位s。
	DurationRange *IntegerRange `json:"DurationRange,omitempty" name:"DurationRange"`

	// 按照媒体创建时间检索。
	CreateTimeRange *TimeRange `json:"CreateTimeRange,omitempty" name:"CreateTimeRange"`

	// 按标签检索，填入检索的标签名。
	Tags []*string `json:"Tags,omitempty" name:"Tags"`

	// 排序方式。Sort.Field 可选值：CreateTime。指定 Text 搜索时，将根据匹配度排序，该字段无效。
	Sort *SortBy `json:"Sort,omitempty" name:"Sort"`

	// 偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回记录条数，默认值：50。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以搜索任意媒体的信息。如果指定操作者，则操作者必须对媒体有读权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type SearchMaterialRequest struct {
	*tchttp.BaseRequest
	
	// 平台 Id，指定访问的平台。关于平台概念，请参见文档 [平台](https://cloud.tencent.com/document/product/1156/43767)。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 指定搜索空间，数组长度不得超过5。
	SearchScopes []*SearchScope `json:"SearchScopes,omitempty" name:"SearchScopes"`

	// 媒体类型，可取值有：
	// <li>AUDIO：音频；</li>
	// <li>VIDEO：视频 ；</li>
	// <li>IMAGE：图片；</li>
	// <li>VIDEO_EDIT_TEMPLATE：剪辑模板。</li>
	MaterialTypes []*string `json:"MaterialTypes,omitempty" name:"MaterialTypes"`

	// 搜索文本，模糊匹配媒体名称或描述信息，匹配项越多，匹配度越高，排序越优先。长度限制：15个字符。
	Text *string `json:"Text,omitempty" name:"Text"`

	// 按画质检索，取值为：LD/SD/HD/FHD/2K/4K。
	Resolution *string `json:"Resolution,omitempty" name:"Resolution"`

	// 按媒体时长检索，单位s。
	DurationRange *IntegerRange `json:"DurationRange,omitempty" name:"DurationRange"`

	// 按照媒体创建时间检索。
	CreateTimeRange *TimeRange `json:"CreateTimeRange,omitempty" name:"CreateTimeRange"`

	// 按标签检索，填入检索的标签名。
	Tags []*string `json:"Tags,omitempty" name:"Tags"`

	// 排序方式。Sort.Field 可选值：CreateTime。指定 Text 搜索时，将根据匹配度排序，该字段无效。
	Sort *SortBy `json:"Sort,omitempty" name:"Sort"`

	// 偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回记录条数，默认值：50。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 操作者。如不填，默认为 `cmeid_system`，表示平台管理员操作，可以搜索任意媒体的信息。如果指定操作者，则操作者必须对媒体有读权限。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *SearchMaterialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchMaterialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchMaterialResponseParams struct {
	// 符合记录总条数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 媒体信息，仅返回基础信息。
	MaterialInfoSet []*MaterialInfo `json:"MaterialInfoSet,omitempty" name:"MaterialInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SearchMaterialResponse struct {
	*tchttp.BaseResponse
	Response *SearchMaterialResponseParams `json:"Response"`
}

func (r *SearchMaterialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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

	// 卡槽类型，可取值有：
	// <li> AUDIO：音频卡槽，可替换素材类型为 AUDIO 的音频素材;</li>
	// <li> VIDEO：视频卡槽，可替换素材类型为 VIDEO 的视频素材;</li>
	// <li> IMAGE：图片卡槽，可替换素材类型为 IMAGE 的图片素材;</li>
	// <li> TEXT：文本卡槽，可替换文本内容。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 默认素材ID。当卡槽类型为 AUDIO，VIDEO，或 IMAGE 中的一种时有效。
	DefaultMaterialId *string `json:"DefaultMaterialId,omitempty" name:"DefaultMaterialId"`

	// 默认文本卡槽信息。当卡槽类型为 TEXT 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultTextSlotInfo *TextSlotInfo `json:"DefaultTextSlotInfo,omitempty" name:"DefaultTextSlotInfo"`

	// 素材时长，单位秒。
	Duration *float64 `json:"Duration,omitempty" name:"Duration"`
}

type SlotReplacementInfo struct {
	// 卡槽 Id。
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 替换类型，可取值有：
	// <li> AUDIO ：音频；</li>
	// <li> VIDEO ：视频；</li>
	// <li> IMAGE ：图片；</li>
	// <li> TEXT ：文本。</li>
	// 注意：这里必须保证替换的素材类型与模板轨道数据的素材类型一致。如果替换的类型为Text,，则必须保证模板轨道数据中相应卡槽的位置标记的是文本。
	ReplacementType *string `json:"ReplacementType,omitempty" name:"ReplacementType"`

	// 媒体替换信息，仅当要替换的媒体类型为音频、视频、图片时有效。
	MediaReplacementInfo *MediaReplacementInfo `json:"MediaReplacementInfo,omitempty" name:"MediaReplacementInfo"`

	// 文本替换信息，仅当要替换的卡槽类型为文本时有效。
	TextReplacementInfo *TextReplacementInfo `json:"TextReplacementInfo,omitempty" name:"TextReplacementInfo"`
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

	// 操作者 Id。（废弃，请勿使用）
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 操作类型，可取值有：
	// <li>Upload：本地上传；</li>
	// <li>PullUpload：拉取上传；</li>
	// <li>VideoEdit：视频剪辑；</li>
	// <li>LiveStreamClip：直播流剪辑；</li>
	// <li>LiveStreamRecord：直播流录制。</li>
	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`

	// 媒体归属。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 媒体分类路径。
	ClassPath *string `json:"ClassPath,omitempty" name:"ClassPath"`

	// 生成文件的任务 Id。当生成新文件是拉取上传、视频剪辑、直播流剪辑时为任务 Id。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 来源上下文信息。视频剪辑生成新文件时此字段为项目 Id；直播流剪辑或者直播流录制生成新文件则为原始流地址。
	SourceContext *string `json:"SourceContext,omitempty" name:"SourceContext"`
}

type StreamConnectInputInterruptInfo struct {
	// 云转推输入源标识，取值有：
	// <li>Main：主源；</li>
	// <li>Backup：备源。</li>
	EndPoint *string `json:"EndPoint,omitempty" name:"EndPoint"`
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

type StreamConnectOutputInterruptInfo struct {
	// 云转推输出标识。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 云转推输出名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 云转推输出地址。
	Url *string `json:"Url,omitempty" name:"Url"`
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
	OutputSet []*StreamConnectOutputInfo `json:"OutputSet,omitempty" name:"OutputSet"`
}

type StreamConnectProjectInput struct {
	// 云转推主输入源信息。
	MainInput *StreamInputInfo `json:"MainInput,omitempty" name:"MainInput"`

	// 云转推备输入源信息。
	BackupInput *StreamInputInfo `json:"BackupInput,omitempty" name:"BackupInput"`

	// 云转推输出源信息。
	Outputs []*StreamConnectOutput `json:"Outputs,omitempty" name:"Outputs"`
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

type TextReplacementInfo struct {
	// 替换的文本信息。
	Text *string `json:"Text,omitempty" name:"Text"`
}

type TextSlotInfo struct {
	// 文本内容。
	Text *string `json:"Text,omitempty" name:"Text"`
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

	// 腾讯云对象存储发布信息， 如果使用的发布通道为腾讯云对象存储时必填。
	CosPublishInfo *CosPublishInputInfo `json:"CosPublishInfo,omitempty" name:"CosPublishInfo"`
}

type TimeRange struct {
	// 开始时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#I)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type VODExportInfo struct {
	// 导出的媒资名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 导出的媒资分类 Id。
	ClassId *uint64 `json:"ClassId,omitempty" name:"ClassId"`

	// 第三方平台发布信息列表。暂未正式对外，请勿使用。
	ThirdPartyPublishInfos []*ThirdPartyPublishInfo `json:"ThirdPartyPublishInfos,omitempty" name:"ThirdPartyPublishInfos"`
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
	InitTracks []*MediaTrack `json:"InitTracks,omitempty" name:"InitTracks"`
}

type VideoEditProjectOutput struct {
	// 导出的多媒体创作引擎媒体 Id，仅当导出目标为多媒体创作引擎媒体时有效。
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
	SlotSet []*SlotInfo `json:"SlotSet,omitempty" name:"SlotSet"`

	// 模板预览视频 URL 地址 。
	PreviewVideoUrl *string `json:"PreviewVideoUrl,omitempty" name:"PreviewVideoUrl"`
}

type VideoEncodingPreset struct {
	// 配置 ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 配置名。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 封装格式，可选值：
	// <li>mp4 ；</li>
	// <li>mov 。</li>
	Container *string `json:"Container,omitempty" name:"Container"`

	// 是否去除视频数据，可选值：
	// <li>0：保留；</li>
	// <li>1：去除。</li>
	// 默认值：0。
	RemoveVideo *int64 `json:"RemoveVideo,omitempty" name:"RemoveVideo"`

	// 是否去除音频数据，可选值：
	// <li>0：保留；</li>
	// <li>1：去除。</li>
	// 默认值：0。
	RemoveAudio *int64 `json:"RemoveAudio,omitempty" name:"RemoveAudio"`

	// 视频编码配置中的视频设置。
	VideoSetting *VideoEncodingPresetVideoSetting `json:"VideoSetting,omitempty" name:"VideoSetting"`

	// 视频编码配置中的音频设置。
	AudioSetting *VideoEncodingPresetAudioSetting `json:"AudioSetting,omitempty" name:"AudioSetting"`
}

type VideoEncodingPresetAudioSetting struct {
	// 音频流的编码格式，可选值：
	// AAC：AAC 编码。
	// 
	// 默认值：AAC。
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// 音频码率，单位：bps。
	// 默认值：64K。
	Bitrate *uint64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 音频声道数，可选值： 
	// <li>1：单声道；</li>
	// <li>2：双声道。</li> 
	// 默认值：2。
	Channels *uint64 `json:"Channels,omitempty" name:"Channels"`

	// 音频流的采样率，仅支持 16000； 32000； 44100； 48000。单位：Hz。 
	// 默认值：16000。
	SampleRate *uint64 `json:"SampleRate,omitempty" name:"SampleRate"`
}

type VideoEncodingPresetAudioSettingForUpdate struct {
	// 音频码率，单位：bps。
	// 不填则不修改。
	Bitrate *string `json:"Bitrate,omitempty" name:"Bitrate"`

	// 音频声道数，可选值： 
	// <li>1：单声道；</li>
	// <li>2：双声道。</li> 
	// 不填则不修改。
	Channels *uint64 `json:"Channels,omitempty" name:"Channels"`

	// 音频流的采样率，目前仅支持： 16000； 32000； 44100； 48000。单位：Hz。
	// 不填则不修改。
	SampleRate *uint64 `json:"SampleRate,omitempty" name:"SampleRate"`
}

type VideoEncodingPresetVideoSetting struct {
	// 视频流的编码格式，可选值：
	// <li>H264：H.264 编码。</li>
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// 视频短边尺寸，取值范围： [128, 4096]，单位：px。
	// 视频最后的分辨率，根据短边尺寸和宽高比进行计算。
	// 例：如果项目的宽高比是 16：9 ：
	// <li>短边尺寸为 1080，则导出视频的分辨率为 1920 * 1080。</li>
	// <li>短边尺寸为 720，则导出视频的分辨率为 1280 * 720。</li>
	// 如果项目的宽高比是 9：16 ：
	// <li>短边尺寸为 1080，则导出视频的分辨率为 1080 * 1920。</li>
	// <li>短边尺寸为 720，则导出视频的分辨率为 720 * 1280。</li>
	// 默认值：1080。
	ShortEdge *uint64 `json:"ShortEdge,omitempty" name:"ShortEdge"`

	// 指定码率，单位 bps。当该参数为'0'时则不强制限定码率。
	// 默认值：0。
	Bitrate *uint64 `json:"Bitrate,omitempty" name:"Bitrate"`
}

type VideoEncodingPresetVideoSettingForUpdate struct {
	// 视频短边尺寸，取值范围： [128, 4096]，单位：px。
	// 视频最后的分辨率，根据短边尺寸和宽高比进行计算。
	// 例：如果项目的宽高比是 16：9 ：
	// <li>短边尺寸为 1080，则导出视频的分辨率为 1920 * 1080。</li>
	// <li>短边尺寸为 720，则导出视频的分辨率为 1280 * 720。</li>
	// 如果项目的宽高比是 9：16 ：
	// <li>短边尺寸为 1080，则导出视频的分辨率为 1080 * 1920。</li>
	// <li>短边尺寸为 720，则导出视频的分辨率为 720 * 1280。</li>
	// 不填则不修改。
	ShortEdge *uint64 `json:"ShortEdge,omitempty" name:"ShortEdge"`

	// 指定码率，单位 bps。当该参数为'0' 时则不强制限定码率。
	// 不填则不修改。
	Bitrate *uint64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 指定帧率。单位 Hz。
	// 不填则不修改。
	FrameRate *float64 `json:"FrameRate,omitempty" name:"FrameRate"`
}

type VideoExportCompletedEvent struct {
	// 任务 Id。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务状态，取值有：
	// <li>SUCCESS：成功；</li>
	// <li>FAIL：失败。</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，取值有：
	// <li>0：成功；</li>
	// <li>其他值：失败。</li>
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *VideoEditProjectOutput `json:"Output,omitempty" name:"Output"`
}

type VideoExportExtensionArgs struct {
	// 封装格式，可选值：
	// <li>mp4 </li>
	// <li>mov </li>
	// 不填则使用视频导出编码配置。
	Container *string `json:"Container,omitempty" name:"Container"`

	// 视频短边尺寸，取值范围： [128, 4096]，单位：px。
	// 视频最后的分辨率，根据短边尺寸和宽高比进行计算。
	// 例如：项目的宽高比是 16：9 ：
	// <li>短边尺寸为 1080，则导出视频的分辨率为 1920 * 1080。</li>
	// <li>短边尺寸为 720，则导出视频的分辨率为 1280 * 720</li>
	// 不填则使用视频导出编码配置。
	ShortEdge *uint64 `json:"ShortEdge,omitempty" name:"ShortEdge"`

	// 指定码率，单位 bps。当该参数为 0 时则不强制限定码率。
	// 不填则使用视频导出编码配置。
	VideoBitrate *uint64 `json:"VideoBitrate,omitempty" name:"VideoBitrate"`

	// 帧率。取值范围：[15, 60]，不填默认值为 25。
	FrameRate *float64 `json:"FrameRate,omitempty" name:"FrameRate"`

	// 是否去除视频数据，可选值：
	// <li>0：保留；</li>
	// <li>1：去除。</li>
	// 不填则使用视频导出编码配置。
	RemoveVideo *int64 `json:"RemoveVideo,omitempty" name:"RemoveVideo"`

	// 是否去除音频数据，可选值：
	// <li>0：保留；</li>
	// <li>1：去除。</li>
	// 不填则使用视频导出编码配置。
	RemoveAudio *int64 `json:"RemoveAudio,omitempty" name:"RemoveAudio"`

	// 片段起始时间，单位：毫秒。
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 片段结束时间，单位：毫秒。
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
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
	// <li>EXTERNAL ：视频来源于媒资绑定，如果媒体不是存储在腾讯云点播中或者云创中，都需要使用媒资绑定。</li>
	// </ul>
	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`

	// 视频媒体，可取值为：
	// <ul>
	// <li>当 SourceType 为 VOD 时，参数填云点播 FileId ；</li>
	// <li>当 SourceType 为 CME 时，参数填多媒体创作引擎媒体 Id；</li>
	// <li>当 SourceType 为 EXTERNAL 时，目前仅支持外部媒体 URL(如`https://www.example.com/a.mp4`)，参数填写规则请参见注意事项。</li>
	// </ul>
	// 
	// 注意：
	// <li>当 SourceType 为 EXTERNAL 并且媒体 URL Scheme 为 `https` 时(如：`https://www.example.com/a.mp4`)，参数为：`1000000:www.example.com/a.mp4`。</li>
	// <li>当 SourceType 为 EXTERNAL 并且媒体 URL Scheme 为 `http` 时(如：`http://www.example.com/b.mp4`)，参数为：`1000001:www.example.com/b.mp4`。</li>
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
	InputUrls []*string `json:"InputUrls,omitempty" name:"InputUrls"`

	// 播放次数，取值有：
	// <li>-1 : 循环播放，直到转推结束；</li>
	// <li>0 : 不循环；</li>
	// <li>大于0 : 具体循环次数，次数和时间以先结束的为准。</li>
	// 默认不循环。
	LoopTimes *int64 `json:"LoopTimes,omitempty" name:"LoopTimes"`
}

type VodPullInputPlayInfo struct {
	// 当前正在播放文件 Url 。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 点播文件已播放时长，单位：秒。
	TimeOffset *float64 `json:"TimeOffset,omitempty" name:"TimeOffset"`
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