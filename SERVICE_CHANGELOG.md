# Release v1.3.113

## 运维安全中心（堡垒机）(bh) 版本：2023-04-18

### 第 32 次发布

发布时间：2026-06-10 01:14:19

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [AppAsset](https://cloud.tencent.com/document/api/1025/74416#AppAsset)

	* 新增成员：AccountCount, AgentInputType, AgentInputSubmit, UserNameType, UserNameValue, PasswordType, PasswordValue, SubmitType, SubmitValue




## 费用中心(billing) 版本：2018-07-09

### 第 89 次发布

发布时间：2026-06-09 01:19:55

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeOrgMemberAccountBalance](https://cloud.tencent.com/document/api/555/132645)

新增数据结构：

* [DescribeOrgMemberAccountBalanceData](https://cloud.tencent.com/document/api/555/19183#DescribeOrgMemberAccountBalanceData)



## 云硬盘(cbs) 版本：2017-03-12

### 第 75 次发布

发布时间：2026-06-10 01:18:58

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [AttachRemoteDisks](https://cloud.tencent.com/document/api/362/132688)
* [CreateRemoteDisks](https://cloud.tencent.com/document/api/362/132687)
* [DescribeRemoteDiskConfigQuota](https://cloud.tencent.com/document/api/362/132686)
* [DescribeRemoteDisks](https://cloud.tencent.com/document/api/362/132685)
* [DescribeRemoteDisksDeniedActions](https://cloud.tencent.com/document/api/362/132684)
* [DetachRemoteDisks](https://cloud.tencent.com/document/api/362/132683)
* [InquirePriceCreateRemoteDisks](https://cloud.tencent.com/document/api/362/132682)
* [InquirePriceRenewRemoteDisks](https://cloud.tencent.com/document/api/362/132681)
* [ModifyRemoteDiskAttributes](https://cloud.tencent.com/document/api/362/132680)
* [RenewRemoteDisk](https://cloud.tencent.com/document/api/362/132679)
* [SwitchParameterCreateRemoteDisks](https://cloud.tencent.com/document/api/362/132678)
* [SwitchParameterRenewRemoteDisks](https://cloud.tencent.com/document/api/362/132677)
* [TerminateRemoteDisks](https://cloud.tencent.com/document/api/362/132676)

### 第 74 次发布

发布时间：2026-06-09 01:22:53

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [Snapshot](https://cloud.tencent.com/document/api/362/15669#Snapshot)

	* 新增成员：IsLocked, LatestModifyTime, AutoSnapshotPolicyId




## 云防火墙(cfw) 版本：2019-09-04

### 第 103 次发布

发布时间：2026-06-11 01:11:59

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DeleteBlockIgnoreRuleNew](https://cloud.tencent.com/document/api/1132/104037)

	* <font color="#dd0000">**修改入参**：</font>ShowType

* [ExpandCfwVertical](https://cloud.tencent.com/document/api/1132/54258)

	* 新增入参：ElasticTrafficSwitch


修改数据结构：

* [NatInstanceInfo](https://cloud.tencent.com/document/api/1132/49071#NatInstanceInfo)

	* 新增成员：ElasticTrafficSwitch

* [SerialRegionInfo](https://cloud.tencent.com/document/api/1132/49071#SerialRegionInfo)

	* 新增成员：ElasticTrafficSwitch

* [VpcFwInstanceInfo](https://cloud.tencent.com/document/api/1132/49071#VpcFwInstanceInfo)

	* 新增成员：ElasticTrafficSwitch




## 负载均衡(clb) 版本：2018-03-17

### 第 150 次发布

发布时间：2026-06-10 01:28:51

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateLoadBalancer](https://cloud.tencent.com/document/api/214/30692)

	* 新增入参：AvailableZoneAffinityInfo




## 云安全一体化平台(csip) 版本：2022-11-21

### 第 82 次发布

发布时间：2026-06-11 01:15:55

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateIaCAccessToken](https://cloud.tencent.com/document/api/664/132736)
* [CreateIaCFileExportJob](https://cloud.tencent.com/document/api/664/132735)
* [CreateIaCFileReScanTask](https://cloud.tencent.com/document/api/664/132734)
* [DeleteIaCAccessToken](https://cloud.tencent.com/document/api/664/132733)
* [DeleteIaCFile](https://cloud.tencent.com/document/api/664/132732)
* [DescribeIaCFileList](https://cloud.tencent.com/document/api/664/132731)
* [DescribeIaCFileOverview](https://cloud.tencent.com/document/api/664/132730)
* [DescribeIaCFileReport](https://cloud.tencent.com/document/api/664/132729)
* [DescribeIaCTokenList](https://cloud.tencent.com/document/api/664/132728)
* [ModifyIaCTokenPeriod](https://cloud.tencent.com/document/api/664/132727)

新增数据结构：

* [CICDToken](https://cloud.tencent.com/document/api/664/90825#CICDToken)
* [IaCFile](https://cloud.tencent.com/document/api/664/90825#IaCFile)
* [IaCFileRisk](https://cloud.tencent.com/document/api/664/90825#IaCFileRisk)
* [KeyValueInt](https://cloud.tencent.com/document/api/664/90825#KeyValueInt)

### 第 81 次发布

发布时间：2026-06-09 01:37:26

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [AccessKeyAlarm](https://cloud.tencent.com/document/api/664/90825#AccessKeyAlarm)

	* 新增成员：AIFailedReason

* [SubUserInfo](https://cloud.tencent.com/document/api/664/90825#SubUserInfo)

	* 新增成员：CreateTime




## 云服务器(cvm) 版本：2017-03-12

### 第 163 次发布

发布时间：2026-06-11 01:17:26

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [Instance](https://cloud.tencent.com/document/api/213/15753#Instance)

	* 新增成员：CpuTopology




## 主机安全(cwp) 版本：2018-02-28

### 第 164 次发布

发布时间：2026-06-10 01:37:50

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ExportTasks](https://cloud.tencent.com/document/api/296/52508)

	* 新增出参：FileName




## TDSQL-C MySQL 版(cynosdb) 版本：2019-01-07

### 第 174 次发布

发布时间：2026-06-11 01:20:16

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeServerlessInstanceSpecs](https://cloud.tencent.com/document/api/1003/113699)

	* 新增入参：ClusterLevel


### 第 173 次发布

发布时间：2026-06-10 01:42:24

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateProxyEndPoint](https://cloud.tencent.com/document/api/1003/94133)

	* 新增入参：LoadBalanceMode

* [ModifyProxyRwSplit](https://cloud.tencent.com/document/api/1003/94129)

	* 新增入参：LoadBalanceMode


修改数据结构：

* [ProxyGroupRwInfo](https://cloud.tencent.com/document/api/1003/48097#ProxyGroupRwInfo)

	* 新增成员：LoadBalanceMode




## 腾讯云数据分析智能体(dataagent) 版本：2025-05-13

### 第 17 次发布

发布时间：2026-06-11 01:21:21

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyChunk](https://cloud.tencent.com/document/api/1800/125008)

	* 新增入参：Summary

* [QueryChunkList](https://cloud.tencent.com/document/api/1800/125007)

	* 新增出参：AutoTotal, ManualTotal


修改数据结构：

* [Chunk](https://cloud.tencent.com/document/api/1800/125016#Chunk)

	* 新增成员：ChunkSource

* [FileInfo](https://cloud.tencent.com/document/api/1800/125016#FileInfo)

	* 新增成员：UpdateTime

* [KnowledgeTaskConfig](https://cloud.tencent.com/document/api/1800/125016#KnowledgeTaskConfig)

	* 新增成员：EnableExtractDb




## 弹性 MapReduce(emr) 版本：2019-01-03

### 第 143 次发布

发布时间：2026-06-11 01:26:33

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [ComputeResourceAdvanceParams](https://cloud.tencent.com/document/api/589/33981#ComputeResourceAdvanceParams)

	* 新增成员：NodePoolJoinMode




## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 297 次发布

发布时间：2026-06-09 02:02:01

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeExtendedServiceAuthDetail](https://cloud.tencent.com/document/api/1323/102513)

	* 新增入参：PartnerOrganizationName




## 全球加速(ga2) 版本：2025-01-15

### 第 7 次发布

发布时间：2026-06-10 02:01:24

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateForwardingRule](https://cloud.tencent.com/document/api/1817/130172)

	* 新增入参：ResponseHeaders, HideResponseHeaders

* [ModifyForwardingRule](https://cloud.tencent.com/document/api/1817/130157)

	* 新增入参：ResponseHeaders, HideResponseHeaders


新增数据结构：

* [HideResponseHeaders](https://cloud.tencent.com/document/api/1817/130045#HideResponseHeaders)
* [ResponseHeaders](https://cloud.tencent.com/document/api/1817/130045#ResponseHeaders)

### 第 6 次发布

发布时间：2026-06-09 02:05:47

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateListener](https://cloud.tencent.com/document/api/1817/130171)

	* 新增入参：HttpVersion

* [ModifyEndpointGroup](https://cloud.tencent.com/document/api/1817/130158)

	* 新增入参：HttpVersion


修改数据结构：

* [EndpointGroupConfiguration](https://cloud.tencent.com/document/api/1817/130045#EndpointGroupConfiguration)

	* 新增成员：HttpVersion

* [EndpointGroupConfigurationSet](https://cloud.tencent.com/document/api/1817/130045#EndpointGroupConfigurationSet)

	* 新增成员：HttpVersion




## 物联网开发平台(iotexplorer) 版本：2019-04-23

### 第 143 次发布

发布时间：2026-06-10 02:07:34

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateDevicePublishSDPAnswer](https://cloud.tencent.com/document/api/1081/132691)
* [DeleteDeviceSDP](https://cloud.tencent.com/document/api/1081/132690)

修改接口：

* [CreateDeviceSDPAnswer](https://cloud.tencent.com/document/api/1081/127863)

	* 新增入参：RequesterTag




## 智能视图计算平台(iss) 版本：2023-05-17

### 第 32 次发布

发布时间：2026-06-10 02:12:57

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除接口**：</font>

* AddAITask
* DeleteAITask
* DescribeAITask
* DescribeAITaskResult
* ListAITasks
* UpdateAITask
* UpdateAITaskStatus

<font color="#dd0000">**删除数据结构**：</font>

* AIConfig
* AITaskInfo
* AITaskResultData
* AITaskResultInfo
* AITemplates
* BaseAIResultInfo
* BodyAIResultInfo
* CarAIResultInfo
* ChefClothAIResultInfo
* ChefHatAIResultInfo
* FaceMaskAIResultInfo
* ListAITaskData
* Location
* OperTimeSlot
* PetAIResultInfo
* PhoneCallAIResultInfo
* PlateContent
* SmokingAIResultInfo
* SnapshotConfig



## 云数据库 MongoDB(mongodb) 版本：2019-07-25

### 第 71 次发布

发布时间：2026-06-10 02:21:16

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [PromoteDBInstanceToActive](https://cloud.tencent.com/document/api/240/132692)

修改接口：

* [SetInstanceMaintenance](https://cloud.tencent.com/document/api/240/104304)

	* 新增入参：MaintenanceDays




## 云数据库 MongoDB(mongodb) 版本：2018-04-08



## 腾讯云可观测平台(monitor) 版本：2023-06-16



## 腾讯云可观测平台(monitor) 版本：2018-07-24

### 第 159 次发布

发布时间：2026-06-10 02:22:07

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateOnCallForm](https://cloud.tencent.com/document/api/248/132697)
* [DeleteOnCallForms](https://cloud.tencent.com/document/api/248/132696)
* [DescribeOnCallForm](https://cloud.tencent.com/document/api/248/132695)
* [DescribeOnCallForms](https://cloud.tencent.com/document/api/248/132694)
* [UpdateOnCallForm](https://cloud.tencent.com/document/api/248/132693)

新增数据结构：

* [CoverStaffInfo](https://cloud.tencent.com/document/api/248/30354#CoverStaffInfo)
* [OnCallForm](https://cloud.tencent.com/document/api/248/30354#OnCallForm)
* [OneOnCallForm](https://cloud.tencent.com/document/api/248/30354#OneOnCallForm)
* [StaffInfo](https://cloud.tencent.com/document/api/248/30354#StaffInfo)



## 流计算 Oceanus(oceanus) 版本：2019-04-22

### 第 85 次发布

发布时间：2026-06-09 02:32:16

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [HiveMetastoreInfo](https://cloud.tencent.com/document/api/849/52010#HiveMetastoreInfo)

修改数据结构：

* [Cluster](https://cloud.tencent.com/document/api/849/52010#Cluster)

	* 新增成员：HiveMetastore, SecurityGroupIds, NetEniType




## 文字识别(ocr) 版本：2018-11-19

### 第 250 次发布

发布时间：2026-06-11 01:41:39

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [MarkInfo](https://cloud.tencent.com/document/api/866/33527#MarkInfo)

	* 新增成员：Subject, QuestionType




## 云数据库 PostgreSQL(postgres) 版本：2017-03-12

### 第 66 次发布

发布时间：2026-06-11 01:43:04

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeZones](https://cloud.tencent.com/document/api/409/16769)

	* 新增入参：StorageType




## 前端性能监控(rum) 版本：2021-06-22

### 第 52 次发布

发布时间：2026-06-09 02:39:29

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [ScoreInfoV2](https://cloud.tencent.com/document/api/1464/61476#ScoreInfoV2)

	* 新增成员：Exclusion




## 容器镜像服务(tcr) 版本：2019-09-24

### 第 80 次发布

发布时间：2026-06-10 02:43:34

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DeleteAIModel](https://cloud.tencent.com/document/api/1141/132702)
* [DeleteSkill](https://cloud.tencent.com/document/api/1141/132708)
* [DescribeAIModelVersionDetail](https://cloud.tencent.com/document/api/1141/132701)
* [DescribeSkillDetail](https://cloud.tencent.com/document/api/1141/132707)
* [DescribeSkillDownloadInfo](https://cloud.tencent.com/document/api/1141/132706)
* [ListAIModelVersions](https://cloud.tencent.com/document/api/1141/132700)
* [ListAIModels](https://cloud.tencent.com/document/api/1141/132699)
* [ListSkillVersions](https://cloud.tencent.com/document/api/1141/132705)
* [ListSkills](https://cloud.tencent.com/document/api/1141/132704)

新增数据结构：

* [DeleteModelItem](https://cloud.tencent.com/document/api/1141/41603#DeleteModelItem)
* [ModelDetail](https://cloud.tencent.com/document/api/1141/41603#ModelDetail)
* [ModelList](https://cloud.tencent.com/document/api/1141/41603#ModelList)
* [Skill](https://cloud.tencent.com/document/api/1141/41603#Skill)
* [SkillList](https://cloud.tencent.com/document/api/1141/41603#SkillList)
* [SkillType](https://cloud.tencent.com/document/api/1141/41603#SkillType)
* [SkillVersionList](https://cloud.tencent.com/document/api/1141/41603#SkillVersionList)
* [VersionList](https://cloud.tencent.com/document/api/1141/41603#VersionList)



## 容器安全服务(tcss) 版本：2020-11-01

### 第 96 次发布

发布时间：2026-06-10 02:44:46

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateAssetImageScanSetting](https://cloud.tencent.com/document/api/1285/65591)

	* 新增入参：ClusterIDs

* [CreateAssetImageScanTask](https://cloud.tencent.com/document/api/1285/65515)

	* 新增入参：ClusterIDs

* [CreateVulScanTask](https://cloud.tencent.com/document/api/1285/81658)

	* 新增入参：ExcludeLocalImageIDs, ExcludeRegistryImageIDs, LocalClusterIDs

* [DescribeAssetImageScanSetting](https://cloud.tencent.com/document/api/1285/65487)

	* 新增出参：ClusterIDs


### 第 95 次发布

发布时间：2026-06-09 02:49:22

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeAbnormalProcessRules](https://cloud.tencent.com/document/api/1285/65548)

	* 新增出参：RuleExtSet

* [DescribeAccessControlRules](https://cloud.tencent.com/document/api/1285/65542)

	* 新增出参：RuleExtSet


新增数据结构：

* [AbnormalProcessRuleExtSetItem](https://cloud.tencent.com/document/api/1285/65614#AbnormalProcessRuleExtSetItem)
* [AccessControlRuleExtSetItem](https://cloud.tencent.com/document/api/1285/65614#AccessControlRuleExtSetItem)

修改数据结构：

* [AbnormalProcessRuleInfo](https://cloud.tencent.com/document/api/1285/65614#AbnormalProcessRuleInfo)

	* 新增成员：IsGlobal

* [AccessControlRuleInfo](https://cloud.tencent.com/document/api/1285/65614#AccessControlRuleInfo)

	* 新增成员：IsGlobal

* [AssetClusterListItem](https://cloud.tencent.com/document/api/1285/65614#AssetClusterListItem)

	* 新增成员：BindRuleID

* [K8sApiAbnormalRuleListItem](https://cloud.tencent.com/document/api/1285/65614#K8sApiAbnormalRuleListItem)

	* 新增成员：EffectAllCluster, RuleActions, RuleInfoList

* [RuleBaseInfo](https://cloud.tencent.com/document/api/1285/65614#RuleBaseInfo)

	* 新增成员：IsGlobal




## 消息队列 TDMQ(tdmq) 版本：2020-02-17

### 第 176 次发布

发布时间：2026-06-11 01:50:39

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [PulsarProClusterInfo](https://cloud.tencent.com/document/api/1179/46089#PulsarProClusterInfo)

	* 新增成员：ElasticTpsEnabled, EncryptionStatus

* [PulsarProClusterSpecInfo](https://cloud.tencent.com/document/api/1179/46089#PulsarProClusterSpecInfo)

	* 新增成员：TotalTps




## TDSQL(tdmysql) 版本：2021-11-22

### 第 7 次发布

发布时间：2026-06-09 02:55:13

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateCloneInstance](https://cloud.tencent.com/document/api/1376/132041)

	* 新增入参：SecurityGroupIds

* [CreateDBInstances](https://cloud.tencent.com/document/api/1376/132036)

	* 新增入参：SecurityGroupIds, UserName, Password, EncryptionEnable

* [DescribeDBInstanceDetail](https://cloud.tencent.com/document/api/1376/132035)

	* 新增出参：EncryptionEnable, EncryptionKmsRegion

* [DescribeDBInstances](https://cloud.tencent.com/document/api/1376/132034)

	* 新增入参：EngineType


修改数据结构：

* [InstanceInfo](https://cloud.tencent.com/document/api/1376/128305#InstanceInfo)

	* 新增成员：AnalysisInstanceInfo




## TI-ONE 训练平台(tione) 版本：2021-11-11

### 第 120 次发布

发布时间：2026-06-11 01:53:26

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateDataset](https://cloud.tencent.com/document/api/851/75050)

	* 新增入参：TiProjectId

* [DescribeLogs](https://cloud.tencent.com/document/api/851/117840)

	* 新增入参：TiProjectId




## TI-ONE 训练平台(tione) 版本：2019-10-22



## 容器服务(tke) 版本：2022-05-01



## 容器服务(tke) 版本：2018-05-25

### 第 233 次发布

发布时间：2026-06-10 02:56:02

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**预下线接口**：</font>

* DescribeAvailableTKEEdgeVersion
* DescribeEdgeAvailableExtraArgs
* DescribeTKEEdgeExternalKubeconfig
* DescribeTKEEdgeScript

### 第 232 次发布

发布时间：2026-06-09 03:00:55

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [Extenders](https://cloud.tencent.com/document/api/457/31866#Extenders)

	* 新增成员：Ignorable




## TokenHub(tokenhub) 版本：2026-03-22

### 第 5 次发布

发布时间：2026-06-09 03:04:04

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [BindingItem](https://cloud.tencent.com/document/api/1823/132279#BindingItem)

	* 新增成员：Status




## 实时音视频(trtc) 版本：2019-07-22

### 第 145 次发布

发布时间：2026-06-10 03:01:07

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [TextToSpeech](https://cloud.tencent.com/document/api/647/122475)

	* 新增入参：ExtraParams

* [TextToSpeechSSE](https://cloud.tencent.com/document/api/647/122474)

	* 新增入参：ExtraParams


修改数据结构：

* [Voice](https://cloud.tencent.com/document/api/647/44055#Voice)

	* 新增成员：Emotion


### 第 144 次发布

发布时间：2026-06-09 03:06:13

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateCloudTranscription](https://cloud.tencent.com/document/api/647/129058)

	* 新增入参：TTSParam


新增数据结构：

* [TTSParam](https://cloud.tencent.com/document/api/647/44055#TTSParam)
* [TTSVoice](https://cloud.tencent.com/document/api/647/44055#TTSVoice)
* [TermPair](https://cloud.tencent.com/document/api/647/44055#TermPair)
* [TerminologyItem](https://cloud.tencent.com/document/api/647/44055#TerminologyItem)

修改数据结构：

* [TranslationParam](https://cloud.tencent.com/document/api/647/44055#TranslationParam)

	* 新增成员：Terminologies




## 云点播(vod) 版本：2024-07-18



## 云点播(vod) 版本：2018-07-17

### 第 265 次发布

发布时间：2026-06-10 03:07:25

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeAigcFaceInfoAsync](https://cloud.tencent.com/document/api/266/132709)

### 第 264 次发布

发布时间：2026-06-09 03:12:25

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeTaskDetail](https://cloud.tencent.com/document/api/266/33431)

	* 新增出参：DescribeAigcFaceInfoAsyncTask


新增数据结构：

* [DescribeAigcFaceInfoAsyncInput](https://cloud.tencent.com/document/api/266/31773#DescribeAigcFaceInfoAsyncInput)
* [DescribeAigcFaceInfoAsyncOutput](https://cloud.tencent.com/document/api/266/31773#DescribeAigcFaceInfoAsyncOutput)
* [DescribeAigcFaceInfoAsyncTask](https://cloud.tencent.com/document/api/266/31773#DescribeAigcFaceInfoAsyncTask)

修改数据结构：

* [EventContent](https://cloud.tencent.com/document/api/266/31773#EventContent)

	* 新增成员：CreateAigcAdvancedCustomElementCompleteEvent, CreateAigcCustomVoiceCompleteEvent, DescribeAigcFaceInfoAsyncCompleteEvent




## Web 应用防火墙(waf) 版本：2018-01-25

### 第 158 次发布

发布时间：2026-06-11 02:02:28

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeApiSecEventDetail](https://cloud.tencent.com/document/api/627/132739)
* [DescribeApiSecEventList](https://cloud.tencent.com/document/api/627/132738)

新增数据结构：

* [ApiEvent](https://cloud.tencent.com/document/api/627/53609#ApiEvent)
* [ApiSecAttackSource](https://cloud.tencent.com/document/api/627/53609#ApiSecAttackSource)
* [ApiSecEventChange](https://cloud.tencent.com/document/api/627/53609#ApiSecEventChange)



