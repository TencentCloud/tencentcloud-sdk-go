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

package v20180408

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AdInfo struct {
	// 插播广告列表
	Spots []*PluginInfo `json:"Spots,omitempty" name:"Spots"`

	// 精品推荐广告列表
	BoutiqueRecommands []*PluginInfo `json:"BoutiqueRecommands,omitempty" name:"BoutiqueRecommands"`

	// 悬浮窗广告列表
	FloatWindowses []*PluginInfo `json:"FloatWindowses,omitempty" name:"FloatWindowses"`

	// banner广告列表
	Banners []*PluginInfo `json:"Banners,omitempty" name:"Banners"`

	// 积分墙广告列表
	IntegralWalls []*PluginInfo `json:"IntegralWalls,omitempty" name:"IntegralWalls"`

	// 通知栏广告列表
	NotifyBars []*PluginInfo `json:"NotifyBars,omitempty" name:"NotifyBars"`
}

type AppDetailInfo struct {
	// app的名称
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// app的包名
	AppPkgName *string `json:"AppPkgName,omitempty" name:"AppPkgName"`

	// app的版本号
	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`

	// app的大小
	AppSize *uint64 `json:"AppSize,omitempty" name:"AppSize"`

	// app的md5
	AppMd5 *string `json:"AppMd5,omitempty" name:"AppMd5"`

	// app的图标url
	AppIconUrl *string `json:"AppIconUrl,omitempty" name:"AppIconUrl"`

	// app的文件名称
	FileName *string `json:"FileName,omitempty" name:"FileName"`
}

type AppInfo struct {
	// app的url，必须保证不用权限校验就可以下载
	AppUrl *string `json:"AppUrl,omitempty" name:"AppUrl"`

	// app的md5，需要正确传递
	AppMd5 *string `json:"AppMd5,omitempty" name:"AppMd5"`

	// app的大小
	AppSize *uint64 `json:"AppSize,omitempty" name:"AppSize"`

	// app的文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// app的包名，需要正确的传递此字段
	AppPkgName *string `json:"AppPkgName,omitempty" name:"AppPkgName"`

	// app的版本号
	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`

	// app的图标url
	AppIconUrl *string `json:"AppIconUrl,omitempty" name:"AppIconUrl"`

	// app的名称
	AppName *string `json:"AppName,omitempty" name:"AppName"`
}

type AppScanSet struct {
	// 任务唯一标识
	ItemId *string `json:"ItemId,omitempty" name:"ItemId"`

	// app的名称
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// app的包名
	AppPkgName *string `json:"AppPkgName,omitempty" name:"AppPkgName"`

	// app的版本号
	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`

	// app的md5
	AppMd5 *string `json:"AppMd5,omitempty" name:"AppMd5"`

	// app的大小
	AppSize *uint64 `json:"AppSize,omitempty" name:"AppSize"`

	// 扫描结果返回码
	ScanCode *uint64 `json:"ScanCode,omitempty" name:"ScanCode"`

	// 任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
	TaskStatus *uint64 `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// 提交扫描时间
	TaskTime *uint64 `json:"TaskTime,omitempty" name:"TaskTime"`

	// app的图标url
	AppIconUrl *string `json:"AppIconUrl,omitempty" name:"AppIconUrl"`

	// 标识唯一该app，主要用于删除
	AppSid *string `json:"AppSid,omitempty" name:"AppSid"`

	// 安全类型:1-安全软件，2-风险软件，3病毒软件
	SafeType *uint64 `json:"SafeType,omitempty" name:"SafeType"`

	// 漏洞个数
	VulCount *uint64 `json:"VulCount,omitempty" name:"VulCount"`
}

type AppSetInfo struct {
	// 任务唯一标识
	ItemId *string `json:"ItemId,omitempty" name:"ItemId"`

	// app的名称
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// app的包名
	AppPkgName *string `json:"AppPkgName,omitempty" name:"AppPkgName"`

	// app的版本号
	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`

	// app的md5
	AppMd5 *string `json:"AppMd5,omitempty" name:"AppMd5"`

	// app的大小
	AppSize *uint64 `json:"AppSize,omitempty" name:"AppSize"`

	// 加固服务版本
	ServiceEdition *string `json:"ServiceEdition,omitempty" name:"ServiceEdition"`

	// 加固结果返回码
	ShieldCode *uint64 `json:"ShieldCode,omitempty" name:"ShieldCode"`

	// 加固后的APP下载地址
	AppUrl *string `json:"AppUrl,omitempty" name:"AppUrl"`

	// 任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
	TaskStatus *uint64 `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// 请求的客户端ip
	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`

	// 提交加固时间
	TaskTime *uint64 `json:"TaskTime,omitempty" name:"TaskTime"`

	// app的图标url
	AppIconUrl *string `json:"AppIconUrl,omitempty" name:"AppIconUrl"`

	// 加固后app的md5
	ShieldMd5 *string `json:"ShieldMd5,omitempty" name:"ShieldMd5"`

	// 加固后app的大小
	ShieldSize *uint64 `json:"ShieldSize,omitempty" name:"ShieldSize"`
}

type BindInfo struct {
	// app的icon的url
	AppIconUrl *string `json:"AppIconUrl,omitempty" name:"AppIconUrl"`

	// app的名称
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// app的包名
	AppPkgName *string `json:"AppPkgName,omitempty" name:"AppPkgName"`
}

// Predefined struct for user
type CreateBindInstanceRequestParams struct {
	// 资源id，全局唯一
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// app的icon的url
	AppIconUrl *string `json:"AppIconUrl,omitempty" name:"AppIconUrl"`

	// app的名称
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// app的包名
	AppPkgName *string `json:"AppPkgName,omitempty" name:"AppPkgName"`
}

type CreateBindInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 资源id，全局唯一
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// app的icon的url
	AppIconUrl *string `json:"AppIconUrl,omitempty" name:"AppIconUrl"`

	// app的名称
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// app的包名
	AppPkgName *string `json:"AppPkgName,omitempty" name:"AppPkgName"`
}

func (r *CreateBindInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBindInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceId")
	delete(f, "AppIconUrl")
	delete(f, "AppName")
	delete(f, "AppPkgName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBindInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBindInstanceResponseParams struct {
	// 任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
	Progress *uint64 `json:"Progress,omitempty" name:"Progress"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateBindInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateBindInstanceResponseParams `json:"Response"`
}

func (r *CreateBindInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBindInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCosSecKeyInstanceRequestParams struct {
	// 地域信息，例如广州：ap-guangzhou，上海：ap-shanghai，默认为广州。
	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`

	// 密钥有效时间，默认为1小时。
	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`
}

type CreateCosSecKeyInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 地域信息，例如广州：ap-guangzhou，上海：ap-shanghai，默认为广州。
	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`

	// 密钥有效时间，默认为1小时。
	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`
}

func (r *CreateCosSecKeyInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCosSecKeyInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CosRegion")
	delete(f, "Duration")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCosSecKeyInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCosSecKeyInstanceResponseParams struct {
	// COS密钥对应的AppId
	CosAppid *uint64 `json:"CosAppid,omitempty" name:"CosAppid"`

	// COS密钥对应的存储桶名
	CosBucket *string `json:"CosBucket,omitempty" name:"CosBucket"`

	// 存储桶对应的地域
	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`

	// 密钥过期时间
	ExpireTime *uint64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 密钥ID信息
	CosId *string `json:"CosId,omitempty" name:"CosId"`

	// 密钥KEY信息
	CosKey *string `json:"CosKey,omitempty" name:"CosKey"`

	// 密钥TOCKEN信息
	CosTocken *string `json:"CosTocken,omitempty" name:"CosTocken"`

	// 密钥可访问的文件前缀人。例如：CosPrefix=test/123/666，则该密钥只能操作test/123/666为前缀的文件，例如test/123/666/1.txt
	CosPrefix *string `json:"CosPrefix,omitempty" name:"CosPrefix"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCosSecKeyInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateCosSecKeyInstanceResponseParams `json:"Response"`
}

func (r *CreateCosSecKeyInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCosSecKeyInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateResourceInstancesRequestParams struct {
	// 资源类型id。13624：加固专业版。
	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`

	// 时间单位，取值为d，m，y，分别表示天，月，年。
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 时间数量。
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 资源数量。
	ResourceNum *uint64 `json:"ResourceNum,omitempty" name:"ResourceNum"`
}

type CreateResourceInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 资源类型id。13624：加固专业版。
	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`

	// 时间单位，取值为d，m，y，分别表示天，月，年。
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 时间数量。
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 资源数量。
	ResourceNum *uint64 `json:"ResourceNum,omitempty" name:"ResourceNum"`
}

func (r *CreateResourceInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateResourceInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Pid")
	delete(f, "TimeUnit")
	delete(f, "TimeSpan")
	delete(f, "ResourceNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateResourceInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateResourceInstancesResponseParams struct {
	// 新创建的资源列表。
	ResourceSet []*string `json:"ResourceSet,omitempty" name:"ResourceSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateResourceInstancesResponse struct {
	*tchttp.BaseResponse
	Response *CreateResourceInstancesResponseParams `json:"Response"`
}

func (r *CreateResourceInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateResourceInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateScanInstancesRequestParams struct {
	// 待扫描的app信息列表，一次最多提交20个
	AppInfos []*AppInfo `json:"AppInfos,omitempty" name:"AppInfos"`

	// 扫描信息
	ScanInfo *ScanInfo `json:"ScanInfo,omitempty" name:"ScanInfo"`
}

type CreateScanInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 待扫描的app信息列表，一次最多提交20个
	AppInfos []*AppInfo `json:"AppInfos,omitempty" name:"AppInfos"`

	// 扫描信息
	ScanInfo *ScanInfo `json:"ScanInfo,omitempty" name:"ScanInfo"`
}

func (r *CreateScanInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateScanInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppInfos")
	delete(f, "ScanInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateScanInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateScanInstancesResponseParams struct {
	// 任务唯一标识
	ItemId *string `json:"ItemId,omitempty" name:"ItemId"`

	// 任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
	Progress *uint64 `json:"Progress,omitempty" name:"Progress"`

	// 提交成功的app的md5集合
	AppMd5s []*string `json:"AppMd5s,omitempty" name:"AppMd5s"`

	// 剩余可用次数
	LimitCount *uint64 `json:"LimitCount,omitempty" name:"LimitCount"`

	// 到期时间
	LimitTime *uint64 `json:"LimitTime,omitempty" name:"LimitTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateScanInstancesResponse struct {
	*tchttp.BaseResponse
	Response *CreateScanInstancesResponseParams `json:"Response"`
}

func (r *CreateScanInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateScanInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateShieldInstanceRequestParams struct {
	// 待加固的应用信息
	AppInfo *AppInfo `json:"AppInfo,omitempty" name:"AppInfo"`

	// 加固服务信息
	ServiceInfo *ServiceInfo `json:"ServiceInfo,omitempty" name:"ServiceInfo"`
}

type CreateShieldInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 待加固的应用信息
	AppInfo *AppInfo `json:"AppInfo,omitempty" name:"AppInfo"`

	// 加固服务信息
	ServiceInfo *ServiceInfo `json:"ServiceInfo,omitempty" name:"ServiceInfo"`
}

func (r *CreateShieldInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateShieldInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppInfo")
	delete(f, "ServiceInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateShieldInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateShieldInstanceResponseParams struct {
	// 任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
	Progress *uint64 `json:"Progress,omitempty" name:"Progress"`

	// 任务唯一标识
	ItemId *string `json:"ItemId,omitempty" name:"ItemId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateShieldInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateShieldInstanceResponseParams `json:"Response"`
}

func (r *CreateShieldInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateShieldInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateShieldPlanInstanceRequestParams struct {
	// 资源id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 策略名称
	PlanName *string `json:"PlanName,omitempty" name:"PlanName"`

	// 策略具体信息
	PlanInfo *PlanInfo `json:"PlanInfo,omitempty" name:"PlanInfo"`
}

type CreateShieldPlanInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 资源id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 策略名称
	PlanName *string `json:"PlanName,omitempty" name:"PlanName"`

	// 策略具体信息
	PlanInfo *PlanInfo `json:"PlanInfo,omitempty" name:"PlanInfo"`
}

func (r *CreateShieldPlanInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateShieldPlanInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceId")
	delete(f, "PlanName")
	delete(f, "PlanInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateShieldPlanInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateShieldPlanInstanceResponseParams struct {
	// 策略id
	PlanId *uint64 `json:"PlanId,omitempty" name:"PlanId"`

	// 任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
	Progress *uint64 `json:"Progress,omitempty" name:"Progress"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateShieldPlanInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateShieldPlanInstanceResponseParams `json:"Response"`
}

func (r *CreateShieldPlanInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateShieldPlanInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteScanInstancesRequestParams struct {
	// 删除一个或多个扫描的app，最大支持20个
	AppSids []*string `json:"AppSids,omitempty" name:"AppSids"`
}

type DeleteScanInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 删除一个或多个扫描的app，最大支持20个
	AppSids []*string `json:"AppSids,omitempty" name:"AppSids"`
}

func (r *DeleteScanInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteScanInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppSids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteScanInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteScanInstancesResponseParams struct {
	// 任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
	Progress *uint64 `json:"Progress,omitempty" name:"Progress"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteScanInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteScanInstancesResponseParams `json:"Response"`
}

func (r *DeleteScanInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteScanInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteShieldInstancesRequestParams struct {
	// 任务唯一标识ItemId的列表
	ItemIds []*string `json:"ItemIds,omitempty" name:"ItemIds"`
}

type DeleteShieldInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 任务唯一标识ItemId的列表
	ItemIds []*string `json:"ItemIds,omitempty" name:"ItemIds"`
}

func (r *DeleteShieldInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteShieldInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ItemIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteShieldInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteShieldInstancesResponseParams struct {
	// 任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
	Progress *uint64 `json:"Progress,omitempty" name:"Progress"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteShieldInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteShieldInstancesResponseParams `json:"Response"`
}

func (r *DeleteShieldInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteShieldInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApkDetectionResultRequestParams struct {
	// 软件包的下载链接
	ApkUrl *string `json:"ApkUrl,omitempty" name:"ApkUrl"`

	// 软件包的md5值，具有唯一性。腾讯APK云检测服务会根据md5值来判断该包是否为库中已收集的样本，已存在，则返回检测结果，反之，需要一定时间检测该样本。
	ApkMd5 *string `json:"ApkMd5,omitempty" name:"ApkMd5"`
}

type DescribeApkDetectionResultRequest struct {
	*tchttp.BaseRequest
	
	// 软件包的下载链接
	ApkUrl *string `json:"ApkUrl,omitempty" name:"ApkUrl"`

	// 软件包的md5值，具有唯一性。腾讯APK云检测服务会根据md5值来判断该包是否为库中已收集的样本，已存在，则返回检测结果，反之，需要一定时间检测该样本。
	ApkMd5 *string `json:"ApkMd5,omitempty" name:"ApkMd5"`
}

func (r *DescribeApkDetectionResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApkDetectionResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApkUrl")
	delete(f, "ApkMd5")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApkDetectionResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApkDetectionResultResponseParams struct {
	// 响应结果，ok表示正常，error表示错误
	Result *string `json:"Result,omitempty" name:"Result"`

	// Result为error错误时的原因说明
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// APK检测结果数组
	ResultList []*ResultListItem `json:"ResultList,omitempty" name:"ResultList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApkDetectionResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApkDetectionResultResponseParams `json:"Response"`
}

func (r *DescribeApkDetectionResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApkDetectionResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceInstancesRequestParams struct {
	// 支持CreateTime、ExpireTime、AppName、AppPkgName、BindValue、IsBind过滤
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 数量限制，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 资源类别id数组，13624：加固专业版，12750：企业版。空数组表示返回全部资源。
	Pids []*uint64 `json:"Pids,omitempty" name:"Pids"`

	// 按某个字段排序，目前支持CreateTime、ExpireTime其中的一个排序。
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 升序（asc）还是降序（desc），默认：desc。
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

type DescribeResourceInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 支持CreateTime、ExpireTime、AppName、AppPkgName、BindValue、IsBind过滤
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 数量限制，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 资源类别id数组，13624：加固专业版，12750：企业版。空数组表示返回全部资源。
	Pids []*uint64 `json:"Pids,omitempty" name:"Pids"`

	// 按某个字段排序，目前支持CreateTime、ExpireTime其中的一个排序。
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 升序（asc）还是降序（desc），默认：desc。
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

func (r *DescribeResourceInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Pids")
	delete(f, "OrderField")
	delete(f, "OrderDirection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourceInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceInstancesResponseParams struct {
	// 符合要求的资源数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 符合要求的资源数组
	ResourceSet []*ResourceInfo `json:"ResourceSet,omitempty" name:"ResourceSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeResourceInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourceInstancesResponseParams `json:"Response"`
}

func (r *DescribeResourceInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanInstancesRequestParams struct {
	// 支持通过app名称，app包名进行筛选
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 数量限制，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 可以提供ItemId数组来查询一个或者多个结果。注意不可以同时指定ItemIds和Filters。
	ItemIds []*string `json:"ItemIds,omitempty" name:"ItemIds"`

	// 按某个字段排序，目前仅支持TaskTime排序。
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 升序（asc）还是降序（desc），默认：desc。
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

type DescribeScanInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 支持通过app名称，app包名进行筛选
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 数量限制，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 可以提供ItemId数组来查询一个或者多个结果。注意不可以同时指定ItemIds和Filters。
	ItemIds []*string `json:"ItemIds,omitempty" name:"ItemIds"`

	// 按某个字段排序，目前仅支持TaskTime排序。
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 升序（asc）还是降序（desc），默认：desc。
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

func (r *DescribeScanInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ItemIds")
	delete(f, "OrderField")
	delete(f, "OrderDirection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScanInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanInstancesResponseParams struct {
	// 符合要求的app数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 一个关于app详细信息的结构体，主要包括app的基本信息和扫描状态信息。
	ScanSet []*AppScanSet `json:"ScanSet,omitempty" name:"ScanSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeScanInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScanInstancesResponseParams `json:"Response"`
}

func (r *DescribeScanInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanResultsRequestParams struct {
	// 任务唯一标识
	ItemId *string `json:"ItemId,omitempty" name:"ItemId"`

	// 批量查询一个或者多个app的扫描结果，如果不传表示查询该任务下所提交的所有app
	AppMd5s []*string `json:"AppMd5s,omitempty" name:"AppMd5s"`
}

type DescribeScanResultsRequest struct {
	*tchttp.BaseRequest
	
	// 任务唯一标识
	ItemId *string `json:"ItemId,omitempty" name:"ItemId"`

	// 批量查询一个或者多个app的扫描结果，如果不传表示查询该任务下所提交的所有app
	AppMd5s []*string `json:"AppMd5s,omitempty" name:"AppMd5s"`
}

func (r *DescribeScanResultsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanResultsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ItemId")
	delete(f, "AppMd5s")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScanResultsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanResultsResponseParams struct {
	// 批量扫描的app结果集
	ScanSet []*ScanSetInfo `json:"ScanSet,omitempty" name:"ScanSet"`

	// 批量扫描结果的个数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeScanResultsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScanResultsResponseParams `json:"Response"`
}

func (r *DescribeScanResultsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanResultsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeShieldInstancesRequestParams struct {
	// 支持通过app名称，app包名，加固的服务版本，提交的渠道进行筛选。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 数量限制，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 可以提供ItemId数组来查询一个或者多个结果。注意不可以同时指定ItemIds和Filters。
	ItemIds []*string `json:"ItemIds,omitempty" name:"ItemIds"`

	// 按某个字段排序，目前仅支持TaskTime排序。
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 升序（asc）还是降序（desc），默认：desc。
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

type DescribeShieldInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 支持通过app名称，app包名，加固的服务版本，提交的渠道进行筛选。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 数量限制，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 可以提供ItemId数组来查询一个或者多个结果。注意不可以同时指定ItemIds和Filters。
	ItemIds []*string `json:"ItemIds,omitempty" name:"ItemIds"`

	// 按某个字段排序，目前仅支持TaskTime排序。
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 升序（asc）还是降序（desc），默认：desc。
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

func (r *DescribeShieldInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShieldInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ItemIds")
	delete(f, "OrderField")
	delete(f, "OrderDirection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeShieldInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeShieldInstancesResponseParams struct {
	// 符合要求的app数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 一个关于app详细信息的结构体，主要包括app的基本信息和加固信息。
	AppSet []*AppSetInfo `json:"AppSet,omitempty" name:"AppSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeShieldInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeShieldInstancesResponseParams `json:"Response"`
}

func (r *DescribeShieldInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShieldInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeShieldPlanInstanceRequestParams struct {
	// 资源id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 服务类别id
	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`
}

type DescribeShieldPlanInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 资源id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 服务类别id
	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`
}

func (r *DescribeShieldPlanInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShieldPlanInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceId")
	delete(f, "Pid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeShieldPlanInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeShieldPlanInstanceResponseParams struct {
	// 绑定资源信息
	BindInfo *BindInfo `json:"BindInfo,omitempty" name:"BindInfo"`

	// 加固策略信息
	ShieldPlanInfo *ShieldPlanInfo `json:"ShieldPlanInfo,omitempty" name:"ShieldPlanInfo"`

	// 加固资源信息
	ResourceServiceInfo *ResourceServiceInfo `json:"ResourceServiceInfo,omitempty" name:"ResourceServiceInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeShieldPlanInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeShieldPlanInstanceResponseParams `json:"Response"`
}

func (r *DescribeShieldPlanInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShieldPlanInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeShieldResultRequestParams struct {
	// 任务唯一标识
	ItemId *string `json:"ItemId,omitempty" name:"ItemId"`
}

type DescribeShieldResultRequest struct {
	*tchttp.BaseRequest
	
	// 任务唯一标识
	ItemId *string `json:"ItemId,omitempty" name:"ItemId"`
}

func (r *DescribeShieldResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShieldResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ItemId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeShieldResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeShieldResultResponseParams struct {
	// 任务状态: 0-请返回,1-已完成,2-处理中,3-处理出错,4-处理超时
	TaskStatus *uint64 `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// app加固前的详细信息
	AppDetailInfo *AppDetailInfo `json:"AppDetailInfo,omitempty" name:"AppDetailInfo"`

	// app加固后的详细信息
	ShieldInfo *ShieldInfo `json:"ShieldInfo,omitempty" name:"ShieldInfo"`

	// 状态描述
	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

	// 状态指引
	StatusRef *string `json:"StatusRef,omitempty" name:"StatusRef"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeShieldResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeShieldResultResponseParams `json:"Response"`
}

func (r *DescribeShieldResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShieldResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUrlDetectionResultRequestParams struct {
	// 查询的网址
	Url *string `json:"Url,omitempty" name:"Url"`
}

type DescribeUrlDetectionResultRequest struct {
	*tchttp.BaseRequest
	
	// 查询的网址
	Url *string `json:"Url,omitempty" name:"Url"`
}

func (r *DescribeUrlDetectionResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUrlDetectionResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Url")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUrlDetectionResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUrlDetectionResultResponseParams struct {
	// [查询结果]查询结果；枚举值：0 查询成功，否则查询失败
	ResultCode *int64 `json:"ResultCode,omitempty" name:"ResultCode"`

	// [固定信息]响应协议版本号
	RespVer *int64 `json:"RespVer,omitempty" name:"RespVer"`

	// [查询结果]url恶意状态
	// 枚举值：
	// 0-1：未知，访问暂无风险。
	// 2 ：	风险网址，具体的恶意类型定义参考恶意大类EvilClass字段。
	// 3-4：安全，访问无风险。
	// 
	// 注意：查询结果EvilClass字段在Urltype=2时，才有意义。
	UrlType *int64 `json:"UrlType,omitempty" name:"UrlType"`

	// [查询结果]url恶意类型大类:{
	//     "1": "社工欺诈（仿冒、账号钓鱼、中奖诈骗）",
	//     "2": "信息诈骗（虚假充值、虚假兼职、虚假金融投资、虚假信用卡代办、网络赌博诈骗）",
	//     "3": "虚假销售（男女保健美容减肥产品、电子产品、虚假广告、违法销售）",
	//     "4": "恶意文件（病毒文件，木马文件，恶意apk文件的下载链接以及站点，挂马网站）",
	//     "5": "博彩网站（博彩网站，在线赌博网站）",
	//     "6": "色情网站（涉嫌传播色情内容，提供色情服务的网站）"
	//   }
	EvilClass *int64 `json:"EvilClass,omitempty" name:"EvilClass"`

	// 该字段暂为空
	EvilType *int64 `json:"EvilType,omitempty" name:"EvilType"`

	// 该字段暂为空
	Level *int64 `json:"Level,omitempty" name:"Level"`

	// [查询详情]url检出时间；时间戳
	DetectTime *int64 `json:"DetectTime,omitempty" name:"DetectTime"`

	// 该字段暂为空
	Wording *string `json:"Wording,omitempty" name:"Wording"`

	// 该字段暂为空
	WordingTitle *string `json:"WordingTitle,omitempty" name:"WordingTitle"`

	// [查询结果]url恶意状态说明；为UrlType字段值对应的说明
	UrlTypeDesc *string `json:"UrlTypeDesc,omitempty" name:"UrlTypeDesc"`

	// [查询结果]url恶意大类说明；为EvilClass字段值对应的说明
	EvilClassDesc *string `json:"EvilClassDesc,omitempty" name:"EvilClassDesc"`

	// 该字段暂为空
	EvilTypeDesc *string `json:"EvilTypeDesc,omitempty" name:"EvilTypeDesc"`

	// 该字段暂为空
	LevelDesc *string `json:"LevelDesc,omitempty" name:"LevelDesc"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUrlDetectionResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUrlDetectionResultResponseParams `json:"Response"`
}

func (r *DescribeUrlDetectionResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUrlDetectionResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserBaseInfoInstanceRequestParams struct {

}

type DescribeUserBaseInfoInstanceRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeUserBaseInfoInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserBaseInfoInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserBaseInfoInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserBaseInfoInstanceResponseParams struct {
	// 用户uin信息
	UserUin *uint64 `json:"UserUin,omitempty" name:"UserUin"`

	// 用户APPID信息
	UserAppid *uint64 `json:"UserAppid,omitempty" name:"UserAppid"`

	// 系统时间戳
	TimeStamp *uint64 `json:"TimeStamp,omitempty" name:"TimeStamp"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUserBaseInfoInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserBaseInfoInstanceResponseParams `json:"Response"`
}

func (r *DescribeUserBaseInfoInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserBaseInfoInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 需要过滤的字段
	Name *string `json:"Name,omitempty" name:"Name"`

	// 需要过滤字段的值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type OptPluginListItem struct {
	// 非广告类型
	PluginType *string `json:"PluginType,omitempty" name:"PluginType"`

	// 非广告插件名称
	PluginName *string `json:"PluginName,omitempty" name:"PluginName"`

	// 非广告插件描述
	PluginDesc *string `json:"PluginDesc,omitempty" name:"PluginDesc"`
}

type PlanDetailInfo struct {
	// 默认策略，1为默认，0为非默认
	IsDefault *uint64 `json:"IsDefault,omitempty" name:"IsDefault"`

	// 策略id
	PlanId *uint64 `json:"PlanId,omitempty" name:"PlanId"`

	// 策略名称
	PlanName *string `json:"PlanName,omitempty" name:"PlanName"`

	// 策略信息
	PlanInfo *PlanInfo `json:"PlanInfo,omitempty" name:"PlanInfo"`
}

type PlanInfo struct {
	// apk大小优化，0关闭，1开启
	ApkSizeOpt *uint64 `json:"ApkSizeOpt,omitempty" name:"ApkSizeOpt"`

	// Dex加固，0关闭，1开启
	Dex *uint64 `json:"Dex,omitempty" name:"Dex"`

	// So加固，0关闭，1开启
	So *uint64 `json:"So,omitempty" name:"So"`

	// 数据收集，0关闭，1开启
	Bugly *uint64 `json:"Bugly,omitempty" name:"Bugly"`

	// 防止重打包，0关闭，1开启
	AntiRepack *uint64 `json:"AntiRepack,omitempty" name:"AntiRepack"`

	// Dex分离，0关闭，1开启
	SeperateDex *uint64 `json:"SeperateDex,omitempty" name:"SeperateDex"`

	// 内存保护，0关闭，1开启
	Db *uint64 `json:"Db,omitempty" name:"Db"`

	// Dex签名校验，0关闭，1开启
	DexSig *uint64 `json:"DexSig,omitempty" name:"DexSig"`

	// So文件信息
	SoInfo *SoInfo `json:"SoInfo,omitempty" name:"SoInfo"`

	// vmp，0关闭，1开启
	AntiVMP *uint64 `json:"AntiVMP,omitempty" name:"AntiVMP"`

	// 保护so的强度，
	SoType []*string `json:"SoType,omitempty" name:"SoType"`

	// 防日志泄漏，0关闭，1开启
	AntiLogLeak *uint64 `json:"AntiLogLeak,omitempty" name:"AntiLogLeak"`

	// root检测，0关闭，1开启
	AntiQemuRoot *uint64 `json:"AntiQemuRoot,omitempty" name:"AntiQemuRoot"`

	// 资源防篡改，0关闭，1开启
	AntiAssets *uint64 `json:"AntiAssets,omitempty" name:"AntiAssets"`

	// 防止截屏，0关闭，1开启
	AntiScreenshot *uint64 `json:"AntiScreenshot,omitempty" name:"AntiScreenshot"`

	// SSL证书防窃取，0关闭，1开启
	AntiSSL *uint64 `json:"AntiSSL,omitempty" name:"AntiSSL"`
}

type PluginInfo struct {
	// 插件类型，分别为 1-通知栏广告，2-积分墙广告，3-banner广告，4- 悬浮窗图标广告，5-精品推荐列表广告, 6-插播广告
	PluginType *uint64 `json:"PluginType,omitempty" name:"PluginType"`

	// 插件名称
	PluginName *string `json:"PluginName,omitempty" name:"PluginName"`

	// 插件描述
	PluginDesc *string `json:"PluginDesc,omitempty" name:"PluginDesc"`
}

type PluginListItem struct {
	// 数字类型，分别为 1-通知栏广告，2-积分墙广告，3-banner广告，4- 悬浮窗图标广告，5-精品推荐列表广告, 6-插播广告
	PluginType *string `json:"PluginType,omitempty" name:"PluginType"`

	// 广告插件名称
	PluginName *string `json:"PluginName,omitempty" name:"PluginName"`

	// 广告插件描述
	PluginDesc *string `json:"PluginDesc,omitempty" name:"PluginDesc"`
}

type ResourceInfo struct {
	// 用户购买的资源id，全局唯一
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 资源的pid，MTP加固-12767，应用加固-12750 MTP反作弊-12766 源代码混淆-12736
	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`

	// 购买时间戳
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 到期时间戳
	ExpireTime *uint64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 0-未绑定，1-已绑定
	IsBind *int64 `json:"IsBind,omitempty" name:"IsBind"`

	// 用户绑定app的基本信息
	BindInfo *BindInfo `json:"BindInfo,omitempty" name:"BindInfo"`

	// 资源名称，如应用加固，漏洞扫描
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
}

type ResourceServiceInfo struct {
	// 创建时间戳
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 到期时间戳
	ExpireTime *uint64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 资源名称，如应用加固，源码混淆
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
}

type ResultListItem struct {
	// banner广告软件标记，分别为-1-不确定，0-否，1-是
	Banner *string `json:"Banner,omitempty" name:"Banner"`

	// 精品推荐列表广告标记，分别为-1-不确定，0-否，1-是
	BoutiqueRecommand *string `json:"BoutiqueRecommand,omitempty" name:"BoutiqueRecommand"`

	// 悬浮窗图标广告标记,分别为-1-不确定，0-否，1-是
	FloatWindows *string `json:"FloatWindows,omitempty" name:"FloatWindows"`

	// 积分墙广告软件标记，分别为 -1 -不确定，0-否，1-是
	IntegralWall *string `json:"IntegralWall,omitempty" name:"IntegralWall"`

	// 安装包的md5
	Md5 *string `json:"Md5,omitempty" name:"Md5"`

	// 通知栏广告软件标记，分别为-1-不确定，0-否，1-是
	NotifyBar *string `json:"NotifyBar,omitempty" name:"NotifyBar"`

	// 1表示官方，0表示非官方
	Official *string `json:"Official,omitempty" name:"Official"`

	// 广告插件结果列表
	PluginList []*PluginListItem `json:"PluginList,omitempty" name:"PluginList"`

	// 非广告插件结果列表(SDK、风险插件等)
	OptPluginList []*OptPluginListItem `json:"OptPluginList,omitempty" name:"OptPluginList"`

	// 数字类型，分别为0-未知， 1-安全软件，2-风险软件，3-病毒软件
	SafeType *string `json:"SafeType,omitempty" name:"SafeType"`

	// Session id，合作方可以用来区分回调数据，需要唯一。
	Sid *string `json:"Sid,omitempty" name:"Sid"`

	// 安装包名称
	SoftName *string `json:"SoftName,omitempty" name:"SoftName"`

	// 插播广告软件标记，取值：-1 不确定，0否， 1 是
	Spot *string `json:"Spot,omitempty" name:"Spot"`

	// 病毒名称，utf8编码
	VirusName *string `json:"VirusName,omitempty" name:"VirusName"`

	// 病毒描述，utf8编码
	VirusDesc *string `json:"VirusDesc,omitempty" name:"VirusDesc"`

	// 二次打包状态：0-表示默认；1-表示二次
	RepackageStatus *string `json:"RepackageStatus,omitempty" name:"RepackageStatus"`

	// 应用错误码：0、1-表示正常；                  
	// 
	// 2表示System Error(engine analysis error).
	// 
	// 3表示App analysis error, please confirm it.
	// 
	// 4表示App have not cert, please confirm it.
	// 
	// 5表示App size is zero, please confirm it.
	// 
	// 6表示App have not package name, please confirm it.
	// 
	// 7表示App build time is empty, please confirm it.
	// 
	// 8表示App have not valid cert, please confirm it.
	// 
	// 99表示Other error.
	// 
	// 1000表示App downloadlink download fail, please confirm it.
	// 
	// 1001表示APP md5 different between real md5, please confirm it.
	// 
	// 1002表示App md5 uncollect, please offer downloadlink.
	Errno *string `json:"Errno,omitempty" name:"Errno"`

	// 对应errno的错误信息描述
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
}

type ScanInfo struct {
	// 任务处理完成后的反向通知回调地址,批量提交app每扫描完成一个会通知一次,通知为POST请求，post信息{ItemId:
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// VULSCAN-漏洞扫描信息，VIRUSSCAN-返回病毒扫描信息， ADSCAN-广告扫描信息，PLUGINSCAN-插件扫描信息，PERMISSION-系统权限信息，SENSITIVE-敏感词信息，可以自由组合
	ScanTypes []*string `json:"ScanTypes,omitempty" name:"ScanTypes"`
}

type ScanPermissionInfo struct {
	// 系统权限
	Permission *string `json:"Permission,omitempty" name:"Permission"`
}

type ScanPermissionList struct {
	// 系统权限信息
	PermissionList []*ScanPermissionInfo `json:"PermissionList,omitempty" name:"PermissionList"`
}

type ScanSensitiveInfo struct {
	// 敏感词
	WordList []*string `json:"WordList,omitempty" name:"WordList"`

	// 敏感词对应的文件信息
	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`

	// 文件sha1值
	FileSha *string `json:"FileSha,omitempty" name:"FileSha"`
}

type ScanSensitiveList struct {
	// 敏感词列表
	SensitiveList []*ScanSensitiveInfo `json:"SensitiveList,omitempty" name:"SensitiveList"`
}

type ScanSetInfo struct {
	// 任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
	TaskStatus *uint64 `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// app信息
	AppDetailInfo *AppDetailInfo `json:"AppDetailInfo,omitempty" name:"AppDetailInfo"`

	// 病毒信息
	VirusInfo *VirusInfo `json:"VirusInfo,omitempty" name:"VirusInfo"`

	// 漏洞信息
	VulInfo *VulInfo `json:"VulInfo,omitempty" name:"VulInfo"`

	// 广告插件信息
	AdInfo *AdInfo `json:"AdInfo,omitempty" name:"AdInfo"`

	// 提交扫描的时间
	TaskTime *uint64 `json:"TaskTime,omitempty" name:"TaskTime"`

	// 状态码，成功返回0，失败返回错误码
	StatusCode *uint64 `json:"StatusCode,omitempty" name:"StatusCode"`

	// 状态描述
	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

	// 状态操作指引
	StatusRef *string `json:"StatusRef,omitempty" name:"StatusRef"`

	// 系统权限信息
	PermissionInfo *ScanPermissionList `json:"PermissionInfo,omitempty" name:"PermissionInfo"`

	// 敏感词列表
	SensitiveInfo *ScanSensitiveList `json:"SensitiveInfo,omitempty" name:"SensitiveInfo"`
}

type ServiceInfo struct {
	// 服务版本，基础版basic，专业版professional，企业版enterprise。
	ServiceEdition *string `json:"ServiceEdition,omitempty" name:"ServiceEdition"`

	// 任务处理完成后的反向通知回调地址，如果不需要通知请传递空字符串。通知为POST请求，post包体数据示例{"Response":{"ItemId":"4cdad8fb86f036b06bccb3f58971c306","ShieldCode":0,"ShieldMd5":"78701576793c4a5f04e1c9660de0aa0b","ShieldSize":11997354,"TaskStatus":1,"TaskTime":1539148141}}，调用方需要返回如下信息，{"Result":"ok","Reason":"xxxxx"}，如果Result字段值不等于ok会继续回调。
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 提交来源 YYB-应用宝 RDM-rdm MC-控制台 MAC_TOOL-mac工具 WIN_TOOL-window工具。
	SubmitSource *string `json:"SubmitSource,omitempty" name:"SubmitSource"`

	// 加固策略编号，如果不传则使用系统默认加固策略。如果指定的plan不存在会返回错误。
	PlanId *uint64 `json:"PlanId,omitempty" name:"PlanId"`
}

type ShieldInfo struct {
	// 加固结果的返回码
	ShieldCode *uint64 `json:"ShieldCode,omitempty" name:"ShieldCode"`

	// 加固后app的大小
	ShieldSize *uint64 `json:"ShieldSize,omitempty" name:"ShieldSize"`

	// 加固后app的md5
	ShieldMd5 *string `json:"ShieldMd5,omitempty" name:"ShieldMd5"`

	// 加固后的APP下载地址，该地址有效期为20分钟，请及时下载
	AppUrl *string `json:"AppUrl,omitempty" name:"AppUrl"`

	// 加固的提交时间
	TaskTime *uint64 `json:"TaskTime,omitempty" name:"TaskTime"`

	// 任务唯一标识
	ItemId *string `json:"ItemId,omitempty" name:"ItemId"`

	// 加固版本，basic基础版，professional专业版，enterprise企业版
	ServiceEdition *string `json:"ServiceEdition,omitempty" name:"ServiceEdition"`
}

type ShieldPlanInfo struct {
	// 加固策略数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 加固策略具体信息数组
	PlanSet []*PlanDetailInfo `json:"PlanSet,omitempty" name:"PlanSet"`
}

type SoInfo struct {
	// so文件列表
	SoFileNames []*string `json:"SoFileNames,omitempty" name:"SoFileNames"`
}

type VirusInfo struct {
	// 软件安全类型，分别为0-未知、 1-安全软件、2-风险软件、3-病毒软件
	SafeType *int64 `json:"SafeType,omitempty" name:"SafeType"`

	// 病毒名称， utf8编码，非病毒时值为空
	VirusName *string `json:"VirusName,omitempty" name:"VirusName"`

	// 病毒描述，utf8编码，非病毒时值为空
	VirusDesc *string `json:"VirusDesc,omitempty" name:"VirusDesc"`
}

type VulInfo struct {
	// 漏洞列表
	VulList []*VulList `json:"VulList,omitempty" name:"VulList"`

	// 漏洞文件评分
	VulFileScore *uint64 `json:"VulFileScore,omitempty" name:"VulFileScore"`
}

type VulList struct {
	// 漏洞id
	VulId *string `json:"VulId,omitempty" name:"VulId"`

	// 漏洞名称
	VulName *string `json:"VulName,omitempty" name:"VulName"`

	// 漏洞代码
	VulCode *string `json:"VulCode,omitempty" name:"VulCode"`

	// 漏洞描述
	VulDesc *string `json:"VulDesc,omitempty" name:"VulDesc"`

	// 漏洞解决方案
	VulSolution *string `json:"VulSolution,omitempty" name:"VulSolution"`

	// 漏洞来源类别，0默认自身，1第三方插件
	VulSrcType *int64 `json:"VulSrcType,omitempty" name:"VulSrcType"`

	// 漏洞位置
	VulFilepath *string `json:"VulFilepath,omitempty" name:"VulFilepath"`

	// 风险级别：1 低风险 ；2中等风险；3 高风险
	RiskLevel *uint64 `json:"RiskLevel,omitempty" name:"RiskLevel"`
}