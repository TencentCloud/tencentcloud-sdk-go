# Release v1.3.157

## 运维安全中心（堡垒机）(bh) 版本：2023-04-18

### 第 35 次发布

发布时间：2026-08-12 01:08:34

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DeployResource](https://cloud.tencent.com/document/api/1025/86961)

	* 新增入参：DeploySubnets, IntranetVpcId, IntranetVpcCidrBlock, IntranetVpcName, IntranetSubnets

* [EnableIntranetAccess](https://cloud.tencent.com/document/api/1025/125052)

	* 新增入参：VpcName, VpcRegion, IntranetSubnets


新增数据结构：

* [ParamInitResourceSubnet](https://cloud.tencent.com/document/api/1025/74416#ParamInitResourceSubnet)
* [ResourceDeployZone](https://cloud.tencent.com/document/api/1025/74416#ResourceDeployZone)

修改数据结构：

* [Resource](https://cloud.tencent.com/document/api/1025/74416#Resource)

	* 新增成员：EnabledDomainCount, IntranetSubnetIdSet, DeployCvmCount, ResourceZoneSet




## 访问管理(cam) 版本：2019-01-16

### 第 72 次发布

发布时间：2026-08-12 01:10:18

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [LoginActionFlagIntl](https://cloud.tencent.com/document/api/598/33167#LoginActionFlagIntl)

	* 新增成员：Passkey




## 配置审计(config) 版本：2022-08-02

### 第 13 次发布

发布时间：2026-08-12 01:29:34

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DeleteAggregators](https://cloud.tencent.com/document/api/1579/135878)
* [UpdateAggregator](https://cloud.tencent.com/document/api/1579/135877)



## 云安全一体化平台(csip) 版本：2022-11-21

### 第 95 次发布

发布时间：2026-08-12 01:30:38

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeAILinkSetting](https://cloud.tencent.com/document/api/664/133464)

	* 新增出参：TagIDs, TCSSScope, ClusterIDs, ExcludeClusterIDs, InstanceIds, ExcludeInstanceIds

* [ModifyAILinkSetting](https://cloud.tencent.com/document/api/664/133463)

	* 新增入参：TagIDs, TCSSScope, ClusterIDs, ExcludeClusterIDs, InstanceIds, ExcludeInstanceIds

* [ModifyEDRRule](https://cloud.tencent.com/document/api/664/133466)

	* 新增入参：TagIDs, ClusterIDsWithAppId, ExcludeClusterIDsWithAppId, ImageIDsWithAppId, ConditionMatches


新增数据结构：

* [ClusterIDWithAppIdItem](https://cloud.tencent.com/document/api/664/90825#ClusterIDWithAppIdItem)
* [ConditionMatch](https://cloud.tencent.com/document/api/664/90825#ConditionMatch)
* [EDRRuleTagItem](https://cloud.tencent.com/document/api/664/90825#EDRRuleTagItem)
* [ImageIDWithAppIdItem](https://cloud.tencent.com/document/api/664/90825#ImageIDWithAppIdItem)

修改数据结构：

* [EDRRule](https://cloud.tencent.com/document/api/664/90825#EDRRule)

	* 新增成员：ClusterIDs, ExcludeClusterIDs, ConditionMatches, TagItems

* [EdrAlertDetail](https://cloud.tencent.com/document/api/664/90825#EdrAlertDetail)

	* 新增成员：ContainerName, ImageName, ClusterName, RunStatus, PodName, PodIp, Namespace, PodWorkloadType, ClusterCaMD5, PodUniqueId

* [EdrAlertItem](https://cloud.tencent.com/document/api/664/90825#EdrAlertItem)

	* 新增成员：MachineType, ContainerName, ImageName, ClusterName




## 数据湖计算 DLC(dlc) 版本：2021-01-25

### 第 173 次发布

发布时间：2026-08-12 01:51:34

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [TCHousePInfo](https://cloud.tencent.com/document/api/1342/53778#TCHousePInfo)

修改数据结构：

* [DatasourceConnectionConfig](https://cloud.tencent.com/document/api/1342/53778#DatasourceConnectionConfig)

	* 新增成员：TCHouseP




## 多网聚合加速(mna) 版本：2021-01-19

### 第 37 次发布

发布时间：2026-08-12 02:44:38

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [GatewayInfo](https://cloud.tencent.com/document/api/1385/55846#GatewayInfo)

	* 新增成员：GatewayIp, Username, Token, RegisterCenterUrl, TelemetryUrl




## 媒体处理(mps) 版本：2019-06-12

### 第 231 次发布

发布时间：2026-08-12 02:49:53

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeCloneViralTask](https://cloud.tencent.com/document/api/862/135032)

	* <font color="#dd0000">**修改入参**：</font>TaskId




## 云数据库 PostgreSQL(postgres) 版本：2017-03-12

### 第 72 次发布

发布时间：2026-08-12 03:00:47

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DeleteDatabase](https://cloud.tencent.com/document/api/409/135858)

	* 新增入参：DBInstanceId, DatabaseName




## 风险识别 RCE(rce) 版本：2026-01-30

### 第 5 次发布

发布时间：2026-08-12 03:04:17

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [AssessRisk](https://cloud.tencent.com/document/api/1343/135880)
* [ReportEvent](https://cloud.tencent.com/document/api/1343/135879)

新增数据结构：

* [AddPromotionEvent](https://cloud.tencent.com/document/api/1343/134560#AddPromotionEvent)
* [Address](https://cloud.tencent.com/document/api/1343/134560#Address)
* [Amount](https://cloud.tencent.com/document/api/1343/134560#Amount)
* [App](https://cloud.tencent.com/document/api/1343/134560#App)
* [AssessRiskRsp](https://cloud.tencent.com/document/api/1343/134560#AssessRiskRsp)
* [Billing](https://cloud.tencent.com/document/api/1343/134560#Billing)
* [BrowseEvent](https://cloud.tencent.com/document/api/1343/134560#BrowseEvent)
* [Browser](https://cloud.tencent.com/document/api/1343/134560#Browser)
* [Card](https://cloud.tencent.com/document/api/1343/134560#Card)
* [ChargeBackEvent](https://cloud.tencent.com/document/api/1343/134560#ChargeBackEvent)
* [ClaimRedPacketEvent](https://cloud.tencent.com/document/api/1343/134560#ClaimRedPacketEvent)
* [Coupon](https://cloud.tencent.com/document/api/1343/134560#Coupon)
* [CreateOrderEvent](https://cloud.tencent.com/document/api/1343/134560#CreateOrderEvent)
* [CreditPoint](https://cloud.tencent.com/document/api/1343/134560#CreditPoint)
* [Cust](https://cloud.tencent.com/document/api/1343/134560#Cust)
* [CustEvent](https://cloud.tencent.com/document/api/1343/134560#CustEvent)
* [DataAuthorization](https://cloud.tencent.com/document/api/1343/134560#DataAuthorization)
* [Delivery](https://cloud.tencent.com/document/api/1343/134560#Delivery)
* [DigitalOrder](https://cloud.tencent.com/document/api/1343/134560#DigitalOrder)
* [EventDetail](https://cloud.tencent.com/document/api/1343/134560#EventDetail)
* [InvitationEvent](https://cloud.tencent.com/document/api/1343/134560#InvitationEvent)
* [Inviter](https://cloud.tencent.com/document/api/1343/134560#Inviter)
* [Item](https://cloud.tencent.com/document/api/1343/134560#Item)
* [LoginEvent](https://cloud.tencent.com/document/api/1343/134560#LoginEvent)
* [LogoutEvent](https://cloud.tencent.com/document/api/1343/134560#LogoutEvent)
* [LuckyDrawEvent](https://cloud.tencent.com/document/api/1343/134560#LuckyDrawEvent)
* [Merchant](https://cloud.tencent.com/document/api/1343/134560#Merchant)
* [ModifyAccountEvent](https://cloud.tencent.com/document/api/1343/134560#ModifyAccountEvent)
* [ModifyPasswordEvent](https://cloud.tencent.com/document/api/1343/134560#ModifyPasswordEvent)
* [Order](https://cloud.tencent.com/document/api/1343/134560#Order)
* [PaymentMethod](https://cloud.tencent.com/document/api/1343/134560#PaymentMethod)
* [PaymentResult](https://cloud.tencent.com/document/api/1343/134560#PaymentResult)
* [Person](https://cloud.tencent.com/document/api/1343/134560#Person)
* [Promotion](https://cloud.tencent.com/document/api/1343/134560#Promotion)
* [PromotionCode](https://cloud.tencent.com/document/api/1343/134560#PromotionCode)
* [RedeemEvent](https://cloud.tencent.com/document/api/1343/134560#RedeemEvent)
* [RegisterEvent](https://cloud.tencent.com/document/api/1343/134560#RegisterEvent)
* [Result](https://cloud.tencent.com/document/api/1343/134560#Result)
* [SMSEvent](https://cloud.tencent.com/document/api/1343/134560#SMSEvent)
* [ScanCodeEvent](https://cloud.tencent.com/document/api/1343/134560#ScanCodeEvent)
* [Score](https://cloud.tencent.com/document/api/1343/134560#Score)
* [SecurityVerificationEvent](https://cloud.tencent.com/document/api/1343/134560#SecurityVerificationEvent)
* [TaskEvent](https://cloud.tencent.com/document/api/1343/134560#TaskEvent)
* [TransactionEvent](https://cloud.tencent.com/document/api/1343/134560#TransactionEvent)
* [User](https://cloud.tencent.com/document/api/1343/134560#User)
* [Wallet](https://cloud.tencent.com/document/api/1343/134560#Wallet)
* [WithdrawEvent](https://cloud.tencent.com/document/api/1343/134560#WithdrawEvent)



## 风险识别 RCE(rce) 版本：2025-04-25



## 风险识别 RCE(rce) 版本：2020-11-03



## 云托管 CloudBase Run(tcbr) 版本：2022-02-17

### 第 30 次发布

发布时间：2026-08-12 03:19:09

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [SubmitServerConfigChangeDiff](https://cloud.tencent.com/document/api/1243/135881)

新增数据结构：

* [ServerBaseConfigDiff](https://cloud.tencent.com/document/api/1243/75713#ServerBaseConfigDiff)



## 容器镜像服务(tcr) 版本：2019-09-24

### 第 81 次发布

发布时间：2026-08-12 03:20:23

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [ModifyReplication](https://cloud.tencent.com/document/api/1141/135882)

新增数据结构：

* [ModifyReplicationRule](https://cloud.tencent.com/document/api/1141/41603#ModifyReplicationRule)



## 边缘安全加速平台(teo) 版本：2022-09-01

### 第 155 次发布

发布时间：2026-08-12 03:31:17

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateInferenceService](https://cloud.tencent.com/document/api/1552/134320)

	* 新增入参：AffinityConfig

* [ModifyInferenceService](https://cloud.tencent.com/document/api/1552/134312)

	* 新增入参：AffinityConfig


新增数据结构：

* [InferenceAffinityConfig](https://cloud.tencent.com/document/api/1552/80721#InferenceAffinityConfig)
* [SessionIdAffinityConfig](https://cloud.tencent.com/document/api/1552/80721#SessionIdAffinityConfig)

修改数据结构：

* [InferenceServiceConfig](https://cloud.tencent.com/document/api/1552/80721#InferenceServiceConfig)

	* 新增成员：AffinityConfig




## 边缘安全加速平台(teo) 版本：2022-01-06



## TI-ONE 训练平台(tione) 版本：2021-11-11

### 第 130 次发布

发布时间：2026-08-12 03:36:04

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateTrainingTask](https://cloud.tencent.com/document/api/851/117377)

	* 新增入参：Queues


修改数据结构：

* [CFSConfig](https://cloud.tencent.com/document/api/851/75051#CFSConfig)

	* 新增成员：IsPresetStorage

* [ResourceSupplyAttribute](https://cloud.tencent.com/document/api/851/75051#ResourceSupplyAttribute)

	* 新增成员：ClusterType




## TI-ONE 训练平台(tione) 版本：2019-10-22



