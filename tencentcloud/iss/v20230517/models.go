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

package v20230517

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AIConfig struct {
	// AI 分析类型。可选值为 Facemask(口罩识别)、Chefhat(厨师帽识别)、Smoking(抽烟检测)、Chefcloth(厨师服识别)、PhoneCall(接打电话识别)、Pet(宠物识别)、Body(人体识别)和Car(车辆车牌识别)等
	DetectType *string `json:"DetectType,omitnil" name:"DetectType"`

	// 截图频率。可选值1～20秒
	TimeInterval *uint64 `json:"TimeInterval,omitnil" name:"TimeInterval"`

	// 模板生效的时间段。最多包含5组时间段
	OperTimeSlot []*OperTimeSlot `json:"OperTimeSlot,omitnil" name:"OperTimeSlot"`
}

type AITaskInfo struct {
	// AI 任务 ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// AI 任务名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// AI 任务描述
	Desc *string `json:"Desc,omitnil" name:"Desc"`

	// AI 任务状态。"on"代表开启了 AI 分析任务，"off"代表停止 AI 分析任务
	Status *string `json:"Status,omitnil" name:"Status"`

	// 通道 ID 列表
	ChannelList []*string `json:"ChannelList,omitnil" name:"ChannelList"`

	// AI 结果回调地址
	CallbackUrl *string `json:"CallbackUrl,omitnil" name:"CallbackUrl"`

	// AI 配置列表
	Templates []*AITemplates `json:"Templates,omitnil" name:"Templates"`

	// 创建时间
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// 更新时间
	UpdatedTime *string `json:"UpdatedTime,omitnil" name:"UpdatedTime"`
}

type AITaskResultData struct {
	// AI 任务 ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 在 BeginTime 和 EndTime 时间之内，有识别结果的 AI 调用次数（分页依据此数值）
	AIResultCount *uint64 `json:"AIResultCount,omitnil" name:"AIResultCount"`

	// AI 任务执行结果详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	AIResults *AITaskResultInfo `json:"AIResults,omitnil" name:"AIResults"`
}

type AITaskResultInfo struct {
	// 人体识别结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Body []*BodyAIResultInfo `json:"Body,omitnil" name:"Body"`

	// 宠物识别结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pet []*PetAIResultInfo `json:"Pet,omitnil" name:"Pet"`

	// 车辆车牌识别结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Car []*CarAIResultInfo `json:"Car,omitnil" name:"Car"`

	// 厨师帽结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChefHat []*ChefHatAIResultInfo `json:"ChefHat,omitnil" name:"ChefHat"`

	// 厨师服结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChefCloth []*ChefClothAIResultInfo `json:"ChefCloth,omitnil" name:"ChefCloth"`

	// 口罩识别结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	FaceMask []*FaceMaskAIResultInfo `json:"FaceMask,omitnil" name:"FaceMask"`

	// 抽烟检测结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Smoking []*SmokingAIResultInfo `json:"Smoking,omitnil" name:"Smoking"`

	// 接打电话识别结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneCall []*PhoneCallAIResultInfo `json:"PhoneCall,omitnil" name:"PhoneCall"`
}

type AITemplates struct {
	// AI 类别。可选值 AI(AI 分析)和 Snapshot(截图)，Templates 列表中只能出现一种类型。
	Tag *string `json:"Tag,omitnil" name:"Tag"`

	// AI 分析配置。和"SnapshotConfig"二选一。
	AIConfig *AIConfig `json:"AIConfig,omitnil" name:"AIConfig"`

	// 截图配置。和"AIConfig"二选一。
	SnapshotConfig *SnapshotConfig `json:"SnapshotConfig,omitnil" name:"SnapshotConfig"`
}

// Predefined struct for user
type AddAITaskRequestParams struct {
	// AI 任务名称。仅支持中文、英文、数字、_、-，长度不超过32个字符
	Name *string `json:"Name,omitnil" name:"Name"`

	// 通道 ID 列表。不能添加存在于其他 AI 任务的通道，限制1000个通道。
	ChannelList []*string `json:"ChannelList,omitnil" name:"ChannelList"`

	// AI 配置列表
	Templates []*AITemplates `json:"Templates,omitnil" name:"Templates"`

	// AI 任务描述。仅支持中文、英文、数字、_、-，长度不超过128个字符
	Desc *string `json:"Desc,omitnil" name:"Desc"`

	// AI 结果回调地址。类似 "http://ip:port/xxx或者https://domain/xxx
	CallbackUrl *string `json:"CallbackUrl,omitnil" name:"CallbackUrl"`

	// 是否立即开启 AI 任务。"true"代表立即开启 AI 任务，"false"代表暂不开启 AI 任务，默认为 false。
	IsStartTheTask *bool `json:"IsStartTheTask,omitnil" name:"IsStartTheTask"`
}

type AddAITaskRequest struct {
	*tchttp.BaseRequest
	
	// AI 任务名称。仅支持中文、英文、数字、_、-，长度不超过32个字符
	Name *string `json:"Name,omitnil" name:"Name"`

	// 通道 ID 列表。不能添加存在于其他 AI 任务的通道，限制1000个通道。
	ChannelList []*string `json:"ChannelList,omitnil" name:"ChannelList"`

	// AI 配置列表
	Templates []*AITemplates `json:"Templates,omitnil" name:"Templates"`

	// AI 任务描述。仅支持中文、英文、数字、_、-，长度不超过128个字符
	Desc *string `json:"Desc,omitnil" name:"Desc"`

	// AI 结果回调地址。类似 "http://ip:port/xxx或者https://domain/xxx
	CallbackUrl *string `json:"CallbackUrl,omitnil" name:"CallbackUrl"`

	// 是否立即开启 AI 任务。"true"代表立即开启 AI 任务，"false"代表暂不开启 AI 任务，默认为 false。
	IsStartTheTask *bool `json:"IsStartTheTask,omitnil" name:"IsStartTheTask"`
}

func (r *AddAITaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddAITaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "ChannelList")
	delete(f, "Templates")
	delete(f, "Desc")
	delete(f, "CallbackUrl")
	delete(f, "IsStartTheTask")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddAITaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddAITaskResponseParams struct {
	// AI任务信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *AITaskInfo `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AddAITaskResponse struct {
	*tchttp.BaseResponse
	Response *AddAITaskResponseParams `json:"Response"`
}

func (r *AddAITaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddAITaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddDeviceData struct {
	// 设备iD
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`

	// 设备编码（即我们为设备生成的20位国标编码）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *string `json:"Code,omitnil" name:"Code"`

	// 设备名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 设备接入协议，1:RTMP,2:GB,3:GW 
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessProtocol *int64 `json:"AccessProtocol,omitnil" name:"AccessProtocol"`

	// 设备类型，1:IPC,2:NVR
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *int64 `json:"Type,omitnil" name:"Type"`

	// 设备接入服务节点ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 设备接入服务节点名称
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitnil" name:"ClusterName"`

	// 设备流传输协议，1:UDP,2:TCP 
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransportProtocol *int64 `json:"TransportProtocol,omitnil" name:"TransportProtocol"`

	// 设备密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitnil" name:"Password"`

	// 设备描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 设备状态，0:未注册,1:在线,2:离线,3:禁用
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 设备所属组织ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationId *int64 `json:"OrganizationId,omitnil" name:"OrganizationId"`

	// 设备接入网关ID，从查询网关列表接口中获取（仅网关接入需要）
	// 注意：此字段可能返回 null，表示取不到有效值。
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 网关接入协议类型，1.海康SDK，2.大华SDK，3.宇视SDK，4.Onvif（仅网关接入需要）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProtocolType *int64 `json:"ProtocolType,omitnil" name:"ProtocolType"`

	// 设备接入IP（仅网关接入需要）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 设备Port（仅网关接入需要）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitnil" name:"Port"`

	// 设备用户名（仅网关接入需要）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Username *string `json:"Username,omitnil" name:"Username"`

	// 用户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *int64 `json:"AppId,omitnil" name:"AppId"`
}

type AddOrgData struct {
	// 组织 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationId *string `json:"OrganizationId,omitnil" name:"OrganizationId"`

	// 组织名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 组织父节点 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentId *string `json:"ParentId,omitnil" name:"ParentId"`

	// 组织层级
	// 注意：此字段可能返回 null，表示取不到有效值。
	Level *int64 `json:"Level,omitnil" name:"Level"`

	// 用户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *int64 `json:"AppId,omitnil" name:"AppId"`

	// 组织结构
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentIds *string `json:"ParentIds,omitnil" name:"ParentIds"`

	// 设备总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 设备在线数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Online *int64 `json:"Online,omitnil" name:"Online"`
}

// Predefined struct for user
type AddOrganizationRequestParams struct {
	// 组织名称（仅支持中文、英文、数字、_、-的组合，长度不超过16个字符，且组织名称不能重复）
	Name *string `json:"Name,omitnil" name:"Name"`

	// 组织父节点 ID（从查询组织接口DescribeOrganization中获取，填0代表根组织）
	ParentId *string `json:"ParentId,omitnil" name:"ParentId"`
}

type AddOrganizationRequest struct {
	*tchttp.BaseRequest
	
	// 组织名称（仅支持中文、英文、数字、_、-的组合，长度不超过16个字符，且组织名称不能重复）
	Name *string `json:"Name,omitnil" name:"Name"`

	// 组织父节点 ID（从查询组织接口DescribeOrganization中获取，填0代表根组织）
	ParentId *string `json:"ParentId,omitnil" name:"ParentId"`
}

func (r *AddOrganizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddOrganizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "ParentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddOrganizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddOrganizationResponseParams struct {
	// 增加组织接口返回数据
	Data *AddOrgData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AddOrganizationResponse struct {
	*tchttp.BaseResponse
	Response *AddOrganizationResponseParams `json:"Response"`
}

func (r *AddOrganizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddRecordBackupPlanData struct {
	// 录像上云计划ID
	PlanId *string `json:"PlanId,omitnil" name:"PlanId"`

	// 录像上云计划名称
	PlanName *string `json:"PlanName,omitnil" name:"PlanName"`

	// 录像上云模板ID
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 录像上云计划描述
	Describe *string `json:"Describe,omitnil" name:"Describe"`

	// 云文件生命周期
	LifeCycle *LifeCycleData `json:"LifeCycle,omitnil" name:"LifeCycle"`

	// 录像上云计划状态，1:正常使用中，0:删除中，无法使用
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 通道数量
	ChannelCount *int64 `json:"ChannelCount,omitnil" name:"ChannelCount"`

	// 创建时间
	CreateAt *string `json:"CreateAt,omitnil" name:"CreateAt"`

	// 修改时间
	UpdateAt *string `json:"UpdateAt,omitnil" name:"UpdateAt"`
}

// Predefined struct for user
type AddRecordBackupPlanRequestParams struct {
	// 录制模板ID（录像计划关联的模板ID，从查询录像上云模板列表接口ListRecordBackupTemplates中获取）
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 录像计划名称（仅支持中文、英文、数字、_、-，长度不超过32个字符，计划名称全局唯一，不能为空，不能重复）
	PlanName *string `json:"PlanName,omitnil" name:"PlanName"`

	// 录像计划描述（仅支持中文、英文、数字、_、-，长度不超过128个字符）
	Describe *string `json:"Describe,omitnil" name:"Describe"`

	// 生命周期（录像文件生命周期设置，管理文件冷、热存储的时间）
	LifeCycle *LifeCycleData `json:"LifeCycle,omitnil" name:"LifeCycle"`

	// 通道及通道所属设备（添加录像的设备的通道信息，一次添加通道总数不超过5000个，包括组织目录下的通道数量）
	Channels []*ChannelInfo `json:"Channels,omitnil" name:"Channels"`

	// 添加组织目录下所有设备通道（Json数组，可以为空，通道总数量不超过5000个（包括Channel字段的数量））
	OrganizationId []*string `json:"OrganizationId,omitnil" name:"OrganizationId"`
}

type AddRecordBackupPlanRequest struct {
	*tchttp.BaseRequest
	
	// 录制模板ID（录像计划关联的模板ID，从查询录像上云模板列表接口ListRecordBackupTemplates中获取）
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 录像计划名称（仅支持中文、英文、数字、_、-，长度不超过32个字符，计划名称全局唯一，不能为空，不能重复）
	PlanName *string `json:"PlanName,omitnil" name:"PlanName"`

	// 录像计划描述（仅支持中文、英文、数字、_、-，长度不超过128个字符）
	Describe *string `json:"Describe,omitnil" name:"Describe"`

	// 生命周期（录像文件生命周期设置，管理文件冷、热存储的时间）
	LifeCycle *LifeCycleData `json:"LifeCycle,omitnil" name:"LifeCycle"`

	// 通道及通道所属设备（添加录像的设备的通道信息，一次添加通道总数不超过5000个，包括组织目录下的通道数量）
	Channels []*ChannelInfo `json:"Channels,omitnil" name:"Channels"`

	// 添加组织目录下所有设备通道（Json数组，可以为空，通道总数量不超过5000个（包括Channel字段的数量））
	OrganizationId []*string `json:"OrganizationId,omitnil" name:"OrganizationId"`
}

func (r *AddRecordBackupPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddRecordBackupPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "PlanName")
	delete(f, "Describe")
	delete(f, "LifeCycle")
	delete(f, "Channels")
	delete(f, "OrganizationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddRecordBackupPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddRecordBackupPlanResponseParams struct {
	// 返回数据
	Data *AddRecordBackupPlanData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AddRecordBackupPlanResponse struct {
	*tchttp.BaseResponse
	Response *AddRecordBackupPlanResponseParams `json:"Response"`
}

func (r *AddRecordBackupPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddRecordBackupPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddRecordBackupTemplateData struct {
	// 模板ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 模板名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateName *string `json:"TemplateName,omitnil" name:"TemplateName"`

	// 上云时间段（按周进行设置，支持一天设置多个时间段，每个时间段不小于10分钟）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeSections []*RecordTemplateTimeSections `json:"TimeSections,omitnil" name:"TimeSections"`

	// 录像时间段（按周进行设置，支持一天设置多个时间段，每个时间段不小于10分钟）
	// 注意：此字段可能返回 null，表示取不到有效值。
	DevTimeSections []*RecordTemplateTimeSections `json:"DevTimeSections,omitnil" name:"DevTimeSections"`

	// 上云倍速（支持1，2，4倍速）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Scale *int64 `json:"Scale,omitnil" name:"Scale"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateAt *string `json:"CreateAt,omitnil" name:"CreateAt"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateAt *string `json:"UpdateAt,omitnil" name:"UpdateAt"`
}

// Predefined struct for user
type AddRecordBackupTemplateRequestParams struct {
	// 模板名称（仅支持中文、英文、数字、_、-，长度不超过32个字符，模板名称全局唯一，不能为空，不能重复）
	TemplateName *string `json:"TemplateName,omitnil" name:"TemplateName"`

	// 上云时间段（按周进行设置，支持一天设置多个时间段，每个时间段不小于10分钟）
	TimeSections []*RecordTemplateTimeSections `json:"TimeSections,omitnil" name:"TimeSections"`

	// 录像时间段（按周进行设置，支持一天设置多个时间段，每个时间段不小于10分钟）
	DevTimeSections []*RecordTemplateTimeSections `json:"DevTimeSections,omitnil" name:"DevTimeSections"`

	// 上云倍速（支持1，2，4倍速）
	Scale *int64 `json:"Scale,omitnil" name:"Scale"`
}

type AddRecordBackupTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板名称（仅支持中文、英文、数字、_、-，长度不超过32个字符，模板名称全局唯一，不能为空，不能重复）
	TemplateName *string `json:"TemplateName,omitnil" name:"TemplateName"`

	// 上云时间段（按周进行设置，支持一天设置多个时间段，每个时间段不小于10分钟）
	TimeSections []*RecordTemplateTimeSections `json:"TimeSections,omitnil" name:"TimeSections"`

	// 录像时间段（按周进行设置，支持一天设置多个时间段，每个时间段不小于10分钟）
	DevTimeSections []*RecordTemplateTimeSections `json:"DevTimeSections,omitnil" name:"DevTimeSections"`

	// 上云倍速（支持1，2，4倍速）
	Scale *int64 `json:"Scale,omitnil" name:"Scale"`
}

func (r *AddRecordBackupTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddRecordBackupTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateName")
	delete(f, "TimeSections")
	delete(f, "DevTimeSections")
	delete(f, "Scale")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddRecordBackupTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddRecordBackupTemplateResponseParams struct {
	// 返回数据
	Data *AddRecordBackupTemplateData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AddRecordBackupTemplateResponse struct {
	*tchttp.BaseResponse
	Response *AddRecordBackupTemplateResponseParams `json:"Response"`
}

func (r *AddRecordBackupTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddRecordBackupTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddRecordPlanRequestParams struct {
	// 实时上云计划名称，仅支持中文、英文、数字、_、-，长度不超过32个字符，计划名称全局唯一，不能为空，不能重复
	PlanName *string `json:"PlanName,omitnil" name:"PlanName"`

	// 实时上云模板ID
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 生命周期
	LifeCycle *LifeCycleData `json:"LifeCycle,omitnil" name:"LifeCycle"`

	// 上云计划描述，仅支持中文、英文、数字、_、-，长度不超过128个字符 
	Describe *string `json:"Describe,omitnil" name:"Describe"`

	// 码流类型，default:不指定码流类型，以设备默认推送类型为主， main:主码流，sub:子码流，其他根据设备能力集自定义，不填按默认类型处理，长度不能超过32个字节
	StreamType *string `json:"StreamType,omitnil" name:"StreamType"`

	// 添加录像的设备的通道信息，一次添加通道总数不超过5000个，包括组织目录下的通道数量
	Channels []*ChannelInfo `json:"Channels,omitnil" name:"Channels"`

	// 添加组织目录下所有设备通道，Json数组，可以为空，通道总数量不超过5000个（包括Channel字段的数量）
	OrganizationId []*string `json:"OrganizationId,omitnil" name:"OrganizationId"`
}

type AddRecordPlanRequest struct {
	*tchttp.BaseRequest
	
	// 实时上云计划名称，仅支持中文、英文、数字、_、-，长度不超过32个字符，计划名称全局唯一，不能为空，不能重复
	PlanName *string `json:"PlanName,omitnil" name:"PlanName"`

	// 实时上云模板ID
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 生命周期
	LifeCycle *LifeCycleData `json:"LifeCycle,omitnil" name:"LifeCycle"`

	// 上云计划描述，仅支持中文、英文、数字、_、-，长度不超过128个字符 
	Describe *string `json:"Describe,omitnil" name:"Describe"`

	// 码流类型，default:不指定码流类型，以设备默认推送类型为主， main:主码流，sub:子码流，其他根据设备能力集自定义，不填按默认类型处理，长度不能超过32个字节
	StreamType *string `json:"StreamType,omitnil" name:"StreamType"`

	// 添加录像的设备的通道信息，一次添加通道总数不超过5000个，包括组织目录下的通道数量
	Channels []*ChannelInfo `json:"Channels,omitnil" name:"Channels"`

	// 添加组织目录下所有设备通道，Json数组，可以为空，通道总数量不超过5000个（包括Channel字段的数量）
	OrganizationId []*string `json:"OrganizationId,omitnil" name:"OrganizationId"`
}

func (r *AddRecordPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddRecordPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanName")
	delete(f, "TemplateId")
	delete(f, "LifeCycle")
	delete(f, "Describe")
	delete(f, "StreamType")
	delete(f, "Channels")
	delete(f, "OrganizationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddRecordPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddRecordPlanResponseParams struct {
	// 返回结果
	Data *RecordPlanOptData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AddRecordPlanResponse struct {
	*tchttp.BaseResponse
	Response *AddRecordPlanResponseParams `json:"Response"`
}

func (r *AddRecordPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddRecordPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddRecordRetrieveTaskData struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 任务名称
	TaskName *string `json:"TaskName,omitnil" name:"TaskName"`

	// 取回录像的开始时间
	StartTime *uint64 `json:"StartTime,omitnil" name:"StartTime"`

	// 取回录像的结束时间
	EndTime *uint64 `json:"EndTime,omitnil" name:"EndTime"`

	// 取回模式，1:极速模式，其他暂不支持
	Mode *int64 `json:"Mode,omitnil" name:"Mode"`

	// 副本有效期
	Expiration *int64 `json:"Expiration,omitnil" name:"Expiration"`

	// 任务状态，0:已取回，1:取回中，2:待取回
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 取回容量，单位MB
	Capacity *float64 `json:"Capacity,omitnil" name:"Capacity"`

	// 任务描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Describe *string `json:"Describe,omitnil" name:"Describe"`
}

// Predefined struct for user
type AddRecordRetrieveTaskRequestParams struct {
	// 任务名称，仅支持中文、英文、数字、_、-，长度不超过32个字符，模板名称全局唯一，不能为空，不能重复
	TaskName *string `json:"TaskName,omitnil" name:"TaskName"`

	// 取回录像的开始时间，UTC秒数，例如：1662114146，开始和结束时间段最长为一天，且不能跨天
	StartTime *uint64 `json:"StartTime,omitnil" name:"StartTime"`

	// 取回录像的结束时间，UTC秒数，例如：1662114146，开始和结束时间段最长为一天，且不能跨天
	EndTime *uint64 `json:"EndTime,omitnil" name:"EndTime"`

	// 取回模式， 1:极速模式，其他暂不支持
	Mode *int64 `json:"Mode,omitnil" name:"Mode"`

	// 取回录像副本有效期，最小为1天，最大为365天
	Expiration *int64 `json:"Expiration,omitnil" name:"Expiration"`

	// 设备通道，一个任务最多32个设备通道
	Channels []*ChannelInfo `json:"Channels,omitnil" name:"Channels"`

	// 取回任务描述
	Describe *string `json:"Describe,omitnil" name:"Describe"`
}

type AddRecordRetrieveTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务名称，仅支持中文、英文、数字、_、-，长度不超过32个字符，模板名称全局唯一，不能为空，不能重复
	TaskName *string `json:"TaskName,omitnil" name:"TaskName"`

	// 取回录像的开始时间，UTC秒数，例如：1662114146，开始和结束时间段最长为一天，且不能跨天
	StartTime *uint64 `json:"StartTime,omitnil" name:"StartTime"`

	// 取回录像的结束时间，UTC秒数，例如：1662114146，开始和结束时间段最长为一天，且不能跨天
	EndTime *uint64 `json:"EndTime,omitnil" name:"EndTime"`

	// 取回模式， 1:极速模式，其他暂不支持
	Mode *int64 `json:"Mode,omitnil" name:"Mode"`

	// 取回录像副本有效期，最小为1天，最大为365天
	Expiration *int64 `json:"Expiration,omitnil" name:"Expiration"`

	// 设备通道，一个任务最多32个设备通道
	Channels []*ChannelInfo `json:"Channels,omitnil" name:"Channels"`

	// 取回任务描述
	Describe *string `json:"Describe,omitnil" name:"Describe"`
}

func (r *AddRecordRetrieveTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddRecordRetrieveTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Mode")
	delete(f, "Expiration")
	delete(f, "Channels")
	delete(f, "Describe")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddRecordRetrieveTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddRecordRetrieveTaskResponseParams struct {
	// 返回结果
	Data *AddRecordRetrieveTaskData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AddRecordRetrieveTaskResponse struct {
	*tchttp.BaseResponse
	Response *AddRecordRetrieveTaskResponseParams `json:"Response"`
}

func (r *AddRecordRetrieveTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddRecordRetrieveTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddRecordTemplateRequestParams struct {
	// 模板名称， 仅支持中文、英文、数字、_、-，长度不超过32个字符，模板名称全局唯一，不能为空，不能重复
	TemplateName *string `json:"TemplateName,omitnil" name:"TemplateName"`

	// 上云时间段，按周进行设置，支持一天设置多个时间段，每个时间段不小于10分钟
	TimeSections []*RecordTemplateTimeSections `json:"TimeSections,omitnil" name:"TimeSections"`
}

type AddRecordTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板名称， 仅支持中文、英文、数字、_、-，长度不超过32个字符，模板名称全局唯一，不能为空，不能重复
	TemplateName *string `json:"TemplateName,omitnil" name:"TemplateName"`

	// 上云时间段，按周进行设置，支持一天设置多个时间段，每个时间段不小于10分钟
	TimeSections []*RecordTemplateTimeSections `json:"TimeSections,omitnil" name:"TimeSections"`
}

func (r *AddRecordTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddRecordTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateName")
	delete(f, "TimeSections")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddRecordTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddRecordTemplateResponseParams struct {
	// 返回结果
	Data *RecordTemplateInfo `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AddRecordTemplateResponse struct {
	*tchttp.BaseResponse
	Response *AddRecordTemplateResponseParams `json:"Response"`
}

func (r *AddRecordTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddRecordTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddStreamAuthData struct {
	// 鉴权配置ID（uuid）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil" name:"Id"`

	// 是否开播放鉴权（1:开启,0:关闭）
	// 注意：此字段可能返回 null，表示取不到有效值。
	PullState *int64 `json:"PullState,omitnil" name:"PullState"`

	// 播放密钥（仅支持字母数字，长度0-10位）
	// 注意：此字段可能返回 null，表示取不到有效值。
	PullSecret *string `json:"PullSecret,omitnil" name:"PullSecret"`

	// 播放过期时间（单位：分钟）
	// 注意：此字段可能返回 null，表示取不到有效值。
	PullExpired *int64 `json:"PullExpired,omitnil" name:"PullExpired"`

	// 是否开启推流鉴权（1:开启,0:关闭）
	// 注意：此字段可能返回 null，表示取不到有效值。
	PushState *int64 `json:"PushState,omitnil" name:"PushState"`

	// 推流密钥（仅支持字母数字，长度0-10位）
	// 注意：此字段可能返回 null，表示取不到有效值。
	PushSecret *string `json:"PushSecret,omitnil" name:"PushSecret"`

	// 推流过期时间（单位：分钟）
	// 注意：此字段可能返回 null，表示取不到有效值。
	PushExpired *int64 `json:"PushExpired,omitnil" name:"PushExpired"`

	// 用户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *int64 `json:"AppId,omitnil" name:"AppId"`
}

// Predefined struct for user
type AddStreamAuthRequestParams struct {
	// 鉴权配置ID（uuid）
	Id *string `json:"Id,omitnil" name:"Id"`

	// 是否开播放鉴权（1:开启,0:关闭）
	PullState *int64 `json:"PullState,omitnil" name:"PullState"`

	// 播放密钥（仅支持字母数字，长度0-10位）
	PullSecret *string `json:"PullSecret,omitnil" name:"PullSecret"`

	// 播放过期时间（单位：分钟）
	PullExpired *int64 `json:"PullExpired,omitnil" name:"PullExpired"`

	// 是否开启推流鉴权（1:开启,0:关闭）
	PushState *int64 `json:"PushState,omitnil" name:"PushState"`

	// 推流密钥（仅支持字母数字，长度0-10位）
	PushSecret *string `json:"PushSecret,omitnil" name:"PushSecret"`

	// 推流过期时间（单位：分钟）
	PushExpired *int64 `json:"PushExpired,omitnil" name:"PushExpired"`
}

type AddStreamAuthRequest struct {
	*tchttp.BaseRequest
	
	// 鉴权配置ID（uuid）
	Id *string `json:"Id,omitnil" name:"Id"`

	// 是否开播放鉴权（1:开启,0:关闭）
	PullState *int64 `json:"PullState,omitnil" name:"PullState"`

	// 播放密钥（仅支持字母数字，长度0-10位）
	PullSecret *string `json:"PullSecret,omitnil" name:"PullSecret"`

	// 播放过期时间（单位：分钟）
	PullExpired *int64 `json:"PullExpired,omitnil" name:"PullExpired"`

	// 是否开启推流鉴权（1:开启,0:关闭）
	PushState *int64 `json:"PushState,omitnil" name:"PushState"`

	// 推流密钥（仅支持字母数字，长度0-10位）
	PushSecret *string `json:"PushSecret,omitnil" name:"PushSecret"`

	// 推流过期时间（单位：分钟）
	PushExpired *int64 `json:"PushExpired,omitnil" name:"PushExpired"`
}

func (r *AddStreamAuthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddStreamAuthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "PullState")
	delete(f, "PullSecret")
	delete(f, "PullExpired")
	delete(f, "PushState")
	delete(f, "PushSecret")
	delete(f, "PushExpired")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddStreamAuthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddStreamAuthResponseParams struct {
	// 设置推拉流鉴权返回数据
	Data *AddStreamAuthData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AddStreamAuthResponse struct {
	*tchttp.BaseResponse
	Response *AddStreamAuthResponseParams `json:"Response"`
}

func (r *AddStreamAuthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddStreamAuthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddUserDeviceRequestParams struct {
	// 设备名称，仅支持中文、英文、数字、_、-，长度不超过32个字符；（设备名称无需全局唯一，可以重复）
	Name *string `json:"Name,omitnil" name:"Name"`

	// 设备接入协议（1:RTMP,2:GB,3:GW）
	AccessProtocol *int64 `json:"AccessProtocol,omitnil" name:"AccessProtocol"`

	// 设备类型，1:IPC,2:NVR；（若设备接入协议选择RTMP，则设备类型只能选择IPC）
	Type *int64 `json:"Type,omitnil" name:"Type"`

	// 设备所属组织ID，从查询组织接口DescribeOrganization中获取
	OrganizationId *string `json:"OrganizationId,omitnil" name:"OrganizationId"`

	// 设备接入服务节点ID（从查询设备可用服务节点接口DescribeDeviceRegion中获取的Value字段）
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 设备流传输协议，1:UDP,2:TCP；(国标设备有效，不填写则默认UDP协议)
	TransportProtocol *int64 `json:"TransportProtocol,omitnil" name:"TransportProtocol"`

	// 设备密码（国标，网关设备必填，仅支持数字组合，长度为1-64个字符）
	Password *string `json:"Password,omitnil" name:"Password"`

	// 设备描述，仅支持中文、英文、数字、_、-，长度不超过128个字符
	Description *string `json:"Description,omitnil" name:"Description"`

	// 设备接入网关ID，从查询网关列表接口中ListGateways获取（仅网关接入需要）
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 网关接入协议类型（从查询网关接入协议接口DescribeGatewayProtocol中获取）1.海康SDK，2.大华SDK，3.宇视SDK，4.Onvif（仅网关接入需要）
	ProtocolType *int64 `json:"ProtocolType,omitnil" name:"ProtocolType"`

	// 设备接入IP（仅网关接入需要）
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 设备端口（仅网关接入需要）
	Port *int64 `json:"Port,omitnil" name:"Port"`

	// 设备用户名（仅网关接入需要）
	Username *string `json:"Username,omitnil" name:"Username"`
}

type AddUserDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 设备名称，仅支持中文、英文、数字、_、-，长度不超过32个字符；（设备名称无需全局唯一，可以重复）
	Name *string `json:"Name,omitnil" name:"Name"`

	// 设备接入协议（1:RTMP,2:GB,3:GW）
	AccessProtocol *int64 `json:"AccessProtocol,omitnil" name:"AccessProtocol"`

	// 设备类型，1:IPC,2:NVR；（若设备接入协议选择RTMP，则设备类型只能选择IPC）
	Type *int64 `json:"Type,omitnil" name:"Type"`

	// 设备所属组织ID，从查询组织接口DescribeOrganization中获取
	OrganizationId *string `json:"OrganizationId,omitnil" name:"OrganizationId"`

	// 设备接入服务节点ID（从查询设备可用服务节点接口DescribeDeviceRegion中获取的Value字段）
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 设备流传输协议，1:UDP,2:TCP；(国标设备有效，不填写则默认UDP协议)
	TransportProtocol *int64 `json:"TransportProtocol,omitnil" name:"TransportProtocol"`

	// 设备密码（国标，网关设备必填，仅支持数字组合，长度为1-64个字符）
	Password *string `json:"Password,omitnil" name:"Password"`

	// 设备描述，仅支持中文、英文、数字、_、-，长度不超过128个字符
	Description *string `json:"Description,omitnil" name:"Description"`

	// 设备接入网关ID，从查询网关列表接口中ListGateways获取（仅网关接入需要）
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 网关接入协议类型（从查询网关接入协议接口DescribeGatewayProtocol中获取）1.海康SDK，2.大华SDK，3.宇视SDK，4.Onvif（仅网关接入需要）
	ProtocolType *int64 `json:"ProtocolType,omitnil" name:"ProtocolType"`

	// 设备接入IP（仅网关接入需要）
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 设备端口（仅网关接入需要）
	Port *int64 `json:"Port,omitnil" name:"Port"`

	// 设备用户名（仅网关接入需要）
	Username *string `json:"Username,omitnil" name:"Username"`
}

func (r *AddUserDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddUserDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "AccessProtocol")
	delete(f, "Type")
	delete(f, "OrganizationId")
	delete(f, "ClusterId")
	delete(f, "TransportProtocol")
	delete(f, "Password")
	delete(f, "Description")
	delete(f, "GatewayId")
	delete(f, "ProtocolType")
	delete(f, "Ip")
	delete(f, "Port")
	delete(f, "Username")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddUserDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddUserDeviceResponseParams struct {
	// 增加设备返回数据
	Data *AddDeviceData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AddUserDeviceResponse struct {
	*tchttp.BaseResponse
	Response *AddUserDeviceResponseParams `json:"Response"`
}

func (r *AddUserDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddUserDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BaseAIResultInfo struct {
	// 名称。返回值有人体识别结果名称(person)、宠物识别结果名称(cat和dog) 、车辆车牌识别结果名称(vehicle)
	Name *string `json:"Name,omitnil" name:"Name"`

	// 置信度
	Score *uint64 `json:"Score,omitnil" name:"Score"`

	// 截图中坐标信息
	Location *Location `json:"Location,omitnil" name:"Location"`
}

type BatchOperateDeviceData struct {
	// 任务 ID（用于在查询任务的子任务列表接口ListSubTasks中查询任务进度）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

// Predefined struct for user
type BatchOperateDeviceRequestParams struct {
	// 设备 ID 数组（从获取设备列表接口ListDevices中获取）
	DeviceIds []*string `json:"DeviceIds,omitnil" name:"DeviceIds"`

	// 操作命令（enable：启用；disable：禁用；delete：删除）
	Cmd *string `json:"Cmd,omitnil" name:"Cmd"`
}

type BatchOperateDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 设备 ID 数组（从获取设备列表接口ListDevices中获取）
	DeviceIds []*string `json:"DeviceIds,omitnil" name:"DeviceIds"`

	// 操作命令（enable：启用；disable：禁用；delete：删除）
	Cmd *string `json:"Cmd,omitnil" name:"Cmd"`
}

func (r *BatchOperateDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchOperateDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceIds")
	delete(f, "Cmd")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchOperateDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchOperateDeviceResponseParams struct {
	// 返回结果
	Data *BatchOperateDeviceData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type BatchOperateDeviceResponse struct {
	*tchttp.BaseResponse
	Response *BatchOperateDeviceResponseParams `json:"Response"`
}

func (r *BatchOperateDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchOperateDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BitRateInfo struct {
	// 通道Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// 码率,单位:kbps
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bitrate *float64 `json:"Bitrate,omitnil" name:"Bitrate"`
}

type BodyAIResultInfo struct {
	// 时间字符串
	Time *string `json:"Time,omitnil" name:"Time"`

	// 截图 URL
	Url *string `json:"Url,omitnil" name:"Url"`

	// 人体信息
	BodyInfo []*BaseAIResultInfo `json:"BodyInfo,omitnil" name:"BodyInfo"`
}

type CarAIResultInfo struct {
	// 车系
	Serial *string `json:"Serial,omitnil" name:"Serial"`

	// 车辆品牌
	Brand *string `json:"Brand,omitnil" name:"Brand"`

	// 车辆类型
	Type *string `json:"Type,omitnil" name:"Type"`

	// 车辆颜色
	Color *string `json:"Color,omitnil" name:"Color"`

	// 置信度，0 - 100
	Confidence *int64 `json:"Confidence,omitnil" name:"Confidence"`

	// 年份，识别不出年份时返回0
	Year *int64 `json:"Year,omitnil" name:"Year"`

	// 车牌信息
	PlateContent *PlateContent `json:"PlateContent,omitnil" name:"PlateContent"`

	// 截图中坐标信息
	Location *Location `json:"Location,omitnil" name:"Location"`
}

type ChannelInfo struct {
	// 通道所属的设备ID
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`

	// 设备通道ID，一个设备通道只允许被一个上云计划添加
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`
}

// Predefined struct for user
type CheckDomainRequestParams struct {
	// 播放域名
	PlayDomain *string `json:"PlayDomain,omitnil" name:"PlayDomain"`

	// CNAME 记录值
	InternalDomain *string `json:"InternalDomain,omitnil" name:"InternalDomain"`
}

type CheckDomainRequest struct {
	*tchttp.BaseRequest
	
	// 播放域名
	PlayDomain *string `json:"PlayDomain,omitnil" name:"PlayDomain"`

	// CNAME 记录值
	InternalDomain *string `json:"InternalDomain,omitnil" name:"InternalDomain"`
}

func (r *CheckDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlayDomain")
	delete(f, "InternalDomain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckDomainResponseParams struct {
	// 是否备案
	Data *bool `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CheckDomainResponse struct {
	*tchttp.BaseResponse
	Response *CheckDomainResponseParams `json:"Response"`
}

func (r *CheckDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChefClothAIResultInfo struct {
	// 时间字符串
	Time *string `json:"Time,omitnil" name:"Time"`

	// 截图 URL
	Url *string `json:"Url,omitnil" name:"Url"`

	// 厨师服信息
	ChefClothInfoInfo []*BaseAIResultInfo `json:"ChefClothInfoInfo,omitnil" name:"ChefClothInfoInfo"`
}

type ChefHatAIResultInfo struct {
	// 时间字符串
	Time *string `json:"Time,omitnil" name:"Time"`

	// 截图 URL
	Url *string `json:"Url,omitnil" name:"Url"`

	// 厨师帽信息
	ChefHatInfo []*BaseAIResultInfo `json:"ChefHatInfo,omitnil" name:"ChefHatInfo"`
}

// Predefined struct for user
type ControlDevicePTZRequestParams struct {
	// 通道 ID（从通道查询接口DescribeDeviceChannel中获取）
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// 命令类型（上:up,下:down,左:left,右:right
	// 上左:leftup,上右:rightup,下左:leftdown,下右:rightdown
	// 放大:zoomin,缩小:zoomout
	// 聚焦远:focusfar,聚焦近:focusnear
	// 光圈放大:irisin,光圈缩小:irisout）
	Type *string `json:"Type,omitnil" name:"Type"`

	// 命令描述（速度值范围1-8）
	Speed *int64 `json:"Speed,omitnil" name:"Speed"`
}

type ControlDevicePTZRequest struct {
	*tchttp.BaseRequest
	
	// 通道 ID（从通道查询接口DescribeDeviceChannel中获取）
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// 命令类型（上:up,下:down,左:left,右:right
	// 上左:leftup,上右:rightup,下左:leftdown,下右:rightdown
	// 放大:zoomin,缩小:zoomout
	// 聚焦远:focusfar,聚焦近:focusnear
	// 光圈放大:irisin,光圈缩小:irisout）
	Type *string `json:"Type,omitnil" name:"Type"`

	// 命令描述（速度值范围1-8）
	Speed *int64 `json:"Speed,omitnil" name:"Speed"`
}

func (r *ControlDevicePTZRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlDevicePTZRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelId")
	delete(f, "Type")
	delete(f, "Speed")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ControlDevicePTZRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ControlDevicePTZResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ControlDevicePTZResponse struct {
	*tchttp.BaseResponse
	Response *ControlDevicePTZResponseParams `json:"Response"`
}

func (r *ControlDevicePTZResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlDevicePTZResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ControlDevicePresetRequestParams struct {
	// 通道 ID（从通道查询接口DescribeDeviceChannel中获取）
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// 命令（goto:预置位调用；
	// set:预置位设置；
	// del:预置位删除）
	Cmd *string `json:"Cmd,omitnil" name:"Cmd"`

	// 预置位索引（只支持1-10的索引位置，超出报错）
	Index *int64 `json:"Index,omitnil" name:"Index"`
}

type ControlDevicePresetRequest struct {
	*tchttp.BaseRequest
	
	// 通道 ID（从通道查询接口DescribeDeviceChannel中获取）
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// 命令（goto:预置位调用；
	// set:预置位设置；
	// del:预置位删除）
	Cmd *string `json:"Cmd,omitnil" name:"Cmd"`

	// 预置位索引（只支持1-10的索引位置，超出报错）
	Index *int64 `json:"Index,omitnil" name:"Index"`
}

func (r *ControlDevicePresetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlDevicePresetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelId")
	delete(f, "Cmd")
	delete(f, "Index")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ControlDevicePresetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ControlDevicePresetResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ControlDevicePresetResponse struct {
	*tchttp.BaseResponse
	Response *ControlDevicePresetResponseParams `json:"Response"`
}

func (r *ControlDevicePresetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlDevicePresetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ControlDeviceStreamData struct {
	// flv 流地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Flv *string `json:"Flv,omitnil" name:"Flv"`

	// hls 流地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Hls *string `json:"Hls,omitnil" name:"Hls"`

	// rtmp 流地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rtmp *string `json:"Rtmp,omitnil" name:"Rtmp"`
}

// Predefined struct for user
type ControlDeviceStreamRequestParams struct {
	// 通道 ID（从通道查询接口DescribeDeviceChannel中获取）
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// 流类型（1:主码流；
	// 2:子码流（不可以和 Resolution 同时下发））
	StreamType *string `json:"StreamType,omitnil" name:"StreamType"`

	// 分辨率（1:QCIF；
	// 2:CIF；
	// 3:4CIF；
	// 4:D1；
	// 5:720P；
	// 6:1080P/I；
	// 自定义的19201080等等（需设备支持）（不可以和 StreamType 同时下发））
	Resolution *string `json:"Resolution,omitnil" name:"Resolution"`
}

type ControlDeviceStreamRequest struct {
	*tchttp.BaseRequest
	
	// 通道 ID（从通道查询接口DescribeDeviceChannel中获取）
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// 流类型（1:主码流；
	// 2:子码流（不可以和 Resolution 同时下发））
	StreamType *string `json:"StreamType,omitnil" name:"StreamType"`

	// 分辨率（1:QCIF；
	// 2:CIF；
	// 3:4CIF；
	// 4:D1；
	// 5:720P；
	// 6:1080P/I；
	// 自定义的19201080等等（需设备支持）（不可以和 StreamType 同时下发））
	Resolution *string `json:"Resolution,omitnil" name:"Resolution"`
}

func (r *ControlDeviceStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlDeviceStreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelId")
	delete(f, "StreamType")
	delete(f, "Resolution")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ControlDeviceStreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ControlDeviceStreamResponseParams struct {
	// 返回数据
	Data *ControlDeviceStreamData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ControlDeviceStreamResponse struct {
	*tchttp.BaseResponse
	Response *ControlDeviceStreamResponseParams `json:"Response"`
}

func (r *ControlDeviceStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlDeviceStreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ControlRecordRequestParams struct {
	// 通道ID（录像播放地址格式 https://${domain}/live/${ChannelId}@${Session}）
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// 录像会话 ID （ 录像播放地址格式 https://${domain}/live/${ChannelId}@${Session}）
	Session *string `json:"Session,omitnil" name:"Session"`

	// 录像操作类型 （play:播放；pause:暂停 ；stop:关闭）
	ControlAction *string `json:"ControlAction,omitnil" name:"ControlAction"`

	// 跳转进度 （ 参数应大于等于0，跳转到录像开始时间的相对时间（单位秒），例如0就是跳转到录像开始的时间,不可以和 Scale 参数同时出现）
	Position *int64 `json:"Position,omitnil" name:"Position"`

	// 速度 （ 范围（0.25,0.5,1,2,4,8），不可以和 Pos 参数同时出现）
	Scale *float64 `json:"Scale,omitnil" name:"Scale"`
}

type ControlRecordRequest struct {
	*tchttp.BaseRequest
	
	// 通道ID（录像播放地址格式 https://${domain}/live/${ChannelId}@${Session}）
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// 录像会话 ID （ 录像播放地址格式 https://${domain}/live/${ChannelId}@${Session}）
	Session *string `json:"Session,omitnil" name:"Session"`

	// 录像操作类型 （play:播放；pause:暂停 ；stop:关闭）
	ControlAction *string `json:"ControlAction,omitnil" name:"ControlAction"`

	// 跳转进度 （ 参数应大于等于0，跳转到录像开始时间的相对时间（单位秒），例如0就是跳转到录像开始的时间,不可以和 Scale 参数同时出现）
	Position *int64 `json:"Position,omitnil" name:"Position"`

	// 速度 （ 范围（0.25,0.5,1,2,4,8），不可以和 Pos 参数同时出现）
	Scale *float64 `json:"Scale,omitnil" name:"Scale"`
}

func (r *ControlRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelId")
	delete(f, "Session")
	delete(f, "ControlAction")
	delete(f, "Position")
	delete(f, "Scale")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ControlRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ControlRecordResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ControlRecordResponse struct {
	*tchttp.BaseResponse
	Response *ControlRecordResponseParams `json:"Response"`
}

func (r *ControlRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ControlRecordTimelineRequestParams struct {
	// 通道 ID（从通道查询接口DescribeDeviceChannel中获取）
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// 起始时间
	Start *int64 `json:"Start,omitnil" name:"Start"`

	// 结束时间
	End *int64 `json:"End,omitnil" name:"End"`
}

type ControlRecordTimelineRequest struct {
	*tchttp.BaseRequest
	
	// 通道 ID（从通道查询接口DescribeDeviceChannel中获取）
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// 起始时间
	Start *int64 `json:"Start,omitnil" name:"Start"`

	// 结束时间
	End *int64 `json:"End,omitnil" name:"End"`
}

func (r *ControlRecordTimelineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlRecordTimelineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelId")
	delete(f, "Start")
	delete(f, "End")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ControlRecordTimelineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ControlRecordTimelineResponseParams struct {
	// 返回数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*Timeline `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ControlRecordTimelineResponse struct {
	*tchttp.BaseResponse
	Response *ControlRecordTimelineResponseParams `json:"Response"`
}

func (r *ControlRecordTimelineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlRecordTimelineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAITaskRequestParams struct {
	// AI任务ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

type DeleteAITaskRequest struct {
	*tchttp.BaseRequest
	
	// AI任务ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

func (r *DeleteAITaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAITaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAITaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAITaskResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteAITaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAITaskResponseParams `json:"Response"`
}

func (r *DeleteAITaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAITaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDomainRequestParams struct {
	// 域名 ID
	Id *string `json:"Id,omitnil" name:"Id"`
}

type DeleteDomainRequest struct {
	*tchttp.BaseRequest
	
	// 域名 ID
	Id *string `json:"Id,omitnil" name:"Id"`
}

func (r *DeleteDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDomainResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteDomainResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDomainResponseParams `json:"Response"`
}

func (r *DeleteDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGatewayRequestParams struct {
	// 网关索引ID（从获取网关列表接口ListGateways中获取）
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`
}

type DeleteGatewayRequest struct {
	*tchttp.BaseRequest
	
	// 网关索引ID（从获取网关列表接口ListGateways中获取）
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`
}

func (r *DeleteGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGatewayResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteGatewayResponse struct {
	*tchttp.BaseResponse
	Response *DeleteGatewayResponseParams `json:"Response"`
}

func (r *DeleteGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationRequestParams struct {
	// 组织ID（从查询组织接口DescribeOrganization中获取）
	OrganizationId *string `json:"OrganizationId,omitnil" name:"OrganizationId"`
}

type DeleteOrganizationRequest struct {
	*tchttp.BaseRequest
	
	// 组织ID（从查询组织接口DescribeOrganization中获取）
	OrganizationId *string `json:"OrganizationId,omitnil" name:"OrganizationId"`
}

func (r *DeleteOrganizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrganizationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteOrganizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteOrganizationResponse struct {
	*tchttp.BaseResponse
	Response *DeleteOrganizationResponseParams `json:"Response"`
}

func (r *DeleteOrganizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRecordBackupPlanRequestParams struct {
	// 录像上云计划ID（从查询录像上云计划列表接口ListRecordBackupPlans中获取）
	PlanId *string `json:"PlanId,omitnil" name:"PlanId"`
}

type DeleteRecordBackupPlanRequest struct {
	*tchttp.BaseRequest
	
	// 录像上云计划ID（从查询录像上云计划列表接口ListRecordBackupPlans中获取）
	PlanId *string `json:"PlanId,omitnil" name:"PlanId"`
}

func (r *DeleteRecordBackupPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecordBackupPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRecordBackupPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRecordBackupPlanResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteRecordBackupPlanResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRecordBackupPlanResponseParams `json:"Response"`
}

func (r *DeleteRecordBackupPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecordBackupPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRecordBackupTemplateRequestParams struct {
	// 模板ID（从查询录像上云模板列表接口ListRecordBackupTemplates中获取）
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`
}

type DeleteRecordBackupTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板ID（从查询录像上云模板列表接口ListRecordBackupTemplates中获取）
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`
}

func (r *DeleteRecordBackupTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecordBackupTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRecordBackupTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRecordBackupTemplateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteRecordBackupTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRecordBackupTemplateResponseParams `json:"Response"`
}

func (r *DeleteRecordBackupTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecordBackupTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRecordPlanRequestParams struct {
	// 上云计划ID
	PlanId *string `json:"PlanId,omitnil" name:"PlanId"`
}

type DeleteRecordPlanRequest struct {
	*tchttp.BaseRequest
	
	// 上云计划ID
	PlanId *string `json:"PlanId,omitnil" name:"PlanId"`
}

func (r *DeleteRecordPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecordPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRecordPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRecordPlanResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteRecordPlanResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRecordPlanResponseParams `json:"Response"`
}

func (r *DeleteRecordPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecordPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRecordRetrieveTaskRequestParams struct {
	// 取回任务ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

type DeleteRecordRetrieveTaskRequest struct {
	*tchttp.BaseRequest
	
	// 取回任务ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

func (r *DeleteRecordRetrieveTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecordRetrieveTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRecordRetrieveTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRecordRetrieveTaskResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteRecordRetrieveTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRecordRetrieveTaskResponseParams `json:"Response"`
}

func (r *DeleteRecordRetrieveTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecordRetrieveTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRecordTemplateRequestParams struct {
	// 模板ID
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`
}

type DeleteRecordTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板ID
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`
}

func (r *DeleteRecordTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecordTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRecordTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRecordTemplateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteRecordTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRecordTemplateResponseParams `json:"Response"`
}

func (r *DeleteRecordTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecordTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserDeviceRequestParams struct {
	// 设备ID（从获取设备列表ListDevices接口中获取）
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`
}

type DeleteUserDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 设备ID（从获取设备列表ListDevices接口中获取）
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`
}

func (r *DeleteUserDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUserDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserDeviceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteUserDeviceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUserDeviceResponseParams `json:"Response"`
}

func (r *DeleteUserDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAITaskRequestParams struct {
	// AI任务ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

type DescribeAITaskRequest struct {
	*tchttp.BaseRequest
	
	// AI任务ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

func (r *DescribeAITaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAITaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAITaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAITaskResponseParams struct {
	// AI任务详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *AITaskInfo `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAITaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAITaskResponseParams `json:"Response"`
}

func (r *DescribeAITaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAITaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAITaskResultRequestParams struct {
	// AI 任务 ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 通道ID
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// 桶内文件的路径。
	Object *string `json:"Object,omitnil" name:"Object"`

	// AI 任务识别类型。可选值为 Facemask(口罩识别)、Chefhat(厨师帽识别)、Smoking(抽烟检测)、Chefcloth(厨师服识别)、PhoneCall(接打电话识别)、Pet(宠物识别)、Body(人体识别)和 Car(车辆车牌识别)
	DetectType *string `json:"DetectType,omitnil" name:"DetectType"`

	// 开始时间时间。秒级时间戳。开始时间和结束时间跨度小于等于30天
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 结束时间时间。秒级时间戳。开始时间和结束时间跨度小于等于30天
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 页码。默认为1
	PageNumber *uint64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// 每页 AI 识别结果数量。可选值1～100，默认为10（按时间倒序显示识别结果）
	PageSize *uint64 `json:"PageSize,omitnil" name:"PageSize"`
}

type DescribeAITaskResultRequest struct {
	*tchttp.BaseRequest
	
	// AI 任务 ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 通道ID
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// 桶内文件的路径。
	Object *string `json:"Object,omitnil" name:"Object"`

	// AI 任务识别类型。可选值为 Facemask(口罩识别)、Chefhat(厨师帽识别)、Smoking(抽烟检测)、Chefcloth(厨师服识别)、PhoneCall(接打电话识别)、Pet(宠物识别)、Body(人体识别)和 Car(车辆车牌识别)
	DetectType *string `json:"DetectType,omitnil" name:"DetectType"`

	// 开始时间时间。秒级时间戳。开始时间和结束时间跨度小于等于30天
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 结束时间时间。秒级时间戳。开始时间和结束时间跨度小于等于30天
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 页码。默认为1
	PageNumber *uint64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// 每页 AI 识别结果数量。可选值1～100，默认为10（按时间倒序显示识别结果）
	PageSize *uint64 `json:"PageSize,omitnil" name:"PageSize"`
}

func (r *DescribeAITaskResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAITaskResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "ChannelId")
	delete(f, "Object")
	delete(f, "DetectType")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAITaskResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAITaskResultResponseParams struct {
	// AI识别结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *AITaskResultData `json:"Data,omitnil" name:"Data"`

	// AI识别结果数量
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAITaskResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAITaskResultResponseParams `json:"Response"`
}

func (r *DescribeAITaskResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAITaskResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCNAMERequestParams struct {
	// 服务节点 ID（从查询域名可绑定服务节点接口DescribeDomainRegion中获取）
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

type DescribeCNAMERequest struct {
	*tchttp.BaseRequest
	
	// 服务节点 ID（从查询域名可绑定服务节点接口DescribeDomainRegion中获取）
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

func (r *DescribeCNAMERequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCNAMERequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCNAMERequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCNAMEResponseParams struct {
	// CNAME 记录值
	Data *string `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCNAMEResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCNAMEResponseParams `json:"Response"`
}

func (r *DescribeCNAMEResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCNAMEResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceChannelData struct {
	// 设备 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`

	// 通道 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// 通道编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelCode *string `json:"ChannelCode,omitnil" name:"ChannelCode"`

	// 通道名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 流状态（0:未传输,1:传输中）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 是否可控 Ptz（0:不可控,1:可控）
	// 注意：此字段可能返回 null，表示取不到有效值。
	PTZType *int64 `json:"PTZType,omitnil" name:"PTZType"`

	// 通道厂商
	// 注意：此字段可能返回 null，表示取不到有效值。
	Manufacturer *string `json:"Manufacturer,omitnil" name:"Manufacturer"`

	// 通道支持分辨率（分辨率列表由‘/’隔开，国标协议样例（6/3），自定义样例（12800960/640480））
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resolution *string `json:"Resolution,omitnil" name:"Resolution"`

	// 通道在离线状态（0:离线,1:在线）
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *int64 `json:"State,omitnil" name:"State"`

	// 所在地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil" name:"Region"`
}

// Predefined struct for user
type DescribeDeviceChannelRequestParams struct {
	// 设备ID（从获取设备列表接口ListDevices中获取）
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`
}

type DescribeDeviceChannelRequest struct {
	*tchttp.BaseRequest
	
	// 设备ID（从获取设备列表接口ListDevices中获取）
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`
}

func (r *DescribeDeviceChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceChannelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceChannelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceChannelResponseParams struct {
	// 返回结果
	Data []*DescribeDeviceChannelData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDeviceChannelResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceChannelResponseParams `json:"Response"`
}

func (r *DescribeDeviceChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceData struct {
	// 设备ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`

	// 设备编码（即我们为设备生成的20位国标编码）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *string `json:"Code,omitnil" name:"Code"`

	// 设备名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 设备接入协议，1:RTMP,2:GB,3:GW 
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessProtocol *int64 `json:"AccessProtocol,omitnil" name:"AccessProtocol"`

	// 设备类型，1:IPC,2:NVR
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *int64 `json:"Type,omitnil" name:"Type"`

	// 设备接入服务节点id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 设备接入服务节点名称
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitnil" name:"ClusterName"`

	// 设备流传输协议，1:UDP,2:TCP 
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransportProtocol *int64 `json:"TransportProtocol,omitnil" name:"TransportProtocol"`

	// 设备密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitnil" name:"Password"`

	// 设备描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// sip服务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SipId *string `json:"SipId,omitnil" name:"SipId"`

	// sip服务域
	// 注意：此字段可能返回 null，表示取不到有效值。
	SipDomain *string `json:"SipDomain,omitnil" name:"SipDomain"`

	// sip服务IP地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	SipIp *string `json:"SipIp,omitnil" name:"SipIp"`

	// sip服务端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	SipPort *int64 `json:"SipPort,omitnil" name:"SipPort"`

	// Rtmp设备推流地址(仅rtmp设备有效)
	// 注意：此字段可能返回 null，表示取不到有效值。
	PushStreamUrl *string `json:"PushStreamUrl,omitnil" name:"PushStreamUrl"`

	// 设备状态，0:未注册,1:在线,2:离线,3:禁用
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 设备所属组织ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationId *string `json:"OrganizationId,omitnil" name:"OrganizationId"`

	// 设备接入网关ID，从查询网关列表接口中获取（仅网关接入需要）
	// 注意：此字段可能返回 null，表示取不到有效值。
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 设备所属网关名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GatewayName *string `json:"GatewayName,omitnil" name:"GatewayName"`

	// 设备网关协议名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProtocolTypeName *string `json:"ProtocolTypeName,omitnil" name:"ProtocolTypeName"`

	// 网关接入协议类型，1.海康SDK，2.大华SDK，3.宇视SDK，4.Onvif（仅网关接入需要）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProtocolType *int64 `json:"ProtocolType,omitnil" name:"ProtocolType"`

	// 设备接入IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 设备Port
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitnil" name:"Port"`

	// 设备用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Username *string `json:"Username,omitnil" name:"Username"`

	// 设备地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil" name:"Region"`

	// 设备厂商
	// 注意：此字段可能返回 null，表示取不到有效值。
	Manufacturer *string `json:"Manufacturer,omitnil" name:"Manufacturer"`
}

type DescribeDevicePresetData struct {
	// 预置位索引    只支持1-10的索引
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index *int64 `json:"Index,omitnil" name:"Index"`

	// 预置位名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`
}

// Predefined struct for user
type DescribeDevicePresetRequestParams struct {
	// 通道ID（从通道查询接口DescribeDeviceChannel中获取）
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`
}

type DescribeDevicePresetRequest struct {
	*tchttp.BaseRequest
	
	// 通道ID（从通道查询接口DescribeDeviceChannel中获取）
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`
}

func (r *DescribeDevicePresetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDevicePresetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDevicePresetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDevicePresetResponseParams struct {
	// 返回数据
	Data []*DescribeDevicePresetData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDevicePresetResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDevicePresetResponseParams `json:"Response"`
}

func (r *DescribeDevicePresetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDevicePresetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceRegion struct {
	// 服务节点描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitnil" name:"Label"`

	// 服务节点 ID（对应为其他接口中所需的 ClusterId）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 地域信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil" name:"Region"`
}

// Predefined struct for user
type DescribeDeviceRegionRequestParams struct {

}

type DescribeDeviceRegionRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeDeviceRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceRegionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceRegionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceRegionResponseParams struct {
	// 返回数据
	Data []*DescribeDeviceRegion `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDeviceRegionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceRegionResponseParams `json:"Response"`
}

func (r *DescribeDeviceRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainData struct {
	// 域名ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil" name:"Id"`

	// 播放域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlayDomain *string `json:"PlayDomain,omitnil" name:"PlayDomain"`

	// CNAME 记录值
	// 注意：此字段可能返回 null，表示取不到有效值。
	InternalDomain *string `json:"InternalDomain,omitnil" name:"InternalDomain"`

	// 是否上传证书（0：否，1：是）
	// 注意：此字段可能返回 null，表示取不到有效值。
	HaveCert *int64 `json:"HaveCert,omitnil" name:"HaveCert"`

	// 服务节点 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 服务节点名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitnil" name:"ClusterName"`

	// 用户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *int64 `json:"AppId,omitnil" name:"AppId"`
}

type DescribeDomainRegionData struct {
	// 服务节点描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitnil" name:"Label"`

	// 服务节点 ID（对应为其他接口中所需的 ClusterId）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 地域信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil" name:"Region"`
}

// Predefined struct for user
type DescribeDomainRegionRequestParams struct {

}

type DescribeDomainRegionRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeDomainRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainRegionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainRegionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainRegionResponseParams struct {
	// 返回数据
	Data []*DescribeDomainRegionData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDomainRegionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainRegionResponseParams `json:"Response"`
}

func (r *DescribeDomainRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainRequestParams struct {

}

type DescribeDomainRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainResponseParams struct {
	// 返回数据
	Data []*DescribeDomainData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDomainResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainResponseParams `json:"Response"`
}

func (r *DescribeDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewayData struct {
	// 网关索引ID，用于网关查询，更新，删除操作
	// 注意：此字段可能返回 null，表示取不到有效值。
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 网关编码，由网关设备生成的唯一编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	GwId *string `json:"GwId,omitnil" name:"GwId"`

	// 网关名称，仅支持中文、英文、数字、_、-，长度不超过32个字符
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 网关描述，仅支持中文、英文、数字、_、-，长度不超过128个字符
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 服务节点id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 服务节点名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitnil" name:"ClusterName"`

	// 网关状态，0：离线，1:在线
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 网关版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version []*GatewayVersion `json:"Version,omitnil" name:"Version"`

	// 网关下挂设备数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceNum *int64 `json:"DeviceNum,omitnil" name:"DeviceNum"`

	// 激活时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitnil" name:"CreatedAt"`

	// 所属地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil" name:"Region"`
}

type DescribeGatewayMonitor struct {
	// 设备接入总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceTotal *int64 `json:"DeviceTotal,omitnil" name:"DeviceTotal"`

	// 设备在线数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceOnline *int64 `json:"DeviceOnline,omitnil" name:"DeviceOnline"`

	// 设备离线数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceOffline *int64 `json:"DeviceOffline,omitnil" name:"DeviceOffline"`

	// 视频通道总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelTotal *int64 `json:"ChannelTotal,omitnil" name:"ChannelTotal"`

	// 视频通道在线数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelOnline *int64 `json:"ChannelOnline,omitnil" name:"ChannelOnline"`

	// 视频通道离线数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelOffline *int64 `json:"ChannelOffline,omitnil" name:"ChannelOffline"`

	// 网关上行流量,单位kbps
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpFlow *int64 `json:"UpFlow,omitnil" name:"UpFlow"`

	// 流在传输中的通道数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelPull *int64 `json:"ChannelPull,omitnil" name:"ChannelPull"`

	// 流未传输中的通道数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelUnPull *int64 `json:"ChannelUnPull,omitnil" name:"ChannelUnPull"`
}

// Predefined struct for user
type DescribeGatewayMonitorRequestParams struct {
	// 网关索引ID（从获取网关列表接口ListGateways中获取）
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`
}

type DescribeGatewayMonitorRequest struct {
	*tchttp.BaseRequest
	
	// 网关索引ID（从获取网关列表接口ListGateways中获取）
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`
}

func (r *DescribeGatewayMonitorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatewayMonitorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGatewayMonitorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGatewayMonitorResponseParams struct {
	// 返回数据
	Data *DescribeGatewayMonitor `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeGatewayMonitorResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGatewayMonitorResponseParams `json:"Response"`
}

func (r *DescribeGatewayMonitorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatewayMonitorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewayProtocolData struct {
	// 接入协议的字典码
	// 注意：此字段可能返回 null，表示取不到有效值。
	TypeCode *string `json:"TypeCode,omitnil" name:"TypeCode"`

	// 接入协议类型值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *int64 `json:"Value,omitnil" name:"Value"`

	// 接入协议的类型描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitnil" name:"Label"`
}

// Predefined struct for user
type DescribeGatewayProtocolRequestParams struct {

}

type DescribeGatewayProtocolRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeGatewayProtocolRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatewayProtocolRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGatewayProtocolRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGatewayProtocolResponseParams struct {
	// 返回数据
	Data []*DescribeGatewayProtocolData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeGatewayProtocolResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGatewayProtocolResponseParams `json:"Response"`
}

func (r *DescribeGatewayProtocolResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatewayProtocolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGatewayRequestParams struct {
	// 网关索引ID（从获取网关列表接口ListGateways中获取）
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`
}

type DescribeGatewayRequest struct {
	*tchttp.BaseRequest
	
	// 网关索引ID（从获取网关列表接口ListGateways中获取）
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`
}

func (r *DescribeGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGatewayResponseParams struct {
	// 返回数据
	Data *DescribeGatewayData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeGatewayResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGatewayResponseParams `json:"Response"`
}

func (r *DescribeGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewayVersion struct {
	// 服务名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 服务版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitnil" name:"Version"`

	// 服务最新版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestVersion *string `json:"LatestVersion,omitnil" name:"LatestVersion"`

	// 是否需要更新
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsUpdate *bool `json:"IsUpdate,omitnil" name:"IsUpdate"`

	// 升级信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpgradeInfo []*string `json:"UpgradeInfo,omitnil" name:"UpgradeInfo"`
}

type DescribeGatewayVersionData struct {
	// 网关服务列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Services []*DescribeGatewayVersion `json:"Services,omitnil" name:"Services"`
}

// Predefined struct for user
type DescribeGatewayVersionRequestParams struct {
	// 网关索引ID（从获取网关列表接口ListGateways中获取）
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`
}

type DescribeGatewayVersionRequest struct {
	*tchttp.BaseRequest
	
	// 网关索引ID（从获取网关列表接口ListGateways中获取）
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`
}

func (r *DescribeGatewayVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatewayVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGatewayVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGatewayVersionResponseParams struct {
	// 返回数据
	Data *DescribeGatewayVersionData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeGatewayVersionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGatewayVersionResponseParams `json:"Response"`
}

func (r *DescribeGatewayVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatewayVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationData struct {
	// 组织 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationId *string `json:"OrganizationId,omitnil" name:"OrganizationId"`

	// 组织名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 组织父节点 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentId *string `json:"ParentId,omitnil" name:"ParentId"`

	// 组织层级
	// 注意：此字段可能返回 null，表示取不到有效值。
	Level *int64 `json:"Level,omitnil" name:"Level"`

	// 用户id
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *int64 `json:"AppId,omitnil" name:"AppId"`

	// 组织结构
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentIds *string `json:"ParentIds,omitnil" name:"ParentIds"`

	// 设备总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 设备在线数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Online *int64 `json:"Online,omitnil" name:"Online"`
}

// Predefined struct for user
type DescribeOrganizationRequestParams struct {

}

type DescribeOrganizationRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeOrganizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationResponseParams struct {
	// 返回数据
	Data []*DescribeOrganizationData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeOrganizationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationResponseParams `json:"Response"`
}

func (r *DescribeOrganizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRecordBackupPlanData struct {
	// 录像上云计划ID
	PlanId *string `json:"PlanId,omitnil" name:"PlanId"`

	// 录像上云计划名称
	PlanName *string `json:"PlanName,omitnil" name:"PlanName"`

	// 录像上云模板ID
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 录像上云计划描述
	Describe *string `json:"Describe,omitnil" name:"Describe"`

	// 云文件生命周期
	LifeCycle *LifeCycleData `json:"LifeCycle,omitnil" name:"LifeCycle"`

	// 录像上云计划状态，1:正常使用中，0:删除中，无法使用
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 通道数量
	ChannelCount *int64 `json:"ChannelCount,omitnil" name:"ChannelCount"`

	// 创建时间
	CreateAt *string `json:"CreateAt,omitnil" name:"CreateAt"`

	// 修改时间
	UpdateAt *string `json:"UpdateAt,omitnil" name:"UpdateAt"`
}

// Predefined struct for user
type DescribeRecordBackupPlanRequestParams struct {
	// 录像上云计划ID（从查询录像上云计划列表接口ListRecordBackupPlans中获取）
	PlanId *string `json:"PlanId,omitnil" name:"PlanId"`
}

type DescribeRecordBackupPlanRequest struct {
	*tchttp.BaseRequest
	
	// 录像上云计划ID（从查询录像上云计划列表接口ListRecordBackupPlans中获取）
	PlanId *string `json:"PlanId,omitnil" name:"PlanId"`
}

func (r *DescribeRecordBackupPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordBackupPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordBackupPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordBackupPlanResponseParams struct {
	// 返回数据
	Data *DescribeRecordBackupPlanData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRecordBackupPlanResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordBackupPlanResponseParams `json:"Response"`
}

func (r *DescribeRecordBackupPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordBackupPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRecordBackupTemplateData struct {
	// 模板ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 模板名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateName *string `json:"TemplateName,omitnil" name:"TemplateName"`

	// 上云时间段（按周进行设置，支持一天设置多个时间段，每个时间段不小于10分钟）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeSections []*RecordTemplateTimeSections `json:"TimeSections,omitnil" name:"TimeSections"`

	// 录像时间段（按周进行设置，支持一天设置多个时间段，每个时间段不小于10分钟）
	// 注意：此字段可能返回 null，表示取不到有效值。
	DevTimeSections []*RecordTemplateTimeSections `json:"DevTimeSections,omitnil" name:"DevTimeSections"`

	// 上云倍速（支持1，2，4倍速）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Scale *int64 `json:"Scale,omitnil" name:"Scale"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateAt *string `json:"CreateAt,omitnil" name:"CreateAt"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateAt *string `json:"UpdateAt,omitnil" name:"UpdateAt"`
}

// Predefined struct for user
type DescribeRecordBackupTemplateRequestParams struct {
	// 模板ID（从查询录像上云模板列表接口ListRecordBackupTemplates中获取）
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`
}

type DescribeRecordBackupTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板ID（从查询录像上云模板列表接口ListRecordBackupTemplates中获取）
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`
}

func (r *DescribeRecordBackupTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordBackupTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordBackupTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordBackupTemplateResponseParams struct {
	// 返回数据
	Data *DescribeRecordBackupTemplateData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRecordBackupTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordBackupTemplateResponseParams `json:"Response"`
}

func (r *DescribeRecordBackupTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordBackupTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRecordFileData struct {
	// 提示类型，0:时间段内无归档录像，1:时间段内有归档录像
	Tips *int64 `json:"Tips,omitnil" name:"Tips"`

	// 存在为数组格式，不存在字段内容为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*RecordTimeLine `json:"List,omitnil" name:"List"`
}

// Predefined struct for user
type DescribeRecordFileRequestParams struct {
	// 通道所属设备ID
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`

	// 通道ID
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// 检索开始时间，UTC秒数，例如：1662114146，开始和结束时间段最长为一天，且不能跨天
	StartTime *uint64 `json:"StartTime,omitnil" name:"StartTime"`

	// 检索结束时间，UTC秒数，例如：1662114246，开始和结束时间段最长为一天，且不能跨天
	EndTime *uint64 `json:"EndTime,omitnil" name:"EndTime"`
}

type DescribeRecordFileRequest struct {
	*tchttp.BaseRequest
	
	// 通道所属设备ID
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`

	// 通道ID
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// 检索开始时间，UTC秒数，例如：1662114146，开始和结束时间段最长为一天，且不能跨天
	StartTime *uint64 `json:"StartTime,omitnil" name:"StartTime"`

	// 检索结束时间，UTC秒数，例如：1662114246，开始和结束时间段最长为一天，且不能跨天
	EndTime *uint64 `json:"EndTime,omitnil" name:"EndTime"`
}

func (r *DescribeRecordFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	delete(f, "ChannelId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordFileResponseParams struct {
	// 返回结果
	Data *DescribeRecordFileData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRecordFileResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordFileResponseParams `json:"Response"`
}

func (r *DescribeRecordFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordPlanRequestParams struct {
	// 实时上云计划ID
	PlanId *string `json:"PlanId,omitnil" name:"PlanId"`
}

type DescribeRecordPlanRequest struct {
	*tchttp.BaseRequest
	
	// 实时上云计划ID
	PlanId *string `json:"PlanId,omitnil" name:"PlanId"`
}

func (r *DescribeRecordPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordPlanResponseParams struct {
	// 返回结果
	Data *RecordPlanBaseInfo `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRecordPlanResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordPlanResponseParams `json:"Response"`
}

func (r *DescribeRecordPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordPlaybackUrlRequestParams struct {
	// 设备通道ID
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// 回放开始时间，UTC秒数，例如：1662114146，开始和结束时间段最长为一天，且不能跨天
	StartTime *uint64 `json:"StartTime,omitnil" name:"StartTime"`

	// 回放结束时间，UTC秒数，例如：1662114246，开始和结束时间段最长为一天，且不能跨天
	EndTime *uint64 `json:"EndTime,omitnil" name:"EndTime"`
}

type DescribeRecordPlaybackUrlRequest struct {
	*tchttp.BaseRequest
	
	// 设备通道ID
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// 回放开始时间，UTC秒数，例如：1662114146，开始和结束时间段最长为一天，且不能跨天
	StartTime *uint64 `json:"StartTime,omitnil" name:"StartTime"`

	// 回放结束时间，UTC秒数，例如：1662114246，开始和结束时间段最长为一天，且不能跨天
	EndTime *uint64 `json:"EndTime,omitnil" name:"EndTime"`
}

func (r *DescribeRecordPlaybackUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordPlaybackUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordPlaybackUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordPlaybackUrlResponseParams struct {
	// 返回结果
	Data *RecordPlaybackUrl `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRecordPlaybackUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordPlaybackUrlResponseParams `json:"Response"`
}

func (r *DescribeRecordPlaybackUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordPlaybackUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRecordRetrieveTaskData struct {
	// 取回任务ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 取回任务名称
	TaskName *string `json:"TaskName,omitnil" name:"TaskName"`

	// 取回录像的开始时间
	StartTime *uint64 `json:"StartTime,omitnil" name:"StartTime"`

	// 取回录像的结束时间
	EndTime *uint64 `json:"EndTime,omitnil" name:"EndTime"`

	// 取回模式，1:极速模式，其他暂不支持
	Mode *int64 `json:"Mode,omitnil" name:"Mode"`

	// 副本有效期
	Expiration *int64 `json:"Expiration,omitnil" name:"Expiration"`

	// 任务状态，0:已取回，1:取回中，2:待取回
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 取回容量，单位MB
	Capacity *float64 `json:"Capacity,omitnil" name:"Capacity"`

	// 任务的设备通道id
	Channels []*RecordRetrieveTaskChannelInfo `json:"Channels,omitnil" name:"Channels"`

	// 任务描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Describe *string `json:"Describe,omitnil" name:"Describe"`

	// 任务通道数量
	ChannelCount *int64 `json:"ChannelCount,omitnil" name:"ChannelCount"`
}

// Predefined struct for user
type DescribeRecordRetrieveTaskRequestParams struct {
	// 云录像取回任务ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

type DescribeRecordRetrieveTaskRequest struct {
	*tchttp.BaseRequest
	
	// 云录像取回任务ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

func (r *DescribeRecordRetrieveTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordRetrieveTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordRetrieveTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordRetrieveTaskResponseParams struct {
	// 返回结果
	Data *DescribeRecordRetrieveTaskData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRecordRetrieveTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordRetrieveTaskResponseParams `json:"Response"`
}

func (r *DescribeRecordRetrieveTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordRetrieveTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordSliceRequestParams struct {
	// 通道ID
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// 检索开始时间，UTC秒数，例如：1662114146，开始和结束时间段最长为一天，且不能跨天
	StartTime *uint64 `json:"StartTime,omitnil" name:"StartTime"`

	// 检索结束时间，UTC秒数，例如：1662114246，开始和结束时间段最长为一天，且不能跨天
	EndTime *uint64 `json:"EndTime,omitnil" name:"EndTime"`
}

type DescribeRecordSliceRequest struct {
	*tchttp.BaseRequest
	
	// 通道ID
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// 检索开始时间，UTC秒数，例如：1662114146，开始和结束时间段最长为一天，且不能跨天
	StartTime *uint64 `json:"StartTime,omitnil" name:"StartTime"`

	// 检索结束时间，UTC秒数，例如：1662114246，开始和结束时间段最长为一天，且不能跨天
	EndTime *uint64 `json:"EndTime,omitnil" name:"EndTime"`
}

func (r *DescribeRecordSliceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordSliceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordSliceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordSliceResponseParams struct {
	// 云录像切片信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*RecordSliceInfo `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRecordSliceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordSliceResponseParams `json:"Response"`
}

func (r *DescribeRecordSliceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordSliceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordTemplateRequestParams struct {
	// 模板ID
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`
}

type DescribeRecordTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板ID
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`
}

func (r *DescribeRecordTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordTemplateResponseParams struct {
	// 返回结果
	Data *RecordTemplateInfo `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRecordTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordTemplateResponseParams `json:"Response"`
}

func (r *DescribeRecordTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStreamAuthData struct {
	// 鉴权配置ID（uuid）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil" name:"Id"`

	// 是否开播放鉴权（1:开启,0:关闭）
	// 注意：此字段可能返回 null，表示取不到有效值。
	PullState *int64 `json:"PullState,omitnil" name:"PullState"`

	// 播放密钥（仅支持字母数字，长度0-10位）
	// 注意：此字段可能返回 null，表示取不到有效值。
	PullSecret *string `json:"PullSecret,omitnil" name:"PullSecret"`

	// 播放过期时间（单位：分钟）
	// 注意：此字段可能返回 null，表示取不到有效值。
	PullExpired *int64 `json:"PullExpired,omitnil" name:"PullExpired"`

	// 是否开启推流鉴权（1:开启,0:关闭）
	// 注意：此字段可能返回 null，表示取不到有效值。
	PushState *int64 `json:"PushState,omitnil" name:"PushState"`

	// 推流密钥（仅支持字母数字，长度0-10位）
	// 注意：此字段可能返回 null，表示取不到有效值。
	PushSecret *string `json:"PushSecret,omitnil" name:"PushSecret"`

	// 推流过期时间（单位：分钟）
	// 注意：此字段可能返回 null，表示取不到有效值。
	PushExpired *int64 `json:"PushExpired,omitnil" name:"PushExpired"`

	// 用户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *int64 `json:"AppId,omitnil" name:"AppId"`
}

// Predefined struct for user
type DescribeStreamAuthRequestParams struct {

}

type DescribeStreamAuthRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeStreamAuthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamAuthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamAuthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamAuthResponseParams struct {
	// 返回结果
	Data *DescribeStreamAuthData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeStreamAuthResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamAuthResponseParams `json:"Response"`
}

func (r *DescribeStreamAuthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamAuthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskRequestParams struct {
	// 简单任务或复杂任务ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

type DescribeTaskRequest struct {
	*tchttp.BaseRequest
	
	// 简单任务或复杂任务ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

func (r *DescribeTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskResponseParams struct {
	// 任务详情
	Data *TaskData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskResponseParams `json:"Response"`
}

func (r *DescribeTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserDeviceRequestParams struct {
	// 设备ID（从获取设备列表接口ListDevices中获取）
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`
}

type DescribeUserDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 设备ID（从获取设备列表接口ListDevices中获取）
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`
}

func (r *DescribeUserDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserDeviceResponseParams struct {
	// 返回结果
	Data *DescribeDeviceData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeUserDeviceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserDeviceResponseParams `json:"Response"`
}

func (r *DescribeUserDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVideoBitRateList struct {
	// 通道码率列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	BitRates []*BitRateInfo `json:"BitRates,omitnil" name:"BitRates"`
}

// Predefined struct for user
type DescribeVideoBitRateRequestParams struct {
	// 通道ID列表
	ChannelIds []*string `json:"ChannelIds,omitnil" name:"ChannelIds"`
}

type DescribeVideoBitRateRequest struct {
	*tchttp.BaseRequest
	
	// 通道ID列表
	ChannelIds []*string `json:"ChannelIds,omitnil" name:"ChannelIds"`
}

func (r *DescribeVideoBitRateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoBitRateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVideoBitRateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoBitRateResponseParams struct {
	// 无
	Data *DescribeVideoBitRateList `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeVideoBitRateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVideoBitRateResponseParams `json:"Response"`
}

func (r *DescribeVideoBitRateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoBitRateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVideoDownloadUrlData struct {
	// 录像文件下载 URL
	// 注意：
	// URL 有效期是10分钟，过期后将拒绝访问，若需再用请重新获取 
	// 录像文件下载采用分块传输编码，响应头Transfer-Encoding:chunked 
	// 下载文件命名格式为{ChannelId}-{BeginTime}-{EndTime}.{FileType} 
	Url *string `json:"Url,omitnil" name:"Url"`

	// 实际下载录像的开始时间
	// 注意：当请求中指定IsRespActualTime参数为true时，才有该字段
	ActualBeginTime *string `json:"ActualBeginTime,omitnil" name:"ActualBeginTime"`

	// 实际下载录像的结束时间
	// 注意：当请求中指定IsRespActualTime参数为true时，才有该字段
	ActualEndTime *string `json:"ActualEndTime,omitnil" name:"ActualEndTime"`
}

// Predefined struct for user
type DescribeVideoDownloadUrlRequestParams struct {
	// 通道 ID
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// 下载的开始时间，UTC 秒数，开始和结束时间段最长为30分钟，且不能跨天
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 下载的结束时间，UTC 秒数，开始和结束时间段最长为30分钟，且不能跨天
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 文件格式，"mp4"：mp4格式，"ts"：ts文件格式
	FileType *string `json:"FileType,omitnil" name:"FileType"`

	// 响应data中是否携带实际下载录像的开始时间与结束时间
	IsRespActualTime *bool `json:"IsRespActualTime,omitnil" name:"IsRespActualTime"`
}

type DescribeVideoDownloadUrlRequest struct {
	*tchttp.BaseRequest
	
	// 通道 ID
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// 下载的开始时间，UTC 秒数，开始和结束时间段最长为30分钟，且不能跨天
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 下载的结束时间，UTC 秒数，开始和结束时间段最长为30分钟，且不能跨天
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 文件格式，"mp4"：mp4格式，"ts"：ts文件格式
	FileType *string `json:"FileType,omitnil" name:"FileType"`

	// 响应data中是否携带实际下载录像的开始时间与结束时间
	IsRespActualTime *bool `json:"IsRespActualTime,omitnil" name:"IsRespActualTime"`
}

func (r *DescribeVideoDownloadUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoDownloadUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelId")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "FileType")
	delete(f, "IsRespActualTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVideoDownloadUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoDownloadUrlResponseParams struct {
	// 返回的数据结构
	Data *DescribeVideoDownloadUrlData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeVideoDownloadUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVideoDownloadUrlResponseParams `json:"Response"`
}

func (r *DescribeVideoDownloadUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoDownloadUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FaceMaskAIResultInfo struct {
	// 时间字符串
	Time *string `json:"Time,omitnil" name:"Time"`

	// 截图 URL
	Url *string `json:"Url,omitnil" name:"Url"`

	// 口罩信息
	FaceMaskInfo []*BaseAIResultInfo `json:"FaceMaskInfo,omitnil" name:"FaceMaskInfo"`
}

type GatewayDevice struct {
	// 设备ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`

	// 网关接入协议类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProtocolType *int64 `json:"ProtocolType,omitnil" name:"ProtocolType"`

	// 网关接入协议名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProtocolTypeName *string `json:"ProtocolTypeName,omitnil" name:"ProtocolTypeName"`

	// 设备名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 设备类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *int64 `json:"Type,omitnil" name:"Type"`

	// 设备内网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 设备端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitnil" name:"Port"`

	// 设备下通道数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelNum *int64 `json:"ChannelNum,omitnil" name:"ChannelNum"`

	// 设备状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil" name:"Status"`
}

type GatewayVersion struct {
	// 服务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 服务版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitnil" name:"Version"`
}

type GatewaysData struct {
	// 网关索引ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 网关编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	GwId *string `json:"GwId,omitnil" name:"GwId"`

	// 网关名称，仅支持中文、英文、数字、_、-，长度不超过32个字符
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 网关描述，仅支持中文、英文、数字、_、-，长度不超过128个字符
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 网关所属服务节点ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 网关所属服务节点名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitnil" name:"ClusterName"`

	// 网关所属地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil" name:"Region"`

	// 网关状态，0：离线，1:在线
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 网关激活时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitnil" name:"CreatedAt"`

	// 所属网关设备数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceNum *int64 `json:"DeviceNum,omitnil" name:"DeviceNum"`
}

type LifeCycleData struct {
	// 云文件热存储时长，单位天，最小1天，最大3650天
	Transition *int64 `json:"Transition,omitnil" name:"Transition"`

	// 云文件冷存储时长， 单位天，0表示不设置，设置时最小60天，Expiration字段加Transition字段不超过3650天
	Expiration *int64 `json:"Expiration,omitnil" name:"Expiration"`
}

type ListAITaskData struct {
	// AI任务列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*AITaskInfo `json:"List,omitnil" name:"List"`
}

// Predefined struct for user
type ListAITasksRequestParams struct {
	// 是否包含通道列表。"true"代表包含通道列表，"false"代表不包含通道列表，默认为 false
	IsContainChannelList *bool `json:"IsContainChannelList,omitnil" name:"IsContainChannelList"`

	// 是否包含AI配置。"true"代表包含任务配置，"false"代表不包含任务配置，默认为 false。
	IsContainTemplate *bool `json:"IsContainTemplate,omitnil" name:"IsContainTemplate"`

	// 页码。默认为1
	PageNumber *uint64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// 每页数量。可选值1～200，默认为20
	PageSize *uint64 `json:"PageSize,omitnil" name:"PageSize"`
}

type ListAITasksRequest struct {
	*tchttp.BaseRequest
	
	// 是否包含通道列表。"true"代表包含通道列表，"false"代表不包含通道列表，默认为 false
	IsContainChannelList *bool `json:"IsContainChannelList,omitnil" name:"IsContainChannelList"`

	// 是否包含AI配置。"true"代表包含任务配置，"false"代表不包含任务配置，默认为 false。
	IsContainTemplate *bool `json:"IsContainTemplate,omitnil" name:"IsContainTemplate"`

	// 页码。默认为1
	PageNumber *uint64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// 每页数量。可选值1～200，默认为20
	PageSize *uint64 `json:"PageSize,omitnil" name:"PageSize"`
}

func (r *ListAITasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAITasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IsContainChannelList")
	delete(f, "IsContainTemplate")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAITasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAITasksResponseParams struct {
	// AI 任务数量
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// AI任务列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *ListAITaskData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ListAITasksResponse struct {
	*tchttp.BaseResponse
	Response *ListAITasksResponseParams `json:"Response"`
}

func (r *ListAITasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAITasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListDeviceInfo struct {
	// 设备 ID
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`

	// 设备国标编码
	Code *string `json:"Code,omitnil" name:"Code"`

	// 设备状态。0:未注册，1:在线，2:离线，3:禁用
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 设备流传输协议。1:UDP,2:TCP
	TransportProtocol *uint64 `json:"TransportProtocol,omitnil" name:"TransportProtocol"`

	// 设备名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 设备类型。1:IPC,2:NVR
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// 设备密码
	Password *string `json:"Password,omitnil" name:"Password"`

	// 描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// 设备接入服务节点 ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 服务节点名称
	ClusterName *string `json:"ClusterName,omitnil" name:"ClusterName"`

	// 接入协议。1:RTMP,2:GB,3:GW
	AccessProtocol *uint64 `json:"AccessProtocol,omitnil" name:"AccessProtocol"`

	// 设备所属组织 ID
	OrganizationId *string `json:"OrganizationId,omitnil" name:"OrganizationId"`

	// 通道数量
	ChannelNum *uint64 `json:"ChannelNum,omitnil" name:"ChannelNum"`
}

// Predefined struct for user
type ListDevicesRequestParams struct {
	// 组织ID
	OrganizationId *string `json:"OrganizationId,omitnil" name:"OrganizationId"`

	// 是否获取当前层级及子层级的设备列表，默认false
	IsContainSubLevel *bool `json:"IsContainSubLevel,omitnil" name:"IsContainSubLevel"`

	// 设备接入协议。1:RTMP，2:GB，3:GW
	AccessProtocol *uint64 `json:"AccessProtocol,omitnil" name:"AccessProtocol"`

	// 设备类型。1:IPC，2:NVR
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// 设备状态。0:未注册，1:在线，2:离线，3:禁用	
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 服务节点ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 模糊搜索设备关键字
	Keyword *string `json:"Keyword,omitnil" name:"Keyword"`

	// 当前用户Uin
	CurrentUin *uint64 `json:"CurrentUin,omitnil" name:"CurrentUin"`

	// 页码，默认为1。
	PageNumber *uint64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// 每页数量，默认为20。
	PageSize *uint64 `json:"PageSize,omitnil" name:"PageSize"`
}

type ListDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 组织ID
	OrganizationId *string `json:"OrganizationId,omitnil" name:"OrganizationId"`

	// 是否获取当前层级及子层级的设备列表，默认false
	IsContainSubLevel *bool `json:"IsContainSubLevel,omitnil" name:"IsContainSubLevel"`

	// 设备接入协议。1:RTMP，2:GB，3:GW
	AccessProtocol *uint64 `json:"AccessProtocol,omitnil" name:"AccessProtocol"`

	// 设备类型。1:IPC，2:NVR
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// 设备状态。0:未注册，1:在线，2:离线，3:禁用	
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 服务节点ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 模糊搜索设备关键字
	Keyword *string `json:"Keyword,omitnil" name:"Keyword"`

	// 当前用户Uin
	CurrentUin *uint64 `json:"CurrentUin,omitnil" name:"CurrentUin"`

	// 页码，默认为1。
	PageNumber *uint64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// 每页数量，默认为20。
	PageSize *uint64 `json:"PageSize,omitnil" name:"PageSize"`
}

func (r *ListDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrganizationId")
	delete(f, "IsContainSubLevel")
	delete(f, "AccessProtocol")
	delete(f, "Type")
	delete(f, "Status")
	delete(f, "ClusterId")
	delete(f, "Keyword")
	delete(f, "CurrentUin")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDevicesResponseParams struct {
	// 设备列表详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*ListDeviceInfo `json:"Data,omitnil" name:"Data"`

	// 设备总数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ListDevicesResponse struct {
	*tchttp.BaseResponse
	Response *ListDevicesResponseParams `json:"Response"`
}

func (r *ListDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListGatewayDevicesData struct {
	// 网关下设备列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*GatewayDevice `json:"List,omitnil" name:"List"`

	// 网关下设备总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`
}

// Predefined struct for user
type ListGatewayDevicesRequestParams struct {
	// 网关索引ID（从获取网关列表接口ListGateways中获取）
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 分页页数
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// 分页大小
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`
}

type ListGatewayDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 网关索引ID（从获取网关列表接口ListGateways中获取）
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 分页页数
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// 分页大小
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`
}

func (r *ListGatewayDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListGatewayDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListGatewayDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListGatewayDevicesResponseParams struct {
	// 返回数据
	Data *ListGatewayDevicesData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ListGatewayDevicesResponse struct {
	*tchttp.BaseResponse
	Response *ListGatewayDevicesResponseParams `json:"Response"`
}

func (r *ListGatewayDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListGatewayDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListGatewaysData struct {
	// 网关列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*GatewaysData `json:"List,omitnil" name:"List"`

	// 网关数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`
}

// Predefined struct for user
type ListGatewaysRequestParams struct {
	// 页码，默认为1
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// 每页数量，默认为20
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 网关名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 服务节点ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 网关状态（0：离线，1 ：在线）
	Status *int64 `json:"Status,omitnil" name:"Status"`
}

type ListGatewaysRequest struct {
	*tchttp.BaseRequest
	
	// 页码，默认为1
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// 每页数量，默认为20
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 网关名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 服务节点ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 网关状态（0：离线，1 ：在线）
	Status *int64 `json:"Status,omitnil" name:"Status"`
}

func (r *ListGatewaysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListGatewaysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Name")
	delete(f, "ClusterId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListGatewaysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListGatewaysResponseParams struct {
	// 返回数据
	Data *ListGatewaysData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ListGatewaysResponse struct {
	*tchttp.BaseResponse
	Response *ListGatewaysResponseParams `json:"Response"`
}

func (r *ListGatewaysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListGatewaysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListOrganizationChannelNumbersData struct {
	// 组织下通道总数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 组织下未添加到计划的通道总数
	NotInPlanCount *int64 `json:"NotInPlanCount,omitnil" name:"NotInPlanCount"`
}

// Predefined struct for user
type ListOrganizationChannelNumbersRequestParams struct {
	// 组织ID，json数组格式，最多一次支持10个组织
	OrganizationId []*string `json:"OrganizationId,omitnil" name:"OrganizationId"`
}

type ListOrganizationChannelNumbersRequest struct {
	*tchttp.BaseRequest
	
	// 组织ID，json数组格式，最多一次支持10个组织
	OrganizationId []*string `json:"OrganizationId,omitnil" name:"OrganizationId"`
}

func (r *ListOrganizationChannelNumbersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListOrganizationChannelNumbersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrganizationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListOrganizationChannelNumbersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListOrganizationChannelNumbersResponseParams struct {
	// 返回结果
	Data *ListOrganizationChannelNumbersData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ListOrganizationChannelNumbersResponse struct {
	*tchttp.BaseResponse
	Response *ListOrganizationChannelNumbersResponseParams `json:"Response"`
}

func (r *ListOrganizationChannelNumbersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListOrganizationChannelNumbersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListOrganizationChannelsData struct {
	// 第几页
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// 当前页的设备数量
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 本次查询的设备通道总数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 设备通道信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*OrganizationChannelInfo `json:"List,omitnil" name:"List"`
}

// Predefined struct for user
type ListOrganizationChannelsRequestParams struct {
	// 组织ID
	OrganizationId *string `json:"OrganizationId,omitnil" name:"OrganizationId"`

	// 每页最大数量
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 第几页 
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// 查询条件，则按照设备名称查询
	// 查询条件同时只有一个生效。长度不超过32字节
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`

	// 查询条件，则按照通道名称查询
	// 查询条件同时只有一个生效。长度不超过32字节
	ChannelName *string `json:"ChannelName,omitnil" name:"ChannelName"`
}

type ListOrganizationChannelsRequest struct {
	*tchttp.BaseRequest
	
	// 组织ID
	OrganizationId *string `json:"OrganizationId,omitnil" name:"OrganizationId"`

	// 每页最大数量
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 第几页 
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// 查询条件，则按照设备名称查询
	// 查询条件同时只有一个生效。长度不超过32字节
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`

	// 查询条件，则按照通道名称查询
	// 查询条件同时只有一个生效。长度不超过32字节
	ChannelName *string `json:"ChannelName,omitnil" name:"ChannelName"`
}

func (r *ListOrganizationChannelsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListOrganizationChannelsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrganizationId")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "DeviceName")
	delete(f, "ChannelName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListOrganizationChannelsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListOrganizationChannelsResponseParams struct {
	// 返回结果
	Data *ListOrganizationChannelsData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ListOrganizationChannelsResponse struct {
	*tchttp.BaseResponse
	Response *ListOrganizationChannelsResponseParams `json:"Response"`
}

func (r *ListOrganizationChannelsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListOrganizationChannelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListRecordBackupPlanData struct {
	// 录像上云计划ID
	PlanId *string `json:"PlanId,omitnil" name:"PlanId"`

	// 录像上云计划名称
	PlanName *string `json:"PlanName,omitnil" name:"PlanName"`

	// 录像上云模板ID
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 录像上云计划描述
	Describe *string `json:"Describe,omitnil" name:"Describe"`

	// 云文件生命周期
	LifeCycle *LifeCycleData `json:"LifeCycle,omitnil" name:"LifeCycle"`

	// 录像上云计划状态，1:正常使用中，0:删除中，无法使用
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 通道数量
	ChannelCount *int64 `json:"ChannelCount,omitnil" name:"ChannelCount"`

	// 创建时间
	CreateAt *string `json:"CreateAt,omitnil" name:"CreateAt"`

	// 修改时间
	UpdateAt *string `json:"UpdateAt,omitnil" name:"UpdateAt"`
}

type ListRecordBackupPlanDevicesData struct {
	// 第几页
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// 当前页的设备数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 本次查询的设备通道总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 设备通道信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*RecordPlanChannelInfo `json:"List,omitnil" name:"List"`
}

// Predefined struct for user
type ListRecordBackupPlanDevicesRequestParams struct {
	// 录像计划ID（从查询录像上云计划列表接口ListRecordBackupPlans中获取）
	PlanId *string `json:"PlanId,omitnil" name:"PlanId"`

	// 按照设备名称查询（为空时，不参考该参数）
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`

	// 按照通道名称查询（为空时，不参考该参数）
	ChannelName *string `json:"ChannelName,omitnil" name:"ChannelName"`

	// 按照组织名称查询（为空时，不参考该参数）
	OrganizationName *string `json:"OrganizationName,omitnil" name:"OrganizationName"`

	// 每页最大数量
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 分页页数
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`
}

type ListRecordBackupPlanDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 录像计划ID（从查询录像上云计划列表接口ListRecordBackupPlans中获取）
	PlanId *string `json:"PlanId,omitnil" name:"PlanId"`

	// 按照设备名称查询（为空时，不参考该参数）
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`

	// 按照通道名称查询（为空时，不参考该参数）
	ChannelName *string `json:"ChannelName,omitnil" name:"ChannelName"`

	// 按照组织名称查询（为空时，不参考该参数）
	OrganizationName *string `json:"OrganizationName,omitnil" name:"OrganizationName"`

	// 每页最大数量
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 分页页数
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`
}

func (r *ListRecordBackupPlanDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRecordBackupPlanDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanId")
	delete(f, "DeviceName")
	delete(f, "ChannelName")
	delete(f, "OrganizationName")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListRecordBackupPlanDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRecordBackupPlanDevicesResponseParams struct {
	// 返回数据
	Data *ListRecordBackupPlanDevicesData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ListRecordBackupPlanDevicesResponse struct {
	*tchttp.BaseResponse
	Response *ListRecordBackupPlanDevicesResponseParams `json:"Response"`
}

func (r *ListRecordBackupPlanDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRecordBackupPlanDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRecordBackupPlansRequestParams struct {

}

type ListRecordBackupPlansRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ListRecordBackupPlansRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRecordBackupPlansRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListRecordBackupPlansRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRecordBackupPlansResponseParams struct {
	// 返回数据
	Data []*ListRecordBackupPlanData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ListRecordBackupPlansResponse struct {
	*tchttp.BaseResponse
	Response *ListRecordBackupPlansResponseParams `json:"Response"`
}

func (r *ListRecordBackupPlansResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRecordBackupPlansResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListRecordBackupTemplatesData struct {
	// 模板ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 模板名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateName *string `json:"TemplateName,omitnil" name:"TemplateName"`

	// 上云时间段（按周进行设置，支持一天设置多个时间段，每个时间段不小于10分钟）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeSections []*RecordTemplateTimeSections `json:"TimeSections,omitnil" name:"TimeSections"`

	// 录像时间段（按周进行设置，支持一天设置多个时间段，每个时间段不小于10分钟）
	// 注意：此字段可能返回 null，表示取不到有效值。
	DevTimeSections []*RecordTemplateTimeSections `json:"DevTimeSections,omitnil" name:"DevTimeSections"`

	// 上云倍速（支持1，2，4倍速）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Scale *int64 `json:"Scale,omitnil" name:"Scale"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateAt *string `json:"CreateAt,omitnil" name:"CreateAt"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateAt *string `json:"UpdateAt,omitnil" name:"UpdateAt"`
}

// Predefined struct for user
type ListRecordBackupTemplatesRequestParams struct {

}

type ListRecordBackupTemplatesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ListRecordBackupTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRecordBackupTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListRecordBackupTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRecordBackupTemplatesResponseParams struct {
	// 返回数据
	Data []*ListRecordBackupTemplatesData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ListRecordBackupTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *ListRecordBackupTemplatesResponseParams `json:"Response"`
}

func (r *ListRecordBackupTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRecordBackupTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListRecordPlanChannelsData struct {
	// 用户所有计划下通道id，存在通道是为数组格式，不存在时，字段数据为空 
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*string `json:"List,omitnil" name:"List"`
}

// Predefined struct for user
type ListRecordPlanChannelsRequestParams struct {

}

type ListRecordPlanChannelsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ListRecordPlanChannelsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRecordPlanChannelsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListRecordPlanChannelsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRecordPlanChannelsResponseParams struct {
	// 返回结果
	Data *ListRecordPlanChannelsData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ListRecordPlanChannelsResponse struct {
	*tchttp.BaseResponse
	Response *ListRecordPlanChannelsResponseParams `json:"Response"`
}

func (r *ListRecordPlanChannelsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRecordPlanChannelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListRecordPlanDevicesData struct {
	// 第几页
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// 当前页的设备数量
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 本次查询的设备通道总数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 设备通道信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*RecordPlanChannelInfo `json:"List,omitnil" name:"List"`
}

// Predefined struct for user
type ListRecordPlanDevicesRequestParams struct {
	// 上云计划ID
	PlanId *string `json:"PlanId,omitnil" name:"PlanId"`

	// 每页最大数量
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 第几页
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// 按照设备名称查询，为空时，不参考该参数
	// 通道名称、设备名称、组织名称同时只有一个有效，如果同时多个字段有值，按照通道名称、设备名称、组织名称的优先级顺序查询，如果都为空，则全量查询。长度不超过32字节
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`

	// 按照通道名称查询，为空时，不参考该参数
	// 通道名称、设备名称、组织名称同时只有一个有效，如果同时多个字段有值，按照通道名称、设备名称、组织名称的优先级顺序查询，如果都为空，则全量查询。长度不超过32字节
	ChannelName *string `json:"ChannelName,omitnil" name:"ChannelName"`

	// 按照组织名称查询|，为空时，不参考该参数
	// 通道名称、设备名称、组织名称同时只有一个有效，如果同时多个字段有值，按照通道名称、设备名称、组织名称的优先级顺序查询，如果都为空，则全量查询。长度不超过32字节
	OrganizationName *string `json:"OrganizationName,omitnil" name:"OrganizationName"`
}

type ListRecordPlanDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 上云计划ID
	PlanId *string `json:"PlanId,omitnil" name:"PlanId"`

	// 每页最大数量
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 第几页
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// 按照设备名称查询，为空时，不参考该参数
	// 通道名称、设备名称、组织名称同时只有一个有效，如果同时多个字段有值，按照通道名称、设备名称、组织名称的优先级顺序查询，如果都为空，则全量查询。长度不超过32字节
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`

	// 按照通道名称查询，为空时，不参考该参数
	// 通道名称、设备名称、组织名称同时只有一个有效，如果同时多个字段有值，按照通道名称、设备名称、组织名称的优先级顺序查询，如果都为空，则全量查询。长度不超过32字节
	ChannelName *string `json:"ChannelName,omitnil" name:"ChannelName"`

	// 按照组织名称查询|，为空时，不参考该参数
	// 通道名称、设备名称、组织名称同时只有一个有效，如果同时多个字段有值，按照通道名称、设备名称、组织名称的优先级顺序查询，如果都为空，则全量查询。长度不超过32字节
	OrganizationName *string `json:"OrganizationName,omitnil" name:"OrganizationName"`
}

func (r *ListRecordPlanDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRecordPlanDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanId")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "DeviceName")
	delete(f, "ChannelName")
	delete(f, "OrganizationName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListRecordPlanDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRecordPlanDevicesResponseParams struct {
	// 返回结果
	Data *ListRecordPlanDevicesData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ListRecordPlanDevicesResponse struct {
	*tchttp.BaseResponse
	Response *ListRecordPlanDevicesResponseParams `json:"Response"`
}

func (r *ListRecordPlanDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRecordPlanDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRecordPlansRequestParams struct {

}

type ListRecordPlansRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ListRecordPlansRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRecordPlansRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListRecordPlansRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRecordPlansResponseParams struct {
	// 返回结果，存在计划时，为Json数组格式，不存在计划时，字段数据为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*RecordPlanBaseInfo `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ListRecordPlansResponse struct {
	*tchttp.BaseResponse
	Response *ListRecordPlansResponseParams `json:"Response"`
}

func (r *ListRecordPlansResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRecordPlansResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRecordRetrieveTasksRequestParams struct {

}

type ListRecordRetrieveTasksRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ListRecordRetrieveTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRecordRetrieveTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListRecordRetrieveTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRecordRetrieveTasksResponseParams struct {
	// 返回结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*RecordRetrieveTaskDetailsInfo `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ListRecordRetrieveTasksResponse struct {
	*tchttp.BaseResponse
	Response *ListRecordRetrieveTasksResponseParams `json:"Response"`
}

func (r *ListRecordRetrieveTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRecordRetrieveTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRecordTemplatesRequestParams struct {

}

type ListRecordTemplatesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ListRecordTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRecordTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListRecordTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRecordTemplatesResponseParams struct {
	// 返回结果，存在模板时，为Json数组格式，不存在模板时，字段数据为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*RecordTemplateInfo `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ListRecordTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *ListRecordTemplatesResponseParams `json:"Response"`
}

func (r *ListRecordTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRecordTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSubTasksData struct {
	// 子任务列表
	List []*SubTaskData `json:"List,omitnil" name:"List"`

	// 子任务数量
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`
}

// Predefined struct for user
type ListSubTasksRequestParams struct {
	// 复杂任务ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 页码，默认为1
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// 每页数量，默认为10
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 默认不对该字段进行筛选，否则根据任务状态进行筛选。状态码：1-NEW，2-RUNNING，3-COMPLETED，4-FAILED
	Status *int64 `json:"Status,omitnil" name:"Status"`
}

type ListSubTasksRequest struct {
	*tchttp.BaseRequest
	
	// 复杂任务ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 页码，默认为1
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// 每页数量，默认为10
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 默认不对该字段进行筛选，否则根据任务状态进行筛选。状态码：1-NEW，2-RUNNING，3-COMPLETED，4-FAILED
	Status *int64 `json:"Status,omitnil" name:"Status"`
}

func (r *ListSubTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSubTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListSubTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListSubTasksResponseParams struct {
	// 返回数据
	Data *ListSubTasksData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ListSubTasksResponse struct {
	*tchttp.BaseResponse
	Response *ListSubTasksResponseParams `json:"Response"`
}

func (r *ListSubTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSubTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListTasksData struct {
	// 任务列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*TaskData `json:"List,omitnil" name:"List"`

	// 任务数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`
}

// Predefined struct for user
type ListTasksRequestParams struct {
	// 页码，默认为1
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// 每页数量，默认为10
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 默认不根据该字段进行筛选，否则根据设备操作类型进行筛选，对应任务的Action字段，批量任务操作类型以Batch开头。目前值有：BatchDeleteUserDevice，BatchDisableDevice，BatchEnableDevice，DeleteUserDevice，DisableDevice，EnableDevice
	Operation *string `json:"Operation,omitnil" name:"Operation"`

	// 默认不根据该字段进行筛选，否则根据任务状态进行筛选。状态码：1-NEW，2-RUNNING，3-COMPLETED，4-FAILED
	Status *int64 `json:"Status,omitnil" name:"Status"`
}

type ListTasksRequest struct {
	*tchttp.BaseRequest
	
	// 页码，默认为1
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// 每页数量，默认为10
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 默认不根据该字段进行筛选，否则根据设备操作类型进行筛选，对应任务的Action字段，批量任务操作类型以Batch开头。目前值有：BatchDeleteUserDevice，BatchDisableDevice，BatchEnableDevice，DeleteUserDevice，DisableDevice，EnableDevice
	Operation *string `json:"Operation,omitnil" name:"Operation"`

	// 默认不根据该字段进行筛选，否则根据任务状态进行筛选。状态码：1-NEW，2-RUNNING，3-COMPLETED，4-FAILED
	Status *int64 `json:"Status,omitnil" name:"Status"`
}

func (r *ListTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Operation")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTasksResponseParams struct {
	// 返回数据
	Data *ListTasksData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ListTasksResponse struct {
	*tchttp.BaseResponse
	Response *ListTasksResponseParams `json:"Response"`
}

func (r *ListTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Location struct {
	// 左上角 X 坐标轴
	X *int64 `json:"X,omitnil" name:"X"`

	// 左上角 Y 坐标轴
	Y *int64 `json:"Y,omitnil" name:"Y"`

	// 方框宽
	Width *uint64 `json:"Width,omitnil" name:"Width"`

	// 方框高
	Height *uint64 `json:"Height,omitnil" name:"Height"`
}

type OperTimeSlot struct {
	// 开始时间。格式为"hh:mm:ss"，且 Start 必须小于 End
	Start *string `json:"Start,omitnil" name:"Start"`

	// 结束时间。格式为"hh:mm:ss"，且 Start 必须小于 End
	End *string `json:"End,omitnil" name:"End"`
}

type OrganizationChannelInfo struct {
	// 设备通道所属的设备ID
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`

	// 设备通道所属的设备名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`

	// 设备通道ID
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// 设备通道名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelName *string `json:"ChannelName,omitnil" name:"ChannelName"`

	// 该通道是否在上云计划中，如果是，则不能在添加到其他上云计划|true：在上云计划中，false：不在上云计划中
	InPlan *bool `json:"InPlan,omitnil" name:"InPlan"`
}

type PetAIResultInfo struct {
	// 时间字符串
	Time *string `json:"Time,omitnil" name:"Time"`

	// 截图 URL
	Url *string `json:"Url,omitnil" name:"Url"`

	// 宠物信息
	PetInfo []*BaseAIResultInfo `json:"PetInfo,omitnil" name:"PetInfo"`
}

type PhoneCallAIResultInfo struct {
	// 时间字符串
	Time *string `json:"Time,omitnil" name:"Time"`

	// 截图 URL
	Url *string `json:"Url,omitnil" name:"Url"`

	// 打电话信息
	PhoneCallInfo []*BaseAIResultInfo `json:"PhoneCallInfo,omitnil" name:"PhoneCallInfo"`
}

type PlateContent struct {
	// 车牌号信息
	Plate *string `json:"Plate,omitnil" name:"Plate"`

	// 车牌的颜色
	Color *string `json:"Color,omitnil" name:"Color"`

	// 车牌的种类，例如普通蓝牌
	Type *string `json:"Type,omitnil" name:"Type"`

	// 截图中坐标信息
	Location *Location `json:"Location,omitnil" name:"Location"`
}

type PlayRecordData struct {
	// 录像播放地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Flv *string `json:"Flv,omitnil" name:"Flv"`
}

// Predefined struct for user
type PlayRecordRequestParams struct {
	// 通道 ID（从查询通道DescribeDeviceChannel接口中获取）
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// 起始时间
	Start *int64 `json:"Start,omitnil" name:"Start"`

	//  结束时间
	End *int64 `json:"End,omitnil" name:"End"`

	// 流类型（1:主码流；2:子码流（不可以和 Resolution 同时下发））
	StreamType *int64 `json:"StreamType,omitnil" name:"StreamType"`

	// 分辨率（1:QCIF；2:CIF； 3:4CIF； 4:D1； 5:720P； 6:1080P/I； 自定义的19201080等等（需设备支持）（不可以和 StreamType 同时下发））
	Resolution *string `json:"Resolution,omitnil" name:"Resolution"`
}

type PlayRecordRequest struct {
	*tchttp.BaseRequest
	
	// 通道 ID（从查询通道DescribeDeviceChannel接口中获取）
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// 起始时间
	Start *int64 `json:"Start,omitnil" name:"Start"`

	//  结束时间
	End *int64 `json:"End,omitnil" name:"End"`

	// 流类型（1:主码流；2:子码流（不可以和 Resolution 同时下发））
	StreamType *int64 `json:"StreamType,omitnil" name:"StreamType"`

	// 分辨率（1:QCIF；2:CIF； 3:4CIF； 4:D1； 5:720P； 6:1080P/I； 自定义的19201080等等（需设备支持）（不可以和 StreamType 同时下发））
	Resolution *string `json:"Resolution,omitnil" name:"Resolution"`
}

func (r *PlayRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PlayRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelId")
	delete(f, "Start")
	delete(f, "End")
	delete(f, "StreamType")
	delete(f, "Resolution")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PlayRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PlayRecordResponseParams struct {
	// 返回结果
	Data *PlayRecordData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type PlayRecordResponse struct {
	*tchttp.BaseResponse
	Response *PlayRecordResponseParams `json:"Response"`
}

func (r *PlayRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PlayRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecordPlanBaseInfo struct {
	// 上云计划ID
	PlanId *string `json:"PlanId,omitnil" name:"PlanId"`

	// 上云计划名称
	PlanName *string `json:"PlanName,omitnil" name:"PlanName"`

	// 上云模板ID
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 上云计划描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Describe *string `json:"Describe,omitnil" name:"Describe"`

	// 码流类型，default:设备默认码流类型，main:主码流，sub:子码流，其他根据设备能力集自定义
	// 注意：此字段可能返回 null，表示取不到有效值。
	StreamType *string `json:"StreamType,omitnil" name:"StreamType"`

	// 云文件生命周期
	LifeCycle *LifeCycleData `json:"LifeCycle,omitnil" name:"LifeCycle"`

	// 录像计划状态，1:正常使用中，0:删除中，无法使用
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 通道总数
	ChannelCount *int64 `json:"ChannelCount,omitnil" name:"ChannelCount"`
}

type RecordPlanChannelInfo struct {
	// 设备通道所属的设备ID
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`

	// 设备通道所属的设备名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`

	// 设备通道ID
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// 设备通道名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelName *string `json:"ChannelName,omitnil" name:"ChannelName"`

	// 所属组织名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationName *string `json:"OrganizationName,omitnil" name:"OrganizationName"`
}

type RecordPlanOptData struct {
	// 上云计划ID
	PlanId *string `json:"PlanId,omitnil" name:"PlanId"`

	// 上云计划名称
	PlanName *string `json:"PlanName,omitnil" name:"PlanName"`

	// 上云模板ID
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 上云计划描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Describe *string `json:"Describe,omitnil" name:"Describe"`

	// 云文件生命周期
	LifeCycle *LifeCycleData `json:"LifeCycle,omitnil" name:"LifeCycle"`

	// 码流类型，default:设备默认码流类型，main:主码流，sub:子码流，其他根据设备能力集自定义
	// 注意：此字段可能返回 null，表示取不到有效值。
	StreamType *string `json:"StreamType,omitnil" name:"StreamType"`
}

type RecordPlaybackUrl struct {
	// hls回放url
	Hls *string `json:"Hls,omitnil" name:"Hls"`
}

type RecordRetrieveTaskChannelInfo struct {
	// 设备通道所属的设备ID
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`

	// 设备通道所属的设备名称
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`

	// 设备通道ID
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// 设备通道名称 
	ChannelName *string `json:"ChannelName,omitnil" name:"ChannelName"`

	// 任务状态，0:已取回，1:取回中，2:待取回, 3:无归档录像
	Status *int64 `json:"Status,omitnil" name:"Status"`
}

type RecordRetrieveTaskDetailsInfo struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 任务名称
	TaskName *string `json:"TaskName,omitnil" name:"TaskName"`

	// 取回录像的开始时间 
	StartTime *uint64 `json:"StartTime,omitnil" name:"StartTime"`

	// 取回录像的结束时间
	EndTime *uint64 `json:"EndTime,omitnil" name:"EndTime"`

	// 取回模式，1:极速模式，其他暂不支持
	Mode *int64 `json:"Mode,omitnil" name:"Mode"`

	// 副本有效期
	Expiration *int64 `json:"Expiration,omitnil" name:"Expiration"`

	// 任务状态， 0:已取回，1:取回中，2:待取回
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 取回容量，单位MB
	Capacity *int64 `json:"Capacity,omitnil" name:"Capacity"`

	// 任务描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Describe *string `json:"Describe,omitnil" name:"Describe"`

	// 任务通道数量
	ChannelCount *int64 `json:"ChannelCount,omitnil" name:"ChannelCount"`
}

type RecordSliceInfo struct {
	// 计划ID
	PlanId *string `json:"PlanId,omitnil" name:"PlanId"`

	// 录像切片开始和结束时间列表
	List []*RecordTimeLine `json:"List,omitnil" name:"List"`
}

type RecordTemplateInfo struct {
	// 模板ID
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 模板名称
	TemplateName *string `json:"TemplateName,omitnil" name:"TemplateName"`

	// 上云时间段，按周进行设置，支持一天设置多个时间段，每个时间段不小于10分钟
	TimeSections []*RecordTemplateTimeSections `json:"TimeSections,omitnil" name:"TimeSections"`
}

type RecordTemplateTimeSections struct {
	// 周日期，取值范围1～7（对应周一～周日
	DayOfWeek *int64 `json:"DayOfWeek,omitnil" name:"DayOfWeek"`

	// 开始时间，格式：HH:MM:SS，范围：[00:00:00～23:59:59]
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间，格式：HH:MM:SS，范围：[00:00:00～23:59:59] 
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

type RecordTimeLine struct {
	// 时间片段开始时间，UTC秒数，例如：1662114146
	Begin *uint64 `json:"Begin,omitnil" name:"Begin"`

	// 时间片段结束时间，UTC秒数，例如：1662114146
	End *uint64 `json:"End,omitnil" name:"End"`
}

// Predefined struct for user
type RefreshDeviceChannelRequestParams struct {
	// 设备 ID（从获取设备列表ListDevices接口中获取）
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`
}

type RefreshDeviceChannelRequest struct {
	*tchttp.BaseRequest
	
	// 设备 ID（从获取设备列表ListDevices接口中获取）
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`
}

func (r *RefreshDeviceChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefreshDeviceChannelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RefreshDeviceChannelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RefreshDeviceChannelResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RefreshDeviceChannelResponse struct {
	*tchttp.BaseResponse
	Response *RefreshDeviceChannelResponseParams `json:"Response"`
}

func (r *RefreshDeviceChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefreshDeviceChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SmokingAIResultInfo struct {
	// 时间字符串
	Time *string `json:"Time,omitnil" name:"Time"`

	// 截图 URL
	Url *string `json:"Url,omitnil" name:"Url"`

	// 抽烟信息
	SmokingInfo []*BaseAIResultInfo `json:"SmokingInfo,omitnil" name:"SmokingInfo"`
}

type SnapshotConfig struct {
	// 截图频率。可选值1～20秒
	TimeInterval *uint64 `json:"TimeInterval,omitnil" name:"TimeInterval"`

	// 模板生效的时间段。最多包含5组时间段
	OperTimeSlot []*OperTimeSlot `json:"OperTimeSlot,omitnil" name:"OperTimeSlot"`
}

type SubTaskData struct {
	// 子任务ID
	SubTaskId *string `json:"SubTaskId,omitnil" name:"SubTaskId"`

	// 任务状态1:NEW,2:RUNNING,3:COMPLETED ,4:FAILED
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 任务失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailReason *string `json:"FailReason,omitnil" name:"FailReason"`

	// 任务进度
	Progress *float64 `json:"Progress,omitnil" name:"Progress"`

	// 操作类型
	Action *string `json:"Action,omitnil" name:"Action"`

	// 操作类型中文描述
	ActionZhDesc *string `json:"ActionZhDesc,omitnil" name:"ActionZhDesc"`

	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 启动任务时间
	StartedAt *string `json:"StartedAt,omitnil" name:"StartedAt"`

	// 创建任务时间
	CreatedAt *string `json:"CreatedAt,omitnil" name:"CreatedAt"`

	// 更新任务时间
	UpdatedAt *string `json:"UpdatedAt,omitnil" name:"UpdatedAt"`

	// 任务运行时间，单位ms
	Runtime *int64 `json:"Runtime,omitnil" name:"Runtime"`
}

type TaskData struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 任务状态1:NEW,2:RUNNING,3:COMPLETED ,4:FAILED
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailReason *string `json:"FailReason,omitnil" name:"FailReason"`

	// 进度（0-1）
	Progress *float64 `json:"Progress,omitnil" name:"Progress"`

	// 任务操作类型，批量任务类型以Batch开头
	Action *string `json:"Action,omitnil" name:"Action"`

	// 操作类型中文描述
	ActionZhDesc *string `json:"ActionZhDesc,omitnil" name:"ActionZhDesc"`

	// 任务类型 1.简单 2.复杂 3.子任务
	TaskType *int64 `json:"TaskType,omitnil" name:"TaskType"`

	// 任务资源id（复杂任务该字段无效）
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 总任务数（仅复杂任务有效）
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 成功任务数（仅复杂任务有效）
	SuccessCount *int64 `json:"SuccessCount,omitnil" name:"SuccessCount"`

	// 失败任务数（仅复杂任务有效）
	FailCount *int64 `json:"FailCount,omitnil" name:"FailCount"`

	// 运行任务数（仅复杂任务有效）
	RunningCount *int64 `json:"RunningCount,omitnil" name:"RunningCount"`

	// 启动任务时间
	StartedAt *string `json:"StartedAt,omitnil" name:"StartedAt"`

	// 创建任务时间
	CreatedAt *string `json:"CreatedAt,omitnil" name:"CreatedAt"`

	// 更新任务时间
	UpdatedAt *string `json:"UpdatedAt,omitnil" name:"UpdatedAt"`

	// 任务运行时间，单位ms
	Runtime *int64 `json:"Runtime,omitnil" name:"Runtime"`
}

type Timeline struct {
	// 分片起始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Begin *int64 `json:"Begin,omitnil" name:"Begin"`

	// 分片结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	End *int64 `json:"End,omitnil" name:"End"`
}

// Predefined struct for user
type UpdateAITaskRequestParams struct {
	// AI 任务 ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// AI 任务名称。仅支持中文、英文、数字、_、-，长度不超过32个字符
	Name *string `json:"Name,omitnil" name:"Name"`

	// AI 任务描述。仅支持中文、英文、数字、_、-，长度不超过128个字符
	Desc *string `json:"Desc,omitnil" name:"Desc"`

	// 通道 ID 列表。不能添加存在于其他 AI 任务的通道，限制1000个通道。
	ChannelList []*string `json:"ChannelList,omitnil" name:"ChannelList"`

	// AI 结果回调地址。类似 "http://ip:port/xxx或者https://domain/xxx
	CallbackUrl *string `json:"CallbackUrl,omitnil" name:"CallbackUrl"`

	// 是否立即开启 AI 任务。"true"代表立即开启 AI 任务，"false"代表暂不开启 AI 任务，默认为 false。
	IsStartTheTask *bool `json:"IsStartTheTask,omitnil" name:"IsStartTheTask"`

	// AI 配置列表
	Templates []*AITemplates `json:"Templates,omitnil" name:"Templates"`
}

type UpdateAITaskRequest struct {
	*tchttp.BaseRequest
	
	// AI 任务 ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// AI 任务名称。仅支持中文、英文、数字、_、-，长度不超过32个字符
	Name *string `json:"Name,omitnil" name:"Name"`

	// AI 任务描述。仅支持中文、英文、数字、_、-，长度不超过128个字符
	Desc *string `json:"Desc,omitnil" name:"Desc"`

	// 通道 ID 列表。不能添加存在于其他 AI 任务的通道，限制1000个通道。
	ChannelList []*string `json:"ChannelList,omitnil" name:"ChannelList"`

	// AI 结果回调地址。类似 "http://ip:port/xxx或者https://domain/xxx
	CallbackUrl *string `json:"CallbackUrl,omitnil" name:"CallbackUrl"`

	// 是否立即开启 AI 任务。"true"代表立即开启 AI 任务，"false"代表暂不开启 AI 任务，默认为 false。
	IsStartTheTask *bool `json:"IsStartTheTask,omitnil" name:"IsStartTheTask"`

	// AI 配置列表
	Templates []*AITemplates `json:"Templates,omitnil" name:"Templates"`
}

func (r *UpdateAITaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAITaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "Name")
	delete(f, "Desc")
	delete(f, "ChannelList")
	delete(f, "CallbackUrl")
	delete(f, "IsStartTheTask")
	delete(f, "Templates")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAITaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAITaskResponseParams struct {
	// AI任务信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *AITaskInfo `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpdateAITaskResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAITaskResponseParams `json:"Response"`
}

func (r *UpdateAITaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAITaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAITaskStatusRequestParams struct {
	// AI 任务 ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// AI 任务状态。"on"代表开启了 AI 分析任务，"off"代表停止AI分析任务
	Status *string `json:"Status,omitnil" name:"Status"`
}

type UpdateAITaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// AI 任务 ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// AI 任务状态。"on"代表开启了 AI 分析任务，"off"代表停止AI分析任务
	Status *string `json:"Status,omitnil" name:"Status"`
}

func (r *UpdateAITaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAITaskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAITaskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAITaskStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpdateAITaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAITaskStatusResponseParams `json:"Response"`
}

func (r *UpdateAITaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAITaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDeviceData struct {
	// 设备ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`

	// 设备编码（即我们为设备生成的20位国标编码）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *string `json:"Code,omitnil" name:"Code"`

	// 设备名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 设备接入协议，1:RTMP,2:GB,3:GW 
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessProtocol *int64 `json:"AccessProtocol,omitnil" name:"AccessProtocol"`

	// 设备类型，1:IPC,2:NVR
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *int64 `json:"Type,omitnil" name:"Type"`

	// 设备接入服务节点ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 设备接入服务节点名称
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitnil" name:"ClusterName"`

	// 设备流传输协议，1:UDP,2:TCP 
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransportProtocol *int64 `json:"TransportProtocol,omitnil" name:"TransportProtocol"`

	// 设备密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitnil" name:"Password"`

	// 设备描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 设备状态，0:未注册,1:在线,2:离线,3:禁用
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 设备所属组织ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationId *int64 `json:"OrganizationId,omitnil" name:"OrganizationId"`

	// 设备接入网关ID，从查询网关列表接口中获取（仅网关接入需要）
	// 注意：此字段可能返回 null，表示取不到有效值。
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 网关接入协议类型，1.海康SDK，2.大华SDK，3.宇视SDK，4.Onvif（仅网关接入需要）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProtocolType *int64 `json:"ProtocolType,omitnil" name:"ProtocolType"`

	// 设备接入IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 设备Port
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitnil" name:"Port"`

	// 设备用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Username *string `json:"Username,omitnil" name:"Username"`

	// 用户Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *int64 `json:"AppId,omitnil" name:"AppId"`
}

// Predefined struct for user
type UpdateDeviceOrganizationRequestParams struct {
	// 设备 ID 数组（从获取设备列表接口ListDevices中获取）
	DeviceIds []*string `json:"DeviceIds,omitnil" name:"DeviceIds"`

	// 组织 ID（从查询组织接口DescribeOrganization中获取）
	OrganizationId *string `json:"OrganizationId,omitnil" name:"OrganizationId"`
}

type UpdateDeviceOrganizationRequest struct {
	*tchttp.BaseRequest
	
	// 设备 ID 数组（从获取设备列表接口ListDevices中获取）
	DeviceIds []*string `json:"DeviceIds,omitnil" name:"DeviceIds"`

	// 组织 ID（从查询组织接口DescribeOrganization中获取）
	OrganizationId *string `json:"OrganizationId,omitnil" name:"OrganizationId"`
}

func (r *UpdateDeviceOrganizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDeviceOrganizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceIds")
	delete(f, "OrganizationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateDeviceOrganizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDeviceOrganizationResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpdateDeviceOrganizationResponse struct {
	*tchttp.BaseResponse
	Response *UpdateDeviceOrganizationResponseParams `json:"Response"`
}

func (r *UpdateDeviceOrganizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDeviceOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDeviceStatusRequestParams struct {
	// 设备 ID（从获取设备列表接口ListDevices中获取）
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`

	// 禁用启用状态码（2：启用，3:禁用）
	Status *int64 `json:"Status,omitnil" name:"Status"`
}

type UpdateDeviceStatusRequest struct {
	*tchttp.BaseRequest
	
	// 设备 ID（从获取设备列表接口ListDevices中获取）
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`

	// 禁用启用状态码（2：启用，3:禁用）
	Status *int64 `json:"Status,omitnil" name:"Status"`
}

func (r *UpdateDeviceStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDeviceStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateDeviceStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDeviceStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpdateDeviceStatusResponse struct {
	*tchttp.BaseResponse
	Response *UpdateDeviceStatusResponseParams `json:"Response"`
}

func (r *UpdateDeviceStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDeviceStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateGatewayData struct {
	// 网关索引ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 网关编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	GwId *string `json:"GwId,omitnil" name:"GwId"`

	// 网关名称，仅支持中文、英文、数字、_、-，长度不超过32个字符
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 网关描述，仅支持中文、英文、数字、_、-，长度不超过128个字符
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 服务节点ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 服务节点名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitnil" name:"ClusterName"`

	// 网关状态，0：离线，1:在线
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 激活时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *int64 `json:"CreatedAt,omitnil" name:"CreatedAt"`

	// 网关密钥
	// 注意：此字段可能返回 null，表示取不到有效值。
	Secret *string `json:"Secret,omitnil" name:"Secret"`

	// 网关版本信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitnil" name:"Version"`
}

// Predefined struct for user
type UpdateGatewayRequestParams struct {
	// 网关索引ID（从获取网关列表ListGateways接口中获取）	
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 仅支持中文、英文、数网关名称，字、_、-，长度不超过32个字符
	Name *string `json:"Name,omitnil" name:"Name"`

	// 网关描述，仅支持中文、英文、数字、_、-，长度不超过128个字符
	Description *string `json:"Description,omitnil" name:"Description"`
}

type UpdateGatewayRequest struct {
	*tchttp.BaseRequest
	
	// 网关索引ID（从获取网关列表ListGateways接口中获取）	
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 仅支持中文、英文、数网关名称，字、_、-，长度不超过32个字符
	Name *string `json:"Name,omitnil" name:"Name"`

	// 网关描述，仅支持中文、英文、数字、_、-，长度不超过128个字符
	Description *string `json:"Description,omitnil" name:"Description"`
}

func (r *UpdateGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Name")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateGatewayResponseParams struct {
	// 返回数据
	Data *UpdateGatewayData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpdateGatewayResponse struct {
	*tchttp.BaseResponse
	Response *UpdateGatewayResponseParams `json:"Response"`
}

func (r *UpdateGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOrgData struct {
	// 组织 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationId *string `json:"OrganizationId,omitnil" name:"OrganizationId"`

	// 组织名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 组织父节点 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentId *string `json:"ParentId,omitnil" name:"ParentId"`

	// 组织层级
	// 注意：此字段可能返回 null，表示取不到有效值。
	Level *int64 `json:"Level,omitnil" name:"Level"`

	// 用户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *int64 `json:"AppId,omitnil" name:"AppId"`

	// 组织结构
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentIds *string `json:"ParentIds,omitnil" name:"ParentIds"`

	// 设备总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 设备在线数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Online *int64 `json:"Online,omitnil" name:"Online"`
}

// Predefined struct for user
type UpdateOrganizationRequestParams struct {
	// 组织ID（从查询组织接口DescribeOrganization中获取）
	OrganizationId *string `json:"OrganizationId,omitnil" name:"OrganizationId"`

	// 组织名称
	Name *string `json:"Name,omitnil" name:"Name"`
}

type UpdateOrganizationRequest struct {
	*tchttp.BaseRequest
	
	// 组织ID（从查询组织接口DescribeOrganization中获取）
	OrganizationId *string `json:"OrganizationId,omitnil" name:"OrganizationId"`

	// 组织名称
	Name *string `json:"Name,omitnil" name:"Name"`
}

func (r *UpdateOrganizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOrganizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrganizationId")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateOrganizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateOrganizationResponseParams struct {
	// 返回结果
	Data *UpdateOrgData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpdateOrganizationResponse struct {
	*tchttp.BaseResponse
	Response *UpdateOrganizationResponseParams `json:"Response"`
}

func (r *UpdateOrganizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateRecordBackupPlanData struct {
	// 录像上云计划ID
	PlanId *string `json:"PlanId,omitnil" name:"PlanId"`

	// 录像上云计划名称
	PlanName *string `json:"PlanName,omitnil" name:"PlanName"`

	// 录像上云模板ID
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 录像上云计划描述
	Describe *string `json:"Describe,omitnil" name:"Describe"`

	// 云文件生命周期
	LifeCycle *LifeCycleData `json:"LifeCycle,omitnil" name:"LifeCycle"`

	// 录像上云计划状态，1:正常使用中，0:删除中，无法使用
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 通道数量
	ChannelCount *int64 `json:"ChannelCount,omitnil" name:"ChannelCount"`

	// 创建时间
	CreateAt *string `json:"CreateAt,omitnil" name:"CreateAt"`

	// 修改时间
	UpdateAt *string `json:"UpdateAt,omitnil" name:"UpdateAt"`
}

type UpdateRecordBackupPlanModify struct {
	// 录像计划名称（仅支持中文、英文、数字、_、-，长度不超过32个字符，计划名称全局唯一，不能为空，不能重复，不修改名称时，不需要该字段）
	PlanName *string `json:"PlanName,omitnil" name:"PlanName"`

	// 录制模板ID（从查询录像上云模板列表接口ListRecordBackupTemplates中获取，不修改模板ID时，不需要该字段）
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 录像计划描述（仅支持中文、英文、数字、_、-，长度不超过128个字符， 不修改描述时，不需要该字段）
	Describe *string `json:"Describe,omitnil" name:"Describe"`

	// 生命周期（录像文件生命周期设置，管理文件冷、热存储的时间，不修改生命周期时，不需要该字段）
	LifeCycle *LifeCycleData `json:"LifeCycle,omitnil" name:"LifeCycle"`

	// 要新增的设备通道（Json数组，没有新增时，不需要该字段，一次添加通道总数不超过5000个，包括组织目录下的通道数量）
	Add []*ChannelInfo `json:"Add,omitnil" name:"Add"`

	// 要删除的设备通道（Json数组，内容为要删除的设备通道id，没有删除设备通道时，不需要该字段）
	Del []*string `json:"Del,omitnil" name:"Del"`

	// 添加组织目录下所有设备通道（Json数组，可以为空，并且通道总数量不超过5000个（包括Add字段通道数量））
	OrganizationId []*string `json:"OrganizationId,omitnil" name:"OrganizationId"`
}

// Predefined struct for user
type UpdateRecordBackupPlanRequestParams struct {
	// 计划ID
	PlanId *string `json:"PlanId,omitnil" name:"PlanId"`

	// 修改的内容
	Mod *UpdateRecordBackupPlanModify `json:"Mod,omitnil" name:"Mod"`
}

type UpdateRecordBackupPlanRequest struct {
	*tchttp.BaseRequest
	
	// 计划ID
	PlanId *string `json:"PlanId,omitnil" name:"PlanId"`

	// 修改的内容
	Mod *UpdateRecordBackupPlanModify `json:"Mod,omitnil" name:"Mod"`
}

func (r *UpdateRecordBackupPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRecordBackupPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanId")
	delete(f, "Mod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateRecordBackupPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRecordBackupPlanResponseParams struct {
	// 返回数据
	Data *UpdateRecordBackupPlanData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpdateRecordBackupPlanResponse struct {
	*tchttp.BaseResponse
	Response *UpdateRecordBackupPlanResponseParams `json:"Response"`
}

func (r *UpdateRecordBackupPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRecordBackupPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateRecordBackupTemplateData struct {
	// 模板ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 模板名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateName *string `json:"TemplateName,omitnil" name:"TemplateName"`

	// 上云时间段（按周进行设置，支持一天设置多个时间段，每个时间段不小于10分钟）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeSections []*RecordTemplateTimeSections `json:"TimeSections,omitnil" name:"TimeSections"`

	// 录像时间段（按周进行设置，支持一天设置多个时间段，每个时间段不小于10分钟）
	// 注意：此字段可能返回 null，表示取不到有效值。
	DevTimeSections []*RecordTemplateTimeSections `json:"DevTimeSections,omitnil" name:"DevTimeSections"`

	// 上云倍速（支持1，2，4倍速）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Scale *int64 `json:"Scale,omitnil" name:"Scale"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateAt *string `json:"CreateAt,omitnil" name:"CreateAt"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateAt *string `json:"UpdateAt,omitnil" name:"UpdateAt"`
}

type UpdateRecordBackupTemplateModify struct {
	// 模板名称（不修改名称时，不需要带该字段）
	TemplateName *string `json:"TemplateName,omitnil" name:"TemplateName"`

	// 上云时间段（按周进行设置，支持一天设置多个时间段，每个时间段不小于10分钟）
	TimeSections []*RecordTemplateTimeSections `json:"TimeSections,omitnil" name:"TimeSections"`

	// 录像时间段（按周进行设置，支持一天设置多个时间段，每个时间段不小于10分钟）
	DevTimeSections []*RecordTemplateTimeSections `json:"DevTimeSections,omitnil" name:"DevTimeSections"`

	// 上云倍速（支持1，2，4倍速）
	Scale *int64 `json:"Scale,omitnil" name:"Scale"`
}

// Predefined struct for user
type UpdateRecordBackupTemplateRequestParams struct {
	// 模板ID（从查询录像上云模板列表接口ListRecordBackupTemplates中获取）
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 修改录像上云模板数据
	Mod *UpdateRecordBackupTemplateModify `json:"Mod,omitnil" name:"Mod"`
}

type UpdateRecordBackupTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板ID（从查询录像上云模板列表接口ListRecordBackupTemplates中获取）
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 修改录像上云模板数据
	Mod *UpdateRecordBackupTemplateModify `json:"Mod,omitnil" name:"Mod"`
}

func (r *UpdateRecordBackupTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRecordBackupTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "Mod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateRecordBackupTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRecordBackupTemplateResponseParams struct {
	// 返回数据
	Data *UpdateRecordBackupTemplateData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpdateRecordBackupTemplateResponse struct {
	*tchttp.BaseResponse
	Response *UpdateRecordBackupTemplateResponseParams `json:"Response"`
}

func (r *UpdateRecordBackupTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRecordBackupTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateRecordPlanData struct {
	// 上云计划名称，仅支持中文、英文、数字、_、-，长度不超过32个字符，计划名称全局唯一，不能为空，不能重复，不修改名称时，不需要该字段
	PlanName *string `json:"PlanName,omitnil" name:"PlanName"`

	// 上云模板ID，不修改模板ID时，不需要该字段
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 上云计划描述，仅支持中文、英文、数字、_、-，长度不超过128个字符， 不修改描述时，不需要该字段
	Describe *string `json:"Describe,omitnil" name:"Describe"`

	// 码流类型，default:不指定码流类型，以设备默认推送类型为主， main:主码流，sub:子码流，其他根据设备能力集自定义，长度不能超过32个字节
	StreamType *string `json:"StreamType,omitnil" name:"StreamType"`

	// 生命周期，文件生命周期设置，管理文件冷、热存储的时间，不修改生命周期时，不需要该字段
	LifeCycle *LifeCycleData `json:"LifeCycle,omitnil" name:"LifeCycle"`

	// 要新增的设备通道,Json数组，没有新增时，不需要该字段，一次添加通道总数不超过5000个，包括组织目录下的通道数量
	Add []*ChannelInfo `json:"Add,omitnil" name:"Add"`

	// 要删除的设备通道，Json数组，内容为要删除的设备通道id，没有删除设备通道时，不需要该字段
	Del []*string `json:"Del,omitnil" name:"Del"`

	// 组织目录ID，添加组织目录下所有设备通道，Json数组，可以为空，并且通道总数量不超过5000个（包括Add字段通道数量）
	OrganizationId []*string `json:"OrganizationId,omitnil" name:"OrganizationId"`
}

// Predefined struct for user
type UpdateRecordPlanRequestParams struct {
	// 计划ID
	PlanId *string `json:"PlanId,omitnil" name:"PlanId"`

	// 修改计划的内容
	Mod *UpdateRecordPlanData `json:"Mod,omitnil" name:"Mod"`
}

type UpdateRecordPlanRequest struct {
	*tchttp.BaseRequest
	
	// 计划ID
	PlanId *string `json:"PlanId,omitnil" name:"PlanId"`

	// 修改计划的内容
	Mod *UpdateRecordPlanData `json:"Mod,omitnil" name:"Mod"`
}

func (r *UpdateRecordPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRecordPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanId")
	delete(f, "Mod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateRecordPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRecordPlanResponseParams struct {
	// 返回结果
	Data *RecordPlanOptData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpdateRecordPlanResponse struct {
	*tchttp.BaseResponse
	Response *UpdateRecordPlanResponseParams `json:"Response"`
}

func (r *UpdateRecordPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRecordPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateRecordTemplateData struct {
	// 模板名称， 不修改名称时，不需要带该字段
	TemplateName *string `json:"TemplateName,omitnil" name:"TemplateName"`

	// 上云时间段，不修改名称时，不需要带该字段
	TimeSections []*RecordTemplateTimeSections `json:"TimeSections,omitnil" name:"TimeSections"`
}

// Predefined struct for user
type UpdateRecordTemplateRequestParams struct {
	// 模板ID
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 修改内容
	Mod *UpdateRecordTemplateData `json:"Mod,omitnil" name:"Mod"`
}

type UpdateRecordTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板ID
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 修改内容
	Mod *UpdateRecordTemplateData `json:"Mod,omitnil" name:"Mod"`
}

func (r *UpdateRecordTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRecordTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "Mod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateRecordTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRecordTemplateResponseParams struct {
	// 返回结果
	Data *RecordTemplateInfo `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpdateRecordTemplateResponse struct {
	*tchttp.BaseResponse
	Response *UpdateRecordTemplateResponseParams `json:"Response"`
}

func (r *UpdateRecordTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRecordTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUserDeviceRequestParams struct {
	// 设备ID（从获取设备列表接口ListDevices中获取）
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`

	// 设备名称（仅支持中文、英文、数字、_、-，长度不超过32个字符）
	Name *string `json:"Name,omitnil" name:"Name"`

	// 设备流传输协议，仅国标设备有效，填0则不做更改（1:UDP,2:TCP）
	TransportProtocol *int64 `json:"TransportProtocol,omitnil" name:"TransportProtocol"`

	// 设备密码（仅国标，网关设备支持）
	Password *string `json:"Password,omitnil" name:"Password"`

	// 设备描述（仅支持中文、英文、数字、_、-，长度不超过128位）
	Description *string `json:"Description,omitnil" name:"Description"`

	// 设备接入Ip（仅网关接入支持）
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 设备Port（仅网关接入支持）
	Port *int64 `json:"Port,omitnil" name:"Port"`

	// 设备用户名（仅网关接入支持）
	Username *string `json:"Username,omitnil" name:"Username"`
}

type UpdateUserDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 设备ID（从获取设备列表接口ListDevices中获取）
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`

	// 设备名称（仅支持中文、英文、数字、_、-，长度不超过32个字符）
	Name *string `json:"Name,omitnil" name:"Name"`

	// 设备流传输协议，仅国标设备有效，填0则不做更改（1:UDP,2:TCP）
	TransportProtocol *int64 `json:"TransportProtocol,omitnil" name:"TransportProtocol"`

	// 设备密码（仅国标，网关设备支持）
	Password *string `json:"Password,omitnil" name:"Password"`

	// 设备描述（仅支持中文、英文、数字、_、-，长度不超过128位）
	Description *string `json:"Description,omitnil" name:"Description"`

	// 设备接入Ip（仅网关接入支持）
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 设备Port（仅网关接入支持）
	Port *int64 `json:"Port,omitnil" name:"Port"`

	// 设备用户名（仅网关接入支持）
	Username *string `json:"Username,omitnil" name:"Username"`
}

func (r *UpdateUserDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUserDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	delete(f, "Name")
	delete(f, "TransportProtocol")
	delete(f, "Password")
	delete(f, "Description")
	delete(f, "Ip")
	delete(f, "Port")
	delete(f, "Username")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateUserDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUserDeviceResponseParams struct {
	// 返回数据
	Data *UpdateDeviceData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpdateUserDeviceResponse struct {
	*tchttp.BaseResponse
	Response *UpdateUserDeviceResponseParams `json:"Response"`
}

func (r *UpdateUserDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUserDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeGatewayRequestParams struct {
	// 网关索引ID（从获取网关列表ListGateways接口中获取）
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`
}

type UpgradeGatewayRequest struct {
	*tchttp.BaseRequest
	
	// 网关索引ID（从获取网关列表ListGateways接口中获取）
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`
}

func (r *UpgradeGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeGatewayResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpgradeGatewayResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeGatewayResponseParams `json:"Response"`
}

func (r *UpgradeGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}