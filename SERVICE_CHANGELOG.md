# Release v1.3.7

## 运维安全中心（堡垒机）(bh) 版本：2023-04-18

### 第 20 次发布

发布时间：2025-12-03 01:11:58

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeSecuritySetting](https://cloud.tencent.com/document/api/1025/125050)

	* 新增出参：SecuritySetting


新增数据结构：

* [AuthModeSetting](https://cloud.tencent.com/document/api/1025/74416#AuthModeSetting)
* [ReconnectionSetting](https://cloud.tencent.com/document/api/1025/74416#ReconnectionSetting)
* [SecuritySetting](https://cloud.tencent.com/document/api/1025/74416#SecuritySetting)

修改数据结构：

* [User](https://cloud.tencent.com/document/api/1025/74416#User)

	* 新增成员：RoleArn




## 云数据库 MySQL(cdb) 版本：2017-03-20

### 第 209 次发布

发布时间：2025-12-04 01:17:14

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyAccountPassword](https://cloud.tencent.com/document/api/236/17497)

	* 新增入参：SkipValidatePassword

* [SwitchForUpgrade](https://cloud.tencent.com/document/api/236/15864)

	* 新增入参：IsRelatedSwitch




## 腾讯云数据仓库TCHouse-C(cdwch) 版本：2020-09-15

### 第 34 次发布

发布时间：2025-12-03 01:19:42

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [ClusterConfigsInfoFromEMR](https://cloud.tencent.com/document/api/1299/83429#ClusterConfigsInfoFromEMR)

	* 新增成员：Ip, ConfigLevel




## 负载均衡(clb) 版本：2018-03-17

### 第 141 次发布

发布时间：2025-12-04 01:24:43

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [AvailableZoneAffinityInfo](https://cloud.tencent.com/document/api/214/30694#AvailableZoneAffinityInfo)

修改数据结构：

* [Backend](https://cloud.tencent.com/document/api/214/30694#Backend)

	* 新增成员：Zone

* [LoadBalancer](https://cloud.tencent.com/document/api/214/30694#LoadBalancer)

	* 新增成员：AvailableZoneAffinityInfo

* [LoadBalancerDetail](https://cloud.tencent.com/document/api/214/30694#LoadBalancerDetail)

	* 新增成员：AvailableZoneAffinityInfo




## TDSQL-C MySQL 版(cynosdb) 版本：2019-01-07

### 第 147 次发布

发布时间：2025-12-04 01:30:36

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [ModifyClusterGlobalEncryption](https://cloud.tencent.com/document/api/1003/126134)



## 腾讯云数据分析智能体(dataagent) 版本：2025-05-13

### 第 5 次发布

发布时间：2025-12-03 01:31:49

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [GetUploadJobDetails](https://cloud.tencent.com/document/api/1800/126105)
* [UploadAndCommitFile](https://cloud.tencent.com/document/api/1800/126104)

新增数据结构：

* [CosFileInfo](https://cloud.tencent.com/document/api/1800/125016#CosFileInfo)
* [UploadJob](https://cloud.tencent.com/document/api/1800/125016#UploadJob)



## 云数据库独享集群(dbdc) 版本：2020-10-29

### 第 5 次发布

发布时间：2025-12-04 01:33:26

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeInstanceDetail](https://cloud.tencent.com/document/api/1322/74752)

	* 新增出参：ResourceTags, CpuType


新增数据结构：

* [ResourceTag](https://cloud.tencent.com/document/api/1322/74754#ResourceTag)

修改数据结构：

* [DescribeInstanceDetail](https://cloud.tencent.com/document/api/1322/74754#DescribeInstanceDetail)

	* 新增成员：ResourceTags, CpuType




## 数据传输服务(dts) 版本：2021-12-06

### 第 50 次发布

发布时间：2025-12-04 01:39:23

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [DBOpFilter](https://cloud.tencent.com/document/api/571/82108#DBOpFilter)
* [OpFilter](https://cloud.tencent.com/document/api/571/82108#OpFilter)
* [TableFilter](https://cloud.tencent.com/document/api/571/82108#TableFilter)
* [ViewFilter](https://cloud.tencent.com/document/api/571/82108#ViewFilter)

修改数据结构：

* [Objects](https://cloud.tencent.com/document/api/571/82108#Objects)

	* 新增成员：DatabasesOpFilter




## 数据传输服务(dts) 版本：2018-03-30



## 密钥管理系统(kms) 版本：2019-01-18

### 第 26 次发布

发布时间：2025-12-04 01:56:50

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [GenerateDataKey](https://cloud.tencent.com/document/api/573/34419)

	* 新增入参：Tags

	* 新增出参：TagCode, TagMsg

* [ImportDataKey](https://cloud.tencent.com/document/api/573/120153)

	* 新增入参：Tags

	* 新增出参：TagCode, TagMsg

* [ListDataKeyDetail](https://cloud.tencent.com/document/api/573/120152)

	* 新增入参：TagFilters


修改数据结构：

* [DataKeyMetadata](https://cloud.tencent.com/document/api/573/34431#DataKeyMetadata)

	* 新增成员：KeyName




## 腾讯云智能体开发平台(lke) 版本：2023-11-30

### 第 71 次发布

发布时间：2025-12-03 02:00:43

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeAppAgentList](https://cloud.tencent.com/document/api/1759/120165)

	* 新增出参：HandoffAdvancedSetting


新增数据结构：

* [AgentHandoffAdvancedSetting](https://cloud.tencent.com/document/api/1759/105104#AgentHandoffAdvancedSetting)

修改数据结构：

* [AgentToolInfo](https://cloud.tencent.com/document/api/1759/105104#AgentToolInfo)

	* 新增成员：FinanceType




## 云数据库 MariaDB(mariadb) 版本：2017-03-12

### 第 76 次发布

发布时间：2025-12-04 02:03:47

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeDBInstanceDetail](https://cloud.tencent.com/document/api/237/89390)

	* 新增出参：FlowId




## 消息队列 MQTT 版(mqtt) 版本：2024-05-16

### 第 25 次发布

发布时间：2025-12-03 14:17:37

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateMessageEnrichmentRule](https://cloud.tencent.com/document/api/1778/126116)
* [DeleteMessageEnrichmentRule](https://cloud.tencent.com/document/api/1778/126115)
* [DescribeMessageEnrichmentRules](https://cloud.tencent.com/document/api/1778/126114)
* [ModifyMessageEnrichmentRule](https://cloud.tencent.com/document/api/1778/126113)
* [UpdateMessageEnrichmentRulePriority](https://cloud.tencent.com/document/api/1778/126112)

新增数据结构：

* [MessageEnrichmentRuleItem](https://cloud.tencent.com/document/api/1778/111031#MessageEnrichmentRuleItem)
* [MessageEnrichmentRulePriority](https://cloud.tencent.com/document/api/1778/111031#MessageEnrichmentRulePriority)



## 容器服务(tke) 版本：2022-05-01

### 第 17 次发布

发布时间：2025-12-03 03:05:55

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [ModifyClusterMachine](https://cloud.tencent.com/document/api/457/126106)



## 容器服务(tke) 版本：2018-05-25



## 腾讯混元生视频(vclm) 版本：2024-05-23

### 第 22 次发布

发布时间：2025-12-03 03:18:17

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeVideoFaceFusionJob](https://cloud.tencent.com/document/api/1616/126108)
* [SubmitVideoFaceFusionJob](https://cloud.tencent.com/document/api/1616/126107)

新增数据结构：

* [FaceMergeInfo](https://cloud.tencent.com/document/api/1616/107808#FaceMergeInfo)
* [FaceRect](https://cloud.tencent.com/document/api/1616/107808#FaceRect)
* [FaceTemplateInfo](https://cloud.tencent.com/document/api/1616/107808#FaceTemplateInfo)



## 私有网络(vpc) 版本：2017-03-12

### 第 279 次发布

发布时间：2025-12-04 03:24:44

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateNatGateway](https://cloud.tencent.com/document/api/215/36721)

	* 新增入参：ExclusiveType

* [ResetNatGatewayConnection](https://cloud.tencent.com/document/api/215/36713)

	* 新增入参：ExclusiveType


### 第 278 次发布

发布时间：2025-12-03 03:24:19

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateVpc](https://cloud.tencent.com/document/api/215/15774)

	* 新增入参：EnableRouteVpcPublishIpv6

* [ModifyVpcAttribute](https://cloud.tencent.com/document/api/215/15773)

	* 新增入参：EnableRouteVpcPublishIpv6


新增数据结构：

* [ConnectionStateTimeouts](https://cloud.tencent.com/document/api/215/15824#ConnectionStateTimeouts)

修改数据结构：

* [NatGateway](https://cloud.tencent.com/document/api/215/15824#NatGateway)

	* 新增成员：ConnectionStateTimeouts, ExclusiveType

* [TranslationAclRule](https://cloud.tencent.com/document/api/215/15824#TranslationAclRule)

	* <font color="#dd0000">**修改成员**：</font>SourceCidr, AclRuleId

* [Vpc](https://cloud.tencent.com/document/api/215/15824#Vpc)

	* 新增成员：EnableRouteVpcPublishIpv6




