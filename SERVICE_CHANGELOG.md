# Release v1.0.1182

## 应用性能监控(apm) 版本：2021-06-22

### 第 37 次发布

发布时间：2025-06-09 01:08:56

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeApmInstances](https://cloud.tencent.com/document/api/1463/65103)

	* 新增入参：InstanceId




## 云硬盘(cbs) 版本：2017-03-12

### 第 70 次发布

发布时间：2025-06-09 01:12:50

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**预下线接口**：</font>

* GetSnapOverview



## 弹性 MapReduce(emr) 版本：2019-01-03

### 第 107 次发布

发布时间：2025-06-09 01:32:27

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyPodNum](https://cloud.tencent.com/document/api/589/113700)

	* <font color="#dd0000">**修改入参**：</font>ServiceType, PodNum

* [ScaleOutInstance](https://cloud.tencent.com/document/api/589/34264)

	* 新增入参：ComputeResourceAdvanceParams


新增数据结构：

* [AutoScaleGroupAdvanceAttrs](https://cloud.tencent.com/document/api/589/33981#AutoScaleGroupAdvanceAttrs)
* [ComputeResourceAdvanceParams](https://cloud.tencent.com/document/api/589/33981#ComputeResourceAdvanceParams)
* [Taint](https://cloud.tencent.com/document/api/589/33981#Taint)
* [TkeLabel](https://cloud.tencent.com/document/api/589/33981#TkeLabel)

修改数据结构：

* [AutoScaleResourceConf](https://cloud.tencent.com/document/api/589/33981#AutoScaleResourceConf)

	* 新增成员：ExtraAdvanceAttrs

* [TimeAutoScaleStrategy](https://cloud.tencent.com/document/api/589/33981#TimeAutoScaleStrategy)

	* 新增成员：GraceDownLabel




## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 223 次发布

发布时间：2025-06-09 01:33:47

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateContractDiffTaskWebUrl](https://cloud.tencent.com/document/api/1323/119177)
* [DescribeContractDiffTaskWebUrl](https://cloud.tencent.com/document/api/1323/119176)



## 云游戏(gs) 版本：2019-11-18

### 第 39 次发布

发布时间：2025-06-09 01:36:49

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeAndroidInstancesAppBlacklist](https://cloud.tencent.com/document/api/1162/119182)
* [DescribeAndroidInstancesByApps](https://cloud.tencent.com/document/api/1162/119184)
* [ImportAndroidInstanceImage](https://cloud.tencent.com/document/api/1162/119178)
* [ModifyAndroidInstancesAppBlacklist](https://cloud.tencent.com/document/api/1162/119181)
* [ModifyAndroidInstancesResources](https://cloud.tencent.com/document/api/1162/119183)
* [SetAndroidInstancesBGAppKeepAlive](https://cloud.tencent.com/document/api/1162/119180)
* [SetAndroidInstancesFGAppKeepAlive](https://cloud.tencent.com/document/api/1162/119179)

新增数据结构：

* [AndroidInstanceAppBlacklist](https://cloud.tencent.com/document/api/1162/40743#AndroidInstanceAppBlacklist)



## 文字识别(ocr) 版本：2018-11-19

### 第 192 次发布

发布时间：2025-06-09 01:50:58

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [RecognizeGeneralInvoice](https://cloud.tencent.com/document/api/866/90802)

	* 新增入参：EnableSeal


新增数据结构：

* [InvoiceSealInfo](https://cloud.tencent.com/document/api/866/33527#InvoiceSealInfo)

修改数据结构：

* [InvoiceItem](https://cloud.tencent.com/document/api/866/33527#InvoiceItem)

	* 新增成员：InvoiceSealInfo




## 弹性微服务(tem) 版本：2021-07-01

### 第 49 次发布

发布时间：2025-06-09 02:04:00

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [ModifyGatewayIngress](https://cloud.tencent.com/document/api/1371/119185)



## 弹性微服务(tem) 版本：2020-12-21



## 边缘安全加速平台(teo) 版本：2022-09-01

### 第 105 次发布

发布时间：2025-06-09 02:04:31

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [AdaptiveFrequencyControl](https://cloud.tencent.com/document/api/1552/80721#AdaptiveFrequencyControl)
* [BandwidthAbuseDefense](https://cloud.tencent.com/document/api/1552/80721#BandwidthAbuseDefense)
* [ChallengeActionParameters](https://cloud.tencent.com/document/api/1552/80721#ChallengeActionParameters)
* [ClientFiltering](https://cloud.tencent.com/document/api/1552/80721#ClientFiltering)
* [DenyActionParameters](https://cloud.tencent.com/document/api/1552/80721#DenyActionParameters)
* [ExceptionRule](https://cloud.tencent.com/document/api/1552/80721#ExceptionRule)
* [ExceptionRules](https://cloud.tencent.com/document/api/1552/80721#ExceptionRules)
* [HttpDDoSProtection](https://cloud.tencent.com/document/api/1552/80721#HttpDDoSProtection)
* [MinimalRequestBodyTransferRate](https://cloud.tencent.com/document/api/1552/80721#MinimalRequestBodyTransferRate)
* [RateLimitingRule](https://cloud.tencent.com/document/api/1552/80721#RateLimitingRule)
* [RateLimitingRules](https://cloud.tencent.com/document/api/1552/80721#RateLimitingRules)
* [RequestBodyTransferTimeout](https://cloud.tencent.com/document/api/1552/80721#RequestBodyTransferTimeout)
* [RequestFieldsForException](https://cloud.tencent.com/document/api/1552/80721#RequestFieldsForException)
* [SlowAttackDefense](https://cloud.tencent.com/document/api/1552/80721#SlowAttackDefense)

修改数据结构：

* [SecurityAction](https://cloud.tencent.com/document/api/1552/80721#SecurityAction)

	* 新增成员：DenyActionParameters, ChallengeActionParameters

* [SecurityPolicy](https://cloud.tencent.com/document/api/1552/80721#SecurityPolicy)

	* 新增成员：HttpDDoSProtection, RateLimitingRules, ExceptionRules




## 边缘安全加速平台(teo) 版本：2022-01-06



## 云点播(vod) 版本：2024-07-18



## 云点播(vod) 版本：2018-07-17

### 第 200 次发布

发布时间：2025-06-09 02:12:41

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [MediaContentReviewSegmentItem](https://cloud.tencent.com/document/api/266/31773#MediaContentReviewSegmentItem)




## 数据开发治理平台 WeData(wedata) 版本：2021-08-20

### 第 154 次发布

发布时间：2025-06-09 02:17:59

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeExecutorGroupMetric](https://cloud.tencent.com/document/api/1267/119186)

新增数据结构：

* [ExecutorResourceGroupInfo](https://cloud.tencent.com/document/api/1267/76336#ExecutorResourceGroupInfo)
* [ExecutorResourcePackageExtInfo](https://cloud.tencent.com/document/api/1267/76336#ExecutorResourcePackageExtInfo)
* [ExecutorResourcePackageInfo](https://cloud.tencent.com/document/api/1267/76336#ExecutorResourcePackageInfo)
* [ExecutorResourcePackageUsageInfo](https://cloud.tencent.com/document/api/1267/76336#ExecutorResourcePackageUsageInfo)
* [ExecutorUsageTrendInfo](https://cloud.tencent.com/document/api/1267/76336#ExecutorUsageTrendInfo)
* [MQPackageVO](https://cloud.tencent.com/document/api/1267/76336#MQPackageVO)



