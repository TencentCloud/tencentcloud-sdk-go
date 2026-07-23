# Release v1.3.142

## AI Agent 安全网关(apis) 版本：2024-08-01

### 第 11 次发布

发布时间：2026-07-24 01:12:00

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateMcpServer](https://cloud.tencent.com/document/api/1627/129634)

	* 新增入参：IgnoreHealthCheck

* [ModifyMcpServer](https://cloud.tencent.com/document/api/1627/129630)

	* 新增入参：IgnoreHealthCheck


修改数据结构：

* [DescribeMcpServerResponseVO](https://cloud.tencent.com/document/api/1627/129635#DescribeMcpServerResponseVO)

	* 新增成员：IgnoreHealthCheck




## 云防火墙(cfw) 版本：2019-09-04

### 第 110 次发布

发布时间：2026-07-24 01:25:40

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CheckClusterNatFwPreAccess](https://cloud.tencent.com/document/api/1132/134978)
* [CheckClusterVpcFwPreAccess](https://cloud.tencent.com/document/api/1132/134977)
* [DescribeBlockList](https://cloud.tencent.com/document/api/1132/134984)
* [DescribeFwGroupIdNames](https://cloud.tencent.com/document/api/1132/134983)
* [DescribeIpsRuleListNew](https://cloud.tencent.com/document/api/1132/134982)
* [DescribeNatRuleScopes](https://cloud.tencent.com/document/api/1132/134981)
* [DescribeSecurityGroupRegionList](https://cloud.tencent.com/document/api/1132/134979)
* [DescribeVpcAclEdgeRange](https://cloud.tencent.com/document/api/1132/134980)

新增数据结构：

* [BlockInfo](https://cloud.tencent.com/document/api/1132/49071#BlockInfo)
* [ClusterFwPreAccessCheckItem](https://cloud.tencent.com/document/api/1132/49071#ClusterFwPreAccessCheckItem)
* [EdgeRange](https://cloud.tencent.com/document/api/1132/49071#EdgeRange)
* [FwGroupIdName](https://cloud.tencent.com/document/api/1132/49071#FwGroupIdName)
* [IpsRuleDetailNew](https://cloud.tencent.com/document/api/1132/49071#IpsRuleDetailNew)
* [RuleScopeInfo](https://cloud.tencent.com/document/api/1132/49071#RuleScopeInfo)
* [SecurityGroupRegion](https://cloud.tencent.com/document/api/1132/49071#SecurityGroupRegion)



## 负载均衡(clb) 版本：2018-03-17

### 第 155 次发布

发布时间：2026-07-24 01:29:42

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [UserGroupInfo](https://cloud.tencent.com/document/api/214/30694#UserGroupInfo)

	* 新增成员：PromptId, PromptVersion, PromptName




## 日志服务(cls) 版本：2020-10-16

### 第 169 次发布

发布时间：2026-07-24 01:32:15

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateTopic](https://cloud.tencent.com/document/api/614/56456)

	* 新增入参：CustomKmsInfo

* [ModifyTopic](https://cloud.tencent.com/document/api/614/56453)

	* 新增入参：CustomKmsInfo


新增数据结构：

* [CustomKmsInfo](https://cloud.tencent.com/document/api/614/56471#CustomKmsInfo)

修改数据结构：

* [TopicInfo](https://cloud.tencent.com/document/api/614/56471#TopicInfo)

	* 新增成员：CustomKmsInfo




## 数据加速器 GooseFS(goosefs) 版本：2022-05-19

### 第 20 次发布

发布时间：2026-07-24 02:12:17

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [ModifyDataRepositoryTaskStatus](https://cloud.tencent.com/document/api/1424/134985)



## 媒体处理(mps) 版本：2019-06-12

### 第 223 次发布

发布时间：2026-07-24 02:36:35

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeVoices](https://cloud.tencent.com/document/api/862/130123)

	* 新增出参：TotalCount




## 云点播(vod) 版本：2024-07-18



## 云点播(vod) 版本：2018-07-17

### 第 275 次发布

发布时间：2026-07-24 03:24:46

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateAigcVideoRedrawTask](https://cloud.tencent.com/document/api/266/130726)

	* 新增入参：TaskInfo

* [CreateQualityInspectTemplate](https://cloud.tencent.com/document/api/266/94246)

	* 新增入参：Configs, Strategy

* [ModifyQualityInspectTemplate](https://cloud.tencent.com/document/api/266/94243)

	* 新增入参：Configs, Strategy


新增数据结构：

* [QualityInspectConfig](https://cloud.tencent.com/document/api/266/31773#QualityInspectConfig)
* [QualityInspectContainerDiagnoseResultItem](https://cloud.tencent.com/document/api/266/31773#QualityInspectContainerDiagnoseResultItem)
* [QualityInspectLLMDetectionIssue](https://cloud.tencent.com/document/api/266/31773#QualityInspectLLMDetectionIssue)
* [QualityInspectLLMDetectionReport](https://cloud.tencent.com/document/api/266/31773#QualityInspectLLMDetectionReport)
* [QualityInspectLLMDetectionResultItem](https://cloud.tencent.com/document/api/266/31773#QualityInspectLLMDetectionResultItem)
* [QualityInspectStrategy](https://cloud.tencent.com/document/api/266/31773#QualityInspectStrategy)
* [QualityInspectTimeSpotCheck](https://cloud.tencent.com/document/api/266/31773#QualityInspectTimeSpotCheck)

修改数据结构：

* [QualityInspectTaskOutput](https://cloud.tencent.com/document/api/266/31773#QualityInspectTaskOutput)

	* 新增成员：QualityEvaluationMeanOpinionScore, AestheticEvaluationScore, ContainerDiagnoseResultSet, LLMDetectionReport

* [QualityInspectTemplateItem](https://cloud.tencent.com/document/api/266/31773#QualityInspectTemplateItem)

	* 新增成员：Configs, Strategy




