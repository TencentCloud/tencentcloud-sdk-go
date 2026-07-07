# Release v1.3.130

## 腾讯混元生3D(ai3d) 版本：2025-05-13

### 第 17 次发布

发布时间：2026-07-08 01:09:13

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [SubmitTextureTo3DJob](https://cloud.tencent.com/document/api/1804/126292)

	* 新增入参：EnableKeepUV, TextureSize




## 应用型负载均衡(alb) 版本：2025-10-30

### 第 1 次发布

发布时间：2026-07-07 10:57:21

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [AddTargetsToTargetGroup](https://cloud.tencent.com/document/api/1822/133696)
* [AssociateBandwidthPackageWithLoadBalancer](https://cloud.tencent.com/document/api/1822/133732)
* [AssociateListenerAdditionalCertificates](https://cloud.tencent.com/document/api/1822/133706)
* [CreateHealthCheckTemplate](https://cloud.tencent.com/document/api/1822/133737)
* [CreateListener](https://cloud.tencent.com/document/api/1822/133705)
* [CreateLoadBalancer](https://cloud.tencent.com/document/api/1822/133713)
* [CreateRules](https://cloud.tencent.com/document/api/1822/133686)
* [CreateSecurityPolicy](https://cloud.tencent.com/document/api/1822/133721)
* [CreateTargetGroup](https://cloud.tencent.com/document/api/1822/133695)
* [DeleteHealthCheckTemplates](https://cloud.tencent.com/document/api/1822/133736)
* [DeleteListener](https://cloud.tencent.com/document/api/1822/133704)
* [DeleteLoadBalancers](https://cloud.tencent.com/document/api/1822/133712)
* [DeleteRules](https://cloud.tencent.com/document/api/1822/133685)
* [DeleteSecurityPolicy](https://cloud.tencent.com/document/api/1822/133720)
* [DeleteTargetGroups](https://cloud.tencent.com/document/api/1822/133694)
* [DescribeAsyncJobs](https://cloud.tencent.com/document/api/1822/133729)
* [DescribeHealthCheckTemplates](https://cloud.tencent.com/document/api/1822/133735)
* [DescribeListenerCertificates](https://cloud.tencent.com/document/api/1822/133703)
* [DescribeListenerDetail](https://cloud.tencent.com/document/api/1822/133702)
* [DescribeListenerHealthStatus](https://cloud.tencent.com/document/api/1822/133701)
* [DescribeListeners](https://cloud.tencent.com/document/api/1822/133700)
* [DescribeLoadBalancerDetail](https://cloud.tencent.com/document/api/1822/133711)
* [DescribeLoadBalancers](https://cloud.tencent.com/document/api/1822/133710)
* [DescribeQuota](https://cloud.tencent.com/document/api/1822/133728)
* [DescribeRules](https://cloud.tencent.com/document/api/1822/133684)
* [DescribeSecurityPolicies](https://cloud.tencent.com/document/api/1822/133719)
* [DescribeSecurityPolicyCapabilities](https://cloud.tencent.com/document/api/1822/133718)
* [DescribeSecurityPolicyRelations](https://cloud.tencent.com/document/api/1822/133717)
* [DescribeSystemSecurityPolicies](https://cloud.tencent.com/document/api/1822/133716)
* [DescribeTargetGroupTargets](https://cloud.tencent.com/document/api/1822/133693)
* [DescribeTargetGroups](https://cloud.tencent.com/document/api/1822/133692)
* [DescribeTargetGroupsByTarget](https://cloud.tencent.com/document/api/1822/133691)
* [DescribeZones](https://cloud.tencent.com/document/api/1822/133727)
* [DisassociateBandwidthPackageFromLoadBalancer](https://cloud.tencent.com/document/api/1822/133731)
* [DisassociateListenerAdditionalCertificates](https://cloud.tencent.com/document/api/1822/133699)
* [InquirePriceCreateLoadBalancer](https://cloud.tencent.com/document/api/1822/133726)
* [ModifyHealthCheckTemplate](https://cloud.tencent.com/document/api/1822/133734)
* [ModifyListenerAttributes](https://cloud.tencent.com/document/api/1822/133698)
* [ModifyLoadBalancerAddressType](https://cloud.tencent.com/document/api/1822/133709)
* [ModifyLoadBalancerAttributes](https://cloud.tencent.com/document/api/1822/133708)
* [ModifyLoadBalancerModificationProtection](https://cloud.tencent.com/document/api/1822/133725)
* [ModifyRulesAttributes](https://cloud.tencent.com/document/api/1822/133683)
* [ModifySecurityPolicyAttributes](https://cloud.tencent.com/document/api/1822/133715)
* [ModifyTargetGroupAttributes](https://cloud.tencent.com/document/api/1822/133690)
* [ModifyTargetsInTargetGroup](https://cloud.tencent.com/document/api/1822/133689)
* [NotifyUnbindTarget](https://cloud.tencent.com/document/api/1822/133724)
* [RemoveTargetsFromTargetGroup](https://cloud.tencent.com/document/api/1822/133688)
* [SetLoadBalancerSecurityGroups](https://cloud.tencent.com/document/api/1822/133723)

新增数据结构：

* [AccessLogConfig](https://cloud.tencent.com/document/api/1822/133738#AccessLogConfig)
* [CertificateInfo](https://cloud.tencent.com/document/api/1822/133738#CertificateInfo)
* [DefaultAction](https://cloud.tencent.com/document/api/1822/133738#DefaultAction)
* [DeletionProtectionConfig](https://cloud.tencent.com/document/api/1822/133738#DeletionProtectionConfig)
* [Filter](https://cloud.tencent.com/document/api/1822/133738#Filter)
* [FixedResponseInfo](https://cloud.tencent.com/document/api/1822/133738#FixedResponseInfo)
* [HTTPCookieInfo](https://cloud.tencent.com/document/api/1822/133738#HTTPCookieInfo)
* [HTTPHeaderInfo](https://cloud.tencent.com/document/api/1822/133738#HTTPHeaderInfo)
* [HTTPQueryStringInfo](https://cloud.tencent.com/document/api/1822/133738#HTTPQueryStringInfo)
* [HTTPRedirectInfo](https://cloud.tencent.com/document/api/1822/133738#HTTPRedirectInfo)
* [HTTPRewriteInfo](https://cloud.tencent.com/document/api/1822/133738#HTTPRewriteInfo)
* [HealthCheckConfig](https://cloud.tencent.com/document/api/1822/133738#HealthCheckConfig)
* [HealthCheckTemplate](https://cloud.tencent.com/document/api/1822/133738#HealthCheckTemplate)
* [IPAddressInfo](https://cloud.tencent.com/document/api/1822/133738#IPAddressInfo)
* [InsertHTTPHeaderInfo](https://cloud.tencent.com/document/api/1822/133738#InsertHTTPHeaderInfo)
* [Job](https://cloud.tencent.com/document/api/1822/133738#Job)
* [ListenerOutput](https://cloud.tencent.com/document/api/1822/133738#ListenerOutput)
* [LoadBalancer](https://cloud.tencent.com/document/api/1822/133738#LoadBalancer)
* [LoadBalancerAddress](https://cloud.tencent.com/document/api/1822/133738#LoadBalancerAddress)
* [LoadBalancerBillingConfig](https://cloud.tencent.com/document/api/1822/133738#LoadBalancerBillingConfig)
* [LoadBalancerDetail](https://cloud.tencent.com/document/api/1822/133738#LoadBalancerDetail)
* [LoadBalancerOperationLocksItem](https://cloud.tencent.com/document/api/1822/133738#LoadBalancerOperationLocksItem)
* [ModificationProtectionInfo](https://cloud.tencent.com/document/api/1822/133738#ModificationProtectionInfo)
* [PostPayPriceInfo](https://cloud.tencent.com/document/api/1822/133738#PostPayPriceInfo)
* [Price](https://cloud.tencent.com/document/api/1822/133738#Price)
* [QuotaInfo](https://cloud.tencent.com/document/api/1822/133738#QuotaInfo)
* [RelatedListener](https://cloud.tencent.com/document/api/1822/133738#RelatedListener)
* [RemoveHTTPHeaderInfo](https://cloud.tencent.com/document/api/1822/133738#RemoveHTTPHeaderInfo)
* [RuleAction](https://cloud.tencent.com/document/api/1822/133738#RuleAction)
* [RuleCondition](https://cloud.tencent.com/document/api/1822/133738#RuleCondition)
* [RuleHealthStatusInfo](https://cloud.tencent.com/document/api/1822/133738#RuleHealthStatusInfo)
* [RuleInput](https://cloud.tencent.com/document/api/1822/133738#RuleInput)
* [RuleModify](https://cloud.tencent.com/document/api/1822/133738#RuleModify)
* [RuleOutput](https://cloud.tencent.com/document/api/1822/133738#RuleOutput)
* [SecurityPolicyCapability](https://cloud.tencent.com/document/api/1822/133738#SecurityPolicyCapability)
* [SecurityPolicyInfo](https://cloud.tencent.com/document/api/1822/133738#SecurityPolicyInfo)
* [SecurityPolicyRelations](https://cloud.tencent.com/document/api/1822/133738#SecurityPolicyRelations)
* [StickySessionConfig](https://cloud.tencent.com/document/api/1822/133738#StickySessionConfig)
* [TagInfo](https://cloud.tencent.com/document/api/1822/133738#TagInfo)
* [TargetGroupConfig](https://cloud.tencent.com/document/api/1822/133738#TargetGroupConfig)
* [TargetGroupHealthInfo](https://cloud.tencent.com/document/api/1822/133738#TargetGroupHealthInfo)
* [TargetGroupOutput](https://cloud.tencent.com/document/api/1822/133738#TargetGroupOutput)
* [TargetGroupStickySession](https://cloud.tencent.com/document/api/1822/133738#TargetGroupStickySession)
* [TargetGroupTuple](https://cloud.tencent.com/document/api/1822/133738#TargetGroupTuple)
* [TargetHealthStatusInfo](https://cloud.tencent.com/document/api/1822/133738#TargetHealthStatusInfo)
* [TargetOutput](https://cloud.tencent.com/document/api/1822/133738#TargetOutput)
* [TargetToAdd](https://cloud.tencent.com/document/api/1822/133738#TargetToAdd)
* [TargetToModify](https://cloud.tencent.com/document/api/1822/133738#TargetToModify)
* [TargetToRemove](https://cloud.tencent.com/document/api/1822/133738#TargetToRemove)
* [XForwardedForConfig](https://cloud.tencent.com/document/api/1822/133738#XForwardedForConfig)
* [Zone](https://cloud.tencent.com/document/api/1822/133738#Zone)
* [ZoneMappingInfo](https://cloud.tencent.com/document/api/1822/133738#ZoneMappingInfo)
* [ZoneMappingsItem](https://cloud.tencent.com/document/api/1822/133738#ZoneMappingsItem)



## 访问管理(cam) 版本：2019-01-16

### 第 71 次发布

发布时间：2026-07-08 01:18:25

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [StrategyInfo](https://cloud.tencent.com/document/api/598/33167#StrategyInfo)

	* 新增成员：PermissionLevel




## 云防火墙(cfw) 版本：2019-09-04

### 第 106 次发布

发布时间：2026-07-07 12:58:27

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeOfflineExportTemporaryCredentials](https://cloud.tencent.com/document/api/1132/133526)

	* 新增出参：Link




## 云托付物理服务器(chc) 版本：2023-04-18

### 第 10 次发布

发布时间：2026-07-08 01:27:44

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeWorkOrderCarCollectList](https://cloud.tencent.com/document/api/1448/133765)
* [DescribeWorkOrderContactCollectList](https://cloud.tencent.com/document/api/1448/133764)
* [DescribeWorkOrderPersonnelCollectList](https://cloud.tencent.com/document/api/1448/133763)

新增数据结构：

* [ContactCollectInfo](https://cloud.tencent.com/document/api/1448/117193#ContactCollectInfo)



## 负载均衡(clb) 版本：2018-03-17

### 第 153 次发布

发布时间：2026-07-08 01:29:46

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [Coefficient](https://cloud.tencent.com/document/api/214/30694#Coefficient)

	* <font color="#dd0000">**修改成员**：</font>InputCachedCoefficient, InputCoefficient, OutputCoefficient


### 第 152 次发布

发布时间：2026-07-07 10:18:20

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [AddModelKey](https://cloud.tencent.com/document/api/214/133681)
* [AddModelRewrite](https://cloud.tencent.com/document/api/214/133639)
* [AssociateModelRouterGuardrails](https://cloud.tencent.com/document/api/214/133645)
* [AssociateModelsToModelRouter](https://cloud.tencent.com/document/api/214/133661)
* [ChatCompletions](https://cloud.tencent.com/document/api/214/133663)
* [CreateBYOKNetwork](https://cloud.tencent.com/document/api/214/133680)
* [CreateIntentRouter](https://cloud.tencent.com/document/api/214/133638)
* [CreateModel](https://cloud.tencent.com/document/api/214/133679)
* [CreateModelRouterResourcePackage](https://cloud.tencent.com/document/api/214/133652)
* [DeleteIntentRouter](https://cloud.tencent.com/document/api/214/133637)
* [DeleteModel](https://cloud.tencent.com/document/api/214/133678)
* [DeregisterModelsFromServiceProvider](https://cloud.tencent.com/document/api/214/133677)
* [DescribeAssociatedModelAvailability](https://cloud.tencent.com/document/api/214/133660)
* [DescribeIntentRouterTiers](https://cloud.tencent.com/document/api/214/133636)
* [DescribeIntentRouters](https://cloud.tencent.com/document/api/214/133635)
* [DescribeKeys](https://cloud.tencent.com/document/api/214/133644)
* [DescribeModelAliases](https://cloud.tencent.com/document/api/214/133676)
* [DescribeModelAssociations](https://cloud.tencent.com/document/api/214/133659)
* [DescribeModelKeys](https://cloud.tencent.com/document/api/214/133675)
* [DescribeModelNames](https://cloud.tencent.com/document/api/214/133674)
* [DescribeModelRewrite](https://cloud.tencent.com/document/api/214/133634)
* [DescribeModelRouterGuardrails](https://cloud.tencent.com/document/api/214/133643)
* [DescribeModelRouterLogs](https://cloud.tencent.com/document/api/214/133658)
* [DescribeModelRouterResourcePackageDeduction](https://cloud.tencent.com/document/api/214/133651)
* [DescribeModelRouterResourcePackages](https://cloud.tencent.com/document/api/214/133650)
* [DescribeServiceProviderHealthStatus](https://cloud.tencent.com/document/api/214/133673)
* [DescribeSupportedProviders](https://cloud.tencent.com/document/api/214/133654)
* [DescribeUpperModels](https://cloud.tencent.com/document/api/214/133672)
* [DisassociateModelRouterGuardrails](https://cloud.tencent.com/document/api/214/133642)
* [DisassociateModelsFromModelRouter](https://cloud.tencent.com/document/api/214/133657)
* [InquirePriceCreateModelRouterResourcePackage](https://cloud.tencent.com/document/api/214/133649)
* [InquirePriceRefundModelRouterResourcePackage](https://cloud.tencent.com/document/api/214/133648)
* [ModifyIntentRouterAttribute](https://cloud.tencent.com/document/api/214/133633)
* [ModifyModelAliasAttributes](https://cloud.tencent.com/document/api/214/133671)
* [ModifyModelAttributes](https://cloud.tencent.com/document/api/214/133670)
* [ModifyModelRouterGuardrails](https://cloud.tencent.com/document/api/214/133641)
* [ModifyModelRouterSecurityGroups](https://cloud.tencent.com/document/api/214/133656)
* [ModifyServiceProviderModelAttributes](https://cloud.tencent.com/document/api/214/133669)
* [RefundModelRouterResourcePackage](https://cloud.tencent.com/document/api/214/133647)
* [RegisterModelsToServiceProvider](https://cloud.tencent.com/document/api/214/133668)
* [RemoveModelKey](https://cloud.tencent.com/document/api/214/133667)
* [RemoveModelRewrite](https://cloud.tencent.com/document/api/214/133632)
* [TestModelInputModalities](https://cloud.tencent.com/document/api/214/133666)
* [TestServiceProviderConnection](https://cloud.tencent.com/document/api/214/133665)

新增数据结构：

* [AssociateGuardrailConfig](https://cloud.tencent.com/document/api/214/30694#AssociateGuardrailConfig)
* [AssociatedModelRouterItem](https://cloud.tencent.com/document/api/214/30694#AssociatedModelRouterItem)
* [Coefficient](https://cloud.tencent.com/document/api/214/30694#Coefficient)
* [DisassociateGuardrailConfig](https://cloud.tencent.com/document/api/214/30694#DisassociateGuardrailConfig)
* [GuardrailConfig](https://cloud.tencent.com/document/api/214/30694#GuardrailConfig)
* [IntentRouterItem](https://cloud.tencent.com/document/api/214/30694#IntentRouterItem)
* [IntentRouterTierDictItem](https://cloud.tencent.com/document/api/214/30694#IntentRouterTierDictItem)
* [IntentRouterTierItem](https://cloud.tencent.com/document/api/214/30694#IntentRouterTierItem)
* [KeyDetailItem](https://cloud.tencent.com/document/api/214/30694#KeyDetailItem)
* [KeyInfo](https://cloud.tencent.com/document/api/214/30694#KeyInfo)
* [KeyItem](https://cloud.tencent.com/document/api/214/30694#KeyItem)
* [ModalityProbeDetail](https://cloud.tencent.com/document/api/214/30694#ModalityProbeDetail)
* [ModelAlias](https://cloud.tencent.com/document/api/214/30694#ModelAlias)
* [ModelAssociation](https://cloud.tencent.com/document/api/214/30694#ModelAssociation)
* [ModelAvailability](https://cloud.tencent.com/document/api/214/30694#ModelAvailability)
* [ModelHealthCheckResults](https://cloud.tencent.com/document/api/214/30694#ModelHealthCheckResults)
* [ModelItem](https://cloud.tencent.com/document/api/214/30694#ModelItem)
* [ModelKeyInfoItem](https://cloud.tencent.com/document/api/214/30694#ModelKeyInfoItem)
* [ModelNameAggregatedItem](https://cloud.tencent.com/document/api/214/30694#ModelNameAggregatedItem)
* [ModelRouterLog](https://cloud.tencent.com/document/api/214/30694#ModelRouterLog)
* [ModelRouterModel](https://cloud.tencent.com/document/api/214/30694#ModelRouterModel)
* [ModelRouterPackage](https://cloud.tencent.com/document/api/214/30694#ModelRouterPackage)
* [ModelRouterResourcePackageDeduction](https://cloud.tencent.com/document/api/214/30694#ModelRouterResourcePackageDeduction)
* [ModelRouterResourcePackageRefundPrice](https://cloud.tencent.com/document/api/214/30694#ModelRouterResourcePackageRefundPrice)
* [ModelTestResult](https://cloud.tencent.com/document/api/214/30694#ModelTestResult)
* [MultiModalityAttachments](https://cloud.tencent.com/document/api/214/30694#MultiModalityAttachments)
* [ProviderItem](https://cloud.tencent.com/document/api/214/30694#ProviderItem)
* [ProviderTestConnectionErrorInfo](https://cloud.tencent.com/document/api/214/30694#ProviderTestConnectionErrorInfo)
* [RewriteItem](https://cloud.tencent.com/document/api/214/30694#RewriteItem)
* [ServiceProvider](https://cloud.tencent.com/document/api/214/30694#ServiceProvider)
* [ServiceProviderCoefficient](https://cloud.tencent.com/document/api/214/30694#ServiceProviderCoefficient)
* [ServiceProviderItem](https://cloud.tencent.com/document/api/214/30694#ServiceProviderItem)
* [ServiceProviderModelItem](https://cloud.tencent.com/document/api/214/30694#ServiceProviderModelItem)
* [Sort](https://cloud.tencent.com/document/api/214/30694#Sort)
* [TestConnectionRequestInfo](https://cloud.tencent.com/document/api/214/30694#TestConnectionRequestInfo)
* [TierItem](https://cloud.tencent.com/document/api/214/30694#TierItem)

修改数据结构：

* [BudgetAssociation](https://cloud.tencent.com/document/api/214/30694#BudgetAssociation)

	* 新增成员：ResourceName, UserGroupId

* [ModelRouterDetail](https://cloud.tencent.com/document/api/214/30694#ModelRouterDetail)

	* <font color="#dd0000">**删除成员**：</font>CreditUsage




## 云安全一体化平台(csip) 版本：2022-11-21

### 第 89 次发布

发布时间：2026-07-08 01:35:38

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateDspmAssetIdentifyInfoExportJob](https://cloud.tencent.com/document/api/664/133801)
* [CreateDspmIdentifyCategory](https://cloud.tencent.com/document/api/664/133800)
* [CreateDspmIdentifyComplianceCategoryRelation](https://cloud.tencent.com/document/api/664/133799)
* [CreateDspmIdentifyComplianceGroup](https://cloud.tencent.com/document/api/664/133798)
* [CreateDspmIdentifyComplianceGroupCopy](https://cloud.tencent.com/document/api/664/133797)
* [CreateDspmIdentifyComplianceRuleRelation](https://cloud.tencent.com/document/api/664/133796)
* [CreateDspmIdentifyLevelGroup](https://cloud.tencent.com/document/api/664/133795)
* [CreateDspmIdentifyRule](https://cloud.tencent.com/document/api/664/133794)
* [DeleteDspmIdentifyCategory](https://cloud.tencent.com/document/api/664/133793)
* [DeleteDspmIdentifyComplianceCategoryRelation](https://cloud.tencent.com/document/api/664/133792)
* [DeleteDspmIdentifyComplianceGroup](https://cloud.tencent.com/document/api/664/133791)
* [DeleteDspmIdentifyComplianceRuleRelation](https://cloud.tencent.com/document/api/664/133790)
* [DeleteDspmIdentifyLevelGroup](https://cloud.tencent.com/document/api/664/133789)
* [DeleteDspmIdentifyRule](https://cloud.tencent.com/document/api/664/133788)
* [DescribeCosObjectScanTask](https://cloud.tencent.com/document/api/664/133804)
* [DescribeDspmAssetFieldSamples](https://cloud.tencent.com/document/api/664/133787)
* [DescribeDspmAssetIdentifyInfoList](https://cloud.tencent.com/document/api/664/133786)
* [DescribeDspmIdentifyCategoryList](https://cloud.tencent.com/document/api/664/133785)
* [DescribeDspmIdentifyComplianceCategoryRuleList](https://cloud.tencent.com/document/api/664/133784)
* [DescribeDspmIdentifyComplianceGroupDetail](https://cloud.tencent.com/document/api/664/133783)
* [DescribeDspmIdentifyComplianceGroupList](https://cloud.tencent.com/document/api/664/133782)
* [DescribeDspmIdentifyDistributionStatistics](https://cloud.tencent.com/document/api/664/133781)
* [DescribeDspmIdentifyLevelGroupList](https://cloud.tencent.com/document/api/664/133780)
* [DescribeDspmIdentifyRuleDetail](https://cloud.tencent.com/document/api/664/133779)
* [DescribeDspmIdentifyRuleList](https://cloud.tencent.com/document/api/664/133778)
* [DescribeDspmIdentifyRuleTestResult](https://cloud.tencent.com/document/api/664/133777)
* [ModifyCosAuditObjectIdentifyStatus](https://cloud.tencent.com/document/api/664/133803)
* [ModifyCosAuditObjectSampleRate](https://cloud.tencent.com/document/api/664/133802)
* [ModifyDspmApplyingIdentifyComplianceGroup](https://cloud.tencent.com/document/api/664/133776)
* [ModifyDspmAssetDataScanTaskStatus](https://cloud.tencent.com/document/api/664/133775)
* [ModifyDspmIdentifyCategory](https://cloud.tencent.com/document/api/664/133774)
* [ModifyDspmIdentifyComplianceGroup](https://cloud.tencent.com/document/api/664/133773)
* [ModifyDspmIdentifyComplianceGroupStatus](https://cloud.tencent.com/document/api/664/133772)
* [ModifyDspmIdentifyComplianceRuleLevelInfo](https://cloud.tencent.com/document/api/664/133771)
* [ModifyDspmIdentifyLevelGroup](https://cloud.tencent.com/document/api/664/133770)
* [ModifyDspmIdentifyLevelItem](https://cloud.tencent.com/document/api/664/133769)
* [ModifyDspmIdentifyRule](https://cloud.tencent.com/document/api/664/133768)
* [ModifyDspmIdentifyRuleStatus](https://cloud.tencent.com/document/api/664/133767)

新增数据结构：

* [DspmAddIdentifyLevelItem](https://cloud.tencent.com/document/api/664/90825#DspmAddIdentifyLevelItem)
* [DspmAssetIdentifyInfo](https://cloud.tencent.com/document/api/664/90825#DspmAssetIdentifyInfo)
* [DspmAssetIdentifyTaskStatus](https://cloud.tencent.com/document/api/664/90825#DspmAssetIdentifyTaskStatus)
* [DspmIdentifyCategoryItem](https://cloud.tencent.com/document/api/664/90825#DspmIdentifyCategoryItem)
* [DspmIdentifyCategoryRuleRelateDetailItem](https://cloud.tencent.com/document/api/664/90825#DspmIdentifyCategoryRuleRelateDetailItem)
* [DspmIdentifyCategoryRuleRelateItem](https://cloud.tencent.com/document/api/664/90825#DspmIdentifyCategoryRuleRelateItem)
* [DspmIdentifyComplianceCategoryRelation](https://cloud.tencent.com/document/api/664/90825#DspmIdentifyComplianceCategoryRelation)
* [DspmIdentifyComplianceItem](https://cloud.tencent.com/document/api/664/90825#DspmIdentifyComplianceItem)
* [DspmIdentifyComplianceRuleRelation](https://cloud.tencent.com/document/api/664/90825#DspmIdentifyComplianceRuleRelation)
* [DspmIdentifyLevelGroupItem](https://cloud.tencent.com/document/api/664/90825#DspmIdentifyLevelGroupItem)
* [DspmIdentifyLevelItem](https://cloud.tencent.com/document/api/664/90825#DspmIdentifyLevelItem)
* [DspmIdentifyRefComplianceInfo](https://cloud.tencent.com/document/api/664/90825#DspmIdentifyRefComplianceInfo)
* [DspmIdentifyRuleItem](https://cloud.tencent.com/document/api/664/90825#DspmIdentifyRuleItem)
* [DspmIdentifyRuleStructuredTestItem](https://cloud.tencent.com/document/api/664/90825#DspmIdentifyRuleStructuredTestItem)
* [DspmStatisticsItem](https://cloud.tencent.com/document/api/664/90825#DspmStatisticsItem)



## 云服务器(cvm) 版本：2017-03-12

### 第 167 次发布

发布时间：2026-07-08 01:39:17

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [RunInstances](https://cloud.tencent.com/document/api/213/15730)

	* 新增入参：PartitionNumber


修改数据结构：

* [DisasterRecoverGroup](https://cloud.tencent.com/document/api/213/15753#DisasterRecoverGroup)

	* 新增成员：Strategy, PartitionCount

* [Instance](https://cloud.tencent.com/document/api/213/15753#Instance)

	* 新增成员：PartitionNumber




## TDSQL-C MySQL 版(cynosdb) 版本：2019-01-07

### 第 179 次发布

发布时间：2026-07-08 01:45:03

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeBackupOverview](https://cloud.tencent.com/document/api/1003/133805)

新增数据结构：

* [BackupVolumeInfo](https://cloud.tencent.com/document/api/1003/48097#BackupVolumeInfo)



## 弹性 MapReduce(emr) 版本：2019-01-03

### 第 148 次发布

发布时间：2026-07-08 01:58:40

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DeleteGroupsSTD](https://cloud.tencent.com/document/api/589/118727)

	* 新增出参：FlowId

* [DeleteUserManagerUserList](https://cloud.tencent.com/document/api/589/83729)

	* 新增出参：FlowId

* [ModifyUserGroup](https://cloud.tencent.com/document/api/589/118726)

	* 新增出参：FlowId

* [ModifyUserManagerPwd](https://cloud.tencent.com/document/api/589/99696)

	* 新增出参：FlowId

* [ModifyUsersOfGroupSTD](https://cloud.tencent.com/document/api/589/118725)

	* 新增出参：FlowId




## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 305 次发布

发布时间：2026-07-08 02:00:57

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateUserAutoSignEnableUrl](https://cloud.tencent.com/document/api/1323/87787)

	* 新增入参：EndPoint

* [CreateUserAutoSignSealUrl](https://cloud.tencent.com/document/api/1323/99316)

	* 新增入参：EndPoint

* [UploadFiles](https://cloud.tencent.com/document/api/1323/73066)

	* 新增入参：Deadline

	* 新增出参：Deadline




## 腾讯电子签（基础版）(essbasic) 版本：2021-05-26

### 第 263 次发布

发布时间：2026-07-08 02:02:41

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ChannelCreateUserAutoSignEnableUrl](https://cloud.tencent.com/document/api/1420/96110)

	* 新增入参：EndPoint

* [ChannelCreateUserAutoSignSealUrl](https://cloud.tencent.com/document/api/1420/99317)

	* 新增入参：EndPoint

* [UploadFiles](https://cloud.tencent.com/document/api/1420/71479)

	* 新增入参：Deadline

	* 新增出参：Deadline




## 腾讯电子签（基础版）(essbasic) 版本：2020-12-22



## 全球加速(ga2) 版本：2025-01-15

### 第 8 次发布

发布时间：2026-07-08 02:04:59

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [IpAddressInfoSet](https://cloud.tencent.com/document/api/1817/130045#IpAddressInfoSet)

	* 新增成员：DdosProtectionType




## 媒体处理(mps) 版本：2019-06-12

### 第 214 次发布

发布时间：2026-07-08 02:27:17

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [SmartSubtitleTaskFullTextSegmentItem](https://cloud.tencent.com/document/api/862/37615#SmartSubtitleTaskFullTextSegmentItem)

修改数据结构：

* [SmartSubtitleTaskTextResultOutput](https://cloud.tencent.com/document/api/862/37615#SmartSubtitleTaskTextResultOutput)

	* 新增成员：SegmentSet




## 云数据库Redis(redis) 版本：2018-04-12

### 第 103 次发布

发布时间：2026-07-08 02:36:51

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateExportTask](https://cloud.tencent.com/document/api/239/129831)

	* 新增出参：FileName




## 云开发 CloudBase(tcb) 版本：2018-06-08

### 第 151 次发布

发布时间：2026-07-08 02:45:01

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [PostgreSQLInfo](https://cloud.tencent.com/document/api/876/34822#PostgreSQLInfo)

	* <font color="#dd0000">**删除成员**：</font>Version




## TI-ONE 训练平台(tione) 版本：2021-11-11

### 第 123 次发布

发布时间：2026-07-08 02:57:36

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribePresetImageList](https://cloud.tencent.com/document/api/851/133806)

新增数据结构：

* [PresetImageInfo](https://cloud.tencent.com/document/api/851/75051#PresetImageInfo)
* [RuntimeLib](https://cloud.tencent.com/document/api/851/75051#RuntimeLib)



## TI-ONE 训练平台(tione) 版本：2019-10-22



## 机器翻译(tmt) 版本：2018-03-21

### 第 17 次发布

发布时间：2026-07-08 03:01:55

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除接口**：</font>

* TextTranslate



## TokenHub(tokenhub) 版本：2026-03-22

### 第 11 次发布

发布时间：2026-07-08 03:01:59

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateApiKey](https://cloud.tencent.com/document/api/1823/133810)
* [DeleteApiKey](https://cloud.tencent.com/document/api/1823/133809)
* [ModifyApiKeyInfo](https://cloud.tencent.com/document/api/1823/133808)
* [ModifyApiKeyStatus](https://cloud.tencent.com/document/api/1823/133807)



