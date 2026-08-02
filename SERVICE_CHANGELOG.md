# Release v1.3.150

## 文件存储(cfs) 版本：2019-07-19

### 第 55 次发布

发布时间：2026-08-03 01:24:59

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [MountInfo](https://cloud.tencent.com/document/api/582/38175#MountInfo)

	* 新增成员：ServerList, ServerListTruncated




## 负载均衡(clb) 版本：2018-03-17

### 第 156 次发布

发布时间：2026-08-03 01:29:13

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateModelRouter](https://cloud.tencent.com/document/api/214/133217)

	* 新增入参：ModelRouterBillingConfig, ClientToken

	* <font color="#dd0000">**修改入参**：</font>ClusterInfo

* [DisassociateModelsFromModelRouter](https://cloud.tencent.com/document/api/214/133657)

	* <font color="#dd0000">**修改入参**：</font>Models

* [ModifyModelRouterAttributes](https://cloud.tencent.com/document/api/214/133203)

	* 新增入参：Bandwidth


新增数据结构：

* [ClusterInfoInput](https://cloud.tencent.com/document/api/214/30694#ClusterInfoInput)
* [ModelRouterBillingConfigInput](https://cloud.tencent.com/document/api/214/30694#ModelRouterBillingConfigInput)
* [ModelRouterModelToDisassociate](https://cloud.tencent.com/document/api/214/30694#ModelRouterModelToDisassociate)
* [RoutingStrategyArgs](https://cloud.tencent.com/document/api/214/30694#RoutingStrategyArgs)

修改数据结构：

* [Coefficient](https://cloud.tencent.com/document/api/214/30694#Coefficient)

	* 新增成员：InputCacheCreationCoefficient

* [ModelRouterDetail](https://cloud.tencent.com/document/api/214/30694#ModelRouterDetail)

	* 新增成员：Bandwidth, EipAddressId

* [ModelRouterModel](https://cloud.tencent.com/document/api/214/30694#ModelRouterModel)

	* 新增成员：Order, Weight

* [ModelRouterSet](https://cloud.tencent.com/document/api/214/30694#ModelRouterSet)

	* 新增成员：Bandwidth, EipAddressId

* [RouterSettingWithFallBack](https://cloud.tencent.com/document/api/214/30694#RouterSettingWithFallBack)

	* 新增成员：NumRetries, RoutingStrategyArgs

* [RouterSettingWithoutFallBack](https://cloud.tencent.com/document/api/214/30694#RouterSettingWithoutFallBack)

	* 新增成员：CrossModelGroupRoutingStrategy, RoutingStrategyArgs, NumRetries

* [ServiceProvider](https://cloud.tencent.com/document/api/214/30694#ServiceProvider)

	* 新增成员：Order, Weight, AssociationStatus




## 暴露面管理服务(ctem) 版本：2023-11-28

### 第 20 次发布

发布时间：2026-08-03 01:38:02

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [DisplayHttp](https://cloud.tencent.com/document/api/1755/120320#DisplayHttp)

	* 新增成员：AvailabilityTag

* [DisplaySubDomain](https://cloud.tencent.com/document/api/1755/120320#DisplaySubDomain)

	* 新增成员：AvailabilityTag




## TDSQL-C MySQL 版(cynosdb) 版本：2019-01-07

### 第 184 次发布

发布时间：2026-08-03 01:44:47

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [TransferClusterPrepayToPostpay](https://cloud.tencent.com/document/api/1003/135370)
* [TransferStoragePrepayToPostpay](https://cloud.tencent.com/document/api/1003/135369)



## 云数据库独享集群(dbdc) 版本：2020-10-29

### 第 10 次发布

发布时间：2026-08-03 01:48:49

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [ModifyDBCustomNodeAttributes](https://cloud.tencent.com/document/api/1322/135371)



## 数据传输服务(dts) 版本：2021-12-06

### 第 58 次发布

发布时间：2026-08-03 01:56:41

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [Database](https://cloud.tencent.com/document/api/571/82108#Database)

	* 新增成员：SchemaMode




## 数据传输服务(dts) 版本：2018-03-30



## Elasticsearch Service(es) 版本：2025-01-01



## Elasticsearch Service(es) 版本：2018-04-16

### 第 106 次发布

发布时间：2026-08-03 02:01:02

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CheckUpdateInstance](https://cloud.tencent.com/document/api/845/135373)
* [ModifyAutoScaleDiskInfo](https://cloud.tencent.com/document/api/845/135372)



## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 313 次发布

发布时间：2026-08-03 02:02:19

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateEmbedWebUrl](https://cloud.tencent.com/document/api/1323/95590)

	* 新增入参：ApplicationId




## 物联网开发平台(iotexplorer) 版本：2019-04-23

### 第 152 次发布

发布时间：2026-08-03 02:12:54

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateTWeTalkAgent](https://cloud.tencent.com/document/api/1081/134226)

	* 新增入参：EventCallbackConfig

* [DescribeLicenseOverview](https://cloud.tencent.com/document/api/1081/131699)

	* 新增出参：Data

* [ModifyTWeTalkAgent](https://cloud.tencent.com/document/api/1081/134221)

	* 新增入参：EventCallbackConfig


新增数据结构：

* [ActivationLicense](https://cloud.tencent.com/document/api/1081/34988#ActivationLicense)
* [LicenseOverview](https://cloud.tencent.com/document/api/1081/34988#LicenseOverview)
* [TalkEventCallbackConfig](https://cloud.tencent.com/document/api/1081/34988#TalkEventCallbackConfig)

修改数据结构：

* [TalkAgentInfo](https://cloud.tencent.com/document/api/1081/34988#TalkAgentInfo)

	* 新增成员：EventCallbackConfig




## 媒体处理(mps) 版本：2019-06-12

### 第 228 次发布

发布时间：2026-07-31 15:24:32

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateAiFissionTask](https://cloud.tencent.com/document/api/862/135362)

新增数据结构：

* [AiFissionInput](https://cloud.tencent.com/document/api/862/37615#AiFissionInput)
* [CustomModel](https://cloud.tencent.com/document/api/862/37615#CustomModel)
* [FissionTaskInfo](https://cloud.tencent.com/document/api/862/37615#FissionTaskInfo)

### 第 227 次发布

发布时间：2026-07-31 14:54:27

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CloneVoice](https://cloud.tencent.com/document/api/862/135349)
* [TextToSpeech](https://cloud.tencent.com/document/api/862/135348)



## 容器服务(tke) 版本：2022-05-01

### 第 29 次发布

发布时间：2026-08-03 03:04:23

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateNodePool](https://cloud.tencent.com/document/api/457/106086)

	* 新增入参：SkipValidateOptions




## 容器服务(tke) 版本：2018-05-25

### 第 236 次发布

发布时间：2026-08-03 03:01:47

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateClusterNodePool](https://cloud.tencent.com/document/api/457/49436)

	* 新增入参：SkipValidateOptions




## 云点播(vod) 版本：2024-07-18



## 云点播(vod) 版本：2018-07-17

### 第 278 次发布

发布时间：2026-08-03 03:13:06

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateKnowledgeBase](https://cloud.tencent.com/document/api/266/135377)
* [DeleteKnowledgeBase](https://cloud.tencent.com/document/api/266/135376)
* [DescribeKnowledgeBases](https://cloud.tencent.com/document/api/266/135375)
* [ModifyKnowledgeBase](https://cloud.tencent.com/document/api/266/135374)

修改接口：

* [CreateMPSTemplate](https://cloud.tencent.com/document/api/266/122580)

	* 新增入参：EmbedSubtitleTemplate

* [CreateProcedureTemplate](https://cloud.tencent.com/document/api/266/33897)

	* 新增入参：ImportMediaKnowledgeTaskSet

* [DescribeLLMComprehendTemplates](https://cloud.tencent.com/document/api/266/128361)

	* 新增入参：Sort

* [ImportMediaKnowledge](https://cloud.tencent.com/document/api/266/126286)

	* 新增入参：KnowledgeBaseIds

* [ModifyMPSTemplate](https://cloud.tencent.com/document/api/266/122577)

	* 新增入参：EmbedSubtitleTemplate

* [ModifyMediaInfo](https://cloud.tencent.com/document/api/266/31762)

	* 新增入参：DeleteKnowledgeBases, ClearKnowledgeBases

* [ProcessMediaByProcedure](https://cloud.tencent.com/document/api/266/34782)

	* 新增出参：ImportMediaKnowledgeTaskIdSet

* [ResetProcedureTemplate](https://cloud.tencent.com/document/api/266/33894)

	* 新增入参：ImportMediaKnowledgeTaskSet

* [SearchMediaBySemantics](https://cloud.tencent.com/document/api/266/126287)

	* 新增入参：KnowledgeBaseId


新增数据结构：

* [ImportMediaKnowledgeTaskInput](https://cloud.tencent.com/document/api/266/31773#ImportMediaKnowledgeTaskInput)
* [KnowledgeBaseDetail](https://cloud.tencent.com/document/api/266/31773#KnowledgeBaseDetail)
* [MPSEmbedSubtitleTemplate](https://cloud.tencent.com/document/api/266/31773#MPSEmbedSubtitleTemplate)
* [MPSEmbedSubtitleTemplateForUpdate](https://cloud.tencent.com/document/api/266/31773#MPSEmbedSubtitleTemplateForUpdate)
* [MPSSubtitleBoardConfig](https://cloud.tencent.com/document/api/266/31773#MPSSubtitleBoardConfig)
* [MPSSubtitleEmbedConfig](https://cloud.tencent.com/document/api/266/31773#MPSSubtitleEmbedConfig)
* [MPSSubtitleLayoutConfig](https://cloud.tencent.com/document/api/266/31773#MPSSubtitleLayoutConfig)
* [MPSSubtitleOutlineConfig](https://cloud.tencent.com/document/api/266/31773#MPSSubtitleOutlineConfig)
* [MPSSubtitleShadowConfig](https://cloud.tencent.com/document/api/266/31773#MPSSubtitleShadowConfig)

修改数据结构：

* [ImportMediaKnowledgeTask](https://cloud.tencent.com/document/api/266/31773#ImportMediaKnowledgeTask)

	* 新增成员：FileId, Input

* [KnowledgeBasesInfo](https://cloud.tencent.com/document/api/266/31773#KnowledgeBasesInfo)

	* 新增成员：KnowledgeBaseDetails

* [LLMComprehendTemplateItem](https://cloud.tencent.com/document/api/266/31773#LLMComprehendTemplateItem)

	* 新增成员：Type

* [MPSRawSmartSubtitleParameter](https://cloud.tencent.com/document/api/266/31773#MPSRawSmartSubtitleParameter)

	* 新增成员：SelectingSubtitleAreasConfig, SubtitleEmbedId, SpeakerMode, SpeakerLabel

* [MPSSmartEraseSubtitleConfig](https://cloud.tencent.com/document/api/266/31773#MPSSmartEraseSubtitleConfig)

	* 新增成员：SubtitleEmbedId

* [MPSSmartSubtitleTemplate](https://cloud.tencent.com/document/api/266/31773#MPSSmartSubtitleTemplate)

	* 新增成员：SubtitleEmbedId

* [MPSSmartSubtitleTemplateForUpdate](https://cloud.tencent.com/document/api/266/31773#MPSSmartSubtitleTemplateForUpdate)

	* 新增成员：SubtitleEmbedId

* [ProcedureTemplate](https://cloud.tencent.com/document/api/266/31773#ProcedureTemplate)

	* 新增成员：ImportMediaKnowledgeTaskSet

* [SemanticsSearchResult](https://cloud.tencent.com/document/api/266/31773#SemanticsSearchResult)

	* 新增成员：Title




