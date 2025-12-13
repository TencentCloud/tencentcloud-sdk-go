// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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

package v20211108

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type BackupColumnItem struct {
	// 列名。
	ColumnName *string `json:"ColumnName,omitnil,omitempty" name:"ColumnName"`

	// 重命名后的列名。
	NewColumnName *string `json:"NewColumnName,omitnil,omitempty" name:"NewColumnName"`
}

type BackupEndpoint struct {
	// 数据库类型。目前支持的值["mysql", "mariadb", "percona"]。注意，该值必须和备份计划的类型一致。
	DatabaseType *string `json:"DatabaseType,omitnil,omitempty" name:"DatabaseType"`

	// 实例接入类型，支持的值包括：
	// "extranet" - 外网;
	// "cvm" - cvm自建实例;
	// "dcg" - 专线接入;
	// "vpncloud" - 云vpn接入;
	// "cdb" - 腾讯云数据库实例;
	// "ccn" - 云联网接入。
	AccessType *string `json:"AccessType,omitnil,omitempty" name:"AccessType"`

	// 用户名。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 登录密码。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 接入地域。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 服务提供商，支持的值包括["aliyun", "aws", "others"]。
	Supplier *string `json:"Supplier,omitnil,omitempty" name:"Supplier"`

	// 实例 Ip。
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 实例端口号。
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 云数据库实例ID，格式如：cdb-qcloudtest。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// CVM 实例ID，格式如：ins-olgl39y8，与云服务器控制台页面显示的实例ID相同。如果是CVM自建实例，需要填写该字段。
	CvmInstanceId *string `json:"CvmInstanceId,omitnil,omitempty" name:"CvmInstanceId"`

	// 专线网关ID，格式如：dcg-0rxtqqxb。
	UniqDcgId *string `json:"UniqDcgId,omitnil,omitempty" name:"UniqDcgId"`

	// VPN网关ID，格式如：vpngw-9ghexg7q。
	UniqVpnGwId *string `json:"UniqVpnGwId,omitnil,omitempty" name:"UniqVpnGwId"`

	// 私有网络ID，格式如：vpc-92jblxto。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID，格式如：subnet-3paxmkdz。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 云联网ID，如：ccn-afp6kltc。
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// 数据库版本，当实例为 RDS 或 AWS 实例时才有效，格式如：5.6或者5.7，默认为5.6。
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// mariadb三引擎版本。
	DBKernel *string `json:"DBKernel,omitnil,omitempty" name:"DBKernel"`
}

type BackupObject struct {
	// 备份对象类型，可能的取值为:
	// "all" - 整实例;
	// "partial" - 部分对象。
	ObjectMode *string `json:"ObjectMode,omitnil,omitempty" name:"ObjectMode"`

	// 备份对象详情，当 ObjectMode 为 partial, 即选择部分对象备份时，该字段不能为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjectItems []*BackupObjectItem `json:"ObjectItems,omitnil,omitempty" name:"ObjectItems"`
}

type BackupObjectItem struct {
	// 库名。
	DBName *string `json:"DBName,omitnil,omitempty" name:"DBName"`

	// 重命名后的库名。
	NewDBName *string `json:"NewDBName,omitnil,omitempty" name:"NewDBName"`

	// schema 名。
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// 重命名后的 schema 名。
	NewSchemaName *string `json:"NewSchemaName,omitnil,omitempty" name:"NewSchemaName"`

	// 库选择模式，可能的取值为：
	// "all" - 当前对象下的所有对象;
	// "partial" - 当前对象下的部分对象。
	DbMode *string `json:"DbMode,omitnil,omitempty" name:"DbMode"`

	// 表选择模式，可能的取值为:
	// "all" - 当前对象下的所有对象;
	// "partial" - 当前对象下的部分对象。
	TableMode *string `json:"TableMode,omitnil,omitempty" name:"TableMode"`

	// 表对象详情。当 TableMode 为 partial，即选择部分表备份时，此参数需要填写。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tables []*BackupTableItem `json:"Tables,omitnil,omitempty" name:"Tables"`
}

type BackupPeriod struct {
	// 全量备份频率。目前仅支持"Weekly" - 每星期
	PeriodType *string `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`

	// 全量备份周期。取值范围为：["Monday","Tuesday","Wednesday","Thursday","Friday","Saturday","Sunday"]。
	Day []*string `json:"Day,omitnil,omitempty" name:"Day"`
}

type BackupStrategy struct {
	// 全量备份开始时间。周期性的全量备份将在当天该时间开始。
	BackupStartTime *string `json:"BackupStartTime,omitnil,omitempty" name:"BackupStartTime"`

	// 存储策略。
	StorageStrategy *StorageStrategy `json:"StorageStrategy,omitnil,omitempty" name:"StorageStrategy"`

	// 备份周期。
	BackupPeriod *BackupPeriod `json:"BackupPeriod,omitnil,omitempty" name:"BackupPeriod"`

	// 备份方法。目前仅支持 "logical" - 逻辑备份。
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// 备份周期。支持的值包括：
	// "period" - 周期性备份；
	// "single" - 单次备份。
	// 默认值为"period"。
	StrategyType *string `json:"StrategyType,omitnil,omitempty" name:"StrategyType"`

	// 是否开启增量备份。可能的取值为["true", "false"]。默认值为"true"。
	EnableIncrement *bool `json:"EnableIncrement,omitnil,omitempty" name:"EnableIncrement"`
}

type BackupTableItem struct {
	// 表名。
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 重命名后的表名。
	NewTableName *string `json:"NewTableName,omitnil,omitempty" name:"NewTableName"`

	// 列对象。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Columns []*BackupColumnItem `json:"Columns,omitnil,omitempty" name:"Columns"`
}

// Predefined struct for user
type ConfigureBackupPlanRequestParams struct {
	// 备份计划 ID。
	BackupPlanId *string `json:"BackupPlanId,omitnil,omitempty" name:"BackupPlanId"`

	// 备份计划名称。支持数字、英文大小写字母、中文以及特殊字符_-./()（）[]+=：:@,且长度不能超过60。
	BackupPlanName *string `json:"BackupPlanName,omitnil,omitempty" name:"BackupPlanName"`

	// 备份源实例信息。
	SourceEndPoint *BackupEndpoint `json:"SourceEndPoint,omitnil,omitempty" name:"SourceEndPoint"`

	// 备份对象信息。
	BackupObject *BackupObject `json:"BackupObject,omitnil,omitempty" name:"BackupObject"`

	// 备份策略。
	BackupStrategy *BackupStrategy `json:"BackupStrategy,omitnil,omitempty" name:"BackupStrategy"`

	// 加密信息。当需要使用SSE-KMS需要传入该值，你可以通过 KMS 的 GenerateDataKey 接口生成。
	PlainText *string `json:"PlainText,omitnil,omitempty" name:"PlainText"`
}

type ConfigureBackupPlanRequest struct {
	*tchttp.BaseRequest
	
	// 备份计划 ID。
	BackupPlanId *string `json:"BackupPlanId,omitnil,omitempty" name:"BackupPlanId"`

	// 备份计划名称。支持数字、英文大小写字母、中文以及特殊字符_-./()（）[]+=：:@,且长度不能超过60。
	BackupPlanName *string `json:"BackupPlanName,omitnil,omitempty" name:"BackupPlanName"`

	// 备份源实例信息。
	SourceEndPoint *BackupEndpoint `json:"SourceEndPoint,omitnil,omitempty" name:"SourceEndPoint"`

	// 备份对象信息。
	BackupObject *BackupObject `json:"BackupObject,omitnil,omitempty" name:"BackupObject"`

	// 备份策略。
	BackupStrategy *BackupStrategy `json:"BackupStrategy,omitnil,omitempty" name:"BackupStrategy"`

	// 加密信息。当需要使用SSE-KMS需要传入该值，你可以通过 KMS 的 GenerateDataKey 接口生成。
	PlainText *string `json:"PlainText,omitnil,omitempty" name:"PlainText"`
}

func (r *ConfigureBackupPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ConfigureBackupPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BackupPlanId")
	delete(f, "BackupPlanName")
	delete(f, "SourceEndPoint")
	delete(f, "BackupObject")
	delete(f, "BackupStrategy")
	delete(f, "PlainText")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ConfigureBackupPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ConfigureBackupPlanResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ConfigureBackupPlanResponse struct {
	*tchttp.BaseResponse
	Response *ConfigureBackupPlanResponseParams `json:"Response"`
}

func (r *ConfigureBackupPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ConfigureBackupPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConnectTestJobRequestParams struct {
	// 备份源实例信息。
	Endpoint *BackupEndpoint `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`
}

type CreateConnectTestJobRequest struct {
	*tchttp.BaseRequest
	
	// 备份源实例信息。
	Endpoint *BackupEndpoint `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`
}

func (r *CreateConnectTestJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConnectTestJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Endpoint")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateConnectTestJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConnectTestJobResponseParams struct {
	// 连通性任务 ID。
	ConnTaskId *string `json:"ConnTaskId,omitnil,omitempty" name:"ConnTaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateConnectTestJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateConnectTestJobResponseParams `json:"Response"`
}

func (r *CreateConnectTestJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConnectTestJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupCheckJobRequestParams struct {
	// 备份计划 ID。
	BackupPlanId *string `json:"BackupPlanId,omitnil,omitempty" name:"BackupPlanId"`
}

type DescribeBackupCheckJobRequest struct {
	*tchttp.BaseRequest
	
	// 备份计划 ID。
	BackupPlanId *string `json:"BackupPlanId,omitnil,omitempty" name:"BackupPlanId"`
}

func (r *DescribeBackupCheckJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupCheckJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BackupPlanId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupCheckJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupCheckJobResponseParams struct {
	// 校验任务状态。可能的取值为："finished" - 已完成; "running" - 进行中。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务进度。取值范围为[0, 100]，表示任务完成的百分比。例如：30表示任务完成30%。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 校验是否通过标记。可能的取值为："1" - 校验通过;"0" - 校验未通过。
	CheckFlag *int64 `json:"CheckFlag,omitnil,omitempty" name:"CheckFlag"`

	// 错误信息。
	ErrMessage *string `json:"ErrMessage,omitnil,omitempty" name:"ErrMessage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupCheckJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupCheckJobResponseParams `json:"Response"`
}

func (r *DescribeBackupCheckJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupCheckJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartBackupCheckJobRequestParams struct {
	// 备份计划 ID。
	BackupPlanId *string `json:"BackupPlanId,omitnil,omitempty" name:"BackupPlanId"`
}

type StartBackupCheckJobRequest struct {
	*tchttp.BaseRequest
	
	// 备份计划 ID。
	BackupPlanId *string `json:"BackupPlanId,omitnil,omitempty" name:"BackupPlanId"`
}

func (r *StartBackupCheckJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartBackupCheckJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BackupPlanId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartBackupCheckJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartBackupCheckJobResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartBackupCheckJobResponse struct {
	*tchttp.BaseResponse
	Response *StartBackupCheckJobResponseParams `json:"Response"`
}

func (r *StartBackupCheckJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartBackupCheckJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartBackupPlanRequestParams struct {
	// 备份计划 ID。
	BackupPlanId *string `json:"BackupPlanId,omitnil,omitempty" name:"BackupPlanId"`
}

type StartBackupPlanRequest struct {
	*tchttp.BaseRequest
	
	// 备份计划 ID。
	BackupPlanId *string `json:"BackupPlanId,omitnil,omitempty" name:"BackupPlanId"`
}

func (r *StartBackupPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartBackupPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BackupPlanId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartBackupPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartBackupPlanResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartBackupPlanResponse struct {
	*tchttp.BaseResponse
	Response *StartBackupPlanResponseParams `json:"Response"`
}

func (r *StartBackupPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartBackupPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StorageStrategy struct {
	// 存储类型。目前仅支持 "system" - DBS 内置存储。默认值为 "system"。
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 加密方式。可能的取值为：
	// "UnEncrypted" - 非加密存储;
	// "SSE-COS" - 内置加密存储;
	// 当该参数用作入参时，默认值为 "UnEncrypted"。
	Encryption *string `json:"Encryption,omitnil,omitempty" name:"Encryption"`

	// 日志保留时间，单位为天。取值范围为[7, 3650]，默认值为 30。
	BackupRetentionPeriod *int64 `json:"BackupRetentionPeriod,omitnil,omitempty" name:"BackupRetentionPeriod"`
}