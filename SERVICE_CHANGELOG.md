# Release v1.3.94

## 云数据库 MySQL(cdb) 版本：2017-03-20

### 第 217 次发布

发布时间：2026-05-11 01:12:46

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [UpgradeAnalysisInstanceVersionInfo](https://cloud.tencent.com/document/api/236/15878#UpgradeAnalysisInstanceVersionInfo)

修改数据结构：

* [InstanceInfo](https://cloud.tencent.com/document/api/236/15878#InstanceInfo)

	* 新增成员：AnalysisUpgradeVersionInfo




## 日志服务(cls) 版本：2020-10-16

### 第 159 次发布

发布时间：2026-05-11 01:17:20

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateRecordingRuleTask](https://cloud.tencent.com/document/api/614/131371)
* [CreateRecordingRuleYamlTask](https://cloud.tencent.com/document/api/614/131370)
* [DeleteRecordingRuleTask](https://cloud.tencent.com/document/api/614/131369)
* [DeleteRecordingRuleYamlTask](https://cloud.tencent.com/document/api/614/131368)
* [DescribeRecordingRuleTask](https://cloud.tencent.com/document/api/614/131367)
* [DescribeRecordingRuleYamlTask](https://cloud.tencent.com/document/api/614/131366)
* [ModifyRecordingRuleTask](https://cloud.tencent.com/document/api/614/131365)
* [ModifyRecordingRuleYamlTask](https://cloud.tencent.com/document/api/614/131364)

修改接口：

* [CreateTopic](https://cloud.tencent.com/document/api/614/56456)

	* 新增入参：BillingMode

* [ModifyTopic](https://cloud.tencent.com/document/api/614/56453)

	* 新增入参：BillingMode


新增数据结构：

* [RecordingRuleTaskInfo](https://cloud.tencent.com/document/api/614/56471#RecordingRuleTaskInfo)
* [RecordingRuleYamlTaskInfo](https://cloud.tencent.com/document/api/614/56471#RecordingRuleYamlTaskInfo)

修改数据结构：

* [TopicInfo](https://cloud.tencent.com/document/api/614/56471#TopicInfo)

	* 新增成员：BillingMode, NewBillingMode




## 时序数据库 CTSDB(ctsdb) 版本：2023-02-02

### 第 3 次发布

发布时间：2026-05-11 01:19:46

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeClusterDetail](https://cloud.tencent.com/document/api/652/128076)

	* 新增入参：ClusterID




## 弹性 MapReduce(emr) 版本：2019-01-03

### 第 140 次发布

发布时间：2026-05-11 01:28:04

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateCluster](https://cloud.tencent.com/document/api/589/83953)

	* 新增入参：EnableCbsSysEncryptFlag

* [CreateInstance](https://cloud.tencent.com/document/api/589/34261)

	* 新增入参：CbsSysEncrypt




## 腾讯电子签（基础版）(essbasic) 版本：2021-05-26

### 第 256 次发布

发布时间：2026-05-11 01:29:45

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ChannelCreateEmbedWebUrl](https://cloud.tencent.com/document/api/1420/87919)




## 腾讯电子签（基础版）(essbasic) 版本：2020-12-22



## iOA 零信任安全管理系统(ioa) 版本：2022-06-01

### 第 36 次发布

发布时间：2026-05-11 01:32:42

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [BindBusinessResourceConnectorGroup](https://cloud.tencent.com/document/api/1092/131372)



## 腾讯云智能体开发平台(lke) 版本：2023-11-30

### 第 88 次发布

发布时间：2026-05-11 01:38:17

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [AgentPluginInfo](https://cloud.tencent.com/document/api/1759/105104#AgentPluginInfo)

	* 新增成员：AuthMode, AuthType, AuthConfigStatus, PluginClass, PluginStatus

* [AgentToolInfo](https://cloud.tencent.com/document/api/1759/105104#AgentToolInfo)

	* 新增成员：ToolAccessMode, IsDisabled




## 媒体处理(mps) 版本：2019-06-12

### 第 201 次发布

发布时间：2026-05-11 01:46:38

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [AigcImageExtraParam](https://cloud.tencent.com/document/api/862/37615#AigcImageExtraParam)

	* 新增成员：OutputFormat




## 消息队列 MQTT 版(mqtt) 版本：2024-05-16

### 第 28 次发布

发布时间：2026-05-11 01:49:41

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeSharedSubscriptionClient](https://cloud.tencent.com/document/api/1778/131374)
* [DescribeSharedSubscriptions](https://cloud.tencent.com/document/api/1778/131373)

新增数据结构：

* [SharedSubscriptionClient](https://cloud.tencent.com/document/api/1778/111031#SharedSubscriptionClient)



## 文字识别(ocr) 版本：2018-11-19

### 第 246 次发布

发布时间：2026-05-11 01:52:03

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除接口**：</font>

* BusInvoiceOCR
* DutyPaidProofOCR
* FinanBillOCR
* FinanBillSliceOCR
* InvoiceGeneralOCR
* QuotaInvoiceOCR
* ShipInvoiceOCR
* TollInvoiceOCR
* VatRollInvoiceOCR

<font color="#dd0000">**删除数据结构**：</font>

* BusInvoiceInfo
* DutyPaidProofInfo
* FinanBillInfo
* FinanBillSliceInfo
* InvoiceGeneralInfo
* ShipInvoiceInfo
* TollInvoiceInfo
* VatRollInvoiceInfo



## 云托管 CloudBase Run(tcbr) 版本：2022-02-17

### 第 28 次发布

发布时间：2026-05-11 02:12:53

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [VolumeConf](https://cloud.tencent.com/document/api/1243/75713#VolumeConf)

	* 新增成员：MountIP, ReadOnly, InstanceId




## 腾讯云数据库 AI 服务(tdai) 版本：2025-07-17

### 第 12 次发布

发布时间：2026-05-11 02:19:20

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [CreatingProgress](https://cloud.tencent.com/document/api/1813/123239#CreatingProgress)
* [CreatingStepInfo](https://cloud.tencent.com/document/api/1813/123239#CreatingStepInfo)

修改数据结构：

* [AgentInstance](https://cloud.tencent.com/document/api/1813/123239#AgentInstance)

	* 新增成员：CreatingProgress




## 消息队列 TDMQ(tdmq) 版本：2020-02-17

### 第 172 次发布

发布时间：2026-05-09 09:52:58

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [RouterMessageServiceSource](https://cloud.tencent.com/document/api/1179/46089#RouterMessageServiceSource)
* [RouterMessageServiceTarget](https://cloud.tencent.com/document/api/1179/46089#RouterMessageServiceTarget)
* [RouterRocketMQSource](https://cloud.tencent.com/document/api/1179/46089#RouterRocketMQSource)
* [RouterRocketMQTarget](https://cloud.tencent.com/document/api/1179/46089#RouterRocketMQTarget)
* [RouterTencentRocketMQSource](https://cloud.tencent.com/document/api/1179/46089#RouterTencentRocketMQSource)
* [RouterTencentRocketMQTarget](https://cloud.tencent.com/document/api/1179/46089#RouterTencentRocketMQTarget)

修改数据结构：

* [RocketMQRouterRuleInfo](https://cloud.tencent.com/document/api/1179/46089#RocketMQRouterRuleInfo)

	* 新增成员：SourceType, TargetType, RemarkName, AliRocketMQSource, AliRocketMQTarget, AliMessageServiceSource, AliMessageServiceTarget, TenRocketMQSource, TenRocketMQTarget, AliasName


### 第 171 次发布

发布时间：2026-05-08 15:12:58

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateRocketMQRouterRule](https://cloud.tencent.com/document/api/1179/131328)

新增数据结构：

* [RocketMQRouterRuleInfo](https://cloud.tencent.com/document/api/1179/46089#RocketMQRouterRuleInfo)
* [RouterTencentMQTTSource](https://cloud.tencent.com/document/api/1179/46089#RouterTencentMQTTSource)
* [RouterTencentMQTTTarget](https://cloud.tencent.com/document/api/1179/46089#RouterTencentMQTTTarget)



## TI-ONE 训练平台(tione) 版本：2021-11-11

### 第 115 次发布

发布时间：2026-05-11 02:28:01

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateDataSource](https://cloud.tencent.com/document/api/851/129220)

	* 新增入参：TiProjectId

* [CreateMountLimit](https://cloud.tencent.com/document/api/851/129219)

	* 新增入参：TiProjectId

* [DeleteDataSource](https://cloud.tencent.com/document/api/851/129218)

	* 新增入参：TiProjectId

* [DeleteMountLimit](https://cloud.tencent.com/document/api/851/129217)

	* 新增入参：TiProjectId

* [DescribeDataSources](https://cloud.tencent.com/document/api/851/129215)

	* 新增入参：TiProjectId

* [DescribeMountInstance](https://cloud.tencent.com/document/api/851/129214)

	* 新增入参：TiProjectId

* [DescribeMountInstances](https://cloud.tencent.com/document/api/851/129213)

	* 新增入参：TiProjectId

* [DescribeMountLimits](https://cloud.tencent.com/document/api/851/129212)

	* 新增入参：TiProjectId

* [UpdateDataSource](https://cloud.tencent.com/document/api/851/129211)

	* 新增入参：TiProjectId

* [UpdateMountLimit](https://cloud.tencent.com/document/api/851/129210)

	* 新增入参：TiProjectId




## TI-ONE 训练平台(tione) 版本：2019-10-22



## 实时音视频(trtc) 版本：2019-07-22

### 第 142 次发布

发布时间：2026-05-11 02:36:02

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeAsyncTextToSpeech](https://cloud.tencent.com/document/api/647/130581)

	* 新增出参：TotalDurationMs

* [TextToSpeech](https://cloud.tencent.com/document/api/647/122475)

	* 新增出参：TotalDurationMs




## 云点播(vod) 版本：2024-07-18



## 云点播(vod) 版本：2018-07-17

### 第 254 次发布

发布时间：2026-05-11 02:43:50

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateAigcAudioTask](https://cloud.tencent.com/document/api/266/131376)
* [DeleteAigcAdvancedCustomElement](https://cloud.tencent.com/document/api/266/131377)
* [DescribeAigcAdvancedCustomElements](https://cloud.tencent.com/document/api/266/131375)

修改接口：

* [CreateMPSTemplate](https://cloud.tencent.com/document/api/266/122580)

	* 新增入参：AIAnalysisTemplate, SmartSubtitleTemplate, SmartEraseTemplate

	* <font color="#dd0000">**修改入参**：</font>MPSCreateTemplateParams

* [ModifyMPSTemplate](https://cloud.tencent.com/document/api/266/122577)

	* 新增入参：AIAnalysisTemplate, SmartSubtitleTemplate, SmartEraseTemplate

	* <font color="#dd0000">**修改入参**：</font>MPSModifyTemplateParams

* [ProcessMedia](https://cloud.tencent.com/document/api/266/33427)

	* 新增入参：Url

* [ProcessMediaByMPS](https://cloud.tencent.com/document/api/266/121993)

	* 新增入参：AiAnalysisTask, SmartSubtitlesTask, SmartEraseTask

	* <font color="#dd0000">**修改入参**：</font>MPSProcessMediaParams


新增数据结构：

* [AigcAdvancedCustomElementInfo](https://cloud.tencent.com/document/api/266/31773#AigcAdvancedCustomElementInfo)
* [AigcAudioReferenceAudioInfo](https://cloud.tencent.com/document/api/266/31773#AigcAudioReferenceAudioInfo)
* [AigcAudioReferenceVideoInfo](https://cloud.tencent.com/document/api/266/31773#AigcAudioReferenceVideoInfo)
* [MPSAIAnalysisConfigureInfo](https://cloud.tencent.com/document/api/266/31773#MPSAIAnalysisConfigureInfo)
* [MPSAIAnalysisTemplate](https://cloud.tencent.com/document/api/266/31773#MPSAIAnalysisTemplate)
* [MPSAIAnalysisTemplateForUpdate](https://cloud.tencent.com/document/api/266/31773#MPSAIAnalysisTemplateForUpdate)
* [MPSAiAnalysisTaskInput](https://cloud.tencent.com/document/api/266/31773#MPSAiAnalysisTaskInput)
* [MPSEraseArea](https://cloud.tencent.com/document/api/266/31773#MPSEraseArea)
* [MPSEraseTimeArea](https://cloud.tencent.com/document/api/266/31773#MPSEraseTimeArea)
* [MPSOverrideEraseParameter](https://cloud.tencent.com/document/api/266/31773#MPSOverrideEraseParameter)
* [MPSRawSmartEraseParameter](https://cloud.tencent.com/document/api/266/31773#MPSRawSmartEraseParameter)
* [MPSRawSmartSubtitleParameter](https://cloud.tencent.com/document/api/266/31773#MPSRawSmartSubtitleParameter)
* [MPSSelectingSubtitleAreasConfig](https://cloud.tencent.com/document/api/266/31773#MPSSelectingSubtitleAreasConfig)
* [MPSSmartErasePrivacyConfig](https://cloud.tencent.com/document/api/266/31773#MPSSmartErasePrivacyConfig)
* [MPSSmartEraseSubtitleConfig](https://cloud.tencent.com/document/api/266/31773#MPSSmartEraseSubtitleConfig)
* [MPSSmartEraseTaskInput](https://cloud.tencent.com/document/api/266/31773#MPSSmartEraseTaskInput)
* [MPSSmartEraseTemplate](https://cloud.tencent.com/document/api/266/31773#MPSSmartEraseTemplate)
* [MPSSmartEraseTemplateForUpdate](https://cloud.tencent.com/document/api/266/31773#MPSSmartEraseTemplateForUpdate)
* [MPSSmartEraseWatermarkConfig](https://cloud.tencent.com/document/api/266/31773#MPSSmartEraseWatermarkConfig)
* [MPSSmartSubtitleTemplate](https://cloud.tencent.com/document/api/266/31773#MPSSmartSubtitleTemplate)
* [MPSSmartSubtitleTemplateForUpdate](https://cloud.tencent.com/document/api/266/31773#MPSSmartSubtitleTemplateForUpdate)
* [MPSSmartSubtitlesTaskInput](https://cloud.tencent.com/document/api/266/31773#MPSSmartSubtitlesTaskInput)
* [MPSUpdateSmartEraseSubtitleConfig](https://cloud.tencent.com/document/api/266/31773#MPSUpdateSmartEraseSubtitleConfig)
* [MPSUpdateSmartEraseWatermarkConfig](https://cloud.tencent.com/document/api/266/31773#MPSUpdateSmartEraseWatermarkConfig)



## Web 应用防火墙(waf) 版本：2018-01-25

### 第 150 次发布

发布时间：2026-05-11 02:52:58

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [PromptDetectResult](https://cloud.tencent.com/document/api/627/53609#PromptDetectResult)

	* 新增成员：Category




## 数据开发治理平台 WeData(wedata) 版本：2025-08-06



## 数据开发治理平台 WeData(wedata) 版本：2021-08-20

### 第 190 次发布

发布时间：2026-05-11 02:56:10

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**预下线接口**：</font>

* SubmitWorkflow



