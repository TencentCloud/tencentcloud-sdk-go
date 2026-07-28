# Release v1.3.145

## 运维安全中心（堡垒机）(bh) 版本：2023-04-18

### 第 34 次发布

发布时间：2026-07-29 01:13:59

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [LoginSetting](https://cloud.tencent.com/document/api/1025/74416#LoginSetting)

	* 新增成员：EnableSingleLogin




## 主机安全(cwp) 版本：2018-02-28

### 第 167 次发布

发布时间：2026-07-29 01:44:20

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateMalwareWhiteList](https://cloud.tencent.com/document/api/296/99674)

	* 新增入参：ProcessEventID


修改数据结构：

* [RiskProcessEvent](https://cloud.tencent.com/document/api/296/19867#RiskProcessEvent)

	* 新增成员：QUUID, ExeMd5




## 云数据库独享集群(dbdc) 版本：2020-10-29

### 第 8 次发布

发布时间：2026-07-29 01:54:16

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [AddNodesToDBCustomCluster](https://cloud.tencent.com/document/api/1322/132932)

	* 新增入参：Labels, Taints, HostName, HostNameType, DryRun

* [CreateDBCustomCluster](https://cloud.tencent.com/document/api/1322/132930)

	* 新增入参：DryRun

* [CreateDBCustomNodes](https://cloud.tencent.com/document/api/1322/132929)

	* 新增入参：ChargeType, NetworkMode, SystemDisk, DataDisks, HostName, DryRun, SecurityGroupIds

	* <font color="#dd0000">**修改入参**：</font>Period

* [DescribeDBCustomImages](https://cloud.tencent.com/document/api/1322/132924)

	* 新增入参：Filters

* [RemoveNodesFromDBCustomCluster](https://cloud.tencent.com/document/api/1322/132916)

	* 新增入参：LoginSettings

* [RenewDBCustomNode](https://cloud.tencent.com/document/api/1322/132915)

	* <font color="#dd0000">**修改入参**：</font>Period


新增数据结构：

* [Label](https://cloud.tencent.com/document/api/1322/74754#Label)
* [Taint](https://cloud.tencent.com/document/api/1322/74754#Taint)

修改数据结构：

* [DBCustomClusterNode](https://cloud.tencent.com/document/api/1322/74754#DBCustomClusterNode)

	* 新增成员：NetworkMode, EniIP

* [DBCustomImage](https://cloud.tencent.com/document/api/1322/74754#DBCustomImage)

	* 新增成员：OsType

* [DBCustomNode](https://cloud.tencent.com/document/api/1322/74754#DBCustomNode)

	* 新增成员：NetworkMode, EniIP

* [DataDisk](https://cloud.tencent.com/document/api/1322/74754#DataDisk)

	* <font color="#dd0000">**修改成员**：</font>DiskName




## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 312 次发布

发布时间：2026-07-29 02:08:38

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [ForwardRecord](https://cloud.tencent.com/document/api/1323/70369#ForwardRecord)

修改数据结构：

* [FlowApproverDetail](https://cloud.tencent.com/document/api/1323/70369#FlowApproverDetail)

	* 新增成员：ForwardRecords




## 物联网开发平台(iotexplorer) 版本：2019-04-23

### 第 151 次发布

发布时间：2026-07-29 02:20:47

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateDeviceSDPAnswer](https://cloud.tencent.com/document/api/1081/127863)

	* 新增入参：EnableSubPub




## 消息队列 MQTT 版(mqtt) 版本：2024-05-16

### 第 32 次发布

发布时间：2026-07-29 02:43:30

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeDeviceCertificateBackupHistory](https://cloud.tencent.com/document/api/1778/135211)
* [DescribeDeviceIdentityBackupHistory](https://cloud.tencent.com/document/api/1778/135212)
* [DescribeWillMessage](https://cloud.tencent.com/document/api/1778/135213)

新增数据结构：

* [DeviceCertificateBackupHistoryItem](https://cloud.tencent.com/document/api/1778/111031#DeviceCertificateBackupHistoryItem)
* [DeviceIdentityBackupHistoryItem](https://cloud.tencent.com/document/api/1778/111031#DeviceIdentityBackupHistoryItem)



## 云数据库 PostgreSQL(postgres) 版本：2017-03-12

### 第 69 次发布

发布时间：2026-07-29 02:48:27

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateDBProxy](https://cloud.tencent.com/document/api/409/135221)
* [DescribeDBProxy](https://cloud.tencent.com/document/api/409/135220)
* [DescribeDBProxySpecs](https://cloud.tencent.com/document/api/409/135219)
* [DestroyDBProxy](https://cloud.tencent.com/document/api/409/135218)
* [ModifyDBProxy](https://cloud.tencent.com/document/api/409/135217)
* [ModifyDBProxyAddress](https://cloud.tencent.com/document/api/409/135216)
* [ReloadBalanceDBProxyNode](https://cloud.tencent.com/document/api/409/135215)

新增数据结构：

* [ProxyAddress](https://cloud.tencent.com/document/api/409/16778#ProxyAddress)
* [ProxyGroupInfo](https://cloud.tencent.com/document/api/409/16778#ProxyGroupInfo)
* [ProxyNode](https://cloud.tencent.com/document/api/409/16778#ProxyNode)
* [ProxyNodeCustom](https://cloud.tencent.com/document/api/409/16778#ProxyNodeCustom)
* [ProxyRoute](https://cloud.tencent.com/document/api/409/16778#ProxyRoute)
* [ProxySpecItem](https://cloud.tencent.com/document/api/409/16778#ProxySpecItem)



## 边缘安全加速平台(teo) 版本：2022-09-01

### 第 152 次发布

发布时间：2026-07-29 03:12:00

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [ClientAttester](https://cloud.tencent.com/document/api/1552/80721#ClientAttester)

	* 新增成员：UsageLimit, MaxUsageCount




## 边缘安全加速平台(teo) 版本：2022-01-06



## TI-ONE 训练平台(tione) 版本：2021-11-11

### 第 127 次发布

发布时间：2026-07-29 03:15:46

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateTrainingTask](https://cloud.tencent.com/document/api/851/117377)

	* 新增入参：TrainToolConfig, ResourceSupplyAttribute

* [DescribeLogs](https://cloud.tencent.com/document/api/851/117840)

	* 新增入参：LogStream


新增数据结构：

* [TrainToolConfig](https://cloud.tencent.com/document/api/851/75051#TrainToolConfig)

修改数据结构：

* [LogIdentity](https://cloud.tencent.com/document/api/851/75051#LogIdentity)

	* 新增成员：PkgId, PkgLogId




## TI-ONE 训练平台(tione) 版本：2019-10-22



## 实时音视频(trtc) 版本：2019-07-22

### 第 147 次发布

发布时间：2026-07-29 03:23:29

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateLiveStreamModeration](https://cloud.tencent.com/document/api/647/135224)
* [DeleteLiveStreamModeration](https://cloud.tencent.com/document/api/647/135223)
* [DescribeLiveStreamModeration](https://cloud.tencent.com/document/api/647/135222)

新增数据结构：

* [Input](https://cloud.tencent.com/document/api/647/44055#Input)
* [LiveModerationParams](https://cloud.tencent.com/document/api/647/44055#LiveModerationParams)
* [LiveModerationStorageParams](https://cloud.tencent.com/document/api/647/44055#LiveModerationStorageParams)
* [SourceInfo](https://cloud.tencent.com/document/api/647/44055#SourceInfo)



## 数据开发治理平台 WeData(wedata) 版本：2025-08-06



## 数据开发治理平台 WeData(wedata) 版本：2021-08-20

### 第 199 次发布

发布时间：2026-07-29 03:41:50

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeInstancesByExecutors](https://cloud.tencent.com/document/api/1267/135225)

新增数据结构：

* [ExecutorTaskInstanceCount](https://cloud.tencent.com/document/api/1267/76336#ExecutorTaskInstanceCount)



