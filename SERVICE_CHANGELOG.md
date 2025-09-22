# Release v1.1.31

## 应用性能监控(apm) 版本：2021-06-22

### 第 46 次发布

发布时间：2025-09-22 01:09:49

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeApmServiceMetric](https://cloud.tencent.com/document/api/1463/123528)

新增数据结构：

* [ApmServiceMetric](https://cloud.tencent.com/document/api/1463/64927#ApmServiceMetric)
* [ServiceDetail](https://cloud.tencent.com/document/api/1463/64927#ServiceDetail)



## 费用中心(billing) 版本：2018-07-09

### 第 76 次发布

发布时间：2025-09-23 01:11:20

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateBudget](https://cloud.tencent.com/document/api/555/123661)
* [DeleteBudget](https://cloud.tencent.com/document/api/555/123660)
* [DescribeBudget](https://cloud.tencent.com/document/api/555/123659)
* [DescribeBudgetOperationLog](https://cloud.tencent.com/document/api/555/123658)
* [DescribeBudgetRemindRecordList](https://cloud.tencent.com/document/api/555/123657)
* [ModifyBudget](https://cloud.tencent.com/document/api/555/123656)

新增数据结构：

* [BudgetConditionsForm](https://cloud.tencent.com/document/api/555/19183#BudgetConditionsForm)
* [BudgetExtend](https://cloud.tencent.com/document/api/555/19183#BudgetExtend)
* [BudgetInfoApiResponse](https://cloud.tencent.com/document/api/555/19183#BudgetInfoApiResponse)
* [BudgetInfoDiffEntity](https://cloud.tencent.com/document/api/555/19183#BudgetInfoDiffEntity)
* [BudgetOperationLogEntity](https://cloud.tencent.com/document/api/555/19183#BudgetOperationLogEntity)
* [BudgetPlan](https://cloud.tencent.com/document/api/555/19183#BudgetPlan)
* [BudgetRemindRecordList](https://cloud.tencent.com/document/api/555/19183#BudgetRemindRecordList)
* [BudgetRemindRecords](https://cloud.tencent.com/document/api/555/19183#BudgetRemindRecords)
* [BudgetSendInfoDto](https://cloud.tencent.com/document/api/555/19183#BudgetSendInfoDto)
* [BudgetWarn](https://cloud.tencent.com/document/api/555/19183#BudgetWarn)
* [DataForBudgetInfoPage](https://cloud.tencent.com/document/api/555/19183#DataForBudgetInfoPage)
* [DataForBudgetOperationLogPage](https://cloud.tencent.com/document/api/555/19183#DataForBudgetOperationLogPage)
* [OrderDto](https://cloud.tencent.com/document/api/555/19183#OrderDto)
* [TagsForm](https://cloud.tencent.com/document/api/555/19183#TagsForm)
* [WaveThresholdForm](https://cloud.tencent.com/document/api/555/19183#WaveThresholdForm)



## 云联络中心(ccc) 版本：2020-02-10

### 第 108 次发布

发布时间：2025-09-23 01:14:29

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ControlAIConversation](https://cloud.tencent.com/document/api/679/120723)

	* 新增入参：InvokeLLM


新增数据结构：

* [InvokeLLM](https://cloud.tencent.com/document/api/679/47715#InvokeLLM)

### 第 107 次发布

发布时间：2025-09-22 01:14:27

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateOwnNumberApply](https://cloud.tencent.com/document/api/679/111404)

	* 新增入参：MobileNddPrefix, LocalNumberTrimAC

* [ModifyOwnNumberApply](https://cloud.tencent.com/document/api/679/111403)

	* 新增入参：MobileNddPrefix, LocalNumberTrimAC


修改数据结构：

* [OwnNumberApplyDetailItem](https://cloud.tencent.com/document/api/679/47715#OwnNumberApplyDetailItem)

	* 新增成员：CarrierPhoneNumber




## 云服务器(cvm) 版本：2017-03-12

### 第 156 次发布

发布时间：2025-09-19 11:37:45

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除接口**：</font>

* DescribeReservedInstancesConfigInfos
* DescribeReservedInstancesOfferings
* InquirePricePurchaseReservedInstancesOffering
* PurchaseReservedInstancesOffering

修改接口：

* [CreateLaunchTemplate](https://cloud.tencent.com/document/api/213/66327)

	* 新增入参：EnableJumboFrame

* [CreateLaunchTemplateVersion](https://cloud.tencent.com/document/api/213/66326)

	* 新增入参：EnableJumboFrame

* [RunInstances](https://cloud.tencent.com/document/api/213/15730)

	* 新增入参：MinCount


<font color="#dd0000">**删除数据结构**：</font>

* ReservedInstanceConfigInfoItem
* ReservedInstanceFamilyItem
* ReservedInstancePrice
* ReservedInstancePriceItem
* ReservedInstanceTypeItem
* ReservedInstancesOffering



## 云游戏(gs) 版本：2019-11-18

### 第 56 次发布

发布时间：2025-09-22 01:57:51

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [BackUpAndroidInstance](https://cloud.tencent.com/document/api/1162/123532)
* [DeleteAndroidInstanceBackups](https://cloud.tencent.com/document/api/1162/123531)
* [DescribeAndroidInstanceBackups](https://cloud.tencent.com/document/api/1162/123530)
* [RestoreAndroidInstance](https://cloud.tencent.com/document/api/1162/123529)



## 腾讯云智能体开发平台(lke) 版本：2023-11-30

### 第 65 次发布

发布时间：2025-09-23 01:53:05

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [WorkflowRunNodeInfo](https://cloud.tencent.com/document/api/1759/105104#WorkflowRunNodeInfo)

	* 新增成员：FailCode


### 第 64 次发布

发布时间：2025-09-22 02:24:36

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [ModelInfo](https://cloud.tencent.com/document/api/1759/105104#ModelInfo)

	* 新增成员：IsCloseModelParams, IsDeepThinking




## 多网聚合加速(mna) 版本：2021-01-19

### 第 30 次发布

发布时间：2025-09-23 01:58:19

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [GetFlowStatisticByGroup](https://cloud.tencent.com/document/api/1385/106607)

	* 新增入参：MpApplicationId




## 云数据库 SQL Server(sqlserver) 版本：2018-03-28

### 第 75 次发布

发布时间：2025-09-23 02:25:59

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeDBInstancesAttribute](https://cloud.tencent.com/document/api/238/90299)

	* 新增出参：SlowLogThreshold




## 数据开发治理平台 WeData(wedata) 版本：2025-08-06

### 第 1 次发布

发布时间：2025-09-23 00:17:35

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateCodeFile](https://cloud.tencent.com/document/api/1267/123612)
* [CreateCodeFolder](https://cloud.tencent.com/document/api/1267/123611)
* [CreateDataBackfillPlan](https://cloud.tencent.com/document/api/1267/123592)
* [CreateOpsAlarmRule](https://cloud.tencent.com/document/api/1267/123591)
* [CreateResourceFile](https://cloud.tencent.com/document/api/1267/123642)
* [CreateResourceFolder](https://cloud.tencent.com/document/api/1267/123641)
* [CreateSQLFolder](https://cloud.tencent.com/document/api/1267/123610)
* [CreateSQLScript](https://cloud.tencent.com/document/api/1267/123609)
* [CreateTask](https://cloud.tencent.com/document/api/1267/123640)
* [CreateWorkflow](https://cloud.tencent.com/document/api/1267/123639)
* [CreateWorkflowFolder](https://cloud.tencent.com/document/api/1267/123638)
* [DeleteCodeFile](https://cloud.tencent.com/document/api/1267/123608)
* [DeleteCodeFolder](https://cloud.tencent.com/document/api/1267/123607)
* [DeleteOpsAlarmRule](https://cloud.tencent.com/document/api/1267/123590)
* [DeleteResourceFile](https://cloud.tencent.com/document/api/1267/123637)
* [DeleteResourceFolder](https://cloud.tencent.com/document/api/1267/123636)
* [DeleteSQLFolder](https://cloud.tencent.com/document/api/1267/123606)
* [DeleteSQLScript](https://cloud.tencent.com/document/api/1267/123605)
* [DeleteTask](https://cloud.tencent.com/document/api/1267/123635)
* [DeleteWorkflow](https://cloud.tencent.com/document/api/1267/123634)
* [DeleteWorkflowFolder](https://cloud.tencent.com/document/api/1267/123633)
* [GetAlarmMessage](https://cloud.tencent.com/document/api/1267/123589)
* [GetCodeFile](https://cloud.tencent.com/document/api/1267/123604)
* [GetOpsAlarmRule](https://cloud.tencent.com/document/api/1267/123588)
* [GetOpsAsyncJob](https://cloud.tencent.com/document/api/1267/123587)
* [GetOpsTask](https://cloud.tencent.com/document/api/1267/123586)
* [GetOpsTaskCode](https://cloud.tencent.com/document/api/1267/123585)
* [GetOpsWorkflow](https://cloud.tencent.com/document/api/1267/123584)
* [GetResourceFile](https://cloud.tencent.com/document/api/1267/123632)
* [GetSQLScript](https://cloud.tencent.com/document/api/1267/123603)
* [GetTask](https://cloud.tencent.com/document/api/1267/123631)
* [GetTaskCode](https://cloud.tencent.com/document/api/1267/123630)
* [GetTaskInstance](https://cloud.tencent.com/document/api/1267/123583)
* [GetTaskInstanceLog](https://cloud.tencent.com/document/api/1267/123582)
* [GetTaskVersion](https://cloud.tencent.com/document/api/1267/123629)
* [GetWorkflow](https://cloud.tencent.com/document/api/1267/123628)
* [KillTaskInstancesAsync](https://cloud.tencent.com/document/api/1267/123581)
* [ListAlarmMessages](https://cloud.tencent.com/document/api/1267/123580)
* [ListCodeFolderContents](https://cloud.tencent.com/document/api/1267/123602)
* [ListDataBackfillInstances](https://cloud.tencent.com/document/api/1267/123579)
* [ListDownstreamOpsTasks](https://cloud.tencent.com/document/api/1267/123578)
* [ListDownstreamTaskInstances](https://cloud.tencent.com/document/api/1267/123577)
* [ListDownstreamTasks](https://cloud.tencent.com/document/api/1267/123627)
* [ListOpsAlarmRules](https://cloud.tencent.com/document/api/1267/123576)
* [ListOpsTasks](https://cloud.tencent.com/document/api/1267/123575)
* [ListOpsWorkflows](https://cloud.tencent.com/document/api/1267/123574)
* [ListResourceFiles](https://cloud.tencent.com/document/api/1267/123626)
* [ListResourceFolders](https://cloud.tencent.com/document/api/1267/123625)
* [ListSQLFolderContents](https://cloud.tencent.com/document/api/1267/123601)
* [ListSQLScriptRuns](https://cloud.tencent.com/document/api/1267/123600)
* [ListTaskInstanceExecutions](https://cloud.tencent.com/document/api/1267/123573)
* [ListTaskInstances](https://cloud.tencent.com/document/api/1267/123572)
* [ListTaskVersions](https://cloud.tencent.com/document/api/1267/123624)
* [ListTasks](https://cloud.tencent.com/document/api/1267/123623)
* [ListUpstreamOpsTasks](https://cloud.tencent.com/document/api/1267/123571)
* [ListUpstreamTaskInstances](https://cloud.tencent.com/document/api/1267/123570)
* [ListUpstreamTasks](https://cloud.tencent.com/document/api/1267/123622)
* [ListWorkflowFolders](https://cloud.tencent.com/document/api/1267/123621)
* [ListWorkflows](https://cloud.tencent.com/document/api/1267/123620)
* [PauseOpsTasksAsync](https://cloud.tencent.com/document/api/1267/123569)
* [RerunTaskInstancesAsync](https://cloud.tencent.com/document/api/1267/123568)
* [RunSQLScript](https://cloud.tencent.com/document/api/1267/123599)
* [SetSuccessTaskInstancesAsync](https://cloud.tencent.com/document/api/1267/123567)
* [StopOpsTasksAsync](https://cloud.tencent.com/document/api/1267/123566)
* [StopSQLScriptRun](https://cloud.tencent.com/document/api/1267/123598)
* [SubmitTask](https://cloud.tencent.com/document/api/1267/123619)
* [UpdateCodeFile](https://cloud.tencent.com/document/api/1267/123597)
* [UpdateCodeFolder](https://cloud.tencent.com/document/api/1267/123596)
* [UpdateOpsAlarmRule](https://cloud.tencent.com/document/api/1267/123565)
* [UpdateOpsTasksOwner](https://cloud.tencent.com/document/api/1267/123564)
* [UpdateResourceFile](https://cloud.tencent.com/document/api/1267/123618)
* [UpdateResourceFolder](https://cloud.tencent.com/document/api/1267/123617)
* [UpdateSQLFolder](https://cloud.tencent.com/document/api/1267/123595)
* [UpdateSQLScript](https://cloud.tencent.com/document/api/1267/123594)
* [UpdateTask](https://cloud.tencent.com/document/api/1267/123616)
* [UpdateWorkflow](https://cloud.tencent.com/document/api/1267/123615)
* [UpdateWorkflowFolder](https://cloud.tencent.com/document/api/1267/123614)

新增数据结构：

* [AlarmGroup](https://cloud.tencent.com/document/api/1267/123643#AlarmGroup)
* [AlarmMessage](https://cloud.tencent.com/document/api/1267/123643#AlarmMessage)
* [AlarmQuietInterval](https://cloud.tencent.com/document/api/1267/123643#AlarmQuietInterval)
* [AlarmRuleData](https://cloud.tencent.com/document/api/1267/123643#AlarmRuleData)
* [AlarmRuleDetail](https://cloud.tencent.com/document/api/1267/123643#AlarmRuleDetail)
* [AlarmWayWebHook](https://cloud.tencent.com/document/api/1267/123643#AlarmWayWebHook)
* [BackfillInstance](https://cloud.tencent.com/document/api/1267/123643#BackfillInstance)
* [BackfillInstanceCollection](https://cloud.tencent.com/document/api/1267/123643#BackfillInstanceCollection)
* [ChildDependencyConfigPage](https://cloud.tencent.com/document/api/1267/123643#ChildDependencyConfigPage)
* [CodeFile](https://cloud.tencent.com/document/api/1267/123643#CodeFile)
* [CodeFileConfig](https://cloud.tencent.com/document/api/1267/123643#CodeFileConfig)
* [CodeFolderNode](https://cloud.tencent.com/document/api/1267/123643#CodeFolderNode)
* [CodeStudioFileActionResult](https://cloud.tencent.com/document/api/1267/123643#CodeStudioFileActionResult)
* [CodeStudioFolderActionResult](https://cloud.tencent.com/document/api/1267/123643#CodeStudioFolderActionResult)
* [CodeStudioFolderResult](https://cloud.tencent.com/document/api/1267/123643#CodeStudioFolderResult)
* [CreateAlarmRuleData](https://cloud.tencent.com/document/api/1267/123643#CreateAlarmRuleData)
* [CreateDataReplenishmentPlan](https://cloud.tencent.com/document/api/1267/123643#CreateDataReplenishmentPlan)
* [CreateFolderResult](https://cloud.tencent.com/document/api/1267/123643#CreateFolderResult)
* [CreateResourceFileResult](https://cloud.tencent.com/document/api/1267/123643#CreateResourceFileResult)
* [CreateTaskBaseAttribute](https://cloud.tencent.com/document/api/1267/123643#CreateTaskBaseAttribute)
* [CreateTaskConfiguration](https://cloud.tencent.com/document/api/1267/123643#CreateTaskConfiguration)
* [CreateTaskResult](https://cloud.tencent.com/document/api/1267/123643#CreateTaskResult)
* [CreateTaskSchedulerConfiguration](https://cloud.tencent.com/document/api/1267/123643#CreateTaskSchedulerConfiguration)
* [CreateWorkflowResult](https://cloud.tencent.com/document/api/1267/123643#CreateWorkflowResult)
* [DataBackfillRange](https://cloud.tencent.com/document/api/1267/123643#DataBackfillRange)
* [DeleteAlarmRuleResult](https://cloud.tencent.com/document/api/1267/123643#DeleteAlarmRuleResult)
* [DeleteFolderResult](https://cloud.tencent.com/document/api/1267/123643#DeleteFolderResult)
* [DeleteResourceFileResult](https://cloud.tencent.com/document/api/1267/123643#DeleteResourceFileResult)
* [DeleteTaskResult](https://cloud.tencent.com/document/api/1267/123643#DeleteTaskResult)
* [DeleteWorkflowResult](https://cloud.tencent.com/document/api/1267/123643#DeleteWorkflowResult)
* [DependencyConfigPage](https://cloud.tencent.com/document/api/1267/123643#DependencyConfigPage)
* [DependencyStrategyTask](https://cloud.tencent.com/document/api/1267/123643#DependencyStrategyTask)
* [DependencyTaskBrief](https://cloud.tencent.com/document/api/1267/123643#DependencyTaskBrief)
* [EventListener](https://cloud.tencent.com/document/api/1267/123643#EventListener)
* [InTaskParameter](https://cloud.tencent.com/document/api/1267/123643#InTaskParameter)
* [InstanceExecution](https://cloud.tencent.com/document/api/1267/123643#InstanceExecution)
* [InstanceExecutionPhase](https://cloud.tencent.com/document/api/1267/123643#InstanceExecutionPhase)
* [InstanceLog](https://cloud.tencent.com/document/api/1267/123643#InstanceLog)
* [JobDto](https://cloud.tencent.com/document/api/1267/123643#JobDto)
* [JobExecutionDto](https://cloud.tencent.com/document/api/1267/123643#JobExecutionDto)
* [KVMap](https://cloud.tencent.com/document/api/1267/123643#KVMap)
* [KVPair](https://cloud.tencent.com/document/api/1267/123643#KVPair)
* [ListAlarmMessages](https://cloud.tencent.com/document/api/1267/123643#ListAlarmMessages)
* [ListAlarmRulesResult](https://cloud.tencent.com/document/api/1267/123643#ListAlarmRulesResult)
* [ListOpsTasksPage](https://cloud.tencent.com/document/api/1267/123643#ListOpsTasksPage)
* [ListTaskInfo](https://cloud.tencent.com/document/api/1267/123643#ListTaskInfo)
* [ListTaskVersions](https://cloud.tencent.com/document/api/1267/123643#ListTaskVersions)
* [ListWorkflowInfo](https://cloud.tencent.com/document/api/1267/123643#ListWorkflowInfo)
* [ModifyAlarmRuleResult](https://cloud.tencent.com/document/api/1267/123643#ModifyAlarmRuleResult)
* [NotebookSessionInfo](https://cloud.tencent.com/document/api/1267/123643#NotebookSessionInfo)
* [NotificationFatigue](https://cloud.tencent.com/document/api/1267/123643#NotificationFatigue)
* [OpsAsyncJobDetail](https://cloud.tencent.com/document/api/1267/123643#OpsAsyncJobDetail)
* [OpsAsyncResponse](https://cloud.tencent.com/document/api/1267/123643#OpsAsyncResponse)
* [OpsTaskDepend](https://cloud.tencent.com/document/api/1267/123643#OpsTaskDepend)
* [OpsWorkflow](https://cloud.tencent.com/document/api/1267/123643#OpsWorkflow)
* [OpsWorkflowDetail](https://cloud.tencent.com/document/api/1267/123643#OpsWorkflowDetail)
* [OpsWorkflows](https://cloud.tencent.com/document/api/1267/123643#OpsWorkflows)
* [OutTaskParameter](https://cloud.tencent.com/document/api/1267/123643#OutTaskParameter)
* [ParamInfo](https://cloud.tencent.com/document/api/1267/123643#ParamInfo)
* [ParentDependencyConfigPage](https://cloud.tencent.com/document/api/1267/123643#ParentDependencyConfigPage)
* [ProjectInstanceStatisticsAlarmInfo](https://cloud.tencent.com/document/api/1267/123643#ProjectInstanceStatisticsAlarmInfo)
* [ReconciliationStrategyInfo](https://cloud.tencent.com/document/api/1267/123643#ReconciliationStrategyInfo)
* [ResourceFile](https://cloud.tencent.com/document/api/1267/123643#ResourceFile)
* [ResourceFileItem](https://cloud.tencent.com/document/api/1267/123643#ResourceFileItem)
* [ResourceFilePage](https://cloud.tencent.com/document/api/1267/123643#ResourceFilePage)
* [ResourceFolder](https://cloud.tencent.com/document/api/1267/123643#ResourceFolder)
* [ResourceFolderPage](https://cloud.tencent.com/document/api/1267/123643#ResourceFolderPage)
* [SQLContentActionResult](https://cloud.tencent.com/document/api/1267/123643#SQLContentActionResult)
* [SQLFolderNode](https://cloud.tencent.com/document/api/1267/123643#SQLFolderNode)
* [SQLScript](https://cloud.tencent.com/document/api/1267/123643#SQLScript)
* [SQLScriptConfig](https://cloud.tencent.com/document/api/1267/123643#SQLScriptConfig)
* [SQLStopResult](https://cloud.tencent.com/document/api/1267/123643#SQLStopResult)
* [SqlCreateResult](https://cloud.tencent.com/document/api/1267/123643#SqlCreateResult)
* [SubmitTaskResult](https://cloud.tencent.com/document/api/1267/123643#SubmitTaskResult)
* [Task](https://cloud.tencent.com/document/api/1267/123643#Task)
* [TaskBaseAttribute](https://cloud.tencent.com/document/api/1267/123643#TaskBaseAttribute)
* [TaskCode](https://cloud.tencent.com/document/api/1267/123643#TaskCode)
* [TaskCodeResult](https://cloud.tencent.com/document/api/1267/123643#TaskCodeResult)
* [TaskConfiguration](https://cloud.tencent.com/document/api/1267/123643#TaskConfiguration)
* [TaskDataRegistry](https://cloud.tencent.com/document/api/1267/123643#TaskDataRegistry)
* [TaskDependDto](https://cloud.tencent.com/document/api/1267/123643#TaskDependDto)
* [TaskExtParameter](https://cloud.tencent.com/document/api/1267/123643#TaskExtParameter)
* [TaskInstance](https://cloud.tencent.com/document/api/1267/123643#TaskInstance)
* [TaskInstanceDetail](https://cloud.tencent.com/document/api/1267/123643#TaskInstanceDetail)
* [TaskInstanceExecutions](https://cloud.tencent.com/document/api/1267/123643#TaskInstanceExecutions)
* [TaskInstancePage](https://cloud.tencent.com/document/api/1267/123643#TaskInstancePage)
* [TaskOpsInfo](https://cloud.tencent.com/document/api/1267/123643#TaskOpsInfo)
* [TaskSchedulerConfiguration](https://cloud.tencent.com/document/api/1267/123643#TaskSchedulerConfiguration)
* [TaskSchedulingParameter](https://cloud.tencent.com/document/api/1267/123643#TaskSchedulingParameter)
* [TaskVersion](https://cloud.tencent.com/document/api/1267/123643#TaskVersion)
* [TaskVersionDetail](https://cloud.tencent.com/document/api/1267/123643#TaskVersionDetail)
* [TimeOutStrategyInfo](https://cloud.tencent.com/document/api/1267/123643#TimeOutStrategyInfo)
* [UpdateFolderResult](https://cloud.tencent.com/document/api/1267/123643#UpdateFolderResult)
* [UpdateResourceFileResult](https://cloud.tencent.com/document/api/1267/123643#UpdateResourceFileResult)
* [UpdateTaskBaseAttribute](https://cloud.tencent.com/document/api/1267/123643#UpdateTaskBaseAttribute)
* [UpdateTaskBrief](https://cloud.tencent.com/document/api/1267/123643#UpdateTaskBrief)
* [UpdateTaskResult](https://cloud.tencent.com/document/api/1267/123643#UpdateTaskResult)
* [UpdateTasksOwner](https://cloud.tencent.com/document/api/1267/123643#UpdateTasksOwner)
* [UpdateWorkflowResult](https://cloud.tencent.com/document/api/1267/123643#UpdateWorkflowResult)
* [WorkflowDetail](https://cloud.tencent.com/document/api/1267/123643#WorkflowDetail)
* [WorkflowFolder](https://cloud.tencent.com/document/api/1267/123643#WorkflowFolder)
* [WorkflowFolderPage](https://cloud.tencent.com/document/api/1267/123643#WorkflowFolderPage)
* [WorkflowInfo](https://cloud.tencent.com/document/api/1267/123643#WorkflowInfo)
* [WorkflowSchedulerConfiguration](https://cloud.tencent.com/document/api/1267/123643#WorkflowSchedulerConfiguration)
* [WorkflowSchedulerConfigurationInfo](https://cloud.tencent.com/document/api/1267/123643#WorkflowSchedulerConfigurationInfo)



## 数据开发治理平台 WeData(wedata) 版本：2021-08-20



