# Release v1.1.12

## 日志服务(cls) 版本：2020-10-16

### 第 132 次发布

发布时间：2025-08-18 01:26:01

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeDashboardSubscribes](https://cloud.tencent.com/document/api/614/105779)

	* 新增出参：DashboardSubscribeInfos, TotalCount


新增数据结构：

* [DashboardSubscribeInfo](https://cloud.tencent.com/document/api/614/56471#DashboardSubscribeInfo)



## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 241 次发布

发布时间：2025-08-18 01:46:40

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateSeal](https://cloud.tencent.com/document/api/1323/94136)

	* 新增入参：SealDescription


修改数据结构：

* [OccupiedSeal](https://cloud.tencent.com/document/api/1323/70369#OccupiedSeal)

	* 新增成员：RealWidth, RealHeight, SubSealType, SubSealName, SealDescription

* [OrganizationAuthUrl](https://cloud.tencent.com/document/api/1323/70369#OrganizationAuthUrl)

	* 新增成员：OrganizationName




## 腾讯电子签（基础版）(essbasic) 版本：2021-05-26

### 第 230 次发布

发布时间：2025-08-18 01:48:00

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateSealByImage](https://cloud.tencent.com/document/api/1420/73067)

	* 新增入参：SealDescription


修改数据结构：

* [OccupiedSeal](https://cloud.tencent.com/document/api/1420/61525#OccupiedSeal)

	* 新增成员：RealWidth, RealHeight, SealDescription




## 腾讯电子签（基础版）(essbasic) 版本：2020-12-22



## 腾讯云智能体开发平台(lke) 版本：2023-11-30

### 第 59 次发布

发布时间：2025-08-18 02:04:32

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateApp](https://cloud.tencent.com/document/api/1759/105074)

	* 新增入参：AgentType

* [CreateSharedKnowledge](https://cloud.tencent.com/document/api/1759/119409)

	* 新增入参：KnowledgeType

* [CreateVar](https://cloud.tencent.com/document/api/1759/116688)

	* 新增入参：VarModuleType

* [DeleteVar](https://cloud.tencent.com/document/api/1759/119875)

	* 新增入参：VarModuleType

* [DescribeCallStatsGraph](https://cloud.tencent.com/document/api/1759/111069)

	* 新增入参：SpaceId

* [DescribeConcurrencyUsage](https://cloud.tencent.com/document/api/1759/111068)

	* 新增入参：SpaceId

* [DescribeConcurrencyUsageGraph](https://cloud.tencent.com/document/api/1759/111067)

	* 新增入参：SpaceId

* [DescribeDoc](https://cloud.tencent.com/document/api/1759/105071)

	* 新增出参：SplitRule, UpdatePeriodInfo

* [DescribeKnowledgeUsagePieGraph](https://cloud.tencent.com/document/api/1759/111065)

	* 新增入参：SpaceId

* [DescribeSearchStatsGraph](https://cloud.tencent.com/document/api/1759/111064)

	* 新增入参：SpaceId

* [DescribeTokenUsage](https://cloud.tencent.com/document/api/1759/111063)

	* 新增入参：SpaceId

* [DescribeTokenUsageGraph](https://cloud.tencent.com/document/api/1759/111062)

	* 新增入参：SubScenes

* [GetVarList](https://cloud.tencent.com/document/api/1759/116687)

	* 新增入参：VarModuleType

* [ListApp](https://cloud.tencent.com/document/api/1759/105066)

	* 新增入参：AgentType, AppStatus

* [ListAppKnowledgeDetail](https://cloud.tencent.com/document/api/1759/113596)

	* 新增入参：SpaceId

* [ListSharedKnowledge](https://cloud.tencent.com/document/api/1759/119405)

	* 新增入参：KnowledgeTypes

* [ListUsageCallDetail](https://cloud.tencent.com/document/api/1759/113595)

	* 新增入参：SpaceId

* [ModifyDoc](https://cloud.tencent.com/document/api/1759/105058)

	* 新增入参：ModifyTypes, UpdatePeriodInfo, SplitRule

* [SaveDoc](https://cloud.tencent.com/document/api/1759/105054)

	* 新增入参：SplitRule, UpdatePeriodInfo

* [UpdateVar](https://cloud.tencent.com/document/api/1759/119874)

	* 新增入参：VarModuleType


新增数据结构：

* [AgentPluginQuery](https://cloud.tencent.com/document/api/1759/105104#AgentPluginQuery)
* [KnowledgeAdvancedConfig](https://cloud.tencent.com/document/api/1759/105104#KnowledgeAdvancedConfig)
* [KnowledgeModelConfig](https://cloud.tencent.com/document/api/1759/105104#KnowledgeModelConfig)
* [KnowledgeQaAgent](https://cloud.tencent.com/document/api/1759/105104#KnowledgeQaAgent)
* [KnowledgeQaWorkflowInfo](https://cloud.tencent.com/document/api/1759/105104#KnowledgeQaWorkflowInfo)
* [ModelParams](https://cloud.tencent.com/document/api/1759/105104#ModelParams)
* [OptionCardIndex](https://cloud.tencent.com/document/api/1759/105104#OptionCardIndex)
* [UpdatePeriodInfo](https://cloud.tencent.com/document/api/1759/105104#UpdatePeriodInfo)

修改数据结构：

* [Agent](https://cloud.tencent.com/document/api/1759/105104#Agent)

	* 新增成员：AgentMode

* [AgentDebugInfo](https://cloud.tencent.com/document/api/1759/105104#AgentDebugInfo)

	* 新增成员：ModelName

* [AgentMCPServerInfo](https://cloud.tencent.com/document/api/1759/105104#AgentMCPServerInfo)

	* 新增成员：Query

* [AgentModelInfo](https://cloud.tencent.com/document/api/1759/105104#AgentModelInfo)

	* 新增成员：ModelParams

* [AgentPluginInfo](https://cloud.tencent.com/document/api/1759/105104#AgentPluginInfo)

	* 新增成员：EnableRoleAuth, Query, McpType

* [AgentReference](https://cloud.tencent.com/document/api/1759/105104#AgentReference)

	* 新增成员：KnowledgeName, KnowledgeBizId

* [AgentToolInfo](https://cloud.tencent.com/document/api/1759/105104#AgentToolInfo)

	* 新增成员：Query, FinanceStatus

* [AppInfo](https://cloud.tencent.com/document/api/1759/105104#AppInfo)

	* 新增成员：PermissionIds

* [AppModel](https://cloud.tencent.com/document/api/1759/105104#AppModel)

	* 新增成员：ModelParams

* [KnowledgeBaseInfo](https://cloud.tencent.com/document/api/1759/105104#KnowledgeBaseInfo)

	* 新增成员：KnowledgeType, OwnerStaffId, DocTotal, ProcessingFlags

* [KnowledgeDetailInfo](https://cloud.tencent.com/document/api/1759/105104#KnowledgeDetailInfo)

	* 新增成员：PermissionIds

* [KnowledgeQaConfig](https://cloud.tencent.com/document/api/1759/105104#KnowledgeQaConfig)

	* 新增成员：LongMemoryOpen, LongMemoryDay, Agent, KnowledgeModelConfig, KnowledgeAdvancedConfig

* [KnowledgeUpdateInfo](https://cloud.tencent.com/document/api/1759/105104#KnowledgeUpdateInfo)

	* 新增成员：OwnerStaffId

* [ModelInfo](https://cloud.tencent.com/document/api/1759/105104#ModelInfo)

	* 新增成员：ModelTags, ModelParams, ProviderName, ProviderAliasName, ProviderType

* [ModelParameter](https://cloud.tencent.com/document/api/1759/105104#ModelParameter)

	* 新增成员：Name

* [MsgRecordReference](https://cloud.tencent.com/document/api/1759/105104#MsgRecordReference)

	* 新增成员：Index

	* <font color="#dd0000">**修改成员**：</font>Id, Url, Type, Name, DocId, KnowledgeName, KnowledgeBizId, DocBizId, QaBizId

* [Procedure](https://cloud.tencent.com/document/api/1759/105104#Procedure)

	* 新增成员：InputCount, OutputCount

* [SearchStrategy](https://cloud.tencent.com/document/api/1759/105104#SearchStrategy)

	* 新增成员：EmbeddingModel, RerankModelSwitch, RerankModel

* [ShareKnowledgeBase](https://cloud.tencent.com/document/api/1759/105104#ShareKnowledgeBase)

	* 新增成员：KnowledgeModelConfig, SearchStrategy, Search, ReplyFlexibility, KnowledgeName

* [TaskFLowVar](https://cloud.tencent.com/document/api/1759/105104#TaskFLowVar)

	* 新增成员：VarModuleType

* [WorkFlowSummary](https://cloud.tencent.com/document/api/1759/105104#WorkFlowSummary)

	* 新增成员：PendingMessages, OptionCardIndex

* [WorkflowRunDetail](https://cloud.tencent.com/document/api/1759/105104#WorkflowRunDetail)

	* 新增成员：Output




## 文字识别(ocr) 版本：2018-11-19

### 第 210 次发布

发布时间：2025-08-18 02:20:14

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [VatElectronicInfo](https://cloud.tencent.com/document/api/866/33527#VatElectronicInfo)

	* 新增成员：InvoicePageIndex




## 消息队列 RabbitMQ Serverless 版(trabbit) 版本：2023-04-18

### 第 5 次发布

发布时间：2025-08-18 03:09:26

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyRabbitMQServerlessInstance](https://cloud.tencent.com/document/api/1495/116152)

	* 新增入参：SendReceiveRatio




## TSF-应用管理&Consul(tsf) 版本：2018-03-26

### 第 131 次发布

发布时间：2025-08-18 03:16:01

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateGroup](https://cloud.tencent.com/document/api/649/36074)

	* 新增入参：K8sNamespaceName

* [DeployContainerApplication](https://cloud.tencent.com/document/api/649/120669)

	* 新增入参：K8sNamespaceName, StaticIpEnabled, PodManagementPolicyType, Partition


修改数据结构：

* [GatewayPlugin](https://cloud.tencent.com/document/api/649/36099#GatewayPlugin)

	* 新增成员：DeleteDisabled, DeleteDisabledReason

* [LaneInfo](https://cloud.tencent.com/document/api/649/36099#LaneInfo)

	* 新增成员：DeleteDisabled, DeleteDisabledReason

* [SimpleGroup](https://cloud.tencent.com/document/api/649/36099#SimpleGroup)

	* 新增成员：K8sNamespaceName




## 云点播(vod) 版本：2024-07-18



## 云点播(vod) 版本：2018-07-17

### 第 208 次发布

发布时间：2025-08-18 03:22:04

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateMPSTemplate](https://cloud.tencent.com/document/api/266/122580)
* [DeleteMPSTemplate](https://cloud.tencent.com/document/api/266/122579)
* [DescribeMPSTemplates](https://cloud.tencent.com/document/api/266/122578)
* [ModifyMPSTemplate](https://cloud.tencent.com/document/api/266/122577)

新增数据结构：

* [MPSTemplate](https://cloud.tencent.com/document/api/266/31773#MPSTemplate)



