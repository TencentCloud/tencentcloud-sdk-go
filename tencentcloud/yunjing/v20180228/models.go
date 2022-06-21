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

package v20180228

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Account struct {
	// 唯一ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 云镜客户端唯一Uuid
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 主机内网IP。
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// 主机名称。
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// 帐号名。
	Username *string `json:"Username,omitempty" name:"Username"`

	// 帐号所属组。
	Groups *string `json:"Groups,omitempty" name:"Groups"`

	// 帐号类型。
	// <li>ORDINARY：普通帐号</li>
	// <li>SUPPER：超级管理员帐号</li>
	Privilege *string `json:"Privilege,omitempty" name:"Privilege"`

	// 帐号创建时间。
	AccountCreateTime *string `json:"AccountCreateTime,omitempty" name:"AccountCreateTime"`

	// 帐号最后登录时间。
	LastLoginTime *string `json:"LastLoginTime,omitempty" name:"LastLoginTime"`
}

type AccountStatistics struct {
	// 用户名。
	Username *string `json:"Username,omitempty" name:"Username"`

	// 主机数量。
	MachineNum *uint64 `json:"MachineNum,omitempty" name:"MachineNum"`
}

// Predefined struct for user
type AddLoginWhiteListRequestParams struct {
	// 白名单规则
	Rules *LoginWhiteListsRule `json:"Rules,omitempty" name:"Rules"`
}

type AddLoginWhiteListRequest struct {
	*tchttp.BaseRequest
	
	// 白名单规则
	Rules *LoginWhiteListsRule `json:"Rules,omitempty" name:"Rules"`
}

func (r *AddLoginWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddLoginWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Rules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddLoginWhiteListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddLoginWhiteListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddLoginWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *AddLoginWhiteListResponseParams `json:"Response"`
}

func (r *AddLoginWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddLoginWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddMachineTagRequestParams struct {
	// 云服务器ID
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`

	// 标签ID
	TagId *uint64 `json:"TagId,omitempty" name:"TagId"`

	// 云服务器地区
	MRegion *string `json:"MRegion,omitempty" name:"MRegion"`

	// 云服务器类型(CVM|BM)
	MArea *string `json:"MArea,omitempty" name:"MArea"`
}

type AddMachineTagRequest struct {
	*tchttp.BaseRequest
	
	// 云服务器ID
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`

	// 标签ID
	TagId *uint64 `json:"TagId,omitempty" name:"TagId"`

	// 云服务器地区
	MRegion *string `json:"MRegion,omitempty" name:"MRegion"`

	// 云服务器类型(CVM|BM)
	MArea *string `json:"MArea,omitempty" name:"MArea"`
}

func (r *AddMachineTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddMachineTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Quuid")
	delete(f, "TagId")
	delete(f, "MRegion")
	delete(f, "MArea")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddMachineTagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddMachineTagResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddMachineTagResponse struct {
	*tchttp.BaseResponse
	Response *AddMachineTagResponseParams `json:"Response"`
}

func (r *AddMachineTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddMachineTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AgentVul struct {
	// 漏洞ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 主机IP。
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// 漏洞名称。
	VulName *string `json:"VulName,omitempty" name:"VulName"`

	// 漏洞危害等级。
	// <li>HIGH：高危</li>
	// <li>MIDDLE：中危</li>
	// <li>LOW：低危</li>
	// <li>NOTICE：提示</li>
	VulLevel *string `json:"VulLevel,omitempty" name:"VulLevel"`

	// 最后扫描时间。
	LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`

	// 漏洞描述。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 漏洞种类ID。
	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`

	// 漏洞状态。
	// <li>UN_OPERATED : 待处理</li>
	// <li>FIXED : 已修复</li>
	VulStatus *string `json:"VulStatus,omitempty" name:"VulStatus"`
}

type BashEvent struct {
	// ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 云镜ID
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 主机ID
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`

	// 主机内网IP
	Hostip *string `json:"Hostip,omitempty" name:"Hostip"`

	// 执行用户名
	User *string `json:"User,omitempty" name:"User"`

	// 平台类型
	Platform *uint64 `json:"Platform,omitempty" name:"Platform"`

	// 执行命令
	BashCmd *string `json:"BashCmd,omitempty" name:"BashCmd"`

	// 规则ID
	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`

	// 规则名称
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 规则等级
	RuleLevel *uint64 `json:"RuleLevel,omitempty" name:"RuleLevel"`

	// 处理状态
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 发生时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 主机名
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
}

type BashRule struct {
	// 规则ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 客户端ID
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 规则名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 危险等级(1: 高危 2:中危 3: 低危)
	Level *uint64 `json:"Level,omitempty" name:"Level"`

	// 正则表达式
	Rule *string `json:"Rule,omitempty" name:"Rule"`

	// 规则描述
	Decription *string `json:"Decription,omitempty" name:"Decription"`

	// 操作人
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 是否全局规则
	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`

	// 状态 (0: 有效 1: 无效)
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// 主机IP
	Hostip *string `json:"Hostip,omitempty" name:"Hostip"`
}

type BruteAttack struct {
	// 事件ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 主机IP。
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// 破解事件状态
	// <li>BRUTEATTACK_FAIL_ACCOUNT： 暴力破解事件-失败(存在帐号)  </li>
	// <li>BRUTEATTACK_FAIL_NOACCOUNT：暴力破解事件-失败(帐号不存在)</li>
	// <li>BRUTEATTACK_SUCCESS：暴力破解事件-成功</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 用户名称。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 城市ID。
	City *uint64 `json:"City,omitempty" name:"City"`

	// 国家ID。
	Country *uint64 `json:"Country,omitempty" name:"Country"`

	// 省份ID。
	Province *uint64 `json:"Province,omitempty" name:"Province"`

	// 来源IP。
	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`

	// 尝试破解次数。
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 发生时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 主机名称。
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// 云镜客户端唯一标识UUID。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 是否专业版。
	IsProVersion *bool `json:"IsProVersion,omitempty" name:"IsProVersion"`

	// 阻断状态。
	BanStatus *string `json:"BanStatus,omitempty" name:"BanStatus"`

	// 机器UUID
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

type ChargePrepaid struct {
	// 购买实例的时长，单位：月。取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36。
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 自动续费标识。取值范围：
	// <li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费</li>
	// <li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费</li>
	// <li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费</li>
	// 
	// 默认取值：NOTIFY_AND_MANUAL_RENEW。若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

// Predefined struct for user
type CloseProVersionRequestParams struct {
	// 主机唯一标识Uuid。
	// 黑石的InstanceId，CVM的Uuid
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

type CloseProVersionRequest struct {
	*tchttp.BaseRequest
	
	// 主机唯一标识Uuid。
	// 黑石的InstanceId，CVM的Uuid
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *CloseProVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseProVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Quuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseProVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseProVersionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CloseProVersionResponse struct {
	*tchttp.BaseResponse
	Response *CloseProVersionResponseParams `json:"Response"`
}

func (r *CloseProVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseProVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Component struct {
	// 唯一ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 云镜客户端唯一Uuid。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 主机内网IP。
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// 主机名。
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// 组件版本号。
	ComponentVersion *string `json:"ComponentVersion,omitempty" name:"ComponentVersion"`

	// 组件类型。
	// <li>SYSTEM：系统组件</li>
	// <li>WEB：Web组件</li>
	ComponentType *string `json:"ComponentType,omitempty" name:"ComponentType"`

	// 组件名称。
	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`

	// 组件检测更新时间。
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type ComponentStatistics struct {
	// 组件ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 主机数量。
	MachineNum *uint64 `json:"MachineNum,omitempty" name:"MachineNum"`

	// 组件名称。
	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`

	// 组件类型。
	// <li>WEB：Web组件</li>
	// <li>SYSTEM：系统组件</li>
	ComponentType *string `json:"ComponentType,omitempty" name:"ComponentType"`

	// 组件描述。
	Description *string `json:"Description,omitempty" name:"Description"`
}

// Predefined struct for user
type CreateBaselineStrategyRequestParams struct {
	// 策略名称
	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`

	// 检测周期, 表示每隔多少天进行检测.示例: 2, 表示每2天进行检测一次.
	ScanCycle *uint64 `json:"ScanCycle,omitempty" name:"ScanCycle"`

	// 定期检测时间，该时间下发扫描. 示例:“22:00”, 表示在22:00下发检测
	ScanAt *string `json:"ScanAt,omitempty" name:"ScanAt"`

	// 该策略下选择的基线id数组. 示例: [1,3,5,7]
	CategoryIds []*uint64 `json:"CategoryIds,omitempty" name:"CategoryIds"`

	// 扫描范围是否全部服务器, 1:是  0:否, 为1则为全部专业版主机
	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`

	// 云主机类型：“CVM”：虚拟主机，"BMS"：裸金属，"ECM"：边缘计算主机
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// 主机地域. 示例: "ap-bj"
	RegionCode *string `json:"RegionCode,omitempty" name:"RegionCode"`

	// 主机id数组. 示例: ["quuid1","quuid2"]
	Quuids []*string `json:"Quuids,omitempty" name:"Quuids"`
}

type CreateBaselineStrategyRequest struct {
	*tchttp.BaseRequest
	
	// 策略名称
	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`

	// 检测周期, 表示每隔多少天进行检测.示例: 2, 表示每2天进行检测一次.
	ScanCycle *uint64 `json:"ScanCycle,omitempty" name:"ScanCycle"`

	// 定期检测时间，该时间下发扫描. 示例:“22:00”, 表示在22:00下发检测
	ScanAt *string `json:"ScanAt,omitempty" name:"ScanAt"`

	// 该策略下选择的基线id数组. 示例: [1,3,5,7]
	CategoryIds []*uint64 `json:"CategoryIds,omitempty" name:"CategoryIds"`

	// 扫描范围是否全部服务器, 1:是  0:否, 为1则为全部专业版主机
	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`

	// 云主机类型：“CVM”：虚拟主机，"BMS"：裸金属，"ECM"：边缘计算主机
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// 主机地域. 示例: "ap-bj"
	RegionCode *string `json:"RegionCode,omitempty" name:"RegionCode"`

	// 主机id数组. 示例: ["quuid1","quuid2"]
	Quuids []*string `json:"Quuids,omitempty" name:"Quuids"`
}

func (r *CreateBaselineStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBaselineStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StrategyName")
	delete(f, "ScanCycle")
	delete(f, "ScanAt")
	delete(f, "CategoryIds")
	delete(f, "IsGlobal")
	delete(f, "MachineType")
	delete(f, "RegionCode")
	delete(f, "Quuids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBaselineStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBaselineStrategyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateBaselineStrategyResponse struct {
	*tchttp.BaseResponse
	Response *CreateBaselineStrategyResponseParams `json:"Response"`
}

func (r *CreateBaselineStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBaselineStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOpenPortTaskRequestParams struct {
	// 云镜客户端唯一Uuid。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type CreateOpenPortTaskRequest struct {
	*tchttp.BaseRequest
	
	// 云镜客户端唯一Uuid。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *CreateOpenPortTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOpenPortTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOpenPortTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOpenPortTaskResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateOpenPortTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateOpenPortTaskResponseParams `json:"Response"`
}

func (r *CreateOpenPortTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOpenPortTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProcessTaskRequestParams struct {
	// 云镜客户端唯一Uuid。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type CreateProcessTaskRequest struct {
	*tchttp.BaseRequest
	
	// 云镜客户端唯一Uuid。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *CreateProcessTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProcessTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProcessTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProcessTaskResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateProcessTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateProcessTaskResponseParams `json:"Response"`
}

func (r *CreateProcessTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProcessTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUsualLoginPlacesRequestParams struct {
	// 云镜客户端UUID数组。
	Uuids []*string `json:"Uuids,omitempty" name:"Uuids"`

	// 登录地域信息数组。
	Places []*Place `json:"Places,omitempty" name:"Places"`
}

type CreateUsualLoginPlacesRequest struct {
	*tchttp.BaseRequest
	
	// 云镜客户端UUID数组。
	Uuids []*string `json:"Uuids,omitempty" name:"Uuids"`

	// 登录地域信息数组。
	Places []*Place `json:"Places,omitempty" name:"Places"`
}

func (r *CreateUsualLoginPlacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUsualLoginPlacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuids")
	delete(f, "Places")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUsualLoginPlacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUsualLoginPlacesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateUsualLoginPlacesResponse struct {
	*tchttp.BaseResponse
	Response *CreateUsualLoginPlacesResponseParams `json:"Response"`
}

func (r *CreateUsualLoginPlacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUsualLoginPlacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DefendAttackLog struct {
	// 日志ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 客户端ID
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 来源IP
	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`

	// 来源端口
	SrcPort *uint64 `json:"SrcPort,omitempty" name:"SrcPort"`

	// 攻击方式
	HttpMethod *string `json:"HttpMethod,omitempty" name:"HttpMethod"`

	// 攻击描述
	HttpCgi *string `json:"HttpCgi,omitempty" name:"HttpCgi"`

	// 攻击参数
	HttpParam *string `json:"HttpParam,omitempty" name:"HttpParam"`

	// 威胁类型
	VulType *string `json:"VulType,omitempty" name:"VulType"`

	// 攻击时间
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 目标服务器IP
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// 目标服务器名称
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// 目标IP
	DstIp *string `json:"DstIp,omitempty" name:"DstIp"`

	// 目标端口
	DstPort *uint64 `json:"DstPort,omitempty" name:"DstPort"`

	// 攻击内容
	HttpContent *string `json:"HttpContent,omitempty" name:"HttpContent"`
}

// Predefined struct for user
type DeleteAttackLogsRequestParams struct {
	// 日志ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

type DeleteAttackLogsRequest struct {
	*tchttp.BaseRequest
	
	// 日志ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteAttackLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAttackLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAttackLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAttackLogsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteAttackLogsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAttackLogsResponseParams `json:"Response"`
}

func (r *DeleteAttackLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAttackLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBashEventsRequestParams struct {
	// ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

type DeleteBashEventsRequest struct {
	*tchttp.BaseRequest
	
	// ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteBashEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBashEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBashEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBashEventsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteBashEventsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteBashEventsResponseParams `json:"Response"`
}

func (r *DeleteBashEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBashEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBashRulesRequestParams struct {
	// ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

type DeleteBashRulesRequest struct {
	*tchttp.BaseRequest
	
	// ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteBashRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBashRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBashRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBashRulesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteBashRulesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteBashRulesResponseParams `json:"Response"`
}

func (r *DeleteBashRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBashRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBruteAttacksRequestParams struct {
	// 暴力破解事件Id数组。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

type DeleteBruteAttacksRequest struct {
	*tchttp.BaseRequest
	
	// 暴力破解事件Id数组。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteBruteAttacksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBruteAttacksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBruteAttacksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBruteAttacksResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteBruteAttacksResponse struct {
	*tchttp.BaseResponse
	Response *DeleteBruteAttacksResponseParams `json:"Response"`
}

func (r *DeleteBruteAttacksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBruteAttacksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLoginWhiteListRequestParams struct {
	// 白名单ID
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

type DeleteLoginWhiteListRequest struct {
	*tchttp.BaseRequest
	
	// 白名单ID
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteLoginWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoginWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLoginWhiteListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLoginWhiteListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteLoginWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLoginWhiteListResponseParams `json:"Response"`
}

func (r *DeleteLoginWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoginWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMachineRequestParams struct {
	// 云镜客户端Uuid。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DeleteMachineRequest struct {
	*tchttp.BaseRequest
	
	// 云镜客户端Uuid。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DeleteMachineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMachineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMachineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMachineResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteMachineResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMachineResponseParams `json:"Response"`
}

func (r *DeleteMachineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMachineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMachineTagRequestParams struct {
	// 关联的标签ID
	Rid *uint64 `json:"Rid,omitempty" name:"Rid"`
}

type DeleteMachineTagRequest struct {
	*tchttp.BaseRequest
	
	// 关联的标签ID
	Rid *uint64 `json:"Rid,omitempty" name:"Rid"`
}

func (r *DeleteMachineTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMachineTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Rid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMachineTagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMachineTagResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteMachineTagResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMachineTagResponseParams `json:"Response"`
}

func (r *DeleteMachineTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMachineTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMaliciousRequestsRequestParams struct {
	// 恶意请求记录ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

type DeleteMaliciousRequestsRequest struct {
	*tchttp.BaseRequest
	
	// 恶意请求记录ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteMaliciousRequestsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMaliciousRequestsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMaliciousRequestsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMaliciousRequestsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteMaliciousRequestsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMaliciousRequestsResponseParams `json:"Response"`
}

func (r *DeleteMaliciousRequestsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMaliciousRequestsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMalwaresRequestParams struct {
	// 木马记录ID数组
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

type DeleteMalwaresRequest struct {
	*tchttp.BaseRequest
	
	// 木马记录ID数组
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteMalwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMalwaresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMalwaresRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMalwaresResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteMalwaresResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMalwaresResponseParams `json:"Response"`
}

func (r *DeleteMalwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMalwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNonlocalLoginPlacesRequestParams struct {
	// 异地登录事件ID数组。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

type DeleteNonlocalLoginPlacesRequest struct {
	*tchttp.BaseRequest
	
	// 异地登录事件ID数组。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteNonlocalLoginPlacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNonlocalLoginPlacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNonlocalLoginPlacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNonlocalLoginPlacesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteNonlocalLoginPlacesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNonlocalLoginPlacesResponseParams `json:"Response"`
}

func (r *DeleteNonlocalLoginPlacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNonlocalLoginPlacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrivilegeEventsRequestParams struct {
	// ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

type DeletePrivilegeEventsRequest struct {
	*tchttp.BaseRequest
	
	// ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeletePrivilegeEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrivilegeEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePrivilegeEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrivilegeEventsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeletePrivilegeEventsResponse struct {
	*tchttp.BaseResponse
	Response *DeletePrivilegeEventsResponseParams `json:"Response"`
}

func (r *DeletePrivilegeEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrivilegeEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrivilegeRulesRequestParams struct {
	// ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

type DeletePrivilegeRulesRequest struct {
	*tchttp.BaseRequest
	
	// ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeletePrivilegeRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrivilegeRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePrivilegeRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrivilegeRulesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeletePrivilegeRulesResponse struct {
	*tchttp.BaseResponse
	Response *DeletePrivilegeRulesResponseParams `json:"Response"`
}

func (r *DeletePrivilegeRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrivilegeRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReverseShellEventsRequestParams struct {
	// ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

type DeleteReverseShellEventsRequest struct {
	*tchttp.BaseRequest
	
	// ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteReverseShellEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReverseShellEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteReverseShellEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReverseShellEventsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteReverseShellEventsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteReverseShellEventsResponseParams `json:"Response"`
}

func (r *DeleteReverseShellEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReverseShellEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReverseShellRulesRequestParams struct {
	// ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

type DeleteReverseShellRulesRequest struct {
	*tchttp.BaseRequest
	
	// ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteReverseShellRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReverseShellRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteReverseShellRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReverseShellRulesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteReverseShellRulesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteReverseShellRulesResponseParams `json:"Response"`
}

func (r *DeleteReverseShellRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReverseShellRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTagsRequestParams struct {
	// 标签ID
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

type DeleteTagsRequest struct {
	*tchttp.BaseRequest
	
	// 标签ID
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTagsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteTagsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTagsResponseParams `json:"Response"`
}

func (r *DeleteTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUsualLoginPlacesRequestParams struct {
	// 云镜客户端Uuid
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 已添加常用登录地城市ID数组
	CityIds []*uint64 `json:"CityIds,omitempty" name:"CityIds"`
}

type DeleteUsualLoginPlacesRequest struct {
	*tchttp.BaseRequest
	
	// 云镜客户端Uuid
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 已添加常用登录地城市ID数组
	CityIds []*uint64 `json:"CityIds,omitempty" name:"CityIds"`
}

func (r *DeleteUsualLoginPlacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUsualLoginPlacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	delete(f, "CityIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUsualLoginPlacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUsualLoginPlacesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteUsualLoginPlacesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUsualLoginPlacesResponseParams `json:"Response"`
}

func (r *DeleteUsualLoginPlacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUsualLoginPlacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountStatisticsRequestParams struct {
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Username - String - 是否必填：否 - 帐号用户名</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeAccountStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Username - String - 是否必填：否 - 帐号用户名</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAccountStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountStatisticsResponseParams struct {
	// 帐号统计列表记录总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 帐号统计列表。
	AccountStatistics []*AccountStatistics `json:"AccountStatistics,omitempty" name:"AccountStatistics"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAccountStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccountStatisticsResponseParams `json:"Response"`
}

func (r *DescribeAccountStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountsRequestParams struct {
	// 云镜客户端唯一Uuid。Username和Uuid必填其一，使用Uuid表示，查询该主机下列表信息。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 云镜客户端唯一Uuid。Username和Uuid必填其一，使用Username表示，查询该用户名下列表信息。
	Username *string `json:"Username,omitempty" name:"Username"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Username - String - 是否必填：否 - 帐号名</li>
	// <li>Privilege - String - 是否必填：否 - 帐号类型（ORDINARY: 普通帐号 | SUPPER: 超级管理员帐号）</li>
	// <li>MachineIp - String - 是否必填：否 - 主机内网IP</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeAccountsRequest struct {
	*tchttp.BaseRequest
	
	// 云镜客户端唯一Uuid。Username和Uuid必填其一，使用Uuid表示，查询该主机下列表信息。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 云镜客户端唯一Uuid。Username和Uuid必填其一，使用Username表示，查询该用户名下列表信息。
	Username *string `json:"Username,omitempty" name:"Username"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Username - String - 是否必填：否 - 帐号名</li>
	// <li>Privilege - String - 是否必填：否 - 帐号类型（ORDINARY: 普通帐号 | SUPPER: 超级管理员帐号）</li>
	// <li>MachineIp - String - 是否必填：否 - 主机内网IP</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	delete(f, "Username")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountsResponseParams struct {
	// 帐号列表记录总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 帐号数据列表。
	Accounts []*Account `json:"Accounts,omitempty" name:"Accounts"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAccountsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccountsResponseParams `json:"Response"`
}

func (r *DescribeAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentVulsRequestParams struct {
	// 漏洞类型。
	// <li>WEB: Web应用漏洞</li>
	// <li>SYSTEM：系统组件漏洞</li>
	// <li>BASELINE：安全基线</li>
	VulType *string `json:"VulType,omitempty" name:"VulType"`

	// 客户端UUID。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Status - String - 是否必填：否 - 状态筛选（UN_OPERATED: 待处理 | FIXED：已修复）
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeAgentVulsRequest struct {
	*tchttp.BaseRequest
	
	// 漏洞类型。
	// <li>WEB: Web应用漏洞</li>
	// <li>SYSTEM：系统组件漏洞</li>
	// <li>BASELINE：安全基线</li>
	VulType *string `json:"VulType,omitempty" name:"VulType"`

	// 客户端UUID。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Status - String - 是否必填：否 - 状态筛选（UN_OPERATED: 待处理 | FIXED：已修复）
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAgentVulsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentVulsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VulType")
	delete(f, "Uuid")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgentVulsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentVulsResponseParams struct {
	// 记录总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 主机漏洞信息
	AgentVuls []*AgentVul `json:"AgentVuls,omitempty" name:"AgentVuls"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAgentVulsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAgentVulsResponseParams `json:"Response"`
}

func (r *DescribeAgentVulsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentVulsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmAttributeRequestParams struct {

}

type DescribeAlarmAttributeRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAlarmAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmAttributeResponseParams struct {
	// 防护软件离线告警状态：
	// <li>OPEN：告警已开启</li>
	// <li>CLOSE： 告警已关闭</li>
	Offline *string `json:"Offline,omitempty" name:"Offline"`

	// 发现木马告警状态：
	// <li>OPEN：告警已开启</li>
	// <li>CLOSE： 告警已关闭</li>
	Malware *string `json:"Malware,omitempty" name:"Malware"`

	// 发现异地登录告警状态：
	// <li>OPEN：告警已开启</li>
	// <li>CLOSE： 告警已关闭</li>
	NonlocalLogin *string `json:"NonlocalLogin,omitempty" name:"NonlocalLogin"`

	// 被暴力破解成功告警状态：
	// <li>OPEN：告警已开启</li>
	// <li>CLOSE： 告警已关闭</li>
	CrackSuccess *string `json:"CrackSuccess,omitempty" name:"CrackSuccess"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAlarmAttributeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlarmAttributeResponseParams `json:"Response"`
}

func (r *DescribeAlarmAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAttackLogInfoRequestParams struct {
	// 日志ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

type DescribeAttackLogInfoRequest struct {
	*tchttp.BaseRequest
	
	// 日志ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeAttackLogInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAttackLogInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAttackLogInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAttackLogInfoResponseParams struct {
	// 日志ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 主机ID
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`

	// 攻击来源端口
	SrcPort *uint64 `json:"SrcPort,omitempty" name:"SrcPort"`

	// 攻击来源IP
	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`

	// 攻击目标端口
	DstPort *uint64 `json:"DstPort,omitempty" name:"DstPort"`

	// 攻击目标IP
	DstIp *string `json:"DstIp,omitempty" name:"DstIp"`

	// 攻击方法
	HttpMethod *string `json:"HttpMethod,omitempty" name:"HttpMethod"`

	// 攻击目标主机
	HttpHost *string `json:"HttpHost,omitempty" name:"HttpHost"`

	// 攻击头信息
	HttpHead *string `json:"HttpHead,omitempty" name:"HttpHead"`

	// 攻击者浏览器标识
	HttpUserAgent *string `json:"HttpUserAgent,omitempty" name:"HttpUserAgent"`

	// 请求源
	HttpReferer *string `json:"HttpReferer,omitempty" name:"HttpReferer"`

	// 威胁类型
	VulType *string `json:"VulType,omitempty" name:"VulType"`

	// 攻击路径
	HttpCgi *string `json:"HttpCgi,omitempty" name:"HttpCgi"`

	// 攻击参数
	HttpParam *string `json:"HttpParam,omitempty" name:"HttpParam"`

	// 攻击时间
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 攻击内容
	HttpContent *string `json:"HttpContent,omitempty" name:"HttpContent"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAttackLogInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAttackLogInfoResponseParams `json:"Response"`
}

func (r *DescribeAttackLogInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAttackLogInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAttackLogsRequestParams struct {
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>HttpMethod - String - 是否必填：否 - 攻击方法(POST|GET)</li>
	// <li>MachineIp - String - 是否必填：否 - 主机内网IP</li>
	// <li>DateRange - String - 是否必填：否 - 时间范围(存储最近3个月的数据)，如最近一个月["2019-11-17", "2019-12-17"]</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 主机安全客户端ID
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 云主机机器ID
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

type DescribeAttackLogsRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>HttpMethod - String - 是否必填：否 - 攻击方法(POST|GET)</li>
	// <li>MachineIp - String - 是否必填：否 - 主机内网IP</li>
	// <li>DateRange - String - 是否必填：否 - 时间范围(存储最近3个月的数据)，如最近一个月["2019-11-17", "2019-12-17"]</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 主机安全客户端ID
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 云主机机器ID
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *DescribeAttackLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAttackLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Uuid")
	delete(f, "Quuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAttackLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAttackLogsResponseParams struct {
	// 日志列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttackLogs []*DefendAttackLog `json:"AttackLogs,omitempty" name:"AttackLogs"`

	// 总条数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAttackLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAttackLogsResponseParams `json:"Response"`
}

func (r *DescribeAttackLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAttackLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBashEventsRequestParams struct {
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 关键词(主机内网IP)</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeBashEventsRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 关键词(主机内网IP)</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeBashEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBashEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBashEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBashEventsResponseParams struct {
	// 总条数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 高危命令事件列表
	List []*BashEvent `json:"List,omitempty" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBashEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBashEventsResponseParams `json:"Response"`
}

func (r *DescribeBashEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBashEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBashRulesRequestParams struct {
	// 0-系统规则; 1-用户规则
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 关键字(规则名称)</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeBashRulesRequest struct {
	*tchttp.BaseRequest
	
	// 0-系统规则; 1-用户规则
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 关键字(规则名称)</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeBashRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBashRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBashRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBashRulesResponseParams struct {
	// 列表内容
	List []*BashRule `json:"List,omitempty" name:"List"`

	// 总条数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBashRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBashRulesResponseParams `json:"Response"`
}

func (r *DescribeBashRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBashRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBruteAttacksRequestParams struct {
	// 客户端唯一Uuid。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 -  查询关键字</li>
	// <li>Status - String - 是否必填：否 -  查询状态（FAILED：破解失败 |SUCCESS：破解成功）</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeBruteAttacksRequest struct {
	*tchttp.BaseRequest
	
	// 客户端唯一Uuid。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 -  查询关键字</li>
	// <li>Status - String - 是否必填：否 -  查询状态（FAILED：破解失败 |SUCCESS：破解成功）</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeBruteAttacksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBruteAttacksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBruteAttacksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBruteAttacksResponseParams struct {
	// 事件数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 暴力破解事件列表
	BruteAttacks []*BruteAttack `json:"BruteAttacks,omitempty" name:"BruteAttacks"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBruteAttacksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBruteAttacksResponseParams `json:"Response"`
}

func (r *DescribeBruteAttacksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBruteAttacksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComponentInfoRequestParams struct {
	// 组件ID。
	ComponentId *uint64 `json:"ComponentId,omitempty" name:"ComponentId"`
}

type DescribeComponentInfoRequest struct {
	*tchttp.BaseRequest
	
	// 组件ID。
	ComponentId *uint64 `json:"ComponentId,omitempty" name:"ComponentId"`
}

func (r *DescribeComponentInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComponentInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ComponentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeComponentInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComponentInfoResponseParams struct {
	// 组件ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 组件名称。
	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`

	// 组件类型。
	// <li>WEB：web组件</li>
	// <li>SYSTEM：系统组件</li>
	ComponentType *string `json:"ComponentType,omitempty" name:"ComponentType"`

	// 组件官网。
	Homepage *string `json:"Homepage,omitempty" name:"Homepage"`

	// 组件描述。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeComponentInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeComponentInfoResponseParams `json:"Response"`
}

func (r *DescribeComponentInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComponentInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComponentStatisticsRequestParams struct {
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// ComponentName - String - 是否必填：否 - 组件名称
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeComponentStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// ComponentName - String - 是否必填：否 - 组件名称
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeComponentStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComponentStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeComponentStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComponentStatisticsResponseParams struct {
	// 组件统计列表记录总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 组件统计列表数据数组。
	ComponentStatistics []*ComponentStatistics `json:"ComponentStatistics,omitempty" name:"ComponentStatistics"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeComponentStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeComponentStatisticsResponseParams `json:"Response"`
}

func (r *DescribeComponentStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComponentStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComponentsRequestParams struct {
	// 云镜客户端唯一Uuid。Uuid和ComponentId必填其一，使用Uuid表示，查询该主机列表信息。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 组件ID。Uuid和ComponentId必填其一，使用ComponentId表示，查询该组件列表信息。
	ComponentId *uint64 `json:"ComponentId,omitempty" name:"ComponentId"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>ComponentVersion - String - 是否必填：否 - 组件版本号</li>
	// <li>MachineIp - String - 是否必填：否 - 主机内网IP</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeComponentsRequest struct {
	*tchttp.BaseRequest
	
	// 云镜客户端唯一Uuid。Uuid和ComponentId必填其一，使用Uuid表示，查询该主机列表信息。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 组件ID。Uuid和ComponentId必填其一，使用ComponentId表示，查询该组件列表信息。
	ComponentId *uint64 `json:"ComponentId,omitempty" name:"ComponentId"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>ComponentVersion - String - 是否必填：否 - 组件版本号</li>
	// <li>MachineIp - String - 是否必填：否 - 主机内网IP</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeComponentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComponentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	delete(f, "ComponentId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeComponentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComponentsResponseParams struct {
	// 组件列表记录总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 组件列表数据。
	Components []*Component `json:"Components,omitempty" name:"Components"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeComponentsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeComponentsResponseParams `json:"Response"`
}

func (r *DescribeComponentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComponentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHistoryAccountsRequestParams struct {
	// 云镜客户端唯一Uuid。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Username - String - 是否必填：否 - 帐号名</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeHistoryAccountsRequest struct {
	*tchttp.BaseRequest
	
	// 云镜客户端唯一Uuid。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Username - String - 是否必填：否 - 帐号名</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeHistoryAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHistoryAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHistoryAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHistoryAccountsResponseParams struct {
	// 帐号变更历史列表记录总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 帐号变更历史数据数组。
	HistoryAccounts []*HistoryAccount `json:"HistoryAccounts,omitempty" name:"HistoryAccounts"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeHistoryAccountsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHistoryAccountsResponseParams `json:"Response"`
}

func (r *DescribeHistoryAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHistoryAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImpactedHostsRequestParams struct {
	// 漏洞种类ID。
	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Status - String - 是否必填：否 - 状态筛选（UN_OPERATED：待处理 | FIXED：已修复）</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeImpactedHostsRequest struct {
	*tchttp.BaseRequest
	
	// 漏洞种类ID。
	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Status - String - 是否必填：否 - 状态筛选（UN_OPERATED：待处理 | FIXED：已修复）</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeImpactedHostsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImpactedHostsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VulId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImpactedHostsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImpactedHostsResponseParams struct {
	// 记录总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 漏洞影响机器列表数组
	ImpactedHosts []*ImpactedHost `json:"ImpactedHosts,omitempty" name:"ImpactedHosts"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeImpactedHostsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImpactedHostsResponseParams `json:"Response"`
}

func (r *DescribeImpactedHostsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImpactedHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoginWhiteListRequestParams struct {
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 查询关键字 </li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeLoginWhiteListRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 查询关键字 </li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeLoginWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoginWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLoginWhiteListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoginWhiteListResponseParams struct {
	// 记录总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 异地登录白名单数组
	LoginWhiteLists []*LoginWhiteLists `json:"LoginWhiteLists,omitempty" name:"LoginWhiteLists"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLoginWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLoginWhiteListResponseParams `json:"Response"`
}

func (r *DescribeLoginWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoginWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMachineInfoRequestParams struct {
	// 云镜客户端唯一Uuid。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeMachineInfoRequest struct {
	*tchttp.BaseRequest
	
	// 云镜客户端唯一Uuid。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeMachineInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMachineInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMachineInfoResponseParams struct {
	// 机器ip。
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// 受云镜保护天数。
	ProtectDays *uint64 `json:"ProtectDays,omitempty" name:"ProtectDays"`

	// 操作系统。
	MachineOs *string `json:"MachineOs,omitempty" name:"MachineOs"`

	// 主机名称。
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// 在线状态。
	// <li>ONLINE： 在线</li>
	// <li>OFFLINE：离线</li>
	MachineStatus *string `json:"MachineStatus,omitempty" name:"MachineStatus"`

	// CVM或BM主机唯一标识。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 主机外网IP。
	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`

	// CVM或BM主机唯一Uuid。
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`

	// 云镜客户端唯一Uuid。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 是否开通专业版。
	// <li>true：是</li>
	// <li>false：否</li>
	IsProVersion *bool `json:"IsProVersion,omitempty" name:"IsProVersion"`

	// 专业版开通时间。
	ProVersionOpenDate *string `json:"ProVersionOpenDate,omitempty" name:"ProVersionOpenDate"`

	// 云主机类型。
	// <li>CVM: 虚拟主机</li>
	// <li>BM: 黑石物理机</li>
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// 机器所属地域。如：ap-guangzhou，ap-shanghai
	MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`

	// 主机状态。
	// <li>POSTPAY: 表示后付费，即按量计费  </li>
	// <li>PREPAY: 表示预付费，即包年包月</li>
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 免费木马剩余检测数量。
	FreeMalwaresLeft *uint64 `json:"FreeMalwaresLeft,omitempty" name:"FreeMalwaresLeft"`

	// 免费漏洞剩余检测数量。
	FreeVulsLeft *uint64 `json:"FreeVulsLeft,omitempty" name:"FreeVulsLeft"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeMachineInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMachineInfoResponseParams `json:"Response"`
}

func (r *DescribeMachineInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMachinesRequestParams struct {
	// 云主机类型。
	// <li>CVM：表示虚拟主机</li>
	// <li>BM:  表示黑石物理机</li>
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// 机器所属地域。如：ap-guangzhou，ap-shanghai
	MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 查询关键字 </li>
	// <li>Status - String - 是否必填：否 - 客户端在线状态（OFFLINE: 离线 | ONLINE: 在线 | UNINSTALLED：未安装）</li>
	// <li>Version - String  是否必填：否 - 当前防护版本（ PRO_VERSION：专业版 | BASIC_VERSION：基础版）</li>
	// 每个过滤条件只支持一个值，暂不支持多个值“或”关系查询
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeMachinesRequest struct {
	*tchttp.BaseRequest
	
	// 云主机类型。
	// <li>CVM：表示虚拟主机</li>
	// <li>BM:  表示黑石物理机</li>
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// 机器所属地域。如：ap-guangzhou，ap-shanghai
	MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 查询关键字 </li>
	// <li>Status - String - 是否必填：否 - 客户端在线状态（OFFLINE: 离线 | ONLINE: 在线 | UNINSTALLED：未安装）</li>
	// <li>Version - String  是否必填：否 - 当前防护版本（ PRO_VERSION：专业版 | BASIC_VERSION：基础版）</li>
	// 每个过滤条件只支持一个值，暂不支持多个值“或”关系查询
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeMachinesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachinesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MachineType")
	delete(f, "MachineRegion")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMachinesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMachinesResponseParams struct {
	// 主机列表
	Machines []*Machine `json:"Machines,omitempty" name:"Machines"`

	// 主机数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeMachinesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMachinesResponseParams `json:"Response"`
}

func (r *DescribeMachinesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachinesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMaliciousRequestsRequestParams struct {
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Status - String - 是否必填：否 - 状态筛选（UN_OPERATED: 待处理 | TRUSTED：已信任 | UN_TRUSTED：已取消信任）</li>
	// <li>Domain - String - 是否必填：否 - 恶意请求的域名</li>
	// <li>MachineIp - String - 是否必填：否 - 主机内网IP</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 云镜客户端唯一UUID。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeMaliciousRequestsRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Status - String - 是否必填：否 - 状态筛选（UN_OPERATED: 待处理 | TRUSTED：已信任 | UN_TRUSTED：已取消信任）</li>
	// <li>Domain - String - 是否必填：否 - 恶意请求的域名</li>
	// <li>MachineIp - String - 是否必填：否 - 主机内网IP</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 云镜客户端唯一UUID。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeMaliciousRequestsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMaliciousRequestsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Uuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMaliciousRequestsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMaliciousRequestsResponseParams struct {
	// 记录总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 恶意请求记录数组。
	MaliciousRequests []*MaliciousRequest `json:"MaliciousRequests,omitempty" name:"MaliciousRequests"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeMaliciousRequestsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMaliciousRequestsResponseParams `json:"Response"`
}

func (r *DescribeMaliciousRequestsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMaliciousRequestsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMalwaresRequestParams struct {
	// 客户端唯一Uuid。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 查询关键字 </li>
	// <li>Status - String - 是否必填：否 - 木马状态（UN_OPERATED: 未处理 | SEGREGATED: 已隔离|TRUSTED：信任）</li>
	// 每个过滤条件只支持一个值，暂不支持多个值“或”关系查询。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeMalwaresRequest struct {
	*tchttp.BaseRequest
	
	// 客户端唯一Uuid。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 查询关键字 </li>
	// <li>Status - String - 是否必填：否 - 木马状态（UN_OPERATED: 未处理 | SEGREGATED: 已隔离|TRUSTED：信任）</li>
	// 每个过滤条件只支持一个值，暂不支持多个值“或”关系查询。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeMalwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMalwaresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMalwaresRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMalwaresResponseParams struct {
	// 木马总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Malware数组。
	Malwares []*Malware `json:"Malwares,omitempty" name:"Malwares"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeMalwaresResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMalwaresResponseParams `json:"Response"`
}

func (r *DescribeMalwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMalwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNonlocalLoginPlacesRequestParams struct {
	// 客户端唯一Uuid。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 -  查询关键字</li>
	// <li>Status - String - 是否必填：否 -  登录状态（NON_LOCAL_LOGIN: 异地登录 | NORMAL_LOGIN : 正常登录）</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeNonlocalLoginPlacesRequest struct {
	*tchttp.BaseRequest
	
	// 客户端唯一Uuid。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 -  查询关键字</li>
	// <li>Status - String - 是否必填：否 -  登录状态（NON_LOCAL_LOGIN: 异地登录 | NORMAL_LOGIN : 正常登录）</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeNonlocalLoginPlacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNonlocalLoginPlacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNonlocalLoginPlacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNonlocalLoginPlacesResponseParams struct {
	// 记录总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 异地登录信息数组。
	NonLocalLoginPlaces []*NonLocalLoginPlace `json:"NonLocalLoginPlaces,omitempty" name:"NonLocalLoginPlaces"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeNonlocalLoginPlacesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNonlocalLoginPlacesResponseParams `json:"Response"`
}

func (r *DescribeNonlocalLoginPlacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNonlocalLoginPlacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOpenPortStatisticsRequestParams struct {
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Port - Uint64 - 是否必填：否 - 端口号</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeOpenPortStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Port - Uint64 - 是否必填：否 - 端口号</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeOpenPortStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOpenPortStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOpenPortStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOpenPortStatisticsResponseParams struct {
	// 端口统计列表总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 端口统计数据列表
	OpenPortStatistics []*OpenPortStatistics `json:"OpenPortStatistics,omitempty" name:"OpenPortStatistics"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeOpenPortStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOpenPortStatisticsResponseParams `json:"Response"`
}

func (r *DescribeOpenPortStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOpenPortStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOpenPortTaskStatusRequestParams struct {
	// 云镜客户端唯一Uuid。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeOpenPortTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// 云镜客户端唯一Uuid。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeOpenPortTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOpenPortTaskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOpenPortTaskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOpenPortTaskStatusResponseParams struct {
	// 任务状态。
	// <li>COMPLETE：完成（此时可以调用DescribeOpenPorts接口获取实时进程列表）</li>
	// <li>AGENT_OFFLINE：云镜客户端离线</li>
	// <li>COLLECTING：端口获取中</li>
	// <li>FAILED：进程获取失败</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeOpenPortTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOpenPortTaskStatusResponseParams `json:"Response"`
}

func (r *DescribeOpenPortTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOpenPortTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOpenPortsRequestParams struct {
	// 云镜客户端唯一Uuid。Port和Uuid必填其一，使用Uuid表示，查询该主机列表信息。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 开放端口号。Port和Uuid必填其一，使用Port表示查询该端口的列表信息。
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Port - Uint64 - 是否必填：否 - 端口号</li>
	// <li>ProcessName - String - 是否必填：否 - 进程名</li>
	// <li>MachineIp - String - 是否必填：否 - 主机内网IP</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeOpenPortsRequest struct {
	*tchttp.BaseRequest
	
	// 云镜客户端唯一Uuid。Port和Uuid必填其一，使用Uuid表示，查询该主机列表信息。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 开放端口号。Port和Uuid必填其一，使用Port表示查询该端口的列表信息。
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Port - Uint64 - 是否必填：否 - 端口号</li>
	// <li>ProcessName - String - 是否必填：否 - 进程名</li>
	// <li>MachineIp - String - 是否必填：否 - 主机内网IP</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeOpenPortsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOpenPortsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	delete(f, "Port")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOpenPortsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOpenPortsResponseParams struct {
	// 端口列表记录总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 端口列表。
	OpenPorts []*OpenPort `json:"OpenPorts,omitempty" name:"OpenPorts"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeOpenPortsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOpenPortsResponseParams `json:"Response"`
}

func (r *DescribeOpenPortsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOpenPortsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOverviewStatisticsRequestParams struct {

}

type DescribeOverviewStatisticsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeOverviewStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOverviewStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOverviewStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOverviewStatisticsResponseParams struct {
	// 服务器在线数。
	OnlineMachineNum *uint64 `json:"OnlineMachineNum,omitempty" name:"OnlineMachineNum"`

	// 专业服务器数。
	ProVersionMachineNum *uint64 `json:"ProVersionMachineNum,omitempty" name:"ProVersionMachineNum"`

	// 木马文件数。
	MalwareNum *uint64 `json:"MalwareNum,omitempty" name:"MalwareNum"`

	// 异地登录数。
	NonlocalLoginNum *uint64 `json:"NonlocalLoginNum,omitempty" name:"NonlocalLoginNum"`

	// 暴力破解成功数。
	BruteAttackSuccessNum *uint64 `json:"BruteAttackSuccessNum,omitempty" name:"BruteAttackSuccessNum"`

	// 漏洞数。
	VulNum *uint64 `json:"VulNum,omitempty" name:"VulNum"`

	// 安全基线数。
	BaseLineNum *uint64 `json:"BaseLineNum,omitempty" name:"BaseLineNum"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeOverviewStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOverviewStatisticsResponseParams `json:"Response"`
}

func (r *DescribeOverviewStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOverviewStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrivilegeEventsRequestParams struct {
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 关键词(主机IP)</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribePrivilegeEventsRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 关键词(主机IP)</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribePrivilegeEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivilegeEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrivilegeEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrivilegeEventsResponseParams struct {
	// 数据列表
	List []*PrivilegeEscalationProcess `json:"List,omitempty" name:"List"`

	// 总条数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePrivilegeEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrivilegeEventsResponseParams `json:"Response"`
}

func (r *DescribePrivilegeEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivilegeEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrivilegeRulesRequestParams struct {
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 关键字(进程名称)</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribePrivilegeRulesRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 关键字(进程名称)</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribePrivilegeRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivilegeRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrivilegeRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrivilegeRulesResponseParams struct {
	// 列表内容
	List []*PrivilegeRule `json:"List,omitempty" name:"List"`

	// 总条数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePrivilegeRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrivilegeRulesResponseParams `json:"Response"`
}

func (r *DescribePrivilegeRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivilegeRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProVersionInfoRequestParams struct {

}

type DescribeProVersionInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeProVersionInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProVersionInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProVersionInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProVersionInfoResponseParams struct {
	// 后付费昨日扣费
	PostPayCost *uint64 `json:"PostPayCost,omitempty" name:"PostPayCost"`

	// 新增主机是否自动开通专业版
	IsAutoOpenProVersion *bool `json:"IsAutoOpenProVersion,omitempty" name:"IsAutoOpenProVersion"`

	// 开通专业版主机数
	ProVersionNum *uint64 `json:"ProVersionNum,omitempty" name:"ProVersionNum"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeProVersionInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProVersionInfoResponseParams `json:"Response"`
}

func (r *DescribeProVersionInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProVersionInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProcessStatisticsRequestParams struct {
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>ProcessName - String - 是否必填：否 - 进程名</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeProcessStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>ProcessName - String - 是否必填：否 - 进程名</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeProcessStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProcessStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProcessStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProcessStatisticsResponseParams struct {
	// 进程统计列表记录总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 进程统计列表数据数组。
	ProcessStatistics []*ProcessStatistics `json:"ProcessStatistics,omitempty" name:"ProcessStatistics"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeProcessStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProcessStatisticsResponseParams `json:"Response"`
}

func (r *DescribeProcessStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProcessStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProcessTaskStatusRequestParams struct {
	// 云镜客户端唯一Uuid。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeProcessTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// 云镜客户端唯一Uuid。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeProcessTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProcessTaskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProcessTaskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProcessTaskStatusResponseParams struct {
	// 任务状态。
	// <li>COMPLETE：完成（此时可以调用DescribeProcesses接口获取实时进程列表）</li>
	// <li>AGENT_OFFLINE：云镜客户端离线</li>
	// <li>COLLECTING：进程获取中</li>
	// <li>FAILED：进程获取失败</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeProcessTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProcessTaskStatusResponseParams `json:"Response"`
}

func (r *DescribeProcessTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProcessTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProcessesRequestParams struct {
	// 云镜客户端唯一Uuid。Uuid和ProcessName必填其一，使用Uuid表示，查询该主机列表信息。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 进程名。Uuid和ProcessName必填其一，使用ProcessName表示，查询该进程列表信息。
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>ProcessName - String - 是否必填：否 - 进程名</li>
	// <li>MachineIp - String - 是否必填：否 - 主机内网IP</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeProcessesRequest struct {
	*tchttp.BaseRequest
	
	// 云镜客户端唯一Uuid。Uuid和ProcessName必填其一，使用Uuid表示，查询该主机列表信息。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 进程名。Uuid和ProcessName必填其一，使用ProcessName表示，查询该进程列表信息。
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>ProcessName - String - 是否必填：否 - 进程名</li>
	// <li>MachineIp - String - 是否必填：否 - 主机内网IP</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeProcessesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProcessesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	delete(f, "ProcessName")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProcessesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProcessesResponseParams struct {
	// 进程列表记录总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 进程列表数据数组。
	Processes []*Process `json:"Processes,omitempty" name:"Processes"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeProcessesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProcessesResponseParams `json:"Response"`
}

func (r *DescribeProcessesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProcessesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReverseShellEventsRequestParams struct {
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 关键字(主机内网IP|进程名)</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeReverseShellEventsRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 关键字(主机内网IP|进程名)</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeReverseShellEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReverseShellEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReverseShellEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReverseShellEventsResponseParams struct {
	// 列表内容
	List []*ReverseShell `json:"List,omitempty" name:"List"`

	// 总条数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeReverseShellEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReverseShellEventsResponseParams `json:"Response"`
}

func (r *DescribeReverseShellEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReverseShellEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReverseShellRulesRequestParams struct {
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 关键字(进程名称)</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeReverseShellRulesRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 关键字(进程名称)</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeReverseShellRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReverseShellRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReverseShellRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReverseShellRulesResponseParams struct {
	// 列表内容
	List []*ReverseShellRule `json:"List,omitempty" name:"List"`

	// 总条数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeReverseShellRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReverseShellRulesResponseParams `json:"Response"`
}

func (r *DescribeReverseShellRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReverseShellRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityDynamicsRequestParams struct {
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeSecurityDynamicsRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeSecurityDynamicsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityDynamicsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityDynamicsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityDynamicsResponseParams struct {
	// 安全事件消息数组。
	SecurityDynamics []*SecurityDynamic `json:"SecurityDynamics,omitempty" name:"SecurityDynamics"`

	// 记录总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSecurityDynamicsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityDynamicsResponseParams `json:"Response"`
}

func (r *DescribeSecurityDynamicsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityDynamicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityTrendsRequestParams struct {
	// 开始时间。
	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`

	// 结束时间。
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

type DescribeSecurityTrendsRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`

	// 结束时间。
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

func (r *DescribeSecurityTrendsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityTrendsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginDate")
	delete(f, "EndDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityTrendsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityTrendsResponseParams struct {
	// 木马事件统计数据数组。
	Malwares []*SecurityTrend `json:"Malwares,omitempty" name:"Malwares"`

	// 异地登录事件统计数据数组。
	NonLocalLoginPlaces []*SecurityTrend `json:"NonLocalLoginPlaces,omitempty" name:"NonLocalLoginPlaces"`

	// 密码破解事件统计数据数组。
	BruteAttacks []*SecurityTrend `json:"BruteAttacks,omitempty" name:"BruteAttacks"`

	// 漏洞统计数据数组。
	Vuls []*SecurityTrend `json:"Vuls,omitempty" name:"Vuls"`

	// 基线统计数据数组。
	BaseLines []*SecurityTrend `json:"BaseLines,omitempty" name:"BaseLines"`

	// 恶意请求统计数据数组。
	MaliciousRequests []*SecurityTrend `json:"MaliciousRequests,omitempty" name:"MaliciousRequests"`

	// 高危命令统计数据数组。
	HighRiskBashs []*SecurityTrend `json:"HighRiskBashs,omitempty" name:"HighRiskBashs"`

	// 反弹shell统计数据数组。
	ReverseShells []*SecurityTrend `json:"ReverseShells,omitempty" name:"ReverseShells"`

	// 本地提权统计数据数组。
	PrivilegeEscalations []*SecurityTrend `json:"PrivilegeEscalations,omitempty" name:"PrivilegeEscalations"`

	// 网络攻击统计数据数组。
	CyberAttacks []*SecurityTrend `json:"CyberAttacks,omitempty" name:"CyberAttacks"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSecurityTrendsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityTrendsResponseParams `json:"Response"`
}

func (r *DescribeSecurityTrendsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityTrendsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagMachinesRequestParams struct {
	// 标签ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

type DescribeTagMachinesRequest struct {
	*tchttp.BaseRequest
	
	// 标签ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeTagMachinesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagMachinesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTagMachinesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagMachinesResponseParams struct {
	// 列表数据
	List []*TagMachine `json:"List,omitempty" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTagMachinesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTagMachinesResponseParams `json:"Response"`
}

func (r *DescribeTagMachinesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagMachinesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagsRequestParams struct {
	// 云主机类型。
	// <li>CVM：表示虚拟主机</li>
	// <li>BM:  表示黑石物理机</li>
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// 机器所属地域。如：ap-guangzhou，ap-shanghai
	MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`
}

type DescribeTagsRequest struct {
	*tchttp.BaseRequest
	
	// 云主机类型。
	// <li>CVM：表示虚拟主机</li>
	// <li>BM:  表示黑石物理机</li>
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// 机器所属地域。如：ap-guangzhou，ap-shanghai
	MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`
}

func (r *DescribeTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MachineType")
	delete(f, "MachineRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagsResponseParams struct {
	// 列表信息
	List []*Tag `json:"List,omitempty" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTagsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTagsResponseParams `json:"Response"`
}

func (r *DescribeTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsualLoginPlacesRequestParams struct {
	// 云镜客户端UUID
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeUsualLoginPlacesRequest struct {
	*tchttp.BaseRequest
	
	// 云镜客户端UUID
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeUsualLoginPlacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsualLoginPlacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUsualLoginPlacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsualLoginPlacesResponseParams struct {
	// 常用登录地数组
	UsualLoginPlaces []*UsualPlace `json:"UsualLoginPlaces,omitempty" name:"UsualLoginPlaces"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUsualLoginPlacesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUsualLoginPlacesResponseParams `json:"Response"`
}

func (r *DescribeUsualLoginPlacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsualLoginPlacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulInfoRequestParams struct {
	// 漏洞种类ID。
	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`
}

type DescribeVulInfoRequest struct {
	*tchttp.BaseRequest
	
	// 漏洞种类ID。
	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`
}

func (r *DescribeVulInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VulId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulInfoResponseParams struct {
	// 漏洞种类ID。
	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`

	// 漏洞名称。
	VulName *string `json:"VulName,omitempty" name:"VulName"`

	// 漏洞等级。
	VulLevel *string `json:"VulLevel,omitempty" name:"VulLevel"`

	// 漏洞类型。
	VulType *string `json:"VulType,omitempty" name:"VulType"`

	// 漏洞描述。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 修复方案。
	RepairPlan *string `json:"RepairPlan,omitempty" name:"RepairPlan"`

	// 漏洞CVE。
	CveId *string `json:"CveId,omitempty" name:"CveId"`

	// 参考链接。
	Reference *string `json:"Reference,omitempty" name:"Reference"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVulInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVulInfoResponseParams `json:"Response"`
}

func (r *DescribeVulInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulScanResultRequestParams struct {

}

type DescribeVulScanResultRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeVulScanResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulScanResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulScanResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulScanResultResponseParams struct {
	// 漏洞数量。
	VulNum *uint64 `json:"VulNum,omitempty" name:"VulNum"`

	// 专业版机器数。
	ProVersionNum *uint64 `json:"ProVersionNum,omitempty" name:"ProVersionNum"`

	// 受影响的专业版主机数。
	ImpactedHostNum *uint64 `json:"ImpactedHostNum,omitempty" name:"ImpactedHostNum"`

	// 主机总数。
	HostNum *uint64 `json:"HostNum,omitempty" name:"HostNum"`

	// 基础版机器数。
	BasicVersionNum *uint64 `json:"BasicVersionNum,omitempty" name:"BasicVersionNum"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVulScanResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVulScanResultResponseParams `json:"Response"`
}

func (r *DescribeVulScanResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulScanResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulsRequestParams struct {
	// 漏洞类型。
	// <li>WEB：Web应用漏洞</li>
	// <li>SYSTEM：系统组件漏洞</li>
	// <li>BASELINE：安全基线</li>
	VulType *string `json:"VulType,omitempty" name:"VulType"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Status - String - 是否必填：否 - 状态筛选（UN_OPERATED: 待处理 | FIXED：已修复）
	// 
	// Status过滤条件值只能取其一，不能是“或”逻辑。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeVulsRequest struct {
	*tchttp.BaseRequest
	
	// 漏洞类型。
	// <li>WEB：Web应用漏洞</li>
	// <li>SYSTEM：系统组件漏洞</li>
	// <li>BASELINE：安全基线</li>
	VulType *string `json:"VulType,omitempty" name:"VulType"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Status - String - 是否必填：否 - 状态筛选（UN_OPERATED: 待处理 | FIXED：已修复）
	// 
	// Status过滤条件值只能取其一，不能是“或”逻辑。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeVulsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VulType")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulsResponseParams struct {
	// 漏洞数量。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 漏洞列表数组。
	Vuls []*Vul `json:"Vuls,omitempty" name:"Vuls"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVulsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVulsResponseParams `json:"Response"`
}

func (r *DescribeVulsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWeeklyReportBruteAttacksRequestParams struct {
	// 专业周报开始时间。
	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeWeeklyReportBruteAttacksRequest struct {
	*tchttp.BaseRequest
	
	// 专业周报开始时间。
	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeWeeklyReportBruteAttacksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWeeklyReportBruteAttacksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginDate")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWeeklyReportBruteAttacksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWeeklyReportBruteAttacksResponseParams struct {
	// 专业周报密码破解数组。
	WeeklyReportBruteAttacks []*WeeklyReportBruteAttack `json:"WeeklyReportBruteAttacks,omitempty" name:"WeeklyReportBruteAttacks"`

	// 记录总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWeeklyReportBruteAttacksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWeeklyReportBruteAttacksResponseParams `json:"Response"`
}

func (r *DescribeWeeklyReportBruteAttacksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWeeklyReportBruteAttacksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWeeklyReportInfoRequestParams struct {
	// 专业周报开始时间。
	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`
}

type DescribeWeeklyReportInfoRequest struct {
	*tchttp.BaseRequest
	
	// 专业周报开始时间。
	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`
}

func (r *DescribeWeeklyReportInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWeeklyReportInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWeeklyReportInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWeeklyReportInfoResponseParams struct {
	// 账号所属公司或个人名称。
	CompanyName *string `json:"CompanyName,omitempty" name:"CompanyName"`

	// 机器总数。
	MachineNum *uint64 `json:"MachineNum,omitempty" name:"MachineNum"`

	// 云镜客户端在线数。
	OnlineMachineNum *uint64 `json:"OnlineMachineNum,omitempty" name:"OnlineMachineNum"`

	// 云镜客户端离线数。
	OfflineMachineNum *uint64 `json:"OfflineMachineNum,omitempty" name:"OfflineMachineNum"`

	// 开通云镜专业版数量。
	ProVersionMachineNum *uint64 `json:"ProVersionMachineNum,omitempty" name:"ProVersionMachineNum"`

	// 周报开始时间。
	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`

	// 周报结束时间。
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 安全等级。
	// <li>HIGH：高</li>
	// <li>MIDDLE：中</li>
	// <li>LOW：低</li>
	Level *string `json:"Level,omitempty" name:"Level"`

	// 木马记录数。
	MalwareNum *uint64 `json:"MalwareNum,omitempty" name:"MalwareNum"`

	// 异地登录数。
	NonlocalLoginNum *uint64 `json:"NonlocalLoginNum,omitempty" name:"NonlocalLoginNum"`

	// 密码破解成功数。
	BruteAttackSuccessNum *uint64 `json:"BruteAttackSuccessNum,omitempty" name:"BruteAttackSuccessNum"`

	// 漏洞数。
	VulNum *uint64 `json:"VulNum,omitempty" name:"VulNum"`

	// 导出文件下载地址。
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWeeklyReportInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWeeklyReportInfoResponseParams `json:"Response"`
}

func (r *DescribeWeeklyReportInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWeeklyReportInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWeeklyReportMalwaresRequestParams struct {
	// 专业周报开始时间。
	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeWeeklyReportMalwaresRequest struct {
	*tchttp.BaseRequest
	
	// 专业周报开始时间。
	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeWeeklyReportMalwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWeeklyReportMalwaresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginDate")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWeeklyReportMalwaresRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWeeklyReportMalwaresResponseParams struct {
	// 专业周报木马数据。
	WeeklyReportMalwares []*WeeklyReportMalware `json:"WeeklyReportMalwares,omitempty" name:"WeeklyReportMalwares"`

	// 记录总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWeeklyReportMalwaresResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWeeklyReportMalwaresResponseParams `json:"Response"`
}

func (r *DescribeWeeklyReportMalwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWeeklyReportMalwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWeeklyReportNonlocalLoginPlacesRequestParams struct {
	// 专业周报开始时间。
	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeWeeklyReportNonlocalLoginPlacesRequest struct {
	*tchttp.BaseRequest
	
	// 专业周报开始时间。
	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeWeeklyReportNonlocalLoginPlacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWeeklyReportNonlocalLoginPlacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginDate")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWeeklyReportNonlocalLoginPlacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWeeklyReportNonlocalLoginPlacesResponseParams struct {
	// 专业周报异地登录数据。
	WeeklyReportNonlocalLoginPlaces []*WeeklyReportNonlocalLoginPlace `json:"WeeklyReportNonlocalLoginPlaces,omitempty" name:"WeeklyReportNonlocalLoginPlaces"`

	// 记录总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWeeklyReportNonlocalLoginPlacesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWeeklyReportNonlocalLoginPlacesResponseParams `json:"Response"`
}

func (r *DescribeWeeklyReportNonlocalLoginPlacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWeeklyReportNonlocalLoginPlacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWeeklyReportVulsRequestParams struct {
	// 专业版周报开始时间。
	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeWeeklyReportVulsRequest struct {
	*tchttp.BaseRequest
	
	// 专业版周报开始时间。
	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeWeeklyReportVulsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWeeklyReportVulsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginDate")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWeeklyReportVulsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWeeklyReportVulsResponseParams struct {
	// 专业周报漏洞数据数组。
	WeeklyReportVuls []*WeeklyReportVul `json:"WeeklyReportVuls,omitempty" name:"WeeklyReportVuls"`

	// 记录总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWeeklyReportVulsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWeeklyReportVulsResponseParams `json:"Response"`
}

func (r *DescribeWeeklyReportVulsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWeeklyReportVulsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWeeklyReportsRequestParams struct {
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeWeeklyReportsRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeWeeklyReportsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWeeklyReportsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWeeklyReportsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWeeklyReportsResponseParams struct {
	// 专业周报列表数组。
	WeeklyReports []*WeeklyReport `json:"WeeklyReports,omitempty" name:"WeeklyReports"`

	// 记录总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWeeklyReportsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWeeklyReportsResponseParams `json:"Response"`
}

func (r *DescribeWeeklyReportsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWeeklyReportsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EditBashRuleRequestParams struct {
	// 规则名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 危险等级(1: 高危 2:中危 3: 低危)
	Level *uint64 `json:"Level,omitempty" name:"Level"`

	// 正则表达式
	Rule *string `json:"Rule,omitempty" name:"Rule"`

	// 规则ID(新增时不填)
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 客户端ID(IsGlobal为1时，Uuid或Hostip必填一个)
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 主机IP(IsGlobal为1时，Uuid或Hostip必填一个)
	Hostip *string `json:"Hostip,omitempty" name:"Hostip"`

	// 是否全局规则(默认否)
	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`
}

type EditBashRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 危险等级(1: 高危 2:中危 3: 低危)
	Level *uint64 `json:"Level,omitempty" name:"Level"`

	// 正则表达式
	Rule *string `json:"Rule,omitempty" name:"Rule"`

	// 规则ID(新增时不填)
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 客户端ID(IsGlobal为1时，Uuid或Hostip必填一个)
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 主机IP(IsGlobal为1时，Uuid或Hostip必填一个)
	Hostip *string `json:"Hostip,omitempty" name:"Hostip"`

	// 是否全局规则(默认否)
	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`
}

func (r *EditBashRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EditBashRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Level")
	delete(f, "Rule")
	delete(f, "Id")
	delete(f, "Uuid")
	delete(f, "Hostip")
	delete(f, "IsGlobal")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EditBashRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EditBashRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type EditBashRuleResponse struct {
	*tchttp.BaseResponse
	Response *EditBashRuleResponseParams `json:"Response"`
}

func (r *EditBashRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EditBashRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EditPrivilegeRuleRequestParams struct {
	// 规则ID(新增时请留空)
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 客户端ID(IsGlobal为1时，Uuid或Hostip必填一个)
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 主机IP(IsGlobal为1时，Uuid或Hostip必填一个)
	Hostip *string `json:"Hostip,omitempty" name:"Hostip"`

	// 进程名
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// 是否S权限进程
	SMode *uint64 `json:"SMode,omitempty" name:"SMode"`

	// 是否全局规则(默认否)
	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`
}

type EditPrivilegeRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则ID(新增时请留空)
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 客户端ID(IsGlobal为1时，Uuid或Hostip必填一个)
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 主机IP(IsGlobal为1时，Uuid或Hostip必填一个)
	Hostip *string `json:"Hostip,omitempty" name:"Hostip"`

	// 进程名
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// 是否S权限进程
	SMode *uint64 `json:"SMode,omitempty" name:"SMode"`

	// 是否全局规则(默认否)
	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`
}

func (r *EditPrivilegeRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EditPrivilegeRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Uuid")
	delete(f, "Hostip")
	delete(f, "ProcessName")
	delete(f, "SMode")
	delete(f, "IsGlobal")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EditPrivilegeRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EditPrivilegeRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type EditPrivilegeRuleResponse struct {
	*tchttp.BaseResponse
	Response *EditPrivilegeRuleResponseParams `json:"Response"`
}

func (r *EditPrivilegeRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EditPrivilegeRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EditReverseShellRuleRequestParams struct {
	// 规则ID(新增时请留空)
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 客户端ID(IsGlobal为1时，Uuid或Hostip必填一个)
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 主机IP(IsGlobal为1时，Uuid或Hostip必填一个)
	Hostip *string `json:"Hostip,omitempty" name:"Hostip"`

	// 目标IP
	DestIp *string `json:"DestIp,omitempty" name:"DestIp"`

	// 目标端口
	DestPort *string `json:"DestPort,omitempty" name:"DestPort"`

	// 进程名
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// 是否全局规则(默认否)
	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`
}

type EditReverseShellRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则ID(新增时请留空)
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 客户端ID(IsGlobal为1时，Uuid或Hostip必填一个)
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 主机IP(IsGlobal为1时，Uuid或Hostip必填一个)
	Hostip *string `json:"Hostip,omitempty" name:"Hostip"`

	// 目标IP
	DestIp *string `json:"DestIp,omitempty" name:"DestIp"`

	// 目标端口
	DestPort *string `json:"DestPort,omitempty" name:"DestPort"`

	// 进程名
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// 是否全局规则(默认否)
	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`
}

func (r *EditReverseShellRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EditReverseShellRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Uuid")
	delete(f, "Hostip")
	delete(f, "DestIp")
	delete(f, "DestPort")
	delete(f, "ProcessName")
	delete(f, "IsGlobal")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EditReverseShellRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EditReverseShellRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type EditReverseShellRuleResponse struct {
	*tchttp.BaseResponse
	Response *EditReverseShellRuleResponseParams `json:"Response"`
}

func (r *EditReverseShellRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EditReverseShellRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EditTagsRequestParams struct {
	// 标签名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 标签ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// CVM主机ID
	Quuids []*string `json:"Quuids,omitempty" name:"Quuids"`
}

type EditTagsRequest struct {
	*tchttp.BaseRequest
	
	// 标签名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 标签ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// CVM主机ID
	Quuids []*string `json:"Quuids,omitempty" name:"Quuids"`
}

func (r *EditTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EditTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Id")
	delete(f, "Quuids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EditTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EditTagsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type EditTagsResponse struct {
	*tchttp.BaseResponse
	Response *EditTagsResponseParams `json:"Response"`
}

func (r *EditTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EditTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportAttackLogsRequestParams struct {

}

type ExportAttackLogsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ExportAttackLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportAttackLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportAttackLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportAttackLogsResponseParams struct {
	// 导出文件下载链接地址。
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 导出任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ExportAttackLogsResponse struct {
	*tchttp.BaseResponse
	Response *ExportAttackLogsResponseParams `json:"Response"`
}

func (r *ExportAttackLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportAttackLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportBashEventsRequestParams struct {

}

type ExportBashEventsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ExportBashEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportBashEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportBashEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportBashEventsResponseParams struct {
	// 导出文件下载链接地址。
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ExportBashEventsResponse struct {
	*tchttp.BaseResponse
	Response *ExportBashEventsResponseParams `json:"Response"`
}

func (r *ExportBashEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportBashEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportBruteAttacksRequestParams struct {

}

type ExportBruteAttacksRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ExportBruteAttacksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportBruteAttacksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportBruteAttacksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportBruteAttacksResponseParams struct {
	// 导出文件下载链接地址。
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ExportBruteAttacksResponse struct {
	*tchttp.BaseResponse
	Response *ExportBruteAttacksResponseParams `json:"Response"`
}

func (r *ExportBruteAttacksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportBruteAttacksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportMaliciousRequestsRequestParams struct {

}

type ExportMaliciousRequestsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ExportMaliciousRequestsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportMaliciousRequestsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportMaliciousRequestsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportMaliciousRequestsResponseParams struct {
	// 导出文件下载链接地址。
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ExportMaliciousRequestsResponse struct {
	*tchttp.BaseResponse
	Response *ExportMaliciousRequestsResponseParams `json:"Response"`
}

func (r *ExportMaliciousRequestsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportMaliciousRequestsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportMalwaresRequestParams struct {

}

type ExportMalwaresRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ExportMalwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportMalwaresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportMalwaresRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportMalwaresResponseParams struct {
	// 导出文件下载链接地址。
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ExportMalwaresResponse struct {
	*tchttp.BaseResponse
	Response *ExportMalwaresResponseParams `json:"Response"`
}

func (r *ExportMalwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportMalwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportNonlocalLoginPlacesRequestParams struct {

}

type ExportNonlocalLoginPlacesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ExportNonlocalLoginPlacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportNonlocalLoginPlacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportNonlocalLoginPlacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportNonlocalLoginPlacesResponseParams struct {
	// 导出文件下载链接地址。
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 导出任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ExportNonlocalLoginPlacesResponse struct {
	*tchttp.BaseResponse
	Response *ExportNonlocalLoginPlacesResponseParams `json:"Response"`
}

func (r *ExportNonlocalLoginPlacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportNonlocalLoginPlacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportPrivilegeEventsRequestParams struct {

}

type ExportPrivilegeEventsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ExportPrivilegeEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportPrivilegeEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportPrivilegeEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportPrivilegeEventsResponseParams struct {
	// 导出文件下载链接地址。
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ExportPrivilegeEventsResponse struct {
	*tchttp.BaseResponse
	Response *ExportPrivilegeEventsResponseParams `json:"Response"`
}

func (r *ExportPrivilegeEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportPrivilegeEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportReverseShellEventsRequestParams struct {

}

type ExportReverseShellEventsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ExportReverseShellEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportReverseShellEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportReverseShellEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportReverseShellEventsResponseParams struct {
	// 导出文件下载链接地址。
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ExportReverseShellEventsResponse struct {
	*tchttp.BaseResponse
	Response *ExportReverseShellEventsResponseParams `json:"Response"`
}

func (r *ExportReverseShellEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportReverseShellEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 过滤键的名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 一个或者多个过滤值。
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type HistoryAccount struct {
	// 唯一ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 云镜客户端唯一Uuid。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 主机内网IP。
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// 主机名。
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// 帐号名。
	Username *string `json:"Username,omitempty" name:"Username"`

	// 帐号变更类型。
	// <li>CREATE：表示新增帐号</li>
	// <li>MODIFY：表示修改帐号</li>
	// <li>DELETE：表示删除帐号</li>
	ModifyType *string `json:"ModifyType,omitempty" name:"ModifyType"`

	// 变更时间。
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

// Predefined struct for user
type IgnoreImpactedHostsRequestParams struct {
	// 漏洞ID数组。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

type IgnoreImpactedHostsRequest struct {
	*tchttp.BaseRequest
	
	// 漏洞ID数组。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *IgnoreImpactedHostsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IgnoreImpactedHostsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IgnoreImpactedHostsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IgnoreImpactedHostsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type IgnoreImpactedHostsResponse struct {
	*tchttp.BaseResponse
	Response *IgnoreImpactedHostsResponseParams `json:"Response"`
}

func (r *IgnoreImpactedHostsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IgnoreImpactedHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImpactedHost struct {
	// 漏洞ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 主机IP。
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// 主机名称。
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// 最后检测时间。
	LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`

	// 漏洞状态。
	// <li>UN_OPERATED ：待处理</li>
	// <li>SCANING : 扫描中</li>
	// <li>FIXED : 已修复</li>
	VulStatus *string `json:"VulStatus,omitempty" name:"VulStatus"`

	// 云镜客户端唯一标识UUID。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 漏洞描述。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 漏洞种类ID。
	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`

	// 是否为专业版。
	IsProVersion *bool `json:"IsProVersion,omitempty" name:"IsProVersion"`
}

// Predefined struct for user
type InquiryPriceOpenProVersionPrepaidRequestParams struct {
	// 预付费模式(包年包月)参数设置。
	ChargePrepaid *ChargePrepaid `json:"ChargePrepaid,omitempty" name:"ChargePrepaid"`

	// 需要开通专业版机器列表数组。
	Machines []*ProVersionMachine `json:"Machines,omitempty" name:"Machines"`
}

type InquiryPriceOpenProVersionPrepaidRequest struct {
	*tchttp.BaseRequest
	
	// 预付费模式(包年包月)参数设置。
	ChargePrepaid *ChargePrepaid `json:"ChargePrepaid,omitempty" name:"ChargePrepaid"`

	// 需要开通专业版机器列表数组。
	Machines []*ProVersionMachine `json:"Machines,omitempty" name:"Machines"`
}

func (r *InquiryPriceOpenProVersionPrepaidRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceOpenProVersionPrepaidRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChargePrepaid")
	delete(f, "Machines")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceOpenProVersionPrepaidRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceOpenProVersionPrepaidResponseParams struct {
	// 预支费用的原价，单位：元。
	OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

	// 预支费用的折扣价，单位：元。
	DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type InquiryPriceOpenProVersionPrepaidResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceOpenProVersionPrepaidResponseParams `json:"Response"`
}

func (r *InquiryPriceOpenProVersionPrepaidResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceOpenProVersionPrepaidResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoginWhiteLists struct {
	// 记录ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 云镜客户端ID
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 白名单地域
	Places []*Place `json:"Places,omitempty" name:"Places"`

	// 白名单用户（多个用户逗号隔开）
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 白名单IP（多个IP逗号隔开）
	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`

	// 是否为全局规则
	IsGlobal *bool `json:"IsGlobal,omitempty" name:"IsGlobal"`

	// 创建白名单时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 修改白名单时间
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// 机器名
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// 机器IP
	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`

	// 起始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type LoginWhiteListsRule struct {
	// 加白地域
	Places []*Place `json:"Places,omitempty" name:"Places"`

	// 加白源IP，支持网段，多个IP以逗号隔开
	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`

	// 加白用户名，多个用户名以逗号隔开
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 是否对全局生效
	IsGlobal *bool `json:"IsGlobal,omitempty" name:"IsGlobal"`

	// 白名单生效的机器
	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`

	// 规则ID，用于更新规则
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 起始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type Machine struct {
	// 主机名称。
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// 主机系统。
	MachineOs *string `json:"MachineOs,omitempty" name:"MachineOs"`

	// 主机状态。
	// <li>OFFLINE: 离线  </li>
	// <li>ONLINE: 在线</li>
	// <li>MACHINE_STOPPED: 已关机</li>
	MachineStatus *string `json:"MachineStatus,omitempty" name:"MachineStatus"`

	// 云镜客户端唯一Uuid，若客户端长时间不在线将返回空字符。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// CVM或BM机器唯一Uuid。
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`

	// 漏洞数。
	VulNum *int64 `json:"VulNum,omitempty" name:"VulNum"`

	// 主机IP。
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// 是否是专业版。
	// <li>true： 是</li>
	// <li>false：否</li>
	IsProVersion *bool `json:"IsProVersion,omitempty" name:"IsProVersion"`

	// 主机外网IP。
	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`

	// 主机状态。
	// <li>POSTPAY: 表示后付费，即按量计费  </li>
	// <li>PREPAY: 表示预付费，即包年包月</li>
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 木马数。
	MalwareNum *int64 `json:"MalwareNum,omitempty" name:"MalwareNum"`

	// 标签信息
	Tag []*MachineTag `json:"Tag,omitempty" name:"Tag"`

	// 基线风险数。
	BaselineNum *int64 `json:"BaselineNum,omitempty" name:"BaselineNum"`

	// 网络风险数。
	CyberAttackNum *int64 `json:"CyberAttackNum,omitempty" name:"CyberAttackNum"`

	// 风险状态。
	// <li>SAFE：安全</li>
	// <li>RISK：风险</li>
	// <li>UNKNOWN：未知</li>
	SecurityStatus *string `json:"SecurityStatus,omitempty" name:"SecurityStatus"`

	// 入侵事件数
	InvasionNum *int64 `json:"InvasionNum,omitempty" name:"InvasionNum"`

	// 地域信息
	RegionInfo *RegionInfo `json:"RegionInfo,omitempty" name:"RegionInfo"`
}

type MachineTag struct {
	// 关联标签ID
	Rid *int64 `json:"Rid,omitempty" name:"Rid"`

	// 标签名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 标签ID
	TagId *uint64 `json:"TagId,omitempty" name:"TagId"`
}

type MaliciousRequest struct {
	// 记录ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 云镜客户端UUID。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 主机内网IP。
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// 主机名。
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// 恶意请求域名。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 恶意请求数。
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 进程名。
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// 记录状态。
	// <li>UN_OPERATED：待处理</li>
	// <li>TRUSTED：已信任</li>
	// <li>UN_TRUSTED：已取消信任</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 恶意请求域名描述。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 参考地址。
	Reference *string `json:"Reference,omitempty" name:"Reference"`

	// 发现时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 记录合并时间。
	MergeTime *string `json:"MergeTime,omitempty" name:"MergeTime"`

	// 进程MD5
	// 值。
	ProcessMd5 *string `json:"ProcessMd5,omitempty" name:"ProcessMd5"`

	// 执行命令行。
	CmdLine *string `json:"CmdLine,omitempty" name:"CmdLine"`

	// 进程PID。
	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`
}

type Malware struct {
	// 事件ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 主机IP。
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// 当前木马状态。
	// <li>UN_OPERATED：未处理</li><li>SEGREGATED：已隔离</li><li>TRUSTED：已信任</li>
	// <li>SEPARATING：隔离中</li><li>RECOVERING：恢复中</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 木马所在的路径。
	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`

	// 木马描述。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 主机名称。
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// 木马文件创建时间。
	FileCreateTime *string `json:"FileCreateTime,omitempty" name:"FileCreateTime"`

	// 木马文件修改时间。
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// 云镜客户端唯一标识UUID。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

// Predefined struct for user
type MisAlarmNonlocalLoginPlacesRequestParams struct {
	// 异地登录事件Id数组。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

type MisAlarmNonlocalLoginPlacesRequest struct {
	*tchttp.BaseRequest
	
	// 异地登录事件Id数组。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *MisAlarmNonlocalLoginPlacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MisAlarmNonlocalLoginPlacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MisAlarmNonlocalLoginPlacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MisAlarmNonlocalLoginPlacesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type MisAlarmNonlocalLoginPlacesResponse struct {
	*tchttp.BaseResponse
	Response *MisAlarmNonlocalLoginPlacesResponseParams `json:"Response"`
}

func (r *MisAlarmNonlocalLoginPlacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MisAlarmNonlocalLoginPlacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAlarmAttributeRequestParams struct {
	// 告警项目。
	// <li>Offline：防护软件离线</li>
	// <li>Malware：发现木马文件</li>
	// <li>NonlocalLogin：发现异地登录行为</li>
	// <li>CrackSuccess：被暴力破解成功</li>
	Attribute *string `json:"Attribute,omitempty" name:"Attribute"`

	// 告警项目属性。
	// <li>CLOSE：关闭</li>
	// <li>OPEN：打开</li>
	Value *string `json:"Value,omitempty" name:"Value"`
}

type ModifyAlarmAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 告警项目。
	// <li>Offline：防护软件离线</li>
	// <li>Malware：发现木马文件</li>
	// <li>NonlocalLogin：发现异地登录行为</li>
	// <li>CrackSuccess：被暴力破解成功</li>
	Attribute *string `json:"Attribute,omitempty" name:"Attribute"`

	// 告警项目属性。
	// <li>CLOSE：关闭</li>
	// <li>OPEN：打开</li>
	Value *string `json:"Value,omitempty" name:"Value"`
}

func (r *ModifyAlarmAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Attribute")
	delete(f, "Value")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAlarmAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAlarmAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAlarmAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAlarmAttributeResponseParams `json:"Response"`
}

func (r *ModifyAlarmAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAutoOpenProVersionConfigRequestParams struct {
	// 设置自动开通状态。
	// <li>CLOSE：关闭</li>
	// <li>OPEN：打开</li>
	Status *string `json:"Status,omitempty" name:"Status"`
}

type ModifyAutoOpenProVersionConfigRequest struct {
	*tchttp.BaseRequest
	
	// 设置自动开通状态。
	// <li>CLOSE：关闭</li>
	// <li>OPEN：打开</li>
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyAutoOpenProVersionConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoOpenProVersionConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAutoOpenProVersionConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAutoOpenProVersionConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAutoOpenProVersionConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAutoOpenProVersionConfigResponseParams `json:"Response"`
}

func (r *ModifyAutoOpenProVersionConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoOpenProVersionConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoginWhiteListRequestParams struct {
	// 白名单规则
	Rules *LoginWhiteListsRule `json:"Rules,omitempty" name:"Rules"`
}

type ModifyLoginWhiteListRequest struct {
	*tchttp.BaseRequest
	
	// 白名单规则
	Rules *LoginWhiteListsRule `json:"Rules,omitempty" name:"Rules"`
}

func (r *ModifyLoginWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoginWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Rules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLoginWhiteListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoginWhiteListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyLoginWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLoginWhiteListResponseParams `json:"Response"`
}

func (r *ModifyLoginWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoginWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProVersionRenewFlagRequestParams struct {
	// 自动续费标识。取值范围：
	// <li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费</li>
	// <li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费</li>
	// <li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费</li>
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// 主机唯一ID，对应CVM的uuid、BM的instanceId。
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

type ModifyProVersionRenewFlagRequest struct {
	*tchttp.BaseRequest
	
	// 自动续费标识。取值范围：
	// <li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费</li>
	// <li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费</li>
	// <li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费</li>
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// 主机唯一ID，对应CVM的uuid、BM的instanceId。
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *ModifyProVersionRenewFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProVersionRenewFlagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RenewFlag")
	delete(f, "Quuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyProVersionRenewFlagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProVersionRenewFlagResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyProVersionRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *ModifyProVersionRenewFlagResponseParams `json:"Response"`
}

func (r *ModifyProVersionRenewFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProVersionRenewFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NonLocalLoginPlace struct {
	// 事件ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 主机IP。
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// 登录状态
	// <li>NON_LOCAL_LOGIN：异地登录</li>
	// <li>NORMAL_LOGIN：正常登录</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 用户名。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 城市ID。
	City *uint64 `json:"City,omitempty" name:"City"`

	// 国家ID。
	Country *uint64 `json:"Country,omitempty" name:"Country"`

	// 省份ID。
	Province *uint64 `json:"Province,omitempty" name:"Province"`

	// 登录IP。
	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`

	// 机器名称。
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// 登录时间。
	LoginTime *string `json:"LoginTime,omitempty" name:"LoginTime"`

	// 云镜客户端唯一标识Uuid。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type OpenPort struct {
	// 唯一ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 云镜客户端唯一UUID。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 开放端口号。
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 主机IP。
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// 主机名。
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// 端口对应进程名。
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// 端口对应进程Pid。
	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`

	// 记录创建时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 记录更新时间。
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type OpenPortStatistics struct {
	// 端口号
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 主机数量
	MachineNum *uint64 `json:"MachineNum,omitempty" name:"MachineNum"`
}

// Predefined struct for user
type OpenProVersionPrepaidRequestParams struct {
	// 购买相关参数。
	ChargePrepaid *ChargePrepaid `json:"ChargePrepaid,omitempty" name:"ChargePrepaid"`

	// 需要开通专业版主机信息数组。
	Machines []*ProVersionMachine `json:"Machines,omitempty" name:"Machines"`
}

type OpenProVersionPrepaidRequest struct {
	*tchttp.BaseRequest
	
	// 购买相关参数。
	ChargePrepaid *ChargePrepaid `json:"ChargePrepaid,omitempty" name:"ChargePrepaid"`

	// 需要开通专业版主机信息数组。
	Machines []*ProVersionMachine `json:"Machines,omitempty" name:"Machines"`
}

func (r *OpenProVersionPrepaidRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenProVersionPrepaidRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChargePrepaid")
	delete(f, "Machines")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenProVersionPrepaidRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenProVersionPrepaidResponseParams struct {
	// 订单ID列表。
	DealIds []*string `json:"DealIds,omitempty" name:"DealIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type OpenProVersionPrepaidResponse struct {
	*tchttp.BaseResponse
	Response *OpenProVersionPrepaidResponseParams `json:"Response"`
}

func (r *OpenProVersionPrepaidResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenProVersionPrepaidResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenProVersionRequestParams struct {
	// 云主机类型。
	// <li>CVM：表示虚拟主机</li>
	// <li>BM:  表示黑石物理机</li>
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// 机器所属地域。
	// 如：ap-guangzhou，ap-shanghai
	MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`

	// 主机唯一标识Uuid数组。
	// 黑石的InstanceId，CVM的Uuid
	Quuids []*string `json:"Quuids,omitempty" name:"Quuids"`

	// 活动ID。
	ActivityId *uint64 `json:"ActivityId,omitempty" name:"ActivityId"`
}

type OpenProVersionRequest struct {
	*tchttp.BaseRequest
	
	// 云主机类型。
	// <li>CVM：表示虚拟主机</li>
	// <li>BM:  表示黑石物理机</li>
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// 机器所属地域。
	// 如：ap-guangzhou，ap-shanghai
	MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`

	// 主机唯一标识Uuid数组。
	// 黑石的InstanceId，CVM的Uuid
	Quuids []*string `json:"Quuids,omitempty" name:"Quuids"`

	// 活动ID。
	ActivityId *uint64 `json:"ActivityId,omitempty" name:"ActivityId"`
}

func (r *OpenProVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenProVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MachineType")
	delete(f, "MachineRegion")
	delete(f, "Quuids")
	delete(f, "ActivityId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenProVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenProVersionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type OpenProVersionResponse struct {
	*tchttp.BaseResponse
	Response *OpenProVersionResponseParams `json:"Response"`
}

func (r *OpenProVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenProVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Place struct {
	// 城市 ID。
	CityId *uint64 `json:"CityId,omitempty" name:"CityId"`

	// 省份 ID。
	ProvinceId *uint64 `json:"ProvinceId,omitempty" name:"ProvinceId"`

	// 国家ID，暂只支持国内：1。
	CountryId *uint64 `json:"CountryId,omitempty" name:"CountryId"`
}

type PrivilegeEscalationProcess struct {
	// 数据ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 云镜ID
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 主机ID
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`

	// 主机内网IP
	Hostip *string `json:"Hostip,omitempty" name:"Hostip"`

	// 进程名
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// 进程路径
	FullPath *string `json:"FullPath,omitempty" name:"FullPath"`

	// 执行命令
	CmdLine *string `json:"CmdLine,omitempty" name:"CmdLine"`

	// 用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 用户组
	UserGroup *string `json:"UserGroup,omitempty" name:"UserGroup"`

	// 进程文件权限
	ProcFilePrivilege *string `json:"ProcFilePrivilege,omitempty" name:"ProcFilePrivilege"`

	// 父进程名
	ParentProcName *string `json:"ParentProcName,omitempty" name:"ParentProcName"`

	// 父进程用户名
	ParentProcUser *string `json:"ParentProcUser,omitempty" name:"ParentProcUser"`

	// 父进程用户组
	ParentProcGroup *string `json:"ParentProcGroup,omitempty" name:"ParentProcGroup"`

	// 父进程路径
	ParentProcPath *string `json:"ParentProcPath,omitempty" name:"ParentProcPath"`

	// 进程树
	ProcTree *string `json:"ProcTree,omitempty" name:"ProcTree"`

	// 处理状态
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 发生时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 机器名
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
}

type PrivilegeRule struct {
	// 规则ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 客户端ID
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 进程名
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// 是否S权限
	SMode *uint64 `json:"SMode,omitempty" name:"SMode"`

	// 操作人
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 是否全局规则
	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`

	// 状态(0: 有效 1: 无效)
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// 主机IP
	Hostip *string `json:"Hostip,omitempty" name:"Hostip"`
}

type ProVersionMachine struct {
	// 主机类型。
	// <li>CVM: 虚拟主机</li>
	// <li>BM: 黑石物理机</li>
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// 主机所在地域。
	// 如：ap-guangzhou、ap-beijing
	MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`

	// 主机唯一标识Uuid。
	// 黑石的InstanceId，CVM的Uuid
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

type Process struct {
	// 唯一ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 云镜客户端唯一UUID。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 主机内网IP。
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// 主机名。
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// 进程Pid。
	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`

	// 进程Ppid。
	Ppid *uint64 `json:"Ppid,omitempty" name:"Ppid"`

	// 进程名。
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// 进程用户名。
	Username *string `json:"Username,omitempty" name:"Username"`

	// 所属平台。
	// <li>WIN32：windows32位</li>
	// <li>WIN64：windows64位</li>
	// <li>LINUX32：Linux32位</li>
	// <li>LINUX64：Linux64位</li>
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 进程路径。
	FullPath *string `json:"FullPath,omitempty" name:"FullPath"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type ProcessStatistics struct {
	// 进程名。
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// 主机数量。
	MachineNum *uint64 `json:"MachineNum,omitempty" name:"MachineNum"`
}

// Predefined struct for user
type RecoverMalwaresRequestParams struct {
	// 木马Id数组,单次最大删除不能超过200条
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

type RecoverMalwaresRequest struct {
	*tchttp.BaseRequest
	
	// 木马Id数组,单次最大删除不能超过200条
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *RecoverMalwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecoverMalwaresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecoverMalwaresRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecoverMalwaresResponseParams struct {
	// 恢复成功id数组
	SuccessIds []*uint64 `json:"SuccessIds,omitempty" name:"SuccessIds"`

	// 恢复失败id数组
	FailedIds []*uint64 `json:"FailedIds,omitempty" name:"FailedIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RecoverMalwaresResponse struct {
	*tchttp.BaseResponse
	Response *RecoverMalwaresResponseParams `json:"Response"`
}

func (r *RecoverMalwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecoverMalwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionInfo struct {
	// 地域标志，如 ap-guangzhou，ap-shanghai，ap-beijing
	Region *string `json:"Region,omitempty" name:"Region"`

	// 地域中文名，如华南地区（广州），华东地区（上海金融），华北地区（北京）
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 地域ID
	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`

	// 地域代码，如 gz，sh，bj
	RegionCode *string `json:"RegionCode,omitempty" name:"RegionCode"`
}

// Predefined struct for user
type RenewProVersionRequestParams struct {
	// 购买相关参数。
	ChargePrepaid *ChargePrepaid `json:"ChargePrepaid,omitempty" name:"ChargePrepaid"`

	// 主机唯一ID，对应CVM的uuid、BM的InstanceId。
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

type RenewProVersionRequest struct {
	*tchttp.BaseRequest
	
	// 购买相关参数。
	ChargePrepaid *ChargePrepaid `json:"ChargePrepaid,omitempty" name:"ChargePrepaid"`

	// 主机唯一ID，对应CVM的uuid、BM的InstanceId。
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *RenewProVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewProVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChargePrepaid")
	delete(f, "Quuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewProVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewProVersionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RenewProVersionResponse struct {
	*tchttp.BaseResponse
	Response *RenewProVersionResponseParams `json:"Response"`
}

func (r *RenewProVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewProVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RescanImpactedHostRequestParams struct {
	// 漏洞ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

type RescanImpactedHostRequest struct {
	*tchttp.BaseRequest
	
	// 漏洞ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *RescanImpactedHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RescanImpactedHostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RescanImpactedHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RescanImpactedHostResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RescanImpactedHostResponse struct {
	*tchttp.BaseResponse
	Response *RescanImpactedHostResponseParams `json:"Response"`
}

func (r *RescanImpactedHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RescanImpactedHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReverseShell struct {
	// ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 云镜UUID
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 主机ID
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`

	// 主机内网IP
	Hostip *string `json:"Hostip,omitempty" name:"Hostip"`

	// 目标IP
	DstIp *string `json:"DstIp,omitempty" name:"DstIp"`

	// 目标端口
	DstPort *uint64 `json:"DstPort,omitempty" name:"DstPort"`

	// 进程名
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// 进程路径
	FullPath *string `json:"FullPath,omitempty" name:"FullPath"`

	// 命令详情
	CmdLine *string `json:"CmdLine,omitempty" name:"CmdLine"`

	// 执行用户
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 执行用户组
	UserGroup *string `json:"UserGroup,omitempty" name:"UserGroup"`

	// 父进程名
	ParentProcName *string `json:"ParentProcName,omitempty" name:"ParentProcName"`

	// 父进程用户
	ParentProcUser *string `json:"ParentProcUser,omitempty" name:"ParentProcUser"`

	// 父进程用户组
	ParentProcGroup *string `json:"ParentProcGroup,omitempty" name:"ParentProcGroup"`

	// 父进程路径
	ParentProcPath *string `json:"ParentProcPath,omitempty" name:"ParentProcPath"`

	// 处理状态
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 产生时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 主机名
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// 进程树
	ProcTree *string `json:"ProcTree,omitempty" name:"ProcTree"`
}

type ReverseShellRule struct {
	// 规则ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 客户端ID
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 进程名称
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// 目标IP
	DestIp *string `json:"DestIp,omitempty" name:"DestIp"`

	// 目标端口
	DestPort *string `json:"DestPort,omitempty" name:"DestPort"`

	// 操作人
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 是否全局规则
	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`

	// 状态 (0: 有效 1: 无效)
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// 主机IP
	Hostip *string `json:"Hostip,omitempty" name:"Hostip"`
}

type SecurityDynamic struct {
	// 云镜客户端UUID。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 安全事件发生事件。
	EventTime *string `json:"EventTime,omitempty" name:"EventTime"`

	// 安全事件类型。
	// <li>MALWARE：木马事件</li>
	// <li>NON_LOCAL_LOGIN：异地登录</li>
	// <li>BRUTEATTACK_SUCCESS：密码破解成功</li>
	// <li>VUL：漏洞</li>
	// <li>BASELINE：安全基线</li>
	EventType *string `json:"EventType,omitempty" name:"EventType"`

	// 安全事件消息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 安全事件等级。
	// <li>RISK: 严重</li>
	// <li>HIGH: 高危</li>
	// <li>NORMAL: 中危</li>
	// <li>LOW: 低危</li>
	SecurityLevel *string `json:"SecurityLevel,omitempty" name:"SecurityLevel"`
}

type SecurityTrend struct {
	// 事件时间。
	Date *string `json:"Date,omitempty" name:"Date"`

	// 事件数量。
	EventNum *uint64 `json:"EventNum,omitempty" name:"EventNum"`
}

// Predefined struct for user
type SeparateMalwaresRequestParams struct {
	// 木马事件ID数组。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

type SeparateMalwaresRequest struct {
	*tchttp.BaseRequest
	
	// 木马事件ID数组。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *SeparateMalwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SeparateMalwaresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SeparateMalwaresRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SeparateMalwaresResponseParams struct {
	// 隔离成功的id数组。
	SuccessIds []*uint64 `json:"SuccessIds,omitempty" name:"SuccessIds"`

	// 隔离失败的id数组。
	FailedIds []*uint64 `json:"FailedIds,omitempty" name:"FailedIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SeparateMalwaresResponse struct {
	*tchttp.BaseResponse
	Response *SeparateMalwaresResponseParams `json:"Response"`
}

func (r *SeparateMalwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SeparateMalwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetBashEventsStatusRequestParams struct {
	// ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`

	// 新状态(0-待处理 1-高危 2-正常)
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type SetBashEventsStatusRequest struct {
	*tchttp.BaseRequest
	
	// ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`

	// 新状态(0-待处理 1-高危 2-正常)
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *SetBashEventsStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetBashEventsStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetBashEventsStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetBashEventsStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetBashEventsStatusResponse struct {
	*tchttp.BaseResponse
	Response *SetBashEventsStatusResponseParams `json:"Response"`
}

func (r *SetBashEventsStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetBashEventsStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchBashRulesRequestParams struct {
	// 规则ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 是否禁用
	Disabled *uint64 `json:"Disabled,omitempty" name:"Disabled"`
}

type SwitchBashRulesRequest struct {
	*tchttp.BaseRequest
	
	// 规则ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 是否禁用
	Disabled *uint64 `json:"Disabled,omitempty" name:"Disabled"`
}

func (r *SwitchBashRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchBashRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Disabled")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchBashRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchBashRulesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SwitchBashRulesResponse struct {
	*tchttp.BaseResponse
	Response *SwitchBashRulesResponseParams `json:"Response"`
}

func (r *SwitchBashRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchBashRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {
	// 标签ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 标签名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 服务器数
	Count *uint64 `json:"Count,omitempty" name:"Count"`
}

type TagMachine struct {
	// ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 主机ID
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`

	// 主机名称
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// 主机内网IP
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// 主机外网IP
	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`

	// 主机区域
	MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`

	// 主机区域类型
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
}

// Predefined struct for user
type TrustMaliciousRequestRequestParams struct {
	// 恶意请求记录ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

type TrustMaliciousRequestRequest struct {
	*tchttp.BaseRequest
	
	// 恶意请求记录ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *TrustMaliciousRequestRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TrustMaliciousRequestRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TrustMaliciousRequestRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TrustMaliciousRequestResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type TrustMaliciousRequestResponse struct {
	*tchttp.BaseResponse
	Response *TrustMaliciousRequestResponseParams `json:"Response"`
}

func (r *TrustMaliciousRequestResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TrustMaliciousRequestResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TrustMalwaresRequestParams struct {
	// 木马ID数组。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

type TrustMalwaresRequest struct {
	*tchttp.BaseRequest
	
	// 木马ID数组。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *TrustMalwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TrustMalwaresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TrustMalwaresRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TrustMalwaresResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type TrustMalwaresResponse struct {
	*tchttp.BaseResponse
	Response *TrustMalwaresResponseParams `json:"Response"`
}

func (r *TrustMalwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TrustMalwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UntrustMaliciousRequestRequestParams struct {
	// 受信任记录ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

type UntrustMaliciousRequestRequest struct {
	*tchttp.BaseRequest
	
	// 受信任记录ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *UntrustMaliciousRequestRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UntrustMaliciousRequestRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UntrustMaliciousRequestRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UntrustMaliciousRequestResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UntrustMaliciousRequestResponse struct {
	*tchttp.BaseResponse
	Response *UntrustMaliciousRequestResponseParams `json:"Response"`
}

func (r *UntrustMaliciousRequestResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UntrustMaliciousRequestResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UntrustMalwaresRequestParams struct {
	// 木马ID数组，单次最大处理不能超过200条。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

type UntrustMalwaresRequest struct {
	*tchttp.BaseRequest
	
	// 木马ID数组，单次最大处理不能超过200条。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *UntrustMalwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UntrustMalwaresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UntrustMalwaresRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UntrustMalwaresResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UntrustMalwaresResponse struct {
	*tchttp.BaseResponse
	Response *UntrustMalwaresResponseParams `json:"Response"`
}

func (r *UntrustMalwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UntrustMalwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UsualPlace struct {
	// ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 云镜客户端唯一标识UUID。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 国家 ID。
	CountryId *uint64 `json:"CountryId,omitempty" name:"CountryId"`

	// 省份 ID。
	ProvinceId *uint64 `json:"ProvinceId,omitempty" name:"ProvinceId"`

	// 城市 ID。
	CityId *uint64 `json:"CityId,omitempty" name:"CityId"`
}

type Vul struct {
	// 漏洞种类ID
	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`

	// 漏洞名称
	VulName *string `json:"VulName,omitempty" name:"VulName"`

	// 漏洞危害等级:
	// HIGH：高危
	// MIDDLE：中危
	// LOW：低危
	// NOTICE：提示
	VulLevel *string `json:"VulLevel,omitempty" name:"VulLevel"`

	// 最后扫描时间
	LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`

	// 受影响机器数量
	ImpactedHostNum *uint64 `json:"ImpactedHostNum,omitempty" name:"ImpactedHostNum"`

	// 漏洞状态
	// * UN_OPERATED : 待处理
	// * FIXED : 已修复
	VulStatus *string `json:"VulStatus,omitempty" name:"VulStatus"`
}

type WeeklyReport struct {
	// 周报开始时间。
	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`

	// 周报结束时间。
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

type WeeklyReportBruteAttack struct {
	// 主机IP。
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// 被破解用户名。
	Username *string `json:"Username,omitempty" name:"Username"`

	// 源IP。
	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`

	// 尝试次数。
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 攻击时间。
	AttackTime *string `json:"AttackTime,omitempty" name:"AttackTime"`
}

type WeeklyReportMalware struct {
	// 主机IP。
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// 木马文件路径。
	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`

	// 木马文件MD5值。
	Md5 *string `json:"Md5,omitempty" name:"Md5"`

	// 木马发现时间。
	FindTime *string `json:"FindTime,omitempty" name:"FindTime"`

	// 当前木马状态。
	// <li>UN_OPERATED：未处理</li>
	// <li>SEGREGATED：已隔离</li>
	// <li>TRUSTED：已信任</li>
	// <li>SEPARATING：隔离中</li>
	// <li>RECOVERING：恢复中</li>
	Status *string `json:"Status,omitempty" name:"Status"`
}

type WeeklyReportNonlocalLoginPlace struct {
	// 主机IP。
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// 用户名。
	Username *string `json:"Username,omitempty" name:"Username"`

	// 源IP。
	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`

	// 国家ID。
	Country *uint64 `json:"Country,omitempty" name:"Country"`

	// 省份ID。
	Province *uint64 `json:"Province,omitempty" name:"Province"`

	// 城市ID。
	City *uint64 `json:"City,omitempty" name:"City"`

	// 登录时间。
	LoginTime *string `json:"LoginTime,omitempty" name:"LoginTime"`
}

type WeeklyReportVul struct {
	// 主机内网IP。
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// 漏洞名称。
	VulName *string `json:"VulName,omitempty" name:"VulName"`

	// 漏洞类型。
	// <li> WEB : Web漏洞</li>
	// <li> SYSTEM :系统组件漏洞</li>
	// <li> BASELINE : 安全基线</li>
	VulType *string `json:"VulType,omitempty" name:"VulType"`

	// 漏洞描述。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 漏洞状态。
	// <li> UN_OPERATED : 待处理</li>
	// <li> SCANING : 扫描中</li>
	// <li> FIXED : 已修复</li>
	VulStatus *string `json:"VulStatus,omitempty" name:"VulStatus"`

	// 最后扫描时间。
	LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`
}