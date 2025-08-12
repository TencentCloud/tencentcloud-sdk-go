# Release v1.1.4

## 内容分发网络 CDN(cdn) 版本：2018-06-06

### 第 136 次发布

发布时间：2025-08-01 11:13:02

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除接口**：</font>

* CreateScdnDomain
* CreateScdnFailedLogTask
* CreateScdnLogTask
* DeleteScdnDomain
* DescribeCcData
* DescribeDDoSData
* DescribeEventLogData
* DescribeScdnBotData
* DescribeScdnBotRecords
* DescribeScdnConfig
* DescribeScdnIpStrategy
* DescribeScdnTopData
* DescribeWafData
* ListScdnDomains
* ListScdnLogTasks
* ListScdnTopBotData
* ListTopBotData
* ListTopCcData
* ListTopDDoSData
* ListTopWafData
* StartScdnDomain
* StopScdnDomain
* UpdateScdnDomain

<font color="#dd0000">**删除数据结构**：</font>

* AdvancedCCRules
* AdvancedScdnAclGroup
* AdvancedScdnAclRule
* BotCookie
* BotJavaScript
* BotRecord
* BotSortBy
* BotStatisticsCount
* BotStats
* BotStatsDetailData
* CcTopData
* DDoSAttackBandwidthData
* DDoSAttackIPTopData
* DDoSStatsData
* DDoSTopData
* DomainBotCount
* EventLogStatsData
* ScdnAclConfig
* ScdnAclGroup
* ScdnAclRule
* ScdnBotConfig
* ScdnCCRules
* ScdnConfig
* ScdnDdosConfig
* ScdnDomain
* ScdnErrorPage
* ScdnEventLogConditions
* ScdnIpStrategy
* ScdnIpStrategyFilter
* ScdnLogTaskDetail
* ScdnSevenLayerRules
* ScdnTopData
* ScdnTopDomainData
* ScdnTopUrlData
* ScdnTypeData
* ScdnWafConfig
* ScdnWafRule
* WafSubRuleStatus



## 数据湖计算 DLC(dlc) 版本：2021-01-25

### 第 130 次发布

发布时间：2025-08-01 01:52:14

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [AssociateDatasourceHouse](https://cloud.tencent.com/document/api/1342/122147)
* [CreateStandardEngineResourceGroup](https://cloud.tencent.com/document/api/1342/122138)
* [CreateUserVpcConnection](https://cloud.tencent.com/document/api/1342/122143)
* [DeleteNativeSparkSession](https://cloud.tencent.com/document/api/1342/122149)
* [DeleteStandardEngineResourceGroup](https://cloud.tencent.com/document/api/1342/122137)
* [DeleteUserVpcConnection](https://cloud.tencent.com/document/api/1342/122142)
* [DescribeDataEngineSessionParameters](https://cloud.tencent.com/document/api/1342/122141)
* [DescribeEngineNetworks](https://cloud.tencent.com/document/api/1342/122140)
* [DescribeEngineNodeSpec](https://cloud.tencent.com/document/api/1342/122150)
* [DescribeNativeSparkSessions](https://cloud.tencent.com/document/api/1342/122148)
* [DescribeNetworkConnections](https://cloud.tencent.com/document/api/1342/122146)
* [DescribeSessionImageVersion](https://cloud.tencent.com/document/api/1342/122136)
* [DescribeStandardEngineResourceGroupConfigInfo](https://cloud.tencent.com/document/api/1342/122135)
* [DescribeStandardEngineResourceGroups](https://cloud.tencent.com/document/api/1342/122134)
* [DescribeUserVpcConnection](https://cloud.tencent.com/document/api/1342/122139)
* [LaunchStandardEngineResourceGroups](https://cloud.tencent.com/document/api/1342/122133)
* [PauseStandardEngineResourceGroups](https://cloud.tencent.com/document/api/1342/122132)
* [UnboundDatasourceHouse](https://cloud.tencent.com/document/api/1342/122145)
* [UpdateEngineResourceGroupNetworkConfigInfo](https://cloud.tencent.com/document/api/1342/122131)
* [UpdateNetworkConnection](https://cloud.tencent.com/document/api/1342/122144)
* [UpdateStandardEngineResourceGroupBaseInfo](https://cloud.tencent.com/document/api/1342/122130)
* [UpdateStandardEngineResourceGroupConfigInfo](https://cloud.tencent.com/document/api/1342/122129)
* [UpdateStandardEngineResourceGroupResourceInfo](https://cloud.tencent.com/document/api/1342/122128)

新增数据结构：

* [DataEngineImageSessionParameter](https://cloud.tencent.com/document/api/1342/53778#DataEngineImageSessionParameter)
* [EngineNetworkInfo](https://cloud.tencent.com/document/api/1342/53778#EngineNetworkInfo)
* [EngineResourceGroupConfigPair](https://cloud.tencent.com/document/api/1342/53778#EngineResourceGroupConfigPair)
* [EngineSessionImage](https://cloud.tencent.com/document/api/1342/53778#EngineSessionImage)
* [GatewayInfo](https://cloud.tencent.com/document/api/1342/53778#GatewayInfo)
* [OperateEngineResourceGroupFailMessage](https://cloud.tencent.com/document/api/1342/53778#OperateEngineResourceGroupFailMessage)
* [Param](https://cloud.tencent.com/document/api/1342/53778#Param)
* [SparkSessionInfo](https://cloud.tencent.com/document/api/1342/53778#SparkSessionInfo)
* [SpecInfo](https://cloud.tencent.com/document/api/1342/53778#SpecInfo)
* [StandardEngineResourceGroupConfigInfo](https://cloud.tencent.com/document/api/1342/53778#StandardEngineResourceGroupConfigInfo)
* [StandardEngineResourceGroupInfo](https://cloud.tencent.com/document/api/1342/53778#StandardEngineResourceGroupInfo)
* [UpdateConfContext](https://cloud.tencent.com/document/api/1342/53778#UpdateConfContext)
* [UserVpcConnectionInfo](https://cloud.tencent.com/document/api/1342/53778#UserVpcConnectionInfo)



## Elasticsearch Service(es) 版本：2025-01-01

### 第 7 次发布

发布时间：2025-08-01 02:09:11

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [GetDocumentParseResult](https://cloud.tencent.com/document/api/845/117806)

	* 新增出参：Usage

* [ParseDocument](https://cloud.tencent.com/document/api/845/117804)

	* 新增出参：Usage


新增数据结构：

* [PageUsage](https://cloud.tencent.com/document/api/845/117811#PageUsage)



## Elasticsearch Service(es) 版本：2018-04-16



## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 238 次发布

发布时间：2025-08-01 02:09:30

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateBatchContractReviewTask](https://cloud.tencent.com/document/api/1323/122152)
* [DescribeContractReviewTask](https://cloud.tencent.com/document/api/1323/122151)

新增数据结构：

* [OutputRisk](https://cloud.tencent.com/document/api/1323/70369#OutputRisk)
* [RiskIdentificationRoleInfo](https://cloud.tencent.com/document/api/1323/70369#RiskIdentificationRoleInfo)



## 腾讯混元大模型(hunyuan) 版本：2023-09-01

### 第 45 次发布

发布时间：2025-08-01 02:20:47

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除接口**：</font>

* QueryHunyuanTo3DJob
* SubmitHunyuanTo3DJob

<font color="#dd0000">**删除数据结构**：</font>

* File3D
* File3Ds



## 消息队列 MQTT 版(mqtt) 版本：2024-05-16

### 第 18 次发布

发布时间：2025-08-01 02:54:08

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeInstance](https://cloud.tencent.com/document/api/1778/111030)

	* 新增出参：UseDefaultServerCert, TrustedCaLimit, ServerCertLimit

* [ModifyInstance](https://cloud.tencent.com/document/api/1778/116039)

	* 新增入参：UseDefaultServerCert, X509Mode


新增数据结构：

* [SubscriptionUserProperty](https://cloud.tencent.com/document/api/1778/111031#SubscriptionUserProperty)

修改数据结构：

* [MQTTClientSubscription](https://cloud.tencent.com/document/api/1778/111031#MQTTClientSubscription)

	* 新增成员：UserProperties




## 边缘安全加速平台(teo) 版本：2022-09-01

### 第 115 次发布

发布时间：2025-07-31 16:42:36

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateJustInTimeTranscodeTemplate](https://cloud.tencent.com/document/api/1552/122116)
* [CreateSecurityAPIResource](https://cloud.tencent.com/document/api/1552/122125)
* [CreateSecurityAPIService](https://cloud.tencent.com/document/api/1552/122124)
* [CreateSecurityClientAttester](https://cloud.tencent.com/document/api/1552/122112)
* [CreateSecurityJSInjectionRule](https://cloud.tencent.com/document/api/1552/122111)
* [DeleteJustInTimeTranscodeTemplates](https://cloud.tencent.com/document/api/1552/122115)
* [DeleteSecurityAPIResource](https://cloud.tencent.com/document/api/1552/122123)
* [DeleteSecurityAPIService](https://cloud.tencent.com/document/api/1552/122122)
* [DeleteSecurityClientAttester](https://cloud.tencent.com/document/api/1552/122110)
* [DeleteSecurityJSInjectionRule](https://cloud.tencent.com/document/api/1552/122109)
* [DescribeJustInTimeTranscodeTemplates](https://cloud.tencent.com/document/api/1552/122114)
* [DescribeSecurityAPIResource](https://cloud.tencent.com/document/api/1552/122121)
* [DescribeSecurityAPIService](https://cloud.tencent.com/document/api/1552/122120)
* [DescribeSecurityClientAttester](https://cloud.tencent.com/document/api/1552/122108)
* [DescribeSecurityIPGroupContent](https://cloud.tencent.com/document/api/1552/122107)
* [DescribeSecurityJSInjectionRule](https://cloud.tencent.com/document/api/1552/122106)
* [ModifySecurityAPIResource](https://cloud.tencent.com/document/api/1552/122119)
* [ModifySecurityAPIService](https://cloud.tencent.com/document/api/1552/122118)
* [ModifySecurityClientAttester](https://cloud.tencent.com/document/api/1552/122105)
* [ModifySecurityJSInjectionRule](https://cloud.tencent.com/document/api/1552/122104)

新增数据结构：

* [APIResource](https://cloud.tencent.com/document/api/1552/80721#APIResource)
* [APIService](https://cloud.tencent.com/document/api/1552/80721#APIService)
* [AudioTemplateInfo](https://cloud.tencent.com/document/api/1552/80721#AudioTemplateInfo)
* [ClientAttester](https://cloud.tencent.com/document/api/1552/80721#ClientAttester)
* [JSInjectionRule](https://cloud.tencent.com/document/api/1552/80721#JSInjectionRule)
* [JustInTimeTranscodeTemplate](https://cloud.tencent.com/document/api/1552/80721#JustInTimeTranscodeTemplate)
* [TCCaptchaOption](https://cloud.tencent.com/document/api/1552/80721#TCCaptchaOption)
* [TCRCEOption](https://cloud.tencent.com/document/api/1552/80721#TCRCEOption)
* [VideoTemplateInfo](https://cloud.tencent.com/document/api/1552/80721#VideoTemplateInfo)



## 边缘安全加速平台(teo) 版本：2022-01-06



## TI-ONE 训练平台(tione) 版本：2021-11-11

### 第 89 次发布

发布时间：2025-08-01 03:40:02

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [Service](https://cloud.tencent.com/document/api/851/75051#Service)

	* 新增成员：SubUinName




## TI-ONE 训练平台(tione) 版本：2019-10-22



## 实时音视频(trtc) 版本：2019-07-22

### 第 117 次发布

发布时间：2025-08-01 03:49:38

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ControlAIConversation](https://cloud.tencent.com/document/api/647/109408)

	* 新增入参：InvokeLLM


新增数据结构：

* [InvokeLLM](https://cloud.tencent.com/document/api/647/44055#InvokeLLM)



## Web 应用防火墙(waf) 版本：2018-01-25

### 第 123 次发布

发布时间：2025-08-01 04:09:22

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateOwaspWhiteRule](https://cloud.tencent.com/document/api/627/122157)
* [DeleteOwaspWhiteRule](https://cloud.tencent.com/document/api/627/122156)
* [DescribeOwaspWhiteRules](https://cloud.tencent.com/document/api/627/122155)
* [ModifyOwaspWhiteRule](https://cloud.tencent.com/document/api/627/122154)

新增数据结构：

* [OwaspWhiteRule](https://cloud.tencent.com/document/api/627/53609#OwaspWhiteRule)



