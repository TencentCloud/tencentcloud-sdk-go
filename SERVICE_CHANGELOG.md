# Release v1.0.1177

## TDSQL-C MySQL 版(cynosdb) 版本：2019-01-07

### 第 136 次发布

发布时间：2025-05-30 01:25:57

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeBackupDownloadRestriction](https://cloud.tencent.com/document/api/1003/118922)
* [DescribeBackupDownloadUserRestriction](https://cloud.tencent.com/document/api/1003/118921)
* [DescribeClusterReadOnly](https://cloud.tencent.com/document/api/1003/118924)
* [ModifyBackupDownloadRestriction](https://cloud.tencent.com/document/api/1003/118920)
* [ModifyBackupDownloadUserRestriction](https://cloud.tencent.com/document/api/1003/118919)
* [ModifyClusterReadOnly](https://cloud.tencent.com/document/api/1003/118923)

修改接口：

* [DescribeBackupDownloadUrl](https://cloud.tencent.com/document/api/1003/76374)

	* 新增入参：DownloadRestriction

* [DescribeBinlogDownloadUrl](https://cloud.tencent.com/document/api/1003/76373)

	* 新增入参：DownloadRestriction


新增数据结构：

* [BackupLimitClusterRestriction](https://cloud.tencent.com/document/api/1003/48097#BackupLimitClusterRestriction)
* [BackupLimitRestriction](https://cloud.tencent.com/document/api/1003/48097#BackupLimitRestriction)
* [BackupLimitVpcItem](https://cloud.tencent.com/document/api/1003/48097#BackupLimitVpcItem)
* [ClusterReadOnlyValue](https://cloud.tencent.com/document/api/1003/48097#ClusterReadOnlyValue)
* [ClusterTaskId](https://cloud.tencent.com/document/api/1003/48097#ClusterTaskId)



## 弹性 MapReduce(emr) 版本：2019-01-03

### 第 104 次发布

发布时间：2025-05-30 01:33:30

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**预下线接口**：</font>

* ModifyResourceScheduleConfig
* ModifyYarnDeploy



## 文字识别(ocr) 版本：2018-11-19

### 第 191 次发布

发布时间：2025-05-30 01:51:27

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [RecognizeGeneralInvoice](https://cloud.tencent.com/document/api/866/90802)

	* 新增入参：EnableQRCode


新增数据结构：

* [CustomsDeclaration](https://cloud.tencent.com/document/api/866/33527#CustomsDeclaration)
* [OnlineTaxiItinerary](https://cloud.tencent.com/document/api/866/33527#OnlineTaxiItinerary)
* [OverseasInvoice](https://cloud.tencent.com/document/api/866/33527#OverseasInvoice)
* [SaleInventory](https://cloud.tencent.com/document/api/866/33527#SaleInventory)
* [ShoppingReceipt](https://cloud.tencent.com/document/api/866/33527#ShoppingReceipt)

修改数据结构：

* [InvoiceItem](https://cloud.tencent.com/document/api/866/33527#InvoiceItem)

	* 新增成员：QRCode

* [SingleInvoiceItem](https://cloud.tencent.com/document/api/866/33527#SingleInvoiceItem)

	* 新增成员：OnlineTaxiItinerary, CustomsDeclaration, OverseasInvoice, ShoppingReceipt, SaleInventory




## 邮件推送(ses) 版本：2020-10-02

### 第 31 次发布

发布时间：2025-05-30 01:56:10

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateAddressUnsubscribeConfig](https://cloud.tencent.com/document/api/1288/118927)
* [DeleteAddressUnsubscribeConfig](https://cloud.tencent.com/document/api/1288/118926)
* [UpdateAddressUnsubscribeConfig](https://cloud.tencent.com/document/api/1288/118925)



