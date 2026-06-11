# Release v1.3.115

## 应用性能监控(apm) 版本：2021-06-22

### 第 62 次发布

发布时间：2026-06-12 01:09:24

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeApmInstances](https://cloud.tencent.com/document/api/1463/65103)

	* 新增入参：PageIndex, PageSize, Keyword, OrderDirection, OrderBy

	* 新增出参：TotalCount, PageIndex, PageSize

* [DescribeMetricRecords](https://cloud.tencent.com/document/api/1463/68254)

	* 新增入参：ServiceID

* [DescribeTopologyNew](https://cloud.tencent.com/document/api/1463/126866)

	* 新增入参：EnableResourceLink

	* 新增出参：OverviewStats


新增数据结构：

* [OverviewStats](https://cloud.tencent.com/document/api/1463/64927#OverviewStats)
* [TopologyNodeStats](https://cloud.tencent.com/document/api/1463/64927#TopologyNodeStats)

修改数据结构：

* [TopologyEdgeNew](https://cloud.tencent.com/document/api/1463/64927#TopologyEdgeNew)

	* 新增成员：ReqCnt

* [TopologyNode](https://cloud.tencent.com/document/api/1463/64927#TopologyNode)

	* 新增成员：ReqCnt, ConsumerReqCnt




## 费用中心(billing) 版本：2018-07-09

### 第 90 次发布

发布时间：2026-06-12 01:11:04

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [ConsumptionBusinessSummaryDataItem](https://cloud.tencent.com/document/api/555/19183#ConsumptionBusinessSummaryDataItem)

	* 新增成员：LeftRealTotalCost

* [ConsumptionProjectSummaryDataItem](https://cloud.tencent.com/document/api/555/19183#ConsumptionProjectSummaryDataItem)

	* 新增成员：LeftRealTotalCost

* [ConsumptionRegionSummaryDataItem](https://cloud.tencent.com/document/api/555/19183#ConsumptionRegionSummaryDataItem)

	* 新增成员：LeftRealTotalCost

* [ConsumptionResourceSummaryDataItem](https://cloud.tencent.com/document/api/555/19183#ConsumptionResourceSummaryDataItem)

	* 新增成员：LeftRealTotalCost




## 云硬盘(cbs) 版本：2017-03-12

### 第 76 次发布

发布时间：2026-06-12 01:12:28

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [AttachRemoteDisks](https://cloud.tencent.com/document/api/362/132688)

	* 新增入参：InstanceId, RemoteDiskIds

* [CreateRemoteDisks](https://cloud.tencent.com/document/api/362/132687)

	* 新增入参：DiskChargeType, DiskSize, InstanceId, Placement, DiskChargePrepaid, DiskCount, DiskName

* [DescribeRemoteDisksDeniedActions](https://cloud.tencent.com/document/api/362/132684)

	* 新增入参：RemoteDiskIds

* [DetachRemoteDisks](https://cloud.tencent.com/document/api/362/132683)

	* 新增入参：InstanceId, RemoteDiskIds, ForceDetach

* [InquirePriceCreateRemoteDisks](https://cloud.tencent.com/document/api/362/132682)

	* 新增入参：DiskChargeType, DiskSize, DiskChargePrepaid, DiskCount

* [InquirePriceRenewRemoteDisks](https://cloud.tencent.com/document/api/362/132681)

	* 新增入参：DiskChargePrepaidSet, RemoteDiskIds

* [ModifyRemoteDiskAttributes](https://cloud.tencent.com/document/api/362/132680)

	* 新增入参：RemoteDiskIds, DiskName, ProjectId

* [RenewRemoteDisk](https://cloud.tencent.com/document/api/362/132679)

	* 新增入参：DiskChargePrepaid, RemoteDiskId

* [TerminateRemoteDisks](https://cloud.tencent.com/document/api/362/132676)

	* 新增入参：RemoteDiskIds


新增数据结构：

* [RemoteDiskChargePrepaid](https://cloud.tencent.com/document/api/362/15669#RemoteDiskChargePrepaid)



## 文件存储(cfs) 版本：2019-07-19

### 第 51 次发布

发布时间：2026-06-12 01:15:15

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateLifecycleDataTask](https://cloud.tencent.com/document/api/582/122052)

	* 新增入参：ListPath

	* <font color="#dd0000">**修改入参**：</font>TaskPath


修改数据结构：

* [LifecycleDataTaskInfo](https://cloud.tencent.com/document/api/582/38175#LifecycleDataTaskInfo)

	* 新增成员：ListPath




## TDSQL-C MySQL 版(cynosdb) 版本：2019-01-07

### 第 175 次发布

发布时间：2026-06-12 01:23:31

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CancelClusterServerlessScalePlan](https://cloud.tencent.com/document/api/1003/132827)
* [CreateClusterPeriodScalePolicy](https://cloud.tencent.com/document/api/1003/132826)
* [DeleteClusterPeriodScalePolicy](https://cloud.tencent.com/document/api/1003/132825)
* [DescribeClusterPeriodScalePolicy](https://cloud.tencent.com/document/api/1003/132824)
* [DescribeClusterServerlessScalePlans](https://cloud.tencent.com/document/api/1003/132823)
* [ModifyClusterPeriodScalePolicy](https://cloud.tencent.com/document/api/1003/132822)
* [OpenAIOptimizer](https://cloud.tencent.com/document/api/1003/132828)

修改接口：

* [DescribeInstanceSpecs](https://cloud.tencent.com/document/api/1003/48084)

	* 新增入参：ClusterLevel


新增数据结构：

* [ClusterPeriodScalePolicy](https://cloud.tencent.com/document/api/1003/48097#ClusterPeriodScalePolicy)
* [ClusterServerlessScalePlan](https://cloud.tencent.com/document/api/1003/48097#ClusterServerlessScalePlan)

修改数据结构：

* [CynosdbClusterDetail](https://cloud.tencent.com/document/api/1003/48097#CynosdbClusterDetail)

	* 新增成员：ClusterLevel




## 数据传输服务(dts) 版本：2021-12-06

### 第 53 次发布

发布时间：2026-06-12 01:28:14

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [StepDetailInfo](https://cloud.tencent.com/document/api/571/82108#StepDetailInfo)

	* 新增成员：FinishTime

* [StepInfo](https://cloud.tencent.com/document/api/571/82108#StepInfo)

	* 新增成员：FinishTime




## 数据传输服务(dts) 版本：2018-03-30



## 高性能应用服务(hai) 版本：2023-08-12

### 第 25 次发布

发布时间：2026-06-12 01:33:35

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DeployInferService](https://cloud.tencent.com/document/api/1721/129387)

	* 新增入参：SecurityType




## 智能全局流量管理(igtm) 版本：2023-10-24

### 第 5 次发布

发布时间：2026-06-12 01:34:19

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyAddressPool](https://cloud.tencent.com/document/api/1551/120458)

	* 新增入参：KeepResource

* [ModifyStrategy](https://cloud.tencent.com/document/api/1551/120438)

	* 新增入参：KeepResource




## 媒体处理(mps) 版本：2019-06-12

### 第 208 次发布

发布时间：2026-06-12 01:43:00

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateAigcAudioTask](https://cloud.tencent.com/document/api/862/132830)
* [DescribeAigcAudioTask](https://cloud.tencent.com/document/api/862/132829)

新增数据结构：

* [AiTryOnConfig](https://cloud.tencent.com/document/api/862/37615#AiTryOnConfig)
* [AigcAudioExtraParam](https://cloud.tencent.com/document/api/862/37615#AigcAudioExtraParam)
* [AigcAudioOutputAudioInfo](https://cloud.tencent.com/document/api/862/37615#AigcAudioOutputAudioInfo)
* [AigcAudioOutputVideoInfo](https://cloud.tencent.com/document/api/862/37615#AigcAudioOutputVideoInfo)
* [AigcAudioReferenceAudioInfo](https://cloud.tencent.com/document/api/862/37615#AigcAudioReferenceAudioInfo)
* [AigcAudioReferenceVideoInfo](https://cloud.tencent.com/document/api/862/37615#AigcAudioReferenceVideoInfo)

修改数据结构：

* [ImageTaskInput](https://cloud.tencent.com/document/api/862/37615#ImageTaskInput)

	* 新增成员：AiTryOnConfig




## 云数据库 PostgreSQL(postgres) 版本：2017-03-12

### 第 67 次发布

发布时间：2026-06-12 01:47:54

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateBaseBackup](https://cloud.tencent.com/document/api/409/89028)

	* 新增入参：BackupMethod

* [ModifyBackupPlan](https://cloud.tencent.com/document/api/409/68067)

	* 新增入参：BackupMethod




## 向量数据库(vdb) 版本：2023-06-16

### 第 17 次发布

发布时间：2026-06-12 02:35:31

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [InstanceInfo](https://cloud.tencent.com/document/api/1709/106757#InstanceInfo)

	* 新增成员：UpgradeVersion, IsInternal

* [Network](https://cloud.tencent.com/document/api/1709/106757#Network)

	* 新增成员：IsSSL




## 云点播(vod) 版本：2024-07-18



## 云点播(vod) 版本：2018-07-17

### 第 266 次发布

发布时间：2026-06-12 02:36:21

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [AiTryOnConfig](https://cloud.tencent.com/document/api/266/31773#AiTryOnConfig)

修改数据结构：

* [AigcImageSceneInfo](https://cloud.tencent.com/document/api/266/31773#AigcImageSceneInfo)

	* 新增成员：AiTryOnConfig




## 数据开发治理平台 WeData(wedata) 版本：2025-08-06



## 数据开发治理平台 WeData(wedata) 版本：2021-08-20

### 第 195 次发布

发布时间：2026-06-12 02:48:33

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [InstanceOpsDto](https://cloud.tencent.com/document/api/1267/76336#InstanceOpsDto)

	* 新增成员：ProxyTaskId, WorkflowRunName, ProxyTaskType

* [ManualTriggerRecordOpsDto](https://cloud.tencent.com/document/api/1267/76336#ManualTriggerRecordOpsDto)

	* 新增成员：TriggerSource, TriggerSourceId, ParentSpTaskId, ParentSpInstanceName, ParentSpInstanceDataTime, ScheduleTimeList

* [WorkflowExtOpsDto](https://cloud.tencent.com/document/api/1267/76336#WorkflowExtOpsDto)

	* 新增成员：NestedBySpTaskIds




