# Release v1.1.16

## 应用性能监控(apm) 版本：2021-06-22

### 第 44 次发布

发布时间：2025-08-26 01:10:09

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [ApmApplicationConfigView](https://cloud.tencent.com/document/api/1463/64927#ApmApplicationConfigView)

	* 新增成员：DisableMemoryUsed, DisableCpuUsed




## 数据传输服务(dts) 版本：2021-12-06

### 第 44 次发布

发布时间：2025-08-26 01:42:30

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ConfigureSubscribeJob](https://cloud.tencent.com/document/api/571/102952)

	* 新增入参：ConsumerVpcId, ConsumerSubnetId

* [CreateSubscribe](https://cloud.tencent.com/document/api/571/102950)

	* 新增入参：InstanceClass

* [DescribeSubscribeDetail](https://cloud.tencent.com/document/api/571/102944)

	* 新增出参：SubscribeVersion, ConsumerVpcId, ConsumerSubnetId, InstanceClass


修改数据结构：

* [SubscribeInfo](https://cloud.tencent.com/document/api/571/82108#SubscribeInfo)

	* 新增成员：InstanceClass




## 数据传输服务(dts) 版本：2018-03-30



## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 245 次发布

发布时间：2025-08-26 01:47:18

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeContractReviewTask](https://cloud.tencent.com/document/api/1323/122151)

	* 新增出参：HighRiskCount, TotalRiskCount




## 数据开发治理平台 WeData(wedata) 版本：2021-08-20

### 第 169 次发布

发布时间：2025-08-25 10:47:55

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeTestRunningRecord](https://cloud.tencent.com/document/api/1267/122750)

<font color="#dd0000">**删除接口**：</font>

* DescribeDrInstancePage
* DescribeTaskByCycleReport

修改接口：

* [CommitIntegrationTask](https://cloud.tencent.com/document/api/1267/82526)

	* 新增出参：DataDto

* [DescribeDsTaskVersionList](https://cloud.tencent.com/document/api/1267/120168)

	* 新增入参：IsOnlyProductVersion

* [DescribeIntegrationTask](https://cloud.tencent.com/document/api/1267/82495)

	* 新增入参：ExtConfig

* [RenewWorkflowSchedulerInfoDs](https://cloud.tencent.com/document/api/1267/110374)

	* 新增入参：MainCyclicConfig, SubordinateCyclicConfig

* [RunRerunScheduleInstances](https://cloud.tencent.com/document/api/1267/100207)

	* 新增入参：MapParamList, AppParam


新增数据结构：

* [AlarmReceiverGroup](https://cloud.tencent.com/document/api/1267/76336#AlarmReceiverGroup)
* [CommitTaskDataDto](https://cloud.tencent.com/document/api/1267/76336#CommitTaskDataDto)
* [TestRunningRecord](https://cloud.tencent.com/document/api/1267/76336#TestRunningRecord)
* [TestRunningSubRecord](https://cloud.tencent.com/document/api/1267/76336#TestRunningSubRecord)
* [WorkspaceExt](https://cloud.tencent.com/document/api/1267/76336#WorkspaceExt)

<font color="#dd0000">**删除数据结构**：</font>

* DrInstanceOpsDto
* DrInstanceOpsDtoPage

修改数据结构：

* [DataSourceInfo](https://cloud.tencent.com/document/api/1267/76336#DataSourceInfo)

	* 新增成员：DatasourceType

* [DescribePendingSubmitTaskInfo](https://cloud.tencent.com/document/api/1267/76336#DescribePendingSubmitTaskInfo)

	* 新增成员：TaskTypeId

* [EngineTaskInfo](https://cloud.tencent.com/document/api/1267/76336#EngineTaskInfo)

	* 新增成员：VCoreSeconds, MemorySeconds, EmrUserName, QueryId, ApplicationId

* [Project](https://cloud.tencent.com/document/api/1267/76336#Project)

	* 新增成员：WorkspaceExt

* [TaskAlarmInfo](https://cloud.tencent.com/document/api/1267/76336#TaskAlarmInfo)

	* 新增成员：AlarmReceiverGroups, AlarmReceiverGroupFlag

* [TaskDsDTO](https://cloud.tencent.com/document/api/1267/76336#TaskDsDTO)

	* 新增成员：BundleId, BundleName

* [UploadResourceRequestInfo](https://cloud.tencent.com/document/api/1267/76336#UploadResourceRequestInfo)

	* 新增成员：RemotePath




