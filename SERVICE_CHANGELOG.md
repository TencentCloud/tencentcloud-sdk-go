# Release v1.3.17

## 运维安全中心（堡垒机）(bh) 版本：2023-04-18

### 第 22 次发布

发布时间：2025-12-23 01:12:38

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [EnvInternetAccessSetting](https://cloud.tencent.com/document/api/1025/74416#EnvInternetAccessSetting)

修改数据结构：

* [Resource](https://cloud.tencent.com/document/api/1025/74416#Resource)

	* 新增成员：IntranetSubnetId

* [SecuritySetting](https://cloud.tencent.com/document/api/1025/74416#SecuritySetting)

	* 新增成员：EnvInternetAccess




## 腾讯云数据仓库 TCHouse-D(cdwdoris) 版本：2021-12-28

### 第 59 次发布

发布时间：2025-12-23 01:21:16

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateBackUpSchedule](https://cloud.tencent.com/document/api/1387/109543)

	* 新增入参：BucketType, EnableSecurityLock, GracePeriod

* [DescribeBackUpJob](https://cloud.tencent.com/document/api/1387/109558)

	* 新增入参：EncryptionFilters

	* 新增出参：CurrentTime

* [DescribeBackUpSchedules](https://cloud.tencent.com/document/api/1387/109556)

	* 新增入参：EncryptionFilters

	* 新增出参：CurrentTime, BucketEncryption


新增数据结构：

* [BucketEncryptionInfo](https://cloud.tencent.com/document/api/1387/102385#BucketEncryptionInfo)

修改数据结构：

* [BackUpJobDisplay](https://cloud.tencent.com/document/api/1387/102385#BackUpJobDisplay)

	* 新增成员：EnableSecurityLock, GracePeriod, GraceStartTime, IsWithinGracePeriod, UseManagedBucket, InstanceId, InstanceName, InstanceStatus, InstanceStatusDesc, DataRemoteRegion, BucketEncryption, Encryption, EncryptionEnabled

* [BackupStatus](https://cloud.tencent.com/document/api/1387/102385#BackupStatus)

	* 新增成员：UploadBytes

* [InstanceInfo](https://cloud.tencent.com/document/api/1387/102385#InstanceInfo)

	* 新增成员：InstanceType, MasterInstance, SlaveInstances

* [NodeInfo](https://cloud.tencent.com/document/api/1387/102385#NodeInfo)

	* 新增成员：HasFDB

* [NodeInfos](https://cloud.tencent.com/document/api/1387/102385#NodeInfos)

	* 新增成员：HasFDB




## 日志服务(cls) 版本：2020-10-16

### 第 145 次发布

发布时间：2025-12-23 01:27:20

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CommitConsumerOffsets](https://cloud.tencent.com/document/api/614/126716)
* [CreateConsumerGroup](https://cloud.tencent.com/document/api/614/126715)
* [CreateEsRecharge](https://cloud.tencent.com/document/api/614/126743)
* [CreateHostMetricConfig](https://cloud.tencent.com/document/api/614/126735)
* [CreateMetricConfig](https://cloud.tencent.com/document/api/614/126734)
* [CreateMetricSubscribe](https://cloud.tencent.com/document/api/614/126723)
* [CreateSplunkDeliver](https://cloud.tencent.com/document/api/614/126749)
* [DeleteConsumerGroup](https://cloud.tencent.com/document/api/614/126714)
* [DeleteEsRecharge](https://cloud.tencent.com/document/api/614/126742)
* [DeleteHostMetricConfig](https://cloud.tencent.com/document/api/614/126733)
* [DeleteMetricConfig](https://cloud.tencent.com/document/api/614/126732)
* [DeleteMetricSubscribe](https://cloud.tencent.com/document/api/614/126722)
* [DeleteSplunkDeliver](https://cloud.tencent.com/document/api/614/126748)
* [DescribeClusterBaseMetricConfigs](https://cloud.tencent.com/document/api/614/126731)
* [DescribeClusterMetricConfigs](https://cloud.tencent.com/document/api/614/126730)
* [DescribeConsumerGroups](https://cloud.tencent.com/document/api/614/126713)
* [DescribeConsumerOffsets](https://cloud.tencent.com/document/api/614/126712)
* [DescribeConsumerPreview](https://cloud.tencent.com/document/api/614/126737)
* [DescribeConsumers](https://cloud.tencent.com/document/api/614/126736)
* [DescribeEsRechargePreview](https://cloud.tencent.com/document/api/614/126741)
* [DescribeEsRecharges](https://cloud.tencent.com/document/api/614/126740)
* [DescribeHostMetricConfigs](https://cloud.tencent.com/document/api/614/126729)
* [DescribeKafkaConsumerPreview](https://cloud.tencent.com/document/api/614/126751)
* [DescribeKafkaConsumerTopics](https://cloud.tencent.com/document/api/614/126750)
* [DescribeMetricCorrectDimension](https://cloud.tencent.com/document/api/614/126721)
* [DescribeMetricSubscribePreview](https://cloud.tencent.com/document/api/614/126720)
* [DescribeMetricSubscribes](https://cloud.tencent.com/document/api/614/126719)
* [DescribeSplunkDelivers](https://cloud.tencent.com/document/api/614/126747)
* [DescribeSplunkPreview](https://cloud.tencent.com/document/api/614/126746)
* [DescribeTopicBaseMetricConfigs](https://cloud.tencent.com/document/api/614/126728)
* [DescribeTopicMetricConfigs](https://cloud.tencent.com/document/api/614/126727)
* [GetMetricLabelValues](https://cloud.tencent.com/document/api/614/126717)
* [ModifyConsumerGroup](https://cloud.tencent.com/document/api/614/126711)
* [ModifyEsRecharge](https://cloud.tencent.com/document/api/614/126739)
* [ModifyHostMetricConfig](https://cloud.tencent.com/document/api/614/126726)
* [ModifyMetricConfig](https://cloud.tencent.com/document/api/614/126725)
* [ModifyMetricSubscribe](https://cloud.tencent.com/document/api/614/126718)
* [ModifySplunkDeliver](https://cloud.tencent.com/document/api/614/126745)
* [SendConsumerHeartbeat](https://cloud.tencent.com/document/api/614/126710)

新增数据结构：

* [AppointLabel](https://cloud.tencent.com/document/api/614/56471#AppointLabel)
* [BaseMetricCollectConfig](https://cloud.tencent.com/document/api/614/56471#BaseMetricCollectConfig)
* [ConsumerGroupInfo](https://cloud.tencent.com/document/api/614/56471#ConsumerGroupInfo)
* [ConsumerInfo](https://cloud.tencent.com/document/api/614/56471#ConsumerInfo)
* [CustomLabel](https://cloud.tencent.com/document/api/614/56471#CustomLabel)
* [CustomMetricSpec](https://cloud.tencent.com/document/api/614/56471#CustomMetricSpec)
* [Dimension](https://cloud.tencent.com/document/api/614/56471#Dimension)
* [EsImportInfo](https://cloud.tencent.com/document/api/614/56471#EsImportInfo)
* [EsInfo](https://cloud.tencent.com/document/api/614/56471#EsInfo)
* [EsRechargeInfo](https://cloud.tencent.com/document/api/614/56471#EsRechargeInfo)
* [EsTimeInfo](https://cloud.tencent.com/document/api/614/56471#EsTimeInfo)
* [HostMetricConfig](https://cloud.tencent.com/document/api/614/56471#HostMetricConfig)
* [HostMetricItem](https://cloud.tencent.com/document/api/614/56471#HostMetricItem)
* [Instance](https://cloud.tencent.com/document/api/614/56471#Instance)
* [InstanceConfig](https://cloud.tencent.com/document/api/614/56471#InstanceConfig)
* [InstanceData](https://cloud.tencent.com/document/api/614/56471#InstanceData)
* [Label](https://cloud.tencent.com/document/api/614/56471#Label)
* [MetadataInfo](https://cloud.tencent.com/document/api/614/56471#MetadataInfo)
* [MetricCollectConfig](https://cloud.tencent.com/document/api/614/56471#MetricCollectConfig)
* [MetricConfig](https://cloud.tencent.com/document/api/614/56471#MetricConfig)
* [MetricConfigLabel](https://cloud.tencent.com/document/api/614/56471#MetricConfigLabel)
* [MetricSpec](https://cloud.tencent.com/document/api/614/56471#MetricSpec)
* [MetricSubscribeInfo](https://cloud.tencent.com/document/api/614/56471#MetricSubscribeInfo)
* [MetricYamlSpec](https://cloud.tencent.com/document/api/614/56471#MetricYamlSpec)
* [NetInfo](https://cloud.tencent.com/document/api/614/56471#NetInfo)
* [PartitionOffsetInfo](https://cloud.tencent.com/document/api/614/56471#PartitionOffsetInfo)
* [Relabeling](https://cloud.tencent.com/document/api/614/56471#Relabeling)
* [SplunkDeliverInfo](https://cloud.tencent.com/document/api/614/56471#SplunkDeliverInfo)
* [TopicPartitionInfo](https://cloud.tencent.com/document/api/614/56471#TopicPartitionInfo)
* [TopicPartitionOffsetInfo](https://cloud.tencent.com/document/api/614/56471#TopicPartitionOffsetInfo)



## 多媒体创作引擎(cme) 版本：2019-10-29

### 第 62 次发布

发布时间：2025-12-23 01:29:09

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [MediaReplacementInfo](https://cloud.tencent.com/document/api/1156/40360#MediaReplacementInfo)

	* 新增成员：MuteSwitch




## 云安全一体化平台(csip) 版本：2022-11-21

### 第 64 次发布

发布时间：2025-12-23 01:30:01

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeAssetRiskList](https://cloud.tencent.com/document/api/664/120050)

	* 新增出参：StandardNameList

* [DescribeCheckViewRisks](https://cloud.tencent.com/document/api/664/120049)

	* 新增出参：StandardNameList


新增数据结构：

* [StandardItem](https://cloud.tencent.com/document/api/664/90825#StandardItem)
* [StandardTerm](https://cloud.tencent.com/document/api/664/90825#StandardTerm)

修改数据结构：

* [AssetRiskItem](https://cloud.tencent.com/document/api/664/90825#AssetRiskItem)

	* 新增成员：StandardTerms

* [CheckViewRiskItem](https://cloud.tencent.com/document/api/664/90825#CheckViewRiskItem)

	* 新增成员：StandardTerms




## 腾讯电子签（基础版）(essbasic) 版本：2021-05-26

### 第 249 次发布

发布时间：2025-12-23 01:46:46

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [CommonFlowApprover](https://cloud.tencent.com/document/api/1420/61525#CommonFlowApprover)

	* <font color="#dd0000">**修改成员**：</font>NotChannelOrganization




## 腾讯电子签（基础版）(essbasic) 版本：2020-12-22



## 腾讯云智能体开发平台(lke) 版本：2023-11-30

### 第 74 次发布

发布时间：2025-12-23 02:02:22

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateQA](https://cloud.tencent.com/document/api/1759/105037)

	* 新增入参：EnableScope

* [DescribeDoc](https://cloud.tencent.com/document/api/1759/105071)

	* 新增出参：CateBizIdPath, CateNamePath, EnableScope

* [DescribeQA](https://cloud.tencent.com/document/api/1759/105098)

	* 新增出参：CateBizIdPath, CateNamePath, EnableScope, DocEnableScope

* [ListDoc](https://cloud.tencent.com/document/api/1759/105064)

	* 新增入参：EnableScope

* [ListQA](https://cloud.tencent.com/document/api/1759/105028)

	* 新增入参：EnableScope

* [ModifyDoc](https://cloud.tencent.com/document/api/1759/105058)

	* 新增入参：EnableScope

* [ModifyQA](https://cloud.tencent.com/document/api/1759/105025)

	* 新增入参：EnableScope

* [SaveDoc](https://cloud.tencent.com/document/api/1759/105054)

	* 新增入参：EnableScope


修改数据结构：

* [ListDocItem](https://cloud.tencent.com/document/api/1759/105104#ListDocItem)

	* 新增成员：EnableScope

* [ListQaItem](https://cloud.tencent.com/document/api/1759/105104#ListQaItem)

	* 新增成员：EnableScope, DocEnableScope




## 媒体处理(mps) 版本：2019-06-12

### 第 165 次发布

发布时间：2025-12-23 02:13:10

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ProcessImage](https://cloud.tencent.com/document/api/862/112896)

	* 新增入参：ScheduleId, AddOnParameter


新增数据结构：

* [AddOnImageInput](https://cloud.tencent.com/document/api/862/37615#AddOnImageInput)
* [AddOnParameter](https://cloud.tencent.com/document/api/862/37615#AddOnParameter)
* [ImageProcessOutputConfig](https://cloud.tencent.com/document/api/862/37615#ImageProcessOutputConfig)



## 全栈式风控引擎(rce) 版本：2025-04-25

### 第 1 次发布

发布时间：2025-12-22 16:02:53

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [ManageIPPortraitRisk](https://cloud.tencent.com/document/api/1343/126684)



## 全栈式风控引擎(rce) 版本：2020-11-03



## 游戏数据库 TcaplusDB(tcaplusdb) 版本：2019-08-23

### 第 26 次发布

发布时间：2025-12-23 02:43:44

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除接口**：</font>

* RollbackTables

<font color="#dd0000">**删除数据结构**：</font>

* TableRollbackResultNew



## 云托管 CloudBase Run(tcbr) 版本：2022-02-17

### 第 25 次发布

发布时间：2025-12-23 02:46:09

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DeleteCloudRunServer](https://cloud.tencent.com/document/api/1243/126761)
* [DeleteCloudRunVersions](https://cloud.tencent.com/document/api/1243/126760)
* [DescribeCloudRunDeployRecord](https://cloud.tencent.com/document/api/1243/126759)
* [DescribeCloudRunPodList](https://cloud.tencent.com/document/api/1243/126758)
* [DescribeCloudRunProcessLog](https://cloud.tencent.com/document/api/1243/126757)
* [DescribeReleaseOrder](https://cloud.tencent.com/document/api/1243/126756)
* [DescribeVersionDetail](https://cloud.tencent.com/document/api/1243/126755)
* [SearchClsLog](https://cloud.tencent.com/document/api/1243/126754)
* [SubmitServerRollback](https://cloud.tencent.com/document/api/1243/126753)

新增数据结构：

* [DeployRecord](https://cloud.tencent.com/document/api/1243/75713#DeployRecord)
* [FailDeleteVersions](https://cloud.tencent.com/document/api/1243/75713#FailDeleteVersions)
* [LogObject](https://cloud.tencent.com/document/api/1243/75713#LogObject)
* [LogResObject](https://cloud.tencent.com/document/api/1243/75713#LogResObject)
* [ObjectKVPriority](https://cloud.tencent.com/document/api/1243/75713#ObjectKVPriority)
* [ReleaseOrderInfo](https://cloud.tencent.com/document/api/1243/75713#ReleaseOrderInfo)
* [SimpleVersion](https://cloud.tencent.com/document/api/1243/75713#SimpleVersion)
* [SuccessDeleteVersions](https://cloud.tencent.com/document/api/1243/75713#SuccessDeleteVersions)
* [VersionInfo](https://cloud.tencent.com/document/api/1243/75713#VersionInfo)
* [VersionPodInstance](https://cloud.tencent.com/document/api/1243/75713#VersionPodInstance)



## 腾讯云数据库 AI 服务(tdai) 版本：2025-07-17

### 第 7 次发布

发布时间：2025-12-23 02:54:38

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeChatDetail](https://cloud.tencent.com/document/api/1813/123264)

	* <font color="#dd0000">**修改入参**：</font>BeginStreamingTokenId




## 边缘安全加速平台(teo) 版本：2022-09-01

### 第 129 次发布

发布时间：2025-12-23 03:00:08

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreatePrefetchTask](https://cloud.tencent.com/document/api/1552/80704)

	* 新增入参：Mode


修改数据结构：

* [LoadBalancer](https://cloud.tencent.com/document/api/1552/80721#LoadBalancer)

	* 新增成员：References

* [OriginGroupReference](https://cloud.tencent.com/document/api/1552/80721#OriginGroupReference)

	* 新增成员：ZoneId, ZoneName, AliasZoneName




## 边缘安全加速平台(teo) 版本：2022-01-06



## TSF-应用管理&Consul(tsf) 版本：2018-03-26

### 第 135 次发布

发布时间：2025-12-23 03:20:00

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除接口**：</font>

* DescribeInovcationIndicators

<font color="#dd0000">**删除数据结构**：</font>

* IndicatorCoord
* InvocationIndicator



## 腾讯混元生视频(vclm) 版本：2024-05-23

### 第 24 次发布

发布时间：2025-12-22 11:47:57

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeVideoEditJob](https://cloud.tencent.com/document/api/1616/126680)
* [SubmitVideoEditJob](https://cloud.tencent.com/document/api/1616/126679)



## 数据开发治理平台 WeData(wedata) 版本：2025-08-06

### 第 7 次发布

发布时间：2025-12-23 03:48:18

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [CreateTaskBaseAttribute](https://cloud.tencent.com/document/api/1267/123643#CreateTaskBaseAttribute)

	* 新增成员：TaskFolderPath

* [CreateTaskSchedulerConfiguration](https://cloud.tencent.com/document/api/1267/123643#CreateTaskSchedulerConfiguration)

	* 新增成员：ScheduleType, RunPriorityType, RetryWaitMinute, MaxRetryNumber, ExecutionTTLMinute, WaitExecutionTotalTTLMinute

* [TaskBaseAttribute](https://cloud.tencent.com/document/api/1267/123643#TaskBaseAttribute)

	* 新增成员：TaskFolderPath

* [TaskSchedulerConfiguration](https://cloud.tencent.com/document/api/1267/123643#TaskSchedulerConfiguration)

	* 新增成员：DownstreamDependencyConfigList, ScheduleType, RunPriorityType, RetryWaitMinute, MaxRetryNumber, ExecutionTTLMinute, WaitExecutionTotalTTLMinute

* [UpdateTaskBaseAttribute](https://cloud.tencent.com/document/api/1267/123643#UpdateTaskBaseAttribute)

	* 新增成员：TaskFolderPath




## 数据开发治理平台 WeData(wedata) 版本：2021-08-20

### 第 176 次发布

发布时间：2025-12-23 03:41:57

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeFormVersionParam](https://cloud.tencent.com/document/api/1267/118406)

	* 新增入参：Source


新增数据结构：

* [DependencyConfigTimeoutDTO](https://cloud.tencent.com/document/api/1267/76336#DependencyConfigTimeoutDTO)

修改数据结构：

* [Apply](https://cloud.tencent.com/document/api/1267/76336#Apply)

	* 新增成员：CreateTimestamp, ApproveTimestamp

* [BaseClusterInfo](https://cloud.tencent.com/document/api/1267/76336#BaseClusterInfo)

	* 新增成员：CreateTimestamp, UpdateTimestamp

* [BaseProject](https://cloud.tencent.com/document/api/1267/76336#BaseProject)

	* 新增成员：ScheduleMode

* [DependencyStrategyDs](https://cloud.tencent.com/document/api/1267/76336#DependencyStrategyDs)

	* 新增成员：DependencyConfigTimeoutTypeList

* [Project](https://cloud.tencent.com/document/api/1267/76336#Project)

	* 新增成员：CreateTimestamp

* [ProjectUserRole](https://cloud.tencent.com/document/api/1267/76336#ProjectUserRole)

	* 新增成员：CreateTimestamp, Status




