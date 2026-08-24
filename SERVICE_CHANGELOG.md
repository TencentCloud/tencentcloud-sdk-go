# Release v1.3.168

## Agent 沙箱服务(ags) 版本：2025-09-20

### 第 19 次发布

发布时间：2026-08-25 01:09:44

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [AcquireDeploymentToken](https://cloud.tencent.com/document/api/1814/136842)
* [CreateDeployment](https://cloud.tencent.com/document/api/1814/136841)
* [DeleteDeployment](https://cloud.tencent.com/document/api/1814/136840)
* [DescribeDeployment](https://cloud.tencent.com/document/api/1814/136839)
* [DescribeDeploymentList](https://cloud.tencent.com/document/api/1814/136838)
* [ModifyDeployment](https://cloud.tencent.com/document/api/1814/136837)

新增数据结构：

* [AffinityConfiguration](https://cloud.tencent.com/document/api/1814/124823#AffinityConfiguration)
* [ComputerConfiguration](https://cloud.tencent.com/document/api/1814/124823#ComputerConfiguration)
* [Deployment](https://cloud.tencent.com/document/api/1814/124823#Deployment)
* [LifecycleConfiguration](https://cloud.tencent.com/document/api/1814/124823#LifecycleConfiguration)
* [ScalingConfiguration](https://cloud.tencent.com/document/api/1814/124823#ScalingConfiguration)
* [WAAConfiguration](https://cloud.tencent.com/document/api/1814/124823#WAAConfiguration)

修改数据结构：

* [SandboxInstance](https://cloud.tencent.com/document/api/1814/124823#SandboxInstance)

	* 新增成员：ComputerConfiguration

* [SandboxTool](https://cloud.tencent.com/document/api/1814/124823#SandboxTool)

	* 新增成员：ComputerConfiguration




## 灾备中心(bdrc) 版本：2026-03-30

### 第 1 次发布

发布时间：2026-08-24 21:38:35

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [ApplyBackupGroup](https://cloud.tencent.com/document/api/1824/136823)
* [BindAutoBackupPolicy](https://cloud.tencent.com/document/api/1824/136822)
* [CreateAutoBackupPolicy](https://cloud.tencent.com/document/api/1824/136821)
* [CreateBackupGroup](https://cloud.tencent.com/document/api/1824/136820)
* [CreateBackupVault](https://cloud.tencent.com/document/api/1824/136819)
* [CreateDisasterRecoveryProtectGroup](https://cloud.tencent.com/document/api/1824/136782)
* [CreateDisasterRecoverySitePair](https://cloud.tencent.com/document/api/1824/136781)
* [CreateDisasterRecoveryVpcMapping](https://cloud.tencent.com/document/api/1824/136780)
* [CreateFileBackup](https://cloud.tencent.com/document/api/1824/136818)
* [CreateFileBackupPlan](https://cloud.tencent.com/document/api/1824/136817)
* [CreateFileRestoreTask](https://cloud.tencent.com/document/api/1824/136816)
* [CreateInstanceCopyPair](https://cloud.tencent.com/document/api/1824/136779)
* [CreateInstanceDrillPairs](https://cloud.tencent.com/document/api/1824/136778)
* [CreateSecurityGroupMapping](https://cloud.tencent.com/document/api/1824/136777)
* [DeleteAutoBackupPolicies](https://cloud.tencent.com/document/api/1824/136815)
* [DeleteBackupGroups](https://cloud.tencent.com/document/api/1824/136814)
* [DeleteBackupVaults](https://cloud.tencent.com/document/api/1824/136813)
* [DeleteCopyPairs](https://cloud.tencent.com/document/api/1824/136776)
* [DeleteDisasterRecoveryProtectGroups](https://cloud.tencent.com/document/api/1824/136775)
* [DeleteDisasterRecoverySitePairs](https://cloud.tencent.com/document/api/1824/136774)
* [DeleteDisasterRecoveryVpcMapping](https://cloud.tencent.com/document/api/1824/136773)
* [DeleteDrillPairs](https://cloud.tencent.com/document/api/1824/136772)
* [DeleteFileBackupPlans](https://cloud.tencent.com/document/api/1824/136812)
* [DeleteFileBackups](https://cloud.tencent.com/document/api/1824/136811)
* [DeleteSecurityGroupMapping](https://cloud.tencent.com/document/api/1824/136771)
* [DescribeAutoBackupPolicies](https://cloud.tencent.com/document/api/1824/136810)
* [DescribeBackupGroupRollbackTasks](https://cloud.tencent.com/document/api/1824/136809)
* [DescribeBackupGroups](https://cloud.tencent.com/document/api/1824/136808)
* [DescribeBackupGroupsDeniedActions](https://cloud.tencent.com/document/api/1824/136807)
* [DescribeBackupInstances](https://cloud.tencent.com/document/api/1824/136806)
* [DescribeBackupOverviewGeneral](https://cloud.tencent.com/document/api/1824/136805)
* [DescribeBackupPlans](https://cloud.tencent.com/document/api/1824/136804)
* [DescribeBackupVaults](https://cloud.tencent.com/document/api/1824/136803)
* [DescribeBackupVaultsDeniedActions](https://cloud.tencent.com/document/api/1824/136802)
* [DescribeCommonBackupPoints](https://cloud.tencent.com/document/api/1824/136801)
* [DescribeCopyPairs](https://cloud.tencent.com/document/api/1824/136770)
* [DescribeCopyPairsDeniedActions](https://cloud.tencent.com/document/api/1824/136769)
* [DescribeDisasterRecoveryDrillGroups](https://cloud.tencent.com/document/api/1824/136768)
* [DescribeDisasterRecoveryOverview](https://cloud.tencent.com/document/api/1824/136767)
* [DescribeDisasterRecoveryProtectGroups](https://cloud.tencent.com/document/api/1824/136766)
* [DescribeDisasterRecoverySitePairs](https://cloud.tencent.com/document/api/1824/136765)
* [DescribeDisasterRecoverySitePairsDeniedActions](https://cloud.tencent.com/document/api/1824/136764)
* [DescribeDisasterRecoverySupportRegion](https://cloud.tencent.com/document/api/1824/136763)
* [DescribeDisks](https://cloud.tencent.com/document/api/1824/136762)
* [DescribeDrillPairs](https://cloud.tencent.com/document/api/1824/136761)
* [DescribeDrillPairsDeniedActions](https://cloud.tencent.com/document/api/1824/136760)
* [DescribeFileBackupObjects](https://cloud.tencent.com/document/api/1824/136800)
* [DescribeFileBackupPlans](https://cloud.tencent.com/document/api/1824/136799)
* [DescribeFileBackups](https://cloud.tencent.com/document/api/1824/136798)
* [DescribeFileBackupsDeniedActions](https://cloud.tencent.com/document/api/1824/136797)
* [DescribeFileRestoreTasks](https://cloud.tencent.com/document/api/1824/136796)
* [DescribeJobs](https://cloud.tencent.com/document/api/1824/136795)
* [DescribePriceCreateCopyPairs](https://cloud.tencent.com/document/api/1824/136759)
* [DescribeProtectGroupsDeniedActions](https://cloud.tencent.com/document/api/1824/136758)
* [DescribeProtectedInstances](https://cloud.tencent.com/document/api/1824/136794)
* [DescribeSecurityGroupMappings](https://cloud.tencent.com/document/api/1824/136757)
* [DescribeVpcMappings](https://cloud.tencent.com/document/api/1824/136756)
* [FinishFailoverCopyPairs](https://cloud.tencent.com/document/api/1824/136755)
* [ModifyAutoBackupPolicyAttribute](https://cloud.tencent.com/document/api/1824/136793)
* [ModifyBackupAttribute](https://cloud.tencent.com/document/api/1824/136792)
* [ModifyBackupVaultAttribute](https://cloud.tencent.com/document/api/1824/136791)
* [ModifyCopyPairAttribute](https://cloud.tencent.com/document/api/1824/136754)
* [ModifyDrillGroupAttribute](https://cloud.tencent.com/document/api/1824/136753)
* [ModifyDrillPairAttribute](https://cloud.tencent.com/document/api/1824/136752)
* [ModifyFileBackupAttribute](https://cloud.tencent.com/document/api/1824/136790)
* [ModifyFileBackupPlan](https://cloud.tencent.com/document/api/1824/136789)
* [ModifyProtectGroupAttribute](https://cloud.tencent.com/document/api/1824/136751)
* [ModifySitePairAttribute](https://cloud.tencent.com/document/api/1824/136750)
* [ReportAgentMetrics](https://cloud.tencent.com/document/api/1824/136788)
* [ReportGatewayHeartbeat](https://cloud.tencent.com/document/api/1824/136787)
* [ReportJobProgress](https://cloud.tencent.com/document/api/1824/136786)
* [RunCopyPairTasks](https://cloud.tencent.com/document/api/1824/136749)
* [RunFailoverCopyPairs](https://cloud.tencent.com/document/api/1824/136748)
* [RunInstancesWithBackupGroup](https://cloud.tencent.com/document/api/1824/136785)
* [StopCopyPairTasks](https://cloud.tencent.com/document/api/1824/136747)
* [UnbindAutoBackupPolicy](https://cloud.tencent.com/document/api/1824/136784)

新增数据结构：

* [AdvancedRetentionPolicy](https://cloud.tencent.com/document/api/1824/136824#AdvancedRetentionPolicy)
* [ApplyDisk](https://cloud.tencent.com/document/api/1824/136824#ApplyDisk)
* [AspInfo](https://cloud.tencent.com/document/api/1824/136824#AspInfo)
* [AutoBackupPolicy](https://cloud.tencent.com/document/api/1824/136824#AutoBackupPolicy)
* [AutomationServiceEnabled](https://cloud.tencent.com/document/api/1824/136824#AutomationServiceEnabled)
* [BackupDeniedAction](https://cloud.tencent.com/document/api/1824/136824#BackupDeniedAction)
* [BackupDetail](https://cloud.tencent.com/document/api/1824/136824#BackupDetail)
* [BackupGroup](https://cloud.tencent.com/document/api/1824/136824#BackupGroup)
* [BackupGroupDeniedAction](https://cloud.tencent.com/document/api/1824/136824#BackupGroupDeniedAction)
* [BackupGroupRollbackTask](https://cloud.tencent.com/document/api/1824/136824#BackupGroupRollbackTask)
* [BackupInfo](https://cloud.tencent.com/document/api/1824/136824#BackupInfo)
* [BackupInstance](https://cloud.tencent.com/document/api/1824/136824#BackupInstance)
* [BackupPlan](https://cloud.tencent.com/document/api/1824/136824#BackupPlan)
* [BackupPolicyOverview](https://cloud.tencent.com/document/api/1824/136824#BackupPolicyOverview)
* [BackupVault](https://cloud.tencent.com/document/api/1824/136824#BackupVault)
* [BackupVaultOverview](https://cloud.tencent.com/document/api/1824/136824#BackupVaultOverview)
* [BasicServicesSettings](https://cloud.tencent.com/document/api/1824/136824#BasicServicesSettings)
* [CommonBackupPoint](https://cloud.tencent.com/document/api/1824/136824#CommonBackupPoint)
* [CopyPair](https://cloud.tencent.com/document/api/1824/136824#CopyPair)
* [CopyPairDeniedAction](https://cloud.tencent.com/document/api/1824/136824#CopyPairDeniedAction)
* [CopyPairPrice](https://cloud.tencent.com/document/api/1824/136824#CopyPairPrice)
* [CopyPairPriceDetail](https://cloud.tencent.com/document/api/1824/136824#CopyPairPriceDetail)
* [CreateInstanceModel](https://cloud.tencent.com/document/api/1824/136824#CreateInstanceModel)
* [CrossCloudDetails](https://cloud.tencent.com/document/api/1824/136824#CrossCloudDetails)
* [DeleteDrillPairResult](https://cloud.tencent.com/document/api/1824/136824#DeleteDrillPairResult)
* [DeniedAction](https://cloud.tencent.com/document/api/1824/136824#DeniedAction)
* [DisasterRecoveryDrillGroup](https://cloud.tencent.com/document/api/1824/136824#DisasterRecoveryDrillGroup)
* [DisasterRecoveryOverview](https://cloud.tencent.com/document/api/1824/136824#DisasterRecoveryOverview)
* [DiskCopyPairForCvm](https://cloud.tencent.com/document/api/1824/136824#DiskCopyPairForCvm)
* [DiskInfo](https://cloud.tencent.com/document/api/1824/136824#DiskInfo)
* [DiskModel](https://cloud.tencent.com/document/api/1824/136824#DiskModel)
* [DrillPair](https://cloud.tencent.com/document/api/1824/136824#DrillPair)
* [DrillPairDeniedAction](https://cloud.tencent.com/document/api/1824/136824#DrillPairDeniedAction)
* [DrilledResourceStatus](https://cloud.tencent.com/document/api/1824/136824#DrilledResourceStatus)
* [EnhancedService](https://cloud.tencent.com/document/api/1824/136824#EnhancedService)
* [FileBackupOverview](https://cloud.tencent.com/document/api/1824/136824#FileBackupOverview)
* [FilterModel](https://cloud.tencent.com/document/api/1824/136824#FilterModel)
* [FlowControlRule](https://cloud.tencent.com/document/api/1824/136824#FlowControlRule)
* [InstanceBackupOverview](https://cloud.tencent.com/document/api/1824/136824#InstanceBackupOverview)
* [InstanceChargePrepaid](https://cloud.tencent.com/document/api/1824/136824#InstanceChargePrepaid)
* [InternetAccessible](https://cloud.tencent.com/document/api/1824/136824#InternetAccessible)
* [LoginSettings](https://cloud.tencent.com/document/api/1824/136824#LoginSettings)
* [Placement](https://cloud.tencent.com/document/api/1824/136824#Placement)
* [PlanInfo](https://cloud.tencent.com/document/api/1824/136824#PlanInfo)
* [Policy](https://cloud.tencent.com/document/api/1824/136824#Policy)
* [ProtectGroup](https://cloud.tencent.com/document/api/1824/136824#ProtectGroup)
* [ProtectGroupDeniedAction](https://cloud.tencent.com/document/api/1824/136824#ProtectGroupDeniedAction)
* [ProtectInstance](https://cloud.tencent.com/document/api/1824/136824#ProtectInstance)
* [ProtectedResource](https://cloud.tencent.com/document/api/1824/136824#ProtectedResource)
* [ProtectedResourceOverview](https://cloud.tencent.com/document/api/1824/136824#ProtectedResourceOverview)
* [ProtectedResourceStatus](https://cloud.tencent.com/document/api/1824/136824#ProtectedResourceStatus)
* [ResourcePlan](https://cloud.tencent.com/document/api/1824/136824#ResourcePlan)
* [ResourceProtectStat](https://cloud.tencent.com/document/api/1824/136824#ResourceProtectStat)
* [RestoreTask](https://cloud.tencent.com/document/api/1824/136824#RestoreTask)
* [RunSecurityServiceEnabled](https://cloud.tencent.com/document/api/1824/136824#RunSecurityServiceEnabled)
* [SecurityGroupMapping](https://cloud.tencent.com/document/api/1824/136824#SecurityGroupMapping)
* [SitePair](https://cloud.tencent.com/document/api/1824/136824#SitePair)
* [SitePairDeniedAction](https://cloud.tencent.com/document/api/1824/136824#SitePairDeniedAction)
* [SupportRegionInfo](https://cloud.tencent.com/document/api/1824/136824#SupportRegionInfo)
* [SupportZoneRule](https://cloud.tencent.com/document/api/1824/136824#SupportZoneRule)
* [TypeCount](https://cloud.tencent.com/document/api/1824/136824#TypeCount)
* [VaultDeniedAction](https://cloud.tencent.com/document/api/1824/136824#VaultDeniedAction)
* [VirtualPrivateCloud](https://cloud.tencent.com/document/api/1824/136824#VirtualPrivateCloud)
* [VpcMapping](https://cloud.tencent.com/document/api/1824/136824#VpcMapping)



## 云原生智能网关(cngw) 版本：2023-04-18

### 第 6 次发布

发布时间：2026-08-25 01:38:27

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateCloudNativeAPIGatewayConsumer](https://cloud.tencent.com/document/api/1826/133125)

	* 新增入参：Priority

* [CreateCloudNativeAPIGatewayLLMModelAPI](https://cloud.tencent.com/document/api/1826/133115)

	* 新增入参：MaxDocumentsConfig, SensitiveWordRoute

* [CreateCloudNativeAPIGatewayLLMModelService](https://cloud.tencent.com/document/api/1826/133132)

	* 新增入参：LoadBalanceConfig

* [CreateCloudNativeAPIGatewayMCPServer](https://cloud.tencent.com/document/api/1826/133160)

	* 新增入参：PreserveHost

* [CreateCloudNativeAPIGatewaySecretKey](https://cloud.tencent.com/document/api/1826/133141)

	* 新增入参：AKSKCredentialConfig, CAMCredentialConfig, BearerTokenCredentialConfig, CustomHeaderCredentialConfig, QueryParamCredentialConfig, BasicCredentialConfig

* [DescribeCloudNativeAPIGatewayMCPServerList](https://cloud.tencent.com/document/api/1826/133153)

	* 新增入参：SecretKeyId

* [ModifyCloudNativeAPIGatewayConsumer](https://cloud.tencent.com/document/api/1826/133119)

	* 新增入参：Priority

* [ModifyCloudNativeAPIGatewayLLMModelAPI](https://cloud.tencent.com/document/api/1826/133111)

	* 新增入参：MaxDocumentsConfig, SensitiveWordRoute

* [ModifyCloudNativeAPIGatewayLLMModelService](https://cloud.tencent.com/document/api/1826/133128)

	* 新增入参：CustomProviderName, LoadBalanceConfig

* [ModifyCloudNativeAPIGatewayMCPServer](https://cloud.tencent.com/document/api/1826/133149)

	* 新增入参：PreserveHost


新增数据结构：

* [AIGWAKSKCredentialConfig](https://cloud.tencent.com/document/api/1826/133161#AIGWAKSKCredentialConfig)
* [AIGWAuthModelScopeItem](https://cloud.tencent.com/document/api/1826/133161#AIGWAuthModelScopeItem)
* [AIGWBasicCredentialConfig](https://cloud.tencent.com/document/api/1826/133161#AIGWBasicCredentialConfig)
* [AIGWBearerTokenCredentialConfig](https://cloud.tencent.com/document/api/1826/133161#AIGWBearerTokenCredentialConfig)
* [AIGWCAMCredentialConfig](https://cloud.tencent.com/document/api/1826/133161#AIGWCAMCredentialConfig)
* [AIGWConsumerModelScope](https://cloud.tencent.com/document/api/1826/133161#AIGWConsumerModelScope)
* [AIGWCustomHeaderCredentialConfig](https://cloud.tencent.com/document/api/1826/133161#AIGWCustomHeaderCredentialConfig)
* [AIGWLLMHealthCheckSetting](https://cloud.tencent.com/document/api/1826/133161#AIGWLLMHealthCheckSetting)
* [AIGWLoadBalanceConfig](https://cloud.tencent.com/document/api/1826/133161#AIGWLoadBalanceConfig)
* [AIGWModelScope](https://cloud.tencent.com/document/api/1826/133161#AIGWModelScope)
* [AIGWQueryParamCredentialConfig](https://cloud.tencent.com/document/api/1826/133161#AIGWQueryParamCredentialConfig)
* [AIGWRerankMaxDocumentsConfig](https://cloud.tencent.com/document/api/1826/133161#AIGWRerankMaxDocumentsConfig)
* [AIGWSensitiveWordRoute](https://cloud.tencent.com/document/api/1826/133161#AIGWSensitiveWordRoute)
* [AIGWUpstreamTLSConfig](https://cloud.tencent.com/document/api/1826/133161#AIGWUpstreamTLSConfig)

修改数据结构：

* [AIGWMCPServer](https://cloud.tencent.com/document/api/1826/133161#AIGWMCPServer)

	* 新增成员：PreserveHost

* [AIGWMCPUpstreamInfo](https://cloud.tencent.com/document/api/1826/133161#AIGWMCPUpstreamInfo)

	* 新增成员：TLSConfig

* [AIGWMCPUpstreamInfoDetail](https://cloud.tencent.com/document/api/1826/133161#AIGWMCPUpstreamInfoDetail)

	* 新增成员：TLSConfig

* [CNAPIGwConsumer](https://cloud.tencent.com/document/api/1826/133161#CNAPIGwConsumer)

	* 新增成员：Priority, SyncStatus, SourceType, SyncedVersion

* [CNAPIGwConsumerGroup](https://cloud.tencent.com/document/api/1826/133161#CNAPIGwConsumerGroup)

	* 新增成员：SyncStatus, SourceType, SyncedVersion

* [CNAPIGwSecretKey](https://cloud.tencent.com/document/api/1826/133161#CNAPIGwSecretKey)

	* 新增成员：SyncStatus, SourceType, SyncedVersion, AKSKCredentialConfig, CAMCredentialConfig, BearerTokenCredentialConfig, BasicCredentialConfig, CustomHeaderCredentialConfig, QueryParamCredentialConfig

* [CloudNativeAPIGatewayLLMModelAPI](https://cloud.tencent.com/document/api/1826/133161#CloudNativeAPIGatewayLLMModelAPI)

	* 新增成员：MaxDocumentsConfig, SensitiveWordRoute, ConsumerGroupModelScopes, ConsumerInheritModelScope

* [CloudNativeAPIGatewayLLMModelService](https://cloud.tencent.com/document/api/1826/133161#CloudNativeAPIGatewayLLMModelService)

	* 新增成员：LoadBalanceConfig, CanPublish, PublishStatus, SyncStatus, SourceType, SyncedVersion, Status, EnableHealthCheck, HealthCheck




## 云安全一体化平台(csip) 版本：2022-11-21

### 第 102 次发布

发布时间：2026-08-25 01:40:07

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateDspmAuditFilterStrategy](https://cloud.tencent.com/document/api/664/136864)
* [CreateDspmResource](https://cloud.tencent.com/document/api/664/136863)
* [CreateDspmRiskStrategy](https://cloud.tencent.com/document/api/664/136862)
* [DeleteDspmAuditFilterStrategy](https://cloud.tencent.com/document/api/664/136861)
* [DeleteDspmCkafkaConfig](https://cloud.tencent.com/document/api/664/136856)
* [DeleteDspmRiskStrategy](https://cloud.tencent.com/document/api/664/136860)
* [DeleteDspmShareUserData](https://cloud.tencent.com/document/api/664/136844)
* [DescribeDspmAuditFilterStrategy](https://cloud.tencent.com/document/api/664/136859)
* [DescribeDspmCkafkaRouteList](https://cloud.tencent.com/document/api/664/136855)
* [DescribeDspmCkafkaTopicList](https://cloud.tencent.com/document/api/664/136854)
* [DescribeDspmLogDeliveryType](https://cloud.tencent.com/document/api/664/136853)
* [DescribeDspmLogTypeConfigList](https://cloud.tencent.com/document/api/664/136852)
* [DescribeDspmResource](https://cloud.tencent.com/document/api/664/136858)
* [DescribeDspmSessionList](https://cloud.tencent.com/document/api/664/136851)
* [DescribeDspmUserCkafkaInstanceList](https://cloud.tencent.com/document/api/664/136850)
* [ModifyDspmAuditFilterStrategy](https://cloud.tencent.com/document/api/664/136857)
* [ModifyDspmCkafkaSave](https://cloud.tencent.com/document/api/664/136849)
* [ModifyDspmCkafkaStart](https://cloud.tencent.com/document/api/664/136848)
* [ModifyDspmCkafkaStop](https://cloud.tencent.com/document/api/664/136847)
* [ModifyDspmLogDeliveryType](https://cloud.tencent.com/document/api/664/136846)
* [ModifyShareUserDspm](https://cloud.tencent.com/document/api/664/136843)
* [SendDspmCkafkaTest](https://cloud.tencent.com/document/api/664/136845)

新增数据结构：

* [CkafkaInstance](https://cloud.tencent.com/document/api/664/90825#CkafkaInstance)
* [DescribeDspmAuditFilterStrategy](https://cloud.tencent.com/document/api/664/90825#DescribeDspmAuditFilterStrategy)
* [DspmAuditSessionInfo](https://cloud.tencent.com/document/api/664/90825#DspmAuditSessionInfo)
* [LogDeliveryCkafkaConfig](https://cloud.tencent.com/document/api/664/90825#LogDeliveryCkafkaConfig)
* [LogDeliveryInfo](https://cloud.tencent.com/document/api/664/90825#LogDeliveryInfo)
* [LogDeliveryType](https://cloud.tencent.com/document/api/664/90825#LogDeliveryType)
* [RouteInfo](https://cloud.tencent.com/document/api/664/90825#RouteInfo)
* [TopicInfo](https://cloud.tencent.com/document/api/664/90825#TopicInfo)



## TDSQL-C MySQL 版(cynosdb) 版本：2019-01-07

### 第 190 次发布

发布时间：2026-08-25 01:55:28

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [TransferClusterPrepayToPostpay](https://cloud.tencent.com/document/api/1003/135370)

	* 新增入参：ClusterId

	* 新增出参：BigDealIds, TranId, DealNames, ResourceIds, ClusterIds




## 弹性 MapReduce(emr) 版本：2019-01-03

### 第 150 次发布

发布时间：2026-08-25 02:12:35

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [AutoScaleResourceConf](https://cloud.tencent.com/document/api/589/33981#AutoScaleResourceConf)

	* 新增成员：CustomNodeName

* [NodeResourceSpec](https://cloud.tencent.com/document/api/589/33981#NodeResourceSpec)

	* 新增成员：CustomNodeName

* [OperationLog](https://cloud.tencent.com/document/api/589/33981#OperationLog)

	* 新增成员：OperatorName




## Elasticsearch Service(es) 版本：2025-01-01



## Elasticsearch Service(es) 版本：2018-04-16

### 第 109 次发布

发布时间：2026-08-25 02:14:05

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [InstanceInfo](https://cloud.tencent.com/document/api/845/30634#InstanceInfo)

	* 新增成员：OldEsVip, OldEsPrivateTcpUrl




## 物联网开发平台(iotexplorer) 版本：2019-04-23

### 第 155 次发布

发布时间：2026-08-24 14:23:37

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateTWeSeePerson](https://cloud.tencent.com/document/api/1081/136720)
* [DeleteTWeSeeFace](https://cloud.tencent.com/document/api/1081/136719)
* [DeleteTWeSeePerson](https://cloud.tencent.com/document/api/1081/136718)
* [DescribeTWeSeeFace](https://cloud.tencent.com/document/api/1081/136717)
* [DescribeTWeSeePerson](https://cloud.tencent.com/document/api/1081/136716)
* [ImportTWeSeeFaces](https://cloud.tencent.com/document/api/1081/136715)
* [ListTWeSeePersons](https://cloud.tencent.com/document/api/1081/136714)
* [ModifyTWeSeeFace](https://cloud.tencent.com/document/api/1081/136713)
* [ModifyTWeSeePerson](https://cloud.tencent.com/document/api/1081/136712)

新增数据结构：

* [SeeFaceInfo](https://cloud.tencent.com/document/api/1081/34988#SeeFaceInfo)
* [SeeFaceRecognitionResult](https://cloud.tencent.com/document/api/1081/34988#SeeFaceRecognitionResult)
* [SeePersonInfo](https://cloud.tencent.com/document/api/1081/34988#SeePersonInfo)
* [SeeTaskFaceInfo](https://cloud.tencent.com/document/api/1081/34988#SeeTaskFaceInfo)
* [SeeTaskPersonInfo](https://cloud.tencent.com/document/api/1081/34988#SeeTaskPersonInfo)

修改数据结构：

* [SeeComprehensionConfig](https://cloud.tencent.com/document/api/1081/34988#SeeComprehensionConfig)

	* 新增成员：EnableFaceDetection, InputRotateDegree

* [SeeTaskInfo](https://cloud.tencent.com/document/api/1081/34988#SeeTaskInfo)

	* 新增成员：FaceRecognitionResult




## 文字识别(ocr) 版本：2018-11-19

### 第 261 次发布

发布时间：2026-08-25 02:52:04

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CropEnhanceImageOCR](https://cloud.tencent.com/document/api/866/133908)

	* 新增出参：CroppedImageUrl




## 云数据库Redis(redis) 版本：2018-04-12

### 第 109 次发布

发布时间：2026-08-25 02:58:21

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeInstancePasswordPolicy](https://cloud.tencent.com/document/api/239/136865)



## 云开发 CloudBase(tcb) 版本：2018-06-08

### 第 159 次发布

发布时间：2026-08-25 03:07:47

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [SMSCloudFunctionConfig](https://cloud.tencent.com/document/api/876/34822#SMSCloudFunctionConfig)

修改数据结构：

* [MgoKeySchema](https://cloud.tencent.com/document/api/876/34822#MgoKeySchema)

	* 新增成员：PartialFilterExpression

* [SMSProviderTemplateConfig](https://cloud.tencent.com/document/api/876/34822#SMSProviderTemplateConfig)

	* 新增成员：AuthType, CredentialAuthKeyId

* [VerificationConfig](https://cloud.tencent.com/document/api/876/34822#VerificationConfig)

	* 新增成员：CloudFunction




## 边缘安全加速平台(teo) 版本：2022-09-01

### 第 157 次发布

发布时间：2026-08-25 03:19:48

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeDDoSAttackData](https://cloud.tencent.com/document/api/1552/80660)

	* 新增入参：Filters

	* <font color="#dd0000">**修改入参**：</font>ZoneIds


修改数据结构：

* [DDoSAttackEvent](https://cloud.tencent.com/document/api/1552/80721#DDoSAttackEvent)

	* 新增成员：DDoSAttackDips




## 边缘安全加速平台(teo) 版本：2022-01-06



## TI-ONE 训练平台(tione) 版本：2021-11-11

### 第 135 次发布

发布时间：2026-08-25 03:23:37

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateModelService](https://cloud.tencent.com/document/api/851/82291)

	* 新增入参：InferTemplateId

* [ModifyModelService](https://cloud.tencent.com/document/api/851/83228)

	* 新增入参：InferTemplateId


修改数据结构：

* [ServiceInfo](https://cloud.tencent.com/document/api/851/75051#ServiceInfo)

	* 新增成员：InferTemplateId




## TI-ONE 训练平台(tione) 版本：2019-10-22



## TSF-Polaris&ZK&网关(tse) 版本：2020-12-07

### 第 113 次发布

发布时间：2026-08-25 03:32:18

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateCloudNativeAPIGatewayCertificate](https://cloud.tencent.com/document/api/1364/98591)

	* 新增入参：CertType, CertUsage

	* <font color="#dd0000">**修改入参**：</font>BindDomains

* [CreateCloudNativeAPIGatewayConsumer](https://cloud.tencent.com/document/api/1364/131970)

	* 新增入参：Priority

* [CreateCloudNativeAPIGatewayLLMModelAPI](https://cloud.tencent.com/document/api/1364/131968)

	* 新增入参：MaxDocumentsConfig, SensitiveWordRoute

* [CreateCloudNativeAPIGatewayLLMModelService](https://cloud.tencent.com/document/api/1364/131967)

	* 新增入参：LoadBalanceConfig

* [CreateCloudNativeAPIGatewaySecretKey](https://cloud.tencent.com/document/api/1364/131966)

	* 新增入参：AKSKCredentialConfig, CAMCredentialConfig, BearerTokenCredentialConfig, CustomHeaderCredentialConfig, QueryParamCredentialConfig, BasicCredentialConfig

* [DescribeCloudNativeAPIGatewayCertificates](https://cloud.tencent.com/document/api/1364/98588)

	* 新增入参：CertType, CertUsage

* [ModifyCloudNativeAPIGatewayConsumer](https://cloud.tencent.com/document/api/1364/131949)

	* 新增入参：Priority

* [ModifyCloudNativeAPIGatewayLLMModelAPI](https://cloud.tencent.com/document/api/1364/131947)

	* 新增入参：MaxDocumentsConfig, SensitiveWordRoute

* [ModifyCloudNativeAPIGatewayLLMModelService](https://cloud.tencent.com/document/api/1364/131946)

	* 新增入参：CustomProviderName, LoadBalanceConfig

* [UpdateCloudNativeAPIGatewayCertificateInfo](https://cloud.tencent.com/document/api/1364/98587)

	* <font color="#dd0000">**修改入参**：</font>BindDomains


新增数据结构：

* [AIGWAKSKCredentialConfig](https://cloud.tencent.com/document/api/1364/54942#AIGWAKSKCredentialConfig)
* [AIGWAuthModelScopeItem](https://cloud.tencent.com/document/api/1364/54942#AIGWAuthModelScopeItem)
* [AIGWBasicCredentialConfig](https://cloud.tencent.com/document/api/1364/54942#AIGWBasicCredentialConfig)
* [AIGWBearerTokenCredentialConfig](https://cloud.tencent.com/document/api/1364/54942#AIGWBearerTokenCredentialConfig)
* [AIGWCAMCredentialConfig](https://cloud.tencent.com/document/api/1364/54942#AIGWCAMCredentialConfig)
* [AIGWConsumerModelScope](https://cloud.tencent.com/document/api/1364/54942#AIGWConsumerModelScope)
* [AIGWCustomHeaderCredentialConfig](https://cloud.tencent.com/document/api/1364/54942#AIGWCustomHeaderCredentialConfig)
* [AIGWLLMHealthCheckSetting](https://cloud.tencent.com/document/api/1364/54942#AIGWLLMHealthCheckSetting)
* [AIGWLoadBalanceConfig](https://cloud.tencent.com/document/api/1364/54942#AIGWLoadBalanceConfig)
* [AIGWModelScope](https://cloud.tencent.com/document/api/1364/54942#AIGWModelScope)
* [AIGWQueryParamCredentialConfig](https://cloud.tencent.com/document/api/1364/54942#AIGWQueryParamCredentialConfig)
* [AIGWRerankMaxDocumentsConfig](https://cloud.tencent.com/document/api/1364/54942#AIGWRerankMaxDocumentsConfig)
* [AIGWSensitiveWordRoute](https://cloud.tencent.com/document/api/1364/54942#AIGWSensitiveWordRoute)

修改数据结构：

* [AIGWLogConfig](https://cloud.tencent.com/document/api/1364/54942#AIGWLogConfig)

	* 新增成员：RequestLogPayloadTruncationPolicy, ResponseLogPayloadTruncationPolicy

* [CNAPIGwConsumer](https://cloud.tencent.com/document/api/1364/54942#CNAPIGwConsumer)

	* 新增成员：Priority, SyncStatus, SourceType, SyncedVersion

* [CNAPIGwConsumerGroup](https://cloud.tencent.com/document/api/1364/54942#CNAPIGwConsumerGroup)

	* 新增成员：SyncStatus, SourceType, SyncedVersion

* [CNAPIGwSecretKey](https://cloud.tencent.com/document/api/1364/54942#CNAPIGwSecretKey)

	* 新增成员：AKSKCredentialConfig, CAMCredentialConfig, BearerTokenCredentialConfig, BasicCredentialConfig, CustomHeaderCredentialConfig, QueryParamCredentialConfig, SyncStatus, SourceType, SyncedVersion

* [CloudNativeAPIGatewayLLMModelAPI](https://cloud.tencent.com/document/api/1364/54942#CloudNativeAPIGatewayLLMModelAPI)

	* 新增成员：MaxDocumentsConfig, SensitiveWordRoute, ConsumerGroupModelScopes, ConsumerInheritModelScope

* [CloudNativeAPIGatewayLLMModelService](https://cloud.tencent.com/document/api/1364/54942#CloudNativeAPIGatewayLLMModelService)

	* 新增成员：LoadBalanceConfig, PublishStatus, CanPublish, SyncStatus, SourceType, SyncedVersion, Status, EnableHealthCheck, HealthCheck

* [KongCertificatesPreview](https://cloud.tencent.com/document/api/1364/54942#KongCertificatesPreview)

	* 新增成员：CertType, CertUsage, ReferCount




## Web 应用防火墙(waf) 版本：2018-01-25

### 第 163 次发布

发布时间：2026-08-25 03:45:52

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [EnableClientMsg](https://cloud.tencent.com/document/api/627/136867)
* [QueryClientMsg](https://cloud.tencent.com/document/api/627/136866)



## 数据开发治理平台 WeData(wedata) 版本：2025-08-06

### 第 23 次发布

发布时间：2026-08-25 03:52:48

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [WorkspaceExt](https://cloud.tencent.com/document/api/1267/123643#WorkspaceExt)

修改数据结构：

* [Project](https://cloud.tencent.com/document/api/1267/123643#Project)

	* 新增成员：WorkspaceExt

* [ProjectRequest](https://cloud.tencent.com/document/api/1267/123643#ProjectRequest)

	* 新增成员：ScheduleMode

* [TriggerTaskBrief](https://cloud.tencent.com/document/api/1267/123643#TriggerTaskBrief)

	* 新增成员：CycleType




## 数据开发治理平台 WeData(wedata) 版本：2021-08-20

### 第 202 次发布

发布时间：2026-08-25 03:48:54

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeTableContentPreview](https://cloud.tencent.com/document/api/1267/129275)

	* 新增入参：EngineTypeDetail




