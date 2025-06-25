# Release v1.0.1195

## 运维安全中心（堡垒机）(bh) 版本：2023-04-18

### 第 15 次发布

发布时间：2025-06-26 01:04:54

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CheckLDAPConnection](https://cloud.tencent.com/document/api/1025/120145)
* [DescribeLDAPUnitSet](https://cloud.tencent.com/document/api/1025/120144)
* [ModifyLDAPSetting](https://cloud.tencent.com/document/api/1025/120143)
* [ReplaySession](https://cloud.tencent.com/document/api/1025/120147)
* [SetLDAPSyncFlag](https://cloud.tencent.com/document/api/1025/120142)
* [UnlockUser](https://cloud.tencent.com/document/api/1025/120146)



## 主机安全(cwp) 版本：2018-02-28

### 第 148 次发布

发布时间：2025-06-26 01:13:01

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeJavaMemShellList](https://cloud.tencent.com/document/api/296/81790)

	* 新增入参：Order, By

* [DescribeLicenseGeneral](https://cloud.tencent.com/document/api/296/80398)

	* 新增出参：AutoBindRaspSwitch, AutoOpenRaspSwitch

* [DescribeVulInfoCvss](https://cloud.tencent.com/document/api/296/60898)

	* 新增出参：SupportDefence

* [ExportJavaMemShells](https://cloud.tencent.com/document/api/296/99657)

	* 新增入参：Order, By

* [ModifyAutoOpenProVersionConfig](https://cloud.tencent.com/document/api/296/19863)

	* 新增入参：AutoBindRaspSwitch, AutoOpenRaspSwitch


修改数据结构：

* [VulEmergentMsgInfo](https://cloud.tencent.com/document/api/296/19867#VulEmergentMsgInfo)

	* 新增成员：SupportFix, SupportDefense




## TDSQL-C MySQL 版(cynosdb) 版本：2019-01-07

### 第 137 次发布

发布时间：2025-06-26 01:15:00

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateClusters](https://cloud.tencent.com/document/api/1003/48087)

	* 新增入参：AutoArchive, AutoArchiveDelayHours

* [DescribeInstances](https://cloud.tencent.com/document/api/1003/48334)

	* 新增入参：ClusterType




## 数据安全治理中心(dsgc) 版本：2019-07-23

### 第 33 次发布

发布时间：2025-06-26 01:18:01

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateDSPADiscoveryTask](https://cloud.tencent.com/document/api/1087/97252)

	* 新增入参：ScanRange

* [DescribeDSPACOSDiscoveryTaskFiles](https://cloud.tencent.com/document/api/1087/97228)

	* 新增入参：ScanResultId

	* <font color="#dd0000">**修改入参**：</font>BucketResultId

* [DescribeDSPACOSDiscoveryTaskResult](https://cloud.tencent.com/document/api/1087/97227)

	* 新增入参：StartTime, EndTime, FetchHistory

	* 新增出参：MaxCount

* [DescribeDSPACOSTaskResultDetail](https://cloud.tencent.com/document/api/1087/97225)

	* 新增入参：ScanResultId

	* <font color="#dd0000">**修改入参**：</font>BucketResultId

* [DescribeDSPADiscoveryTaskResult](https://cloud.tencent.com/document/api/1087/97212)

	* 新增入参：StartTime, EndTime, FetchHistory

	* 新增出参：MaxCount

* [DescribeDSPADiscoveryTaskResultDetail](https://cloud.tencent.com/document/api/1087/97211)

	* 新增入参：ScanResultId

* [DescribeDSPADiscoveryTaskTables](https://cloud.tencent.com/document/api/1087/97210)

	* 新增入参：ScanResultId

	* <font color="#dd0000">**修改入参**：</font>DbResultId

* [ModifyDSPADiscoveryTask](https://cloud.tencent.com/document/api/1087/97192)

	* 新增入参：ScanRange


修改数据结构：

* [DspaCOSDiscoveryTaskResult](https://cloud.tencent.com/document/api/1087/96844#DspaCOSDiscoveryTaskResult)

	* 新增成员：TaskInstanceId, StartTime

* [DspaDiscoveryTask](https://cloud.tencent.com/document/api/1087/96844#DspaDiscoveryTask)

	* 新增成员：ScanRange

* [DspaDiscoveryTaskDbResult](https://cloud.tencent.com/document/api/1087/96844#DspaDiscoveryTaskDbResult)

	* 新增成员：TaskInstanceId, StartTime, ScanRange

* [DspaDiscoveryTaskDetail](https://cloud.tencent.com/document/api/1087/96844#DspaDiscoveryTaskDetail)

	* 新增成员：ScanRange




## 弹性 MapReduce(emr) 版本：2019-01-03

### 第 110 次发布

发布时间：2025-06-26 01:19:59

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateCluster](https://cloud.tencent.com/document/api/589/83953)

	* 新增入参：NodeMarks

* [CreateInstance](https://cloud.tencent.com/document/api/589/34261)

	* 新增入参：NodeMarks

* [ScaleOutCluster](https://cloud.tencent.com/document/api/589/83952)

	* 新增入参：NodeMarks

* [ScaleOutInstance](https://cloud.tencent.com/document/api/589/34264)

	* 新增入参：NodeMarks


新增数据结构：

* [NodeMark](https://cloud.tencent.com/document/api/589/33981#NodeMark)

修改数据结构：

* [NodeHardwareInfo](https://cloud.tencent.com/document/api/589/33981#NodeHardwareInfo)

	* 新增成员：NodeMark




## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 231 次发布

发布时间：2025-06-25 10:33:16

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [RegisterInfo](https://cloud.tencent.com/document/api/1323/70369#RegisterInfo)

	* 新增成员：AuthorizationType




## 密钥管理系统(kms) 版本：2019-01-18

### 第 24 次发布

发布时间：2025-06-26 01:26:54

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CancelDataKeyDeletion](https://cloud.tencent.com/document/api/573/120162)
* [DescribeDataKey](https://cloud.tencent.com/document/api/573/120161)
* [DescribeDataKeys](https://cloud.tencent.com/document/api/573/120160)
* [DisableDataKey](https://cloud.tencent.com/document/api/573/120159)
* [DisableDataKeys](https://cloud.tencent.com/document/api/573/120158)
* [EnableDataKey](https://cloud.tencent.com/document/api/573/120157)
* [EnableDataKeys](https://cloud.tencent.com/document/api/573/120156)
* [GetDataKeyCiphertextBlob](https://cloud.tencent.com/document/api/573/120155)
* [GetDataKeyPlaintext](https://cloud.tencent.com/document/api/573/120154)
* [ImportDataKey](https://cloud.tencent.com/document/api/573/120153)
* [ListDataKeyDetail](https://cloud.tencent.com/document/api/573/120152)
* [ListDataKeys](https://cloud.tencent.com/document/api/573/120151)
* [ScheduleDataKeyDeletion](https://cloud.tencent.com/document/api/573/120150)
* [UpdateDataKeyDescription](https://cloud.tencent.com/document/api/573/120149)
* [UpdateDataKeyName](https://cloud.tencent.com/document/api/573/120148)

修改接口：

* [GenerateDataKey](https://cloud.tencent.com/document/api/573/34419)

	* 新增入参：IsHostedByKms, DataKeyName, Description, HsmClusterId

	* 新增出参：DataKeyId

* [GetServiceStatus](https://cloud.tencent.com/document/api/573/34417)

	* 新增出参：IsAllowedDataKeyHosted, DataKeyLimit, FreeDataKeyLimit, DataKeyUsedCount


新增数据结构：

* [DataKey](https://cloud.tencent.com/document/api/573/34431#DataKey)
* [DataKeyMetadata](https://cloud.tencent.com/document/api/573/34431#DataKeyMetadata)



## 腾讯云智能体开发平台(lke) 版本：2023-11-30

### 第 51 次发布

发布时间：2025-06-26 01:28:43

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateAgent](https://cloud.tencent.com/document/api/1759/120167)
* [DeleteAgent](https://cloud.tencent.com/document/api/1759/120166)
* [DescribeAppAgentList](https://cloud.tencent.com/document/api/1759/120165)
* [ModifyAgent](https://cloud.tencent.com/document/api/1759/120164)

修改接口：

* [DescribeDoc](https://cloud.tencent.com/document/api/1759/105071)

	* 新增出参：IsDownload

* [GetDocPreview](https://cloud.tencent.com/document/api/1759/105067)

	* 新增出参：IsDownload

* [SaveDoc](https://cloud.tencent.com/document/api/1759/105054)

	* 新增入参：IsDownload


新增数据结构：

* [Agent](https://cloud.tencent.com/document/api/1759/105104#Agent)
* [AgentInput](https://cloud.tencent.com/document/api/1759/105104#AgentInput)
* [AgentInputUserInputValue](https://cloud.tencent.com/document/api/1759/105104#AgentInputUserInputValue)
* [AgentKnowledgeAttrLabel](https://cloud.tencent.com/document/api/1759/105104#AgentKnowledgeAttrLabel)
* [AgentKnowledgeFilter](https://cloud.tencent.com/document/api/1759/105104#AgentKnowledgeFilter)
* [AgentKnowledgeFilterDocAndAnswer](https://cloud.tencent.com/document/api/1759/105104#AgentKnowledgeFilterDocAndAnswer)
* [AgentKnowledgeFilterTag](https://cloud.tencent.com/document/api/1759/105104#AgentKnowledgeFilterTag)
* [AgentKnowledgeQAPlugin](https://cloud.tencent.com/document/api/1759/105104#AgentKnowledgeQAPlugin)
* [AgentMCPServerInfo](https://cloud.tencent.com/document/api/1759/105104#AgentMCPServerInfo)
* [AgentModelInfo](https://cloud.tencent.com/document/api/1759/105104#AgentModelInfo)
* [AgentPluginHeader](https://cloud.tencent.com/document/api/1759/105104#AgentPluginHeader)
* [AgentPluginInfo](https://cloud.tencent.com/document/api/1759/105104#AgentPluginInfo)
* [AgentToolInfo](https://cloud.tencent.com/document/api/1759/105104#AgentToolInfo)
* [AgentToolReqParam](https://cloud.tencent.com/document/api/1759/105104#AgentToolReqParam)
* [AgentToolRspParam](https://cloud.tencent.com/document/api/1759/105104#AgentToolRspParam)

修改数据结构：

* [DocSegment](https://cloud.tencent.com/document/api/1759/105104#DocSegment)

	* 新增成员：WebUrl

* [MsgRecordReference](https://cloud.tencent.com/document/api/1759/105104#MsgRecordReference)

	* 新增成员：KnowledgeName, KnowledgeBizId, DocBizId, QaBizId




## 容器安全服务(tcss) 版本：2020-11-01

### 第 78 次发布

发布时间：2025-06-26 01:38:43

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeNewestVul](https://cloud.tencent.com/document/api/1285/81625)

	* 新增出参：SupportDefense




## 高性能计算平台(thpc) 版本：2023-03-21

### 第 25 次发布

发布时间：2025-06-26 01:41:55

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeNodes](https://cloud.tencent.com/document/api/1527/89569)

	* <font color="#dd0000">**修改入参**：</font>ClusterId


修改数据结构：

* [NodeOverview](https://cloud.tencent.com/document/api/1527/89579#NodeOverview)

	* 新增成员：ClusterId, NodeName, CreateTime




## 高性能计算平台(thpc) 版本：2022-04-01



## 高性能计算平台(thpc) 版本：2021-11-09



## TI-ONE 训练平台(tione) 版本：2021-11-11

### 第 84 次发布

发布时间：2025-06-26 01:42:11

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [Pod](https://cloud.tencent.com/document/api/851/75051#Pod)

	* 新增成员：NodeIP




## TI-ONE 训练平台(tione) 版本：2019-10-22



## 消息队列 RocketMQ 版(trocket) 版本：2023-03-08

### 第 41 次发布

发布时间：2025-06-25 19:47:54

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeConsumerClientList](https://cloud.tencent.com/document/api/1493/120140)



## 私有网络(vpc) 版本：2017-03-12

### 第 262 次发布

发布时间：2025-06-26 01:47:30

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [AssociateIPv6Address](https://cloud.tencent.com/document/api/215/113678)

	* 新增入参：InstanceId




## Web 应用防火墙(waf) 版本：2018-01-25

### 第 121 次发布

发布时间：2025-06-26 01:49:04

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [AddSpartaProtection](https://cloud.tencent.com/document/api/627/72689)

	* 新增入参：UseCase

* [ModifySpartaProtection](https://cloud.tencent.com/document/api/627/94309)

	* 新增入参：UseCase


修改数据结构：

* [ClbDomainsInfo](https://cloud.tencent.com/document/api/627/53609#ClbDomainsInfo)

	* 新增成员：AccessStatus

* [DomainsPartInfo](https://cloud.tencent.com/document/api/627/53609#DomainsPartInfo)

	* 新增成员：UseCase




## 数据开发治理平台 WeData(wedata) 版本：2021-08-20

### 第 157 次发布

发布时间：2025-06-26 01:49:56

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeDsTaskVersionInfo](https://cloud.tencent.com/document/api/1267/120169)
* [DescribeDsTaskVersionList](https://cloud.tencent.com/document/api/1267/120168)

新增数据结构：

* [EventPublisherDTO](https://cloud.tencent.com/document/api/1267/76336#EventPublisherDTO)
* [TaskVersionDsDTO](https://cloud.tencent.com/document/api/1267/76336#TaskVersionDsDTO)



