# Release v1.3.203

## 腾讯云智能体开发平台(adp) 版本：2026-05-20

### 第 33 次发布

发布时间：2026-10-29 01:04:08

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [Space](https://cloud.tencent.com/document/api/1759/132545#Space)

	* 新增成员：CreateTime, UpdateTime




## 语音识别(asr) 版本：2019-06-14

### 第 50 次发布

发布时间：2026-10-29 01:09:36

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [SpeakerRoleInfo](https://cloud.tencent.com/document/api/1093/37824#SpeakerRoleInfo)

	* 新增成员：VoicePrintId, SpeakerNick

	* <font color="#dd0000">**修改成员**：</font>RoleAudioUrl, RoleName




## 智能顾问-混沌演练(cfg) 版本：2021-08-20

### 第 36 次发布

发布时间：2026-10-29 01:19:26

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeInstanceAgentList](https://cloud.tencent.com/document/api/1500/140308)
* [InstallAgentOpenApi](https://cloud.tencent.com/document/api/1500/140307)
* [UninstallAgentOpenApi](https://cloud.tencent.com/document/api/1500/140306)

新增数据结构：

* [AgentFilter](https://cloud.tencent.com/document/api/1500/71784#AgentFilter)
* [InstanceAgent](https://cloud.tencent.com/document/api/1500/71784#InstanceAgent)
* [InstanceAgentItem](https://cloud.tencent.com/document/api/1500/71784#InstanceAgentItem)



## 暴露面管理服务(ctem) 版本：2023-11-28

### 第 23 次发布

发布时间：2026-10-29 01:34:28

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeIntelligences](https://cloud.tencent.com/document/api/1755/140184)

	* 新增入参：CustomerId, IsActual

	* 新增出参：Total, List


新增数据结构：

* [IntelligenceDesc](https://cloud.tencent.com/document/api/1755/120320#IntelligenceDesc)



## 大数据智能体工作台DataBuddy(databuddy) 版本：2026-07-15

### 第 11 次发布

发布时间：2026-10-29 01:42:35

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [RemoveWorkspaceUsers](https://cloud.tencent.com/document/api/1835/140311)
* [UpdateWorkspaceGroupsTags](https://cloud.tencent.com/document/api/1835/140310)
* [UpdateWorkspaceUsersTags](https://cloud.tencent.com/document/api/1835/140309)

新增数据结构：

* [RemoveWorkspaceUsersRsp](https://cloud.tencent.com/document/api/1835/138006#RemoveWorkspaceUsersRsp)
* [UpdateWorkspaceGroupsTagsRsp](https://cloud.tencent.com/document/api/1835/138006#UpdateWorkspaceGroupsTagsRsp)
* [UpdateWorkspaceUsersTagsRsp](https://cloud.tencent.com/document/api/1835/138006#UpdateWorkspaceUsersTagsRsp)



## 云数据库独享集群(dbdc) 版本：2020-10-29

### 第 18 次发布

发布时间：2026-10-29 01:44:54

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [ModifyDBCustomNodeType](https://cloud.tencent.com/document/api/1322/140317)
* [RebootDBCustomNodes](https://cloud.tencent.com/document/api/1322/140316)
* [ResetDBCustomNode](https://cloud.tencent.com/document/api/1322/140315)
* [ResizeDBCustomNodeDisk](https://cloud.tencent.com/document/api/1322/140314)
* [StartDBCustomNodes](https://cloud.tencent.com/document/api/1322/140313)
* [StopDBCustomNodes](https://cloud.tencent.com/document/api/1322/140312)



## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 327 次发布

发布时间：2026-10-29 01:56:54

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeContractDiffTaskWebUrl](https://cloud.tencent.com/document/api/1323/119176)

	* 新增入参：FormatConfigs


新增数据结构：

* [FormatConfigs](https://cloud.tencent.com/document/api/1323/70369#FormatConfigs)



## 数据加速器 GooseFS(goosefs) 版本：2022-05-19

### 第 26 次发布

发布时间：2026-10-29 02:01:54

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateUserQuota](https://cloud.tencent.com/document/api/1424/140321)
* [DeleteUserQuota](https://cloud.tencent.com/document/api/1424/140320)
* [DescribeUserQuota](https://cloud.tencent.com/document/api/1424/140319)
* [UpdateUserQuota](https://cloud.tencent.com/document/api/1424/140318)

新增数据结构：

* [UserQuotaInfo](https://cloud.tencent.com/document/api/1424/95076#UserQuotaInfo)



## 物联网开发平台(iotexplorer) 版本：2019-04-23

### 第 167 次发布

发布时间：2026-10-29 02:05:33

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除接口**：</font>

* CreateTWeTalkProductConfig
* CreateTWeTalkProductConfigV2
* DeleteTWeTalkProductConfigV2
* DescribeTWeTalkProductConfig
* DescribeTWeTalkProductConfigV2
* GetTWeTalkProductConfigList
* GetTWeTalkProductConfigListV2
* ModifyTWeTalkProductConfig
* ModifyTWeTalkProductConfigV2

<font color="#dd0000">**删除数据结构**：</font>

* IdleResponseInfo
* TalkBasicConfigInfo
* TalkConversationConfigInfo
* TalkIdleDetectionConfigInfo
* TalkLLMConfigInfo
* TalkProductConfigInfo
* TalkProductConfigV2Info
* TalkSTTConfigInfo
* TalkTTSConfigInfo



## 语音合成(tts) 版本：2019-08-23

### 第 16 次发布

发布时间：2026-10-29 02:54:12

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [TextToVoice](https://cloud.tencent.com/document/api/1073/37995)

	* 新增入参：EnableLatex




## Web 应用防火墙(waf) 版本：2018-01-25

### 第 171 次发布

发布时间：2026-10-29 03:00:57

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DeleteLLMSdkAccess](https://cloud.tencent.com/document/api/627/140322)



