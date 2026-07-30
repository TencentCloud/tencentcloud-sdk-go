# Release v1.3.148

## 云数据库 MySQL(cdb) 版本：2017-03-20

### 第 225 次发布

发布时间：2026-07-30 01:20:26

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [UpgradeRoGroup](https://cloud.tencent.com/document/api/236/135261)



## TDSQL-C MySQL 版(cynosdb) 版本：2019-01-07

### 第 182 次发布

发布时间：2026-07-30 01:44:28

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeClusterLevels](https://cloud.tencent.com/document/api/1003/134025)

	* <font color="#dd0000">**修改入参**：</font>Zone

	* 新增出参：Zones




## 数据库智能管家 DBbrain(dbbrain) 版本：2021-05-27

### 第 59 次发布

发布时间：2026-07-30 01:47:23

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateDBDiagReportUrls](https://cloud.tencent.com/document/api/1130/135264)
* [DescribeDBInstances](https://cloud.tencent.com/document/api/1130/135262)

修改接口：

* [DescribeDBDiagReportTasks](https://cloud.tencent.com/document/api/1130/57805)

	* 新增入参：TagFilters


新增数据结构：

* [DiagReportUrlItem](https://cloud.tencent.com/document/api/1130/57812#DiagReportUrlItem)
* [InstanceItem](https://cloud.tencent.com/document/api/1130/57812#InstanceItem)
* [TagFilterGroup](https://cloud.tencent.com/document/api/1130/57812#TagFilterGroup)
* [TagInfo](https://cloud.tencent.com/document/api/1130/57812#TagInfo)
* [TagPair](https://cloud.tencent.com/document/api/1130/57812#TagPair)

修改数据结构：

* [HealthReportTask](https://cloud.tencent.com/document/api/1130/57812#HealthReportTask)

	* 新增成员：Tags




## 数据库智能管家 DBbrain(dbbrain) 版本：2019-10-16



## 云数据库独享集群(dbdc) 版本：2020-10-29

### 第 9 次发布

发布时间：2026-07-30 01:48:21

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeDBCustomClusterNodeConfig](https://cloud.tencent.com/document/api/1322/135268)
* [DescribeDBCustomClusterNodeResources](https://cloud.tencent.com/document/api/1322/135267)
* [DescribeDBCustomClusterResources](https://cloud.tencent.com/document/api/1322/135266)
* [DescribeDBCustomNodeSecurityGroups](https://cloud.tencent.com/document/api/1322/135270)
* [DescribeDBCustomNodeTypes](https://cloud.tencent.com/document/api/1322/135273)
* [DescribeDBCustomRegions](https://cloud.tencent.com/document/api/1322/135272)
* [DescribeDBCustomZones](https://cloud.tencent.com/document/api/1322/135271)
* [ModifyDBCustomClusterNodeConfig](https://cloud.tencent.com/document/api/1322/135265)
* [ModifyDBCustomNodeSecurityGroups](https://cloud.tencent.com/document/api/1322/135269)

新增数据结构：

* [DBCustomClusterNodeConfig](https://cloud.tencent.com/document/api/1322/74754#DBCustomClusterNodeConfig)
* [DBCustomClusterNodeResource](https://cloud.tencent.com/document/api/1322/74754#DBCustomClusterNodeResource)
* [DBCustomNodeTypeInfo](https://cloud.tencent.com/document/api/1322/74754#DBCustomNodeTypeInfo)
* [MetaResource](https://cloud.tencent.com/document/api/1322/74754#MetaResource)
* [PolicyRule](https://cloud.tencent.com/document/api/1322/74754#PolicyRule)
* [RegionInfo](https://cloud.tencent.com/document/api/1322/74754#RegionInfo)
* [SecurityGroup](https://cloud.tencent.com/document/api/1322/74754#SecurityGroup)
* [ZoneInfo](https://cloud.tencent.com/document/api/1322/74754#ZoneInfo)



## 数据湖计算 DLC(dlc) 版本：2021-01-25

### 第 168 次发布

发布时间：2026-07-30 01:49:59

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateSparkAppForTDLC](https://cloud.tencent.com/document/api/1342/135114)

	* 新增入参：AppName, AppType, DataEngine, AppFile, RoleArn, AppDriverSize, AppExecutorSize, AppExecutorNums, Eni, IsLocal, MainClass, AppConf, IsLocalJars, AppJars, IsLocalFiles, AppFiles, CmdArgs, MaxRetries, DataSource, IsLocalPythonFiles, AppPythonFiles, IsLocalArchives, AppArchives, SparkImage, SparkImageVersion, AppExecutorMaxNumbers, SessionId, IsInherit, IsSessionStarted, DependencyPackages




## iOA 零信任安全管理系统(ioa) 版本：2022-06-01

### 第 41 次发布

发布时间：2026-07-30 02:09:24

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [DescribeDeviceHardwareInfoItem](https://cloud.tencent.com/document/api/1092/102488#DescribeDeviceHardwareInfoItem)

	* 新增成员：BiosUuid

* [DeviceDetail](https://cloud.tencent.com/document/api/1092/102488#DeviceDetail)

	* 新增成员：BiosUuid




## 密钥管理系统(kms) 版本：2019-01-18

### 第 29 次发布

发布时间：2026-07-30 02:17:36

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [GetServiceStatus](https://cloud.tencent.com/document/api/573/34417)

	* 新增出参：ResourceZone, ResourceRegion


修改数据结构：

* [DataKeyMetadata](https://cloud.tencent.com/document/api/573/34431#DataKeyMetadata)

	* 新增成员：CreatorUinString

* [KeyMetadata](https://cloud.tencent.com/document/api/573/34431#KeyMetadata)

	* 新增成员：CreatorUinString




## 腾讯云可观测平台(monitor) 版本：2023-06-16



## 腾讯云可观测平台(monitor) 版本：2018-07-24

### 第 164 次发布

发布时间：2026-07-30 02:25:23

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeAlarmNoticeOnCallUsersFromPrometheusAlertID](https://cloud.tencent.com/document/api/248/135274)

新增数据结构：

* [NoticeOnCallUsersInfo](https://cloud.tencent.com/document/api/248/30354#NoticeOnCallUsersInfo)
* [NoticeSendGroup](https://cloud.tencent.com/document/api/248/30354#NoticeSendGroup)
* [NoticeUserInfo](https://cloud.tencent.com/document/api/248/30354#NoticeUserInfo)



## 云数据库Redis(redis) 版本：2018-04-12

### 第 105 次发布

发布时间：2026-07-30 02:36:05

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ChangeReplicaToMaster](https://cloud.tencent.com/document/api/239/56698)

	* 新增入参：Emergency




## SSL 证书(ssl) 版本：2019-12-05

### 第 99 次发布

发布时间：2026-07-30 02:41:31

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [UploadUpdateCertificateInstance](https://cloud.tencent.com/document/api/400/119791)

	* 新增入参：CertificateId


修改数据结构：

* [CertificateExtra](https://cloud.tencent.com/document/api/400/41679#CertificateExtra)

	* 新增成员：ServiceRenewCertificateId, ServiceOriginCertificateId

* [Certificates](https://cloud.tencent.com/document/api/400/41679#Certificates)

	* 新增成员：CertServiceShareEnabled, CertServiceValidCertificateCount




## 云开发 CloudBase(tcb) 版本：2018-06-08

### 第 154 次发布

发布时间：2026-07-30 02:44:22

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateCloudApp](https://cloud.tencent.com/document/api/876/135281)
* [DeleteCloudApp](https://cloud.tencent.com/document/api/876/135280)
* [DeleteCloudAppVersion](https://cloud.tencent.com/document/api/876/135279)
* [DescribeCloudAppCosInfo](https://cloud.tencent.com/document/api/876/135278)
* [DescribeCloudAppInfo](https://cloud.tencent.com/document/api/876/135277)
* [DescribeCloudAppVersion](https://cloud.tencent.com/document/api/876/135276)
* [DescribeCloudAppVersionList](https://cloud.tencent.com/document/api/876/135275)

新增数据结构：

* [BuildCommands](https://cloud.tencent.com/document/api/876/34822#BuildCommands)
* [BuildSecret](https://cloud.tencent.com/document/api/876/34822#BuildSecret)
* [BuildSource](https://cloud.tencent.com/document/api/876/34822#BuildSource)
* [BuildStep](https://cloud.tencent.com/document/api/876/34822#BuildStep)
* [BuildStepStatus](https://cloud.tencent.com/document/api/876/34822#BuildStepStatus)
* [CloudAppVersionItem](https://cloud.tencent.com/document/api/876/34822#CloudAppVersionItem)
* [StaticCmd](https://cloud.tencent.com/document/api/876/34822#StaticCmd)
* [StaticConfig](https://cloud.tencent.com/document/api/876/34822#StaticConfig)
* [StaticEnvironment](https://cloud.tencent.com/document/api/876/34822#StaticEnvironment)
* [Variable](https://cloud.tencent.com/document/api/876/34822#Variable)



## 边缘安全加速平台(teo) 版本：2022-09-01

### 第 153 次发布

发布时间：2026-07-30 02:54:10

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeIPGroupReferences](https://cloud.tencent.com/document/api/1552/135282)

新增数据结构：

* [IPGroupReference](https://cloud.tencent.com/document/api/1552/80721#IPGroupReference)

修改数据结构：

* [IPGroup](https://cloud.tencent.com/document/api/1552/80721#IPGroup)

	* 新增成员：RefCount




## 边缘安全加速平台(teo) 版本：2022-01-06



## 实时音视频(trtc) 版本：2019-07-22

### 第 148 次发布

发布时间：2026-07-30 03:04:09

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除接口**：</font>

* CreateBasicModeration
* DeleteBasicModeration

<font color="#dd0000">**删除数据结构**：</font>

* AuditStorageParams
* CloudAuditStorage



## 云点播(vod) 版本：2024-07-18



## 云点播(vod) 版本：2018-07-17

### 第 276 次发布

发布时间：2026-07-30 03:10:07

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [AdaptiveDynamicStreamingTemplate](https://cloud.tencent.com/document/api/266/31773#AdaptiveDynamicStreamingTemplate)

	* 新增成员：SegmentDuration




