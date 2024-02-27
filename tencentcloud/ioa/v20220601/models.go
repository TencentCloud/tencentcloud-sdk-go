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

	// 私有化默认分组id-名称-操作系统
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

	// 系统类型（0: win，1：linux，2: mac，3: win_srv，4：android，5：ios   默认值0）
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`

	// 在线状态 2 在线 0，1 离线
	OnlineStatus *int64 `json:"OnlineStatus,omitnil,omitempty" name:"OnlineStatus"`

	// 过滤条件--兼容旧接口,参数同Condition
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段--兼容旧接口,参数同Condition
	Sort *Sort `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 获取第几页--兼容旧接口,参数同Condition(只支持32位)
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 每页获取数--兼容旧接口,参数同Condition(只支持32位)
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

	// 私有化默认分组id-名称-操作系统
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

	// 系统类型（0: win，1：linux，2: mac，3: win_srv，4：android，5：ios   默认值0）
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`

	// 在线状态 2 在线 0，1 离线
	OnlineStatus *int64 `json:"OnlineStatus,omitnil,omitempty" name:"OnlineStatus"`

	// 过滤条件--兼容旧接口,参数同Condition
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段--兼容旧接口,参数同Condition
	Sort *Sort `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 获取第几页--兼容旧接口,参数同Condition(只支持32位)
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 每页获取数--兼容旧接口,参数同Condition(只支持32位)
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

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
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

type DeviceDetail struct {
	// 设备ID(只支持32位)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 设备唯一标识符
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mid *string `json:"Mid,omitnil,omitempty" name:"Mid"`

	// 终端名（设备名）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 设备所在分组ID(只支持32位)
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// OS平台(只支持32位)
	// 注意：此字段可能返回 null，表示取不到有效值。
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`

	// 设备IP地址（出口IP）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 在线状态 2 在线 0，1 离线(只支持32位)
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

	// 设备是否加锁 1 锁定 0 2 非锁定(只支持32位)
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

	// 设备名 和Name相同，保留参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComputerName *string `json:"ComputerName,omitnil,omitempty" name:"ComputerName"`

	// 登录域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// MAC地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	MacAddr *string `json:"MacAddr,omitnil,omitempty" name:"MacAddr"`

	// 漏洞数(只支持32位)
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulCount *int64 `json:"VulCount,omitnil,omitempty" name:"VulCount"`

	// 病毒风险数(只支持32位)
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

	// 防火墙状态(只支持32位)
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

	// 最近登录账号
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

	// 登录账号姓名
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