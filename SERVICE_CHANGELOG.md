# Release v1.3.86

## 云联络中心(ccc) 版本：2020-02-10

### 第 127 次发布

发布时间：2026-04-24 01:12:00

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateAutoCalloutTask](https://cloud.tencent.com/document/api/679/69194)

	* 新增入参：RetryHangupTypes, RetryTags, AvailableWorkTimeConfig, TriggerStrategy


新增数据结构：

* [AvailableTimeConfig](https://cloud.tencent.com/document/api/679/47715#AvailableTimeConfig)
* [BasicAuth](https://cloud.tencent.com/document/api/679/47715#BasicAuth)
* [BearerAuth](https://cloud.tencent.com/document/api/679/47715#BearerAuth)
* [HeaderParams](https://cloud.tencent.com/document/api/679/47715#HeaderParams)
* [HttpCallbackConfig](https://cloud.tencent.com/document/api/679/47715#HttpCallbackConfig)
* [HttpParams](https://cloud.tencent.com/document/api/679/47715#HttpParams)
* [OauthConfig](https://cloud.tencent.com/document/api/679/47715#OauthConfig)
* [RetryTagItem](https://cloud.tencent.com/document/api/679/47715#RetryTagItem)
* [ReturnKey](https://cloud.tencent.com/document/api/679/47715#ReturnKey)
* [TriggerStrategyItem](https://cloud.tencent.com/document/api/679/47715#TriggerStrategyItem)



## 日志服务(cls) 版本：2020-10-16

### 第 158 次发布

发布时间：2026-04-24 01:16:43

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateSearchView](https://cloud.tencent.com/document/api/614/130943)
* [DeleteSearchView](https://cloud.tencent.com/document/api/614/130942)
* [DescribeSearchViews](https://cloud.tencent.com/document/api/614/130941)
* [ModifySearchView](https://cloud.tencent.com/document/api/614/130940)

新增数据结构：

* [SearchViewInfo](https://cloud.tencent.com/document/api/614/56471#SearchViewInfo)
* [ViewSearchTopic](https://cloud.tencent.com/document/api/614/56471#ViewSearchTopic)



## TDSQL-C MySQL 版(cynosdb) 版本：2019-01-07

### 第 162 次发布

发布时间：2026-04-24 01:21:33

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [AssociateSecurityGroups](https://cloud.tencent.com/document/api/1003/70119)

	* <font color="#dd0000">**修改入参**：</font>Zone

* [DisassociateSecurityGroups](https://cloud.tencent.com/document/api/1003/70117)

	* <font color="#dd0000">**修改入参**：</font>Zone

* [ModifyDBInstanceSecurityGroups](https://cloud.tencent.com/document/api/1003/48096)

	* <font color="#dd0000">**修改入参**：</font>Zone


修改数据结构：

* [CynosdbInstanceDetail](https://cloud.tencent.com/document/api/1003/48097#CynosdbInstanceDetail)

	* 新增成员：MasterZone




## 腾讯云智能体开发平台(lke) 版本：2023-11-30

### 第 85 次发布

发布时间：2026-04-24 01:37:18

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateRejectedQuestion](https://cloud.tencent.com/document/api/1759/105102)

	* 新增入参：EnableScope

* [CreateSharedKnowledge](https://cloud.tencent.com/document/api/1759/119409)

	* 新增入参：EsConfig

* [ListRejectedQuestion](https://cloud.tencent.com/document/api/1759/105095)

	* 新增入参：Filters

* [ListWorkflowRuns](https://cloud.tencent.com/document/api/1759/119567)

	* 新增入参：Query


新增数据结构：

* [ESConfig](https://cloud.tencent.com/document/api/1759/105104#ESConfig)
* [FilterItem](https://cloud.tencent.com/document/api/1759/105104#FilterItem)

修改数据结构：

* [MsgRecord](https://cloud.tencent.com/document/api/1759/105104#MsgRecord)

	* 新增成员：OptionMode

* [RejectedQuestion](https://cloud.tencent.com/document/api/1759/105104#RejectedQuestion)

	* 新增成员：EnableScope

* [WorkFlowSummary](https://cloud.tencent.com/document/api/1759/105104#WorkFlowSummary)

	* 新增成员：OptionMode

* [WorkflowInfo](https://cloud.tencent.com/document/api/1759/105104#WorkflowInfo)

	* 新增成员：OptionMode




## 腾讯云可观测平台(monitor) 版本：2023-06-16



## 腾讯云可观测平台(monitor) 版本：2018-07-24

### 第 154 次发布

发布时间：2026-04-24 01:43:13

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [RemoteWrite](https://cloud.tencent.com/document/api/248/30354#RemoteWrite)

	* 新增成员：RemoteWriteType




## 媒体处理(mps) 版本：2019-06-12

### 第 197 次发布

发布时间：2026-04-24 01:45:42

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeTextToSpeechAsyncTask](https://cloud.tencent.com/document/api/862/130945)
* [TextToSpeechAsync](https://cloud.tencent.com/document/api/862/130944)

修改数据结构：

* [AigcImageExtraParam](https://cloud.tencent.com/document/api/862/37615#AigcImageExtraParam)

	* 新增成员：LogoAdd

* [SSAIChannelInfo](https://cloud.tencent.com/document/api/862/37615#SSAIChannelInfo)

	* 新增成员：HlsPlaybackPrefix, DashPlaybackPrefix

* [SSAIConf](https://cloud.tencent.com/document/api/862/37615#SSAIConf)

	* 新增成员：DashOriginManifestType, SlateOnEmptyVast, SCTEMarkerDuration, SecurityGroupId




## 短信(sms) 版本：2021-01-11

### 第 7 次发布

发布时间：2026-04-24 02:03:49

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [SendMultiGlobalSms](https://cloud.tencent.com/document/api/382/130946)

新增数据结构：

* [MultiSmsInfo](https://cloud.tencent.com/document/api/382/52068#MultiSmsInfo)
* [SendMultiStatus](https://cloud.tencent.com/document/api/382/52068#SendMultiStatus)



## 短信(sms) 版本：2019-07-11



## 容器安全服务(tcss) 版本：2020-11-01

### 第 92 次发布

发布时间：2026-04-24 02:14:05

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [AddOrModifyMaliciousConnectionWhiteList](https://cloud.tencent.com/document/api/1285/130947)
* [AddOrModifyVirusWhiteListRule](https://cloud.tencent.com/document/api/1285/130952)
* [DeleteVirusWhiteListRule](https://cloud.tencent.com/document/api/1285/130951)
* [DescribeVirusMonitorConfig](https://cloud.tencent.com/document/api/1285/130950)
* [DescribeVirusScanConfig](https://cloud.tencent.com/document/api/1285/130949)
* [DescribeVirusWhiteListRules](https://cloud.tencent.com/document/api/1285/130948)

新增数据结构：

* [ScanRangeInfo](https://cloud.tencent.com/document/api/1285/65614#ScanRangeInfo)
* [VirusWhiteListRuleInfo](https://cloud.tencent.com/document/api/1285/65614#VirusWhiteListRuleInfo)

修改数据结构：

* [HostInfo](https://cloud.tencent.com/document/api/1285/65614#HostInfo)

	* 新增成员：ClusterAccessedSubStatus, ClusterAccessedErrorReason

* [SuperNodeListItem](https://cloud.tencent.com/document/api/1285/65614#SuperNodeListItem)

	* 新增成员：ClusterAccessedSubStatus, ClusterAccessedErrorReason




## 高性能计算平台(thpc) 版本：2023-03-21

### 第 36 次发布

发布时间：2026-04-24 02:25:51

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [JobView](https://cloud.tencent.com/document/api/1527/89579#JobView)

	* 新增成员：Creator




## 高性能计算平台(thpc) 版本：2022-04-01



## 高性能计算平台(thpc) 版本：2021-11-09



## TI-ONE 训练平台(tione) 版本：2021-11-11

### 第 111 次发布

发布时间：2026-04-24 02:26:47

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeBillingResourceGroupAttachedWorkspaces](https://cloud.tencent.com/document/api/851/130953)



## TI-ONE 训练平台(tione) 版本：2019-10-22



## 容器服务(tke) 版本：2022-05-01



## 容器服务(tke) 版本：2018-05-25

### 第 228 次发布

发布时间：2026-04-24 02:28:23

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateRollOutSequence](https://cloud.tencent.com/document/api/457/125944)

	* 新增出参：ID




## 实时音视频(trtc) 版本：2019-07-22

### 第 141 次发布

发布时间：2026-04-24 02:34:37

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [AsrParam](https://cloud.tencent.com/document/api/647/44055#AsrParam)

	* 新增成员：FilterDirty, FilterModal, FilterPunc




