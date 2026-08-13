# Release v1.3.159

## 费用中心(billing) 版本：2018-07-09

### 第 94 次发布

发布时间：2026-08-14 01:12:49

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeCostDetail](https://cloud.tencent.com/document/api/555/41010)

	* 新增入参：BusinessCode, ProjectId, RegionId

* [DescribeCostSummaryByResource](https://cloud.tencent.com/document/api/555/41006)

	* 新增入参：TagKey, TagValue




## 日志服务(cls) 版本：2020-10-16

### 第 173 次发布

发布时间：2026-08-14 01:20:23

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateAlarmNotice](https://cloud.tencent.com/document/api/614/56465)

	* 新增入参：SecureDetailStatus

* [ModifyAlarmNotice](https://cloud.tencent.com/document/api/614/56458)

	* 新增入参：SecureDetailStatus


修改数据结构：

* [AlarmNotice](https://cloud.tencent.com/document/api/614/56471#AlarmNotice)

	* 新增成员：SecureDetailStatus




## 云安全一体化平台(csip) 版本：2022-11-21

### 第 97 次发布

发布时间：2026-08-14 01:22:20

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CancelEdrAlertIgnore](https://cloud.tencent.com/document/api/664/135933)
* [CreateCSIPManualMalwareScan](https://cloud.tencent.com/document/api/664/135953)
* [CreateEDRManualScan](https://cloud.tencent.com/document/api/664/135965)
* [CreateEdrAlertExportJob](https://cloud.tencent.com/document/api/664/135952)
* [DeleteCSIPMalwareScanTask](https://cloud.tencent.com/document/api/664/135951)
* [DeleteEDRRules](https://cloud.tencent.com/document/api/664/135950)
* [DeleteEDRScanTask](https://cloud.tencent.com/document/api/664/135964)
* [DescribeCSIPMalwareScanTaskDetail](https://cloud.tencent.com/document/api/664/135949)
* [DescribeCSIPMalwareScanTaskProgress](https://cloud.tencent.com/document/api/664/135948)
* [DescribeEDRScanRecordList](https://cloud.tencent.com/document/api/664/135963)
* [DescribeEDRScanTaskDetail](https://cloud.tencent.com/document/api/664/135962)
* [DescribeEdrAlertCountForAsset](https://cloud.tencent.com/document/api/664/135947)
* [DescribeEdrAlertCountForContainer](https://cloud.tencent.com/document/api/664/135946)
* [DescribeEdrAlertMultiAttackStages](https://cloud.tencent.com/document/api/664/135945)
* [DescribeEdrAlertSummary](https://cloud.tencent.com/document/api/664/135944)
* [DescribeEdrExportJobDownloadURL](https://cloud.tencent.com/document/api/664/135943)
* [DescribeEdrExportJobList](https://cloud.tencent.com/document/api/664/135942)
* [DescribeNetAttackSetting](https://cloud.tencent.com/document/api/664/135958)
* [DescribeReverseShellSystemPolicyConfig](https://cloud.tencent.com/document/api/664/135957)
* [ExportCSIPMalwareScanTaskDetail](https://cloud.tencent.com/document/api/664/135941)
* [ExportEDRRules](https://cloud.tencent.com/document/api/664/135940)
* [ModifyEDRRuleStatus](https://cloud.tencent.com/document/api/664/135939)
* [ModifyEDRRulesAction](https://cloud.tencent.com/document/api/664/135938)
* [ModifyEdrAlertIsolation](https://cloud.tencent.com/document/api/664/135937)
* [ModifyEdrAlertStatus](https://cloud.tencent.com/document/api/664/135936)
* [ModifyNetAttackSetting](https://cloud.tencent.com/document/api/664/135956)
* [ModifyReverseShellSystemPolicyConfig](https://cloud.tencent.com/document/api/664/135955)
* [ScanCSIPTaskAgain](https://cloud.tencent.com/document/api/664/135935)
* [ScanEDRTaskAgain](https://cloud.tencent.com/document/api/664/135961)
* [StopCSIPManualMalwareScan](https://cloud.tencent.com/document/api/664/135934)
* [StopEDRScanTask](https://cloud.tencent.com/document/api/664/135960)

新增数据结构：

* [CSIPMachineExtraInfo](https://cloud.tencent.com/document/api/664/90825#CSIPMachineExtraInfo)
* [CSIPMalwareScanUuidDetailItem](https://cloud.tencent.com/document/api/664/90825#CSIPMalwareScanUuidDetailItem)
* [ClusterWithAppIdItem](https://cloud.tencent.com/document/api/664/90825#ClusterWithAppIdItem)
* [CreatedTaskItem](https://cloud.tencent.com/document/api/664/90825#CreatedTaskItem)
* [EDRExportJobItem](https://cloud.tencent.com/document/api/664/90825#EDRExportJobItem)
* [EDRFilters](https://cloud.tencent.com/document/api/664/90825#EDRFilters)
* [EDRScanRecordItem](https://cloud.tencent.com/document/api/664/90825#EDRScanRecordItem)
* [EDRScanTaskContainerItem](https://cloud.tencent.com/document/api/664/90825#EDRScanTaskContainerItem)
* [EDRScanTaskHostItem](https://cloud.tencent.com/document/api/664/90825#EDRScanTaskHostItem)
* [EdrAlertCountItem](https://cloud.tencent.com/document/api/664/90825#EdrAlertCountItem)
* [EdrContainerAlertCountItem](https://cloud.tencent.com/document/api/664/90825#EdrContainerAlertCountItem)
* [EdrContainerGlobalCount](https://cloud.tencent.com/document/api/664/90825#EdrContainerGlobalCount)
* [MultiAttackStageItem](https://cloud.tencent.com/document/api/664/90825#MultiAttackStageItem)



## 云服务器(cvm) 版本：2017-03-12

### 第 170 次发布

发布时间：2026-08-14 01:24:32

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [ChcHost](https://cloud.tencent.com/document/api/213/15753#ChcHost)

	* 新增成员：ChcGatewayId, DedicatedClusterId, NetworkMode




## TDSQL-C MySQL 版(cynosdb) 版本：2019-01-07

### 第 186 次发布

发布时间：2026-08-14 01:27:14

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [CynosdbClusterDetail](https://cloud.tencent.com/document/api/1003/48097#CynosdbClusterDetail)

	* 新增成员：RealZone

* [CynosdbInstance](https://cloud.tencent.com/document/api/1003/48097#CynosdbInstance)

	* 新增成员：RealZone




## 腾讯云数据分析智能体(dataagent) 版本：2025-05-13

### 第 22 次发布

发布时间：2026-08-14 01:28:23

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [FileInfo](https://cloud.tencent.com/document/api/1800/125016#FileInfo)

	* 新增成员：EnableGraphBuild, EnableTreeBuild, GraphBuildStatus, TreeBuildStatus

* [KnowledgeTaskConfig](https://cloud.tencent.com/document/api/1800/125016#KnowledgeTaskConfig)

	* 新增成员：EnableGraphBuild, EnableTreeBuild

* [SearchConfig](https://cloud.tencent.com/document/api/1800/125016#SearchConfig)

	* 新增成员：EnableGraphSearch, EnableTreeSearch




## 数据湖计算 DLC(dlc) 版本：2021-01-25

### 第 174 次发布

发布时间：2026-08-14 01:29:54

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [ListImages](https://cloud.tencent.com/document/api/1342/135967)

新增数据结构：

* [ImageDto](https://cloud.tencent.com/document/api/1342/53778#ImageDto)

修改数据结构：

* [ResourceSaleInfo](https://cloud.tencent.com/document/api/1342/53778#ResourceSaleInfo)

	* 新增成员：StatusCategory




## TI-ONE 训练平台(tione) 版本：2021-11-11

### 第 132 次发布

发布时间：2026-08-14 02:02:56

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeExport](https://cloud.tencent.com/document/api/851/124490)

	* 新增入参：TiProjectId




## TI-ONE 训练平台(tione) 版本：2019-10-22



## 文本内容安全(tms) 版本：2020-12-29

### 第 18 次发布

发布时间：2026-08-14 02:04:55

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [TextModeration](https://cloud.tencent.com/document/api/1124/51860)

	* 新增入参：BizTag


修改数据结构：

* [DetailResults](https://cloud.tencent.com/document/api/1124/51861#DetailResults)

	* 新增成员：HitSnippetInfos




## 文本内容安全(tms) 版本：2020-07-13



## Web 应用防火墙(waf) 版本：2018-01-25

### 第 162 次发布

发布时间：2026-08-14 02:11:59

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [ClbObject](https://cloud.tencent.com/document/api/627/53609#ClbObject)

	* 新增成员：LLMStatus




