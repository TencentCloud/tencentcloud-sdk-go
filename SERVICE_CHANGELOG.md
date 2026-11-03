# Release v1.3.206

## 运维安全中心（堡垒机）(bh) 版本：2023-04-18

### 第 38 次发布

发布时间：2026-11-03 01:10:05

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateChangePwdTask](https://cloud.tencent.com/document/api/1025/106579)

	* 新增入参：Category, Protocol, Port

* [ModifyChangePwdTask](https://cloud.tencent.com/document/api/1025/106575)

	* 新增入参：Category, Protocol, Port


修改数据结构：

* [ChangePwdTaskInfo](https://cloud.tencent.com/document/api/1025/74416#ChangePwdTaskInfo)

	* 新增成员：Category, Protocol, Port

* [TaskResult](https://cloud.tencent.com/document/api/1025/74416#TaskResult)

	* 新增成员：Category




## 云原生智能网关(cngw) 版本：2023-04-18

### 第 12 次发布

发布时间：2026-11-03 01:26:43

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateAIGWCustomGuardrailPolicy](https://cloud.tencent.com/document/api/1826/140460)
* [CreateOrModifyCloudNativeAPIGatewayIPRestriction](https://cloud.tencent.com/document/api/1826/140459)
* [DeleteAIGWCustomGuardrailPolicy](https://cloud.tencent.com/document/api/1826/140458)
* [DeleteCloudNativeAPIGatewayIPRestriction](https://cloud.tencent.com/document/api/1826/140457)
* [DescribeAIGWCustomGuardrailBindableResources](https://cloud.tencent.com/document/api/1826/140456)
* [DescribeAIGWCustomGuardrailCapabilities](https://cloud.tencent.com/document/api/1826/140455)
* [DescribeAIGWCustomGuardrailHits](https://cloud.tencent.com/document/api/1826/140454)
* [DescribeAIGWCustomGuardrailPolicy](https://cloud.tencent.com/document/api/1826/140453)
* [DescribeAIGWCustomGuardrailPolicyList](https://cloud.tencent.com/document/api/1826/140452)
* [DescribeCloudNativeAPIGatewayIPRestriction](https://cloud.tencent.com/document/api/1826/140451)
* [ExportAIGWCustomGuardrailHits](https://cloud.tencent.com/document/api/1826/140450)
* [ModifyAIGWCustomGuardrailPolicy](https://cloud.tencent.com/document/api/1826/140449)

新增数据结构：

* [CreateAIGWCustomGuardrailPolicyResult](https://cloud.tencent.com/document/api/1826/133161#CreateAIGWCustomGuardrailPolicyResult)
* [CustomGuardrailAuthConfig](https://cloud.tencent.com/document/api/1826/133161#CustomGuardrailAuthConfig)
* [CustomGuardrailAuthConfigInput](https://cloud.tencent.com/document/api/1826/133161#CustomGuardrailAuthConfigInput)
* [CustomGuardrailAuthConfigPatch](https://cloud.tencent.com/document/api/1826/133161#CustomGuardrailAuthConfigPatch)
* [CustomGuardrailBehaviorConfig](https://cloud.tencent.com/document/api/1826/133161#CustomGuardrailBehaviorConfig)
* [CustomGuardrailBehaviorConfigInput](https://cloud.tencent.com/document/api/1826/133161#CustomGuardrailBehaviorConfigInput)
* [CustomGuardrailBehaviorConfigPatch](https://cloud.tencent.com/document/api/1826/133161#CustomGuardrailBehaviorConfigPatch)
* [CustomGuardrailBindableResource](https://cloud.tencent.com/document/api/1826/133161#CustomGuardrailBindableResource)
* [CustomGuardrailCapabilityLimits](https://cloud.tencent.com/document/api/1826/133161#CustomGuardrailCapabilityLimits)
* [CustomGuardrailDecisionConfig](https://cloud.tencent.com/document/api/1826/133161#CustomGuardrailDecisionConfig)
* [CustomGuardrailDecisionConfigInput](https://cloud.tencent.com/document/api/1826/133161#CustomGuardrailDecisionConfigInput)
* [CustomGuardrailDecisionConfigPatch](https://cloud.tencent.com/document/api/1826/133161#CustomGuardrailDecisionConfigPatch)
* [CustomGuardrailFieldMapping](https://cloud.tencent.com/document/api/1826/133161#CustomGuardrailFieldMapping)
* [CustomGuardrailFieldMappingInput](https://cloud.tencent.com/document/api/1826/133161#CustomGuardrailFieldMappingInput)
* [CustomGuardrailFieldMappingPatch](https://cloud.tencent.com/document/api/1826/133161#CustomGuardrailFieldMappingPatch)
* [CustomGuardrailFilterConfig](https://cloud.tencent.com/document/api/1826/133161#CustomGuardrailFilterConfig)
* [CustomGuardrailFilterConfigInput](https://cloud.tencent.com/document/api/1826/133161#CustomGuardrailFilterConfigInput)
* [CustomGuardrailFilterConfigPatch](https://cloud.tencent.com/document/api/1826/133161#CustomGuardrailFilterConfigPatch)
* [CustomGuardrailHit](https://cloud.tencent.com/document/api/1826/133161#CustomGuardrailHit)
* [CustomGuardrailHitFilter](https://cloud.tencent.com/document/api/1826/133161#CustomGuardrailHitFilter)
* [CustomGuardrailPolicy](https://cloud.tencent.com/document/api/1826/133161#CustomGuardrailPolicy)
* [CustomGuardrailPolicyBinding](https://cloud.tencent.com/document/api/1826/133161#CustomGuardrailPolicyBinding)
* [CustomGuardrailPolicyFilter](https://cloud.tencent.com/document/api/1826/133161#CustomGuardrailPolicyFilter)
* [CustomGuardrailResourceRef](https://cloud.tencent.com/document/api/1826/133161#CustomGuardrailResourceRef)
* [CustomGuardrailResourceRefInput](https://cloud.tencent.com/document/api/1826/133161#CustomGuardrailResourceRefInput)
* [CustomGuardrailScope](https://cloud.tencent.com/document/api/1826/133161#CustomGuardrailScope)
* [CustomGuardrailScopeInput](https://cloud.tencent.com/document/api/1826/133161#CustomGuardrailScopeInput)
* [CustomGuardrailServiceConfig](https://cloud.tencent.com/document/api/1826/133161#CustomGuardrailServiceConfig)
* [CustomGuardrailServiceConfigInput](https://cloud.tencent.com/document/api/1826/133161#CustomGuardrailServiceConfigInput)
* [CustomGuardrailServiceConfigPatch](https://cloud.tencent.com/document/api/1826/133161#CustomGuardrailServiceConfigPatch)
* [CustomGuardrailSyncFailure](https://cloud.tencent.com/document/api/1826/133161#CustomGuardrailSyncFailure)
* [DeleteAIGWCustomGuardrailPolicyResult](https://cloud.tencent.com/document/api/1826/133161#DeleteAIGWCustomGuardrailPolicyResult)
* [DescribeAIGWCustomGuardrailBindableResourcesResult](https://cloud.tencent.com/document/api/1826/133161#DescribeAIGWCustomGuardrailBindableResourcesResult)
* [DescribeAIGWCustomGuardrailCapabilitiesResult](https://cloud.tencent.com/document/api/1826/133161#DescribeAIGWCustomGuardrailCapabilitiesResult)
* [DescribeAIGWCustomGuardrailHitsResult](https://cloud.tencent.com/document/api/1826/133161#DescribeAIGWCustomGuardrailHitsResult)
* [DescribeAIGWCustomGuardrailPolicyListResult](https://cloud.tencent.com/document/api/1826/133161#DescribeAIGWCustomGuardrailPolicyListResult)
* [DescribeAIGWCustomGuardrailPolicyResult](https://cloud.tencent.com/document/api/1826/133161#DescribeAIGWCustomGuardrailPolicyResult)
* [DescribeKongIpRestrictionResult](https://cloud.tencent.com/document/api/1826/133161#DescribeKongIpRestrictionResult)
* [ExportAIGWCustomGuardrailHitsResult](https://cloud.tencent.com/document/api/1826/133161#ExportAIGWCustomGuardrailHitsResult)
* [ModifyAIGWCustomGuardrailPolicyResult](https://cloud.tencent.com/document/api/1826/133161#ModifyAIGWCustomGuardrailPolicyResult)



## 云安全一体化平台(csip) 版本：2022-11-21

### 第 114 次发布

发布时间：2026-11-03 01:28:15

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeClusterSummary](https://cloud.tencent.com/document/api/664/136217)

	* 新增出参：NormalNodeCount

* [DescribeImageRegistryAssetOverview](https://cloud.tencent.com/document/api/664/136642)

	* 新增出参：RiskImageCnt, CICDImageCnt, CICDImageScannedCnt, RiskCICDImageCnt, SecretFileCnt, BuildRiskCnt

* [DescribeMalwareTimingScanSetting](https://cloud.tencent.com/document/api/664/136267)

	* 新增出参：QuaraTypes, QuaraTypeDict

* [ModifyEDRRule](https://cloud.tencent.com/document/api/664/133466)

	* 新增入参：K8sApiRule

* [ModifyMalwareTimingScanSettings](https://cloud.tencent.com/document/api/664/136247)

	* 新增入参：QuaraTypes

* [ModifyPayConfig](https://cloud.tencent.com/document/api/664/136525)

	* 新增入参：ContainerConfig


新增数据结构：

* [ContainerAutoScaleConfig](https://cloud.tencent.com/document/api/664/90825#ContainerAutoScaleConfig)
* [QuaraTypeDictItem](https://cloud.tencent.com/document/api/664/90825#QuaraTypeDictItem)
* [RuleContentK8sAPI](https://cloud.tencent.com/document/api/664/90825#RuleContentK8sAPI)

修改数据结构：

* [AssetComponent](https://cloud.tencent.com/document/api/664/90825#AssetComponent)

	* 新增成员：AffectedRegistryImageCount, AffectedCICDImageCount

* [ClusterListItem](https://cloud.tencent.com/document/api/664/90825#ClusterListItem)

	* 新增成员：NormalNodeCount, AuditLogFailedInfo

* [EDRRule](https://cloud.tencent.com/document/api/664/90825#EDRRule)

	* 新增成员：K8sApiRule

* [HostAutoScaleConfig](https://cloud.tencent.com/document/api/664/90825#HostAutoScaleConfig)

	* 新增成员：AssetRange, TagIds

* [IaCFile](https://cloud.tencent.com/document/api/664/90825#IaCFile)

	* 新增成员：FailReason, FailSolution

* [ImageAssetDetail](https://cloud.tencent.com/document/api/664/90825#ImageAssetDetail)

	* 新增成员：SecretFileCnt, BuildRiskCnt

* [ImageAssetInfo](https://cloud.tencent.com/document/api/664/90825#ImageAssetInfo)

	* 新增成员：CICDName, CICDId, IsCICD, SecretFileCnt, BuildRiskCnt, HighLevelSecretFileCnt, CriticalLevelSecretFileCnt, HighLevelBuildRiskCnt, CriticalLevelBuildRiskCnt

* [ImageLayer](https://cloud.tencent.com/document/api/664/90825#ImageLayer)

	* 新增成员：SecretFileCnt, BuildRiskCnt

* [ImageSensitiveInfo](https://cloud.tencent.com/document/api/664/90825#ImageSensitiveInfo)

	* 新增成员：AffectedRegistryImageCount, AffectedCICDImageCount

* [ImageVirus](https://cloud.tencent.com/document/api/664/90825#ImageVirus)

	* 新增成员：AffectedImageCount, AffectedRegistryImageCount, AffectedCICDImageCount

* [ImageVirusWhitelist](https://cloud.tencent.com/document/api/664/90825#ImageVirusWhitelist)

	* 新增成员：Md5

* [ImageVulSummary](https://cloud.tencent.com/document/api/664/90825#ImageVulSummary)

	* 新增成员：AffectedRegistryImageCount, AffectedCICDImageCount




## 暴露面管理服务(ctem) 版本：2023-11-28

### 第 24 次发布

发布时间：2026-11-03 01:34:12

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateSeeds](https://cloud.tencent.com/document/api/1755/125848)

	* 新增入参：KeywordPurposes


修改数据结构：

* [DisplaySeed](https://cloud.tencent.com/document/api/1755/120320#DisplaySeed)

	* 新增成员：Purposes




## 主机安全(cwp) 版本：2018-02-28

### 第 172 次发布

发布时间：2026-11-03 01:36:18

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeMalwareTimingScanSetting](https://cloud.tencent.com/document/api/296/58240)

	* 新增出参：QuaraTypes, QuaraTypeDict

* [ModifyMalwareTimingScanSettings](https://cloud.tencent.com/document/api/296/52509)

	* 新增入参：QuaraTypes


新增数据结构：

* [QuaraTypeDictItem](https://cloud.tencent.com/document/api/296/19867#QuaraTypeDictItem)



## TDSQL-C MySQL 版(cynosdb) 版本：2019-01-07

### 第 201 次发布

发布时间：2026-11-03 01:40:10

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeClusterDatabases](https://cloud.tencent.com/document/api/1003/101559)

	* 新增入参：FuzzyDBNames




## 大数据智能体工作台DataBuddy(databuddy) 版本：2026-07-15

### 第 12 次发布

发布时间：2026-11-03 01:42:22

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateLocation](https://cloud.tencent.com/document/api/1835/140012)

	* 新增入参：FileEventSubscription

* [UpdateLocation](https://cloud.tencent.com/document/api/1835/139995)

	* 新增入参：FileEventSubscription


新增数据结构：

* [CosEventSubscription](https://cloud.tencent.com/document/api/1835/138006#CosEventSubscription)
* [ListLocationItem](https://cloud.tencent.com/document/api/1835/138006#ListLocationItem)

修改数据结构：

* [GetLocationRsp](https://cloud.tencent.com/document/api/1835/138006#GetLocationRsp)

	* 新增成员：FileEventSubscription

* [ListLocationsRsp](https://cloud.tencent.com/document/api/1835/138006#ListLocationsRsp)

	* 新增成员：Items

* [UpdateLocationRsp](https://cloud.tencent.com/document/api/1835/138006#UpdateLocationRsp)

	* 新增成员：FileEventSubscription




## 专线接入(dc) 版本：2018-04-10

### 第 41 次发布

发布时间：2026-11-03 01:44:58

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DeleteDirectConnect](https://cloud.tencent.com/document/api/216/34828)

	* 新增入参：PlannedRepealDate




## 弹性 MapReduce(emr) 版本：2019-01-03

### 第 161 次发布

发布时间：2026-11-03 01:54:39

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeInstancesList](https://cloud.tencent.com/document/api/589/74070)

	* 新增入参：NeedMcpPasswordInfo


新增数据结构：

* [McpPasswordInfo](https://cloud.tencent.com/document/api/589/33981#McpPasswordInfo)

修改数据结构：

* [EmrListInstance](https://cloud.tencent.com/document/api/589/33981#EmrListInstance)

	* 新增成员：Soft, IsTerminateProtection, McpPasswordInfo




## 防火墙管理(fwm) 版本：2025-06-11

### 第 7 次发布

发布时间：2026-11-03 01:59:56

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [OrganMemberItem](https://cloud.tencent.com/document/api/1812/133595#OrganMemberItem)

	* 新增成员：CfwSharerMemberId, SgFreeClaimed

* [StrategyResp](https://cloud.tencent.com/document/api/1812/133595#StrategyResp)

	* 新增成员：Product




## 高性能应用服务(hai) 版本：2023-08-12

### 第 33 次发布

发布时间：2026-11-03 02:02:58

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [RoleStatusInfo](https://cloud.tencent.com/document/api/1721/101518#RoleStatusInfo)

修改数据结构：

* [ServiceDetail](https://cloud.tencent.com/document/api/1721/101518#ServiceDetail)

	* 新增成员：RoleStatuses




## SSL 证书(ssl) 版本：2019-12-05

### 第 100 次发布

发布时间：2026-11-03 02:35:05

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [BatchDeleteCSR](https://cloud.tencent.com/document/api/400/140466)
* [CreateCSR](https://cloud.tencent.com/document/api/400/140465)
* [DescribeCSR](https://cloud.tencent.com/document/api/400/140464)
* [DescribeCSRSet](https://cloud.tencent.com/document/api/400/140463)
* [ModifyCSR](https://cloud.tencent.com/document/api/400/140462)

修改接口：

* [DescribeCertificate](https://cloud.tencent.com/document/api/400/41674)

	* 新增出参：ResubmitAble

* [DescribeCertificateDetail](https://cloud.tencent.com/document/api/400/41673)

	* 新增出参：ResubmitAble


新增数据结构：

* [CSRItem](https://cloud.tencent.com/document/api/400/41679#CSRItem)

修改数据结构：

* [Certificates](https://cloud.tencent.com/document/api/400/41679#Certificates)

	* 新增成员：ResubmitAble




## TSF-Polaris&ZK&网关(tse) 版本：2020-12-07

### 第 116 次发布

发布时间：2026-11-03 02:53:31

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateAIGWCustomGuardrailPolicy](https://cloud.tencent.com/document/api/1364/140475)
* [DeleteAIGWCustomGuardrailPolicy](https://cloud.tencent.com/document/api/1364/140474)
* [DescribeAIGWCustomGuardrailBindableResources](https://cloud.tencent.com/document/api/1364/140473)
* [DescribeAIGWCustomGuardrailCapabilities](https://cloud.tencent.com/document/api/1364/140472)
* [DescribeAIGWCustomGuardrailHits](https://cloud.tencent.com/document/api/1364/140471)
* [DescribeAIGWCustomGuardrailPolicy](https://cloud.tencent.com/document/api/1364/140470)
* [DescribeAIGWCustomGuardrailPolicyList](https://cloud.tencent.com/document/api/1364/140469)
* [ExportAIGWCustomGuardrailHits](https://cloud.tencent.com/document/api/1364/140468)
* [ModifyAIGWCustomGuardrailPolicy](https://cloud.tencent.com/document/api/1364/140467)

新增数据结构：

* [CreateAIGWCustomGuardrailPolicyResult](https://cloud.tencent.com/document/api/1364/54942#CreateAIGWCustomGuardrailPolicyResult)
* [CustomGuardrailAuthConfig](https://cloud.tencent.com/document/api/1364/54942#CustomGuardrailAuthConfig)
* [CustomGuardrailAuthConfigInput](https://cloud.tencent.com/document/api/1364/54942#CustomGuardrailAuthConfigInput)
* [CustomGuardrailAuthConfigPatch](https://cloud.tencent.com/document/api/1364/54942#CustomGuardrailAuthConfigPatch)
* [CustomGuardrailBehaviorConfig](https://cloud.tencent.com/document/api/1364/54942#CustomGuardrailBehaviorConfig)
* [CustomGuardrailBehaviorConfigInput](https://cloud.tencent.com/document/api/1364/54942#CustomGuardrailBehaviorConfigInput)
* [CustomGuardrailBehaviorConfigPatch](https://cloud.tencent.com/document/api/1364/54942#CustomGuardrailBehaviorConfigPatch)
* [CustomGuardrailBindableResource](https://cloud.tencent.com/document/api/1364/54942#CustomGuardrailBindableResource)
* [CustomGuardrailCapabilityLimits](https://cloud.tencent.com/document/api/1364/54942#CustomGuardrailCapabilityLimits)
* [CustomGuardrailDecisionConfig](https://cloud.tencent.com/document/api/1364/54942#CustomGuardrailDecisionConfig)
* [CustomGuardrailDecisionConfigInput](https://cloud.tencent.com/document/api/1364/54942#CustomGuardrailDecisionConfigInput)
* [CustomGuardrailDecisionConfigPatch](https://cloud.tencent.com/document/api/1364/54942#CustomGuardrailDecisionConfigPatch)
* [CustomGuardrailFieldMapping](https://cloud.tencent.com/document/api/1364/54942#CustomGuardrailFieldMapping)
* [CustomGuardrailFieldMappingInput](https://cloud.tencent.com/document/api/1364/54942#CustomGuardrailFieldMappingInput)
* [CustomGuardrailFieldMappingPatch](https://cloud.tencent.com/document/api/1364/54942#CustomGuardrailFieldMappingPatch)
* [CustomGuardrailFilterConfig](https://cloud.tencent.com/document/api/1364/54942#CustomGuardrailFilterConfig)
* [CustomGuardrailFilterConfigInput](https://cloud.tencent.com/document/api/1364/54942#CustomGuardrailFilterConfigInput)
* [CustomGuardrailFilterConfigPatch](https://cloud.tencent.com/document/api/1364/54942#CustomGuardrailFilterConfigPatch)
* [CustomGuardrailHit](https://cloud.tencent.com/document/api/1364/54942#CustomGuardrailHit)
* [CustomGuardrailHitFilter](https://cloud.tencent.com/document/api/1364/54942#CustomGuardrailHitFilter)
* [CustomGuardrailPolicy](https://cloud.tencent.com/document/api/1364/54942#CustomGuardrailPolicy)
* [CustomGuardrailPolicyBinding](https://cloud.tencent.com/document/api/1364/54942#CustomGuardrailPolicyBinding)
* [CustomGuardrailPolicyFilter](https://cloud.tencent.com/document/api/1364/54942#CustomGuardrailPolicyFilter)
* [CustomGuardrailResourceRef](https://cloud.tencent.com/document/api/1364/54942#CustomGuardrailResourceRef)
* [CustomGuardrailResourceRefInput](https://cloud.tencent.com/document/api/1364/54942#CustomGuardrailResourceRefInput)
* [CustomGuardrailScope](https://cloud.tencent.com/document/api/1364/54942#CustomGuardrailScope)
* [CustomGuardrailScopeInput](https://cloud.tencent.com/document/api/1364/54942#CustomGuardrailScopeInput)
* [CustomGuardrailServiceConfig](https://cloud.tencent.com/document/api/1364/54942#CustomGuardrailServiceConfig)
* [CustomGuardrailServiceConfigInput](https://cloud.tencent.com/document/api/1364/54942#CustomGuardrailServiceConfigInput)
* [CustomGuardrailServiceConfigPatch](https://cloud.tencent.com/document/api/1364/54942#CustomGuardrailServiceConfigPatch)
* [CustomGuardrailSyncFailure](https://cloud.tencent.com/document/api/1364/54942#CustomGuardrailSyncFailure)
* [DeleteAIGWCustomGuardrailPolicyResult](https://cloud.tencent.com/document/api/1364/54942#DeleteAIGWCustomGuardrailPolicyResult)
* [DescribeAIGWCustomGuardrailBindableResourcesResult](https://cloud.tencent.com/document/api/1364/54942#DescribeAIGWCustomGuardrailBindableResourcesResult)
* [DescribeAIGWCustomGuardrailCapabilitiesResult](https://cloud.tencent.com/document/api/1364/54942#DescribeAIGWCustomGuardrailCapabilitiesResult)
* [DescribeAIGWCustomGuardrailHitsResult](https://cloud.tencent.com/document/api/1364/54942#DescribeAIGWCustomGuardrailHitsResult)
* [DescribeAIGWCustomGuardrailPolicyListResult](https://cloud.tencent.com/document/api/1364/54942#DescribeAIGWCustomGuardrailPolicyListResult)
* [DescribeAIGWCustomGuardrailPolicyResult](https://cloud.tencent.com/document/api/1364/54942#DescribeAIGWCustomGuardrailPolicyResult)
* [ExportAIGWCustomGuardrailHitsResult](https://cloud.tencent.com/document/api/1364/54942#ExportAIGWCustomGuardrailHitsResult)
* [MaintenanceWindow](https://cloud.tencent.com/document/api/1364/54942#MaintenanceWindow)
* [ModifyAIGWCustomGuardrailPolicyResult](https://cloud.tencent.com/document/api/1364/54942#ModifyAIGWCustomGuardrailPolicyResult)

修改数据结构：

* [SREInstance](https://cloud.tencent.com/document/api/1364/54942#SREInstance)

	* 新增成员：MaintenanceWindow




## WorkBuddy Enterprise(workbuddyenterprise) 版本：2026-07-09

### 第 7 次发布

发布时间：2026-11-03 03:11:03

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateScheduledTask](https://cloud.tencent.com/document/api/1831/140484)
* [DeleteScheduledTask](https://cloud.tencent.com/document/api/1831/140483)
* [DescribeScheduledTask](https://cloud.tencent.com/document/api/1831/140482)
* [DescribeScheduledTaskList](https://cloud.tencent.com/document/api/1831/140481)
* [DescribeScheduledTaskLog](https://cloud.tencent.com/document/api/1831/140480)
* [DescribeScheduledTaskLogList](https://cloud.tencent.com/document/api/1831/140479)
* [ModifyScheduledTask](https://cloud.tencent.com/document/api/1831/140478)
* [TriggerScheduledTask](https://cloud.tencent.com/document/api/1831/140477)

新增数据结构：

* [AgentConfig](https://cloud.tencent.com/document/api/1831/138580#AgentConfig)
* [DeliveryTarget](https://cloud.tencent.com/document/api/1831/138580#DeliveryTarget)
* [Header](https://cloud.tencent.com/document/api/1831/138580#Header)
* [ScheduledTaskLogStats](https://cloud.tencent.com/document/api/1831/138580#ScheduledTaskLogStats)
* [ScheduledTaskLogSummary](https://cloud.tencent.com/document/api/1831/138580#ScheduledTaskLogSummary)
* [ScheduledTaskSummary](https://cloud.tencent.com/document/api/1831/138580#ScheduledTaskSummary)
* [TargetOptions](https://cloud.tencent.com/document/api/1831/138580#TargetOptions)



