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

package v20201028

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AutomationAgentInfo struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Agent 版本号。
	Version *string `json:"Version,omitnil" name:"Version"`

	// 上次心跳时间
	LastHeartbeatTime *string `json:"LastHeartbeatTime,omitnil" name:"LastHeartbeatTime"`

	// Agent状态，取值范围：
	// <li> Online：在线
	// <li> Offline：离线
	AgentStatus *string `json:"AgentStatus,omitnil" name:"AgentStatus"`

	// Agent运行环境，取值范围：
	// <li> Linux：Linux实例
	// <li> Windows：Windows实例
	Environment *string `json:"Environment,omitnil" name:"Environment"`

	// Agent 支持的功能列表。
	SupportFeatures []*string `json:"SupportFeatures,omitnil" name:"SupportFeatures"`
}

// Predefined struct for user
type CancelInvocationRequestParams struct {
	// 执行活动ID
	InvocationId *string `json:"InvocationId,omitnil" name:"InvocationId"`

	// 实例ID列表，上限100。支持实例类型：
	// <li> CVM
	// <li> LIGHTHOUSE
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
}

type CancelInvocationRequest struct {
	*tchttp.BaseRequest
	
	// 执行活动ID
	InvocationId *string `json:"InvocationId,omitnil" name:"InvocationId"`

	// 实例ID列表，上限100。支持实例类型：
	// <li> CVM
	// <li> LIGHTHOUSE
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
}

func (r *CancelInvocationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelInvocationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InvocationId")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelInvocationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelInvocationResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CancelInvocationResponse struct {
	*tchttp.BaseResponse
	Response *CancelInvocationResponseParams `json:"Response"`
}

func (r *CancelInvocationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelInvocationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Command struct {
	// 命令ID。
	CommandId *string `json:"CommandId,omitnil" name:"CommandId"`

	// 命令名称。
	CommandName *string `json:"CommandName,omitnil" name:"CommandName"`

	// 命令描述。
	Description *string `json:"Description,omitnil" name:"Description"`

	// Base64编码后的命令内容。
	Content *string `json:"Content,omitnil" name:"Content"`

	// 命令类型。
	CommandType *string `json:"CommandType,omitnil" name:"CommandType"`

	// 命令执行路径。
	WorkingDirectory *string `json:"WorkingDirectory,omitnil" name:"WorkingDirectory"`

	// 命令超时时间。
	Timeout *uint64 `json:"Timeout,omitnil" name:"Timeout"`

	// 命令创建时间。
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// 命令更新时间。
	UpdatedTime *string `json:"UpdatedTime,omitnil" name:"UpdatedTime"`

	// 是否启用自定义参数功能。
	EnableParameter *bool `json:"EnableParameter,omitnil" name:"EnableParameter"`

	// 自定义参数的默认取值。
	DefaultParameters *string `json:"DefaultParameters,omitnil" name:"DefaultParameters"`

	// 自定义参数的默认取值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultParameterConfs []*DefaultParameterConf `json:"DefaultParameterConfs,omitnil" name:"DefaultParameterConfs"`

	// 命令的结构化描述。公共命令有值，用户命令为空字符串。
	FormattedDescription *string `json:"FormattedDescription,omitnil" name:"FormattedDescription"`

	// 命令创建者。TAT 代表公共命令，USER 代表个人命令。
	CreatedBy *string `json:"CreatedBy,omitnil" name:"CreatedBy"`

	// 命令关联的标签列表。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// 在实例上执行命令的用户名。
	Username *string `json:"Username,omitnil" name:"Username"`

	// 日志上传的cos bucket 地址。
	OutputCOSBucketUrl *string `json:"OutputCOSBucketUrl,omitnil" name:"OutputCOSBucketUrl"`

	// 日志在cos bucket中的目录。
	OutputCOSKeyPrefix *string `json:"OutputCOSKeyPrefix,omitnil" name:"OutputCOSKeyPrefix"`
}

type CommandDocument struct {
	// Base64 编码后的执行命令。
	Content *string `json:"Content,omitnil" name:"Content"`

	// 命令类型。
	CommandType *string `json:"CommandType,omitnil" name:"CommandType"`

	// 超时时间。
	Timeout *uint64 `json:"Timeout,omitnil" name:"Timeout"`

	// 执行路径。
	WorkingDirectory *string `json:"WorkingDirectory,omitnil" name:"WorkingDirectory"`

	// 执行用户。
	Username *string `json:"Username,omitnil" name:"Username"`

	// 保存输出的 COS Bucket 链接。
	OutputCOSBucketUrl *string `json:"OutputCOSBucketUrl,omitnil" name:"OutputCOSBucketUrl"`

	// 保存输出的文件名称前缀。
	OutputCOSKeyPrefix *string `json:"OutputCOSKeyPrefix,omitnil" name:"OutputCOSKeyPrefix"`
}

// Predefined struct for user
type CreateCommandRequestParams struct {
	// 命令名称。名称仅支持中文、英文、数字、下划线、分隔符"-"、小数点，最大长度不能超60个字节。
	CommandName *string `json:"CommandName,omitnil" name:"CommandName"`

	// Base64编码后的命令内容，长度不可超过64KB。
	Content *string `json:"Content,omitnil" name:"Content"`

	// 命令描述。不超过120字符。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 命令类型，目前支持取值：SHELL、POWERSHELL。默认：SHELL。
	CommandType *string `json:"CommandType,omitnil" name:"CommandType"`

	// 命令执行路径，对于 SHELL 命令默认为 /root，对于 POWERSHELL 命令默认为 C:\Program Files\qcloud\tat_agent\workdir。
	WorkingDirectory *string `json:"WorkingDirectory,omitnil" name:"WorkingDirectory"`

	// 命令超时时间，默认60秒。取值范围[1, 86400]。
	Timeout *uint64 `json:"Timeout,omitnil" name:"Timeout"`

	// 是否启用自定义参数功能。
	// 一旦创建，此值不提供修改。
	// 默认值：false。
	EnableParameter *bool `json:"EnableParameter,omitnil" name:"EnableParameter"`

	// 启用自定义参数功能时，自定义参数的默认取值。字段类型为json encoded string。如：{\"varA\": \"222\"}。
	// key为自定义参数名称，value为该参数的默认取值。kv均为字符串型。
	// 如果InvokeCommand时未提供参数取值，将使用这里的默认值进行替换。
	// 自定义参数最多20个。
	// 自定义参数名称需符合以下规范：字符数目上限64，可选范围【a-zA-Z0-9-_】。
	DefaultParameters *string `json:"DefaultParameters,omitnil" name:"DefaultParameters"`

	// 自定义参数数组。
	// 如果InvokeCommand时未提供参数取值，将使用这里的默认值进行替换。
	// 自定义参数最多20个。
	DefaultParameterConfs []*DefaultParameterConf `json:"DefaultParameterConfs,omitnil" name:"DefaultParameterConfs"`

	// 为命令关联的标签，列表长度不超过10。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// 在 CVM 或 Lighthouse 实例中执行命令的用户名称。
	// 使用最小权限执行命令是权限管理的最佳实践，建议您以普通用户身份执行云助手命令。默认情况下，在 Linux 实例中以 root 用户执行命令；在Windows 实例中以 System 用户执行命令。
	Username *string `json:"Username,omitnil" name:"Username"`

	// 指定日志上传的cos bucket 地址，必须以https开头，如 https://BucketName-123454321.cos.ap-beijing.myqcloud.com。
	OutputCOSBucketUrl *string `json:"OutputCOSBucketUrl,omitnil" name:"OutputCOSBucketUrl"`

	// 指定日志在cos bucket中的目录，目录命名有如下规则：
	// 1. 可用数字、中英文和可见字符的组合，长度最多为60。
	// 2. 用 / 分割路径，可快速创建子目录。
	// 3. 不允许连续 / ；不允许以 / 开头；不允许以..作为文件夹名称
	OutputCOSKeyPrefix *string `json:"OutputCOSKeyPrefix,omitnil" name:"OutputCOSKeyPrefix"`
}

type CreateCommandRequest struct {
	*tchttp.BaseRequest
	
	// 命令名称。名称仅支持中文、英文、数字、下划线、分隔符"-"、小数点，最大长度不能超60个字节。
	CommandName *string `json:"CommandName,omitnil" name:"CommandName"`

	// Base64编码后的命令内容，长度不可超过64KB。
	Content *string `json:"Content,omitnil" name:"Content"`

	// 命令描述。不超过120字符。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 命令类型，目前支持取值：SHELL、POWERSHELL。默认：SHELL。
	CommandType *string `json:"CommandType,omitnil" name:"CommandType"`

	// 命令执行路径，对于 SHELL 命令默认为 /root，对于 POWERSHELL 命令默认为 C:\Program Files\qcloud\tat_agent\workdir。
	WorkingDirectory *string `json:"WorkingDirectory,omitnil" name:"WorkingDirectory"`

	// 命令超时时间，默认60秒。取值范围[1, 86400]。
	Timeout *uint64 `json:"Timeout,omitnil" name:"Timeout"`

	// 是否启用自定义参数功能。
	// 一旦创建，此值不提供修改。
	// 默认值：false。
	EnableParameter *bool `json:"EnableParameter,omitnil" name:"EnableParameter"`

	// 启用自定义参数功能时，自定义参数的默认取值。字段类型为json encoded string。如：{\"varA\": \"222\"}。
	// key为自定义参数名称，value为该参数的默认取值。kv均为字符串型。
	// 如果InvokeCommand时未提供参数取值，将使用这里的默认值进行替换。
	// 自定义参数最多20个。
	// 自定义参数名称需符合以下规范：字符数目上限64，可选范围【a-zA-Z0-9-_】。
	DefaultParameters *string `json:"DefaultParameters,omitnil" name:"DefaultParameters"`

	// 自定义参数数组。
	// 如果InvokeCommand时未提供参数取值，将使用这里的默认值进行替换。
	// 自定义参数最多20个。
	DefaultParameterConfs []*DefaultParameterConf `json:"DefaultParameterConfs,omitnil" name:"DefaultParameterConfs"`

	// 为命令关联的标签，列表长度不超过10。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// 在 CVM 或 Lighthouse 实例中执行命令的用户名称。
	// 使用最小权限执行命令是权限管理的最佳实践，建议您以普通用户身份执行云助手命令。默认情况下，在 Linux 实例中以 root 用户执行命令；在Windows 实例中以 System 用户执行命令。
	Username *string `json:"Username,omitnil" name:"Username"`

	// 指定日志上传的cos bucket 地址，必须以https开头，如 https://BucketName-123454321.cos.ap-beijing.myqcloud.com。
	OutputCOSBucketUrl *string `json:"OutputCOSBucketUrl,omitnil" name:"OutputCOSBucketUrl"`

	// 指定日志在cos bucket中的目录，目录命名有如下规则：
	// 1. 可用数字、中英文和可见字符的组合，长度最多为60。
	// 2. 用 / 分割路径，可快速创建子目录。
	// 3. 不允许连续 / ；不允许以 / 开头；不允许以..作为文件夹名称
	OutputCOSKeyPrefix *string `json:"OutputCOSKeyPrefix,omitnil" name:"OutputCOSKeyPrefix"`
}

func (r *CreateCommandRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCommandRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CommandName")
	delete(f, "Content")
	delete(f, "Description")
	delete(f, "CommandType")
	delete(f, "WorkingDirectory")
	delete(f, "Timeout")
	delete(f, "EnableParameter")
	delete(f, "DefaultParameters")
	delete(f, "DefaultParameterConfs")
	delete(f, "Tags")
	delete(f, "Username")
	delete(f, "OutputCOSBucketUrl")
	delete(f, "OutputCOSKeyPrefix")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCommandRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCommandResponseParams struct {
	// 命令ID。
	CommandId *string `json:"CommandId,omitnil" name:"CommandId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateCommandResponse struct {
	*tchttp.BaseResponse
	Response *CreateCommandResponseParams `json:"Response"`
}

func (r *CreateCommandResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCommandResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInvokerRequestParams struct {
	// 执行器名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 执行器类型，当前仅支持周期类型执行器，取值：`SCHEDULE` 。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 远程命令ID。
	CommandId *string `json:"CommandId,omitnil" name:"CommandId"`

	// 触发器关联的实例ID。列表上限 100。
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// 命令执行用户。
	Username *string `json:"Username,omitnil" name:"Username"`

	// 命令自定义参数。
	Parameters *string `json:"Parameters,omitnil" name:"Parameters"`

	// 周期执行器设置，当创建周期执行器时，必须指定此参数。
	ScheduleSettings *ScheduleSettings `json:"ScheduleSettings,omitnil" name:"ScheduleSettings"`
}

type CreateInvokerRequest struct {
	*tchttp.BaseRequest
	
	// 执行器名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 执行器类型，当前仅支持周期类型执行器，取值：`SCHEDULE` 。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 远程命令ID。
	CommandId *string `json:"CommandId,omitnil" name:"CommandId"`

	// 触发器关联的实例ID。列表上限 100。
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// 命令执行用户。
	Username *string `json:"Username,omitnil" name:"Username"`

	// 命令自定义参数。
	Parameters *string `json:"Parameters,omitnil" name:"Parameters"`

	// 周期执行器设置，当创建周期执行器时，必须指定此参数。
	ScheduleSettings *ScheduleSettings `json:"ScheduleSettings,omitnil" name:"ScheduleSettings"`
}

func (r *CreateInvokerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInvokerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "CommandId")
	delete(f, "InstanceIds")
	delete(f, "Username")
	delete(f, "Parameters")
	delete(f, "ScheduleSettings")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInvokerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInvokerResponseParams struct {
	// 执行器ID。
	InvokerId *string `json:"InvokerId,omitnil" name:"InvokerId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateInvokerResponse struct {
	*tchttp.BaseResponse
	Response *CreateInvokerResponseParams `json:"Response"`
}

func (r *CreateInvokerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInvokerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRegisterCodeRequestParams struct {
	// 注册码描述。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 注册实列名称前缀。
	InstanceNamePrefix *string `json:"InstanceNamePrefix,omitnil" name:"InstanceNamePrefix"`

	// 该注册码允许注册的实列数目。默认限制为10个。
	RegisterLimit *int64 `json:"RegisterLimit,omitnil" name:"RegisterLimit"`

	// 该注册码的有效时间，单位为小时。默认为4小时。
	EffectiveTime *int64 `json:"EffectiveTime,omitnil" name:"EffectiveTime"`

	// 该注册码限制tat_agent只能从IpAddressRange所描述公网出口进行注册。默认不做限制。
	IpAddressRange *string `json:"IpAddressRange,omitnil" name:"IpAddressRange"`
}

type CreateRegisterCodeRequest struct {
	*tchttp.BaseRequest
	
	// 注册码描述。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 注册实列名称前缀。
	InstanceNamePrefix *string `json:"InstanceNamePrefix,omitnil" name:"InstanceNamePrefix"`

	// 该注册码允许注册的实列数目。默认限制为10个。
	RegisterLimit *int64 `json:"RegisterLimit,omitnil" name:"RegisterLimit"`

	// 该注册码的有效时间，单位为小时。默认为4小时。
	EffectiveTime *int64 `json:"EffectiveTime,omitnil" name:"EffectiveTime"`

	// 该注册码限制tat_agent只能从IpAddressRange所描述公网出口进行注册。默认不做限制。
	IpAddressRange *string `json:"IpAddressRange,omitnil" name:"IpAddressRange"`
}

func (r *CreateRegisterCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRegisterCodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Description")
	delete(f, "InstanceNamePrefix")
	delete(f, "RegisterLimit")
	delete(f, "EffectiveTime")
	delete(f, "IpAddressRange")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRegisterCodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRegisterCodeResponseParams struct {
	// 注册码ID。
	RegisterCodeId *string `json:"RegisterCodeId,omitnil" name:"RegisterCodeId"`

	// 注册码值。
	RegisterCodeValue *string `json:"RegisterCodeValue,omitnil" name:"RegisterCodeValue"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateRegisterCodeResponse struct {
	*tchttp.BaseResponse
	Response *CreateRegisterCodeResponseParams `json:"Response"`
}

func (r *CreateRegisterCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRegisterCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DefaultParameterConf struct {
	// 参数名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParameterName *string `json:"ParameterName,omitnil" name:"ParameterName"`

	// 参数默认值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParameterValue *string `json:"ParameterValue,omitnil" name:"ParameterValue"`

	// 参数描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParameterDescription *string `json:"ParameterDescription,omitnil" name:"ParameterDescription"`
}

// Predefined struct for user
type DeleteCommandRequestParams struct {
	// 待删除的命令ID。
	CommandId *string `json:"CommandId,omitnil" name:"CommandId"`
}

type DeleteCommandRequest struct {
	*tchttp.BaseRequest
	
	// 待删除的命令ID。
	CommandId *string `json:"CommandId,omitnil" name:"CommandId"`
}

func (r *DeleteCommandRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCommandRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CommandId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCommandRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCommandResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteCommandResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCommandResponseParams `json:"Response"`
}

func (r *DeleteCommandResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCommandResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCommandsRequestParams struct {
	// 待删除命令id
	CommandIds []*string `json:"CommandIds,omitnil" name:"CommandIds"`
}

type DeleteCommandsRequest struct {
	*tchttp.BaseRequest
	
	// 待删除命令id
	CommandIds []*string `json:"CommandIds,omitnil" name:"CommandIds"`
}

func (r *DeleteCommandsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCommandsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CommandIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCommandsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCommandsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteCommandsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCommandsResponseParams `json:"Response"`
}

func (r *DeleteCommandsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCommandsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInvokerRequestParams struct {
	// 待删除的执行器ID。
	InvokerId *string `json:"InvokerId,omitnil" name:"InvokerId"`
}

type DeleteInvokerRequest struct {
	*tchttp.BaseRequest
	
	// 待删除的执行器ID。
	InvokerId *string `json:"InvokerId,omitnil" name:"InvokerId"`
}

func (r *DeleteInvokerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInvokerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InvokerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteInvokerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInvokerResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteInvokerResponse struct {
	*tchttp.BaseResponse
	Response *DeleteInvokerResponseParams `json:"Response"`
}

func (r *DeleteInvokerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInvokerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRegisterCodesRequestParams struct {
	// 注册码ID列表。限制输入的注册码ID数量大于0小于100。
	RegisterCodeIds []*string `json:"RegisterCodeIds,omitnil" name:"RegisterCodeIds"`
}

type DeleteRegisterCodesRequest struct {
	*tchttp.BaseRequest
	
	// 注册码ID列表。限制输入的注册码ID数量大于0小于100。
	RegisterCodeIds []*string `json:"RegisterCodeIds,omitnil" name:"RegisterCodeIds"`
}

func (r *DeleteRegisterCodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRegisterCodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegisterCodeIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRegisterCodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRegisterCodesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteRegisterCodesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRegisterCodesResponseParams `json:"Response"`
}

func (r *DeleteRegisterCodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRegisterCodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRegisterInstanceRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DeleteRegisterInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DeleteRegisterInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRegisterInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRegisterInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRegisterInstanceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteRegisterInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRegisterInstanceResponseParams `json:"Response"`
}

func (r *DeleteRegisterInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRegisterInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutomationAgentStatusRequestParams struct {
	// 待查询的实例ID列表。
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// 过滤条件。<br> <li> agent-status - String - 是否必填：否 -（过滤条件）按照agent状态过滤，取值：Online 在线，Offline 离线。<br> <li> environment - String - 是否必填：否 -（过滤条件）按照agent运行环境查询，取值：Linux, Windows。<br> <li> instance-id - String - 是否必填：否 -（过滤条件）按照实例ID过滤。 <br>每次请求的 `Filters` 的上限为10， `Filter.Values` 的上限为5。参数不支持同时指定 `InstanceIds` 和 `Filters` 。
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 返回数量，默认为20，最大值为100。关于 `Limit` 的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。关于 `Offset` 的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeAutomationAgentStatusRequest struct {
	*tchttp.BaseRequest
	
	// 待查询的实例ID列表。
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// 过滤条件。<br> <li> agent-status - String - 是否必填：否 -（过滤条件）按照agent状态过滤，取值：Online 在线，Offline 离线。<br> <li> environment - String - 是否必填：否 -（过滤条件）按照agent运行环境查询，取值：Linux, Windows。<br> <li> instance-id - String - 是否必填：否 -（过滤条件）按照实例ID过滤。 <br>每次请求的 `Filters` 的上限为10， `Filter.Values` 的上限为5。参数不支持同时指定 `InstanceIds` 和 `Filters` 。
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 返回数量，默认为20，最大值为100。关于 `Limit` 的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。关于 `Offset` 的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeAutomationAgentStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutomationAgentStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAutomationAgentStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutomationAgentStatusResponseParams struct {
	// Agent 信息列表。
	AutomationAgentSet []*AutomationAgentInfo `json:"AutomationAgentSet,omitnil" name:"AutomationAgentSet"`

	// 符合条件的 Agent 总数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAutomationAgentStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAutomationAgentStatusResponseParams `json:"Response"`
}

func (r *DescribeAutomationAgentStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutomationAgentStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCommandsRequestParams struct {
	// 命令ID列表，每次请求的上限为100。参数不支持同时指定 `CommandIds` 和 `Filters` 。
	CommandIds []*string `json:"CommandIds,omitnil" name:"CommandIds"`

	// 过滤条件。
	// <li> command-id - String - 是否必填：否 -（过滤条件）按照命令ID过滤。
	// <li> command-name - String - 是否必填：否 -（过滤条件）按照命令名称过滤。
	// <li> command-type - String - 是否必填：否 -（过滤条件）按照命令类型过滤，取值为 SHELL 或 POWERSHELL。
	// <li> created-by - String - 是否必填：否 -（过滤条件）按照命令创建者过滤，取值为 TAT 或 USER，TAT 代表公共命令，USER 代表由用户创建的命令。
	// <li> tag-key - String - 是否必填：否 -（过滤条件）按照标签键进行过滤。</li>
	// <li> tag-value - String - 是否必填：否 -（过滤条件）按照标签值进行过滤。</li>
	// <li> tag:tag-key - String - 是否必填：否 -（过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。使用请参考示例4</li>
	// 
	// 每次请求的 `Filters` 的上限为10， `Filter.Values` 的上限为5。参数不支持同时指定 `CommandIds` 和 `Filters` 。
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 返回数量，默认为20，最大值为100。关于 `Limit` 的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。关于 `Offset` 的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeCommandsRequest struct {
	*tchttp.BaseRequest
	
	// 命令ID列表，每次请求的上限为100。参数不支持同时指定 `CommandIds` 和 `Filters` 。
	CommandIds []*string `json:"CommandIds,omitnil" name:"CommandIds"`

	// 过滤条件。
	// <li> command-id - String - 是否必填：否 -（过滤条件）按照命令ID过滤。
	// <li> command-name - String - 是否必填：否 -（过滤条件）按照命令名称过滤。
	// <li> command-type - String - 是否必填：否 -（过滤条件）按照命令类型过滤，取值为 SHELL 或 POWERSHELL。
	// <li> created-by - String - 是否必填：否 -（过滤条件）按照命令创建者过滤，取值为 TAT 或 USER，TAT 代表公共命令，USER 代表由用户创建的命令。
	// <li> tag-key - String - 是否必填：否 -（过滤条件）按照标签键进行过滤。</li>
	// <li> tag-value - String - 是否必填：否 -（过滤条件）按照标签值进行过滤。</li>
	// <li> tag:tag-key - String - 是否必填：否 -（过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。使用请参考示例4</li>
	// 
	// 每次请求的 `Filters` 的上限为10， `Filter.Values` 的上限为5。参数不支持同时指定 `CommandIds` 和 `Filters` 。
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 返回数量，默认为20，最大值为100。关于 `Limit` 的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。关于 `Offset` 的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeCommandsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCommandsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CommandIds")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCommandsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCommandsResponseParams struct {
	// 符合条件的命令总数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 命令详情列表。
	CommandSet []*Command `json:"CommandSet,omitnil" name:"CommandSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCommandsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCommandsResponseParams `json:"Response"`
}

func (r *DescribeCommandsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCommandsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInvocationTasksRequestParams struct {
	// 执行任务ID列表，每次请求的上限为100。参数不支持同时指定 `InvocationTaskIds` 和 `Filters`。
	InvocationTaskIds []*string `json:"InvocationTaskIds,omitnil" name:"InvocationTaskIds"`

	// 过滤条件。<br> <li> invocation-id - String - 是否必填：否 -（过滤条件）按照执行活动ID过滤。<br> <li> invocation-task-id - String - 是否必填：否 -（过滤条件）按照执行任务ID过滤。<br> <li> instance-id - String - 是否必填：否 -（过滤条件）按照实例ID过滤。 <br> <li> command-id - String - 是否必填：否 -（过滤条件）按照命令ID过滤。 <br>每次请求的 `Filters` 的上限为10， `Filter.Values` 的上限为5。参数不支持同时指定 `InvocationTaskIds` 和 `Filters` 。
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 返回数量，默认为20，最大值为100。关于 `Limit` 的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。关于 `Offset` 的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 是否隐藏输出，取值范围：<br><li>true：隐藏输出 <br><li>false：不隐藏 <br>默认为 false。
	HideOutput *bool `json:"HideOutput,omitnil" name:"HideOutput"`
}

type DescribeInvocationTasksRequest struct {
	*tchttp.BaseRequest
	
	// 执行任务ID列表，每次请求的上限为100。参数不支持同时指定 `InvocationTaskIds` 和 `Filters`。
	InvocationTaskIds []*string `json:"InvocationTaskIds,omitnil" name:"InvocationTaskIds"`

	// 过滤条件。<br> <li> invocation-id - String - 是否必填：否 -（过滤条件）按照执行活动ID过滤。<br> <li> invocation-task-id - String - 是否必填：否 -（过滤条件）按照执行任务ID过滤。<br> <li> instance-id - String - 是否必填：否 -（过滤条件）按照实例ID过滤。 <br> <li> command-id - String - 是否必填：否 -（过滤条件）按照命令ID过滤。 <br>每次请求的 `Filters` 的上限为10， `Filter.Values` 的上限为5。参数不支持同时指定 `InvocationTaskIds` 和 `Filters` 。
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 返回数量，默认为20，最大值为100。关于 `Limit` 的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。关于 `Offset` 的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 是否隐藏输出，取值范围：<br><li>true：隐藏输出 <br><li>false：不隐藏 <br>默认为 false。
	HideOutput *bool `json:"HideOutput,omitnil" name:"HideOutput"`
}

func (r *DescribeInvocationTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInvocationTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InvocationTaskIds")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "HideOutput")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInvocationTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInvocationTasksResponseParams struct {
	// 符合条件的执行任务总数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 执行任务列表。
	InvocationTaskSet []*InvocationTask `json:"InvocationTaskSet,omitnil" name:"InvocationTaskSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInvocationTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInvocationTasksResponseParams `json:"Response"`
}

func (r *DescribeInvocationTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInvocationTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInvocationsRequestParams struct {
	// 执行活动ID列表，每次请求的上限为100。参数不支持同时指定 `InvocationIds` 和 `Filters`。
	InvocationIds []*string `json:"InvocationIds,omitnil" name:"InvocationIds"`

	// 过滤条件。<br> <li> invocation-id - String - 是否必填：否 -（过滤条件）按照执行活动ID过滤。<br> 
	// <li> command-id - String - 是否必填：否 -（过滤条件）按照命令ID过滤。 
	// <li> command-created-by - String - 是否必填：否 -（过滤条件）按照执行的命令类型过滤，取值为 TAT 或 USER，TAT 代表公共命令，USER 代表由用户创建的命令。
	// <li> instance-kind - String - 是否必填：否 -（过滤条件）按照运行实例类型过滤，取值为 CVM 或 LIGHTHOUSE，CVM 代表实例为云服务器， LIGHTHOUSE 代表实例为轻量应用服务器。
	// <br>每次请求的 `Filters` 的上限为10， `Filter.Values` 的上限为5。参数不支持同时指定 `InvocationIds` 和 `Filters` 。
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 返回数量，默认为20，最大值为100。关于 `Limit` 的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。关于 `Offset` 的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeInvocationsRequest struct {
	*tchttp.BaseRequest
	
	// 执行活动ID列表，每次请求的上限为100。参数不支持同时指定 `InvocationIds` 和 `Filters`。
	InvocationIds []*string `json:"InvocationIds,omitnil" name:"InvocationIds"`

	// 过滤条件。<br> <li> invocation-id - String - 是否必填：否 -（过滤条件）按照执行活动ID过滤。<br> 
	// <li> command-id - String - 是否必填：否 -（过滤条件）按照命令ID过滤。 
	// <li> command-created-by - String - 是否必填：否 -（过滤条件）按照执行的命令类型过滤，取值为 TAT 或 USER，TAT 代表公共命令，USER 代表由用户创建的命令。
	// <li> instance-kind - String - 是否必填：否 -（过滤条件）按照运行实例类型过滤，取值为 CVM 或 LIGHTHOUSE，CVM 代表实例为云服务器， LIGHTHOUSE 代表实例为轻量应用服务器。
	// <br>每次请求的 `Filters` 的上限为10， `Filter.Values` 的上限为5。参数不支持同时指定 `InvocationIds` 和 `Filters` 。
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 返回数量，默认为20，最大值为100。关于 `Limit` 的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。关于 `Offset` 的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeInvocationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInvocationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InvocationIds")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInvocationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInvocationsResponseParams struct {
	// 符合条件的执行活动总数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 执行活动列表。
	InvocationSet []*Invocation `json:"InvocationSet,omitnil" name:"InvocationSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInvocationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInvocationsResponseParams `json:"Response"`
}

func (r *DescribeInvocationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInvocationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInvokerRecordsRequestParams struct {
	// 执行器ID列表。列表上限 100。
	InvokerIds []*string `json:"InvokerIds,omitnil" name:"InvokerIds"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeInvokerRecordsRequest struct {
	*tchttp.BaseRequest
	
	// 执行器ID列表。列表上限 100。
	InvokerIds []*string `json:"InvokerIds,omitnil" name:"InvokerIds"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeInvokerRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInvokerRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InvokerIds")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInvokerRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInvokerRecordsResponseParams struct {
	// 符合条件的历史记录数量。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 执行器执行历史记录。
	InvokerRecordSet []*InvokerRecord `json:"InvokerRecordSet,omitnil" name:"InvokerRecordSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInvokerRecordsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInvokerRecordsResponseParams `json:"Response"`
}

func (r *DescribeInvokerRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInvokerRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInvokersRequestParams struct {
	// 执行器ID列表。
	InvokerIds []*string `json:"InvokerIds,omitnil" name:"InvokerIds"`

	// 过滤条件：
	// 
	// <li> invoker-id - String - 是否必填：否 - （过滤条件）按执行器ID过滤。
	// <li> command-id - String - 是否必填：否 - （过滤条件）按命令ID过滤。
	// <li> type - String - 是否必填：否 - （过滤条件）按执行器类型过滤。
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeInvokersRequest struct {
	*tchttp.BaseRequest
	
	// 执行器ID列表。
	InvokerIds []*string `json:"InvokerIds,omitnil" name:"InvokerIds"`

	// 过滤条件：
	// 
	// <li> invoker-id - String - 是否必填：否 - （过滤条件）按执行器ID过滤。
	// <li> command-id - String - 是否必填：否 - （过滤条件）按命令ID过滤。
	// <li> type - String - 是否必填：否 - （过滤条件）按执行器类型过滤。
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeInvokersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInvokersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InvokerIds")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInvokersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInvokersResponseParams struct {
	// 满足条件的执行器数量。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 执行器信息。
	InvokerSet []*Invoker `json:"InvokerSet,omitnil" name:"InvokerSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInvokersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInvokersResponseParams `json:"Response"`
}

func (r *DescribeInvokersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInvokersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQuotasRequestParams struct {
	// 资源名称，目前有"COMMAND","REGISTER_CODE" 这两个指标
	ResourceNames []*string `json:"ResourceNames,omitnil" name:"ResourceNames"`
}

type DescribeQuotasRequest struct {
	*tchttp.BaseRequest
	
	// 资源名称，目前有"COMMAND","REGISTER_CODE" 这两个指标
	ResourceNames []*string `json:"ResourceNames,omitnil" name:"ResourceNames"`
}

func (r *DescribeQuotasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQuotasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeQuotasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQuotasResponseParams struct {
	// 资源额度列表
	GeneralResourceQuotaSet []*GeneralResourceQuotaSet `json:"GeneralResourceQuotaSet,omitnil" name:"GeneralResourceQuotaSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeQuotasResponse struct {
	*tchttp.BaseResponse
	Response *DescribeQuotasResponseParams `json:"Response"`
}

func (r *DescribeQuotasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQuotasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegionsRequestParams struct {

}

type DescribeRegionsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegionsResponseParams struct {
	// 地域数量
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 地域信息列表
	RegionSet []*RegionInfo `json:"RegionSet,omitnil" name:"RegionSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRegionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRegionsResponseParams `json:"Response"`
}

func (r *DescribeRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegisterCodesRequestParams struct {
	// 注册码ID。
	RegisterCodeIds []*string `json:"RegisterCodeIds,omitnil" name:"RegisterCodeIds"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeRegisterCodesRequest struct {
	*tchttp.BaseRequest
	
	// 注册码ID。
	RegisterCodeIds []*string `json:"RegisterCodeIds,omitnil" name:"RegisterCodeIds"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeRegisterCodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegisterCodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegisterCodeIds")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRegisterCodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegisterCodesResponseParams struct {
	// 查询到的注册码总数。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 注册码信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegisterCodeSet []*RegisterCodeInfo `json:"RegisterCodeSet,omitnil" name:"RegisterCodeSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRegisterCodesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRegisterCodesResponseParams `json:"Response"`
}

func (r *DescribeRegisterCodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegisterCodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegisterInstancesRequestParams struct {
	// 实例id。
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// 过滤器列表。
	// 
	// - instance-name
	// 
	// 按照【实例名称】进行过滤。
	// 类型：String
	// 必选：否
	// 
	// - instance-id
	// 
	// 按照【实例ID】进行过滤。
	// 类型：String
	// 必选：否
	// 
	// - register-code-id
	// 
	// 按照【注册码ID】进行过滤。
	// 类型：String
	// 必选：否
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeRegisterInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例id。
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// 过滤器列表。
	// 
	// - instance-name
	// 
	// 按照【实例名称】进行过滤。
	// 类型：String
	// 必选：否
	// 
	// - instance-id
	// 
	// 按照【实例ID】进行过滤。
	// 类型：String
	// 必选：否
	// 
	// - register-code-id
	// 
	// 按照【注册码ID】进行过滤。
	// 类型：String
	// 必选：否
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeRegisterInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegisterInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRegisterInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegisterInstancesResponseParams struct {
	// 该实例注册过的注册码总数。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 被托管的实例信息的列表。
	RegisterInstanceSet []*RegisterInstanceInfo `json:"RegisterInstanceSet,omitnil" name:"RegisterInstanceSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRegisterInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRegisterInstancesResponseParams `json:"Response"`
}

func (r *DescribeRegisterInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegisterInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableInvokerRequestParams struct {
	// 待停止的执行器ID。
	InvokerId *string `json:"InvokerId,omitnil" name:"InvokerId"`
}

type DisableInvokerRequest struct {
	*tchttp.BaseRequest
	
	// 待停止的执行器ID。
	InvokerId *string `json:"InvokerId,omitnil" name:"InvokerId"`
}

func (r *DisableInvokerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableInvokerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InvokerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableInvokerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableInvokerResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DisableInvokerResponse struct {
	*tchttp.BaseResponse
	Response *DisableInvokerResponseParams `json:"Response"`
}

func (r *DisableInvokerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableInvokerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableRegisterCodesRequestParams struct {
	// 注册码ID。
	RegisterCodeIds []*string `json:"RegisterCodeIds,omitnil" name:"RegisterCodeIds"`
}

type DisableRegisterCodesRequest struct {
	*tchttp.BaseRequest
	
	// 注册码ID。
	RegisterCodeIds []*string `json:"RegisterCodeIds,omitnil" name:"RegisterCodeIds"`
}

func (r *DisableRegisterCodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableRegisterCodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegisterCodeIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableRegisterCodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableRegisterCodesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DisableRegisterCodesResponse struct {
	*tchttp.BaseResponse
	Response *DisableRegisterCodesResponseParams `json:"Response"`
}

func (r *DisableRegisterCodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableRegisterCodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableInvokerRequestParams struct {
	// 待启用的执行器ID。
	InvokerId *string `json:"InvokerId,omitnil" name:"InvokerId"`
}

type EnableInvokerRequest struct {
	*tchttp.BaseRequest
	
	// 待启用的执行器ID。
	InvokerId *string `json:"InvokerId,omitnil" name:"InvokerId"`
}

func (r *EnableInvokerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableInvokerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InvokerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableInvokerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableInvokerResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type EnableInvokerResponse struct {
	*tchttp.BaseResponse
	Response *EnableInvokerResponseParams `json:"Response"`
}

func (r *EnableInvokerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableInvokerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 需要过滤的字段。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 字段的过滤值。
	Values []*string `json:"Values,omitnil" name:"Values"`
}

type GeneralResourceQuotaSet struct {
	// 资源名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceName *string `json:"ResourceName,omitnil" name:"ResourceName"`

	// 已使用额度
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceQuotaUsed *int64 `json:"ResourceQuotaUsed,omitnil" name:"ResourceQuotaUsed"`

	// 总额度
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceQuotaTotal *int64 `json:"ResourceQuotaTotal,omitnil" name:"ResourceQuotaTotal"`
}

type Invocation struct {
	// 执行活动ID。
	InvocationId *string `json:"InvocationId,omitnil" name:"InvocationId"`

	// 命令ID。
	CommandId *string `json:"CommandId,omitnil" name:"CommandId"`

	// 执行任务状态。取值范围：
	// <li> PENDING：等待下发 
	// <li> RUNNING：命令运行中
	// <li> SUCCESS：命令成功
	// <li> FAILED：命令失败
	// <li> TIMEOUT：命令超时
	// <li> PARTIAL_FAILED：命令部分失败
	InvocationStatus *string `json:"InvocationStatus,omitnil" name:"InvocationStatus"`

	// 执行任务信息列表。
	InvocationTaskBasicInfoSet []*InvocationTaskBasicInfo `json:"InvocationTaskBasicInfoSet,omitnil" name:"InvocationTaskBasicInfoSet"`

	// 执行活动描述。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 执行活动开始时间。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 执行活动结束时间。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 执行活动创建时间。
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// 执行活动更新时间。
	UpdatedTime *string `json:"UpdatedTime,omitnil" name:"UpdatedTime"`

	// 自定义参数取值。
	Parameters *string `json:"Parameters,omitnil" name:"Parameters"`

	// 自定义参数的默认取值。
	DefaultParameters *string `json:"DefaultParameters,omitnil" name:"DefaultParameters"`

	// 执行命令的实例类型，取值范围：CVM、LIGHTHOUSE。
	InstanceKind *string `json:"InstanceKind,omitnil" name:"InstanceKind"`

	// 在实例上执行命令时使用的用户名。
	Username *string `json:"Username,omitnil" name:"Username"`

	// 调用来源。
	InvocationSource *string `json:"InvocationSource,omitnil" name:"InvocationSource"`

	// base64编码的命令内容
	CommandContent *string `json:"CommandContent,omitnil" name:"CommandContent"`

	// 命令类型
	CommandType *string `json:"CommandType,omitnil" name:"CommandType"`

	// 执行命令过期时间， 单位秒
	Timeout *uint64 `json:"Timeout,omitnil" name:"Timeout"`

	// 执行命令的工作路径
	WorkingDirectory *string `json:"WorkingDirectory,omitnil" name:"WorkingDirectory"`

	// 日志上传的cos bucket 地址。
	OutputCOSBucketUrl *string `json:"OutputCOSBucketUrl,omitnil" name:"OutputCOSBucketUrl"`

	// 日志在cos bucket中的目录。
	OutputCOSKeyPrefix *string `json:"OutputCOSKeyPrefix,omitnil" name:"OutputCOSKeyPrefix"`
}

type InvocationTask struct {
	// 执行活动ID。
	InvocationId *string `json:"InvocationId,omitnil" name:"InvocationId"`

	// 执行任务ID。
	InvocationTaskId *string `json:"InvocationTaskId,omitnil" name:"InvocationTaskId"`

	// 命令ID。
	CommandId *string `json:"CommandId,omitnil" name:"CommandId"`

	// 执行任务状态。取值范围：
	// <li> PENDING：等待下发 
	// <li> DELIVERING：下发中
	// <li> DELIVER_DELAYED：延时下发 
	// <li> DELIVER_FAILED：下发失败
	// <li> START_FAILED：命令启动失败
	// <li> RUNNING：命令运行中
	// <li> SUCCESS：命令成功
	// <li> FAILED：命令执行失败，执行完退出码不为 0
	// <li> TIMEOUT：命令超时
	// <li> TASK_TIMEOUT：执行任务超时
	// <li> CANCELLING：取消中
	// <li> CANCELLED：已取消（命令启动前就被取消）
	// <li> TERMINATED：已中止（命令执行期间被取消）
	TaskStatus *string `json:"TaskStatus,omitnil" name:"TaskStatus"`

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 执行结果。
	TaskResult *TaskResult `json:"TaskResult,omitnil" name:"TaskResult"`

	// 执行任务开始时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 执行任务结束时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 创建时间。
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// 更新时间。
	UpdatedTime *string `json:"UpdatedTime,omitnil" name:"UpdatedTime"`

	// 执行任务所执行的命令详情。
	CommandDocument *CommandDocument `json:"CommandDocument,omitnil" name:"CommandDocument"`

	// 执行任务失败时的错误信息。
	ErrorInfo *string `json:"ErrorInfo,omitnil" name:"ErrorInfo"`

	// 调用来源。
	InvocationSource *string `json:"InvocationSource,omitnil" name:"InvocationSource"`
}

type InvocationTaskBasicInfo struct {
	// 执行任务ID。
	InvocationTaskId *string `json:"InvocationTaskId,omitnil" name:"InvocationTaskId"`

	// 执行任务状态。取值范围：
	// <li> PENDING：等待下发 
	// <li> DELIVERING：下发中
	// <li> DELIVER_DELAYED：延时下发 
	// <li> DELIVER_FAILED：下发失败
	// <li> START_FAILED：命令启动失败
	// <li> RUNNING：命令运行中
	// <li> SUCCESS：命令成功
	// <li> FAILED：命令执行失败，执行完退出码不为 0
	// <li> TIMEOUT：命令超时
	// <li> TASK_TIMEOUT：执行任务超时
	// <li> CANCELLING：取消中
	// <li> CANCELLED：已取消（命令启动前就被取消）
	// <li> TERMINATED：已中止（命令执行期间被取消）
	TaskStatus *string `json:"TaskStatus,omitnil" name:"TaskStatus"`

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

// Predefined struct for user
type InvokeCommandRequestParams struct {
	// 待触发的命令ID。
	CommandId *string `json:"CommandId,omitnil" name:"CommandId"`

	// 待执行命令的实例ID列表，上限200。
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Command 的自定义参数。字段类型为json encoded string。如：{\"varA\": \"222\"}。
	// key为自定义参数名称，value为该参数的默认取值。kv均为字符串型。
	// 如果未提供该参数取值，将使用 Command 的 DefaultParameters 进行替换。
	// 自定义参数最多20个。
	// 自定义参数名称需符合以下规范：字符数目上限64，可选范围【a-zA-Z0-9-_】。
	Parameters *string `json:"Parameters,omitnil" name:"Parameters"`

	// 在 CVM 或 Lighthouse 实例中执行命令的用户名称。
	// 使用最小权限执行命令是权限管理的最佳实践，建议您以普通用户身份执行云助手命令。若不填，默认以 Command 配置的 Username 执行。
	Username *string `json:"Username,omitnil" name:"Username"`

	// 命令执行路径, 默认以Command配置的WorkingDirectory执行。
	WorkingDirectory *string `json:"WorkingDirectory,omitnil" name:"WorkingDirectory"`

	// 命令超时时间，取值范围[1, 86400]。默认以Command配置的Timeout执行。
	Timeout *uint64 `json:"Timeout,omitnil" name:"Timeout"`

	// 指定日志上传的cos bucket 地址，必须以https开头，如 https://BucketName-123454321.cos.ap-beijing.myqcloud.com。
	OutputCOSBucketUrl *string `json:"OutputCOSBucketUrl,omitnil" name:"OutputCOSBucketUrl"`

	// 指定日志在cos bucket中的目录，目录命名有如下规则：
	// 1. 可用数字、中英文和可见字符的组合，长度最多为60。
	// 2. 用 / 分割路径，可快速创建子目录。
	// 3. 不允许连续 / ；不允许以 / 开头；不允许以..作为文件夹名称。
	OutputCOSKeyPrefix *string `json:"OutputCOSKeyPrefix,omitnil" name:"OutputCOSKeyPrefix"`
}

type InvokeCommandRequest struct {
	*tchttp.BaseRequest
	
	// 待触发的命令ID。
	CommandId *string `json:"CommandId,omitnil" name:"CommandId"`

	// 待执行命令的实例ID列表，上限200。
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Command 的自定义参数。字段类型为json encoded string。如：{\"varA\": \"222\"}。
	// key为自定义参数名称，value为该参数的默认取值。kv均为字符串型。
	// 如果未提供该参数取值，将使用 Command 的 DefaultParameters 进行替换。
	// 自定义参数最多20个。
	// 自定义参数名称需符合以下规范：字符数目上限64，可选范围【a-zA-Z0-9-_】。
	Parameters *string `json:"Parameters,omitnil" name:"Parameters"`

	// 在 CVM 或 Lighthouse 实例中执行命令的用户名称。
	// 使用最小权限执行命令是权限管理的最佳实践，建议您以普通用户身份执行云助手命令。若不填，默认以 Command 配置的 Username 执行。
	Username *string `json:"Username,omitnil" name:"Username"`

	// 命令执行路径, 默认以Command配置的WorkingDirectory执行。
	WorkingDirectory *string `json:"WorkingDirectory,omitnil" name:"WorkingDirectory"`

	// 命令超时时间，取值范围[1, 86400]。默认以Command配置的Timeout执行。
	Timeout *uint64 `json:"Timeout,omitnil" name:"Timeout"`

	// 指定日志上传的cos bucket 地址，必须以https开头，如 https://BucketName-123454321.cos.ap-beijing.myqcloud.com。
	OutputCOSBucketUrl *string `json:"OutputCOSBucketUrl,omitnil" name:"OutputCOSBucketUrl"`

	// 指定日志在cos bucket中的目录，目录命名有如下规则：
	// 1. 可用数字、中英文和可见字符的组合，长度最多为60。
	// 2. 用 / 分割路径，可快速创建子目录。
	// 3. 不允许连续 / ；不允许以 / 开头；不允许以..作为文件夹名称。
	OutputCOSKeyPrefix *string `json:"OutputCOSKeyPrefix,omitnil" name:"OutputCOSKeyPrefix"`
}

func (r *InvokeCommandRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InvokeCommandRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CommandId")
	delete(f, "InstanceIds")
	delete(f, "Parameters")
	delete(f, "Username")
	delete(f, "WorkingDirectory")
	delete(f, "Timeout")
	delete(f, "OutputCOSBucketUrl")
	delete(f, "OutputCOSKeyPrefix")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InvokeCommandRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InvokeCommandResponseParams struct {
	// 执行活动ID。
	InvocationId *string `json:"InvocationId,omitnil" name:"InvocationId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type InvokeCommandResponse struct {
	*tchttp.BaseResponse
	Response *InvokeCommandResponseParams `json:"Response"`
}

func (r *InvokeCommandResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InvokeCommandResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Invoker struct {
	// 执行器ID。
	InvokerId *string `json:"InvokerId,omitnil" name:"InvokerId"`

	// 执行器名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 执行器类型。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 命令ID。
	CommandId *string `json:"CommandId,omitnil" name:"CommandId"`

	// 用户名。
	Username *string `json:"Username,omitnil" name:"Username"`

	// 自定义参数。
	Parameters *string `json:"Parameters,omitnil" name:"Parameters"`

	// 实例ID列表。
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// 执行器是否启用。
	Enable *bool `json:"Enable,omitnil" name:"Enable"`

	// 执行器周期计划。周期执行器会返回此字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScheduleSettings *ScheduleSettings `json:"ScheduleSettings,omitnil" name:"ScheduleSettings"`

	// 创建时间。
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// 修改时间。
	UpdatedTime *string `json:"UpdatedTime,omitnil" name:"UpdatedTime"`
}

type InvokerRecord struct {
	// 执行器ID。
	InvokerId *string `json:"InvokerId,omitnil" name:"InvokerId"`

	// 执行时间。
	InvokeTime *string `json:"InvokeTime,omitnil" name:"InvokeTime"`

	// 执行原因。
	Reason *string `json:"Reason,omitnil" name:"Reason"`

	// 命令执行ID。
	InvocationId *string `json:"InvocationId,omitnil" name:"InvocationId"`

	// 触发结果。
	Result *string `json:"Result,omitnil" name:"Result"`
}

// Predefined struct for user
type ModifyCommandRequestParams struct {
	// 命令ID。
	CommandId *string `json:"CommandId,omitnil" name:"CommandId"`

	// 命令名称。名称仅支持中文、英文、数字、下划线、分隔符"-"、小数点，最大长度不能超60个字节。
	CommandName *string `json:"CommandName,omitnil" name:"CommandName"`

	// 命令描述。不超过120字符。
	Description *string `json:"Description,omitnil" name:"Description"`

	// Base64编码后的命令内容，长度不可超过64KB。
	Content *string `json:"Content,omitnil" name:"Content"`

	// 命令类型，目前支持取值：SHELL、POWERSHELL。
	CommandType *string `json:"CommandType,omitnil" name:"CommandType"`

	// 命令执行路径。
	WorkingDirectory *string `json:"WorkingDirectory,omitnil" name:"WorkingDirectory"`

	// 命令超时时间。取值范围[1, 86400]。
	Timeout *uint64 `json:"Timeout,omitnil" name:"Timeout"`

	// 启用自定义参数功能时，自定义参数的默认取值。字段类型为json encoded string。如：{\"varA\": \"222\"}。
	// 采取整体全覆盖式修改，即修改时必须提供所有新默认值。
	// 必须 Command 的 EnableParameter 为 true 时，才允许修改这个值。
	// key为自定义参数名称，value为该参数的默认取值。kv均为字符串型。
	// 自定义参数最多20个。
	// 自定义参数名称需符合以下规范：字符数目上限64，可选范围【a-zA-Z0-9-_】。
	DefaultParameters *string `json:"DefaultParameters,omitnil" name:"DefaultParameters"`

	// 自定义参数数组。
	// 如果InvokeCommand时未提供参数取值，将使用这里的默认值进行替换。
	// 自定义参数最多20个。
	DefaultParameterConfs []*DefaultParameterConf `json:"DefaultParameterConfs,omitnil" name:"DefaultParameterConfs"`

	// 在 CVM 或 Lighthouse 实例中执行命令的用户名称。
	// 使用最小权限执行命令是权限管理的最佳实践，建议您以普通用户身份执行云助手命令。
	Username *string `json:"Username,omitnil" name:"Username"`

	// 指定日志上传的cos bucket 地址，必须以https开头，如 https://BucketName-123454321.cos.ap-beijing.myqcloud.com。
	OutputCOSBucketUrl *string `json:"OutputCOSBucketUrl,omitnil" name:"OutputCOSBucketUrl"`

	// 指定日志在cos bucket中的目录，目录命名有如下规则：
	// 1. 可用数字、中英文和可见字符的组合，长度最多为60。
	// 2. 用 / 分割路径，可快速创建子目录。
	// 3. 不允许连续 / ；不允许以 / 开头；不允许以..作为文件夹名称。
	OutputCOSKeyPrefix *string `json:"OutputCOSKeyPrefix,omitnil" name:"OutputCOSKeyPrefix"`
}

type ModifyCommandRequest struct {
	*tchttp.BaseRequest
	
	// 命令ID。
	CommandId *string `json:"CommandId,omitnil" name:"CommandId"`

	// 命令名称。名称仅支持中文、英文、数字、下划线、分隔符"-"、小数点，最大长度不能超60个字节。
	CommandName *string `json:"CommandName,omitnil" name:"CommandName"`

	// 命令描述。不超过120字符。
	Description *string `json:"Description,omitnil" name:"Description"`

	// Base64编码后的命令内容，长度不可超过64KB。
	Content *string `json:"Content,omitnil" name:"Content"`

	// 命令类型，目前支持取值：SHELL、POWERSHELL。
	CommandType *string `json:"CommandType,omitnil" name:"CommandType"`

	// 命令执行路径。
	WorkingDirectory *string `json:"WorkingDirectory,omitnil" name:"WorkingDirectory"`

	// 命令超时时间。取值范围[1, 86400]。
	Timeout *uint64 `json:"Timeout,omitnil" name:"Timeout"`

	// 启用自定义参数功能时，自定义参数的默认取值。字段类型为json encoded string。如：{\"varA\": \"222\"}。
	// 采取整体全覆盖式修改，即修改时必须提供所有新默认值。
	// 必须 Command 的 EnableParameter 为 true 时，才允许修改这个值。
	// key为自定义参数名称，value为该参数的默认取值。kv均为字符串型。
	// 自定义参数最多20个。
	// 自定义参数名称需符合以下规范：字符数目上限64，可选范围【a-zA-Z0-9-_】。
	DefaultParameters *string `json:"DefaultParameters,omitnil" name:"DefaultParameters"`

	// 自定义参数数组。
	// 如果InvokeCommand时未提供参数取值，将使用这里的默认值进行替换。
	// 自定义参数最多20个。
	DefaultParameterConfs []*DefaultParameterConf `json:"DefaultParameterConfs,omitnil" name:"DefaultParameterConfs"`

	// 在 CVM 或 Lighthouse 实例中执行命令的用户名称。
	// 使用最小权限执行命令是权限管理的最佳实践，建议您以普通用户身份执行云助手命令。
	Username *string `json:"Username,omitnil" name:"Username"`

	// 指定日志上传的cos bucket 地址，必须以https开头，如 https://BucketName-123454321.cos.ap-beijing.myqcloud.com。
	OutputCOSBucketUrl *string `json:"OutputCOSBucketUrl,omitnil" name:"OutputCOSBucketUrl"`

	// 指定日志在cos bucket中的目录，目录命名有如下规则：
	// 1. 可用数字、中英文和可见字符的组合，长度最多为60。
	// 2. 用 / 分割路径，可快速创建子目录。
	// 3. 不允许连续 / ；不允许以 / 开头；不允许以..作为文件夹名称。
	OutputCOSKeyPrefix *string `json:"OutputCOSKeyPrefix,omitnil" name:"OutputCOSKeyPrefix"`
}

func (r *ModifyCommandRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCommandRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CommandId")
	delete(f, "CommandName")
	delete(f, "Description")
	delete(f, "Content")
	delete(f, "CommandType")
	delete(f, "WorkingDirectory")
	delete(f, "Timeout")
	delete(f, "DefaultParameters")
	delete(f, "DefaultParameterConfs")
	delete(f, "Username")
	delete(f, "OutputCOSBucketUrl")
	delete(f, "OutputCOSKeyPrefix")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCommandRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCommandResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyCommandResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCommandResponseParams `json:"Response"`
}

func (r *ModifyCommandResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCommandResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInvokerRequestParams struct {
	// 待修改的执行器ID。
	InvokerId *string `json:"InvokerId,omitnil" name:"InvokerId"`

	// 待修改的执行器名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 执行器类型，当前仅支持周期类型执行器，取值：`SCHEDULE` 。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 待修改的命令ID。
	CommandId *string `json:"CommandId,omitnil" name:"CommandId"`

	// 待修改的用户名。
	Username *string `json:"Username,omitnil" name:"Username"`

	// 待修改的自定义参数。
	Parameters *string `json:"Parameters,omitnil" name:"Parameters"`

	// 待修改的实例ID列表。列表长度上限100。
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// 待修改的周期执行器设置。
	ScheduleSettings *ScheduleSettings `json:"ScheduleSettings,omitnil" name:"ScheduleSettings"`
}

type ModifyInvokerRequest struct {
	*tchttp.BaseRequest
	
	// 待修改的执行器ID。
	InvokerId *string `json:"InvokerId,omitnil" name:"InvokerId"`

	// 待修改的执行器名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 执行器类型，当前仅支持周期类型执行器，取值：`SCHEDULE` 。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 待修改的命令ID。
	CommandId *string `json:"CommandId,omitnil" name:"CommandId"`

	// 待修改的用户名。
	Username *string `json:"Username,omitnil" name:"Username"`

	// 待修改的自定义参数。
	Parameters *string `json:"Parameters,omitnil" name:"Parameters"`

	// 待修改的实例ID列表。列表长度上限100。
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// 待修改的周期执行器设置。
	ScheduleSettings *ScheduleSettings `json:"ScheduleSettings,omitnil" name:"ScheduleSettings"`
}

func (r *ModifyInvokerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInvokerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InvokerId")
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "CommandId")
	delete(f, "Username")
	delete(f, "Parameters")
	delete(f, "InstanceIds")
	delete(f, "ScheduleSettings")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInvokerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInvokerResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyInvokerResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInvokerResponseParams `json:"Response"`
}

func (r *ModifyInvokerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInvokerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRegisterInstanceRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 实例名。
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`
}

type ModifyRegisterInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 实例名。
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`
}

func (r *ModifyRegisterInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRegisterInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRegisterInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRegisterInstanceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyRegisterInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRegisterInstanceResponseParams `json:"Response"`
}

func (r *ModifyRegisterInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRegisterInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PreviewReplacedCommandContentRequestParams struct {
	// 本次预览采用的自定义参数。字段类型为 json encoded string，如：{\"varA\": \"222\"}。
	// key 为自定义参数名称，value 为该参数的取值。kv 均为字符串型。
	// 自定义参数最多 20 个。
	// 自定义参数名称需符合以下规范：字符数目上限 64，可选范围【a-zA-Z0-9-_】。
	// 如果将预览的 CommandId 设置过 DefaultParameters，本参数可以为空。
	Parameters *string `json:"Parameters,omitnil" name:"Parameters"`

	// 要进行替换预览的命令，如果有设置过 DefaultParameters，会与 Parameters 进行叠加，后者覆盖前者。
	// CommandId 与 Content，必须且只能提供一个。
	CommandId *string `json:"CommandId,omitnil" name:"CommandId"`

	// 要预览的命令内容，经 Base64 编码，长度不可超过 64KB。
	// CommandId 与 Content，必须且只能提供一个。
	Content *string `json:"Content,omitnil" name:"Content"`
}

type PreviewReplacedCommandContentRequest struct {
	*tchttp.BaseRequest
	
	// 本次预览采用的自定义参数。字段类型为 json encoded string，如：{\"varA\": \"222\"}。
	// key 为自定义参数名称，value 为该参数的取值。kv 均为字符串型。
	// 自定义参数最多 20 个。
	// 自定义参数名称需符合以下规范：字符数目上限 64，可选范围【a-zA-Z0-9-_】。
	// 如果将预览的 CommandId 设置过 DefaultParameters，本参数可以为空。
	Parameters *string `json:"Parameters,omitnil" name:"Parameters"`

	// 要进行替换预览的命令，如果有设置过 DefaultParameters，会与 Parameters 进行叠加，后者覆盖前者。
	// CommandId 与 Content，必须且只能提供一个。
	CommandId *string `json:"CommandId,omitnil" name:"CommandId"`

	// 要预览的命令内容，经 Base64 编码，长度不可超过 64KB。
	// CommandId 与 Content，必须且只能提供一个。
	Content *string `json:"Content,omitnil" name:"Content"`
}

func (r *PreviewReplacedCommandContentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PreviewReplacedCommandContentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Parameters")
	delete(f, "CommandId")
	delete(f, "Content")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PreviewReplacedCommandContentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PreviewReplacedCommandContentResponseParams struct {
	// 自定义参数替换后的，经Base64编码的命令内容。
	ReplacedContent *string `json:"ReplacedContent,omitnil" name:"ReplacedContent"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type PreviewReplacedCommandContentResponse struct {
	*tchttp.BaseResponse
	Response *PreviewReplacedCommandContentResponseParams `json:"Response"`
}

func (r *PreviewReplacedCommandContentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PreviewReplacedCommandContentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionInfo struct {
	// 地域名称，例如，ap-guangzhou
	Region *string `json:"Region,omitnil" name:"Region"`

	// 地域描述，例如: 广州
	RegionName *string `json:"RegionName,omitnil" name:"RegionName"`

	// 地域是否可用状态，AVAILABLE 代表可用
	RegionState *string `json:"RegionState,omitnil" name:"RegionState"`
}

type RegisterCodeInfo struct {
	// 注册码ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegisterCodeId *string `json:"RegisterCodeId,omitnil" name:"RegisterCodeId"`

	// 注册码描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 注册实例名称前缀。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceNamePrefix *string `json:"InstanceNamePrefix,omitnil" name:"InstanceNamePrefix"`

	// 该注册码允许注册的实列数目。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegisterLimit *int64 `json:"RegisterLimit,omitnil" name:"RegisterLimit"`

	// 该注册码的过期时间，按照 ISO8601 标准表示，并且使用 UTC 时间。 
	// 格式为： YYYY-MM-DDThh:mm:ssZ。
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpiredTime *string `json:"ExpiredTime,omitnil" name:"ExpiredTime"`

	// 该注册码限制tat_agent只能从IpAddressRange所描述公网出口进行注册。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpAddressRange *string `json:"IpAddressRange,omitnil" name:"IpAddressRange"`

	// 该注册码是否可用。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enabled *bool `json:"Enabled,omitnil" name:"Enabled"`

	// 该注册码已注册数目。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegisteredCount *int64 `json:"RegisteredCount,omitnil" name:"RegisteredCount"`

	// 注册码创建时间，按照 ISO8601 标准表示，并且使用 UTC 时间。 
	// 格式为： YYYY-MM-DDThh:mm:ssZ。
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// 注册码最近一次更新时间，按照 ISO8601 标准表示，并且使用 UTC 时间。 
	// 格式为： YYYY-MM-DDThh:mm:ssZ。
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedTime *string `json:"UpdatedTime,omitnil" name:"UpdatedTime"`
}

type RegisterInstanceInfo struct {
	// 注册码ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegisterCodeId *string `json:"RegisterCodeId,omitnil" name:"RegisterCodeId"`

	// 实例ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 实例名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 机器ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineId *string `json:"MachineId,omitnil" name:"MachineId"`

	// 系统名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SystemName *string `json:"SystemName,omitnil" name:"SystemName"`

	// 主机IP。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostName *string `json:"HostName,omitnil" name:"HostName"`

	// 内网IP。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocalIp *string `json:"LocalIp,omitnil" name:"LocalIp"`

	// 公钥。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicKey *string `json:"PublicKey,omitnil" name:"PublicKey"`

	// 托管状态。
	// 返回Online表示实例正在托管，返回Offline表示实例未托管。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// 上次更新时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedTime *string `json:"UpdatedTime,omitnil" name:"UpdatedTime"`
}

// Predefined struct for user
type RunCommandRequestParams struct {
	// Base64编码后的命令内容，长度不可超过64KB。
	Content *string `json:"Content,omitnil" name:"Content"`

	// 待执行命令的实例ID列表，上限200。支持实例类型：
	// <li> CVM
	// <li> LIGHTHOUSE
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// 命令名称。名称仅支持中文、英文、数字、下划线、分隔符"-"、小数点，最大长度不能超60个字节。
	CommandName *string `json:"CommandName,omitnil" name:"CommandName"`

	// 命令描述。不超过120字符。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 命令类型，目前支持取值：SHELL、POWERSHELL。默认：SHELL。
	CommandType *string `json:"CommandType,omitnil" name:"CommandType"`

	// 命令执行路径，对于 SHELL 命令默认为 /root，对于 POWERSHELL 命令默认为 C:\Program Files\qcloud\tat_agent\workdir。
	WorkingDirectory *string `json:"WorkingDirectory,omitnil" name:"WorkingDirectory"`

	// 命令超时时间，默认60秒。取值范围[1, 86400]。
	Timeout *uint64 `json:"Timeout,omitnil" name:"Timeout"`

	// 是否保存命令，取值范围：
	// <li> true：保存
	// <li> false：不保存
	// 默认为 false。
	SaveCommand *bool `json:"SaveCommand,omitnil" name:"SaveCommand"`

	// 是否启用自定义参数功能。
	// 一旦创建，此值不提供修改。
	// 取值范围：
	// <li> true：启用
	// <li> false：不启用
	// 默认值：false。
	EnableParameter *bool `json:"EnableParameter,omitnil" name:"EnableParameter"`

	// 启用自定义参数功能时，自定义参数的默认取值。字段类型为json encoded string。如：{\"varA\": \"222\"}。
	// key为自定义参数名称，value为该参数的默认取值。kv均为字符串型。
	// 如果 Parameters 未提供，将使用这里的默认值进行替换。
	// 自定义参数最多20个。
	// 自定义参数名称需符合以下规范：字符数目上限64，可选范围【a-zA-Z0-9-_】。
	DefaultParameters *string `json:"DefaultParameters,omitnil" name:"DefaultParameters"`

	// 自定义参数数组。 如果 Parameters 未提供，将使用这里的默认值进行替换。 自定义参数最多20个。
	DefaultParameterConfs []*DefaultParameterConf `json:"DefaultParameterConfs,omitnil" name:"DefaultParameterConfs"`

	// Command 的自定义参数。字段类型为json encoded string。如：{\"varA\": \"222\"}。
	// key为自定义参数名称，value为该参数的默认取值。kv均为字符串型。
	// 如果未提供该参数取值，将使用 DefaultParameters 进行替换。
	// 自定义参数最多20个。
	// 自定义参数名称需符合以下规范：字符数目上限64，可选范围【a-zA-Z0-9-_】。
	Parameters *string `json:"Parameters,omitnil" name:"Parameters"`

	// 如果保存命令，可为命令设置标签。列表长度不超过10。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// 在 CVM 或 Lighthouse 实例中执行命令的用户名称。
	// 使用最小权限执行命令是权限管理的最佳实践，建议您以普通用户身份执行云助手命令。默认情况下，在 Linux 实例中以 root 用户执行命令；在Windows 实例中以 System 用户执行命令。
	Username *string `json:"Username,omitnil" name:"Username"`

	// 指定日志上传的cos bucket 地址，必须以https开头，如 https://BucketName-123454321.cos.ap-beijing.myqcloud.com。
	OutputCOSBucketUrl *string `json:"OutputCOSBucketUrl,omitnil" name:"OutputCOSBucketUrl"`

	// 指定日志在cos bucket中的目录，目录命名有如下规则：
	// 1. 可用数字、中英文和可见字符的组合，长度最多为60。
	// 2. 用 / 分割路径，可快速创建子目录。
	// 3. 不允许连续 / ；不允许以 / 开头；不允许以..作为文件夹名称。
	OutputCOSKeyPrefix *string `json:"OutputCOSKeyPrefix,omitnil" name:"OutputCOSKeyPrefix"`
}

type RunCommandRequest struct {
	*tchttp.BaseRequest
	
	// Base64编码后的命令内容，长度不可超过64KB。
	Content *string `json:"Content,omitnil" name:"Content"`

	// 待执行命令的实例ID列表，上限200。支持实例类型：
	// <li> CVM
	// <li> LIGHTHOUSE
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// 命令名称。名称仅支持中文、英文、数字、下划线、分隔符"-"、小数点，最大长度不能超60个字节。
	CommandName *string `json:"CommandName,omitnil" name:"CommandName"`

	// 命令描述。不超过120字符。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 命令类型，目前支持取值：SHELL、POWERSHELL。默认：SHELL。
	CommandType *string `json:"CommandType,omitnil" name:"CommandType"`

	// 命令执行路径，对于 SHELL 命令默认为 /root，对于 POWERSHELL 命令默认为 C:\Program Files\qcloud\tat_agent\workdir。
	WorkingDirectory *string `json:"WorkingDirectory,omitnil" name:"WorkingDirectory"`

	// 命令超时时间，默认60秒。取值范围[1, 86400]。
	Timeout *uint64 `json:"Timeout,omitnil" name:"Timeout"`

	// 是否保存命令，取值范围：
	// <li> true：保存
	// <li> false：不保存
	// 默认为 false。
	SaveCommand *bool `json:"SaveCommand,omitnil" name:"SaveCommand"`

	// 是否启用自定义参数功能。
	// 一旦创建，此值不提供修改。
	// 取值范围：
	// <li> true：启用
	// <li> false：不启用
	// 默认值：false。
	EnableParameter *bool `json:"EnableParameter,omitnil" name:"EnableParameter"`

	// 启用自定义参数功能时，自定义参数的默认取值。字段类型为json encoded string。如：{\"varA\": \"222\"}。
	// key为自定义参数名称，value为该参数的默认取值。kv均为字符串型。
	// 如果 Parameters 未提供，将使用这里的默认值进行替换。
	// 自定义参数最多20个。
	// 自定义参数名称需符合以下规范：字符数目上限64，可选范围【a-zA-Z0-9-_】。
	DefaultParameters *string `json:"DefaultParameters,omitnil" name:"DefaultParameters"`

	// 自定义参数数组。 如果 Parameters 未提供，将使用这里的默认值进行替换。 自定义参数最多20个。
	DefaultParameterConfs []*DefaultParameterConf `json:"DefaultParameterConfs,omitnil" name:"DefaultParameterConfs"`

	// Command 的自定义参数。字段类型为json encoded string。如：{\"varA\": \"222\"}。
	// key为自定义参数名称，value为该参数的默认取值。kv均为字符串型。
	// 如果未提供该参数取值，将使用 DefaultParameters 进行替换。
	// 自定义参数最多20个。
	// 自定义参数名称需符合以下规范：字符数目上限64，可选范围【a-zA-Z0-9-_】。
	Parameters *string `json:"Parameters,omitnil" name:"Parameters"`

	// 如果保存命令，可为命令设置标签。列表长度不超过10。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// 在 CVM 或 Lighthouse 实例中执行命令的用户名称。
	// 使用最小权限执行命令是权限管理的最佳实践，建议您以普通用户身份执行云助手命令。默认情况下，在 Linux 实例中以 root 用户执行命令；在Windows 实例中以 System 用户执行命令。
	Username *string `json:"Username,omitnil" name:"Username"`

	// 指定日志上传的cos bucket 地址，必须以https开头，如 https://BucketName-123454321.cos.ap-beijing.myqcloud.com。
	OutputCOSBucketUrl *string `json:"OutputCOSBucketUrl,omitnil" name:"OutputCOSBucketUrl"`

	// 指定日志在cos bucket中的目录，目录命名有如下规则：
	// 1. 可用数字、中英文和可见字符的组合，长度最多为60。
	// 2. 用 / 分割路径，可快速创建子目录。
	// 3. 不允许连续 / ；不允许以 / 开头；不允许以..作为文件夹名称。
	OutputCOSKeyPrefix *string `json:"OutputCOSKeyPrefix,omitnil" name:"OutputCOSKeyPrefix"`
}

func (r *RunCommandRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunCommandRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Content")
	delete(f, "InstanceIds")
	delete(f, "CommandName")
	delete(f, "Description")
	delete(f, "CommandType")
	delete(f, "WorkingDirectory")
	delete(f, "Timeout")
	delete(f, "SaveCommand")
	delete(f, "EnableParameter")
	delete(f, "DefaultParameters")
	delete(f, "DefaultParameterConfs")
	delete(f, "Parameters")
	delete(f, "Tags")
	delete(f, "Username")
	delete(f, "OutputCOSBucketUrl")
	delete(f, "OutputCOSKeyPrefix")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunCommandRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunCommandResponseParams struct {
	// 命令ID。
	CommandId *string `json:"CommandId,omitnil" name:"CommandId"`

	// 执行活动ID。
	InvocationId *string `json:"InvocationId,omitnil" name:"InvocationId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RunCommandResponse struct {
	*tchttp.BaseResponse
	Response *RunCommandResponseParams `json:"Response"`
}

func (r *RunCommandResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunCommandResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScheduleSettings struct {
	// 执行策略：
	// <br><li>ONCE：单次执行
	// <br><li>RECURRENCE：周期执行
	Policy *string `json:"Policy,omitnil" name:"Policy"`

	// 触发 Crontab 表达式。Policy 为 RECURRENCE 时，需要指定此字段。Crontab 按北京时间解析。
	Recurrence *string `json:"Recurrence,omitnil" name:"Recurrence"`

	// 执行器下次执行时间。Policy 为 ONCE 时，需要指定此字段。
	InvokeTime *string `json:"InvokeTime,omitnil" name:"InvokeTime"`
}

type Tag struct {
	// 标签键。
	Key *string `json:"Key,omitnil" name:"Key"`

	// 标签值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type TaskResult struct {
	// 命令执行ExitCode。
	ExitCode *int64 `json:"ExitCode,omitnil" name:"ExitCode"`

	// Base64编码后的命令输出。最大长度24KB。
	Output *string `json:"Output,omitnil" name:"Output"`

	// 命令执行开始时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecStartTime *string `json:"ExecStartTime,omitnil" name:"ExecStartTime"`

	// 命令执行结束时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecEndTime *string `json:"ExecEndTime,omitnil" name:"ExecEndTime"`

	// 命令最终输出被截断的字节数。
	Dropped *uint64 `json:"Dropped,omitnil" name:"Dropped"`

	// 日志在cos中的地址
	OutputUrl *string `json:"OutputUrl,omitnil" name:"OutputUrl"`

	// 日志上传cos的错误信息。
	OutputUploadCOSErrorInfo *string `json:"OutputUploadCOSErrorInfo,omitnil" name:"OutputUploadCOSErrorInfo"`
}