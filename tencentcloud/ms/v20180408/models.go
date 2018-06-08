// Copyright 1999-2018 Tencent Ltd.
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

package v20180408

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AdInfo struct {
	// 插播广告列表
	Spots []*PluginInfo `json:"Spots" name:"Spots" list`
	// 精品推荐广告列表
	BoutiqueRecommands []*PluginInfo `json:"BoutiqueRecommands" name:"BoutiqueRecommands" list`
	// 悬浮窗广告列表
	FloatWindowses []*PluginInfo `json:"FloatWindowses" name:"FloatWindowses" list`
	// banner广告列表
	Banners []*PluginInfo `json:"Banners" name:"Banners" list`
	// 积分墙广告列表
	IntegralWalls []*PluginInfo `json:"IntegralWalls" name:"IntegralWalls" list`
	// 通知栏广告列表
	NotifyBars []*PluginInfo `json:"NotifyBars" name:"NotifyBars" list`
}

type AppDetailInfo struct {
	// app的名称
	AppName *string `json:"AppName" name:"AppName"`
	// app的包名
	AppPkgName *string `json:"AppPkgName" name:"AppPkgName"`
	// app的版本号
	AppVersion *string `json:"AppVersion" name:"AppVersion"`
	// app的大小
	AppSize *uint64 `json:"AppSize" name:"AppSize"`
	// app的md5
	AppMd5 *string `json:"AppMd5" name:"AppMd5"`
	// app的图标url
	AppIconUrl *string `json:"AppIconUrl" name:"AppIconUrl"`
	// app的文件名称
	FileName *string `json:"FileName" name:"FileName"`
}

type AppInfo struct {
	// app的url，必须保证不用权限校验就可以下载
	AppUrl *string `json:"AppUrl" name:"AppUrl"`
	// app的md5
	AppMd5 *string `json:"AppMd5" name:"AppMd5"`
	// app的大小
	AppSize *uint64 `json:"AppSize" name:"AppSize"`
	// app的文件名，指定后加固后的文件名是{FileName}_legu.apk
	FileName *string `json:"FileName" name:"FileName"`
	// app的包名
	AppPkgName *string `json:"AppPkgName" name:"AppPkgName"`
	// app的版本号
	AppVersion *string `json:"AppVersion" name:"AppVersion"`
	// app的图标url
	AppIconUrl *string `json:"AppIconUrl" name:"AppIconUrl"`
	// app的名称
	AppName *string `json:"AppName" name:"AppName"`
}

type AppScanSet struct {
	// 任务唯一标识
	ItemId *string `json:"ItemId" name:"ItemId"`
	// app的名称
	AppName *string `json:"AppName" name:"AppName"`
	// app的包名
	AppPkgName *string `json:"AppPkgName" name:"AppPkgName"`
	// app的版本号
	AppVersion *string `json:"AppVersion" name:"AppVersion"`
	// app的md5
	AppMd5 *string `json:"AppMd5" name:"AppMd5"`
	// app的大小
	AppSize *uint64 `json:"AppSize" name:"AppSize"`
	// 扫描结果返回码
	ScanCode *uint64 `json:"ScanCode" name:"ScanCode"`
	// 任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
	TaskStatus *uint64 `json:"TaskStatus" name:"TaskStatus"`
	// 提交扫描时间
	TaskTime *uint64 `json:"TaskTime" name:"TaskTime"`
	// app的图标url
	AppIconUrl *string `json:"AppIconUrl" name:"AppIconUrl"`
	// 标识唯一该app，主要用于删除
	AppSid *string `json:"AppSid" name:"AppSid"`
	// 安全类型:1-安全软件，2-风险软件，3病毒软件
	SafeType *uint64 `json:"SafeType" name:"SafeType"`
	// 漏洞个数
	VulCount *uint64 `json:"VulCount" name:"VulCount"`
}

type AppSetInfo struct {
	// 任务唯一标识
	ItemId *string `json:"ItemId" name:"ItemId"`
	// app的名称
	AppName *string `json:"AppName" name:"AppName"`
	// app的包名
	AppPkgName *string `json:"AppPkgName" name:"AppPkgName"`
	// app的版本号
	AppVersion *string `json:"AppVersion" name:"AppVersion"`
	// app的md5
	AppMd5 *string `json:"AppMd5" name:"AppMd5"`
	// app的大小
	AppSize *uint64 `json:"AppSize" name:"AppSize"`
	// 加固服务版本
	ServiceEdition *string `json:"ServiceEdition" name:"ServiceEdition"`
	// 加固结果返回码
	ShieldCode *uint64 `json:"ShieldCode" name:"ShieldCode"`
	// 加固后的APP下载地址
	AppUrl *string `json:"AppUrl" name:"AppUrl"`
	// 任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
	TaskStatus *uint64 `json:"TaskStatus" name:"TaskStatus"`
	// 请求的客户端ip
	ClientIp *string `json:"ClientIp" name:"ClientIp"`
	// 提交加固时间
	TaskTime *uint64 `json:"TaskTime" name:"TaskTime"`
	// app的图标url
	AppIconUrl *string `json:"AppIconUrl" name:"AppIconUrl"`
	// 加固后app的md5
	ShieldMd5 *string `json:"ShieldMd5" name:"ShieldMd5"`
	// 加固后app的大小
	ShieldSize *uint64 `json:"ShieldSize" name:"ShieldSize"`
}

type CreateScanInstancesRequest struct {
	*tchttp.BaseRequest
	// 待扫描的app信息列表，一次最多提交20个
	AppInfos []*AppInfo `json:"AppInfos" name:"AppInfos" list`
	// 扫描信息
	ScanInfo *ScanInfo `json:"ScanInfo" name:"ScanInfo"`
}

func (r *CreateScanInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateScanInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateScanInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 任务唯一标识
		ItemId *string `json:"ItemId" name:"ItemId"`
		// 任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
		Progress *uint64 `json:"Progress" name:"Progress"`
		// 提交成功的app的md5集合
		AppMd5s []*string `json:"AppMd5s" name:"AppMd5s" list`
		// 剩余可用次数
		LimitCount *uint64 `json:"LimitCount" name:"LimitCount"`
		// 到期时间
		LimitTime *uint64 `json:"LimitTime" name:"LimitTime"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateScanInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateScanInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateShieldInstanceRequest struct {
	*tchttp.BaseRequest
	// 待加固的应用信息
	AppInfo *AppInfo `json:"AppInfo" name:"AppInfo"`
	// 加固服务信息
	ServiceInfo *ServiceInfo `json:"ServiceInfo" name:"ServiceInfo"`
}

func (r *CreateShieldInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateShieldInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateShieldInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
		Progress *uint64 `json:"Progress" name:"Progress"`
		// 任务唯一标识
		ItemId *string `json:"ItemId" name:"ItemId"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateShieldInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateShieldInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteScanInstancesRequest struct {
	*tchttp.BaseRequest
	// 删除一个或多个扫描的app，最大支持20个
	AppSids []*string `json:"AppSids" name:"AppSids" list`
}

func (r *DeleteScanInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteScanInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteScanInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
		Progress *uint64 `json:"Progress" name:"Progress"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteScanInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteScanInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteShieldInstancesRequest struct {
	*tchttp.BaseRequest
	// 任务唯一标识ItemId的列表
	ItemIds []*string `json:"ItemIds" name:"ItemIds" list`
}

func (r *DeleteShieldInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteShieldInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteShieldInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
		Progress *uint64 `json:"Progress" name:"Progress"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteShieldInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteShieldInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeScanInstancesRequest struct {
	*tchttp.BaseRequest
	// 支持通过app名称，app包名进行筛选
	Filters []*Filters `json:"Filters" name:"Filters" list`
	// 偏移量，默认为0
	Offset *uint64 `json:"Offset" name:"Offset"`
	// 数量限制，默认为20，最大值为100。
	Limit *uint64 `json:"Limit" name:"Limit"`
	// 可以提供ItemId数组来查询一个或者多个结果。注意不可以同时指定ItemIds和Filters。
	ItemIds []*string `json:"ItemIds" name:"ItemIds" list`
	// 按某个字段排序，目前仅支持TaskTime排序。
	OrderField *string `json:"OrderField" name:"OrderField"`
	// 升序（asc）还是降序（desc），默认：desc。
	OrderDirection *string `json:"OrderDirection" name:"OrderDirection"`
}

func (r *DescribeScanInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeScanInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeScanInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 符合要求的app数量
		TotalCount *uint64 `json:"TotalCount" name:"TotalCount"`
		// 一个关于app详细信息的结构体，主要包括app的基本信息和扫描状态信息。
		ScanSet []*AppScanSet `json:"ScanSet" name:"ScanSet" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScanInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeScanInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeScanResultsRequest struct {
	*tchttp.BaseRequest
	// 任务唯一标识
	ItemId *string `json:"ItemId" name:"ItemId"`
	// 批量查询一个或者多个app的扫描结果，如果不传表示查询该任务下所提交的所有app
	AppMd5s []*string `json:"AppMd5s" name:"AppMd5s" list`
}

func (r *DescribeScanResultsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeScanResultsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeScanResultsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 批量扫描的app结果集
		ScanSet []*ScanSetInfo `json:"ScanSet" name:"ScanSet" list`
		// 批量扫描结果的个数
		TotalCount *uint64 `json:"TotalCount" name:"TotalCount"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScanResultsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeScanResultsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeShieldInstancesRequest struct {
	*tchttp.BaseRequest
	// 支持通过app名称，app包名，加固的服务版本，提交的渠道进行筛选。
	Filters []*Filters `json:"Filters" name:"Filters" list`
	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset" name:"Offset"`
	// 数量限制，默认为20，最大值为100。
	Limit *uint64 `json:"Limit" name:"Limit"`
	// 可以提供ItemId数组来查询一个或者多个结果。注意不可以同时指定ItemIds和Filters。
	ItemIds []*string `json:"ItemIds" name:"ItemIds" list`
	// 按某个字段排序，目前仅支持TaskTime排序。
	OrderField *string `json:"OrderField" name:"OrderField"`
	// 升序（asc）还是降序（desc），默认：desc。
	OrderDirection *string `json:"OrderDirection" name:"OrderDirection"`
}

func (r *DescribeShieldInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeShieldInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeShieldInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 符合要求的app数量
		TotalCount *uint64 `json:"TotalCount" name:"TotalCount"`
		// 一个关于app详细信息的结构体，主要包括app的基本信息和加固信息。
		AppSet []*AppSetInfo `json:"AppSet" name:"AppSet" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeShieldInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeShieldInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeShieldResultRequest struct {
	*tchttp.BaseRequest
	// 任务唯一标识
	ItemId *string `json:"ItemId" name:"ItemId"`
}

func (r *DescribeShieldResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeShieldResultRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeShieldResultResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
		TaskStatus *uint64 `json:"TaskStatus" name:"TaskStatus"`
		// app加固前的详细信息
		AppDetailInfo *AppDetailInfo `json:"AppDetailInfo" name:"AppDetailInfo"`
		// app加固后的详细信息
		ShieldInfo *ShieldInfo `json:"ShieldInfo" name:"ShieldInfo"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeShieldResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeShieldResultResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Filters struct {
	// 需要过滤的字段
	Name *string `json:"Name" name:"Name"`
	// 需要过滤字段的值
	Value *string `json:"Value" name:"Value"`
}

type PluginInfo struct {
	// 插件类型，分别为 1-通知栏广告，2-积分墙广告，3-banner广告，4- 悬浮窗图标广告，5-精品推荐列表广告, 6-插播广告
	PluginType *uint64 `json:"PluginType" name:"PluginType"`
	// 插件名称
	PluginName *string `json:"PluginName" name:"PluginName"`
	// 插件描述
	PluginDesc *string `json:"PluginDesc" name:"PluginDesc"`
}

type ScanInfo struct {
	// 任务处理完成后的反向通知回调地址,批量提交app每扫描完成一个会通知一次,通知为POST请求，post信息{ItemId:
	CallbackUrl *string `json:"CallbackUrl" name:"CallbackUrl"`
	// VULSCAN-漏洞扫描信息，VIRUSSCAN-返回病毒扫描信息， ADSCAN-广告扫描信息，PLUGINSCAN-插件扫描信息，可以自由组合
	ScanTypes []*string `json:"ScanTypes" name:"ScanTypes" list`
}

type ScanSetInfo struct {
	// 任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
	TaskStatus *uint64 `json:"TaskStatus" name:"TaskStatus"`
	// app信息
	AppDetailInfo *AppDetailInfo `json:"AppDetailInfo" name:"AppDetailInfo"`
	// 病毒信息
	VirusInfo *VirusInfo `json:"VirusInfo" name:"VirusInfo"`
	// 漏洞信息
	VulInfo *VulInfo `json:"VulInfo" name:"VulInfo"`
	// 广告插件信息
	AdInfo *AdInfo `json:"AdInfo" name:"AdInfo"`
	// 提交扫描的时间
	TaskTime *uint64 `json:"TaskTime" name:"TaskTime"`
}

type ServiceInfo struct {
	// 服务版本，基础版basic,专业版Professional
	ServiceEdition *string `json:"ServiceEdition" name:"ServiceEdition"`
	// 任务处理完成后的反向通知回调地址,通知为POST请求，post信息{ItemId:"xxxduuyt-ugusg"}
	CallbackUrl *string `json:"CallbackUrl" name:"CallbackUrl"`
	// 提交来源 YYB-应用宝 RDM-rdm MC-控制台 MAC_TOOL-mac工具 WIN_TOOL-window工具
	SubmitSource *string `json:"SubmitSource" name:"SubmitSource"`
}

type ShieldInfo struct {
	// 加固结果的返回码
	ShieldCode *uint64 `json:"ShieldCode" name:"ShieldCode"`
	// 加固后app的大小
	ShieldSize *uint64 `json:"ShieldSize" name:"ShieldSize"`
	// 加固后app的md5
	ShieldMd5 *string `json:"ShieldMd5" name:"ShieldMd5"`
	// 加固后的APP下载地址
	AppUrl *string `json:"AppUrl" name:"AppUrl"`
	// 加固的提交时间
	TaskTime *uint64 `json:"TaskTime" name:"TaskTime"`
	// 任务唯一标识
	ItemId *string `json:"ItemId" name:"ItemId"`
}

type VirusInfo struct {
	// 软件安全类型，分别为0-未知、 1-安全软件、2-风险软件、3-病毒软件
	SafeType *int64 `json:"SafeType" name:"SafeType"`
	// 病毒名称， utf8编码，非病毒时值为空
	VirusName *string `json:"VirusName" name:"VirusName"`
	// 病毒描述，utf8编码，非病毒时值为空
	VirusDesc *string `json:"VirusDesc" name:"VirusDesc"`
}

type VulInfo struct {
	// 漏洞列表
	VulList []*VulList `json:"VulList" name:"VulList" list`
	// 漏洞文件评分
	VulFileScore *uint64 `json:"VulFileScore" name:"VulFileScore"`
}

type VulList struct {
	// 漏洞id
	VulId *string `json:"VulId" name:"VulId"`
	// 漏洞名称
	VulName *string `json:"VulName" name:"VulName"`
	// 漏洞代码
	VulCode *string `json:"VulCode" name:"VulCode"`
	// 漏洞描述
	VulDesc *string `json:"VulDesc" name:"VulDesc"`
	// 漏洞解决方案
	VulSolution *string `json:"VulSolution" name:"VulSolution"`
	// 漏洞来源类别，0默认自身，1第三方插件
	VulSrcType *int64 `json:"VulSrcType" name:"VulSrcType"`
	// 漏洞位置
	VulFilepath *string `json:"VulFilepath" name:"VulFilepath"`
	// 风险级别：1 低风险 ；2中等风险；3 高风险
	RiskLevel *uint64 `json:"RiskLevel" name:"RiskLevel"`
}
