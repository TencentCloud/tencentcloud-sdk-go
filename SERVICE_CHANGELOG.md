# Release v1.3.156

## AI Agent 安全网关(apis) 版本：2024-08-01

### 第 12 次发布

发布时间：2026-08-11 01:12:33

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateModelService](https://cloud.tencent.com/document/api/1627/129615)

	* 新增入参：RouteStrategy, TokenLengthRoute, TaskComplexityRoute

* [ModifyModelService](https://cloud.tencent.com/document/api/1627/129611)

	* 新增入参：RouteStrategy, TokenLengthRoute, TaskComplexityRoute


新增数据结构：

* [AgentCredentialApiKeyDTO](https://cloud.tencent.com/document/api/1627/129635#AgentCredentialApiKeyDTO)
* [FaultToleranceDTO](https://cloud.tencent.com/document/api/1627/129635#FaultToleranceDTO)
* [TaskComplexityRouteDTO](https://cloud.tencent.com/document/api/1627/129635#TaskComplexityRouteDTO)
* [TokenLengthRouteDTO](https://cloud.tencent.com/document/api/1627/129635#TokenLengthRouteDTO)

修改数据结构：

* [AgentCredentialContentDTO](https://cloud.tencent.com/document/api/1627/129635#AgentCredentialContentDTO)

	* 新增成员：ApiKeys, FaultTolerance

* [DescribeModelServiceResponseVO](https://cloud.tencent.com/document/api/1627/129635#DescribeModelServiceResponseVO)

	* 新增成员：RouteStrategy, TokenLengthRoute, TaskComplexityRoute




## 云联络中心(ccc) 版本：2020-02-10

### 第 133 次发布

发布时间：2026-08-11 01:20:28

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeTelCdr](https://cloud.tencent.com/document/api/679/47714)


修改数据结构：

* [TelCdrInfo](https://cloud.tencent.com/document/api/679/47715#TelCdrInfo)

	* 新增成员：PostIVRKeyPressedEx




## 云数据库 MySQL(cdb) 版本：2017-03-20

### 第 226 次发布

发布时间：2026-08-11 01:21:27

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [ModifyInstanceChargeType](https://cloud.tencent.com/document/api/236/135816)
* [ModifyInstanceDestroyProtect](https://cloud.tencent.com/document/api/236/135817)

修改接口：

* [UpgradeDBInstance](https://cloud.tencent.com/document/api/236/15876)

	* 新增出参：JobId


新增数据结构：

* [InstanceChargePrepaid](https://cloud.tencent.com/document/api/236/15878#InstanceChargePrepaid)



## T-Sec-数据安全审计（DSA）(cds) 版本：2018-04-20

### 第 5 次发布

发布时间：2026-08-11 01:24:14

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateReportPdf](https://cloud.tencent.com/document/api/856/135822)
* [CreateTimerReport](https://cloud.tencent.com/document/api/856/135821)
* [DescribeAssetsList](https://cloud.tencent.com/document/api/856/135820)
* [DescribeReportList](https://cloud.tencent.com/document/api/856/135819)
* [DescribeReportMissionList](https://cloud.tencent.com/document/api/856/135818)

新增数据结构：

* [AssetsInfo](https://cloud.tencent.com/document/api/856/33913#AssetsInfo)
* [AuditCapability](https://cloud.tencent.com/document/api/856/33913#AuditCapability)
* [DsgcBindingInfo](https://cloud.tencent.com/document/api/856/33913#DsgcBindingInfo)
* [IdWithName](https://cloud.tencent.com/document/api/856/33913#IdWithName)
* [NameValueString](https://cloud.tencent.com/document/api/856/33913#NameValueString)
* [ReportMission](https://cloud.tencent.com/document/api/856/33913#ReportMission)
* [Reports](https://cloud.tencent.com/document/api/856/33913#Reports)



## 负载均衡(clb) 版本：2018-03-17

### 第 157 次发布

发布时间：2026-08-11 01:30:25

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateModel](https://cloud.tencent.com/document/api/214/133679)

	* 新增入参：HealthCheckConfig

* [DescribeModelNames](https://cloud.tencent.com/document/api/214/133674)

	* 新增入参：Filters


新增数据结构：

* [ServiceProviderHealthCheckConfigInput](https://cloud.tencent.com/document/api/214/30694#ServiceProviderHealthCheckConfigInput)
* [ServiceProviderHealthCheckConfigOutput](https://cloud.tencent.com/document/api/214/30694#ServiceProviderHealthCheckConfigOutput)

修改数据结构：

* [ModelKeyInfoItem](https://cloud.tencent.com/document/api/214/30694#ModelKeyInfoItem)

	* 新增成员：HealthCheckConfig




## 日志服务(cls) 版本：2020-10-16

### 第 172 次发布

发布时间：2026-08-11 01:32:38

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateAgentApplication](https://cloud.tencent.com/document/api/614/134210)

	* <font color="#dd0000">**修改入参**：</font>LogsetId




## 云安全一体化平台(csip) 版本：2022-11-21

### 第 94 次发布

发布时间：2026-08-11 01:36:35

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [AddVulWhitelist](https://cloud.tencent.com/document/api/664/135857)
* [CreateHostVulExportJob](https://cloud.tencent.com/document/api/664/135856)
* [CreateVulFixRetryTask](https://cloud.tencent.com/document/api/664/135833)
* [CreateVulFixTask](https://cloud.tencent.com/document/api/664/135832)
* [CreateVulFixedExportJob](https://cloud.tencent.com/document/api/664/135831)
* [CreateVulReScan](https://cloud.tencent.com/document/api/664/135855)
* [CreateVulScanManual](https://cloud.tencent.com/document/api/664/135854)
* [DeleteVulWhitelist](https://cloud.tencent.com/document/api/664/135853)
* [DescribeHostKBRiskList](https://cloud.tencent.com/document/api/664/135852)
* [DescribeHostVulItemVPRInfo](https://cloud.tencent.com/document/api/664/135851)
* [DescribeHostVulOverview](https://cloud.tencent.com/document/api/664/135850)
* [DescribeHostVulRiskList](https://cloud.tencent.com/document/api/664/135849)
* [DescribeKBDetail](https://cloud.tencent.com/document/api/664/135830)
* [DescribeKBUpdatableMachineList](https://cloud.tencent.com/document/api/664/135829)
* [DescribeVulComponentRelateHost](https://cloud.tencent.com/document/api/664/135848)
* [DescribeVulFixTaskDetail](https://cloud.tencent.com/document/api/664/135828)
* [DescribeVulFixTaskList](https://cloud.tencent.com/document/api/664/135827)
* [DescribeVulFixableMachineList](https://cloud.tencent.com/document/api/664/135826)
* [DescribeVulFixedHostDetail](https://cloud.tencent.com/document/api/664/135825)
* [DescribeVulFixedList](https://cloud.tencent.com/document/api/664/135824)
* [DescribeVulHostRelateComponent](https://cloud.tencent.com/document/api/664/135847)
* [DescribeVulIgnoreRuleList](https://cloud.tencent.com/document/api/664/135846)
* [DescribeVulItemList](https://cloud.tencent.com/document/api/664/135845)
* [DescribeVulLabelList](https://cloud.tencent.com/document/api/664/135844)
* [DescribeVulRiskRelateComponent](https://cloud.tencent.com/document/api/664/135843)
* [DescribeVulRiskRelateHost](https://cloud.tencent.com/document/api/664/135842)
* [DescribeVulScanPeriodic](https://cloud.tencent.com/document/api/664/135841)
* [DescribeVulScanTaskDetail](https://cloud.tencent.com/document/api/664/135840)
* [DescribeVulScanTaskList](https://cloud.tencent.com/document/api/664/135839)
* [ModifyVulScanPeriodic](https://cloud.tencent.com/document/api/664/135838)
* [ModifyVulWhitelistConfig](https://cloud.tencent.com/document/api/664/135837)
* [ModifyVulWhitelistSwitch](https://cloud.tencent.com/document/api/664/135836)
* [StopVulScanTask](https://cloud.tencent.com/document/api/664/135835)

新增数据结构：

* [AccountBriefInfo](https://cloud.tencent.com/document/api/664/90825#AccountBriefInfo)
* [ComponentDetailItem](https://cloud.tencent.com/document/api/664/90825#ComponentDetailItem)
* [HostBriefInfo](https://cloud.tencent.com/document/api/664/90825#HostBriefInfo)
* [HostKBRisk](https://cloud.tencent.com/document/api/664/90825#HostKBRisk)
* [HostVulComponent](https://cloud.tencent.com/document/api/664/90825#HostVulComponent)
* [HostVulOverview](https://cloud.tencent.com/document/api/664/90825#HostVulOverview)
* [HostVulRisk](https://cloud.tencent.com/document/api/664/90825#HostVulRisk)
* [KBDetail](https://cloud.tencent.com/document/api/664/90825#KBDetail)
* [KBFixSummaryItem](https://cloud.tencent.com/document/api/664/90825#KBFixSummaryItem)
* [KBUpdateMachineItem](https://cloud.tencent.com/document/api/664/90825#KBUpdateMachineItem)
* [VPRLabel](https://cloud.tencent.com/document/api/664/90825#VPRLabel)
* [VPRRatingInfo](https://cloud.tencent.com/document/api/664/90825#VPRRatingInfo)
* [VPRRatingStage](https://cloud.tencent.com/document/api/664/90825#VPRRatingStage)
* [VulBriefInfo](https://cloud.tencent.com/document/api/664/90825#VulBriefInfo)
* [VulComponentSummary](https://cloud.tencent.com/document/api/664/90825#VulComponentSummary)
* [VulDetailInfo](https://cloud.tencent.com/document/api/664/90825#VulDetailInfo)
* [VulFixItem](https://cloud.tencent.com/document/api/664/90825#VulFixItem)
* [VulFixStatusItem](https://cloud.tencent.com/document/api/664/90825#VulFixStatusItem)
* [VulFixSummaryItem](https://cloud.tencent.com/document/api/664/90825#VulFixSummaryItem)
* [VulFixTaskDetailItem](https://cloud.tencent.com/document/api/664/90825#VulFixTaskDetailItem)
* [VulFixTaskInfo](https://cloud.tencent.com/document/api/664/90825#VulFixTaskInfo)
* [VulFixTaskItem](https://cloud.tencent.com/document/api/664/90825#VulFixTaskItem)
* [VulFixableMachineItem](https://cloud.tencent.com/document/api/664/90825#VulFixableMachineItem)
* [VulFixedItem](https://cloud.tencent.com/document/api/664/90825#VulFixedItem)
* [VulHostBriefInfo](https://cloud.tencent.com/document/api/664/90825#VulHostBriefInfo)
* [VulScanTask](https://cloud.tencent.com/document/api/664/90825#VulScanTask)
* [VulScanTaskDetail](https://cloud.tencent.com/document/api/664/90825#VulScanTaskDetail)
* [VulSpreadTrend](https://cloud.tencent.com/document/api/664/90825#VulSpreadTrend)
* [VulVendorProduct](https://cloud.tencent.com/document/api/664/90825#VulVendorProduct)
* [VulWhitelist](https://cloud.tencent.com/document/api/664/90825#VulWhitelist)



## 数据湖计算 DLC(dlc) 版本：2021-01-25

### 第 172 次发布

发布时间：2026-08-11 01:51:49

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [TaskFullRespInfo](https://cloud.tencent.com/document/api/1342/53778#TaskFullRespInfo)

	* 新增成员：ResourceGroupType




## Elasticsearch Service(es) 版本：2025-01-01



## Elasticsearch Service(es) 版本：2018-04-16

### 第 108 次发布

发布时间：2026-08-11 02:02:21

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeUserCosSnapshotList](https://cloud.tencent.com/document/api/845/104321)

	* 新增入参：PaasEsRepository




## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 316 次发布

发布时间：2026-08-11 02:03:35

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreatePartnerAutoSignAuthUrl](https://cloud.tencent.com/document/api/1323/109065)

	* 新增入参：LimitAuthType

* [ModifyPartnerAutoSignAuthUrl](https://cloud.tencent.com/document/api/1323/120051)

	* 新增入参：LimitAuthType




## 腾讯电子签（基础版）(essbasic) 版本：2021-05-26

### 第 271 次发布

发布时间：2026-08-11 02:05:21

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreatePartnerAutoSignAuthUrl](https://cloud.tencent.com/document/api/1420/102515)

	* 新增入参：LimitAuthType

* [ModifyPartnerAutoSignAuthUrl](https://cloud.tencent.com/document/api/1420/120052)

	* 新增入参：LimitAuthType




## 腾讯电子签（基础版）(essbasic) 版本：2020-12-22



## 媒体处理(mps) 版本：2019-06-12

### 第 230 次发布

发布时间：2026-08-11 02:30:24

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [UnderstandImageConfig](https://cloud.tencent.com/document/api/862/37615#UnderstandImageConfig)

修改数据结构：

* [ImageTaskInput](https://cloud.tencent.com/document/api/862/37615#ImageTaskInput)

	* 新增成员：UnderstandImageConfig

* [VoiceInfo](https://cloud.tencent.com/document/api/862/37615#VoiceInfo)

	* 新增成员：Engine




## 云数据库 PostgreSQL(postgres) 版本：2017-03-12

### 第 71 次发布

发布时间：2026-08-11 02:38:17

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DeleteDatabase](https://cloud.tencent.com/document/api/409/135858)



## TI-ONE 训练平台(tione) 版本：2021-11-11

### 第 129 次发布

发布时间：2026-08-11 03:01:18

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [PrivateLinkInfo](https://cloud.tencent.com/document/api/851/75051#PrivateLinkInfo)

	* 新增成员：CreatedBy, CreateTime, SubUinName




## TI-ONE 训练平台(tione) 版本：2019-10-22



