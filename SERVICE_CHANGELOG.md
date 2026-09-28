# Release v1.3.188

## Agent 沙箱服务(ags) 版本：2025-09-20

### 第 26 次发布

发布时间：2026-09-29 01:08:45

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreatePreCacheImageTask](https://cloud.tencent.com/document/api/1814/127508)

	* 新增出参：PreCacheImageId

* [DescribePreCacheImageTask](https://cloud.tencent.com/document/api/1814/127507)

	* 新增入参：PreCacheImageId

	* <font color="#dd0000">**修改入参**：</font>Image, ImageRegistryType, ImageDigest

	* 新增出参：CreateTime, PreCacheImageId, SourceType, CachedImageSizeBytes, LastUsedTime




## 腾讯云数据仓库 TCHouse-D(cdwdoris) 版本：2021-12-28

### 第 65 次发布

发布时间：2026-09-29 01:24:49

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateInstanceNew](https://cloud.tencent.com/document/api/1387/102611)

	* 新增入参：DiskEncrypt




## 腾讯云数据仓库TCHouse-P(cdwpg) 版本：2020-12-30

### 第 17 次发布

发布时间：2026-09-29 01:25:29

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeInstanceState](https://cloud.tencent.com/document/api/878/100160)

	* 新增入参：InstanceIds

	* <font color="#dd0000">**修改入参**：</font>InstanceId

	* 新增出参：InstanceStates


新增数据结构：

* [InstanceStateItem](https://cloud.tencent.com/document/api/878/98895#InstanceStateItem)



## 云加密机(cloudhsm) 版本：2019-11-12

### 第 12 次发布

发布时间：2026-09-29 01:32:24

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeVsmAttributes](https://cloud.tencent.com/document/api/639/41444)

	* 新增出参：ClusterId, ClusterRole

* [DescribeVsms](https://cloud.tencent.com/document/api/639/41443)

	* 新增入参：ClusterId

* [GetVsmMonitorInfo](https://cloud.tencent.com/document/api/639/88095)

	* 新增出参：DigestList, InitStatus


新增数据结构：

* [VsmDigestItem](https://cloud.tencent.com/document/api/639/41450#VsmDigestItem)

修改数据结构：

* [ResourceInfo](https://cloud.tencent.com/document/api/639/41450#ResourceInfo)

	* 新增成员：Version, ClusterId, ClusterRole




## 云安全一体化平台(csip) 版本：2022-11-21

### 第 106 次发布

发布时间：2026-09-29 01:37:23

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [NotifyAssetConfigItem](https://cloud.tencent.com/document/api/664/90825#NotifyAssetConfigItem)

	* 新增成员：ProjectIds

* [WebhookAssetScope](https://cloud.tencent.com/document/api/664/90825#WebhookAssetScope)

	* 新增成员：ProjectIds




## 大数据智能体工作台DataBuddy(databuddy) 版本：2026-07-15

### 第 5 次发布

发布时间：2026-09-29 01:53:21

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateFolder](https://cloud.tencent.com/document/api/1835/138899)
* [CreateWorkspace](https://cloud.tencent.com/document/api/1835/138904)
* [DeleteFolder](https://cloud.tencent.com/document/api/1835/138898)
* [DeleteWorkspace](https://cloud.tencent.com/document/api/1835/138903)
* [GetFolder](https://cloud.tencent.com/document/api/1835/138897)
* [GetWorkspace](https://cloud.tencent.com/document/api/1835/138902)
* [ListFiles](https://cloud.tencent.com/document/api/1835/138896)
* [UpdateFolder](https://cloud.tencent.com/document/api/1835/138895)
* [UpdateWorkspace](https://cloud.tencent.com/document/api/1835/138901)

新增数据结构：

* [CreateFolderRsp](https://cloud.tencent.com/document/api/1835/138006#CreateFolderRsp)
* [CreateWorkspaceRsp](https://cloud.tencent.com/document/api/1835/138006#CreateWorkspaceRsp)
* [DeleteFolderRsp](https://cloud.tencent.com/document/api/1835/138006#DeleteFolderRsp)
* [DeleteWorkspaceRsp](https://cloud.tencent.com/document/api/1835/138006#DeleteWorkspaceRsp)
* [FileMeta](https://cloud.tencent.com/document/api/1835/138006#FileMeta)
* [FileNode](https://cloud.tencent.com/document/api/1835/138006#FileNode)
* [FolderLocator](https://cloud.tencent.com/document/api/1835/138006#FolderLocator)
* [GetFolderRsp](https://cloud.tencent.com/document/api/1835/138006#GetFolderRsp)
* [GetWorkspaceRsp](https://cloud.tencent.com/document/api/1835/138006#GetWorkspaceRsp)
* [GitRepoConfig](https://cloud.tencent.com/document/api/1835/138006#GitRepoConfig)
* [ListFilesRsp](https://cloud.tencent.com/document/api/1835/138006#ListFilesRsp)
* [SparseCheckoutConfig](https://cloud.tencent.com/document/api/1835/138006#SparseCheckoutConfig)
* [StandardUserInfo](https://cloud.tencent.com/document/api/1835/138006#StandardUserInfo)
* [UpdateFolderRsp](https://cloud.tencent.com/document/api/1835/138006#UpdateFolderRsp)
* [UpdateWorkspaceRsp](https://cloud.tencent.com/document/api/1835/138006#UpdateWorkspaceRsp)
* [UserInfo](https://cloud.tencent.com/document/api/1835/138006#UserInfo)
* [WorkspaceInfo](https://cloud.tencent.com/document/api/1835/138006#WorkspaceInfo)



## 数据库智能管家 DBbrain(dbbrain) 版本：2021-05-27

### 第 65 次发布

发布时间：2026-09-29 01:54:17

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeDeadLockLogs](https://cloud.tencent.com/document/api/1130/138906)

新增数据结构：

* [DeadLockLogItem](https://cloud.tencent.com/document/api/1130/57812#DeadLockLogItem)
* [DeadlockFrame](https://cloud.tencent.com/document/api/1130/57812#DeadlockFrame)
* [DeadlockResource](https://cloud.tencent.com/document/api/1130/57812#DeadlockResource)
* [DeadlockSession](https://cloud.tencent.com/document/api/1130/57812#DeadlockSession)
* [DeadlockTransaction](https://cloud.tencent.com/document/api/1130/57812#DeadlockTransaction)
* [OwnerItem](https://cloud.tencent.com/document/api/1130/57812#OwnerItem)
* [WaiterItem](https://cloud.tencent.com/document/api/1130/57812#WaiterItem)



## 数据库智能管家 DBbrain(dbbrain) 版本：2019-10-16



## 弹性 MapReduce(emr) 版本：2019-01-03

### 第 159 次发布

发布时间：2026-09-29 02:07:12

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateCloudInstance](https://cloud.tencent.com/document/api/589/113701)

	* 新增入参：AirflowDagSource, AirflowGitCredential


新增数据结构：

* [AirflowCfsSource](https://cloud.tencent.com/document/api/589/33981#AirflowCfsSource)
* [AirflowDagSourceInput](https://cloud.tencent.com/document/api/589/33981#AirflowDagSourceInput)
* [AirflowGitCredentialInput](https://cloud.tencent.com/document/api/589/33981#AirflowGitCredentialInput)
* [AirflowGitSource](https://cloud.tencent.com/document/api/589/33981#AirflowGitSource)



## 图片内容安全(ims) 版本：2020-12-29

### 第 15 次发布

发布时间：2026-09-29 02:18:42

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ImageModeration](https://cloud.tencent.com/document/api/1125/53273)

	* 新增出参：StoreUrl, Reason




## 图片内容安全(ims) 版本：2020-07-13



## 媒体处理(mps) 版本：2019-06-12

### 第 252 次发布

发布时间：2026-09-29 02:37:37

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [AiComposeConfig](https://cloud.tencent.com/document/api/862/37615#AiComposeConfig)
* [ImageComposeCanvas](https://cloud.tencent.com/document/api/862/37615#ImageComposeCanvas)
* [ImageComposeLayer](https://cloud.tencent.com/document/api/862/37615#ImageComposeLayer)

修改数据结构：

* [ImageTaskInput](https://cloud.tencent.com/document/api/862/37615#ImageTaskInput)

	* 新增成员：AiComposeConfig




## 文字识别(ocr) 版本：2018-11-19

### 第 269 次发布

发布时间：2026-09-29 02:42:30

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [RecognizeThaiIDCardOCR](https://cloud.tencent.com/document/api/866/48475)

	* 新增出参：ThaiFirstName, ThaiLastName




## 云数据库 SQL Server(sqlserver) 版本：2018-03-28

### 第 80 次发布

发布时间：2026-09-29 02:52:15

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateExportTask](https://cloud.tencent.com/document/api/238/129961)

	* 新增出参：FileName


修改数据结构：

* [ExportFile](https://cloud.tencent.com/document/api/238/19976#ExportFile)

	* 新增成员：LogStartTime, LogEndTime, LogFilter

* [LogResult](https://cloud.tencent.com/document/api/238/19976#LogResult)

	* 新增成员：EventName




## 云开发 CloudBase(tcb) 版本：2018-06-08

### 第 170 次发布

发布时间：2026-09-29 02:56:41

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreatePlatformHTTPServiceRoute](https://cloud.tencent.com/document/api/876/138911)
* [DeletePlatformHTTPServiceRoute](https://cloud.tencent.com/document/api/876/138910)
* [DescribePlatformHTTPServiceRoute](https://cloud.tencent.com/document/api/876/138909)
* [ModifyPlatformHTTPServiceRoute](https://cloud.tencent.com/document/api/876/138908)
* [VerifyPlatformHTTPServiceRoute](https://cloud.tencent.com/document/api/876/138907)

修改接口：

* [CreateCloudApp](https://cloud.tencent.com/document/api/876/135281)

	* 新增入参：Trigger, ServiceList, WorkingDir, Routes, PromoteType, ClientToken, PreDeployCommand, PostDeployCommand

* [DescribeCloudAppInfo](https://cloud.tencent.com/document/api/876/135277)

	* 新增出参：BuildConfig, CurrentVersion, PreviewDomain

* [DescribeCloudAppList](https://cloud.tencent.com/document/api/876/132936)

	* 新增入参：Filter

* [DescribeCloudAppVersion](https://cloud.tencent.com/document/api/876/135276)

	* 新增出参：Snapshot, TrafficPercent, VersionDomain, Resources, Artifacts


新增数据结构：

* [BuildArtifactInfo](https://cloud.tencent.com/document/api/876/34822#BuildArtifactInfo)
* [BuildContext](https://cloud.tencent.com/document/api/876/34822#BuildContext)
* [CloudAppFilter](https://cloud.tencent.com/document/api/876/34822#CloudAppFilter)
* [CloudAppLinkService](https://cloud.tencent.com/document/api/876/34822#CloudAppLinkService)
* [CloudAppResourceItem](https://cloud.tencent.com/document/api/876/34822#CloudAppResourceItem)
* [CloudAppRoute](https://cloud.tencent.com/document/api/876/34822#CloudAppRoute)
* [CloudAppTrigger](https://cloud.tencent.com/document/api/876/34822#CloudAppTrigger)
* [CloudAppWebHook](https://cloud.tencent.com/document/api/876/34822#CloudAppWebHook)

修改数据结构：

* [BuildSource](https://cloud.tencent.com/document/api/876/34822#BuildSource)

	* 新增成员：PackageFileName

* [CloudAppServiceItem](https://cloud.tencent.com/document/api/876/34822#CloudAppServiceItem)

	* 新增成员：BuildConfig, CurrentVersion

* [CloudAppVersionItem](https://cloud.tencent.com/document/api/876/34822#CloudAppVersionItem)

	* 新增成员：Snapshot, VersionDomain, TrafficPercent, Resources, Artifacts




