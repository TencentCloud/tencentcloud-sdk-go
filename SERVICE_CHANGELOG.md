# Release v1.3.160

## 应用性能监控(apm) 版本：2021-06-22

### 第 67 次发布

发布时间：2026-08-17 01:12:03

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyApmApplicationConfig](https://cloud.tencent.com/document/api/1463/125072)

	* 新增入参：CrossAccountStatus, CrossAccountPeerId

* [ModifyApmInstance](https://cloud.tencent.com/document/api/1463/89002)

	* 新增入参：CrossAccountStatus, CrossAccountPeerId


修改数据结构：

* [ApmAppConfig](https://cloud.tencent.com/document/api/1463/64927#ApmAppConfig)

	* 新增成员：CrossAccountStatus, CrossAccountPeerId

* [ApmInstanceDetail](https://cloud.tencent.com/document/api/1463/64927#ApmInstanceDetail)

	* 新增成员：CrossAccountStatus, CrossAccountPeerId




## 负载均衡(clb) 版本：2018-03-17

### 第 158 次发布

发布时间：2026-08-17 01:29:20

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateModelRouter](https://cloud.tencent.com/document/api/214/133217)

	* 新增入参：EipAddressId, Bandwidth


新增数据结构：

* [ModelRouterBillingConfigOutput](https://cloud.tencent.com/document/api/214/30694#ModelRouterBillingConfigOutput)
* [StickyConfig](https://cloud.tencent.com/document/api/214/30694#StickyConfig)

修改数据结构：

* [ModelRouterDetail](https://cloud.tencent.com/document/api/214/30694#ModelRouterDetail)

	* 新增成员：BillingConfig

* [ModelRouterSet](https://cloud.tencent.com/document/api/214/30694#ModelRouterSet)

	* 新增成员：BillingConfig

* [RouterSettingWithFallBack](https://cloud.tencent.com/document/api/214/30694#RouterSettingWithFallBack)

	* 新增成员：StickyConfig

* [RouterSettingWithoutFallBack](https://cloud.tencent.com/document/api/214/30694#RouterSettingWithoutFallBack)

	* 新增成员：StickyConfig




## 数据库智能管家 DBbrain(dbbrain) 版本：2021-05-27

### 第 60 次发布

发布时间：2026-08-17 01:48:29

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeDBDiagEvents](https://cloud.tencent.com/document/api/1130/65947)

	* 新增入参：DiagItems




## 数据库智能管家 DBbrain(dbbrain) 版本：2019-10-16



## 数据湖计算 DLC(dlc) 版本：2021-01-25

### 第 175 次发布

发布时间：2026-08-17 01:51:02

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateInferenceModel](https://cloud.tencent.com/document/api/1342/135526)

	* 新增入参：ResourceTags, GooseFSConfig, StorageType

	* 新增出参：ResourceTags

* [CreateInferenceService](https://cloud.tencent.com/document/api/1342/135625)

	* 新增入参：AdvancedOptions, ResourceTags, IsCustom, RuntimeEnv

	* 新增出参：AdvancedOptions, ResourceTags, DeploymentMode, IsCustom

* [CreateModelVersion](https://cloud.tencent.com/document/api/1342/135618)

	* 新增入参：GooseFSConfig, StorageType

* [DescribeMCPTaskResult](https://cloud.tencent.com/document/api/1342/134616)

	* 新增入参：NextToken

* [GetInferenceModel](https://cloud.tencent.com/document/api/1342/135525)

	* 新增出参：ResourceTags

* [GetInferenceService](https://cloud.tencent.com/document/api/1342/135624)

	* 新增出参：DeploymentMode, IsCustom, ResourceTags

* [RestartInferenceService](https://cloud.tencent.com/document/api/1342/135621)

	* 新增出参：ResourceTags

* [StopInferenceService](https://cloud.tencent.com/document/api/1342/135620)

	* 新增出参：ResourceTags, DeploymentMode, IsCustom

* [UpdateInferenceModel](https://cloud.tencent.com/document/api/1342/135523)

	* 新增入参：ResourceTags

	* 新增出参：ResourceTags


新增数据结构：

* [GooseFSConfig](https://cloud.tencent.com/document/api/1342/53778#GooseFSConfig)

修改数据结构：

* [GpuSummaryItem](https://cloud.tencent.com/document/api/1342/53778#GpuSummaryItem)

	* 新增成员：GpuType, GpuCount, Replicas

* [InferenceModelInfo](https://cloud.tencent.com/document/api/1342/53778#InferenceModelInfo)

	* 新增成员：ResourceTags




## 物联网开发平台(iotexplorer) 版本：2019-04-23

### 第 153 次发布

发布时间：2026-08-17 02:13:07

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateTWeSeeDirectUploadCredential](https://cloud.tencent.com/document/api/1081/133452)

	* 新增入参：UploadTarget

* [ListTWeSeeTasks](https://cloud.tencent.com/document/api/1081/132116)

	* 新增入参：Filters


新增数据结构：

* [VisionRecognitionTaskFilter](https://cloud.tencent.com/document/api/1081/34988#VisionRecognitionTaskFilter)

修改数据结构：

* [SeeTaskInfo](https://cloud.tencent.com/document/api/1081/34988#SeeTaskInfo)

	* 新增成员：COSURI




## 多网聚合加速(mna) 版本：2021-01-19

### 第 38 次发布

发布时间：2026-08-17 02:26:15

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [AddCustomerGatewayCluster](https://cloud.tencent.com/document/api/1385/136014)
* [AddGateway](https://cloud.tencent.com/document/api/1385/136013)
* [DeleteCustomerGatewayCluster](https://cloud.tencent.com/document/api/1385/136012)
* [DeleteGateway](https://cloud.tencent.com/document/api/1385/136011)
* [DescribeAccessPointList](https://cloud.tencent.com/document/api/1385/136010)
* [GetCustomerGatewayClusterList](https://cloud.tencent.com/document/api/1385/136009)
* [ModifyDeviceAccessScope](https://cloud.tencent.com/document/api/1385/136008)
* [UpdateCustomerGatewayCluster](https://cloud.tencent.com/document/api/1385/136007)

新增数据结构：

* [AccessPointInfo](https://cloud.tencent.com/document/api/1385/55846#AccessPointInfo)
* [GatewayClusterInfo](https://cloud.tencent.com/document/api/1385/55846#GatewayClusterInfo)



## TokenHub(tokenhub) 版本：2026-03-22

### 第 18 次发布

发布时间：2026-08-17 03:05:54

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [Model](https://cloud.tencent.com/document/api/1823/132279#Model)

	* 新增成员：ExtraModelIds

* [ModelChargingItem](https://cloud.tencent.com/document/api/1823/132279#ModelChargingItem)

	* 新增成员：Specification, Usage, ReferencePrice

* [ModelEndpointView](https://cloud.tencent.com/document/api/1823/132279#ModelEndpointView)

	* 新增成员：ExtraModelIds, ModelStatus




