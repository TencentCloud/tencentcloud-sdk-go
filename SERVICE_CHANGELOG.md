# Release v1.3.173

## 弹性伸缩(as) 版本：2018-04-19

### 第 93 次发布

发布时间：2026-09-04 01:10:37

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ClearLaunchConfigurationAttributes](https://cloud.tencent.com/document/api/377/54255)

	* 新增入参：ClearNetworkInterfaces

* [CreateLaunchConfiguration](https://cloud.tencent.com/document/api/377/20447)

	* 新增入参：NetworkInterfaces

* [ModifyLaunchConfigurationAttributes](https://cloud.tencent.com/document/api/377/31298)

	* 新增入参：NetworkInterfaces


新增数据结构：

* [NetworkInterface](https://cloud.tencent.com/document/api/377/20453#NetworkInterface)

修改数据结构：

* [DataDisk](https://cloud.tencent.com/document/api/377/20453#DataDisk)

	* 新增成员：KmsKeyId

* [LaunchConfiguration](https://cloud.tencent.com/document/api/377/20453#LaunchConfiguration)

	* 新增成员：NetworkInterfaces

* [SystemDisk](https://cloud.tencent.com/document/api/377/20453#SystemDisk)

	* 新增成员：Encrypt, KmsKeyId




## 费用中心(billing) 版本：2018-07-09

### 第 95 次发布

发布时间：2026-09-04 01:12:35

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeAccountWarning](https://cloud.tencent.com/document/api/555/137478)
* [ModifyAccountWarning](https://cloud.tencent.com/document/api/555/137477)



## 负载均衡(clb) 版本：2018-03-17

### 第 162 次发布

发布时间：2026-09-04 01:19:07

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyModelAttributes](https://cloud.tencent.com/document/api/214/133670)

	* 新增入参：ApiBases




## 数据湖计算 DLC(dlc) 版本：2021-01-25

### 第 179 次发布

发布时间：2026-09-04 01:28:06

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [LakeFsInfo](https://cloud.tencent.com/document/api/1342/53778#LakeFsInfo)

	* 新增成员：MultiAZ, Configuration

* [SparkSessionInfo](https://cloud.tencent.com/document/api/1342/53778#SparkSessionInfo)

	* 新增成员：ApplicationId, ApplicationStartTime




## 数字版权管理(drm) 版本：2018-11-15

### 第 8 次发布

发布时间：2026-09-04 01:30:35

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [PlaybackPolicy](https://cloud.tencent.com/document/api/1000/30712#PlaybackPolicy)

	* 新增成员：CanPersistent




## 弹性 MapReduce(emr) 版本：2019-01-03

### 第 155 次发布

发布时间：2026-09-04 01:32:45

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeDynamicInstanceDetail](https://cloud.tencent.com/document/api/589/132013)

	* 新增出参：EnableHistoryServer, TensorBoardUrl

* [DescribeNodeSpec](https://cloud.tencent.com/document/api/589/120504)

	* 新增出参：Architectures

* [InquiryPriceScaleOutInstance](https://cloud.tencent.com/document/api/589/34265)

	* 新增入参：NodeGroupId


新增数据结构：

* [ArchitectureInfo](https://cloud.tencent.com/document/api/589/33981#ArchitectureInfo)

修改数据结构：

* [ModifyDynamicInstanceForm](https://cloud.tencent.com/document/api/589/33981#ModifyDynamicInstanceForm)

	* 新增成员：EnableHistoryServer

* [NodeSpecInstanceType](https://cloud.tencent.com/document/api/589/33981#NodeSpecInstanceType)

	* 新增成员：GpuResourceKey, GpuNum

* [RayCluster](https://cloud.tencent.com/document/api/589/33981#RayCluster)

	* 新增成员：StorageCount




## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 319 次发布

发布时间：2026-09-04 01:33:54

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [ApproverOption](https://cloud.tencent.com/document/api/1323/70369#ApproverOption)

	* 新增成员：AddSignComponentUseSealSize




## 腾讯电子签（基础版）(essbasic) 版本：2021-05-26

### 第 275 次发布

发布时间：2026-09-04 01:34:46

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [ApproverOption](https://cloud.tencent.com/document/api/1420/61525#ApproverOption)

	* 新增成员：AddSignComponentUseSealSize




## 腾讯电子签（基础版）(essbasic) 版本：2020-12-22



## iOA 零信任安全管理系统(ioa) 版本：2022-06-01

### 第 43 次发布

发布时间：2026-09-04 01:38:11

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DeleteAccountGroup](https://cloud.tencent.com/document/api/1092/137480)

修改接口：

* [CreateCompanyDirectoryConfig](https://cloud.tencent.com/document/api/1092/131718)

	* 新增入参：NameI18n

* [DescribeAccountGroups](https://cloud.tencent.com/document/api/1092/107711)

	* 新增入参：DomainInstanceId

* [ModifyCompanyDirectoryConfig](https://cloud.tencent.com/document/api/1092/131716)

	* 新增入参：NameI18n


新增数据结构：

* [I18nString](https://cloud.tencent.com/document/api/1092/102488#I18nString)

修改数据结构：

* [DescribeDLPEdgeNodeGroupsRspItem](https://cloud.tencent.com/document/api/1092/102488#DescribeDLPEdgeNodeGroupsRspItem)

	* 新增成员：GroupNameI18n

* [DescribeDeviceDetailListData](https://cloud.tencent.com/document/api/1092/102488#DescribeDeviceDetailListData)

	* <font color="#dd0000">**修改成员**：</font>UserName, ComputerName, Name, AccountGroupIdPath, AccountGroupId, GroupNamePath, Ip, AccountGroupName, GroupIdPath, Mid, IoaUserName, GroupId, GroupName, Mac, Version, AccountGroupNamePath, Id

* [DirectoryConfigData](https://cloud.tencent.com/document/api/1092/102488#DirectoryConfigData)

	* 新增成员：NameI18n

* [DirectoryConfigResultData](https://cloud.tencent.com/document/api/1092/102488#DirectoryConfigResultData)

	* 新增成员：NameI18n




## 云数据库 MongoDB(mongodb) 版本：2019-07-25

### 第 76 次发布

发布时间：2026-09-04 01:45:09

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeDBInstanceNodeProperty](https://cloud.tencent.com/document/api/240/82022)

	* 新增出参：DynamoProxies




## 云数据库 MongoDB(mongodb) 版本：2018-04-08



## 媒体处理(mps) 版本：2019-06-12

### 第 244 次发布

发布时间：2026-09-04 01:46:37

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateAigcAudioTask](https://cloud.tencent.com/document/api/862/132830)

	* <font color="#dd0000">**删除入参**：</font>OutputAudioFormat

* [CreateAigcVideoTask](https://cloud.tencent.com/document/api/862/126965)

	* 新增入参：SubjectInfos


新增数据结构：

* [AigcVideoReferenceSubjectInfo](https://cloud.tencent.com/document/api/862/37615#AigcVideoReferenceSubjectInfo)

修改数据结构：

* [DescribeOutput](https://cloud.tencent.com/document/api/862/37615#DescribeOutput)

	* 新增成员：State




## TDSQL(tdmysql) 版本：2021-11-22

### 第 15 次发布

发布时间：2026-09-04 01:59:02

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [BreakStandbyDBInstanceRelation](https://cloud.tencent.com/document/api/1376/137484)
* [CreateStandbyDBInstance](https://cloud.tencent.com/document/api/1376/137483)
* [DescribeStandbyDBInstanceRelationDetail](https://cloud.tencent.com/document/api/1376/137482)

新增数据结构：

* [StandbyDBInstanceRelation](https://cloud.tencent.com/document/api/1376/128305#StandbyDBInstanceRelation)



## 高性能计算平台(thpc) 版本：2023-03-21

### 第 44 次发布

发布时间：2026-09-04 02:00:50

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [BindClusterVpc](https://cloud.tencent.com/document/api/1527/137489)
* [DescribeClusterDedicatedProxy](https://cloud.tencent.com/document/api/1527/137488)
* [DisableClusterDedicatedProxy](https://cloud.tencent.com/document/api/1527/137487)
* [EnableClusterDedicatedProxy](https://cloud.tencent.com/document/api/1527/137486)
* [GenerateRegisterCode](https://cloud.tencent.com/document/api/1527/137490)
* [GenerateRegisterCommand](https://cloud.tencent.com/document/api/1527/137485)



## 高性能计算平台(thpc) 版本：2022-04-01



## 高性能计算平台(thpc) 版本：2021-11-09



## TokenHub(tokenhub) 版本：2026-03-22

### 第 21 次发布

发布时间：2026-09-04 02:03:15

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [EndpointDetail](https://cloud.tencent.com/document/api/1823/132279#EndpointDetail)

	* 新增成员：ModelStatus




## 云点播(vod) 版本：2024-07-18



## 云点播(vod) 版本：2018-07-17

### 第 285 次发布

发布时间：2026-09-04 02:07:12

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CloneVoiceSync](https://cloud.tencent.com/document/api/266/137492)
* [TextToSpeechSync](https://cloud.tencent.com/document/api/266/137491)

新增数据结构：

* [TextToSpeechSyncOutputOption](https://cloud.tencent.com/document/api/266/31773#TextToSpeechSyncOutputOption)



