// Copyright 1999-2018 Tencent Ltd.
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
	JobId *string `json:"JobId" name:"JobId"`
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
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
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
	// 1-100的整数值，select(*)对比时每张表的抽样行数比例
	SelectRowsPerTable *int64 `json:"SelectRowsPerTable" name:"SelectRowsPerTable"`
	// 1-100的整数值，select(*)对比的表的比例
	TablesSelectAll *int64 `json:"TablesSelectAll" name:"TablesSelectAll"`
	// 1-100的整数值，select count(*)对比的表的比例
	TablesSelectCount *int64 `json:"TablesSelectCount" name:"TablesSelectCount"`
}

type CreateMigrateCheckJobRequest struct {
	*tchttp.BaseRequest
	// 数据迁移任务ID
	JobId *string `json:"JobId" name:"JobId"`
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
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
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
	JobName *string `json:"JobName" name:"JobName"`
	// 迁移任务配置选项
	MigrateOption *MigrateOption `json:"MigrateOption" name:"MigrateOption"`
	// 源实例数据库类型:mysql,redis,mongodb
	SrcDatabaseType *string `json:"SrcDatabaseType" name:"SrcDatabaseType"`
	// 源实例接入类型，值包括：extranet(外网),cvm(cvm自建实例),dcg(专线接入的实例),vpncloud(云vpn接入的实例),vpnselfbuild(自建vpn接入的实例)，cdb(云上cdb实例)
	SrcAccessType *string `json:"SrcAccessType" name:"SrcAccessType"`
	// 源实例信息，具体内容跟迁移任务类型相关
	SrcInfo *SrcInfo `json:"SrcInfo" name:"SrcInfo"`
	// 目标实例数据库类型,mysql,redis,mongodb
	DstDatabaseType *string `json:"DstDatabaseType" name:"DstDatabaseType"`
	// 目标实例接入类型，值包括：extranet(外网),cvm(cvm自建实例),dcg(专线接入的实例),vpncloud(云vpn接入的实例),vpnselfbuild(自建vpn接入的实例)，cdb(云上cdb实例). 目前只支持cdb.
	DstAccessType *string `json:"DstAccessType" name:"DstAccessType"`
	// 目标实例信息
	DstInfo *DstInfo `json:"DstInfo" name:"DstInfo"`
	// 需要迁移的源数据库表信息，用json格式的字符串描述。
	// 对于database-table两级结构的数据库：
	// [{Database:db1,Table:[table1,table2]},{Database:db2}]
	// 对于database-schema-table三级结构：
	// [{Database:db1,Schema:s1
	// Table:[table1,table2]},{Database:db1,Schema:s2
	// Table:[table1,table2]},{Database:db2,Schema:s1
	// Table:[table1,table2]},{Database:db3},{Database:db4
	// Schema:s1}]
	DatabaseInfo *string `json:"DatabaseInfo" name:"DatabaseInfo"`
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
		JobId *string `json:"JobId" name:"JobId"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateMigrateJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMigrateJobRequest struct {
	*tchttp.BaseRequest
	// 数据迁移任务ID
	JobId *string `json:"JobId" name:"JobId"`
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
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMigrateJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMigrateCheckJobRequest struct {
	*tchttp.BaseRequest
	// 数据迁移任务ID
	JobId *string `json:"JobId" name:"JobId"`
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
		Status *string `json:"Status" name:"Status"`
		// 任务的错误码
		ErrorCode *int64 `json:"ErrorCode" name:"ErrorCode"`
		// 任务的错误信息
		ErrorMessage *string `json:"ErrorMessage" name:"ErrorMessage"`
		// Check任务总进度,如："30"表示30%
		Progress *string `json:"Progress" name:"Progress"`
		// 校验是否通过,0-未通过，1-校验通过, 3-未校验
		CheckFlag *int64 `json:"CheckFlag" name:"CheckFlag"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
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
	JobId *string `json:"JobId" name:"JobId"`
	// 数据迁移任务名称
	JobName *string `json:"JobName" name:"JobName"`
	// 排序字段，可以取值为JobId、Status、JobName、MigrateType、RunMode、CreateTime
	Order *string `json:"Order" name:"Order"`
	// 排序方式，升序为ASC，降序为DESC
	OrderSeq *string `json:"OrderSeq" name:"OrderSeq"`
	// 偏移量，默认为0
	Offset *uint64 `json:"Offset" name:"Offset"`
	// 返回实例数量，默认20，有效区间[1,100]
	Limit *uint64 `json:"Limit" name:"Limit"`
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
		TotalCount *uint64 `json:"TotalCount" name:"TotalCount"`
		// 任务详情数组
		JobList []*MigrateJobInfo `json:"JobList" name:"JobList" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMigrateJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMigrateJobsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DstInfo struct {
	// 目标实例Id
	InstanceId *string `json:"InstanceId" name:"InstanceId"`
	// 目标实例vip
	Ip *string `json:"Ip" name:"Ip"`
	// 目标实例vport
	Port *int64 `json:"Port" name:"Port"`
	// 目标实例Id
	Region *string `json:"Region" name:"Region"`
	// 只读开关
	ReadOnly *int64 `json:"ReadOnly" name:"ReadOnly"`
}

type MigrateDetailInfo struct {
	// 总步骤数
	StepAll *int64 `json:"StepAll" name:"StepAll"`
	// 当前步骤
	StepNow *int64 `json:"StepNow" name:"StepNow"`
	// 总进度,如：
	Progress *string `json:"Progress" name:"Progress"`
	// 当前步骤进度,如:
	CurrentStepProgress *string `json:"CurrentStepProgress" name:"CurrentStepProgress"`
	// 主从差距，MB
	MasterSlaveDistance *int64 `json:"MasterSlaveDistance" name:"MasterSlaveDistance"`
	// 主从差距，秒
	SecondsBehindMaster *int64 `json:"SecondsBehindMaster" name:"SecondsBehindMaster"`
	// 步骤信息
	StepInfo []*MigrateStepDetailInfo `json:"StepInfo" name:"StepInfo" list`
}

type MigrateJobInfo struct {
	// 数据迁移任务ID
	JobId *string `json:"JobId" name:"JobId"`
	// 数据迁移任务名称
	JobName *string `json:"JobName" name:"JobName"`
	// 迁移任务配置选项
	MigrateOption *MigrateOption `json:"MigrateOption" name:"MigrateOption"`
	// 源实例数据库类型:mysql,redis,percona,mongodb,postgresql,sqlserver,mariadb
	SrcDatabaseType *string `json:"SrcDatabaseType" name:"SrcDatabaseType"`
	// 源实例接入类型，值包括：extranet(外网),cvm(cvm自建实例),dcg(专线接入的实例),vpncloud(云vpn接入的实例),vpnselfbuild(自建vpn接入的实例)，cdb(云上cdb实例)
	SrcAccessType *string `json:"SrcAccessType" name:"SrcAccessType"`
	// 源实例信息，具体内容跟迁移任务类型相关
	SrcInfo *SrcInfo `json:"SrcInfo" name:"SrcInfo"`
	// 目标实例数据库类型,mysql,redis,percona,mongodb,postgresql,sqlserver,mariadb
	DstDatabaseType *string `json:"DstDatabaseType" name:"DstDatabaseType"`
	// 源实例接入类型，值包括：extranet(外网),cvm(cvm自建实例),dcg(专线接入的实例),vpncloud(云vpn接入的实例),vpnselfbuild(自建vpn接入的实例)，cdb(云上cdb实例)
	DstAccessType *string `json:"DstAccessType" name:"DstAccessType"`
	// 目的实例信息
	DstInfo *DstInfo `json:"DstInfo" name:"DstInfo"`
	// 需要迁移的源数据库表信息，如果需要迁移的是整个实例，该字段为[]
	DatabaseInfo *string `json:"DatabaseInfo" name:"DatabaseInfo"`
	// 任务创建(提交)时间
	CreateTime *string `json:"CreateTime" name:"CreateTime"`
	// 任务开始执行时间
	StartTime *string `json:"StartTime" name:"StartTime"`
	// 任务执行结束时间
	EndTime *string `json:"EndTime" name:"EndTime"`
	// 任务状态,取值为：1-创建中(Creating),2-创建完成(Created),3-校验中(Checking)4-校验通过(CheckPass),5-校验不通过（CheckNotPass）,6-准备运行(ReadyRun),7-任务运行(Running),8-准备完成（ReadyComplete）,9-任务成功（Success）,10-任务失败（Failed）,11-中止中（Stoping）,12-完成中（Completing）
	Status *int64 `json:"Status" name:"Status"`
	// 任务详情
	Detail *MigrateDetailInfo `json:"Detail" name:"Detail"`
}

type MigrateOption struct {
	// 任务运行模式，值包括：1-立即执行，2-定时执行
	RunMode *int64 `json:"RunMode" name:"RunMode"`
	// 期望执行时间，当runMode=2时，该字段必填，时间格式：yyyy-mm-dd hh:mm:ss
	ExpectTime *string `json:"ExpectTime" name:"ExpectTime"`
	// 数据迁移类型，值包括：1-结构迁移,2-全量迁移,3-全量+增量迁移
	MigrateType *int64 `json:"MigrateType" name:"MigrateType"`
	// 迁移对象，1-整个实例，2-指定库表
	MigrateObject *int64 `json:"MigrateObject" name:"MigrateObject"`
	// 数据对比类型，1-未配置,2-全量检测,3-抽样检测, 4-仅校验不一致表,5-不检测
	ConsistencyType *int64 `json:"ConsistencyType" name:"ConsistencyType"`
	// 是否用源库Root账户覆盖目标库，值包括：0-不覆盖，1-覆盖，选择库表或者结构迁移时应该为0
	IsOverrideRoot *int64 `json:"IsOverrideRoot" name:"IsOverrideRoot"`
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
	ExternParams *string `json:"ExternParams" name:"ExternParams"`
	// 抽样检验时的抽样参数
	ConsistencyParams *ConsistencyParams `json:"ConsistencyParams" name:"ConsistencyParams"`
}

type MigrateStepDetailInfo struct {
	// 步骤序列
	StepNo *int64 `json:"StepNo" name:"StepNo"`
	// 步骤展现名称
	StepName *string `json:"StepName" name:"StepName"`
	// 步骤英文标识
	StepId *string `json:"StepId" name:"StepId"`
	// 步骤状态:0-默认值,1-成功,2-失败,3-执行中,4-未执行
	Status *int64 `json:"Status" name:"Status"`
}

type ModifyMigrateJobRequest struct {
	*tchttp.BaseRequest
	// 待修改的数据迁移任务ID
	JobId *string `json:"JobId" name:"JobId"`
	// 数据迁移任务名称
	JobName *string `json:"JobName" name:"JobName"`
	// 迁移任务配置选项
	MigrateOption *MigrateOption `json:"MigrateOption" name:"MigrateOption"`
	// 源实例接入类型，值包括：extranet(外网),cvm(cvm自建实例),dcg(专线接入的实例),vpncloud(云vpn接入的实例),vpnselfbuild(自建vpn接入的实例)，cdb(云上cdb实例)
	SrcAccessType *string `json:"SrcAccessType" name:"SrcAccessType"`
	// 源实例信息，具体内容跟迁移任务类型相关
	SrcInfo *SrcInfo `json:"SrcInfo" name:"SrcInfo"`
	// 目标实例接入类型，值包括：extranet(外网),cvm(cvm自建实例),dcg(专线接入的实例),vpncloud(云vpn接入的实例),vpnselfbuild(自建vpn接入的实例)，cdb(云上cdb实例). 目前只支持cdb.
	DstAccessType *string `json:"DstAccessType" name:"DstAccessType"`
	// 目标实例信息, 其中目标实例地域不允许修改.
	DstInfo *DstInfo `json:"DstInfo" name:"DstInfo"`
	// 当选择'指定库表'迁移的时候, 需要设置待迁移的源数据库表信息,用符合json数组格式的字符串描述, 如下所例。
	// 
	// 对于database-table两级结构的数据库：
	// [{"Database":"db1","Table":["table1","table2"]},{"Database":"db2"}]
	// 对于database-schema-table三级结构：
	// [{"Database":"db1","Schema":"s1","Table":["table1","table2"]},{"Database":"db1","Schema":"s2","Table":["table1","table2"]},{"Database":"db2","Schema":"s1","Table":["table1","table2"]},{"Database":"db3"},{"Database":"db4","Schema":"s1"}]
	// 
	// 如果是'整个实例'的迁移模式,不需设置该字段
	DatabaseInfo *string `json:"DatabaseInfo" name:"DatabaseInfo"`
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
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMigrateJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SrcInfo struct {
	// 阿里云AccessKey
	AccessKey *string `json:"AccessKey" name:"AccessKey"`
	// 实例的IP地址
	Ip *string `json:"Ip" name:"Ip"`
	// 实例的端口
	Port *int64 `json:"Port" name:"Port"`
	// 实例的用户名
	User *string `json:"User" name:"User"`
	// 实例的密码
	Password *string `json:"Password" name:"Password"`
	// 阿里云rds实例id
	RdsInstanceId *string `json:"RdsInstanceId" name:"RdsInstanceId"`
	// CVM实例短ID，格式如：ins-olgl89y8，与云主机控制台页面显示的实例ID相同，如果是CVM自建实例或者通过自建VPN接入的公网实例，需要传递此字段
	CvmInstanceId *string `json:"CvmInstanceId" name:"CvmInstanceId"`
	// 专线网关ID
	UniqDcgId *string `json:"UniqDcgId" name:"UniqDcgId"`
	// 私有网络ID，和原来的数字vpcId对应，需要调vpc的接口去转换
	VpcId *string `json:"VpcId" name:"VpcId"`
	// 私有网络下的子网ID, 和原来的数字子网ID对应，需要调用vpc的接口去转换
	SubnetId *string `json:"SubnetId" name:"SubnetId"`
	// 系统分配的VPN网关ID
	UniqVpnGwId *string `json:"UniqVpnGwId" name:"UniqVpnGwId"`
	// 实例短Id
	InstanceId *string `json:"InstanceId" name:"InstanceId"`
	// 地域英文名，如：ap-guangzhou
	Region *string `json:"Region" name:"Region"`
	// 服务提供商，如:aliyun,others
	Supplier *string `json:"Supplier" name:"Supplier"`
}

type StartMigrateJobRequest struct {
	*tchttp.BaseRequest
	// 数据迁移任务ID
	JobId *string `json:"JobId" name:"JobId"`
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
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartMigrateJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopMigrateJobRequest struct {
	*tchttp.BaseRequest
	// 数据迁移任务ID
	JobId *string `json:"JobId" name:"JobId"`
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
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopMigrateJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
