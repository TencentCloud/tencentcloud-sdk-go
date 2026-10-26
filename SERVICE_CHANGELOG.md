# Release v1.3.200

## 腾讯云智能体开发平台(adp) 版本：2026-05-20

### 第 32 次发布

发布时间：2026-10-27 01:03:41

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [ConversationExperience](https://cloud.tencent.com/document/api/1759/132545#ConversationExperience)

	* 新增成员：EnableNewMessageInterrupt




## 腾讯云数据仓库 TCHouse-D(cdwdoris) 版本：2021-12-28

### 第 68 次发布

发布时间：2026-10-27 01:17:48

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyUserPrivilegesV3](https://cloud.tencent.com/document/api/1387/109432)

	* 新增入参：GrantPrivileges, RevokePrivileges

	* <font color="#dd0000">**修改入参**：</font>UserPrivileges

	* 新增出参：SkippedItems


新增数据结构：

* [SkippedPrivilegeItem](https://cloud.tencent.com/document/api/1387/102385#SkippedPrivilegeItem)



## 消息队列 CKafka 版(ckafka) 版本：2019-08-19

### 第 154 次发布

发布时间：2026-10-27 01:21:30

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [CosParam](https://cloud.tencent.com/document/api/597/40861#CosParam)

	* 新增成员：Timezone, SourceDataFormat

* [KafkaConnectParam](https://cloud.tencent.com/document/api/597/40861#KafkaConnectParam)

	* 新增成员：OriginalBrokerAddresses




## 云原生智能网关(cngw) 版本：2023-04-18

### 第 10 次发布

发布时间：2026-10-27 01:26:34

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateCloudNativeAPIGatewayConsumer](https://cloud.tencent.com/document/api/1826/133125)

	* 新增入参：OperationSource, Status, SecretBindType, AutoSecretKey

* [CreateCloudNativeAPIGatewayConsumerGroup](https://cloud.tencent.com/document/api/1826/133124)

	* 新增入参：OperationSource, ParentConsumerGroupIds

* [CreateCloudNativeAPIGatewayLLMModelAPI](https://cloud.tencent.com/document/api/1826/133115)

	* 新增入参：PriorityScheduleConfig, ModelNameRoute, StandardModelName

* [CreateCloudNativeAPIGatewayLLMModelService](https://cloud.tencent.com/document/api/1826/133132)

	* 新增入参：SmartBalance, LengthLimitConfig

* [CreateCloudNativeAPIGatewayMCPServer](https://cloud.tencent.com/document/api/1826/133160)

	* 新增入参：ParamRewriteConfig

* [DeleteCloudNativeAPIGatewayConsumer](https://cloud.tencent.com/document/api/1826/133123)

	* 新增入参：OperationSource

* [DeleteCloudNativeAPIGatewayConsumerGroup](https://cloud.tencent.com/document/api/1826/133122)

	* 新增入参：OperationSource

* [DescribeCloudNativeAPIGatewayLLMModelAPIs](https://cloud.tencent.com/document/api/1826/133112)

	* 新增入参：SecretKeyId

* [DescribeCloudNativeAPIGatewaySecretKey](https://cloud.tencent.com/document/api/1826/133139)

	* 新增入参：IncludeCreationInfo

* [DescribeCloudNativeAPIGatewaySecretKeyList](https://cloud.tencent.com/document/api/1826/137931)

	* 新增入参：IncludeCreationInfo

* [ModifyCloudNativeAPIGatewayLLMModelAPI](https://cloud.tencent.com/document/api/1826/133111)

	* 新增入参：ModelNameRoute, PriorityScheduleConfig, StandardModelName

* [ModifyCloudNativeAPIGatewayLLMModelService](https://cloud.tencent.com/document/api/1826/133128)

	* 新增入参：SmartBalance

* [ModifyCloudNativeAPIGatewayMCPServer](https://cloud.tencent.com/document/api/1826/133149)

	* 新增入参：ParamRewriteConfig


新增数据结构：

* [AIGWGatewayVerifyConfig](https://cloud.tencent.com/document/api/1826/133161#AIGWGatewayVerifyConfig)
* [AIGWGatewayVerifyOAuthConfig](https://cloud.tencent.com/document/api/1826/133161#AIGWGatewayVerifyOAuthConfig)
* [AIGWLLMResourceDimension](https://cloud.tencent.com/document/api/1826/133161#AIGWLLMResourceDimension)
* [AIGWLengthLimitConfig](https://cloud.tencent.com/document/api/1826/133161#AIGWLengthLimitConfig)
* [AIGWMCPServerParamRewriteConfig](https://cloud.tencent.com/document/api/1826/133161#AIGWMCPServerParamRewriteConfig)
* [AIGWMCPServerParamRewriteRule](https://cloud.tencent.com/document/api/1826/133161#AIGWMCPServerParamRewriteRule)
* [AIGWMCPServerRequestRewriteConfig](https://cloud.tencent.com/document/api/1826/133161#AIGWMCPServerRequestRewriteConfig)
* [AIGWModelNameRoute](https://cloud.tencent.com/document/api/1826/133161#AIGWModelNameRoute)
* [AIGWPriorityScheduleConfig](https://cloud.tencent.com/document/api/1826/133161#AIGWPriorityScheduleConfig)
* [AIGWQuotaSubItem](https://cloud.tencent.com/document/api/1826/133161#AIGWQuotaSubItem)
* [AIGWSmartBalanceConfig](https://cloud.tencent.com/document/api/1826/133161#AIGWSmartBalanceConfig)
* [AutoSecretKey](https://cloud.tencent.com/document/api/1826/133161#AutoSecretKey)

修改数据结构：

* [AIGWConsumerGroupBrief](https://cloud.tencent.com/document/api/1826/133161#AIGWConsumerGroupBrief)

	* 新增成员：ConsumerGroupPath

* [AIGWIntentRoute](https://cloud.tencent.com/document/api/1826/133161#AIGWIntentRoute)

	* 新增成员：DefaultStandardModelName, IntentStandardModelName

* [AIGWIntentRouteRule](https://cloud.tencent.com/document/api/1826/133161#AIGWIntentRouteRule)

	* 新增成员：IntentId, IntentName, StandardModelName

* [AIGWLLMTokenUsageItem](https://cloud.tencent.com/document/api/1826/133161#AIGWLLMTokenUsageItem)

	* 新增成员：DimensionList

* [AIGWLLMTokenUsageListResult](https://cloud.tencent.com/document/api/1826/133161#AIGWLLMTokenUsageListResult)

	* 新增成员：OrgAttributionIncomplete

* [AIGWLLMTokenUsageStatisticsResult](https://cloud.tencent.com/document/api/1826/133161#AIGWLLMTokenUsageStatisticsResult)

	* 新增成员：TotalTokens, OrgAttributionIncomplete

* [AIGWMCPServer](https://cloud.tencent.com/document/api/1826/133161#AIGWMCPServer)

	* 新增成员：ParamRewriteConfig, GatewayVerifyConfig

* [AIGWQuota](https://cloud.tencent.com/document/api/1826/133161#AIGWQuota)

	* 新增成员：ExceedPercentage, SetExceedPercentage, Scope, SubResourceType, SubResourceCount, ConsumerGroupPath, ConsumerGroupActive, Used, UsedExact, UsageRate, AlarmLevel, SecretKeyCount

* [AIGWQuotaDetail](https://cloud.tencent.com/document/api/1826/133161#AIGWQuotaDetail)

	* 新增成员：SetExceedPercentage, ExceedPercentage, UsedExact, ConsumerGroupPath, ConsumerGroupActive, SubResources

* [AIGWSensitiveWordRoute](https://cloud.tencent.com/document/api/1826/133161#AIGWSensitiveWordRoute)

	* 新增成员：StandardModelName

* [CNAPIGwConsumer](https://cloud.tencent.com/document/api/1826/133161#CNAPIGwConsumer)

	* 新增成员：DisplayName, UserName, Status, Origin, SyncSourceId, ExternalId

* [CNAPIGwConsumerGroup](https://cloud.tencent.com/document/api/1826/133161#CNAPIGwConsumerGroup)

	* 新增成员：ExternalId

* [CNAPIGwSecretKey](https://cloud.tencent.com/document/api/1826/133161#CNAPIGwSecretKey)

	* 新增成员：CredentialGenerateType, CreationSource, BatchId

* [CloudNativeAPIGatewayLLMModelAPI](https://cloud.tencent.com/document/api/1826/133161#CloudNativeAPIGatewayLLMModelAPI)

	* 新增成员：ConsumerDirectModelScopes, SecretKeyModelScopes, PriorityScheduleConfig, PlaygroundAvailable, ModelNameRoute, StandardModelName

* [CloudNativeAPIGatewayLLMModelService](https://cloud.tencent.com/document/api/1826/133161#CloudNativeAPIGatewayLLMModelService)

	* 新增成员：LengthLimitConfig, SmartBalance

* [CloudNativeAPIGatewayLLMModelServiceRouteModelNameStrategy](https://cloud.tencent.com/document/api/1826/133161#CloudNativeAPIGatewayLLMModelServiceRouteModelNameStrategy)

	* 新增成员：StandardModelName

* [CloudNativeAPIGatewayLLMModelServiceRouteWeightedStrategy](https://cloud.tencent.com/document/api/1826/133161#CloudNativeAPIGatewayLLMModelServiceRouteWeightedStrategy)

	* 新增成员：StandardModelName




## 云安全一体化平台(csip) 版本：2022-11-21

### 第 112 次发布

发布时间：2026-10-27 01:28:11

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyVulScanPeriodic](https://cloud.tencent.com/document/api/664/135838)

	* 新增入参：PreferCloudSource




## TDSQL-C MySQL 版(cynosdb) 版本：2019-01-07

### 第 199 次发布

发布时间：2026-10-27 01:38:53

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [AddClusterSlaveZone](https://cloud.tencent.com/document/api/1003/79555)

	* 新增出参：BigDealIds, DealNames, TranId




## 云数据库独享集群(dbdc) 版本：2020-10-29

### 第 17 次发布

发布时间：2026-10-27 01:42:57

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeDBCustomTaskStatus](https://cloud.tencent.com/document/api/1322/132922)

	* 新增出参：ResourceStatusSet


新增数据结构：

* [TaskResourceStatus](https://cloud.tencent.com/document/api/1322/74754#TaskResourceStatus)

修改数据结构：

* [DBCustomNode](https://cloud.tencent.com/document/api/1322/74754#DBCustomNode)

	* 新增成员：BusinessType

* [DataDisk](https://cloud.tencent.com/document/api/1322/74754#DataDisk)

	* 新增成员：Encrypt, DiskId, SnapshotId, ClientDiskKey

* [SystemDisk](https://cloud.tencent.com/document/api/1322/74754#SystemDisk)

	* 新增成员：DiskId




## 数据加速器 GooseFS(goosefs) 版本：2022-05-19

### 第 25 次发布

发布时间：2026-10-27 01:59:14

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeDataRepositoryTasks](https://cloud.tencent.com/document/api/1424/140070)

新增数据结构：

* [DataRepositoryTaskAttribute](https://cloud.tencent.com/document/api/1424/95076#DataRepositoryTaskAttribute)



## 云直播CSS(live) 版本：2018-08-01

### 第 193 次发布

发布时间：2026-10-27 02:11:26

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [SnapshotTemplateInfo](https://cloud.tencent.com/document/api/267/20474#SnapshotTemplateInfo)

	* 新增成员：CosSecretId, CosSecretKey




## 消息队列 MQTT 版(mqtt) 版本：2024-05-16

### 第 36 次发布

发布时间：2026-10-27 02:20:08

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeMessageByTopic](https://cloud.tencent.com/document/api/1778/118936)

	* 新增入参：FilterExpression

* [DescribeMessageDetails](https://cloud.tencent.com/document/api/1778/120249)

	* 新增出参：Retain

* [RegisterCaCertificate](https://cloud.tencent.com/document/api/1778/116205)

	* <font color="#dd0000">**修改入参**：</font>VerificationCertificate




## 私有域解析 Private DNS(privatedns) 版本：2020-10-28

### 第 28 次发布

发布时间：2026-10-27 02:25:22

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeExtendEndpointList](https://cloud.tencent.com/document/api/1338/124330)

	* 新增出参：TotalCount, OutboundEndpointSet


新增数据结构：

* [EndpointService](https://cloud.tencent.com/document/api/1338/55947#EndpointService)
* [OutboundEndpoint](https://cloud.tencent.com/document/api/1338/55947#OutboundEndpoint)



## 云开发 CloudBase(tcb) 版本：2018-06-08

### 第 174 次发布

发布时间：2026-10-27 02:33:57

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除接口**：</font>

* CreateUserKey
* DeleteUserKey
* GetOneUserKey
* GetUserKeyList
* UpdateUserKey

<font color="#dd0000">**删除数据结构**：</font>

* APIsUserKeyInfo
* AgentRuntimeCode
* AgentRuntimeConfig
* AgentRuntimeSessionConfig
* BindService
* BindServiceList

修改数据结构：

* [HTTPServiceRoute](https://cloud.tencent.com/document/api/876/34822#HTTPServiceRoute)

	* 新增成员：UpstreamResourceVersion

* [HTTPServiceRouteParam](https://cloud.tencent.com/document/api/876/34822#HTTPServiceRouteParam)

	* 新增成员：UpstreamResourceVersion




## TSF-Polaris&ZK&网关(tse) 版本：2020-12-07

### 第 115 次发布

发布时间：2026-10-27 02:49:24

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateCloudNativeAPIGatewayConsumer](https://cloud.tencent.com/document/api/1364/131970)

	* 新增入参：SecretBindType, AutoSecretKey, OperationSource, Status

* [CreateCloudNativeAPIGatewayConsumerGroup](https://cloud.tencent.com/document/api/1364/131969)

	* 新增入参：OperationSource, ParentConsumerGroupIds

* [CreateCloudNativeAPIGatewayLLMModelAPI](https://cloud.tencent.com/document/api/1364/131968)

	* 新增入参：PriorityScheduleConfig, ModelNameRoute, StandardModelName

* [CreateCloudNativeAPIGatewayLLMModelService](https://cloud.tencent.com/document/api/1364/131967)

	* 新增入参：LengthLimitConfig, SmartBalance

* [DeleteCloudNativeAPIGatewayConsumer](https://cloud.tencent.com/document/api/1364/131965)

	* 新增入参：OperationSource

* [DeleteCloudNativeAPIGatewayConsumerGroup](https://cloud.tencent.com/document/api/1364/131964)

	* 新增入参：OperationSource

* [DescribeCloudNativeAPIGatewayCanaryRules](https://cloud.tencent.com/document/api/1364/94844)

	* 新增入参：OrderField, OrderType, Filters

* [DescribeCloudNativeAPIGatewayConsumerList](https://cloud.tencent.com/document/api/1364/131957)

	* 新增入参：Origin, Status, SyncSourceId, WithoutSyncSource

* [DescribeCloudNativeAPIGatewayLLMModelAPIs](https://cloud.tencent.com/document/api/1364/131955)

	* 新增入参：SecretKeyId

* [DescribeCloudNativeAPIGatewaySecretKey](https://cloud.tencent.com/document/api/1364/131952)

	* 新增入参：IncludeCreationInfo

* [DescribeCloudNativeAPIGatewaySecretKeyList](https://cloud.tencent.com/document/api/1364/131951)

	* 新增入参：ResourceId, UseToBind, Filters, Keyword, IncludeCreationInfo

* [DescribeCloudNativeAPIGatewaySecretKeyValue](https://cloud.tencent.com/document/api/1364/131950)

	* 新增入参：SecretValueType

* [ModifyCloudNativeAPIGatewayLLMModelAPI](https://cloud.tencent.com/document/api/1364/131947)

	* 新增入参：ModelNameRoute, PriorityScheduleConfig, StandardModelName

* [ModifyCloudNativeAPIGatewayLLMModelService](https://cloud.tencent.com/document/api/1364/131946)

	* 新增入参：LengthLimitConfig, SmartBalance


新增数据结构：

* [AIGWLengthLimitConfig](https://cloud.tencent.com/document/api/1364/54942#AIGWLengthLimitConfig)
* [AIGWModelNameRoute](https://cloud.tencent.com/document/api/1364/54942#AIGWModelNameRoute)
* [AIGWPriorityScheduleConfig](https://cloud.tencent.com/document/api/1364/54942#AIGWPriorityScheduleConfig)
* [AIGWSimpleSecretKey](https://cloud.tencent.com/document/api/1364/54942#AIGWSimpleSecretKey)
* [AIGWSmartBalanceConfig](https://cloud.tencent.com/document/api/1364/54942#AIGWSmartBalanceConfig)
* [AutoSecretKey](https://cloud.tencent.com/document/api/1364/54942#AutoSecretKey)

修改数据结构：

* [AIGWIntentRoute](https://cloud.tencent.com/document/api/1364/54942#AIGWIntentRoute)

	* 新增成员：DefaultStandardModelName, IntentStandardModelName

* [AIGWIntentRouteRule](https://cloud.tencent.com/document/api/1364/54942#AIGWIntentRouteRule)

	* 新增成员：IntentId, IntentName, StandardModelName

* [AIGWLatencyPriorityConfig](https://cloud.tencent.com/document/api/1364/54942#AIGWLatencyPriorityConfig)

	* 新增成员：ProbeIntervalMs, ProbeTimeoutMs

* [AIGWOAuthCredentialConfig](https://cloud.tencent.com/document/api/1364/54942#AIGWOAuthCredentialConfig)

	* 新增成员：RedirectURIs

* [AIGWSensitiveWordRoute](https://cloud.tencent.com/document/api/1364/54942#AIGWSensitiveWordRoute)

	* 新增成员：StandardModelName

* [CNAPIGwSecretKey](https://cloud.tencent.com/document/api/1364/54942#CNAPIGwSecretKey)

	* 新增成员：BoundModelSecretKeys, BoundConsumerSecretKeys, CredentialGenerateType, CreationSource, BatchId

* [CloudNativeAPIGatewayLLMModelAPI](https://cloud.tencent.com/document/api/1364/54942#CloudNativeAPIGatewayLLMModelAPI)

	* 新增成员：SecretKeyModelScopes, PlaygroundAvailable, PriorityScheduleConfig, ConsumerDirectModelScopes, ModelNameRoute, StandardModelName

* [CloudNativeAPIGatewayLLMModelService](https://cloud.tencent.com/document/api/1364/54942#CloudNativeAPIGatewayLLMModelService)

	* 新增成员：LengthLimitConfig, SmartBalance

* [CloudNativeAPIGatewayLLMModelServiceRouteModelNameStrategy](https://cloud.tencent.com/document/api/1364/54942#CloudNativeAPIGatewayLLMModelServiceRouteModelNameStrategy)

	* 新增成员：StandardModelName

* [CloudNativeAPIGatewayLLMModelServiceRouteWeightedStrategy](https://cloud.tencent.com/document/api/1364/54942#CloudNativeAPIGatewayLLMModelServiceRouteWeightedStrategy)

	* 新增成员：StandardModelName

* [DeleteCloudNativeAPIGatewayResult](https://cloud.tencent.com/document/api/1364/54942#DeleteCloudNativeAPIGatewayResult)

	* 新增成员：TaskId




## 数据开发治理平台 WeData(wedata) 版本：2025-08-06



## 数据开发治理平台 WeData(wedata) 版本：2021-08-20

### 第 208 次发布

发布时间：2026-10-27 03:01:31

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [InstanceOpsDto](https://cloud.tencent.com/document/api/1267/76336#InstanceOpsDto)

	* 新增成员：TriggerSource, ParentSpTaskId, ParentSpInstanceName, ParentSpInstanceDataTime




