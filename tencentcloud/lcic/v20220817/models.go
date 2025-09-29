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

package v20220817

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type AddGroupMemberRequestParams struct {
	// 群组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 成员列表，最大值200
	MemberIds []*string `json:"MemberIds,omitnil,omitempty" name:"MemberIds"`
}

type AddGroupMemberRequest struct {
	*tchttp.BaseRequest
	
	// 群组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 成员列表，最大值200
	MemberIds []*string `json:"MemberIds,omitnil,omitempty" name:"MemberIds"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 答案（按照位表示是否选择，如0x1表示选择A，0x11表示选择AB）
	Answer *uint64 `json:"Answer,omitnil,omitempty" name:"Answer"`

	// 答题用时
	CostTime *uint64 `json:"CostTime,omitnil,omitempty" name:"CostTime"`

	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 答案是否正确（1正确0错误）
	IsCorrect *uint64 `json:"IsCorrect,omitnil,omitempty" name:"IsCorrect"`
}

type AnswerStat struct {
	// 选项（按照位表示是否选择，如0x1表示选择A，0x11表示选择AB）
	Answer *int64 `json:"Answer,omitnil,omitempty" name:"Answer"`

	// 答题人数
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type AppConfig struct {
	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 应用名称
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 应用状态 1正常 2停用
	State *int64 `json:"State,omitnil,omitempty" name:"State"`

	// 1试用 2轻量版 3标准版 4旗舰版
	AppVersion *int64 `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 回调
	Callback *string `json:"Callback,omitnil,omitempty" name:"Callback"`

	// 回调Key
	CallbackKey *string `json:"CallbackKey,omitnil,omitempty" name:"CallbackKey"`
}

type AppCustomContent struct {
	// 场景参数，一个应用下可以设置多个不同场景。
	Scene *string `json:"Scene,omitnil,omitempty" name:"Scene"`

	// logo地址，用于上课时展示的课堂或平台图标，支持开发商自定义业务品牌展示。
	LogoUrl *string `json:"LogoUrl,omitnil,omitempty" name:"LogoUrl"`

	// HomeUrl：主页地址，用于上课结束后课堂跳转，支持跳转到自己的业务系统。如果配置为空则下课后关闭课堂页面。
	HomeUrl *string `json:"HomeUrl,omitnil,omitempty" name:"HomeUrl"`

	// JsUrl ：自定义js。针对应用用于开发上自定义课堂界面、模块功能、监控操作，支持数据请求与响应处理。
	JsUrl *string `json:"JsUrl,omitnil,omitempty" name:"JsUrl"`

	// Css : 自定义的css。针对应用用于支持课堂界面的、模块的UI渲染修改、皮肤配色修改、功能模块的隐藏和展示。
	CssUrl *string `json:"CssUrl,omitnil,omitempty" name:"CssUrl"`
}

type BackgroundPictureConfig struct {
	// 背景图片的url
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

// Predefined struct for user
type BatchAddGroupMemberRequestParams struct {
	// 待添加群组ID列表，最大值100
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 待添加成员列表，最大值200
	MemberIds []*string `json:"MemberIds,omitnil,omitempty" name:"MemberIds"`
}

type BatchAddGroupMemberRequest struct {
	*tchttp.BaseRequest
	
	// 待添加群组ID列表，最大值100
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 待添加成员列表，最大值200
	MemberIds []*string `json:"MemberIds,omitnil,omitempty" name:"MemberIds"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 批量创建群组基础信息，最大长度限制256
	GroupBaseInfos []*GroupBaseInfo `json:"GroupBaseInfos,omitnil,omitempty" name:"GroupBaseInfos"`

	// 群组绑定的成员列表，一次性最多200个
	MemberIds []*string `json:"MemberIds,omitnil,omitempty" name:"MemberIds"`
}

type BatchCreateGroupWithMembersRequest struct {
	*tchttp.BaseRequest
	
	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 批量创建群组基础信息，最大长度限制256
	GroupBaseInfos []*GroupBaseInfo `json:"GroupBaseInfos,omitnil,omitempty" name:"GroupBaseInfos"`

	// 群组绑定的成员列表，一次性最多200个
	MemberIds []*string `json:"MemberIds,omitnil,omitempty" name:"MemberIds"`
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
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 批量创建课堂的配置信息
	RoomInfos []*RoomInfo `json:"RoomInfos,omitnil,omitempty" name:"RoomInfos"`
}

type BatchCreateRoomRequest struct {
	*tchttp.BaseRequest
	
	// 低代码平台的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 批量创建课堂的配置信息
	RoomInfos []*RoomInfo `json:"RoomInfos,omitnil,omitempty" name:"RoomInfos"`
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
	RoomIds []*uint64 `json:"RoomIds,omitnil,omitempty" name:"RoomIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 待添加成员列表，最大值256
	MemberIds []*string `json:"MemberIds,omitnil,omitempty" name:"MemberIds"`
}

type BatchDeleteGroupMemberRequest struct {
	*tchttp.BaseRequest
	
	// 待添加群组ID列表，最大值100
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 待添加成员列表，最大值256
	MemberIds []*string `json:"MemberIds,omitnil,omitempty" name:"MemberIds"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	RoomIds []*int64 `json:"RoomIds,omitnil,omitempty" name:"RoomIds"`

	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
}

type BatchDeleteRecordRequest struct {
	*tchttp.BaseRequest
	
	// 房间ID列表
	RoomIds []*int64 `json:"RoomIds,omitnil,omitempty" name:"RoomIds"`

	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
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
	RoomIds []*int64 `json:"RoomIds,omitnil,omitempty" name:"RoomIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 分页查询当前页数，从1开始递增
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 每页数据量，最大200
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 课件权限。[0]：获取owner的私有课件；[1]：获取owner的公开课件; [0,1]：则获取owner的私有课件和公开课件；[2]：获取owner的私有课件和所有人(包括owner)的公开课件
	Permission []*uint64 `json:"Permission,omitnil,omitempty" name:"Permission"`

	// 课件所有者的user_id，不填默认获取SdkAppId下所有课件
	Owner *string `json:"Owner,omitnil,omitempty" name:"Owner"`

	// 课件名称搜索词
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 课件id列表，从列表中查询，忽略错误的id
	DocumentId []*string `json:"DocumentId,omitnil,omitempty" name:"DocumentId"`
}

type BatchDescribeDocumentRequest struct {
	*tchttp.BaseRequest
	
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 分页查询当前页数，从1开始递增
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 每页数据量，最大200
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 课件权限。[0]：获取owner的私有课件；[1]：获取owner的公开课件; [0,1]：则获取owner的私有课件和公开课件；[2]：获取owner的私有课件和所有人(包括owner)的公开课件
	Permission []*uint64 `json:"Permission,omitnil,omitempty" name:"Permission"`

	// 课件所有者的user_id，不填默认获取SdkAppId下所有课件
	Owner *string `json:"Owner,omitnil,omitempty" name:"Owner"`

	// 课件名称搜索词
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 课件id列表，从列表中查询，忽略错误的id
	DocumentId []*string `json:"DocumentId,omitnil,omitempty" name:"DocumentId"`
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
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 文档信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Documents []*DocumentInfo `json:"Documents,omitnil,omitempty" name:"Documents"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Users []*BatchUserRequest `json:"Users,omitnil,omitempty" name:"Users"`
}

type BatchRegisterRequest struct {
	*tchttp.BaseRequest
	
	// 批量注册用户信息列表
	Users []*BatchUserRequest `json:"Users,omitnil,omitempty" name:"Users"`
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
	Users []*BatchUserInfo `json:"Users,omitnil,omitempty" name:"Users"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 用户ID。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户在客户系统的Id。 若用户注册时该字段为空，则默认为 UserId 值一致。
	OriginId *string `json:"OriginId,omitnil,omitempty" name:"OriginId"`
}

type BatchUserRequest struct {
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 用户名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 用户在客户系统的Id，需要在同一应用下唯一。入参为空时默认赋值为UserId
	// 。
	OriginId *string `json:"OriginId,omitnil,omitempty" name:"OriginId"`

	// 用户头像。
	Avatar *string `json:"Avatar,omitnil,omitempty" name:"Avatar"`
}

// Predefined struct for user
type BindDocumentToRoomRequestParams struct {
	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 文档ID。
	DocumentId *string `json:"DocumentId,omitnil,omitempty" name:"DocumentId"`

	// 绑定类型。后台可透传到客户端，默认为0。除以下例值外支持自定义该字段，并在前端实现相应业务逻辑，示例参考：
	// 示例值：0，仅绑定课件到房间
	// 示例值：1，绑定课件到房间后，默认展示课件
	BindType *uint64 `json:"BindType,omitnil,omitempty" name:"BindType"`
}

type BindDocumentToRoomRequest struct {
	*tchttp.BaseRequest
	
	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 文档ID。
	DocumentId *string `json:"DocumentId,omitnil,omitempty" name:"DocumentId"`

	// 绑定类型。后台可透传到客户端，默认为0。除以下例值外支持自定义该字段，并在前端实现相应业务逻辑，示例参考：
	// 示例值：0，仅绑定课件到房间
	// 示例值：1，绑定课件到房间后，默认展示课件
	BindType *uint64 `json:"BindType,omitnil,omitempty" name:"BindType"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 评分时间
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 课堂评分
	Score *uint64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 课堂评价
	ScoreMsg *string `json:"ScoreMsg,omitnil,omitempty" name:"ScoreMsg"`
}

// Predefined struct for user
type CreateDocumentRequestParams struct {
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 文档地址。
	DocumentUrl *string `json:"DocumentUrl,omitnil,omitempty" name:"DocumentUrl"`

	// 文档名称。
	DocumentName *string `json:"DocumentName,omitnil,omitempty" name:"DocumentName"`

	// 文档所有者的Id
	Owner *string `json:"Owner,omitnil,omitempty" name:"Owner"`

	// 转码类型，可以有如下取值：
	// 0 无需转码（默认），bmp，jpg，jpeg，png，gif
	// 1 需要转码的文档，ppt，pptx，pdf，doc，docx，xls，xlsx
	// 2 需要转码的视频，mp4，3pg，mpeg，avi，flv，wmv，rm，h264等
	// 2 需要转码的音频，mp3，wav，wma，aac，flac，opus
	// 请注意，待录制的页面中任何视频的分辨率不能超过页面录制最大分辨率（1920*1080），否则将导致录制失败。
	//  - ppt课件内嵌视频或纯视频课件，在上传课件时，云api会进行转码，以确保视频分辨率不超过页面录制最大分辨率。
	//  - h5课件中内嵌音视频内容时，由于平台无法获取视频内容，因此在制作环节需确保视频分辨率不超过页面录制最大分辨率。
	TranscodeType *uint64 `json:"TranscodeType,omitnil,omitempty" name:"TranscodeType"`

	// 权限，可以有如下取值：
	// 0 私有文档（默认）
	// 1 公共文档
	Permission *uint64 `json:"Permission,omitnil,omitempty" name:"Permission"`

	// 文档后缀名。
	DocumentType *string `json:"DocumentType,omitnil,omitempty" name:"DocumentType"`

	// 文档大小，单位 字节
	DocumentSize *uint64 `json:"DocumentSize,omitnil,omitempty" name:"DocumentSize"`

	// 是否对不支持元素开启自动处理的功能。默认关闭。
	// 自动处理的元素如下：
	// 1. 墨迹：移除不支持的墨迹（例如WPS墨迹）
	// 2. 自动翻页：移除PPT上所有自动翻页设置，并设置为单击鼠标翻页
	// 3. 已损坏音视频：移除PPT上对损坏音视频的引用
	AutoHandleUnsupportedElement *bool `json:"AutoHandleUnsupportedElement,omitnil,omitempty" name:"AutoHandleUnsupportedElement"`

	// 转码后文档的最小分辨率，不传、传空字符串或分辨率格式错误则使用文档原分辨率。该参数仅对TranscodeType=1的课件生效。示例：1280x720，注意分辨率宽高中间为英文字母"xyz"的"x"
	// 示例值：1280x720
	MinScaleResolution *string `json:"MinScaleResolution,omitnil,omitempty" name:"MinScaleResolution"`
}

type CreateDocumentRequest struct {
	*tchttp.BaseRequest
	
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 文档地址。
	DocumentUrl *string `json:"DocumentUrl,omitnil,omitempty" name:"DocumentUrl"`

	// 文档名称。
	DocumentName *string `json:"DocumentName,omitnil,omitempty" name:"DocumentName"`

	// 文档所有者的Id
	Owner *string `json:"Owner,omitnil,omitempty" name:"Owner"`

	// 转码类型，可以有如下取值：
	// 0 无需转码（默认），bmp，jpg，jpeg，png，gif
	// 1 需要转码的文档，ppt，pptx，pdf，doc，docx，xls，xlsx
	// 2 需要转码的视频，mp4，3pg，mpeg，avi，flv，wmv，rm，h264等
	// 2 需要转码的音频，mp3，wav，wma，aac，flac，opus
	// 请注意，待录制的页面中任何视频的分辨率不能超过页面录制最大分辨率（1920*1080），否则将导致录制失败。
	//  - ppt课件内嵌视频或纯视频课件，在上传课件时，云api会进行转码，以确保视频分辨率不超过页面录制最大分辨率。
	//  - h5课件中内嵌音视频内容时，由于平台无法获取视频内容，因此在制作环节需确保视频分辨率不超过页面录制最大分辨率。
	TranscodeType *uint64 `json:"TranscodeType,omitnil,omitempty" name:"TranscodeType"`

	// 权限，可以有如下取值：
	// 0 私有文档（默认）
	// 1 公共文档
	Permission *uint64 `json:"Permission,omitnil,omitempty" name:"Permission"`

	// 文档后缀名。
	DocumentType *string `json:"DocumentType,omitnil,omitempty" name:"DocumentType"`

	// 文档大小，单位 字节
	DocumentSize *uint64 `json:"DocumentSize,omitnil,omitempty" name:"DocumentSize"`

	// 是否对不支持元素开启自动处理的功能。默认关闭。
	// 自动处理的元素如下：
	// 1. 墨迹：移除不支持的墨迹（例如WPS墨迹）
	// 2. 自动翻页：移除PPT上所有自动翻页设置，并设置为单击鼠标翻页
	// 3. 已损坏音视频：移除PPT上对损坏音视频的引用
	AutoHandleUnsupportedElement *bool `json:"AutoHandleUnsupportedElement,omitnil,omitempty" name:"AutoHandleUnsupportedElement"`

	// 转码后文档的最小分辨率，不传、传空字符串或分辨率格式错误则使用文档原分辨率。该参数仅对TranscodeType=1的课件生效。示例：1280x720，注意分辨率宽高中间为英文字母"xyz"的"x"
	// 示例值：1280x720
	MinScaleResolution *string `json:"MinScaleResolution,omitnil,omitempty" name:"MinScaleResolution"`
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
	delete(f, "MinScaleResolution")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDocumentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDocumentResponseParams struct {
	// 文档ID。
	DocumentId *string `json:"DocumentId,omitnil,omitempty" name:"DocumentId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 默认绑定主讲老师ID
	TeacherId *string `json:"TeacherId,omitnil,omitempty" name:"TeacherId"`

	// 群组成员列表,一次性最多200个
	MemberIds []*string `json:"MemberIds,omitnil,omitempty" name:"MemberIds"`
}

type CreateGroupWithMembersRequest struct {
	*tchttp.BaseRequest
	
	// 待创建群组名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 默认绑定主讲老师ID
	TeacherId *string `json:"TeacherId,omitnil,omitempty" name:"TeacherId"`

	// 群组成员列表,一次性最多200个
	MemberIds []*string `json:"MemberIds,omitnil,omitempty" name:"MemberIds"`
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
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 子群组ID列表，子群组ID不能重复，最多40个
	SubGroupIds []*string `json:"SubGroupIds,omitnil,omitempty" name:"SubGroupIds"`

	// 群组默认主讲老师ID
	TeacherId *string `json:"TeacherId,omitnil,omitempty" name:"TeacherId"`
}

type CreateGroupWithSubGroupRequest struct {
	*tchttp.BaseRequest
	
	// 待创建的新群组名
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 子群组ID列表，子群组ID不能重复，最多40个
	SubGroupIds []*string `json:"SubGroupIds,omitnil,omitempty" name:"SubGroupIds"`

	// 群组默认主讲老师ID
	TeacherId *string `json:"TeacherId,omitnil,omitempty" name:"TeacherId"`
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
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 课堂名称。
	// 字符数不超过256
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 预定的课堂开始时间，unix时间戳（秒）。
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 预定的课堂结束时间，unix时间戳（秒）。
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 头像区域，摄像头视频画面的分辨率。可以有如下取值：
	// 1 标清
	// 2 高清
	// 3 全高清
	// 注意：连麦人数（MaxMicNumber）>6时，仅可使用标清
	Resolution *uint64 `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// 设置课堂同时最大可与老师进行连麦互动的人数，该参数支持正式上课/开播前调用修改房间修改。小班课取值范围[0,16]，大班课取值范围[0,1]，当取值为0时表示当前课堂/直播，不支持连麦互动。该取值影响计费，请根据业务实际情况设置。计费规则见“购买指南”下“计费概述”。
	MaxMicNumber *uint64 `json:"MaxMicNumber,omitnil,omitempty" name:"MaxMicNumber"`

	// 课堂子类型，可以有以下取值：videodoc 文档+视频video 纯视频
	SubType *string `json:"SubType,omitnil,omitempty" name:"SubType"`

	// 老师ID。通过[注册用户]接口获取的UserId。指定后该用户在房间内拥有老师权限。
	TeacherId *string `json:"TeacherId,omitnil,omitempty" name:"TeacherId"`

	// 进入课堂时是否自动连麦。可以有以下取值：
	// 0 不自动连麦（需要手动申请上麦，默认值）
	// 1 自动连麦
	AutoMic *uint64 `json:"AutoMic,omitnil,omitempty" name:"AutoMic"`

	// 释放音视频权限后是否自动取消连麦。可以有以下取值：
	// 0 自动取消连麦（默认值）
	// 1 保持连麦状态
	TurnOffMic *uint64 `json:"TurnOffMic,omitnil,omitempty" name:"TurnOffMic"`

	// 声音音质。可以有以下取值：
	// 0：流畅模式（默认值），占用更小的带宽、拥有更好的降噪效果，适用于1对1、小班教学、多人音视频会议等场景。
	// 1：高音质模式，适合需要高保真传输音乐的场景，但降噪效果会被削弱，适用于音乐教学场景。
	AudioQuality *uint64 `json:"AudioQuality,omitnil,omitempty" name:"AudioQuality"`

	// 录制方式，可以有以下取值：0 开启自动录制（默认值）1  禁止录制2 开启手动录制 注： - 如果该配置取值为0，录制将从上课后开始，课堂结束后停止。 - 如果该配置取值为2，需通过startRecord、stopRecord接口控制录制的开始和结束。 
	DisableRecord *uint64 `json:"DisableRecord,omitnil,omitempty" name:"DisableRecord"`

	// 助教Id列表。通过[注册用户]接口获取的UserId。指定后该用户在房间内拥有助教权限。
	Assistants []*string `json:"Assistants,omitnil,omitempty" name:"Assistants"`

	// rtc人数。
	//
	// Deprecated: RTCAudienceNumber is deprecated.
	RTCAudienceNumber *uint64 `json:"RTCAudienceNumber,omitnil,omitempty" name:"RTCAudienceNumber"`

	// 观看类型。互动观看 （默认）
	AudienceType *uint64 `json:"AudienceType,omitnil,omitempty" name:"AudienceType"`

	// 录制模板。未配置时默认取值0。录制模板枚举值参考：https://cloud.tencent.com/document/product/1639/89744
	RecordLayout *uint64 `json:"RecordLayout,omitnil,omitempty" name:"RecordLayout"`

	// 课堂绑定的群组ID,非空时限制组成员进入
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 是否允许老师/助教直接控制学生的摄像头/麦克风。可以有以下取值：
	// 0 不允许直接控制（需同意，默认值）
	// 1 允许直接控制（无需同意）
	EnableDirectControl *uint64 `json:"EnableDirectControl,omitnil,omitempty" name:"EnableDirectControl"`

	// 开启专注模式。
	// 0 收看全部角色音视频(默认)
	// 1 只看老师和助教
	InteractionMode *int64 `json:"InteractionMode,omitnil,omitempty" name:"InteractionMode"`

	// 横竖屏。0：横屏开播（默认值）; 1：竖屏开播，当前仅支持移动端的纯视频类型
	VideoOrientation *uint64 `json:"VideoOrientation,omitnil,omitempty" name:"VideoOrientation"`

	// 开启课后评分。 0：不开启(默认)  1：开启
	IsGradingRequiredPostClass *int64 `json:"IsGradingRequiredPostClass,omitnil,omitempty" name:"IsGradingRequiredPostClass"`

	// 课堂类型: 0 小班课（默认值）; 1 大班课; 2 1V1 (预留参数，暂未开放); 3 圆桌会议 注：大班课的布局(layout)只有三分屏
	RoomType *int64 `json:"RoomType,omitnil,omitempty" name:"RoomType"`

	// 嘉宾Id列表。当圆桌会议模式（RoomType==3）时生效
	Guests []*string `json:"Guests,omitnil,omitempty" name:"Guests"`

	// 拖堂时间：单位分钟，0为不限制(默认值), -1为不能拖堂，大于0为拖堂的时间，最大值120分钟
	EndDelayTime *int64 `json:"EndDelayTime,omitnil,omitempty" name:"EndDelayTime"`

	// 直播类型：0 常规（默认）1 伪直播 2 RTMP推流直播
	LiveType *uint64 `json:"LiveType,omitnil,omitempty" name:"LiveType"`

	// 伪直播链接。 支持的协议以及格式： 协议：HTTP、HTTPS、RTMP、HLS 。格式：FLV、MP3、MP4、MPEG-TS、MOV、MKV、M4A。视频编码：H.264、VP8。音频编码：AAC、OPUS。
	RecordLiveUrl *string `json:"RecordLiveUrl,omitnil,omitempty" name:"RecordLiveUrl"`

	// 是否自动开始上课：0 不自动上课（默认） 1 自动上课 live_type=1或2的时候有效
	EnableAutoStart *uint64 `json:"EnableAutoStart,omitnil,omitempty" name:"EnableAutoStart"`

	// 录制文件背景图片，支持png、jpg、jpeg、bmp格式，暂不支持透明通道
	RecordBackground *string `json:"RecordBackground,omitnil,omitempty" name:"RecordBackground"`

	// 录制自定义场景。注意：仅recordlayout=9的时候此参数有效。需注意各类参数配置正确能够生效。不然会造成录制失败，失败后无法补救。数据内容为用户自定义场景参数，数据格式为json键值对方式，其中键值对的value为string类型。
	// 
	// 自定义场景参数的含义。如下：
	//      scene：自定义js/css对应的场景值。如scene=recordScene，会加载 recordScene 场景对应的 js/css，这样就可以自定义录制页面的元素。 
	//     lng：录制页面对应的语种。如lng=en，则录制界面为en。（枚举值：en,zh，zh-TW，jp，ar，kr，vi）
	//      customToken：录制页面中涉及客户自己的服务需要鉴权时进行配置。一般情况下，无需配置。
	RecordScene *string `json:"RecordScene,omitnil,omitempty" name:"RecordScene"`

	// 录制自定义语言，仅recordlayout=9的时候此参数有效
	//
	// Deprecated: RecordLang is deprecated.
	RecordLang *string `json:"RecordLang,omitnil,omitempty" name:"RecordLang"`

	// 录制类型 0 仅录制混流（默认） ;1 录制混流+单流，该模式下除混流录制基础上，分别录制老师、台上学生的音视频流，每路录制都会产生相应的录制费用 。示例：0
	RecordStream *uint64 `json:"RecordStream,omitnil,omitempty" name:"RecordStream"`

	// 板书截图生成类型。0 不生成板书（默认）；1 全量模式；2 单页去重模式
	WhiteBoardSnapshotMode *uint64 `json:"WhiteBoardSnapshotMode,omitnil,omitempty" name:"WhiteBoardSnapshotMode"`

	// 字幕转写功能开关。可以有以下取值：
	// 0 不开启字幕转写功能（默认值）
	// 1 自动转写模式：上课自动开启，下课自动停止
	SubtitlesTranscription *uint64 `json:"SubtitlesTranscription,omitnil,omitempty" name:"SubtitlesTranscription"`

	// 录制文件合并开关。0 关闭 1 开启 注：只有在一节课多次启用手动录制时，此功能才有效
	RecordMerge *uint64 `json:"RecordMerge,omitnil,omitempty" name:"RecordMerge"`
}

type CreateRoomRequest struct {
	*tchttp.BaseRequest
	
	// 课堂名称。
	// 字符数不超过256
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 预定的课堂开始时间，unix时间戳（秒）。
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 预定的课堂结束时间，unix时间戳（秒）。
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 头像区域，摄像头视频画面的分辨率。可以有如下取值：
	// 1 标清
	// 2 高清
	// 3 全高清
	// 注意：连麦人数（MaxMicNumber）>6时，仅可使用标清
	Resolution *uint64 `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// 设置课堂同时最大可与老师进行连麦互动的人数，该参数支持正式上课/开播前调用修改房间修改。小班课取值范围[0,16]，大班课取值范围[0,1]，当取值为0时表示当前课堂/直播，不支持连麦互动。该取值影响计费，请根据业务实际情况设置。计费规则见“购买指南”下“计费概述”。
	MaxMicNumber *uint64 `json:"MaxMicNumber,omitnil,omitempty" name:"MaxMicNumber"`

	// 课堂子类型，可以有以下取值：videodoc 文档+视频video 纯视频
	SubType *string `json:"SubType,omitnil,omitempty" name:"SubType"`

	// 老师ID。通过[注册用户]接口获取的UserId。指定后该用户在房间内拥有老师权限。
	TeacherId *string `json:"TeacherId,omitnil,omitempty" name:"TeacherId"`

	// 进入课堂时是否自动连麦。可以有以下取值：
	// 0 不自动连麦（需要手动申请上麦，默认值）
	// 1 自动连麦
	AutoMic *uint64 `json:"AutoMic,omitnil,omitempty" name:"AutoMic"`

	// 释放音视频权限后是否自动取消连麦。可以有以下取值：
	// 0 自动取消连麦（默认值）
	// 1 保持连麦状态
	TurnOffMic *uint64 `json:"TurnOffMic,omitnil,omitempty" name:"TurnOffMic"`

	// 声音音质。可以有以下取值：
	// 0：流畅模式（默认值），占用更小的带宽、拥有更好的降噪效果，适用于1对1、小班教学、多人音视频会议等场景。
	// 1：高音质模式，适合需要高保真传输音乐的场景，但降噪效果会被削弱，适用于音乐教学场景。
	AudioQuality *uint64 `json:"AudioQuality,omitnil,omitempty" name:"AudioQuality"`

	// 录制方式，可以有以下取值：0 开启自动录制（默认值）1  禁止录制2 开启手动录制 注： - 如果该配置取值为0，录制将从上课后开始，课堂结束后停止。 - 如果该配置取值为2，需通过startRecord、stopRecord接口控制录制的开始和结束。 
	DisableRecord *uint64 `json:"DisableRecord,omitnil,omitempty" name:"DisableRecord"`

	// 助教Id列表。通过[注册用户]接口获取的UserId。指定后该用户在房间内拥有助教权限。
	Assistants []*string `json:"Assistants,omitnil,omitempty" name:"Assistants"`

	// rtc人数。
	RTCAudienceNumber *uint64 `json:"RTCAudienceNumber,omitnil,omitempty" name:"RTCAudienceNumber"`

	// 观看类型。互动观看 （默认）
	AudienceType *uint64 `json:"AudienceType,omitnil,omitempty" name:"AudienceType"`

	// 录制模板。未配置时默认取值0。录制模板枚举值参考：https://cloud.tencent.com/document/product/1639/89744
	RecordLayout *uint64 `json:"RecordLayout,omitnil,omitempty" name:"RecordLayout"`

	// 课堂绑定的群组ID,非空时限制组成员进入
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 是否允许老师/助教直接控制学生的摄像头/麦克风。可以有以下取值：
	// 0 不允许直接控制（需同意，默认值）
	// 1 允许直接控制（无需同意）
	EnableDirectControl *uint64 `json:"EnableDirectControl,omitnil,omitempty" name:"EnableDirectControl"`

	// 开启专注模式。
	// 0 收看全部角色音视频(默认)
	// 1 只看老师和助教
	InteractionMode *int64 `json:"InteractionMode,omitnil,omitempty" name:"InteractionMode"`

	// 横竖屏。0：横屏开播（默认值）; 1：竖屏开播，当前仅支持移动端的纯视频类型
	VideoOrientation *uint64 `json:"VideoOrientation,omitnil,omitempty" name:"VideoOrientation"`

	// 开启课后评分。 0：不开启(默认)  1：开启
	IsGradingRequiredPostClass *int64 `json:"IsGradingRequiredPostClass,omitnil,omitempty" name:"IsGradingRequiredPostClass"`

	// 课堂类型: 0 小班课（默认值）; 1 大班课; 2 1V1 (预留参数，暂未开放); 3 圆桌会议 注：大班课的布局(layout)只有三分屏
	RoomType *int64 `json:"RoomType,omitnil,omitempty" name:"RoomType"`

	// 嘉宾Id列表。当圆桌会议模式（RoomType==3）时生效
	Guests []*string `json:"Guests,omitnil,omitempty" name:"Guests"`

	// 拖堂时间：单位分钟，0为不限制(默认值), -1为不能拖堂，大于0为拖堂的时间，最大值120分钟
	EndDelayTime *int64 `json:"EndDelayTime,omitnil,omitempty" name:"EndDelayTime"`

	// 直播类型：0 常规（默认）1 伪直播 2 RTMP推流直播
	LiveType *uint64 `json:"LiveType,omitnil,omitempty" name:"LiveType"`

	// 伪直播链接。 支持的协议以及格式： 协议：HTTP、HTTPS、RTMP、HLS 。格式：FLV、MP3、MP4、MPEG-TS、MOV、MKV、M4A。视频编码：H.264、VP8。音频编码：AAC、OPUS。
	RecordLiveUrl *string `json:"RecordLiveUrl,omitnil,omitempty" name:"RecordLiveUrl"`

	// 是否自动开始上课：0 不自动上课（默认） 1 自动上课 live_type=1或2的时候有效
	EnableAutoStart *uint64 `json:"EnableAutoStart,omitnil,omitempty" name:"EnableAutoStart"`

	// 录制文件背景图片，支持png、jpg、jpeg、bmp格式，暂不支持透明通道
	RecordBackground *string `json:"RecordBackground,omitnil,omitempty" name:"RecordBackground"`

	// 录制自定义场景。注意：仅recordlayout=9的时候此参数有效。需注意各类参数配置正确能够生效。不然会造成录制失败，失败后无法补救。数据内容为用户自定义场景参数，数据格式为json键值对方式，其中键值对的value为string类型。
	// 
	// 自定义场景参数的含义。如下：
	//      scene：自定义js/css对应的场景值。如scene=recordScene，会加载 recordScene 场景对应的 js/css，这样就可以自定义录制页面的元素。 
	//     lng：录制页面对应的语种。如lng=en，则录制界面为en。（枚举值：en,zh，zh-TW，jp，ar，kr，vi）
	//      customToken：录制页面中涉及客户自己的服务需要鉴权时进行配置。一般情况下，无需配置。
	RecordScene *string `json:"RecordScene,omitnil,omitempty" name:"RecordScene"`

	// 录制自定义语言，仅recordlayout=9的时候此参数有效
	RecordLang *string `json:"RecordLang,omitnil,omitempty" name:"RecordLang"`

	// 录制类型 0 仅录制混流（默认） ;1 录制混流+单流，该模式下除混流录制基础上，分别录制老师、台上学生的音视频流，每路录制都会产生相应的录制费用 。示例：0
	RecordStream *uint64 `json:"RecordStream,omitnil,omitempty" name:"RecordStream"`

	// 板书截图生成类型。0 不生成板书（默认）；1 全量模式；2 单页去重模式
	WhiteBoardSnapshotMode *uint64 `json:"WhiteBoardSnapshotMode,omitnil,omitempty" name:"WhiteBoardSnapshotMode"`

	// 字幕转写功能开关。可以有以下取值：
	// 0 不开启字幕转写功能（默认值）
	// 1 自动转写模式：上课自动开启，下课自动停止
	SubtitlesTranscription *uint64 `json:"SubtitlesTranscription,omitnil,omitempty" name:"SubtitlesTranscription"`

	// 录制文件合并开关。0 关闭 1 开启 注：只有在一节课多次启用手动录制时，此功能才有效
	RecordMerge *uint64 `json:"RecordMerge,omitnil,omitempty" name:"RecordMerge"`
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
	delete(f, "Guests")
	delete(f, "EndDelayTime")
	delete(f, "LiveType")
	delete(f, "RecordLiveUrl")
	delete(f, "EnableAutoStart")
	delete(f, "RecordBackground")
	delete(f, "RecordScene")
	delete(f, "RecordLang")
	delete(f, "RecordStream")
	delete(f, "WhiteBoardSnapshotMode")
	delete(f, "SubtitlesTranscription")
	delete(f, "RecordMerge")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRoomRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRoomResponseParams struct {
	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 用户ID列表。
	Users []*string `json:"Users,omitnil,omitempty" name:"Users"`
}

type CreateSupervisorRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 用户ID列表。
	Users []*string `json:"Users,omitnil,omitempty" name:"Users"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type CustomMsgContent struct {
	// 自定义消息数据。
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 自定义消息描述信息。
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 扩展字段。
	Ext *string `json:"Ext,omitnil,omitempty" name:"Ext"`
}

type CustomRecordInfo struct {
	// 开始时间
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	StopTime *uint64 `json:"StopTime,omitnil,omitempty" name:"StopTime"`

	// 总时长
	Duration *uint64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 文件格式
	FileFormat *string `json:"FileFormat,omitnil,omitempty" name:"FileFormat"`

	// 流url
	RecordUrl *string `json:"RecordUrl,omitnil,omitempty" name:"RecordUrl"`

	// 流大小
	RecordSize *uint64 `json:"RecordSize,omitnil,omitempty" name:"RecordSize"`

	// 流ID
	VideoId *string `json:"VideoId,omitnil,omitempty" name:"VideoId"`

	// 任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

// Predefined struct for user
type DeleteAppCustomContentRequestParams struct {
	// 应用ID。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 指定需要删除的已设置的scene场景自定义元素，如果为空则删除应用下已设置的所有自定义元素。
	Scenes []*string `json:"Scenes,omitnil,omitempty" name:"Scenes"`
}

type DeleteAppCustomContentRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 指定需要删除的已设置的scene场景自定义元素，如果为空则删除应用下已设置的所有自定义元素。
	Scenes []*string `json:"Scenes,omitnil,omitempty" name:"Scenes"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	DocumentId *string `json:"DocumentId,omitnil,omitempty" name:"DocumentId"`
}

type DeleteDocumentRequest struct {
	*tchttp.BaseRequest
	
	// 文档ID。
	DocumentId *string `json:"DocumentId,omitnil,omitempty" name:"DocumentId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 成员列表，最大值200
	MemberIds []*string `json:"MemberIds,omitnil,omitempty" name:"MemberIds"`
}

type DeleteGroupMemberRequest struct {
	*tchttp.BaseRequest
	
	// 群组ID，联合群组无法删除群组成员
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 成员列表，最大值200
	MemberIds []*string `json:"MemberIds,omitnil,omitempty" name:"MemberIds"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
}

type DeleteGroupRequest struct {
	*tchttp.BaseRequest
	
	// 待删除群组ID列表
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	RoomId *int64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 低代码互动课堂的SdkAppId。
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
}

type DeleteRecordRequest struct {
	*tchttp.BaseRequest
	
	// 房间Id。
	RoomId *int64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 低代码互动课堂的SdkAppId。
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 课堂ID。
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

type DeleteRoomRequest struct {
	*tchttp.BaseRequest
	
	// 课堂ID。
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 用户ID列表
	Users []*string `json:"Users,omitnil,omitempty" name:"Users"`
}

type DeleteSupervisorRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 用户ID列表
	Users []*string `json:"Users,omitnil,omitempty" name:"Users"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type DeleteUserRequest struct {
	*tchttp.BaseRequest
	
	// 待删除用户的ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteWhiteBoardSnapshotRequestParams struct {
	// 课堂ID
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

type DeleteWhiteBoardSnapshotRequest struct {
	*tchttp.BaseRequest
	
	// 课堂ID
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

func (r *DeleteWhiteBoardSnapshotRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWhiteBoardSnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWhiteBoardSnapshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWhiteBoardSnapshotResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteWhiteBoardSnapshotResponse struct {
	*tchttp.BaseResponse
	Response *DeleteWhiteBoardSnapshotResponseParams `json:"Response"`
}

func (r *DeleteWhiteBoardSnapshotResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWhiteBoardSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAnswerListRequestParams struct {
	// 问题ID
	QuestionId *string `json:"QuestionId,omitnil,omitempty" name:"QuestionId"`

	// 1
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeAnswerListRequest struct {
	*tchttp.BaseRequest
	
	// 问题ID
	QuestionId *string `json:"QuestionId,omitnil,omitempty" name:"QuestionId"`

	// 1
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 房间提问答案列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnswerInfo []*AnswerInfo `json:"AnswerInfo,omitnil,omitempty" name:"AnswerInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 开发商ID
	DeveloperId *string `json:"DeveloperId,omitnil,omitempty" name:"DeveloperId"`
}

type DescribeAppDetailRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID。低代码互动课堂的SdkAppId。
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 开发商ID
	DeveloperId *string `json:"DeveloperId,omitnil,omitempty" name:"DeveloperId"`
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
	SdkAppId *string `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 应用配置
	AppConfig *AppConfig `json:"AppConfig,omitnil,omitempty" name:"AppConfig"`

	// 场景配置
	SceneConfig []*SceneItem `json:"SceneConfig,omitnil,omitempty" name:"SceneConfig"`

	// 转存配置
	TransferConfig *TransferItem `json:"TransferConfig,omitnil,omitempty" name:"TransferConfig"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 课堂Id。
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 分页查询当前页数，从1开始递增。
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 每页数据量，最大1000。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeCurrentMemberListRequest struct {
	*tchttp.BaseRequest
	
	// 课堂Id。
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 分页查询当前页数，从1开始递增。
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 每页数据量，最大1000。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 成员记录列表。
	MemberRecords []*MemberRecord `json:"MemberRecords,omitnil,omitempty" name:"MemberRecords"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	DeveloperId *string `json:"DeveloperId,omitnil,omitempty" name:"DeveloperId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	DocumentId *string `json:"DocumentId,omitnil,omitempty" name:"DocumentId"`
}

type DescribeDocumentRequest struct {
	*tchttp.BaseRequest
	
	// 文档Id（唯一id）
	DocumentId *string `json:"DocumentId,omitnil,omitempty" name:"DocumentId"`
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
	DocumentId *string `json:"DocumentId,omitnil,omitempty" name:"DocumentId"`

	// 文档原址url
	DocumentUrl *string `json:"DocumentUrl,omitnil,omitempty" name:"DocumentUrl"`

	// 文档名称
	DocumentName *string `json:"DocumentName,omitnil,omitempty" name:"DocumentName"`

	// 文档所有者UserId
	Owner *string `json:"Owner,omitnil,omitempty" name:"Owner"`

	// 应用Id
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 文档权限
	Permission *uint64 `json:"Permission,omitnil,omitempty" name:"Permission"`

	// 转码结果，无需转码为空，转码成功为结果url，转码失败为错误码
	TranscodeResult *string `json:"TranscodeResult,omitnil,omitempty" name:"TranscodeResult"`

	// 转码类型
	TranscodeType *uint64 `json:"TranscodeType,omitnil,omitempty" name:"TranscodeType"`

	// 转码进度， 0 - 100 表示（0% - 100%）
	TranscodeProgress *uint64 `json:"TranscodeProgress,omitnil,omitempty" name:"TranscodeProgress"`

	// 转码状态，0为无需转码，1为正在转码，2为转码失败，3为转码成功
	TranscodeState *uint64 `json:"TranscodeState,omitnil,omitempty" name:"TranscodeState"`

	// 转码失败后的错误信息
	TranscodeInfo *string `json:"TranscodeInfo,omitnil,omitempty" name:"TranscodeInfo"`

	// 文档类型
	DocumentType *string `json:"DocumentType,omitnil,omitempty" name:"DocumentType"`

	// 文档大小，单位：字节
	DocumentSize *uint64 `json:"DocumentSize,omitnil,omitempty" name:"DocumentSize"`

	// 更新的UNIX时间戳
	UpdateTime *uint64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 课件页数
	Pages *uint64 `json:"Pages,omitnil,omitempty" name:"Pages"`

	// 课件预览地址
	Preview *string `json:"Preview,omitnil,omitempty" name:"Preview"`

	// 文档的分辨率
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// 转码后文档的最小分辨率，和创建文档时传入的参数一致。
	MinScaleResolution *string `json:"MinScaleResolution,omitnil,omitempty" name:"MinScaleResolution"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 低代码互动课堂的SdkAppId
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 分页查询当前页数，从1开始递增，默认值为1
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 每页数据量，最大1000，默认值为100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 课件权限。
	// [0]：获取owner的私有课件；
	// [1]：获取owner的公开课件;
	// [0,1]：则获取owner的私有课件和公开课件；
	// [2]：获取owner的私有课件和所有人(包括owner)的公开课件。
	// 默认值为[2]
	Permission []*uint64 `json:"Permission,omitnil,omitempty" name:"Permission"`

	// 文档所有者的user_id，不填默认获取SdkAppId下所有课件
	Owner *string `json:"Owner,omitnil,omitempty" name:"Owner"`
}

type DescribeDocumentsByRoomRequest struct {
	*tchttp.BaseRequest
	
	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 低代码互动课堂的SdkAppId
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 分页查询当前页数，从1开始递增，默认值为1
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 每页数据量，最大1000，默认值为100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 课件权限。
	// [0]：获取owner的私有课件；
	// [1]：获取owner的公开课件;
	// [0,1]：则获取owner的私有课件和公开课件；
	// [2]：获取owner的私有课件和所有人(包括owner)的公开课件。
	// 默认值为[2]
	Permission []*uint64 `json:"Permission,omitnil,omitempty" name:"Permission"`

	// 文档所有者的user_id，不填默认获取SdkAppId下所有课件
	Owner *string `json:"Owner,omitnil,omitempty" name:"Owner"`
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
	Documents []*DocumentInfo `json:"Documents,omitnil,omitempty" name:"Documents"`

	// 符合查询条件文档总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	SchoolId *uint64 `json:"SchoolId,omitnil,omitempty" name:"SchoolId"`

	// 分页查询当前页数，从1开始递增
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 每页数据量，最大1000
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 课件权限。[0]：获取owner的私有课件；[1]：获取owner的公开课件; [0,1]：则获取owner的私有课件和公开课件；[2]：获取owner的私有课件和所有人(包括owner)的公开课件
	Permission []*uint64 `json:"Permission,omitnil,omitempty" name:"Permission"`

	// 课件所有者的user_id，不填默认获取school_id下所有课件
	Owner *string `json:"Owner,omitnil,omitempty" name:"Owner"`

	// 课件名称搜索词
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 课件id列表，从列表中查询，忽略错误的id
	DocumentId []*string `json:"DocumentId,omitnil,omitempty" name:"DocumentId"`
}

type DescribeDocumentsRequest struct {
	*tchttp.BaseRequest
	
	// 学校id
	SchoolId *uint64 `json:"SchoolId,omitnil,omitempty" name:"SchoolId"`

	// 分页查询当前页数，从1开始递增
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 每页数据量，最大1000
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 课件权限。[0]：获取owner的私有课件；[1]：获取owner的公开课件; [0,1]：则获取owner的私有课件和公开课件；[2]：获取owner的私有课件和所有人(包括owner)的公开课件
	Permission []*uint64 `json:"Permission,omitnil,omitempty" name:"Permission"`

	// 课件所有者的user_id，不填默认获取school_id下所有课件
	Owner *string `json:"Owner,omitnil,omitempty" name:"Owner"`

	// 课件名称搜索词
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 课件id列表，从列表中查询，忽略错误的id
	DocumentId []*string `json:"DocumentId,omitnil,omitempty" name:"DocumentId"`
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
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 文档信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Documents []*DocumentInfo `json:"Documents,omitnil,omitempty" name:"Documents"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 分页查询当前页数，默认从1开始递增。
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 每页数据量，默认20，最大1000。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 主讲人ID筛选群组，与MemberId有且只有一个,都传时以此字段获取
	TeacherId *string `json:"TeacherId,omitnil,omitempty" name:"TeacherId"`

	// 成员ID刷选群组，与TeacherId有且只有一个
	MemberId *string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type DescribeGroupListRequest struct {
	*tchttp.BaseRequest
	
	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 分页查询当前页数，默认从1开始递增。
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 每页数据量，默认20，最大1000。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 主讲人ID筛选群组，与MemberId有且只有一个,都传时以此字段获取
	TeacherId *string `json:"TeacherId,omitnil,omitempty" name:"TeacherId"`

	// 成员ID刷选群组，与TeacherId有且只有一个
	MemberId *string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
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
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 群组信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupInfos []*GroupInfo `json:"GroupInfos,omitnil,omitempty" name:"GroupInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 分页值，默认1
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 每页数据量，默认20，最大1000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeGroupMemberListRequest struct {
	*tchttp.BaseRequest
	
	// 群组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 分页值，默认1
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 每页数据量，默认20，最大1000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 查询成员列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemberIds []*string `json:"MemberIds,omitnil,omitempty" name:"MemberIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
}

type DescribeGroupRequest struct {
	*tchttp.BaseRequest
	
	// 群组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
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
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 群组名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 群主主讲人ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TeacherId *string `json:"TeacherId,omitnil,omitempty" name:"TeacherId"`

	// 群组类型
	// 0-基础群组
	// 1-组合群组，若为1时会返回子群组ID
	GroupType *uint64 `json:"GroupType,omitnil,omitempty" name:"GroupType"`

	// 子群组ID列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubGroupIds []*string `json:"SubGroupIds,omitnil,omitempty" name:"SubGroupIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeMarqueeRequestParams struct {
	// 学校ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 房间号
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

type DescribeMarqueeRequest struct {
	*tchttp.BaseRequest
	
	// 学校ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 房间号
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

func (r *DescribeMarqueeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMarqueeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMarqueeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMarqueeResponseParams struct {
	//  跑马灯类型：1为固定值，2为用户昵称，3为固定值+用户昵称，4为用户ID，5为originId+固定值，6为用户昵称（originId）
	MarqueeType *uint64 `json:"MarqueeType,omitnil,omitempty" name:"MarqueeType"`

	// 固定值内容（当MarqueeType=1或5，则展示固定值内容）
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 字体大小（数字，像素单位，范围：10到24）
	FontSize *uint64 `json:"FontSize,omitnil,omitempty" name:"FontSize"`

	// 字体粗细：1为粗体，0为细体
	FontWeight *uint64 `json:"FontWeight,omitnil,omitempty" name:"FontWeight"`

	// 字体颜色（十六进制颜色值）
	FontColor *string `json:"FontColor,omitnil,omitempty" name:"FontColor"`

	// 字体透明度（数字，范围 0.0 到 1.0）
	FontOpacity *float64 `json:"FontOpacity,omitnil,omitempty" name:"FontOpacity"`

	// 背景颜色（十六进制颜色值）
	BackgroundColor *string `json:"BackgroundColor,omitnil,omitempty" name:"BackgroundColor"`

	// 背景透明度（数字，范围 0.0 到 1.0）
	BackgroundOpacity *float64 `json:"BackgroundOpacity,omitnil,omitempty" name:"BackgroundOpacity"`

	// 显示方式：1为滚动，2为闪烁
	DisplayMode *uint64 `json:"DisplayMode,omitnil,omitempty" name:"DisplayMode"`

	// 停留时长（秒，整数，范围 1～10）
	Duration *uint64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 跑马灯个数：目前仅支持1或2, 对应显示单排或双排
	MarqueeCount *uint64 `json:"MarqueeCount,omitnil,omitempty" name:"MarqueeCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMarqueeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMarqueeResponseParams `json:"Response"`
}

func (r *DescribeMarqueeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMarqueeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQuestionListRequestParams struct {
	// 房间ID
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 分页查询当前页数，从1开始递增，默认值为1
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 分页查询当前页数，从1开始递增，默认值为1
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeQuestionListRequest struct {
	*tchttp.BaseRequest
	
	// 房间ID
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 分页查询当前页数，从1开始递增，默认值为1
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 分页查询当前页数，从1开始递增，默认值为1
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 房间问答问题列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	QuestionInfo []*QuestionInfo `json:"QuestionInfo,omitnil,omitempty" name:"QuestionInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeRecordRequestParams struct {
	// 学校ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 房间ID
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

type DescribeRecordRequest struct {
	*tchttp.BaseRequest
	
	// 学校ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 房间ID
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

func (r *DescribeRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordResponseParams struct {
	// 学校ID
	SchoolId *uint64 `json:"SchoolId,omitnil,omitempty" name:"SchoolId"`

	// 课堂ID
	ClassId *uint64 `json:"ClassId,omitnil,omitempty" name:"ClassId"`

	// 录制信息
	RecordInfo []*CustomRecordInfo `json:"RecordInfo,omitnil,omitempty" name:"RecordInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRecordResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordResponseParams `json:"Response"`
}

func (r *DescribeRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordStreamRequestParams struct {
	// 学校ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 房间ID
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

type DescribeRecordStreamRequest struct {
	*tchttp.BaseRequest
	
	// 学校ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 房间ID
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

func (r *DescribeRecordStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordStreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordStreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordStreamResponseParams struct {
	// 学校ID
	SchoolId *uint64 `json:"SchoolId,omitnil,omitempty" name:"SchoolId"`

	// 课堂ID
	ClassId *uint64 `json:"ClassId,omitnil,omitempty" name:"ClassId"`

	// 课堂类型
	ClassType *uint64 `json:"ClassType,omitnil,omitempty" name:"ClassType"`

	// 用户流信息
	StreamInfo []*SingleStreamInfo `json:"StreamInfo,omitnil,omitempty" name:"StreamInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRecordStreamResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordStreamResponseParams `json:"Response"`
}

func (r *DescribeRecordStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordStreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordTaskRequestParams struct {
	// 学校ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 房间ID
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

type DescribeRecordTaskRequest struct {
	*tchttp.BaseRequest
	
	// 学校ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 房间ID
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

func (r *DescribeRecordTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordTaskResponseParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRecordTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordTaskResponseParams `json:"Response"`
}

func (r *DescribeRecordTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoomForbiddenUserRequestParams struct {
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

type DescribeRoomForbiddenUserRequest struct {
	*tchttp.BaseRequest
	
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`
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
	MutedAccountList []*MutedAccountList `json:"MutedAccountList,omitnil,omitempty" name:"MutedAccountList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 课堂Id。
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 请求RTMP推流链接，0：否，1：是，默认为0。
	RTMPStreamingURL *uint64 `json:"RTMPStreamingURL,omitnil,omitempty" name:"RTMPStreamingURL"`
}

type DescribeRoomRequest struct {
	*tchttp.BaseRequest
	
	// 课堂Id。
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 请求RTMP推流链接，0：否，1：是，默认为0。
	RTMPStreamingURL *uint64 `json:"RTMPStreamingURL,omitnil,omitempty" name:"RTMPStreamingURL"`
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
	delete(f, "RTMPStreamingURL")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRoomRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoomResponseParams struct {
	// 课堂名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 预定的课堂开始时间，unix时间戳（秒）。
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 预定的课堂结束时间，unix时间戳（秒）。
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 老师的UserId。
	TeacherId *string `json:"TeacherId,omitnil,omitempty" name:"TeacherId"`

	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 观看类型。互动观看 （默认）	
	AudienceType *uint64 `json:"AudienceType,omitnil,omitempty" name:"AudienceType"`

	// 头像区域，摄像头视频画面的分辨率。可以有如下取值：
	// 1 标清
	// 2 高清
	// 3 全高清
	Resolution *uint64 `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// 设置课堂同时最大可与老师进行连麦互动的人数，该参数支持正式上课/开播前调用修改房间修改。小班课取值范围[0,16]，大班课取值范围[0,1]，当取值为0时表示当前课堂/直播，不支持连麦互动。
	MaxMicNumber *uint64 `json:"MaxMicNumber,omitnil,omitempty" name:"MaxMicNumber"`

	// 进入课堂时是否自动连麦。可以有以下取值：
	// 0 不自动连麦（需要手动申请上麦，默认值）
	// 1 自动连麦
	AutoMic *uint64 `json:"AutoMic,omitnil,omitempty" name:"AutoMic"`

	// 高音质模式。可以有以下取值：
	// 0 不开启高音质（默认值）
	// 1 开启高音质
	AudioQuality *uint64 `json:"AudioQuality,omitnil,omitempty" name:"AudioQuality"`

	// 课堂子类型，可以有以下取值：videodoc 文档+视频video 纯视频
	SubType *string `json:"SubType,omitnil,omitempty" name:"SubType"`

	// 上课后是否禁止自动录制。可以有以下取值：
	// 0 不禁止录制（自动开启录制，默认值）
	// 1 禁止录制
	// 注：如果该配置取值为0，录制将从上课后开始，课堂结束后停止。
	DisableRecord *uint64 `json:"DisableRecord,omitnil,omitempty" name:"DisableRecord"`

	// 助教UserId列表。
	Assistants []*string `json:"Assistants,omitnil,omitempty" name:"Assistants"`

	// 录制地址（协议为https)。仅在房间结束后存在。
	RecordUrl *string `json:"RecordUrl,omitnil,omitempty" name:"RecordUrl"`

	// 课堂状态。0为未开始，1为已开始，2为已结束，3为已过期。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 课堂绑定的群组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 打开学生麦克风/摄像头的授权开关
	EnableDirectControl *uint64 `json:"EnableDirectControl,omitnil,omitempty" name:"EnableDirectControl"`

	// 开启专注模式。
	// 0 收看全部角色音视频(默认)
	// 1 只看老师和助教
	InteractionMode *int64 `json:"InteractionMode,omitnil,omitempty" name:"InteractionMode"`

	// 横竖屏。0：横屏开播（默认值）; 1：竖屏开播，当前仅支持移动端的纯视频类型
	VideoOrientation *uint64 `json:"VideoOrientation,omitnil,omitempty" name:"VideoOrientation"`

	// 该课堂是否开启了课后评分功能。0：未开启  1：开启
	IsGradingRequiredPostClass *int64 `json:"IsGradingRequiredPostClass,omitnil,omitempty" name:"IsGradingRequiredPostClass"`

	// 课堂类型: 0 小班课（默认值）; 1 大班课; 2 1V1 (预留参数，暂未开放); 3 圆桌会议 注：大班课的布局(layout)只有三分屏
	RoomType *int64 `json:"RoomType,omitnil,omitempty" name:"RoomType"`

	// 录制时长
	VideoDuration *uint64 `json:"VideoDuration,omitnil,omitempty" name:"VideoDuration"`

	// 拖堂时间：单位分钟，0为不限制(默认值), -1为不能拖堂，大于0为拖堂的时间，最大值120分钟
	EndDelayTime *int64 `json:"EndDelayTime,omitnil,omitempty" name:"EndDelayTime"`

	// 直播类型：0 常规（默认）1 伪直播 2 RTMP推流直播
	LiveType *uint64 `json:"LiveType,omitnil,omitempty" name:"LiveType"`

	// 伪直播链接
	RecordLiveUrl *string `json:"RecordLiveUrl,omitnil,omitempty" name:"RecordLiveUrl"`

	// 是否自动开始上课：0 不自动上课（默认） 1 自动上课 live_type=1的时候有效
	EnableAutoStart *uint64 `json:"EnableAutoStart,omitnil,omitempty" name:"EnableAutoStart"`

	// 录制文件背景图片，支持png、jpg、jpeg、bmp格式，暂不支持透明通道
	RecordBackground *string `json:"RecordBackground,omitnil,omitempty" name:"RecordBackground"`

	// RTMP推流链接
	RTMPStreamingURL *string `json:"RTMPStreamingURL,omitnil,omitempty" name:"RTMPStreamingURL"`

	// 录制自定义场景。注意：仅recordlayout=9的时候此参数有效。需注意各类参数配置正确能够生效。不然会造成录制失败，失败后无法补救。数据内容为用户自定义场景参数，数据格式为json键值对方式，其中键值对的value为string类型。自定义场景参数的含义。如下：     scene：自定义js/css对应的场景值。如scene=recordScene，会加载 recordScene 场景对应的 js/css，这样就可以自定义录制页面的元素。     lng：录制页面对应的语种。如lng=en，则录制界面为en。（枚举值：en,zh，zh-TW，jp，ar，kr，vi）     customToken：录制页面中涉及客户自己的服务需要鉴权时进行配置。一般情况下，无需配置。
	RecordScene *string `json:"RecordScene,omitnil,omitempty" name:"RecordScene"`

	// 录制自定义语言，仅recordlayout=9的时候此参数有效
	RecordLang *string `json:"RecordLang,omitnil,omitempty" name:"RecordLang"`

	// 录制类型 0 仅录制混流（默认） ;1 录制混流+单流，该模式下除混流录制基础上，分别录制老师、台上学生的音视频流，每路录制都会产生相应的录制费用 。示例：0
	RecordStream *uint64 `json:"RecordStream,omitnil,omitempty" name:"RecordStream"`

	// 录制模板。房间子类型为视频+白板（SubType=videodoc）时默认为3，房间子类型为纯视频（SubType=video）时默认为0。录制模板枚举值参考：https://cloud.tencent.com/document/product/1639/89744
	RecordLayout *uint64 `json:"RecordLayout,omitnil,omitempty" name:"RecordLayout"`

	// 板书截图生成类型。0 不生成板书；1 全量模式；2 单页去重模式
	WhiteBoardSnapshotMode *uint64 `json:"WhiteBoardSnapshotMode,omitnil,omitempty" name:"WhiteBoardSnapshotMode"`

	// 字幕转写功能开关：0关闭，1开启，默认关闭
	SubtitlesTranscription *uint64 `json:"SubtitlesTranscription,omitnil,omitempty" name:"SubtitlesTranscription"`

	// 嘉宾Id列表。当圆桌会议模式（RoomType==3）时生效
	Guests []*string `json:"Guests,omitnil,omitempty" name:"Guests"`

	// 录制文件合并开关。0 关闭 1 开启 注：只有在一节课多次启用手动录制时，此功能才有效
	RecordMerge *uint64 `json:"RecordMerge,omitnil,omitempty" name:"RecordMerge"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 课堂Id。
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 分页查询当前页数，从1开始递增。
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 每页数据量，最大1000。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeRoomStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 课堂Id。
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 分页查询当前页数，从1开始递增。
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 每页数据量，最大1000。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	PeakMemberNumber *uint64 `json:"PeakMemberNumber,omitnil,omitempty" name:"PeakMemberNumber"`

	// 累计在线人数。
	MemberNumber *uint64 `json:"MemberNumber,omitnil,omitempty" name:"MemberNumber"`

	// 记录总数。包含进入房间或者应到未到的。
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 成员记录列表。
	MemberRecords []*MemberRecord `json:"MemberRecords,omitnil,omitempty" name:"MemberRecords"`

	// 秒级unix时间戳，实际房间开始时间。
	RealStartTime *uint64 `json:"RealStartTime,omitnil,omitempty" name:"RealStartTime"`

	// 秒级unix时间戳，实际房间结束时间。
	RealEndTime *uint64 `json:"RealEndTime,omitnil,omitempty" name:"RealEndTime"`

	// 课堂消息总数。
	MessageCount *uint64 `json:"MessageCount,omitnil,omitempty" name:"MessageCount"`

	// 课堂连麦总数。
	MicCount *uint64 `json:"MicCount,omitnil,omitempty" name:"MicCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 分页查询当前页数，从1开始递增
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 默认是10条
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeScoreListRequest struct {
	*tchttp.BaseRequest
	
	// 课堂ID
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 分页查询当前页数，从1开始递增
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 默认是10条
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 课堂评分列表
	Scores []*ClassScoreItem `json:"Scores,omitnil,omitempty" name:"Scores"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 分页，默认值为1
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 分页数据限制，默认值为20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeSdkAppIdUsersRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 分页，默认值为1
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 分页数据限制，默认值为20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 当前获取用户信息数组列表
	Users []*UserInfo `json:"Users,omitnil,omitempty" name:"Users"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 每页数据量，最大100。 不填默认20.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页查询当前页数，从1开始递增，不填默认为1。
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`
}

type DescribeSupervisorsRequest struct {
	*tchttp.BaseRequest
	
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 每页数据量，最大100。 不填默认20.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页查询当前页数，从1开始递增，不填默认为1。
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`
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
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 分页查询当前页数
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 当前页数据量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 巡课列表
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeUserDetailRequestParams struct {
	// 用户id。支持通过 user_id 或 OriginId 查询用户信息，优先使用 user_id 进行查询。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户在客户系统的Id。支持通过 user_id 或 OriginId 查询用户信息，优先使用 user_id 进行查询（UserId不为空时，OriginId不生效）。
	OriginId *string `json:"OriginId,omitnil,omitempty" name:"OriginId"`
}

type DescribeUserDetailRequest struct {
	*tchttp.BaseRequest
	
	// 用户id。支持通过 user_id 或 OriginId 查询用户信息，优先使用 user_id 进行查询。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户在客户系统的Id。支持通过 user_id 或 OriginId 查询用户信息，优先使用 user_id 进行查询（UserId不为空时，OriginId不生效）。
	OriginId *string `json:"OriginId,omitnil,omitempty" name:"OriginId"`
}

func (r *DescribeUserDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "OriginId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserDetailResponseParams struct {
	// 当前获取用户信息数组列表
	Users []*UserInfo `json:"Users,omitnil,omitempty" name:"Users"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserDetailResponseParams `json:"Response"`
}

func (r *DescribeUserDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserRequestParams struct {
	// 用户id。支持通过 user_id 或 OriginId 查询用户信息，优先使用 user_id 进行查询。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户在客户系统的Id。支持通过 user_id 或 OriginId 查询用户信息，优先使用 user_id 进行查询（UserId不为空时，OriginId不生效）。
	OriginId *string `json:"OriginId,omitnil,omitempty" name:"OriginId"`
}

type DescribeUserRequest struct {
	*tchttp.BaseRequest
	
	// 用户id。支持通过 user_id 或 OriginId 查询用户信息，优先使用 user_id 进行查询。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户在客户系统的Id。支持通过 user_id 或 OriginId 查询用户信息，优先使用 user_id 进行查询（UserId不为空时，OriginId不生效）。
	OriginId *string `json:"OriginId,omitnil,omitempty" name:"OriginId"`
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
	delete(f, "OriginId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserResponseParams struct {
	// 应用Id。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 用户Id。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户昵称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 用户头像Url。
	Avatar *string `json:"Avatar,omitnil,omitempty" name:"Avatar"`

	// 用户在客户系统的Id
	OriginId *string `json:"OriginId,omitnil,omitempty" name:"OriginId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type DescribeWhiteBoardSnapshotRequestParams struct {
	// 房间ID
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

type DescribeWhiteBoardSnapshotRequest struct {
	*tchttp.BaseRequest
	
	// 房间ID
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

func (r *DescribeWhiteBoardSnapshotRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWhiteBoardSnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWhiteBoardSnapshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWhiteBoardSnapshotResponseParams struct {
	// 板书截图生成类型。0 不生成板书；1 全量模式；2 单页去重模式
	WhiteBoardSnapshotMode *uint64 `json:"WhiteBoardSnapshotMode,omitnil,omitempty" name:"WhiteBoardSnapshotMode"`

	// 板书任务状态，0：未开始，1：进行中，2：失败，3：成功，4：已删除
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 板书截图链接
	Result []*string `json:"Result,omitnil,omitempty" name:"Result"`

	// 总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWhiteBoardSnapshotResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWhiteBoardSnapshotResponseParams `json:"Response"`
}

func (r *DescribeWhiteBoardSnapshotResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWhiteBoardSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DocumentInfo struct {
	// 文档Id
	DocumentId *string `json:"DocumentId,omitnil,omitempty" name:"DocumentId"`

	// 文档原址url
	DocumentUrl *string `json:"DocumentUrl,omitnil,omitempty" name:"DocumentUrl"`

	// 文档名称
	DocumentName *string `json:"DocumentName,omitnil,omitempty" name:"DocumentName"`

	// 文档所有者UserId
	Owner *string `json:"Owner,omitnil,omitempty" name:"Owner"`

	// 应用Id
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 文档权限，0：私有课件 1：公共课件
	Permission *uint64 `json:"Permission,omitnil,omitempty" name:"Permission"`

	// 转码结果，无需转码为空，转码成功为结果url，转码失败为错误码
	TranscodeResult *string `json:"TranscodeResult,omitnil,omitempty" name:"TranscodeResult"`

	// 转码类型
	TranscodeType *uint64 `json:"TranscodeType,omitnil,omitempty" name:"TranscodeType"`

	// 转码进度， 0 - 100 表示（0% - 100%）
	TranscodeProgress *uint64 `json:"TranscodeProgress,omitnil,omitempty" name:"TranscodeProgress"`

	// 转码状态，0为无需转码，1为正在转码，2为转码失败，3为转码成功
	TranscodeState *uint64 `json:"TranscodeState,omitnil,omitempty" name:"TranscodeState"`

	// 转码失败后的错误信息
	TranscodeInfo *string `json:"TranscodeInfo,omitnil,omitempty" name:"TranscodeInfo"`

	// 文档类型
	DocumentType *string `json:"DocumentType,omitnil,omitempty" name:"DocumentType"`

	// 文档大小，单位：字节
	DocumentSize *uint64 `json:"DocumentSize,omitnil,omitempty" name:"DocumentSize"`

	// 更新的UNIX时间戳
	UpdateTime *uint64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 课件页数
	Pages *uint64 `json:"Pages,omitnil,omitempty" name:"Pages"`

	// 宽，仅在静态转码的课件有效
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 高，仅在静态转码的课件有效
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 封面，仅转码的课件会生成封面
	Cover *string `json:"Cover,omitnil,omitempty" name:"Cover"`

	// 课件预览地址
	Preview *string `json:"Preview,omitnil,omitempty" name:"Preview"`

	// 文档的分辨率
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// 转码后文档的最小分辨率，和创建文档时传入的参数一致。
	MinScaleResolution *string `json:"MinScaleResolution,omitnil,omitempty" name:"MinScaleResolution"`
}

// Predefined struct for user
type EndRoomRequestParams struct {
	// 课堂ID
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

type EndRoomRequest struct {
	*tchttp.BaseRequest
	
	// 课堂ID
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 事件发生的用户。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户设备类型。0: Unknown; 1: Windows; 2: macOS; 3: Android; 4: iOS; 5: Web; 6: Mobile webpage; 7: Weixin Mini Program.
	Device *uint64 `json:"Device,omitnil,omitempty" name:"Device"`

	// 录制时长。单位：秒
	Duration *uint64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 录制文件大小
	RecordSize *uint64 `json:"RecordSize,omitnil,omitempty" name:"RecordSize"`

	// 录制url
	RecordUrl *string `json:"RecordUrl,omitnil,omitempty" name:"RecordUrl"`
}

type EventInfo struct {
	// 事件发生的秒级unix时间戳。
	Timestamp *uint64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

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
	EventType *string `json:"EventType,omitnil,omitempty" name:"EventType"`

	// 事件详细内容，包含房间号,成员类型事件包含用户Id。
	EventData *EventDataInfo `json:"EventData,omitnil,omitempty" name:"EventData"`
}

type FaceMsgContent struct {
	// 表情索引，用户自定义。
	Index *int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 额外数据。
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`
}

// Predefined struct for user
type ForbidSendMsgRequestParams struct {
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 课堂ID
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 需要禁言的用户账号，最多支持500个账号
	MembersAccount []*string `json:"MembersAccount,omitnil,omitempty" name:"MembersAccount"`

	// 需禁言时间，单位为秒，为0时表示取消禁言，4294967295为永久禁言。
	MuteTime *uint64 `json:"MuteTime,omitnil,omitempty" name:"MuteTime"`
}

type ForbidSendMsgRequest struct {
	*tchttp.BaseRequest
	
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 课堂ID
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 需要禁言的用户账号，最多支持500个账号
	MembersAccount []*string `json:"MembersAccount,omitnil,omitempty" name:"MembersAccount"`

	// 需禁言时间，单位为秒，为0时表示取消禁言，4294967295为永久禁言。
	MuteTime *uint64 `json:"MuteTime,omitnil,omitempty" name:"MuteTime"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 课堂Id。
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 应用Id。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 起始页，1开始。keyword为空时有效。
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 每页个数。keyword为空时有效。一次性最多100条。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

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
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

type GetRoomEventRequest struct {
	*tchttp.BaseRequest
	
	// 课堂Id。
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 应用Id。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 起始页，1开始。keyword为空时有效。
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 每页个数。keyword为空时有效。一次性最多100条。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

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
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
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
	// 该课堂的事件总数，keyword搜索不影响该值。
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 详细事件内容。包含相应的类型、发生的时间戳。
	Events []*EventInfo `json:"Events,omitnil,omitempty" name:"Events"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 课堂Id。	
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 消息序列。获取该序列以前的消息(不包含该seq消息)
	Seq *int64 `json:"Seq,omitnil,omitempty" name:"Seq"`

	// 消息拉取的条数。最大数量不能超过套餐包限制。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 请求消息的userId
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type GetRoomMessageRequest struct {
	*tchttp.BaseRequest
	
	// 低代码互动课堂的SdkAppId。
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 课堂Id。	
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 消息序列。获取该序列以前的消息(不包含该seq消息)
	Seq *int64 `json:"Seq,omitnil,omitempty" name:"Seq"`

	// 消息拉取的条数。最大数量不能超过套餐包限制。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 请求消息的userId
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
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
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetRoomMessageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRoomMessageResponseParams struct {
	// 消息列表
	Messages []*MessageList `json:"Messages,omitnil,omitempty" name:"Messages"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 开始时间。默认以当前时间减去半小时作为开始时间。
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。默认以当前时间加上半小时作为结束时间。
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页查询当前页数，从1开始递增
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 默认10条，最大上限为100条
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 课堂状态。默认展示所有课堂，0为未开始，1为正在上课，2为已结束，3为已过期
	Status []*uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type GetRoomsRequest struct {
	*tchttp.BaseRequest
	
	// 低代码平台的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 开始时间。默认以当前时间减去半小时作为开始时间。
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。默认以当前时间加上半小时作为结束时间。
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页查询当前页数，从1开始递增
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 默认10条，最大上限为100条
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 课堂状态。默认展示所有课堂，0为未开始，1为正在上课，2为已结束，3为已过期
	Status []*uint64 `json:"Status,omitnil,omitempty" name:"Status"`
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
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 课堂列表
	Rooms []*RoomItem `json:"Rooms,omitnil,omitempty" name:"Rooms"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
}

type GetWatermarkRequest struct {
	*tchttp.BaseRequest
	
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
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
	TeacherLogo *WatermarkConfig `json:"TeacherLogo,omitnil,omitempty" name:"TeacherLogo"`

	// 白板区域的水印参数配置
	BoardLogo *WatermarkConfig `json:"BoardLogo,omitnil,omitempty" name:"BoardLogo"`

	// 背景图片配置
	BackgroundPicture *BackgroundPictureConfig `json:"BackgroundPicture,omitnil,omitempty" name:"BackgroundPicture"`

	// 文字水印配置
	Text *TextMarkConfig `json:"Text,omitnil,omitempty" name:"Text"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 群组主讲人ID
	TeacherId *string `json:"TeacherId,omitnil,omitempty" name:"TeacherId"`
}

type GroupInfo struct {
	// 群组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 群组名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 群组主讲人ID
	TeacherId *string `json:"TeacherId,omitnil,omitempty" name:"TeacherId"`

	// 群组类型 
	// 0-基础群组 
	// 1-组合群组，若为1时会返回子群组ID列表
	GroupType *uint64 `json:"GroupType,omitnil,omitempty" name:"GroupType"`

	// 子群组ID列表，如有。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubGroupIds *string `json:"SubGroupIds,omitnil,omitempty" name:"SubGroupIds"`
}

type ImageInfo struct {
	// 图片类型：
	// 1-原图
	// 2-大图
	// 3-缩略图
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 图片数据大小，单位：字节。
	Size *uint64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 图片宽度，单位为像素。
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 图片高度，单位为像素。
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 图片下载地址。
	URL *string `json:"URL,omitnil,omitempty" name:"URL"`
}

type ImageMsgContent struct {
	// 图片的唯一标识，客户端用于索引图片的键值。
	UUID *string `json:"UUID,omitnil,omitempty" name:"UUID"`

	// 图片格式。
	// JPG = 1
	// GIF = 2
	// PNG = 3
	// BMP = 4
	// 其他 = 255
	ImageFormat *uint64 `json:"ImageFormat,omitnil,omitempty" name:"ImageFormat"`

	// 图片信息
	ImageInfoList []*ImageInfo `json:"ImageInfoList,omitnil,omitempty" name:"ImageInfoList"`
}

// Predefined struct for user
type KickUserFromRoomRequestParams struct {
	// 课堂Id。
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 低代码平台的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 需要踢出成员Id
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 踢出类型：
	// 1：临时踢出，可以使用Duration参数指定污点时间，污点时间间隔内用户无法进入房间。
	// 2：永久踢出
	KickType *uint64 `json:"KickType,omitnil,omitempty" name:"KickType"`

	// 污点时间(单位秒)，KickType = 1时生效，默认为0
	Duration *uint64 `json:"Duration,omitnil,omitempty" name:"Duration"`
}

type KickUserFromRoomRequest struct {
	*tchttp.BaseRequest
	
	// 课堂Id。
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 低代码平台的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 需要踢出成员Id
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 踢出类型：
	// 1：临时踢出，可以使用Duration参数指定污点时间，污点时间间隔内用户无法进入房间。
	// 2：永久踢出
	KickType *uint64 `json:"KickType,omitnil,omitempty" name:"KickType"`

	// 污点时间(单位秒)，KickType = 1时生效，默认为0
	Duration *uint64 `json:"Duration,omitnil,omitempty" name:"Duration"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 用户在客户系统的Id，需要在同一应用下唯一。
	OriginId *string `json:"OriginId,omitnil,omitempty" name:"OriginId"`
}

type LoginOriginIdRequest struct {
	*tchttp.BaseRequest
	
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 用户在客户系统的Id，需要在同一应用下唯一。
	OriginId *string `json:"OriginId,omitnil,omitempty" name:"OriginId"`
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
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 登录/注册成功后返回登录态token。有效期7天。
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type LoginUserRequest struct {
	*tchttp.BaseRequest
	
	// 注册获取的用户id。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
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
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 注册成功后返回登录态token，有效期7天。token过期后可以通过调用“登录”或“源账号登录”进行更新。
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户名称。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 在线时长，单位秒。
	PresentTime *uint64 `json:"PresentTime,omitnil,omitempty" name:"PresentTime"`

	// 是否开启摄像头。
	Camera *uint64 `json:"Camera,omitnil,omitempty" name:"Camera"`

	// 是否开启麦克风。
	Mic *uint64 `json:"Mic,omitnil,omitempty" name:"Mic"`

	// 是否禁言。
	Silence *uint64 `json:"Silence,omitnil,omitempty" name:"Silence"`

	// 回答问题数量。
	AnswerQuestions *uint64 `json:"AnswerQuestions,omitnil,omitempty" name:"AnswerQuestions"`

	// 举手数量。
	HandUps *uint64 `json:"HandUps,omitnil,omitempty" name:"HandUps"`

	// 首次进入房间的unix时间戳。
	FirstJoinTimestamp *uint64 `json:"FirstJoinTimestamp,omitnil,omitempty" name:"FirstJoinTimestamp"`

	// 最后一次退出房间的unix时间戳。
	LastQuitTimestamp *uint64 `json:"LastQuitTimestamp,omitnil,omitempty" name:"LastQuitTimestamp"`

	// 奖励次数。
	Rewords *uint64 `json:"Rewords,omitnil,omitempty" name:"Rewords"`

	// 用户IP。
	IPAddress *string `json:"IPAddress,omitnil,omitempty" name:"IPAddress"`

	// 用户位置信息。
	Location *string `json:"Location,omitnil,omitempty" name:"Location"`

	// 用户设备平台信息。0:unknown  1:windows  2:mac  3:android  4:ios  5:web   6:h5   7:miniprogram （小程序）
	Device *int64 `json:"Device,omitnil,omitempty" name:"Device"`

	// 每个成员上麦次数。
	PerMemberMicCount *int64 `json:"PerMemberMicCount,omitnil,omitempty" name:"PerMemberMicCount"`

	// 每个成员发送消息数量。
	PerMemberMessageCount *int64 `json:"PerMemberMessageCount,omitnil,omitempty" name:"PerMemberMessageCount"`

	// 用户角色。0代表学生；1代表老师； 2助教；3巡课。
	Role *uint64 `json:"Role,omitnil,omitempty" name:"Role"`

	// 上课班号
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 子上课班号
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubGroupId []*string `json:"SubGroupId,omitnil,omitempty" name:"SubGroupId"`

	// 用户的上台状态
	Stage *int64 `json:"Stage,omitnil,omitempty" name:"Stage"`

	// 用户状态。0为未到，1为在线，2为离线，3为被踢，4为永久被踢，5为暂时掉线
	CurrentState *uint64 `json:"CurrentState,omitnil,omitempty" name:"CurrentState"`
}

type MessageItem struct {
	// 消息类型。0表示文本消息，1表示图片消息
	MessageType *int64 `json:"MessageType,omitnil,omitempty" name:"MessageType"`

	// 文本消息内容。message type为0时有效。
	TextMessage *string `json:"TextMessage,omitnil,omitempty" name:"TextMessage"`

	// 图片消息URL。 message type为1时有效。
	ImageMessage *string `json:"ImageMessage,omitnil,omitempty" name:"ImageMessage"`

	// 自定义消息内容。message type为2时有效。
	CustomMessage *CustomMsgContent `json:"CustomMessage,omitnil,omitempty" name:"CustomMessage"`
}

type MessageList struct {
	// 消息时间戳
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 消息发送者
	FromAccount *string `json:"FromAccount,omitnil,omitempty" name:"FromAccount"`

	// 消息序列号，当前课堂内唯一且单调递增
	Seq *int64 `json:"Seq,omitnil,omitempty" name:"Seq"`

	// 历史消息列表
	MessageBody []*MessageItem `json:"MessageBody,omitnil,omitempty" name:"MessageBody"`
}

// Predefined struct for user
type ModifyAppRequestParams struct {
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 回调地址。
	Callback *string `json:"Callback,omitnil,omitempty" name:"Callback"`

	// 回调key。
	CallbackKey *string `json:"CallbackKey,omitnil,omitempty" name:"CallbackKey"`

	// 转存id
	TransferId *string `json:"TransferId,omitnil,omitempty" name:"TransferId"`

	// 转存地址
	TransferUrl *string `json:"TransferUrl,omitnil,omitempty" name:"TransferUrl"`
}

type ModifyAppRequest struct {
	*tchttp.BaseRequest
	
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 回调地址。
	Callback *string `json:"Callback,omitnil,omitempty" name:"Callback"`

	// 回调key。
	CallbackKey *string `json:"CallbackKey,omitnil,omitempty" name:"CallbackKey"`

	// 转存id
	TransferId *string `json:"TransferId,omitnil,omitempty" name:"TransferId"`

	// 转存地址
	TransferUrl *string `json:"TransferUrl,omitnil,omitempty" name:"TransferUrl"`
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
	delete(f, "TransferId")
	delete(f, "TransferUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAppResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 默认绑定主讲老师ID
	TeacherId *string `json:"TeacherId,omitnil,omitempty" name:"TeacherId"`

	// 待修改的群组名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`
}

type ModifyGroupRequest struct {
	*tchttp.BaseRequest
	
	// 需要修改的群组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 低代码平台应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 默认绑定主讲老师ID
	TeacherId *string `json:"TeacherId,omitnil,omitempty" name:"TeacherId"`

	// 待修改的群组名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 低代码互动课堂的SdkAppId
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 预定的房间开始时间，unix时间戳（秒）。直播开始后不允许修改。
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 预定的房间结束时间，unix时间戳（秒）。直播开始后不允许修改。
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 老师ID。直播开始后不允许修改。
	TeacherId *string `json:"TeacherId,omitnil,omitempty" name:"TeacherId"`

	// 房间名称。
	// 字符数不超过256
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分辨率。可以有如下取值：
	// 1 标清
	// 2 高清
	// 3 全高清
	// 直播开始后不允许修改。
	Resolution *uint64 `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// 设置房间/课堂同时最大可与老师进行连麦互动的人数，该参数支持正式上课/开播前调用修改房间修改。小班课取值范围[0,16]，大班课取值范围[0,1]，当取值为0时表示当前课堂/直播，不支持连麦互动。
	MaxMicNumber *uint64 `json:"MaxMicNumber,omitnil,omitempty" name:"MaxMicNumber"`

	// 进入房间时是否自动连麦。可以有以下取值：
	// 0 不自动连麦（默认值）
	// 1 自动连麦
	// 直播开始后不允许修改。
	AutoMic *uint64 `json:"AutoMic,omitnil,omitempty" name:"AutoMic"`

	// 高音质模式。可以有以下取值：
	// 0 不开启高音质（默认值）
	// 1 开启高音质
	// 直播开始后不允许修改。
	AudioQuality *uint64 `json:"AudioQuality,omitnil,omitempty" name:"AudioQuality"`

	// 房间子类型，可以有以下取值：
	// videodoc 文档+视频
	// video 纯视频
	// 直播开始后不允许修改。
	SubType *string `json:"SubType,omitnil,omitempty" name:"SubType"`

	// 禁止录制。可以有以下取值：
	// 0 不禁止录制（默认值）
	// 1 禁止录制
	// 直播开始后不允许修改。
	DisableRecord *uint64 `json:"DisableRecord,omitnil,omitempty" name:"DisableRecord"`

	// 助教Id列表。直播开始后不允许修改。
	Assistants []*string `json:"Assistants,omitnil,omitempty" name:"Assistants"`

	// 房间绑定的群组ID。直播开始后不允许修改。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 打开学生麦克风/摄像头的授权开关。直播开始后不允许修改。
	EnableDirectControl *uint64 `json:"EnableDirectControl,omitnil,omitempty" name:"EnableDirectControl"`

	// 开启专注模式。
	// 0 收看全部角色音视频(默认)
	// 1 只看老师和助教
	InteractionMode *uint64 `json:"InteractionMode,omitnil,omitempty" name:"InteractionMode"`

	// 横竖屏。0：横屏开播（默认值）; 1：竖屏开播，当前仅支持移动端的纯视频类型
	VideoOrientation *uint64 `json:"VideoOrientation,omitnil,omitempty" name:"VideoOrientation"`

	// 开启课后评分。 0：不开启(默认)  1：开启
	IsGradingRequiredPostClass *uint64 `json:"IsGradingRequiredPostClass,omitnil,omitempty" name:"IsGradingRequiredPostClass"`

	// 房间类型: 0 小班课（默认值）; 1 大班课; 2 1V1 （预留参数、暂未开放)
	// 注：大班课的布局(layout)只有三分屏
	RoomType *uint64 `json:"RoomType,omitnil,omitempty" name:"RoomType"`

	// 录制模板。仅可修改还未开始的房间。录制模板枚举值参考：https://cloud.tencent.com/document/product/1639/89744
	RecordLayout *uint64 `json:"RecordLayout,omitnil,omitempty" name:"RecordLayout"`

	// 拖堂时间：单位分钟，0为不限制(默认值), -1为不能拖堂，大于0为拖堂的时间，最大值120分钟
	EndDelayTime *int64 `json:"EndDelayTime,omitnil,omitempty" name:"EndDelayTime"`

	// 直播方式：0 常规模式（默认）1 回放直播模式（伪直播）。 目前支持从回放直播模式（伪直播）改为常规模式，不支持从常规模式改为回放直播模式（伪直播）
	LiveType *uint64 `json:"LiveType,omitnil,omitempty" name:"LiveType"`

	// 伪直播链接
	RecordLiveUrl *string `json:"RecordLiveUrl,omitnil,omitempty" name:"RecordLiveUrl"`

	// 是否自动开始上课：0 不自动上课（默认） 1 自动上课 live_type=1的时候有效
	EnableAutoStart *uint64 `json:"EnableAutoStart,omitnil,omitempty" name:"EnableAutoStart"`

	// 录制自定义场景，仅recordlayout=9的时候此参数有效,数据内容为用户自定义场景参数，数据格式为json键值对方式，其中键值对的value为string类型。
	RecordScene *string `json:"RecordScene,omitnil,omitempty" name:"RecordScene"`

	// 录制自定义语言，仅recordlayout=9的时候此参数有效
	//
	// Deprecated: RecordLang is deprecated.
	RecordLang *string `json:"RecordLang,omitnil,omitempty" name:"RecordLang"`

	// 板书截图生成类型。0 不生成板书；1 全量模式；2 单页去重模式
	WhiteBoardSnapshotMode *uint64 `json:"WhiteBoardSnapshotMode,omitnil,omitempty" name:"WhiteBoardSnapshotMode"`

	// 字幕转写功能开关。可以有以下取值：
	// 0 不开启字幕转写功能（默认值）
	// 1 自动转写模式：上课自动开启，下课自动停止
	SubtitlesTranscription *uint64 `json:"SubtitlesTranscription,omitnil,omitempty" name:"SubtitlesTranscription"`

	// 嘉宾Id列表。当圆桌会议模式（RoomType==3）时生效
	Guests []*string `json:"Guests,omitnil,omitempty" name:"Guests"`

	// 录制文件合并开关。0 关闭 1 开启 注：只有在一节课多次启用手动录制时，此功能才有效
	RecordMerge *uint64 `json:"RecordMerge,omitnil,omitempty" name:"RecordMerge"`
}

type ModifyRoomRequest struct {
	*tchttp.BaseRequest
	
	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 低代码互动课堂的SdkAppId
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 预定的房间开始时间，unix时间戳（秒）。直播开始后不允许修改。
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 预定的房间结束时间，unix时间戳（秒）。直播开始后不允许修改。
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 老师ID。直播开始后不允许修改。
	TeacherId *string `json:"TeacherId,omitnil,omitempty" name:"TeacherId"`

	// 房间名称。
	// 字符数不超过256
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分辨率。可以有如下取值：
	// 1 标清
	// 2 高清
	// 3 全高清
	// 直播开始后不允许修改。
	Resolution *uint64 `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// 设置房间/课堂同时最大可与老师进行连麦互动的人数，该参数支持正式上课/开播前调用修改房间修改。小班课取值范围[0,16]，大班课取值范围[0,1]，当取值为0时表示当前课堂/直播，不支持连麦互动。
	MaxMicNumber *uint64 `json:"MaxMicNumber,omitnil,omitempty" name:"MaxMicNumber"`

	// 进入房间时是否自动连麦。可以有以下取值：
	// 0 不自动连麦（默认值）
	// 1 自动连麦
	// 直播开始后不允许修改。
	AutoMic *uint64 `json:"AutoMic,omitnil,omitempty" name:"AutoMic"`

	// 高音质模式。可以有以下取值：
	// 0 不开启高音质（默认值）
	// 1 开启高音质
	// 直播开始后不允许修改。
	AudioQuality *uint64 `json:"AudioQuality,omitnil,omitempty" name:"AudioQuality"`

	// 房间子类型，可以有以下取值：
	// videodoc 文档+视频
	// video 纯视频
	// 直播开始后不允许修改。
	SubType *string `json:"SubType,omitnil,omitempty" name:"SubType"`

	// 禁止录制。可以有以下取值：
	// 0 不禁止录制（默认值）
	// 1 禁止录制
	// 直播开始后不允许修改。
	DisableRecord *uint64 `json:"DisableRecord,omitnil,omitempty" name:"DisableRecord"`

	// 助教Id列表。直播开始后不允许修改。
	Assistants []*string `json:"Assistants,omitnil,omitempty" name:"Assistants"`

	// 房间绑定的群组ID。直播开始后不允许修改。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 打开学生麦克风/摄像头的授权开关。直播开始后不允许修改。
	EnableDirectControl *uint64 `json:"EnableDirectControl,omitnil,omitempty" name:"EnableDirectControl"`

	// 开启专注模式。
	// 0 收看全部角色音视频(默认)
	// 1 只看老师和助教
	InteractionMode *uint64 `json:"InteractionMode,omitnil,omitempty" name:"InteractionMode"`

	// 横竖屏。0：横屏开播（默认值）; 1：竖屏开播，当前仅支持移动端的纯视频类型
	VideoOrientation *uint64 `json:"VideoOrientation,omitnil,omitempty" name:"VideoOrientation"`

	// 开启课后评分。 0：不开启(默认)  1：开启
	IsGradingRequiredPostClass *uint64 `json:"IsGradingRequiredPostClass,omitnil,omitempty" name:"IsGradingRequiredPostClass"`

	// 房间类型: 0 小班课（默认值）; 1 大班课; 2 1V1 （预留参数、暂未开放)
	// 注：大班课的布局(layout)只有三分屏
	RoomType *uint64 `json:"RoomType,omitnil,omitempty" name:"RoomType"`

	// 录制模板。仅可修改还未开始的房间。录制模板枚举值参考：https://cloud.tencent.com/document/product/1639/89744
	RecordLayout *uint64 `json:"RecordLayout,omitnil,omitempty" name:"RecordLayout"`

	// 拖堂时间：单位分钟，0为不限制(默认值), -1为不能拖堂，大于0为拖堂的时间，最大值120分钟
	EndDelayTime *int64 `json:"EndDelayTime,omitnil,omitempty" name:"EndDelayTime"`

	// 直播方式：0 常规模式（默认）1 回放直播模式（伪直播）。 目前支持从回放直播模式（伪直播）改为常规模式，不支持从常规模式改为回放直播模式（伪直播）
	LiveType *uint64 `json:"LiveType,omitnil,omitempty" name:"LiveType"`

	// 伪直播链接
	RecordLiveUrl *string `json:"RecordLiveUrl,omitnil,omitempty" name:"RecordLiveUrl"`

	// 是否自动开始上课：0 不自动上课（默认） 1 自动上课 live_type=1的时候有效
	EnableAutoStart *uint64 `json:"EnableAutoStart,omitnil,omitempty" name:"EnableAutoStart"`

	// 录制自定义场景，仅recordlayout=9的时候此参数有效,数据内容为用户自定义场景参数，数据格式为json键值对方式，其中键值对的value为string类型。
	RecordScene *string `json:"RecordScene,omitnil,omitempty" name:"RecordScene"`

	// 录制自定义语言，仅recordlayout=9的时候此参数有效
	RecordLang *string `json:"RecordLang,omitnil,omitempty" name:"RecordLang"`

	// 板书截图生成类型。0 不生成板书；1 全量模式；2 单页去重模式
	WhiteBoardSnapshotMode *uint64 `json:"WhiteBoardSnapshotMode,omitnil,omitempty" name:"WhiteBoardSnapshotMode"`

	// 字幕转写功能开关。可以有以下取值：
	// 0 不开启字幕转写功能（默认值）
	// 1 自动转写模式：上课自动开启，下课自动停止
	SubtitlesTranscription *uint64 `json:"SubtitlesTranscription,omitnil,omitempty" name:"SubtitlesTranscription"`

	// 嘉宾Id列表。当圆桌会议模式（RoomType==3）时生效
	Guests []*string `json:"Guests,omitnil,omitempty" name:"Guests"`

	// 录制文件合并开关。0 关闭 1 开启 注：只有在一节课多次启用手动录制时，此功能才有效
	RecordMerge *uint64 `json:"RecordMerge,omitnil,omitempty" name:"RecordMerge"`
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
	delete(f, "RecordScene")
	delete(f, "RecordLang")
	delete(f, "WhiteBoardSnapshotMode")
	delete(f, "SubtitlesTranscription")
	delete(f, "Guests")
	delete(f, "RecordMerge")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRoomRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRoomResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 待修改的用户名。对应注册用户下“Name“字段，本次修改是对此内容进行修改。
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// 待修改头像url
	Avatar *string `json:"Avatar,omitnil,omitempty" name:"Avatar"`
}

type ModifyUserProfileRequest struct {
	*tchttp.BaseRequest
	
	// 待修改用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 待修改的用户名。对应注册用户下“Name“字段，本次修改是对此内容进行修改。
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// 待修改头像url
	Avatar *string `json:"Avatar,omitnil,omitempty" name:"Avatar"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// TIMCustomElem（自定义消息）
	MsgType *string `json:"MsgType,omitnil,omitempty" name:"MsgType"`

	// 文本消息，当MsgType 为TIMTextElem（文本消息）必选。
	TextMsgContent *TextMsgContent `json:"TextMsgContent,omitnil,omitempty" name:"TextMsgContent"`

	// 表情消息，当MsgType 为TIMFaceElem（表情消息）必选。
	FaceMsgContent *FaceMsgContent `json:"FaceMsgContent,omitnil,omitempty" name:"FaceMsgContent"`

	// 图像消息，当MsgType为TIMImageElem（图像消息）必选。
	ImageMsgContent *ImageMsgContent `json:"ImageMsgContent,omitnil,omitempty" name:"ImageMsgContent"`

	// 自定义消息，TIMCustomElem（自定义消息）必选。
	CustomMsgContent *CustomMsgContent `json:"CustomMsgContent,omitnil,omitempty" name:"CustomMsgContent"`
}

type MutedAccountList struct {
	// 用户 ID
	MemberAccount *string `json:"MemberAccount,omitnil,omitempty" name:"MemberAccount"`

	// 禁言到的时间（使用 UTC 时间，即世界协调时间）
	MutedUntil *uint64 `json:"MutedUntil,omitnil,omitempty" name:"MutedUntil"`
}

type QuestionInfo struct {
	// 问题ID
	QuestionId *string `json:"QuestionId,omitnil,omitempty" name:"QuestionId"`

	// 问题内容
	QuestionContent *string `json:"QuestionContent,omitnil,omitempty" name:"QuestionContent"`

	// 倒计时答题设置的秒数（0 表示不计时）
	Duration *uint64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 正确答案（按照位表示是否选择，如0x1表示选择A，0x11表示选择AB）
	CorrectAnswer *int64 `json:"CorrectAnswer,omitnil,omitempty" name:"CorrectAnswer"`

	// 每个选项答题人数统计
	AnswerStats []*AnswerStat `json:"AnswerStats,omitnil,omitempty" name:"AnswerStats"`
}

// Predefined struct for user
type RegisterUserRequestParams struct {
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 对应用户昵称。对应修改用户信息下“nickname“字段，在修改用户信息时，可以对该字段进行调整，从而更改用户的昵称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 用户在客户系统的Id，需要在同一应用下唯一。入参为空时默认赋值为UserId。
	OriginId *string `json:"OriginId,omitnil,omitempty" name:"OriginId"`

	// 用户头像。
	Avatar *string `json:"Avatar,omitnil,omitempty" name:"Avatar"`
}

type RegisterUserRequest struct {
	*tchttp.BaseRequest
	
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 对应用户昵称。对应修改用户信息下“nickname“字段，在修改用户信息时，可以对该字段进行调整，从而更改用户的昵称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 用户在客户系统的Id，需要在同一应用下唯一。入参为空时默认赋值为UserId。
	OriginId *string `json:"OriginId,omitnil,omitempty" name:"OriginId"`

	// 用户头像。
	Avatar *string `json:"Avatar,omitnil,omitempty" name:"Avatar"`
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
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 登录/注册成功后返回登录态token。有效期7天。
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 字符数不超过256
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 预定的房间开始时间，unix时间戳。
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 预定的房间结束时间，unix时间戳。
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 头像区域，摄像头视频画面的分辨率。可以有如下取值：
	// 1 标清
	// 2 高清
	// 3 全高清
	Resolution *uint64 `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// 设置房间/课堂同时最大可与老师进行连麦互动的人数，该参数支持正式上课/开播前调用修改房间修改。小班课取值范围[0,16]，大班课取值范围[0,1]，当取值为0时表示当前课堂/直播，不支持连麦互动。
	MaxMicNumber *uint64 `json:"MaxMicNumber,omitnil,omitempty" name:"MaxMicNumber"`

	// 房间子类型，可以有以下取值： videodoc 文档+视频 video 纯视频
	SubType *string `json:"SubType,omitnil,omitempty" name:"SubType"`

	// 老师ID。通过[注册用户]接口获取的UserId。
	TeacherId *string `json:"TeacherId,omitnil,omitempty" name:"TeacherId"`

	// 进入课堂时是否自动连麦。可以有以下取值： 0 不自动连麦（需要手动申请上麦，默认值） 1 自动连麦
	AutoMic *uint64 `json:"AutoMic,omitnil,omitempty" name:"AutoMic"`

	// 释放音视频权限后是否自动取消连麦。可以有以下取值： 0 自动取消连麦（默认值） 1 保持连麦状态
	TurnOffMic *uint64 `json:"TurnOffMic,omitnil,omitempty" name:"TurnOffMic"`

	// 高音质模式。可以有以下取值： 0 不开启高音质（默认值） 1 开启高音质
	AudioQuality *uint64 `json:"AudioQuality,omitnil,omitempty" name:"AudioQuality"`

	// 上课后是否禁止自动录制。可以有以下取值： 0 不禁止录制（自动开启录制，默认值） 1 禁止录制 注：如果该配置取值为0，录制将从上课后开始，课堂结束后停止。
	DisableRecord *uint64 `json:"DisableRecord,omitnil,omitempty" name:"DisableRecord"`

	// 助教Id列表。通过[注册用户]接口获取的UserId。
	Assistants []*string `json:"Assistants,omitnil,omitempty" name:"Assistants"`

	// rtc人数。
	//
	// Deprecated: RTCAudienceNumber is deprecated.
	RTCAudienceNumber *uint64 `json:"RTCAudienceNumber,omitnil,omitempty" name:"RTCAudienceNumber"`

	// 观看类型。
	AudienceType *uint64 `json:"AudienceType,omitnil,omitempty" name:"AudienceType"`

	// 录制布局。
	RecordLayout *uint64 `json:"RecordLayout,omitnil,omitempty" name:"RecordLayout"`

	// 房间绑定的群组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 打开学生麦克风/摄像头的授权开关
	EnableDirectControl *uint64 `json:"EnableDirectControl,omitnil,omitempty" name:"EnableDirectControl"`

	// 开启专注模式。 0 收看全部角色音视频(默认) 1 只看老师和助教
	InteractionMode *int64 `json:"InteractionMode,omitnil,omitempty" name:"InteractionMode"`

	// 横竖屏。0：横屏开播（默认值）; 1：竖屏开播，当前仅支持移动端的纯视频类型
	VideoOrientation *int64 `json:"VideoOrientation,omitnil,omitempty" name:"VideoOrientation"`

	// 开启课后评分。 0：不开启(默认)  1：开启
	IsGradingRequiredPostClass *int64 `json:"IsGradingRequiredPostClass,omitnil,omitempty" name:"IsGradingRequiredPostClass"`

	// 课堂类型: 0 小班课（默认值）; 1 大班课; 2 1V1 (预留参数，暂未开放); 3 圆桌会议 注：大班课的布局(layout)只有三分屏
	RoomType *int64 `json:"RoomType,omitnil,omitempty" name:"RoomType"`

	// 拖堂时间：单位分钟，0为不限制(默认值), -1为不能拖堂，大于0为拖堂的时间，最大值120分钟
	EndDelayTime *int64 `json:"EndDelayTime,omitnil,omitempty" name:"EndDelayTime"`

	// 直播类型：0 常规（默认）1 伪直播 2 RTMP推流直播
	LiveType *uint64 `json:"LiveType,omitnil,omitempty" name:"LiveType"`

	// 伪直播回放链接
	RecordLiveUrl *string `json:"RecordLiveUrl,omitnil,omitempty" name:"RecordLiveUrl"`

	// 是否自动开始上课：0 不自动上课（默认） 1 自动上课 live_type=1或2的时候有效
	EnableAutoStart *uint64 `json:"EnableAutoStart,omitnil,omitempty" name:"EnableAutoStart"`

	// 录制文件背景图片，支持png、jpg、jpeg、bmp格式，暂不支持透明通道
	RecordBackground *string `json:"RecordBackground,omitnil,omitempty" name:"RecordBackground"`

	// 录制自定义场景。注意：仅recordlayout=9的时候此参数有效。需注意各类参数配置正确能够生效。不然会造成录制失败，失败后无法补救。数据内容为用户自定义场景参数，数据格式为json键值对方式，其中键值对的value为string类型。自定义场景参数的含义。如下：     scene：自定义js/css对应的场景值。如scene=recordScene，会加载 recordScene 场景对应的 js/css，这样就可以自定义录制页面的元素。     lng：录制页面对应的语种。如lng=en，则录制界面为en。（枚举值：en,zh，zh-TW，jp，ar，kr，vi）     customToken：录制页面中涉及客户自己的服务需要鉴权时进行配置。一般情况下，无需配置。
	RecordScene *string `json:"RecordScene,omitnil,omitempty" name:"RecordScene"`

	// 录制自定义语言，仅recordlayout=9的时候此参数有效
	//
	// Deprecated: RecordLang is deprecated.
	RecordLang *string `json:"RecordLang,omitnil,omitempty" name:"RecordLang"`

	// 录制类型 0 仅录制混流（默认） ;1 录制混流+单流，该模式下除混流录制基础上，分别录制老师、台上学生的音视频流，每路录制都会产生相应的录制费用 。示例：0
	RecordStream *uint64 `json:"RecordStream,omitnil,omitempty" name:"RecordStream"`

	// 板书截图生成类型。0 不生成板书（默认）；1 全量模式；2 单页去重模式
	WhiteBoardSnapshotMode *uint64 `json:"WhiteBoardSnapshotMode,omitnil,omitempty" name:"WhiteBoardSnapshotMode"`

	// 字幕转写功能开关。可以有以下取值：
	// 0 不开启字幕转写功能（默认值）
	// 1 自动转写模式：上课自动开启，下课自动停止
	SubtitlesTranscription *uint64 `json:"SubtitlesTranscription,omitnil,omitempty" name:"SubtitlesTranscription"`

	// 嘉宾Id列表。当圆桌会议模式（RoomType==3）时生效
	Guests []*string `json:"Guests,omitnil,omitempty" name:"Guests"`

	// 录制文件合并开关。0 关闭 1 开启 注：只有在一节课多次启用手动录制时，此功能才有效
	RecordMerge *uint64 `json:"RecordMerge,omitnil,omitempty" name:"RecordMerge"`
}

type RoomItem struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 房间ID
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 房间状态。0 未开始 ；1进行中  ；2 已结束；3已过期
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 开始时间
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 实际开始时间
	RealStartTime *uint64 `json:"RealStartTime,omitnil,omitempty" name:"RealStartTime"`

	// 实际结束时间
	RealEndTime *uint64 `json:"RealEndTime,omitnil,omitempty" name:"RealEndTime"`

	// 头像区域，摄像头视频画面的分辨率。可以有如下取值：
	// 1 标清
	// 2 高清
	// 3 全高清
	Resolution *uint64 `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// 最大允许连麦人数。已废弃，使用字段 MaxMicNumber
	MaxRTCMember *uint64 `json:"MaxRTCMember,omitnil,omitempty" name:"MaxRTCMember"`

	// 房间录制地址。已废弃，使用新字段 RecordUrl
	ReplayUrl *string `json:"ReplayUrl,omitnil,omitempty" name:"ReplayUrl"`

	// 录制地址（协议为https)。仅在房间结束后存在。
	RecordUrl *string `json:"RecordUrl,omitnil,omitempty" name:"RecordUrl"`

	// 课堂同时最大可与老师进行连麦互动的人数，该参数支持正式上课/开播前调用修改房间修改。小班课取值范围[0,16]，大班课取值范围[0,1]，当取值为0时表示当前课堂/直播，不支持连麦互动。
	MaxMicNumber *uint64 `json:"MaxMicNumber,omitnil,omitempty" name:"MaxMicNumber"`

	// 打开学生麦克风/摄像头的授权开关 
	EnableDirectControl *uint64 `json:"EnableDirectControl,omitnil,omitempty" name:"EnableDirectControl"`

	// 开启专注模式。 0 收看全部角色音视频(默认) 1 只看老师和助教
	InteractionMode *int64 `json:"InteractionMode,omitnil,omitempty" name:"InteractionMode"`

	// 横竖屏。0：横屏开播（默认值）; 1：竖屏开播，当前仅支持移动端的纯视频类型
	VideoOrientation *int64 `json:"VideoOrientation,omitnil,omitempty" name:"VideoOrientation"`

	// 开启课后评分。 0：不开启(默认)  1：开启
	IsGradingRequiredPostClass *int64 `json:"IsGradingRequiredPostClass,omitnil,omitempty" name:"IsGradingRequiredPostClass"`

	// 房间类型。0:小班课（默认值）；1:大班课；2:1V1（后续扩展）
	// 注：大班课的布局(layout)只有三分屏
	RoomType *int64 `json:"RoomType,omitnil,omitempty" name:"RoomType"`

	// 拖堂时间：单位分钟，0为不限制(默认值), -1为不能拖堂，大于0为拖堂的时间，最大值120分钟
	EndDelayTime *int64 `json:"EndDelayTime,omitnil,omitempty" name:"EndDelayTime"`

	// 直播类型：0 常规（默认）1 伪直播
	LiveType *uint64 `json:"LiveType,omitnil,omitempty" name:"LiveType"`

	// 伪直播回放链接	
	RecordLiveUrl *string `json:"RecordLiveUrl,omitnil,omitempty" name:"RecordLiveUrl"`

	// 是否自动开始上课：0 不自动上课（默认） 1 自动上课 live_type=1的时候有效	
	EnableAutoStart *uint64 `json:"EnableAutoStart,omitnil,omitempty" name:"EnableAutoStart"`

	// 录制文件背景图片，支持png、jpg、jpeg、bmp格式，暂不支持透明通道
	RecordBackground *string `json:"RecordBackground,omitnil,omitempty" name:"RecordBackground"`

	// 录制自定义场景，仅recordlayout=9的时候此参数有效,数据内容为用户自定义场景参数，数据格式为json键值对方式，其中键值对的value为string类型。
	RecordScene *string `json:"RecordScene,omitnil,omitempty" name:"RecordScene"`

	// 录制自定义语言，仅recordlayout=9的时候此参数有效
	RecordLang *string `json:"RecordLang,omitnil,omitempty" name:"RecordLang"`

	// 板书截图生成类型。0 不生成板书；1 全量模式；2 单页去重模式
	WhiteBoardSnapshotMode *uint64 `json:"WhiteBoardSnapshotMode,omitnil,omitempty" name:"WhiteBoardSnapshotMode"`

	// 字幕转写功能开关：0关闭，1开启，默认关闭
	SubtitlesTranscription *uint64 `json:"SubtitlesTranscription,omitnil,omitempty" name:"SubtitlesTranscription"`
}

type SceneItem struct {
	// 场景名称
	Scene *string `json:"Scene,omitnil,omitempty" name:"Scene"`

	// logo地址
	LogoUrl *string `json:"LogoUrl,omitnil,omitempty" name:"LogoUrl"`

	// 主页地址
	HomeUrl *string `json:"HomeUrl,omitnil,omitempty" name:"HomeUrl"`

	// 自定义的js
	JSUrl *string `json:"JSUrl,omitnil,omitempty" name:"JSUrl"`

	// 自定义的css
	CSSUrl *string `json:"CSSUrl,omitnil,omitempty" name:"CSSUrl"`
}

// Predefined struct for user
type SendRoomNormalMessageRequestParams struct {
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 管理员指定消息发送方账号（若需设置 FromAccount 信息，则该参数取值不能为空）
	FromAccount *string `json:"FromAccount,omitnil,omitempty" name:"FromAccount"`

	// 自定义消息
	MsgBody []*MsgBody `json:"MsgBody,omitnil,omitempty" name:"MsgBody"`

	// 消息自定义数据（云端保存，会发送到对端，程序卸载重装后还能拉取到）。
	CloudCustomData *string `json:"CloudCustomData,omitnil,omitempty" name:"CloudCustomData"`

	// 昵称，当FromAccount没有在房间中，需要填写NickName，当FromAccount在房间中，填写NickName无意义
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 消息的优先级，默认优先级 Normal。
	// 可以指定3种优先级，从高到低依次为 High、Normal 和 Low，区分大小写。
	Priority *string `json:"Priority,omitnil,omitempty" name:"Priority"`
}

type SendRoomNormalMessageRequest struct {
	*tchttp.BaseRequest
	
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 管理员指定消息发送方账号（若需设置 FromAccount 信息，则该参数取值不能为空）
	FromAccount *string `json:"FromAccount,omitnil,omitempty" name:"FromAccount"`

	// 自定义消息
	MsgBody []*MsgBody `json:"MsgBody,omitnil,omitempty" name:"MsgBody"`

	// 消息自定义数据（云端保存，会发送到对端，程序卸载重装后还能拉取到）。
	CloudCustomData *string `json:"CloudCustomData,omitnil,omitempty" name:"CloudCustomData"`

	// 昵称，当FromAccount没有在房间中，需要填写NickName，当FromAccount在房间中，填写NickName无意义
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 消息的优先级，默认优先级 Normal。
	// 可以指定3种优先级，从高到低依次为 High、Normal 和 Low，区分大小写。
	Priority *string `json:"Priority,omitnil,omitempty" name:"Priority"`
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
	delete(f, "Priority")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendRoomNormalMessageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendRoomNormalMessageResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 消息。
	MsgContent *string `json:"MsgContent,omitnil,omitempty" name:"MsgContent"`
}

type SendRoomNotificationMessageRequest struct {
	*tchttp.BaseRequest
	
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 消息。
	MsgContent *string `json:"MsgContent,omitnil,omitempty" name:"MsgContent"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	CustomContent []*AppCustomContent `json:"CustomContent,omitnil,omitempty" name:"CustomContent"`

	// 应用ID。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
}

type SetAppCustomContentRequest struct {
	*tchttp.BaseRequest
	
	// 自定义内容。
	CustomContent []*AppCustomContent `json:"CustomContent,omitnil,omitempty" name:"CustomContent"`

	// 应用ID。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type SetMarqueeRequestParams struct {
	// 学校ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 房间号
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	//  跑马灯类型：1为固定值，2为用户昵称，3为固定值+用户昵称，4为用户ID，5为originId+固定值，6为用户昵称（originId）
	MarqueeType *uint64 `json:"MarqueeType,omitnil,omitempty" name:"MarqueeType"`

	// 显示方式：1为滚动，2为闪烁
	DisplayMode *uint64 `json:"DisplayMode,omitnil,omitempty" name:"DisplayMode"`

	// 固定值内容（当MarqueeType=1或5，则展示固定值内容）
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 字体大小（数字，像素单位，范围：10到24）。
	FontSize *uint64 `json:"FontSize,omitnil,omitempty" name:"FontSize"`

	// 字体粗细：1为粗体，0为细体
	FontWeight *uint64 `json:"FontWeight,omitnil,omitempty" name:"FontWeight"`

	// 字体颜色（十六进制颜色值，例如：#00FF00（绿色））
	FontColor *string `json:"FontColor,omitnil,omitempty" name:"FontColor"`

	// 字体透明度（数字，范围 0.0 到 1.0）
	FontOpacity *float64 `json:"FontOpacity,omitnil,omitempty" name:"FontOpacity"`

	// 背景颜色（十六进制颜色值，例如：#FFFF00（黄色））
	BackgroundColor *string `json:"BackgroundColor,omitnil,omitempty" name:"BackgroundColor"`

	// 背景透明度（数字，范围 0.0 到 1.0）
	BackgroundOpacity *float64 `json:"BackgroundOpacity,omitnil,omitempty" name:"BackgroundOpacity"`

	// 跑马灯文字移动/闪烁指定像素所需时间，范围：1-10；数值越小，跑马灯滚动/闪烁速度越快
	Duration *uint64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 跑马灯个数：目前仅支持1或2, 对应显示单排或双排
	MarqueeCount *uint64 `json:"MarqueeCount,omitnil,omitempty" name:"MarqueeCount"`
}

type SetMarqueeRequest struct {
	*tchttp.BaseRequest
	
	// 学校ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 房间号
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	//  跑马灯类型：1为固定值，2为用户昵称，3为固定值+用户昵称，4为用户ID，5为originId+固定值，6为用户昵称（originId）
	MarqueeType *uint64 `json:"MarqueeType,omitnil,omitempty" name:"MarqueeType"`

	// 显示方式：1为滚动，2为闪烁
	DisplayMode *uint64 `json:"DisplayMode,omitnil,omitempty" name:"DisplayMode"`

	// 固定值内容（当MarqueeType=1或5，则展示固定值内容）
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 字体大小（数字，像素单位，范围：10到24）。
	FontSize *uint64 `json:"FontSize,omitnil,omitempty" name:"FontSize"`

	// 字体粗细：1为粗体，0为细体
	FontWeight *uint64 `json:"FontWeight,omitnil,omitempty" name:"FontWeight"`

	// 字体颜色（十六进制颜色值，例如：#00FF00（绿色））
	FontColor *string `json:"FontColor,omitnil,omitempty" name:"FontColor"`

	// 字体透明度（数字，范围 0.0 到 1.0）
	FontOpacity *float64 `json:"FontOpacity,omitnil,omitempty" name:"FontOpacity"`

	// 背景颜色（十六进制颜色值，例如：#FFFF00（黄色））
	BackgroundColor *string `json:"BackgroundColor,omitnil,omitempty" name:"BackgroundColor"`

	// 背景透明度（数字，范围 0.0 到 1.0）
	BackgroundOpacity *float64 `json:"BackgroundOpacity,omitnil,omitempty" name:"BackgroundOpacity"`

	// 跑马灯文字移动/闪烁指定像素所需时间，范围：1-10；数值越小，跑马灯滚动/闪烁速度越快
	Duration *uint64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 跑马灯个数：目前仅支持1或2, 对应显示单排或双排
	MarqueeCount *uint64 `json:"MarqueeCount,omitnil,omitempty" name:"MarqueeCount"`
}

func (r *SetMarqueeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetMarqueeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	delete(f, "MarqueeType")
	delete(f, "DisplayMode")
	delete(f, "Content")
	delete(f, "FontSize")
	delete(f, "FontWeight")
	delete(f, "FontColor")
	delete(f, "FontOpacity")
	delete(f, "BackgroundColor")
	delete(f, "BackgroundOpacity")
	delete(f, "Duration")
	delete(f, "MarqueeCount")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetMarqueeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetMarqueeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetMarqueeResponse struct {
	*tchttp.BaseResponse
	Response *SetMarqueeResponseParams `json:"Response"`
}

func (r *SetMarqueeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetMarqueeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetWatermarkRequestParams struct {
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 老师视频区域的水印参数地址，设置为空字符串表示删除
	TeacherUrl *string `json:"TeacherUrl,omitnil,omitempty" name:"TeacherUrl"`

	// 白板视频区域的水印参数地址，设置为空字符串表示删除
	BoardUrl *string `json:"BoardUrl,omitnil,omitempty" name:"BoardUrl"`

	// 视频默认图片（在没有视频流的时候显示），设置为空字符串表示删除
	VideoUrl *string `json:"VideoUrl,omitnil,omitempty" name:"VideoUrl"`

	// 白板区域水印的宽度，取值:0-100，默认为0，表示区域X方向的百分比
	BoardW *float64 `json:"BoardW,omitnil,omitempty" name:"BoardW"`

	// 白板区域水印的高度，取值:0-100，默认为0, 表示区域Y方向的百分比
	BoardH *float64 `json:"BoardH,omitnil,omitempty" name:"BoardH"`

	// 白板区域水印X偏移, 取值:0-100, 表示区域X方向的百分比。比如50，则表示位于X轴中间
	BoardX *float64 `json:"BoardX,omitnil,omitempty" name:"BoardX"`

	// 白板区域水印Y偏移, 取值:0-100, 表示区域X方向的百分比。比如50，则表示位于X轴中间
	BoardY *float64 `json:"BoardY,omitnil,omitempty" name:"BoardY"`

	// 老师视频区域水印的宽度，取值:0-100，默认为0，表示区域X方向的百分比
	TeacherW *float64 `json:"TeacherW,omitnil,omitempty" name:"TeacherW"`

	// 老师视频区域水印的高度，取值:0-100，默认为0, 表示区域Y方向的百分比
	TeacherH *float64 `json:"TeacherH,omitnil,omitempty" name:"TeacherH"`

	// 老师视频区域水印X偏移, 取值:0-100, 表示区域X方向的百分比。比如50，则表示位于X轴中间
	TeacherX *float64 `json:"TeacherX,omitnil,omitempty" name:"TeacherX"`

	// 老师视频区域水印Y偏移, 取值:0-100, 表示区域X方向的百分比。比如50，则表示位于X轴中间
	TeacherY *float64 `json:"TeacherY,omitnil,omitempty" name:"TeacherY"`

	// 文字水印内容，设置为空字符串表示删除
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 文字水印颜色
	TextColor *string `json:"TextColor,omitnil,omitempty" name:"TextColor"`
}

type SetWatermarkRequest struct {
	*tchttp.BaseRequest
	
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 老师视频区域的水印参数地址，设置为空字符串表示删除
	TeacherUrl *string `json:"TeacherUrl,omitnil,omitempty" name:"TeacherUrl"`

	// 白板视频区域的水印参数地址，设置为空字符串表示删除
	BoardUrl *string `json:"BoardUrl,omitnil,omitempty" name:"BoardUrl"`

	// 视频默认图片（在没有视频流的时候显示），设置为空字符串表示删除
	VideoUrl *string `json:"VideoUrl,omitnil,omitempty" name:"VideoUrl"`

	// 白板区域水印的宽度，取值:0-100，默认为0，表示区域X方向的百分比
	BoardW *float64 `json:"BoardW,omitnil,omitempty" name:"BoardW"`

	// 白板区域水印的高度，取值:0-100，默认为0, 表示区域Y方向的百分比
	BoardH *float64 `json:"BoardH,omitnil,omitempty" name:"BoardH"`

	// 白板区域水印X偏移, 取值:0-100, 表示区域X方向的百分比。比如50，则表示位于X轴中间
	BoardX *float64 `json:"BoardX,omitnil,omitempty" name:"BoardX"`

	// 白板区域水印Y偏移, 取值:0-100, 表示区域X方向的百分比。比如50，则表示位于X轴中间
	BoardY *float64 `json:"BoardY,omitnil,omitempty" name:"BoardY"`

	// 老师视频区域水印的宽度，取值:0-100，默认为0，表示区域X方向的百分比
	TeacherW *float64 `json:"TeacherW,omitnil,omitempty" name:"TeacherW"`

	// 老师视频区域水印的高度，取值:0-100，默认为0, 表示区域Y方向的百分比
	TeacherH *float64 `json:"TeacherH,omitnil,omitempty" name:"TeacherH"`

	// 老师视频区域水印X偏移, 取值:0-100, 表示区域X方向的百分比。比如50，则表示位于X轴中间
	TeacherX *float64 `json:"TeacherX,omitnil,omitempty" name:"TeacherX"`

	// 老师视频区域水印Y偏移, 取值:0-100, 表示区域X方向的百分比。比如50，则表示位于X轴中间
	TeacherY *float64 `json:"TeacherY,omitnil,omitempty" name:"TeacherY"`

	// 文字水印内容，设置为空字符串表示删除
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 文字水印颜色
	TextColor *string `json:"TextColor,omitnil,omitempty" name:"TextColor"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type SingleStreamInfo struct {
	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 开始时间
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	StopTime *uint64 `json:"StopTime,omitnil,omitempty" name:"StopTime"`

	// 总时长
	Duration *uint64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 文件格式
	FileFormat *string `json:"FileFormat,omitnil,omitempty" name:"FileFormat"`

	// 流url
	RecordUrl *string `json:"RecordUrl,omitnil,omitempty" name:"RecordUrl"`

	// 流大小
	RecordSize *uint64 `json:"RecordSize,omitnil,omitempty" name:"RecordSize"`

	// 流ID
	VideoId *string `json:"VideoId,omitnil,omitempty" name:"VideoId"`

	// 流类型
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`
}

// Predefined struct for user
type StartRecordRequestParams struct {
	// 学校ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 房间ID
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

type StartRecordRequest struct {
	*tchttp.BaseRequest
	
	// 学校ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 房间ID
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

func (r *StartRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartRecordResponseParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartRecordResponse struct {
	*tchttp.BaseResponse
	Response *StartRecordResponseParams `json:"Response"`
}

func (r *StartRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartRoomRequestParams struct {
	// 课堂ID
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

type StartRoomRequest struct {
	*tchttp.BaseRequest
	
	// 课堂ID
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type StopRecordRequestParams struct {
	// 学校ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 课堂ID
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type StopRecordRequest struct {
	*tchttp.BaseRequest
	
	// 学校ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 课堂ID
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *StopRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopRecordResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopRecordResponse struct {
	*tchttp.BaseResponse
	Response *StopRecordResponseParams `json:"Response"`
}

func (r *StopRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TextMarkConfig struct {
	// 文字水印内容
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 文字水印颜色
	Color *string `json:"Color,omitnil,omitempty" name:"Color"`
}

type TextMsgContent struct {
	// 文本消息。
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type TransferItem struct {
	// 转存状态， 1正常 2停用
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *uint64 `json:"State,omitnil,omitempty" name:"State"`
}

// Predefined struct for user
type UnbindDocumentFromRoomRequestParams struct {
	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 文档ID。
	DocumentId *string `json:"DocumentId,omitnil,omitempty" name:"DocumentId"`
}

type UnbindDocumentFromRoomRequest struct {
	*tchttp.BaseRequest
	
	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 文档ID。
	DocumentId *string `json:"DocumentId,omitnil,omitempty" name:"DocumentId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type UnblockKickedUserRequestParams struct {
	// 低代码平台的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 课堂Id。
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 需要解禁踢出的成员Id。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type UnblockKickedUserRequest struct {
	*tchttp.BaseRequest
	
	// 低代码平台的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 课堂Id。
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 需要解禁踢出的成员Id。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *UnblockKickedUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnblockKickedUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnblockKickedUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnblockKickedUserResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UnblockKickedUserResponse struct {
	*tchttp.BaseResponse
	Response *UnblockKickedUserResponseParams `json:"Response"`
}

func (r *UnblockKickedUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnblockKickedUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserInfo struct {
	// 应用Id。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 用户Id。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户昵称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 用户头像Url。
	Avatar *string `json:"Avatar,omitnil,omitempty" name:"Avatar"`

	// 用户在客户系统的Id
	OriginId *string `json:"OriginId,omitnil,omitempty" name:"OriginId"`
}

type WatermarkConfig struct {
	// 水印图片的url
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 水印宽。为比例值
	Width *float64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 水印高。为比例值
	Height *float64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 水印X偏移, 取值:0-100, 表示区域X方向的百分比。比如50，则表示位于X轴中间。
	LocationX *float64 `json:"LocationX,omitnil,omitempty" name:"LocationX"`

	// 水印Y偏移, 取值:0-100, 表示区域Y方向的百分比。比如50，则表示位于Y轴中间。
	LocationY *float64 `json:"LocationY,omitnil,omitempty" name:"LocationY"`
}