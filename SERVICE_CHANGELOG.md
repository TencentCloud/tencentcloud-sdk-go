# Release v1.3.99

## AI Agent 安全网关(apis) 版本：2024-08-01

### 第 6 次发布

发布时间：2026-05-18 01:07:56

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateModel](https://cloud.tencent.com/document/api/1627/129609)

	* 新增入参：ModelID, Description

* [CreateModelService](https://cloud.tencent.com/document/api/1627/129615)

	* 新增入参：ModelProtocol

* [DescribeModelServices](https://cloud.tencent.com/document/api/1627/129612)

	* 新增入参：ModelProtocol

* [ModifyModel](https://cloud.tencent.com/document/api/1627/129605)

	* 新增入参：ModelID, Description

* [ModifyModelService](https://cloud.tencent.com/document/api/1627/129611)

	* 新增入参：ModelProtocol


修改数据结构：

* [DescribeModelResponseVO](https://cloud.tencent.com/document/api/1627/129635#DescribeModelResponseVO)

	* 新增成员：ModelID, Description

* [DescribeModelServiceResponseVO](https://cloud.tencent.com/document/api/1627/129635#DescribeModelServiceResponseVO)

	* 新增成员：ModelProtocol

* [PromptModerateConfigDTO](https://cloud.tencent.com/document/api/1627/129635#PromptModerateConfigDTO)

	* 新增成员：ContextScope

* [SensitiveDataCheckConfigDTO](https://cloud.tencent.com/document/api/1627/129635#SensitiveDataCheckConfigDTO)

	* 新增成员：ContextScope

* [TmsConfigDTO](https://cloud.tencent.com/document/api/1627/129635#TmsConfigDTO)

	* 新增成员：ContextScope




## 云安全一体化平台(csip) 版本：2022-11-21

### 第 78 次发布

发布时间：2026-05-18 01:29:39

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateCosAssetSyncTask](https://cloud.tencent.com/document/api/664/131698)
* [CreateCosObjectScanTask](https://cloud.tencent.com/document/api/664/131697)
* [CreateCosPolicy](https://cloud.tencent.com/document/api/664/131696)
* [CreateCosRiskScanTask](https://cloud.tencent.com/document/api/664/131695)
* [DeleteCosAkAsset](https://cloud.tencent.com/document/api/664/131694)
* [DeleteCosPolicy](https://cloud.tencent.com/document/api/664/131693)
* [DescribeBucketInvokeIpList](https://cloud.tencent.com/document/api/664/131692)
* [DescribeCosAccessPermission](https://cloud.tencent.com/document/api/664/131691)
* [DescribeCosAccessPermissions](https://cloud.tencent.com/document/api/664/131690)
* [DescribeCosActionList](https://cloud.tencent.com/document/api/664/131689)
* [DescribeCosAkAsset](https://cloud.tencent.com/document/api/664/131688)
* [DescribeCosAkInvokeIpList](https://cloud.tencent.com/document/api/664/131687)
* [DescribeCosAlarmList](https://cloud.tencent.com/document/api/664/131686)
* [DescribeCosAlarmTrendData](https://cloud.tencent.com/document/api/664/131685)
* [DescribeCosAsset](https://cloud.tencent.com/document/api/664/131684)
* [DescribeCosAssetSyncTask](https://cloud.tencent.com/document/api/664/131683)
* [DescribeCosAuditAppIdList](https://cloud.tencent.com/document/api/664/131682)
* [DescribeCosAuditDictionaryList](https://cloud.tencent.com/document/api/664/131681)
* [DescribeCosAuditPayInfo](https://cloud.tencent.com/document/api/664/131680)
* [DescribeCosBucketBillingInfo](https://cloud.tencent.com/document/api/664/131679)
* [DescribeCosBucketList](https://cloud.tencent.com/document/api/664/131678)
* [DescribeCosBucketRisk](https://cloud.tencent.com/document/api/664/131677)
* [DescribeCosIdentifyFileList](https://cloud.tencent.com/document/api/664/131676)
* [DescribeCosInvokeUa](https://cloud.tencent.com/document/api/664/131675)
* [DescribeCosIpInvokeLog](https://cloud.tencent.com/document/api/664/131674)
* [DescribeCosIpInvokeRecordFile](https://cloud.tencent.com/document/api/664/131673)
* [DescribeCosOverview](https://cloud.tencent.com/document/api/664/131672)
* [DescribeCosPolicy](https://cloud.tencent.com/document/api/664/131671)
* [DescribeCosRiskActionList](https://cloud.tencent.com/document/api/664/131670)
* [DescribeCosRiskEvidence](https://cloud.tencent.com/document/api/664/131669)
* [DescribeCosRiskScanTask](https://cloud.tencent.com/document/api/664/131668)
* [DescribeCosRoleAccessPermission](https://cloud.tencent.com/document/api/664/131667)
* [DescribeCosRoleAccessPermissions](https://cloud.tencent.com/document/api/664/131666)
* [DescribeCosSourceIp](https://cloud.tencent.com/document/api/664/131665)
* [DescribeIpInvokeRecord](https://cloud.tencent.com/document/api/664/131664)
* [DescribeIpInvokeRecordDetail](https://cloud.tencent.com/document/api/664/131663)
* [DescribePolicyHitData](https://cloud.tencent.com/document/api/664/131662)
* [DescribeRiskBucketList](https://cloud.tencent.com/document/api/664/131661)
* [DescribeRiskItemList](https://cloud.tencent.com/document/api/664/131660)
* [DescribeRiskTrendData](https://cloud.tencent.com/document/api/664/131659)
* [ModifyAlarmRiskStatus](https://cloud.tencent.com/document/api/664/131658)
* [ModifyCosAuditMonitorAccount](https://cloud.tencent.com/document/api/664/131657)
* [ModifyCosMarkInfo](https://cloud.tencent.com/document/api/664/131656)
* [ModifyPolicyStatus](https://cloud.tencent.com/document/api/664/131655)

新增数据结构：

* [CosAccessInfo](https://cloud.tencent.com/document/api/664/90825#CosAccessInfo)
* [CosActionInfo](https://cloud.tencent.com/document/api/664/90825#CosActionInfo)
* [CosAkAssetInfo](https://cloud.tencent.com/document/api/664/90825#CosAkAssetInfo)
* [CosAkSet](https://cloud.tencent.com/document/api/664/90825#CosAkSet)
* [CosAlarmInfo](https://cloud.tencent.com/document/api/664/90825#CosAlarmInfo)
* [CosAlarmRiskIdInfo](https://cloud.tencent.com/document/api/664/90825#CosAlarmRiskIdInfo)
* [CosAlarmTrendInfo](https://cloud.tencent.com/document/api/664/90825#CosAlarmTrendInfo)
* [CosAssetDataScanDetail](https://cloud.tencent.com/document/api/664/90825#CosAssetDataScanDetail)
* [CosAssetFileIdentifyInfo](https://cloud.tencent.com/document/api/664/90825#CosAssetFileIdentifyInfo)
* [CosAssetInfo](https://cloud.tencent.com/document/api/664/90825#CosAssetInfo)
* [CosAssetSyncTaskInfo](https://cloud.tencent.com/document/api/664/90825#CosAssetSyncTaskInfo)
* [CosAuditPayInfo](https://cloud.tencent.com/document/api/664/90825#CosAuditPayInfo)
* [CosBucketAccessWay](https://cloud.tencent.com/document/api/664/90825#CosBucketAccessWay)
* [CosBucketBillingInfo](https://cloud.tencent.com/document/api/664/90825#CosBucketBillingInfo)
* [CosBucketId](https://cloud.tencent.com/document/api/664/90825#CosBucketId)
* [CosBucketInfo](https://cloud.tencent.com/document/api/664/90825#CosBucketInfo)
* [CosBucketTaskInfo](https://cloud.tencent.com/document/api/664/90825#CosBucketTaskInfo)
* [CosDictionary](https://cloud.tencent.com/document/api/664/90825#CosDictionary)
* [CosIdentifyCategoryDetail](https://cloud.tencent.com/document/api/664/90825#CosIdentifyCategoryDetail)
* [CosIdentifyRuleDetail](https://cloud.tencent.com/document/api/664/90825#CosIdentifyRuleDetail)
* [CosInvokeDetailInfo](https://cloud.tencent.com/document/api/664/90825#CosInvokeDetailInfo)
* [CosInvokeIpVpcInfo](https://cloud.tencent.com/document/api/664/90825#CosInvokeIpVpcInfo)
* [CosInvokeLog](https://cloud.tencent.com/document/api/664/90825#CosInvokeLog)
* [CosOverview](https://cloud.tencent.com/document/api/664/90825#CosOverview)
* [CosPermissionInfo](https://cloud.tencent.com/document/api/664/90825#CosPermissionInfo)
* [CosPolicyInfo](https://cloud.tencent.com/document/api/664/90825#CosPolicyInfo)
* [CosRiskActionInfo](https://cloud.tencent.com/document/api/664/90825#CosRiskActionInfo)
* [CosRiskAlarmInfo](https://cloud.tencent.com/document/api/664/90825#CosRiskAlarmInfo)
* [CosRiskBucketInfo](https://cloud.tencent.com/document/api/664/90825#CosRiskBucketInfo)
* [CosRiskInfo](https://cloud.tencent.com/document/api/664/90825#CosRiskInfo)
* [CosRiskTrendInfo](https://cloud.tencent.com/document/api/664/90825#CosRiskTrendInfo)
* [CosRiskViewInfo](https://cloud.tencent.com/document/api/664/90825#CosRiskViewInfo)
* [CosRoleAccessInfo](https://cloud.tencent.com/document/api/664/90825#CosRoleAccessInfo)
* [CosSourceIpInfo](https://cloud.tencent.com/document/api/664/90825#CosSourceIpInfo)



## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 294 次发布

发布时间：2026-05-18 01:53:25

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateSeal](https://cloud.tencent.com/document/api/1323/94136)

	* 新增出参：ImageUrl




## 腾讯电子签（基础版）(essbasic) 版本：2021-05-26

### 第 258 次发布

发布时间：2026-05-18 01:55:11

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateSealByImage](https://cloud.tencent.com/document/api/1420/73067)

	* 新增出参：PreviewFileUrl, PreviewPdfUrl




## 腾讯电子签（基础版）(essbasic) 版本：2020-12-22



## 物联网开发平台(iotexplorer) 版本：2019-04-23

### 第 138 次发布

发布时间：2026-05-18 02:03:12

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeLicenseOverview](https://cloud.tencent.com/document/api/1081/131699)



## TI-ONE 训练平台(tione) 版本：2021-11-11

### 第 117 次发布

发布时间：2026-05-18 03:23:37

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeBillingSpecs](https://cloud.tencent.com/document/api/851/112648)

	* 新增入参：TiProjectId


修改数据结构：

* [Service](https://cloud.tencent.com/document/api/851/75051#Service)

	* 新增成员：Changer, ChangerName

* [ServiceGroup](https://cloud.tencent.com/document/api/851/75051#ServiceGroup)

	* 新增成员：Changer, ChangerName




## TI-ONE 训练平台(tione) 版本：2019-10-22



## 云点播(vod) 版本：2024-07-18



## 云点播(vod) 版本：2018-07-17

### 第 257 次发布

发布时间：2026-05-18 03:45:46

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeProcedureTemplates](https://cloud.tencent.com/document/api/266/33895)

	* 新增入参：SortBy




## Web 应用防火墙(waf) 版本：2018-01-25

### 第 152 次发布

发布时间：2026-05-18 03:58:46

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [ModifyOwaspDomainUpdateStatus](https://cloud.tencent.com/document/api/627/131700)



