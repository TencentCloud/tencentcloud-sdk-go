# Release v1.0.1169

## 云硬盘(cbs) 版本：2017-03-12

### 第 69 次发布

发布时间：2025-05-20 01:09:48

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [UnbindAutoSnapshotPolicy](https://cloud.tencent.com/document/api/362/33554)

	* 新增入参：InstanceIds




## 云数据库 MySQL(cdb) 版本：2017-03-20

### 第 192 次发布

发布时间：2025-05-20 01:10:11

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeBackupConfig](https://cloud.tencent.com/document/api/236/15837)

* [IsolateDBInstance](https://cloud.tencent.com/document/api/236/15869)

* [ModifyDBInstanceVipVport](https://cloud.tencent.com/document/api/236/15867)

* [ModifyRoGroupInfo](https://cloud.tencent.com/document/api/236/40938)




## 腾讯云数据仓库 TCHouse-D(cdwdoris) 版本：2021-12-28

### 第 47 次发布

发布时间：2025-05-20 01:11:08

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeInstanceState](https://cloud.tencent.com/document/api/1387/102615)

	* 新增出参：ProcessId, JobName




## TDSQL-C MySQL 版(cynosdb) 版本：2019-01-07

### 第 133 次发布

发布时间：2025-05-20 01:14:53

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateAuditLogFile](https://cloud.tencent.com/document/api/1003/81440)

* [CreateClusters](https://cloud.tencent.com/document/api/1003/48087)

	* 新增入参：GdnId, ProxyConfig

* [DescribeAuditLogs](https://cloud.tencent.com/document/api/1003/81437)


新增数据结构：

* [ProxyConfig](https://cloud.tencent.com/document/api/1003/48097#ProxyConfig)

修改数据结构：

* [CynosdbCluster](https://cloud.tencent.com/document/api/1003/48097#CynosdbCluster)

	* 新增成员：GdnId

* [CynosdbClusterDetail](https://cloud.tencent.com/document/api/1003/48097#CynosdbClusterDetail)

	* 新增成员：GdnId, GdnRole

* [CynosdbInstance](https://cloud.tencent.com/document/api/1003/48097#CynosdbInstance)

	* 新增成员：GdnId




## DNSPod(dnspod) 版本：2021-03-23

### 第 42 次发布

发布时间：2025-05-20 01:16:29

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [BatchSearchRecordInfo](https://cloud.tencent.com/document/api/1427/56185#BatchSearchRecordInfo)

	* <font color="#dd0000">**修改成员**：</font>RecordId, Remark, Enabled, Weight, GroupId, MX




## 弹性 MapReduce(emr) 版本：2019-01-03

### 第 101 次发布

发布时间：2025-05-20 01:18:04

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeClusterFlowStatusDetail](https://cloud.tencent.com/document/api/589/107224)

	* 新增出参：FlowInfo




## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 221 次发布

发布时间：2025-05-20 01:18:34

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [CreateFlowOption](https://cloud.tencent.com/document/api/1323/70369#CreateFlowOption)

	* 新增成员：ForbidEditWatermark

* [EmbedUrlOption](https://cloud.tencent.com/document/api/1323/70369#EmbedUrlOption)

	* 新增成员：ForbidEditWatermark




## 腾讯电子签（基础版）(essbasic) 版本：2021-05-26

### 第 218 次发布

发布时间：2025-05-20 01:19:01

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [CreateFlowOption](https://cloud.tencent.com/document/api/1420/61525#CreateFlowOption)

	* 新增成员：ForbidEditWatermark

* [EmbedUrlOption](https://cloud.tencent.com/document/api/1420/61525#EmbedUrlOption)

	* 新增成员：ForbidEditWatermark




## 腾讯电子签（基础版）(essbasic) 版本：2020-12-22



## 云游戏(gs) 版本：2019-11-18

### 第 33 次发布

发布时间：2025-05-20 01:19:48

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [AndroidInstance](https://cloud.tencent.com/document/api/1162/40743#AndroidInstance)

	* 新增成员：CreateTime, HostServerSerialNumber




## 腾讯云智能体开发平台(lke) 版本：2023-11-30

### 第 42 次发布

发布时间：2025-05-20 01:23:15

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [AICallConfig](https://cloud.tencent.com/document/api/1759/105104#AICallConfig)
* [DigitalHumanConfig](https://cloud.tencent.com/document/api/1759/105104#DigitalHumanConfig)
* [VoiceConfig](https://cloud.tencent.com/document/api/1759/105104#VoiceConfig)
* [WorkflowInfo](https://cloud.tencent.com/document/api/1759/105104#WorkflowInfo)

修改数据结构：

* [AgentProcedure](https://cloud.tencent.com/document/api/1759/105104#AgentProcedure)

	* 新增成员：AgentIcon

* [AgentProcedureDebugging](https://cloud.tencent.com/document/api/1759/105104#AgentProcedureDebugging)

	* 新增成员：DisplayStatus, SandboxUrl, DisplayUrl

* [AgentThought](https://cloud.tencent.com/document/api/1759/105104#AgentThought)

	* 新增成员：Files

* [Credentials](https://cloud.tencent.com/document/api/1759/105104#Credentials)

	* 新增成员：AppId

* [FileInfo](https://cloud.tencent.com/document/api/1759/105104#FileInfo)

	* 新增成员：CreatedAt

* [KnowledgeQaConfig](https://cloud.tencent.com/document/api/1759/105104#KnowledgeQaConfig)

	* 新增成员：AiCall

* [KnowledgeQaSingleWorkflow](https://cloud.tencent.com/document/api/1759/105104#KnowledgeQaSingleWorkflow)

	* 新增成员：AsyncWorkflow

* [ModelInfo](https://cloud.tencent.com/document/api/1759/105104#ModelInfo)

	* 新增成员：SupportAiCallStatus

* [MsgRecord](https://cloud.tencent.com/document/api/1759/105104#MsgRecord)

	* 新增成员：WorkFlow

* [PluginToolReqParam](https://cloud.tencent.com/document/api/1759/105104#PluginToolReqParam)

	* 新增成员：OneOf, AnyOf

* [ProcedureDebugging](https://cloud.tencent.com/document/api/1759/105104#ProcedureDebugging)

	* 新增成员：CustomVariables

* [WorkFlowSummary](https://cloud.tencent.com/document/api/1759/105104#WorkFlowSummary)

	* 新增成员：OptionCards, Outputs, WorkflowReleaseTime




## 云数据库 PostgreSQL(postgres) 版本：2017-03-12

### 第 55 次发布

发布时间：2025-05-20 01:26:21

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除接口**：</font>

* CloseServerlessDBExtranetAccess
* DeleteServerlessDBInstance



## 云压测(pts) 版本：2021-07-28

### 第 21 次发布

发布时间：2025-05-20 01:26:43

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [CustomSampleMatrix](https://cloud.tencent.com/document/api/1484/78100#CustomSampleMatrix)

	* 新增成员：Step




## 容器服务(tke) 版本：2022-05-01



## 容器服务(tke) 版本：2018-05-25

### 第 196 次发布

发布时间：2025-05-20 01:31:45

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除接口**：</font>

* ForwardApplicationRequestV3



## 文本内容安全(tms) 版本：2020-12-29

### 第 14 次发布

发布时间：2025-05-20 01:32:22

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [TextModeration](https://cloud.tencent.com/document/api/1124/51860)

	* 新增入参：Type

	* 新增出参：HitType




## 文本内容安全(tms) 版本：2020-07-13



