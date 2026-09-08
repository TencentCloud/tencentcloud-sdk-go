# Release v1.3.176

## Agent 沙箱服务(ags) 版本：2025-09-20

### 第 21 次发布

发布时间：2026-09-09 01:10:05

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [OSWorldConfiguration](https://cloud.tencent.com/document/api/1814/124823#OSWorldConfiguration)

修改数据结构：

* [ComputerConfiguration](https://cloud.tencent.com/document/api/1814/124823#ComputerConfiguration)

	* 新增成员：OSWorldConfiguration




## 音频内容安全(ams) 版本：2020-12-29

### 第 20 次发布

发布时间：2026-09-09 01:11:20

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeTaskDetail](https://cloud.tencent.com/document/api/1219/53256)

	* 新增出参：HitSnippetInfos


新增数据结构：

* [Duration](https://cloud.tencent.com/document/api/1219/53259#Duration)
* [HitSnippetInfos](https://cloud.tencent.com/document/api/1219/53259#HitSnippetInfos)

修改数据结构：

* [InputInfo](https://cloud.tencent.com/document/api/1219/53259#InputInfo)

	* 新增成员：Title, Extra

* [SpeakerResults](https://cloud.tencent.com/document/api/1219/53259#SpeakerResults)

	* <font color="#dd0000">**修改成员**：</font>EndTime




## 音频内容安全(ams) 版本：2020-06-08



## 负载均衡(clb) 版本：2018-03-17

### 第 164 次发布

发布时间：2026-09-09 01:32:41

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [TestServiceProviderConnection](https://cloud.tencent.com/document/api/214/133665)

	* 新增入参：EndpointPath




## 日志服务(cls) 版本：2020-10-16

### 第 178 次发布

发布时间：2026-09-08 14:54:08

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateResourceGraph](https://cloud.tencent.com/document/api/614/137648)
* [CreateResourceGraphProductIngestTask](https://cloud.tencent.com/document/api/614/137662)
* [DeleteResourceGraph](https://cloud.tencent.com/document/api/614/137647)
* [DeleteResourceGraphProductIngestTask](https://cloud.tencent.com/document/api/614/137661)
* [DescribeResourceGraphDetail](https://cloud.tencent.com/document/api/614/137646)
* [DescribeResourceGraphEntities](https://cloud.tencent.com/document/api/614/137653)
* [DescribeResourceGraphEntityDependency](https://cloud.tencent.com/document/api/614/137652)
* [DescribeResourceGraphEntityDetail](https://cloud.tencent.com/document/api/614/137651)
* [DescribeResourceGraphFailureDetail](https://cloud.tencent.com/document/api/614/137645)
* [DescribeResourceGraphIngestTaskFailureDetail](https://cloud.tencent.com/document/api/614/137660)
* [DescribeResourceGraphProductIngestTaskDetail](https://cloud.tencent.com/document/api/614/137659)
* [DescribeResourceGraphProductIngestTaskList](https://cloud.tencent.com/document/api/614/137658)
* [DescribeResourceGraphTkeClusterStatus](https://cloud.tencent.com/document/api/614/137657)
* [DescribeResourceGraphs](https://cloud.tencent.com/document/api/614/137644)
* [ModifyResourceGraph](https://cloud.tencent.com/document/api/614/137643)
* [ModifyResourceGraphEntityTopicsRelation](https://cloud.tencent.com/document/api/614/137650)
* [ModifyResourceGraphProductIngestTask](https://cloud.tencent.com/document/api/614/137656)
* [RetryResourceGraph](https://cloud.tencent.com/document/api/614/137642)
* [RetryResourceGraphProductIngestTask](https://cloud.tencent.com/document/api/614/137655)

修改接口：

* [CreateConsumer](https://cloud.tencent.com/document/api/614/66228)

	* 新增入参：DSLFilter

* [DescribeConsumer](https://cloud.tencent.com/document/api/614/66226)

	* 新增出参：CreateTime, RoleArn, ExternalId, TaskStatus, AdvancedConfig, DSLFilter

* [DescribeConsumerOffsets](https://cloud.tencent.com/document/api/614/126712)

	* 新增入参：OffsetType

* [ModifyConsumer](https://cloud.tencent.com/document/api/614/66225)

	* 新增入参：DSLFilter


新增数据结构：

* [EBPFCollectFilters](https://cloud.tencent.com/document/api/614/56471#EBPFCollectFilters)
* [EBPFCollectRule](https://cloud.tencent.com/document/api/614/56471#EBPFCollectRule)
* [EBPFDNSFilter](https://cloud.tencent.com/document/api/614/56471#EBPFDNSFilter)
* [EBPFDestEndpointFilter](https://cloud.tencent.com/document/api/614/56471#EBPFDestEndpointFilter)
* [EBPFEndpoint](https://cloud.tencent.com/document/api/614/56471#EBPFEndpoint)
* [EBPFProcessNameFilter](https://cloud.tencent.com/document/api/614/56471#EBPFProcessNameFilter)
* [ProductIngestTaskDetail](https://cloud.tencent.com/document/api/614/56471#ProductIngestTaskDetail)
* [ProductIngestTaskItem](https://cloud.tencent.com/document/api/614/56471#ProductIngestTaskItem)
* [RelationLogset](https://cloud.tencent.com/document/api/614/56471#RelationLogset)
* [RelationTopic](https://cloud.tencent.com/document/api/614/56471#RelationTopic)
* [ResourceGraphDetailInfo](https://cloud.tencent.com/document/api/614/56471#ResourceGraphDetailInfo)
* [ResourceGraphEntityRelatedTopic](https://cloud.tencent.com/document/api/614/56471#ResourceGraphEntityRelatedTopic)
* [ResourceGraphTkeClusterInfo](https://cloud.tencent.com/document/api/614/56471#ResourceGraphTkeClusterInfo)

修改数据结构：

* [ConsumerContent](https://cloud.tencent.com/document/api/614/56471#ConsumerContent)

	* 新增成员：AutoConvertNumber

* [ConsumerInfo](https://cloud.tencent.com/document/api/614/56471#ConsumerInfo)

	* 新增成员：DSLFilter

* [KafkaConsumerContent](https://cloud.tencent.com/document/api/614/56471#KafkaConsumerContent)

	* 新增成员：AutoConvertNumber




## 配置审计(config) 版本：2022-08-02

### 第 14 次发布

发布时间：2026-09-09 01:38:48

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [OpenConfigRecorder](https://cloud.tencent.com/document/api/1579/129753)

	* 新增入参：FromMode




## 云服务器(cvm) 版本：2017-03-12

### 第 171 次发布

发布时间：2026-09-09 01:40:24

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyInstancesAttribute](https://cloud.tencent.com/document/api/213/15739)

	* 新增入参：EnableJumboFrame




## 主机安全(cwp) 版本：2018-02-28

### 第 169 次发布

发布时间：2026-09-09 01:41:38

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateScanMalwareSetting](https://cloud.tencent.com/document/api/296/58241)

	* 新增入参：CustomPaths

* [DescribeMalwareTimingScanSetting](https://cloud.tencent.com/document/api/296/58240)

	* 新增出参：CustomPaths, MonitorCustomPaths

* [ModifyMalwareTimingScanSettings](https://cloud.tencent.com/document/api/296/52509)

	* 新增入参：CustomPaths, MonitorCustomPaths




## TDSQL-C MySQL 版(cynosdb) 版本：2019-01-07

### 第 194 次发布

发布时间：2026-09-09 01:46:26

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeInstanceSpecs](https://cloud.tencent.com/document/api/1003/48084)

	* 新增入参：Zone




## Elasticsearch Service(es) 版本：2025-01-01



## Elasticsearch Service(es) 版本：2018-04-16

### 第 111 次发布

发布时间：2026-09-09 02:03:44

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [InstanceInfo](https://cloud.tencent.com/document/api/845/30634#InstanceInfo)

	* 新增成员：NeedOfflineOldNodesNotify




## 人脸核身(faceid) 版本：2018-03-01

### 第 92 次发布

发布时间：2026-09-09 02:08:09

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除接口**：</font>

* ImageRecognition



## 流计算 Oceanus(oceanus) 版本：2019-04-22

### 第 93 次发布

发布时间：2026-09-09 02:37:31

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeJobDetail](https://cloud.tencent.com/document/api/849/137733)



## 云数据库 PostgreSQL(postgres) 版本：2017-03-12

### 第 77 次发布

发布时间：2026-09-09 02:41:28

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CloseMem0Service](https://cloud.tencent.com/document/api/409/137740)
* [ClosePostgRESTService](https://cloud.tencent.com/document/api/409/137739)
* [DescribeMem0Service](https://cloud.tencent.com/document/api/409/137738)
* [DescribePostgRESTService](https://cloud.tencent.com/document/api/409/137737)
* [OpenMem0Service](https://cloud.tencent.com/document/api/409/137736)
* [OpenPostgRESTService](https://cloud.tencent.com/document/api/409/137735)

修改数据结构：

* [AccountInfo](https://cloud.tencent.com/document/api/409/16778#AccountInfo)

	* 新增成员：PGRoles




## 云开发 CloudBase(tcb) 版本：2018-06-08

### 第 165 次发布

发布时间：2026-09-09 02:52:37

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeHTTPServiceCachePurgeTask](https://cloud.tencent.com/document/api/876/137742)
* [PurgeHTTPServiceCache](https://cloud.tencent.com/document/api/876/137741)

<font color="#dd0000">**删除接口**：</font>

* CreateVmInstance
* DeleteVmInstance
* DescribeVmInstances
* DescribeVmSpec
* InquireVmPrice

新增数据结构：

* [HTTPServiceCachePurgeTask](https://cloud.tencent.com/document/api/876/34822#HTTPServiceCachePurgeTask)

<font color="#dd0000">**删除数据结构**：</font>

* VMLoginConfiguration
* VMPrice
* VMSpec
* VMSpecLightHouse
* VmInstance



## 高性能计算平台(thpc) 版本：2023-03-21

### 第 45 次发布

发布时间：2026-09-09 03:04:20

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeClusterMonitorStatus](https://cloud.tencent.com/document/api/1527/137745)
* [DisableClusterMonitor](https://cloud.tencent.com/document/api/1527/137744)
* [EnableClusterMonitor](https://cloud.tencent.com/document/api/1527/137743)



## 高性能计算平台(thpc) 版本：2022-04-01



## 高性能计算平台(thpc) 版本：2021-11-09



## 视频内容安全(vm) 版本：2021-09-22

### 第 13 次发布

发布时间：2026-09-09 03:17:40

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeTaskDetail](https://cloud.tencent.com/document/api/1265/80016)

	* 新增出参：HitSnippetInfos


新增数据结构：

* [Duration](https://cloud.tencent.com/document/api/1265/80019#Duration)
* [HitSnippetInfo](https://cloud.tencent.com/document/api/1265/80019#HitSnippetInfo)
* [Position](https://cloud.tencent.com/document/api/1265/80019#Position)
* [Rect](https://cloud.tencent.com/document/api/1265/80019#Rect)

修改数据结构：

* [AudioResult](https://cloud.tencent.com/document/api/1265/80019#AudioResult)

	* 新增成员：HitSnippetInfos

* [ImageResult](https://cloud.tencent.com/document/api/1265/80019#ImageResult)

	* 新增成员：HitSnippetInfos

* [InputInfo](https://cloud.tencent.com/document/api/1265/80019#InputInfo)

	* 新增成员：Extra

* [StorageInfo](https://cloud.tencent.com/document/api/1265/80019#StorageInfo)

	* 新增成员：Extra

* [VideoLLMDetail](https://cloud.tencent.com/document/api/1265/80019#VideoLLMDetail)

	* 新增成员：Extra




## 视频内容安全(vm) 版本：2020-12-29



## 视频内容安全(vm) 版本：2020-07-09



