# Release v1.3.132

## TDSQL-C MySQL 版(cynosdb) 版本：2019-01-07

### 第 181 次发布

发布时间：2026-07-13 01:44:36

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeClusterLevels](https://cloud.tencent.com/document/api/1003/134025)



## 数据湖计算 DLC(dlc) 版本：2021-01-25

### 第 163 次发布

发布时间：2026-07-13 01:49:57

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [AttachUserPolicy](https://cloud.tencent.com/document/api/1342/58468)

	* 新增出参：PolicySet

* [AttachWorkGroupPolicy](https://cloud.tencent.com/document/api/1342/58467)

	* 新增出参：PolicySet

* [DescribeUserInfo](https://cloud.tencent.com/document/api/1342/99262)

	* 新增入参：PolicyId

* [DescribeWorkGroupInfo](https://cloud.tencent.com/document/api/1342/99260)

	* 新增入参：PolicyId

* [DetachUserPolicy](https://cloud.tencent.com/document/api/1342/58458)

	* 新增入参：PolicyIds

* [DetachWorkGroupPolicy](https://cloud.tencent.com/document/api/1342/58457)

	* 新增入参：PolicyIds


修改数据结构：

* [Policy](https://cloud.tencent.com/document/api/1342/53778#Policy)

	* 新增成员：PolicyId




## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 306 次发布

发布时间：2026-07-13 02:00:21

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateContractComparisonTask](https://cloud.tencent.com/document/api/1323/124366)

	* 新增入参：RevisionOperation




## 人脸核身(faceid) 版本：2018-03-01

### 第 90 次发布

发布时间：2026-07-13 02:03:27

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [IdCardOCRVerification](https://cloud.tencent.com/document/api/1007/37980)

	* 新增入参：Config

	* 新增出参：Portrait, Warnings, Quality




## 高性能应用服务(hai) 版本：2023-08-12

### 第 27 次发布

发布时间：2026-07-13 02:07:23

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateInferServiceByTemplate](https://cloud.tencent.com/document/api/1721/129388)

	* 新增入参：ServiceChargePrepaid

* [DeployInferService](https://cloud.tencent.com/document/api/1721/129387)

	* 新增入参：ServiceChargePrepaid


新增数据结构：

* [ServiceChargePrepaid](https://cloud.tencent.com/document/api/1721/101518#ServiceChargePrepaid)

修改数据结构：

* [ComputeDetail](https://cloud.tencent.com/document/api/1721/101518#ComputeDetail)

	* 新增成员：PrepaidEnable, PostpaidEnable

* [ServiceDetail](https://cloud.tencent.com/document/api/1721/101518#ServiceDetail)

	* 新增成员：ChargeType, ExpireTime, RenewFlag, RestrictState, IsCustomDeploy




## 媒体处理(mps) 版本：2019-06-12

### 第 215 次发布

发布时间：2026-07-13 02:26:32

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateAigcVideoTask](https://cloud.tencent.com/document/api/862/126965)

	* 新增入参：AudioInfos


新增数据结构：

* [AigcVideoReferenceAudioInfo](https://cloud.tencent.com/document/api/862/37615#AigcVideoReferenceAudioInfo)



## TDSQL(tdmysql) 版本：2021-11-22

### 第 12 次发布

发布时间：2026-07-13 02:52:24

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeFlow](https://cloud.tencent.com/document/api/1376/128281)

	* 新增出参：Status




## 文本内容安全(tms) 版本：2020-12-29

### 第 17 次发布

发布时间：2026-07-13 03:00:40

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [TextModeration](https://cloud.tencent.com/document/api/1124/51860)

	* 新增出参：HitSnippetInfos


新增数据结构：

* [HitPosition](https://cloud.tencent.com/document/api/1124/51861#HitPosition)
* [HitSnippetInfo](https://cloud.tencent.com/document/api/1124/51861#HitSnippetInfo)



## 文本内容安全(tms) 版本：2020-07-13



## TokenHub(tokenhub) 版本：2026-03-22

### 第 13 次发布

发布时间：2026-07-13 03:00:50

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateApiKey](https://cloud.tencent.com/document/api/1823/133810)

	* 新增入参：ApiKeyName, Platform, BindType, Remark, Status, Bindings, IpWhitelist, Quotas

	* 新增出参：ApiKeyId

* [DeleteApiKey](https://cloud.tencent.com/document/api/1823/133809)

	* 新增入参：ApiKeyId, Platform

* [ModifyApiKeyInfo](https://cloud.tencent.com/document/api/1823/133808)

	* 新增入参：ApiKeyId, Platform, ApiKeyName, Remark, IpWhitelist, QuotasDesired

* [ModifyApiKeyStatus](https://cloud.tencent.com/document/api/1823/133807)

	* 新增入参：ApiKeyId, Platform, Status


新增数据结构：

* [QuotaCreateItem](https://cloud.tencent.com/document/api/1823/132279#QuotaCreateItem)
* [QuotasDesired](https://cloud.tencent.com/document/api/1823/132279#QuotasDesired)



## TSF-Polaris&ZK&网关(tse) 版本：2020-12-07

### 第 110 次发布

发布时间：2026-07-13 03:03:59

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [BetaLabel](https://cloud.tencent.com/document/api/1364/54942#BetaLabel)
* [BetaLabelMatchString](https://cloud.tencent.com/document/api/1364/54942#BetaLabelMatchString)

修改数据结构：

* [ConfigFileRelease](https://cloud.tencent.com/document/api/1364/54942#ConfigFileRelease)

	* 新增成员：BetaLabels, ReleaseType




## 云点播(vod) 版本：2024-07-18



## 云点播(vod) 版本：2018-07-17

### 第 270 次发布

发布时间：2026-07-13 03:08:57

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [BeautyEffectItem](https://cloud.tencent.com/document/api/266/31773#BeautyEffectItem)
* [BeautyFilterItem](https://cloud.tencent.com/document/api/266/31773#BeautyFilterItem)
* [ImageBeautyConfig](https://cloud.tencent.com/document/api/266/31773#ImageBeautyConfig)

修改数据结构：

* [ProcessImageAsyncInput](https://cloud.tencent.com/document/api/266/31773#ProcessImageAsyncInput)

	* 新增成员：Url

* [ProcessImageAsyncTask](https://cloud.tencent.com/document/api/266/31773#ProcessImageAsyncTask)

	* 新增成员：BeautyConfig




