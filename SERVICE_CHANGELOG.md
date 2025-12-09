# Release v1.3.9

## 腾讯云数据分析智能体(dataagent) 版本：2025-05-13

### 第 7 次发布

发布时间：2025-12-09 01:32:56

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [GetKnowledgeBaseFileList](https://cloud.tencent.com/document/api/1800/126199)

	* 新增出参：FileList, Total


新增数据结构：

* [ColumnInfo](https://cloud.tencent.com/document/api/1800/125016#ColumnInfo)
* [FileInfo](https://cloud.tencent.com/document/api/1800/125016#FileInfo)
* [KnowledgeTaskConfig](https://cloud.tencent.com/document/api/1800/125016#KnowledgeTaskConfig)



## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 266 次发布

发布时间：2025-12-09 01:42:04

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateBatchQuickSignUrl](https://cloud.tencent.com/document/api/1323/101058)

	* 新增入参：CanSkipReadFlow

* [CreateEmployeeQualificationSealQrCode](https://cloud.tencent.com/document/api/1323/108596)

	* 新增入参：UserData




## 腾讯电子签（基础版）(essbasic) 版本：2021-05-26

### 第 248 次发布

发布时间：2025-12-09 01:43:30

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateEmployeeQualificationSealQrCode](https://cloud.tencent.com/document/api/1420/108597)

	* 新增入参：UserData




## 腾讯电子签（基础版）(essbasic) 版本：2020-12-22



## 容器服务(tke) 版本：2022-05-01



## 容器服务(tke) 版本：2018-05-25

### 第 211 次发布

发布时间：2025-12-09 03:02:10

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateClusterVirtualNodePool](https://cloud.tencent.com/document/api/457/85354)

	* <font color="#dd0000">**修改入参**：</font>SecurityGroupIds




## 云点播(vod) 版本：2024-07-18



## 云点播(vod) 版本：2018-07-17

### 第 210 次发布

发布时间：2025-12-09 03:20:36

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateAigcImageTask](https://cloud.tencent.com/document/api/266/126240)
* [CreateAigcVideoTask](https://cloud.tencent.com/document/api/266/126239)

修改接口：

* [DescribeTaskDetail](https://cloud.tencent.com/document/api/266/33431)

	* 新增出参：AigcImageTask, AigcVideoTask


新增数据结构：

* [AigcImageOutputConfig](https://cloud.tencent.com/document/api/266/31773#AigcImageOutputConfig)
* [AigcImageTask](https://cloud.tencent.com/document/api/266/31773#AigcImageTask)
* [AigcImageTaskInput](https://cloud.tencent.com/document/api/266/31773#AigcImageTaskInput)
* [AigcImageTaskInputFileInfo](https://cloud.tencent.com/document/api/266/31773#AigcImageTaskInputFileInfo)
* [AigcImageTaskOutput](https://cloud.tencent.com/document/api/266/31773#AigcImageTaskOutput)
* [AigcImageTaskOutputFileInfo](https://cloud.tencent.com/document/api/266/31773#AigcImageTaskOutputFileInfo)
* [AigcVideoOutputConfig](https://cloud.tencent.com/document/api/266/31773#AigcVideoOutputConfig)
* [AigcVideoTask](https://cloud.tencent.com/document/api/266/31773#AigcVideoTask)
* [AigcVideoTaskInput](https://cloud.tencent.com/document/api/266/31773#AigcVideoTaskInput)
* [AigcVideoTaskInputFileInfo](https://cloud.tencent.com/document/api/266/31773#AigcVideoTaskInputFileInfo)
* [AigcVideoTaskOutput](https://cloud.tencent.com/document/api/266/31773#AigcVideoTaskOutput)
* [AigcVideoTaskOutputFileInfo](https://cloud.tencent.com/document/api/266/31773#AigcVideoTaskOutputFileInfo)

修改数据结构：

* [EventContent](https://cloud.tencent.com/document/api/266/31773#EventContent)

	* 新增成员：AigcImageCompleteEvent, AigcVideoCompleteEvent




## 数据开发治理平台 WeData(wedata) 版本：2025-08-06



## 数据开发治理平台 WeData(wedata) 版本：2021-08-20

### 第 171 次发布

发布时间：2025-12-09 11:17:55

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeOperateOpsTasks](https://cloud.tencent.com/document/api/1267/95254)

	* 新增入参：IncludeManualTask, CheckPrivilege

* [DescribeOpsWorkflows](https://cloud.tencent.com/document/api/1267/95200)

	* 新增入参：CheckPrivilege


新增数据结构：

* [AssetDim](https://cloud.tencent.com/document/api/1267/76336#AssetDim)
* [AssetDimTableColumn](https://cloud.tencent.com/document/api/1267/76336#AssetDimTableColumn)
* [DimTableLink](https://cloud.tencent.com/document/api/1267/76336#DimTableLink)
* [IndicatorBaseSimpleInfo](https://cloud.tencent.com/document/api/1267/76336#IndicatorBaseSimpleInfo)
* [PermissionStatus](https://cloud.tencent.com/document/api/1267/76336#PermissionStatus)

修改数据结构：

* [ColumnMeta](https://cloud.tencent.com/document/api/1267/76336#ColumnMeta)

	* 新增成员：CategoryName, OriginType, IndicatorBase, AssetDim

* [CommonIdOpsDto](https://cloud.tencent.com/document/api/1267/76336#CommonIdOpsDto)

	* 新增成员：MakeId

* [DataSourceEnvInfo](https://cloud.tencent.com/document/api/1267/76336#DataSourceEnvInfo)

	* 新增成员：TcCatalogOpen

* [DataSourceInfo](https://cloud.tencent.com/document/api/1267/76336#DataSourceInfo)

	* 新增成员：TcCatalogOpen

* [InstanceOpsDto](https://cloud.tencent.com/document/api/1267/76336#InstanceOpsDto)

	* 新增成员：Privileges

* [ProdSchedulerTask](https://cloud.tencent.com/document/api/1267/76336#ProdSchedulerTask)

	* 新增成员：InChargeIdList, InChargeNameList

* [SuccessorTaskInfo](https://cloud.tencent.com/document/api/1267/76336#SuccessorTaskInfo)

	* 新增成员：Privileges

* [TagVoteSum](https://cloud.tencent.com/document/api/1267/76336#TagVoteSum)

	* 新增成员：TagValueId, TagValue, TagIsDeleted, TagValueIsDeleted

	* <font color="#dd0000">**修改成员**：</font>VoteSum, Status, TagDesc

* [TaskOpsDto](https://cloud.tencent.com/document/api/1267/76336#TaskOpsDto)

	* 新增成员：Privileges, BundleId, BundleInfo

* [WorkflowExtOpsDto](https://cloud.tencent.com/document/api/1267/76336#WorkflowExtOpsDto)

	* 新增成员：BundleId, BundleInfo




