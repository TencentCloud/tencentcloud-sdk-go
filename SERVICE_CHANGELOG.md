# Release v1.0.1164

## 应用性能监控(apm) 版本：2021-06-22

### 第 36 次发布

发布时间：2025-05-14 01:09:06

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [ApmInstanceDetail](https://cloud.tencent.com/document/api/1463/64927#ApmInstanceDetail)

	* 新增成员：LogIndexType, LogTraceIdKey




## 弹性伸缩(as) 版本：2018-04-19

### 第 84 次发布

发布时间：2025-05-14 01:09:15

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeAutoScalingGroupLastActivities](https://cloud.tencent.com/document/api/377/37652)

	* 新增入参：ExcludeCancelledActivity




## 费用中心(billing) 版本：2018-07-09

### 第 75 次发布

发布时间：2025-05-14 01:11:06

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateAllocationRule](https://cloud.tencent.com/document/api/555/118451)
* [CreateAllocationUnit](https://cloud.tencent.com/document/api/555/118450)
* [CreateGatherRule](https://cloud.tencent.com/document/api/555/118449)
* [DeleteAllocationRule](https://cloud.tencent.com/document/api/555/118448)
* [DeleteAllocationUnit](https://cloud.tencent.com/document/api/555/118447)
* [DeleteGatherRule](https://cloud.tencent.com/document/api/555/118446)
* [DescribeAllocationRuleDetail](https://cloud.tencent.com/document/api/555/118445)
* [DescribeAllocationRuleSummary](https://cloud.tencent.com/document/api/555/118444)
* [DescribeAllocationTree](https://cloud.tencent.com/document/api/555/118443)
* [DescribeAllocationUnitDetail](https://cloud.tencent.com/document/api/555/118442)
* [DescribeGatherRuleDetail](https://cloud.tencent.com/document/api/555/118441)
* [ModifyAllocationRule](https://cloud.tencent.com/document/api/555/118440)
* [ModifyAllocationUnit](https://cloud.tencent.com/document/api/555/118439)
* [ModifyGatherRule](https://cloud.tencent.com/document/api/555/118438)

修改接口：

* [DescribeBillDetail](https://cloud.tencent.com/document/api/555/19182)


新增数据结构：

* [AllocationRationExpression](https://cloud.tencent.com/document/api/555/19183#AllocationRationExpression)
* [AllocationRuleExpression](https://cloud.tencent.com/document/api/555/19183#AllocationRuleExpression)
* [AllocationRuleOverview](https://cloud.tencent.com/document/api/555/19183#AllocationRuleOverview)
* [AllocationRulesSummary](https://cloud.tencent.com/document/api/555/19183#AllocationRulesSummary)
* [AllocationTree](https://cloud.tencent.com/document/api/555/19183#AllocationTree)
* [AllocationUnit](https://cloud.tencent.com/document/api/555/19183#AllocationUnit)
* [GatherRuleSummary](https://cloud.tencent.com/document/api/555/19183#GatherRuleSummary)



## 腾讯混元大模型(hunyuan) 版本：2023-09-01

### 第 41 次发布

发布时间：2025-05-14 01:36:21

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ChatCompletions](https://cloud.tencent.com/document/api/1729/105701)

	* 新增入参：WebSearchOptions


新增数据结构：

* [Approximate](https://cloud.tencent.com/document/api/1729/101838#Approximate)
* [Knowledge](https://cloud.tencent.com/document/api/1729/101838#Knowledge)
* [UserLocation](https://cloud.tencent.com/document/api/1729/101838#UserLocation)
* [WebSearchOptions](https://cloud.tencent.com/document/api/1729/101838#WebSearchOptions)



## 大模型知识引擎(lke) 版本：2023-11-30

### 第 41 次发布

发布时间：2025-05-14 01:44:28

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeAttributeLabel](https://cloud.tencent.com/document/api/1759/105043)

	* 新增入参：QueryScope

* [DescribeDoc](https://cloud.tencent.com/document/api/1759/105071)

	* 新增出参：CustomerKnowledgeId, AttributeFlags

* [ListAttributeLabel](https://cloud.tencent.com/document/api/1759/105041)

	* 新增入参：LabelSize

* [ModifyDoc](https://cloud.tencent.com/document/api/1759/105058)

	* 新增入参：CustomerKnowledgeId, AttributeFlags

* [SaveDoc](https://cloud.tencent.com/document/api/1759/105054)

	* 新增入参：CustomerKnowledgeId, AttributeFlags


修改数据结构：

* [AttrLabelDetail](https://cloud.tencent.com/document/api/1759/105104#AttrLabelDetail)

	* 新增成员：LabelTotalCount

* [ListDocItem](https://cloud.tencent.com/document/api/1759/105104#ListDocItem)

	* 新增成员：CustomerKnowledgeId, AttributeFlags

* [ModelInfo](https://cloud.tencent.com/document/api/1759/105104#ModelInfo)

	* 新增成员：IsExclusive




## 云开发 CloudBase(tcb) 版本：2018-06-08

### 第 112 次发布

发布时间：2025-05-14 01:58:00

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeEnvs](https://cloud.tencent.com/document/api/876/34820)

	* 新增入参：Limit, Offset




## TI-ONE 训练平台(tione) 版本：2021-11-11

### 第 79 次发布

发布时间：2025-05-14 02:04:39

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [ExecAction](https://cloud.tencent.com/document/api/851/75051#ExecAction)
* [NumOrPercent](https://cloud.tencent.com/document/api/851/75051#NumOrPercent)
* [RollingUpdate](https://cloud.tencent.com/document/api/851/75051#RollingUpdate)
* [TCPSocketAction](https://cloud.tencent.com/document/api/851/75051#TCPSocketAction)

修改数据结构：

* [HTTPGetAction](https://cloud.tencent.com/document/api/851/75051#HTTPGetAction)

	* 新增成员：Port

* [ProbeAction](https://cloud.tencent.com/document/api/851/75051#ProbeAction)

	* 新增成员：Exec, TCPSocket, ActionType

* [ServiceGroup](https://cloud.tencent.com/document/api/851/75051#ServiceGroup)

	* 新增成员：MonitorSource

* [ServiceInfo](https://cloud.tencent.com/document/api/851/75051#ServiceInfo)

	* 新增成员：RollingUpdate

* [SpecPrice](https://cloud.tencent.com/document/api/851/75051#SpecPrice)

	* <font color="#dd0000">**修改成员**：</font>SpecCount




## TI-ONE 训练平台(tione) 版本：2019-10-22



## 实时音视频(trtc) 版本：2019-07-22

### 第 108 次发布

发布时间：2025-05-14 02:08:03

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [RecognizeConfig](https://cloud.tencent.com/document/api/647/44055#RecognizeConfig)

	* 新增成员：HotWordList




