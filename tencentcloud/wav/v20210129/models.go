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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type ActivityDetail struct {
	// 活动id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActivityId *int64 `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// 活动名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActivityName *string `json:"ActivityName,omitnil,omitempty" name:"ActivityName"`

	// 活动状态，10:未开始状态、20:已开始（进行中）状态、30:已结束状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActivityState *int64 `json:"ActivityState,omitnil,omitempty" name:"ActivityState"`

	// 活动类型，100:留资活动
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActivityType *int64 `json:"ActivityType,omitnil,omitempty" name:"ActivityType"`

	// 活动开始时间戳，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 活动结束时间戳，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 活动主图
	// 注意：此字段可能返回 null，表示取不到有效值。
	MainPhoto *string `json:"MainPhoto,omitnil,omitempty" name:"MainPhoto"`

	// 协议编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivacyAgreementId *string `json:"PrivacyAgreementId,omitnil,omitempty" name:"PrivacyAgreementId"`

	// 活动更新时间戳，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 活动数据列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActivityDataList *string `json:"ActivityDataList,omitnil,omitempty" name:"ActivityDataList"`
}

type ActivityJoinDetail struct {
	// 活动id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActivityId *int64 `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// 活动名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActivityName *string `json:"ActivityName,omitnil,omitempty" name:"ActivityName"`

	// 销售姓名
	// 注意：此字段可能返回 null，表示取不到有效值。
	SalesName *string `json:"SalesName,omitnil,omitempty" name:"SalesName"`

	// 销售电话
	// 注意：此字段可能返回 null，表示取不到有效值。
	SalesPhone *string `json:"SalesPhone,omitnil,omitempty" name:"SalesPhone"`

	// 参与id
	// 注意：此字段可能返回 null，表示取不到有效值。
	JoinId *int64 `json:"JoinId,omitnil,omitempty" name:"JoinId"`

	// 活码id
	// 注意：此字段可能返回 null，表示取不到有效值。
	LiveCodeId *int64 `json:"LiveCodeId,omitnil,omitempty" name:"LiveCodeId"`

	// 用户电话
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserPhone *string `json:"UserPhone,omitnil,omitempty" name:"UserPhone"`

	// 用户姓名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 活动数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActivityData *string `json:"ActivityData,omitnil,omitempty" name:"ActivityData"`

	// 线索id
	// 注意：此字段可能返回 null，表示取不到有效值。
	LeadId *int64 `json:"LeadId,omitnil,omitempty" name:"LeadId"`

	// 参与时间戳，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	JoinTime *int64 `json:"JoinTime,omitnil,omitempty" name:"JoinTime"`

	// 线索是否是重复创建， 0 ：新建、 1：合并、 2：重复， 默认为0
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duplicate *int64 `json:"Duplicate,omitnil,omitempty" name:"Duplicate"`

	// 重复线索id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DuplicateLeadId *int64 `json:"DuplicateLeadId,omitnil,omitempty" name:"DuplicateLeadId"`

	// 是否为参与多次活动， 1：参与一次、2、参与多次，默认为0
	// 注意：此字段可能返回 null，表示取不到有效值。
	JoinState *int64 `json:"JoinState,omitnil,omitempty" name:"JoinState"`

	// 创建时间戳，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间戳，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type ArrivalInfo struct {
	// 线索id
	ClueId *uint64 `json:"ClueId,omitnil,omitempty" name:"ClueId"`

	// 客户id
	CustomerId *uint64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 客户姓名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 客户的手机号
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 是否首次到店，0否，1是
	FirstArrival *int64 `json:"FirstArrival,omitnil,omitempty" name:"FirstArrival"`

	// 到店时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ArrivalTime *uint64 `json:"ArrivalTime,omitnil,omitempty" name:"ArrivalTime"`

	// 发生事件
	EventType *int64 `json:"EventType,omitnil,omitempty" name:"EventType"`

	// 发生事件名称
	EventTypeName *string `json:"EventTypeName,omitnil,omitempty" name:"EventTypeName"`

	// 跟进记录
	FollowRecord *string `json:"FollowRecord,omitnil,omitempty" name:"FollowRecord"`
}

type ChannelCodeInnerDetail struct {
	// 渠道活码id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 欢迎语类型，0：普通欢迎语、1:渠道欢迎语
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 渠道来源
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 渠道来源名称
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// 二维码名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 使用成员用户id集
	UseUserIdList []*int64 `json:"UseUserIdList,omitnil,omitempty" name:"UseUserIdList"`

	// 使用成员企微账号id集
	UseUserOpenIdList []*string `json:"UseUserOpenIdList,omitnil,omitempty" name:"UseUserOpenIdList"`

	// 标签
	TagList []*WeComTagDetail `json:"TagList,omitnil,omitempty" name:"TagList"`

	// 自动通过好友，0：开启、1：关闭，默认0开启
	SkipVerify *int64 `json:"SkipVerify,omitnil,omitempty" name:"SkipVerify"`

	// 添加好友人数
	Friends *int64 `json:"Friends,omitnil,omitempty" name:"Friends"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 欢迎语id（通过欢迎语新增返回的id）
	MsgId *int64 `json:"MsgId,omitnil,omitempty" name:"MsgId"`

	// 联系我config_id
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`

	// 联系我二维码地址
	QrCodeUrl *string `json:"QrCodeUrl,omitnil,omitempty" name:"QrCodeUrl"`

	// 记录状态， 0：有效、1：无效
	RecStatus *int64 `json:"RecStatus,omitnil,omitempty" name:"RecStatus"`

	// 应用ID
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`
}

type ChannelTag struct {
	// 该客户档案当前已成功关联的渠道标签名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagName *string `json:"TagName,omitnil,omitempty" name:"TagName"`

	// 该客户档案当前已成功关联的渠道标签的id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagId *string `json:"TagId,omitnil,omitempty" name:"TagId"`
}

type ChatArchivingDetail struct {
	// 消息id
	MsgId *string `json:"MsgId,omitnil,omitempty" name:"MsgId"`

	// 动作名称，switch表示切换企微账号，send表示企微普通消息
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 消息类型，当Action != "switch"时存在，例如video, text, voice 等，和企微开放文档一一对应
	// https://open.work.weixin.qq.com/api/doc/90000/90135/91774
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgType *string `json:"MsgType,omitnil,omitempty" name:"MsgType"`

	// 消息发送人
	// 注意：此字段可能返回 null，表示取不到有效值。
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 消息接收人列表，注意接收人可能只有一个
	// 注意：此字段可能返回 null，表示取不到有效值。
	ToList []*string `json:"ToList,omitnil,omitempty" name:"ToList"`

	// 如果是群消息，则不为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 消息发送的时间戳，单位为秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgTime *uint64 `json:"MsgTime,omitnil,omitempty" name:"MsgTime"`

	// MsgType=video时的消息体，忽略此字段，见BodyJson字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	Video *ChatArchivingMsgTypeVideo `json:"Video,omitnil,omitempty" name:"Video"`

	// 根据MsgType的不同取值，解析内容不同，参考：https://open.work.weixin.qq.com/api/doc/90000/90135/91774
	// 注意：此字段可能返回 null，表示取不到有效值。
	BodyJson *string `json:"BodyJson,omitnil,omitempty" name:"BodyJson"`
}

type ChatArchivingMsgTypeVideo struct {
	// 视频时长，单位秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlayLength *uint64 `json:"PlayLength,omitnil,omitempty" name:"PlayLength"`

	// 文件大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileSize *uint64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// 视频资源对象Cos下载地址
	CosKey *string `json:"CosKey,omitnil,omitempty" name:"CosKey"`
}

type ClueInfoDetail struct {
	// 线索id，线索唯一识别编码
	ClueId *string `json:"ClueId,omitnil,omitempty" name:"ClueId"`

	// 接待客户经销商顾问所属经销商code
	DealerId *string `json:"DealerId,omitnil,omitempty" name:"DealerId"`

	// 线索获取时间，用户添加企业微信时间，单位是秒
	EnquireTime *uint64 `json:"EnquireTime,omitnil,omitempty" name:"EnquireTime"`

	// 客户在微信生态中唯一识别码
	UnionId *string `json:"UnionId,omitnil,omitempty" name:"UnionId"`

	// 微信昵称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 联系方式
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 车系编号
	SeriesCode *string `json:"SeriesCode,omitnil,omitempty" name:"SeriesCode"`

	// 车型编号
	ModelCode *string `json:"ModelCode,omitnil,omitempty" name:"ModelCode"`

	// 省份编号
	ProvinceCode *string `json:"ProvinceCode,omitnil,omitempty" name:"ProvinceCode"`

	// 城市编号
	CityCode *string `json:"CityCode,omitnil,omitempty" name:"CityCode"`

	// 顾问名称
	SalesName *string `json:"SalesName,omitnil,omitempty" name:"SalesName"`

	// 顾问电话
	SalesPhone *string `json:"SalesPhone,omitnil,omitempty" name:"SalesPhone"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 标签
	TagList []*string `json:"TagList,omitnil,omitempty" name:"TagList"`

	// 客户姓名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 线索属性，0：个人，1：企业
	LeadUserType *int64 `json:"LeadUserType,omitnil,omitempty" name:"LeadUserType"`

	// 线索来源类型，1：线上，2：线下
	LeadType *int64 `json:"LeadType,omitnil,omitempty" name:"LeadType"`

	// 线索渠道对应ID
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// 线索渠道类型，与线索来源对应的渠道名称
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`

	// 线索渠道名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceChannelName *string `json:"SourceChannelName,omitnil,omitempty" name:"SourceChannelName"`

	// 0：未知，1：男，2：女
	Gender *int64 `json:"Gender,omitnil,omitempty" name:"Gender"`

	// 线索创建时间戳，单位：秒
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 线索创建时间戳，单位：秒
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 线索所处状态，101-待分配 201-待建档 301-已建档 401-已邀约 501-跟进中 601-已下订单 701-已成交 801-战败申请中 901-已战败 1001-未知状态 1101-转移申请中 1201-已完成
	LeadStatus *int64 `json:"LeadStatus,omitnil,omitempty" name:"LeadStatus"`

	// 线索意向等级
	LevelCode *string `json:"LevelCode,omitnil,omitempty" name:"LevelCode"`

	// 线索成功导入的时间戳，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImportAtTime *uint64 `json:"ImportAtTime,omitnil,omitempty" name:"ImportAtTime"`

	// 完成线索分配的时间戳，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	DistributeTime *uint64 `json:"DistributeTime,omitnil,omitempty" name:"DistributeTime"`

	// 获取线索的时间戳，单位：秒
	CreateAtTime *uint64 `json:"CreateAtTime,omitnil,omitempty" name:"CreateAtTime"`

	// 客户微信id
	WxId *string `json:"WxId,omitnil,omitempty" name:"WxId"`

	// 意向车型对应品牌code
	BrandCode *string `json:"BrandCode,omitnil,omitempty" name:"BrandCode"`

	// 建档时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	BuildTime *uint64 `json:"BuildTime,omitnil,omitempty" name:"BuildTime"`

	// 下订时间，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderTime *uint64 `json:"OrderTime,omitnil,omitempty" name:"OrderTime"`

	// 到店时间，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	ArrivalTime *uint64 `json:"ArrivalTime,omitnil,omitempty" name:"ArrivalTime"`

	// 交车时间，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeliveryTime *uint64 `json:"DeliveryTime,omitnil,omitempty" name:"DeliveryTime"`

	// 上次跟进时间，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	FollowTime *uint64 `json:"FollowTime,omitnil,omitempty" name:"FollowTime"`

	// 下次跟进时间，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextFollowTime *uint64 `json:"NextFollowTime,omitnil,omitempty" name:"NextFollowTime"`

	// 线索所属组织id
	OrgId *uint64 `json:"OrgId,omitnil,omitempty" name:"OrgId"`

	// 线索所属组织名称
	OrgName *string `json:"OrgName,omitnil,omitempty" name:"OrgName"`

	// 介绍人姓名
	Introducer *string `json:"Introducer,omitnil,omitempty" name:"Introducer"`

	// 介绍人电话
	IntroducerPhone *string `json:"IntroducerPhone,omitnil,omitempty" name:"IntroducerPhone"`

	// 是否关联微信 1 是 0 否
	IsBindWx *int64 `json:"IsBindWx,omitnil,omitempty" name:"IsBindWx"`

	// 是否经过合并 1 是 0 否
	IsMerge *int64 `json:"IsMerge,omitnil,omitempty" name:"IsMerge"`

	// 是否无效  1 是 0 否
	IsInvalid *int64 `json:"IsInvalid,omitnil,omitempty" name:"IsInvalid"`

	// 无效类型
	InvalidType *string `json:"InvalidType,omitnil,omitempty" name:"InvalidType"`

	// 无效类型枚举：
	// 无意向购买、空错号、未接听、其他
	InvalidTypeName *string `json:"InvalidTypeName,omitnil,omitempty" name:"InvalidTypeName"`

	// 由顾问手动输入的无效原因文字
	InvalidRemark *string `json:"InvalidRemark,omitnil,omitempty" name:"InvalidRemark"`

	// 无效时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvalidTime *uint64 `json:"InvalidTime,omitnil,omitempty" name:"InvalidTime"`

	// 经销商名称
	DealerName *string `json:"DealerName,omitnil,omitempty" name:"DealerName"`

	// 经销商下级门店ID
	ShopId *uint64 `json:"ShopId,omitnil,omitempty" name:"ShopId"`

	// 经销商下级门店名称
	ShopName *string `json:"ShopName,omitnil,omitempty" name:"ShopName"`

	// 职位
	Position *string `json:"Position,omitnil,omitempty" name:"Position"`

	// 自定义的门店id
	CorpShopId *string `json:"CorpShopId,omitnil,omitempty" name:"CorpShopId"`
}

type CorpUserInfo struct {
	// 企业成员UserId
	UserId *uint64 `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 企业成员在SaaS名片内填写的姓名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 企业成员在企微原生通讯录内的id
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserOpenId *string `json:"UserOpenId,omitnil,omitempty" name:"UserOpenId"`

	// 成员所属经销商id，可为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	DealerId *uint64 `json:"DealerId,omitnil,omitempty" name:"DealerId"`

	// 成员所属门店id，可为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShopId *uint64 `json:"ShopId,omitnil,omitempty" name:"ShopId"`

	// 企业成员手机号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 成员所属部门id列表，仅返回该应用有查看权限的部门id；成员授权模式下，固定返回根部门id，即固定为1；多个部门使用逗号分割
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrgIds *string `json:"OrgIds,omitnil,omitempty" name:"OrgIds"`

	// 主部门，仅当应用对主部门有查看权限时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	MainDepartment *string `json:"MainDepartment,omitnil,omitempty" name:"MainDepartment"`

	// 是否为部门负责人，第三方应用可为空。与orgIds值一一对应，多个部门使用逗号隔开，0-否， 1-是
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsLeaderInDept *string `json:"IsLeaderInDept,omitnil,omitempty" name:"IsLeaderInDept"`

	// 激活状态: 0=已激活，1=已禁用，-1=退出企业"
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 工号
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobNumber *string `json:"JobNumber,omitnil,omitempty" name:"JobNumber"`
}

// Predefined struct for user
type CreateChannelCodeRequestParams struct {
	// 欢迎语类型:0普通欢迎语,1渠道欢迎语
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 使用成员用户id集
	UseUserId []*int64 `json:"UseUserId,omitnil,omitempty" name:"UseUserId"`

	// 使用成员企微账号id集
	UseUserOpenId []*string `json:"UseUserOpenId,omitnil,omitempty" name:"UseUserOpenId"`

	// 应用ID,字典表中的APP_TYPE值,多个已逗号分隔
	AppIds *string `json:"AppIds,omitnil,omitempty" name:"AppIds"`

	// 渠道来源
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 渠道来源名称
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// 二维码名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 标签
	Tag []*WeComTagDetail `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 自动通过好友：0开启 1关闭, 默认开启
	SkipVerify *int64 `json:"SkipVerify,omitnil,omitempty" name:"SkipVerify"`

	// 欢迎语id（通过欢迎语新增返回的id）
	MsgId *int64 `json:"MsgId,omitnil,omitempty" name:"MsgId"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 渠道类型 0 未知 1 公域 2私域
	SourceType *int64 `json:"SourceType,omitnil,omitempty" name:"SourceType"`
}

type CreateChannelCodeRequest struct {
	*tchttp.BaseRequest
	
	// 欢迎语类型:0普通欢迎语,1渠道欢迎语
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 使用成员用户id集
	UseUserId []*int64 `json:"UseUserId,omitnil,omitempty" name:"UseUserId"`

	// 使用成员企微账号id集
	UseUserOpenId []*string `json:"UseUserOpenId,omitnil,omitempty" name:"UseUserOpenId"`

	// 应用ID,字典表中的APP_TYPE值,多个已逗号分隔
	AppIds *string `json:"AppIds,omitnil,omitempty" name:"AppIds"`

	// 渠道来源
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 渠道来源名称
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// 二维码名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 标签
	Tag []*WeComTagDetail `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 自动通过好友：0开启 1关闭, 默认开启
	SkipVerify *int64 `json:"SkipVerify,omitnil,omitempty" name:"SkipVerify"`

	// 欢迎语id（通过欢迎语新增返回的id）
	MsgId *int64 `json:"MsgId,omitnil,omitempty" name:"MsgId"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 渠道类型 0 未知 1 公域 2私域
	SourceType *int64 `json:"SourceType,omitnil,omitempty" name:"SourceType"`
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

// Predefined struct for user
type CreateChannelCodeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateChannelCodeResponse struct {
	*tchttp.BaseResponse
	Response *CreateChannelCodeResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateCorpTagRequestParams struct {
	// 标签组名称，最长为15个字符
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 标签信息数组
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 标签组次序值。sort值大的排序靠前。有效的值范围是[0, 2^32)
	Sort *uint64 `json:"Sort,omitnil,omitempty" name:"Sort"`
}

type CreateCorpTagRequest struct {
	*tchttp.BaseRequest
	
	// 标签组名称，最长为15个字符
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 标签信息数组
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 标签组次序值。sort值大的排序靠前。有效的值范围是[0, 2^32)
	Sort *uint64 `json:"Sort,omitnil,omitempty" name:"Sort"`
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

// Predefined struct for user
type CreateCorpTagResponseParams struct {
	// 标签组信息
	TagGroup *TagGroup `json:"TagGroup,omitnil,omitempty" name:"TagGroup"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCorpTagResponse struct {
	*tchttp.BaseResponse
	Response *CreateCorpTagResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateLeadRequestParams struct {
	// 来源ID
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// 来源名称
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`

	// 创建时间， 单位毫秒
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 线索类型：1-400呼入，2-常规留资
	SourceType *int64 `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 经销商id
	DealerId *uint64 `json:"DealerId,omitnil,omitempty" name:"DealerId"`

	// 品牌id
	BrandId *uint64 `json:"BrandId,omitnil,omitempty" name:"BrandId"`

	// 车系id
	SeriesId *uint64 `json:"SeriesId,omitnil,omitempty" name:"SeriesId"`

	// 客户姓名
	CustomerName *string `json:"CustomerName,omitnil,omitempty" name:"CustomerName"`

	// 客户手机号
	CustomerPhone *string `json:"CustomerPhone,omitnil,omitempty" name:"CustomerPhone"`

	// 车型id
	ModelId *uint64 `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 客户性别: 0-未知, 1-男, 2-女
	CustomerSex *int64 `json:"CustomerSex,omitnil,omitempty" name:"CustomerSex"`

	// 销售姓名
	SalesName *string `json:"SalesName,omitnil,omitempty" name:"SalesName"`

	// 销售手机号
	SalesPhone *string `json:"SalesPhone,omitnil,omitempty" name:"SalesPhone"`

	// Cc坐席姓名
	CcName *string `json:"CcName,omitnil,omitempty" name:"CcName"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type CreateLeadRequest struct {
	*tchttp.BaseRequest
	
	// 来源ID
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// 来源名称
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`

	// 创建时间， 单位毫秒
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 线索类型：1-400呼入，2-常规留资
	SourceType *int64 `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 经销商id
	DealerId *uint64 `json:"DealerId,omitnil,omitempty" name:"DealerId"`

	// 品牌id
	BrandId *uint64 `json:"BrandId,omitnil,omitempty" name:"BrandId"`

	// 车系id
	SeriesId *uint64 `json:"SeriesId,omitnil,omitempty" name:"SeriesId"`

	// 客户姓名
	CustomerName *string `json:"CustomerName,omitnil,omitempty" name:"CustomerName"`

	// 客户手机号
	CustomerPhone *string `json:"CustomerPhone,omitnil,omitempty" name:"CustomerPhone"`

	// 车型id
	ModelId *uint64 `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 客户性别: 0-未知, 1-男, 2-女
	CustomerSex *int64 `json:"CustomerSex,omitnil,omitempty" name:"CustomerSex"`

	// 销售姓名
	SalesName *string `json:"SalesName,omitnil,omitempty" name:"SalesName"`

	// 销售手机号
	SalesPhone *string `json:"SalesPhone,omitnil,omitempty" name:"SalesPhone"`

	// Cc坐席姓名
	CcName *string `json:"CcName,omitnil,omitempty" name:"CcName"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *CreateLeadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLeadRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelId")
	delete(f, "ChannelName")
	delete(f, "CreateTime")
	delete(f, "SourceType")
	delete(f, "DealerId")
	delete(f, "BrandId")
	delete(f, "SeriesId")
	delete(f, "CustomerName")
	delete(f, "CustomerPhone")
	delete(f, "ModelId")
	delete(f, "CustomerSex")
	delete(f, "SalesName")
	delete(f, "SalesPhone")
	delete(f, "CcName")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLeadRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLeadResponseParams struct {
	// 线索处理状态码： 0-表示创建成功， 1-表示线索合并，2-表示线索重复
	BusinessCode *int64 `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// 线索处理结果描述
	BusinessMsg *string `json:"BusinessMsg,omitnil,omitempty" name:"BusinessMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLeadResponse struct {
	*tchttp.BaseResponse
	Response *CreateLeadResponseParams `json:"Response"`
}

func (r *CreateLeadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLeadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CrmStatisticsData struct {
	// 新增线索
	LeadCnt *uint64 `json:"LeadCnt,omitnil,omitempty" name:"LeadCnt"`

	// 新增建档
	BuildCnt *uint64 `json:"BuildCnt,omitnil,omitempty" name:"BuildCnt"`

	// 新增到店
	InvitedCnt *uint64 `json:"InvitedCnt,omitnil,omitempty" name:"InvitedCnt"`

	// 新增下订
	OrderedCnt *uint64 `json:"OrderedCnt,omitnil,omitempty" name:"OrderedCnt"`

	// 新增成交
	DeliveredCnt *uint64 `json:"DeliveredCnt,omitnil,omitempty" name:"DeliveredCnt"`

	// 新增战败
	DefeatCnt *uint64 `json:"DefeatCnt,omitnil,omitempty" name:"DefeatCnt"`

	// 新增好友
	NewContactCnt *uint64 `json:"NewContactCnt,omitnil,omitempty" name:"NewContactCnt"`

	// 统计时间, 单位：天
	StatisticalTime *string `json:"StatisticalTime,omitnil,omitempty" name:"StatisticalTime"`
}

type CustomerActionEventDetail struct {
	// 事件码
	EventCode *string `json:"EventCode,omitnil,omitempty" name:"EventCode"`

	// 事件类型
	EventType *int64 `json:"EventType,omitnil,omitempty" name:"EventType"`

	// 事件来源
	EventSource *int64 `json:"EventSource,omitnil,omitempty" name:"EventSource"`

	// 外部联系人id
	ExternalUserId *string `json:"ExternalUserId,omitnil,omitempty" name:"ExternalUserId"`

	// 销售顾问id
	SalesId *uint64 `json:"SalesId,omitnil,omitempty" name:"SalesId"`

	// 素材类型
	MaterialType *int64 `json:"MaterialType,omitnil,omitempty" name:"MaterialType"`

	// 素材编号id
	MaterialId *uint64 `json:"MaterialId,omitnil,omitempty" name:"MaterialId"`

	// 事件上报时间，单位：秒
	EventTime *uint64 `json:"EventTime,omitnil,omitempty" name:"EventTime"`
}

type CustomerProfile struct {
	// 客户档案id，客户唯一识别编码
	CustomerId *uint64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 所属经销商id
	DealerCode *string `json:"DealerCode,omitnil,omitempty" name:"DealerCode"`

	// 客户在微信生态中唯一识别码
	UnionId *string `json:"UnionId,omitnil,omitempty" name:"UnionId"`

	// 档案创建时间戳，单位：秒
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 客户姓名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 0未知，1：男，2：女
	Gender *int64 `json:"Gender,omitnil,omitempty" name:"Gender"`

	// 客户联系手机号
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 客户年龄段名称
	AgeRangeName *string `json:"AgeRangeName,omitnil,omitempty" name:"AgeRangeName"`

	// 客户行业类型名称信息，如教师、医生
	JobTypeName *string `json:"JobTypeName,omitnil,omitempty" name:"JobTypeName"`

	// 客户居住地址
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// 客户所处状态
	//  0:已分配 1:未分配 1 待建档, 2 已建档， 3 已到店 4 已经试驾 5 战败申请中 6 已战败 7 已成交 
	LeadsProcessStatus *int64 `json:"LeadsProcessStatus,omitnil,omitempty" name:"LeadsProcessStatus"`

	// 客户来源类型，1：线上，2：线下
	LeadType *int64 `json:"LeadType,omitnil,omitempty" name:"LeadType"`

	// 与客户来源类型对应的渠道名称
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// 客户购车的意向等级
	LeadsLevelCode *string `json:"LeadsLevelCode,omitnil,omitempty" name:"LeadsLevelCode"`

	// 客户意向品牌编号
	VehicleBrandCode *string `json:"VehicleBrandCode,omitnil,omitempty" name:"VehicleBrandCode"`

	// 客户意向车系编号
	VehicleSeriesCode *string `json:"VehicleSeriesCode,omitnil,omitempty" name:"VehicleSeriesCode"`

	// 客户意向车型编号
	VehicleTypeCode *string `json:"VehicleTypeCode,omitnil,omitempty" name:"VehicleTypeCode"`

	// 购车用途信息
	VehiclePurpose *VehiclePurpose `json:"VehiclePurpose,omitnil,omitempty" name:"VehiclePurpose"`

	// 购车关注点信息
	PurchaseConcern []*PurchaseConcern `json:"PurchaseConcern,omitnil,omitempty" name:"PurchaseConcern"`

	// 客户所属顾问姓名
	SalesName *string `json:"SalesName,omitnil,omitempty" name:"SalesName"`

	// 客户所属顾问手机号
	SalesPhone *string `json:"SalesPhone,omitnil,omitempty" name:"SalesPhone"`

	// 客户实际到店时间戳，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealArrivalTime *uint64 `json:"RealArrivalTime,omitnil,omitempty" name:"RealArrivalTime"`

	// 客户到店完成试乘试驾时间戳，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompleteTestDriveTime *string `json:"CompleteTestDriveTime,omitnil,omitempty" name:"CompleteTestDriveTime"`

	// 客户完成下订的时间戳，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderTime *uint64 `json:"OrderTime,omitnil,omitempty" name:"OrderTime"`

	// 客户成交的时间戳，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeliveryTime *uint64 `json:"DeliveryTime,omitnil,omitempty" name:"DeliveryTime"`

	// 开票时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvoiceTime *uint64 `json:"InvoiceTime,omitnil,omitempty" name:"InvoiceTime"`

	// 完成对此客户战败审批的时间戳，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoseTime *uint64 `json:"LoseTime,omitnil,omitempty" name:"LoseTime"`

	// 线索成功获取的时间戳，单位：秒
	CreatedAtTime *uint64 `json:"CreatedAtTime,omitnil,omitempty" name:"CreatedAtTime"`

	// 线索成功导入的时间戳，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImportAtTime *uint64 `json:"ImportAtTime,omitnil,omitempty" name:"ImportAtTime"`

	// 完成线索分配的时间戳，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	DistributeTime *uint64 `json:"DistributeTime,omitnil,omitempty" name:"DistributeTime"`

	// 线索成功创建的时间戳，单位：秒
	LeadCreateTime *uint64 `json:"LeadCreateTime,omitnil,omitempty" name:"LeadCreateTime"`

	// 线索关联微信昵称
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// 线索归属部门节点
	OrgIdList []*string `json:"OrgIdList,omitnil,omitempty" name:"OrgIdList"`

	// 客户的介绍人姓名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Introducer *string `json:"Introducer,omitnil,omitempty" name:"Introducer"`

	// 客户的介绍人手机号码
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntroducerPhone *string `json:"IntroducerPhone,omitnil,omitempty" name:"IntroducerPhone"`

	// 最近一次完成跟进的时间戳，单位：秒
	FollowTime *uint64 `json:"FollowTime,omitnil,omitempty" name:"FollowTime"`

	// 最近一次计划跟进的时间戳，单位：秒
	NextFollowTime *uint64 `json:"NextFollowTime,omitnil,omitempty" name:"NextFollowTime"`

	// 已为该客户添加的企业标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnterpriseTags []*EnterpriseTag `json:"EnterpriseTags,omitnil,omitempty" name:"EnterpriseTags"`

	// 已为该客户添加的渠道标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelTags []*ChannelTag `json:"ChannelTags,omitnil,omitempty" name:"ChannelTags"`

	// 关联线索id
	LeadId *uint64 `json:"LeadId,omitnil,omitempty" name:"LeadId"`

	// 客户微信id
	WxId *string `json:"WxId,omitnil,omitempty" name:"WxId"`

	// 顾问职位
	Position *string `json:"Position,omitnil,omitempty" name:"Position"`

	// 是否关联微信 1 是 0 否
	IsBindWx *int64 `json:"IsBindWx,omitnil,omitempty" name:"IsBindWx"`

	// 是否无效
	IsInvalid *int64 `json:"IsInvalid,omitnil,omitempty" name:"IsInvalid"`

	// 无效类型
	InvalidType *string `json:"InvalidType,omitnil,omitempty" name:"InvalidType"`

	// 无效类型名称
	InvalidTypeName *string `json:"InvalidTypeName,omitnil,omitempty" name:"InvalidTypeName"`

	// 无效时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvalidTime *uint64 `json:"InvalidTime,omitnil,omitempty" name:"InvalidTime"`

	// 由顾问手动输入的无效原因文字
	InvalidRemark *string `json:"InvalidRemark,omitnil,omitempty" name:"InvalidRemark"`

	// 线索是否战败
	IsLose *int64 `json:"IsLose,omitnil,omitempty" name:"IsLose"`

	// 战败类型
	LoseType *string `json:"LoseType,omitnil,omitempty" name:"LoseType"`

	// 战败类型名称
	LoseTypeName *string `json:"LoseTypeName,omitnil,omitempty" name:"LoseTypeName"`

	// 战败申请原因
	LoseRemark *string `json:"LoseRemark,omitnil,omitempty" name:"LoseRemark"`
}

type DealerInfo struct {
	// 企微SaaS平台经销商id
	DealerId *uint64 `json:"DealerId,omitnil,omitempty" name:"DealerId"`

	// 经销商名称
	DealerName *string `json:"DealerName,omitnil,omitempty" name:"DealerName"`

	// 所属省份编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProvinceCode *string `json:"ProvinceCode,omitnil,omitempty" name:"ProvinceCode"`

	// 所属城市编号列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	CityCodeList []*string `json:"CityCodeList,omitnil,omitempty" name:"CityCodeList"`

	// 业务覆盖品牌id列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	BrandIdList []*string `json:"BrandIdList,omitnil,omitempty" name:"BrandIdList"`
}

type EnterpriseTag struct {
	// 该客户档案当前已成功关联的企业标签分组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 该客户档案当前已成功关联的企业标签名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagName *string `json:"TagName,omitnil,omitempty" name:"TagName"`

	// 该客户档案当前已成功关联的企业标签的id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagId *string `json:"TagId,omitnil,omitempty" name:"TagId"`
}

type ExternalContact struct {
	// 外部联系人的userId
	ExternalUserId *string `json:"ExternalUserId,omitnil,omitempty" name:"ExternalUserId"`

	// 外部联系人性别 0-未知 1-男性 2-女性
	Gender *int64 `json:"Gender,omitnil,omitempty" name:"Gender"`

	// 外部联系人的名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 外部联系人的类型，1表示该外部联系人是微信用户，2表示该外部联系人是企业微信用户
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 外部联系人在微信开放平台的唯一身份标识（微信unionid），通过此字段企业可将外部联系人与公众号/小程序用户关联起来。仅当联系人类型是微信用户，且企业或第三方服务商绑定了微信开发者ID有此字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnionId *string `json:"UnionId,omitnil,omitempty" name:"UnionId"`

	// 外部联系人联系电话
	// 注意：此字段可能返回 null，表示取不到有效值。
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`
}

type ExternalContactDetailPro struct {
	// 客户信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Customer *ExternalContact `json:"Customer,omitnil,omitempty" name:"Customer"`

	// 添加了此外部联系人的企业成员信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FollowUser []*FollowUserPro `json:"FollowUser,omitnil,omitempty" name:"FollowUser"`
}

type ExternalContactSimpleInfo struct {
	// 外部联系人的userId
	ExternalUserId *string `json:"ExternalUserId,omitnil,omitempty" name:"ExternalUserId"`

	// 添加了此外部联系人的企业成员userId
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 添加了此外部联系人的企业成员的姓名
	// 注意：此字段可能返回 null，表示取不到有效值。
	SalesName *string `json:"SalesName,omitnil,omitempty" name:"SalesName"`

	// 添加了此外部联系人的企业成员的归属部门id列表
	DepartmentIdList []*int64 `json:"DepartmentIdList,omitnil,omitempty" name:"DepartmentIdList"`
}

type ExternalContactTag struct {
	// 该成员添加此外部联系人所打标签的分组名称（标签功能需要企业微信升级到2.7.5及以上版本）
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 该成员添加此外部联系人所打标签名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagName *string `json:"TagName,omitnil,omitempty" name:"TagName"`

	// 该成员添加此外部联系人所打标签类型, 1-企业设置, 2-用户自定义
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 该成员添加此外部联系人所打企业标签的id，仅企业设置（type为1）的标签返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagId *string `json:"TagId,omitnil,omitempty" name:"TagId"`
}

type ExternalUserEventInfo struct {
	// 事件编码, 添加外部联系人(ADD_EXTERNAL_CUSTOMER)/成员删除外部联系人(DELETE_EXTERNAL_CUSTOMER)/外部联系人删除成员(DELETE_FOLLOW_USER)
	EventCode *string `json:"EventCode,omitnil,omitempty" name:"EventCode"`

	// 外部联系人id
	ExternalUserId *string `json:"ExternalUserId,omitnil,omitempty" name:"ExternalUserId"`

	// 企微SaaS的成员id
	SalesId *string `json:"SalesId,omitnil,omitempty" name:"SalesId"`

	// 事件上报时间戳，单位：秒
	EventTime *uint64 `json:"EventTime,omitnil,omitempty" name:"EventTime"`
}

type ExternalUserMappingInfo struct {
	// 企业主体对应的外部联系人userId
	CorpExternalUserId *string `json:"CorpExternalUserId,omitnil,omitempty" name:"CorpExternalUserId"`

	// 乐销车应用主体对应的外部联系人, 当不存在好友关系时，该字段值为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalUserId *string `json:"ExternalUserId,omitnil,omitempty" name:"ExternalUserId"`
}

type FollowInfo struct {
	// 线索id
	ClueId *uint64 `json:"ClueId,omitnil,omitempty" name:"ClueId"`

	// 客户档案id
	CustomerId *uint64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 客户姓名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 客户的手机号
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 是否逾期
	IsOverdue *int64 `json:"IsOverdue,omitnil,omitempty" name:"IsOverdue"`

	// 逾期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	OverdueTime *uint64 `json:"OverdueTime,omitnil,omitempty" name:"OverdueTime"`

	// 发生事件
	EventType *int64 `json:"EventType,omitnil,omitempty" name:"EventType"`

	// 发生事件名称
	EventTypeName *string `json:"EventTypeName,omitnil,omitempty" name:"EventTypeName"`

	// 跟进方式
	FollowWayType *string `json:"FollowWayType,omitnil,omitempty" name:"FollowWayType"`

	// 跟进方式名称
	FollowWayName *string `json:"FollowWayName,omitnil,omitempty" name:"FollowWayName"`

	// 本次跟进时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FollowTime *uint64 `json:"FollowTime,omitnil,omitempty" name:"FollowTime"`

	// 下次跟进时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextFollowTime *uint64 `json:"NextFollowTime,omitnil,omitempty" name:"NextFollowTime"`

	// 跟进记录
	FollowRecord *string `json:"FollowRecord,omitnil,omitempty" name:"FollowRecord"`
}

type FollowUser struct {
	// 添加了此外部联系人的企业成员userid
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 该成员对此外部联系人的备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 该成员对此外部联系人的描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 该成员添加此外部联系人的时间戳，单位为秒
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 该成员添加此客户的来源，具体含义详见<a href="https://work.weixin.qq.com/api/doc/90000/90135/92114#%E6%9D%A5%E6%BA%90%E5%AE%9A%E4%B9%89">来源定义</a>
	AddWay *int64 `json:"AddWay,omitnil,omitempty" name:"AddWay"`

	// 发起添加的userid，如果成员主动添加，为成员的userid；如果是客户主动添加，则为客户的外部联系人userid；如果是内部成员共享/管理员分配，则为对应的成员/管理员userid
	OperUserId *string `json:"OperUserId,omitnil,omitempty" name:"OperUserId"`

	// 该成员添加此外部联系人所打标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*ExternalContactTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type FollowUserPro struct {
	// 添加了此外部联系人的企业成员userid
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 该成员对此外部联系人的备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 该成员对此外部联系人的描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 该成员添加此外部联系人的时间戳，单位为秒
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 该成员添加此客户的来源，具体含义详见<a href="https://work.weixin.qq.com/api/doc/90000/90135/92114#%E6%9D%A5%E6%BA%90%E5%AE%9A%E4%B9%89">来源定义</a>
	AddWay *int64 `json:"AddWay,omitnil,omitempty" name:"AddWay"`

	// 发起添加的userid，如果成员主动添加，为成员的userid；如果是客户主动添加，则为客户的外部联系人userid；如果是内部成员共享/管理员分配，则为对应的成员/管理员userid
	OperUserId *string `json:"OperUserId,omitnil,omitempty" name:"OperUserId"`

	// 该成员添加此外部联系人所打标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*ExternalContactTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 添加了此外部联系人的企业成员的姓名
	// 注意：此字段可能返回 null，表示取不到有效值。
	SalesName *string `json:"SalesName,omitnil,omitempty" name:"SalesName"`

	// 企业成员的归属部门id列表
	DepartmentIdList []*int64 `json:"DepartmentIdList,omitnil,omitempty" name:"DepartmentIdList"`
}

type LicenseInfo struct {
	// license编号
	License *string `json:"License,omitnil,omitempty" name:"License"`

	// license版本；1-基础版，2-标准版，3-增值版
	LicenseEdition *int64 `json:"LicenseEdition,omitnil,omitempty" name:"LicenseEdition"`

	// 生效开始时间, 格式yyyy-MM-dd HH:mm:ss
	ResourceStartTime *string `json:"ResourceStartTime,omitnil,omitempty" name:"ResourceStartTime"`

	// 生效结束时间, 格式yyyy-MM-dd HH:mm:ss
	ResourceEndTime *string `json:"ResourceEndTime,omitnil,omitempty" name:"ResourceEndTime"`

	// 隔离截止时间, 格式yyyy-MM-dd HH:mm:ss
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsolationDeadline *string `json:"IsolationDeadline,omitnil,omitempty" name:"IsolationDeadline"`

	// 资源计划销毁时间, 格式yyyy-MM-dd HH:mm:ss
	// 注意：此字段可能返回 null，表示取不到有效值。
	DestroyTime *string `json:"DestroyTime,omitnil,omitempty" name:"DestroyTime"`

	// 资源状态，1.正常，2.隔离，3.销毁
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 扩展信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`
}

type LiveCodeDetail struct {
	// 活码id
	LiveCodeId *uint64 `json:"LiveCodeId,omitnil,omitempty" name:"LiveCodeId"`

	// 活码名称
	LiveCodeName *string `json:"LiveCodeName,omitnil,omitempty" name:"LiveCodeName"`

	// 短链url
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShortChainAddress *string `json:"ShortChainAddress,omitnil,omitempty" name:"ShortChainAddress"`

	// 活码二维码
	// 注意：此字段可能返回 null，表示取不到有效值。
	LiveCodePreview *string `json:"LiveCodePreview,omitnil,omitempty" name:"LiveCodePreview"`

	// 活动id
	ActivityId *int64 `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// 活动名称
	ActivityName *string `json:"ActivityName,omitnil,omitempty" name:"ActivityName"`

	// 活码状态，-1：删除，0：启用，1禁用，默认为0
	LiveCodeState *int64 `json:"LiveCodeState,omitnil,omitempty" name:"LiveCodeState"`

	// 活码参数，每个活码参数都是不一样的， 这个的值对应的是字符串json类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	LiveCodeData *string `json:"LiveCodeData,omitnil,omitempty" name:"LiveCodeData"`

	// 创建时间戳，单位为秒
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间戳，单位为秒
	UpdateTime *uint64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type MaterialInfo struct {
	// 素材id
	MaterialId *uint64 `json:"MaterialId,omitnil,omitempty" name:"MaterialId"`

	// 素材名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaterialName *string `json:"MaterialName,omitnil,omitempty" name:"MaterialName"`

	// 素材状态, -1: 删除 0: 启用 1: 禁用
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type MiniAppCodeInfo struct {
	// 主键id
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 小程序名称
	MiniAppName *string `json:"MiniAppName,omitnil,omitempty" name:"MiniAppName"`

	// 小程序logo
	MiniAppLogo *string `json:"MiniAppLogo,omitnil,omitempty" name:"MiniAppLogo"`

	// 小程序管理端地址
	MiniAdminUrl *string `json:"MiniAdminUrl,omitnil,omitempty" name:"MiniAdminUrl"`

	// 状态：0正常，1删除
	State *int64 `json:"State,omitnil,omitempty" name:"State"`

	// 创建时间戳，单位为秒
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间戳，单位为秒
	UpdateTime *uint64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type PurchaseConcern struct {
	// 购车关注点code
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 购车关注点描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

// Predefined struct for user
type QueryActivityJoinListRequestParams struct {
	// 活动id
	ActivityId *int64 `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// 分页游标，对应结果返回的NextCursor,首次请求保持为空
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// 单页数据限制
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type QueryActivityJoinListRequest struct {
	*tchttp.BaseRequest
	
	// 活动id
	ActivityId *int64 `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// 分页游标，对应结果返回的NextCursor,首次请求保持为空
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// 单页数据限制
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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

// Predefined struct for user
type QueryActivityJoinListResponseParams struct {
	// 分页游标
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextCursor *string `json:"NextCursor,omitnil,omitempty" name:"NextCursor"`

	// 活码列表响应参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageData []*ActivityJoinDetail `json:"PageData,omitnil,omitempty" name:"PageData"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryActivityJoinListResponse struct {
	*tchttp.BaseResponse
	Response *QueryActivityJoinListResponseParams `json:"Response"`
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

// Predefined struct for user
type QueryActivityListRequestParams struct {
	// 分页游标，对应结果返回的NextCursor,首次请求保持为空
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// 单页数据限制
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type QueryActivityListRequest struct {
	*tchttp.BaseRequest
	
	// 分页游标，对应结果返回的NextCursor,首次请求保持为空
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// 单页数据限制
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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

// Predefined struct for user
type QueryActivityListResponseParams struct {
	// 分页游标
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextCursor *string `json:"NextCursor,omitnil,omitempty" name:"NextCursor"`

	// 活码列表响应参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageData []*ActivityDetail `json:"PageData,omitnil,omitempty" name:"PageData"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryActivityListResponse struct {
	*tchttp.BaseResponse
	Response *QueryActivityListResponseParams `json:"Response"`
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

// Predefined struct for user
type QueryActivityLiveCodeListRequestParams struct {
	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// 返回的最大记录数，整型，最大值100，默认值50，超过最大值时取最大值
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type QueryActivityLiveCodeListRequest struct {
	*tchttp.BaseRequest
	
	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// 返回的最大记录数，整型，最大值100，默认值50，超过最大值时取最大值
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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

// Predefined struct for user
type QueryActivityLiveCodeListResponseParams struct {
	// 分页游标，再下次请求时填写以获取之后分页的记录，如果已经没有更多的数据则返回空
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextCursor *string `json:"NextCursor,omitnil,omitempty" name:"NextCursor"`

	// 活码列表响应参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageData []*LiveCodeDetail `json:"PageData,omitnil,omitempty" name:"PageData"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryActivityLiveCodeListResponse struct {
	*tchttp.BaseResponse
	Response *QueryActivityLiveCodeListResponseParams `json:"Response"`
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

// Predefined struct for user
type QueryArrivalListRequestParams struct {
	// 分页，预期请求的数据量，取值范围 1 ~ 1000
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询开始时间， 单位秒
	BeginTime *uint64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 查询结束时间， 单位秒
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`
}

type QueryArrivalListRequest struct {
	*tchttp.BaseRequest
	
	// 分页，预期请求的数据量，取值范围 1 ~ 1000
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询开始时间， 单位秒
	BeginTime *uint64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 查询结束时间， 单位秒
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`
}

func (r *QueryArrivalListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryArrivalListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "Cursor")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryArrivalListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryArrivalListResponseParams struct {
	// 分页游标，下次调用带上该值，则从当前的位置继续往后拉，以实现增量拉取。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextCursor *string `json:"NextCursor,omitnil,omitempty" name:"NextCursor"`

	// 潜客客户存档信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageData []*ArrivalInfo `json:"PageData,omitnil,omitempty" name:"PageData"`

	// 是否还有更多数据。0-否；1-是。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HasMore *uint64 `json:"HasMore,omitnil,omitempty" name:"HasMore"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryArrivalListResponse struct {
	*tchttp.BaseResponse
	Response *QueryArrivalListResponseParams `json:"Response"`
}

func (r *QueryArrivalListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryArrivalListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryChannelCodeListRequestParams struct {
	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// 返回的最大记录数，整型，最大值100，默认值50，超过最大值时取最大值
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type QueryChannelCodeListRequest struct {
	*tchttp.BaseRequest
	
	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// 返回的最大记录数，整型，最大值100，默认值50，超过最大值时取最大值
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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

// Predefined struct for user
type QueryChannelCodeListResponseParams struct {
	// 分页游标，再下次请求时填写以获取之后分页的记录，如果已经没有更多的数据则返回空
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextCursor *string `json:"NextCursor,omitnil,omitempty" name:"NextCursor"`

	// 活码列表响应参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageData []*ChannelCodeInnerDetail `json:"PageData,omitnil,omitempty" name:"PageData"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryChannelCodeListResponse struct {
	*tchttp.BaseResponse
	Response *QueryChannelCodeListResponseParams `json:"Response"`
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

// Predefined struct for user
type QueryChatArchivingListRequestParams struct {
	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// 返回的最大记录数，整型，最大值100，默认值50，超过最大值时取最大值
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type QueryChatArchivingListRequest struct {
	*tchttp.BaseRequest
	
	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// 返回的最大记录数，整型，最大值100，默认值50，超过最大值时取最大值
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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

// Predefined struct for user
type QueryChatArchivingListResponseParams struct {
	// 分页游标，再下次请求时填写以获取之后分页的记录，如果已经没有更多的数据则返回空
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextCursor *string `json:"NextCursor,omitnil,omitempty" name:"NextCursor"`

	// 会话存档列表响应参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageData []*ChatArchivingDetail `json:"PageData,omitnil,omitempty" name:"PageData"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryChatArchivingListResponse struct {
	*tchttp.BaseResponse
	Response *QueryChatArchivingListResponseParams `json:"Response"`
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

// Predefined struct for user
type QueryClueInfoListRequestParams struct {
	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// 返回的最大记录数，整型，最大值100，默认值50，超过最大值时取最大值
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询开始时间， 单位秒
	BeginTime *uint64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 查询结束时间， 单位秒
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type QueryClueInfoListRequest struct {
	*tchttp.BaseRequest
	
	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// 返回的最大记录数，整型，最大值100，默认值50，超过最大值时取最大值
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询开始时间， 单位秒
	BeginTime *uint64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 查询结束时间， 单位秒
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *QueryClueInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryClueInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Cursor")
	delete(f, "Limit")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryClueInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryClueInfoListResponseParams struct {
	// 线索信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageData []*ClueInfoDetail `json:"PageData,omitnil,omitempty" name:"PageData"`

	// 分页游标，下次调用带上该值，则从当前的位置继续往后拉，以实现增量拉取。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextCursor *string `json:"NextCursor,omitnil,omitempty" name:"NextCursor"`

	// 是否还有更多数据。0-否；1-是。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HasMore *uint64 `json:"HasMore,omitnil,omitempty" name:"HasMore"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryClueInfoListResponse struct {
	*tchttp.BaseResponse
	Response *QueryClueInfoListResponseParams `json:"Response"`
}

func (r *QueryClueInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryClueInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryCrmStatisticsRequestParams struct {
	// 查询开始时间， 单位秒
	BeginTime *uint64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 查询结束时间， 单位秒
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// 返回的最大记录数，整型，最大值100，默认值50，超过最大值时取最大值
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 请求的企业成员id，为空时默认全租户
	SalesId *string `json:"SalesId,omitnil,omitempty" name:"SalesId"`

	// 请求的部门id，为空时默认全租户
	OrgId *uint64 `json:"OrgId,omitnil,omitempty" name:"OrgId"`
}

type QueryCrmStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 查询开始时间， 单位秒
	BeginTime *uint64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 查询结束时间， 单位秒
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// 返回的最大记录数，整型，最大值100，默认值50，超过最大值时取最大值
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 请求的企业成员id，为空时默认全租户
	SalesId *string `json:"SalesId,omitnil,omitempty" name:"SalesId"`

	// 请求的部门id，为空时默认全租户
	OrgId *uint64 `json:"OrgId,omitnil,omitempty" name:"OrgId"`
}

func (r *QueryCrmStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCrmStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "Cursor")
	delete(f, "Limit")
	delete(f, "SalesId")
	delete(f, "OrgId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryCrmStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryCrmStatisticsResponseParams struct {
	// 分页游标，在下次请求时填写以获取之后分页的记录，如果已经没有更多的数据则返回空
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextCursor *string `json:"NextCursor,omitnil,omitempty" name:"NextCursor"`

	// CRM统计响应数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageData []*CrmStatisticsData `json:"PageData,omitnil,omitempty" name:"PageData"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryCrmStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *QueryCrmStatisticsResponseParams `json:"Response"`
}

func (r *QueryCrmStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCrmStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryCustomerEventDetailStatisticsRequestParams struct {
	// 查询开始时间， 单位秒
	BeginTime *uint64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 查询结束时间， 单位秒
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// 返回的最大记录数，整型，最大值100，默认值50，超过最大值时取最大值
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type QueryCustomerEventDetailStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 查询开始时间， 单位秒
	BeginTime *uint64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 查询结束时间， 单位秒
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// 返回的最大记录数，整型，最大值100，默认值50，超过最大值时取最大值
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *QueryCustomerEventDetailStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCustomerEventDetailStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "Cursor")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryCustomerEventDetailStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryCustomerEventDetailStatisticsResponseParams struct {
	// 分页游标，在下次请求时填写以获取之后分页的记录，如果已经没有更多的数据则返回空
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextCursor *string `json:"NextCursor,omitnil,omitempty" name:"NextCursor"`

	// 外部联系人SaaS使用明细统计响应数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageData []*CustomerActionEventDetail `json:"PageData,omitnil,omitempty" name:"PageData"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryCustomerEventDetailStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *QueryCustomerEventDetailStatisticsResponseParams `json:"Response"`
}

func (r *QueryCustomerEventDetailStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCustomerEventDetailStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryCustomerProfileListRequestParams struct {
	// 分页，预期请求的数据量，取值范围 1 ~ 1000
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询开始时间， 单位秒
	BeginTime *uint64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 查询结束时间， 单位秒
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`
}

type QueryCustomerProfileListRequest struct {
	*tchttp.BaseRequest
	
	// 分页，预期请求的数据量，取值范围 1 ~ 1000
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询开始时间， 单位秒
	BeginTime *uint64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 查询结束时间， 单位秒
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`
}

func (r *QueryCustomerProfileListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCustomerProfileListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "Cursor")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryCustomerProfileListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryCustomerProfileListResponseParams struct {
	// 分页游标，下次调用带上该值，则从当前的位置继续往后拉，以实现增量拉取。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextCursor *string `json:"NextCursor,omitnil,omitempty" name:"NextCursor"`

	// 潜客客户存档信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageData []*CustomerProfile `json:"PageData,omitnil,omitempty" name:"PageData"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryCustomerProfileListResponse struct {
	*tchttp.BaseResponse
	Response *QueryCustomerProfileListResponseParams `json:"Response"`
}

func (r *QueryCustomerProfileListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCustomerProfileListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryDealerInfoListRequestParams struct {
	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// 返回的最大记录数，整型，最大值100，默认值50，超过最大值时取最大值
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type QueryDealerInfoListRequest struct {
	*tchttp.BaseRequest
	
	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// 返回的最大记录数，整型，最大值100，默认值50，超过最大值时取最大值
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *QueryDealerInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryDealerInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Cursor")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryDealerInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryDealerInfoListResponseParams struct {
	// 经销商信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageData []*DealerInfo `json:"PageData,omitnil,omitempty" name:"PageData"`

	// 分页游标，下次调用带上该值，则从当前的位置继续往后拉取新增的数据，以实现增量拉取。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextCursor *string `json:"NextCursor,omitnil,omitempty" name:"NextCursor"`

	// 是否还有更多数据。0-否；1-是。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HasMore *uint64 `json:"HasMore,omitnil,omitempty" name:"HasMore"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryDealerInfoListResponse struct {
	*tchttp.BaseResponse
	Response *QueryDealerInfoListResponseParams `json:"Response"`
}

func (r *QueryDealerInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryDealerInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryExternalContactDetailByDateRequestParams struct {
	// 查询结束时间， 单位秒
	BeginTime *uint64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 查询结束时间， 单位秒
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// 返回的最大记录数，整型，最大值100，默认值50，超过最大值时取最大值
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type QueryExternalContactDetailByDateRequest struct {
	*tchttp.BaseRequest
	
	// 查询结束时间， 单位秒
	BeginTime *uint64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 查询结束时间， 单位秒
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// 返回的最大记录数，整型，最大值100，默认值50，超过最大值时取最大值
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *QueryExternalContactDetailByDateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryExternalContactDetailByDateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "Cursor")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryExternalContactDetailByDateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryExternalContactDetailByDateResponseParams struct {
	// 分页游标，再下次请求时填写以获取之后分页的记录，如果已经没有更多的数据则返回空
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextCursor *string `json:"NextCursor,omitnil,omitempty" name:"NextCursor"`

	// 外部联系人详细信息
	PageData []*ExternalContactDetailPro `json:"PageData,omitnil,omitempty" name:"PageData"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryExternalContactDetailByDateResponse struct {
	*tchttp.BaseResponse
	Response *QueryExternalContactDetailByDateResponseParams `json:"Response"`
}

func (r *QueryExternalContactDetailByDateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryExternalContactDetailByDateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryExternalContactDetailRequestParams struct {
	// 外部联系人的userid，注意不是企业成员的账号
	ExternalUserId *string `json:"ExternalUserId,omitnil,omitempty" name:"ExternalUserId"`

	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填。当客户在企业内的跟进人超过500人时需要使用cursor参数进行分页获取
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// 当前接口Limit不需要传参， 保留Limit只是为了保持向后兼容性， Limit默认值为500，当返回结果超过500时， NextCursor才有返回值
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type QueryExternalContactDetailRequest struct {
	*tchttp.BaseRequest
	
	// 外部联系人的userid，注意不是企业成员的账号
	ExternalUserId *string `json:"ExternalUserId,omitnil,omitempty" name:"ExternalUserId"`

	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填。当客户在企业内的跟进人超过500人时需要使用cursor参数进行分页获取
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// 当前接口Limit不需要传参， 保留Limit只是为了保持向后兼容性， Limit默认值为500，当返回结果超过500时， NextCursor才有返回值
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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

// Predefined struct for user
type QueryExternalContactDetailResponseParams struct {
	// 分页游标，再下次请求时填写以获取之后分页的记录，如果已经没有更多的数据则返回空
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextCursor *string `json:"NextCursor,omitnil,omitempty" name:"NextCursor"`

	// 客户信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Customer *ExternalContact `json:"Customer,omitnil,omitempty" name:"Customer"`

	// 添加了此外部联系人的企业成员信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FollowUser []*FollowUser `json:"FollowUser,omitnil,omitempty" name:"FollowUser"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryExternalContactDetailResponse struct {
	*tchttp.BaseResponse
	Response *QueryExternalContactDetailResponseParams `json:"Response"`
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

// Predefined struct for user
type QueryExternalContactListRequestParams struct {
	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// 返回的最大记录数，整型，最大值100，默认值50，超过最大值时取最大值
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type QueryExternalContactListRequest struct {
	*tchttp.BaseRequest
	
	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// 返回的最大记录数，整型，最大值100，默认值50，超过最大值时取最大值
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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

// Predefined struct for user
type QueryExternalContactListResponseParams struct {
	// 外部联系人信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageData []*ExternalContactSimpleInfo `json:"PageData,omitnil,omitempty" name:"PageData"`

	// 分页游标，再下次请求时填写以获取之后分页的记录，如果已经没有更多的数据则返回空
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextCursor *string `json:"NextCursor,omitnil,omitempty" name:"NextCursor"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryExternalContactListResponse struct {
	*tchttp.BaseResponse
	Response *QueryExternalContactListResponseParams `json:"Response"`
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

// Predefined struct for user
type QueryExternalUserEventListRequestParams struct {
	// 查询开始时间， 单位秒
	BeginTime *uint64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 查询结束时间， 单位秒
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// 返回的最大记录数，整型，最大值100，默认值50，超过最大值时取最大值
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type QueryExternalUserEventListRequest struct {
	*tchttp.BaseRequest
	
	// 查询开始时间， 单位秒
	BeginTime *uint64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 查询结束时间， 单位秒
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// 返回的最大记录数，整型，最大值100，默认值50，超过最大值时取最大值
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *QueryExternalUserEventListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryExternalUserEventListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "Cursor")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryExternalUserEventListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryExternalUserEventListResponseParams struct {
	// 分页游标，再下次请求时填写以获取之后分页的记录，如果已经没有更多的数据则返回空
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextCursor *string `json:"NextCursor,omitnil,omitempty" name:"NextCursor"`

	// 外部联系人事件信息响应数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageData []*ExternalUserEventInfo `json:"PageData,omitnil,omitempty" name:"PageData"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryExternalUserEventListResponse struct {
	*tchttp.BaseResponse
	Response *QueryExternalUserEventListResponseParams `json:"Response"`
}

func (r *QueryExternalUserEventListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryExternalUserEventListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryExternalUserMappingInfoRequestParams struct {
	// 企业主体对应的外部联系人id列表，列表长度限制最大为50。
	CorpExternalUserIdList []*string `json:"CorpExternalUserIdList,omitnil,omitempty" name:"CorpExternalUserIdList"`
}

type QueryExternalUserMappingInfoRequest struct {
	*tchttp.BaseRequest
	
	// 企业主体对应的外部联系人id列表，列表长度限制最大为50。
	CorpExternalUserIdList []*string `json:"CorpExternalUserIdList,omitnil,omitempty" name:"CorpExternalUserIdList"`
}

func (r *QueryExternalUserMappingInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryExternalUserMappingInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CorpExternalUserIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryExternalUserMappingInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryExternalUserMappingInfoResponseParams struct {
	// 外部联系人映射信息, 只返回映射成功的记录
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalUserIdMapping []*ExternalUserMappingInfo `json:"ExternalUserIdMapping,omitnil,omitempty" name:"ExternalUserIdMapping"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryExternalUserMappingInfoResponse struct {
	*tchttp.BaseResponse
	Response *QueryExternalUserMappingInfoResponseParams `json:"Response"`
}

func (r *QueryExternalUserMappingInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryExternalUserMappingInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryFollowListRequestParams struct {
	// 分页，预期请求的数据量，取值范围 1 ~ 1000
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询开始时间， 单位秒
	BeginTime *uint64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 查询结束时间， 单位秒
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`
}

type QueryFollowListRequest struct {
	*tchttp.BaseRequest
	
	// 分页，预期请求的数据量，取值范围 1 ~ 1000
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询开始时间， 单位秒
	BeginTime *uint64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 查询结束时间， 单位秒
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`
}

func (r *QueryFollowListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryFollowListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "Cursor")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryFollowListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryFollowListResponseParams struct {
	// 分页游标，下次调用带上该值，则从当前的位置继续往后拉，以实现增量拉取。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextCursor *string `json:"NextCursor,omitnil,omitempty" name:"NextCursor"`

	// 潜客客户存档信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageData []*FollowInfo `json:"PageData,omitnil,omitempty" name:"PageData"`

	// 是否还有更多数据。0-否；1-是。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HasMore *uint64 `json:"HasMore,omitnil,omitempty" name:"HasMore"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryFollowListResponse struct {
	*tchttp.BaseResponse
	Response *QueryFollowListResponseParams `json:"Response"`
}

func (r *QueryFollowListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryFollowListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryLicenseInfoRequestParams struct {
	// license编号
	License *string `json:"License,omitnil,omitempty" name:"License"`
}

type QueryLicenseInfoRequest struct {
	*tchttp.BaseRequest
	
	// license编号
	License *string `json:"License,omitnil,omitempty" name:"License"`
}

func (r *QueryLicenseInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryLicenseInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "License")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryLicenseInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryLicenseInfoResponseParams struct {
	// license响应信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	LicenseInfo *LicenseInfo `json:"LicenseInfo,omitnil,omitempty" name:"LicenseInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryLicenseInfoResponse struct {
	*tchttp.BaseResponse
	Response *QueryLicenseInfoResponseParams `json:"Response"`
}

func (r *QueryLicenseInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryLicenseInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryMaterialListRequestParams struct {
	// 素材类型：0-图片，1-视频，3-文章，10-车型，11-名片
	MaterialType *int64 `json:"MaterialType,omitnil,omitempty" name:"MaterialType"`

	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// 返回的最大记录数，整型，最大值100，默认值50，超过最大值时取最大值
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type QueryMaterialListRequest struct {
	*tchttp.BaseRequest
	
	// 素材类型：0-图片，1-视频，3-文章，10-车型，11-名片
	MaterialType *int64 `json:"MaterialType,omitnil,omitempty" name:"MaterialType"`

	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// 返回的最大记录数，整型，最大值100，默认值50，超过最大值时取最大值
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *QueryMaterialListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryMaterialListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MaterialType")
	delete(f, "Cursor")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryMaterialListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryMaterialListResponseParams struct {
	// 分页游标，再下次请求时填写以获取之后分页的记录，如果已经没有更多的数据则返回空
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextCursor *string `json:"NextCursor,omitnil,omitempty" name:"NextCursor"`

	// 企业素材列表响应数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageData []*MaterialInfo `json:"PageData,omitnil,omitempty" name:"PageData"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryMaterialListResponse struct {
	*tchttp.BaseResponse
	Response *QueryMaterialListResponseParams `json:"Response"`
}

func (r *QueryMaterialListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryMaterialListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryMiniAppCodeListRequestParams struct {
	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// 返回的最大记录数，整型，最大值100，默认值50，超过最大值时取最大值
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type QueryMiniAppCodeListRequest struct {
	*tchttp.BaseRequest
	
	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// 返回的最大记录数，整型，最大值100，默认值50，超过最大值时取最大值
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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

// Predefined struct for user
type QueryMiniAppCodeListResponseParams struct {
	// 分页游标，再下次请求时填写以获取之后分页的记录，如果已经没有更多的数据则返回空
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextCursor *string `json:"NextCursor,omitnil,omitempty" name:"NextCursor"`

	// 小程序码列表响应数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageData []*MiniAppCodeInfo `json:"PageData,omitnil,omitempty" name:"PageData"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryMiniAppCodeListResponse struct {
	*tchttp.BaseResponse
	Response *QueryMiniAppCodeListResponseParams `json:"Response"`
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

// Predefined struct for user
type QueryStaffEventDetailStatisticsRequestParams struct {
	// 查询开始时间， 单位秒
	BeginTime *uint64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 查询结束时间， 单位秒
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// 返回的最大记录数，整型，最大值100，默认值50，超过最大值时取最大值
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type QueryStaffEventDetailStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 查询开始时间， 单位秒
	BeginTime *uint64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 查询结束时间， 单位秒
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// 返回的最大记录数，整型，最大值100，默认值50，超过最大值时取最大值
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *QueryStaffEventDetailStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryStaffEventDetailStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "Cursor")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryStaffEventDetailStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryStaffEventDetailStatisticsResponseParams struct {
	// 分页游标，在下次请求时填写以获取之后分页的记录，如果已经没有更多的数据则返回空
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextCursor *string `json:"NextCursor,omitnil,omitempty" name:"NextCursor"`

	// 企业成员SaaS使用明细统计响应数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageData []*SalesActionEventDetail `json:"PageData,omitnil,omitempty" name:"PageData"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryStaffEventDetailStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *QueryStaffEventDetailStatisticsResponseParams `json:"Response"`
}

func (r *QueryStaffEventDetailStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryStaffEventDetailStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryUserInfoListRequestParams struct {
	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// 返回的最大记录数，整型，最大值100，默认值50，超过最大值时取最大值
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type QueryUserInfoListRequest struct {
	*tchttp.BaseRequest
	
	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// 返回的最大记录数，整型，最大值100，默认值50，超过最大值时取最大值
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *QueryUserInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryUserInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Cursor")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryUserInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryUserInfoListResponseParams struct {
	// 分页游标，再下次请求时填写以获取之后分页的记录，如果已经没有更多的数据则返回空
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextCursor *string `json:"NextCursor,omitnil,omitempty" name:"NextCursor"`

	// 企业成员信息列表响应数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageData []*CorpUserInfo `json:"PageData,omitnil,omitempty" name:"PageData"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryUserInfoListResponse struct {
	*tchttp.BaseResponse
	Response *QueryUserInfoListResponseParams `json:"Response"`
}

func (r *QueryUserInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryUserInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryVehicleInfoListRequestParams struct {
	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// 返回的最大记录数，整型，最大值100，默认值50，超过最大值时取最大值
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type QueryVehicleInfoListRequest struct {
	*tchttp.BaseRequest
	
	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// 返回的最大记录数，整型，最大值100，默认值50，超过最大值时取最大值
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *QueryVehicleInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryVehicleInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Cursor")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryVehicleInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryVehicleInfoListResponseParams struct {
	// 车系车型信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageData []*VehicleInfo `json:"PageData,omitnil,omitempty" name:"PageData"`

	// 分页游标，下次调用带上该值，则从当前的位置继续往后拉取新增的数据，以实现增量拉取。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextCursor *string `json:"NextCursor,omitnil,omitempty" name:"NextCursor"`

	// 是否还有更多数据。0-否；1-是。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HasMore *uint64 `json:"HasMore,omitnil,omitempty" name:"HasMore"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryVehicleInfoListResponse struct {
	*tchttp.BaseResponse
	Response *QueryVehicleInfoListResponseParams `json:"Response"`
}

func (r *QueryVehicleInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryVehicleInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SalesActionEventDetail struct {
	// 事件码
	EventCode *string `json:"EventCode,omitnil,omitempty" name:"EventCode"`

	// 事件类型
	EventType *int64 `json:"EventType,omitnil,omitempty" name:"EventType"`

	// 事件来源
	EventSource *int64 `json:"EventSource,omitnil,omitempty" name:"EventSource"`

	// 销售顾问id
	SalesId *uint64 `json:"SalesId,omitnil,omitempty" name:"SalesId"`

	// 素材类型
	MaterialType *int64 `json:"MaterialType,omitnil,omitempty" name:"MaterialType"`

	// 素材编号id
	MaterialId *uint64 `json:"MaterialId,omitnil,omitempty" name:"MaterialId"`

	// 事件上报时间，单位：秒
	EventTime *uint64 `json:"EventTime,omitnil,omitempty" name:"EventTime"`
}

type TagDetailInfo struct {
	// 标签名称
	TagName *string `json:"TagName,omitnil,omitempty" name:"TagName"`

	// 标签业务ID
	BizTagId *string `json:"BizTagId,omitnil,omitempty" name:"BizTagId"`

	// 企微标签ID
	TagId *string `json:"TagId,omitnil,omitempty" name:"TagId"`

	// 标签排序的次序值，sort值大的排序靠前。有效的值范围是[0, 2^32)
	Sort *uint64 `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 标签创建时间,单位为秒
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type TagGroup struct {
	// 企微标签组id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 标签组业务id
	BizGroupId *string `json:"BizGroupId,omitnil,omitempty" name:"BizGroupId"`

	// 企微标签组名称，不能超过15个字符
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 标签组次序值。sort值大的排序靠前。有效的值范围是[0, 2^32)
	Sort *uint64 `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 标签组创建时间,单位为秒
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 标签组内的标签列表, 上限为20
	Tags []*TagDetailInfo `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type TagInfo struct {
	// 标签名称, 最大长度限制15个字符
	TagName *string `json:"TagName,omitnil,omitempty" name:"TagName"`

	// 标签组排序,值越大,排序越靠前
	Sort *uint64 `json:"Sort,omitnil,omitempty" name:"Sort"`
}

type VehicleInfo struct {
	// 品牌id
	BrandId *uint64 `json:"BrandId,omitnil,omitempty" name:"BrandId"`

	// 品牌名称
	BrandName *string `json:"BrandName,omitnil,omitempty" name:"BrandName"`

	// 车系id
	SeriesId *uint64 `json:"SeriesId,omitnil,omitempty" name:"SeriesId"`

	// 车系名称
	SeriesName *string `json:"SeriesName,omitnil,omitempty" name:"SeriesName"`

	// 车型id
	ModelId *uint64 `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 车型名称
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`
}

type VehiclePurpose struct {
	// 购车用途code
	VehiclePurposeCode *string `json:"VehiclePurposeCode,omitnil,omitempty" name:"VehiclePurposeCode"`

	// 购车用途名称
	VehiclePurposeName *string `json:"VehiclePurposeName,omitnil,omitempty" name:"VehiclePurposeName"`
}

type WeComTagDetail struct {
	// 标签分组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 标签分组业务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	BizGroupId *string `json:"BizGroupId,omitnil,omitempty" name:"BizGroupId"`

	// 标签名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagName *string `json:"TagName,omitnil,omitempty" name:"TagName"`

	// 标签ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagId *string `json:"TagId,omitnil,omitempty" name:"TagId"`

	// 标签业务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	BizTagId *string `json:"BizTagId,omitnil,omitempty" name:"BizTagId"`

	// 标签分类，1：企业设置、2：用户自定义
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 标签业务ID字符串格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	BizTagIdStr *string `json:"BizTagIdStr,omitnil,omitempty" name:"BizTagIdStr"`
}