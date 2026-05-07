# Release v1.3.93

## 应用性能监控(apm) 版本：2021-06-22

### 第 60 次发布

发布时间：2026-05-08 01:11:02

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [ModifyApmService](https://cloud.tencent.com/document/api/1463/131311)



## 运维安全中心（堡垒机）(bh) 版本：2023-04-18

### 第 31 次发布

发布时间：2026-05-08 01:12:58

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [BindDeviceAccountKubeconfig](https://cloud.tencent.com/document/api/1025/131312)



## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 292 次发布

发布时间：2026-05-08 01:53:21

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateMultiFlowSignQRCode](https://cloud.tencent.com/document/api/1323/75450)

	* 新增入参：FlowDisplayType


修改数据结构：

* [RecipientComponentInfo](https://cloud.tencent.com/document/api/1323/70369#RecipientComponentInfo)

	* 新增成员：SignComponents




## 腾讯电子签（基础版）(essbasic) 版本：2021-05-26

### 第 255 次发布

发布时间：2026-05-08 01:55:06

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ChannelCreateMultiFlowSignQRCode](https://cloud.tencent.com/document/api/1420/75452)

	* 新增入参：FlowDisplayType




## 腾讯电子签（基础版）(essbasic) 版本：2020-12-22



## iOA 零信任安全管理系统(ioa) 版本：2022-06-01

### 第 35 次发布

发布时间：2026-05-08 02:01:56

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [ModifyBusinessResource](https://cloud.tencent.com/document/api/1092/131313)



## 云直播CSS(live) 版本：2018-08-01

### 第 175 次发布

发布时间：2026-05-08 02:12:30

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeOriginStreamInfo](https://cloud.tencent.com/document/api/267/130621)

	* 新增入参：DomainName

	* 新增出参：Status, CdnStreamPlayType, OriginStreamType, OriginStreamPlayType, OriginAddressType, OriginAddress, OriginTimeout, OriginRetryTimes, TimeJitter, HlsPlayFragmentCount, HlsPlayFragmentDuration, PassThroughHttpHeader, PassThroughResponse, PassThroughParam, OriginHost, IndexerCache, FragmentCache, DomainName, UsingHttps, CacheFollowOrigin, CacheStatusCode, UrlReplaceRules, OptionsRequest, FollowRedirect, StreamPackageRegion, CustomerName, IndexerKeepParam, FragmentKeepParam, MediaPackageType, MediaPackageChannelTypes, IndexerHeader, FragmentHeader, CustomizationRules

* [ModifyOriginStreamInfo](https://cloud.tencent.com/document/api/267/130620)

	* 新增入参：DomainName, OriginStreamPlayType, CdnStreamPlayType, OriginStreamType, OriginAddress, OriginAddressType, CustomerName, OriginHost, OriginTimeout, OriginRetryTimes, PassThroughHttpHeader, PassThroughResponse, PassThroughParam, IndexerCache, FragmentCache, HlsPlayFragmentCount, HlsPlayFragmentDuration, TimeJitter, UsingHttps, CacheFollowOrigin, CacheStatusCode, UrlReplaceRules, OptionsRequest, FollowRedirect, IndexerKeepParam, FragmentKeepParam, MediaPackageType, MediaPackageChannelTypes, IndexerHeader, FragmentHeader, CustomizationRules, CacheFormatRule


新增数据结构：

* [OriginStreamCustomizationRule](https://cloud.tencent.com/document/api/267/20474#OriginStreamCustomizationRule)



## 云数据库 MongoDB(mongodb) 版本：2019-07-25

### 第 69 次发布

发布时间：2026-05-08 02:21:29

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [EnableWanService](https://cloud.tencent.com/document/api/240/131314)

新增数据结构：

* [WanServiceNodeList](https://cloud.tencent.com/document/api/240/38576#WanServiceNodeList)



## 云数据库 MongoDB(mongodb) 版本：2018-04-08



## 媒体处理(mps) 版本：2019-06-12

### 第 200 次发布

发布时间：2026-05-08 02:26:22

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateAigcImageTask](https://cloud.tencent.com/document/api/862/126966)

	* 新增入参：SceneType




## 短信(sms) 版本：2021-01-11

### 第 8 次发布

发布时间：2026-05-08 02:50:59

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeSmsSignList](https://cloud.tencent.com/document/api/382/55969)

	* 新增入参：Limit, Offset

	* <font color="#dd0000">**修改入参**：</font>SignIdSet




## 短信(sms) 版本：2019-07-11



## 云开发 CloudBase(tcb) 版本：2018-06-08

### 第 137 次发布

发布时间：2026-05-08 03:00:24

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateAIModel](https://cloud.tencent.com/document/api/876/131320)
* [DeleteAIModel](https://cloud.tencent.com/document/api/876/131319)
* [DescribeAIModels](https://cloud.tencent.com/document/api/876/131318)
* [DescribeManagedAIModelList](https://cloud.tencent.com/document/api/876/131317)
* [UpdateAIModel](https://cloud.tencent.com/document/api/876/131316)

新增数据结构：

* [AIModel](https://cloud.tencent.com/document/api/876/34822#AIModel)
* [AIModelGroup](https://cloud.tencent.com/document/api/876/34822#AIModelGroup)
* [AIModelSecret](https://cloud.tencent.com/document/api/876/34822#AIModelSecret)
* [ManagedAIModel](https://cloud.tencent.com/document/api/876/34822#ManagedAIModel)
* [ManagedAIModelChargingInfo](https://cloud.tencent.com/document/api/876/34822#ManagedAIModelChargingInfo)
* [ManagedAIModelGroup](https://cloud.tencent.com/document/api/876/34822#ManagedAIModelGroup)
* [ManagedAIModelSpec](https://cloud.tencent.com/document/api/876/34822#ManagedAIModelSpec)



## 云点播(vod) 版本：2024-07-18



## 云点播(vod) 版本：2018-07-17

### 第 253 次发布

发布时间：2026-05-08 03:44:09

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeTaskDetail](https://cloud.tencent.com/document/api/266/33431)

	* 新增出参：AigcAudioTask


新增数据结构：

* [AigcAudioOutputConfig](https://cloud.tencent.com/document/api/266/31773#AigcAudioOutputConfig)
* [AigcAudioTask](https://cloud.tencent.com/document/api/266/31773#AigcAudioTask)
* [AigcAudioTaskInput](https://cloud.tencent.com/document/api/266/31773#AigcAudioTaskInput)
* [AigcAudioTaskOutput](https://cloud.tencent.com/document/api/266/31773#AigcAudioTaskOutput)
* [AigcAudioTaskOutputFileInfo](https://cloud.tencent.com/document/api/266/31773#AigcAudioTaskOutputFileInfo)



