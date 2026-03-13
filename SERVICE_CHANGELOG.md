# Release v1.3.56

## 数据开发治理平台 WeData(wedata) 版本：2025-08-06

### 第 15 次发布

发布时间：2026-03-13 12:01:25

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateTriggerWorkflowRun](https://cloud.tencent.com/document/api/1267/129123)

新增数据结构：

* [CreateTriggerWorkflowRunResult](https://cloud.tencent.com/document/api/1267/123643#CreateTriggerWorkflowRunResult)



## 数据开发治理平台 WeData(wedata) 版本：2021-08-20

### 第 179 次发布

发布时间：2026-03-13 12:02:16

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeDataAssets](https://cloud.tencent.com/document/api/1267/129140)

修改接口：

* [DescribeTableMeta](https://cloud.tencent.com/document/api/1267/102541)

	* <font color="#dd0000">**修改出参**：</font>TagVoteSumList


新增数据结构：

* [AssetDimSimpleVO](https://cloud.tencent.com/document/api/1267/76336#AssetDimSimpleVO)
* [ChangeLog](https://cloud.tencent.com/document/api/1267/76336#ChangeLog)
* [DataAssetOption](https://cloud.tencent.com/document/api/1267/76336#DataAssetOption)
* [DataSetRecord](https://cloud.tencent.com/document/api/1267/76336#DataSetRecord)
* [IndicatorBaseInfo](https://cloud.tencent.com/document/api/1267/76336#IndicatorBaseInfo)
* [RegisteredModelAlias](https://cloud.tencent.com/document/api/1267/76336#RegisteredModelAlias)
* [RegisteredModelTag](https://cloud.tencent.com/document/api/1267/76336#RegisteredModelTag)

修改数据结构：

* [DataServiceRequestParam](https://cloud.tencent.com/document/api/1267/76336#DataServiceRequestParam)

	* 新增成员：StdCode

* [IntegrationTaskInfo](https://cloud.tencent.com/document/api/1267/76336#IntegrationTaskInfo)

	* 新增成员：TaskSubType, NotExistsCheckPoint, SavePointId, SavePointPath

* [TaskDataRegistryDTO](https://cloud.tencent.com/document/api/1267/76336#TaskDataRegistryDTO)

	* 新增成员：CatalogName, DatasourceName, QualifiedName

* [TaskDsDTO](https://cloud.tencent.com/document/api/1267/76336#TaskDsDTO)

	* 新增成员：AllowDownstreamDependency, DependencyTriggerPolicy

* [TaskImportInfo](https://cloud.tencent.com/document/api/1267/76336#TaskImportInfo)

	* 新增成员：WorkFlowFolderPath




