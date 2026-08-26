# Release v1.3.169

## Agent 沙箱服务(ags) 版本：2025-09-20

### 第 20 次发布

发布时间：2026-08-26 01:09:41

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateSandboxTool](https://cloud.tencent.com/document/api/1814/124812)

	* 新增入参：ComputerConfiguration

* [UpdateSandboxTool](https://cloud.tencent.com/document/api/1814/124809)

	* 新增入参：ComputerConfiguration




## 灾备中心(bdrc) 版本：2026-03-30

### 第 2 次发布

发布时间：2026-08-26 01:12:32

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateFileBackupPlan](https://cloud.tencent.com/document/api/1824/136817)

	* 新增入参：ResourceType

	* <font color="#dd0000">**修改入参**：</font>BackupStorageId




## 云联络中心(ccc) 版本：2020-02-10

### 第 134 次发布

发布时间：2026-08-25 16:12:39

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateAIAgentCall](https://cloud.tencent.com/document/api/679/115681)

	* 新增入参：AcquireTimeoutSecond




## 云原生智能网关(cngw) 版本：2023-04-18

### 第 7 次发布

发布时间：2026-08-26 01:23:06

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [AIGWLogConfig](https://cloud.tencent.com/document/api/1826/133161#AIGWLogConfig)

	* 新增成员：RequestLogPayloadTruncationPolicy, ResponseLogPayloadTruncationPolicy




## 云安全一体化平台(csip) 版本：2022-11-21

### 第 103 次发布

发布时间：2026-08-26 01:23:49

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateHostImageListExportJob](https://cloud.tencent.com/document/api/664/136900)
* [CreateSandboxACLRule](https://cloud.tencent.com/document/api/664/136923)
* [CreateSandboxDLPRule](https://cloud.tencent.com/document/api/664/136922)
* [CreateSandboxFileRule](https://cloud.tencent.com/document/api/664/136921)
* [CreateSandboxLLMAuditRule](https://cloud.tencent.com/document/api/664/136920)
* [DeleteSandboxACLRule](https://cloud.tencent.com/document/api/664/136919)
* [DeleteSandboxDLPRule](https://cloud.tencent.com/document/api/664/136918)
* [DeleteSandboxFileRule](https://cloud.tencent.com/document/api/664/136917)
* [DescribeAccessKeyWhiteList](https://cloud.tencent.com/document/api/664/136902)
* [DescribeCSCPayInfo](https://cloud.tencent.com/document/api/664/136896)
* [DescribeCSPMPayInfo](https://cloud.tencent.com/document/api/664/136895)
* [DescribeSandboxDLPAlertList](https://cloud.tencent.com/document/api/664/136916)
* [DescribeSandboxDLPRuleList](https://cloud.tencent.com/document/api/664/136915)
* [DescribeSandboxLLMAuditAlertList](https://cloud.tencent.com/document/api/664/136914)
* [DescribeSandboxLLMAuditRuleList](https://cloud.tencent.com/document/api/664/136913)
* [DescribeSandboxLLMAuditSystemRuleList](https://cloud.tencent.com/document/api/664/136912)
* [DescribeSourceIPDetail](https://cloud.tencent.com/document/api/664/136901)
* [DescribeUserAKInfoList](https://cloud.tencent.com/document/api/664/136894)
* [DescribeVoucherEligibility](https://cloud.tencent.com/document/api/664/136893)
* [InstallKeySandboxSkill](https://cloud.tencent.com/document/api/664/136898)
* [InstallSandboxPlugin](https://cloud.tencent.com/document/api/664/136911)
* [ModifyClusterDefendStatus](https://cloud.tencent.com/document/api/664/136899)
* [ModifySandboxACLRule](https://cloud.tencent.com/document/api/664/136910)
* [ModifySandboxACLRuleStatus](https://cloud.tencent.com/document/api/664/136909)
* [ModifySandboxAlertStatus](https://cloud.tencent.com/document/api/664/136908)
* [ModifySandboxDLPRule](https://cloud.tencent.com/document/api/664/136907)
* [ModifySandboxDLPRuleStatus](https://cloud.tencent.com/document/api/664/136906)
* [ModifySandboxFileRule](https://cloud.tencent.com/document/api/664/136905)
* [ModifySandboxFileRuleStatus](https://cloud.tencent.com/document/api/664/136904)
* [ModifySandboxLLMAuditRule](https://cloud.tencent.com/document/api/664/136903)
* [ModifyShareUserAK](https://cloud.tencent.com/document/api/664/136892)
* [UninstallKeySandboxSkill](https://cloud.tencent.com/document/api/664/136897)

修改接口：

* [CreateDspmIdentifyCategory](https://cloud.tencent.com/document/api/664/133800)

	* 新增入参：OperationSource

* [CreateDspmIdentifyComplianceCategoryRelation](https://cloud.tencent.com/document/api/664/133799)

	* 新增入参：OperationSource

* [CreateDspmIdentifyComplianceGroup](https://cloud.tencent.com/document/api/664/133798)

	* 新增入参：OperationSource

* [CreateDspmIdentifyComplianceGroupCopy](https://cloud.tencent.com/document/api/664/133797)

	* 新增入参：OperationSource

* [CreateDspmIdentifyComplianceRuleRelation](https://cloud.tencent.com/document/api/664/133796)

	* 新增入参：OperationSource

* [CreateDspmIdentifyLevelGroup](https://cloud.tencent.com/document/api/664/133795)

	* 新增入参：OperationSource

* [CreateDspmIdentifyRule](https://cloud.tencent.com/document/api/664/133794)

	* 新增入参：OperationSource

* [DeleteDspmIdentifyCategory](https://cloud.tencent.com/document/api/664/133793)

	* 新增入参：OperationSource

* [DeleteDspmIdentifyComplianceCategoryRelation](https://cloud.tencent.com/document/api/664/133792)

	* 新增入参：OperationSource

* [DeleteDspmIdentifyComplianceGroup](https://cloud.tencent.com/document/api/664/133791)

	* 新增入参：OperationSource

* [DeleteDspmIdentifyComplianceRuleRelation](https://cloud.tencent.com/document/api/664/133790)

	* 新增入参：OperationSource

* [DeleteDspmIdentifyLevelGroup](https://cloud.tencent.com/document/api/664/133789)

	* 新增入参：OperationSource

* [DeleteDspmIdentifyRule](https://cloud.tencent.com/document/api/664/133788)

	* 新增入参：OperationSource

* [DescribeDspmDictionaryList](https://cloud.tencent.com/document/api/664/131495)

	* 新增入参：OperationSource

* [DescribeDspmIdentifyCategoryList](https://cloud.tencent.com/document/api/664/133785)

	* 新增入参：OperationSource

* [DescribeDspmIdentifyComplianceCategoryRuleList](https://cloud.tencent.com/document/api/664/133784)

	* 新增入参：OperationSource

* [DescribeDspmIdentifyComplianceGroupDetail](https://cloud.tencent.com/document/api/664/133783)

	* 新增入参：OperationSource

* [DescribeDspmIdentifyComplianceGroupList](https://cloud.tencent.com/document/api/664/133782)

	* 新增入参：OperationSource

* [DescribeDspmIdentifyLevelGroupList](https://cloud.tencent.com/document/api/664/133780)

	* 新增入参：OperationSource

* [DescribeDspmIdentifyRuleDetail](https://cloud.tencent.com/document/api/664/133779)

	* 新增入参：OperationSource

* [DescribeDspmIdentifyRuleList](https://cloud.tencent.com/document/api/664/133778)

	* 新增入参：OperationSource

* [DescribeDspmIdentifyRuleTestResult](https://cloud.tencent.com/document/api/664/133777)

	* 新增入参：OperationSource

* [ModifyCosAuditObjectSampleRate](https://cloud.tencent.com/document/api/664/133802)

	* 新增入参：TargetAppId, DefaultSampleRate

	* <font color="#dd0000">**修改入参**：</font>BucketIdSet, SampleRateSet

	* 新增出参：DefaultSampleRateUpdated, DefaultSampleRate, UpdatedBucketCount

* [ModifyDspmApplyingIdentifyComplianceGroup](https://cloud.tencent.com/document/api/664/133776)

	* 新增入参：OperationSource

* [ModifyDspmIdentifyCategory](https://cloud.tencent.com/document/api/664/133774)

	* 新增入参：OperationSource

* [ModifyDspmIdentifyComplianceGroup](https://cloud.tencent.com/document/api/664/133773)

	* 新增入参：OperationSource

* [ModifyDspmIdentifyComplianceGroupStatus](https://cloud.tencent.com/document/api/664/133772)

	* 新增入参：OperationSource

* [ModifyDspmIdentifyComplianceRuleLevelInfo](https://cloud.tencent.com/document/api/664/133771)

	* 新增入参：OperationSource

* [ModifyDspmIdentifyLevelGroup](https://cloud.tencent.com/document/api/664/133770)

	* 新增入参：OperationSource

* [ModifyDspmIdentifyLevelItem](https://cloud.tencent.com/document/api/664/133769)

	* 新增入参：OperationSource

* [ModifyDspmIdentifyRule](https://cloud.tencent.com/document/api/664/133768)

	* 新增入参：OperationSource

* [ModifyDspmIdentifyRuleStatus](https://cloud.tencent.com/document/api/664/133767)

	* 新增入参：OperationSource


新增数据结构：

* [AccessKeyWhiteList](https://cloud.tencent.com/document/api/664/90825#AccessKeyWhiteList)
* [CosIdentifyLevelDetail](https://cloud.tencent.com/document/api/664/90825#CosIdentifyLevelDetail)
* [OrderQuotaInfo](https://cloud.tencent.com/document/api/664/90825#OrderQuotaInfo)
* [TrafficSandboxDLPAlertInfo](https://cloud.tencent.com/document/api/664/90825#TrafficSandboxDLPAlertInfo)
* [TrafficSandboxDLPFileSizeRange](https://cloud.tencent.com/document/api/664/90825#TrafficSandboxDLPFileSizeRange)
* [TrafficSandboxDLPRuleContentItem](https://cloud.tencent.com/document/api/664/90825#TrafficSandboxDLPRuleContentItem)
* [TrafficSandboxDLPRuleInfo](https://cloud.tencent.com/document/api/664/90825#TrafficSandboxDLPRuleInfo)
* [TrafficSandboxDLPTrafficRuleItem](https://cloud.tencent.com/document/api/664/90825#TrafficSandboxDLPTrafficRuleItem)
* [TrafficSandboxDLPURLRuleItem](https://cloud.tencent.com/document/api/664/90825#TrafficSandboxDLPURLRuleItem)
* [TrafficSandboxDLPUserRuleInfo](https://cloud.tencent.com/document/api/664/90825#TrafficSandboxDLPUserRuleInfo)
* [TrafficSandboxLLMAuditAlertInfo](https://cloud.tencent.com/document/api/664/90825#TrafficSandboxLLMAuditAlertInfo)
* [TrafficSandboxLLMAuditRuleInfo](https://cloud.tencent.com/document/api/664/90825#TrafficSandboxLLMAuditRuleInfo)
* [TrafficSandboxLLMAuditRuleRef](https://cloud.tencent.com/document/api/664/90825#TrafficSandboxLLMAuditRuleRef)
* [TrafficSandboxLLMAuditSystemRuleItem](https://cloud.tencent.com/document/api/664/90825#TrafficSandboxLLMAuditSystemRuleItem)
* [UserAKInfo](https://cloud.tencent.com/document/api/664/90825#UserAKInfo)

修改数据结构：

* [CosAssetDataScanDetail](https://cloud.tencent.com/document/api/664/90825#CosAssetDataScanDetail)

	* 新增成员：RuleDetails, LevelDetails, IsFullScanned

* [CosAssetFileIdentifyInfo](https://cloud.tencent.com/document/api/664/90825#CosAssetFileIdentifyInfo)

	* 新增成员：BucketInfo, LastScanTime, HandleStatus, ResultId, RuleDetails, LevelDetails, ComplianceId

* [CosBucketBillingInfo](https://cloud.tencent.com/document/api/664/90825#CosBucketBillingInfo)

	* 新增成员：DefaultSampleRate, BucketSamplingRateWhitelist

* [CosBucketTaskInfo](https://cloud.tencent.com/document/api/664/90825#CosBucketTaskInfo)

	* 新增成员：IdentifyFileCount, SensitiveFileCount

* [CosIdentifyCategoryDetail](https://cloud.tencent.com/document/api/664/90825#CosIdentifyCategoryDetail)

* [CosOverview](https://cloud.tencent.com/document/api/664/90825#CosOverview)

	* 新增成员：HighLevelSensitiveFileCount




## TDSQL MySQL 版(dcdb) 版本：2018-04-11

### 第 83 次发布

发布时间：2026-08-26 01:33:20

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [UpgradeDCDBInstance](https://cloud.tencent.com/document/api/557/16136)

	* 新增入参：SwitchInterval

* [UpgradeDedicatedDCDBInstance](https://cloud.tencent.com/document/api/557/92616)

	* 新增入参：SwitchInterval

* [UpgradeHourDCDBInstance](https://cloud.tencent.com/document/api/557/85926)

	* 新增入参：SwitchInterval




## 弹性 MapReduce(emr) 版本：2019-01-03

### 第 151 次发布

发布时间：2026-08-26 01:38:43

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateCloudInstance](https://cloud.tencent.com/document/api/589/113701)

	* 新增入参：ComputeResourceGroupIds, TerminateProtection


新增数据结构：

* [GpuImageDriverSpec](https://cloud.tencent.com/document/api/589/33981#GpuImageDriverSpec)

修改数据结构：

* [DynamicInstanceForm](https://cloud.tencent.com/document/api/589/33981#DynamicInstanceForm)

	* 新增成员：EnableHistoryServer

* [Resource](https://cloud.tencent.com/document/api/589/33981#Resource)

	* 新增成员：CustomNodeName, GpuImageDriver

* [UserManagerUserBriefInfo](https://cloud.tencent.com/document/api/589/33981#UserManagerUserBriefInfo)

	* 新增成员：Groups, Uin, State, DisplayPasswdUpdateTime, PasswdUpdateTime, PasswdUsedDay, PasswdUsedHour




## 腾讯云可观测平台(monitor) 版本：2023-06-16

### 第 10 次发布

发布时间：2026-08-26 01:55:00

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CancelAIWorkbenchChat](https://cloud.tencent.com/document/api/248/136950)
* [CreateAIWorkbenchAgent](https://cloud.tencent.com/document/api/248/136949)
* [CreateAIWorkbenchTask](https://cloud.tencent.com/document/api/248/136948)
* [DeleteAIWorkbenchAgent](https://cloud.tencent.com/document/api/248/136947)
* [DeleteAIWorkbenchTask](https://cloud.tencent.com/document/api/248/136946)
* [DescribeAIWorkbenchAgent](https://cloud.tencent.com/document/api/248/136945)
* [DescribeAIWorkbenchArtifact](https://cloud.tencent.com/document/api/248/136944)
* [DescribeAIWorkbenchExecution](https://cloud.tencent.com/document/api/248/136943)
* [DescribeAIWorkbenchSession](https://cloud.tencent.com/document/api/248/136942)
* [DescribeAIWorkbenchSkill](https://cloud.tencent.com/document/api/248/136941)
* [GetAIWorkbenchArtifactDownloadURL](https://cloud.tencent.com/document/api/248/136940)
* [ListAIWorkbenchAgents](https://cloud.tencent.com/document/api/248/136939)
* [ListAIWorkbenchArtifacts](https://cloud.tencent.com/document/api/248/136938)
* [ListAIWorkbenchExecutions](https://cloud.tencent.com/document/api/248/136937)
* [ListAIWorkbenchMCPs](https://cloud.tencent.com/document/api/248/136936)
* [ListAIWorkbenchMessages](https://cloud.tencent.com/document/api/248/136935)
* [ListAIWorkbenchResourceInstances](https://cloud.tencent.com/document/api/248/136934)
* [ListAIWorkbenchResourceMaps](https://cloud.tencent.com/document/api/248/136933)
* [ListAIWorkbenchSessions](https://cloud.tencent.com/document/api/248/136932)
* [ListAIWorkbenchSkills](https://cloud.tencent.com/document/api/248/136931)
* [ListAIWorkbenchTasks](https://cloud.tencent.com/document/api/248/136930)
* [TriggerAIWorkbenchTask](https://cloud.tencent.com/document/api/248/136929)
* [UpdateAIWorkbenchAgent](https://cloud.tencent.com/document/api/248/136928)

新增数据结构：

* [AgentInfo](https://cloud.tencent.com/document/api/248/115881#AgentInfo)
* [ArtifactInfo](https://cloud.tencent.com/document/api/248/115881#ArtifactInfo)
* [ContentBlockInfo](https://cloud.tencent.com/document/api/248/115881#ContentBlockInfo)
* [EnvEntry](https://cloud.tencent.com/document/api/248/115881#EnvEntry)
* [EnvVar](https://cloud.tencent.com/document/api/248/115881#EnvVar)
* [ExecutionInfo](https://cloud.tencent.com/document/api/248/115881#ExecutionInfo)
* [InstructionConfig](https://cloud.tencent.com/document/api/248/115881#InstructionConfig)
* [MCPInfo](https://cloud.tencent.com/document/api/248/115881#MCPInfo)
* [MessageInfo](https://cloud.tencent.com/document/api/248/115881#MessageInfo)
* [PageByNumParams](https://cloud.tencent.com/document/api/248/115881#PageByNumParams)
* [PageByNumResult](https://cloud.tencent.com/document/api/248/115881#PageByNumResult)
* [ResourceInstance](https://cloud.tencent.com/document/api/248/115881#ResourceInstance)
* [ResourceMapInfo](https://cloud.tencent.com/document/api/248/115881#ResourceMapInfo)
* [SessionInfo](https://cloud.tencent.com/document/api/248/115881#SessionInfo)
* [SkillInfo](https://cloud.tencent.com/document/api/248/115881#SkillInfo)
* [Tag](https://cloud.tencent.com/document/api/248/115881#Tag)
* [TaskInfo](https://cloud.tencent.com/document/api/248/115881#TaskInfo)



## 腾讯云可观测平台(monitor) 版本：2018-07-24

### 第 167 次发布

发布时间：2026-08-26 01:53:56

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateAlarmHistoryShield](https://cloud.tencent.com/document/api/248/136927)
* [DeleteAlarmHistoryShields](https://cloud.tencent.com/document/api/248/136926)
* [DescribeAlarmHistoryShield](https://cloud.tencent.com/document/api/248/136925)
* [ModifyAlarmHistoryShield](https://cloud.tencent.com/document/api/248/136924)

新增数据结构：

* [ShieldMetric](https://cloud.tencent.com/document/api/248/30354#ShieldMetric)

修改数据结构：

* [PrometheusClusterAgentBasic](https://cloud.tencent.com/document/api/248/30354#PrometheusClusterAgentBasic)

	* 新增成员：CollectAll




## 媒体处理(mps) 版本：2019-06-12

### 第 237 次发布

发布时间：2026-08-26 01:55:24

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [MediaUsageItem](https://cloud.tencent.com/document/api/862/37615#MediaUsageItem)

修改数据结构：

* [MediaTranscodeItem](https://cloud.tencent.com/document/api/862/37615#MediaTranscodeItem)

	* 新增成员：Usage




## Web 应用防火墙(waf) 版本：2018-01-25

### 第 164 次发布

发布时间：2026-08-26 02:22:36

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeLLMContentSecCheck](https://cloud.tencent.com/document/api/627/129968)

	* 新增入参：ClientIP




## 数据开发治理平台 WeData(wedata) 版本：2025-08-06



## 数据开发治理平台 WeData(wedata) 版本：2021-08-20

### 第 203 次发布

发布时间：2026-08-26 02:24:04

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeTableContentPreview](https://cloud.tencent.com/document/api/1267/129275)

	* 新增入参：DatasourceId




