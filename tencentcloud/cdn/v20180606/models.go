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

package v20180606

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AddCdnDomainRequest struct {
	*tchttp.BaseRequest

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 加速域名业务类型
	// web：静态加速
	// download：下载加速
	// media：流媒体点播加速
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 源站配置
	Origin *Origin `json:"Origin,omitempty" name:"Origin"`

	// 项目 ID，默认为 0，代表【默认项目】
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// IP 黑白名单配置
	IpFilter *IpFilter `json:"IpFilter,omitempty" name:"IpFilter"`

	// IP 限频配置
	IpFreqLimit *IpFreqLimit `json:"IpFreqLimit,omitempty" name:"IpFreqLimit"`

	// 状态码缓存配置
	StatusCodeCache *StatusCodeCache `json:"StatusCodeCache,omitempty" name:"StatusCodeCache"`

	// 智能压缩配置
	Compression *Compression `json:"Compression,omitempty" name:"Compression"`

	// 带宽封顶配置
	BandwidthAlert *BandwidthAlert `json:"BandwidthAlert,omitempty" name:"BandwidthAlert"`

	// Range 回源配置
	RangeOriginPull *RangeOriginPull `json:"RangeOriginPull,omitempty" name:"RangeOriginPull"`

	// 301/302 回源跟随配置。
	FollowRedirect *FollowRedirect `json:"FollowRedirect,omitempty" name:"FollowRedirect"`

	// 错误码重定向配置（功能灰度中，尚未全量）
	ErrorPage *ErrorPage `json:"ErrorPage,omitempty" name:"ErrorPage"`

	// 请求头部配置
	RequestHeader *RequestHeader `json:"RequestHeader,omitempty" name:"RequestHeader"`

	// 响应头部配置
	ResponseHeader *ResponseHeader `json:"ResponseHeader,omitempty" name:"ResponseHeader"`

	// 下载速度配置
	DownstreamCapping *DownstreamCapping `json:"DownstreamCapping,omitempty" name:"DownstreamCapping"`

	// 节点缓存键配置
	CacheKey *CacheKey `json:"CacheKey,omitempty" name:"CacheKey"`

	// 头部缓存配置
	ResponseHeaderCache *ResponseHeaderCache `json:"ResponseHeaderCache,omitempty" name:"ResponseHeaderCache"`

	// 视频拖拽配置
	VideoSeek *VideoSeek `json:"VideoSeek,omitempty" name:"VideoSeek"`

	// 缓存过期时间配置
	Cache *Cache `json:"Cache,omitempty" name:"Cache"`

	// 跨国链路优化配置
	OriginPullOptimization *OriginPullOptimization `json:"OriginPullOptimization,omitempty" name:"OriginPullOptimization"`

	// Https 加速配置
	Https *Https `json:"Https,omitempty" name:"Https"`

	// 时间戳防盗链配置
	Authentication *Authentication `json:"Authentication,omitempty" name:"Authentication"`

	// SEO 优化配置
	Seo *Seo `json:"Seo,omitempty" name:"Seo"`

	// 访问协议强制跳转配置
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitempty" name:"ForceRedirect"`

	// Referer 防盗链配置
	Referer *Referer `json:"Referer,omitempty" name:"Referer"`

	// 浏览器缓存配置（功能灰度中，尚未全量）
	MaxAge *MaxAge `json:"MaxAge,omitempty" name:"MaxAge"`

	// Ipv6 配置（功能灰度中，尚未全量）
	Ipv6 *Ipv6 `json:"Ipv6,omitempty" name:"Ipv6"`

	// 地域属性特殊配置
	// 适用于域名境内加速、境外加速配置不一致场景
	SpecificConfig *SpecificConfig `json:"SpecificConfig,omitempty" name:"SpecificConfig"`

	// 域名加速区域
	// mainland：中国境内加速
	// overseas：中国境外加速
	// global：全球加速
	// 使用中国境外加速、全球加速时，需要先开通中国境外加速服务
	Area *string `json:"Area,omitempty" name:"Area"`

	// 回源超时配置
	OriginPullTimeout *OriginPullTimeout `json:"OriginPullTimeout,omitempty" name:"OriginPullTimeout"`
}

func (r *AddCdnDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddCdnDomainRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddCdnDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddCdnDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddCdnDomainResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AdvanceCacheRule struct {

	// 规则类型：
	// all：所有文件生效
	// file：指定文件后缀生效
	// directory：指定路径生效
	// path：指定绝对路径生效
	// default：源站未返回 max-age 情况下的缓存规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheType *string `json:"CacheType,omitempty" name:"CacheType"`

	// 对应类型下的匹配内容：
	// all 时填充 *
	// file 时填充后缀名，如 jpg、txt
	// directory 时填充路径，如 /xxx/test/
	// path 时填充绝对路径，如 /xxx/test.html
	// default 时填充 "no max-age"
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheContents []*string `json:"CacheContents,omitempty" name:"CacheContents" list`

	// 缓存过期时间
	// 单位为秒，最大可设置为 365 天
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheTime *int64 `json:"CacheTime,omitempty" name:"CacheTime"`
}

type AdvancedCache struct {

	// 缓存过期规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheRules []*AdvanceCacheRule `json:"CacheRules,omitempty" name:"CacheRules" list`

	// 强制缓存配置
	// on：开启
	// off：关闭
	// 开启时，源站返回 no-cache、no-store 头部时，仍按照缓存过期规则进行节点缓存
	// 默认为关闭状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	IgnoreCacheControl *string `json:"IgnoreCacheControl,omitempty" name:"IgnoreCacheControl"`

	// 忽略源站的 Set-Cookie 头部
	// on：开启
	// off：关闭
	// 默认为关闭状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	IgnoreSetCookie *string `json:"IgnoreSetCookie,omitempty" name:"IgnoreSetCookie"`
}

type Authentication struct {

	// 防盗链配置开关
	// on：开启
	// off：关闭
	// 开启时必须且只配置一种模式，其余模式需要设置为 null
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 时间戳防盗链模式 A 配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	TypeA *AuthenticationTypeA `json:"TypeA,omitempty" name:"TypeA"`

	// 时间戳防盗链模式 B 配置（模式 B 后台升级中，暂时不支持配置）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TypeB *AuthenticationTypeB `json:"TypeB,omitempty" name:"TypeB"`

	// 时间戳防盗链模式 C 配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	TypeC *AuthenticationTypeC `json:"TypeC,omitempty" name:"TypeC"`

	// 时间戳防盗链模式 D 配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	TypeD *AuthenticationTypeD `json:"TypeD,omitempty" name:"TypeD"`
}

type AuthenticationTypeA struct {

	// 计算签名的密钥
	// 仅允许大小写字母与数字，长度 6~32 位
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`

	// 签名参数名设置
	// 仅允许大小写字母、数字或下划线，长度 1~100 位，不能以数字开头
	SignParam *string `json:"SignParam,omitempty" name:"SignParam"`

	// 签名过期时间设置
	// 单位为秒，最大可设置为 31536000
	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 鉴权/不做鉴权的文件扩展名列表设置
	// 如果包含字符 *  则表示所有文件
	FileExtensions []*string `json:"FileExtensions,omitempty" name:"FileExtensions" list`

	// whitelist：白名单，表示对除了 FileExtensions 列表之外的所有类型进行鉴权
	// blacklist：黑名单，表示仅对 FileExtensions 中的类型进行鉴权
	FilterType *string `json:"FilterType,omitempty" name:"FilterType"`
}

type AuthenticationTypeB struct {

	// 计算签名的密钥
	// 仅允许大小写字母与数字，长度 6~32 位
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`

	// 签名过期时间设置
	// 单位为秒，最大可设置为 31536000
	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 鉴权/不做鉴权的文件扩展名列表设置
	// 如果包含字符 *  则表示所有文件
	FileExtensions []*string `json:"FileExtensions,omitempty" name:"FileExtensions" list`

	// whitelist：白名单，表示对除了 FileExtensions 列表之外的所有类型进行鉴权
	// blacklist：黑名单，表示仅对 FileExtensions 中的类型进行鉴权
	FilterType *string `json:"FilterType,omitempty" name:"FilterType"`
}

type AuthenticationTypeC struct {

	// 计算签名的密钥
	// 仅允许大小写字母与数字，长度 6~32 位
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`

	// 签名过期时间设置
	// 单位为秒，最大可设置为 31536000
	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 鉴权/不做鉴权的文件扩展名列表设置
	// 如果包含字符 *  则表示所有文件
	FileExtensions []*string `json:"FileExtensions,omitempty" name:"FileExtensions" list`

	// whitelist：白名单，表示对除了 FileExtensions 列表之外的所有类型进行鉴权
	// blacklist：黑名单，表示仅对 FileExtensions 中的类型进行鉴权
	FilterType *string `json:"FilterType,omitempty" name:"FilterType"`
}

type AuthenticationTypeD struct {

	// 计算签名的密钥
	// 仅允许大小写字母与数字，长度 6~32 位
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`

	// 签名过期时间设置
	// 单位为秒，最大可设置为 31536000
	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 鉴权/不做鉴权的文件扩展名列表设置
	// 如果包含字符 *  则表示所有文件
	FileExtensions []*string `json:"FileExtensions,omitempty" name:"FileExtensions" list`

	// whitelist：白名单，表示对除了 FileExtensions 列表之外的所有类型进行鉴权
	// blacklist：黑名单，表示仅对 FileExtensions 中的类型进行鉴权
	FilterType *string `json:"FilterType,omitempty" name:"FilterType"`

	// 签名参数名设置
	// 仅允许大小写字母、数字或下划线，长度 1~100 位，不能以数字开头
	SignParam *string `json:"SignParam,omitempty" name:"SignParam"`

	// 时间戳参数名设置
	// 仅允许大小写字母、数字或下划线，长度 1~100 位，不能以数字开头
	TimeParam *string `json:"TimeParam,omitempty" name:"TimeParam"`

	// 时间戳进制设置
	// dec：十进制
	// hex：十六进制
	TimeFormat *string `json:"TimeFormat,omitempty" name:"TimeFormat"`
}

type AwsPrivateAccess struct {

	// 开关，on/off。
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 访问ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessKey *string `json:"AccessKey,omitempty" name:"AccessKey"`

	// 密钥。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`
}

type BandwidthAlert struct {

	// 带宽封顶配置开关
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 带宽封顶阈值，单位为bps
	// 注意：此字段可能返回 null，表示取不到有效值。
	BpsThreshold *int64 `json:"BpsThreshold,omitempty" name:"BpsThreshold"`

	// 达到阈值后的操作
	// RESOLVE_DNS_TO_ORIGIN：直接回源，仅自有源站域名支持
	// RETURN_404：全部请求返回 404
	// 注意：此字段可能返回 null，表示取不到有效值。
	CounterMeasure *string `json:"CounterMeasure,omitempty" name:"CounterMeasure"`

	// 上次触发带宽封顶阈值的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastTriggerTime *string `json:"LastTriggerTime,omitempty" name:"LastTriggerTime"`
}

type BriefDomain struct {

	// 域名 ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 腾讯云账号 ID
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 加速域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名对应的 CNAME 地址
	Cname *string `json:"Cname,omitempty" name:"Cname"`

	// 加速服务状态
	// rejected：域名审核未通过，域名备案过期/被注销导致
	// processing：部署中
	// online：已启动
	// offline：已关闭
	Status *string `json:"Status,omitempty" name:"Status"`

	// 项目 ID，可前往腾讯云项目管理页面查看
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 域名业务类型
	// web：静态加速
	// download：下载加速
	// media：流媒体点播加速
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 域名创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 域名更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 源站配置详情
	Origin *Origin `json:"Origin,omitempty" name:"Origin"`

	// 域名封禁状态
	// normal：正常状态
	// overdue：账号欠费导致域名关闭，充值完成后可自行启动加速服务
	// malicious：域名出现恶意行为，强制关闭加速服务
	// ddos：域名被大规模 DDoS 攻击，关闭加速服务
	// idle：域名超过 90 天内无任何操作、数据产生，判定为不活跃域名自动关闭加速服务，可自行启动加速服务
	// unlicensed：域名未备案/备案注销，自动关闭加速服务，备案完成后可自行启动加速服务
	// capping：触发配置的带宽阈值上限
	// readonly：域名存在特殊配置，被锁定
	Disable *string `json:"Disable,omitempty" name:"Disable"`

	// 加速区域
	// mainland：中国境内加速
	// overseas：中国境外加速
	// global：全球加速
	Area *string `json:"Area,omitempty" name:"Area"`

	// 域名锁定状态
	// normal：未锁定
	// mainland：中国境内锁定
	// overseas：中国境外锁定
	// global：全球锁定
	Readonly *string `json:"Readonly,omitempty" name:"Readonly"`
}

type Cache struct {

	// 基础缓存过期时间配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	SimpleCache *SimpleCache `json:"SimpleCache,omitempty" name:"SimpleCache"`

	// 高级缓存过期时间配置（功能灰度中，尚未全量）
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdvancedCache *AdvancedCache `json:"AdvancedCache,omitempty" name:"AdvancedCache"`
}

type CacheKey struct {

	// 是否开启全路径缓存
	// on：开启全路径缓存（即关闭参数过滤）
	// off：关闭全路径缓存（即开启参数过滤）
	FullUrlCache *string `json:"FullUrlCache,omitempty" name:"FullUrlCache"`

	// 是否使用请求参数作为CacheKey的一部分
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueryString *QueryStringKey `json:"QueryString,omitempty" name:"QueryString"`

	// 是否使用请求头部作为CacheKey的一部分
	// 注意：此字段可能返回 null，表示取不到有效值。
	Header *HeaderKey `json:"Header,omitempty" name:"Header"`

	// 是否使用Cookie作为CacheKey的一部分
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cookie *CookieKey `json:"Cookie,omitempty" name:"Cookie"`

	// 是否使用请求协议作为CacheKey的一部分
	// 注意：此字段可能返回 null，表示取不到有效值。
	Scheme *SchemeKey `json:"Scheme,omitempty" name:"Scheme"`

	// 是否使用自定义字符串作为CacheKey的一部分
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheTag *CacheTagKey `json:"CacheTag,omitempty" name:"CacheTag"`
}

type CacheOptResult struct {

	// 成功的url列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessUrls []*string `json:"SuccessUrls,omitempty" name:"SuccessUrls" list`

	// 失败的url列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailUrls []*string `json:"FailUrls,omitempty" name:"FailUrls" list`
}

type CacheTagKey struct {

	// 是否使用CacheTag作为CacheKey的一部分
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 自定义CacheTag的值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type CappingRule struct {

	// 规则类型：
	// all：所有文件生效
	// file：指定文件后缀生效
	// directory：指定路径生效
	// path：指定绝对路径生效
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`

	// RuleType 对应类型下的匹配内容： 
	// all 时填充 *
	// file 时填充后缀名，如 jpg、txt
	// directory 时填充路径，如 /xxx/test/
	// path 时填充绝对路径，如 /xxx/test.html
	RulePaths []*string `json:"RulePaths,omitempty" name:"RulePaths" list`

	// 下行速度值设置，单位为 KB/s
	KBpsThreshold *int64 `json:"KBpsThreshold,omitempty" name:"KBpsThreshold"`
}

type CdnData struct {

	// 查询指定的指标名称：
	// flux：流量，单位为 byte
	// bandwidth：带宽，单位为 bps
	// request：请求数，单位为 次
	// fluxHitRate：流量命中率，单位为 %
	// statusCode：状态码，返回 2XX、3XX、4XX、5XX 汇总数据，单位为 个
	// 2XX：返回 2XX 状态码汇总及各 2 开头状态码数据，单位为 个
	// 3XX：返回 3XX 状态码汇总及各 3 开头状态码数据，单位为 个
	// 4XX：返回 4XX 状态码汇总及各 4 开头状态码数据，单位为 个
	// 5XX：返回 5XX 状态码汇总及各 5 开头状态码数据，单位为 个
	// 或指定查询的某一具体状态码
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// 明细数据组合
	DetailData []*TimestampData `json:"DetailData,omitempty" name:"DetailData" list`

	// 汇总数据组合
	SummarizedData *SummarizedData `json:"SummarizedData,omitempty" name:"SummarizedData"`
}

type CdnIp struct {

	// 指定查询的 IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// IP 归属：
	// yes：节点归属于腾讯云 CDN
	// no：节点不属于腾讯云 CDN
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 节点所处的省份/国家
	// unknown 表示节点位置未知
	Location *string `json:"Location,omitempty" name:"Location"`

	// 节点上下线历史记录
	History []*CdnIpHistory `json:"History,omitempty" name:"History" list`

	// 节点的所在区域
	// mainland：中国境内加速节点
	// overseas：中国境外加速节点
	// unknown：服务地域无法获取
	Area *string `json:"Area,omitempty" name:"Area"`
}

type CdnIpHistory struct {

	// 操作类型
	// online：节点上线
	// offline：节点下线
	Status *string `json:"Status,omitempty" name:"Status"`

	// 操作类型对应的操作时间
	// 当该值为 null 时表示无历史状态变更记录
	// 注意：此字段可能返回 null，表示取不到有效值。
	Datetime *string `json:"Datetime,omitempty" name:"Datetime"`
}

type ClientCert struct {

	// 客户端证书
	// PEM 格式，需要进行 Base 64 编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Certificate *string `json:"Certificate,omitempty" name:"Certificate"`

	// 客户端证书名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertName *string `json:"CertName,omitempty" name:"CertName"`

	// 证书过期时间
	// 作为入参时无需填充
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 证书颁发时间
	// 作为入参时无需填充
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployTime *string `json:"DeployTime,omitempty" name:"DeployTime"`
}

type ClsLogObject struct {

	// 主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 主题名字
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 日志时间
	Timestamp *string `json:"Timestamp,omitempty" name:"Timestamp"`

	// 日志内容
	Content *string `json:"Content,omitempty" name:"Content"`

	// 采集路径
	Filename *string `json:"Filename,omitempty" name:"Filename"`

	// 日志来源设备
	Source *string `json:"Source,omitempty" name:"Source"`
}

type ClsSearchLogs struct {

	// 获取更多检索结果的游标
	Context *string `json:"Context,omitempty" name:"Context"`

	// 搜索结果是否已经全部返回
	Listover *bool `json:"Listover,omitempty" name:"Listover"`

	// 日志内容信息
	Results []*ClsLogObject `json:"Results,omitempty" name:"Results" list`
}

type Compatibility struct {

	// 兼容标志状态码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *int64 `json:"Code,omitempty" name:"Code"`
}

type Compression struct {

	// 智能压缩配置开关
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 压缩规则数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompressionRules []*CompressionRule `json:"CompressionRules,omitempty" name:"CompressionRules" list`
}

type CompressionRule struct {

	// true：需要设置为 ture，启用压缩
	// 注意：此字段可能返回 null，表示取不到有效值。
	Compress *bool `json:"Compress,omitempty" name:"Compress"`

	// 根据文件后缀类型压缩
	// 例如 jpg、txt
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileExtensions []*string `json:"FileExtensions,omitempty" name:"FileExtensions" list`

	// 触发压缩的文件长度最小值，单位为字节数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinLength *int64 `json:"MinLength,omitempty" name:"MinLength"`

	// 触发压缩的文件长度最大值，单位为字节数
	// 最大可设置为 30MB
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxLength *int64 `json:"MaxLength,omitempty" name:"MaxLength"`

	// 文件压缩算法
	// gzip：指定 GZIP 压缩
	// brotli：需要同时指定 GZIP 压缩才可启用
	// 注意：此字段可能返回 null，表示取不到有效值。
	Algorithms []*string `json:"Algorithms,omitempty" name:"Algorithms" list`
}

type CookieKey struct {

	// on | off 是否使用Cookie作为Cache的一部分
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 使用的cookie 逗号分割
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type CreateClsLogTopicRequest struct {
	*tchttp.BaseRequest

	// 日志主题名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 日志集ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// 接入渠道，默认值为cdn
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// 域名区域信息
	DomainAreaConfigs []*DomainAreaConfig `json:"DomainAreaConfigs,omitempty" name:"DomainAreaConfigs" list`
}

func (r *CreateClsLogTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateClsLogTopicRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateClsLogTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 主题ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClsLogTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateClsLogTopicResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteCdnDomainRequest struct {
	*tchttp.BaseRequest

	// 域名
	// 域名状态需要为【已停用】
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DeleteCdnDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteCdnDomainRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteCdnDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCdnDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteCdnDomainResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteClsLogTopicRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 日志集ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// 接入渠道，默认值为cdn
	Channel *string `json:"Channel,omitempty" name:"Channel"`
}

func (r *DeleteClsLogTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClsLogTopicRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteClsLogTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClsLogTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClsLogTopicResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillingDataRequest struct {
	*tchttp.BaseRequest

	// 查询起始时间，如：2018-09-04 10:40:00，返回结果大于等于指定时间
	// 根据指定时间粒度参数不同，会进行向前取整，如指定起始时间为 2018-09-04 10:40:00 按小时粒度查询，返回的第一个数据对应时间点为 2018-09-04 10:00:00
	// 起始时间与结束时间间隔小于等于 90 天
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，如：2018-09-04 10:40:00，返回结果小于等于指定时间
	// 根据指定时间粒度参数不同，会进行向前取整，如指定结束时间为  2018-09-04 10:40:00 按小时粒度查询时，返回的最后一个数据对应时间点为 2018-09-04 10:00:00
	// 起始时间与结束时间间隔小于等于 90 天
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 时间粒度，支持模式如下：
	// min：1 分钟粒度，查询区间需要小于等于 24 小时
	// 5min：5 分钟粒度，查询区间需要小于等于 31 天
	// hour：1 小时粒度，查询区间需要小于等于 31 天内
	// day：天粒度，查询区间需要大于 31 天
	// 
	// Area 字段为 overseas 时暂不支持1分钟粒度数据查询
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 指定域名查询计费数据
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 指定项目 ID 查询，[前往查看项目 ID](https://console.cloud.tencent.com/project)
	// 若 Domain 参数填充了具体域名信息，则返回该域名的计费数据，而非指定项目计费数据
	Project *int64 `json:"Project,omitempty" name:"Project"`

	// 指定加速区域查询计费数据：
	// mainland：中国境内
	// overseas：中国境外
	// 不填充时，默认为 mainland
	Area *string `json:"Area,omitempty" name:"Area"`

	// Area 为 overseas 时，指定国家/地区查询
	// 省份、国家/地区编码可以查看 [省份编码映射](https://cloud.tencent.com/document/product/228/6316#.E7.9C.81.E4.BB.BD.E6.98.A0.E5.B0.84)
	// 不填充时，查询所有国家/地区
	District *int64 `json:"District,omitempty" name:"District"`

	// 计费统计类型
	// flux：计费流量
	// bandwidth：计费带宽
	// 默认为 bandwidth
	Metric *string `json:"Metric,omitempty" name:"Metric"`
}

func (r *DescribeBillingDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillingDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillingDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 时间粒度，根据查询时传递参数指定：
	// min：1 分钟粒度
	// 5min：5 分钟粒度
	// hour：1 小时粒度
	// day：天粒度
		Interval *string `json:"Interval,omitempty" name:"Interval"`

		// 数据明细
		Data []*ResourceBillingData `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillingDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillingDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCdnDataRequest struct {
	*tchttp.BaseRequest

	// 查询起始时间，如：2018-09-04 10:40:00，返回结果大于等于指定时间
	// 根据指定时间粒度不同，会进行向前归整，如 2018-09-04 10:40:00 在按 1 小时的时间粒度查询时，返回的第一个数据对应时间点为 2018-09-04 10:00:00
	// 起始时间与结束时间间隔小于等于 90 天
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，如：2018-09-04 10:40:00，返回结果小于等于指定时间
	// 根据指定时间粒度不同，会进行向前归整，如 2018-09-04 10:40:00 在按 1 小时的时间粒度查询时，返回的最后一个数据对应时间点为 2018-09-04 10:00:00
	// 起始时间与结束时间间隔小于等于 90 天
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 指定查询指标，支持的类型有：
	// flux：流量，单位为 byte
	// bandwidth：带宽，单位为 bps
	// request：请求数，单位为 次
	// fluxHitRate：流量命中率，单位为 %
	// statusCode：状态码，返回 2xx、3xx、4xx、5xx 汇总数据，单位为 个
	// 2xx：返回 2xx 状态码汇总及各 2 开头状态码数据，单位为 个
	// 3xx：返回 3xx 状态码汇总及各 3 开头状态码数据，单位为 个
	// 4xx：返回 4xx 状态码汇总及各 4 开头状态码数据，单位为 个
	// 5xx：返回 5xx 状态码汇总及各 5 开头状态码数据，单位为 个
	// 支持指定具体状态码查询，若未产生过，则返回为空
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// 指定查询域名列表
	// 最多可一次性查询 30 个加速域名明细
	Domains []*string `json:"Domains,omitempty" name:"Domains" list`

	// 指定要查询的项目 ID，[前往查看项目 ID](https://console.cloud.tencent.com/project)
	// 未填充域名情况下，指定项目查询，若填充了具体域名信息，以域名为主
	Project *int64 `json:"Project,omitempty" name:"Project"`

	// 时间粒度，支持以下几种模式：
	// min：1 分钟粒度，指定查询区间 24 小时内（含 24 小时），可返回 1 分钟粒度明细数据（指定查询服务地域为中国境外时不支持 1 分钟粒度）
	// 5min：5 分钟粒度，指定查询区间 31 天内（含 31 天），可返回 5 分钟粒度明细数据
	// hour：1 小时粒度，指定查询区间 31 天内（含 31 天），可返回 1 小时粒度明细数据
	// day：天粒度，指定查询区间大于 31 天，可返回天粒度明细数据
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 多域名查询时，默认（false)返回多个域名的汇总数据
	// 可按需指定为 true，返回每一个 Domain 的明细数据（statusCode 指标暂不支持）
	Detail *bool `json:"Detail,omitempty" name:"Detail"`

	// 查询中国境内CDN数据时，指定运营商查询，不填充表示查询所有运营商
	// 运营商编码可以查看 [运营商编码映射](https://cloud.tencent.com/document/product/228/6316#.E5.8C.BA.E5.9F.9F-.2F-.E8.BF.90.E8.90.A5.E5.95.86.E6.98.A0.E5.B0.84.E8.A1.A8)
	// 指定运营商查询时，不可同时指定省份、IP协议查询
	Isp *int64 `json:"Isp,omitempty" name:"Isp"`

	// 查询中国境内CDN数据时，指定省份查询，不填充表示查询所有省份
	// 查询中国境外CDN数据时，指定国家/地区查询，不填充表示查询所有国家/地区
	// 省份、国家/地区编码可以查看 [省份编码映射](https://cloud.tencent.com/document/product/228/6316#.E5.8C.BA.E5.9F.9F-.2F-.E8.BF.90.E8.90.A5.E5.95.86.E6.98.A0.E5.B0.84.E8.A1.A8)
	// 指定（中国境内）省份查询时，不可同时指定运营商、IP协议查询
	District *int64 `json:"District,omitempty" name:"District"`

	// 指定协议查询，不填充表示查询所有协议
	// all：所有协议
	// http：指定查询 HTTP 对应指标
	// https：指定查询 HTTPS 对应指标
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 指定数据源查询，白名单功能
	DataSource *string `json:"DataSource,omitempty" name:"DataSource"`

	// 指定IP协议查询，不填充表示查询所有协议
	// all：所有协议
	// ipv4：指定查询 ipv4 对应指标
	// ipv6：指定查询 ipv6 对应指标
	// 指定IP协议查询时，不可同时指定省份、运营商查询
	IpProtocol *string `json:"IpProtocol,omitempty" name:"IpProtocol"`

	// 指定服务地域查询，不填充表示查询中国境内CDN数据
	// mainland：指定查询中国境内 CDN 数据
	// overseas：指定查询中国境外 CDN 数据
	Area *string `json:"Area,omitempty" name:"Area"`

	// 查询中国境外CDN数据时，可指定地区类型查询，不填充表示查询服务地区数据（仅在 Area 为 overseas 时可用）
	// server：指定查询服务地区（腾讯云 CDN 节点服务器所在地区）数据
	// client：指定查询客户端地区（用户请求终端所在地区）数据
	AreaType *string `json:"AreaType,omitempty" name:"AreaType"`
}

func (r *DescribeCdnDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCdnDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCdnDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回数据的时间粒度，查询时指定：
	// min：1 分钟粒度
	// 5min：5 分钟粒度
	// hour：1 小时粒度
	// day：天粒度
		Interval *string `json:"Interval,omitempty" name:"Interval"`

		// 指定条件查询得到的数据明细
		Data []*ResourceData `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCdnDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCdnDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCdnDomainLogsRequest struct {
	*tchttp.BaseRequest

	// 指定域名查询
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 开始时间，如 2019-09-04 00:00:00
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，如 2019-09-04 12:00:00
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 分页查询偏移量，默认为 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询限制数目，默认为 100，最大为 1000
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 指定区域下载日志
	// mainland：获取境内加速日志包下载链接
	// overseas：获取境外加速日志包下载链接
	// global：同时获取境内、境外加速日志包下载链接（分开打包）
	// 不指定时默认为 mainland
	Area *string `json:"Area,omitempty" name:"Area"`

	// 指定下载日志的类型。
	// access：获取访问日志
	LogType *string `json:"LogType,omitempty" name:"LogType"`
}

func (r *DescribeCdnDomainLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCdnDomainLogsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCdnDomainLogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志包下载链接
		DomainLogs []*DomainLog `json:"DomainLogs,omitempty" name:"DomainLogs" list`

		// 查询到的总条数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCdnDomainLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCdnDomainLogsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCdnIpRequest struct {
	*tchttp.BaseRequest

	// 需要查询的 IP 列表
	Ips []*string `json:"Ips,omitempty" name:"Ips" list`
}

func (r *DescribeCdnIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCdnIpRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCdnIpResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询的节点归属详情。
		Ips []*CdnIp `json:"Ips,omitempty" name:"Ips" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCdnIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCdnIpResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCertDomainsRequest struct {
	*tchttp.BaseRequest

	// PEM格式证书Base64编码后的字符串
	Cert *string `json:"Cert,omitempty" name:"Cert"`
}

func (r *DescribeCertDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCertDomainsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCertDomainsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 已接入CDN的域名列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Domains []*string `json:"Domains,omitempty" name:"Domains" list`

		// 已配置证书的CDN域名列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		CertifiedDomains []*string `json:"CertifiedDomains,omitempty" name:"CertifiedDomains" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCertDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCertDomainsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainsConfigRequest struct {
	*tchttp.BaseRequest

	// 分页查询偏移量，默认为 0 （第一页）
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询限制数目，默认为 100，最大可设置为 1000
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 查询条件过滤器，复杂类型
	Filters []*DomainFilter `json:"Filters,omitempty" name:"Filters" list`

	// 排序规则
	Sort *Sort `json:"Sort,omitempty" name:"Sort"`
}

func (r *DescribeDomainsConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDomainsConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainsConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 域名列表
		Domains []*DetailDomain `json:"Domains,omitempty" name:"Domains" list`

		// 符合查询条件的域名总数
	// 用于分页查询
		TotalNumber *int64 `json:"TotalNumber,omitempty" name:"TotalNumber"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDomainsConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDomainsConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainsRequest struct {
	*tchttp.BaseRequest

	// 分页查询偏移量，默认为 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询限制数目，默认为 100，最大可设置为 1000
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 查询条件过滤器，复杂类型
	Filters []*DomainFilter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDomainsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 域名列表
		Domains []*BriefDomain `json:"Domains,omitempty" name:"Domains" list`

		// 符合查询条件的域名总数
	// 用于分页查询
		TotalNumber *int64 `json:"TotalNumber,omitempty" name:"TotalNumber"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDomainsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageConfigRequest struct {
	*tchttp.BaseRequest

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeImageConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// WebpAdapter配置
	// 注意：此字段可能返回 null，表示取不到有效值。
		WebpAdapter *WebpAdapter `json:"WebpAdapter,omitempty" name:"WebpAdapter"`

		// TpgAdapter配置
	// 注意：此字段可能返回 null，表示取不到有效值。
		TpgAdapter *TpgAdapter `json:"TpgAdapter,omitempty" name:"TpgAdapter"`

		// GuetzliAdapter配置
	// 注意：此字段可能返回 null，表示取不到有效值。
		GuetzliAdapter *GuetzliAdapter `json:"GuetzliAdapter,omitempty" name:"GuetzliAdapter"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeIpStatusRequest struct {
	*tchttp.BaseRequest

	// 加速域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 节点类型：
	// edge：表示边缘节点
	// last：表示回源层节点
	// 不填充情况下，默认返回边缘节点信息
	Layer *string `json:"Layer,omitempty" name:"Layer"`
}

func (r *DescribeIpStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeIpStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeIpStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 节点列表
		Ips []*IpStatus `json:"Ips,omitempty" name:"Ips" list`

		// 节点总个数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIpStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeIpStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeIpVisitRequest struct {
	*tchttp.BaseRequest

	// 查询起始时间，如：2018-09-04 10:40:10，返回结果大于等于指定时间
	// 根据指定时间粒度不同，会进行向前归整，如 2018-09-04 10:40:10 在按 5 分钟的时间粒度查询时，返回的第一个数据对应时间点为 2018-09-04 10:40:00
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，如：2018-09-04 10:40:10，返回结果小于等于指定时间
	// 根据指定时间粒度不同，会进行向前归整，如 2018-09-04 10:40:10 在按 5 分钟的时间粒度查询时，返回的最后一个数据对应时间点为 2018-09-04 10:40:00
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 指定查询域名列表，最多可一次性查询 30 个加速域名明细
	Domains []*string `json:"Domains,omitempty" name:"Domains" list`

	// 指定要查询的项目 ID，[前往查看项目 ID](https://console.cloud.tencent.com/project)
	// 未填充域名情况下，指定项目查询，若填充了具体域名信息，以域名为主
	Project *int64 `json:"Project,omitempty" name:"Project"`

	// 时间粒度，支持以下几种模式：
	// 5min：5 分钟粒度，查询时间区间 24 小时内，默认返回 5 分钟粒度活跃用户数
	// day：天粒度，查询时间区间大于 1 天时，默认返回天粒度活跃用户数
	Interval *string `json:"Interval,omitempty" name:"Interval"`
}

func (r *DescribeIpVisitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeIpVisitRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeIpVisitResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据统计的时间粒度，支持5min,  day，分别表示5分钟，1天的时间粒度。
		Interval *string `json:"Interval,omitempty" name:"Interval"`

		// 各个资源的回源数据详情。
		Data []*ResourceData `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIpVisitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeIpVisitResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMapInfoRequest struct {
	*tchttp.BaseRequest

	// 映射查询类别：
	// isp：运营商映射查询
	// district：省份（中国境内）、国家/地区（中国境外）映射查询
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeMapInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMapInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMapInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 映射关系数组。
		MapInfoList []*MapInfo `json:"MapInfoList,omitempty" name:"MapInfoList" list`

		// 服务端区域id和子区域id的映射关系。
	// 注意：此字段可能返回 null，表示取不到有效值。
		ServerRegionRelation []*RegionMapRelation `json:"ServerRegionRelation,omitempty" name:"ServerRegionRelation" list`

		// 客户端区域id和子区域id的映射关系。
	// 注意：此字段可能返回 null，表示取不到有效值。
		ClientRegionRelation []*RegionMapRelation `json:"ClientRegionRelation,omitempty" name:"ClientRegionRelation" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMapInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMapInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeOriginDataRequest struct {
	*tchttp.BaseRequest

	// 查询起始时间，如：2018-09-04 10:40:00，返回结果大于等于指定时间
	// 根据指定时间粒度不同，会进行向前归整，如 2018-09-04 10:40:00 在按 1 小时的时间粒度查询时，返回的第一个数据对应时间点为 2018-09-04 10:00:00
	// 起始时间与结束时间间隔小于等于 90 天
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，如：2018-09-04 10:40:00，返回结果小于等于指定时间
	// 根据指定时间粒度不同，会进行向前归整，如 2018-09-04 10:40:00 在按 1 小时的时间粒度查询时，返回的最后一个数据对应时间点为 2018-09-04 10:00:00
	// 起始时间与结束时间间隔小于等于 90 天
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 指定查询指标，支持的类型有：
	// flux：回源流量，单位为 byte
	// bandwidth：回源带宽，单位为 bps
	// request：回源请求数，单位为 次
	// failRequest：回源失败请求数，单位为 次
	// failRate：回源失败率，单位为 %
	// statusCode：回源状态码，返回 2xx、3xx、4xx、5xx 汇总数据，单位为 个
	// 2xx：返回 2xx 回源状态码汇总及各 2 开头回源状态码数据，单位为 个
	// 3xx：返回 3xx 回源状态码汇总及各 3 开头回源状态码数据，单位为 个
	// 4xx：返回 4xx 回源状态码汇总及各 4 开头回源状态码数据，单位为 个
	// 5xx：返回 5xx 回源状态码汇总及各 5 开头回源状态码数据，单位为 个
	// 支持指定具体状态码查询，若未产生过，则返回为空
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// 指定查询域名列表，最多可一次性查询 30 个加速域名明细
	Domains []*string `json:"Domains,omitempty" name:"Domains" list`

	// 指定要查询的项目 ID，[前往查看项目 ID](https://console.cloud.tencent.com/project)
	// 未填充域名情况下，指定项目查询，最多可一次性查询 30 个加速域名明细
	// 若填充了具体域名信息，以域名为主
	Project *int64 `json:"Project,omitempty" name:"Project"`

	// 时间粒度，支持以下几种模式：
	// min：1 分钟粒度，指定查询区间 24 小时内（含 24 小时），可返回 1 分钟粒度明细数据（指定查询服务地域为中国境外时不支持 1 分钟粒度）
	// 5min：5 分钟粒度，指定查询区间 31 天内（含 31 天），可返回 5 分钟粒度明细数据
	// hour：1 小时粒度，指定查询区间 31 天内（含 31 天），可返回 1 小时粒度明细数据
	// day：天粒度，指定查询区间大于 31 天，可返回天粒度明细数据
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Domains 传入多个时，默认（false)返回多个域名的汇总数据
	// 可按需指定为 true，返回每一个 Domain 的明细数据（statusCode 指标暂不支持）
	Detail *bool `json:"Detail,omitempty" name:"Detail"`

	// 指定服务地域查询，不填充表示查询中国境内 CDN 数据
	// mainland：指定查询中国境内 CDN 数据
	// overseas：指定查询中国境外 CDN 数据
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeOriginDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeOriginDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeOriginDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据统计的时间粒度，支持min, 5min, hour, day，分别表示1分钟，5分钟，1小时和1天的时间粒度。
		Interval *string `json:"Interval,omitempty" name:"Interval"`

		// 各个资源的回源数据详情。
		Data []*ResourceOriginData `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOriginDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeOriginDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePayTypeRequest struct {
	*tchttp.BaseRequest

	// 指定服务地域查询
	// mainland：境内计费方式查询
	// overseas：境外计费方式查询
	// 未填充时默认为 mainland
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribePayTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePayTypeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePayTypeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 计费类型：
	// flux：流量计费
	// bandwidth：带宽计费
	// 日结计费方式切换时，若当日产生消耗，则此字段表示第二天即将生效的计费方式，若未产生消耗，则表示已经生效的计费方式。
		PayType *string `json:"PayType,omitempty" name:"PayType"`

		// 计费周期：
	// day：日结计费
	// month：月结计费
		BillingCycle *string `json:"BillingCycle,omitempty" name:"BillingCycle"`

		// 计费方式：
	// monthMax：日峰值月平均计费，月结模式
	// day95：日 95 带宽计费，月结模式
	// month95：月95带宽计费，月结模式
	// sum：总流量计费，日结与月结均有流量计费模式
	// max：峰值带宽计费，日结模式
		StatType *string `json:"StatType,omitempty" name:"StatType"`

		// 境外计费类型：
	// all：全地区统一计费
	// multiple：分地区计费
		RegionType *string `json:"RegionType,omitempty" name:"RegionType"`

		// 当前生效计费类型：
	// flux：流量计费
	// bandwidth：带宽计费
		CurrentPayType *string `json:"CurrentPayType,omitempty" name:"CurrentPayType"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePayTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePayTypeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePurgeQuotaRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribePurgeQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePurgeQuotaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePurgeQuotaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// URL刷新用量及配额。
		UrlPurge []*Quota `json:"UrlPurge,omitempty" name:"UrlPurge" list`

		// 目录刷新用量及配额。
		PathPurge []*Quota `json:"PathPurge,omitempty" name:"PathPurge" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePurgeQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePurgeQuotaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePurgeTasksRequest struct {
	*tchttp.BaseRequest

	// 指定刷新类型查询
	// url：url 刷新记录
	// path：目录刷新记录
	PurgeType *string `json:"PurgeType,omitempty" name:"PurgeType"`

	// 根据时间区间查询时，填充开始时间，如 2018-08-08 00:00:00
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 根据时间区间查询时，填充结束时间，如 2018-08-08 23:59:59
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 根据任务 ID 查询时，填充任务 ID
	// 查询时任务 ID 与起始时间必须填充一项
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 分页查询偏移量，默认为 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询限制数目，默认为 20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 支持域名过滤，或 http(s):// 开头完整 URL 过滤
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// 指定任务状态查询
	// fail：刷新失败
	// done：刷新成功
	// process：刷新中
	Status *string `json:"Status,omitempty" name:"Status"`

	// 指定刷新地域查询
	// mainland：境内
	// overseas：境外
	// global：全球
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribePurgeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePurgeTasksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePurgeTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 详细刷新记录
	// 注意：此字段可能返回 null，表示取不到有效值。
		PurgeLogs []*PurgeTask `json:"PurgeLogs,omitempty" name:"PurgeLogs" list`

		// 任务总数，用于分页
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePurgeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePurgeTasksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePushQuotaRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribePushQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePushQuotaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePushQuotaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Url预热用量及配额。
		UrlPush []*Quota `json:"UrlPush,omitempty" name:"UrlPush" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePushQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePushQuotaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePushTasksRequest struct {
	*tchttp.BaseRequest

	// 开始时间，如2018-08-08 00:00:00。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，如2018-08-08 23:59:59。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 指定任务 ID 查询
	// TaskId 和起始时间必须指定一项
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 查询关键字，请输入域名或 http(s):// 开头完整 URL
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// 分页查询偏移量，默认为 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询限制数目，默认为 20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 指定地区查询预热纪录
	// mainland：境内
	// overseas：境外
	// global：全球
	Area *string `json:"Area,omitempty" name:"Area"`

	// 指定任务状态查询
	// fail：预热失败
	// done：预热成功
	// process：预热中
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *DescribePushTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePushTasksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePushTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 预热历史记录
	// 注意：此字段可能返回 null，表示取不到有效值。
		PushLogs []*PushTask `json:"PushLogs,omitempty" name:"PushLogs" list`

		// 任务总数，用于分页
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePushTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePushTasksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeReportDataRequest struct {
	*tchttp.BaseRequest

	// 查询起始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 报表类型
	// daily：日报表
	// weekly：周报表
	// monthly：月报表
	ReportType *string `json:"ReportType,omitempty" name:"ReportType"`

	// 域名加速区域
	// mainland：中国境内
	// overseas：中国境外
	Area *string `json:"Area,omitempty" name:"Area"`

	// 偏移量，默认0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 数据个数，默认1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 按项目ID筛选
	Project *int64 `json:"Project,omitempty" name:"Project"`
}

func (r *DescribeReportDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeReportDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeReportDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 域名维度数据详情
		DomainReport []*ReportData `json:"DomainReport,omitempty" name:"DomainReport" list`

		// 项目维度数据详情
		ProjectReport []*ReportData `json:"ProjectReport,omitempty" name:"ProjectReport" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReportDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeReportDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTrafficPackagesRequest struct {
	*tchttp.BaseRequest

	// 分页查询起始地址，默认 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询记录个数，默认100，最大1000
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTrafficPackagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTrafficPackagesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTrafficPackagesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 流量包总个数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 流量包详情
		TrafficPackages []*TrafficPackage `json:"TrafficPackages,omitempty" name:"TrafficPackages" list`

		// 即将过期的流量包个数（7天内）
		ExpiringCount *int64 `json:"ExpiringCount,omitempty" name:"ExpiringCount"`

		// 有效流量包个数
		EnabledCount *int64 `json:"EnabledCount,omitempty" name:"EnabledCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTrafficPackagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTrafficPackagesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUrlViolationsRequest struct {
	*tchttp.BaseRequest

	// 分页查询偏移量，默认为 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询限制数目，默认为 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 指定的域名查询
	Domains []*string `json:"Domains,omitempty" name:"Domains" list`
}

func (r *DescribeUrlViolationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUrlViolationsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUrlViolationsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 违规 URL 详情
	// 注意：此字段可能返回 null，表示取不到有效值。
		UrlRecordList []*ViolationUrl `json:"UrlRecordList,omitempty" name:"UrlRecordList" list`

		// 记录总数，用于分页
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUrlViolationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUrlViolationsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetailDomain struct {

	// 域名 ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 腾讯云账号ID
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 加速域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名对应的 CNAME 地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cname *string `json:"Cname,omitempty" name:"Cname"`

	// 加速服务状态
	// rejected：域名审核未通过，域名备案过期/被注销导致
	// processing：部署中
	// online：已启动
	// offline：已关闭
	Status *string `json:"Status,omitempty" name:"Status"`

	// 项目 ID，可前往腾讯云项目管理页面查看
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 域名业务类型
	// web：静态加速
	// download：下载加速
	// media：流媒体点播加速
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 域名创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 域名更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 源站配置
	Origin *Origin `json:"Origin,omitempty" name:"Origin"`

	// IP 黑白名单配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpFilter *IpFilter `json:"IpFilter,omitempty" name:"IpFilter"`

	// IP 访问限频配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpFreqLimit *IpFreqLimit `json:"IpFreqLimit,omitempty" name:"IpFreqLimit"`

	// 状态码缓存配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusCodeCache *StatusCodeCache `json:"StatusCodeCache,omitempty" name:"StatusCodeCache"`

	// 智能压缩配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Compression *Compression `json:"Compression,omitempty" name:"Compression"`

	// 带宽封顶配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	BandwidthAlert *BandwidthAlert `json:"BandwidthAlert,omitempty" name:"BandwidthAlert"`

	// Range 回源配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	RangeOriginPull *RangeOriginPull `json:"RangeOriginPull,omitempty" name:"RangeOriginPull"`

	// 301/302 回源自动跟随配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	FollowRedirect *FollowRedirect `json:"FollowRedirect,omitempty" name:"FollowRedirect"`

	// 自定义错误页面配置（功能灰度中，敬请期待）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorPage *ErrorPage `json:"ErrorPage,omitempty" name:"ErrorPage"`

	// 自定义请求头部配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	RequestHeader *RequestHeader `json:"RequestHeader,omitempty" name:"RequestHeader"`

	// 自定义响应头部配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResponseHeader *ResponseHeader `json:"ResponseHeader,omitempty" name:"ResponseHeader"`

	// 单链接下行限速配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownstreamCapping *DownstreamCapping `json:"DownstreamCapping,omitempty" name:"DownstreamCapping"`

	// 带参/不带参缓存配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheKey *CacheKey `json:"CacheKey,omitempty" name:"CacheKey"`

	// 源站头部缓存配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResponseHeaderCache *ResponseHeaderCache `json:"ResponseHeaderCache,omitempty" name:"ResponseHeaderCache"`

	// 视频拖拽配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoSeek *VideoSeek `json:"VideoSeek,omitempty" name:"VideoSeek"`

	// 节点缓存过期规则配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cache *Cache `json:"Cache,omitempty" name:"Cache"`

	// 跨国链路优化配置（功能灰度中，敬请期待）
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginPullOptimization *OriginPullOptimization `json:"OriginPullOptimization,omitempty" name:"OriginPullOptimization"`

	// Https 加速相关配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Https *Https `json:"Https,omitempty" name:"Https"`

	// 时间戳防盗链配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Authentication *Authentication `json:"Authentication,omitempty" name:"Authentication"`

	// SEO 优化配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Seo *Seo `json:"Seo,omitempty" name:"Seo"`

	// 域名封禁状态
	// normal：正常状态
	// overdue：账号欠费导致域名关闭，充值完成后可自行启动加速服务
	// malicious：域名出现恶意行为，强制关闭加速服务
	// ddos：域名被大规模 DDoS 攻击，关闭加速服务
	// idle：域名超过 90 天内无任何操作、数据产生，判定为不活跃域名自动关闭加速服务，可自行启动加速服务
	// unlicensed：域名未备案/备案注销，自动关闭加速服务，备案完成后可自行启动加速服务
	// capping：触发配置的带宽阈值上限
	// readonly：域名存在特殊配置，被锁定
	// 注意：此字段可能返回 null，表示取不到有效值。
	Disable *string `json:"Disable,omitempty" name:"Disable"`

	// 访问协议强制跳转配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitempty" name:"ForceRedirect"`

	// Referer 防盗链配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Referer *Referer `json:"Referer,omitempty" name:"Referer"`

	// 浏览器缓存过期规则配置（功能灰度中，敬请期待）
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxAge *MaxAge `json:"MaxAge,omitempty" name:"MaxAge"`

	// Ipv6 配置（功能灰度中，敬请期待）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ipv6 *Ipv6 `json:"Ipv6,omitempty" name:"Ipv6"`

	// 是否兼容旧版本配置（内部兼容性字段）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Compatibility *Compatibility `json:"Compatibility,omitempty" name:"Compatibility"`

	// 区域特殊配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecificConfig *SpecificConfig `json:"SpecificConfig,omitempty" name:"SpecificConfig"`

	// 加速区域
	// mainland：中国境内加速
	// overseas：中国境外加速
	// global：全球加速
	// 注意：此字段可能返回 null，表示取不到有效值。
	Area *string `json:"Area,omitempty" name:"Area"`

	// 域名锁定状态
	// normal：未锁定
	// mainland：中国境内锁定
	// overseas：中国境外锁定
	// global：全球锁定
	// 注意：此字段可能返回 null，表示取不到有效值。
	Readonly *string `json:"Readonly,omitempty" name:"Readonly"`

	// 回源超时配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginPullTimeout *OriginPullTimeout `json:"OriginPullTimeout,omitempty" name:"OriginPullTimeout"`

	// 回源S3鉴权配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	AwsPrivateAccess *AwsPrivateAccess `json:"AwsPrivateAccess,omitempty" name:"AwsPrivateAccess"`

	// Scdn配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityConfig *SecurityConfig `json:"SecurityConfig,omitempty" name:"SecurityConfig"`

	// ImageOptimization配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageOptimization *ImageOptimization `json:"ImageOptimization,omitempty" name:"ImageOptimization"`

	// UA黑白名单配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserAgentFilter *UserAgentFilter `json:"UserAgentFilter,omitempty" name:"UserAgentFilter"`
}

type DisableCachesRequest struct {
	*tchttp.BaseRequest

	// 需要禁用的 URL 列表
	// 每次最多可提交 100 条，每日最多可提交 3000 条
	Urls []*string `json:"Urls,omitempty" name:"Urls" list`
}

func (r *DisableCachesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableCachesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableCachesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 提交结果
	// 注意：此字段可能返回 null，表示取不到有效值。
		CacheOptResult *CacheOptResult `json:"CacheOptResult,omitempty" name:"CacheOptResult"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableCachesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableCachesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableClsLogTopicRequest struct {
	*tchttp.BaseRequest

	// 日志集ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// 日志主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 接入渠道，默认值为cdn
	Channel *string `json:"Channel,omitempty" name:"Channel"`
}

func (r *DisableClsLogTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableClsLogTopicRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableClsLogTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableClsLogTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableClsLogTopicResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DomainAreaConfig struct {

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 地区列表，其中元素可为mainland/overseas
	Area []*string `json:"Area,omitempty" name:"Area" list`
}

type DomainFilter struct {

	// 过滤字段名，支持的列表如下：
	// - origin：主源站。
	// - domain：域名。
	// - resourceId：域名id。
	// - status：域名状态，online，offline或processing。
	// - serviceType：业务类型，web，download或media。
	// - projectId：项目ID。
	// - domainType：主源站类型，cname表示自有源，cos表示cos接入。
	// - fullUrlCache：全路径缓存，on或off。
	// - https：是否配置https，on，off或processing。
	// - originPullProtocol：回源协议类型，支持http，follow或https。
	// - tagKey：标签键。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 过滤字段值。
	Value []*string `json:"Value,omitempty" name:"Value" list`

	// 是否启用模糊查询，仅支持过滤字段名为origin，domain。
	// 模糊查询时，Value长度最大为1，否则Value长度最大为5。
	Fuzzy *bool `json:"Fuzzy,omitempty" name:"Fuzzy"`
}

type DomainLog struct {

	// 日志包起始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 日志包结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 日志包下载链接
	LogPath *string `json:"LogPath,omitempty" name:"LogPath"`

	// 日志包对应加速区域
	// mainland：境内
	// overseas：境外
	Area *string `json:"Area,omitempty" name:"Area"`

	// 日志包文件名
	LogName *string `json:"LogName,omitempty" name:"LogName"`
}

type DownstreamCapping struct {

	// 下行速度配置开关
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 下行限速规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	CappingRules []*CappingRule `json:"CappingRules,omitempty" name:"CappingRules" list`
}

type EnableCachesRequest struct {
	*tchttp.BaseRequest

	// 解封 URL 列表
	Urls []*string `json:"Urls,omitempty" name:"Urls" list`
}

func (r *EnableCachesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableCachesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableCachesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		CacheOptResult *CacheOptResult `json:"CacheOptResult,omitempty" name:"CacheOptResult"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableCachesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableCachesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableClsLogTopicRequest struct {
	*tchttp.BaseRequest

	// 日志集ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// 日志主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 接入渠道，默认值为cdn
	Channel *string `json:"Channel,omitempty" name:"Channel"`
}

func (r *EnableClsLogTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableClsLogTopicRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableClsLogTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableClsLogTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableClsLogTopicResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ErrorPage struct {

	// 状态码重定向配置开关
	// on：开启
	// off：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 状态码重定向规则配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageRules []*ErrorPageRule `json:"PageRules,omitempty" name:"PageRules" list`
}

type ErrorPageRule struct {

	// 状态码
	// 支持 400、403、404、500
	StatusCode *int64 `json:"StatusCode,omitempty" name:"StatusCode"`

	// 重定向状态码设置
	// 支持 301 或 302
	RedirectCode *int64 `json:"RedirectCode,omitempty" name:"RedirectCode"`

	// 重定向 URL
	// 需要为完整跳转路径，如 https://www.test.com/error.html
	RedirectUrl *string `json:"RedirectUrl,omitempty" name:"RedirectUrl"`
}

type FollowRedirect struct {

	// 回源跟随开关
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type ForceRedirect struct {

	// 访问强制跳转配置开关
	// on：开启
	// off：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 访问强制跳转类型
	// http：强制 http 跳转
	// https：强制 https 跳转
	// 注意：此字段可能返回 null，表示取不到有效值。
	RedirectType *string `json:"RedirectType,omitempty" name:"RedirectType"`

	// 强制跳转时返回状态码 
	// 支持 301、302
	// 注意：此字段可能返回 null，表示取不到有效值。
	RedirectStatusCode *int64 `json:"RedirectStatusCode,omitempty" name:"RedirectStatusCode"`
}

type GetDisableRecordsRequest struct {
	*tchttp.BaseRequest

	// 开始时间，如：2018-12-12 10:24:00。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，如：2018-12-14 10:24:00。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 指定 URL 查询
	Url *string `json:"Url,omitempty" name:"Url"`

	// URL 当前状态
	// disable：当前仍为禁用状态，访问返回 403
	// enable：当前为可用状态，已解禁，可正常访问
	Status *string `json:"Status,omitempty" name:"Status"`

	// 分页查询偏移量，默认为 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询限制数目，默认为20。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetDisableRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDisableRecordsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetDisableRecordsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 封禁历史记录
	// 注意：此字段可能返回 null，表示取不到有效值。
		UrlRecordList []*UrlRecord `json:"UrlRecordList,omitempty" name:"UrlRecordList" list`

		// 任务总数，用于分页
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDisableRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDisableRecordsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GuetzliAdapter struct {

	// 开关，"on/off"
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type HeaderKey struct {

	// 是否组成Cachekey
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 组成CacheKey的header 逗号分隔
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type HttpHeaderPathRule struct {

	// http 头部设置方式
	// add：添加头部，若已存在头部，则会存在重复头部
	// set：仅回源头部配置支持，若头部已存在则会覆盖原有头部值，若不存在，则会增加该头部及值
	// del：删除头部
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeaderMode *string `json:"HeaderMode,omitempty" name:"HeaderMode"`

	// http 头部名称，最多可设置 100 个字符
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeaderName *string `json:"HeaderName,omitempty" name:"HeaderName"`

	// http 头部值，最多可设置 1000 个字符
	// Mode 为 del 时非必填
	// Mode 为 add/set 时必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeaderValue *string `json:"HeaderValue,omitempty" name:"HeaderValue"`

	// 规则类型：
	// all：所有文件生效
	// file：指定文件后缀生效
	// directory：指定路径生效
	// path：指定绝对路径生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`

	// RuleType 对应类型下的匹配内容：
	// all 时填充 *
	// file 时填充后缀名，如 jpg、txt
	// directory 时填充路径，如 /xxx/test/
	// path 时填充绝对路径，如 /xxx/test.html
	// 注意：此字段可能返回 null，表示取不到有效值。
	RulePaths []*string `json:"RulePaths,omitempty" name:"RulePaths" list`
}

type Https struct {

	// https 配置开关
	// on：开启
	// off：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// http2 配置开关
	// on：开启
	// off：关闭
	// 初次启用 https 加速会默认开启 http2 配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Http2 *string `json:"Http2,omitempty" name:"Http2"`

	// OCSP 配置开关
	// on：开启
	// off：关闭
	// 默认为关闭状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcspStapling *string `json:"OcspStapling,omitempty" name:"OcspStapling"`

	// 客户端证书校验功能
	// on：开启
	// off：关闭
	// 默认为关闭状态，开启时需要上传客户端证书信息，该配置项目前在灰度中，尚未全量
	// 注意：此字段可能返回 null，表示取不到有效值。
	VerifyClient *string `json:"VerifyClient,omitempty" name:"VerifyClient"`

	// 服务端证书配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertInfo *ServerCert `json:"CertInfo,omitempty" name:"CertInfo"`

	// 客户端证书配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientCertInfo *ClientCert `json:"ClientCertInfo,omitempty" name:"ClientCertInfo"`

	// Spdy 配置开关
	// on：开启
	// off：关闭
	// 默认为关闭状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Spdy *string `json:"Spdy,omitempty" name:"Spdy"`

	// https 证书部署状态
	// closed：已关闭
	// deploying：部署中
	// deployed：部署成功
	// failed：部署失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	SslStatus *string `json:"SslStatus,omitempty" name:"SslStatus"`
}

type ImageOptimization struct {

	// WebpAdapter配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebpAdapter *WebpAdapter `json:"WebpAdapter,omitempty" name:"WebpAdapter"`

	// TpgAdapter配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	TpgAdapter *TpgAdapter `json:"TpgAdapter,omitempty" name:"TpgAdapter"`

	// GuetzliAdapter配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	GuetzliAdapter *GuetzliAdapter `json:"GuetzliAdapter,omitempty" name:"GuetzliAdapter"`
}

type IpFilter struct {

	// IP 黑白名单配置开关
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// IP 黑白名单类型
	// whitelist：白名单
	// blacklist：黑名单
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilterType *string `json:"FilterType,omitempty" name:"FilterType"`

	// IP 黑白名单列表
	// 支持 X.X.X.X 形式 IP，或 /8、 /16、/24 形式网段
	// 最多可填充 50 个白名单或 50 个黑名单
	// 注意：此字段可能返回 null，表示取不到有效值。
	Filters []*string `json:"Filters,omitempty" name:"Filters" list`
}

type IpFreqLimit struct {

	// IP 限频配置开关
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 设置每秒请求数限制
	// 超出限制的请求会直接返回 514
	// 注意：此字段可能返回 null，表示取不到有效值。
	Qps *int64 `json:"Qps,omitempty" name:"Qps"`
}

type IpStatus struct {

	// 节点 IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 节点所属区域
	District *string `json:"District,omitempty" name:"District"`

	// 节点所属运营商
	Isp *string `json:"Isp,omitempty" name:"Isp"`

	// 节点所在城市
	City *string `json:"City,omitempty" name:"City"`

	// 节点状态
	// online：上线状态，正常调度服务中
	// offline：下线状态
	Status *string `json:"Status,omitempty" name:"Status"`
}

type Ipv6 struct {

	// 域名是否开启ipv6功能，on或off。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type ListClsLogTopicsRequest struct {
	*tchttp.BaseRequest

	// 接入渠道，默认值为cdn
	Channel *string `json:"Channel,omitempty" name:"Channel"`
}

func (r *ListClsLogTopicsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListClsLogTopicsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListClsLogTopicsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志集信息
		Logset *LogSetInfo `json:"Logset,omitempty" name:"Logset"`

		// 日志主题信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Topics []*TopicInfo `json:"Topics,omitempty" name:"Topics" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListClsLogTopicsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListClsLogTopicsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListClsTopicDomainsRequest struct {
	*tchttp.BaseRequest

	// 日志集ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// 日志主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 接入渠道，默认值为cdn
	Channel *string `json:"Channel,omitempty" name:"Channel"`
}

func (r *ListClsTopicDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListClsTopicDomainsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListClsTopicDomainsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 开发者ID
		AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

		// 渠道
		Channel *string `json:"Channel,omitempty" name:"Channel"`

		// 日志集ID
		LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

		// 日志主题ID
		TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

		// 域名区域配置，其中可能含有已删除的域名，如果要再传回ManageClsTopicDomains接口，需要结合ListCdnDomains接口排除掉已删除的域名
		DomainAreaConfigs []*DomainAreaConfig `json:"DomainAreaConfigs,omitempty" name:"DomainAreaConfigs" list`

		// 日志主题名称
		TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

		// 日志主题最近更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
		UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListClsTopicDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListClsTopicDomainsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListTopDataRequest struct {
	*tchttp.BaseRequest

	// 查询起始日期，如：2018-09-09
	// 仅支持按天粒度的数据查询，取入参中的天信息作为起始日期
	// 返回大于等于起始日期当天 00:00:00 点产生的数据
	// 仅支持 90 天内数据查询
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束日期，如：2018-09-10
	// 仅支持按天粒度的数据查询，取入参中的天信息作为结束日期
	// 返回小于等于结束日期当天 23:59:59 产生的数据
	// EndTime 需要大于等于 StartTime
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 排序对象，支持以下几种形式：
	// url：访问 URL 排序，带参数统计，支持的 Filter 为 flux、request
	// path：访问 URL 排序，不带参数统计，支持的 Filter 为 flux、request（白名单功能）
	// district：省份、国家/地区排序，支持的 Filter 为 flux、request
	// isp：运营商排序，支持的 Filter 为 flux、request
	// host：域名访问数据排序，支持的 Filter 为：flux、request、bandwidth、fluxHitRate、2XX、3XX、4XX、5XX、statusCode
	// originHost：域名回源数据排序，支持的 Filter 为 flux、request、bandwidth、origin_2XX、origin_3XX、origin_4XX、origin_5XX、OriginStatusCode
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// 排序使用的指标名称：
	// flux：Metric 为 host 时指代访问流量，originHost 时指代回源流量
	// bandwidth：Metric 为 host 时指代访问带宽，originHost 时指代回源带宽
	// request：Metric 为 host 时指代访问请求数，originHost 时指代回源请求数
	// fluxHitRate：平均流量命中率
	// 2XX：访问 2XX 状态码
	// 3XX：访问 3XX 状态码
	// 4XX：访问 4XX 状态码
	// 5XX：访问 5XX 状态码
	// origin_2XX：回源 2XX 状态码
	// origin_3XX：回源 3XX 状态码
	// origin_4XX：回源 4XX 状态码
	// origin_5XX：回源 5XX 状态码
	// statusCode：指定访问状态码统计，在 Code 参数中填充指定状态码
	// OriginStatusCode：指定回源状态码统计，在 Code 参数中填充指定状态码
	Filter *string `json:"Filter,omitempty" name:"Filter"`

	// 指定查询域名列表，最多可一次性查询 30 个加速域名明细
	Domains []*string `json:"Domains,omitempty" name:"Domains" list`

	// 指定要查询的项目 ID，[前往查看项目 ID](https://console.cloud.tencent.com/project)
	// 未填充域名情况下，指定项目查询，若填充了具体域名信息，以域名为主
	Project *int64 `json:"Project,omitempty" name:"Project"`

	// 多域名查询时，默认（false)返回所有域名汇总排序结果
	// Metric 为 url、path、district、isp，Filter 为 flux、request 时，可设置为 true，返回每一个 Domain 的排序数据
	Detail *bool `json:"Detail,omitempty" name:"Detail"`

	// Filter 为 statusCode、OriginStatusCode 时，填充指定状态码查询排序结果
	Code *string `json:"Code,omitempty" name:"Code"`

	// 指定服务地域查询，不填充表示查询中国境内 CDN 数据
	// mainland：指定查询中国境内 CDN 数据
	// overseas：指定查询中国境外 CDN 数据，支持的 Metric 为 url、district、host、originHost，当 Metric 为 originHost 时仅支持 flux、request、bandwidth Filter
	Area *string `json:"Area,omitempty" name:"Area"`

	// 查询中国境外CDN数据，且仅当 Metric 为 district 或 host 时，可指定地区类型查询，不填充表示查询服务地区数据（仅在 Area 为 overseas，且 Metric 是 district 或 host 时可用）
	// server：指定查询服务地区（腾讯云 CDN 节点服务器所在地区）数据
	// client：指定查询客户端地区（用户请求终端所在地区）数据，当 Metric 为 host 时仅支持 flux、request、bandwidth Filter
	AreaType *string `json:"AreaType,omitempty" name:"AreaType"`
}

func (r *ListTopDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListTopDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListTopDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 各个资源的Top 访问数据详情。
		Data []*TopData `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListTopDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListTopDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LogSetInfo struct {

	// 开发者ID
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// 渠道
	// 注意：此字段可能返回 null，表示取不到有效值。
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// 日志集ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// 日志集名字
	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`

	// 是否默认日志集
	IsDefault *uint64 `json:"IsDefault,omitempty" name:"IsDefault"`

	// 日志保存时间，单位为天
	LogsetSavePeriod *uint64 `json:"LogsetSavePeriod,omitempty" name:"LogsetSavePeriod"`

	// 创建日期
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 区域
	Region *string `json:"Region,omitempty" name:"Region"`
}

type MainlandConfig struct {

	// 时间戳防盗链配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Authentication *Authentication `json:"Authentication,omitempty" name:"Authentication"`

	// 带宽封顶配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BandwidthAlert *BandwidthAlert `json:"BandwidthAlert,omitempty" name:"BandwidthAlert"`

	// 缓存规则配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cache *Cache `json:"Cache,omitempty" name:"Cache"`

	// 缓存相关配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheKey *CacheKey `json:"CacheKey,omitempty" name:"CacheKey"`

	// 智能压缩配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Compression *Compression `json:"Compression,omitempty" name:"Compression"`

	// 下载限速配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownstreamCapping *DownstreamCapping `json:"DownstreamCapping,omitempty" name:"DownstreamCapping"`

	// 错误码重定向配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorPage *ErrorPage `json:"ErrorPage,omitempty" name:"ErrorPage"`

	// 301和302自动回源跟随配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FollowRedirect *FollowRedirect `json:"FollowRedirect,omitempty" name:"FollowRedirect"`

	// 访问协议强制跳转配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitempty" name:"ForceRedirect"`

	// Https配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Https *Https `json:"Https,omitempty" name:"Https"`

	// IP黑白名单配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpFilter *IpFilter `json:"IpFilter,omitempty" name:"IpFilter"`

	// IP限频配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpFreqLimit *IpFreqLimit `json:"IpFreqLimit,omitempty" name:"IpFreqLimit"`

	// 浏览器缓存规则配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxAge *MaxAge `json:"MaxAge,omitempty" name:"MaxAge"`

	// 源站配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Origin *Origin `json:"Origin,omitempty" name:"Origin"`

	// 跨国优化配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginPullOptimization *OriginPullOptimization `json:"OriginPullOptimization,omitempty" name:"OriginPullOptimization"`

	// Range回源配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RangeOriginPull *RangeOriginPull `json:"RangeOriginPull,omitempty" name:"RangeOriginPull"`

	// 防盗链配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Referer *Referer `json:"Referer,omitempty" name:"Referer"`

	// 回源请求头部配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RequestHeader *RequestHeader `json:"RequestHeader,omitempty" name:"RequestHeader"`

	// 源站响应头部配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResponseHeader *ResponseHeader `json:"ResponseHeader,omitempty" name:"ResponseHeader"`

	// 遵循源站缓存头部配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResponseHeaderCache *ResponseHeaderCache `json:"ResponseHeaderCache,omitempty" name:"ResponseHeaderCache"`

	// seo优化配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Seo *Seo `json:"Seo,omitempty" name:"Seo"`

	// 域名业务类型，web，download，media分别表示静态加速，下载加速和流媒体加速。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 状态码缓存配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusCodeCache *StatusCodeCache `json:"StatusCodeCache,omitempty" name:"StatusCodeCache"`

	// 视频拖拽配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoSeek *VideoSeek `json:"VideoSeek,omitempty" name:"VideoSeek"`
}

type ManageClsTopicDomainsRequest struct {
	*tchttp.BaseRequest

	// 日志集ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// 日志主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 接入渠道，默认值为cdn
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// 域名区域配置，注意：如果此字段为空，则表示解绑对应主题下的所有域名
	DomainAreaConfigs []*DomainAreaConfig `json:"DomainAreaConfigs,omitempty" name:"DomainAreaConfigs" list`
}

func (r *ManageClsTopicDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ManageClsTopicDomainsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ManageClsTopicDomainsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ManageClsTopicDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ManageClsTopicDomainsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MapInfo struct {

	// 对象 Id
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 对象名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

type MaxAge struct {

	// 浏览器缓存配置开关
	// on：开启
	// off：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// MaxAge 规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxAgeRules []*MaxAgeRule `json:"MaxAgeRules,omitempty" name:"MaxAgeRules" list`
}

type MaxAgeRule struct {

	// 规则类型：
	// all：所有文件生效
	// file：指定文件后缀生效
	// directory：指定路径生效
	// path：指定绝对路径生效
	MaxAgeType *string `json:"MaxAgeType,omitempty" name:"MaxAgeType"`

	// MaxAgeType 对应类型下的匹配内容：
	// all 时填充 *
	// file 时填充后缀名，如 jpg、txt
	// directory 时填充路径，如 /xxx/test/
	// path 时填充绝对路径，如 /xxx/test.html
	MaxAgeContents []*string `json:"MaxAgeContents,omitempty" name:"MaxAgeContents" list`

	// MaxAge 时间设置，单位秒
	MaxAgeTime *int64 `json:"MaxAgeTime,omitempty" name:"MaxAgeTime"`
}

type Origin struct {

	// 主源站列表
	// 修改源站时，需要同时填充对应的 OriginType
	// 注意：此字段可能返回 null，表示取不到有效值。
	Origins []*string `json:"Origins,omitempty" name:"Origins" list`

	// 主源站类型
	// 入参支持以下几种类型：
	// domain：域名类型
	// cos：对象存储源站
	// ip：IP 列表作为源站
	// ipv6：源站列表为一个单独的 IPv6 地址
	// ip_ipv6：源站列表为多个 IPv4 地址和一个 IPv6 地址
	// 出参增加以下几种类型：
	// image：数据万象源站
	// ftp：历史 FTP 托管源源站，现已不维护
	// 修改 Origins 时需要同时填充对应的 OriginType
	// IPv6 功能目前尚未全量，需要先申请试用
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// 回主源站时 Host 头部，不填充则默认为加速域名
	// 若接入的是泛域名，则回源 Host 默认为访问时的子域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// OriginType 为对象存储（COS）时，可以指定是否允许访问私有 bucket
	// 注意：需要先授权 CDN 访问该私有 Bucket 的权限后，才可开启此配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosPrivateAccess *string `json:"CosPrivateAccess,omitempty" name:"CosPrivateAccess"`

	// 回源协议配置
	// http：强制 http 回源
	// follow：协议跟随回源
	// https：强制 https 回源，https 回源时仅支持源站 443 端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginPullProtocol *string `json:"OriginPullProtocol,omitempty" name:"OriginPullProtocol"`

	// 备源站列表
	// 修改备源站时，需要同时填充对应的 BackupOriginType
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupOrigins []*string `json:"BackupOrigins,omitempty" name:"BackupOrigins" list`

	// 备源站类型，支持以下类型：
	// domain：域名类型
	// ip：IP 列表作为源站
	// 修改 BackupOrigins 时需要同时填充对应的 BackupOriginType
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupOriginType *string `json:"BackupOriginType,omitempty" name:"BackupOriginType"`

	// 回备源站时 Host 头部，不填充则默认为主源站的 ServerName
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupServerName *string `json:"BackupServerName,omitempty" name:"BackupServerName"`
}

type OriginPullOptimization struct {

	// 跨国回源优化配置开关
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 跨国类型
	// OVToCN：境外回源境内
	// CNToOV：境内回源境外
	// 注意：此字段可能返回 null，表示取不到有效值。
	OptimizationType *string `json:"OptimizationType,omitempty" name:"OptimizationType"`
}

type OriginPullTimeout struct {

	// 回源建连超时时间，单位为秒，要求5~60之间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConnectTimeout *uint64 `json:"ConnectTimeout,omitempty" name:"ConnectTimeout"`

	// 回源接收超时时间，单位为秒，要求10 ~ 60之间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReceiveTimeout *uint64 `json:"ReceiveTimeout,omitempty" name:"ReceiveTimeout"`
}

type OverseaConfig struct {

	// 时间戳防盗链配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Authentication *Authentication `json:"Authentication,omitempty" name:"Authentication"`

	// 带宽封顶配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BandwidthAlert *BandwidthAlert `json:"BandwidthAlert,omitempty" name:"BandwidthAlert"`

	// 缓存规则配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cache *Cache `json:"Cache,omitempty" name:"Cache"`

	// 缓存相关配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheKey *CacheKey `json:"CacheKey,omitempty" name:"CacheKey"`

	// 智能压缩配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Compression *Compression `json:"Compression,omitempty" name:"Compression"`

	// 下载限速配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownstreamCapping *DownstreamCapping `json:"DownstreamCapping,omitempty" name:"DownstreamCapping"`

	// 错误码重定向配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorPage *ErrorPage `json:"ErrorPage,omitempty" name:"ErrorPage"`

	// 301和302自动回源跟随配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FollowRedirect *FollowRedirect `json:"FollowRedirect,omitempty" name:"FollowRedirect"`

	// 访问协议强制跳转配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitempty" name:"ForceRedirect"`

	// Https配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Https *Https `json:"Https,omitempty" name:"Https"`

	// IP黑白名单配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpFilter *IpFilter `json:"IpFilter,omitempty" name:"IpFilter"`

	// IP限频配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpFreqLimit *IpFreqLimit `json:"IpFreqLimit,omitempty" name:"IpFreqLimit"`

	// 浏览器缓存规则配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxAge *MaxAge `json:"MaxAge,omitempty" name:"MaxAge"`

	// 源站配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Origin *Origin `json:"Origin,omitempty" name:"Origin"`

	// 跨国优化配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginPullOptimization *OriginPullOptimization `json:"OriginPullOptimization,omitempty" name:"OriginPullOptimization"`

	// Range回源配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RangeOriginPull *RangeOriginPull `json:"RangeOriginPull,omitempty" name:"RangeOriginPull"`

	// 防盗链配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Referer *Referer `json:"Referer,omitempty" name:"Referer"`

	// 回源请求头部配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RequestHeader *RequestHeader `json:"RequestHeader,omitempty" name:"RequestHeader"`

	// 源站响应头部配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResponseHeader *ResponseHeader `json:"ResponseHeader,omitempty" name:"ResponseHeader"`

	// 遵循源站缓存头部配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResponseHeaderCache *ResponseHeaderCache `json:"ResponseHeaderCache,omitempty" name:"ResponseHeaderCache"`

	// seo优化配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Seo *Seo `json:"Seo,omitempty" name:"Seo"`

	// 域名业务类型，web，download，media分别表示静态加速，下载加速和流媒体加速。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 状态码缓存配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusCodeCache *StatusCodeCache `json:"StatusCodeCache,omitempty" name:"StatusCodeCache"`

	// 视频拖拽配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoSeek *VideoSeek `json:"VideoSeek,omitempty" name:"VideoSeek"`
}

type PurgePathCacheRequest struct {
	*tchttp.BaseRequest

	// 目录列表，需要包含协议头部 http:// 或 https://
	Paths []*string `json:"Paths,omitempty" name:"Paths" list`

	// 刷新类型
	// flush：刷新产生更新的资源
	// delete：刷新全部资源
	FlushType *string `json:"FlushType,omitempty" name:"FlushType"`
}

func (r *PurgePathCacheRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PurgePathCacheRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PurgePathCacheResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 刷新任务 ID，同一批次提交的目录共用一个任务 ID
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PurgePathCacheResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PurgePathCacheResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PurgeTask struct {

	// 刷新任务 ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 刷新 URL
	Url *string `json:"Url,omitempty" name:"Url"`

	// 刷新任务状态
	// fail：刷新失败
	// done：刷新成功
	// process：刷新中
	Status *string `json:"Status,omitempty" name:"Status"`

	// 刷新类型
	// url：URL 刷新
	// path：目录刷新
	PurgeType *string `json:"PurgeType,omitempty" name:"PurgeType"`

	// 刷新方式
	// flush：刷新更新资源（仅目录刷新时有此类型）
	// delete：刷新全部资源
	FlushType *string `json:"FlushType,omitempty" name:"FlushType"`

	// 刷新任务提交时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type PurgeUrlsCacheRequest struct {
	*tchttp.BaseRequest

	// URL 列表，需要包含协议头部 http:// 或 https://
	Urls []*string `json:"Urls,omitempty" name:"Urls" list`
}

func (r *PurgeUrlsCacheRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PurgeUrlsCacheRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PurgeUrlsCacheResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 刷新任务 ID，同一批次提交的 URL 共用一个任务 ID
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PurgeUrlsCacheResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PurgeUrlsCacheResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PushTask struct {

	// 预热任务 ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 预热 URL
	Url *string `json:"Url,omitempty" name:"Url"`

	// 预热任务状态
	// fail：预热失败
	// done：预热成功
	// process：预热中
	Status *string `json:"Status,omitempty" name:"Status"`

	// 预热进度百分比
	Percent *int64 `json:"Percent,omitempty" name:"Percent"`

	// 预热任务提交时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 预热区域
	// mainland：境内
	// overseas：境外
	// global：全球
	Area *string `json:"Area,omitempty" name:"Area"`

	// 预热任务更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type PushUrlsCacheRequest struct {
	*tchttp.BaseRequest

	// URL 列表，需要包含协议头部 http:// 或 https://
	Urls []*string `json:"Urls,omitempty" name:"Urls" list`

	// 指定预热请求回源时 HTTP 请求的 User-Agent 头部
	// 默认为 TencentCdn
	UserAgent *string `json:"UserAgent,omitempty" name:"UserAgent"`

	// 预热生效区域
	// mainland：预热至境内节点
	// overseas：预热至境外节点
	// global：预热全球节点
	// 不填充情况下，默认为 mainland， URL 中域名必须在对应区域启用了加速服务才能提交对应区域的预热任务
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *PushUrlsCacheRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PushUrlsCacheRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PushUrlsCacheResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 此批提交的任务 ID
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PushUrlsCacheResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PushUrlsCacheResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryStringKey struct {

	// on | off CacheKey是否由QueryString组成
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 是否重新排序
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reorder *string `json:"Reorder,omitempty" name:"Reorder"`

	// includeAll | excludeAll | includeCustom | excludeAll 使用/排除部分url参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Action *string `json:"Action,omitempty" name:"Action"`

	// 使用/排除的url参数名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type Quota struct {

	// 单次批量提交配额上限。
	Batch *int64 `json:"Batch,omitempty" name:"Batch"`

	// 每日提交配额上限。
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 每日剩余的可提交配额。
	Available *int64 `json:"Available,omitempty" name:"Available"`

	// 配额的区域。
	Area *string `json:"Area,omitempty" name:"Area"`
}

type RangeOriginPull struct {

	// 分片回源配置开关
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type Referer struct {

	// referer 黑白名单配置开关
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// referer 黑白名单配置规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	RefererRules []*RefererRule `json:"RefererRules,omitempty" name:"RefererRules" list`
}

type RefererRule struct {

	// 规则类型：
	// all：所有文件生效
	// file：指定文件后缀生效
	// directory：指定路径生效
	// path：指定绝对路径生效
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`

	// RuleType 对应类型下的匹配内容：
	// all 时填充 *
	// file 时填充后缀名，如 jpg、txt
	// directory 时填充路径，如 /xxx/test/
	// path 时填充绝对路径，如 /xxx/test.html
	RulePaths []*string `json:"RulePaths,omitempty" name:"RulePaths" list`

	// referer 配置类型
	// whitelist：白名单
	// blacklist：黑名单
	RefererType *string `json:"RefererType,omitempty" name:"RefererType"`

	// referer 内容列表列表
	Referers []*string `json:"Referers,omitempty" name:"Referers" list`

	// 是否允许空 referer
	// true：允许空 referer
	// false：不允许空 referer
	AllowEmpty *bool `json:"AllowEmpty,omitempty" name:"AllowEmpty"`
}

type RegionMapRelation struct {

	// 区域ID。
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// 子区域ID列表
	SubRegionIdList []*int64 `json:"SubRegionIdList,omitempty" name:"SubRegionIdList" list`
}

type ReportData struct {

	// 项目ID/域名ID。
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 项目名称/域名。
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// 流量总和/带宽最大值，单位分别为bytes，bps。
	Value *int64 `json:"Value,omitempty" name:"Value"`

	// 单个资源占总体百分比。
	Percentage *float64 `json:"Percentage,omitempty" name:"Percentage"`

	// 计费流量总和/计费带宽最大值，单位分别为bytes，bps。
	BillingValue *int64 `json:"BillingValue,omitempty" name:"BillingValue"`

	// 计费数值占总体百分比。
	BillingPercentage *float64 `json:"BillingPercentage,omitempty" name:"BillingPercentage"`
}

type RequestHeader struct {

	// 自定义请求头配置开关
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 自定义请求头配置规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeaderRules []*HttpHeaderPathRule `json:"HeaderRules,omitempty" name:"HeaderRules" list`
}

type ResourceBillingData struct {

	// 资源名称，根据查询条件不同分为以下几类：
	// 某一个具体域名：表示该域名明细数据
	// multiDomains：表示多域名汇总明细数据
	// 某一个项目 ID：指定项目查询时，显示为项目 ID
	// all：账号维度数据明细
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// 计费数据详情
	BillingData []*CdnData `json:"BillingData,omitempty" name:"BillingData" list`
}

type ResourceData struct {

	// 资源名称，根据查询条件不同分为以下几类：
	// 具体域名：表示该域名明细数据
	// multiDomains：表示多域名汇总明细数据
	// 项目 ID：指定项目查询时，显示为项目 ID
	// all：账号维度明细数据
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// 资源对应的数据明细
	CdnData []*CdnData `json:"CdnData,omitempty" name:"CdnData" list`
}

type ResourceOriginData struct {

	// 资源名称，根据查询条件不同分为以下几类：
	// 具体域名：表示该域名明细数据
	// multiDomains：表示多域名汇总明细数据
	// 项目 ID：指定项目查询时，显示为项目 ID
	// all：账号维度明细数据
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// 回源数据详情
	OriginData []*CdnData `json:"OriginData,omitempty" name:"OriginData" list`
}

type ResponseHeader struct {

	// 自定义响应头开关
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 自定义响应头规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeaderRules []*HttpHeaderPathRule `json:"HeaderRules,omitempty" name:"HeaderRules" list`
}

type ResponseHeaderCache struct {

	// 源站头部缓存开关
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type SchemeKey struct {

	// on | off 是否使用scheme作为cache key的一部分
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type SearchClsLogRequest struct {
	*tchttp.BaseRequest

	// 需要查询的日志集ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// 需要查询的日志主题ID组合，以逗号分隔
	TopicIds *string `json:"TopicIds,omitempty" name:"TopicIds"`

	// 需要查询的日志的起始时间，格式 YYYY-mm-dd HH:MM:SS
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 需要查询的日志的结束时间，格式 YYYY-mm-dd HH:MM:SS
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 单次要返回的日志条数，单次返回的最大条数为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 接入渠道，默认值为cdn
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// 需要查询的内容，详情请参考https://cloud.tencent.com/document/product/614/16982
	Query *string `json:"Query,omitempty" name:"Query"`

	// 加载更多使用，透传上次返回的 context 值，获取后续的日志内容，通过游标最多可获取10000条，请尽可能缩小时间范围
	Context *string `json:"Context,omitempty" name:"Context"`

	// 按日志时间排序， asc（升序）或者 desc（降序），默认为 desc
	Sort *string `json:"Sort,omitempty" name:"Sort"`
}

func (r *SearchClsLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchClsLogRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchClsLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询结果
		Logs *ClsSearchLogs `json:"Logs,omitempty" name:"Logs"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchClsLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchClsLogResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SecurityConfig struct {

	// on|off
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type Seo struct {

	// SEO 配置开关
	// on：开启
	// off：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type ServerCert struct {

	// 服务器证书 ID
	// 在 SSL 证书管理进行证书托管时自动生成
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// 服务器证书名称
	// 在 SSL 证书管理进行证书托管时自动生成
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertName *string `json:"CertName,omitempty" name:"CertName"`

	// 服务器证书信息
	// 上传自有证书时必填，需要包含完整的证书链
	// 注意：此字段可能返回 null，表示取不到有效值。
	Certificate *string `json:"Certificate,omitempty" name:"Certificate"`

	// 服务器密钥信息
	// 上传自有证书时必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`

	// 证书过期时间
	// 作为入参配置时无需填充
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 证书颁发时间
	// 作为入参配置时无需填充
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployTime *string `json:"DeployTime,omitempty" name:"DeployTime"`

	// 证书备注信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`
}

type SimpleCache struct {

	// 缓存过期时间规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheRules []*SimpleCacheRule `json:"CacheRules,omitempty" name:"CacheRules" list`

	// 遵循源站 Cache-Control: max-age 配置
	// on：开启
	// off：关闭
	// 开启后，未能匹配 CacheRules 规则的资源将根据源站返回的 max-age 值进行节点缓存；匹配了 CacheRules 规则的资源将按照 CacheRules 中设置的缓存过期时间在节点进行缓存
	// 与 CompareMaxAge 冲突，不能同时开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	FollowOrigin *string `json:"FollowOrigin,omitempty" name:"FollowOrigin"`

	// 强制缓存
	// on：开启
	// off：关闭
	// 默认为关闭状态，开启后，源站发挥的 no-store、no-cache 资源，也将按照 CacheRules 规则进行缓存
	// 注意：此字段可能返回 null，表示取不到有效值。
	IgnoreCacheControl *string `json:"IgnoreCacheControl,omitempty" name:"IgnoreCacheControl"`

	// 忽略源站的Set-Cookie头部
	// on：开启
	// off：关闭
	// 默认为关闭状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	IgnoreSetCookie *string `json:"IgnoreSetCookie,omitempty" name:"IgnoreSetCookie"`

	// 高级缓存过期配置，开启时会对比源站返回的 max-age 值与 CacheRules 中设置的缓存过期时间，取最小值在节点进行缓存
	// on：开启
	// off：关闭
	// 默认为关闭状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompareMaxAge *string `json:"CompareMaxAge,omitempty" name:"CompareMaxAge"`
}

type SimpleCacheRule struct {

	// 规则类型：
	// all：所有文件生效
	// file：指定文件后缀生效
	// directory：指定路径生效
	// path：指定绝对路径生效
	// index：首页
	CacheType *string `json:"CacheType,omitempty" name:"CacheType"`

	// CacheType 对应类型下的匹配内容：
	// all 时填充 *
	// file 时填充后缀名，如 jpg、txt
	// directory 时填充路径，如 /xxx/test/
	// path 时填充绝对路径，如 /xxx/test.html
	// index 时填充 /
	CacheContents []*string `json:"CacheContents,omitempty" name:"CacheContents" list`

	// 缓存过期时间设置
	// 单位为秒，最大可设置为 365 天
	CacheTime *int64 `json:"CacheTime,omitempty" name:"CacheTime"`
}

type Sort struct {

	// 排序字段，当前支持：
	// createTime，域名创建时间
	// certExpireTime，证书过期时间
	// 默认createTime。
	Key *string `json:"Key,omitempty" name:"Key"`

	// asc/desc，默认desc。
	Sequence *string `json:"Sequence,omitempty" name:"Sequence"`
}

type SpecificConfig struct {

	// 国内特殊配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mainland *MainlandConfig `json:"Mainland,omitempty" name:"Mainland"`

	// 海外特殊配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Overseas *OverseaConfig `json:"Overseas,omitempty" name:"Overseas"`
}

type StartCdnDomainRequest struct {
	*tchttp.BaseRequest

	// 域名
	// 域名状态需要为【已停用】
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *StartCdnDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartCdnDomainRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartCdnDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartCdnDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartCdnDomainResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StatusCodeCache struct {

	// 状态码缓存过期配置开关
	// on：开启
	// off：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 状态码缓存过期规则明细
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheRules []*StatusCodeCacheRule `json:"CacheRules,omitempty" name:"CacheRules" list`
}

type StatusCodeCacheRule struct {

	// http 状态码
	// 支持 403、404 状态码
	StatusCode *string `json:"StatusCode,omitempty" name:"StatusCode"`

	// 状态码缓存过期时间，单位秒
	CacheTime *int64 `json:"CacheTime,omitempty" name:"CacheTime"`
}

type StopCdnDomainRequest struct {
	*tchttp.BaseRequest

	// 域名
	// 域名需要为【已启动】状态
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *StopCdnDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopCdnDomainRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopCdnDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopCdnDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopCdnDomainResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SummarizedData struct {

	// 汇总方式，存在以下几种：
	// sum：累加求和
	// max：最大值，带宽模式下，采用 5 分钟粒度汇总数据，计算峰值带宽
	// avg：平均值
	Name *string `json:"Name,omitempty" name:"Name"`

	// 汇总后的数据值
	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type TimestampData struct {

	// 数据统计时间点，采用向前汇总模式
	// 以 5 分钟粒度为例，13:35:00 时间点代表的统计数据区间为 13:35:00 至 13:39:59
	Time *string `json:"Time,omitempty" name:"Time"`

	// 数据值
	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type TopData struct {

	// 资源名称，根据查询条件不同分为以下几类：
	// 具体域名：表示该域名明细数据
	// multiDomains：表示多域名汇总明细数据
	// 项目 ID：指定项目查询时，显示为项目 ID
	// all：账号维度明细数据
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// 排序结果详情
	DetailData []*TopDetailData `json:"DetailData,omitempty" name:"DetailData" list`
}

type TopDetailData struct {

	// 数据类型的名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 数据值
	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type TopicInfo struct {

	// 主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 主题名字
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 是否启用投递
	Enabled *int64 `json:"Enabled,omitempty" name:"Enabled"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type TpgAdapter struct {

	// 开关，"on/off"
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type TrafficPackage struct {

	// 流量包 Id
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 流量包类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 流量包大小（单位为 Byte）
	Bytes *int64 `json:"Bytes,omitempty" name:"Bytes"`

	// 已消耗流量（单位为 Byte）
	BytesUsed *int64 `json:"BytesUsed,omitempty" name:"BytesUsed"`

	// 流量包状态
	// enabled：已启用
	// expired：已过期
	// disabled：未启用
	Status *string `json:"Status,omitempty" name:"Status"`

	// 流量包发放时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 流量包生效时间
	EnableTime *string `json:"EnableTime,omitempty" name:"EnableTime"`

	// 流量包过期时间
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 流量包是否续订
	ContractExtension *bool `json:"ContractExtension,omitempty" name:"ContractExtension"`

	// 流量包是否自动续订
	AutoExtension *bool `json:"AutoExtension,omitempty" name:"AutoExtension"`

	// 流量包来源
	Channel *string `json:"Channel,omitempty" name:"Channel"`
}

type UpdateDomainConfigRequest struct {
	*tchttp.BaseRequest

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 项目 ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 源站配置
	Origin *Origin `json:"Origin,omitempty" name:"Origin"`

	// IP 黑白名单配置
	IpFilter *IpFilter `json:"IpFilter,omitempty" name:"IpFilter"`

	// IP 限频配置
	IpFreqLimit *IpFreqLimit `json:"IpFreqLimit,omitempty" name:"IpFreqLimit"`

	// 状态码缓存配置
	StatusCodeCache *StatusCodeCache `json:"StatusCodeCache,omitempty" name:"StatusCodeCache"`

	// 智能压缩配置
	Compression *Compression `json:"Compression,omitempty" name:"Compression"`

	// 带宽封顶配置
	BandwidthAlert *BandwidthAlert `json:"BandwidthAlert,omitempty" name:"BandwidthAlert"`

	// Range 回源配置
	RangeOriginPull *RangeOriginPull `json:"RangeOriginPull,omitempty" name:"RangeOriginPull"`

	// 301/302 回源跟随配置
	FollowRedirect *FollowRedirect `json:"FollowRedirect,omitempty" name:"FollowRedirect"`

	// 错误码重定向配置（功能灰度中，尚未全量）
	ErrorPage *ErrorPage `json:"ErrorPage,omitempty" name:"ErrorPage"`

	// 请求头部配置
	RequestHeader *RequestHeader `json:"RequestHeader,omitempty" name:"RequestHeader"`

	// 响应头部配置
	ResponseHeader *ResponseHeader `json:"ResponseHeader,omitempty" name:"ResponseHeader"`

	// 下载速度配置
	DownstreamCapping *DownstreamCapping `json:"DownstreamCapping,omitempty" name:"DownstreamCapping"`

	// 节点缓存键配置
	CacheKey *CacheKey `json:"CacheKey,omitempty" name:"CacheKey"`

	// 头部缓存配置
	ResponseHeaderCache *ResponseHeaderCache `json:"ResponseHeaderCache,omitempty" name:"ResponseHeaderCache"`

	// 视频拖拽配置
	VideoSeek *VideoSeek `json:"VideoSeek,omitempty" name:"VideoSeek"`

	// 缓存过期时间配置
	Cache *Cache `json:"Cache,omitempty" name:"Cache"`

	// 跨国链路优化配置
	OriginPullOptimization *OriginPullOptimization `json:"OriginPullOptimization,omitempty" name:"OriginPullOptimization"`

	// Https 加速配置
	Https *Https `json:"Https,omitempty" name:"Https"`

	// 时间戳防盗链配置
	Authentication *Authentication `json:"Authentication,omitempty" name:"Authentication"`

	// SEO 优化配置
	Seo *Seo `json:"Seo,omitempty" name:"Seo"`

	// 访问协议强制跳转配置
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitempty" name:"ForceRedirect"`

	// Referer 防盗链配置
	Referer *Referer `json:"Referer,omitempty" name:"Referer"`

	// 浏览器缓存配置（功能灰度中，尚未全量）
	MaxAge *MaxAge `json:"MaxAge,omitempty" name:"MaxAge"`

	// 域名业务类型
	// web：静态加速
	// download：下载加速
	// media：流媒体点播加速
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 地域属性特殊配置
	// 适用于域名境内加速、境外加速配置不一致场景
	SpecificConfig *SpecificConfig `json:"SpecificConfig,omitempty" name:"SpecificConfig"`

	// 域名加速区域
	// mainland：中国境内加速
	// overseas：中国境外加速
	// global：全球加速
	Area *string `json:"Area,omitempty" name:"Area"`

	// 回源超时配置
	OriginPullTimeout *OriginPullTimeout `json:"OriginPullTimeout,omitempty" name:"OriginPullTimeout"`

	// 回源S3私有鉴权
	AwsPrivateAccess *AwsPrivateAccess `json:"AwsPrivateAccess,omitempty" name:"AwsPrivateAccess"`
}

func (r *UpdateDomainConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateDomainConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateDomainConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateDomainConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateDomainConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateImageConfigRequest struct {
	*tchttp.BaseRequest

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// WebpAdapter配置项
	WebpAdapter *WebpAdapter `json:"WebpAdapter,omitempty" name:"WebpAdapter"`

	// TpgAdapter配置项
	TpgAdapter *TpgAdapter `json:"TpgAdapter,omitempty" name:"TpgAdapter"`

	// GuetzliAdapter配置项
	GuetzliAdapter *GuetzliAdapter `json:"GuetzliAdapter,omitempty" name:"GuetzliAdapter"`
}

func (r *UpdateImageConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateImageConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateImageConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateImageConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateImageConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdatePayTypeRequest struct {
	*tchttp.BaseRequest

	// 计费区域，mainland或overseas。
	Area *string `json:"Area,omitempty" name:"Area"`

	// 计费类型，flux或bandwidth。
	PayType *string `json:"PayType,omitempty" name:"PayType"`
}

func (r *UpdatePayTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdatePayTypeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdatePayTypeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePayTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdatePayTypeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UrlRecord struct {

	// 状态(disable表示封禁，enable表示解封)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 对应的url
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealUrl *string `json:"RealUrl,omitempty" name:"RealUrl"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type UserAgentFilter struct {

	// 开关，on或off
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// UA黑白名单生效规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilterRules []*UserAgentFilterRule `json:"FilterRules,omitempty" name:"FilterRules" list`
}

type UserAgentFilterRule struct {

	// 访问路径生效类型
	// all: 所有访问路径生效
	// file: 根据文件后缀类型生效
	// directory: 根据目录生效
	// path: 根据完整访问路径生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`

	// 访问路径生效内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	RulePaths []*string `json:"RulePaths,omitempty" name:"RulePaths" list`

	// UserAgent列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserAgents []*string `json:"UserAgents,omitempty" name:"UserAgents" list`

	// 黑名单或白名单，blacklist或whitelist
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilterType *string `json:"FilterType,omitempty" name:"FilterType"`
}

type VideoSeek struct {

	// 视频拖拽开关
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type ViolationUrl struct {

	// ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 违规资源原始访问 URL
	RealUrl *string `json:"RealUrl,omitempty" name:"RealUrl"`

	// 快照路径，用于控制台展示违规内容快照
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 违规资源当前状态
	// forbid：已封禁
	// release：已解封
	// delay ： 延迟处理
	// reject ：申诉驳回，状态仍为封禁态
	// complain：申诉进行中
	UrlStatus *string `json:"UrlStatus,omitempty" name:"UrlStatus"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type WebpAdapter struct {

	// 开关，"on/off"
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}
