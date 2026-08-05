# Release v1.3.153

## 云联络中心(ccc) 版本：2020-02-10

### 第 132 次发布

发布时间：2026-08-06 01:14:47

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateAICall](https://cloud.tencent.com/document/api/679/111211)

	* 新增入参：TransferToAgentEnable, TransferToAgentItems, AcquireTimeoutSecond, CustomSTTConfig


新增数据结构：

* [TransferToAgentItem](https://cloud.tencent.com/document/api/679/47715#TransferToAgentItem)



## 消息队列 CKafka 版(ckafka) 版本：2019-08-19

### 第 149 次发布

发布时间：2026-08-06 01:19:19

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateThrottleRule](https://cloud.tencent.com/document/api/597/135609)
* [DeleteThrottleRule](https://cloud.tencent.com/document/api/597/135608)
* [DescribeThrottleRules](https://cloud.tencent.com/document/api/597/135607)
* [ModifyThrottleRule](https://cloud.tencent.com/document/api/597/135606)

新增数据结构：

* [ThrottleRuleDetail](https://cloud.tencent.com/document/api/597/40861#ThrottleRuleDetail)
* [ThrottleRuleResult](https://cloud.tencent.com/document/api/597/40861#ThrottleRuleResult)



## 日志服务(cls) 版本：2020-10-16

### 第 171 次发布

发布时间：2026-08-06 01:21:06

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DeleteLog](https://cloud.tencent.com/document/api/614/135612)
* [ModifyLog](https://cloud.tencent.com/document/api/614/135611)



## 云数据库独享集群(dbdc) 版本：2020-10-29

### 第 12 次发布

发布时间：2026-08-06 01:29:39

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [DBCustomClusterNode](https://cloud.tencent.com/document/api/1322/74754#DBCustomClusterNode)

	* 新增成员：SecurityGroupIds

* [DBCustomNode](https://cloud.tencent.com/document/api/1322/74754#DBCustomNode)

	* 新增成员：SecurityGroupIds




## 数据湖计算 DLC(dlc) 版本：2021-01-25

### 第 171 次发布

发布时间：2026-08-06 01:30:30

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateInferenceService](https://cloud.tencent.com/document/api/1342/135625)
* [CreateModelVersion](https://cloud.tencent.com/document/api/1342/135618)
* [GetInferenceService](https://cloud.tencent.com/document/api/1342/135624)
* [GetModelConfig](https://cloud.tencent.com/document/api/1342/135617)
* [GetModelFiles](https://cloud.tencent.com/document/api/1342/135616)
* [GetModelReadme](https://cloud.tencent.com/document/api/1342/135615)
* [ListInferenceEngines](https://cloud.tencent.com/document/api/1342/135623)
* [ListInferenceServices](https://cloud.tencent.com/document/api/1342/135622)
* [ListModelVersions](https://cloud.tencent.com/document/api/1342/135614)
* [QueryDashboardOverview](https://cloud.tencent.com/document/api/1342/135629)
* [QueryDashboardServiceList](https://cloud.tencent.com/document/api/1342/135628)
* [QueryMonitorOverview](https://cloud.tencent.com/document/api/1342/135627)
* [RestartInferenceService](https://cloud.tencent.com/document/api/1342/135621)
* [StopInferenceService](https://cloud.tencent.com/document/api/1342/135620)

新增数据结构：

* [CpuSummaryItem](https://cloud.tencent.com/document/api/1342/53778#CpuSummaryItem)
* [EngineCapabilities](https://cloud.tencent.com/document/api/1342/53778#EngineCapabilities)
* [FileNode](https://cloud.tencent.com/document/api/1342/53778#FileNode)
* [GpuSummaryItem](https://cloud.tencent.com/document/api/1342/53778#GpuSummaryItem)
* [InferenceEngineInfo](https://cloud.tencent.com/document/api/1342/53778#InferenceEngineInfo)
* [InferenceServiceInfo](https://cloud.tencent.com/document/api/1342/53778#InferenceServiceInfo)
* [LinkedServiceInfo](https://cloud.tencent.com/document/api/1342/53778#LinkedServiceInfo)
* [MetricsData](https://cloud.tencent.com/document/api/1342/53778#MetricsData)
* [ModelVersionInfo](https://cloud.tencent.com/document/api/1342/53778#ModelVersionInfo)
* [OverviewItem](https://cloud.tencent.com/document/api/1342/53778#OverviewItem)
* [ParallelKeyMapping](https://cloud.tencent.com/document/api/1342/53778#ParallelKeyMapping)
* [ReplicaInfo](https://cloud.tencent.com/document/api/1342/53778#ReplicaInfo)
* [ServiceMetricsItem](https://cloud.tencent.com/document/api/1342/53778#ServiceMetricsItem)



## 数据传输服务(dts) 版本：2021-12-06

### 第 60 次发布

发布时间：2026-08-06 01:34:50

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DeleteConsumerGroup](https://cloud.tencent.com/document/api/571/102948)

	* 新增入参：BackendJobId


修改数据结构：

* [SubscribeInfo](https://cloud.tencent.com/document/api/571/82108#SubscribeInfo)

	* 新增成员：ConsumerRoutePhase




## 数据传输服务(dts) 版本：2018-03-30



## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 315 次发布

发布时间：2026-08-06 01:37:40

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateSeal](https://cloud.tencent.com/document/api/1323/94136)

	* 新增入参：SubSealType




## 腾讯电子签（基础版）(essbasic) 版本：2021-05-26

### 第 270 次发布

发布时间：2026-08-06 01:38:38

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateSealByImage](https://cloud.tencent.com/document/api/1420/73067)

	* 新增入参：SubSealType




## 腾讯电子签（基础版）(essbasic) 版本：2020-12-22



## 媒体处理(mps) 版本：2019-06-12

### 第 229 次发布

发布时间：2026-08-06 01:51:41

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateDocToVideoTask](https://cloud.tencent.com/document/api/862/134552)

	* 新增入参：ResourceId


修改数据结构：

* [AiCutoutConfig](https://cloud.tencent.com/document/api/862/37615#AiCutoutConfig)

	* 新增成员：Model




## 流计算 Oceanus(oceanus) 版本：2019-04-22

### 第 90 次发布

发布时间：2026-08-06 01:53:40

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [JobV1](https://cloud.tencent.com/document/api/849/52010#JobV1)

	* 新增成员：HealthScore, LastDiagnoseTime, ManagerUin




## 云数据库Redis(redis) 版本：2018-04-12

### 第 106 次发布

发布时间：2026-08-06 01:56:17

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CloneInstances](https://cloud.tencent.com/document/api/239/89391)

	* 新增入参：ProductVersion




## 云开发 CloudBase(tcb) 版本：2018-06-08

### 第 155 次发布

发布时间：2026-08-06 02:00:39

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [VerifyHTTPServiceRoute](https://cloud.tencent.com/document/api/876/135630)

新增数据结构：

* [VerifyHTTPServiceRouteCheckItem](https://cloud.tencent.com/document/api/876/34822#VerifyHTTPServiceRouteCheckItem)



## TokenHub(tokenhub) 版本：2026-03-22

### 第 17 次发布

发布时间：2026-08-06 02:09:06

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [Model](https://cloud.tencent.com/document/api/1823/132279#Model)

	* 新增成员：DiscontinuedAt




## 实时音视频(trtc) 版本：2019-07-22

### 第 149 次发布

发布时间：2026-08-06 02:10:14

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateLiveStreamModeration](https://cloud.tencent.com/document/api/647/135224)

	* <font color="#dd0000">**修改入参**：</font>DataId


新增数据结构：

* [AgoraParam](https://cloud.tencent.com/document/api/647/44055#AgoraParam)

修改数据结构：

* [Input](https://cloud.tencent.com/document/api/647/44055#Input)

	* 新增成员：AgoraParam

	* <font color="#dd0000">**修改成员**：</font>Url




