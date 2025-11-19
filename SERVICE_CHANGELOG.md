# Release v1.2.3

## 消息队列 CKafka 版(ckafka) 版本：2019-08-19

### 第 136 次发布

发布时间：2025-11-20 01:14:29

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除接口**：</font>

* DescribeAppInfo

<font color="#dd0000">**删除数据结构**：</font>

* AppIdResponse



## 数据湖计算 DLC(dlc) 版本：2021-01-25

### 第 141 次发布

发布时间：2025-11-20 01:18:53

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [AttachUserPolicy](https://cloud.tencent.com/document/api/1342/58468)

	* 新增入参：AccountType

* [CreateUser](https://cloud.tencent.com/document/api/1342/58465)

	* 新增入参：AccountType

* [DeleteUser](https://cloud.tencent.com/document/api/1342/58463)

	* 新增入参：AccountType

* [DescribeUserInfo](https://cloud.tencent.com/document/api/1342/99262)

	* 新增入参：AccountType

* [DescribeUserType](https://cloud.tencent.com/document/api/1342/99261)

	* 新增入参：AccountType

* [DescribeUsers](https://cloud.tencent.com/document/api/1342/58460)

	* 新增入参：AccountType

* [DetachUserPolicy](https://cloud.tencent.com/document/api/1342/58458)

	* 新增入参：AccountType

* [ModifyUser](https://cloud.tencent.com/document/api/1342/58456)

	* 新增入参：AccountType

* [ModifyUserType](https://cloud.tencent.com/document/api/1342/99259)

	* 新增入参：AccountType


修改数据结构：

* [Policy](https://cloud.tencent.com/document/api/1342/53778#Policy)

	* 新增成员：Model

* [UserDetailInfo](https://cloud.tencent.com/document/api/1342/53778#UserDetailInfo)

	* 新增成员：ModelPolicyInfo

* [UserMessage](https://cloud.tencent.com/document/api/1342/53778#UserMessage)

	* 新增成员：AccountType




## 数据传输服务(dts) 版本：2021-12-06

### 第 48 次发布

发布时间：2025-11-20 01:20:59

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [KafkaOption](https://cloud.tencent.com/document/api/571/82108#KafkaOption)

	* 新增成员：DataOption




## 数据传输服务(dts) 版本：2018-03-30



## 物联网开发平台(iotexplorer) 版本：2019-04-23

### 第 124 次发布

发布时间：2025-11-19 17:22:41

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [BatchCreateTWeSeeRecognitionTask](https://cloud.tencent.com/document/api/1081/125438)
* [BatchInvokeTWeSeeRecognitionTask](https://cloud.tencent.com/document/api/1081/125437)
* [CreateTWeSeeRecognitionTaskWithFile](https://cloud.tencent.com/document/api/1081/125436)
* [DeleteTWeTalkProductConfigV2](https://cloud.tencent.com/document/api/1081/125433)
* [InvokeTWeSeeRecognitionTaskWithFile](https://cloud.tencent.com/document/api/1081/125435)

<font color="#dd0000">**删除接口**：</font>

* DescribeDeviceLocationSolve

新增数据结构：

* [CreateVisionRecognitionTaskInput](https://cloud.tencent.com/document/api/1081/34988#CreateVisionRecognitionTaskInput)
* [CreateVisionRecognitionTaskOutput](https://cloud.tencent.com/document/api/1081/34988#CreateVisionRecognitionTaskOutput)
* [InvokeVisionRecognitionTaskOutput](https://cloud.tencent.com/document/api/1081/34988#InvokeVisionRecognitionTaskOutput)
* [VisionCustomDetectQuery](https://cloud.tencent.com/document/api/1081/34988#VisionCustomDetectQuery)
* [VisionDetectedObject](https://cloud.tencent.com/document/api/1081/34988#VisionDetectedObject)

<font color="#dd0000">**删除数据结构**：</font>

* WifiInfo

修改数据结构：

* [VisionRecognitionResult](https://cloud.tencent.com/document/api/1081/34988#VisionRecognitionResult)

	* 新增成员：DetectedObjects

* [VisionSummaryConfig](https://cloud.tencent.com/document/api/1081/34988#VisionSummaryConfig)

	* 新增成员：DetectTypes, CustomDetectQueries




## 媒体处理(mps) 版本：2019-06-12

### 第 157 次发布

发布时间：2025-11-20 01:37:36

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeSmartEraseTemplates](https://cloud.tencent.com/document/api/862/123733)

	* 新增入参：EraseType


新增数据结构：

* [LiveAiAnalysisDescriptionItem](https://cloud.tencent.com/document/api/862/37615#LiveAiAnalysisDescriptionItem)
* [LiveAiParagraphInfo](https://cloud.tencent.com/document/api/862/37615#LiveAiParagraphInfo)

修改数据结构：

* [LiveStreamAiAnalysisResultInfo](https://cloud.tencent.com/document/api/862/37615#LiveStreamAiAnalysisResultInfo)

	* <font color="#dd0000">**修改成员**：</font>ResultSet

* [LiveStreamAiAnalysisResultItem](https://cloud.tencent.com/document/api/862/37615#LiveStreamAiAnalysisResultItem)

	* 新增成员：DescriptionResult




## 音视频终端引擎(vcube) 版本：2022-04-10

### 第 4 次发布

发布时间：2025-11-20 02:27:28

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateApplicationAndVideo](https://cloud.tencent.com/document/api/1449/113532)

	* 新增出参：LicenseId

* [CreateApplicationAndWebPlayerLicense](https://cloud.tencent.com/document/api/1449/113536)

	* 新增出参：LicenseId

* [DescribeVcubeApplicationAndPlayList](https://cloud.tencent.com/document/api/1449/113527)

	* 新增入参：LicenseId




## 私有网络(vpc) 版本：2017-03-12

### 第 277 次发布

发布时间：2025-11-20 02:31:45

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DeletePrivateNatGateway](https://cloud.tencent.com/document/api/215/107900)

	* 新增入参：IgnoreOperationRisk




