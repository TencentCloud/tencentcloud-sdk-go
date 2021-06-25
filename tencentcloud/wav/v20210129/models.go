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

package v20210129

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type ActivityDetail struct {

	// 活动id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActivityId *int64 `json:"ActivityId,omitempty" name:"ActivityId"`

	// 活动名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActivityName *string `json:"ActivityName,omitempty" name:"ActivityName"`

	// 活动状态，10:未开始状态、20:已开始（进行中）状态、30:已结束状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActivityState *int64 `json:"ActivityState,omitempty" name:"ActivityState"`

	// 活动类型，100:留资活动
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActivityType *int64 `json:"ActivityType,omitempty" name:"ActivityType"`

	// 活动开始时间戳，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 活动结束时间戳，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 活动主图
	// 注意：此字段可能返回 null，表示取不到有效值。
	MainPhoto *string `json:"MainPhoto,omitempty" name:"MainPhoto"`

	// 协议编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivacyAgreementId *string `json:"PrivacyAgreementId,omitempty" name:"PrivacyAgreementId"`

	// 活动更新时间戳，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 活动数据列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActivityDataList *string `json:"ActivityDataList,omitempty" name:"ActivityDataList"`
}

type ActivityJoinDetail struct {

	// 活动id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActivityId *int64 `json:"ActivityId,omitempty" name:"ActivityId"`

	// 活动名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActivityName *string `json:"ActivityName,omitempty" name:"ActivityName"`

	// 销售姓名
	// 注意：此字段可能返回 null，表示取不到有效值。
	SalesName *string `json:"SalesName,omitempty" name:"SalesName"`

	// 销售电话
	// 注意：此字段可能返回 null，表示取不到有效值。
	SalesPhone *string `json:"SalesPhone,omitempty" name:"SalesPhone"`

	// 参与id
	// 注意：此字段可能返回 null，表示取不到有效值。
	JoinId *int64 `json:"JoinId,omitempty" name:"JoinId"`

	// 活码id
	// 注意：此字段可能返回 null，表示取不到有效值。
	LiveCodeId *int64 `json:"LiveCodeId,omitempty" name:"LiveCodeId"`

	// 用户电话
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserPhone *string `json:"UserPhone,omitempty" name:"UserPhone"`

	// 用户姓名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 活动数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActivityData *string `json:"ActivityData,omitempty" name:"ActivityData"`

	// 线索id
	// 注意：此字段可能返回 null，表示取不到有效值。
	LeadId *int64 `json:"LeadId,omitempty" name:"LeadId"`

	// 参与时间戳，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	JoinTime *int64 `json:"JoinTime,omitempty" name:"JoinTime"`

	// 线索是否是重复创建， 0 ：新建、 1：合并、 2：重复， 默认为0
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duplicate *int64 `json:"Duplicate,omitempty" name:"Duplicate"`

	// 重复线索id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DuplicateLeadId *int64 `json:"DuplicateLeadId,omitempty" name:"DuplicateLeadId"`

	// 是否为参与多次活动， 1：参与一次、2、参与多次，默认为0
	// 注意：此字段可能返回 null，表示取不到有效值。
	JoinState *int64 `json:"JoinState,omitempty" name:"JoinState"`

	// 创建时间戳，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间戳，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type ChannelCodeInnerDetail struct {

	// 渠道活码id
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 欢迎语类型，0：普通欢迎语、1:渠道欢迎语
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 渠道来源
	Source *string `json:"Source,omitempty" name:"Source"`

	// 渠道来源名称
	SourceName *string `json:"SourceName,omitempty" name:"SourceName"`

	// 二维码名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 使用成员用户id集
	UseUserIdList []*int64 `json:"UseUserIdList,omitempty" name:"UseUserIdList"`

	// 使用成员企微账号id集
	UseUserOpenIdList []*string `json:"UseUserOpenIdList,omitempty" name:"UseUserOpenIdList"`

	// 标签
	TagList []*WeComTagDetail `json:"TagList,omitempty" name:"TagList"`

	// 自动通过好友，0：开启、1：关闭，默认0开启
	SkipVerify *int64 `json:"SkipVerify,omitempty" name:"SkipVerify"`

	// 添加好友人数
	Friends *int64 `json:"Friends,omitempty" name:"Friends"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 欢迎语id（通过欢迎语新增返回的id）
	MsgId *int64 `json:"MsgId,omitempty" name:"MsgId"`

	// 联系我config_id
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 联系我二维码地址
	QrCodeUrl *string `json:"QrCodeUrl,omitempty" name:"QrCodeUrl"`

	// 记录状态， 0：有效、1：无效
	RecStatus *int64 `json:"RecStatus,omitempty" name:"RecStatus"`

	// 应用ID
	AppId *string `json:"AppId,omitempty" name:"AppId"`
}

type ChatArchivingDetail struct {

	// 消息id
	MsgId *string `json:"MsgId,omitempty" name:"MsgId"`

	// 动作名称，switch表示切换企微账号，send表示企微普通消息
	Action *string `json:"Action,omitempty" name:"Action"`

	// 消息类型，当Action != "switch"时存在，例如video, text, voice 等，和企微开放文档一一对应
	// https://open.work.weixin.qq.com/api/doc/90000/90135/91774
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgType *string `json:"MsgType,omitempty" name:"MsgType"`

	// 消息发送人
	// 注意：此字段可能返回 null，表示取不到有效值。
	From *string `json:"From,omitempty" name:"From"`

	// 消息接收人列表，注意接收人可能只有一个
	// 注意：此字段可能返回 null，表示取不到有效值。
	ToList []*string `json:"ToList,omitempty" name:"ToList"`

	// 如果是群消息，则不为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// 消息发送的时间戳，单位为秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgTime *uint64 `json:"MsgTime,omitempty" name:"MsgTime"`

	// MsgType=video时的消息体，忽略此字段，见BodyJson字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	Video *ChatArchivingMsgTypeVideo `json:"Video,omitempty" name:"Video"`

	// 根据MsgType的不同取值，解析内容不同，参考：https://open.work.weixin.qq.com/api/doc/90000/90135/91774
	// 注意：此字段可能返回 null，表示取不到有效值。
	BodyJson *string `json:"BodyJson,omitempty" name:"BodyJson"`
}

type ChatArchivingMsgTypeVideo struct {

	// 视频时长，单位秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlayLength *uint64 `json:"PlayLength,omitempty" name:"PlayLength"`

	// 文件大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileSize *uint64 `json:"FileSize,omitempty" name:"FileSize"`

	// 视频资源对象Cos下载地址
	CosKey *string `json:"CosKey,omitempty" name:"CosKey"`
}

type CreateChannelCodeRequest struct {
	*tchttp.BaseRequest

	// 欢迎语类型:0普通欢迎语,1渠道欢迎语
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 使用成员用户id集
	UseUserId []*int64 `json:"UseUserId,omitempty" name:"UseUserId"`

	// 使用成员企微账号id集
	UseUserOpenId []*string `json:"UseUserOpenId,omitempty" name:"UseUserOpenId"`

	// 应用ID,字典表中的APP_TYPE值,多个已逗号分隔
	AppIds *string `json:"AppIds,omitempty" name:"AppIds"`

	// 渠道来源
	Source *string `json:"Source,omitempty" name:"Source"`

	// 渠道来源名称
	SourceName *string `json:"SourceName,omitempty" name:"SourceName"`

	// 二维码名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 标签
	Tag []*WeComTagDetail `json:"Tag,omitempty" name:"Tag"`

	// 自动通过好友：0开启 1关闭, 默认开启
	SkipVerify *int64 `json:"SkipVerify,omitempty" name:"SkipVerify"`

	// 欢迎语id（通过欢迎语新增返回的id）
	MsgId *int64 `json:"MsgId,omitempty" name:"MsgId"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 渠道类型 0 未知 1 公域 2私域
	SourceType *int64 `json:"SourceType,omitempty" name:"SourceType"`
}

func (r *CreateChannelCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateChannelCodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "UseUserId")
	delete(f, "UseUserOpenId")
	delete(f, "AppIds")
	delete(f, "Source")
	delete(f, "SourceName")
	delete(f, "Name")
	delete(f, "Tag")
	delete(f, "SkipVerify")
	delete(f, "MsgId")
	delete(f, "Remark")
	delete(f, "SourceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateChannelCodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateChannelCodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateChannelCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateChannelCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCorpTagRequest struct {
	*tchttp.BaseRequest

	// 标签组名称，最长为15个字符
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 标签信息数组
	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`

	// 标签组次序值。sort值大的排序靠前。有效的值范围是[0, 2^32)
	Sort *uint64 `json:"Sort,omitempty" name:"Sort"`
}

func (r *CreateCorpTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCorpTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupName")
	delete(f, "Tags")
	delete(f, "Sort")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCorpTagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateCorpTagResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 标签组信息
		TagGroup *TagGroup `json:"TagGroup,omitempty" name:"TagGroup"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCorpTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCorpTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExternalContact struct {

	// 外部联系人的userId
	ExternalUserId *string `json:"ExternalUserId,omitempty" name:"ExternalUserId"`

	// 外部联系人性别 0-未知 1-男性 2-女性
	Gender *int64 `json:"Gender,omitempty" name:"Gender"`

	// 外部联系人的名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 外部联系人的类型，1表示该外部联系人是微信用户，2表示该外部联系人是企业微信用户
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 外部联系人在微信开放平台的唯一身份标识（微信unionid），通过此字段企业可将外部联系人与公众号/小程序用户关联起来。仅当联系人类型是微信用户，且企业或第三方服务商绑定了微信开发者ID有此字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnionId *string `json:"UnionId,omitempty" name:"UnionId"`

	// 外部联系人联系电话
	// 注意：此字段可能返回 null，表示取不到有效值。
	Phone *string `json:"Phone,omitempty" name:"Phone"`
}

type ExternalContactSimpleInfo struct {

	// 外部联系人的userId
	ExternalUserId *string `json:"ExternalUserId,omitempty" name:"ExternalUserId"`

	// 添加了此外部联系人的企业成员userId
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 添加了此外部联系人的企业成员的姓名
	// 注意：此字段可能返回 null，表示取不到有效值。
	SalesName *string `json:"SalesName,omitempty" name:"SalesName"`
}

type ExternalContactTag struct {

	// 该成员添加此外部联系人所打标签的分组名称（标签功能需要企业微信升级到2.7.5及以上版本）
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 该成员添加此外部联系人所打标签名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagName *string `json:"TagName,omitempty" name:"TagName"`

	// 该成员添加此外部联系人所打标签类型, 1-企业设置, 2-用户自定义
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 该成员添加此外部联系人所打企业标签的id，仅企业设置（type为1）的标签返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagId *string `json:"TagId,omitempty" name:"TagId"`
}

type FollowUser struct {

	// 添加了此外部联系人的企业成员userid
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 该成员对此外部联系人的备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 该成员对此外部联系人的描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 该成员添加此外部联系人的时间戳，单位为秒
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 该成员添加此客户的来源，具体含义详见<a href="https://work.weixin.qq.com/api/doc/90000/90135/92114#%E6%9D%A5%E6%BA%90%E5%AE%9A%E4%B9%89">来源定义</a>
	AddWay *int64 `json:"AddWay,omitempty" name:"AddWay"`

	// 发起添加的userid，如果成员主动添加，为成员的userid；如果是客户主动添加，则为客户的外部联系人userid；如果是内部成员共享/管理员分配，则为对应的成员/管理员userid
	OperUserId *string `json:"OperUserId,omitempty" name:"OperUserId"`

	// 该成员添加此外部联系人所打标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*ExternalContactTag `json:"Tags,omitempty" name:"Tags"`
}

type LiveCodeDetail struct {

	// 活码id
	LiveCodeId *uint64 `json:"LiveCodeId,omitempty" name:"LiveCodeId"`

	// 活码名称
	LiveCodeName *string `json:"LiveCodeName,omitempty" name:"LiveCodeName"`

	// 短链url
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShortChainAddress *string `json:"ShortChainAddress,omitempty" name:"ShortChainAddress"`

	// 活码二维码
	// 注意：此字段可能返回 null，表示取不到有效值。
	LiveCodePreview *string `json:"LiveCodePreview,omitempty" name:"LiveCodePreview"`

	// 活动id
	ActivityId *int64 `json:"ActivityId,omitempty" name:"ActivityId"`

	// 活动名称
	ActivityName *string `json:"ActivityName,omitempty" name:"ActivityName"`

	// 活码状态，-1：删除，0：启用，1禁用，默认为0
	LiveCodeState *int64 `json:"LiveCodeState,omitempty" name:"LiveCodeState"`

	// 活码参数，每个活码参数都是不一样的， 这个的值对应的是字符串json类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	LiveCodeData *string `json:"LiveCodeData,omitempty" name:"LiveCodeData"`

	// 创建时间戳，单位为秒
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间戳，单位为秒
	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type MiniAppCodeInfo struct {

	// 主键id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 小程序名称
	MiniAppName *string `json:"MiniAppName,omitempty" name:"MiniAppName"`

	// 小程序logo
	MiniAppLogo *string `json:"MiniAppLogo,omitempty" name:"MiniAppLogo"`

	// 小程序管理端地址
	MiniAdminUrl *string `json:"MiniAdminUrl,omitempty" name:"MiniAdminUrl"`

	// 状态：0正常，1删除
	State *int64 `json:"State,omitempty" name:"State"`

	// 创建时间戳，单位为秒
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间戳，单位为秒
	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type QueryActivityJoinListRequest struct {
	*tchttp.BaseRequest

	// 活动id
	ActivityId *int64 `json:"ActivityId,omitempty" name:"ActivityId"`

	// 分页游标，对应结果返回的NextCursor,首次请求保持为空
	Cursor *string `json:"Cursor,omitempty" name:"Cursor"`

	// 单页数据限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *QueryActivityJoinListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryActivityJoinListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ActivityId")
	delete(f, "Cursor")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryActivityJoinListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryActivityJoinListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分页游标
	// 注意：此字段可能返回 null，表示取不到有效值。
		NextCursor *string `json:"NextCursor,omitempty" name:"NextCursor"`

		// 活码列表响应参数
	// 注意：此字段可能返回 null，表示取不到有效值。
		PageData []*ActivityJoinDetail `json:"PageData,omitempty" name:"PageData"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryActivityJoinListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryActivityJoinListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryActivityListRequest struct {
	*tchttp.BaseRequest

	// 分页游标，对应结果返回的NextCursor,首次请求保持为空
	Cursor *string `json:"Cursor,omitempty" name:"Cursor"`

	// 单页数据限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *QueryActivityListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryActivityListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Cursor")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryActivityListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryActivityListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分页游标
	// 注意：此字段可能返回 null，表示取不到有效值。
		NextCursor *string `json:"NextCursor,omitempty" name:"NextCursor"`

		// 活码列表响应参数
	// 注意：此字段可能返回 null，表示取不到有效值。
		PageData []*ActivityDetail `json:"PageData,omitempty" name:"PageData"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryActivityListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryActivityListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryActivityLiveCodeListRequest struct {
	*tchttp.BaseRequest

	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor *string `json:"Cursor,omitempty" name:"Cursor"`

	// 返回的最大记录数，整型，最大值100，默认值50，超过最大值时取最大值
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *QueryActivityLiveCodeListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryActivityLiveCodeListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Cursor")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryActivityLiveCodeListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryActivityLiveCodeListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分页游标，再下次请求时填写以获取之后分页的记录，如果已经没有更多的数据则返回空
	// 注意：此字段可能返回 null，表示取不到有效值。
		NextCursor *string `json:"NextCursor,omitempty" name:"NextCursor"`

		// 活码列表响应参数
	// 注意：此字段可能返回 null，表示取不到有效值。
		PageData []*LiveCodeDetail `json:"PageData,omitempty" name:"PageData"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryActivityLiveCodeListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryActivityLiveCodeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryChannelCodeListRequest struct {
	*tchttp.BaseRequest

	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor *string `json:"Cursor,omitempty" name:"Cursor"`

	// 返回的最大记录数，整型，最大值100，默认值50，超过最大值时取最大值
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *QueryChannelCodeListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryChannelCodeListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Cursor")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryChannelCodeListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryChannelCodeListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分页游标，再下次请求时填写以获取之后分页的记录，如果已经没有更多的数据则返回空
	// 注意：此字段可能返回 null，表示取不到有效值。
		NextCursor *string `json:"NextCursor,omitempty" name:"NextCursor"`

		// 活码列表响应参数
	// 注意：此字段可能返回 null，表示取不到有效值。
		PageData []*ChannelCodeInnerDetail `json:"PageData,omitempty" name:"PageData"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryChannelCodeListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryChannelCodeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryChatArchivingListRequest struct {
	*tchttp.BaseRequest

	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor *string `json:"Cursor,omitempty" name:"Cursor"`

	// 返回的最大记录数，整型，最大值100，默认值50，超过最大值时取最大值
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *QueryChatArchivingListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryChatArchivingListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Cursor")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryChatArchivingListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryChatArchivingListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分页游标，再下次请求时填写以获取之后分页的记录，如果已经没有更多的数据则返回空
	// 注意：此字段可能返回 null，表示取不到有效值。
		NextCursor *string `json:"NextCursor,omitempty" name:"NextCursor"`

		// 会话存档列表响应参数
	// 注意：此字段可能返回 null，表示取不到有效值。
		PageData []*ChatArchivingDetail `json:"PageData,omitempty" name:"PageData"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryChatArchivingListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryChatArchivingListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryExternalContactDetailRequest struct {
	*tchttp.BaseRequest

	// 外部联系人的userid，注意不是企业成员的帐号
	ExternalUserId *string `json:"ExternalUserId,omitempty" name:"ExternalUserId"`

	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填。当客户在企业内的跟进人超过500人时需要使用cursor参数进行分页获取
	Cursor *string `json:"Cursor,omitempty" name:"Cursor"`

	// 当前接口Limit不需要传参， 保留Limit只是为了保持向后兼容性， Limit默认值为500，当返回结果超过500时， NextCursor才有返回值
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *QueryExternalContactDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryExternalContactDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ExternalUserId")
	delete(f, "Cursor")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryExternalContactDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryExternalContactDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分页游标，再下次请求时填写以获取之后分页的记录，如果已经没有更多的数据则返回空
	// 注意：此字段可能返回 null，表示取不到有效值。
		NextCursor *string `json:"NextCursor,omitempty" name:"NextCursor"`

		// 客户信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Customer *ExternalContact `json:"Customer,omitempty" name:"Customer"`

		// 添加了此外部联系人的企业成员信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		FollowUser []*FollowUser `json:"FollowUser,omitempty" name:"FollowUser"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryExternalContactDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryExternalContactDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryExternalContactListRequest struct {
	*tchttp.BaseRequest

	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor *string `json:"Cursor,omitempty" name:"Cursor"`

	// 返回的最大记录数，整型，最大值100，默认值50，超过最大值时取最大值
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *QueryExternalContactListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryExternalContactListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Cursor")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryExternalContactListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryExternalContactListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 外部联系人信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		PageData []*ExternalContactSimpleInfo `json:"PageData,omitempty" name:"PageData"`

		// 分页游标，再下次请求时填写以获取之后分页的记录，如果已经没有更多的数据则返回空
	// 注意：此字段可能返回 null，表示取不到有效值。
		NextCursor *string `json:"NextCursor,omitempty" name:"NextCursor"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryExternalContactListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryExternalContactListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryMiniAppCodeListRequest struct {
	*tchttp.BaseRequest

	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor *string `json:"Cursor,omitempty" name:"Cursor"`

	// 返回的最大记录数，整型，最大值100，默认值50，超过最大值时取最大值
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *QueryMiniAppCodeListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryMiniAppCodeListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Cursor")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryMiniAppCodeListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryMiniAppCodeListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分页游标，再下次请求时填写以获取之后分页的记录，如果已经没有更多的数据则返回空
	// 注意：此字段可能返回 null，表示取不到有效值。
		NextCursor *string `json:"NextCursor,omitempty" name:"NextCursor"`

		// 小程序码列表响应数据
	// 注意：此字段可能返回 null，表示取不到有效值。
		PageData []*MiniAppCodeInfo `json:"PageData,omitempty" name:"PageData"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryMiniAppCodeListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryMiniAppCodeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TagDetailInfo struct {

	// 标签名称
	TagName *string `json:"TagName,omitempty" name:"TagName"`

	// 标签业务ID
	BizTagId *string `json:"BizTagId,omitempty" name:"BizTagId"`

	// 企微标签ID
	TagId *string `json:"TagId,omitempty" name:"TagId"`

	// 标签排序的次序值，sort值大的排序靠前。有效的值范围是[0, 2^32)
	Sort *uint64 `json:"Sort,omitempty" name:"Sort"`

	// 标签创建时间,单位为秒
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
}

type TagGroup struct {

	// 企微标签组id
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 标签组业务id
	BizGroupId *string `json:"BizGroupId,omitempty" name:"BizGroupId"`

	// 企微标签组名称，不能超过15个字符
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 标签组次序值。sort值大的排序靠前。有效的值范围是[0, 2^32)
	Sort *uint64 `json:"Sort,omitempty" name:"Sort"`

	// 标签组创建时间,单位为秒
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 标签组内的标签列表, 上限为20
	Tags []*TagDetailInfo `json:"Tags,omitempty" name:"Tags"`
}

type TagInfo struct {

	// 标签名称, 最大长度限制15个字符
	TagName *string `json:"TagName,omitempty" name:"TagName"`

	// 标签组排序,值越大,排序越靠前
	Sort *uint64 `json:"Sort,omitempty" name:"Sort"`
}

type WeComTagDetail struct {

	// 标签分组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 标签分组业务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	BizGroupId *string `json:"BizGroupId,omitempty" name:"BizGroupId"`

	// 标签名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagName *string `json:"TagName,omitempty" name:"TagName"`

	// 标签ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagId *string `json:"TagId,omitempty" name:"TagId"`

	// 标签业务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	BizTagId *string `json:"BizTagId,omitempty" name:"BizTagId"`

	// 标签分类，1：企业设置、2：用户自定义
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 标签业务ID字符串格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	BizTagIdStr *string `json:"BizTagIdStr,omitempty" name:"BizTagIdStr"`
}
