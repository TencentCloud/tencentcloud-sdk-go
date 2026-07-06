# Release v1.3.129

## 日志服务(cls) 版本：2020-10-16

### 第 165 次发布

发布时间：2026-07-07 01:29:54

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeCloudProductLogTasks](https://cloud.tencent.com/document/api/614/117419)

	* 新增入参：WithTags

	* 新增出参：Message

* [ModifyCloudProductLogCollection](https://cloud.tencent.com/document/api/614/117418)

	* 新增入参：Tags

	* 新增出参：Message


新增数据结构：

* [DlcFailHandle](https://cloud.tencent.com/document/api/614/56471#DlcFailHandle)
* [DlcFailTableInfo](https://cloud.tencent.com/document/api/614/56471#DlcFailTableInfo)

修改数据结构：

* [CloudProductLogTaskInfo](https://cloud.tencent.com/document/api/614/56471#CloudProductLogTaskInfo)

	* 新增成员：TopicTags, LogsetTags

* [DlcDeliverInfo](https://cloud.tencent.com/document/api/614/56471#DlcDeliverInfo)

	* 新增成员：AutoCreateField, DlcFailHandle, DSLFilter




## 弹性 MapReduce(emr) 版本：2019-01-03

### 第 147 次发布

发布时间：2026-07-07 01:55:35

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [InquiryPriceCreateInstance](https://cloud.tencent.com/document/api/589/33980)

	* 新增入参：MetaDBGroupInfo

* [InstallSoftware](https://cloud.tencent.com/document/api/589/130892)

	* 新增入参：MetaDBGroupInfo


新增数据结构：

* [DiskHealthIssue](https://cloud.tencent.com/document/api/589/33981#DiskHealthIssue)

修改数据结构：

* [AutoScaleRecord](https://cloud.tencent.com/document/api/589/33981#AutoScaleRecord)

	* 新增成员：ShortageClass

* [ClusterInstancesInfo](https://cloud.tencent.com/document/api/589/33981#ClusterInstancesInfo)

	* 新增成员：IsIOHungSelfRecovery, MetaDBGroupInfo

* [NodeHardwareInfo](https://cloud.tencent.com/document/api/589/33981#NodeHardwareInfo)

	* 新增成员：DiskHealthIssues




## 防火墙管理(fwm) 版本：2025-06-11

### 第 1 次发布

发布时间：2026-07-06 15:58:06

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CancelIgnorePolicyRisk](https://cloud.tencent.com/document/api/1812/133565)
* [CreateAnalyzePolicyTask](https://cloud.tencent.com/document/api/1812/133564)
* [CreateEdgeAclRule](https://cloud.tencent.com/document/api/1812/133582)
* [CreateEdgeAclRuleGroup](https://cloud.tencent.com/document/api/1812/133550)
* [CreateNatAclRule](https://cloud.tencent.com/document/api/1812/133594)
* [CreateNatAclRuleGroup](https://cloud.tencent.com/document/api/1812/133549)
* [CreateSecurityGroupRule](https://cloud.tencent.com/document/api/1812/133576)
* [CreateSecurityGroupRuleGroup](https://cloud.tencent.com/document/api/1812/133548)
* [CreateStrategy](https://cloud.tencent.com/document/api/1812/133558)
* [CreateVpcAclRule](https://cloud.tencent.com/document/api/1812/133588)
* [CreateVpcAclRuleGroup](https://cloud.tencent.com/document/api/1812/133547)
* [DeleteEdgeAclRule](https://cloud.tencent.com/document/api/1812/133581)
* [DeleteNatAclRule](https://cloud.tencent.com/document/api/1812/133593)
* [DeleteRuleGroup](https://cloud.tencent.com/document/api/1812/133546)
* [DeleteSecurityGroupRule](https://cloud.tencent.com/document/api/1812/133575)
* [DeleteStrategy](https://cloud.tencent.com/document/api/1812/133557)
* [DeleteVpcAclRule](https://cloud.tencent.com/document/api/1812/133587)
* [DescribeEdgeAclRules](https://cloud.tencent.com/document/api/1812/133580)
* [DescribeNatAclRules](https://cloud.tencent.com/document/api/1812/133592)
* [DescribeOrganMembers](https://cloud.tencent.com/document/api/1812/133568)
* [DescribeOrganSummary](https://cloud.tencent.com/document/api/1812/133567)
* [DescribePolicyRiskAccountProductStats](https://cloud.tencent.com/document/api/1812/133563)
* [DescribeRiskAnalysisDetails](https://cloud.tencent.com/document/api/1812/133562)
* [DescribeRiskCategoryStats](https://cloud.tencent.com/document/api/1812/133561)
* [DescribeRiskList](https://cloud.tencent.com/document/api/1812/133560)
* [DescribeSecurityGroupRule](https://cloud.tencent.com/document/api/1812/133574)
* [DescribeSecurityGroupRules](https://cloud.tencent.com/document/api/1812/133573)
* [DescribeStrategies](https://cloud.tencent.com/document/api/1812/133556)
* [DescribeStrategy](https://cloud.tencent.com/document/api/1812/133555)
* [DescribeStrategyAccounts](https://cloud.tencent.com/document/api/1812/133543)
* [DescribeStrategyDispatchStatus](https://cloud.tencent.com/document/api/1812/133572)
* [DescribeVpcAclRules](https://cloud.tencent.com/document/api/1812/133586)
* [DispatchStrategy](https://cloud.tencent.com/document/api/1812/133554)
* [IgnorePolicyRisk](https://cloud.tencent.com/document/api/1812/133571)
* [ModifyEdgeAclRule](https://cloud.tencent.com/document/api/1812/133579)
* [ModifyEdgeAclRuleSequence](https://cloud.tencent.com/document/api/1812/133578)
* [ModifyNatAclRule](https://cloud.tencent.com/document/api/1812/133591)
* [ModifyNatAclRuleSequence](https://cloud.tencent.com/document/api/1812/133590)
* [ModifyRuleGroup](https://cloud.tencent.com/document/api/1812/133545)
* [ModifySecurityGroupRule](https://cloud.tencent.com/document/api/1812/133570)
* [ModifyStrategy](https://cloud.tencent.com/document/api/1812/133553)
* [ModifyStrategySequence](https://cloud.tencent.com/document/api/1812/133552)
* [ModifyVpcAclRule](https://cloud.tencent.com/document/api/1812/133585)
* [ModifyVpcAclRuleSequence](https://cloud.tencent.com/document/api/1812/133584)

新增数据结构：

* [Account](https://cloud.tencent.com/document/api/1812/133595#Account)
* [AccountGroupInfo](https://cloud.tencent.com/document/api/1812/133595#AccountGroupInfo)
* [AccountGroupQuotaDetail](https://cloud.tencent.com/document/api/1812/133595#AccountGroupQuotaDetail)
* [AccountProductDetailStats](https://cloud.tencent.com/document/api/1812/133595#AccountProductDetailStats)
* [AccountStatsGroup](https://cloud.tencent.com/document/api/1812/133595#AccountStatsGroup)
* [AddressTemplateSpecification](https://cloud.tencent.com/document/api/1812/133595#AddressTemplateSpecification)
* [AnalysisSgRuleInfoResp](https://cloud.tencent.com/document/api/1812/133595#AnalysisSgRuleInfoResp)
* [CommonFilter](https://cloud.tencent.com/document/api/1812/133595#CommonFilter)
* [EdgeAclRuleInfo](https://cloud.tencent.com/document/api/1812/133595#EdgeAclRuleInfo)
* [EdgeAclRuleResp](https://cloud.tencent.com/document/api/1812/133595#EdgeAclRuleResp)
* [MemberInfo](https://cloud.tencent.com/document/api/1812/133595#MemberInfo)
* [NatAclRule](https://cloud.tencent.com/document/api/1812/133595#NatAclRule)
* [NatAclRuleResp](https://cloud.tencent.com/document/api/1812/133595#NatAclRuleResp)
* [OrganMemberItem](https://cloud.tencent.com/document/api/1812/133595#OrganMemberItem)
* [OrganSummary](https://cloud.tencent.com/document/api/1812/133595#OrganSummary)
* [PolicyRisk](https://cloud.tencent.com/document/api/1812/133595#PolicyRisk)
* [ReceiveAccount](https://cloud.tencent.com/document/api/1812/133595#ReceiveAccount)
* [RiskCategoryItem](https://cloud.tencent.com/document/api/1812/133595#RiskCategoryItem)
* [SecGroupRuleResp](https://cloud.tencent.com/document/api/1812/133595#SecGroupRuleResp)
* [SecurityGroupRiskPolicy](https://cloud.tencent.com/document/api/1812/133595#SecurityGroupRiskPolicy)
* [SecurityGroupRule](https://cloud.tencent.com/document/api/1812/133595#SecurityGroupRule)
* [SequenceIndex](https://cloud.tencent.com/document/api/1812/133595#SequenceIndex)
* [ServiceTemplateSpecification](https://cloud.tencent.com/document/api/1812/133595#ServiceTemplateSpecification)
* [SgDnsParseCount](https://cloud.tencent.com/document/api/1812/133595#SgDnsParseCount)
* [SgRuleResp](https://cloud.tencent.com/document/api/1812/133595#SgRuleResp)
* [StrategyReq](https://cloud.tencent.com/document/api/1812/133595#StrategyReq)
* [StrategyResp](https://cloud.tencent.com/document/api/1812/133595#StrategyResp)
* [VpcAclRule](https://cloud.tencent.com/document/api/1812/133595#VpcAclRule)
* [VpcAclRuleResp](https://cloud.tencent.com/document/api/1812/133595#VpcAclRuleResp)



## 云数据库 KeeWiDB(keewidb) 版本：2022-03-08

### 第 11 次发布

发布时间：2026-07-07 02:13:54

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [InstanceNodeInfo](https://cloud.tencent.com/document/api/1520/86230#InstanceNodeInfo)

	* 新增成员：ZoneId

* [ProxyNodeInfo](https://cloud.tencent.com/document/api/1520/86230#ProxyNodeInfo)

	* 新增成员：ZoneId




## 云直播CSS(live) 版本：2018-08-01

### 第 180 次发布

发布时间：2026-07-07 02:16:43

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateLiveTranscodeTemplate](https://cloud.tencent.com/document/api/267/32646)

	* 新增入参：AudienceDrivenTranscode, AudienceThreshold

* [ModifyLiveTranscodeTemplate](https://cloud.tencent.com/document/api/267/32640)

	* 新增入参：AudienceDrivenTranscode, AudienceThreshold


修改数据结构：

* [TemplateInfo](https://cloud.tencent.com/document/api/267/20474#TemplateInfo)

	* 新增成员：AudienceDrivenTranscode, AudienceThreshold




## 媒体处理(mps) 版本：2019-06-12

### 第 213 次发布

发布时间：2026-07-07 02:23:36

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateAigcImageTask](https://cloud.tencent.com/document/api/862/126966)

	* 新增入参：OutputImageCount




## 邮件推送(ses) 版本：2020-10-02

### 第 37 次发布

发布时间：2026-07-07 02:35:50

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [GetEmailIdentity](https://cloud.tencent.com/document/api/1288/51046)

	* 新增出参：TagList




## 云开发 CloudBase(tcb) 版本：2018-06-08

### 第 150 次发布

发布时间：2026-07-07 02:41:07

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除接口**：</font>

* RollbackPGUserMigrations



## 容器服务(tke) 版本：2022-05-01

### 第 28 次发布

发布时间：2026-07-07 02:57:08

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyClusterMachine](https://cloud.tencent.com/document/api/457/126106)

	* 新增入参：InstanceChargeType, ModifyPortableDataDisk


修改数据结构：

* [MachineSetScaling](https://cloud.tencent.com/document/api/457/103206#MachineSetScaling)

	* 新增成员：ScaleDownMode




## 容器服务(tke) 版本：2018-05-25



## TokenHub(tokenhub) 版本：2026-03-22

### 第 10 次发布

发布时间：2026-07-07 02:57:34

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeModelList](https://cloud.tencent.com/document/api/1823/132614)

	* 新增入参：ModelIds, ModelNames, ModelTypes, Tags, Limit, Offset




## 消息队列 RocketMQ 版(trocket) 版本：2023-03-08

### 第 55 次发布

发布时间：2026-07-07 02:58:14

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateConsumerLabel](https://cloud.tencent.com/document/api/1493/133620)
* [DeleteConsumerLabel](https://cloud.tencent.com/document/api/1493/133619)
* [DeleteConsumerRouteConfig](https://cloud.tencent.com/document/api/1493/133618)
* [DescribeConsumerLabel](https://cloud.tencent.com/document/api/1493/133617)
* [DescribeConsumerLabelList](https://cloud.tencent.com/document/api/1493/133616)
* [DescribeConsumerRouteConfig](https://cloud.tencent.com/document/api/1493/133615)
* [DescribeConsumerRouteVersionList](https://cloud.tencent.com/document/api/1493/133614)
* [PutConsumerRouteConfig](https://cloud.tencent.com/document/api/1493/133613)

新增数据结构：

* [ConsumerLabel](https://cloud.tencent.com/document/api/1493/96031#ConsumerLabel)
* [RouteRule](https://cloud.tencent.com/document/api/1493/96031#RouteRule)
* [RouteRuleVersion](https://cloud.tencent.com/document/api/1493/96031#RouteRuleVersion)



