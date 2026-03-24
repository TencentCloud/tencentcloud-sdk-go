# Release v1.3.64

## 应用性能监控(apm) 版本：2021-06-22

### 第 58 次发布

发布时间：2026-03-25 01:11:00

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyApmApplicationConfig](https://cloud.tencent.com/document/api/1463/125072)

	* 新增入参：UseDefaultFuseConfig


修改数据结构：

* [ApmAppConfig](https://cloud.tencent.com/document/api/1463/64927#ApmAppConfig)

	* 新增成员：UseDefaultFuseConfig

* [ApmApplicationConfigView](https://cloud.tencent.com/document/api/1463/64927#ApmApplicationConfigView)

	* 新增成员：UseDefaultFuseConfig




## 费用中心(billing) 版本：2018-07-09

### 第 86 次发布

发布时间：2026-03-25 01:14:23

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeVoucherInfo](https://cloud.tencent.com/document/api/555/70813)

	* 新增入参：Lang




## 配置审计(config) 版本：2022-08-02

### 第 6 次发布

发布时间：2026-03-25 01:30:44

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [AddAggregateCompliancePack](https://cloud.tencent.com/document/api/1579/129751)
* [AddAggregateConfigRule](https://cloud.tencent.com/document/api/1579/129727)
* [AddCompliancePack](https://cloud.tencent.com/document/api/1579/129750)
* [AddConfigRule](https://cloud.tencent.com/document/api/1579/129726)
* [CloseAggregateConfigRule](https://cloud.tencent.com/document/api/1579/129725)
* [CloseConfigRecorder](https://cloud.tencent.com/document/api/1579/129701)
* [CloseConfigRule](https://cloud.tencent.com/document/api/1579/129724)
* [CreateAggregator](https://cloud.tencent.com/document/api/1579/129705)
* [CreateRemediation](https://cloud.tencent.com/document/api/1579/129723)
* [DeleteAggregateCompliancePack](https://cloud.tencent.com/document/api/1579/129749)
* [DeleteAggregateConfigRule](https://cloud.tencent.com/document/api/1579/129722)
* [DeleteCompliancePack](https://cloud.tencent.com/document/api/1579/129748)
* [DeleteConfigRule](https://cloud.tencent.com/document/api/1579/129721)
* [DeleteRemediations](https://cloud.tencent.com/document/api/1579/129720)
* [DescribeAggregateCompliancePack](https://cloud.tencent.com/document/api/1579/129747)
* [DescribeAggregateConfigDeliver](https://cloud.tencent.com/document/api/1579/129734)
* [DescribeAggregateConfigRule](https://cloud.tencent.com/document/api/1579/129719)
* [DescribeAggregator](https://cloud.tencent.com/document/api/1579/129704)
* [DescribeCompliancePack](https://cloud.tencent.com/document/api/1579/129746)
* [DescribeConfigDeliver](https://cloud.tencent.com/document/api/1579/129733)
* [DescribeConfigRecorder](https://cloud.tencent.com/document/api/1579/129700)
* [DescribeConfigRule](https://cloud.tencent.com/document/api/1579/129718)
* [DescribeSystemCompliancePack](https://cloud.tencent.com/document/api/1579/129745)
* [DescribeSystemRule](https://cloud.tencent.com/document/api/1579/129729)
* [DetachAggregateConfigRuleToCompliancePack](https://cloud.tencent.com/document/api/1579/129744)
* [DetachConfigRuleToCompliancePack](https://cloud.tencent.com/document/api/1579/129743)
* [ListAggregateCompliancePacks](https://cloud.tencent.com/document/api/1579/129742)
* [ListAggregateConfigRuleEvaluationResults](https://cloud.tencent.com/document/api/1579/129717)
* [ListAggregators](https://cloud.tencent.com/document/api/1579/129703)
* [ListCompliancePacks](https://cloud.tencent.com/document/api/1579/129741)
* [ListConfigRuleEvaluationResults](https://cloud.tencent.com/document/api/1579/129716)
* [ListRemediationExecutions](https://cloud.tencent.com/document/api/1579/129699)
* [ListRemediations](https://cloud.tencent.com/document/api/1579/129715)
* [ListResourceTypes](https://cloud.tencent.com/document/api/1579/129698)
* [ListSystemCompliancePacks](https://cloud.tencent.com/document/api/1579/129740)
* [ListSystemRules](https://cloud.tencent.com/document/api/1579/129714)
* [OpenAggregateConfigRule](https://cloud.tencent.com/document/api/1579/129713)
* [OpenConfigRecorder](https://cloud.tencent.com/document/api/1579/129753)
* [OpenConfigRule](https://cloud.tencent.com/document/api/1579/129712)
* [StartAggregateConfigRuleEvaluation](https://cloud.tencent.com/document/api/1579/129711)
* [StartConfigRuleEvaluation](https://cloud.tencent.com/document/api/1579/129710)
* [StartRemediation](https://cloud.tencent.com/document/api/1579/129709)
* [UpdateAggregateCompliancePack](https://cloud.tencent.com/document/api/1579/129739)
* [UpdateAggregateCompliancePackStatus](https://cloud.tencent.com/document/api/1579/129738)
* [UpdateAggregateConfigDeliver](https://cloud.tencent.com/document/api/1579/129732)
* [UpdateAggregateConfigRule](https://cloud.tencent.com/document/api/1579/129708)
* [UpdateCompliancePack](https://cloud.tencent.com/document/api/1579/129737)
* [UpdateCompliancePackStatus](https://cloud.tencent.com/document/api/1579/129736)
* [UpdateConfigDeliver](https://cloud.tencent.com/document/api/1579/129731)
* [UpdateConfigRecorder](https://cloud.tencent.com/document/api/1579/129697)
* [UpdateConfigRule](https://cloud.tencent.com/document/api/1579/129707)
* [UpdateRemediation](https://cloud.tencent.com/document/api/1579/129706)

新增数据结构：

* [AggregateEvaluationResult](https://cloud.tencent.com/document/api/1579/101783#AggregateEvaluationResult)
* [Aggregator](https://cloud.tencent.com/document/api/1579/101783#Aggregator)
* [AggregatorAccount](https://cloud.tencent.com/document/api/1579/101783#AggregatorAccount)
* [ComplianceConfigRule](https://cloud.tencent.com/document/api/1579/101783#ComplianceConfigRule)
* [CompliancePackRule](https://cloud.tencent.com/document/api/1579/101783#CompliancePackRule)
* [CompliancePackRuleForManage](https://cloud.tencent.com/document/api/1579/101783#CompliancePackRuleForManage)
* [CompliancePackRules](https://cloud.tencent.com/document/api/1579/101783#CompliancePackRules)
* [ConfigCompliancePack](https://cloud.tencent.com/document/api/1579/101783#ConfigCompliancePack)
* [ConfigResource](https://cloud.tencent.com/document/api/1579/101783#ConfigResource)
* [Control](https://cloud.tencent.com/document/api/1579/101783#Control)
* [EvaluationResult](https://cloud.tencent.com/document/api/1579/101783#EvaluationResult)
* [Remediation](https://cloud.tencent.com/document/api/1579/101783#Remediation)
* [RemediationExecutions](https://cloud.tencent.com/document/api/1579/101783#RemediationExecutions)
* [SystemCompliancePack](https://cloud.tencent.com/document/api/1579/101783#SystemCompliancePack)
* [SystemConfigRule](https://cloud.tencent.com/document/api/1579/101783#SystemConfigRule)
* [UserConfigResource](https://cloud.tencent.com/document/api/1579/101783#UserConfigResource)



## 云安全一体化平台(csip) 版本：2022-11-21

### 第 70 次发布

发布时间：2026-03-25 01:31:38

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [SkillState](https://cloud.tencent.com/document/api/664/90825#SkillState)

修改数据结构：

* [AIAgentAsset](https://cloud.tencent.com/document/api/664/90825#AIAgentAsset)

	* 新增成员：SkillState




## 数据湖计算 DLC(dlc) 版本：2021-01-25

### 第 153 次发布

发布时间：2026-03-25 01:44:07

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [DatabaseResponseInfo](https://cloud.tencent.com/document/api/1342/53778#DatabaseResponseInfo)

	* 新增成员：CatalogName, CatalogType, IsInformationSchema




## DNSPod(dnspod) 版本：2021-03-23

### 第 51 次发布

发布时间：2026-03-25 01:46:09

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateAndPayDeal](https://cloud.tencent.com/document/api/1427/129754)



## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 283 次发布

发布时间：2026-03-25 01:52:30

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeContractReviewMarkedRiskExportTask](https://cloud.tencent.com/document/api/1323/129756)
* [ExportContractReviewMarkedRisk](https://cloud.tencent.com/document/api/1323/129755)



## 物联网开发平台(iotexplorer) 版本：2019-04-23

### 第 135 次发布

发布时间：2026-03-24 17:53:25

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeProductDynamicRegister](https://cloud.tencent.com/document/api/1081/129696)
* [ModifyProductDynamicRegister](https://cloud.tencent.com/document/api/1081/129695)



## 云直播CSS(live) 版本：2018-08-01

### 第 167 次发布

发布时间：2026-03-25 02:10:49

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateVideoRedrawTask](https://cloud.tencent.com/document/api/267/129758)

新增数据结构：

* [VideoRedrawCosInfo](https://cloud.tencent.com/document/api/267/20474#VideoRedrawCosInfo)
* [VideoRedrawInput](https://cloud.tencent.com/document/api/267/20474#VideoRedrawInput)



## 云开发 CloudBase(tcb) 版本：2018-06-08

### 第 130 次发布

发布时间：2026-03-25 02:56:08

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DeleteVmInstance](https://cloud.tencent.com/document/api/876/129761)
* [DescribeVmInstances](https://cloud.tencent.com/document/api/876/129760)
* [InquireVmPrice](https://cloud.tencent.com/document/api/876/129759)

新增数据结构：

* [VmInstance](https://cloud.tencent.com/document/api/876/34822#VmInstance)



## 容器镜像服务(tcr) 版本：2019-09-24

### 第 76 次发布

发布时间：2026-03-25 02:58:40

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyNamespace](https://cloud.tencent.com/document/api/1141/42727)

	* 新增入参：TagSpecification




## 数据开发治理平台 WeData(wedata) 版本：2025-08-06



## 数据开发治理平台 WeData(wedata) 版本：2021-08-20

### 第 184 次发布

发布时间：2026-03-25 03:54:25

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [ProdSchedulerTask](https://cloud.tencent.com/document/api/1267/76336#ProdSchedulerTask)

	* 新增成员：TaskStatus

* [RuleGroupPage](https://cloud.tencent.com/document/api/1267/76336#RuleGroupPage)

	* 新增成员：MonitorEnabledCount




