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

type CreateStructureTaskInfo struct {

	// 任务类型
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

type CreateStructureTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 创建的主任务号，用于查询结果
		MainTaskId *string `json:"MainTaskId,omitempty" name:"MainTaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateUnderwriteTaskByIdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 核保任务ID数据
		UnderwriteTaskIds []*string `json:"UnderwriteTaskIds,omitempty" name:"UnderwriteTaskIds"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeMachineUnderwriteResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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

		// 状态码
		Status *uint64 `json:"Status,omitempty" name:"Status"`

		// 机器核保结果
		UnderwriteResults []*MachineUnderwriteOutput `json:"UnderwriteResults,omitempty" name:"UnderwriteResults"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeStructCompareDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type DescribeStructureDifferenceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type DescribeStructureResultResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type DescribeStructureTaskResultResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type InsuranceResult struct {

	// 险种
	InsuranceType *string `json:"InsuranceType,omitempty" name:"InsuranceType"`

	// 对应险种的机器核保结果
	Result []*MachinePredict `json:"Result,omitempty" name:"Result"`
}

type MachinePredict struct {

	// 核保引擎名称
	Title *string `json:"Title,omitempty" name:"Title"`

	// 核保结论
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

type PerStructDifference struct {

	// 子任务ID
	SubTaskId *string `json:"SubTaskId,omitempty" name:"SubTaskId"`

	// 任务类型
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 修改的项
	ModifyItems []*StructureModifyItem `json:"ModifyItems,omitempty" name:"ModifyItems"`

	// 新增的项
	NewItems []*StructureOneItem `json:"NewItems,omitempty" name:"NewItems"`

	// 删除的项
	RemoveItems []*StructureOneItem `json:"RemoveItems,omitempty" name:"RemoveItems"`
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

	// 任务类型
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

	// 0表示正常返回
	Code *uint64 `json:"Code,omitempty" name:"Code"`

	// 报告类型
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 结构化结果
	StructureResult *string `json:"StructureResult,omitempty" name:"StructureResult"`

	// 子任务ID
	SubTaskId *string `json:"SubTaskId,omitempty" name:"SubTaskId"`
}

type UnderwriteItem struct {

	// 字段名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 结果
	Result *string `json:"Result,omitempty" name:"Result"`

	// 风险值或者说明
	Value *string `json:"Value,omitempty" name:"Value"`
}
