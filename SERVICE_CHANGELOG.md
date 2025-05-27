# Release v1.0.1175

## 运维安全中心（堡垒机）(bh) 版本：2023-04-18

### 第 12 次发布

发布时间：2025-05-28 01:08:27

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [BindDeviceResource](https://cloud.tencent.com/document/api/1025/74792)

	* 新增入参：ManageDimension, ManageAccountId, ManageAccount, ManageKubeconfig, Namespace, Workload

* [CreateResource](https://cloud.tencent.com/document/api/1025/100130)

	* 新增入参：ShareClb


修改数据结构：

* [Device](https://cloud.tencent.com/document/api/1025/74416#Device)

	* 新增成员：ManageDimension, ManageAccountId, Namespace, Workload, SyncPodCount, TotalPodCount

* [DeviceAccount](https://cloud.tencent.com/document/api/1025/74416#DeviceAccount)

	* 新增成员：BoundKubeconfig, IsK8SManageAccount

* [Resource](https://cloud.tencent.com/document/api/1025/74416#Resource)

	* 新增成员：IOAResourceId

* [SearchCommandResult](https://cloud.tencent.com/document/api/1025/74416#SearchCommandResult)

	* 新增成员：DeviceKind

* [SessionResult](https://cloud.tencent.com/document/api/1025/74416#SessionResult)

	* 新增成员：DeviceKind, Namespace, Workload, PodName




## 云数据库 MySQL(cdb) 版本：2017-03-20

### 第 194 次发布

发布时间：2025-05-28 01:09:51

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeCpuExpandHistory](https://cloud.tencent.com/document/api/236/118750)

修改接口：

* [StartCpuExpand](https://cloud.tencent.com/document/api/236/96017)

	* 新增入参：TimeIntervalStrategy, PeriodStrategy


新增数据结构：

* [AnalysisNodeInfo](https://cloud.tencent.com/document/api/236/15878#AnalysisNodeInfo)
* [HistoryJob](https://cloud.tencent.com/document/api/236/15878#HistoryJob)
* [PeriodStrategy](https://cloud.tencent.com/document/api/236/15878#PeriodStrategy)
* [TImeCycle](https://cloud.tencent.com/document/api/236/15878#TImeCycle)
* [TimeInterval](https://cloud.tencent.com/document/api/236/15878#TimeInterval)
* [TimeIntervalStrategy](https://cloud.tencent.com/document/api/236/15878#TimeIntervalStrategy)

修改数据结构：

* [AccountInfo](https://cloud.tencent.com/document/api/236/15878#AccountInfo)

* [InstanceInfo](https://cloud.tencent.com/document/api/236/15878#InstanceInfo)

	* 新增成员：AnalysisNodeInfos, DeviceBandwidth




## 腾讯云数据仓库 TCHouse-D(cdwdoris) 版本：2021-12-28

### 第 51 次发布

发布时间：2025-05-28 01:10:44

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateInstanceNew](https://cloud.tencent.com/document/api/1387/102611)

	* 新增入参：CacheDataDiskSize




## TDSQL-C MySQL 版(cynosdb) 版本：2019-01-07

### 第 135 次发布

发布时间：2025-05-28 01:14:12

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeServerlessStrategy](https://cloud.tencent.com/document/api/1003/110601)

	* 新增出参：AutoArchive

* [ModifyServerlessStrategy](https://cloud.tencent.com/document/api/1003/110600)

	* 新增入参：AutoArchive

* [RollbackToNewCluster](https://cloud.tencent.com/document/api/1003/104727)

	* 新增入参：AutoArchive


新增数据结构：

* [GdnTaskInfo](https://cloud.tencent.com/document/api/1003/48097#GdnTaskInfo)

修改数据结构：

* [BizTaskInfo](https://cloud.tencent.com/document/api/1003/48097#BizTaskInfo)

	* 新增成员：GdnTaskInfo

* [CynosdbClusterDetail](https://cloud.tencent.com/document/api/1003/48097#CynosdbClusterDetail)

	* 新增成员：UsedArchiveStorage, ArchiveStatus, ArchiveProgress




## 媒体处理(mps) 版本：2019-06-12

### 第 125 次发布

发布时间：2025-05-28 01:23:03

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeStreamLinkFlowMediaStatistics](https://cloud.tencent.com/document/api/862/82438)

	* 新增入参：RemoteIp

* [DescribeStreamLinkFlowSRTStatistics](https://cloud.tencent.com/document/api/862/82436)

	* 新增入参：RemoteIp

* [DescribeStreamLinkFlowStatistics](https://cloud.tencent.com/document/api/862/82435)

	* 新增入参：RemoteIp




## 流计算 Oceanus(oceanus) 版本：2019-04-22

### 第 68 次发布

发布时间：2025-05-28 01:23:41

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeJobs](https://cloud.tencent.com/document/api/849/52008)

	* 新增入参：ConnectorOptions


新增数据结构：

* [ResourceRefLatest](https://cloud.tencent.com/document/api/849/52010#ResourceRefLatest)
* [Setats](https://cloud.tencent.com/document/api/849/52010#Setats)
* [SetatsCvmInfo](https://cloud.tencent.com/document/api/849/52010#SetatsCvmInfo)
* [SetatsDisk](https://cloud.tencent.com/document/api/849/52010#SetatsDisk)
* [Warehouse](https://cloud.tencent.com/document/api/849/52010#Warehouse)

修改数据结构：

* [Cluster](https://cloud.tencent.com/document/api/849/52010#Cluster)

	* 新增成员：Setats




## 邮件推送(ses) 版本：2020-10-02

### 第 30 次发布

发布时间：2025-05-28 01:25:38

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [ListAddressUnsubscribeConfig](https://cloud.tencent.com/document/api/1288/118751)

新增数据结构：

* [AddressUnsubscribeConfigData](https://cloud.tencent.com/document/api/1288/51053#AddressUnsubscribeConfigData)



## SSL 证书(ssl) 版本：2019-12-05

### 第 83 次发布

发布时间：2025-05-28 01:26:15

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [ClbInstanceDetail](https://cloud.tencent.com/document/api/400/41679#ClbInstanceDetail)

	* 新增成员：Forward




## TI-ONE 训练平台(tione) 版本：2021-11-11

### 第 82 次发布

发布时间：2025-05-28 01:29:29

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [Pod](https://cloud.tencent.com/document/api/851/75051#Pod)

	* 新增成员：StartScheduleTime, Message

* [Service](https://cloud.tencent.com/document/api/851/75051#Service)

	* 新增成员：MonitorSource

* [ServiceGroup](https://cloud.tencent.com/document/api/851/75051#ServiceGroup)

	* 新增成员：SubUinName




## TI-ONE 训练平台(tione) 版本：2019-10-22



## 私有网络(vpc) 版本：2017-03-12

### 第 259 次发布

发布时间：2025-05-28 01:32:35

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [AllocateAddresses](https://cloud.tencent.com/document/api/215/16699)

* [ModifyAddressesBandwidth](https://cloud.tencent.com/document/api/215/19214)




## 数据开发治理平台 WeData(wedata) 版本：2021-08-20

### 第 153 次发布

发布时间：2025-05-28 01:33:59

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateOpsMakePlan](https://cloud.tencent.com/document/api/1267/95231)

	* 新增入参：MakeType, StatusList, MakeCheckEventType

* [DownloadLogByLine](https://cloud.tencent.com/document/api/1267/118484)

	* 新增入参：QueryFileFlag, ExtInfo


新增数据结构：

* [ExtensionInfoVO](https://cloud.tencent.com/document/api/1267/76336#ExtensionInfoVO)

修改数据结构：

* [InstanceLogByLine](https://cloud.tencent.com/document/api/1267/76336#InstanceLogByLine)

	* 新增成员：ExecutionExtendedProps, ExtInfo




