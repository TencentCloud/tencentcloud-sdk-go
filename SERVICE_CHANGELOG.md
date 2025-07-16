# Release v1.0.1208

## 暴露面管理服务(ctem) 版本：2023-11-28

### 第 4 次发布

发布时间：2025-07-16 01:29:54

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeFakeApps](https://cloud.tencent.com/document/api/1755/121355)
* [DescribeFakeMiniPrograms](https://cloud.tencent.com/document/api/1755/121354)
* [DescribeFakeWebsites](https://cloud.tencent.com/document/api/1755/121353)
* [DescribeFakeWechatOfficials](https://cloud.tencent.com/document/api/1755/121352)

新增数据结构：

* [DisplayFakeApp](https://cloud.tencent.com/document/api/1755/120320#DisplayFakeApp)
* [DisplayFakeMiniProgram](https://cloud.tencent.com/document/api/1755/120320#DisplayFakeMiniProgram)
* [DisplayFakeWebsite](https://cloud.tencent.com/document/api/1755/120320#DisplayFakeWebsite)
* [DisplayFakeWechatOfficial](https://cloud.tencent.com/document/api/1755/120320#DisplayFakeWechatOfficial)



## 主机安全(cwp) 版本：2018-02-28

### 第 151 次发布

发布时间：2025-07-17 01:25:19

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [Machine](https://cloud.tencent.com/document/api/296/19867#Machine)

	* 新增成员：AgentVersion




## TDSQL-C MySQL 版(cynosdb) 版本：2019-01-07

### 第 138 次发布

发布时间：2025-07-17 01:29:44

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [AddClusterSlaveZone](https://cloud.tencent.com/document/api/1003/79555)

	* 新增入参：SemiSyncTimeout

* [ModifyClusterSlaveZone](https://cloud.tencent.com/document/api/1003/79554)

	* 新增入参：SemiSyncTimeout


修改数据结构：

* [SlaveZoneAttrItem](https://cloud.tencent.com/document/api/1003/48097#SlaveZoneAttrItem)

	* 新增成员：SemiSyncTimeout




## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 235 次发布

发布时间：2025-07-17 01:43:04

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateBatchInformationExtractionTask](https://cloud.tencent.com/document/api/1323/121392)
* [DescribeInformationExtractionTask](https://cloud.tencent.com/document/api/1323/121391)

新增数据结构：

* [ExtractionField](https://cloud.tencent.com/document/api/1323/70369#ExtractionField)



## 云直播CSS(live) 版本：2018-08-01

### 第 154 次发布

发布时间：2025-07-17 01:58:05

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateAuditKeywords](https://cloud.tencent.com/document/api/267/121248)

	* 新增入参：Keywords, LibId

	* 新增出参：KeywordIds, DupInfos

* [DeleteAuditKeywords](https://cloud.tencent.com/document/api/267/121247)

	* 新增入参：KeywordIds, LibId

	* 新增出参：SuccessCount, Infos

* [DescribeAuditKeywords](https://cloud.tencent.com/document/api/267/121246)

	* 新增入参：Offset, Limit, LibId, Content, Labels

	* 新增出参：Total, Infos


新增数据结构：

* [AuditKeyword](https://cloud.tencent.com/document/api/267/20474#AuditKeyword)
* [AuditKeywordDeleteDetail](https://cloud.tencent.com/document/api/267/20474#AuditKeywordDeleteDetail)
* [AuditKeywordInfo](https://cloud.tencent.com/document/api/267/20474#AuditKeywordInfo)

### 第 153 次发布

发布时间：2025-07-16 02:34:29

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateLiveTranscodeTemplate](https://cloud.tencent.com/document/api/267/32646)

	* 新增入参：IsAdaptiveBitRate, AdaptiveChildren

* [ModifyLiveTranscodeTemplate](https://cloud.tencent.com/document/api/267/32640)

	* 新增入参：IsAdaptiveBitRate, AdaptiveChildren




## 知识引擎原子能力(lkeap) 版本：2024-05-22

### 第 13 次发布

发布时间：2025-07-17 02:01:03

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [GetReconstructDocumentResult](https://cloud.tencent.com/document/api/1772/115342)

	* 新增出参：Usage

* [ReconstructDocumentSSE](https://cloud.tencent.com/document/api/1772/115340)

	* 新增出参：FailPageNum, SuccessPageNum


修改数据结构：

* [CreateReconstructDocumentFlowConfig](https://cloud.tencent.com/document/api/1772/115364#CreateReconstructDocumentFlowConfig)

	* 新增成员：IgnoreFailedPage

* [CreateSplitDocumentFlowConfig](https://cloud.tencent.com/document/api/1772/115364#CreateSplitDocumentFlowConfig)

	* 新增成员：IgnoreFailedPage

* [DocumentUsage](https://cloud.tencent.com/document/api/1772/115364#DocumentUsage)

	* 新增成员：SuccessPageNum, FailPageNum

* [ReconstructDocumentSSEConfig](https://cloud.tencent.com/document/api/1772/115364#ReconstructDocumentSSEConfig)

	* 新增成员：IgnoreFailedPage




## 腾讯云可观测平台(monitor) 版本：2023-06-16



## 腾讯云可观测平台(monitor) 版本：2018-07-24

### 第 140 次发布

发布时间：2025-07-17 02:03:42

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [ModifyConditionsTemplateRequestEventCondition](https://cloud.tencent.com/document/api/248/30354#ModifyConditionsTemplateRequestEventCondition)

	* 新增成员：MetricName, Description

	* <font color="#dd0000">**修改成员**：</font>AlarmNotifyPeriod, AlarmNotifyType, EventID




## 媒体处理(mps) 版本：2019-06-12

### 第 133 次发布

发布时间：2025-07-17 02:05:23

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [EvaluationTaskInput](https://cloud.tencent.com/document/api/862/37615#EvaluationTaskInput)


### 第 132 次发布

发布时间：2025-07-16 02:47:54

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateMediaEvaluation](https://cloud.tencent.com/document/api/862/121357)

新增数据结构：

* [EvaluationMediaInputInfo](https://cloud.tencent.com/document/api/862/37615#EvaluationMediaInputInfo)
* [EvaluationTaskInput](https://cloud.tencent.com/document/api/862/37615#EvaluationTaskInput)
* [EvaluationTemplateInputInfo](https://cloud.tencent.com/document/api/862/37615#EvaluationTemplateInputInfo)



## 文字识别(ocr) 版本：2018-11-19

### 第 202 次发布

发布时间：2025-07-16 02:54:00

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [HKIDCardOCR](https://cloud.tencent.com/document/api/866/46919)

	* 新增出参：SmallHeadImage, WindowEmbeddedText




## 腾讯健康组学平台(omics) 版本：2022-11-28

### 第 22 次发布

发布时间：2025-07-17 02:10:02

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [ClusterOption](https://cloud.tencent.com/document/api/1643/89100#ClusterOption)

	* 新增成员：AutoUpgradeClusterLevel




## 边缘安全加速平台(teo) 版本：2022-09-01

### 第 111 次发布

发布时间：2025-07-15 15:58:34

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateMultiPathGateway](https://cloud.tencent.com/document/api/1552/121343)
* [CreateMultiPathGatewayLine](https://cloud.tencent.com/document/api/1552/121342)
* [CreateMultiPathGatewaySecretKey](https://cloud.tencent.com/document/api/1552/121341)
* [DeleteMultiPathGateway](https://cloud.tencent.com/document/api/1552/121340)
* [DeleteMultiPathGatewayLine](https://cloud.tencent.com/document/api/1552/121339)
* [DescribeMultiPathGateway](https://cloud.tencent.com/document/api/1552/121338)
* [DescribeMultiPathGatewayLine](https://cloud.tencent.com/document/api/1552/121337)
* [DescribeMultiPathGatewayRegions](https://cloud.tencent.com/document/api/1552/121336)
* [DescribeMultiPathGatewaySecretKey](https://cloud.tencent.com/document/api/1552/121335)
* [DescribeMultiPathGateways](https://cloud.tencent.com/document/api/1552/121334)
* [ModifyMultiPathGateway](https://cloud.tencent.com/document/api/1552/121333)
* [ModifyMultiPathGatewayLine](https://cloud.tencent.com/document/api/1552/121332)
* [ModifyMultiPathGatewaySecretKey](https://cloud.tencent.com/document/api/1552/121331)
* [RefreshMultiPathGatewaySecretKey](https://cloud.tencent.com/document/api/1552/121330)

新增数据结构：

* [GatewayRegion](https://cloud.tencent.com/document/api/1552/80721#GatewayRegion)
* [MultiPathGateway](https://cloud.tencent.com/document/api/1552/80721#MultiPathGateway)
* [MultiPathGatewayLine](https://cloud.tencent.com/document/api/1552/80721#MultiPathGatewayLine)



## 边缘安全加速平台(teo) 版本：2022-01-06



## TI-ONE 训练平台(tione) 版本：2021-11-11

### 第 87 次发布

发布时间：2025-07-17 02:31:57

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyModelService](https://cloud.tencent.com/document/api/851/83228)

	* 新增入参：ResourceGroupId




## TI-ONE 训练平台(tione) 版本：2019-10-22



