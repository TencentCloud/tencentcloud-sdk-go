# Release v1.3.105

## 云数据库 MySQL(cdb) 版本：2017-03-20

### 第 220 次发布

发布时间：2026-05-28 01:19:21

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeInstanceUpgradeType](https://cloud.tencent.com/document/api/236/107630)

	* 新增入参：DstFourthZone

* [UpgradeDBInstance](https://cloud.tencent.com/document/api/236/15876)

	* 新增入参：FourthZone




## 负载均衡(clb) 版本：2018-03-17

### 第 149 次发布

发布时间：2026-05-28 01:27:48

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [TargetGroupInfo](https://cloud.tencent.com/document/api/214/30694#TargetGroupInfo)

	* 新增成员：SnatEnable




## 云安全一体化平台(csip) 版本：2022-11-21

### 第 79 次发布

发布时间：2026-05-28 01:32:43

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateCosObjectScanTask](https://cloud.tencent.com/document/api/664/131697)

	* 新增入参：TaskArgs, IsScanAll, DeleteBucketSet

	* 新增出参：TaskId


修改数据结构：

* [CosAssetInfo](https://cloud.tencent.com/document/api/664/90825#CosAssetInfo)

	* 新增成员：BucketAzType, BucketStorageSize, BucketObjectCount, IdentifySampleRate

* [CosAuditPayInfo](https://cloud.tencent.com/document/api/664/90825#CosAuditPayInfo)

	* 新增成员：PostProductStatusList, PostProductBuyStatusList, NewPostPayResourceId

* [CosBucketBillingInfo](https://cloud.tencent.com/document/api/664/90825#CosBucketBillingInfo)

	* 新增成员：LogFeatureWhitelist, IsHaveNewPostOrder, IsHaveOldPostOrder, PostProductList




## TDSQL-C MySQL 版(cynosdb) 版本：2019-01-07

### 第 170 次发布

发布时间：2026-05-28 01:41:24

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [AddServerlessRoInstances](https://cloud.tencent.com/document/api/1003/132223)
* [DescribeSQLExecutionPlan](https://cloud.tencent.com/document/api/1003/132225)

新增数据结构：

* [ExecutionPlanDetail](https://cloud.tencent.com/document/api/1003/48097#ExecutionPlanDetail)
* [ExplainRow](https://cloud.tencent.com/document/api/1003/48097#ExplainRow)



## 全球加速(ga2) 版本：2025-01-15

### 第 5 次发布

发布时间：2026-05-28 01:59:23

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateForwardingPolicy](https://cloud.tencent.com/document/api/1817/132229)
* [DeleteForwardingPolicy](https://cloud.tencent.com/document/api/1817/132228)
* [DescribeForwardingPolicy](https://cloud.tencent.com/document/api/1817/132227)
* [ModifyForwardingPolicy](https://cloud.tencent.com/document/api/1817/132226)

新增数据结构：

* [ForwardingPolicySet](https://cloud.tencent.com/document/api/1817/130045#ForwardingPolicySet)



## 高性能应用服务(hai) 版本：2023-08-12

### 第 23 次发布

发布时间：2026-05-28 02:02:24

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [ServiceDetail](https://cloud.tencent.com/document/api/1721/101518#ServiceDetail)

	* 新增成员：SecurityType, RoleComputeSet, TargetReplicas

* [TemplateDetail](https://cloud.tencent.com/document/api/1721/101518#TemplateDetail)

	* 新增成员：RoleComputeSet




## SSL 证书(ssl) 版本：2019-12-05

### 第 98 次发布

发布时间：2026-05-28 02:38:15

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeCertificateDetail](https://cloud.tencent.com/document/api/400/41673)

	* 新增出参：HostingStatus




## 云开发 CloudBase(tcb) 版本：2018-06-08

### 第 142 次发布

发布时间：2026-05-28 02:41:04

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [SMSProviderTemplateConfig](https://cloud.tencent.com/document/api/876/34822#SMSProviderTemplateConfig)
* [SMSTemplateParams](https://cloud.tencent.com/document/api/876/34822#SMSTemplateParams)

修改数据结构：

* [VerificationConfig](https://cloud.tencent.com/document/api/876/34822#VerificationConfig)

	* 新增成员：TemplateProvider




## Web 应用防火墙(waf) 版本：2018-01-25

### 第 156 次发布

发布时间：2026-05-28 03:12:16

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeIpAccessControl](https://cloud.tencent.com/document/api/627/72645)

	* 新增入参：IpList




