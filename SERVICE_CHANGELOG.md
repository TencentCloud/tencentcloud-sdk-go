# Release v1.3.40

## 云硬盘(cbs) 版本：2017-03-12

### 第 73 次发布

发布时间：2026-01-28 01:17:02

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CopyAutoSnapshotPolicyCrossAccount](https://cloud.tencent.com/document/api/362/127847)



## TDSQL-C MySQL 版(cynosdb) 版本：2019-01-07

### 第 153 次发布

发布时间：2026-01-28 01:36:20

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DeleteClusterSaveBackup](https://cloud.tencent.com/document/api/1003/127851)
* [DescribeRedoLogs](https://cloud.tencent.com/document/api/1003/127850)
* [DescribeSaveBackupClusters](https://cloud.tencent.com/document/api/1003/127849)
* [ModifySnapBackupCrossRegionConfig](https://cloud.tencent.com/document/api/1003/127848)

修改接口：

* [DescribeBackupConfig](https://cloud.tencent.com/document/api/1003/48094)

	* 新增出参：SnapshotSecondaryBackupConfig

* [IsolateCluster](https://cloud.tencent.com/document/api/1003/48082)

	* 新增入参：SaveBackup

* [IsolateInstance](https://cloud.tencent.com/document/api/1003/48081)

	* 新增入参：SaveBackup

* [ModifyBackupConfig](https://cloud.tencent.com/document/api/1003/48090)

	* 新增入参：SnapshotSecondaryBackupConfig

* [RollbackToNewCluster](https://cloud.tencent.com/document/api/1003/104727)

	* 新增入参：FromSaveBackup


新增数据结构：

* [BackupConfigInfo](https://cloud.tencent.com/document/api/1003/48097#BackupConfigInfo)
* [BackupRegionAndIds](https://cloud.tencent.com/document/api/1003/48097#BackupRegionAndIds)
* [QuerySimpleFilter](https://cloud.tencent.com/document/api/1003/48097#QuerySimpleFilter)
* [RedoLogItem](https://cloud.tencent.com/document/api/1003/48097#RedoLogItem)
* [SaveBackupClusterInfo](https://cloud.tencent.com/document/api/1003/48097#SaveBackupClusterInfo)
* [SnapshotBackupConfig](https://cloud.tencent.com/document/api/1003/48097#SnapshotBackupConfig)

修改数据结构：

* [CynosdbClusterDetail](https://cloud.tencent.com/document/api/1003/48097#CynosdbClusterDetail)

	* 新增成员：IsOpenTDE




## Elasticsearch Service(es) 版本：2025-01-01



## Elasticsearch Service(es) 版本：2018-04-16

### 第 95 次发布

发布时间：2026-01-28 01:48:30

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateCollector](https://cloud.tencent.com/document/api/845/127853)

新增数据结构：

* [CollectorConfigInfo](https://cloud.tencent.com/document/api/845/30634#CollectorConfigInfo)
* [CollectorOutputInstance](https://cloud.tencent.com/document/api/845/30634#CollectorOutputInstance)
* [CollectorTarget](https://cloud.tencent.com/document/api/845/30634#CollectorTarget)
* [Namespaces](https://cloud.tencent.com/document/api/845/30634#Namespaces)
* [PodLabel](https://cloud.tencent.com/document/api/845/30634#PodLabel)



## 云游戏(gs) 版本：2019-11-18

### 第 64 次发布

发布时间：2026-01-28 01:54:10

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [InstallAndroidInstancesAppWithURL](https://cloud.tencent.com/document/api/1162/118979)

	* 新增入参：AndroidAppMD5




## 云数据库 MongoDB(mongodb) 版本：2019-07-25

### 第 63 次发布

发布时间：2026-01-28 02:13:30

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeSRVConnectionDomain](https://cloud.tencent.com/document/api/240/127857)
* [DisableSRVConnectionUrl](https://cloud.tencent.com/document/api/240/127856)
* [EnableSRVConnectionUrl](https://cloud.tencent.com/document/api/240/127855)
* [ModifySRVConnectionUrl](https://cloud.tencent.com/document/api/240/127854)



## 云数据库 MongoDB(mongodb) 版本：2018-04-08



## NLP 技术(nlp) 版本：2019-04-08

### 第 26 次发布

发布时间：2026-01-28 02:23:53

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除接口**：</font>

* ClassifyContent
* ComposeCouplet
* EvaluateSentenceSimilarity

<font color="#dd0000">**删除数据结构**：</font>

* Category
* SentencePair



## 容器镜像服务(tcr) 版本：2019-09-24

### 第 74 次发布

发布时间：2026-01-28 02:52:46

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateTagRetentionRule](https://cloud.tencent.com/document/api/1141/54803)

	* 新增入参：AdvancedRuleItems

	* <font color="#dd0000">**修改入参**：</font>RetentionRule

* [DeleteRepository](https://cloud.tencent.com/document/api/1141/42724)

	* 新增入参：ForceDelete

* [ModifyTagRetentionRule](https://cloud.tencent.com/document/api/1141/54798)

	* 新增入参：AdvancedRuleItems

	* <font color="#dd0000">**修改入参**：</font>RetentionRule


新增数据结构：

* [FilterSelector](https://cloud.tencent.com/document/api/1141/41603#FilterSelector)
* [RetentionRuleItem](https://cloud.tencent.com/document/api/1141/41603#RetentionRuleItem)

修改数据结构：

* [RetentionPolicy](https://cloud.tencent.com/document/api/1141/41603#RetentionPolicy)

	* 新增成员：AdvancedRuleItems




## 容器安全服务(tcss) 版本：2020-11-01

### 第 86 次发布

发布时间：2026-01-28 02:54:42

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**预下线接口**：</font>

* AddAndPublishNetworkFirewallPolicyDetail
* AddAndPublishNetworkFirewallPolicyYamlDetail
* AddNetworkFirewallPolicyDetail
* AddNetworkFirewallPolicyYamlDetail
* CheckNetworkFirewallPolicyYaml
* ConfirmNetworkFirewallPolicy
* CreateNetworkFirewallClusterRefresh
* CreateNetworkFirewallPolicyDiscover
* CreateNetworkFirewallPublish
* CreateNetworkFirewallUndoPublish
* DeleteNetworkFirewallPolicyDetail
* DescribeNetworkFirewallAuditRecord
* DescribeNetworkFirewallClusterList
* DescribeNetworkFirewallClusterRefreshStatus
* DescribeNetworkFirewallNamespaceLabelList
* DescribeNetworkFirewallNamespaceList
* DescribeNetworkFirewallPodLabelsList
* DescribeNetworkFirewallPolicyDetail
* DescribeNetworkFirewallPolicyDiscover
* DescribeNetworkFirewallPolicyList
* DescribeNetworkFirewallPolicyStatus
* DescribeNetworkFirewallPolicyYamlDetail
* UpdateAndPublishNetworkFirewallPolicyDetail
* UpdateAndPublishNetworkFirewallPolicyYamlDetail
* UpdateNetworkFirewallPolicyDetail
* UpdateNetworkFirewallPolicyYamlDetail



## 云点播(vod) 版本：2024-07-18



## 云点播(vod) 版本：2018-07-17

### 第 226 次发布

发布时间：2026-01-28 03:33:19

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateProcessImageAsyncTemplate](https://cloud.tencent.com/document/api/266/127862)
* [DeleteProcessImageAsyncTemplate](https://cloud.tencent.com/document/api/266/127861)
* [DescribeProcessImageAsyncTemplates](https://cloud.tencent.com/document/api/266/127860)
* [ModifyProcessImageAsyncTemplate](https://cloud.tencent.com/document/api/266/127859)
* [ProcessImageAsync](https://cloud.tencent.com/document/api/266/127858)

修改接口：

* [DescribeTaskDetail](https://cloud.tencent.com/document/api/266/33431)

	* 新增出参：ProcessImageAsyncTask


新增数据结构：

* [AdvancedSuperResolutionConfig](https://cloud.tencent.com/document/api/266/31773#AdvancedSuperResolutionConfig)
* [ImageDenoiseConfig](https://cloud.tencent.com/document/api/266/31773#ImageDenoiseConfig)
* [ImageEncodeConfig](https://cloud.tencent.com/document/api/266/31773#ImageEncodeConfig)
* [ImageEnhanceConfig](https://cloud.tencent.com/document/api/266/31773#ImageEnhanceConfig)
* [ImageSceneAigcEncodeConfig](https://cloud.tencent.com/document/api/266/31773#ImageSceneAigcEncodeConfig)
* [ProcessImageAsync](https://cloud.tencent.com/document/api/266/31773#ProcessImageAsync)
* [ProcessImageAsyncInput](https://cloud.tencent.com/document/api/266/31773#ProcessImageAsyncInput)
* [ProcessImageAsyncOutput](https://cloud.tencent.com/document/api/266/31773#ProcessImageAsyncOutput)
* [ProcessImageAsyncOutputConfig](https://cloud.tencent.com/document/api/266/31773#ProcessImageAsyncOutputConfig)
* [ProcessImageAsyncOutputFileInfo](https://cloud.tencent.com/document/api/266/31773#ProcessImageAsyncOutputFileInfo)
* [ProcessImageAsyncTask](https://cloud.tencent.com/document/api/266/31773#ProcessImageAsyncTask)
* [ProcessImageAsyncTaskInput](https://cloud.tencent.com/document/api/266/31773#ProcessImageAsyncTaskInput)
* [ProcessImageAsyncTemplateItem](https://cloud.tencent.com/document/api/266/31773#ProcessImageAsyncTemplateItem)

修改数据结构：

* [ChangeClothesConfig](https://cloud.tencent.com/document/api/266/31773#ChangeClothesConfig)

	* 新增成员：Prompt

* [SceneAigcImageOutputConfig](https://cloud.tencent.com/document/api/266/31773#SceneAigcImageOutputConfig)

	* 新增成员：EncodeConfig




## 私有网络(vpc) 版本：2017-03-12

### 第 293 次发布

发布时间：2026-01-28 03:38:00

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [ConnectionStateTimeouts](https://cloud.tencent.com/document/api/215/15824#ConnectionStateTimeouts)

	* 新增成员：TCPTimeWaitTimeout

	* <font color="#dd0000">**删除成员**：</font>TcpTimeWaitTimeout




