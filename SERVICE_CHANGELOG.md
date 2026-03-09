# Release v1.3.53

## 暴露面管理服务(ctem) 版本：2023-11-28

### 第 15 次发布

发布时间：2026-03-10 01:16:32

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeApiSecs](https://cloud.tencent.com/document/api/1755/125480)

	* 新增入参：OrderBy




## 主机安全(cwp) 版本：2018-02-28

### 第 158 次发布

发布时间：2026-03-10 01:17:22

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateWhiteListOrder](https://cloud.tencent.com/document/api/296/99541)

	* <font color="#dd0000">**修改入参**：</font>SourceType

* [DescribeLogStorageConfig](https://cloud.tencent.com/document/api/296/91086)

	* 新增出参：MsgLanguage

* [ModifyLogStorageConfig](https://cloud.tencent.com/document/api/296/91084)

	* 新增入参：MsgLanguage




## 容器安全服务(tcss) 版本：2020-11-01

### 第 88 次发布

发布时间：2026-03-10 02:08:56

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeSecLogJoinTypeList](https://cloud.tencent.com/document/api/1285/81727)

	* 新增出参：MsgLanguage


修改数据结构：

* [AbnormalProcessChildRuleInfo](https://cloud.tencent.com/document/api/1285/65614#AbnormalProcessChildRuleInfo)

	* 新增成员：CmdLine

* [AbnormalProcessEventInfo](https://cloud.tencent.com/document/api/1285/65614#AbnormalProcessEventInfo)

	* 新增成员：CmdLine

* [AccessControlChildRuleInfo](https://cloud.tencent.com/document/api/1285/65614#AccessControlChildRuleInfo)

	* 新增成员：CmdLine

* [AccessControlEventInfo](https://cloud.tencent.com/document/api/1285/65614#AccessControlEventInfo)

	* 新增成员：CmdLine




## 消息队列 TDMQ(tdmq) 版本：2020-02-17

### 第 167 次发布

发布时间：2026-03-10 02:14:03

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除接口**：</font>

* CreateCluster



## TI-ONE 训练平台(tione) 版本：2021-11-11

### 第 105 次发布

发布时间：2026-03-10 02:21:54

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateNotebook](https://cloud.tencent.com/document/api/851/95658)

	* 新增入参：Description

* [DescribeTrainingTask](https://cloud.tencent.com/document/api/851/75089)

	* 新增入参：InstanceId

* [ModifyNotebook](https://cloud.tencent.com/document/api/851/127303)

	* 新增入参：Description


修改数据结构：

* [AuthTokenBase](https://cloud.tencent.com/document/api/851/75051#AuthTokenBase)

	* 新增成员：Id

* [NotebookDetail](https://cloud.tencent.com/document/api/851/75051#NotebookDetail)

	* 新增成员：Description

* [NotebookSetItem](https://cloud.tencent.com/document/api/851/75051#NotebookSetItem)

	* 新增成员：Description




## TI-ONE 训练平台(tione) 版本：2019-10-22



## 容器服务(tke) 版本：2022-05-01

### 第 23 次发布

发布时间：2026-03-10 02:26:29

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyClusterMachine](https://cloud.tencent.com/document/api/457/126106)

	* 新增入参：SecurityGroupIDs


修改数据结构：

* [UpdateNativeNodePoolParam](https://cloud.tencent.com/document/api/457/103206#UpdateNativeNodePoolParam)

	* 新增成员：UpdateMachineManagement




## 容器服务(tke) 版本：2018-05-25



## 云点播(vod) 版本：2024-07-18



## 云点播(vod) 版本：2018-07-17

### 第 235 次发布

发布时间：2026-03-10 02:37:38

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateLLMComprehendTemplate](https://cloud.tencent.com/document/api/266/128363)

	* 新增入参：FaceRecognition

* [ModifyLLMComprehendTemplate](https://cloud.tencent.com/document/api/266/128360)

	* 新增入参：FaceRecognition

* [SearchMediaBySemantics](https://cloud.tencent.com/document/api/266/126287)

	* 新增入参：Persons


新增数据结构：

* [LLMComprehendFaceRecognition](https://cloud.tencent.com/document/api/266/31773#LLMComprehendFaceRecognition)
* [LLMComprehendFaceRecognitionForUpdate](https://cloud.tencent.com/document/api/266/31773#LLMComprehendFaceRecognitionForUpdate)



## Web 应用防火墙(waf) 版本：2018-01-25

### 第 143 次发布

发布时间：2026-03-10 02:46:34

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeApiSecSensitiveRuleList](https://cloud.tencent.com/document/api/627/129030)

新增数据结构：

* [ApiSecSensitiveRule](https://cloud.tencent.com/document/api/627/53609#ApiSecSensitiveRule)



## 企业微信汽车行业版(wav) 版本：2021-01-29

### 第 19 次发布

发布时间：2026-03-10 02:49:06

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**预下线接口**：</font>

* CreateChannelCode
* CreateLead
* QueryActivityJoinList
* QueryActivityList
* QueryActivityLiveCodeList
* QueryArrivalList
* QueryChannelCodeList
* QueryChatArchivingList
* QueryClueInfoList
* QueryCrmStatistics
* QueryCustomerEventDetailStatistics
* QueryCustomerProfileList
* QueryDealerInfoList
* QueryExternalContactDetail
* QueryExternalContactDetailByDate
* QueryExternalContactList
* QueryExternalUserEventList
* QueryExternalUserMappingInfo
* QueryFollowList
* QueryLicenseInfo
* QueryMaterialList
* QueryMiniAppCodeList
* QueryStaffEventDetailStatistics
* QueryUserInfoList
* QueryVehicleInfoList



