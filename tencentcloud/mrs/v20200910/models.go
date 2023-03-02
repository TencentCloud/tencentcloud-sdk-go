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

type AdmissionConditionBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 原文
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 对外输出值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type AdmissionDiagnosisBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 归一化值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Norm *string `json:"Norm,omitempty" name:"Norm"`

	// 对外输出值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

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

type BaseInfo struct {
	// 项目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 索引
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitempty" name:"Index"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 标准值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 原文对应坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitempty" name:"Coords"`
}

type BaseItem struct {
	// 类型名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 原始文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 归一化后值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 四点坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitempty" name:"Coords"`
}

type BaseItem2 struct {
	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 原始文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 归一化后值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 四点坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitempty" name:"Coords"`
}

type BaseItem3 struct {
	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 原始文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 归一化后值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 四点坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitempty" name:"Coords"`

	// 第几次
	// 注意：此字段可能返回 null，表示取不到有效值。
	Order *int64 `json:"Order,omitempty" name:"Order"`
}

type BiopsyPart struct {
	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitempty" name:"Coords"`
}

type BirthCert struct {
	// 新生儿信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	NeonatalInfo *NeonatalInfo `json:"NeonatalInfo,omitempty" name:"NeonatalInfo"`

	// 母亲信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MotherInfo *ParentInfo `json:"MotherInfo,omitempty" name:"MotherInfo"`

	// 父亲信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FatherInfo *ParentInfo `json:"FatherInfo,omitempty" name:"FatherInfo"`

	// 签发信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	IssueInfo *IssueInfo `json:"IssueInfo,omitempty" name:"IssueInfo"`
}

type BirthPlaceBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 对外输出值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
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

	// 大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size []*Size `json:"Size,omitempty" name:"Size"`
}

type BlockInfoV2 struct {
	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitempty" name:"Index"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 疾病编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *string `json:"Code,omitempty" name:"Code"`
}

type BloodPressureBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 归一化值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Norm *string `json:"Norm,omitempty" name:"Norm"`

	// 单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// 对外输出值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 舒张压
	// 注意：此字段可能返回 null，表示取不到有效值。
	NormDiastolic *string `json:"NormDiastolic,omitempty" name:"NormDiastolic"`

	// 收缩压
	// 注意：此字段可能返回 null，表示取不到有效值。
	NormSystolic *string `json:"NormSystolic,omitempty" name:"NormSystolic"`
}

type BodyExaminationBlock struct {
	// 体温
	// 注意：此字段可能返回 null，表示取不到有效值。
	BodyTemperature *BodyTemperatureBlock `json:"BodyTemperature,omitempty" name:"BodyTemperature"`

	// 脉搏
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pulse *BodyTemperatureBlock `json:"Pulse,omitempty" name:"Pulse"`

	// 呼吸
	// 注意：此字段可能返回 null，表示取不到有效值。
	Breathe *BodyTemperatureBlock `json:"Breathe,omitempty" name:"Breathe"`

	// 血压
	// 注意：此字段可能返回 null，表示取不到有效值。
	BloodPressure *BloodPressureBlock `json:"BloodPressure,omitempty" name:"BloodPressure"`
}

type BodyTemperatureBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 归一化值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Norm *string `json:"Norm,omitempty" name:"Norm"`

	// 单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// 对外输出值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type Check struct {
	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *Desc `json:"Desc,omitempty" name:"Desc"`

	// 结论
	// 注意：此字段可能返回 null，表示取不到有效值。
	Summary *Summary `json:"Summary,omitempty" name:"Summary"`
}

type ChiefComplaintBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 单位输出值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 主诉详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Detail []*ChiefComplaintDetailBlock `json:"Detail,omitempty" name:"Detail"`
}

type ChiefComplaintDetailBlock struct {
	// 疾病名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiseaseName *string `json:"DiseaseName,omitempty" name:"DiseaseName"`

	// 部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Part *string `json:"Part,omitempty" name:"Part"`

	// 时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Time *string `json:"Time,omitempty" name:"Time"`

	// 时间类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeType *string `json:"TimeType,omitempty" name:"TimeType"`
}

type ClinicalStaging struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type Coord struct {
	// 坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Points []*Point `json:"Points,omitempty" name:"Points"`
}

type Coordinate struct {
	// 左上角x坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	X *int64 `json:"X,omitempty" name:"X"`

	// 左上角y坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Y *int64 `json:"Y,omitempty" name:"Y"`

	// 宽度，单位像素
	// 注意：此字段可能返回 null，表示取不到有效值。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 高度，单位像素
	// 注意：此字段可能返回 null，表示取不到有效值。
	Height *int64 `json:"Height,omitempty" name:"Height"`
}

type CovidItem struct {
	// 采样时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	SampleTime *BaseItem `json:"SampleTime,omitempty" name:"SampleTime"`

	// 检测时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TestTime *BaseItem `json:"TestTime,omitempty" name:"TestTime"`

	// 检测机构
	// 注意：此字段可能返回 null，表示取不到有效值。
	TestOrganization *BaseItem `json:"TestOrganization,omitempty" name:"TestOrganization"`

	// 检测结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	TestResult *BaseItem `json:"TestResult,omitempty" name:"TestResult"`

	// 健康码颜色
	// 注意：此字段可能返回 null，表示取不到有效值。
	CodeColor *BaseItem `json:"CodeColor,omitempty" name:"CodeColor"`
}

type CovidItemsInfo struct {
	// 核酸报告结论
	// 注意：此字段可能返回 null，表示取不到有效值。
	CovidItems []*CovidItem `json:"CovidItems,omitempty" name:"CovidItems"`

	// 版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitempty" name:"Version"`
}

type DeathDateBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 归一化值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Norm *string `json:"Norm,omitempty" name:"Norm"`

	// 单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// 时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *string `json:"Timestamp,omitempty" name:"Timestamp"`

	// 对外输出值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
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

	// 坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitempty" name:"Coords"`
}

type DescInfo struct {
	// 描述段落文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *BaseInfo `json:"Text,omitempty" name:"Text"`

	// 描述段落详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Infos []*DetailInformation `json:"Infos,omitempty" name:"Infos"`
}

type DetailInformation struct {
	// 部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Part *Part `json:"Part,omitempty" name:"Part"`

	// 组织大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	TissueSizes []*Size `json:"TissueSizes,omitempty" name:"TissueSizes"`

	// 结节大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	TuberSizes []*Size `json:"TuberSizes,omitempty" name:"TuberSizes"`

	// 肿瘤大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	CancerSizes []*Size `json:"CancerSizes,omitempty" name:"CancerSizes"`

	// 组织学等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	HistologyLevel *BaseInfo `json:"HistologyLevel,omitempty" name:"HistologyLevel"`

	// 组织学类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	HistologyType *HistologyTypeV2 `json:"HistologyType,omitempty" name:"HistologyType"`

	// 侵犯
	// 注意：此字段可能返回 null，表示取不到有效值。
	Invasive []*InvasiveV2 `json:"Invasive,omitempty" name:"Invasive"`

	// pTNM
	// 注意：此字段可能返回 null，表示取不到有效值。
	PTNM *PTNM `json:"PTNM,omitempty" name:"PTNM"`

	// 浸润深度
	// 注意：此字段可能返回 null，表示取不到有效值。
	InfiltrationDepth *BaseInfo `json:"InfiltrationDepth,omitempty" name:"InfiltrationDepth"`

	// 结节数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TuberNum *BaseInfo `json:"TuberNum,omitempty" name:"TuberNum"`

	// 钙化
	// 注意：此字段可能返回 null，表示取不到有效值。
	Calcification *BaseInfo `json:"Calcification,omitempty" name:"Calcification"`

	// 坏死
	// 注意：此字段可能返回 null，表示取不到有效值。
	Necrosis *BaseInfo `json:"Necrosis,omitempty" name:"Necrosis"`

	// 异形
	// 注意：此字段可能返回 null，表示取不到有效值。
	Abnormity *BaseInfo `json:"Abnormity,omitempty" name:"Abnormity"`

	// 断链
	// 注意：此字段可能返回 null，表示取不到有效值。
	Breaked *BaseInfo `json:"Breaked,omitempty" name:"Breaked"`
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

type DischargeConditionBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 对外输出值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 归一化值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Norm *string `json:"Norm,omitempty" name:"Norm"`
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

type DischargeDiagnosisBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 归一化值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Norm *string `json:"Norm,omitempty" name:"Norm"`

	// 对外输出值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type DischargeInfoBlock struct {
	// 疾病史
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiseaseHistory *DiseaseHistoryBlock `json:"DiseaseHistory,omitempty" name:"DiseaseHistory"`

	// 个人史
	// 注意：此字段可能返回 null，表示取不到有效值。
	PersonalHistory *PersonalHistoryBlock `json:"PersonalHistory,omitempty" name:"PersonalHistory"`

	// 药物史
	// 注意：此字段可能返回 null，表示取不到有效值。
	DrugHistory *DrugHistoryBlock `json:"DrugHistory,omitempty" name:"DrugHistory"`

	// 治疗相关
	// 注意：此字段可能返回 null，表示取不到有效值。
	TreatmentRecord *TreatmentRecordBlock `json:"TreatmentRecord,omitempty" name:"TreatmentRecord"`

	// 文本段落
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParagraphBlock *ParagraphBlock `json:"ParagraphBlock,omitempty" name:"ParagraphBlock"`
}

type DiseaseHistoryBlock struct {
	// 主要病史
	// 注意：此字段可能返回 null，表示取不到有效值。
	MainDiseaseHistory *MainDiseaseHistoryBlock `json:"MainDiseaseHistory,omitempty" name:"MainDiseaseHistory"`

	// 过敏史
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllergyHistory *MainDiseaseHistoryBlock `json:"AllergyHistory,omitempty" name:"AllergyHistory"`

	// 注射史
	// 注意：此字段可能返回 null，表示取不到有效值。
	InfectHistory *MainDiseaseHistoryBlock `json:"InfectHistory,omitempty" name:"InfectHistory"`

	// 手术史
	// 注意：此字段可能返回 null，表示取不到有效值。
	SurgeryHistory *SurgeryHistoryBlock `json:"SurgeryHistory,omitempty" name:"SurgeryHistory"`

	// 输血史
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransfusionHistory *TransfusionHistoryBlock `json:"TransfusionHistory,omitempty" name:"TransfusionHistory"`
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

type DiseasePresentBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 归一化
	// 注意：此字段可能返回 null，表示取不到有效值。
	Norm *string `json:"Norm,omitempty" name:"Norm"`
}

type DosageBlock struct {
	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 单次计量
	// 注意：此字段可能返回 null，表示取不到有效值。
	SingleMeasurement *string `json:"SingleMeasurement,omitempty" name:"SingleMeasurement"`

	// 频次
	// 注意：此字段可能返回 null，表示取不到有效值。
	Frequency *string `json:"Frequency,omitempty" name:"Frequency"`

	// 给药途径
	// 注意：此字段可能返回 null，表示取不到有效值。
	DrugDeliveryRoute *string `json:"DrugDeliveryRoute,omitempty" name:"DrugDeliveryRoute"`
}

type DrugHistoryBlock struct {
	// 药品名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 药物列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DrugList []*DrugListBlock `json:"DrugList,omitempty" name:"DrugList"`

	// 归一化值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type DrugListBlock struct {
	// 通用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	CommonName *string `json:"CommonName,omitempty" name:"CommonName"`

	// 商品名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TradeName *string `json:"TradeName,omitempty" name:"TradeName"`

	// 用法用量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Dosage *DosageBlock `json:"Dosage,omitempty" name:"Dosage"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
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

	// 坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitempty" name:"Coords"`
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

	// 坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitempty" name:"Coords"`
}

type Exame struct {
	// 结论段落
	// 注意：此字段可能返回 null，表示取不到有效值。
	OverView []*ResultInfo `json:"OverView,omitempty" name:"OverView"`

	// 异常与建议段落
	// 注意：此字段可能返回 null，表示取不到有效值。
	Abnormality []*ResultInfo `json:"Abnormality,omitempty" name:"Abnormality"`
}

type EyeChildItem struct {
	// 球镜
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sph []*BaseItem3 `json:"Sph,omitempty" name:"Sph"`

	// 柱镜
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cyl []*BaseItem3 `json:"Cyl,omitempty" name:"Cyl"`

	// 轴位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ax []*BaseItem3 `json:"Ax,omitempty" name:"Ax"`

	// 等效球镜
	// 注意：此字段可能返回 null，表示取不到有效值。
	Se *BaseItem2 `json:"Se,omitempty" name:"Se"`
}

type EyeItem struct {
	// 左眼
	// 注意：此字段可能返回 null，表示取不到有效值。
	Left *EyeChildItem `json:"Left,omitempty" name:"Left"`

	// 右眼
	// 注意：此字段可能返回 null，表示取不到有效值。
	Right *EyeChildItem `json:"Right,omitempty" name:"Right"`

	// 瞳距
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pd *BaseItem2 `json:"Pd,omitempty" name:"Pd"`
}

type EyeItemsInfo struct {
	// 眼科报告
	// 注意：此字段可能返回 null，表示取不到有效值。
	EyeItems *EyeItem `json:"EyeItems,omitempty" name:"EyeItems"`

	// 版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitempty" name:"Version"`
}

type FamilyHistoryBlock struct {
	// 家庭成员
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelativeHistory *RelativeHistoryBlock `json:"RelativeHistory,omitempty" name:"RelativeHistory"`

	// 家族肿瘤史
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelativeCancerHistory *RelativeCancerHistoryBlock `json:"RelativeCancerHistory,omitempty" name:"RelativeCancerHistory"`

	// 家族遗传史
	// 注意：此字段可能返回 null，表示取不到有效值。
	GeneticHistory *GeneticHistoryBlock `json:"GeneticHistory,omitempty" name:"GeneticHistory"`
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

type FertilityHistoryBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *string `json:"State,omitempty" name:"State"`

	// 归一化值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Norm *string `json:"Norm,omitempty" name:"Norm"`

	// 对外输出值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 妊娠次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PregCount *string `json:"PregCount,omitempty" name:"PregCount"`

	// 生产次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProduCount *string `json:"ProduCount,omitempty" name:"ProduCount"`
}

type Fetus struct {
	// 双顶径
	// 注意：此字段可能返回 null，表示取不到有效值。
	BPD *FieldInfo `json:"BPD,omitempty" name:"BPD"`

	// 腹前后径
	// 注意：此字段可能返回 null，表示取不到有效值。
	APTD *FieldInfo `json:"APTD,omitempty" name:"APTD"`

	// 腹左右径
	// 注意：此字段可能返回 null，表示取不到有效值。
	TTD *FieldInfo `json:"TTD,omitempty" name:"TTD"`

	// 头臀径
	// 注意：此字段可能返回 null，表示取不到有效值。
	CRL *FieldInfo `json:"CRL,omitempty" name:"CRL"`

	// 头围
	// 注意：此字段可能返回 null，表示取不到有效值。
	HC *FieldInfo `json:"HC,omitempty" name:"HC"`

	// 腹围
	// 注意：此字段可能返回 null，表示取不到有效值。
	AC *FieldInfo `json:"AC,omitempty" name:"AC"`

	// 股骨长
	// 注意：此字段可能返回 null，表示取不到有效值。
	FL *FieldInfo `json:"FL,omitempty" name:"FL"`

	// 肱骨长
	// 注意：此字段可能返回 null，表示取不到有效值。
	HL *FieldInfo `json:"HL,omitempty" name:"HL"`

	// 胎儿重量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Weight *FieldInfo `json:"Weight,omitempty" name:"Weight"`

	// 颈项透明层
	// 注意：此字段可能返回 null，表示取不到有效值。
	NT *FieldInfo `json:"NT,omitempty" name:"NT"`

	// 脐动脉血流
	// 注意：此字段可能返回 null，表示取不到有效值。
	UmbilicalCord *FieldInfo `json:"UmbilicalCord,omitempty" name:"UmbilicalCord"`

	// 羊水最大深度
	// 注意：此字段可能返回 null，表示取不到有效值。
	WaterDeep *FieldInfo `json:"WaterDeep,omitempty" name:"WaterDeep"`

	// 羊水四象限测量
	// 注意：此字段可能返回 null，表示取不到有效值。
	WaterQuad *FieldInfo `json:"WaterQuad,omitempty" name:"WaterQuad"`

	// 羊水指数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AFI *FieldInfo `json:"AFI,omitempty" name:"AFI"`

	// 胎心
	// 注意：此字段可能返回 null，表示取不到有效值。
	FHR *FieldInfo `json:"FHR,omitempty" name:"FHR"`

	// 胎动
	// 注意：此字段可能返回 null，表示取不到有效值。
	Movement *FieldInfo `json:"Movement,omitempty" name:"Movement"`

	// 胎数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Num *FieldInfo `json:"Num,omitempty" name:"Num"`

	// 胎位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Position *FieldInfo `json:"Position,omitempty" name:"Position"`

	// 是否活胎
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alive *FieldInfo `json:"Alive,omitempty" name:"Alive"`

	// 胎盘位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlacentaLocation *FieldInfo `json:"PlacentaLocation,omitempty" name:"PlacentaLocation"`

	// 胎盘厚度
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlacentaThickness *FieldInfo `json:"PlacentaThickness,omitempty" name:"PlacentaThickness"`

	// 胎盘成熟度
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlacentaGrade *FieldInfo `json:"PlacentaGrade,omitempty" name:"PlacentaGrade"`

	// 妊娠时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	GestationTime *FieldInfo `json:"GestationTime,omitempty" name:"GestationTime"`

	// 妊娠周期
	// 注意：此字段可能返回 null，表示取不到有效值。
	GestationPeriod *FieldInfo `json:"GestationPeriod,omitempty" name:"GestationPeriod"`

	// 绕颈
	// 注意：此字段可能返回 null，表示取不到有效值。
	AroundNeck *FieldInfo `json:"AroundNeck,omitempty" name:"AroundNeck"`

	// 病变
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sym []*FieldInfo `json:"Sym,omitempty" name:"Sym"`

	// 原文内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`
}

type FieldInfo struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 数值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Nums []*NumValue `json:"Nums,omitempty" name:"Nums"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`
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

	// 受伤中毒的外部原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	DamagePoi *BlockInfoV2 `json:"DamagePoi,omitempty" name:"DamagePoi"`

	// 病案首页第二页
	// 注意：此字段可能返回 null，表示取不到有效值。
	Fp2NdItems []*Fp2NdItem `json:"Fp2NdItems,omitempty" name:"Fp2NdItems"`
}

type Fp2NdItem struct {
	// 手术编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *BaseItem `json:"Code,omitempty" name:"Code"`

	// 手术名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *BaseItem `json:"Name,omitempty" name:"Name"`

	// 手术开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *BaseItem `json:"StartTime,omitempty" name:"StartTime"`

	// 手术结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *BaseItem `json:"EndTime,omitempty" name:"EndTime"`

	// 手术等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	Level *BaseItem `json:"Level,omitempty" name:"Level"`

	// 手术类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *BaseItem `json:"Type,omitempty" name:"Type"`

	// 醉愈合方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncisionHealingGrade *BaseItem `json:"IncisionHealingGrade,omitempty" name:"IncisionHealingGrade"`

	// 麻醉方法
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnesthesiaMethod *BaseItem `json:"AnesthesiaMethod,omitempty" name:"AnesthesiaMethod"`
}

type GeneticHistoryBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 遗传列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	GeneticList *string `json:"GeneticList,omitempty" name:"GeneticList"`

	// 对外输出值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
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

type HistologyClass struct {
	// 项目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 索引
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitempty" name:"Index"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 归一化值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 原文对应坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitempty" name:"Coords"`
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

type HistologyTypeV2 struct {
	// 浸润
	// 注意：此字段可能返回 null，表示取不到有效值。
	Infiltration *string `json:"Infiltration,omitempty" name:"Infiltration"`

	// 索引
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitempty" name:"Index"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 归一化后的组织学类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 项目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 原文对应坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitempty" name:"Coords"`
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

type IHCBlock struct {
	// 索引
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitempty" name:"Index"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 具体值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *ValueBlock `json:"Value,omitempty" name:"Value"`

	// 坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitempty" name:"Coords"`
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

type IHCV2 struct {
	// 索引
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitempty" name:"Index"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// ihc归一化
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// ihc详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *Value `json:"Value,omitempty" name:"Value"`

	// 原文对应坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitempty" name:"Coords"`
}

type ImageInfo struct {
	// 图片id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 图片url
	Url *string `json:"Url,omitempty" name:"Url"`

	// 图片base64编码
	Base64 *string `json:"Base64,omitempty" name:"Base64"`
}

// Predefined struct for user
type ImageToClassRequestParams struct {
	// 图片列表，允许传入多张图片，支持传入图片的base64编码，暂不支持图片url
	ImageInfoList []*ImageInfo `json:"ImageInfoList,omitempty" name:"ImageInfoList"`

	// 图片处理参数
	HandleParam *HandleParam `json:"HandleParam,omitempty" name:"HandleParam"`

	// 不填，默认为0
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 后付费的用户类型，新客户传1，老客户可不传或传 0。2022 年 12 月 15 新增了计费项，在此时间之前已经通过商务指定优惠价格的大客户，请不传这个字段或传 0，如果传 1 会导致以前获得的折扣价格失效。在 2022 年 12 月 15 日之后，通过商务指定优惠价格的大客户请传 1。
	UserType *uint64 `json:"UserType,omitempty" name:"UserType"`
}

type ImageToClassRequest struct {
	*tchttp.BaseRequest
	
	// 图片列表，允许传入多张图片，支持传入图片的base64编码，暂不支持图片url
	ImageInfoList []*ImageInfo `json:"ImageInfoList,omitempty" name:"ImageInfoList"`

	// 图片处理参数
	HandleParam *HandleParam `json:"HandleParam,omitempty" name:"HandleParam"`

	// 不填，默认为0
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 后付费的用户类型，新客户传1，老客户可不传或传 0。2022 年 12 月 15 新增了计费项，在此时间之前已经通过商务指定优惠价格的大客户，请不传这个字段或传 0，如果传 1 会导致以前获得的折扣价格失效。在 2022 年 12 月 15 日之后，通过商务指定优惠价格的大客户请传 1。
	UserType *uint64 `json:"UserType,omitempty" name:"UserType"`
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
	delete(f, "UserType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImageToClassRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImageToClassResponseParams struct {
	// 分类结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	TextTypeList []*TextType `json:"TextTypeList,omitempty" name:"TextTypeList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ImageToClassResponse struct {
	*tchttp.BaseResponse
	Response *ImageToClassResponseParams `json:"Response"`
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

// Predefined struct for user
type ImageToObjectRequestParams struct {
	// 图片列表，允许传入多张图片，目前只支持传入图片base64编码，图片url暂不支持
	ImageInfoList []*ImageInfo `json:"ImageInfoList,omitempty" name:"ImageInfoList"`

	// 图片处理参数
	HandleParam *HandleParam `json:"HandleParam,omitempty" name:"HandleParam"`

	// 报告类型，目前支持11（检验报告），12（检查报告），15（病理报告），28（出院报告），29（入院报告），210（门诊病历），212（手术记录），218（诊断证明），363（心电图），27（内窥镜检查），215（处方单），219（免疫接种证明），301（C14呼气试验）。如果不清楚报告类型，可以使用分类引擎，该字段传0（同时IsUsedClassify字段必须为True，否则无法输出结果）
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 是否使用分类引擎，当不确定报告类型时，可以使用收费的报告分类引擎服务。若该字段为 False，则 Type 字段不能为 0，否则无法输出结果。
	// 注意：当 IsUsedClassify 为True 时，表示使用收费的报告分类服务，将会产生额外的费用，具体收费标准参见 [购买指南的产品价格](https://cloud.tencent.com/document/product/1314/54264)。
	IsUsedClassify *bool `json:"IsUsedClassify,omitempty" name:"IsUsedClassify"`

	// 后付费的用户类型，新客户传1，老客户可不传或传 0。2022 年 12 月 15 新增了计费项，在此时间之前已经通过商务指定优惠价格的大客户，请不传这个字段或传 0，如果传 1 会导致以前获得的折扣价格失效。在 2022 年 12 月 15 日之后，通过商务指定优惠价格的大客户请传 1。
	UserType *int64 `json:"UserType,omitempty" name:"UserType"`

	// 可选。用于指定不同报告使用的结构化引擎版本，不同版本返回的JSON 数据结果不兼容。若不指定版本号，就默认用旧的版本号。
	// （1）检验报告 11，默认使用 V2，最高支持 V3。
	// （2）病理报告 15，默认使用 V1，最高支持 V2。
	// （3）入院记录29、出院记录 28、病理记录 216、病程记录 217、门诊记录 210，默认使用 V1，最高支持 V2。
	ReportTypeVersion []*ReportTypeVersion `json:"ReportTypeVersion,omitempty" name:"ReportTypeVersion"`
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

	// 后付费的用户类型，新客户传1，老客户可不传或传 0。2022 年 12 月 15 新增了计费项，在此时间之前已经通过商务指定优惠价格的大客户，请不传这个字段或传 0，如果传 1 会导致以前获得的折扣价格失效。在 2022 年 12 月 15 日之后，通过商务指定优惠价格的大客户请传 1。
	UserType *int64 `json:"UserType,omitempty" name:"UserType"`

	// 可选。用于指定不同报告使用的结构化引擎版本，不同版本返回的JSON 数据结果不兼容。若不指定版本号，就默认用旧的版本号。
	// （1）检验报告 11，默认使用 V2，最高支持 V3。
	// （2）病理报告 15，默认使用 V1，最高支持 V2。
	// （3）入院记录29、出院记录 28、病理记录 216、病程记录 217、门诊记录 210，默认使用 V1，最高支持 V2。
	ReportTypeVersion []*ReportTypeVersion `json:"ReportTypeVersion,omitempty" name:"ReportTypeVersion"`
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
	delete(f, "UserType")
	delete(f, "ReportTypeVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImageToObjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImageToObjectResponseParams struct {
	// 报告结构化结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Template *Template `json:"Template,omitempty" name:"Template"`

	// 多级分类结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	TextTypeList []*TextType `json:"TextTypeList,omitempty" name:"TextTypeList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ImageToObjectResponse struct {
	*tchttp.BaseResponse
	Response *ImageToObjectResponseParams `json:"Response"`
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

type ImmunohistochemistryBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 免疫组化详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value []*IHCBlock `json:"Value,omitempty" name:"Value"`
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

	// 指标项ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 指标项坐标位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords *Coordinate `json:"Coords,omitempty" name:"Coords"`

	// 推测结果是否异常
	// 注意：此字段可能返回 null，表示取不到有效值。
	InferNormal *string `json:"InferNormal,omitempty" name:"InferNormal"`
}

type IndicatorItemV2 struct {
	// 项目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Item *BaseItem `json:"Item,omitempty" name:"Item"`

	// 英文编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *BaseItem `json:"Code,omitempty" name:"Code"`

	// 结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *BaseItem `json:"Result,omitempty" name:"Result"`

	// 单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *BaseItem `json:"Unit,omitempty" name:"Unit"`

	// 参考范围
	// 注意：此字段可能返回 null，表示取不到有效值。
	Range *BaseItem `json:"Range,omitempty" name:"Range"`

	// 上下箭头
	// 注意：此字段可能返回 null，表示取不到有效值。
	Arrow *BaseItem `json:"Arrow,omitempty" name:"Arrow"`

	// 检测方法
	// 注意：此字段可能返回 null，表示取不到有效值。
	Method *BaseItem `json:"Method,omitempty" name:"Method"`

	// 结果是否异常
	// 注意：此字段可能返回 null，表示取不到有效值。
	Normal *bool `json:"Normal,omitempty" name:"Normal"`

	// ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 序号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Order *int64 `json:"Order,omitempty" name:"Order"`

	// 推测结果是否异常
	// 注意：此字段可能返回 null，表示取不到有效值。
	InferNormal *string `json:"InferNormal,omitempty" name:"InferNormal"`
}

type IndicatorV3 struct {
	// 检验报告V3结论
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableIndictors []*TableIndicators `json:"TableIndictors,omitempty" name:"TableIndictors"`

	// 版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitempty" name:"Version"`
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

type InvasiveV2 struct {
	// 索引
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitempty" name:"Index"`

	// 部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Part *Part `json:"Part,omitempty" name:"Part"`

	// 阴性或阳性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Positive *string `json:"Positive,omitempty" name:"Positive"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 原文对应坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitempty" name:"Coords"`
}

type IssueInfo struct {
	// 编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertNumber *string `json:"CertNumber,omitempty" name:"CertNumber"`

	// 签发机构
	// 注意：此字段可能返回 null，表示取不到有效值。
	IssuedAuthority *string `json:"IssuedAuthority,omitempty" name:"IssuedAuthority"`

	// 签发日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	IssuedDate *string `json:"IssuedDate,omitempty" name:"IssuedDate"`
}

type LastMenstrualPeriodBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 归一化值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Norm *string `json:"Norm,omitempty" name:"Norm"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *string `json:"Timestamp,omitempty" name:"Timestamp"`

	// 单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// 对外输出值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
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

type LymphNode struct {
	// 项目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 索引
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitempty" name:"Index"`

	// 部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Part *Part `json:"Part,omitempty" name:"Part"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 总数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 转移数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferNum *int64 `json:"TransferNum,omitempty" name:"TransferNum"`

	// 淋巴结大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sizes []*int64 `json:"Sizes,omitempty" name:"Sizes"`

	// 原文对应坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitempty" name:"Coords"`
}

type LymphTotal struct {
	// 项目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 转移数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferNum *string `json:"TransferNum,omitempty" name:"TransferNum"`

	// 总数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *string `json:"Total,omitempty" name:"Total"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 索引
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitempty" name:"Index"`

	// 原文对应坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitempty" name:"Coords"`
}

type MainDiseaseHistoryBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *bool `json:"State,omitempty" name:"State"`

	// 对外输出值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 否定列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Neglist *NeglistBlock `json:"Neglist,omitempty" name:"Neglist"`

	// 肯定列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Poslist *PoslistBlock `json:"Poslist,omitempty" name:"Poslist"`
}

type Maternity struct {
	// 描述部分
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *MaternityDesc `json:"Desc,omitempty" name:"Desc"`

	// 结论部分
	// 注意：此字段可能返回 null，表示取不到有效值。
	Summary *MaternitySummary `json:"Summary,omitempty" name:"Summary"`

	// 报告原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrText *string `json:"OcrText,omitempty" name:"OcrText"`
}

type MaternityDesc struct {
	// 胎儿数据结构
	// 注意：此字段可能返回 null，表示取不到有效值。
	Fetus []*Fetus `json:"Fetus,omitempty" name:"Fetus"`

	// 胎儿数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	FetusNum *FieldInfo `json:"FetusNum,omitempty" name:"FetusNum"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitempty" name:"Text"`

	// 坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitempty" name:"Coords"`
}

type MaternitySummary struct {
	// 胎儿数据结构
	// 注意：此字段可能返回 null，表示取不到有效值。
	Fetus []*Fetus `json:"Fetus,omitempty" name:"Fetus"`

	// 胎儿数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	FetusNum *FieldInfo `json:"FetusNum,omitempty" name:"FetusNum"`

	// 病变
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sym []*FieldInfo `json:"Sym,omitempty" name:"Sym"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitempty" name:"Text"`

	// 坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitempty" name:"Coords"`
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

type MenstrualFlowBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 对外输出值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type MenstrualHistoryBlock struct {
	// 末次月经
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastMenstrualPeriod *LastMenstrualPeriodBlock `json:"LastMenstrualPeriod,omitempty" name:"LastMenstrualPeriod"`

	// 月经量
	// 注意：此字段可能返回 null，表示取不到有效值。
	MenstrualFlow *MenstrualFlowBlock `json:"MenstrualFlow,omitempty" name:"MenstrualFlow"`

	// 初潮时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	MenarcheAge *LastMenstrualPeriodBlock `json:"MenarcheAge,omitempty" name:"MenarcheAge"`

	// 是否绝经
	// 注意：此字段可能返回 null，表示取不到有效值。
	MenstruationOrNot *MenstruationOrNotBlock `json:"MenstruationOrNot,omitempty" name:"MenstruationOrNot"`

	// 月经周期
	// 注意：此字段可能返回 null，表示取不到有效值。
	MenstrualCycles *LastMenstrualPeriodBlock `json:"MenstrualCycles,omitempty" name:"MenstrualCycles"`

	// 月经经期
	// 注意：此字段可能返回 null，表示取不到有效值。
	MenstrualPeriod *MenstrualPeriodBlock `json:"MenstrualPeriod,omitempty" name:"MenstrualPeriod"`
}

type MenstrualHistoryDetailBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *string `json:"State,omitempty" name:"State"`

	// 归一化值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Norm *string `json:"Norm,omitempty" name:"Norm"`

	// 时间类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeType *string `json:"TimeType,omitempty" name:"TimeType"`

	// 时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *string `json:"Timestamp,omitempty" name:"Timestamp"`

	// 单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// 对外输出值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
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

type MenstrualPeriodBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 归一化值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Norm *string `json:"Norm,omitempty" name:"Norm"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *string `json:"Timestamp,omitempty" name:"Timestamp"`

	// 单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// 对外输出值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type MenstruationOrNotBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 归一化值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Norm *string `json:"Norm,omitempty" name:"Norm"`

	// 时间类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeType *string `json:"TimeType,omitempty" name:"TimeType"`

	// 时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *string `json:"Timestamp,omitempty" name:"Timestamp"`

	// 对外输出值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type Molecular struct {
	// 索引
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitempty" name:"Index"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 基因名称标注化
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分子病理详细信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *MolecularValue `json:"Value,omitempty" name:"Value"`

	// 原文对应坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitempty" name:"Coords"`
}

type MolecularValue struct {
	// 外显子
	// 注意：此字段可能返回 null，表示取不到有效值。
	Exon *string `json:"Exon,omitempty" name:"Exon"`

	// 点位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Position *string `json:"Position,omitempty" name:"Position"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 阳性或阴性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Positive *string `json:"Positive,omitempty" name:"Positive"`

	// 基因名称原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`
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

type NeglistBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type NeonatalInfo struct {
	// 新生儿名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	NeonatalName *string `json:"NeonatalName,omitempty" name:"NeonatalName"`

	// 新生儿性别
	// 注意：此字段可能返回 null，表示取不到有效值。
	NeonatalGender *string `json:"NeonatalGender,omitempty" name:"NeonatalGender"`

	// 出生身长
	// 注意：此字段可能返回 null，表示取不到有效值。
	BirthLength *string `json:"BirthLength,omitempty" name:"BirthLength"`

	// 出生体重
	// 注意：此字段可能返回 null，表示取不到有效值。
	BirthWeight *string `json:"BirthWeight,omitempty" name:"BirthWeight"`

	// 出生孕周
	// 注意：此字段可能返回 null，表示取不到有效值。
	GestationalAge *string `json:"GestationalAge,omitempty" name:"GestationalAge"`

	// 出生时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	BirthTime *string `json:"BirthTime,omitempty" name:"BirthTime"`

	// 出生地点
	// 注意：此字段可能返回 null，表示取不到有效值。
	BirthPlace *string `json:"BirthPlace,omitempty" name:"BirthPlace"`

	// 医疗机构
	// 注意：此字段可能返回 null，表示取不到有效值。
	MedicalInstitutions *string `json:"MedicalInstitutions,omitempty" name:"MedicalInstitutions"`
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

	// 部位详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	PartDetail *PartDesc `json:"PartDetail,omitempty" name:"PartDetail"`
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

	// 归一化值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Impl *string `json:"Impl,omitempty" name:"Impl"`
}

type NumValue struct {
	// 数值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Num *string `json:"Num,omitempty" name:"Num"`

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

type ObstetricalHistoryBlock struct {
	// 婚姻史
	// 注意：此字段可能返回 null，表示取不到有效值。
	MarriageHistory *MenstrualHistoryDetailBlock `json:"MarriageHistory,omitempty" name:"MarriageHistory"`

	// 婚育史
	// 注意：此字段可能返回 null，表示取不到有效值。
	FertilityHistory *FertilityHistoryBlock `json:"FertilityHistory,omitempty" name:"FertilityHistory"`
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

	// 透声度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Transparent *BlockInfo `json:"Transparent,omitempty" name:"Transparent"`

	// MRI ADC
	// 注意：此字段可能返回 null，表示取不到有效值。
	MriAdc *BlockInfo `json:"MriAdc,omitempty" name:"MriAdc"`

	// MRI DWI
	// 注意：此字段可能返回 null，表示取不到有效值。
	MriDwi *BlockInfo `json:"MriDwi,omitempty" name:"MriDwi"`

	// MRI T1信号
	// 注意：此字段可能返回 null，表示取不到有效值。
	MriT1 *BlockInfo `json:"MriT1,omitempty" name:"MriT1"`

	// MRI T2信号
	// 注意：此字段可能返回 null，表示取不到有效值。
	MriT2 *BlockInfo `json:"MriT2,omitempty" name:"MriT2"`

	// CT HU值
	// 注意：此字段可能返回 null，表示取不到有效值。
	CtHu *BlockInfo `json:"CtHu,omitempty" name:"CtHu"`

	// SUmax值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suvmax *BlockInfo `json:"Suvmax,omitempty" name:"Suvmax"`

	// 代谢情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	Metabolism *BlockInfo `json:"Metabolism,omitempty" name:"Metabolism"`

	// 放射性摄取
	// 注意：此字段可能返回 null，表示取不到有效值。
	RadioactiveUptake *BlockInfo `json:"RadioactiveUptake,omitempty" name:"RadioactiveUptake"`

	// 淋巴结情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	LymphEnlargement *BlockInfo `json:"LymphEnlargement,omitempty" name:"LymphEnlargement"`

	// 影像特征
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageFeature *BlockInfo `json:"ImageFeature,omitempty" name:"ImageFeature"`

	// 导管
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duct *BlockInfo `json:"Duct,omitempty" name:"Duct"`

	// 趋势
	// 注意：此字段可能返回 null，表示取不到有效值。
	Trend *BlockInfo `json:"Trend,omitempty" name:"Trend"`

	// 手术情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operation *BlockInfo `json:"Operation,omitempty" name:"Operation"`

	// 器官在报告图片中的坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitempty" name:"Coords"`
}

type OtherInfo struct {
	// 麻醉方法
	// 注意：此字段可能返回 null，表示取不到有效值。
	Anesthesia *SurgeryAttr `json:"Anesthesia,omitempty" name:"Anesthesia"`

	// 术中出血
	// 注意：此字段可能返回 null，表示取不到有效值。
	BloodLoss *SurgeryAttr `json:"BloodLoss,omitempty" name:"BloodLoss"`

	// 输血
	// 注意：此字段可能返回 null，表示取不到有效值。
	BloodTransfusion *SurgeryAttr `json:"BloodTransfusion,omitempty" name:"BloodTransfusion"`

	// 手术用时
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *SurgeryAttr `json:"Duration,omitempty" name:"Duration"`

	// 手术开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *SurgeryAttr `json:"EndTime,omitempty" name:"EndTime"`

	// 手术结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *SurgeryAttr `json:"StartTime,omitempty" name:"StartTime"`
}

type PTNM struct {
	// 项目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 索引
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitempty" name:"Index"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 归一化值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// pT
	// 注意：此字段可能返回 null，表示取不到有效值。
	PT *string `json:"PT,omitempty" name:"PT"`

	// pN
	// 注意：此字段可能返回 null，表示取不到有效值。
	PN *string `json:"PN,omitempty" name:"PN"`

	// pM
	// 注意：此字段可能返回 null，表示取不到有效值。
	PM *string `json:"PM,omitempty" name:"PM"`

	// 原文对应坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitempty" name:"Coords"`
}

type PTNMBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// PTNM分期
	// 注意：此字段可能返回 null，表示取不到有效值。
	PTNMM *string `json:"PTNMM,omitempty" name:"PTNMM"`

	// PTNM分期
	// 注意：此字段可能返回 null，表示取不到有效值。
	PTNMN *string `json:"PTNMN,omitempty" name:"PTNMN"`

	// PTNM分期
	// 注意：此字段可能返回 null，表示取不到有效值。
	PTNMT *string `json:"PTNMT,omitempty" name:"PTNMT"`
}

type ParagraphBlock struct {
	// 切口愈合情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncisionHealingText *string `json:"IncisionHealingText,omitempty" name:"IncisionHealingText"`

	// 辅助检查
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuxiliaryExaminationText *string `json:"AuxiliaryExaminationText,omitempty" name:"AuxiliaryExaminationText"`

	// 特殊检查
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecialExamText *string `json:"SpecialExamText,omitempty" name:"SpecialExamText"`

	// 门诊诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutpatientDiagnosisText *string `json:"OutpatientDiagnosisText,omitempty" name:"OutpatientDiagnosisText"`

	// 入院情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdmissionConditionText *string `json:"AdmissionConditionText,omitempty" name:"AdmissionConditionText"`

	// 诊疗经过
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckAndTreatmentProcessText *string `json:"CheckAndTreatmentProcessText,omitempty" name:"CheckAndTreatmentProcessText"`

	// 体征
	// 注意：此字段可能返回 null，表示取不到有效值。
	SymptomsAndSignsText *string `json:"SymptomsAndSignsText,omitempty" name:"SymptomsAndSignsText"`

	// 出院医嘱
	// 注意：此字段可能返回 null，表示取不到有效值。
	DischargeInstructionsText *string `json:"DischargeInstructionsText,omitempty" name:"DischargeInstructionsText"`

	// 入院诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdmissionDiagnosisText *string `json:"AdmissionDiagnosisText,omitempty" name:"AdmissionDiagnosisText"`

	// 手术情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	SurgeryConditionText *string `json:"SurgeryConditionText,omitempty" name:"SurgeryConditionText"`

	// 病理诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	PathologicalDiagnosisText *string `json:"PathologicalDiagnosisText,omitempty" name:"PathologicalDiagnosisText"`

	// 出院情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	DischargeConditionText *string `json:"DischargeConditionText,omitempty" name:"DischargeConditionText"`

	// 检查记录
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckRecordText *string `json:"CheckRecordText,omitempty" name:"CheckRecordText"`

	// 主诉
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChiefComplaintText *string `json:"ChiefComplaintText,omitempty" name:"ChiefComplaintText"`

	// 出院诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	DischargeDiagnosisText *string `json:"DischargeDiagnosisText,omitempty" name:"DischargeDiagnosisText"`

	// 既往史
	// 注意：此字段可能返回 null，表示取不到有效值。
	MainDiseaseHistoryText *string `json:"MainDiseaseHistoryText,omitempty" name:"MainDiseaseHistoryText"`

	// 现病史
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiseasePresentText *string `json:"DiseasePresentText,omitempty" name:"DiseasePresentText"`

	// 个人史
	// 注意：此字段可能返回 null，表示取不到有效值。
	PersonalHistoryText *string `json:"PersonalHistoryText,omitempty" name:"PersonalHistoryText"`

	// 月经史
	// 注意：此字段可能返回 null，表示取不到有效值。
	MenstruallHistoryText *string `json:"MenstruallHistoryText,omitempty" name:"MenstruallHistoryText"`

	// 婚育史
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObstericalHistoryText *string `json:"ObstericalHistoryText,omitempty" name:"ObstericalHistoryText"`

	// 家族史
	// 注意：此字段可能返回 null，表示取不到有效值。
	FamilyHistoryText *string `json:"FamilyHistoryText,omitempty" name:"FamilyHistoryText"`

	// 过敏史
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllergyHistoryText *string `json:"AllergyHistoryText,omitempty" name:"AllergyHistoryText"`

	// 病史信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiseaseHistoryText *string `json:"DiseaseHistoryText,omitempty" name:"DiseaseHistoryText"`

	// 其它诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	OtherDiagnosisText *string `json:"OtherDiagnosisText,omitempty" name:"OtherDiagnosisText"`

	// 体格检查
	// 注意：此字段可能返回 null，表示取不到有效值。
	BodyExaminationText *string `json:"BodyExaminationText,omitempty" name:"BodyExaminationText"`

	// 专科检查
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecialistExaminationText *string `json:"SpecialistExaminationText,omitempty" name:"SpecialistExaminationText"`

	// 治疗结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	TreatmentResultText *string `json:"TreatmentResultText,omitempty" name:"TreatmentResultText"`
}

type ParentInfo struct {
	// 名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 年龄
	// 注意：此字段可能返回 null，表示取不到有效值。
	Age *string `json:"Age,omitempty" name:"Age"`

	// 证件号
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdCard *string `json:"IdCard,omitempty" name:"IdCard"`

	// 民族
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ethnicity *string `json:"Ethnicity,omitempty" name:"Ethnicity"`

	// 国籍
	// 注意：此字段可能返回 null，表示取不到有效值。
	Nationality *string `json:"Nationality,omitempty" name:"Nationality"`

	// 地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Address *string `json:"Address,omitempty" name:"Address"`
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

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValueBrief *string `json:"ValueBrief,omitempty" name:"ValueBrief"`
}

type PartDesc struct {
	// 主要部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	MainDir *string `json:"MainDir,omitempty" name:"MainDir"`

	// 部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Part *string `json:"Part,omitempty" name:"Part"`

	// 次要部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecondaryDir *string `json:"SecondaryDir,omitempty" name:"SecondaryDir"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`
}

type PathologicalDiagnosisBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 病理详细
	// 注意：此字段可能返回 null，表示取不到有效值。
	Detail []*PathologicalDiagnosisDetailBlock `json:"Detail,omitempty" name:"Detail"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type PathologicalDiagnosisDetailBlock struct {
	// 部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Part *string `json:"Part,omitempty" name:"Part"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	HistologicalType *string `json:"HistologicalType,omitempty" name:"HistologicalType"`

	// 等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	HistologicalGrade *string `json:"HistologicalGrade,omitempty" name:"HistologicalGrade"`
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

type PathologyV2 struct {
	// 报告类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	PathologicalReportType *Report `json:"PathologicalReportType,omitempty" name:"PathologicalReportType"`

	// 描述段落
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *DescInfo `json:"Desc,omitempty" name:"Desc"`

	// 诊断结论
	// 注意：此字段可能返回 null，表示取不到有效值。
	Summary *SummaryInfo `json:"Summary,omitempty" name:"Summary"`

	// 报告全文
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportText *string `json:"ReportText,omitempty" name:"ReportText"`

	// 淋巴结总计转移信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	LymphTotal []*LymphTotal `json:"LymphTotal,omitempty" name:"LymphTotal"`

	// 单淋巴结转移信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	LymphNodes []*LymphNode `json:"LymphNodes,omitempty" name:"LymphNodes"`

	// ihc信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ihc []*IHCV2 `json:"Ihc,omitempty" name:"Ihc"`

	// 临床诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	Clinical *BaseInfo `json:"Clinical,omitempty" name:"Clinical"`

	// 是否癌前病变
	// 注意：此字段可能返回 null，表示取不到有效值。
	Precancer *HistologyClass `json:"Precancer,omitempty" name:"Precancer"`

	// 是否恶性肿瘤
	// 注意：此字段可能返回 null，表示取不到有效值。
	Malignant *HistologyClass `json:"Malignant,omitempty" name:"Malignant"`

	// 是否良性肿瘤
	// 注意：此字段可能返回 null，表示取不到有效值。
	Benigntumor *HistologyClass `json:"Benigntumor,omitempty" name:"Benigntumor"`

	// 送检材料
	// 注意：此字段可能返回 null，表示取不到有效值。
	SampleType *BaseInfo `json:"SampleType,omitempty" name:"SampleType"`

	// 淋巴结大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	LymphSize []*Size `json:"LymphSize,omitempty" name:"LymphSize"`

	// 分子病理
	// 注意：此字段可能返回 null，表示取不到有效值。
	Molecular []*Molecular `json:"Molecular,omitempty" name:"Molecular"`
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

	// 床号
	// 注意：此字段可能返回 null，表示取不到有效值。
	BedNo *string `json:"BedNo,omitempty" name:"BedNo"`
}

type PersonalHistoryBlock struct {
	// 出生地
	// 注意：此字段可能返回 null，表示取不到有效值。
	BirthPlace *BirthPlaceBlock `json:"BirthPlace,omitempty" name:"BirthPlace"`

	// 居住地
	// 注意：此字段可能返回 null，表示取不到有效值。
	LivePlace *BirthPlaceBlock `json:"LivePlace,omitempty" name:"LivePlace"`

	// 职业
	// 注意：此字段可能返回 null，表示取不到有效值。
	Job *BirthPlaceBlock `json:"Job,omitempty" name:"Job"`

	// 吸烟
	// 注意：此字段可能返回 null，表示取不到有效值。
	SmokeHistory *SmokeHistoryBlock `json:"SmokeHistory,omitempty" name:"SmokeHistory"`

	// 喝酒
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlcoholicHistory *SmokeHistoryBlock `json:"AlcoholicHistory,omitempty" name:"AlcoholicHistory"`

	// 月经史
	// 注意：此字段可能返回 null，表示取不到有效值。
	MenstrualHistory *MenstrualHistoryBlock `json:"MenstrualHistory,omitempty" name:"MenstrualHistory"`

	// 婚姻-生育史
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObstericalHistory *ObstetricalHistoryBlock `json:"ObstericalHistory,omitempty" name:"ObstericalHistory"`

	// 家族史
	// 注意：此字段可能返回 null，表示取不到有效值。
	FamilyHistory *FamilyHistoryBlock `json:"FamilyHistory,omitempty" name:"FamilyHistory"`
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

type Point struct {
	// x坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	X *int64 `json:"X,omitempty" name:"X"`

	// y坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Y *int64 `json:"Y,omitempty" name:"Y"`
}

type PoslistBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type Prescription struct {
	// 药品列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	MedicineList []*Medicine `json:"MedicineList,omitempty" name:"MedicineList"`
}

type Rectangle struct {
	// x坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	X *int64 `json:"X,omitempty" name:"X"`

	// y坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Y *int64 `json:"Y,omitempty" name:"Y"`

	// 宽
	// 注意：此字段可能返回 null，表示取不到有效值。
	W *int64 `json:"W,omitempty" name:"W"`

	// 高
	// 注意：此字段可能返回 null，表示取不到有效值。
	H *int64 `json:"H,omitempty" name:"H"`
}

type RelapseDateBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 疾病名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiseaseName *string `json:"DiseaseName,omitempty" name:"DiseaseName"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 归一化值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Norm *string `json:"Norm,omitempty" name:"Norm"`

	// 单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// 时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *string `json:"Timestamp,omitempty" name:"Timestamp"`

	// 对外输出值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type RelativeCancerHistoryBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 肿瘤史列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelativeCancerList *string `json:"RelativeCancerList,omitempty" name:"RelativeCancerList"`

	// 对外输出值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type RelativeHistoryBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 成员列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Detail []*RelativeHistoryDetailBlock `json:"Detail,omitempty" name:"Detail"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 对外输出值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type RelativeHistoryDetailBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 关系
	// 注意：此字段可能返回 null，表示取不到有效值。
	Relation *string `json:"Relation,omitempty" name:"Relation"`

	// 死亡时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeOfDeath *string `json:"TimeOfDeath,omitempty" name:"TimeOfDeath"`

	// 时间类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeType *string `json:"TimeType,omitempty" name:"TimeType"`
}

type Report struct {
	// 项目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 索引
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitempty" name:"Index"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 报告类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 原文对应坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitempty" name:"Coords"`
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

	// 检查项目
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckItem *string `json:"CheckItem,omitempty" name:"CheckItem"`

	// 检查方法
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckMethod *string `json:"CheckMethod,omitempty" name:"CheckMethod"`

	// 诊断时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiagnoseTime *string `json:"DiagnoseTime,omitempty" name:"DiagnoseTime"`

	// 体检号
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthCheckupNum *string `json:"HealthCheckupNum,omitempty" name:"HealthCheckupNum"`

	// 其它时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	OtherTime *string `json:"OtherTime,omitempty" name:"OtherTime"`

	// 打印时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrintTime *string `json:"PrintTime,omitempty" name:"PrintTime"`

	// 未归类时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Times []*Time `json:"Times,omitempty" name:"Times"`
}

type ReportTypeVersion struct {
	// 检验报告
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportType *int64 `json:"ReportType,omitempty" name:"ReportType"`

	// 版本2
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *int64 `json:"Version,omitempty" name:"Version"`
}

type ResultInfo struct {
	// 段落文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *BaseInfo `json:"Text,omitempty" name:"Text"`

	// 结论详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*BaseInfo `json:"Items,omitempty" name:"Items"`
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

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`
}

type SmokeHistoryBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 时间单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 时间归一化
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeNorm *string `json:"TimeNorm,omitempty" name:"TimeNorm"`

	// 吸烟量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Amount *string `json:"Amount,omitempty" name:"Amount"`

	// 戒烟状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	QuitState *bool `json:"QuitState,omitempty" name:"QuitState"`

	// 是否吸烟
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *bool `json:"State,omitempty" name:"State"`

	// 对外输出值
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

	// 坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitempty" name:"Coords"`
}

type SummaryInfo struct {
	// 诊断结论文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *BaseInfo `json:"Text,omitempty" name:"Text"`

	// 诊断结论详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Infos []*DetailInformation `json:"Infos,omitempty" name:"Infos"`
}

type Surgery struct {
	// 手术史
	// 注意：此字段可能返回 null，表示取不到有效值。
	SurgeryHistory *SurgeryHistory `json:"SurgeryHistory,omitempty" name:"SurgeryHistory"`

	// 其他信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	OtherInfo *OtherInfo `json:"OtherInfo,omitempty" name:"OtherInfo"`
}

type SurgeryAttr struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type SurgeryConditionBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 手术历史
	// 注意：此字段可能返回 null，表示取不到有效值。
	SurgeryList []*SurgeryListBlock `json:"SurgeryList,omitempty" name:"SurgeryList"`

	// 对外输出值
	// 
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

type SurgeryHistoryBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 手术列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Surgerylist []*SurgeryListBlock `json:"Surgerylist,omitempty" name:"Surgerylist"`
}

type SurgeryListBlock struct {
	// 时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Time *string `json:"Time,omitempty" name:"Time"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeType *string `json:"TimeType,omitempty" name:"TimeType"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name []*string `json:"Name,omitempty" name:"Name"`

	// 部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Part *string `json:"Part,omitempty" name:"Part"`
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

	// 坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitempty" name:"Coords"`
}

type TableIndicators struct {
	// 项目列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Indicators []*IndicatorItemV2 `json:"Indicators,omitempty" name:"Indicators"`

	// 采样标本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sample *BaseItem `json:"Sample,omitempty" name:"Sample"`
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

	// 疫苗接种凭证
	// 注意：此字段可能返回 null，表示取不到有效值。
	VaccineCertificate *VaccineCertificate `json:"VaccineCertificate,omitempty" name:"VaccineCertificate"`

	// OCR文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrText *string `json:"OcrText,omitempty" name:"OcrText"`

	// OCR拼接后文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrResult *string `json:"OcrResult,omitempty" name:"OcrResult"`

	// 报告类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportTypeDesc *string `json:"ReportTypeDesc,omitempty" name:"ReportTypeDesc"`

	// 病理报告v2
	// 注意：此字段可能返回 null，表示取不到有效值。
	PathologyV2 *PathologyV2 `json:"PathologyV2,omitempty" name:"PathologyV2"`

	// 碳14尿素呼气试验
	// 注意：此字段可能返回 null，表示取不到有效值。
	C14 *Indicator `json:"C14,omitempty" name:"C14"`

	// 体检结论
	// 注意：此字段可能返回 null，表示取不到有效值。
	Exame *Exame `json:"Exame,omitempty" name:"Exame"`

	// 出院报告v2，入院报告v2，门诊病历v2
	// 注意：此字段可能返回 null，表示取不到有效值。
	MedDocV2 *DischargeInfoBlock `json:"MedDocV2,omitempty" name:"MedDocV2"`

	// 检验报告v3
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndicatorV3 *IndicatorV3 `json:"IndicatorV3,omitempty" name:"IndicatorV3"`

	// 核酸报告
	// 注意：此字段可能返回 null，表示取不到有效值。
	Covid *CovidItemsInfo `json:"Covid,omitempty" name:"Covid"`

	// 孕产报告
	// 注意：此字段可能返回 null，表示取不到有效值。
	Maternity *Maternity `json:"Maternity,omitempty" name:"Maternity"`

	// 眼科报告
	// 注意：此字段可能返回 null，表示取不到有效值。
	Eye *EyeItemsInfo `json:"Eye,omitempty" name:"Eye"`

	// 出生证明
	// 注意：此字段可能返回 null，表示取不到有效值。
	BirthCert *BirthCert `json:"BirthCert,omitempty" name:"BirthCert"`

	// 时间轴
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timeline *TimelineInformation `json:"Timeline,omitempty" name:"Timeline"`
}

// Predefined struct for user
type TextToClassRequestParams struct {
	// 报告文本
	Text *string `json:"Text,omitempty" name:"Text"`

	// 后付费的用户类型，新客户传1，老客户可不传或传 0。2022 年 12 月 15 新增了计费项，在此时间之前已经通过商务指定优惠价格的大客户，请不传这个字段或传 0，如果传 1 会导致以前获得的折扣价格失效。在 2022 年 12 月 15 日之后，通过商务指定优惠价格的大客户请传 1。
	UserType *uint64 `json:"UserType,omitempty" name:"UserType"`
}

type TextToClassRequest struct {
	*tchttp.BaseRequest
	
	// 报告文本
	Text *string `json:"Text,omitempty" name:"Text"`

	// 后付费的用户类型，新客户传1，老客户可不传或传 0。2022 年 12 月 15 新增了计费项，在此时间之前已经通过商务指定优惠价格的大客户，请不传这个字段或传 0，如果传 1 会导致以前获得的折扣价格失效。在 2022 年 12 月 15 日之后，通过商务指定优惠价格的大客户请传 1。
	UserType *uint64 `json:"UserType,omitempty" name:"UserType"`
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
	delete(f, "UserType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TextToClassRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextToClassResponseParams struct {
	// 分类结果
	TextTypeList []*TextType `json:"TextTypeList,omitempty" name:"TextTypeList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type TextToClassResponse struct {
	*tchttp.BaseResponse
	Response *TextToClassResponseParams `json:"Response"`
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

// Predefined struct for user
type TextToObjectRequestParams struct {
	// 报告文本
	Text *string `json:"Text,omitempty" name:"Text"`

	// 报告类型，目前支持12（检查报告），15（病理报告），28（出院报告），29（入院报告），210（门诊病历），212（手术记录），218（诊断证明），363（心电图），27（内窥镜检查），215（处方单），219（免疫接种证明），301（C14呼气试验）。如果不清楚报告类型，可以使用分类引擎，该字段传0（同时IsUsedClassify字段必须为True，否则无法输出结果）
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 是否使用分类引擎，当不确定报告类型时，可以使用收费的报告分类引擎服务。若该字段为False，则Type字段不能为0，否则无法输出结果。
	// 注意：当 IsUsedClassify 为True 时，表示使用收费的报告分类服务，将会产生额外的费用，具体收费标准参见 [购买指南的产品价格](https://cloud.tencent.com/document/product/1314/54264)。
	IsUsedClassify *bool `json:"IsUsedClassify,omitempty" name:"IsUsedClassify"`

	// 后付费的用户类型，新客户传1，老客户可不传或传 0。2022 年 12 月 15 新增了计费项，在此时间之前已经通过商务指定优惠价格的大客户，请不传这个字段或传 0，如果传 1 会导致以前获得的折扣价格失效。在 2022 年 12 月 15 日之后，通过商务指定优惠价格的大客户请传 1。
	UserType *uint64 `json:"UserType,omitempty" name:"UserType"`

	// 可选。用于指定不同报告使用的结构化引擎版本，不同版本返回的JSON 数据结果不兼容。若不指定版本号，就默认用旧的版本号。
	// （1）检验报告 11，默认使用 V2，最高支持 V3。
	// （2）病理报告 15，默认使用 V1，最高支持 V2。
	// （3）入院记录29、出院记录 28、病理记录 216、病程记录 217、门诊记录 210，默认使用 V1，最高支持 V2。
	ReportTypeVersion []*ReportTypeVersion `json:"ReportTypeVersion,omitempty" name:"ReportTypeVersion"`
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

	// 后付费的用户类型，新客户传1，老客户可不传或传 0。2022 年 12 月 15 新增了计费项，在此时间之前已经通过商务指定优惠价格的大客户，请不传这个字段或传 0，如果传 1 会导致以前获得的折扣价格失效。在 2022 年 12 月 15 日之后，通过商务指定优惠价格的大客户请传 1。
	UserType *uint64 `json:"UserType,omitempty" name:"UserType"`

	// 可选。用于指定不同报告使用的结构化引擎版本，不同版本返回的JSON 数据结果不兼容。若不指定版本号，就默认用旧的版本号。
	// （1）检验报告 11，默认使用 V2，最高支持 V3。
	// （2）病理报告 15，默认使用 V1，最高支持 V2。
	// （3）入院记录29、出院记录 28、病理记录 216、病程记录 217、门诊记录 210，默认使用 V1，最高支持 V2。
	ReportTypeVersion []*ReportTypeVersion `json:"ReportTypeVersion,omitempty" name:"ReportTypeVersion"`
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
	delete(f, "UserType")
	delete(f, "ReportTypeVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TextToObjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextToObjectResponseParams struct {
	// 报告结构化结果
	Template *Template `json:"Template,omitempty" name:"Template"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type TextToObjectResponse struct {
	*tchttp.BaseResponse
	Response *TextToObjectResponseParams `json:"Response"`
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

type Time struct {
	// 具体时间类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 时间值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type TimelineEvent struct {
	// 事件类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 原文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 事件子类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubType *string `json:"SubType,omitempty" name:"SubType"`

	// 事件发生时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Time *string `json:"Time,omitempty" name:"Time"`

	// 事件值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 位置坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rectangle *Rectangle `json:"Rectangle,omitempty" name:"Rectangle"`

	// 事件发生地点
	// 注意：此字段可能返回 null，表示取不到有效值。
	Place *string `json:"Place,omitempty" name:"Place"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type TimelineInformation struct {
	// 时间轴
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timeline []*TimelineEvent `json:"Timeline,omitempty" name:"Timeline"`
}

type TransfusionHistoryBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *bool `json:"State,omitempty" name:"State"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
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

type TreatmentRecordBlock struct {
	// 免疫组化
	// 注意：此字段可能返回 null，表示取不到有效值。
	Immunohistochemistry *ImmunohistochemistryBlock `json:"Immunohistochemistry,omitempty" name:"Immunohistochemistry"`

	// 主诉
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChiefComplaint *ChiefComplaintBlock `json:"ChiefComplaint,omitempty" name:"ChiefComplaint"`

	// 入院情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdmissionCondition *AdmissionConditionBlock `json:"AdmissionCondition,omitempty" name:"AdmissionCondition"`

	// 查体
	// 注意：此字段可能返回 null，表示取不到有效值。
	BodyExamination *BodyExaminationBlock `json:"BodyExamination,omitempty" name:"BodyExamination"`

	// 入院诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdmissionDiagnosis *AdmissionDiagnosisBlock `json:"AdmissionDiagnosis,omitempty" name:"AdmissionDiagnosis"`

	// 入院中医诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdmissionTraditionalDiagnosis *AdmissionDiagnosisBlock `json:"AdmissionTraditionalDiagnosis,omitempty" name:"AdmissionTraditionalDiagnosis"`

	// 入院西医诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdmissionModernDiagnosis *AdmissionDiagnosisBlock `json:"AdmissionModernDiagnosis,omitempty" name:"AdmissionModernDiagnosis"`

	// 病理诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	PathologicalDiagnosis *PathologicalDiagnosisBlock `json:"PathologicalDiagnosis,omitempty" name:"PathologicalDiagnosis"`

	// 现病史
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiseasePresent *DiseasePresentBlock `json:"DiseasePresent,omitempty" name:"DiseasePresent"`

	// 体征
	// 注意：此字段可能返回 null，表示取不到有效值。
	SymptomsAndSigns *DiseasePresentBlock `json:"SymptomsAndSigns,omitempty" name:"SymptomsAndSigns"`

	// 辅助检查
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuxiliaryExamination *DiseasePresentBlock `json:"AuxiliaryExamination,omitempty" name:"AuxiliaryExamination"`

	// 特殊检查
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecialistExamination *DiseasePresentBlock `json:"SpecialistExamination,omitempty" name:"SpecialistExamination"`

	// 精神检查
	// 注意：此字段可能返回 null，表示取不到有效值。
	MentalExamination *DiseasePresentBlock `json:"MentalExamination,omitempty" name:"MentalExamination"`

	// 检查记录
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckRecord *DiseasePresentBlock `json:"CheckRecord,omitempty" name:"CheckRecord"`

	// 检查结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	InspectResult *DiseasePresentBlock `json:"InspectResult,omitempty" name:"InspectResult"`

	// 治疗经过
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckAndTreatmentProcess *DiseasePresentBlock `json:"CheckAndTreatmentProcess,omitempty" name:"CheckAndTreatmentProcess"`

	// 手术经过
	// 注意：此字段可能返回 null，表示取不到有效值。
	SurgeryCondition *SurgeryConditionBlock `json:"SurgeryCondition,omitempty" name:"SurgeryCondition"`

	// 切口愈合
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncisionHealing *DiseasePresentBlock `json:"IncisionHealing,omitempty" name:"IncisionHealing"`

	// 出院诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	DischargeDiagnosis *DischargeDiagnosisBlock `json:"DischargeDiagnosis,omitempty" name:"DischargeDiagnosis"`

	// 出院中医诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	DischargeTraditionalDiagnosis *DiseasePresentBlock `json:"DischargeTraditionalDiagnosis,omitempty" name:"DischargeTraditionalDiagnosis"`

	// 出院西医诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	DischargeModernDiagnosis *DischargeDiagnosisBlock `json:"DischargeModernDiagnosis,omitempty" name:"DischargeModernDiagnosis"`

	// 出院情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	DischargeCondition *DischargeConditionBlock `json:"DischargeCondition,omitempty" name:"DischargeCondition"`

	// 出院医嘱
	// 注意：此字段可能返回 null，表示取不到有效值。
	DischargeInstructions *DiseasePresentBlock `json:"DischargeInstructions,omitempty" name:"DischargeInstructions"`

	// 治疗建议
	// 注意：此字段可能返回 null，表示取不到有效值。
	TreatmentSuggestion *DiseasePresentBlock `json:"TreatmentSuggestion,omitempty" name:"TreatmentSuggestion"`

	// 随访
	// 注意：此字段可能返回 null，表示取不到有效值。
	FollowUpRequirements *DiseasePresentBlock `json:"FollowUpRequirements,omitempty" name:"FollowUpRequirements"`

	// 治疗情况变化
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConditionChanges *DiseasePresentBlock `json:"ConditionChanges,omitempty" name:"ConditionChanges"`

	// 收缩压
	// 注意：此字段可能返回 null，表示取不到有效值。
	PulmonaryArterySystolicPressure *DiseasePresentBlock `json:"PulmonaryArterySystolicPressure,omitempty" name:"PulmonaryArterySystolicPressure"`

	// bclc分期
	// 注意：此字段可能返回 null，表示取不到有效值。
	BCLC *DiseasePresentBlock `json:"BCLC,omitempty" name:"BCLC"`

	// PTNM分期
	// 注意：此字段可能返回 null，表示取不到有效值。
	PTNM *PTNMBlock `json:"PTNM,omitempty" name:"PTNM"`

	// ECOG评分
	// 注意：此字段可能返回 null，表示取不到有效值。
	ECOG *DiseasePresentBlock `json:"ECOG,omitempty" name:"ECOG"`

	// NRS评分
	// 注意：此字段可能返回 null，表示取不到有效值。
	NRS *DiseasePresentBlock `json:"NRS,omitempty" name:"NRS"`

	// kps评分
	// 注意：此字段可能返回 null，表示取不到有效值。
	KPS *DiseasePresentBlock `json:"KPS,omitempty" name:"KPS"`

	// 癌症分期
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cancerstaging *ClinicalStaging `json:"Cancerstaging,omitempty" name:"Cancerstaging"`

	// 死亡时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeathDate *DeathDateBlock `json:"DeathDate,omitempty" name:"DeathDate"`

	// 复发日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelapseDate *RelapseDateBlock `json:"RelapseDate,omitempty" name:"RelapseDate"`

	// 观察日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObservationDays *DeathDateBlock `json:"ObservationDays,omitempty" name:"ObservationDays"`

	// 切口愈合情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncisionHealingText *string `json:"IncisionHealingText,omitempty" name:"IncisionHealingText"`

	// 辅助检查
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuxiliaryExaminationText *string `json:"AuxiliaryExaminationText,omitempty" name:"AuxiliaryExaminationText"`

	// 特殊检查
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecialExamText *string `json:"SpecialExamText,omitempty" name:"SpecialExamText"`

	// 门诊诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutpatientDiagnosisText *string `json:"OutpatientDiagnosisText,omitempty" name:"OutpatientDiagnosisText"`

	// 入院情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdmissionConditionText *string `json:"AdmissionConditionText,omitempty" name:"AdmissionConditionText"`

	// 诊疗经过
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckAndTreatmentProcessText *string `json:"CheckAndTreatmentProcessText,omitempty" name:"CheckAndTreatmentProcessText"`

	// 体征
	// 注意：此字段可能返回 null，表示取不到有效值。
	SymptomsAndSignsText *string `json:"SymptomsAndSignsText,omitempty" name:"SymptomsAndSignsText"`

	// 出院医嘱
	// 注意：此字段可能返回 null，表示取不到有效值。
	DischargeInstructionsText *string `json:"DischargeInstructionsText,omitempty" name:"DischargeInstructionsText"`

	// 入院诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdmissionDiagnosisText *string `json:"AdmissionDiagnosisText,omitempty" name:"AdmissionDiagnosisText"`

	// 手术情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	SurgeryConditionText *string `json:"SurgeryConditionText,omitempty" name:"SurgeryConditionText"`

	// 病理诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	PathologicalDiagnosisText *string `json:"PathologicalDiagnosisText,omitempty" name:"PathologicalDiagnosisText"`

	// 出院情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	DischargeConditionText *string `json:"DischargeConditionText,omitempty" name:"DischargeConditionText"`

	// 检查记录
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckRecordText *string `json:"CheckRecordText,omitempty" name:"CheckRecordText"`

	// 主诉
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChiefComplaintText *string `json:"ChiefComplaintText,omitempty" name:"ChiefComplaintText"`

	// 出院诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	DischargeDiagnosisText *string `json:"DischargeDiagnosisText,omitempty" name:"DischargeDiagnosisText"`
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

	// 透声度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Transparent *BlockInfo `json:"Transparent,omitempty" name:"Transparent"`

	// MRI ADC
	// 注意：此字段可能返回 null，表示取不到有效值。
	MriAdc *BlockInfo `json:"MriAdc,omitempty" name:"MriAdc"`

	// MRI DWI
	// 注意：此字段可能返回 null，表示取不到有效值。
	MriDwi *BlockInfo `json:"MriDwi,omitempty" name:"MriDwi"`

	// MRI T1信号
	// 注意：此字段可能返回 null，表示取不到有效值。
	MriT1 *BlockInfo `json:"MriT1,omitempty" name:"MriT1"`

	// MRI T2信号
	// 注意：此字段可能返回 null，表示取不到有效值。
	MriT2 *BlockInfo `json:"MriT2,omitempty" name:"MriT2"`

	// CT HU值
	// 注意：此字段可能返回 null，表示取不到有效值。
	CtHu *BlockInfo `json:"CtHu,omitempty" name:"CtHu"`

	// SUmax值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suvmax *BlockInfo `json:"Suvmax,omitempty" name:"Suvmax"`

	// 代谢情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	Metabolism *BlockInfo `json:"Metabolism,omitempty" name:"Metabolism"`

	// 放射性摄取
	// 注意：此字段可能返回 null，表示取不到有效值。
	RadioactiveUptake *BlockInfo `json:"RadioactiveUptake,omitempty" name:"RadioactiveUptake"`

	// 病变
	// 注意：此字段可能返回 null，表示取不到有效值。
	SymDesc *BlockInfo `json:"SymDesc,omitempty" name:"SymDesc"`

	// 影像特征
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageFeature *BlockInfo `json:"ImageFeature,omitempty" name:"ImageFeature"`

	// 在报告图片中的坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitempty" name:"Coords"`
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

	// 疫苗批号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Lot *string `json:"Lot,omitempty" name:"Lot"`
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

type ValueBlock struct {
	// 等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	Grade *string `json:"Grade,omitempty" name:"Grade"`

	// 百分比
	// 注意：此字段可能返回 null，表示取不到有效值。
	Percent []*float64 `json:"Percent,omitempty" name:"Percent"`

	// 阳性阴性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Positive *string `json:"Positive,omitempty" name:"Positive"`
}