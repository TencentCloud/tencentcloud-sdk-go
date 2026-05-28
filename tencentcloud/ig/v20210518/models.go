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

package v20210518

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type DescribeIgOrderListRequestParams struct {
	// <p>页码</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页数目</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>产品类型</p><p>枚举值：</p><ul><li>ig： 导诊</li><li>ipc： 预问诊</li></ul><p>默认值：ig</p>
	ProductType *string `json:"ProductType,omitnil,omitempty" name:"ProductType"`

	// <p>订单状态</p><p>枚举值：</p><ul><li>0： 无状态</li><li>1： 未激活</li><li>2： 使用中</li><li>3： 暂停使用</li><li>4： 已到期</li><li>5： 已删除</li><li>6： 已失效</li></ul><p>默认值：0</p>
	OrderStatus *int64 `json:"OrderStatus,omitnil,omitempty" name:"OrderStatus"`

	// <p>搜索关键字</p>
	KeyWord *string `json:"KeyWord,omitnil,omitempty" name:"KeyWord"`
}

type DescribeIgOrderListRequest struct {
	*tchttp.BaseRequest
	
	// <p>页码</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页数目</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>产品类型</p><p>枚举值：</p><ul><li>ig： 导诊</li><li>ipc： 预问诊</li></ul><p>默认值：ig</p>
	ProductType *string `json:"ProductType,omitnil,omitempty" name:"ProductType"`

	// <p>订单状态</p><p>枚举值：</p><ul><li>0： 无状态</li><li>1： 未激活</li><li>2： 使用中</li><li>3： 暂停使用</li><li>4： 已到期</li><li>5： 已删除</li><li>6： 已失效</li></ul><p>默认值：0</p>
	OrderStatus *int64 `json:"OrderStatus,omitnil,omitempty" name:"OrderStatus"`

	// <p>搜索关键字</p>
	KeyWord *string `json:"KeyWord,omitnil,omitempty" name:"KeyWord"`
}

func (r *DescribeIgOrderListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIgOrderListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "ProductType")
	delete(f, "OrderStatus")
	delete(f, "KeyWord")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIgOrderListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIgOrderListResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIgOrderListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIgOrderListResponseParams `json:"Response"`
}

func (r *DescribeIgOrderListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIgOrderListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DrugCardInfo struct {
	// 药品id
	DrugId *string `json:"DrugId,omitnil,omitempty" name:"DrugId"`

	// 药品名称
	DrugName *string `json:"DrugName,omitnil,omitempty" name:"DrugName"`

	// 商品名称
	TradeName *string `json:"TradeName,omitnil,omitempty" name:"TradeName"`

	// 规格
	Specification *string `json:"Specification,omitnil,omitempty" name:"Specification"`

	// 生产厂商
	Manufacturer *string `json:"Manufacturer,omitnil,omitempty" name:"Manufacturer"`

	// 是否处方药，0:非处方药，1:处方药
	DrugRxType *int64 `json:"DrugRxType,omitnil,omitempty" name:"DrugRxType"`

	// 药品说明书URL
	DocUrl *string `json:"DocUrl,omitnil,omitempty" name:"DocUrl"`

	// 识别来源，1:药品图片，2:用户输入
	IdentifySource *int64 `json:"IdentifySource,omitnil,omitempty" name:"IdentifySource"`
}

type DrugInstructionInfo struct {
	// 药品id
	DrugId *string `json:"DrugId,omitnil,omitempty" name:"DrugId"`

	// 药品名称
	DrugName *string `json:"DrugName,omitnil,omitempty" name:"DrugName"`

	// 商品名称
	TradeName *string `json:"TradeName,omitnil,omitempty" name:"TradeName"`

	// 规格
	Specification *string `json:"Specification,omitnil,omitempty" name:"Specification"`

	// 生产企业
	Manufacturer *string `json:"Manufacturer,omitnil,omitempty" name:"Manufacturer"`

	// 用法用量
	DrugUsage *string `json:"DrugUsage,omitnil,omitempty" name:"DrugUsage"`

	// 适应症
	Indication *string `json:"Indication,omitnil,omitempty" name:"Indication"`

	// 不良反应
	AdverseReaction *string `json:"AdverseReaction,omitnil,omitempty" name:"AdverseReaction"`

	// 是否处方药，0:非处方药，1:处方药
	DrugRxType *int64 `json:"DrugRxType,omitnil,omitempty" name:"DrugRxType"`

	// 药品说明书URL
	DocUrl *string `json:"DocUrl,omitnil,omitempty" name:"DocUrl"`
}

// Predefined struct for user
type GetLLMDiagnosisDrugChatRequestParams struct {
	// <p>合作方ID</p>
	PartnerId *string `json:"PartnerId,omitnil,omitempty" name:"PartnerId"`

	// <p>合作方密钥</p>
	PartnerSecret *string `json:"PartnerSecret,omitnil,omitempty" name:"PartnerSecret"`

	// <p>医院应用ID</p>
	HospitalAppId *string `json:"HospitalAppId,omitnil,omitempty" name:"HospitalAppId"`

	// <p>对话ID，由接入方生成，相同对话ID会保持相同的上下文，接入方根据业务场景使用相同或新建对话ID（建议使用UUID值）</p><p>入参限制：长度限制10-50</p>
	DialogueId *string `json:"DialogueId,omitnil,omitempty" name:"DialogueId"`

	// <p>本次问答用户输入的问题，（最大长度1000）</p>
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// <p>返回类型：0:正常返回； 1:流式返回； 默认：0</p>
	ResultType *int64 `json:"ResultType,omitnil,omitempty" name:"ResultType"`

	// <p>提示词</p>
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`
}

type GetLLMDiagnosisDrugChatRequest struct {
	*tchttp.BaseRequest
	
	// <p>合作方ID</p>
	PartnerId *string `json:"PartnerId,omitnil,omitempty" name:"PartnerId"`

	// <p>合作方密钥</p>
	PartnerSecret *string `json:"PartnerSecret,omitnil,omitempty" name:"PartnerSecret"`

	// <p>医院应用ID</p>
	HospitalAppId *string `json:"HospitalAppId,omitnil,omitempty" name:"HospitalAppId"`

	// <p>对话ID，由接入方生成，相同对话ID会保持相同的上下文，接入方根据业务场景使用相同或新建对话ID（建议使用UUID值）</p><p>入参限制：长度限制10-50</p>
	DialogueId *string `json:"DialogueId,omitnil,omitempty" name:"DialogueId"`

	// <p>本次问答用户输入的问题，（最大长度1000）</p>
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// <p>返回类型：0:正常返回； 1:流式返回； 默认：0</p>
	ResultType *int64 `json:"ResultType,omitnil,omitempty" name:"ResultType"`

	// <p>提示词</p>
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`
}

func (r *GetLLMDiagnosisDrugChatRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetLLMDiagnosisDrugChatRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PartnerId")
	delete(f, "PartnerSecret")
	delete(f, "HospitalAppId")
	delete(f, "DialogueId")
	delete(f, "Message")
	delete(f, "ResultType")
	delete(f, "Prompt")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetLLMDiagnosisDrugChatRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetLLMDiagnosisDrugChatResponseParams struct {
	// <p>错误码，0成功，非0失败</p>
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// <p>错误信息</p>
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// <p>返回数据</p>
	Data *LLMDiagnosisDrugData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。本接口为流式响应接口，当请求成功时，RequestId 会被放在 HTTP 响应的 Header "X-TC-RequestId" 中。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetLLMDiagnosisDrugChatResponse struct {
	tchttp.BaseSSEResponse `json:"-"`
	Response *GetLLMDiagnosisDrugChatResponseParams `json:"Response"`
}

func (r *GetLLMDiagnosisDrugChatResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetLLMDiagnosisDrugChatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetLLMDiagnosisDrugRequestParams struct {
	// <p>合作方ID</p>
	PartnerId *string `json:"PartnerId,omitnil,omitempty" name:"PartnerId"`

	// <p>合作方密钥</p>
	PartnerSecret *string `json:"PartnerSecret,omitnil,omitempty" name:"PartnerSecret"`

	// <p>医院应用ID</p>
	HospitalAppId *string `json:"HospitalAppId,omitnil,omitempty" name:"HospitalAppId"`

	// <p>对话ID，由接入方生成，相同对话ID会保持相同的上下文，接入方根据业务场景使用相同或新建对话ID（建议使用UUID值）</p><p>入参限制：长度限制10-50</p>
	DialogueId *string `json:"DialogueId,omitnil,omitempty" name:"DialogueId"`

	// <p>本次问答用户输入的问题，（最大长度1000，传了图片链接，此参数可为空）</p>
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// <p>药品图片链接</p>
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// <p>返回类型：0:正常返回； 1:流式返回； 默认：0</p>
	ResultType *int64 `json:"ResultType,omitnil,omitempty" name:"ResultType"`

	// <p>提示词</p>
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`
}

type GetLLMDiagnosisDrugRequest struct {
	*tchttp.BaseRequest
	
	// <p>合作方ID</p>
	PartnerId *string `json:"PartnerId,omitnil,omitempty" name:"PartnerId"`

	// <p>合作方密钥</p>
	PartnerSecret *string `json:"PartnerSecret,omitnil,omitempty" name:"PartnerSecret"`

	// <p>医院应用ID</p>
	HospitalAppId *string `json:"HospitalAppId,omitnil,omitempty" name:"HospitalAppId"`

	// <p>对话ID，由接入方生成，相同对话ID会保持相同的上下文，接入方根据业务场景使用相同或新建对话ID（建议使用UUID值）</p><p>入参限制：长度限制10-50</p>
	DialogueId *string `json:"DialogueId,omitnil,omitempty" name:"DialogueId"`

	// <p>本次问答用户输入的问题，（最大长度1000，传了图片链接，此参数可为空）</p>
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// <p>药品图片链接</p>
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// <p>返回类型：0:正常返回； 1:流式返回； 默认：0</p>
	ResultType *int64 `json:"ResultType,omitnil,omitempty" name:"ResultType"`

	// <p>提示词</p>
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`
}

func (r *GetLLMDiagnosisDrugRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetLLMDiagnosisDrugRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PartnerId")
	delete(f, "PartnerSecret")
	delete(f, "HospitalAppId")
	delete(f, "DialogueId")
	delete(f, "Message")
	delete(f, "ImageUrl")
	delete(f, "ResultType")
	delete(f, "Prompt")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetLLMDiagnosisDrugRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetLLMDiagnosisDrugResponseParams struct {
	// <p>错误码，0成功，非0失败</p>
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// <p>错误信息</p>
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// <p>返回数据</p>
	Data *LLMDiagnosisDrugData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。本接口为流式响应接口，当请求成功时，RequestId 会被放在 HTTP 响应的 Header "X-TC-RequestId" 中。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetLLMDiagnosisDrugResponse struct {
	tchttp.BaseSSEResponse `json:"-"`
	Response *GetLLMDiagnosisDrugResponseParams `json:"Response"`
}

func (r *GetLLMDiagnosisDrugResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetLLMDiagnosisDrugResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetLLMDiagnosisHealthRequestParams struct {
	// <p>合作方ID</p>
	PartnerId *string `json:"PartnerId,omitnil,omitempty" name:"PartnerId"`

	// <p>合作方密钥</p>
	PartnerSecret *string `json:"PartnerSecret,omitnil,omitempty" name:"PartnerSecret"`

	// <p>医院应用ID</p>
	HospitalAppId *string `json:"HospitalAppId,omitnil,omitempty" name:"HospitalAppId"`

	// <p>对话ID，由接入方生成，相同对话ID会保持相同的上下文，接入方根据业务场景使用相同或新建对话ID（建议使用UUID值）</p><p>入参限制：长度限制10-50</p>
	DialogueId *string `json:"DialogueId,omitnil,omitempty" name:"DialogueId"`

	// <p>本次问答用户输入的问题，（最大长度1000）</p>
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// <p>返回类型：0:正常返回； 1:流式返回； 默认：0</p>
	ResultType *int64 `json:"ResultType,omitnil,omitempty" name:"ResultType"`

	// <p>性别，0:未知，1:男，2:女</p>
	Sex *int64 `json:"Sex,omitnil,omitempty" name:"Sex"`

	// <p>年龄，0:未知，最大值：120</p>
	Age *int64 `json:"Age,omitnil,omitempty" name:"Age"`

	// <p>追问轮数，-1:不追问；0:使用默认规则； 1-10:按指定轮数追问；（最大值10轮）； 默认：0</p>
	RoundNumber *int64 `json:"RoundNumber,omitnil,omitempty" name:"RoundNumber"`

	// <p>是否追问结构化结果，0：否，1：是；</p>
	OutputStructured *int64 `json:"OutputStructured,omitnil,omitempty" name:"OutputStructured"`
}

type GetLLMDiagnosisHealthRequest struct {
	*tchttp.BaseRequest
	
	// <p>合作方ID</p>
	PartnerId *string `json:"PartnerId,omitnil,omitempty" name:"PartnerId"`

	// <p>合作方密钥</p>
	PartnerSecret *string `json:"PartnerSecret,omitnil,omitempty" name:"PartnerSecret"`

	// <p>医院应用ID</p>
	HospitalAppId *string `json:"HospitalAppId,omitnil,omitempty" name:"HospitalAppId"`

	// <p>对话ID，由接入方生成，相同对话ID会保持相同的上下文，接入方根据业务场景使用相同或新建对话ID（建议使用UUID值）</p><p>入参限制：长度限制10-50</p>
	DialogueId *string `json:"DialogueId,omitnil,omitempty" name:"DialogueId"`

	// <p>本次问答用户输入的问题，（最大长度1000）</p>
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// <p>返回类型：0:正常返回； 1:流式返回； 默认：0</p>
	ResultType *int64 `json:"ResultType,omitnil,omitempty" name:"ResultType"`

	// <p>性别，0:未知，1:男，2:女</p>
	Sex *int64 `json:"Sex,omitnil,omitempty" name:"Sex"`

	// <p>年龄，0:未知，最大值：120</p>
	Age *int64 `json:"Age,omitnil,omitempty" name:"Age"`

	// <p>追问轮数，-1:不追问；0:使用默认规则； 1-10:按指定轮数追问；（最大值10轮）； 默认：0</p>
	RoundNumber *int64 `json:"RoundNumber,omitnil,omitempty" name:"RoundNumber"`

	// <p>是否追问结构化结果，0：否，1：是；</p>
	OutputStructured *int64 `json:"OutputStructured,omitnil,omitempty" name:"OutputStructured"`
}

func (r *GetLLMDiagnosisHealthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetLLMDiagnosisHealthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PartnerId")
	delete(f, "PartnerSecret")
	delete(f, "HospitalAppId")
	delete(f, "DialogueId")
	delete(f, "Message")
	delete(f, "ResultType")
	delete(f, "Sex")
	delete(f, "Age")
	delete(f, "RoundNumber")
	delete(f, "OutputStructured")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetLLMDiagnosisHealthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetLLMDiagnosisHealthResponseParams struct {
	// <p>错误码，0成功，非0失败</p>
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// <p>错误信息</p>
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// <p>返回数据</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *LLMDiagnosisHealthData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。本接口为流式响应接口，当请求成功时，RequestId 会被放在 HTTP 响应的 Header "X-TC-RequestId" 中。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetLLMDiagnosisHealthResponse struct {
	tchttp.BaseSSEResponse `json:"-"`
	Response *GetLLMDiagnosisHealthResponseParams `json:"Response"`
}

func (r *GetLLMDiagnosisHealthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetLLMDiagnosisHealthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetLLMReportInterpretationRequestParams struct {
	// <p>合作方ID</p>
	PartnerId *string `json:"PartnerId,omitnil,omitempty" name:"PartnerId"`

	// <p>合作方密钥</p>
	PartnerSecret *string `json:"PartnerSecret,omitnil,omitempty" name:"PartnerSecret"`

	// <p>接入医院id</p>
	HospitalId *string `json:"HospitalId,omitnil,omitempty" name:"HospitalId"`

	// <p>对话ID，由接入方生成，相同对话ID会保持相同的上下文，接入方根据业务场景使用相同或新建对话ID（建议使用UUID值）</p><p>入参限制：长度限制10-50</p>
	DialogueId *string `json:"DialogueId,omitnil,omitempty" name:"DialogueId"`

	// <p>本次问答用户输入的问题</p>
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// <p>报告文件信息</p>
	ReportFileInfos []*ReportFileInfoReq `json:"ReportFileInfos,omitnil,omitempty" name:"ReportFileInfos"`

	// <p>返回类型：0:正常返回； 1:流式返回； 默认：0</p>
	ResultType *int64 `json:"ResultType,omitnil,omitempty" name:"ResultType"`

	// <p>报告解读提示词</p>
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// <p>无报告问答提示词</p>
	QaPrompt *string `json:"QaPrompt,omitnil,omitempty" name:"QaPrompt"`
}

type GetLLMReportInterpretationRequest struct {
	*tchttp.BaseRequest
	
	// <p>合作方ID</p>
	PartnerId *string `json:"PartnerId,omitnil,omitempty" name:"PartnerId"`

	// <p>合作方密钥</p>
	PartnerSecret *string `json:"PartnerSecret,omitnil,omitempty" name:"PartnerSecret"`

	// <p>接入医院id</p>
	HospitalId *string `json:"HospitalId,omitnil,omitempty" name:"HospitalId"`

	// <p>对话ID，由接入方生成，相同对话ID会保持相同的上下文，接入方根据业务场景使用相同或新建对话ID（建议使用UUID值）</p><p>入参限制：长度限制10-50</p>
	DialogueId *string `json:"DialogueId,omitnil,omitempty" name:"DialogueId"`

	// <p>本次问答用户输入的问题</p>
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// <p>报告文件信息</p>
	ReportFileInfos []*ReportFileInfoReq `json:"ReportFileInfos,omitnil,omitempty" name:"ReportFileInfos"`

	// <p>返回类型：0:正常返回； 1:流式返回； 默认：0</p>
	ResultType *int64 `json:"ResultType,omitnil,omitempty" name:"ResultType"`

	// <p>报告解读提示词</p>
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// <p>无报告问答提示词</p>
	QaPrompt *string `json:"QaPrompt,omitnil,omitempty" name:"QaPrompt"`
}

func (r *GetLLMReportInterpretationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetLLMReportInterpretationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PartnerId")
	delete(f, "PartnerSecret")
	delete(f, "HospitalId")
	delete(f, "DialogueId")
	delete(f, "Message")
	delete(f, "ReportFileInfos")
	delete(f, "ResultType")
	delete(f, "Prompt")
	delete(f, "QaPrompt")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetLLMReportInterpretationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetLLMReportInterpretationResponseParams struct {
	// <p>错误码，0成功，非0失败</p>
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// <p>错误信息</p>
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// <p>返回数据</p>
	Data *LLMReportInterpretationData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。本接口为流式响应接口，当请求成功时，RequestId 会被放在 HTTP 响应的 Header "X-TC-RequestId" 中。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetLLMReportInterpretationResponse struct {
	tchttp.BaseSSEResponse `json:"-"`
	Response *GetLLMReportInterpretationResponseParams `json:"Response"`
}

func (r *GetLLMReportInterpretationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetLLMReportInterpretationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GuessQuestion struct {
	// 问题标题
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`
}

type HealthFollowUpQuestion struct {
	// <p>追问类型</p>
	PromptType *string `json:"PromptType,omitnil,omitempty" name:"PromptType"`

	// <p>追问选项列表</p>
	Choices []*string `json:"Choices,omitnil,omitempty" name:"Choices"`

	// <p>追问题型，单选或多选</p>
	QuestionType *string `json:"QuestionType,omitnil,omitempty" name:"QuestionType"`

	// <p>追问特殊类型，如部位题、时间题</p>
	SpecialType *string `json:"SpecialType,omitnil,omitempty" name:"SpecialType"`

	// <p>追问问题内容</p>
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`
}

type HighlightWordInfo struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 类型 1:疾病，2:检验/检查，3:药品，4:科室，5:自定义配置
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 跳转类型 h5:h5类型，mini_wx:微信小程序
	JumpType *string `json:"JumpType,omitnil,omitempty" name:"JumpType"`

	// 跳转URL
	JumpUrl *string `json:"JumpUrl,omitnil,omitempty" name:"JumpUrl"`

	// 跳转小程序Appid
	JumpAppid *string `json:"JumpAppid,omitnil,omitempty" name:"JumpAppid"`

	// 跳转小程序原始ID
	JumpOriginId *string `json:"JumpOriginId,omitnil,omitempty" name:"JumpOriginId"`
}

type LLMDiagnosisDrugData struct {
	// <p>大模型返回问答内容</p>
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// <p>大模型返回思考内容</p>
	Think *string `json:"Think,omitnil,omitempty" name:"Think"`

	// <p>序号，无业务含义，标识流式结果的序号</p>
	Sort *int64 `json:"Sort,omitnil,omitempty" name:"Sort"`

	// <p>流式输出结束标识，false:未结束，true:结束</p>
	IsFinish *bool `json:"IsFinish,omitnil,omitempty" name:"IsFinish"`

	// <p>是否命中敏感词，false:否，true:是；</p>
	IsSensitive *bool `json:"IsSensitive,omitnil,omitempty" name:"IsSensitive"`

	// <p>引用资料列表，流式输出方式，只在流式输出第一次返回；</p>
	ReferResourceItems []*ReferResourceInfo `json:"ReferResourceItems,omitnil,omitempty" name:"ReferResourceItems"`

	// <p>猜你想问列表，流式输出方式，只在流式结束输出结果；</p>
	GuessQuestions []*GuessQuestion `json:"GuessQuestions,omitnil,omitempty" name:"GuessQuestions"`

	// <p>高亮词列表，流式输出方式在流式过程中输出结果。</p>
	HighlightWords []*HighlightWordInfo `json:"HighlightWords,omitnil,omitempty" name:"HighlightWords"`

	// <p>药品列表，流式输出方式，只在流式结束输出结果。</p>
	DrugList []*StandardDrugCardInfo `json:"DrugList,omitnil,omitempty" name:"DrugList"`
}

type LLMDiagnosisHealthData struct {
	// <p>大模型返回问答内容</p>
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// <p>大模型返回思考内容</p>
	Think *string `json:"Think,omitnil,omitempty" name:"Think"`

	// <p>序号，无业务含义，标识流式结果的序号</p>
	Sort *int64 `json:"Sort,omitnil,omitempty" name:"Sort"`

	// <p>流式输出结束标识，false:未结束，true:结束</p>
	IsFinish *bool `json:"IsFinish,omitnil,omitempty" name:"IsFinish"`

	// <p>是否命中敏感词，false:否，true:是；</p>
	IsSensitive *bool `json:"IsSensitive,omitnil,omitempty" name:"IsSensitive"`

	// <p>结果类型，0：正常结果，1：追问问题</p>
	Kind *int64 `json:"Kind,omitnil,omitempty" name:"Kind"`

	// <p>风险提示</p>
	RiskReminder *string `json:"RiskReminder,omitnil,omitempty" name:"RiskReminder"`

	// <p>引用资料列表，流式输出方式，只在流式输出第一次返回；</p>
	ReferResourceItems []*ReferResourceInfo `json:"ReferResourceItems,omitnil,omitempty" name:"ReferResourceItems"`

	// <p>猜你想问列表，流式输出方式，只在流式结束输出结果；</p>
	GuessQuestions []*GuessQuestion `json:"GuessQuestions,omitnil,omitempty" name:"GuessQuestions"`

	// <p>高亮词列表，流式输出方式在流式过程中输出结果。</p>
	HighlightWords []*HighlightWordInfo `json:"HighlightWords,omitnil,omitempty" name:"HighlightWords"`

	// <p>追问问题列表，流式输出方式，只在流式结束输出结果。只会在kind=1追问类型时有结果。</p>
	FollowUpQuestions []*HealthFollowUpQuestion `json:"FollowUpQuestions,omitnil,omitempty" name:"FollowUpQuestions"`

	// <p>药品列表，流式输出方式，只在流式结束输出结果。</p>
	DrugList []*StandardDrugCardInfo `json:"DrugList,omitnil,omitempty" name:"DrugList"`
}

type LLMReportInterpretationData struct {
	// <p>大模型返回问答内容</p>
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// <p>大模型返回思考内容</p>
	Think *string `json:"Think,omitnil,omitempty" name:"Think"`

	// <p>序号，无业务含义，标识流式结果的序号</p>
	Sort *int64 `json:"Sort,omitnil,omitempty" name:"Sort"`

	// <p>流式输出结束标识，false:未结束，true:结束</p>
	IsFinish *bool `json:"IsFinish,omitnil,omitempty" name:"IsFinish"`

	// <p>是否命中敏感词，false:否，true:是；</p>
	IsSensitive *bool `json:"IsSensitive,omitnil,omitempty" name:"IsSensitive"`

	// <p>是否支持报告类型，false:否，true:是； 默认返回true，不支持的报告类型返回false</p>
	IsSupportFile *bool `json:"IsSupportFile,omitnil,omitempty" name:"IsSupportFile"`

	// <p>返回的报告列表，和传入报告文件顺序一致</p>
	ReportInfos []*ReportFileInfoRsp `json:"ReportInfos,omitnil,omitempty" name:"ReportInfos"`

	// <p>引用资料列表，流式输出方式，只在流式输出第一次返回；</p>
	ReferResourceItems []*ReferResourceInfo `json:"ReferResourceItems,omitnil,omitempty" name:"ReferResourceItems"`

	// <p>猜你想问列表，流式输出方式，只在流式结束输出结果；</p>
	GuessQuestions []*GuessQuestion `json:"GuessQuestions,omitnil,omitempty" name:"GuessQuestions"`

	// <p>高亮词列表，流式输出方式在流式过程中输出结果。</p>
	HighlightWords []*HighlightWordInfo `json:"HighlightWords,omitnil,omitempty" name:"HighlightWords"`
}

// Predefined struct for user
type QueryDrugInstructionsRequestParams struct {
	// <p>合作方ID</p>
	PartnerId *string `json:"PartnerId,omitnil,omitempty" name:"PartnerId"`

	// <p>合作方密钥</p>
	PartnerSecret *string `json:"PartnerSecret,omitnil,omitempty" name:"PartnerSecret"`

	// <p>医院应用ID</p>
	HospitalAppId *string `json:"HospitalAppId,omitnil,omitempty" name:"HospitalAppId"`

	// <p>本次问答用户输入的问题，（最大长度1000）</p>
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

type QueryDrugInstructionsRequest struct {
	*tchttp.BaseRequest
	
	// <p>合作方ID</p>
	PartnerId *string `json:"PartnerId,omitnil,omitempty" name:"PartnerId"`

	// <p>合作方密钥</p>
	PartnerSecret *string `json:"PartnerSecret,omitnil,omitempty" name:"PartnerSecret"`

	// <p>医院应用ID</p>
	HospitalAppId *string `json:"HospitalAppId,omitnil,omitempty" name:"HospitalAppId"`

	// <p>本次问答用户输入的问题，（最大长度1000）</p>
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

func (r *QueryDrugInstructionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryDrugInstructionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PartnerId")
	delete(f, "PartnerSecret")
	delete(f, "HospitalAppId")
	delete(f, "Message")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryDrugInstructionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryDrugInstructionsResponseParams struct {
	// <p>错误码，0成功，非0失败</p>
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// <p>错误信息</p>
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// <p>返回数据</p>
	Data []*StandardDrugInstructionInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryDrugInstructionsResponse struct {
	*tchttp.BaseResponse
	Response *QueryDrugInstructionsResponseParams `json:"Response"`
}

func (r *QueryDrugInstructionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryDrugInstructionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReferResourceInfo struct {
	// 资料标题
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 资料来源类型，1:问答库，2:文档，3:医典百科，4:临床指南，5:药品说明书
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 资料详情链接
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type ReportFileInfoReq struct {
	// <p>报告文件链接</p>
	ReportFileUrl *string `json:"ReportFileUrl,omitnil,omitempty" name:"ReportFileUrl"`

	// <p>报告文件类型，1:pdf，2:图片</p>
	ReportFileType *int64 `json:"ReportFileType,omitnil,omitempty" name:"ReportFileType"`

	// <p>报告id</p>
	ReportId *string `json:"ReportId,omitnil,omitempty" name:"ReportId"`
}

type ReportFileInfoRsp struct {
	// <p>报告id</p>
	ReportId *string `json:"ReportId,omitnil,omitempty" name:"ReportId"`

	// <p>是否解析成功</p>
	IsAnalysis *bool `json:"IsAnalysis,omitnil,omitempty" name:"IsAnalysis"`

	// <p>解析错误信息</p>
	ErrInfo *string `json:"ErrInfo,omitnil,omitempty" name:"ErrInfo"`
}

type StandardDrugCardInfo struct {
	// 标准药品名
	StandardDrugName *string `json:"StandardDrugName,omitnil,omitempty" name:"StandardDrugName"`

	// 药品列表
	DrugInfos []*DrugCardInfo `json:"DrugInfos,omitnil,omitempty" name:"DrugInfos"`
}

type StandardDrugInstructionInfo struct {
	// 标准药品名
	StandardDrugName *string `json:"StandardDrugName,omitnil,omitempty" name:"StandardDrugName"`

	// 药品列表
	DrugInfos []*DrugInstructionInfo `json:"DrugInfos,omitnil,omitempty" name:"DrugInfos"`
}