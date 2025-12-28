# Release v1.3.22

## 应用性能监控(apm) 版本：2021-06-22

### 第 51 次发布

发布时间：2025-12-29 01:08:56

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyApmApplicationConfig](https://cloud.tencent.com/document/api/1463/125072)

	* 新增入参：EnableDesensitizationRule, DesensitizationRule


修改数据结构：

* [ApmAppConfig](https://cloud.tencent.com/document/api/1463/64927#ApmAppConfig)

	* 新增成员：EnableDesensitizationRule, DesensitizationRule

* [ApmApplicationConfigView](https://cloud.tencent.com/document/api/1463/64927#ApmApplicationConfigView)

	* 新增成员：EnableDesensitizationRule, DesensitizationRule




## 云数据库 MySQL(cdb) 版本：2017-03-20

### 第 211 次发布

发布时间：2025-12-29 01:11:42

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateCloneInstance](https://cloud.tencent.com/document/api/236/50424)

	* 新增入参：MasterZone




## 内容分发网络 CDN(cdn) 版本：2018-06-06

### 第 142 次发布

发布时间：2025-12-29 01:12:28

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除接口**：</font>

* CreateDiagnoseUrl
* DescribeDiagnoseReport
* ListDiagnoseReport

<font color="#dd0000">**删除数据结构**：</font>

* ClientInfo
* DiagnoseData
* DiagnoseInfo
* DiagnoseList
* DiagnoseUnit



## 云防火墙(cfw) 版本：2019-09-04

### 第 88 次发布

发布时间：2025-12-29 01:13:28

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除接口**：</font>

* DeleteAllAccessControlRule



## 负载均衡(clb) 版本：2018-03-17

### 第 144 次发布

发布时间：2025-12-29 01:14:29

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [RenewLoadBalancers](https://cloud.tencent.com/document/api/214/126911)



## 媒体处理(mps) 版本：2019-06-12

### 第 167 次发布

发布时间：2025-12-29 01:36:04

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [AiAnalysisTaskCutoutInput](https://cloud.tencent.com/document/api/862/37615#AiAnalysisTaskCutoutInput)
* [AiAnalysisTaskCutoutOutput](https://cloud.tencent.com/document/api/862/37615#AiAnalysisTaskCutoutOutput)
* [AiAnalysisTaskCutoutResult](https://cloud.tencent.com/document/api/862/37615#AiAnalysisTaskCutoutResult)
* [AiAnalysisTaskReelInput](https://cloud.tencent.com/document/api/862/37615#AiAnalysisTaskReelInput)
* [AiAnalysisTaskReelOutput](https://cloud.tencent.com/document/api/862/37615#AiAnalysisTaskReelOutput)
* [AiAnalysisTaskReelResult](https://cloud.tencent.com/document/api/862/37615#AiAnalysisTaskReelResult)
* [SelectingSubtitleAreasConfig](https://cloud.tencent.com/document/api/862/37615#SelectingSubtitleAreasConfig)

修改数据结构：

* [AiAnalysisResult](https://cloud.tencent.com/document/api/862/37615#AiAnalysisResult)

	* 新增成员：CutoutTask, ReelTask

* [RawSmartSubtitleParameter](https://cloud.tencent.com/document/api/862/37615#RawSmartSubtitleParameter)

	* 新增成员：SelectingSubtitleAreasConfig

* [SmartSubtitleTemplateItem](https://cloud.tencent.com/document/api/862/37615#SmartSubtitleTemplateItem)

	* 新增成员：SelectingSubtitleAreasConfig




## 文字识别(ocr) 版本：2018-11-19

### 第 227 次发布

发布时间：2025-12-26 17:17:41

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [LicensePlateOCR](https://cloud.tencent.com/document/api/866/36211)

	* 新增出参：LicensePlateCategory


修改数据结构：

* [LicensePlateInfo](https://cloud.tencent.com/document/api/866/33527#LicensePlateInfo)

	* 新增成员：LicensePlateCategory




## 云点播(vod) 版本：2024-07-18



## 云点播(vod) 版本：2018-07-17

### 第 215 次发布

发布时间：2025-12-29 02:30:15

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateAigcVideoTask](https://cloud.tencent.com/document/api/266/126239)

	* 新增入参：InputRegion


修改数据结构：

* [AigcVideoTaskInputFileInfo](https://cloud.tencent.com/document/api/266/31773#AigcVideoTaskInputFileInfo)

	* 新增成员：ReferenceType, ObjectId, VoiceId




## 私有网络(vpc) 版本：2017-03-12

### 第 283 次发布

发布时间：2025-12-29 02:33:24

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [NatGateway](https://cloud.tencent.com/document/api/215/15824#NatGateway)

	* 新增成员：StrictSnatMode, AutoScaling, ICMPProxy, PublicAddressAffinity




