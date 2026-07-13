# Release v1.3.133

## 应用性能监控(apm) 版本：2021-06-22

### 第 64 次发布

发布时间：2026-07-14 01:12:08

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyApmInstance](https://cloud.tencent.com/document/api/1463/89002)

	* 新增入参：EnableHeadSampler, HeadSamplerType, HeadSamplerArg




## 云数据库 MySQL(cdb) 版本：2017-03-20

### 第 222 次发布

发布时间：2026-07-14 01:20:33

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [DeliverSummary](https://cloud.tencent.com/document/api/236/15878#DeliverSummary)

修改数据结构：

* [InstanceDbAuditStatus](https://cloud.tencent.com/document/api/236/15878#InstanceDbAuditStatus)

	* 新增成员：DeliverSummary, Deliver




## 消息队列 CKafka 版(ckafka) 版本：2019-08-19

### 第 148 次发布

发布时间：2026-07-14 01:28:10

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [IsolatedInstancePre](https://cloud.tencent.com/document/api/597/134029)



## 负载均衡(clb) 版本：2018-03-17

### 第 154 次发布

发布时间：2026-07-14 01:29:22

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [ClusterInfo](https://cloud.tencent.com/document/api/214/30694#ClusterInfo)

	* 新增成员：ClusterName




## 云安全一体化平台(csip) 版本：2022-11-21

### 第 90 次发布

发布时间：2026-07-14 01:35:30

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [CosBucketId](https://cloud.tencent.com/document/api/664/90825#CosBucketId)

	* 新增成员：IsAutoMonitor




## 数据加速器 GooseFS(goosefs) 版本：2022-05-19

### 第 19 次发布

发布时间：2026-07-14 02:06:09

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [BuildCustomerCluster](https://cloud.tencent.com/document/api/1424/129191)

	* 新增入参：Zone




## 云数据库 MongoDB(mongodb) 版本：2019-07-25

### 第 73 次发布

发布时间：2026-07-14 02:24:29

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateDBInstance](https://cloud.tencent.com/document/api/240/38571)

	* 新增入参：DataEncryption, EncryptionKeySource, KeyId, KmsRegion

* [CreateDBInstanceHour](https://cloud.tencent.com/document/api/240/38570)

	* 新增入参：DataEncryption, EncryptionKeySource, KeyId, KmsRegion




## 云数据库 MongoDB(mongodb) 版本：2018-04-08



## 云数据库Redis(redis) 版本：2018-04-12

### 第 104 次发布

发布时间：2026-07-14 02:36:58

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeBackupDetail](https://cloud.tencent.com/document/api/239/107146)

	* 新增出参：Encrypted, DecryptKey, KmsKeyId, KeyAlgorithm


修改数据结构：

* [RedisBackupSet](https://cloud.tencent.com/document/api/239/20022#RedisBackupSet)

	* 新增成员：Encrypted




## TokenHub(tokenhub) 版本：2026-03-22

### 第 14 次发布

发布时间：2026-07-14 03:02:16

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateEndpoint](https://cloud.tencent.com/document/api/1823/134036)
* [DeleteEndpoint](https://cloud.tencent.com/document/api/1823/134035)
* [DescribeEndpoint](https://cloud.tencent.com/document/api/1823/134034)
* [DescribeModelEndpointList](https://cloud.tencent.com/document/api/1823/134033)
* [ModifyEndpoint](https://cloud.tencent.com/document/api/1823/134032)

新增数据结构：

* [EndpointCreateDetail](https://cloud.tencent.com/document/api/1823/132279#EndpointCreateDetail)
* [EndpointDetail](https://cloud.tencent.com/document/api/1823/132279#EndpointDetail)
* [ModelEndpointView](https://cloud.tencent.com/document/api/1823/132279#ModelEndpointView)



## TSF-Polaris&ZK&网关(tse) 版本：2020-12-07

### 第 111 次发布

发布时间：2026-07-14 03:05:35

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateCloudNativeAPIGatewayLLMModelService](https://cloud.tencent.com/document/api/1364/131967)

	* 新增入参：ModelRewriteRules, SourceId, Namespace, ServiceName, Protocol, ExtParams, CustomProviderName, KeyRotationEnabled, KeyRotationPeriodDays, ExternalInstanceId

* [CreateCloudNativeAPIGatewaySecretKey](https://cloud.tencent.com/document/api/1364/131966)

	* 新增入参：JWTCredentialConfig, OAuthCredentialConfig, OIDCCredentialConfig, Provider

* [DescribeCloudNativeAPIGatewayConsumerList](https://cloud.tencent.com/document/api/1364/131957)

	* 新增入参：ResourceType, ResourceId

* [DescribeCloudNativeAPIGatewayLLMModelAPIs](https://cloud.tencent.com/document/api/1364/131955)

	* 新增入参：ConsumerId

* [ModifyCloudNativeAPIGatewayLLMModelService](https://cloud.tencent.com/document/api/1364/131946)

	* 新增入参：ModelRewriteRules, SourceId, Namespace, ServiceName, Protocol, ExtParams, KeyRotationEnabled, KeyRotationPeriodDays, ExternalInstanceId


新增数据结构：

* [AIGWCacheAwareRouteCandidate](https://cloud.tencent.com/document/api/1364/54942#AIGWCacheAwareRouteCandidate)
* [AIGWCacheAwareRouteConfig](https://cloud.tencent.com/document/api/1364/54942#AIGWCacheAwareRouteConfig)
* [AIGWJWTCredentialConfig](https://cloud.tencent.com/document/api/1364/54942#AIGWJWTCredentialConfig)
* [AIGWLLMModelServiceSubRoute](https://cloud.tencent.com/document/api/1364/54942#AIGWLLMModelServiceSubRoute)
* [AIGWModelRewriteRule](https://cloud.tencent.com/document/api/1364/54942#AIGWModelRewriteRule)
* [AIGWOAuthCredentialConfig](https://cloud.tencent.com/document/api/1364/54942#AIGWOAuthCredentialConfig)
* [AIGWOIDCCredentialConfig](https://cloud.tencent.com/document/api/1364/54942#AIGWOIDCCredentialConfig)
* [AIGWRouteModelServiceConfig](https://cloud.tencent.com/document/api/1364/54942#AIGWRouteModelServiceConfig)
* [AIGWTokenLengthRoute](https://cloud.tencent.com/document/api/1364/54942#AIGWTokenLengthRoute)
* [AIGWTokenLengthRouteRule](https://cloud.tencent.com/document/api/1364/54942#AIGWTokenLengthRouteRule)

修改数据结构：

* [AIGWLLMQuotaLimit](https://cloud.tencent.com/document/api/1364/54942#AIGWLLMQuotaLimit)

	* 新增成员：ConcurrentCountLimit

* [AIGWLogConfig](https://cloud.tencent.com/document/api/1364/54942#AIGWLogConfig)

	* 新增成员：RequestLogPayloadMode, ResponseLogPayloadMode

* [CNAPIGwSecretKey](https://cloud.tencent.com/document/api/1364/54942#CNAPIGwSecretKey)

	* 新增成员：JWTCredentialConfig, OAuthCredentialConfig, OIDCCredentialConfig, Provider

* [CloudNativeAPIGatewayLLMModelService](https://cloud.tencent.com/document/api/1364/54942#CloudNativeAPIGatewayLLMModelService)

	* 新增成员：ModelRewriteRules, SourceId, Namespace, ServiceName, Protocol, ExtParams, CustomProviderName, KeyRotationEnabled, KeyRotationPeriodDays, ExternalInstanceId

* [CloudNativeAPIGatewayLLMModelServiceRoute](https://cloud.tencent.com/document/api/1364/54942#CloudNativeAPIGatewayLLMModelServiceRoute)

	* 新增成员：CacheAwareRouteConfig, TokenLengthRouteConfig

* [DescribeCloudNativeAPIGatewayResult](https://cloud.tencent.com/document/api/1364/54942#DescribeCloudNativeAPIGatewayResult)

	* 新增成员：ForceHTTPSRedirect




## TSF-应用管理&Consul(tsf) 版本：2018-03-26

### 第 144 次发布

发布时间：2026-07-14 03:07:18

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [ApplicationAttribute](https://cloud.tencent.com/document/api/649/36099#ApplicationAttribute)

	* 新增成员：ImageTagCount




## 云点播(vod) 版本：2024-07-18



## 云点播(vod) 版本：2018-07-17

### 第 271 次发布

发布时间：2026-07-14 03:10:41

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [AigcVideoTaskUsage](https://cloud.tencent.com/document/api/266/31773#AigcVideoTaskUsage)

修改数据结构：

* [AigcVideoTaskOutput](https://cloud.tencent.com/document/api/266/31773#AigcVideoTaskOutput)

	* 新增成员：Usage




## Web 应用防火墙(waf) 版本：2018-01-25

### 第 160 次发布

发布时间：2026-07-14 03:17:50

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [LLMPkg](https://cloud.tencent.com/document/api/627/53609#LLMPkg)

	* 新增成员：InquireNum




