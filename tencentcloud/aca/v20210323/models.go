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

package v20210323

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type Abnormals struct {
	// 类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 标题
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 异常原因 PatientInfo 病人信息缺失；OrderInfo 医嘱信息缺失； PrescriptionError 处方异常提醒
	AbnormalReason *string `json:"AbnormalReason,omitnil,omitempty" name:"AbnormalReason"`
}

type CommonHeader struct {
	// 机构ID
	HospitalId *string `json:"HospitalId,omitnil,omitempty" name:"HospitalId"`

	// 合作方ID
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`
}

type CriticalInfo struct {
	// 危急重症类型 0:文字描述类 1:数值检查类
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 提示
	Tips *string `json:"Tips,omitnil,omitempty" name:"Tips"`
}

type CurrentDiagnosis struct {
	// 诊断疾病
	DiagnoseDisease *string `json:"DiagnoseDisease,omitnil,omitempty" name:"DiagnoseDisease"`

	// 疾病指南信息
	DiseaseGuideInfo *string `json:"DiseaseGuideInfo,omitnil,omitempty" name:"DiseaseGuideInfo"`

	// 标准名称
	StandardName *string `json:"StandardName,omitnil,omitempty" name:"StandardName"`
}

type Department struct {
	// 科室ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 科室名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 科室类型 0:门诊  1:住院  2:门诊+住院
	// 注意：此字段可能返回 null，表示取不到有效值。
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// 门诊区开关 true:此科室对应的门诊区开启智能审方功能, false:此科室对应的门诊区关闭智能审方功能; 仅对scope为0/2的科室生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutpatientOn *bool `json:"OutpatientOn,omitnil,omitempty" name:"OutpatientOn"`

	// 住院区开关 true:此科室对应的住院区开启智能审方功能, false:此科室对应的住院区关闭智能审方功能; 仅对scope为1/2的科室生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	InHospitalOn *bool `json:"InHospitalOn,omitnil,omitempty" name:"InHospitalOn"`
}

type DiagnosisInfo struct {
	// 默认0，0:初诊-常规诊疗 1:复诊 2:检验检查/取药/咨询/疫苗 3:信息缺失 4:信息错误
	IntentType *int64 `json:"IntentType,omitnil,omitempty" name:"IntentType"`

	// 诊断风险
	RiskInfo *string `json:"RiskInfo,omitnil,omitempty" name:"RiskInfo"`

	// 疑似诊断列表
	SuspectedDiagnosis []*SuspectedDiagnosis `json:"SuspectedDiagnosis,omitnil,omitempty" name:"SuspectedDiagnosis"`

	// 转诊提醒
	ReferralInfo *ReferralInfo `json:"ReferralInfo,omitnil,omitempty" name:"ReferralInfo"`

	// 危急重症
	CriticalInfo []*CriticalInfo `json:"CriticalInfo,omitnil,omitempty" name:"CriticalInfo"`

	// 生命体征风险
	VitalSignsInfo *VitalSignsInfo `json:"VitalSignsInfo,omitnil,omitempty" name:"VitalSignsInfo"`

	// 鉴别诊断
	DifferDiagnosis []*DifferDiagnosis `json:"DifferDiagnosis,omitnil,omitempty" name:"DifferDiagnosis"`

	// 病历质控
	RecordQuality *RecordQuality `json:"RecordQuality,omitnil,omitempty" name:"RecordQuality"`

	// 当前诊断
	CurrentDiagnosis []*CurrentDiagnosis `json:"CurrentDiagnosis,omitnil,omitempty" name:"CurrentDiagnosis"`

	// 治疗方案
	TreatmentGuide []*TreatmentGuide `json:"TreatmentGuide,omitnil,omitempty" name:"TreatmentGuide"`

	// 病历质控
	EmrQuality *EmrQuality `json:"EmrQuality,omitnil,omitempty" name:"EmrQuality"`

	// 健康处方
	HealthPrescriptions []*HealthPrescriptions `json:"HealthPrescriptions,omitnil,omitempty" name:"HealthPrescriptions"`
}

type Dict struct {
	// 给药频次编码
	FreqCode *string `json:"FreqCode,omitnil,omitempty" name:"FreqCode"`

	// 给药频次名称
	FreqName *string `json:"FreqName,omitnil,omitempty" name:"FreqName"`

	// 是否禁用 0-启用 1-禁用
	Disable *int64 `json:"Disable,omitnil,omitempty" name:"Disable"`

	// 给药途径编码
	UsageCode *string `json:"UsageCode,omitnil,omitempty" name:"UsageCode"`

	// 给药途径名称
	UsageName *string `json:"UsageName,omitnil,omitempty" name:"UsageName"`

	// 科室ID
	DeptId *string `json:"DeptId,omitnil,omitempty" name:"DeptId"`

	// 科室名称
	DeptName *string `json:"DeptName,omitnil,omitempty" name:"DeptName"`

	// 科室区域类型 0:门诊  1:住院  2:门诊+住院
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// 门诊开关
	OutpatientOn *bool `json:"OutpatientOn,omitnil,omitempty" name:"OutpatientOn"`

	// 住院
	InHospitalOn *bool `json:"InHospitalOn,omitnil,omitempty" name:"InHospitalOn"`

	// 诊断编码
	DiagCode *string `json:"DiagCode,omitnil,omitempty" name:"DiagCode"`

	// 诊断名称
	DiagName *string `json:"DiagName,omitnil,omitempty" name:"DiagName"`

	// ICD代码
	IcdCode *string `json:"IcdCode,omitnil,omitempty" name:"IcdCode"`
}

type DifferDiagnosis struct {
	// 鉴别名称
	DifferName *string `json:"DifferName,omitnil,omitempty" name:"DifferName"`

	// 鉴别提示
	DifferTips *string `json:"DifferTips,omitnil,omitempty" name:"DifferTips"`

	// 疾病指南信息
	DiseaseGuideInfo *string `json:"DiseaseGuideInfo,omitnil,omitempty" name:"DiseaseGuideInfo"`
}

type DocInfo struct {
	// 药品ID
	DrugId *string `json:"DrugId,omitnil,omitempty" name:"DrugId"`

	// 药品名称
	DrugName *string `json:"DrugName,omitnil,omitempty" name:"DrugName"`

	// 说明书地址
	DocUrl *string `json:"DocUrl,omitnil,omitempty" name:"DocUrl"`
}

type DoctorInfo struct {
	// 医生ID
	DoctorId *string `json:"DoctorId,omitnil,omitempty" name:"DoctorId"`

	// 医生姓名
	DoctorName *string `json:"DoctorName,omitnil,omitempty" name:"DoctorName"`

	// 医生电话
	DoctorPhone *string `json:"DoctorPhone,omitnil,omitempty" name:"DoctorPhone"`
}

type Drug struct {
	// 医院药品id
	DrugOrgId *string `json:"DrugOrgId,omitnil,omitempty" name:"DrugOrgId"`

	// 医院药品通用名
	DrugName *string `json:"DrugName,omitnil,omitempty" name:"DrugName"`

	// 医院药品商品名
	DrugCommodityName *string `json:"DrugCommodityName,omitnil,omitempty" name:"DrugCommodityName"`

	// 医院药品规格
	Specifications *string `json:"Specifications,omitnil,omitempty" name:"Specifications"`

	// 医院药品批准文号
	ApprovalNumber *string `json:"ApprovalNumber,omitnil,omitempty" name:"ApprovalNumber"`

	// 生产厂商
	Manufacturer *string `json:"Manufacturer,omitnil,omitempty" name:"Manufacturer"`

	// 剂型
	DosageForm *string `json:"DosageForm,omitnil,omitempty" name:"DosageForm"`

	// 使用状态 0:启用 1:停用
	Unuse *int64 `json:"Unuse,omitnil,omitempty" name:"Unuse"`

	// 剂型编码
	DosageFormCode *string `json:"DosageFormCode,omitnil,omitempty" name:"DosageFormCode"`

	// 抗菌药DDD值
	DefinedDailyDose *string `json:"DefinedDailyDose,omitnil,omitempty" name:"DefinedDailyDose"`

	// 药品单价
	Amount *string `json:"Amount,omitnil,omitempty" name:"Amount"`

	// 国家医保编码
	YbCode *string `json:"YbCode,omitnil,omitempty" name:"YbCode"`

	// 药品本位码
	DrugBasicCode *string `json:"DrugBasicCode,omitnil,omitempty" name:"DrugBasicCode"`

	// 药品属性
	PropertyInfo *DurgPropertyInfo `json:"PropertyInfo,omitnil,omitempty" name:"PropertyInfo"`
}

type DrugList struct {
	// 药品ID
	DrugId *string `json:"DrugId,omitnil,omitempty" name:"DrugId"`

	// 药品名称
	DrugName *string `json:"DrugName,omitnil,omitempty" name:"DrugName"`

	// 文档地址
	DocUrl *string `json:"DocUrl,omitnil,omitempty" name:"DocUrl"`

	// 是否找到
	NotFound *bool `json:"NotFound,omitnil,omitempty" name:"NotFound"`
}

type DrugUsage struct {
	// 药品ID
	DrugId *string `json:"DrugId,omitnil,omitempty" name:"DrugId"`

	// 药品名称
	DrugName *string `json:"DrugName,omitnil,omitempty" name:"DrugName"`

	// 日服药频次
	TimePerDay *string `json:"TimePerDay,omitnil,omitempty" name:"TimePerDay"`

	// 给药途径
	Usage *string `json:"Usage,omitnil,omitempty" name:"Usage"`

	// 处方ID，药品不同分组是传不同的处方ID
	PrescriptionId *string `json:"PrescriptionId,omitnil,omitempty" name:"PrescriptionId"`

	// 每次剂量
	DosagePerTime *string `json:"DosagePerTime,omitnil,omitempty" name:"DosagePerTime"`

	// 每次剂量单位
	DosagePerTimeUnit *string `json:"DosagePerTimeUnit,omitnil,omitempty" name:"DosagePerTimeUnit"`

	// 单次服药时间
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 给药周期
	Cycle *string `json:"Cycle,omitnil,omitempty" name:"Cycle"`

	// 单日剂量
	DosagePerDay *string `json:"DosagePerDay,omitnil,omitempty" name:"DosagePerDay"`

	// 疗程
	Course *string `json:"Course,omitnil,omitempty" name:"Course"`

	// 给药速度
	Speed *string `json:"Speed,omitnil,omitempty" name:"Speed"`

	// 处方生效时间戳，住院医嘱必须传(caseType =1)
	BeginTime *int64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 处方失效时间戳，住院医嘱必须传(caseType =1)
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 开具数量
	Package *string `json:"Package,omitnil,omitempty" name:"Package"`

	// 开具数量单位
	PackageUnit *string `json:"PackageUnit,omitnil,omitempty" name:"PackageUnit"`

	// 相同标志液体间进行配伍禁忌审核，不同标志间液体不进行配伍禁忌审核
	GroupInj *string `json:"GroupInj,omitnil,omitempty" name:"GroupInj"`

	// 处方金额
	PrescriptionCharge *string `json:"PrescriptionCharge,omitnil,omitempty" name:"PrescriptionCharge"`

	// 用药天数
	MedicationDays *string `json:"MedicationDays,omitnil,omitempty" name:"MedicationDays"`
}

type DurgPropertyInfo struct {
	// 药品类型 1:西药,2:中成药,3:中药,4:化学药品,5:生物制药
	DrugType *int64 `json:"DrugType,omitnil,omitempty" name:"DrugType"`

	// 抗菌药分类 1:抗真菌药物, 2:抗细菌药物, 3:抗结核药物, 4:其他抗菌药, 0:普通药品
	AntibacterialType *int64 `json:"AntibacterialType,omitnil,omitempty" name:"AntibacterialType"`

	// 抗菌药级别 1:非限制级, 2:限制级, 3:特殊使用级 
	AntibacterialClass *int64 `json:"AntibacterialClass,omitnil,omitempty" name:"AntibacterialClass"`

	// 特殊药品类型 1:毒性药品, 2:麻醉药品, 3:放射药品, 4:精神一类药品, 5:精神二类药品, 6:其他特管药品， 7:贵重药品
	SpeciallyDrugType *int64 `json:"SpeciallyDrugType,omitnil,omitempty" name:"SpeciallyDrugType"`

	// 是否为基本药物 1:是, 2:否, 0:未知
	IsBasicDrug *int64 `json:"IsBasicDrug,omitnil,omitempty" name:"IsBasicDrug"`

	// 社保药品 1:甲类药品, 2:乙类药品, 3:双跨药品, 4:自费药品, 0:未知
	ChargeType *int64 `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`
}

type EmrDiagnosises struct {
	// 诊断名称
	DiagnosisName *string `json:"DiagnosisName,omitnil,omitempty" name:"DiagnosisName"`

	// ICD代码
	IcdCode *string `json:"IcdCode,omitnil,omitempty" name:"IcdCode"`
}

type EmrQuality struct {
	// 缺失体格检查项目
	MissPhysicalExamination []*string `json:"MissPhysicalExamination,omitnil,omitempty" name:"MissPhysicalExamination"`
}

type GetDrugIndicationsReqData struct {
	// 查询药品列表
	Drugs []*IndicationsDrug `json:"Drugs,omitnil,omitempty" name:"Drugs"`
}

// Predefined struct for user
type GetDrugIndicationsRequestParams struct {
	// 请求头
	Header *CommonHeader `json:"Header,omitnil,omitempty" name:"Header"`

	// 获取适应症请求
	Data *GetDrugIndicationsReqData `json:"Data,omitnil,omitempty" name:"Data"`
}

type GetDrugIndicationsRequest struct {
	*tchttp.BaseRequest
	
	// 请求头
	Header *CommonHeader `json:"Header,omitnil,omitempty" name:"Header"`

	// 获取适应症请求
	Data *GetDrugIndicationsReqData `json:"Data,omitnil,omitempty" name:"Data"`
}

func (r *GetDrugIndicationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDrugIndicationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Header")
	delete(f, "Data")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDrugIndicationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetDrugIndicationsResp struct {
	// 适应症
	Indications []*string `json:"Indications,omitnil,omitempty" name:"Indications"`

	// 药品说明
	DocInfos []*DocInfo `json:"DocInfos,omitnil,omitempty" name:"DocInfos"`
}

// Predefined struct for user
type GetDrugIndicationsResponseParams struct {
	// 错误码
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 错误信息
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 适应症响应
	Data *GetDrugIndicationsResp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetDrugIndicationsResponse struct {
	*tchttp.BaseResponse
	Response *GetDrugIndicationsResponseParams `json:"Response"`
}

func (r *GetDrugIndicationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDrugIndicationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HealthPrescriptions struct {
	// 标题
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 健康处方链接
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 关键信息
	KeyInformation *string `json:"KeyInformation,omitnil,omitempty" name:"KeyInformation"`
}

type IndicationsDrug struct {
	// 药品名称
	DrugName *string `json:"DrugName,omitnil,omitempty" name:"DrugName"`

	// 规格
	Specifications *string `json:"Specifications,omitnil,omitempty" name:"Specifications"`

	// 批准文号
	ApprovalNumber *string `json:"ApprovalNumber,omitnil,omitempty" name:"ApprovalNumber"`

	// 生产厂家
	Manufacturer *string `json:"Manufacturer,omitnil,omitempty" name:"Manufacturer"`

	// 药品ID
	DrugId *string `json:"DrugId,omitnil,omitempty" name:"DrugId"`

	// 商品名
	TradeName *string `json:"TradeName,omitnil,omitempty" name:"TradeName"`

	// 类型 默认0 0-西药 2-中药
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`
}

type LoginData struct {
	// 医生ID
	DoctorId *string `json:"DoctorId,omitnil,omitempty" name:"DoctorId"`

	// 医生名称
	DoctorName *string `json:"DoctorName,omitnil,omitempty" name:"DoctorName"`

	// 医生职级 主治医师、副主任医师、主任医师
	DoctorLevel *string `json:"DoctorLevel,omitnil,omitempty" name:"DoctorLevel"`

	// 医生科室 当前登录科室
	DoctorDepartment *string `json:"DoctorDepartment,omitnil,omitempty" name:"DoctorDepartment"`

	// 科室ID
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`
}

type LoginDataResp struct {
	// token
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// 过期时间
	ExpiresIn *int64 `json:"ExpiresIn,omitnil,omitempty" name:"ExpiresIn"`

	// 服务端时间戳，时间戳校验失败时返回
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`
}

type LoginHeader struct {
	// 机构ID
	HospitalId *string `json:"HospitalId,omitnil,omitempty" name:"HospitalId"`

	// 合作方ID
	PartnerId *string `json:"PartnerId,omitnil,omitempty" name:"PartnerId"`

	// 加密时间戳毫秒
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 签名数据
	Signature *string `json:"Signature,omitnil,omitempty" name:"Signature"`

	// 平台ID，平台版登录时传入
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

// Predefined struct for user
type LoginHisToolRequestParams struct {
	// 请求头
	Header *LoginHeader `json:"Header,omitnil,omitempty" name:"Header"`

	// 请求体
	Data *LoginData `json:"Data,omitnil,omitempty" name:"Data"`
}

type LoginHisToolRequest struct {
	*tchttp.BaseRequest
	
	// 请求头
	Header *LoginHeader `json:"Header,omitnil,omitempty" name:"Header"`

	// 请求体
	Data *LoginData `json:"Data,omitnil,omitempty" name:"Data"`
}

func (r *LoginHisToolRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LoginHisToolRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Header")
	delete(f, "Data")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "LoginHisToolRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type LoginHisToolResponseParams struct {
	// 错误码
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 错误信息
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 登录返回token信息
	Data *LoginDataResp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type LoginHisToolResponse struct {
	*tchttp.BaseResponse
	Response *LoginHisToolResponseParams `json:"Response"`
}

func (r *LoginHisToolResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LoginHisToolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoginOutData struct {
	// 登录获取的token
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`
}

type LoginOutHeader struct {
	// 合作方ID
	PartnerId *string `json:"PartnerId,omitnil,omitempty" name:"PartnerId"`

	// 时间戳毫秒数
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 签名值
	Signature *string `json:"Signature,omitnil,omitempty" name:"Signature"`

	// 医院ID 单院版传这个
	HospitalId *string `json:"HospitalId,omitnil,omitempty" name:"HospitalId"`

	// 平台ID 平台版传这个
	PlatformId *string `json:"PlatformId,omitnil,omitempty" name:"PlatformId"`
}

// Predefined struct for user
type LoginOutHisToolRequestParams struct {
	// 登出请求header
	Header *LoginOutHeader `json:"Header,omitnil,omitempty" name:"Header"`

	// 登出请求数据
	Data *LoginOutData `json:"Data,omitnil,omitempty" name:"Data"`
}

type LoginOutHisToolRequest struct {
	*tchttp.BaseRequest
	
	// 登出请求header
	Header *LoginOutHeader `json:"Header,omitnil,omitempty" name:"Header"`

	// 登出请求数据
	Data *LoginOutData `json:"Data,omitnil,omitempty" name:"Data"`
}

func (r *LoginOutHisToolRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LoginOutHisToolRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Header")
	delete(f, "Data")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "LoginOutHisToolRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type LoginOutHisToolResponseParams struct {
	// 错误码
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 错误信息
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 响应数据
	Data *LoginOutResponseData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type LoginOutHisToolResponse struct {
	*tchttp.BaseResponse
	Response *LoginOutHisToolResponseParams `json:"Response"`
}

func (r *LoginOutHisToolResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LoginOutHisToolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoginOutResponseData struct {
	// 服务器时间戳毫秒
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`
}

type OperateResp struct {
	// 操作结果
	Dummy *bool `json:"Dummy,omitnil,omitempty" name:"Dummy"`
}

type PatientBaseInfo struct {
	// 性别
	Sex *string `json:"Sex,omitnil,omitempty" name:"Sex"`

	// 身高 单位cm
	Height *string `json:"Height,omitnil,omitempty" name:"Height"`

	// 体重 单位kg
	Weight *string `json:"Weight,omitnil,omitempty" name:"Weight"`

	// 患者ID
	PatientId *string `json:"PatientId,omitnil,omitempty" name:"PatientId"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 年龄
	Age *string `json:"Age,omitnil,omitempty" name:"Age"`

	// 出生地
	BirthPlace *string `json:"BirthPlace,omitnil,omitempty" name:"BirthPlace"`

	// 居住地
	LivePlace *string `json:"LivePlace,omitnil,omitempty" name:"LivePlace"`

	// 出生日期和年龄必须传一个
	BirthDay *string `json:"BirthDay,omitnil,omitempty" name:"BirthDay"`
}

type PatientFamilyHistory struct {
	// 家族病史 如家族病史不能分开，可传入此字段
	FamilyDiseaseHistory *string `json:"FamilyDiseaseHistory,omitnil,omitempty" name:"FamilyDiseaseHistory"`

	// 关系
	Relation *string `json:"Relation,omitnil,omitempty" name:"Relation"`

	// 当前情况
	CurrentSituation *string `json:"CurrentSituation,omitnil,omitempty" name:"CurrentSituation"`
}

type PatientHistory struct {
	// 病史
	DiseaseHistory *string `json:"DiseaseHistory,omitnil,omitempty" name:"DiseaseHistory"`

	// 治疗史
	TreatmentHistory *string `json:"TreatmentHistory,omitnil,omitempty" name:"TreatmentHistory"`
}

type PhysicalExam struct {
	// 脉搏，单位：次/分
	Pulse *string `json:"Pulse,omitnil,omitempty" name:"Pulse"`

	// 呼吸，单位：次/分
	Breathe *string `json:"Breathe,omitnil,omitempty" name:"Breathe"`

	// 体重，单位KG
	Weight *string `json:"Weight,omitnil,omitempty" name:"Weight"`

	// 体温，单位：℃ 
	BodyTemperature *string `json:"BodyTemperature,omitnil,omitempty" name:"BodyTemperature"`

	// 舒张压， 单位：mmHg
	DiastolicPressure *string `json:"DiastolicPressure,omitnil,omitempty" name:"DiastolicPressure"`

	// 收缩压， 单位：mmHg
	SystolicPressure *string `json:"SystolicPressure,omitnil,omitempty" name:"SystolicPressure"`
}

type RationalDrugInfo struct {
	// 是否有风险
	Hit *bool `json:"Hit,omitnil,omitempty" name:"Hit"`

	// 药品用量风险
	DrugUsages []*RiskInfo `json:"DrugUsages,omitnil,omitempty" name:"DrugUsages"`

	// 重复用药风险
	DrugRepeats []*RiskInfo `json:"DrugRepeats,omitnil,omitempty" name:"DrugRepeats"`

	// 用药途径风险
	DrugRoutes []*RiskInfo `json:"DrugRoutes,omitnil,omitempty" name:"DrugRoutes"`

	// 特殊人群风险
	SpecialPopulations []*RiskInfo `json:"SpecialPopulations,omitnil,omitempty" name:"SpecialPopulations"`

	// 禁忌症风险
	DrugTaboos []*RiskInfo `json:"DrugTaboos,omitnil,omitempty" name:"DrugTaboos"`

	// 相互作用风险
	DrugInteractions []*RiskInfo `json:"DrugInteractions,omitnil,omitempty" name:"DrugInteractions"`

	// 配伍禁忌风险
	DrugIncompatibility []*RiskInfo `json:"DrugIncompatibility,omitnil,omitempty" name:"DrugIncompatibility"`

	// 过敏风险
	DrugAllergys []*RiskInfo `json:"DrugAllergys,omitnil,omitempty" name:"DrugAllergys"`

	// 适应症风险
	DrugIndications []*RiskInfo `json:"DrugIndications,omitnil,omitempty" name:"DrugIndications"`

	// 异常提醒
	Abnormals []*Abnormals `json:"Abnormals,omitnil,omitempty" name:"Abnormals"`

	// 药品列表
	DrugList []*DrugList `json:"DrugList,omitnil,omitempty" name:"DrugList"`
}

type RecommendedUsage struct {
	// 给药途径
	UsageRoute *string `json:"UsageRoute,omitnil,omitempty" name:"UsageRoute"`

	// 给药频率 格式为“最小频次,最大频次,频次单位,频次周期”，如"1,2,次,2"， 表示2天内最少给药1次，最大给药2次。
	Frequency *string `json:"Frequency,omitnil,omitempty" name:"Frequency"`

	// 给药剂量 格式为"最小剂量,最大剂量,剂量单位"，如"10,10,ml"，<br>表示每次最大给药量为10ml，最小给药量为10ml。
	SingleDose *string `json:"SingleDose,omitnil,omitempty" name:"SingleDose"`
}

type RecordQuality struct {
	// 病历是否有问题
	Hit *bool `json:"Hit,omitnil,omitempty" name:"Hit"`

	// 完整性问题
	Completeness *string `json:"Completeness,omitnil,omitempty" name:"Completeness"`

	// 及时性问题
	Timeliness *string `json:"Timeliness,omitnil,omitempty" name:"Timeliness"`

	// 逻辑性问题
	Logical *string `json:"Logical,omitnil,omitempty" name:"Logical"`
}

type ReferralInfo struct {
	// 命中
	Hit *bool `json:"Hit,omitnil,omitempty" name:"Hit"`

	// 提示
	Tips *string `json:"Tips,omitnil,omitempty" name:"Tips"`
}

type RequestCase struct {
	// 处方类型 0:门诊处方；1:住院医嘱；2:急诊处方 
	CaseType *int64 `json:"CaseType,omitnil,omitempty" name:"CaseType"`

	// 主诉
	ChiefComplaint *string `json:"ChiefComplaint,omitnil,omitempty" name:"ChiefComplaint"`

	// 科室
	Department *string `json:"Department,omitnil,omitempty" name:"Department"`

	// 病历文书ID
	// 医生每次书写病历文书的ID，文书内容包含主诉，病史，当前诊断等内容<br />门诊场景：门诊病历文书（带有主诉、病史、诊断及药品的）只有一份，这个编号只有一个。<br/>住院场景：假设住院3天，医生每天都会写一份病历文书（带有主诉、病史、诊断及医嘱药品的文书），那么有对应三个病历文书编号，每次调用预测接口都要传入不同的病历文书编号。注意：如两次调用预测接口，传相同的caseid，则在药师端管理平台的上一次审方记录中的诊断会被本次接口传入的诊断更新。
	CaseId *string `json:"CaseId,omitnil,omitempty" name:"CaseId"`

	// 病历更新时间
	CaseTime *string `json:"CaseTime,omitnil,omitempty" name:"CaseTime"`

	// 就诊ID
	// 门诊处方传门诊号，住院医嘱传住院号；备注：门诊场景：用户挂一次号，看一个医生，这时候会有一个代表变成就诊的编号，下一次挂号就诊，这个编号会变。<br/>住院场景：用户本次办理入院，会有一个住院编号，仅代表本次住院，如果下次再住院，这个编号会变。
	VisitId *string `json:"VisitId,omitnil,omitempty" name:"VisitId"`

	// 患者信息
	PatientBaseinfo *PatientBaseInfo `json:"PatientBaseinfo,omitnil,omitempty" name:"PatientBaseinfo"`

	// 医生信息
	DoctorInfo *DoctorInfo `json:"DoctorInfo,omitnil,omitempty" name:"DoctorInfo"`

	// 现病史
	PresentIllness *string `json:"PresentIllness,omitnil,omitempty" name:"PresentIllness"`

	// 患者其他信息，包含过敏史等
	PatientOther *string `json:"PatientOther,omitnil,omitempty" name:"PatientOther"`

	// 患者过往病史
	PatientHistory *PatientHistory `json:"PatientHistory,omitnil,omitempty" name:"PatientHistory"`

	// 患者家族病史
	PatientFamilyHistory *PatientFamilyHistory `json:"PatientFamilyHistory,omitnil,omitempty" name:"PatientFamilyHistory"`

	// 体格检查
	PhysicalExam *PhysicalExam `json:"PhysicalExam,omitnil,omitempty" name:"PhysicalExam"`

	// 诊断列表，第一个为首要诊断
	EmrDiagnosises []*EmrDiagnosises `json:"EmrDiagnosises,omitnil,omitempty" name:"EmrDiagnosises"`

	// 处方列表
	Prescriptions []*DrugUsage `json:"Prescriptions,omitnil,omitempty" name:"Prescriptions"`
}

type RiskInfo struct {
	// 药品ID
	DrugId *string `json:"DrugId,omitnil,omitempty" name:"DrugId"`

	// 药品名称
	DrugName *string `json:"DrugName,omitnil,omitempty" name:"DrugName"`

	// 风险等级：低级风险、中级风险、高级风险
	RiskLevel *string `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 风险提示
	RiskTips *string `json:"RiskTips,omitnil,omitempty" name:"RiskTips"`

	// FDA分级
	FdaLevel *string `json:"FdaLevel,omitnil,omitempty" name:"FdaLevel"`

	// 关联药品名称
	RelatedDrugName *string `json:"RelatedDrugName,omitnil,omitempty" name:"RelatedDrugName"`

	// 关联处方ID
	RelatedPrescriptionId *string `json:"RelatedPrescriptionId,omitnil,omitempty" name:"RelatedPrescriptionId"`
}

type SmartDrugInfoData struct {
	// 药品名称
	DrugName *string `json:"DrugName,omitnil,omitempty" name:"DrugName"`

	// 规格
	Specifications *string `json:"Specifications,omitnil,omitempty" name:"Specifications"`

	// 批准文号
	ApprovalNumber *string `json:"ApprovalNumber,omitnil,omitempty" name:"ApprovalNumber"`

	// 生产厂家
	Manufacturer *string `json:"Manufacturer,omitnil,omitempty" name:"Manufacturer"`

	// 药品ID
	DrugId *string `json:"DrugId,omitnil,omitempty" name:"DrugId"`

	// 诊断
	Diagnosis *string `json:"Diagnosis,omitnil,omitempty" name:"Diagnosis"`

	// 年龄
	Age *float64 `json:"Age,omitnil,omitempty" name:"Age"`
}

// Predefined struct for user
type SmartDrugInfoRequestParams struct {
	// 请求头
	Header *CommonHeader `json:"Header,omitnil,omitempty" name:"Header"`

	// 药品信息
	Data *SmartDrugInfoData `json:"Data,omitnil,omitempty" name:"Data"`
}

type SmartDrugInfoRequest struct {
	*tchttp.BaseRequest
	
	// 请求头
	Header *CommonHeader `json:"Header,omitnil,omitempty" name:"Header"`

	// 药品信息
	Data *SmartDrugInfoData `json:"Data,omitnil,omitempty" name:"Data"`
}

func (r *SmartDrugInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SmartDrugInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Header")
	delete(f, "Data")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SmartDrugInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SmartDrugInfoResp struct {
	// 药品ID
	DrugId *string `json:"DrugId,omitnil,omitempty" name:"DrugId"`

	// 序列ID
	SequenceId *int64 `json:"SequenceId,omitnil,omitempty" name:"SequenceId"`

	// 药品哈希ID
	DrugHashId *string `json:"DrugHashId,omitnil,omitempty" name:"DrugHashId"`

	// 图片URL
	ImgUrl *string `json:"ImgUrl,omitnil,omitempty" name:"ImgUrl"`

	// 药品名称
	DrugName *string `json:"DrugName,omitnil,omitempty" name:"DrugName"`

	// 商品名
	TradeName *string `json:"TradeName,omitnil,omitempty" name:"TradeName"`

	// 英文名称
	EnglishName *string `json:"EnglishName,omitnil,omitempty" name:"EnglishName"`

	// 英文商品名
	EnglishTradeName *string `json:"EnglishTradeName,omitnil,omitempty" name:"EnglishTradeName"`

	// 拼音
	Pinyin *string `json:"Pinyin,omitnil,omitempty" name:"Pinyin"`

	// 其他名称
	OtherNames *string `json:"OtherNames,omitnil,omitempty" name:"OtherNames"`

	// 化学名称
	ChemicalName *string `json:"ChemicalName,omitnil,omitempty" name:"ChemicalName"`

	// 英文化学名称
	EnglishChemicalName *string `json:"EnglishChemicalName,omitnil,omitempty" name:"EnglishChemicalName"`

	// 批准文号
	ApprovalNumber *string `json:"ApprovalNumber,omitnil,omitempty" name:"ApprovalNumber"`

	// 药品属性标签 多个标签时 | 分隔，如抗菌药|抗生素|贵重药品
	Property *string `json:"Property,omitnil,omitempty" name:"Property"`

	// 药品成分
	Ingredients *string `json:"Ingredients,omitnil,omitempty" name:"Ingredients"`

	// 药品性状
	PhenotypicTrait *string `json:"PhenotypicTrait,omitnil,omitempty" name:"PhenotypicTrait"`

	// 适应症
	Indications *string `json:"Indications,omitnil,omitempty" name:"Indications"`

	// 规格
	Specifications *string `json:"Specifications,omitnil,omitempty" name:"Specifications"`

	// 用法用量
	UsageAndDosage *string `json:"UsageAndDosage,omitnil,omitempty" name:"UsageAndDosage"`

	// 推荐用法
	RecommendedUsage *RecommendedUsage `json:"RecommendedUsage,omitnil,omitempty" name:"RecommendedUsage"`

	// 不良反应
	AdverseReaction *string `json:"AdverseReaction,omitnil,omitempty" name:"AdverseReaction"`

	// 禁忌
	Contraindication *string `json:"Contraindication,omitnil,omitempty" name:"Contraindication"`

	// 注意事项
	Attentions *string `json:"Attentions,omitnil,omitempty" name:"Attentions"`

	// 药物过量
	Overdose *string `json:"Overdose,omitnil,omitempty" name:"Overdose"`

	// 孕妇及哺乳期妇女用药
	PregnantAndLactatingWomen *string `json:"PregnantAndLactatingWomen,omitnil,omitempty" name:"PregnantAndLactatingWomen"`

	// 老年患者用药
	ElderlyPatients *string `json:"ElderlyPatients,omitnil,omitempty" name:"ElderlyPatients"`

	// 儿童用药
	PediatricDrugs *string `json:"PediatricDrugs,omitnil,omitempty" name:"PediatricDrugs"`

	// 药物相互作用
	Interactions *string `json:"Interactions,omitnil,omitempty" name:"Interactions"`

	// 临床研究
	ClinicalResearch *string `json:"ClinicalResearch,omitnil,omitempty" name:"ClinicalResearch"`

	// 药理毒理
	PharmacologyToxicology *string `json:"PharmacologyToxicology,omitnil,omitempty" name:"PharmacologyToxicology"`

	// 药代动力学
	Pharmacokinetics *string `json:"Pharmacokinetics,omitnil,omitempty" name:"Pharmacokinetics"`

	// 警告
	Warning *string `json:"Warning,omitnil,omitempty" name:"Warning"`

	// 有效期
	ExpireDate *string `json:"ExpireDate,omitnil,omitempty" name:"ExpireDate"`

	// 贮藏
	Storage *string `json:"Storage,omitnil,omitempty" name:"Storage"`

	// 包装
	Pack *string `json:"Pack,omitnil,omitempty" name:"Pack"`

	// 生产企业
	Manufacturer *string `json:"Manufacturer,omitnil,omitempty" name:"Manufacturer"`

	// 生产企业地址
	ManufacturerAddress *string `json:"ManufacturerAddress,omitnil,omitempty" name:"ManufacturerAddress"`

	// 生产企业电话
	ManufacturerPhone *string `json:"ManufacturerPhone,omitnil,omitempty" name:"ManufacturerPhone"`

	// 生产企业邮箱
	ManufacturerEmail *string `json:"ManufacturerEmail,omitnil,omitempty" name:"ManufacturerEmail"`

	// 生产企业网站
	ManufacturerWebsite *string `json:"ManufacturerWebsite,omitnil,omitempty" name:"ManufacturerWebsite"`

	// 说明书制定和修订时间
	DocRevisionTime *string `json:"DocRevisionTime,omitnil,omitempty" name:"DocRevisionTime"`

	// 参考文献
	References *string `json:"References,omitnil,omitempty" name:"References"`

	// 剂型
	DrugDosageForm *string `json:"DrugDosageForm,omitnil,omitempty" name:"DrugDosageForm"`

	// 给药途径
	DrugRoute *string `json:"DrugRoute,omitnil,omitempty" name:"DrugRoute"`

	// 药品本位码
	DrugBasicCode *string `json:"DrugBasicCode,omitnil,omitempty" name:"DrugBasicCode"`

	// OCT标签
	OctTag *string `json:"OctTag,omitnil,omitempty" name:"OctTag"`
}

// Predefined struct for user
type SmartDrugInfoResponseParams struct {
	// 错误码
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 错误信息
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 智能用药信息
	Data *SmartDrugInfoResp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SmartDrugInfoResponse struct {
	*tchttp.BaseResponse
	Response *SmartDrugInfoResponseParams `json:"Response"`
}

func (r *SmartDrugInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SmartDrugInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SmartPredictReqData struct {
	// 病历和处方信息
	RequestCase *RequestCase `json:"RequestCase,omitnil,omitempty" name:"RequestCase"`

	// 0--默认值，同时返回**疾病预测**和**用药审查**结果<br>1--仅返回**疾病预测**结果<br>2--仅返回**用药审查**结果<br>已同时激活两个模块时，可按需使用 
	RequestType *int64 `json:"RequestType,omitnil,omitempty" name:"RequestType"`
}

// Predefined struct for user
type SmartPredictRequestParams struct {
	// 请求头
	Header *CommonHeader `json:"Header,omitnil,omitempty" name:"Header"`

	// 请求体
	Data *SmartPredictReqData `json:"Data,omitnil,omitempty" name:"Data"`
}

type SmartPredictRequest struct {
	*tchttp.BaseRequest
	
	// 请求头
	Header *CommonHeader `json:"Header,omitnil,omitempty" name:"Header"`

	// 请求体
	Data *SmartPredictReqData `json:"Data,omitnil,omitempty" name:"Data"`
}

func (r *SmartPredictRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SmartPredictRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Header")
	delete(f, "Data")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SmartPredictRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SmartPredictRespData struct {
	// 诊断辅助内容
	DiagnosisInfo *DiagnosisInfo `json:"DiagnosisInfo,omitnil,omitempty" name:"DiagnosisInfo"`

	// 用药风险信息
	RationalDrugInfo *RationalDrugInfo `json:"RationalDrugInfo,omitnil,omitempty" name:"RationalDrugInfo"`
}

// Predefined struct for user
type SmartPredictResponseParams struct {
	// 错误码
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 错误信息
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 智能预测内容
	Data *SmartPredictRespData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SmartPredictResponse struct {
	*tchttp.BaseResponse
	Response *SmartPredictResponseParams `json:"Response"`
}

func (r *SmartPredictResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SmartPredictResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SuspectedDiagnosis struct {
	// 疾病名称
	DiseaseName *string `json:"DiseaseName,omitnil,omitempty" name:"DiseaseName"`

	// ICD代码
	IcdCode *string `json:"IcdCode,omitnil,omitempty" name:"IcdCode"`

	// 症状
	Symptom *string `json:"Symptom,omitnil,omitempty" name:"Symptom"`

	// 体征
	Physi *string `json:"Physi,omitnil,omitempty" name:"Physi"`

	// 检查
	Inspection *string `json:"Inspection,omitnil,omitempty" name:"Inspection"`

	// 疾病指南信息
	DiseaseGuideInfo *string `json:"DiseaseGuideInfo,omitnil,omitempty" name:"DiseaseGuideInfo"`

	// 置信度
	Probability *float64 `json:"Probability,omitnil,omitempty" name:"Probability"`
}

type SyncDepartmentData struct {
	// 操作类型 1:获取科室列表  2:同步科室信息（增、改） 3:删除科室
	Cmd *int64 `json:"Cmd,omitnil,omitempty" name:"Cmd"`

	// 科室列表
	List []*Department `json:"List,omitnil,omitempty" name:"List"`
}

// Predefined struct for user
type SyncDepartmentRequestParams struct {
	// 请求头
	Header *CommonHeader `json:"Header,omitnil,omitempty" name:"Header"`

	// 同步数据
	Data *SyncDepartmentData `json:"Data,omitnil,omitempty" name:"Data"`
}

type SyncDepartmentRequest struct {
	*tchttp.BaseRequest
	
	// 请求头
	Header *CommonHeader `json:"Header,omitnil,omitempty" name:"Header"`

	// 同步数据
	Data *SyncDepartmentData `json:"Data,omitnil,omitempty" name:"Data"`
}

func (r *SyncDepartmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncDepartmentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Header")
	delete(f, "Data")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SyncDepartmentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SyncDepartmentRespData struct {
	// 科室列表
	List []*Department `json:"List,omitnil,omitempty" name:"List"`
}

// Predefined struct for user
type SyncDepartmentResponseParams struct {
	// 错误码
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 错误信息
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 响应数据
	Data *SyncDepartmentRespData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SyncDepartmentResponse struct {
	*tchttp.BaseResponse
	Response *SyncDepartmentResponseParams `json:"Response"`
}

func (r *SyncDepartmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncDepartmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncDictData struct {
	// 医院ID
	HospitalId *string `json:"HospitalId,omitnil,omitempty" name:"HospitalId"`

	// 字典类型 1-给药频次 2-给药途径 3-科室 4-诊断
	DictType *int64 `json:"DictType,omitnil,omitempty" name:"DictType"`

	// 字典数据
	Dicts []*Dict `json:"Dicts,omitnil,omitempty" name:"Dicts"`
}

// Predefined struct for user
type SyncStandardDictRequestParams struct {
	// 请求头
	Header *CommonHeader `json:"Header,omitnil,omitempty" name:"Header"`

	// 字典数据
	Data *SyncDictData `json:"Data,omitnil,omitempty" name:"Data"`
}

type SyncStandardDictRequest struct {
	*tchttp.BaseRequest
	
	// 请求头
	Header *CommonHeader `json:"Header,omitnil,omitempty" name:"Header"`

	// 字典数据
	Data *SyncDictData `json:"Data,omitnil,omitempty" name:"Data"`
}

func (r *SyncStandardDictRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncStandardDictRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Header")
	delete(f, "Data")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SyncStandardDictRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SyncStandardDictResponseParams struct {
	// 错误码
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 错误信息
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SyncStandardDictResponse struct {
	*tchttp.BaseResponse
	Response *SyncStandardDictResponseParams `json:"Response"`
}

func (r *SyncStandardDictResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncStandardDictResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TreatmentGuide struct {
	// 医生诊断
	DoctorDiagnosis *string `json:"DoctorDiagnosis,omitnil,omitempty" name:"DoctorDiagnosis"`

	// 疾病名称
	DiseaseName *string `json:"DiseaseName,omitnil,omitempty" name:"DiseaseName"`

	// 治疗详情链接
	TreatDetailUrl *string `json:"TreatDetailUrl,omitnil,omitempty" name:"TreatDetailUrl"`

	// 治疗方案
	TreatPlan *string `json:"TreatPlan,omitnil,omitempty" name:"TreatPlan"`

	// 治疗原则
	TreatPrinciple *string `json:"TreatPrinciple,omitnil,omitempty" name:"TreatPrinciple"`
}

type UploadDrugData struct {
	// 药品列表
	Drugs []*Drug `json:"Drugs,omitnil,omitempty" name:"Drugs"`
}

// Predefined struct for user
type UploadDrugsRequestParams struct {
	// 请求头数据
	Header *CommonHeader `json:"Header,omitnil,omitempty" name:"Header"`

	// 药品上传数据
	Data *UploadDrugData `json:"Data,omitnil,omitempty" name:"Data"`
}

type UploadDrugsRequest struct {
	*tchttp.BaseRequest
	
	// 请求头数据
	Header *CommonHeader `json:"Header,omitnil,omitempty" name:"Header"`

	// 药品上传数据
	Data *UploadDrugData `json:"Data,omitnil,omitempty" name:"Data"`
}

func (r *UploadDrugsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadDrugsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Header")
	delete(f, "Data")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadDrugsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadDrugsResponseParams struct {
	// 错误码
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 错误信息
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 操作信息
	Data *OperateResp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UploadDrugsResponse struct {
	*tchttp.BaseResponse
	Response *UploadDrugsResponseParams `json:"Response"`
}

func (r *UploadDrugsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadDrugsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VitalSignsInfo struct {
	// 是否包含风险
	Hit *bool `json:"Hit,omitnil,omitempty" name:"Hit"`

	// 提示
	Tips *string `json:"Tips,omitnil,omitempty" name:"Tips"`
}