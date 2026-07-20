# Release v1.3.139

## 腾讯云智能体开发平台(adp) 版本：2026-05-20

### 第 10 次发布

发布时间：2026-07-21 01:08:24

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeAccountList](https://cloud.tencent.com/document/api/1759/134611)
* [DescribeAuditLogList](https://cloud.tencent.com/document/api/1759/134610)
* [DescribeAuditLogMeta](https://cloud.tencent.com/document/api/1759/134609)

新增数据结构：

* [AccountInfo](https://cloud.tencent.com/document/api/1759/132545#AccountInfo)
* [AuditLog](https://cloud.tencent.com/document/api/1759/132545#AuditLog)
* [AuditLogMetaField](https://cloud.tencent.com/document/api/1759/132545#AuditLogMetaField)

### 第 9 次发布

发布时间：2026-07-20 14:34:22

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreatePlugin](https://cloud.tencent.com/document/api/1759/133501)

	* 新增入参：LoginUin, LoginSubAccountUin

	* <font color="#dd0000">**修改入参**：</font>ToolList

* [CreateRelease](https://cloud.tencent.com/document/api/1759/132538)

	* 新增入参：AppShareAccessControl, CorpShareConfig

* [DeleteApp](https://cloud.tencent.com/document/api/1759/132505)

	* 新增入参：Reason

* [DeletePlugin](https://cloud.tencent.com/document/api/1759/133500)

	* 新增入参：LoginUin, LoginSubAccountUin

* [ModifyApp](https://cloud.tencent.com/document/api/1759/132502)

	* <font color="#dd0000">**删除入参**：</font>ShareConfig

* [ModifyPlugin](https://cloud.tencent.com/document/api/1759/133498)

	* 新增入参：LoginUin, LoginSubAccountUin


新增数据结构：

* [ClawAgentAgentTeamConfig](https://cloud.tencent.com/document/api/1759/132545#ClawAgentAgentTeamConfig)
* [ClawAgentLongMemoryConfig](https://cloud.tencent.com/document/api/1759/132545#ClawAgentLongMemoryConfig)
* [CorpShareConfig](https://cloud.tencent.com/document/api/1759/132545#CorpShareConfig)

修改数据结构：

* [App](https://cloud.tencent.com/document/api/1759/132545#App)

	* 新增成员：CorpShareConfig

* [AppShareAccessControl](https://cloud.tencent.com/document/api/1759/132545#AppShareAccessControl)

	* <font color="#dd0000">**修改成员**：</font>AccessType, Enabled

* [AppShareWhitelistItem](https://cloud.tencent.com/document/api/1759/132545#AppShareWhitelistItem)

	* <font color="#dd0000">**修改成员**：</font>Type

* [ClawAgentConfig](https://cloud.tencent.com/document/api/1759/132545#ClawAgentConfig)

	* 新增成员：AgentTeamConfig, LongMemoryConfig

* [ReleaseSummary](https://cloud.tencent.com/document/api/1759/132545#ReleaseSummary)

	* 新增成员：AppShareAccessControl, CorpShareConfig

* [Variable](https://cloud.tencent.com/document/api/1759/132545#Variable)

	* 新增成员：EnableEndpoints, EndpointList




## 云数据库 MySQL(cdb) 版本：2017-03-20

### 第 224 次发布

发布时间：2026-07-21 01:21:13

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateDBInstance](https://cloud.tencent.com/document/api/236/15871)

	* 新增入参：DiskEncryption

* [CreateDBInstanceHour](https://cloud.tencent.com/document/api/236/15865)

	* 新增入参：DiskEncryption




## 腾讯云数据仓库 TCHouse-D(cdwdoris) 版本：2021-12-28

### 第 63 次发布

发布时间：2026-07-21 01:24:26

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [WorkloadGroupConfig](https://cloud.tencent.com/document/api/1387/102385#WorkloadGroupConfig)

	* 新增成员：MinCpuPercent, MinMemoryPercent, MaxConcurrencyNum, MaxQueueSize, QueueTimeout




## 云防火墙(cfw) 版本：2019-09-04

### 第 108 次发布

发布时间：2026-07-21 01:26:34

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateAlertCenterRuleAsync](https://cloud.tencent.com/document/api/1132/134613)
* [ModifyIsolateTable](https://cloud.tencent.com/document/api/1132/134612)

修改接口：

* [DescribeAssetSync](https://cloud.tencent.com/document/api/1132/97917)

	* 新增出参：CVMCount

* [DescribeCfwAlerts](https://cloud.tencent.com/document/api/1132/134463)

	* 新增入参：StartTime, EndTime, Level, Direction, ActionStatus, KillChain, AttackResult, Strategy, EventName, EventId, SrcIp, DstIp, InstanceId, OrderBy, Order

* [DescribeCfwAssets](https://cloud.tencent.com/document/api/1132/134461)

	* 新增入参：AssetType, Ip, InstanceId, VpcId, SubnetId, InstanceType, NextToken

* [DescribeCfwRiskOverview](https://cloud.tencent.com/document/api/1132/134459)

	* 新增入参：StartTime, EndTime

* [DescribeCfwRuleOptimization](https://cloud.tencent.com/document/api/1132/134458)

	* 新增入参：RuleType, Dimensions

* [DescribeCfwRules](https://cloud.tencent.com/document/api/1132/134457)

	* 新增入参：Enabled, IncludeDisabled, RuleUuid, Protocol, SrcIp, DstIp, Description, Keyword, InstanceId, ExpandNames




## 云安全一体化平台(csip) 版本：2022-11-21

### 第 92 次发布

发布时间：2026-07-21 01:36:56

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [Machine](https://cloud.tencent.com/document/api/664/90825#Machine)

	* 新增成员：MigrationRequired, IsSupportXSPM, CanUnbind




## 腾讯云数据分析智能体(dataagent) 版本：2025-05-13

### 第 21 次发布

发布时间：2026-07-21 01:48:39

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [QueryModels](https://cloud.tencent.com/document/api/1800/134615)

新增数据结构：

* [ModelList](https://cloud.tencent.com/document/api/1800/125016#ModelList)



## 数据湖计算 DLC(dlc) 版本：2021-01-25

### 第 164 次发布

发布时间：2026-07-21 01:51:53

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeMCPTask](https://cloud.tencent.com/document/api/1342/134617)
* [DescribeMCPTaskResult](https://cloud.tencent.com/document/api/1342/134616)



## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 309 次发布

发布时间：2026-07-21 02:02:56

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateOrganizationAuthUrl](https://cloud.tencent.com/document/api/1323/105134)

	* 新增入参：BizLicenseSame




## 全球加速(ga2) 版本：2025-01-15

### 第 11 次发布

发布时间：2026-07-21 02:07:17

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateGlobalAcceleratorAclPolicy](https://cloud.tencent.com/document/api/1817/134623)
* [CreateGlobalAcceleratorAclRule](https://cloud.tencent.com/document/api/1817/134622)
* [DeleteGlobalAcceleratorAclPolicy](https://cloud.tencent.com/document/api/1817/134621)
* [DeleteGlobalAcceleratorAclRule](https://cloud.tencent.com/document/api/1817/134620)
* [ModifyGlobalAcceleratorAclPolicy](https://cloud.tencent.com/document/api/1817/134619)
* [ModifyGlobalAcceleratorAclRule](https://cloud.tencent.com/document/api/1817/134618)

新增数据结构：

* [AclEntries](https://cloud.tencent.com/document/api/1817/130045#AclEntries)



## 物联网开发平台(iotexplorer) 版本：2019-04-23

### 第 150 次发布

发布时间：2026-07-21 02:13:36

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [ADPConfig](https://cloud.tencent.com/document/api/1081/34988#ADPConfig)

修改数据结构：

* [TalkLLMConfig](https://cloud.tencent.com/document/api/1081/34988#TalkLLMConfig)

	* 新增成员：ADP




## 媒体处理(mps) 版本：2019-06-12

### 第 220 次发布

发布时间：2026-07-21 02:29:55

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateVideoRedrawTask](https://cloud.tencent.com/document/api/862/134232)

	* 新增入参：TaskInfo


新增数据结构：

* [VideoRedrawTaskInfo](https://cloud.tencent.com/document/api/862/37615#VideoRedrawTaskInfo)



## 边缘安全加速平台(teo) 版本：2022-09-01

### 第 151 次发布

发布时间：2026-07-21 02:56:43

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateConfigGroupVersion](https://cloud.tencent.com/document/api/1552/101867)

	* 新增入参：SourceVersion


修改数据结构：

* [ConfigGroupVersionInfo](https://cloud.tencent.com/document/api/1552/80721#ConfigGroupVersionInfo)

	* 新增成员：SourceVersion




## 边缘安全加速平台(teo) 版本：2022-01-06



## Web 应用防火墙(waf) 版本：2018-01-25

### 第 161 次发布

发布时间：2026-07-21 03:19:33

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [AddBatchCustomRule](https://cloud.tencent.com/document/api/627/130908)

	* 新增入参：GroupIds


新增数据结构：

* [ApiEventSample](https://cloud.tencent.com/document/api/627/53609#ApiEventSample)

修改数据结构：

* [ApiSecAttackSource](https://cloud.tencent.com/document/api/627/53609#ApiSecAttackSource)

	* 新增成员：EventDescription, EventDescriptionEng, Sample

* [BatchCustomRuleListItem](https://cloud.tencent.com/document/api/627/53609#BatchCustomRuleListItem)

	* 新增成员：GroupIds

* [ClbObject](https://cloud.tencent.com/document/api/627/53609#ClbObject)

	* 新增成员：Note

* [IntentDetectResult](https://cloud.tencent.com/document/api/627/53609#IntentDetectResult)

	* 新增成员：Category




