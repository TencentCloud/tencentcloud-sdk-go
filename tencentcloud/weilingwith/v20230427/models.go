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

package v20230427

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type Action struct {
	// 动作id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 动作名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type ActionDetail struct {
	// 动作id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 动作名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 动作类型
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 动作说明
	ActionDesc *string `json:"ActionDesc,omitnil,omitempty" name:"ActionDesc"`

	// 消息类型，orgin/custom/model
	MsgType *string `json:"MsgType,omitnil,omitempty" name:"MsgType"`

	// 消息内容,有效值为x-json:后的字段
	MsgContent *string `json:"MsgContent,omitnil,omitempty" name:"MsgContent"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 设备唯一标识
	WID *string `json:"WID,omitnil,omitempty" name:"WID"`

	// 关联故障列表
	LinkRuleSet []*LinkRule `json:"LinkRuleSet,omitnil,omitempty" name:"LinkRuleSet"`

	// 动作下沉配置,有效值为x-json:后的字段
	SinkConfig *string `json:"SinkConfig,omitnil,omitempty" name:"SinkConfig"`
}

type ActionObj struct {
	// 动作id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 动作名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 动作类型。（app/推送消息至应用-携带空间设备：无,appWithNearbyDevices/推送至应用-携带空间设备：携带,device/推送消息至设备-指定设备,nearbyDevices/推送消息至设备-事件所在范围内的设备,toAlarm/转换为告警,toNotification/转换为通知）
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 动作说明
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 消息类型，orgin/custom/model
	MsgType *string `json:"MsgType,omitnil,omitempty" name:"MsgType"`

	// 消息内容
	MsgContent *string `json:"MsgContent,omitnil,omitempty" name:"MsgContent"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 动作下沉配置
	SinkConfig *string `json:"SinkConfig,omitnil,omitempty" name:"SinkConfig"`

	//  具体应用（appid）/具体设备（DIN/subID）
	ApplyDevice *string `json:"ApplyDevice,omitnil,omitempty" name:"ApplyDevice"`
}

// Predefined struct for user
type AddAlarmProcessRecordRequestParams struct {
	// 处理记录项
	RecordSet []*ProcessRecordInfo `json:"RecordSet,omitnil,omitempty" name:"RecordSet"`

	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 非孪生平台外部应用appid
	ApplicationId *int64 `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 此字段填写的是非孪生中台的用户id（多个用逗号分隔），如果是非孪生中台用户必填该字段
	ExtendOne *string `json:"ExtendOne,omitnil,omitempty" name:"ExtendOne"`
}

type AddAlarmProcessRecordRequest struct {
	*tchttp.BaseRequest
	
	// 处理记录项
	RecordSet []*ProcessRecordInfo `json:"RecordSet,omitnil,omitempty" name:"RecordSet"`

	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 非孪生平台外部应用appid
	ApplicationId *int64 `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 此字段填写的是非孪生中台的用户id（多个用逗号分隔），如果是非孪生中台用户必填该字段
	ExtendOne *string `json:"ExtendOne,omitnil,omitempty" name:"ExtendOne"`
}

func (r *AddAlarmProcessRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddAlarmProcessRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RecordSet")
	delete(f, "WorkspaceId")
	delete(f, "ApplicationToken")
	delete(f, "ApplicationId")
	delete(f, "ExtendOne")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddAlarmProcessRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddAlarmProcessRecordResponseParams struct {
	// 添加告警处理记录结果
	Result *EmptyRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddAlarmProcessRecordResponse struct {
	*tchttp.BaseResponse
	Response *AddAlarmProcessRecordResponseParams `json:"Response"`
}

func (r *AddAlarmProcessRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddAlarmProcessRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddDeviceInfo struct {
	// 产品id
	ProductId *int64 `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备sn序列号
	SN *string `json:"SN,omitnil,omitempty" name:"SN"`

	// 父设备wid，不为空表示导入子设备
	ParentWID *string `json:"ParentWID,omitnil,omitempty" name:"ParentWID"`

	// 密钥来源：0-使用产品密钥 1-使用设备特有的密钥
	KeySource *int64 `json:"KeySource,omitnil,omitempty" name:"KeySource"`
}

type AdministrationData struct {
	// 行政区划编码
	AdministrationCode *string `json:"AdministrationCode,omitnil,omitempty" name:"AdministrationCode"`

	// 行政区划名称
	AdministrationName *string `json:"AdministrationName,omitnil,omitempty" name:"AdministrationName"`
}

type AdministrativeDetail struct {
	// 行政区域类型编码
	AdministrativeTypeCode *string `json:"AdministrativeTypeCode,omitnil,omitempty" name:"AdministrativeTypeCode"`

	// 行政区域编码
	AdministrativeCode *string `json:"AdministrativeCode,omitnil,omitempty" name:"AdministrativeCode"`

	// 行政区域名称
	AdministrativeName *string `json:"AdministrativeName,omitnil,omitempty" name:"AdministrativeName"`
}

type AlarmInfo struct {
	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 告警ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 告警状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 告警时间
	Time *int64 `json:"Time,omitnil,omitempty" name:"Time"`

	// 告警业务类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 告警业务类型名称
	TypeName *string `json:"TypeName,omitnil,omitempty" name:"TypeName"`

	// 子告警类型
	SubType *string `json:"SubType,omitnil,omitempty" name:"SubType"`

	// 子告警类型名称
	SubTypeName *string `json:"SubTypeName,omitnil,omitempty" name:"SubTypeName"`

	// 告警级别id
	Level *int64 `json:"Level,omitnil,omitempty" name:"Level"`

	// 告警级别名称
	LevelName *string `json:"LevelName,omitnil,omitempty" name:"LevelName"`

	// 上报应用appid
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 设备wid
	WID *string `json:"WID,omitnil,omitempty" name:"WID"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 空间位置
	Position *string `json:"Position,omitnil,omitempty" name:"Position"`

	// 上报图片
	ReportImg *ReportImg `json:"ReportImg,omitnil,omitempty" name:"ReportImg"`

	// 告警描述
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 处理人
	HandlePersonSet []*HandlerPersonInfo `json:"HandlePersonSet,omitnil,omitempty" name:"HandlePersonSet"`

	// 处理记录
	HandleRecordSet []*HandleRecordInfo `json:"HandleRecordSet,omitnil,omitempty" name:"HandleRecordSet"`

	// 扩展信息
	Extend *string `json:"Extend,omitnil,omitempty" name:"Extend"`

	// 应用扩展字段1
	ExtendOne *string `json:"ExtendOne,omitnil,omitempty" name:"ExtendOne"`

	// 应用扩展字段2
	ExtendTwo *string `json:"ExtendTwo,omitnil,omitempty" name:"ExtendTwo"`

	// 应用透传字段,有效字段为x-json后的字段
	Echo *string `json:"Echo,omitnil,omitempty" name:"Echo"`
}

type AlarmLevelInfo struct {
	// 级别id
	LevelId *int64 `json:"LevelId,omitnil,omitempty" name:"LevelId"`

	// 级别名称
	LevelName *string `json:"LevelName,omitnil,omitempty" name:"LevelName"`
}

type AlarmStatusData struct {
	// 告警状态ID
	StatusID *string `json:"StatusID,omitnil,omitempty" name:"StatusID"`

	// 告警状态名称
	StatusName *string `json:"StatusName,omitnil,omitempty" name:"StatusName"`

	// 告警状态类型
	StatusType *string `json:"StatusType,omitnil,omitempty" name:"StatusType"`
}

type AlarmTypeDetailInfo struct {
	// 告警类型id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 父节点id
	ParentId *int64 `json:"ParentId,omitnil,omitempty" name:"ParentId"`

	// 0-标准告警类型，1-自定义告警类型
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 告警名称类型
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 告警类型英文名称
	EnglishName *string `json:"EnglishName,omitnil,omitempty" name:"EnglishName"`
}

type AlarmTypeInfo struct {
	// 告警父类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 告警子类型(如果传告警子类型，则必传父类型)
	SubType *string `json:"SubType,omitnil,omitempty" name:"SubType"`
}

type ApiContent struct {
	// 所属API的id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 参数名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 参数类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 是否为动态值
	Dynamic *bool `json:"Dynamic,omitnil,omitempty" name:"Dynamic"`

	// 是否必填
	Required *bool `json:"Required,omitnil,omitempty" name:"Required"`

	// 参数值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 默认值
	DefaultValue *string `json:"DefaultValue,omitnil,omitempty" name:"DefaultValue"`
}

type ApiInfo struct {
	// API的id
	ApiId *string `json:"ApiId,omitnil,omitempty" name:"ApiId"`

	// API名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// API所属应用的id
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// API所属的项目空间的id
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// API所属目录的编码
	PoiCode *string `json:"PoiCode,omitnil,omitempty" name:"PoiCode"`

	//  接口分类0. 其他服务 1. IOT服务 2. 空间服务 3.微应用服务 4.场景服务 5.AI算法服务 6.任务算法服务 7.第三方服务
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 数据授权 0:否 1:是
	DataAudit *uint64 `json:"DataAudit,omitnil,omitempty" name:"DataAudit"`

	// 是否需要申请 0:否 1:是
	ApplyAudit *uint64 `json:"ApplyAudit,omitnil,omitempty" name:"ApplyAudit"`

	// API详情
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// API地址
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// 请求方法类型
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// API状态
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// API预览地址
	PreviewUrl *string `json:"PreviewUrl,omitnil,omitempty" name:"PreviewUrl"`

	// query参数
	QueryParams []*ApiContent `json:"QueryParams,omitnil,omitempty" name:"QueryParams"`

	// 路径参数
	PathParams []*ApiContent `json:"PathParams,omitnil,omitempty" name:"PathParams"`

	// 请求头
	RequestHeaders []*ApiContent `json:"RequestHeaders,omitnil,omitempty" name:"RequestHeaders"`

	// 响应头
	ResponseHeaders []*ApiContent `json:"ResponseHeaders,omitnil,omitempty" name:"ResponseHeaders"`

	// 是否为公共空间接口
	IsCommonSpace *bool `json:"IsCommonSpace,omitnil,omitempty" name:"IsCommonSpace"`

	// 请求体（base64编码）
	Body *string `json:"Body,omitnil,omitempty" name:"Body"`

	// 响应体（base64编码）
	ResponseBody *string `json:"ResponseBody,omitnil,omitempty" name:"ResponseBody"`

	// 接口方式 1.http 2消息通知服务
	Style *uint64 `json:"Style,omitnil,omitempty" name:"Style"`
}

type ApiInfoList struct {
	// API列表
	ApiInfo []*ApiInfo `json:"ApiInfo,omitnil,omitempty" name:"ApiInfo"`

	// 数据总条数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type ApplicationInfo struct {
	// 应用分配的appId
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 应用中文名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 应用地址
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// 应用logo
	ApplicationLogo *ApplicationLogo `json:"ApplicationLogo,omitnil,omitempty" name:"ApplicationLogo"`

	// 应用类型，0:saas应用 1:平台应用
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// engine
	EnglishName *string `json:"EnglishName,omitnil,omitempty" name:"EnglishName"`

	// 能源管理应用
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ApplicationList struct {
	// 应用列表
	ApplicationInfoList []*ApplicationInfo `json:"ApplicationInfoList,omitnil,omitempty" name:"ApplicationInfoList"`

	// 当前查询条件命中的数据总条数
	TotalCount *string `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type ApplicationLogo struct {
	// logo图片对应的fileId
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// logo图片地址
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type ApplicationTokenInfo struct {
	// 应用申请调用API的令牌
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`
}

// Predefined struct for user
type BatchCreateDeviceRequestParams struct {
	// 工作空间Id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 设备信息项
	AddDeviceSet []*AddDeviceInfo `json:"AddDeviceSet,omitnil,omitempty" name:"AddDeviceSet"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

type BatchCreateDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 工作空间Id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 设备信息项
	AddDeviceSet []*AddDeviceInfo `json:"AddDeviceSet,omitnil,omitempty" name:"AddDeviceSet"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

func (r *BatchCreateDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchCreateDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "AddDeviceSet")
	delete(f, "ApplicationToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchCreateDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type BatchCreateDeviceRes struct {
	// 新增成功的设备列表
	SuccessSet []*CreateDeviceSucceeded `json:"SuccessSet,omitnil,omitempty" name:"SuccessSet"`

	// 新增失败的设备列表
	FailSet []*CreateDeviceFailed `json:"FailSet,omitnil,omitempty" name:"FailSet"`
}

// Predefined struct for user
type BatchCreateDeviceResponseParams struct {
	// 批量新增设备返回结果
	Result *BatchCreateDeviceRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BatchCreateDeviceResponse struct {
	*tchttp.BaseResponse
	Response *BatchCreateDeviceResponseParams `json:"Response"`
}

func (r *BatchCreateDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchCreateDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchDeleteDeviceRequestParams struct {
	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 设备wid数组列表
	WIDSet []*string `json:"WIDSet,omitnil,omitempty" name:"WIDSet"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

type BatchDeleteDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 设备wid数组列表
	WIDSet []*string `json:"WIDSet,omitnil,omitempty" name:"WIDSet"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

func (r *BatchDeleteDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDeleteDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "WIDSet")
	delete(f, "ApplicationToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchDeleteDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchDeleteDeviceResponseParams struct {
	// 返回请求结果
	Result *EmptyRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BatchDeleteDeviceResponse struct {
	*tchttp.BaseResponse
	Response *BatchDeleteDeviceResponseParams `json:"Response"`
}

func (r *BatchDeleteDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDeleteDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchKillAlarmRequestParams struct {
	// 告警开始时间，必填,时间戳秒
	BeginTime *int64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 告警结束时间，必填，时间戳秒
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 告警状态: unprocessed processing
	StatusSet []*string `json:"StatusSet,omitnil,omitempty" name:"StatusSet"`

	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 当前操作用户id
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 当前操作用户名称
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 当前告警处理人，填孪生中台用户id,多个用逗号分隔
	ProcessorId *string `json:"ProcessorId,omitnil,omitempty" name:"ProcessorId"`

	// 告警子类型(如果传告警子类型，则必传父类型)
	AlarmTypeSet []*AlarmTypeInfo `json:"AlarmTypeSet,omitnil,omitempty" name:"AlarmTypeSet"`

	// 告警级别,包括1~5
	LevelSet []*int64 `json:"LevelSet,omitnil,omitempty" name:"LevelSet"`

	// 设备id
	WIDSet []*string `json:"WIDSet,omitnil,omitempty" name:"WIDSet"`

	// 告警id
	IdSet []*string `json:"IdSet,omitnil,omitempty" name:"IdSet"`

	// 告警处理的说明
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`
}

type BatchKillAlarmRequest struct {
	*tchttp.BaseRequest
	
	// 告警开始时间，必填,时间戳秒
	BeginTime *int64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 告警结束时间，必填，时间戳秒
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 告警状态: unprocessed processing
	StatusSet []*string `json:"StatusSet,omitnil,omitempty" name:"StatusSet"`

	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 当前操作用户id
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 当前操作用户名称
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 当前告警处理人，填孪生中台用户id,多个用逗号分隔
	ProcessorId *string `json:"ProcessorId,omitnil,omitempty" name:"ProcessorId"`

	// 告警子类型(如果传告警子类型，则必传父类型)
	AlarmTypeSet []*AlarmTypeInfo `json:"AlarmTypeSet,omitnil,omitempty" name:"AlarmTypeSet"`

	// 告警级别,包括1~5
	LevelSet []*int64 `json:"LevelSet,omitnil,omitempty" name:"LevelSet"`

	// 设备id
	WIDSet []*string `json:"WIDSet,omitnil,omitempty" name:"WIDSet"`

	// 告警id
	IdSet []*string `json:"IdSet,omitnil,omitempty" name:"IdSet"`

	// 告警处理的说明
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`
}

func (r *BatchKillAlarmRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchKillAlarmRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "StatusSet")
	delete(f, "WorkspaceId")
	delete(f, "UserId")
	delete(f, "UserName")
	delete(f, "ApplicationToken")
	delete(f, "ProcessorId")
	delete(f, "AlarmTypeSet")
	delete(f, "LevelSet")
	delete(f, "WIDSet")
	delete(f, "IdSet")
	delete(f, "Desc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchKillAlarmRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchKillAlarmResponseParams struct {
	// 批量消警结果
	Result *EmptyRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BatchKillAlarmResponse struct {
	*tchttp.BaseResponse
	Response *BatchKillAlarmResponseParams `json:"Response"`
}

func (r *BatchKillAlarmResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchKillAlarmResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchReportAppMessageRequestParams struct {
	// 工作空间Id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 消息上报请求列表
	ReportSet []*ReportAppMessage `json:"ReportSet,omitnil,omitempty" name:"ReportSet"`
}

type BatchReportAppMessageRequest struct {
	*tchttp.BaseRequest
	
	// 工作空间Id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 消息上报请求列表
	ReportSet []*ReportAppMessage `json:"ReportSet,omitnil,omitempty" name:"ReportSet"`
}

func (r *BatchReportAppMessageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchReportAppMessageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "ApplicationToken")
	delete(f, "ReportSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchReportAppMessageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type BatchReportAppMessageRes struct {
	// 上报数量
	TotalElements *int64 `json:"TotalElements,omitnil,omitempty" name:"TotalElements"`

	// 提交数量（推送成功）
	Commit *int64 `json:"Commit,omitnil,omitempty" name:"Commit"`

	// 消息推送结果列表
	SpanMap []*ReportMsgRes `json:"SpanMap,omitnil,omitempty" name:"SpanMap"`
}

// Predefined struct for user
type BatchReportAppMessageResponseParams struct {
	// 批量消息上报结果
	Result *BatchReportAppMessageRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BatchReportAppMessageResponse struct {
	*tchttp.BaseResponse
	Response *BatchReportAppMessageResponseParams `json:"Response"`
}

func (r *BatchReportAppMessageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchReportAppMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BuildingListRes struct {
	// 建筑列表
	BuildingProfileList []*BuildingProfile `json:"BuildingProfileList,omitnil,omitempty" name:"BuildingProfileList"`
}

type BuildingModel struct {
	// 构件ID
	ElementId *string `json:"ElementId,omitnil,omitempty" name:"ElementId"`

	// 构件名称
	ElementName *string `json:"ElementName,omitnil,omitempty" name:"ElementName"`

	// 模型类型
	ModelType *string `json:"ModelType,omitnil,omitempty" name:"ModelType"`

	// 模型URL
	ModelUrl *string `json:"ModelUrl,omitnil,omitempty" name:"ModelUrl"`
}

type BuildingModelRes struct {
	// 建模模型信息出参
	Models []*BuildingModel `json:"Models,omitnil,omitempty" name:"Models"`
}

type BuildingProfile struct {
	// 建筑id
	BuildingId *string `json:"BuildingId,omitnil,omitempty" name:"BuildingId"`

	// 建筑名称
	BuildingName *string `json:"BuildingName,omitnil,omitempty" name:"BuildingName"`

	// 空间编码
	SpaceCode *string `json:"SpaceCode,omitnil,omitempty" name:"SpaceCode"`

	// 经度
	Longitude *float64 `json:"Longitude,omitnil,omitempty" name:"Longitude"`

	// 纬度
	Latitude *float64 `json:"Latitude,omitnil,omitempty" name:"Latitude"`

	// 地址
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`
}

type BuildingProfileRes struct {
	// 建筑概要信息
	BuildingProfile *BuildingProfile `json:"BuildingProfile,omitnil,omitempty" name:"BuildingProfile"`
}

type CameraExtendInfoRes struct {
	// 存储方式 (nvr或cosmtav)
	SaveType *string `json:"SaveType,omitnil,omitempty" name:"SaveType"`

	// 云存储天数（save_type是cosmtav时这个参数才有效）
	SaveDay *int64 `json:"SaveDay,omitnil,omitempty" name:"SaveDay"`

	// 实时分辨率
	LiveResolution *int64 `json:"LiveResolution,omitnil,omitempty" name:"LiveResolution"`

	// 历史分辨率
	HistoryResolution *int64 `json:"HistoryResolution,omitnil,omitempty" name:"HistoryResolution"`
}

// Predefined struct for user
type ChangeAlarmStatusRequestParams struct {
	// 告警的id，返回的列表中的id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 告警状态 processed unprocessed processing misreport shield
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 告警处理时间
	ProcessTime *int64 `json:"ProcessTime,omitnil,omitempty" name:"ProcessTime"`

	// 处理类型
	ProcessType *string `json:"ProcessType,omitnil,omitempty" name:"ProcessType"`

	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 当前操作用户id
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 当前操作用户名称
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 平台告警处理人
	Processor *string `json:"Processor,omitnil,omitempty" name:"Processor"`

	// 告警处理的描述信息
	ProcessDescription *string `json:"ProcessDescription,omitnil,omitempty" name:"ProcessDescription"`

	// 告警处理的扩展信息，可以为JSON字符串
	ProcessExtend *string `json:"ProcessExtend,omitnil,omitempty" name:"ProcessExtend"`

	// 扩展字段1，目前存的应用告警处理人
	ExtendOne *string `json:"ExtendOne,omitnil,omitempty" name:"ExtendOne"`

	// 应用id
	ApplicationId *int64 `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`
}

type ChangeAlarmStatusRequest struct {
	*tchttp.BaseRequest
	
	// 告警的id，返回的列表中的id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 告警状态 processed unprocessed processing misreport shield
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 告警处理时间
	ProcessTime *int64 `json:"ProcessTime,omitnil,omitempty" name:"ProcessTime"`

	// 处理类型
	ProcessType *string `json:"ProcessType,omitnil,omitempty" name:"ProcessType"`

	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 当前操作用户id
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 当前操作用户名称
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 平台告警处理人
	Processor *string `json:"Processor,omitnil,omitempty" name:"Processor"`

	// 告警处理的描述信息
	ProcessDescription *string `json:"ProcessDescription,omitnil,omitempty" name:"ProcessDescription"`

	// 告警处理的扩展信息，可以为JSON字符串
	ProcessExtend *string `json:"ProcessExtend,omitnil,omitempty" name:"ProcessExtend"`

	// 扩展字段1，目前存的应用告警处理人
	ExtendOne *string `json:"ExtendOne,omitnil,omitempty" name:"ExtendOne"`

	// 应用id
	ApplicationId *int64 `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`
}

func (r *ChangeAlarmStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeAlarmStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Status")
	delete(f, "ProcessTime")
	delete(f, "ProcessType")
	delete(f, "WorkspaceId")
	delete(f, "UserId")
	delete(f, "UserName")
	delete(f, "ApplicationToken")
	delete(f, "Processor")
	delete(f, "ProcessDescription")
	delete(f, "ProcessExtend")
	delete(f, "ExtendOne")
	delete(f, "ApplicationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChangeAlarmStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChangeAlarmStatusResponseParams struct {
	// 返回空结果
	Result *EmptyRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChangeAlarmStatusResponse struct {
	*tchttp.BaseResponse
	Response *ChangeAlarmStatusResponseParams `json:"Response"`
}

func (r *ChangeAlarmStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeAlarmStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ControlCameraPTZRequestParams struct {
	// 设备的唯一标识
	WID *string `json:"WID,omitnil,omitempty" name:"WID"`

	// ptz指令
	// left - 向左移动
	// right - 向右移动
	// up - 向上移动
	// down - 向下
	// zoomOut - 镜头缩小
	// zoomIn - 镜头放大
	// irisIn - 光圈缩小
	// irisOut - 光圈放大
	// focusIn - 焦距变近
	// focusOut - 焦距变远
	CMD *string `json:"CMD,omitnil,omitempty" name:"CMD"`

	// 工作空间Id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

type ControlCameraPTZRequest struct {
	*tchttp.BaseRequest
	
	// 设备的唯一标识
	WID *string `json:"WID,omitnil,omitempty" name:"WID"`

	// ptz指令
	// left - 向左移动
	// right - 向右移动
	// up - 向上移动
	// down - 向下
	// zoomOut - 镜头缩小
	// zoomIn - 镜头放大
	// irisIn - 光圈缩小
	// irisOut - 光圈放大
	// focusIn - 焦距变近
	// focusOut - 焦距变远
	CMD *string `json:"CMD,omitnil,omitempty" name:"CMD"`

	// 工作空间Id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

func (r *ControlCameraPTZRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlCameraPTZRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WID")
	delete(f, "CMD")
	delete(f, "WorkspaceId")
	delete(f, "ApplicationToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ControlCameraPTZRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ControlCameraPTZResponseParams struct {
	// 控制摄像头结果返回
	Result *EmptyRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ControlCameraPTZResponse struct {
	*tchttp.BaseResponse
	Response *ControlCameraPTZResponseParams `json:"Response"`
}

func (r *ControlCameraPTZResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlCameraPTZResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ControlDeviceRequestParams struct {
	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 设备wid，最大100个
	WIDSet []*string `json:"WIDSet,omitnil,omitempty" name:"WIDSet"`

	// 控制内容
	ControlData *string `json:"ControlData,omitnil,omitempty" name:"ControlData"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 是否同步返回设备下控ack结果
	IsSynchronized *bool `json:"IsSynchronized,omitnil,omitempty" name:"IsSynchronized"`
}

type ControlDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 设备wid，最大100个
	WIDSet []*string `json:"WIDSet,omitnil,omitempty" name:"WIDSet"`

	// 控制内容
	ControlData *string `json:"ControlData,omitnil,omitempty" name:"ControlData"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 是否同步返回设备下控ack结果
	IsSynchronized *bool `json:"IsSynchronized,omitnil,omitempty" name:"IsSynchronized"`
}

func (r *ControlDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "WIDSet")
	delete(f, "ControlData")
	delete(f, "ApplicationToken")
	delete(f, "IsSynchronized")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ControlDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ControlDeviceRes struct {
	// 设备Id
	WID *string `json:"WID,omitnil,omitempty" name:"WID"`

	// 指令接受, 0表示成功
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 控制结果
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 批量大于1时，可用此seq进行链路追踪
	Seq *string `json:"Seq,omitnil,omitempty" name:"Seq"`
}

// Predefined struct for user
type ControlDeviceResponseParams struct {
	// 设备控制后结果集
	Result *ControlDeviceSet `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ControlDeviceResponse struct {
	*tchttp.BaseResponse
	Response *ControlDeviceResponseParams `json:"Response"`
}

func (r *ControlDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ControlDeviceSet struct {
	// 设备控制后返回结果集合
	Set []*ControlDeviceRes `json:"Set,omitnil,omitempty" name:"Set"`
}

// Predefined struct for user
type CreateApplicationTokenRequestParams struct {
	// 应用id
	ApplicationId *uint64 `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 一个随机数或者时间戳，用于防止重放攻击，每个请求唯一，建议用uuid
	Nonce *string `json:"Nonce,omitnil,omitempty" name:"Nonce"`

	// 租户id
	TenantId *uint64 `json:"TenantId,omitnil,omitempty" name:"TenantId"`

	// 请求时间，当前时间的unix毫秒时间戳
	RequestTime *uint64 `json:"RequestTime,omitnil,omitempty" name:"RequestTime"`

	// 签名方法见用户使用文档
	Signature *string `json:"Signature,omitnil,omitempty" name:"Signature"`
}

type CreateApplicationTokenRequest struct {
	*tchttp.BaseRequest
	
	// 应用id
	ApplicationId *uint64 `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 一个随机数或者时间戳，用于防止重放攻击，每个请求唯一，建议用uuid
	Nonce *string `json:"Nonce,omitnil,omitempty" name:"Nonce"`

	// 租户id
	TenantId *uint64 `json:"TenantId,omitnil,omitempty" name:"TenantId"`

	// 请求时间，当前时间的unix毫秒时间戳
	RequestTime *uint64 `json:"RequestTime,omitnil,omitempty" name:"RequestTime"`

	// 签名方法见用户使用文档
	Signature *string `json:"Signature,omitnil,omitempty" name:"Signature"`
}

func (r *CreateApplicationTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "Nonce")
	delete(f, "TenantId")
	delete(f, "RequestTime")
	delete(f, "Signature")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApplicationTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationTokenResponseParams struct {
	// 应用令牌信息
	Result *ApplicationTokenInfo `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateApplicationTokenResponse struct {
	*tchttp.BaseResponse
	Response *CreateApplicationTokenResponseParams `json:"Response"`
}

func (r *CreateApplicationTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDeviceFailed struct {
	// 产品id
	ProductId *int64 `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 父设备wid，不为空表示导入自设备
	ParentWID *string `json:"ParentWID,omitnil,omitempty" name:"ParentWID"`

	// 失败原因
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 设备sn序列号
	SN *string `json:"SN,omitnil,omitempty" name:"SN"`
}

type CreateDeviceSucceeded struct {
	// 产品id
	ProductId *int64 `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 父设备wid，不为空表示导入自设备
	ParentWID *string `json:"ParentWID,omitnil,omitempty" name:"ParentWID"`

	// 设备编码
	WID *string `json:"WID,omitnil,omitempty" name:"WID"`

	// 设备sn序列号
	SN *string `json:"SN,omitnil,omitempty" name:"SN"`
}

type CustomField struct {
	// 字段id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 字段内容
	Val *string `json:"Val,omitnil,omitempty" name:"Val"`
}

type CustomFieldInfo struct {
	// 字段id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 字段key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 字段名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 字段值
	Val *string `json:"Val,omitnil,omitempty" name:"Val"`
}

// Predefined struct for user
type DeleteDeviceGroupRequestParams struct {
	// 设备分组的id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 工作空间的id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

type DeleteDeviceGroupRequest struct {
	*tchttp.BaseRequest
	
	// 设备分组的id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 工作空间的id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

func (r *DeleteDeviceGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDeviceGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "WorkspaceId")
	delete(f, "ApplicationToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDeviceGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDeviceGroupResponseParams struct {
	// 无返回信息
	Result *EmptyRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDeviceGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDeviceGroupResponseParams `json:"Response"`
}

func (r *DeleteDeviceGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDeviceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeActionListRequestParams struct {
	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 分页查询，第几页，必传，大于0
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页条数，必传大于0
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 动作类型，（app,device,toAlarm,toNotification）
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 事件id详情
	IdSet []*int64 `json:"IdSet,omitnil,omitempty" name:"IdSet"`
}

type DescribeActionListRequest struct {
	*tchttp.BaseRequest
	
	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 分页查询，第几页，必传，大于0
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页条数，必传大于0
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 动作类型，（app,device,toAlarm,toNotification）
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 事件id详情
	IdSet []*int64 `json:"IdSet,omitnil,omitempty" name:"IdSet"`
}

func (r *DescribeActionListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeActionListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "ApplicationToken")
	delete(f, "ActionType")
	delete(f, "IdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeActionListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeActionListRes struct {
	// 第几页
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页条数
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 总页数
	TotalPage *int64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// 总条数
	TotalRow *int64 `json:"TotalRow,omitnil,omitempty" name:"TotalRow"`

	// 动作列表查询集合
	ActionDetailSet []*ActionDetail `json:"ActionDetailSet,omitnil,omitempty" name:"ActionDetailSet"`
}

// Predefined struct for user
type DescribeActionListResponseParams struct {
	// 动作列表查询结果
	Result *DescribeActionListRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeActionListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeActionListResponseParams `json:"Response"`
}

func (r *DescribeActionListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeActionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAdministrationByTagRequestParams struct {
	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 工作空间ID
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 标签值
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`
}

type DescribeAdministrationByTagRequest struct {
	*tchttp.BaseRequest
	
	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 工作空间ID
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 标签值
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`
}

func (r *DescribeAdministrationByTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAdministrationByTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationToken")
	delete(f, "WorkspaceId")
	delete(f, "Tag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAdministrationByTagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAdministrationByTagRes struct {
	// 行政区划列表
	List []*AdministrationData `json:"List,omitnil,omitempty" name:"List"`
}

// Predefined struct for user
type DescribeAdministrationByTagResponseParams struct {
	// 行政区划返回结构
	Result *DescribeAdministrationByTagRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAdministrationByTagResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAdministrationByTagResponseParams `json:"Response"`
}

func (r *DescribeAdministrationByTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAdministrationByTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmLevelListRequestParams struct {
	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

type DescribeAlarmLevelListRequest struct {
	*tchttp.BaseRequest
	
	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

func (r *DescribeAlarmLevelListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmLevelListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "ApplicationToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmLevelListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmLevelListRes struct {
	// 告警级别枚举获取数组
	AlarmLevelSet []*AlarmLevelInfo `json:"AlarmLevelSet,omitnil,omitempty" name:"AlarmLevelSet"`
}

// Predefined struct for user
type DescribeAlarmLevelListResponseParams struct {
	// 告警级别列表查询结果
	Result *DescribeAlarmLevelListRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAlarmLevelListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlarmLevelListResponseParams `json:"Response"`
}

func (r *DescribeAlarmLevelListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmLevelListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmListRequestParams struct {
	// 告警开始时间，必填,时间戳秒
	BeginTime *int64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 告警结束时间，必填，时间戳秒
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页查询，第几页
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页条数
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 告警状态
	Statuses []*string `json:"Statuses,omitnil,omitempty" name:"Statuses"`

	// 告警类型
	AlarmTypeSet []*AlarmTypeInfo `json:"AlarmTypeSet,omitnil,omitempty" name:"AlarmTypeSet"`

	// 告警级别id
	LevelSet []*int64 `json:"LevelSet,omitnil,omitempty" name:"LevelSet"`

	// 告警id
	IdSet []*string `json:"IdSet,omitnil,omitempty" name:"IdSet"`

	// 应用id
	AppIdSet []*int64 `json:"AppIdSet,omitnil,omitempty" name:"AppIdSet"`

	// 设备id
	WIDSet []*string `json:"WIDSet,omitnil,omitempty" name:"WIDSet"`

	// 空间层级
	SpaceCodeSet []*string `json:"SpaceCodeSet,omitnil,omitempty" name:"SpaceCodeSet"`

	// 应用扩展字段1
	ExtendOne []*string `json:"ExtendOne,omitnil,omitempty" name:"ExtendOne"`

	// 应用扩展字段2
	ExtendTwo []*string `json:"ExtendTwo,omitnil,omitempty" name:"ExtendTwo"`

	// 当前告警处理人，填孪生中台的用户id
	ProcessorSet []*string `json:"ProcessorSet,omitnil,omitempty" name:"ProcessorSet"`

	// 分组id
	GroupIdSet []*int64 `json:"GroupIdSet,omitnil,omitempty" name:"GroupIdSet"`
}

type DescribeAlarmListRequest struct {
	*tchttp.BaseRequest
	
	// 告警开始时间，必填,时间戳秒
	BeginTime *int64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 告警结束时间，必填，时间戳秒
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页查询，第几页
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页条数
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 告警状态
	Statuses []*string `json:"Statuses,omitnil,omitempty" name:"Statuses"`

	// 告警类型
	AlarmTypeSet []*AlarmTypeInfo `json:"AlarmTypeSet,omitnil,omitempty" name:"AlarmTypeSet"`

	// 告警级别id
	LevelSet []*int64 `json:"LevelSet,omitnil,omitempty" name:"LevelSet"`

	// 告警id
	IdSet []*string `json:"IdSet,omitnil,omitempty" name:"IdSet"`

	// 应用id
	AppIdSet []*int64 `json:"AppIdSet,omitnil,omitempty" name:"AppIdSet"`

	// 设备id
	WIDSet []*string `json:"WIDSet,omitnil,omitempty" name:"WIDSet"`

	// 空间层级
	SpaceCodeSet []*string `json:"SpaceCodeSet,omitnil,omitempty" name:"SpaceCodeSet"`

	// 应用扩展字段1
	ExtendOne []*string `json:"ExtendOne,omitnil,omitempty" name:"ExtendOne"`

	// 应用扩展字段2
	ExtendTwo []*string `json:"ExtendTwo,omitnil,omitempty" name:"ExtendTwo"`

	// 当前告警处理人，填孪生中台的用户id
	ProcessorSet []*string `json:"ProcessorSet,omitnil,omitempty" name:"ProcessorSet"`

	// 分组id
	GroupIdSet []*int64 `json:"GroupIdSet,omitnil,omitempty" name:"GroupIdSet"`
}

func (r *DescribeAlarmListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "WorkspaceId")
	delete(f, "ApplicationToken")
	delete(f, "Statuses")
	delete(f, "AlarmTypeSet")
	delete(f, "LevelSet")
	delete(f, "IdSet")
	delete(f, "AppIdSet")
	delete(f, "WIDSet")
	delete(f, "SpaceCodeSet")
	delete(f, "ExtendOne")
	delete(f, "ExtendTwo")
	delete(f, "ProcessorSet")
	delete(f, "GroupIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmListRes struct {
	// 第几页
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页条数
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 总页数
	TotalPage *int64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// 总条数
	TotalRow *int64 `json:"TotalRow,omitnil,omitempty" name:"TotalRow"`

	// 告警列表集合
	AlarmInfoSet []*AlarmInfo `json:"AlarmInfoSet,omitnil,omitempty" name:"AlarmInfoSet"`
}

// Predefined struct for user
type DescribeAlarmListResponseParams struct {
	// 告警列表查询结果
	Result *DescribeAlarmListRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAlarmListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlarmListResponseParams `json:"Response"`
}

func (r *DescribeAlarmListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmStatusListRequestParams struct {
	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 工作空间ID
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`
}

type DescribeAlarmStatusListRequest struct {
	*tchttp.BaseRequest
	
	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 工作空间ID
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`
}

func (r *DescribeAlarmStatusListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmStatusListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationToken")
	delete(f, "WorkspaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmStatusListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmStatusListRes struct {
	// 告警状态返回结构
	List []*AlarmStatusData `json:"List,omitnil,omitempty" name:"List"`
}

// Predefined struct for user
type DescribeAlarmStatusListResponseParams struct {
	// 告警状态返回结构
	Result *DescribeAlarmStatusListRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAlarmStatusListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlarmStatusListResponseParams `json:"Response"`
}

func (r *DescribeAlarmStatusListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmStatusListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmTypeListRequestParams struct {
	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 一级类型
	ParentType *string `json:"ParentType,omitnil,omitempty" name:"ParentType"`
}

type DescribeAlarmTypeListRequest struct {
	*tchttp.BaseRequest
	
	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 一级类型
	ParentType *string `json:"ParentType,omitnil,omitempty" name:"ParentType"`
}

func (r *DescribeAlarmTypeListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmTypeListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "ApplicationToken")
	delete(f, "ParentType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmTypeListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmTypeListRes struct {
	// 告警类型查询列表
	AlarmTypeSet []*AlarmTypeDetailInfo `json:"AlarmTypeSet,omitnil,omitempty" name:"AlarmTypeSet"`
}

// Predefined struct for user
type DescribeAlarmTypeListResponseParams struct {
	// 告警类型列表查询
	Result *DescribeAlarmTypeListRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAlarmTypeListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlarmTypeListResponseParams `json:"Response"`
}

func (r *DescribeAlarmTypeListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmTypeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationListRequestParams struct {
	// 项目空间id，本次查询返回的应用均关联至该空间
	WorkspaceId *uint64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 应用id数组，可选，填了则表示根据id批量查询
	ApplicationId []*uint64 `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 请求页号
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 页容量，默认为10
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeApplicationListRequest struct {
	*tchttp.BaseRequest
	
	// 项目空间id，本次查询返回的应用均关联至该空间
	WorkspaceId *uint64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 应用id数组，可选，填了则表示根据id批量查询
	ApplicationId []*uint64 `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 请求页号
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 页容量，默认为10
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *DescribeApplicationListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "ApplicationToken")
	delete(f, "ApplicationId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationListResponseParams struct {
	// 应用列表
	Result *ApplicationList `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeApplicationListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationListResponseParams `json:"Response"`
}

func (r *DescribeApplicationListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBuildingListRequestParams struct {
	// 项目空间id
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 是否有模型文件
	HasModel *bool `json:"HasModel,omitnil,omitempty" name:"HasModel"`

	// 空间编码
	SpaceCodes []*string `json:"SpaceCodes,omitnil,omitempty" name:"SpaceCodes"`
}

type DescribeBuildingListRequest struct {
	*tchttp.BaseRequest
	
	// 项目空间id
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 是否有模型文件
	HasModel *bool `json:"HasModel,omitnil,omitempty" name:"HasModel"`

	// 空间编码
	SpaceCodes []*string `json:"SpaceCodes,omitnil,omitempty" name:"SpaceCodes"`
}

func (r *DescribeBuildingListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBuildingListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "ApplicationToken")
	delete(f, "HasModel")
	delete(f, "SpaceCodes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBuildingListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBuildingListResponseParams struct {
	// 查询建筑列表出参
	Result *BuildingListRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBuildingListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBuildingListResponseParams `json:"Response"`
}

func (r *DescribeBuildingListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBuildingListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBuildingModelRequestParams struct {
	// 建筑id
	BuildingId *string `json:"BuildingId,omitnil,omitempty" name:"BuildingId"`

	// 项目空间id
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

type DescribeBuildingModelRequest struct {
	*tchttp.BaseRequest
	
	// 建筑id
	BuildingId *string `json:"BuildingId,omitnil,omitempty" name:"BuildingId"`

	// 项目空间id
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

func (r *DescribeBuildingModelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBuildingModelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BuildingId")
	delete(f, "WorkspaceId")
	delete(f, "ApplicationToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBuildingModelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBuildingModelResponseParams struct {
	// 建模模型信息出参
	Result *BuildingModelRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBuildingModelResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBuildingModelResponseParams `json:"Response"`
}

func (r *DescribeBuildingModelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBuildingModelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBuildingProfileRequestParams struct {
	// 建筑id
	BuildingId *string `json:"BuildingId,omitnil,omitempty" name:"BuildingId"`

	// 项目空间id
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

type DescribeBuildingProfileRequest struct {
	*tchttp.BaseRequest
	
	// 建筑id
	BuildingId *string `json:"BuildingId,omitnil,omitempty" name:"BuildingId"`

	// 项目空间id
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

func (r *DescribeBuildingProfileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBuildingProfileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BuildingId")
	delete(f, "WorkspaceId")
	delete(f, "ApplicationToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBuildingProfileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBuildingProfileResponseParams struct {
	// 查询建筑信息出参
	Result *BuildingProfileRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBuildingProfileResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBuildingProfileResponseParams `json:"Response"`
}

func (r *DescribeBuildingProfileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBuildingProfileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCameraExtendInfoRequestParams struct {
	// 设备的唯一标识
	WID *string `json:"WID,omitnil,omitempty" name:"WID"`

	// 工作空间Id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

type DescribeCameraExtendInfoRequest struct {
	*tchttp.BaseRequest
	
	// 设备的唯一标识
	WID *string `json:"WID,omitnil,omitempty" name:"WID"`

	// 工作空间Id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

func (r *DescribeCameraExtendInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCameraExtendInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WID")
	delete(f, "WorkspaceId")
	delete(f, "ApplicationToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCameraExtendInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCameraExtendInfoResponseParams struct {
	// 获取视频扩展信息结果
	Result *CameraExtendInfoRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCameraExtendInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCameraExtendInfoResponseParams `json:"Response"`
}

func (r *DescribeCameraExtendInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCameraExtendInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCityWorkspaceListRequestParams struct {
	// 行政区编码集合
	AdministrativeCodeSet []*string `json:"AdministrativeCodeSet,omitnil,omitempty" name:"AdministrativeCodeSet"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

type DescribeCityWorkspaceListRequest struct {
	*tchttp.BaseRequest
	
	// 行政区编码集合
	AdministrativeCodeSet []*string `json:"AdministrativeCodeSet,omitnil,omitempty" name:"AdministrativeCodeSet"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

func (r *DescribeCityWorkspaceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCityWorkspaceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AdministrativeCodeSet")
	delete(f, "ApplicationToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCityWorkspaceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCityWorkspaceListRes struct {
	// 通过城市id查询工作空间列表结果
	WorkspaceSet []*WorkspaceInfo `json:"WorkspaceSet,omitnil,omitempty" name:"WorkspaceSet"`
}

// Predefined struct for user
type DescribeCityWorkspaceListResponseParams struct {
	// 工作空间信息集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *DescribeCityWorkspaceListRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCityWorkspaceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCityWorkspaceListResponseParams `json:"Response"`
}

func (r *DescribeCityWorkspaceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCityWorkspaceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceGroupListRequestParams struct {
	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 工作空间ID
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 分组id, 不传默认全部
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

type DescribeDeviceGroupListRequest struct {
	*tchttp.BaseRequest
	
	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 工作空间ID
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 分组id, 不传默认全部
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

func (r *DescribeDeviceGroupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceGroupListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationToken")
	delete(f, "WorkspaceId")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceGroupListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceGroupListRes struct {
	// 设备分组list
	List []*DescribeGroupInfo `json:"List,omitnil,omitempty" name:"List"`
}

// Predefined struct for user
type DescribeDeviceGroupListResponseParams struct {
	// 分组信息
	Result *DescribeDeviceGroupListRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDeviceGroupListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceGroupListResponseParams `json:"Response"`
}

func (r *DescribeDeviceGroupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceListRequestParams struct {
	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 分页查询，第几页，必传，大于0
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页条数，必传大于0
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 设备类型，非必填
	DeviceTypeSet []*string `json:"DeviceTypeSet,omitnil,omitempty" name:"DeviceTypeSet"`

	// 产品 pid，非必填
	ProductIdSet []*int64 `json:"ProductIdSet,omitnil,omitempty" name:"ProductIdSet"`

	// 设备标签，非必填
	TagIdSet []*int64 `json:"TagIdSet,omitnil,omitempty" name:"TagIdSet"`

	// 空间层级
	SpaceCodeSet []*string `json:"SpaceCodeSet,omitnil,omitempty" name:"SpaceCodeSet"`

	// 设备标签名，非必填
	DeviceTagSet []*string `json:"DeviceTagSet,omitnil,omitempty" name:"DeviceTagSet"`

	// 设备wid,非必填
	WIDSet []*string `json:"WIDSet,omitnil,omitempty" name:"WIDSet"`

	// 自定义字段
	Field *CustomField `json:"Field,omitnil,omitempty" name:"Field"`

	// 分组id列表，非必填
	GroupIdSet []*int64 `json:"GroupIdSet,omitnil,omitempty" name:"GroupIdSet"`

	// 是否激活，默认全部，"1"激活，"0"未激活
	IsActive *string `json:"IsActive,omitnil,omitempty" name:"IsActive"`

	// 是否为摄像头，默认全部，"true"摄像头，"false"非摄像头
	IsCamera *string `json:"IsCamera,omitnil,omitempty" name:"IsCamera"`
}

type DescribeDeviceListRequest struct {
	*tchttp.BaseRequest
	
	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 分页查询，第几页，必传，大于0
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页条数，必传大于0
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 设备类型，非必填
	DeviceTypeSet []*string `json:"DeviceTypeSet,omitnil,omitempty" name:"DeviceTypeSet"`

	// 产品 pid，非必填
	ProductIdSet []*int64 `json:"ProductIdSet,omitnil,omitempty" name:"ProductIdSet"`

	// 设备标签，非必填
	TagIdSet []*int64 `json:"TagIdSet,omitnil,omitempty" name:"TagIdSet"`

	// 空间层级
	SpaceCodeSet []*string `json:"SpaceCodeSet,omitnil,omitempty" name:"SpaceCodeSet"`

	// 设备标签名，非必填
	DeviceTagSet []*string `json:"DeviceTagSet,omitnil,omitempty" name:"DeviceTagSet"`

	// 设备wid,非必填
	WIDSet []*string `json:"WIDSet,omitnil,omitempty" name:"WIDSet"`

	// 自定义字段
	Field *CustomField `json:"Field,omitnil,omitempty" name:"Field"`

	// 分组id列表，非必填
	GroupIdSet []*int64 `json:"GroupIdSet,omitnil,omitempty" name:"GroupIdSet"`

	// 是否激活，默认全部，"1"激活，"0"未激活
	IsActive *string `json:"IsActive,omitnil,omitempty" name:"IsActive"`

	// 是否为摄像头，默认全部，"true"摄像头，"false"非摄像头
	IsCamera *string `json:"IsCamera,omitnil,omitempty" name:"IsCamera"`
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
	delete(f, "WorkspaceId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "ApplicationToken")
	delete(f, "DeviceTypeSet")
	delete(f, "ProductIdSet")
	delete(f, "TagIdSet")
	delete(f, "SpaceCodeSet")
	delete(f, "DeviceTagSet")
	delete(f, "WIDSet")
	delete(f, "Field")
	delete(f, "GroupIdSet")
	delete(f, "IsActive")
	delete(f, "IsCamera")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceListRes struct {
	// 第几页
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页条数
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 总页数
	TotalPage *int64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// 总条数
	TotalRow *int64 `json:"TotalRow,omitnil,omitempty" name:"TotalRow"`

	// 设备信息集合
	DeviceDataSet []*DeviceDataInfo `json:"DeviceDataSet,omitnil,omitempty" name:"DeviceDataSet"`
}

// Predefined struct for user
type DescribeDeviceListResponseParams struct {
	// 查询设备列表结果
	Result *DescribeDeviceListRes `json:"Result,omitnil,omitempty" name:"Result"`

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
type DescribeDeviceShadowListRequestParams struct {
	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// WID
	WIDSet []*string `json:"WIDSet,omitnil,omitempty" name:"WIDSet"`

	// 分页查询，第几页
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页条数
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 设备类型code
	DeviceTypeSet []*string `json:"DeviceTypeSet,omitnil,omitempty" name:"DeviceTypeSet"`

	// 产品 pid
	ProductIdSet []*int64 `json:"ProductIdSet,omitnil,omitempty" name:"ProductIdSet"`

	// 设备标签id
	TagIdSet []*int64 `json:"TagIdSet,omitnil,omitempty" name:"TagIdSet"`

	// 空间层级，（支持空间多层，比如具体建筑、具体楼层）
	SpaceCodeSet []*string `json:"SpaceCodeSet,omitnil,omitempty" name:"SpaceCodeSet"`

	// 设备标签名
	DeviceTagSet []*string `json:"DeviceTagSet,omitnil,omitempty" name:"DeviceTagSet"`
}

type DescribeDeviceShadowListRequest struct {
	*tchttp.BaseRequest
	
	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// WID
	WIDSet []*string `json:"WIDSet,omitnil,omitempty" name:"WIDSet"`

	// 分页查询，第几页
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页条数
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 设备类型code
	DeviceTypeSet []*string `json:"DeviceTypeSet,omitnil,omitempty" name:"DeviceTypeSet"`

	// 产品 pid
	ProductIdSet []*int64 `json:"ProductIdSet,omitnil,omitempty" name:"ProductIdSet"`

	// 设备标签id
	TagIdSet []*int64 `json:"TagIdSet,omitnil,omitempty" name:"TagIdSet"`

	// 空间层级，（支持空间多层，比如具体建筑、具体楼层）
	SpaceCodeSet []*string `json:"SpaceCodeSet,omitnil,omitempty" name:"SpaceCodeSet"`

	// 设备标签名
	DeviceTagSet []*string `json:"DeviceTagSet,omitnil,omitempty" name:"DeviceTagSet"`
}

func (r *DescribeDeviceShadowListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceShadowListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "WIDSet")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "ApplicationToken")
	delete(f, "DeviceTypeSet")
	delete(f, "ProductIdSet")
	delete(f, "TagIdSet")
	delete(f, "SpaceCodeSet")
	delete(f, "DeviceTagSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceShadowListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceShadowListResponseParams struct {
	// 获取设备影子结果
	Result *DeviceShadowRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDeviceShadowListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceShadowListResponseParams `json:"Response"`
}

func (r *DescribeDeviceShadowListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceShadowListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceStatusListRequestParams struct {
	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 分页查询，第几页，必传，大于0
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页条数，必传大于0
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 设备类型
	DeviceTypeSet []*string `json:"DeviceTypeSet,omitnil,omitempty" name:"DeviceTypeSet"`

	// 产品 pid
	ProductIdSet []*int64 `json:"ProductIdSet,omitnil,omitempty" name:"ProductIdSet"`

	// 设备标签id
	TagIdSet []*int64 `json:"TagIdSet,omitnil,omitempty" name:"TagIdSet"`

	// 空间位置（支持空间多层，比如具体建筑、具体楼层）
	SpaceCodeSet []*string `json:"SpaceCodeSet,omitnil,omitempty" name:"SpaceCodeSet"`

	// 设备编号列表
	WIDSet []*string `json:"WIDSet,omitnil,omitempty" name:"WIDSet"`

	// 设备标签名，非必填
	DeviceTagSet []*string `json:"DeviceTagSet,omitnil,omitempty" name:"DeviceTagSet"`

	// 通信在/离线状态（online=normal+fault、offline）
	DeviceStatusSet []*string `json:"DeviceStatusSet,omitnil,omitempty" name:"DeviceStatusSet"`

	// 设备业务状态
	// （正常-normal、故障-fault、离线-offline）
	StatusSet []*string `json:"StatusSet,omitnil,omitempty" name:"StatusSet"`

	// 推流状态，推流中-true，未推流-false 仅摄像头有的状态
	IsAlive []*bool `json:"IsAlive,omitnil,omitempty" name:"IsAlive"`
}

type DescribeDeviceStatusListRequest struct {
	*tchttp.BaseRequest
	
	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 分页查询，第几页，必传，大于0
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页条数，必传大于0
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 设备类型
	DeviceTypeSet []*string `json:"DeviceTypeSet,omitnil,omitempty" name:"DeviceTypeSet"`

	// 产品 pid
	ProductIdSet []*int64 `json:"ProductIdSet,omitnil,omitempty" name:"ProductIdSet"`

	// 设备标签id
	TagIdSet []*int64 `json:"TagIdSet,omitnil,omitempty" name:"TagIdSet"`

	// 空间位置（支持空间多层，比如具体建筑、具体楼层）
	SpaceCodeSet []*string `json:"SpaceCodeSet,omitnil,omitempty" name:"SpaceCodeSet"`

	// 设备编号列表
	WIDSet []*string `json:"WIDSet,omitnil,omitempty" name:"WIDSet"`

	// 设备标签名，非必填
	DeviceTagSet []*string `json:"DeviceTagSet,omitnil,omitempty" name:"DeviceTagSet"`

	// 通信在/离线状态（online=normal+fault、offline）
	DeviceStatusSet []*string `json:"DeviceStatusSet,omitnil,omitempty" name:"DeviceStatusSet"`

	// 设备业务状态
	// （正常-normal、故障-fault、离线-offline）
	StatusSet []*string `json:"StatusSet,omitnil,omitempty" name:"StatusSet"`

	// 推流状态，推流中-true，未推流-false 仅摄像头有的状态
	IsAlive []*bool `json:"IsAlive,omitnil,omitempty" name:"IsAlive"`
}

func (r *DescribeDeviceStatusListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceStatusListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "ApplicationToken")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "DeviceTypeSet")
	delete(f, "ProductIdSet")
	delete(f, "TagIdSet")
	delete(f, "SpaceCodeSet")
	delete(f, "WIDSet")
	delete(f, "DeviceTagSet")
	delete(f, "DeviceStatusSet")
	delete(f, "StatusSet")
	delete(f, "IsAlive")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceStatusListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceStatusListResponseParams struct {
	// 查询设备状态结果
	Result *DeviceStatusRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDeviceStatusListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceStatusListResponseParams `json:"Response"`
}

func (r *DescribeDeviceStatusListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceStatusListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceStatusStatRequestParams struct {
	// 所属空间地理层级，必填。0表示查询所有层级（1、2）的设备状态，1表示楼栋，2表示楼层
	Level *int64 `json:"Level,omitnil,omitempty" name:"Level"`

	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 空间位置，非必填。为空表示查询所有（1，2）层级
	SpaceCodeSet []*string `json:"SpaceCodeSet,omitnil,omitempty" name:"SpaceCodeSet"`

	// 设备类型，非必填。为空表示查询所有设备类型
	DeviceTypeSet []*string `json:"DeviceTypeSet,omitnil,omitempty" name:"DeviceTypeSet"`
}

type DescribeDeviceStatusStatRequest struct {
	*tchttp.BaseRequest
	
	// 所属空间地理层级，必填。0表示查询所有层级（1、2）的设备状态，1表示楼栋，2表示楼层
	Level *int64 `json:"Level,omitnil,omitempty" name:"Level"`

	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 空间位置，非必填。为空表示查询所有（1，2）层级
	SpaceCodeSet []*string `json:"SpaceCodeSet,omitnil,omitempty" name:"SpaceCodeSet"`

	// 设备类型，非必填。为空表示查询所有设备类型
	DeviceTypeSet []*string `json:"DeviceTypeSet,omitnil,omitempty" name:"DeviceTypeSet"`
}

func (r *DescribeDeviceStatusStatRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceStatusStatRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Level")
	delete(f, "WorkspaceId")
	delete(f, "ApplicationToken")
	delete(f, "SpaceCodeSet")
	delete(f, "DeviceTypeSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceStatusStatRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceStatusStatResponseParams struct {
	// 设备状态统计结果
	Result *DeviceStatusStatRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDeviceStatusStatResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceStatusStatResponseParams `json:"Response"`
}

func (r *DescribeDeviceStatusStatResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceStatusStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceTagListRequestParams struct {
	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 分页查询，第几页，必传，大于0
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页条数，必传大于0
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

type DescribeDeviceTagListRequest struct {
	*tchttp.BaseRequest
	
	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 分页查询，第几页，必传，大于0
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页条数，必传大于0
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

func (r *DescribeDeviceTagListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceTagListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "ApplicationToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceTagListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceTagListResponseParams struct {
	// 设备标签查询结果
	Result *DeviceTagRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDeviceTagListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceTagListResponseParams `json:"Response"`
}

func (r *DescribeDeviceTagListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceTagListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceTypeListRequestParams struct {
	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 默认0只拉取设备列表关联的设备类型列表；1拉取空间下所有的设备类型列表
	Flag *int64 `json:"Flag,omitnil,omitempty" name:"Flag"`
}

type DescribeDeviceTypeListRequest struct {
	*tchttp.BaseRequest
	
	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 默认0只拉取设备列表关联的设备类型列表；1拉取空间下所有的设备类型列表
	Flag *int64 `json:"Flag,omitnil,omitempty" name:"Flag"`
}

func (r *DescribeDeviceTypeListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceTypeListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "ApplicationToken")
	delete(f, "Flag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceTypeListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceTypeListResponseParams struct {
	// 设备的设备类型列表
	Result *DeviceTypeSet `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDeviceTypeListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceTypeListResponseParams `json:"Response"`
}

func (r *DescribeDeviceTypeListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceTypeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeApplicationTokenRequestParams struct {
	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 是否刷新token，默认为false
	Refresh *bool `json:"Refresh,omitnil,omitempty" name:"Refresh"`
}

type DescribeEdgeApplicationTokenRequest struct {
	*tchttp.BaseRequest
	
	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 是否刷新token，默认为false
	Refresh *bool `json:"Refresh,omitnil,omitempty" name:"Refresh"`
}

func (r *DescribeEdgeApplicationTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeApplicationTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationToken")
	delete(f, "Refresh")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdgeApplicationTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeApplicationTokenResponseParams struct {
	// 边缘应用令牌信息	
	Result *ApplicationTokenInfo `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEdgeApplicationTokenResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdgeApplicationTokenResponseParams `json:"Response"`
}

func (r *DescribeEdgeApplicationTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeApplicationTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeElementProfilePageRequestParams struct {
	// 建筑id
	BuildingId *string `json:"BuildingId,omitnil,omitempty" name:"BuildingId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 页容量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 项目空间id
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 父级构件id
	ParentElementIds []*string `json:"ParentElementIds,omitnil,omitempty" name:"ParentElementIds"`

	// 空间层级
	Level *uint64 `json:"Level,omitnil,omitempty" name:"Level"`

	// 空间分类代码
	SpaceTypeCode *string `json:"SpaceTypeCode,omitnil,omitempty" name:"SpaceTypeCode"`

	// 构件类型
	EntityTypes []*string `json:"EntityTypes,omitnil,omitempty" name:"EntityTypes"`

	// 是否包含已删除构件
	IncludeDelete *bool `json:"IncludeDelete,omitnil,omitempty" name:"IncludeDelete"`

	// 时间范围-开始
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 时间范围-结束
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeElementProfilePageRequest struct {
	*tchttp.BaseRequest
	
	// 建筑id
	BuildingId *string `json:"BuildingId,omitnil,omitempty" name:"BuildingId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 页容量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 项目空间id
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 父级构件id
	ParentElementIds []*string `json:"ParentElementIds,omitnil,omitempty" name:"ParentElementIds"`

	// 空间层级
	Level *uint64 `json:"Level,omitnil,omitempty" name:"Level"`

	// 空间分类代码
	SpaceTypeCode *string `json:"SpaceTypeCode,omitnil,omitempty" name:"SpaceTypeCode"`

	// 构件类型
	EntityTypes []*string `json:"EntityTypes,omitnil,omitempty" name:"EntityTypes"`

	// 是否包含已删除构件
	IncludeDelete *bool `json:"IncludeDelete,omitnil,omitempty" name:"IncludeDelete"`

	// 时间范围-开始
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 时间范围-结束
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *DescribeElementProfilePageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeElementProfilePageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BuildingId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "WorkspaceId")
	delete(f, "ApplicationToken")
	delete(f, "ParentElementIds")
	delete(f, "Level")
	delete(f, "SpaceTypeCode")
	delete(f, "EntityTypes")
	delete(f, "IncludeDelete")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeElementProfilePageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeElementProfilePageResponseParams struct {
	// 分页查询构件出参
	Result *ElementProfilePageRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeElementProfilePageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeElementProfilePageResponseParams `json:"Response"`
}

func (r *DescribeElementProfilePageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeElementProfilePageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeElementProfileTreeRequestParams struct {
	// 建筑id
	BuildingId *string `json:"BuildingId,omitnil,omitempty" name:"BuildingId"`

	// 项目空间id
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 父级构件id
	ElementId *string `json:"ElementId,omitnil,omitempty" name:"ElementId"`

	// 构件级别
	Level *uint64 `json:"Level,omitnil,omitempty" name:"Level"`

	// 空间分类代码
	SpaceTypeCode *string `json:"SpaceTypeCode,omitnil,omitempty" name:"SpaceTypeCode"`
}

type DescribeElementProfileTreeRequest struct {
	*tchttp.BaseRequest
	
	// 建筑id
	BuildingId *string `json:"BuildingId,omitnil,omitempty" name:"BuildingId"`

	// 项目空间id
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 父级构件id
	ElementId *string `json:"ElementId,omitnil,omitempty" name:"ElementId"`

	// 构件级别
	Level *uint64 `json:"Level,omitnil,omitempty" name:"Level"`

	// 空间分类代码
	SpaceTypeCode *string `json:"SpaceTypeCode,omitnil,omitempty" name:"SpaceTypeCode"`
}

func (r *DescribeElementProfileTreeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeElementProfileTreeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BuildingId")
	delete(f, "WorkspaceId")
	delete(f, "ApplicationToken")
	delete(f, "ElementId")
	delete(f, "Level")
	delete(f, "SpaceTypeCode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeElementProfileTreeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeElementProfileTreeResponseParams struct {
	// 构件树出参
	Result *ElementProfileTreeRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeElementProfileTreeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeElementProfileTreeResponseParams `json:"Response"`
}

func (r *DescribeElementProfileTreeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeElementProfileTreeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEventListRequestParams struct {
	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 分页查询，第几页，必传，大于0
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页条数，必传大于0
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 事件触发类型，(app, device, timer)
	TriggerType *string `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// 事件id详情
	IdSet []*int64 `json:"IdSet,omitnil,omitempty" name:"IdSet"`
}

type DescribeEventListRequest struct {
	*tchttp.BaseRequest
	
	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 分页查询，第几页，必传，大于0
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页条数，必传大于0
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 事件触发类型，(app, device, timer)
	TriggerType *string `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// 事件id详情
	IdSet []*int64 `json:"IdSet,omitnil,omitempty" name:"IdSet"`
}

func (r *DescribeEventListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEventListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "ApplicationToken")
	delete(f, "TriggerType")
	delete(f, "IdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEventListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEventListRes struct {
	// 第几页
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页条数
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 总页数
	TotalPage *int64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// 总条数
	TotalRow *int64 `json:"TotalRow,omitnil,omitempty" name:"TotalRow"`

	// 事件信息列表
	EventDetailSet []*EventDetail `json:"EventDetailSet,omitnil,omitempty" name:"EventDetailSet"`
}

// Predefined struct for user
type DescribeEventListResponseParams struct {
	// 事件列表查询结果
	Result *DescribeEventListRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEventListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEventListResponseParams `json:"Response"`
}

func (r *DescribeEventListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEventListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileDownloadURLRequestParams struct {
	// 工作空间Id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 文件Id
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

type DescribeFileDownloadURLRequest struct {
	*tchttp.BaseRequest
	
	// 工作空间Id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 文件Id
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

func (r *DescribeFileDownloadURLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileDownloadURLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "FileId")
	delete(f, "ApplicationToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFileDownloadURLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileDownloadURLResponseParams struct {
	// 文件下载URL地址
	Result *FileDownloadURL `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFileDownloadURLResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFileDownloadURLResponseParams `json:"Response"`
}

func (r *DescribeFileDownloadURLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileDownloadURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileUploadURLRequestParams struct {
	// 工作空间Id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 文件名称
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文件大小
	FileSize *int64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 文件MD5
	FileMD5 *string `json:"FileMD5,omitnil,omitempty" name:"FileMD5"`

	// 存储类型
	SaveType *string `json:"SaveType,omitnil,omitempty" name:"SaveType"`

	// 过期时间，过期时间戳，精确到秒的时间戳，noExpireFlag为false时必填
	FileExpireTime *int64 `json:"FileExpireTime,omitnil,omitempty" name:"FileExpireTime"`

	// 永不过期标记
	NoExpireFlag *bool `json:"NoExpireFlag,omitnil,omitempty" name:"NoExpireFlag"`
}

type DescribeFileUploadURLRequest struct {
	*tchttp.BaseRequest
	
	// 工作空间Id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 文件名称
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文件大小
	FileSize *int64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 文件MD5
	FileMD5 *string `json:"FileMD5,omitnil,omitempty" name:"FileMD5"`

	// 存储类型
	SaveType *string `json:"SaveType,omitnil,omitempty" name:"SaveType"`

	// 过期时间，过期时间戳，精确到秒的时间戳，noExpireFlag为false时必填
	FileExpireTime *int64 `json:"FileExpireTime,omitnil,omitempty" name:"FileExpireTime"`

	// 永不过期标记
	NoExpireFlag *bool `json:"NoExpireFlag,omitnil,omitempty" name:"NoExpireFlag"`
}

func (r *DescribeFileUploadURLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileUploadURLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "FileName")
	delete(f, "FileSize")
	delete(f, "ApplicationToken")
	delete(f, "FileMD5")
	delete(f, "SaveType")
	delete(f, "FileExpireTime")
	delete(f, "NoExpireFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFileUploadURLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileUploadURLResponseParams struct {
	// 获取文件上传地址结果
	Result *FileUploadURL `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFileUploadURLResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFileUploadURLResponseParams `json:"Response"`
}

func (r *DescribeFileUploadURLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileUploadURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupInfo struct {
	// 分组
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 设备分组名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分组描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 分组父级ID
	ParentId *int64 `json:"ParentId,omitnil,omitempty" name:"ParentId"`
}

// Predefined struct for user
type DescribeInterfaceListRequestParams struct {
	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 请求页号
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 请求页容量，默认全量返回
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询关键字
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 接口方式 1.http 2消息通知服务
	Style []*uint64 `json:"Style,omitnil,omitempty" name:"Style"`

	// 接口分类0. 其他服务 1. IOT服务 2. 空间服务 3.微应用服务 4.场景服务 5.AI算法服务 6.任务算法服务 7.第三方服务 8.3DTiles服务
	Type []*uint64 `json:"Type,omitnil,omitempty" name:"Type"`
}

type DescribeInterfaceListRequest struct {
	*tchttp.BaseRequest
	
	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 请求页号
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 请求页容量，默认全量返回
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询关键字
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 接口方式 1.http 2消息通知服务
	Style []*uint64 `json:"Style,omitnil,omitempty" name:"Style"`

	// 接口分类0. 其他服务 1. IOT服务 2. 空间服务 3.微应用服务 4.场景服务 5.AI算法服务 6.任务算法服务 7.第三方服务 8.3DTiles服务
	Type []*uint64 `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *DescribeInterfaceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInterfaceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationToken")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Keyword")
	delete(f, "Style")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInterfaceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInterfaceListResponseParams struct {
	// API列表
	Result *ApiInfoList `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInterfaceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInterfaceListResponseParams `json:"Response"`
}

func (r *DescribeInterfaceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInterfaceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLinkRuleListRequestParams struct {
	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 分页查询，第几页，必传，大于0
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页条数，必传大于0
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 事件触发类型
	TriggerType *string `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// 联动id
	IdSet []*int64 `json:"IdSet,omitnil,omitempty" name:"IdSet"`
}

type DescribeLinkRuleListRequest struct {
	*tchttp.BaseRequest
	
	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 分页查询，第几页，必传，大于0
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页条数，必传大于0
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 事件触发类型
	TriggerType *string `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// 联动id
	IdSet []*int64 `json:"IdSet,omitnil,omitempty" name:"IdSet"`
}

func (r *DescribeLinkRuleListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLinkRuleListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "ApplicationToken")
	delete(f, "TriggerType")
	delete(f, "IdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLinkRuleListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLinkRuleListRes struct {
	// 第几页
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页条数
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 总页数
	TotalPage *int64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// 总条数
	TotalRow *int64 `json:"TotalRow,omitnil,omitempty" name:"TotalRow"`

	// 联动规则列表
	LinkRuleSet []*LinkRuleInfo `json:"LinkRuleSet,omitnil,omitempty" name:"LinkRuleSet"`
}

// Predefined struct for user
type DescribeLinkRuleListResponseParams struct {
	// 联动规则列表查询结果
	Result *DescribeLinkRuleListRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLinkRuleListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLinkRuleListResponseParams `json:"Response"`
}

func (r *DescribeLinkRuleListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLinkRuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelListRequestParams struct {
	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 分页查询，第几页，大于0
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页条数，大于0
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 设备类型
	DeviceTypeSet []*string `json:"DeviceTypeSet,omitnil,omitempty" name:"DeviceTypeSet"`

	// 产品 pid
	ProductIdSet []*int64 `json:"ProductIdSet,omitnil,omitempty" name:"ProductIdSet"`

	// 模型id
	ModelIdSet []*string `json:"ModelIdSet,omitnil,omitempty" name:"ModelIdSet"`
}

type DescribeModelListRequest struct {
	*tchttp.BaseRequest
	
	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 分页查询，第几页，大于0
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页条数，大于0
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 设备类型
	DeviceTypeSet []*string `json:"DeviceTypeSet,omitnil,omitempty" name:"DeviceTypeSet"`

	// 产品 pid
	ProductIdSet []*int64 `json:"ProductIdSet,omitnil,omitempty" name:"ProductIdSet"`

	// 模型id
	ModelIdSet []*string `json:"ModelIdSet,omitnil,omitempty" name:"ModelIdSet"`
}

func (r *DescribeModelListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "ApplicationToken")
	delete(f, "DeviceTypeSet")
	delete(f, "ProductIdSet")
	delete(f, "ModelIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelListResponseParams struct {
	// 模型列表查询结果
	Result *ModelSet `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModelListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelListResponseParams `json:"Response"`
}

func (r *DescribeModelListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductListRequestParams struct {
	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 分页查询，第几页
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页条数，大于0
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 设备类型
	DeviceTypeSet []*string `json:"DeviceTypeSet,omitnil,omitempty" name:"DeviceTypeSet"`

	// 产品 pid
	ProductIdSet []*int64 `json:"ProductIdSet,omitnil,omitempty" name:"ProductIdSet"`

	// 模型id
	ModelIdSet []*string `json:"ModelIdSet,omitnil,omitempty" name:"ModelIdSet"`
}

type DescribeProductListRequest struct {
	*tchttp.BaseRequest
	
	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 分页查询，第几页
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页条数，大于0
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 设备类型
	DeviceTypeSet []*string `json:"DeviceTypeSet,omitnil,omitempty" name:"DeviceTypeSet"`

	// 产品 pid
	ProductIdSet []*int64 `json:"ProductIdSet,omitnil,omitempty" name:"ProductIdSet"`

	// 模型id
	ModelIdSet []*string `json:"ModelIdSet,omitnil,omitempty" name:"ModelIdSet"`
}

func (r *DescribeProductListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "ApplicationToken")
	delete(f, "DeviceTypeSet")
	delete(f, "ProductIdSet")
	delete(f, "ModelIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProductListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductListResponseParams struct {
	// 产品列表查询结果
	Result *ProductSet `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProductListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProductListResponseParams `json:"Response"`
}

func (r *DescribeProductListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePropertyListRequestParams struct {
	// 建筑id
	BuildingId *string `json:"BuildingId,omitnil,omitempty" name:"BuildingId"`

	// 构件id
	ElementId *string `json:"ElementId,omitnil,omitempty" name:"ElementId"`

	// 项目空间id
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

type DescribePropertyListRequest struct {
	*tchttp.BaseRequest
	
	// 建筑id
	BuildingId *string `json:"BuildingId,omitnil,omitempty" name:"BuildingId"`

	// 构件id
	ElementId *string `json:"ElementId,omitnil,omitempty" name:"ElementId"`

	// 项目空间id
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

func (r *DescribePropertyListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePropertyListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BuildingId")
	delete(f, "ElementId")
	delete(f, "WorkspaceId")
	delete(f, "ApplicationToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePropertyListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePropertyListResponseParams struct {
	// 构件属性信息出参
	Result *ElementPropertyRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePropertyListResponse struct {
	*tchttp.BaseResponse
	Response *DescribePropertyListResponseParams `json:"Response"`
}

func (r *DescribePropertyListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePropertyListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleDetailRequestParams struct {
	// 工作空间id
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 联动id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

type DescribeRuleDetailRequest struct {
	*tchttp.BaseRequest
	
	// 工作空间id
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 联动id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

func (r *DescribeRuleDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "Id")
	delete(f, "ApplicationToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRuleDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleDetailResponseParams struct {
	// 规则详情查询结果
	Result *RuleDetailRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRuleDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRuleDetailResponseParams `json:"Response"`
}

func (r *DescribeRuleDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSceneListRequestParams struct {
	// 项目空间id
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

type DescribeSceneListRequest struct {
	*tchttp.BaseRequest
	
	// 项目空间id
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

func (r *DescribeSceneListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSceneListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "ApplicationToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSceneListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSceneListResponseParams struct {
	// 场景列表出参
	Result *SceneListRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSceneListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSceneListResponseParams `json:"Response"`
}

func (r *DescribeSceneListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSceneListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpaceDeviceIdListRequestParams struct {
	// 构件id列表
	ElementIds []*string `json:"ElementIds,omitnil,omitempty" name:"ElementIds"`

	// 是否级联
	IsCascade *bool `json:"IsCascade,omitnil,omitempty" name:"IsCascade"`

	// 项目空间id
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 页容量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

type DescribeSpaceDeviceIdListRequest struct {
	*tchttp.BaseRequest
	
	// 构件id列表
	ElementIds []*string `json:"ElementIds,omitnil,omitempty" name:"ElementIds"`

	// 是否级联
	IsCascade *bool `json:"IsCascade,omitnil,omitempty" name:"IsCascade"`

	// 项目空间id
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 页容量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

func (r *DescribeSpaceDeviceIdListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpaceDeviceIdListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ElementIds")
	delete(f, "IsCascade")
	delete(f, "WorkspaceId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "ApplicationToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSpaceDeviceIdListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpaceDeviceIdListResponseParams struct {
	// 设备ID列表
	Result *SpaceDeviceIdListRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSpaceDeviceIdListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSpaceDeviceIdListResponseParams `json:"Response"`
}

func (r *DescribeSpaceDeviceIdListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpaceDeviceIdListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpaceDeviceRelationListRequestParams struct {
	// 构件id列表
	ElementIds []*string `json:"ElementIds,omitnil,omitempty" name:"ElementIds"`

	// 是否级联
	IsCascade *bool `json:"IsCascade,omitnil,omitempty" name:"IsCascade"`

	// 项目空间id
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 页容量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

type DescribeSpaceDeviceRelationListRequest struct {
	*tchttp.BaseRequest
	
	// 构件id列表
	ElementIds []*string `json:"ElementIds,omitnil,omitempty" name:"ElementIds"`

	// 是否级联
	IsCascade *bool `json:"IsCascade,omitnil,omitempty" name:"IsCascade"`

	// 项目空间id
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 页容量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

func (r *DescribeSpaceDeviceRelationListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpaceDeviceRelationListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ElementIds")
	delete(f, "IsCascade")
	delete(f, "WorkspaceId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "ApplicationToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSpaceDeviceRelationListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpaceDeviceRelationListResponseParams struct {
	// 查询指定空间下设备与构件绑定关系列表出参
	Result *SpaceDeviceRelationRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSpaceDeviceRelationListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSpaceDeviceRelationListResponseParams `json:"Response"`
}

func (r *DescribeSpaceDeviceRelationListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpaceDeviceRelationListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpaceInfoByDeviceIdRequestParams struct {
	// 设备id
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 项目空间id
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

type DescribeSpaceInfoByDeviceIdRequest struct {
	*tchttp.BaseRequest
	
	// 设备id
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 项目空间id
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

func (r *DescribeSpaceInfoByDeviceIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpaceInfoByDeviceIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	delete(f, "WorkspaceId")
	delete(f, "ApplicationToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSpaceInfoByDeviceIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpaceInfoByDeviceIdResponseParams struct {
	// 设备绑定的空间信息出参
	Result *DeviceSpaceInfoRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSpaceInfoByDeviceIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSpaceInfoByDeviceIdResponseParams `json:"Response"`
}

func (r *DescribeSpaceInfoByDeviceIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpaceInfoByDeviceIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpaceRelationByDeviceIdRequestParams struct {
	// 设备id
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 工作空间id
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

type DescribeSpaceRelationByDeviceIdRequest struct {
	*tchttp.BaseRequest
	
	// 设备id
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 工作空间id
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

func (r *DescribeSpaceRelationByDeviceIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpaceRelationByDeviceIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	delete(f, "WorkspaceId")
	delete(f, "ApplicationToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSpaceRelationByDeviceIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpaceRelationByDeviceIdResponseParams struct {
	// 空间层级关系出参
	Result *SpaceRelationRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSpaceRelationByDeviceIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSpaceRelationByDeviceIdResponseParams `json:"Response"`
}

func (r *DescribeSpaceRelationByDeviceIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpaceRelationByDeviceIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpaceTypeListRequestParams struct {
	// 项目空间ID
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 页容量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

type DescribeSpaceTypeListRequest struct {
	*tchttp.BaseRequest
	
	// 项目空间ID
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 页容量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

func (r *DescribeSpaceTypeListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpaceTypeListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "ApplicationToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSpaceTypeListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpaceTypeListResponseParams struct {
	// 空间分类列表出参
	Result *SpaceTypeListRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSpaceTypeListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSpaceTypeListResponseParams `json:"Response"`
}

func (r *DescribeSpaceTypeListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpaceTypeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTenantBuildingCountAndAreaRequestParams struct {
	// 租户所有工作空间ID列表
	WorkspaceIdList []*string `json:"WorkspaceIdList,omitnil,omitempty" name:"WorkspaceIdList"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

type DescribeTenantBuildingCountAndAreaRequest struct {
	*tchttp.BaseRequest
	
	// 租户所有工作空间ID列表
	WorkspaceIdList []*string `json:"WorkspaceIdList,omitnil,omitempty" name:"WorkspaceIdList"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

func (r *DescribeTenantBuildingCountAndAreaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTenantBuildingCountAndAreaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceIdList")
	delete(f, "ApplicationToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTenantBuildingCountAndAreaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTenantBuildingCountAndAreaResponseParams struct {
	// 租户所有项目空间楼栋数量与建筑面积统计结果
	Result *SpaceDataTotalStatsRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTenantBuildingCountAndAreaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTenantBuildingCountAndAreaResponseParams `json:"Response"`
}

func (r *DescribeTenantBuildingCountAndAreaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTenantBuildingCountAndAreaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTenantDepartmentListRequestParams struct {
	// 翻页页码
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 翻页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 租户ID
	TenantId *string `json:"TenantId,omitnil,omitempty" name:"TenantId"`

	// 更新时间戳，单位秒
	UpdateAt *int64 `json:"UpdateAt,omitnil,omitempty" name:"UpdateAt"`

	// 部门ID
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`

	// 用户id
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`
}

type DescribeTenantDepartmentListRequest struct {
	*tchttp.BaseRequest
	
	// 翻页页码
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 翻页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 租户ID
	TenantId *string `json:"TenantId,omitnil,omitempty" name:"TenantId"`

	// 更新时间戳，单位秒
	UpdateAt *int64 `json:"UpdateAt,omitnil,omitempty" name:"UpdateAt"`

	// 部门ID
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`

	// 用户id
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`
}

func (r *DescribeTenantDepartmentListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTenantDepartmentListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ApplicationToken")
	delete(f, "TenantId")
	delete(f, "UpdateAt")
	delete(f, "DepartmentId")
	delete(f, "Cursor")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTenantDepartmentListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTenantDepartmentListResponseParams struct {
	// 返回数据
	Result *SsoDepartmentsResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTenantDepartmentListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTenantDepartmentListResponseParams `json:"Response"`
}

func (r *DescribeTenantDepartmentListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTenantDepartmentListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTenantUserListRequestParams struct {
	// 翻页页码
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 翻页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 租户ID
	TenantId *string `json:"TenantId,omitnil,omitempty" name:"TenantId"`

	// 更新时间戳，单位秒
	UpdateAt *int64 `json:"UpdateAt,omitnil,omitempty" name:"UpdateAt"`

	// 部门ID
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`

	// 用户id
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// 状态，0，获取所有数据，1正常启用，2禁用
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 项目空间id
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 关键词
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 是否递归获取子级数据，0需要，1不需要，默认为0
	NoRecursive *string `json:"NoRecursive,omitnil,omitempty" name:"NoRecursive"`
}

type DescribeTenantUserListRequest struct {
	*tchttp.BaseRequest
	
	// 翻页页码
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 翻页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 租户ID
	TenantId *string `json:"TenantId,omitnil,omitempty" name:"TenantId"`

	// 更新时间戳，单位秒
	UpdateAt *int64 `json:"UpdateAt,omitnil,omitempty" name:"UpdateAt"`

	// 部门ID
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`

	// 用户id
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// 状态，0，获取所有数据，1正常启用，2禁用
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 项目空间id
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 关键词
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 是否递归获取子级数据，0需要，1不需要，默认为0
	NoRecursive *string `json:"NoRecursive,omitnil,omitempty" name:"NoRecursive"`
}

func (r *DescribeTenantUserListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTenantUserListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ApplicationToken")
	delete(f, "TenantId")
	delete(f, "UpdateAt")
	delete(f, "DepartmentId")
	delete(f, "Cursor")
	delete(f, "Status")
	delete(f, "WorkspaceId")
	delete(f, "Keyword")
	delete(f, "NoRecursive")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTenantUserListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTenantUserListResponseParams struct {
	// 返回数据
	Result *SsoUserResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTenantUserListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTenantUserListResponseParams `json:"Response"`
}

func (r *DescribeTenantUserListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTenantUserListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoCloudRecordRequestParams struct {
	// 设备的唯一标识
	WID *string `json:"WID,omitnil,omitempty" name:"WID"`

	// 录像开始时间（s）
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 录像结束时间（s）
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 工作空间Id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

type DescribeVideoCloudRecordRequest struct {
	*tchttp.BaseRequest
	
	// 设备的唯一标识
	WID *string `json:"WID,omitnil,omitempty" name:"WID"`

	// 录像开始时间（s）
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 录像结束时间（s）
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 工作空间Id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

func (r *DescribeVideoCloudRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoCloudRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WID")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "WorkspaceId")
	delete(f, "ApplicationToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVideoCloudRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoCloudRecordResponseParams struct {
	// 获取云录像结果
	Result *VideoCloudRecordRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVideoCloudRecordResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVideoCloudRecordResponseParams `json:"Response"`
}

func (r *DescribeVideoCloudRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoCloudRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoLiveStreamRequestParams struct {
	// 设备的唯一标识
	WID *string `json:"WID,omitnil,omitempty" name:"WID"`

	// 枚举如下：
	// flv
	// rtmp
	// hls
	// webrtc
	// raw (视频原始帧)
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 工作空间Id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 主码流传0，子码流传1，默认主码流
	StreamId *int64 `json:"StreamId,omitnil,omitempty" name:"StreamId"`

	// 设备所在环境，公有云私有化项目传0或者不传，混合云项目一般传空间id
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

type DescribeVideoLiveStreamRequest struct {
	*tchttp.BaseRequest
	
	// 设备的唯一标识
	WID *string `json:"WID,omitnil,omitempty" name:"WID"`

	// 枚举如下：
	// flv
	// rtmp
	// hls
	// webrtc
	// raw (视频原始帧)
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 工作空间Id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 主码流传0，子码流传1，默认主码流
	StreamId *int64 `json:"StreamId,omitnil,omitempty" name:"StreamId"`

	// 设备所在环境，公有云私有化项目传0或者不传，混合云项目一般传空间id
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

func (r *DescribeVideoLiveStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoLiveStreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WID")
	delete(f, "Protocol")
	delete(f, "WorkspaceId")
	delete(f, "ApplicationToken")
	delete(f, "StreamId")
	delete(f, "Env")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVideoLiveStreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoLiveStreamResponseParams struct {
	// 视频实时流获取结果
	Result *VideoRecordStreamRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVideoLiveStreamResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVideoLiveStreamResponseParams `json:"Response"`
}

func (r *DescribeVideoLiveStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoLiveStreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoRecordStreamRequestParams struct {
	// 设备唯一标识
	WID *string `json:"WID,omitnil,omitempty" name:"WID"`

	// 枚举如下：
	// flvsh
	// rtmp
	// hls
	// webrtc
	// raw (视频原始帧)
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 开始时间（精确到毫秒）
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间（精确到毫秒）
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 倍速 0.5、1、2、4
	PlayBackRate *float64 `json:"PlayBackRate,omitnil,omitempty" name:"PlayBackRate"`

	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 流的唯一标识，播放链接尾缀
	Stream *string `json:"Stream,omitnil,omitempty" name:"Stream"`

	// 公有云私有化项目传0或者不传；混合云项目一般传空间id
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

type DescribeVideoRecordStreamRequest struct {
	*tchttp.BaseRequest
	
	// 设备唯一标识
	WID *string `json:"WID,omitnil,omitempty" name:"WID"`

	// 枚举如下：
	// flvsh
	// rtmp
	// hls
	// webrtc
	// raw (视频原始帧)
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 开始时间（精确到毫秒）
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间（精确到毫秒）
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 倍速 0.5、1、2、4
	PlayBackRate *float64 `json:"PlayBackRate,omitnil,omitempty" name:"PlayBackRate"`

	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 流的唯一标识，播放链接尾缀
	Stream *string `json:"Stream,omitnil,omitempty" name:"Stream"`

	// 公有云私有化项目传0或者不传；混合云项目一般传空间id
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

func (r *DescribeVideoRecordStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoRecordStreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WID")
	delete(f, "Protocol")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PlayBackRate")
	delete(f, "WorkspaceId")
	delete(f, "ApplicationToken")
	delete(f, "Stream")
	delete(f, "Env")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVideoRecordStreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoRecordStreamResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVideoRecordStreamResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVideoRecordStreamResponseParams `json:"Response"`
}

func (r *DescribeVideoRecordStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoRecordStreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkSpaceBuildingCountAndAreaRequestParams struct {
	// 工作空间ID列表
	WorkspaceIdList []*string `json:"WorkspaceIdList,omitnil,omitempty" name:"WorkspaceIdList"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

type DescribeWorkSpaceBuildingCountAndAreaRequest struct {
	*tchttp.BaseRequest
	
	// 工作空间ID列表
	WorkspaceIdList []*string `json:"WorkspaceIdList,omitnil,omitempty" name:"WorkspaceIdList"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

func (r *DescribeWorkSpaceBuildingCountAndAreaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkSpaceBuildingCountAndAreaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceIdList")
	delete(f, "ApplicationToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWorkSpaceBuildingCountAndAreaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkSpaceBuildingCountAndAreaResponseParams struct {
	// 查询项目空间楼栋数量与建筑面积出参
	Result *SpaceDataListStatsRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWorkSpaceBuildingCountAndAreaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWorkSpaceBuildingCountAndAreaResponseParams `json:"Response"`
}

func (r *DescribeWorkSpaceBuildingCountAndAreaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkSpaceBuildingCountAndAreaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkspaceListRequestParams struct {
	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 工作空间id，非必填，填了则表示根据id进行批量查询
	WorkspaceId *uint64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`
}

type DescribeWorkspaceListRequest struct {
	*tchttp.BaseRequest
	
	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 工作空间id，非必填，填了则表示根据id进行批量查询
	WorkspaceId *uint64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`
}

func (r *DescribeWorkspaceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkspaceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationToken")
	delete(f, "WorkspaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWorkspaceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkspaceListResponseParams struct {
	// 项目空间列表
	Result *WorkspaceInfoList `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWorkspaceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWorkspaceListResponseParams `json:"Response"`
}

func (r *DescribeWorkspaceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkspaceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkspaceUserListRequestParams struct {
	// 翻页页码
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 翻页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 工作空间ID
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 租户ID
	TenantId *string `json:"TenantId,omitnil,omitempty" name:"TenantId"`

	// 更新时间戳，单位秒
	UpdateAt *int64 `json:"UpdateAt,omitnil,omitempty" name:"UpdateAt"`
}

type DescribeWorkspaceUserListRequest struct {
	*tchttp.BaseRequest
	
	// 翻页页码
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 翻页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 工作空间ID
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 租户ID
	TenantId *string `json:"TenantId,omitnil,omitempty" name:"TenantId"`

	// 更新时间戳，单位秒
	UpdateAt *int64 `json:"UpdateAt,omitnil,omitempty" name:"UpdateAt"`
}

func (r *DescribeWorkspaceUserListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkspaceUserListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "WorkspaceId")
	delete(f, "ApplicationToken")
	delete(f, "TenantId")
	delete(f, "UpdateAt")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWorkspaceUserListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkspaceUserListResponseParams struct {
	// 返回数据
	Result *SsoTeamUserResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWorkspaceUserListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWorkspaceUserListResponseParams `json:"Response"`
}

func (r *DescribeWorkspaceUserListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkspaceUserListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeviceDataInfo struct {
	// 设备ID， wid
	WID *string `json:"WID,omitnil,omitempty" name:"WID"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 设备类型Id
	DeviceTypeCode *string `json:"DeviceTypeCode,omitnil,omitempty" name:"DeviceTypeCode"`

	// 设备类型名称
	DeviceTypeName *string `json:"DeviceTypeName,omitnil,omitempty" name:"DeviceTypeName"`

	// 产品Id
	ProductId *int64 `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 产品名称
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 产品能力:信令数据、音视频。二进制数值中第0位表示信令数据、第1位表示音视频 。1（信令数据），3（具有信令数据以及音视频能力）。
	ProductAbility *int64 `json:"ProductAbility,omitnil,omitempty" name:"ProductAbility"`

	// 设备位置信息
	SpaceInfoSet []*DeviceSpaceInfo `json:"SpaceInfoSet,omitnil,omitempty" name:"SpaceInfoSet"`

	// 模型id
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 模型名称
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 设备标签名，非必填
	DeviceTagSet []*string `json:"DeviceTagSet,omitnil,omitempty" name:"DeviceTagSet"`

	// 激活状态（1激活、0未激活）
	IsActive *int64 `json:"IsActive,omitnil,omitempty" name:"IsActive"`

	//  激活时间
	ActiveTime *string `json:"ActiveTime,omitnil,omitempty" name:"ActiveTime"`

	// 推流状态（推流中、未推流） 仅摄像机有的状态
	IsLive *bool `json:"IsLive,omitnil,omitempty" name:"IsLive"`

	// 设备所属父设备id（子设备才有）
	ParentWID *string `json:"ParentWID,omitnil,omitempty" name:"ParentWID"`

	// 设备所有父设备名称（子设备才有）
	ParentWIDName *string `json:"ParentWIDName,omitnil,omitempty" name:"ParentWIDName"`

	// 序列号
	SN *string `json:"SN,omitnil,omitempty" name:"SN"`

	// 设备点位坐标值
	Location *DeviceLocation `json:"Location,omitnil,omitempty" name:"Location"`

	// 自定义字段
	FieldList []*CustomFieldInfo `json:"FieldList,omitnil,omitempty" name:"FieldList"`

	// 分组信息
	GroupInfo *string `json:"GroupInfo,omitnil,omitempty" name:"GroupInfo"`

	// 通信在/离线状态（online=normal+fault，offline）
	DeviceStatus *string `json:"DeviceStatus,omitnil,omitempty" name:"DeviceStatus"`

	// 设备业务状态（normal、fault、offline）
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type DeviceLocation struct {
	// 点位X坐标值
	X *float64 `json:"X,omitnil,omitempty" name:"X"`

	// 点位Y坐标值
	Y *float64 `json:"Y,omitnil,omitempty" name:"Y"`

	// 点位Z坐标值
	Z *float64 `json:"Z,omitnil,omitempty" name:"Z"`
}

type DeviceModifyInfo struct {
	// 设备id
	WID *string `json:"WID,omitnil,omitempty" name:"WID"`

	// 修改后的设备名字
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`
}

type DeviceShadowInfo struct {
	// 设备ID
	WID *string `json:"WID,omitnil,omitempty" name:"WID"`

	// 设备影子数据,返回有效数据为"x-json:"后字段
	DeviceShadow *string `json:"DeviceShadow,omitnil,omitempty" name:"DeviceShadow"`

	// 设备影子更新时间
	DeviceShadowUpdateTime *string `json:"DeviceShadowUpdateTime,omitnil,omitempty" name:"DeviceShadowUpdateTime"`
}

type DeviceShadowRes struct {
	// 第几页
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页条数
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 总页数
	TotalPage *int64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// 总条数
	TotalRow *int64 `json:"TotalRow,omitnil,omitempty" name:"TotalRow"`

	// 设备影子列表
	Set []*DeviceShadowInfo `json:"Set,omitnil,omitempty" name:"Set"`
}

type DeviceSpaceInfo struct {
	// 空间Id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 空间名字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 空间级别
	Level *int64 `json:"Level,omitnil,omitempty" name:"Level"`

	// 空间编码
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`
}

type DeviceSpaceInfoRes struct {
	// 建筑id
	BuildingId *string `json:"BuildingId,omitnil,omitempty" name:"BuildingId"`

	// 构件id
	ElementId *string `json:"ElementId,omitnil,omitempty" name:"ElementId"`

	// 构件类型
	EntityType *string `json:"EntityType,omitnil,omitempty" name:"EntityType"`

	// 构件名称
	ElementName *string `json:"ElementName,omitnil,omitempty" name:"ElementName"`

	// 构件级别
	Level *uint64 `json:"Level,omitnil,omitempty" name:"Level"`

	// 底部标高（单位mm）
	BottomHeight *int64 `json:"BottomHeight,omitnil,omitempty" name:"BottomHeight"`

	// 空间编码
	SpaceCode *string `json:"SpaceCode,omitnil,omitempty" name:"SpaceCode"`
}

type DeviceStatusInfo struct {
	// 设备ID
	WID *string `json:"WID,omitnil,omitempty" name:"WID"`

	// 设备状态（online=normal+fault、offline）
	DeviceStatus *string `json:"DeviceStatus,omitnil,omitempty" name:"DeviceStatus"`

	// 设备状态更新时间
	DeviceStatusUpdateTime *string `json:"DeviceStatusUpdateTime,omitnil,omitempty" name:"DeviceStatusUpdateTime"`

	// 设备业务状态（normal、fault、offline）
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 推流状态。推流中-true，未推流-false
	IsAlive *bool `json:"IsAlive,omitnil,omitempty" name:"IsAlive"`
}

type DeviceStatusRes struct {
	// 第几页
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页条数
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 总页数
	TotalPage *int64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// 总条数
	TotalRow *int64 `json:"TotalRow,omitnil,omitempty" name:"TotalRow"`

	// 设备状态信息列表
	DeviceStatusSet []*DeviceStatusInfo `json:"DeviceStatusSet,omitnil,omitempty" name:"DeviceStatusSet"`
}

type DeviceStatusStatRes struct {
	// 工作空间Id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 汇总数。在线（正常+故障） + 离线
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 正常数
	NormalSum *int64 `json:"NormalSum,omitnil,omitempty" name:"NormalSum"`

	// 离线数
	OfflineSum *int64 `json:"OfflineSum,omitnil,omitempty" name:"OfflineSum"`

	// 故障数
	FaultSum *int64 `json:"FaultSum,omitnil,omitempty" name:"FaultSum"`

	// 设备类型概览列表
	DeviceTypeOverviewSet []*DeviceTypeOverview `json:"DeviceTypeOverviewSet,omitnil,omitempty" name:"DeviceTypeOverviewSet"`

	// 设备类型统计列表
	StatLevelSet []*StatLevel `json:"StatLevelSet,omitnil,omitempty" name:"StatLevelSet"`
}

type DeviceTagInfo struct {
	// 标签Id
	TagId *int64 `json:"TagId,omitnil,omitempty" name:"TagId"`

	// 标签名字
	TagName *string `json:"TagName,omitnil,omitempty" name:"TagName"`
}

type DeviceTagRes struct {
	// 第几页
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页条数
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 总页数
	TotalPage *int64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// 总条数
	TotalRow *int64 `json:"TotalRow,omitnil,omitempty" name:"TotalRow"`

	// 设备标签列表
	Set []*DeviceTagInfo `json:"Set,omitnil,omitempty" name:"Set"`
}

type DeviceType struct {
	// 设备类型编码
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 设备类型名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 父设备类型编码
	ParentCode *string `json:"ParentCode,omitnil,omitempty" name:"ParentCode"`

	// 父设备类型名称
	ParentName *string `json:"ParentName,omitnil,omitempty" name:"ParentName"`

	// 是否子系统，1是
	IsSubsystem *int64 `json:"IsSubsystem,omitnil,omitempty" name:"IsSubsystem"`
}

type DeviceTypeOverview struct {
	// 设备类型值
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 设备类型名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 汇总数。在线（正常+故障） + 离线
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 正常数
	Normal *int64 `json:"Normal,omitnil,omitempty" name:"Normal"`

	// 离线数
	Offline *int64 `json:"Offline,omitnil,omitempty" name:"Offline"`

	// 故障数
	Fault *int64 `json:"Fault,omitnil,omitempty" name:"Fault"`
}

type DeviceTypeSet struct {
	// 设备类型列表
	Set []*DeviceType `json:"Set,omitnil,omitempty" name:"Set"`
}

type ElementCoordinates struct {
	// 经度
	Longitude *float64 `json:"Longitude,omitnil,omitempty" name:"Longitude"`

	// 纬度
	Latitude *float64 `json:"Latitude,omitnil,omitempty" name:"Latitude"`

	// 高程
	Altitude *float64 `json:"Altitude,omitnil,omitempty" name:"Altitude"`
}

type ElementProfile struct {
	// 建筑id
	BuildingId *string `json:"BuildingId,omitnil,omitempty" name:"BuildingId"`

	// 构件id
	ElementId *string `json:"ElementId,omitnil,omitempty" name:"ElementId"`

	// 构件类型
	EntityType *string `json:"EntityType,omitnil,omitempty" name:"EntityType"`

	// 构件名称
	ElementName *string `json:"ElementName,omitnil,omitempty" name:"ElementName"`

	// 构件空间级别
	Level *uint64 `json:"Level,omitnil,omitempty" name:"Level"`

	// 底部标高（单位mm）
	BottomHeight *int64 `json:"BottomHeight,omitnil,omitempty" name:"BottomHeight"`

	// 排序
	Sort *uint64 `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 空间编码
	SpaceCode *string `json:"SpaceCode,omitnil,omitempty" name:"SpaceCode"`

	// 空间分类编码
	SpaceTypeCode *string `json:"SpaceTypeCode,omitnil,omitempty" name:"SpaceTypeCode"`

	// 空间分类名称
	SpaceTypeName *string `json:"SpaceTypeName,omitnil,omitempty" name:"SpaceTypeName"`

	// 父级构件id
	ParentElementId *string `json:"ParentElementId,omitnil,omitempty" name:"ParentElementId"`

	// 空间层级类型编码
	SpacePoiId *string `json:"SpacePoiId,omitnil,omitempty" name:"SpacePoiId"`

	// 构件描述
	ElementDesc *string `json:"ElementDesc,omitnil,omitempty" name:"ElementDesc"`

	// 删除标记
	IsDelete *uint64 `json:"IsDelete,omitnil,omitempty" name:"IsDelete"`
}

type ElementProfilePageRes struct {
	// 构件总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 构件列表
	List []*ElementProfile `json:"List,omitnil,omitempty" name:"List"`
}

type ElementProfileTreeNode struct {
	// 构件概要信息
	ElementProfile *ElementProfile `json:"ElementProfile,omitnil,omitempty" name:"ElementProfile"`

	// 子节点信息
	Children []*ElementProfileTreeNode `json:"Children,omitnil,omitempty" name:"Children"`
}

type ElementProfileTreeRes struct {
	// 建筑id
	BuildingId *string `json:"BuildingId,omitnil,omitempty" name:"BuildingId"`

	// 父级构件id
	ParentElementId *string `json:"ParentElementId,omitnil,omitempty" name:"ParentElementId"`

	// 构件树
	Root *ElementProfileTreeNode `json:"Root,omitnil,omitempty" name:"Root"`
}

type ElementProperty struct {
	// 属性名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 属性描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 属性内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type ElementPropertyRes struct {
	// 建筑id
	BuildingId *string `json:"BuildingId,omitnil,omitempty" name:"BuildingId"`

	// 构件id
	ElementId *string `json:"ElementId,omitnil,omitempty" name:"ElementId"`

	// 构件属性集合
	PropertySet []*ElementProperty `json:"PropertySet,omitnil,omitempty" name:"PropertySet"`

	// 构件地理坐标
	Coordinates *ElementCoordinates `json:"Coordinates,omitnil,omitempty" name:"Coordinates"`

	// 构件偏移量
	Translate *ElementTranslate `json:"Translate,omitnil,omitempty" name:"Translate"`

	// 构件名称
	ElementName *string `json:"ElementName,omitnil,omitempty" name:"ElementName"`

	// 构件类型代码
	EntityTypeCode *string `json:"EntityTypeCode,omitnil,omitempty" name:"EntityTypeCode"`

	// 构件类型名称
	EntityTypeName *string `json:"EntityTypeName,omitnil,omitempty" name:"EntityTypeName"`
}

type ElementTranslate struct {
	// X方向偏移量
	X *float64 `json:"X,omitnil,omitempty" name:"X"`

	// Y方向偏移量
	Y *float64 `json:"Y,omitnil,omitempty" name:"Y"`

	// Z方向偏移量
	Z *float64 `json:"Z,omitnil,omitempty" name:"Z"`
}

type EmptyRes struct {
	// 返回请求状态,成功ok，失败error
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`
}

type Event struct {
	// 事件id或动作Id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 事件名称或动作名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type EventDetail struct {
	// 事件id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 事件名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 事件触发类型
	TriggerType *string `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// 事件触发条件，返回为x-json后的字段
	TriggerCondition *string `json:"TriggerCondition,omitnil,omitempty" name:"TriggerCondition"`

	// 有效期
	ValidPeriod *string `json:"ValidPeriod,omitnil,omitempty" name:"ValidPeriod"`

	// 关联规则列表
	LinkRuleSet []*LinkRule `json:"LinkRuleSet,omitnil,omitempty" name:"LinkRuleSet"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 设备类型，当触发类型为deviceType时返回
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 设备的wid，当触发类型是device返回
	WID *string `json:"WID,omitnil,omitempty" name:"WID"`
}

type EventObj struct {
	// 事件id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 事件名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 事件触发类型名称
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 时间触发条件
	Condition *string `json:"Condition,omitnil,omitempty" name:"Condition"`
}

type FileDownloadURL struct {
	// 下载地址
	FileURL *string `json:"FileURL,omitnil,omitempty" name:"FileURL"`
}

type FileInfo struct {
	// 文件id
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 名称
	ReportName *string `json:"ReportName,omitnil,omitempty" name:"ReportName"`
}

type FileUploadURL struct {
	// 上传地址
	UploadURL *string `json:"UploadURL,omitnil,omitempty" name:"UploadURL"`

	// 文件Id
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 下载地址
	DownloadURL *string `json:"DownloadURL,omitnil,omitempty" name:"DownloadURL"`
}

type HandleRecordInfo struct {
	// 告警处理记录id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 操作类型
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`

	// 处理时间
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 文件列表
	FileSet []*FileInfo `json:"FileSet,omitnil,omitempty" name:"FileSet"`

	// 应用appid
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 扩展字段1，存非孪生中台用户id
	ExtendOne *string `json:"ExtendOne,omitnil,omitempty" name:"ExtendOne"`
}

type HandlerPersonInfo struct {
	// 用户id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 用户名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type LinkRule struct {
	// 关联联动规则id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 关联联动规则名字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type LinkRuleInfo struct {
	// 联动id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 联动名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 事件列表
	EventSet []*Event `json:"EventSet,omitnil,omitempty" name:"EventSet"`

	// 动作列表
	ActionSet []*Action `json:"ActionSet,omitnil,omitempty" name:"ActionSet"`

	// 状态：0开，-1关
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 起始时间
	BeginDate *string `json:"BeginDate,omitnil,omitempty" name:"BeginDate"`

	// 结束时间
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// 有效周期内容,有效字段为x-json后的字段
	ValidPeriod *string `json:"ValidPeriod,omitnil,omitempty" name:"ValidPeriod"`
}

type MessageProfile struct {
	// 应用类型
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`

	// 模型Id
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 设备类型
	PoiCode *string `json:"PoiCode,omitnil,omitempty" name:"PoiCode"`
}

type ModelInfo struct {
	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 模型名称
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 物模型id
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 关联产品信息
	RelatedProduct []*RelatedProduct `json:"RelatedProduct,omitnil,omitempty" name:"RelatedProduct"`

	// 设备类型名
	DeviceTypeName *string `json:"DeviceTypeName,omitnil,omitempty" name:"DeviceTypeName"`

	// 设备类型id
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 物模型类型，产品模型/标准模型
	ModelType *int64 `json:"ModelType,omitnil,omitempty" name:"ModelType"`

	// 模型参数内容,有效字段为"x-json:"后的字段
	ModelParams *string `json:"ModelParams,omitnil,omitempty" name:"ModelParams"`
}

type ModelSet struct {
	// 第几页
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页条数
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 总页数
	TotalPage *int64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// 总条数
	TotalRow *int64 `json:"TotalRow,omitnil,omitempty" name:"TotalRow"`

	// 模型基础信息
	Set []*ModelInfo `json:"Set,omitnil,omitempty" name:"Set"`
}

type ModifyDeviceFieldInfo struct {
	// 设备id
	WID *string `json:"WID,omitnil,omitempty" name:"WID"`

	// 自定义字段key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 自定义字段值
	Val *string `json:"Val,omitnil,omitempty" name:"Val"`
}

// Predefined struct for user
type ModifyDeviceFieldRequestParams struct {
	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 设备自定义字段修改信息集合
	Set []*ModifyDeviceFieldInfo `json:"Set,omitnil,omitempty" name:"Set"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

type ModifyDeviceFieldRequest struct {
	*tchttp.BaseRequest
	
	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 设备自定义字段修改信息集合
	Set []*ModifyDeviceFieldInfo `json:"Set,omitnil,omitempty" name:"Set"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

func (r *ModifyDeviceFieldRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDeviceFieldRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "Set")
	delete(f, "ApplicationToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDeviceFieldRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDeviceFieldResponseParams struct {
	// 返回请求结果
	Result *EmptyRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDeviceFieldResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDeviceFieldResponseParams `json:"Response"`
}

func (r *ModifyDeviceFieldResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDeviceFieldResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDeviceGroupInfo struct {
	// 设备id
	WID *string `json:"WID,omitnil,omitempty" name:"WID"`

	// 设备分组id
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

// Predefined struct for user
type ModifyDeviceGroupRequestParams struct {
	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 设备组修改信息集合	
	Set []*ModifyDeviceGroupInfo `json:"Set,omitnil,omitempty" name:"Set"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

type ModifyDeviceGroupRequest struct {
	*tchttp.BaseRequest
	
	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 设备组修改信息集合	
	Set []*ModifyDeviceGroupInfo `json:"Set,omitnil,omitempty" name:"Set"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

func (r *ModifyDeviceGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDeviceGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "Set")
	delete(f, "ApplicationToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDeviceGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDeviceGroupResponseParams struct {
	// 返回请求结果
	Result *EmptyRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDeviceGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDeviceGroupResponseParams `json:"Response"`
}

func (r *ModifyDeviceGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDeviceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDeviceNameRequestParams struct {
	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 设备修改信息集合
	Set []*DeviceModifyInfo `json:"Set,omitnil,omitempty" name:"Set"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

type ModifyDeviceNameRequest struct {
	*tchttp.BaseRequest
	
	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 设备修改信息集合
	Set []*DeviceModifyInfo `json:"Set,omitnil,omitempty" name:"Set"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

func (r *ModifyDeviceNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDeviceNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "Set")
	delete(f, "ApplicationToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDeviceNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDeviceNameResponseParams struct {
	// 返回请求结果
	Result *EmptyRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDeviceNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDeviceNameResponseParams `json:"Response"`
}

func (r *ModifyDeviceNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDeviceNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDeviceTagInfo struct {
	// 设备id
	WID *string `json:"WID,omitnil,omitempty" name:"WID"`

	// 设备标签名称集合
	NameSet []*string `json:"NameSet,omitnil,omitempty" name:"NameSet"`
}

// Predefined struct for user
type ModifyDeviceTagRequestParams struct {
	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 设备标签修改信息集合
	Set []*ModifyDeviceTagInfo `json:"Set,omitnil,omitempty" name:"Set"`

	// 应用token	
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

type ModifyDeviceTagRequest struct {
	*tchttp.BaseRequest
	
	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 设备标签修改信息集合
	Set []*ModifyDeviceTagInfo `json:"Set,omitnil,omitempty" name:"Set"`

	// 应用token	
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

func (r *ModifyDeviceTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDeviceTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "Set")
	delete(f, "ApplicationToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDeviceTagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDeviceTagResponseParams struct {
	// 返回请求结果
	Result *EmptyRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDeviceTagResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDeviceTagResponseParams `json:"Response"`
}

func (r *ModifyDeviceTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDeviceTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProcessRecordInfo struct {
	// 告警的id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 处理时间，毫秒
	ProcessTime *int64 `json:"ProcessTime,omitnil,omitempty" name:"ProcessTime"`

	// 处理类型，此处操作类型固定填add_record
	ProcessType *string `json:"ProcessType,omitnil,omitempty" name:"ProcessType"`

	// 注:此字段填写的是孪生中台的用户，非孪生中台用户不填该字段
	// [{\"id\":\"123\",\"name\":\"tes\"}]
	Processor *string `json:"Processor,omitnil,omitempty" name:"Processor"`

	// 处理描述信息
	ProcessDescription *string `json:"ProcessDescription,omitnil,omitempty" name:"ProcessDescription"`

	// 附加文件标识
	AttachedFileId *string `json:"AttachedFileId,omitnil,omitempty" name:"AttachedFileId"`
}

type ProductInfo struct {
	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 产品PID
	ProductId *int64 `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 产品名称
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 设备类型
	DeviceTypeName *string `json:"DeviceTypeName,omitnil,omitempty" name:"DeviceTypeName"`

	// 设备类型id
	DeviceTypeId *string `json:"DeviceTypeId,omitnil,omitempty" name:"DeviceTypeId"`

	// 产品属性，如：网关（1）、直连设备（2）、子设备（3）
	Attribute *int64 `json:"Attribute,omitnil,omitempty" name:"Attribute"`

	// 产品型号
	ProductType *string `json:"ProductType,omitnil,omitempty" name:"ProductType"`

	// 产品能力:信令数据、音视频。二进制数值中第0位表示信令数据、第1位表示音视频 。1（信令数据），3（具有信令数据以及音视频能力）。
	ProductAbility *int64 `json:"ProductAbility,omitnil,omitempty" name:"ProductAbility"`

	// 生产厂商
	Manufacturer *string `json:"Manufacturer,omitnil,omitempty" name:"Manufacturer"`

	// 维保厂商
	MaintenanceMfr *string `json:"MaintenanceMfr,omitnil,omitempty" name:"MaintenanceMfr"`

	// 物模型名称
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 物模型id
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 物模型类型，产品模型/标准模型
	ModelType *int64 `json:"ModelType,omitnil,omitempty" name:"ModelType"`
}

type ProductSet struct {
	// 第几页
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页条数
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 总页数
	TotalPage *int64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// 总条数
	TotalRow *int64 `json:"TotalRow,omitnil,omitempty" name:"TotalRow"`

	// 产品信息列表
	Product []*ProductInfo `json:"Product,omitnil,omitempty" name:"Product"`
}

type RawInfo struct {
	// 加密向量（如果视频网关选择流为非加密传输这个参数可忽略）
	SM4Vector *string `json:"SM4Vector,omitnil,omitempty" name:"SM4Vector"`

	// 专线ip (非专线接入可忽略)
	NATIP *string `json:"NATIP,omitnil,omitempty" name:"NATIP"`

	// 客户端握手鉴权参数
	StreamToken *string `json:"StreamToken,omitnil,omitempty" name:"StreamToken"`

	// 拉流端口
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 视频流加密key,目前为AES128加密KEY（如果视频网关选择流为非加密传输这个参数可忽略）
	StreamEnKey *string `json:"StreamEnKey,omitnil,omitempty" name:"StreamEnKey"`

	// 拉流公网地址（非公网接入时，这个地址是内网地址）
	IP *string `json:"IP,omitnil,omitempty" name:"IP"`

	// 拉流内网地址
	InnerIP *string `json:"InnerIP,omitnil,omitempty" name:"InnerIP"`
}

type RecordInfo struct {
	// 本录像片段开始时间（s）
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 本录像片段结束时间（s）
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 录像片段文件url
	VideoURL *string `json:"VideoURL,omitnil,omitempty" name:"VideoURL"`
}

type RelatedProduct struct {
	// 关联产品pid
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 关联产品名字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type ReportAppMessage struct {
	// 工作空间Id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 消息定义
	Profile *MessageProfile `json:"Profile,omitnil,omitempty" name:"Profile"`

	// 数据上报时间
	ReportTs *int64 `json:"ReportTs,omitnil,omitempty" name:"ReportTs"`

	// 属性定义 - KV，若为json格式，需在前加上x-json:，有效字段为x-json:后的字段
	Properties *string `json:"Properties,omitnil,omitempty" name:"Properties"`

	// 事件定义 - KKV，若为json格式，需在前加上x-json:，有效字段为x-json:后的字段
	EventSet *string `json:"EventSet,omitnil,omitempty" name:"EventSet"`

	// 服务定义 - KKV,若为json格式，需在前加上x-json:，有效字段为x-json:后的字段
	ServiceSet *string `json:"ServiceSet,omitnil,omitempty" name:"ServiceSet"`

	// 扩展字段2，如：算法上报id，若为json格式，需在前加上x-json:
	ExtendTwo *string `json:"ExtendTwo,omitnil,omitempty" name:"ExtendTwo"`

	// 透传信息，若为json格式，需在前加上x-json:，有效字段为x-json:后的字段
	Echo *string `json:"Echo,omitnil,omitempty" name:"Echo"`
}

// Predefined struct for user
type ReportAppMessageRequestParams struct {
	// 工作空间Id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 消息定义
	Profile *MessageProfile `json:"Profile,omitnil,omitempty" name:"Profile"`

	// 数据上报时间
	ReportTs *int64 `json:"ReportTs,omitnil,omitempty" name:"ReportTs"`

	// 属性定义 - KV的json格式,有效字段为x-json:后的字段
	Properties *string `json:"Properties,omitnil,omitempty" name:"Properties"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 事件定义 - KKV的json格式,有效字段为x-json:后的字段
	EventSet *string `json:"EventSet,omitnil,omitempty" name:"EventSet"`

	// 服务定义 - KKV的json格式,有效字段为x-json:后的字段
	ServiceSet *string `json:"ServiceSet,omitnil,omitempty" name:"ServiceSet"`

	// 扩展字段2，如：算法上报id，若为json格式，需传x-json:{}，有效字段为x-json:后的字段
	ExtendTwo *string `json:"ExtendTwo,omitnil,omitempty" name:"ExtendTwo"`

	// 透传信息，若为json格式，需传x-json:{}，有效字段为x-json:后的字段
	Echo *string `json:"Echo,omitnil,omitempty" name:"Echo"`
}

type ReportAppMessageRequest struct {
	*tchttp.BaseRequest
	
	// 工作空间Id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 消息定义
	Profile *MessageProfile `json:"Profile,omitnil,omitempty" name:"Profile"`

	// 数据上报时间
	ReportTs *int64 `json:"ReportTs,omitnil,omitempty" name:"ReportTs"`

	// 属性定义 - KV的json格式,有效字段为x-json:后的字段
	Properties *string `json:"Properties,omitnil,omitempty" name:"Properties"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 事件定义 - KKV的json格式,有效字段为x-json:后的字段
	EventSet *string `json:"EventSet,omitnil,omitempty" name:"EventSet"`

	// 服务定义 - KKV的json格式,有效字段为x-json:后的字段
	ServiceSet *string `json:"ServiceSet,omitnil,omitempty" name:"ServiceSet"`

	// 扩展字段2，如：算法上报id，若为json格式，需传x-json:{}，有效字段为x-json:后的字段
	ExtendTwo *string `json:"ExtendTwo,omitnil,omitempty" name:"ExtendTwo"`

	// 透传信息，若为json格式，需传x-json:{}，有效字段为x-json:后的字段
	Echo *string `json:"Echo,omitnil,omitempty" name:"Echo"`
}

func (r *ReportAppMessageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReportAppMessageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "Profile")
	delete(f, "ReportTs")
	delete(f, "Properties")
	delete(f, "ApplicationToken")
	delete(f, "EventSet")
	delete(f, "ServiceSet")
	delete(f, "ExtendTwo")
	delete(f, "Echo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReportAppMessageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReportAppMessageResponseParams struct {
	// 上报单条信息结果
	Result *EmptyRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReportAppMessageResponse struct {
	*tchttp.BaseResponse
	Response *ReportAppMessageResponseParams `json:"Response"`
}

func (r *ReportAppMessageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReportAppMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReportImg struct {
	// 类型
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 数据
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`
}

type ReportMsgRes struct {
	// 上报消息对应下标的16位标识Id, 即第几个消息
	ReportId *string `json:"ReportId,omitnil,omitempty" name:"ReportId"`

	// 上报消息结果，1表示成功推送，0表示推送失败
	ReportStatus *int64 `json:"ReportStatus,omitnil,omitempty" name:"ReportStatus"`
}

type RuleDetailRes struct {
	// 联动id
	RuleId *int64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 联动名称
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 联动说明
	RuleDesc *string `json:"RuleDesc,omitnil,omitempty" name:"RuleDesc"`

	// 1 全天有效，0：固定时间段有效
	ValidType *int64 `json:"ValidType,omitnil,omitempty" name:"ValidType"`

	// 有效期，json字符串（全天有效时为空）
	ValidPeriod *string `json:"ValidPeriod,omitnil,omitempty" name:"ValidPeriod"`

	// 起始时间
	BeginDate *string `json:"BeginDate,omitnil,omitempty" name:"BeginDate"`

	// 结束时间
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// 启用状态。1-启用，0-停用
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 触发规则，事件的组合
	EventRule *string `json:"EventRule,omitnil,omitempty" name:"EventRule"`

	// 事件对象集合
	EventInfoSet []*EventObj `json:"EventInfoSet,omitnil,omitempty" name:"EventInfoSet"`

	// 动作对象集合
	ActionInfoSet []*ActionObj `json:"ActionInfoSet,omitnil,omitempty" name:"ActionInfoSet"`
}

// Predefined struct for user
type SaveDeviceGroupRequestParams struct {
	// 分组名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分组描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 分组id, 携带则为修改, 不携带则为新增
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 分组父级id
	ParentId *int64 `json:"ParentId,omitnil,omitempty" name:"ParentId"`
}

type SaveDeviceGroupRequest struct {
	*tchttp.BaseRequest
	
	// 分组名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分组描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 分组id, 携带则为修改, 不携带则为新增
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 分组父级id
	ParentId *int64 `json:"ParentId,omitnil,omitempty" name:"ParentId"`
}

func (r *SaveDeviceGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SaveDeviceGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "WorkspaceId")
	delete(f, "ApplicationToken")
	delete(f, "Id")
	delete(f, "ParentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SaveDeviceGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SaveDeviceGroupRes struct {
	// 保存or修改设备分组回包信息
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`
}

// Predefined struct for user
type SaveDeviceGroupResponseParams struct {
	// 新增/修改的设备分组记录的id
	Result *SaveDeviceGroupRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SaveDeviceGroupResponse struct {
	*tchttp.BaseResponse
	Response *SaveDeviceGroupResponseParams `json:"Response"`
}

func (r *SaveDeviceGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SaveDeviceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SceneInfo struct {
	// 场景id
	SceneId *string `json:"SceneId,omitnil,omitempty" name:"SceneId"`

	// 场景名称
	SceneName *string `json:"SceneName,omitnil,omitempty" name:"SceneName"`

	// 场景版本
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`
}

type SceneListRes struct {
	// 场景列表
	SceneList []*SceneInfo `json:"SceneList,omitnil,omitempty" name:"SceneList"`
}

type SpaceDataListStatsRes struct {
	// 楼栋数量与建筑面积列表
	List []*SpaceDataStats `json:"List,omitnil,omitempty" name:"List"`
}

type SpaceDataStats struct {
	// 工作空间ID
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 工作空间名称
	WorkspaceName *string `json:"WorkspaceName,omitnil,omitempty" name:"WorkspaceName"`

	// 楼栋数量
	BuildingCount *uint64 `json:"BuildingCount,omitnil,omitempty" name:"BuildingCount"`

	// 建筑面积
	BuildingArea *float64 `json:"BuildingArea,omitnil,omitempty" name:"BuildingArea"`
}

type SpaceDataTotalStatsRes struct {
	// 总楼栋数量
	BuildingCount *uint64 `json:"BuildingCount,omitnil,omitempty" name:"BuildingCount"`

	// 总建筑面积
	BuildingArea *float64 `json:"BuildingArea,omitnil,omitempty" name:"BuildingArea"`
}

type SpaceDeviceIdListRes struct {
	// 设备id列表
	DeviceIds []*string `json:"DeviceIds,omitnil,omitempty" name:"DeviceIds"`
}

type SpaceDeviceRelation struct {
	// 设备id
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 构件id
	ElementId *string `json:"ElementId,omitnil,omitempty" name:"ElementId"`
}

type SpaceDeviceRelationRes struct {
	// 设备空间绑定关系列表
	SpaceDeviceRelationList []*SpaceDeviceRelation `json:"SpaceDeviceRelationList,omitnil,omitempty" name:"SpaceDeviceRelationList"`
}

type SpaceInfo struct {
	// 项目空间id
	WorkspaceId *uint64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 租户id
	TenantId *uint64 `json:"TenantId,omitnil,omitempty" name:"TenantId"`

	// 英文名
	EnglishName *string `json:"EnglishName,omitnil,omitempty" name:"EnglishName"`

	// 中文名
	ChineseName *string `json:"ChineseName,omitnil,omitempty" name:"ChineseName"`

	// 项目空间描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 项目空间状态:0 启用 1 停用 -1 已删除
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 是否是公共空间
	IsCommWorkspace *bool `json:"IsCommWorkspace,omitnil,omitempty" name:"IsCommWorkspace"`

	// 有效期开始时间
	ValidityStartTime *string `json:"ValidityStartTime,omitnil,omitempty" name:"ValidityStartTime"`

	// 有效期结束时间
	ValidityEndTime *string `json:"ValidityEndTime,omitnil,omitempty" name:"ValidityEndTime"`

	// 选中状态
	Selected *uint64 `json:"Selected,omitnil,omitempty" name:"Selected"`

	// 系统生成状态
	IsSystem *uint64 `json:"IsSystem,omitnil,omitempty" name:"IsSystem"`
}

type SpaceRelation struct {
	// 构件id
	ElementId *string `json:"ElementId,omitnil,omitempty" name:"ElementId"`

	// 构件名称
	ElementName *string `json:"ElementName,omitnil,omitempty" name:"ElementName"`

	// 空间层级
	Level *uint64 `json:"Level,omitnil,omitempty" name:"Level"`

	// 空间编码
	SpaceCode *string `json:"SpaceCode,omitnil,omitempty" name:"SpaceCode"`

	// 父级空间编码
	ParentSpaceCode *string `json:"ParentSpaceCode,omitnil,omitempty" name:"ParentSpaceCode"`

	// 子构件信息
	Children []*SpaceRelation `json:"Children,omitnil,omitempty" name:"Children"`
}

type SpaceRelationRes struct {
	// 空间层级关系
	SpaceRelation *SpaceRelation `json:"SpaceRelation,omitnil,omitempty" name:"SpaceRelation"`
}

type SpaceType struct {
	// 空间分类编码
	SpaceTypeCode *string `json:"SpaceTypeCode,omitnil,omitempty" name:"SpaceTypeCode"`

	// 空间分类名称
	SpaceTypeName *string `json:"SpaceTypeName,omitnil,omitempty" name:"SpaceTypeName"`
}

type SpaceTypeListRes struct {
	// 空间分类列表
	SpaceTypeList []*SpaceType `json:"SpaceTypeList,omitnil,omitempty" name:"SpaceTypeList"`
}

type SsoDepartment struct {
	// 部门ID
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`

	// 部门名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 父级部门ID
	ParentDepartmentId *string `json:"ParentDepartmentId,omitnil,omitempty" name:"ParentDepartmentId"`
}

type SsoDepartmentsResult struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 部门列表
	Departments []*SsoDepartment `json:"Departments,omitnil,omitempty" name:"Departments"`
}

type SsoTeamUser struct {
	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户名称
	RealName *string `json:"RealName,omitnil,omitempty" name:"RealName"`

	// 用户类型，1-超级管理员；2-1号管理员；3-普通管理员；99-普通用户
	UserType *string `json:"UserType,omitnil,omitempty" name:"UserType"`

	// 所属租户ID
	TenantId *string `json:"TenantId,omitnil,omitempty" name:"TenantId"`

	// 邮箱
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 电话
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 用户状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建时间
	CreateAt *int64 `json:"CreateAt,omitnil,omitempty" name:"CreateAt"`

	// 部门ID
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`

	// 部门名称
	DepartmentName *string `json:"DepartmentName,omitnil,omitempty" name:"DepartmentName"`

	// 是否关联权限
	LinkFilter *int64 `json:"LinkFilter,omitnil,omitempty" name:"LinkFilter"`
}

type SsoTeamUserResult struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 部门用户列表
	Users []*SsoTeamUser `json:"Users,omitnil,omitempty" name:"Users"`
}

type SsoUser struct {
	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户昵称
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 用户名称
	RealName *string `json:"RealName,omitnil,omitempty" name:"RealName"`

	// 用户类型，1-超级管理员；2-1号管理员；3-普通管理员；99-普通用户
	UserType *string `json:"UserType,omitnil,omitempty" name:"UserType"`

	// 所属租户ID
	TenantId *string `json:"TenantId,omitnil,omitempty" name:"TenantId"`

	// 所属组ID
	UserGroup *string `json:"UserGroup,omitnil,omitempty" name:"UserGroup"`

	// 邮箱
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 电话
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 用户状态，0待审核，1正常启用，2禁用
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建时间
	CreateAt *int64 `json:"CreateAt,omitnil,omitempty" name:"CreateAt"`

	// 更新时间
	UpdateAt *int64 `json:"UpdateAt,omitnil,omitempty" name:"UpdateAt"`

	// 是否属于团队，0不可用，1属于，2不属
	BelongTeam *int64 `json:"BelongTeam,omitnil,omitempty" name:"BelongTeam"`

	// 部门ID
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`

	// 部门名称
	DepartmentName *string `json:"DepartmentName,omitnil,omitempty" name:"DepartmentName"`

	// 子账户ID
	DepartmentUserId *int64 `json:"DepartmentUserId,omitnil,omitempty" name:"DepartmentUserId"`

	// 密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

type SsoUserResult struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 租户人员数据
	Users []*SsoUser `json:"Users,omitnil,omitempty" name:"Users"`
}

type StatDeviceType struct {
	// 汇总数。在线（正常+故障） + 离线
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 正常数
	Normal *int64 `json:"Normal,omitnil,omitempty" name:"Normal"`

	// 离线数
	Offline *int64 `json:"Offline,omitnil,omitempty" name:"Offline"`

	// 故障数
	Fault *int64 `json:"Fault,omitnil,omitempty" name:"Fault"`

	// 设备名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 设备类型
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`
}

type StatLevel struct {
	// 汇总数。在线（正常+故障） + 离线
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 正常数
	NormalSum *int64 `json:"NormalSum,omitnil,omitempty" name:"NormalSum"`

	// 离线数
	OfflineSum *int64 `json:"OfflineSum,omitnil,omitempty" name:"OfflineSum"`

	// 故障数
	FaultSum *int64 `json:"FaultSum,omitnil,omitempty" name:"FaultSum"`

	// 空间id
	SpaceCode *string `json:"SpaceCode,omitnil,omitempty" name:"SpaceCode"`

	// 设备类型统计列表
	StatDeviceTypeSet []*StatDeviceType `json:"StatDeviceTypeSet,omitnil,omitempty" name:"StatDeviceTypeSet"`
}

// Predefined struct for user
type StopVideoStreamingRequestParams struct {
	// 流的唯一标识，播放链接尾缀
	Stream *string `json:"Stream,omitnil,omitempty" name:"Stream"`

	// 设备的唯一标识
	WID *string `json:"WID,omitnil,omitempty" name:"WID"`

	// 工作空间Id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

type StopVideoStreamingRequest struct {
	*tchttp.BaseRequest
	
	// 流的唯一标识，播放链接尾缀
	Stream *string `json:"Stream,omitnil,omitempty" name:"Stream"`

	// 设备的唯一标识
	WID *string `json:"WID,omitnil,omitempty" name:"WID"`

	// 工作空间Id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`
}

func (r *StopVideoStreamingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopVideoStreamingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Stream")
	delete(f, "WID")
	delete(f, "WorkspaceId")
	delete(f, "ApplicationToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopVideoStreamingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopVideoStreamingResponseParams struct {
	// 停流接口返回结果
	Result *EmptyRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopVideoStreamingResponse struct {
	*tchttp.BaseResponse
	Response *StopVideoStreamingResponseParams `json:"Response"`
}

func (r *StopVideoStreamingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopVideoStreamingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateWorkspaceParkAttributesRequestParams struct {
	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 园区简称
	ParkName *string `json:"ParkName,omitnil,omitempty" name:"ParkName"`

	// 园区编号
	ParkNum *string `json:"ParkNum,omitnil,omitempty" name:"ParkNum"`
}

type UpdateWorkspaceParkAttributesRequest struct {
	*tchttp.BaseRequest
	
	// 工作空间id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 应用token
	ApplicationToken *string `json:"ApplicationToken,omitnil,omitempty" name:"ApplicationToken"`

	// 园区简称
	ParkName *string `json:"ParkName,omitnil,omitempty" name:"ParkName"`

	// 园区编号
	ParkNum *string `json:"ParkNum,omitnil,omitempty" name:"ParkNum"`
}

func (r *UpdateWorkspaceParkAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateWorkspaceParkAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "ApplicationToken")
	delete(f, "ParkName")
	delete(f, "ParkNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateWorkspaceParkAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateWorkspaceParkAttributesResponseParams struct {
	// 修改工作空间园区属性结果
	Result *EmptyRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateWorkspaceParkAttributesResponse struct {
	*tchttp.BaseResponse
	Response *UpdateWorkspaceParkAttributesResponseParams `json:"Response"`
}

func (r *UpdateWorkspaceParkAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateWorkspaceParkAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VideoCloudRecordRes struct {
	// 录像信息总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 录像信息列表
	RecordSet []*RecordInfo `json:"RecordSet,omitnil,omitempty" name:"RecordSet"`
}

type VideoRecordStreamRes struct {
	// FLV协议格式视频流
	FLV *string `json:"FLV,omitnil,omitempty" name:"FLV"`

	// RTMP协议格式视频流
	RTMP *string `json:"RTMP,omitnil,omitempty" name:"RTMP"`

	// HLS协议格式视频流
	HLS *string `json:"HLS,omitnil,omitempty" name:"HLS"`

	// WebRtc协议格式视频流
	WebRTC *string `json:"WebRTC,omitnil,omitempty" name:"WebRTC"`

	// RAW协议格式视频流
	RAW *RawInfo `json:"RAW,omitnil,omitempty" name:"RAW"`

	// 视频流的唯一标识
	Stream *string `json:"Stream,omitnil,omitempty" name:"Stream"`
}

type WorkspaceInfo struct {
	// 工作空间Id
	WorkspaceId *int64 `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 工作空间中文名字
	ChineseName *string `json:"ChineseName,omitnil,omitempty" name:"ChineseName"`

	// 工作空间描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 工作空间是否删除状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 该工作空间绑定的区/县的行政区名字
	ParkName *string `json:"ParkName,omitnil,omitempty" name:"ParkName"`

	// 该工作空间绑定的区/县的行政区编码
	ParkNum *string `json:"ParkNum,omitnil,omitempty" name:"ParkNum"`

	// 获取该工作空间绑定的区/县的上级行政区划信息
	AdministrativeDetailSet []*AdministrativeDetail `json:"AdministrativeDetailSet,omitnil,omitempty" name:"AdministrativeDetailSet"`
}

type WorkspaceInfoList struct {
	// 项目空间列表
	List []*SpaceInfo `json:"List,omitnil,omitempty" name:"List"`
}