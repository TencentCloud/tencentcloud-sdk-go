# Release v1.3.52

## 腾讯混元生图(aiart) 版本：2022-12-29

### 第 30 次发布

发布时间：2026-03-05 01:08:05

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [SubmitTextToImageJob](https://cloud.tencent.com/document/api/1668/124632)

	* <font color="#dd0000">**删除入参**：</font>SkillType, GeneratedImageCount




## 应用性能监控(apm) 版本：2021-06-22

### 第 56 次发布

发布时间：2026-03-03 01:09:20

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyApmApplicationConfig](https://cloud.tencent.com/document/api/1463/125072)

	* 新增入参：EnableThresholdConfig, ErrRateThreshold, ResponseDurationWarningThreshold


修改数据结构：

* [ApmAppConfig](https://cloud.tencent.com/document/api/1463/64927#ApmAppConfig)

	* 新增成员：EnableThresholdConfig, ErrRateThreshold, ResponseDurationWarningThreshold

* [ApmApplicationConfigView](https://cloud.tencent.com/document/api/1463/64927#ApmApplicationConfigView)

	* 新增成员：AgentIgnoreOperation, EnableSecurityConfig, IsSqlInjectionAnalysis, IsInstrumentationVulnerabilityScan, IsRemoteCommandExecutionAnalysis, IsMemoryHijackingAnalysis, IsDeleteAnyFileAnalysis, IsReadAnyFileAnalysis, IsUploadAnyFileAnalysis, IsIncludeAnyFileAnalysis, IsDirectoryTraversalAnalysis, IsTemplateEngineInjectionAnalysis, IsScriptEngineInjectionAnalysis, IsExpressionInjectionAnalysis, IsJndiInjectionAnalysis, IsJniInjectionAnalysis, IsWebshellBackdoorAnalysis, IsDeserializationAnalysis, EnableDashboardConfig, IsRelatedDashboard, DashboardTopicID, EnableThresholdConfig, ErrRateThreshold, ResponseDurationWarningThreshold




## 语音识别(asr) 版本：2019-06-14

### 第 47 次发布

发布时间：2026-03-05 01:09:25

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [SentenceDetail](https://cloud.tencent.com/document/api/1093/37824#SentenceDetail)

	* 新增成员：LangType, SpeakerRoleName




## 云联络中心(ccc) 版本：2020-02-10

### 第 124 次发布

发布时间：2026-03-09 01:17:57

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeSessionDetail](https://cloud.tencent.com/document/api/679/123050)

	* 新增出参：SysHangupReason, SysHangupReasonString




## 云原生xa0etcd(cetcd) 版本：2022-03-25

### 第 1 次发布

发布时间：2026-03-04 15:12:45

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateEtcdInstance](https://cloud.tencent.com/document/api/1577/128924)
* [CreateEtcdSnapshot](https://cloud.tencent.com/document/api/1577/128923)
* [CreateEtcdSnapshotPolicy](https://cloud.tencent.com/document/api/1577/128922)
* [DescribeEtcdAvailableVersions](https://cloud.tencent.com/document/api/1577/128921)
* [DescribeEtcdCreatingProgress](https://cloud.tencent.com/document/api/1577/128920)
* [DescribeEtcdCredentials](https://cloud.tencent.com/document/api/1577/128919)
* [DescribeEtcdInstances](https://cloud.tencent.com/document/api/1577/128905)
* [DescribeEtcdQuota](https://cloud.tencent.com/document/api/1577/128918)
* [DescribeEtcdRegions](https://cloud.tencent.com/document/api/1577/128917)
* [DescribeEtcdSnapshotPolicies](https://cloud.tencent.com/document/api/1577/128916)
* [DescribeEtcdSnapshots](https://cloud.tencent.com/document/api/1577/128915)
* [DescribeEtcdTasks](https://cloud.tencent.com/document/api/1577/128914)
* [DescribeRPCMethodList](https://cloud.tencent.com/document/api/1577/128904)
* [DisableEtcdInstanceDeletionProtection](https://cloud.tencent.com/document/api/1577/128913)
* [EnableEtcdInstanceDeletionProtection](https://cloud.tencent.com/document/api/1577/128912)
* [ModifyEtcdAttribute](https://cloud.tencent.com/document/api/1577/128911)
* [ModifyEtcdConfiguration](https://cloud.tencent.com/document/api/1577/128910)
* [ModifyEtcdSnapshotPolicy](https://cloud.tencent.com/document/api/1577/128909)
* [ResizeEtcdInstance](https://cloud.tencent.com/document/api/1577/128908)
* [UpgradeEtcdInstance](https://cloud.tencent.com/document/api/1577/128907)

新增数据结构：

* [ChargePrepaidConfig](https://cloud.tencent.com/document/api/1577/128925#ChargePrepaidConfig)
* [Etcd](https://cloud.tencent.com/document/api/1577/128925#Etcd)
* [EtcdAdvancedSettings](https://cloud.tencent.com/document/api/1577/128925#EtcdAdvancedSettings)
* [EtcdAutoCompactionSettings](https://cloud.tencent.com/document/api/1577/128925#EtcdAutoCompactionSettings)
* [EtcdBackupSettings](https://cloud.tencent.com/document/api/1577/128925#EtcdBackupSettings)
* [EtcdCredential](https://cloud.tencent.com/document/api/1577/128925#EtcdCredential)
* [EtcdMember](https://cloud.tencent.com/document/api/1577/128925#EtcdMember)
* [EtcdMonitorSettings](https://cloud.tencent.com/document/api/1577/128925#EtcdMonitorSettings)
* [EtcdSecuritySettings](https://cloud.tencent.com/document/api/1577/128925#EtcdSecuritySettings)
* [EtcdSnapshot](https://cloud.tencent.com/document/api/1577/128925#EtcdSnapshot)
* [EtcdSnapshotPolicy](https://cloud.tencent.com/document/api/1577/128925#EtcdSnapshotPolicy)
* [EtcdTaskInfo](https://cloud.tencent.com/document/api/1577/128925#EtcdTaskInfo)
* [Filter](https://cloud.tencent.com/document/api/1577/128925#Filter)
* [InstanceConfig](https://cloud.tencent.com/document/api/1577/128925#InstanceConfig)
* [PrometheusCreationParam](https://cloud.tencent.com/document/api/1577/128925#PrometheusCreationParam)
* [RPCMethod](https://cloud.tencent.com/document/api/1577/128925#RPCMethod)
* [RegionInstance](https://cloud.tencent.com/document/api/1577/128925#RegionInstance)
* [TaskStepInfo](https://cloud.tencent.com/document/api/1577/128925#TaskStepInfo)
* [VersionInstance](https://cloud.tencent.com/document/api/1577/128925#VersionInstance)



## 云防火墙(cfw) 版本：2019-09-04

### 第 93 次发布

发布时间：2026-03-03 01:14:02

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [EnterpriseSecurityGroupRuleRuleInfo](https://cloud.tencent.com/document/api/1132/49071#EnterpriseSecurityGroupRuleRuleInfo)

	* 新增成员：RulePartition, Scope

* [SecurityGroupRule](https://cloud.tencent.com/document/api/1132/49071#SecurityGroupRule)

	* 新增成员：Scope

* [SecurityGroupSimplifyRule](https://cloud.tencent.com/document/api/1132/49071#SecurityGroupSimplifyRule)

	* 新增成员：Scope




## 消息队列 CKafka 版(ckafka) 版本：2019-08-19

### 第 141 次发布

发布时间：2026-03-02 12:03:10

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeModifyType](https://cloud.tencent.com/document/api/597/125916)

	* 新增入参：ModifyZone

* [ModifyInstanceAttributes](https://cloud.tencent.com/document/api/597/40832)

	* 新增入参：RetentionBytes, AdminSecurity, TransactionalIdExpirationMs


修改数据结构：

* [CreateInstancePostData](https://cloud.tencent.com/document/api/597/40861#CreateInstancePostData)

	* 新增成员：EventId

* [CreateInstancePreData](https://cloud.tencent.com/document/api/597/40861#CreateInstancePreData)

	* 新增成员：EventId

* [DescModifyType](https://cloud.tencent.com/document/api/597/40861#DescModifyType)

	* <font color="#dd0000">**删除成员**：</font>MigrateCostTime, UpgradeStrategy, MigrateCostTimeHighSpeed

* [InstanceAttributesResponse](https://cloud.tencent.com/document/api/597/40861#InstanceAttributesResponse)

	* 新增成员：RetentionBytes, TransactionalIdExpirationMs

* [InstanceDetail](https://cloud.tencent.com/document/api/597/40861#InstanceDetail)

	* 新增成员：RetentionBytes

* [InstanceVersion](https://cloud.tencent.com/document/api/597/40861#InstanceVersion)

	* 新增成员：AllowModifyTxnIdExpiration




## 日志服务(cls) 版本：2020-10-16

### 第 149 次发布

发布时间：2026-03-09 01:28:42

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateNetworkApplication](https://cloud.tencent.com/document/api/614/129006)
* [DeleteNetworkApplication](https://cloud.tencent.com/document/api/614/129005)
* [DescribeNetworkApplicationDetail](https://cloud.tencent.com/document/api/614/129004)
* [DescribeNetworkApplications](https://cloud.tencent.com/document/api/614/129003)
* [ModifyNetworkApplication](https://cloud.tencent.com/document/api/614/129002)

修改接口：

* [GetMetricLabelValues](https://cloud.tencent.com/document/api/614/126717)

	* 新增入参：TopicId, LabelName, Start, End, Match


新增数据结构：

* [NetworkApplicationDetail](https://cloud.tencent.com/document/api/614/56471#NetworkApplicationDetail)
* [NetworkApplicationInfo](https://cloud.tencent.com/document/api/614/56471#NetworkApplicationInfo)



## 暴露面管理服务(ctem) 版本：2023-11-28

### 第 14 次发布

发布时间：2026-03-09 01:32:19

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeJobRecordDetails](https://cloud.tencent.com/document/api/1755/120306)

	* 新增出参：EnterpriseEquityPath


新增数据结构：

* [Equity](https://cloud.tencent.com/document/api/1755/120320#Equity)



## 主机安全(cwp) 版本：2018-02-28

### 第 157 次发布

发布时间：2026-03-09 01:34:13

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribePatchEffectHostList](https://cloud.tencent.com/document/api/296/129011)
* [DescribePatchInfo](https://cloud.tencent.com/document/api/296/129010)
* [DescribeWindowsPatchList](https://cloud.tencent.com/document/api/296/129009)
* [ExportPatchEffectHostList](https://cloud.tencent.com/document/api/296/129008)
* [ExportWindowsPatchList](https://cloud.tencent.com/document/api/296/129007)

新增数据结构：

* [EventPatchInfo](https://cloud.tencent.com/document/api/296/19867#EventPatchInfo)
* [PatchEffectHostList](https://cloud.tencent.com/document/api/296/19867#PatchEffectHostList)
* [RelateVulInfo](https://cloud.tencent.com/document/api/296/19867#RelateVulInfo)



## 数据库智能管家 DBbrain(dbbrain) 版本：2021-05-27

### 第 53 次发布

发布时间：2026-03-06 01:40:09

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateMongoDBKillTask](https://cloud.tencent.com/document/api/1130/128958)



## 数据库智能管家 DBbrain(dbbrain) 版本：2019-10-16



## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 276 次发布

发布时间：2026-03-03 01:24:44

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeContractComparisonTask](https://cloud.tencent.com/document/api/1323/123953)

	* 新增入参：Filters


修改数据结构：

* [ComparisonDetail](https://cloud.tencent.com/document/api/1323/70369#ComparisonDetail)

	* 新增成员：FormatType




## 图片内容安全(ims) 版本：2020-12-29

### 第 11 次发布

发布时间：2026-03-04 01:26:55

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateImageModerationAsyncTask](https://cloud.tencent.com/document/api/1125/87021)

	* 新增入参：Type




## 图片内容安全(ims) 版本：2020-07-13



## 物联网开发平台(iotexplorer) 版本：2019-04-23

### 第 132 次发布

发布时间：2026-03-09 02:01:02

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateTWeTalkAIBot](https://cloud.tencent.com/document/api/1081/125990)

	* 新增入参：WebhookTools, BotType, RAGConfig

* [ModifyTWeTalkAIBot](https://cloud.tencent.com/document/api/1081/125986)

	* 新增入参：WebhookTools, BotType, RAGConfig


修改数据结构：

* [TalkAIBotInfo](https://cloud.tencent.com/document/api/1081/34988#TalkAIBotInfo)

	* 新增成员：WebhookTools, BotType, RAGConfig

* [TalkAgentConfigInfo](https://cloud.tencent.com/document/api/1081/34988#TalkAgentConfigInfo)

	* 新增成员：SubtitleCallbackUrl, SubtitleCallbackSignKey, SubtitleCallbackTimeout




## 云直播CSS(live) 版本：2018-08-01

### 第 165 次发布

发布时间：2026-03-09 02:09:49

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateAuditKeywordLib](https://cloud.tencent.com/document/api/267/127523)

	* 新增入参：Name, Description, Suggestion, MatchType

* [DescribeDeliverLogDownList](https://cloud.tencent.com/document/api/267/97727)

	* 新增入参：StartTime, EndTime, DeliverDomains




## 多网聚合加速(mna) 版本：2021-01-19

### 第 32 次发布

发布时间：2026-03-05 01:36:12

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [AddApplication](https://cloud.tencent.com/document/api/1385/128944)
* [DeleteApplication](https://cloud.tencent.com/document/api/1385/128943)
* [GetApplication](https://cloud.tencent.com/document/api/1385/128942)
* [UpdateApplicationInfo](https://cloud.tencent.com/document/api/1385/128941)
* [UpdateApplicationKey](https://cloud.tencent.com/document/api/1385/128940)

新增数据结构：

* [DelApplicationList](https://cloud.tencent.com/document/api/1385/55846#DelApplicationList)



## 媒体处理(mps) 版本：2019-06-12

### 第 183 次发布

发布时间：2026-03-03 01:40:33

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [EditMedia](https://cloud.tencent.com/document/api/862/43010)

	* 新增入参：ResourceId

* [ProcessLiveStream](https://cloud.tencent.com/document/api/862/39227)

	* 新增入参：ResourceId




## 文字识别(ocr) 版本：2018-11-19

### 第 236 次发布

发布时间：2026-03-03 01:45:35

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除接口**：</font>

* RecognizeTableMultiOCR



## SSL 证书(ssl) 版本：2019-12-05

### 第 97 次发布

发布时间：2026-03-06 02:48:37

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateCertificate](https://cloud.tencent.com/document/api/400/49427)

	* 新增出参：ResourceIds




## 云开发 CloudBase(tcb) 版本：2018-06-08

### 第 125 次发布

发布时间：2026-03-09 02:54:39

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [RunCommands](https://cloud.tencent.com/document/api/876/129012)

修改接口：

* [CreateEnv](https://cloud.tencent.com/document/api/876/128592)

	* 新增入参：RenewFlag


新增数据结构：

* [MgoCommandParam](https://cloud.tencent.com/document/api/876/34822#MgoCommandParam)

### 第 124 次发布

发布时间：2026-03-06 02:54:12

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DeleteAuthDomain](https://cloud.tencent.com/document/api/876/128960)
* [ModifySafeRule](https://cloud.tencent.com/document/api/876/128959)

### 第 123 次发布

发布时间：2026-03-02 12:47:38

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除接口**：</font>

* BindEnvGateway
* CommonServiceAPI
* CreateAndDeployCloudBaseProject
* CreateCloudBaseRunResource
* CreateCloudBaseRunServerVersion
* CreatePostpayPackage
* DeleteCloudBaseProjectLatestVersion
* DeleteCloudBaseRunServerVersion
* DeleteGatewayVersion
* DeleteWxGatewayRoute
* DescribeActivityRecord
* DescribeBillingInfo
* DescribeCbrServerVersion
* DescribeCloudBaseProjectLatestVersionList
* DescribeCloudBaseProjectVersionList
* DescribeCloudBaseRunResource
* DescribeCloudBaseRunResourceForExtend
* DescribeCloudBaseRunServer
* DescribeCloudBaseRunServerVersion
* DescribeCloudBaseRunVersion
* DescribeCloudBaseRunVersionSnapshot
* DescribeCurveData
* DescribeDownloadFile
* DescribeEnvDealRegion
* DescribeEnvFreeQuota
* DescribeEnvPostpaidDeduct
* DescribeExtensionUploadInfo
* DescribeExtraPkgBillingInfo
* DescribeGatewayCurveData
* DescribeGatewayVersions
* DescribeGraphData
* DescribePostpayFreeQuotas
* DescribePostpayPackageFreeQuotas
* DescribeSmsQuotas
* DescribeSpecialCostItems
* DescribeUserActivityInfo
* DescribeWxGatewayRoutes
* DescribeWxGateways
* EstablishCloudBaseRunServer
* EstablishWxGatewayRoute
* FreezeCloudBaseRunServers
* ModifyCloudBaseRunServerFlowConf
* ModifyCloudBaseRunServerVersion
* ModifyGatewayVersionTraffic
* ReplaceActivityRecord
* UnfreezeCloudBaseRunServers

<font color="#dd0000">**删除数据结构**：</font>

* ActivityRecordItem
* BanConfig
* CbrPackageInfo
* CbrRepoInfo
* CloudBaseCapabilities
* CloudBaseCodeRepoDetail
* CloudBaseCodeRepoName
* CloudBaseEsInfo
* CloudBaseProjectVersion
* CloudBaseRunEmptyDirVolumeSource
* CloudBaseRunImageInfo
* CloudBaseRunImageSecretInfo
* CloudBaseRunKVPriority
* CloudBaseRunNfsVolumeSource
* CloudBaseRunServerVersionItem
* CloudBaseRunServiceVolumeHostPath
* CloudBaseRunServiceVolumeMount
* CloudBaseRunSideSpec
* CloudBaseRunVersionFlowItem
* CloudBaseRunVolumeMount
* CloudBaseRunVpcInfo
* CloudBaseRunVpcSubnet
* CloudBaseSecurityContext
* CloudRunServiceSimpleVersionSnapshot
* CloudRunServiceVolume
* CodeSource
* CustomHeader
* CustomLogConfig
* CustomRequestToAdd
* EnvBillingInfoItem
* ExtensionFile
* ExtensionFileInfo
* FreequotaInfo
* FrequencyLimitConfig
* GatewayItem
* GatewayVersionItem
* HpaPolicy
* ObjectKV
* OrderInfo
* PackageFreeQuotaInfo
* PostPaidEnvDeductInfo
* PostpayEnvQuota
* SmsFreeQuota
* SpecialCostItem
* TkeClusterInfo
* WxGatewayCustomConfig
* WxGatewayRountItem



## TI-ONE 训练平台(tione) 版本：2021-11-11

### 第 104 次发布

发布时间：2026-03-09 03:15:50

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateModelService](https://cloud.tencent.com/document/api/851/82291)

	* 新增入参：GatewayLogConfig


修改数据结构：

* [ServiceGroup](https://cloud.tencent.com/document/api/851/75051#ServiceGroup)

	* 新增成员：GatewayLogConfig




## TI-ONE 训练平台(tione) 版本：2019-10-22



## 容器服务(tke) 版本：2022-05-01

### 第 22 次发布

发布时间：2026-03-06 03:21:13

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [Disk](https://cloud.tencent.com/document/api/457/103206#Disk)

	* 新增成员：Encrypt, KmsKeyId




## 容器服务(tke) 版本：2018-05-25

### 第 221 次发布

发布时间：2026-03-06 03:17:01

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [AddExistedInstances](https://cloud.tencent.com/document/api/457/31865)

	* 新增入参：NodeType




## TSF-Polaris&ZK&网关(tse) 版本：2020-12-07

### 第 104 次发布

发布时间：2026-03-09 03:27:55

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateGovernanceLaneGroups](https://cloud.tencent.com/document/api/1364/129016)
* [DeleteGovernanceLaneGroups](https://cloud.tencent.com/document/api/1364/129015)
* [DescribeGovernanceLaneGroups](https://cloud.tencent.com/document/api/1364/129014)
* [ModifyGovernanceLaneGroups](https://cloud.tencent.com/document/api/1364/129013)

新增数据结构：

* [Argument](https://cloud.tencent.com/document/api/1364/54942#Argument)
* [ArgumentValue](https://cloud.tencent.com/document/api/1364/54942#ArgumentValue)
* [DeleteGovernanceLaneGroup](https://cloud.tencent.com/document/api/1364/54942#DeleteGovernanceLaneGroup)
* [GovernanceLaneGroup](https://cloud.tencent.com/document/api/1364/54942#GovernanceLaneGroup)
* [GovernanceLaneRule](https://cloud.tencent.com/document/api/1364/54942#GovernanceLaneRule)
* [GovernanceServiceDestination](https://cloud.tencent.com/document/api/1364/54942#GovernanceServiceDestination)
* [Label](https://cloud.tencent.com/document/api/1364/54942#Label)
* [LaneTrafficEntry](https://cloud.tencent.com/document/api/1364/54942#LaneTrafficEntry)
* [RoutingDestinationRuleLabel](https://cloud.tencent.com/document/api/1364/54942#RoutingDestinationRuleLabel)
* [ServiceGatewaySelector](https://cloud.tencent.com/document/api/1364/54942#ServiceGatewaySelector)
* [ServiceSelector](https://cloud.tencent.com/document/api/1364/54942#ServiceSelector)
* [TSEGatewaySelector](https://cloud.tencent.com/document/api/1364/54942#TSEGatewaySelector)
* [TrafficGray](https://cloud.tencent.com/document/api/1364/54942#TrafficGray)



## 云点播(vod) 版本：2024-07-18



## 云点播(vod) 版本：2018-07-17

### 第 234 次发布

发布时间：2026-03-06 03:35:53

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeTaskDetail](https://cloud.tencent.com/document/api/266/33431)

	* 新增出参：CreateAigcAdvancedCustomElementTask, CreateAigcCustomVoiceTask

* [RemoveWatermark](https://cloud.tencent.com/document/api/266/79257)


新增数据结构：

* [AdvancedElementInfo](https://cloud.tencent.com/document/api/266/31773#AdvancedElementInfo)
* [CreateAigcAdvancedCustomElementInput](https://cloud.tencent.com/document/api/266/31773#CreateAigcAdvancedCustomElementInput)
* [CreateAigcAdvancedCustomElementOutput](https://cloud.tencent.com/document/api/266/31773#CreateAigcAdvancedCustomElementOutput)
* [CreateAigcAdvancedCustomElementTask](https://cloud.tencent.com/document/api/266/31773#CreateAigcAdvancedCustomElementTask)
* [CreateAigcCustomVoiceInput](https://cloud.tencent.com/document/api/266/31773#CreateAigcCustomVoiceInput)
* [CreateAigcCustomVoiceOutput](https://cloud.tencent.com/document/api/266/31773#CreateAigcCustomVoiceOutput)
* [CreateAigcCustomVoiceTask](https://cloud.tencent.com/document/api/266/31773#CreateAigcCustomVoiceTask)
* [CustomVoiceInfo](https://cloud.tencent.com/document/api/266/31773#CustomVoiceInfo)

### 第 233 次发布

发布时间：2026-03-04 02:33:36

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [SearchMedia](https://cloud.tencent.com/document/api/266/31813)

	* 新增入参：StreamDomains, StreamPaths


修改数据结构：

* [LiveRecordInfo](https://cloud.tencent.com/document/api/266/31773#LiveRecordInfo)

	* 新增成员：Domain, Path




## 私有网络(vpc) 版本：2017-03-12

### 第 297 次发布

发布时间：2026-03-05 02:38:33

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [NatRegionInfoWithArea](https://cloud.tencent.com/document/api/215/15824#NatRegionInfoWithArea)

	* 新增成员：Region


### 第 296 次发布

发布时间：2026-03-03 02:38:05

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeDesignatedZones](https://cloud.tencent.com/document/api/215/128830)

修改接口：

* [CreateVpnGatewaySslServer](https://cloud.tencent.com/document/api/215/70289)

	* 新增入参：DnsServers

* [ModifyVpnGatewaySslServer](https://cloud.tencent.com/document/api/215/100955)

	* 新增入参：DnsServers


新增数据结构：

* [DesignatedZoneInfoDict](https://cloud.tencent.com/document/api/215/15824#DesignatedZoneInfoDict)
* [DnsServers](https://cloud.tencent.com/document/api/215/15824#DnsServers)

修改数据结构：

* [SslVpnSever](https://cloud.tencent.com/document/api/215/15824#SslVpnSever)

	* 新增成员：DnsServers




## 企业微信汽车行业版(wav) 版本：2021-01-29

### 第 18 次发布

发布时间：2026-03-05 02:46:38

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**预下线接口**：</font>

* CreateCorpTag



