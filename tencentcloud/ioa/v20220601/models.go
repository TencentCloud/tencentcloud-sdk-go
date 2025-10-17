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

package v20220601

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AggrCategorySoftDetailRow struct {
	// ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 软件名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 盗版风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	PiracyRisk *int64 `json:"PiracyRisk,omitnil,omitempty" name:"PiracyRisk"`

	// 系统平台
	// 注意：此字段可能返回 null，表示取不到有效值。
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`

	// 企业名
	// 注意：此字段可能返回 null，表示取不到有效值。
	CorpName *string `json:"CorpName,omitnil,omitempty" name:"CorpName"`

	// 安装设备数量(只支持32位)
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstalledDeviceNum *int64 `json:"InstalledDeviceNum,omitnil,omitempty" name:"InstalledDeviceNum"`

	// 盗版安装设备数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PiracyInstalledDeviceNum *int64 `json:"PiracyInstalledDeviceNum,omitnil,omitempty" name:"PiracyInstalledDeviceNum"`

	// 已安装用户数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstalledUserNum *int64 `json:"InstalledUserNum,omitnil,omitempty" name:"InstalledUserNum"`

	// 盗版软件用户数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PiracyInstalledUserNum *int64 `json:"PiracyInstalledUserNum,omitnil,omitempty" name:"PiracyInstalledUserNum"`

	// 授权总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthNum *int64 `json:"AuthNum,omitnil,omitempty" name:"AuthNum"`

	// 正版率
	// 注意：此字段可能返回 null，表示取不到有效值。
	GenuineRate *float64 `json:"GenuineRate,omitnil,omitempty" name:"GenuineRate"`

	// 有新版本可升级的设备数量
	UpgradableDeviceNum *int64 `json:"UpgradableDeviceNum,omitnil,omitempty" name:"UpgradableDeviceNum"`
}

type AggrSoftDeviceRow struct {
	// 终端名
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 最近登录账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastLoginAccount *string `json:"LastLoginAccount,omitnil,omitempty" name:"LastLoginAccount"`

	// 终端用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceUserName *string `json:"DeviceUserName,omitnil,omitempty" name:"DeviceUserName"`

	// 软件版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 是否盗版
	// 注意：此字段可能返回 null，表示取不到有效值。
	PiracyRisk *int64 `json:"PiracyRisk,omitnil,omitempty" name:"PiracyRisk"`

	// 盗版原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	PiracyReason *string `json:"PiracyReason,omitnil,omitempty" name:"PiracyReason"`

	// 安装时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstallTime *string `json:"InstallTime,omitnil,omitempty" name:"InstallTime"`

	// 用户目录
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserPath *string `json:"UserPath,omitnil,omitempty" name:"UserPath"`

	// 所在分组
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserGroup *string `json:"UserGroup,omitnil,omitempty" name:"UserGroup"`

	// IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	IP *string `json:"IP,omitnil,omitempty" name:"IP"`

	// MAC
	// 注意：此字段可能返回 null，表示取不到有效值。
	MAC *string `json:"MAC,omitnil,omitempty" name:"MAC"`

	// 使用时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	UseTime *int64 `json:"UseTime,omitnil,omitempty" name:"UseTime"`

	// 设备ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceId *int64 `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 软件全名
	// 注意：此字段可能返回 null，表示取不到有效值。
	FullSoftName *string `json:"FullSoftName,omitnil,omitempty" name:"FullSoftName"`

	// 数据ID（唯一）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 该终端此款软件可升级到的目标版本号
	NewVersion *string `json:"NewVersion,omitnil,omitempty" name:"NewVersion"`

	// 该软件对应运营配置的可升级id
	UpgradeSoftId *int64 `json:"UpgradeSoftId,omitnil,omitempty" name:"UpgradeSoftId"`

	// 终端备注名
	RemarkName *string `json:"RemarkName,omitnil,omitempty" name:"RemarkName"`
}

type ComplexRule struct {
	// 简单规则表达式
	SimpleRules []*SimpleRule `json:"SimpleRules,omitnil,omitempty" name:"SimpleRules"`

	// 表达式间逻辑关系
	Relation *string `json:"Relation,omitnil,omitempty" name:"Relation"`
}

type Condition struct {
	// Filters 条件过滤
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// FilterGroups 条件过滤组
	FilterGroups []*FilterGroup `json:"FilterGroups,omitnil,omitempty" name:"FilterGroups"`

	// Sort 排序字段
	Sort *Sort `json:"Sort,omitnil,omitempty" name:"Sort"`

	// PageSize 每页获取数(只支持32位)
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// PageNum 获取第几页(只支持32位)
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`
}

type CreateDLPFileDetectTaskData struct {
	// 任务请求唯一Id
	TaskRequestId []*string `json:"TaskRequestId,omitnil,omitempty" name:"TaskRequestId"`
}

// Predefined struct for user
type CreateDLPFileDetectTaskRequestParams struct {
	// 文件下载Url
	DownloadUrl *string `json:"DownloadUrl,omitnil,omitempty" name:"DownloadUrl"`

	// 文件名
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文件Md5
	FileMd5 *string `json:"FileMd5,omitnil,omitempty" name:"FileMd5"`

	// 负载类型  1 从GroupId中选一节点 鉴定  2使用所有SelectNodeIds节点鉴定
	BalanceType *int64 `json:"BalanceType,omitnil,omitempty" name:"BalanceType"`

	// 管理域实例ID，用于CAM管理域权限分配。若企业未进行管理域的划分，可直接传入根域"1"，此时表示针对当前企业的全部设备和账号进行接口CRUD，具体CRUD的影响范围限制于相应接口的入参。
	DomainInstanceId *string `json:"DomainInstanceId,omitnil,omitempty" name:"DomainInstanceId"`

	// 选中节点唯一Id列表,BalanceType=2时必填
	SelectNodeIds []*string `json:"SelectNodeIds,omitnil,omitempty" name:"SelectNodeIds"`

	// 节点组唯一Id,BalanceType=1时必填
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

type CreateDLPFileDetectTaskRequest struct {
	*tchttp.BaseRequest
	
	// 文件下载Url
	DownloadUrl *string `json:"DownloadUrl,omitnil,omitempty" name:"DownloadUrl"`

	// 文件名
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文件Md5
	FileMd5 *string `json:"FileMd5,omitnil,omitempty" name:"FileMd5"`

	// 负载类型  1 从GroupId中选一节点 鉴定  2使用所有SelectNodeIds节点鉴定
	BalanceType *int64 `json:"BalanceType,omitnil,omitempty" name:"BalanceType"`

	// 管理域实例ID，用于CAM管理域权限分配。若企业未进行管理域的划分，可直接传入根域"1"，此时表示针对当前企业的全部设备和账号进行接口CRUD，具体CRUD的影响范围限制于相应接口的入参。
	DomainInstanceId *string `json:"DomainInstanceId,omitnil,omitempty" name:"DomainInstanceId"`

	// 选中节点唯一Id列表,BalanceType=2时必填
	SelectNodeIds []*string `json:"SelectNodeIds,omitnil,omitempty" name:"SelectNodeIds"`

	// 节点组唯一Id,BalanceType=1时必填
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

func (r *CreateDLPFileDetectTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDLPFileDetectTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DownloadUrl")
	delete(f, "FileName")
	delete(f, "FileMd5")
	delete(f, "BalanceType")
	delete(f, "DomainInstanceId")
	delete(f, "SelectNodeIds")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDLPFileDetectTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDLPFileDetectTaskResponseParams struct {
	// 创建文件鉴定任务数据
	Data *CreateDLPFileDetectTaskData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDLPFileDetectTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateDLPFileDetectTaskResponseParams `json:"Response"`
}

func (r *CreateDLPFileDetectTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDLPFileDetectTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDLPFileDetectionTaskData struct {
	// 提交任务生成的id，也即requestID。用于后续查询
	DLPFileDetectionTaskID *string `json:"DLPFileDetectionTaskID,omitnil,omitempty" name:"DLPFileDetectionTaskID"`
}

// Predefined struct for user
type CreateDLPFileDetectionTaskRequestParams struct {
	// 文件下载链接，要求公网可访问，GET方式访问后为文件
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 文件名，带后缀
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	//  文件md5，传入相同md5会直接使用之前缓存的结果。
	// 
	// > 请注意：不同文件使用相同md5送检，会命中缓存得到旧的检测结果
	FileMd5 *string `json:"FileMd5,omitnil,omitempty" name:"FileMd5"`

	// 管理域实例ID，用于CAM管理域权限分配
	DomainInstanceId *string `json:"DomainInstanceId,omitnil,omitempty" name:"DomainInstanceId"`

	// 回调地址，暂时未使用
	CallBackUrl *string `json:"CallBackUrl,omitnil,omitempty" name:"CallBackUrl"`
}

type CreateDLPFileDetectionTaskRequest struct {
	*tchttp.BaseRequest
	
	// 文件下载链接，要求公网可访问，GET方式访问后为文件
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 文件名，带后缀
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	//  文件md5，传入相同md5会直接使用之前缓存的结果。
	// 
	// > 请注意：不同文件使用相同md5送检，会命中缓存得到旧的检测结果
	FileMd5 *string `json:"FileMd5,omitnil,omitempty" name:"FileMd5"`

	// 管理域实例ID，用于CAM管理域权限分配
	DomainInstanceId *string `json:"DomainInstanceId,omitnil,omitempty" name:"DomainInstanceId"`

	// 回调地址，暂时未使用
	CallBackUrl *string `json:"CallBackUrl,omitnil,omitempty" name:"CallBackUrl"`
}

func (r *CreateDLPFileDetectionTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDLPFileDetectionTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Url")
	delete(f, "FileName")
	delete(f, "FileMd5")
	delete(f, "DomainInstanceId")
	delete(f, "CallBackUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDLPFileDetectionTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDLPFileDetectionTaskResponseParams struct {
	// 创建送检任务响应数据
	Data *CreateDLPFileDetectionTaskData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDLPFileDetectionTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateDLPFileDetectionTaskResponseParams `json:"Response"`
}

func (r *CreateDLPFileDetectionTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDLPFileDetectionTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDeviceTaskRequestParams struct {
	// 管理域实例ID，用于CAM管理域权限分配。若企业未进行管理域的划分，可直接传入根域"1"，此时表示针对当前企业的全部设备和账号进行接口CRUD，具体CRUD的影响范围限制于相应接口的入参。
	DomainInstanceId *string `json:"DomainInstanceId,omitnil,omitempty" name:"DomainInstanceId"`

	// 终端id
	Mid *string `json:"Mid,omitnil,omitempty" name:"Mid"`
}

type CreateDeviceTaskRequest struct {
	*tchttp.BaseRequest
	
	// 管理域实例ID，用于CAM管理域权限分配。若企业未进行管理域的划分，可直接传入根域"1"，此时表示针对当前企业的全部设备和账号进行接口CRUD，具体CRUD的影响范围限制于相应接口的入参。
	DomainInstanceId *string `json:"DomainInstanceId,omitnil,omitempty" name:"DomainInstanceId"`

	// 终端id
	Mid *string `json:"Mid,omitnil,omitempty" name:"Mid"`
}

func (r *CreateDeviceTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDeviceTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainInstanceId")
	delete(f, "Mid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDeviceTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDeviceTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDeviceTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateDeviceTaskResponseParams `json:"Response"`
}

func (r *CreateDeviceTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDeviceTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDeviceVirtualGroupRequestParams struct {
	// 必填，终端自定义分组名
	DeviceVirtualGroupName *string `json:"DeviceVirtualGroupName,omitnil,omitempty" name:"DeviceVirtualGroupName"`

	// 管理域实例ID，用于CAM管理域权限分配。若企业未进行管理域的划分，可直接传入根域"1"，此时表示针对当前企业的全部设备和账号进行接口CRUD，具体CRUD的影响范围限制于相应接口的入参。
	DomainInstanceId *string `json:"DomainInstanceId,omitnil,omitempty" name:"DomainInstanceId"`

	// 详情
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 系统类型（0: win，1：linux，2: mac，4：android，5：ios ； 默认值0）(只支持32位)
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`

	// 分组类型（0:手动分组；非0为自动划分分组；具体枚举值为：1:自动每小时划分分组、2:自动每天划分分组、3:自定义时间划分分组； 默认值0）(只支持32位)
	TimeType *int64 `json:"TimeType,omitnil,omitempty" name:"TimeType"`

	// 选填，TimeType=3时的自动划分时间，其他情况为0（单位min）(只支持32位)
	AutoMinute *int64 `json:"AutoMinute,omitnil,omitempty" name:"AutoMinute"`

	// 选填，手动分组不填，自动划分分组的划分规则数据
	AutoRules *ComplexRule `json:"AutoRules,omitnil,omitempty" name:"AutoRules"`
}

type CreateDeviceVirtualGroupRequest struct {
	*tchttp.BaseRequest
	
	// 必填，终端自定义分组名
	DeviceVirtualGroupName *string `json:"DeviceVirtualGroupName,omitnil,omitempty" name:"DeviceVirtualGroupName"`

	// 管理域实例ID，用于CAM管理域权限分配。若企业未进行管理域的划分，可直接传入根域"1"，此时表示针对当前企业的全部设备和账号进行接口CRUD，具体CRUD的影响范围限制于相应接口的入参。
	DomainInstanceId *string `json:"DomainInstanceId,omitnil,omitempty" name:"DomainInstanceId"`

	// 详情
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 系统类型（0: win，1：linux，2: mac，4：android，5：ios ； 默认值0）(只支持32位)
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`

	// 分组类型（0:手动分组；非0为自动划分分组；具体枚举值为：1:自动每小时划分分组、2:自动每天划分分组、3:自定义时间划分分组； 默认值0）(只支持32位)
	TimeType *int64 `json:"TimeType,omitnil,omitempty" name:"TimeType"`

	// 选填，TimeType=3时的自动划分时间，其他情况为0（单位min）(只支持32位)
	AutoMinute *int64 `json:"AutoMinute,omitnil,omitempty" name:"AutoMinute"`

	// 选填，手动分组不填，自动划分分组的划分规则数据
	AutoRules *ComplexRule `json:"AutoRules,omitnil,omitempty" name:"AutoRules"`
}

func (r *CreateDeviceVirtualGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDeviceVirtualGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceVirtualGroupName")
	delete(f, "DomainInstanceId")
	delete(f, "Description")
	delete(f, "OsType")
	delete(f, "TimeType")
	delete(f, "AutoMinute")
	delete(f, "AutoRules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDeviceVirtualGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDeviceVirtualGroupResponseParams struct {
	// 响应返回的data
	Data *CreateDeviceVirtualGroupRspData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDeviceVirtualGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateDeviceVirtualGroupResponseParams `json:"Response"`
}

func (r *CreateDeviceVirtualGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDeviceVirtualGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDeviceVirtualGroupRspData struct {
	// 返回的自定义分组id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`
}

// Predefined struct for user
type CreatePrivilegeCodeRequestParams struct {
	// 必填；设备唯一标识符;
	Mid *string `json:"Mid,omitnil,omitempty" name:"Mid"`

	// 管理域实例ID，用于CAM管理域权限分配。若企业未进行管理域的划分，可直接传入根域"1"，此时表示针对当前企业的全部设备和账号进行接口CRUD，具体CRUD的影响范围限制于相应接口的入参。
	DomainInstanceId *string `json:"DomainInstanceId,omitnil,omitempty" name:"DomainInstanceId"`

	// 系统类型（0: win，1：linux，2: mac，4：android，5：ios ）；默认值0
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`
}

type CreatePrivilegeCodeRequest struct {
	*tchttp.BaseRequest
	
	// 必填；设备唯一标识符;
	Mid *string `json:"Mid,omitnil,omitempty" name:"Mid"`

	// 管理域实例ID，用于CAM管理域权限分配。若企业未进行管理域的划分，可直接传入根域"1"，此时表示针对当前企业的全部设备和账号进行接口CRUD，具体CRUD的影响范围限制于相应接口的入参。
	DomainInstanceId *string `json:"DomainInstanceId,omitnil,omitempty" name:"DomainInstanceId"`

	// 系统类型（0: win，1：linux，2: mac，4：android，5：ios ）；默认值0
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`
}

func (r *CreatePrivilegeCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrivilegeCodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Mid")
	delete(f, "DomainInstanceId")
	delete(f, "OsType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrivilegeCodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrivilegeCodeResponseParams struct {
	// 业务响应数据
	Data *CreatePrivilegeCodeRspData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePrivilegeCodeResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrivilegeCodeResponseParams `json:"Response"`
}

func (r *CreatePrivilegeCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrivilegeCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePrivilegeCodeRspData struct {
	// 特权码数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`
}

type DescribeAccountGroupsData struct {
	// 账号分组名全路径，点分格式
	NamePath *string `json:"NamePath,omitnil,omitempty" name:"NamePath"`

	// 账号分组ID全路径，数组格式
	IdPathArr []*int64 `json:"IdPathArr,omitnil,omitempty" name:"IdPathArr"`

	// 扩展信息
	ExtraInfo *string `json:"ExtraInfo,omitnil,omitempty" name:"ExtraInfo"`

	// 最后更新时间
	Utime *string `json:"Utime,omitnil,omitempty" name:"Utime"`

	// 父分组ID
	ParentId *int64 `json:"ParentId,omitnil,omitempty" name:"ParentId"`

	// 源账号组织ID。使用第三方导入用户源时，记录该分组在源组织架构下的分组ID
	OrgId *string `json:"OrgId,omitnil,omitempty" name:"OrgId"`

	// 分组名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分组ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 分组描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 同步数据源
	Source *int64 `json:"Source,omitnil,omitempty" name:"Source"`

	// 账号分组ID全路径，点分格式
	IdPath *string `json:"IdPath,omitnil,omitempty" name:"IdPath"`

	// 创建时间
	Itime *string `json:"Itime,omitnil,omitempty" name:"Itime"`

	// 父源账号组织ID。使用第三方导入用户源时，记录该分组在源组织架构下的分组ID
	ParentOrgId *string `json:"ParentOrgId,omitnil,omitempty" name:"ParentOrgId"`

	// 导入类型
	ImportType *string `json:"ImportType,omitnil,omitempty" name:"ImportType"`

	// miniIAM id
	MiniIamId *string `json:"MiniIamId,omitnil,omitempty" name:"MiniIamId"`

	// 该分组下含子组的所有用户总数
	UserTotal *int64 `json:"UserTotal,omitnil,omitempty" name:"UserTotal"`

	// 是否叶子节点
	IsLeaf *bool `json:"IsLeaf,omitnil,omitempty" name:"IsLeaf"`

	// 是否该账户的直接权限
	ReadOnly *bool `json:"ReadOnly,omitnil,omitempty" name:"ReadOnly"`

	// 最新一次同步任务的结果
	LatestSyncResult *string `json:"LatestSyncResult,omitnil,omitempty" name:"LatestSyncResult"`

	// 最新一次同步任务的结束时间
	LatestSyncTime *string `json:"LatestSyncTime,omitnil,omitempty" name:"LatestSyncTime"`

	// 分组名称数组
	NamePathArr []*string `json:"NamePathArr,omitnil,omitempty" name:"NamePathArr"`
}

type DescribeAccountGroupsPageResp struct {
	// 账户分响应对象集合
	Items []*DescribeAccountGroupsData `json:"Items,omitnil,omitempty" name:"Items"`

	// 分页公共对象
	Page *Paging `json:"Page,omitnil,omitempty" name:"Page"`
}

// Predefined struct for user
type DescribeAccountGroupsRequestParams struct {
	// 搜索范围：0-仅当前分组的直接子组，1-当前分组的所有子组。默认为0。
	Deepin *int64 `json:"Deepin,omitnil,omitempty" name:"Deepin"`

	// 查询条件
	// 
	// 过滤参数
	// 1、Name，string类型，按分组名过滤
	// 是否必填：否
	// 操作符: like
	// 
	// 排序条件
	// 1、Itime，string类型，按分组创建时间排序
	// 是否必填：否
	// 2、Utime，string类型，按分组更新时间排序
	// 是否必填：否
	Condition *Condition `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 父分组ID，获取该分组下的子组信息。默认查询全网根分组下子组信息。
	ParentId *int64 `json:"ParentId,omitnil,omitempty" name:"ParentId"`
}

type DescribeAccountGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 搜索范围：0-仅当前分组的直接子组，1-当前分组的所有子组。默认为0。
	Deepin *int64 `json:"Deepin,omitnil,omitempty" name:"Deepin"`

	// 查询条件
	// 
	// 过滤参数
	// 1、Name，string类型，按分组名过滤
	// 是否必填：否
	// 操作符: like
	// 
	// 排序条件
	// 1、Itime，string类型，按分组创建时间排序
	// 是否必填：否
	// 2、Utime，string类型，按分组更新时间排序
	// 是否必填：否
	Condition *Condition `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 父分组ID，获取该分组下的子组信息。默认查询全网根分组下子组信息。
	ParentId *int64 `json:"ParentId,omitnil,omitempty" name:"ParentId"`
}

func (r *DescribeAccountGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Deepin")
	delete(f, "Condition")
	delete(f, "ParentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountGroupsResponseParams struct {
	// 账号分组详情响应数据
	Data *DescribeAccountGroupsPageResp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAccountGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccountGroupsResponseParams `json:"Response"`
}

func (r *DescribeAccountGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAggrSoftCategorySoftListData struct {
	// 分页公共对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Page *Paging `json:"Page,omitnil,omitempty" name:"Page"`

	// 总数(只支持32位)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 行数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	AggrSoftCategorySoftList []*AggrCategorySoftDetailRow `json:"AggrSoftCategorySoftList,omitnil,omitempty" name:"AggrSoftCategorySoftList"`
}

// Predefined struct for user
type DescribeAggrSoftCategorySoftListRequestParams struct {
	// 过滤条件、分页参数<li>Name - String - 过滤支持：是 - 操作符:eq,neq,like,ilike,nlike - 排序支持：是 - 按类别名称过滤或排序。</li><li>CorpName - String - 过滤支持：是 - 操作符:eq,neq,like,ilike,nlike - 排序支持：是 - 按CorpName过滤或排序。</li><li>Version - String - 过滤支持：否 - 操作符:eq,like - 排序支持：是 - 按版本排序。</li><li>InstalledDeviceCount - int - 过滤支持：否 - 操作符:eq,like - 排序支持：是 - 按安装设备数量排序。</li><li>GenuineRate - float - 过滤支持：否 - 操作符:eq,like - 排序支持：是 - 按正版率排序。</li><li>AuthNum - int - 过滤支持：否 - 操作符:eq,like - 排序支持：是 - 按授权数量排序。</li><li>CategoryNamePath - String - 过滤支持：否 - 操作符:eq,like - 排序支持：是 - 按类别路径名排序。</li>
	Condition *Condition `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 操作系统类型（0: win，1：linux，2: mac，4：android，5：ios 默认值0）
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`
}

type DescribeAggrSoftCategorySoftListRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件、分页参数<li>Name - String - 过滤支持：是 - 操作符:eq,neq,like,ilike,nlike - 排序支持：是 - 按类别名称过滤或排序。</li><li>CorpName - String - 过滤支持：是 - 操作符:eq,neq,like,ilike,nlike - 排序支持：是 - 按CorpName过滤或排序。</li><li>Version - String - 过滤支持：否 - 操作符:eq,like - 排序支持：是 - 按版本排序。</li><li>InstalledDeviceCount - int - 过滤支持：否 - 操作符:eq,like - 排序支持：是 - 按安装设备数量排序。</li><li>GenuineRate - float - 过滤支持：否 - 操作符:eq,like - 排序支持：是 - 按正版率排序。</li><li>AuthNum - int - 过滤支持：否 - 操作符:eq,like - 排序支持：是 - 按授权数量排序。</li><li>CategoryNamePath - String - 过滤支持：否 - 操作符:eq,like - 排序支持：是 - 按类别路径名排序。</li>
	Condition *Condition `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 操作系统类型（0: win，1：linux，2: mac，4：android，5：ios 默认值0）
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`
}

func (r *DescribeAggrSoftCategorySoftListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAggrSoftCategorySoftListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Condition")
	delete(f, "OsType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAggrSoftCategorySoftListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAggrSoftCategorySoftListResponseParams struct {
	// 数据
	Data *DescribeAggrSoftCategorySoftListData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAggrSoftCategorySoftListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAggrSoftCategorySoftListResponseParams `json:"Response"`
}

func (r *DescribeAggrSoftCategorySoftListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAggrSoftCategorySoftListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAggrSoftDetailData struct {
	// 软件名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 平台
	// 注意：此字段可能返回 null，表示取不到有效值。
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`

	// 盗版风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	PiracyRisk *int64 `json:"PiracyRisk,omitnil,omitempty" name:"PiracyRisk"`

	// 厂商
	// 注意：此字段可能返回 null，表示取不到有效值。
	Corp *string `json:"Corp,omitnil,omitempty" name:"Corp"`

	// 已安装版本分布
	// 注意：此字段可能返回 null，表示取不到有效值。
	SoftVersionDist []*SoftVersionAndNum `json:"SoftVersionDist,omitnil,omitempty" name:"SoftVersionDist"`

	// 盗版版本安装
	// 注意：此字段可能返回 null，表示取不到有效值。
	PiracyVersionDist []*SoftVersionAndNum `json:"PiracyVersionDist,omitnil,omitempty" name:"PiracyVersionDist"`

	// 安装设备数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstalledDeviceNum *int64 `json:"InstalledDeviceNum,omitnil,omitempty" name:"InstalledDeviceNum"`

	// 盗版安装设备数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PiracyInstalledDeviceNum *int64 `json:"PiracyInstalledDeviceNum,omitnil,omitempty" name:"PiracyInstalledDeviceNum"`

	// 安装用户数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstalledUserNum *int64 `json:"InstalledUserNum,omitnil,omitempty" name:"InstalledUserNum"`

	// 盗版安装用户数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PiracyInstalledUserNum *int64 `json:"PiracyInstalledUserNum,omitnil,omitempty" name:"PiracyInstalledUserNum"`

	// 授权数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthNum *int64 `json:"AuthNum,omitnil,omitempty" name:"AuthNum"`

	// 正版率
	// 注意：此字段可能返回 null，表示取不到有效值。
	GenuineRate *float64 `json:"GenuineRate,omitnil,omitempty" name:"GenuineRate"`

	// 有新版本可升级的设备数量
	UpgradableDeviceNum *int64 `json:"UpgradableDeviceNum,omitnil,omitempty" name:"UpgradableDeviceNum"`

	// 当前可升级的最新版本信息, 每一项均为json字符串
	UpgradableVersions []*string `json:"UpgradableVersions,omitnil,omitempty" name:"UpgradableVersions"`
}

// Predefined struct for user
type DescribeAggrSoftDetailRequestParams struct {
	// 软件名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 操作系统
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`
}

type DescribeAggrSoftDetailRequest struct {
	*tchttp.BaseRequest
	
	// 软件名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 操作系统
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`
}

func (r *DescribeAggrSoftDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAggrSoftDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "OsType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAggrSoftDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAggrSoftDetailResponseParams struct {
	// 数据
	Data *DescribeAggrSoftDetailData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAggrSoftDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAggrSoftDetailResponseParams `json:"Response"`
}

func (r *DescribeAggrSoftDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAggrSoftDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAggrSoftDeviceListData struct {
	// 分页公共对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Page *Paging `json:"Page,omitnil,omitempty" name:"Page"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	AggrSoftDeviceList []*AggrSoftDeviceRow `json:"AggrSoftDeviceList,omitnil,omitempty" name:"AggrSoftDeviceList"`
}

// Predefined struct for user
type DescribeAggrSoftDeviceListRequestParams struct {
	// 过滤条件
	Condition *Condition `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 软件名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 0:win 2:mac
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`
}

type DescribeAggrSoftDeviceListRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件
	Condition *Condition `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 软件名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 0:win 2:mac
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`
}

func (r *DescribeAggrSoftDeviceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAggrSoftDeviceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Condition")
	delete(f, "Name")
	delete(f, "OsType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAggrSoftDeviceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAggrSoftDeviceListResponseParams struct {
	// 已安装终端列表
	Data *DescribeAggrSoftDeviceListData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAggrSoftDeviceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAggrSoftDeviceListResponseParams `json:"Response"`
}

func (r *DescribeAggrSoftDeviceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAggrSoftDeviceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDLPEdgeNodeGroupsRequestParams struct {
	// 管理域实例ID，用于CAM管理域权限分配。若企业未进行管理域的划分，可直接传入根域"1"，此时表示针对当前企业的全部设备和账号进行接口CRUD，具体CRUD的影响范围限制于相应接口的入参。
	DomainInstanceId *string `json:"DomainInstanceId,omitnil,omitempty" name:"DomainInstanceId"`

	// 过滤条件
	Condition *Condition `json:"Condition,omitnil,omitempty" name:"Condition"`
}

type DescribeDLPEdgeNodeGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 管理域实例ID，用于CAM管理域权限分配。若企业未进行管理域的划分，可直接传入根域"1"，此时表示针对当前企业的全部设备和账号进行接口CRUD，具体CRUD的影响范围限制于相应接口的入参。
	DomainInstanceId *string `json:"DomainInstanceId,omitnil,omitempty" name:"DomainInstanceId"`

	// 过滤条件
	Condition *Condition `json:"Condition,omitnil,omitempty" name:"Condition"`
}

func (r *DescribeDLPEdgeNodeGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDLPEdgeNodeGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainInstanceId")
	delete(f, "Condition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDLPEdgeNodeGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDLPEdgeNodeGroupsResponseParams struct {
	// 业务响应数据
	Data *DescribeDLPEdgeNodeGroupsRspData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDLPEdgeNodeGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDLPEdgeNodeGroupsResponseParams `json:"Response"`
}

func (r *DescribeDLPEdgeNodeGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDLPEdgeNodeGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDLPEdgeNodeGroupsRspData struct {
	// 分组信息
	Items []*DescribeDLPEdgeNodeGroupsRspItem `json:"Items,omitnil,omitempty" name:"Items"`

	// 分页信息
	Page *Paging `json:"Page,omitnil,omitempty" name:"Page"`
}

type DescribeDLPEdgeNodeGroupsRspItem struct {
	// 自增id，数据库中唯一
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 节点分组名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 节点分组id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 包含边缘节点数量
	EdgeCount *int64 `json:"EdgeCount,omitnil,omitempty" name:"EdgeCount"`
}

type DescribeDLPEdgeNodesPageData struct {
	// 分页信息
	Page *Paging `json:"Page,omitnil,omitempty" name:"Page"`

	// 节点列表
	Items []*DescribeDLPEdgeNodesRspItem `json:"Items,omitnil,omitempty" name:"Items"`
}

// Predefined struct for user
type DescribeDLPEdgeNodesRequestParams struct {
	// 管理域实例ID，用于CAM管理域权限分配。若企业未进行管理域的划分，可直接传入根域"1"，此时表示针对当前企业的全部设备和账号进行接口CRUD，具体CRUD的影响范围限制于相应接口的入参。
	DomainInstanceId *string `json:"DomainInstanceId,omitnil,omitempty" name:"DomainInstanceId"`

	// 过滤条件、分页参数<li>EdgeNodeName - string - 是否必填：否 - 操作符: ilike  - 排序支持：否- 按节点名称过滤。</li>
	Condition *Condition `json:"Condition,omitnil,omitempty" name:"Condition"`
}

type DescribeDLPEdgeNodesRequest struct {
	*tchttp.BaseRequest
	
	// 管理域实例ID，用于CAM管理域权限分配。若企业未进行管理域的划分，可直接传入根域"1"，此时表示针对当前企业的全部设备和账号进行接口CRUD，具体CRUD的影响范围限制于相应接口的入参。
	DomainInstanceId *string `json:"DomainInstanceId,omitnil,omitempty" name:"DomainInstanceId"`

	// 过滤条件、分页参数<li>EdgeNodeName - string - 是否必填：否 - 操作符: ilike  - 排序支持：否- 按节点名称过滤。</li>
	Condition *Condition `json:"Condition,omitnil,omitempty" name:"Condition"`
}

func (r *DescribeDLPEdgeNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDLPEdgeNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainInstanceId")
	delete(f, "Condition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDLPEdgeNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDLPEdgeNodesResponseParams struct {
	// 业务响应数据
	Data *DescribeDLPEdgeNodesPageData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDLPEdgeNodesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDLPEdgeNodesResponseParams `json:"Response"`
}

func (r *DescribeDLPEdgeNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDLPEdgeNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDLPEdgeNodesRspItem struct {
	// 自增id，数据库中唯一
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 节点分组唯一id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 节点id
	EdgeNodeId *string `json:"EdgeNodeId,omitnil,omitempty" name:"EdgeNodeId"`

	// 节点名称
	EdgeNodeName *string `json:"EdgeNodeName,omitnil,omitempty" name:"EdgeNodeName"`

	// 是否活跃/连通
	IsActive *bool `json:"IsActive,omitnil,omitempty" name:"IsActive"`

	// 节点分组名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 节点IP
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 节点版本
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 是否支持升级连接器
	IsUpgradeEnable *bool `json:"IsUpgradeEnable,omitnil,omitempty" name:"IsUpgradeEnable"`

	// 升级状态: 0(升级中) , 1(升级失败) 或 2(升级成功)
	UpgradeStatus *int64 `json:"UpgradeStatus,omitnil,omitempty" name:"UpgradeStatus"`

	// 升级状态描述
	UpgradeDescription *string `json:"UpgradeDescription,omitnil,omitempty" name:"UpgradeDescription"`

	// 规则版本
	RuleVersion *string `json:"RuleVersion,omitnil,omitempty" name:"RuleVersion"`
}

type DescribeDLPFileDetectResultData struct {
	// 提交任务时的文件md5
	FileMd5 *string `json:"FileMd5,omitnil,omitempty" name:"FileMd5"`

	// 提交任务时的文件名
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 状态：等待检测->正在检测->检测失败/检测成功。或任务不存在
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 文件检测结果，json字符串。包含文件基本信息如type，path，md5以及命中的信息。其中State为检测状态，0为待解析文件，1为检测中，2为检测完成；FileAbstract为命中的上下文摘要信息，HitRuleid是命中的规则唯一ID，HitRuleCategoryId是规则分类唯一id，HitLevel是文件的等级，HitRuleDesc是规则的名称，HitContent是具体命中的规则以及词库信息，以及命中的内容。EngineConfigVersion是当前词库版本号
	DetectResult *string `json:"DetectResult,omitnil,omitempty" name:"DetectResult"`
}

// Predefined struct for user
type DescribeDLPFileDetectResultRequestParams struct {
	// 管理域实例ID，用于CAM管理域权限分配。若企业未进行管理域的划分，可直接传入根域"1"，此时表示针对当前企业的全部设备和账号进行接口CRUD，具体CRUD的影响范围限制于相应接口的入参。
	DomainInstanceId *string `json:"DomainInstanceId,omitnil,omitempty" name:"DomainInstanceId"`

	// 查询ID，即提交送检任务接口（CreateDLPFileDetectionTask）返回的任务ID（DLPFileDetectionTaskID）
	QueryID *string `json:"QueryID,omitnil,omitempty" name:"QueryID"`
}

type DescribeDLPFileDetectResultRequest struct {
	*tchttp.BaseRequest
	
	// 管理域实例ID，用于CAM管理域权限分配。若企业未进行管理域的划分，可直接传入根域"1"，此时表示针对当前企业的全部设备和账号进行接口CRUD，具体CRUD的影响范围限制于相应接口的入参。
	DomainInstanceId *string `json:"DomainInstanceId,omitnil,omitempty" name:"DomainInstanceId"`

	// 查询ID，即提交送检任务接口（CreateDLPFileDetectionTask）返回的任务ID（DLPFileDetectionTaskID）
	QueryID *string `json:"QueryID,omitnil,omitempty" name:"QueryID"`
}

func (r *DescribeDLPFileDetectResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDLPFileDetectResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainInstanceId")
	delete(f, "QueryID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDLPFileDetectResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDLPFileDetectResultResponseParams struct {
	// 文件鉴定任务结果数据。详情查看具体数据结构
	Data *DescribeDLPFileDetectResultData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDLPFileDetectResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDLPFileDetectResultResponseParams `json:"Response"`
}

func (r *DescribeDLPFileDetectResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDLPFileDetectResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDLPFileDetectTaskResult struct {
	// 提交任务时的文件md5
	FileMd5 *string `json:"FileMd5,omitnil,omitempty" name:"FileMd5"`

	// 提交任务时的文件名
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 检测执行状态：0未执行 1等待执行 2执行中 3执行失败 4执行完成 
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	//     FileAbstract:文件摘要
	//     FileAttr:文件属性
	//     FileCategory:命中分级分类 array
	//     FileContent:命中信息json(array)
	// 	            RuleId:规则Id
	// 				RuleName:规则名称
	// 				RuleLevel:规则等级
	// 				Hits：命中词库内容
	// 				    LibraryId：词库Id
	// 					LibraryType:词库类型
	// 					LibraryName:词库名称
	// 					Attribute: 命中属性 doc.Content文件内容|doc.FileSize文件大小|doc.Name文件名|doc.Type文件类型
	// 					String  待匹配内容
	// 					Content 命中内容
	//                 HitsTotal 规则命中次数
	//     FileMd5 文件ND5
	//     FileName 文件名
	//     FileSize 文件大小
	//     FileType 文件后缀
	//     FileTypeName 文件类型名称
	//     FinalDataLevel 命中最高等级
	//     NodeId 节点唯一Id
	//     NodeIp 节点IP
	//     NodeName 节点名称
	//     OperateTime 文件操作时间
	//     Url 文件下载Url
	DetectResult *string `json:"DetectResult,omitnil,omitempty" name:"DetectResult"`

	// 检测执行状态描述
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

// Predefined struct for user
type DescribeDLPFileDetectTaskResultRequestParams struct {
	// 管理域实例ID，用于CAM管理域权限分配。若企业未进行管理域的划分，可直接传入根域"1"，此时表示针对当前企业的全部设备和账号进行接口CRUD，具体CRUD的影响范围限制于相应接口的入参。
	DomainInstanceId *string `json:"DomainInstanceId,omitnil,omitempty" name:"DomainInstanceId"`

	// 任务请求Id
	TaskRequestId *string `json:"TaskRequestId,omitnil,omitempty" name:"TaskRequestId"`
}

type DescribeDLPFileDetectTaskResultRequest struct {
	*tchttp.BaseRequest
	
	// 管理域实例ID，用于CAM管理域权限分配。若企业未进行管理域的划分，可直接传入根域"1"，此时表示针对当前企业的全部设备和账号进行接口CRUD，具体CRUD的影响范围限制于相应接口的入参。
	DomainInstanceId *string `json:"DomainInstanceId,omitnil,omitempty" name:"DomainInstanceId"`

	// 任务请求Id
	TaskRequestId *string `json:"TaskRequestId,omitnil,omitempty" name:"TaskRequestId"`
}

func (r *DescribeDLPFileDetectTaskResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDLPFileDetectTaskResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainInstanceId")
	delete(f, "TaskRequestId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDLPFileDetectTaskResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDLPFileDetectTaskResultResponseParams struct {
	// 文件鉴定任务结果数据。详情查看具体数据结构
	Data *DescribeDLPFileDetectTaskResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDLPFileDetectTaskResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDLPFileDetectTaskResultResponseParams `json:"Response"`
}

func (r *DescribeDLPFileDetectTaskResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDLPFileDetectTaskResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceChildGroupsRequestParams struct {
	// 管理域实例ID，用于CAM管理域权限分配。若企业未进行管理域的划分，可直接传入根域"1"，此时表示针对当前企业的全部设备和账号进行接口CRUD，具体CRUD的影响范围限制于相应接口的入参。
	DomainInstanceId *string `json:"DomainInstanceId,omitnil,omitempty" name:"DomainInstanceId"`

	// 过滤条件参数（字段含义请参考接口返回值）
	// - Name, 类型String，支持操作：【like，ilike】，支持排序
	// 
	// 
	// 
	// 
	// 分页参数
	// - PageNum 从1开始，小于等于0时使用默认参数
	// - PageSize 最大值5000，最好不超过100
	Condition *Condition `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 父分组id，默认0：表示获取全网终端分组
	ParentId *int64 `json:"ParentId,omitnil,omitempty" name:"ParentId"`

	// 操作系统类型（0：win，1：linux，2：mac，4：android，5：ios；默认0：系统win）
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`
}

type DescribeDeviceChildGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 管理域实例ID，用于CAM管理域权限分配。若企业未进行管理域的划分，可直接传入根域"1"，此时表示针对当前企业的全部设备和账号进行接口CRUD，具体CRUD的影响范围限制于相应接口的入参。
	DomainInstanceId *string `json:"DomainInstanceId,omitnil,omitempty" name:"DomainInstanceId"`

	// 过滤条件参数（字段含义请参考接口返回值）
	// - Name, 类型String，支持操作：【like，ilike】，支持排序
	// 
	// 
	// 
	// 
	// 分页参数
	// - PageNum 从1开始，小于等于0时使用默认参数
	// - PageSize 最大值5000，最好不超过100
	Condition *Condition `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 父分组id，默认0：表示获取全网终端分组
	ParentId *int64 `json:"ParentId,omitnil,omitempty" name:"ParentId"`

	// 操作系统类型（0：win，1：linux，2：mac，4：android，5：ios；默认0：系统win）
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`
}

func (r *DescribeDeviceChildGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceChildGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainInstanceId")
	delete(f, "Condition")
	delete(f, "ParentId")
	delete(f, "OsType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceChildGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceChildGroupsResponseParams struct {
	// 查询设备组子分组详情响应结构
	Data *DescribeDeviceChildGroupsRspData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDeviceChildGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceChildGroupsResponseParams `json:"Response"`
}

func (r *DescribeDeviceChildGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceChildGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceChildGroupsRspData struct {
	// 返回的数组列表
	Items []*DeviceGroupDetail `json:"Items,omitnil,omitempty" name:"Items"`
}

type DescribeDeviceDetailListData struct {
	// 账号名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 计算机名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComputerName *string `json:"ComputerName,omitnil,omitempty" name:"ComputerName"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 用户组IdPath
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountGroupIdPath *string `json:"AccountGroupIdPath,omitnil,omitempty" name:"AccountGroupIdPath"`

	// 用户组id(只支持32位)
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountGroupId *int64 `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// 终端组名path
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupNamePath *string `json:"GroupNamePath,omitnil,omitempty" name:"GroupNamePath"`

	// Ip地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 用户组名
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountGroupName *string `json:"AccountGroupName,omitnil,omitempty" name:"AccountGroupName"`

	// 终端组IdPath
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupIdPath *string `json:"GroupIdPath,omitnil,omitempty" name:"GroupIdPath"`

	// 唯一标识Mid
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mid *string `json:"Mid,omitnil,omitempty" name:"Mid"`

	// IOA账号名
	// 注意：此字段可能返回 null，表示取不到有效值。
	IoaUserName *string `json:"IoaUserName,omitnil,omitempty" name:"IoaUserName"`

	// 所在分组Id(只支持32位)
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 所在分组Name
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Mac地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mac *string `json:"Mac,omitnil,omitempty" name:"Mac"`

	// 软件版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 用户组名Path
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountGroupNamePath *string `json:"AccountGroupNamePath,omitnil,omitempty" name:"AccountGroupNamePath"`

	// 列表Id(只支持32位)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`
}

type DescribeDeviceDetailListPageData struct {
	// 终端详情响应对象集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*DescribeDeviceDetailListData `json:"Items,omitnil,omitempty" name:"Items"`

	// 分页公共对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Page *Paging `json:"Page,omitnil,omitempty" name:"Page"`
}

// Predefined struct for user
type DescribeDeviceDetailListRequestParams struct {
	// 系统类型(只支持32位)
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`

	// 终端分组id(只支持32位)
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 过滤条件、分页参数
	// <li>Name - String - 过滤支持：是 - 操作符:eq,like - 排序支持：是 。</li>
	Condition *Condition `json:"Condition,omitnil,omitempty" name:"Condition"`
}

type DescribeDeviceDetailListRequest struct {
	*tchttp.BaseRequest
	
	// 系统类型(只支持32位)
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`

	// 终端分组id(只支持32位)
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 过滤条件、分页参数
	// <li>Name - String - 过滤支持：是 - 操作符:eq,like - 排序支持：是 。</li>
	Condition *Condition `json:"Condition,omitnil,omitempty" name:"Condition"`
}

func (r *DescribeDeviceDetailListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceDetailListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OsType")
	delete(f, "GroupId")
	delete(f, "Condition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceDetailListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceDetailListResponseParams struct {
	// 业务响应数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *DescribeDeviceDetailListPageData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDeviceDetailListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceDetailListResponseParams `json:"Response"`
}

func (r *DescribeDeviceDetailListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceDetailListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceHardwareInfoItem struct {
	// 设备ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 设备唯一标识符
	Mid *string `json:"Mid,omitnil,omitempty" name:"Mid"`

	// OS平台 0 Windows 1 Linux 2 macOS 4 Android 5 iOS
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`

	// 终端名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 终端用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 授权状态（ 4未授权 5已授权）
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 设备所属分组ID
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 设备所属分组名
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 设备所属分组路径
	GroupNamePath *string `json:"GroupNamePath,omitnil,omitempty" name:"GroupNamePath"`

	// 最近登录账户的姓名
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// 出口IP
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// MAC地址
	MacAddr *string `json:"MacAddr,omitnil,omitempty" name:"MacAddr"`

	// CPU品牌型号
	Cpu *string `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存信息
	Memory *string `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 硬盘信息
	HardDiskSize *string `json:"HardDiskSize,omitnil,omitempty" name:"HardDiskSize"`

	// 显示器品牌型号
	Monitor *string `json:"Monitor,omitnil,omitempty" name:"Monitor"`

	// 终端备注名
	RemarkName *string `json:"RemarkName,omitnil,omitempty" name:"RemarkName"`
}

// Predefined struct for user
type DescribeDeviceHardwareInfoListRequestParams struct {
	// 【必填】设备分组id（需要和OsType匹配），下面是私有化场景下默认id：id-名称-操作系统1	全网终端	Win2	未分组终端	Win30000000	服务器	Win40000101	全网终端	Linux40000102	未分组终端	Linux40000103	服务器	Linux40000201	全网终端	macOS40000202	未分组终端	macOS40000203	服务器	macOS40000401	全网终端	Android40000402	未分组终端	Android40000501	全网终端	iOS40000502	未分组终端	iOSSaaS需要调用分组接口DescribeDeviceChildGroups获取对应分组id
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 【必填】操作系统类型（0: win，1：linux，2: mac，4：android，5：ios   默认值0），需要和GroupId或者GroupIds匹配
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`

	// 管理域实例ID，用于CAM管理域权限分配。若企业未进行管理域的划分，可直接传入根域"1"，此时表示针对当前企业的全部设备和账号进行接口CRUD，具体CRUD的影响范围限制于相应接口的入参。
	DomainInstanceId *string `json:"DomainInstanceId,omitnil,omitempty" name:"DomainInstanceId"`

	// 过滤条件参数（字段含义请参考接口返回值）  - Name, 类型String，支持操作：【eq，like，ilike】，支持排序  - UserName, 类型String，支持操作：【eq，like，ilike】，支持排序  - IoaUserName，类型String，支持操作：【eq，like，ilike】，支持排序  - MacAddr, 类型String，支持操作：【eq，like，ilike】，支持排序  - Ip, 类型String，支持操作：【eq，like，ilike】，支持排序  - Mid, 类型String，支持操作：【eq，like，ilike】，支持排序  ，支持排序分页参数  - PageNum 从1开始，小于等于0时使用默认参数 - PageSize 最大值5000，最好不超过100
	Condition *Condition `json:"Condition,omitnil,omitempty" name:"Condition"`
}

type DescribeDeviceHardwareInfoListRequest struct {
	*tchttp.BaseRequest
	
	// 【必填】设备分组id（需要和OsType匹配），下面是私有化场景下默认id：id-名称-操作系统1	全网终端	Win2	未分组终端	Win30000000	服务器	Win40000101	全网终端	Linux40000102	未分组终端	Linux40000103	服务器	Linux40000201	全网终端	macOS40000202	未分组终端	macOS40000203	服务器	macOS40000401	全网终端	Android40000402	未分组终端	Android40000501	全网终端	iOS40000502	未分组终端	iOSSaaS需要调用分组接口DescribeDeviceChildGroups获取对应分组id
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 【必填】操作系统类型（0: win，1：linux，2: mac，4：android，5：ios   默认值0），需要和GroupId或者GroupIds匹配
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`

	// 管理域实例ID，用于CAM管理域权限分配。若企业未进行管理域的划分，可直接传入根域"1"，此时表示针对当前企业的全部设备和账号进行接口CRUD，具体CRUD的影响范围限制于相应接口的入参。
	DomainInstanceId *string `json:"DomainInstanceId,omitnil,omitempty" name:"DomainInstanceId"`

	// 过滤条件参数（字段含义请参考接口返回值）  - Name, 类型String，支持操作：【eq，like，ilike】，支持排序  - UserName, 类型String，支持操作：【eq，like，ilike】，支持排序  - IoaUserName，类型String，支持操作：【eq，like，ilike】，支持排序  - MacAddr, 类型String，支持操作：【eq，like，ilike】，支持排序  - Ip, 类型String，支持操作：【eq，like，ilike】，支持排序  - Mid, 类型String，支持操作：【eq，like，ilike】，支持排序  ，支持排序分页参数  - PageNum 从1开始，小于等于0时使用默认参数 - PageSize 最大值5000，最好不超过100
	Condition *Condition `json:"Condition,omitnil,omitempty" name:"Condition"`
}

func (r *DescribeDeviceHardwareInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceHardwareInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "OsType")
	delete(f, "DomainInstanceId")
	delete(f, "Condition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceHardwareInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceHardwareInfoListResponseParams struct {
	// 分页的data数据
	Data *DescribeDeviceHardwareInfoListRspData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDeviceHardwareInfoListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceHardwareInfoListResponseParams `json:"Response"`
}

func (r *DescribeDeviceHardwareInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceHardwareInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceHardwareInfoListRspData struct {
	// 分页数据
	Page *Paging `json:"Page,omitnil,omitempty" name:"Page"`

	// 终端硬件信息数据数组
	Items []*DescribeDeviceHardwareInfoItem `json:"Items,omitnil,omitempty" name:"Items"`
}

// Predefined struct for user
type DescribeDeviceInfoRequestParams struct {
	// 管理域实例ID，用于CAM管理域权限分配。若企业未进行管理域的划分，可直接传入根域"1"，此时表示针对当前企业的全部设备和账号进行接口CRUD，具体CRUD的影响范围限制于相应接口的入参。
	DomainInstanceId *string `json:"DomainInstanceId,omitnil,omitempty" name:"DomainInstanceId"`

	// 终端id
	Mid *string `json:"Mid,omitnil,omitempty" name:"Mid"`

	// 查询类型  process_list network_list service_list
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type DescribeDeviceInfoRequest struct {
	*tchttp.BaseRequest
	
	// 管理域实例ID，用于CAM管理域权限分配。若企业未进行管理域的划分，可直接传入根域"1"，此时表示针对当前企业的全部设备和账号进行接口CRUD，具体CRUD的影响范围限制于相应接口的入参。
	DomainInstanceId *string `json:"DomainInstanceId,omitnil,omitempty" name:"DomainInstanceId"`

	// 终端id
	Mid *string `json:"Mid,omitnil,omitempty" name:"Mid"`

	// 查询类型  process_list network_list service_list
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *DescribeDeviceInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainInstanceId")
	delete(f, "Mid")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceInfoResponseParams struct {
	// 业务响应数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *DescribeDeviceInfoRspData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDeviceInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceInfoResponseParams `json:"Response"`
}

func (r *DescribeDeviceInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceInfoRspData struct {
	// 分页的具体数据对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessList []*DeviceProcessInfo `json:"ProcessList,omitnil,omitempty" name:"ProcessList"`

	// 分页的具体数据对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkList []*DeviceNetworkInfo `json:"NetworkList,omitnil,omitempty" name:"NetworkList"`

	// 分页的具体数据对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceList []*DeviceServiceInfo `json:"ServiceList,omitnil,omitempty" name:"ServiceList"`
}

type DescribeDeviceVirtualGroupsPageRsp struct {
	// 分页公共对象
	Page *Paging `json:"Page,omitnil,omitempty" name:"Page"`

	// 终端自定义分组列表数据
	Items []*DeviceVirtualDeviceGroupsDetail `json:"Items,omitnil,omitempty" name:"Items"`
}

// Predefined struct for user
type DescribeDeviceVirtualGroupsRequestParams struct {
	// 管理域实例ID，用于CAM管理域权限分配。若企业未进行管理域的划分，可直接传入根域"1"，此时表示针对当前企业的全部设备和账号进行接口CRUD，具体CRUD的影响范围限制于相应接口的入参。
	DomainInstanceId *string `json:"DomainInstanceId,omitnil,omitempty" name:"DomainInstanceId"`

	// 滤条件、分页参数 <li>Name - String - 是否必填：否 - 操作符: like  - 排序支持：否- 按终端自定义分组过滤。</li> <li>DeviceVirtualGroupName - String - 是否必填：否 - 操作符: like  - 排序支持：否- 按终端自定义分组过滤。</li>
	Condition *Condition `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 系统类型（0: win，1：linux，2: mac，4：android，5：ios   默认值0）
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`

	// 非必填，自定义分组ids
	VirtualGroupIds []*int64 `json:"VirtualGroupIds,omitnil,omitempty" name:"VirtualGroupIds"`
}

type DescribeDeviceVirtualGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 管理域实例ID，用于CAM管理域权限分配。若企业未进行管理域的划分，可直接传入根域"1"，此时表示针对当前企业的全部设备和账号进行接口CRUD，具体CRUD的影响范围限制于相应接口的入参。
	DomainInstanceId *string `json:"DomainInstanceId,omitnil,omitempty" name:"DomainInstanceId"`

	// 滤条件、分页参数 <li>Name - String - 是否必填：否 - 操作符: like  - 排序支持：否- 按终端自定义分组过滤。</li> <li>DeviceVirtualGroupName - String - 是否必填：否 - 操作符: like  - 排序支持：否- 按终端自定义分组过滤。</li>
	Condition *Condition `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 系统类型（0: win，1：linux，2: mac，4：android，5：ios   默认值0）
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`

	// 非必填，自定义分组ids
	VirtualGroupIds []*int64 `json:"VirtualGroupIds,omitnil,omitempty" name:"VirtualGroupIds"`
}

func (r *DescribeDeviceVirtualGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceVirtualGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainInstanceId")
	delete(f, "Condition")
	delete(f, "OsType")
	delete(f, "VirtualGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceVirtualGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceVirtualGroupsResponseParams struct {
	// 查询终端自定义分组的Data数据
	Data *DescribeDeviceVirtualGroupsPageRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDeviceVirtualGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceVirtualGroupsResponseParams `json:"Response"`
}

func (r *DescribeDeviceVirtualGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceVirtualGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDevicesPageRsp struct {
	// 数据分页信息
	Paging *Paging `json:"Paging,omitnil,omitempty" name:"Paging"`

	// 业务响应数据
	Items []*DeviceDetail `json:"Items,omitnil,omitempty" name:"Items"`
}

// Predefined struct for user
type DescribeDevicesRequestParams struct {
	// 管理域实例ID，用于CAM管理域权限分配。若企业未进行管理域的划分，可直接传入根域"1"，此时表示针对当前企业的全部设备和账号进行接口CRUD，具体CRUD的影响范围限制于相应接口的入参。
	DomainInstanceId *string `json:"DomainInstanceId,omitnil,omitempty" name:"DomainInstanceId"`

	// 过滤条件参数（字段含义请参考接口返回值）
	// 
	// - Mid, 类型String，支持操作：【eq，like，ilike】，支持排序
	// - Name, 类型String，支持操作：【eq，like，ilike】，支持排序
	// - Itime, 类型String，支持操作：【eq，like，ilike】，支持排序
	// - UserName, 类型String，支持操作：【eq，like，ilike】，支持排序
	// - MacAddr, 类型String，支持操作：【eq，like，ilike】，支持排序
	// - UserId, 类型String，支持操作：【eq，like，ilike】，支持排序
	// - Ip, 类型String，支持操作：【eq，like，ilike】，支持排序
	// - Tags，类型String，支持操作：【eq，like，ilike】，支持排序
	// - LocalIpList，类型String，支持操作：【eq，like，ilike】，支持排序
	// - SerialNum，类型String，支持操作：【eq，like，ilike】，支持排序
	// - Version，类型String，支持操作：【eq，like，ilike】，支持排序
	// - StrVersion，类型String，支持操作：【eq，like，ilike】，支持排序
	// - RtpStatus，类型String，支持操作：【eq，like，ilike】，**不支持排序**
	// - HostName，类型String，支持操作：【eq，like，ilike】，支持排序
	// - IoaUserName，类型String，支持操作：【eq，like，ilike】，支持排序
	// - GroupName，类型String，支持操作：【eq，like，ilike】，支持排序
	// - CriticalVulListCount，**类型Int**，支持操作：【eq】，**不支持排序**
	// - RiskCount，**类型Int**，支持操作：【eq】，**不支持排序**
	// - VulVersion，类型String，支持操作：【eq，like，ilike】，**不支持排序**
	// - Virusver，类型String，支持操作：【eq，like，ilike】，**不支持排序**
	// - SysRepver，类型String，支持操作：【eq，like，ilike】，**不支持排序**
	// - BaseBoardSn，类型String，支持操作：【eq，like，ilike】，支持排序
	// - Os，类型String，支持操作：【eq，like，ilike】，支持排序
	// - ConnActiveTime，类型String，支持操作：【eq，like，ilike】，**不支持排序**
	// - FirewallStatus，**类型Int**，支持操作：【eq】，**不支持排序**
	// - ProfileName，类型String，支持操作：【eq，like，ilike】，支持排序
	// - DomainName，类型String，支持操作：【eq，like，ilike】，支持排序
	// - SysRepVersion，类型String，支持操作：【eq，like，ilike】，支持排序
	// - VirusVer，类型String，支持操作：【eq，like，ilike】，支持排序
	// - Cpu，类型String，支持操作：【eq，like，ilike】，支持排序
	// - Memory，类型String，支持操作：【eq，like，ilike】，支持排序
	// - HardDiskSize，类型String，支持操作：【eq，like，ilike】，支持排序
	// - HardwareChangeCount，**类型Int**，支持操作：【eq】，支持排序
	// - AccountName，类型String，支持操作：【like.ilike】，支持排序
	// - AccountGroupName，类型String，支持操作：【like.ilike】，支持排序
	// - ScreenRecordingPermission，**类型Int**，支持操作：【eq】，支持排序
	// - DiskAccessPermission，**类型Int**，支持操作：【eq】，支持排序
	// 
	// 
	// 
	// 
	// 
	// 分页参数
	// - PageNum 从1开始，小于等于0时使用默认参数
	// - PageSize 最大值5000，最好不超过100
	Condition *Condition `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 【和GroupIds必须有一个填写】设备分组id（需要和OsType匹配），下面是私有化场景下默认id：
	// id-名称-操作系统
	// 1	全网终端	Win
	// 2	未分组终端	Win
	// 30000000	服务器	Win
	// 40000101	全网终端	Linux
	// 40000102	未分组终端	Linux
	// 40000103	服务器	Linux
	// 40000201	全网终端	macOS
	// 40000202	未分组终端	macOS
	// 40000203	服务器	macOS
	// 40000401	全网终端	Android
	// 40000402	未分组终端	Android
	// 40000501	全网终端	iOS
	// 40000502	未分组终端	iOS
	// 
	// 
	// SaaS需要调用分组接口DescribeDeviceChildGroups获取对应分组id
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 操作系统类型（0: win，1：linux，2: mac，4：android，5：ios   默认值0），需要和GroupId或者GroupIds匹配
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`

	// 在线状态 （2表示在线，0或者1表示离线）
	OnlineStatus *int64 `json:"OnlineStatus,omitnil,omitempty" name:"OnlineStatus"`

	// 过滤条件--兼容旧接口,参数同Condition
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段--兼容旧接口,参数同Condition
	Sort *Sort `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 获取第几页--兼容旧接口,参数同Condition
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 每页获取数--兼容旧接口,参数同Condition
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 授权状态： 4基础授权 5高级授权
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 【和GroupId必须有一个填写】设备分组id列表（需要和OsType匹配）
	GroupIds []*int64 `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`
}

type DescribeDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 管理域实例ID，用于CAM管理域权限分配。若企业未进行管理域的划分，可直接传入根域"1"，此时表示针对当前企业的全部设备和账号进行接口CRUD，具体CRUD的影响范围限制于相应接口的入参。
	DomainInstanceId *string `json:"DomainInstanceId,omitnil,omitempty" name:"DomainInstanceId"`

	// 过滤条件参数（字段含义请参考接口返回值）
	// 
	// - Mid, 类型String，支持操作：【eq，like，ilike】，支持排序
	// - Name, 类型String，支持操作：【eq，like，ilike】，支持排序
	// - Itime, 类型String，支持操作：【eq，like，ilike】，支持排序
	// - UserName, 类型String，支持操作：【eq，like，ilike】，支持排序
	// - MacAddr, 类型String，支持操作：【eq，like，ilike】，支持排序
	// - UserId, 类型String，支持操作：【eq，like，ilike】，支持排序
	// - Ip, 类型String，支持操作：【eq，like，ilike】，支持排序
	// - Tags，类型String，支持操作：【eq，like，ilike】，支持排序
	// - LocalIpList，类型String，支持操作：【eq，like，ilike】，支持排序
	// - SerialNum，类型String，支持操作：【eq，like，ilike】，支持排序
	// - Version，类型String，支持操作：【eq，like，ilike】，支持排序
	// - StrVersion，类型String，支持操作：【eq，like，ilike】，支持排序
	// - RtpStatus，类型String，支持操作：【eq，like，ilike】，**不支持排序**
	// - HostName，类型String，支持操作：【eq，like，ilike】，支持排序
	// - IoaUserName，类型String，支持操作：【eq，like，ilike】，支持排序
	// - GroupName，类型String，支持操作：【eq，like，ilike】，支持排序
	// - CriticalVulListCount，**类型Int**，支持操作：【eq】，**不支持排序**
	// - RiskCount，**类型Int**，支持操作：【eq】，**不支持排序**
	// - VulVersion，类型String，支持操作：【eq，like，ilike】，**不支持排序**
	// - Virusver，类型String，支持操作：【eq，like，ilike】，**不支持排序**
	// - SysRepver，类型String，支持操作：【eq，like，ilike】，**不支持排序**
	// - BaseBoardSn，类型String，支持操作：【eq，like，ilike】，支持排序
	// - Os，类型String，支持操作：【eq，like，ilike】，支持排序
	// - ConnActiveTime，类型String，支持操作：【eq，like，ilike】，**不支持排序**
	// - FirewallStatus，**类型Int**，支持操作：【eq】，**不支持排序**
	// - ProfileName，类型String，支持操作：【eq，like，ilike】，支持排序
	// - DomainName，类型String，支持操作：【eq，like，ilike】，支持排序
	// - SysRepVersion，类型String，支持操作：【eq，like，ilike】，支持排序
	// - VirusVer，类型String，支持操作：【eq，like，ilike】，支持排序
	// - Cpu，类型String，支持操作：【eq，like，ilike】，支持排序
	// - Memory，类型String，支持操作：【eq，like，ilike】，支持排序
	// - HardDiskSize，类型String，支持操作：【eq，like，ilike】，支持排序
	// - HardwareChangeCount，**类型Int**，支持操作：【eq】，支持排序
	// - AccountName，类型String，支持操作：【like.ilike】，支持排序
	// - AccountGroupName，类型String，支持操作：【like.ilike】，支持排序
	// - ScreenRecordingPermission，**类型Int**，支持操作：【eq】，支持排序
	// - DiskAccessPermission，**类型Int**，支持操作：【eq】，支持排序
	// 
	// 
	// 
	// 
	// 
	// 分页参数
	// - PageNum 从1开始，小于等于0时使用默认参数
	// - PageSize 最大值5000，最好不超过100
	Condition *Condition `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 【和GroupIds必须有一个填写】设备分组id（需要和OsType匹配），下面是私有化场景下默认id：
	// id-名称-操作系统
	// 1	全网终端	Win
	// 2	未分组终端	Win
	// 30000000	服务器	Win
	// 40000101	全网终端	Linux
	// 40000102	未分组终端	Linux
	// 40000103	服务器	Linux
	// 40000201	全网终端	macOS
	// 40000202	未分组终端	macOS
	// 40000203	服务器	macOS
	// 40000401	全网终端	Android
	// 40000402	未分组终端	Android
	// 40000501	全网终端	iOS
	// 40000502	未分组终端	iOS
	// 
	// 
	// SaaS需要调用分组接口DescribeDeviceChildGroups获取对应分组id
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 操作系统类型（0: win，1：linux，2: mac，4：android，5：ios   默认值0），需要和GroupId或者GroupIds匹配
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`

	// 在线状态 （2表示在线，0或者1表示离线）
	OnlineStatus *int64 `json:"OnlineStatus,omitnil,omitempty" name:"OnlineStatus"`

	// 过滤条件--兼容旧接口,参数同Condition
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段--兼容旧接口,参数同Condition
	Sort *Sort `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 获取第几页--兼容旧接口,参数同Condition
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 每页获取数--兼容旧接口,参数同Condition
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 授权状态： 4基础授权 5高级授权
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 【和GroupId必须有一个填写】设备分组id列表（需要和OsType匹配）
	GroupIds []*int64 `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`
}

func (r *DescribeDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainInstanceId")
	delete(f, "Condition")
	delete(f, "GroupId")
	delete(f, "OsType")
	delete(f, "OnlineStatus")
	delete(f, "Filters")
	delete(f, "Sort")
	delete(f, "PageNum")
	delete(f, "PageSize")
	delete(f, "Status")
	delete(f, "GroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDevicesResponseParams struct {
	// 分页的data数据
	Data *DescribeDevicesPageRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDevicesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDevicesResponseParams `json:"Response"`
}

func (r *DescribeDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLocalAccountAccountGroupsData struct {
	// 组Id(只支持32位)
	AccountGroupId *int64 `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`
}

type DescribeLocalAccountsData struct {
	// uid，数据库中唯一
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 账号，登录账号
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 账号id，同Id字段
	AccountId *int64 `json:"AccountId,omitnil,omitempty" name:"AccountId"`

	// 账号所在的分组id
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 账号所在的分组名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 账号所在的分组名称路径，用英文.分割
	NamePath *string `json:"NamePath,omitnil,omitempty" name:"NamePath"`

	// 账号来源,0表示本地账号(只支持32位)
	Source *int64 `json:"Source,omitnil,omitempty" name:"Source"`

	// 账号状态,0禁用，1启用(只支持32位)
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 账号的创建时间
	Itime *string `json:"Itime,omitnil,omitempty" name:"Itime"`

	// 账号的最后更新时间
	Utime *string `json:"Utime,omitnil,omitempty" name:"Utime"`

	// 账号的扩展信息，包含邮箱、手机号、身份证、职位等信息
	ExtraInfo *string `json:"ExtraInfo,omitnil,omitempty" name:"ExtraInfo"`

	// 用户风险等级，枚举：none, low, middle, high
	RiskLevel *string `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 所属组
	AccountGroups []*DescribeLocalAccountAccountGroupsData `json:"AccountGroups,omitnil,omitempty" name:"AccountGroups"`

	// 绑定手机端设备数
	MobileBindNum *int64 `json:"MobileBindNum,omitnil,omitempty" name:"MobileBindNum"`

	// 绑定Pc端设备数
	PcBindNum *int64 `json:"PcBindNum,omitnil,omitempty" name:"PcBindNum"`

	// 账号在线状态 1：在线 2：离线
	OnlineStatus *int64 `json:"OnlineStatus,omitnil,omitempty" name:"OnlineStatus"`

	// 账号活跃状态 1：活跃 2：非活跃
	ActiveStatus *int64 `json:"ActiveStatus,omitnil,omitempty" name:"ActiveStatus"`

	// 账号登录时间
	LoginTime *string `json:"LoginTime,omitnil,omitempty" name:"LoginTime"`

	// 账号登出时间
	LogoutTime *string `json:"LogoutTime,omitnil,omitempty" name:"LogoutTime"`
}

type DescribeLocalAccountsPage struct {
	// 公共分页对象
	Page *Paging `json:"Page,omitnil,omitempty" name:"Page"`

	// 获取账号列表响应的单个对象
	Items []*DescribeLocalAccountsData `json:"Items,omitnil,omitempty" name:"Items"`
}

// Predefined struct for user
type DescribeLocalAccountsRequestParams struct {
	// 查询条件：过滤或排序
	// 1、UserName，string类型，姓名
	// 是否必填：否
	// 过滤支持：是，支持eq、like、ilike
	// 排序支持：否
	// 2、UserId，string类型，账户
	// 是否必填：否
	// 过滤支持：是，支持eq、like、ilike
	// 排序支持：否
	// 3、Phone，string类型，手机号
	// 是否必填：否
	// 过滤支持：是，支持eq、like、ilike
	// 排序支持：否
	Condition *Condition `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 获取账号的分组ID，不传默认获取全网根账号组
	AccountGroupId *int64 `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// 是否仅展示当前目录下用户 1： 递归显示 2：仅显示当前目录下用户
	ShowFlag *int64 `json:"ShowFlag,omitnil,omitempty" name:"ShowFlag"`
}

type DescribeLocalAccountsRequest struct {
	*tchttp.BaseRequest
	
	// 查询条件：过滤或排序
	// 1、UserName，string类型，姓名
	// 是否必填：否
	// 过滤支持：是，支持eq、like、ilike
	// 排序支持：否
	// 2、UserId，string类型，账户
	// 是否必填：否
	// 过滤支持：是，支持eq、like、ilike
	// 排序支持：否
	// 3、Phone，string类型，手机号
	// 是否必填：否
	// 过滤支持：是，支持eq、like、ilike
	// 排序支持：否
	Condition *Condition `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 获取账号的分组ID，不传默认获取全网根账号组
	AccountGroupId *int64 `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// 是否仅展示当前目录下用户 1： 递归显示 2：仅显示当前目录下用户
	ShowFlag *int64 `json:"ShowFlag,omitnil,omitempty" name:"ShowFlag"`
}

func (r *DescribeLocalAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLocalAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Condition")
	delete(f, "AccountGroupId")
	delete(f, "ShowFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLocalAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLocalAccountsResponseParams struct {
	// 获取账号列表响应的分页对象
	Data *DescribeLocalAccountsPage `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLocalAccountsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLocalAccountsResponseParams `json:"Response"`
}

func (r *DescribeLocalAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLocalAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRootAccountGroupRequestParams struct {

}

type DescribeRootAccountGroupRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeRootAccountGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRootAccountGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRootAccountGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRootAccountGroupResponseParams struct {
	// 账号根分组响应详情
	Data *GetAccountGroupData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRootAccountGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRootAccountGroupResponseParams `json:"Response"`
}

func (r *DescribeRootAccountGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRootAccountGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSoftCensusListByDeviceData struct {
	// 终端用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// mac地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	MacAddr *string `json:"MacAddr,omitnil,omitempty" name:"MacAddr"`

	// 终端计算机名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 终端组路径名
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupNamePath *string `json:"GroupNamePath,omitnil,omitempty" name:"GroupNamePath"`

	// IP地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 唯一标识Mid
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mid *string `json:"Mid,omitnil,omitempty" name:"Mid"`

	// 企业账户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	IoaUserName *string `json:"IoaUserName,omitnil,omitempty" name:"IoaUserName"`

	// 终端分组Id(只支持32位)
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 终端组名
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 终端列表Id(只支持32位)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 软件数量(只支持32位)
	// 注意：此字段可能返回 null，表示取不到有效值。
	SoftNum *int64 `json:"SoftNum,omitnil,omitempty" name:"SoftNum"`

	// 盗版风险（1=风险;2=未知）
	// 注意：此字段可能返回 null，表示取不到有效值。
	PiracyRisk *int64 `json:"PiracyRisk,omitnil,omitempty" name:"PiracyRisk"`

	// 终端备注名
	RemarkName *string `json:"RemarkName,omitnil,omitempty" name:"RemarkName"`
}

type DescribeSoftCensusListByDevicePageData struct {
	// 软件统计响应对象集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*DescribeSoftCensusListByDeviceData `json:"Items,omitnil,omitempty" name:"Items"`

	// 分页公共对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Page *Paging `json:"Page,omitnil,omitempty" name:"Page"`
}

// Predefined struct for user
type DescribeSoftCensusListByDeviceRequestParams struct {
	// 必填，终端分组ID
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 管理域实例ID，用于CAM管理域权限分配。若企业未进行管理域的划分，可直接传入根域"1"，此时表示针对当前企业的全部设备和账号进行接口CRUD，具体CRUD的影响范围限制于相应接口的入参。
	DomainInstanceId *string `json:"DomainInstanceId,omitnil,omitempty" name:"DomainInstanceId"`

	// 系统类型（0: win，1：linux，2: mac，4：android，5：ios  ）；默认值0
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`

	// 过滤条件、分页参数   <li>Name - String - 是否必填：否 - 操作符: eq,like,ilike  - 排序支持：否 - 备注：字段含义，终端名。</li> 	<li>UserName - String - 是否必填：否 - 操作符: eq,like,ilike  - 排序支持：否 - 备注：字段含义，终端用户名。</li> 	<li>IoaUserName - String - 是否必填：否 - 操作符: eq,like,ilike  - 排序支持：否 - 备注：字段含义，最近登录账号。</li> 	<li>Ip - String - 是否必填：否 - 操作符: eq,like,ilike  - 排序支持：否 - 备注：字段含义，IP地址。</li> 	<li>MacAddr - String - 是否必填：否 - 操作符: eq,like,ilike  - 排序支持：否 - 备注：字段含义，MAC地址。</li>
	Condition *Condition `json:"Condition,omitnil,omitempty" name:"Condition"`
}

type DescribeSoftCensusListByDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 必填，终端分组ID
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 管理域实例ID，用于CAM管理域权限分配。若企业未进行管理域的划分，可直接传入根域"1"，此时表示针对当前企业的全部设备和账号进行接口CRUD，具体CRUD的影响范围限制于相应接口的入参。
	DomainInstanceId *string `json:"DomainInstanceId,omitnil,omitempty" name:"DomainInstanceId"`

	// 系统类型（0: win，1：linux，2: mac，4：android，5：ios  ）；默认值0
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`

	// 过滤条件、分页参数   <li>Name - String - 是否必填：否 - 操作符: eq,like,ilike  - 排序支持：否 - 备注：字段含义，终端名。</li> 	<li>UserName - String - 是否必填：否 - 操作符: eq,like,ilike  - 排序支持：否 - 备注：字段含义，终端用户名。</li> 	<li>IoaUserName - String - 是否必填：否 - 操作符: eq,like,ilike  - 排序支持：否 - 备注：字段含义，最近登录账号。</li> 	<li>Ip - String - 是否必填：否 - 操作符: eq,like,ilike  - 排序支持：否 - 备注：字段含义，IP地址。</li> 	<li>MacAddr - String - 是否必填：否 - 操作符: eq,like,ilike  - 排序支持：否 - 备注：字段含义，MAC地址。</li>
	Condition *Condition `json:"Condition,omitnil,omitempty" name:"Condition"`
}

func (r *DescribeSoftCensusListByDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSoftCensusListByDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "DomainInstanceId")
	delete(f, "OsType")
	delete(f, "Condition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSoftCensusListByDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSoftCensusListByDeviceResponseParams struct {
	// 业务响应数据
	Data *DescribeSoftCensusListByDevicePageData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSoftCensusListByDeviceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSoftCensusListByDeviceResponseParams `json:"Response"`
}

func (r *DescribeSoftCensusListByDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSoftCensusListByDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSoftwareInformationPageData struct {
	// 软件详情响应对象集合
	Items []*SoftwareInformationData `json:"Items,omitnil,omitempty" name:"Items"`

	// 分页公共对象
	Page *Paging `json:"Page,omitnil,omitempty" name:"Page"`
}

// Predefined struct for user
type DescribeSoftwareInformationRequestParams struct {
	// 终端唯一标识Mid
	Mid *string `json:"Mid,omitnil,omitempty" name:"Mid"`

	// 过滤条件、分页参数
	// <li>Name - String - 过滤支持：是 - 操作符:eq,like - 排序支持：是 。</li>
	Condition *Condition `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 0:win 2:mac
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`
}

type DescribeSoftwareInformationRequest struct {
	*tchttp.BaseRequest
	
	// 终端唯一标识Mid
	Mid *string `json:"Mid,omitnil,omitempty" name:"Mid"`

	// 过滤条件、分页参数
	// <li>Name - String - 过滤支持：是 - 操作符:eq,like - 排序支持：是 。</li>
	Condition *Condition `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 0:win 2:mac
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`
}

func (r *DescribeSoftwareInformationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSoftwareInformationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Mid")
	delete(f, "Condition")
	delete(f, "OsType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSoftwareInformationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSoftwareInformationResponseParams struct {
	// 业务响应数据
	Data *DescribeSoftwareInformationPageData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSoftwareInformationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSoftwareInformationResponseParams `json:"Response"`
}

func (r *DescribeSoftwareInformationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSoftwareInformationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirtualDevicesPageRsp struct {
	// 数据分页信息
	Paging *Paging `json:"Paging,omitnil,omitempty" name:"Paging"`

	// 设备列表
	Items []*DeviceDetail `json:"Items,omitnil,omitempty" name:"Items"`
}

// Predefined struct for user
type DescribeVirtualDevicesRequestParams struct {
	// 管理域实例ID，用于CAM管理域权限分配。若企业未进行管理域的划分，可直接传入根域"1"，此时表示针对当前企业的全部设备和账号进行接口CRUD，具体CRUD的影响范围限制于相应接口的入参。
	DomainInstanceId *string `json:"DomainInstanceId,omitnil,omitempty" name:"DomainInstanceId"`

	// 过滤条件参数（字段含义请参考接口返回值）- Mid, 类型String，支持操作：【eq，like，ilike】，支持排序- Name, 类型String，支持操作：【eq，like，ilike】，支持排序- Itime, 类型String，支持操作：【eq，like，ilike】，支持排序- UserName, 类型String，支持操作：【eq，like，ilike】，支持排序- MacAddr, 类型String，支持操作：【eq，like，ilike】，支持排序- UserId, 类型String，支持操作：【eq，like，ilike】，支持排序- Ip, 类型String，支持操作：【eq，like，ilike】，支持排序- Tags，类型String，支持操作：【eq，like，ilike】，支持排序- LocalIpList，类型String，支持操作：【eq，like，ilike】，支持排序- SerialNum，类型String，支持操作：【eq，like，ilike】，支持排序- Version，类型String，支持操作：【eq，like，ilike】，支持排序- StrVersion，类型String，支持操作：【eq，like，ilike】，支持排序- RtpStatus，类型String，支持操作：【eq，like，ilike】，**不支持排序**- HostName，类型String，支持操作：【eq，like，ilike】，支持排序- IoaUserName，类型String，支持操作：【eq，like，ilike】，支持排序- GroupName，类型String，支持操作：【eq，like，ilike】，支持排序- CriticalVulListCount，**类型Int**，支持操作：【eq】，**不支持排序**- RiskCount，**类型Int**，支持操作：【eq】，**不支持排序**- VulVersion，类型String，支持操作：【eq，like，ilike】，**不支持排序**- Virusver，类型String，支持操作：【eq，like，ilike】，**不支持排序**- SysRepver，类型String，支持操作：【eq，like，ilike】，**不支持排序**- BaseBoardSn，类型String，支持操作：【eq，like，ilike】，支持排序- Os，类型String，支持操作：【eq，like，ilike】，支持排序- ConnActiveTime，类型String，支持操作：【eq，like，ilike】，**不支持排序**- FirewallStatus，**类型Int**，支持操作：【eq】，**不支持排序**- ProfileName，类型String，支持操作：【eq，like，ilike】，支持排序- DomainName，类型String，支持操作：【eq，like，ilike】，支持排序- SysRepVersion，类型String，支持操作：【eq，like，ilike】，支持排序- VirusVer，类型String，支持操作：【eq，like，ilike】，支持排序- Cpu，类型String，支持操作：【eq，like，ilike】，支持排序- Memory，类型String，支持操作：【eq，like，ilike】，支持排序- HardDiskSize，类型String，支持操作：【eq，like，ilike】，支持排序- HardwareChangeCount，**类型Int**，支持操作：【eq】，支持排序- AccountName，类型String，支持操作：【like.ilike】，支持排序- AccountGroupName，类型String，支持操作：【like.ilike】，支持排序- ScreenRecordingPermission，**类型Int**，支持操作：【eq】，支持排序- DiskAccessPermission，**类型Int**，支持操作：【eq】，支持排序分页参数- PageNum 从1开始，小于等于0时使用默认参数- PageSize 最大值5000，最好不超过100
	Condition *Condition `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 终端自定义分组ID（0：获取租户全部自定义分组下的终端数据；其他值：获取具体ID分组下的终端数据）
	DeviceVirtualGroupId *int64 `json:"DeviceVirtualGroupId,omitnil,omitempty" name:"DeviceVirtualGroupId"`

	// 系统类型（0: win，1：linux，2: mac，4：android，5：ios   默认值0）
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`

	// 选填，在线状态 （2表示在线，0或者1表示离线）
	OnlineStatus *int64 `json:"OnlineStatus,omitnil,omitempty" name:"OnlineStatus"`
}

type DescribeVirtualDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 管理域实例ID，用于CAM管理域权限分配。若企业未进行管理域的划分，可直接传入根域"1"，此时表示针对当前企业的全部设备和账号进行接口CRUD，具体CRUD的影响范围限制于相应接口的入参。
	DomainInstanceId *string `json:"DomainInstanceId,omitnil,omitempty" name:"DomainInstanceId"`

	// 过滤条件参数（字段含义请参考接口返回值）- Mid, 类型String，支持操作：【eq，like，ilike】，支持排序- Name, 类型String，支持操作：【eq，like，ilike】，支持排序- Itime, 类型String，支持操作：【eq，like，ilike】，支持排序- UserName, 类型String，支持操作：【eq，like，ilike】，支持排序- MacAddr, 类型String，支持操作：【eq，like，ilike】，支持排序- UserId, 类型String，支持操作：【eq，like，ilike】，支持排序- Ip, 类型String，支持操作：【eq，like，ilike】，支持排序- Tags，类型String，支持操作：【eq，like，ilike】，支持排序- LocalIpList，类型String，支持操作：【eq，like，ilike】，支持排序- SerialNum，类型String，支持操作：【eq，like，ilike】，支持排序- Version，类型String，支持操作：【eq，like，ilike】，支持排序- StrVersion，类型String，支持操作：【eq，like，ilike】，支持排序- RtpStatus，类型String，支持操作：【eq，like，ilike】，**不支持排序**- HostName，类型String，支持操作：【eq，like，ilike】，支持排序- IoaUserName，类型String，支持操作：【eq，like，ilike】，支持排序- GroupName，类型String，支持操作：【eq，like，ilike】，支持排序- CriticalVulListCount，**类型Int**，支持操作：【eq】，**不支持排序**- RiskCount，**类型Int**，支持操作：【eq】，**不支持排序**- VulVersion，类型String，支持操作：【eq，like，ilike】，**不支持排序**- Virusver，类型String，支持操作：【eq，like，ilike】，**不支持排序**- SysRepver，类型String，支持操作：【eq，like，ilike】，**不支持排序**- BaseBoardSn，类型String，支持操作：【eq，like，ilike】，支持排序- Os，类型String，支持操作：【eq，like，ilike】，支持排序- ConnActiveTime，类型String，支持操作：【eq，like，ilike】，**不支持排序**- FirewallStatus，**类型Int**，支持操作：【eq】，**不支持排序**- ProfileName，类型String，支持操作：【eq，like，ilike】，支持排序- DomainName，类型String，支持操作：【eq，like，ilike】，支持排序- SysRepVersion，类型String，支持操作：【eq，like，ilike】，支持排序- VirusVer，类型String，支持操作：【eq，like，ilike】，支持排序- Cpu，类型String，支持操作：【eq，like，ilike】，支持排序- Memory，类型String，支持操作：【eq，like，ilike】，支持排序- HardDiskSize，类型String，支持操作：【eq，like，ilike】，支持排序- HardwareChangeCount，**类型Int**，支持操作：【eq】，支持排序- AccountName，类型String，支持操作：【like.ilike】，支持排序- AccountGroupName，类型String，支持操作：【like.ilike】，支持排序- ScreenRecordingPermission，**类型Int**，支持操作：【eq】，支持排序- DiskAccessPermission，**类型Int**，支持操作：【eq】，支持排序分页参数- PageNum 从1开始，小于等于0时使用默认参数- PageSize 最大值5000，最好不超过100
	Condition *Condition `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 终端自定义分组ID（0：获取租户全部自定义分组下的终端数据；其他值：获取具体ID分组下的终端数据）
	DeviceVirtualGroupId *int64 `json:"DeviceVirtualGroupId,omitnil,omitempty" name:"DeviceVirtualGroupId"`

	// 系统类型（0: win，1：linux，2: mac，4：android，5：ios   默认值0）
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`

	// 选填，在线状态 （2表示在线，0或者1表示离线）
	OnlineStatus *int64 `json:"OnlineStatus,omitnil,omitempty" name:"OnlineStatus"`
}

func (r *DescribeVirtualDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVirtualDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainInstanceId")
	delete(f, "Condition")
	delete(f, "DeviceVirtualGroupId")
	delete(f, "OsType")
	delete(f, "OnlineStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVirtualDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVirtualDevicesResponseParams struct {
	// 返回的具体Data数据
	Data *DescribeVirtualDevicesPageRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVirtualDevicesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVirtualDevicesResponseParams `json:"Response"`
}

func (r *DescribeVirtualDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVirtualDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeviceDetail struct {
	// 设备ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 设备唯一标识码，在ioa中每个设备有唯一标识码
	Mid *string `json:"Mid,omitnil,omitempty" name:"Mid"`

	// 终端名（设备名）
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 设备所在分组ID
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// OS平台，0：Windows 、1： Linux、 2：macOS 、4： Android、 5: iOS。默认是0
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`

	// 设备IP地址（出口IP）
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 在线状态，2：在线、0或者1:离线
	OnlineStatus *int64 `json:"OnlineStatus,omitnil,omitempty" name:"OnlineStatus"`

	// 客户端版本号-大整数
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 客户端版本号-点分字符串
	StrVersion *string `json:"StrVersion,omitnil,omitempty" name:"StrVersion"`

	// 首次在线时间
	Itime *string `json:"Itime,omitnil,omitempty" name:"Itime"`

	// 最后一次在线时间
	ConnActiveTime *string `json:"ConnActiveTime,omitnil,omitempty" name:"ConnActiveTime"`

	// 设备是否加锁 ，1：锁定 0或者2：未锁定。
	Locked *int64 `json:"Locked,omitnil,omitempty" name:"Locked"`

	// 设备本地IP列表, 包括IP
	LocalIpList *string `json:"LocalIpList,omitnil,omitempty" name:"LocalIpList"`

	// 宿主机id（需要宿主机也安装iOA才能显示）
	HostId *int64 `json:"HostId,omitnil,omitempty" name:"HostId"`

	// 设备所属分组名
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 设备所属分组路径
	GroupNamePath *string `json:"GroupNamePath,omitnil,omitempty" name:"GroupNamePath"`

	// 未修复高危漏洞数(只支持32位)
	CriticalVulListCount *int64 `json:"CriticalVulListCount,omitnil,omitempty" name:"CriticalVulListCount"`

	// 操作系统名称
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 操作系统位数
	OsBits *int64 `json:"OsBits,omitnil,omitempty" name:"OsBits"`

	// 操作系统版本
	OsVersion *string `json:"OsVersion,omitnil,omitempty" name:"OsVersion"`

	// 操作系统语言
	OsLanguage *string `json:"OsLanguage,omitnil,omitempty" name:"OsLanguage"`

	// 操作系统安装时间
	OsInstallDate *string `json:"OsInstallDate,omitnil,omitempty" name:"OsInstallDate"`

	// 设备名，和Name相同
	ComputerName *string `json:"ComputerName,omitnil,omitempty" name:"ComputerName"`

	// 登录域名
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// MAC地址
	MacAddr *string `json:"MacAddr,omitnil,omitempty" name:"MacAddr"`

	// 漏洞数
	VulCount *int64 `json:"VulCount,omitnil,omitempty" name:"VulCount"`

	// 病毒风险数
	RiskCount *int64 `json:"RiskCount,omitnil,omitempty" name:"RiskCount"`

	// 病毒库版本
	VirusVer *string `json:"VirusVer,omitnil,omitempty" name:"VirusVer"`

	// 漏洞库版本
	VulVersion *string `json:"VulVersion,omitnil,omitempty" name:"VulVersion"`

	// 系统修复引擎版本
	SysRepVersion *string `json:"SysRepVersion,omitnil,omitempty" name:"SysRepVersion"`

	// 高危补丁列表
	VulCriticalList []*string `json:"VulCriticalList,omitnil,omitempty" name:"VulCriticalList"`

	// 标签
	Tags *string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 终端用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 防火墙状态，不等于0表示开启
	FirewallStatus *int64 `json:"FirewallStatus,omitnil,omitempty" name:"FirewallStatus"`

	// SN序列号
	SerialNum *string `json:"SerialNum,omitnil,omitempty" name:"SerialNum"`

	// 设备管控策略版本
	DeviceStrategyVer *string `json:"DeviceStrategyVer,omitnil,omitempty" name:"DeviceStrategyVer"`

	// NGN策略版本
	NGNStrategyVer *string `json:"NGNStrategyVer,omitnil,omitempty" name:"NGNStrategyVer"`

	// 最近登录账户的账号(账号系统用户账号)
	IOAUserName *string `json:"IOAUserName,omitnil,omitempty" name:"IOAUserName"`

	// 设备管控新策略
	DeviceNewStrategyVer *string `json:"DeviceNewStrategyVer,omitnil,omitempty" name:"DeviceNewStrategyVer"`

	// NGN策略新版本
	NGNNewStrategyVer *string `json:"NGNNewStrategyVer,omitnil,omitempty" name:"NGNNewStrategyVer"`

	// 宿主机名称（需要宿主机也安装iOA才能显示）
	HostName *string `json:"HostName,omitnil,omitempty" name:"HostName"`

	// 主板序列号
	BaseBoardSn *string `json:"BaseBoardSn,omitnil,omitempty" name:"BaseBoardSn"`

	// 绑定账户名称
	AccountUsers *string `json:"AccountUsers,omitnil,omitempty" name:"AccountUsers"`

	// 身份策略版本
	IdentityStrategyVer *string `json:"IdentityStrategyVer,omitnil,omitempty" name:"IdentityStrategyVer"`

	// 身份策略新版本
	IdentityNewStrategyVer *string `json:"IdentityNewStrategyVer,omitnil,omitempty" name:"IdentityNewStrategyVer"`

	// 最近登录账号部门
	AccountGroupName *string `json:"AccountGroupName,omitnil,omitempty" name:"AccountGroupName"`

	// 最近登录账户的姓名(账号系统用户姓名)
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// 账号组id
	AccountGroupId *int64 `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// 终端备注名
	RemarkName *string `json:"RemarkName,omitnil,omitempty" name:"RemarkName"`
}

type DeviceGroupDetail struct {
	// 设备组id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 设备组名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 设备组描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 父节点id
	ParentId *int64 `json:"ParentId,omitnil,omitempty" name:"ParentId"`

	// 基于id的节点路径
	IdPath *string `json:"IdPath,omitnil,omitempty" name:"IdPath"`

	// 基于名称的节点路径
	NamePath *string `json:"NamePath,omitnil,omitempty" name:"NamePath"`

	// 分组锁定状态
	Locked *int64 `json:"Locked,omitnil,omitempty" name:"Locked"`

	// 系统类型（0: win，1：linux，2: mac，4：android，5：ios   ）
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`

	// 排序
	Sort *int64 `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 是否自动调整
	FromAuto *int64 `json:"FromAuto,omitnil,omitempty" name:"FromAuto"`

	// 子节点数量
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 图标
	Icon *string `json:"Icon,omitnil,omitempty" name:"Icon"`

	// 是否有ip
	WithIp *int64 `json:"WithIp,omitnil,omitempty" name:"WithIp"`

	// 是否有组ip
	HasIp *bool `json:"HasIp,omitnil,omitempty" name:"HasIp"`

	// 是否是叶子节点
	IsLeaf *bool `json:"IsLeaf,omitnil,omitempty" name:"IsLeaf"`

	// 是否只读
	ReadOnly *bool `json:"ReadOnly,omitnil,omitempty" name:"ReadOnly"`

	// 对应绑定的账号id
	BindAccount *int64 `json:"BindAccount,omitnil,omitempty" name:"BindAccount"`

	// 绑定账号的用户名
	BindAccountName *string `json:"BindAccountName,omitnil,omitempty" name:"BindAccountName"`
}

type DeviceNetworkInfo struct {
	// 本地地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocalAddr *string `json:"LocalAddr,omitnil,omitempty" name:"LocalAddr"`

	// 本地端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocalPort *int64 `json:"LocalPort,omitnil,omitempty" name:"LocalPort"`

	// 进程id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessId *int64 `json:"ProcessId,omitnil,omitempty" name:"ProcessId"`

	// 进程名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessName *string `json:"ProcessName,omitnil,omitempty" name:"ProcessName"`

	// 协议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 远程地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	RemoteAddr *string `json:"RemoteAddr,omitnil,omitempty" name:"RemoteAddr"`

	// 远程端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	RemotePort *int64 `json:"RemotePort,omitnil,omitempty" name:"RemotePort"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *int64 `json:"State,omitnil,omitempty" name:"State"`
}

type DeviceProcessInfo struct {
	// 命令行
	// 注意：此字段可能返回 null，表示取不到有效值。
	CmdLine *string `json:"CmdLine,omitnil,omitempty" name:"CmdLine"`

	// 内存
	// 注意：此字段可能返回 null，表示取不到有效值。
	Memory *string `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 进程id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessId *int64 `json:"ProcessId,omitnil,omitempty" name:"ProcessId"`

	// 启动用户
	// 注意：此字段可能返回 null，表示取不到有效值。
	User *string `json:"User,omitnil,omitempty" name:"User"`
}

type DeviceServiceInfo struct {
	// 命令行
	// 注意：此字段可能返回 null，表示取不到有效值。
	CmdLine *string `json:"CmdLine,omitnil,omitempty" name:"CmdLine"`

	// 内存
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 进程id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessId *int64 `json:"ProcessId,omitnil,omitempty" name:"ProcessId"`

	// 启动类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartType *int64 `json:"StartType,omitnil,omitempty" name:"StartType"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *int64 `json:"State,omitnil,omitempty" name:"State"`

	// 启动用户
	// 注意：此字段可能返回 null，表示取不到有效值。
	User *string `json:"User,omitnil,omitempty" name:"User"`
}

type DeviceVirtualDeviceGroupsDetail struct {
	// 终端自定义分组id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 自定义分组名称
	DeviceVirtualGroupName *string `json:"DeviceVirtualGroupName,omitnil,omitempty" name:"DeviceVirtualGroupName"`

	// 设备数
	DeviceCount *int64 `json:"DeviceCount,omitnil,omitempty" name:"DeviceCount"`

	// 系统类型（0: win，1：linux，2: mac，4：android，5：ios  ）
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`

	// 创建时间
	Itime *string `json:"Itime,omitnil,omitempty" name:"Itime"`

	// 更新时间
	Utime *string `json:"Utime,omitnil,omitempty" name:"Utime"`
}

type ExportSoftwareDownloadUrlRspData struct {
	// 下载的url
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadURL *string `json:"DownloadURL,omitnil,omitempty" name:"DownloadURL"`

	// 超过一定时间走异步任务
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

// Predefined struct for user
type ExportSoftwareInformationListRequestParams struct {
	// 终端唯一标识Mid
	Mid *string `json:"Mid,omitnil,omitempty" name:"Mid"`

	// 过滤条件、分页参数
	// <li>Name - String - 过滤支持：是 - 操作符:eq,like - 排序支持：是 。</li>
	Condition *Condition `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 系统类型0:win 2:mac
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`
}

type ExportSoftwareInformationListRequest struct {
	*tchttp.BaseRequest
	
	// 终端唯一标识Mid
	Mid *string `json:"Mid,omitnil,omitempty" name:"Mid"`

	// 过滤条件、分页参数
	// <li>Name - String - 过滤支持：是 - 操作符:eq,like - 排序支持：是 。</li>
	Condition *Condition `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 系统类型0:win 2:mac
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`
}

func (r *ExportSoftwareInformationListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportSoftwareInformationListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Mid")
	delete(f, "Condition")
	delete(f, "OsType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportSoftwareInformationListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportSoftwareInformationListResponseParams struct {
	// 业务响应数据
	Data *ExportSoftwareDownloadUrlRspData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ExportSoftwareInformationListResponse struct {
	*tchttp.BaseResponse
	Response *ExportSoftwareInformationListResponseParams `json:"Response"`
}

func (r *ExportSoftwareInformationListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportSoftwareInformationListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 过滤字段
	Field *string `json:"Field,omitnil,omitempty" name:"Field"`

	// 过滤方式： eq:等于,net:不等于,like,nlike,gt:大于,lt:小于,egt:大于等于,elt:小于等于。具体支持哪些过滤方式，结合具体接口字段描述来定
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 过滤条件
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type FilterGroup struct {
	// Filters 条件过滤
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type GetAccountGroupData struct {
	// 分组名称全路径，点分格式
	NamePath *string `json:"NamePath,omitnil,omitempty" name:"NamePath"`

	// 分组ID全路径，数组格式
	IdPathArr []*int64 `json:"IdPathArr,omitnil,omitempty" name:"IdPathArr"`

	// 分组扩展信息
	ExtraInfo *string `json:"ExtraInfo,omitnil,omitempty" name:"ExtraInfo"`

	// 最后更新时间
	Utime *string `json:"Utime,omitnil,omitempty" name:"Utime"`

	// 当前分组的父分组ID
	ParentId *uint64 `json:"ParentId,omitnil,omitempty" name:"ParentId"`

	// 源账号组ID，该字段仅适用于第三方同步的组织架构，通过OrgId-Id构成源组织架构分组ID-现组织架构分组ID映射关系
	OrgId *string `json:"OrgId,omitnil,omitempty" name:"OrgId"`

	// 分组名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分组ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 分组描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 分组导入源(只支持32位)
	Source *uint64 `json:"Source,omitnil,omitempty" name:"Source"`

	// 分组ID全路径，点分格式
	IdPath *string `json:"IdPath,omitnil,omitempty" name:"IdPath"`

	// 创建时间
	Itime *string `json:"Itime,omitnil,omitempty" name:"Itime"`

	// 父源账号组ID，该字段仅适用于第三方同步的组织架构
	ParentOrgId *string `json:"ParentOrgId,omitnil,omitempty" name:"ParentOrgId"`

	// 导入信息,json格式
	Import *string `json:"Import,omitnil,omitempty" name:"Import"`

	// 是否开启导入架构
	ImportEnable *bool `json:"ImportEnable,omitnil,omitempty" name:"ImportEnable"`

	// 导入类型
	ImportType *string `json:"ImportType,omitnil,omitempty" name:"ImportType"`

	// miniIAMId，MiniIAM源才有
	MiniIamId *string `json:"MiniIamId,omitnil,omitempty" name:"MiniIamId"`
}

type ModifyVirtualDeviceGroupsReqItem struct {
	// 设备mid
	DeviceMid *string `json:"DeviceMid,omitnil,omitempty" name:"DeviceMid"`

	// 操作标识  0:删除设备 1:添加设备
	Operation *int64 `json:"Operation,omitnil,omitempty" name:"Operation"`
}

// Predefined struct for user
type ModifyVirtualDeviceGroupsRequestParams struct {
	// 必填，操作的设备列表数据
	DeviceList []*ModifyVirtualDeviceGroupsReqItem `json:"DeviceList,omitnil,omitempty" name:"DeviceList"`

	// 管理域实例ID，用于CAM管理域权限分配。若企业未进行管理域的划分，可直接传入根域"1"，此时表示针对当前企业的全部设备和账号进行接口CRUD，具体CRUD的影响范围限制于相应接口的入参。
	DomainInstanceId *string `json:"DomainInstanceId,omitnil,omitempty" name:"DomainInstanceId"`

	// 添加到的终端自定义分组id。和DeviceVirtualGroupIds互斥，必填其一，优先使用本参数
	DeviceVirtualGroupId *int64 `json:"DeviceVirtualGroupId,omitnil,omitempty" name:"DeviceVirtualGroupId"`

	// 要添加的终端自定义分组id列表
	DeviceVirtualGroupIds []*int64 `json:"DeviceVirtualGroupIds,omitnil,omitempty" name:"DeviceVirtualGroupIds"`

	// 系统类型（0: win，1：linux，2: mac，4：android，5：ios   默认值0）
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`
}

type ModifyVirtualDeviceGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 必填，操作的设备列表数据
	DeviceList []*ModifyVirtualDeviceGroupsReqItem `json:"DeviceList,omitnil,omitempty" name:"DeviceList"`

	// 管理域实例ID，用于CAM管理域权限分配。若企业未进行管理域的划分，可直接传入根域"1"，此时表示针对当前企业的全部设备和账号进行接口CRUD，具体CRUD的影响范围限制于相应接口的入参。
	DomainInstanceId *string `json:"DomainInstanceId,omitnil,omitempty" name:"DomainInstanceId"`

	// 添加到的终端自定义分组id。和DeviceVirtualGroupIds互斥，必填其一，优先使用本参数
	DeviceVirtualGroupId *int64 `json:"DeviceVirtualGroupId,omitnil,omitempty" name:"DeviceVirtualGroupId"`

	// 要添加的终端自定义分组id列表
	DeviceVirtualGroupIds []*int64 `json:"DeviceVirtualGroupIds,omitnil,omitempty" name:"DeviceVirtualGroupIds"`

	// 系统类型（0: win，1：linux，2: mac，4：android，5：ios   默认值0）
	OsType *int64 `json:"OsType,omitnil,omitempty" name:"OsType"`
}

func (r *ModifyVirtualDeviceGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVirtualDeviceGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceList")
	delete(f, "DomainInstanceId")
	delete(f, "DeviceVirtualGroupId")
	delete(f, "DeviceVirtualGroupIds")
	delete(f, "OsType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVirtualDeviceGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVirtualDeviceGroupsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyVirtualDeviceGroupsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVirtualDeviceGroupsResponseParams `json:"Response"`
}

func (r *ModifyVirtualDeviceGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVirtualDeviceGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Paging struct {
	// 每页条数
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页码
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 总页数
	PageCount *uint64 `json:"PageCount,omitnil,omitempty" name:"PageCount"`

	// 记录总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`
}

type RuleExpression struct {
	// 规则元数据
	Items []*RuleItem `json:"Items,omitnil,omitempty" name:"Items"`

	// 关系
	Relation *string `json:"Relation,omitnil,omitempty" name:"Relation"`
}

type RuleItem struct {
	// 字段名称
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 操作关系（等于、不等于、包含、不包含）
	Operate *string `json:"Operate,omitnil,omitempty" name:"Operate"`

	// 内容
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 内容，v2多值版本使用
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type SimpleRule struct {
	// 规则表达式
	Expressions []*RuleExpression `json:"Expressions,omitnil,omitempty" name:"Expressions"`

	// 表达式间逻辑关系
	Relation *string `json:"Relation,omitnil,omitempty" name:"Relation"`
}

type SoftVersionAndNum struct {
	// 软件版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 安装数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Num *int64 `json:"Num,omitnil,omitempty" name:"Num"`
}

type SoftwareInformationData struct {
	// 软件名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 安装时间
	InstallDate *string `json:"InstallDate,omitnil,omitempty" name:"InstallDate"`

	// 软件列表id(只支持32位)
	SoftwareId *int64 `json:"SoftwareId,omitnil,omitempty" name:"SoftwareId"`

	// 唯一标识Mid
	Mid *string `json:"Mid,omitnil,omitempty" name:"Mid"`

	// 软件版本
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 公司名
	CorpName *string `json:"CorpName,omitnil,omitempty" name:"CorpName"`

	// 列表Id(只支持32位)
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 盗版风险（0:未支持，1:风险，2:未发现，3:未开启）
	PiracyRisk *int64 `json:"PiracyRisk,omitnil,omitempty" name:"PiracyRisk"`
}

type Sort struct {
	// 排序字段
	Field *string `json:"Field,omitnil,omitempty" name:"Field"`

	// 排序方式
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}