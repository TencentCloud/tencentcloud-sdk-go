# Release v1.0.1161

## 应用性能监控(apm) 版本：2021-06-22

### 第 34 次发布

发布时间：2025-05-09 01:09:12

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyApmInstance](https://cloud.tencent.com/document/api/1463/89002)

	* 新增入参：IsRemoteCommandExecutionAnalysis, IsMemoryHijackingAnalysis


修改数据结构：

* [ApmInstanceDetail](https://cloud.tencent.com/document/api/1463/64927#ApmInstanceDetail)

	* 新增成员：IsRemoteCommandExecutionAnalysis, IsMemoryHijackingAnalysis




## 语音识别(asr) 版本：2019-06-14

### 第 42 次发布

发布时间：2025-05-09 01:09:46

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [GetAsrVocabList](https://cloud.tencent.com/document/api/1093/41484)


修改数据结构：

* [Model](https://cloud.tencent.com/document/api/1093/37824#Model)




## 负载均衡(clb) 版本：2018-03-17

### 第 128 次发布

发布时间：2025-05-09 01:18:54

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateTargetGroup](https://cloud.tencent.com/document/api/214/40559)

	* 新增入参：KeepaliveEnable, SessionExpireTime

* [ModifyTargetGroupAttribute](https://cloud.tencent.com/document/api/214/40552)

	* 新增入参：KeepaliveEnable, SessionExpireTime




## TDSQL-C MySQL 版(cynosdb) 版本：2019-01-07

### 第 131 次发布

发布时间：2025-05-09 01:25:13

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CloseSSL](https://cloud.tencent.com/document/api/1003/118379)



## 数据安全治理中心(dsgc) 版本：2019-07-23

### 第 29 次发布

发布时间：2025-05-09 01:29:45

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeDSPADiscoveryTasks](https://cloud.tencent.com/document/api/1087/118380)

新增数据结构：

* [DspaDiscoveryTask](https://cloud.tencent.com/document/api/1087/96844#DspaDiscoveryTask)



## 文字识别(ocr) 版本：2018-11-19

### 第 189 次发布

发布时间：2025-05-08 14:12:26

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [ResultList](https://cloud.tencent.com/document/api/866/33527#ResultList)

	* 新增成员：Parse




## 消息队列 TDMQ(tdmq) 版本：2020-02-17

### 第 148 次发布

发布时间：2025-05-09 02:01:57

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateRocketMQCluster](https://cloud.tencent.com/document/api/1179/63429)

	* 新增入参：TagList

* [CreateTopic](https://cloud.tencent.com/document/api/1179/46088)

	* 新增入参：IsolateConsumerEnable, AckTimeOut

* [DescribeMqMsgTrace](https://cloud.tencent.com/document/api/1179/104106)

	* 新增入参：ProduceTime

* [ModifyTopic](https://cloud.tencent.com/document/api/1179/46085)

	* 新增入参：IsolateConsumerEnable, AckTimeOut


修改数据结构：

* [RocketMQClusterConfig](https://cloud.tencent.com/document/api/1179/46089#RocketMQClusterConfig)

	* 新增成员：MaxRoleNum, MaxTpsLimit

* [RocketMQClusterInfo](https://cloud.tencent.com/document/api/1179/46089#RocketMQClusterInfo)

	* 新增成员：IsFrozen

* [Topic](https://cloud.tencent.com/document/api/1179/46089#Topic)

	* 新增成员：IsolateConsumerEnable, AckTimeOut




## 消息队列 RocketMQ 版(trocket) 版本：2023-03-08

### 第 37 次发布

发布时间：2025-05-09 02:07:19

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateConsumerGroup](https://cloud.tencent.com/document/api/1493/97943)

	* 新增入参：TagList

* [CreateTopic](https://cloud.tencent.com/document/api/1493/97947)

	* 新增入参：TagList


新增数据结构：

* [ClientSubscriptionInfo](https://cloud.tencent.com/document/api/1493/96031#ClientSubscriptionInfo)

修改数据结构：

* [FusionInstanceItem](https://cloud.tencent.com/document/api/1493/96031#FusionInstanceItem)

	* 新增成员：ZoneIds

* [InstanceItemExtraInfo](https://cloud.tencent.com/document/api/1493/96031#InstanceItemExtraInfo)

	* 新增成员：IsFrozen

* [SubscriptionData](https://cloud.tencent.com/document/api/1493/96031#SubscriptionData)

	* 新增成员：ClientSubscriptionInfos




## 私有网络(vpc) 版本：2017-03-12

### 第 257 次发布

发布时间：2025-05-09 02:13:10

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [InquiryPriceAllocateAddresses](https://cloud.tencent.com/document/api/215/114855)

	* <font color="#dd0000">**修改入参**：</font>InternetChargeType


修改数据结构：

* [ServiceTemplateSpecification](https://cloud.tencent.com/document/api/215/15824#ServiceTemplateSpecification)

	* <font color="#dd0000">**修改成员**：</font>ServiceId, ServiceGroupId




## Web 应用防火墙(waf) 版本：2018-01-25

### 第 116 次发布

发布时间：2025-05-09 02:15:30

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [ModifyInstanceAttackLogPost](https://cloud.tencent.com/document/api/627/118381)



## 数据开发治理平台 WeData(wedata) 版本：2021-08-20

### 第 147 次发布

发布时间：2025-05-09 02:16:48

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateTaskNew](https://cloud.tencent.com/document/api/1267/118386)
* [ModifyTaskLinksDs](https://cloud.tencent.com/document/api/1267/118385)
* [RegisterDsEvent](https://cloud.tencent.com/document/api/1267/118384)
* [RenewWorkflowOwnerDs](https://cloud.tencent.com/document/api/1267/118383)
* [UpdateWorkflowInfo](https://cloud.tencent.com/document/api/1267/118382)

新增数据结构：

* [EventDsDto](https://cloud.tencent.com/document/api/1267/76336#EventDsDto)
* [EventListenerDTO](https://cloud.tencent.com/document/api/1267/76336#EventListenerDTO)
* [EventListenerTaskInfo](https://cloud.tencent.com/document/api/1267/76336#EventListenerTaskInfo)



