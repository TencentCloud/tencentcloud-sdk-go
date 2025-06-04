# Release v1.0.1180

## 本地专用集群(cdc) 版本：2020-12-14

### 第 13 次发布

发布时间：2025-06-05 01:14:47

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [DedicatedClusterOrder](https://cloud.tencent.com/document/api/1346/73770#DedicatedClusterOrder)

* [HostInfo](https://cloud.tencent.com/document/api/1346/73770#HostInfo)




## 负载均衡(clb) 版本：2018-03-17

### 第 130 次发布

发布时间：2025-06-05 01:18:53

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeLBOperateProtect](https://cloud.tencent.com/document/api/214/118976)

新增数据结构：

* [LBOperateProtectInfo](https://cloud.tencent.com/document/api/214/30694#LBOperateProtectInfo)



## 弹性 MapReduce(emr) 版本：2019-01-03

### 第 106 次发布

发布时间：2025-06-05 01:32:33

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [InspectionTaskSettings](https://cloud.tencent.com/document/api/589/33981#InspectionTaskSettings)

	* 新增成员：SettingsJson

* [OverviewRow](https://cloud.tencent.com/document/api/589/33981#OverviewRow)

	* 新增成员：StoreFileNum




## 云游戏(gs) 版本：2019-11-18

### 第 37 次发布

发布时间：2025-06-05 01:36:51

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateAndroidInstanceADB](https://cloud.tencent.com/document/api/1162/118978)
* [InstallAndroidInstancesAppWithURL](https://cloud.tencent.com/document/api/1162/118979)
* [ModifyAndroidInstancesProperties](https://cloud.tencent.com/document/api/1162/118980)

修改接口：

* [CreateAndroidAppVersion](https://cloud.tencent.com/document/api/1162/117913)

	* 新增入参：CleanupMode


新增数据结构：

* [AndroidInstanceDevice](https://cloud.tencent.com/document/api/1162/40743#AndroidInstanceDevice)
* [AndroidInstanceProperty](https://cloud.tencent.com/document/api/1162/40743#AndroidInstanceProperty)

修改数据结构：

* [AndroidAppVersionInfo](https://cloud.tencent.com/document/api/1162/40743#AndroidAppVersionInfo)

	* 新增成员：CleanupMode




## 云开发低码(lowcode) 版本：2021-01-08

### 第 5 次发布

发布时间：2025-06-05 01:46:17

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [KnowledgeDocumentSet](https://cloud.tencent.com/document/api/1301/68878#KnowledgeDocumentSet)

	* 新增成员：DocStatus, ErrMsg, FileId

* [QureyKnowledgeDocumentSet](https://cloud.tencent.com/document/api/1301/68878#QureyKnowledgeDocumentSet)

	* 新增成员：DocStatus, ErrMsg, FileId




## 腾讯云可观测平台(monitor) 版本：2023-06-16



## 腾讯云可观测平台(monitor) 版本：2018-07-24

### 第 138 次发布

发布时间：2025-06-05 01:47:47

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateExternalCluster](https://cloud.tencent.com/document/api/248/118983)
* [DescribeExternalClusterUninstallCommand](https://cloud.tencent.com/document/api/248/118982)



## 流计算 Oceanus(oceanus) 版本：2019-04-22

### 第 69 次发布

发布时间：2025-06-05 01:50:56

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeJobEvents](https://cloud.tencent.com/document/api/849/102554)

	* 新增出参：Versions




## 私有域解析 Private DNS(privatedns) 版本：2020-10-28

### 第 25 次发布

发布时间：2025-06-05 01:53:52

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeRecord](https://cloud.tencent.com/document/api/1338/118984)

新增数据结构：

* [RecordInfo](https://cloud.tencent.com/document/api/1338/55947#RecordInfo)



## 容器安全服务(tcss) 版本：2020-11-01

### 第 75 次发布

发布时间：2025-06-05 02:00:59

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DeleteMachine](https://cloud.tencent.com/document/api/1285/81687)

	* 新增入参：NodeUniqueIds, UUIDs

* [DescribeAccessControlEvents](https://cloud.tencent.com/document/api/1285/65545)

	* 新增出参：SupportCoreVersion, InterceptionFailureTip

* [DescribeAssetImageRegistryListExport](https://cloud.tencent.com/document/api/1285/65499)

	* <font color="#dd0000">**修改入参**：</font>ExportField

	* 新增出参：JobId

* [DescribeAssetImageRegistryVirusListExport](https://cloud.tencent.com/document/api/1285/65492)

	* <font color="#dd0000">**修改入参**：</font>ExportField

	* 新增出参：JobId

* [DescribeAssetImageRegistryVulListExport](https://cloud.tencent.com/document/api/1285/65490)

	* <font color="#dd0000">**修改入参**：</font>ExportField

	* 新增出参：JobId

* [DescribeAssetSyncLastTime](https://cloud.tencent.com/document/api/1285/81683)

	* 新增出参：FailedHostCount, TaskId

* [DescribeClusterDetail](https://cloud.tencent.com/document/api/1285/65448)

	* 新增出参：OwnerName

* [DescribeVirusDetail](https://cloud.tencent.com/document/api/1285/65583)

	* 新增出参：ContainerStatus


修改数据结构：

* [ClusterInfoItem](https://cloud.tencent.com/document/api/1285/65614#ClusterInfoItem)

	* 新增成员：OwnerName

* [RiskDnsEventInfo](https://cloud.tencent.com/document/api/1285/65614#RiskDnsEventInfo)

	* 新增成员：ImageId, ContainerId

* [RunTimeEventBaseInfo](https://cloud.tencent.com/document/api/1285/65614#RunTimeEventBaseInfo)

	* 新增成员：ContainerStatus

* [VirusInfo](https://cloud.tencent.com/document/api/1285/65614#VirusInfo)

	* 新增成员：HostIP




