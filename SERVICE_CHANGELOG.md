# Release v1.3.45

## 腾讯混元生图(aiart) 版本：2022-12-29

### 第 28 次发布

发布时间：2026-02-05 01:08:37

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除接口**：</font>

* DescribeTemplateToImageJob
* SubmitTemplateToImageJob



## 应用性能监控(apm) 版本：2021-06-22

### 第 54 次发布

发布时间：2026-02-05 01:10:56

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeApmAllVulCount](https://cloud.tencent.com/document/api/1463/128213)
* [DescribeApmVulnerabilityCount](https://cloud.tencent.com/document/api/1463/128212)
* [DescribeApmVulnerabilityDetail](https://cloud.tencent.com/document/api/1463/128211)
* [DescribeOPRAllVulCount](https://cloud.tencent.com/document/api/1463/128210)

新增数据结构：

* [ApmVulnerabilityServiceDetail](https://cloud.tencent.com/document/api/1463/64927#ApmVulnerabilityServiceDetail)



## 运维安全中心（堡垒机）(bh) 版本：2023-04-18

### 第 25 次发布

发布时间：2026-02-05 01:12:52

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [RunOperationTask](https://cloud.tencent.com/document/api/1025/117829)

	* 新增出参：SubTaskId




## 腾讯云数据仓库TCHouse-P(cdwpg) 版本：2020-12-30

### 第 16 次发布

发布时间：2026-02-05 01:22:09

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeInstance](https://cloud.tencent.com/document/api/878/100195)

	* 新增出参：ErrorMsg

* [DescribeInstanceOperations](https://cloud.tencent.com/document/api/878/116664)

	* 新增出参：ErrorMsg

* [DescribeInstanceState](https://cloud.tencent.com/document/api/878/100160)

	* 新增出参：BackupOpenStatus

* [DescribeUpgradeList](https://cloud.tencent.com/document/api/878/116662)

	* 新增出参：ErrorMsg


修改数据结构：

* [InstanceInfo](https://cloud.tencent.com/document/api/878/98895#InstanceInfo)

	* 新增成员：IsAz, SecondaryZone, SecondarySubnet, AccessInfo, GTMNodes

* [InstanceNode](https://cloud.tencent.com/document/api/878/98895#InstanceNode)

	* 新增成员：PrivateNetworkIp, NodeRole, NodeName, SpecName, Cpu, Memory, DataDiskCount, DataDiskSize, DataDiskType, UUID, Zone

* [ParamDetail](https://cloud.tencent.com/document/api/878/98895#ParamDetail)

	* 新增成员：LatestValue




## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 275 次发布

发布时间：2026-02-05 01:50:43

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [OutputRisk](https://cloud.tencent.com/document/api/1323/70369#OutputRisk)

	* 新增成员：IsMark, IsIgnore, RiskLevelAliasName




## 云直播CSS(live) 版本：2018-08-01

### 第 164 次发布

发布时间：2026-02-05 02:08:29

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [CommonMixLayoutParams](https://cloud.tencent.com/document/api/267/20474#CommonMixLayoutParams)

	* 新增成员：WebPageUrl




## 媒体处理(mps) 版本：2019-06-12

### 第 181 次发布

发布时间：2026-02-05 02:21:18

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateBlindWatermarkTemplate](https://cloud.tencent.com/document/api/862/125874)

	* 新增入参：Strength

* [ModifyBlindWatermarkTemplate](https://cloud.tencent.com/document/api/862/125871)

	* 新增入参：Strength

* [ProcessLiveStream](https://cloud.tencent.com/document/api/862/39227)

	* 新增入参：SmartSubtitlesTask


新增数据结构：

* [LiveSmartSubtitlesTaskInput](https://cloud.tencent.com/document/api/862/37615#LiveSmartSubtitlesTaskInput)

修改数据结构：

* [BlindWatermarkTemplate](https://cloud.tencent.com/document/api/862/37615#BlindWatermarkTemplate)

	* 新增成员：Strength




## Web 应用防火墙(waf) 版本：2018-01-25

### 第 141 次发布

发布时间：2026-02-05 03:45:27

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateRateLimitV2](https://cloud.tencent.com/document/api/627/125859)

	* 新增出参：LimitRuleID, Domain

* [DescribeApiDetail](https://cloud.tencent.com/document/api/627/111557)

	* 新增出参：FullReqLog

* [ImportIpAccessControl](https://cloud.tencent.com/document/api/627/105913)

	* 新增出参：SuccessCount, TotalCount, Timestamp

* [ModifyApiSecSensitiveRule](https://cloud.tencent.com/document/api/627/125861)

	* 新增入参：CustomApiExcludeRule, ApiExcludeRuleName

* [UpdateRateLimitV2](https://cloud.tencent.com/document/api/627/125855)

	* 新增出参：LimitRuleID, Domain

* [UpsertCCRule](https://cloud.tencent.com/document/api/627/97646)

	* 新增入参：JobType, JobDateTime, ExpireTime, ValidStatus


新增数据结构：

* [ApiSecExcludeRule](https://cloud.tencent.com/document/api/627/53609#ApiSecExcludeRule)
* [DedicatedIPPkg](https://cloud.tencent.com/document/api/627/53609#DedicatedIPPkg)

修改数据结构：

* [ApiDetailSampleHistory](https://cloud.tencent.com/document/api/627/53609#ApiDetailSampleHistory)

	* 新增成员：FullReqLog

* [BotToken](https://cloud.tencent.com/document/api/627/53609#BotToken)

	* 新增成员：DisableMultiJson

* [FieldWriteConfig](https://cloud.tencent.com/document/api/627/53609#FieldWriteConfig)

	* 新增成员：EnableResponse

* [InstanceInfo](https://cloud.tencent.com/document/api/627/53609#InstanceInfo)

	* 新增成员：DedicatedIPPkg, DedicatedIPCount

* [IpAccessControlParam](https://cloud.tencent.com/document/api/627/53609#IpAccessControlParam)

	* 新增成员：JobType, JobDateTime, ValidStatus

* [MiniPkg](https://cloud.tencent.com/document/api/627/53609#MiniPkg)

	* 新增成员：GatewayType

* [PostCLSFlowInfo](https://cloud.tencent.com/document/api/627/53609#PostCLSFlowInfo)

	* 新增成员：WriteConfig




