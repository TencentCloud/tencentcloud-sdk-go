# Release v1.3.140

## 腾讯云智能体开发平台(adp) 版本：2026-05-20

### 第 11 次发布

发布时间：2026-07-22 01:08:33

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeConversationMessageList](https://cloud.tencent.com/document/api/1759/132517)

	* 新增出参：ResetInfo

* [DescribePlugin](https://cloud.tencent.com/document/api/1759/132500)

	* 新增入参：Module


新增数据结构：

* [ConversationResetInfo](https://cloud.tencent.com/document/api/1759/132545#ConversationResetInfo)
* [ToolSummary](https://cloud.tencent.com/document/api/1759/132545#ToolSummary)

修改数据结构：

* [PluginSummary](https://cloud.tencent.com/document/api/1759/132545#PluginSummary)

	* 新增成员：ToolList




## 云硬盘(cbs) 版本：2017-03-12

### 第 77 次发布

发布时间：2026-07-22 01:20:26

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyRemoteDiskAttributes](https://cloud.tencent.com/document/api/362/132680)

	* 新增入参：AutoRenewFlag


修改数据结构：

* [Snapshot](https://cloud.tencent.com/document/api/362/15669#Snapshot)

	* 新增成员：SnapshotMode




## 云防火墙(cfw) 版本：2019-09-04

### 第 109 次发布

发布时间：2026-07-22 01:27:18

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeAclRegInfo](https://cloud.tencent.com/document/api/1132/134684)

新增数据结构：

* [AclRegInfo](https://cloud.tencent.com/document/api/1132/49071#AclRegInfo)



## 数据湖计算 DLC(dlc) 版本：2021-01-25

### 第 165 次发布

发布时间：2026-07-22 01:52:02

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeMCPTask](https://cloud.tencent.com/document/api/1342/134617)

	* 新增入参：TaskId

	* 新增出参：TaskInfo

* [DescribeMCPTaskResult](https://cloud.tencent.com/document/api/1342/134616)

	* 新增入参：TaskId

	* 新增出参：TaskResult


新增数据结构：

* [MCPTaskInfo](https://cloud.tencent.com/document/api/1342/53778#MCPTaskInfo)
* [MCPTaskResultInfo](https://cloud.tencent.com/document/api/1342/53778#MCPTaskResultInfo)



## 全球加速(ga2) 版本：2025-01-15

### 第 12 次发布

发布时间：2026-07-22 02:06:45

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateGlobalAcceleratorAccessLog](https://cloud.tencent.com/document/api/1817/134692)
* [DeleteGlobalAcceleratorAccessLog](https://cloud.tencent.com/document/api/1817/134691)
* [DescribeAccessLogParam](https://cloud.tencent.com/document/api/1817/134690)
* [DescribeGlobalAcceleratorAccessLog](https://cloud.tencent.com/document/api/1817/134689)
* [DescribeGlobalAcceleratorAclPolicies](https://cloud.tencent.com/document/api/1817/134688)
* [DescribeGlobalAcceleratorAclRules](https://cloud.tencent.com/document/api/1817/134687)
* [ModifyAccessLogStatus](https://cloud.tencent.com/document/api/1817/134686)
* [ModifyGlobalAcceleratorAccessLog](https://cloud.tencent.com/document/api/1817/134685)

新增数据结构：

* [GlobalAcceleratorAccessLog](https://cloud.tencent.com/document/api/1817/130045#GlobalAcceleratorAccessLog)
* [GlobalAcceleratorAclPolicies](https://cloud.tencent.com/document/api/1817/130045#GlobalAcceleratorAclPolicies)
* [GlobalAcceleratorAclRuleSet](https://cloud.tencent.com/document/api/1817/130045#GlobalAcceleratorAclRuleSet)

修改数据结构：

* [ForwardingRuleSet](https://cloud.tencent.com/document/api/1817/130045#ForwardingRuleSet)

	* 新增成员：HideResponseHeaders, ResponseHeaders




## 云直播CSS(live) 版本：2018-08-01

### 第 182 次发布

发布时间：2026-07-22 02:22:34

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [AvatarScriptInfo](https://cloud.tencent.com/document/api/267/20474#AvatarScriptInfo)

	* 新增成员：MediaUrl




## 媒体处理(mps) 版本：2019-06-12

### 第 221 次发布

发布时间：2026-07-22 02:29:33

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [AiRestorationConfig](https://cloud.tencent.com/document/api/862/37615#AiRestorationConfig)

修改数据结构：

* [AiAnalysisTaskDelLogoOutput](https://cloud.tencent.com/document/api/862/37615#AiAnalysisTaskDelLogoOutput)

	* 新增成员：FileId, OriginSubtitleFileId, TranslateSubtitleFileId

* [AiAnalysisTaskDubbingOutput](https://cloud.tencent.com/document/api/862/37615#AiAnalysisTaskDubbingOutput)

	* 新增成员：VideoFileId, SpeakerFileId

* [ImageProcessTaskOutput](https://cloud.tencent.com/document/api/862/37615#ImageProcessTaskOutput)

	* 新增成员：FileId

* [SubtitleResult](https://cloud.tencent.com/document/api/862/37615#SubtitleResult)

	* 新增成员：SubtitleFileId

* [SubtitleTransResultItem](https://cloud.tencent.com/document/api/862/37615#SubtitleTransResultItem)

	* 新增成员：SubtitleFileId

* [VODInputInfo](https://cloud.tencent.com/document/api/862/37615#VODInputInfo)

	* 新增成员：VodBasic, FileId

* [VODOutputStorage](https://cloud.tencent.com/document/api/862/37615#VODOutputStorage)

	* 新增成员：VodBasic

* [VideoEnhanceConfig](https://cloud.tencent.com/document/api/862/37615#VideoEnhanceConfig)

	* 新增成员：AiRestoration




## 流计算 Oceanus(oceanus) 版本：2019-04-22

### 第 89 次发布

发布时间：2026-07-22 02:33:28

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [SqlGatewayEndpoint](https://cloud.tencent.com/document/api/849/52010#SqlGatewayEndpoint)

修改数据结构：

* [SqlGatewayItem](https://cloud.tencent.com/document/api/849/52010#SqlGatewayItem)

	* 新增成员：SessionClusterId, PgUser, Endpoints




## 风险识别 RCE(rce) 版本：2026-01-30

### 第 2 次发布

发布时间：2026-07-22 02:38:11

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [AssessDeviceRiskPremiumPro](https://cloud.tencent.com/document/api/1343/134694)
* [AssessDeviceRiskPro](https://cloud.tencent.com/document/api/1343/134693)

新增数据结构：

* [AssessDeviceRiskPremiumRsp](https://cloud.tencent.com/document/api/1343/134560#AssessDeviceRiskPremiumRsp)
* [AssessDeviceRiskRsp](https://cloud.tencent.com/document/api/1343/134560#AssessDeviceRiskRsp)
* [Decision](https://cloud.tencent.com/document/api/1343/134560#Decision)
* [Device](https://cloud.tencent.com/document/api/1343/134560#Device)



## 风险识别 RCE(rce) 版本：2025-04-25



## 风险识别 RCE(rce) 版本：2020-11-03



## TI-ONE 训练平台(tione) 版本：2021-11-11

### 第 124 次发布

发布时间：2026-07-22 02:59:07

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DeleteModelService](https://cloud.tencent.com/document/api/851/82290)

	* 新增入参：TiProjectId

* [DeleteModelServiceGroup](https://cloud.tencent.com/document/api/851/82289)

	* 新增入参：TiProjectId

* [DescribeModelServiceCallInfo](https://cloud.tencent.com/document/api/851/82286)

	* 新增入参：TiProjectId

* [ModifyModelService](https://cloud.tencent.com/document/api/851/83228)

	* 新增入参：TiProjectId


修改数据结构：

* [Service](https://cloud.tencent.com/document/api/851/75051#Service)

	* 新增成员：ProjectId




## TI-ONE 训练平台(tione) 版本：2019-10-22



