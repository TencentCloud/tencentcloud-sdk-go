# Release v1.3.13

## 运维安全中心（堡垒机）(bh) 版本：2023-04-18

### 第 21 次发布

发布时间：2025-12-17 01:17:16

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeDepartments](https://cloud.tencent.com/document/api/1025/126463)

修改接口：

* [CreateUserDirectory](https://cloud.tencent.com/document/api/1025/125473)

	* <font color="#dd0000">**修改入参**：</font>UserCount


新增数据结构：

* [Departments](https://cloud.tencent.com/document/api/1025/74416#Departments)



## 费用中心(billing) 版本：2018-07-09

### 第 78 次发布

发布时间：2025-12-17 01:18:41

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [CostDetail](https://cloud.tencent.com/document/api/555/19183#CostDetail)

	* 新增成员：Tags




## 本地专用集群(cdc) 版本：2020-12-14

### 第 14 次发布

发布时间：2025-12-17 01:24:18

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeDedicatedClusterOverview](https://cloud.tencent.com/document/api/1346/73760)

	* 新增出参：HostAbnormalCount




## 负载均衡(clb) 版本：2018-03-17

### 第 142 次发布

发布时间：2025-12-15 01:25:20

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeTargetGroupInstanceStatus](https://cloud.tencent.com/document/api/214/126441)

新增数据结构：

* [TargetGroupInstanceStatus](https://cloud.tencent.com/document/api/214/30694#TargetGroupInstanceStatus)



## 日志服务(cls) 版本：2020-10-16

### 第 143 次发布

发布时间：2025-12-16 01:26:54

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateAlarmNotice](https://cloud.tencent.com/document/api/614/56465)

	* 新增入参：CallbackPrioritize

* [CreateConsumer](https://cloud.tencent.com/document/api/614/66228)

	* 新增入参：RoleArn, ExternalId, AdvancedConfig

* [CreateIndex](https://cloud.tencent.com/document/api/614/56445)

	* 新增入参：CoverageField

* [CreateShipper](https://cloud.tencent.com/document/api/614/58747)

	* 新增入参：RoleArn, ExternalId

* [DescribeAlarmNotices](https://cloud.tencent.com/document/api/614/56462)

	* 新增入参：HasAlarmShieldCount

* [DescribeIndex](https://cloud.tencent.com/document/api/614/56443)

	* 新增出参：CoverageField

* [ModifyAlarmNotice](https://cloud.tencent.com/document/api/614/56458)

	* 新增入参：CallbackPrioritize

* [ModifyConsumer](https://cloud.tencent.com/document/api/614/66225)

	* 新增入参：RoleArn, ExternalId, AdvancedConfig

* [ModifyIndex](https://cloud.tencent.com/document/api/614/56442)

	* 新增入参：CoverageField

* [ModifyShipper](https://cloud.tencent.com/document/api/614/58743)

	* 新增入参：RoleArn, ExternalId

* [ModifyTopic](https://cloud.tencent.com/document/api/614/56453)

	* 新增入参：StorageType, Encryption, IsSourceFrom


新增数据结构：

* [AdvancedConsumerConfiguration](https://cloud.tencent.com/document/api/614/56471#AdvancedConsumerConfiguration)
* [AlarmShieldCount](https://cloud.tencent.com/document/api/614/56471#AlarmShieldCount)

修改数据结构：

* [AlarmNotice](https://cloud.tencent.com/document/api/614/56471#AlarmNotice)

	* 新增成员：DeliverStatus, DeliverFlag, AlarmShieldCount, CallbackPrioritize

* [LogsetInfo](https://cloud.tencent.com/document/api/614/56471#LogsetInfo)

	* 新增成员：AssumerUin, MetricTopicCount

* [ShipperInfo](https://cloud.tencent.com/document/api/614/56471#ShipperInfo)

	* 新增成员：RoleArn, ExternalId, TaskStatus

* [TopicInfo](https://cloud.tencent.com/document/api/614/56471#TopicInfo)

	* 新增成员：AssumerUin, RoleName, IsSourceFrom




## Elasticsearch Service(es) 版本：2025-01-01



## Elasticsearch Service(es) 版本：2018-04-16

### 第 91 次发布

发布时间：2025-12-16 01:43:28

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateClusterSnapshot](https://cloud.tencent.com/document/api/845/112044)

	* 新增入参：EsRepositoryType, UserEsRepository, StorageDuration, CosRetention, RetainUntilDate, RetentionGraceTime, RemoteCos, RemoteCosRegion


修改数据结构：

* [CosBackup](https://cloud.tencent.com/document/api/845/30634#CosBackup)

	* 新增成员：SnapshotName, PaasEsRepository, CosRetention, RetainUntilDate, RetentionGraceTime, RemoteCos, RemoteCosRegion, StrategyName, Indices, CreateTime

* [Snapshots](https://cloud.tencent.com/document/api/845/30634#Snapshots)

	* 新增成员：EsRepositoryType, PaasEsRepository, UserEsRepository, StorageDuration, AutoBackupInterval, CosRetention, RetainUntilDate, RetentionGraceTime, IsLocked, RemoteCos, RemoteCosRegion, CosEncryption, KmsKey, StrategyName




## 事件中心(evt) 版本：2025-02-17

### 第 3 次发布

发布时间：2025-12-15 11:22:42

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CompleteApproval](https://cloud.tencent.com/document/api/1808/126452)

### 第 2 次发布

发布时间：2025-12-15 01:47:01

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateRoleUser](https://cloud.tencent.com/document/api/1808/126075)

	* 新增入参：TencentUin




## 云游戏(gs) 版本：2019-11-18

### 第 61 次发布

发布时间：2025-12-17 01:53:12

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateAndroidInstanceAcceleratorToken](https://cloud.tencent.com/document/api/1162/125484)

	* <font color="#dd0000">**修改入参**：</font>UserIP


### 第 60 次发布

发布时间：2025-12-16 01:48:49

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateAndroidInstanceAcceleratorToken](https://cloud.tencent.com/document/api/1162/125484)

	* 新增出参：Token




## 智能全局流量管理(igtm) 版本：2023-10-24

### 第 4 次发布

发布时间：2025-12-17 01:55:21

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreatePackageAndPay](https://cloud.tencent.com/document/api/1551/126465)
* [ModifyPackageAutoRenew](https://cloud.tencent.com/document/api/1551/126464)



## 腾讯云智能体开发平台(lke) 版本：2023-11-30

### 第 72 次发布

发布时间：2025-12-16 02:00:55

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [GenerateQA](https://cloud.tencent.com/document/api/1759/105069)

	* <font color="#dd0000">**修改入参**：</font>DocBizIds

* [GetWsToken](https://cloud.tencent.com/document/api/1759/105089)

	* 新增出参：VisionModelInputLimit

* [ListUnsatisfiedReply](https://cloud.tencent.com/document/api/1759/105026)

	* 新增入参：HandlingStatuses


新增数据结构：

* [AppModelDetailInfo](https://cloud.tencent.com/document/api/1759/105104#AppModelDetailInfo)
* [NL2SQLModelConfig](https://cloud.tencent.com/document/api/1759/105104#NL2SQLModelConfig)

修改数据结构：

* [AttrLabelDetail](https://cloud.tencent.com/document/api/1759/105104#AttrLabelDetail)

* [Filters](https://cloud.tencent.com/document/api/1759/105104#Filters)

	* 新增成员：HandlingStatuses

* [SearchStrategy](https://cloud.tencent.com/document/api/1759/105104#SearchStrategy)

	* 新增成员：NatureLanguageToSqlModelConfig




## 知识引擎原子能力(lkeap) 版本：2024-05-22

### 第 19 次发布

发布时间：2025-12-17 14:13:05

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CancelTask](https://cloud.tencent.com/document/api/1772/126467)



## 流计算 Oceanus(oceanus) 版本：2019-04-22

### 第 77 次发布

发布时间：2025-12-17 02:21:44

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [Cluster](https://cloud.tencent.com/document/api/849/52010#Cluster)

	* 新增成员：CdcId




## 文字识别(ocr) 版本：2018-11-19

### 第 226 次发布

发布时间：2025-12-15 02:18:53

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [ExtractDocAgent](https://cloud.tencent.com/document/api/866/126442)



## 云开发 CloudBase(tcb) 版本：2018-06-08

### 第 115 次发布

发布时间：2025-12-16 02:42:42

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [ActivityRecordItem](https://cloud.tencent.com/document/api/876/34822#ActivityRecordItem)

	* 新增成员：EnvId, ActivityName, CreateTime




## 边缘安全加速平台(teo) 版本：2022-09-01

### 第 128 次发布

发布时间：2025-12-16 02:57:47

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeTopL7AnalysisData](https://cloud.tencent.com/document/api/1552/80646)

	* <font color="#dd0000">**修改入参**：</font>ZoneIds




## 边缘安全加速平台(teo) 版本：2022-01-06



## 容器服务(tke) 版本：2022-05-01

### 第 19 次发布

发布时间：2025-12-16 03:09:01

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [CreateNativeNodePoolParam](https://cloud.tencent.com/document/api/457/103206#CreateNativeNodePoolParam)

	* 新增成员：Password

* [UpdateNativeNodePoolParam](https://cloud.tencent.com/document/api/457/103206#UpdateNativeNodePoolParam)

	* 新增成员：Password




## 容器服务(tke) 版本：2018-05-25

### 第 212 次发布

发布时间：2025-12-15 03:05:39

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeControlPlaneLogs](https://cloud.tencent.com/document/api/457/126445)
* [DisableControlPlaneLogs](https://cloud.tencent.com/document/api/457/126444)
* [EnableControlPlaneLogs](https://cloud.tencent.com/document/api/457/126443)

新增数据结构：

* [ComponentLogConfig](https://cloud.tencent.com/document/api/457/31866#ComponentLogConfig)



## 云点播(vod) 版本：2024-07-18



## 云点播(vod) 版本：2018-07-17

### 第 212 次发布

发布时间：2025-12-15 03:24:22

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeAigcUsageData](https://cloud.tencent.com/document/api/266/126446)

新增数据结构：

* [AigcUsageDataItem](https://cloud.tencent.com/document/api/266/31773#AigcUsageDataItem)



## 私有网络(vpc) 版本：2017-03-12

### 第 282 次发布

发布时间：2025-12-17 03:32:37

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeRoutePolicies](https://cloud.tencent.com/document/api/215/126466)

修改接口：

* [CloneSecurityGroup](https://cloud.tencent.com/document/api/215/51025)

	* 新增入参：IsCloneTags

* [CreateVpcEndPointService](https://cloud.tencent.com/document/api/215/54684)

* [DescribeVpcEndPointServiceWhiteList](https://cloud.tencent.com/document/api/215/54677)

* [ReplaceHighPriorityRouteTableAssociation](https://cloud.tencent.com/document/api/215/111312)

	* <font color="#dd0000">**修改入参**：</font>SubnetId

* [ReplaceRouteTableAssociation](https://cloud.tencent.com/document/api/215/15767)

	* <font color="#dd0000">**修改入参**：</font>SubnetId

* [ReplaceSecurityGroupPolicies](https://cloud.tencent.com/document/api/215/88451)

	* 新增入参：UpdateType




## 数据开发治理平台 WeData(wedata) 版本：2025-08-06



## 数据开发治理平台 WeData(wedata) 版本：2021-08-20

### 第 174 次发布

发布时间：2025-12-16 03:39:05

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeTaskInstancesStatus](https://cloud.tencent.com/document/api/1267/126460)

新增数据结构：

* [DescribeTaskInstancesStatusDto](https://cloud.tencent.com/document/api/1267/76336#DescribeTaskInstancesStatusDto)
* [ParamGetTaskInstancesStatusInfoResponseInstance](https://cloud.tencent.com/document/api/1267/76336#ParamGetTaskInstancesStatusInfoResponseInstance)



