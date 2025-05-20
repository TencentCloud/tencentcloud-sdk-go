# Release v1.0.1170

## 云数据库 MySQL(cdb) 版本：2017-03-20

### 第 193 次发布

发布时间：2025-05-21 01:10:22

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [AnalyzeAuditLogs](https://cloud.tencent.com/document/api/236/89112)

* [CreateAuditLogFile](https://cloud.tencent.com/document/api/236/45461)

* [ModifyAuditService](https://cloud.tencent.com/document/api/236/101809)

* [OpenAuditService](https://cloud.tencent.com/document/api/236/76408)




## TDSQL-C MySQL 版(cynosdb) 版本：2019-01-07

### 第 134 次发布

发布时间：2025-05-21 01:15:11

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [CynosdbCluster](https://cloud.tencent.com/document/api/1003/48097#CynosdbCluster)

	* 新增成员：CynosVersionTag, GdnRole




## Elasticsearch Service(es) 版本：2025-01-01



## Elasticsearch Service(es) 版本：2018-04-16

### 第 81 次发布

发布时间：2025-05-21 01:18:51

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateLogstashInstance](https://cloud.tencent.com/document/api/845/77244)

	* 新增入参：MultiZoneInfo, DeployMode


修改数据结构：

* [LogstashInstanceInfo](https://cloud.tencent.com/document/api/845/30634#LogstashInstanceInfo)

	* 新增成员：DeployMode, MultiZoneInfo

* [ZoneDetail](https://cloud.tencent.com/document/api/845/30634#ZoneDetail)

	* 新增成员：Hidden




## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 222 次发布

发布时间：2025-05-21 01:19:08

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateUserNameChangeUrl](https://cloud.tencent.com/document/api/1323/118524)



## 知识引擎原子能力(lkeap) 版本：2024-05-22

### 第 11 次发布

发布时间：2025-05-21 01:24:26

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [RetrieveKnowledgeRealtime](https://cloud.tencent.com/document/api/1772/118525)

<font color="#dd0000">**删除接口**：</font>

* UploadDocRealtime



## 媒体处理(mps) 版本：2019-06-12

### 第 123 次发布

发布时间：2025-05-21 01:25:35

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateTranscodeTemplate](https://cloud.tencent.com/document/api/862/37605)

	* 新增入参：StdExtInfo

* [DescribeImageTaskDetail](https://cloud.tencent.com/document/api/862/118509)

	* 新增出参：ImageProcessTaskResultSet


新增数据结构：

* [ImageDenoiseConfig](https://cloud.tencent.com/document/api/862/37615#ImageDenoiseConfig)
* [ImageProcessTaskOutput](https://cloud.tencent.com/document/api/862/37615#ImageProcessTaskOutput)
* [ImageProcessTaskResult](https://cloud.tencent.com/document/api/862/37615#ImageProcessTaskResult)

修改数据结构：

* [ImageEnhanceConfig](https://cloud.tencent.com/document/api/862/37615#ImageEnhanceConfig)

	* 新增成员：Denoise, LowLightEnhance

* [RawTranscodeParameter](https://cloud.tencent.com/document/api/862/37615#RawTranscodeParameter)

	* 新增成员：StdExtInfo, EnhanceConfig




## 医疗报告结构化(mrs) 版本：2020-09-10

### 第 35 次发布

发布时间：2025-05-21 01:26:13

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [Block](https://cloud.tencent.com/document/api/1314/56230#Block)

	* 新增成员：EndoscopyV2

* [IndicatorItem](https://cloud.tencent.com/document/api/1314/56230#IndicatorItem)

	* 新增成员：Sample

* [Template](https://cloud.tencent.com/document/api/1314/56230#Template)

	* 新增成员：EndoscopyV2




## 云数据库 PostgreSQL(postgres) 版本：2017-03-12

### 第 56 次发布

发布时间：2025-05-21 01:27:30

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除接口**：</font>

* InitDBInstances



## 边缘安全加速平台(teo) 版本：2022-09-01

### 第 104 次发布

发布时间：2025-05-21 01:32:21

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreatePrefetchTask](https://cloud.tencent.com/document/api/1552/80704)

	* 新增入参：PrefetchMediaSegments


修改数据结构：

* [OriginDetail](https://cloud.tencent.com/document/api/1552/80721#OriginDetail)

	* 新增成员：HostHeader

* [OriginInfo](https://cloud.tencent.com/document/api/1552/80721#OriginInfo)

	* 新增成员：HostHeader




## 边缘安全加速平台(teo) 版本：2022-01-06



## 容器服务(tke) 版本：2022-05-01

### 第 10 次发布

发布时间：2025-05-21 01:33:41

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [NativeNodeInfo](https://cloud.tencent.com/document/api/457/103206#NativeNodeInfo)

	* 新增成员：SystemDisk, WanIp, KeyIds, GPUParams, DataDisks, MachineType




## 容器服务(tke) 版本：2018-05-25



## Web 应用防火墙(waf) 版本：2018-01-25

### 第 120 次发布

发布时间：2025-05-21 01:37:18

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [AddCustomWhiteRule](https://cloud.tencent.com/document/api/627/90325)

	* 新增入参：LogicalOp

* [ModifyCustomWhiteRule](https://cloud.tencent.com/document/api/627/90323)

	* 新增入参：LogicalOp

* [UpsertCCRule](https://cloud.tencent.com/document/api/627/97646)

	* 新增入参：CelRule, LogicalOp

	* <font color="#dd0000">**修改入参**：</font>Url, MatchFunc


修改数据结构：

* [CCRuleItems](https://cloud.tencent.com/document/api/627/53609#CCRuleItems)

	* 新增成员：CelRule, LogicalOp




