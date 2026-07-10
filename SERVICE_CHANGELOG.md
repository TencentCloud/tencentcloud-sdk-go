# Release v1.3.131

## Agent 沙箱服务(ags) 版本：2025-09-20

### 第 16 次发布

发布时间：2026-07-10 01:08:34

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreatePreCacheImageTask](https://cloud.tencent.com/document/api/1814/127508)

	* 新增入参：TimeoutMinutes




## 应用型负载均衡(alb) 版本：2025-10-30

### 第 2 次发布

发布时间：2026-07-09 01:09:48

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [AddTargetsToTargetGroup](https://cloud.tencent.com/document/api/1822/133696)

	* <font color="#dd0000">**修改入参**：</font>Targets

* [CreateRules](https://cloud.tencent.com/document/api/1822/133686)

	* <font color="#dd0000">**修改入参**：</font>Rules

* [DeleteHealthCheckTemplates](https://cloud.tencent.com/document/api/1822/133736)

	* <font color="#dd0000">**修改入参**：</font>HealthCheckTemplateIds

* [ModifyLoadBalancerModificationProtection](https://cloud.tencent.com/document/api/1822/133725)

	* <font color="#dd0000">**修改入参**：</font>ModificationProtectionEnabled

* [ModifyTargetsInTargetGroup](https://cloud.tencent.com/document/api/1822/133689)

	* <font color="#dd0000">**修改入参**：</font>Targets

* [RemoveTargetsFromTargetGroup](https://cloud.tencent.com/document/api/1822/133688)

	* <font color="#dd0000">**修改入参**：</font>Targets




## 语音识别(asr) 版本：2019-06-14

### 第 48 次发布

发布时间：2026-07-09 01:17:43

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [VoicePrintGroupList](https://cloud.tencent.com/document/api/1093/133831)

新增数据结构：

* [VoicePrintGroupList](https://cloud.tencent.com/document/api/1093/37824#VoicePrintGroupList)



## 文件存储(cfs) 版本：2019-07-19

### 第 53 次发布

发布时间：2026-07-10 01:25:00

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [RunDataRetrievalTask](https://cloud.tencent.com/document/api/582/132383)

	* 新增入参：DataRetrievalId




## 日志服务(cls) 版本：2020-10-16

### 第 166 次发布

发布时间：2026-07-10 01:31:10

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateDlcDeliver](https://cloud.tencent.com/document/api/614/125886)

	* 新增入参：AutoCreateField, DlcFailHandle, DSLFilter

* [ModifyDlcDeliver](https://cloud.tencent.com/document/api/614/125883)

	* 新增入参：AutoCreateField, DlcFailHandle, DSLFilter




## TDSQL-C MySQL 版(cynosdb) 版本：2019-01-07

### 第 180 次发布

发布时间：2026-07-10 01:44:12

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [InquirePriceCreate](https://cloud.tencent.com/document/api/1003/77738)

	* 新增入参：StorageVersion, IsMultiAz

* [ModifyBinlogSaveDays](https://cloud.tencent.com/document/api/1003/92725)

	* 新增入参：BinlogCrossRegionSaveDays

* [ModifySnapBackupCrossRegionConfig](https://cloud.tencent.com/document/api/1003/127848)

	* 新增入参：CrossRegionSaveDays


新增数据结构：

* [BinlogRegionInfo](https://cloud.tencent.com/document/api/1003/48097#BinlogRegionInfo)

修改数据结构：

* [BackupConfigInfo](https://cloud.tencent.com/document/api/1003/48097#BackupConfigInfo)

	* 新增成员：CrossRegionSaveDays

* [BinlogConfigInfo](https://cloud.tencent.com/document/api/1003/48097#BinlogConfigInfo)

	* 新增成员：BinlogCrossRegionSaveDays

* [BinlogItem](https://cloud.tencent.com/document/api/1003/48097#BinlogItem)

	* 新增成员：ExistRegions

* [LogicBackupConfigInfo](https://cloud.tencent.com/document/api/1003/48097#LogicBackupConfigInfo)

	* 新增成员：LogicCrossRegionSaveDays




## 数据库智能管家 DBbrain(dbbrain) 版本：2021-05-27

### 第 58 次发布

发布时间：2026-07-09 02:21:01

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateIgnoreDiagRecord](https://cloud.tencent.com/document/api/1130/133835)



## 数据库智能管家 DBbrain(dbbrain) 版本：2019-10-16



## 数据湖计算 DLC(dlc) 版本：2021-01-25

### 第 162 次发布

发布时间：2026-07-09 02:25:43

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeDataEngineEvents](https://cloud.tencent.com/document/api/1342/101534)

	* 新增入参：StartTime, EndTime




## DNSPod(dnspod) 版本：2021-03-23

### 第 54 次发布

发布时间：2026-07-10 01:51:38

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateSubdomainValidateTXTValue](https://cloud.tencent.com/document/api/1427/113012)

	* 新增出参：SubDomain

* [DescribeDomainAnalytics](https://cloud.tencent.com/document/api/1427/75753)

	* 新增入参：DNSFormat

* [DescribeRecordList](https://cloud.tencent.com/document/api/1427/56166)

	* 新增入参：SubDomain

* [DescribeResolveCount](https://cloud.tencent.com/document/api/1427/121749)

	* 新增入参：DNSFormat

	* <font color="#dd0000">**修改入参**：</font>DnsFormat

* [DescribeSubdomainAnalytics](https://cloud.tencent.com/document/api/1427/75752)

	* 新增入参：SubDomain, DNSFormat

	* <font color="#dd0000">**修改入参**：</font>Subdomain

* [DescribeVasList](https://cloud.tencent.com/document/api/1427/119490)

	* 新增出参：VASList

* [ModifyDynamicDNS](https://cloud.tencent.com/document/api/1427/56158)

	* 新增入参：TTL


修改数据结构：

* [DomainAnalyticsInfo](https://cloud.tencent.com/document/api/1427/56185#DomainAnalyticsInfo)

	* 新增成员：DNSFormat, DNSTotal

* [DomainInfo](https://cloud.tencent.com/document/api/1427/56185#DomainInfo)

	* 新增成员：DNSStatus, CNAMESpeedup

* [ResolveCountInfo](https://cloud.tencent.com/document/api/1427/56185#ResolveCountInfo)

	* 新增成员：DNSTotal, DNSFormat

* [SubdomainAnalyticsInfo](https://cloud.tencent.com/document/api/1427/56185#SubdomainAnalyticsInfo)

	* 新增成员：DNSFormat, DNSTotal, SubDomain




## 数据传输服务(dts) 版本：2021-12-06

### 第 57 次发布

发布时间：2026-07-09 02:35:20

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateSubscribe](https://cloud.tencent.com/document/api/571/102950)

	* 新增入参：SubscribeVersion




## 数据传输服务(dts) 版本：2018-03-30



## 弹性 MapReduce(emr) 版本：2019-01-03

### 第 149 次发布

发布时间：2026-07-10 01:57:59

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [Resource](https://cloud.tencent.com/document/api/589/33981#Resource)

	* 新增成员：HCCHpcClusterId




## 全球加速(ga2) 版本：2025-01-15

### 第 9 次发布

发布时间：2026-07-10 02:04:33

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeAccelerateAreas](https://cloud.tencent.com/document/api/1817/130165)

	* 新增入参：Filters




## 云直播CSS(live) 版本：2018-08-01

### 第 181 次发布

发布时间：2026-07-09 03:23:56

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateAuditKeywords](https://cloud.tencent.com/document/api/267/121248)

	* 新增出参：Keywords




## 腾讯云可观测平台(monitor) 版本：2023-06-16



## 腾讯云可观测平台(monitor) 版本：2018-07-24

### 第 162 次发布

发布时间：2026-07-09 03:33:51

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [EnablePredefinedPolicies](https://cloud.tencent.com/document/api/248/133836)



## 文字识别(ocr) 版本：2018-11-19

### 第 257 次发布

发布时间：2026-07-10 02:32:49

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CropEnhanceImageOCR](https://cloud.tencent.com/document/api/866/133908)
* [EraseHandwrittenImageOCR](https://cloud.tencent.com/document/api/866/133907)



## 云开发 CloudBase(tcb) 版本：2018-06-08

### 第 152 次发布

发布时间：2026-07-09 04:12:10

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [PreviewPGUserMigrations](https://cloud.tencent.com/document/api/876/132260)

	* 新增入参：IncludeAll

* [PushPGUserMigrations](https://cloud.tencent.com/document/api/876/132259)

	* 新增入参：IncludeAll




## 消息队列 TDMQ(tdmq) 版本：2020-02-17

### 第 179 次发布

发布时间：2026-07-09 04:24:38

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateRabbitMQVipInstance](https://cloud.tencent.com/document/api/1179/88134)

	* 新增入参：ClusterType, CdcClusterId


修改数据结构：

* [RabbitMQClusterInfo](https://cloud.tencent.com/document/api/1179/46089#RabbitMQClusterInfo)

	* 新增成员：ClusterType, CdcClusterId

* [RabbitMQVipInstance](https://cloud.tencent.com/document/api/1179/46089#RabbitMQVipInstance)

	* 新增成员：ClusterType, CdcClusterId




## TDSQL(tdmysql) 版本：2021-11-22

### 第 11 次发布

发布时间：2026-07-10 02:55:33

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeFlow](https://cloud.tencent.com/document/api/1376/128281)

	* 新增入参：FlowId




## 高性能计算平台(thpc) 版本：2023-03-21

### 第 39 次发布

发布时间：2026-07-09 04:34:36

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [SpaceInfo](https://cloud.tencent.com/document/api/1527/89579#SpaceInfo)

	* 新增成员：SpaceClass, PrivateIpAddresses




## 高性能计算平台(thpc) 版本：2022-04-01



## 高性能计算平台(thpc) 版本：2021-11-09



## TokenHub(tokenhub) 版本：2026-03-22

### 第 12 次发布

发布时间：2026-07-10 03:04:21

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [ModelChargingItem](https://cloud.tencent.com/document/api/1823/132279#ModelChargingItem)

	* 新增成员：PeakPrice




## 数据开发治理平台 WeData(wedata) 版本：2025-08-06



## 数据开发治理平台 WeData(wedata) 版本：2021-08-20

### 第 198 次发布

发布时间：2026-07-10 03:22:11

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [TaskOpsDto](https://cloud.tencent.com/document/api/1267/76336#TaskOpsDto)

	* 新增成员：ProxyTaskId, ProxyTaskTypeId




