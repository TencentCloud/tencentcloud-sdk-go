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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type AddGroupMemberRequestParams struct {
	// 群组ID
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 成员列表，最大值200
	MemberIds []*string `json:"MemberIds,omitnil" name:"MemberIds"`
}

type AddGroupMemberRequest struct {
	*tchttp.BaseRequest
	
	// 群组ID
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 成员列表，最大值200
	MemberIds []*string `json:"MemberIds,omitnil" name:"MemberIds"`
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
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

type AnswerInfo struct {
	// 用户名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 答案（按照位表示是否选择，如0x1表示选择A，0x11表示选择AB）
	Answer *uint64 `json:"Answer,omitnil" name:"Answer"`

	// 答题用时
	CostTime *uint64 `json:"CostTime,omitnil" name:"CostTime"`

	// 用户ID
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 答案是否正确（1正确0错误）
	IsCorrect *uint64 `json:"IsCorrect,omitnil" name:"IsCorrect"`
}

type AnswerStat struct {
	// 选项（按照位表示是否选择，如0x1表示选择A，0x11表示选择AB）
	Answer *int64 `json:"Answer,omitnil" name:"Answer"`

	// 答题人数
	Count *int64 `json:"Count,omitnil" name:"Count"`
}

type AppConfig struct {

}

type AppCustomContent struct {
	// 场景参数，一个应用下可以设置多个不同场景。
	Scene *string `json:"Scene,omitnil" name:"Scene"`

	// logo地址，用于上课时展示的课堂或平台图标，支持开发商自定义业务品牌展示。
	LogoUrl *string `json:"LogoUrl,omitnil" name:"LogoUrl"`

	// HomeUrl：主页地址，用于上课结束后课堂跳转，支持跳转到自己的业务系统。如果配置为空则下课后关闭课堂页面。
	HomeUrl *string `json:"HomeUrl,omitnil" name:"HomeUrl"`

	// JsUrl ：自定义js。针对应用用于开发上自定义课堂界面、模块功能、监控操作，支持数据请求与响应处理。
	JsUrl *string `json:"JsUrl,omitnil" name:"JsUrl"`

	// Css : 自定义的css。针对应用用于支持课堂界面的、模块的UI渲染修改、皮肤配色修改、功能模块的隐藏和展示。
	CssUrl *string `json:"CssUrl,omitnil" name:"CssUrl"`
}

type BackgroundPictureConfig struct {
	// 背景图片的url
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil" name:"Url"`
}

// Predefined struct for user
type BatchAddGroupMemberRequestParams struct {
	// 待添加群组ID列表，最大值100
	GroupIds []*string `json:"GroupIds,omitnil" name:"GroupIds"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 待添加成员列表，最大值200
	MemberIds []*string `json:"MemberIds,omitnil" name:"MemberIds"`
}

type BatchAddGroupMemberRequest struct {
	*tchttp.BaseRequest
	
	// 待添加群组ID列表，最大值100
	GroupIds []*string `json:"GroupIds,omitnil" name:"GroupIds"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 待添加成员列表，最大值200
	MemberIds []*string `json:"MemberIds,omitnil" name:"MemberIds"`
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
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 批量创建群组基础信息，最大长度限制256
	GroupBaseInfos []*GroupBaseInfo `json:"GroupBaseInfos,omitnil" name:"GroupBaseInfos"`

	// 群组绑定的成员列表，一次性最多200个
	MemberIds []*string `json:"MemberIds,omitnil" name:"MemberIds"`
}

type BatchCreateGroupWithMembersRequest struct {
	*tchttp.BaseRequest
	
	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 批量创建群组基础信息，最大长度限制256
	GroupBaseInfos []*GroupBaseInfo `json:"GroupBaseInfos,omitnil" name:"GroupBaseInfos"`

	// 群组绑定的成员列表，一次性最多200个
	MemberIds []*string `json:"MemberIds,omitnil" name:"MemberIds"`
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
	GroupIds []*string `json:"GroupIds,omitnil" name:"GroupIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 创建房间ID列表
	RoomInfos []*RoomInfo `json:"RoomInfos,omitnil" name:"RoomInfos"`
}

type BatchCreateRoomRequest struct {
	*tchttp.BaseRequest
	
	// 低代码平台的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 创建房间ID列表
	RoomInfos []*RoomInfo `json:"RoomInfos,omitnil" name:"RoomInfos"`
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
	RoomIds []*uint64 `json:"RoomIds,omitnil" name:"RoomIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	GroupIds []*string `json:"GroupIds,omitnil" name:"GroupIds"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 待添加成员列表，最大值256
	MemberIds []*string `json:"MemberIds,omitnil" name:"MemberIds"`
}

type BatchDeleteGroupMemberRequest struct {
	*tchttp.BaseRequest
	
	// 待添加群组ID列表，最大值100
	GroupIds []*string `json:"GroupIds,omitnil" name:"GroupIds"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 待添加成员列表，最大值256
	MemberIds []*string `json:"MemberIds,omitnil" name:"MemberIds"`
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
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	RoomIds []*int64 `json:"RoomIds,omitnil" name:"RoomIds"`

	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`
}

type BatchDeleteRecordRequest struct {
	*tchttp.BaseRequest
	
	// 房间ID列表
	RoomIds []*int64 `json:"RoomIds,omitnil" name:"RoomIds"`

	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`
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
	RoomIds []*int64 `json:"RoomIds,omitnil" name:"RoomIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type BatchDescribeDocumentRequestParams struct {
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 分页查询当前页数，从1开始递增
	Page *int64 `json:"Page,omitnil" name:"Page"`

	// 每页数据量，最大1000
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 课件权限。[0]：获取owner的私有课件；[1]：获取owner的公开课件; [0,1]：则获取owner的私有课件和公开课件；[2]：获取owner的私有课件和所有人(包括owner)的公开课件
	Permission []*uint64 `json:"Permission,omitnil" name:"Permission"`

	// 课件所有者的user_id，不填默认获取SdkAppId下所有课件
	Owner *string `json:"Owner,omitnil" name:"Owner"`

	// 课件名称搜索词
	Keyword *string `json:"Keyword,omitnil" name:"Keyword"`

	// 课件id列表，从列表中查询，忽略错误的id
	DocumentId []*string `json:"DocumentId,omitnil" name:"DocumentId"`
}

type BatchDescribeDocumentRequest struct {
	*tchttp.BaseRequest
	
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 分页查询当前页数，从1开始递增
	Page *int64 `json:"Page,omitnil" name:"Page"`

	// 每页数据量，最大1000
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 课件权限。[0]：获取owner的私有课件；[1]：获取owner的公开课件; [0,1]：则获取owner的私有课件和公开课件；[2]：获取owner的私有课件和所有人(包括owner)的公开课件
	Permission []*uint64 `json:"Permission,omitnil" name:"Permission"`

	// 课件所有者的user_id，不填默认获取SdkAppId下所有课件
	Owner *string `json:"Owner,omitnil" name:"Owner"`

	// 课件名称搜索词
	Keyword *string `json:"Keyword,omitnil" name:"Keyword"`

	// 课件id列表，从列表中查询，忽略错误的id
	DocumentId []*string `json:"DocumentId,omitnil" name:"DocumentId"`
}

func (r *BatchDescribeDocumentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDescribeDocumentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Page")
	delete(f, "Limit")
	delete(f, "Permission")
	delete(f, "Owner")
	delete(f, "Keyword")
	delete(f, "DocumentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchDescribeDocumentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchDescribeDocumentResponseParams struct {
	// 符合查询条件文档总数
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 文档信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Documents []*DocumentInfo `json:"Documents,omitnil" name:"Documents"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type BatchDescribeDocumentResponse struct {
	*tchttp.BaseResponse
	Response *BatchDescribeDocumentResponseParams `json:"Response"`
}

func (r *BatchDescribeDocumentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDescribeDocumentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchRegisterRequestParams struct {
	// 批量注册用户信息列表
	Users []*BatchUserRequest `json:"Users,omitnil" name:"Users"`
}

type BatchRegisterRequest struct {
	*tchttp.BaseRequest
	
	// 批量注册用户信息列表
	Users []*BatchUserRequest `json:"Users,omitnil" name:"Users"`
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
	Users []*BatchUserInfo `json:"Users,omitnil" name:"Users"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 用户ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 用户在客户系统的Id。 若用户注册时该字段为空，则默认为 UserId 值一致。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginId *string `json:"OriginId,omitnil" name:"OriginId"`
}

type BatchUserRequest struct {
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 用户名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 用户在客户系统的Id，需要在同一应用下唯一。
	OriginId *string `json:"OriginId,omitnil" name:"OriginId"`

	// 用户头像。
	Avatar *string `json:"Avatar,omitnil" name:"Avatar"`
}

// Predefined struct for user
type BindDocumentToRoomRequestParams struct {
	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`

	// 文档ID。
	DocumentId *string `json:"DocumentId,omitnil" name:"DocumentId"`

	// 绑定类型。后台可透传到客户端，默认为0。客户端可以根据这个字段实现业务逻辑。
	BindType *uint64 `json:"BindType,omitnil" name:"BindType"`
}

type BindDocumentToRoomRequest struct {
	*tchttp.BaseRequest
	
	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`

	// 文档ID。
	DocumentId *string `json:"DocumentId,omitnil" name:"DocumentId"`

	// 绑定类型。后台可透传到客户端，默认为0。客户端可以根据这个字段实现业务逻辑。
	BindType *uint64 `json:"BindType,omitnil" name:"BindType"`
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
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

type ClassScoreItem struct {
	// 课堂iD
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`

	// 用户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 评分时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *uint64 `json:"CreateTime,omitnil" name:"CreateTime"`

	// 课堂评分
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *uint64 `json:"Score,omitnil" name:"Score"`

	// 课堂评价
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScoreMsg *string `json:"ScoreMsg,omitnil" name:"ScoreMsg"`
}

// Predefined struct for user
type CreateDocumentRequestParams struct {
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 文档地址。
	DocumentUrl *string `json:"DocumentUrl,omitnil" name:"DocumentUrl"`

	// 文档名称。
	DocumentName *string `json:"DocumentName,omitnil" name:"DocumentName"`

	// 文档所有者的Id
	Owner *string `json:"Owner,omitnil" name:"Owner"`

	// 转码类型，可以有如下取值：
	// 0 无需转码（默认），bmp，jpg，jpeg，png，gif
	// 1 需要转码的文档，ppt，pptx，pdf，doc，docx，xls，xlsx
	// 2 需要转码的视频，mp4，3pg，mpeg，avi，flv，wmv，rm，h264等
	// 2 需要转码的音频，mp3，wav，wma，aac，flac，opus
	TranscodeType *uint64 `json:"TranscodeType,omitnil" name:"TranscodeType"`

	// 权限，可以有如下取值：
	// 0 私有文档（默认）
	// 1 公共文档
	Permission *uint64 `json:"Permission,omitnil" name:"Permission"`

	// 文档后缀名。
	DocumentType *string `json:"DocumentType,omitnil" name:"DocumentType"`

	// 文档大小，单位 字节
	DocumentSize *uint64 `json:"DocumentSize,omitnil" name:"DocumentSize"`

	// 是否对不支持元素开启自动处理的功能。默认关闭。
	// 自动处理的元素如下：
	// 1. 墨迹：移除不支持的墨迹（例如WPS墨迹）
	// 2. 自动翻页：移除PPT上所有自动翻页设置，并设置为单击鼠标翻页
	// 3. 已损坏音视频：移除PPT上对损坏音视频的引用
	AutoHandleUnsupportedElement *bool `json:"AutoHandleUnsupportedElement,omitnil" name:"AutoHandleUnsupportedElement"`
}

type CreateDocumentRequest struct {
	*tchttp.BaseRequest
	
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 文档地址。
	DocumentUrl *string `json:"DocumentUrl,omitnil" name:"DocumentUrl"`

	// 文档名称。
	DocumentName *string `json:"DocumentName,omitnil" name:"DocumentName"`

	// 文档所有者的Id
	Owner *string `json:"Owner,omitnil" name:"Owner"`

	// 转码类型，可以有如下取值：
	// 0 无需转码（默认），bmp，jpg，jpeg，png，gif
	// 1 需要转码的文档，ppt，pptx，pdf，doc，docx，xls，xlsx
	// 2 需要转码的视频，mp4，3pg，mpeg，avi，flv，wmv，rm，h264等
	// 2 需要转码的音频，mp3，wav，wma，aac，flac，opus
	TranscodeType *uint64 `json:"TranscodeType,omitnil" name:"TranscodeType"`

	// 权限，可以有如下取值：
	// 0 私有文档（默认）
	// 1 公共文档
	Permission *uint64 `json:"Permission,omitnil" name:"Permission"`

	// 文档后缀名。
	DocumentType *string `json:"DocumentType,omitnil" name:"DocumentType"`

	// 文档大小，单位 字节
	DocumentSize *uint64 `json:"DocumentSize,omitnil" name:"DocumentSize"`

	// 是否对不支持元素开启自动处理的功能。默认关闭。
	// 自动处理的元素如下：
	// 1. 墨迹：移除不支持的墨迹（例如WPS墨迹）
	// 2. 自动翻页：移除PPT上所有自动翻页设置，并设置为单击鼠标翻页
	// 3. 已损坏音视频：移除PPT上对损坏音视频的引用
	AutoHandleUnsupportedElement *bool `json:"AutoHandleUnsupportedElement,omitnil" name:"AutoHandleUnsupportedElement"`
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
	delete(f, "AutoHandleUnsupportedElement")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDocumentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDocumentResponseParams struct {
	// 文档ID。
	DocumentId *string `json:"DocumentId,omitnil" name:"DocumentId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	GroupName *string `json:"GroupName,omitnil" name:"GroupName"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 默认绑定主讲老师ID
	TeacherId *string `json:"TeacherId,omitnil" name:"TeacherId"`

	// 群组成员列表,一次性最多200个
	MemberIds []*string `json:"MemberIds,omitnil" name:"MemberIds"`
}

type CreateGroupWithMembersRequest struct {
	*tchttp.BaseRequest
	
	// 待创建群组名称
	GroupName *string `json:"GroupName,omitnil" name:"GroupName"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 默认绑定主讲老师ID
	TeacherId *string `json:"TeacherId,omitnil" name:"TeacherId"`

	// 群组成员列表,一次性最多200个
	MemberIds []*string `json:"MemberIds,omitnil" name:"MemberIds"`
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
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	GroupName *string `json:"GroupName,omitnil" name:"GroupName"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 子群组ID列表，子群组ID不能重复，最多40个
	SubGroupIds []*string `json:"SubGroupIds,omitnil" name:"SubGroupIds"`

	// 群组默认主讲老师ID
	TeacherId *string `json:"TeacherId,omitnil" name:"TeacherId"`
}

type CreateGroupWithSubGroupRequest struct {
	*tchttp.BaseRequest
	
	// 待创建的新群组名
	GroupName *string `json:"GroupName,omitnil" name:"GroupName"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 子群组ID列表，子群组ID不能重复，最多40个
	SubGroupIds []*string `json:"SubGroupIds,omitnil" name:"SubGroupIds"`

	// 群组默认主讲老师ID
	TeacherId *string `json:"TeacherId,omitnil" name:"TeacherId"`
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
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Name *string `json:"Name,omitnil" name:"Name"`

	// 预定的房间开始时间，unix时间戳（秒）。
	StartTime *uint64 `json:"StartTime,omitnil" name:"StartTime"`

	// 预定的房间结束时间，unix时间戳（秒）。
	EndTime *uint64 `json:"EndTime,omitnil" name:"EndTime"`

	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 分辨率。可以有如下取值：
	// 1 标清
	// 2 高清
	// 3 全高清
	Resolution *uint64 `json:"Resolution,omitnil" name:"Resolution"`

	// 设置房间/课堂同时最大可与老师进行连麦互动的人数，该参数支持正式上课/开播前调用修改房间修改。取值范围[0,16]，当取值为0时表示当前课堂/直播，不支持连麦互动。
	MaxMicNumber *uint64 `json:"MaxMicNumber,omitnil" name:"MaxMicNumber"`

	// 房间子类型，可以有以下取值：
	// videodoc 文档+视频
	// video 纯视频
	SubType *string `json:"SubType,omitnil" name:"SubType"`

	// 老师ID。通过[注册用户]接口获取的UserId。指定后该用户在房间内拥有老师权限。
	TeacherId *string `json:"TeacherId,omitnil" name:"TeacherId"`

	// 进入课堂时是否自动连麦。可以有以下取值：
	// 0 不自动连麦（需要手动申请上麦，默认值）
	// 1 自动连麦
	AutoMic *uint64 `json:"AutoMic,omitnil" name:"AutoMic"`

	// 释放音视频权限后是否自动取消连麦。可以有以下取值：
	// 0 自动取消连麦（默认值）
	// 1 保持连麦状态
	TurnOffMic *uint64 `json:"TurnOffMic,omitnil" name:"TurnOffMic"`

	// 声音音质。可以有以下取值：
	// 0：流畅模式（默认值），占用更小的带宽、拥有更好的降噪效果，适用于1对1、小班教学、多人音视频会议等场景。
	// 1：高音质模式，适合需要高保真传输音乐的场景，但降噪效果会被削弱，适用于音乐教学场景。
	AudioQuality *uint64 `json:"AudioQuality,omitnil" name:"AudioQuality"`

	// 上课后是否禁止自动录制。可以有以下取值：
	// 0 不禁止录制（自动开启录制，默认值）
	// 1 禁止录制
	// 注：如果该配置取值为0，录制将从上课后开始，课堂结束后停止。
	DisableRecord *uint64 `json:"DisableRecord,omitnil" name:"DisableRecord"`

	// 助教Id列表。通过[注册用户]接口获取的UserId。指定后该用户在房间内拥有助教权限。
	Assistants []*string `json:"Assistants,omitnil" name:"Assistants"`

	// rtc人数。
	RTCAudienceNumber *uint64 `json:"RTCAudienceNumber,omitnil" name:"RTCAudienceNumber"`

	// 观看类型。互动观看 （默认）
	AudienceType *uint64 `json:"AudienceType,omitnil" name:"AudienceType"`

	// 录制模板。录制模板枚举值参考：https://cloud.tencent.com/document/product/1639/89744
	RecordLayout *uint64 `json:"RecordLayout,omitnil" name:"RecordLayout"`

	// 房间绑定的群组ID,非空时限制组成员进入
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// 是否允许老师/助教直接控制学生的摄像头/麦克风。可以有以下取值：
	// 0 不允许直接控制（需同意，默认值）
	// 1 允许直接控制（无需同意）
	EnableDirectControl *uint64 `json:"EnableDirectControl,omitnil" name:"EnableDirectControl"`

	// 开启专注模式。
	// 0 收看全部角色音视频(默认)
	// 1 只看老师和助教
	InteractionMode *int64 `json:"InteractionMode,omitnil" name:"InteractionMode"`

	// 横竖屏。0：横屏开播（默认值）; 1：竖屏开播，当前仅支持移动端的纯视频类型
	VideoOrientation *uint64 `json:"VideoOrientation,omitnil" name:"VideoOrientation"`

	// 开启课后评分。 0：不开启(默认)  1：开启
	IsGradingRequiredPostClass *int64 `json:"IsGradingRequiredPostClass,omitnil" name:"IsGradingRequiredPostClass"`

	// 房间类型: 0 小班课（默认值）; 1 大班课; 2 1V1 (预留参数，暂未开放)
	RoomType *int64 `json:"RoomType,omitnil" name:"RoomType"`

	// 拖堂时间：单位分钟，0为不限制(默认值), -1为不能拖堂，大于0为拖堂的时间，最大值120分钟
	EndDelayTime *int64 `json:"EndDelayTime,omitnil" name:"EndDelayTime"`

	// 直播方式：0 常规模式（默认）1 回放直播模式（伪直播）
	LiveType *uint64 `json:"LiveType,omitnil" name:"LiveType"`

	// 伪直播链接
	RecordLiveUrl *string `json:"RecordLiveUrl,omitnil" name:"RecordLiveUrl"`

	// 是否自动开始上课：0 不自动上课（默认） 1 自动上课 live_type=1的时候有效
	EnableAutoStart *uint64 `json:"EnableAutoStart,omitnil" name:"EnableAutoStart"`
}

type CreateRoomRequest struct {
	*tchttp.BaseRequest
	
	// 房间名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 预定的房间开始时间，unix时间戳（秒）。
	StartTime *uint64 `json:"StartTime,omitnil" name:"StartTime"`

	// 预定的房间结束时间，unix时间戳（秒）。
	EndTime *uint64 `json:"EndTime,omitnil" name:"EndTime"`

	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 分辨率。可以有如下取值：
	// 1 标清
	// 2 高清
	// 3 全高清
	Resolution *uint64 `json:"Resolution,omitnil" name:"Resolution"`

	// 设置房间/课堂同时最大可与老师进行连麦互动的人数，该参数支持正式上课/开播前调用修改房间修改。取值范围[0,16]，当取值为0时表示当前课堂/直播，不支持连麦互动。
	MaxMicNumber *uint64 `json:"MaxMicNumber,omitnil" name:"MaxMicNumber"`

	// 房间子类型，可以有以下取值：
	// videodoc 文档+视频
	// video 纯视频
	SubType *string `json:"SubType,omitnil" name:"SubType"`

	// 老师ID。通过[注册用户]接口获取的UserId。指定后该用户在房间内拥有老师权限。
	TeacherId *string `json:"TeacherId,omitnil" name:"TeacherId"`

	// 进入课堂时是否自动连麦。可以有以下取值：
	// 0 不自动连麦（需要手动申请上麦，默认值）
	// 1 自动连麦
	AutoMic *uint64 `json:"AutoMic,omitnil" name:"AutoMic"`

	// 释放音视频权限后是否自动取消连麦。可以有以下取值：
	// 0 自动取消连麦（默认值）
	// 1 保持连麦状态
	TurnOffMic *uint64 `json:"TurnOffMic,omitnil" name:"TurnOffMic"`

	// 声音音质。可以有以下取值：
	// 0：流畅模式（默认值），占用更小的带宽、拥有更好的降噪效果，适用于1对1、小班教学、多人音视频会议等场景。
	// 1：高音质模式，适合需要高保真传输音乐的场景，但降噪效果会被削弱，适用于音乐教学场景。
	AudioQuality *uint64 `json:"AudioQuality,omitnil" name:"AudioQuality"`

	// 上课后是否禁止自动录制。可以有以下取值：
	// 0 不禁止录制（自动开启录制，默认值）
	// 1 禁止录制
	// 注：如果该配置取值为0，录制将从上课后开始，课堂结束后停止。
	DisableRecord *uint64 `json:"DisableRecord,omitnil" name:"DisableRecord"`

	// 助教Id列表。通过[注册用户]接口获取的UserId。指定后该用户在房间内拥有助教权限。
	Assistants []*string `json:"Assistants,omitnil" name:"Assistants"`

	// rtc人数。
	RTCAudienceNumber *uint64 `json:"RTCAudienceNumber,omitnil" name:"RTCAudienceNumber"`

	// 观看类型。互动观看 （默认）
	AudienceType *uint64 `json:"AudienceType,omitnil" name:"AudienceType"`

	// 录制模板。录制模板枚举值参考：https://cloud.tencent.com/document/product/1639/89744
	RecordLayout *uint64 `json:"RecordLayout,omitnil" name:"RecordLayout"`

	// 房间绑定的群组ID,非空时限制组成员进入
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// 是否允许老师/助教直接控制学生的摄像头/麦克风。可以有以下取值：
	// 0 不允许直接控制（需同意，默认值）
	// 1 允许直接控制（无需同意）
	EnableDirectControl *uint64 `json:"EnableDirectControl,omitnil" name:"EnableDirectControl"`

	// 开启专注模式。
	// 0 收看全部角色音视频(默认)
	// 1 只看老师和助教
	InteractionMode *int64 `json:"InteractionMode,omitnil" name:"InteractionMode"`

	// 横竖屏。0：横屏开播（默认值）; 1：竖屏开播，当前仅支持移动端的纯视频类型
	VideoOrientation *uint64 `json:"VideoOrientation,omitnil" name:"VideoOrientation"`

	// 开启课后评分。 0：不开启(默认)  1：开启
	IsGradingRequiredPostClass *int64 `json:"IsGradingRequiredPostClass,omitnil" name:"IsGradingRequiredPostClass"`

	// 房间类型: 0 小班课（默认值）; 1 大班课; 2 1V1 (预留参数，暂未开放)
	RoomType *int64 `json:"RoomType,omitnil" name:"RoomType"`

	// 拖堂时间：单位分钟，0为不限制(默认值), -1为不能拖堂，大于0为拖堂的时间，最大值120分钟
	EndDelayTime *int64 `json:"EndDelayTime,omitnil" name:"EndDelayTime"`

	// 直播方式：0 常规模式（默认）1 回放直播模式（伪直播）
	LiveType *uint64 `json:"LiveType,omitnil" name:"LiveType"`

	// 伪直播链接
	RecordLiveUrl *string `json:"RecordLiveUrl,omitnil" name:"RecordLiveUrl"`

	// 是否自动开始上课：0 不自动上课（默认） 1 自动上课 live_type=1的时候有效
	EnableAutoStart *uint64 `json:"EnableAutoStart,omitnil" name:"EnableAutoStart"`
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
	delete(f, "TurnOffMic")
	delete(f, "AudioQuality")
	delete(f, "DisableRecord")
	delete(f, "Assistants")
	delete(f, "RTCAudienceNumber")
	delete(f, "AudienceType")
	delete(f, "RecordLayout")
	delete(f, "GroupId")
	delete(f, "EnableDirectControl")
	delete(f, "InteractionMode")
	delete(f, "VideoOrientation")
	delete(f, "IsGradingRequiredPostClass")
	delete(f, "RoomType")
	delete(f, "EndDelayTime")
	delete(f, "LiveType")
	delete(f, "RecordLiveUrl")
	delete(f, "EnableAutoStart")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRoomRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRoomResponseParams struct {
	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 用户ID列表。
	Users []*string `json:"Users,omitnil" name:"Users"`
}

type CreateSupervisorRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID。
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 用户ID列表。
	Users []*string `json:"Users,omitnil" name:"Users"`
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
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 指定需要删除的已设置的scene场景自定义元素，如果为空则删除应用下已设置的所有自定义元素。
	Scenes []*string `json:"Scenes,omitnil" name:"Scenes"`
}

type DeleteAppCustomContentRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID。
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 指定需要删除的已设置的scene场景自定义元素，如果为空则删除应用下已设置的所有自定义元素。
	Scenes []*string `json:"Scenes,omitnil" name:"Scenes"`
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
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	DocumentId *string `json:"DocumentId,omitnil" name:"DocumentId"`
}

type DeleteDocumentRequest struct {
	*tchttp.BaseRequest
	
	// 文档ID。
	DocumentId *string `json:"DocumentId,omitnil" name:"DocumentId"`
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
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 成员列表，最大值200
	MemberIds []*string `json:"MemberIds,omitnil" name:"MemberIds"`
}

type DeleteGroupMemberRequest struct {
	*tchttp.BaseRequest
	
	// 群组ID，联合群组无法删除群组成员
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 成员列表，最大值200
	MemberIds []*string `json:"MemberIds,omitnil" name:"MemberIds"`
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
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	GroupIds []*string `json:"GroupIds,omitnil" name:"GroupIds"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`
}

type DeleteGroupRequest struct {
	*tchttp.BaseRequest
	
	// 待删除群组ID列表
	GroupIds []*string `json:"GroupIds,omitnil" name:"GroupIds"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`
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
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	RoomId *int64 `json:"RoomId,omitnil" name:"RoomId"`

	// 低代码互动课堂的SdkAppId。
	SdkAppId *int64 `json:"SdkAppId,omitnil" name:"SdkAppId"`
}

type DeleteRecordRequest struct {
	*tchttp.BaseRequest
	
	// 房间Id。
	RoomId *int64 `json:"RoomId,omitnil" name:"RoomId"`

	// 低代码互动课堂的SdkAppId。
	SdkAppId *int64 `json:"SdkAppId,omitnil" name:"SdkAppId"`
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
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`
}

type DeleteRoomRequest struct {
	*tchttp.BaseRequest
	
	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`
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
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DeleteSupervisorRequestParams struct {
	// 应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 用户ID列表
	Users []*string `json:"Users,omitnil" name:"Users"`
}

type DeleteSupervisorRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 用户ID列表
	Users []*string `json:"Users,omitnil" name:"Users"`
}

func (r *DeleteSupervisorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSupervisorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Users")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSupervisorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSupervisorResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteSupervisorResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSupervisorResponseParams `json:"Response"`
}

func (r *DeleteSupervisorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSupervisorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserRequestParams struct {
	// 待删除用户的ID
	UserId *string `json:"UserId,omitnil" name:"UserId"`
}

type DeleteUserRequest struct {
	*tchttp.BaseRequest
	
	// 待删除用户的ID
	UserId *string `json:"UserId,omitnil" name:"UserId"`
}

func (r *DeleteUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteUserResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUserResponseParams `json:"Response"`
}

func (r *DeleteUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAnswerListRequestParams struct {
	// 问题ID
	QuestionId *string `json:"QuestionId,omitnil" name:"QuestionId"`

	// 1
	Page *int64 `json:"Page,omitnil" name:"Page"`

	// 100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeAnswerListRequest struct {
	*tchttp.BaseRequest
	
	// 问题ID
	QuestionId *string `json:"QuestionId,omitnil" name:"QuestionId"`

	// 1
	Page *int64 `json:"Page,omitnil" name:"Page"`

	// 100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeAnswerListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAnswerListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "QuestionId")
	delete(f, "Page")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAnswerListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAnswerListResponseParams struct {
	// 符合查询条件的房间答案总数
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 房间提问答案列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnswerInfo []*AnswerInfo `json:"AnswerInfo,omitnil" name:"AnswerInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAnswerListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAnswerListResponseParams `json:"Response"`
}

func (r *DescribeAnswerListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAnswerListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppDetailRequestParams struct {
	// 应用ID。低代码互动课堂的SdkAppId。
	ApplicationId *string `json:"ApplicationId,omitnil" name:"ApplicationId"`

	// 开发商ID
	DeveloperId *string `json:"DeveloperId,omitnil" name:"DeveloperId"`
}

type DescribeAppDetailRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID。低代码互动课堂的SdkAppId。
	ApplicationId *string `json:"ApplicationId,omitnil" name:"ApplicationId"`

	// 开发商ID
	DeveloperId *string `json:"DeveloperId,omitnil" name:"DeveloperId"`
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
	SdkAppId *string `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 应用配置
	AppConfig *AppConfig `json:"AppConfig,omitnil" name:"AppConfig"`

	// 场景配置
	SceneConfig []*SceneItem `json:"SceneConfig,omitnil" name:"SceneConfig"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`

	// 分页查询当前页数，从1开始递增。
	Page *uint64 `json:"Page,omitnil" name:"Page"`

	// 每页数据量，最大1000。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeCurrentMemberListRequest struct {
	*tchttp.BaseRequest
	
	// 房间Id。
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`

	// 分页查询当前页数，从1开始递增。
	Page *uint64 `json:"Page,omitnil" name:"Page"`

	// 每页数据量，最大1000。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
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
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 成员记录列表。
	MemberRecords []*MemberRecord `json:"MemberRecords,omitnil" name:"MemberRecords"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DescribeDeveloperRequestParams struct {

}

type DescribeDeveloperRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeDeveloperRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeveloperRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeveloperRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeveloperResponseParams struct {
	// 开发商ID
	DeveloperId *string `json:"DeveloperId,omitnil" name:"DeveloperId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDeveloperResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeveloperResponseParams `json:"Response"`
}

func (r *DescribeDeveloperResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeveloperResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDocumentRequestParams struct {
	// 文档Id（唯一id）
	DocumentId *string `json:"DocumentId,omitnil" name:"DocumentId"`
}

type DescribeDocumentRequest struct {
	*tchttp.BaseRequest
	
	// 文档Id（唯一id）
	DocumentId *string `json:"DocumentId,omitnil" name:"DocumentId"`
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
	DocumentId *string `json:"DocumentId,omitnil" name:"DocumentId"`

	// 文档原址url
	DocumentUrl *string `json:"DocumentUrl,omitnil" name:"DocumentUrl"`

	// 文档名称
	DocumentName *string `json:"DocumentName,omitnil" name:"DocumentName"`

	// 文档所有者UserId
	Owner *string `json:"Owner,omitnil" name:"Owner"`

	// 应用Id
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 文档权限
	Permission *uint64 `json:"Permission,omitnil" name:"Permission"`

	// 转码结果，无需转码为空，转码成功为结果url，转码失败为错误码
	TranscodeResult *string `json:"TranscodeResult,omitnil" name:"TranscodeResult"`

	// 转码类型
	TranscodeType *uint64 `json:"TranscodeType,omitnil" name:"TranscodeType"`

	// 转码进度， 0 - 100 表示（0% - 100%）
	TranscodeProgress *uint64 `json:"TranscodeProgress,omitnil" name:"TranscodeProgress"`

	// 转码状态，0为无需转码，1为正在转码，2为转码失败，3为转码成功
	TranscodeState *uint64 `json:"TranscodeState,omitnil" name:"TranscodeState"`

	// 转码失败后的错误信息
	TranscodeInfo *string `json:"TranscodeInfo,omitnil" name:"TranscodeInfo"`

	// 文档类型
	DocumentType *string `json:"DocumentType,omitnil" name:"DocumentType"`

	// 文档大小，单位：字节
	DocumentSize *uint64 `json:"DocumentSize,omitnil" name:"DocumentSize"`

	// 更新的UNIX时间戳
	UpdateTime *uint64 `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 课件页数
	Pages *uint64 `json:"Pages,omitnil" name:"Pages"`

	// 课件预览地址
	Preview *string `json:"Preview,omitnil" name:"Preview"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`

	// 低代码互动课堂的SdkAppId
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 分页查询当前页数，从1开始递增，默认值为1
	Page *int64 `json:"Page,omitnil" name:"Page"`

	// 每页数据量，最大1000，默认值为100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 课件权限。
	// [0]：获取owner的私有课件；
	// [1]：获取owner的公开课件;
	// [0,1]：则获取owner的私有课件和公开课件；
	// [2]：获取owner的私有课件和所有人(包括owner)的公开课件。
	// 默认值为[2]
	Permission []*uint64 `json:"Permission,omitnil" name:"Permission"`

	// 文档所有者的user_id，不填默认获取SdkAppId下所有课件
	Owner *string `json:"Owner,omitnil" name:"Owner"`
}

type DescribeDocumentsByRoomRequest struct {
	*tchttp.BaseRequest
	
	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`

	// 低代码互动课堂的SdkAppId
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 分页查询当前页数，从1开始递增，默认值为1
	Page *int64 `json:"Page,omitnil" name:"Page"`

	// 每页数据量，最大1000，默认值为100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 课件权限。
	// [0]：获取owner的私有课件；
	// [1]：获取owner的公开课件;
	// [0,1]：则获取owner的私有课件和公开课件；
	// [2]：获取owner的私有课件和所有人(包括owner)的公开课件。
	// 默认值为[2]
	Permission []*uint64 `json:"Permission,omitnil" name:"Permission"`

	// 文档所有者的user_id，不填默认获取SdkAppId下所有课件
	Owner *string `json:"Owner,omitnil" name:"Owner"`
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
	Documents []*DocumentInfo `json:"Documents,omitnil" name:"Documents"`

	// 符合查询条件文档总数
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DescribeDocumentsRequestParams struct {
	// 学校id
	SchoolId *uint64 `json:"SchoolId,omitnil" name:"SchoolId"`

	// 分页查询当前页数，从1开始递增
	Page *int64 `json:"Page,omitnil" name:"Page"`

	// 每页数据量，最大1000
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 课件权限。[0]：获取owner的私有课件；[1]：获取owner的公开课件; [0,1]：则获取owner的私有课件和公开课件；[2]：获取owner的私有课件和所有人(包括owner)的公开课件
	Permission []*uint64 `json:"Permission,omitnil" name:"Permission"`

	// 课件所有者的user_id，不填默认获取school_id下所有课件
	Owner *string `json:"Owner,omitnil" name:"Owner"`

	// 课件名称搜索词
	Keyword *string `json:"Keyword,omitnil" name:"Keyword"`

	// 课件id列表，从列表中查询，忽略错误的id
	DocumentId []*string `json:"DocumentId,omitnil" name:"DocumentId"`
}

type DescribeDocumentsRequest struct {
	*tchttp.BaseRequest
	
	// 学校id
	SchoolId *uint64 `json:"SchoolId,omitnil" name:"SchoolId"`

	// 分页查询当前页数，从1开始递增
	Page *int64 `json:"Page,omitnil" name:"Page"`

	// 每页数据量，最大1000
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 课件权限。[0]：获取owner的私有课件；[1]：获取owner的公开课件; [0,1]：则获取owner的私有课件和公开课件；[2]：获取owner的私有课件和所有人(包括owner)的公开课件
	Permission []*uint64 `json:"Permission,omitnil" name:"Permission"`

	// 课件所有者的user_id，不填默认获取school_id下所有课件
	Owner *string `json:"Owner,omitnil" name:"Owner"`

	// 课件名称搜索词
	Keyword *string `json:"Keyword,omitnil" name:"Keyword"`

	// 课件id列表，从列表中查询，忽略错误的id
	DocumentId []*string `json:"DocumentId,omitnil" name:"DocumentId"`
}

func (r *DescribeDocumentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDocumentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SchoolId")
	delete(f, "Page")
	delete(f, "Limit")
	delete(f, "Permission")
	delete(f, "Owner")
	delete(f, "Keyword")
	delete(f, "DocumentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDocumentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDocumentsResponseParams struct {
	// 符合查询条件文档总数
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 文档信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Documents []*DocumentInfo `json:"Documents,omitnil" name:"Documents"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDocumentsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDocumentsResponseParams `json:"Response"`
}

func (r *DescribeDocumentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDocumentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGroupListRequestParams struct {
	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 分页查询当前页数，默认从1开始递增。
	Page *uint64 `json:"Page,omitnil" name:"Page"`

	// 每页数据量，默认20，最大1000。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 主讲人ID筛选群组，与MemberId有且只有一个,都传时以此字段获取
	TeacherId *string `json:"TeacherId,omitnil" name:"TeacherId"`

	// 成员ID刷选群组，与TeacherId有且只有一个
	MemberId *string `json:"MemberId,omitnil" name:"MemberId"`
}

type DescribeGroupListRequest struct {
	*tchttp.BaseRequest
	
	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 分页查询当前页数，默认从1开始递增。
	Page *uint64 `json:"Page,omitnil" name:"Page"`

	// 每页数据量，默认20，最大1000。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 主讲人ID筛选群组，与MemberId有且只有一个,都传时以此字段获取
	TeacherId *string `json:"TeacherId,omitnil" name:"TeacherId"`

	// 成员ID刷选群组，与TeacherId有且只有一个
	MemberId *string `json:"MemberId,omitnil" name:"MemberId"`
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
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 群组信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupInfos []*GroupInfo `json:"GroupInfos,omitnil" name:"GroupInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 分页值，默认1
	Page *uint64 `json:"Page,omitnil" name:"Page"`

	// 每页数据量，默认20，最大1000
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeGroupMemberListRequest struct {
	*tchttp.BaseRequest
	
	// 群组ID
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 分页值，默认1
	Page *uint64 `json:"Page,omitnil" name:"Page"`

	// 每页数据量，默认20，最大1000
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
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
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 查询成员列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemberIds []*string `json:"MemberIds,omitnil" name:"MemberIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`
}

type DescribeGroupRequest struct {
	*tchttp.BaseRequest
	
	// 群组ID
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`
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
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// 群组名称
	GroupName *string `json:"GroupName,omitnil" name:"GroupName"`

	// 群主主讲人ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TeacherId *string `json:"TeacherId,omitnil" name:"TeacherId"`

	// 群组类型
	// 0-基础群组
	// 1-组合群组，若为1时会返回子群组ID
	GroupType *uint64 `json:"GroupType,omitnil" name:"GroupType"`

	// 子群组ID列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubGroupIds []*string `json:"SubGroupIds,omitnil" name:"SubGroupIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DescribeQuestionListRequestParams struct {
	// 房间ID
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`

	// 分页查询当前页数，从1开始递增，默认值为1
	Page *int64 `json:"Page,omitnil" name:"Page"`

	// 分页查询当前页数，从1开始递增，默认值为1
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeQuestionListRequest struct {
	*tchttp.BaseRequest
	
	// 房间ID
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`

	// 分页查询当前页数，从1开始递增，默认值为1
	Page *int64 `json:"Page,omitnil" name:"Page"`

	// 分页查询当前页数，从1开始递增，默认值为1
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeQuestionListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQuestionListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoomId")
	delete(f, "Page")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeQuestionListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQuestionListResponseParams struct {
	// 符合查询条件的房间问答问题总数
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 房间问答问题列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	QuestionInfo []*QuestionInfo `json:"QuestionInfo,omitnil" name:"QuestionInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeQuestionListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeQuestionListResponseParams `json:"Response"`
}

func (r *DescribeQuestionListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQuestionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoomForbiddenUserRequestParams struct {
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`
}

type DescribeRoomForbiddenUserRequest struct {
	*tchttp.BaseRequest
	
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`
}

func (r *DescribeRoomForbiddenUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoomForbiddenUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRoomForbiddenUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoomForbiddenUserResponseParams struct {
	// 禁言用户信息数组，内容包括被禁言的成员 ID，及其被禁言到的时间（使用 UTC 时间，即世界协调时间）
	// 注意：此字段可能返回 null，表示取不到有效值。
	MutedAccountList []*MutedAccountList `json:"MutedAccountList,omitnil" name:"MutedAccountList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRoomForbiddenUserResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRoomForbiddenUserResponseParams `json:"Response"`
}

func (r *DescribeRoomForbiddenUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoomForbiddenUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoomRequestParams struct {
	// 房间Id。
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`
}

type DescribeRoomRequest struct {
	*tchttp.BaseRequest
	
	// 房间Id。
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`
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
	Name *string `json:"Name,omitnil" name:"Name"`

	// 预定的房间开始时间，unix时间戳（秒）。
	StartTime *uint64 `json:"StartTime,omitnil" name:"StartTime"`

	// 预定的房间结束时间，unix时间戳（秒）。
	EndTime *uint64 `json:"EndTime,omitnil" name:"EndTime"`

	// 老师的UserId。
	TeacherId *string `json:"TeacherId,omitnil" name:"TeacherId"`

	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 观看类型。互动观看 （默认）	
	AudienceType *uint64 `json:"AudienceType,omitnil" name:"AudienceType"`

	// 分辨率。可以有如下取值：
	// 1 标清
	// 2 高清
	// 3 全高清
	Resolution *uint64 `json:"Resolution,omitnil" name:"Resolution"`

	// 最大连麦人数（不包括老师）。取值范围[0, 16]
	MaxMicNumber *uint64 `json:"MaxMicNumber,omitnil" name:"MaxMicNumber"`

	// 进入课堂时是否自动连麦。可以有以下取值：
	// 0 不自动连麦（需要手动申请上麦，默认值）
	// 1 自动连麦
	AutoMic *uint64 `json:"AutoMic,omitnil" name:"AutoMic"`

	// 高音质模式。可以有以下取值：
	// 0 不开启高音质（默认值）
	// 1 开启高音质
	AudioQuality *uint64 `json:"AudioQuality,omitnil" name:"AudioQuality"`

	// 房间子类型，可以有以下取值：
	// videodoc 文档+视频
	// video 纯视频
	SubType *string `json:"SubType,omitnil" name:"SubType"`

	// 上课后是否禁止自动录制。可以有以下取值：
	// 0 不禁止录制（自动开启录制，默认值）
	// 1 禁止录制
	// 注：如果该配置取值为0，录制将从上课后开始，课堂结束后停止。
	DisableRecord *uint64 `json:"DisableRecord,omitnil" name:"DisableRecord"`

	// 助教UserId列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Assistants []*string `json:"Assistants,omitnil" name:"Assistants"`

	// 录制地址（协议为https)。仅在房间结束后存在。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordUrl *string `json:"RecordUrl,omitnil" name:"RecordUrl"`

	// 课堂状态。0为未开始，1为已开始，2为已结束，3为已过期。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 房间绑定的群组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// 打开学生麦克风/摄像头的授权开关
	EnableDirectControl *uint64 `json:"EnableDirectControl,omitnil" name:"EnableDirectControl"`

	// 开启专注模式。
	// 0 收看全部角色音视频(默认)
	// 1 只看老师和助教
	InteractionMode *int64 `json:"InteractionMode,omitnil" name:"InteractionMode"`

	// 横竖屏。0：横屏开播（默认值）; 1：竖屏开播，当前仅支持移动端的纯视频类型
	VideoOrientation *uint64 `json:"VideoOrientation,omitnil" name:"VideoOrientation"`

	// 该房间是否开启了课后评分功能。0：未开启  1：开启
	IsGradingRequiredPostClass *int64 `json:"IsGradingRequiredPostClass,omitnil" name:"IsGradingRequiredPostClass"`

	// 房间类型: 0 小班课（默认值）; 1 大班课; 2 1V1 (后续扩展)
	RoomType *int64 `json:"RoomType,omitnil" name:"RoomType"`

	// 录制时长
	VideoDuration *uint64 `json:"VideoDuration,omitnil" name:"VideoDuration"`

	// 拖堂时间：单位分钟，0为不限制(默认值), -1为不能拖堂，大于0为拖堂的时间，最大值120分钟
	EndDelayTime *int64 `json:"EndDelayTime,omitnil" name:"EndDelayTime"`

	// 直播方式：0 常规模式（默认）1 回放直播模式（伪直播）
	LiveType *uint64 `json:"LiveType,omitnil" name:"LiveType"`

	// 伪直播链接
	RecordLiveUrl *string `json:"RecordLiveUrl,omitnil" name:"RecordLiveUrl"`

	// 是否自动开始上课：0 不自动上课（默认） 1 自动上课 live_type=1的时候有效
	EnableAutoStart *uint64 `json:"EnableAutoStart,omitnil" name:"EnableAutoStart"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`

	// 分页查询当前页数，从1开始递增。
	Page *uint64 `json:"Page,omitnil" name:"Page"`

	// 每页数据量，最大1000。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeRoomStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 房间Id。
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`

	// 分页查询当前页数，从1开始递增。
	Page *uint64 `json:"Page,omitnil" name:"Page"`

	// 每页数据量，最大1000。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
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
	PeakMemberNumber *uint64 `json:"PeakMemberNumber,omitnil" name:"PeakMemberNumber"`

	// 累计在线人数。
	MemberNumber *uint64 `json:"MemberNumber,omitnil" name:"MemberNumber"`

	// 记录总数。包含进入房间或者应到未到的。
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 成员记录列表。
	MemberRecords []*MemberRecord `json:"MemberRecords,omitnil" name:"MemberRecords"`

	// 秒级unix时间戳，实际房间开始时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealStartTime *uint64 `json:"RealStartTime,omitnil" name:"RealStartTime"`

	// 秒级unix时间戳，实际房间结束时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealEndTime *uint64 `json:"RealEndTime,omitnil" name:"RealEndTime"`

	// 房间消息总数。
	MessageCount *uint64 `json:"MessageCount,omitnil" name:"MessageCount"`

	// 房间连麦总数。
	MicCount *uint64 `json:"MicCount,omitnil" name:"MicCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DescribeScoreListRequestParams struct {
	// 课堂ID
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`

	// 分页查询当前页数，从1开始递增
	Page *uint64 `json:"Page,omitnil" name:"Page"`

	// 默认是10条
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeScoreListRequest struct {
	*tchttp.BaseRequest
	
	// 课堂ID
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`

	// 分页查询当前页数，从1开始递增
	Page *uint64 `json:"Page,omitnil" name:"Page"`

	// 默认是10条
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeScoreListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScoreListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoomId")
	delete(f, "Page")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScoreListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScoreListResponseParams struct {
	// 总数
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 课堂评分列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Scores []*ClassScoreItem `json:"Scores,omitnil" name:"Scores"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeScoreListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScoreListResponseParams `json:"Response"`
}

func (r *DescribeScoreListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScoreListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSdkAppIdUsersRequestParams struct {
	// 应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 分页，默认值为1
	Page *int64 `json:"Page,omitnil" name:"Page"`

	// 分页数据限制，默认值为20
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeSdkAppIdUsersRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 分页，默认值为1
	Page *int64 `json:"Page,omitnil" name:"Page"`

	// 分页数据限制，默认值为20
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
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
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 当前获取用户信息数组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Users []*UserInfo `json:"Users,omitnil" name:"Users"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DescribeSupervisorsRequestParams struct {
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 每页数据量，最大100。 不填默认20.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 分页查询当前页数，从1开始递增，不填默认为1。
	Page *uint64 `json:"Page,omitnil" name:"Page"`
}

type DescribeSupervisorsRequest struct {
	*tchttp.BaseRequest
	
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 每页数据量，最大100。 不填默认20.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 分页查询当前页数，从1开始递增，不填默认为1。
	Page *uint64 `json:"Page,omitnil" name:"Page"`
}

func (r *DescribeSupervisorsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSupervisorsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Limit")
	delete(f, "Page")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSupervisorsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSupervisorsResponseParams struct {
	// 数据总量
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 分页查询当前页数
	Page *uint64 `json:"Page,omitnil" name:"Page"`

	// 当前页数据量
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 巡课列表
	UserIds []*string `json:"UserIds,omitnil" name:"UserIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSupervisorsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSupervisorsResponseParams `json:"Response"`
}

func (r *DescribeSupervisorsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSupervisorsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserRequestParams struct {
	// 用户Id。
	UserId *string `json:"UserId,omitnil" name:"UserId"`
}

type DescribeUserRequest struct {
	*tchttp.BaseRequest
	
	// 用户Id。
	UserId *string `json:"UserId,omitnil" name:"UserId"`
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
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 用户Id。
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 用户昵称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 用户头像Url。
	Avatar *string `json:"Avatar,omitnil" name:"Avatar"`

	// 用户在客户系统的Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginId *string `json:"OriginId,omitnil" name:"OriginId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	DocumentId *string `json:"DocumentId,omitnil" name:"DocumentId"`

	// 文档原址url
	// 注意：此字段可能返回 null，表示取不到有效值。
	DocumentUrl *string `json:"DocumentUrl,omitnil" name:"DocumentUrl"`

	// 文档名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DocumentName *string `json:"DocumentName,omitnil" name:"DocumentName"`

	// 文档所有者UserId
	// 注意：此字段可能返回 null，表示取不到有效值。
	Owner *string `json:"Owner,omitnil" name:"Owner"`

	// 应用Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 文档权限，0：私有课件 1：公共课件
	// 注意：此字段可能返回 null，表示取不到有效值。
	Permission *uint64 `json:"Permission,omitnil" name:"Permission"`

	// 转码结果，无需转码为空，转码成功为结果url，转码失败为错误码
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranscodeResult *string `json:"TranscodeResult,omitnil" name:"TranscodeResult"`

	// 转码类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranscodeType *uint64 `json:"TranscodeType,omitnil" name:"TranscodeType"`

	// 转码进度， 0 - 100 表示（0% - 100%）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranscodeProgress *uint64 `json:"TranscodeProgress,omitnil" name:"TranscodeProgress"`

	// 转码状态，0为无需转码，1为正在转码，2为转码失败，3为转码成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranscodeState *uint64 `json:"TranscodeState,omitnil" name:"TranscodeState"`

	// 转码失败后的错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranscodeInfo *string `json:"TranscodeInfo,omitnil" name:"TranscodeInfo"`

	// 文档类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DocumentType *string `json:"DocumentType,omitnil" name:"DocumentType"`

	// 文档大小，单位：字节
	// 注意：此字段可能返回 null，表示取不到有效值。
	DocumentSize *uint64 `json:"DocumentSize,omitnil" name:"DocumentSize"`

	// 更新的UNIX时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *uint64 `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 课件页数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pages *uint64 `json:"Pages,omitnil" name:"Pages"`

	// 宽，仅在静态转码的课件有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	Width *uint64 `json:"Width,omitnil" name:"Width"`

	// 高，仅在静态转码的课件有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	Height *uint64 `json:"Height,omitnil" name:"Height"`

	// 封面，仅转码的课件会生成封面
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cover *string `json:"Cover,omitnil" name:"Cover"`

	// 课件预览地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Preview *string `json:"Preview,omitnil" name:"Preview"`
}

// Predefined struct for user
type EndRoomRequestParams struct {
	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`
}

type EndRoomRequest struct {
	*tchttp.BaseRequest
	
	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`
}

func (r *EndRoomRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EndRoomRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EndRoomRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EndRoomResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type EndRoomResponse struct {
	*tchttp.BaseResponse
	Response *EndRoomResponseParams `json:"Response"`
}

func (r *EndRoomResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EndRoomResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EventDataInfo struct {
	// 事件发生的房间号。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`

	// 事件发生的用户。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 用户设备类型。0: Unknown; 1: Windows; 2: macOS; 3: Android; 4: iOS; 5: Web; 6: Mobile webpage; 7: Weixin Mini Program.
	// 注意：此字段可能返回 null，表示取不到有效值。
	Device *uint64 `json:"Device,omitnil" name:"Device"`

	// 录制时长。单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *uint64 `json:"Duration,omitnil" name:"Duration"`

	// 录制文件大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordSize *uint64 `json:"RecordSize,omitnil" name:"RecordSize"`

	// 录制url
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordUrl *string `json:"RecordUrl,omitnil" name:"RecordUrl"`
}

type EventInfo struct {
	// 事件发生的秒级unix时间戳。
	Timestamp *uint64 `json:"Timestamp,omitnil" name:"Timestamp"`

	// 事件类型,有以下值:
	// RoomStart:房间开始 RoomEnd:房间结束 MemberJoin:成员加入 MemberQuit:成员退出 RecordFinish:录制结束
	// CameraOn: 摄像头打开
	// CameraOff: 摄像头关闭
	// MicOn: 麦克风打开
	// MicOff: 麦克风关闭
	// ScreenOn: 屏幕共享打开
	// ScreenOff: 屏幕共享关闭
	// VisibleOn: 页面可见
	// VisibleOff: 页面不可见
	EventType *string `json:"EventType,omitnil" name:"EventType"`

	// 事件详细内容，包含房间号,成员类型事件包含用户Id。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventData *EventDataInfo `json:"EventData,omitnil" name:"EventData"`
}

type FaceMsgContent struct {
	// 表情索引，用户自定义。
	Index *int64 `json:"Index,omitnil" name:"Index"`

	// 额外数据。
	Data *string `json:"Data,omitnil" name:"Data"`
}

// Predefined struct for user
type ForbidSendMsgRequestParams struct {
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`

	// 需要禁言的用户账号，最多支持500个账号
	MembersAccount []*string `json:"MembersAccount,omitnil" name:"MembersAccount"`

	// 需禁言时间，单位为秒，为0时表示取消禁言，4294967295为永久禁言。
	MuteTime *uint64 `json:"MuteTime,omitnil" name:"MuteTime"`
}

type ForbidSendMsgRequest struct {
	*tchttp.BaseRequest
	
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`

	// 需要禁言的用户账号，最多支持500个账号
	MembersAccount []*string `json:"MembersAccount,omitnil" name:"MembersAccount"`

	// 需禁言时间，单位为秒，为0时表示取消禁言，4294967295为永久禁言。
	MuteTime *uint64 `json:"MuteTime,omitnil" name:"MuteTime"`
}

func (r *ForbidSendMsgRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ForbidSendMsgRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	delete(f, "MembersAccount")
	delete(f, "MuteTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ForbidSendMsgRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ForbidSendMsgResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ForbidSendMsgResponse struct {
	*tchttp.BaseResponse
	Response *ForbidSendMsgResponseParams `json:"Response"`
}

func (r *ForbidSendMsgResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ForbidSendMsgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRoomEventRequestParams struct {
	// 房间Id。
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`

	// 应用Id。
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 起始页，1开始。keyword为空时有效。
	Page *uint64 `json:"Page,omitnil" name:"Page"`

	// 每页个数。keyword为空时有效。一次性最多200条。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 搜索事件类型。有以下事件类型:
	// RoomStart:房间开始
	// RoomEnd:房间结束
	// MemberJoin:成员加入
	// MemberQuit:成员退出
	// RecordFinish:录制结束
	// CameraOn: 摄像头打开
	// CameraOff: 摄像头关闭
	// MicOn: 麦克风打开
	// MicOff: 麦克风关闭
	// ScreenOn: 屏幕共享打开
	// ScreenOff: 屏幕共享关闭
	// VisibleOn: 页面可见
	// VisibleOff: 页面不可见
	Keyword *string `json:"Keyword,omitnil" name:"Keyword"`
}

type GetRoomEventRequest struct {
	*tchttp.BaseRequest
	
	// 房间Id。
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`

	// 应用Id。
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 起始页，1开始。keyword为空时有效。
	Page *uint64 `json:"Page,omitnil" name:"Page"`

	// 每页个数。keyword为空时有效。一次性最多200条。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 搜索事件类型。有以下事件类型:
	// RoomStart:房间开始
	// RoomEnd:房间结束
	// MemberJoin:成员加入
	// MemberQuit:成员退出
	// RecordFinish:录制结束
	// CameraOn: 摄像头打开
	// CameraOff: 摄像头关闭
	// MicOn: 麦克风打开
	// MicOff: 麦克风关闭
	// ScreenOn: 屏幕共享打开
	// ScreenOff: 屏幕共享关闭
	// VisibleOn: 页面可见
	// VisibleOff: 页面不可见
	Keyword *string `json:"Keyword,omitnil" name:"Keyword"`
}

func (r *GetRoomEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRoomEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoomId")
	delete(f, "SdkAppId")
	delete(f, "Page")
	delete(f, "Limit")
	delete(f, "Keyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetRoomEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRoomEventResponseParams struct {
	// 该房间的事件总数，keyword搜索不影响该值。
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 详细事件内容。包含相应的类型、发生的时间戳。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Events []*EventInfo `json:"Events,omitnil" name:"Events"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetRoomEventResponse struct {
	*tchttp.BaseResponse
	Response *GetRoomEventResponseParams `json:"Response"`
}

func (r *GetRoomEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRoomEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRoomMessageRequestParams struct {
	// 低代码互动课堂的SdkAppId。
	SdkAppId *int64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 房间Id。	
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`

	// 消息序列。获取该序列以前的消息(不包含该seq消息)
	Seq *int64 `json:"Seq,omitnil" name:"Seq"`

	// 消息拉取的条数。最大数量不能超过套餐包限制。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type GetRoomMessageRequest struct {
	*tchttp.BaseRequest
	
	// 低代码互动课堂的SdkAppId。
	SdkAppId *int64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 房间Id。	
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`

	// 消息序列。获取该序列以前的消息(不包含该seq消息)
	Seq *int64 `json:"Seq,omitnil" name:"Seq"`

	// 消息拉取的条数。最大数量不能超过套餐包限制。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
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
	Messages []*MessageList `json:"Messages,omitnil" name:"Messages"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type GetRoomsRequestParams struct {
	// 低代码平台的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 开始时间。默认以当前时间减去半小时作为开始时间。
	StartTime *uint64 `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间。默认以当前时间加上半小时作为结束时间。
	EndTime *uint64 `json:"EndTime,omitnil" name:"EndTime"`

	// 分页查询当前页数，从1开始递增
	Page *uint64 `json:"Page,omitnil" name:"Page"`

	// 默认是10条
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 课堂状态。默认展示所有课堂，0为未开始，1为正在上课，2为已结束，3为已过期
	Status []*uint64 `json:"Status,omitnil" name:"Status"`
}

type GetRoomsRequest struct {
	*tchttp.BaseRequest
	
	// 低代码平台的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 开始时间。默认以当前时间减去半小时作为开始时间。
	StartTime *uint64 `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间。默认以当前时间加上半小时作为结束时间。
	EndTime *uint64 `json:"EndTime,omitnil" name:"EndTime"`

	// 分页查询当前页数，从1开始递增
	Page *uint64 `json:"Page,omitnil" name:"Page"`

	// 默认是10条
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 课堂状态。默认展示所有课堂，0为未开始，1为正在上课，2为已结束，3为已过期
	Status []*uint64 `json:"Status,omitnil" name:"Status"`
}

func (r *GetRoomsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRoomsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Page")
	delete(f, "Limit")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetRoomsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRoomsResponseParams struct {
	// 总数
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 房间列表
	Rooms []*RoomItem `json:"Rooms,omitnil" name:"Rooms"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetRoomsResponse struct {
	*tchttp.BaseResponse
	Response *GetRoomsResponseParams `json:"Response"`
}

func (r *GetRoomsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRoomsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetWatermarkRequestParams struct {
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`
}

type GetWatermarkRequest struct {
	*tchttp.BaseRequest
	
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`
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
	TeacherLogo *WatermarkConfig `json:"TeacherLogo,omitnil" name:"TeacherLogo"`

	// 白板区域的水印参数配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	BoardLogo *WatermarkConfig `json:"BoardLogo,omitnil" name:"BoardLogo"`

	// 背景图片配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackgroundPicture *BackgroundPictureConfig `json:"BackgroundPicture,omitnil" name:"BackgroundPicture"`

	// 文字水印配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *TextMarkConfig `json:"Text,omitnil" name:"Text"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	GroupName *string `json:"GroupName,omitnil" name:"GroupName"`

	// 群组主讲人ID
	TeacherId *string `json:"TeacherId,omitnil" name:"TeacherId"`
}

type GroupInfo struct {
	// 群组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// 群组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitnil" name:"GroupName"`

	// 群组主讲人ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TeacherId *string `json:"TeacherId,omitnil" name:"TeacherId"`

	// 群组类型 
	// 0-基础群组 
	// 1-组合群组，若为1时会返回子群组ID列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupType *uint64 `json:"GroupType,omitnil" name:"GroupType"`

	// 子群组ID列表，如有。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubGroupIds *string `json:"SubGroupIds,omitnil" name:"SubGroupIds"`
}

type ImageInfo struct {
	// 图片类型：
	// 1-原图
	// 2-大图
	// 3-缩略图
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// 图片数据大小，单位：字节。
	Size *uint64 `json:"Size,omitnil" name:"Size"`

	// 图片宽度，单位为像素。
	Width *uint64 `json:"Width,omitnil" name:"Width"`

	// 图片高度，单位为像素。
	Height *uint64 `json:"Height,omitnil" name:"Height"`

	// 图片下载地址。
	URL *string `json:"URL,omitnil" name:"URL"`
}

type ImageMsgContent struct {
	// 图片的唯一标识，客户端用于索引图片的键值。
	UUID *string `json:"UUID,omitnil" name:"UUID"`

	// 图片格式。
	// JPG = 1
	// GIF = 2
	// PNG = 3
	// BMP = 4
	// 其他 = 255
	ImageFormat *uint64 `json:"ImageFormat,omitnil" name:"ImageFormat"`

	// 图片信息
	ImageInfoList []*ImageInfo `json:"ImageInfoList,omitnil" name:"ImageInfoList"`
}

// Predefined struct for user
type KickUserFromRoomRequestParams struct {
	// 房间Id。
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`

	// 低代码平台的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 需要踢出成员Id
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 踢出类型：
	// 1：临时踢出，可以使用Duration参数指定污点时间，污点时间间隔内用户无法进入房间。
	// 2：永久踢出
	KickType *uint64 `json:"KickType,omitnil" name:"KickType"`

	// 污点时间(单位秒)，KickType = 1时生效，默认为0
	Duration *uint64 `json:"Duration,omitnil" name:"Duration"`
}

type KickUserFromRoomRequest struct {
	*tchttp.BaseRequest
	
	// 房间Id。
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`

	// 低代码平台的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 需要踢出成员Id
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 踢出类型：
	// 1：临时踢出，可以使用Duration参数指定污点时间，污点时间间隔内用户无法进入房间。
	// 2：永久踢出
	KickType *uint64 `json:"KickType,omitnil" name:"KickType"`

	// 污点时间(单位秒)，KickType = 1时生效，默认为0
	Duration *uint64 `json:"Duration,omitnil" name:"Duration"`
}

func (r *KickUserFromRoomRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *KickUserFromRoomRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoomId")
	delete(f, "SdkAppId")
	delete(f, "UserId")
	delete(f, "KickType")
	delete(f, "Duration")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "KickUserFromRoomRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type KickUserFromRoomResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type KickUserFromRoomResponse struct {
	*tchttp.BaseResponse
	Response *KickUserFromRoomResponseParams `json:"Response"`
}

func (r *KickUserFromRoomResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *KickUserFromRoomResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type LoginOriginIdRequestParams struct {
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 用户在客户系统的Id，需要在同一应用下唯一。
	OriginId *string `json:"OriginId,omitnil" name:"OriginId"`
}

type LoginOriginIdRequest struct {
	*tchttp.BaseRequest
	
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 用户在客户系统的Id，需要在同一应用下唯一。
	OriginId *string `json:"OriginId,omitnil" name:"OriginId"`
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
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 登录/注册成功后返回登录态token。有效期7天。
	Token *string `json:"Token,omitnil" name:"Token"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	UserId *string `json:"UserId,omitnil" name:"UserId"`
}

type LoginUserRequest struct {
	*tchttp.BaseRequest
	
	// 注册获取的用户id。
	UserId *string `json:"UserId,omitnil" name:"UserId"`
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
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 注册成功后返回登录态token，有效期7天。token过期后可以通过调用“登录”或“源账号登录”进行更新。
	Token *string `json:"Token,omitnil" name:"Token"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 用户名称。
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 在线时长，单位秒。
	PresentTime *uint64 `json:"PresentTime,omitnil" name:"PresentTime"`

	// 是否开启摄像头。
	Camera *uint64 `json:"Camera,omitnil" name:"Camera"`

	// 是否开启麦克风。
	Mic *uint64 `json:"Mic,omitnil" name:"Mic"`

	// 是否禁言。
	Silence *uint64 `json:"Silence,omitnil" name:"Silence"`

	// 回答问题数量。
	AnswerQuestions *uint64 `json:"AnswerQuestions,omitnil" name:"AnswerQuestions"`

	// 举手数量。
	HandUps *uint64 `json:"HandUps,omitnil" name:"HandUps"`

	// 首次进入房间的unix时间戳。
	FirstJoinTimestamp *uint64 `json:"FirstJoinTimestamp,omitnil" name:"FirstJoinTimestamp"`

	// 最后一次退出房间的unix时间戳。
	LastQuitTimestamp *uint64 `json:"LastQuitTimestamp,omitnil" name:"LastQuitTimestamp"`

	// 奖励次数。
	Rewords *uint64 `json:"Rewords,omitnil" name:"Rewords"`

	// 用户IP。
	IPAddress *string `json:"IPAddress,omitnil" name:"IPAddress"`

	// 用户位置信息。
	Location *string `json:"Location,omitnil" name:"Location"`

	// 用户设备平台信息。0:unknown  1:windows  2:mac  3:android  4:ios  5:web   6:h5   7:miniprogram （小程序）
	Device *int64 `json:"Device,omitnil" name:"Device"`

	// 每个成员上麦次数。
	PerMemberMicCount *int64 `json:"PerMemberMicCount,omitnil" name:"PerMemberMicCount"`

	// 每个成员发送消息数量。
	PerMemberMessageCount *int64 `json:"PerMemberMessageCount,omitnil" name:"PerMemberMessageCount"`

	// 用户角色。0代表学生；1代表老师； 2助教；3巡课。
	Role *uint64 `json:"Role,omitnil" name:"Role"`

	// 上课班号
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// 子上课班号
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubGroupId []*string `json:"SubGroupId,omitnil" name:"SubGroupId"`

	// 用户的上台状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Stage *int64 `json:"Stage,omitnil" name:"Stage"`

	// 用户状态。0为未到，1为在线，2为离线，3为被踢，4为永久被踢，5为暂时掉线
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentState *uint64 `json:"CurrentState,omitnil" name:"CurrentState"`
}

type MessageItem struct {
	// 消息类型。0表示文本消息，1表示图片消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MessageType *int64 `json:"MessageType,omitnil" name:"MessageType"`

	// 文本消息内容。message type为0时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TextMessage *string `json:"TextMessage,omitnil" name:"TextMessage"`

	// 图片消息URL。 message type为1时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageMessage *string `json:"ImageMessage,omitnil" name:"ImageMessage"`
}

type MessageList struct {
	// 消息时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *int64 `json:"Timestamp,omitnil" name:"Timestamp"`

	// 消息发送者
	// 注意：此字段可能返回 null，表示取不到有效值。
	FromAccount *string `json:"FromAccount,omitnil" name:"FromAccount"`

	// 消息序列号，当前课堂内唯一且单调递增
	// 注意：此字段可能返回 null，表示取不到有效值。
	Seq *int64 `json:"Seq,omitnil" name:"Seq"`

	// 历史消息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	MessageBody []*MessageItem `json:"MessageBody,omitnil" name:"MessageBody"`
}

// Predefined struct for user
type ModifyAppRequestParams struct {
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 回调地址。
	Callback *string `json:"Callback,omitnil" name:"Callback"`

	// 回调key。
	CallbackKey *string `json:"CallbackKey,omitnil" name:"CallbackKey"`
}

type ModifyAppRequest struct {
	*tchttp.BaseRequest
	
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 回调地址。
	Callback *string `json:"Callback,omitnil" name:"Callback"`

	// 回调key。
	CallbackKey *string `json:"CallbackKey,omitnil" name:"CallbackKey"`
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
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 默认绑定主讲老师ID
	TeacherId *string `json:"TeacherId,omitnil" name:"TeacherId"`

	// 待修改的群组名称
	GroupName *string `json:"GroupName,omitnil" name:"GroupName"`
}

type ModifyGroupRequest struct {
	*tchttp.BaseRequest
	
	// 需要修改的群组ID
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 默认绑定主讲老师ID
	TeacherId *string `json:"TeacherId,omitnil" name:"TeacherId"`

	// 待修改的群组名称
	GroupName *string `json:"GroupName,omitnil" name:"GroupName"`
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
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`

	// 低代码互动课堂的SdkAppId
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 预定的房间开始时间，unix时间戳（秒）。直播开始后不允许修改。
	StartTime *uint64 `json:"StartTime,omitnil" name:"StartTime"`

	// 预定的房间结束时间，unix时间戳（秒）。直播开始后不允许修改。
	EndTime *uint64 `json:"EndTime,omitnil" name:"EndTime"`

	// 老师ID。直播开始后不允许修改。
	TeacherId *string `json:"TeacherId,omitnil" name:"TeacherId"`

	// 房间名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 分辨率。可以有如下取值：
	// 1 标清
	// 2 高清
	// 3 全高清
	// 直播开始后不允许修改。
	Resolution *uint64 `json:"Resolution,omitnil" name:"Resolution"`

	// 设置房间/课堂同时最大可与老师进行连麦互动的人数，该参数支持正式上课/开播前调用修改房间修改。
	// 取值范围[0,16]，当取值为0时表示当前课堂/直播，不支持连麦互动。
	MaxMicNumber *uint64 `json:"MaxMicNumber,omitnil" name:"MaxMicNumber"`

	// 进入房间时是否自动连麦。可以有以下取值：
	// 0 不自动连麦（默认值）
	// 1 自动连麦
	// 直播开始后不允许修改。
	AutoMic *uint64 `json:"AutoMic,omitnil" name:"AutoMic"`

	// 高音质模式。可以有以下取值：
	// 0 不开启高音质（默认值）
	// 1 开启高音质
	// 直播开始后不允许修改。
	AudioQuality *uint64 `json:"AudioQuality,omitnil" name:"AudioQuality"`

	// 房间子类型，可以有以下取值：
	// videodoc 文档+视频
	// video 纯视频
	// 直播开始后不允许修改。
	SubType *string `json:"SubType,omitnil" name:"SubType"`

	// 禁止录制。可以有以下取值：
	// 0 不禁止录制（默认值）
	// 1 禁止录制
	// 直播开始后不允许修改。
	DisableRecord *uint64 `json:"DisableRecord,omitnil" name:"DisableRecord"`

	// 助教Id列表。直播开始后不允许修改。
	Assistants []*string `json:"Assistants,omitnil" name:"Assistants"`

	// 房间绑定的群组ID。直播开始后不允许修改。
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// 打开学生麦克风/摄像头的授权开关。直播开始后不允许修改。
	EnableDirectControl *uint64 `json:"EnableDirectControl,omitnil" name:"EnableDirectControl"`

	// 开启专注模式。
	// 0 收看全部角色音视频(默认)
	// 1 只看老师和助教
	InteractionMode *uint64 `json:"InteractionMode,omitnil" name:"InteractionMode"`

	// 横竖屏。0：横屏开播（默认值）; 1：竖屏开播，当前仅支持移动端的纯视频类型
	VideoOrientation *uint64 `json:"VideoOrientation,omitnil" name:"VideoOrientation"`

	// 开启课后评分。 0：不开启(默认)  1：开启
	IsGradingRequiredPostClass *uint64 `json:"IsGradingRequiredPostClass,omitnil" name:"IsGradingRequiredPostClass"`

	// 房间类型: 0 小班课（默认值）; 1 大班课; 2 1V1 （预留参数、暂未开放)
	RoomType *uint64 `json:"RoomType,omitnil" name:"RoomType"`

	// 录制模板。仅可修改还未开始的房间。录制模板枚举值参考：https://cloud.tencent.com/document/product/1639/89744
	RecordLayout *uint64 `json:"RecordLayout,omitnil" name:"RecordLayout"`

	// 拖堂时间：单位分钟，0为不限制(默认值), -1为不能拖堂，大于0为拖堂的时间，最大值120分钟
	EndDelayTime *int64 `json:"EndDelayTime,omitnil" name:"EndDelayTime"`

	// 直播方式：0 常规模式（默认）1 回放直播模式（伪直播）
	LiveType *uint64 `json:"LiveType,omitnil" name:"LiveType"`

	// 伪直播链接
	RecordLiveUrl *string `json:"RecordLiveUrl,omitnil" name:"RecordLiveUrl"`

	// 是否自动开始上课：0 不自动上课（默认） 1 自动上课 live_type=1的时候有效
	EnableAutoStart *uint64 `json:"EnableAutoStart,omitnil" name:"EnableAutoStart"`
}

type ModifyRoomRequest struct {
	*tchttp.BaseRequest
	
	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`

	// 低代码互动课堂的SdkAppId
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 预定的房间开始时间，unix时间戳（秒）。直播开始后不允许修改。
	StartTime *uint64 `json:"StartTime,omitnil" name:"StartTime"`

	// 预定的房间结束时间，unix时间戳（秒）。直播开始后不允许修改。
	EndTime *uint64 `json:"EndTime,omitnil" name:"EndTime"`

	// 老师ID。直播开始后不允许修改。
	TeacherId *string `json:"TeacherId,omitnil" name:"TeacherId"`

	// 房间名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 分辨率。可以有如下取值：
	// 1 标清
	// 2 高清
	// 3 全高清
	// 直播开始后不允许修改。
	Resolution *uint64 `json:"Resolution,omitnil" name:"Resolution"`

	// 设置房间/课堂同时最大可与老师进行连麦互动的人数，该参数支持正式上课/开播前调用修改房间修改。
	// 取值范围[0,16]，当取值为0时表示当前课堂/直播，不支持连麦互动。
	MaxMicNumber *uint64 `json:"MaxMicNumber,omitnil" name:"MaxMicNumber"`

	// 进入房间时是否自动连麦。可以有以下取值：
	// 0 不自动连麦（默认值）
	// 1 自动连麦
	// 直播开始后不允许修改。
	AutoMic *uint64 `json:"AutoMic,omitnil" name:"AutoMic"`

	// 高音质模式。可以有以下取值：
	// 0 不开启高音质（默认值）
	// 1 开启高音质
	// 直播开始后不允许修改。
	AudioQuality *uint64 `json:"AudioQuality,omitnil" name:"AudioQuality"`

	// 房间子类型，可以有以下取值：
	// videodoc 文档+视频
	// video 纯视频
	// 直播开始后不允许修改。
	SubType *string `json:"SubType,omitnil" name:"SubType"`

	// 禁止录制。可以有以下取值：
	// 0 不禁止录制（默认值）
	// 1 禁止录制
	// 直播开始后不允许修改。
	DisableRecord *uint64 `json:"DisableRecord,omitnil" name:"DisableRecord"`

	// 助教Id列表。直播开始后不允许修改。
	Assistants []*string `json:"Assistants,omitnil" name:"Assistants"`

	// 房间绑定的群组ID。直播开始后不允许修改。
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// 打开学生麦克风/摄像头的授权开关。直播开始后不允许修改。
	EnableDirectControl *uint64 `json:"EnableDirectControl,omitnil" name:"EnableDirectControl"`

	// 开启专注模式。
	// 0 收看全部角色音视频(默认)
	// 1 只看老师和助教
	InteractionMode *uint64 `json:"InteractionMode,omitnil" name:"InteractionMode"`

	// 横竖屏。0：横屏开播（默认值）; 1：竖屏开播，当前仅支持移动端的纯视频类型
	VideoOrientation *uint64 `json:"VideoOrientation,omitnil" name:"VideoOrientation"`

	// 开启课后评分。 0：不开启(默认)  1：开启
	IsGradingRequiredPostClass *uint64 `json:"IsGradingRequiredPostClass,omitnil" name:"IsGradingRequiredPostClass"`

	// 房间类型: 0 小班课（默认值）; 1 大班课; 2 1V1 （预留参数、暂未开放)
	RoomType *uint64 `json:"RoomType,omitnil" name:"RoomType"`

	// 录制模板。仅可修改还未开始的房间。录制模板枚举值参考：https://cloud.tencent.com/document/product/1639/89744
	RecordLayout *uint64 `json:"RecordLayout,omitnil" name:"RecordLayout"`

	// 拖堂时间：单位分钟，0为不限制(默认值), -1为不能拖堂，大于0为拖堂的时间，最大值120分钟
	EndDelayTime *int64 `json:"EndDelayTime,omitnil" name:"EndDelayTime"`

	// 直播方式：0 常规模式（默认）1 回放直播模式（伪直播）
	LiveType *uint64 `json:"LiveType,omitnil" name:"LiveType"`

	// 伪直播链接
	RecordLiveUrl *string `json:"RecordLiveUrl,omitnil" name:"RecordLiveUrl"`

	// 是否自动开始上课：0 不自动上课（默认） 1 自动上课 live_type=1的时候有效
	EnableAutoStart *uint64 `json:"EnableAutoStart,omitnil" name:"EnableAutoStart"`
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
	delete(f, "EnableDirectControl")
	delete(f, "InteractionMode")
	delete(f, "VideoOrientation")
	delete(f, "IsGradingRequiredPostClass")
	delete(f, "RoomType")
	delete(f, "RecordLayout")
	delete(f, "EndDelayTime")
	delete(f, "LiveType")
	delete(f, "RecordLiveUrl")
	delete(f, "EnableAutoStart")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRoomRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRoomResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 待修改的用户名
	Nickname *string `json:"Nickname,omitnil" name:"Nickname"`

	// 待修改头像url
	Avatar *string `json:"Avatar,omitnil" name:"Avatar"`
}

type ModifyUserProfileRequest struct {
	*tchttp.BaseRequest
	
	// 待修改用户ID
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 待修改的用户名
	Nickname *string `json:"Nickname,omitnil" name:"Nickname"`

	// 待修改头像url
	Avatar *string `json:"Avatar,omitnil" name:"Avatar"`
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
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

type MsgBody struct {
	// TIM 消息对象类型，目前支持的消息对象包括：
	// TIMTextElem（文本消息）
	// TIMFaceElem（表情消息）
	// TIMImageElem（图像消息）
	MsgType *string `json:"MsgType,omitnil" name:"MsgType"`

	// 文本消息，当MsgType 为TIMTextElem（文本消息）必选。
	TextMsgContent *TextMsgContent `json:"TextMsgContent,omitnil" name:"TextMsgContent"`

	// 表情消息，当MsgType 为TIMFaceElem（表情消息）必选。
	FaceMsgContent *FaceMsgContent `json:"FaceMsgContent,omitnil" name:"FaceMsgContent"`

	// 图像消息，当MsgType为TIMImageElem（图像消息）必选。
	ImageMsgContent *ImageMsgContent `json:"ImageMsgContent,omitnil" name:"ImageMsgContent"`
}

type MutedAccountList struct {
	// 用户 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemberAccount *string `json:"MemberAccount,omitnil" name:"MemberAccount"`

	// 禁言到的时间（使用 UTC 时间，即世界协调时间）
	// 注意：此字段可能返回 null，表示取不到有效值。
	MutedUntil *uint64 `json:"MutedUntil,omitnil" name:"MutedUntil"`
}

type QuestionInfo struct {
	// 问题ID
	QuestionId *string `json:"QuestionId,omitnil" name:"QuestionId"`

	// 问题内容
	QuestionContent *string `json:"QuestionContent,omitnil" name:"QuestionContent"`

	// 倒计时答题设置的秒数（0 表示不计时）
	Duration *uint64 `json:"Duration,omitnil" name:"Duration"`

	// 正确答案（按照位表示是否选择，如0x1表示选择A，0x11表示选择AB）
	CorrectAnswer *int64 `json:"CorrectAnswer,omitnil" name:"CorrectAnswer"`

	// 每个选项答题人数统计
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnswerStats []*AnswerStat `json:"AnswerStats,omitnil" name:"AnswerStats"`
}

// Predefined struct for user
type RegisterUserRequestParams struct {
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 用户名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 用户在客户系统的Id，需要在同一应用下唯一。
	OriginId *string `json:"OriginId,omitnil" name:"OriginId"`

	// 用户头像。
	Avatar *string `json:"Avatar,omitnil" name:"Avatar"`
}

type RegisterUserRequest struct {
	*tchttp.BaseRequest
	
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 用户名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 用户在客户系统的Id，需要在同一应用下唯一。
	OriginId *string `json:"OriginId,omitnil" name:"OriginId"`

	// 用户头像。
	Avatar *string `json:"Avatar,omitnil" name:"Avatar"`
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
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 登录/注册成功后返回登录态token。有效期7天。
	Token *string `json:"Token,omitnil" name:"Token"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Name *string `json:"Name,omitnil" name:"Name"`

	// 预定的房间开始时间，unix时间戳。
	StartTime *uint64 `json:"StartTime,omitnil" name:"StartTime"`

	// 预定的房间结束时间，unix时间戳。
	EndTime *uint64 `json:"EndTime,omitnil" name:"EndTime"`

	// 分辨率。可以有如下取值： 1 标清 2 高清 3 全高清
	Resolution *uint64 `json:"Resolution,omitnil" name:"Resolution"`

	// 最大连麦人数（不包括老师）。取值范围[0, 16]
	MaxMicNumber *uint64 `json:"MaxMicNumber,omitnil" name:"MaxMicNumber"`

	// 房间子类型，可以有以下取值： videodoc 文档+视频 video 纯视频
	SubType *string `json:"SubType,omitnil" name:"SubType"`

	// 老师ID。通过[注册用户]接口获取的UserId。
	TeacherId *string `json:"TeacherId,omitnil" name:"TeacherId"`

	// 进入课堂时是否自动连麦。可以有以下取值： 0 不自动连麦（需要手动申请上麦，默认值） 1 自动连麦
	AutoMic *uint64 `json:"AutoMic,omitnil" name:"AutoMic"`

	// 释放音视频权限后是否自动取消连麦。可以有以下取值： 0 自动取消连麦（默认值） 1 保持连麦状态
	TurnOffMic *uint64 `json:"TurnOffMic,omitnil" name:"TurnOffMic"`

	// 高音质模式。可以有以下取值： 0 不开启高音质（默认值） 1 开启高音质
	AudioQuality *uint64 `json:"AudioQuality,omitnil" name:"AudioQuality"`

	// 上课后是否禁止自动录制。可以有以下取值： 0 不禁止录制（自动开启录制，默认值） 1 禁止录制 注：如果该配置取值为0，录制将从上课后开始，课堂结束后停止。
	DisableRecord *uint64 `json:"DisableRecord,omitnil" name:"DisableRecord"`

	// 助教Id列表。通过[注册用户]接口获取的UserId。
	Assistants []*string `json:"Assistants,omitnil" name:"Assistants"`

	// rtc人数。
	RTCAudienceNumber *uint64 `json:"RTCAudienceNumber,omitnil" name:"RTCAudienceNumber"`

	// 观看类型。
	AudienceType *uint64 `json:"AudienceType,omitnil" name:"AudienceType"`

	// 录制布局。
	RecordLayout *uint64 `json:"RecordLayout,omitnil" name:"RecordLayout"`

	// 房间绑定的群组ID
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// 打开学生麦克风/摄像头的授权开关
	EnableDirectControl *uint64 `json:"EnableDirectControl,omitnil" name:"EnableDirectControl"`

	// 开启专注模式。 0 收看全部角色音视频(默认) 1 只看老师和助教
	InteractionMode *int64 `json:"InteractionMode,omitnil" name:"InteractionMode"`

	// 横竖屏。0：横屏开播（默认值）; 1：竖屏开播，当前仅支持移动端的纯视频类型
	VideoOrientation *int64 `json:"VideoOrientation,omitnil" name:"VideoOrientation"`

	// 开启课后评分。 0：不开启(默认)  1：开启
	IsGradingRequiredPostClass *int64 `json:"IsGradingRequiredPostClass,omitnil" name:"IsGradingRequiredPostClass"`

	// 房间类型: 0 小班课（默认值）; 1 大班课; 2 1V1 (后续扩展)
	RoomType *int64 `json:"RoomType,omitnil" name:"RoomType"`

	// 拖堂时间：单位分钟，0为不限制(默认值), -1为不能拖堂，大于0为拖堂的时间，最大值120分钟
	EndDelayTime *int64 `json:"EndDelayTime,omitnil" name:"EndDelayTime"`

	// 直播方式：0 常规模式（默认）1 回放直播模式（伪直播）
	LiveType *uint64 `json:"LiveType,omitnil" name:"LiveType"`

	// 伪直播回放链接
	RecordLiveUrl *string `json:"RecordLiveUrl,omitnil" name:"RecordLiveUrl"`

	// 是否自动开始上课：0 不自动上课（默认） 1 自动上课 live_type=1的时候有效
	EnableAutoStart *uint64 `json:"EnableAutoStart,omitnil" name:"EnableAutoStart"`
}

type RoomItem struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 房间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`

	// 房间状态。0 未开始 ；1进行中  ；2 已结束；3已过期
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *uint64 `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *uint64 `json:"EndTime,omitnil" name:"EndTime"`

	// 实际开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealStartTime *uint64 `json:"RealStartTime,omitnil" name:"RealStartTime"`

	// 实际结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealEndTime *uint64 `json:"RealEndTime,omitnil" name:"RealEndTime"`

	// 分辨率。1 标清
	// 2 高清
	// 3 全高清
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resolution *uint64 `json:"Resolution,omitnil" name:"Resolution"`

	// 最大允许连麦人数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxRTCMember *uint64 `json:"MaxRTCMember,omitnil" name:"MaxRTCMember"`

	// 房间录制地址。已废弃，使用新字段 RecordUrl
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReplayUrl *string `json:"ReplayUrl,omitnil" name:"ReplayUrl"`

	// 录制地址（协议为https)。仅在房间结束后存在。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordUrl *string `json:"RecordUrl,omitnil" name:"RecordUrl"`

	// 最高房间内人数（不包括老师），0表示不限制，默认为0
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxMicNumber *uint64 `json:"MaxMicNumber,omitnil" name:"MaxMicNumber"`

	// 打开学生麦克风/摄像头的授权开关 
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableDirectControl *uint64 `json:"EnableDirectControl,omitnil" name:"EnableDirectControl"`

	// 开启专注模式。 0 收看全部角色音视频(默认) 1 只看老师和助教
	// 注意：此字段可能返回 null，表示取不到有效值。
	InteractionMode *int64 `json:"InteractionMode,omitnil" name:"InteractionMode"`

	// 横竖屏。0：横屏开播（默认值）; 1：竖屏开播，当前仅支持移动端的纯视频类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoOrientation *int64 `json:"VideoOrientation,omitnil" name:"VideoOrientation"`

	// 开启课后评分。 0：不开启(默认)  1：开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsGradingRequiredPostClass *int64 `json:"IsGradingRequiredPostClass,omitnil" name:"IsGradingRequiredPostClass"`

	// 房间类型。0:小班课（默认值）；1:大班课；2:1V1（后续扩展）
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoomType *int64 `json:"RoomType,omitnil" name:"RoomType"`

	// 拖堂时间：单位分钟，0为不限制(默认值), -1为不能拖堂，大于0为拖堂的时间，最大值120分钟
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndDelayTime *int64 `json:"EndDelayTime,omitnil" name:"EndDelayTime"`

	// 直播方式：0 常规模式（默认）1 回放直播模式（伪直播）	
	// 注意：此字段可能返回 null，表示取不到有效值。
	LiveType *uint64 `json:"LiveType,omitnil" name:"LiveType"`

	// 伪直播回放链接	
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordLiveUrl *string `json:"RecordLiveUrl,omitnil" name:"RecordLiveUrl"`

	// 是否自动开始上课：0 不自动上课（默认） 1 自动上课 live_type=1的时候有效	
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableAutoStart *uint64 `json:"EnableAutoStart,omitnil" name:"EnableAutoStart"`

	// 录制文件背景图片，支持png、jpg、jpeg、bmp格式，暂不支持透明通道
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordBackground *string `json:"RecordBackground,omitnil" name:"RecordBackground"`
}

type SceneItem struct {

}

// Predefined struct for user
type SendRoomNormalMessageRequestParams struct {
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`

	// 管理员指定消息发送方账号（若需设置 FromAccount 信息，则该参数取值不能为空）
	FromAccount *string `json:"FromAccount,omitnil" name:"FromAccount"`

	// 自定义消息
	MsgBody []*MsgBody `json:"MsgBody,omitnil" name:"MsgBody"`

	// 消息自定义数据（云端保存，会发送到对端，程序卸载重装后还能拉取到）。
	CloudCustomData *string `json:"CloudCustomData,omitnil" name:"CloudCustomData"`

	// 昵称，当FromAccount没有在房间中，需要填写NickName，当FromAccount在房间中，填写NickName无意义
	NickName *string `json:"NickName,omitnil" name:"NickName"`
}

type SendRoomNormalMessageRequest struct {
	*tchttp.BaseRequest
	
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`

	// 管理员指定消息发送方账号（若需设置 FromAccount 信息，则该参数取值不能为空）
	FromAccount *string `json:"FromAccount,omitnil" name:"FromAccount"`

	// 自定义消息
	MsgBody []*MsgBody `json:"MsgBody,omitnil" name:"MsgBody"`

	// 消息自定义数据（云端保存，会发送到对端，程序卸载重装后还能拉取到）。
	CloudCustomData *string `json:"CloudCustomData,omitnil" name:"CloudCustomData"`

	// 昵称，当FromAccount没有在房间中，需要填写NickName，当FromAccount在房间中，填写NickName无意义
	NickName *string `json:"NickName,omitnil" name:"NickName"`
}

func (r *SendRoomNormalMessageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendRoomNormalMessageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	delete(f, "FromAccount")
	delete(f, "MsgBody")
	delete(f, "CloudCustomData")
	delete(f, "NickName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendRoomNormalMessageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendRoomNormalMessageResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SendRoomNormalMessageResponse struct {
	*tchttp.BaseResponse
	Response *SendRoomNormalMessageResponseParams `json:"Response"`
}

func (r *SendRoomNormalMessageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendRoomNormalMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendRoomNotificationMessageRequestParams struct {
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`

	// 消息。
	MsgContent *string `json:"MsgContent,omitnil" name:"MsgContent"`
}

type SendRoomNotificationMessageRequest struct {
	*tchttp.BaseRequest
	
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`

	// 消息。
	MsgContent *string `json:"MsgContent,omitnil" name:"MsgContent"`
}

func (r *SendRoomNotificationMessageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendRoomNotificationMessageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	delete(f, "MsgContent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendRoomNotificationMessageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendRoomNotificationMessageResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SendRoomNotificationMessageResponse struct {
	*tchttp.BaseResponse
	Response *SendRoomNotificationMessageResponseParams `json:"Response"`
}

func (r *SendRoomNotificationMessageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendRoomNotificationMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetAppCustomContentRequestParams struct {
	// 自定义内容。
	CustomContent []*AppCustomContent `json:"CustomContent,omitnil" name:"CustomContent"`

	// 应用ID。
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`
}

type SetAppCustomContentRequest struct {
	*tchttp.BaseRequest
	
	// 自定义内容。
	CustomContent []*AppCustomContent `json:"CustomContent,omitnil" name:"CustomContent"`

	// 应用ID。
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`
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
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 老师视频区域的水印参数地址，设置为空字符串表示删除
	TeacherUrl *string `json:"TeacherUrl,omitnil" name:"TeacherUrl"`

	// 白板视频区域的水印参数地址，设置为空字符串表示删除
	BoardUrl *string `json:"BoardUrl,omitnil" name:"BoardUrl"`

	// 视频默认图片（在没有视频流的时候显示），设置为空字符串表示删除
	VideoUrl *string `json:"VideoUrl,omitnil" name:"VideoUrl"`

	// 白板区域水印的宽度，取值:0-100，默认为0，表示区域X方向的百分比
	BoardW *float64 `json:"BoardW,omitnil" name:"BoardW"`

	// 白板区域水印的高度，取值:0-100，默认为0, 表示区域Y方向的百分比
	BoardH *float64 `json:"BoardH,omitnil" name:"BoardH"`

	// 白板区域水印X偏移, 取值:0-100, 表示区域X方向的百分比。比如50，则表示位于X轴中间
	BoardX *float64 `json:"BoardX,omitnil" name:"BoardX"`

	// 白板区域水印Y偏移, 取值:0-100, 表示区域X方向的百分比。比如50，则表示位于X轴中间
	BoardY *float64 `json:"BoardY,omitnil" name:"BoardY"`

	// 老师视频区域水印的宽度，取值:0-100，默认为0，表示区域X方向的百分比
	TeacherW *float64 `json:"TeacherW,omitnil" name:"TeacherW"`

	// 老师视频区域水印的高度，取值:0-100，默认为0, 表示区域Y方向的百分比
	TeacherH *float64 `json:"TeacherH,omitnil" name:"TeacherH"`

	// 老师视频区域水印X偏移, 取值:0-100, 表示区域X方向的百分比。比如50，则表示位于X轴中间
	TeacherX *float64 `json:"TeacherX,omitnil" name:"TeacherX"`

	// 老师视频区域水印Y偏移, 取值:0-100, 表示区域X方向的百分比。比如50，则表示位于X轴中间
	TeacherY *float64 `json:"TeacherY,omitnil" name:"TeacherY"`

	// 文字水印内容，设置为空字符串表示删除
	Text *string `json:"Text,omitnil" name:"Text"`

	// 文字水印颜色
	TextColor *string `json:"TextColor,omitnil" name:"TextColor"`
}

type SetWatermarkRequest struct {
	*tchttp.BaseRequest
	
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 老师视频区域的水印参数地址，设置为空字符串表示删除
	TeacherUrl *string `json:"TeacherUrl,omitnil" name:"TeacherUrl"`

	// 白板视频区域的水印参数地址，设置为空字符串表示删除
	BoardUrl *string `json:"BoardUrl,omitnil" name:"BoardUrl"`

	// 视频默认图片（在没有视频流的时候显示），设置为空字符串表示删除
	VideoUrl *string `json:"VideoUrl,omitnil" name:"VideoUrl"`

	// 白板区域水印的宽度，取值:0-100，默认为0，表示区域X方向的百分比
	BoardW *float64 `json:"BoardW,omitnil" name:"BoardW"`

	// 白板区域水印的高度，取值:0-100，默认为0, 表示区域Y方向的百分比
	BoardH *float64 `json:"BoardH,omitnil" name:"BoardH"`

	// 白板区域水印X偏移, 取值:0-100, 表示区域X方向的百分比。比如50，则表示位于X轴中间
	BoardX *float64 `json:"BoardX,omitnil" name:"BoardX"`

	// 白板区域水印Y偏移, 取值:0-100, 表示区域X方向的百分比。比如50，则表示位于X轴中间
	BoardY *float64 `json:"BoardY,omitnil" name:"BoardY"`

	// 老师视频区域水印的宽度，取值:0-100，默认为0，表示区域X方向的百分比
	TeacherW *float64 `json:"TeacherW,omitnil" name:"TeacherW"`

	// 老师视频区域水印的高度，取值:0-100，默认为0, 表示区域Y方向的百分比
	TeacherH *float64 `json:"TeacherH,omitnil" name:"TeacherH"`

	// 老师视频区域水印X偏移, 取值:0-100, 表示区域X方向的百分比。比如50，则表示位于X轴中间
	TeacherX *float64 `json:"TeacherX,omitnil" name:"TeacherX"`

	// 老师视频区域水印Y偏移, 取值:0-100, 表示区域X方向的百分比。比如50，则表示位于X轴中间
	TeacherY *float64 `json:"TeacherY,omitnil" name:"TeacherY"`

	// 文字水印内容，设置为空字符串表示删除
	Text *string `json:"Text,omitnil" name:"Text"`

	// 文字水印颜色
	TextColor *string `json:"TextColor,omitnil" name:"TextColor"`
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
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

// Predefined struct for user
type StartRoomRequestParams struct {
	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`
}

type StartRoomRequest struct {
	*tchttp.BaseRequest
	
	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`
}

func (r *StartRoomRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartRoomRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartRoomRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartRoomResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type StartRoomResponse struct {
	*tchttp.BaseResponse
	Response *StartRoomResponseParams `json:"Response"`
}

func (r *StartRoomResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartRoomResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TextMarkConfig struct {
	// 文字水印内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitnil" name:"Text"`

	// 文字水印颜色
	// 注意：此字段可能返回 null，表示取不到有效值。
	Color *string `json:"Color,omitnil" name:"Color"`
}

type TextMsgContent struct {
	// 文本消息。
	Text *string `json:"Text,omitnil" name:"Text"`
}

// Predefined struct for user
type UnbindDocumentFromRoomRequestParams struct {
	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`

	// 文档ID。
	DocumentId *string `json:"DocumentId,omitnil" name:"DocumentId"`
}

type UnbindDocumentFromRoomRequest struct {
	*tchttp.BaseRequest
	
	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`

	// 文档ID。
	DocumentId *string `json:"DocumentId,omitnil" name:"DocumentId"`
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
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 用户Id。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 用户昵称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 用户头像Url。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Avatar *string `json:"Avatar,omitnil" name:"Avatar"`

	// 用户在客户系统的Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginId *string `json:"OriginId,omitnil" name:"OriginId"`
}

type WatermarkConfig struct {
	// 水印图片的url
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil" name:"Url"`

	// 水印宽。为比例值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Width *float64 `json:"Width,omitnil" name:"Width"`

	// 水印高。为比例值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Height *float64 `json:"Height,omitnil" name:"Height"`

	// 水印X偏移, 取值:0-100, 表示区域X方向的百分比。比如50，则表示位于X轴中间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocationX *float64 `json:"LocationX,omitnil" name:"LocationX"`

	// 水印Y偏移, 取值:0-100, 表示区域Y方向的百分比。比如50，则表示位于Y轴中间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocationY *float64 `json:"LocationY,omitnil" name:"LocationY"`
}