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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSREInstanceAccessAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSREInstanceAccessAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSREInstanceAccessAddressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 内网访问地址
		IntranetAddress *string `json:"IntranetAddress,omitempty" name:"IntranetAddress"`

		// 公网访问地址
		InternetAddress *string `json:"InternetAddress,omitempty" name:"InternetAddress"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSREInstanceAccessAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSREInstanceAccessAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSREInstancesRequest struct {
	*tchttp.BaseRequest

	// 请求过滤参数
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 翻页单页查询限制数量[0,1000], 默认值0
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 翻页单页偏移量，默认值0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询类型
	QueryType *string `json:"QueryType,omitempty" name:"QueryType"`
}

func (r *DescribeSREInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSREInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "QueryType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSREInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSREInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 实例记录
		Content []*SREInstance `json:"Content,omitempty" name:"Content"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSREInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSREInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnvInfo struct {

	// 环境名称
	EnvName *string `json:"EnvName,omitempty" name:"EnvName"`

	// 环境对应的网络信息
	VpcInfos []*VpcInfo `json:"VpcInfos,omitempty" name:"VpcInfos"`

	// 云硬盘容量
	StorageCapacity *int64 `json:"StorageCapacity,omitempty" name:"StorageCapacity"`

	// 运行状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// Admin service 访问地址
	AdminServiceIp *string `json:"AdminServiceIp,omitempty" name:"AdminServiceIp"`

	// Config service访问地址
	ConfigServiceIp *string `json:"ConfigServiceIp,omitempty" name:"ConfigServiceIp"`
}

type Filter struct {

	// 过滤参数名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 过滤参数值
	Values []*string `json:"Values,omitempty" name:"Values"`
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
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`

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

	// EKS集群的ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	EKSClusterID *string `json:"EKSClusterID,omitempty" name:"EKSClusterID"`

	// 集群创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 环境配置信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvInfos []*EnvInfo `json:"EnvInfos,omitempty" name:"EnvInfos"`
}

type VpcInfo struct {

	// Vpc Id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}
