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

package v20220817

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type AddGroupMemberRequestParams struct {
	// 群组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 成员列表，最大值200
	MemberIds []*string `json:"MemberIds,omitempty" name:"MemberIds"`
}

type AddGroupMemberRequest struct {
	*tchttp.BaseRequest
	
	// 群组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 成员列表，最大值200
	MemberIds []*string `json:"MemberIds,omitempty" name:"MemberIds"`
}

func (r *AddGroupMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddGroupMemberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "SdkAppId")
	delete(f, "MemberIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddGroupMemberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddGroupMemberResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddGroupMemberResponse struct {
	*tchttp.BaseResponse
	Response *AddGroupMemberResponseParams `json:"Response"`
}

func (r *AddGroupMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddGroupMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AppConfig struct {

}

type AppCustomContent struct {
	// 场景参数，一个应用下可以设置多个不同场景。
	Scene *string `json:"Scene,omitempty" name:"Scene"`

	// logo地址。
	LogoUrl *string `json:"LogoUrl,omitempty" name:"LogoUrl"`

	// 主页地址，可设置用于跳转。
	HomeUrl *string `json:"HomeUrl,omitempty" name:"HomeUrl"`

	// 自定义的js。
	JsUrl *string `json:"JsUrl,omitempty" name:"JsUrl"`

	// 自定义的css。
	CssUrl *string `json:"CssUrl,omitempty" name:"CssUrl"`
}

type BackgroundPictureConfig struct {
	// 背景图片的url
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitempty" name:"Url"`
}

// Predefined struct for user
type BatchAddGroupMemberRequestParams struct {
	// 待添加群组ID列表，最大值100
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 待添加成员列表，最大值200
	MemberIds []*string `json:"MemberIds,omitempty" name:"MemberIds"`
}

type BatchAddGroupMemberRequest struct {
	*tchttp.BaseRequest
	
	// 待添加群组ID列表，最大值100
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 待添加成员列表，最大值200
	MemberIds []*string `json:"MemberIds,omitempty" name:"MemberIds"`
}

func (r *BatchAddGroupMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchAddGroupMemberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupIds")
	delete(f, "SdkAppId")
	delete(f, "MemberIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchAddGroupMemberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchAddGroupMemberResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BatchAddGroupMemberResponse struct {
	*tchttp.BaseResponse
	Response *BatchAddGroupMemberResponseParams `json:"Response"`
}

func (r *BatchAddGroupMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchAddGroupMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchCreateGroupWithMembersRequestParams struct {
	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 批量创建群组基础信息，最大长度限制256
	GroupBaseInfos []*GroupBaseInfo `json:"GroupBaseInfos,omitempty" name:"GroupBaseInfos"`

	// 群组绑定的成员列表，一次性最多200个
	MemberIds []*string `json:"MemberIds,omitempty" name:"MemberIds"`
}

type BatchCreateGroupWithMembersRequest struct {
	*tchttp.BaseRequest
	
	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 批量创建群组基础信息，最大长度限制256
	GroupBaseInfos []*GroupBaseInfo `json:"GroupBaseInfos,omitempty" name:"GroupBaseInfos"`

	// 群组绑定的成员列表，一次性最多200个
	MemberIds []*string `json:"MemberIds,omitempty" name:"MemberIds"`
}

func (r *BatchCreateGroupWithMembersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchCreateGroupWithMembersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "GroupBaseInfos")
	delete(f, "MemberIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchCreateGroupWithMembersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchCreateGroupWithMembersResponseParams struct {
	// 新创建群组ID列表，与输入创建参数顺序一致
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BatchCreateGroupWithMembersResponse struct {
	*tchttp.BaseResponse
	Response *BatchCreateGroupWithMembersResponseParams `json:"Response"`
}

func (r *BatchCreateGroupWithMembersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchCreateGroupWithMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchCreateRoomRequestParams struct {
	// 低代码平台的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 创建房间ID列表
	RoomInfos []*RoomInfo `json:"RoomInfos,omitempty" name:"RoomInfos"`
}

type BatchCreateRoomRequest struct {
	*tchttp.BaseRequest
	
	// 低代码平台的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 创建房间ID列表
	RoomInfos []*RoomInfo `json:"RoomInfos,omitempty" name:"RoomInfos"`
}

func (r *BatchCreateRoomRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchCreateRoomRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomInfos")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchCreateRoomRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchCreateRoomResponseParams struct {
	// 创建成功课堂ID，与传入课堂信息顺序一致
	RoomIds []*uint64 `json:"RoomIds,omitempty" name:"RoomIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BatchCreateRoomResponse struct {
	*tchttp.BaseResponse
	Response *BatchCreateRoomResponseParams `json:"Response"`
}

func (r *BatchCreateRoomResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchCreateRoomResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchDeleteGroupMemberRequestParams struct {
	// 待添加群组ID列表，最大值100
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 待添加成员列表，最大值256
	MemberIds []*string `json:"MemberIds,omitempty" name:"MemberIds"`
}

type BatchDeleteGroupMemberRequest struct {
	*tchttp.BaseRequest
	
	// 待添加群组ID列表，最大值100
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 待添加成员列表，最大值256
	MemberIds []*string `json:"MemberIds,omitempty" name:"MemberIds"`
}

func (r *BatchDeleteGroupMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDeleteGroupMemberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupIds")
	delete(f, "SdkAppId")
	delete(f, "MemberIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchDeleteGroupMemberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchDeleteGroupMemberResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BatchDeleteGroupMemberResponse struct {
	*tchttp.BaseResponse
	Response *BatchDeleteGroupMemberResponseParams `json:"Response"`
}

func (r *BatchDeleteGroupMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDeleteGroupMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchDeleteRecordRequestParams struct {
	// 房间ID列表
	RoomIds []*int64 `json:"RoomIds,omitempty" name:"RoomIds"`

	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

type BatchDeleteRecordRequest struct {
	*tchttp.BaseRequest
	
	// 房间ID列表
	RoomIds []*int64 `json:"RoomIds,omitempty" name:"RoomIds"`

	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

func (r *BatchDeleteRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDeleteRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoomIds")
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchDeleteRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchDeleteRecordResponseParams struct {
	// 本次操作删除成功的房间ID列表。如果入参列表中某个房间ID的录制文件已经删除，则出参列表中无对应的房间ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoomIds []*int64 `json:"RoomIds,omitempty" name:"RoomIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BatchDeleteRecordResponse struct {
	*tchttp.BaseResponse
	Response *BatchDeleteRecordResponseParams `json:"Response"`
}

func (r *BatchDeleteRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDeleteRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchRegisterRequestParams struct {
	// 批量注册用户信息列表
	Users []*BatchUserRequest `json:"Users,omitempty" name:"Users"`
}

type BatchRegisterRequest struct {
	*tchttp.BaseRequest
	
	// 批量注册用户信息列表
	Users []*BatchUserRequest `json:"Users,omitempty" name:"Users"`
}

func (r *BatchRegisterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchRegisterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Users")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchRegisterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchRegisterResponseParams struct {
	// 注册成功的用户列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Users []*BatchUserInfo `json:"Users,omitempty" name:"Users"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BatchRegisterResponse struct {
	*tchttp.BaseResponse
	Response *BatchRegisterResponseParams `json:"Response"`
}

func (r *BatchRegisterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchRegisterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchUserInfo struct {
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 用户ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户在客户系统的Id。 若用户注册时该字段为空，则默认为 UserId 值一致。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`
}

type BatchUserRequest struct {
	// 低代码互动课堂的SdkAppId。
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 用户名称。
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 用户在客户系统的Id，需要在同一应用下唯一。
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`

	// 用户头像。
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Avatar *string `json:"Avatar,omitempty" name:"Avatar"`
}

// Predefined struct for user
type BindDocumentToRoomRequestParams struct {
	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// 文档ID。
	DocumentId *string `json:"DocumentId,omitempty" name:"DocumentId"`

	// 绑定类型。后台可透传到客户端，默认为0。客户端可以根据这个字段实现业务逻辑。
	BindType *uint64 `json:"BindType,omitempty" name:"BindType"`
}

type BindDocumentToRoomRequest struct {
	*tchttp.BaseRequest
	
	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// 文档ID。
	DocumentId *string `json:"DocumentId,omitempty" name:"DocumentId"`

	// 绑定类型。后台可透传到客户端，默认为0。客户端可以根据这个字段实现业务逻辑。
	BindType *uint64 `json:"BindType,omitempty" name:"BindType"`
}

func (r *BindDocumentToRoomRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindDocumentToRoomRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoomId")
	delete(f, "DocumentId")
	delete(f, "BindType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindDocumentToRoomRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindDocumentToRoomResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BindDocumentToRoomResponse struct {
	*tchttp.BaseResponse
	Response *BindDocumentToRoomResponseParams `json:"Response"`
}

func (r *BindDocumentToRoomResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindDocumentToRoomResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDocumentRequestParams struct {
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 文档地址。
	DocumentUrl *string `json:"DocumentUrl,omitempty" name:"DocumentUrl"`

	// 文档名称。
	DocumentName *string `json:"DocumentName,omitempty" name:"DocumentName"`

	// 文档所有者的Id
	Owner *string `json:"Owner,omitempty" name:"Owner"`

	// 转码类型，可以有如下取值：
	// 0 无需转码（默认）
	// 1 需要转码的文档，ppt，pptx，pdf，doc，docx
	// 2 需要转码的视频，mp4，3pg，mpeg，avi，flv，wmv，rm，h264等
	// 2 需要转码的音频，mp3，wav，wma，aac，flac，opus
	TranscodeType *uint64 `json:"TranscodeType,omitempty" name:"TranscodeType"`

	// 权限，可以有如下取值：
	// 0 私有文档（默认）
	// 1 公共文档
	Permission *uint64 `json:"Permission,omitempty" name:"Permission"`

	// 文档后缀名。
	DocumentType *string `json:"DocumentType,omitempty" name:"DocumentType"`

	// 文档大小，单位 字节
	DocumentSize *uint64 `json:"DocumentSize,omitempty" name:"DocumentSize"`
}

type CreateDocumentRequest struct {
	*tchttp.BaseRequest
	
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 文档地址。
	DocumentUrl *string `json:"DocumentUrl,omitempty" name:"DocumentUrl"`

	// 文档名称。
	DocumentName *string `json:"DocumentName,omitempty" name:"DocumentName"`

	// 文档所有者的Id
	Owner *string `json:"Owner,omitempty" name:"Owner"`

	// 转码类型，可以有如下取值：
	// 0 无需转码（默认）
	// 1 需要转码的文档，ppt，pptx，pdf，doc，docx
	// 2 需要转码的视频，mp4，3pg，mpeg，avi，flv，wmv，rm，h264等
	// 2 需要转码的音频，mp3，wav，wma，aac，flac，opus
	TranscodeType *uint64 `json:"TranscodeType,omitempty" name:"TranscodeType"`

	// 权限，可以有如下取值：
	// 0 私有文档（默认）
	// 1 公共文档
	Permission *uint64 `json:"Permission,omitempty" name:"Permission"`

	// 文档后缀名。
	DocumentType *string `json:"DocumentType,omitempty" name:"DocumentType"`

	// 文档大小，单位 字节
	DocumentSize *uint64 `json:"DocumentSize,omitempty" name:"DocumentSize"`
}

func (r *CreateDocumentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDocumentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "DocumentUrl")
	delete(f, "DocumentName")
	delete(f, "Owner")
	delete(f, "TranscodeType")
	delete(f, "Permission")
	delete(f, "DocumentType")
	delete(f, "DocumentSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDocumentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDocumentResponseParams struct {
	// 文档ID。
	DocumentId *string `json:"DocumentId,omitempty" name:"DocumentId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDocumentResponse struct {
	*tchttp.BaseResponse
	Response *CreateDocumentResponseParams `json:"Response"`
}

func (r *CreateDocumentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDocumentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGroupWithMembersRequestParams struct {
	// 待创建群组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 默认绑定主讲老师ID
	TeacherId *string `json:"TeacherId,omitempty" name:"TeacherId"`

	// 群组成员列表,一次性最多200个
	MemberIds []*string `json:"MemberIds,omitempty" name:"MemberIds"`
}

type CreateGroupWithMembersRequest struct {
	*tchttp.BaseRequest
	
	// 待创建群组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 默认绑定主讲老师ID
	TeacherId *string `json:"TeacherId,omitempty" name:"TeacherId"`

	// 群组成员列表,一次性最多200个
	MemberIds []*string `json:"MemberIds,omitempty" name:"MemberIds"`
}

func (r *CreateGroupWithMembersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGroupWithMembersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupName")
	delete(f, "SdkAppId")
	delete(f, "TeacherId")
	delete(f, "MemberIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateGroupWithMembersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGroupWithMembersResponseParams struct {
	// 创建成功群组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateGroupWithMembersResponse struct {
	*tchttp.BaseResponse
	Response *CreateGroupWithMembersResponseParams `json:"Response"`
}

func (r *CreateGroupWithMembersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGroupWithMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGroupWithSubGroupRequestParams struct {
	// 待创建的新群组名
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 子群组ID列表，子群组ID不能重复，最多40个
	SubGroupIds []*string `json:"SubGroupIds,omitempty" name:"SubGroupIds"`

	// 群组默认主讲老师ID
	TeacherId *string `json:"TeacherId,omitempty" name:"TeacherId"`
}

type CreateGroupWithSubGroupRequest struct {
	*tchttp.BaseRequest
	
	// 待创建的新群组名
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 子群组ID列表，子群组ID不能重复，最多40个
	SubGroupIds []*string `json:"SubGroupIds,omitempty" name:"SubGroupIds"`

	// 群组默认主讲老师ID
	TeacherId *string `json:"TeacherId,omitempty" name:"TeacherId"`
}

func (r *CreateGroupWithSubGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGroupWithSubGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupName")
	delete(f, "SdkAppId")
	delete(f, "SubGroupIds")
	delete(f, "TeacherId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateGroupWithSubGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGroupWithSubGroupResponseParams struct {
	// 新创建群组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateGroupWithSubGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateGroupWithSubGroupResponseParams `json:"Response"`
}

func (r *CreateGroupWithSubGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGroupWithSubGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRoomRequestParams struct {
	// 房间名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 预定的房间开始时间，unix时间戳。
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 预定的房间结束时间，unix时间戳。
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 分辨率。可以有如下取值：
	// 1 标清
	// 2 高清
	// 3 全高清
	Resolution *uint64 `json:"Resolution,omitempty" name:"Resolution"`

	// 最大连麦人数（不包括老师）。取值范围[0, 16]
	MaxMicNumber *uint64 `json:"MaxMicNumber,omitempty" name:"MaxMicNumber"`

	// 房间子类型，可以有以下取值：
	// videodoc 文档+视频
	// video 纯视频
	SubType *string `json:"SubType,omitempty" name:"SubType"`

	// 老师ID。通过[注册用户]接口获取的UserId。指定后该用户在房间内拥有老师权限。
	TeacherId *string `json:"TeacherId,omitempty" name:"TeacherId"`

	// 进入课堂时是否自动连麦。可以有以下取值：
	// 0 不自动连麦（需要手动申请上麦，默认值）
	// 1 自动连麦
	AutoMic *uint64 `json:"AutoMic,omitempty" name:"AutoMic"`

	// 高音质模式。可以有以下取值：
	// 0 不开启高音质（默认值）
	// 1 开启高音质
	AudioQuality *uint64 `json:"AudioQuality,omitempty" name:"AudioQuality"`

	// 上课后是否禁止自动录制。可以有以下取值：
	// 0 不禁止录制（自动开启录制，默认值）
	// 1 禁止录制
	// 注：如果该配置取值为0，录制将从上课后开始，课堂结束后停止。
	DisableRecord *uint64 `json:"DisableRecord,omitempty" name:"DisableRecord"`

	// 助教Id列表。通过[注册用户]接口获取的UserId。指定后该用户在房间内拥有助教权限。
	Assistants []*string `json:"Assistants,omitempty" name:"Assistants"`

	// 录制布局。
	RecordLayout *uint64 `json:"RecordLayout,omitempty" name:"RecordLayout"`

	// 房间绑定的群组ID,非空时限制组成员进入
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type CreateRoomRequest struct {
	*tchttp.BaseRequest
	
	// 房间名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 预定的房间开始时间，unix时间戳。
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 预定的房间结束时间，unix时间戳。
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 分辨率。可以有如下取值：
	// 1 标清
	// 2 高清
	// 3 全高清
	Resolution *uint64 `json:"Resolution,omitempty" name:"Resolution"`

	// 最大连麦人数（不包括老师）。取值范围[0, 16]
	MaxMicNumber *uint64 `json:"MaxMicNumber,omitempty" name:"MaxMicNumber"`

	// 房间子类型，可以有以下取值：
	// videodoc 文档+视频
	// video 纯视频
	SubType *string `json:"SubType,omitempty" name:"SubType"`

	// 老师ID。通过[注册用户]接口获取的UserId。指定后该用户在房间内拥有老师权限。
	TeacherId *string `json:"TeacherId,omitempty" name:"TeacherId"`

	// 进入课堂时是否自动连麦。可以有以下取值：
	// 0 不自动连麦（需要手动申请上麦，默认值）
	// 1 自动连麦
	AutoMic *uint64 `json:"AutoMic,omitempty" name:"AutoMic"`

	// 高音质模式。可以有以下取值：
	// 0 不开启高音质（默认值）
	// 1 开启高音质
	AudioQuality *uint64 `json:"AudioQuality,omitempty" name:"AudioQuality"`

	// 上课后是否禁止自动录制。可以有以下取值：
	// 0 不禁止录制（自动开启录制，默认值）
	// 1 禁止录制
	// 注：如果该配置取值为0，录制将从上课后开始，课堂结束后停止。
	DisableRecord *uint64 `json:"DisableRecord,omitempty" name:"DisableRecord"`

	// 助教Id列表。通过[注册用户]接口获取的UserId。指定后该用户在房间内拥有助教权限。
	Assistants []*string `json:"Assistants,omitempty" name:"Assistants"`

	// 录制布局。
	RecordLayout *uint64 `json:"RecordLayout,omitempty" name:"RecordLayout"`

	// 房间绑定的群组ID,非空时限制组成员进入
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *CreateRoomRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRoomRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SdkAppId")
	delete(f, "Resolution")
	delete(f, "MaxMicNumber")
	delete(f, "SubType")
	delete(f, "TeacherId")
	delete(f, "AutoMic")
	delete(f, "AudioQuality")
	delete(f, "DisableRecord")
	delete(f, "Assistants")
	delete(f, "RecordLayout")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRoomRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRoomResponseParams struct {
	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateRoomResponse struct {
	*tchttp.BaseResponse
	Response *CreateRoomResponseParams `json:"Response"`
}

func (r *CreateRoomResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRoomResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSupervisorRequestParams struct {
	// 应用ID。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 用户ID列表。
	Users []*string `json:"Users,omitempty" name:"Users"`
}

type CreateSupervisorRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 用户ID列表。
	Users []*string `json:"Users,omitempty" name:"Users"`
}

func (r *CreateSupervisorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSupervisorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Users")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSupervisorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSupervisorResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSupervisorResponse struct {
	*tchttp.BaseResponse
	Response *CreateSupervisorResponseParams `json:"Response"`
}

func (r *CreateSupervisorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSupervisorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAppCustomContentRequestParams struct {
	// 应用ID。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 指定需要删除的已设置的scene场景自定义元素，如果为空则删除应用下已设置的所有自定义元素。
	Scenes []*string `json:"Scenes,omitempty" name:"Scenes"`
}

type DeleteAppCustomContentRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 指定需要删除的已设置的scene场景自定义元素，如果为空则删除应用下已设置的所有自定义元素。
	Scenes []*string `json:"Scenes,omitempty" name:"Scenes"`
}

func (r *DeleteAppCustomContentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAppCustomContentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Scenes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAppCustomContentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAppCustomContentResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteAppCustomContentResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAppCustomContentResponseParams `json:"Response"`
}

func (r *DeleteAppCustomContentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAppCustomContentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDocumentRequestParams struct {
	// 文档ID。
	DocumentId *string `json:"DocumentId,omitempty" name:"DocumentId"`
}

type DeleteDocumentRequest struct {
	*tchttp.BaseRequest
	
	// 文档ID。
	DocumentId *string `json:"DocumentId,omitempty" name:"DocumentId"`
}

func (r *DeleteDocumentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDocumentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DocumentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDocumentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDocumentResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteDocumentResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDocumentResponseParams `json:"Response"`
}

func (r *DeleteDocumentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDocumentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGroupMemberRequestParams struct {
	// 群组ID，联合群组无法删除群组成员
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 成员列表，最大值200
	MemberIds []*string `json:"MemberIds,omitempty" name:"MemberIds"`
}

type DeleteGroupMemberRequest struct {
	*tchttp.BaseRequest
	
	// 群组ID，联合群组无法删除群组成员
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 成员列表，最大值200
	MemberIds []*string `json:"MemberIds,omitempty" name:"MemberIds"`
}

func (r *DeleteGroupMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGroupMemberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "SdkAppId")
	delete(f, "MemberIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteGroupMemberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGroupMemberResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteGroupMemberResponse struct {
	*tchttp.BaseResponse
	Response *DeleteGroupMemberResponseParams `json:"Response"`
}

func (r *DeleteGroupMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGroupMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGroupRequestParams struct {
	// 待删除群组ID列表
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

type DeleteGroupRequest struct {
	*tchttp.BaseRequest
	
	// 待删除群组ID列表
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

func (r *DeleteGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupIds")
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGroupResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteGroupResponseParams `json:"Response"`
}

func (r *DeleteGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRecordRequestParams struct {
	// 房间Id。
	RoomId *int64 `json:"RoomId,omitempty" name:"RoomId"`

	// 低代码互动课堂的SdkAppId。
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

type DeleteRecordRequest struct {
	*tchttp.BaseRequest
	
	// 房间Id。
	RoomId *int64 `json:"RoomId,omitempty" name:"RoomId"`

	// 低代码互动课堂的SdkAppId。
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

func (r *DeleteRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoomId")
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRecordResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteRecordResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRecordResponseParams `json:"Response"`
}

func (r *DeleteRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRoomRequestParams struct {
	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`
}

type DeleteRoomRequest struct {
	*tchttp.BaseRequest
	
	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`
}

func (r *DeleteRoomRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRoomRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRoomRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRoomResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteRoomResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRoomResponseParams `json:"Response"`
}

func (r *DeleteRoomResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRoomResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppDetailRequestParams struct {
	// 应用ID。低代码互动课堂的SdkAppId。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 开发商ID
	DeveloperId *string `json:"DeveloperId,omitempty" name:"DeveloperId"`
}

type DescribeAppDetailRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID。低代码互动课堂的SdkAppId。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 开发商ID
	DeveloperId *string `json:"DeveloperId,omitempty" name:"DeveloperId"`
}

func (r *DescribeAppDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "DeveloperId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAppDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppDetailResponseParams struct {
	// SDK 对应的AppId 
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 应用配置
	AppConfig *AppConfig `json:"AppConfig,omitempty" name:"AppConfig"`

	// 场景配置
	SceneConfig []*SceneItem `json:"SceneConfig,omitempty" name:"SceneConfig"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAppDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAppDetailResponseParams `json:"Response"`
}

func (r *DescribeAppDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCurrentMemberListRequestParams struct {
	// 房间Id。
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// 分页查询当前页数，从1开始递增。
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// 每页数据量，最大1000。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeCurrentMemberListRequest struct {
	*tchttp.BaseRequest
	
	// 房间Id。
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// 分页查询当前页数，从1开始递增。
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// 每页数据量，最大1000。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeCurrentMemberListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCurrentMemberListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoomId")
	delete(f, "Page")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCurrentMemberListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCurrentMemberListResponseParams struct {
	// 记录总数。当前房间的总人数。
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 成员记录列表。
	MemberRecords []*MemberRecord `json:"MemberRecords,omitempty" name:"MemberRecords"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCurrentMemberListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCurrentMemberListResponseParams `json:"Response"`
}

func (r *DescribeCurrentMemberListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCurrentMemberListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDocumentRequestParams struct {
	// 文档Id（唯一id）
	DocumentId *string `json:"DocumentId,omitempty" name:"DocumentId"`
}

type DescribeDocumentRequest struct {
	*tchttp.BaseRequest
	
	// 文档Id（唯一id）
	DocumentId *string `json:"DocumentId,omitempty" name:"DocumentId"`
}

func (r *DescribeDocumentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDocumentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DocumentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDocumentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDocumentResponseParams struct {
	// 文档Id
	DocumentId *string `json:"DocumentId,omitempty" name:"DocumentId"`

	// 文档原址url
	DocumentUrl *string `json:"DocumentUrl,omitempty" name:"DocumentUrl"`

	// 文档名称
	DocumentName *string `json:"DocumentName,omitempty" name:"DocumentName"`

	// 文档所有者UserId
	Owner *string `json:"Owner,omitempty" name:"Owner"`

	// 应用Id
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 文档权限
	Permission *uint64 `json:"Permission,omitempty" name:"Permission"`

	// 转码结果，无需转码为空，转码成功为结果url，转码失败为错误码
	TranscodeResult *string `json:"TranscodeResult,omitempty" name:"TranscodeResult"`

	// 转码类型
	TranscodeType *uint64 `json:"TranscodeType,omitempty" name:"TranscodeType"`

	// 转码进度， 0 - 100 表示（0% - 100%）
	TranscodeProgress *uint64 `json:"TranscodeProgress,omitempty" name:"TranscodeProgress"`

	// 转码状态，0为无需转码，1为正在转码，2为转码失败，3为转码成功
	TranscodeState *uint64 `json:"TranscodeState,omitempty" name:"TranscodeState"`

	// 转码失败后的错误信息
	TranscodeInfo *string `json:"TranscodeInfo,omitempty" name:"TranscodeInfo"`

	// 文档类型
	DocumentType *string `json:"DocumentType,omitempty" name:"DocumentType"`

	// 文档大小，单位：字节
	DocumentSize *uint64 `json:"DocumentSize,omitempty" name:"DocumentSize"`

	// 更新的UNIX时间戳
	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDocumentResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDocumentResponseParams `json:"Response"`
}

func (r *DescribeDocumentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDocumentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDocumentsByRoomRequestParams struct {
	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// 低代码互动课堂的SdkAppId
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 分页查询当前页数，从1开始递增，默认值为1
	Page *int64 `json:"Page,omitempty" name:"Page"`

	// 每页数据量，最大1000，默认值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 课件权限。
	// [0]：获取owner的私有课件；
	// [1]：获取owner的公开课件;
	// [0,1]：则获取owner的私有课件和公开课件；
	// [2]：获取owner的私有课件和所有人(包括owner)的公开课件。
	// 默认值为[2]
	Permission []*uint64 `json:"Permission,omitempty" name:"Permission"`

	// 文档所有者的user_id，不填默认获取SdkAppId下所有课件
	Owner *string `json:"Owner,omitempty" name:"Owner"`
}

type DescribeDocumentsByRoomRequest struct {
	*tchttp.BaseRequest
	
	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// 低代码互动课堂的SdkAppId
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 分页查询当前页数，从1开始递增，默认值为1
	Page *int64 `json:"Page,omitempty" name:"Page"`

	// 每页数据量，最大1000，默认值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 课件权限。
	// [0]：获取owner的私有课件；
	// [1]：获取owner的公开课件;
	// [0,1]：则获取owner的私有课件和公开课件；
	// [2]：获取owner的私有课件和所有人(包括owner)的公开课件。
	// 默认值为[2]
	Permission []*uint64 `json:"Permission,omitempty" name:"Permission"`

	// 文档所有者的user_id，不填默认获取SdkAppId下所有课件
	Owner *string `json:"Owner,omitempty" name:"Owner"`
}

func (r *DescribeDocumentsByRoomRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDocumentsByRoomRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoomId")
	delete(f, "SdkAppId")
	delete(f, "Page")
	delete(f, "Limit")
	delete(f, "Permission")
	delete(f, "Owner")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDocumentsByRoomRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDocumentsByRoomResponseParams struct {
	// 文档信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Documents []*DocumentInfo `json:"Documents,omitempty" name:"Documents"`

	// 符合查询条件文档总数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDocumentsByRoomResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDocumentsByRoomResponseParams `json:"Response"`
}

func (r *DescribeDocumentsByRoomResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDocumentsByRoomResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGroupListRequestParams struct {
	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 分页查询当前页数，默认从1开始递增。
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// 每页数据量，默认20，最大1000。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 主讲人ID筛选群组，与MemberId有且只有一个,都传时以此字段获取
	TeacherId *string `json:"TeacherId,omitempty" name:"TeacherId"`

	// 成员ID刷选群组，与TeacherId有且只有一个
	MemberId *string `json:"MemberId,omitempty" name:"MemberId"`
}

type DescribeGroupListRequest struct {
	*tchttp.BaseRequest
	
	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 分页查询当前页数，默认从1开始递增。
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// 每页数据量，默认20，最大1000。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 主讲人ID筛选群组，与MemberId有且只有一个,都传时以此字段获取
	TeacherId *string `json:"TeacherId,omitempty" name:"TeacherId"`

	// 成员ID刷选群组，与TeacherId有且只有一个
	MemberId *string `json:"MemberId,omitempty" name:"MemberId"`
}

func (r *DescribeGroupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Page")
	delete(f, "Limit")
	delete(f, "TeacherId")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGroupListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGroupListResponseParams struct {
	// 记录总数。当前匹配群组总数。
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 群组信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupInfos []*GroupInfo `json:"GroupInfos,omitempty" name:"GroupInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeGroupListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGroupListResponseParams `json:"Response"`
}

func (r *DescribeGroupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGroupMemberListRequestParams struct {
	// 群组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 分页值，默认1
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// 每页数据量，默认20，最大1000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeGroupMemberListRequest struct {
	*tchttp.BaseRequest
	
	// 群组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 分页值，默认1
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// 每页数据量，默认20，最大1000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeGroupMemberListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupMemberListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "SdkAppId")
	delete(f, "Page")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGroupMemberListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGroupMemberListResponseParams struct {
	// 符合查询条件总条数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 查询成员列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemberIds []*string `json:"MemberIds,omitempty" name:"MemberIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeGroupMemberListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGroupMemberListResponseParams `json:"Response"`
}

func (r *DescribeGroupMemberListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupMemberListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGroupRequestParams struct {
	// 群组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

type DescribeGroupRequest struct {
	*tchttp.BaseRequest
	
	// 群组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

func (r *DescribeGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGroupResponseParams struct {
	// 群组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 群组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 群主主讲人ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TeacherId *string `json:"TeacherId,omitempty" name:"TeacherId"`

	// 群组类型
	// 0-基础群组
	// 1-组合群组，若为1时会返回子群组ID
	GroupType *uint64 `json:"GroupType,omitempty" name:"GroupType"`

	// 子群组ID列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubGroupIds []*string `json:"SubGroupIds,omitempty" name:"SubGroupIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGroupResponseParams `json:"Response"`
}

func (r *DescribeGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoomRequestParams struct {
	// 房间Id。
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`
}

type DescribeRoomRequest struct {
	*tchttp.BaseRequest
	
	// 房间Id。
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`
}

func (r *DescribeRoomRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoomRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRoomRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoomResponseParams struct {
	// 房间名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 预定的房间开始时间，unix时间戳。
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 预定的房间结束时间，unix时间戳。
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 老师的UserId。
	TeacherId *string `json:"TeacherId,omitempty" name:"TeacherId"`

	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 分辨率。可以有如下取值：
	// 1 标清
	// 2 高清
	// 3 全高清
	Resolution *uint64 `json:"Resolution,omitempty" name:"Resolution"`

	// 最大连麦人数（不包括老师）。取值范围[0, 16]
	MaxMicNumber *uint64 `json:"MaxMicNumber,omitempty" name:"MaxMicNumber"`

	// 进入课堂时是否自动连麦。可以有以下取值：
	// 0 不自动连麦（需要手动申请上麦，默认值）
	// 1 自动连麦
	AutoMic *uint64 `json:"AutoMic,omitempty" name:"AutoMic"`

	// 高音质模式。可以有以下取值：
	// 0 不开启高音质（默认值）
	// 1 开启高音质
	AudioQuality *uint64 `json:"AudioQuality,omitempty" name:"AudioQuality"`

	// 房间子类型，可以有以下取值：
	// videodoc 文档+视频
	// video 纯视频
	SubType *string `json:"SubType,omitempty" name:"SubType"`

	// 上课后是否禁止自动录制。可以有以下取值：
	// 0 不禁止录制（自动开启录制，默认值）
	// 1 禁止录制
	// 注：如果该配置取值为0，录制将从上课后开始，课堂结束后停止。
	DisableRecord *uint64 `json:"DisableRecord,omitempty" name:"DisableRecord"`

	// 助教UserId列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Assistants []*string `json:"Assistants,omitempty" name:"Assistants"`

	// 录制地址（协议为https)。仅在房间结束后存在。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordUrl *string `json:"RecordUrl,omitempty" name:"RecordUrl"`

	// 课堂状态。0为未开始，1为已开始，2为已结束，3为已过期。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 房间绑定的群组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRoomResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRoomResponseParams `json:"Response"`
}

func (r *DescribeRoomResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoomResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoomStatisticsRequestParams struct {
	// 房间Id。
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// 分页查询当前页数，从1开始递增。
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// 每页数据量，最大1000。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeRoomStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 房间Id。
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// 分页查询当前页数，从1开始递增。
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// 每页数据量，最大1000。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeRoomStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoomStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoomId")
	delete(f, "Page")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRoomStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoomStatisticsResponseParams struct {
	// 峰值在线成员人数。
	PeakMemberNumber *uint64 `json:"PeakMemberNumber,omitempty" name:"PeakMemberNumber"`

	// 累计在线人数。
	MemberNumber *uint64 `json:"MemberNumber,omitempty" name:"MemberNumber"`

	// 记录总数。包含进入房间或者应到未到的。
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 成员记录列表。
	MemberRecords []*MemberRecord `json:"MemberRecords,omitempty" name:"MemberRecords"`

	// 秒级unix时间戳，实际房间开始时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealStartTime *uint64 `json:"RealStartTime,omitempty" name:"RealStartTime"`

	// 秒级unix时间戳，实际房间结束时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealEndTime *uint64 `json:"RealEndTime,omitempty" name:"RealEndTime"`

	// 房间消息总数。
	MessageCount *uint64 `json:"MessageCount,omitempty" name:"MessageCount"`

	// 房间连麦总数。
	MicCount *uint64 `json:"MicCount,omitempty" name:"MicCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRoomStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRoomStatisticsResponseParams `json:"Response"`
}

func (r *DescribeRoomStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoomStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSdkAppIdUsersRequestParams struct {
	// 应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 分页，默认值为1
	Page *int64 `json:"Page,omitempty" name:"Page"`

	// 分页数据限制，默认值为20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeSdkAppIdUsersRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 分页，默认值为1
	Page *int64 `json:"Page,omitempty" name:"Page"`

	// 分页数据限制，默认值为20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSdkAppIdUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSdkAppIdUsersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Page")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSdkAppIdUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSdkAppIdUsersResponseParams struct {
	// 用户总数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 当前获取用户信息数组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Users []*UserInfo `json:"Users,omitempty" name:"Users"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSdkAppIdUsersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSdkAppIdUsersResponseParams `json:"Response"`
}

func (r *DescribeSdkAppIdUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSdkAppIdUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserRequestParams struct {
	// 用户Id。
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

type DescribeUserRequest struct {
	*tchttp.BaseRequest
	
	// 用户Id。
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

func (r *DescribeUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserResponseParams struct {
	// 应用Id。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 用户Id。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户昵称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 用户头像Url。
	Avatar *string `json:"Avatar,omitempty" name:"Avatar"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUserResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserResponseParams `json:"Response"`
}

func (r *DescribeUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DocumentInfo struct {
	// 文档Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DocumentId *string `json:"DocumentId,omitempty" name:"DocumentId"`

	// 文档原址url
	// 注意：此字段可能返回 null，表示取不到有效值。
	DocumentUrl *string `json:"DocumentUrl,omitempty" name:"DocumentUrl"`

	// 文档名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DocumentName *string `json:"DocumentName,omitempty" name:"DocumentName"`

	// 文档所有者UserId
	// 注意：此字段可能返回 null，表示取不到有效值。
	Owner *string `json:"Owner,omitempty" name:"Owner"`

	// 应用Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 文档权限，0：私有课件 1：公共课件
	// 注意：此字段可能返回 null，表示取不到有效值。
	Permission *uint64 `json:"Permission,omitempty" name:"Permission"`

	// 转码结果，无需转码为空，转码成功为结果url，转码失败为错误码
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranscodeResult *string `json:"TranscodeResult,omitempty" name:"TranscodeResult"`

	// 转码类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranscodeType *uint64 `json:"TranscodeType,omitempty" name:"TranscodeType"`

	// 转码进度， 0 - 100 表示（0% - 100%）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranscodeProgress *uint64 `json:"TranscodeProgress,omitempty" name:"TranscodeProgress"`

	// 转码状态，0为无需转码，1为正在转码，2为转码失败，3为转码成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranscodeState *uint64 `json:"TranscodeState,omitempty" name:"TranscodeState"`

	// 转码失败后的错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranscodeInfo *string `json:"TranscodeInfo,omitempty" name:"TranscodeInfo"`

	// 文档类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DocumentType *string `json:"DocumentType,omitempty" name:"DocumentType"`

	// 文档大小，单位：字节
	// 注意：此字段可能返回 null，表示取不到有效值。
	DocumentSize *uint64 `json:"DocumentSize,omitempty" name:"DocumentSize"`

	// 更新的UNIX时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

// Predefined struct for user
type GetRoomMessageRequestParams struct {
	// 低代码互动课堂的SdkAppId。
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 房间Id。	
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// 消息序列。获取该序列以前前的消息(不包含该seq消息)
	Seq *int64 `json:"Seq,omitempty" name:"Seq"`

	// 消息拉取的条数。最大数量不能超过套餐包限制。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type GetRoomMessageRequest struct {
	*tchttp.BaseRequest
	
	// 低代码互动课堂的SdkAppId。
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 房间Id。	
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// 消息序列。获取该序列以前前的消息(不包含该seq消息)
	Seq *int64 `json:"Seq,omitempty" name:"Seq"`

	// 消息拉取的条数。最大数量不能超过套餐包限制。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetRoomMessageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRoomMessageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	delete(f, "Seq")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetRoomMessageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRoomMessageResponseParams struct {
	// 消息列表
	Messages []*MessageList `json:"Messages,omitempty" name:"Messages"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetRoomMessageResponse struct {
	*tchttp.BaseResponse
	Response *GetRoomMessageResponseParams `json:"Response"`
}

func (r *GetRoomMessageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRoomMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetWatermarkRequestParams struct {
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

type GetWatermarkRequest struct {
	*tchttp.BaseRequest
	
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

func (r *GetWatermarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetWatermarkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetWatermarkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetWatermarkResponseParams struct {
	// 老师视频区域的水印参数配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	TeacherLogo *WatermarkConfig `json:"TeacherLogo,omitempty" name:"TeacherLogo"`

	// 白板区域的水印参数配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	BoardLogo *WatermarkConfig `json:"BoardLogo,omitempty" name:"BoardLogo"`

	// 背景图片配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackgroundPicture *BackgroundPictureConfig `json:"BackgroundPicture,omitempty" name:"BackgroundPicture"`

	// 文字水印配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *TextMarkConfig `json:"Text,omitempty" name:"Text"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetWatermarkResponse struct {
	*tchttp.BaseResponse
	Response *GetWatermarkResponseParams `json:"Response"`
}

func (r *GetWatermarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetWatermarkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GroupBaseInfo struct {
	// 待创建群组名
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 群组主讲人ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TeacherId *string `json:"TeacherId,omitempty" name:"TeacherId"`
}

type GroupInfo struct {
	// 群组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 群组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 群组主讲人ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TeacherId *string `json:"TeacherId,omitempty" name:"TeacherId"`

	// 群组类型 
	// 0-基础群组 
	// 1-组合群组，若为1时会返回子群组ID列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupType *uint64 `json:"GroupType,omitempty" name:"GroupType"`

	// 子群组ID列表，如有。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubGroupIds *string `json:"SubGroupIds,omitempty" name:"SubGroupIds"`
}

// Predefined struct for user
type LoginOriginIdRequestParams struct {
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 用户在客户系统的Id，需要在同一应用下唯一。
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`
}

type LoginOriginIdRequest struct {
	*tchttp.BaseRequest
	
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 用户在客户系统的Id，需要在同一应用下唯一。
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`
}

func (r *LoginOriginIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LoginOriginIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "OriginId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "LoginOriginIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type LoginOriginIdResponseParams struct {
	// 用户Id。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 登录/注册成功后返回登录态token。有效期7天。
	Token *string `json:"Token,omitempty" name:"Token"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type LoginOriginIdResponse struct {
	*tchttp.BaseResponse
	Response *LoginOriginIdResponseParams `json:"Response"`
}

func (r *LoginOriginIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LoginOriginIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type LoginUserRequestParams struct {
	// 注册获取的用户id。
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

type LoginUserRequest struct {
	*tchttp.BaseRequest
	
	// 注册获取的用户id。
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

func (r *LoginUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LoginUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "LoginUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type LoginUserResponseParams struct {
	// 用户Id。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 登录/注册成功后返回登录态token。有效期7天。
	Token *string `json:"Token,omitempty" name:"Token"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type LoginUserResponse struct {
	*tchttp.BaseResponse
	Response *LoginUserResponseParams `json:"Response"`
}

func (r *LoginUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LoginUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MemberRecord struct {
	// 用户ID。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户名称。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 在线时长，单位秒。
	PresentTime *uint64 `json:"PresentTime,omitempty" name:"PresentTime"`

	// 是否开启摄像头。
	Camera *uint64 `json:"Camera,omitempty" name:"Camera"`

	// 是否开启麦克风。
	Mic *uint64 `json:"Mic,omitempty" name:"Mic"`

	// 是否禁言。
	Silence *uint64 `json:"Silence,omitempty" name:"Silence"`

	// 回答问题数量。
	AnswerQuestions *uint64 `json:"AnswerQuestions,omitempty" name:"AnswerQuestions"`

	// 举手数量。
	HandUps *uint64 `json:"HandUps,omitempty" name:"HandUps"`

	// 首次进入房间的unix时间戳。
	FirstJoinTimestamp *uint64 `json:"FirstJoinTimestamp,omitempty" name:"FirstJoinTimestamp"`

	// 最后一次退出房间的unix时间戳。
	LastQuitTimestamp *uint64 `json:"LastQuitTimestamp,omitempty" name:"LastQuitTimestamp"`

	// 奖励次数。
	Rewords *uint64 `json:"Rewords,omitempty" name:"Rewords"`

	// 用户IP。
	IPAddress *string `json:"IPAddress,omitempty" name:"IPAddress"`

	// 用户位置信息。
	Location *string `json:"Location,omitempty" name:"Location"`

	// 用户设备平台信息。0:unknown  1:windows  2:mac  3:android  4:ios  5:web   6:h5   7:miniprogram （小程序）
	Device *int64 `json:"Device,omitempty" name:"Device"`

	// 每个成员上麦次数。
	PerMemberMicCount *int64 `json:"PerMemberMicCount,omitempty" name:"PerMemberMicCount"`

	// 每个成员发送消息数量。
	PerMemberMessageCount *int64 `json:"PerMemberMessageCount,omitempty" name:"PerMemberMessageCount"`
}

type MessageItem struct {
	// 消息类型。0表示文本消息，1表示图片消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MessageType *int64 `json:"MessageType,omitempty" name:"MessageType"`

	// 文本消息内容。message type为0时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TextMessage *string `json:"TextMessage,omitempty" name:"TextMessage"`

	// 图片消息URL。 message type为1时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageMessage *string `json:"ImageMessage,omitempty" name:"ImageMessage"`
}

type MessageList struct {
	// 消息时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// 消息发送者
	// 注意：此字段可能返回 null，表示取不到有效值。
	FromAccount *string `json:"FromAccount,omitempty" name:"FromAccount"`

	// 消息序列号，当前课堂内唯一且单调递增
	// 注意：此字段可能返回 null，表示取不到有效值。
	Seq *int64 `json:"Seq,omitempty" name:"Seq"`

	// 历史消息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	MessageBody []*MessageItem `json:"MessageBody,omitempty" name:"MessageBody"`
}

// Predefined struct for user
type ModifyAppRequestParams struct {
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 回调地址。
	Callback *string `json:"Callback,omitempty" name:"Callback"`

	// 回调key。
	CallbackKey *string `json:"CallbackKey,omitempty" name:"CallbackKey"`
}

type ModifyAppRequest struct {
	*tchttp.BaseRequest
	
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 回调地址。
	Callback *string `json:"Callback,omitempty" name:"Callback"`

	// 回调key。
	CallbackKey *string `json:"CallbackKey,omitempty" name:"CallbackKey"`
}

func (r *ModifyAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Callback")
	delete(f, "CallbackKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAppResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAppResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAppResponseParams `json:"Response"`
}

func (r *ModifyAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGroupRequestParams struct {
	// 需要修改的群组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 默认绑定主讲老师ID
	TeacherId *string `json:"TeacherId,omitempty" name:"TeacherId"`

	// 待修改的群组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

type ModifyGroupRequest struct {
	*tchttp.BaseRequest
	
	// 需要修改的群组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 默认绑定主讲老师ID
	TeacherId *string `json:"TeacherId,omitempty" name:"TeacherId"`

	// 待修改的群组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

func (r *ModifyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "SdkAppId")
	delete(f, "TeacherId")
	delete(f, "GroupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGroupResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyGroupResponseParams `json:"Response"`
}

func (r *ModifyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRoomRequestParams struct {
	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// 低代码互动课堂的SdkAppId
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 预定的房间开始时间，unix时间戳。直播开始后不允许修改。
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 预定的房间结束时间，unix时间戳。直播开始后不允许修改。
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 老师ID。直播开始后不允许修改。
	TeacherId *string `json:"TeacherId,omitempty" name:"TeacherId"`

	// 房间名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分辨率。可以有如下取值：
	// 1 标清
	// 2 高清
	// 3 全高清
	// 直播开始后不允许修改。
	Resolution *uint64 `json:"Resolution,omitempty" name:"Resolution"`

	// 最大连麦人数（不包括老师）。取值范围[0, 17)
	// 直播开始后不允许修改。
	MaxMicNumber *uint64 `json:"MaxMicNumber,omitempty" name:"MaxMicNumber"`

	// 进入房间时是否自动连麦。可以有以下取值：
	// 0 不自动连麦（默认值）
	// 1 自动连麦
	// 直播开始后不允许修改。
	AutoMic *uint64 `json:"AutoMic,omitempty" name:"AutoMic"`

	// 高音质模式。可以有以下取值：
	// 0 不开启高音质（默认值）
	// 1 开启高音质
	// 直播开始后不允许修改。
	AudioQuality *uint64 `json:"AudioQuality,omitempty" name:"AudioQuality"`

	// 房间子类型，可以有以下取值：
	// videodoc 文档+视频
	// video 纯视频
	// coteaching 双师
	// 直播开始后不允许修改。
	SubType *string `json:"SubType,omitempty" name:"SubType"`

	// 禁止录制。可以有以下取值：
	// 0 不禁止录制（默认值）
	// 1 禁止录制
	// 直播开始后不允许修改。
	DisableRecord *uint64 `json:"DisableRecord,omitempty" name:"DisableRecord"`

	// 助教Id列表。直播开始后不允许修改。
	Assistants []*string `json:"Assistants,omitempty" name:"Assistants"`

	// 房间绑定的群组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type ModifyRoomRequest struct {
	*tchttp.BaseRequest
	
	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// 低代码互动课堂的SdkAppId
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 预定的房间开始时间，unix时间戳。直播开始后不允许修改。
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 预定的房间结束时间，unix时间戳。直播开始后不允许修改。
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 老师ID。直播开始后不允许修改。
	TeacherId *string `json:"TeacherId,omitempty" name:"TeacherId"`

	// 房间名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分辨率。可以有如下取值：
	// 1 标清
	// 2 高清
	// 3 全高清
	// 直播开始后不允许修改。
	Resolution *uint64 `json:"Resolution,omitempty" name:"Resolution"`

	// 最大连麦人数（不包括老师）。取值范围[0, 17)
	// 直播开始后不允许修改。
	MaxMicNumber *uint64 `json:"MaxMicNumber,omitempty" name:"MaxMicNumber"`

	// 进入房间时是否自动连麦。可以有以下取值：
	// 0 不自动连麦（默认值）
	// 1 自动连麦
	// 直播开始后不允许修改。
	AutoMic *uint64 `json:"AutoMic,omitempty" name:"AutoMic"`

	// 高音质模式。可以有以下取值：
	// 0 不开启高音质（默认值）
	// 1 开启高音质
	// 直播开始后不允许修改。
	AudioQuality *uint64 `json:"AudioQuality,omitempty" name:"AudioQuality"`

	// 房间子类型，可以有以下取值：
	// videodoc 文档+视频
	// video 纯视频
	// coteaching 双师
	// 直播开始后不允许修改。
	SubType *string `json:"SubType,omitempty" name:"SubType"`

	// 禁止录制。可以有以下取值：
	// 0 不禁止录制（默认值）
	// 1 禁止录制
	// 直播开始后不允许修改。
	DisableRecord *uint64 `json:"DisableRecord,omitempty" name:"DisableRecord"`

	// 助教Id列表。直播开始后不允许修改。
	Assistants []*string `json:"Assistants,omitempty" name:"Assistants"`

	// 房间绑定的群组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *ModifyRoomRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRoomRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoomId")
	delete(f, "SdkAppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "TeacherId")
	delete(f, "Name")
	delete(f, "Resolution")
	delete(f, "MaxMicNumber")
	delete(f, "AutoMic")
	delete(f, "AudioQuality")
	delete(f, "SubType")
	delete(f, "DisableRecord")
	delete(f, "Assistants")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRoomRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRoomResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyRoomResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRoomResponseParams `json:"Response"`
}

func (r *ModifyRoomResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRoomResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserProfileRequestParams struct {
	// 待修改用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 待修改的用户名
	Nickname *string `json:"Nickname,omitempty" name:"Nickname"`

	// 待修改头像url
	Avatar *string `json:"Avatar,omitempty" name:"Avatar"`
}

type ModifyUserProfileRequest struct {
	*tchttp.BaseRequest
	
	// 待修改用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 待修改的用户名
	Nickname *string `json:"Nickname,omitempty" name:"Nickname"`

	// 待修改头像url
	Avatar *string `json:"Avatar,omitempty" name:"Avatar"`
}

func (r *ModifyUserProfileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserProfileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "Nickname")
	delete(f, "Avatar")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserProfileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserProfileResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyUserProfileResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUserProfileResponseParams `json:"Response"`
}

func (r *ModifyUserProfileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserProfileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterUserRequestParams struct {
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 用户名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 用户在客户系统的Id，需要在同一应用下唯一。
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`

	// 用户头像。
	Avatar *string `json:"Avatar,omitempty" name:"Avatar"`
}

type RegisterUserRequest struct {
	*tchttp.BaseRequest
	
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 用户名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 用户在客户系统的Id，需要在同一应用下唯一。
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`

	// 用户头像。
	Avatar *string `json:"Avatar,omitempty" name:"Avatar"`
}

func (r *RegisterUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Name")
	delete(f, "OriginId")
	delete(f, "Avatar")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RegisterUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterUserResponseParams struct {
	// 用户Id。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 登录/注册成功后返回登录态token。有效期7天。
	Token *string `json:"Token,omitempty" name:"Token"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RegisterUserResponse struct {
	*tchttp.BaseResponse
	Response *RegisterUserResponseParams `json:"Response"`
}

func (r *RegisterUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RoomInfo struct {
	// 房间名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 预定的房间开始时间，unix时间戳。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 预定的房间结束时间，unix时间戳。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 分辨率。可以有如下取值： 1 标清 2 高清 3 全高清
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resolution *uint64 `json:"Resolution,omitempty" name:"Resolution"`

	// 最大连麦人数（不包括老师）。取值范围[0, 16]
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxMicNumber *uint64 `json:"MaxMicNumber,omitempty" name:"MaxMicNumber"`

	// 房间子类型，可以有以下取值： videodoc 文档+视频 video 纯视频
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubType *string `json:"SubType,omitempty" name:"SubType"`

	// 老师ID。通过[注册用户]接口获取的UserId。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TeacherId *string `json:"TeacherId,omitempty" name:"TeacherId"`

	// 进入课堂时是否自动连麦。可以有以下取值： 0 不自动连麦（需要手动申请上麦，默认值） 1 自动连麦
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoMic *uint64 `json:"AutoMic,omitempty" name:"AutoMic"`

	// 释放音视频权限后是否自动取消连麦。可以有以下取值： 0 自动取消连麦（默认值） 1 保持连麦状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	TurnOffMic *uint64 `json:"TurnOffMic,omitempty" name:"TurnOffMic"`

	// 高音质模式。可以有以下取值： 0 不开启高音质（默认值） 1 开启高音质
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioQuality *uint64 `json:"AudioQuality,omitempty" name:"AudioQuality"`

	// 上课后是否禁止自动录制。可以有以下取值： 0 不禁止录制（自动开启录制，默认值） 1 禁止录制 注：如果该配置取值为0，录制将从上课后开始，课堂结束后停止。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisableRecord *uint64 `json:"DisableRecord,omitempty" name:"DisableRecord"`

	// 助教Id列表。通过[注册用户]接口获取的UserId。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Assistants []*string `json:"Assistants,omitempty" name:"Assistants"`

	// rtc人数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RTCAudienceNumber *uint64 `json:"RTCAudienceNumber,omitempty" name:"RTCAudienceNumber"`

	// 观看类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudienceType *uint64 `json:"AudienceType,omitempty" name:"AudienceType"`

	// 录制布局。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordLayout *uint64 `json:"RecordLayout,omitempty" name:"RecordLayout"`

	// 房间绑定的群组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type SceneItem struct {

}

// Predefined struct for user
type SetAppCustomContentRequestParams struct {
	// 自定义内容。
	CustomContent []*AppCustomContent `json:"CustomContent,omitempty" name:"CustomContent"`

	// 应用ID。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

type SetAppCustomContentRequest struct {
	*tchttp.BaseRequest
	
	// 自定义内容。
	CustomContent []*AppCustomContent `json:"CustomContent,omitempty" name:"CustomContent"`

	// 应用ID。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

func (r *SetAppCustomContentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetAppCustomContentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomContent")
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetAppCustomContentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetAppCustomContentResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetAppCustomContentResponse struct {
	*tchttp.BaseResponse
	Response *SetAppCustomContentResponseParams `json:"Response"`
}

func (r *SetAppCustomContentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetAppCustomContentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetWatermarkRequestParams struct {
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 老师视频区域的水印参数地址，设置为空字符串表示删除
	TeacherUrl *string `json:"TeacherUrl,omitempty" name:"TeacherUrl"`

	// 白板视频区域的水印参数地址，设置为空字符串表示删除
	BoardUrl *string `json:"BoardUrl,omitempty" name:"BoardUrl"`

	// 视频默认图片（在没有视频流的时候显示），设置为空字符串表示删除
	VideoUrl *string `json:"VideoUrl,omitempty" name:"VideoUrl"`

	// 白板区域水印的宽度，取值:0-100，默认为0，表示区域X方向的百分比
	BoardW *float64 `json:"BoardW,omitempty" name:"BoardW"`

	// 白板区域水印的高度，取值:0-100，默认为0, 表示区域Y方向的百分比
	BoardH *float64 `json:"BoardH,omitempty" name:"BoardH"`

	// 白板区域水印X偏移, 取值:0-100, 表示区域X方向的百分比。比如50，则表示位于X轴中间
	BoardX *float64 `json:"BoardX,omitempty" name:"BoardX"`

	// 白板区域水印Y偏移, 取值:0-100, 表示区域X方向的百分比。比如50，则表示位于X轴中间
	BoardY *float64 `json:"BoardY,omitempty" name:"BoardY"`

	// 老师视频区域水印的宽度，取值:0-100，默认为0，表示区域X方向的百分比
	TeacherW *float64 `json:"TeacherW,omitempty" name:"TeacherW"`

	// 老师视频区域水印的高度，取值:0-100，默认为0, 表示区域Y方向的百分比
	TeacherH *float64 `json:"TeacherH,omitempty" name:"TeacherH"`

	// 老师视频区域水印X偏移, 取值:0-100, 表示区域X方向的百分比。比如50，则表示位于X轴中间
	TeacherX *float64 `json:"TeacherX,omitempty" name:"TeacherX"`

	// 老师视频区域水印Y偏移, 取值:0-100, 表示区域X方向的百分比。比如50，则表示位于X轴中间
	TeacherY *float64 `json:"TeacherY,omitempty" name:"TeacherY"`

	// 文字水印内容，设置为空字符串表示删除
	Text *string `json:"Text,omitempty" name:"Text"`

	// 文字水印颜色
	TextColor *string `json:"TextColor,omitempty" name:"TextColor"`
}

type SetWatermarkRequest struct {
	*tchttp.BaseRequest
	
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 老师视频区域的水印参数地址，设置为空字符串表示删除
	TeacherUrl *string `json:"TeacherUrl,omitempty" name:"TeacherUrl"`

	// 白板视频区域的水印参数地址，设置为空字符串表示删除
	BoardUrl *string `json:"BoardUrl,omitempty" name:"BoardUrl"`

	// 视频默认图片（在没有视频流的时候显示），设置为空字符串表示删除
	VideoUrl *string `json:"VideoUrl,omitempty" name:"VideoUrl"`

	// 白板区域水印的宽度，取值:0-100，默认为0，表示区域X方向的百分比
	BoardW *float64 `json:"BoardW,omitempty" name:"BoardW"`

	// 白板区域水印的高度，取值:0-100，默认为0, 表示区域Y方向的百分比
	BoardH *float64 `json:"BoardH,omitempty" name:"BoardH"`

	// 白板区域水印X偏移, 取值:0-100, 表示区域X方向的百分比。比如50，则表示位于X轴中间
	BoardX *float64 `json:"BoardX,omitempty" name:"BoardX"`

	// 白板区域水印Y偏移, 取值:0-100, 表示区域X方向的百分比。比如50，则表示位于X轴中间
	BoardY *float64 `json:"BoardY,omitempty" name:"BoardY"`

	// 老师视频区域水印的宽度，取值:0-100，默认为0，表示区域X方向的百分比
	TeacherW *float64 `json:"TeacherW,omitempty" name:"TeacherW"`

	// 老师视频区域水印的高度，取值:0-100，默认为0, 表示区域Y方向的百分比
	TeacherH *float64 `json:"TeacherH,omitempty" name:"TeacherH"`

	// 老师视频区域水印X偏移, 取值:0-100, 表示区域X方向的百分比。比如50，则表示位于X轴中间
	TeacherX *float64 `json:"TeacherX,omitempty" name:"TeacherX"`

	// 老师视频区域水印Y偏移, 取值:0-100, 表示区域X方向的百分比。比如50，则表示位于X轴中间
	TeacherY *float64 `json:"TeacherY,omitempty" name:"TeacherY"`

	// 文字水印内容，设置为空字符串表示删除
	Text *string `json:"Text,omitempty" name:"Text"`

	// 文字水印颜色
	TextColor *string `json:"TextColor,omitempty" name:"TextColor"`
}

func (r *SetWatermarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetWatermarkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TeacherUrl")
	delete(f, "BoardUrl")
	delete(f, "VideoUrl")
	delete(f, "BoardW")
	delete(f, "BoardH")
	delete(f, "BoardX")
	delete(f, "BoardY")
	delete(f, "TeacherW")
	delete(f, "TeacherH")
	delete(f, "TeacherX")
	delete(f, "TeacherY")
	delete(f, "Text")
	delete(f, "TextColor")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetWatermarkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetWatermarkResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetWatermarkResponse struct {
	*tchttp.BaseResponse
	Response *SetWatermarkResponseParams `json:"Response"`
}

func (r *SetWatermarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetWatermarkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TextMarkConfig struct {
	// 文字水印内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitempty" name:"Text"`

	// 文字水印颜色
	// 注意：此字段可能返回 null，表示取不到有效值。
	Color *string `json:"Color,omitempty" name:"Color"`
}

// Predefined struct for user
type UnbindDocumentFromRoomRequestParams struct {
	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// 文档ID。
	DocumentId *string `json:"DocumentId,omitempty" name:"DocumentId"`
}

type UnbindDocumentFromRoomRequest struct {
	*tchttp.BaseRequest
	
	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// 文档ID。
	DocumentId *string `json:"DocumentId,omitempty" name:"DocumentId"`
}

func (r *UnbindDocumentFromRoomRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindDocumentFromRoomRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoomId")
	delete(f, "DocumentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindDocumentFromRoomRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindDocumentFromRoomResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnbindDocumentFromRoomResponse struct {
	*tchttp.BaseResponse
	Response *UnbindDocumentFromRoomResponseParams `json:"Response"`
}

func (r *UnbindDocumentFromRoomResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindDocumentFromRoomResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserInfo struct {
	// 应用Id。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 用户Id。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户昵称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 用户头像Url。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Avatar *string `json:"Avatar,omitempty" name:"Avatar"`
}

type WatermarkConfig struct {
	// 水印图片的url
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 水印宽。为比例值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Width *float64 `json:"Width,omitempty" name:"Width"`

	// 水印高。为比例值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Height *float64 `json:"Height,omitempty" name:"Height"`

	// 水印X偏移, 取值:0-100, 表示区域X方向的百分比。比如50，则表示位于X轴中间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocationX *float64 `json:"LocationX,omitempty" name:"LocationX"`

	// 水印Y偏移, 取值:0-100, 表示区域Y方向的百分比。比如50，则表示位于Y轴中间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocationY *float64 `json:"LocationY,omitempty" name:"LocationY"`
}