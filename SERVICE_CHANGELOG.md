# Release v1.3.209

## 腾讯云智能体开发平台(adp) 版本：2026-05-20

### 第 35 次发布

发布时间：2026-11-06 01:03:32

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateRelease](https://cloud.tencent.com/document/api/1759/132538)

	* 新增入参：A2AConfig


新增数据结构：

* [A2AConfig](https://cloud.tencent.com/document/api/1759/132545#A2AConfig)

修改数据结构：

* [App](https://cloud.tencent.com/document/api/1759/132545#App)

	* 新增成员：A2AConfig

* [ReleaseSummary](https://cloud.tencent.com/document/api/1759/132545#ReleaseSummary)

	* 新增成员：A2AConfig




## Agent 沙箱服务(ags) 版本：2025-09-20

### 第 32 次发布

发布时间：2026-11-06 01:04:49

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribePreCacheImageTask](https://cloud.tencent.com/document/api/1814/127507)

	* 新增出参：ReadyTime


修改数据结构：

* [PreCacheImage](https://cloud.tencent.com/document/api/1814/124823#PreCacheImage)

	* 新增成员：ReadyTime




## 运维安全中心（堡垒机）(bh) 版本：2023-04-18

### 第 40 次发布

发布时间：2026-11-06 01:10:09

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateResource](https://cloud.tencent.com/document/api/1025/100130)

	* 新增入参：IOAResource, PackageIOAUserCount, PackageIOABandwidth

* [ModifyResource](https://cloud.tencent.com/document/api/1025/100129)

	* 新增入参：IOAResource, PackageIOAUserCount, PackageIOABandwidth




## 云安全一体化平台(csip) 版本：2022-11-21

### 第 116 次发布

发布时间：2026-11-06 01:28:11

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [VulFixableMachineAction](https://cloud.tencent.com/document/api/664/90825#VulFixableMachineAction)

修改数据结构：

* [HostKBRisk](https://cloud.tencent.com/document/api/664/90825#HostKBRisk)

	* 新增成员：IsNew, FixTaskId

* [HostVulRisk](https://cloud.tencent.com/document/api/664/90825#HostVulRisk)

	* 新增成员：FixTaskId, SupportAutoFix

* [KBUpdateMachineItem](https://cloud.tencent.com/document/api/664/90825#KBUpdateMachineItem)

	* 新增成员：HostStatus, NotFixableAction

* [VPRExplainDimensionItem](https://cloud.tencent.com/document/api/664/90825#VPRExplainDimensionItem)

	* 新增成员：Level

* [VulDetailInfo](https://cloud.tencent.com/document/api/664/90825#VulDetailInfo)

	* 新增成员：VPRExplainInfo, ScanTaskId, AffectAsset

* [VulFixTaskDetailItem](https://cloud.tencent.com/document/api/664/90825#VulFixTaskDetailItem)

	* 新增成员：ErrorCode, ErrorReason, ErrorSolution, RawErrorMessage, MachineType, SnapshotReused

* [VulFixableMachineItem](https://cloud.tencent.com/document/api/664/90825#VulFixableMachineItem)

	* 新增成员：HostStatus, NotFixableReasonCode, NotFixableAction

* [VulFixedItem](https://cloud.tencent.com/document/api/664/90825#VulFixedItem)

	* 新增成员：RepairMethod

* [VulHostBriefInfo](https://cloud.tencent.com/document/api/664/90825#VulHostBriefInfo)

	* 新增成员：VPRExplainInfo




## 云服务器(cvm) 版本：2017-03-12

### 第 174 次发布

发布时间：2026-11-06 01:34:46

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [Instance](https://cloud.tencent.com/document/api/213/15753#Instance)

	* 新增成员：SuperNodeId




## TDSQL-C MySQL 版(cynosdb) 版本：2019-01-07

### 第 202 次发布

发布时间：2026-11-06 01:39:19

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeInstanceServerlessQuota](https://cloud.tencent.com/document/api/1003/140604)
* [SetInstanceServerlessQuota](https://cloud.tencent.com/document/api/1003/140603)

修改接口：

* [CreateClusters](https://cloud.tencent.com/document/api/1003/48087)

	* 新增入参：ServerlessQuotaConfigs


新增数据结构：

* [ClusterServerlessQuotaConfig](https://cloud.tencent.com/document/api/1003/48097#ClusterServerlessQuotaConfig)
* [InstanceServerlessQuotaInfo](https://cloud.tencent.com/document/api/1003/48097#InstanceServerlessQuotaInfo)
* [ServerlessQuotaConfig](https://cloud.tencent.com/document/api/1003/48097#ServerlessQuotaConfig)

修改数据结构：

* [ClusterInstanceDetail](https://cloud.tencent.com/document/api/1003/48097#ClusterInstanceDetail)

	* 新增成员：ServerlessQuotaInfo

* [CynosdbInstance](https://cloud.tencent.com/document/api/1003/48097#CynosdbInstance)

	* 新增成员：ServerlessQuotaInfo

* [CynosdbInstanceDetail](https://cloud.tencent.com/document/api/1003/48097#CynosdbInstanceDetail)

	* 新增成员：ServerlessQuotaInfo




## 腾讯云数据分析智能体(dataagent) 版本：2025-05-13

### 第 26 次发布

发布时间：2026-11-06 01:41:20

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DeleteDataAgentSession](https://cloud.tencent.com/document/api/1800/125013)

	* 新增入参：SurfaceContext

* [QueryUserSessionDetail](https://cloud.tencent.com/document/api/1800/137839)

	* 新增入参：SurfaceContext

	* 新增出参：SessionTargetInfo


新增数据结构：

* [SessionTargetInfo](https://cloud.tencent.com/document/api/1800/125016#SessionTargetInfo)
* [SurfaceContext](https://cloud.tencent.com/document/api/1800/125016#SurfaceContext)



## 数据湖计算 DLC(dlc) 版本：2021-01-25

### 第 189 次发布

发布时间：2026-11-06 01:44:58

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [AddDeployment](https://cloud.tencent.com/document/api/1342/137397)

	* 新增入参：LoraBindings, CodePackageOriginalName, CodePackageUri

	* 新增出参：MinReplicas, MaxReplicas, AutoscalerOptions, RayDashboardUrl, CodePackageUri, CodePackageOriginalName, LoraBindings

* [CreateBenchmarkTask](https://cloud.tencent.com/document/api/1342/137403)

	* 新增入参：ModelIdentifier

	* 新增出参：ModelIdentifier

* [CreateInferenceModel](https://cloud.tencent.com/document/api/1342/135526)

	* 新增入参：ModelSource, RegisterMeta, FileList, BaseModelUid, BaseModelVersionUid, ModelCategory

	* 新增出参：ModelSource, ModelCategory, BaseModelUid, LatestVersionUid, BaseModelVersionUid, BaseModelName, BaseModelVersion

* [CreateInferenceService](https://cloud.tencent.com/document/api/1342/135625)

	* 新增入参：LoraBindings

* [CreateMlflowServer](https://cloud.tencent.com/document/api/1342/137414)

	* 新增入参：ServerArgs, Description

* [CreateModelVersion](https://cloud.tencent.com/document/api/1342/135618)

	* 新增入参：ModelSource, RegisterMeta, FileList

	* 新增出参：RegisterMeta, VersionUid, LoraMeta, PoolImportStatus

* [DescribeTrainingCheckpoints](https://cloud.tencent.com/document/api/1342/137376)

	* 新增出参：ArtifactRootUri

* [GetInferenceModel](https://cloud.tencent.com/document/api/1342/135525)

	* 新增出参：ModelCategory, BaseModelUid, LatestVersionUid, BaseModelVersionUid, BaseModelName, BaseModelVersion, ModelSource

* [RerunBenchmarkTask](https://cloud.tencent.com/document/api/1342/137399)

	* 新增出参：ModelIdentifier

* [RestartDeployment](https://cloud.tencent.com/document/api/1342/137390)

	* 新增出参：LoraBindings

* [ResumeTrainingJobInstance](https://cloud.tencent.com/document/api/1342/137370)

	* 新增入参：ResumeCheckpointSubDir

* [StopBenchmarkTask](https://cloud.tencent.com/document/api/1342/137398)

	* 新增出参：ModelIdentifier

* [StopDeployment](https://cloud.tencent.com/document/api/1342/137389)

	* 新增出参：LoraBindings

* [UpdateDeployment](https://cloud.tencent.com/document/api/1342/137388)

	* 新增入参：CodePackageUri, CodePackageOriginalName, CrOverrides, HeadHighAvailabilityEnabled, LoraBindings, RuntimeEnv

	* 新增出参：LoraBindings

* [UpdateInferenceModel](https://cloud.tencent.com/document/api/1342/135523)

	* 新增出参：ModelCategory, BaseModelUid, LatestVersionUid, BaseModelVersionUid, BaseModelName, BaseModelVersion, ModelSource


新增数据结构：

* [LoraAdapterBinding](https://cloud.tencent.com/document/api/1342/53778#LoraAdapterBinding)
* [LoraItem](https://cloud.tencent.com/document/api/1342/53778#LoraItem)
* [LoraVersionMeta](https://cloud.tencent.com/document/api/1342/53778#LoraVersionMeta)
* [RegisterMeta](https://cloud.tencent.com/document/api/1342/53778#RegisterMeta)

修改数据结构：

* [InferenceModelInfo](https://cloud.tencent.com/document/api/1342/53778#InferenceModelInfo)

	* 新增成员：ModelSource, ModelCategory, BaseModelUid, LatestVersionUid, BaseModelVersionUid, BaseModelName, BaseModelVersion

* [MlFlowServerInfo](https://cloud.tencent.com/document/api/1342/53778#MlFlowServerInfo)

	* 新增成员：ServerArgs, Description, SubAccountUin, SubAccountName, LastProbeSuccessTime, ConsecutiveProbeFailures, LastProbeError, TotalInstanceCount, RunningInstanceCount

* [TrainingJobInstance](https://cloud.tencent.com/document/api/1342/53778#TrainingJobInstance)

	* 新增成员：SubAccountUin, SubAccountName, NetworkConfig, ResumeSourceInstanceId

* [TrainingTuningParams](https://cloud.tencent.com/document/api/1342/53778#TrainingTuningParams)

	* 新增成员：PPOMaxTokenLenPerGpu, LogProbMaxTokenLenPerGpu, ActorParamOffload, ActorOptimizerOffload, RefParamOffload, PPOMicroBatchSizePerGpu, RollOutLogProbMicroBatchSizePerGpu, GRPORawOverrides




## 人脸识别(iai) 版本：2020-03-03

### 第 15 次发布

发布时间：2026-11-06 02:01:51

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateFace](https://cloud.tencent.com/document/api/867/45016)

	* 新增入参：Encryption, EncryptedBody

* [CreatePerson](https://cloud.tencent.com/document/api/867/45014)

	* 新增入参：Encryption, EncryptedBody

* [SearchFaces](https://cloud.tencent.com/document/api/867/44994)

	* 新增入参：Encryption, EncryptedBody

* [SearchPersons](https://cloud.tencent.com/document/api/867/44992)

	* 新增入参：Encryption, EncryptedBody


新增数据结构：

* [Encryption](https://cloud.tencent.com/document/api/867/45020#Encryption)



## 人脸识别(iai) 版本：2018-03-01



## 多网聚合加速(mna) 版本：2021-01-19

### 第 40 次发布

发布时间：2026-11-06 02:14:41

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DownloadActiveDeviceCount](https://cloud.tencent.com/document/api/1385/116202)

	* 新增入参：ApplicationId

* [GetActiveDeviceCount](https://cloud.tencent.com/document/api/1385/116201)

	* 新增入参：ApplicationId




## 文字识别(ocr) 版本：2018-11-19

### 第 276 次发布

发布时间：2026-11-06 02:21:28

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [QueryEnterpriseBasicInfo](https://cloud.tencent.com/document/api/866/140605)



## 云开发 CloudBase(tcb) 版本：2018-06-08

### 第 178 次发布

发布时间：2026-11-06 02:33:44

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeEnvInfo](https://cloud.tencent.com/document/api/876/140606)



## TDSQL(tdmysql) 版本：2021-11-22

### 第 19 次发布

发布时间：2026-11-06 02:39:38

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateDBInstances](https://cloud.tencent.com/document/api/1376/132036)

	* 新增入参：PasswordComplexity

* [DescribeDBCharsets](https://cloud.tencent.com/document/api/1376/137573)

	* 新增出参：CharsetCollations

* [DescribeDBInstanceDetail](https://cloud.tencent.com/document/api/1376/132035)

	* 新增出参：SSLStatus, Collation

* [DescribeSpecs](https://cloud.tencent.com/document/api/1376/132042)

	* 新增入参：InstanceId

* [ResetDbaAdminPrivileges](https://cloud.tencent.com/document/api/1376/137570)

	* 新增入参：Host


新增数据结构：

* [CharsetCollationItem](https://cloud.tencent.com/document/api/1376/128305#CharsetCollationItem)
* [PasswordComplexity](https://cloud.tencent.com/document/api/1376/128305#PasswordComplexity)



## 边缘安全加速平台(teo) 版本：2022-09-01

### 第 171 次发布

发布时间：2026-11-06 02:40:47

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [AccessURLRedirectPort](https://cloud.tencent.com/document/api/1552/80721#AccessURLRedirectPort)

修改数据结构：

* [AccessURLRedirectParameters](https://cloud.tencent.com/document/api/1552/80721#AccessURLRedirectParameters)

	* 新增成员：Port




## 边缘安全加速平台(teo) 版本：2022-01-06



## TI-ONE 训练平台(tione) 版本：2021-11-11

### 第 145 次发布

发布时间：2026-11-06 02:43:30

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateNotebook](https://cloud.tencent.com/document/api/851/95658)

	* 新增入参：ResourceSupplyAttribute, Queues


修改数据结构：

* [SmartRouterResourceSpec](https://cloud.tencent.com/document/api/851/75051#SmartRouterResourceSpec)

	* 新增成员：ResourceGroupSWType




## TI-ONE 训练平台(tione) 版本：2019-10-22



## 容器服务(tke) 版本：2022-05-01



## 容器服务(tke) 版本：2018-05-25

### 第 240 次发布

发布时间：2026-11-06 02:44:44

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [AddExistedInstances](https://cloud.tencent.com/document/api/457/31865)

	* 新增入参：SkipOptions




## 云点播(vod) 版本：2024-07-18



## 云点播(vod) 版本：2018-07-17

### 第 295 次发布

发布时间：2026-11-06 02:52:41

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CloneVoiceAsync](https://cloud.tencent.com/document/api/266/137251)

	* 新增入参：ModelName

* [CloneVoiceSync](https://cloud.tencent.com/document/api/266/137492)

	* 新增入参：ModelName

* [DeleteVoice](https://cloud.tencent.com/document/api/266/137250)

	* 新增入参：ModelName

* [DescribeVoices](https://cloud.tencent.com/document/api/266/137249)

	* 新增入参：ModelName

	* 新增出参：VoiceSlot

* [DesignVoiceAsync](https://cloud.tencent.com/document/api/266/137248)

	* 新增入参：ModelName

* [TextToSpeechAsync](https://cloud.tencent.com/document/api/266/137247)

	* 新增入参：ModelName, ModelVersion

* [TextToSpeechSync](https://cloud.tencent.com/document/api/266/137491)

	* 新增入参：ModelName, ModelVersion

* [UpdateVoice](https://cloud.tencent.com/document/api/266/137246)

	* 新增入参：ModelName


新增数据结构：

* [VoiceSlotInfo](https://cloud.tencent.com/document/api/266/31773#VoiceSlotInfo)



