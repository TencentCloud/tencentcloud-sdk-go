# Release v1.3.69

## 批量计算(batch) 版本：2017-03-12

### 第 63 次发布

发布时间：2026-04-01 01:10:50

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [Placement](https://cloud.tencent.com/document/api/599/15912#Placement)

	* 新增成员：DedicatedResourcePackTenancy, DedicatedResourcePackIds




## 文件存储(cfs) 版本：2019-07-19

### 第 48 次发布

发布时间：2026-04-01 01:14:58

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateCfsFileSystem](https://cloud.tencent.com/document/api/582/38174)

	* 新增入参：Encrypted




## 消息队列 CKafka 版(ckafka) 版本：2019-08-19

### 第 142 次发布

发布时间：2026-04-01 01:15:59

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateRoute](https://cloud.tencent.com/document/api/597/70172)

	* 新增入参：IpWhitelist


新增数据结构：

* [IpWhitelistDTO](https://cloud.tencent.com/document/api/597/40861#IpWhitelistDTO)



## 日志服务(cls) 版本：2020-10-16

### 第 154 次发布

发布时间：2026-04-01 01:17:04

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateShipper](https://cloud.tencent.com/document/api/614/58747)

	* 新增入参：TimeZone, DSLFilter

* [ModifyShipper](https://cloud.tencent.com/document/api/614/58743)

	* 新增入参：TimeZone, DSLFilter

* [SearchLog](https://cloud.tencent.com/document/api/614/56447)

	* 新增入参：QueryString, QuerySyntax

	* <font color="#dd0000">**修改入参**：</font>Query


修改数据结构：

* [ExtractRuleInfo](https://cloud.tencent.com/document/api/614/56471#ExtractRuleInfo)

	* 新增成员：RawLogKey

* [ShipperInfo](https://cloud.tencent.com/document/api/614/56471#ShipperInfo)

	* 新增成员：TimeZone, DSLFilter




## 云服务器(cvm) 版本：2017-03-12

### 第 159 次发布

发布时间：2026-04-01 01:19:03

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [Placement](https://cloud.tencent.com/document/api/213/15753#Placement)

	* 新增成员：DedicatedResourcePackTenancy, DedicatedResourcePackIds




## Elasticsearch Service(es) 版本：2025-01-01



## Elasticsearch Service(es) 版本：2018-04-16

### 第 97 次发布

发布时间：2026-04-01 01:25:58

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [UpdateLogstashInstance](https://cloud.tencent.com/document/api/845/77233)

	* 新增入参：UserDnsIp


修改数据结构：

* [LogstashInstanceInfo](https://cloud.tencent.com/document/api/845/30634#LogstashInstanceInfo)

	* 新增成员：UserDnsIp




## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 286 次发布

发布时间：2026-04-01 01:26:23

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeDraftContractByPromptsTask](https://cloud.tencent.com/document/api/1323/129859)

	* 新增出参：ContractUrl




## 设备安全(tds) 版本：2022-08-01

### 第 7 次发布

发布时间：2026-04-01 02:17:49

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeFinanceFraudUltimate](https://cloud.tencent.com/document/api/1628/118247)

	* 新增出参：RiskCheckTimestamp, ExtraInfos

* [DescribeFraudUltimate](https://cloud.tencent.com/document/api/1628/83695)

	* 新增出参：RiskCheckTimestamp, ExtraInfos




## 容器服务(tke) 版本：2022-05-01



## 容器服务(tke) 版本：2018-05-25

### 第 223 次发布

发布时间：2026-04-01 02:24:43

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [GetMostSuitableImageCache](https://cloud.tencent.com/document/api/457/70858)

	* 新增入参：Snapshotter


修改数据结构：

* [ImageCache](https://cloud.tencent.com/document/api/457/31866#ImageCache)

	* 新增成员：ImageCacheType, Snapshotter




## 私有网络(vpc) 版本：2017-03-12

### 第 299 次发布

发布时间：2026-04-01 02:41:46

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [TrafficMirror](https://cloud.tencent.com/document/api/215/15824#TrafficMirror)

	* 新增成员：IngressFilterRules, EgressFilterRules

* [TrafficMirrorFilter](https://cloud.tencent.com/document/api/215/15824#TrafficMirrorFilter)

	* 新增成员：TrafficMirrorFilterRuleId, Priority, Action, Description, CreatedTime

	* <font color="#dd0000">**修改成员**：</font>SrcNet, DstNet, Protocol




