# Release v1.3.208

## Agent 沙箱服务(ags) 版本：2025-09-20

### 第 31 次发布

发布时间：2026-11-05 01:04:52

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeRegistryRecordVersionList](https://cloud.tencent.com/document/api/1814/138456)

	* <font color="#dd0000">**修改入参**：</font>RecordId


新增数据结构：

* [HarnessVersionReference](https://cloud.tencent.com/document/api/1814/124823#HarnessVersionReference)
* [MetadataEntry](https://cloud.tencent.com/document/api/1814/124823#MetadataEntry)

修改数据结构：

* [CloudRecordVersion](https://cloud.tencent.com/document/api/1814/124823#CloudRecordVersion)

	* 新增成员：SupportedHarnessVersions, Metadata, Name, Description, DescriptorType




## 灾备中心(bdrc) 版本：2026-03-30

### 第 4 次发布

发布时间：2026-11-05 01:10:02

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [BindAutoBackupPolicy](https://cloud.tencent.com/document/api/1824/136822)

	* 新增入参：AutoBackupPolicyId, InstanceIds

* [CreateFileBackup](https://cloud.tencent.com/document/api/1824/136818)

	* 新增入参：ResourceType, EnableObjectAcl, EnableObjectTags

* [CreateFileRestoreTask](https://cloud.tencent.com/document/api/1824/136816)

	* 新增入参：BackupId, TargetResourceId, RestorePaths, TargetLocation

	* 新增出参：TaskId

* [DescribeBackupOverviewGeneral](https://cloud.tencent.com/document/api/1824/136805)

	* 新增出参：CFSAgentBackupOverview, COSAgentBackupOverview

* [DescribeFileBackupObjects](https://cloud.tencent.com/document/api/1824/136800)

	* 新增入参：BackupId, Path, Offset, Limit, OrderField, Order, ObjectType

	* 新增出参：DirList, FileList

* [DescribeFileBackupPlans](https://cloud.tencent.com/document/api/1824/136799)

	* 新增入参：ResourceType

* [DescribeFileBackups](https://cloud.tencent.com/document/api/1824/136798)

	* 新增入参：ResourceType

* [DescribeFileRestoreTasks](https://cloud.tencent.com/document/api/1824/136796)

	* 新增入参：ResourceType

* [ModifyFileBackupPlan](https://cloud.tencent.com/document/api/1824/136789)

	* 新增入参：EnableObjectAcl, EnableObjectTags


新增数据结构：

* [BackupResourceInfo](https://cloud.tencent.com/document/api/1824/136824#BackupResourceInfo)
* [FileInfo](https://cloud.tencent.com/document/api/1824/136824#FileInfo)

修改数据结构：

* [AutoBackupPolicy](https://cloud.tencent.com/document/api/1824/136824#AutoBackupPolicy)

	* 新增成员：BackupResourceInfos

* [BackupInfo](https://cloud.tencent.com/document/api/1824/136824#BackupInfo)

	* 新增成员：IsInventoryPending, EnableObjectAcl, EnableObjectTags, CloudResourceId, ProductType

* [BackupInstance](https://cloud.tencent.com/document/api/1824/136824#BackupInstance)

	* 新增成员：BackupSizeMb

* [PlanInfo](https://cloud.tencent.com/document/api/1824/136824#PlanInfo)

	* 新增成员：EnableObjectAcl, EnableObjectTags, ProductType

* [ProtectInstance](https://cloud.tencent.com/document/api/1824/136824#ProtectInstance)

	* 新增成员：BackupSize, ProductType

* [ProtectedResourceOverview](https://cloud.tencent.com/document/api/1824/136824#ProtectedResourceOverview)

	* 新增成员：CFSAgent, COSAgent

* [ResourcePlan](https://cloud.tencent.com/document/api/1824/136824#ResourcePlan)

	* 新增成员：EnableObjectAcl, EnableObjectTags




## 云防火墙(cfw) 版本：2019-09-04

### 第 121 次发布

发布时间：2026-11-05 01:20:14

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateAlertCenterRuleAsync](https://cloud.tencent.com/document/api/1132/134613)

	* 新增入参：ExcludeIdList




## 大数据智能体工作台DataBuddy(databuddy) 版本：2026-07-15

### 第 14 次发布

发布时间：2026-11-05 01:42:23

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [FileArrivalConfig](https://cloud.tencent.com/document/api/1835/138006#FileArrivalConfig)

修改数据结构：

* [GetWorkflowTaskRunRsp](https://cloud.tencent.com/document/api/1835/138006#GetWorkflowTaskRunRsp)

	* 新增成员：FileArrivalConfig

* [WorkflowRun](https://cloud.tencent.com/document/api/1835/138006#WorkflowRun)

	* 新增成员：FileArrivalConfig

* [WorkflowTaskRun](https://cloud.tencent.com/document/api/1835/138006#WorkflowTaskRun)

	* 新增成员：FileArrivalConfig

* [WorkflowTriggerConfiguration](https://cloud.tencent.com/document/api/1835/138006#WorkflowTriggerConfiguration)

	* 新增成员：FileArrivalConfig




## 数据湖计算 DLC(dlc) 版本：2021-01-25

### 第 188 次发布

发布时间：2026-11-05 01:45:59

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CopyJobSpec](https://cloud.tencent.com/document/api/1342/135459)

	* 新增出参：JobPackageSource, SubAccountUin, SubAccountName

* [CreateLab](https://cloud.tencent.com/document/api/1342/135449)

	* 新增入参：AutoscalerOptions

	* 新增出参：UpdateTime, StartTime, StopTime, DeleteTime, Token, AutoscalerOptions

* [GetLabDetail](https://cloud.tencent.com/document/api/1342/135447)

	* 新增出参：CreateTime, UpdateTime, Token, AutoscalerOptions

* [GetRayJob](https://cloud.tencent.com/document/api/1342/135469)

	* 新增出参：ClusterId, SubmissionId, SubAccountName

* [ListRayClusterJobs](https://cloud.tencent.com/document/api/1342/135414)

	* 新增出参：Total

* [ListRayJobs](https://cloud.tencent.com/document/api/1342/135462)

	* 新增出参：Total

* [RunJobSpec](https://cloud.tencent.com/document/api/1342/135453)

	* 新增出参：ClusterId, ClusterName, ResourcePartitionName, SubmissionId, HistoryUrl, FinishTime, ErrorMessage

* [StartLab](https://cloud.tencent.com/document/api/1342/135438)

	* 新增出参：UpdateTime, StopTime, DeleteTime, Token, AutoscalerOptions

* [StopLab](https://cloud.tencent.com/document/api/1342/135437)

	* 新增出参：CodeArchiveUrl, LabImagePullPolicy, Description, StatusMessage, UpdateTime, GroupName, GroupId, ResourceConfigId, ImagePullPolicy, StopTime, DeleteTime, AutoscalerOptions

* [SubmitRayJob](https://cloud.tencent.com/document/api/1342/140518)

	* 新增入参：ClusterId, DispatchStrategy, JobPackage, JobPackageName, JobPackageSource, JobSource, FlowId, ExecutionId

	* 新增出参：ClusterId, ResourcePartitionName, SubmissionId, SubAccountName, HistoryUrl, FinishTime, ErrorMessage, SpecId, SpecName

* [UpdateJobSpecPriority](https://cloud.tencent.com/document/api/1342/135451)

	* 新增出参：SubmissionTarget, GroupName, ClusterName, SubAccountName

* [UpdateLab](https://cloud.tencent.com/document/api/1342/135436)

	* 新增入参：Id, ResourceConfig, Catalog, ServiceTypes, AdvancedOptions, AutoscalerOptions

	* 新增出参：SubAccountUin, StatusMessage, CreateTime, UpdateTime, GroupName, Catalog, HistoryUrl, AdvancedOptions, StartTime, StopTime, DeleteTime, Tags, Token, AutoscalerOptions

* [UpdateRayJobPriority](https://cloud.tencent.com/document/api/1342/135461)

	* 新增出参：ClusterId, SubmissionId, SubAccountName, ErrorMessage, RuntimeEnv, AutoscalerOptions, SpecName, AdvancedOptions, ClusterName




## 数据传输服务(dts) 版本：2021-12-06

### 第 62 次发布

发布时间：2026-11-05 01:52:10

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [StopMigrateJob](https://cloud.tencent.com/document/api/571/82071)

	* 新增入参：MigrateAdvancedObject


修改数据结构：

* [MigrateDBItem](https://cloud.tencent.com/document/api/571/82108#MigrateDBItem)

	* 新增成员：StorageType




## 数据传输服务(dts) 版本：2018-03-30



## 云游戏(gs) 版本：2019-11-18

### 第 68 次发布

发布时间：2026-11-05 02:02:27

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateAndroidInstanceMCPToken](https://cloud.tencent.com/document/api/1162/140547)
* [DeleteAndroidInstanceMCPToken](https://cloud.tencent.com/document/api/1162/140546)
* [DescribeAndroidInstanceMCPTokens](https://cloud.tencent.com/document/api/1162/140545)
* [RenewAndroidInstanceMCPToken](https://cloud.tencent.com/document/api/1162/140544)



## 文字识别(ocr) 版本：2018-11-19

### 第 275 次发布

发布时间：2026-11-05 02:24:24

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [GenerateTextEmbedding](https://cloud.tencent.com/document/api/866/140549)
* [GenerateVLEmbedding](https://cloud.tencent.com/document/api/866/140548)

新增数据结构：

* [EmbeddingObject](https://cloud.tencent.com/document/api/866/33527#EmbeddingObject)



## 云数据库Redis(redis) 版本：2018-04-12

### 第 112 次发布

发布时间：2026-11-05 02:29:20

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeSSLStatus](https://cloud.tencent.com/document/api/239/88703)

	* 新增入参：ServiceType

	* 新增出参：CertSource, CertStatus, CertId, CertAlias, CertExpireTime, Domain, SSLMinVersion, EncryptionMode

* [OpenSSL](https://cloud.tencent.com/document/api/239/81239)

	* 新增入参：CertSource, SwitchOption, CertId, SSLMinVersion, EncryptionMode


修改数据结构：

* [InstanceEnumParam](https://cloud.tencent.com/document/api/239/20022#InstanceEnumParam)

	* 新增成员：ParamType, FeatureSupported, NewFeatureOnlineTime, Pattern, InputDescription

* [InstanceIntegerParam](https://cloud.tencent.com/document/api/239/20022#InstanceIntegerParam)

	* 新增成员：ParamType, FeatureSupported, NewFeatureOnlineTime, Pattern, InputDescription

* [InstanceMultiParam](https://cloud.tencent.com/document/api/239/20022#InstanceMultiParam)

	* 新增成员：ParamType, FeatureSupported, NewFeatureOnlineTime, Pattern, InputDescription

* [InstanceTextParam](https://cloud.tencent.com/document/api/239/20022#InstanceTextParam)

	* 新增成员：ParamType, FeatureSupported, NewFeatureOnlineTime, Pattern, InputDescription

* [LogResult](https://cloud.tencent.com/document/api/239/20022#LogResult)

	* 新增成员：Node




## TDSQL(tdmysql) 版本：2021-11-22

### 第 18 次发布

发布时间：2026-11-05 02:43:21

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [DBEngineInfo](https://cloud.tencent.com/document/api/1376/128305#DBEngineInfo)

	* 新增成员：IsSupportTDE, IsSupportPasswordComplexity

* [SlowLogData](https://cloud.tencent.com/document/api/1376/128305#SlowLogData)

	* 新增成员：ExplainTree




## TI-ONE 训练平台(tione) 版本：2021-11-11

### 第 144 次发布

发布时间：2026-11-05 02:47:22

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateRole](https://cloud.tencent.com/document/api/851/140564)
* [DeleteRole](https://cloud.tencent.com/document/api/851/140563)
* [DescribeAccessRules](https://cloud.tencent.com/document/api/851/140562)
* [DescribeRole](https://cloud.tencent.com/document/api/851/140561)
* [DescribeRoleUsers](https://cloud.tencent.com/document/api/851/140560)
* [DescribeRoles](https://cloud.tencent.com/document/api/851/140559)
* [DescribeWorkspace](https://cloud.tencent.com/document/api/851/140558)
* [DescribeWorkspaceMembers](https://cloud.tencent.com/document/api/851/140557)
* [DescribeWorkspaceQuotaUsage](https://cloud.tencent.com/document/api/851/140551)
* [DescribeWorkspaceQuotas](https://cloud.tencent.com/document/api/851/140550)
* [ModifyRole](https://cloud.tencent.com/document/api/851/140556)
* [ModifyRoleAccessRules](https://cloud.tencent.com/document/api/851/140555)
* [ModifyRoleUserGrant](https://cloud.tencent.com/document/api/851/140554)
* [ModifyUserGrant](https://cloud.tencent.com/document/api/851/140553)
* [ModifyWorkspaceMembers](https://cloud.tencent.com/document/api/851/140552)

新增数据结构：

* [AccessRule](https://cloud.tencent.com/document/api/851/75051#AccessRule)
* [DataSourceBrief](https://cloud.tencent.com/document/api/851/75051#DataSourceBrief)
* [MemberInfo](https://cloud.tencent.com/document/api/851/75051#MemberInfo)
* [ResourcePermission](https://cloud.tencent.com/document/api/851/75051#ResourcePermission)
* [RoleBriefInfo](https://cloud.tencent.com/document/api/851/75051#RoleBriefInfo)
* [RoleInfo](https://cloud.tencent.com/document/api/851/75051#RoleInfo)
* [UserBriefInfo](https://cloud.tencent.com/document/api/851/75051#UserBriefInfo)
* [UserRoleInfo](https://cloud.tencent.com/document/api/851/75051#UserRoleInfo)
* [WorkspaceTaskInfo](https://cloud.tencent.com/document/api/851/75051#WorkspaceTaskInfo)



## TI-ONE 训练平台(tione) 版本：2019-10-22



## 音视频终端引擎(vcube) 版本：2022-04-10

### 第 10 次发布

发布时间：2026-11-05 02:56:15

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [ApplicationInfo](https://cloud.tencent.com/document/api/1449/113537#ApplicationInfo)

	* 新增成员：LinuxProcessName




## 私有网络(vpc) 版本：2017-03-12

### 第 315 次发布

发布时间：2026-11-05 02:59:19

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DeleteHaVip](https://cloud.tencent.com/document/api/215/30651)

	* 新增入参：SkipTerminationProtection

* [ModifyHaVipAttribute](https://cloud.tencent.com/document/api/215/30647)

	* 新增入参：TerminationProtection, TrafficProtection




## 元宝搜索(wsa) 版本：2025-05-08

### 第 6 次发布

发布时间：2026-11-05 03:10:39

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [SearchPro](https://cloud.tencent.com/document/api/1806/121811)

	* <font color="#dd0000">**修改入参**：</font>Cnt




