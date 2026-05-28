# Release v1.3.106

## 云防火墙(cfw) 版本：2019-09-04

### 第 101 次发布

发布时间：2026-05-29 01:25:13

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [CcnAssociatedInstance](https://cloud.tencent.com/document/api/1132/49071#CcnAssociatedInstance)

	* 新增成员：IsCrossInstance




## 边缘计算机器(ecm) 版本：2019-07-19

### 第 78 次发布

发布时间：2026-05-28 12:13:23

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除接口**：</font>

* AttachDisks
* CreateDisks
* DeleteSnapshots
* DescribeDisks
* DescribeSnapshots
* DetachDisks
* TerminateDisks

<font color="#dd0000">**删除数据结构**：</font>

* Disk
* DiskChargePrepaid
* Placement
* Snapshot

修改数据结构：

* [Position](https://cloud.tencent.com/document/api/1108/42574#Position)

	* <font color="#dd0000">**修改成员**：</font>Ipv6Supported




## 事件中心(evt) 版本：2025-02-17

### 第 7 次发布

发布时间：2026-05-29 02:03:11

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [PutMessage](https://cloud.tencent.com/document/api/1808/126898)

	* 新增入参：PluginId

	* <font color="#dd0000">**修改入参**：</font>EventId




## 智能导诊(ig) 版本：2021-05-18

### 第 2 次发布

发布时间：2026-05-28 11:33:44

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [GetLLMDiagnosisDrug](https://cloud.tencent.com/document/api/1273/128049)
* [GetLLMDiagnosisDrugChat](https://cloud.tencent.com/document/api/1273/128048)
* [GetLLMDiagnosisHealth](https://cloud.tencent.com/document/api/1273/128047)
* [GetLLMReportInterpretation](https://cloud.tencent.com/document/api/1273/128046)
* [QueryDrugInstructions](https://cloud.tencent.com/document/api/1273/128045)

新增数据结构：

* [DrugCardInfo](https://cloud.tencent.com/document/api/1273/128050#DrugCardInfo)
* [DrugInstructionInfo](https://cloud.tencent.com/document/api/1273/128050#DrugInstructionInfo)
* [GuessQuestion](https://cloud.tencent.com/document/api/1273/128050#GuessQuestion)
* [HealthFollowUpQuestion](https://cloud.tencent.com/document/api/1273/128050#HealthFollowUpQuestion)
* [HighlightWordInfo](https://cloud.tencent.com/document/api/1273/128050#HighlightWordInfo)
* [LLMDiagnosisDrugData](https://cloud.tencent.com/document/api/1273/128050#LLMDiagnosisDrugData)
* [LLMDiagnosisHealthData](https://cloud.tencent.com/document/api/1273/128050#LLMDiagnosisHealthData)
* [LLMReportInterpretationData](https://cloud.tencent.com/document/api/1273/128050#LLMReportInterpretationData)
* [ReferResourceInfo](https://cloud.tencent.com/document/api/1273/128050#ReferResourceInfo)
* [ReportFileInfoReq](https://cloud.tencent.com/document/api/1273/128050#ReportFileInfoReq)
* [ReportFileInfoRsp](https://cloud.tencent.com/document/api/1273/128050#ReportFileInfoRsp)
* [StandardDrugCardInfo](https://cloud.tencent.com/document/api/1273/128050#StandardDrugCardInfo)
* [StandardDrugInstructionInfo](https://cloud.tencent.com/document/api/1273/128050#StandardDrugInstructionInfo)



## 实时互动-教育版(lcic) 版本：2022-08-17

### 第 85 次发布

发布时间：2026-05-29 02:19:37

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [EventDataInfo](https://cloud.tencent.com/document/api/1639/81423#EventDataInfo)

	* 新增成员：Role




## 云开发 CloudBase(tcb) 版本：2018-06-08

### 第 143 次发布

发布时间：2026-05-29 02:49:59

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribePGUserMigration](https://cloud.tencent.com/document/api/876/132262)
* [DescribeResourcePermission](https://cloud.tencent.com/document/api/876/132256)
* [ListPGUserMigrations](https://cloud.tencent.com/document/api/876/132261)
* [ModifyResourcePermission](https://cloud.tencent.com/document/api/876/132255)
* [PreviewPGUserMigrations](https://cloud.tencent.com/document/api/876/132260)
* [PushPGUserMigrations](https://cloud.tencent.com/document/api/876/132259)
* [RepairPGUserMigrationHistory](https://cloud.tencent.com/document/api/876/132258)
* [RollbackPGUserMigrations](https://cloud.tencent.com/document/api/876/132257)

新增数据结构：

* [DescribeResourcePermissionResult](https://cloud.tencent.com/document/api/876/34822#DescribeResourcePermissionResult)
* [MigrationConflict](https://cloud.tencent.com/document/api/876/34822#MigrationConflict)
* [MigrationInput](https://cloud.tencent.com/document/api/876/34822#MigrationInput)
* [MigrationPlanItem](https://cloud.tencent.com/document/api/876/34822#MigrationPlanItem)
* [MigrationSummary](https://cloud.tencent.com/document/api/876/34822#MigrationSummary)
* [ModifyResourcePermissionResult](https://cloud.tencent.com/document/api/876/34822#ModifyResourcePermissionResult)
* [ResourcePermission](https://cloud.tencent.com/document/api/876/34822#ResourcePermission)



## 高性能计算平台(thpc) 版本：2023-03-21

### 第 37 次发布

发布时间：2026-05-29 03:03:00

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [CosOption](https://cloud.tencent.com/document/api/1527/89579#CosOption)

修改数据结构：

* [GooseFSOption](https://cloud.tencent.com/document/api/1527/89579#GooseFSOption)

	* 新增成员：FileSystemId

	* <font color="#dd0000">**修改成员**：</font>Masters

* [StorageOption](https://cloud.tencent.com/document/api/1527/89579#StorageOption)

	* 新增成员：CosOptions




## 高性能计算平台(thpc) 版本：2022-04-01



## 高性能计算平台(thpc) 版本：2021-11-09



## 容器服务(tke) 版本：2022-05-01



## 容器服务(tke) 版本：2018-05-25

### 第 230 次发布

发布时间：2026-05-29 03:05:07

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateClusterVirtualNodePool](https://cloud.tencent.com/document/api/457/85354)

	* 新增入参：SubnetAllocationPolicy, AgentPlugin


新增数据结构：

* [AgentPluginConfig](https://cloud.tencent.com/document/api/457/31866#AgentPluginConfig)



## 云点播(vod) 版本：2024-07-18



## 云点播(vod) 版本：2018-07-17

### 第 259 次发布

发布时间：2026-05-29 03:16:59

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeAigcUsageData](https://cloud.tencent.com/document/api/266/126446)

	* 新增入参：APIKey




