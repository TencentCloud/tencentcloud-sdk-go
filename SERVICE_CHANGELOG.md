# Release v1.3.178

## 腾讯云智能体开发平台(adp) 版本：2026-05-20

### 第 22 次发布

发布时间：2026-09-11 01:08:40

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateSkillShare](https://cloud.tencent.com/document/api/1759/133495)

	* 新增入参：CorpShareConfig

* [DescribePluginSummaryList](https://cloud.tencent.com/document/api/1759/132499)

	* 新增入参：PluginSpaceRelation

* [DescribeSkillReferenceList](https://cloud.tencent.com/document/api/1759/133490)

	* 新增出参：AllowForceModify


新增数据结构：

* [AccessKeyAuthConfig](https://cloud.tencent.com/document/api/1759/132545#AccessKeyAuthConfig)
* [AccessKeyParamConfig](https://cloud.tencent.com/document/api/1759/132545#AccessKeyParamConfig)
* [AccessKeyPassThroughConfig](https://cloud.tencent.com/document/api/1759/132545#AccessKeyPassThroughConfig)

修改数据结构：

* [AuthConfig](https://cloud.tencent.com/document/api/1759/132545#AuthConfig)

	* 新增成员：AccessKeyAuthConfig

* [CamAuthConfig](https://cloud.tencent.com/document/api/1759/132545#CamAuthConfig)

	* 新增成员：ParamList, SupportRoleAuth

* [PluginSummary](https://cloud.tencent.com/document/api/1759/132545#PluginSummary)

	* 新增成员：IsShared, SpaceId, UpdateTime




## 云防火墙(cfw) 版本：2019-09-04

### 第 114 次发布

发布时间：2026-09-11 01:28:40

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeNDRDataLeakOutAlertList](https://cloud.tencent.com/document/api/1132/137922)

新增数据结构：

* [DataLeakOutAlertEvent](https://cloud.tencent.com/document/api/1132/49071#DataLeakOutAlertEvent)



## 云原生智能网关(cngw) 版本：2023-04-18

### 第 8 次发布

发布时间：2026-09-11 01:37:47

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CheckCloudNativeAPIGatewayMCPRouteMatch](https://cloud.tencent.com/document/api/1826/137944)
* [CheckCloudNativeAPIGatewayMCPToolVersionExist](https://cloud.tencent.com/document/api/1826/137943)
* [CompareCloudNativeAPIGatewayMCPToolVersion](https://cloud.tencent.com/document/api/1826/137942)
* [CreateCloudNativeAPIGatewayAIServiceSource](https://cloud.tencent.com/document/api/1826/137930)
* [CreateCloudNativeAPIGatewayMCPRoute](https://cloud.tencent.com/document/api/1826/137941)
* [DeleteCloudNativeAPIGatewayAIServiceSource](https://cloud.tencent.com/document/api/1826/137929)
* [DeleteCloudNativeAPIGatewayMCPRoute](https://cloud.tencent.com/document/api/1826/137940)
* [DeleteCloudNativeAPIGatewayMCPToolVersion](https://cloud.tencent.com/document/api/1826/137939)
* [DescribeCloudNativeAPIGatewayAIQuota](https://cloud.tencent.com/document/api/1826/137925)
* [DescribeCloudNativeAPIGatewayAIQuotaList](https://cloud.tencent.com/document/api/1826/137924)
* [DescribeCloudNativeAPIGatewayAIServiceSourceList](https://cloud.tencent.com/document/api/1826/137928)
* [DescribeCloudNativeAPIGatewayMCPRouteList](https://cloud.tencent.com/document/api/1826/137938)
* [DescribeCloudNativeAPIGatewayMCPToolImportTask](https://cloud.tencent.com/document/api/1826/137937)
* [DescribeCloudNativeAPIGatewayMCPToolVersion](https://cloud.tencent.com/document/api/1826/137936)
* [DescribeCloudNativeAPIGatewayMCPToolVersionList](https://cloud.tencent.com/document/api/1826/137935)
* [DescribeCloudNativeAPIGatewaySecretKeyList](https://cloud.tencent.com/document/api/1826/137931)
* [ModifyCloudNativeAPIGatewayAIServiceSource](https://cloud.tencent.com/document/api/1826/137927)
* [ModifyCloudNativeAPIGatewayMCPRoute](https://cloud.tencent.com/document/api/1826/137934)
* [ModifyCloudNativeAPIGatewayMCPRouteStatus](https://cloud.tencent.com/document/api/1826/137933)
* [RollbackCloudNativeAPIGatewayMCPToolVersion](https://cloud.tencent.com/document/api/1826/137932)

新增数据结构：

* [AIGWChangeSummary](https://cloud.tencent.com/document/api/1826/133161#AIGWChangeSummary)
* [AIGWCreateMCPRouteResult](https://cloud.tencent.com/document/api/1826/133161#AIGWCreateMCPRouteResult)
* [AIGWHeaderRule](https://cloud.tencent.com/document/api/1826/133161#AIGWHeaderRule)
* [AIGWMCPRoute](https://cloud.tencent.com/document/api/1826/133161#AIGWMCPRoute)
* [AIGWMCPRouteCheckResult](https://cloud.tencent.com/document/api/1826/133161#AIGWMCPRouteCheckResult)
* [AIGWMCPRouteListResult](https://cloud.tencent.com/document/api/1826/133161#AIGWMCPRouteListResult)
* [AIGWMCPToolVersion](https://cloud.tencent.com/document/api/1826/133161#AIGWMCPToolVersion)
* [AIGWMCPToolVersionList](https://cloud.tencent.com/document/api/1826/133161#AIGWMCPToolVersionList)
* [AIGWQuota](https://cloud.tencent.com/document/api/1826/133161#AIGWQuota)
* [AIGWQuotaDetail](https://cloud.tencent.com/document/api/1826/133161#AIGWQuotaDetail)
* [AIGWQuotaList](https://cloud.tencent.com/document/api/1826/133161#AIGWQuotaList)
* [CNAPIGwAIServiceSource](https://cloud.tencent.com/document/api/1826/133161#CNAPIGwAIServiceSource)
* [CNAPIGwAIServiceSourceAuth](https://cloud.tencent.com/document/api/1826/133161#CNAPIGwAIServiceSourceAuth)
* [CNAPIGwAIServiceSourceInfo](https://cloud.tencent.com/document/api/1826/133161#CNAPIGwAIServiceSourceInfo)
* [CNAPIGwAIServiceSourceList](https://cloud.tencent.com/document/api/1826/133161#CNAPIGwAIServiceSourceList)
* [CNAPIGwMCPToolImportResult](https://cloud.tencent.com/document/api/1826/133161#CNAPIGwMCPToolImportResult)
* [CNAPIGwMCPToolImportTaskResult](https://cloud.tencent.com/document/api/1826/133161#CNAPIGwMCPToolImportTaskResult)
* [CNAPIGwSecretKeyList](https://cloud.tencent.com/document/api/1826/133161#CNAPIGwSecretKeyList)



## 云数据库 KeeWiDB(keewidb) 版本：2022-03-08

### 第 12 次发布

发布时间：2026-09-11 02:23:09

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ChangeInstanceMaster](https://cloud.tencent.com/document/api/1520/86229)

	* 新增入参：GroupId

	* <font color="#dd0000">**修改入参**：</font>NodeId

* [CreateInstances](https://cloud.tencent.com/document/api/1520/86207)

	* 新增入参：NodeSet

* [UpgradeInstance](https://cloud.tencent.com/document/api/1520/86190)

	* 新增入参：ReplicasNum, NodeSet




## 媒体处理(mps) 版本：2019-06-12

### 第 247 次发布

发布时间：2026-09-11 02:33:49

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ProcessMedia](https://cloud.tencent.com/document/api/862/37578)

	* 新增入参：AIDubbingTask


新增数据结构：

* [AIDubbingTaskInput](https://cloud.tencent.com/document/api/862/37615#AIDubbingTaskInput)
* [AIDubbingTaskOutput](https://cloud.tencent.com/document/api/862/37615#AIDubbingTaskOutput)
* [AIDubbingTaskResult](https://cloud.tencent.com/document/api/862/37615#AIDubbingTaskResult)
* [DstSubtitleInput](https://cloud.tencent.com/document/api/862/37615#DstSubtitleInput)
* [DubbingConfig](https://cloud.tencent.com/document/api/862/37615#DubbingConfig)
* [DubbingEmbedSubtitleConfig](https://cloud.tencent.com/document/api/862/37615#DubbingEmbedSubtitleConfig)
* [DubbingOutputConfig](https://cloud.tencent.com/document/api/862/37615#DubbingOutputConfig)
* [DubbingSubtitleConfig](https://cloud.tencent.com/document/api/862/37615#DubbingSubtitleConfig)
* [DubbingTranslateConfig](https://cloud.tencent.com/document/api/862/37615#DubbingTranslateConfig)
* [OverrideAIDubbingParameter](https://cloud.tencent.com/document/api/862/37615#OverrideAIDubbingParameter)
* [RawAIDubbingParameter](https://cloud.tencent.com/document/api/862/37615#RawAIDubbingParameter)

修改数据结构：

* [ActivityPara](https://cloud.tencent.com/document/api/862/37615#ActivityPara)

	* 新增成员：AIDubbingTask

* [ActivityResItem](https://cloud.tencent.com/document/api/862/37615#ActivityResItem)

	* 新增成员：AIDubbingTask

* [WorkflowTask](https://cloud.tencent.com/document/api/862/37615#WorkflowTask)

	* 新增成员：AiDubbingTaskResult




## 流计算 Oceanus(oceanus) 版本：2019-04-22

### 第 94 次发布

发布时间：2026-09-11 02:38:04

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [JobV1](https://cloud.tencent.com/document/api/849/52010#JobV1)

	* 新增成员：StartupPoint, IsEvent, IsAlarm




## 云开发 CloudBase(tcb) 版本：2018-06-08

### 第 166 次发布

发布时间：2026-09-11 02:52:24

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateFunction](https://cloud.tencent.com/document/api/876/137951)
* [DeleteFunction](https://cloud.tencent.com/document/api/876/137950)
* [DownloadFunction](https://cloud.tencent.com/document/api/876/137949)
* [GetFunction](https://cloud.tencent.com/document/api/876/137948)
* [ListFunctions](https://cloud.tencent.com/document/api/876/137947)
* [UpdateFunctionCode](https://cloud.tencent.com/document/api/876/137952)
* [UpdateFunctionConfiguration](https://cloud.tencent.com/document/api/876/137946)

新增数据结构：

* [AgentRuntimeCodeImageConfig](https://cloud.tencent.com/document/api/876/34822#AgentRuntimeCodeImageConfig)
* [CodeReq](https://cloud.tencent.com/document/api/876/34822#CodeReq)
* [Function](https://cloud.tencent.com/document/api/876/34822#Function)
* [FunctionEipConfig](https://cloud.tencent.com/document/api/876/34822#FunctionEipConfig)
* [FunctionEipConfigFixed](https://cloud.tencent.com/document/api/876/34822#FunctionEipConfigFixed)
* [FunctionEnvironment](https://cloud.tencent.com/document/api/876/34822#FunctionEnvironment)
* [FunctionLayer](https://cloud.tencent.com/document/api/876/34822#FunctionLayer)
* [FunctionPublicNetConfig](https://cloud.tencent.com/document/api/876/34822#FunctionPublicNetConfig)
* [FunctionTrigger](https://cloud.tencent.com/document/api/876/34822#FunctionTrigger)
* [FunctionVpcConfig](https://cloud.tencent.com/document/api/876/34822#FunctionVpcConfig)
* [PrivateConfig](https://cloud.tencent.com/document/api/876/34822#PrivateConfig)
* [StatusReason](https://cloud.tencent.com/document/api/876/34822#StatusReason)



## 边缘安全加速平台(teo) 版本：2022-09-01

### 第 160 次发布

发布时间：2026-09-11 03:01:17

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DownloadL4Logs](https://cloud.tencent.com/document/api/1552/80636)

	* <font color="#dd0000">**修改入参**：</font>ZoneIds

* [DownloadL7Logs](https://cloud.tencent.com/document/api/1552/80635)

	* <font color="#dd0000">**修改入参**：</font>ZoneIds




## 边缘安全加速平台(teo) 版本：2022-01-06



## 容器服务(tke) 版本：2022-05-01

### 第 31 次发布

发布时间：2026-09-11 03:08:42

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [CreateNativeNodePoolParam](https://cloud.tencent.com/document/api/457/103206#CreateNativeNodePoolParam)

	* 新增成员：CustomImage

* [NativeNodePoolInfo](https://cloud.tencent.com/document/api/457/103206#NativeNodePoolInfo)

	* 新增成员：CustomImage

* [UpdateNativeNodePoolParam](https://cloud.tencent.com/document/api/457/103206#UpdateNativeNodePoolParam)

	* 新增成员：CustomImage




## 容器服务(tke) 版本：2018-05-25



## TokenHub(tokenhub) 版本：2026-03-22

### 第 23 次发布

发布时间：2026-09-11 03:09:09

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [UsageSeries](https://cloud.tencent.com/document/api/1823/132279#UsageSeries)

	* 新增成员：RequestCount, RequestFailCount

* [UsageStats](https://cloud.tencent.com/document/api/1823/132279#UsageStats)

	* 新增成员：RequestCount, RequestFailCount




## 云点播(vod) 版本：2024-07-18



## 云点播(vod) 版本：2018-07-17

### 第 287 次发布

发布时间：2026-09-11 03:17:29

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [AiCutOutConfig](https://cloud.tencent.com/document/api/266/31773#AiCutOutConfig)
* [PatternConfig](https://cloud.tencent.com/document/api/266/31773#PatternConfig)

修改数据结构：

* [ProcessImageAsyncTask](https://cloud.tencent.com/document/api/266/31773#ProcessImageAsyncTask)

	* 新增成员：AiCutOutConfig




