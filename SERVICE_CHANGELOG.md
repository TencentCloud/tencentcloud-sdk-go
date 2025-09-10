# Release v1.1.26

## 云托付物理服务器(chc) 版本：2023-04-18

### 第 8 次发布

发布时间：2025-09-11 01:22:31

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreatePersonnelVisitWorkOrder](https://cloud.tencent.com/document/api/1448/117187)

	* 新增入参：CarSet

* [DescribePersonnelVisitWorkOrderDetail](https://cloud.tencent.com/document/api/1448/117170)

	* 新增出参：CarSet


新增数据结构：

* [PersonnelVisitCar](https://cloud.tencent.com/document/api/1448/117193#PersonnelVisitCar)

修改数据结构：

* [WorkOrderData](https://cloud.tencent.com/document/api/1448/117193#WorkOrderData)

	* 新增成员：TicketId




## 消息队列 CKafka 版(ckafka) 版本：2019-08-19

### 第 129 次发布

发布时间：2025-09-11 01:23:24

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreatePostPaidInstance](https://cloud.tencent.com/document/api/597/94259)

	* <font color="#dd0000">**修改入参**：</font>SubnetId

* [CreateRoute](https://cloud.tencent.com/document/api/597/70172)

	* 新增入参：Note


修改数据结构：

* [Route](https://cloud.tencent.com/document/api/597/40861#Route)

	* 新增成员：Note

* [ZoneResponse](https://cloud.tencent.com/document/api/597/40861#ZoneResponse)

	* <font color="#dd0000">**删除成员**：</font>Version




## 日志服务(cls) 版本：2020-10-16

### 第 134 次发布

发布时间：2025-09-11 01:26:02

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateAlarm](https://cloud.tencent.com/document/api/614/56466)

	* <font color="#dd0000">**修改入参**：</font>AlarmNoticeIds


新增数据结构：

* [MonitorNotice](https://cloud.tencent.com/document/api/614/56471#MonitorNotice)
* [MonitorNoticeRule](https://cloud.tencent.com/document/api/614/56471#MonitorNoticeRule)

修改数据结构：

* [AlarmInfo](https://cloud.tencent.com/document/api/614/56471#AlarmInfo)

	* 新增成员：MonitorNotice

* [AlertHistoryRecord](https://cloud.tencent.com/document/api/614/56471#AlertHistoryRecord)

	* 新增成员：SendType




## 暴露面管理服务(ctem) 版本：2023-11-28

### 第 8 次发布

发布时间：2025-09-11 01:28:52

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [DisplayAsset](https://cloud.tencent.com/document/api/1755/120320#DisplayAsset)

	* 新增成员：Ports, Services, Domains, LastModify




## 专线接入(dc) 版本：2018-04-10

### 第 36 次发布

发布时间：2025-09-11 01:36:19

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [DirectConnectTunnel](https://cloud.tencent.com/document/api/216/18418#DirectConnectTunnel)

	* 新增成员：AccessPointName, AccessPointId




## 弹性 MapReduce(emr) 版本：2019-01-03

### 第 119 次发布

发布时间：2025-09-11 01:44:56

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateCloudInstance](https://cloud.tencent.com/document/api/589/113701)

	* 新增入参：DefaultMetaVersion, NeedCdbAudit

* [CreateCluster](https://cloud.tencent.com/document/api/589/83953)

	* 新增入参：DefaultMetaVersion, NeedCdbAudit

* [CreateInstance](https://cloud.tencent.com/document/api/589/34261)

	* 新增入参：DefaultMetaVersion, NeedCdbAudit

* [DescribeNodeDataDisks](https://cloud.tencent.com/document/api/589/115191)

	* 新增入参：Scene

	* 新增出参：MaxThroughputPerformance

* [InquiryPriceCreateInstance](https://cloud.tencent.com/document/api/589/33980)

	* 新增入参：DefaultMetaVersion, NeedCdbAudit


新增数据结构：

* [TagInfo](https://cloud.tencent.com/document/api/589/33981#TagInfo)

修改数据结构：

* [CBSInstance](https://cloud.tencent.com/document/api/589/33981#CBSInstance)

	* 新增成员：Tags, ThroughputPerformance

* [LoadAutoScaleStrategy](https://cloud.tencent.com/document/api/589/33981#LoadAutoScaleStrategy)

	* 新增成员：GraceDownProtectFlag

* [TimeAutoScaleStrategy](https://cloud.tencent.com/document/api/589/33981#TimeAutoScaleStrategy)

	* 新增成员：GraceDownProtectFlag




## Elasticsearch Service(es) 版本：2025-01-01

### 第 8 次发布

发布时间：2025-09-11 01:46:43

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [ToolCall](https://cloud.tencent.com/document/api/845/117811#ToolCall)
* [ToolCallFunction](https://cloud.tencent.com/document/api/845/117811#ToolCallFunction)

修改数据结构：

* [Message](https://cloud.tencent.com/document/api/845/117811#Message)

	* 新增成员：ToolCallId, ToolCalls

* [OutputMessage](https://cloud.tencent.com/document/api/845/117811#OutputMessage)

	* 新增成员：ToolCalls




## Elasticsearch Service(es) 版本：2018-04-16



## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 246 次发布

发布时间：2025-09-11 01:46:53

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateMultiFlowSignQRCode](https://cloud.tencent.com/document/api/1323/75450)

	* 新增入参：QrCodeName, QrCodeExpiredOn


修改数据结构：

* [MiniAppCreateFlowOption](https://cloud.tencent.com/document/api/1323/70369#MiniAppCreateFlowOption)

	* 新增成员：ForbidEditFlow




## 腾讯电子签（基础版）(essbasic) 版本：2021-05-26

### 第 234 次发布

发布时间：2025-09-11 01:48:14

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ChannelCreateMultiFlowSignQRCode](https://cloud.tencent.com/document/api/1420/75452)

	* 新增入参：QrCodeName, QrCodeExpiredOn




## 腾讯电子签（基础版）(essbasic) 版本：2020-12-22



## 游戏多媒体引擎(gme) 版本：2018-07-11

### 第 42 次发布

发布时间：2025-09-11 01:50:50

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [ControlAIConversation](https://cloud.tencent.com/document/api/607/123144)
* [DeleteVoicePrint](https://cloud.tencent.com/document/api/607/123143)
* [DescribeAIConversation](https://cloud.tencent.com/document/api/607/123142)
* [DescribeVoicePrint](https://cloud.tencent.com/document/api/607/123141)
* [RegisterVoicePrint](https://cloud.tencent.com/document/api/607/123140)
* [StartAIConversation](https://cloud.tencent.com/document/api/607/123139)
* [StopAIConversation](https://cloud.tencent.com/document/api/607/123138)
* [UpdateAIConversation](https://cloud.tencent.com/document/api/607/123137)
* [UpdateVoicePrint](https://cloud.tencent.com/document/api/607/123136)

新增数据结构：

* [AgentConfig](https://cloud.tencent.com/document/api/607/35375#AgentConfig)
* [AmbientSound](https://cloud.tencent.com/document/api/607/35375#AmbientSound)
* [InvokeLLM](https://cloud.tencent.com/document/api/607/35375#InvokeLLM)
* [STTConfig](https://cloud.tencent.com/document/api/607/35375#STTConfig)
* [ServerPushText](https://cloud.tencent.com/document/api/607/35375#ServerPushText)
* [TurnDetection](https://cloud.tencent.com/document/api/607/35375#TurnDetection)
* [VoicePrint](https://cloud.tencent.com/document/api/607/35375#VoicePrint)
* [VoicePrintInfo](https://cloud.tencent.com/document/api/607/35375#VoicePrintInfo)



## 腾讯混元大模型(hunyuan) 版本：2023-09-01

### 第 46 次发布

发布时间：2025-09-11 01:52:34

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [WebSearchOptions](https://cloud.tencent.com/document/api/1729/101838#WebSearchOptions)

	* 新增成员：EnableImage, EnableMusic




## iOA 零信任安全管理系统(ioa) 版本：2022-06-01

### 第 25 次发布

发布时间：2025-09-11 01:54:03

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeAggrSoftCategorySoftList](https://cloud.tencent.com/document/api/1092/122707)

	* 新增入参：Condition

* [DescribeAggrSoftDeviceList](https://cloud.tencent.com/document/api/1092/122976)

	* 新增入参：Condition

* [DescribeSoftwareInformation](https://cloud.tencent.com/document/api/1092/119493)

	* 新增入参：OsType

* [ExportSoftwareInformationList](https://cloud.tencent.com/document/api/1092/122974)

	* 新增入参：OsType




## 物联网开发平台(iotexplorer) 版本：2019-04-23

### 第 116 次发布

发布时间：2025-09-11 01:54:59

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [BatchUpdateFirmware](https://cloud.tencent.com/document/api/1081/123152)
* [CreateOtaModule](https://cloud.tencent.com/document/api/1081/123151)
* [DeleteOtaModule](https://cloud.tencent.com/document/api/1081/123150)
* [DescribeFirmwareTaskDevices](https://cloud.tencent.com/document/api/1081/123146)
* [DescribeFirmwareTasks](https://cloud.tencent.com/document/api/1081/123145)
* [ListOtaModules](https://cloud.tencent.com/document/api/1081/123149)
* [ListProductOtaModules](https://cloud.tencent.com/document/api/1081/123148)
* [UpdateOtaModule](https://cloud.tencent.com/document/api/1081/123147)

新增数据结构：

* [DeviceUpdateStatus](https://cloud.tencent.com/document/api/1081/34988#DeviceUpdateStatus)
* [FirmwareTaskInfo](https://cloud.tencent.com/document/api/1081/34988#FirmwareTaskInfo)
* [OtaModuleInfo](https://cloud.tencent.com/document/api/1081/34988#OtaModuleInfo)



## 轻量应用服务器(lighthouse) 版本：2020-03-24

### 第 73 次发布

发布时间：2025-09-11 02:02:26

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [RebootInstances](https://cloud.tencent.com/document/api/1207/47572)

	* 新增入参：StopType

* [StopInstances](https://cloud.tencent.com/document/api/1207/47569)

	* 新增入参：StopType




## 知识引擎原子能力(lkeap) 版本：2024-05-22

### 第 16 次发布

发布时间：2025-09-11 02:07:02

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [GetEmbedding](https://cloud.tencent.com/document/api/1772/115343)

	* 新增入参：TextType, Instruction




## 消息队列 MQTT 版(mqtt) 版本：2024-05-16

### 第 20 次发布

发布时间：2025-09-11 02:17:53

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeInstance](https://cloud.tencent.com/document/api/1778/111030)

	* 新增出参：TopicPrefixSlashLimit, MessageRate

* [DescribeMessageDetails](https://cloud.tencent.com/document/api/1778/120249)

	* 新增出参：ContentType, PayloadFormatIndicator, MessageExpiryInterval, ResponseTopic, CorrelationData, SubscriptionIdentifier

* [ModifyInstance](https://cloud.tencent.com/document/api/1778/116039)

	* 新增入参：MessageRate




## 文字识别(ocr) 版本：2018-11-19

### 第 214 次发布

发布时间：2025-09-11 02:21:15

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeExtractDocAgentJob](https://cloud.tencent.com/document/api/866/123153)

### 第 213 次发布

发布时间：2025-09-10 19:17:52

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [SubmitExtractDocAgentJob](https://cloud.tencent.com/document/api/866/123134)

新增数据结构：

* [ItemNames](https://cloud.tencent.com/document/api/866/33527#ItemNames)



