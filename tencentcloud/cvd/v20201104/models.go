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

package v20201104

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type AssociateSecurityGroupRequestParams struct {

}

type AssociateSecurityGroupRequest struct {
	*tchttp.BaseRequest
	
}

func (r *AssociateSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateSecurityGroupResponseParams struct {
	// 关联成功的桌面列表
	SuccessSet []*string `json:"SuccessSet,omitnil,omitempty" name:"SuccessSet"`

	// 关联失败的桌面列表
	FailureSet []*string `json:"FailureSet,omitnil,omitempty" name:"FailureSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AssociateSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *AssociateSecurityGroupResponseParams `json:"Response"`
}

func (r *AssociateSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttributeInfo struct {
	// <p>实例属性名</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>实例属性值</p>
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type BindInfo struct {
	// <p>实例绑定的用户ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

// Predefined struct for user
type BindInstancesRequestParams struct {
	// <p>指定得到云桌面用户id</p>
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// <p>待分配桌面实例id</p>
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// <p>待分配桌面类型</p>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// <p>待分配桌面所在区域</p>
	InstanceRegion *string `json:"InstanceRegion,omitnil,omitempty" name:"InstanceRegion"`
}

type BindInstancesRequest struct {
	*tchttp.BaseRequest
	
	// <p>指定得到云桌面用户id</p>
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// <p>待分配桌面实例id</p>
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// <p>待分配桌面类型</p>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// <p>待分配桌面所在区域</p>
	InstanceRegion *string `json:"InstanceRegion,omitnil,omitempty" name:"InstanceRegion"`
}

func (r *BindInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserIds")
	delete(f, "InstanceIds")
	delete(f, "InstanceType")
	delete(f, "InstanceRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindInstancesResponseParams struct {
	// <p>待分配桌面状态</p>
	BindInfoList []*BindStatusInfo `json:"BindInfoList,omitnil,omitempty" name:"BindInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BindInstancesResponse struct {
	*tchttp.BaseResponse
	Response *BindInstancesResponseParams `json:"Response"`
}

func (r *BindInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindStatusInfo struct {
	// <p>桌面实例id</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>实例的绑定状态。取值：<br>UNBOUND - 未绑定<br>BINDING - 绑定中<br>BOUND - 绑定<br>UNBINDING-解绑中</p>
	BindStatus *string `json:"BindStatus,omitnil,omitempty" name:"BindStatus"`

	// <p>绑定的用户id【BOUND状态不允许为空】</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>绑定用户名称【BOUND状态不允许为空】</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// <p>绑定用户显示名称【BOUND状态不允许为空】</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`
}

type ChargePrepaid struct {
	// 购买实例的时长，单位：月。取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12。
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 自动续费标识。取值范围：
	// NOTIFY_AND_AUTO_RENEW：通知过期且自动续费
	// NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费
	// DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费
	// 
	// 默认取值：NOTIFY_AND_AUTO_RENEW。若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
}

type CommandDetail struct {
	// 租户 Uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 租户 APPID
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 任务类型，取值有：
	// PowerShellScript
	// Bat
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// RUNNING - 执行中；FINISHED - 执行完成
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务脚本，Base64 编码
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// 任务唯一 ID
	CommandId *string `json:"CommandId,omitnil,omitempty" name:"CommandId"`

	// 任务完成时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 任务完成时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`

	// 各实例运行结果集
	Invocations []*InvocationDetail `json:"Invocations,omitnil,omitempty" name:"Invocations"`

	// 是否以绑定用户运行任务，默认 false：system 用户执行
	RunAsBoundUser *bool `json:"RunAsBoundUser,omitnil,omitempty" name:"RunAsBoundUser"`

	// 是否前台运行任务，默认 false：后台运行任务
	RunForeground *bool `json:"RunForeground,omitnil,omitempty" name:"RunForeground"`
}

// Predefined struct for user
type CreateCommandTaskRequestParams struct {
	// 脚本任务，需要 Base64 编码
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// 脚本任务类型：windows 桌面可选有：    PowerShellScript — powershell脚本    BatScript — Bat 脚本
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 执行脚本任务的实例列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 是否以绑定用户身份执行，默认 System身份运行false： System身份运行，默认行为true：以绑定用户身份执行
	RunAsBoundUser *bool `json:"RunAsBoundUser,omitnil,omitempty" name:"RunAsBoundUser"`

	// 是否前台运行false：后台运行，默认行为true：前台运行
	RunForeground *bool `json:"RunForeground,omitnil,omitempty" name:"RunForeground"`
}

type CreateCommandTaskRequest struct {
	*tchttp.BaseRequest
	
	// 脚本任务，需要 Base64 编码
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// 脚本任务类型：windows 桌面可选有：    PowerShellScript — powershell脚本    BatScript — Bat 脚本
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 执行脚本任务的实例列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 是否以绑定用户身份执行，默认 System身份运行false： System身份运行，默认行为true：以绑定用户身份执行
	RunAsBoundUser *bool `json:"RunAsBoundUser,omitnil,omitempty" name:"RunAsBoundUser"`

	// 是否前台运行false：后台运行，默认行为true：前台运行
	RunForeground *bool `json:"RunForeground,omitnil,omitempty" name:"RunForeground"`
}

func (r *CreateCommandTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCommandTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Command")
	delete(f, "Type")
	delete(f, "InstanceIds")
	delete(f, "RunAsBoundUser")
	delete(f, "RunForeground")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCommandTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCommandTaskResponseParams struct {
	// 脚本任务唯一 ID
	CommandId *string `json:"CommandId,omitnil,omitempty" name:"CommandId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCommandTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateCommandTaskResponseParams `json:"Response"`
}

func (r *CreateCommandTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCommandTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDataDiskRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据盘类型，有效值CLOUD_SSD/CLOUD_PREMIUM
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 数据盘大小
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 硬盘快照id
	SnapShotId *string `json:"SnapShotId,omitnil,omitempty" name:"SnapShotId"`
}

type CreateDataDiskRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据盘类型，有效值CLOUD_SSD/CLOUD_PREMIUM
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 数据盘大小
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 硬盘快照id
	SnapShotId *string `json:"SnapShotId,omitnil,omitempty" name:"SnapShotId"`
}

func (r *CreateDataDiskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDataDiskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DiskType")
	delete(f, "DiskSize")
	delete(f, "SnapShotId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDataDiskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDataDiskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDataDiskResponse struct {
	*tchttp.BaseResponse
	Response *CreateDataDiskResponseParams `json:"Response"`
}

func (r *CreateDataDiskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDataDiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDomainControlPolicyRequestParams struct {
	// <p>策略名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>策略描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>策略类型，取值范围[&quot;WHITE&quot;,&quot;BLACK&quot;]，WHITE - 白名单策略；BLACK - 黑名单策略；默认 值为BLACK</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>域名控制规则详情</p>
	DetailSet []*DomainControlDetail `json:"DetailSet,omitnil,omitempty" name:"DetailSet"`

	// <p>IP控制规则详情</p>
	IpDetailSet []*IpControlDetail `json:"IpDetailSet,omitnil,omitempty" name:"IpDetailSet"`

	// <p>作用到的实例列表</p>
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type CreateDomainControlPolicyRequest struct {
	*tchttp.BaseRequest
	
	// <p>策略名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>策略描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>策略类型，取值范围[&quot;WHITE&quot;,&quot;BLACK&quot;]，WHITE - 白名单策略；BLACK - 黑名单策略；默认 值为BLACK</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>域名控制规则详情</p>
	DetailSet []*DomainControlDetail `json:"DetailSet,omitnil,omitempty" name:"DetailSet"`

	// <p>IP控制规则详情</p>
	IpDetailSet []*IpControlDetail `json:"IpDetailSet,omitnil,omitempty" name:"IpDetailSet"`

	// <p>作用到的实例列表</p>
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *CreateDomainControlPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDomainControlPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "Type")
	delete(f, "DetailSet")
	delete(f, "IpDetailSet")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDomainControlPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDomainControlPolicyResponseParams struct {
	// <p>策略唯一ID</p>
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDomainControlPolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreateDomainControlPolicyResponseParams `json:"Response"`
}

func (r *CreateDomainControlPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDomainControlPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateImageRequestParams struct {
	// <p>云桌面实例ID。</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>镜像名称。只支持中文、字母、数字或连接符号&quot;-&quot;，且最多60个字符。<br>默认值：空。</p>
	ImageName *string `json:"ImageName,omitnil,omitempty" name:"ImageName"`

	// <p>镜像的描述信息。支持英文或中文字符，最多60个字符。<br>默认值：空。</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>基于实例创建整机镜像时，指定包含在镜像里的数据盘ID。</p>
	DataDiskIds []*string `json:"DataDiskIds,omitnil,omitempty" name:"DataDiskIds"`
}

type CreateImageRequest struct {
	*tchttp.BaseRequest
	
	// <p>云桌面实例ID。</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>镜像名称。只支持中文、字母、数字或连接符号&quot;-&quot;，且最多60个字符。<br>默认值：空。</p>
	ImageName *string `json:"ImageName,omitnil,omitempty" name:"ImageName"`

	// <p>镜像的描述信息。支持英文或中文字符，最多60个字符。<br>默认值：空。</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>基于实例创建整机镜像时，指定包含在镜像里的数据盘ID。</p>
	DataDiskIds []*string `json:"DataDiskIds,omitnil,omitempty" name:"DataDiskIds"`
}

func (r *CreateImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ImageName")
	delete(f, "Description")
	delete(f, "DataDiskIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateImageResponseParams struct {
	// <p>镜像ID。</p>
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateImageResponse struct {
	*tchttp.BaseResponse
	Response *CreateImageResponseParams `json:"Response"`
}

func (r *CreateImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstancesRequestParams struct {
	// 实例类型。<li>  S3L.MEDIUM4: 云桌面-标准S3_2核4G </li><li> S3L.LARGE8: 云桌面-标准S3_4核8G </li><li> S3L.LARGE16: 云桌面-标准S3_4核16G </li><li> S3L.2XLARGE16: 云桌面-标准S3_8核16G </li><li> S1L.LARGE8：云桌面-标准S1_4核8G </li><li> S1L.LARGE16：标准S1_4核16G </li><li> S2L.2XLARGE16：标准S2_8核16G </li><li> G1L.LARGE16：图形G1_4核16G  </li><li> G1L.2XLARGE32：图形G1_8核32G </li><li> G1L.4XLARGE64：图形G1_16核64G </li><li> G2L.LARGE24: 云桌面-图形G2_6核24G </li><li> G2L.3XLARGE58: 云桌面-图形G2_14核58G </li><li> G2L.3XLARGE44: 云桌面-图形G2_12核44G </li><li> G3L.LARGE16: 云桌面-图形G3_4核16G </li><li> G3L.5XMEDIUM40: 云桌面-图形G3_10核40G </li><li> G3L.5XLARGE80: 云桌面-图形G3_20核80G </li><li> G4L.5XMEDIUM40: 云桌面-图形G4_10核40G </li><li> G4L.11XMEDIUM80: 云桌面-图形G4_22核80G </li><li> G4L.11XLARGE160: 云桌面-图形G4_44核160G </li><li> G5L.5XMEDIUM40: 云桌面-图形G5_10核40G </li><li> G5L.5XLARGE80: 云桌面-图形G5_22核80G </li><li> G5L.21XMEDIUM160: 云桌面-图形G5_42核160G </li><li> G5L.21XLARGE320: 云桌面-图形G5_84核320G </li>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 系统盘信息
	SystemDisk *Disk `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// 镜像ID
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 私有网络相关信息配置。通过该参数可以指定私有网络的ID，子网ID等信息。
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitnil,omitempty" name:"VirtualPrivateCloud"`

	// 实例显示名称。
	// <li> 如果该字段为空，则采用默认格式：云桌面型号+月日+三位随机数+镜像ID。</li>
	// <li> 购买多台实例，如果指定模式串{R:x}，表示生成数字[x, x+n-1]，其中n表示购买实例的数量，例如server_{R:3}，购买1台时，实例显示名称为server_3；购买2台时，实例显示名称分别为server_3，server_4。支持指定多个模式串{R:x}。</li>
	// <li> 购买多台实例，如果不指定模式串，则在实例显示名称添加后缀1、2...n，其中n表示购买实例的数量，例如server_，购买2台时，实例显示名称分别为server_1，server_2。</li>
	// <li> 最多支持60个字符（包含模式串）。</li>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 购买实例数量。取值范围[1,50]。默认为1
	InstanceCount *uint64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// 数据盘信息。目前只支持一块数据盘。
	DataDisks []*Disk `json:"DataDisks,omitnil,omitempty" name:"DataDisks"`

	// 实例计费类型。
	// <li> PREPAID：预付费，即包年包月</li>
	// <li> POSTPAID_BY_HOUR：按小时后付费</li>
	// 默认为PREPAID
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	ChargePrepaid *ChargePrepaid `json:"ChargePrepaid,omitnil,omitempty" name:"ChargePrepaid"`

	// 是否只预检此次请求。
	// <li> true：发送检查请求，不会创建实例。检查项包括是否填写了必需参数，请求格式，业务限制和库存。
	// 如果检查通过，则返回的错误码为DryRunOperation；如果返回其他错误码表示检查没有通过；</li>
	// <li> false（默认）：发送正常请求，通过检查后直接创建实例</li>
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 允许绑定用户前的匿名访问。默认值：false。
	AnonymousAccess *bool `json:"AnonymousAccess,omitnil,omitempty" name:"AnonymousAccess"`

	// 关联的安全组ID列表
	SecurityGroups []*string `json:"SecurityGroups,omitnil,omitempty" name:"SecurityGroups"`

	// 桌面池ID
	PoolId *string `json:"PoolId,omitnil,omitempty" name:"PoolId"`

	// 域名管控策略ID
	DomainControlPolicies []*string `json:"DomainControlPolicies,omitnil,omitempty" name:"DomainControlPolicies"`

	// 标签，最多10个
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

type CreateInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例类型。<li>  S3L.MEDIUM4: 云桌面-标准S3_2核4G </li><li> S3L.LARGE8: 云桌面-标准S3_4核8G </li><li> S3L.LARGE16: 云桌面-标准S3_4核16G </li><li> S3L.2XLARGE16: 云桌面-标准S3_8核16G </li><li> S1L.LARGE8：云桌面-标准S1_4核8G </li><li> S1L.LARGE16：标准S1_4核16G </li><li> S2L.2XLARGE16：标准S2_8核16G </li><li> G1L.LARGE16：图形G1_4核16G  </li><li> G1L.2XLARGE32：图形G1_8核32G </li><li> G1L.4XLARGE64：图形G1_16核64G </li><li> G2L.LARGE24: 云桌面-图形G2_6核24G </li><li> G2L.3XLARGE58: 云桌面-图形G2_14核58G </li><li> G2L.3XLARGE44: 云桌面-图形G2_12核44G </li><li> G3L.LARGE16: 云桌面-图形G3_4核16G </li><li> G3L.5XMEDIUM40: 云桌面-图形G3_10核40G </li><li> G3L.5XLARGE80: 云桌面-图形G3_20核80G </li><li> G4L.5XMEDIUM40: 云桌面-图形G4_10核40G </li><li> G4L.11XMEDIUM80: 云桌面-图形G4_22核80G </li><li> G4L.11XLARGE160: 云桌面-图形G4_44核160G </li><li> G5L.5XMEDIUM40: 云桌面-图形G5_10核40G </li><li> G5L.5XLARGE80: 云桌面-图形G5_22核80G </li><li> G5L.21XMEDIUM160: 云桌面-图形G5_42核160G </li><li> G5L.21XLARGE320: 云桌面-图形G5_84核320G </li>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 系统盘信息
	SystemDisk *Disk `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// 镜像ID
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 私有网络相关信息配置。通过该参数可以指定私有网络的ID，子网ID等信息。
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitnil,omitempty" name:"VirtualPrivateCloud"`

	// 实例显示名称。
	// <li> 如果该字段为空，则采用默认格式：云桌面型号+月日+三位随机数+镜像ID。</li>
	// <li> 购买多台实例，如果指定模式串{R:x}，表示生成数字[x, x+n-1]，其中n表示购买实例的数量，例如server_{R:3}，购买1台时，实例显示名称为server_3；购买2台时，实例显示名称分别为server_3，server_4。支持指定多个模式串{R:x}。</li>
	// <li> 购买多台实例，如果不指定模式串，则在实例显示名称添加后缀1、2...n，其中n表示购买实例的数量，例如server_，购买2台时，实例显示名称分别为server_1，server_2。</li>
	// <li> 最多支持60个字符（包含模式串）。</li>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 购买实例数量。取值范围[1,50]。默认为1
	InstanceCount *uint64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// 数据盘信息。目前只支持一块数据盘。
	DataDisks []*Disk `json:"DataDisks,omitnil,omitempty" name:"DataDisks"`

	// 实例计费类型。
	// <li> PREPAID：预付费，即包年包月</li>
	// <li> POSTPAID_BY_HOUR：按小时后付费</li>
	// 默认为PREPAID
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	ChargePrepaid *ChargePrepaid `json:"ChargePrepaid,omitnil,omitempty" name:"ChargePrepaid"`

	// 是否只预检此次请求。
	// <li> true：发送检查请求，不会创建实例。检查项包括是否填写了必需参数，请求格式，业务限制和库存。
	// 如果检查通过，则返回的错误码为DryRunOperation；如果返回其他错误码表示检查没有通过；</li>
	// <li> false（默认）：发送正常请求，通过检查后直接创建实例</li>
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 允许绑定用户前的匿名访问。默认值：false。
	AnonymousAccess *bool `json:"AnonymousAccess,omitnil,omitempty" name:"AnonymousAccess"`

	// 关联的安全组ID列表
	SecurityGroups []*string `json:"SecurityGroups,omitnil,omitempty" name:"SecurityGroups"`

	// 桌面池ID
	PoolId *string `json:"PoolId,omitnil,omitempty" name:"PoolId"`

	// 域名管控策略ID
	DomainControlPolicies []*string `json:"DomainControlPolicies,omitnil,omitempty" name:"DomainControlPolicies"`

	// 标签，最多10个
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

func (r *CreateInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceType")
	delete(f, "SystemDisk")
	delete(f, "ImageId")
	delete(f, "VirtualPrivateCloud")
	delete(f, "InstanceName")
	delete(f, "InstanceCount")
	delete(f, "DataDisks")
	delete(f, "ChargeType")
	delete(f, "ChargePrepaid")
	delete(f, "DryRun")
	delete(f, "AnonymousAccess")
	delete(f, "SecurityGroups")
	delete(f, "PoolId")
	delete(f, "DomainControlPolicies")
	delete(f, "Tags")
	delete(f, "Zone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstancesResponseParams struct {
	// 当通过本接口来创建实例时会返回该参数，表示一个或多个实例ID。返回实例ID列表并不代表实例创建成功。通过进一步查询返回的InstancesSet中对应实例的ID的状态来判断创建是否完成；如果实例状态由“准备中”变为“正在运行”，则为创建成功。
	InstanceIdSet []*string `json:"InstanceIdSet,omitnil,omitempty" name:"InstanceIdSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateInstancesResponse struct {
	*tchttp.BaseResponse
	Response *CreateInstancesResponseParams `json:"Response"`
}

func (r *CreateInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePolicyRequestParams struct {
	// <p>策略名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>策略详情</p>
	PolicyInfoList []*PolicyInfo `json:"PolicyInfoList,omitnil,omitempty" name:"PolicyInfoList"`

	// <p>策略描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>策略作用的用户id</p>
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`
}

type CreatePolicyRequest struct {
	*tchttp.BaseRequest
	
	// <p>策略名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>策略详情</p>
	PolicyInfoList []*PolicyInfo `json:"PolicyInfoList,omitnil,omitempty" name:"PolicyInfoList"`

	// <p>策略描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>策略作用的用户id</p>
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`
}

func (r *CreatePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "PolicyInfoList")
	delete(f, "Description")
	delete(f, "UserIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreatePolicyResponseParams `json:"Response"`
}

func (r *CreatePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataPoint struct {
	// <p>实例对象维度组合</p>
	Dimensions []*Dimension `json:"Dimensions,omitnil,omitempty" name:"Dimensions"`

	// <p>时间戳数组</p>
	Timestamps []*float64 `json:"Timestamps,omitnil,omitempty" name:"Timestamps"`

	// <p>监控值数组，与Timestamps一一对应</p>
	Values []*float64 `json:"Values,omitnil,omitempty" name:"Values"`

	// <p>MaxValues 最大值数组，与Timestamps一一对应</p>
	MaxValues []*float64 `json:"MaxValues,omitnil,omitempty" name:"MaxValues"`

	// <p>MinValues 最小值数组，与Timestamps一一对应</p>
	MinValues []*float64 `json:"MinValues,omitnil,omitempty" name:"MinValues"`

	// <p>AvgValues 平均值数组，与Timestamps一一对应</p>
	AvgValues []*float64 `json:"AvgValues,omitnil,omitempty" name:"AvgValues"`
}

// Predefined struct for user
type DeleteCommandTasksRequestParams struct {
	// <p>删除命令任务条件<br>Filter.name 目前支持 CommandId</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DeleteCommandTasksRequest struct {
	*tchttp.BaseRequest
	
	// <p>删除命令任务条件<br>Filter.name 目前支持 CommandId</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DeleteCommandTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCommandTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCommandTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCommandTasksResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCommandTasksResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCommandTasksResponseParams `json:"Response"`
}

func (r *DeleteCommandTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCommandTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDomainControlPolicyRequestParams struct {
	// 策略唯一ID
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`
}

type DeleteDomainControlPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 策略唯一ID
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`
}

func (r *DeleteDomainControlPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDomainControlPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDomainControlPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDomainControlPolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDomainControlPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDomainControlPolicyResponseParams `json:"Response"`
}

func (r *DeleteDomainControlPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDomainControlPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImagesRequestParams struct {
	// <p>准备删除的镜像Id列表</p>
	ImageIds []*string `json:"ImageIds,omitnil,omitempty" name:"ImageIds"`
}

type DeleteImagesRequest struct {
	*tchttp.BaseRequest
	
	// <p>准备删除的镜像Id列表</p>
	ImageIds []*string `json:"ImageIds,omitnil,omitempty" name:"ImageIds"`
}

func (r *DeleteImagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteImagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImagesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteImagesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteImagesResponseParams `json:"Response"`
}

func (r *DeleteImagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCommandTasksRequestParams struct {
	// <p>页号</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页大小，默认 20，最大 100</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>拉取任务的过滤器 ，Filter.Name 取值：CommandId - 任务 lD；Staus - 任务状态， RUNNING - 执行中；FINISHED - 执行完成（不代表成功，只表示任务执行完成）</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeCommandTasksRequest struct {
	*tchttp.BaseRequest
	
	// <p>页号</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页大小，默认 20，最大 100</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>拉取任务的过滤器 ，Filter.Name 取值：CommandId - 任务 lD；Staus - 任务状态， RUNNING - 执行中；FINISHED - 执行完成（不代表成功，只表示任务执行完成）</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeCommandTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCommandTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCommandTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCommandTasksResponseParams struct {
	// <p>拉取任务详情数</p>
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// <p>任务描述</p>
	CommandDetails []*CommandDetail `json:"CommandDetails,omitnil,omitempty" name:"CommandDetails"`

	// <p>任务状态，RUNNING -执行中；FINISHED-执行完成； STOPPED-终止任务</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCommandTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCommandTasksResponseParams `json:"Response"`
}

func (r *DescribeCommandTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCommandTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImagesRequestParams struct {
	// <p>拉取页号</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页拉取个数</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>Name取值：[&quot;ImageID&quot;,&quot;Name&quot;,&quot;ImageType&quot;,&quot;ImageState&quot;,&quot;Region&quot;,&quot;InstanceType&quot;]<br>注意：filter中没有ImageID或ImageState过滤条件，后台默认不返回INVISIBLE的镜像信息</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeImagesRequest struct {
	*tchttp.BaseRequest
	
	// <p>拉取页号</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页拉取个数</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>Name取值：[&quot;ImageID&quot;,&quot;Name&quot;,&quot;ImageType&quot;,&quot;ImageState&quot;,&quot;Region&quot;,&quot;InstanceType&quot;]<br>注意：filter中没有ImageID或ImageState过滤条件，后台默认不返回INVISIBLE的镜像信息</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeImagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImagesResponseParams struct {
	// <p>返回查询到的所有镜像信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Images []*ImageInfo `json:"Images,omitnil,omitempty" name:"Images"`

	// <p>返回当前查询到的镜像总数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeImagesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImagesResponseParams `json:"Response"`
}

func (r *DescribeImagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceFetchMsgRequestParams struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 绑定实例的UserID，为后续多会话预留参数
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 拉取桌面title，用户可自行定义，默认为云桌面实例名称
	DesktopName *string `json:"DesktopName,omitnil,omitempty" name:"DesktopName"`

	// 是否抢占桌面，即桌面使用过程中，抢占登录。
	// true：抢占登录
	// false：非抢占登录
	// 默认为抢占登录
	Force *bool `json:"Force,omitnil,omitempty" name:"Force"`

	// 是否内网登录桌面
	IsInternal *bool `json:"IsInternal,omitnil,omitempty" name:"IsInternal"`
}

type DescribeInstanceFetchMsgRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 绑定实例的UserID，为后续多会话预留参数
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 拉取桌面title，用户可自行定义，默认为云桌面实例名称
	DesktopName *string `json:"DesktopName,omitnil,omitempty" name:"DesktopName"`

	// 是否抢占桌面，即桌面使用过程中，抢占登录。
	// true：抢占登录
	// false：非抢占登录
	// 默认为抢占登录
	Force *bool `json:"Force,omitnil,omitempty" name:"Force"`

	// 是否内网登录桌面
	IsInternal *bool `json:"IsInternal,omitnil,omitempty" name:"IsInternal"`
}

func (r *DescribeInstanceFetchMsgRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceFetchMsgRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserId")
	delete(f, "DesktopName")
	delete(f, "Force")
	delete(f, "IsInternal")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceFetchMsgRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceFetchMsgResponseParams struct {
	// 桌面登录ticket状态描述
	FetchStatusTicket *string `json:"FetchStatusTicket,omitnil,omitempty" name:"FetchStatusTicket"`

	// 桌面登录ticket描述
	FetchTicket *string `json:"FetchTicket,omitnil,omitempty" name:"FetchTicket"`

	// 桌面登录launch url
	FetchUrl *string `json:"FetchUrl,omitnil,omitempty" name:"FetchUrl"`

	// 桌面传输协议
	ProtocolVersion *string `json:"ProtocolVersion,omitnil,omitempty" name:"ProtocolVersion"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceFetchMsgResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceFetchMsgResponseParams `json:"Response"`
}

func (r *DescribeInstanceFetchMsgResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceFetchMsgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceGraphDataRequestParams struct {
	// 云桌面实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 查询起始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeInstanceGraphDataRequest struct {
	*tchttp.BaseRequest
	
	// 云桌面实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 查询起始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *DescribeInstanceGraphDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceGraphDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceGraphDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceGraphDataResponseParams struct {
	// 查询起始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 周期
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 指标名
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// 实例磁盘数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Partitions []*Partition `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceGraphDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceGraphDataResponseParams `json:"Response"`
}

func (r *DescribeInstanceGraphDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceGraphDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceLoginMsgRequestParams struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 绑定实例的UserID，为后续多会话预留参数
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 拉取桌面title，用户可自行定义，默认为云桌面实例名称
	DesktopName *string `json:"DesktopName,omitnil,omitempty" name:"DesktopName"`

	// 是否抢占桌面，即桌面使用过程中，抢占登录。
	// true：抢占登录
	// false：非抢占登录
	// 默认为抢占登录
	Force *bool `json:"Force,omitnil,omitempty" name:"Force"`

	// 是否内网方式
	IsInternal *bool `json:"IsInternal,omitnil,omitempty" name:"IsInternal"`
}

type DescribeInstanceLoginMsgRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 绑定实例的UserID，为后续多会话预留参数
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 拉取桌面title，用户可自行定义，默认为云桌面实例名称
	DesktopName *string `json:"DesktopName,omitnil,omitempty" name:"DesktopName"`

	// 是否抢占桌面，即桌面使用过程中，抢占登录。
	// true：抢占登录
	// false：非抢占登录
	// 默认为抢占登录
	Force *bool `json:"Force,omitnil,omitempty" name:"Force"`

	// 是否内网方式
	IsInternal *bool `json:"IsInternal,omitnil,omitempty" name:"IsInternal"`
}

func (r *DescribeInstanceLoginMsgRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceLoginMsgRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserId")
	delete(f, "DesktopName")
	delete(f, "Force")
	delete(f, "IsInternal")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceLoginMsgRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceLoginMsgResponseParams struct {
	// 桌面登录login信息
	LoginMsg *string `json:"LoginMsg,omitnil,omitempty" name:"LoginMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceLoginMsgResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceLoginMsgResponseParams `json:"Response"`
}

func (r *DescribeInstanceLoginMsgResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceLoginMsgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorDataRequestParams struct {
	// <p>实例维度组合数组</p>
	Instances []*Instance `json:"Instances,omitnil,omitempty" name:"Instances"`

	// <p>命名空间。如QCE/CVM、QCE/BLOCK_STORAGE等</p>
	NameSpace *string `json:"NameSpace,omitnil,omitempty" name:"NameSpace"`

	// <p>指标名</p>
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// <p>统计周期(时间粒度), 单位s。</p>
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// <p>开始时间</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>多种统计方式数据。avg, max, min (1,2,4)可以自由组合</p>
	SpecifyStatistics *int64 `json:"SpecifyStatistics,omitnil,omitempty" name:"SpecifyStatistics"`
}

type DescribeInstanceMonitorDataRequest struct {
	*tchttp.BaseRequest
	
	// <p>实例维度组合数组</p>
	Instances []*Instance `json:"Instances,omitnil,omitempty" name:"Instances"`

	// <p>命名空间。如QCE/CVM、QCE/BLOCK_STORAGE等</p>
	NameSpace *string `json:"NameSpace,omitnil,omitempty" name:"NameSpace"`

	// <p>指标名</p>
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// <p>统计周期(时间粒度), 单位s。</p>
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// <p>开始时间</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>多种统计方式数据。avg, max, min (1,2,4)可以自由组合</p>
	SpecifyStatistics *int64 `json:"SpecifyStatistics,omitnil,omitempty" name:"SpecifyStatistics"`
}

func (r *DescribeInstanceMonitorDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Instances")
	delete(f, "NameSpace")
	delete(f, "MetricName")
	delete(f, "Period")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SpecifyStatistics")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceMonitorDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorDataResponseParams struct {
	// <p>指标名</p>
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// <p>统计周期</p>
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// <p>开始时间</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>数据点数组</p>
	DataPoints []*DataPoint `json:"DataPoints,omitnil,omitempty" name:"DataPoints"`

	// <p>message信息</p>
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceMonitorDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceMonitorDataResponseParams `json:"Response"`
}

func (r *DescribeInstanceMonitorDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceRegionDistributionRequestParams struct {
	// <p>查询区域列表</p>
	Regions []*string `json:"Regions,omitnil,omitempty" name:"Regions"`

	// <p>查询实例类型列表</p>
	Types []*string `json:"Types,omitnil,omitempty" name:"Types"`

	// <p>查询用户拥有的云桌面地域分布信息</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type DescribeInstanceRegionDistributionRequest struct {
	*tchttp.BaseRequest
	
	// <p>查询区域列表</p>
	Regions []*string `json:"Regions,omitnil,omitempty" name:"Regions"`

	// <p>查询实例类型列表</p>
	Types []*string `json:"Types,omitnil,omitempty" name:"Types"`

	// <p>查询用户拥有的云桌面地域分布信息</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *DescribeInstanceRegionDistributionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceRegionDistributionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Regions")
	delete(f, "Types")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceRegionDistributionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceRegionDistributionResponseParams struct {
	// <p>各区域实例个数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DistributionInfo []*DistributionInfo `json:"DistributionInfo,omitnil,omitempty" name:"DistributionInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceRegionDistributionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceRegionDistributionResponseParams `json:"Response"`
}

func (r *DescribeInstanceRegionDistributionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceRegionDistributionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesBindStatusRequestParams struct {
	// <p>拉取的桌面id列表</p>
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// <p>实例所在区域</p>
	InstanceRegion *string `json:"InstanceRegion,omitnil,omitempty" name:"InstanceRegion"`
}

type DescribeInstancesBindStatusRequest struct {
	*tchttp.BaseRequest
	
	// <p>拉取的桌面id列表</p>
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// <p>实例所在区域</p>
	InstanceRegion *string `json:"InstanceRegion,omitnil,omitempty" name:"InstanceRegion"`
}

func (r *DescribeInstancesBindStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesBindStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "InstanceRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesBindStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesBindStatusResponseParams struct {
	// <p>桌面实例状态列表</p>
	InstanceBindStatusList []*BindStatusInfo `json:"InstanceBindStatusList,omitnil,omitempty" name:"InstanceBindStatusList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstancesBindStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesBindStatusResponseParams `json:"Response"`
}

func (r *DescribeInstancesBindStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesBindStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesRequestParams struct {
	// 拉取页号
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页拉取个数
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Name取值范围："InstanceId"      按照云桌面ID搜索，云桌面ID示例：cdi-l0dfy594htvbn37"Name"     按照云桌面名称搜索 "Region"     按照云桌面所在地域搜索，目前云桌面所支持的地域：        ap-beijing：北京        ap-shanghai：上海        ap-guangzhou：广州        ap-shanghai-adc：上海自动驾驶云"Type"    按照云桌面类型搜索，类型取值：        S3L.MEDIUM4：云桌面-标准S3_2核4G        S3L.LARGE8：云桌面-标准S3_4核8G        S3L.LARGE16：云桌面-标准S3_4核16G        S3L.2XLARGE16：云桌面-标准S3_8核16G        S1L.LARGE8：云桌面-标准S1_4核8G        S1L.LARGE16：云桌面-标准S1_4核16G        S2L.2XLARGE16：云桌面-标准S2_8核16G        G1L.LARGE16：云桌面-图形G1_4核16G        G1L.2XLARGE32：云桌面-图形G1_8核32G        G1L.4XLARGE64：云桌面-图形G1_16核64G        G2L.LARGE24：云桌面-图形G2_6核24G        G2L.3XLARGE58：云桌面-图形G2_14核58G	        G2L.3XLARGE44：云桌面-图形G2_12核44G         G3L.LARGE16：云桌面-图形G3_4核16G	        G3L.5XMEDIUM40：云桌面-图形G3_10核40G	        G3L.5XLARGE80：云桌面-图形G3_20核80G        G4L.5XMEDIUM40：云桌面-图形G4_10核40G        G4L.11XMEDIUM80：云桌面-图形G4_22核80G        G4L.11XLARGE160：云桌面-图形G4_44核160G        G5L.5XMEDIUM40：云桌面-图形G5_10核40G	        G5L.5XLARGE80：云桌面-图形G5_22核80G        G5L.21XMEDIUM160：云桌面-图形G5_42核160G        G5L.21XLARGE320：云桌面-图形G5_84核320G	"ImageId"    按照镜像ID搜索，镜像ID示例：img-ed892xcd2"ChargeType"    按着计费模式搜索，目前支持的计费模式为：         PREPAID：包年包月        POSTPAID_BY_HOUR：按量计费 "HostName"    按照云桌面机器名称搜索，云桌面机器名示例：l0dfy594htvbn37"Status"    按照桌面状态搜索，云桌面状态：        PENDING：创建中        STARTING：开机中        REBOOTING：重启中        ROLLBACKING：回滚中        RESETTING：重置中        STOPPING：关机中        TERMINATING：销毁中        RUNNING：运行中        USING：使用中        STOPPED：关机        ISOLATING ：隔离中，只返回最近24小时数据        SHUTDOWN：销毁，只返回最近24小时数据        DELETED：销毁和SHUTDOWN意义一致，只返回最近24小时数据        LAUNCH_FAILED：创建失败，只返回最近24小时数据"UserId"    按照绑定的用户ID搜索，用户ID示例: us-6b36f283c302481bb83bsdswdefqdqwerqw"UserVpcId"    按照云桌面所属vpc搜索，VPCID示例: vpc-0ji679xf"UserSubnetId"    按照云桌面所属子网搜索，SubnetID示例: subnet-0ji679xf"UserPrivateIp"    按照云桌面IP搜索"Keyword"    按照云桌面绑定的用户名搜索，Values每次只取第一个值"SecurityGroupId"    按照云桌面关联安全组搜索，Values每次只取第一个值,安全组ID实例：sg-0ji679xf"AssociateSG"    按照可以关联安全组搜索，Values每次只取第一个值,安全组ID实例：sg-0ji679xf"DomainControlPolicyId"    按照云桌面关联域名管控策略搜索，Values每次只取第一个值,安全组ID实例：domain-0ji679xf "AssociateDomainControlPolicy"    按照可以关联域名策略搜索        ①若已有域名策略关联桌面情况：Values每次只取第一个值,安全组ID实例：domain-0ji679xf         ②若新建域名策略关联桌面情况，Values传空"Uin"    外部租户请勿使用"AppId"    外部租户请勿使用"ImageName"    废弃
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 拉取页号
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页拉取个数
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Name取值范围："InstanceId"      按照云桌面ID搜索，云桌面ID示例：cdi-l0dfy594htvbn37"Name"     按照云桌面名称搜索 "Region"     按照云桌面所在地域搜索，目前云桌面所支持的地域：        ap-beijing：北京        ap-shanghai：上海        ap-guangzhou：广州        ap-shanghai-adc：上海自动驾驶云"Type"    按照云桌面类型搜索，类型取值：        S3L.MEDIUM4：云桌面-标准S3_2核4G        S3L.LARGE8：云桌面-标准S3_4核8G        S3L.LARGE16：云桌面-标准S3_4核16G        S3L.2XLARGE16：云桌面-标准S3_8核16G        S1L.LARGE8：云桌面-标准S1_4核8G        S1L.LARGE16：云桌面-标准S1_4核16G        S2L.2XLARGE16：云桌面-标准S2_8核16G        G1L.LARGE16：云桌面-图形G1_4核16G        G1L.2XLARGE32：云桌面-图形G1_8核32G        G1L.4XLARGE64：云桌面-图形G1_16核64G        G2L.LARGE24：云桌面-图形G2_6核24G        G2L.3XLARGE58：云桌面-图形G2_14核58G	        G2L.3XLARGE44：云桌面-图形G2_12核44G         G3L.LARGE16：云桌面-图形G3_4核16G	        G3L.5XMEDIUM40：云桌面-图形G3_10核40G	        G3L.5XLARGE80：云桌面-图形G3_20核80G        G4L.5XMEDIUM40：云桌面-图形G4_10核40G        G4L.11XMEDIUM80：云桌面-图形G4_22核80G        G4L.11XLARGE160：云桌面-图形G4_44核160G        G5L.5XMEDIUM40：云桌面-图形G5_10核40G	        G5L.5XLARGE80：云桌面-图形G5_22核80G        G5L.21XMEDIUM160：云桌面-图形G5_42核160G        G5L.21XLARGE320：云桌面-图形G5_84核320G	"ImageId"    按照镜像ID搜索，镜像ID示例：img-ed892xcd2"ChargeType"    按着计费模式搜索，目前支持的计费模式为：         PREPAID：包年包月        POSTPAID_BY_HOUR：按量计费 "HostName"    按照云桌面机器名称搜索，云桌面机器名示例：l0dfy594htvbn37"Status"    按照桌面状态搜索，云桌面状态：        PENDING：创建中        STARTING：开机中        REBOOTING：重启中        ROLLBACKING：回滚中        RESETTING：重置中        STOPPING：关机中        TERMINATING：销毁中        RUNNING：运行中        USING：使用中        STOPPED：关机        ISOLATING ：隔离中，只返回最近24小时数据        SHUTDOWN：销毁，只返回最近24小时数据        DELETED：销毁和SHUTDOWN意义一致，只返回最近24小时数据        LAUNCH_FAILED：创建失败，只返回最近24小时数据"UserId"    按照绑定的用户ID搜索，用户ID示例: us-6b36f283c302481bb83bsdswdefqdqwerqw"UserVpcId"    按照云桌面所属vpc搜索，VPCID示例: vpc-0ji679xf"UserSubnetId"    按照云桌面所属子网搜索，SubnetID示例: subnet-0ji679xf"UserPrivateIp"    按照云桌面IP搜索"Keyword"    按照云桌面绑定的用户名搜索，Values每次只取第一个值"SecurityGroupId"    按照云桌面关联安全组搜索，Values每次只取第一个值,安全组ID实例：sg-0ji679xf"AssociateSG"    按照可以关联安全组搜索，Values每次只取第一个值,安全组ID实例：sg-0ji679xf"DomainControlPolicyId"    按照云桌面关联域名管控策略搜索，Values每次只取第一个值,安全组ID实例：domain-0ji679xf "AssociateDomainControlPolicy"    按照可以关联域名策略搜索        ①若已有域名策略关联桌面情况：Values每次只取第一个值,安全组ID实例：domain-0ji679xf         ②若新建域名策略关联桌面情况，Values传空"Uin"    外部租户请勿使用"AppId"    外部租户请勿使用"ImageName"    废弃
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesResponseParams struct {
	// 返回查询到的所有实例信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceList []*InstanceInfo `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// 返回当前查询到的实例总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesResponseParams `json:"Response"`
}

func (r *DescribeInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesStatusRequestParams struct {
	// <p>查询的实例id列表</p>
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type DescribeInstancesStatusRequest struct {
	*tchttp.BaseRequest
	
	// <p>查询的实例id列表</p>
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *DescribeInstancesStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesStatusResponseParams struct {
	// <p>实例状态列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceStatusList []*InstanceStatus `json:"InstanceStatusList,omitnil,omitempty" name:"InstanceStatusList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstancesStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesStatusResponseParams `json:"Response"`
}

func (r *DescribeInstancesStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInvocationDetailRequestParams struct {
	// <p>云桌面实例ID。</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>任务ID。</p>
	InvokeId *string `json:"InvokeId,omitnil,omitempty" name:"InvokeId"`

	// <p>需要返回的任务输出行号起始索引。</p>
	OutputLinesOffset *uint64 `json:"OutputLinesOffset,omitnil,omitempty" name:"OutputLinesOffset"`

	// <p>需要返回的任务输出行数。</p>
	OutputLinesLimit *uint64 `json:"OutputLinesLimit,omitnil,omitempty" name:"OutputLinesLimit"`

	// <p>任务的输出内容是否合并成一个字符串。默认为true。</p>
	MergeOutputLines *bool `json:"MergeOutputLines,omitnil,omitempty" name:"MergeOutputLines"`
}

type DescribeInvocationDetailRequest struct {
	*tchttp.BaseRequest
	
	// <p>云桌面实例ID。</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>任务ID。</p>
	InvokeId *string `json:"InvokeId,omitnil,omitempty" name:"InvokeId"`

	// <p>需要返回的任务输出行号起始索引。</p>
	OutputLinesOffset *uint64 `json:"OutputLinesOffset,omitnil,omitempty" name:"OutputLinesOffset"`

	// <p>需要返回的任务输出行数。</p>
	OutputLinesLimit *uint64 `json:"OutputLinesLimit,omitnil,omitempty" name:"OutputLinesLimit"`

	// <p>任务的输出内容是否合并成一个字符串。默认为true。</p>
	MergeOutputLines *bool `json:"MergeOutputLines,omitnil,omitempty" name:"MergeOutputLines"`
}

func (r *DescribeInvocationDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInvocationDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InvokeId")
	delete(f, "OutputLinesOffset")
	delete(f, "OutputLinesLimit")
	delete(f, "MergeOutputLines")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInvocationDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInvocationDetailResponseParams struct {
	// <p>实例上任务执行的详细情况。</p>
	Detail *Invocation `json:"Detail,omitnil,omitempty" name:"Detail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInvocationDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInvocationDetailResponseParams `json:"Response"`
}

func (r *DescribeInvocationDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInvocationDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInvocationsRequestParams struct {
	// <p>云桌面实例ID。</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInvocationsRequest struct {
	*tchttp.BaseRequest
	
	// <p>云桌面实例ID。</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInvocationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInvocationsResponseParams struct {
	// <p>该实例上的任务列表。不会返回Output字段。</p>
	Invocations []*Invocation `json:"Invocations,omitnil,omitempty" name:"Invocations"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribePolicyDetailRequestParams struct {
	// <p>拉取页号</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页最大显示个数，最大值不超过100</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>查询条件（暂时无效，保留字段）</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribePolicyDetailRequest struct {
	*tchttp.BaseRequest
	
	// <p>拉取页号</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页最大显示个数，最大值不超过100</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>查询条件（暂时无效，保留字段）</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribePolicyDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePolicyDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePolicyDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePolicyDetailResponseParams struct {
	// <p>策略详细信息列表</p>
	PolicyList []*PolicyDetail `json:"PolicyList,omitnil,omitempty" name:"PolicyList"`

	// <p>当前查询到的策略总数</p>
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePolicyDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribePolicyDetailResponseParams `json:"Response"`
}

func (r *DescribePolicyDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePolicyDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupsRequestParams struct {
	// <p>页号，从1开始</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页最大展示数量，默认20，最大100</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>地域</p>
	SecurityGroupRegion *string `json:"SecurityGroupRegion,omitnil,omitempty" name:"SecurityGroupRegion"`
}

type DescribeSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// <p>页号，从1开始</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页最大展示数量，默认20，最大100</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>地域</p>
	SecurityGroupRegion *string `json:"SecurityGroupRegion,omitnil,omitempty" name:"SecurityGroupRegion"`
}

func (r *DescribeSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "SecurityGroupRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityGroupsResponseParams `json:"Response"`
}

func (r *DescribeSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotsRequestParams struct {
	// <p>快照所在的地域</p>
	SnapshotRegion *string `json:"SnapshotRegion,omitnil,omitempty" name:"SnapshotRegion"`

	// <p>页号</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页大小，最大100，默认20</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>检索条件，Filter的name取值：<br>InstanceId：实例ID，返回当前实例挂载的云硬盘关联的快照信息<br>DiskId：云硬盘ID，返回当前云硬盘关联的快照信息</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeSnapshotsRequest struct {
	*tchttp.BaseRequest
	
	// <p>快照所在的地域</p>
	SnapshotRegion *string `json:"SnapshotRegion,omitnil,omitempty" name:"SnapshotRegion"`

	// <p>页号</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页大小，最大100，默认20</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>检索条件，Filter的name取值：<br>InstanceId：实例ID，返回当前实例挂载的云硬盘关联的快照信息<br>DiskId：云硬盘ID，返回当前云硬盘关联的快照信息</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeSnapshotsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotRegion")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSnapshotsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotsResponseParams struct {
	// <p>当前检索条件可查询到的快照总数</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>快照明细列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SnapshotList []*SnapshotInfo `json:"SnapshotList,omitnil,omitempty" name:"SnapshotList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSnapshotsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSnapshotsResponseParams `json:"Response"`
}

func (r *DescribeSnapshotsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsageRequestParams struct {
	// 查询起始时间，最小不能超过当前时间 90 天前，取 unix 时间戳
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询截止时间，取 unix 时间戳
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 每页查询信息条数，默认 20，最大 100
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页号，起始页为 1，不填默认为 1
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 查询key，每次查询会返回
	QueryKey *string `json:"QueryKey,omitnil,omitempty" name:"QueryKey"`

	// 查询条件，Name可选为：
	// InstanceId按照实例ID搜索
	// UserName按照绑定的用户搜索，只取Values的第一个值
	// ChareType按照计费类型搜索
	// InstanceType按照实例类型搜索
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 指定的排序方式，默认不排序
	// 1:使用时间，升序
	// 2:使用时间，降序
	// 3:空闲时间，升序
	// 4:空闲时间，降序
	SortType *int64 `json:"SortType,omitnil,omitempty" name:"SortType"`
}

type DescribeUsageRequest struct {
	*tchttp.BaseRequest
	
	// 查询起始时间，最小不能超过当前时间 90 天前，取 unix 时间戳
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询截止时间，取 unix 时间戳
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 每页查询信息条数，默认 20，最大 100
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页号，起始页为 1，不填默认为 1
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 查询key，每次查询会返回
	QueryKey *string `json:"QueryKey,omitnil,omitempty" name:"QueryKey"`

	// 查询条件，Name可选为：
	// InstanceId按照实例ID搜索
	// UserName按照绑定的用户搜索，只取Values的第一个值
	// ChareType按照计费类型搜索
	// InstanceType按照实例类型搜索
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 指定的排序方式，默认不排序
	// 1:使用时间，升序
	// 2:使用时间，降序
	// 3:空闲时间，升序
	// 4:空闲时间，降序
	SortType *int64 `json:"SortType,omitnil,omitempty" name:"SortType"`
}

func (r *DescribeUsageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "QueryKey")
	delete(f, "Filters")
	delete(f, "SortType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUsageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsageResponseParams struct {
	// 检索出相关使用详情总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 查询 key，作为翻页查询入参
	QueryKey *string `json:"QueryKey,omitnil,omitempty" name:"QueryKey"`

	// 使用时长详情
	UsageSet []*Usage `json:"UsageSet,omitnil,omitempty" name:"UsageSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUsageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUsageResponseParams `json:"Response"`
}

func (r *DescribeUsageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsersRequestParams struct {
	// <p>搜索关键字。</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>页面编号。从1开始</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>页面大小。默认20，不超过100。</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>部门ID。</p>
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`
}

type DescribeUsersRequest struct {
	*tchttp.BaseRequest
	
	// <p>搜索关键字。</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>页面编号。从1开始</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>页面大小。默认20，不超过100。</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>部门ID。</p>
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`
}

func (r *DescribeUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Keyword")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "DepartmentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsersResponseParams struct {
	// <p>符合条件的用户个数。</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>用户列表。</p>
	Users []*User `json:"Users,omitnil,omitempty" name:"Users"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUsersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUsersResponseParams `json:"Response"`
}

func (r *DescribeUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Dimension struct {
	// 实例维度名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 实例维度值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type DisassociateSecurityGroupRequestParams struct {

}

type DisassociateSecurityGroupRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DisassociateSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateSecurityGroupResponseParams struct {
	// 解关联失败的云桌面列表
	FailureSet []*string `json:"FailureSet,omitnil,omitempty" name:"FailureSet"`

	// 解关联成功的云桌面列表
	SuccessSet []*string `json:"SuccessSet,omitnil,omitempty" name:"SuccessSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisassociateSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateSecurityGroupResponseParams `json:"Response"`
}

func (r *DisassociateSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Disk struct {
	// <p>磁盘类型：<br>CLOUD_SSD：SSD云硬盘<br>CLOUD_PREMIUM：高性能云硬盘</p>
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// <p>磁盘大小（单位G）<br>系统盘最小值50，最大值500<br>数据盘最小值10，最大值32000。最小调整步长为10G</p>
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// <p>磁盘id</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`
}

type DistributionInfo struct {
	// <p>地域描述</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>实例个数</p>
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type DomainControlDetail struct {
	// 配置域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 此参数必填，新版本无异议，请求包和回包强制填写空字符串,做校验
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// 规则描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type Filter struct {
	// 要匹配的参数名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 要匹配的参数值
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type ImageInfo struct {
	// <p>镜像id</p>
	ImageID *string `json:"ImageID,omitnil,omitempty" name:"ImageID"`

	// <p>镜像名称</p>
	ImageName *string `json:"ImageName,omitnil,omitempty" name:"ImageName"`

	// <p>镜像类型：<br>PUBLIC_IMAGE-公共镜像<br>PRIVATE_IMAGE-自定义镜像<br>TENCENT_OFFICE_IMAGE-腾讯办公镜像</p>
	ImageType *string `json:"ImageType,omitnil,omitempty" name:"ImageType"`

	// <p>镜像状态：<br>CREATING-创建中<br>NORMAL-正常<br>CREATEFAILED-创建失败<br>USING-使用中<br>SYNCING-同步中<br>IMPORTING-导入中<br>IMPORTFAILED-导入失败<br>INVISIBLE-淘汰的镜像<br>DELETIING-删除中<br>DELETED-删除的</p>
	ImageState *string `json:"ImageState,omitnil,omitempty" name:"ImageState"`

	// <p>镜像操作系统</p>
	OsName *string `json:"OsName,omitnil,omitempty" name:"OsName"`

	// <p>镜像大小（单位G）</p>
	ImageSize *int64 `json:"ImageSize,omitnil,omitempty" name:"ImageSize"`

	// <p>镜像描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>镜像所在区域</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>内核类型：<br>CPU<br>GPU<br>GPU_A10<br>GPU_RTX6000<br>GPU_2080Ti<br>GPU_3080Ti</p>
	CoreType *string `json:"CoreType,omitnil,omitempty" name:"CoreType"`

	// <p>cvm侧镜像id</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CvmImageId *string `json:"CvmImageId,omitnil,omitempty" name:"CvmImageId"`

	// <p>是否是整机镜像，整机镜像无法用于桌面重置<br>true：整机镜像<br>false: 不是整机镜像</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsWhole *bool `json:"IsWhole,omitnil,omitempty" name:"IsWhole"`

	// <p>整机镜像中数据盘大小</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataDiskSize *int64 `json:"DataDiskSize,omitnil,omitempty" name:"DataDiskSize"`

	// <p>镜像创建时间</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type Instance struct {
	// 实例的维度组合
	Dimensions []*Dimension `json:"Dimensions,omitnil,omitempty" name:"Dimensions"`
}

type InstanceInfo struct {
	// <p>实例id</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>实例名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>实例状态<br>PENDING：表示创建中<br>LAUNCH_FAILED：表示创建失败<br>RUNNING：表示运行中<br>STOPPED：表示关机<br>STARTING：表示开机中<br>STOPPING：表示关机中<br>REBOOTING：表示重启中<br>SHUTDOWN：表示停止待销毁<br>TERMINATING：表示销毁中<br>DELETED：表示资源已经销毁，这个状态是业务层面的状态，cvm如果销毁了，在cvm侧资源已经不存在了<br>FAILED：表示资源创建失败.这个是业务层面的状态<br>ROLLBACKING：表示回滚中<br>USING：表示使用中<br>RESETTING：表示重置中.</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>实例拥有者名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// <p>实例</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>实例规格<br>S1.LARGE8：4VCPU 2.5GHZ主频 8G内存 无显卡<br>S1.LARGE16：4VCPU 2.5GHZ主频 16G内存 无显卡<br>S2.2XLARGE16：8VCPU 3.2GHZ主频 16G内存 无显卡<br>G1.LARGE16：4VCPU 2.5GHZ主频 16G内存 1/4 NVIDIA T4<br>G1.2XLARGE32：8VCPU 2.5GHZ主频 32G内存 1/2 NVIDIA T4<br>G1.4XLARGE64：16VCPU 2.5GHZ主频 64G内存 1 NVIDIA T4</p>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// <p>系统盘信息</p>
	SysDisk *Disk `json:"SysDisk,omitnil,omitempty" name:"SysDisk"`

	// <p>挂载的数据盘信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataDiskList []*Disk `json:"DataDiskList,omitnil,omitempty" name:"DataDiskList"`

	// <p>桌面实例所在的区域</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>桌面计费类型</p>
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// <p>自动续费标识。取值范围：<br>NOTIFY_AND_MANUAL_RENEW：表示通知即将过期，但不自动续费<br>NOTIFY_AND_AUTO_RENEW：表示通知即将过期，而且自动续费<br>DISABLE_NOTIFY_AND_MANUAL_RENEW：表示不通知即将过期，也不自动续费。</p><p>注意：后付费模式本项为null</p>
	AutoRenewFlag *string `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// <p>桌面创建时间</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>桌面过期时间</p>
	ExpiredTime *string `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// <p>桌面描述</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>桌面hostname</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostName *string `json:"HostName,omitnil,omitempty" name:"HostName"`

	// <p>桌面拥有者显示名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// <p>桌面绑定状态。取值：<br>UNBOUND - 未绑定<br>BINDING - 绑定中<br>BOUND - 绑定</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindStatus *string `json:"BindStatus,omitnil,omitempty" name:"BindStatus"`

	// <p>创建桌面的镜像id</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// <p>用户桌面vpc id</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserVpcId *string `json:"UserVpcId,omitnil,omitempty" name:"UserVpcId"`

	// <p>用户桌面子网 id</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserSubnetId *string `json:"UserSubnetId,omitnil,omitempty" name:"UserSubnetId"`

	// <p>用户桌面ip</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserPrivateIp *string `json:"UserPrivateIp,omitnil,omitempty" name:"UserPrivateIp"`

	// <p>机器回收站到期时间，只有机器状态为SHUTDOWN该值才有意义</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsolatedTime *string `json:"IsolatedTime,omitnil,omitempty" name:"IsolatedTime"`

	// <p>最近一次状态变更原因, AUTO:被动变更,如欠费,如包月实例到期, MANUAL:主动变更,如主动退还实例，当前只有机器处于ISOLATING、SHUTDOWN、TERMINATING、DELETED该字段有意义</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastStatusChangeReason *string `json:"LastStatusChangeReason,omitnil,omitempty" name:"LastStatusChangeReason"`

	// <p>机器最近一次操作</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestOperation *string `json:"LatestOperation,omitnil,omitempty" name:"LatestOperation"`

	// <p>机器最近一次操作状态</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestOperationState *string `json:"LatestOperationState,omitnil,omitempty" name:"LatestOperationState"`

	// <p>机器最近一次操作请求ReqId</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestOperationRequestID *string `json:"LatestOperationRequestID,omitnil,omitempty" name:"LatestOperationRequestID"`

	// <p>机器操作系统名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	OsName *string `json:"OsName,omitnil,omitempty" name:"OsName"`

	// <p>实例逻辑set</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SetId *string `json:"SetId,omitnil,omitempty" name:"SetId"`

	// <p>应用id</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>账号id</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// <p>关联的安全组ID列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityGroups []*string `json:"SecurityGroups,omitnil,omitempty" name:"SecurityGroups"`

	// <p>预付费实例降配次数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownGradeTimes *int64 `json:"DownGradeTimes,omitnil,omitempty" name:"DownGradeTimes"`

	// <p>桌面绑定的域名策略ID列表，最多五个</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainControlPolicies []*string `json:"DomainControlPolicies,omitnil,omitempty" name:"DomainControlPolicies"`

	// <p>桌面绑定的水印策略</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WaterMarkPolicy *string `json:"WaterMarkPolicy,omitnil,omitempty" name:"WaterMarkPolicy"`

	// <p>桌面绑定的录屏策略ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScreenRecordingPolicy *string `json:"ScreenRecordingPolicy,omitnil,omitempty" name:"ScreenRecordingPolicy"`

	// <p>WHITE --- 白名单策略<br>BLACK --- 黑名单策略</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainControlPolicyType *string `json:"DomainControlPolicyType,omitnil,omitempty" name:"DomainControlPolicyType"`

	// <p>桌面绑定的运行管控策略</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunControlPolicy *string `json:"RunControlPolicy,omitnil,omitempty" name:"RunControlPolicy"`

	// <p>桌面绑定的安装管控策略</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstallControlPolicy *string `json:"InstallControlPolicy,omitnil,omitempty" name:"InstallControlPolicy"`

	// <p>桌面是否绑定匿名用户</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnonymousAccess *bool `json:"AnonymousAccess,omitnil,omitempty" name:"AnonymousAccess"`

	// <p>实例绑定的标签</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type InstanceStatus struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例状态
	// PENDING：表示创建中
	// LAUNCH_FAILED：表示创建失败
	// RUNNING：表示运行中
	// STOPPED：表示关机
	// STARTING：表示开机中
	// STOPPING：表示关机中
	// REBOOTING：表示重启中
	// SHUTDOWN：表示停止待销毁
	// TERMINATING：表示销毁中
	// DELETED：表示资源已经销毁，这个状态是业务层面的状态，cvm如果销毁了，在cvm测资源已经不存在了
	// FAILED：表示资源创建失败.这个是业务层面的状态
	// ROLLBACKING：表示机器正在回滚中
	// RESETTING：表示机器正在重置中
	// USING：表示机器正在使用中
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 只有在Status为RUNNING时候，当前值才有意义
	// 取值范围：
	// 0:Unknown
	// 1:Connected 
	// 2:Disconnected 
	// 3:Terminated
	// 4:PreparingSession
	// 5:Active
	// 6:Reconnecting
	// 7:NonBrokeredSession
	// 8:Other
	// 9:Pending
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionState *int64 `json:"SessionState,omitnil,omitempty" name:"SessionState"`

	// 此字段只有在Status为USING时，才有效。
	// 字段意义：表示桌面当前session空闲时间（客户没有鼠标，键盘等操作），单位秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdleSeconds *float64 `json:"IdleSeconds,omitnil,omitempty" name:"IdleSeconds"`

	// 当前状态改变的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusChangeTime *int64 `json:"StatusChangeTime,omitnil,omitempty" name:"StatusChangeTime"`

	// 实例最近一次操作类型
	LatestOperation *string `json:"LatestOperation,omitnil,omitempty" name:"LatestOperation"`

	// 实例最近一次操作状态
	LatestOperationState *string `json:"LatestOperationState,omitnil,omitempty" name:"LatestOperationState"`
}

type Invocation struct {
	// <p>脚本执行ID。</p>
	InvokeId *string `json:"InvokeId,omitnil,omitempty" name:"InvokeId"`

	// <p>运行状态。<li> Running：正在执行</li> <li>Finished：已经完成</li> <li>Failed：失败</li>  <li> Stopped：已停止</li></p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>任务输出。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *InvocationOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// <p>任务失败时的错误码。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorCode *string `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// <p>任务失败时的错误信息描述。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *string `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// <p>任务退出状态码</p>
	ExitCode *int64 `json:"ExitCode,omitnil,omitempty" name:"ExitCode"`
}

type InvocationDetail struct {
	// 实例 id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例所在区域，目前支持的区域有：ap-beijing、ap-shanghai、ap-guangzhou、ap-shanghai-adc
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 脚本任务执行状态，取值有：
	// Running-执行中
	// Finished-完成
	// Failed-失败
	// Stopped-终止
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 实例上任务唯一 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvokeId *string `json:"InvokeId,omitnil,omitempty" name:"InvokeId"`

	// 脚本任务输出
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *string `json:"Output,omitnil,omitempty" name:"Output"`

	// 任务退出码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExitCode *int64 `json:"ExitCode,omitnil,omitempty" name:"ExitCode"`

	// 任务错误码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorCode *string `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// 执行任务错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *string `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`
}

type InvocationOutput struct {
	// <p>任务的标准输出。</p>
	Lines []*string `json:"Lines,omitnil,omitempty" name:"Lines"`

	// <p>本次返回的输出第一行对应的索引编号。</p>
	LinesOffset *uint64 `json:"LinesOffset,omitnil,omitempty" name:"LinesOffset"`

	// <p>本次返回的输出的行数。</p>
	LinesNum *uint64 `json:"LinesNum,omitnil,omitempty" name:"LinesNum"`

	// <p>输出的编码方式：<br>&lt;li&gt;PlainText 未经过Base64编码。</p>
	Encoding *string `json:"Encoding,omitnil,omitempty" name:"Encoding"`

	// <p>当前有效的输出起始编号，从0开始。输出过多，会冲掉旧的输出。</p>
	AllLinesOffset *uint64 `json:"AllLinesOffset,omitnil,omitempty" name:"AllLinesOffset"`

	// <p>当前有效的标准输出总行数。</p>
	AllLinesNum *uint64 `json:"AllLinesNum,omitnil,omitempty" name:"AllLinesNum"`
}

type IpControlDetail struct {
	// <p>配置的IP。支持单IP和CIDR，暂不支持IP范围</p>
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// <p>IP的描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

// Predefined struct for user
type ModifyDiskConfigurationRequestParams struct {
	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>新系统盘配置</p>
	NewSystemDisk *Disk `json:"NewSystemDisk,omitnil,omitempty" name:"NewSystemDisk"`

	// <p>新数据盘配置</p>
	NewDataDisks []*Disk `json:"NewDataDisks,omitnil,omitempty" name:"NewDataDisks"`
}

type ModifyDiskConfigurationRequest struct {
	*tchttp.BaseRequest
	
	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>新系统盘配置</p>
	NewSystemDisk *Disk `json:"NewSystemDisk,omitnil,omitempty" name:"NewSystemDisk"`

	// <p>新数据盘配置</p>
	NewDataDisks []*Disk `json:"NewDataDisks,omitnil,omitempty" name:"NewDataDisks"`
}

func (r *ModifyDiskConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDiskConfigurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "NewSystemDisk")
	delete(f, "NewDataDisks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDiskConfigurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDiskConfigurationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDiskConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDiskConfigurationResponseParams `json:"Response"`
}

func (r *ModifyDiskConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDiskConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyImageAttrsRequestParams struct {
	// <p>镜像id</p>
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// <p>支持修改的属性列表为：{ImageName}</p>
	Attrs []*AttributeInfo `json:"Attrs,omitnil,omitempty" name:"Attrs"`
}

type ModifyImageAttrsRequest struct {
	*tchttp.BaseRequest
	
	// <p>镜像id</p>
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// <p>支持修改的属性列表为：{ImageName}</p>
	Attrs []*AttributeInfo `json:"Attrs,omitnil,omitempty" name:"Attrs"`
}

func (r *ModifyImageAttrsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyImageAttrsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageId")
	delete(f, "Attrs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyImageAttrsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyImageAttrsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyImageAttrsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyImageAttrsResponseParams `json:"Response"`
}

func (r *ModifyImageAttrsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyImageAttrsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceBasicAttributeRequestParams struct {
	// 云桌面实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 修改云桌面实例的属性信息
	AttributeInfoList []*AttributeInfo `json:"AttributeInfoList,omitnil,omitempty" name:"AttributeInfoList"`
}

type ModifyInstanceBasicAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 云桌面实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 修改云桌面实例的属性信息
	AttributeInfoList []*AttributeInfo `json:"AttributeInfoList,omitnil,omitempty" name:"AttributeInfoList"`
}

func (r *ModifyInstanceBasicAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceBasicAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AttributeInfoList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceBasicAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceBasicAttributeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceBasicAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceBasicAttributeResponseParams `json:"Response"`
}

func (r *ModifyInstanceBasicAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceBasicAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstancesTypeRequestParams struct {
	// <p>实例id数组</p>
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// <p>变更的实例类型</p>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// <p>付费类型</p>
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`
}

type ModifyInstancesTypeRequest struct {
	*tchttp.BaseRequest
	
	// <p>实例id数组</p>
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// <p>变更的实例类型</p>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// <p>付费类型</p>
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`
}

func (r *ModifyInstancesTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancesTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "InstanceType")
	delete(f, "ChargeType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstancesTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstancesTypeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstancesTypeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstancesTypeResponseParams `json:"Response"`
}

func (r *ModifyInstancesTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancesTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstancesVpcRequestParams struct {
	// <p>实例ID数组</p>
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// <p>私有网络ID</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>子网ID</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>租户vpc的dns</p><p>参数格式：“189.0.01”</p>
	DnsServerSet []*string `json:"DnsServerSet,omitnil,omitempty" name:"DnsServerSet"`
}

type ModifyInstancesVpcRequest struct {
	*tchttp.BaseRequest
	
	// <p>实例ID数组</p>
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// <p>私有网络ID</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>子网ID</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>租户vpc的dns</p><p>参数格式：“189.0.01”</p>
	DnsServerSet []*string `json:"DnsServerSet,omitnil,omitempty" name:"DnsServerSet"`
}

func (r *ModifyInstancesVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancesVpcRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "DnsServerSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstancesVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstancesVpcResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstancesVpcResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstancesVpcResponseParams `json:"Response"`
}

func (r *ModifyInstancesVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancesVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Partition struct {
	// 磁盘名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 总容量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *float64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 使用量,39.31%
	// 注意：此字段可能返回 null，表示取不到有效值。
	Usage *float64 `json:"Usage,omitnil,omitempty" name:"Usage"`
}

type PolicyDetail struct {
	// <p>策略id</p>
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// <p>策略名称</p>
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// <p>策略描述</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>策略作用用户id列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// <p>策略的租户uin</p>
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// <p>文件重定向策略id。取值范围：<br>1：读写<br>2：只读<br>3：禁止读写</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilePolicyId *int64 `json:"FilePolicyId,omitnil,omitempty" name:"FilePolicyId"`

	// <p>剪切板重定向策略id。取值范围：<br>1：开启云桌面到客户端<br>2：开启客户端到云桌面<br>3：双向开启<br>4：禁用</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClipPolicyId *int64 `json:"ClipPolicyId,omitnil,omitempty" name:"ClipPolicyId"`

	// <p>外设重定向策略id。取值范围：<br>1：允许usb重定向<br>2：允许端口重定向<br>3：允许usb和端口重定向<br>4：禁用</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DrivePolicyId *int64 `json:"DrivePolicyId,omitnil,omitempty" name:"DrivePolicyId"`

	// <p>水印重定向策略id。取值范围（离散值）：<br>4:   当前时间，居中策略<br>10: 当前时间，平铺策略<br>15:  云桌面ID， 居中策略<br>16:  云桌面ID+当前时间， 居中策略<br>17:  云桌面ID， 平铺策略<br>18:  云桌面ID+当前时间， 平铺策略</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WaterMarkPolicyId *int64 `json:"WaterMarkPolicyId,omitnil,omitempty" name:"WaterMarkPolicyId"`

	// <p>是否是默认策略</p>
	IsDefault *bool `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`

	// <p>注意：默认策略的当前值无效</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalUser *int64 `json:"TotalUser,omitnil,omitempty" name:"TotalUser"`
}

type PolicyInfo struct {
	// 策略类型。取值：
	// 1：文件重定向策略
	// 2：剪切板重定向策略
	// 3：外设重定向策略
	// 4：水印策略
	PolicyType *int64 `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// 策略值。取值：
	// PolicyType对应个子取值范围：
	// 	①：PolicyType为1-文件重定向策略	
	// 					1：读写
	// 					2：只读
	// 					3：禁止读写
	// 	②：剪切板重定向策略
	// 					1：开启云桌面到客户端
	// 					2：开启客户端到云桌面
	// 					3：双向开启
	// 					4：禁用
	// 	③：外设重定向策略
	// 					1：允许usb重定向
	// 					2：允许端口重定向
	// 					3：允许usb和端口重定向
	// 					4：禁用
	// 	④：水印策略
	// 					4:   当前时间，居中策略
	// 					10: 当前时间，平铺策略
	// 					15:  云桌面ID， 居中策略
	//                                         16:  云桌面ID+当前时间， 居中策略
	//                                         17:  云桌面ID， 平铺策略
	//                                         18:  云桌面ID+当前时间， 平铺策略
	PolicyValue *int64 `json:"PolicyValue,omitnil,omitempty" name:"PolicyValue"`
}

// Predefined struct for user
type RebootInstancesRequestParams struct {
	// 重启实例id列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 实例所在区域，保证传入的列表在统一区域，否则报错
	InstanceRegion *string `json:"InstanceRegion,omitnil,omitempty" name:"InstanceRegion"`

	// 是否在正常重启失败后选择强制重启实例。取值范围：
	// TRUE：表示在正常重启失败后进行强制重启
	// FALSE：表示在正常重启失败后不进行强制重启
	// 
	// 默认取值：FALSE。
	Force *bool `json:"Force,omitnil,omitempty" name:"Force"`

	// 重启是否需要预登录
	PreLogin *bool `json:"PreLogin,omitnil,omitempty" name:"PreLogin"`
}

type RebootInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 重启实例id列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 实例所在区域，保证传入的列表在统一区域，否则报错
	InstanceRegion *string `json:"InstanceRegion,omitnil,omitempty" name:"InstanceRegion"`

	// 是否在正常重启失败后选择强制重启实例。取值范围：
	// TRUE：表示在正常重启失败后进行强制重启
	// FALSE：表示在正常重启失败后不进行强制重启
	// 
	// 默认取值：FALSE。
	Force *bool `json:"Force,omitnil,omitempty" name:"Force"`

	// 重启是否需要预登录
	PreLogin *bool `json:"PreLogin,omitnil,omitempty" name:"PreLogin"`
}

func (r *RebootInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RebootInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "InstanceRegion")
	delete(f, "Force")
	delete(f, "PreLogin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RebootInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RebootInstancesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RebootInstancesResponse struct {
	*tchttp.BaseResponse
	Response *RebootInstancesResponseParams `json:"Response"`
}

func (r *RebootInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RebootInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveDataDisksRequestParams struct {
	// 实例 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据盘磁盘 ID，可以通过 DescribeInstances 查下实例数据盘
	DataDiskIds []*string `json:"DataDiskIds,omitnil,omitempty" name:"DataDiskIds"`
}

type RemoveDataDisksRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据盘磁盘 ID，可以通过 DescribeInstances 查下实例数据盘
	DataDiskIds []*string `json:"DataDiskIds,omitnil,omitempty" name:"DataDiskIds"`
}

func (r *RemoveDataDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveDataDisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DataDiskIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveDataDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveDataDisksResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemoveDataDisksResponse struct {
	*tchttp.BaseResponse
	Response *RemoveDataDisksResponseParams `json:"Response"`
}

func (r *RemoveDataDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveDataDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemovePolicyRequestParams struct {
	// <p>待删除策略id</p>
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`
}

type RemovePolicyRequest struct {
	*tchttp.BaseRequest
	
	// <p>待删除策略id</p>
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`
}

func (r *RemovePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemovePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemovePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemovePolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemovePolicyResponse struct {
	*tchttp.BaseResponse
	Response *RemovePolicyResponseParams `json:"Response"`
}

func (r *RemovePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemovePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewInstancesRequestParams struct {
	// 一个或多个待操作的实例ID。每次请求批量实例的上限为50。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的续费时长、是否设置自动续费等属性。包年包月实例该参数为必传参数。
	ChargePrepaid *ChargePrepaid `json:"ChargePrepaid,omitnil,omitempty" name:"ChargePrepaid"`
}

type RenewInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 一个或多个待操作的实例ID。每次请求批量实例的上限为50。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的续费时长、是否设置自动续费等属性。包年包月实例该参数为必传参数。
	ChargePrepaid *ChargePrepaid `json:"ChargePrepaid,omitnil,omitempty" name:"ChargePrepaid"`
}

func (r *RenewInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "ChargePrepaid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewInstancesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RenewInstancesResponse struct {
	*tchttp.BaseResponse
	Response *RenewInstancesResponseParams `json:"Response"`
}

func (r *RenewInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetInstancesRequestParams struct {
	// 需要重置的实例id列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 重置的目标镜像id
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 重置实例的区域
	InstanceRegion *string `json:"InstanceRegion,omitnil,omitempty" name:"InstanceRegion"`

	// 是否解绑机器原有绑定关系
	IsUnbind *bool `json:"IsUnbind,omitnil,omitempty" name:"IsUnbind"`
}

type ResetInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 需要重置的实例id列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 重置的目标镜像id
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 重置实例的区域
	InstanceRegion *string `json:"InstanceRegion,omitnil,omitempty" name:"InstanceRegion"`

	// 是否解绑机器原有绑定关系
	IsUnbind *bool `json:"IsUnbind,omitnil,omitempty" name:"IsUnbind"`
}

func (r *ResetInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "ImageId")
	delete(f, "InstanceRegion")
	delete(f, "IsUnbind")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetInstancesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResetInstancesResponse struct {
	*tchttp.BaseResponse
	Response *ResetInstancesResponseParams `json:"Response"`
}

func (r *ResetInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunCommandRequestParams struct {
	// <p>云桌面实例ID。</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>执行的脚本内容，编码方式由ContentEncoding指定。<br>&lt;li&gt;脚本内容在Base64编码后的大小不能超过16 KB。</p>
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// <p>脚本种类。取值有：&lt;li&gt;PowerShellScript Powershell脚本 &lt;li&gt;BatScript Bat脚本</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>是否同步等待脚本的执行结果。&lt;li&gt;false为异步模式，通过此接口触发任务的执行，之后异步获取执行结果。默认为false。此为推荐的模式。&lt;li&gt;true为同步等待执行结果，只有当脚本执行时间在4秒之内时，才适合将该参数置为true。如果接口返回时任务尚未执行完毕的话，也需要通过异步接口来获取执行结果。</p>
	WaitUntilDone *bool `json:"WaitUntilDone,omitnil,omitempty" name:"WaitUntilDone"`

	// <p>Content的编码。&lt;li&gt;Base64 Base64编码。 默认为Base64</p>
	ContentEncoding *string `json:"ContentEncoding,omitnil,omitempty" name:"ContentEncoding"`

	// <p>任务的输出内容是否合并成一个字符串。&lt;li&gt;true表示所有的输出行拼接在一起。默认为true。&lt;li&gt;false表示输出内容通过数组返回,数组的每一个元素为一行。</p>
	MergeOutputLines *bool `json:"MergeOutputLines,omitnil,omitempty" name:"MergeOutputLines"`

	// <p>执行任务的用户ID。&lt;li&gt;由SYSTEM身份执行。此场景下不需要指定该字段，这是默认情况。&lt;li&gt;需要以绑定到实例的用户执行任务，通过该字段指定用户ID。 &lt;li&gt;匿名访问，需要以实际登录系统的用户执行的场景，该字段填实例ID。</p>
	RunAs *string `json:"RunAs,omitnil,omitempty" name:"RunAs"`

	// <p>任务是否为前台运行。仅在RunAs场景下有效。仅当任务会呼起图形界面时有意义。&lt;li&gt;true表示用户可以看到图形窗口。&lt;li&gt;false表示图形窗口不可见。默认为false。</p>
	RunForeground *bool `json:"RunForeground,omitnil,omitempty" name:"RunForeground"`
}

type RunCommandRequest struct {
	*tchttp.BaseRequest
	
	// <p>云桌面实例ID。</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>执行的脚本内容，编码方式由ContentEncoding指定。<br>&lt;li&gt;脚本内容在Base64编码后的大小不能超过16 KB。</p>
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// <p>脚本种类。取值有：&lt;li&gt;PowerShellScript Powershell脚本 &lt;li&gt;BatScript Bat脚本</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>是否同步等待脚本的执行结果。&lt;li&gt;false为异步模式，通过此接口触发任务的执行，之后异步获取执行结果。默认为false。此为推荐的模式。&lt;li&gt;true为同步等待执行结果，只有当脚本执行时间在4秒之内时，才适合将该参数置为true。如果接口返回时任务尚未执行完毕的话，也需要通过异步接口来获取执行结果。</p>
	WaitUntilDone *bool `json:"WaitUntilDone,omitnil,omitempty" name:"WaitUntilDone"`

	// <p>Content的编码。&lt;li&gt;Base64 Base64编码。 默认为Base64</p>
	ContentEncoding *string `json:"ContentEncoding,omitnil,omitempty" name:"ContentEncoding"`

	// <p>任务的输出内容是否合并成一个字符串。&lt;li&gt;true表示所有的输出行拼接在一起。默认为true。&lt;li&gt;false表示输出内容通过数组返回,数组的每一个元素为一行。</p>
	MergeOutputLines *bool `json:"MergeOutputLines,omitnil,omitempty" name:"MergeOutputLines"`

	// <p>执行任务的用户ID。&lt;li&gt;由SYSTEM身份执行。此场景下不需要指定该字段，这是默认情况。&lt;li&gt;需要以绑定到实例的用户执行任务，通过该字段指定用户ID。 &lt;li&gt;匿名访问，需要以实际登录系统的用户执行的场景，该字段填实例ID。</p>
	RunAs *string `json:"RunAs,omitnil,omitempty" name:"RunAs"`

	// <p>任务是否为前台运行。仅在RunAs场景下有效。仅当任务会呼起图形界面时有意义。&lt;li&gt;true表示用户可以看到图形窗口。&lt;li&gt;false表示图形窗口不可见。默认为false。</p>
	RunForeground *bool `json:"RunForeground,omitnil,omitempty" name:"RunForeground"`
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
	delete(f, "InstanceId")
	delete(f, "Content")
	delete(f, "Type")
	delete(f, "WaitUntilDone")
	delete(f, "ContentEncoding")
	delete(f, "MergeOutputLines")
	delete(f, "RunAs")
	delete(f, "RunForeground")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunCommandRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunCommandResponseParams struct {
	// <p>本次任务对应的任务ID。方便异步查询执行状态。</p>
	InvokeId *string `json:"InvokeId,omitnil,omitempty" name:"InvokeId"`

	// <p>执行详细情况。WaitUntilDone为true，在接口返回时的执行状态。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Detail *Invocation `json:"Detail,omitnil,omitempty" name:"Detail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type SnapshotInfo struct {
	// <p>快照唯一ID</p>
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// <p>快照对应的云硬盘ID</p>
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`

	// <p>快照状态，取值如下：<br>NORMAL：正常<br>CREATING：创建中<br>ROLLBACKING：回滚中<br>COPYING_FROM_REMOTE：跨地域复制中<br>CHECKING_COPIED：复制校验中<br>TORECYCLE：待回收。</p>
	SnapshotState *string `json:"SnapshotState,omitnil,omitempty" name:"SnapshotState"`

	// <p>创建此快照的云硬盘类型。取值范围：<br>SYSTEM_DISK：系统盘<br>DATA_DISK：数据盘。</p>
	DiskUsage *string `json:"DiskUsage,omitnil,omitempty" name:"DiskUsage"`

	// <p>快照创建的时间</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>快照到期时间。如果快照为永久保留，此字段为空。</p>
	DeadlineTime *string `json:"DeadlineTime,omitnil,omitempty" name:"DeadlineTime"`

	// <p>快照创建进度百分比，快照创建成功后此字段恒为100。</p>
	Percent *uint64 `json:"Percent,omitnil,omitempty" name:"Percent"`

	// <p>是否为永久快照。取值范围：<br>true：永久快照<br>false：非永久快照。</p>
	IsPermanent *bool `json:"IsPermanent,omitnil,omitempty" name:"IsPermanent"`

	// <p>快照创建方式，取值:<br>ManualSnapshot：手动创建<br>AutoSnapshot：快照策略自动创建</p>
	CreateType *string `json:"CreateType,omitnil,omitempty" name:"CreateType"`

	// <p>上一次快照回滚完成时间</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastRollbackTime *string `json:"LastRollbackTime,omitnil,omitempty" name:"LastRollbackTime"`

	// <p>上一次快照回滚状态，取值：<br>成功：SUCCESS<br>失败：FAILED</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastRollbackState *string `json:"LastRollbackState,omitnil,omitempty" name:"LastRollbackState"`
}

// Predefined struct for user
type StartInstancesRequestParams struct {
	// 需要开启的实例id列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 实例所在区域，保证传入的列表再统一区域，否则报错
	InstanceRegion *string `json:"InstanceRegion,omitnil,omitempty" name:"InstanceRegion"`
}

type StartInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 需要开启的实例id列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 实例所在区域，保证传入的列表再统一区域，否则报错
	InstanceRegion *string `json:"InstanceRegion,omitnil,omitempty" name:"InstanceRegion"`
}

func (r *StartInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "InstanceRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartInstancesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartInstancesResponse struct {
	*tchttp.BaseResponse
	Response *StartInstancesResponseParams `json:"Response"`
}

func (r *StartInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopCommandTasksRequestParams struct {
	// <p>删除命令任务条件<br>Filter.name 目前支持 CommandId</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type StopCommandTasksRequest struct {
	*tchttp.BaseRequest
	
	// <p>删除命令任务条件<br>Filter.name 目前支持 CommandId</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *StopCommandTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopCommandTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopCommandTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopCommandTasksResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopCommandTasksResponse struct {
	*tchttp.BaseResponse
	Response *StopCommandTasksResponseParams `json:"Response"`
}

func (r *StopCommandTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopCommandTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopInstancesRequestParams struct {
	// 需要重启的实例id列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 实例所在区域，保证传入的列表在统一区域，否则报错
	InstanceRegion *string `json:"InstanceRegion,omitnil,omitempty" name:"InstanceRegion"`

	// 是否在正常重启失败后选择强制重启实例。取值范围：
	// TRUE：表示在正常重启失败后进行强制重启
	// FALSE：表示在正常重启失败后不进行强制重启
	// 
	// 默认取值：FALSE。
	Force *bool `json:"Force,omitnil,omitempty" name:"Force"`
}

type StopInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 需要重启的实例id列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 实例所在区域，保证传入的列表在统一区域，否则报错
	InstanceRegion *string `json:"InstanceRegion,omitnil,omitempty" name:"InstanceRegion"`

	// 是否在正常重启失败后选择强制重启实例。取值范围：
	// TRUE：表示在正常重启失败后进行强制重启
	// FALSE：表示在正常重启失败后不进行强制重启
	// 
	// 默认取值：FALSE。
	Force *bool `json:"Force,omitnil,omitempty" name:"Force"`
}

func (r *StopInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "InstanceRegion")
	delete(f, "Force")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopInstancesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopInstancesResponse struct {
	*tchttp.BaseResponse
	Response *StopInstancesResponseParams `json:"Response"`
}

func (r *StopInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopInvocationRequestParams struct {
	// <p>云桌面实例ID。</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>任务ID。</p>
	InvokeId *string `json:"InvokeId,omitnil,omitempty" name:"InvokeId"`
}

type StopInvocationRequest struct {
	*tchttp.BaseRequest
	
	// <p>云桌面实例ID。</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>任务ID。</p>
	InvokeId *string `json:"InvokeId,omitnil,omitempty" name:"InvokeId"`
}

func (r *StopInvocationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopInvocationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InvokeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopInvocationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopInvocationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopInvocationResponse struct {
	*tchttp.BaseResponse
	Response *StopInvocationResponseParams `json:"Response"`
}

func (r *StopInvocationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopInvocationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopSessionRequestParams struct {
	// <p>机器所在区域</p>
	InstanceRegion *string `json:"InstanceRegion,omitnil,omitempty" name:"InstanceRegion"`

	// <p>机器hostname</p>
	HostName *string `json:"HostName,omitnil,omitempty" name:"HostName"`

	// <p>IsDisconnect为true，表示断开当前会话<br>IsDisconnect为false，表示注销当前会话</p>
	IsDisconnect *bool `json:"IsDisconnect,omitnil,omitempty" name:"IsDisconnect"`

	// <p>选填，为多会话预留字段</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type StopSessionRequest struct {
	*tchttp.BaseRequest
	
	// <p>机器所在区域</p>
	InstanceRegion *string `json:"InstanceRegion,omitnil,omitempty" name:"InstanceRegion"`

	// <p>机器hostname</p>
	HostName *string `json:"HostName,omitnil,omitempty" name:"HostName"`

	// <p>IsDisconnect为true，表示断开当前会话<br>IsDisconnect为false，表示注销当前会话</p>
	IsDisconnect *bool `json:"IsDisconnect,omitnil,omitempty" name:"IsDisconnect"`

	// <p>选填，为多会话预留字段</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *StopSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceRegion")
	delete(f, "HostName")
	delete(f, "IsDisconnect")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopSessionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopSessionResponse struct {
	*tchttp.BaseResponse
	Response *StopSessionResponseParams `json:"Response"`
}

func (r *StopSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {
	// 标签键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

// Predefined struct for user
type TerminateInstancesRequestParams struct {
	// <p>一个或多个待操作的实例ID。每次请求批量实例的上限为10。</p>
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// <p>是否立即销毁实例。默认false。 立即销毁的实例不会进入隔离期。</p>
	DestroyImmediately *bool `json:"DestroyImmediately,omitnil,omitempty" name:"DestroyImmediately"`
}

type TerminateInstancesRequest struct {
	*tchttp.BaseRequest
	
	// <p>一个或多个待操作的实例ID。每次请求批量实例的上限为10。</p>
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// <p>是否立即销毁实例。默认false。 立即销毁的实例不会进入隔离期。</p>
	DestroyImmediately *bool `json:"DestroyImmediately,omitnil,omitempty" name:"DestroyImmediately"`
}

func (r *TerminateInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "DestroyImmediately")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateInstancesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TerminateInstancesResponse struct {
	*tchttp.BaseResponse
	Response *TerminateInstancesResponseParams `json:"Response"`
}

func (r *TerminateInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindInstancesRequestParams struct {
	// 实例所在地域
	InstanceRegion *string `json:"InstanceRegion,omitnil,omitempty" name:"InstanceRegion"`

	// 解除绑定关系列表
	BindInfos []*BindInfo `json:"BindInfos,omitnil,omitempty" name:"BindInfos"`
}

type UnbindInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例所在地域
	InstanceRegion *string `json:"InstanceRegion,omitnil,omitempty" name:"InstanceRegion"`

	// 解除绑定关系列表
	BindInfos []*BindInfo `json:"BindInfos,omitnil,omitempty" name:"BindInfos"`
}

func (r *UnbindInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceRegion")
	delete(f, "BindInfos")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindInstancesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UnbindInstancesResponse struct {
	*tchttp.BaseResponse
	Response *UnbindInstancesResponseParams `json:"Response"`
}

func (r *UnbindInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdatePolicyRequestParams struct {
	// <p>待更新策略id</p>
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// <p>更新策略后策略名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>更新策略后策略描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>更新策略的具体修改值（全集）</p>
	PolicyInfoList []*PolicyInfo `json:"PolicyInfoList,omitnil,omitempty" name:"PolicyInfoList"`

	// <p>新增绑定用户列表</p>
	BindUserIds []*string `json:"BindUserIds,omitnil,omitempty" name:"BindUserIds"`

	// <p>移除绑定的用户列表</p>
	RemoveUserIds []*string `json:"RemoveUserIds,omitnil,omitempty" name:"RemoveUserIds"`
}

type UpdatePolicyRequest struct {
	*tchttp.BaseRequest
	
	// <p>待更新策略id</p>
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// <p>更新策略后策略名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>更新策略后策略描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>更新策略的具体修改值（全集）</p>
	PolicyInfoList []*PolicyInfo `json:"PolicyInfoList,omitnil,omitempty" name:"PolicyInfoList"`

	// <p>新增绑定用户列表</p>
	BindUserIds []*string `json:"BindUserIds,omitnil,omitempty" name:"BindUserIds"`

	// <p>移除绑定的用户列表</p>
	RemoveUserIds []*string `json:"RemoveUserIds,omitnil,omitempty" name:"RemoveUserIds"`
}

func (r *UpdatePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdatePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "PolicyInfoList")
	delete(f, "BindUserIds")
	delete(f, "RemoveUserIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdatePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdatePolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdatePolicyResponse struct {
	*tchttp.BaseResponse
	Response *UpdatePolicyResponseParams `json:"Response"`
}

func (r *UpdatePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdatePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Usage struct {
	// 桌面实例 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 桌面名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 桌面当前状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 桌面绑定的用户名，若未绑定，为空
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 桌面计费类型
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 桌面类型
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 桌面操作系统
	OsName *string `json:"OsName,omitnil,omitempty" name:"OsName"`

	// 桌面使用时长，h:m:s
	ActiveTime *string `json:"ActiveTime,omitnil,omitempty" name:"ActiveTime"`

	// 桌面空闲时长，h:m:s
	IdleTime *string `json:"IdleTime,omitnil,omitempty" name:"IdleTime"`

	// 是否匿名访问
	AnonymousAccess *bool `json:"AnonymousAccess,omitnil,omitempty" name:"AnonymousAccess"`
}

type User struct {
	// <p>用户ID。</p>
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>名字。</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>显示名字。</p>
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// <p>邮箱。</p>
	Mail *string `json:"Mail,omitnil,omitempty" name:"Mail"`

	// <p>手机号。</p>
	PhoneNum *string `json:"PhoneNum,omitnil,omitempty" name:"PhoneNum"`

	// <p>状态。<br>ACTIVE:启用<br>PENDING_ACTIVATE:待激活<br>PENDING_ONBOARD:待启用<br>SUSPENDED:已停用<br>INACTIVE:已离职<br>PASSWORD_RESET:重置密码<br>LOCKED_OUT:锁定</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>部门ID。</p>
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`

	// <p>策略ID。</p>
	StrategyId *string `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`

	// <p>策略名称。</p>
	StrategyName *string `json:"StrategyName,omitnil,omitempty" name:"StrategyName"`

	// <p>当前用户是否绑定桌面，true：绑定，false：未绑定任何桌面</p>
	BindStatus *bool `json:"BindStatus,omitnil,omitempty" name:"BindStatus"`
}

type VirtualPrivateCloud struct {
	// <p>私有网络ID。</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>私有网络子网ID。</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`
}