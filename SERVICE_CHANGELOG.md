# Release v1.0.1205

## 弹性伸缩(as) 版本：2018-04-19

### 第 85 次发布

发布时间：2025-07-11 01:08:53

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [RefreshSettings](https://cloud.tencent.com/document/api/377/20453#RefreshSettings)

	* 新增成员：CheckInstanceTargetHealthTimeout




## 腾讯云数据仓库 TCHouse-D(cdwdoris) 版本：2021-12-28

### 第 53 次发布

发布时间：2025-07-11 01:13:34

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateBackUpSchedule](https://cloud.tencent.com/document/api/1387/109543)

	* 新增入参：SnapshotRemainPolicy, DataRemoteRegion

* [DescribeSqlApis](https://cloud.tencent.com/document/api/1387/109441)

	* 新增入参：UserNames

* [ModifyUserPrivilegesV3](https://cloud.tencent.com/document/api/1387/109432)

	* 新增入参：DefaultComputeGroup


新增数据结构：

* [SnapshotRemainPolicy](https://cloud.tencent.com/document/api/1387/102385#SnapshotRemainPolicy)

修改数据结构：

* [BackUpJobDisplay](https://cloud.tencent.com/document/api/1387/102385#BackUpJobDisplay)

	* 新增成员：SnapshotRemainPolicy

* [BackupCosInfo](https://cloud.tencent.com/document/api/1387/102385#BackupCosInfo)

	* 新增成员：Region

* [InstanceNode](https://cloud.tencent.com/document/api/1387/102385#InstanceNode)

	* 新增成员：VirtualZone

* [NodeInfo](https://cloud.tencent.com/document/api/1387/102385#NodeInfo)

	* 新增成员：VirtualZone

* [NodeInfos](https://cloud.tencent.com/document/api/1387/102385#NodeInfos)

	* 新增成员：VirtualZone

* [RestoreStatus](https://cloud.tencent.com/document/api/1387/102385#RestoreStatus)

	* 新增成员：ID




## 混沌演练平台(cfg) 版本：2021-08-20

### 第 30 次发布

发布时间：2025-07-11 01:14:03

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [Template](https://cloud.tencent.com/document/api/1500/71784#Template)

	* 新增成员：TemplateScenario, TemplatePurpose




## 云游戏(gs) 版本：2019-11-18

### 第 49 次发布

发布时间：2025-07-11 01:28:18

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateAndroidInstanceImage](https://cloud.tencent.com/document/api/1162/117236)

	* 新增入参：AndroidInstanceImageDescription

* [DescribeAndroidInstanceImages](https://cloud.tencent.com/document/api/1162/117234)

	* 新增入参：Filters


修改数据结构：

* [AndroidInstanceImage](https://cloud.tencent.com/document/api/1162/40743#AndroidInstanceImage)

	* 新增成员：AndroidInstanceImageDescription, CreateTime




## 智能全局流量管理(igtm) 版本：2023-10-24

### 第 2 次发布

发布时间：2025-07-11 01:29:22

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeDetectPackageDetail](https://cloud.tencent.com/document/api/1551/121139)
* [DescribeDetectTaskPackageList](https://cloud.tencent.com/document/api/1551/121138)
* [DescribeInstancePackageList](https://cloud.tencent.com/document/api/1551/121137)

新增数据结构：

* [CostItem](https://cloud.tencent.com/document/api/1551/120465#CostItem)
* [DetectTaskPackage](https://cloud.tencent.com/document/api/1551/120465#DetectTaskPackage)
* [InstancePackage](https://cloud.tencent.com/document/api/1551/120465#InstancePackage)



## 物联网开发平台(iotexplorer) 版本：2019-04-23

### 第 109 次发布

发布时间：2025-07-11 01:30:01

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**预下线接口**：</font>

* CancelAssignTWeCallLicense



## 腾讯健康组学平台(omics) 版本：2022-11-28

### 第 21 次发布

发布时间：2025-07-11 01:39:59

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [ClusterOption](https://cloud.tencent.com/document/api/1643/89100#ClusterOption)

	* 新增成员：SystemNodeInstanceType, SystemNodeCount

* [ResourceIds](https://cloud.tencent.com/document/api/1643/89100#ResourceIds)

	* 新增成员：TKEId, TKESystemNodePoolId




## 邮件推送(ses) 版本：2020-10-02

### 第 33 次发布

发布时间：2025-07-11 01:42:46

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateEmailIdentity](https://cloud.tencent.com/document/api/1288/51048)

	* 新增入参：TagList

* [ListEmailIdentities](https://cloud.tencent.com/document/api/1288/51045)

	* 新增入参：TagList, Limit, Offset

	* 新增出参：Total


新增数据结构：

* [TagList](https://cloud.tencent.com/document/api/1288/51053#TagList)

修改数据结构：

* [EmailIdentity](https://cloud.tencent.com/document/api/1288/51053#EmailIdentity)

	* 新增成员：TagList




## 消息队列 TDMQ(tdmq) 版本：2020-02-17

### 第 153 次发布

发布时间：2025-07-11 01:47:41

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [PulsarProClusterInfo](https://cloud.tencent.com/document/api/1179/46089#PulsarProClusterInfo)

	* 新增成员：DeleteProtection




## 边缘安全加速平台(teo) 版本：2022-09-01

### 第 110 次发布

发布时间：2025-07-10 14:38:08

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DeleteWebSecurityTemplate](https://cloud.tencent.com/document/api/1552/121064)

### 第 109 次发布

发布时间：2025-07-10 14:17:53

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateWebSecurityTemplate](https://cloud.tencent.com/document/api/1552/121063)
* [DescribeWebSecurityTemplate](https://cloud.tencent.com/document/api/1552/121062)
* [DescribeWebSecurityTemplates](https://cloud.tencent.com/document/api/1552/121061)
* [ModifyWebSecurityTemplate](https://cloud.tencent.com/document/api/1552/121060)

新增数据结构：

* [BindDomainInfo](https://cloud.tencent.com/document/api/1552/80721#BindDomainInfo)
* [SecurityPolicyTemplateInfo](https://cloud.tencent.com/document/api/1552/80721#SecurityPolicyTemplateInfo)



## 边缘安全加速平台(teo) 版本：2022-01-06



