# Release v1.3.88

## 商业智能分析 BI(bi) 版本：2022-01-05

### 第 34 次发布

发布时间：2026-04-28 01:11:02

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateAuthApiKey](https://cloud.tencent.com/document/api/590/131087)
* [DeleteAuthApiKey](https://cloud.tencent.com/document/api/590/131086)
* [DescribeAuthApiKeyInfo](https://cloud.tencent.com/document/api/590/131085)
* [DescribeAuthApiKeyList](https://cloud.tencent.com/document/api/590/131082)
* [ModifyAuthApiKey](https://cloud.tencent.com/document/api/590/131084)

新增数据结构：

* [ApiKeyAuthApplyVO](https://cloud.tencent.com/document/api/590/73726#ApiKeyAuthApplyVO)
* [ApiKeyAuthApplyVOList](https://cloud.tencent.com/document/api/590/73726#ApiKeyAuthApplyVOList)



## 腾讯云数据仓库TCHouse-C(cdwch) 版本：2020-09-15

### 第 42 次发布

发布时间：2026-04-28 01:14:56

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [InstanceInfo](https://cloud.tencent.com/document/api/1299/83429#InstanceInfo)

	* 新增成员：HttpsEnabled




## 暴露面管理服务(ctem) 版本：2023-11-28

### 第 16 次发布

发布时间：2026-04-28 01:20:17

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateSubDomain](https://cloud.tencent.com/document/api/1755/125838)

	* 新增入参：DnsType, DnsValue


修改数据结构：

* [DisplaySubDomain](https://cloud.tencent.com/document/api/1755/120320#DisplaySubDomain)

	* 新增成员：DnsType, DnsValue




## TDSQL-C MySQL 版(cynosdb) 版本：2019-01-07

### 第 163 次发布

发布时间：2026-04-28 01:23:29

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeBackupConfig](https://cloud.tencent.com/document/api/1003/48094)

	* 新增出参：SparseBackupConfig

* [ModifyBackupConfig](https://cloud.tencent.com/document/api/1003/48090)

	* 新增入参：SparseBackupConfig


新增数据结构：

* [MonthDay](https://cloud.tencent.com/document/api/1003/48097#MonthDay)
* [SparseBackupConfig](https://cloud.tencent.com/document/api/1003/48097#SparseBackupConfig)
* [SparseBackupConfigInfo](https://cloud.tencent.com/document/api/1003/48097#SparseBackupConfigInfo)
* [SparseBackupConfigRsp](https://cloud.tencent.com/document/api/1003/48097#SparseBackupConfigRsp)
* [SparsePeriodTime](https://cloud.tencent.com/document/api/1003/48097#SparsePeriodTime)

修改数据结构：

* [BackupFileInfo](https://cloud.tencent.com/document/api/1003/48097#BackupFileInfo)

	* 新增成员：BackupPeriodStrategy

* [DeliverSummary](https://cloud.tencent.com/document/api/1003/48097#DeliverSummary)

	* 新增成员：DeliverError




## 数据湖计算 DLC(dlc) 版本：2021-01-25

### 第 157 次发布

发布时间：2026-04-28 01:25:59

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [AnalysisTaskResults](https://cloud.tencent.com/document/api/1342/53778#AnalysisTaskResults)

	* 新增成员：ShuffleWriteBytesSum

* [TaskFullRespInfo](https://cloud.tencent.com/document/api/1342/53778#TaskFullRespInfo)

	* 新增成员：ShuffleWriteBytesSum, ActiveCore




## iOA 零信任安全管理系统(ioa) 版本：2022-06-01

### 第 33 次发布

发布时间：2026-04-28 01:35:01

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeLocalAccounts](https://cloud.tencent.com/document/api/1092/107710)

	* 新增入参：DomainInstanceId




## 云直播CSS(live) 版本：2018-08-01

### 第 174 次发布

发布时间：2026-04-28 01:40:05

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateLivePadTemplate](https://cloud.tencent.com/document/api/267/93529)

	* 新增入参：TriggerCondition

* [ModifyLivePadTemplate](https://cloud.tencent.com/document/api/267/93523)

	* 新增入参：TriggerCondition




## 腾讯云智能体开发平台(lke) 版本：2023-11-30

### 第 87 次发布

发布时间：2026-04-28 01:41:08

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeSearchStatsGraph](https://cloud.tencent.com/document/api/1759/111064)

* [ListDoc](https://cloud.tencent.com/document/api/1759/105064)

	* 新增入参：UpdateTime

* [ListQA](https://cloud.tencent.com/document/api/1759/105028)

	* 新增入参：CreateTime, UpdateTime


新增数据结构：

* [TimeRange](https://cloud.tencent.com/document/api/1759/105104#TimeRange)

修改数据结构：

* [QAQuery](https://cloud.tencent.com/document/api/1759/105104#QAQuery)

	* 新增成员：CreateTime, UpdateTime




## 云数据库 MongoDB(mongodb) 版本：2019-07-25

### 第 68 次发布

发布时间：2026-04-28 01:46:08

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateDBInstance](https://cloud.tencent.com/document/api/240/38571)

	* 新增入参：CpuCore

* [CreateDBInstanceHour](https://cloud.tencent.com/document/api/240/38570)

	* 新增入参：CpuCore

* [InquirePriceCreateDBInstances](https://cloud.tencent.com/document/api/240/43666)

	* 新增入参：ReadonlyNodeNum, Cpu

* [InquirePriceModifyDBInstanceSpec](https://cloud.tencent.com/document/api/240/43665)

	* 新增入参：Cpu

* [ModifyDBInstanceSpec](https://cloud.tencent.com/document/api/240/38565)

	* 新增入参：Cpu, MachineCode




## 云数据库 MongoDB(mongodb) 版本：2018-04-08



## 腾讯云可观测平台(monitor) 版本：2023-06-16



## 腾讯云可观测平台(monitor) 版本：2018-07-24

### 第 155 次发布

发布时间：2026-04-28 01:47:13

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CheckAddressByPrometheus](https://cloud.tencent.com/document/api/248/131088)



## 媒体处理(mps) 版本：2019-06-12

### 第 198 次发布

发布时间：2026-04-28 01:49:48

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateSmartSubtitleTemplate](https://cloud.tencent.com/document/api/862/117004)

	* 新增入参：SpeakerMode, SpeakerLabel

* [ModifySmartSubtitleTemplate](https://cloud.tencent.com/document/api/862/117001)

	* 新增入参：SpeakerMode, SpeakerLabel


修改数据结构：

* [AnimatedGraphicTaskInput](https://cloud.tencent.com/document/api/862/37615#AnimatedGraphicTaskInput)

	* 新增成员：ExtInfo

* [ImageSpriteTaskInput](https://cloud.tencent.com/document/api/862/37615#ImageSpriteTaskInput)

	* 新增成员：ExtInfo

* [RawSmartSubtitleParameter](https://cloud.tencent.com/document/api/862/37615#RawSmartSubtitleParameter)

	* 新增成员：SpeakerMode, SpeakerLabel

* [SampleSnapshotTaskInput](https://cloud.tencent.com/document/api/862/37615#SampleSnapshotTaskInput)

	* 新增成员：ExtInfo

* [SmartSubtitleTaskResultInput](https://cloud.tencent.com/document/api/862/37615#SmartSubtitleTaskResultInput)

	* 新增成员：UserExtPara

* [SmartSubtitleTemplateItem](https://cloud.tencent.com/document/api/862/37615#SmartSubtitleTemplateItem)

	* 新增成员：SubtitleEmbedId, SpeakerMode, SpeakerLabel

* [SnapshotByTimeOffsetTaskInput](https://cloud.tencent.com/document/api/862/37615#SnapshotByTimeOffsetTaskInput)

	* 新增成员：ExtInfo




## 云数据库 PostgreSQL(postgres) 版本：2017-03-12

### 第 65 次发布

发布时间：2026-04-28 02:00:13

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CloseAuditService](https://cloud.tencent.com/document/api/409/131097)
* [CreateAuditLogFile](https://cloud.tencent.com/document/api/409/131096)
* [DeleteAuditLogFile](https://cloud.tencent.com/document/api/409/131095)
* [DescribeAuditInstanceList](https://cloud.tencent.com/document/api/409/131094)
* [DescribeAuditLogFiles](https://cloud.tencent.com/document/api/409/131093)
* [DescribeAuditLogs](https://cloud.tencent.com/document/api/409/131092)
* [ModifyAuditService](https://cloud.tencent.com/document/api/409/131091)
* [OpenAuditService](https://cloud.tencent.com/document/api/409/131090)

新增数据结构：

* [AuditInstanceInfo](https://cloud.tencent.com/document/api/409/16778#AuditInstanceInfo)
* [AuditLog](https://cloud.tencent.com/document/api/409/16778#AuditLog)
* [AuditLogFile](https://cloud.tencent.com/document/api/409/16778#AuditLogFile)
* [AuditLogFilter](https://cloud.tencent.com/document/api/409/16778#AuditLogFilter)
* [DeliverSummary](https://cloud.tencent.com/document/api/409/16778#DeliverSummary)
* [LogInstanceInfo](https://cloud.tencent.com/document/api/409/16778#LogInstanceInfo)



## TI-ONE 训练平台(tione) 版本：2021-11-11

### 第 113 次发布

发布时间：2026-04-28 02:32:45

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeBillingResourceGroup](https://cloud.tencent.com/document/api/851/101969)

	* 新增入参：TiProjectId

* [DescribeBillingResourceGroupAttachedWorkspaces](https://cloud.tencent.com/document/api/851/130953)

	* 新增入参：TiProjectId

* [DescribeBillingResourceGroups](https://cloud.tencent.com/document/api/851/75065)

	* 新增入参：TiProjectId

* [DescribeBillingResourceInstanceRunningJobs](https://cloud.tencent.com/document/api/851/101008)

	* 新增入参：TiProjectId

* [DescribeDataSource](https://cloud.tencent.com/document/api/851/129216)

	* 新增入参：TiProjectId

* [DescribeModelServiceGroups](https://cloud.tencent.com/document/api/851/82284)

	* 新增入参：TiProjectId

	* 新增出参：GlobalTotalCount

* [DescribeNotebook](https://cloud.tencent.com/document/api/851/95662)

	* 新增入参：TiProjectId

* [DescribeNotebooks](https://cloud.tencent.com/document/api/851/95653)

	* 新增入参：TiProjectId

* [DescribeTrainingTask](https://cloud.tencent.com/document/api/851/75089)

	* 新增入参：TiProjectId

* [DescribeTrainingTaskPods](https://cloud.tencent.com/document/api/851/75088)

	* 新增入参：TiProjectId

* [DescribeTrainingTasks](https://cloud.tencent.com/document/api/851/75087)

	* 新增入参：TiProjectId


修改数据结构：

* [NotebookDetail](https://cloud.tencent.com/document/api/851/75051#NotebookDetail)

	* 新增成员：LatestOperatorInfo

* [NotebookSetItem](https://cloud.tencent.com/document/api/851/75051#NotebookSetItem)

	* 新增成员：LatestOperatorInfo




## TI-ONE 训练平台(tione) 版本：2019-10-22



## 云点播(vod) 版本：2024-07-18



## 云点播(vod) 版本：2018-07-17

### 第 250 次发布

发布时间：2026-04-28 02:49:07

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeTaskDetail](https://cloud.tencent.com/document/api/266/33431)

	* 新增出参：AigcVideoRedrawTask


新增数据结构：

* [AigcVideoRedrawOutputFileInfo](https://cloud.tencent.com/document/api/266/31773#AigcVideoRedrawOutputFileInfo)
* [AigcVideoRedrawTask](https://cloud.tencent.com/document/api/266/31773#AigcVideoRedrawTask)
* [AigcVideoRedrawTaskInput](https://cloud.tencent.com/document/api/266/31773#AigcVideoRedrawTaskInput)
* [AigcVideoRedrawTaskOutput](https://cloud.tencent.com/document/api/266/31773#AigcVideoRedrawTaskOutput)



## Web 应用防火墙(waf) 版本：2018-01-25

### 第 148 次发布

发布时间：2026-04-28 02:58:31

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [GenerateLLMSecAnswer](https://cloud.tencent.com/document/api/627/131098)

修改接口：

* [DescribeLLMContentSecCheck](https://cloud.tencent.com/document/api/627/129968)

	* 新增入参：ImageEncode

	* <font color="#dd0000">**修改入参**：</font>Content


新增数据结构：

* [ImageResult](https://cloud.tencent.com/document/api/627/53609#ImageResult)
* [SSEClientMessage](https://cloud.tencent.com/document/api/627/53609#SSEClientMessage)

修改数据结构：

* [LLMDetectResult](https://cloud.tencent.com/document/api/627/53609#LLMDetectResult)

	* 新增成员：ImageResult




## 数据开发治理平台 WeData(wedata) 版本：2025-08-06



## 数据开发治理平台 WeData(wedata) 版本：2021-08-20

### 第 189 次发布

发布时间：2026-04-28 03:01:52

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [DatabaseRealViewVOPageVO](https://cloud.tencent.com/document/api/1267/76336#DatabaseRealViewVOPageVO)

	* 新增成员：SnapshotId




