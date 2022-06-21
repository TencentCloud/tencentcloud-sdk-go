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

package v20180330

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type ActivateSubscribeRequestParams struct {
	// 订阅实例ID。
	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`

	// 数据库实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 数据订阅类型0-全实例订阅，1数据订阅，2结构订阅，3数据订阅与结构订阅
	SubscribeObjectType *int64 `json:"SubscribeObjectType,omitempty" name:"SubscribeObjectType"`

	// 订阅对象
	Objects *SubscribeObject `json:"Objects,omitempty" name:"Objects"`

	// 数据订阅服务所在子网。默认为数据库实例所在的子网内。
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// 订阅服务端口；默认为7507
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
}

type ActivateSubscribeRequest struct {
	*tchttp.BaseRequest
	
	// 订阅实例ID。
	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`

	// 数据库实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 数据订阅类型0-全实例订阅，1数据订阅，2结构订阅，3数据订阅与结构订阅
	SubscribeObjectType *int64 `json:"SubscribeObjectType,omitempty" name:"SubscribeObjectType"`

	// 订阅对象
	Objects *SubscribeObject `json:"Objects,omitempty" name:"Objects"`

	// 数据订阅服务所在子网。默认为数据库实例所在的子网内。
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// 订阅服务端口；默认为7507
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
}

func (r *ActivateSubscribeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ActivateSubscribeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubscribeId")
	delete(f, "InstanceId")
	delete(f, "SubscribeObjectType")
	delete(f, "Objects")
	delete(f, "UniqSubnetId")
	delete(f, "Vport")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ActivateSubscribeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ActivateSubscribeResponseParams struct {
	// 配置数据订阅任务ID。
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ActivateSubscribeResponse struct {
	*tchttp.BaseResponse
	Response *ActivateSubscribeResponseParams `json:"Response"`
}

func (r *ActivateSubscribeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ActivateSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CompleteMigrateJobRequestParams struct {
	// 数据迁移任务ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 完成任务的方式,仅支持旧版MySQL迁移任务。waitForSync-等待主从差距为0才停止,immediately-立即完成，不会等待主从差距一致。默认为waitForSync
	CompleteMode *string `json:"CompleteMode,omitempty" name:"CompleteMode"`
}

type CompleteMigrateJobRequest struct {
	*tchttp.BaseRequest
	
	// 数据迁移任务ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 完成任务的方式,仅支持旧版MySQL迁移任务。waitForSync-等待主从差距为0才停止,immediately-立即完成，不会等待主从差距一致。默认为waitForSync
	CompleteMode *string `json:"CompleteMode,omitempty" name:"CompleteMode"`
}

func (r *CompleteMigrateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CompleteMigrateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "CompleteMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CompleteMigrateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CompleteMigrateJobResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CompleteMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *CompleteMigrateJobResponseParams `json:"Response"`
}

func (r *CompleteMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CompleteMigrateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConsistencyParams struct {
	// 数据内容检测参数。表中选出用来数据对比的行，占表的总行数的百分比。取值范围是整数[1-100]
	SelectRowsPerTable *int64 `json:"SelectRowsPerTable,omitempty" name:"SelectRowsPerTable"`

	// 数据内容检测参数。迁移库表中，要进行数据内容检测的表，占所有表的百分比。取值范围是整数[1-100]
	TablesSelectAll *int64 `json:"TablesSelectAll,omitempty" name:"TablesSelectAll"`

	// 数据数量检测，检测表行数是否一致。迁移库表中，要进行数据数量检测的表，占所有表的百分比。取值范围是整数[1-100]
	TablesSelectCount *int64 `json:"TablesSelectCount,omitempty" name:"TablesSelectCount"`
}

// Predefined struct for user
type CreateMigrateCheckJobRequestParams struct {
	// 数据迁移任务ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

type CreateMigrateCheckJobRequest struct {
	*tchttp.BaseRequest
	
	// 数据迁移任务ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *CreateMigrateCheckJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMigrateCheckJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMigrateCheckJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMigrateCheckJobResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateMigrateCheckJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateMigrateCheckJobResponseParams `json:"Response"`
}

func (r *CreateMigrateCheckJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMigrateCheckJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMigrateJobRequestParams struct {
	// 数据迁移任务名称
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// 迁移任务配置选项
	MigrateOption *MigrateOption `json:"MigrateOption,omitempty" name:"MigrateOption"`

	// 源实例数据库类型，目前支持：mysql，redis，mongodb，postgresql，mariadb，percona，sqlserver 不同地域数据库类型的具体支持情况，请参考控制台创建迁移页面。
	SrcDatabaseType *string `json:"SrcDatabaseType,omitempty" name:"SrcDatabaseType"`

	// 源实例接入类型，值包括：extranet(外网),cvm(CVM自建实例),dcg(专线接入的实例),vpncloud(云VPN接入的实例),cdb(腾讯云数据库实例),ccn(云联网实例)
	SrcAccessType *string `json:"SrcAccessType,omitempty" name:"SrcAccessType"`

	// 源实例信息，具体内容跟迁移任务类型相关
	SrcInfo *SrcInfo `json:"SrcInfo,omitempty" name:"SrcInfo"`

	// 目标实例数据库类型，目前支持：mysql，redis，mongodb，postgresql，mariadb，percona，sqlserver，cynosdbmysql。不同地域数据库类型的具体支持情况，请参考控制台创建迁移页面。
	DstDatabaseType *string `json:"DstDatabaseType,omitempty" name:"DstDatabaseType"`

	// 目标实例接入类型，目前支持：cdb（腾讯云数据库实例）
	DstAccessType *string `json:"DstAccessType,omitempty" name:"DstAccessType"`

	// 目标实例信息
	DstInfo *DstInfo `json:"DstInfo,omitempty" name:"DstInfo"`

	// 需要迁移的源数据库表信息，用json格式的字符串描述。当MigrateOption.MigrateObject配置为2（指定库表迁移）时必填。
	// 对于database-table两级结构的数据库：
	// [{"Database":"db1","Table":["table1","table2"]},{"Database":"db2"}]
	// 对于database-schema-table三级结构：
	// [{"Database":"db1","Schema":"s1","Table":["table1","table2"]},{"Database":"db1","Schema":"s2","Table":["table1","table2"]},{"Database":"db2","Schema":"s1","Table":["table1","table2"]},{"Database":"db3"},{"Database":"db4","Schema":"s1"}]
	DatabaseInfo *string `json:"DatabaseInfo,omitempty" name:"DatabaseInfo"`

	// 迁移实例的tag
	Tags []*TagItem `json:"Tags,omitempty" name:"Tags"`

	// 源实例类型: ""或者"simple":主从节点，"cluster": 集群节点
	SrcNodeType *string `json:"SrcNodeType,omitempty" name:"SrcNodeType"`

	// 源实例信息，具体内容跟迁移任务类型相关
	SrcInfoMulti []*SrcInfo `json:"SrcInfoMulti,omitempty" name:"SrcInfoMulti"`
}

type CreateMigrateJobRequest struct {
	*tchttp.BaseRequest
	
	// 数据迁移任务名称
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// 迁移任务配置选项
	MigrateOption *MigrateOption `json:"MigrateOption,omitempty" name:"MigrateOption"`

	// 源实例数据库类型，目前支持：mysql，redis，mongodb，postgresql，mariadb，percona，sqlserver 不同地域数据库类型的具体支持情况，请参考控制台创建迁移页面。
	SrcDatabaseType *string `json:"SrcDatabaseType,omitempty" name:"SrcDatabaseType"`

	// 源实例接入类型，值包括：extranet(外网),cvm(CVM自建实例),dcg(专线接入的实例),vpncloud(云VPN接入的实例),cdb(腾讯云数据库实例),ccn(云联网实例)
	SrcAccessType *string `json:"SrcAccessType,omitempty" name:"SrcAccessType"`

	// 源实例信息，具体内容跟迁移任务类型相关
	SrcInfo *SrcInfo `json:"SrcInfo,omitempty" name:"SrcInfo"`

	// 目标实例数据库类型，目前支持：mysql，redis，mongodb，postgresql，mariadb，percona，sqlserver，cynosdbmysql。不同地域数据库类型的具体支持情况，请参考控制台创建迁移页面。
	DstDatabaseType *string `json:"DstDatabaseType,omitempty" name:"DstDatabaseType"`

	// 目标实例接入类型，目前支持：cdb（腾讯云数据库实例）
	DstAccessType *string `json:"DstAccessType,omitempty" name:"DstAccessType"`

	// 目标实例信息
	DstInfo *DstInfo `json:"DstInfo,omitempty" name:"DstInfo"`

	// 需要迁移的源数据库表信息，用json格式的字符串描述。当MigrateOption.MigrateObject配置为2（指定库表迁移）时必填。
	// 对于database-table两级结构的数据库：
	// [{"Database":"db1","Table":["table1","table2"]},{"Database":"db2"}]
	// 对于database-schema-table三级结构：
	// [{"Database":"db1","Schema":"s1","Table":["table1","table2"]},{"Database":"db1","Schema":"s2","Table":["table1","table2"]},{"Database":"db2","Schema":"s1","Table":["table1","table2"]},{"Database":"db3"},{"Database":"db4","Schema":"s1"}]
	DatabaseInfo *string `json:"DatabaseInfo,omitempty" name:"DatabaseInfo"`

	// 迁移实例的tag
	Tags []*TagItem `json:"Tags,omitempty" name:"Tags"`

	// 源实例类型: ""或者"simple":主从节点，"cluster": 集群节点
	SrcNodeType *string `json:"SrcNodeType,omitempty" name:"SrcNodeType"`

	// 源实例信息，具体内容跟迁移任务类型相关
	SrcInfoMulti []*SrcInfo `json:"SrcInfoMulti,omitempty" name:"SrcInfoMulti"`
}

func (r *CreateMigrateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMigrateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobName")
	delete(f, "MigrateOption")
	delete(f, "SrcDatabaseType")
	delete(f, "SrcAccessType")
	delete(f, "SrcInfo")
	delete(f, "DstDatabaseType")
	delete(f, "DstAccessType")
	delete(f, "DstInfo")
	delete(f, "DatabaseInfo")
	delete(f, "Tags")
	delete(f, "SrcNodeType")
	delete(f, "SrcInfoMulti")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMigrateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMigrateJobResponseParams struct {
	// 数据迁移任务ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateMigrateJobResponseParams `json:"Response"`
}

func (r *CreateMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMigrateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSubscribeRequestParams struct {
	// 订阅的数据库类型，目前支持的有 mysql
	Product *string `json:"Product,omitempty" name:"Product"`

	// 实例付费类型，1小时计费，0包年包月
	PayType *int64 `json:"PayType,omitempty" name:"PayType"`

	// 购买时长。PayType为0时必填。单位为月，最大支持120
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`

	// 购买数量,默认为1，最大为10
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 是否自动续费，0表示不自动续费，1表示自动续费，默认为0。小时计费实例设置该标识无效。
	AutoRenew *int64 `json:"AutoRenew,omitempty" name:"AutoRenew"`

	// 实例资源标签
	Tags []*TagItem `json:"Tags,omitempty" name:"Tags"`

	// 用户自定义实例名
	Name *string `json:"Name,omitempty" name:"Name"`
}

type CreateSubscribeRequest struct {
	*tchttp.BaseRequest
	
	// 订阅的数据库类型，目前支持的有 mysql
	Product *string `json:"Product,omitempty" name:"Product"`

	// 实例付费类型，1小时计费，0包年包月
	PayType *int64 `json:"PayType,omitempty" name:"PayType"`

	// 购买时长。PayType为0时必填。单位为月，最大支持120
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`

	// 购买数量,默认为1，最大为10
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 是否自动续费，0表示不自动续费，1表示自动续费，默认为0。小时计费实例设置该标识无效。
	AutoRenew *int64 `json:"AutoRenew,omitempty" name:"AutoRenew"`

	// 实例资源标签
	Tags []*TagItem `json:"Tags,omitempty" name:"Tags"`

	// 用户自定义实例名
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *CreateSubscribeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSubscribeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "PayType")
	delete(f, "Duration")
	delete(f, "Count")
	delete(f, "AutoRenew")
	delete(f, "Tags")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSubscribeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSubscribeResponseParams struct {
	// 数据订阅实例的ID数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubscribeIds []*string `json:"SubscribeIds,omitempty" name:"SubscribeIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSubscribeResponse struct {
	*tchttp.BaseResponse
	Response *CreateSubscribeResponseParams `json:"Response"`
}

func (r *CreateSubscribeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMigrateJobRequestParams struct {
	// 数据迁移任务ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

type DeleteMigrateJobRequest struct {
	*tchttp.BaseRequest
	
	// 数据迁移任务ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *DeleteMigrateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMigrateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMigrateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMigrateJobResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMigrateJobResponseParams `json:"Response"`
}

func (r *DeleteMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMigrateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAsyncRequestInfoRequestParams struct {
	// 任务 ID
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`
}

type DescribeAsyncRequestInfoRequest struct {
	*tchttp.BaseRequest
	
	// 任务 ID
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`
}

func (r *DescribeAsyncRequestInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAsyncRequestInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AsyncRequestId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAsyncRequestInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAsyncRequestInfoResponseParams struct {
	// 任务执行结果信息
	Info *string `json:"Info,omitempty" name:"Info"`

	// 任务执行状态，可能的值有：success，failed，running
	Status *string `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAsyncRequestInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAsyncRequestInfoResponseParams `json:"Response"`
}

func (r *DescribeAsyncRequestInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAsyncRequestInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigrateCheckJobRequestParams struct {
	// 数据迁移任务ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

type DescribeMigrateCheckJobRequest struct {
	*tchttp.BaseRequest
	
	// 数据迁移任务ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *DescribeMigrateCheckJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrateCheckJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMigrateCheckJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigrateCheckJobResponseParams struct {
	// 校验任务状态：unavailable(当前不可用), starting(开始中)，running(校验中)，finished(校验完成)
	Status *string `json:"Status,omitempty" name:"Status"`

	// 任务的错误码
	ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`

	// 任务的错误信息
	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`

	// Check任务总进度,如："30"表示30%
	Progress *string `json:"Progress,omitempty" name:"Progress"`

	// 校验是否通过,0-未通过，1-校验通过, 3-未校验
	CheckFlag *int64 `json:"CheckFlag,omitempty" name:"CheckFlag"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeMigrateCheckJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMigrateCheckJobResponseParams `json:"Response"`
}

func (r *DescribeMigrateCheckJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrateCheckJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigrateJobsRequestParams struct {
	// 数据迁移任务ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 数据迁移任务名称
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// 排序字段，可以取值为JobId、Status、JobName、MigrateType、RunMode、CreateTime
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序方式，升序为ASC，降序为DESC
	OrderSeq *string `json:"OrderSeq,omitempty" name:"OrderSeq"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回实例数量，默认20，有效区间[1,100]
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 标签过滤条件
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`
}

type DescribeMigrateJobsRequest struct {
	*tchttp.BaseRequest
	
	// 数据迁移任务ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 数据迁移任务名称
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// 排序字段，可以取值为JobId、Status、JobName、MigrateType、RunMode、CreateTime
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序方式，升序为ASC，降序为DESC
	OrderSeq *string `json:"OrderSeq,omitempty" name:"OrderSeq"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回实例数量，默认20，有效区间[1,100]
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 标签过滤条件
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`
}

func (r *DescribeMigrateJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrateJobsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "JobName")
	delete(f, "Order")
	delete(f, "OrderSeq")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "TagFilters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMigrateJobsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigrateJobsResponseParams struct {
	// 任务数目
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 任务详情数组
	JobList []*MigrateJobInfo `json:"JobList,omitempty" name:"JobList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeMigrateJobsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMigrateJobsResponseParams `json:"Response"`
}

func (r *DescribeMigrateJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrateJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegionConfRequestParams struct {

}

type DescribeRegionConfRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeRegionConfRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionConfRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRegionConfRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegionConfResponseParams struct {
	// 可售卖地域的数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 可售卖地域详情
	Items []*SubscribeRegionConf `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRegionConfResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRegionConfResponseParams `json:"Response"`
}

func (r *DescribeRegionConfResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionConfResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubscribeConfRequestParams struct {
	// 订阅实例ID
	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
}

type DescribeSubscribeConfRequest struct {
	*tchttp.BaseRequest
	
	// 订阅实例ID
	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
}

func (r *DescribeSubscribeConfRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubscribeConfRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubscribeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSubscribeConfRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubscribeConfResponseParams struct {
	// 订阅实例ID
	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`

	// 订阅实例名称
	SubscribeName *string `json:"SubscribeName,omitempty" name:"SubscribeName"`

	// 订阅通道
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 订阅数据库类型
	Product *string `json:"Product,omitempty" name:"Product"`

	// 被订阅的实例
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 被订阅的实例的状态，可能的值有running,offline,isolate
	InstanceStatus *string `json:"InstanceStatus,omitempty" name:"InstanceStatus"`

	// 订阅实例状态，可能的值有unconfigure-未配置，configuring-配置中，configured-已配置
	SubsStatus *string `json:"SubsStatus,omitempty" name:"SubsStatus"`

	// 订阅实例生命周期状态，可能的值有：normal-正常，isolating-隔离中，isolated-已隔离，offlining-下线中
	Status *string `json:"Status,omitempty" name:"Status"`

	// 订阅实例创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 订阅实例被隔离时间
	IsolateTime *string `json:"IsolateTime,omitempty" name:"IsolateTime"`

	// 订阅实例到期时间
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 订阅实例下线时间
	OfflineTime *string `json:"OfflineTime,omitempty" name:"OfflineTime"`

	// 订阅实例消费时间起点。
	ConsumeStartTime *string `json:"ConsumeStartTime,omitempty" name:"ConsumeStartTime"`

	// 订阅实例计费类型，1-小时计费，0-包年包月
	PayType *int64 `json:"PayType,omitempty" name:"PayType"`

	// 订阅通道Vip
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 订阅通道Port
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// 订阅通道所在VpcId
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// 订阅通道所在SubnetId
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// 当前SDK消费时间位点
	SdkConsumedTime *string `json:"SdkConsumedTime,omitempty" name:"SdkConsumedTime"`

	// 订阅SDK IP地址
	SdkHost *string `json:"SdkHost,omitempty" name:"SdkHost"`

	// 订阅对象类型0-全实例订阅，1-DDL数据订阅，2-DML结构订阅，3-DDL数据订阅+DML结构订阅
	SubscribeObjectType *int64 `json:"SubscribeObjectType,omitempty" name:"SubscribeObjectType"`

	// 订阅对象，当SubscribeObjectType 为0时，此字段为空数组
	SubscribeObjects []*SubscribeObject `json:"SubscribeObjects,omitempty" name:"SubscribeObjects"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 订阅实例的标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*TagItem `json:"Tags,omitempty" name:"Tags"`

	// 自动续费标识,0-不自动续费，1-自动续费
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 数据订阅版本。老版订阅填txdts，kafka版填kafka
	SubscribeVersion *string `json:"SubscribeVersion,omitempty" name:"SubscribeVersion"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Errors []*SubsErr `json:"Errors,omitempty" name:"Errors"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSubscribeConfResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSubscribeConfResponseParams `json:"Response"`
}

func (r *DescribeSubscribeConfResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubscribeConfResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubscribesRequestParams struct {
	// 数据订阅的实例ID
	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`

	// 数据订阅的实例名称
	SubscribeName *string `json:"SubscribeName,omitempty" name:"SubscribeName"`

	// 绑定数据库实例的ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 数据订阅实例的通道ID
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 计费模式筛选，可能的值：0-包年包月，1-按量计费
	PayType *string `json:"PayType,omitempty" name:"PayType"`

	// 订阅的数据库产品，如mysql
	Product *string `json:"Product,omitempty" name:"Product"`

	// 数据订阅实例的状态，creating - 创建中，normal - 正常运行，isolating - 隔离中，isolated - 已隔离，offlining - 下线中
	Status []*string `json:"Status,omitempty" name:"Status"`

	// 数据订阅实例的配置状态，unconfigure - 未配置， configuring - 配置中，configured - 已配置
	SubsStatus []*string `json:"SubsStatus,omitempty" name:"SubsStatus"`

	// 返回记录的起始偏移量，默认为0。请输入非负整数
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 单次返回的记录数量，默认20。请输入1到100的整数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 排序方向，可选的值为"DESC"和"ASC"，默认为"DESC"，按创建时间逆序排序
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`

	// 标签过滤条件
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`

	// 订阅实例版本;txdts-旧版数据订阅，kafka-kafka版本数据订阅
	SubscribeVersion *string `json:"SubscribeVersion,omitempty" name:"SubscribeVersion"`
}

type DescribeSubscribesRequest struct {
	*tchttp.BaseRequest
	
	// 数据订阅的实例ID
	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`

	// 数据订阅的实例名称
	SubscribeName *string `json:"SubscribeName,omitempty" name:"SubscribeName"`

	// 绑定数据库实例的ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 数据订阅实例的通道ID
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 计费模式筛选，可能的值：0-包年包月，1-按量计费
	PayType *string `json:"PayType,omitempty" name:"PayType"`

	// 订阅的数据库产品，如mysql
	Product *string `json:"Product,omitempty" name:"Product"`

	// 数据订阅实例的状态，creating - 创建中，normal - 正常运行，isolating - 隔离中，isolated - 已隔离，offlining - 下线中
	Status []*string `json:"Status,omitempty" name:"Status"`

	// 数据订阅实例的配置状态，unconfigure - 未配置， configuring - 配置中，configured - 已配置
	SubsStatus []*string `json:"SubsStatus,omitempty" name:"SubsStatus"`

	// 返回记录的起始偏移量，默认为0。请输入非负整数
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 单次返回的记录数量，默认20。请输入1到100的整数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 排序方向，可选的值为"DESC"和"ASC"，默认为"DESC"，按创建时间逆序排序
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`

	// 标签过滤条件
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`

	// 订阅实例版本;txdts-旧版数据订阅，kafka-kafka版本数据订阅
	SubscribeVersion *string `json:"SubscribeVersion,omitempty" name:"SubscribeVersion"`
}

func (r *DescribeSubscribesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubscribesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubscribeId")
	delete(f, "SubscribeName")
	delete(f, "InstanceId")
	delete(f, "ChannelId")
	delete(f, "PayType")
	delete(f, "Product")
	delete(f, "Status")
	delete(f, "SubsStatus")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderDirection")
	delete(f, "TagFilters")
	delete(f, "SubscribeVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSubscribesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubscribesResponseParams struct {
	// 符合查询条件的实例总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 数据订阅实例的信息列表
	Items []*SubscribeInfo `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSubscribesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSubscribesResponseParams `json:"Response"`
}

func (r *DescribeSubscribesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubscribesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DstInfo struct {
	// 目标实例地域，如ap-guangzhou
	Region *string `json:"Region,omitempty" name:"Region"`

	// 目标实例ID，如cdb-jd92ijd8
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 目标实例vip。已废弃，无需填写
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 目标实例vport。已废弃，无需填写
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 目前只对MySQL有效。当为整实例迁移时，1-只读，0-可读写。
	ReadOnly *int64 `json:"ReadOnly,omitempty" name:"ReadOnly"`

	// 目标数据库账号
	User *string `json:"User,omitempty" name:"User"`

	// 目标数据库密码
	Password *string `json:"Password,omitempty" name:"Password"`
}

type ErrorInfo struct {
	// 具体的报错日志, 包含错误码和错误信息
	ErrorLog *string `json:"ErrorLog,omitempty" name:"ErrorLog"`

	// 报错对应的帮助文档Ur
	HelpDoc *string `json:"HelpDoc,omitempty" name:"HelpDoc"`
}

// Predefined struct for user
type IsolateSubscribeRequestParams struct {
	// 订阅实例ID
	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
}

type IsolateSubscribeRequest struct {
	*tchttp.BaseRequest
	
	// 订阅实例ID
	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
}

func (r *IsolateSubscribeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateSubscribeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubscribeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IsolateSubscribeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateSubscribeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type IsolateSubscribeResponse struct {
	*tchttp.BaseResponse
	Response *IsolateSubscribeResponseParams `json:"Response"`
}

func (r *IsolateSubscribeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MigrateDetailInfo struct {
	// 总步骤数
	StepAll *int64 `json:"StepAll,omitempty" name:"StepAll"`

	// 当前步骤
	StepNow *int64 `json:"StepNow,omitempty" name:"StepNow"`

	// 总进度,如："10"
	Progress *string `json:"Progress,omitempty" name:"Progress"`

	// 当前步骤进度,如:"1"
	CurrentStepProgress *string `json:"CurrentStepProgress,omitempty" name:"CurrentStepProgress"`

	// 主从差距，MB；在增量同步阶段有效，目前支持产品为：redis和mysql
	MasterSlaveDistance *int64 `json:"MasterSlaveDistance,omitempty" name:"MasterSlaveDistance"`

	// 主从差距，秒；在增量同步阶段有效，目前支持产品为：mysql
	SecondsBehindMaster *int64 `json:"SecondsBehindMaster,omitempty" name:"SecondsBehindMaster"`

	// 步骤信息
	StepInfo []*MigrateStepDetailInfo `json:"StepInfo,omitempty" name:"StepInfo"`
}

type MigrateJobInfo struct {
	// 数据迁移任务ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 数据迁移任务名称
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// 迁移任务配置选项
	MigrateOption *MigrateOption `json:"MigrateOption,omitempty" name:"MigrateOption"`

	// 源实例数据库类型:mysql，redis，mongodb，postgresql，mariadb，percona
	SrcDatabaseType *string `json:"SrcDatabaseType,omitempty" name:"SrcDatabaseType"`

	// 源实例接入类型，值包括：extranet(外网),cvm(cvm自建实例),dcg(专线接入的实例),vpncloud(云vpn接入的实例),cdb(腾讯云数据库实例),ccn(云联网实例)
	SrcAccessType *string `json:"SrcAccessType,omitempty" name:"SrcAccessType"`

	// 源实例信息，具体内容跟迁移任务类型相关
	SrcInfo *SrcInfo `json:"SrcInfo,omitempty" name:"SrcInfo"`

	// 目标实例数据库类型:mysql，redis，mongodb，postgresql，mariadb，percona
	DstDatabaseType *string `json:"DstDatabaseType,omitempty" name:"DstDatabaseType"`

	// 目标实例接入类型，目前支持：cdb(腾讯云数据库实例)
	DstAccessType *string `json:"DstAccessType,omitempty" name:"DstAccessType"`

	// 目标实例信息
	DstInfo *DstInfo `json:"DstInfo,omitempty" name:"DstInfo"`

	// 需要迁移的源数据库表信息，如果需要迁移的是整个实例，该字段为[]
	DatabaseInfo *string `json:"DatabaseInfo,omitempty" name:"DatabaseInfo"`

	// 任务创建(提交)时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 任务开始执行时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 任务执行结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 任务状态,取值为：1-创建中(Creating),3-校验中(Checking)4-校验通过(CheckPass),5-校验不通过（CheckNotPass）,7-任务运行(Running),8-准备完成（ReadyComplete）,9-任务成功（Success）,10-任务失败（Failed）,11-撤销中（Stopping）,12-完成中（Completing）
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 任务详情
	Detail *MigrateDetailInfo `json:"Detail,omitempty" name:"Detail"`

	// 任务错误信息提示，当任务发生错误时，不为null或者空值
	ErrorInfo []*ErrorInfo `json:"ErrorInfo,omitempty" name:"ErrorInfo"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*TagItem `json:"Tags,omitempty" name:"Tags"`

	// 源实例为集群时且接入为非cdb时源实例信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SrcInfoMulti []*SrcInfo `json:"SrcInfoMulti,omitempty" name:"SrcInfoMulti"`
}

type MigrateOption struct {
	// 任务运行模式，值包括：1-立即执行，2-定时执行
	RunMode *int64 `json:"RunMode,omitempty" name:"RunMode"`

	// 期望执行时间，当runMode=2时，该字段必填，时间格式：yyyy-mm-dd hh:mm:ss
	ExpectTime *string `json:"ExpectTime,omitempty" name:"ExpectTime"`

	// 数据迁移类型，值包括：1-结构迁移,2-全量迁移,3-全量+增量迁移
	MigrateType *int64 `json:"MigrateType,omitempty" name:"MigrateType"`

	// 迁移对象，1-整个实例，2-指定库表
	MigrateObject *int64 `json:"MigrateObject,omitempty" name:"MigrateObject"`

	// 抽样数据一致性检测参数，1-未配置,2-全量检测,3-抽样检测, 4-仅校验不一致表,5-不检测
	ConsistencyType *int64 `json:"ConsistencyType,omitempty" name:"ConsistencyType"`

	// 是否用源库Root账户覆盖目标库，值包括：0-不覆盖，1-覆盖，选择库表或者结构迁移时应该为0
	IsOverrideRoot *int64 `json:"IsOverrideRoot,omitempty" name:"IsOverrideRoot"`

	// 不同数据库用到的额外参数.以JSON格式描述. 
	// Redis可定义如下的参数: 
	// { 
	// 	"ClientOutputBufferHardLimit":512, 	从机缓冲区的硬性容量限制(MB) 
	// 	"ClientOutputBufferSoftLimit":512, 	从机缓冲区的软性容量限制(MB) 
	// 	"ClientOutputBufferPersistTime":60, 从机缓冲区的软性限制持续时间(秒) 
	// 	"ReplBacklogSize":512, 	环形缓冲区容量限制(MB) 
	// 	"ReplTimeout":120，		复制超时时间(秒) 
	// }
	// MongoDB可定义如下的参数: 
	// {
	// 	'SrcAuthDatabase':'admin', 
	// 	'SrcAuthFlag': "1", 
	// 	'SrcAuthMechanism':"SCRAM-SHA-1"
	// }
	// MySQL暂不支持额外参数设置。
	ExternParams *string `json:"ExternParams,omitempty" name:"ExternParams"`

	// 仅用于“抽样数据一致性检测”，ConsistencyType配置为抽样检测时，必选
	ConsistencyParams *ConsistencyParams `json:"ConsistencyParams,omitempty" name:"ConsistencyParams"`
}

type MigrateStepDetailInfo struct {
	// 步骤序列
	StepNo *int64 `json:"StepNo,omitempty" name:"StepNo"`

	// 步骤展现名称
	StepName *string `json:"StepName,omitempty" name:"StepName"`

	// 步骤英文标识
	StepId *string `json:"StepId,omitempty" name:"StepId"`

	// 步骤状态:0-默认值,1-成功,2-失败,3-执行中,4-未执行
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 当前步骤开始的时间，格式为"yyyy-mm-dd hh:mm:ss"，该字段不存在或者为空是无意义
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

// Predefined struct for user
type ModifyMigrateJobRequestParams struct {
	// 待修改的数据迁移任务ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 数据迁移任务名称
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// 迁移任务配置选项
	MigrateOption *MigrateOption `json:"MigrateOption,omitempty" name:"MigrateOption"`

	// 源实例接入类型，值包括：extranet(外网),cvm(CVM自建实例),dcg(专线接入的实例),vpncloud(云VPN接入的实例),cdb(云上CDB实例)
	SrcAccessType *string `json:"SrcAccessType,omitempty" name:"SrcAccessType"`

	// 源实例信息，具体内容跟迁移任务类型相关
	SrcInfo *SrcInfo `json:"SrcInfo,omitempty" name:"SrcInfo"`

	// 目标实例接入类型，值包括：extranet(外网),cvm(CVM自建实例),dcg(专线接入的实例),vpncloud(云VPN接入的实例)，cdb(云上CDB实例). 目前只支持cdb.
	DstAccessType *string `json:"DstAccessType,omitempty" name:"DstAccessType"`

	// 目标实例信息, 其中目标实例地域不允许修改.
	DstInfo *DstInfo `json:"DstInfo,omitempty" name:"DstInfo"`

	// 当选择'指定库表'迁移的时候, 需要设置待迁移的源数据库表信息,用符合json数组格式的字符串描述, 如下所例。
	// 
	// 对于database-table两级结构的数据库：
	// [{"Database":"db1","Table":["table1","table2"]},{"Database":"db2"}]
	// 对于database-schema-table三级结构：
	// [{"Database":"db1","Schema":"s1","Table":["table1","table2"]},{"Database":"db1","Schema":"s2","Table":["table1","table2"]},{"Database":"db2","Schema":"s1","Table":["table1","table2"]},{"Database":"db3"},{"Database":"db4","Schema":"s1"}]
	// 
	// 如果是'整个实例'的迁移模式,不需设置该字段
	DatabaseInfo *string `json:"DatabaseInfo,omitempty" name:"DatabaseInfo"`

	// 源实例类型: ""或者"simple":主从节点，"cluster": 集群节点
	SrcNodeType *string `json:"SrcNodeType,omitempty" name:"SrcNodeType"`

	// 源实例信息，具体内容跟迁移任务类型相关
	SrcInfoMulti []*SrcInfo `json:"SrcInfoMulti,omitempty" name:"SrcInfoMulti"`
}

type ModifyMigrateJobRequest struct {
	*tchttp.BaseRequest
	
	// 待修改的数据迁移任务ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 数据迁移任务名称
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// 迁移任务配置选项
	MigrateOption *MigrateOption `json:"MigrateOption,omitempty" name:"MigrateOption"`

	// 源实例接入类型，值包括：extranet(外网),cvm(CVM自建实例),dcg(专线接入的实例),vpncloud(云VPN接入的实例),cdb(云上CDB实例)
	SrcAccessType *string `json:"SrcAccessType,omitempty" name:"SrcAccessType"`

	// 源实例信息，具体内容跟迁移任务类型相关
	SrcInfo *SrcInfo `json:"SrcInfo,omitempty" name:"SrcInfo"`

	// 目标实例接入类型，值包括：extranet(外网),cvm(CVM自建实例),dcg(专线接入的实例),vpncloud(云VPN接入的实例)，cdb(云上CDB实例). 目前只支持cdb.
	DstAccessType *string `json:"DstAccessType,omitempty" name:"DstAccessType"`

	// 目标实例信息, 其中目标实例地域不允许修改.
	DstInfo *DstInfo `json:"DstInfo,omitempty" name:"DstInfo"`

	// 当选择'指定库表'迁移的时候, 需要设置待迁移的源数据库表信息,用符合json数组格式的字符串描述, 如下所例。
	// 
	// 对于database-table两级结构的数据库：
	// [{"Database":"db1","Table":["table1","table2"]},{"Database":"db2"}]
	// 对于database-schema-table三级结构：
	// [{"Database":"db1","Schema":"s1","Table":["table1","table2"]},{"Database":"db1","Schema":"s2","Table":["table1","table2"]},{"Database":"db2","Schema":"s1","Table":["table1","table2"]},{"Database":"db3"},{"Database":"db4","Schema":"s1"}]
	// 
	// 如果是'整个实例'的迁移模式,不需设置该字段
	DatabaseInfo *string `json:"DatabaseInfo,omitempty" name:"DatabaseInfo"`

	// 源实例类型: ""或者"simple":主从节点，"cluster": 集群节点
	SrcNodeType *string `json:"SrcNodeType,omitempty" name:"SrcNodeType"`

	// 源实例信息，具体内容跟迁移任务类型相关
	SrcInfoMulti []*SrcInfo `json:"SrcInfoMulti,omitempty" name:"SrcInfoMulti"`
}

func (r *ModifyMigrateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMigrateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "JobName")
	delete(f, "MigrateOption")
	delete(f, "SrcAccessType")
	delete(f, "SrcInfo")
	delete(f, "DstAccessType")
	delete(f, "DstInfo")
	delete(f, "DatabaseInfo")
	delete(f, "SrcNodeType")
	delete(f, "SrcInfoMulti")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMigrateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMigrateJobResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMigrateJobResponseParams `json:"Response"`
}

func (r *ModifyMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMigrateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubscribeAutoRenewFlagRequestParams struct {
	// 订阅实例ID，例如：subs-8uey736k
	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`

	// 自动续费标识。1-自动续费，0-不自动续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
}

type ModifySubscribeAutoRenewFlagRequest struct {
	*tchttp.BaseRequest
	
	// 订阅实例ID，例如：subs-8uey736k
	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`

	// 自动续费标识。1-自动续费，0-不自动续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
}

func (r *ModifySubscribeAutoRenewFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubscribeAutoRenewFlagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubscribeId")
	delete(f, "AutoRenewFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySubscribeAutoRenewFlagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubscribeAutoRenewFlagResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySubscribeAutoRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *ModifySubscribeAutoRenewFlagResponseParams `json:"Response"`
}

func (r *ModifySubscribeAutoRenewFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubscribeAutoRenewFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubscribeConsumeTimeRequestParams struct {
	// 数据订阅实例的ID
	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`

	// 消费时间起点，也即是指定订阅数据的时间起点，时间格式如：Y-m-d h:m:s，取值范围为过去24小时之内
	ConsumeStartTime *string `json:"ConsumeStartTime,omitempty" name:"ConsumeStartTime"`
}

type ModifySubscribeConsumeTimeRequest struct {
	*tchttp.BaseRequest
	
	// 数据订阅实例的ID
	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`

	// 消费时间起点，也即是指定订阅数据的时间起点，时间格式如：Y-m-d h:m:s，取值范围为过去24小时之内
	ConsumeStartTime *string `json:"ConsumeStartTime,omitempty" name:"ConsumeStartTime"`
}

func (r *ModifySubscribeConsumeTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubscribeConsumeTimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubscribeId")
	delete(f, "ConsumeStartTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySubscribeConsumeTimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubscribeConsumeTimeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySubscribeConsumeTimeResponse struct {
	*tchttp.BaseResponse
	Response *ModifySubscribeConsumeTimeResponseParams `json:"Response"`
}

func (r *ModifySubscribeConsumeTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubscribeConsumeTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubscribeNameRequestParams struct {
	// 数据订阅实例的ID
	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`

	// 数据订阅实例的名称，长度限制为[1,60]
	SubscribeName *string `json:"SubscribeName,omitempty" name:"SubscribeName"`
}

type ModifySubscribeNameRequest struct {
	*tchttp.BaseRequest
	
	// 数据订阅实例的ID
	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`

	// 数据订阅实例的名称，长度限制为[1,60]
	SubscribeName *string `json:"SubscribeName,omitempty" name:"SubscribeName"`
}

func (r *ModifySubscribeNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubscribeNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubscribeId")
	delete(f, "SubscribeName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySubscribeNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubscribeNameResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySubscribeNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifySubscribeNameResponseParams `json:"Response"`
}

func (r *ModifySubscribeNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubscribeNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubscribeObjectsRequestParams struct {
	// 数据订阅实例的ID
	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`

	// 数据订阅的类型，可选的值有：0 - 全实例订阅；1 - 数据订阅；2 - 结构订阅；3 - 数据订阅+结构订阅
	SubscribeObjectType *int64 `json:"SubscribeObjectType,omitempty" name:"SubscribeObjectType"`

	// 订阅的数据库表信息
	Objects []*SubscribeObject `json:"Objects,omitempty" name:"Objects"`
}

type ModifySubscribeObjectsRequest struct {
	*tchttp.BaseRequest
	
	// 数据订阅实例的ID
	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`

	// 数据订阅的类型，可选的值有：0 - 全实例订阅；1 - 数据订阅；2 - 结构订阅；3 - 数据订阅+结构订阅
	SubscribeObjectType *int64 `json:"SubscribeObjectType,omitempty" name:"SubscribeObjectType"`

	// 订阅的数据库表信息
	Objects []*SubscribeObject `json:"Objects,omitempty" name:"Objects"`
}

func (r *ModifySubscribeObjectsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubscribeObjectsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubscribeId")
	delete(f, "SubscribeObjectType")
	delete(f, "Objects")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySubscribeObjectsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubscribeObjectsResponseParams struct {
	// 异步任务的ID
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySubscribeObjectsResponse struct {
	*tchttp.BaseResponse
	Response *ModifySubscribeObjectsResponseParams `json:"Response"`
}

func (r *ModifySubscribeObjectsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubscribeObjectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubscribeVipVportRequestParams struct {
	// 数据订阅实例的ID
	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`

	// 指定目的子网，如果传此参数，DstIp必须在目的子网内
	DstUniqSubnetId *string `json:"DstUniqSubnetId,omitempty" name:"DstUniqSubnetId"`

	// 目标IP，与DstPort至少传一个
	DstIp *string `json:"DstIp,omitempty" name:"DstIp"`

	// 目标PORT，支持范围为：[1025-65535]
	DstPort *int64 `json:"DstPort,omitempty" name:"DstPort"`
}

type ModifySubscribeVipVportRequest struct {
	*tchttp.BaseRequest
	
	// 数据订阅实例的ID
	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`

	// 指定目的子网，如果传此参数，DstIp必须在目的子网内
	DstUniqSubnetId *string `json:"DstUniqSubnetId,omitempty" name:"DstUniqSubnetId"`

	// 目标IP，与DstPort至少传一个
	DstIp *string `json:"DstIp,omitempty" name:"DstIp"`

	// 目标PORT，支持范围为：[1025-65535]
	DstPort *int64 `json:"DstPort,omitempty" name:"DstPort"`
}

func (r *ModifySubscribeVipVportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubscribeVipVportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubscribeId")
	delete(f, "DstUniqSubnetId")
	delete(f, "DstIp")
	delete(f, "DstPort")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySubscribeVipVportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubscribeVipVportResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySubscribeVipVportResponse struct {
	*tchttp.BaseResponse
	Response *ModifySubscribeVipVportResponseParams `json:"Response"`
}

func (r *ModifySubscribeVipVportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubscribeVipVportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OfflineIsolatedSubscribeRequestParams struct {
	// 数据订阅实例的ID
	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
}

type OfflineIsolatedSubscribeRequest struct {
	*tchttp.BaseRequest
	
	// 数据订阅实例的ID
	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
}

func (r *OfflineIsolatedSubscribeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OfflineIsolatedSubscribeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubscribeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OfflineIsolatedSubscribeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OfflineIsolatedSubscribeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type OfflineIsolatedSubscribeResponse struct {
	*tchttp.BaseResponse
	Response *OfflineIsolatedSubscribeResponseParams `json:"Response"`
}

func (r *OfflineIsolatedSubscribeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OfflineIsolatedSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetSubscribeRequestParams struct {
	// 数据订阅实例的ID
	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
}

type ResetSubscribeRequest struct {
	*tchttp.BaseRequest
	
	// 数据订阅实例的ID
	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
}

func (r *ResetSubscribeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetSubscribeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubscribeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetSubscribeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetSubscribeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ResetSubscribeResponse struct {
	*tchttp.BaseResponse
	Response *ResetSubscribeResponseParams `json:"Response"`
}

func (r *ResetSubscribeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SrcInfo struct {
	// 阿里云AccessKey。源库是阿里云RDS5.6适用
	AccessKey *string `json:"AccessKey,omitempty" name:"AccessKey"`

	// 实例的IP地址
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 实例的端口
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 实例的用户名
	User *string `json:"User,omitempty" name:"User"`

	// 实例的密码
	Password *string `json:"Password,omitempty" name:"Password"`

	// 阿里云RDS实例ID。源库是阿里云RDS5.6/5.6适用
	RdsInstanceId *string `json:"RdsInstanceId,omitempty" name:"RdsInstanceId"`

	// CVM实例短ID，格式如：ins-olgl39y8，与云服务器控制台页面显示的实例ID相同。如果是CVM自建实例，需要传递此字段
	CvmInstanceId *string `json:"CvmInstanceId,omitempty" name:"CvmInstanceId"`

	// 专线网关ID，格式如：dcg-0rxtqqxb
	UniqDcgId *string `json:"UniqDcgId,omitempty" name:"UniqDcgId"`

	// 私有网络ID，格式如：vpc-92jblxto
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 私有网络下的子网ID，格式如：subnet-3paxmkdz
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// VPN网关ID，格式如：vpngw-9ghexg7q
	UniqVpnGwId *string `json:"UniqVpnGwId,omitempty" name:"UniqVpnGwId"`

	// 数据库实例ID，格式如：cdb-powiqx8q
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 地域英文名，如：ap-guangzhou
	Region *string `json:"Region,omitempty" name:"Region"`

	// 当实例为RDS实例时，填写为aliyun, 其他情况均填写others
	Supplier *string `json:"Supplier,omitempty" name:"Supplier"`

	// 云联网ID，如：ccn-afp6kltc
	// 注意：此字段可能返回 null，表示取不到有效值。
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`

	// 数据库版本，当实例为RDS实例时才有效，格式如：5.6或者5.7，默认为5.6
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`
}

// Predefined struct for user
type StartMigrateJobRequestParams struct {
	// 数据迁移任务ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

type StartMigrateJobRequest struct {
	*tchttp.BaseRequest
	
	// 数据迁移任务ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *StartMigrateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartMigrateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartMigrateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartMigrateJobResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StartMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *StartMigrateJobResponseParams `json:"Response"`
}

func (r *StartMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartMigrateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopMigrateJobRequestParams struct {
	// 数据迁移任务ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

type StopMigrateJobRequest struct {
	*tchttp.BaseRequest
	
	// 数据迁移任务ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *StopMigrateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopMigrateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopMigrateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopMigrateJobResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StopMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *StopMigrateJobResponseParams `json:"Response"`
}

func (r *StopMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopMigrateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubsErr struct {
	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`
}

type SubscribeInfo struct {
	// 数据订阅的实例ID
	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`

	// 数据订阅实例的名称
	SubscribeName *string `json:"SubscribeName,omitempty" name:"SubscribeName"`

	// 数据订阅实例绑定的通道ID
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 数据订阅绑定实例对应的产品名称
	Product *string `json:"Product,omitempty" name:"Product"`

	// 数据订阅实例绑定的数据库实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 数据订阅实例绑定的数据库实例状态
	InstanceStatus *string `json:"InstanceStatus,omitempty" name:"InstanceStatus"`

	// 数据订阅实例的配置状态，unconfigure - 未配置， configuring - 配置中，configured - 已配置
	SubsStatus *string `json:"SubsStatus,omitempty" name:"SubsStatus"`

	// 上次修改时间
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 隔离时间
	IsolateTime *string `json:"IsolateTime,omitempty" name:"IsolateTime"`

	// 到期时间
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 下线时间
	OfflineTime *string `json:"OfflineTime,omitempty" name:"OfflineTime"`

	// 最近一次修改的消费时间起点，如果从未修改则为零值
	ConsumeStartTime *string `json:"ConsumeStartTime,omitempty" name:"ConsumeStartTime"`

	// 数据订阅实例所属地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 计费方式，0 - 包年包月，1 - 按量计费
	PayType *int64 `json:"PayType,omitempty" name:"PayType"`

	// 数据订阅实例的Vip
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 数据订阅实例的Vport
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// 数据订阅实例Vip所在VPC的唯一ID
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// 数据订阅实例Vip所在子网的唯一ID
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// 数据订阅实例的状态，creating - 创建中，normal - 正常运行，isolating - 隔离中，isolated - 已隔离，offlining - 下线中，offline - 已下线
	Status *string `json:"Status,omitempty" name:"Status"`

	// SDK最后一条确认消息的时间戳，如果SDK一直消费，也可以作为SDK当前消费时间点
	SdkConsumedTime *string `json:"SdkConsumedTime,omitempty" name:"SdkConsumedTime"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*TagItem `json:"Tags,omitempty" name:"Tags"`

	// 自动续费标识。0-不自动续费，1-自动续费
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 订阅实例版本；txdts-旧版数据订阅,kafka-kafka版本数据订阅
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubscribeVersion *string `json:"SubscribeVersion,omitempty" name:"SubscribeVersion"`
}

type SubscribeObject struct {
	// 数据订阅对象的类型，0-数据库，1-数据库内的表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjectsType *int64 `json:"ObjectsType,omitempty" name:"ObjectsType"`

	// 订阅数据库的名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 订阅数据库中表名称数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableNames []*string `json:"TableNames,omitempty" name:"TableNames"`
}

type SubscribeRegionConf struct {
	// 地域名称，如广州
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 地区标识，如ap-guangzhou
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 地域名称，如华南地区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Area *string `json:"Area,omitempty" name:"Area"`

	// 是否为默认地域，0 - 不是，1 - 是的
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDefaultRegion *int64 `json:"IsDefaultRegion,omitempty" name:"IsDefaultRegion"`

	// 当前地域的售卖情况，1 - 正常， 2-灰度，3 - 停售
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type TagFilter struct {
	// 标签键值
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue []*string `json:"TagValue,omitempty" name:"TagValue"`
}

type TagItem struct {
	// 标签键值
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}