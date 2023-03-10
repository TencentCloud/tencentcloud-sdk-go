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

package v20220805

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type DescribeGoodsRecommendRequestParams struct {
	// 实例ID，在控制台获取
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 场景ID，在控制台创建场景后获取
	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`

	// 用户唯一ID，客户自定义用户ID，作为一个用户的唯一标识，需和行为数据上报接口中的UserId一致，否则影响特征关联
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户设备ID数组，可传入用户的多个类型ID，用于关联画像信息
	UserIdList []*StrUserIdInfo `json:"UserIdList,omitempty" name:"UserIdList"`

	// 推荐返回数量，默认10个，最多支持50个的内容返回。如果有更多数量要求，<a href="https://console.cloud.tencent.com/workorder/category" target="_blank">提单</a>沟通解决
	GoodsCnt *int64 `json:"GoodsCnt,omitempty" name:"GoodsCnt"`

	// 当场景是相关推荐时该值必填，场景是非相关推荐时该值无效
	CurrentGoodsId *string `json:"CurrentGoodsId,omitempty" name:"CurrentGoodsId"`

	// 用户的实时特征信息，用作特征
	UserPortraitInfo *UserPortraitInfo `json:"UserPortraitInfo,omitempty" name:"UserPortraitInfo"`

	// 本次请求针对该用户需要过滤的物品列表(不超过100个)
	BlackGoodsList []*string `json:"BlackGoodsList,omitempty" name:"BlackGoodsList"`

	// json字符串，扩展字段
	Extension *string `json:"Extension,omitempty" name:"Extension"`
}

type DescribeGoodsRecommendRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，在控制台获取
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 场景ID，在控制台创建场景后获取
	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`

	// 用户唯一ID，客户自定义用户ID，作为一个用户的唯一标识，需和行为数据上报接口中的UserId一致，否则影响特征关联
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户设备ID数组，可传入用户的多个类型ID，用于关联画像信息
	UserIdList []*StrUserIdInfo `json:"UserIdList,omitempty" name:"UserIdList"`

	// 推荐返回数量，默认10个，最多支持50个的内容返回。如果有更多数量要求，<a href="https://console.cloud.tencent.com/workorder/category" target="_blank">提单</a>沟通解决
	GoodsCnt *int64 `json:"GoodsCnt,omitempty" name:"GoodsCnt"`

	// 当场景是相关推荐时该值必填，场景是非相关推荐时该值无效
	CurrentGoodsId *string `json:"CurrentGoodsId,omitempty" name:"CurrentGoodsId"`

	// 用户的实时特征信息，用作特征
	UserPortraitInfo *UserPortraitInfo `json:"UserPortraitInfo,omitempty" name:"UserPortraitInfo"`

	// 本次请求针对该用户需要过滤的物品列表(不超过100个)
	BlackGoodsList []*string `json:"BlackGoodsList,omitempty" name:"BlackGoodsList"`

	// json字符串，扩展字段
	Extension *string `json:"Extension,omitempty" name:"Extension"`
}

func (r *DescribeGoodsRecommendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGoodsRecommendRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SceneId")
	delete(f, "UserId")
	delete(f, "UserIdList")
	delete(f, "GoodsCnt")
	delete(f, "CurrentGoodsId")
	delete(f, "UserPortraitInfo")
	delete(f, "BlackGoodsList")
	delete(f, "Extension")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGoodsRecommendRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGoodsRecommendResponseParams struct {
	// 推荐返回的商品信息列表
	DataList []*RecGoodsData `json:"DataList,omitempty" name:"DataList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeGoodsRecommendResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGoodsRecommendResponseParams `json:"Response"`
}

func (r *DescribeGoodsRecommendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGoodsRecommendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DislikeInfo struct {
	// 过滤的类别：<br>● author 作者名<br/>（如当前类型不满足，请<a href="https://console.cloud.tencent.com/workorder/category" target="_blank">提单</a>沟通解决方案）
	Type *string `json:"Type,omitempty" name:"Type"`

	// Type对应字段名的值，如：需要过滤的作者名
	Value *string `json:"Value,omitempty" name:"Value"`
}

type DocItem struct {
	// 内容唯一id，建议限制在128字符以内
	ItemId *string `json:"ItemId,omitempty" name:"ItemId"`

	// 内容类型：<br/>● article -图文<br>● text -纯文本<br/>● video -视频<br/>● short_video -时长15秒以内的视频<br/>● mini_video -竖屏视频<br/>● image -纯图片<br/>（如当前类型不满足，请登录控制台进入对应项目，在<b>物料管理->物料类型管理</b>中添加）
	ItemType *string `json:"ItemType,omitempty" name:"ItemType"`

	// 内容状态：
	// ● 1 - 上架 
	// ● 2 - 下架 
	// Status=2的内容不会在推荐结果中出现 
	// 需要下架内容时，把Status的值修改为2即可
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 内容生成时间，秒级时间戳（1639624786），需大于0，<b>用作特征和物料管理</b>
	PublishTimestamp *int64 `json:"PublishTimestamp,omitempty" name:"PublishTimestamp"`

	// 内容过期时间，秒级时间戳（1639624786），如未填，则默认PublishTimestamp往后延一年，用作特征，过期则不会被推荐，<b>强烈建议</b>
	ExpireTimestamp *int64 `json:"ExpireTimestamp,omitempty" name:"ExpireTimestamp"`

	// 类目层级数，例如3级类目，则填3，和CategoryPath字段的类数据匹配，<b>强烈建议</b>
	CategoryLevel *int64 `json:"CategoryLevel,omitempty" name:"CategoryLevel"`

	// 类目路径，一级二级三级等依次用英文冒号联接，和CategoryLevel字段值匹配，如体育：“足球:巴塞罗那”。<b>用于物料池管理，强烈建议</b>
	CategoryPath *string `json:"CategoryPath,omitempty" name:"CategoryPath"`

	// 内容标签，多个标签用英文冒号联接，<b>用作特征，强烈建议</b>
	Tags *string `json:"Tags,omitempty" name:"Tags"`

	// 作者名，需保证作者名唯一，若有重名需要加编号区分。<b>用于召回过滤、规则打散，强烈建议</b>
	Author *string `json:"Author,omitempty" name:"Author"`

	// 内容来源类型，客户自定义，<b>用于物料池管理</b>
	SourceId *string `json:"SourceId,omitempty" name:"SourceId"`

	// 内容标题，<b>主要用于语义分析</b>
	Title *string `json:"Title,omitempty" name:"Title"`

	// 正文关键片段，建议控制在500字符以内，<b>主要用于语义分析</b>
	Content *string `json:"Content,omitempty" name:"Content"`

	// 正文详情，主要用于语义分析，当内容过大时建议用ContentUrl传递，<b>与Content可二选一</b>
	ContentUrl *string `json:"ContentUrl,omitempty" name:"ContentUrl"`

	// 视频时长，时间秒，大于等于0，小于 3600 * 10。<b>视频内容必填，其它内容非必填，用作特征</b>
	VideoDuration *int64 `json:"VideoDuration,omitempty" name:"VideoDuration"`

	// 国家，ISO 3166-1 alpha-2编码，参考<a href="https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2" target="_blank">ISO 3166-1 alpha-2</a>，中国：“CN”，<b>用作特征</b>
	Country *string `json:"Country,omitempty" name:"Country"`

	// 省份，ISO 3166-2行政区编码，如中国参考<a href="https://zh.wikipedia.org/wiki/ISO_3166-2:CN" target="_blank">ISO_3166-2:CN</a>，广东省：“CN-GD”，<b>用作特征</b>
	Province *string `json:"Province,omitempty" name:"Province"`

	// 城市地区，统一用国家最新标准地区行政编码，如：<a href="https://www.mca.gov.cn/article/sj/xzqh/2020/" target="_blank">2020年行政区编码</a>，其他国家统一用国际公认城市简称或者城市编码，<b>用作特征</b>
	City *string `json:"City,omitempty" name:"City"`

	// 作者粉丝数，<b>用作特征</b>
	AuthorFans *int64 `json:"AuthorFans,omitempty" name:"AuthorFans"`

	// 作者评级，<b>用作特征</b>
	AuthorLevel *string `json:"AuthorLevel,omitempty" name:"AuthorLevel"`

	// 内容累计收藏次数，<b>用作特征</b>
	CollectCnt *int64 `json:"CollectCnt,omitempty" name:"CollectCnt"`

	// 内容累积点赞次数，<b>用作特征</b>
	PraiseCnt *int64 `json:"PraiseCnt,omitempty" name:"PraiseCnt"`

	// 内容累计评论次数，<b>用作特征</b>
	CommentCnt *int64 `json:"CommentCnt,omitempty" name:"CommentCnt"`

	// 内容累计分享次数，<b>用作特征</b>
	ShareCnt *int64 `json:"ShareCnt,omitempty" name:"ShareCnt"`

	// 内容累积打赏数，<b>用作特征</b>
	RewardCnt *int64 `json:"RewardCnt,omitempty" name:"RewardCnt"`

	// 内容质量评分，<b>用作特征</b>
	Score *float64 `json:"Score,omitempty" name:"Score"`

	// json字符串，<b>用于物料池管理的自定义扩展</b>，需要base64加密
	Extension *string `json:"Extension,omitempty" name:"Extension"`
}

type FeedBehaviorInfo struct {
	// 用户唯一ID，客户自定义用户ID，作为一个用户的唯一标识
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 内容唯一id
	ItemId *string `json:"ItemId,omitempty" name:"ItemId"`

	// 行为类型：<br> ● expose - 曝光，<b>必须</b><br> ● click - 点击，<b>必须</b><br/>  ● stay - 详情页停留时长，<b>强烈建议</b><br/>  ● videoover - 视频播放时长，<b>强烈建议</b><br/> ●  like - 点赞&喜欢，<b>正效果</b><br/> ● collect - 收藏，<b>正效果</b><br/> ●  share - 转发&分享，<b>正效果</b><br/> ● reward - 打赏，<b>正效果</b><br/> ● unlike - 踩&不喜欢，<b>负效果</b><br/> ●  comment - 评论<br/> 不支持的行为类型，可以映射到未被使用的其他行为类型。如实际业务数据中有私信行为，没有收藏行为，可以将私信行为映射到收藏行为
	BehaviorType *string `json:"BehaviorType,omitempty" name:"BehaviorType"`

	// 行为类型对应的行为值：<br/> ● expose - 曝光，固定填1<br/> ● click - 点击，固定填1<br/>  ● stay - 详情页停留时长，填停留秒数，取值[1-86400]<br/>  ● videoover - 视频播放时长，填播放结束的秒数，取值[1-86400]<br/> ●  like - 点赞&喜欢，固定填1<br/> ● collect - 收藏，固定填1<br/> ●  share - 转发&分享，固定填1<br/> ● reward - 打赏，填打赏金额，没有则填1<br/> ● unlike - 踩&不喜欢，填不喜欢的原因，没有则填1<br/> ●  comment - 评论，填评论内容，如“上海加油”
	BehaviorValue *string `json:"BehaviorValue,omitempty" name:"BehaviorValue"`

	// 行为发生的时间戳： 秒级时间戳，尽量实时上报，最长不超过半小时否则会影响推荐结果的准确性
	BehaviorTimestamp *int64 `json:"BehaviorTimestamp,omitempty" name:"BehaviorTimestamp"`

	// 行为发生的场景ID，在控制台创建场景后获取
	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`

	// 推荐追踪ID，使用推荐结果中返回的ItemTraceId填入。 
	// 注意：如果和推荐结果中的ItemTraceId不同，会影响行为特征归因，影响推荐算法效果
	ItemTraceId *string `json:"ItemTraceId,omitempty" name:"ItemTraceId"`

	// 内容类型，跟内容上报类型一致，用于效果分析，不做内容校验，<b>强烈建议</b>
	ItemType *string `json:"ItemType,omitempty" name:"ItemType"`

	// 相关推荐场景点击进入详情页的内容id，该字段用来注明行为发生于哪个内容的详情页推荐中，<b>相关推荐场景强烈建议</b>
	ReferrerItemId *string `json:"ReferrerItemId,omitempty" name:"ReferrerItemId"`

	// 用户设备ID数组，可传入用户的多个类型ID，详见UserIdInfo结构体，建议补齐，<b>用于构建用户画像信息</b>
	UserIdList []*UserIdInfo `json:"UserIdList,omitempty" name:"UserIdList"`

	// 算法来源： <br>● business 业务自己的算法对照组<br/> ● tencent 腾讯算法<br/> ● other 其他算法<br/>默认为tencent，区分行为来源于哪个算法，<b>用于Poc阶段的效果对比验证</b>
	Source *string `json:"Source,omitempty" name:"Source"`

	// 行为发生时的国家，ISO 3166-1 alpha-2编码，参考<a href="https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2" target="_blank">ISO 3166-1 alpha-2</a>，中国：“CN”，<b>用作特征</b>
	Country *string `json:"Country,omitempty" name:"Country"`

	// 行为发生时的省份，ISO 3166-2行政区编码，如中国参考<a href="https://zh.wikipedia.org/wiki/ISO_3166-2:CN" target="_blank">ISO_3166-2:CN</a>，广东省：“CN-GD”，<b>用作特征</b>
	Province *string `json:"Province,omitempty" name:"Province"`

	// 行为发生时的城市地区，统一用国家最新标准地区行政编码，如：<a href="https://www.mca.gov.cn/article/sj/xzqh/2020/" target="_blank">2020年行政区编码</a>，其他国家统一用国际公认城市简称或者城市编码，<b>用作特征</b>
	City *string `json:"City,omitempty" name:"City"`

	// 行为发生时的客户端ip，<b>用作特征</b>
	IP *string `json:"IP,omitempty" name:"IP"`

	// 行为发生时的客户端网络类型，<b>用作特征</b>
	Network *string `json:"Network,omitempty" name:"Network"`

	// 行为发生时的客户端平台，ios/android/h5，<b>用作特征</b>
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 行为发生时的客户端app版本，<b>用作特征</b>
	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`

	// 行为发生时的操作系统版本，<b>用作特征</b>
	OsVersion *string `json:"OsVersion,omitempty" name:"OsVersion"`

	// 行为发生时的机型，<b>用作特征</b>
	DeviceModel *string `json:"DeviceModel,omitempty" name:"DeviceModel"`

	// json字符串，<b>用于行为数据的扩展</b>，需要base64加密
	Extension *string `json:"Extension,omitempty" name:"Extension"`
}

// Predefined struct for user
type FeedRecommendRequestParams struct {
	// 实例ID，在控制台获取
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 场景ID，在控制台创建场景后获取
	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`

	// 用户唯一ID，客户自定义用户ID，作为一个用户的唯一标识
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户设备ID数组，可传入用户的多个类型ID，用于关联画像信息
	UserIdList []*UserIdInfo `json:"UserIdList,omitempty" name:"UserIdList"`

	// 推荐返回数量，默认10个，最多支持50个的内容返回。如果有更多数量要求，<a href="https://console.cloud.tencent.com/workorder/category" target="_blank">提单</a>沟通解决
	ItemCnt *int64 `json:"ItemCnt,omitempty" name:"ItemCnt"`

	// 当场景是相关推荐时该值必填，场景是非相关推荐时该值无效
	CurrentItemId *string `json:"CurrentItemId,omitempty" name:"CurrentItemId"`
}

type FeedRecommendRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，在控制台获取
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 场景ID，在控制台创建场景后获取
	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`

	// 用户唯一ID，客户自定义用户ID，作为一个用户的唯一标识
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户设备ID数组，可传入用户的多个类型ID，用于关联画像信息
	UserIdList []*UserIdInfo `json:"UserIdList,omitempty" name:"UserIdList"`

	// 推荐返回数量，默认10个，最多支持50个的内容返回。如果有更多数量要求，<a href="https://console.cloud.tencent.com/workorder/category" target="_blank">提单</a>沟通解决
	ItemCnt *int64 `json:"ItemCnt,omitempty" name:"ItemCnt"`

	// 当场景是相关推荐时该值必填，场景是非相关推荐时该值无效
	CurrentItemId *string `json:"CurrentItemId,omitempty" name:"CurrentItemId"`
}

func (r *FeedRecommendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FeedRecommendRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SceneId")
	delete(f, "UserId")
	delete(f, "UserIdList")
	delete(f, "ItemCnt")
	delete(f, "CurrentItemId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FeedRecommendRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FeedRecommendResponseParams struct {
	// 推荐返回的内容信息列表
	DataList []*RecItemData `json:"DataList,omitempty" name:"DataList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type FeedRecommendResponse struct {
	*tchttp.BaseResponse
	Response *FeedRecommendResponseParams `json:"Response"`
}

func (r *FeedRecommendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FeedRecommendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FeedUserInfo struct {
	// 用户唯一ID，客户自定义用户ID，作为一个用户的唯一标识
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户设备ID数组，可传入用户的多个类型ID，详见UserIdInfo结构体，建议补齐，<b>用于构建用户画像信息</b>
	UserIdList []*UserIdInfo `json:"UserIdList,omitempty" name:"UserIdList"`

	// 用户标签，多个标签用英文冒号联接，<b>用作特征，强烈建议</b>
	Tags *string `json:"Tags,omitempty" name:"Tags"`

	// 过滤列表，<b>会在推荐结果里过滤掉这类内容</b>
	DislikeInfoList []*DislikeInfo `json:"DislikeInfoList,omitempty" name:"DislikeInfoList"`

	// 用户年龄
	Age *int64 `json:"Age,omitempty" name:"Age"`

	// 用户性别： 0 - 未知 1 - 男 2 - 女
	Gender *int64 `json:"Gender,omitempty" name:"Gender"`

	// 用户学历 ：小学，初中，高中，大专，本科，硕士，博士
	Degree *string `json:"Degree,omitempty" name:"Degree"`

	// 用户毕业学校全称
	School *string `json:"School,omitempty" name:"School"`

	// 用户职业
	Occupation *string `json:"Occupation,omitempty" name:"Occupation"`

	// 用户所属行业
	Industry *string `json:"Industry,omitempty" name:"Industry"`

	// 用户常驻国家，ISO 3166-1 alpha-2编码，参考<a href="https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2" target="_blank">ISO 3166-1 alpha-2</a>，中国：“CN”
	ResidentCountry *string `json:"ResidentCountry,omitempty" name:"ResidentCountry"`

	// 用户常驻省份，ISO 3166-2行政区编码，如中国参考<a href="https://zh.wikipedia.org/wiki/ISO_3166-2:CN" target="_blank">ISO_3166-2:CN</a>，广东省：“CN-GD”
	ResidentProvince *string `json:"ResidentProvince,omitempty" name:"ResidentProvince"`

	// 用户常驻城市，统一用国家最新标准地区行政编码，如：<a href="https://www.mca.gov.cn/article/sj/xzqh/2020/" target="_blank">2020年行政区编码</a>，
	ResidentCity *string `json:"ResidentCity,omitempty" name:"ResidentCity"`

	// 用户注册时间，秒级时间戳（1639624786）
	RegisterTimestamp *int64 `json:"RegisterTimestamp,omitempty" name:"RegisterTimestamp"`

	// 用户会员等级
	MembershipLevel *string `json:"MembershipLevel,omitempty" name:"MembershipLevel"`

	// 用户上一次登录时间，秒级时间戳（1639624786）
	LastLoginTimestamp *int64 `json:"LastLoginTimestamp,omitempty" name:"LastLoginTimestamp"`

	// 用户上一次登录的ip
	LastLoginIp *string `json:"LastLoginIp,omitempty" name:"LastLoginIp"`

	// 用户信息的最后修改时间戳，秒级时间戳（1639624786）
	LastModifyTimestamp *int64 `json:"LastModifyTimestamp,omitempty" name:"LastModifyTimestamp"`

	// json字符串，用于画像数据的扩展，需要base64加密
	Extension *string `json:"Extension,omitempty" name:"Extension"`
}

type GoodsBehaviorInfo struct {
	// 用户唯一ID，客户自定义用户ID，作为一个用户的唯一标识
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 商品唯一ID，skuId或spuId，客户根据需求自行决定商品主键粒度
	GoodsId *string `json:"GoodsId,omitempty" name:"GoodsId"`

	// 行为类型：<br> ● expose - 曝光，<b>必须</b><br> ● click - 点击，<b>必须</b><br/>  ● stay - 详情页停留时长，<b>强烈建议</b><br/>  ● videoover - 视频播放时长，<b>强烈建议</b><br/> ●  like - 点赞&喜欢，<b>正效果</b><br/> ● collect - 收藏，<b>正效果</b><br/> ●  share - 转发&分享，<b>正效果</b><br/> ● reward - 打赏，<b>正效果</b><br/> ● unlike - 踩&不喜欢，<b>负效果</b><br/> ●  comment - 评论<br/> ●  order - 下单<br/> ●  buy - 购买成功<br/> ●  addcart - 加入购物车<br/>   
	// 不支持的行为类型，可以映射到未被使用的其他行为类型。如实际业务数据中有私信行为，没有收藏行为，可以将私信行为映射到收藏行为
	BehaviorType *string `json:"BehaviorType,omitempty" name:"BehaviorType"`

	// 行为类型对应的行为值：<br/> ● expose - 曝光，固定填1<br/> ● click - 点击，固定填1<br/>  ● stay - 详情页停留时长，填停留秒数，取值[1-86400]<br/>  ● videoover - 视频播放时长，填播放结束的秒数，取值[1-86400]<br/> ●  like - 点赞&喜欢，固定填1<br/> ● collect - 收藏，固定填1<br/> ●  share - 转发&分享，固定填1<br/> ● reward - 打赏，填打赏金额，没有则填1<br/> ● unlike - 踩&不喜欢，填不喜欢的原因，没有则填1<br/> ●  comment - 评论，填评论内容，如“上海加油”<br/> ●  order - 下单，固定填1<br/> ●  buy - 购买成功，固定填1<br/> ●  addcart - 加入购物车，固定填1
	BehaviorValue *string `json:"BehaviorValue,omitempty" name:"BehaviorValue"`

	// 行为发生的时间戳： 秒级时间戳，尽量实时上报，最长不超过半小时否则会影响推荐结果的准确性
	BehaviorTimestamp *int64 `json:"BehaviorTimestamp,omitempty" name:"BehaviorTimestamp"`

	// 行为发生的场景ID，在控制台创建场景后获取
	SceneId *string `json:"SceneId,omitempty" name:"SceneId"`

	// 算法来源： <br>● business 业务自己的算法对照组<br/> ● tencent 腾讯算法<br/> ● other 其他算法<br/>默认为tencent，区分行为来源于哪个算法，<b>用于Poc阶段的效果对比验证</b>
	Source *string `json:"Source,omitempty" name:"Source"`

	// 标识行为发生在app内哪个页面，取值客户自定，可以是明文或id，建议传明文便于理解、分析，如首页，发现页，用户中心等
	// <b>用作上下文特征，刻画不同场景用户行为分布的差异</b>
	Page *string `json:"Page,omitempty" name:"Page"`

	// 标识行为发生在页面的哪一区块，取值客户自定，可以是明文或id，建议传明文便于理解、分析，如横幅、广告位、猜你喜欢等
	// <b>用作上下文特征，刻画不同模块用户行为分布的差异</b>
	Module *string `json:"Module,omitempty" name:"Module"`

	// 推荐追踪ID，使用推荐结果中返回的GoodsTraceId填入。 
	// 注意：如果和推荐结果中的GoodsTraceId不同，会影响行为特征归因，影响推荐算法效果。<b>强烈建议</b>
	GoodsTraceId *string `json:"GoodsTraceId,omitempty" name:"GoodsTraceId"`

	// 相关推荐场景点击进入详情页的内容id，该字段用来注明行为发生于哪个内容的详情页推荐中，<b>相关推荐场景强烈建议</b>
	ReferrerGoodsId *string `json:"ReferrerGoodsId,omitempty" name:"ReferrerGoodsId"`

	// 订单商品购买个数，当behaviorType=order，buy或addcart时有值，<b>用作特征</b>
	OrderGoodsCnt *int64 `json:"OrderGoodsCnt,omitempty" name:"OrderGoodsCnt"`

	// 订单总金额，当behaviorType=order或buy时有值（单位：元，统一货币体系，如统一为RMB，美元等），<b>用作特征</b>
	OrderAmount *float64 `json:"OrderAmount,omitempty" name:"OrderAmount"`

	// 用户设备ID数组，可传入用户的多个类型ID，详见UserIdInfo结构体，建议补齐，<b>用于构建用户画像信息</b>
	UserIdList []*StrUserIdInfo `json:"UserIdList,omitempty" name:"UserIdList"`

	// 行为发生时用户基础特征信息，<b>用作特征</b>
	UserPortraitInfo *UserPortraitInfo `json:"UserPortraitInfo,omitempty" name:"UserPortraitInfo"`

	// 标识行为发生在模块内的具体位置，如1、2、...
	// <b>用作上下文特征，刻画不同位置用户行为分布的差异</b>
	Position *int64 `json:"Position,omitempty" name:"Position"`

	// json字符串，<b>用于行为数据的扩展</b>
	Extension *string `json:"Extension,omitempty" name:"Extension"`
}

type GoodsInfo struct {
	// 商品唯一ID，skuId或spuId，客户根据需求自行决定商品主键粒度。建议限制在128字符以内
	GoodsId *string `json:"GoodsId,omitempty" name:"GoodsId"`

	// 商品物料展示类型：<br/>● article -图文<br>● text -纯文本<br/>● video -视频<br/>● short_video -时长15秒以内的视频<br/>● mini_video -竖屏视频<br/>● image -纯图片<br/>（如当前类型不满足，请<a href="https://console.cloud.tencent.com/workorder/category" target="_blank">提单</a>沟通解决方案）
	GoodsType *string `json:"GoodsType,omitempty" name:"GoodsType"`

	// 商品状态：
	// ● 1 - 上架 
	// ● 2 - 下架 
	// Status=2的内容不会在推荐结果中出现 
	// 需要下架内容时，把Status的值修改为2即可
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 商品生成时间，秒级时间戳（1639624786），需大于0，<b>用作特征和物料管理</b>
	PublishTimestamp *int64 `json:"PublishTimestamp,omitempty" name:"PublishTimestamp"`

	// 商品过期时间，秒级时间戳（1639624786），如未填，则默认PublishTimestamp往后延一年，<b>用作特征</b>，过期则不会被推荐，<b>强烈建议</b>
	ExpireTimestamp *int64 `json:"ExpireTimestamp,omitempty" name:"ExpireTimestamp"`

	// spu((Standard Product Unit))维度id，商品聚合信息的最小单位，<b>强烈建议</b>
	SpuId *string `json:"SpuId,omitempty" name:"SpuId"`

	// 类目层级数，例如3级类目，则填3，和CategoryPath字段的类数据匹配，<b>强烈建议</b>
	CategoryLevel *int64 `json:"CategoryLevel,omitempty" name:"CategoryLevel"`

	// 类目路径，一级二级三级等依次用英文冒号联接，和CategoryLevel字段值匹配，如体育：“女装:裙子:半身裙”。<b>用于物料池管理，强烈建议</b>
	CategoryPath *string `json:"CategoryPath,omitempty" name:"CategoryPath"`

	// 商品标题，<b>主要用于语义分析</b>，<b>强烈建议</b>
	Title *string `json:"Title,omitempty" name:"Title"`

	// 商品标签，多个标签用英文冒号联接，<b>用作特征，强烈建议</b>
	Tags *string `json:"Tags,omitempty" name:"Tags"`

	// 商品对应的品牌，取值用户自定义，可以是品牌id或品牌明文，<b>用作特征以及打散/过滤规则，强烈建议</b>
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// 商品所属店铺ID，取值客户自定义，<b>用作特征，强烈建议</b>
	ShopId *string `json:"ShopId,omitempty" name:"ShopId"`

	// 商品原始价格（单位：元，统一货币体系，如统一为RMB或美元等），<b>用作特征，强烈建议</b>
	OrgPrice *float64 `json:"OrgPrice,omitempty" name:"OrgPrice"`

	// 商品当前价格（单位：元，统一货币体系，如统一为RMB或美元等），<b>用作特征，强烈建议</b>
	CurPrice *float64 `json:"CurPrice,omitempty" name:"CurPrice"`

	// 商品来源类型，客户自定义，<b>用于物料池管理</b>
	SourceId *string `json:"SourceId,omitempty" name:"SourceId"`

	// 商品正文关键片段，建议控制在500字符以内，<b>主要用于语义分析</b>
	Content *string `json:"Content,omitempty" name:"Content"`

	// 商品正文详情，主要用于语义分析，当内容过大时建议用ContentUrl传递，<b>与Content可二选一</b>
	ContentUrl *string `json:"ContentUrl,omitempty" name:"ContentUrl"`

	// 商品封面url，不超过10个，<b>用作特征</b>
	PicUrlList []*string `json:"PicUrlList,omitempty" name:"PicUrlList"`

	// 卖家所在国家，ISO 3166-1 alpha-2编码，参考<a href="https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2" target="_blank">ISO 3166-1 alpha-2</a>，中国：“CN”，<b>用作特征</b>
	Country *string `json:"Country,omitempty" name:"Country"`

	// 卖家所在省份，ISO 3166-2行政区编码，如中国参考<a href="https://zh.wikipedia.org/wiki/ISO_3166-2:CN" target="_blank">ISO_3166-2:CN</a>，广东省：“CN-GD”，<b>用作特征</b>
	Province *string `json:"Province,omitempty" name:"Province"`

	// 卖家所在城市地区，统一用国家最新标准地区行政编码，如：<a href="https://www.mca.gov.cn/article/sj/xzqh/2020/" target="_blank">2020年行政区编码</a>，其他国家统一用国际公认城市简称或者城市编码，<b>用作特征</b>
	City *string `json:"City,omitempty" name:"City"`

	// 商品是否包邮；1:包邮；2:不包邮；3:满足条件包邮，<b>用作特征</b>
	FreeShipping *int64 `json:"FreeShipping,omitempty" name:"FreeShipping"`

	// 商品邮费（单位：元，统一货币体系，如统一为RMB或美元等），<b>用作特征</b>
	ShippingPrice *float64 `json:"ShippingPrice,omitempty" name:"ShippingPrice"`

	// 商品累计好评次数，<b>用作特征</b>
	PraiseCnt *int64 `json:"PraiseCnt,omitempty" name:"PraiseCnt"`

	// 商品累计评论次数，<b>用作特征</b>
	CommentCnt *int64 `json:"CommentCnt,omitempty" name:"CommentCnt"`

	// 商品累计分享次数，<b>用作特征</b>
	ShareCnt *int64 `json:"ShareCnt,omitempty" name:"ShareCnt"`

	// 商品累计收藏次数，<b>用作特征</b>
	CollectCnt *int64 `json:"CollectCnt,omitempty" name:"CollectCnt"`

	// 商品累积成交次数，<b>用作特征</b>
	OrderCnt *int64 `json:"OrderCnt,omitempty" name:"OrderCnt"`

	// 商品平均客户评分，取值范围用户自定，<b>用作特征</b>
	Score *float64 `json:"Score,omitempty" name:"Score"`

	// json字符串，<b>用于物料池管理的自定义扩展</b>
	Extension *string `json:"Extension,omitempty" name:"Extension"`
}

type RecGoodsData struct {
	// 推荐返回的商品ID
	GoodsId *string `json:"GoodsId,omitempty" name:"GoodsId"`

	// 推荐结果分，取值范围[0,1000000]
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *float64 `json:"Score,omitempty" name:"Score"`

	// 推荐追踪id，本次推荐内容产生的后续行为上报均要用该GoodsTraceId上报。每次接口调用返回的GoodsTraceId不同
	// 注意：此字段可能返回 null，表示取不到有效值。
	GoodsTraceId *string `json:"GoodsTraceId,omitempty" name:"GoodsTraceId"`

	// 商品所在位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Position *int64 `json:"Position,omitempty" name:"Position"`
}

type RecItemData struct {
	// 推荐的内容ID
	ItemId *string `json:"ItemId,omitempty" name:"ItemId"`

	// 内容类型，同内容上报类型一致
	// 注意：此字段可能返回 null，表示取不到有效值。
	ItemType *string `json:"ItemType,omitempty" name:"ItemType"`

	// 推荐追踪id，本次推荐内容产生的后续行为上报均要用该ItemTraceId上报。每次接口调用返回的ItemTraceId不同
	// 注意：此字段可能返回 null，表示取不到有效值。
	ItemTraceId *string `json:"ItemTraceId,omitempty" name:"ItemTraceId"`

	// 推荐预测分，分值越高被推荐的理由越充分，取值范围[0,1000000]，用于做二次排序的参考
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *float64 `json:"Score,omitempty" name:"Score"`
}

// Predefined struct for user
type ReportFeedBehaviorRequestParams struct {
	// 实例ID，在控制台获取
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 上报的行为数据数组，数量不超过50
	FeedBehaviorList []*FeedBehaviorInfo `json:"FeedBehaviorList,omitempty" name:"FeedBehaviorList"`
}

type ReportFeedBehaviorRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，在控制台获取
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 上报的行为数据数组，数量不超过50
	FeedBehaviorList []*FeedBehaviorInfo `json:"FeedBehaviorList,omitempty" name:"FeedBehaviorList"`
}

func (r *ReportFeedBehaviorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReportFeedBehaviorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "FeedBehaviorList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReportFeedBehaviorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReportFeedBehaviorResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ReportFeedBehaviorResponse struct {
	*tchttp.BaseResponse
	Response *ReportFeedBehaviorResponseParams `json:"Response"`
}

func (r *ReportFeedBehaviorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReportFeedBehaviorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReportFeedItemRequestParams struct {
	// 实例ID，在控制台获取
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 上报的信息流内容数组，一次数量不超过50
	FeedItemList []*DocItem `json:"FeedItemList,omitempty" name:"FeedItemList"`
}

type ReportFeedItemRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，在控制台获取
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 上报的信息流内容数组，一次数量不超过50
	FeedItemList []*DocItem `json:"FeedItemList,omitempty" name:"FeedItemList"`
}

func (r *ReportFeedItemRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReportFeedItemRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "FeedItemList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReportFeedItemRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReportFeedItemResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ReportFeedItemResponse struct {
	*tchttp.BaseResponse
	Response *ReportFeedItemResponseParams `json:"Response"`
}

func (r *ReportFeedItemResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReportFeedItemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReportFeedUserRequestParams struct {
	// 实例ID，在控制台获取
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 上报的用户信息数组，数量不超过50
	FeedUserList []*FeedUserInfo `json:"FeedUserList,omitempty" name:"FeedUserList"`
}

type ReportFeedUserRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，在控制台获取
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 上报的用户信息数组，数量不超过50
	FeedUserList []*FeedUserInfo `json:"FeedUserList,omitempty" name:"FeedUserList"`
}

func (r *ReportFeedUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReportFeedUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "FeedUserList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReportFeedUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReportFeedUserResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ReportFeedUserResponse struct {
	*tchttp.BaseResponse
	Response *ReportFeedUserResponseParams `json:"Response"`
}

func (r *ReportFeedUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReportFeedUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReportGoodsBehaviorRequestParams struct {
	// 实例ID，在控制台获取
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 上报的商品对应的用户行为数据数组，数量不超过50
	GoodsBehaviorList []*GoodsBehaviorInfo `json:"GoodsBehaviorList,omitempty" name:"GoodsBehaviorList"`
}

type ReportGoodsBehaviorRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，在控制台获取
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 上报的商品对应的用户行为数据数组，数量不超过50
	GoodsBehaviorList []*GoodsBehaviorInfo `json:"GoodsBehaviorList,omitempty" name:"GoodsBehaviorList"`
}

func (r *ReportGoodsBehaviorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReportGoodsBehaviorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "GoodsBehaviorList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReportGoodsBehaviorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReportGoodsBehaviorResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ReportGoodsBehaviorResponse struct {
	*tchttp.BaseResponse
	Response *ReportGoodsBehaviorResponseParams `json:"Response"`
}

func (r *ReportGoodsBehaviorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReportGoodsBehaviorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReportGoodsInfoRequestParams struct {
	// 实例ID，在控制台获取
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 上报的商品数组，一次数量不超过50
	GoodsList []*GoodsInfo `json:"GoodsList,omitempty" name:"GoodsList"`
}

type ReportGoodsInfoRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，在控制台获取
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 上报的商品数组，一次数量不超过50
	GoodsList []*GoodsInfo `json:"GoodsList,omitempty" name:"GoodsList"`
}

func (r *ReportGoodsInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReportGoodsInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "GoodsList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReportGoodsInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReportGoodsInfoResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ReportGoodsInfoResponse struct {
	*tchttp.BaseResponse
	Response *ReportGoodsInfoResponseParams `json:"Response"`
}

func (r *ReportGoodsInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReportGoodsInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StrUserIdInfo struct {

}

type UserIdInfo struct {
	// 用户ID类型： <br/>● qq: qq号码 <br/>● qq_md5：qq的md5值 <br/>● imei：设备imei <br/>● imei_md5：imei的md5值 <br/>● idfa: Apple 向用户设备随机分配的设备标识符 <br/>● idfa_md5：idfa的md5值 <br/>● oaid：安卓10之后一种非永久性设备标识符 <br/>● oaid_md5：md5后的oaid <br/>● wx_openid：微信openid <br/>● qq_openid：QQ的openid <br/>● phone：电话号码 <br/>● phone_md5：md5后的电话号码 <br/>● phone_sha256：SHA256加密的手机号 <br/>● phone_sm3：国密SM3加密的手机号 <br/>（如当前类型不满足，请<a href="https://console.cloud.tencent.com/workorder/category" target="_blank">提单</a>沟通解决方案）
	Type *string `json:"Type,omitempty" name:"Type"`

	// 用户ID值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type UserPortraitInfo struct {

}