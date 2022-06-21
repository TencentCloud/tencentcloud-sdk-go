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

package v20180420

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type CdsAuditInstance struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 用户AppId
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// 用户Uin
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 项目ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 续费标识
	RenewFlag *uint64 `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// 所属地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 付费模式（数据安全审计只支持预付费：1）
	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`

	// 实例状态： 0，未生效；1：正常运行； 2：被隔离； 3，已过期
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 实例被隔离时间，格式：yyyy-mm-dd HH:ii:ss
	IsolatedTimestamp *string `json:"IsolatedTimestamp,omitempty" name:"IsolatedTimestamp"`

	// 实例创建时间，格式： yyyy-mm-dd HH:ii:ss
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 实例过期时间，格式：yyyy-mm-dd HH:ii:ss
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例公网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`

	// 实例私网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`

	// 实例类型（版本）
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 实例域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pdomain *string `json:"Pdomain,omitempty" name:"Pdomain"`
}

type DbauditTypesInfo struct {
	// 规格描述
	InstanceVersionName *string `json:"InstanceVersionName,omitempty" name:"InstanceVersionName"`

	// 规格名称
	InstanceVersionKey *string `json:"InstanceVersionKey,omitempty" name:"InstanceVersionKey"`

	// 最大吞吐量
	Qps *uint64 `json:"Qps,omitempty" name:"Qps"`

	// 最大实例数
	MaxInstances *uint64 `json:"MaxInstances,omitempty" name:"MaxInstances"`

	// 入库速率（每小时）
	InsertSpeed *uint64 `json:"InsertSpeed,omitempty" name:"InsertSpeed"`

	// 最大在线存储量，单位：条
	OnlineStorageCapacity *uint64 `json:"OnlineStorageCapacity,omitempty" name:"OnlineStorageCapacity"`

	// 最大归档存储量，单位：条
	ArchivingStorageCapacity *uint64 `json:"ArchivingStorageCapacity,omitempty" name:"ArchivingStorageCapacity"`
}

// Predefined struct for user
type DescribeDasbImageIdsRequestParams struct {

}

type DescribeDasbImageIdsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeDasbImageIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDasbImageIdsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDasbImageIdsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDasbImageIdsResponseParams struct {
	// 基础镜像ID
	BaseImageId *string `json:"BaseImageId,omitempty" name:"BaseImageId"`

	// AI镜像ID
	AiImageId *string `json:"AiImageId,omitempty" name:"AiImageId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDasbImageIdsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDasbImageIdsResponseParams `json:"Response"`
}

func (r *DescribeDasbImageIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDasbImageIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDbauditInstanceTypeRequestParams struct {

}

type DescribeDbauditInstanceTypeRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeDbauditInstanceTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDbauditInstanceTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDbauditInstanceTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDbauditInstanceTypeResponseParams struct {
	// 数据安全审计产品规格信息列表
	DbauditTypesSet []*DbauditTypesInfo `json:"DbauditTypesSet,omitempty" name:"DbauditTypesSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDbauditInstanceTypeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDbauditInstanceTypeResponseParams `json:"Response"`
}

func (r *DescribeDbauditInstanceTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDbauditInstanceTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDbauditInstancesRequestParams struct {
	// 查询条件地域
	SearchRegion *string `json:"SearchRegion,omitempty" name:"SearchRegion"`

	// 限制数目，默认10， 最大50
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认1
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeDbauditInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 查询条件地域
	SearchRegion *string `json:"SearchRegion,omitempty" name:"SearchRegion"`

	// 限制数目，默认10， 最大50
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认1
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeDbauditInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDbauditInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SearchRegion")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDbauditInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDbauditInstancesResponseParams struct {
	// 总实例数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 数据安全审计实例信息列表
	CdsAuditInstanceSet []*CdsAuditInstance `json:"CdsAuditInstanceSet,omitempty" name:"CdsAuditInstanceSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDbauditInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDbauditInstancesResponseParams `json:"Response"`
}

func (r *DescribeDbauditInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDbauditInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDbauditUsedRegionsRequestParams struct {

}

type DescribeDbauditUsedRegionsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeDbauditUsedRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDbauditUsedRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDbauditUsedRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDbauditUsedRegionsResponseParams struct {
	// 可售卖地域信息列表
	RegionSet []*RegionInfo `json:"RegionSet,omitempty" name:"RegionSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDbauditUsedRegionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDbauditUsedRegionsResponseParams `json:"Response"`
}

func (r *DescribeDbauditUsedRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDbauditUsedRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceDbauditInstanceRequestParams struct {
	// 实例规格，取值范围： cdsaudit，cdsaudit_adv， cdsaudit_ent 分别为合规版，高级版，企业版
	InstanceVersion *string `json:"InstanceVersion,omitempty" name:"InstanceVersion"`

	// 询价类型： renew，续费；newbuy，新购
	InquiryType *string `json:"InquiryType,omitempty" name:"InquiryType"`

	// 购买实例的时长。取值范围：1（y/m），2（y/m）,，3（y/m），4（m）， 5（m），6（m）， 7（m），8（m），9（m）， 10（m）
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 购买时长单位，y：年；m：月
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 实例所在地域
	ServiceRegion *string `json:"ServiceRegion,omitempty" name:"ServiceRegion"`
}

type InquiryPriceDbauditInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例规格，取值范围： cdsaudit，cdsaudit_adv， cdsaudit_ent 分别为合规版，高级版，企业版
	InstanceVersion *string `json:"InstanceVersion,omitempty" name:"InstanceVersion"`

	// 询价类型： renew，续费；newbuy，新购
	InquiryType *string `json:"InquiryType,omitempty" name:"InquiryType"`

	// 购买实例的时长。取值范围：1（y/m），2（y/m）,，3（y/m），4（m）， 5（m），6（m）， 7（m），8（m），9（m）， 10（m）
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 购买时长单位，y：年；m：月
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 实例所在地域
	ServiceRegion *string `json:"ServiceRegion,omitempty" name:"ServiceRegion"`
}

func (r *InquiryPriceDbauditInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceDbauditInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceVersion")
	delete(f, "InquiryType")
	delete(f, "TimeSpan")
	delete(f, "TimeUnit")
	delete(f, "ServiceRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceDbauditInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceDbauditInstanceResponseParams struct {
	// 总价，单位：元
	TotalPrice *float64 `json:"TotalPrice,omitempty" name:"TotalPrice"`

	// 真实价钱，预支费用的折扣价，单位：元
	RealTotalCost *float64 `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type InquiryPriceDbauditInstanceResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceDbauditInstanceResponseParams `json:"Response"`
}

func (r *InquiryPriceDbauditInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceDbauditInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDbauditInstancesRenewFlagRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 0，表示默认状态(用户未设置，即初始状态)；1，表示自动续费；2，表示明确不自动续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
}

type ModifyDbauditInstancesRenewFlagRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 0，表示默认状态(用户未设置，即初始状态)；1，表示自动续费；2，表示明确不自动续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
}

func (r *ModifyDbauditInstancesRenewFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDbauditInstancesRenewFlagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AutoRenewFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDbauditInstancesRenewFlagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDbauditInstancesRenewFlagResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDbauditInstancesRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDbauditInstancesRenewFlagResponseParams `json:"Response"`
}

func (r *ModifyDbauditInstancesRenewFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDbauditInstancesRenewFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionInfo struct {
	// 地域ID
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// 地域名称
	Region *string `json:"Region,omitempty" name:"Region"`

	// 地域描述
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 地域可用状态
	RegionState *int64 `json:"RegionState,omitempty" name:"RegionState"`
}