# Release v1.3.77

## 媒体处理(mps) 版本：2019-06-12

### 第 191 次发布

发布时间：2026-04-10 21:47:50

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [BatchProcessMedia](https://cloud.tencent.com/document/api/862/118508)

	* 新增入参：TaskMode

* [DescribeImageTaskDetail](https://cloud.tencent.com/document/api/862/118509)

	* 新增出参：Definition, ImageTask, InputInfo




## 文字识别(ocr) 版本：2018-11-19

### 第 243 次发布

发布时间：2026-04-13 02:32:37

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [SubmitQuestionMarkAgentJob](https://cloud.tencent.com/document/api/866/128273)

	* 新增入参：ImageBase64List, ImageUrlList


修改数据结构：

* [MarkInfo](https://cloud.tencent.com/document/api/866/33527#MarkInfo)

	* 新增成员：QuestionPositions, QuestionImagePositions, RightAnswer




## 腾讯健康组学平台(omics) 版本：2022-11-28

### 第 28 次发布

发布时间：2026-04-13 02:35:19

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribePublicApplications](https://cloud.tencent.com/document/api/1643/130431)

新增数据结构：

* [PublicApplication](https://cloud.tencent.com/document/api/1643/89100#PublicApplication)
* [ToolRepoTag](https://cloud.tencent.com/document/api/1643/89100#ToolRepoTag)



## 前端性能监控(rum) 版本：2021-06-22

### 第 50 次发布

发布时间：2026-04-13 02:45:29

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeApplicationExitReportDetail](https://cloud.tencent.com/document/api/1464/130450)
* [DescribeApplicationExitReportList](https://cloud.tencent.com/document/api/1464/130449)
* [DescribeExceptionDetail](https://cloud.tencent.com/document/api/1464/130448)
* [DescribeExceptionReportList](https://cloud.tencent.com/document/api/1464/130447)
* [DescribeFOOMMallocProblemDetail](https://cloud.tencent.com/document/api/1464/130446)
* [DescribeFOOMMallocProblemList](https://cloud.tencent.com/document/api/1464/130445)
* [DescribeFOOMMallocReportList](https://cloud.tencent.com/document/api/1464/130444)
* [DescribeFOOMProblemDetail](https://cloud.tencent.com/document/api/1464/130443)
* [DescribeFOOMProblemList](https://cloud.tencent.com/document/api/1464/130442)
* [DescribeFOOMReportList](https://cloud.tencent.com/document/api/1464/130441)
* [DescribeIssuesDistribution](https://cloud.tencent.com/document/api/1464/130440)
* [DescribeIssuesList](https://cloud.tencent.com/document/api/1464/130439)
* [DescribeIssuesStatisticsTrend](https://cloud.tencent.com/document/api/1464/130438)
* [DescribeLagANRProblemAccountDetail](https://cloud.tencent.com/document/api/1464/130437)
* [DescribeLagANRProblemFeatureAccounts](https://cloud.tencent.com/document/api/1464/130436)
* [DescribeLagANRProblemList](https://cloud.tencent.com/document/api/1464/130435)
* [DescribeToken](https://cloud.tencent.com/document/api/1464/130434)
* [DescribeTopIssues](https://cloud.tencent.com/document/api/1464/130433)

新增数据结构：

* [CompareCondition](https://cloud.tencent.com/document/api/1464/61476#CompareCondition)
* [Filters](https://cloud.tencent.com/document/api/1464/61476#Filters)



## TI-ONE 训练平台(tione) 版本：2021-11-11

### 第 110 次发布

发布时间：2026-04-13 03:21:26

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateTrainingTask](https://cloud.tencent.com/document/api/851/117377)

	* 新增入参：Envs


修改数据结构：

* [TrainingTaskSetItem](https://cloud.tencent.com/document/api/851/75051#TrainingTaskSetItem)

	* 新增成员：Envs




## TI-ONE 训练平台(tione) 版本：2019-10-22



## TSF-应用管理&Consul(tsf) 版本：2018-03-26

### 第 139 次发布

发布时间：2026-04-13 03:36:12

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyLaneRule](https://cloud.tencent.com/document/api/649/44502)

	* <font color="#dd0000">**修改入参**：</font>RuleTagRelationship

* [UpdateApiRateLimitRules](https://cloud.tencent.com/document/api/649/50623)

	* 新增入参：UsePathAndMethodFormat


修改数据结构：

* [ApiDetailInfo](https://cloud.tencent.com/document/api/649/36099#ApiDetailInfo)

	* 新增成员：PathMappingUnsupported, PathMappingUnsupportedMsg

* [ApiRateLimitRule](https://cloud.tencent.com/document/api/649/36099#ApiRateLimitRule)

	* 新增成员：UsePathAndMethodFormat

* [Microservice](https://cloud.tencent.com/document/api/649/36099#Microservice)

	* 新增成员：DeleteDisabled, DeleteDisabledReason

* [ServiceSetting](https://cloud.tencent.com/document/api/649/36099#ServiceSetting)

	* 新增成员：EnableGlobalService

* [SimpleGroup](https://cloud.tencent.com/document/api/649/36099#SimpleGroup)

	* 新增成员：CreateTime, UpdatedTime




## 音视频终端引擎(vcube) 版本：2022-04-10

### 第 8 次发布

发布时间：2026-04-13 03:40:59

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeLicenseList](https://cloud.tencent.com/document/api/1449/113509)

	* 新增入参：PageNumber, PageSize, Platform




