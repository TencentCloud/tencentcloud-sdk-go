# Release v1.3.16

## Agent 沙箱服务(ags) 版本：2025-09-20

### 第 2 次发布

发布时间：2025-12-22 01:08:09

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateSandboxTool](https://cloud.tencent.com/document/api/1814/124812)

	* 新增入参：RoleArn, StorageMounts

* [StartSandboxInstance](https://cloud.tencent.com/document/api/1814/124816)

	* 新增入参：MountOptions


新增数据结构：

* [CosStorageSource](https://cloud.tencent.com/document/api/1814/124823#CosStorageSource)
* [MountOption](https://cloud.tencent.com/document/api/1814/124823#MountOption)
* [StorageMount](https://cloud.tencent.com/document/api/1814/124823#StorageMount)
* [StorageSource](https://cloud.tencent.com/document/api/1814/124823#StorageSource)
* [VPCConfig](https://cloud.tencent.com/document/api/1814/124823#VPCConfig)

修改数据结构：

* [NetworkConfiguration](https://cloud.tencent.com/document/api/1814/124823#NetworkConfiguration)

	* 新增成员：VpcConfig

* [SandboxInstance](https://cloud.tencent.com/document/api/1814/124823#SandboxInstance)

	* 新增成员：MountOptions

* [SandboxTool](https://cloud.tencent.com/document/api/1814/124823#SandboxTool)

	* 新增成员：RoleArn, StorageMounts




## 弹性伸缩(as) 版本：2018-04-19

### 第 91 次发布

发布时间：2025-12-22 01:11:12

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateScheduledAction](https://cloud.tencent.com/document/api/377/20452)

	* 新增入参：DisableUpdateDesiredCapacity

* [ModifyScheduledAction](https://cloud.tencent.com/document/api/377/20449)

	* 新增入参：DisableUpdateDesiredCapacity


修改数据结构：

* [ScheduledAction](https://cloud.tencent.com/document/api/377/20453#ScheduledAction)

	* 新增成员：DisableUpdateDesiredCapacity




## 消息队列 CKafka 版(ckafka) 版本：2019-08-19

### 第 140 次发布

发布时间：2025-12-22 01:25:12

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateInstancePre](https://cloud.tencent.com/document/api/597/45847)

	* 新增入参：CustomSSLCertId

* [CreatePostPaidInstance](https://cloud.tencent.com/document/api/597/94259)

	* 新增入参：CustomSSLCertId

* [DescribeCkafkaVersion](https://cloud.tencent.com/document/api/597/125479)

	* 新增出参：Result


新增数据结构：

* [InstanceVersion](https://cloud.tencent.com/document/api/597/40861#InstanceVersion)
* [LatestBrokerVersion](https://cloud.tencent.com/document/api/597/40861#LatestBrokerVersion)



## 弹性 MapReduce(emr) 版本：2019-01-03

### 第 126 次发布

发布时间：2025-12-22 01:43:30

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeHBaseTableRequestMetric](https://cloud.tencent.com/document/api/589/126601)
* [DescribeHBaseTableStoreSizeMetric](https://cloud.tencent.com/document/api/589/126600)

新增数据结构：

* [HBaseMetricData](https://cloud.tencent.com/document/api/589/33981#HBaseMetricData)



## 云游戏(gs) 版本：2019-11-18

### 第 62 次发布

发布时间：2025-12-22 01:50:13

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeAndroidInstanceBackups](https://cloud.tencent.com/document/api/1162/123530)

	* 新增出参：Backups, TotalCount


新增数据结构：

* [AndroidInstanceBackup](https://cloud.tencent.com/document/api/1162/40743#AndroidInstanceBackup)



## 容器镜像服务(tcr) 版本：2019-09-24

### 第 72 次发布

发布时间：2025-12-22 02:46:36

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DeleteReplicationRule](https://cloud.tencent.com/document/api/1141/126603)
* [DescribeReplicationPolicies](https://cloud.tencent.com/document/api/1141/126602)

新增数据结构：

* [PolicyFilter](https://cloud.tencent.com/document/api/1141/41603#PolicyFilter)
* [ReplicationPolicyInfo](https://cloud.tencent.com/document/api/1141/41603#ReplicationPolicyInfo)



## 腾讯云数据库 AI 服务(tdai) 版本：2025-07-17

### 第 6 次发布

发布时间：2025-12-22 02:53:53

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [StartAgentTask](https://cloud.tencent.com/document/api/1813/126604)



## 容器服务(tke) 版本：2022-05-01

### 第 20 次发布

发布时间：2025-12-22 03:10:39

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeClusterMachines](https://cloud.tencent.com/document/api/457/126605)

新增数据结构：

* [Machine](https://cloud.tencent.com/document/api/457/103206#Machine)



## 容器服务(tke) 版本：2018-05-25



## TSF-应用管理&Consul(tsf) 版本：2018-03-26

### 第 134 次发布

发布时间：2025-12-22 03:18:45

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [BusinessLogConfig](https://cloud.tencent.com/document/api/649/36099#BusinessLogConfig)

	* 新增成员：CustomMultilinePattern




## 云点播(vod) 版本：2024-07-18



## 云点播(vod) 版本：2018-07-17

### 第 213 次发布

发布时间：2025-12-22 03:25:10

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [AigcImageTaskInputFileInfo](https://cloud.tencent.com/document/api/266/31773#AigcImageTaskInputFileInfo)

	* 新增成员：Text




