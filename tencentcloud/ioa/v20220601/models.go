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

package v20220601

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type ComplexRule struct {
	// 简单规则表达式
	SimpleRules []*SimpleRule `json:"SimpleRules,omitnil,omitempty" name:"SimpleRules"`

	// 表达式间逻辑关系
	Relation *string `json:"Relation,omitnil,omitempty" name:"Relation"`
}

type Condition struct {
	// Filters 条件过滤
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// FilterGroups 条件过滤组
	FilterGroups []*FilterGroup `json:"FilterGroups,omitnil,omitempty" name:"FilterGroups"`

	// Sort 排序字段
	Sort *Sort `json:"Sort,omitnil,omitempty" name:"Sort"`

	// PageSize 每页获取数(只支持32位)
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// PageNum 获取第几页(只支持32位)
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`
}

// Predefined struct for user
type CreateDeviceVirtualGroupRequestParams struct {
	// 必填，终端自定义分组名
	DeviceVirtualGroupName *string `json:"DeviceVirtualGroupName,omitnil,omitempty" name:"DeviceVirtualGroupName"`

	// 详情
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 必填，系统类型（0: win，1：linux，2: mac，3: win_srv，4：android，5：ios ）(只支持32位)
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`

	// 必填，分组类型（0:手动分组；非0为自动划分分组；具体枚举值为：1:自动每小时划分分组、2:自动每天划分分组、3:自定义时间划分分组）(只支持32位)
	TimeType *int64 `json:"TimeType,omitnil,omitempty" name:"TimeType"`

	// 选填，TimeType=3时的自动划分时间，其他情况为0（单位min）(只支持32位)
	AutoMinute *int64 `json:"AutoMinute,omitnil,omitempty" name:"AutoMinute"`

	// 选填，手动分组不填，自动划分分组的划分规则数据
	AutoRules *ComplexRule `json:"AutoRules,omitnil,omitempty" name:"AutoRules"`
}

type CreateDeviceVirtualGroupRequest struct {
	*tchttp.BaseRequest
	
	// 必填，终端自定义分组名
	DeviceVirtualGroupName *string `json:"DeviceVirtualGroupName,omitnil,omitempty" name:"DeviceVirtualGroupName"`

	// 详情
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 必填，系统类型（0: win，1：linux，2: mac，3: win_srv，4：android，5：ios ）(只支持32位)
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`

	// 必填，分组类型（0:手动分组；非0为自动划分分组；具体枚举值为：1:自动每小时划分分组、2:自动每天划分分组、3:自定义时间划分分组）(只支持32位)
	TimeType *int64 `json:"TimeType,omitnil,omitempty" name:"TimeType"`

	// 选填，TimeType=3时的自动划分时间，其他情况为0（单位min）(只支持32位)
	AutoMinute *int64 `json:"AutoMinute,omitnil,omitempty" name:"AutoMinute"`

	// 选填，手动分组不填，自动划分分组的划分规则数据
	AutoRules *ComplexRule `json:"AutoRules,omitnil,omitempty" name:"AutoRules"`
}

func (r *CreateDeviceVirtualGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDeviceVirtualGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceVirtualGroupName")
	delete(f, "Description")
	delete(f, "OsType")
	delete(f, "TimeType")
	delete(f, "AutoMinute")
	delete(f, "AutoRules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDeviceVirtualGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDeviceVirtualGroupResponseParams struct {
	// 响应返回的data
	Data *CreateDeviceVirtualGroupRspData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDeviceVirtualGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateDeviceVirtualGroupResponseParams `json:"Response"`
}

func (r *CreateDeviceVirtualGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDeviceVirtualGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDeviceVirtualGroupRspData struct {
	// 返回的自定义分组id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`
}

type DescribeAccountGroupsData struct {
	// 账号分组名全路径，点分格式
	NamePath *string `json:"NamePath,omitnil,omitempty" name:"NamePath"`

	// 账号分组ID全路径，数组格式
	IdPathArr []*int64 `json:"IdPathArr,omitnil,omitempty" name:"IdPathArr"`

	// 扩展信息
	ExtraInfo *string `json:"ExtraInfo,omitnil,omitempty" name:"ExtraInfo"`

	// 最后更新时间
	Utime *string `json:"Utime,omitnil,omitempty" name:"Utime"`

	// 父分组ID
	ParentId *int64 `json:"ParentId,omitnil,omitempty" name:"ParentId"`

	// 源账号组织ID。使用第三方导入用户源时，记录该分组在源组织架构下的分组ID
	OrgId *string `json:"OrgId,omitnil,omitempty" name:"OrgId"`

	// 分组名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分组ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 分组描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 同步数据源
	Source *int64 `json:"Source,omitnil,omitempty" name:"Source"`

	// 账号分组ID全路径，点分格式
	IdPath *string `json:"IdPath,omitnil,omitempty" name:"IdPath"`

	// 创建时间
	Itime *string `json:"Itime,omitnil,omitempty" name:"Itime"`

	// 父源账号组织ID。使用第三方导入用户源时，记录该分组在源组织架构下的分组ID
	ParentOrgId *string `json:"ParentOrgId,omitnil,omitempty" name:"ParentOrgId"`

	// 导入类型
	ImportType *string `json:"ImportType,omitnil,omitempty" name:"ImportType"`

	// miniIAM id
	MiniIamId *string `json:"MiniIamId,omitnil,omitempty" name:"MiniIamId"`

	// 该分组下含子组的所有用户总数
	UserTotal *int64 `json:"UserTotal,omitnil,omitempty" name:"UserTotal"`

	// 是否叶子节点
	IsLeaf *bool `json:"IsLeaf,omitnil,omitempty" name:"IsLeaf"`

	// 是否该账户的直接权限
	ReadOnly *bool `json:"ReadOnly,omitnil,omitempty" name:"ReadOnly"`

	// 最新一次同步任务的结果
	LatestSyncResult *string `json:"LatestSyncResult,omitnil,omitempty" name:"LatestSyncResult"`

	// 最新一次同步任务的结束时间
	LatestSyncTime *string `json:"LatestSyncTime,omitnil,omitempty" name:"LatestSyncTime"`
}

type DescribeAccountGroupsPageResp struct {
	// 账户分响应对象集合
	Items []*DescribeAccountGroupsData `json:"Items,omitnil,omitempty" name:"Items"`

	// 分页公共对象
	Page *Paging `json:"Page,omitnil,omitempty" name:"Page"`
}

// Predefined struct for user
type DescribeAccountGroupsRequestParams struct {
	// 搜索范围：0-仅当前分组的直接子组，1-当前分组的所有子组。默认为0。
	Deepin *int64 `json:"Deepin,omitnil,omitempty" name:"Deepin"`

	// 查询条件
	// 
	// 过滤参数
	// 1、Name，string类型，按分组名过滤
	// 是否必填：否
	// 操作符: like
	// 
	// 排序条件
	// 1、Itime，string类型，按分组创建时间排序
	// 是否必填：否
	// 2、Utime，string类型，按分组更新时间排序
	// 是否必填：否
	Condition *Condition `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 父分组ID，获取该分组下的子组信息。默认查询全网根分组下子组信息。
	ParentId *int64 `json:"ParentId,omitnil,omitempty" name:"ParentId"`
}

type DescribeAccountGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 搜索范围：0-仅当前分组的直接子组，1-当前分组的所有子组。默认为0。
	Deepin *int64 `json:"Deepin,omitnil,omitempty" name:"Deepin"`

	// 查询条件
	// 
	// 过滤参数
	// 1、Name，string类型，按分组名过滤
	// 是否必填：否
	// 操作符: like
	// 
	// 排序条件
	// 1、Itime，string类型，按分组创建时间排序
	// 是否必填：否
	// 2、Utime，string类型，按分组更新时间排序
	// 是否必填：否
	Condition *Condition `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 父分组ID，获取该分组下的子组信息。默认查询全网根分组下子组信息。
	ParentId *int64 `json:"ParentId,omitnil,omitempty" name:"ParentId"`
}

func (r *DescribeAccountGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Deepin")
	delete(f, "Condition")
	delete(f, "ParentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountGroupsResponseParams struct {
	// 账号分组详情响应数据
	Data *DescribeAccountGroupsPageResp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAccountGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccountGroupsResponseParams `json:"Response"`
}

func (r *DescribeAccountGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDevicesPageRsp struct {
	// 数据分页信息
	Paging *Paging `json:"Paging,omitnil,omitempty" name:"Paging"`

	// 业务响应数据
	Items []*DeviceDetail `json:"Items,omitnil,omitempty" name:"Items"`
}

// Predefined struct for user
type DescribeDevicesRequestParams struct {
	// 过滤条件参数（字段含义请参考接口返回值）
	// 
	// - Mid, 类型String，支持操作：【eq，like，ilike】，支持排序
	// - Name, 类型String，支持操作：【eq，like，ilike】，支持排序
	// - Itime, 类型String，支持操作：【eq，like，ilike】，支持排序
	// - UserName, 类型String，支持操作：【eq，like，ilike】，支持排序
	// - MacAddr, 类型String，支持操作：【eq，like，ilike】，支持排序
	// - UserId, 类型String，支持操作：【eq，like，ilike】，支持排序
	// - Ip, 类型String，支持操作：【eq，like，ilike】，支持排序
	// - Tags，类型String，支持操作：【eq，like，ilike】，支持排序
	// - LocalIpList，类型String，支持操作：【eq，like，ilike】，支持排序
	// - SerialNum，类型String，支持操作：【eq，like，ilike】，支持排序
	// - Version，类型String，支持操作：【eq，like，ilike】，支持排序
	// - StrVersion，类型String，支持操作：【eq，like，ilike】，支持排序
	// - RtpStatus，类型String，支持操作：【eq，like，ilike】，**不支持排序**
	// - HostName，类型String，支持操作：【eq，like，ilike】，支持排序
	// - IoaUserName，类型String，支持操作：【eq，like，ilike】，支持排序
	// - GroupName，类型String，支持操作：【eq，like，ilike】，支持排序
	// - CriticalVulListCount，**类型Int**，支持操作：【eq】，**不支持排序**
	// - RiskCount，**类型Int**，支持操作：【eq】，**不支持排序**
	// - VulVersion，类型String，支持操作：【eq，like，ilike】，**不支持排序**
	// - Virusver，类型String，支持操作：【eq，like，ilike】，**不支持排序**
	// - SysRepver，类型String，支持操作：【eq，like，ilike】，**不支持排序**
	// - BaseBoardSn，类型String，支持操作：【eq，like，ilike】，支持排序
	// - Os，类型String，支持操作：【eq，like，ilike】，支持排序
	// - ConnActiveTime，类型String，支持操作：【eq，like，ilike】，**不支持排序**
	// - FirewallStatus，**类型Int**，支持操作：【eq】，**不支持排序**
	// - ProfileName，类型String，支持操作：【eq，like，ilike】，支持排序
	// - DomainName，类型String，支持操作：【eq，like，ilike】，支持排序
	// - SysRepVersion，类型String，支持操作：【eq，like，ilike】，支持排序
	// - VirusVer，类型String，支持操作：【eq，like，ilike】，支持排序
	// - Cpu，类型String，支持操作：【eq，like，ilike】，支持排序
	// - Memory，类型String，支持操作：【eq，like，ilike】，支持排序
	// - HardDiskSize，类型String，支持操作：【eq，like，ilike】，支持排序
	// - HardwareChangeCount，**类型Int**，支持操作：【eq】，支持排序
	// - AccountName，类型String，支持操作：【like.ilike】，支持排序
	// - AccountGroupName，类型String，支持操作：【like.ilike】，支持排序
	// - ScreenRecordingPermission，**类型Int**，支持操作：【eq】，支持排序
	// - DiskAccessPermission，**类型Int**，支持操作：【eq】，支持排序
	// 
	// 
	// 
	// 
	// 
	// 分页参数
	// - PageNum 从1开始，小于等于0时使用默认参数
	// - PageSize 最大值5000，最好不超过100
	Condition *Condition `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 【和GroupIds必须有一个填写】设备分组id（需要和OsType匹配），下面是私有化场景下默认id：
	// id-名称-操作系统
	// 1	全网终端	Win
	// 2	未分组终端	Win
	// 30000000	服务器	Win
	// 40000101	全网终端	Linux
	// 40000102	未分组终端	Linux
	// 40000103	服务器	Linux
	// 40000201	全网终端	macOS
	// 40000202	未分组终端	macOS
	// 40000203	服务器	macOS
	// 40000401	全网终端	Android
	// 40000402	未分组终端	Android
	// 40000501	全网终端	iOS
	// 40000502	未分组终端	iOS
	// 
	// 
	// SaaS需要调用分组接口DescribeDeviceChildGroups获取对应分组id
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 【必填】操作系统类型（0: win，1：linux，2: mac，4：android，5：ios   默认值0），需要和GroupId或者GroupIds匹配
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`

	// 在线状态 （2表示在线，0或者1表示离线）
	OnlineStatus *int64 `json:"OnlineStatus,omitnil,omitempty" name:"OnlineStatus"`

	// 过滤条件--兼容旧接口,参数同Condition
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段--兼容旧接口,参数同Condition
	Sort *Sort `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 获取第几页--兼容旧接口,参数同Condition
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 每页获取数--兼容旧接口,参数同Condition
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 授权状态： 4基础授权 5高级授权
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 【和GroupId必须有一个填写】设备分组id列表（需要和OsType匹配）
	GroupIds []*int64 `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`
}

type DescribeDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件参数（字段含义请参考接口返回值）
	// 
	// - Mid, 类型String，支持操作：【eq，like，ilike】，支持排序
	// - Name, 类型String，支持操作：【eq，like，ilike】，支持排序
	// - Itime, 类型String，支持操作：【eq，like，ilike】，支持排序
	// - UserName, 类型String，支持操作：【eq，like，ilike】，支持排序
	// - MacAddr, 类型String，支持操作：【eq，like，ilike】，支持排序
	// - UserId, 类型String，支持操作：【eq，like，ilike】，支持排序
	// - Ip, 类型String，支持操作：【eq，like，ilike】，支持排序
	// - Tags，类型String，支持操作：【eq，like，ilike】，支持排序
	// - LocalIpList，类型String，支持操作：【eq，like，ilike】，支持排序
	// - SerialNum，类型String，支持操作：【eq，like，ilike】，支持排序
	// - Version，类型String，支持操作：【eq，like，ilike】，支持排序
	// - StrVersion，类型String，支持操作：【eq，like，ilike】，支持排序
	// - RtpStatus，类型String，支持操作：【eq，like，ilike】，**不支持排序**
	// - HostName，类型String，支持操作：【eq，like，ilike】，支持排序
	// - IoaUserName，类型String，支持操作：【eq，like，ilike】，支持排序
	// - GroupName，类型String，支持操作：【eq，like，ilike】，支持排序
	// - CriticalVulListCount，**类型Int**，支持操作：【eq】，**不支持排序**
	// - RiskCount，**类型Int**，支持操作：【eq】，**不支持排序**
	// - VulVersion，类型String，支持操作：【eq，like，ilike】，**不支持排序**
	// - Virusver，类型String，支持操作：【eq，like，ilike】，**不支持排序**
	// - SysRepver，类型String，支持操作：【eq，like，ilike】，**不支持排序**
	// - BaseBoardSn，类型String，支持操作：【eq，like，ilike】，支持排序
	// - Os，类型String，支持操作：【eq，like，ilike】，支持排序
	// - ConnActiveTime，类型String，支持操作：【eq，like，ilike】，**不支持排序**
	// - FirewallStatus，**类型Int**，支持操作：【eq】，**不支持排序**
	// - ProfileName，类型String，支持操作：【eq，like，ilike】，支持排序
	// - DomainName，类型String，支持操作：【eq，like，ilike】，支持排序
	// - SysRepVersion，类型String，支持操作：【eq，like，ilike】，支持排序
	// - VirusVer，类型String，支持操作：【eq，like，ilike】，支持排序
	// - Cpu，类型String，支持操作：【eq，like，ilike】，支持排序
	// - Memory，类型String，支持操作：【eq，like，ilike】，支持排序
	// - HardDiskSize，类型String，支持操作：【eq，like，ilike】，支持排序
	// - HardwareChangeCount，**类型Int**，支持操作：【eq】，支持排序
	// - AccountName，类型String，支持操作：【like.ilike】，支持排序
	// - AccountGroupName，类型String，支持操作：【like.ilike】，支持排序
	// - ScreenRecordingPermission，**类型Int**，支持操作：【eq】，支持排序
	// - DiskAccessPermission，**类型Int**，支持操作：【eq】，支持排序
	// 
	// 
	// 
	// 
	// 
	// 分页参数
	// - PageNum 从1开始，小于等于0时使用默认参数
	// - PageSize 最大值5000，最好不超过100
	Condition *Condition `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 【和GroupIds必须有一个填写】设备分组id（需要和OsType匹配），下面是私有化场景下默认id：
	// id-名称-操作系统
	// 1	全网终端	Win
	// 2	未分组终端	Win
	// 30000000	服务器	Win
	// 40000101	全网终端	Linux
	// 40000102	未分组终端	Linux
	// 40000103	服务器	Linux
	// 40000201	全网终端	macOS
	// 40000202	未分组终端	macOS
	// 40000203	服务器	macOS
	// 40000401	全网终端	Android
	// 40000402	未分组终端	Android
	// 40000501	全网终端	iOS
	// 40000502	未分组终端	iOS
	// 
	// 
	// SaaS需要调用分组接口DescribeDeviceChildGroups获取对应分组id
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 【必填】操作系统类型（0: win，1：linux，2: mac，4：android，5：ios   默认值0），需要和GroupId或者GroupIds匹配
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`

	// 在线状态 （2表示在线，0或者1表示离线）
	OnlineStatus *int64 `json:"OnlineStatus,omitnil,omitempty" name:"OnlineStatus"`

	// 过滤条件--兼容旧接口,参数同Condition
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段--兼容旧接口,参数同Condition
	Sort *Sort `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 获取第几页--兼容旧接口,参数同Condition
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 每页获取数--兼容旧接口,参数同Condition
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 授权状态： 4基础授权 5高级授权
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 【和GroupId必须有一个填写】设备分组id列表（需要和OsType匹配）
	GroupIds []*int64 `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`
}

func (r *DescribeDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Condition")
	delete(f, "GroupId")
	delete(f, "OsType")
	delete(f, "OnlineStatus")
	delete(f, "Filters")
	delete(f, "Sort")
	delete(f, "PageNum")
	delete(f, "PageSize")
	delete(f, "Status")
	delete(f, "GroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDevicesResponseParams struct {
	// 分页的data数据
	Data *DescribeDevicesPageRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDevicesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDevicesResponseParams `json:"Response"`
}

func (r *DescribeDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLocalAccountAccountGroupsData struct {
	// 组Id(只支持32位)
	AccountGroupId *int64 `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`
}

type DescribeLocalAccountsData struct {
	// uid，数据库中唯一
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 账号，登录账号
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 账号id，同Id字段
	AccountId *int64 `json:"AccountId,omitnil,omitempty" name:"AccountId"`

	// 账号所在的分组id
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 账号所在的分组名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 账号所在的分组名称路径，用英文.分割
	NamePath *string `json:"NamePath,omitnil,omitempty" name:"NamePath"`

	// 账号来源,0表示本地账号(只支持32位)
	Source *int64 `json:"Source,omitnil,omitempty" name:"Source"`

	// 账号状态,0禁用，1启用(只支持32位)
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 账号的创建时间
	Itime *string `json:"Itime,omitnil,omitempty" name:"Itime"`

	// 账号的最后更新时间
	Utime *string `json:"Utime,omitnil,omitempty" name:"Utime"`

	// 账号的扩展信息，包含邮箱、手机号、身份证、职位等信息
	ExtraInfo *string `json:"ExtraInfo,omitnil,omitempty" name:"ExtraInfo"`

	// 用户风险等级，枚举：none, low, middle, high
	RiskLevel *string `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 所属组
	AccountGroups []*DescribeLocalAccountAccountGroupsData `json:"AccountGroups,omitnil,omitempty" name:"AccountGroups"`

	// 绑定手机端设备数
	MobileBindNum *int64 `json:"MobileBindNum,omitnil,omitempty" name:"MobileBindNum"`

	// 绑定Pc端设备数
	PcBindNum *int64 `json:"PcBindNum,omitnil,omitempty" name:"PcBindNum"`

	// 账号在线状态 1：在线 2：离线
	OnlineStatus *int64 `json:"OnlineStatus,omitnil,omitempty" name:"OnlineStatus"`

	// 账号活跃状态 1：活跃 2：非活跃
	ActiveStatus *int64 `json:"ActiveStatus,omitnil,omitempty" name:"ActiveStatus"`

	// 账号登录时间
	LoginTime *string `json:"LoginTime,omitnil,omitempty" name:"LoginTime"`

	// 账号登出时间
	LogoutTime *string `json:"LogoutTime,omitnil,omitempty" name:"LogoutTime"`
}

type DescribeLocalAccountsPage struct {
	// 公共分页对象
	Page *Paging `json:"Page,omitnil,omitempty" name:"Page"`

	// 获取账号列表响应的单个对象
	Items []*DescribeLocalAccountsData `json:"Items,omitnil,omitempty" name:"Items"`
}

// Predefined struct for user
type DescribeLocalAccountsRequestParams struct {
	// 查询条件：过滤或排序
	// 1、UserName，string类型，姓名
	// 是否必填：否
	// 过滤支持：是，支持eq、like、ilike
	// 排序支持：否
	// 2、UserId，string类型，账户
	// 是否必填：否
	// 过滤支持：是，支持eq、like、ilike
	// 排序支持：否
	// 3、Phone，string类型，手机号
	// 是否必填：否
	// 过滤支持：是，支持eq、like、ilike
	// 排序支持：否
	Condition *Condition `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 获取账号的分组ID，不传默认获取全网根账号组
	AccountGroupId *int64 `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// 是否仅展示当前目录下用户 1： 递归显示 2：仅显示当前目录下用户
	ShowFlag *int64 `json:"ShowFlag,omitnil,omitempty" name:"ShowFlag"`
}

type DescribeLocalAccountsRequest struct {
	*tchttp.BaseRequest
	
	// 查询条件：过滤或排序
	// 1、UserName，string类型，姓名
	// 是否必填：否
	// 过滤支持：是，支持eq、like、ilike
	// 排序支持：否
	// 2、UserId，string类型，账户
	// 是否必填：否
	// 过滤支持：是，支持eq、like、ilike
	// 排序支持：否
	// 3、Phone，string类型，手机号
	// 是否必填：否
	// 过滤支持：是，支持eq、like、ilike
	// 排序支持：否
	Condition *Condition `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 获取账号的分组ID，不传默认获取全网根账号组
	AccountGroupId *int64 `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// 是否仅展示当前目录下用户 1： 递归显示 2：仅显示当前目录下用户
	ShowFlag *int64 `json:"ShowFlag,omitnil,omitempty" name:"ShowFlag"`
}

func (r *DescribeLocalAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLocalAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Condition")
	delete(f, "AccountGroupId")
	delete(f, "ShowFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLocalAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLocalAccountsResponseParams struct {
	// 获取账号列表响应的分页对象
	Data *DescribeLocalAccountsPage `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLocalAccountsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLocalAccountsResponseParams `json:"Response"`
}

func (r *DescribeLocalAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLocalAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRootAccountGroupRequestParams struct {

}

type DescribeRootAccountGroupRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeRootAccountGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRootAccountGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRootAccountGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRootAccountGroupResponseParams struct {
	// 账号根分组响应详情
	Data *GetAccountGroupData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRootAccountGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRootAccountGroupResponseParams `json:"Response"`
}

func (r *DescribeRootAccountGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRootAccountGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeviceDetail struct {
	// 设备ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 设备唯一标识码，在ioa中每个设备有唯一标识码
	Mid *string `json:"Mid,omitnil,omitempty" name:"Mid"`

	// 终端名（设备名）
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 设备所在分组ID
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// OS平台，0：Windows 、1： Linux、 2：macOS 、4： Android、 5: iOS。默认是0
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`

	// 设备IP地址（出口IP）
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 在线状态，2：在线、0或者1:离线
	OnlineStatus *int64 `json:"OnlineStatus,omitnil,omitempty" name:"OnlineStatus"`

	// 客户端版本号-大整数
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 客户端版本号-点分字符串
	StrVersion *string `json:"StrVersion,omitnil,omitempty" name:"StrVersion"`

	// 首次在线时间
	Itime *string `json:"Itime,omitnil,omitempty" name:"Itime"`

	// 最后一次在线时间
	ConnActiveTime *string `json:"ConnActiveTime,omitnil,omitempty" name:"ConnActiveTime"`

	// 设备是否加锁 ，1：锁定 0或者2：未锁定。
	Locked *int64 `json:"Locked,omitnil,omitempty" name:"Locked"`

	// 设备本地IP列表, 包括IP
	LocalIpList *string `json:"LocalIpList,omitnil,omitempty" name:"LocalIpList"`

	// 宿主机id（需要宿主机也安装iOA才能显示）
	HostId *int64 `json:"HostId,omitnil,omitempty" name:"HostId"`

	// 设备所属分组名
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 设备所属分组路径
	GroupNamePath *string `json:"GroupNamePath,omitnil,omitempty" name:"GroupNamePath"`

	// 未修复高危漏洞数(只支持32位)
	CriticalVulListCount *int64 `json:"CriticalVulListCount,omitnil,omitempty" name:"CriticalVulListCount"`

	// 设备名，和Name相同
	ComputerName *string `json:"ComputerName,omitnil,omitempty" name:"ComputerName"`

	// 登录域名
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// MAC地址
	MacAddr *string `json:"MacAddr,omitnil,omitempty" name:"MacAddr"`

	// 漏洞数
	VulCount *int64 `json:"VulCount,omitnil,omitempty" name:"VulCount"`

	// 病毒风险数
	RiskCount *int64 `json:"RiskCount,omitnil,omitempty" name:"RiskCount"`

	// 病毒库版本
	VirusVer *string `json:"VirusVer,omitnil,omitempty" name:"VirusVer"`

	// 漏洞库版本
	VulVersion *string `json:"VulVersion,omitnil,omitempty" name:"VulVersion"`

	// 系统修复引擎版本
	SysRepVersion *string `json:"SysRepVersion,omitnil,omitempty" name:"SysRepVersion"`

	// 高危补丁列表
	VulCriticalList []*string `json:"VulCriticalList,omitnil,omitempty" name:"VulCriticalList"`

	// 标签
	Tags *string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 终端用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 防火墙状态，不等于0表示开启
	FirewallStatus *int64 `json:"FirewallStatus,omitnil,omitempty" name:"FirewallStatus"`

	// SN序列号
	SerialNum *string `json:"SerialNum,omitnil,omitempty" name:"SerialNum"`

	// 设备管控策略版本
	DeviceStrategyVer *string `json:"DeviceStrategyVer,omitnil,omitempty" name:"DeviceStrategyVer"`

	// NGN策略版本
	NGNStrategyVer *string `json:"NGNStrategyVer,omitnil,omitempty" name:"NGNStrategyVer"`

	// 最近登录账户的账号
	IOAUserName *string `json:"IOAUserName,omitnil,omitempty" name:"IOAUserName"`

	// 设备管控新策略
	DeviceNewStrategyVer *string `json:"DeviceNewStrategyVer,omitnil,omitempty" name:"DeviceNewStrategyVer"`

	// NGN策略新版本
	NGNNewStrategyVer *string `json:"NGNNewStrategyVer,omitnil,omitempty" name:"NGNNewStrategyVer"`

	// 宿主机名称（需要宿主机也安装iOA才能显示）
	HostName *string `json:"HostName,omitnil,omitempty" name:"HostName"`

	// 主板序列号
	BaseBoardSn *string `json:"BaseBoardSn,omitnil,omitempty" name:"BaseBoardSn"`

	// 绑定账户名称
	AccountUsers *string `json:"AccountUsers,omitnil,omitempty" name:"AccountUsers"`

	// 身份策略版本
	IdentityStrategyVer *string `json:"IdentityStrategyVer,omitnil,omitempty" name:"IdentityStrategyVer"`

	// 身份策略新版本
	IdentityNewStrategyVer *string `json:"IdentityNewStrategyVer,omitnil,omitempty" name:"IdentityNewStrategyVer"`

	// 最近登录账号部门
	AccountGroupName *string `json:"AccountGroupName,omitnil,omitempty" name:"AccountGroupName"`

	// 最近登录账户的姓名
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// 账号组id
	AccountGroupId *int64 `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`
}

type Filter struct {
	// 过滤字段
	Field *string `json:"Field,omitnil,omitempty" name:"Field"`

	// 过滤方式： eq:等于,net:不等于,like,nlike,gt:大于,lt:小于,egt:大于等于,elt:小于等于。具体支持哪些过滤方式，结合具体接口字段描述来定
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 过滤条件
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type FilterGroup struct {
	// Filters 条件过滤
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type GetAccountGroupData struct {
	// 分组名称全路径，点分格式
	NamePath *string `json:"NamePath,omitnil,omitempty" name:"NamePath"`

	// 分组ID全路径，数组格式
	IdPathArr []*int64 `json:"IdPathArr,omitnil,omitempty" name:"IdPathArr"`

	// 分组扩展信息
	ExtraInfo *string `json:"ExtraInfo,omitnil,omitempty" name:"ExtraInfo"`

	// 最后更新时间
	Utime *string `json:"Utime,omitnil,omitempty" name:"Utime"`

	// 当前分组的父分组ID
	ParentId *uint64 `json:"ParentId,omitnil,omitempty" name:"ParentId"`

	// 源账号组ID，该字段仅适用于第三方同步的组织架构，通过OrgId-Id构成源组织架构分组ID-现组织架构分组ID映射关系
	OrgId *string `json:"OrgId,omitnil,omitempty" name:"OrgId"`

	// 分组名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分组ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 分组描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 分组导入源(只支持32位)
	Source *uint64 `json:"Source,omitnil,omitempty" name:"Source"`

	// 分组ID全路径，点分格式
	IdPath *string `json:"IdPath,omitnil,omitempty" name:"IdPath"`

	// 创建时间
	Itime *string `json:"Itime,omitnil,omitempty" name:"Itime"`

	// 父源账号组ID，该字段仅适用于第三方同步的组织架构
	ParentOrgId *string `json:"ParentOrgId,omitnil,omitempty" name:"ParentOrgId"`

	// 导入信息,json格式
	Import *string `json:"Import,omitnil,omitempty" name:"Import"`

	// 是否开启导入架构
	ImportEnable *bool `json:"ImportEnable,omitnil,omitempty" name:"ImportEnable"`

	// 导入类型
	ImportType *string `json:"ImportType,omitnil,omitempty" name:"ImportType"`

	// miniIAMId，MiniIAM源才有
	MiniIamId *string `json:"MiniIamId,omitnil,omitempty" name:"MiniIamId"`
}

type Paging struct {
	// 每页条数
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页码
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 总页数
	PageCount *uint64 `json:"PageCount,omitnil,omitempty" name:"PageCount"`

	// 记录总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`
}

type RuleExpression struct {
	// 规则元数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*RuleItem `json:"Items,omitnil,omitempty" name:"Items"`

	// 关系
	// 注意：此字段可能返回 null，表示取不到有效值。
	Relation *string `json:"Relation,omitnil,omitempty" name:"Relation"`
}

type RuleItem struct {
	// 字段名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 操作关系（等于、不等于、包含、不包含）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operate *string `json:"Operate,omitnil,omitempty" name:"Operate"`

	// 内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 内容，v2多值版本使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type SimpleRule struct {
	// 规则表达式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Expressions []*RuleExpression `json:"Expressions,omitnil,omitempty" name:"Expressions"`

	// 表达式间逻辑关系
	// 注意：此字段可能返回 null，表示取不到有效值。
	Relation *string `json:"Relation,omitnil,omitempty" name:"Relation"`
}

type Sort struct {
	// 排序字段
	Field *string `json:"Field,omitnil,omitempty" name:"Field"`

	// 排序方式
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}