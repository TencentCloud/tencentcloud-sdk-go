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

package v20230418

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AvailableModelVersion struct {
	// 带有版本号的设备型号
	ModelVersion *string `json:"ModelVersion,omitnil,omitempty" name:"ModelVersion"`

	// 设备高度
	DevHeight *string `json:"DevHeight,omitnil,omitempty" name:"DevHeight"`

	// 设备类型，server 服务器，netDevice 网络设备
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`
}

type Cage struct {
	// 围笼名称
	CageName *string `json:"CageName,omitnil,omitempty" name:"CageName"`

	// 围笼审核人账号ID
	CheckerSet []*string `json:"CheckerSet,omitnil,omitempty" name:"CheckerSet"`
}

type Campus struct {
	// 园区ID
	CampusId *uint64 `json:"CampusId,omitnil,omitempty" name:"CampusId"`

	// 园区名称
	CampusName *string `json:"CampusName,omitnil,omitempty" name:"CampusName"`
}

type CommonServiceBaseInfo struct {
	// 机房名称
	IdcName *string `json:"IdcName,omitnil,omitempty" name:"IdcName"`

	// 业务联系人
	ContactName *string `json:"ContactName,omitnil,omitempty" name:"ContactName"`

	// 联系人电话
	ContactPhone *string `json:"ContactPhone,omitnil,omitempty" name:"ContactPhone"`

	// 操作说明
	Instructions *string `json:"Instructions,omitnil,omitempty" name:"Instructions"`

	// 1 、2 、3 分别代表 L1、L2、L3
	ServiceLevel *uint64 `json:"ServiceLevel,omitnil,omitempty" name:"ServiceLevel"`

	// 操作预授权
	PreAuthorization *bool `json:"PreAuthorization,omitnil,omitempty" name:"PreAuthorization"`
}

// Predefined struct for user
type ConfirmCommonServiceWorkOrderRequestParams struct {
	// 工单ID
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`
}

type ConfirmCommonServiceWorkOrderRequest struct {
	*tchttp.BaseRequest
	
	// 工单ID
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`
}

func (r *ConfirmCommonServiceWorkOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ConfirmCommonServiceWorkOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrderId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ConfirmCommonServiceWorkOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ConfirmCommonServiceWorkOrderResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ConfirmCommonServiceWorkOrderResponse struct {
	*tchttp.BaseResponse
	Response *ConfirmCommonServiceWorkOrderResponseParams `json:"Response"`
}

func (r *ConfirmCommonServiceWorkOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ConfirmCommonServiceWorkOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCommonServiceWorkOrderRequestParams struct {
	// 设备及位置信息列表
	DevicePositionSet []*DevicePosition `json:"DevicePositionSet,omitnil,omitempty" name:"DevicePositionSet"`

	// 服务级别，只支持传入 1、2、3，分别对应 L1、L2、L3
	ServiceLevel *uint64 `json:"ServiceLevel,omitnil,omitempty" name:"ServiceLevel"`

	// 操作预授权
	PreAuthorization *bool `json:"PreAuthorization,omitnil,omitempty" name:"PreAuthorization"`

	// 业务联系人
	ContactName *string `json:"ContactName,omitnil,omitempty" name:"ContactName"`

	// 联系人电话
	ContactPhone *string `json:"ContactPhone,omitnil,omitempty" name:"ContactPhone"`

	// 设备类型 server 服务器，netDevice 网络设备	
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 操作说明
	Instructions *string `json:"Instructions,omitnil,omitempty" name:"Instructions"`
}

type CreateCommonServiceWorkOrderRequest struct {
	*tchttp.BaseRequest
	
	// 设备及位置信息列表
	DevicePositionSet []*DevicePosition `json:"DevicePositionSet,omitnil,omitempty" name:"DevicePositionSet"`

	// 服务级别，只支持传入 1、2、3，分别对应 L1、L2、L3
	ServiceLevel *uint64 `json:"ServiceLevel,omitnil,omitempty" name:"ServiceLevel"`

	// 操作预授权
	PreAuthorization *bool `json:"PreAuthorization,omitnil,omitempty" name:"PreAuthorization"`

	// 业务联系人
	ContactName *string `json:"ContactName,omitnil,omitempty" name:"ContactName"`

	// 联系人电话
	ContactPhone *string `json:"ContactPhone,omitnil,omitempty" name:"ContactPhone"`

	// 设备类型 server 服务器，netDevice 网络设备	
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 操作说明
	Instructions *string `json:"Instructions,omitnil,omitempty" name:"Instructions"`
}

func (r *CreateCommonServiceWorkOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCommonServiceWorkOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DevicePositionSet")
	delete(f, "ServiceLevel")
	delete(f, "PreAuthorization")
	delete(f, "ContactName")
	delete(f, "ContactPhone")
	delete(f, "DeviceType")
	delete(f, "Instructions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCommonServiceWorkOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCommonServiceWorkOrderResponseParams struct {
	// 创建的工单信息
	WorkOrderSet []*WorkOrderTinyInfo `json:"WorkOrderSet,omitnil,omitempty" name:"WorkOrderSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCommonServiceWorkOrderResponse struct {
	*tchttp.BaseResponse
	Response *CreateCommonServiceWorkOrderResponseParams `json:"Response"`
}

func (r *CreateCommonServiceWorkOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCommonServiceWorkOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateModelEvaluationWorkOrderRequestParams struct {
	// 设备名称及型号
	ModelSet []*ModelVersion `json:"ModelSet,omitnil,omitempty" name:"ModelSet"`

	// 园区ID
	CampusId *uint64 `json:"CampusId,omitnil,omitempty" name:"CampusId"`

	// 只支持传入 server 和 netDevice
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type CreateModelEvaluationWorkOrderRequest struct {
	*tchttp.BaseRequest
	
	// 设备名称及型号
	ModelSet []*ModelVersion `json:"ModelSet,omitnil,omitempty" name:"ModelSet"`

	// 园区ID
	CampusId *uint64 `json:"CampusId,omitnil,omitempty" name:"CampusId"`

	// 只支持传入 server 和 netDevice
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *CreateModelEvaluationWorkOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateModelEvaluationWorkOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelSet")
	delete(f, "CampusId")
	delete(f, "DeviceType")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateModelEvaluationWorkOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateModelEvaluationWorkOrderResponseParams struct {
	// 创建的工单信息
	WorkOrderSet []*WorkOrderTinyInfo `json:"WorkOrderSet,omitnil,omitempty" name:"WorkOrderSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateModelEvaluationWorkOrderResponse struct {
	*tchttp.BaseResponse
	Response *CreateModelEvaluationWorkOrderResponseParams `json:"Response"`
}

func (r *CreateModelEvaluationWorkOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateModelEvaluationWorkOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMovingWorkOrderRequestParams struct {
	// 机房id
	IdcId *uint64 `json:"IdcId,omitnil,omitempty" name:"IdcId"`

	// 设备类型，server, netDevice
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 上架后是否开电
	WithPowerOn *bool `json:"WithPowerOn,omitnil,omitempty" name:"WithPowerOn"`

	// 设备搬迁信息列表
	DeviceMovingList []*DeviceRackOn `json:"DeviceMovingList,omitnil,omitempty" name:"DeviceMovingList"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type CreateMovingWorkOrderRequest struct {
	*tchttp.BaseRequest
	
	// 机房id
	IdcId *uint64 `json:"IdcId,omitnil,omitempty" name:"IdcId"`

	// 设备类型，server, netDevice
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 上架后是否开电
	WithPowerOn *bool `json:"WithPowerOn,omitnil,omitempty" name:"WithPowerOn"`

	// 设备搬迁信息列表
	DeviceMovingList []*DeviceRackOn `json:"DeviceMovingList,omitnil,omitempty" name:"DeviceMovingList"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *CreateMovingWorkOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMovingWorkOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdcId")
	delete(f, "DeviceType")
	delete(f, "WithPowerOn")
	delete(f, "DeviceMovingList")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMovingWorkOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMovingWorkOrderResponseParams struct {
	// 创建的工单信息
	WorkOrderSet []*WorkOrderTinyInfo `json:"WorkOrderSet,omitnil,omitempty" name:"WorkOrderSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateMovingWorkOrderResponse struct {
	*tchttp.BaseResponse
	Response *CreateMovingWorkOrderResponseParams `json:"Response"`
}

func (r *CreateMovingWorkOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMovingWorkOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNetDeviceModelRequestParams struct {
	// 网络设备型号
	ModelDetail *NetDeviceModel `json:"ModelDetail,omitnil,omitempty" name:"ModelDetail"`
}

type CreateNetDeviceModelRequest struct {
	*tchttp.BaseRequest
	
	// 网络设备型号
	ModelDetail *NetDeviceModel `json:"ModelDetail,omitnil,omitempty" name:"ModelDetail"`
}

func (r *CreateNetDeviceModelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNetDeviceModelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelDetail")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNetDeviceModelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNetDeviceModelResponseParams struct {
	// 型号名称
	DevModel *string `json:"DevModel,omitnil,omitempty" name:"DevModel"`

	// 版本号
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateNetDeviceModelResponse struct {
	*tchttp.BaseResponse
	Response *CreateNetDeviceModelResponseParams `json:"Response"`
}

func (r *CreateNetDeviceModelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNetDeviceModelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePersonnelVisitWorkOrderRequestParams struct {
	// 到访人员信息
	PersonnelSet []*Personnel `json:"PersonnelSet,omitnil,omitempty" name:"PersonnelSet"`

	// 机房 ID
	IdcId *uint64 `json:"IdcId,omitnil,omitempty" name:"IdcId"`

	// 机房管理单元列表
	IdcUnitIdSet []*uint64 `json:"IdcUnitIdSet,omitnil,omitempty" name:"IdcUnitIdSet"`

	// 到访开始时间
	EnterStartTime *string `json:"EnterStartTime,omitnil,omitempty" name:"EnterStartTime"`

	// 到访结束时间
	EnterEndTime *string `json:"EnterEndTime,omitnil,omitempty" name:"EnterEndTime"`

	// 到访原因，映射关系：IT_OPERATION IT运维 OTHER 其他
	VisitReason []*string `json:"VisitReason,omitnil,omitempty" name:"VisitReason"`

	// 到访说明
	VisitRemark *string `json:"VisitRemark,omitnil,omitempty" name:"VisitRemark"`
}

type CreatePersonnelVisitWorkOrderRequest struct {
	*tchttp.BaseRequest
	
	// 到访人员信息
	PersonnelSet []*Personnel `json:"PersonnelSet,omitnil,omitempty" name:"PersonnelSet"`

	// 机房 ID
	IdcId *uint64 `json:"IdcId,omitnil,omitempty" name:"IdcId"`

	// 机房管理单元列表
	IdcUnitIdSet []*uint64 `json:"IdcUnitIdSet,omitnil,omitempty" name:"IdcUnitIdSet"`

	// 到访开始时间
	EnterStartTime *string `json:"EnterStartTime,omitnil,omitempty" name:"EnterStartTime"`

	// 到访结束时间
	EnterEndTime *string `json:"EnterEndTime,omitnil,omitempty" name:"EnterEndTime"`

	// 到访原因，映射关系：IT_OPERATION IT运维 OTHER 其他
	VisitReason []*string `json:"VisitReason,omitnil,omitempty" name:"VisitReason"`

	// 到访说明
	VisitRemark *string `json:"VisitRemark,omitnil,omitempty" name:"VisitRemark"`
}

func (r *CreatePersonnelVisitWorkOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePersonnelVisitWorkOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PersonnelSet")
	delete(f, "IdcId")
	delete(f, "IdcUnitIdSet")
	delete(f, "EnterStartTime")
	delete(f, "EnterEndTime")
	delete(f, "VisitReason")
	delete(f, "VisitRemark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePersonnelVisitWorkOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePersonnelVisitWorkOrderResponseParams struct {
	// 创建的工单信息
	WorkOrderSet []*WorkOrderTinyInfo `json:"WorkOrderSet,omitnil,omitempty" name:"WorkOrderSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePersonnelVisitWorkOrderResponse struct {
	*tchttp.BaseResponse
	Response *CreatePersonnelVisitWorkOrderResponseParams `json:"Response"`
}

func (r *CreatePersonnelVisitWorkOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePersonnelVisitWorkOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePowerOffWorkOrderRequestParams struct {
	// 机房id
	IdcId *uint64 `json:"IdcId,omitnil,omitempty" name:"IdcId"`

	// 设备类型，server, netDevice
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 关电确认，1.授权时关电 2.关电前需要确认
	IsPowerOffConfirm *string `json:"IsPowerOffConfirm,omitnil,omitempty" name:"IsPowerOffConfirm"`

	// 设备sn列表
	DeviceSnList []*string `json:"DeviceSnList,omitnil,omitempty" name:"DeviceSnList"`

	// 关电前需要确认时必填
	PowerOffConfirmInfo *PowerOffConfirm `json:"PowerOffConfirmInfo,omitnil,omitempty" name:"PowerOffConfirmInfo"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type CreatePowerOffWorkOrderRequest struct {
	*tchttp.BaseRequest
	
	// 机房id
	IdcId *uint64 `json:"IdcId,omitnil,omitempty" name:"IdcId"`

	// 设备类型，server, netDevice
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 关电确认，1.授权时关电 2.关电前需要确认
	IsPowerOffConfirm *string `json:"IsPowerOffConfirm,omitnil,omitempty" name:"IsPowerOffConfirm"`

	// 设备sn列表
	DeviceSnList []*string `json:"DeviceSnList,omitnil,omitempty" name:"DeviceSnList"`

	// 关电前需要确认时必填
	PowerOffConfirmInfo *PowerOffConfirm `json:"PowerOffConfirmInfo,omitnil,omitempty" name:"PowerOffConfirmInfo"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *CreatePowerOffWorkOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePowerOffWorkOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdcId")
	delete(f, "DeviceType")
	delete(f, "IsPowerOffConfirm")
	delete(f, "DeviceSnList")
	delete(f, "PowerOffConfirmInfo")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePowerOffWorkOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePowerOffWorkOrderResponseParams struct {
	// 创建的工单信息
	WorkOrderSet []*WorkOrderTinyInfo `json:"WorkOrderSet,omitnil,omitempty" name:"WorkOrderSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePowerOffWorkOrderResponse struct {
	*tchttp.BaseResponse
	Response *CreatePowerOffWorkOrderResponseParams `json:"Response"`
}

func (r *CreatePowerOffWorkOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePowerOffWorkOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePowerOnWorkOrderRequestParams struct {
	// 机房id
	IdcId *uint64 `json:"IdcId,omitnil,omitempty" name:"IdcId"`

	// 设备类型，server, netDevice
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 设备sn列表
	DeviceSnList []*string `json:"DeviceSnList,omitnil,omitempty" name:"DeviceSnList"`
}

type CreatePowerOnWorkOrderRequest struct {
	*tchttp.BaseRequest
	
	// 机房id
	IdcId *uint64 `json:"IdcId,omitnil,omitempty" name:"IdcId"`

	// 设备类型，server, netDevice
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 设备sn列表
	DeviceSnList []*string `json:"DeviceSnList,omitnil,omitempty" name:"DeviceSnList"`
}

func (r *CreatePowerOnWorkOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePowerOnWorkOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdcId")
	delete(f, "DeviceType")
	delete(f, "DeviceSnList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePowerOnWorkOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePowerOnWorkOrderResponseParams struct {
	// 创建的工单信息
	WorkOrderSet []*WorkOrderTinyInfo `json:"WorkOrderSet,omitnil,omitempty" name:"WorkOrderSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePowerOnWorkOrderResponse struct {
	*tchttp.BaseResponse
	Response *CreatePowerOnWorkOrderResponseParams `json:"Response"`
}

func (r *CreatePowerOnWorkOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePowerOnWorkOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateQuitWorkOrderRequestParams struct {
	// 机房id
	IdcId *uint64 `json:"IdcId,omitnil,omitempty" name:"IdcId"`

	// 设备类型，server, netDevice, otherDevice
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 下架选择 1.自行解决 2.由腾讯IDC负责
	StuffOption *string `json:"StuffOption,omitnil,omitempty" name:"StuffOption"`

	// 关电确认 1.授权时关电 2.关电前需要确认
	IsPowerOffConfirm *string `json:"IsPowerOffConfirm,omitnil,omitempty" name:"IsPowerOffConfirm"`

	// 设备sn列表
	DeviceSnList []*string `json:"DeviceSnList,omitnil,omitempty" name:"DeviceSnList"`

	// 交接方式 1.物流上门收货 2.客户上门自提
	HandoverMethod *string `json:"HandoverMethod,omitnil,omitempty" name:"HandoverMethod"`

	// 自行解决必填
	SelfOperationInfo *SelfOperation `json:"SelfOperationInfo,omitnil,omitempty" name:"SelfOperationInfo"`

	// 关电前需要确认时必填
	PowerOffConfirmInfo *PowerOffConfirm `json:"PowerOffConfirmInfo,omitnil,omitempty" name:"PowerOffConfirmInfo"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 物流上门收货必传
	LogisticsReceipt *LogisticsReceipt `json:"LogisticsReceipt,omitnil,omitempty" name:"LogisticsReceipt"`

	// 客户上门自提必传
	CustomerReceipt *CustomerReceipt `json:"CustomerReceipt,omitnil,omitempty" name:"CustomerReceipt"`
}

type CreateQuitWorkOrderRequest struct {
	*tchttp.BaseRequest
	
	// 机房id
	IdcId *uint64 `json:"IdcId,omitnil,omitempty" name:"IdcId"`

	// 设备类型，server, netDevice, otherDevice
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 下架选择 1.自行解决 2.由腾讯IDC负责
	StuffOption *string `json:"StuffOption,omitnil,omitempty" name:"StuffOption"`

	// 关电确认 1.授权时关电 2.关电前需要确认
	IsPowerOffConfirm *string `json:"IsPowerOffConfirm,omitnil,omitempty" name:"IsPowerOffConfirm"`

	// 设备sn列表
	DeviceSnList []*string `json:"DeviceSnList,omitnil,omitempty" name:"DeviceSnList"`

	// 交接方式 1.物流上门收货 2.客户上门自提
	HandoverMethod *string `json:"HandoverMethod,omitnil,omitempty" name:"HandoverMethod"`

	// 自行解决必填
	SelfOperationInfo *SelfOperation `json:"SelfOperationInfo,omitnil,omitempty" name:"SelfOperationInfo"`

	// 关电前需要确认时必填
	PowerOffConfirmInfo *PowerOffConfirm `json:"PowerOffConfirmInfo,omitnil,omitempty" name:"PowerOffConfirmInfo"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 物流上门收货必传
	LogisticsReceipt *LogisticsReceipt `json:"LogisticsReceipt,omitnil,omitempty" name:"LogisticsReceipt"`

	// 客户上门自提必传
	CustomerReceipt *CustomerReceipt `json:"CustomerReceipt,omitnil,omitempty" name:"CustomerReceipt"`
}

func (r *CreateQuitWorkOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateQuitWorkOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdcId")
	delete(f, "DeviceType")
	delete(f, "StuffOption")
	delete(f, "IsPowerOffConfirm")
	delete(f, "DeviceSnList")
	delete(f, "HandoverMethod")
	delete(f, "SelfOperationInfo")
	delete(f, "PowerOffConfirmInfo")
	delete(f, "Remark")
	delete(f, "LogisticsReceipt")
	delete(f, "CustomerReceipt")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateQuitWorkOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateQuitWorkOrderResponseParams struct {
	// 创建的工单信息
	WorkOrderSet []*WorkOrderTinyInfo `json:"WorkOrderSet,omitnil,omitempty" name:"WorkOrderSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateQuitWorkOrderResponse struct {
	*tchttp.BaseResponse
	Response *CreateQuitWorkOrderResponseParams `json:"Response"`
}

func (r *CreateQuitWorkOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateQuitWorkOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRackOffWorkOrderRequestParams struct {
	// 机房id
	IdcId *uint64 `json:"IdcId,omitnil,omitempty" name:"IdcId"`

	// 设备类型，server, netDevice
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 下架人员 1.自行解决 2.由腾讯IDC负责
	StuffOption *string `json:"StuffOption,omitnil,omitempty" name:"StuffOption"`

	// 关电确认 1.授权时关电 2.关电前需要确认
	IsPowerOffConfirm *string `json:"IsPowerOffConfirm,omitnil,omitempty" name:"IsPowerOffConfirm"`

	// 设备sn列表
	DeviceSnList []*string `json:"DeviceSnList,omitnil,omitempty" name:"DeviceSnList"`

	// 自行解决必填
	SelfOperationInfo *SelfOperation `json:"SelfOperationInfo,omitnil,omitempty" name:"SelfOperationInfo"`

	// 关电前需要确认时必填
	PowerOffConfirmInfo *PowerOffConfirm `json:"PowerOffConfirmInfo,omitnil,omitempty" name:"PowerOffConfirmInfo"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type CreateRackOffWorkOrderRequest struct {
	*tchttp.BaseRequest
	
	// 机房id
	IdcId *uint64 `json:"IdcId,omitnil,omitempty" name:"IdcId"`

	// 设备类型，server, netDevice
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 下架人员 1.自行解决 2.由腾讯IDC负责
	StuffOption *string `json:"StuffOption,omitnil,omitempty" name:"StuffOption"`

	// 关电确认 1.授权时关电 2.关电前需要确认
	IsPowerOffConfirm *string `json:"IsPowerOffConfirm,omitnil,omitempty" name:"IsPowerOffConfirm"`

	// 设备sn列表
	DeviceSnList []*string `json:"DeviceSnList,omitnil,omitempty" name:"DeviceSnList"`

	// 自行解决必填
	SelfOperationInfo *SelfOperation `json:"SelfOperationInfo,omitnil,omitempty" name:"SelfOperationInfo"`

	// 关电前需要确认时必填
	PowerOffConfirmInfo *PowerOffConfirm `json:"PowerOffConfirmInfo,omitnil,omitempty" name:"PowerOffConfirmInfo"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *CreateRackOffWorkOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRackOffWorkOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdcId")
	delete(f, "DeviceType")
	delete(f, "StuffOption")
	delete(f, "IsPowerOffConfirm")
	delete(f, "DeviceSnList")
	delete(f, "SelfOperationInfo")
	delete(f, "PowerOffConfirmInfo")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRackOffWorkOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRackOffWorkOrderResponseParams struct {
	// 创建的工单信息
	WorkOrderSet []*WorkOrderTinyInfo `json:"WorkOrderSet,omitnil,omitempty" name:"WorkOrderSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRackOffWorkOrderResponse struct {
	*tchttp.BaseResponse
	Response *CreateRackOffWorkOrderResponseParams `json:"Response"`
}

func (r *CreateRackOffWorkOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRackOffWorkOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRackOnWorkOrderRequestParams struct {
	// 机房id
	IdcId *uint64 `json:"IdcId,omitnil,omitempty" name:"IdcId"`

	// 设备类型，server, netDevice
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 上架人员 1.自行解决 2.由腾讯IDC负责
	StuffOption *string `json:"StuffOption,omitnil,omitempty" name:"StuffOption"`

	// 上架后是否开电
	WithPowerOn *bool `json:"WithPowerOn,omitnil,omitempty" name:"WithPowerOn"`

	// 服务器收货列表
	DeviceRackOnList []*DeviceRackOn `json:"DeviceRackOnList,omitnil,omitempty" name:"DeviceRackOnList"`

	// 自行解决必填
	SelfOperationInfo *SelfOperation `json:"SelfOperationInfo,omitnil,omitempty" name:"SelfOperationInfo"`
}

type CreateRackOnWorkOrderRequest struct {
	*tchttp.BaseRequest
	
	// 机房id
	IdcId *uint64 `json:"IdcId,omitnil,omitempty" name:"IdcId"`

	// 设备类型，server, netDevice
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 上架人员 1.自行解决 2.由腾讯IDC负责
	StuffOption *string `json:"StuffOption,omitnil,omitempty" name:"StuffOption"`

	// 上架后是否开电
	WithPowerOn *bool `json:"WithPowerOn,omitnil,omitempty" name:"WithPowerOn"`

	// 服务器收货列表
	DeviceRackOnList []*DeviceRackOn `json:"DeviceRackOnList,omitnil,omitempty" name:"DeviceRackOnList"`

	// 自行解决必填
	SelfOperationInfo *SelfOperation `json:"SelfOperationInfo,omitnil,omitempty" name:"SelfOperationInfo"`
}

func (r *CreateRackOnWorkOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRackOnWorkOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdcId")
	delete(f, "DeviceType")
	delete(f, "StuffOption")
	delete(f, "WithPowerOn")
	delete(f, "DeviceRackOnList")
	delete(f, "SelfOperationInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRackOnWorkOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRackOnWorkOrderResponseParams struct {
	// 创建的工单信息
	WorkOrderSet []*WorkOrderTinyInfo `json:"WorkOrderSet,omitnil,omitempty" name:"WorkOrderSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRackOnWorkOrderResponse struct {
	*tchttp.BaseResponse
	Response *CreateRackOnWorkOrderResponseParams `json:"Response"`
}

func (r *CreateRackOnWorkOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRackOnWorkOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReceivingWorkOrderRequestParams struct {
	// 机房id
	IdcId *uint64 `json:"IdcId,omitnil,omitempty" name:"IdcId"`

	// 设备类型，server, netDevice, wire, otherDevice
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 进入时间
	EntryTime *string `json:"EntryTime,omitnil,omitempty" name:"EntryTime"`

	// 1.收货-仅核对外包装完整和数量，不开箱 2.验收-开箱核对型号SN一致
	ReceivingOperation *string `json:"ReceivingOperation,omitnil,omitempty" name:"ReceivingOperation"`

	// 是否快递寄件
	IsExpressDelivery *bool `json:"IsExpressDelivery,omitnil,omitempty" name:"IsExpressDelivery"`

	// 快递寄件信息,快递寄件必填
	ExpressInfo *ExpressDelivery `json:"ExpressInfo,omitnil,omitempty" name:"ExpressInfo"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 服务器收货列表
	ServerDeviceList []*ServerReceivingInfo `json:"ServerDeviceList,omitnil,omitempty" name:"ServerDeviceList"`

	// 网络设备收货列表
	NetDeviceList []*NetReceivingInfo `json:"NetDeviceList,omitnil,omitempty" name:"NetDeviceList"`

	// 线材收货列表
	WireDeviceList []*WireReceivingInfo `json:"WireDeviceList,omitnil,omitempty" name:"WireDeviceList"`

	// 其他设备收货列表
	OtherDeviceList []*OtherDevReceivingInfo `json:"OtherDeviceList,omitnil,omitempty" name:"OtherDeviceList"`
}

type CreateReceivingWorkOrderRequest struct {
	*tchttp.BaseRequest
	
	// 机房id
	IdcId *uint64 `json:"IdcId,omitnil,omitempty" name:"IdcId"`

	// 设备类型，server, netDevice, wire, otherDevice
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 进入时间
	EntryTime *string `json:"EntryTime,omitnil,omitempty" name:"EntryTime"`

	// 1.收货-仅核对外包装完整和数量，不开箱 2.验收-开箱核对型号SN一致
	ReceivingOperation *string `json:"ReceivingOperation,omitnil,omitempty" name:"ReceivingOperation"`

	// 是否快递寄件
	IsExpressDelivery *bool `json:"IsExpressDelivery,omitnil,omitempty" name:"IsExpressDelivery"`

	// 快递寄件信息,快递寄件必填
	ExpressInfo *ExpressDelivery `json:"ExpressInfo,omitnil,omitempty" name:"ExpressInfo"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 服务器收货列表
	ServerDeviceList []*ServerReceivingInfo `json:"ServerDeviceList,omitnil,omitempty" name:"ServerDeviceList"`

	// 网络设备收货列表
	NetDeviceList []*NetReceivingInfo `json:"NetDeviceList,omitnil,omitempty" name:"NetDeviceList"`

	// 线材收货列表
	WireDeviceList []*WireReceivingInfo `json:"WireDeviceList,omitnil,omitempty" name:"WireDeviceList"`

	// 其他设备收货列表
	OtherDeviceList []*OtherDevReceivingInfo `json:"OtherDeviceList,omitnil,omitempty" name:"OtherDeviceList"`
}

func (r *CreateReceivingWorkOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReceivingWorkOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdcId")
	delete(f, "DeviceType")
	delete(f, "EntryTime")
	delete(f, "ReceivingOperation")
	delete(f, "IsExpressDelivery")
	delete(f, "ExpressInfo")
	delete(f, "Remark")
	delete(f, "ServerDeviceList")
	delete(f, "NetDeviceList")
	delete(f, "WireDeviceList")
	delete(f, "OtherDeviceList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateReceivingWorkOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReceivingWorkOrderResponseParams struct {
	// 创建的工单信息
	WorkOrderSet []*WorkOrderTinyInfo `json:"WorkOrderSet,omitnil,omitempty" name:"WorkOrderSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateReceivingWorkOrderResponse struct {
	*tchttp.BaseResponse
	Response *CreateReceivingWorkOrderResponseParams `json:"Response"`
}

func (r *CreateReceivingWorkOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReceivingWorkOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateServerModelRequestParams struct {
	// 设备型号详情
	ModelDetail *ServerModel `json:"ModelDetail,omitnil,omitempty" name:"ModelDetail"`
}

type CreateServerModelRequest struct {
	*tchttp.BaseRequest
	
	// 设备型号详情
	ModelDetail *ServerModel `json:"ModelDetail,omitnil,omitempty" name:"ModelDetail"`
}

func (r *CreateServerModelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServerModelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelDetail")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateServerModelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateServerModelResponseParams struct {
	// 型号名称
	DevModel *string `json:"DevModel,omitnil,omitempty" name:"DevModel"`

	// 版本
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateServerModelResponse struct {
	*tchttp.BaseResponse
	Response *CreateServerModelResponseParams `json:"Response"`
}

func (r *CreateServerModelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServerModelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSpeciallyQuitWorkOrderRequestParams struct {
	// 机房id
	IdcId *uint64 `json:"IdcId,omitnil,omitempty" name:"IdcId"`

	// 设备类型：otherDevice。此接口只支持其他设备
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 交接方式 1.物流上门收货 2.客户上门自提
	HandoverMethod *string `json:"HandoverMethod,omitnil,omitempty" name:"HandoverMethod"`

	// 物流上门收货必传
	LogisticsReceipt *LogisticsReceipt `json:"LogisticsReceipt,omitnil,omitempty" name:"LogisticsReceipt"`

	// 客户上门自提必传
	CustomerReceipt *CustomerReceipt `json:"CustomerReceipt,omitnil,omitempty" name:"CustomerReceipt"`

	// 备注信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 当设备类型为otherDevice，此参数必传
	OtherDeviceList []*OtherDevReceivingInfo `json:"OtherDeviceList,omitnil,omitempty" name:"OtherDeviceList"`
}

type CreateSpeciallyQuitWorkOrderRequest struct {
	*tchttp.BaseRequest
	
	// 机房id
	IdcId *uint64 `json:"IdcId,omitnil,omitempty" name:"IdcId"`

	// 设备类型：otherDevice。此接口只支持其他设备
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 交接方式 1.物流上门收货 2.客户上门自提
	HandoverMethod *string `json:"HandoverMethod,omitnil,omitempty" name:"HandoverMethod"`

	// 物流上门收货必传
	LogisticsReceipt *LogisticsReceipt `json:"LogisticsReceipt,omitnil,omitempty" name:"LogisticsReceipt"`

	// 客户上门自提必传
	CustomerReceipt *CustomerReceipt `json:"CustomerReceipt,omitnil,omitempty" name:"CustomerReceipt"`

	// 备注信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 当设备类型为otherDevice，此参数必传
	OtherDeviceList []*OtherDevReceivingInfo `json:"OtherDeviceList,omitnil,omitempty" name:"OtherDeviceList"`
}

func (r *CreateSpeciallyQuitWorkOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSpeciallyQuitWorkOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdcId")
	delete(f, "DeviceType")
	delete(f, "HandoverMethod")
	delete(f, "LogisticsReceipt")
	delete(f, "CustomerReceipt")
	delete(f, "Remark")
	delete(f, "OtherDeviceList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSpeciallyQuitWorkOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSpeciallyQuitWorkOrderResponseParams struct {
	// 创建的工单信息
	WorkOrderSet []*WorkOrderTinyInfo `json:"WorkOrderSet,omitnil,omitempty" name:"WorkOrderSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSpeciallyQuitWorkOrderResponse struct {
	*tchttp.BaseResponse
	Response *CreateSpeciallyQuitWorkOrderResponseParams `json:"Response"`
}

func (r *CreateSpeciallyQuitWorkOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSpeciallyQuitWorkOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomerInfo struct {
	// 公司全称
	CustomerName *string `json:"CustomerName,omitnil,omitempty" name:"CustomerName"`

	// 公司简称
	ShortCustomerName *string `json:"ShortCustomerName,omitnil,omitempty" name:"ShortCustomerName"`

	// 是否全托管用户
	WholeFlag *bool `json:"WholeFlag,omitnil,omitempty" name:"WholeFlag"`
}

type CustomerReceipt struct {
	// 自提人员姓名
	PickUpStuff *string `json:"PickUpStuff,omitnil,omitempty" name:"PickUpStuff"`

	// 自提人电话
	PickUpStuffContact *string `json:"PickUpStuffContact,omitnil,omitempty" name:"PickUpStuffContact"`

	// 自提人身份证号
	PickUpStuffIDCard *string `json:"PickUpStuffIDCard,omitnil,omitempty" name:"PickUpStuffIDCard"`

	// 自提时间
	PickUpTime *string `json:"PickUpTime,omitnil,omitempty" name:"PickUpTime"`
}

// Predefined struct for user
type DescribeAvailableModelListRequestParams struct {
	// 机房ID
	IdcId *uint64 `json:"IdcId,omitnil,omitempty" name:"IdcId"`

	// server 服务器，netDevice 网络设备
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`
}

type DescribeAvailableModelListRequest struct {
	*tchttp.BaseRequest
	
	// 机房ID
	IdcId *uint64 `json:"IdcId,omitnil,omitempty" name:"IdcId"`

	// server 服务器，netDevice 网络设备
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`
}

func (r *DescribeAvailableModelListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAvailableModelListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdcId")
	delete(f, "DeviceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAvailableModelListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAvailableModelListResponseParams struct {
	// 机房内可用的设备型号及版本列表
	ModelVersionSet []*AvailableModelVersion `json:"ModelVersionSet,omitnil,omitempty" name:"ModelVersionSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAvailableModelListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAvailableModelListResponseParams `json:"Response"`
}

func (r *DescribeAvailableModelListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAvailableModelListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCampusListRequestParams struct {

}

type DescribeCampusListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeCampusListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCampusListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCampusListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCampusListResponseParams struct {
	// 客户可操作的园区列表
	CampusSet []*Campus `json:"CampusSet,omitnil,omitempty" name:"CampusSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCampusListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCampusListResponseParams `json:"Response"`
}

func (r *DescribeCampusListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCampusListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCommonServiceWorkOrderDetailRequestParams struct {
	// 工单ID
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`
}

type DescribeCommonServiceWorkOrderDetailRequest struct {
	*tchttp.BaseRequest
	
	// 工单ID
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`
}

func (r *DescribeCommonServiceWorkOrderDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCommonServiceWorkOrderDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrderId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCommonServiceWorkOrderDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCommonServiceWorkOrderDetailResponseParams struct {
	// 进度
	StepSet []*OrderStep `json:"StepSet,omitnil,omitempty" name:"StepSet"`

	// 基本信息
	BaseInfo *CommonServiceBaseInfo `json:"BaseInfo,omitnil,omitempty" name:"BaseInfo"`

	// 设备信息
	DeviceSet []*DevicePosition `json:"DeviceSet,omitnil,omitempty" name:"DeviceSet"`

	// 工单状态
	OrderStatus *string `json:"OrderStatus,omitnil,omitempty" name:"OrderStatus"`

	// 拒绝原因
	RejectReason *string `json:"RejectReason,omitnil,omitempty" name:"RejectReason"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCommonServiceWorkOrderDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCommonServiceWorkOrderDetailResponseParams `json:"Response"`
}

func (r *DescribeCommonServiceWorkOrderDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCommonServiceWorkOrderDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomerInfoRequestParams struct {

}

type DescribeCustomerInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeCustomerInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomerInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomerInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomerInfoResponseParams struct {
	// 客户信息
	CustomerInfo *CustomerInfo `json:"CustomerInfo,omitnil,omitempty" name:"CustomerInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCustomerInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomerInfoResponseParams `json:"Response"`
}

func (r *DescribeCustomerInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomerInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceListRequestParams struct {
	// 设备类型 server 服务器，netDevice 网络设备，otherDevice 其他设备
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 
	// <li><strong>rack-id</strong></li> <p style="padding-left: 30px;">按照【<strong>机架ID</strong>】进行过滤。例如：15082。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p><p style="padding-left: 30px;"></p> <li><strong> sn</strong></li> <p style="padding-left: 30px;">按照【<strong>设备 SN 码</strong>】进行过滤，设备 SN 例如：TEN948P004。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong> idc-id</strong></li> <p style="padding-left: 30px;">按照【<strong>机房ID</strong>】进行过滤，机房ID例如：159。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>  <li><strong>idc-unit-id </strong></li> <p style="padding-left: 30px;">按照【<strong>机房管理单元ID</strong>】进行过滤，机房管理ID例如：568。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>server-type-id</strong></li> <p style="padding-left: 30px;">按照【<strong>机器子类型</strong>】进行过滤，只包含以下几种：1:服务器, 2:Twins主机, 3:Twins子机,5:虚拟机, 6:2U4S主机, 7:2U4S子机,8 Rack主机,9 Rack子机，例如： 1。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>status</strong></li> <p style="padding-left: 30px;">按照【<strong>设备状态</strong>】进行过滤，操作状态只包含：POWER_ON 设备开电，POWER_OFF 设备关电，RACK_OFF 未上架，MOVING 搬迁中 。例如： POWER_OFF。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>svr-is-special</strong></li> <p style="padding-left: 30px;">按照【<strong>是否</strong>】进行过滤，支持 0：自有，1 租用。例如： 1。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为1000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 传入目标服务，返回允许进行此服务的设备列表；可以和Filters一起使用。允许的值：('rackOn', 'powerOn', 'powerOff', 'rackOff', 'quit', 'moving'，'netDeviceCommon', 'serverCommon')
	DstService *string `json:"DstService,omitnil,omitempty" name:"DstService"`
}

type DescribeDeviceListRequest struct {
	*tchttp.BaseRequest
	
	// 设备类型 server 服务器，netDevice 网络设备，otherDevice 其他设备
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 
	// <li><strong>rack-id</strong></li> <p style="padding-left: 30px;">按照【<strong>机架ID</strong>】进行过滤。例如：15082。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p><p style="padding-left: 30px;"></p> <li><strong> sn</strong></li> <p style="padding-left: 30px;">按照【<strong>设备 SN 码</strong>】进行过滤，设备 SN 例如：TEN948P004。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong> idc-id</strong></li> <p style="padding-left: 30px;">按照【<strong>机房ID</strong>】进行过滤，机房ID例如：159。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>  <li><strong>idc-unit-id </strong></li> <p style="padding-left: 30px;">按照【<strong>机房管理单元ID</strong>】进行过滤，机房管理ID例如：568。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>server-type-id</strong></li> <p style="padding-left: 30px;">按照【<strong>机器子类型</strong>】进行过滤，只包含以下几种：1:服务器, 2:Twins主机, 3:Twins子机,5:虚拟机, 6:2U4S主机, 7:2U4S子机,8 Rack主机,9 Rack子机，例如： 1。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>status</strong></li> <p style="padding-left: 30px;">按照【<strong>设备状态</strong>】进行过滤，操作状态只包含：POWER_ON 设备开电，POWER_OFF 设备关电，RACK_OFF 未上架，MOVING 搬迁中 。例如： POWER_OFF。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>svr-is-special</strong></li> <p style="padding-left: 30px;">按照【<strong>是否</strong>】进行过滤，支持 0：自有，1 租用。例如： 1。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为1000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 传入目标服务，返回允许进行此服务的设备列表；可以和Filters一起使用。允许的值：('rackOn', 'powerOn', 'powerOff', 'rackOff', 'quit', 'moving'，'netDeviceCommon', 'serverCommon')
	DstService *string `json:"DstService,omitnil,omitempty" name:"DstService"`
}

func (r *DescribeDeviceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceType")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "DstService")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceListResponseParams struct {
	// 总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 服务器列表
	DeviceSet []*Device `json:"DeviceSet,omitnil,omitempty" name:"DeviceSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDeviceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceListResponseParams `json:"Response"`
}

func (r *DescribeDeviceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceWorkOrderDetailRequestParams struct {
	// 工单ID
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`
}

type DescribeDeviceWorkOrderDetailRequest struct {
	*tchttp.BaseRequest
	
	// 工单ID
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`
}

func (r *DescribeDeviceWorkOrderDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceWorkOrderDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrderId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceWorkOrderDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceWorkOrderDetailResponseParams struct {
	// 工单ID
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// 服务类型
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// 工单类型
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// 工单状态
	OrderStatus *string `json:"OrderStatus,omitnil,omitempty" name:"OrderStatus"`

	// 工单流程状态
	StepSet []*OrderStep `json:"StepSet,omitnil,omitempty" name:"StepSet"`

	// 工单设备信息
	DeviceSet []*DeviceHistory `json:"DeviceSet,omitnil,omitempty" name:"DeviceSet"`

	// 工单的入参信息
	BaseInfo *DeviceOrderBaseInfo `json:"BaseInfo,omitnil,omitempty" name:"BaseInfo"`

	// 工单的拒绝原因，工单状态为reject的时候返回
	RejectReason *string `json:"RejectReason,omitnil,omitempty" name:"RejectReason"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDeviceWorkOrderDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceWorkOrderDetailResponseParams `json:"Response"`
}

func (r *DescribeDeviceWorkOrderDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceWorkOrderDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIdcUnitAssetDetailRequestParams struct {
	// 机房管理单元ID
	IdcUnitId *uint64 `json:"IdcUnitId,omitnil,omitempty" name:"IdcUnitId"`
}

type DescribeIdcUnitAssetDetailRequest struct {
	*tchttp.BaseRequest
	
	// 机房管理单元ID
	IdcUnitId *uint64 `json:"IdcUnitId,omitnil,omitempty" name:"IdcUnitId"`
}

func (r *DescribeIdcUnitAssetDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIdcUnitAssetDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdcUnitId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIdcUnitAssetDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIdcUnitAssetDetailResponseParams struct {
	// 机房管理单元详情
	IdcUnitDetail *IdcUnitInfo `json:"IdcUnitDetail,omitnil,omitempty" name:"IdcUnitDetail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIdcUnitAssetDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIdcUnitAssetDetailResponseParams `json:"Response"`
}

func (r *DescribeIdcUnitAssetDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIdcUnitAssetDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIdcUnitDetailRequestParams struct {
	// 机房管理单元ID
	IdcUnitId *uint64 `json:"IdcUnitId,omitnil,omitempty" name:"IdcUnitId"`
}

type DescribeIdcUnitDetailRequest struct {
	*tchttp.BaseRequest
	
	// 机房管理单元ID
	IdcUnitId *uint64 `json:"IdcUnitId,omitnil,omitempty" name:"IdcUnitId"`
}

func (r *DescribeIdcUnitDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIdcUnitDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdcUnitId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIdcUnitDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIdcUnitDetailResponseParams struct {
	// 机房管理单元详情
	IdcUnitDetail *IdcUnitInfo `json:"IdcUnitDetail,omitnil,omitempty" name:"IdcUnitDetail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIdcUnitDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIdcUnitDetailResponseParams `json:"Response"`
}

func (r *DescribeIdcUnitDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIdcUnitDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIdcsRequestParams struct {

}

type DescribeIdcsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeIdcsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIdcsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIdcsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIdcsResponseParams struct {
	// 机房管理单元列表
	IdcSet []*Idc `json:"IdcSet,omitnil,omitempty" name:"IdcSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIdcsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIdcsResponseParams `json:"Response"`
}

func (r *DescribeIdcsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIdcsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelEvaluationWorkOrderDetailRequestParams struct {
	// 工单ID
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`
}

type DescribeModelEvaluationWorkOrderDetailRequest struct {
	*tchttp.BaseRequest
	
	// 工单ID
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`
}

func (r *DescribeModelEvaluationWorkOrderDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelEvaluationWorkOrderDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrderId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelEvaluationWorkOrderDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelEvaluationWorkOrderDetailResponseParams struct {
	// 工单进度
	StepSet []*OrderStep `json:"StepSet,omitnil,omitempty" name:"StepSet"`

	// 工单详情
	BaseInfo *ModelEvaluationBaseInfo `json:"BaseInfo,omitnil,omitempty" name:"BaseInfo"`

	// 评估中的网络设备型号详情
	NetDeviceModelSet []*ModelVersionDetail `json:"NetDeviceModelSet,omitnil,omitempty" name:"NetDeviceModelSet"`

	// 评估中的服务器型号详情
	ServerModelSet []*ModelVersionDetail `json:"ServerModelSet,omitnil,omitempty" name:"ServerModelSet"`

	// 订单状态, process 处理中 ，reject 已拒绝 ，finish 已完成，exception 异常
	OrderStatus *string `json:"OrderStatus,omitnil,omitempty" name:"OrderStatus"`

	// 工单拒绝原因
	RejectReason *string `json:"RejectReason,omitnil,omitempty" name:"RejectReason"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModelEvaluationWorkOrderDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelEvaluationWorkOrderDetailResponseParams `json:"Response"`
}

func (r *DescribeModelEvaluationWorkOrderDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelEvaluationWorkOrderDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelRequestParams struct {
	// 服务器设备型号
	DevModel *string `json:"DevModel,omitnil,omitempty" name:"DevModel"`

	// 园区ID
	CampusId *uint64 `json:"CampusId,omitnil,omitempty" name:"CampusId"`

	// 设备类型，服务器传入 server，网络设备传入 netDevice
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 是否只返回已评估的版本
	Checked *bool `json:"Checked,omitnil,omitempty" name:"Checked"`
}

type DescribeModelRequest struct {
	*tchttp.BaseRequest
	
	// 服务器设备型号
	DevModel *string `json:"DevModel,omitnil,omitempty" name:"DevModel"`

	// 园区ID
	CampusId *uint64 `json:"CampusId,omitnil,omitempty" name:"CampusId"`

	// 设备类型，服务器传入 server，网络设备传入 netDevice
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 是否只返回已评估的版本
	Checked *bool `json:"Checked,omitnil,omitempty" name:"Checked"`
}

func (r *DescribeModelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DevModel")
	delete(f, "CampusId")
	delete(f, "DeviceType")
	delete(f, "Checked")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelResponseParams struct {
	// 设备型号详情
	ModelSet []*ModelVersionDetail `json:"ModelSet,omitnil,omitempty" name:"ModelSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModelResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelResponseParams `json:"Response"`
}

func (r *DescribeModelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelTemplateRequestParams struct {
	// 型号类型，只支持传入 server 和 netDevice
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`
}

type DescribeModelTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 型号类型，只支持传入 server 和 netDevice
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`
}

func (r *DescribeModelTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelTemplateResponseParams struct {
	// 该型号模板的选项列表
	TemplateDetail []*TemplateOption `json:"TemplateDetail,omitnil,omitempty" name:"TemplateDetail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModelTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelTemplateResponseParams `json:"Response"`
}

func (r *DescribeModelTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelVersionListRequestParams struct {
	// 型号类型，只支持传入 netDevice 和 server
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// model-name  型号名称  类型：String  必选：否
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否已评估
	Checked *bool `json:"Checked,omitnil,omitempty" name:"Checked"`

	// 园区ID，当 Checked 参数传 True 时，该参数必须传值
	CampusId *uint64 `json:"CampusId,omitnil,omitempty" name:"CampusId"`
}

type DescribeModelVersionListRequest struct {
	*tchttp.BaseRequest
	
	// 型号类型，只支持传入 netDevice 和 server
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// model-name  型号名称  类型：String  必选：否
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否已评估
	Checked *bool `json:"Checked,omitnil,omitempty" name:"Checked"`

	// 园区ID，当 Checked 参数传 True 时，该参数必须传值
	CampusId *uint64 `json:"CampusId,omitnil,omitempty" name:"CampusId"`
}

func (r *DescribeModelVersionListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelVersionListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceType")
	delete(f, "Filters")
	delete(f, "Checked")
	delete(f, "CampusId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelVersionListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelVersionListResponseParams struct {
	// 型号和对应的版本数量
	ModelVersionSet []*ModelVersionCount `json:"ModelVersionSet,omitnil,omitempty" name:"ModelVersionSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModelVersionListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelVersionListResponseParams `json:"Response"`
}

func (r *DescribeModelVersionListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelVersionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePersonnelVisitWorkOrderDetailRequestParams struct {
	// 工单ID
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`
}

type DescribePersonnelVisitWorkOrderDetailRequest struct {
	*tchttp.BaseRequest
	
	// 工单ID
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`
}

func (r *DescribePersonnelVisitWorkOrderDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePersonnelVisitWorkOrderDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrderId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePersonnelVisitWorkOrderDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePersonnelVisitWorkOrderDetailResponseParams struct {
	// 工单进度	
	StepSet []*OrderStep `json:"StepSet,omitnil,omitempty" name:"StepSet"`

	// 工单详情
	BaseInfo *PersonnelVisitBaseInfo `json:"BaseInfo,omitnil,omitempty" name:"BaseInfo"`

	// 到访人员详情
	PersonnelSet []*Personnel `json:"PersonnelSet,omitnil,omitempty" name:"PersonnelSet"`

	// 工单状态 订单状态, processing 处理中 ，reject 已拒绝 ，finish 已完成，exception 异常
	OrderStatus *string `json:"OrderStatus,omitnil,omitempty" name:"OrderStatus"`

	// 拒绝原因
	RejectReason *string `json:"RejectReason,omitnil,omitempty" name:"RejectReason"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePersonnelVisitWorkOrderDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribePersonnelVisitWorkOrderDetailResponseParams `json:"Response"`
}

func (r *DescribePersonnelVisitWorkOrderDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePersonnelVisitWorkOrderDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePositionStatusSummaryRequestParams struct {
	// <li><strong>rack-id</strong></li> <p style="padding-left: 30px;">按照【<strong>机架ID</strong>】进行过滤。例如：15082。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p><p style="padding-left: 30px;"></p> <li><strong> rack-name</strong></li> <p style="padding-left: 30px;">按照【<strong>机架名称</strong>】进行过滤，机架名称例如：M301-E10。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong> idc-id</strong></li> <p style="padding-left: 30px;">按照【<strong>机房ID</strong>】进行过滤，机房ID例如：159。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>  <li><strong>idc-unit-id </strong></li> <p style="padding-left: 30px;">按照【<strong>机房管理单元ID</strong>】进行过滤，机房管理ID例如：568。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>position-status</strong></li> <p style="padding-left: 30px;">按照【<strong>机位状态</strong>】进行过滤，机位状态只包含以下四种：机位状态,0 空闲,1 已用,2 不可用,3 预占用,4 预留，例如： 0。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>op-status</strong></li> <p style="padding-left: 30px;">按照【<strong>操作状态</strong>】进行过滤，操作状态只包含两种：FINISH 操作完成，PENDING 操作中，例如： PENDING。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribePositionStatusSummaryRequest struct {
	*tchttp.BaseRequest
	
	// <li><strong>rack-id</strong></li> <p style="padding-left: 30px;">按照【<strong>机架ID</strong>】进行过滤。例如：15082。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p><p style="padding-left: 30px;"></p> <li><strong> rack-name</strong></li> <p style="padding-left: 30px;">按照【<strong>机架名称</strong>】进行过滤，机架名称例如：M301-E10。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong> idc-id</strong></li> <p style="padding-left: 30px;">按照【<strong>机房ID</strong>】进行过滤，机房ID例如：159。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>  <li><strong>idc-unit-id </strong></li> <p style="padding-left: 30px;">按照【<strong>机房管理单元ID</strong>】进行过滤，机房管理ID例如：568。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>position-status</strong></li> <p style="padding-left: 30px;">按照【<strong>机位状态</strong>】进行过滤，机位状态只包含以下四种：机位状态,0 空闲,1 已用,2 不可用,3 预占用,4 预留，例如： 0。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>op-status</strong></li> <p style="padding-left: 30px;">按照【<strong>操作状态</strong>】进行过滤，操作状态只包含两种：FINISH 操作完成，PENDING 操作中，例如： PENDING。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribePositionStatusSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePositionStatusSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePositionStatusSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePositionStatusSummaryResponseParams struct {
	// 总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 状态及对应数量
	StatusCountSet []*PositionStatusItem `json:"StatusCountSet,omitnil,omitempty" name:"StatusCountSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePositionStatusSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribePositionStatusSummaryResponseParams `json:"Response"`
}

func (r *DescribePositionStatusSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePositionStatusSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePositionsRequestParams struct {
	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API 简介中的相关小节。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <li><strong>rack-id</strong></li> <p style="padding-left: 30px;">按照【<strong>机架ID</strong>】进行过滤。例如：15082。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p><p style="padding-left: 30px;"></p> <li><strong> rack-name</strong></li> <p style="padding-left: 30px;">按照【<strong>机架名称</strong>】进行过滤，机架名称例如：M301-E10。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong> idc-id</strong></li> <p style="padding-left: 30px;">按照【<strong>机房ID</strong>】进行过滤，机房ID例如：159。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>  <li><strong>idc-unit-id </strong></li> <p style="padding-left: 30px;">按照【<strong>机房管理单元ID</strong>】进行过滤，机房管理ID例如：568。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>position-status</strong></li> <p style="padding-left: 30px;">按照【<strong>机位状态</strong>】进行过滤，机位状态只包含以下四种：机位状态,0 空闲,1 已用,2 不可用,3 预占用,4 预留，例如： 0。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>op-status</strong></li> <p style="padding-left: 30px;">按照【<strong>操作状态</strong>】进行过滤，操作状态只包含两种：FINISH 操作完成，PENDING 操作中，例如： PENDING。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribePositionsRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API 简介中的相关小节。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <li><strong>rack-id</strong></li> <p style="padding-left: 30px;">按照【<strong>机架ID</strong>】进行过滤。例如：15082。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p><p style="padding-left: 30px;"></p> <li><strong> rack-name</strong></li> <p style="padding-left: 30px;">按照【<strong>机架名称</strong>】进行过滤，机架名称例如：M301-E10。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong> idc-id</strong></li> <p style="padding-left: 30px;">按照【<strong>机房ID</strong>】进行过滤，机房ID例如：159。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>  <li><strong>idc-unit-id </strong></li> <p style="padding-left: 30px;">按照【<strong>机房管理单元ID</strong>】进行过滤，机房管理ID例如：568。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>position-status</strong></li> <p style="padding-left: 30px;">按照【<strong>机位状态</strong>】进行过滤，机位状态只包含以下四种：机位状态,0 空闲,1 已用,2 不可用,3 预占用,4 预留，例如： 0。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>op-status</strong></li> <p style="padding-left: 30px;">按照【<strong>操作状态</strong>】进行过滤，操作状态只包含两种：FINISH 操作完成，PENDING 操作中，例如： PENDING。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribePositionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePositionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePositionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePositionsResponseParams struct {
	// 客户拥有的机位列表
	PositionSet []*Position `json:"PositionSet,omitnil,omitempty" name:"PositionSet"`

	// 总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePositionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePositionsResponseParams `json:"Response"`
}

func (r *DescribePositionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePositionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRacksDistributionRequestParams struct {
	// 机房管理单元ID
	IdcUnitId *uint64 `json:"IdcUnitId,omitnil,omitempty" name:"IdcUnitId"`
}

type DescribeRacksDistributionRequest struct {
	*tchttp.BaseRequest
	
	// 机房管理单元ID
	IdcUnitId *uint64 `json:"IdcUnitId,omitnil,omitempty" name:"IdcUnitId"`
}

func (r *DescribeRacksDistributionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRacksDistributionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdcUnitId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRacksDistributionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRacksDistributionResponseParams struct {
	// 机架的用量分布
	DistributionSet []*Distribution `json:"DistributionSet,omitnil,omitempty" name:"DistributionSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRacksDistributionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRacksDistributionResponseParams `json:"Response"`
}

func (r *DescribeRacksDistributionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRacksDistributionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRacksRequestParams struct {
	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API 简介中的相关小节。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件
	// 
	// rack-id
	// 按照机架id进行过滤。
	// 类型：String
	// 必选：否
	// 
	// rack-name
	// 按照机架名称进行过滤。
	// 类型：String
	// 必选：否
	// 
	// idc-id
	// 按照机房id进行过滤。
	// 类型：String
	// 必选：否
	// 
	// idc-unit-id
	// 按照机房管理单元id过滤
	// 类型： String
	// 必选： 否
	// 
	// is-power-on
	// 按照是否开电进行过滤，支持传入 0 和 1。1 代表开电，0 代表关电
	// 类型： String
	// 必选： 否
	// 
	// hosting-type
	// 按照托管类型进行过滤。支持传入 CUSTOMER_MIX 代表客户混合，CUSTOMER_ONLY 代表客户独享 ，NOT_INIT 代表未初始化
	// 类型： String
	// 必选： 否
	// 
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 传入目标服务，返回允许进行此服务的机架列表；可以和Filters一起使用。允许的值：('rackPowerOn', 'rackPowerOff')
	DstService *string `json:"DstService,omitnil,omitempty" name:"DstService"`
}

type DescribeRacksRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API 简介中的相关小节。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件
	// 
	// rack-id
	// 按照机架id进行过滤。
	// 类型：String
	// 必选：否
	// 
	// rack-name
	// 按照机架名称进行过滤。
	// 类型：String
	// 必选：否
	// 
	// idc-id
	// 按照机房id进行过滤。
	// 类型：String
	// 必选：否
	// 
	// idc-unit-id
	// 按照机房管理单元id过滤
	// 类型： String
	// 必选： 否
	// 
	// is-power-on
	// 按照是否开电进行过滤，支持传入 0 和 1。1 代表开电，0 代表关电
	// 类型： String
	// 必选： 否
	// 
	// hosting-type
	// 按照托管类型进行过滤。支持传入 CUSTOMER_MIX 代表客户混合，CUSTOMER_ONLY 代表客户独享 ，NOT_INIT 代表未初始化
	// 类型： String
	// 必选： 否
	// 
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 传入目标服务，返回允许进行此服务的机架列表；可以和Filters一起使用。允许的值：('rackPowerOn', 'rackPowerOff')
	DstService *string `json:"DstService,omitnil,omitempty" name:"DstService"`
}

func (r *DescribeRacksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRacksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "DstService")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRacksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRacksResponseParams struct {
	// 客户拥有的机架列表
	RackSet []*Rack `json:"RackSet,omitnil,omitempty" name:"RackSet"`

	// 总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRacksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRacksResponseParams `json:"Response"`
}

func (r *DescribeRacksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRacksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceUsageRequestParams struct {
	// 过滤条件
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeResourceUsageRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeResourceUsageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceUsageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourceUsageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceUsageResponseParams struct {
	// 托管服务器数量
	HostingServerCount *uint64 `json:"HostingServerCount,omitnil,omitempty" name:"HostingServerCount"`

	// 租用服务器数量
	RentServerCount *uint64 `json:"RentServerCount,omitnil,omitempty" name:"RentServerCount"`

	// 网络设备数量
	NetDeviceCount *uint64 `json:"NetDeviceCount,omitnil,omitempty" name:"NetDeviceCount"`

	// 机架总数
	RackTotalCount *uint64 `json:"RackTotalCount,omitnil,omitempty" name:"RackTotalCount"`

	// 开电机架总数
	RackPowerOnCount *uint64 `json:"RackPowerOnCount,omitnil,omitempty" name:"RackPowerOnCount"`

	// 机位使用数量
	PositionUsedCount *uint64 `json:"PositionUsedCount,omitnil,omitempty" name:"PositionUsedCount"`

	// 机位总数
	PositionTotalCount *uint64 `json:"PositionTotalCount,omitnil,omitempty" name:"PositionTotalCount"`

	// 机架开电率，保留一位小数
	RackPowerOnRate *string `json:"RackPowerOnRate,omitnil,omitempty" name:"RackPowerOnRate"`

	// 机位使用率，
	PositionUsedRate *string `json:"PositionUsedRate,omitnil,omitempty" name:"PositionUsedRate"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeResourceUsageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourceUsageResponseParams `json:"Response"`
}

func (r *DescribeResourceUsageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkOrderListRequestParams struct {
	// 过滤条件。支持：service-type、order-type、order-status、order-id
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 通过sn过滤工单，上限10个
	SnList []*string `json:"SnList,omitnil,omitempty" name:"SnList"`

	// 起始
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 长度
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeWorkOrderListRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。支持：service-type、order-type、order-status、order-id
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 通过sn过滤工单，上限10个
	SnList []*string `json:"SnList,omitnil,omitempty" name:"SnList"`

	// 起始
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 长度
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeWorkOrderListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkOrderListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "SnList")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWorkOrderListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkOrderListResponseParams struct {
	// 总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 查询结果
	WorkOrderSet []*WorkOrderData `json:"WorkOrderSet,omitnil,omitempty" name:"WorkOrderSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWorkOrderListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWorkOrderListResponseParams `json:"Response"`
}

func (r *DescribeWorkOrderListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkOrderListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkOrderStatisticsRequestParams struct {

}

type DescribeWorkOrderStatisticsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeWorkOrderStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkOrderStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWorkOrderStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkOrderStatisticsResponseParams struct {
	// 总工单数量
	TotalNum *uint64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// 进行中的工单数量
	ProcessingNum *uint64 `json:"ProcessingNum,omitnil,omitempty" name:"ProcessingNum"`

	// 待确认的工单数量
	ConfirmingNum *uint64 `json:"ConfirmingNum,omitnil,omitempty" name:"ConfirmingNum"`

	// 完成的工单数量
	FinishNum *uint64 `json:"FinishNum,omitnil,omitempty" name:"FinishNum"`

	// 拒绝的工单数量
	RejectNum *uint64 `json:"RejectNum,omitnil,omitempty" name:"RejectNum"`

	// 异常的工单数量
	ExceptionNum *uint64 `json:"ExceptionNum,omitnil,omitempty" name:"ExceptionNum"`

	// 客户取消的工单数量
	CancelNum *uint64 `json:"CancelNum,omitnil,omitempty" name:"CancelNum"`

	// 围笼进出审核的工单数量
	CheckingNum *uint64 `json:"CheckingNum,omitnil,omitempty" name:"CheckingNum"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWorkOrderStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWorkOrderStatisticsResponseParams `json:"Response"`
}

func (r *DescribeWorkOrderStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkOrderStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkOrderTypesRequestParams struct {

}

type DescribeWorkOrderTypesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeWorkOrderTypesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkOrderTypesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWorkOrderTypesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkOrderTypesResponseParams struct {
	// 已收藏的工单类型列表
	CollectedWorkOderTypeSet []*WorkOrderTypeDetail `json:"CollectedWorkOderTypeSet,omitnil,omitempty" name:"CollectedWorkOderTypeSet"`

	// 全部工单类型列表
	WorkOrderFamilySet []*WorkOrderFamilyDetail `json:"WorkOrderFamilySet,omitnil,omitempty" name:"WorkOrderFamilySet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWorkOrderTypesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWorkOrderTypesResponseParams `json:"Response"`
}

func (r *DescribeWorkOrderTypesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkOrderTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Device struct {
	// 设备 SN 码
	Sn *string `json:"Sn,omitnil,omitempty" name:"Sn"`

	// 设备型号版本
	ModelVersion *string `json:"ModelVersion,omitnil,omitempty" name:"ModelVersion"`

	// 设备固资号。只有设备类型为服务器时才返回
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 0 自有，1 租用。只有设备类型为服务器时才返回
	SvrIsSpecial *uint64 `json:"SvrIsSpecial,omitnil,omitempty" name:"SvrIsSpecial"`

	// IP。
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 设备所属的机房名称
	IdcName *string `json:"IdcName,omitnil,omitempty" name:"IdcName"`

	// 设备所属的机房ID
	IdcId *uint64 `json:"IdcId,omitnil,omitempty" name:"IdcId"`

	// 设备所属的机房管理单元ID
	IdcUnitId *uint64 `json:"IdcUnitId,omitnil,omitempty" name:"IdcUnitId"`

	// 设备所属的机房管理单元名称
	IdcUnitName *string `json:"IdcUnitName,omitnil,omitempty" name:"IdcUnitName"`

	// 已上架设备所在的机架ID，未上架设备不返回
	RackId *uint64 `json:"RackId,omitnil,omitempty" name:"RackId"`

	// 服务器类型， 1 代表服务器， 7 代表 2U4S。只有设备类型为服务器时才返回
	ServerTypeId *uint64 `json:"ServerTypeId,omitnil,omitempty" name:"ServerTypeId"`

	// 已上架设备所在的机架名称，未上架设备不返回
	RackName *string `json:"RackName,omitnil,omitempty" name:"RackName"`

	// 已上架设备所在的机位编号，未上架设备不返回。只有设备类型为服务器时才返回
	PositionCode *uint64 `json:"PositionCode,omitnil,omitempty" name:"PositionCode"`

	// 设备状态：POWER_ON 已开电 POWER_OFF 未开电 RACK_OFF 未上架 MOVING 搬迁中
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 设备最近一次的开电时间，YYYY-MM-DD 格式。
	PowerOnTime *string `json:"PowerOnTime,omitnil,omitempty" name:"PowerOnTime"`

	// 设备最近一次的上架时间，YYYY-MM-DD 格式。
	OnshelfDate *string `json:"OnshelfDate,omitnil,omitempty" name:"OnshelfDate"`

	// 设备类型 server 服务器，netDevice 网络设备
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 厂商
	Manufacturer *string `json:"Manufacturer,omitnil,omitempty" name:"Manufacturer"`

	// 其他设备-设备子类型
	TypeName *string `json:"TypeName,omitnil,omitempty" name:"TypeName"`

	// 硬件备注
	HardwareMemo *string `json:"HardwareMemo,omitnil,omitempty" name:"HardwareMemo"`
}

type DeviceHistory struct {
	// 设备sn
	Sn *string `json:"Sn,omitnil,omitempty" name:"Sn"`

	// 设备类型
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 机架名
	RackName *string `json:"RackName,omitnil,omitempty" name:"RackName"`

	// 机位号
	PositionCode *uint64 `json:"PositionCode,omitnil,omitempty" name:"PositionCode"`

	// 机房id
	IdcId *uint64 `json:"IdcId,omitnil,omitempty" name:"IdcId"`

	// 机房名称
	IdcName *string `json:"IdcName,omitnil,omitempty" name:"IdcName"`

	// 机房管理单元id
	IdcUnitId *uint64 `json:"IdcUnitId,omitnil,omitempty" name:"IdcUnitId"`

	// 机房管理单元名称
	IdcUnitName *string `json:"IdcUnitName,omitnil,omitempty" name:"IdcUnitName"`

	// 固资号
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 设备型号-版本，只有收货单详情返回
	ModelVersion *string `json:"ModelVersion,omitnil,omitempty" name:"ModelVersion"`

	// 设备高度，只有收货单详情返回
	DeviceHeight *string `json:"DeviceHeight,omitnil,omitempty" name:"DeviceHeight"`

	// 需要万兆机位，只有收货单详情返回
	Need10GbSlot *string `json:"Need10GbSlot,omitnil,omitempty" name:"Need10GbSlot"`

	// 需要直流电，只有收货单详情返回
	NeedDCPower *string `json:"NeedDCPower,omitnil,omitempty" name:"NeedDCPower"`

	// 需要外网，只有收货单详情返回
	NeedExtranet *string `json:"NeedExtranet,omitnil,omitempty" name:"NeedExtranet"`

	// 需要虚拟化，只有收货单详情返回
	NeedVirtualization *string `json:"NeedVirtualization,omitnil,omitempty" name:"NeedVirtualization"`

	// 厂商,只有收货单详情返回
	Manufacturer *string `json:"Manufacturer,omitnil,omitempty" name:"Manufacturer"`

	// 硬件备注
	HardwareMemo *string `json:"HardwareMemo,omitnil,omitempty" name:"HardwareMemo"`

	// 目标机架
	DstRackName *string `json:"DstRackName,omitnil,omitempty" name:"DstRackName"`

	// 目标机位
	DstPositionCode *string `json:"DstPositionCode,omitnil,omitempty" name:"DstPositionCode"`

	// 目标ip
	DstIp *string `json:"DstIp,omitnil,omitempty" name:"DstIp"`

	// 设备子类型, 其他设备/线材用
	TypeName *string `json:"TypeName,omitnil,omitempty" name:"TypeName"`

	// 线材-数量，只有收货单详情返回
	Quantity *uint64 `json:"Quantity,omitnil,omitempty" name:"Quantity"`

	// 计量单位，，只有收货单详情返回，'箱', '卷', '套'
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

	// 线材-收货凭证号，只有收货单详情返回
	ReceiptNumber *string `json:"ReceiptNumber,omitnil,omitempty" name:"ReceiptNumber"`
}

type DeviceOrderBaseInfo struct {
	// 机房id
	IdcId *uint64 `json:"IdcId,omitnil,omitempty" name:"IdcId"`

	// 机房名称
	IdcName *string `json:"IdcName,omitnil,omitempty" name:"IdcName"`

	// 设备类型
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 1.收货-仅核对外包装完整和数量，不开箱 2.验收-开箱核对型号SN一致
	ReceivingOperation *string `json:"ReceivingOperation,omitnil,omitempty" name:"ReceivingOperation"`

	// 设备收货-进入时间
	EntryTime *string `json:"EntryTime,omitnil,omitempty" name:"EntryTime"`

	// 设备收货-是否快递寄件
	IsExpressDelivery *bool `json:"IsExpressDelivery,omitnil,omitempty" name:"IsExpressDelivery"`

	// 设备收货-快递寄件信息
	ExpressInfo *ExpressDelivery `json:"ExpressInfo,omitnil,omitempty" name:"ExpressInfo"`

	// 上/下架人员 1.自行解决 2.由腾讯IDC负责
	StuffOption *string `json:"StuffOption,omitnil,omitempty" name:"StuffOption"`

	// 上/下架自行解决信息
	SelfOperationInfo *SelfOperation `json:"SelfOperationInfo,omitnil,omitempty" name:"SelfOperationInfo"`

	// 上架后开电
	WithPowerOn *bool `json:"WithPowerOn,omitnil,omitempty" name:"WithPowerOn"`

	// 关电确认 1.授权时关电 2.关电前需要确认
	IsPowerOffConfirm *string `json:"IsPowerOffConfirm,omitnil,omitempty" name:"IsPowerOffConfirm"`

	// 关电前需要确认信息
	PowerOffConfirmInfo *PowerOffConfirm `json:"PowerOffConfirmInfo,omitnil,omitempty" name:"PowerOffConfirmInfo"`

	// 退出交接方式 1.物流上门收货 2.客户上门自提
	HandoverMethod *string `json:"HandoverMethod,omitnil,omitempty" name:"HandoverMethod"`

	// 客户上门自提信息
	CustomerReceipt *CustomerReceipt `json:"CustomerReceipt,omitnil,omitempty" name:"CustomerReceipt"`

	// 物流上门收货信息
	LogisticsReceipt *LogisticsReceipt `json:"LogisticsReceipt,omitnil,omitempty" name:"LogisticsReceipt"`
}

type DevicePosition struct {
	// 设备SN
	Sn *string `json:"Sn,omitnil,omitempty" name:"Sn"`

	// 机架名称
	RackName *string `json:"RackName,omitnil,omitempty" name:"RackName"`

	// 机房管理单元ID
	IdcUnitId *uint64 `json:"IdcUnitId,omitnil,omitempty" name:"IdcUnitId"`

	// 机房名称
	IdcName *string `json:"IdcName,omitnil,omitempty" name:"IdcName"`

	// 机房管理单元名称
	IdcUnitName *string `json:"IdcUnitName,omitnil,omitempty" name:"IdcUnitName"`

	// 设备固资。只有服务器设备才需要这个值
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 机位编号，只有服务器设备才需要传这个值
	PositionCode *uint64 `json:"PositionCode,omitnil,omitempty" name:"PositionCode"`

	// server 代表服务器，netDevice 代表网络设备
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`
}

type DeviceRackOn struct {
	// 设备sn
	DeviceSn *string `json:"DeviceSn,omitnil,omitempty" name:"DeviceSn"`

	// 目标机架
	DstRackName *string `json:"DstRackName,omitnil,omitempty" name:"DstRackName"`

	// 目标机位,服务器必传,网络设备不用传
	DstPositionCode *string `json:"DstPositionCode,omitnil,omitempty" name:"DstPositionCode"`

	// 设备ip
	DstIp *string `json:"DstIp,omitnil,omitempty" name:"DstIp"`
}

type Distribution struct {
	// 机架编号
	RackNumber *string `json:"RackNumber,omitnil,omitempty" name:"RackNumber"`

	// 机架的用量分布
	RackUsageSet []*RackUsage `json:"RackUsageSet,omitnil,omitempty" name:"RackUsageSet"`
}

type ExpressDelivery struct {
	// 物流公司
	LogisticsCompany *string `json:"LogisticsCompany,omitnil,omitempty" name:"LogisticsCompany"`

	// 快递单号
	ExpressNumber *string `json:"ExpressNumber,omitnil,omitempty" name:"ExpressNumber"`
}

type Filter struct {
	// 需要过滤的字段。	
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 字段的过滤值。	
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type Idc struct {
	// 机房名称
	IdcName *string `json:"IdcName,omitnil,omitempty" name:"IdcName"`

	// 机房ID
	IdcId *uint64 `json:"IdcId,omitnil,omitempty" name:"IdcId"`

	// 机房管理单元详情列表
	IdcUnitSet []*IdcUnit `json:"IdcUnitSet,omitnil,omitempty" name:"IdcUnitSet"`
}

type IdcUnit struct {
	// 机房管理单元 ID
	IdcUnitId *uint64 `json:"IdcUnitId,omitnil,omitempty" name:"IdcUnitId"`

	// 机房管理单元名称
	IdcUnitName *string `json:"IdcUnitName,omitnil,omitempty" name:"IdcUnitName"`

	// 围笼列表
	CageSet []*Cage `json:"CageSet,omitnil,omitempty" name:"CageSet"`
}

type IdcUnitInfo struct {
	// 机房管理单元地址
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// 机房经理
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 联系电话
	TelNumber *string `json:"TelNumber,omitnil,omitempty" name:"TelNumber"`

	// 资产管理员
	AssetManager *string `json:"AssetManager,omitnil,omitempty" name:"AssetManager"`

	// 资产管理员电话
	AssetManagerTelNumber *string `json:"AssetManagerTelNumber,omitnil,omitempty" name:"AssetManagerTelNumber"`
}

type LogisticsReceipt struct {
	// 物流预计上门时间
	LogisticsArrivalTime *string `json:"LogisticsArrivalTime,omitnil,omitempty" name:"LogisticsArrivalTime"`

	// 物流公司名称
	LogisticsCompany *string `json:"LogisticsCompany,omitnil,omitempty" name:"LogisticsCompany"`

	// 物流联系人
	LogisticsStuff *string `json:"LogisticsStuff,omitnil,omitempty" name:"LogisticsStuff"`

	// 物流电话
	LogisticsStuffContact *string `json:"LogisticsStuffContact,omitnil,omitempty" name:"LogisticsStuffContact"`

	// 收货人电话
	ReceiverContact *string `json:"ReceiverContact,omitnil,omitempty" name:"ReceiverContact"`

	// 收货人姓名
	ReceiverName *string `json:"ReceiverName,omitnil,omitempty" name:"ReceiverName"`

	// 收货地址
	ShippingAddress *string `json:"ShippingAddress,omitnil,omitempty" name:"ShippingAddress"`
}

type ModelEvaluationBaseInfo struct {
	// 客户名称
	CustomerName *string `json:"CustomerName,omitnil,omitempty" name:"CustomerName"`

	// server 服务器  netDevice 网络设备
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 园区名称
	CampusName *string `json:"CampusName,omitnil,omitempty" name:"CampusName"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type ModelVersion struct {
	// 型号名称
	DevModel *string `json:"DevModel,omitnil,omitempty" name:"DevModel"`

	// 版本
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`
}

type ModelVersionCount struct {
	// 型号名称
	DevModel *string `json:"DevModel,omitnil,omitempty" name:"DevModel"`

	// 版本数量
	VersionCount *uint64 `json:"VersionCount,omitnil,omitempty" name:"VersionCount"`
}

type ModelVersionDetail struct {
	// 版本号
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 0 代表在当前园区没评估过，1 代表完全满足IDC准入标准 2 代表存在托管风险 3 代表不满足IDC准入标准
	CheckResult *uint64 `json:"CheckResult,omitnil,omitempty" name:"CheckResult"`

	// 型号各个配置项的详情
	OptionSet []*TemplateOption `json:"OptionSet,omitnil,omitempty" name:"OptionSet"`

	// 设备型号名称及版本
	ModelVersion *string `json:"ModelVersion,omitnil,omitempty" name:"ModelVersion"`
}

// Predefined struct for user
type ModifyWorkOrderTypeCollectFlagRequestParams struct {
	// 工单类型的唯一英文标识
	WorkOrderType *string `json:"WorkOrderType,omitnil,omitempty" name:"WorkOrderType"`
}

type ModifyWorkOrderTypeCollectFlagRequest struct {
	*tchttp.BaseRequest
	
	// 工单类型的唯一英文标识
	WorkOrderType *string `json:"WorkOrderType,omitnil,omitempty" name:"WorkOrderType"`
}

func (r *ModifyWorkOrderTypeCollectFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWorkOrderTypeCollectFlagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkOrderType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyWorkOrderTypeCollectFlagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWorkOrderTypeCollectFlagResponseParams struct {
	// 工单类型当前的收藏状态
	CurrentCollectFlag *bool `json:"CurrentCollectFlag,omitnil,omitempty" name:"CurrentCollectFlag"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyWorkOrderTypeCollectFlagResponse struct {
	*tchttp.BaseResponse
	Response *ModifyWorkOrderTypeCollectFlagResponseParams `json:"Response"`
}

func (r *ModifyWorkOrderTypeCollectFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWorkOrderTypeCollectFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetDeviceModel struct {
	// 版本号
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 型号和版本的组合名称
	ModelVersion *string `json:"ModelVersion,omitnil,omitempty" name:"ModelVersion"`

	// 型号名
	DevModel *string `json:"DevModel,omitnil,omitempty" name:"DevModel"`

	// 宽度
	DevWidth *string `json:"DevWidth,omitnil,omitempty" name:"DevWidth"`

	// 深度
	DevDepth *string `json:"DevDepth,omitnil,omitempty" name:"DevDepth"`

	// 重量
	DevWeight *string `json:"DevWeight,omitnil,omitempty" name:"DevWeight"`

	// 是否携带挂耳
	MountEar *string `json:"MountEar,omitnil,omitempty" name:"MountEar"`

	// 是否符合CCC认证
	AccordCCC *string `json:"AccordCCC,omitnil,omitempty" name:"AccordCCC"`

	// 是否通过入网许可认证
	PassNetwork *string `json:"PassNetwork,omitnil,omitempty" name:"PassNetwork"`

	// 电源接口型号
	PowerportType *string `json:"PowerportType,omitnil,omitempty" name:"PowerportType"`

	// 电源模块
	PowerModule *string `json:"PowerModule,omitnil,omitempty" name:"PowerModule"`

	// 电源模块数量
	PowermoduleNum *string `json:"PowermoduleNum,omitnil,omitempty" name:"PowermoduleNum"`

	// 电源模块位置
	PowermodulePosition *string `json:"PowermodulePosition,omitnil,omitempty" name:"PowermodulePosition"`

	// 高压直流自适应
	HighVoltageAdapt *string `json:"HighVoltageAdapt,omitnil,omitempty" name:"HighVoltageAdapt"`

	// 实际工作功耗(W)
	PowerEnergy *string `json:"PowerEnergy,omitnil,omitempty" name:"PowerEnergy"`

	// 进风口位置
	InwindPosition *string `json:"InwindPosition,omitnil,omitempty" name:"InwindPosition"`

	// 出风口位置
	OutwindPosition *string `json:"OutwindPosition,omitnil,omitempty" name:"OutwindPosition"`

	// 业务端口位置
	BusinessPortPosition *string `json:"BusinessPortPosition,omitnil,omitempty" name:"BusinessPortPosition"`

	// 带有理线器
	LineManager *string `json:"LineManager,omitnil,omitempty" name:"LineManager"`

	// 0 代表在当前园区没评估过，1 代表完全满足IDC准入标准  2 代表存在托管风险  3 代表不满足IDC准入标准
	CheckResult *uint64 `json:"CheckResult,omitnil,omitempty" name:"CheckResult"`

	// 设备高度
	DevHeight *string `json:"DevHeight,omitnil,omitempty" name:"DevHeight"`
}

type NetReceivingInfo struct {
	// 设备sn
	DeviceSn *string `json:"DeviceSn,omitnil,omitempty" name:"DeviceSn"`

	// 设备型号-版本
	ModelVersion *string `json:"ModelVersion,omitnil,omitempty" name:"ModelVersion"`

	// 硬件备注
	HardwareMemo *string `json:"HardwareMemo,omitnil,omitempty" name:"HardwareMemo"`

	// 厂商
	Manufacturer *string `json:"Manufacturer,omitnil,omitempty" name:"Manufacturer"`
}

type OptionValueItem struct {
	// 选项的值
	OptionValue *string `json:"OptionValue,omitnil,omitempty" name:"OptionValue"`

	// 是否默认选中
	Selected *bool `json:"Selected,omitnil,omitempty" name:"Selected"`
}

type OrderStep struct {
	// 步骤名
	StepName *string `json:"StepName,omitnil,omitempty" name:"StepName"`

	// 处理人
	OwnerName *string `json:"OwnerName,omitnil,omitempty" name:"OwnerName"`

	// 完成时间
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`

	// 此步骤状态
	StepStatus *string `json:"StepStatus,omitnil,omitempty" name:"StepStatus"`
}

type OtherDevReceivingInfo struct {
	// 设备sn
	DeviceSn *string `json:"DeviceSn,omitnil,omitempty" name:"DeviceSn"`

	// 'USB', '移动硬盘', '网络设备板卡', '网络设备模块', '服务器硬盘', '服务器内存', '其他'
	TypeName *string `json:"TypeName,omitnil,omitempty" name:"TypeName"`

	// 厂商
	Manufacturer *string `json:"Manufacturer,omitnil,omitempty" name:"Manufacturer"`

	// 硬件备注
	HardwareMemo *string `json:"HardwareMemo,omitnil,omitempty" name:"HardwareMemo"`
}

type Personnel struct {
	// 证件号码
	IDCardNumber *string `json:"IDCardNumber,omitnil,omitempty" name:"IDCardNumber"`

	// 证件类型。对应关系如下：IDENTITY_CARD: 身份证,
	// HONG_KONG_AND_MACAO_PASS: 港澳通行证',
	// PASSPORT: 护照,
	// DRIVING_LICENSE: 驾照,
	// OTHER: 其他
	IDCardType *string `json:"IDCardType,omitnil,omitempty" name:"IDCardType"`

	// 公司名称
	Company *string `json:"Company,omitnil,omitempty" name:"Company"`

	// 语言。对应关系：ENGLISH: 英文, CHINESE: 中文
	LanguageType *string `json:"LanguageType,omitnil,omitempty" name:"LanguageType"`

	// 姓名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 电话
	TelNumber *string `json:"TelNumber,omitnil,omitempty" name:"TelNumber"`

	// 职位
	Position *string `json:"Position,omitnil,omitempty" name:"Position"`

	// 微信
	Wechat *string `json:"Wechat,omitnil,omitempty" name:"Wechat"`

	// 邮箱
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`
}

type PersonnelVisitBaseInfo struct {
	// 机房名称
	IdcName *string `json:"IdcName,omitnil,omitempty" name:"IdcName"`

	// 到访原因。到访原因，映射关系：DEVICE_MAINTENANCE 设备维护 DEVICE_MOVE 设备收货上下架 CHECK 盘点 OTHER 其他
	VisitReason []*string `json:"VisitReason,omitnil,omitempty" name:"VisitReason"`

	// 到访原因
	VisitRemark *string `json:"VisitRemark,omitnil,omitempty" name:"VisitRemark"`

	// 到访结束时间
	EnterStartTime *string `json:"EnterStartTime,omitnil,omitempty" name:"EnterStartTime"`

	// 到访开始时间
	EnterEndTime *string `json:"EnterEndTime,omitnil,omitempty" name:"EnterEndTime"`

	// 机房管理单元列表
	IdcUnitNameList []*string `json:"IdcUnitNameList,omitnil,omitempty" name:"IdcUnitNameList"`
}

type Position struct {
	// 机位ID
	PositionId *uint64 `json:"PositionId,omitnil,omitempty" name:"PositionId"`

	// 机位高度
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 机位编号
	PositionCode *string `json:"PositionCode,omitnil,omitempty" name:"PositionCode"`

	// 机位状态,0 空闲,1 已用,2 不可用,3 预占用,4 预留
	PositionStatus *uint64 `json:"PositionStatus,omitnil,omitempty" name:"PositionStatus"`

	// 设备规划类型ID
	PlanDeviceType *int64 `json:"PlanDeviceType,omitnil,omitempty" name:"PlanDeviceType"`

	// 机位所属的机房管理单元ID
	IdcUnitId *uint64 `json:"IdcUnitId,omitnil,omitempty" name:"IdcUnitId"`

	// 机位所属的机架ID
	RackId *uint64 `json:"RackId,omitnil,omitempty" name:"RackId"`

	// 机位所属的机架名称
	RackName *string `json:"RackName,omitnil,omitempty" name:"RackName"`

	// 机位所属的机房管理单元名称
	IdcUnitName *string `json:"IdcUnitName,omitnil,omitempty" name:"IdcUnitName"`

	// 机位所属的机房名称
	IdcName *string `json:"IdcName,omitnil,omitempty" name:"IdcName"`

	// 机位所属的机房ID
	IdcId *uint64 `json:"IdcId,omitnil,omitempty" name:"IdcId"`

	// 机位上如果有设备，该字段代表设备的 SN 码，如果是空闲机位，不返回该字段。
	Sn *string `json:"Sn,omitnil,omitempty" name:"Sn"`

	// 机位上如果有设备，该字段代表设备的固资号，如果是空闲机位，不返回该字段。
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 机位上如果有设备，该字段代表设备的设备型号加版本号，如果是空闲机位，不返回该字段。
	ModelVersion *string `json:"ModelVersion,omitnil,omitempty" name:"ModelVersion"`
}

type PositionStatusItem struct {
	// 状态值
	PositionStatus *uint64 `json:"PositionStatus,omitnil,omitempty" name:"PositionStatus"`

	// 对应的机位数量
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type PowerOffConfirm struct {
	// 联系人
	ConfirmContact *string `json:"ConfirmContact,omitnil,omitempty" name:"ConfirmContact"`

	// 联系人电话
	ConfirmContactNumber *string `json:"ConfirmContactNumber,omitnil,omitempty" name:"ConfirmContactNumber"`
}

type Rack struct {
	// 机架名称
	RackName *string `json:"RackName,omitnil,omitempty" name:"RackName"`

	// 机架所属的机房管理单元ID
	IdcUnitId *uint64 `json:"IdcUnitId,omitnil,omitempty" name:"IdcUnitId"`

	// 机架所属的机房管理单元名称
	IdcUnitName *string `json:"IdcUnitName,omitnil,omitempty" name:"IdcUnitName"`

	// 机架所属的机房名称
	IdcName *string `json:"IdcName,omitnil,omitempty" name:"IdcName"`

	// 机架所属的机房ID
	IdcId *uint64 `json:"IdcId,omitnil,omitempty" name:"IdcId"`

	// 机架ID
	RackId *uint64 `json:"RackId,omitnil,omitempty" name:"RackId"`

	// 是否开电
	IsPowerOn *bool `json:"IsPowerOn,omitnil,omitempty" name:"IsPowerOn"`

	// 机架最近一次开电时间，YYYY-MM-DD 格式
	RackOpenTime *string `json:"RackOpenTime,omitnil,omitempty" name:"RackOpenTime"`

	// 托管类型
	HostingType *string `json:"HostingType,omitnil,omitempty" name:"HostingType"`
}

type RackUsage struct {
	// 机架ID
	RackId *uint64 `json:"RackId,omitnil,omitempty" name:"RackId"`

	// 已使用的机位数量
	UsedNum *uint64 `json:"UsedNum,omitnil,omitempty" name:"UsedNum"`

	// 空闲机位数量
	UnusedNum *uint64 `json:"UnusedNum,omitnil,omitempty" name:"UnusedNum"`

	// 机架简称
	RackShortName *string `json:"RackShortName,omitnil,omitempty" name:"RackShortName"`

	// 机位总数
	TotalNum *uint64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// 机位使用率
	UsedRate *float64 `json:"UsedRate,omitnil,omitempty" name:"UsedRate"`
}

type SelfOperation struct {
	// 联系人员电话
	StuffContact *string `json:"StuffContact,omitnil,omitempty" name:"StuffContact"`

	// 身份证号
	StuffIDCard *string `json:"StuffIDCard,omitnil,omitempty" name:"StuffIDCard"`

	// 人员姓名
	StuffName *string `json:"StuffName,omitnil,omitempty" name:"StuffName"`

	// 上门时间
	OperationTime *string `json:"OperationTime,omitnil,omitempty" name:"OperationTime"`
}

type ServerModel struct {
	// 型号名称
	DevModel *string `json:"DevModel,omitnil,omitempty" name:"DevModel"`

	// 节点数
	DevNode *string `json:"DevNode,omitnil,omitempty" name:"DevNode"`

	// 设备高度
	DevHeight *string `json:"DevHeight,omitnil,omitempty" name:"DevHeight"`

	// 功耗
	PowerEnergy *string `json:"PowerEnergy,omitnil,omitempty" name:"PowerEnergy"`

	// 电源接口型号
	PowerportType *string `json:"PowerportType,omitnil,omitempty" name:"PowerportType"`

	// 电源模块数量
	PowermoduleNum *string `json:"PowermoduleNum,omitnil,omitempty" name:"PowermoduleNum"`

	// 进风口位置
	InwindPosition *string `json:"InwindPosition,omitnil,omitempty" name:"InwindPosition"`

	// 出风口位置
	OutwindPosition *string `json:"OutwindPosition,omitnil,omitempty" name:"OutwindPosition"`

	// 网卡接口位置
	NetportPosition *string `json:"NetportPosition,omitnil,omitempty" name:"NetportPosition"`

	// 宽度
	DevWidth *string `json:"DevWidth,omitnil,omitempty" name:"DevWidth"`

	// 深度
	DevDepth *string `json:"DevDepth,omitnil,omitempty" name:"DevDepth"`

	// 重量
	DevWeight *string `json:"DevWeight,omitnil,omitempty" name:"DevWeight"`

	// 电源模块
	PowerModule *string `json:"PowerModule,omitnil,omitempty" name:"PowerModule"`

	// 电源模块位置
	PowermodulePosition *string `json:"PowermodulePosition,omitnil,omitempty" name:"PowermodulePosition"`

	// 网络接口类型
	NetportType *string `json:"NetportType,omitnil,omitempty" name:"NetportType"`

	// 网卡速率
	NetSpeed *string `json:"NetSpeed,omitnil,omitempty" name:"NetSpeed"`

	// 0 代表在当前园区没评估过，1 代表完全满足IDC准入标准 2 代表存在托管风险 3 代表不满足IDC准入标准
	CheckResult *uint64 `json:"CheckResult,omitnil,omitempty" name:"CheckResult"`

	// 版本号
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 型号和版本的组合名称
	ModelVersion *string `json:"ModelVersion,omitnil,omitempty" name:"ModelVersion"`
}

type ServerReceivingInfo struct {
	// 设备sn
	DeviceSn *string `json:"DeviceSn,omitnil,omitempty" name:"DeviceSn"`

	// 设备型号-版本
	ModelVersion *string `json:"ModelVersion,omitnil,omitempty" name:"ModelVersion"`

	// 需要万兆机位
	Need10GbSlot *string `json:"Need10GbSlot,omitnil,omitempty" name:"Need10GbSlot"`

	// 需要直流电
	NeedDCPower *string `json:"NeedDCPower,omitnil,omitempty" name:"NeedDCPower"`

	// 需要外网
	NeedExtranet *string `json:"NeedExtranet,omitnil,omitempty" name:"NeedExtranet"`

	// 需要虚拟化
	NeedVirtualization *string `json:"NeedVirtualization,omitnil,omitempty" name:"NeedVirtualization"`

	// 硬件备注
	HardwareMemo *string `json:"HardwareMemo,omitnil,omitempty" name:"HardwareMemo"`
}

type TemplateOption struct {
	// 选项名称
	OptionName *string `json:"OptionName,omitnil,omitempty" name:"OptionName"`

	// 腾讯的标准
	Standard *string `json:"Standard,omitnil,omitempty" name:"Standard"`

	// 腾讯标准的展示字段
	StandardInfo *string `json:"StandardInfo,omitnil,omitempty" name:"StandardInfo"`

	// 选项的唯一标识
	OptionKey *string `json:"OptionKey,omitnil,omitempty" name:"OptionKey"`

	// 输入的类型
	InputType *string `json:"InputType,omitnil,omitempty" name:"InputType"`

	// 输入值的类型
	ValueType *string `json:"ValueType,omitnil,omitempty" name:"ValueType"`

	// 是否符合腾讯标准的比较方式，-- 为不做比较
	CompareType *string `json:"CompareType,omitnil,omitempty" name:"CompareType"`

	// 下拉列表中填充的选项值
	OptionValueSet []*OptionValueItem `json:"OptionValueSet,omitnil,omitempty" name:"OptionValueSet"`

	// 输入框中的占位的提示符
	InputHint *string `json:"InputHint,omitnil,omitempty" name:"InputHint"`

	// 输入框下方的提示文案
	InputInfo *string `json:"InputInfo,omitnil,omitempty" name:"InputInfo"`

	// 客户写入的值
	OptionValue *string `json:"OptionValue,omitnil,omitempty" name:"OptionValue"`
}

type WireReceivingInfo struct {
	// '光纤', '网线', '电源线'
	TypeName *string `json:"TypeName,omitnil,omitempty" name:"TypeName"`

	// 数量
	Quantity *uint64 `json:"Quantity,omitnil,omitempty" name:"Quantity"`

	// '箱', '卷', '套'
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

	// 收货凭证号
	ReceiptNumber *string `json:"ReceiptNumber,omitnil,omitempty" name:"ReceiptNumber"`

	// 硬件备注
	HardwareMemo *string `json:"HardwareMemo,omitnil,omitempty" name:"HardwareMemo"`
}

type WorkOrderData struct {
	// 工单号
	WorkOrderId *string `json:"WorkOrderId,omitnil,omitempty" name:"WorkOrderId"`

	// 服务类型，一个服务可能会产生多个工单
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// 工单类型
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// 工单状态
	OrderStatus *string `json:"OrderStatus,omitnil,omitempty" name:"OrderStatus"`

	// 工单创建人
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 工单创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 工单完成时间
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`
}

type WorkOrderFamilyDetail struct {
	// 工单类型大类的名称
	WorkOrderFamily *string `json:"WorkOrderFamily,omitnil,omitempty" name:"WorkOrderFamily"`

	// 工单类型详情列表
	WorkOrderTypeSet []*WorkOrderTypeDetail `json:"WorkOrderTypeSet,omitnil,omitempty" name:"WorkOrderTypeSet"`
}

type WorkOrderTinyInfo struct {
	// 工单id
	WorkOrderId *string `json:"WorkOrderId,omitnil,omitempty" name:"WorkOrderId"`

	// 服务类型，一个服务可能会产生多个工单
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// 工单类型
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`
}

type WorkOrderTypeDetail struct {
	// 工单类型所属的大类
	WorkOrderFamily *string `json:"WorkOrderFamily,omitnil,omitempty" name:"WorkOrderFamily"`

	// 工单类型名称
	WorkOrderName *string `json:"WorkOrderName,omitnil,omitempty" name:"WorkOrderName"`

	// 工单类型的唯一英文标识
	WorkOrderType *string `json:"WorkOrderType,omitnil,omitempty" name:"WorkOrderType"`

	// 工单类型简述
	WorkOrderDescription *string `json:"WorkOrderDescription,omitnil,omitempty" name:"WorkOrderDescription"`

	// 是否被收藏
	CollectFlag *bool `json:"CollectFlag,omitnil,omitempty" name:"CollectFlag"`

	// 服务时效
	SlaMessage *string `json:"SlaMessage,omitnil,omitempty" name:"SlaMessage"`
}