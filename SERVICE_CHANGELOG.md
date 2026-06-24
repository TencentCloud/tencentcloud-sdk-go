# Release v1.3.121

## 腾讯混元生3D(ai3d) 版本：2025-05-13

### 第 16 次发布

发布时间：2026-06-24 01:08:31

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [QueryHunyuan3DPartJob](https://cloud.tencent.com/document/api/1804/126296)

	* 新增出参：PartSegmentationInfo

* [SubmitHunyuan3DPartJob](https://cloud.tencent.com/document/api/1804/126295)

	* 新增入参：PartSegmentationInfo, EnableStagedGeneration




## 费用中心(billing) 版本：2018-07-09

### 第 91 次发布

发布时间：2026-06-24 01:15:23

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [CostComponentSet](https://cloud.tencent.com/document/api/555/19183#CostComponentSet)

	* 新增成员：ComponentCode, ItemCode

* [CostDetail](https://cloud.tencent.com/document/api/555/19183#CostDetail)

	* 新增成员：BusinessCode




## 访问管理(cam) 版本：2019-01-16

### 第 68 次发布

发布时间：2026-06-24 01:17:05

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [ListAccounts](https://cloud.tencent.com/document/api/598/133082)

新增数据结构：

* [ListAllUser](https://cloud.tencent.com/document/api/598/33167#ListAllUser)



## 弹性 MapReduce(emr) 版本：2019-01-03

### 第 145 次发布

发布时间：2026-06-24 01:55:37

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateCloudInstance](https://cloud.tencent.com/document/api/589/113701)

	* 新增入参：EnableSparkAppMonitorInfo

* [DescribeDynamicInstanceDetail](https://cloud.tencent.com/document/api/589/132013)

	* 新增出参：ImageInfoV2

* [DescribeDynamicInstanceList](https://cloud.tencent.com/document/api/589/131177)

	* 新增出参：WebUIInfos


新增数据结构：

* [EnableSparkAppMonitorInfo](https://cloud.tencent.com/document/api/589/33981#EnableSparkAppMonitorInfo)
* [GooseFSVolume](https://cloud.tencent.com/document/api/589/33981#GooseFSVolume)
* [ImageInfoV2](https://cloud.tencent.com/document/api/589/33981#ImageInfoV2)
* [WebUIInfo](https://cloud.tencent.com/document/api/589/33981#WebUIInfo)

修改数据结构：

* [CloudResource](https://cloud.tencent.com/document/api/589/33981#CloudResource)

	* 新增成员：ImageInfoV2, DynamicInstanceForm

* [CustomMetaDBInfo](https://cloud.tencent.com/document/api/589/33981#CustomMetaDBInfo)

	* 新增成员：Components, DefaultMetaVersion, LinkInstanceId

* [DynamicInstanceForm](https://cloud.tencent.com/document/api/589/33981#DynamicInstanceForm)

	* 新增成员：ImageInfoV2, GooseFSVolumes

* [DynamicInstanceGroup](https://cloud.tencent.com/document/api/589/33981#DynamicInstanceGroup)

	* 新增成员：GooseFSVolumes, PreStartCommand, RayStartParams

* [DynamicInstanceGroupSpec](https://cloud.tencent.com/document/api/589/33981#DynamicInstanceGroupSpec)

	* 新增成员：PreStartCommand, RayStartParams

* [ModifyDynamicInstanceForm](https://cloud.tencent.com/document/api/589/33981#ModifyDynamicInstanceForm)

	* 新增成员：CustomImage, ImageInfoV2, GooseFSVolumes

* [NodeSpecInstanceType](https://cloud.tencent.com/document/api/589/33981#NodeSpecInstanceType)

	* 新增成员：NeedHpcClusterId, IsGpuInstance

* [PersistentVolume](https://cloud.tencent.com/document/api/589/33981#PersistentVolume)

	* 新增成员：GooseFSVolumes

* [RayCluster](https://cloud.tencent.com/document/api/589/33981#RayCluster)

	* 新增成员：Namespace




## 高性能应用服务(hai) 版本：2023-08-12

### 第 26 次发布

发布时间：2026-06-24 02:04:23

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [TcpSocketConfig](https://cloud.tencent.com/document/api/1721/101518#TcpSocketConfig)

修改数据结构：

* [ProbeConfig](https://cloud.tencent.com/document/api/1721/101518#ProbeConfig)

	* 新增成员：TcpSocket




## 知识引擎原子能力(lkeap) 版本：2024-05-22

### 第 22 次发布

发布时间：2026-06-24 02:18:54

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [ReconstructDocumentSSEConfig](https://cloud.tencent.com/document/api/1772/115364#ReconstructDocumentSSEConfig)

	* 新增成员：ResultType




## 文字识别(ocr) 版本：2018-11-19

### 第 253 次发布

发布时间：2026-06-24 02:28:08

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除接口**：</font>

* RecognizeContainerOCR



## 容器服务(tke) 版本：2022-05-01



## 容器服务(tke) 版本：2018-05-25

### 第 234 次发布

发布时间：2026-06-24 02:55:39

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyClusterTags](https://cloud.tencent.com/document/api/457/111752)

	* 新增入参：SyncNodePoolTags




## 机器翻译(tmt) 版本：2018-03-21

### 第 16 次发布

发布时间：2026-06-24 02:58:33

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ImageTranslateLLM](https://cloud.tencent.com/document/api/551/118482)

	* 新增入参：Mode




## 向量数据库(vdb) 版本：2023-06-16

### 第 18 次发布

发布时间：2026-06-24 03:06:13

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateInstance](https://cloud.tencent.com/document/api/1709/117430)

	* 新增入参：EnableEncryption




## Web 应用防火墙(waf) 版本：2018-01-25

### 第 159 次发布

发布时间：2026-06-24 03:13:35

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeLLMContentSecCheck](https://cloud.tencent.com/document/api/627/129968)

	* 新增入参：SessionId, IntentContent


新增数据结构：

* [IntentContent](https://cloud.tencent.com/document/api/627/53609#IntentContent)
* [IntentDetectResult](https://cloud.tencent.com/document/api/627/53609#IntentDetectResult)

修改数据结构：

* [LLMDetectResult](https://cloud.tencent.com/document/api/627/53609#LLMDetectResult)

	* 新增成员：IntentDetectResult




