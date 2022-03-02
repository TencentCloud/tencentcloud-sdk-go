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

package v20211118

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Account struct {

	// 数据库账号名
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 数据库账号描述
	AccountDescription *string `json:"AccountDescription,omitempty" name:"AccountDescription"`

	// 数据库账号创建时间。按照RFC3339标准表示，并且使用东八区时区时间，格式为：YYYY-MM-DDThh:mm:ss+08:00。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 数据库账号信息更新时间。按照RFC3339标准表示，并且使用东八区时区时间，格式为：YYYY-MM-DDThh:mm:ss+08:00。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type AvailableRecoveryTimeRange struct {

	// 可回档起始时间。按照RFC3339标准表示，并且使用东八区时区时间，格式为：YYYY-MM-DDThh:mm:ss+08:00。
	AvailableBeginTime *string `json:"AvailableBeginTime,omitempty" name:"AvailableBeginTime"`

	// 可回档结束时间。按照RFC3339标准表示，并且使用东八区时区时间，格式为：YYYY-MM-DDThh:mm:ss+08:00。
	AvailableEndTime *string `json:"AvailableEndTime,omitempty" name:"AvailableEndTime"`
}

type Backup struct {

	// 备份集ID，集群内唯一
	BackupId *int64 `json:"BackupId,omitempty" name:"BackupId"`

	// 备份集类型，目前只支持 SNAPSHOT：快照
	BackupType *string `json:"BackupType,omitempty" name:"BackupType"`

	// 备份集产生的方案，目前只支持 AUTO：自动
	BackupMethod *string `json:"BackupMethod,omitempty" name:"BackupMethod"`

	// 备份集对应的数据时间。按照RFC3339标准表示，并且使用东八区时区时间，格式为：YYYY-MM-DDThh:mm:ss+08:00。
	BackupDataTime *string `json:"BackupDataTime,omitempty" name:"BackupDataTime"`

	// 备份集数据大小，单位GiB
	BackupDataSize *int64 `json:"BackupDataSize,omitempty" name:"BackupDataSize"`

	// 备份集对应的任务开始时间。按照RFC3339标准表示，并且使用东八区时区时间，格式为：YYYY-MM-DDThh:mm:ss+08:00。
	BackupTaskStartTime *string `json:"BackupTaskStartTime,omitempty" name:"BackupTaskStartTime"`

	// 备份集对应的任务结束时间。按照RFC3339标准表示，并且使用东八区时区时间，格式为：YYYY-MM-DDThh:mm:ss+08:00。
	BackupTaskEndTime *string `json:"BackupTaskEndTime,omitempty" name:"BackupTaskEndTime"`

	// 备份集对应的任务状态  SUCCESS：成功
	BackupTaskStatus *string `json:"BackupTaskStatus,omitempty" name:"BackupTaskStatus"`
}

type CloneClusterToPointInTimeRequest struct {
	*tchttp.BaseRequest

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 数据库版本，目前仅支持 10.17
	DBVersion *string `json:"DBVersion,omitempty" name:"DBVersion"`

	// CPU核数。取值参考文档【购买指南】
	CPU *uint64 `json:"CPU,omitempty" name:"CPU"`

	// 内存大小，单位GiB。取值参考文档【购买指南】
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 已配置的私有网络中的子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 集群付费模式
	// - PREPAID：预付费，即包年包月
	// - POSTPAID_BY_HOUR：按小时后付费
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 对应的备份数据来源集群ID
	SourceClusterId *string `json:"SourceClusterId,omitempty" name:"SourceClusterId"`

	// 对应的备份数据时间点。按照RFC3339标准表示，并且使用东八区时区时间。格式为：YYYY-MM-DDThh:mm:ss+08:00。
	SourceDataPoint *string `json:"SourceDataPoint,omitempty" name:"SourceDataPoint"`

	// 集群名，1-60个字符，可以包含中文、英文、数字和符号"-"、"_"、"."。不输入此参数时默认与ClusterId保持一致。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 项目Id，默认为0表示默认项目
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 连接数据库时，Endpoint使用的端口。取值范围为[1,65534]，默认值为5432
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 集群下实例数量。取值范围为[1,4]，默认值为1
	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 购买时长，单位：月。取值范围为[1,60]，默认值为1。
	// 只有当PayMode为PREPAID时生效。
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 是否自动续费，0-不 1-是。默认为0，只有当PayMode为PREPAID时生效。
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
}

func (r *CloneClusterToPointInTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloneClusterToPointInTimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "DBVersion")
	delete(f, "CPU")
	delete(f, "Memory")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "PayMode")
	delete(f, "SourceClusterId")
	delete(f, "SourceDataPoint")
	delete(f, "ClusterName")
	delete(f, "ProjectId")
	delete(f, "Port")
	delete(f, "InstanceCount")
	delete(f, "Period")
	delete(f, "AutoRenewFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloneClusterToPointInTimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CloneClusterToPointInTimeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 订单号
		DealNameSet []*string `json:"DealNameSet,omitempty" name:"DealNameSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CloneClusterToPointInTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloneClusterToPointInTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Cluster struct {

	// 集群ID，集群的唯一标识
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名字，不修改时默认和集群ID相同
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 数据库版本
	DBVersion *string `json:"DBVersion,omitempty" name:"DBVersion"`

	// 项目ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 集群状态。目前包括
	//  - creating ：创建中
	//  - running : 运行中
	//  - isolating : 隔离中
	//  - isolated : 已隔离
	//  - recovering : 恢复中
	//  - deleting : 删除中
	//  - deleted : 已删除
	Status *string `json:"Status,omitempty" name:"Status"`

	// 集群状态中文含义
	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

	// 集群创建时间。按照RFC3339标准表示，并且使用东八区时区时间，格式为：YYYY-MM-DDThh:mm:ss+08:00。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 存储当前使用量，单位GiB
	StorageUsed *float64 `json:"StorageUsed,omitempty" name:"StorageUsed"`

	// 存储最大使用量，单位GiB
	StorageLimit *uint64 `json:"StorageLimit,omitempty" name:"StorageLimit"`

	// 付费模式：
	//  - PREPAID : 预付费，即包年包月
	//  - POSTPAID_BY_HOUR : 按小时结算后付费
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 预付费集群到期时间。按照RFC3339标准表示，并且使用东八区时区时间，格式为：YYYY-MM-DDThh:mm:ss+08:00。
	PayPeriodEndTime *string `json:"PayPeriodEndTime,omitempty" name:"PayPeriodEndTime"`

	// 预付费集群自动续费标签
	//  - 0 : 到期不自动续费
	//  - 1 : 到期自动续费
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 数据库字符集
	DBCharset *string `json:"DBCharset,omitempty" name:"DBCharset"`

	// 集群内实例的数量
	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 集群内访问点信息
	EndpointSet []*Endpoint `json:"EndpointSet,omitempty" name:"EndpointSet"`
}

type CreateClusterInstancesRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// CPU核数。取值参考文档【购买指南】
	CPU *uint64 `json:"CPU,omitempty" name:"CPU"`

	// 内存大小，单位GiB。取值参考文档【购买指南】
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// 实例名，1-60个字符，可以包含中文、英文、数字和符号"-"、"_"、"."。不输入此参数时默认与InstanceId一致。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 新建实例的数量，默认为1。单集群下实例数量目前不能超过4个。
	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`
}

func (r *CreateClusterInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "CPU")
	delete(f, "Memory")
	delete(f, "InstanceName")
	delete(f, "InstanceCount")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClusterInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateClusterInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 订单号
		DealNameSet []*string `json:"DealNameSet,omitempty" name:"DealNameSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClusterInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateClusterRequest struct {
	*tchttp.BaseRequest

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 数据库版本，目前仅支持 10.17
	DBVersion *string `json:"DBVersion,omitempty" name:"DBVersion"`

	// 数据库用户密码，必须满足 8-64个字符，至少包含 大写字母、小写字母、数字和符号~!@#$%^&*_-+=`|\(){}[]:;'<>,.?/中的任意三种
	MasterUserPassword *string `json:"MasterUserPassword,omitempty" name:"MasterUserPassword"`

	// CPU核数。取值参考文档【购买指南】
	CPU *uint64 `json:"CPU,omitempty" name:"CPU"`

	// 内存大小，单位GiB。取值参考文档【购买指南】
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 已配置的私有网络中的子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 集群付费模式
	//  - PREPAID：预付费，即包年包月
	//  - POSTPAID_BY_HOUR：按小时后付费
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 集群名，1-60个字符，可以包含中文、英文、数字和符号"-"、"_"、"."。不输入此参数时默认与ClusterId保持一致
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 项目Id，默认为0表示默认项目
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 连接数据库时，Endpoint使用的端口。取值范围为[1,65534]，默认值为5432
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 集群下实例数量。取值范围为[1,4]，默认值为1
	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 购买时长，单位：月。取值范围为[1,60]，默认值为1。
	// 只有当PayMode为PREPAID时生效。
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 是否自动续费，0-不 1-是。默认值为0，只有当PayMode为PREPAID时生效。
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
}

func (r *CreateClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "DBVersion")
	delete(f, "MasterUserPassword")
	delete(f, "CPU")
	delete(f, "Memory")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "PayMode")
	delete(f, "ClusterName")
	delete(f, "ProjectId")
	delete(f, "Port")
	delete(f, "InstanceCount")
	delete(f, "Period")
	delete(f, "AutoRenewFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 订单号
		DealNameSet []*string `json:"DealNameSet,omitempty" name:"DealNameSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterInstancesRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 实例ID列表
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`
}

func (r *DeleteClusterInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteClusterInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClusterInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DeleteClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountsRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总条数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 账号信息列表
		AccountSet []*Account `json:"AccountSet,omitempty" name:"AccountSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterBackupsRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 页码，取值范围为[1,INF)，默认值为1
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 每页个数，取值范围为默认为[1,100]，默认值为20
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeClusterBackupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterBackupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterBackupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterBackupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总条数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 备份列表信息
		BackupSet []*Backup `json:"BackupSet,omitempty" name:"BackupSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterBackupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterBackupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterEndpointsRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterEndpointsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterEndpointsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterEndpointsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterEndpointsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总条数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 接入点列表
		EndpointSet []*Endpoint `json:"EndpointSet,omitempty" name:"EndpointSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterEndpointsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterEndpointsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterInstancesRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 页码，取值范围为[1,INF)，默认值为1
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 每页个数，取值范围为默认为[1,100]，默认值为20
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 目前支持查询条件包括：
	//  - InstanceId : 实例ID
	//  - InstanceName : 实例名
	//  - EndpointId : 接入点ID
	//  - Status : 实例状态
	//  - InstanceType : 实例类型
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 排序字段，可选字段：
	// - CreateTime : 实例创建时间(默认值)
	// - PayPeriodEndTime : 实例过期时间
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，可选字段：
	// - DESC : 降序(默认值)
	// - ASC : 升序
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
}

func (r *DescribeClusterInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Filters")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总条数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 实例列表信息
		InstanceSet []*Instance `json:"InstanceSet,omitempty" name:"InstanceSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterRecoveryTimeRangeRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 期望的回档时间点，传入从集群创建时间到当前时间之间的时间点。按照RFC3339标准表示，并且使用东八区时区时间，格式为：YYYY-MM-DDThh:mm:ss+08:00。
	DataPoint *string `json:"DataPoint,omitempty" name:"DataPoint"`
}

func (r *DescribeClusterRecoveryTimeRangeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterRecoveryTimeRangeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "DataPoint")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterRecoveryTimeRangeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterRecoveryTimeRangeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 可回档时间范围列表
		AvailableRecoveryTimeRangeSet []*AvailableRecoveryTimeRange `json:"AvailableRecoveryTimeRangeSet,omitempty" name:"AvailableRecoveryTimeRangeSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterRecoveryTimeRangeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterRecoveryTimeRangeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClustersRequest struct {
	*tchttp.BaseRequest

	// 页码，取值范围为[1,INF)，默认值为1
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 每页条数，取值范围为默认为[1,100]，默认值为20
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 目前支持查询条件包括：
	//  - ClusterId : 集群ID
	//  - ClusterName : 集群名
	//  - ProjectId : 项目ID
	//  - Status : 集群状态
	//  - PayMode : 付费模式
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 排序字段，可选字段：
	//  - CreateTime : 集群创建时间(默认值)
	//  - PayPeriodEndTime : 集群过期时间
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，可选字段：
	//  - DESC : 降序(默认值)
	//  - ASC : 升序
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
}

func (r *DescribeClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Filters")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClustersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总条数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 集群列表信息
		ClusterSet []*Cluster `json:"ClusterSet,omitempty" name:"ClusterSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourcesByDealNameRequest struct {
	*tchttp.BaseRequest

	// 计费订单id（如果计费还没回调业务发货，可能出现错误码InvalidParameterValue.DealNameNotFound，这种情况需要业务重试DescribeResourcesByDealName接口直到成功）
	DealName *string `json:"DealName,omitempty" name:"DealName"`
}

func (r *DescribeResourcesByDealNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcesByDealNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DealName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourcesByDealNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourcesByDealNameResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 资源ID信息列表
		ResourceIdInfoSet []*ResourceIdInfo `json:"ResourceIdInfoSet,omitempty" name:"ResourceIdInfoSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourcesByDealNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcesByDealNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Endpoint struct {

	// 连接点ID，集群内唯一
	EndpointId *string `json:"EndpointId,omitempty" name:"EndpointId"`

	// 连接点所属的集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 连接点名字，默认和连接点ID一致
	EndpointName *string `json:"EndpointName,omitempty" name:"EndpointName"`

	// 连接点类型
	//  - RW : 读写
	//  - RO : 只读
	EndpointType *string `json:"EndpointType,omitempty" name:"EndpointType"`

	// 私有网络VPC实例ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 私有网络VPC下子网实例ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 私有网络VPC下用于访问数据库的IP
	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`

	// 私有网络VPC下用于访问数据库的端口
	PrivatePort *uint64 `json:"PrivatePort,omitempty" name:"PrivatePort"`

	// 公共网络用户访问数据库的IP
	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`

	// 公共网络用户访问数据库的端口
	WanPort *uint64 `json:"WanPort,omitempty" name:"WanPort"`

	// 公共网络用户访问数据库的域名
	WanDomain *string `json:"WanDomain,omitempty" name:"WanDomain"`
}

type Filter struct {

	// 过滤条件名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 过滤条件值数组
	Values []*string `json:"Values,omitempty" name:"Values"`

	// true:精确匹配(默认值) false:(模糊匹配)
	ExactMatch *bool `json:"ExactMatch,omitempty" name:"ExactMatch"`
}

type Instance struct {

	// 实例ID，集群下唯一
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名字，默认和实例ID一致
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 实例所在的访问点ID
	EndpointId *string `json:"EndpointId,omitempty" name:"EndpointId"`

	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 数据库版本
	DBVersion *string `json:"DBVersion,omitempty" name:"DBVersion"`

	// 实例状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 实例状态中文含义
	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

	// 实例创建时间。按照RFC3339标准表示，并且使用东八区时区时间，格式为：YYYY-MM-DDThh:mm:ss+08:00。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 付费模式：
	// - PREPAID : 预付费
	// - POSTPAID_BY_HOUR : 按小时结算后付费
	// 
	// 同一集群下付费模式需要保持一致。
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 实例到期时间。同一集群下到期时间需要保持一致。按照RFC3339标准表示，并且使用东八区时区时间，格式为：YYYY-MM-DDThh:mm:ss+08:00。
	PayPeriodEndTime *string `json:"PayPeriodEndTime,omitempty" name:"PayPeriodEndTime"`

	// CPU核数
	CPU *uint64 `json:"CPU,omitempty" name:"CPU"`

	// 内存大小，单位GiB
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// 实例类型
	//  - RW：读写实例
	//  - RO：只读实例
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
}

type IsolateClusterInstancesRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 实例ID列表
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`
}

func (r *IsolateClusterInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateClusterInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IsolateClusterInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type IsolateClusterInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *IsolateClusterInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateClusterInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IsolateClusterRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *IsolateClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IsolateClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type IsolateClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *IsolateClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccountDescriptionRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 账号名字
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

	// 账号描述，0-256个字符
	AccountDescription *string `json:"AccountDescription,omitempty" name:"AccountDescription"`
}

func (r *ModifyAccountDescriptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountDescriptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "AccountName")
	delete(f, "AccountDescription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccountDescriptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccountDescriptionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccountDescriptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountDescriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterEndpointWanStatusRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 接入点ID
	EndpointId *string `json:"EndpointId,omitempty" name:"EndpointId"`

	// 取值为： 
	//  - OPEN：开启外网 
	//  - CLOSE：关闭外网
	WanStatus *string `json:"WanStatus,omitempty" name:"WanStatus"`
}

func (r *ModifyClusterEndpointWanStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterEndpointWanStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "EndpointId")
	delete(f, "WanStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterEndpointWanStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterEndpointWanStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyClusterEndpointWanStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterEndpointWanStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterInstancesSpecRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 实例ID列表，目前只支持单个实例修改
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// 修改后的CPU核数。取值参考文档【购买指南】
	CPU *uint64 `json:"CPU,omitempty" name:"CPU"`

	// 修改后的内存大小，单位GiB。取值参考文档【购买指南】
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// 操作时机
	//  - IMMEDIATE：立即执行 
	//  - MAINTAIN_PERIOD：维护窗口期执行
	OperationTiming *string `json:"OperationTiming,omitempty" name:"OperationTiming"`
}

func (r *ModifyClusterInstancesSpecRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterInstancesSpecRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceIdSet")
	delete(f, "CPU")
	delete(f, "Memory")
	delete(f, "OperationTiming")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterInstancesSpecRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterInstancesSpecResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyClusterInstancesSpecResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterInstancesSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterNameRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名，1-60个字符，可以包含中文、英文、数字和符号"-"、"_"、"."
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
}

func (r *ModifyClusterNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ClusterName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterNameResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyClusterNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyClustersAutoRenewFlagRequest struct {
	*tchttp.BaseRequest

	// 集群ID列表
	ClusterIdSet []*string `json:"ClusterIdSet,omitempty" name:"ClusterIdSet"`

	// 是否自动续费，0-不 1-是。默认为0，只有当集群的PayMode为PREPAID时生效。
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
}

func (r *ModifyClustersAutoRenewFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClustersAutoRenewFlagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterIdSet")
	delete(f, "AutoRenewFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClustersAutoRenewFlagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyClustersAutoRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyClustersAutoRenewFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClustersAutoRenewFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecoverClusterInstancesRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 实例ID列表
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// 购买时长，单位：月。取值范围为[1,60]，默认值为1。
	// 只有当PayMode为PREPAID时生效。
	Period *uint64 `json:"Period,omitempty" name:"Period"`
}

func (r *RecoverClusterInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecoverClusterInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceIdSet")
	delete(f, "Period")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecoverClusterInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RecoverClusterInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RecoverClusterInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecoverClusterInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecoverClusterRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 购买时长，单位：月。取值范围为[1,60]，默认值为1。
	// 只有当PayMode为PREPAID时生效。
	Period *uint64 `json:"Period,omitempty" name:"Period"`
}

func (r *RecoverClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecoverClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Period")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecoverClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RecoverClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RecoverClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecoverClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RenewClusterRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 续费时间，单位：月。取值范围为[1,60]，默认值为1。
	Period *uint64 `json:"Period,omitempty" name:"Period"`
}

func (r *RenewClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Period")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RenewClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RenewClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetAccountPasswordRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 账号名字
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

	// 数据库用户密码，必须满足 8-64个字符，至少包含 大写字母、小写字母、数字和符号~!@#$%^&*_-+=`|(){}[]:;'<>,.?/中的任意三种
	AccountPassword *string `json:"AccountPassword,omitempty" name:"AccountPassword"`
}

func (r *ResetAccountPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetAccountPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "AccountName")
	delete(f, "AccountPassword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetAccountPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ResetAccountPasswordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetAccountPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetAccountPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceIdInfo struct {

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 实例ID列表
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`
}

type RestartClusterInstancesRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 实例ID列表，目前只支持单个实例重启
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`
}

func (r *RestartClusterInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartClusterInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartClusterInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RestartClusterInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RestartClusterInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartClusterInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TransformClusterPayModeRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 当前付费模式，目前只支持：POSTPAID_BY_HOUR(按小时后付费)
	CurrentPayMode *string `json:"CurrentPayMode,omitempty" name:"CurrentPayMode"`

	// 目标付费模式，目前只支持：PREPAID(预付费)
	TargetPayMode *string `json:"TargetPayMode,omitempty" name:"TargetPayMode"`

	// 购买时长，单位：月。取值范围为[1,60]，默认值为1。
	Period *uint64 `json:"Period,omitempty" name:"Period"`
}

func (r *TransformClusterPayModeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TransformClusterPayModeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "CurrentPayMode")
	delete(f, "TargetPayMode")
	delete(f, "Period")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TransformClusterPayModeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type TransformClusterPayModeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TransformClusterPayModeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TransformClusterPayModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
