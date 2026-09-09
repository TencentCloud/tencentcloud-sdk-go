# Release v1.3.177

## 腾讯云智能体开发平台(adp) 版本：2026-05-20

### 第 21 次发布

发布时间：2026-09-10 01:08:41

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [AgentPluginCredentialConfig](https://cloud.tencent.com/document/api/1759/132545#AgentPluginCredentialConfig)
* [AgentPluginCredentialParam](https://cloud.tencent.com/document/api/1759/132545#AgentPluginCredentialParam)
* [SkillCorpShareConfig](https://cloud.tencent.com/document/api/1759/132545#SkillCorpShareConfig)

修改数据结构：

* [AgentPluginConfig](https://cloud.tencent.com/document/api/1759/132545#AgentPluginConfig)

	* 新增成员：CredentialConfig

* [SkillProfile](https://cloud.tencent.com/document/api/1759/132545#SkillProfile)

	* 新增成员：SpaceId

* [SkillShare](https://cloud.tencent.com/document/api/1759/132545#SkillShare)

	* 新增成员：CorpShareConfig

* [SkillVersion](https://cloud.tencent.com/document/api/1759/132545#SkillVersion)

	* 新增成员：Updater

* [Variable](https://cloud.tencent.com/document/api/1759/132545#Variable)

	* 新增成员：IsBuiltin, EnableSandbox


### 第 20 次发布

发布时间：2026-09-09 10:33:57

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateChannel](https://cloud.tencent.com/document/api/1759/137760)
* [DeleteChannel](https://cloud.tencent.com/document/api/1759/137759)
* [DescribeChannel](https://cloud.tencent.com/document/api/1759/137758)
* [DescribeChannelList](https://cloud.tencent.com/document/api/1759/137757)
* [ModifyChannel](https://cloud.tencent.com/document/api/1759/137756)

新增数据结构：

* [CallbackConfig](https://cloud.tencent.com/document/api/1759/132545#CallbackConfig)
* [Channel](https://cloud.tencent.com/document/api/1759/132545#Channel)
* [ChannelSpec](https://cloud.tencent.com/document/api/1759/132545#ChannelSpec)
* [DingTalkChannelConfig](https://cloud.tencent.com/document/api/1759/132545#DingTalkChannelConfig)
* [LarkChannelConfig](https://cloud.tencent.com/document/api/1759/132545#LarkChannelConfig)
* [LineChannelConfig](https://cloud.tencent.com/document/api/1759/132545#LineChannelConfig)
* [TelegramChannelConfig](https://cloud.tencent.com/document/api/1759/132545#TelegramChannelConfig)
* [UserAgentReference](https://cloud.tencent.com/document/api/1759/132545#UserAgentReference)
* [WechatChannelConfig](https://cloud.tencent.com/document/api/1759/132545#WechatChannelConfig)
* [WechatClawBotChannelConfig](https://cloud.tencent.com/document/api/1759/132545#WechatClawBotChannelConfig)
* [WechatCustomerServiceChannelConfig](https://cloud.tencent.com/document/api/1759/132545#WechatCustomerServiceChannelConfig)
* [WecomAppChannelConfig](https://cloud.tencent.com/document/api/1759/132545#WecomAppChannelConfig)
* [WecomRobotCallbackAccess](https://cloud.tencent.com/document/api/1759/132545#WecomRobotCallbackAccess)
* [WecomRobotChannelConfig](https://cloud.tencent.com/document/api/1759/132545#WecomRobotChannelConfig)
* [WecomRobotWebsocketAccess](https://cloud.tencent.com/document/api/1759/132545#WecomRobotWebsocketAccess)



## Cloud Studio（云端 IDE）(cloudstudio) 版本：2023-05-08

### 第 8 次发布

发布时间：2026-09-10 01:34:45

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateWorkspace](https://cloud.tencent.com/document/api/1039/94096)

	* <font color="#dd0000">**删除入参**：</font>Region




## 云服务器(cvm) 版本：2017-03-12

### 第 172 次发布

发布时间：2026-09-10 01:40:57

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [Instance](https://cloud.tencent.com/document/api/213/15753#Instance)

	* 新增成员：EnableJumboFrame




## 腾讯云数据分析智能体(dataagent) 版本：2025-05-13

### 第 24 次发布

发布时间：2026-09-10 01:49:26

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [QueryUserSessionDetail](https://cloud.tencent.com/document/api/1800/137839)

<font color="#dd0000">**删除接口**：</font>

* GetSessionDetails

新增数据结构：

* [RecordList](https://cloud.tencent.com/document/api/1800/125016#RecordList)

<font color="#dd0000">**删除数据结构**：</font>

* Record
* StepExpand
* StepInfo
* Task



## 云数据库独享集群(dbdc) 版本：2020-10-29

### 第 16 次发布

发布时间：2026-09-10 01:51:04

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyDBCustomClusterAttributes](https://cloud.tencent.com/document/api/1322/135706)

	* 新增入参：ClusterIds, ClusterName, ClusterDescription

	* <font color="#dd0000">**修改入参**：</font>ClusterId

* [ModifyDBCustomNodeAttributes](https://cloud.tencent.com/document/api/1322/135371)

	* 新增入参：NodeIds

	* <font color="#dd0000">**修改入参**：</font>NodeId


修改数据结构：

* [DBCustomClusterNode](https://cloud.tencent.com/document/api/1322/74754#DBCustomClusterNode)

	* 新增成员：LatestRunningTaskType

* [DBCustomNode](https://cloud.tencent.com/document/api/1322/74754#DBCustomNode)

	* 新增成员：LatestRunningTaskType




## 弹性 MapReduce(emr) 版本：2019-01-03

### 第 157 次发布

发布时间：2026-09-10 02:03:04

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeExportConfs](https://cloud.tencent.com/document/api/589/137841)
* [ModifyServiceParamsByExportConfs](https://cloud.tencent.com/document/api/589/137840)

新增数据结构：

* [ConfSubContext](https://cloud.tencent.com/document/api/589/33981#ConfSubContext)
* [ExportConfContext](https://cloud.tencent.com/document/api/589/33981#ExportConfContext)
* [ExportConfMeta](https://cloud.tencent.com/document/api/589/33981#ExportConfMeta)



## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 321 次发布

发布时间：2026-09-10 02:05:41

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreatePreparedPersonalEsign](https://cloud.tencent.com/document/api/1323/89386)

* [DescribeFlowTemplates](https://cloud.tencent.com/document/api/1323/74803)

	* 新增入参：ShowPreviewComponents


修改数据结构：

* [TemplateUserFlowType](https://cloud.tencent.com/document/api/1323/70369#TemplateUserFlowType)

	* 新增成员：Status




## 腾讯电子签（基础版）(essbasic) 版本：2021-05-26

### 第 277 次发布

发布时间：2026-09-10 02:07:38

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ChannelCreatePreparedPersonalEsign](https://cloud.tencent.com/document/api/1420/96160)

* [DescribeTemplates](https://cloud.tencent.com/document/api/1420/61521)

	* 新增入参：ShowPreviewComponents




## 腾讯电子签（基础版）(essbasic) 版本：2020-12-22



## iOA 零信任安全管理系统(ioa) 版本：2022-06-01

### 第 44 次发布

发布时间：2026-09-10 02:15:17

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [DescribeBusinessResourceData](https://cloud.tencent.com/document/api/1092/102488#DescribeBusinessResourceData)

	* 新增成员：ConnectivityCheckSwitch, ConnectivityCheckInterval, ConnectivityCheckIntervalUnit, URLAuditState, URLAuditId, URLPath, ReachableType, APISecretName, APISecretKey, EnableSensitiveRes, EnableIPPolicy, IPPolicyAttr, IPPolicyIds, IPPolicyNames, EnableUserAgent, UserAgentAttr, UserAgentIds, UserAgentNames




## 文字识别(ocr) 版本：2018-11-19

### 第 266 次发布

发布时间：2026-09-10 02:39:11

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [VerifyScenePhoto](https://cloud.tencent.com/document/api/866/131193)

	* 新增入参：ReasoningPrompt, ReasoningConfig

	* 新增出参：Template, ReasoningResult


新增数据结构：

* [ReasoningConfig](https://cloud.tencent.com/document/api/866/33527#ReasoningConfig)
* [ReasoningResult](https://cloud.tencent.com/document/api/866/33527#ReasoningResult)



## 实时音视频(trtc) 版本：2019-07-22

### 第 153 次发布

发布时间：2026-09-10 03:12:28

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ControlAIConversation](https://cloud.tencent.com/document/api/647/109408)

	* 新增入参：TransparentData


新增数据结构：

* [TransparentData](https://cloud.tencent.com/document/api/647/44055#TransparentData)



