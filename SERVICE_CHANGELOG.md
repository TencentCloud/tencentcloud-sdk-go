# Release v1.1.33

## 内容分发网络 CDN(cdn) 版本：2018-06-06

### 第 138 次发布

发布时间：2025-09-25 01:16:54

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除接口**：</font>

* DisableCaches
* EnableCaches
* GetDisableRecords

<font color="#dd0000">**删除数据结构**：</font>

* CacheOptResult
* UrlRecord



## 物联网开发平台(iotexplorer) 版本：2019-04-23

### 第 119 次发布

发布时间：2025-09-25 01:42:36

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateTWeTalkProductConfig](https://cloud.tencent.com/document/api/1081/122778)

	* 新增入参：FastVoiceType

* [ModifyTWeTalkProductConfig](https://cloud.tencent.com/document/api/1081/122775)

	* 新增入参：FastVoiceType




## 腾讯云智能体开发平台(lke) 版本：2023-11-30

### 第 66 次发布

发布时间：2025-09-25 01:52:27

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除接口**：</font>

* CreateAgent
* ModifyAgent



## 云数据库 MongoDB(mongodb) 版本：2019-07-25

### 第 56 次发布

发布时间：2025-09-25 01:58:45

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateAccountUser](https://cloud.tencent.com/document/api/240/84704)

	* <font color="#dd0000">**修改入参**：</font>MongoUserPassword




## 云数据库 MongoDB(mongodb) 版本：2018-04-08



## 媒体处理(mps) 版本：2019-06-12

### 第 145 次发布

发布时间：2025-09-24 16:03:12

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateSmartEraseTemplate](https://cloud.tencent.com/document/api/862/123735)
* [DeleteSmartEraseTemplate](https://cloud.tencent.com/document/api/862/123734)
* [DescribeSmartEraseTemplates](https://cloud.tencent.com/document/api/862/123733)
* [ModifySmartEraseTemplate](https://cloud.tencent.com/document/api/862/123732)

新增数据结构：

* [SmartEraseTemplateItem](https://cloud.tencent.com/document/api/862/37615#SmartEraseTemplateItem)



## 容器安全服务(tcss) 版本：2020-11-01

### 第 80 次发布

发布时间：2025-09-25 02:37:43

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [VulInfo](https://cloud.tencent.com/document/api/1285/65614#VulInfo)

	* 新增成员：RaspOpenNodeCount, RaspClosedNodeCount




## 边缘安全加速平台(teo) 版本：2022-09-01

### 第 120 次发布

发布时间：2025-09-25 02:45:59

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [MultiPathGateway](https://cloud.tencent.com/document/api/1552/80721#MultiPathGateway)

	* 新增成员：NeedConfirm




## 边缘安全加速平台(teo) 版本：2022-01-06



## TSF-Polaris&ZK&网关(tse) 版本：2020-12-07

### 第 100 次发布

发布时间：2025-09-25 03:02:40

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateNativeGatewayServiceSource](https://cloud.tencent.com/document/api/1364/113754)

	* 新增出参：SourceID

* [DescribeNativeGatewayServiceSources](https://cloud.tencent.com/document/api/1364/113752)

	* 新增入参：SourceID

* [ModifyNetworkBasicInfo](https://cloud.tencent.com/document/api/1364/102774)

	* 新增入参：SlaType


修改数据结构：

* [CloudNativeAPIGatewayConfig](https://cloud.tencent.com/document/api/1364/54942#CloudNativeAPIGatewayConfig)

	* 新增成员：CustomizedConfigContent




## Web 应用防火墙(waf) 版本：2018-01-25

### 第 128 次发布

发布时间：2025-09-25 03:22:18

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [AddCustomRule](https://cloud.tencent.com/document/api/627/53608)

	* 新增入参：ActionRatio

* [AddSpartaProtection](https://cloud.tencent.com/document/api/627/72689)

	* 新增入参：ProxyConnectTimeout, Gzip

	* <font color="#dd0000">**修改入参**：</font>HttpsRewrite, IsHttp2, ActiveCheck, CipherTemplate

* [DeleteSpartaProtection](https://cloud.tencent.com/document/api/627/96949)

	* <font color="#dd0000">**修改入参**：</font>InstanceID

* [DescribeBotSceneOverview](https://cloud.tencent.com/document/api/627/118058)

	* 新增出参：TldStatus

* [ModifyCustomRule](https://cloud.tencent.com/document/api/627/97442)

	* 新增入参：ActionRatio

* [ModifySpartaProtection](https://cloud.tencent.com/document/api/627/94309)

	* 新增入参：ProxyConnectTimeout, Gzip

* [UpsertCCRule](https://cloud.tencent.com/document/api/627/97646)

	* 新增入参：PageId, ActionRatio

* [UpsertSession](https://cloud.tencent.com/document/api/627/97645)

	* 新增入参：Key


新增数据结构：

* [JWTConfig](https://cloud.tencent.com/document/api/627/53609#JWTConfig)
* [LLMPkg](https://cloud.tencent.com/document/api/627/53609#LLMPkg)
* [RCEPkg](https://cloud.tencent.com/document/api/627/53609#RCEPkg)
* [SecretInfo](https://cloud.tencent.com/document/api/627/53609#SecretInfo)
* [TokenDisplaySetting](https://cloud.tencent.com/document/api/627/53609#TokenDisplaySetting)
* [TokenRuleEntry](https://cloud.tencent.com/document/api/627/53609#TokenRuleEntry)
* [TokenRuleEntryValue](https://cloud.tencent.com/document/api/627/53609#TokenRuleEntryValue)
* [TokenValidation](https://cloud.tencent.com/document/api/627/53609#TokenValidation)
* [TokenVerifyRule](https://cloud.tencent.com/document/api/627/53609#TokenVerifyRule)

修改数据结构：

* [ApiDataFilter](https://cloud.tencent.com/document/api/627/53609#ApiDataFilter)

	* 新增成员：ValueList

* [BotToken](https://cloud.tencent.com/document/api/627/53609#BotToken)

	* 新增成员：TokenValidation

* [CCRuleItems](https://cloud.tencent.com/document/api/627/53609#CCRuleItems)

	* 新增成员：PageId, ActionRatio

* [DescribeCustomRulesRspRuleListItem](https://cloud.tencent.com/document/api/627/53609#DescribeCustomRulesRspRuleListItem)

	* 新增成员：ActionRatio

* [DomainsPartInfo](https://cloud.tencent.com/document/api/627/53609#DomainsPartInfo)

	* 新增成员：ProxyConnectTimeout, Gzip

* [InstanceInfo](https://cloud.tencent.com/document/api/627/53609#InstanceInfo)

	* 新增成员：RCEPkg, ExceedPolicy, LLMPkg, ElasticResourceId

* [SessionItem](https://cloud.tencent.com/document/api/627/53609#SessionItem)

	* 新增成员：Key




