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
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// FilterGroups 条件过滤组
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilterGroups []*FilterGroup `json:"FilterGroups,omitnil" name:"FilterGroups"`

	// Sort 排序字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sort *Sort `json:"Sort,omitnil" name:"Sort"`

	// PageSize 每页获取数(只支持32位)
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// PageNum 获取第几页(只支持32位)
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNum *int64 `json:"PageNum,omitnil" name:"PageNum"`
}

type DescribeDevicesPageRsp struct {
	// 数据分页信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Paging *Paging `json:"Paging,omitnil" name:"Paging"`
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
	Condition *Condition `json:"Condition,omitnil" name:"Condition"`

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
	GroupId *int64 `json:"GroupId,omitnil" name:"GroupId"`

	// 系统类型（0: win，1：linux，2: mac，3: win_srv，4：android，5：ios   默认值0）
	OsType *int64 `json:"OsType,omitnil" name:"OsType"`

	// 在线状态 2 在线 0，1 离线
	OnlineStatus *int64 `json:"OnlineStatus,omitnil" name:"OnlineStatus"`

	// 过滤条件--兼容旧接口,参数同Condition
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 排序字段--兼容旧接口,参数同Condition
	Sort *Sort `json:"Sort,omitnil" name:"Sort"`

	// 获取第几页--兼容旧接口,参数同Condition(只支持32位)
	PageNum *int64 `json:"PageNum,omitnil" name:"PageNum"`

	// 每页获取数--兼容旧接口,参数同Condition(只支持32位)
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 授权状态 4未授权 5已授权
	Status *int64 `json:"Status,omitnil" name:"Status"`
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
	Condition *Condition `json:"Condition,omitnil" name:"Condition"`

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
	GroupId *int64 `json:"GroupId,omitnil" name:"GroupId"`

	// 系统类型（0: win，1：linux，2: mac，3: win_srv，4：android，5：ios   默认值0）
	OsType *int64 `json:"OsType,omitnil" name:"OsType"`

	// 在线状态 2 在线 0，1 离线
	OnlineStatus *int64 `json:"OnlineStatus,omitnil" name:"OnlineStatus"`

	// 过滤条件--兼容旧接口,参数同Condition
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 排序字段--兼容旧接口,参数同Condition
	Sort *Sort `json:"Sort,omitnil" name:"Sort"`

	// 获取第几页--兼容旧接口,参数同Condition(只支持32位)
	PageNum *int64 `json:"PageNum,omitnil" name:"PageNum"`

	// 每页获取数--兼容旧接口,参数同Condition(只支持32位)
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 授权状态 4未授权 5已授权
	Status *int64 `json:"Status,omitnil" name:"Status"`
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
	Data *DescribeDevicesPageRsp `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

type Filter struct {
	// 过滤字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	Field *string `json:"Field,omitnil" name:"Field"`

	// 过滤方式 eq:等于,net:不等于,like,nlike,gt:大于,lt:小于,egt:大于等于,elt:小于等于
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operator *string `json:"Operator,omitnil" name:"Operator"`

	// 过滤条件
	// 注意：此字段可能返回 null，表示取不到有效值。
	Values []*string `json:"Values,omitnil" name:"Values"`
}

type FilterGroup struct {
	// Filters 条件过滤
	// 注意：此字段可能返回 null，表示取不到有效值。
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type Paging struct {
	// 每页条数(只支持32位)
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *uint64 `json:"PageSize,omitnil" name:"PageSize"`

	// 页码(只支持32位)
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNum *uint64 `json:"PageNum,omitnil" name:"PageNum"`

	// 总页数(只支持32位)
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageCount *uint64 `json:"PageCount,omitnil" name:"PageCount"`

	// 记录总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitnil" name:"Total"`
}

type Sort struct {
	// 排序字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	Field *string `json:"Field,omitnil" name:"Field"`

	// 排序方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Order *string `json:"Order,omitnil" name:"Order"`
}