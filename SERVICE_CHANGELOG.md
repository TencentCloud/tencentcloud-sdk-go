# Release v1.3.212

## Agent 沙箱服务(ags) 版本：2025-09-20

### 第 34 次发布

发布时间：2026-11-10 01:05:04

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeEventCenterEvents](https://cloud.tencent.com/document/api/1814/140711)

新增数据结构：

* [EventItem](https://cloud.tencent.com/document/api/1814/124823#EventItem)



## 应用性能监控(apm) 版本：2021-06-22

### 第 74 次发布

发布时间：2026-11-10 01:08:53

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [ApmApplicationConfigView](https://cloud.tencent.com/document/api/1463/64927#ApmApplicationConfigView)

	* 新增成员：IsJNDIInjectionAnalysis, IsJNIInjectionAnalysis

* [CVMMeta](https://cloud.tencent.com/document/api/1463/64927#CVMMeta)

	* 新增成员：InstanceId

* [ChartData](https://cloud.tencent.com/document/api/1463/64927#ChartData)

* [ServiceDetail](https://cloud.tencent.com/document/api/1463/64927#ServiceDetail)

	* 新增成员：AppId, CreateUin

* [TopologyNode](https://cloud.tencent.com/document/api/1463/64927#TopologyNode)

	* 新增成员：ServiceID




## 灾备中心(bdrc) 版本：2026-03-30

### 第 6 次发布

发布时间：2026-11-10 01:10:45

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [BackupGroup](https://cloud.tencent.com/document/api/1824/136824#BackupGroup)

	* 新增成员：StartTime, EndTime

* [BackupInfo](https://cloud.tencent.com/document/api/1824/136824#BackupInfo)

	* 新增成员：BackupCompletedSize, BackupCompletedSizeFormatted

* [RestoreTask](https://cloud.tencent.com/document/api/1824/136824#RestoreTask)

	* 新增成员：IsOriginalPath, CloudResourceId, TargetCloudResourceId




## 云数据库 MySQL(cdb) 版本：2017-03-20

### 第 237 次发布

发布时间：2026-11-10 01:17:09

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateCloneInstance](https://cloud.tencent.com/document/api/236/50424)

	* 新增入参：UseLatestMinorVersion




## 日志服务(cls) 版本：2020-10-16

### 第 186 次发布

发布时间：2026-11-10 01:28:01

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [DynamicIndex](https://cloud.tencent.com/document/api/614/56471#DynamicIndex)

	* 新增成员：Tokenizer




## 云安全中心(csip) 版本：2022-11-21

### 第 118 次发布

发布时间：2026-11-11 01:28:26

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeDspmAccessTopologySummary](https://cloud.tencent.com/document/api/664/140766)

### 第 117 次发布

发布时间：2026-11-10 01:32:23

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateAssetBaselineRiskExportJob](https://cloud.tencent.com/document/api/664/140717)
* [DeleteBaselineIgnorePolicy](https://cloud.tencent.com/document/api/664/140716)
* [DescribeAssetBaselineRiskList](https://cloud.tencent.com/document/api/664/140715)
* [DescribeBaselineIgnorePolicyList](https://cloud.tencent.com/document/api/664/140714)
* [DescribeClusterBaselineRiskPolicyList](https://cloud.tencent.com/document/api/664/140713)
* [DescribeHostVulRiskEventList](https://cloud.tencent.com/document/api/664/140721)
* [DescribeHostVulScanStatus](https://cloud.tencent.com/document/api/664/140720)
* [ModifyBaselineIgnorePolicy](https://cloud.tencent.com/document/api/664/140712)
* [ModifyClusterAuditLogStatus](https://cloud.tencent.com/document/api/664/140718)

新增数据结构：

* [AssetBaselineRisk](https://cloud.tencent.com/document/api/664/90825#AssetBaselineRisk)
* [BaselineIgnorePolicy](https://cloud.tencent.com/document/api/664/90825#BaselineIgnorePolicy)
* [ClusterBaselinePolicy](https://cloud.tencent.com/document/api/664/90825#ClusterBaselinePolicy)
* [HostVulRiskEventItem](https://cloud.tencent.com/document/api/664/90825#HostVulRiskEventItem)
* [HostVulScanStatus](https://cloud.tencent.com/document/api/664/90825#HostVulScanStatus)



## 云服务器(cvm) 版本：2017-03-12

### 第 175 次发布

发布时间：2026-11-11 01:38:07

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ConfigureChcAssistVpc](https://cloud.tencent.com/document/api/213/76417)

	* 新增入参：SkipBmcCheck




## 大数据智能体工作台DataBuddy(databuddy) 版本：2026-07-15

### 第 15 次发布

发布时间：2026-11-11 01:44:58

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [GetSessionStatus](https://cloud.tencent.com/document/api/1835/140769)
* [ReadProductJson](https://cloud.tencent.com/document/api/1835/140768)

新增数据结构：

* [GetSessionStatusRsp](https://cloud.tencent.com/document/api/1835/138006#GetSessionStatusRsp)
* [ReadProductJsonRsp](https://cloud.tencent.com/document/api/1835/138006#ReadProductJsonRsp)
* [SessionStatusItem](https://cloud.tencent.com/document/api/1835/138006#SessionStatusItem)
* [SessionStatusQueryItem](https://cloud.tencent.com/document/api/1835/138006#SessionStatusQueryItem)



## 数据库智能管家 DBbrain(dbbrain) 版本：2021-05-27

### 第 70 次发布

发布时间：2026-11-11 01:46:10

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeSqlExplainV2](https://cloud.tencent.com/document/api/1130/140770)

新增数据结构：

* [SqlExplainFieldDescV2](https://cloud.tencent.com/document/api/1130/57812#SqlExplainFieldDescV2)
* [SqlExplainResultV2](https://cloud.tencent.com/document/api/1130/57812#SqlExplainResultV2)
* [SqlExplainRowV2](https://cloud.tencent.com/document/api/1130/57812#SqlExplainRowV2)
* [SqlExplainTableRefV2](https://cloud.tencent.com/document/api/1130/57812#SqlExplainTableRefV2)



## 数据库智能管家 DBbrain(dbbrain) 版本：2019-10-16



## TDSQL MySQL 版(dcdb) 版本：2018-04-11

### 第 84 次发布

发布时间：2026-11-10 01:54:07

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [UpgradeDedicatedDCDBInstance](https://cloud.tencent.com/document/api/557/92616)

	* 新增入参：MasterHostId, SlaveHostIds




## Elasticsearch Service(es) 版本：2025-01-01



## Elasticsearch Service(es) 版本：2018-04-16

### 第 113 次发布

发布时间：2026-11-10 02:05:06

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DeleteIndex](https://cloud.tencent.com/document/api/845/74383)

	* 新增入参：IndexNames

	* 新增出参：TotalCount, SuccessCount, FailedCount, Results

* [DescribeIndexList](https://cloud.tencent.com/document/api/845/74382)

	* 新增入参：FilterIndexNames

* [UpdateInstance](https://cloud.tencent.com/document/api/845/30629)

	* 新增入参：WaitGreen


新增数据结构：

* [BatchIndexResult](https://cloud.tencent.com/document/api/845/30634#BatchIndexResult)

修改数据结构：

* [EventDataInfoOverview](https://cloud.tencent.com/document/api/845/30634#EventDataInfoOverview)

	* 新增成员：ExpandDetail

* [IndexMetaField](https://cloud.tencent.com/document/api/845/30634#IndexMetaField)

	* 新增成员：FieldCount, FieldLimit, FieldUsageRatio, OccupiedShards, IndexOpenStatus, FieldLimitHardCap, FieldLimitMin, FieldAutoAdjustRemaining

* [IndexSettingsField](https://cloud.tencent.com/document/api/845/30634#IndexSettingsField)

	* 新增成员：Dynamic, TotalFieldsLimit




## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 329 次发布

发布时间：2026-11-11 01:58:38

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [MiniAppCreateApproverInfo](https://cloud.tencent.com/document/api/1323/70369#MiniAppCreateApproverInfo)

	* 新增成员：PinCode




## 云直播CSS(live) 版本：2018-08-01

### 第 194 次发布

发布时间：2026-11-10 02:24:55

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateAvatarEditingLive](https://cloud.tencent.com/document/api/267/140725)
* [DescribeAvatarEditingLive](https://cloud.tencent.com/document/api/267/140724)
* [StopAvatarEditingLive](https://cloud.tencent.com/document/api/267/140723)
* [SwitchAvatarEditingScene](https://cloud.tencent.com/document/api/267/140722)



## 云数据库 MariaDB(mariadb) 版本：2017-03-12

### 第 77 次发布

发布时间：2026-11-10 02:27:32

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [UpgradeDedicatedDBInstance](https://cloud.tencent.com/document/api/237/92618)

	* 新增入参：MasterHostId, SlaveHostIds




## 媒体处理(mps) 版本：2019-06-12

### 第 259 次发布

发布时间：2026-11-10 02:32:03

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [AudioSRConfig](https://cloud.tencent.com/document/api/862/37615#AudioSRConfig)
* [SpatialAudioConfig](https://cloud.tencent.com/document/api/862/37615#SpatialAudioConfig)

修改数据结构：

* [AudioEnhanceConfig](https://cloud.tencent.com/document/api/862/37615#AudioEnhanceConfig)

	* 新增成员：AudioSR, SpatialAudio

* [MediaVideoStreamItem](https://cloud.tencent.com/document/api/862/37615#MediaVideoStreamItem)

	* 新增成员：Duration


### 第 258 次发布

发布时间：2026-11-09 16:54:49

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateModelChatTask](https://cloud.tencent.com/document/api/862/140685)
* [DescribeModelChatTask](https://cloud.tencent.com/document/api/862/140684)

新增数据结构：

* [ModelChatCompletionChoice](https://cloud.tencent.com/document/api/862/37615#ModelChatCompletionChoice)
* [ModelChatCompletionUsage](https://cloud.tencent.com/document/api/862/37615#ModelChatCompletionUsage)
* [ModelChatContentPart](https://cloud.tencent.com/document/api/862/37615#ModelChatContentPart)
* [ModelChatDataInputInfo](https://cloud.tencent.com/document/api/862/37615#ModelChatDataInputInfo)
* [ModelChatInputAudio](https://cloud.tencent.com/document/api/862/37615#ModelChatInputAudio)
* [ModelChatInputFile](https://cloud.tencent.com/document/api/862/37615#ModelChatInputFile)
* [ModelChatInputImage](https://cloud.tencent.com/document/api/862/37615#ModelChatInputImage)
* [ModelChatInputVideo](https://cloud.tencent.com/document/api/862/37615#ModelChatInputVideo)
* [ModelChatMediaInputInfo](https://cloud.tencent.com/document/api/862/37615#ModelChatMediaInputInfo)
* [ModelChatMessageItem](https://cloud.tencent.com/document/api/862/37615#ModelChatMessageItem)
* [ModelChatResponseMessage](https://cloud.tencent.com/document/api/862/37615#ModelChatResponseMessage)
* [ModelChatTaskNotifyConfig](https://cloud.tencent.com/document/api/862/37615#ModelChatTaskNotifyConfig)
* [ModelChatUrlInputInfo](https://cloud.tencent.com/document/api/862/37615#ModelChatUrlInputInfo)



## 腾讯健康组学平台(omics) 版本：2022-11-28

### 第 32 次发布

发布时间：2026-11-10 02:37:25

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [HPCCluster](https://cloud.tencent.com/document/api/1643/89100#HPCCluster)

	* 新增成员：StatusReason, IsolateTime, Message

* [HPCNode](https://cloud.tencent.com/document/api/1643/89100#HPCNode)

	* 新增成员：IsExistingNode




## 云数据库Redis(redis) 版本：2018-04-12

### 第 117 次发布

发布时间：2026-11-10 17:43:28

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeModules](https://cloud.tencent.com/document/api/239/140754)

新增数据结构：

* [Module](https://cloud.tencent.com/document/api/239/20022#Module)

### 第 116 次发布

发布时间：2026-11-10 02:41:44

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [ModifyModule](https://cloud.tencent.com/document/api/239/140726)

### 第 115 次发布

发布时间：2026-11-09 15:08:36

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [ModifyValkeyModule](https://cloud.tencent.com/document/api/239/140680)

### 第 114 次发布

发布时间：2026-11-09 11:33:30

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [ModifySSLEncryptionMode](https://cloud.tencent.com/document/api/239/140670)
* [UpdateSSLCert](https://cloud.tencent.com/document/api/239/140669)



## 边缘安全加速平台(teo) 版本：2022-09-01

### 第 173 次发布

发布时间：2026-11-11 02:44:31

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [Plan](https://cloud.tencent.com/document/api/1552/80721#Plan)

	* 新增成员：ClientAuthTrafficCapacity




## 边缘安全加速平台(teo) 版本：2022-01-06



## 容器服务(tke) 版本：2022-05-01



## 容器服务(tke) 版本：2018-05-25

### 第 241 次发布

发布时间：2026-11-11 02:48:28

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [RotateClusterToken](https://cloud.tencent.com/document/api/457/133506)

	* 新增入参：ClusterId, ClusterType, Operation, ClusterToken, Force




## 实时音视频(trtc) 版本：2019-07-22

### 第 156 次发布

发布时间：2026-11-10 03:08:04

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [ModifyCloudTranscription](https://cloud.tencent.com/document/api/647/140727)

新增数据结构：

* [TranscriptionSubscribeParam](https://cloud.tencent.com/document/api/647/44055#TranscriptionSubscribeParam)



