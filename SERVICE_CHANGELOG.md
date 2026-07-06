# Release v1.3.128

## 腾讯云智能体开发平台(adp) 版本：2026-05-20

### 第 8 次发布

发布时间：2026-07-06 11:28:02

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DeleteAgent](https://cloud.tencent.com/document/api/1759/133531)
* [DescribeAgentSummaryList](https://cloud.tencent.com/document/api/1759/133530)

<font color="#dd0000">**删除接口**：</font>

* DescribeApprovalSkillDetail

新增数据结构：

* [AgentSummary](https://cloud.tencent.com/document/api/1759/132545#AgentSummary)

### 第 7 次发布

发布时间：2026-07-02 20:03:14

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreatePlugin](https://cloud.tencent.com/document/api/1759/133501)
* [DeletePlugin](https://cloud.tencent.com/document/api/1759/133500)
* [FavoritePlugin](https://cloud.tencent.com/document/api/1759/133499)
* [ModifyPlugin](https://cloud.tencent.com/document/api/1759/133498)
* [UnfavoritePlugin](https://cloud.tencent.com/document/api/1759/133497)

### 第 6 次发布

发布时间：2026-07-02 17:53:12

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateSkill](https://cloud.tencent.com/document/api/1759/133496)
* [CreateSkillShare](https://cloud.tencent.com/document/api/1759/133495)
* [DeleteSkill](https://cloud.tencent.com/document/api/1759/133494)
* [DeleteSkillShare](https://cloud.tencent.com/document/api/1759/133493)
* [DescribeApprovalSkillDetail](https://cloud.tencent.com/document/api/#/#)
* [DescribeSkillDetail](https://cloud.tencent.com/document/api/1759/133491)
* [DescribeSkillReferenceList](https://cloud.tencent.com/document/api/1759/133490)
* [FavoriteSkill](https://cloud.tencent.com/document/api/1759/133489)
* [ModifySkill](https://cloud.tencent.com/document/api/1759/133488)
* [ReleaseSkill](https://cloud.tencent.com/document/api/1759/133487)
* [UnfavoriteSkill](https://cloud.tencent.com/document/api/1759/133486)

新增数据结构：

* [SkillDetail](https://cloud.tencent.com/document/api/1759/132545#SkillDetail)
* [SkillReferenceGroup](https://cloud.tencent.com/document/api/1759/132545#SkillReferenceGroup)
* [SkillReferenceSummary](https://cloud.tencent.com/document/api/1759/132545#SkillReferenceSummary)



## 云防火墙(cfw) 版本：2019-09-04

### 第 105 次发布

发布时间：2026-07-06 12:01:54

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeOfflineExportTask](https://cloud.tencent.com/document/api/1132/133527)
* [DescribeOfflineExportTemporaryCredentials](https://cloud.tencent.com/document/api/1132/133526)
* [ExportLogsOffline](https://cloud.tencent.com/document/api/1132/133525)
* [RemoveOfflineExportTask](https://cloud.tencent.com/document/api/1132/133524)

修改接口：

* [DescribeLogStorageStatistic](https://cloud.tencent.com/document/api/1132/120865)

	* 新增入参：StartTime, EndTime

* [DescribeSerialRegion](https://cloud.tencent.com/document/api/1132/131411)

	* 新增出参：EdgeWidth, EdgeElasticSwitch, EdgeElasticBandwidth, EdgeElasticBandwidthLimit, EdgeElasticTrafficSwitch

* [ModifyBlockIgnoreList](https://cloud.tencent.com/document/api/1132/55248)

	* 新增入参：IsFromWeChat


新增数据结构：

* [ClusterFwPreAccessCheckResult](https://cloud.tencent.com/document/api/1132/49071#ClusterFwPreAccessCheckResult)
* [ClusterFwPreAccessCheckStage](https://cloud.tencent.com/document/api/1132/49071#ClusterFwPreAccessCheckStage)
* [OfflineExportTask](https://cloud.tencent.com/document/api/1132/49071#OfflineExportTask)
* [PolicyRoutePreCheckReport](https://cloud.tencent.com/document/api/1132/49071#PolicyRoutePreCheckReport)

修改数据结构：

* [ClusterSwitchDetail](https://cloud.tencent.com/document/api/1132/49071#ClusterSwitchDetail)

	* 新增成员：CheckResult

* [NatFwSwitchDetailS](https://cloud.tencent.com/document/api/1132/49071#NatFwSwitchDetailS)

	* 新增成员：CheckResult




## 日志服务(cls) 版本：2020-10-16

### 第 164 次发布

发布时间：2026-07-06 12:06:35

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateScheduledSql](https://cloud.tencent.com/document/api/614/95138)

	* 新增入参：HasServicesLog, FullQuery

* [ModifyScheduledSql](https://cloud.tencent.com/document/api/614/95518)

	* 新增入参：HasServicesLog, FullQuery




## 云安全一体化平台(csip) 版本：2022-11-21

### 第 88 次发布

发布时间：2026-07-06 12:12:00

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeDspmDictionaryList](https://cloud.tencent.com/document/api/664/131495)

	* 新增入参：MemberId

* [DescribeDspmStatistics](https://cloud.tencent.com/document/api/664/131484)

	* 新增入参：AssetType


修改数据结构：

* [DspmAssetDataScanDetail](https://cloud.tencent.com/document/api/664/90825#DspmAssetDataScanDetail)

	* 新增成员：TaskId

* [DspmDbAsset](https://cloud.tencent.com/document/api/664/90825#DspmDbAsset)

	* 新增成员：ClusterType




## 云服务器(cvm) 版本：2017-03-12

### 第 166 次发布

发布时间：2026-07-06 11:09:34

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateDisasterRecoverGroup](https://cloud.tencent.com/document/api/213/17813)

	* 新增入参：Strategy, PartitionCount

	* 新增出参：Strategy, PartitionCount




## 数据传输服务(dts) 版本：2021-12-06

### 第 56 次发布

发布时间：2026-07-06 12:37:40

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [CompareTableInfo](https://cloud.tencent.com/document/api/571/82108#CompareTableInfo)
* [CompareTableResult](https://cloud.tencent.com/document/api/571/82108#CompareTableResult)

修改数据结构：

* [CompareDetailInfo](https://cloud.tencent.com/document/api/571/82108#CompareDetailInfo)

	* 新增成员：FullProgress, IncDifference




## 数据传输服务(dts) 版本：2018-03-30



## Elasticsearch Service(es) 版本：2025-01-01



## Elasticsearch Service(es) 版本：2018-04-16

### 第 103 次发布

发布时间：2026-07-06 12:44:13

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [InstanceInfo](https://cloud.tencent.com/document/api/845/30634#InstanceInfo)

	* 新增成员：RealLicenseType, EnableAutoReplace, OpenMTLS, ServerCertSource




## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 304 次发布

发布时间：2026-07-06 12:45:36

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateFileConvertTask](https://cloud.tencent.com/document/api/1323/133503)
* [DescribeFileConvertTask](https://cloud.tencent.com/document/api/1323/133502)

修改接口：

* [CreatePreparedPersonalEsign](https://cloud.tencent.com/document/api/1323/89386)




## 腾讯电子签（基础版）(essbasic) 版本：2021-05-26

### 第 262 次发布

发布时间：2026-07-06 12:48:33

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateFileConvertTask](https://cloud.tencent.com/document/api/1420/133505)
* [DescribeFileConvertTask](https://cloud.tencent.com/document/api/1420/133504)

修改接口：

* [ChannelCreatePreparedPersonalEsign](https://cloud.tencent.com/document/api/1420/96160)

* [CreateEmployeeChangeUrl](https://cloud.tencent.com/document/api/1420/116528)

	* 新增出参：LongUrl, ShortUrl




## 腾讯电子签（基础版）(essbasic) 版本：2020-12-22



## 数据加速器 GooseFS(goosefs) 版本：2022-05-19

### 第 18 次发布

发布时间：2026-07-06 12:54:27

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateDataRepositoryTask](https://cloud.tencent.com/document/api/1424/85978)

	* 新增入参：IsTaskPathAbsolute, EnableCustomDestPath, CustomDestPath


修改数据结构：

* [ClientNodeAttribute](https://cloud.tencent.com/document/api/1424/95076#ClientNodeAttribute)

	* 新增成员：ClusterId




## 物联网开发平台(iotexplorer) 版本：2019-04-23

### 第 146 次发布

发布时间：2026-07-06 13:01:35

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [SeeComprehensionConfig](https://cloud.tencent.com/document/api/1081/34988#SeeComprehensionConfig)

	* 新增成员：SummaryPrompt




## 腾讯云可观测平台(monitor) 版本：2023-06-16



## 腾讯云可观测平台(monitor) 版本：2018-07-24

### 第 161 次发布

发布时间：2026-07-06 13:22:22

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**预下线接口**：</font>

* DescribeProductEventList



## 媒体处理(mps) 版本：2019-06-12

### 第 212 次发布

发布时间：2026-07-06 13:25:07

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [AddOnAudio](https://cloud.tencent.com/document/api/862/37615#AddOnAudio)

修改数据结构：

* [AdaptiveDynamicStreamingTaskInput](https://cloud.tencent.com/document/api/862/37615#AdaptiveDynamicStreamingTaskInput)

	* 新增成员：AddOnAudios

* [AddOnSubtitle](https://cloud.tencent.com/document/api/862/37615#AddOnSubtitle)

	* 新增成员：SubtitleLanguage

* [AudioTemplateInfo](https://cloud.tencent.com/document/api/862/37615#AudioTemplateInfo)

	* 新增成员：AudioLanguage, AudioName, DefaultTrack




## 消息队列 MQTT 版(mqtt) 版本：2024-05-16

### 第 31 次发布

发布时间：2026-07-06 13:28:20

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyHttpAuthenticator](https://cloud.tencent.com/document/api/1778/116786)

	* 新增入参：IncludingUserProperties




## 流计算 Oceanus(oceanus) 版本：2019-04-22

### 第 87 次发布

发布时间：2026-07-06 13:30:09

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeJobEvents](https://cloud.tencent.com/document/api/849/102554)

	* 新增入参：Limit, Offset


修改数据结构：

* [JobEvent](https://cloud.tencent.com/document/api/849/52010#JobEvent)

	* 新增成员：CauseAnalysis, Solution




## 文字识别(ocr) 版本：2018-11-19

### 第 256 次发布

发布时间：2026-07-06 13:30:54

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**预下线接口**：</font>

* FormulaOCR
* RecognizeFormulaOCR



## 风险识别 RCE(rce) 版本：2025-04-25

### 第 3 次发布

发布时间：2026-07-06 13:38:54

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [IpLocationInfo](https://cloud.tencent.com/document/api/1343/126967#IpLocationInfo)
* [IpNetworkInfo](https://cloud.tencent.com/document/api/1343/126967#IpNetworkInfo)

修改数据结构：

* [ManageIPPortraitRiskValueOutput](https://cloud.tencent.com/document/api/1343/126967#ManageIPPortraitRiskValueOutput)

	* 新增成员：IpLocation, IpNetwork




## 风险识别 RCE(rce) 版本：2020-11-03



## 凭据管理系统(ssm) 版本：2019-09-23

### 第 17 次发布

发布时间：2026-07-06 13:48:38

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [GetServiceStatus](https://cloud.tencent.com/document/api/1140/40521)

	* 新增出参：ResourceRegion




## 云开发 CloudBase(tcb) 版本：2018-06-08

### 第 149 次发布

发布时间：2026-07-06 13:51:45

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [AllocateEnv](https://cloud.tencent.com/document/api/876/131594)

	* 新增入参：ExternalTag, RequireFunction




## 容器安全服务(tcss) 版本：2020-11-01

### 第 97 次发布

发布时间：2026-07-06 13:55:32

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribePurchaseStateInfo](https://cloud.tencent.com/document/api/1285/65466)

	* 新增出参：NoContainerCoresCnt


修改数据结构：

* [ClusterInfoItem](https://cloud.tencent.com/document/api/1285/65614#ClusterInfoItem)

	* 新增成员：ClusterCAMD5




## 腾讯云数据库 AI 服务(tdai) 版本：2025-07-17

### 第 15 次发布

发布时间：2026-07-06 13:59:59

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateMemoryPlusSpace](https://cloud.tencent.com/document/api/1813/132186)

	* 新增入参：PayMode, PayPeriod, AutoRenew

* [DescribeMemoryPlusSpace](https://cloud.tencent.com/document/api/1813/132184)

	* 新增出参：AutoRenew

* [RecoverMemoryPlusSpace](https://cloud.tencent.com/document/api/1813/132178)

	* 新增入参：PayPeriod




## 容器服务(tke) 版本：2022-05-01



## 容器服务(tke) 版本：2018-05-25

### 第 235 次发布

发布时间：2026-07-06 14:11:21

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [RotateClusterToken](https://cloud.tencent.com/document/api/457/133506)

修改接口：

* [DisableEventPersistence](https://cloud.tencent.com/document/api/457/73905)

	* 新增入参：ClusterType

* [EnableClusterAudit](https://cloud.tencent.com/document/api/457/73900)

	* 新增入参：ClusterType

* [EnableEventPersistence](https://cloud.tencent.com/document/api/457/73904)

	* 新增入参：ClusterType

* [UninstallLogAgent](https://cloud.tencent.com/document/api/457/73902)

	* 新增入参：ClusterType




## TokenHub(tokenhub) 版本：2026-03-22

### 第 9 次发布

发布时间：2026-07-06 14:15:19

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeModelList](https://cloud.tencent.com/document/api/1823/132614)

	* 新增出参：ModelSet, TotalCount


新增数据结构：

* [Model](https://cloud.tencent.com/document/api/1823/132279#Model)
* [ModelAccessInfo](https://cloud.tencent.com/document/api/1823/132279#ModelAccessInfo)
* [ModelChargingInfo](https://cloud.tencent.com/document/api/1823/132279#ModelChargingInfo)
* [ModelChargingItem](https://cloud.tencent.com/document/api/1823/132279#ModelChargingItem)
* [ModelFreeTrialInfo](https://cloud.tencent.com/document/api/1823/132279#ModelFreeTrialInfo)
* [ModelImage](https://cloud.tencent.com/document/api/1823/132279#ModelImage)
* [ModelSiteRegion](https://cloud.tencent.com/document/api/1823/132279#ModelSiteRegion)
* [ModelSpec](https://cloud.tencent.com/document/api/1823/132279#ModelSpec)



## 腾讯混元生视频(vclm) 版本：2024-05-23

### 第 36 次发布

发布时间：2026-07-06 14:24:34

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除接口**：</font>

* CheckAnimateImageJob
* DescribeAigcVideoJob
* DescribeImageAnimateJob
* DescribeVideoEditJob
* DescribeVideoStylizationJob
* DescribeVideoVoiceJob
* SubmitAigcVideoJob
* SubmitImageAnimateJob
* SubmitVideoEditJob
* SubmitVideoStylizationJob
* SubmitVideoVoiceJob

<font color="#dd0000">**删除数据结构**：</font>

* VideoEditParam



## 云点播(vod) 版本：2024-07-18



## 云点播(vod) 版本：2018-07-17

### 第 269 次发布

发布时间：2026-07-06 14:26:30

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeAigcUsageData](https://cloud.tencent.com/document/api/266/126446)

	* 新增入参：ScrollToken, PageSize, ReqId

	* 新增出参：AigcTextDetails


新增数据结构：

* [AigcTextDetail](https://cloud.tencent.com/document/api/266/31773#AigcTextDetail)
* [AigcTextDetailData](https://cloud.tencent.com/document/api/266/31773#AigcTextDetailData)



## 私有网络(vpc) 版本：2017-03-12

### 第 306 次发布

发布时间：2026-07-06 14:30:09

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateBandwidthPackage](https://cloud.tencent.com/document/api/215/19212)

	* 新增入参：DesignatedZone




