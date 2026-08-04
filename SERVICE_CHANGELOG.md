# Release v1.3.152

## 腾讯混元生3D(ai3d) 版本：2025-05-13

### 第 18 次发布

发布时间：2026-08-05 01:09:18

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [QueryHunyuan3DPartJob](https://cloud.tencent.com/document/api/1804/126296)

	* 新增出参：PartSegmentationInfoUrl

* [SubmitHunyuan3DPartJob](https://cloud.tencent.com/document/api/1804/126295)

	* 新增入参：EnablePostProcess




## 云托付物理服务器(chc) 版本：2023-04-18

### 第 11 次发布

发布时间：2026-08-05 01:30:28

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeDeviceWorkOrderDetail](https://cloud.tencent.com/document/api/1448/117177)

	* 新增出参：SLAInfo, PreOrderSet


新增数据结构：

* [SLAInfo](https://cloud.tencent.com/document/api/1448/117193#SLAInfo)

修改数据结构：

* [WorkOrderData](https://cloud.tencent.com/document/api/1448/117193#WorkOrderData)

	* 新增成员：SLAInfo




## 云安全一体化平台(csip) 版本：2022-11-21

### 第 93 次发布

发布时间：2026-08-05 01:40:36

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateDspmWhitelistStrategy](https://cloud.tencent.com/document/api/664/131523)

	* 新增入参：WhitelistType

	* <font color="#dd0000">**修改入参**：</font>StrategyType

* [ModifyDspmRiskStrategy](https://cloud.tencent.com/document/api/664/131468)

	* 新增入参：Name, Description, DbTypes


修改数据结构：

* [DspmRiskStrategy](https://cloud.tencent.com/document/api/664/90825#DspmRiskStrategy)

	* 新增成员：RuleSource, AssetTypes, RiskDescription

* [DspmRiskStrategyGroup](https://cloud.tencent.com/document/api/664/90825#DspmRiskStrategyGroup)

	* 新增成员：RuleSource

* [DspmWhitelistStrategy](https://cloud.tencent.com/document/api/664/90825#DspmWhitelistStrategy)

	* 新增成员：WhitelistType




## 云数据库独享集群(dbdc) 版本：2020-10-29

### 第 11 次发布

发布时间：2026-08-05 01:55:50

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateDBCustomCluster](https://cloud.tencent.com/document/api/1322/132930)

	* 新增入参：DeletionProtection

* [DescribeDBCustomClusterDetail](https://cloud.tencent.com/document/api/1322/132928)

	* 新增出参：DeletionProtection


修改数据结构：

* [DBCustomCluster](https://cloud.tencent.com/document/api/1322/74754#DBCustomCluster)

	* 新增成员：DeletionProtection




## 数据湖计算 DLC(dlc) 版本：2021-01-25

### 第 170 次发布

发布时间：2026-08-05 01:57:43

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CheckModifyPartition](https://cloud.tencent.com/document/api/1342/135521)
* [CheckQueueName](https://cloud.tencent.com/document/api/1342/135520)
* [CheckResourceName](https://cloud.tencent.com/document/api/1342/135519)
* [CreateInferenceModel](https://cloud.tencent.com/document/api/1342/135526)
* [CreatePartition](https://cloud.tencent.com/document/api/1342/135518)
* [CreatePartitionQueue](https://cloud.tencent.com/document/api/1342/135517)
* [DeletePartitionQueue](https://cloud.tencent.com/document/api/1342/135516)
* [DescribeFlowDetailList](https://cloud.tencent.com/document/api/1342/135515)
* [DescribeFlowList](https://cloud.tencent.com/document/api/1342/135514)
* [DescribePartitionDetail](https://cloud.tencent.com/document/api/1342/135513)
* [DescribePartitionQueues](https://cloud.tencent.com/document/api/1342/135512)
* [DescribePartitions](https://cloud.tencent.com/document/api/1342/135511)
* [DescribeSaleRegions](https://cloud.tencent.com/document/api/1342/135510)
* [DescribeSaleResourceInfo](https://cloud.tencent.com/document/api/1342/135509)
* [GetInferenceModel](https://cloud.tencent.com/document/api/1342/135525)
* [ListInferenceModels](https://cloud.tencent.com/document/api/1342/135524)
* [ModifyPartitionDescription](https://cloud.tencent.com/document/api/1342/135508)
* [ModifyPartitionQueue](https://cloud.tencent.com/document/api/1342/135507)
* [UpdateInferenceModel](https://cloud.tencent.com/document/api/1342/135523)

新增数据结构：

* [FlowActivityDetail](https://cloud.tencent.com/document/api/1342/53778#FlowActivityDetail)
* [FlowDetail](https://cloud.tencent.com/document/api/1342/53778#FlowDetail)
* [FlowInfo](https://cloud.tencent.com/document/api/1342/53778#FlowInfo)
* [InferenceModelInfo](https://cloud.tencent.com/document/api/1342/53778#InferenceModelInfo)
* [MessageItem](https://cloud.tencent.com/document/api/1342/53778#MessageItem)
* [PartitionDetail](https://cloud.tencent.com/document/api/1342/53778#PartitionDetail)
* [PartitionInfo](https://cloud.tencent.com/document/api/1342/53778#PartitionInfo)
* [QueueInfo](https://cloud.tencent.com/document/api/1342/53778#QueueInfo)
* [RegionInfo](https://cloud.tencent.com/document/api/1342/53778#RegionInfo)
* [ResourceQuota](https://cloud.tencent.com/document/api/1342/53778#ResourceQuota)
* [ResourceSaleInfo](https://cloud.tencent.com/document/api/1342/53778#ResourceSaleInfo)
* [ResourceSpec](https://cloud.tencent.com/document/api/1342/53778#ResourceSpec)
* [ResourceUsage](https://cloud.tencent.com/document/api/1342/53778#ResourceUsage)



## 数据传输服务(dts) 版本：2021-12-06

### 第 59 次发布

发布时间：2026-08-05 02:05:00

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [OnlineDDL](https://cloud.tencent.com/document/api/571/82108#OnlineDDL)

修改数据结构：

* [Objects](https://cloud.tencent.com/document/api/571/82108#Objects)

	* 新增成员：OnlineDDL




## 数据传输服务(dts) 版本：2018-03-30



## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 314 次发布

发布时间：2026-08-05 02:11:21

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreatePrepareFlowGroup](https://cloud.tencent.com/document/api/1323/118401)

	* 新增入参：FlowGroupType, FlowGroupDeadline


修改数据结构：

* [FlowGroupOptions](https://cloud.tencent.com/document/api/1323/70369#FlowGroupOptions)

	* 新增成员：NoEditFlowName, NoEditFlowType, NoEditDeadline, SignComponentConfig, ForbidEditWatermark, HideSignCodeAfterStart, SignAfterStart, PreviewAfterStart




## 腾讯电子签（基础版）(essbasic) 版本：2021-05-26

### 第 269 次发布

发布时间：2026-08-05 02:13:23

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ChannelCreatePrepareFlowGroup](https://cloud.tencent.com/document/api/1420/118402)

	* 新增入参：FlowGroupOptions, FlowGroupType, FlowGroupDeadline


修改数据结构：

* [FlowGroupOptions](https://cloud.tencent.com/document/api/1420/61525#FlowGroupOptions)

	* 新增成员：NoEditFlowName, NoEditFlowType, NoEditDeadline, SignComponentConfig, ForbidEditWatermark, HideSignCodeAfterStart, SignAfterStart, PreviewAfterStart




## 腾讯电子签（基础版）(essbasic) 版本：2020-12-22



## 密钥管理系统(kms) 版本：2019-01-18

### 第 30 次发布

发布时间：2026-08-05 02:31:03

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [RotateKey](https://cloud.tencent.com/document/api/573/135527)



## 消息队列 MQTT 版(mqtt) 版本：2024-05-16

### 第 33 次发布

发布时间：2026-08-05 02:44:54

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateBlockRule](https://cloud.tencent.com/document/api/1778/135532)
* [DeleteBlockRule](https://cloud.tencent.com/document/api/1778/135531)
* [DescribeBlockRuleList](https://cloud.tencent.com/document/api/1778/135530)
* [ModifyBlockRule](https://cloud.tencent.com/document/api/1778/135529)

新增数据结构：

* [BlockRuleItem](https://cloud.tencent.com/document/api/1778/111031#BlockRuleItem)



## 腾讯健康组学平台(omics) 版本：2022-11-28

### 第 31 次发布

发布时间：2026-08-05 02:47:16

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DeleteEnvironmentCache](https://cloud.tencent.com/document/api/1643/135533)



## 私有网络(vpc) 版本：2017-03-12

### 第 308 次发布

发布时间：2026-08-05 03:34:41

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateVpnConnection](https://cloud.tencent.com/document/api/215/17522)




