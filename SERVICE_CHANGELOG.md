# Release v1.3.154

## 费用中心(billing) 版本：2018-07-09

### 第 93 次发布

发布时间：2026-08-07 01:12:28

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [GatherResourceSummary](https://cloud.tencent.com/document/api/555/19183#GatherResourceSummary)

	* 新增成员：EffectiveMode




## 云防火墙(cfw) 版本：2019-09-04

### 第 112 次发布

发布时间：2026-08-07 01:16:57

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateWhiteRule](https://cloud.tencent.com/document/api/1132/135701)
* [DeleteWhiteRule](https://cloud.tencent.com/document/api/1132/135699)
* [ModifyWhiteRule](https://cloud.tencent.com/document/api/1132/135700)

新增数据结构：

* [IdsWhiteRule](https://cloud.tencent.com/document/api/1132/49071#IdsWhiteRule)
* [WhiteRuleInfo](https://cloud.tencent.com/document/api/1132/49071#WhiteRuleInfo)



## 云原生智能网关(cngw) 版本：2023-04-18

### 第 5 次发布

发布时间：2026-08-07 01:21:12

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeCNGWServicesWithRoutes](https://cloud.tencent.com/document/api/1826/135703)

新增数据结构：

* [KVMapping](https://cloud.tencent.com/document/api/1826/133161#KVMapping)
* [KongRoutePreview](https://cloud.tencent.com/document/api/1826/133161#KongRoutePreview)
* [KongServicePreview](https://cloud.tencent.com/document/api/1826/133161#KongServicePreview)
* [KongServiceRoute](https://cloud.tencent.com/document/api/1826/133161#KongServiceRoute)
* [KongServiceWithRoutes](https://cloud.tencent.com/document/api/1826/133161#KongServiceWithRoutes)
* [KongTarget](https://cloud.tencent.com/document/api/1826/133161#KongTarget)
* [KongUpstreamInfo](https://cloud.tencent.com/document/api/1826/133161#KongUpstreamInfo)
* [ListFilter](https://cloud.tencent.com/document/api/1826/133161#ListFilter)



## 配置审计(config) 版本：2022-08-02

### 第 12 次发布

发布时间：2026-08-07 01:21:31

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [ListAggregateConfigRuleResourceEvaluationResults](https://cloud.tencent.com/document/api/1579/135705)
* [ListConfigRuleResourceEvaluationResults](https://cloud.tencent.com/document/api/1579/135704)

新增数据结构：

* [ConfigRuleResourceEvaluationResult](https://cloud.tencent.com/document/api/1579/101783#ConfigRuleResourceEvaluationResult)



## 时序数据库 CTSDB(ctsdb) 版本：2023-02-02

### 第 4 次发布

发布时间：2026-08-07 01:23:24

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [Database](https://cloud.tencent.com/document/api/652/121580#Database)

	* 新增成员：CoolDownTime




## 云服务器(cvm) 版本：2017-03-12

### 第 168 次发布

发布时间：2026-08-07 01:23:27

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [InstanceTypeConfig](https://cloud.tencent.com/document/api/213/15753#InstanceTypeConfig)

	* 新增成员：GpuType, GpuMemory




## 云数据库独享集群(dbdc) 版本：2020-10-29

### 第 13 次发布

发布时间：2026-08-07 01:27:45

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [ModifyDBCustomClusterAttributes](https://cloud.tencent.com/document/api/1322/135706)



## Elasticsearch Service(es) 版本：2025-01-01



## Elasticsearch Service(es) 版本：2018-04-16

### 第 107 次发布

发布时间：2026-08-07 01:33:10

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [KibanaView](https://cloud.tencent.com/document/api/845/30634#KibanaView)

	* 新增成员：UserIp




## 文字识别(ocr) 版本：2018-11-19

### 第 258 次发布

发布时间：2026-08-06 11:53:45

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除接口**：</font>

* EduPaperOCR
* ExtractDocMultiPro
* FlightInvoiceOCR
* FormulaOCR
* InsuranceBillOCR

修改接口：

* [ExtractDocMulti](https://cloud.tencent.com/document/api/866/119451)

	* 新增入参：NewItemNames, MultiModelVersion

* [MultimodalDocParse](https://cloud.tencent.com/document/api/866/133224)

	* 新增入参：TaskType

* [SubmitQuestionMarkAgentJob](https://cloud.tencent.com/document/api/866/128273)

	* 新增入参：AssistMarkType, AnswerAssistMap

	* 新增出参：OriginalImageUrl


<font color="#dd0000">**删除数据结构**：</font>

* FlightInvoiceInfo
* InsuranceBillInfo
* QuestionBlockObj
* QuestionObj
* TextEduPaper
* TextFormula



## 云数据库 PostgreSQL(postgres) 版本：2017-03-12

### 第 70 次发布

发布时间：2026-08-07 01:49:38

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateReadOnlyDBInstance](https://cloud.tencent.com/document/api/409/52602)

	* 新增入参：Tags

* [ModifyDBInstanceSpec](https://cloud.tencent.com/document/api/409/63689)

	* 新增入参：SyncModifyParams




## 风险识别 RCE(rce) 版本：2026-01-30

### 第 3 次发布

发布时间：2026-08-07 01:50:44

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [DataScore](https://cloud.tencent.com/document/api/1343/134560#DataScore)

	* 新增成员：RiskScore

* [Device](https://cloud.tencent.com/document/api/1343/134560#Device)

	* 新增成员：SignToken, TokenTime




## 风险识别 RCE(rce) 版本：2025-04-25



## 风险识别 RCE(rce) 版本：2020-11-03



## 云数据库Redis(redis) 版本：2018-04-12

### 第 107 次发布

发布时间：2026-08-07 01:50:47

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateInstances](https://cloud.tencent.com/document/api/239/20026)

	* 新增入参：PasswordPolicy, EnableSSL, SSLBindPrivateIPv4, ConnectionMode




## 云开发 CloudBase(tcb) 版本：2018-06-08

### 第 156 次发布

发布时间：2026-08-07 01:54:46

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeCloudBaseRunBuildLog](https://cloud.tencent.com/document/api/876/135707)

新增数据结构：

* [CloudBaseRunBuildLog](https://cloud.tencent.com/document/api/876/34822#CloudBaseRunBuildLog)



## TSF-Polaris&ZK&网关(tse) 版本：2020-12-07

### 第 112 次发布

发布时间：2026-08-07 02:03:38

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeCNGWServicesWithRoutes](https://cloud.tencent.com/document/api/1364/135708)

修改接口：

* [DescribeCloudNativeAPIGatewayRoutes](https://cloud.tencent.com/document/api/1364/94842)

	* 新增入参：RouteTypes, GrayRoutesFirst, OrderField, OrderType

* [DescribeCloudNativeAPIGatewayServices](https://cloud.tencent.com/document/api/1364/94840)

	* 新增入参：OrderField, OrderType


新增数据结构：

* [KongServiceRoute](https://cloud.tencent.com/document/api/1364/54942#KongServiceRoute)
* [KongServiceWithRoutes](https://cloud.tencent.com/document/api/1364/54942#KongServiceWithRoutes)

修改数据结构：

* [KongActiveHealthCheck](https://cloud.tencent.com/document/api/1364/54942#KongActiveHealthCheck)

	* 新增成员：HostHeader

* [KongRoutePreview](https://cloud.tencent.com/document/api/1364/54942#KongRoutePreview)

	* 新增成员：RouteSource




## TSF-应用管理&Consul(tsf) 版本：2018-03-26

### 第 145 次发布

发布时间：2026-08-07 02:05:41

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DeployContainerApplication](https://cloud.tencent.com/document/api/649/120669)

	* 新增入参：MeshSidecarVersion




## 云点播(vod) 版本：2024-07-18



## 云点播(vod) 版本：2018-07-17

### 第 280 次发布

发布时间：2026-08-07 02:10:19

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [AigcVideoTaskUsage](https://cloud.tencent.com/document/api/266/31773#AigcVideoTaskUsage)

	* 新增成员：InputImageCount, InputSeconds, OutputSeconds, TotalSeconds




