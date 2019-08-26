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

	// 源实例接入类型，值包括：extranet(外网),cvm(cvm自建实例),dcg(专线接入的实例),vpncloud(云vpn接入的实例),cdb(腾讯云数据库实例),ccn(云联网实例)
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

	// 目标实例Id，如cdb-jd92ijd8
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 目标实例地域，如ap-guangzhou
	Region *string `json:"Region,omitempty" name:"Region"`

	// 目标实例vip。已废弃，无需填写
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 目标实例vport。已废弃，无需填写
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 目前只对MySQL有效。当为整实例迁移时，1-只读，0-可读写。
	ReadOnly *int64 `json:"ReadOnly,omitempty" name:"ReadOnly"`
}

type MigrateDetailInfo struct {

	// 总步骤数
	StepAll *int64 `json:"StepAll,omitempty" name:"StepAll"`

	// 当前步骤
	StepNow *int64 `json:"StepNow,omitempty" name:"StepNow"`

	// 总进度,如：
	Progress *string `json:"Progress,omitempty" name:"Progress"`

	// 当前步骤进度,如:
	CurrentStepProgress *string `json:"CurrentStepProgress,omitempty" name:"CurrentStepProgress"`

	// 主从差距，MB
	MasterSlaveDistance *int64 `json:"MasterSlaveDistance,omitempty" name:"MasterSlaveDistance"`

	// 主从差距，秒
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
}

type ModifyMigrateJobRequest struct {
	*tchttp.BaseRequest

	// 待修改的数据迁移任务ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 数据迁移任务名称
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// 迁移任务配置选项
	MigrateOption *MigrateOption `json:"MigrateOption,omitempty" name:"MigrateOption"`

	// 源实例接入类型，值包括：extranet(外网),cvm(cvm自建实例),dcg(专线接入的实例),vpncloud(云vpn接入的实例),vpnselfbuild(自建vpn接入的实例)，cdb(云上cdb实例)
	SrcAccessType *string `json:"SrcAccessType,omitempty" name:"SrcAccessType"`

	// 源实例信息，具体内容跟迁移任务类型相关
	SrcInfo *SrcInfo `json:"SrcInfo,omitempty" name:"SrcInfo"`

	// 目标实例接入类型，值包括：extranet(外网),cvm(cvm自建实例),dcg(专线接入的实例),vpncloud(云vpn接入的实例),vpnselfbuild(自建vpn接入的实例)，cdb(云上cdb实例). 目前只支持cdb.
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

	// 阿里云RDS实例ID。源库是阿里云RDS5.6适用
	RdsInstanceId *string `json:"RdsInstanceId,omitempty" name:"RdsInstanceId"`

	// CVM实例短ID，格式如：ins-olgl39y8，与云主机控制台页面显示的实例ID相同。如果是CVM自建实例，需要传递此字段
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

	// 实例短Id
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
