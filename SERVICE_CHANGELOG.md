# Release v1.3.111

## 腾讯云智能体开发平台(adp) 版本：2026-05-20

### 第 2 次发布

发布时间：2026-06-04 19:43:26

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateAgent](https://cloud.tencent.com/document/api/1759/132576)

<font color="#dd0000">**删除接口**：</font>

* DescribeMsgRecord

修改接口：

* [CreateVariable](https://cloud.tencent.com/document/api/1759/132530)

	* 新增入参：Variable

* [DescribeApp](https://cloud.tencent.com/document/api/1759/132504)

	* 新增入参：FieldMask

	* 新增出参：App

* [DescribeLatestRelease](https://cloud.tencent.com/document/api/1759/132536)

	* 新增出参：ReleaseSummary

* [DescribeReleaseSummary](https://cloud.tencent.com/document/api/1759/132534)

	* 新增出参：ReleaseSummary

* [DescribeVariable](https://cloud.tencent.com/document/api/1759/132527)

	* 新增出参：Variable

* [ModifyApp](https://cloud.tencent.com/document/api/1759/132502)

	* 新增入参：Config, ShareConfig, UpdateMask

* [ModifyVariable](https://cloud.tencent.com/document/api/1759/132525)

	* 新增入参：Variable


新增数据结构：

* [AICallConfig](https://cloud.tencent.com/document/api/1759/132545#AICallConfig)
* [AIOptimizeModel](https://cloud.tencent.com/document/api/1759/132545#AIOptimizeModel)
* [AgentCollaborationConfig](https://cloud.tencent.com/document/api/1759/132545#AgentCollaborationConfig)
* [App](https://cloud.tencent.com/document/api/1759/132545#App)
* [AppAdvancedConf](https://cloud.tencent.com/document/api/1759/132545#AppAdvancedConf)
* [AppAppeal](https://cloud.tencent.com/document/api/1759/132545#AppAppeal)
* [AppAuxiliaryInfo](https://cloud.tencent.com/document/api/1759/132545#AppAuxiliaryInfo)
* [AppConfig](https://cloud.tencent.com/document/api/1759/132545#AppConfig)
* [AppExperienceConfig](https://cloud.tencent.com/document/api/1759/132545#AppExperienceConfig)
* [AppGreetingConfig](https://cloud.tencent.com/document/api/1759/132545#AppGreetingConfig)
* [AppMemoryConfig](https://cloud.tencent.com/document/api/1759/132545#AppMemoryConfig)
* [AppMetadata](https://cloud.tencent.com/document/api/1759/132545#AppMetadata)
* [AppModeConfig](https://cloud.tencent.com/document/api/1759/132545#AppModeConfig)
* [AppModelConfig](https://cloud.tencent.com/document/api/1759/132545#AppModelConfig)
* [AppOperation](https://cloud.tencent.com/document/api/1759/132545#AppOperation)
* [AppSecretInfo](https://cloud.tencent.com/document/api/1759/132545#AppSecretInfo)
* [AppShareAccessControl](https://cloud.tencent.com/document/api/1759/132545#AppShareAccessControl)
* [AppShareURLInfo](https://cloud.tencent.com/document/api/1759/132545#AppShareURLInfo)
* [AppShareWhitelistItem](https://cloud.tencent.com/document/api/1759/132545#AppShareWhitelistItem)
* [AppSharedKnowledgeInfo](https://cloud.tencent.com/document/api/1759/132545#AppSharedKnowledgeInfo)
* [AppStatusInfo](https://cloud.tencent.com/document/api/1759/132545#AppStatusInfo)
* [AppSubStatusInfo](https://cloud.tencent.com/document/api/1759/132545#AppSubStatusInfo)
* [AppWebSearchConfig](https://cloud.tencent.com/document/api/1759/132545#AppWebSearchConfig)
* [AppWorkflowConfig](https://cloud.tencent.com/document/api/1759/132545#AppWorkflowConfig)
* [AppealingStatus](https://cloud.tencent.com/document/api/1759/132545#AppealingStatus)
* [BackgroundImage](https://cloud.tencent.com/document/api/1759/132545#BackgroundImage)
* [ConversationExperience](https://cloud.tencent.com/document/api/1759/132545#ConversationExperience)
* [DigitalHumanConfig](https://cloud.tencent.com/document/api/1759/132545#DigitalHumanConfig)
* [FileParseModel](https://cloud.tencent.com/document/api/1759/132545#FileParseModel)
* [GenerateModel](https://cloud.tencent.com/document/api/1759/132545#GenerateModel)
* [InputBoxConfig](https://cloud.tencent.com/document/api/1759/132545#InputBoxConfig)
* [IntentAchievementInfo](https://cloud.tencent.com/document/api/1759/132545#IntentAchievementInfo)
* [ModelBasic](https://cloud.tencent.com/document/api/1759/132545#ModelBasic)
* [ModelDetailInfo](https://cloud.tencent.com/document/api/1759/132545#ModelDetailInfo)
* [ModelLimit](https://cloud.tencent.com/document/api/1759/132545#ModelLimit)
* [ModelProviderBasic](https://cloud.tencent.com/document/api/1759/132545#ModelProviderBasic)
* [ModelStatus](https://cloud.tencent.com/document/api/1759/132545#ModelStatus)
* [MultiAgentConfig](https://cloud.tencent.com/document/api/1759/132545#MultiAgentConfig)
* [MultiModalQAModel](https://cloud.tencent.com/document/api/1759/132545#MultiModalQAModel)
* [MultiModalUnderstandingModel](https://cloud.tencent.com/document/api/1759/132545#MultiModalUnderstandingModel)
* [PromptRewriteModel](https://cloud.tencent.com/document/api/1759/132545#PromptRewriteModel)
* [ReleaseSummary](https://cloud.tencent.com/document/api/1759/132545#ReleaseSummary)
* [RoleConfig](https://cloud.tencent.com/document/api/1759/132545#RoleConfig)
* [SearchResourceStatusInfo](https://cloud.tencent.com/document/api/1759/132545#SearchResourceStatusInfo)
* [SingleWorkflowConfig](https://cloud.tencent.com/document/api/1759/132545#SingleWorkflowConfig)
* [SkillAnalysisInfo](https://cloud.tencent.com/document/api/1759/132545#SkillAnalysisInfo)
* [SpecialStatusInfo](https://cloud.tencent.com/document/api/1759/132545#SpecialStatusInfo)
* [SupportedFileType](https://cloud.tencent.com/document/api/1759/132545#SupportedFileType)
* [TemplatePublishInfo](https://cloud.tencent.com/document/api/1759/132545#TemplatePublishInfo)
* [ThinkModel](https://cloud.tencent.com/document/api/1759/132545#ThinkModel)
* [VoiceConfig](https://cloud.tencent.com/document/api/1759/132545#VoiceConfig)

<font color="#dd0000">**删除数据结构**：</font>

* MsgRecord
* MsgRecordResult
* MsgRecordSource

修改数据结构：

* [AgentToolBasicConfig](https://cloud.tencent.com/document/api/1759/132545#AgentToolBasicConfig)

	* 新增成员：Description

* [AppSummary](https://cloud.tencent.com/document/api/1759/132545#AppSummary)

	* 新增成员：OperationInfo, Status, SubStatus

* [Model](https://cloud.tencent.com/document/api/1759/132545#Model)

	* 新增成员：LimitInfo, ModelBasic, ProviderInfo, StatusInfo

* [SkillVersion](https://cloud.tencent.com/document/api/1759/132545#SkillVersion)

	* 新增成员：AnalysisInfo


### 第 1 次发布

发布时间：2026-06-04 11:56:25

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CopyApp](https://cloud.tencent.com/document/api/1759/132507)
* [CreateApp](https://cloud.tencent.com/document/api/1759/132506)
* [CreateConversation](https://cloud.tencent.com/document/api/1759/132523)
* [CreateRelease](https://cloud.tencent.com/document/api/1759/132538)
* [CreateSpace](https://cloud.tencent.com/document/api/1759/132513)
* [CreateVariable](https://cloud.tencent.com/document/api/1759/132530)
* [CreateWebSocketToken](https://cloud.tencent.com/document/api/1759/132522)
* [CreateWorkspaceCredential](https://cloud.tencent.com/document/api/1759/132521)
* [DeleteApp](https://cloud.tencent.com/document/api/1759/132505)
* [DeleteConversation](https://cloud.tencent.com/document/api/1759/132520)
* [DeleteSpace](https://cloud.tencent.com/document/api/1759/132512)
* [DeleteVariable](https://cloud.tencent.com/document/api/1759/132529)
* [DescribeAgentDetail](https://cloud.tencent.com/document/api/1759/132544)
* [DescribeAgentReleasePreviewList](https://cloud.tencent.com/document/api/1759/132537)
* [DescribeApp](https://cloud.tencent.com/document/api/1759/132504)
* [DescribeAppSummaryList](https://cloud.tencent.com/document/api/1759/132503)
* [DescribeConversation](https://cloud.tencent.com/document/api/1759/132519)
* [DescribeConversationList](https://cloud.tencent.com/document/api/1759/132518)
* [DescribeConversationMessageList](https://cloud.tencent.com/document/api/1759/132517)
* [DescribeLatestRelease](https://cloud.tencent.com/document/api/1759/132536)
* [DescribeModelList](https://cloud.tencent.com/document/api/1759/132497)
* [DescribeMsgRecord](https://cloud.tencent.com/document/api/#/#)
* [DescribePlugin](https://cloud.tencent.com/document/api/1759/132500)
* [DescribePluginSummaryList](https://cloud.tencent.com/document/api/1759/132499)
* [DescribeReleaseList](https://cloud.tencent.com/document/api/1759/132535)
* [DescribeReleaseSummary](https://cloud.tencent.com/document/api/1759/132534)
* [DescribeSkillCategoryList](https://cloud.tencent.com/document/api/1759/132541)
* [DescribeSkillSummaryList](https://cloud.tencent.com/document/api/1759/132540)
* [DescribeSpaceList](https://cloud.tencent.com/document/api/1759/132510)
* [DescribeSystemVariableList](https://cloud.tencent.com/document/api/1759/132528)
* [DescribeVariable](https://cloud.tencent.com/document/api/1759/132527)
* [DescribeVariableList](https://cloud.tencent.com/document/api/1759/132526)
* [ModifyAgent](https://cloud.tencent.com/document/api/1759/132543)
* [ModifyApp](https://cloud.tencent.com/document/api/1759/132502)
* [ModifyConversation](https://cloud.tencent.com/document/api/1759/132516)
* [ModifySpace](https://cloud.tencent.com/document/api/1759/132509)
* [ModifyVariable](https://cloud.tencent.com/document/api/1759/132525)
* [ResetConversation](https://cloud.tencent.com/document/api/1759/132515)
* [RetryRelease](https://cloud.tencent.com/document/api/1759/132533)
* [RollbackRelease](https://cloud.tencent.com/document/api/1759/132532)

新增数据结构：

* [AgentAdvancedConfig](https://cloud.tencent.com/document/api/1759/132545#AgentAdvancedConfig)
* [AgentDetail](https://cloud.tencent.com/document/api/1759/132545#AgentDetail)
* [AgentInput](https://cloud.tencent.com/document/api/1759/132545#AgentInput)
* [AgentModelConfig](https://cloud.tencent.com/document/api/1759/132545#AgentModelConfig)
* [AgentPlugin](https://cloud.tencent.com/document/api/1759/132545#AgentPlugin)
* [AgentPluginConfig](https://cloud.tencent.com/document/api/1759/132545#AgentPluginConfig)
* [AgentPluginParameter](https://cloud.tencent.com/document/api/1759/132545#AgentPluginParameter)
* [AgentProfile](https://cloud.tencent.com/document/api/1759/132545#AgentProfile)
* [AgentRelease](https://cloud.tencent.com/document/api/1759/132545#AgentRelease)
* [AgentReleasePreview](https://cloud.tencent.com/document/api/1759/132545#AgentReleasePreview)
* [AgentSkill](https://cloud.tencent.com/document/api/1759/132545#AgentSkill)
* [AgentSkillConfig](https://cloud.tencent.com/document/api/1759/132545#AgentSkillConfig)
* [AgentSpec](https://cloud.tencent.com/document/api/1759/132545#AgentSpec)
* [AgentSystemVariable](https://cloud.tencent.com/document/api/1759/132545#AgentSystemVariable)
* [AgentTool](https://cloud.tencent.com/document/api/1759/132545#AgentTool)
* [AgentToolBasicConfig](https://cloud.tencent.com/document/api/1759/132545#AgentToolBasicConfig)
* [AgentToolConfig](https://cloud.tencent.com/document/api/1759/132545#AgentToolConfig)
* [AgentToolInputParameter](https://cloud.tencent.com/document/api/1759/132545#AgentToolInputParameter)
* [AgentToolOutputParameter](https://cloud.tencent.com/document/api/1759/132545#AgentToolOutputParameter)
* [AgentUserInputValue](https://cloud.tencent.com/document/api/1759/132545#AgentUserInputValue)
* [ApiToolConfig](https://cloud.tencent.com/document/api/1759/132545#ApiToolConfig)
* [AppSummary](https://cloud.tencent.com/document/api/1759/132545#AppSummary)
* [AppToolConfig](https://cloud.tencent.com/document/api/1759/132545#AppToolConfig)
* [CodeToolConfig](https://cloud.tencent.com/document/api/1759/132545#CodeToolConfig)
* [Conversation](https://cloud.tencent.com/document/api/1759/132545#Conversation)
* [ConversationAgentTask](https://cloud.tencent.com/document/api/1759/132545#ConversationAgentTask)
* [ConversationContent](https://cloud.tencent.com/document/api/1759/132545#ConversationContent)
* [ConversationMessage](https://cloud.tencent.com/document/api/1759/132545#ConversationMessage)
* [ConversationQuoteInfo](https://cloud.tencent.com/document/api/1759/132545#ConversationQuoteInfo)
* [ConversationReference](https://cloud.tencent.com/document/api/1759/132545#ConversationReference)
* [ConversationWorkspace](https://cloud.tencent.com/document/api/1759/132545#ConversationWorkspace)
* [FieldMask](https://cloud.tencent.com/document/api/1759/132545#FieldMask)
* [Filter](https://cloud.tencent.com/document/api/1759/132545#Filter)
* [MCPToolConfig](https://cloud.tencent.com/document/api/1759/132545#MCPToolConfig)
* [Model](https://cloud.tencent.com/document/api/1759/132545#Model)
* [ModelBadge](https://cloud.tencent.com/document/api/1759/132545#ModelBadge)
* [ModelParameter](https://cloud.tencent.com/document/api/1759/132545#ModelParameter)
* [ModelParams](https://cloud.tencent.com/document/api/1759/132545#ModelParams)
* [ModelProperty](https://cloud.tencent.com/document/api/1759/132545#ModelProperty)
* [MsgRecord](https://cloud.tencent.com/document/api/1759/132545#MsgRecord)
* [MsgRecordResult](https://cloud.tencent.com/document/api/1759/132545#MsgRecordResult)
* [MsgRecordSource](https://cloud.tencent.com/document/api/1759/132545#MsgRecordSource)
* [Plugin](https://cloud.tencent.com/document/api/1759/132545#Plugin)
* [PluginOperation](https://cloud.tencent.com/document/api/1759/132545#PluginOperation)
* [PluginProfile](https://cloud.tencent.com/document/api/1759/132545#PluginProfile)
* [PluginStatistics](https://cloud.tencent.com/document/api/1759/132545#PluginStatistics)
* [PluginSummary](https://cloud.tencent.com/document/api/1759/132545#PluginSummary)
* [PluginUserState](https://cloud.tencent.com/document/api/1759/132545#PluginUserState)
* [ReleaseRecord](https://cloud.tencent.com/document/api/1759/132545#ReleaseRecord)
* [RequestParam](https://cloud.tencent.com/document/api/1759/132545#RequestParam)
* [ResponseParam](https://cloud.tencent.com/document/api/1759/132545#ResponseParam)
* [SkillCategory](https://cloud.tencent.com/document/api/1759/132545#SkillCategory)
* [SkillClassification](https://cloud.tencent.com/document/api/1759/132545#SkillClassification)
* [SkillProfile](https://cloud.tencent.com/document/api/1759/132545#SkillProfile)
* [SkillShare](https://cloud.tencent.com/document/api/1759/132545#SkillShare)
* [SkillSummary](https://cloud.tencent.com/document/api/1759/132545#SkillSummary)
* [SkillVersion](https://cloud.tencent.com/document/api/1759/132545#SkillVersion)
* [Space](https://cloud.tencent.com/document/api/1759/132545#Space)
* [SystemVariable](https://cloud.tencent.com/document/api/1759/132545#SystemVariable)
* [Tool](https://cloud.tencent.com/document/api/1759/132545#Tool)
* [ToolBilling](https://cloud.tencent.com/document/api/1759/132545#ToolBilling)
* [ToolConfig](https://cloud.tencent.com/document/api/1759/132545#ToolConfig)
* [ToolExample](https://cloud.tencent.com/document/api/1759/132545#ToolExample)
* [Variable](https://cloud.tencent.com/document/api/1759/132545#Variable)



## 云防火墙(cfw) 版本：2019-09-04

### 第 102 次发布

发布时间：2026-06-05 01:24:34

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CloseClusterNatFwSwitch](https://cloud.tencent.com/document/api/1132/132582)
* [DescribeClusterNatCcnFwSwitchList](https://cloud.tencent.com/document/api/1132/132581)
* [DescribeNatCcnFwSwitch](https://cloud.tencent.com/document/api/1132/132580)
* [DescribeNatFwClusterRegionStatus](https://cloud.tencent.com/document/api/1132/132584)
* [ModifyClusterFwBypass](https://cloud.tencent.com/document/api/1132/132579)
* [ModifyClusterNatFwSwitch](https://cloud.tencent.com/document/api/1132/132578)
* [OpenClusterNatFwSwitch](https://cloud.tencent.com/document/api/1132/132577)

新增数据结构：

* [FilterDataObject](https://cloud.tencent.com/document/api/1132/49071#FilterDataObject)
* [NatCcnSwitchConfig](https://cloud.tencent.com/document/api/1132/49071#NatCcnSwitchConfig)
* [NatClusterRegionStatusQuery](https://cloud.tencent.com/document/api/1132/49071#NatClusterRegionStatusQuery)
* [NatFwClusterRegionStatus](https://cloud.tencent.com/document/api/1132/49071#NatFwClusterRegionStatus)
* [NatFwSwitchDetailS](https://cloud.tencent.com/document/api/1132/49071#NatFwSwitchDetailS)



## 主机安全(cwp) 版本：2018-02-28

### 第 162 次发布

发布时间：2026-06-05 01:36:40

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [Machine](https://cloud.tencent.com/document/api/296/19867#Machine)

	* 新增成员：AppId, CSIPProtectType




## Web 应用防火墙(waf) 版本：2018-01-25

### 第 157 次发布

发布时间：2026-06-05 03:13:27

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [QpsData](https://cloud.tencent.com/document/api/627/53609#QpsData)

	* 新增成员：ElasticPrepaidRatio




