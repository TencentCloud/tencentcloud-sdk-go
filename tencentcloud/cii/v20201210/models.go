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

package v20201210

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type CompareMetricsData struct {
	// 短文准确率
	ShortStructAccuracy *string `json:"ShortStructAccuracy,omitempty" name:"ShortStructAccuracy"`

	// 短文召回率
	ShortStructRecall *string `json:"ShortStructRecall,omitempty" name:"ShortStructRecall"`

	// 长文结构化准确率
	LongStructAccuracy *string `json:"LongStructAccuracy,omitempty" name:"LongStructAccuracy"`

	// 长文结构化召回率
	LongStructRecall *string `json:"LongStructRecall,omitempty" name:"LongStructRecall"`

	// 长文提取准确率
	LongContentAccuracy *string `json:"LongContentAccuracy,omitempty" name:"LongContentAccuracy"`

	// 长文提取召回率
	LongContentRecall *string `json:"LongContentRecall,omitempty" name:"LongContentRecall"`
}

// Predefined struct for user
type CreateStructureTaskRequestParams struct {
	// 保单号
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 客户号
	CustomerId *string `json:"CustomerId,omitempty" name:"CustomerId"`

	// 客户姓名
	CustomerName *string `json:"CustomerName,omitempty" name:"CustomerName"`

	// 文件类型，目前只支持体检报告类型，对应的值为：HealthReport
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 报告年份
	Year *string `json:"Year,omitempty" name:"Year"`

	// 报告文件上传的地址列表，需按顺序排列。如果使用ImageList参数，置为空数组即可
	FileList []*string `json:"FileList,omitempty" name:"FileList"`

	// 险种，如果是体检报告类型，此参数是必填，类型说明如下：
	// CriticalDiseaseInsurance:重疾险
	// LifeInsurance：寿险
	// AccidentInsurance：意外险
	InsuranceTypes []*string `json:"InsuranceTypes,omitempty" name:"InsuranceTypes"`

	// 报告上传的图片内容数组，图片内容采用base64编码，需按顺序排列
	ImageList []*string `json:"ImageList,omitempty" name:"ImageList"`
}

type CreateStructureTaskRequest struct {
	*tchttp.BaseRequest
	
	// 保单号
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 客户号
	CustomerId *string `json:"CustomerId,omitempty" name:"CustomerId"`

	// 客户姓名
	CustomerName *string `json:"CustomerName,omitempty" name:"CustomerName"`

	// 文件类型，目前只支持体检报告类型，对应的值为：HealthReport
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 报告年份
	Year *string `json:"Year,omitempty" name:"Year"`

	// 报告文件上传的地址列表，需按顺序排列。如果使用ImageList参数，置为空数组即可
	FileList []*string `json:"FileList,omitempty" name:"FileList"`

	// 险种，如果是体检报告类型，此参数是必填，类型说明如下：
	// CriticalDiseaseInsurance:重疾险
	// LifeInsurance：寿险
	// AccidentInsurance：意外险
	InsuranceTypes []*string `json:"InsuranceTypes,omitempty" name:"InsuranceTypes"`

	// 报告上传的图片内容数组，图片内容采用base64编码，需按顺序排列
	ImageList []*string `json:"ImageList,omitempty" name:"ImageList"`
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
	delete(f, "PolicyId")
	delete(f, "CustomerId")
	delete(f, "CustomerName")
	delete(f, "TaskType")
	delete(f, "Year")
	delete(f, "FileList")
	delete(f, "InsuranceTypes")
	delete(f, "ImageList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStructureTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStructureTaskResponseParams struct {
	// 本次结构化任务的ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

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
type DescribeStructCompareDataRequestParams struct {
	// 结构化任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeStructCompareDataRequest struct {
	*tchttp.BaseRequest
	
	// 结构化任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
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
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStructCompareDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStructCompareDataResponseParams struct {
	// 保单号
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 结构化任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 客户号
	CustomerId *string `json:"CustomerId,omitempty" name:"CustomerId"`

	// 客户姓名
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
type DescribeStructureTaskResultRequestParams struct {
	// 结构化任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeStructureTaskResultRequest struct {
	*tchttp.BaseRequest
	
	// 结构化任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
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
	delete(f, "TaskId")
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

type ResultObject struct {
	// 图片质量分
	Quality *float64 `json:"Quality,omitempty" name:"Quality"`

	// 由结构化算法结构化json转换的字符串，具体协议参见算法结构化结果协议
	StructureResult *string `json:"StructureResult,omitempty" name:"StructureResult"`
}