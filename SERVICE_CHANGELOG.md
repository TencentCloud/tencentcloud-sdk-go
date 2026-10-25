# Release v1.3.199

## 负载均衡(clb) 版本：2018-03-17

### 第 173 次发布

发布时间：2026-10-26 01:32:34

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyModelRouterAttributes](https://cloud.tencent.com/document/api/214/133203)

	* 新增入参：RedirectConfig, Schema, Port


新增数据结构：

* [RedirectConfig](https://cloud.tencent.com/document/api/214/30694#RedirectConfig)

修改数据结构：

* [IntentRouterDecisionModelConfig](https://cloud.tencent.com/document/api/214/30694#IntentRouterDecisionModelConfig)

	* 新增成员：ModelName

* [ModelRouterDetail](https://cloud.tencent.com/document/api/214/30694#ModelRouterDetail)

	* 新增成员：RedirectConfig

* [ModelRouterSet](https://cloud.tencent.com/document/api/214/30694#ModelRouterSet)

	* 新增成员：RedirectConfig




## 大数据智能体工作台DataBuddy(databuddy) 版本：2026-07-15

### 第 10 次发布

发布时间：2026-10-26 01:55:40

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateLocation](https://cloud.tencent.com/document/api/1835/140012)
* [CreateModel](https://cloud.tencent.com/document/api/1835/139987)
* [CreateModelVersion](https://cloud.tencent.com/document/api/1835/139986)
* [CreateVolume](https://cloud.tencent.com/document/api/1835/140011)
* [DeleteLocation](https://cloud.tencent.com/document/api/1835/140010)
* [DeleteModel](https://cloud.tencent.com/document/api/1835/139985)
* [DeleteView](https://cloud.tencent.com/document/api/1835/140009)
* [DeleteVolume](https://cloud.tencent.com/document/api/1835/140008)
* [GetLocation](https://cloud.tencent.com/document/api/1835/140007)
* [GetModel](https://cloud.tencent.com/document/api/1835/139984)
* [GetModelVersion](https://cloud.tencent.com/document/api/1835/139983)
* [GetView](https://cloud.tencent.com/document/api/1835/140006)
* [GetVolume](https://cloud.tencent.com/document/api/1835/140005)
* [ListCatalogNames](https://cloud.tencent.com/document/api/1835/139989)
* [ListFeatureTables](https://cloud.tencent.com/document/api/1835/139978)
* [ListLocations](https://cloud.tencent.com/document/api/1835/140004)
* [ListModels](https://cloud.tencent.com/document/api/1835/139982)
* [ListSchemaNames](https://cloud.tencent.com/document/api/1835/140003)
* [ListTableNames](https://cloud.tencent.com/document/api/1835/140002)
* [ListViewNames](https://cloud.tencent.com/document/api/1835/140001)
* [ListViews](https://cloud.tencent.com/document/api/1835/140000)
* [ListVolumes](https://cloud.tencent.com/document/api/1835/139999)
* [SearchAsset](https://cloud.tencent.com/document/api/1835/139998)
* [UpdateCatalog](https://cloud.tencent.com/document/api/1835/139997)
* [UpdateFunction](https://cloud.tencent.com/document/api/1835/139996)
* [UpdateLocation](https://cloud.tencent.com/document/api/1835/139995)
* [UpdateModel](https://cloud.tencent.com/document/api/1835/139981)
* [UpdateModelVersion](https://cloud.tencent.com/document/api/1835/139980)
* [UpdateSchema](https://cloud.tencent.com/document/api/1835/139994)
* [UpdateTable](https://cloud.tencent.com/document/api/1835/139993)
* [UpdateTableColumn](https://cloud.tencent.com/document/api/1835/139992)
* [UpdateView](https://cloud.tencent.com/document/api/1835/139991)
* [UpdateVolume](https://cloud.tencent.com/document/api/1835/139990)

新增数据结构：

* [AssetTagVote](https://cloud.tencent.com/document/api/1835/138006#AssetTagVote)
* [ColumnBrief](https://cloud.tencent.com/document/api/1835/138006#ColumnBrief)
* [ColumnErrorMessages](https://cloud.tencent.com/document/api/1835/138006#ColumnErrorMessages)
* [CreateModelRsp](https://cloud.tencent.com/document/api/1835/138006#CreateModelRsp)
* [CreateModelVersionRsp](https://cloud.tencent.com/document/api/1835/138006#CreateModelVersionRsp)
* [CreateVolumeRsp](https://cloud.tencent.com/document/api/1835/138006#CreateVolumeRsp)
* [DeleteLocationRsp](https://cloud.tencent.com/document/api/1835/138006#DeleteLocationRsp)
* [DeleteModelRsp](https://cloud.tencent.com/document/api/1835/138006#DeleteModelRsp)
* [DeleteViewRsp](https://cloud.tencent.com/document/api/1835/138006#DeleteViewRsp)
* [DeleteVolumeRsp](https://cloud.tencent.com/document/api/1835/138006#DeleteVolumeRsp)
* [FeatureTable](https://cloud.tencent.com/document/api/1835/138006#FeatureTable)
* [FieldInfo](https://cloud.tencent.com/document/api/1835/138006#FieldInfo)
* [FieldTagInfo](https://cloud.tencent.com/document/api/1835/138006#FieldTagInfo)
* [Fragment](https://cloud.tencent.com/document/api/1835/138006#Fragment)
* [FunctionInfo](https://cloud.tencent.com/document/api/1835/138006#FunctionInfo)
* [FunctionResource](https://cloud.tencent.com/document/api/1835/138006#FunctionResource)
* [GetLocationRsp](https://cloud.tencent.com/document/api/1835/138006#GetLocationRsp)
* [GetModelRsp](https://cloud.tencent.com/document/api/1835/138006#GetModelRsp)
* [GetModelVersionRsp](https://cloud.tencent.com/document/api/1835/138006#GetModelVersionRsp)
* [GetViewRsp](https://cloud.tencent.com/document/api/1835/138006#GetViewRsp)
* [GetVolumeRsp](https://cloud.tencent.com/document/api/1835/138006#GetVolumeRsp)
* [Highlighting](https://cloud.tencent.com/document/api/1835/138006#Highlighting)
* [HighlightingInfo](https://cloud.tencent.com/document/api/1835/138006#HighlightingInfo)
* [ListCatalogNamesRsp](https://cloud.tencent.com/document/api/1835/138006#ListCatalogNamesRsp)
* [ListFeatureTableResp](https://cloud.tencent.com/document/api/1835/138006#ListFeatureTableResp)
* [ListLocationsRsp](https://cloud.tencent.com/document/api/1835/138006#ListLocationsRsp)
* [ListModelsRsp](https://cloud.tencent.com/document/api/1835/138006#ListModelsRsp)
* [ListSchemaNamesRsp](https://cloud.tencent.com/document/api/1835/138006#ListSchemaNamesRsp)
* [ListTableNamesRsp](https://cloud.tencent.com/document/api/1835/138006#ListTableNamesRsp)
* [ListViewNamesRsp](https://cloud.tencent.com/document/api/1835/138006#ListViewNamesRsp)
* [ListViewsRsp](https://cloud.tencent.com/document/api/1835/138006#ListViewsRsp)
* [ListVolumesRsp](https://cloud.tencent.com/document/api/1835/138006#ListVolumesRsp)
* [LocationInfo](https://cloud.tencent.com/document/api/1835/138006#LocationInfo)
* [LocationUpdate](https://cloud.tencent.com/document/api/1835/138006#LocationUpdate)
* [ModelInfo](https://cloud.tencent.com/document/api/1835/138006#ModelInfo)
* [ModelVersionInfo](https://cloud.tencent.com/document/api/1835/138006#ModelVersionInfo)
* [NameIdentifier](https://cloud.tencent.com/document/api/1835/138006#NameIdentifier)
* [Owner](https://cloud.tencent.com/document/api/1835/138006#Owner)
* [SearchAssetRsp](https://cloud.tencent.com/document/api/1835/138006#SearchAssetRsp)
* [SearchResult](https://cloud.tencent.com/document/api/1835/138006#SearchResult)
* [Tag](https://cloud.tencent.com/document/api/1835/138006#Tag)
* [TagInfo](https://cloud.tencent.com/document/api/1835/138006#TagInfo)
* [UpdateCatalogRsp](https://cloud.tencent.com/document/api/1835/138006#UpdateCatalogRsp)
* [UpdateFunctionRsp](https://cloud.tencent.com/document/api/1835/138006#UpdateFunctionRsp)
* [UpdateLocationRsp](https://cloud.tencent.com/document/api/1835/138006#UpdateLocationRsp)
* [UpdateModelRsp](https://cloud.tencent.com/document/api/1835/138006#UpdateModelRsp)
* [UpdateModelVersionRsp](https://cloud.tencent.com/document/api/1835/138006#UpdateModelVersionRsp)
* [UpdateSchemaRsp](https://cloud.tencent.com/document/api/1835/138006#UpdateSchemaRsp)
* [UpdateTableColumnRsp](https://cloud.tencent.com/document/api/1835/138006#UpdateTableColumnRsp)
* [UpdateTableRsp](https://cloud.tencent.com/document/api/1835/138006#UpdateTableRsp)
* [UpdateViewRsp](https://cloud.tencent.com/document/api/1835/138006#UpdateViewRsp)
* [UpdateVolumeRsp](https://cloud.tencent.com/document/api/1835/138006#UpdateVolumeRsp)
* [UserActivityInfo](https://cloud.tencent.com/document/api/1835/138006#UserActivityInfo)
* [ViewInfo](https://cloud.tencent.com/document/api/1835/138006#ViewInfo)
* [Volume](https://cloud.tencent.com/document/api/1835/138006#Volume)

修改数据结构：

* [FeatureTableDetail](https://cloud.tencent.com/document/api/1835/138006#FeatureTableDetail)

	* 新增成员：OnlineTableType, ConnectionId

* [MetaOwner](https://cloud.tencent.com/document/api/1835/138006#MetaOwner)

	* 新增成员：CatalogSource




## 数据库智能管家 DBbrain(dbbrain) 版本：2021-05-27

### 第 68 次发布

发布时间：2026-10-26 01:57:20

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeBlockingLockTree](https://cloud.tencent.com/document/api/1130/140014)
* [DescribeBlockingTrend](https://cloud.tencent.com/document/api/1130/140013)
* [DescribeTableSpaceDataV2](https://cloud.tencent.com/document/api/1130/140015)

新增数据结构：

* [BlockingChainNode](https://cloud.tencent.com/document/api/1130/57812#BlockingChainNode)
* [BlockingChainTree](https://cloud.tencent.com/document/api/1130/57812#BlockingChainTree)
* [BlockingTrendPoint](https://cloud.tencent.com/document/api/1130/57812#BlockingTrendPoint)
* [MongodbSpaceObjectItem](https://cloud.tencent.com/document/api/1130/57812#MongodbSpaceObjectItem)
* [SqlServerSpaceObjectItem](https://cloud.tencent.com/document/api/1130/57812#SqlServerSpaceObjectItem)

修改数据结构：

* [MysqlSpaceObjectItem](https://cloud.tencent.com/document/api/1130/57812#MysqlSpaceObjectItem)

	* 新增成员：TableStatus

* [PostgresSpaceObjectItem](https://cloud.tencent.com/document/api/1130/57812#PostgresSpaceObjectItem)

	* 新增成员：TableStatus




## 数据库智能管家 DBbrain(dbbrain) 版本：2019-10-16



## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 326 次发布

发布时间：2026-10-26 02:13:12

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeBillUsage](https://cloud.tencent.com/document/api/1323/105296)

	* 新增入参：Operator




## 物联网开发平台(iotexplorer) 版本：2019-04-23

### 第 166 次发布

发布时间：2026-10-23 15:24:06

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreatePrivateCA](https://cloud.tencent.com/document/api/1081/139885)
* [DeletePrivateCA](https://cloud.tencent.com/document/api/1081/139884)
* [DeleteProductPrivateCA](https://cloud.tencent.com/document/api/1081/139883)
* [DescribePrivateCA](https://cloud.tencent.com/document/api/1081/139881)
* [DescribePrivateCABindedProducts](https://cloud.tencent.com/document/api/1081/139880)
* [DescribePrivateCAs](https://cloud.tencent.com/document/api/1081/139879)
* [DescribeProductCA](https://cloud.tencent.com/document/api/1081/139878)
* [UpdatePrivateCA](https://cloud.tencent.com/document/api/1081/139877)
* [UpdateProductPrivateCA](https://cloud.tencent.com/document/api/1081/139876)

修改接口：

* [CreateDevice](https://cloud.tencent.com/document/api/1081/37178)

	* 新增入参：TlsCrt

* [CreateStudioProduct](https://cloud.tencent.com/document/api/1081/34987)

	* 新增入参：PrivateCAName


新增数据结构：

* [CertInfo](https://cloud.tencent.com/document/api/1081/34988#CertInfo)
* [PrivateCABindProductInfo](https://cloud.tencent.com/document/api/1081/34988#PrivateCABindProductInfo)

修改数据结构：

* [ProductEntry](https://cloud.tencent.com/document/api/1081/34988#ProductEntry)

	* 新增成员：Uin, AppId, PrivateCAName, IsPrivateCA, CertName, CertSN, IssuerName, Subject, EffectiveTime, ExpireTime, CertText




## 前端性能监控(rum) 版本：2021-06-22

### 第 55 次发布

发布时间：2026-10-23 21:23:22

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateFieldTransformRule](https://cloud.tencent.com/document/api/1464/139963)
* [CreateTDEMApplication](https://cloud.tencent.com/document/api/1464/139636)
* [CreateTDEMBusinessSystem](https://cloud.tencent.com/document/api/1464/139635)
* [CreateTDEMDialTask](https://cloud.tencent.com/document/api/1464/139962)
* [CreateTDEMFieldTransformRule](https://cloud.tencent.com/document/api/1464/139961)
* [CreateTDEMFunnel](https://cloud.tencent.com/document/api/1464/139960)
* [DeleteFieldTransformRule](https://cloud.tencent.com/document/api/1464/139959)
* [DeleteMiniProgramReplayPack](https://cloud.tencent.com/document/api/1464/139958)
* [DeleteSymbolArtifact](https://cloud.tencent.com/document/api/1464/139957)
* [DeleteTDEMApplication](https://cloud.tencent.com/document/api/1464/139956)
* [DeleteTDEMBusinessSystem](https://cloud.tencent.com/document/api/1464/139955)
* [DeleteTDEMDialTask](https://cloud.tencent.com/document/api/1464/139954)
* [DeleteTDEMFieldTransformRule](https://cloud.tencent.com/document/api/1464/139953)
* [DeleteTDEMFunnel](https://cloud.tencent.com/document/api/1464/139952)
* [DescribeFieldTransformRules](https://cloud.tencent.com/document/api/1464/139951)
* [DescribeFieldValues](https://cloud.tencent.com/document/api/1464/139950)
* [DescribeIssueAffectedUsers](https://cloud.tencent.com/document/api/1464/139949)
* [DescribeIssueBreadcrumbs](https://cloud.tencent.com/document/api/1464/139948)
* [DescribeIssueCrashLogDownloadURL](https://cloud.tencent.com/document/api/1464/139947)
* [DescribeIssueDetail](https://cloud.tencent.com/document/api/1464/139946)
* [DescribeIssueEventDetail](https://cloud.tencent.com/document/api/1464/139945)
* [DescribeIssueEvents](https://cloud.tencent.com/document/api/1464/139944)
* [DescribeIssueTrend](https://cloud.tencent.com/document/api/1464/139943)
* [DescribeIssues](https://cloud.tencent.com/document/api/1464/139942)
* [DescribeMiniProgramReplayPack](https://cloud.tencent.com/document/api/1464/139941)
* [DescribeMiniProgramReplayPacks](https://cloud.tencent.com/document/api/1464/139940)
* [DescribePerformanceMetrics](https://cloud.tencent.com/document/api/1464/139939)
* [DescribeReplayData](https://cloud.tencent.com/document/api/1464/139938)
* [DescribeReplayDetails](https://cloud.tencent.com/document/api/1464/139937)
* [DescribeReplayMeta](https://cloud.tencent.com/document/api/1464/139936)
* [DescribeSampleRate](https://cloud.tencent.com/document/api/1464/139935)
* [DescribeSessionList](https://cloud.tencent.com/document/api/1464/139934)
* [DescribeStruggleDistribution](https://cloud.tencent.com/document/api/1464/139933)
* [DescribeSymbolArtifactUploadCredential](https://cloud.tencent.com/document/api/1464/139932)
* [DescribeSymbolArtifacts](https://cloud.tencent.com/document/api/1464/139931)
* [DescribeTDEMApplications](https://cloud.tencent.com/document/api/1464/139930)
* [DescribeTDEMBusinessSystems](https://cloud.tencent.com/document/api/1464/139634)
* [DescribeTDEMDialDimensions](https://cloud.tencent.com/document/api/1464/139929)
* [DescribeTDEMDialExecutionDetail](https://cloud.tencent.com/document/api/1464/139928)
* [DescribeTDEMDialExecutionList](https://cloud.tencent.com/document/api/1464/139927)
* [DescribeTDEMDialOverview](https://cloud.tencent.com/document/api/1464/139676)
* [DescribeTDEMDialTasks](https://cloud.tencent.com/document/api/1464/139675)
* [DescribeTDEMDialTrend](https://cloud.tencent.com/document/api/1464/139674)
* [DescribeTDEMEventDetail](https://cloud.tencent.com/document/api/1464/139673)
* [DescribeTDEMEventRank](https://cloud.tencent.com/document/api/1464/139672)
* [DescribeTDEMEventSearch](https://cloud.tencent.com/document/api/1464/139671)
* [DescribeTDEMEventSearchSummary](https://cloud.tencent.com/document/api/1464/139670)
* [DescribeTDEMFieldTransformRules](https://cloud.tencent.com/document/api/1464/139669)
* [DescribeTDEMFunnelDetail](https://cloud.tencent.com/document/api/1464/139668)
* [DescribeTDEMFunnelDimensions](https://cloud.tencent.com/document/api/1464/139667)
* [DescribeTDEMFunnelList](https://cloud.tencent.com/document/api/1464/139666)
* [DescribeTDEMFunnelReport](https://cloud.tencent.com/document/api/1464/139665)
* [DescribeTDEMFunnelStepSessionList](https://cloud.tencent.com/document/api/1464/139664)
* [DescribeTDEMFunnelTrend](https://cloud.tencent.com/document/api/1464/139663)
* [DescribeTDEMMetricDistribution](https://cloud.tencent.com/document/api/1464/139662)
* [DescribeTDEMMetricSummary](https://cloud.tencent.com/document/api/1464/139661)
* [DescribeTDEMMetricTrend](https://cloud.tencent.com/document/api/1464/139660)
* [DescribeTDEMOverviewScores](https://cloud.tencent.com/document/api/1464/139659)
* [DescribeTDEMRelatedEvents](https://cloud.tencent.com/document/api/1464/139658)
* [DescribeTDEMSSEDimensions](https://cloud.tencent.com/document/api/1464/139926)
* [DescribeTDEMSSEDistribution](https://cloud.tencent.com/document/api/1464/139925)
* [DescribeTDEMSSEPercentile](https://cloud.tencent.com/document/api/1464/139924)
* [DescribeTDEMSSESummary](https://cloud.tencent.com/document/api/1464/139923)
* [DescribeTDEMSSETrend](https://cloud.tencent.com/document/api/1464/139922)
* [DescribeTDEMSampleConfig](https://cloud.tencent.com/document/api/1464/139657)
* [DescribeTDEMSampleRate](https://cloud.tencent.com/document/api/1464/139656)
* [DescribeTDEMSessionReplayTimeline](https://cloud.tencent.com/document/api/1464/139655)
* [DescribeTDEMUsageDetail](https://cloud.tencent.com/document/api/1464/139654)
* [DescribeTDEMUsageSummary](https://cloud.tencent.com/document/api/1464/139653)
* [DescribeUserJourneySanKey](https://cloud.tencent.com/document/api/1464/139652)
* [ModifyFieldTransformRule](https://cloud.tencent.com/document/api/1464/139651)
* [ModifyIssueStatus](https://cloud.tencent.com/document/api/1464/139650)
* [ModifyMiniProgramReplayPack](https://cloud.tencent.com/document/api/1464/139649)
* [ModifySampleRate](https://cloud.tencent.com/document/api/1464/139648)
* [ModifyTDEMApplication](https://cloud.tencent.com/document/api/1464/139647)
* [ModifyTDEMBusinessSystem](https://cloud.tencent.com/document/api/1464/139646)
* [ModifyTDEMDialTask](https://cloud.tencent.com/document/api/1464/139645)
* [ModifyTDEMFieldTransformRule](https://cloud.tencent.com/document/api/1464/139644)
* [ModifyTDEMFunnel](https://cloud.tencent.com/document/api/1464/139643)
* [ModifyTDEMSampleConfig](https://cloud.tencent.com/document/api/1464/139642)
* [ModifyTDEMSampleRate](https://cloud.tencent.com/document/api/1464/139641)
* [PreviewTDEMFunnelReport](https://cloud.tencent.com/document/api/1464/139640)
* [ResumeTDEMApplication](https://cloud.tencent.com/document/api/1464/139639)
* [StopTDEMApplication](https://cloud.tencent.com/document/api/1464/139638)
* [UploadMiniProgramSource](https://cloud.tencent.com/document/api/1464/139637)
* [UploadSymbolArtifact](https://cloud.tencent.com/document/api/1464/139919)

新增数据结构：

* [AppOverviewScore](https://cloud.tencent.com/document/api/1464/61476#AppOverviewScore)
* [ApplicationItem](https://cloud.tencent.com/document/api/1464/61476#ApplicationItem)
* [BusinessSystemItem](https://cloud.tencent.com/document/api/1464/61476#BusinessSystemItem)
* [DialDimensionItem](https://cloud.tencent.com/document/api/1464/61476#DialDimensionItem)
* [DialExecution](https://cloud.tencent.com/document/api/1464/61476#DialExecution)
* [DialProbe](https://cloud.tencent.com/document/api/1464/61476#DialProbe)
* [DialTargetIpGeo](https://cloud.tencent.com/document/api/1464/61476#DialTargetIpGeo)
* [DialTask](https://cloud.tencent.com/document/api/1464/61476#DialTask)
* [DialTrendPoint](https://cloud.tencent.com/document/api/1464/61476#DialTrendPoint)
* [EventFrame](https://cloud.tencent.com/document/api/1464/61476#EventFrame)
* [EventRankItem](https://cloud.tencent.com/document/api/1464/61476#EventRankItem)
* [EventRankPoint](https://cloud.tencent.com/document/api/1464/61476#EventRankPoint)
* [EventSearchSeriesPoint](https://cloud.tencent.com/document/api/1464/61476#EventSearchSeriesPoint)
* [EventStack](https://cloud.tencent.com/document/api/1464/61476#EventStack)
* [EventView](https://cloud.tencent.com/document/api/1464/61476#EventView)
* [ExceptionBucket](https://cloud.tencent.com/document/api/1464/61476#ExceptionBucket)
* [ExceptionDimensionRow](https://cloud.tencent.com/document/api/1464/61476#ExceptionDimensionRow)
* [ExceptionDistribution](https://cloud.tencent.com/document/api/1464/61476#ExceptionDistribution)
* [ExplorationEventItem](https://cloud.tencent.com/document/api/1464/61476#ExplorationEventItem)
* [ExplorationExceptionStack](https://cloud.tencent.com/document/api/1464/61476#ExplorationExceptionStack)
* [FieldTransformRuleItem](https://cloud.tencent.com/document/api/1464/61476#FieldTransformRuleItem)
* [FloatKeyValue](https://cloud.tencent.com/document/api/1464/61476#FloatKeyValue)
* [FloatNameValue](https://cloud.tencent.com/document/api/1464/61476#FloatNameValue)
* [FunnelStepInput](https://cloud.tencent.com/document/api/1464/61476#FunnelStepInput)
* [IntKeyValue](https://cloud.tencent.com/document/api/1464/61476#IntKeyValue)
* [MetricDistributionBatchItem](https://cloud.tencent.com/document/api/1464/61476#MetricDistributionBatchItem)
* [MetricDistributionItem](https://cloud.tencent.com/document/api/1464/61476#MetricDistributionItem)
* [MetricQueryItem](https://cloud.tencent.com/document/api/1464/61476#MetricQueryItem)
* [MetricSummaryBatchItem](https://cloud.tencent.com/document/api/1464/61476#MetricSummaryBatchItem)
* [MetricTrendBatchItem](https://cloud.tencent.com/document/api/1464/61476#MetricTrendBatchItem)
* [MetricTrendPoint](https://cloud.tencent.com/document/api/1464/61476#MetricTrendPoint)
* [OverviewScoreItem](https://cloud.tencent.com/document/api/1464/61476#OverviewScoreItem)
* [StringKeyValue](https://cloud.tencent.com/document/api/1464/61476#StringKeyValue)
* [SymbolicatedView](https://cloud.tencent.com/document/api/1464/61476#SymbolicatedView)
* [TDEMSSEDimensionRow](https://cloud.tencent.com/document/api/1464/61476#TDEMSSEDimensionRow)
* [TDEMSSEDistributionBin](https://cloud.tencent.com/document/api/1464/61476#TDEMSSEDistributionBin)
* [TDEMSSEPercentilePoint](https://cloud.tencent.com/document/api/1464/61476#TDEMSSEPercentilePoint)
* [TDEMSSETrendPoint](https://cloud.tencent.com/document/api/1464/61476#TDEMSSETrendPoint)
* [TDEMSSETrendSeries](https://cloud.tencent.com/document/api/1464/61476#TDEMSSETrendSeries)
* [UsageDetailRow](https://cloud.tencent.com/document/api/1464/61476#UsageDetailRow)
* [UsageTimePoint](https://cloud.tencent.com/document/api/1464/61476#UsageTimePoint)



## 凭据管理系统(ssm) 版本：2019-09-23

### 第 20 次发布

发布时间：2026-10-26 03:01:17

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateService](https://cloud.tencent.com/document/api/1140/140016)



## 实时互动-工业能源版(trro) 版本：2022-03-25

### 第 15 次发布

发布时间：2026-10-26 03:21:52

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [Job](https://cloud.tencent.com/document/api/1584/89828#Job)

	* 新增成员：TaskStatusCount

* [ProcessParams](https://cloud.tencent.com/document/api/1584/89828#ProcessParams)

	* 新增成员：Edition, ExtraParams




## 云点播(vod) 版本：2024-07-18



## 云点播(vod) 版本：2018-07-17

### 第 292 次发布

发布时间：2026-10-26 03:27:11

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [CharacterSceneConfig](https://cloud.tencent.com/document/api/266/31773#CharacterSceneConfig)

修改数据结构：

* [AigcVideoSceneInfo](https://cloud.tencent.com/document/api/266/31773#AigcVideoSceneInfo)

	* 新增成员：CharacterSceneConfig

* [SceneAigcVideoOutputConfig](https://cloud.tencent.com/document/api/266/31773#SceneAigcVideoOutputConfig)

	* 新增成员：Resolution




## WorkBuddy Enterprise(workbuddyenterprise) 版本：2026-07-09

### 第 5 次发布

发布时间：2026-10-26 03:43:36

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [ConnectorInfo](https://cloud.tencent.com/document/api/1831/138580#ConnectorInfo)

	* 新增成员：DisplayName




