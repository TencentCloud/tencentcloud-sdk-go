# Release v1.0.1200

## 应用性能监控(apm) 版本：2021-06-22

### 第 41 次发布

发布时间：2025-07-04 01:05:53

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyApmInstance](https://cloud.tencent.com/document/api/1463/89002)

	* 新增入参：IsDeleteAnyFileAnalysis, IsReadAnyFileAnalysis, IsUploadAnyFileAnalysis, IsIncludeAnyFileAnalysis, IsDirectoryTraversalAnalysis, IsTemplateEngineInjectionAnalysis, IsScriptEngineInjectionAnalysis, IsExpressionInjectionAnalysis, IsJNDIInjectionAnalysis, IsJNIInjectionAnalysis, IsWebshellBackdoorAnalysis, IsDeserializationAnalysis


修改数据结构：

* [ApmInstanceDetail](https://cloud.tencent.com/document/api/1463/64927#ApmInstanceDetail)

	* 新增成员：IsDeleteAnyFileAnalysis, IsReadAnyFileAnalysis, IsUploadAnyFileAnalysis, IsIncludeAnyFileAnalysis, IsDirectoryTraversalAnalysis, IsTemplateEngineInjectionAnalysis, IsScriptEngineInjectionAnalysis, IsExpressionInjectionAnalysis, IsJNDIInjectionAnalysis, IsJNIInjectionAnalysis, IsWebshellBackdoorAnalysis, IsDeserializationAnalysis




## 云联络中心(ccc) 版本：2020-02-10

### 第 101 次发布

发布时间：2025-07-04 01:11:58

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [ControlAIConversation](https://cloud.tencent.com/document/api/679/120723)

新增数据结构：

* [ServerPushText](https://cloud.tencent.com/document/api/679/47715#ServerPushText)



## 云安全一体化平台(csip) 版本：2022-11-21

### 第 56 次发布

发布时间：2025-07-04 01:23:34

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeAccessKeyAsset](https://cloud.tencent.com/document/api/664/120727)
* [DescribeCallRecord](https://cloud.tencent.com/document/api/664/120726)
* [DescribeSourceIPAsset](https://cloud.tencent.com/document/api/664/120725)

新增数据结构：

* [AKInfo](https://cloud.tencent.com/document/api/664/90825#AKInfo)
* [AccessKeyAlarmInfo](https://cloud.tencent.com/document/api/664/90825#AccessKeyAlarmInfo)
* [AccessKeyAsset](https://cloud.tencent.com/document/api/664/90825#AccessKeyAsset)
* [CallRecord](https://cloud.tencent.com/document/api/664/90825#CallRecord)
* [SourceIPAsset](https://cloud.tencent.com/document/api/664/90825#SourceIPAsset)



## 弹性 MapReduce(emr) 版本：2019-01-03

### 第 114 次发布

发布时间：2025-07-04 01:40:58

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [PriceResource](https://cloud.tencent.com/document/api/589/33981#PriceResource)

	* 新增成员：GpuDesc

* [Resource](https://cloud.tencent.com/document/api/589/33981#Resource)

	* 新增成员：GpuDesc




## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 233 次发布

发布时间：2025-07-04 01:42:54

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [FlowBrief](https://cloud.tencent.com/document/api/1323/70369#FlowBrief)

	* 新增成员：UserFlowType, TemplateId

* [FlowDetailInfo](https://cloud.tencent.com/document/api/1323/70369#FlowDetailInfo)

	* 新增成员：UserFlowType, TemplateId




## 腾讯电子签（基础版）(essbasic) 版本：2021-05-26

### 第 225 次发布

发布时间：2025-07-04 01:44:09

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [FlowDetailInfo](https://cloud.tencent.com/document/api/1420/61525#FlowDetailInfo)

	* 新增成员：UserFlowType, TemplateId




## 腾讯电子签（基础版）(essbasic) 版本：2020-12-22



## 云游戏(gs) 版本：2019-11-18

### 第 47 次发布

发布时间：2025-07-04 01:47:22

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [AndroidInstanceLabel](https://cloud.tencent.com/document/api/1162/40743#AndroidInstanceLabel)

	* <font color="#dd0000">**修改成员**：</font>Value




## 流计算 Oceanus(oceanus) 版本：2019-04-22

### 第 70 次发布

发布时间：2025-07-04 02:08:45

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateJobConfig](https://cloud.tencent.com/document/api/849/52004)

	* 新增入参：UseOldSystemConnector, ProgramArgsAfterGzip, CheckpointTimeoutSecond

* [ModifyJob](https://cloud.tencent.com/document/api/849/75553)

	* 新增入参：ContinueAlarm


新增数据结构：

* [HadoopYarnItem](https://cloud.tencent.com/document/api/849/52010#HadoopYarnItem)

修改数据结构：

* [Cluster](https://cloud.tencent.com/document/api/849/52010#Cluster)

	* 新增成员：Yarns

* [JobConfig](https://cloud.tencent.com/document/api/849/52010#JobConfig)

	* 新增成员：CheckpointTimeoutSecond

* [JobV1](https://cloud.tencent.com/document/api/849/52010#JobV1)

	* 新增成员：ContinueAlarm




## 云数据库Redis(redis) 版本：2018-04-12

### 第 88 次发布

发布时间：2025-07-04 02:14:00

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyAutoBackupConfig](https://cloud.tencent.com/document/api/239/20016)

	* 新增入参：BackupStorageDays




## 实时音视频(trtc) 版本：2019-07-22

### 第 114 次发布

发布时间：2025-07-04 02:36:36

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [STTConfig](https://cloud.tencent.com/document/api/647/44055#STTConfig)

	* 新增成员：VadLevel

* [ServerPushText](https://cloud.tencent.com/document/api/647/44055#ServerPushText)

	* 新增成员：AddHistory

* [TranscriptionParams](https://cloud.tencent.com/document/api/647/44055#TranscriptionParams)

	* 新增成员：VoicePrint




## Web 应用防火墙(waf) 版本：2018-01-25

### 第 122 次发布

发布时间：2025-07-04 02:47:20

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [NetworkConfig](https://cloud.tencent.com/document/api/627/53609#NetworkConfig)

修改数据结构：

* [InstanceInfo](https://cloud.tencent.com/document/api/627/53609#InstanceInfo)

	* 新增成员：NetworkConfig




