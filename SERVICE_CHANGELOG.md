# Release v1.3.89

## T-Sec-DDoS防护(Anti-DDoS)(antiddos) 版本：2025-09-03

### 第 1 次发布

发布时间：2026-04-28 11:48:16

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeDDoSBlockRecords](https://cloud.tencent.com/document/api/297/131105)
* [UnblockResources](https://cloud.tencent.com/document/api/297/131104)

新增数据结构：

* [DDoSBlockRecord](https://cloud.tencent.com/document/api/297/131106#DDoSBlockRecord)
* [DDoSUnblockQuota](https://cloud.tencent.com/document/api/297/131106#DDoSUnblockQuota)
* [Filter](https://cloud.tencent.com/document/api/297/131106#Filter)



## T-Sec-DDoS防护(Anti-DDoS)(antiddos) 版本：2020-03-09



## AI Agent 安全网关(apis) 版本：2024-08-01

### 第 5 次发布

发布时间：2026-04-29 01:10:59

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateModelService](https://cloud.tencent.com/document/api/1627/129615)

	* 新增入参：TargetSelect, FindHostKeyMethod, HostKeyHeaderName, FallbackStatus, FallbackModels

* [ModifyModelService](https://cloud.tencent.com/document/api/1627/129611)

	* 新增入参：TargetSelect, FindHostKeyMethod, HostKeyHeaderName, FallbackStatus, FallbackModels


修改数据结构：

* [DescribeModelServiceResponseVO](https://cloud.tencent.com/document/api/1627/129635#DescribeModelServiceResponseVO)

	* 新增成员：TargetSelect, FindHostKeyMethod, HostKeyHeaderName, FallbackStatus, FallbackModels




## 云防火墙(cfw) 版本：2019-09-04

### 第 98 次发布

发布时间：2026-04-29 01:24:48

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [AssociatedInstanceInfo](https://cloud.tencent.com/document/api/1132/49071#AssociatedInstanceInfo)

	* 新增成员：TkeClusterId




## 云服务器(cvm) 版本：2017-03-12

### 第 162 次发布

发布时间：2026-04-29 01:34:47

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [ModifyChcNetworkMode](https://cloud.tencent.com/document/api/213/131129)



## TDSQL-C MySQL 版(cynosdb) 版本：2019-01-07

### 第 164 次发布

发布时间：2026-04-29 01:40:47

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeLibraDBInstanceDetail](https://cloud.tencent.com/document/api/1003/128085)

	* 新增出参：AnalysisUpgradeVersionInfo


新增数据结构：

* [UpgradeAnalysisInstanceVersionInfo](https://cloud.tencent.com/document/api/1003/48097#UpgradeAnalysisInstanceVersionInfo)

修改数据结构：

* [LibraDBClusterDetail](https://cloud.tencent.com/document/api/1003/48097#LibraDBClusterDetail)

	* 新增成员：AnalysisUpgradeVersionInfo




## 腾讯云数据分析智能体(dataagent) 版本：2025-05-13

### 第 15 次发布

发布时间：2026-04-29 01:43:05

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyKnowledgeBase](https://cloud.tencent.com/document/api/1800/125004)

	* 新增入参：Config


修改数据结构：

* [FileInfo](https://cloud.tencent.com/document/api/1800/125016#FileInfo)

	* 新增成员：Capabilities

* [KnowledgeBase](https://cloud.tencent.com/document/api/1800/125016#KnowledgeBase)

	* 新增成员：Config

* [KnowledgeTaskConfig](https://cloud.tencent.com/document/api/1800/125016#KnowledgeTaskConfig)

	* 新增成员：EnableImageUnderstanding

* [Scene](https://cloud.tencent.com/document/api/1800/125016#Scene)

	* 新增成员：Knowledge




## 文字识别(ocr) 版本：2018-11-19

### 第 244 次发布

发布时间：2026-04-29 02:36:04

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除接口**：</font>

* CropEnhanceImageOCR
* EraseHandwrittenImageOCR

修改接口：

* [GeneralAccurateOCR](https://cloud.tencent.com/document/api/866/34937)

	* 新增入参：WordsType




## 集团账号管理(organization) 版本：2021-03-31

### 第 57 次发布

发布时间：2026-04-29 02:39:50

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [GetIPWhitelist](https://cloud.tencent.com/document/api/850/131130)



## 集团账号管理(organization) 版本：2018-12-25



## 容器镜像服务(tcr) 版本：2019-09-24

### 第 78 次发布

发布时间：2026-04-29 03:05:38

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateInstance](https://cloud.tencent.com/document/api/1141/41572)

	* 新增入参：EnableCosVersioning


修改数据结构：

* [Registry](https://cloud.tencent.com/document/api/1141/41603#Registry)

	* 新增成员：EnableCosMAZ, EnableCosVersioning




## 容器安全服务(tcss) 版本：2020-11-01

### 第 93 次发布

发布时间：2026-04-29 03:07:33

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeAssetImageDetail](https://cloud.tencent.com/document/api/1285/65506)

	* 新增出参：RepoDigests




## TI-ONE 训练平台(tione) 版本：2021-11-11

### 第 114 次发布

发布时间：2026-04-29 03:25:29

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeAnnotatedTaskList](https://cloud.tencent.com/document/api/851/131132)
* [DescribeWorkspaces](https://cloud.tencent.com/document/api/851/131131)

修改接口：

* [CreateTrainingTask](https://cloud.tencent.com/document/api/851/117377)

	* 新增入参：TiProjectId

* [DeleteTrainingTask](https://cloud.tencent.com/document/api/851/117376)

	* 新增入参：TiProjectId

* [PushTrainingMetrics](https://cloud.tencent.com/document/api/851/75086)

	* 新增入参：TiProjectId

* [StartTrainingTask](https://cloud.tencent.com/document/api/851/117375)

	* 新增入参：TiProjectId

* [StopTrainingTask](https://cloud.tencent.com/document/api/851/117374)

	* 新增入参：TiProjectId


新增数据结构：

* [AnnotationTaskInfo](https://cloud.tencent.com/document/api/851/75051#AnnotationTaskInfo)
* [CamTag](https://cloud.tencent.com/document/api/851/75051#CamTag)
* [LabelValue](https://cloud.tencent.com/document/api/851/75051#LabelValue)
* [ResourceGroupInWorkspace](https://cloud.tencent.com/document/api/851/75051#ResourceGroupInWorkspace)
* [Workspace](https://cloud.tencent.com/document/api/851/75051#Workspace)



## TI-ONE 训练平台(tione) 版本：2019-10-22



## 云点播(vod) 版本：2024-07-18



## 云点播(vod) 版本：2018-07-17

### 第 251 次发布

发布时间：2026-04-29 03:47:32

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyEventConfig](https://cloud.tencent.com/document/api/266/55244)

	* 新增入参：SignKey




