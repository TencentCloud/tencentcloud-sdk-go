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

package v20220726

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AccountInfo struct {
	// 用户账号类型：
	// 1-手机号；
	// 2-IMEI；
	// 3-IDFA；
	// 100-SSID类型
	AccountType *int64 `json:"AccountType,omitempty" name:"AccountType"`

	// 通用账号信息，当AccountType为1、2、3、100时必填
	UniversalAccount *UniversalAccountInfo `json:"UniversalAccount,omitempty" name:"UniversalAccount"`
}

type DeviceDetailInfo struct {
	// mac地址或唯一设备标识
	MacAddress *string `json:"MacAddress,omitempty" name:"MacAddress"`

	// 手机型号
	Model *string `json:"Model,omitempty" name:"Model"`

	// 操作系统(unknown，android，ios，windows)
	OsSystem *string `json:"OsSystem,omitempty" name:"OsSystem"`

	// 操作系统版本
	OsSystemVersion *string `json:"OsSystemVersion,omitempty" name:"OsSystemVersion"`

	// 竞价底价
	BidFloor *int64 `json:"BidFloor,omitempty" name:"BidFloor"`

	// 设备版本
	DeviceVersion *string `json:"DeviceVersion,omitempty" name:"DeviceVersion"`

	// 设备制造商
	Maker *string `json:"Maker,omitempty" name:"Maker"`

	// 设备类型（PHONE,TABLET）
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// 运营商；-1: 获取失败，0: 其他，1: 移动，2: 联通，3: 电信，4: 铁通
	Carrier *string `json:"Carrier,omitempty" name:"Carrier"`

	// 入网方式(wifi,5g,4g,3g,2g)
	AccessMode *string `json:"AccessMode,omitempty" name:"AccessMode"`

	// 手机芯片信息
	PhoneChipInfo *string `json:"PhoneChipInfo,omitempty" name:"PhoneChipInfo"`

	// CPU 型号
	CpuModel *string `json:"CpuModel,omitempty" name:"CpuModel"`

	// CPU 核数
	CpuCore *string `json:"CpuCore,omitempty" name:"CpuCore"`

	// 内存容量，单位转换为 GB
	Memory *string `json:"Memory,omitempty" name:"Memory"`

	// 系统语言
	Language *string `json:"Language,omitempty" name:"Language"`

	// 手机音量
	Volume *string `json:"Volume,omitempty" name:"Volume"`

	// 电池电量
	BatteryPower *string `json:"BatteryPower,omitempty" name:"BatteryPower"`

	// 屏幕分辨率宽，保留整数
	ResolutionWidth *int64 `json:"ResolutionWidth,omitempty" name:"ResolutionWidth"`

	// 屏幕分辨率高，保留整数
	ResolutionHeight *int64 `json:"ResolutionHeight,omitempty" name:"ResolutionHeight"`

	// 浏览器类型
	Ua *string `json:"Ua,omitempty" name:"Ua"`

	// 客户端应用
	App *string `json:"App,omitempty" name:"App"`

	// 应用包名
	AppPackageName *string `json:"AppPackageName,omitempty" name:"AppPackageName"`

	// 设备序列号
	// Android设备
	SerialNumber *string `json:"SerialNumber,omitempty" name:"SerialNumber"`

	// netOperator MCC+MNC
	// Android设备
	MobileCountryAndNetworkCode *string `json:"MobileCountryAndNetworkCode,omitempty" name:"MobileCountryAndNetworkCode"`

	// 设备品牌 “华为”“oppo”“小米”
	// Android设备
	VendorId *string `json:"VendorId,omitempty" name:"VendorId"`

	// Android API等级
	// Android设备
	AndroidApiLevel *string `json:"AndroidApiLevel,omitempty" name:"AndroidApiLevel"`

	// 屏幕亮度
	// Android设备
	Brightness *string `json:"Brightness,omitempty" name:"Brightness"`

	// 蓝牙地址
	// Android设备
	BluetoothAddress *string `json:"BluetoothAddress,omitempty" name:"BluetoothAddress"`

	// 基带版本
	// Android设备
	BaseBandVersion *string `json:"BaseBandVersion,omitempty" name:"BaseBandVersion"`

	// kernel 版本
	// Android设备
	KernelVersion *string `json:"KernelVersion,omitempty" name:"KernelVersion"`

	// 存储容量，单位转换为 GB
	// Android设备
	Storage *string `json:"Storage,omitempty" name:"Storage"`

	// 软件包名
	// Android设备
	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`

	// app 版本号
	// Android设备
	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`

	// app 显示名称
	// Android设备
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 是否 debug；0 为正常模式，1 为 debug 模式；其他值无效
	// Android设备
	IsDebug *string `json:"IsDebug,omitempty" name:"IsDebug"`

	// 是否越狱；0 为正常，1 为越狱；其他值无效
	// Android设备
	IsRoot *string `json:"IsRoot,omitempty" name:"IsRoot"`

	// 是否启动代理；0 为未开启，1 为开启；其他值无效
	// Android设备
	IsProxy *string `json:"IsProxy,omitempty" name:"IsProxy"`

	// 是否模拟器；0 为未开启，1 为开启；其他值无效
	// Android设备
	IsEmulator *string `json:"IsEmulator,omitempty" name:"IsEmulator"`

	// 充电状态；1-不在充电，2-USB充电，3-电源充电
	// Android设备
	ChargeStatus *string `json:"ChargeStatus,omitempty" name:"ChargeStatus"`

	// 网络类型：2G/3G/4G/5G/Wi-Fi/WWAN/other
	// Android设备
	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`

	// Wi-Fi MAC地址
	// Android设备
	WifiMac *string `json:"WifiMac,omitempty" name:"WifiMac"`

	// 设备名称 "xxx 的 iPhone"，"xxx's IPhone" 等等
	// IOS设备
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 开机时间
	// IOS设备
	StartupTime *string `json:"StartupTime,omitempty" name:"StartupTime"`

	// 所在经度
	Lon *string `json:"Lon,omitempty" name:"Lon"`

	// 所在纬度
	Lat *string `json:"Lat,omitempty" name:"Lat"`
}

type DeviceFingerprintInfo struct {
	// 设备指纹Token
	DeviceToken *string `json:"DeviceToken,omitempty" name:"DeviceToken"`

	// 设备指纹的客户端SDK对应渠道
	SdkChannel *string `json:"SdkChannel,omitempty" name:"SdkChannel"`
}

// Predefined struct for user
type EvaluateUserRiskRequestParams struct {
	// 账号信息
	Account *AccountInfo `json:"Account,omitempty" name:"Account"`

	// 用户信息
	User *UserInfo `json:"User,omitempty" name:"User"`

	// 模型ID
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 设备指纹信息
	DeviceFingerprint *DeviceFingerprintInfo `json:"DeviceFingerprint,omitempty" name:"DeviceFingerprint"`

	// 场景Code，不传默认活动防刷；
	// e_activity_antirush；活动防刷场景
	// e_login_protection；登录保护场景
	// e_register_protection：注册保护场景
	SceneCode *string `json:"SceneCode,omitempty" name:"SceneCode"`

	// 设备详情
	DeviceDetail *DeviceDetailInfo `json:"DeviceDetail,omitempty" name:"DeviceDetail"`

	// 营销信息
	Marketing *MarketingInfo `json:"Marketing,omitempty" name:"Marketing"`
}

type EvaluateUserRiskRequest struct {
	*tchttp.BaseRequest
	
	// 账号信息
	Account *AccountInfo `json:"Account,omitempty" name:"Account"`

	// 用户信息
	User *UserInfo `json:"User,omitempty" name:"User"`

	// 模型ID
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 设备指纹信息
	DeviceFingerprint *DeviceFingerprintInfo `json:"DeviceFingerprint,omitempty" name:"DeviceFingerprint"`

	// 场景Code，不传默认活动防刷；
	// e_activity_antirush；活动防刷场景
	// e_login_protection；登录保护场景
	// e_register_protection：注册保护场景
	SceneCode *string `json:"SceneCode,omitempty" name:"SceneCode"`

	// 设备详情
	DeviceDetail *DeviceDetailInfo `json:"DeviceDetail,omitempty" name:"DeviceDetail"`

	// 营销信息
	Marketing *MarketingInfo `json:"Marketing,omitempty" name:"Marketing"`
}

func (r *EvaluateUserRiskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EvaluateUserRiskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Account")
	delete(f, "User")
	delete(f, "ModelId")
	delete(f, "DeviceFingerprint")
	delete(f, "SceneCode")
	delete(f, "DeviceDetail")
	delete(f, "Marketing")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EvaluateUserRiskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EvaluateUserRiskResponseParams struct {
	// 评估结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	EvaluationResult *EvaluationResult `json:"EvaluationResult,omitempty" name:"EvaluationResult"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type EvaluateUserRiskResponse struct {
	*tchttp.BaseResponse
	Response *EvaluateUserRiskResponseParams `json:"Response"`
}

func (r *EvaluateUserRiskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EvaluateUserRiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EvaluationResult struct {
	// SSID值
	SSID *string `json:"SSID,omitempty" name:"SSID"`

	// 风险价值分
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *float64 `json:"Score,omitempty" name:"Score"`

	// 风险标签，请参考官网风险类型
	// 账号风险
	// 1 账号信用低 账号近期存在因恶意被处罚历史，网络低活跃，被举报等因素
	// 11 疑似低活跃账号 账号活跃度与正常用户有差异
	// 2 垃圾账号 疑似批量注册小号，近期存在严重违规或大量举报
	// 21 疑似小号 账号有疑似线上养号，小号等行为
	// 22 疑似违规账号 账号曾有违规行为、曾被举报过、曾因违规被处罚过等
	// 3 无效账号 送检账号参数无法成功解析，请检查微信 Openid 是否有误/Appid与QQopenid无法关联/微信Openid权限是否有开通/手机号是否为中国大陆手机号
	// 4 黑名单 该账号在业务侧有过拉黑记录
	// 5 白名单 业务自行有添加过白名单记录
	// 行为风险
	// 101 批量操作存在 IP/设备/环境等因素的聚集性异常
	// 1011 疑似 IP 属性聚集 出现 IP 聚集
	// 1012 疑似设备属性聚集 出现设备聚集
	// 102 自动机 疑似自动机批量请求
	// 103 恶意行为-网赚 疑似网赚
	// 104 微信登录态无效 检查 WeChatAccessToken 参数，是否已经失效
	// 201 环境风险 环境异常 操作 IP/设备/环境存在异常。当前 IP 为非常用 IP 或恶意 IP 段
	// 2011 疑似非常用 IP 请求 当前请求 IP 非该账号常用 IP
	// 2012 疑似 IP 异常 使用 idc 机房 IP 或使用代理 IP 或使用恶意 IP 等
	// 205 非公网有效IP 传进来的 IP 地址为内网 IP 地址或者 IP 保留地址
	// 设备风险
	// 206 设备异常 该设备存在异常的使用行为
	// 2061 疑似非常用设备 当前请求的设备非该账号常用设备
	// 2062 疑似虚拟设备 请求设备为模拟器、脚本、云设备等虚拟设备
	// 2063 疑似群控设备 请求设备为猫池、手机墙等群控设备
	// 10201 设备处于开发者模式 来源于Android
	// 10202 设备疑似 Root 来源于Android
	// 10203 疑似应用被注 来源于Android
	// 10204 疑似应用被重打包 来源于Android
	// 10205 疑似使用 hook 技术 来源于Android
	// 10206 疑似应用被双开 来源于Android
	// 10207 疑似设备存在风险进程 来源于Android
	// 10208 疑似伪造地理位置 来源于Android
	// 10209 疑似使用 VPN 软件 来源于Android
	// 10210 疑似使用代理软件 来源于Android
	// 10211 疑似设备处于调试模式 来源于Android
	// 10212 疑似高危 ROM 来源于Android
	// 10213 疑似检测到系统劫持 来源于Android
	// 10003 疑似模拟器 来源于Android
	// 10301 疑似主流模拟器 来源于Android
	// 10302 疑似云模拟器 来源于Android
	// 10303 疑似开发板设备 来源于Android
	// 10004 疑似群控设备风险 来源于Android
	// 10401 疑似使用自动化软件 来源于Android
	// 10402 疑似群控自动化操作 来源于Android
	// 10501 疑似参数异常-IMEI 来源于Android
	// 10502 疑似参数异常-FP 来源于Android
	// 10504 疑似参数异常-IMSI 来源于Android
	// 10801 疑似存在刷量安装应用的行为 来源于Android
	// 10901 疑似多账号异常 来源于Android
	// 11001 疑似设备参数篡改 来源于Android
	// 11002 疑似存在风险软件 来源于Android
	// 11003 疑似应用被调试 来源于Android
	// 11100 疑似云真机 来源于Android
	// 25001 设备疑似越狱 来源于IOS
	// 25004 进程疑似有注入文件 来源于IOS
	// 25005 疑似使用代理软件 来源于IOS
	// 25006 疑似使用 VPN 软件 来源于IOS
	// 25007 疑似被 Hook 来源于IOS
	// 25008 疑似进程被调试 来源于IOS
	// 25009 疑似多开 来源于IOS
	// 25010 疑似改机 来源于IOS
	// 25011 疑似应用二次打包 来源于IOS
	// 25012 疑似模拟器 来源于IOS
	// 25013 疑似云真机 来源于IOS
	// 25014 疑似云模拟器 来源于IOS
	// 25015 疑似伪造地理位置 来源于IOS
	// 25016 疑似使用自动化脚本 来源于IOS
	// 25017 疑似群控自动化操作 来源于IOS
	// 30001 疑似虚拟浏览器 来源于H5
	// 30002 疑似安装作弊插件 来源于H5
	// 30003 疑似浏览器参数遭篡改 来源于H5
	// 30004 疑似 JS 代码被篡改 来源于H5
	// 30005 疑似 JS 被调试 来源于H5
	// 30006 Cookies 被禁用 来源于H5
	// 40001 疑似伪造地理位置 来源于小程序
	// 40002 疑似被调试 来源于小程序
	// 40003 疑似模拟器 来源于小程序
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskLabels []*int64 `json:"RiskLabels,omitempty" name:"RiskLabels"`
}

type MarketingInfo struct {
	// 投放模式（0=PDB，1=PD，2=RTB，10=其他）
	DeliveryMode *int64 `json:"DeliveryMode,omitempty" name:"DeliveryMode"`

	// 广告位类型 （0=前贴片，1=开屏广告，2=网页头部广告、3=网页中部广告、4=网页底部广告、5=悬浮广告、10=其它）
	AdvertisingType *int64 `json:"AdvertisingType,omitempty" name:"AdvertisingType"`

	// 是否全屏插广告（0-否，1-是）
	FullScreen *int64 `json:"FullScreen,omitempty" name:"FullScreen"`

	// 广告位宽度
	AdvertisingSpaceWidth *int64 `json:"AdvertisingSpaceWidth,omitempty" name:"AdvertisingSpaceWidth"`

	// 广告位高度
	AdvertisingSpaceHeight *int64 `json:"AdvertisingSpaceHeight,omitempty" name:"AdvertisingSpaceHeight"`

	// 网址
	Url *string `json:"Url,omitempty" name:"Url"`
}

type UniversalAccountInfo struct {
	// 账号值：
	// 当账户类型为1时，填入手机号，如135****3695；
	// 当账户类型为2、3或100时，填入对应的值。
	AccountId *string `json:"AccountId,omitempty" name:"AccountId"`
}

type UserInfo struct {
	// 用户外网IP地址
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 来源渠道编码
	ChannelSource *string `json:"ChannelSource,omitempty" name:"ChannelSource"`

	// 用户登录平台。1：Android 2：iOS 3：H5 4：小程序
	Platform *int64 `json:"Platform,omitempty" name:"Platform"`

	// 姓名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 年龄
	Age *int64 `json:"Age,omitempty" name:"Age"`

	// 性别：
	// male（男）
	// female（女）
	Gender *string `json:"Gender,omitempty" name:"Gender"`

	// 身份证号
	ResidentIdentityCard *string `json:"ResidentIdentityCard,omitempty" name:"ResidentIdentityCard"`

	// 邮箱地址
	Email *string `json:"Email,omitempty" name:"Email"`

	// 用户地址
	Address *string `json:"Address,omitempty" name:"Address"`

	// 用户昵称
	Nickname *string `json:"Nickname,omitempty" name:"Nickname"`
}