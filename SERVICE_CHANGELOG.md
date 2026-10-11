# Release v1.3.193

## Agent 沙箱服务(ags) 版本：2025-09-20

### 第 27 次发布

发布时间：2026-10-12 01:08:37

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateSandboxSnapshot](https://cloud.tencent.com/document/api/1814/139277)
* [DeleteSandboxSnapshot](https://cloud.tencent.com/document/api/1814/139276)
* [DescribeSandboxSnapshotList](https://cloud.tencent.com/document/api/1814/139275)

修改接口：

* [StartSandboxInstance](https://cloud.tencent.com/document/api/1814/124816)

	* 新增入参：SourceInstanceId


新增数据结构：

* [SandboxSnapshot](https://cloud.tencent.com/document/api/1814/124823#SandboxSnapshot)



## 应用性能监控(apm) 版本：2021-06-22

### 第 71 次发布

发布时间：2026-10-12 01:12:48

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeSpanTreeByID](https://cloud.tencent.com/document/api/1463/139278)

新增数据结构：

* [ChartData](https://cloud.tencent.com/document/api/1463/64927#ChartData)
* [ErrorSpan](https://cloud.tencent.com/document/api/1463/64927#ErrorSpan)
* [Series](https://cloud.tencent.com/document/api/1463/64927#Series)
* [SpanTreeView](https://cloud.tencent.com/document/api/1463/64927#SpanTreeView)
* [TagValues](https://cloud.tencent.com/document/api/1463/64927#TagValues)
* [TraceOverview](https://cloud.tencent.com/document/api/1463/64927#TraceOverview)
* [YAxis](https://cloud.tencent.com/document/api/1463/64927#YAxis)



## 暴露面管理服务(ctem) 版本：2023-11-28

### 第 21 次发布

发布时间：2026-10-12 01:44:28

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [DisplayConfig](https://cloud.tencent.com/document/api/1755/120320#DisplayConfig)

	* 新增成员：Status, LastCheckTime

* [DisplaySensitiveInfo](https://cloud.tencent.com/document/api/1755/120320#DisplaySensitiveInfo)

	* 新增成员：Status, LastCheckTime




## 大数据智能体工作台DataBuddy(databuddy) 版本：2026-07-15

### 第 7 次发布

发布时间：2026-10-12 01:53:43

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [AddWorkspaceGroups](https://cloud.tencent.com/document/api/1835/139286)
* [AddWorkspaceUsers](https://cloud.tencent.com/document/api/1835/139285)
* [DeleteTable](https://cloud.tencent.com/document/api/1835/139293)
* [GetCatalog](https://cloud.tencent.com/document/api/1835/139292)
* [GetSchema](https://cloud.tencent.com/document/api/1835/139291)
* [GetTable](https://cloud.tencent.com/document/api/1835/139290)
* [GetTables](https://cloud.tencent.com/document/api/1835/139289)
* [ListCatalogs](https://cloud.tencent.com/document/api/1835/139288)
* [ListTables](https://cloud.tencent.com/document/api/1835/139287)
* [ListWorkspaceGroupUsers](https://cloud.tencent.com/document/api/1835/139284)
* [ListWorkspaceGroups](https://cloud.tencent.com/document/api/1835/139283)
* [ListWorkspaceUsers](https://cloud.tencent.com/document/api/1835/139282)
* [RemoveWorkspaceGroups](https://cloud.tencent.com/document/api/1835/139281)
* [UpdateWorkspaceGroupsRoles](https://cloud.tencent.com/document/api/1835/139280)
* [UpdateWorkspaceUsersRoles](https://cloud.tencent.com/document/api/1835/139279)

修改接口：

* [CreateCatalog](https://cloud.tencent.com/document/api/1835/139028)

	* 新增入参：ConnectionCatalogName

* [CreateSchema](https://cloud.tencent.com/document/api/1835/139027)

	* 新增入参：CatalogSource

	* <font color="#dd0000">**修改入参**：</font>WorkspaceId

* [DeleteCatalog](https://cloud.tencent.com/document/api/1835/139026)

	* 新增入参：CatalogSource

* [DeleteSchema](https://cloud.tencent.com/document/api/1835/139025)

	* 新增入参：CatalogSource

* [ListSchemas](https://cloud.tencent.com/document/api/1835/139024)

	* 新增入参：CatalogSource


新增数据结构：

* [BucketPartitioning](https://cloud.tencent.com/document/api/1835/138006#BucketPartitioning)
* [Catalog](https://cloud.tencent.com/document/api/1835/138006#Catalog)
* [ColumnInfo](https://cloud.tencent.com/document/api/1835/138006#ColumnInfo)
* [DataMaskGroup](https://cloud.tencent.com/document/api/1835/138006#DataMaskGroup)
* [DataMaskLabelInfo](https://cloud.tencent.com/document/api/1835/138006#DataMaskLabelInfo)
* [DataMaskStrategyInfo](https://cloud.tencent.com/document/api/1835/138006#DataMaskStrategyInfo)
* [DeleteTableRsp](https://cloud.tencent.com/document/api/1835/138006#DeleteTableRsp)
* [DictDimensionInfo](https://cloud.tencent.com/document/api/1835/138006#DictDimensionInfo)
* [DictItem](https://cloud.tencent.com/document/api/1835/138006#DictItem)
* [FeatureTableDetail](https://cloud.tencent.com/document/api/1835/138006#FeatureTableDetail)
* [GetCatalogRsp](https://cloud.tencent.com/document/api/1835/138006#GetCatalogRsp)
* [GetSchemaRsp](https://cloud.tencent.com/document/api/1835/138006#GetSchemaRsp)
* [GetTableRsp](https://cloud.tencent.com/document/api/1835/138006#GetTableRsp)
* [GetTablesRsp](https://cloud.tencent.com/document/api/1835/138006#GetTablesRsp)
* [GroupUserInfo](https://cloud.tencent.com/document/api/1835/138006#GroupUserInfo)
* [Index](https://cloud.tencent.com/document/api/1835/138006#Index)
* [ListCatalogsRsp](https://cloud.tencent.com/document/api/1835/138006#ListCatalogsRsp)
* [ListPartition](https://cloud.tencent.com/document/api/1835/138006#ListPartition)
* [ListPartitioning](https://cloud.tencent.com/document/api/1835/138006#ListPartitioning)
* [ListTablesRsp](https://cloud.tencent.com/document/api/1835/138006#ListTablesRsp)
* [ListWorkspaceGroupUsersRsp](https://cloud.tencent.com/document/api/1835/138006#ListWorkspaceGroupUsersRsp)
* [ListWorkspaceGroupsRsp](https://cloud.tencent.com/document/api/1835/138006#ListWorkspaceGroupsRsp)
* [ListWorkspaceUsersRsp](https://cloud.tencent.com/document/api/1835/138006#ListWorkspaceUsersRsp)
* [Literal](https://cloud.tencent.com/document/api/1835/138006#Literal)
* [Partitioning](https://cloud.tencent.com/document/api/1835/138006#Partitioning)
* [PlatformStatusRsp](https://cloud.tencent.com/document/api/1835/138006#PlatformStatusRsp)
* [RangePartition](https://cloud.tencent.com/document/api/1835/138006#RangePartition)
* [RangePartitioning](https://cloud.tencent.com/document/api/1835/138006#RangePartitioning)
* [SingleFieldPartitioning](https://cloud.tencent.com/document/api/1835/138006#SingleFieldPartitioning)
* [TableInfo](https://cloud.tencent.com/document/api/1835/138006#TableInfo)
* [TagValueInfo](https://cloud.tencent.com/document/api/1835/138006#TagValueInfo)
* [TruncatePartitioning](https://cloud.tencent.com/document/api/1835/138006#TruncatePartitioning)
* [UserTagInfo](https://cloud.tencent.com/document/api/1835/138006#UserTagInfo)
* [WorkspaceGroupInfo](https://cloud.tencent.com/document/api/1835/138006#WorkspaceGroupInfo)
* [WorkspaceUserInfo](https://cloud.tencent.com/document/api/1835/138006#WorkspaceUserInfo)

修改数据结构：

* [FileConfig](https://cloud.tencent.com/document/api/1835/138006#FileConfig)

	* 新增成员：DefaultCatalogSource

* [FileInfo](https://cloud.tencent.com/document/api/1835/138006#FileInfo)

	* 新增成员：VersionId

* [TaskType](https://cloud.tencent.com/document/api/1835/138006#TaskType)

* [TaskTypeNotebookExt](https://cloud.tencent.com/document/api/1835/138006#TaskTypeNotebookExt)

* [WorkflowBaseInfo](https://cloud.tencent.com/document/api/1835/138006#WorkflowBaseInfo)




## 数据湖计算 DLC(dlc) 版本：2021-01-25

### 第 184 次发布

发布时间：2026-10-12 01:57:46

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateJob](https://cloud.tencent.com/document/api/1342/138819)

	* 新增入参：Spec, Dependencies, Entrypoint

* [DescribeJobDetail](https://cloud.tencent.com/document/api/1342/138814)

	* 新增出参：JobDetail

* [DescribeJobResult](https://cloud.tencent.com/document/api/1342/138811)

	* 新增出参：Download


新增数据结构：

* [DownloadInfo](https://cloud.tencent.com/document/api/1342/53778#DownloadInfo)
* [EngineConfigDependencies](https://cloud.tencent.com/document/api/1342/53778#EngineConfigDependencies)
* [EntrypointInfo](https://cloud.tencent.com/document/api/1342/53778#EntrypointInfo)
* [JobDetailInfo](https://cloud.tencent.com/document/api/1342/53778#JobDetailInfo)
* [SparkJobSpec](https://cloud.tencent.com/document/api/1342/53778#SparkJobSpec)
* [Spec](https://cloud.tencent.com/document/api/1342/53778#Spec)

修改数据结构：

* [TaskFullRespInfo](https://cloud.tencent.com/document/api/1342/53778#TaskFullRespInfo)

	* 新增成员：Tag




## 数据加速器 GooseFS(goosefs) 版本：2022-05-19

### 第 24 次发布

发布时间：2026-10-12 02:16:38

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CancelPolicyRuleExecuteRecord](https://cloud.tencent.com/document/api/1424/139301)
* [CreatePolicyRule](https://cloud.tencent.com/document/api/1424/139300)
* [DeletePolicyRule](https://cloud.tencent.com/document/api/1424/139299)
* [DescribePolicyRule](https://cloud.tencent.com/document/api/1424/139298)
* [DescribePolicyRuleExecuteRecord](https://cloud.tencent.com/document/api/1424/139297)
* [DisablePolicyRule](https://cloud.tencent.com/document/api/1424/139296)
* [EnablePolicyRule](https://cloud.tencent.com/document/api/1424/139295)

修改接口：

* [ExpandCapacity](https://cloud.tencent.com/document/api/1424/95787)

	* 新增入参：RapidExpansion


新增数据结构：

* [ExecutePeriodAttribute](https://cloud.tencent.com/document/api/1424/95076#ExecutePeriodAttribute)
* [ExecutionPlan](https://cloud.tencent.com/document/api/1424/95076#ExecutionPlan)
* [ManifestAttr](https://cloud.tencent.com/document/api/1424/95076#ManifestAttr)
* [PolicyRuleOperator](https://cloud.tencent.com/document/api/1424/95076#PolicyRuleOperator)
* [PolicyRuleTaskAttribute](https://cloud.tencent.com/document/api/1424/95076#PolicyRuleTaskAttribute)
* [RuleExecuteAttribute](https://cloud.tencent.com/document/api/1424/95076#RuleExecuteAttribute)



## iOA 零信任安全管理系统(ioa) 版本：2022-06-01

### 第 47 次发布

发布时间：2026-10-12 02:19:55

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeAccountVirtualGroups](https://cloud.tencent.com/document/api/1092/139302)

新增数据结构：

* [DescribeAccountVirtualGroupsData](https://cloud.tencent.com/document/api/1092/102488#DescribeAccountVirtualGroupsData)
* [DescribeAccountVirtualGroupsPageData](https://cloud.tencent.com/document/api/1092/102488#DescribeAccountVirtualGroupsPageData)



## 云数据库 MongoDB(mongodb) 版本：2019-07-25

### 第 79 次发布

发布时间：2026-10-12 02:36:55

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateInstanceRole](https://cloud.tencent.com/document/api/240/139310)
* [CreateInstanceUser](https://cloud.tencent.com/document/api/240/139309)
* [DeleteInstanceRole](https://cloud.tencent.com/document/api/240/139308)
* [DeleteInstanceUser](https://cloud.tencent.com/document/api/240/139307)
* [DescribeInstanceRoles](https://cloud.tencent.com/document/api/240/139306)
* [DescribeInstanceUsers](https://cloud.tencent.com/document/api/240/139305)
* [UpdateInstanceRole](https://cloud.tencent.com/document/api/240/139304)
* [UpdateInstanceUser](https://cloud.tencent.com/document/api/240/139303)

新增数据结构：

* [InstanceUserInfo](https://cloud.tencent.com/document/api/240/38576#InstanceUserInfo)
* [Role](https://cloud.tencent.com/document/api/240/38576#Role)
* [RoleInfo](https://cloud.tencent.com/document/api/240/38576#RoleInfo)
* [RolePrivilege](https://cloud.tencent.com/document/api/240/38576#RolePrivilege)
* [RoleResource](https://cloud.tencent.com/document/api/240/38576#RoleResource)



## 云数据库 MongoDB(mongodb) 版本：2018-04-08



## 腾讯云可观测平台(monitor) 版本：2023-06-16

### 第 13 次发布

发布时间：2026-10-12 02:39:42

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeAlarmNotifyHistories](https://cloud.tencent.com/document/api/248/118735)

	* 新增入参：NotifyId


新增数据结构：

* [ChannelContent](https://cloud.tencent.com/document/api/248/115881#ChannelContent)
* [ReceiverGroup](https://cloud.tencent.com/document/api/248/115881#ReceiverGroup)

修改数据结构：

* [AlarmNotifyHistory](https://cloud.tencent.com/document/api/248/115881#AlarmNotifyHistory)

	* 新增成员：ReceiverGroups

* [ChannelsReceivers](https://cloud.tencent.com/document/api/248/115881#ChannelsReceivers)

	* 新增成员：SendStatusCode, SendTime, SubChannels, ExecDetail, RealDeliverStatus, RealDeliverTime, RealDeliverErrMsg, Contents




## 腾讯云可观测平台(monitor) 版本：2018-07-24

### 第 170 次发布

发布时间：2026-10-12 02:37:56

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyPrometheusConfig](https://cloud.tencent.com/document/api/248/86031)




## 媒体处理(mps) 版本：2019-06-12

### 第 254 次发布

发布时间：2026-10-12 02:40:11

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [AiDecomposeConfig](https://cloud.tencent.com/document/api/862/37615#AiDecomposeConfig)
* [AudioEnhanceEditConfig](https://cloud.tencent.com/document/api/862/37615#AudioEnhanceEditConfig)
* [DiffusionEnhanceV2Config](https://cloud.tencent.com/document/api/862/37615#DiffusionEnhanceV2Config)

修改数据结构：

* [AudioEnhanceConfig](https://cloud.tencent.com/document/api/862/37615#AudioEnhanceConfig)

	* 新增成员：AudioEnhanceEdit

* [ImageTaskInput](https://cloud.tencent.com/document/api/862/37615#ImageTaskInput)

	* 新增成员：AiDecomposeConfig

* [VideoEnhanceConfig](https://cloud.tencent.com/document/api/862/37615#VideoEnhanceConfig)

	* 新增成员：DiffusionEnhanceV2




## 消息队列 MQTT 版(mqtt) 版本：2024-05-16

### 第 35 次发布

发布时间：2026-10-09 18:04:13

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeClientEvents](https://cloud.tencent.com/document/api/1778/139235)

新增数据结构：

* [EventInfo](https://cloud.tencent.com/document/api/1778/111031#EventInfo)
* [EventItem](https://cloud.tencent.com/document/api/1778/111031#EventItem)



## 流计算 Oceanus(oceanus) 版本：2019-04-22

### 第 95 次发布

发布时间：2026-10-12 02:44:19

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [CCN](https://cloud.tencent.com/document/api/849/52010#CCN)

	* <font color="#dd0000">**修改成员**：</font>VpcId, SubnetId, CcnId

* [ClusterSession](https://cloud.tencent.com/document/api/849/52010#ClusterSession)

	* 新增成员：Description, Name, SerialId, ClusterGroupName, TotalCpu, TotalMem, WorkSpaceId, WorkSpaceName, Tags, LogConfig, UpdateUin

* [Setats](https://cloud.tencent.com/document/api/849/52010#Setats)

	* 新增成员：SecurityGroupIds, CCNs

* [Warehouse](https://cloud.tencent.com/document/api/849/52010#Warehouse)

	* 新增成员：StorageSerialId




## 文字识别(ocr) 版本：2018-11-19

### 第 271 次发布

发布时间：2026-10-12 02:44:58

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeMultimodalDocParseJob](https://cloud.tencent.com/document/api/866/139313)
* [RunRerank](https://cloud.tencent.com/document/api/866/139311)
* [SubmitMultimodalDocParseJob](https://cloud.tencent.com/document/api/866/139312)

新增数据结构：

* [ReconstructDocumentFailedPage](https://cloud.tencent.com/document/api/866/33527#ReconstructDocumentFailedPage)
* [RerankResult](https://cloud.tencent.com/document/api/866/33527#RerankResult)



## 云数据库Redis(redis) 版本：2018-04-12

### 第 110 次发布

发布时间：2026-10-12 02:50:34

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [UpgradeInstanceProductVersion](https://cloud.tencent.com/document/api/239/139314)



## TDSQL(tdmysql) 版本：2021-11-22

### 第 17 次发布

发布时间：2026-10-12 03:06:34

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeSpecs](https://cloud.tencent.com/document/api/1376/132042)

	* 新增入参：Zones

* [ModifyInstanceSSLStatus](https://cloud.tencent.com/document/api/1376/133010)

	* 新增入参：CertificateType, CertificateId


修改数据结构：

* [AutoScalingConfig](https://cloud.tencent.com/document/api/1376/128305#AutoScalingConfig)

	* 新增成员：ScaleUpEnabled, ScaleDownEnabled, ThresholdPercent, Step




## TI-ONE 训练平台(tione) 版本：2021-11-11

### 第 138 次发布

发布时间：2026-10-12 03:11:36

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [EnableTrainingTaskClsConfig](https://cloud.tencent.com/document/api/851/139315)



## TI-ONE 训练平台(tione) 版本：2019-10-22



## 容器服务(tke) 版本：2022-05-01



## 容器服务(tke) 版本：2018-05-25

### 第 238 次发布

发布时间：2026-10-12 03:13:00

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**预下线接口**：</font>

* DescribeClusterSecurity



