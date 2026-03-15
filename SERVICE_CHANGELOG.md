# Release v1.3.57

## TDSQL-C MySQL 版(cynosdb) 版本：2019-01-07

### 第 156 次发布

发布时间：2026-03-16 01:38:40

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyServerlessStrategy](https://cloud.tencent.com/document/api/1003/110600)

	* 新增入参：SecurityGroupIdsForNewRo


新增数据结构：

* [CreateBackupVaultItem](https://cloud.tencent.com/document/api/1003/48097#CreateBackupVaultItem)
* [VaultInfo](https://cloud.tencent.com/document/api/1003/48097#VaultInfo)

修改数据结构：

* [BackupConfigInfo](https://cloud.tencent.com/document/api/1003/48097#BackupConfigInfo)

	* 新增成员：AutoCopyVaults

* [BackupFileInfo](https://cloud.tencent.com/document/api/1003/48097#BackupFileInfo)

	* 新增成员：CopyStatus, EncryptKeyId, EncryptRegion, VaultInfos

* [BinlogConfigInfo](https://cloud.tencent.com/document/api/1003/48097#BinlogConfigInfo)

	* 新增成员：AutoCopyVaults

* [BinlogItem](https://cloud.tencent.com/document/api/1003/48097#BinlogItem)

	* 新增成员：CopyStatus, VaultInfos, EncryptKeyId, EncryptRegion

* [BizTaskInfo](https://cloud.tencent.com/document/api/1003/48097#BizTaskInfo)

	* 新增成员：VaultId, VaultName

* [LogicBackupConfigInfo](https://cloud.tencent.com/document/api/1003/48097#LogicBackupConfigInfo)

	* 新增成员：AutoCopyVaults

* [RedoLogItem](https://cloud.tencent.com/document/api/1003/48097#RedoLogItem)

	* 新增成员：VaultInfos, CopyStatus, EncryptKeyId, EncryptRegion

* [SnapshotBackupConfig](https://cloud.tencent.com/document/api/1003/48097#SnapshotBackupConfig)

	* 新增成员：AutoCopyVaults




## 数据加速器 GooseFS(goosefs) 版本：2022-05-19

### 第 17 次发布

发布时间：2026-03-16 01:57:34

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [BuildCustomerCluster](https://cloud.tencent.com/document/api/1424/129191)
* [DeleteCustomerCluster](https://cloud.tencent.com/document/api/1424/129190)
* [DescribeCustomerCluster](https://cloud.tencent.com/document/api/1424/129189)
* [MountMultipleStorageFileSystem](https://cloud.tencent.com/document/api/1424/129188)
* [QueryClientNodeMountCommand](https://cloud.tencent.com/document/api/1424/129187)

新增数据结构：

* [ClusterMountAttr](https://cloud.tencent.com/document/api/1424/95076#ClusterMountAttr)
* [CustomerClusterAttr](https://cloud.tencent.com/document/api/1424/95076#CustomerClusterAttr)



## TDSQL(tdmysql) 版本：2021-11-22

### 第 2 次发布

发布时间：2026-03-16 03:11:13

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除接口**：</font>

* DescribeBillingEnable
* DescribeDatabaseTable
* ModifyBinlogStatus

<font color="#dd0000">**删除数据结构**：</font>

* TableColumn



## 边缘安全加速平台(teo) 版本：2022-09-01

### 第 139 次发布

发布时间：2026-03-16 03:13:12

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [EnableOriginACL](https://cloud.tencent.com/document/api/1552/120406)

	* 新增入参：OriginACLFamily

* [ModifyOriginACL](https://cloud.tencent.com/document/api/1552/120405)

	* 新增入参：OriginACLFamily


修改数据结构：

* [OriginACLInfo](https://cloud.tencent.com/document/api/1552/80721#OriginACLInfo)

	* 新增成员：OriginACLFamily




## 边缘安全加速平台(teo) 版本：2022-01-06



## 云点播(vod) 版本：2024-07-18



## 云点播(vod) 版本：2018-07-17

### 第 237 次发布

发布时间：2026-03-16 03:39:38

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateAigcSubject](https://cloud.tencent.com/document/api/266/129192)

修改接口：

* [CreateAigcVideoTask](https://cloud.tencent.com/document/api/266/126239)

	* 新增入参：SubjectInfos

* [DescribeTaskDetail](https://cloud.tencent.com/document/api/266/33431)

	* 新增出参：CreateAigcSubjectTask

* [ProcessImageAsync](https://cloud.tencent.com/document/api/266/127858)

	* 新增入参：Base64


新增数据结构：

* [AigcVideoTaskInputSubjectInfo](https://cloud.tencent.com/document/api/266/31773#AigcVideoTaskInputSubjectInfo)
* [CreateAigcSubjectInput](https://cloud.tencent.com/document/api/266/31773#CreateAigcSubjectInput)
* [CreateAigcSubjectOutput](https://cloud.tencent.com/document/api/266/31773#CreateAigcSubjectOutput)
* [CreateAigcSubjectTask](https://cloud.tencent.com/document/api/266/31773#CreateAigcSubjectTask)

修改数据结构：

* [AigcVideoTaskInput](https://cloud.tencent.com/document/api/266/31773#AigcVideoTaskInput)

	* 新增成员：SubjectInfos




