# Release v1.3.211

## 腾讯云智能体开发平台(adp) 版本：2026-05-20

### 第 36 次发布

发布时间：2026-11-09 01:03:37

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeMetricOverviewList](https://cloud.tencent.com/document/api/1759/136062)

	* 新增入参：SpaceId

* [DescribeUsageDetailList](https://cloud.tencent.com/document/api/1759/136061)

	* 新增入参：SpaceId

* [DescribeUsageSummaryList](https://cloud.tencent.com/document/api/1759/136060)

	* 新增入参：SpaceId


新增数据结构：

* [DocumentParseUsageDetail](https://cloud.tencent.com/document/api/1759/132545#DocumentParseUsageDetail)

修改数据结构：

* [CallSource](https://cloud.tencent.com/document/api/1759/132545#CallSource)

	* 新增成员：DeploymentUuid, SpaceId, SpaceName

* [ModelUsageDetail](https://cloud.tencent.com/document/api/1759/132545#ModelUsageDetail)

	* 新增成员：ApiKeyDisplay, BillingType

* [UsageDetail](https://cloud.tencent.com/document/api/1759/132545#UsageDetail)

	* 新增成员：DocumentParse




## Agent 沙箱服务(ags) 版本：2025-09-20

### 第 33 次发布

发布时间：2026-11-09 01:04:56

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribePreCacheImageTask](https://cloud.tencent.com/document/api/1814/127507)

	* 新增入参：PreCacheType

* [StartSandboxInstance](https://cloud.tencent.com/document/api/1814/124816)

	* 新增入参：StartMode


修改数据结构：

* [ResourceConfiguration](https://cloud.tencent.com/document/api/1814/124823#ResourceConfiguration)

	* 新增成员：GPUResourceGroupId




## AI Agent 安全网关(apis) 版本：2024-08-01

### 第 16 次发布

发布时间：2026-11-09 01:07:51

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [AgentAppMcpServerVO](https://cloud.tencent.com/document/api/1627/129635#AgentAppMcpServerVO)

	* 新增成员：AgentAppVO

* [AgentAppServiceVO](https://cloud.tencent.com/document/api/1627/129635#AgentAppServiceVO)

	* 新增成员：AgentAppVO




## 灾备中心(bdrc) 版本：2026-03-30

### 第 5 次发布

发布时间：2026-11-09 01:09:49

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyCopyPairAttribute](https://cloud.tencent.com/document/api/1824/136754)

	* 新增入参：PrivateIpAddress, UseDefaultPrivateIp, LoginSettings, UseDefaultPassword

* [RunFailoverCopyPairs](https://cloud.tencent.com/document/api/1824/136748)

	* 新增入参：StopSourceInstance




## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 328 次发布

发布时间：2026-11-09 01:55:50

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [ContractReviewWebUrlOption](https://cloud.tencent.com/document/api/1323/70369#ContractReviewWebUrlOption)

	* 新增成员：DisableIgnoreRisk, DisableRevisedRisk




## 云游戏(gs) 版本：2019-11-18

### 第 69 次发布

发布时间：2026-11-09 02:01:07

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateAndroidInstanceMCPToken](https://cloud.tencent.com/document/api/1162/140547)

	* 新增入参：Role, ExpireSeconds, TokenName, AllowedInstanceIds

	* 新增出参：TokenId, Token

* [DeleteAndroidInstanceMCPToken](https://cloud.tencent.com/document/api/1162/140546)

	* 新增入参：TokenId

* [DescribeAndroidInstanceMCPTokens](https://cloud.tencent.com/document/api/1162/140545)

	* 新增入参：Offset, Limit

	* 新增出参：Tokens, TotalCount

* [RenewAndroidInstanceMCPToken](https://cloud.tencent.com/document/api/1162/140544)

	* 新增入参：TokenId, ExpireSeconds

	* 新增出参：ExpiresTime


新增数据结构：

* [AndroidInstanceMCPToken](https://cloud.tencent.com/document/api/1162/40743#AndroidInstanceMCPToken)



## 智能视图计算平台(iss) 版本：2023-05-17

### 第 38 次发布

发布时间：2026-11-09 02:09:43

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [ListVLPlanResults](https://cloud.tencent.com/document/api/1344/140649)

新增数据结构：

* [VLPlanResultInfo](https://cloud.tencent.com/document/api/1344/95952#VLPlanResultInfo)



## 文字识别(ocr) 版本：2018-11-19

### 第 277 次发布

发布时间：2026-11-09 02:23:09

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeExtractDocAgentJob](https://cloud.tencent.com/document/api/866/123153)

	* 新增出参：AgentResultJson

* [SubmitExtractDocAgentJob](https://cloud.tencent.com/document/api/866/123134)

	* 新增入参：AgentType




## 云数据库Redis(redis) 版本：2018-04-12

### 第 113 次发布

发布时间：2026-11-09 02:28:00

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyInstanceAccount](https://cloud.tencent.com/document/api/239/38923)

	* 新增入参：DisableCommands, ModifyDisableCommands




## TDSQL(tdmysql) 版本：2021-11-22

### 第 20 次发布

发布时间：2026-11-09 02:41:56

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [ModifyDBInstance](https://cloud.tencent.com/document/api/1376/140650)



## 边缘安全加速平台(teo) 版本：2022-09-01

### 第 172 次发布

发布时间：2026-11-09 02:43:09

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeTimingFunctionAnalysisData](https://cloud.tencent.com/document/api/1552/140652)
* [DescribeTopFunctionAnalysisData](https://cloud.tencent.com/document/api/1552/140651)



## 边缘安全加速平台(teo) 版本：2022-01-06



## TokenHub(tokenhub) 版本：2026-03-22

### 第 27 次发布

发布时间：2026-11-09 02:49:36

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [ModifyApiKeyBindings](https://cloud.tencent.com/document/api/1823/140653)



