# Release v1.1.13

## 云硬盘(cbs) 版本：2017-03-12

### 第 72 次发布

发布时间：2025-08-19 01:15:53

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [InquiryPriceResizeDisk](https://cloud.tencent.com/document/api/362/16320)

	* 新增入参：DiskIds

	* <font color="#dd0000">**修改入参**：</font>DiskId




## 云联络中心(ccc) 版本：2020-02-10

### 第 105 次发布

发布时间：2025-08-19 01:16:20

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateAICall](https://cloud.tencent.com/document/api/679/111211)

	* 新增入参：EnableComplianceAudio




## 数据湖计算 DLC(dlc) 版本：2021-01-25

### 第 134 次发布

发布时间：2025-08-19 01:37:33

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [AttachDataMaskPolicy](https://cloud.tencent.com/document/api/1342/122615)
* [CreateDataMaskStrategy](https://cloud.tencent.com/document/api/1342/122619)
* [DeleteDataMaskStrategy](https://cloud.tencent.com/document/api/1342/122618)
* [DescribeDataMaskStrategies](https://cloud.tencent.com/document/api/1342/122617)
* [DescribeUDFPolicy](https://cloud.tencent.com/document/api/1342/122614)
* [UpdateDataMaskStrategy](https://cloud.tencent.com/document/api/1342/122616)
* [UpdateUDFPolicy](https://cloud.tencent.com/document/api/1342/122613)

新增数据结构：

* [DataMaskStrategy](https://cloud.tencent.com/document/api/1342/53778#DataMaskStrategy)
* [DataMaskStrategyPolicy](https://cloud.tencent.com/document/api/1342/53778#DataMaskStrategyPolicy)
* [UDFPolicyInfo](https://cloud.tencent.com/document/api/1342/53778#UDFPolicyInfo)



## Elasticsearch Service(es) 版本：2025-01-01



## Elasticsearch Service(es) 版本：2018-04-16

### 第 86 次发布

发布时间：2025-08-19 01:46:17

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [GpuInfo](https://cloud.tencent.com/document/api/845/30634#GpuInfo)

修改数据结构：

* [InstanceInfo](https://cloud.tencent.com/document/api/845/30634#InstanceInfo)

	* 新增成员：IsCdzLite

* [NodeInfo](https://cloud.tencent.com/document/api/845/30634#NodeInfo)

	* 新增成员：GpuInfo




## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 242 次发布

发布时间：2025-08-19 01:47:07

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [EmbedUrlOption](https://cloud.tencent.com/document/api/1323/70369#EmbedUrlOption)

	* 新增成员：SealDescription, ForbidEditSealDescription




## 腾讯电子签（基础版）(essbasic) 版本：2021-05-26

### 第 231 次发布

发布时间：2025-08-19 01:48:23

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [EmbedUrlOption](https://cloud.tencent.com/document/api/1420/61525#EmbedUrlOption)

	* 新增成员：SealDescription, ForbidEditSealDescription

* [OrganizationAuthUrl](https://cloud.tencent.com/document/api/1420/61525#OrganizationAuthUrl)

	* 新增成员：OrganizationName, SubTaskId




## 腾讯电子签（基础版）(essbasic) 版本：2020-12-22



## iOA 零信任安全管理系统(ioa) 版本：2022-06-01

### 第 19 次发布

发布时间：2025-08-19 01:53:55

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateDeviceVirtualGroup](https://cloud.tencent.com/document/api/1092/112071)

	* <font color="#dd0000">**修改入参**：</font>DeviceVirtualGroupName

* [CreatePrivilegeCode](https://cloud.tencent.com/document/api/1092/118218)

	* 新增入参：OsType

	* <font color="#dd0000">**修改入参**：</font>Mid

* [ModifyVirtualDeviceGroups](https://cloud.tencent.com/document/api/1092/120505)

	* <font color="#dd0000">**修改入参**：</font>DeviceList


修改数据结构：

* [DescribeDeviceHardwareInfoItem](https://cloud.tencent.com/document/api/1092/102488#DescribeDeviceHardwareInfoItem)

	* 新增成员：RemarkName




## 媒体处理(mps) 版本：2019-06-12

### 第 139 次发布

发布时间：2025-08-19 02:14:32

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateLiveRecordTemplate](https://cloud.tencent.com/document/api/862/114299)

	* 新增入参：RecordType

* [ModifyLiveRecordTemplate](https://cloud.tencent.com/document/api/862/114296)

	* 新增入参：RecordType


修改数据结构：

* [LiveRecordTemplate](https://cloud.tencent.com/document/api/862/37615#LiveRecordTemplate)

	* 新增成员：RecordType




## 游戏数据库 TcaplusDB(tcaplusdb) 版本：2019-08-23

### 第 23 次发布

发布时间：2025-08-19 02:43:21

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [CompareTablesInfo](https://cloud.tencent.com/document/api/596/39686#CompareTablesInfo)

	* <font color="#dd0000">**修改成员**：</font>SrcTableClusterId, SrcTableGroupId, SrcTableName, DstTableClusterId, DstTableGroupId, DstTableName, SrcTableInstanceId, DstTableInstanceId

* [KafkaInfo](https://cloud.tencent.com/document/api/596/39686#KafkaInfo)

	* <font color="#dd0000">**修改成员**：</font>Address, Topic, User, Password, Instance, IsVpc

* [SnapshotResult](https://cloud.tencent.com/document/api/596/39686#SnapshotResult)

	* <font color="#dd0000">**修改成员**：</font>ApplicationId

* [TaskInfoNew](https://cloud.tencent.com/document/api/596/39686#TaskInfoNew)

	* <font color="#dd0000">**修改成员**：</font>TableGroupId, TableGroupName, TableName




## 消息队列 TDMQ(tdmq) 版本：2020-02-17

### 第 156 次发布

发布时间：2025-08-19 02:54:43

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [ExecuteDisasterRecovery](https://cloud.tencent.com/document/api/1179/122621)

修改接口：

* [DescribeEnvironmentRoles](https://cloud.tencent.com/document/api/1179/49051)

	* <font color="#dd0000">**修改入参**：</font>ClusterId




## TSF-Polaris&ZK&网关(tse) 版本：2020-12-07

### 第 99 次发布

发布时间：2025-08-19 03:14:59

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DeleteCloudNativeAPIGatewayService](https://cloud.tencent.com/document/api/1364/94846)

	* 新增入参：DeleteRoutes




