# Release v1.3.95

## 云防火墙(cfw) 版本：2019-09-04

### 第 99 次发布

发布时间：2026-05-12 01:15:30

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeSerialRegion](https://cloud.tencent.com/document/api/1132/131411)

新增数据结构：

* [SerialRegionInfo](https://cloud.tencent.com/document/api/1132/49071#SerialRegionInfo)



## TDSQL-C MySQL 版(cynosdb) 版本：2019-01-07

### 第 165 次发布

发布时间：2026-05-12 01:22:46

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [ProxyConfig](https://cloud.tencent.com/document/api/1003/48097#ProxyConfig)




## Elasticsearch Service(es) 版本：2025-01-01



## Elasticsearch Service(es) 版本：2018-04-16

### 第 100 次发布

发布时间：2026-05-12 01:28:53

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateClusterSnapshot](https://cloud.tencent.com/document/api/845/112044)

	* 新增入参：MaxSnapshotPerSec

* [DeleteInstance](https://cloud.tencent.com/document/api/845/30632)

	* 新增入参：LockEnabled, LockDuration

* [RestoreClusterSnapshot](https://cloud.tencent.com/document/api/845/112040)

	* 新增入参：MaxRestorePerSec


修改数据结构：

* [CosBackup](https://cloud.tencent.com/document/api/845/30634#CosBackup)

	* 新增成员：MaxSnapshotPerSec, MaxRestorePerSec, InstanceId

* [InstanceInfo](https://cloud.tencent.com/document/api/845/30634#InstanceInfo)

	* 新增成员：IsInRecycleBin, RecycleLockEnabled, MayDestroyPoint, DelayDestroyInterval

* [Operation](https://cloud.tencent.com/document/api/845/30634#Operation)

	* 新增成员：SuspendedReason

* [ProcessDetail](https://cloud.tencent.com/document/api/845/30634#ProcessDetail)

	* 新增成员：EstimatedTimeRemaining

* [Snapshots](https://cloud.tencent.com/document/api/845/30634#Snapshots)

	* 新增成员：MaxSnapshotPerSec, InstanceId




## 云直播CSS(live) 版本：2018-08-01

### 第 176 次发布

发布时间：2026-05-12 01:37:55

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateSceneVideoTask](https://cloud.tencent.com/document/api/267/131414)
* [DescribeSceneVideoTask](https://cloud.tencent.com/document/api/267/131413)

新增数据结构：

* [SceneStoreCosParam](https://cloud.tencent.com/document/api/267/20474#SceneStoreCosParam)
* [SceneVideoExtraParam](https://cloud.tencent.com/document/api/267/20474#SceneVideoExtraParam)
* [SceneVideoOutputInfo](https://cloud.tencent.com/document/api/267/20474#SceneVideoOutputInfo)
* [SceneVideoReferenceImageInfo](https://cloud.tencent.com/document/api/267/20474#SceneVideoReferenceImageInfo)
* [SceneVideoReferenceVideoInfo](https://cloud.tencent.com/document/api/267/20474#SceneVideoReferenceVideoInfo)



## 媒体处理(mps) 版本：2019-06-12

### 第 202 次发布

发布时间：2026-05-12 01:47:24

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DetectVideoSubtitleArea](https://cloud.tencent.com/document/api/862/131415)

新增数据结构：

* [SubtitleArea](https://cloud.tencent.com/document/api/862/37615#SubtitleArea)



## 实时音视频(trtc) 版本：2019-07-22

### 第 143 次发布

发布时间：2026-05-12 02:37:21

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeTRTCAIRecognitionUsage](https://cloud.tencent.com/document/api/647/131418)
* [DescribeTRTCDedicatedCloudAccUsage](https://cloud.tencent.com/document/api/647/131417)
* [DescribeTRTCSegmentModerationUsage](https://cloud.tencent.com/document/api/647/131416)

新增数据结构：

* [UsageList](https://cloud.tencent.com/document/api/647/44055#UsageList)



## TSF-Polaris&ZK&网关(tse) 版本：2020-12-07

### 第 107 次发布

发布时间：2026-05-12 02:38:32

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeGovernanceServices](https://cloud.tencent.com/document/api/1364/104380)

	* 新增入参：Type


修改数据结构：

* [GovernanceService](https://cloud.tencent.com/document/api/1364/54942#GovernanceService)

	* 新增成员：Type

* [GovernanceServiceInput](https://cloud.tencent.com/document/api/1364/54942#GovernanceServiceInput)

	* 新增成员：Type




## TSF-应用管理&Consul(tsf) 版本：2018-03-26

### 第 140 次发布

发布时间：2026-05-12 02:40:17

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribePrograms](https://cloud.tencent.com/document/api/649/73477)

	* 新增入参：SearchFilters


新增数据结构：

* [SearchFiltersProgram](https://cloud.tencent.com/document/api/649/36099#SearchFiltersProgram)



## Web 应用防火墙(waf) 版本：2018-01-25

### 第 151 次发布

发布时间：2026-05-12 02:54:35

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeQClawContentSecCheck](https://cloud.tencent.com/document/api/627/129967)

新增数据结构：

* [ApiGuardContent](https://cloud.tencent.com/document/api/627/53609#ApiGuardContent)
* [ClawRiskItem](https://cloud.tencent.com/document/api/627/53609#ClawRiskItem)
* [LLMRisks](https://cloud.tencent.com/document/api/627/53609#LLMRisks)



## 数据开发治理平台 WeData(wedata) 版本：2025-08-06



## 数据开发治理平台 WeData(wedata) 版本：2021-08-20

### 第 191 次发布

发布时间：2026-05-12 02:57:54

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CommitIntegrationTask](https://cloud.tencent.com/document/api/1267/82526)

	* 新增入参：CurrentStatus




