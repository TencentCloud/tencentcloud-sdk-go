# Release v1.3.123

## API 网关(apigateway) 版本：2018-08-08

### 第 53 次发布

发布时间：2026-06-26 01:08:56

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除接口**：</font>

* DescribeApiApp



## 运维安全中心（堡垒机）(bh) 版本：2023-04-18

### 第 33 次发布

发布时间：2026-06-26 01:10:33

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateUserDirectory](https://cloud.tencent.com/document/api/1025/125473)

	* 新增入参：AutoSync, SyncCron

* [ModifyUserDirectory](https://cloud.tencent.com/document/api/1025/125469)

	* 新增入参：AutoSync, SyncCron


修改数据结构：

* [IOAUserGroup](https://cloud.tencent.com/document/api/1025/74416#IOAUserGroup)

	* 新增成员：UserDirName

* [UserDirectory](https://cloud.tencent.com/document/api/1025/74416#UserDirectory)

	* 新增成员：AutoSync, SyncCron, NextSyncTime

* [UserOrg](https://cloud.tencent.com/document/api/1025/74416#UserOrg)

	* 新增成员：BindGroupIds




## 费用中心(billing) 版本：2018-07-09

### 第 92 次发布

发布时间：2026-06-26 01:11:26

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [BillDetail](https://cloud.tencent.com/document/api/555/19183#BillDetail)

	* 新增成员：ExtendField




## 负载均衡(clb) 版本：2018-03-17

### 第 151 次发布

发布时间：2026-06-26 01:17:21

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [AssociateBudget](https://cloud.tencent.com/document/api/214/133221)
* [CreateBudget](https://cloud.tencent.com/document/api/214/133220)
* [CreateKey](https://cloud.tencent.com/document/api/214/133219)
* [CreateKeys](https://cloud.tencent.com/document/api/214/133218)
* [CreateModelRouter](https://cloud.tencent.com/document/api/214/133217)
* [CreateUserGroup](https://cloud.tencent.com/document/api/214/133200)
* [DeleteBudgets](https://cloud.tencent.com/document/api/214/133216)
* [DeleteKeys](https://cloud.tencent.com/document/api/214/133215)
* [DeleteModelRouters](https://cloud.tencent.com/document/api/214/133214)
* [DeleteUserGroups](https://cloud.tencent.com/document/api/214/133199)
* [DescribeAsyncJobs](https://cloud.tencent.com/document/api/214/133213)
* [DescribeBudgetAssociations](https://cloud.tencent.com/document/api/214/133212)
* [DescribeBudgets](https://cloud.tencent.com/document/api/214/133211)
* [DescribeModelRouterDetail](https://cloud.tencent.com/document/api/214/133210)
* [DescribeModelRouterQuota](https://cloud.tencent.com/document/api/214/133209)
* [DescribeModelRouters](https://cloud.tencent.com/document/api/214/133208)
* [DescribeUserGroups](https://cloud.tencent.com/document/api/214/133198)
* [DisassociateBudget](https://cloud.tencent.com/document/api/214/133207)
* [ModifyBudgetAttributes](https://cloud.tencent.com/document/api/214/133206)
* [ModifyKeyAttributes](https://cloud.tencent.com/document/api/214/133205)
* [ModifyKeysBlockStatus](https://cloud.tencent.com/document/api/214/133204)
* [ModifyKeysUserGroup](https://cloud.tencent.com/document/api/214/133197)
* [ModifyModelRouterAttributes](https://cloud.tencent.com/document/api/214/133203)
* [ModifyUserGroupAttributes](https://cloud.tencent.com/document/api/214/133196)
* [RegenerateKeys](https://cloud.tencent.com/document/api/214/133202)

修改接口：

* [SetCustomizedConfigForLoadBalancer](https://cloud.tencent.com/document/api/214/60008)

	* 新增入参：Tags


新增数据结构：

* [BudgetAssociation](https://cloud.tencent.com/document/api/214/30694#BudgetAssociation)
* [BudgetConfig](https://cloud.tencent.com/document/api/214/30694#BudgetConfig)
* [BudgetConfigInput](https://cloud.tencent.com/document/api/214/30694#BudgetConfigInput)
* [BudgetInfo](https://cloud.tencent.com/document/api/214/30694#BudgetInfo)
* [BudgetResource](https://cloud.tencent.com/document/api/214/30694#BudgetResource)
* [ClusterInfo](https://cloud.tencent.com/document/api/214/30694#ClusterInfo)
* [CreatedKey](https://cloud.tencent.com/document/api/214/30694#CreatedKey)
* [CreditUsage](https://cloud.tencent.com/document/api/214/30694#CreditUsage)
* [FallBackItem](https://cloud.tencent.com/document/api/214/30694#FallBackItem)
* [InputKeyInfo](https://cloud.tencent.com/document/api/214/30694#InputKeyInfo)
* [Job](https://cloud.tencent.com/document/api/214/30694#Job)
* [ModelRouterDetail](https://cloud.tencent.com/document/api/214/30694#ModelRouterDetail)
* [ModelRouterQuota](https://cloud.tencent.com/document/api/214/30694#ModelRouterQuota)
* [ModelRouterSet](https://cloud.tencent.com/document/api/214/30694#ModelRouterSet)
* [RateLimitConfigForBudget](https://cloud.tencent.com/document/api/214/30694#RateLimitConfigForBudget)
* [RateLimitConfigForKey](https://cloud.tencent.com/document/api/214/30694#RateLimitConfigForKey)
* [RateLimitConfigForModelRouter](https://cloud.tencent.com/document/api/214/30694#RateLimitConfigForModelRouter)
* [RegeneratedKey](https://cloud.tencent.com/document/api/214/30694#RegeneratedKey)
* [RouterSettingWithFallBack](https://cloud.tencent.com/document/api/214/30694#RouterSettingWithFallBack)
* [RouterSettingWithoutFallBack](https://cloud.tencent.com/document/api/214/30694#RouterSettingWithoutFallBack)
* [ServiceEndPoints](https://cloud.tencent.com/document/api/214/30694#ServiceEndPoints)
* [UserGroupInfo](https://cloud.tencent.com/document/api/214/30694#UserGroupInfo)



## TDSQL-C MySQL 版(cynosdb) 版本：2019-01-07

### 第 176 次发布

发布时间：2026-06-26 01:23:45

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [QueryFilter](https://cloud.tencent.com/document/api/1003/48097#QueryFilter)

	* <font color="#dd0000">**修改成员**：</font>Names




## Elasticsearch Service(es) 版本：2025-01-01



## Elasticsearch Service(es) 版本：2018-04-16

### 第 102 次发布

发布时间：2026-06-26 01:30:12

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateInstance](https://cloud.tencent.com/document/api/845/30633)

	* 新增入参：EnableCosBackup




## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 301 次发布

发布时间：2026-06-26 01:30:39

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CancelOrganizationFlows](https://cloud.tencent.com/document/api/1323/133222)



## 腾讯电子签（基础版）(essbasic) 版本：2021-05-26

### 第 259 次发布

发布时间：2026-06-26 01:31:28

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CancelOrganizationFlows](https://cloud.tencent.com/document/api/1420/133223)



## 腾讯电子签（基础版）(essbasic) 版本：2020-12-22



## 云直播CSS(live) 版本：2018-08-01

### 第 178 次发布

发布时间：2026-06-26 01:39:08

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeAIGCTaskStatus](https://cloud.tencent.com/document/api/267/129793)

	* 新增出参：RequestBody, TaskType




## 文字识别(ocr) 版本：2018-11-19

### 第 255 次发布

发布时间：2026-06-26 01:44:24

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [MultimodalDocParse](https://cloud.tencent.com/document/api/866/133224)

修改接口：

* [ExtractDocAgent](https://cloud.tencent.com/document/api/866/126442)

	* 新增入参：EnableCoord, EnableAudit




## 集团账号管理(organization) 版本：2021-03-31

### 第 58 次发布

发布时间：2026-06-26 01:45:13

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [NotAllowReason](https://cloud.tencent.com/document/api/850/67060#NotAllowReason)

	* 新增成员：CICRoleConfig, CICUserConfig




## 集团账号管理(organization) 版本：2018-12-25



## 云数据库 PostgreSQL(postgres) 版本：2017-03-12

### 第 68 次发布

发布时间：2026-06-26 01:45:59

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [BackupPlan](https://cloud.tencent.com/document/api/409/16778#BackupPlan)

	* 新增成员：BackupMethod




## 云数据库Redis(redis) 版本：2018-04-12

### 第 102 次发布

发布时间：2026-06-26 01:47:00

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CloseSSL](https://cloud.tencent.com/document/api/239/81240)

	* 新增入参：AddressType




## 云开发 CloudBase(tcb) 版本：2018-06-08

### 第 148 次发布

发布时间：2026-06-26 01:50:40

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyEnv](https://cloud.tencent.com/document/api/876/34818)

	* 新增入参：CustomQps




## TDSQL(tdmysql) 版本：2021-11-22

### 第 9 次发布

发布时间：2026-06-25 19:37:52

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [ResetUsersPassword](https://cloud.tencent.com/document/api/1376/133194)

新增数据结构：

* [ResetUserPasswordInfo](https://cloud.tencent.com/document/api/1376/128305#ResetUserPasswordInfo)



## 容器服务(tke) 版本：2022-05-01

### 第 27 次发布

发布时间：2026-06-26 01:58:12

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeNodePoolsElasticityStrength](https://cloud.tencent.com/document/api/457/133225)

修改数据结构：

* [CreateNativeNodePoolParam](https://cloud.tencent.com/document/api/457/103206#CreateNativeNodePoolParam)

	* 新增成员：GPUConfigs




## 容器服务(tke) 版本：2018-05-25



## 实时音视频(trtc) 版本：2019-07-22

### 第 146 次发布

发布时间：2026-06-26 01:59:27

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [WordItem](https://cloud.tencent.com/document/api/647/44055#WordItem)

修改数据结构：

* [AlignmentItem](https://cloud.tencent.com/document/api/647/44055#AlignmentItem)

	* 新增成员：Text, Words

* [McuLayout](https://cloud.tencent.com/document/api/647/44055#McuLayout)

	* 新增成员：EnableStreamSEI

* [McuLayoutParams](https://cloud.tencent.com/document/api/647/44055#McuLayoutParams)

	* 新增成员：EnableStreamSEI




