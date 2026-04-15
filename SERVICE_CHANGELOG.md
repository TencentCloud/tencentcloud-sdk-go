# Release v1.3.80

## 腾讯云数据仓库 TCHouse-D(cdwdoris) 版本：2021-12-28

### 第 61 次发布

发布时间：2026-04-16 01:14:37

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeAreaRegion](https://cloud.tencent.com/document/api/1387/109559)

	* 新增出参：IsolationDays

* [DescribeBackUpSchedules](https://cloud.tencent.com/document/api/1387/109556)

	* 新增入参：ScheduleId

* [DescribeClusterConfigs](https://cloud.tencent.com/document/api/1387/102619)

	* 新增出参：IPDBFileSizeLimit

* [DescribeWorkloadGroup](https://cloud.tencent.com/document/api/1387/109437)

	* 新增出参：MonitorStatus


修改数据结构：

* [BackUpJobDisplay](https://cloud.tencent.com/document/api/1387/102385#BackUpJobDisplay)

	* 新增成员：ScheduleId

* [InstanceInfo](https://cloud.tencent.com/document/api/1387/102385#InstanceInfo)

	* 新增成员：EnableSqlConv, TimeZone

* [SnapshotRemainPolicy](https://cloud.tencent.com/document/api/1387/102385#SnapshotRemainPolicy)

	* 新增成员：RemainDaysUnit




## 云防火墙(cfw) 版本：2019-09-04

### 第 96 次发布

发布时间：2026-04-16 01:15:34

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeNDRAssetIdentificationList](https://cloud.tencent.com/document/api/1132/130555)

修改接口：

* [DescribeBlockIgnoreList](https://cloud.tencent.com/document/api/1132/86691)

	* <font color="#dd0000">**修改入参**：</font>ShowType


新增数据结构：

* [FieldOption](https://cloud.tencent.com/document/api/1132/49071#FieldOption)
* [NDRAssetCategoryStats](https://cloud.tencent.com/document/api/1132/49071#NDRAssetCategoryStats)
* [NDRAssetServiceInfo](https://cloud.tencent.com/document/api/1132/49071#NDRAssetServiceInfo)
* [NDRAssetServiceStats](https://cloud.tencent.com/document/api/1132/49071#NDRAssetServiceStats)
* [OperatorFilter](https://cloud.tencent.com/document/api/1132/49071#OperatorFilter)



## 配置审计(config) 版本：2022-08-02

### 第 10 次发布

发布时间：2026-04-16 01:19:12

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ListAlarmPolicy](https://cloud.tencent.com/document/api/1579/130466)

	* 新增入参：Limit




## 云安全一体化平台(csip) 版本：2022-11-21

### 第 74 次发布

发布时间：2026-04-16 01:19:30

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [RiskDetailItem](https://cloud.tencent.com/document/api/664/90825#RiskDetailItem)

	* 新增成员：AppID




## 数据传输服务(dts) 版本：2021-12-06

### 第 51 次发布

发布时间：2026-04-15 12:18:00

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除数据结构**：</font>

* OnlineDDL

修改数据结构：

* [Objects](https://cloud.tencent.com/document/api/571/82108#Objects)

	* <font color="#dd0000">**删除成员**：</font>OnlineDDL




## 数据传输服务(dts) 版本：2018-03-30



## 弹性 MapReduce(emr) 版本：2019-01-03

### 第 136 次发布

发布时间：2026-04-16 01:27:20

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeBootScript](https://cloud.tencent.com/document/api/589/130557)
* [ModifyBootScript](https://cloud.tencent.com/document/api/589/130556)

新增数据结构：

* [DescribeBootScriptRsp](https://cloud.tencent.com/document/api/589/33981#DescribeBootScriptRsp)
* [PreExecuteFileSetting](https://cloud.tencent.com/document/api/589/33981#PreExecuteFileSetting)



## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 288 次发布

发布时间：2026-04-16 01:28:16

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateBatchQuickSignUrl](https://cloud.tencent.com/document/api/1323/101058)

	* 新增入参：ExpiredOn

* [CreateFlowSignUrl](https://cloud.tencent.com/document/api/1323/85818)

	* 新增入参：ExpiredOn


修改数据结构：

* [ApproverOption](https://cloud.tencent.com/document/api/1323/70369#ApproverOption)

	* 新增成员：ApproverMobileMode

* [FlowCreateApprover](https://cloud.tencent.com/document/api/1323/70369#FlowCreateApprover)

	* 新增成员：ApproverEmail




## 腾讯电子签（基础版）(essbasic) 版本：2021-05-26

### 第 254 次发布

发布时间：2026-04-16 01:29:00

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ChannelCreateBatchQuickSignUrl](https://cloud.tencent.com/document/api/1420/101059)

	* 新增入参：ExpiredOn

* [ChannelCreateFlowSignUrl](https://cloud.tencent.com/document/api/1420/85819)

	* 新增入参：ExpiredOn


修改数据结构：

* [ApproverOption](https://cloud.tencent.com/document/api/1420/61525#ApproverOption)

	* 新增成员：ApproverMobileMode

* [FlowApproverInfo](https://cloud.tencent.com/document/api/1420/61525#FlowApproverInfo)

	* 新增成员：ApproverEmail




## 腾讯电子签（基础版）(essbasic) 版本：2020-12-22



## 智能视图计算平台(iss) 版本：2023-05-17

### 第 31 次发布

发布时间：2026-04-16 01:34:11

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [AddRecordBackupTemplate](https://cloud.tencent.com/document/api/1344/95913)

	* 新增入参：DayOffset


修改数据结构：

* [AddRecordBackupTemplateData](https://cloud.tencent.com/document/api/1344/95952#AddRecordBackupTemplateData)

	* 新增成员：DayOffset

* [DescribeRecordBackupTemplateData](https://cloud.tencent.com/document/api/1344/95952#DescribeRecordBackupTemplateData)

	* 新增成员：DayOffset

* [ListRecordBackupTemplatesData](https://cloud.tencent.com/document/api/1344/95952#ListRecordBackupTemplatesData)

	* 新增成员：DayOffset

* [UpdateRecordBackupTemplateData](https://cloud.tencent.com/document/api/1344/95952#UpdateRecordBackupTemplateData)

	* 新增成员：DayOffset




## 媒体处理(mps) 版本：2019-06-12

### 第 193 次发布

发布时间：2026-04-16 01:45:15

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [TermBase](https://cloud.tencent.com/document/api/862/37615#TermBase)

	* 新增成员：SrcLanguage, DstLanguage




## 医疗报告结构化(mrs) 版本：2020-09-10

### 第 39 次发布

发布时间：2026-04-16 01:49:19

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [ItemCoordinate](https://cloud.tencent.com/document/api/1314/56230#ItemCoordinate)

修改数据结构：

* [IndicatorItem](https://cloud.tencent.com/document/api/1314/56230#IndicatorItem)

	* 新增成员：ItemCoords




## 流计算 Oceanus(oceanus) 版本：2019-04-22

### 第 82 次发布

发布时间：2026-04-16 01:50:30

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [Cluster](https://cloud.tencent.com/document/api/849/52010#Cluster)

	* 新增成员：ClusterProcessMsg, MaxCuPerJob

* [Setats](https://cloud.tencent.com/document/api/849/52010#Setats)

	* 新增成员：WebUIType




## 云开发 CloudBase(tcb) 版本：2018-06-08

### 第 136 次发布

发布时间：2026-04-16 02:11:53

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [StaticStorageInfo](https://cloud.tencent.com/document/api/876/34822#StaticStorageInfo)

	* 新增成员：AccessExpire, ExternalStorage

* [StorageInfo](https://cloud.tencent.com/document/api/876/34822#StorageInfo)

	* 新增成员：ExternalStorage




## 云托管 CloudBase Run(tcbr) 版本：2022-02-17

### 第 27 次发布

发布时间：2026-04-16 02:12:55

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [StartVersionInstance](https://cloud.tencent.com/document/api/1243/130559)
* [StopVersionInstance](https://cloud.tencent.com/document/api/1243/130558)



## 消息队列 TDMQ(tdmq) 版本：2020-02-17

### 第 170 次发布

发布时间：2026-04-16 02:20:34

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateProCluster](https://cloud.tencent.com/document/api/1179/101433)

	* 新增入参：InstanceVersion


修改数据结构：

* [Role](https://cloud.tencent.com/document/api/1179/46089#Role)

	* 新增成员：TokenType, SecretName, RotateFreq




## TDSQL(tdmysql) 版本：2021-11-22

### 第 3 次发布

发布时间：2026-04-16 02:22:32

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeDBSArchiveLogs](https://cloud.tencent.com/document/api/1376/130560)

新增数据结构：

* [ArchiveLogModel](https://cloud.tencent.com/document/api/1376/128305#ArchiveLogModel)



## 腾讯混元生视频(vclm) 版本：2024-05-23

### 第 33 次发布

发布时间：2026-04-16 02:41:25

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeImageToVideoJob](https://cloud.tencent.com/document/api/1616/130575)
* [DescribeImageToVideoViduJob](https://cloud.tencent.com/document/api/1616/130574)
* [DescribeMotionControlKlingJob](https://cloud.tencent.com/document/api/1616/130573)
* [DescribeReferenceToVideoViduJob](https://cloud.tencent.com/document/api/1616/130572)
* [DescribeTextToVideoJob](https://cloud.tencent.com/document/api/1616/130571)
* [DescribeTextToVideoViduJob](https://cloud.tencent.com/document/api/1616/130570)
* [DescribeVideoEditKlingJob](https://cloud.tencent.com/document/api/1616/130569)
* [DescribeVideoExtendKlingJob](https://cloud.tencent.com/document/api/1616/130568)
* [SubmitImageToVideoJob](https://cloud.tencent.com/document/api/1616/130567)
* [SubmitMotionControlKlingJob](https://cloud.tencent.com/document/api/1616/130566)
* [SubmitReferenceToVideoViduJob](https://cloud.tencent.com/document/api/1616/130565)
* [SubmitTextToVideoJob](https://cloud.tencent.com/document/api/1616/130564)
* [SubmitTextToVideoViduJob](https://cloud.tencent.com/document/api/1616/130563)
* [SubmitVideoEditKlingJob](https://cloud.tencent.com/document/api/1616/130562)
* [SubmitVideoExtendKlingJob](https://cloud.tencent.com/document/api/1616/130561)

新增数据结构：

* [CameraControl](https://cloud.tencent.com/document/api/1616/107808#CameraControl)
* [CameraControlConfig](https://cloud.tencent.com/document/api/1616/107808#CameraControlConfig)
* [DynamicMask](https://cloud.tencent.com/document/api/1616/107808#DynamicMask)
* [Element](https://cloud.tencent.com/document/api/1616/107808#Element)
* [ImageInfo](https://cloud.tencent.com/document/api/1616/107808#ImageInfo)
* [MultiPrompt](https://cloud.tencent.com/document/api/1616/107808#MultiPrompt)
* [ReferVideoInfo](https://cloud.tencent.com/document/api/1616/107808#ReferVideoInfo)
* [ReferenceSubject](https://cloud.tencent.com/document/api/1616/107808#ReferenceSubject)
* [Trajectory](https://cloud.tencent.com/document/api/1616/107808#Trajectory)
* [Voice](https://cloud.tencent.com/document/api/1616/107808#Voice)



## 云点播(vod) 版本：2024-07-18



## 云点播(vod) 版本：2018-07-17

### 第 248 次发布

发布时间：2026-04-16 02:43:21

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [AigcImageOutputConfig](https://cloud.tencent.com/document/api/266/31773#AigcImageOutputConfig)

	* 新增成员：OutputImageCount

* [AigcVideoOutputConfig](https://cloud.tencent.com/document/api/266/31773#AigcVideoOutputConfig)

	* 新增成员：EnableBGM




## 私有网络(vpc) 版本：2017-03-12

### 第 302 次发布

发布时间：2026-04-16 02:46:55

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [ModifyNatGatewayAdvancedAttribute](https://cloud.tencent.com/document/api/215/130576)



