# Release v1.3.59

## 应用性能监控(apm) 版本：2021-06-22

### 第 57 次发布

发布时间：2026-03-18 01:10:31

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [ServiceDetail](https://cloud.tencent.com/document/api/1463/64927#ServiceDetail)

	* 新增成员：EnableThresholdConfig, ErrRateThreshold, ResponseDurationWarningThreshold




## 运维安全中心（堡垒机）(bh) 版本：2023-04-18

### 第 27 次发布

发布时间：2026-03-18 01:12:29

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateAcl](https://cloud.tencent.com/document/api/1025/74411)

	* 新增入参：MaxAccessCredentialDuration

* [ModifyAcl](https://cloud.tencent.com/document/api/1025/74408)

	* 新增入参：MaxAccessCredentialDuration


修改数据结构：

* [Acl](https://cloud.tencent.com/document/api/1025/74416#Acl)

	* 新增成员：MaxAccessCredentialDuration

* [Resource](https://cloud.tencent.com/document/api/1025/74416#Resource)

	* 新增成员：ResourceEdition, TimeUnit, TimeSpan, PayMode




## 主机安全(cwp) 版本：2018-02-28

### 第 159 次发布

发布时间：2026-03-18 01:33:01

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeAttackType](https://cloud.tencent.com/document/api/296/129253)
* [DescribeInjectRiskyServiceSwitch](https://cloud.tencent.com/document/api/296/129241)
* [DescribeLoginTypeGlobalConf](https://cloud.tencent.com/document/api/296/129244)
* [DescribeLoginTypeHost](https://cloud.tencent.com/document/api/296/129243)
* [DescribeMemShellRules](https://cloud.tencent.com/document/api/296/129257)
* [DescribeRaspEventCWP](https://cloud.tencent.com/document/api/296/129252)
* [DescribeRaspEventDetailCWP](https://cloud.tencent.com/document/api/296/129251)
* [DescribeRaspEventDetailTCSS](https://cloud.tencent.com/document/api/296/129250)
* [DescribeRaspEventTCSS](https://cloud.tencent.com/document/api/296/129249)
* [DescribeRaspLicenseList](https://cloud.tencent.com/document/api/296/129238)
* [DescribeRaspMemShellDetailTCSS](https://cloud.tencent.com/document/api/296/129248)
* [DescribeRaspMemShellListTCSS](https://cloud.tencent.com/document/api/296/129247)
* [DescribeRaspPluginList](https://cloud.tencent.com/document/api/296/129237)
* [DescribeReverseShellRulesAggregation](https://cloud.tencent.com/document/api/296/129256)
* [DescribeReverseShellSystemPolicyConfig](https://cloud.tencent.com/document/api/296/129255)
* [DescribeShellPolicyList](https://cloud.tencent.com/document/api/296/129254)
* [DescribeVulDefenceOverviewCount](https://cloud.tencent.com/document/api/296/129240)
* [DescribeVulDefenceSettingList](https://cloud.tencent.com/document/api/296/129239)
* [DescribeYDRaspBlackWhite](https://cloud.tencent.com/document/api/296/129246)
* [RaspEventOverview](https://cloud.tencent.com/document/api/296/129245)

新增数据结构：

* [ClientSettingHost](https://cloud.tencent.com/document/api/296/19867#ClientSettingHost)
* [MemShellRule](https://cloud.tencent.com/document/api/296/19867#MemShellRule)
* [OrderDetail](https://cloud.tencent.com/document/api/296/19867#OrderDetail)
* [RaspAttackTypeListItem](https://cloud.tencent.com/document/api/296/19867#RaspAttackTypeListItem)
* [RaspEvent](https://cloud.tencent.com/document/api/296/19867#RaspEvent)
* [RaspEventDetail](https://cloud.tencent.com/document/api/296/19867#RaspEventDetail)
* [RaspEventOverview](https://cloud.tencent.com/document/api/296/19867#RaspEventOverview)
* [RaspLicenseList](https://cloud.tencent.com/document/api/296/19867#RaspLicenseList)
* [RaspLicensePlugin](https://cloud.tencent.com/document/api/296/19867#RaspLicensePlugin)
* [RaspMemShellDetail](https://cloud.tencent.com/document/api/296/19867#RaspMemShellDetail)
* [RaspMemShellEvent](https://cloud.tencent.com/document/api/296/19867#RaspMemShellEvent)
* [ReverseShellRuleAggregation](https://cloud.tencent.com/document/api/296/19867#ReverseShellRuleAggregation)
* [RiskMainClass](https://cloud.tencent.com/document/api/296/19867#RiskMainClass)
* [ShellPolicyList](https://cloud.tencent.com/document/api/296/19867#ShellPolicyList)
* [UuidHostip](https://cloud.tencent.com/document/api/296/19867#UuidHostip)
* [VulDefenceSetting](https://cloud.tencent.com/document/api/296/19867#VulDefenceSetting)
* [YDRaspBlackWhiteListItem](https://cloud.tencent.com/document/api/296/19867#YDRaspBlackWhiteListItem)



## TDSQL-C MySQL 版(cynosdb) 版本：2019-01-07

### 第 157 次发布

发布时间：2026-03-18 01:37:47

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateBackup](https://cloud.tencent.com/document/api/1003/80212)

	* 新增入参：Vaults

* [DescribeRollbackTimeRange](https://cloud.tencent.com/document/api/1003/48092)

	* 新增入参：VaultId, VaultRegion

* [RollBackCluster](https://cloud.tencent.com/document/api/1003/70115)

	* 新增入参：VaultId




## DNSPod(dnspod) 版本：2021-03-23

### 第 50 次发布

发布时间：2026-03-18 01:44:41

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeRecordList](https://cloud.tencent.com/document/api/1427/56166)

	* 新增入参：ErrorOnEmpty


修改数据结构：

* [Deals](https://cloud.tencent.com/document/api/1427/56185#Deals)

	* 新增成员：ResourceId




## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 279 次发布

发布时间：2026-03-18 01:51:46

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateBatchContractReviewTask](https://cloud.tencent.com/document/api/1323/122152)

	* 新增入参：Roles, ChecklistIds

* [CreateFlow](https://cloud.tencent.com/document/api/1323/70361)

	* 新增入参：FlowOperateLimit

* [CreateFlowByFiles](https://cloud.tencent.com/document/api/1323/70360)

	* 新增入参：FlowOperateLimit

* [DescribeContractReviewTask](https://cloud.tencent.com/document/api/1323/122151)

	* 新增出参：ChecklistIds, Roles


新增数据结构：

* [FlowOperateLimit](https://cloud.tencent.com/document/api/1323/70369#FlowOperateLimit)



## 腾讯电子签（基础版）(essbasic) 版本：2021-05-26

### 第 252 次发布

发布时间：2026-03-18 01:53:25

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ChannelCreateFlowByFiles](https://cloud.tencent.com/document/api/1420/73068)

	* 新增入参：FlowOperateLimit


新增数据结构：

* [FlowOperateLimit](https://cloud.tencent.com/document/api/1420/61525#FlowOperateLimit)

修改数据结构：

* [FlowInfo](https://cloud.tencent.com/document/api/1420/61525#FlowInfo)

	* 新增成员：FlowOperateLimit




## 腾讯电子签（基础版）(essbasic) 版本：2020-12-22



## 云数据库 SQL Server(sqlserver) 版本：2018-03-28

### 第 77 次发布

发布时间：2026-03-18 02:46:35

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DeleteDB](https://cloud.tencent.com/document/api/238/19971)

	* 新增入参：NoDoBackup




## 云开发 CloudBase(tcb) 版本：2018-06-08

### 第 128 次发布

发布时间：2026-03-18 02:54:47

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeCurveData](https://cloud.tencent.com/document/api/876/129258)



## 容器安全服务(tcss) 版本：2020-11-01

### 第 89 次发布

发布时间：2026-03-18 02:58:57

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeImageDenyEventDetail](https://cloud.tencent.com/document/api/1285/129264)
* [DescribeImageDenyEventList](https://cloud.tencent.com/document/api/1285/129263)
* [DescribeImageDenyEventTendency](https://cloud.tencent.com/document/api/1285/129262)
* [DescribeImageDenyRuleDetail](https://cloud.tencent.com/document/api/1285/129261)
* [DescribeImageDenyRuleList](https://cloud.tencent.com/document/api/1285/129260)
* [DescribeImageDenyRuleSummary](https://cloud.tencent.com/document/api/1285/129259)
* [DescribeMaliciousConnectionBlackList](https://cloud.tencent.com/document/api/1285/129266)
* [DescribeMaliciousConnectionWhiteList](https://cloud.tencent.com/document/api/1285/129265)
* [DescribeReverseShellRegexpWhiteList](https://cloud.tencent.com/document/api/1285/129268)
* [DescribeReverseShellRegexpWhiteListInfo](https://cloud.tencent.com/document/api/1285/129267)

新增数据结构：

* [ImageDenyEvent](https://cloud.tencent.com/document/api/1285/65614#ImageDenyEvent)
* [ImageDenyEventTendency](https://cloud.tencent.com/document/api/1285/65614#ImageDenyEventTendency)
* [ImageDenyRule](https://cloud.tencent.com/document/api/1285/65614#ImageDenyRule)
* [MaliciousConnectionRuleInfo](https://cloud.tencent.com/document/api/1285/65614#MaliciousConnectionRuleInfo)
* [RegexpRuleInfo](https://cloud.tencent.com/document/api/1285/65614#RegexpRuleInfo)
* [RegexpRuleListItem](https://cloud.tencent.com/document/api/1285/65614#RegexpRuleListItem)
* [WhiteListRegexpExpressionInfo](https://cloud.tencent.com/document/api/1285/65614#WhiteListRegexpExpressionInfo)



## 容器服务(tke) 版本：2022-05-01

### 第 24 次发布

发布时间：2026-03-18 03:22:48

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeGPUInfo](https://cloud.tencent.com/document/api/457/129269)
* [DescribeZoneInstanceConfigInfos](https://cloud.tencent.com/document/api/457/129270)



## 容器服务(tke) 版本：2018-05-25



## T-Sec-安心平台(RP)(trp) 版本：2021-05-15

### 第 38 次发布

发布时间：2026-03-18 03:25:26

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [ReportScanDetail](https://cloud.tencent.com/document/api/1458/129271)

新增数据结构：

* [ReportScanDetailResult](https://cloud.tencent.com/document/api/1458/75030#ReportScanDetailResult)
* [ScanDetailItem](https://cloud.tencent.com/document/api/1458/75030#ScanDetailItem)



## 数据开发治理平台 WeData(wedata) 版本：2025-08-06



## 数据开发治理平台 WeData(wedata) 版本：2021-08-20

### 第 181 次发布

发布时间：2026-03-18 03:53:25

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeDatabaseByName](https://cloud.tencent.com/document/api/1267/129278)
* [DescribeDatabaseMeta](https://cloud.tencent.com/document/api/1267/129277)
* [DescribeLineageColumns](https://cloud.tencent.com/document/api/1267/129276)
* [DescribeRealViewDatabasePage](https://cloud.tencent.com/document/api/1267/129272)
* [DescribeTableContentPreview](https://cloud.tencent.com/document/api/1267/129275)
* [DescribeTableDdl](https://cloud.tencent.com/document/api/1267/129274)
* [DescribeTableSelect](https://cloud.tencent.com/document/api/1267/129273)

新增数据结构：

* [BizParams](https://cloud.tencent.com/document/api/1267/76336#BizParams)
* [DatabaseRealViewVO](https://cloud.tencent.com/document/api/1267/76336#DatabaseRealViewVO)
* [DatabaseRealViewVOPageVO](https://cloud.tencent.com/document/api/1267/76336#DatabaseRealViewVOPageVO)
* [TableRecord](https://cloud.tencent.com/document/api/1267/76336#TableRecord)
* [TableRecordField](https://cloud.tencent.com/document/api/1267/76336#TableRecordField)



