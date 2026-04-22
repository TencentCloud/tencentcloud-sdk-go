# Release v1.3.85

## 腾讯云数据仓库TCHouse-C(cdwch) 版本：2020-09-15

### 第 41 次发布

发布时间：2026-04-23 01:22:13

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [InstanceConfigInfo](https://cloud.tencent.com/document/api/1299/83429#InstanceConfigInfo)

	* 新增成员：ConfigEffective




## 弹性 MapReduce(emr) 版本：2019-01-03

### 第 138 次发布

发布时间：2026-04-23 01:52:30

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [InstallSoftware](https://cloud.tencent.com/document/api/589/130892)

修改接口：

* [CreateCloudInstance](https://cloud.tencent.com/document/api/589/113701)

	* 新增入参：ContainerExtraConf

* [TerminateInstance](https://cloud.tencent.com/document/api/589/34260)

	* 新增入参：RetainTkeCluster


新增数据结构：

* [ComponentDeployInfo](https://cloud.tencent.com/document/api/589/33981#ComponentDeployInfo)
* [ContainerExtraConf](https://cloud.tencent.com/document/api/589/33981#ContainerExtraConf)
* [LabelSelector](https://cloud.tencent.com/document/api/589/33981#LabelSelector)
* [LabelSelectorRequirement](https://cloud.tencent.com/document/api/589/33981#LabelSelectorRequirement)
* [PodAffinitySpec](https://cloud.tencent.com/document/api/589/33981#PodAffinitySpec)
* [PodAffinityTerm](https://cloud.tencent.com/document/api/589/33981#PodAffinityTerm)
* [ServiceDeployInfo](https://cloud.tencent.com/document/api/589/33981#ServiceDeployInfo)
* [StringMap](https://cloud.tencent.com/document/api/589/33981#StringMap)
* [TopologySpreadConstraint](https://cloud.tencent.com/document/api/589/33981#TopologySpreadConstraint)
* [WeightedPodAffinityTerm](https://cloud.tencent.com/document/api/589/33981#WeightedPodAffinityTerm)

修改数据结构：

* [CloudResource](https://cloud.tencent.com/document/api/589/33981#CloudResource)

	* 新增成员：PodAffinity, PodAntiAffinity, TopologySpreadConstraints, PodLabels, EnableDefaultRayCluster

* [ComputeResourceAdvanceParams](https://cloud.tencent.com/document/api/589/33981#ComputeResourceAdvanceParams)

	* 新增成员：TkeClusterNodePool




## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 290 次发布

发布时间：2026-04-23 01:54:51

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [OperateFlowRemarks](https://cloud.tencent.com/document/api/1323/130893)

修改接口：

* [DescribeFlowInfo](https://cloud.tencent.com/document/api/1323/80032)

	* 新增出参：FlowGroupRemarks


新增数据结构：

* [FlowRemarkItem](https://cloud.tencent.com/document/api/1323/70369#FlowRemarkItem)

修改数据结构：

* [FlowDetailInfo](https://cloud.tencent.com/document/api/1323/70369#FlowDetailInfo)

	* 新增成员：FlowRemarks




## 图片内容安全(ims) 版本：2020-12-29

### 第 12 次发布

发布时间：2026-04-23 02:03:23

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateImageModerationAsyncTask](https://cloud.tencent.com/document/api/1125/87021)

	* 新增入参：FileUrlList, TextContent




## 图片内容安全(ims) 版本：2020-07-13



## 腾讯云可观测平台(monitor) 版本：2023-06-16



## 腾讯云可观测平台(monitor) 版本：2018-07-24

### 第 153 次发布

发布时间：2026-04-23 02:24:11

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateAlarmNotice](https://cloud.tencent.com/document/api/248/51288)

	* 新增入参：TimeZoneName

* [ModifyAlarmNotice](https://cloud.tencent.com/document/api/248/51277)

	* 新增入参：TimeZoneName


修改数据结构：

* [AlarmNotice](https://cloud.tencent.com/document/api/248/30354#AlarmNotice)

	* 新增成员：TimeZoneName




## 云数据库 PostgreSQL(postgres) 版本：2017-03-12

### 第 64 次发布

发布时间：2026-04-23 02:41:49

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeDBErrlogs](https://cloud.tencent.com/document/api/409/18093)

	* 新增入参：LogFilters


新增数据结构：

* [LogFilter](https://cloud.tencent.com/document/api/409/16778#LogFilter)

修改数据结构：

* [ErrLogDetail](https://cloud.tencent.com/document/api/409/16778#ErrLogDetail)

	* 新增成员：ProcessId, ClientAddr, SessionId, SessionStartTime, VirtualTransactionId, SqlStateCode, ApplicationName

* [RawSlowQuery](https://cloud.tencent.com/document/api/409/16778#RawSlowQuery)

	* 新增成员：ProcessId, SessionId, VirtualTransactionId, SqlStateCode, ApplicationName




## Web 应用防火墙(waf) 版本：2018-01-25

### 第 147 次发布

发布时间：2026-04-23 03:57:54

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [AddBatchCustomRule](https://cloud.tencent.com/document/api/627/130908)
* [AddBatchCustomWhiteRule](https://cloud.tencent.com/document/api/627/130907)
* [CreateProtectGroup](https://cloud.tencent.com/document/api/627/130906)
* [DeleteBatchCustomRule](https://cloud.tencent.com/document/api/627/130905)
* [DeleteBatchCustomWhiteRule](https://cloud.tencent.com/document/api/627/130904)
* [DeleteProtectGroup](https://cloud.tencent.com/document/api/627/130903)
* [DeleteProtectGroupDomain](https://cloud.tencent.com/document/api/627/130902)
* [DescribeBatchCustomRuleList](https://cloud.tencent.com/document/api/627/130901)
* [DescribeBatchCustomWhiteRules](https://cloud.tencent.com/document/api/627/130900)
* [DescribeProtectGroup](https://cloud.tencent.com/document/api/627/130899)
* [ModifyBatchCustomRule](https://cloud.tencent.com/document/api/627/130898)
* [ModifyBatchCustomRuleStatus](https://cloud.tencent.com/document/api/627/130897)
* [ModifyBatchCustomWhiteRule](https://cloud.tencent.com/document/api/627/130896)
* [ModifyBatchCustomWhiteRuleStatus](https://cloud.tencent.com/document/api/627/130895)
* [ModifyProtectGroup](https://cloud.tencent.com/document/api/627/130894)

新增数据结构：

* [BatchCustomRuleListData](https://cloud.tencent.com/document/api/627/53609#BatchCustomRuleListData)
* [BatchCustomRuleListItem](https://cloud.tencent.com/document/api/627/53609#BatchCustomRuleListItem)
* [BatchCustomWhiteRule](https://cloud.tencent.com/document/api/627/53609#BatchCustomWhiteRule)
* [InstanceBriefInfo](https://cloud.tencent.com/document/api/627/53609#InstanceBriefInfo)
* [ProtectGroupDomainInfo](https://cloud.tencent.com/document/api/627/53609#ProtectGroupDomainInfo)
* [ProtectGroupInfo](https://cloud.tencent.com/document/api/627/53609#ProtectGroupInfo)



