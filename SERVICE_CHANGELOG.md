# Release v1.3.83

## AI Agent 安全网关(apis) 版本：2024-08-01

### 第 4 次发布

发布时间：2026-04-21 01:10:59

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateModelService](https://cloud.tencent.com/document/api/1627/129615)

	* 新增入参：SensitiveDataCheckStatus, SensitiveDataCheckConfig

* [ModifyModelService](https://cloud.tencent.com/document/api/1627/129611)

	* 新增入参：SensitiveDataCheckStatus, SensitiveDataCheckConfig


新增数据结构：

* [SensitiveDataCheckConfigDTO](https://cloud.tencent.com/document/api/1627/129635#SensitiveDataCheckConfigDTO)

修改数据结构：

* [DescribeModelServiceResponseVO](https://cloud.tencent.com/document/api/1627/129635#DescribeModelServiceResponseVO)

	* 新增成员：SensitiveDataCheckStatus, SensitiveDataCheckConfig




## 消息队列 CKafka 版(ckafka) 版本：2019-08-19

### 第 144 次发布

发布时间：2026-04-21 01:26:34

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeAccessPolicy](https://cloud.tencent.com/document/api/597/130721)
* [ModifyAccessPolicy](https://cloud.tencent.com/document/api/597/130720)

修改接口：

* [CreateConnectResource](https://cloud.tencent.com/document/api/597/77737)

	* 新增入参：Tags

* [ModifyDatahubTask](https://cloud.tencent.com/document/api/597/77793)

	* 新增入参：TasksMax, SyncThrottleLimit, AutoExpandFlag


新增数据结构：

* [ExternalAccessInfoWrapper](https://cloud.tencent.com/document/api/597/40861#ExternalAccessInfoWrapper)

修改数据结构：

* [DatahubTaskInfo](https://cloud.tencent.com/document/api/597/40861#DatahubTaskInfo)

	* 新增成员：TaskMax, SyncThrottleLimit, AutoExpandFlag

* [DescribeConnectResource](https://cloud.tencent.com/document/api/597/40861#DescribeConnectResource)

	* 新增成员：Tags

* [DescribeConnectResourceResp](https://cloud.tencent.com/document/api/597/40861#DescribeConnectResourceResp)

	* 新增成员：Tags

* [DescribeDatahubTaskRes](https://cloud.tencent.com/document/api/597/40861#DescribeDatahubTaskRes)

	* 新增成员：TaskMax, SyncThrottleLimit, AutoExpandFlag

* [KafkaParam](https://cloud.tencent.com/document/api/597/40861#KafkaParam)

	* 新增成员：TopicList




## 负载均衡(clb) 版本：2018-03-17

### 第 148 次发布

发布时间：2026-04-21 01:27:51

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyLoadBalancerSla](https://cloud.tencent.com/document/api/214/63892)

	* 新增出参：DealName




## 日志服务(cls) 版本：2020-10-16

### 第 156 次发布

发布时间：2026-04-21 01:29:26

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeKafkaConsumer](https://cloud.tencent.com/document/api/614/95719)

	* 新增出参：HasServicesLog, ScopeType

* [ModifyKafkaConsumer](https://cloud.tencent.com/document/api/614/95720)

	* 新增入参：HasServicesLog, ScopeType

* [OpenKafkaConsumer](https://cloud.tencent.com/document/api/614/72339)

	* 新增入参：HasServicesLog, ScopeType




## 多媒体创作引擎(cme) 版本：2019-10-29

### 第 64 次发布

发布时间：2026-04-21 01:31:15

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DeleteAccount](https://cloud.tencent.com/document/api/1156/130723)
* [ForbidAccount](https://cloud.tencent.com/document/api/1156/130722)



## Elasticsearch Service(es) 版本：2025-01-01



## Elasticsearch Service(es) 版本：2018-04-16

### 第 98 次发布

发布时间：2026-04-21 01:52:59

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [UpdateIndex](https://cloud.tencent.com/document/api/845/74380)

	* 新增入参：MountIndex, IndexUuid, BackingIndexName


修改数据结构：

* [BackingIndexMetaField](https://cloud.tencent.com/document/api/845/30634#BackingIndexMetaField)

	* 新增成员：IndexUuid

* [IndexOptionsField](https://cloud.tencent.com/document/api/845/30634#IndexOptionsField)

	* 新增成员：FullOffloadedEnable, FullOffloadedMaxAge, FullOffloadedRetrieveMaxAge




## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 289 次发布

发布时间：2026-04-21 01:54:01

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateFlowReminds](https://cloud.tencent.com/document/api/1323/86002)

	* 新增入参：RemindTypes, RemindEmailInfos


新增数据结构：

* [RemindEmailInfo](https://cloud.tencent.com/document/api/1323/70369#RemindEmailInfo)

修改数据结构：

* [ApproverInfo](https://cloud.tencent.com/document/api/1323/70369#ApproverInfo)

	* 新增成员：ApproverEmail




## 媒体处理(mps) 版本：2019-06-12

### 第 196 次发布

发布时间：2026-04-21 02:26:34

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateStreamPackageLinearAssemblyProgram](https://cloud.tencent.com/document/api/862/130315)

	* 新增入参：VodAcquisitionMethod

* [ModifyStreamPackageLinearAssemblyProgram](https://cloud.tencent.com/document/api/862/130294)

	* 新增入参：VodAcquisitionMethod


修改数据结构：

* [LinearAssemblyProgramInfo](https://cloud.tencent.com/document/api/862/37615#LinearAssemblyProgramInfo)

	* 新增成员：VodAcquisitionMethod




## 消息队列 MQTT 版(mqtt) 版本：2024-05-16

### 第 27 次发布

发布时间：2026-04-21 02:30:57

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeSharedSubscriptionGroups](https://cloud.tencent.com/document/api/1778/130725)

新增数据结构：

* [SharedGroup](https://cloud.tencent.com/document/api/1778/111031#SharedGroup)



## 腾讯云数据库 AI 服务(tdai) 版本：2025-07-17

### 第 11 次发布

发布时间：2026-04-21 03:11:20

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateAgentInstance](https://cloud.tencent.com/document/api/1813/123274)

	* 新增入参：InstanceType, TemplateId, Skills, SoulId, Description

* [DescribeAgentInstances](https://cloud.tencent.com/document/api/1813/123272)

	* 新增入参：InstanceType

	* 新增出参：StatusCounts


新增数据结构：

* [ClawConfigInfo](https://cloud.tencent.com/document/api/1813/123239#ClawConfigInfo)
* [ClawDeployInfo](https://cloud.tencent.com/document/api/1813/123239#ClawDeployInfo)
* [StatusItem](https://cloud.tencent.com/document/api/1813/123239#StatusItem)

修改数据结构：

* [AgentInstance](https://cloud.tencent.com/document/api/1813/123239#AgentInstance)

	* 新增成员：DeployPlace, PolicyIds, ClawConfig, InstanceType, AllowedActions, LastActiveTime, Description




## 边缘安全加速平台(teo) 版本：2022-09-01

### 第 144 次发布

发布时间：2026-04-21 03:17:27

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [ShieldParameters](https://cloud.tencent.com/document/api/1552/80721#ShieldParameters)

修改数据结构：

* [RuleEngineAction](https://cloud.tencent.com/document/api/1552/80721#RuleEngineAction)

	* 新增成员：ShieldParameters




## 边缘安全加速平台(teo) 版本：2022-01-06



## 音视频终端引擎(vcube) 版本：2022-04-10

### 第 9 次发布

发布时间：2026-04-21 03:43:43

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateApplicationAndBindLicense](https://cloud.tencent.com/document/api/1449/113533)

	* 新增入参：BundleName

* [CreateTrialApplicationAndLicense](https://cloud.tencent.com/document/api/1449/113512)

	* 新增入参：BundleName




## 云点播(vod) 版本：2024-07-18



## 云点播(vod) 版本：2018-07-17

### 第 249 次发布

发布时间：2026-04-21 03:45:32

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateAigcVideoRedrawTask](https://cloud.tencent.com/document/api/266/130726)

新增数据结构：

* [AigcVideoRedrawOutputConfig](https://cloud.tencent.com/document/api/266/31773#AigcVideoRedrawOutputConfig)
* [AigcVideoRedrawTaskInputFileInfo](https://cloud.tencent.com/document/api/266/31773#AigcVideoRedrawTaskInputFileInfo)



## 私有网络(vpc) 版本：2017-03-12

### 第 303 次发布

发布时间：2026-04-21 03:50:36

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [VpnGateway](https://cloud.tencent.com/document/api/215/15824#VpnGateway)

	* 新增成员：TagSet




## 数据开发治理平台 WeData(wedata) 版本：2025-08-06



## 数据开发治理平台 WeData(wedata) 版本：2021-08-20

### 第 188 次发布

发布时间：2026-04-21 04:02:46

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [SystemRole](https://cloud.tencent.com/document/api/1267/76336#SystemRole)

	* 新增成员：CreateTimeStr, Creator




