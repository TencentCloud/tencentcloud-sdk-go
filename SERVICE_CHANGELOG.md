# Release v1.3.134

## 应用型负载均衡(alb) 版本：2025-10-30

### 第 3 次发布

发布时间：2026-07-15 01:08:49

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyHealthCheckTemplate](https://cloud.tencent.com/document/api/1822/133734)

	* <font color="#dd0000">**删除入参**：</font>Tags




## 云数据库 MySQL(cdb) 版本：2017-03-20

### 第 223 次发布

发布时间：2026-07-15 01:20:20

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [SwitchDBInstanceMasterSlave](https://cloud.tencent.com/document/api/236/52171)


修改数据结构：

* [ProxyAddress](https://cloud.tencent.com/document/api/236/15878#ProxyAddress)

	* 新增成员：Region




## 日志服务(cls) 版本：2020-10-16

### 第 167 次发布

发布时间：2026-07-15 01:31:30

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateScheduledSql](https://cloud.tencent.com/document/api/614/95138)

	* 新增入参：ProcessPeriodUnit

* [ModifyScheduledSql](https://cloud.tencent.com/document/api/614/95518)

	* 新增入参：ProcessPeriodUnit


修改数据结构：

* [ScheduledSqlTaskInfo](https://cloud.tencent.com/document/api/614/56471#ScheduledSqlTaskInfo)

	* 新增成员：ProcessPeriodUnit




## 云安全一体化平台(csip) 版本：2022-11-21

### 第 91 次发布

发布时间：2026-07-15 01:35:26

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [ModifyCosAuditBucketMonitorStatus](https://cloud.tencent.com/document/api/664/134143)



## Elasticsearch Service(es) 版本：2025-01-01



## Elasticsearch Service(es) 版本：2018-04-16

### 第 104 次发布

发布时间：2026-07-15 01:59:52

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeAutoScaleDiskInfo](https://cloud.tencent.com/document/api/845/134148)
* [DescribeClusterDiskRange](https://cloud.tencent.com/document/api/845/134156)
* [DescribeCosBackupStrategyViews](https://cloud.tencent.com/document/api/845/134145)
* [DescribeEsInstanceEventLists](https://cloud.tencent.com/document/api/845/134155)
* [DescribeEventDataDetail](https://cloud.tencent.com/document/api/845/134154)
* [DescribeEventInfoList](https://cloud.tencent.com/document/api/845/134153)
* [DescribeForceMergeTask](https://cloud.tencent.com/document/api/845/134147)
* [DescribeLogstashViews](https://cloud.tencent.com/document/api/845/134157)
* [DescribeRegions](https://cloud.tencent.com/document/api/845/134152)
* [DescribeRequestInstancePolicy](https://cloud.tencent.com/document/api/845/134146)
* [DescribeSnapshotViews](https://cloud.tencent.com/document/api/845/134144)
* [DescribeUpgrade](https://cloud.tencent.com/document/api/845/134151)
* [QueryZoneResource](https://cloud.tencent.com/document/api/845/134150)
* [QueryZoneResourceForLogstash](https://cloud.tencent.com/document/api/845/134149)

修改接口：

* [RequestInstancesByGet](https://cloud.tencent.com/document/api/845/130846)

	* 新增入参：InstanceId, Uri, Caller

	* 新增出参：Detail


新增数据结构：

* [CrontabTaskInfo](https://cloud.tencent.com/document/api/845/30634#CrontabTaskInfo)
* [DiskCountRange](https://cloud.tencent.com/document/api/845/30634#DiskCountRange)
* [DiskSizeRange](https://cloud.tencent.com/document/api/845/30634#DiskSizeRange)
* [EventDataDetail](https://cloud.tencent.com/document/api/845/30634#EventDataDetail)
* [EventDataInfoOverview](https://cloud.tencent.com/document/api/845/30634#EventDataInfoOverview)
* [EventTypeInfo](https://cloud.tencent.com/document/api/845/30634#EventTypeInfo)
* [LogstashNodeTypeResource](https://cloud.tencent.com/document/api/845/30634#LogstashNodeTypeResource)
* [LogstashNodeView](https://cloud.tencent.com/document/api/845/30634#LogstashNodeView)
* [LogstashZoneResource](https://cloud.tencent.com/document/api/845/30634#LogstashZoneResource)
* [NodeTypeDiskSizeRange](https://cloud.tencent.com/document/api/845/30634#NodeTypeDiskSizeRange)
* [NodeTypeInfo](https://cloud.tencent.com/document/api/845/30634#NodeTypeInfo)
* [NodeTypeResource](https://cloud.tencent.com/document/api/845/30634#NodeTypeResource)
* [RegionsData](https://cloud.tencent.com/document/api/845/30634#RegionsData)
* [ZoneResource](https://cloud.tencent.com/document/api/845/30634#ZoneResource)



## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 307 次发布

发布时间：2026-07-15 02:01:12

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateOrganizationAuthUrl](https://cloud.tencent.com/document/api/1323/105134)

	* 新增入参：AddressSame




## 腾讯电子签（基础版）(essbasic) 版本：2021-05-26

### 第 264 次发布

发布时间：2026-07-15 02:03:00

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateBatchAdminChangeInvitations](https://cloud.tencent.com/document/api/1420/134159)
* [CreateBatchAdminChangeInvitationsUrl](https://cloud.tencent.com/document/api/1420/134158)
* [ModifyOrganizationBusinessInfo](https://cloud.tencent.com/document/api/1420/134160)

新增数据结构：

* [AdminChangeInvitationInfo](https://cloud.tencent.com/document/api/1420/61525#AdminChangeInvitationInfo)

修改数据结构：

* [OrganizationAuthorizationOptions](https://cloud.tencent.com/document/api/1420/61525#OrganizationAuthorizationOptions)

	* 新增成员：AddressSame




## 腾讯电子签（基础版）(essbasic) 版本：2020-12-22



## 媒体处理(mps) 版本：2019-06-12

### 第 216 次发布

发布时间：2026-07-15 02:28:01

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeImageTasks](https://cloud.tencent.com/document/api/862/134161)

新增数据结构：

* [ImageTaskInfo](https://cloud.tencent.com/document/api/862/37615#ImageTaskInfo)



## 边缘安全加速平台(teo) 版本：2022-09-01

### 第 149 次发布

发布时间：2026-07-15 03:15:09

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateLogAnalysisDownloadTask](https://cloud.tencent.com/document/api/1552/134164)
* [DescribeLogAnalysisDetail](https://cloud.tencent.com/document/api/1552/134163)
* [DescribeLogAnalysisDownloadTasks](https://cloud.tencent.com/document/api/1552/134162)

新增数据结构：

* [LogAnalysisDownloadTask](https://cloud.tencent.com/document/api/1552/80721#LogAnalysisDownloadTask)
* [LogItem](https://cloud.tencent.com/document/api/1552/80721#LogItem)



## 边缘安全加速平台(teo) 版本：2022-01-06



## 云点播(vod) 版本：2024-07-18



## 云点播(vod) 版本：2018-07-17

### 第 272 次发布

发布时间：2026-07-15 03:44:08

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [OverrideTranscodeParameter](https://cloud.tencent.com/document/api/266/31773#OverrideTranscodeParameter)

修改数据结构：

* [AudioTemplateInfoForUpdate](https://cloud.tencent.com/document/api/266/31773#AudioTemplateInfoForUpdate)

	* 新增成员：StreamSelects

* [TranscodeTaskInput](https://cloud.tencent.com/document/api/266/31773#TranscodeTaskInput)

	* 新增成员：OverrideParameter




