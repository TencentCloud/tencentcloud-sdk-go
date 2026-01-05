# Release v1.3.27

## 腾讯混元生3D(ai3d) 版本：2025-05-13

### 第 8 次发布

发布时间：2026-01-06 01:07:34

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [SubmitHunyuanTo3DProJob](https://cloud.tencent.com/document/api/1804/123447)

	* 新增入参：ResultFormat




## 负载均衡(clb) 版本：2018-03-17

### 第 145 次发布

发布时间：2026-01-06 01:14:00

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyTargetGroupAttribute](https://cloud.tencent.com/document/api/214/40552)

	* 新增入参：SnatEnable




## 云安全一体化平台(csip) 版本：2022-11-21

### 第 66 次发布

发布时间：2026-01-06 01:15:31

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [AccessKeyAlarm](https://cloud.tencent.com/document/api/664/90825#AccessKeyAlarm)

	* 新增成员：AIStatus, FirstAlarmTimestamp, LastAlarmTimestamp




## 主机安全(cwp) 版本：2018-02-28

### 第 154 次发布

发布时间：2026-01-05 17:13:11

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除接口**：</font>

* CancelIgnoreVul
* DescribeAvailableExpertServiceDetail
* DescribeEmergencyResponseList
* DescribeExpertServiceList
* DescribeExpertServiceOrderList
* DescribeMonthInspectionReport
* DescribeProtectNetList
* IgnoreImpactedHosts

修改接口：

* [DescribeLicenseGeneral](https://cloud.tencent.com/document/api/296/80398)

	* 新增出参：AutoDowngradeSwitch

* [DescribeScanState](https://cloud.tencent.com/document/api/296/60923)

	* 新增出参：KBNumber

* [DescribeScanTaskDetails](https://cloud.tencent.com/document/api/296/58238)

	* 新增出参：PatchInfo

* [DescribeVulFixStatus](https://cloud.tencent.com/document/api/296/99551)

	* 新增入参：KbId

* [ModifyAutoOpenProVersionConfig](https://cloud.tencent.com/document/api/296/19863)

	* 新增入参：ProtectType, AutoDowngradeSwitch

* [ModifyWebHookPolicy](https://cloud.tencent.com/document/api/296/99582)

	* 新增入参：MsgLanguage

* [ModifyWebHookReceiver](https://cloud.tencent.com/document/api/296/99580)

	* 新增入参：Type, SCFRegion, Namespace, FunctionName, FunctionVersion, Alias

* [RetryVulFix](https://cloud.tencent.com/document/api/296/99543)

	* 新增入参：KbId

	* <font color="#dd0000">**修改入参**：</font>VulId

* [ScanVul](https://cloud.tencent.com/document/api/296/57375)

	* 新增入参：KBNumber

* [ScanVulAgain](https://cloud.tencent.com/document/api/296/58236)

	* 新增入参：EventType

	* 新增出参：SuccessCount, BasicVersionCount


新增数据结构：

* [IPAnalyse](https://cloud.tencent.com/document/api/296/19867#IPAnalyse)
* [PatchInfoDetail](https://cloud.tencent.com/document/api/296/19867#PatchInfoDetail)

<font color="#dd0000">**删除数据结构**：</font>

* EmergencyResponseInfo
* ExpertServiceOrderInfo
* MonthInspectionReport
* ProtectNetInfo
* SecurityButlerInfo

修改数据结构：

* [BruteAttackInfo](https://cloud.tencent.com/document/api/296/19867#BruteAttackInfo)

	* 新增成员：IPAnalyse

* [CreateVulFixTaskQuuids](https://cloud.tencent.com/document/api/296/19867#CreateVulFixTaskQuuids)

	* 新增成员：KbId

	* <font color="#dd0000">**修改成员**：</font>VulId

* [HostLoginList](https://cloud.tencent.com/document/api/296/19867#HostLoginList)

	* 新增成员：IPAnalyse

* [NetAttackEvent](https://cloud.tencent.com/document/api/296/19867#NetAttackEvent)

	* 新增成员：RaspOpen, IPAnalyse

* [NetAttackEventInfo](https://cloud.tencent.com/document/api/296/19867#NetAttackEventInfo)

	* 新增成员：IPAnalyse

* [VulEmergentMsgInfo](https://cloud.tencent.com/document/api/296/19867#VulEmergentMsgInfo)

	* 新增成员：KbId, KbNumber

* [VulFixStatusInfo](https://cloud.tencent.com/document/api/296/19867#VulFixStatusInfo)

	* 新增成员：KbId, KbNumber, KbName, PreKbList

* [VulInfoHostInfo](https://cloud.tencent.com/document/api/296/19867#VulInfoHostInfo)

	* 新增成员：AgentStatus

* [VulInfoList](https://cloud.tencent.com/document/api/296/19867#VulInfoList)

	* 新增成员：RaspOpenNodeCount, RaspClosedNodeCount

* [WebHookPolicy](https://cloud.tencent.com/document/api/296/19867#WebHookPolicy)

	* 新增成员：MsgLanguage

* [WebHookReceiver](https://cloud.tencent.com/document/api/296/19867#WebHookReceiver)

	* 新增成员：Type, SCFRegion, Namespace, FunctionName, FunctionVersion, Alias




## 实时互动-教育版(lcic) 版本：2022-08-17

### 第 77 次发布

发布时间：2026-01-06 01:28:13

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [EventDataInfo](https://cloud.tencent.com/document/api/1639/81423#EventDataInfo)

	* 新增成员：Reason




## 媒体处理(mps) 版本：2019-06-12

### 第 169 次发布

发布时间：2026-01-06 01:37:56

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateAigcVideoTask](https://cloud.tencent.com/document/api/862/126965)

	* 新增入参：AdditionalParameters




## 流计算 Oceanus(oceanus) 版本：2019-04-22

### 第 78 次发布

发布时间：2026-01-06 01:42:28

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateJob](https://cloud.tencent.com/document/api/849/52009)

	* 新增入参：Uid, JdkVersion

* [CreateJobConfig](https://cloud.tencent.com/document/api/849/52004)

	* 新增入参：JdkVersion


新增数据结构：

* [FlinkJdkVersion](https://cloud.tencent.com/document/api/849/52010#FlinkJdkVersion)

修改数据结构：

* [ClusterSession](https://cloud.tencent.com/document/api/849/52010#ClusterSession)

	* 新增成员：JdkVersion

* [ClusterVersion](https://cloud.tencent.com/document/api/849/52010#ClusterVersion)

	* 新增成员：JdkSupportVersion

* [JobConfig](https://cloud.tencent.com/document/api/849/52010#JobConfig)

	* 新增成员：JdkVersion

* [JobV1](https://cloud.tencent.com/document/api/849/52010#JobV1)

	* 新增成员：ExpectJobDefaultAlarmStatus, JdkVersion

* [ResourceRefDetail](https://cloud.tencent.com/document/api/849/52010#ResourceRefDetail)

	* 新增成员：ConnectorVersion

* [SqlGatewayItem](https://cloud.tencent.com/document/api/849/52010#SqlGatewayItem)

	* 新增成员：JdkVersion




## 云数据库Redis(redis) 版本：2018-04-12

### 第 96 次发布

发布时间：2026-01-06 01:51:13

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [Account](https://cloud.tencent.com/document/api/239/20022#Account)

	* 新增成员：PasswordLastModifiedTime




## 节省计划(svp) 版本：2024-01-25

### 第 6 次发布

发布时间：2026-01-06 02:00:37

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [SavingPlanUsageDetail](https://cloud.tencent.com/document/api/1761/106543#SavingPlanUsageDetail)

	* 新增成员：SpId




## 实时音视频(trtc) 版本：2019-07-22

### 第 131 次发布

发布时间：2026-01-06 02:28:03

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeTRTCMarketQualityData](https://cloud.tencent.com/document/api/647/97441)

	* 新增入参：IsFloat


修改数据结构：

* [RowValues](https://cloud.tencent.com/document/api/647/44055#RowValues)

	* 新增成员：RowValueFloat




## TSF-应用管理&Consul(tsf) 版本：2018-03-26

### 第 137 次发布

发布时间：2026-01-06 02:30:56

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**预下线接口**：</font>

* ContinueRunFailedTaskBatch
* CreateTask
* CreateTaskFlow
* DeleteTask
* DescribeFlowLastBatchState
* DescribeTaskDetail
* DescribeTaskLastStatus
* DescribeTaskRecords
* DisableTask
* DisableTaskFlow
* EnableTask
* EnableTaskFlow
* ExecuteTask
* ExecuteTaskFlow
* ModifyTask
* RedoTask
* RedoTaskBatch
* RedoTaskExecute
* RedoTaskFlowBatch
* StopTaskBatch
* StopTaskExecute
* TerminateTaskFlowBatch



## 向量数据库(vdb) 版本：2023-06-16

### 第 16 次发布

发布时间：2026-01-06 02:35:21

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribePriceCreateInstance](https://cloud.tencent.com/document/api/1709/127211)
* [DescribePriceRenewInstance](https://cloud.tencent.com/document/api/1709/127210)
* [DescribePriceResizeInstance](https://cloud.tencent.com/document/api/1709/127209)



## 云点播(vod) 版本：2024-07-18



## 云点播(vod) 版本：2018-07-17

### 第 218 次发布

发布时间：2026-01-06 02:36:15

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [ProductImageConfig](https://cloud.tencent.com/document/api/266/31773#ProductImageConfig)

修改数据结构：

* [AigcImageSceneInfo](https://cloud.tencent.com/document/api/266/31773#AigcImageSceneInfo)

	* 新增成员：ProductImageConfig

* [SceneAigcImageOutputConfig](https://cloud.tencent.com/document/api/266/31773#SceneAigcImageOutputConfig)

	* 新增成员：AspectRatio




## Web 应用防火墙(waf) 版本：2018-01-25

### 第 138 次发布

发布时间：2026-01-05 15:03:11

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeBotIdRule](https://cloud.tencent.com/document/api/627/127099)

新增数据结构：

* [BotIdDetail](https://cloud.tencent.com/document/api/627/53609#BotIdDetail)
* [BotIdStat](https://cloud.tencent.com/document/api/627/53609#BotIdStat)



## 数据开发治理平台 WeData(wedata) 版本：2025-08-06

### 第 9 次发布

发布时间：2026-01-06 02:53:26

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [AuthorizeDataSource](https://cloud.tencent.com/document/api/1267/127235)
* [CreateCodePermissions](https://cloud.tencent.com/document/api/1267/127239)
* [CreateQualityRuleGroup](https://cloud.tencent.com/document/api/1267/127231)
* [CreateTaskFolder](https://cloud.tencent.com/document/api/1267/127269)
* [CreateTriggerTask](https://cloud.tencent.com/document/api/1267/127268)
* [CreateTriggerWorkflow](https://cloud.tencent.com/document/api/1267/127267)
* [CreateWorkflowPermissions](https://cloud.tencent.com/document/api/1267/127266)
* [DeleteCodePermissions](https://cloud.tencent.com/document/api/1267/127238)
* [DeleteDataBackfillPlanAsync](https://cloud.tencent.com/document/api/1267/127220)
* [DeleteQualityRule](https://cloud.tencent.com/document/api/1267/127230)
* [DeleteQualityRuleGroup](https://cloud.tencent.com/document/api/1267/127229)
* [DeleteTaskFolder](https://cloud.tencent.com/document/api/1267/127265)
* [DeleteTriggerTask](https://cloud.tencent.com/document/api/1267/127264)
* [DeleteTriggerWorkflow](https://cloud.tencent.com/document/api/1267/127263)
* [DeleteWorkflowPermissions](https://cloud.tencent.com/document/api/1267/127262)
* [DescribeDataSourceAuthority](https://cloud.tencent.com/document/api/1267/127234)
* [GetMyCodeMaxPermission](https://cloud.tencent.com/document/api/1267/127237)
* [GetMyWorkflowMaxPermission](https://cloud.tencent.com/document/api/1267/127261)
* [GetOpsTriggerWorkflow](https://cloud.tencent.com/document/api/1267/127219)
* [GetResourceFolder](https://cloud.tencent.com/document/api/1267/127260)
* [GetTaskFolder](https://cloud.tencent.com/document/api/1267/127259)
* [GetTriggerTask](https://cloud.tencent.com/document/api/1267/127258)
* [GetTriggerTaskCode](https://cloud.tencent.com/document/api/1267/127257)
* [GetTriggerTaskRun](https://cloud.tencent.com/document/api/1267/127218)
* [GetTriggerTaskVersion](https://cloud.tencent.com/document/api/1267/127256)
* [GetTriggerWorkflow](https://cloud.tencent.com/document/api/1267/127255)
* [GetTriggerWorkflowRun](https://cloud.tencent.com/document/api/1267/127217)
* [GetWorkflowFolder](https://cloud.tencent.com/document/api/1267/127254)
* [KillTriggerWorkflowRuns](https://cloud.tencent.com/document/api/1267/127216)
* [ListCodePermissions](https://cloud.tencent.com/document/api/1267/127236)
* [ListDownstreamTriggerTasks](https://cloud.tencent.com/document/api/1267/127253)
* [ListOpsTriggerWorkflows](https://cloud.tencent.com/document/api/1267/127215)
* [ListQualityRuleGroupExecResultsByPage](https://cloud.tencent.com/document/api/1267/127228)
* [ListQualityRuleGroupsTable](https://cloud.tencent.com/document/api/1267/127227)
* [ListQualityRuleTemplates](https://cloud.tencent.com/document/api/1267/127222)
* [ListQualityRules](https://cloud.tencent.com/document/api/1267/127226)
* [ListTaskFolders](https://cloud.tencent.com/document/api/1267/127252)
* [ListTriggerTaskVersions](https://cloud.tencent.com/document/api/1267/127251)
* [ListTriggerTasks](https://cloud.tencent.com/document/api/1267/127250)
* [ListTriggerWorkflowRuns](https://cloud.tencent.com/document/api/1267/127214)
* [ListTriggerWorkflows](https://cloud.tencent.com/document/api/1267/127249)
* [ListUpstreamTriggerTasks](https://cloud.tencent.com/document/api/1267/127248)
* [ListWorkflowPermissions](https://cloud.tencent.com/document/api/1267/127247)
* [ModifyQualityRule](https://cloud.tencent.com/document/api/1267/127225)
* [ModifyQualityRuleGroup](https://cloud.tencent.com/document/api/1267/127224)
* [RerunTriggerWorkflowRunAsync](https://cloud.tencent.com/document/api/1267/127213)
* [RevokeDataSourceAuthorization](https://cloud.tencent.com/document/api/1267/127233)
* [SubmitTriggerTask](https://cloud.tencent.com/document/api/1267/127246)
* [UpdateOpsTriggerTasksOwner](https://cloud.tencent.com/document/api/1267/127212)
* [UpdateTaskFolder](https://cloud.tencent.com/document/api/1267/127245)
* [UpdateTaskPartially](https://cloud.tencent.com/document/api/1267/127244)
* [UpdateTriggerTask](https://cloud.tencent.com/document/api/1267/127243)
* [UpdateTriggerTaskPartially](https://cloud.tencent.com/document/api/1267/127242)
* [UpdateTriggerWorkflow](https://cloud.tencent.com/document/api/1267/127241)
* [UpdateTriggerWorkflowPartially](https://cloud.tencent.com/document/api/1267/127240)

新增数据结构：

* [AuthInfo](https://cloud.tencent.com/document/api/1267/123643#AuthInfo)
* [BatchOperationOpsDto](https://cloud.tencent.com/document/api/1267/123643#BatchOperationOpsDto)
* [BizStateEnumInfoBrief](https://cloud.tencent.com/document/api/1267/123643#BizStateEnumInfoBrief)
* [CodePermissionsResultItem](https://cloud.tencent.com/document/api/1267/123643#CodePermissionsResultItem)
* [CodeStudioMaxPermission](https://cloud.tencent.com/document/api/1267/123643#CodeStudioMaxPermission)
* [CommonQualityOperateResult](https://cloud.tencent.com/document/api/1267/123643#CommonQualityOperateResult)
* [CompareQualityResult](https://cloud.tencent.com/document/api/1267/123643#CompareQualityResult)
* [CompareQualityResultItem](https://cloud.tencent.com/document/api/1267/123643#CompareQualityResultItem)
* [CreateQualityRuleGroupResultVO](https://cloud.tencent.com/document/api/1267/123643#CreateQualityRuleGroupResultVO)
* [CreateTaskFolderResult](https://cloud.tencent.com/document/api/1267/123643#CreateTaskFolderResult)
* [CreateTriggerTaskBaseAttribute](https://cloud.tencent.com/document/api/1267/123643#CreateTriggerTaskBaseAttribute)
* [CreateTriggerTaskConfiguration](https://cloud.tencent.com/document/api/1267/123643#CreateTriggerTaskConfiguration)
* [CreateTriggerTaskSchedulerConfiguration](https://cloud.tencent.com/document/api/1267/123643#CreateTriggerTaskSchedulerConfiguration)
* [CreateTriggerWorkflowResult](https://cloud.tencent.com/document/api/1267/123643#CreateTriggerWorkflowResult)
* [CreateWorkflowPermissionsResult](https://cloud.tencent.com/document/api/1267/123643#CreateWorkflowPermissionsResult)
* [DeleteQualityRuleGroupResultVO](https://cloud.tencent.com/document/api/1267/123643#DeleteQualityRuleGroupResultVO)
* [DeleteTaskFolderResult](https://cloud.tencent.com/document/api/1267/123643#DeleteTaskFolderResult)
* [DeleteTriggerWorkflowResult](https://cloud.tencent.com/document/api/1267/123643#DeleteTriggerWorkflowResult)
* [DeleteWorkflowPermission](https://cloud.tencent.com/document/api/1267/123643#DeleteWorkflowPermission)
* [DeleteWorkflowPermissionsResult](https://cloud.tencent.com/document/api/1267/123643#DeleteWorkflowPermissionsResult)
* [DependencyTriggerTaskBrief](https://cloud.tencent.com/document/api/1267/123643#DependencyTriggerTaskBrief)
* [ExecutionActionBrief](https://cloud.tencent.com/document/api/1267/123643#ExecutionActionBrief)
* [ExploreAuthorizationObject](https://cloud.tencent.com/document/api/1267/123643#ExploreAuthorizationObject)
* [ExploreAuthorizationRecycleObject](https://cloud.tencent.com/document/api/1267/123643#ExploreAuthorizationRecycleObject)
* [ExploreAuthorizeSubject](https://cloud.tencent.com/document/api/1267/123643#ExploreAuthorizeSubject)
* [ExploreFilePermissionsPage](https://cloud.tencent.com/document/api/1267/123643#ExploreFilePermissionsPage)
* [ExploreFilePrivilegeItem](https://cloud.tencent.com/document/api/1267/123643#ExploreFilePrivilegeItem)
* [ExploreFileResource](https://cloud.tencent.com/document/api/1267/123643#ExploreFileResource)
* [Filter](https://cloud.tencent.com/document/api/1267/123643#Filter)
* [ListQualityRuleGroupExecResult](https://cloud.tencent.com/document/api/1267/123643#ListQualityRuleGroupExecResult)
* [ListQualityRuleGroupExecResultPage](https://cloud.tencent.com/document/api/1267/123643#ListQualityRuleGroupExecResultPage)
* [ListTriggerTaskInfo](https://cloud.tencent.com/document/api/1267/123643#ListTriggerTaskInfo)
* [ListTriggerTaskVersions](https://cloud.tencent.com/document/api/1267/123643#ListTriggerTaskVersions)
* [ListTriggerWorkflowInfo](https://cloud.tencent.com/document/api/1267/123643#ListTriggerWorkflowInfo)
* [ModifyQualityRuleGroupResultVO](https://cloud.tencent.com/document/api/1267/123643#ModifyQualityRuleGroupResultVO)
* [OrderField](https://cloud.tencent.com/document/api/1267/123643#OrderField)
* [QualityColumnValueConfig](https://cloud.tencent.com/document/api/1267/123643#QualityColumnValueConfig)
* [QualityCompareRule](https://cloud.tencent.com/document/api/1267/123643#QualityCompareRule)
* [QualityCompareRuleItem](https://cloud.tencent.com/document/api/1267/123643#QualityCompareRuleItem)
* [QualityFieldConfig](https://cloud.tencent.com/document/api/1267/123643#QualityFieldConfig)
* [QualityProdSchedulerTask](https://cloud.tencent.com/document/api/1267/123643#QualityProdSchedulerTask)
* [QualityRule](https://cloud.tencent.com/document/api/1267/123643#QualityRule)
* [QualityRuleExecResult](https://cloud.tencent.com/document/api/1267/123643#QualityRuleExecResult)
* [QualityRuleFieldConfig](https://cloud.tencent.com/document/api/1267/123643#QualityRuleFieldConfig)
* [QualityRuleGroupConfig](https://cloud.tencent.com/document/api/1267/123643#QualityRuleGroupConfig)
* [QualityRuleGroupExecStrategy](https://cloud.tencent.com/document/api/1267/123643#QualityRuleGroupExecStrategy)
* [QualityRuleGroupResult](https://cloud.tencent.com/document/api/1267/123643#QualityRuleGroupResult)
* [QualityRuleGroupSubscribe](https://cloud.tencent.com/document/api/1267/123643#QualityRuleGroupSubscribe)
* [QualityRuleGroupTableV2](https://cloud.tencent.com/document/api/1267/123643#QualityRuleGroupTableV2)
* [QualityRuleGroupsTableVO](https://cloud.tencent.com/document/api/1267/123643#QualityRuleGroupsTableVO)
* [QualityRulePage](https://cloud.tencent.com/document/api/1267/123643#QualityRulePage)
* [QualityRuleTemplate](https://cloud.tencent.com/document/api/1267/123643#QualityRuleTemplate)
* [QualityRuleTemplatePage](https://cloud.tencent.com/document/api/1267/123643#QualityRuleTemplatePage)
* [QualitySqlExpression](https://cloud.tencent.com/document/api/1267/123643#QualitySqlExpression)
* [QualitySqlExpressionTable](https://cloud.tencent.com/document/api/1267/123643#QualitySqlExpressionTable)
* [QualitySubscribeReceiver](https://cloud.tencent.com/document/api/1267/123643#QualitySubscribeReceiver)
* [QualitySubscribeWebHook](https://cloud.tencent.com/document/api/1267/123643#QualitySubscribeWebHook)
* [QualityTableConfig](https://cloud.tencent.com/document/api/1267/123643#QualityTableConfig)
* [QualityThresholdValue](https://cloud.tencent.com/document/api/1267/123643#QualityThresholdValue)
* [ResourceFolderDetail](https://cloud.tencent.com/document/api/1267/123643#ResourceFolderDetail)
* [SchedulingParameter](https://cloud.tencent.com/document/api/1267/123643#SchedulingParameter)
* [TaskFolder](https://cloud.tencent.com/document/api/1267/123643#TaskFolder)
* [TaskFolderDetail](https://cloud.tencent.com/document/api/1267/123643#TaskFolderDetail)
* [TaskFolderPage](https://cloud.tencent.com/document/api/1267/123643#TaskFolderPage)
* [TriggerDependencyConfigPage](https://cloud.tencent.com/document/api/1267/123643#TriggerDependencyConfigPage)
* [TriggerTask](https://cloud.tencent.com/document/api/1267/123643#TriggerTask)
* [TriggerTaskBaseAttribute](https://cloud.tencent.com/document/api/1267/123643#TriggerTaskBaseAttribute)
* [TriggerTaskBrief](https://cloud.tencent.com/document/api/1267/123643#TriggerTaskBrief)
* [TriggerTaskConfiguration](https://cloud.tencent.com/document/api/1267/123643#TriggerTaskConfiguration)
* [TriggerTaskDAGBrief](https://cloud.tencent.com/document/api/1267/123643#TriggerTaskDAGBrief)
* [TriggerTaskDependDto](https://cloud.tencent.com/document/api/1267/123643#TriggerTaskDependDto)
* [TriggerTaskLinkBrief](https://cloud.tencent.com/document/api/1267/123643#TriggerTaskLinkBrief)
* [TriggerTaskRunBrief](https://cloud.tencent.com/document/api/1267/123643#TriggerTaskRunBrief)
* [TriggerTaskSchedulerConfiguration](https://cloud.tencent.com/document/api/1267/123643#TriggerTaskSchedulerConfiguration)
* [TriggerTaskVersion](https://cloud.tencent.com/document/api/1267/123643#TriggerTaskVersion)
* [TriggerTaskVersionDetail](https://cloud.tencent.com/document/api/1267/123643#TriggerTaskVersionDetail)
* [TriggerWorkflowBrief](https://cloud.tencent.com/document/api/1267/123643#TriggerWorkflowBrief)
* [TriggerWorkflowDetail](https://cloud.tencent.com/document/api/1267/123643#TriggerWorkflowDetail)
* [TriggerWorkflowInfo](https://cloud.tencent.com/document/api/1267/123643#TriggerWorkflowInfo)
* [TriggerWorkflowResult](https://cloud.tencent.com/document/api/1267/123643#TriggerWorkflowResult)
* [TriggerWorkflowRunBrief](https://cloud.tencent.com/document/api/1267/123643#TriggerWorkflowRunBrief)
* [TriggerWorkflowRunResult](https://cloud.tencent.com/document/api/1267/123643#TriggerWorkflowRunResult)
* [TriggerWorkflowTaskRunDetailBrief](https://cloud.tencent.com/document/api/1267/123643#TriggerWorkflowTaskRunDetailBrief)
* [UpdateTaskBaseAttributePart](https://cloud.tencent.com/document/api/1267/123643#UpdateTaskBaseAttributePart)
* [UpdateTaskFolderResult](https://cloud.tencent.com/document/api/1267/123643#UpdateTaskFolderResult)
* [UpdateTaskPart](https://cloud.tencent.com/document/api/1267/123643#UpdateTaskPart)
* [UpdateTriggerTaskBaseAttribute](https://cloud.tencent.com/document/api/1267/123643#UpdateTriggerTaskBaseAttribute)
* [UpdateTriggerTaskBaseAttributePart](https://cloud.tencent.com/document/api/1267/123643#UpdateTriggerTaskBaseAttributePart)
* [UpdateTriggerTaskBrief](https://cloud.tencent.com/document/api/1267/123643#UpdateTriggerTaskBrief)
* [UpdateTriggerTaskPart](https://cloud.tencent.com/document/api/1267/123643#UpdateTriggerTaskPart)
* [UpdateTriggerWorkflowPartially](https://cloud.tencent.com/document/api/1267/123643#UpdateTriggerWorkflowPartially)
* [UpdateTriggerWorkflowResult](https://cloud.tencent.com/document/api/1267/123643#UpdateTriggerWorkflowResult)
* [WorkflowFolderDetail](https://cloud.tencent.com/document/api/1267/123643#WorkflowFolderDetail)
* [WorkflowGeneralTaskParam](https://cloud.tencent.com/document/api/1267/123643#WorkflowGeneralTaskParam)
* [WorkflowMaxPermission](https://cloud.tencent.com/document/api/1267/123643#WorkflowMaxPermission)
* [WorkflowPermission](https://cloud.tencent.com/document/api/1267/123643#WorkflowPermission)
* [WorkflowPermissionPage](https://cloud.tencent.com/document/api/1267/123643#WorkflowPermissionPage)
* [WorkflowTriggerConfig](https://cloud.tencent.com/document/api/1267/123643#WorkflowTriggerConfig)



## 数据开发治理平台 WeData(wedata) 版本：2021-08-20



## 联网图像搜索(wimgs) 版本：2025-11-06

### 第 1 次发布

发布时间：2026-01-05 15:03:02

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [SearchByText](https://cloud.tencent.com/document/api/1815/127087)



