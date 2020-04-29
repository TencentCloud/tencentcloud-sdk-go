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

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

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

func (r *ActivateSubscribeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ActivateSubscribeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 配置数据订阅任务ID。
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ActivateSubscribeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ActivateSubscribeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CompleteMigrateJobRequest struct {
	*tchttp.BaseRequest

	// 数据迁移任务ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *CompleteMigrateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CompleteMigrateJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CompleteMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CompleteMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

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

type CreateMigrateCheckJobRequest struct {
	*tchttp.BaseRequest

	// 数据迁移任务ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *CreateMigrateCheckJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateMigrateCheckJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateMigrateCheckJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMigrateCheckJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateMigrateCheckJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateMigrateJobRequest struct {
	*tchttp.BaseRequest

	// 数据迁移任务名称
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// 迁移任务配置选项
	MigrateOption *MigrateOption `json:"MigrateOption,omitempty" name:"MigrateOption"`

	// 源实例数据库类型，目前支持：mysql，redis，mongodb，postgresql，mariadb，percona。不同地域数据库类型的具体支持情况，请参考控制台创建迁移页面。
	SrcDatabaseType *string `json:"SrcDatabaseType,omitempty" name:"SrcDatabaseType"`

	// 源实例接入类型，值包括：extranet(外网),cvm(CVM自建实例),dcg(专线接入的实例),vpncloud(云VPN接入的实例),cdb(腾讯云数据库实例),ccn(云联网实例)
	SrcAccessType *string `json:"SrcAccessType,omitempty" name:"SrcAccessType"`

	// 源实例信息，具体内容跟迁移任务类型相关
	SrcInfo *SrcInfo `json:"SrcInfo,omitempty" name:"SrcInfo"`

	// 目标实例数据库类型，目前支持：mysql，redis，mongodb，postgresql，mariadb，percona。不同地域数据库类型的具体支持情况，请参考控制台创建迁移页面。
	DstDatabaseType *string `json:"DstDatabaseType,omitempty" name:"DstDatabaseType"`

	// 目标实例接入类型，目前支持：cdb（腾讯云数据库实例）
	DstAccessType *string `json:"DstAccessType,omitempty" name:"DstAccessType"`

	// 目标实例信息
	DstInfo *DstInfo `json:"DstInfo,omitempty" name:"DstInfo"`

	// 需要迁移的源数据库表信息，用json格式的字符串描述。当MigrateOption.MigrateObject配置为2（指定库表迁移）时必填。
	// 对于database-table两级结构的数据库：
	// [{Database:db1,Table:[table1,table2]},{Database:db2}]
	// 对于database-schema-table三级结构：
	// [{Database:db1,Schema:s1
	// Table:[table1,table2]},{Database:db1,Schema:s2
	// Table:[table1,table2]},{Database:db2,Schema:s1
	// Table:[table1,table2]},{Database:db3},{Database:db4
	// Schema:s1}]
	DatabaseInfo *string `json:"DatabaseInfo,omitempty" name:"DatabaseInfo"`
}

func (r *CreateMigrateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateMigrateJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据迁移任务ID
		JobId *string `json:"JobId,omitempty" name:"JobId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateMigrateJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

	// 是否自动续费，默认为0，1表示自动续费。小时计费实例设置该标识无效。
	AutoRenew *int64 `json:"AutoRenew,omitempty" name:"AutoRenew"`
}

func (r *CreateSubscribeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSubscribeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSubscribeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据订阅实例的ID数组
	// 注意：此字段可能返回 null，表示取不到有效值。
		SubscribeIds []*string `json:"SubscribeIds,omitempty" name:"SubscribeIds" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSubscribeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSubscribeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSyncCheckJobRequest struct {
	*tchttp.BaseRequest

	// 灾备同步任务ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *CreateSyncCheckJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSyncCheckJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSyncCheckJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSyncCheckJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSyncCheckJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSyncJobRequest struct {
	*tchttp.BaseRequest

	// 灾备同步任务名
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// 灾备同步任务配置选项
	SyncOption *SyncOption `json:"SyncOption,omitempty" name:"SyncOption"`

	// 源实例数据库类型，目前仅包括：mysql
	SrcDatabaseType *string `json:"SrcDatabaseType,omitempty" name:"SrcDatabaseType"`

	// 源实例接入类型，目前仅包括：cdb(云上cdb实例)
	SrcAccessType *string `json:"SrcAccessType,omitempty" name:"SrcAccessType"`

	// 源实例信息
	SrcInfo *SyncInstanceInfo `json:"SrcInfo,omitempty" name:"SrcInfo"`

	// 目标实例数据库类型，目前仅包括：mysql
	DstDatabaseType *string `json:"DstDatabaseType,omitempty" name:"DstDatabaseType"`

	// 目标实例接入类型，目前仅包括：cdb(云上cdb实例)
	DstAccessType *string `json:"DstAccessType,omitempty" name:"DstAccessType"`

	// 目标实例信息
	DstInfo *SyncInstanceInfo `json:"DstInfo,omitempty" name:"DstInfo"`

	// 需要同步的源数据库表信息，用json格式的字符串描述。
	// 对于database-table两级结构的数据库：
	// [{Database:db1,Table:[table1,table2]},{Database:db2}]
	DatabaseInfo *string `json:"DatabaseInfo,omitempty" name:"DatabaseInfo"`
}

func (r *CreateSyncJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSyncJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSyncJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 灾备同步任务ID
		JobId *string `json:"JobId,omitempty" name:"JobId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSyncJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSyncJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DeleteMigrateJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMigrateJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSyncJobRequest struct {
	*tchttp.BaseRequest

	// 待删除的灾备同步任务ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *DeleteSyncJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSyncJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSyncJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSyncJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSyncJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeAsyncRequestInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAsyncRequestInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务执行结果信息
		Info *string `json:"Info,omitempty" name:"Info"`

		// 任务执行状态，可能的值有：success，failed，running
		Status *string `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAsyncRequestInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAsyncRequestInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeMigrateCheckJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMigrateCheckJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
}

func (r *DescribeMigrateCheckJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMigrateCheckJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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
}

func (r *DescribeMigrateJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMigrateJobsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMigrateJobsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务数目
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 任务详情数组
		JobList []*MigrateJobInfo `json:"JobList,omitempty" name:"JobList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMigrateJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMigrateJobsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionConfRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeRegionConfRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRegionConfRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionConfResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 可售卖地域的数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 可售卖地域详情
		Items []*SubscribeRegionConf `json:"Items,omitempty" name:"Items" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRegionConfResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRegionConfResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeSubscribeConfRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscribeConfResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
		SubscribeObjects []*SubscribeObject `json:"SubscribeObjects,omitempty" name:"SubscribeObjects" list`

		// 修改时间
		ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

		// 地域
		Region *string `json:"Region,omitempty" name:"Region"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubscribeConfResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSubscribeConfResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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
	Status []*string `json:"Status,omitempty" name:"Status" list`

	// 数据订阅实例的配置状态，unconfigure - 未配置， configuring - 配置中，configured - 已配置
	SubsStatus []*string `json:"SubsStatus,omitempty" name:"SubsStatus" list`

	// 返回记录的起始偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 单次返回的记录数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 排序方向，可选的值为"DESC"和"ASC"，默认为"DESC"，按创建时间逆序排序
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

func (r *DescribeSubscribesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSubscribesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscribesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合查询条件的实例总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 数据订阅实例的信息列表
		Items []*SubscribeInfo `json:"Items,omitempty" name:"Items" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubscribesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSubscribesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSyncCheckJobRequest struct {
	*tchttp.BaseRequest

	// 要查询的灾备同步任务ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *DescribeSyncCheckJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSyncCheckJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSyncCheckJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务校验状态： starting(开始中)，running(校验中)，finished(校验完成)
		Status *string `json:"Status,omitempty" name:"Status"`

		// 任务校验结果代码
		ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`

		// 提示信息
		ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`

		// 任务执行步骤描述
		StepInfo []*SyncCheckStepInfo `json:"StepInfo,omitempty" name:"StepInfo" list`

		// 校验标志：0（尚未校验成功） ， 1（校验成功）
		CheckFlag *int64 `json:"CheckFlag,omitempty" name:"CheckFlag"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSyncCheckJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSyncCheckJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSyncJobsRequest struct {
	*tchttp.BaseRequest

	// 灾备同步任务ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 灾备同步任务名
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// 排序字段，可以取值为JobId、Status、JobName、CreateTime
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序方式，升序为ASC，降序为DESC
	OrderSeq *string `json:"OrderSeq,omitempty" name:"OrderSeq"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回实例数量，默认20，有效区间[1,100]
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSyncJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSyncJobsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSyncJobsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务数目
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 任务详情数组
		JobList []*SyncJobInfo `json:"JobList,omitempty" name:"JobList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSyncJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSyncJobsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DstInfo struct {

	// 目标实例ID，如cdb-jd92ijd8
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 目标实例地域，如ap-guangzhou
	Region *string `json:"Region,omitempty" name:"Region"`

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

type IsolateSubscribeRequest struct {
	*tchttp.BaseRequest

	// 订阅实例ID
	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
}

func (r *IsolateSubscribeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *IsolateSubscribeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type IsolateSubscribeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *IsolateSubscribeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

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
	StepInfo []*MigrateStepDetailInfo `json:"StepInfo,omitempty" name:"StepInfo" list`
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
	ErrorInfo []*ErrorInfo `json:"ErrorInfo,omitempty" name:"ErrorInfo" list`
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
}

func (r *ModifyMigrateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMigrateJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMigrateJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *ModifySubscribeConsumeTimeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySubscribeConsumeTimeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySubscribeConsumeTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySubscribeConsumeTimeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *ModifySubscribeNameRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySubscribeNameResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySubscribeNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySubscribeNameResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySubscribeObjectsRequest struct {
	*tchttp.BaseRequest

	// 数据订阅实例的ID
	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`

	// 数据订阅的类型，可选的值有：0 - 全实例订阅；1 - 数据订阅；2 - 结构订阅；3 - 数据订阅+结构订阅
	SubscribeObjectType *int64 `json:"SubscribeObjectType,omitempty" name:"SubscribeObjectType"`

	// 订阅的数据库表信息
	Objects []*SubscribeObject `json:"Objects,omitempty" name:"Objects" list`
}

func (r *ModifySubscribeObjectsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySubscribeObjectsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySubscribeObjectsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务的ID
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySubscribeObjectsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySubscribeObjectsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *ModifySubscribeVipVportRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySubscribeVipVportResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySubscribeVipVportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySubscribeVipVportResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySyncJobRequest struct {
	*tchttp.BaseRequest

	// 待修改的灾备同步任务ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 灾备同步任务名称
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// 灾备同步任务配置选项
	SyncOption *SyncOption `json:"SyncOption,omitempty" name:"SyncOption"`

	// 当选择'指定库表'灾备同步的时候, 需要设置待同步的源数据库表信息,用符合json数组格式的字符串描述, 如下所例。
	// 对于database-table两级结构的数据库：
	// [{"Database":"db1","Table":["table1","table2"]},{"Database":"db2"}]
	DatabaseInfo *string `json:"DatabaseInfo,omitempty" name:"DatabaseInfo"`
}

func (r *ModifySyncJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySyncJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySyncJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySyncJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySyncJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *OfflineIsolatedSubscribeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OfflineIsolatedSubscribeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OfflineIsolatedSubscribeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OfflineIsolatedSubscribeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *ResetSubscribeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetSubscribeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetSubscribeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

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

type StartMigrateJobRequest struct {
	*tchttp.BaseRequest

	// 数据迁移任务ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *StartMigrateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartMigrateJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartMigrateJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartSyncJobRequest struct {
	*tchttp.BaseRequest

	// 灾备同步任务ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *StartSyncJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartSyncJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartSyncJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartSyncJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartSyncJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *StopMigrateJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopMigrateJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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
	TableNames []*string `json:"TableNames,omitempty" name:"TableNames" list`
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

type SwitchDrToMasterRequest struct {
	*tchttp.BaseRequest

	// 灾备实例的信息
	DstInfo *SyncInstanceInfo `json:"DstInfo,omitempty" name:"DstInfo"`

	// 数据库的类型  （如 mysql）
	DatabaseType *string `json:"DatabaseType,omitempty" name:"DatabaseType"`
}

func (r *SwitchDrToMasterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SwitchDrToMasterRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SwitchDrToMasterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 后台异步任务请求id
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchDrToMasterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SwitchDrToMasterResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SyncCheckStepInfo struct {

	// 步骤序列
	StepNo *uint64 `json:"StepNo,omitempty" name:"StepNo"`

	// 步骤展现名称
	StepName *string `json:"StepName,omitempty" name:"StepName"`

	// 步骤执行结果代码
	StepCode *int64 `json:"StepCode,omitempty" name:"StepCode"`

	// 步骤执行结果提示
	StepMessage *string `json:"StepMessage,omitempty" name:"StepMessage"`
}

type SyncDetailInfo struct {

	// 总步骤数
	StepAll *int64 `json:"StepAll,omitempty" name:"StepAll"`

	// 当前步骤
	StepNow *int64 `json:"StepNow,omitempty" name:"StepNow"`

	// 总进度
	Progress *string `json:"Progress,omitempty" name:"Progress"`

	// 当前步骤进度
	CurrentStepProgress *string `json:"CurrentStepProgress,omitempty" name:"CurrentStepProgress"`

	// 主从差距，MB
	MasterSlaveDistance *int64 `json:"MasterSlaveDistance,omitempty" name:"MasterSlaveDistance"`

	// 主从差距，秒
	SecondsBehindMaster *int64 `json:"SecondsBehindMaster,omitempty" name:"SecondsBehindMaster"`

	// 步骤信息
	StepInfo []*SyncStepDetailInfo `json:"StepInfo,omitempty" name:"StepInfo" list`
}

type SyncInstanceInfo struct {

	// 地域英文名，如：ap-guangzhou
	Region *string `json:"Region,omitempty" name:"Region"`

	// 实例短ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type SyncJobInfo struct {

	// 灾备任务id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 灾备任务名
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// 任务同步
	SyncOption *SyncOption `json:"SyncOption,omitempty" name:"SyncOption"`

	// 源接入类型
	SrcAccessType *string `json:"SrcAccessType,omitempty" name:"SrcAccessType"`

	// 源数据类型
	SrcDatabaseType *string `json:"SrcDatabaseType,omitempty" name:"SrcDatabaseType"`

	// 源实例信息
	SrcInfo *SyncInstanceInfo `json:"SrcInfo,omitempty" name:"SrcInfo"`

	// 灾备接入类型
	DstAccessType *string `json:"DstAccessType,omitempty" name:"DstAccessType"`

	// 灾备数据类型
	DstDatabaseType *string `json:"DstDatabaseType,omitempty" name:"DstDatabaseType"`

	// 灾备实例信息
	DstInfo *SyncInstanceInfo `json:"DstInfo,omitempty" name:"DstInfo"`

	// 任务信息
	Detail *SyncDetailInfo `json:"Detail,omitempty" name:"Detail"`

	// 任务状态
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 迁移库表
	DatabaseInfo *string `json:"DatabaseInfo,omitempty" name:"DatabaseInfo"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type SyncOption struct {

	// 同步对象，1-整个实例，2-指定库表
	SyncObject *uint64 `json:"SyncObject,omitempty" name:"SyncObject"`

	// 同步开始设置，1-立即开始
	RunMode *uint64 `json:"RunMode,omitempty" name:"RunMode"`

	// 同步模式， 3-全量且增量同步
	SyncType *uint64 `json:"SyncType,omitempty" name:"SyncType"`

	// 数据一致性检测， 1-无需配置
	ConsistencyType *uint64 `json:"ConsistencyType,omitempty" name:"ConsistencyType"`
}

type SyncStepDetailInfo struct {

	// 步骤编号
	StepNo *uint64 `json:"StepNo,omitempty" name:"StepNo"`

	// 步骤名
	StepName *string `json:"StepName,omitempty" name:"StepName"`

	// 能否中止
	CanStop *int64 `json:"CanStop,omitempty" name:"CanStop"`

	// 步骤号
	StepId *int64 `json:"StepId,omitempty" name:"StepId"`
}
