# Release v1.2.1

## 腾讯混元生3D(ai3d) 版本：2025-05-13

### 第 5 次发布

发布时间：2025-11-14 01:07:42

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除接口**：</font>

* QueryHunyuanTo3DJob
* SubmitHunyuanTo3DJob



## 运维安全中心（堡垒机）(bh) 版本：2023-04-18

### 第 18 次发布

发布时间：2025-11-14 10:47:44

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeAssetSyncFlag](https://cloud.tencent.com/document/api/1025/125051)
* [DescribeSecuritySetting](https://cloud.tencent.com/document/api/1025/125050)
* [DisableExternalAccess](https://cloud.tencent.com/document/api/1025/125055)
* [DisableIntranetAccess](https://cloud.tencent.com/document/api/1025/125054)
* [EnableExternalAccess](https://cloud.tencent.com/document/api/1025/125053)
* [EnableIntranetAccess](https://cloud.tencent.com/document/api/1025/125052)
* [ModifyAccessWhiteListAutoStatus](https://cloud.tencent.com/document/api/1025/125046)
* [ModifyAccessWhiteListRule](https://cloud.tencent.com/document/api/1025/125045)
* [ModifyAccessWhiteListStatus](https://cloud.tencent.com/document/api/1025/125044)
* [ModifyAssetSyncFlag](https://cloud.tencent.com/document/api/1025/125049)
* [ModifyAuthModeSetting](https://cloud.tencent.com/document/api/1025/125048)
* [ModifyReconnectionSetting](https://cloud.tencent.com/document/api/1025/125047)

修改接口：

* [DescribeAcls](https://cloud.tencent.com/document/api/1025/74409)

	* 新增入参：StatusSet

* [DescribeCmdTemplates](https://cloud.tencent.com/document/api/1025/86962)

	* 新增入参：TypeSet

* [DescribeDevices](https://cloud.tencent.com/document/api/1025/74415)

	* 新增入参：AccountIdSet, ProviderTypeSet, CloudDeviceStatusSet

* [DescribeLoginEvent](https://cloud.tencent.com/document/api/1025/90039)

	* 新增入参：EntrySet, ResultSet

* [DescribeOperationEvent](https://cloud.tencent.com/document/api/1025/90038)

	* 新增入参：KindSet, ResultSet

* [ImportExternalDevice](https://cloud.tencent.com/document/api/1025/86967)

	* 新增入参：AccountId

* [SearchFileBySid](https://cloud.tencent.com/document/api/1025/112712)

	* 新增入参：AuditActionSet

* [SearchSession](https://cloud.tencent.com/document/api/1025/112711)

	* 新增入参：StatusSet, DeviceKindSet


新增数据结构：

* [AssetSyncFlags](https://cloud.tencent.com/document/api/1025/74416#AssetSyncFlags)

修改数据结构：

* [Acl](https://cloud.tencent.com/document/api/1025/74416#Acl)

	* 新增成员：AclType, TicketId, TicketName

* [Device](https://cloud.tencent.com/document/api/1025/74416#Device)

	* 新增成员：ApName, CloudAccountId, CloudAccountName, ProviderType, ProviderName, SyncCloudDeviceStatus

* [ExternalDevice](https://cloud.tencent.com/document/api/1025/74416#ExternalDevice)

	* 新增成员：InstanceId, ApCode, ApName, VpcId, SubnetId, PublicIp

* [Resource](https://cloud.tencent.com/document/api/1025/74416#Resource)

	* 新增成员：DomainName




## 费用中心(billing) 版本：2018-07-09

### 第 77 次发布

发布时间：2025-11-14 01:09:30

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeBillSummary](https://cloud.tencent.com/document/api/555/93162)

	* 新增入参：OperateUin




## 日志服务(cls) 版本：2020-10-16

### 第 140 次发布

发布时间：2025-11-14 01:14:46

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CheckRechargeKafkaServer](https://cloud.tencent.com/document/api/614/94449)

	* 新增入参：UserKafkaMeta

* [CreateConfig](https://cloud.tencent.com/document/api/614/58620)

	* 新增入参：InputType

* [CreateKafkaRecharge](https://cloud.tencent.com/document/api/614/94448)

	* 新增入参：UserKafkaMeta

* [ModifyConfig](https://cloud.tencent.com/document/api/614/58614)

	* 新增入参：InputType

* [ModifyKafkaRecharge](https://cloud.tencent.com/document/api/614/94445)

	* 新增入参：UserKafkaMeta

* [PreviewKafkaRecharge](https://cloud.tencent.com/document/api/614/94444)

	* 新增入参：UserKafkaMeta


新增数据结构：

* [UserKafkaMeta](https://cloud.tencent.com/document/api/614/56471#UserKafkaMeta)

修改数据结构：

* [ConfigInfo](https://cloud.tencent.com/document/api/614/56471#ConfigInfo)

	* 新增成员：InputType

* [KafkaRechargeInfo](https://cloud.tencent.com/document/api/614/56471#KafkaRechargeInfo)

	* 新增成员：UserKafkaMeta

* [LogRechargeRuleInfo](https://cloud.tencent.com/document/api/614/56471#LogRechargeRuleInfo)

	* 新增成员：Delimiter




## 域名注册(domain) 版本：2018-08-08

### 第 43 次发布

发布时间：2025-11-14 01:19:10

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeBiddingSuccessfulDetail](https://cloud.tencent.com/document/api/242/106597)

	* 新增出参：ModifyOwnerEndTime


修改数据结构：

* [BiddingSuccessfulResult](https://cloud.tencent.com/document/api/242/38895#BiddingSuccessfulResult)

	* 新增成员：ModifyOwnerEndTime




## 腾讯混元大模型(hunyuan) 版本：2023-09-01

### 第 47 次发布

发布时间：2025-11-14 01:23:21

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [VideoFrames](https://cloud.tencent.com/document/api/1729/101838#VideoFrames)
* [VideoUrl](https://cloud.tencent.com/document/api/1729/101838#VideoUrl)

修改数据结构：

* [Content](https://cloud.tencent.com/document/api/1729/101838#Content)

	* 新增成员：VideoUrl, VideoFrames




## 腾讯云可观测平台(monitor) 版本：2023-06-16



## 腾讯云可观测平台(monitor) 版本：2018-07-24

### 第 144 次发布

发布时间：2025-11-14 01:33:41

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeAlarmHistories](https://cloud.tencent.com/document/api/248/48684)

	* 新增入参：ShieldStatus




## 边缘安全加速平台(teo) 版本：2022-09-01

### 第 125 次发布

发布时间：2025-11-14 02:08:13

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [OriginCertificateVerify](https://cloud.tencent.com/document/api/1552/80721#OriginCertificateVerify)

修改数据结构：

* [UpstreamCertInfo](https://cloud.tencent.com/document/api/1552/80721#UpstreamCertInfo)

	* 新增成员：UpstreamCertificateVerify




## 边缘安全加速平台(teo) 版本：2022-01-06



## 消息队列 RocketMQ 版(trocket) 版本：2023-03-08

### 第 46 次发布

发布时间：2025-11-14 02:17:57

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeConsumerGroupList](https://cloud.tencent.com/document/api/1493/101535)

	* 新增入参：TagFilters

* [DescribeInstance](https://cloud.tencent.com/document/api/1493/97866)

	* 新增出参：NodeCount, ZoneScheduledList

* [DescribeTopicList](https://cloud.tencent.com/document/api/1493/96030)

	* 新增入参：TagFilters


新增数据结构：

* [ZoneScheduledItem](https://cloud.tencent.com/document/api/1493/96031#ZoneScheduledItem)

修改数据结构：

* [ConsumeGroupItem](https://cloud.tencent.com/document/api/1493/96031#ConsumeGroupItem)

	* 新增成员：TagList

* [FusionInstanceItem](https://cloud.tencent.com/document/api/1493/96031#FusionInstanceItem)

	* 新增成员：CreateTime, ScaledTpsEnabled

* [TopicItem](https://cloud.tencent.com/document/api/1493/96031#TopicItem)

	* 新增成员：TagList




## TSF-Polaris&ZK&网关(tse) 版本：2020-12-07

### 第 102 次发布

发布时间：2025-11-14 02:21:12

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateOrModifyCloudNativeAPIGatewayIPRestriction](https://cloud.tencent.com/document/api/1364/125038)
* [DeleteCloudNativeAPIGatewayIPRestriction](https://cloud.tencent.com/document/api/1364/125037)
* [DescribeCloudNativeAPIGatewayIPRestriction](https://cloud.tencent.com/document/api/1364/125036)

新增数据结构：

* [DescribeKongIpRestrictionResult](https://cloud.tencent.com/document/api/1364/54942#DescribeKongIpRestrictionResult)

修改数据结构：

* [DescribeCloudNativeAPIGatewayResult](https://cloud.tencent.com/document/api/1364/54942#DescribeCloudNativeAPIGatewayResult)

	* 新增成员：AvailableVersions




## TSF-应用管理&Consul(tsf) 版本：2018-03-26

### 第 133 次发布

发布时间：2025-11-14 02:22:51

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [ModifyGroupLane](https://cloud.tencent.com/document/api/649/125039)



