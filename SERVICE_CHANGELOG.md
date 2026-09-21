# Release v1.3.184

## 腾讯混元生图(aiart) 版本：2022-12-29

### 第 31 次发布

发布时间：2026-09-22 01:09:53

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除接口**：</font>

* ImageToImage
* QueryTextToImageJob
* SubmitTextToImageJob
* TextToImageLite
* TextToImageRapid

<font color="#dd0000">**删除数据结构**：</font>

* Image
* ResultConfig



## 负载均衡(clb) 版本：2018-03-17

### 第 166 次发布

发布时间：2026-09-22 01:19:26

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyModelAliasAttributes](https://cloud.tencent.com/document/api/214/133671)

	* 新增入参：CoefficientTiers, CoefficientSchedule

	* <font color="#dd0000">**修改入参**：</font>Coefficient


新增数据结构：

* [CoefficientScheduleRule](https://cloud.tencent.com/document/api/214/30694#CoefficientScheduleRule)
* [CoefficientTier](https://cloud.tencent.com/document/api/214/30694#CoefficientTier)
* [CoefficientTierCondition](https://cloud.tencent.com/document/api/214/30694#CoefficientTierCondition)

修改数据结构：

* [Coefficient](https://cloud.tencent.com/document/api/214/30694#Coefficient)

	* 新增成员：InputImageCoefficient, InputVideoSecondCoefficient, OutputVideoSecondCoefficient

* [ModelAlias](https://cloud.tencent.com/document/api/214/30694#ModelAlias)

	* 新增成员：CoefficientTiers, CoefficientSchedule

* [ServiceProviderCoefficient](https://cloud.tencent.com/document/api/214/30694#ServiceProviderCoefficient)

	* 新增成员：CoefficientTiers, CoefficientSchedule




## 云安全一体化平台(csip) 版本：2022-11-21

### 第 105 次发布

发布时间：2026-09-22 01:22:32

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [BackupLog](https://cloud.tencent.com/document/api/664/90825#BackupLog)

	* 新增成员：InstanceId, InstanceName, AssetType




## 大数据智能体工作台DataBuddy(databuddy) 版本：2026-07-15

### 第 4 次发布

发布时间：2026-09-22 01:29:45

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateConsoleGroup](https://cloud.tencent.com/document/api/1835/138543)
* [DeleteConsoleGroups](https://cloud.tencent.com/document/api/1835/138542)
* [ListConsoleGroupUsers](https://cloud.tencent.com/document/api/1835/138541)
* [ListConsoleGroups](https://cloud.tencent.com/document/api/1835/138540)
* [ListConsoleRoles](https://cloud.tencent.com/document/api/1835/138539)
* [UpdateConsoleGroup](https://cloud.tencent.com/document/api/1835/138538)

新增数据结构：

* [ConsoleGroupInfo](https://cloud.tencent.com/document/api/1835/138006#ConsoleGroupInfo)
* [ConsoleGroupUserInfo](https://cloud.tencent.com/document/api/1835/138006#ConsoleGroupUserInfo)
* [ConsoleRoleInfo](https://cloud.tencent.com/document/api/1835/138006#ConsoleRoleInfo)
* [CreateConsoleGroupRsp](https://cloud.tencent.com/document/api/1835/138006#CreateConsoleGroupRsp)
* [DeleteConsoleGroupsRsp](https://cloud.tencent.com/document/api/1835/138006#DeleteConsoleGroupsRsp)
* [ListConsoleGroupUsersRsp](https://cloud.tencent.com/document/api/1835/138006#ListConsoleGroupUsersRsp)
* [ListConsoleGroupsRsp](https://cloud.tencent.com/document/api/1835/138006#ListConsoleGroupsRsp)
* [ListConsoleRolesRsp](https://cloud.tencent.com/document/api/1835/138006#ListConsoleRolesRsp)
* [RoleMetaData](https://cloud.tencent.com/document/api/1835/138006#RoleMetaData)
* [RolePermission](https://cloud.tencent.com/document/api/1835/138006#RolePermission)
* [UpdateConsoleGroupRsp](https://cloud.tencent.com/document/api/1835/138006#UpdateConsoleGroupRsp)



## 数据库智能管家 DBbrain(dbbrain) 版本：2021-05-27

### 第 63 次发布

发布时间：2026-09-22 01:30:08

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeTopSpaceTablesV2](https://cloud.tencent.com/document/api/1130/138545)

新增数据结构：

* [MongoCollectionDetail](https://cloud.tencent.com/document/api/1130/57812#MongoCollectionDetail)
* [MongoDBTableSpaceItem](https://cloud.tencent.com/document/api/1130/57812#MongoDBTableSpaceItem)
* [MysqlSpaceObjectItem](https://cloud.tencent.com/document/api/1130/57812#MysqlSpaceObjectItem)
* [PostgresSpaceObjectItem](https://cloud.tencent.com/document/api/1130/57812#PostgresSpaceObjectItem)

修改数据结构：

* [SlowLogInfoItem](https://cloud.tencent.com/document/api/1130/57812#SlowLogInfoItem)

	* 新增成员：InstanceId

* [SlowLogTopSqlItem](https://cloud.tencent.com/document/api/1130/57812#SlowLogTopSqlItem)

	* 新增成员：SqlType, InstanceId




## 数据库智能管家 DBbrain(dbbrain) 版本：2019-10-16



## 数据湖计算 DLC(dlc) 版本：2021-01-25

### 第 182 次发布

发布时间：2026-09-22 01:31:23

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [SparkJobInfo](https://cloud.tencent.com/document/api/1342/53778#SparkJobInfo)

	* 新增成员：DependencyPackages, RunAsIdentity




## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 322 次发布

发布时间：2026-09-22 01:37:05

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [ComparisonDetail](https://cloud.tencent.com/document/api/1323/70369#ComparisonDetail)

	* 新增成员：PageNumber

* [PdfVerifyResult](https://cloud.tencent.com/document/api/1323/70369#PdfVerifyResult)

	* 新增成员：CertProvider, IsTimestampTrust




## 高性能应用服务(hai) 版本：2023-08-12

### 第 29 次发布

发布时间：2026-09-22 01:40:19

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [GetServicePodLogs](https://cloud.tencent.com/document/api/1721/138547)



## 智能视图计算平台(iss) 版本：2023-05-17

### 第 34 次发布

发布时间：2026-09-22 01:44:12

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [UpdateUserDevice](https://cloud.tencent.com/document/api/1344/95867)

	* 新增入参：TimeSyncSwitch


新增数据结构：

* [SipCarrierEndpoints](https://cloud.tencent.com/document/api/1344/95952#SipCarrierEndpoints)

修改数据结构：

* [DescribeDeviceData](https://cloud.tencent.com/document/api/1344/95952#DescribeDeviceData)

	* 新增成员：SipFQDN, SipCarrierEndpoints, TimeSyncSwitch




## 媒体处理(mps) 版本：2019-06-12

### 第 250 次发布

发布时间：2026-09-22 01:49:28

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [BeautyConfig](https://cloud.tencent.com/document/api/862/37615#BeautyConfig)

	* 新增成员：Type

* [ImageEraseLogoConfig](https://cloud.tencent.com/document/api/862/37615#ImageEraseLogoConfig)

	* 新增成员：EraseStrength, WatermarkType




## 文字识别(ocr) 版本：2018-11-19

### 第 268 次发布

发布时间：2026-09-22 01:51:33

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [PassInvoiceInfo](https://cloud.tencent.com/document/api/866/33527#PassInvoiceInfo)

	* <font color="#dd0000">**删除成员**：</font>CarType, PlateNumber




## 边缘安全加速平台(teo) 版本：2022-09-01

### 第 164 次发布

发布时间：2026-09-22 02:01:52

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeZoneCustomVariables](https://cloud.tencent.com/document/api/1552/138549)
* [ModifyZoneCustomVariables](https://cloud.tencent.com/document/api/1552/138548)



## 边缘安全加速平台(teo) 版本：2022-01-06



## TI-ONE 训练平台(tione) 版本：2021-11-11

### 第 136 次发布

发布时间：2026-09-22 02:03:35

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateModelService](https://cloud.tencent.com/document/api/851/82291)

	* 新增入参：Priority

* [ModifyModelService](https://cloud.tencent.com/document/api/851/83228)

	* 新增入参：Priority


修改数据结构：

* [ResourceInfo](https://cloud.tencent.com/document/api/851/75051#ResourceInfo)

	* 新增成员：RdmaNumber, Rdma




## TI-ONE 训练平台(tione) 版本：2019-10-22



## 消息队列 RocketMQ 版(trocket) 版本：2023-03-08

### 第 58 次发布

发布时间：2026-09-22 02:05:59

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateConsumerLabels](https://cloud.tencent.com/document/api/1493/138556)
* [DeleteConsumerLabels](https://cloud.tencent.com/document/api/1493/138555)
* [DeleteConsumerRouteConfigs](https://cloud.tencent.com/document/api/1493/138554)
* [DescribeConsumerLabelLists](https://cloud.tencent.com/document/api/1493/138553)
* [DescribeConsumerLabelRoutes](https://cloud.tencent.com/document/api/1493/138552)
* [DescribeConsumerRouteConfigs](https://cloud.tencent.com/document/api/1493/138551)
* [PutConsumerRouteConfigs](https://cloud.tencent.com/document/api/1493/138550)

新增数据结构：

* [ConsumerLabelFailure](https://cloud.tencent.com/document/api/1493/96031#ConsumerLabelFailure)
* [ConsumerLabelItem](https://cloud.tencent.com/document/api/1493/96031#ConsumerLabelItem)
* [ConsumerLabelKey](https://cloud.tencent.com/document/api/1493/96031#ConsumerLabelKey)
* [ConsumerLabelList](https://cloud.tencent.com/document/api/1493/96031#ConsumerLabelList)
* [ConsumerLabelRoute](https://cloud.tencent.com/document/api/1493/96031#ConsumerLabelRoute)
* [ConsumerLabelRouteItem](https://cloud.tencent.com/document/api/1493/96031#ConsumerLabelRouteItem)
* [ConsumerRouteKey](https://cloud.tencent.com/document/api/1493/96031#ConsumerRouteKey)
* [ConsumerRouteLabelKey](https://cloud.tencent.com/document/api/1493/96031#ConsumerRouteLabelKey)
* [DeleteConsumerRouteConfigFailure](https://cloud.tencent.com/document/api/1493/96031#DeleteConsumerRouteConfigFailure)
* [DescribeConsumerRouteConfigItem](https://cloud.tencent.com/document/api/1493/96031#DescribeConsumerRouteConfigItem)
* [ErrorInfo](https://cloud.tencent.com/document/api/1493/96031#ErrorInfo)
* [PutConsumerRouteConfigFailure](https://cloud.tencent.com/document/api/1493/96031#PutConsumerRouteConfigFailure)
* [PutConsumerRouteConfigItem](https://cloud.tencent.com/document/api/1493/96031#PutConsumerRouteConfigItem)



## 私有网络(vpc) 版本：2017-03-12

### 第 312 次发布

发布时间：2026-09-22 02:10:22

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateAndAttachNetworkInterface](https://cloud.tencent.com/document/api/215/43370)

	* 新增入参：Ipv6Addresses, Ipv6AddressCount, TerminationProtection, TrafficProtection

* [CreateNetworkInterface](https://cloud.tencent.com/document/api/215/15818)

	* 新增入参：Ipv6Addresses, Ipv6AddressCount, TerminationProtection, TrafficProtection

* [CreateSubnet](https://cloud.tencent.com/document/api/215/15782)

	* 新增入参：StackType, Ipv6CidrBlock

	* <font color="#dd0000">**修改入参**：</font>CidrBlock

* [CreateVpc](https://cloud.tencent.com/document/api/215/15774)

	* 新增入参：StackType, Ipv6CidrBlock, AddressType

	* <font color="#dd0000">**修改入参**：</font>CidrBlock

* [SetCcnRegionBandwidthLimits](https://cloud.tencent.com/document/api/215/19194)

	* 新增入参：SetQosDefaultLimitFlag


修改数据结构：

* [CrossBorderCompliance](https://cloud.tencent.com/document/api/215/15824#CrossBorderCompliance)

	* 新增成员：WhiteListFlag

* [NetworkInterface](https://cloud.tencent.com/document/api/215/15824#NetworkInterface)

	* 新增成员：Ipv6Addresses, Ipv6AddressCount

* [Subnet](https://cloud.tencent.com/document/api/215/15824#Subnet)

	* 新增成员：StackType

* [SubnetInput](https://cloud.tencent.com/document/api/215/15824#SubnetInput)

	* 新增成员：StackType, Ipv6CidrBlock

	* <font color="#dd0000">**修改成员**：</font>CidrBlock

* [Vpc](https://cloud.tencent.com/document/api/215/15824#Vpc)

	* 新增成员：StackType




