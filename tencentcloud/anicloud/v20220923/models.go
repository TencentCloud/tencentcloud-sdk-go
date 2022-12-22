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

package v20220923

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type CheckAppidExistRequestParams struct {
	// 业务的appid
	SDKAppid *string `json:"SDKAppid,omitempty" name:"SDKAppid"`

	// sub product code
	Type *string `json:"Type,omitempty" name:"Type"`
}

type CheckAppidExistRequest struct {
	*tchttp.BaseRequest
	
	// 业务的appid
	SDKAppid *string `json:"SDKAppid,omitempty" name:"SDKAppid"`

	// sub product code
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *CheckAppidExistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckAppidExistRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SDKAppid")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckAppidExistRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckAppidExistResponseParams struct {
	// appid是否存在
	Exist *bool `json:"Exist,omitempty" name:"Exist"`

	// 请求是否成功
	HasError *bool `json:"HasError,omitempty" name:"HasError"`

	// 出错消息
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CheckAppidExistResponse struct {
	*tchttp.BaseResponse
	Response *CheckAppidExistResponseParams `json:"Response"`
}

func (r *CheckAppidExistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckAppidExistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GoodsDetail struct {
	// 按照四层接入的产品需要传入产品标签,例如:p_cvm
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 四层定义的子产品标签,例如:sp_cvm_s1
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`

	// 资源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type []*string `json:"Type,omitempty" name:"Type"`

	// 资源数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`
}

// Predefined struct for user
type QueryResourceInfoRequestParams struct {
	// 资源id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type QueryResourceInfoRequest struct {
	*tchttp.BaseRequest
	
	// 资源id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *QueryResourceInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryResourceInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryResourceInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryResourceInfoResponseParams struct {
	// 资源信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resource *Resource `json:"Resource,omitempty" name:"Resource"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryResourceInfoResponse struct {
	*tchttp.BaseResponse
	Response *QueryResourceInfoResponseParams `json:"Response"`
}

func (r *QueryResourceInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryResourceInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryResourceRequestParams struct {
	// 0: sdk 1:material
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 分页起始页
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

type QueryResourceRequest struct {
	*tchttp.BaseRequest
	
	// 0: sdk 1:material
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 分页起始页
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *QueryResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryResourceResponseParams struct {
	// 资源信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resources []*Resource `json:"Resources,omitempty" name:"Resources"`

	// 总量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryResourceResponse struct {
	*tchttp.BaseResponse
	Response *QueryResourceResponseParams `json:"Response"`
}

func (r *QueryResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Resource struct {
	// 资源拥有者
	// 注意：此字段可能返回 null，表示取不到有效值。
	UIN *string `json:"UIN,omitempty" name:"UIN"`

	// 云平台应用ID，一般来说与uin存在一一对应的关系
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 资源id，会展示到通知内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 区域ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 资源状态，1正常，2隔离，3销毁(如果资源已经删除或销毁，不需要返回)，4冻结/封禁
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 资源隔离时间，未隔离传"0000-00-00 00:00:00"，资源由隔离变回正常传"0000-00-00 00:00:00"
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsolatedTimestamp *string `json:"IsolatedTimestamp,omitempty" name:"IsolatedTimestamp"`

	// 资源创建时间，用于更新新购订单的资源开始时间，按时长退费时的资源结束时间取自订单的资源结束时间，
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 0后付费 1预付费 2预留实例
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// 资源名称，账单中资源别名，生命周期通知中的资源名称，不返回则通知中展示为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 购买详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	GoodsDetail *GoodsDetail `json:"GoodsDetail,omitempty" name:"GoodsDetail"`

	// 预付费必填 ，自动续费标记，0表示默认状态(用户未设置，即初始状态即手动续费，用户开通了预付费不停服特权也会进行自动续费)， 1表示自动续费，2表示明确不自动续费(用户设置)，若业务无续费概念或无需自动续费，需要设置为0
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewFlag *int64 `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// （仅预付费）资源到期时间，无到期概念传"0000-00-00 00:00:00"
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 地域ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *int64 `json:"Region,omitempty" name:"Region"`

	// sdk appid
	// 注意：此字段可能返回 null，表示取不到有效值。
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// app名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// app的package名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`

	// 资源链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	URL *string `json:"URL,omitempty" name:"URL"`

	// app的entry
	// 注意：此字段可能返回 null，表示取不到有效值。
	Entry *string `json:"Entry,omitempty" name:"Entry"`

	// 0：sdk 1：素材
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstType *int64 `json:"InstType,omitempty" name:"InstType"`

	// license的秘钥
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitempty" name:"Key"`
}