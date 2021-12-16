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

package v20200910

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Advice struct {

	// 文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitempty" name:"Text"`
}

type AspectRatio struct {

	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitempty" name:"Index"`

	// 数值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Number *string `json:"Number,omitempty" name:"Number"`

	// 关系
	// 注意：此字段可能返回 null，表示取不到有效值。
	Relation *string `json:"Relation,omitempty" name:"Relation"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type BiopsyPart struct {

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`
}

type BlockInfo struct {

	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitempty" name:"Index"`

	// 阳性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Positive *string `json:"Positive,omitempty" name:"Positive"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`
}

type Check struct {

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *Desc `json:"Desc,omitempty" name:"Desc"`

	// 结论
	// 注意：此字段可能返回 null，表示取不到有效值。
	Summary *Summary `json:"Summary,omitempty" name:"Summary"`
}

type Desc struct {

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitempty" name:"Text"`

	// 器官
	// 注意：此字段可能返回 null，表示取不到有效值。
	Organ []*Organ `json:"Organ,omitempty" name:"Organ"`

	// 结节
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tuber []*TuberInfo `json:"Tuber,omitempty" name:"Tuber"`
}

type DiagCert struct {

	// 建议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Advice *Advice `json:"Advice,omitempty" name:"Advice"`

	// 诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	Diagnosis []*DiagCertItem `json:"Diagnosis,omitempty" name:"Diagnosis"`
}

type DiagCertItem struct {

	// 文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitempty" name:"Text"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value []*string `json:"Value,omitempty" name:"Value"`
}

type DischargeDiagnosis struct {

	// 表格位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableIndex *int64 `json:"TableIndex,omitempty" name:"TableIndex"`

	// 出院诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutDiagnosis *BlockInfo `json:"OutDiagnosis,omitempty" name:"OutDiagnosis"`

	// 疾病编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiseaseCode *BlockInfo `json:"DiseaseCode,omitempty" name:"DiseaseCode"`

	// 入院情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	InStatus *BlockInfo `json:"InStatus,omitempty" name:"InStatus"`

	// 出院情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutStatus *BlockInfo `json:"OutStatus,omitempty" name:"OutStatus"`
}

type DiseaseMedicalHistory struct {

	// 主病史
	// 注意：此字段可能返回 null，表示取不到有效值。
	MainDiseaseHistory *string `json:"MainDiseaseHistory,omitempty" name:"MainDiseaseHistory"`

	// 过敏史
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllergyHistory *string `json:"AllergyHistory,omitempty" name:"AllergyHistory"`

	// 传染疾病史
	// 注意：此字段可能返回 null，表示取不到有效值。
	InfectHistory *string `json:"InfectHistory,omitempty" name:"InfectHistory"`

	// 手术史
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperationHistory *string `json:"OperationHistory,omitempty" name:"OperationHistory"`

	// 输血史
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransfusionHistory *string `json:"TransfusionHistory,omitempty" name:"TransfusionHistory"`
}

type EcgDescription struct {

	// 心率
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeartRate *EcgItem `json:"HeartRate,omitempty" name:"HeartRate"`

	// 心房率
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuricularRate *EcgItem `json:"AuricularRate,omitempty" name:"AuricularRate"`

	// 心室率
	// 注意：此字段可能返回 null，表示取不到有效值。
	VentricularRate *EcgItem `json:"VentricularRate,omitempty" name:"VentricularRate"`

	// 节律
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rhythm *EcgItem `json:"Rhythm,omitempty" name:"Rhythm"`

	// P波时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	PDuration *EcgItem `json:"PDuration,omitempty" name:"PDuration"`

	// QRS时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	QrsDuration *EcgItem `json:"QrsDuration,omitempty" name:"QrsDuration"`

	// QRS电轴
	// 注意：此字段可能返回 null，表示取不到有效值。
	QrsAxis *EcgItem `json:"QrsAxis,omitempty" name:"QrsAxis"`

	// P-R间期
	// 注意：此字段可能返回 null，表示取不到有效值。
	PRInterval *EcgItem `json:"PRInterval,omitempty" name:"PRInterval"`

	// P-P间期
	// 注意：此字段可能返回 null，表示取不到有效值。
	PPInterval *EcgItem `json:"PPInterval,omitempty" name:"PPInterval"`

	// R-R间期
	// 注意：此字段可能返回 null，表示取不到有效值。
	RRInterval *EcgItem `json:"RRInterval,omitempty" name:"RRInterval"`

	// P-J间期
	// 注意：此字段可能返回 null，表示取不到有效值。
	PJInterval *EcgItem `json:"PJInterval,omitempty" name:"PJInterval"`

	// Q-T间期
	// 注意：此字段可能返回 null，表示取不到有效值。
	QTInterval *EcgItem `json:"QTInterval,omitempty" name:"QTInterval"`

	// qt/qtc间期
	// 注意：此字段可能返回 null，表示取不到有效值。
	QTCInterval *EcgItem `json:"QTCInterval,omitempty" name:"QTCInterval"`

	// RV5/SV1振幅
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rv5SV1Amplitude *EcgItem `json:"Rv5SV1Amplitude,omitempty" name:"Rv5SV1Amplitude"`

	// RV5+SV1振幅
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rv5AddSV1Amplitude *EcgItem `json:"Rv5AddSV1Amplitude,omitempty" name:"Rv5AddSV1Amplitude"`

	// PRT电轴
	// 注意：此字段可能返回 null，表示取不到有效值。
	PRTAxis *EcgItem `json:"PRTAxis,omitempty" name:"PRTAxis"`

	// RV5振幅
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rv5Amplitude *EcgItem `json:"Rv5Amplitude,omitempty" name:"Rv5Amplitude"`

	// SV1振幅
	// 注意：此字段可能返回 null，表示取不到有效值。
	SV1Amplitude *EcgItem `json:"SV1Amplitude,omitempty" name:"SV1Amplitude"`

	// RV6/SV2
	// 注意：此字段可能返回 null，表示取不到有效值。
	RV6SV2 *EcgItem `json:"RV6SV2,omitempty" name:"RV6SV2"`

	// P/QRS/T电轴
	// 注意：此字段可能返回 null，表示取不到有效值。
	PQRSTAxis *EcgItem `json:"PQRSTAxis,omitempty" name:"PQRSTAxis"`
}

type EcgDiagnosis struct {

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value []*string `json:"Value,omitempty" name:"Value"`
}

type EcgItem struct {

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`
}

type Elastic struct {

	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitempty" name:"Index"`

	// 分数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *string `json:"Score,omitempty" name:"Score"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`
}

type Electrocardiogram struct {

	// 心电图详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	EcgDescription *EcgDescription `json:"EcgDescription,omitempty" name:"EcgDescription"`

	// 心电图诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	EcgDiagnosis *EcgDiagnosis `json:"EcgDiagnosis,omitempty" name:"EcgDiagnosis"`
}

type Endoscopy struct {

	// 活检部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	BiopsyPart *BiopsyPart `json:"BiopsyPart,omitempty" name:"BiopsyPart"`

	// 可见描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *EndoscopyDesc `json:"Desc,omitempty" name:"Desc"`

	// 结论
	// 注意：此字段可能返回 null，表示取不到有效值。
	Summary *Summary `json:"Summary,omitempty" name:"Summary"`
}

type EndoscopyDesc struct {

	// 描述内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitempty" name:"Text"`

	// 器官
	// 注意：此字段可能返回 null，表示取不到有效值。
	Organ []*EndoscopyOrgan `json:"Organ,omitempty" name:"Organ"`
}

type EndoscopyOrgan struct {

	// 部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Part *Part `json:"Part,omitempty" name:"Part"`

	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitempty" name:"Index"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 部位别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	PartAlias *string `json:"PartAlias,omitempty" name:"PartAlias"`

	// 症状描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	SymDescList []*BlockInfo `json:"SymDescList,omitempty" name:"SymDescList"`
}

type FamilyMedicalHistory struct {

	// 家族成员史
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelativeHistory *string `json:"RelativeHistory,omitempty" name:"RelativeHistory"`

	// 家族肿瘤史
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelativeCancerHistory *string `json:"RelativeCancerHistory,omitempty" name:"RelativeCancerHistory"`

	// 家族遗传史
	// 注意：此字段可能返回 null，表示取不到有效值。
	GeneticHistory *string `json:"GeneticHistory,omitempty" name:"GeneticHistory"`
}

type FirstPage struct {

	// 出入院诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	DischargeDiagnosis []*DischargeDiagnosis `json:"DischargeDiagnosis,omitempty" name:"DischargeDiagnosis"`

	// 病理诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	PathologicalDiagnosis *BlockInfo `json:"PathologicalDiagnosis,omitempty" name:"PathologicalDiagnosis"`

	// 临床诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClinicalDiagnosis *BlockInfo `json:"ClinicalDiagnosis,omitempty" name:"ClinicalDiagnosis"`
}

type HandleParam struct {

	// ocr引擎
	OcrEngineType *int64 `json:"OcrEngineType,omitempty" name:"OcrEngineType"`

	// 是否返回分行文本内容
	IsReturnText *bool `json:"IsReturnText,omitempty" name:"IsReturnText"`

	// 顺时针旋转角度
	RotateTheAngle *float64 `json:"RotateTheAngle,omitempty" name:"RotateTheAngle"`

	// 自动适配方向,仅支持优图引擎
	AutoFitDirection *bool `json:"AutoFitDirection,omitempty" name:"AutoFitDirection"`

	// 坐标优化
	AutoOptimizeCoordinate *bool `json:"AutoOptimizeCoordinate,omitempty" name:"AutoOptimizeCoordinate"`

	// 是否开启图片压缩，开启时imageOriginalSize必须正确填写
	IsScale *bool `json:"IsScale,omitempty" name:"IsScale"`

	// 原始图片大小(单位byes),用来判断该图片是否需要压缩
	ImageOriginalSize *uint64 `json:"ImageOriginalSize,omitempty" name:"ImageOriginalSize"`

	// 采用后台默认值(2048Kb)
	ScaleTargetSize *uint64 `json:"ScaleTargetSize,omitempty" name:"ScaleTargetSize"`
}

type HistologyLevel struct {

	// 等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	Grade *string `json:"Grade,omitempty" name:"Grade"`

	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitempty" name:"Index"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`
}

type HistologyType struct {

	// 浸润
	// 注意：此字段可能返回 null，表示取不到有效值。
	Infiltration *string `json:"Infiltration,omitempty" name:"Infiltration"`

	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitempty" name:"Index"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`
}

type Hospitalization struct {

	// 入院时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdmissionTime *string `json:"AdmissionTime,omitempty" name:"AdmissionTime"`

	// 出院时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DischargeTime *string `json:"DischargeTime,omitempty" name:"DischargeTime"`

	// 住院天数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdmissionDays *string `json:"AdmissionDays,omitempty" name:"AdmissionDays"`

	// 入院诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdmissionDignosis *string `json:"AdmissionDignosis,omitempty" name:"AdmissionDignosis"`

	// 入院情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdmissionCondition *string `json:"AdmissionCondition,omitempty" name:"AdmissionCondition"`

	// 诊疗经过
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiagnosisTreatment *string `json:"DiagnosisTreatment,omitempty" name:"DiagnosisTreatment"`

	// 出院诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	DischargeDiagnosis *string `json:"DischargeDiagnosis,omitempty" name:"DischargeDiagnosis"`

	// 出院医嘱
	// 注意：此字段可能返回 null，表示取不到有效值。
	DischargeInstruction *string `json:"DischargeInstruction,omitempty" name:"DischargeInstruction"`
}

type IHCInfo struct {

	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitempty" name:"Index"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 值
	Value *Value `json:"Value,omitempty" name:"Value"`
}

type ImageInfo struct {

	// 图片id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 图片url
	Url *string `json:"Url,omitempty" name:"Url"`

	// 图片base64编码
	Base64 *string `json:"Base64,omitempty" name:"Base64"`
}

type ImageToClassRequest struct {
	*tchttp.BaseRequest

	// 图片列表，允许传入多张图片，支持传入图片的base64编码，暂不支持图片url
	ImageInfoList []*ImageInfo `json:"ImageInfoList,omitempty" name:"ImageInfoList"`

	// 图片处理参数
	HandleParam *HandleParam `json:"HandleParam,omitempty" name:"HandleParam"`

	// 不填，默认为0
	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

func (r *ImageToClassRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageToClassRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageInfoList")
	delete(f, "HandleParam")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImageToClassRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ImageToClassResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分类结果
	// 注意：此字段可能返回 null，表示取不到有效值。
		TextTypeList []*TextType `json:"TextTypeList,omitempty" name:"TextTypeList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImageToClassResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageToClassResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageToObjectRequest struct {
	*tchttp.BaseRequest

	// 图片列表，允许传入多张图片，目前只支持传入图片base64编码，图片url暂不支持
	ImageInfoList []*ImageInfo `json:"ImageInfoList,omitempty" name:"ImageInfoList"`

	// 图片处理参数
	HandleParam *HandleParam `json:"HandleParam,omitempty" name:"HandleParam"`

	// 报告类型，目前支持11（检验报告），12（检查报告），15（病理报告），28（出院报告），29（入院报告），210（门诊病历），212（手术记录），218（诊断证明），363（心电图），27（内窥镜检查），215（处方单），219（免疫接种证明），301（C14呼气试验）。如果不清楚报告类型，可以使用分类引擎，该字段传0（同时IsUsedClassify字段必须为True，否则无法输出结果）
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 是否使用分类引擎，当不确定报告类型时，可以使用收费的报告分类引擎服务。若该字段为 False，则 Type 字段不能为 0，否则无法输出结果。
	// 注意：当 IsUsedClassify 为True 时，表示使用收费的报告分类服务，将会产生额外的费用，具体收费标准参见 [购买指南的产品价格](https://cloud.tencent.com/document/product/1314/54264)。
	IsUsedClassify *bool `json:"IsUsedClassify,omitempty" name:"IsUsedClassify"`
}

func (r *ImageToObjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageToObjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageInfoList")
	delete(f, "HandleParam")
	delete(f, "Type")
	delete(f, "IsUsedClassify")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImageToObjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ImageToObjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 报告结构化结果
	// 注意：此字段可能返回 null，表示取不到有效值。
		Template *Template `json:"Template,omitempty" name:"Template"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImageToObjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageToObjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Indicator struct {

	// 检验指标项
	// 注意：此字段可能返回 null，表示取不到有效值。
	Indicators []*IndicatorItem `json:"Indicators,omitempty" name:"Indicators"`
}

type IndicatorItem struct {

	// 英文缩写
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *string `json:"Code,omitempty" name:"Code"`

	// 标准缩写
	// 注意：此字段可能返回 null，表示取不到有效值。
	Scode *string `json:"Scode,omitempty" name:"Scode"`

	// 项目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 标准名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sname *string `json:"Sname,omitempty" name:"Sname"`

	// 结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// 参考范围
	// 注意：此字段可能返回 null，表示取不到有效值。
	Range *string `json:"Range,omitempty" name:"Range"`

	// 上下箭头
	// 注意：此字段可能返回 null，表示取不到有效值。
	Arrow *string `json:"Arrow,omitempty" name:"Arrow"`

	// 是否正常
	// 注意：此字段可能返回 null，表示取不到有效值。
	Normal *bool `json:"Normal,omitempty" name:"Normal"`

	// 项目原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	ItemString *string `json:"ItemString,omitempty" name:"ItemString"`
}

type Invas struct {

	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitempty" name:"Index"`

	// 部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Part *Part `json:"Part,omitempty" name:"Part"`

	// 阳性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Positive *string `json:"Positive,omitempty" name:"Positive"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`
}

type Lymph struct {

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitempty" name:"Index"`

	// 部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Part *Part `json:"Part,omitempty" name:"Part"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 转移数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferNum *int64 `json:"TransferNum,omitempty" name:"TransferNum"`
}

type MedDoc struct {

	// 建议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Advice *Advice `json:"Advice,omitempty" name:"Advice"`

	// 诊断结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Diagnosis []*DiagCertItem `json:"Diagnosis,omitempty" name:"Diagnosis"`

	// 疾病史
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiseaseMedicalHistory *DiseaseMedicalHistory `json:"DiseaseMedicalHistory,omitempty" name:"DiseaseMedicalHistory"`

	// 个人史
	PersonalMedicalHistory *PersonalMedicalHistory `json:"PersonalMedicalHistory,omitempty" name:"PersonalMedicalHistory"`

	// 婚孕史
	ObstericalMedicalHistory *ObstericalMedicalHistory `json:"ObstericalMedicalHistory,omitempty" name:"ObstericalMedicalHistory"`

	// 家族史
	FamilyMedicalHistory *FamilyMedicalHistory `json:"FamilyMedicalHistory,omitempty" name:"FamilyMedicalHistory"`

	// 月经史
	MenstrualMedicalHistory *MenstrualMedicalHistory `json:"MenstrualMedicalHistory,omitempty" name:"MenstrualMedicalHistory"`

	// 诊疗记录
	TreatmentRecord *TreatmentRecord `json:"TreatmentRecord,omitempty" name:"TreatmentRecord"`
}

type MedicalRecordInfo struct {

	// 就诊日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiagnosisTime *string `json:"DiagnosisTime,omitempty" name:"DiagnosisTime"`

	// 就诊科室
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiagnosisDepartmentName *string `json:"DiagnosisDepartmentName,omitempty" name:"DiagnosisDepartmentName"`

	// 就诊医生
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiagnosisDoctorName *string `json:"DiagnosisDoctorName,omitempty" name:"DiagnosisDoctorName"`

	// 临床诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClinicalDiagnosis *string `json:"ClinicalDiagnosis,omitempty" name:"ClinicalDiagnosis"`

	// 主述
	// 注意：此字段可能返回 null，表示取不到有效值。
	MainNarration *string `json:"MainNarration,omitempty" name:"MainNarration"`

	// 体格检查
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhysicalExamination *string `json:"PhysicalExamination,omitempty" name:"PhysicalExamination"`

	// 检查结论
	// 注意：此字段可能返回 null，表示取不到有效值。
	InspectionFindings *string `json:"InspectionFindings,omitempty" name:"InspectionFindings"`

	// 治疗意见
	// 注意：此字段可能返回 null，表示取不到有效值。
	TreatmentOpinion *string `json:"TreatmentOpinion,omitempty" name:"TreatmentOpinion"`
}

type Medicine struct {

	// 药品名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 商品名
	// 注意：此字段可能返回 null，表示取不到有效值。
	TradeName *string `json:"TradeName,omitempty" name:"TradeName"`

	// 厂商
	// 注意：此字段可能返回 null，表示取不到有效值。
	Firm *string `json:"Firm,omitempty" name:"Firm"`

	// 医保类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Category *string `json:"Category,omitempty" name:"Category"`

	// 规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	Specification *string `json:"Specification,omitempty" name:"Specification"`

	// 最小包装数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinQuantity *string `json:"MinQuantity,omitempty" name:"MinQuantity"`

	// 最小制剂单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	DosageUnit *string `json:"DosageUnit,omitempty" name:"DosageUnit"`

	// 最小包装单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackingUnit *string `json:"PackingUnit,omitempty" name:"PackingUnit"`
}

type MenstrualMedicalHistory struct {

	// 末次月经时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastMenstrualPeriod *string `json:"LastMenstrualPeriod,omitempty" name:"LastMenstrualPeriod"`

	// 经量
	// 注意：此字段可能返回 null，表示取不到有效值。
	MenstrualFlow *string `json:"MenstrualFlow,omitempty" name:"MenstrualFlow"`

	// 月经初潮年龄
	// 注意：此字段可能返回 null，表示取不到有效值。
	MenarcheAge *string `json:"MenarcheAge,omitempty" name:"MenarcheAge"`

	// 是否来月经
	// 注意：此字段可能返回 null，表示取不到有效值。
	MenstruationOrNot *string `json:"MenstruationOrNot,omitempty" name:"MenstruationOrNot"`

	// 月经周期
	// 注意：此字段可能返回 null，表示取不到有效值。
	MenstrualCycles *string `json:"MenstrualCycles,omitempty" name:"MenstrualCycles"`

	// 月经持续天数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MenstrualPeriod *string `json:"MenstrualPeriod,omitempty" name:"MenstrualPeriod"`
}

type Multiple struct {

	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitempty" name:"Index"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`
}

type NormPart struct {

	// 部位值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Part *string `json:"Part,omitempty" name:"Part"`

	// 部位方向
	// 注意：此字段可能返回 null，表示取不到有效值。
	PartDirection *string `json:"PartDirection,omitempty" name:"PartDirection"`

	// 组织值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tissue *string `json:"Tissue,omitempty" name:"Tissue"`

	// 组织方向
	// 注意：此字段可能返回 null，表示取不到有效值。
	TissueDirection *string `json:"TissueDirection,omitempty" name:"TissueDirection"`

	// 上级部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Upper *string `json:"Upper,omitempty" name:"Upper"`
}

type NormSize struct {

	// 数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Number []*string `json:"Number,omitempty" name:"Number"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *string `json:"Unit,omitempty" name:"Unit"`
}

type ObstericalMedicalHistory struct {

	// 婚史
	// 注意：此字段可能返回 null，表示取不到有效值。
	MarriageHistory *string `json:"MarriageHistory,omitempty" name:"MarriageHistory"`

	// 孕史
	// 注意：此字段可能返回 null，表示取不到有效值。
	FertilityHistory *string `json:"FertilityHistory,omitempty" name:"FertilityHistory"`
}

type Organ struct {

	// 部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Part *Part `json:"Part,omitempty" name:"Part"`

	// 大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size []*Size `json:"Size,omitempty" name:"Size"`

	// 包膜
	// 注意：此字段可能返回 null，表示取不到有效值。
	Envelope *BlockInfo `json:"Envelope,omitempty" name:"Envelope"`

	// 边缘
	// 注意：此字段可能返回 null，表示取不到有效值。
	Edge *BlockInfo `json:"Edge,omitempty" name:"Edge"`

	// 内部回声
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerEcho *BlockInfo `json:"InnerEcho,omitempty" name:"InnerEcho"`

	// 腺体
	// 注意：此字段可能返回 null，表示取不到有效值。
	Gland *BlockInfo `json:"Gland,omitempty" name:"Gland"`

	// 形状
	// 注意：此字段可能返回 null，表示取不到有效值。
	Shape *BlockInfo `json:"Shape,omitempty" name:"Shape"`

	// 厚度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Thickness *BlockInfo `json:"Thickness,omitempty" name:"Thickness"`

	// 形态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShapeAttr *BlockInfo `json:"ShapeAttr,omitempty" name:"ShapeAttr"`

	// 血液cdfi
	// 注意：此字段可能返回 null，表示取不到有效值。
	CDFI *BlockInfo `json:"CDFI,omitempty" name:"CDFI"`

	// 描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SymDesc *BlockInfo `json:"SymDesc,omitempty" name:"SymDesc"`

	// 大小状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	SizeStatus *BlockInfo `json:"SizeStatus,omitempty" name:"SizeStatus"`

	// 轮廓
	// 注意：此字段可能返回 null，表示取不到有效值。
	Outline *BlockInfo `json:"Outline,omitempty" name:"Outline"`

	// 结构
	// 注意：此字段可能返回 null，表示取不到有效值。
	Structure *BlockInfo `json:"Structure,omitempty" name:"Structure"`

	// 密度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Density *BlockInfo `json:"Density,omitempty" name:"Density"`

	// 血管
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vas *BlockInfo `json:"Vas,omitempty" name:"Vas"`

	// 囊壁
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cysticwall *BlockInfo `json:"Cysticwall,omitempty" name:"Cysticwall"`

	// 被膜
	// 注意：此字段可能返回 null，表示取不到有效值。
	Capsule *BlockInfo `json:"Capsule,omitempty" name:"Capsule"`

	// 峡部厚度
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsthmusThicknese *Size `json:"IsthmusThicknese,omitempty" name:"IsthmusThicknese"`

	// 内部回声分布
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerEchoDistribution *BlockInfo `json:"InnerEchoDistribution,omitempty" name:"InnerEchoDistribution"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitempty" name:"Index"`
}

type Part struct {

	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitempty" name:"Index"`

	// 部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	NormPart *NormPart `json:"NormPart,omitempty" name:"NormPart"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`
}

type PathologyReport struct {

	// 癌症部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	CancerPart *Part `json:"CancerPart,omitempty" name:"CancerPart"`

	// 癌症部位大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	CancerSize []*Size `json:"CancerSize,omitempty" name:"CancerSize"`

	// 描述文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	DescText *string `json:"DescText,omitempty" name:"DescText"`

	// 组织学等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	HistologyLevel *HistologyLevel `json:"HistologyLevel,omitempty" name:"HistologyLevel"`

	// 组织学类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	HistologyType *HistologyType `json:"HistologyType,omitempty" name:"HistologyType"`

	// IHC信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	IHC []*IHCInfo `json:"IHC,omitempty" name:"IHC"`

	// 浸润深度
	// 注意：此字段可能返回 null，表示取不到有效值。
	InfiltrationDepth *BlockInfo `json:"InfiltrationDepth,omitempty" name:"InfiltrationDepth"`

	// 肿瘤扩散
	// 注意：此字段可能返回 null，表示取不到有效值。
	Invasive []*Invas `json:"Invasive,omitempty" name:"Invasive"`

	// 淋巴结
	// 注意：此字段可能返回 null，表示取不到有效值。
	LymphNodes []*Lymph `json:"LymphNodes,omitempty" name:"LymphNodes"`

	// PTNM信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PTNM *BlockInfo `json:"PTNM,omitempty" name:"PTNM"`

	// 病理报告类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	PathologicalReportType *BlockInfo `json:"PathologicalReportType,omitempty" name:"PathologicalReportType"`

	// 报告原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportText *string `json:"ReportText,omitempty" name:"ReportText"`

	// 标本类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	SampleType *BlockInfo `json:"SampleType,omitempty" name:"SampleType"`

	// 结论文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	SummaryText *string `json:"SummaryText,omitempty" name:"SummaryText"`
}

type PatientInfo struct {

	// 患者姓名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 患者性别
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sex *string `json:"Sex,omitempty" name:"Sex"`

	// 患者年龄
	// 注意：此字段可能返回 null，表示取不到有效值。
	Age *string `json:"Age,omitempty" name:"Age"`

	// 患者手机号码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 患者地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Address *string `json:"Address,omitempty" name:"Address"`

	// 患者身份证
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdCard *string `json:"IdCard,omitempty" name:"IdCard"`

	// 健康卡号
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthCardNo *string `json:"HealthCardNo,omitempty" name:"HealthCardNo"`

	// 社保卡号
	// 注意：此字段可能返回 null，表示取不到有效值。
	SocialSecurityCardNo *string `json:"SocialSecurityCardNo,omitempty" name:"SocialSecurityCardNo"`

	// 出生日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	Birthday *string `json:"Birthday,omitempty" name:"Birthday"`

	// 民族
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ethnicity *string `json:"Ethnicity,omitempty" name:"Ethnicity"`

	// 婚姻状况
	// 注意：此字段可能返回 null，表示取不到有效值。
	Married *string `json:"Married,omitempty" name:"Married"`

	// 职业
	// 注意：此字段可能返回 null，表示取不到有效值。
	Profession *string `json:"Profession,omitempty" name:"Profession"`

	// 教育程度
	// 注意：此字段可能返回 null，表示取不到有效值。
	EducationBackground *string `json:"EducationBackground,omitempty" name:"EducationBackground"`

	// 国籍
	// 注意：此字段可能返回 null，表示取不到有效值。
	Nationality *string `json:"Nationality,omitempty" name:"Nationality"`

	// 籍贯
	// 注意：此字段可能返回 null，表示取不到有效值。
	BirthPlace *string `json:"BirthPlace,omitempty" name:"BirthPlace"`

	// 医保类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	MedicalInsuranceType *string `json:"MedicalInsuranceType,omitempty" name:"MedicalInsuranceType"`

	// 标准化年龄
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgeNorm *string `json:"AgeNorm,omitempty" name:"AgeNorm"`

	// 民族
	// 注意：此字段可能返回 null，表示取不到有效值。
	Nation *string `json:"Nation,omitempty" name:"Nation"`

	// 婚姻代码
	// 注意：此字段可能返回 null，表示取不到有效值。
	MarriedCode *string `json:"MarriedCode,omitempty" name:"MarriedCode"`

	// 职业代码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProfessionCode *string `json:"ProfessionCode,omitempty" name:"ProfessionCode"`

	// 居民医保代码
	// 注意：此字段可能返回 null，表示取不到有效值。
	MedicalInsuranceTypeCode *string `json:"MedicalInsuranceTypeCode,omitempty" name:"MedicalInsuranceTypeCode"`
}

type PersonalMedicalHistory struct {

	// 出生史
	// 注意：此字段可能返回 null，表示取不到有效值。
	BirthPlace *string `json:"BirthPlace,omitempty" name:"BirthPlace"`

	// 居住史
	// 注意：此字段可能返回 null，表示取不到有效值。
	LivePlace *string `json:"LivePlace,omitempty" name:"LivePlace"`

	// 工作史
	// 注意：此字段可能返回 null，表示取不到有效值。
	Job *string `json:"Job,omitempty" name:"Job"`

	// 吸烟史
	// 注意：此字段可能返回 null，表示取不到有效值。
	SmokeHistory *string `json:"SmokeHistory,omitempty" name:"SmokeHistory"`

	// 饮酒史
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlcoholicHistory *string `json:"AlcoholicHistory,omitempty" name:"AlcoholicHistory"`
}

type Prescription struct {

	// 药品列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	MedicineList []*Medicine `json:"MedicineList,omitempty" name:"MedicineList"`
}

type ReportInfo struct {

	// 医院名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Hospital *string `json:"Hospital,omitempty" name:"Hospital"`

	// 科室名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DepartmentName *string `json:"DepartmentName,omitempty" name:"DepartmentName"`

	// 申请时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillingTime *string `json:"BillingTime,omitempty" name:"BillingTime"`

	// 报告时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportTime *string `json:"ReportTime,omitempty" name:"ReportTime"`

	// 检查时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	InspectTime *string `json:"InspectTime,omitempty" name:"InspectTime"`

	// 检查号
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckNum *string `json:"CheckNum,omitempty" name:"CheckNum"`

	// 影像号
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageNum *string `json:"ImageNum,omitempty" name:"ImageNum"`

	// 放射号
	// 注意：此字段可能返回 null，表示取不到有效值。
	RadiationNum *string `json:"RadiationNum,omitempty" name:"RadiationNum"`

	// 检验号
	// 注意：此字段可能返回 null，表示取不到有效值。
	TestNum *string `json:"TestNum,omitempty" name:"TestNum"`

	// 门诊号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutpatientNum *string `json:"OutpatientNum,omitempty" name:"OutpatientNum"`

	// 病理号
	// 注意：此字段可能返回 null，表示取不到有效值。
	PathologyNum *string `json:"PathologyNum,omitempty" name:"PathologyNum"`

	// 住院号
	// 注意：此字段可能返回 null，表示取不到有效值。
	InHospitalNum *string `json:"InHospitalNum,omitempty" name:"InHospitalNum"`

	// 样本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	SampleNum *string `json:"SampleNum,omitempty" name:"SampleNum"`

	// 标本种类
	// 注意：此字段可能返回 null，表示取不到有效值。
	SampleType *string `json:"SampleType,omitempty" name:"SampleType"`

	// 病历号
	// 注意：此字段可能返回 null，表示取不到有效值。
	MedicalRecordNum *string `json:"MedicalRecordNum,omitempty" name:"MedicalRecordNum"`

	// 报告名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportName *string `json:"ReportName,omitempty" name:"ReportName"`

	// 超声号
	// 注意：此字段可能返回 null，表示取不到有效值。
	UltraNum *string `json:"UltraNum,omitempty" name:"UltraNum"`

	// 临床诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	Diagnose *string `json:"Diagnose,omitempty" name:"Diagnose"`
}

type Size struct {

	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitempty" name:"Index"`

	// 标准大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	NormSize *NormSize `json:"NormSize,omitempty" name:"NormSize"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type Summary struct {

	// 症状
	// 注意：此字段可能返回 null，表示取不到有效值。
	Symptom []*SymptomInfo `json:"Symptom,omitempty" name:"Symptom"`

	// 文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitempty" name:"Text"`
}

type Surgery struct {

	// 手术史
	// 注意：此字段可能返回 null，表示取不到有效值。
	SurgeryHistory *SurgeryHistory `json:"SurgeryHistory,omitempty" name:"SurgeryHistory"`
}

type SurgeryAttr struct {

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type SurgeryHistory struct {

	// 手术名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SurgeryName *SurgeryAttr `json:"SurgeryName,omitempty" name:"SurgeryName"`

	// 手术日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	SurgeryDate *SurgeryAttr `json:"SurgeryDate,omitempty" name:"SurgeryDate"`

	// 术前诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	PreoperativePathology *SurgeryAttr `json:"PreoperativePathology,omitempty" name:"PreoperativePathology"`

	// 术中诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntraoperativePathology *SurgeryAttr `json:"IntraoperativePathology,omitempty" name:"IntraoperativePathology"`

	// 术后诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	PostoperativePathology *SurgeryAttr `json:"PostoperativePathology,omitempty" name:"PostoperativePathology"`

	// 出院诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	DischargeDiagnosis *SurgeryAttr `json:"DischargeDiagnosis,omitempty" name:"DischargeDiagnosis"`
}

type SymptomInfo struct {

	// 等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	Grade *BlockInfo `json:"Grade,omitempty" name:"Grade"`

	// 部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Part *Part `json:"Part,omitempty" name:"Part"`

	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitempty" name:"Index"`

	// 病变
	// 注意：此字段可能返回 null，表示取不到有效值。
	Symptom *BlockInfo `json:"Symptom,omitempty" name:"Symptom"`

	// 属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Attrs []*BlockInfo `json:"Attrs,omitempty" name:"Attrs"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`
}

type Template struct {

	// 患者信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PatientInfo *PatientInfo `json:"PatientInfo,omitempty" name:"PatientInfo"`

	// 报告信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportInfo *ReportInfo `json:"ReportInfo,omitempty" name:"ReportInfo"`

	// 检查报告
	// 注意：此字段可能返回 null，表示取不到有效值。
	Check *Check `json:"Check,omitempty" name:"Check"`

	// 病理报告
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pathology *PathologyReport `json:"Pathology,omitempty" name:"Pathology"`

	// 出院报告，入院报告，门诊病历
	// 注意：此字段可能返回 null，表示取不到有效值。
	MedDoc *MedDoc `json:"MedDoc,omitempty" name:"MedDoc"`

	// 诊断证明
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiagCert *DiagCert `json:"DiagCert,omitempty" name:"DiagCert"`

	// 病案首页
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstPage *FirstPage `json:"FirstPage,omitempty" name:"FirstPage"`

	// 检验报告
	// 注意：此字段可能返回 null，表示取不到有效值。
	Indicator *Indicator `json:"Indicator,omitempty" name:"Indicator"`

	// 报告类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportType *string `json:"ReportType,omitempty" name:"ReportType"`

	// 门诊病历信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MedicalRecordInfo *MedicalRecordInfo `json:"MedicalRecordInfo,omitempty" name:"MedicalRecordInfo"`

	// 出入院信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Hospitalization *Hospitalization `json:"Hospitalization,omitempty" name:"Hospitalization"`

	// 手术记录
	// 注意：此字段可能返回 null，表示取不到有效值。
	Surgery *Surgery `json:"Surgery,omitempty" name:"Surgery"`

	// 心电图报告
	// 注意：此字段可能返回 null，表示取不到有效值。
	Electrocardiogram *Electrocardiogram `json:"Electrocardiogram,omitempty" name:"Electrocardiogram"`

	// 内窥镜报告
	// 注意：此字段可能返回 null，表示取不到有效值。
	Endoscopy *Endoscopy `json:"Endoscopy,omitempty" name:"Endoscopy"`

	// 处方单
	// 注意：此字段可能返回 null，表示取不到有效值。
	Prescription *Prescription `json:"Prescription,omitempty" name:"Prescription"`

	// 免疫接种证明
	// 注意：此字段可能返回 null，表示取不到有效值。
	VaccineCertificate *VaccineCertificate `json:"VaccineCertificate,omitempty" name:"VaccineCertificate"`

	// OCR文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrText *string `json:"OcrText,omitempty" name:"OcrText"`

	// OCR拼接后文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrResult *string `json:"OcrResult,omitempty" name:"OcrResult"`
}

type TextToClassRequest struct {
	*tchttp.BaseRequest

	// 报告文本
	Text *string `json:"Text,omitempty" name:"Text"`
}

func (r *TextToClassRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextToClassRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Text")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TextToClassRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type TextToClassResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分类结果
		TextTypeList []*TextType `json:"TextTypeList,omitempty" name:"TextTypeList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TextToClassResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextToClassResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TextToObjectRequest struct {
	*tchttp.BaseRequest

	// 报告文本
	Text *string `json:"Text,omitempty" name:"Text"`

	// 报告类型，目前支持12（检查报告），15（病理报告），28（出院报告），29（入院报告），210（门诊病历），212（手术记录），218（诊断证明），363（心电图），27（内窥镜检查），215（处方单），219（免疫接种证明），301（C14呼气试验）。如果不清楚报告类型，可以使用分类引擎，该字段传0（同时IsUsedClassify字段必须为True，否则无法输出结果）
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 是否使用分类引擎，当不确定报告类型时，可以使用收费的报告分类引擎服务。若该字段为False，则Type字段不能为0，否则无法输出结果。
	// 注意：当 IsUsedClassify 为True 时，表示使用收费的报告分类服务，将会产生额外的费用，具体收费标准参见 [购买指南的产品价格](https://cloud.tencent.com/document/product/1314/54264)。
	IsUsedClassify *bool `json:"IsUsedClassify,omitempty" name:"IsUsedClassify"`
}

func (r *TextToObjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextToObjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Text")
	delete(f, "Type")
	delete(f, "IsUsedClassify")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TextToObjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type TextToObjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 报告结构化结果
		Template *Template `json:"Template,omitempty" name:"Template"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TextToObjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextToObjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TextType struct {

	// 类别Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 类别层级
	// 注意：此字段可能返回 null，表示取不到有效值。
	Level *int64 `json:"Level,omitempty" name:"Level"`

	// 类别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`
}

type TreatmentRecord struct {

	// 入院
	// 注意：此字段可能返回 null，表示取不到有效值。
	DmissionCondition *string `json:"DmissionCondition,omitempty" name:"DmissionCondition"`

	// 主诉
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChiefComplaint *string `json:"ChiefComplaint,omitempty" name:"ChiefComplaint"`

	// 现病史
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiseasePresent *string `json:"DiseasePresent,omitempty" name:"DiseasePresent"`

	// 主要症状体征
	// 注意：此字段可能返回 null，表示取不到有效值。
	SymptomsAndSigns *string `json:"SymptomsAndSigns,omitempty" name:"SymptomsAndSigns"`

	// 辅助检查
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuxiliaryExamination *string `json:"AuxiliaryExamination,omitempty" name:"AuxiliaryExamination"`

	// 体格检查
	// 注意：此字段可能返回 null，表示取不到有效值。
	BodyExamination *string `json:"BodyExamination,omitempty" name:"BodyExamination"`

	// 专科检查
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecialistExamination *string `json:"SpecialistExamination,omitempty" name:"SpecialistExamination"`

	// 精神检查
	// 注意：此字段可能返回 null，表示取不到有效值。
	MentalExamination *string `json:"MentalExamination,omitempty" name:"MentalExamination"`

	// 检查记录
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckRecord *string `json:"CheckRecord,omitempty" name:"CheckRecord"`

	// 化验结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	InspectResult *string `json:"InspectResult,omitempty" name:"InspectResult"`

	// 切口愈合情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncisionHealing *string `json:"IncisionHealing,omitempty" name:"IncisionHealing"`

	// 处理意见
	// 注意：此字段可能返回 null，表示取不到有效值。
	TreatmentSuggestion *string `json:"TreatmentSuggestion,omitempty" name:"TreatmentSuggestion"`

	// 门诊随访要求
	// 注意：此字段可能返回 null，表示取不到有效值。
	FollowUpRequirements *string `json:"FollowUpRequirements,omitempty" name:"FollowUpRequirements"`

	// 诊疗经过
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckAndTreatmentProcess *string `json:"CheckAndTreatmentProcess,omitempty" name:"CheckAndTreatmentProcess"`

	// 手术经过
	// 注意：此字段可能返回 null，表示取不到有效值。
	SurgeryCondition *string `json:"SurgeryCondition,omitempty" name:"SurgeryCondition"`

	// 入院情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConditionChanges *string `json:"ConditionChanges,omitempty" name:"ConditionChanges"`

	// 出院情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	DischargeCondition *string `json:"DischargeCondition,omitempty" name:"DischargeCondition"`

	// pTNM信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PTNM *string `json:"PTNM,omitempty" name:"PTNM"`

	// pTNMM信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PTNMM *string `json:"PTNMM,omitempty" name:"PTNMM"`

	// pTNMN信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PTNMN *string `json:"PTNMN,omitempty" name:"PTNMN"`

	// pTNMT信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PTNMT *string `json:"PTNMT,omitempty" name:"PTNMT"`

	// ECOG信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ECOG *string `json:"ECOG,omitempty" name:"ECOG"`

	// NRS信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	NRS *string `json:"NRS,omitempty" name:"NRS"`

	// KPS信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	KPS *string `json:"KPS,omitempty" name:"KPS"`

	// 死亡日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeathDate *string `json:"DeathDate,omitempty" name:"DeathDate"`

	// 复发日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelapseDate *string `json:"RelapseDate,omitempty" name:"RelapseDate"`

	// 观测天数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObservationDays *string `json:"ObservationDays,omitempty" name:"ObservationDays"`
}

type TuberInfo struct {

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *BlockInfo `json:"Type,omitempty" name:"Type"`

	// 部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Part *Part `json:"Part,omitempty" name:"Part"`

	// 大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size []*Size `json:"Size,omitempty" name:"Size"`

	// 多发
	// 注意：此字段可能返回 null，表示取不到有效值。
	Multiple *Multiple `json:"Multiple,omitempty" name:"Multiple"`

	// 纵横比
	// 注意：此字段可能返回 null，表示取不到有效值。
	AspectRatio *AspectRatio `json:"AspectRatio,omitempty" name:"AspectRatio"`

	// 边缘
	// 注意：此字段可能返回 null，表示取不到有效值。
	Edge *BlockInfo `json:"Edge,omitempty" name:"Edge"`

	// 内部回声
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerEcho *BlockInfo `json:"InnerEcho,omitempty" name:"InnerEcho"`

	// 外部回声
	// 注意：此字段可能返回 null，表示取不到有效值。
	RearEcho *BlockInfo `json:"RearEcho,omitempty" name:"RearEcho"`

	// 弹性质地
	// 注意：此字段可能返回 null，表示取不到有效值。
	Elastic *Elastic `json:"Elastic,omitempty" name:"Elastic"`

	// 形状
	// 注意：此字段可能返回 null，表示取不到有效值。
	Shape *BlockInfo `json:"Shape,omitempty" name:"Shape"`

	// 形态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShapeAttr *BlockInfo `json:"ShapeAttr,omitempty" name:"ShapeAttr"`

	// 皮髓质信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SkinMedulla *BlockInfo `json:"SkinMedulla,omitempty" name:"SkinMedulla"`

	// 变化趋势
	// 注意：此字段可能返回 null，表示取不到有效值。
	Trend *BlockInfo `json:"Trend,omitempty" name:"Trend"`

	// 钙化
	// 注意：此字段可能返回 null，表示取不到有效值。
	Calcification *BlockInfo `json:"Calcification,omitempty" name:"Calcification"`

	// 包膜
	// 注意：此字段可能返回 null，表示取不到有效值。
	Envelope *BlockInfo `json:"Envelope,omitempty" name:"Envelope"`

	// 强化
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enhancement *BlockInfo `json:"Enhancement,omitempty" name:"Enhancement"`

	// 淋巴结
	// 注意：此字段可能返回 null，表示取不到有效值。
	LymphEnlargement *BlockInfo `json:"LymphEnlargement,omitempty" name:"LymphEnlargement"`

	// 淋巴门
	// 注意：此字段可能返回 null，表示取不到有效值。
	LymphDoor *BlockInfo `json:"LymphDoor,omitempty" name:"LymphDoor"`

	// 活动度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Activity *BlockInfo `json:"Activity,omitempty" name:"Activity"`

	// 手术情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operation *BlockInfo `json:"Operation,omitempty" name:"Operation"`

	// 血液cdfi
	// 注意：此字段可能返回 null，表示取不到有效值。
	CDFI *BlockInfo `json:"CDFI,omitempty" name:"CDFI"`

	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitempty" name:"Index"`

	// 大小状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	SizeStatus *BlockInfo `json:"SizeStatus,omitempty" name:"SizeStatus"`

	// 内部回声分布
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerEchoDistribution *BlockInfo `json:"InnerEchoDistribution,omitempty" name:"InnerEchoDistribution"`

	// 内部回声类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerEchoType []*BlockInfo `json:"InnerEchoType,omitempty" name:"InnerEchoType"`

	// 轮廓
	// 注意：此字段可能返回 null，表示取不到有效值。
	Outline *BlockInfo `json:"Outline,omitempty" name:"Outline"`

	// 结构
	// 注意：此字段可能返回 null，表示取不到有效值。
	Structure *BlockInfo `json:"Structure,omitempty" name:"Structure"`

	// 密度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Density *BlockInfo `json:"Density,omitempty" name:"Density"`

	// 血管
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vas *BlockInfo `json:"Vas,omitempty" name:"Vas"`

	// 囊壁
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cysticwall *BlockInfo `json:"Cysticwall,omitempty" name:"Cysticwall"`

	// 被膜
	// 注意：此字段可能返回 null，表示取不到有效值。
	Capsule *BlockInfo `json:"Capsule,omitempty" name:"Capsule"`

	// 峡部厚度
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsthmusThicknese *Size `json:"IsthmusThicknese,omitempty" name:"IsthmusThicknese"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`
}

type Vaccination struct {

	// 序号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 疫苗名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vaccine *string `json:"Vaccine,omitempty" name:"Vaccine"`

	// 剂次
	// 注意：此字段可能返回 null，表示取不到有效值。
	Dose *string `json:"Dose,omitempty" name:"Dose"`

	// 接种日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	Date *string `json:"Date,omitempty" name:"Date"`

	// 疫苗批号
	// 注意：此字段可能返回 null，表示取不到有效值。
	LotNumber *string `json:"LotNumber,omitempty" name:"LotNumber"`

	// 生产企业
	// 注意：此字段可能返回 null，表示取不到有效值。
	Manufacturer *string `json:"Manufacturer,omitempty" name:"Manufacturer"`

	// 接种单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Clinic *string `json:"Clinic,omitempty" name:"Clinic"`

	// 接种部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Site *string `json:"Site,omitempty" name:"Site"`

	// 接种者
	// 注意：此字段可能返回 null，表示取不到有效值。
	Provider *string `json:"Provider,omitempty" name:"Provider"`
}

type VaccineCertificate struct {

	// 免疫接种列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	VaccineList []*Vaccination `json:"VaccineList,omitempty" name:"VaccineList"`
}

type Value struct {

	// 等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	Grade *string `json:"Grade,omitempty" name:"Grade"`

	// 百分比
	// 注意：此字段可能返回 null，表示取不到有效值。
	Percent []*float64 `json:"Percent,omitempty" name:"Percent"`

	// 阳性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Positive *string `json:"Positive,omitempty" name:"Positive"`
}
