# Release v1.3.186

## Agent 沙箱服务(ags) 版本：2025-09-20

### 第 25 次发布

发布时间：2026-09-25 01:08:26

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ApproveRegistryRecord](https://cloud.tencent.com/document/api/1814/138463)

	* 新增入参：RegistryId, RecordId, VersionId, Comment

	* 新增出参：Version

* [CancelRegistryRecord](https://cloud.tencent.com/document/api/1814/138462)

	* 新增入参：RegistryId, RecordId, VersionId, Comment

	* 新增出参：Version

* [CreateRegistry](https://cloud.tencent.com/document/api/1814/138469)

	* 新增入参：Name, Description, ApprovalMode, Tags

	* 新增出参：RegistryId, Registry

* [CreateRegistryRecord](https://cloud.tencent.com/document/api/1814/138461)

	* 新增入参：RegistryId, Name, DescriptorType, Description, VersionName, MCPSource, AgentSource, SkillSource, CustomDescriptors

	* 新增出参：RecordId, Record, Version, UploadURL, ExpireTime, ContentStatus

* [DeleteRegistry](https://cloud.tencent.com/document/api/1814/138468)

	* 新增入参：RegistryId

* [DeleteRegistryRecord](https://cloud.tencent.com/document/api/1814/138460)

	* 新增入参：RegistryId, RecordId, VersionId, Reason

* [DescribeRegistry](https://cloud.tencent.com/document/api/1814/138467)

	* 新增入参：RegistryId

	* 新增出参：Registry

* [DescribeRegistryAuditLogList](https://cloud.tencent.com/document/api/1814/138459)

	* 新增入参：RegistryId, RecordId, VersionId, ActionFilter, Actor, StartTime, EndTime, Offset, Limit

	* 新增出参：AuditLogSet, TotalCount

* [DescribeRegistryList](https://cloud.tencent.com/document/api/1814/138466)

	* 新增入参：Offset, Limit, Filters

	* 新增出参：RegistrySet, TotalCount

* [DescribeRegistryRecord](https://cloud.tencent.com/document/api/1814/138458)

	* 新增入参：RegistryId, RecordId, VersionId, Label

	* 新增出参：Record, Version, ResolvedBy, ResolvedLabel

* [DescribeRegistryRecordList](https://cloud.tencent.com/document/api/1814/138457)

	* 新增入参：RegistryId, Offset, Limit, Filters

	* 新增出参：RecordSet, TotalCount

* [DescribeRegistryRecordVersionList](https://cloud.tencent.com/document/api/1814/138456)

	* 新增入参：RegistryId, RecordId, Offset, Limit, Filters

	* 新增出参：VersionSet, TotalCount

* [GetSkillPackageDownloadURL](https://cloud.tencent.com/document/api/1814/138455)

	* 新增入参：RegistryId, RecordId, VersionId, Label

	* 新增出参：DownloadURL, ExpireTime, SHA256, ResolvedVersionId

* [GetSkillPackageUploadURL](https://cloud.tencent.com/document/api/1814/138454)

	* 新增入参：RegistryId, RecordId, VersionId

	* 新增出参：Version, UploadURL, ContentStatus, ExpireTime

* [PreviewRegistryRecord](https://cloud.tencent.com/document/api/1814/138453)

	* 新增入参：RegistryId, RecordId, VersionId, Label

	* 新增出参：PreviewResult, ResolvedVersionId

* [RejectRegistryRecord](https://cloud.tencent.com/document/api/1814/138452)

	* 新增入参：RegistryId, RecordId, VersionId, Comment

	* 新增出参：Version

* [SyncRegistryRecord](https://cloud.tencent.com/document/api/1814/138451)

	* 新增入参：RegistryId, RecordId, VersionId, Label, ChangeLog

	* 新增出参：SyncStatus, ResolvedVersionId, CreatedVersion, Record, LastSyncTime, ErrorCode, ErrorMessage

* [UpdateRegistry](https://cloud.tencent.com/document/api/1814/138465)

	* 新增入参：RegistryId, Description

	* 新增出参：Registry

* [UpdateRegistryRecord](https://cloud.tencent.com/document/api/1814/138450)

	* 新增入参：RegistryId, RecordId, Description, VersionName, ChangeLog, MCPSource, AgentSource, SkillSource, CustomDescriptors, LabelMutations

	* 新增出参：Record, Version, UploadURL, ExpireTime, ContentStatus


新增数据结构：

* [CloudAgentSourceInput](https://cloud.tencent.com/document/api/1814/124823#CloudAgentSourceInput)
* [CloudAuditLog](https://cloud.tencent.com/document/api/1814/124823#CloudAuditLog)
* [CloudFilter](https://cloud.tencent.com/document/api/1814/124823#CloudFilter)
* [CloudMCPSourceInput](https://cloud.tencent.com/document/api/1814/124823#CloudMCPSourceInput)
* [CloudRecord](https://cloud.tencent.com/document/api/1814/124823#CloudRecord)
* [CloudRecordLabelMutation](https://cloud.tencent.com/document/api/1814/124823#CloudRecordLabelMutation)
* [CloudRecordVersion](https://cloud.tencent.com/document/api/1814/124823#CloudRecordVersion)
* [CloudRegistry](https://cloud.tencent.com/document/api/1814/124823#CloudRegistry)
* [CloudSkillSourceInput](https://cloud.tencent.com/document/api/1814/124823#CloudSkillSourceInput)
* [CloudTag](https://cloud.tencent.com/document/api/1814/124823#CloudTag)
* [CloudVersionApprovalAction](https://cloud.tencent.com/document/api/1814/124823#CloudVersionApprovalAction)



## 应用性能监控(apm) 版本：2021-06-22

### 第 70 次发布

发布时间：2026-09-25 01:12:40

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyApmInstance](https://cloud.tencent.com/document/api/1463/89002)

	* 新增入参：TokenDisplayProtection


修改数据结构：

* [ApmInstanceDetail](https://cloud.tencent.com/document/api/1463/64927#ApmInstanceDetail)

	* 新增成员：TokenDisplayProtection




## 访问管理(cam) 版本：2019-01-16

### 第 73 次发布

发布时间：2026-09-25 01:18:55

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [AddUser](https://cloud.tencent.com/document/api/598/34595)

	* 新增出参：PhoneNumVerifyLink

* [CreateMessageReceiver](https://cloud.tencent.com/document/api/598/96136)

	* 新增出参：PhoneNumVerifyLink




## 负载均衡(clb) 版本：2018-03-17

### 第 168 次发布

发布时间：2026-09-25 01:30:48

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateModelRouter](https://cloud.tencent.com/document/api/214/133217)

	* 新增入参：VideoConfig, RerankConfig, DecisionsConfig

* [DescribeModelAssociations](https://cloud.tencent.com/document/api/214/133659)

	* 新增入参：Capabilities

* [ModifyModelRouterAttributes](https://cloud.tencent.com/document/api/214/133203)

	* 新增入参：VideoConfig, RerankConfig, DecisionsConfig


新增数据结构：

* [DecisionsConfig](https://cloud.tencent.com/document/api/214/30694#DecisionsConfig)
* [RerankConfig](https://cloud.tencent.com/document/api/214/30694#RerankConfig)
* [VideoConfig](https://cloud.tencent.com/document/api/214/30694#VideoConfig)

修改数据结构：

* [ModelRouterDetail](https://cloud.tencent.com/document/api/214/30694#ModelRouterDetail)

	* 新增成员：LoadBalancerId, VideoConfig, RerankConfig, DecisionsConfig




## 日志服务(cls) 版本：2020-10-16

### 第 182 次发布

发布时间：2026-09-25 01:33:11

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeKafkaConsumer](https://cloud.tencent.com/document/api/614/95719)

	* 新增出参：EnableInternetConsume, EnableIntranetConsume

* [ModifyKafkaConsumer](https://cloud.tencent.com/document/api/614/95720)

	* 新增入参：EnableInternetConsume, EnableIntranetConsume

* [OpenKafkaConsumer](https://cloud.tencent.com/document/api/614/72339)

	* 新增入参：EnableInternetConsume, EnableIntranetConsume




## 主机安全(cwp) 版本：2018-02-28

### 第 170 次发布

发布时间：2026-09-25 01:46:18

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [RecentLoginItem](https://cloud.tencent.com/document/api/296/19867#RecentLoginItem)

修改数据结构：

* [HostLoginList](https://cloud.tencent.com/document/api/296/19867#HostLoginList)

	* 新增成员：HitRule, HitRuleName, AlertCount, FirstDiscoverTime, LastDiscoverTime, HarmDescribe, SuggestScheme, RecentLoginList




## 腾讯云数据分析智能体(dataagent) 版本：2025-05-13

### 第 25 次发布

发布时间：2026-09-25 01:53:14

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [Thinking](https://cloud.tencent.com/document/api/1800/125016#Thinking)

修改数据结构：

* [ModelList](https://cloud.tencent.com/document/api/1800/125016#ModelList)

	* 新增成员：DisplayName, Description, ContextWindow, IconUrl, CreditMultiplier, Thinking




## 数据库智能管家 DBbrain(dbbrain) 版本：2021-05-27

### 第 64 次发布

发布时间：2026-09-25 01:54:12

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [SlowLogInfoItem](https://cloud.tencent.com/document/api/1130/57812#SlowLogInfoItem)

	* 新增成员：ClientAppName, ClientHostName




## 数据库智能管家 DBbrain(dbbrain) 版本：2019-10-16



## 数据湖计算 DLC(dlc) 版本：2021-01-25

### 第 183 次发布

发布时间：2026-09-25 01:56:53

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateJob](https://cloud.tencent.com/document/api/1342/138819)
* [CreateJobDefinition](https://cloud.tencent.com/document/api/1342/138818)
* [CreateJobFromDefinition](https://cloud.tencent.com/document/api/1342/138817)
* [DescribeCatalogTableInfo](https://cloud.tencent.com/document/api/1342/138724)
* [DescribeCatalogTableNames](https://cloud.tencent.com/document/api/1342/138723)
* [DescribeCatalogTableNamesPage](https://cloud.tencent.com/document/api/1342/138722)
* [DescribeCatalogs](https://cloud.tencent.com/document/api/1342/138721)
* [DescribeJobDefinitionDetail](https://cloud.tencent.com/document/api/1342/138816)
* [DescribeJobDefinitions](https://cloud.tencent.com/document/api/1342/138815)
* [DescribeJobDetail](https://cloud.tencent.com/document/api/1342/138814)
* [DescribeJobList](https://cloud.tencent.com/document/api/1342/138813)
* [DescribeJobLog](https://cloud.tencent.com/document/api/1342/138812)
* [DescribeJobResult](https://cloud.tencent.com/document/api/1342/138811)
* [DescribeSchemaNamesPage](https://cloud.tencent.com/document/api/1342/138720)
* [DescribeWarehouses](https://cloud.tencent.com/document/api/1342/138808)
* [ModifyJobDefinition](https://cloud.tencent.com/document/api/1342/138810)

新增数据结构：

* [Audit](https://cloud.tencent.com/document/api/1342/53778#Audit)
* [BucketPartitioning](https://cloud.tencent.com/document/api/1342/53778#BucketPartitioning)
* [CatalogConfig](https://cloud.tencent.com/document/api/1342/53778#CatalogConfig)
* [CatalogTaleInfo](https://cloud.tencent.com/document/api/1342/53778#CatalogTaleInfo)
* [ClsLogEntry](https://cloud.tencent.com/document/api/1342/53778#ClsLogEntry)
* [ColumnInfo](https://cloud.tencent.com/document/api/1342/53778#ColumnInfo)
* [ConnectionConfig](https://cloud.tencent.com/document/api/1342/53778#ConnectionConfig)
* [DlcConnection](https://cloud.tencent.com/document/api/1342/53778#DlcConnection)
* [DorisConnection](https://cloud.tencent.com/document/api/1342/53778#DorisConnection)
* [HiveConnection](https://cloud.tencent.com/document/api/1342/53778#HiveConnection)
* [IndexInfo](https://cloud.tencent.com/document/api/1342/53778#IndexInfo)
* [JobBriefInfo](https://cloud.tencent.com/document/api/1342/53778#JobBriefInfo)
* [JobDefinitionItemInfo](https://cloud.tencent.com/document/api/1342/53778#JobDefinitionItemInfo)
* [LakeHouseConnection](https://cloud.tencent.com/document/api/1342/53778#LakeHouseConnection)
* [ListPartition](https://cloud.tencent.com/document/api/1342/53778#ListPartition)
* [ListPartitioning](https://cloud.tencent.com/document/api/1342/53778#ListPartitioning)
* [Literal](https://cloud.tencent.com/document/api/1342/53778#Literal)
* [MysqlConnection](https://cloud.tencent.com/document/api/1342/53778#MysqlConnection)
* [NameIdentifier](https://cloud.tencent.com/document/api/1342/53778#NameIdentifier)
* [Partitioning](https://cloud.tencent.com/document/api/1342/53778#Partitioning)
* [PostgreSQLConnection](https://cloud.tencent.com/document/api/1342/53778#PostgreSQLConnection)
* [RangePartition](https://cloud.tencent.com/document/api/1342/53778#RangePartition)
* [RangePartitioning](https://cloud.tencent.com/document/api/1342/53778#RangePartitioning)
* [RangerConnection](https://cloud.tencent.com/document/api/1342/53778#RangerConnection)
* [ResultColumn](https://cloud.tencent.com/document/api/1342/53778#ResultColumn)
* [ResultRow](https://cloud.tencent.com/document/api/1342/53778#ResultRow)
* [SingleFieldPartitioning](https://cloud.tencent.com/document/api/1342/53778#SingleFieldPartitioning)
* [TruncatePartitioning](https://cloud.tencent.com/document/api/1342/53778#TruncatePartitioning)
* [VolumeConnection](https://cloud.tencent.com/document/api/1342/53778#VolumeConnection)
* [WarehouseInfo](https://cloud.tencent.com/document/api/1342/53778#WarehouseInfo)



## 高性能应用服务(hai) 版本：2023-08-12

### 第 31 次发布

发布时间：2026-09-25 02:16:38

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeDeployTemplates](https://cloud.tencent.com/document/api/1721/129386)

	* 新增入参：ServiceId

* [DescribeModels](https://cloud.tencent.com/document/api/1721/129385)

	* 新增入参：ServiceId




## iOA 零信任安全管理系统(ioa) 版本：2022-06-01

### 第 46 次发布

发布时间：2026-09-25 02:18:30

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeDeviceSecurityInfo](https://cloud.tencent.com/document/api/1092/138821)

修改接口：

* [BindVirtualAccounts](https://cloud.tencent.com/document/api/1092/138748)

	* 新增入参：AccountUserList

	* 新增出参：Data

* [UnbindVirtualAccounts](https://cloud.tencent.com/document/api/1092/138746)

	* 新增入参：AccountUserList

	* 新增出参：Data


新增数据结构：

* [AccountUserIdItem](https://cloud.tencent.com/document/api/1092/102488#AccountUserIdItem)
* [BindVirtualAccountData](https://cloud.tencent.com/document/api/1092/102488#BindVirtualAccountData)
* [BindVirtualAccountResultData](https://cloud.tencent.com/document/api/1092/102488#BindVirtualAccountResultData)
* [DescribeDeviceSecurityInfoData](https://cloud.tencent.com/document/api/1092/102488#DescribeDeviceSecurityInfoData)
* [UnbindVirtualAccountData](https://cloud.tencent.com/document/api/1092/102488#UnbindVirtualAccountData)

### 第 45 次发布

发布时间：2026-09-24 12:03:54

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [BindVirtualAccounts](https://cloud.tencent.com/document/api/1092/138748)
* [DescribeProfileFieldsMenu](https://cloud.tencent.com/document/api/1092/138750)
* [DescribeVirtualAccounts](https://cloud.tencent.com/document/api/1092/138747)
* [UnbindVirtualAccounts](https://cloud.tencent.com/document/api/1092/138746)

新增数据结构：

* [DescribeAccountAccountGroupsData](https://cloud.tencent.com/document/api/1092/102488#DescribeAccountAccountGroupsData)
* [DescribeProfileFieldsRspData](https://cloud.tencent.com/document/api/1092/102488#DescribeProfileFieldsRspData)
* [DescribeVirtualAccountsData](https://cloud.tencent.com/document/api/1092/102488#DescribeVirtualAccountsData)
* [DescribeVirtualAccountsPageData](https://cloud.tencent.com/document/api/1092/102488#DescribeVirtualAccountsPageData)
* [DeviceProfile](https://cloud.tencent.com/document/api/1092/102488#DeviceProfile)
* [OptionsItem](https://cloud.tencent.com/document/api/1092/102488#OptionsItem)
* [ProfileFieldItem](https://cloud.tencent.com/document/api/1092/102488#ProfileFieldItem)
* [ProfileTips](https://cloud.tencent.com/document/api/1092/102488#ProfileTips)

修改数据结构：

* [DeviceDetail](https://cloud.tencent.com/document/api/1092/102488#DeviceDetail)

	* 新增成员：Profiles, InstallationStatus




## 物联网开发平台(iotexplorer) 版本：2019-04-23

### 第 160 次发布

发布时间：2026-09-24 22:23:10

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [BatchPublishMessage](https://cloud.tencent.com/document/api/1081/138806)

修改接口：

* [BatchUpdateFirmware](https://cloud.tencent.com/document/api/1081/123152)

	* 新增入参：EndTime, StartTime

* [DescribeCloudStorageEventsByTWeSeePerson](https://cloud.tencent.com/document/api/1081/138360)

	* 新增入参：StartTime, EndTime

	* 新增出参：VideoURL

* [DescribeFirmwareTask](https://cloud.tencent.com/document/api/1081/53875)

	* 新增出参：EndTime, StartTime


新增数据结构：

* [DeviceResult](https://cloud.tencent.com/document/api/1081/34988#DeviceResult)
* [SeeExtendedOutput](https://cloud.tencent.com/document/api/1081/34988#SeeExtendedOutput)
* [SeeExtendedOutputPrompt](https://cloud.tencent.com/document/api/1081/34988#SeeExtendedOutputPrompt)

修改数据结构：

* [SeeComprehensionConfig](https://cloud.tencent.com/document/api/1081/34988#SeeComprehensionConfig)

	* 新增成员：EnableExtendedOutput, ExtendedOutputPrompts

* [SeeComprehensionResult](https://cloud.tencent.com/document/api/1081/34988#SeeComprehensionResult)

	* 新增成员：ExtendedOutput

* [SeeEventIdFilterConfig](https://cloud.tencent.com/document/api/1081/34988#SeeEventIdFilterConfig)

	* 新增成员：TriggerAt




## 腾讯云数据库 AI 服务(tdai) 版本：2025-07-17

### 第 20 次发布

发布时间：2026-09-25 03:02:21

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [AgentMemInfo](https://cloud.tencent.com/document/api/1813/123239#AgentMemInfo)

修改数据结构：

* [AgentInstance](https://cloud.tencent.com/document/api/1813/123239#AgentInstance)

	* 新增成员：AgentMem




## 边缘安全加速平台(teo) 版本：2022-09-01

### 第 166 次发布

发布时间：2026-09-25 03:04:32

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateInferenceDomain](https://cloud.tencent.com/document/api/1552/138824)
* [DescribeInferenceDomains](https://cloud.tencent.com/document/api/1552/138823)
* [OperateInferenceDomain](https://cloud.tencent.com/document/api/1552/138822)

新增数据结构：

* [HostCertInfo](https://cloud.tencent.com/document/api/1552/80721#HostCertInfo)
* [HostsCertificate](https://cloud.tencent.com/document/api/1552/80721#HostsCertificate)
* [InferenceDomain](https://cloud.tencent.com/document/api/1552/80721#InferenceDomain)
* [ZoneCustomVariables](https://cloud.tencent.com/document/api/1552/80721#ZoneCustomVariables)

修改数据结构：

* [ZoneFullConfig](https://cloud.tencent.com/document/api/1552/80721#ZoneFullConfig)

	* 新增成员：ZoneCustomVariables




## 边缘安全加速平台(teo) 版本：2022-01-06



## TI-ONE 训练平台(tione) 版本：2021-11-11

### 第 137 次发布

发布时间：2026-09-25 03:08:26

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [EnvVar](https://cloud.tencent.com/document/api/851/75051#EnvVar)

	* 新增成员：IsPrivate




## TI-ONE 训练平台(tione) 版本：2019-10-22



## Web 应用防火墙(waf) 版本：2018-01-25

### 第 166 次发布

发布时间：2026-09-25 03:27:08

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**预下线接口**：</font>

* ModifyInstanceElasticMode



## WorkBuddy Enterprise(workbuddyenterprise) 版本：2026-07-09

### 第 3 次发布

发布时间：2026-09-25 03:35:43

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeUserAccessToken](https://cloud.tencent.com/document/api/1831/138826)



