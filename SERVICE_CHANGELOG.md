# Release v1.1.49

## 云数据库 MySQL(cdb) 版本：2017-03-20

### 第 207 次发布

发布时间：2025-11-05 01:12:10

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyBackupEncryptionStatus](https://cloud.tencent.com/document/api/236/86507)

	* 新增入参：BinlogEncryptionStatus




## 事件总线(eb) 版本：2021-04-16

### 第 17 次发布

发布时间：2025-11-05 01:23:38

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [CkafkaTargetParams](https://cloud.tencent.com/document/api/1359/67704#CkafkaTargetParams)

	* 新增成员：EventDeliveryFormat




## Elasticsearch Service(es) 版本：2025-01-01



## Elasticsearch Service(es) 版本：2018-04-16

### 第 89 次发布

发布时间：2025-11-05 01:24:39

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [UpdateLogstashInstance](https://cloud.tencent.com/document/api/845/77233)

	* 新增入参：MultiZoneInfo


修改数据结构：

* [InstanceInfo](https://cloud.tencent.com/document/api/845/30634#InstanceInfo)

	* 新增成员：EsPrivateTcpUrl




## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 257 次发布

发布时间：2025-11-05 01:25:07

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [CreateFlowOption](https://cloud.tencent.com/document/api/1323/70369#CreateFlowOption)

	* 新增成员：HideOperationInstructions, HideOperationSteps, SelfName, HideSignCodeAfterStart




## 腾讯电子签（基础版）(essbasic) 版本：2021-05-26

### 第 241 次发布

发布时间：2025-11-05 01:25:50

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [CreateFlowOption](https://cloud.tencent.com/document/api/1420/61525#CreateFlowOption)

	* 新增成员：HideOperationSteps, SelfName, HideSignCodeAfterStart




## 腾讯电子签（基础版）(essbasic) 版本：2020-12-22



## 文字识别(ocr) 版本：2018-11-19

### 第 220 次发布

发布时间：2025-11-05 01:46:42

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [OnlineTaxiItinerary](https://cloud.tencent.com/document/api/866/33527#OnlineTaxiItinerary)

	* <font color="#dd0000">**修改成员**：</font>Content




## 集团账号管理(organization) 版本：2021-03-31

### 第 53 次发布

发布时间：2025-11-05 01:49:26

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateUser](https://cloud.tencent.com/document/api/850/109928)

	* 新增入参：NeedResetPassword

* [UpdateUser](https://cloud.tencent.com/document/api/850/109924)

	* 新增入参：NeedResetPassword


修改数据结构：

* [UserInfo](https://cloud.tencent.com/document/api/850/67060#UserInfo)

	* 新增成员：NeedResetPassword




## 集团账号管理(organization) 版本：2018-12-25



## 容器服务(tke) 版本：2022-05-01

### 第 16 次发布

发布时间：2025-11-05 02:25:36

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeClusters](https://cloud.tencent.com/document/api/457/124673)

新增数据结构：

* [Cluster](https://cloud.tencent.com/document/api/457/103206#Cluster)



## 容器服务(tke) 版本：2018-05-25



## 消息队列 RabbitMQ Serverless 版(trabbit) 版本：2023-04-18

### 第 6 次发布

发布时间：2025-11-05 02:26:32

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeRabbitMQServerlessConnection](https://cloud.tencent.com/document/api/1495/116138)

	* 新增入参：SortElement, SortType, Offset, Limit, Name

* [DescribeRabbitMQServerlessConsumers](https://cloud.tencent.com/document/api/1495/116137)

	* 新增入参：Channel

	* <font color="#dd0000">**修改入参**：</font>VirtualHost, QueueName

* [ModifyRabbitMQServerlessInstance](https://cloud.tencent.com/document/api/1495/116152)

	* 新增入参：DeleteAllTags, InstanceTags

* [ModifyRabbitMQServerlessQueue](https://cloud.tencent.com/document/api/1495/116127)

	* 新增入参：MessageTTL, DeadLetterExchange, DeadLetterRoutingKey


新增数据结构：

* [RabbitMQServerlessTag](https://cloud.tencent.com/document/api/1495/116155#RabbitMQServerlessTag)

修改数据结构：

* [RabbitMQClusterInfo](https://cloud.tencent.com/document/api/1495/116155#RabbitMQClusterInfo)

	* 新增成员：Tags

* [RabbitMQConnection](https://cloud.tencent.com/document/api/1495/116155#RabbitMQConnection)

	* 新增成员：IncomingBytes, OutgoingBytes, Heartbeat, MaxChannel, IdleSince

* [RabbitMQConsumersListInfo](https://cloud.tencent.com/document/api/1495/116155#RabbitMQConsumersListInfo)

	* 新增成员：QueueName, AckRequired, PrefetchCount, Active, LastDeliveredTime, UnAckMsgCount, ChannelName

* [RabbitMQServerlessInstance](https://cloud.tencent.com/document/api/1495/116155#RabbitMQServerlessInstance)

	* 新增成员：Tags




