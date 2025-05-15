# Release v1.0.1166

## 运维安全中心（堡垒机）(bh) 版本：2023-04-18

### 第 11 次发布

发布时间：2025-05-16 01:10:31

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeUsers](https://cloud.tencent.com/document/api/1025/74404)

	* 新增入参：IsCamUser, UserFromSet


新增数据结构：

* [IOAUserGroup](https://cloud.tencent.com/document/api/1025/74416#IOAUserGroup)

修改数据结构：

* [AssetSyncStatus](https://cloud.tencent.com/document/api/1025/74416#AssetSyncStatus)

	* 新增成员：ErrMsg

* [Device](https://cloud.tencent.com/document/api/1025/74416#Device)

	* 新增成员：IOAId

* [Resource](https://cloud.tencent.com/document/api/1025/74416#Resource)

	* 新增成员：IOAResource, PackageIOAUserCount, PackageIOABandwidth

* [User](https://cloud.tencent.com/document/api/1025/74416#User)

	* 新增成员：UserFrom, IOAUserGroup




## 商业智能分析 BI(bi) 版本：2022-01-05

### 第 26 次发布

发布时间：2025-05-16 01:11:06

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateEmbedToken](https://cloud.tencent.com/document/api/590/73724)

	* 新增入参：ConfigParam


修改数据结构：

* [EmbedTokenInfo](https://cloud.tencent.com/document/api/590/73726#EmbedTokenInfo)

	* 新增成员：ConfigParam




## 云数据库 MySQL(cdb) 版本：2017-03-20

### 第 191 次发布

发布时间：2025-05-16 01:14:00

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [Inbound](https://cloud.tencent.com/document/api/236/15878#Inbound)

	* 新增成员：Id

	* <font color="#dd0000">**修改成员**：</font>Action, CidrIp, PortRange, IpProtocol, Dir

* [RoInstanceInfo](https://cloud.tencent.com/document/api/236/15878#RoInstanceInfo)

	* 新增成员：ReplicationStatus




## TDSQL-C MySQL 版(cynosdb) 版本：2019-01-07

### 第 132 次发布

发布时间：2025-05-16 01:25:17

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [OpenAuditService](https://cloud.tencent.com/document/api/1003/84696)


修改数据结构：

* [AuditLogFile](https://cloud.tencent.com/document/api/1003/48097#AuditLogFile)

	* 新增成员：ProgressRate




## 弹性 MapReduce(emr) 版本：2019-01-03

### 第 100 次发布

发布时间：2025-05-16 01:32:47

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [UserInfoForUserManager](https://cloud.tencent.com/document/api/589/33981#UserInfoForUserManager)

	* 新增成员：Groups

* [UserManagerFilter](https://cloud.tencent.com/document/api/589/33981#UserManagerFilter)

	* 新增成员：Groups




## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 220 次发布

发布时间：2025-05-16 01:34:04

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateBatchSignUrl](https://cloud.tencent.com/document/api/1323/98670)

	* 新增入参：SignatureTypes




## 云数据库 PostgreSQL(postgres) 版本：2017-03-12

### 第 54 次发布

发布时间：2025-05-16 01:52:28

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除接口**：</font>

* OpenServerlessDBExtranetAccess

<font color="#dd0000">**预下线接口**：</font>

* CreateDBInstances
* DescribeDBSlowlogs
* InitDBInstances
* UpgradeDBInstance



## 腾讯云区块链服务平台 TBaaS(tbaas) 版本：2018-04-16

### 第 21 次发布

发布时间：2025-05-16 01:58:04

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除接口**：</font>

* GetLatesdTransactionList



## 容器服务(tke) 版本：2022-05-01



## 容器服务(tke) 版本：2018-05-25

### 第 195 次发布

发布时间：2025-05-16 02:05:38

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [EtcdOverrideConfig](https://cloud.tencent.com/document/api/457/31866#EtcdOverrideConfig)

修改数据结构：

* [ClusterAdvancedSettings](https://cloud.tencent.com/document/api/457/31866#ClusterAdvancedSettings)

	* 新增成员：EtcdOverrideConfigs




## 数据开发治理平台 WeData(wedata) 版本：2021-08-20

### 第 151 次发布

发布时间：2025-05-16 02:16:57

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeCodeTemplateDetail](https://cloud.tencent.com/document/api/1267/118493)
* [GetBatchDetailErrorLog](https://cloud.tencent.com/document/api/1267/118492)
* [ListBatchDetail](https://cloud.tencent.com/document/api/1267/118491)

新增数据结构：

* [AsyncResourceVO](https://cloud.tencent.com/document/api/1267/76336#AsyncResourceVO)

修改数据结构：

* [CodeTemplateDetail](https://cloud.tencent.com/document/api/1267/76336#CodeTemplateDetail)

	* 新增成员：Content




