# Release v1.3.194

## 腾讯云智能体开发平台(adp) 版本：2026-05-20

### 第 30 次发布

发布时间：2026-10-15 01:08:37

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ResetConversation](https://cloud.tencent.com/document/api/1759/132515)

	* 新增出参：ResetTime, ResetThroughRecordId


修改数据结构：

* [ChannelSpec](https://cloud.tencent.com/document/api/1759/132545#ChannelSpec)

	* 新增成员：ShareGroupConversation, IsDebug




## Agent 沙箱服务(ags) 版本：2025-09-20

### 第 29 次发布

发布时间：2026-10-16 01:09:57

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [Capabilities](https://cloud.tencent.com/document/api/1814/124823#Capabilities)
* [SecurityContext](https://cloud.tencent.com/document/api/1814/124823#SecurityContext)
* [VolumeRef](https://cloud.tencent.com/document/api/1814/124823#VolumeRef)
* [VolumeTemplateRef](https://cloud.tencent.com/document/api/1814/124823#VolumeTemplateRef)

修改数据结构：

* [CustomConfiguration](https://cloud.tencent.com/document/api/1814/124823#CustomConfiguration)

	* 新增成员：SecurityContext

* [CustomConfigurationDetail](https://cloud.tencent.com/document/api/1814/124823#CustomConfigurationDetail)

	* 新增成员：SecurityContext

* [StorageMount](https://cloud.tencent.com/document/api/1814/124823#StorageMount)

	* 新增成员：Volume, VolumeTemplate




## 应用型负载均衡(alb) 版本：2025-10-30

### 第 5 次发布

发布时间：2026-10-15 01:09:49

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [RuleHealthStatusInfo](https://cloud.tencent.com/document/api/1822/133738#RuleHealthStatusInfo)

	* 新增成员：Priority




## AI Agent 安全网关(apis) 版本：2024-08-01

### 第 15 次发布

发布时间：2026-10-19 01:10:41

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateMcpServer](https://cloud.tencent.com/document/api/1627/129634)

* [CreateModelService](https://cloud.tencent.com/document/api/1627/129615)

* [DescribeAgentAppServices](https://cloud.tencent.com/document/api/1627/132420)

	* <font color="#dd0000">**修改入参**：</font>InstanceID

* [ModifyMcpServer](https://cloud.tencent.com/document/api/1627/129630)

* [ModifyModelService](https://cloud.tencent.com/document/api/1627/129611)

* [ModifyService](https://cloud.tencent.com/document/api/1627/132415)


修改数据结构：

* [AgentAppModelServiceDTO](https://cloud.tencent.com/document/api/1627/129635#AgentAppModelServiceDTO)

* [AgentAppServiceDTO](https://cloud.tencent.com/document/api/1627/129635#AgentAppServiceDTO)

* [DescribeMcpServerResponseVO](https://cloud.tencent.com/document/api/1627/129635#DescribeMcpServerResponseVO)

* [DescribeModelServiceResponseVO](https://cloud.tencent.com/document/api/1627/129635#DescribeModelServiceResponseVO)

* [ServiceVO](https://cloud.tencent.com/document/api/1627/129635#ServiceVO)

* [ToolConfigDTO](https://cloud.tencent.com/document/api/1627/129635#ToolConfigDTO)




## 应用性能监控(apm) 版本：2021-06-22

### 第 72 次发布

发布时间：2026-10-19 01:10:57

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyApmApplicationConfig](https://cloud.tencent.com/document/api/1463/125072)

	* 新增入参：EnableSpanStatusConfig, SpanStatusConfig

* [ModifyApmInstance](https://cloud.tencent.com/document/api/1463/89002)

	* 新增入参：SpanStatusConfig


新增数据结构：

* [ErrorSpanConf](https://cloud.tencent.com/document/api/1463/64927#ErrorSpanConf)
* [SpanStatusConfig](https://cloud.tencent.com/document/api/1463/64927#SpanStatusConfig)

修改数据结构：

* [ApmAppConfig](https://cloud.tencent.com/document/api/1463/64927#ApmAppConfig)

	* 新增成员：SpanStatusConfig, EnableSpanStatusConfig

* [ApmInstanceDetail](https://cloud.tencent.com/document/api/1463/64927#ApmInstanceDetail)

	* 新增成员：SpanStatusConfig




## 运维安全中心（堡垒机）(bh) 版本：2023-04-18

### 第 37 次发布

发布时间：2026-10-15 01:12:25

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeDevices](https://cloud.tencent.com/document/api/1025/74415)

	* 新增入参：OwnerAppIds


修改数据结构：

* [Device](https://cloud.tencent.com/document/api/1025/74416#Device)

	* 新增成员：OwnerAppId, OwnerAccountUin, OwnerAccountName

* [Resource](https://cloud.tencent.com/document/api/1025/74416#Resource)

	* 新增成员：OrgAccountQuota




## 云数据库 MySQL(cdb) 版本：2017-03-20

### 第 233 次发布

发布时间：2026-10-16 01:16:33

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [SwitchDrInstanceToMaster](https://cloud.tencent.com/document/api/236/52562)

	* 新增入参：IgnoreSyncDelay




## 本地专用集群(cdc) 版本：2020-12-14

### 第 15 次发布

发布时间：2026-10-16 01:17:25

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeEdgeHostVm](https://cloud.tencent.com/document/api/1346/139515)



## 文件存储(cfs) 版本：2019-07-19

### 第 59 次发布

发布时间：2026-10-15 01:17:17

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateAccessCert](https://cloud.tencent.com/document/api/582/115736)

	* 新增入参：Type




## 云防火墙(cfw) 版本：2019-09-04

### 第 118 次发布

发布时间：2026-10-15 01:17:34

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeOverviewResourceQuota](https://cloud.tencent.com/document/api/1132/139467)

新增数据结构：

* [OverviewBlockListQuota](https://cloud.tencent.com/document/api/1132/49071#OverviewBlockListQuota)
* [OverviewBlockListQuotaDetail](https://cloud.tencent.com/document/api/1132/49071#OverviewBlockListQuotaDetail)
* [OverviewCommonRuleQuotaDetail](https://cloud.tencent.com/document/api/1132/49071#OverviewCommonRuleQuotaDetail)
* [OverviewEastWestDetail](https://cloud.tencent.com/document/api/1132/49071#OverviewEastWestDetail)
* [OverviewLogStorageDetail](https://cloud.tencent.com/document/api/1132/49071#OverviewLogStorageDetail)
* [OverviewNDRDetail](https://cloud.tencent.com/document/api/1132/49071#OverviewNDRDetail)
* [OverviewNorthSouthDetail](https://cloud.tencent.com/document/api/1132/49071#OverviewNorthSouthDetail)
* [OverviewQuotaDetail](https://cloud.tencent.com/document/api/1132/49071#OverviewQuotaDetail)
* [OverviewResourceQuotaDataLite](https://cloud.tencent.com/document/api/1132/49071#OverviewResourceQuotaDataLite)
* [OverviewResourceQuotaDetail](https://cloud.tencent.com/document/api/1132/49071#OverviewResourceQuotaDetail)
* [OverviewVpcRuleQuotaDetail](https://cloud.tencent.com/document/api/1132/49071#OverviewVpcRuleQuotaDetail)
* [VolumeUsageDetail](https://cloud.tencent.com/document/api/1132/49071#VolumeUsageDetail)



## 负载均衡(clb) 版本：2018-03-17

### 第 170 次发布

发布时间：2026-10-16 01:20:59

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateModelRouter](https://cloud.tencent.com/document/api/214/133217)

	* 新增入参：ImageConfig

* [ModifyModelRouterAttributes](https://cloud.tencent.com/document/api/214/133203)

	* 新增入参：ImageConfig


新增数据结构：

* [ImageConfig](https://cloud.tencent.com/document/api/214/30694#ImageConfig)

修改数据结构：

* [Coefficient](https://cloud.tencent.com/document/api/214/30694#Coefficient)

	* 新增成员：OutputImageCoefficient

* [ModelRouterDetail](https://cloud.tencent.com/document/api/214/30694#ModelRouterDetail)

	* 新增成员：ImageConfig

* [ServiceProviderHealthCheckConfigItemInput](https://cloud.tencent.com/document/api/214/30694#ServiceProviderHealthCheckConfigItemInput)

	* 新增成员：HealthCheckTimeout

* [ServiceProviderHealthCheckConfigItemOutput](https://cloud.tencent.com/document/api/214/30694#ServiceProviderHealthCheckConfigItemOutput)

	* 新增成员：HealthCheckTimeout




## 云安全一体化平台(csip) 版本：2022-11-21

### 第 109 次发布

发布时间：2026-10-19 01:22:53

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [OrganizationUserInfo](https://cloud.tencent.com/document/api/664/90825#OrganizationUserInfo)

	* 新增成员：SharedModules, SecurityScore


### 第 108 次发布

发布时间：2026-10-15 01:22:19

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [DspmDbAsset](https://cloud.tencent.com/document/api/664/90825#DspmDbAsset)

	* 新增成员：LogDeliveryDisableMessage, DbVersion

* [DspmRiskCount](https://cloud.tencent.com/document/api/664/90825#DspmRiskCount)

	* 新增成员：AccountRisk, ConfigRisk

* [UserDspmInfo](https://cloud.tencent.com/document/api/664/90825#UserDspmInfo)

	* 新增成员：Role, NodeName




## TDSQL-C MySQL 版(cynosdb) 版本：2019-01-07

### 第 197 次发布

发布时间：2026-10-16 01:33:32

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateProxyEndPoint](https://cloud.tencent.com/document/api/1003/94133)

	* 新增入参：TransAllSplit

* [ModifyProxyRwSplit](https://cloud.tencent.com/document/api/1003/94129)

	* 新增入参：TransAllSplit


修改数据结构：

* [ClusterInstanceDetail](https://cloud.tencent.com/document/api/1003/48097#ClusterInstanceDetail)

	* 新增成员：Uin, AppId, ClusterId, ClusterName, Cpu, MinCpu, MaxCpu, Memory, ProjectId, Region, Zone, PhysicalZone, Status, StatusDesc, DbType, DbVersion, Storage, StorageId, StoragePayMode, CreateTime, UpdateTime, IsolateTime, DestroyTime, DestroyDeadlineText, PeriodEndTime, VpcId, SubnetId, Vip, Vport, NetType, WanIP, WanPort, WanStatus, WanDomain, InstanceNetInfo, PayMode, ProcessingTask, RenewFlag, Tasks, IsFreeze, ResourceTags, InstanceAbility, DeviceType, GdnId

* [ProxyGroupRwInfo](https://cloud.tencent.com/document/api/1003/48097#ProxyGroupRwInfo)

	* 新增成员：TransAllSplit




## 数据湖计算 DLC(dlc) 版本：2021-01-25

### 第 186 次发布

发布时间：2026-10-16 01:36:25

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [AddDeployment](https://cloud.tencent.com/document/api/1342/137397)

	* 新增入参：NetworkConfig

* [CreateInferenceService](https://cloud.tencent.com/document/api/1342/135625)

	* 新增入参：NetworkConfig

	* 新增出参：NetworkConfig

* [CreateJobSpec](https://cloud.tencent.com/document/api/1342/135458)

	* 新增入参：NetworkConfig

	* 新增出参：NetworkConfig

* [CreateLab](https://cloud.tencent.com/document/api/1342/135449)

	* 新增入参：NetworkConfig

	* 新增出参：NetworkConfig

* [CreateRayCluster](https://cloud.tencent.com/document/api/1342/135422)

	* 新增入参：NetworkConfig

	* 新增出参：NetworkConfig

* [GetJobSpec](https://cloud.tencent.com/document/api/1342/135456)

	* 新增出参：NetworkConfig

* [GetLabDetail](https://cloud.tencent.com/document/api/1342/135447)

	* 新增出参：NetworkConfig

* [GetRayCluster](https://cloud.tencent.com/document/api/1342/135420)

	* 新增出参：NetworkConfig

* [ModifyTrainingJobSpec](https://cloud.tencent.com/document/api/1342/137371)

	* 新增入参：NetworkConfig

* [StartLab](https://cloud.tencent.com/document/api/1342/135438)

	* 新增出参：NetworkConfig

* [StartRayCluster](https://cloud.tencent.com/document/api/1342/135411)

	* 新增出参：NetworkConfig

* [StopLab](https://cloud.tencent.com/document/api/1342/135437)

	* 新增出参：NetworkConfig

* [StopRayCluster](https://cloud.tencent.com/document/api/1342/135410)

	* 新增出参：NetworkConfig

* [SubmitTrainingJob](https://cloud.tencent.com/document/api/1342/137369)

	* 新增入参：NetworkConfig

* [UpdateJobSpec](https://cloud.tencent.com/document/api/1342/135452)

	* 新增入参：NetworkConfig

	* 新增出参：NetworkConfig

* [UpdateLab](https://cloud.tencent.com/document/api/1342/135436)

	* 新增入参：NetworkConfig

	* 新增出参：NetworkConfig

* [UpdateRayCluster](https://cloud.tencent.com/document/api/1342/135409)

	* 新增入参：NetworkConfig

	* 新增出参：NetworkConfig


修改数据结构：

* [JobSpec](https://cloud.tencent.com/document/api/1342/53778#JobSpec)

	* 新增成员：NetworkConfig

* [LabResponse](https://cloud.tencent.com/document/api/1342/53778#LabResponse)

	* 新增成员：NetworkConfig

* [RayClusterEntity](https://cloud.tencent.com/document/api/1342/53778#RayClusterEntity)

	* 新增成员：NetworkConfig

* [TrainingJobSpec](https://cloud.tencent.com/document/api/1342/53778#TrainingJobSpec)

	* 新增成员：SubAccountUin, SubAccountName, NetworkConfig




## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 324 次发布

发布时间：2026-10-15 01:36:45

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateDraftContractByPromptsTask](https://cloud.tencent.com/document/api/1323/129860)

	* 新增入参：ReviewTaskId




## 防火墙管理(fwm) 版本：2025-06-11

### 第 4 次发布

发布时间：2026-10-19 15:31:48

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateRuleGroup](https://cloud.tencent.com/document/api/1812/139554)

### 第 3 次发布

发布时间：2026-10-19 15:19:40

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateEdgeAclRule](https://cloud.tencent.com/document/api/1812/133582)

	* 新增出参：RuleIds

* [CreateNatAclRule](https://cloud.tencent.com/document/api/1812/133594)

	* 新增出参：RuleIds

* [CreateSecurityGroupRule](https://cloud.tencent.com/document/api/1812/133576)

	* 新增出参：RuleIds

* [CreateStrategy](https://cloud.tencent.com/document/api/1812/133558)

	* 新增出参：StrategyIds

* [CreateVpcAclRule](https://cloud.tencent.com/document/api/1812/133588)

	* 新增出参：RuleIds


修改数据结构：

* [EdgeAclRuleResp](https://cloud.tencent.com/document/api/1812/133595#EdgeAclRuleResp)

	* 新增成员：HitCount, LastHitTime

* [NatAclRuleResp](https://cloud.tencent.com/document/api/1812/133595#NatAclRuleResp)

	* 新增成员：HitCount, LastHitTime

* [OrganSummary](https://cloud.tencent.com/document/api/1812/133595#OrganSummary)

	* 新增成员：FwmShareUserQuota

* [VpcAclRuleResp](https://cloud.tencent.com/document/api/1812/133595#VpcAclRuleResp)

	* 新增成员：HitCount, LastHitTime




## 物联网开发平台(iotexplorer) 版本：2019-04-23

### 第 163 次发布

发布时间：2026-10-19 01:44:42

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除接口**：</font>

* BindTWeTalkAIBot
* CreateTWeTalkAIBot
* DeleteTWeTalkAIBot
* DescribeTWeTalkAIBot
* GetTWeTalkAIBotList
* ModifyTWeTalkAIBot
* UnbindTWeTalkAIBot

<font color="#dd0000">**删除数据结构**：</font>

* TalkAIBotInfo
* TalkAgentConfigInfo
* TalkProductInfo

### 第 162 次发布

发布时间：2026-10-16 01:46:59

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ResetCloudStorage](https://cloud.tencent.com/document/api/1081/104635)

	* 新增入参：OrderId

* [TransferCloudStorage](https://cloud.tencent.com/document/api/1081/104633)

	* 新增入参：ChannelId, ToChannelId




## 云数据库 KeeWiDB(keewidb) 版本：2022-03-08

### 第 13 次发布

发布时间：2026-10-19 01:48:27

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyConnectionConfig](https://cloud.tencent.com/document/api/1520/86197)




## 云数据库 MongoDB(mongodb) 版本：2019-07-25

### 第 80 次发布

发布时间：2026-10-15 01:47:59

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [ModifyNetworkAddress](https://cloud.tencent.com/document/api/240/38576#ModifyNetworkAddress)

	* 新增成员：NodeType




## 云数据库 MongoDB(mongodb) 版本：2018-04-08



## 媒体处理(mps) 版本：2019-06-12

### 第 255 次发布

发布时间：2026-10-19 01:53:44

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CloneViral](https://cloud.tencent.com/document/api/862/135033)

	* <font color="#dd0000">**修改入参**：</font>Product




## 文字识别(ocr) 版本：2018-11-19

### 第 272 次发布

发布时间：2026-10-15 01:51:23

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**预下线接口**：</font>

* ImageEnhancement



## 云数据库Redis(redis) 版本：2018-04-12

### 第 111 次发布

发布时间：2026-10-19 01:58:50

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeBackupEncryptConfig](https://cloud.tencent.com/document/api/239/139546)
* [ModifyBackupEncrypt](https://cloud.tencent.com/document/api/239/139545)



## 前端性能监控(rum) 版本：2021-06-22

### 第 54 次发布

发布时间：2026-10-15 01:54:31

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeFileCertificate](https://cloud.tencent.com/document/api/1464/139468)



## 腾讯云数据库 AI 服务(tdai) 版本：2025-07-17

### 第 22 次发布

发布时间：2026-10-16 02:07:59

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [AgentInstance](https://cloud.tencent.com/document/api/1813/123239#AgentInstance)

	* 新增成员：MemInstanceId


### 第 21 次发布

发布时间：2026-10-15 02:01:20

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeMemoryPlusSpace](https://cloud.tencent.com/document/api/1813/132184)

	* 新增出参：Compressed

* [ModifyMemoryPlusSpace](https://cloud.tencent.com/document/api/1813/132179)

	* 新增入参：Compressed


修改数据结构：

* [MemoryPlusInfo](https://cloud.tencent.com/document/api/1813/123239#MemoryPlusInfo)

	* 新增成员：Compressed




## 边缘安全加速平台(teo) 版本：2022-09-01

### 第 167 次发布

发布时间：2026-10-16 02:09:04

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [TCEOCaptchaOption](https://cloud.tencent.com/document/api/1552/80721#TCEOCaptchaOption)

	* 新增成员：MobileCaptchaMode




## 边缘安全加速平台(teo) 版本：2022-01-06



## TI-ONE 训练平台(tione) 版本：2021-11-11

### 第 141 次发布

发布时间：2026-10-19 02:09:16

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateNotebookImage](https://cloud.tencent.com/document/api/851/139547)

### 第 140 次发布

发布时间：2026-10-16 02:10:53

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateTrainingModel](https://cloud.tencent.com/document/api/851/75074)

	* 新增入参：TiProjectId

* [DeleteTrainingModel](https://cloud.tencent.com/document/api/851/75073)

	* 新增入参：TiProjectId

* [DeleteTrainingModelVersion](https://cloud.tencent.com/document/api/851/75072)

	* 新增入参：TiProjectId

* [DescribeModelAccelerateTask](https://cloud.tencent.com/document/api/851/82606)

	* 新增入参：TiProjectId

* [DescribeModelAccelerateVersions](https://cloud.tencent.com/document/api/851/102555)

	* 新增入参：TiProjectId

* [DescribePlatformImages](https://cloud.tencent.com/document/api/851/120508)

	* 新增入参：TiProjectId

* [DescribeTrainingModelVersion](https://cloud.tencent.com/document/api/851/75071)

	* 新增入参：TiProjectId

* [DescribeTrainingModelVersions](https://cloud.tencent.com/document/api/851/75070)

	* 新增入参：TiProjectId

* [StopModelAccelerateTask](https://cloud.tencent.com/document/api/851/82603)

	* 新增入参：TiProjectId




## TI-ONE 训练平台(tione) 版本：2019-10-22



## 容器服务(tke) 版本：2022-05-01

### 第 32 次发布

发布时间：2026-10-19 02:11:16

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [CreateNativeNodePoolParam](https://cloud.tencent.com/document/api/457/103206#CreateNativeNodePoolParam)

	* 新增成员：RuntimeVersion




## 容器服务(tke) 版本：2018-05-25



## 云点播(vod) 版本：2024-07-18



## 云点播(vod) 版本：2018-07-17

### 第 290 次发布

发布时间：2026-10-19 02:14:55

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeTaskDetail](https://cloud.tencent.com/document/api/266/33431)

	* 新增出参：VideoDepthEstimationTask


新增数据结构：

* [VideoDepthEstimationInput](https://cloud.tencent.com/document/api/266/31773#VideoDepthEstimationInput)
* [VideoDepthEstimationOutput](https://cloud.tencent.com/document/api/266/31773#VideoDepthEstimationOutput)
* [VideoDepthEstimationTask](https://cloud.tencent.com/document/api/266/31773#VideoDepthEstimationTask)



## Web 应用防火墙(waf) 版本：2018-01-25

### 第 169 次发布

发布时间：2026-10-16 02:19:23

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [AddLLMSdkAccess](https://cloud.tencent.com/document/api/627/139517)
* [DescribeAfterPayInstances](https://cloud.tencent.com/document/api/627/139516)

新增数据结构：

* [AfterPayInstanceInfo](https://cloud.tencent.com/document/api/627/53609#AfterPayInstanceInfo)
* [LLMImageTrendItem](https://cloud.tencent.com/document/api/627/53609#LLMImageTrendItem)
* [LLMTokenTrendItem](https://cloud.tencent.com/document/api/627/53609#LLMTokenTrendItem)
* [LLMTokenUsage](https://cloud.tencent.com/document/api/627/53609#LLMTokenUsage)

### 第 168 次发布

发布时间：2026-10-15 02:12:22

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [LLMSdkAccessItem](https://cloud.tencent.com/document/api/627/53609#LLMSdkAccessItem)

	* 新增成员：UrlDetectStatus




## 数据开发治理平台 WeData(wedata) 版本：2025-08-06

### 第 27 次发布

发布时间：2026-10-19 02:22:35

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [CreateTaskSchedulerConfiguration](https://cloud.tencent.com/document/api/1267/123643#CreateTaskSchedulerConfiguration)

	* 新增成员：MinuteOffset

* [TaskSchedulerConfiguration](https://cloud.tencent.com/document/api/1267/123643#TaskSchedulerConfiguration)

	* 新增成员：MinuteOffset

* [WorkflowSchedulerConfiguration](https://cloud.tencent.com/document/api/1267/123643#WorkflowSchedulerConfiguration)

	* 新增成员：MinuteOffset

* [WorkflowSchedulerConfigurationInfo](https://cloud.tencent.com/document/api/1267/123643#WorkflowSchedulerConfigurationInfo)

	* 新增成员：MinuteOffset




## 数据开发治理平台 WeData(wedata) 版本：2021-08-20

### 第 206 次发布

发布时间：2026-10-19 02:20:47

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [TaskDsDTO](https://cloud.tencent.com/document/api/1267/76336#TaskDsDTO)

	* 新增成员：MinuteOffset

* [WorkflowScheduleDtoDs](https://cloud.tencent.com/document/api/1267/76336#WorkflowScheduleDtoDs)

	* 新增成员：MinuteOffset




