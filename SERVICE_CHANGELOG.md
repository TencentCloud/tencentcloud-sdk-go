# Release v1.3.135

## 日志服务(cls) 版本：2020-10-16

### 第 168 次发布

发布时间：2026-07-16 01:33:04

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateAgentApplication](https://cloud.tencent.com/document/api/614/134210)
* [CreateRemoteWriteTask](https://cloud.tencent.com/document/api/614/134203)
* [CreateS3Recharge](https://cloud.tencent.com/document/api/614/134216)
* [DeleteAgentApplication](https://cloud.tencent.com/document/api/614/134209)
* [DeleteRemoteWriteTask](https://cloud.tencent.com/document/api/614/134202)
* [DeleteS3Recharge](https://cloud.tencent.com/document/api/614/134215)
* [DescribeAgentApplications](https://cloud.tencent.com/document/api/614/134208)
* [DescribeAgentConfigs](https://cloud.tencent.com/document/api/614/134205)
* [DescribeRemoteWriteTasks](https://cloud.tencent.com/document/api/614/134201)
* [DescribeS3Recharges](https://cloud.tencent.com/document/api/614/134214)
* [ModifyAgentApplication](https://cloud.tencent.com/document/api/614/134207)
* [ModifyRemoteWriteTask](https://cloud.tencent.com/document/api/614/134200)
* [ModifyS3Recharge](https://cloud.tencent.com/document/api/614/134213)
* [SearchS3RechargeInfo](https://cloud.tencent.com/document/api/614/134212)

新增数据结构：

* [AgentApplicationInfo](https://cloud.tencent.com/document/api/614/56471#AgentApplicationInfo)
* [AgentTopicInfo](https://cloud.tencent.com/document/api/614/56471#AgentTopicInfo)
* [LogConfigInfo](https://cloud.tencent.com/document/api/614/56471#LogConfigInfo)
* [RemoteWriteAuthInfo](https://cloud.tencent.com/document/api/614/56471#RemoteWriteAuthInfo)
* [RemoteWriteInfo](https://cloud.tencent.com/document/api/614/56471#RemoteWriteInfo)
* [S3RechargeInfo](https://cloud.tencent.com/document/api/614/56471#S3RechargeInfo)
* [ServiceLogConfigInfo](https://cloud.tencent.com/document/api/614/56471#ServiceLogConfigInfo)



## 云原生智能网关(cngw) 版本：2023-04-18

### 第 3 次发布

发布时间：2026-07-16 01:36:23

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateCloudNativeAPIGatewayLLMModelService](https://cloud.tencent.com/document/api/1826/133132)

	* 新增入参：ModelRewriteRules, CustomProviderName, ExternalInstanceId, ExtParams, KeyRotationEnabled, KeyRotationPeriodDays

* [CreateCloudNativeAPIGatewaySecretKey](https://cloud.tencent.com/document/api/1826/133141)

	* 新增入参：JWTCredentialConfig, OAuthCredentialConfig, OIDCCredentialConfig

* [ModifyCloudNativeAPIGatewayLLMModelService](https://cloud.tencent.com/document/api/1826/133128)

	* 新增入参：ModelRewriteRules, ExternalInstanceId, ExtParams, KeyRotationEnabled, KeyRotationPeriodDays

* [ModifyCloudNativeAPIGatewayMCPServerAuth](https://cloud.tencent.com/document/api/1826/133147)

	* 新增入参：JWTAuthConfig, OAuthAuthConfig, OIDCAuthConfig


新增数据结构：

* [AIGWCacheAwareRouteCandidate](https://cloud.tencent.com/document/api/1826/133161#AIGWCacheAwareRouteCandidate)
* [AIGWCacheAwareRouteConfig](https://cloud.tencent.com/document/api/1826/133161#AIGWCacheAwareRouteConfig)
* [AIGWJWTAuthPluginConfig](https://cloud.tencent.com/document/api/1826/133161#AIGWJWTAuthPluginConfig)
* [AIGWJWTCredentialConfig](https://cloud.tencent.com/document/api/1826/133161#AIGWJWTCredentialConfig)
* [AIGWLLMModelServiceSubRoute](https://cloud.tencent.com/document/api/1826/133161#AIGWLLMModelServiceSubRoute)
* [AIGWModelRewriteRule](https://cloud.tencent.com/document/api/1826/133161#AIGWModelRewriteRule)
* [AIGWOAuthAuthPluginConfig](https://cloud.tencent.com/document/api/1826/133161#AIGWOAuthAuthPluginConfig)
* [AIGWOAuthCredentialConfig](https://cloud.tencent.com/document/api/1826/133161#AIGWOAuthCredentialConfig)
* [AIGWOIDCAuthPluginConfig](https://cloud.tencent.com/document/api/1826/133161#AIGWOIDCAuthPluginConfig)
* [AIGWOIDCCredentialConfig](https://cloud.tencent.com/document/api/1826/133161#AIGWOIDCCredentialConfig)
* [AIGWRouteModelServiceConfig](https://cloud.tencent.com/document/api/1826/133161#AIGWRouteModelServiceConfig)
* [AIGWTokenLengthRoute](https://cloud.tencent.com/document/api/1826/133161#AIGWTokenLengthRoute)
* [AIGWTokenLengthRouteRule](https://cloud.tencent.com/document/api/1826/133161#AIGWTokenLengthRouteRule)
* [KeyValue](https://cloud.tencent.com/document/api/1826/133161#KeyValue)

修改数据结构：

* [AIGWLogConfig](https://cloud.tencent.com/document/api/1826/133161#AIGWLogConfig)

	* 新增成员：RequestLogPayloadMode, ResponseLogPayloadMode

* [AIGWMCPServerAuthResult](https://cloud.tencent.com/document/api/1826/133161#AIGWMCPServerAuthResult)

	* 新增成员：JWTAuthConfig, OAuthAuthConfig, OIDCAuthConfig

* [CNAPIGwSecretKey](https://cloud.tencent.com/document/api/1826/133161#CNAPIGwSecretKey)

	* 新增成员：JWTCredentialConfig, OAuthCredentialConfig, OIDCCredentialConfig, Provider

* [CloudNativeAPIGatewayLLMModelService](https://cloud.tencent.com/document/api/1826/133161#CloudNativeAPIGatewayLLMModelService)

	* 新增成员：ModelRewriteRules, SourceId, Namespace, ServiceName, Protocol, ExtParams, CustomProviderName, KeyRotationEnabled, KeyRotationPeriodDays, ExternalInstanceId

* [CloudNativeAPIGatewayLLMModelServiceRoute](https://cloud.tencent.com/document/api/1826/133161#CloudNativeAPIGatewayLLMModelServiceRoute)

	* 新增成员：CacheAwareRouteConfig, TokenLengthRouteConfig




## 暴露面管理服务(ctem) 版本：2023-11-28

### 第 19 次发布

发布时间：2026-07-16 01:40:56

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateCustomer](https://cloud.tencent.com/document/api/1755/120315)

	* 新增入参：ScanPriority

* [CreateJobRecord](https://cloud.tencent.com/document/api/1755/120319)

	* 新增入参：ScanPriority

* [ModifyCustomer](https://cloud.tencent.com/document/api/1755/120313)

	* 新增入参：ScanPriority


新增数据结构：

* [ScanPriorityDisplay](https://cloud.tencent.com/document/api/1755/120320#ScanPriorityDisplay)
* [ScanPriorityReq](https://cloud.tencent.com/document/api/1755/120320#ScanPriorityReq)

修改数据结构：

* [Customer](https://cloud.tencent.com/document/api/1755/120320#Customer)

	* 新增成员：ScanPriority

* [DisplayLeakageCode](https://cloud.tencent.com/document/api/1755/120320#DisplayLeakageCode)

	* 新增成员：RepoNamespace, RepoName, AuthorName




## 腾讯云数据分析智能体(dataagent) 版本：2025-05-13

### 第 20 次发布

发布时间：2026-07-16 01:51:13

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ChatAI](https://cloud.tencent.com/document/api/1800/125015)

	* 新增入参：ArchVersion




## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 308 次发布

发布时间：2026-07-16 02:06:56

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateSealPolicy](https://cloud.tencent.com/document/api/1323/86183)

	* 新增入参：AuthorizationFlows


新增数据结构：

* [SealPolicyAuthorizationFlows](https://cloud.tencent.com/document/api/1323/70369#SealPolicyAuthorizationFlows)



## 腾讯电子签（基础版）(essbasic) 版本：2021-05-26

### 第 265 次发布

发布时间：2026-07-16 02:09:02

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyOrganizationBusinessInfo](https://cloud.tencent.com/document/api/1420/134160)

	* 新增入参：NewLegalMobile




## 腾讯电子签（基础版）(essbasic) 版本：2020-12-22



## 全球加速(ga2) 版本：2025-01-15

### 第 10 次发布

发布时间：2026-07-16 02:12:00

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateListenerAdditionalCert](https://cloud.tencent.com/document/api/1817/134219)
* [DeleteListenerAdditionalCert](https://cloud.tencent.com/document/api/1817/134218)
* [ReplaceListenerAdditionalCert](https://cloud.tencent.com/document/api/1817/134217)



## 物联网开发平台(iotexplorer) 版本：2019-04-23

### 第 147 次发布

发布时间：2026-07-16 02:18:50

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [BindTWeTalkAgent](https://cloud.tencent.com/document/api/1081/134227)
* [CreateTWeTalkAgent](https://cloud.tencent.com/document/api/1081/134226)
* [DeleteTWeTalkAgent](https://cloud.tencent.com/document/api/1081/134225)
* [DescribeTWeTalkAgent](https://cloud.tencent.com/document/api/1081/134224)
* [DescribeTWeTalkAgentBinding](https://cloud.tencent.com/document/api/1081/134223)
* [DescribeTWeTalkAgentList](https://cloud.tencent.com/document/api/1081/134222)
* [ModifyTWeTalkAgent](https://cloud.tencent.com/document/api/1081/134221)
* [UnbindTWeTalkAgent](https://cloud.tencent.com/document/api/1081/134220)

新增数据结构：

* [TalkAgentBinding](https://cloud.tencent.com/document/api/1081/34988#TalkAgentBinding)
* [TalkAgentInfo](https://cloud.tencent.com/document/api/1081/34988#TalkAgentInfo)
* [TalkConversationConfig](https://cloud.tencent.com/document/api/1081/34988#TalkConversationConfig)
* [TalkIOTTool](https://cloud.tencent.com/document/api/1081/34988#TalkIOTTool)
* [TalkLLMConfig](https://cloud.tencent.com/document/api/1081/34988#TalkLLMConfig)
* [TalkMemoryConfig](https://cloud.tencent.com/document/api/1081/34988#TalkMemoryConfig)
* [TalkSTTConfig](https://cloud.tencent.com/document/api/1081/34988#TalkSTTConfig)
* [TalkSTTTRTC](https://cloud.tencent.com/document/api/1081/34988#TalkSTTTRTC)
* [TalkTTSConfig](https://cloud.tencent.com/document/api/1081/34988#TalkTTSConfig)
* [TalkTTSFlow](https://cloud.tencent.com/document/api/1081/34988#TalkTTSFlow)
* [TalkWebhookAuth](https://cloud.tencent.com/document/api/1081/34988#TalkWebhookAuth)
* [TalkWebhookEndpoint](https://cloud.tencent.com/document/api/1081/34988#TalkWebhookEndpoint)
* [TalkWebhookTool](https://cloud.tencent.com/document/api/1081/34988#TalkWebhookTool)



## 实时互动-教育版(lcic) 版本：2022-08-17

### 第 86 次发布

发布时间：2026-07-16 02:27:34

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeEditVersions](https://cloud.tencent.com/document/api/1639/134230)
* [GetEditVersionToken](https://cloud.tencent.com/document/api/1639/134229)
* [SetMainEditVersion](https://cloud.tencent.com/document/api/1639/134228)

新增数据结构：

* [EditVersions](https://cloud.tencent.com/document/api/1639/81423#EditVersions)



## 云数据库 MongoDB(mongodb) 版本：2019-07-25

### 第 74 次发布

发布时间：2026-07-16 02:34:51

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DeleteAccountUser](https://cloud.tencent.com/document/api/240/90048)

	* <font color="#dd0000">**修改入参**：</font>MongoUserPassword




## 云数据库 MongoDB(mongodb) 版本：2018-04-08



## 媒体处理(mps) 版本：2019-06-12

### 第 217 次发布

发布时间：2026-07-16 02:37:55

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateAiDramaTask](https://cloud.tencent.com/document/api/862/134233)
* [CreateVideoRedrawTask](https://cloud.tencent.com/document/api/862/134232)

新增数据结构：

* [AiDramaInput](https://cloud.tencent.com/document/api/862/37615#AiDramaInput)
* [VideoDramaCosInfo](https://cloud.tencent.com/document/api/862/37615#VideoDramaCosInfo)
* [VideoRedrawCosInfo](https://cloud.tencent.com/document/api/862/37615#VideoRedrawCosInfo)
* [VideoRedrawInput](https://cloud.tencent.com/document/api/862/37615#VideoRedrawInput)

修改数据结构：

* [AdaptiveDynamicStreamingTaskInput](https://cloud.tencent.com/document/api/862/37615#AdaptiveDynamicStreamingTaskInput)

	* 新增成员：StdExtStreamInfos




## 流计算 Oceanus(oceanus) 版本：2019-04-22

### 第 88 次发布

发布时间：2026-07-16 02:42:07

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateOceanusCluster](https://cloud.tencent.com/document/api/849/134237)
* [DeleteOceanusCluster](https://cloud.tencent.com/document/api/849/134236)
* [RenewOceanusCluster](https://cloud.tencent.com/document/api/849/134235)
* [ScaleOceanusCluster](https://cloud.tencent.com/document/api/849/134234)

新增数据结构：

* [SlaveVpcDescriptions](https://cloud.tencent.com/document/api/849/52010#SlaveVpcDescriptions)
* [VPCDescription](https://cloud.tencent.com/document/api/849/52010#VPCDescription)



## 云点播(vod) 版本：2024-07-18



## 云点播(vod) 版本：2018-07-17

### 第 273 次发布

发布时间：2026-07-16 03:22:34

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [UpdateAigcApiToken](https://cloud.tencent.com/document/api/266/134238)

修改接口：

* [DescribeAigcApiTokens](https://cloud.tencent.com/document/api/266/128052)

	* 新增出参：ExtInfos




