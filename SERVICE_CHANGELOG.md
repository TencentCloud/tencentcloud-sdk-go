# Release v1.3.151

## 腾讯云智能体开发平台(adp) 版本：2026-05-20

### 第 13 次发布

发布时间：2026-08-03 11:34:14

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除接口**：</font>

* CreateTimerTask
* DeleteTimerTask
* DescribeTimerTask
* DescribeTimerTaskRunLogList
* DescribeTimerTaskSummaryList
* MarkAppTriggerRunLogRead
* MarkTimerTaskRunLogRead
* ModifyTimerTask
* PauseTimerTask
* ResumeTimerTask
* RunTimerTaskNow

修改接口：

* [CreateAppTrigger](https://cloud.tencent.com/document/api/1759/135012)

	* 新增入参：Scope, UserId

* [DeleteAppTrigger](https://cloud.tencent.com/document/api/1759/135010)

	* 新增入参：Scope, UserId

* [DescribeAppTrigger](https://cloud.tencent.com/document/api/1759/135008)

	* 新增入参：Scope, UserId

* [DescribeAppTriggerInstance](https://cloud.tencent.com/document/api/1759/135007)

	* 新增入参：Scope, UserId

* [DescribeAppTriggerRunLogList](https://cloud.tencent.com/document/api/1759/135006)

	* 新增入参：Scope, UserId

* [DescribeAppTriggerSummaryList](https://cloud.tencent.com/document/api/1759/135005)

	* 新增入参：Scope, UserId

* [DescribeConversationMessageList](https://cloud.tencent.com/document/api/1759/132517)

	* 新增出参：RecordSummaryList

* [ModifyAppTrigger](https://cloud.tencent.com/document/api/1759/134999)

	* 新增入参：Scope, UserId

* [PauseAppTrigger](https://cloud.tencent.com/document/api/1759/134997)

	* 新增入参：Scope, UserId

* [ResumeAppTrigger](https://cloud.tencent.com/document/api/1759/134995)

	* 新增入参：Scope, UserId

* [RunAppTriggerNow](https://cloud.tencent.com/document/api/1759/134993)

	* 新增入参：Scope, UserId


新增数据结构：

* [ConversationRecordErrorInfo](https://cloud.tencent.com/document/api/1759/132545#ConversationRecordErrorInfo)
* [ConversationRecordSummary](https://cloud.tencent.com/document/api/1759/132545#ConversationRecordSummary)
* [ConversationRecordTimeUsage](https://cloud.tencent.com/document/api/1759/132545#ConversationRecordTimeUsage)
* [ConversationRecordTokenUsage](https://cloud.tencent.com/document/api/1759/132545#ConversationRecordTokenUsage)

<font color="#dd0000">**删除数据结构**：</font>

* TimerConfig
* TimerProfile
* TimerStatus
* TimerTask
* TimerTaskSummary

修改数据结构：

* [AppTrigger](https://cloud.tencent.com/document/api/1759/132545#AppTrigger)

	* 新增成员：Scope, UserId

* [AppTriggerInstance](https://cloud.tencent.com/document/api/1759/132545#AppTriggerInstance)

	* 新增成员：Scope, UserId

* [AppTriggerRunLog](https://cloud.tencent.com/document/api/1759/132545#AppTriggerRunLog)

	* 新增成员：Scope, UserId

* [AppTriggerSummary](https://cloud.tencent.com/document/api/1759/132545#AppTriggerSummary)

	* 新增成员：Scope, UserId

* [TriggerStatus](https://cloud.tencent.com/document/api/1759/132545#TriggerStatus)

	* 新增成员：Scope, UserId




## Agent 沙箱服务(ags) 版本：2025-09-20

### 第 18 次发布

发布时间：2026-08-04 01:05:36

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [PauseSandboxInstance](https://cloud.tencent.com/document/api/1814/127876)

	* 新增入参：Memory

* [ResumeSandboxInstance](https://cloud.tencent.com/document/api/1814/127875)

	* 新增入参：Timeout


### 第 17 次发布

发布时间：2026-08-03 20:48:43

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreatePreCacheImageTask](https://cloud.tencent.com/document/api/1814/127508)

	* <font color="#dd0000">**删除入参**：</font>TimeoutMinutes


新增数据结构：

* [AgentBucketStorageSource](https://cloud.tencent.com/document/api/1814/124823#AgentBucketStorageSource)

修改数据结构：

* [StorageSource](https://cloud.tencent.com/document/api/1814/124823#StorageSource)

	* 新增成员：AgentBucket




## 日志服务(cls) 版本：2020-10-16

### 第 170 次发布

发布时间：2026-08-04 01:28:30

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [ToolCall](https://cloud.tencent.com/document/api/614/56471#ToolCall)

	* 新增成员：ThoughtSignature




## 云原生智能网关(cngw) 版本：2023-04-18

### 第 4 次发布

发布时间：2026-08-04 01:31:14

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeCloudNativeAPIGatewayMCPToolsFromFile](https://cloud.tencent.com/document/api/1826/135400)
* [UpdateCloudNativeAPIGatewayMCPTools](https://cloud.tencent.com/document/api/1826/135399)

修改接口：

* [CreateCloudNativeAPIGatewayLLMModelService](https://cloud.tencent.com/document/api/1826/133132)

	* 新增入参：SourceId, Namespace, ServiceName, Protocol

* [CreateCloudNativeAPIGatewaySecretKey](https://cloud.tencent.com/document/api/1826/133141)

	* 新增入参：Provider

* [DescribeCloudNativeAPIGatewayLLMModelAPIs](https://cloud.tencent.com/document/api/1826/133112)

	* 新增入参：ConsumerId

* [ModifyCloudNativeAPIGatewayLLMModelService](https://cloud.tencent.com/document/api/1826/133128)

	* 新增入参：SourceId, Namespace, ServiceName, Protocol


新增数据结构：

* [CNAPIGwMCPToolPreview](https://cloud.tencent.com/document/api/1826/133161#CNAPIGwMCPToolPreview)
* [CNAPIGwParseMCPToolsResult](https://cloud.tencent.com/document/api/1826/133161#CNAPIGwParseMCPToolsResult)

修改数据结构：

* [AIGWLLMQuotaLimit](https://cloud.tencent.com/document/api/1826/133161#AIGWLLMQuotaLimit)

	* 新增成员：ConcurrentCountLimit

* [AIGWOAuthCredentialConfig](https://cloud.tencent.com/document/api/1826/133161#AIGWOAuthCredentialConfig)

	* 新增成员：RedirectURIs

* [AIGWRedisConfig](https://cloud.tencent.com/document/api/1826/133161#AIGWRedisConfig)

	* 新增成员：RedisConfigId, Type

	* <font color="#dd0000">**修改成员**：</font>Host, Port




## TDSQL-C MySQL 版(cynosdb) 版本：2019-01-07

### 第 185 次发布

发布时间：2026-08-04 01:41:44

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateAccounts](https://cloud.tencent.com/document/api/1003/71660)

	* 新增出参：TaskId




## 数据湖计算 DLC(dlc) 版本：2021-01-25

### 第 169 次发布

发布时间：2026-08-04 01:47:11

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CancelRayJob](https://cloud.tencent.com/document/api/1342/135471)
* [CopyJobSpec](https://cloud.tencent.com/document/api/1342/135459)
* [CreateClusterGroup](https://cloud.tencent.com/document/api/1342/135407)
* [CreateJobSpec](https://cloud.tencent.com/document/api/1342/135458)
* [CreateLab](https://cloud.tencent.com/document/api/1342/135449)
* [CreateRayCluster](https://cloud.tencent.com/document/api/1342/135422)
* [CreateResourceConfig](https://cloud.tencent.com/document/api/1342/135428)
* [DeleteClusterGroup](https://cloud.tencent.com/document/api/1342/135406)
* [DeleteJobSpec](https://cloud.tencent.com/document/api/1342/135457)
* [DeleteLab](https://cloud.tencent.com/document/api/1342/135448)
* [DeleteRayCluster](https://cloud.tencent.com/document/api/1342/135421)
* [DeleteRayJob](https://cloud.tencent.com/document/api/1342/135470)
* [DeleteResourceConfig](https://cloud.tencent.com/document/api/1342/135427)
* [DescribeClusterGroup](https://cloud.tencent.com/document/api/1342/135405)
* [DescribeClusterGroupClusters](https://cloud.tencent.com/document/api/1342/135404)
* [GetExampleDetail](https://cloud.tencent.com/document/api/1342/135434)
* [GetJobSpec](https://cloud.tencent.com/document/api/1342/135456)
* [GetLabDetail](https://cloud.tencent.com/document/api/1342/135447)
* [GetLabEvent](https://cloud.tencent.com/document/api/1342/135446)
* [GetLabHistory](https://cloud.tencent.com/document/api/1342/135445)
* [GetLabPodYaml](https://cloud.tencent.com/document/api/1342/135444)
* [GetLabPods](https://cloud.tencent.com/document/api/1342/135443)
* [GetLabServiceUrls](https://cloud.tencent.com/document/api/1342/135442)
* [GetLabYaml](https://cloud.tencent.com/document/api/1342/135441)
* [GetRayCluster](https://cloud.tencent.com/document/api/1342/135420)
* [GetRayClusterEvent](https://cloud.tencent.com/document/api/1342/135419)
* [GetRayClusterHistory](https://cloud.tencent.com/document/api/1342/135418)
* [GetRayClusterPodYaml](https://cloud.tencent.com/document/api/1342/135417)
* [GetRayClusterPods](https://cloud.tencent.com/document/api/1342/135416)
* [GetRayClusterYaml](https://cloud.tencent.com/document/api/1342/135415)
* [GetRayJob](https://cloud.tencent.com/document/api/1342/135469)
* [GetRayJobEvent](https://cloud.tencent.com/document/api/1342/135468)
* [GetRayJobEventLog](https://cloud.tencent.com/document/api/1342/135467)
* [GetRayJobHistory](https://cloud.tencent.com/document/api/1342/135466)
* [GetRayJobPodYaml](https://cloud.tencent.com/document/api/1342/135465)
* [GetRayJobPods](https://cloud.tencent.com/document/api/1342/135464)
* [GetRayJobYaml](https://cloud.tencent.com/document/api/1342/135463)
* [GetResourceConfig](https://cloud.tencent.com/document/api/1342/135426)
* [ListClusterGroups](https://cloud.tencent.com/document/api/1342/135403)
* [ListExampleCategories](https://cloud.tencent.com/document/api/1342/135433)
* [ListExampleDifficulties](https://cloud.tencent.com/document/api/1342/135432)
* [ListExampleTags](https://cloud.tencent.com/document/api/1342/135431)
* [ListExamples](https://cloud.tencent.com/document/api/1342/135430)
* [ListJobSpecs](https://cloud.tencent.com/document/api/1342/135455)
* [ListJobsBySpec](https://cloud.tencent.com/document/api/1342/135454)
* [ListLabs](https://cloud.tencent.com/document/api/1342/135440)
* [ListRayClusterJobs](https://cloud.tencent.com/document/api/1342/135414)
* [ListRayClusters](https://cloud.tencent.com/document/api/1342/135413)
* [ListRayJobs](https://cloud.tencent.com/document/api/1342/135462)
* [ListResourceConfigs](https://cloud.tencent.com/document/api/1342/135425)
* [ModifyClusterPriority](https://cloud.tencent.com/document/api/1342/135412)
* [ModifyLabPriority](https://cloud.tencent.com/document/api/1342/135439)
* [RunJobSpec](https://cloud.tencent.com/document/api/1342/135453)
* [StartLab](https://cloud.tencent.com/document/api/1342/135438)
* [StartRayCluster](https://cloud.tencent.com/document/api/1342/135411)
* [StopLab](https://cloud.tencent.com/document/api/1342/135437)
* [StopRayCluster](https://cloud.tencent.com/document/api/1342/135410)
* [UpdateClusterGroup](https://cloud.tencent.com/document/api/1342/135402)
* [UpdateJobSpec](https://cloud.tencent.com/document/api/1342/135452)
* [UpdateJobSpecPriority](https://cloud.tencent.com/document/api/1342/135451)
* [UpdateLab](https://cloud.tencent.com/document/api/1342/135436)
* [UpdateRayCluster](https://cloud.tencent.com/document/api/1342/135409)
* [UpdateRayJobPriority](https://cloud.tencent.com/document/api/1342/135461)
* [UpdateResourceConfig](https://cloud.tencent.com/document/api/1342/135424)

新增数据结构：

* [ClusterGroup](https://cloud.tencent.com/document/api/1342/53778#ClusterGroup)
* [ClusterPod](https://cloud.tencent.com/document/api/1342/53778#ClusterPod)
* [Env](https://cloud.tencent.com/document/api/1342/53778#Env)
* [EventItem](https://cloud.tencent.com/document/api/1342/53778#EventItem)
* [EventLogItem](https://cloud.tencent.com/document/api/1342/53778#EventLogItem)
* [ExampleCategories](https://cloud.tencent.com/document/api/1342/53778#ExampleCategories)
* [ExampleDifficulties](https://cloud.tencent.com/document/api/1342/53778#ExampleDifficulties)
* [ExampleEntity](https://cloud.tencent.com/document/api/1342/53778#ExampleEntity)
* [ExampleTag](https://cloud.tencent.com/document/api/1342/53778#ExampleTag)
* [HeadSpecDTO](https://cloud.tencent.com/document/api/1342/53778#HeadSpecDTO)
* [JobPodEntity](https://cloud.tencent.com/document/api/1342/53778#JobPodEntity)
* [JobSpec](https://cloud.tencent.com/document/api/1342/53778#JobSpec)
* [JobStatusHistory](https://cloud.tencent.com/document/api/1342/53778#JobStatusHistory)
* [LabResponse](https://cloud.tencent.com/document/api/1342/53778#LabResponse)
* [Label](https://cloud.tencent.com/document/api/1342/53778#Label)
* [PersistentWorkDir](https://cloud.tencent.com/document/api/1342/53778#PersistentWorkDir)
* [RayClusterEntity](https://cloud.tencent.com/document/api/1342/53778#RayClusterEntity)
* [RayClusterHistory](https://cloud.tencent.com/document/api/1342/53778#RayClusterHistory)
* [RayJobEventItem](https://cloud.tencent.com/document/api/1342/53778#RayJobEventItem)
* [RayJobSubmitEntity](https://cloud.tencent.com/document/api/1342/53778#RayJobSubmitEntity)
* [ResourceConfig](https://cloud.tencent.com/document/api/1342/53778#ResourceConfig)
* [SortField](https://cloud.tencent.com/document/api/1342/53778#SortField)
* [Tag](https://cloud.tencent.com/document/api/1342/53778#Tag)
* [TypeKVPair](https://cloud.tencent.com/document/api/1342/53778#TypeKVPair)
* [WorkerSpecDTO](https://cloud.tencent.com/document/api/1342/53778#WorkerSpecDTO)



## 全球加速(ga2) 版本：2025-01-15

### 第 13 次发布

发布时间：2026-08-04 02:03:02

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [AcceleratorAreas](https://cloud.tencent.com/document/api/1817/130045#AcceleratorAreas)

	* <font color="#dd0000">**修改成员**：</font>Bandwidth




## 边缘安全加速平台(teo) 版本：2022-09-01

### 第 154 次发布

发布时间：2026-08-04 02:51:56

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [ExceptionRule](https://cloud.tencent.com/document/api/1552/80721#ExceptionRule)

	* 新增成员：WebSecuritySubmodulesForException

* [IPSSLConfig](https://cloud.tencent.com/document/api/1552/80721#IPSSLConfig)

	* 新增成员：ZoneId




## 边缘安全加速平台(teo) 版本：2022-01-06



## TI-ONE 训练平台(tione) 版本：2021-11-11

### 第 128 次发布

发布时间：2026-08-04 02:55:09

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DeleteDataset](https://cloud.tencent.com/document/api/851/75080)

	* 新增入参：TiProjectId

* [DescribeAnnotatedTaskList](https://cloud.tencent.com/document/api/851/131132)

	* 新增入参：TiProjectId




## TI-ONE 训练平台(tione) 版本：2019-10-22



## 云点播(vod) 版本：2024-07-18



## 云点播(vod) 版本：2018-07-17

### 第 279 次发布

发布时间：2026-08-04 03:07:30

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [SearchMedia](https://cloud.tencent.com/document/api/266/31813)

	* 新增入参：KnowledgeBases




