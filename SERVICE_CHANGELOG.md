# Release v1.3.185

## 应用性能监控(apm) 版本：2021-06-22

### 第 69 次发布

发布时间：2026-09-23 01:10:50

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [OpenApmPaidVersion](https://cloud.tencent.com/document/api/1463/138606)



## 云防火墙(cfw) 版本：2019-09-04

### 第 116 次发布

发布时间：2026-09-23 01:30:25

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyStorageSetting](https://cloud.tencent.com/document/api/1132/86751)

	* 新增出参：Status




## 负载均衡(clb) 版本：2018-03-17

### 第 167 次发布

发布时间：2026-09-23 01:35:36

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [ServiceProviderHealthCheckConfigItemInput](https://cloud.tencent.com/document/api/214/30694#ServiceProviderHealthCheckConfigItemInput)

	* 新增成员：HealthCheckPath, HealthCheckMethod

* [ServiceProviderHealthCheckConfigItemOutput](https://cloud.tencent.com/document/api/214/30694#ServiceProviderHealthCheckConfigItemOutput)

	* 新增成员：HealthCheckPath, HealthCheckMethod




## 日志服务(cls) 版本：2020-10-16

### 第 181 次发布

发布时间：2026-09-23 01:38:58

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateResourceGraph](https://cloud.tencent.com/document/api/614/137648)

	* 新增出参：ResourceGraphId

* [DescribeResourceGraphEntities](https://cloud.tencent.com/document/api/614/137653)

	* 新增出参：EntityInfos, HasMore

* [DescribeResourceGraphEntityDependency](https://cloud.tencent.com/document/api/614/137652)

	* 新增出参：Topology

* [DescribeResourceGraphEntityDetail](https://cloud.tencent.com/document/api/614/137651)

	* 新增出参：EntityInfo

* [DescribeResourceGraphs](https://cloud.tencent.com/document/api/614/137644)

	* 新增入参：Filters, Offset, Limit

	* 新增出参：ResourceGraphInfos, TotalCount


新增数据结构：

* [DependencyTopology](https://cloud.tencent.com/document/api/614/56471#DependencyTopology)
* [EntityAttribute](https://cloud.tencent.com/document/api/614/56471#EntityAttribute)
* [EntityInfo](https://cloud.tencent.com/document/api/614/56471#EntityInfo)
* [RelatedTopicItem](https://cloud.tencent.com/document/api/614/56471#RelatedTopicItem)
* [ResourceGraphInfo](https://cloud.tencent.com/document/api/614/56471#ResourceGraphInfo)
* [TopologyEdge](https://cloud.tencent.com/document/api/614/56471#TopologyEdge)
* [TopologyNode](https://cloud.tencent.com/document/api/614/56471#TopologyNode)



## TDSQL-C MySQL 版(cynosdb) 版本：2019-01-07

### 第 196 次发布

发布时间：2026-09-23 02:02:27

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [CynosdbInstanceDetail](https://cloud.tencent.com/document/api/1003/48097#CynosdbInstanceDetail)

	* 新增成员：RealZone, SlaveZones, StorageVersion




## 高性能应用服务(hai) 版本：2023-08-12

### 第 30 次发布

发布时间：2026-09-23 02:38:46

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [GetServicePodLogs](https://cloud.tencent.com/document/api/1721/138547)

	* 新增入参：ServiceId, PodName, TailLines

	* 新增出参：LogLines




## 腾讯混元大模型(hunyuan) 版本：2023-09-01

### 第 53 次发布

发布时间：2026-09-23 02:39:19

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**预下线接口**：</font>

* ChatCompletions
* ChatTranslations
* CreateGlossary
* CreateGlossaryEntry
* CreateThread
* DeleteGlossary
* DeleteGlossaryEntry
* FilesDeletions
* FilesList
* FilesUploads
* GetEmbedding
* GetThread
* GetThreadMessage
* GetThreadMessageList
* GetTokenCount
* GroupChatCompletions
* ImageQuestion
* ListGlossary
* ListGlossaryEntry
* RunThread
* UpdateGlossaryEntry



## 云直播CSS(live) 版本：2018-08-01

### 第 188 次发布

发布时间：2026-09-23 02:55:56

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateAuditRule](https://cloud.tencent.com/document/api/267/138616)
* [CreateAuditTemplate](https://cloud.tencent.com/document/api/267/138615)
* [CreateLiveSmartEraseTemplate](https://cloud.tencent.com/document/api/267/138608)
* [DeleteAuditRule](https://cloud.tencent.com/document/api/267/138614)
* [DeleteAuditTemplate](https://cloud.tencent.com/document/api/267/138613)
* [DescribeAuditRules](https://cloud.tencent.com/document/api/267/138612)
* [DescribeAuditTemplate](https://cloud.tencent.com/document/api/267/138611)
* [DescribeAuditTemplates](https://cloud.tencent.com/document/api/267/138610)
* [ModifyAuditTemplate](https://cloud.tencent.com/document/api/267/138609)

新增数据结构：

* [AuditTemplate](https://cloud.tencent.com/document/api/267/20474#AuditTemplate)
* [CMSBizInfo](https://cloud.tencent.com/document/api/267/20474#CMSBizInfo)
* [CMSSceneDetail](https://cloud.tencent.com/document/api/267/20474#CMSSceneDetail)



## 媒体处理(mps) 版本：2019-06-12

### 第 251 次发布

发布时间：2026-09-23 03:06:36

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [AigcAudioExtraParam](https://cloud.tencent.com/document/api/862/37615#AigcAudioExtraParam)

	* 新增成员：OutputAudioFormat




## 云开发 CloudBase(tcb) 版本：2018-06-08

### 第 169 次发布

发布时间：2026-09-23 03:33:12

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeTaskResult](https://cloud.tencent.com/document/api/876/138619)
* [ResetPGAccountPassword](https://cloud.tencent.com/document/api/876/138618)
* [UpgradePGInstanceToDedicated](https://cloud.tencent.com/document/api/876/138617)

新增数据结构：

* [ObjectKV](https://cloud.tencent.com/document/api/876/34822#ObjectKV)



## 腾讯云数据库 AI 服务(tdai) 版本：2025-07-17

### 第 19 次发布

发布时间：2026-09-23 03:41:07

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateAgentInstance](https://cloud.tencent.com/document/api/1813/123274)

	* 新增入参：DeploymentFree, EnableMemory




## 边缘安全加速平台(teo) 版本：2022-09-01

### 第 165 次发布

发布时间：2026-09-23 03:44:06

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [AccountProtectionSettings](https://cloud.tencent.com/document/api/1552/80721#AccountProtectionSettings)
* [SecurityHeadersToOrigin](https://cloud.tencent.com/document/api/1552/80721#SecurityHeadersToOrigin)
* [UserRiskProfile](https://cloud.tencent.com/document/api/1552/80721#UserRiskProfile)

修改数据结构：

* [ClientAttestationRule](https://cloud.tencent.com/document/api/1552/80721#ClientAttestationRule)

	* 新增成员：AccountProtectionSettings

* [SecurityPolicy](https://cloud.tencent.com/document/api/1552/80721#SecurityPolicy)

	* 新增成员：SecurityHeadersToOrigin




## 边缘安全加速平台(teo) 版本：2022-01-06



## WorkBuddy Enterprise(workbuddyenterprise) 版本：2026-07-09

### 第 2 次发布

发布时间：2026-09-23 04:24:50

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [BindExternalAgent](https://cloud.tencent.com/document/api/1831/138635)
* [CreateAgentSession](https://cloud.tencent.com/document/api/1831/138630)
* [DescribeAgentSession](https://cloud.tencent.com/document/api/1831/138629)
* [DescribeAgentSessionList](https://cloud.tencent.com/document/api/1831/138628)
* [DescribeBuiltinModelList](https://cloud.tencent.com/document/api/1831/138624)
* [DescribeConnectorList](https://cloud.tencent.com/document/api/1831/138621)
* [DescribeExpertList](https://cloud.tencent.com/document/api/1831/138639)
* [DescribeExternalAgent](https://cloud.tencent.com/document/api/1831/138634)
* [DescribeExternalAgentList](https://cloud.tencent.com/document/api/1831/138633)
* [DescribeMessageEventList](https://cloud.tencent.com/document/api/1831/138627)
* [DescribeSkillList](https://cloud.tencent.com/document/api/1831/138637)
* [MigrateAgentSession](https://cloud.tencent.com/document/api/1831/138626)
* [ModifyAgentVersion](https://cloud.tencent.com/document/api/1831/138622)
* [UnbindExternalAgent](https://cloud.tencent.com/document/api/1831/138632)

修改接口：

* [CreateAgentVersion](https://cloud.tencent.com/document/api/1831/138571)

	* <font color="#dd0000">**删除入参**：</font>IsTest

* [CreateAgentVersionFromSource](https://cloud.tencent.com/document/api/1831/138570)

	* <font color="#dd0000">**删除入参**：</font>IsTest


新增数据结构：

* [A2ASkillItem](https://cloud.tencent.com/document/api/1831/138580#A2ASkillItem)
* [BuiltinModel](https://cloud.tencent.com/document/api/1831/138580#BuiltinModel)
* [ChatEndpoint](https://cloud.tencent.com/document/api/1831/138580#ChatEndpoint)
* [ConnectorInfo](https://cloud.tencent.com/document/api/1831/138580#ConnectorInfo)
* [ExpertCounts](https://cloud.tencent.com/document/api/1831/138580#ExpertCounts)
* [ExpertItem](https://cloud.tencent.com/document/api/1831/138580#ExpertItem)
* [ExternalAgentInfo](https://cloud.tencent.com/document/api/1831/138580#ExternalAgentInfo)
* [MessageEvent](https://cloud.tencent.com/document/api/1831/138580#MessageEvent)
* [MessageEventMessage](https://cloud.tencent.com/document/api/1831/138580#MessageEventMessage)
* [MessageEventToolCall](https://cloud.tencent.com/document/api/1831/138580#MessageEventToolCall)
* [SessionItem](https://cloud.tencent.com/document/api/1831/138580#SessionItem)
* [SkillCounts](https://cloud.tencent.com/document/api/1831/138580#SkillCounts)
* [SkillItem](https://cloud.tencent.com/document/api/1831/138580#SkillItem)
* [TokenUsage](https://cloud.tencent.com/document/api/1831/138580#TokenUsage)

### 第 1 次发布

发布时间：2026-09-22 11:14:11

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateAgent](https://cloud.tencent.com/document/api/1831/138579)
* [CreateAgentVersion](https://cloud.tencent.com/document/api/1831/138571)
* [CreateAgentVersionFromSource](https://cloud.tencent.com/document/api/1831/138570)
* [DeleteAgent](https://cloud.tencent.com/document/api/1831/138578)
* [DescribeAgent](https://cloud.tencent.com/document/api/1831/138577)
* [DescribeAgentList](https://cloud.tencent.com/document/api/1831/138576)
* [DescribeAgentVersion](https://cloud.tencent.com/document/api/1831/138569)
* [DescribeAgentVersionList](https://cloud.tencent.com/document/api/1831/138592)
* [ModifyAgent](https://cloud.tencent.com/document/api/1831/138575)
* [ModifyAgentA2AConfig](https://cloud.tencent.com/document/api/1831/138574)
* [ModifyAgentRouting](https://cloud.tencent.com/document/api/1831/138573)

新增数据结构：

* [A2AConfig](https://cloud.tencent.com/document/api/1831/138580#A2AConfig)
* [A2ASkillInput](https://cloud.tencent.com/document/api/1831/138580#A2ASkillInput)
* [AgentItem](https://cloud.tencent.com/document/api/1831/138580#AgentItem)
* [AgentVersionItem](https://cloud.tencent.com/document/api/1831/138580#AgentVersionItem)
* [ConnectorRefInput](https://cloud.tencent.com/document/api/1831/138580#ConnectorRefInput)
* [Filter](https://cloud.tencent.com/document/api/1831/138580#Filter)
* [RoutingItem](https://cloud.tencent.com/document/api/1831/138580#RoutingItem)



