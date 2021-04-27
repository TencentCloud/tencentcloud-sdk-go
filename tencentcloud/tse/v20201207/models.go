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

package v20201207

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type DescribeSREInstanceAccessAddressRequest struct {
	*tchttp.BaseRequest

	// 注册引擎实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeSREInstanceAccessAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSREInstanceAccessAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return errors.New("DescribeSREInstanceAccessAddressRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSREInstanceAccessAddressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 内网访问地址
		IntranetAddress *string `json:"IntranetAddress,omitempty" name:"IntranetAddress"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSREInstanceAccessAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSREInstanceAccessAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSREInstancesRequest struct {
	*tchttp.BaseRequest

	// 请求过滤参数
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 翻页单页查询限制数量[0,1000], 默认值0
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 翻页单页偏移量，默认值0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeSREInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSREInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return errors.New("DescribeSREInstancesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSREInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 实例记录
		Content []*SREInstance `json:"Content,omitempty" name:"Content" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSREInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSREInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// 过滤参数名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 过滤参数值
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type ManageConfigRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 配置中心类型（consul、zookeeper、apollo等）
	Type *string `json:"Type,omitempty" name:"Type"`

	// 请求命名 PUT GET DELETE
	Command *string `json:"Command,omitempty" name:"Command"`

	// 配置的Key
	Key *string `json:"Key,omitempty" name:"Key"`

	// 配置的Value
	Value *string `json:"Value,omitempty" name:"Value"`
}

func (r *ManageConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManageConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Type")
	delete(f, "Command")
	delete(f, "Key")
	delete(f, "Value")
	if len(f) > 0 {
		return errors.New("ManageConfigRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ManageConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 对配置中心操作配置之后的返回值
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ManageConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManageConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SREInstance struct {

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 版本号
	Edition *string `json:"Edition,omitempty" name:"Edition"`

	// 状态, 枚举值:creating/create_fail/running/updating/update_fail/restarting/restart_fail/destroying/destroy_fail
	Status *string `json:"Status,omitempty" name:"Status"`

	// 规格ID
	SpecId *string `json:"SpecId,omitempty" name:"SpecId"`

	// 副本数
	Replica *int64 `json:"Replica,omitempty" name:"Replica"`

	// 类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// Vpc iD
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds" list`

	// 是否开启持久化存储
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableStorage *bool `json:"EnableStorage,omitempty" name:"EnableStorage"`

	// 数据存储方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`

	// 云硬盘容量
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageCapacity *int64 `json:"StorageCapacity,omitempty" name:"StorageCapacity"`

	// 计费方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Paymode *string `json:"Paymode,omitempty" name:"Paymode"`
}
