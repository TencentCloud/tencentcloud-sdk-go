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

type Condition struct {
	// Filters 条件过滤
	// 注意：此字段可能返回 null，表示取不到有效值。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// FilterGroups 条件过滤组
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilterGroups []*FilterGroup `json:"FilterGroups,omitnil,omitempty" name:"FilterGroups"`

	// Sort 排序字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sort *Sort `json:"Sort,omitnil,omitempty" name:"Sort"`

	// PageSize 每页获取数(只支持32位)
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// PageNum 获取第几页(只支持32位)
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`
}

type DescribeAccountGroupsData struct {
	// 名称path
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamePath *string `json:"NamePath,omitnil,omitempty" name:"NamePath"`

	// id patch数组(只支持32位)
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdPathArr []*int64 `json:"IdPathArr,omitnil,omitempty" name:"IdPathArr"`

	// 扩展信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraInfo *string `json:"ExtraInfo,omitnil,omitempty" name:"ExtraInfo"`

	// 最后更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Utime *string `json:"Utime,omitnil,omitempty" name:"Utime"`

	// 父id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentId *int64 `json:"ParentId,omitnil,omitempty" name:"ParentId"`

	// 组织id
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrgId *string `json:"OrgId,omitnil,omitempty" name:"OrgId"`

	// 账户组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 同步数据源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *int64 `json:"Source,omitnil,omitempty" name:"Source"`

	// id path
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdPath *string `json:"IdPath,omitnil,omitempty" name:"IdPath"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Itime *string `json:"Itime,omitnil,omitempty" name:"Itime"`

	// 父组织id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentOrgId *string `json:"ParentOrgId,omitnil,omitempty" name:"ParentOrgId"`

	// 导入类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImportType *string `json:"ImportType,omitnil,omitempty" name:"ImportType"`

	// miniIAM id
	// 注意：此字段可能返回 null，表示取不到有效值。
	MiniIamId *string `json:"MiniIamId,omitnil,omitempty" name:"MiniIamId"`

	// 该分组下用户总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserTotal *int64 `json:"UserTotal,omitnil,omitempty" name:"UserTotal"`

	// 是否叶子节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsLeaf *bool `json:"IsLeaf,omitnil,omitempty" name:"IsLeaf"`

	// 是否该账户的直接权限
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReadOnly *bool `json:"ReadOnly,omitnil,omitempty" name:"ReadOnly"`

	// 最新一次同步任务的结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestSyncResult *string `json:"LatestSyncResult,omitnil,omitempty" name:"LatestSyncResult"`

	// 最新一次同步任务的结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestSyncTime *string `json:"LatestSyncTime,omitnil,omitempty" name:"LatestSyncTime"`
}

type DescribeAccountGroupsPageResp struct {
	// 账户分响应对象集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*DescribeAccountGroupsData `json:"Items,omitnil,omitempty" name:"Items"`

	// 分页公共对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Page *Paging `json:"Page,omitnil,omitempty" name:"Page"`
}

// Predefined struct for user
type DescribeAccountGroupsRequestParams struct {
	// 搜索范围,0-仅搜直接子组,1-深层搜索(只支持32位)
	Deepin *int64 `json:"Deepin,omitnil,omitempty" name:"Deepin"`

	// 滤条件、分页参数
	// <li>Name - String - 是否必填：否 - 操作符: like  - 排序支持：否- 按账号分组过滤。</li>
	// 排序条件
	// <li>Itime - string - 是否必填：否 - 排序支持：是 - 按账号分组创建时间排序。</li>
	// <li>Utime - string - 是否必填：否 - 排序支持：是 - 按账号分组更新时间排序。</li>
	Condition *Condition `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 父分组id
	ParentId *int64 `json:"ParentId,omitnil,omitempty" name:"ParentId"`
}

type DescribeAccountGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 搜索范围,0-仅搜直接子组,1-深层搜索(只支持32位)
	Deepin *int64 `json:"Deepin,omitnil,omitempty" name:"Deepin"`

	// 滤条件、分页参数
	// <li>Name - String - 是否必填：否 - 操作符: like  - 排序支持：否- 按账号分组过滤。</li>
	// 排序条件
	// <li>Itime - string - 是否必填：否 - 排序支持：是 - 按账号分组创建时间排序。</li>
	// <li>Utime - string - 是否必填：否 - 排序支持：是 - 按账号分组更新时间排序。</li>
	Condition *Condition `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 父分组id
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
	// 账户分组详情响应数据
	// 注意：此字段可能返回 null，表示取不到有效值。
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	Paging *Paging `json:"Paging,omitnil,omitempty" name:"Paging"`

	// 业务响应数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*DeviceDetail `json:"Items,omitnil,omitempty" name:"Items"`
}

// Predefined struct for user
type DescribeDevicesRequestParams struct {
	// 过滤条件<br>
	// <li>Ip - String - 是否必填：否 - 操作符: eq  - 排序支持：否- 按照Ip进行过滤。</li>
	// <li>MacAddr - String - 是否必填：否 - 操作符: eq  - 排序支持：否- 按照mac地址进行过滤。</li>
	// <li>IoaUserName - String - 是否必填：否 - 操作符: eq  - 排序支持：否- 按照ioa用户名进行过滤。</li>
	// 分页参数<br>
	// <li>PageNum 从1开始，小于等于0时使用默认参数。</li>
	// <li>PageSize 最大值5000，最好不超过100。</li>
	Condition *Condition `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 【和GroupIds必须有一个填写】设备分组id（需要和OsType匹配）
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
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 【必填】操作系统类型（0: win，1：linux，2: mac，3: win_srv，4：android，5：ios   默认值0），需要和GroupId或者GroupIds匹配
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

	// 授权状态 4未授权 5已授权
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type DescribeDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件<br>
	// <li>Ip - String - 是否必填：否 - 操作符: eq  - 排序支持：否- 按照Ip进行过滤。</li>
	// <li>MacAddr - String - 是否必填：否 - 操作符: eq  - 排序支持：否- 按照mac地址进行过滤。</li>
	// <li>IoaUserName - String - 是否必填：否 - 操作符: eq  - 排序支持：否- 按照ioa用户名进行过滤。</li>
	// 分页参数<br>
	// <li>PageNum 从1开始，小于等于0时使用默认参数。</li>
	// <li>PageSize 最大值5000，最好不超过100。</li>
	Condition *Condition `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 【和GroupIds必须有一个填写】设备分组id（需要和OsType匹配）
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
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 【必填】操作系统类型（0: win，1：linux，2: mac，3: win_srv，4：android，5：ios   默认值0），需要和GroupId或者GroupIds匹配
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

	// 授权状态 4未授权 5已授权
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDevicesResponseParams struct {
	// 分页的data数据
	// 注意：此字段可能返回 null，表示取不到有效值。
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountGroupId *int64 `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`
}

type DescribeLocalAccountsData struct {
	// uid，数据库中唯一
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 账号，登录账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 账号id，同Id字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountId *int64 `json:"AccountId,omitnil,omitempty" name:"AccountId"`

	// 账号所在的分组id
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 账号所在的分组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 账号所在的分组名称路径，用英文.分割
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamePath *string `json:"NamePath,omitnil,omitempty" name:"NamePath"`

	// 账号来源,0表示本地账号(只支持32位)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *int64 `json:"Source,omitnil,omitempty" name:"Source"`

	// 账号状态,0禁用，1启用(只支持32位)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 账号的创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Itime *string `json:"Itime,omitnil,omitempty" name:"Itime"`

	// 账号的最后更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Utime *string `json:"Utime,omitnil,omitempty" name:"Utime"`

	// 账号的扩展信息，包含邮箱、手机号、身份证、职位等信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraInfo *string `json:"ExtraInfo,omitnil,omitempty" name:"ExtraInfo"`

	// 用户风险等级，枚举：none, low, middle, high
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskLevel *string `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 所属组
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountGroups []*DescribeLocalAccountAccountGroupsData `json:"AccountGroups,omitnil,omitempty" name:"AccountGroups"`

	// 绑定手机端设备数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MobileBindNum *int64 `json:"MobileBindNum,omitnil,omitempty" name:"MobileBindNum"`

	// 绑定Pc端设备数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PcBindNum *int64 `json:"PcBindNum,omitnil,omitempty" name:"PcBindNum"`

	// 账号在线状态 1：在线 2：离线
	// 注意：此字段可能返回 null，表示取不到有效值。
	OnlineStatus *int64 `json:"OnlineStatus,omitnil,omitempty" name:"OnlineStatus"`

	// 账号活跃状态 1：活跃 2：非活跃
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActiveStatus *int64 `json:"ActiveStatus,omitnil,omitempty" name:"ActiveStatus"`

	// 账号登录时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoginTime *string `json:"LoginTime,omitnil,omitempty" name:"LoginTime"`

	// 账号登出时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogoutTime *string `json:"LogoutTime,omitnil,omitempty" name:"LogoutTime"`
}

type DescribeLocalAccountsPage struct {
	// 公共分页对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Page *Paging `json:"Page,omitnil,omitempty" name:"Page"`

	// 获取账号列表响应的单个对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*DescribeLocalAccountsData `json:"Items,omitnil,omitempty" name:"Items"`
}

// Predefined struct for user
type DescribeLocalAccountsRequestParams struct {
	// 滤条件、分页参数
	// <li>UserName - String - 是否必填：否 - 操作符: eq,like  - 排序支持：否- 按账号UserName过滤。</li>
	// <li>UserId - string - 是否必填：否 - 操作符: eq,like  - 排序支持：否 - 按账号UserNd过滤。</li>
	// <li>Phone - string - 是否必填：否 - 操作符: eq,like - 排序支持：否 - 按手机号过滤。</li>
	Condition *Condition `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 获取账号的分组Id，不传默认获取全部(只支持32位)
	AccountGroupId *int64 `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// 是否仅展示当前目录下用户 1： 递归显示 2：仅显示当前目录下用户(只支持32位)
	ShowFlag *int64 `json:"ShowFlag,omitnil,omitempty" name:"ShowFlag"`
}

type DescribeLocalAccountsRequest struct {
	*tchttp.BaseRequest
	
	// 滤条件、分页参数
	// <li>UserName - String - 是否必填：否 - 操作符: eq,like  - 排序支持：否- 按账号UserName过滤。</li>
	// <li>UserId - string - 是否必填：否 - 操作符: eq,like  - 排序支持：否 - 按账号UserNd过滤。</li>
	// <li>Phone - string - 是否必填：否 - 操作符: eq,like - 排序支持：否 - 按手机号过滤。</li>
	Condition *Condition `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 获取账号的分组Id，不传默认获取全部(只支持32位)
	AccountGroupId *int64 `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// 是否仅展示当前目录下用户 1： 递归显示 2：仅显示当前目录下用户(只支持32位)
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
	// 注意：此字段可能返回 null，表示取不到有效值。
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
	// 账户分组详情响应数据
	// 注意：此字段可能返回 null，表示取不到有效值。
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 设备唯一标识码，在ioa中每个设备有唯一标识码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mid *string `json:"Mid,omitnil,omitempty" name:"Mid"`

	// 终端名（设备名）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 设备所在分组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// OS平台，0：Windows 、1： Linux、 2：macOS 、4： Android、 5: iOS。默认是0
	// 注意：此字段可能返回 null，表示取不到有效值。
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`

	// 设备IP地址（出口IP）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 在线状态，2：在线、0或者1:离线
	// 注意：此字段可能返回 null，表示取不到有效值。
	OnlineStatus *int64 `json:"OnlineStatus,omitnil,omitempty" name:"OnlineStatus"`

	// 客户端版本号-大整数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 客户端版本号-点分字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	StrVersion *string `json:"StrVersion,omitnil,omitempty" name:"StrVersion"`

	// 首次在线时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Itime *string `json:"Itime,omitnil,omitempty" name:"Itime"`

	// 最后一次在线时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConnActiveTime *string `json:"ConnActiveTime,omitnil,omitempty" name:"ConnActiveTime"`

	// 设备是否加锁 ，1：锁定 0或者2：未锁定。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Locked *int64 `json:"Locked,omitnil,omitempty" name:"Locked"`

	// 设备本地IP列表, 包括IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocalIpList *string `json:"LocalIpList,omitnil,omitempty" name:"LocalIpList"`

	// 主机ID(只支持32位)
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostId *int64 `json:"HostId,omitnil,omitempty" name:"HostId"`

	// 设备所属分组名
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 设备所属分组路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupNamePath *string `json:"GroupNamePath,omitnil,omitempty" name:"GroupNamePath"`

	// 未修复高危漏洞数(只支持32位)
	// 注意：此字段可能返回 null，表示取不到有效值。
	CriticalVulListCount *int64 `json:"CriticalVulListCount,omitnil,omitempty" name:"CriticalVulListCount"`

	// 设备名，和Name相同
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComputerName *string `json:"ComputerName,omitnil,omitempty" name:"ComputerName"`

	// 登录域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// MAC地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	MacAddr *string `json:"MacAddr,omitnil,omitempty" name:"MacAddr"`

	// 漏洞数
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulCount *int64 `json:"VulCount,omitnil,omitempty" name:"VulCount"`

	// 病毒风险数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskCount *int64 `json:"RiskCount,omitnil,omitempty" name:"RiskCount"`

	// 病毒库版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	VirusVer *string `json:"VirusVer,omitnil,omitempty" name:"VirusVer"`

	// 漏洞库版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulVersion *string `json:"VulVersion,omitnil,omitempty" name:"VulVersion"`

	// 系统修复引擎版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	SysRepVersion *string `json:"SysRepVersion,omitnil,omitempty" name:"SysRepVersion"`

	// 高危补丁列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulCriticalList []*string `json:"VulCriticalList,omitnil,omitempty" name:"VulCriticalList"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags *string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 终端用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 防火墙状态，不等于0表示开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirewallStatus *int64 `json:"FirewallStatus,omitnil,omitempty" name:"FirewallStatus"`

	// SN序列号
	// 注意：此字段可能返回 null，表示取不到有效值。
	SerialNum *string `json:"SerialNum,omitnil,omitempty" name:"SerialNum"`

	// 设备管控策略版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceStrategyVer *string `json:"DeviceStrategyVer,omitnil,omitempty" name:"DeviceStrategyVer"`

	// NGN策略版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	NGNStrategyVer *string `json:"NGNStrategyVer,omitnil,omitempty" name:"NGNStrategyVer"`

	// 最近登录账户的账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	IOAUserName *string `json:"IOAUserName,omitnil,omitempty" name:"IOAUserName"`

	// 设备管控新策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceNewStrategyVer *string `json:"DeviceNewStrategyVer,omitnil,omitempty" name:"DeviceNewStrategyVer"`

	// NGN策略新版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	NGNNewStrategyVer *string `json:"NGNNewStrategyVer,omitnil,omitempty" name:"NGNNewStrategyVer"`

	// 主机名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostName *string `json:"HostName,omitnil,omitempty" name:"HostName"`

	// 主板序列号
	// 注意：此字段可能返回 null，表示取不到有效值。
	BaseBoardSn *string `json:"BaseBoardSn,omitnil,omitempty" name:"BaseBoardSn"`

	// 绑定账户只有名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountUsers *string `json:"AccountUsers,omitnil,omitempty" name:"AccountUsers"`

	// 身份策略版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentityStrategyVer *string `json:"IdentityStrategyVer,omitnil,omitempty" name:"IdentityStrategyVer"`

	// 身份策略新版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentityNewStrategyVer *string `json:"IdentityNewStrategyVer,omitnil,omitempty" name:"IdentityNewStrategyVer"`

	// 最近登录账号部门
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountGroupName *string `json:"AccountGroupName,omitnil,omitempty" name:"AccountGroupName"`

	// 最近登录账户的姓名
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// 账号组id
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountGroupId *int64 `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`
}

type Filter struct {
	// 过滤字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	Field *string `json:"Field,omitnil,omitempty" name:"Field"`

	// 过滤方式 eq:等于,net:不等于,like,nlike,gt:大于,lt:小于,egt:大于等于,elt:小于等于
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 过滤条件
	// 注意：此字段可能返回 null，表示取不到有效值。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type FilterGroup struct {
	// Filters 条件过滤
	// 注意：此字段可能返回 null，表示取不到有效值。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type GetAccountGroupData struct {
	// 分组Namepath
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamePath *string `json:"NamePath,omitnil,omitempty" name:"NamePath"`

	// 分组Id path arr(只支持32位)
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdPathArr []*int64 `json:"IdPathArr,omitnil,omitempty" name:"IdPathArr"`

	// 分组扩展信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraInfo *string `json:"ExtraInfo,omitnil,omitempty" name:"ExtraInfo"`

	// 最后更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Utime *string `json:"Utime,omitnil,omitempty" name:"Utime"`

	// 父分组id(只支持32位)
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentId *uint64 `json:"ParentId,omitnil,omitempty" name:"ParentId"`

	// 组织id
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrgId *string `json:"OrgId,omitnil,omitempty" name:"OrgId"`

	// 分组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分组id(只支持32位)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 分组导入源(只支持32位)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *uint64 `json:"Source,omitnil,omitempty" name:"Source"`

	// Id Path
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdPath *string `json:"IdPath,omitnil,omitempty" name:"IdPath"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Itime *string `json:"Itime,omitnil,omitempty" name:"Itime"`

	// 父组织id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentOrgId *string `json:"ParentOrgId,omitnil,omitempty" name:"ParentOrgId"`

	// 导入信息,json格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Import *string `json:"Import,omitnil,omitempty" name:"Import"`

	// 是否开启导入架构
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImportEnable *bool `json:"ImportEnable,omitnil,omitempty" name:"ImportEnable"`

	// 导入类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImportType *string `json:"ImportType,omitnil,omitempty" name:"ImportType"`

	// miniIAMId，MiniIAM源才有
	// 注意：此字段可能返回 null，表示取不到有效值。
	MiniIamId *string `json:"MiniIamId,omitnil,omitempty" name:"MiniIamId"`
}

type Paging struct {
	// 每页条数(只支持32位)
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页码(只支持32位)
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 总页数(只支持32位)
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageCount *uint64 `json:"PageCount,omitnil,omitempty" name:"PageCount"`

	// 记录总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`
}

type Sort struct {
	// 排序字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	Field *string `json:"Field,omitnil,omitempty" name:"Field"`

	// 排序方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}