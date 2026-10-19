# Release v1.3.195

## 腾讯云智能体开发平台(adp) 版本：2026-05-20

### 第 31 次发布

发布时间：2026-10-20 01:09:13

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CallMCPTool](https://cloud.tencent.com/document/api/1759/139588)
* [ReadMCPResource](https://cloud.tencent.com/document/api/1759/139587)

新增数据结构：

* [MCPResourceContents](https://cloud.tencent.com/document/api/1759/132545#MCPResourceContents)



## 云数据库 MySQL(cdb) 版本：2017-03-20

### 第 234 次发布

发布时间：2026-10-20 01:16:21

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateAccounts](https://cloud.tencent.com/document/api/236/17502)

	* <font color="#dd0000">**修改入参**：</font>Accounts, Password

* [DeleteAccounts](https://cloud.tencent.com/document/api/236/17501)

	* <font color="#dd0000">**修改入参**：</font>Accounts




## 云防火墙(cfw) 版本：2019-09-04

### 第 119 次发布

发布时间：2026-10-20 01:19:04

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [NatFwSwitchDetailS](https://cloud.tencent.com/document/api/1132/49071#NatFwSwitchDetailS)

	* 新增成员：ForbidMode, ForbidCode




## 日志服务(cls) 版本：2020-10-16

### 第 183 次发布

发布时间：2026-10-20 01:22:11

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [AnonymousLoginInfo](https://cloud.tencent.com/document/api/614/56471#AnonymousLoginInfo)

	* 新增成员：RoleArn

	* <font color="#dd0000">**修改成员**：</font>SecretId

* [AuthRoleInfo](https://cloud.tencent.com/document/api/614/56471#AuthRoleInfo)

	* 新增成员：RoleArn

	* <font color="#dd0000">**修改成员**：</font>SecretId

* [ConsoleAccount](https://cloud.tencent.com/document/api/614/56471#ConsoleAccount)

	* 新增成员：RoleArn

	* <font color="#dd0000">**修改成员**：</font>SecretId, SecretKey




## 云安全一体化平台(csip) 版本：2022-11-21

### 第 110 次发布

发布时间：2026-10-20 01:24:33

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeCSCPayInfo](https://cloud.tencent.com/document/api/664/136896)

	* 新增入参：MemberId, NeedQuotaUsed

	* 新增出参：QuotaUsedError

* [DescribeEDRRuleList](https://cloud.tencent.com/document/api/664/133469)

	* 新增入参：FileOperationRules

* [ExportEDRRules](https://cloud.tencent.com/document/api/664/135940)

	* 新增入参：FileOperationRules

* [ModifyEDRRule](https://cloud.tencent.com/document/api/664/133466)

	* 新增入参：FileOperationRules


新增数据结构：

* [BaselinePolicyBrief](https://cloud.tencent.com/document/api/664/90825#BaselinePolicyBrief)
* [RuleContentFileOperation](https://cloud.tencent.com/document/api/664/90825#RuleContentFileOperation)

修改数据结构：

* [BaselineAggregatedPolicy](https://cloud.tencent.com/document/api/664/90825#BaselineAggregatedPolicy)

	* 新增成员：AppID

* [BaselineMainTask](https://cloud.tencent.com/document/api/664/90825#BaselineMainTask)

	* 新增成员：PolicyList

* [EDRRule](https://cloud.tencent.com/document/api/664/90825#EDRRule)

	* 新增成员：FileOperationRules

* [LicenseStatusItem](https://cloud.tencent.com/document/api/664/90825#LicenseStatusItem)

	* 新增成员：OrderType

* [OrderQuotaInfo](https://cloud.tencent.com/document/api/664/90825#OrderQuotaInfo)

	* 新增成员：QuotaGranted

	* <font color="#dd0000">**修改成员**：</font>QuotaKey, QuotaNum, QuotaUsed

* [WebhookNotifyItem](https://cloud.tencent.com/document/api/664/90825#WebhookNotifyItem)

	* 新增成员：Options




## 主机安全(cwp) 版本：2018-02-28

### 第 171 次发布

发布时间：2026-10-20 01:29:04

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [BashEventNew](https://cloud.tencent.com/document/api/296/19867#BashEventNew)

	* 新增成员：EventCount, FirstDetectTime, LatestDetectTime

* [BashEventsInfoNew](https://cloud.tencent.com/document/api/296/19867#BashEventsInfoNew)

	* 新增成员：EventCount, FirstDetectTime, LatestDetectTime, FileProcessMd5, FileAction, FilePath, FileMd5, FileName, FileSize, FileModifyTime, FilePermission, FileAccessTime, BashType

* [BashPolicy](https://cloud.tencent.com/document/api/296/19867#BashPolicy)

	* 新增成员：MatchType

* [PolicyRules](https://cloud.tencent.com/document/api/296/19867#PolicyRules)

	* 新增成员：FilePath, FileActions




## 数据库智能管家 DBbrain(dbbrain) 版本：2021-05-27

### 第 67 次发布

发布时间：2026-10-20 01:33:28

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeBlockingLogs](https://cloud.tencent.com/document/api/1130/139590)

新增数据结构：

* [BlockingLogItem](https://cloud.tencent.com/document/api/1130/57812#BlockingLogItem)



## 数据库智能管家 DBbrain(dbbrain) 版本：2019-10-16



## 数据传输服务(dts) 版本：2021-12-06

### 第 61 次发布

发布时间：2026-10-20 01:38:22

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [SkipDdlOnError](https://cloud.tencent.com/document/api/571/82108#SkipDdlOnError)

修改数据结构：

* [MigrateOption](https://cloud.tencent.com/document/api/571/82108#MigrateOption)

	* 新增成员：SkipDdlOnError

* [Options](https://cloud.tencent.com/document/api/571/82108#Options)

	* 新增成员：SkipDdlOnError




## 数据传输服务(dts) 版本：2018-03-30



## 智能视图计算平台(iss) 版本：2023-05-17

### 第 36 次发布

发布时间：2026-10-20 01:49:21

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [ControlDeviceStreamData](https://cloud.tencent.com/document/api/1344/95952#ControlDeviceStreamData)

	* 新增成员：FMp4

* [PlayRecordData](https://cloud.tencent.com/document/api/1344/95952#PlayRecordData)

	* 新增成员：FMp4




## 文字识别(ocr) 版本：2018-11-19

### 第 273 次发布

发布时间：2026-10-20 01:57:37

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeQuestionMarkAgentJob](https://cloud.tencent.com/document/api/866/128274)

	* 新增出参：QuestionInfo

* [SubmitQuestionMarkAgentJob](https://cloud.tencent.com/document/api/866/128273)

	* 新增入参：EnableImageCrop

* [SubmitQuestionSplitJob](https://cloud.tencent.com/document/api/866/137427)

	* 新增入参：EnableImageCrop


新增数据结构：

* [AnswerElement](https://cloud.tencent.com/document/api/866/33527#AnswerElement)

修改数据结构：

* [Element](https://cloud.tencent.com/document/api/866/33527#Element)

	* 新增成员：AnswerElement

* [QuestionInfo](https://cloud.tencent.com/document/api/866/33527#QuestionInfo)

	* 新增成员：ReturnImageUrl

* [ResultList](https://cloud.tencent.com/document/api/866/33527#ResultList)

	* 新增成员：RightAnswer, Subject, KnowledgePoints




## 云开发 CloudBase(tcb) 版本：2018-06-08

### 第 171 次发布

发布时间：2026-10-20 02:04:51

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [LocalizedTemplate](https://cloud.tencent.com/document/api/876/34822#LocalizedTemplate)

	* 新增成员：SubjectZhCN, SubjectEnUS, SenderNameZhCN, SenderNameEnUS




## 高性能计算平台(thpc) 版本：2023-03-21

### 第 46 次发布

发布时间：2026-10-20 02:10:33

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [TemplateOverrides](https://cloud.tencent.com/document/api/1527/89579#TemplateOverrides)

	* 新增成员：EnableJobLoadAutoscaling




## 高性能计算平台(thpc) 版本：2022-04-01



## 高性能计算平台(thpc) 版本：2021-11-09



## TI-ONE 训练平台(tione) 版本：2021-11-11

### 第 142 次发布

发布时间：2026-10-20 02:10:58

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateNotebookImage](https://cloud.tencent.com/document/api/851/139547)

	* 新增入参：TiProjectId, RegisterToImageCenter

* [CreatePresignedNotebookUrl](https://cloud.tencent.com/document/api/851/104306)

	* 新增入参：TiProjectId

* [DeleteNotebook](https://cloud.tencent.com/document/api/851/95663)

	* 新增入参：TiProjectId

* [DescribeBillingSpecsPrice](https://cloud.tencent.com/document/api/851/112647)

	* 新增入参：TiProjectId

* [DescribeEvents](https://cloud.tencent.com/document/api/851/118481)

	* 新增入参：TiProjectId

* [DescribePublicAlgoVersionList](https://cloud.tencent.com/document/api/851/129362)

	* 新增入参：TiProjectId

* [DescribeSubAccountLinuxUserInfos](https://cloud.tencent.com/document/api/851/127579)

	* 新增入参：TiProjectId

* [DescribeTrainingTaskPodUrl](https://cloud.tencent.com/document/api/851/131906)

	* 新增入参：TiProjectId

* [EnableTrainingTaskClsConfig](https://cloud.tencent.com/document/api/851/139315)

	* 新增入参：TiProjectId

* [ModifyModelServiceAuthorization](https://cloud.tencent.com/document/api/851/119881)

	* 新增入参：TiProjectId

* [ModifyNotebook](https://cloud.tencent.com/document/api/851/127303)

	* 新增入参：TiProjectId

* [ModifyNotebookTags](https://cloud.tencent.com/document/api/851/118331)

	* 新增入参：TiProjectId

* [ModifyServiceGroupWeights](https://cloud.tencent.com/document/api/851/127689)

	* 新增入参：TiProjectId

* [StartNotebook](https://cloud.tencent.com/document/api/851/95650)

	* 新增入参：TiProjectId

* [StopNotebook](https://cloud.tencent.com/document/api/851/95660)

	* 新增入参：TiProjectId

* [UpdateSubAccountLinuxUserInfo](https://cloud.tencent.com/document/api/851/127578)

	* 新增入参：TiProjectId


新增数据结构：

* [SmartRouter](https://cloud.tencent.com/document/api/851/75051#SmartRouter)
* [SmartRouterComponent](https://cloud.tencent.com/document/api/851/75051#SmartRouterComponent)
* [SmartRouterResourceConfig](https://cloud.tencent.com/document/api/851/75051#SmartRouterResourceConfig)
* [SmartRouterResourceSpec](https://cloud.tencent.com/document/api/851/75051#SmartRouterResourceSpec)

修改数据结构：

* [ServiceGroup](https://cloud.tencent.com/document/api/851/75051#ServiceGroup)

	* 新增成员：SmartRouter




## TI-ONE 训练平台(tione) 版本：2019-10-22



## 消息队列 RocketMQ 版(trocket) 版本：2023-03-08

### 第 59 次发布

发布时间：2026-10-20 02:13:44

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateConsumerGroup](https://cloud.tencent.com/document/api/1493/97943)

	* 新增入参：DeadLetterTargetTopic

* [DescribeConsumerGroup](https://cloud.tencent.com/document/api/1493/97941)

	* 新增出参：DeadLetterTargetTopic

* [DescribeMessageTrace](https://cloud.tencent.com/document/api/1493/114302)

	* 新增入参：QueryRetryMessage, GroupId

* [DescribeTopic](https://cloud.tencent.com/document/api/1493/97945)

	* 新增出参：QueueNum

* [ModifyConsumerGroup](https://cloud.tencent.com/document/api/1493/97940)

	* 新增入参：DeadLetterTargetTopic

* [SendMessage](https://cloud.tencent.com/document/api/1493/126164)

	* 新增入参：Priority


修改数据结构：

* [ConsumeGroupItem](https://cloud.tencent.com/document/api/1493/96031#ConsumeGroupItem)

	* 新增成员：ConsumeModel, DeadLetterTargetTopic

* [TopicItem](https://cloud.tencent.com/document/api/1493/96031#TopicItem)

	* 新增成员：AutoExpireDelete, AutoExpireTime




## 云点播(vod) 版本：2024-07-18



## 云点播(vod) 版本：2018-07-17

### 第 291 次发布

发布时间：2026-10-20 02:17:02

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [VideoDepthEstimation](https://cloud.tencent.com/document/api/266/139591)

新增数据结构：

* [VideoDepthEstimationOutputConfig](https://cloud.tencent.com/document/api/266/31773#VideoDepthEstimationOutputConfig)



## 数据开发治理平台 WeData(wedata) 版本：2025-08-06

### 第 28 次发布

发布时间：2026-10-20 02:24:43

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateResourceGroup](https://cloud.tencent.com/document/api/1267/123984)

	* 新增入参：Tags


新增数据结构：

* [ResourceTag](https://cloud.tencent.com/document/api/1267/123643#ResourceTag)

修改数据结构：

* [LineageNodeInfo](https://cloud.tencent.com/document/api/1267/123643#LineageNodeInfo)

	* 新增成员：SourceColumnName, ColumnName

* [Project](https://cloud.tencent.com/document/api/1267/123643#Project)

	* 新增成员：DisplayStatus




## 数据开发治理平台 WeData(wedata) 版本：2021-08-20

### 第 207 次发布

发布时间：2026-10-20 02:22:54

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyTaskInfoDs](https://cloud.tencent.com/document/api/1267/120021)

	* 新增入参：MinuteOffset

* [RenewWorkflowSchedulerInfoDs](https://cloud.tencent.com/document/api/1267/110374)

	* 新增入参：MinuteOffset


修改数据结构：

* [Apply](https://cloud.tencent.com/document/api/1267/76336#Apply)

	* 新增成员：ApproveCode




