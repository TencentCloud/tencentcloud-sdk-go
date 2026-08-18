# Release v1.3.164

## 腾讯云智能体开发平台(adp) 版本：2026-05-20

### 第 16 次发布

发布时间：2026-08-19 01:08:08

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [Identity](https://cloud.tencent.com/document/api/1759/132545#Identity)

修改数据结构：

* [CorpShareConfig](https://cloud.tencent.com/document/api/1759/132545#CorpShareConfig)

	* 新增成员：ShareScopeList


### 第 15 次发布

发布时间：2026-08-18 17:29:42

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeConcurrencyLimitDetailList](https://cloud.tencent.com/document/api/1759/136064)
* [DescribeConsumptionDetailList](https://cloud.tencent.com/document/api/1759/136063)
* [DescribeMetricOverviewList](https://cloud.tencent.com/document/api/1759/136062)
* [DescribeUsageDetailList](https://cloud.tencent.com/document/api/1759/136061)
* [DescribeUsageSummaryList](https://cloud.tencent.com/document/api/1759/136060)

新增数据结构：

* [CallSource](https://cloud.tencent.com/document/api/1759/132545#CallSource)
* [ConcurrencyLimitDetail](https://cloud.tencent.com/document/api/1759/132545#ConcurrencyLimitDetail)
* [ConsumptionClassification](https://cloud.tencent.com/document/api/1759/132545#ConsumptionClassification)
* [ConsumptionDetail](https://cloud.tencent.com/document/api/1759/132545#ConsumptionDetail)
* [ConsumptionUsage](https://cloud.tencent.com/document/api/1759/132545#ConsumptionUsage)
* [MetricOverview](https://cloud.tencent.com/document/api/1759/132545#MetricOverview)
* [ModelUsageDetail](https://cloud.tencent.com/document/api/1759/132545#ModelUsageDetail)
* [ModelUsageSummary](https://cloud.tencent.com/document/api/1759/132545#ModelUsageSummary)
* [PlatformUsageSummary](https://cloud.tencent.com/document/api/1759/132545#PlatformUsageSummary)
* [PluginUsageDetail](https://cloud.tencent.com/document/api/1759/132545#PluginUsageDetail)
* [PluginUsageSummary](https://cloud.tencent.com/document/api/1759/132545#PluginUsageSummary)
* [ResourceConsumption](https://cloud.tencent.com/document/api/1759/132545#ResourceConsumption)
* [TimeRange](https://cloud.tencent.com/document/api/1759/132545#TimeRange)
* [UsageDetail](https://cloud.tencent.com/document/api/1759/132545#UsageDetail)
* [UsageSummary](https://cloud.tencent.com/document/api/1759/132545#UsageSummary)
* [ViewScope](https://cloud.tencent.com/document/api/1759/132545#ViewScope)

<font color="#dd0000">**删除数据结构**：</font>

* ClawAgentCustomConfig

修改数据结构：

* [ClawAgentConfig](https://cloud.tencent.com/document/api/1759/132545#ClawAgentConfig)

	* <font color="#dd0000">**删除成员**：</font>CustomConfig




## 云硬盘(cbs) 版本：2017-03-12

### 第 79 次发布

发布时间：2026-08-19 01:19:17

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeDedicatedClusterDiskStatistics](https://cloud.tencent.com/document/api/362/136032)

	* 新增入参：DedicatedClusterId

	* 新增出参：DedicatedClusterDiskStatisticSet

* [DescribeRemoteDisks](https://cloud.tencent.com/document/api/362/132685)

	* 新增出参：RemoteDiskSet, TotalCount


新增数据结构：

* [DedicatedClusterDiskStatistic](https://cloud.tencent.com/document/api/362/15669#DedicatedClusterDiskStatistic)
* [RemoteDiskDetail](https://cloud.tencent.com/document/api/362/15669#RemoteDiskDetail)



## 云数据库 MySQL(cdb) 版本：2017-03-20

### 第 228 次发布

发布时间：2026-08-19 01:20:47

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [RoGroup](https://cloud.tencent.com/document/api/236/15878#RoGroup)

	* 新增成员：RoGroupType




## 云托付物理服务器(chc) 版本：2023-04-18

### 第 12 次发布

发布时间：2026-08-19 01:27:28

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ExportCustomerWorkOrderDetail](https://cloud.tencent.com/document/api/1448/118996)

	* <font color="#dd0000">**修改入参**：</font>WorkOrderType




## 消息队列 CKafka 版(ckafka) 版本：2019-08-19

### 第 151 次发布

发布时间：2026-08-19 01:28:28

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeTopicDetail](https://cloud.tencent.com/document/api/597/40845)

	* 新增入参：SearchWordIgnoreCaseFlag


修改数据结构：

* [DatahubTaskInfo](https://cloud.tencent.com/document/api/597/40861#DatahubTaskInfo)

	* 新增成员：WarnMessage

* [DescribeDatahubTaskRes](https://cloud.tencent.com/document/api/597/40861#DescribeDatahubTaskRes)

	* 新增成员：WarnMessage




## 日志服务(cls) 版本：2020-10-16

### 第 174 次发布

发布时间：2026-08-19 01:31:50

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [ExtractRuleInfo](https://cloud.tencent.com/document/api/614/56471#ExtractRuleInfo)

	* 新增成员：Units, IncludeKernel, UseJournalTime, KeysDelimiter, KeysFlag




## TDSQL-C MySQL 版(cynosdb) 版本：2019-01-07

### 第 187 次发布

发布时间：2026-08-19 01:45:21

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [RollBackCluster](https://cloud.tencent.com/document/api/1003/70115)

	* <font color="#dd0000">**修改入参**：</font>RollbackId




## 腾讯云数据分析智能体(dataagent) 版本：2025-05-13

### 第 23 次发布

发布时间：2026-08-19 01:47:41

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除接口**：</font>

* AddScene
* DeleteScene
* QuerySceneList
* UpdateScene

<font color="#dd0000">**删除数据结构**：</font>

* ExampleQA
* Scene
* SearchConfig



## 数据湖计算 DLC(dlc) 版本：2021-01-25

### 第 177 次发布

发布时间：2026-08-19 01:50:47

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateLab](https://cloud.tencent.com/document/api/1342/135449)

	* <font color="#dd0000">**修改入参**：</font>Image

* [UpdateLab](https://cloud.tencent.com/document/api/1342/135436)

	* <font color="#dd0000">**修改入参**：</font>Image




## 数据加速器 GooseFS(goosefs) 版本：2022-05-19

### 第 21 次发布

发布时间：2026-08-19 02:09:17

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [MountPointEntry](https://cloud.tencent.com/document/api/1424/95076#MountPointEntry)

修改数据结构：

* [ClientNodeAttribute](https://cloud.tencent.com/document/api/1424/95076#ClientNodeAttribute)

	* 新增成员：MountPoints

* [CustomerClusterAttr](https://cloud.tencent.com/document/api/1424/95076#CustomerClusterAttr)

	* 新增成员：Zone, MountStorageNum, StorageFileSystemId




## 云数据库 MongoDB(mongodb) 版本：2019-07-25

### 第 75 次发布

发布时间：2026-08-19 02:26:57

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyDBInstanceSpec](https://cloud.tencent.com/document/api/240/38565)

	* 新增入参：ModifyShardList


新增数据结构：

* [ModifyShardSpecInfo](https://cloud.tencent.com/document/api/240/38576#ModifyShardSpecInfo)



## 云数据库 MongoDB(mongodb) 版本：2018-04-08



## 腾讯云可观测平台(monitor) 版本：2023-06-16



## 腾讯云可观测平台(monitor) 版本：2018-07-24

### 第 165 次发布

发布时间：2026-08-19 02:27:49

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [ModifyPrometheusInstanceAccessPoints](https://cloud.tencent.com/document/api/248/136108)



## 流计算 Oceanus(oceanus) 版本：2019-04-22

### 第 92 次发布

发布时间：2026-08-19 02:33:36

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [JobConfig](https://cloud.tencent.com/document/api/849/52010#JobConfig)

	* 新增成员：IsLocked




## 云数据库 PostgreSQL(postgres) 版本：2017-03-12

### 第 73 次发布

发布时间：2026-08-19 02:37:21

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeDBInstanceSSLConfig](https://cloud.tencent.com/document/api/409/116172)

	* 新增出参：CACert, CAJKS, CAP7B




## 凭据管理系统(ssm) 版本：2019-09-23

### 第 19 次发布

发布时间：2026-08-19 02:45:55

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeSecret](https://cloud.tencent.com/document/api/1140/40526)

	* 新增出参：NextRotationTime

* [GetSecretValue](https://cloud.tencent.com/document/api/1140/40522)

	* 新增入参：EncryptionPublicKey, EncryptionAlgorithm

* [ListSecrets](https://cloud.tencent.com/document/api/1140/40519)

	* 新增入参：InstanceID




## TokenHub(tokenhub) 版本：2026-03-22

### 第 19 次发布

发布时间：2026-08-19 03:04:39

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeModelQuota](https://cloud.tencent.com/document/api/1823/136110)



## 数据开发治理平台 WeData(wedata) 版本：2025-08-06

### 第 22 次发布

发布时间：2026-08-19 03:24:47

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [TriggerTaskRunBrief](https://cloud.tencent.com/document/api/1267/123643#TriggerTaskRunBrief)

	* 新增成员：AssociatedEntityExist, ScheduleRunType




## 数据开发治理平台 WeData(wedata) 版本：2021-08-20

### 第 201 次发布

发布时间：2026-08-19 03:21:16

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeIntegrationTask](https://cloud.tencent.com/document/api/1267/82495)

	* 新增出参：TaskVersionList

* [DescribeStreamTaskLogList](https://cloud.tencent.com/document/api/1267/82487)

	* 新增入参：Context

	* 新增出参：Context

* [DescribeTableContentPreview](https://cloud.tencent.com/document/api/1267/129275)

	* 新增入参：ResourceGroupId, Sql, EngineId


新增数据结构：

* [RealtimeTaskInstanceVO](https://cloud.tencent.com/document/api/1267/76336#RealtimeTaskInstanceVO)



## 联网搜索API(wsa) 版本：2025-05-08

### 第 5 次发布

发布时间：2026-08-19 03:27:25

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [SearchPro](https://cloud.tencent.com/document/api/1806/121811)

	* <font color="#dd0000">**删除入参**：</font>FromTime, ToTime




