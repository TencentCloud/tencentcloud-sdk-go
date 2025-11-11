# Release v1.1.54

## 日志服务(cls) 版本：2020-10-16

### 第 139 次发布

发布时间：2025-11-12 01:14:43

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateDataTransform](https://cloud.tencent.com/document/api/614/72184)

	* 新增入参：BackupGiveUpData, HasServicesLog, ProcessFromTimestamp, ProcessToTimestamp, TaskId, DataTransformSqlDataSources, EnvInfos

* [ModifyDataTransform](https://cloud.tencent.com/document/api/614/72181)

	* 新增入参：BackupGiveUpData, KeepFailureLog, FailureLogKey, DataTransformSqlDataSources, EnvInfos


新增数据结构：

* [DataTransformSqlDataSource](https://cloud.tencent.com/document/api/614/56471#DataTransformSqlDataSource)
* [EnvInfo](https://cloud.tencent.com/document/api/614/56471#EnvInfo)

修改数据结构：

* [DataTransformTaskInfo](https://cloud.tencent.com/document/api/614/56471#DataTransformTaskInfo)

	* 新增成员：BackupTopicID, BackupGiveUpData, HasServicesLog, TaskDstCount, ProcessFromTimestamp, ProcessToTimestamp, DataTransformSqlDataSources, EnvInfos




## 弹性 MapReduce(emr) 版本：2019-01-03

### 第 123 次发布

发布时间：2025-11-12 01:20:50

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [NodeSpecDisk](https://cloud.tencent.com/document/api/589/33981#NodeSpecDisk)

	* 新增成员：IsSpecialDisk




## 云游戏(gs) 版本：2019-11-18

### 第 57 次发布

发布时间：2025-11-12 01:23:15

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateAndroidInstancesAccessToken](https://cloud.tencent.com/document/api/1162/119708)

	* 新增入参：Mode, UserIP




## iOA 零信任安全管理系统(ioa) 版本：2022-06-01

### 第 29 次发布

发布时间：2025-11-12 01:24:20

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [DeviceDetail](https://cloud.tencent.com/document/api/1092/102488#DeviceDetail)

	* 新增成员：ScreenRecordingPermission, DiskAccessPermission




## 媒体处理(mps) 版本：2019-06-12

### 第 156 次发布

发布时间：2025-11-12 01:36:01

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateStreamLinkFlow](https://cloud.tencent.com/document/api/862/82445)

	* <font color="#dd0000">**修改入参**：</font>OutputGroup


修改数据结构：

* [AdaptiveDynamicStreamingTaskInput](https://cloud.tencent.com/document/api/862/37615#AdaptiveDynamicStreamingTaskInput)

	* 新增成员：KeyPTSList




## 私有网络(vpc) 版本：2017-03-12

### 第 274 次发布

发布时间：2025-11-12 02:30:40

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [AllocateAddresses](https://cloud.tencent.com/document/api/215/16699)

	* 新增入参：IPChargeType

* [InquiryPriceAllocateAddresses](https://cloud.tencent.com/document/api/215/114855)

	* 新增入参：IPChargeType


修改数据结构：

* [InternetPrice](https://cloud.tencent.com/document/api/215/15824#InternetPrice)

	* 新增成员：IPPrice, OriginalPrice, DiscountPrice




