# Release v1.3.124

## 腾讯云智能体开发平台(adp) 版本：2026-05-20

### 第 3 次发布

发布时间：2026-06-26 14:45:55

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateAgent](https://cloud.tencent.com/document/api/1759/132576)

	* 新增入参：Kind

* [CreateConversation](https://cloud.tencent.com/document/api/1759/132523)

	* 新增入参：AgentId

* [DescribeConversation](https://cloud.tencent.com/document/api/1759/132519)

	* 新增出参：AgentId

* [DescribeConversationList](https://cloud.tencent.com/document/api/1759/132518)

	* 新增入参：AgentId

* [DescribeModelList](https://cloud.tencent.com/document/api/1759/132497)

	* 新增入参：Query, PageNumber, PageSize, FilterList

	* 新增出参：TotalCount

* [DescribePlugin](https://cloud.tencent.com/document/api/1759/132500)

	* 新增入参：FieldMask

* [ModifyApp](https://cloud.tencent.com/document/api/1759/132502)

	* 新增入参：SharedKbIdList

	* <font color="#dd0000">**删除入参**：</font>SharedKnowledgeIdList


新增数据结构：

* [ApiKeyAuthConfig](https://cloud.tencent.com/document/api/1759/132545#ApiKeyAuthConfig)
* [ApiPluginConfig](https://cloud.tencent.com/document/api/1759/132545#ApiPluginConfig)
* [AppPluginConfig](https://cloud.tencent.com/document/api/1759/132545#AppPluginConfig)
* [AppSharedKbInfo](https://cloud.tencent.com/document/api/1759/132545#AppSharedKbInfo)
* [AuthConfig](https://cloud.tencent.com/document/api/1759/132545#AuthConfig)
* [BasicBilling](https://cloud.tencent.com/document/api/1759/132545#BasicBilling)
* [BillingAttribute](https://cloud.tencent.com/document/api/1759/132545#BillingAttribute)
* [CamAuthConfig](https://cloud.tencent.com/document/api/1759/132545#CamAuthConfig)
* [ClawAgentConfig](https://cloud.tencent.com/document/api/1759/132545#ClawAgentConfig)
* [ClawAgentCustomConfig](https://cloud.tencent.com/document/api/1759/132545#ClawAgentCustomConfig)
* [ComplexBilling](https://cloud.tencent.com/document/api/1759/132545#ComplexBilling)
* [ComplexBillingItem](https://cloud.tencent.com/document/api/1759/132545#ComplexBillingItem)
* [DuplexBilling](https://cloud.tencent.com/document/api/1759/132545#DuplexBilling)
* [MCPPluginConfig](https://cloud.tencent.com/document/api/1759/132545#MCPPluginConfig)
* [ModelDeveloperBasic](https://cloud.tencent.com/document/api/1759/132545#ModelDeveloperBasic)
* [OAuthConfig](https://cloud.tencent.com/document/api/1759/132545#OAuthConfig)
* [PluginConfig](https://cloud.tencent.com/document/api/1759/132545#PluginConfig)
* [PluginParam](https://cloud.tencent.com/document/api/1759/132545#PluginParam)
* [SkillNotice](https://cloud.tencent.com/document/api/1759/132545#SkillNotice)

<font color="#dd0000">**删除数据结构**：</font>

* AppSharedKnowledgeInfo
* TemplatePublishInfo

修改数据结构：

* [AgentProfile](https://cloud.tencent.com/document/api/1759/132545#AgentProfile)

	* 新增成员：Role, Description, AppName, Developer, ParentAgentId

* [App](https://cloud.tencent.com/document/api/1759/132545#App)

	* 新增成员：SharedKbList

	* <font color="#dd0000">**删除成员**：</font>SharedKnowledgeList

* [AppAuxiliaryInfo](https://cloud.tencent.com/document/api/1759/132545#AppAuxiliaryInfo)

	* <font color="#dd0000">**删除成员**：</font>TemplatePublish

* [AppModeConfig](https://cloud.tencent.com/document/api/1759/132545#AppModeConfig)

	* 新增成员：ClawAgentConfig

* [CodeToolConfig](https://cloud.tencent.com/document/api/1759/132545#CodeToolConfig)

	* <font color="#dd0000">**修改成员**：</font>Code

* [Conversation](https://cloud.tencent.com/document/api/1759/132545#Conversation)

	* 新增成员：AgentId

* [Filter](https://cloud.tencent.com/document/api/1759/132545#Filter)

	* 新增成员：Operator

* [Model](https://cloud.tencent.com/document/api/1759/132545#Model)

	* 新增成员：DeveloperInfo

* [ModelDetailInfo](https://cloud.tencent.com/document/api/1759/132545#ModelDetailInfo)

	* <font color="#dd0000">**修改成员**：</font>Alias, HistoryLimit, ModelId

* [ModelParams](https://cloud.tencent.com/document/api/1759/132545#ModelParams)

	* <font color="#dd0000">**修改成员**：</font>DeepThinking, ReasoningEffort, ReplyFormat

* [Plugin](https://cloud.tencent.com/document/api/1759/132545#Plugin)

	* 新增成员：Config

* [PluginSummary](https://cloud.tencent.com/document/api/1759/132545#PluginSummary)

	* 新增成员：Config

* [RequestParam](https://cloud.tencent.com/document/api/1759/132545#RequestParam)

	* <font color="#dd0000">**修改成员**：</font>DefaultValue, Description, IsGlobalHidden, IsRequired, Name, Type

* [ResponseParam](https://cloud.tencent.com/document/api/1759/132545#ResponseParam)

	* <font color="#dd0000">**修改成员**：</font>Description, Name, Type

* [SkillSummary](https://cloud.tencent.com/document/api/1759/132545#SkillSummary)

	* 新增成员：NoticeList, PermissionIdList, SkillStatus

* [SkillVersion](https://cloud.tencent.com/document/api/1759/132545#SkillVersion)

	* 新增成员：VersionStatus, SkillMd5, SkillUrl, CreateTime, SkillMarkdownUrl, UpdateDesc

* [ToolBilling](https://cloud.tencent.com/document/api/1759/132545#ToolBilling)

	* 新增成员：BasicBilling, ComplexBilling, DuplexBilling




## 应用性能监控(apm) 版本：2021-06-22

### 第 63 次发布

发布时间：2026-06-29 01:12:10

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyApmApplicationConfig](https://cloud.tencent.com/document/api/1463/125072)

	* 新增入参：EnableHeadSampler, HeadSamplerType, HeadSamplerArg


修改数据结构：

* [ApmAppConfig](https://cloud.tencent.com/document/api/1463/64927#ApmAppConfig)

	* 新增成员：AnalysisAutoEnable, EnableHeadSampler, HeadSamplerType, HeadSamplerArg




## 云防火墙(cfw) 版本：2019-09-04

### 第 104 次发布

发布时间：2026-06-29 01:26:00

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeNDRAssetIdentificationCursorList](https://cloud.tencent.com/document/api/1132/133337)



## 日志服务(cls) 版本：2020-10-16

### 第 163 次发布

发布时间：2026-06-29 01:31:12

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [JsonExpandInfo](https://cloud.tencent.com/document/api/614/56471#JsonExpandInfo)

修改数据结构：

* [LogRechargeRuleInfo](https://cloud.tencent.com/document/api/614/56471#LogRechargeRuleInfo)

	* 新增成员：JsonExpand




## 云安全一体化平台(csip) 版本：2022-11-21

### 第 84 次发布

发布时间：2026-06-29 01:35:00

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeCLSLogIndexV3](https://cloud.tencent.com/document/api/664/133340)
* [DescribeCLSLogListV3](https://cloud.tencent.com/document/api/664/133339)

新增数据结构：

* [LogCLSFilter](https://cloud.tencent.com/document/api/664/90825#LogCLSFilter)
* [LogColumn](https://cloud.tencent.com/document/api/664/90825#LogColumn)
* [LogContextInfo](https://cloud.tencent.com/document/api/664/90825#LogContextInfo)
* [LogDynamicIndex](https://cloud.tencent.com/document/api/664/90825#LogDynamicIndex)
* [LogFullTextInfo](https://cloud.tencent.com/document/api/664/90825#LogFullTextInfo)
* [LogHighLightItem](https://cloud.tencent.com/document/api/664/90825#LogHighLightItem)
* [LogIndexRuleInfo](https://cloud.tencent.com/document/api/664/90825#LogIndexRuleInfo)
* [LogItem](https://cloud.tencent.com/document/api/664/90825#LogItem)
* [LogItems](https://cloud.tencent.com/document/api/664/90825#LogItems)
* [LogKeyValueInfo](https://cloud.tencent.com/document/api/664/90825#LogKeyValueInfo)
* [LogRuleKeyValueInfo](https://cloud.tencent.com/document/api/664/90825#LogRuleKeyValueInfo)
* [LogSearchErrors](https://cloud.tencent.com/document/api/664/90825#LogSearchErrors)
* [LogSearchInfos](https://cloud.tencent.com/document/api/664/90825#LogSearchInfos)
* [LogSearchResult](https://cloud.tencent.com/document/api/664/90825#LogSearchResult)
* [LogSearchTopics](https://cloud.tencent.com/document/api/664/90825#LogSearchTopics)
* [LogTopicIndexInfo](https://cloud.tencent.com/document/api/664/90825#LogTopicIndexInfo)
* [LogValueInfo](https://cloud.tencent.com/document/api/664/90825#LogValueInfo)



## 主机安全(cwp) 版本：2018-02-28

### 第 165 次发布

发布时间：2026-06-29 01:39:29

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [AddVulIgnoreRule](https://cloud.tencent.com/document/api/296/133341)

新增数据结构：

* [VulIgnoreRule](https://cloud.tencent.com/document/api/296/19867#VulIgnoreRule)



## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 302 次发布

发布时间：2026-06-29 01:59:52

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeCancelFlowsTask](https://cloud.tencent.com/document/api/1323/104494)

	* 新增入参：CancelType


修改数据结构：

* [CancelFailureFlow](https://cloud.tencent.com/document/api/1323/70369#CancelFailureFlow)

	* 新增成员：FlowName




## 腾讯电子签（基础版）(essbasic) 版本：2021-05-26

### 第 260 次发布

发布时间：2026-06-29 02:01:42

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeCancelFlowsTask](https://cloud.tencent.com/document/api/1420/104511)

	* 新增入参：CancelType


修改数据结构：

* [CancelFailureFlow](https://cloud.tencent.com/document/api/1420/61525#CancelFailureFlow)

	* 新增成员：FlowName




## 腾讯电子签（基础版）(essbasic) 版本：2020-12-22



## 腾讯云可观测平台(monitor) 版本：2023-06-16



## 腾讯云可观测平台(monitor) 版本：2018-07-24

### 第 160 次发布

发布时间：2026-06-29 02:24:33

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribePrometheusCreateInstanceQuota](https://cloud.tencent.com/document/api/248/133342)

新增数据结构：

* [PrometheusInstanceQuotaDetail](https://cloud.tencent.com/document/api/248/30354#PrometheusInstanceQuotaDetail)



## 媒体处理(mps) 版本：2019-06-12

### 第 210 次发布

发布时间：2026-06-29 02:26:39

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [AiPosterSuiteConfig](https://cloud.tencent.com/document/api/862/37615#AiPosterSuiteConfig)
* [CustomVariable](https://cloud.tencent.com/document/api/862/37615#CustomVariable)
* [RecipeItem](https://cloud.tencent.com/document/api/862/37615#RecipeItem)

修改数据结构：

* [ImageTaskInput](https://cloud.tencent.com/document/api/862/37615#ImageTaskInput)

	* 新增成员：AiPosterSuiteConfig




## 前端性能监控(rum) 版本：2021-06-22

### 第 53 次发布

发布时间：2026-06-29 02:37:57

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeReleaseFiles](https://cloud.tencent.com/document/api/1464/69215)

	* 新增入参：IgnoreDefaultTimeRange


修改数据结构：

* [ReleaseFile](https://cloud.tencent.com/document/api/1464/61476#ReleaseFile)

	* 新增成员：CreatedAt




## TDSQL(tdmysql) 版本：2021-11-22

### 第 10 次发布

发布时间：2026-06-29 02:53:54

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除接口**：</font>

* ResetUserPassword



## TokenHub(tokenhub) 版本：2026-03-22

### 第 8 次发布

发布时间：2026-06-29 03:02:45

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [UsageSeries](https://cloud.tencent.com/document/api/1823/132279#UsageSeries)

	* 新增成员：CacheTotalToken, SearchRequestCount, SearchCount

* [UsageStats](https://cloud.tencent.com/document/api/1823/132279#UsageStats)

	* 新增成员：CacheTotalToken, SearchRequestCount, SearchCount




## 云点播(vod) 版本：2024-07-18



## 云点播(vod) 版本：2018-07-17

### 第 267 次发布

发布时间：2026-06-29 03:11:01

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [KnowledgeAnalysisInfo](https://cloud.tencent.com/document/api/266/31773#KnowledgeAnalysisInfo)
* [KnowledgeAnalysisResult](https://cloud.tencent.com/document/api/266/31773#KnowledgeAnalysisResult)

修改数据结构：

* [KnowledgeBasesInfo](https://cloud.tencent.com/document/api/266/31773#KnowledgeBasesInfo)

	* 新增成员：KnowledgeAnalysisInfos




