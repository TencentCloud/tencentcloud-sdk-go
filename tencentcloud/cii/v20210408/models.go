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

package v20210408

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type AddSubStructureTasksRequestParams struct {
	// 主任务id
	MainTaskId *string `json:"MainTaskId,omitempty" name:"MainTaskId"`

	// 子任务信息数组
	TaskInfos []*CreateStructureTaskInfo `json:"TaskInfos,omitempty" name:"TaskInfos"`
}

type AddSubStructureTasksRequest struct {
	*tchttp.BaseRequest
	
	// 主任务id
	MainTaskId *string `json:"MainTaskId,omitempty" name:"MainTaskId"`

	// 子任务信息数组
	TaskInfos []*CreateStructureTaskInfo `json:"TaskInfos,omitempty" name:"TaskInfos"`
}

func (r *AddSubStructureTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddSubStructureTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MainTaskId")
	delete(f, "TaskInfos")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddSubStructureTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddSubStructureTasksResponseParams struct {
	// 增量子任务id数组
	SubTaskIds []*string `json:"SubTaskIds,omitempty" name:"SubTaskIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddSubStructureTasksResponse struct {
	*tchttp.BaseResponse
	Response *AddSubStructureTasksResponseParams `json:"Response"`
}

func (r *AddSubStructureTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddSubStructureTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClassifiedReports struct {
	// 报告类型
	ReportType *string `json:"ReportType,omitempty" name:"ReportType"`

	// 文件列表
	FileList []*string `json:"FileList,omitempty" name:"FileList"`
}

type ClassifyInfo struct {
	// 一级分类
	FirstClass *string `json:"FirstClass,omitempty" name:"FirstClass"`

	// 二级分类
	SecondClass *string `json:"SecondClass,omitempty" name:"SecondClass"`

	// 三级分类
	ThirdClass *string `json:"ThirdClass,omitempty" name:"ThirdClass"`

	// 一级分类序号
	FirstClassId *uint64 `json:"FirstClassId,omitempty" name:"FirstClassId"`

	// 二级分类序号
	SecondClassId *uint64 `json:"SecondClassId,omitempty" name:"SecondClassId"`

	// 三级分类序号
	ThirdClassId *uint64 `json:"ThirdClassId,omitempty" name:"ThirdClassId"`
}

type CompareMetricsData struct {
	// 短文准确率
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShortStructAccuracy *string `json:"ShortStructAccuracy,omitempty" name:"ShortStructAccuracy"`

	// 短文召回率
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShortStructRecall *string `json:"ShortStructRecall,omitempty" name:"ShortStructRecall"`

	// 长文结构化准确率
	// 注意：此字段可能返回 null，表示取不到有效值。
	LongStructAccuracy *string `json:"LongStructAccuracy,omitempty" name:"LongStructAccuracy"`

	// 长文结构化召回率
	// 注意：此字段可能返回 null，表示取不到有效值。
	LongStructRecall *string `json:"LongStructRecall,omitempty" name:"LongStructRecall"`

	// 长文提取准确率
	// 注意：此字段可能返回 null，表示取不到有效值。
	LongContentAccuracy *string `json:"LongContentAccuracy,omitempty" name:"LongContentAccuracy"`

	// 长文提取召回率
	// 注意：此字段可能返回 null，表示取不到有效值。
	LongContentRecall *string `json:"LongContentRecall,omitempty" name:"LongContentRecall"`
}

type CreateAutoClassifyStructureTaskInfo struct {
	// 报告文件上传的地址列表，需按顺序排列。如果使用ImageList参数，置为空数组即可
	FileList []*string `json:"FileList,omitempty" name:"FileList"`

	// 客户号
	CustomerId *string `json:"CustomerId,omitempty" name:"CustomerId"`

	// 客户姓名
	CustomerName *string `json:"CustomerName,omitempty" name:"CustomerName"`

	// 报告上传的图片内容数组，图片内容采用base64编码，需按顺序排列
	ImageList []*string `json:"ImageList,omitempty" name:"ImageList"`
}

// Predefined struct for user
type CreateAutoClassifyStructureTaskRequestParams struct {
	// 服务类型
	// Structured 仅结构化
	// Underwrite 结构化+核保
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 创建任务时可以上传多个报告，后台生成多个识别子任务，子任务的详细信息
	TaskInfos []*CreateAutoClassifyStructureTaskInfo `json:"TaskInfos,omitempty" name:"TaskInfos"`

	// 保单号
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 核保触发方式
	// Auto 自动
	// Manual 手动
	TriggerType *string `json:"TriggerType,omitempty" name:"TriggerType"`

	// 险种，如果是体检报告类型，此参数是必填，类型说明如下：
	// CriticalDiseaseInsurance:重疾险
	// LifeInsurance：寿险
	// AccidentInsurance：意外险
	InsuranceTypes []*string `json:"InsuranceTypes,omitempty" name:"InsuranceTypes"`

	// 回调地址，接收Post请求传送结果
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`
}

type CreateAutoClassifyStructureTaskRequest struct {
	*tchttp.BaseRequest
	
	// 服务类型
	// Structured 仅结构化
	// Underwrite 结构化+核保
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 创建任务时可以上传多个报告，后台生成多个识别子任务，子任务的详细信息
	TaskInfos []*CreateAutoClassifyStructureTaskInfo `json:"TaskInfos,omitempty" name:"TaskInfos"`

	// 保单号
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 核保触发方式
	// Auto 自动
	// Manual 手动
	TriggerType *string `json:"TriggerType,omitempty" name:"TriggerType"`

	// 险种，如果是体检报告类型，此参数是必填，类型说明如下：
	// CriticalDiseaseInsurance:重疾险
	// LifeInsurance：寿险
	// AccidentInsurance：意外险
	InsuranceTypes []*string `json:"InsuranceTypes,omitempty" name:"InsuranceTypes"`

	// 回调地址，接收Post请求传送结果
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`
}

func (r *CreateAutoClassifyStructureTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAutoClassifyStructureTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceType")
	delete(f, "TaskInfos")
	delete(f, "PolicyId")
	delete(f, "TriggerType")
	delete(f, "InsuranceTypes")
	delete(f, "CallbackUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAutoClassifyStructureTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAutoClassifyStructureTaskResponseParams struct {
	// 创建的主任务号，用于查询结果
	MainTaskId *string `json:"MainTaskId,omitempty" name:"MainTaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAutoClassifyStructureTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateAutoClassifyStructureTaskResponseParams `json:"Response"`
}

func (r *CreateAutoClassifyStructureTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAutoClassifyStructureTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateStructureTaskInfo struct {
	// 任务类型:HealthReport(体检报告); BUltraReport(B超报告);MedCheckReport(检查报告);LaboratoryReport(检验报告); PathologyReport(病理报告);AdmissionReport(入院记录);DischargeReport(出院记录); DischargeSummary(出院小结);DiagnosisReport(诊断证明); MedicalRecordFront(病案首页);OperationReport(手术记录);OutpatientMedicalRecord(门诊病历)
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 报告文件上传的地址列表，需按顺序排列。如果使用ImageList参数，置为空数组即可
	FileList []*string `json:"FileList,omitempty" name:"FileList"`

	// 客户号
	CustomerId *string `json:"CustomerId,omitempty" name:"CustomerId"`

	// 客户姓名
	CustomerName *string `json:"CustomerName,omitempty" name:"CustomerName"`

	// 报告上传的图片内容数组，图片内容采用base64编码，需按顺序排列
	ImageList []*string `json:"ImageList,omitempty" name:"ImageList"`

	// 报告年份
	Year *string `json:"Year,omitempty" name:"Year"`
}

// Predefined struct for user
type CreateStructureTaskRequestParams struct {
	// 服务类型
	// Structured 仅结构化
	// Underwrite 结构化+核保
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 创建任务时可以上传多个报告，后台生成多个识别子任务，子任务的详细信息
	TaskInfos []*CreateStructureTaskInfo `json:"TaskInfos,omitempty" name:"TaskInfos"`

	// 保单号
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 核保触发方式
	// Auto 自动
	// Manual 手动
	TriggerType *string `json:"TriggerType,omitempty" name:"TriggerType"`

	// 险种，如果是体检报告类型，此参数是必填，类型说明如下：
	// CriticalDiseaseInsurance:重疾险
	// LifeInsurance：寿险
	// AccidentInsurance：意外险
	InsuranceTypes []*string `json:"InsuranceTypes,omitempty" name:"InsuranceTypes"`

	// 回调地址，接收Post请求传送结果
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`
}

type CreateStructureTaskRequest struct {
	*tchttp.BaseRequest
	
	// 服务类型
	// Structured 仅结构化
	// Underwrite 结构化+核保
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 创建任务时可以上传多个报告，后台生成多个识别子任务，子任务的详细信息
	TaskInfos []*CreateStructureTaskInfo `json:"TaskInfos,omitempty" name:"TaskInfos"`

	// 保单号
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 核保触发方式
	// Auto 自动
	// Manual 手动
	TriggerType *string `json:"TriggerType,omitempty" name:"TriggerType"`

	// 险种，如果是体检报告类型，此参数是必填，类型说明如下：
	// CriticalDiseaseInsurance:重疾险
	// LifeInsurance：寿险
	// AccidentInsurance：意外险
	InsuranceTypes []*string `json:"InsuranceTypes,omitempty" name:"InsuranceTypes"`

	// 回调地址，接收Post请求传送结果
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`
}

func (r *CreateStructureTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStructureTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceType")
	delete(f, "TaskInfos")
	delete(f, "PolicyId")
	delete(f, "TriggerType")
	delete(f, "InsuranceTypes")
	delete(f, "CallbackUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStructureTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStructureTaskResponseParams struct {
	// 创建的主任务号，用于查询结果
	MainTaskId *string `json:"MainTaskId,omitempty" name:"MainTaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateStructureTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateStructureTaskResponseParams `json:"Response"`
}

func (r *CreateStructureTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStructureTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUnderwriteTaskByIdRequestParams struct {
	// 主任务ID数组，
	MainTaskIds []*string `json:"MainTaskIds,omitempty" name:"MainTaskIds"`

	// 回调地址，可不传（提供轮询机制）。
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`
}

type CreateUnderwriteTaskByIdRequest struct {
	*tchttp.BaseRequest
	
	// 主任务ID数组，
	MainTaskIds []*string `json:"MainTaskIds,omitempty" name:"MainTaskIds"`

	// 回调地址，可不传（提供轮询机制）。
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`
}

func (r *CreateUnderwriteTaskByIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUnderwriteTaskByIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MainTaskIds")
	delete(f, "CallbackUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUnderwriteTaskByIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUnderwriteTaskByIdResponseParams struct {
	// 核保任务ID数据
	UnderwriteTaskIds []*string `json:"UnderwriteTaskIds,omitempty" name:"UnderwriteTaskIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateUnderwriteTaskByIdResponse struct {
	*tchttp.BaseResponse
	Response *CreateUnderwriteTaskByIdResponseParams `json:"Response"`
}

func (r *CreateUnderwriteTaskByIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUnderwriteTaskByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMachineUnderwriteRequestParams struct {
	// 核保任务ID
	UnderwriteTaskId *string `json:"UnderwriteTaskId,omitempty" name:"UnderwriteTaskId"`
}

type DescribeMachineUnderwriteRequest struct {
	*tchttp.BaseRequest
	
	// 核保任务ID
	UnderwriteTaskId *string `json:"UnderwriteTaskId,omitempty" name:"UnderwriteTaskId"`
}

func (r *DescribeMachineUnderwriteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineUnderwriteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UnderwriteTaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMachineUnderwriteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMachineUnderwriteResponseParams struct {
	// 腾讯云主账号ID
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 操作人子账户ID
	SubAccountUin *string `json:"SubAccountUin,omitempty" name:"SubAccountUin"`

	// 保单ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 主任务ID
	MainTaskId *string `json:"MainTaskId,omitempty" name:"MainTaskId"`

	// 核保任务ID
	UnderwriteTaskId *string `json:"UnderwriteTaskId,omitempty" name:"UnderwriteTaskId"`

	// 结果状态：
	// 0：返回成功
	// 1：结果未生成
	// 2：结果生成失败
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 机器核保结果
	UnderwriteResults []*MachineUnderwriteOutput `json:"UnderwriteResults,omitempty" name:"UnderwriteResults"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeMachineUnderwriteResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMachineUnderwriteResponseParams `json:"Response"`
}

func (r *DescribeMachineUnderwriteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineUnderwriteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQualityScoreRequestParams struct {
	// 文件二进制数据
	File *string `json:"File,omitempty" name:"File"`
}

type DescribeQualityScoreRequest struct {
	*tchttp.BaseRequest
	
	// 文件二进制数据
	File *string `json:"File,omitempty" name:"File"`
}

func (r *DescribeQualityScoreRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQualityScoreRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "File")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeQualityScoreRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQualityScoreResponseParams struct {
	// 质量分
	QualityScore *float64 `json:"QualityScore,omitempty" name:"QualityScore"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeQualityScoreResponse struct {
	*tchttp.BaseResponse
	Response *DescribeQualityScoreResponseParams `json:"Response"`
}

func (r *DescribeQualityScoreResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQualityScoreResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReportClassifyRequestParams struct {
	// 服务类型
	// Structured 仅结构化
	// Underwrite 结构化+核保
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 文件地址数组
	FileList []*string `json:"FileList,omitempty" name:"FileList"`
}

type DescribeReportClassifyRequest struct {
	*tchttp.BaseRequest
	
	// 服务类型
	// Structured 仅结构化
	// Underwrite 结构化+核保
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 文件地址数组
	FileList []*string `json:"FileList,omitempty" name:"FileList"`
}

func (r *DescribeReportClassifyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReportClassifyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceType")
	delete(f, "FileList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReportClassifyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReportClassifyResponseParams struct {
	// 报告分类结果
	Reports []*ClassifiedReports `json:"Reports,omitempty" name:"Reports"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeReportClassifyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReportClassifyResponseParams `json:"Response"`
}

func (r *DescribeReportClassifyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReportClassifyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStructCompareDataRequestParams struct {
	// 主任务号
	MainTaskId *string `json:"MainTaskId,omitempty" name:"MainTaskId"`

	// 子任务号
	SubTaskId *string `json:"SubTaskId,omitempty" name:"SubTaskId"`
}

type DescribeStructCompareDataRequest struct {
	*tchttp.BaseRequest
	
	// 主任务号
	MainTaskId *string `json:"MainTaskId,omitempty" name:"MainTaskId"`

	// 子任务号
	SubTaskId *string `json:"SubTaskId,omitempty" name:"SubTaskId"`
}

func (r *DescribeStructCompareDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStructCompareDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MainTaskId")
	delete(f, "SubTaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStructCompareDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStructCompareDataResponseParams struct {
	// 保单号
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 主任务号
	MainTaskId *string `json:"MainTaskId,omitempty" name:"MainTaskId"`

	// 客户号
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomerId *string `json:"CustomerId,omitempty" name:"CustomerId"`

	// 客户姓名
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomerName *string `json:"CustomerName,omitempty" name:"CustomerName"`

	// 复核时间
	ReviewTime *string `json:"ReviewTime,omitempty" name:"ReviewTime"`

	// 算法识别结果
	MachineResult *string `json:"MachineResult,omitempty" name:"MachineResult"`

	// 人工复核结果
	ManualResult *string `json:"ManualResult,omitempty" name:"ManualResult"`

	// 结构化对比指标数据
	Metrics *CompareMetricsData `json:"Metrics,omitempty" name:"Metrics"`

	// 新增项
	NewItems *string `json:"NewItems,omitempty" name:"NewItems"`

	// 修改项
	ModifyItems *string `json:"ModifyItems,omitempty" name:"ModifyItems"`

	// 子任务号
	SubTaskId *string `json:"SubTaskId,omitempty" name:"SubTaskId"`

	// 所有的子任务
	AllTasks []*ReviewDataTaskInfo `json:"AllTasks,omitempty" name:"AllTasks"`

	// 任务类型
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeStructCompareDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStructCompareDataResponseParams `json:"Response"`
}

func (r *DescribeStructCompareDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStructCompareDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStructureDifferenceRequestParams struct {
	// 主任务号
	MainTaskId *string `json:"MainTaskId,omitempty" name:"MainTaskId"`

	// 子任务号
	SubTaskId *string `json:"SubTaskId,omitempty" name:"SubTaskId"`
}

type DescribeStructureDifferenceRequest struct {
	*tchttp.BaseRequest
	
	// 主任务号
	MainTaskId *string `json:"MainTaskId,omitempty" name:"MainTaskId"`

	// 子任务号
	SubTaskId *string `json:"SubTaskId,omitempty" name:"SubTaskId"`
}

func (r *DescribeStructureDifferenceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStructureDifferenceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MainTaskId")
	delete(f, "SubTaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStructureDifferenceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStructureDifferenceResponseParams struct {
	// 主任务号
	MainTaskId *string `json:"MainTaskId,omitempty" name:"MainTaskId"`

	// 结果状态：
	// 0：返回成功
	// 1：结果未生成
	// 2：结果生成失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 差异的结果数组
	Results []*PerStructDifference `json:"Results,omitempty" name:"Results"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeStructureDifferenceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStructureDifferenceResponseParams `json:"Response"`
}

func (r *DescribeStructureDifferenceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStructureDifferenceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStructureResultRequestParams struct {
	// 创建任务时返回的主任务ID
	MainTaskId *string `json:"MainTaskId,omitempty" name:"MainTaskId"`
}

type DescribeStructureResultRequest struct {
	*tchttp.BaseRequest
	
	// 创建任务时返回的主任务ID
	MainTaskId *string `json:"MainTaskId,omitempty" name:"MainTaskId"`
}

func (r *DescribeStructureResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStructureResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MainTaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStructureResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStructureResultResponseParams struct {
	// 结果状态：
	// 0：返回成功
	// 1：结果未生成
	// 2：结果生成失败
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 结构化结果
	Results []*StructureResultObject `json:"Results,omitempty" name:"Results"`

	// 主任务ID
	MainTaskId *string `json:"MainTaskId,omitempty" name:"MainTaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeStructureResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStructureResultResponseParams `json:"Response"`
}

func (r *DescribeStructureResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStructureResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStructureTaskResultRequestParams struct {
	// 结构化任务ID
	MainTaskId *string `json:"MainTaskId,omitempty" name:"MainTaskId"`
}

type DescribeStructureTaskResultRequest struct {
	*tchttp.BaseRequest
	
	// 结构化任务ID
	MainTaskId *string `json:"MainTaskId,omitempty" name:"MainTaskId"`
}

func (r *DescribeStructureTaskResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStructureTaskResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MainTaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStructureTaskResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStructureTaskResultResponseParams struct {
	// 结果状态：
	// 0：返回成功
	// 1：结果未生成
	// 2：结果生成失败
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 结构化识别结果数组，每个数组元素对应一个图片的结构化结果，顺序和输入参数的ImageList或FileList对应。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Results []*ResultObject `json:"Results,omitempty" name:"Results"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeStructureTaskResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStructureTaskResultResponseParams `json:"Response"`
}

func (r *DescribeStructureTaskResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStructureTaskResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUnderwriteTaskRequestParams struct {
	// 任务ID
	UnderwriteTaskId *string `json:"UnderwriteTaskId,omitempty" name:"UnderwriteTaskId"`
}

type DescribeUnderwriteTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	UnderwriteTaskId *string `json:"UnderwriteTaskId,omitempty" name:"UnderwriteTaskId"`
}

func (r *DescribeUnderwriteTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUnderwriteTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UnderwriteTaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUnderwriteTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUnderwriteTaskResponseParams struct {
	// 腾讯云主账号ID
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 操作人子账户ID
	SubAccountUin *string `json:"SubAccountUin,omitempty" name:"SubAccountUin"`

	// 保单ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 主任务ID
	MainTaskId *string `json:"MainTaskId,omitempty" name:"MainTaskId"`

	// 核保任务ID
	UnderwriteTaskId *string `json:"UnderwriteTaskId,omitempty" name:"UnderwriteTaskId"`

	// 结果状态：
	// 0：返回成功
	// 1：结果未生成
	// 2：结果生成失败
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 核保结果
	UnderwriteResults []*UnderwriteOutput `json:"UnderwriteResults,omitempty" name:"UnderwriteResults"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUnderwriteTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUnderwriteTaskResponseParams `json:"Response"`
}

func (r *DescribeUnderwriteTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUnderwriteTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InsuranceResult struct {
	// 险种:CriticalDiseaseInsurance(重疾险);LifeInsurance(寿险);AccidentInsurance(意外险);MedicalInsurance(医疗险)
	InsuranceType *string `json:"InsuranceType,omitempty" name:"InsuranceType"`

	// 对应险种的机器核保结果
	Result []*MachinePredict `json:"Result,omitempty" name:"Result"`
}

type Location struct {
	// 位置信息
	Points []*Point `json:"Points,omitempty" name:"Points"`
}

type MachinePredict struct {
	// 核保引擎名称
	Title *string `json:"Title,omitempty" name:"Title"`

	// 核保结论：加费、承保、拒保、延期、除外、加费+除外
	Conclusion *string `json:"Conclusion,omitempty" name:"Conclusion"`

	// AI决策树解释
	Explanation []*UnderwriteItem `json:"Explanation,omitempty" name:"Explanation"`

	// 疾病指标
	Disease []*UnderwriteItem `json:"Disease,omitempty" name:"Disease"`

	// 检查异常
	Laboratory []*UnderwriteItem `json:"Laboratory,omitempty" name:"Laboratory"`
}

type MachineUnderwriteOutput struct {
	// 客户号
	CustomerId *string `json:"CustomerId,omitempty" name:"CustomerId"`

	// 客户姓名
	CustomerName *string `json:"CustomerName,omitempty" name:"CustomerName"`

	// 各个险种的结果
	Results []*InsuranceResult `json:"Results,omitempty" name:"Results"`
}

type OcrRecognise struct {
	// 原文字段
	OriginalField *string `json:"OriginalField,omitempty" name:"OriginalField"`

	// 识别结果
	Value *string `json:"Value,omitempty" name:"Value"`

	// 置信度
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 位置信息
	Location *Location `json:"Location,omitempty" name:"Location"`

	// 字段名
	Field *string `json:"Field,omitempty" name:"Field"`
}

type PerStructDifference struct {
	// 子任务ID
	SubTaskId *string `json:"SubTaskId,omitempty" name:"SubTaskId"`

	// 任务类型:HealthReport(体检报告); BUltraReport(B超报告);MedCheckReport(检查报告);LaboratoryReport(检验报告); PathologyReport(病理报告);AdmissionReport(入院记录);DischargeReport(出院记录); DischargeSummary(出院小结);DiagnosisReport(诊断证明); MedicalRecordFront(病案首页);OperationReport(手术记录);OutpatientMedicalRecord(门诊病历)
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 修改的项
	ModifyItems []*StructureModifyItem `json:"ModifyItems,omitempty" name:"ModifyItems"`

	// 新增的项
	NewItems []*StructureOneItem `json:"NewItems,omitempty" name:"NewItems"`

	// 删除的项
	RemoveItems []*StructureOneItem `json:"RemoveItems,omitempty" name:"RemoveItems"`
}

type Point struct {
	// x坐标
	XCoordinate *int64 `json:"XCoordinate,omitempty" name:"XCoordinate"`

	// y坐标
	YCoordinate *int64 `json:"YCoordinate,omitempty" name:"YCoordinate"`

	// 页码
	Page *int64 `json:"Page,omitempty" name:"Page"`
}

type ResultObject struct {
	// 图片质量分
	Quality *float64 `json:"Quality,omitempty" name:"Quality"`

	// 由结构化算法结构化json转换的字符串，具体协议参见算法结构化结果协议
	StructureResult *string `json:"StructureResult,omitempty" name:"StructureResult"`

	// 报告分类信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportType []*ClassifyInfo `json:"ReportType,omitempty" name:"ReportType"`
}

type ReviewDataTaskInfo struct {
	// 主任务号
	MainTaskId *string `json:"MainTaskId,omitempty" name:"MainTaskId"`

	// 子任务号
	SubTaskId *string `json:"SubTaskId,omitempty" name:"SubTaskId"`

	// 任务名
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 任务类型:HealthReport(体检报告); BUltraReport(B超报告);MedCheckReport(检查报告);LaboratoryReport(检验报告); PathologyReport(病理报告);AdmissionReport(入院记录);DischargeReport(出院记录); DischargeSummary(出院小结);DiagnosisReport(诊断证明); MedicalRecordFront(病案首页);OperationReport(手术记录);OutpatientMedicalRecord(门诊病历)
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`
}

type StructureModifyItem struct {
	// 修改的字段的路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitempty" name:"Path"`

	// 机器结果的值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Machine *string `json:"Machine,omitempty" name:"Machine"`

	// 人工结果的值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Manual *string `json:"Manual,omitempty" name:"Manual"`
}

type StructureOneItem struct {
	// 新字段的路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitempty" name:"Path"`

	// 字段的值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type StructureResultObject struct {
	// 0表示正常返回；1代表结果未生成；2代表任务执行失败
	Code *uint64 `json:"Code,omitempty" name:"Code"`

	// 报告类型:HealthReport(体检报告); BUltraReport(B超报告);MedCheckReport(检查报告);LaboratoryReport(检验报告); PathologyReport(病理报告);AdmissionReport(入院记录);DischargeReport(出院记录); DischargeSummary(出院小结);DiagnosisReport(诊断证明); MedicalRecordFront(病案首页);OperationReport(手术记录);OutpatientMedicalRecord(门诊病历)
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 结构化结果
	StructureResult *string `json:"StructureResult,omitempty" name:"StructureResult"`

	// 子任务ID
	SubTaskId *string `json:"SubTaskId,omitempty" name:"SubTaskId"`

	// 任务文件列表
	TaskFiles []*string `json:"TaskFiles,omitempty" name:"TaskFiles"`

	// 结构化字段结果数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultFields []*OcrRecognise `json:"ResultFields,omitempty" name:"ResultFields"`
}

type UnderwriteConclusion struct {
	// 类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 结论
	Conclusion *string `json:"Conclusion,omitempty" name:"Conclusion"`

	// 解释
	Explanation *string `json:"Explanation,omitempty" name:"Explanation"`
}

type UnderwriteItem struct {
	// 字段名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 结果
	Result *string `json:"Result,omitempty" name:"Result"`

	// 风险值或者说明
	Value *string `json:"Value,omitempty" name:"Value"`

	// 参考范围
	// 注意：此字段可能返回 null，表示取不到有效值。
	Range *string `json:"Range,omitempty" name:"Range"`

	// 报告时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportDate []*string `json:"ReportDate,omitempty" name:"ReportDate"`

	// 文件类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 检查项目
	// 注意：此字段可能返回 null，表示取不到有效值。
	InspectProject *string `json:"InspectProject,omitempty" name:"InspectProject"`

	// 单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// 原名
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginName *string `json:"OriginName,omitempty" name:"OriginName"`

	// 阴阳性
	// 注意：此字段可能返回 null，表示取不到有效值。
	YinYang *string `json:"YinYang,omitempty" name:"YinYang"`
}

type UnderwriteOutput struct {
	// 客户ID
	CustomerId *string `json:"CustomerId,omitempty" name:"CustomerId"`

	// 客户姓名
	CustomerName *string `json:"CustomerName,omitempty" name:"CustomerName"`

	// 结果
	Results []*InsuranceResult `json:"Results,omitempty" name:"Results"`

	// 复核时间
	ReviewTime *string `json:"ReviewTime,omitempty" name:"ReviewTime"`

	// 人工复核结果
	ManualDetail []*UnderwriteConclusion `json:"ManualDetail,omitempty" name:"ManualDetail"`
}

// Predefined struct for user
type UploadMedicalFileRequestParams struct {
	// 文件的字节内容。File与FileURL有一个不为空即可，若FileURL参数也存在，会只取File的内容。
	File *string `json:"File,omitempty" name:"File"`

	// 文件的URL地址。File与FileURL不能同时为空，若File参数也存在，会只取File的内容。
	FileURL *string `json:"FileURL,omitempty" name:"FileURL"`
}

type UploadMedicalFileRequest struct {
	*tchttp.BaseRequest
	
	// 文件的字节内容。File与FileURL有一个不为空即可，若FileURL参数也存在，会只取File的内容。
	File *string `json:"File,omitempty" name:"File"`

	// 文件的URL地址。File与FileURL不能同时为空，若File参数也存在，会只取File的内容。
	FileURL *string `json:"FileURL,omitempty" name:"FileURL"`
}

func (r *UploadMedicalFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadMedicalFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "File")
	delete(f, "FileURL")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadMedicalFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadMedicalFileResponseParams struct {
	// 文件存储的key，可以用来创建结构化任务。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileKey *string `json:"FileKey,omitempty" name:"FileKey"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UploadMedicalFileResponse struct {
	*tchttp.BaseResponse
	Response *UploadMedicalFileResponseParams `json:"Response"`
}

func (r *UploadMedicalFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadMedicalFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}