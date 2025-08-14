# Release v1.1.11

## 负载均衡(clb) 版本：2018-03-17

### 第 134 次发布

发布时间：2025-08-15 01:24:35

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateTargetGroup](https://cloud.tencent.com/document/api/214/40559)

	* 新增入参：HealthCheck, ScheduleAlgorithm

* [ModifyTargetGroupAttribute](https://cloud.tencent.com/document/api/214/40552)

	* 新增入参：ScheduleAlgorithm, HealthCheck


新增数据结构：

* [TargetGroupHealthCheck](https://cloud.tencent.com/document/api/214/30694#TargetGroupHealthCheck)

修改数据结构：

* [TargetGroupInfo](https://cloud.tencent.com/document/api/214/30694#TargetGroupInfo)

	* 新增成员：ScheduleAlgorithm, HealthCheck




## 实时互动-教育版(lcic) 版本：2022-08-17

### 第 75 次发布

发布时间：2025-08-15 02:01:54

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateRoom](https://cloud.tencent.com/document/api/1639/80942)

	* 新增入参：Guests

* [DescribeRoom](https://cloud.tencent.com/document/api/1639/80941)

	* 新增出参：Guests


修改数据结构：

* [RoomInfo](https://cloud.tencent.com/document/api/1639/81423#RoomInfo)

	* 新增成员：Guests




## 媒体处理(mps) 版本：2019-06-12

### 第 138 次发布

发布时间：2025-08-15 02:10:40

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [HighlightSegmentItem](https://cloud.tencent.com/document/api/862/37615#HighlightSegmentItem)

	* 新增成员：Title, Summary

* [MediaAiAnalysisHighlightItem](https://cloud.tencent.com/document/api/862/37615#MediaAiAnalysisHighlightItem)

	* 新增成员：HighlightUrl, CovImgUrl




## 文字识别(ocr) 版本：2018-11-19

### 第 209 次发布

发布时间：2025-08-14 15:33:55

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [GeneralAccurateOCR](https://cloud.tencent.com/document/api/866/34937)

	* 新增入参：ConfigID




## 云数据库 PostgreSQL(postgres) 版本：2017-03-12

### 第 59 次发布

发布时间：2025-08-14 21:32:45

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除接口**：</font>

* CreateServerlessDBInstance
* DescribeDBSlowlogs
* DescribeServerlessDBInstances

<font color="#dd0000">**删除数据结构**：</font>

* NormalQueryItem
* ServerlessDBAccount
* ServerlessDBInstance
* ServerlessDBInstanceNetInfo
* SlowlogDetail



## 云托管 CloudBase Run(tcbr) 版本：2022-02-17

### 第 21 次发布

发布时间：2025-08-15 02:27:09

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [VolumeConf](https://cloud.tencent.com/document/api/1243/75713#VolumeConf)

修改数据结构：

* [DiffConfigItem](https://cloud.tencent.com/document/api/1243/75713#DiffConfigItem)

	* 新增成员：VolumesConf

* [ServerBaseConfig](https://cloud.tencent.com/document/api/1243/75713#ServerBaseConfig)

	* 新增成员：VolumesConf




## 容器安全服务(tcss) 版本：2020-11-01

### 第 79 次发布

发布时间：2025-08-15 02:28:36

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeAssetHostDetail](https://cloud.tencent.com/document/api/1285/65508)

	* 新增出参：AssetSyncTime

* [DescribeClusterDetail](https://cloud.tencent.com/document/api/1285/65448)

	* 新增出参：CheckFailReason


修改数据结构：

* [ClusterInfoItem](https://cloud.tencent.com/document/api/1285/65614#ClusterInfoItem)

	* 新增成员：AccessedErrorReason




## 容器服务(tke) 版本：2022-05-01



## 容器服务(tke) 版本：2018-05-25

### 第 203 次发布

发布时间：2025-08-15 02:38:07

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [ClusterAdvancedSettings](https://cloud.tencent.com/document/api/457/31866#ClusterAdvancedSettings)

	* 新增成员：DataPlaneV2




## 消息队列 RabbitMQ Serverless 版(trabbit) 版本：2023-04-18

### 第 4 次发布

发布时间：2025-08-15 02:40:37

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateRabbitMQServerlessExchange](https://cloud.tencent.com/document/api/1495/116149)

	* 新增入参：DelayedExchangeType


修改数据结构：

* [RabbitMQClusterInfo](https://cloud.tencent.com/document/api/1495/116155#RabbitMQClusterInfo)

	* 新增成员：SendReceiveRatio, TraceTime

* [RabbitMQServerlessAccessInfo](https://cloud.tencent.com/document/api/1495/116155#RabbitMQServerlessAccessInfo)

	* 新增成员：PublicClbId




## 消息队列 RocketMQ 版(trocket) 版本：2023-03-08

### 第 45 次发布

发布时间：2025-08-15 02:40:59

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeProducerList](https://cloud.tencent.com/document/api/1493/122548)

新增数据结构：

* [ProducerInfo](https://cloud.tencent.com/document/api/1493/96031#ProducerInfo)



## TSF-应用管理&Consul(tsf) 版本：2018-03-26

### 第 130 次发布

发布时间：2025-08-15 02:44:27

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateNamespace](https://cloud.tencent.com/document/api/649/36098)

	* 新增入参：CreateK8sNamespaceFlag


新增数据结构：

* [VolumeClaimTemplatesOption](https://cloud.tencent.com/document/api/649/36099#VolumeClaimTemplatesOption)

修改数据结构：

* [BusinessLogConfig](https://cloud.tencent.com/document/api/649/36099#BusinessLogConfig)

	* 新增成员：FilebeatIgnoreOlder, FilebeatHarvesterLimit, FilebeatCloseInactive, FilebeatCleanInactive

* [VolumeInfo](https://cloud.tencent.com/document/api/649/36099#VolumeInfo)

	* 新增成员：VolumeClaimTemplateOption




## 云点播(vod) 版本：2024-07-18



## 云点播(vod) 版本：2018-07-17

### 第 207 次发布

发布时间：2025-08-15 02:48:05

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ProcessMediaByMPS](https://cloud.tencent.com/document/api/266/121993)

	* 新增入参：ExtInfo


修改数据结构：

* [EventContent](https://cloud.tencent.com/document/api/266/31773#EventContent)

	* 新增成员：ProcessMediaByMPSCompleteEvent

* [MPSOutputFile](https://cloud.tencent.com/document/api/266/31773#MPSOutputFile)

	* 新增成员：Definition




## Web 应用防火墙(waf) 版本：2018-01-25

### 第 124 次发布

发布时间：2025-08-14 21:41:01

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DeleteOwaspRuleStatus](https://cloud.tencent.com/document/api/627/122547)
* [DescribeOwaspRuleTypes](https://cloud.tencent.com/document/api/627/122546)
* [DescribeOwaspRules](https://cloud.tencent.com/document/api/627/122545)
* [ModifyOwaspRuleStatus](https://cloud.tencent.com/document/api/627/122544)
* [ModifyOwaspRuleTypeAction](https://cloud.tencent.com/document/api/627/122543)
* [ModifyOwaspRuleTypeLevel](https://cloud.tencent.com/document/api/627/122542)
* [ModifyOwaspRuleTypeStatus](https://cloud.tencent.com/document/api/627/122541)

新增数据结构：

* [OwaspRule](https://cloud.tencent.com/document/api/627/53609#OwaspRule)
* [OwaspRuleType](https://cloud.tencent.com/document/api/627/53609#OwaspRuleType)



## 数据开发治理平台 WeData(wedata) 版本：2021-08-20

### 第 168 次发布

发布时间：2025-08-15 02:56:00

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [CompareResult](https://cloud.tencent.com/document/api/1267/76336#CompareResult)

	* 新增成员：ComputeExpression

* [RuleExecResult](https://cloud.tencent.com/document/api/1267/76336#RuleExecResult)

	* 新增成员：RuleGroupName, DatasourceId, DatasourceName, DatabaseName, SchemaName, TableName, RuleGroupExist, DatasourceType, RuleGroupTableId, MonitorType, FinishTime




