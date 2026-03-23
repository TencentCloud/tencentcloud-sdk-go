# Release v1.3.63

## AI Agent 安全网关(apis) 版本：2024-08-01

### 第 2 次发布

发布时间：2026-03-24 01:11:04

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateAgentCredential](https://cloud.tencent.com/document/api/1627/129675)
* [DeleteAgentCredential](https://cloud.tencent.com/document/api/1627/129674)
* [DescribeAgentCredential](https://cloud.tencent.com/document/api/1627/129673)
* [DescribeAgentCredentials](https://cloud.tencent.com/document/api/1627/129672)
* [ModifyAgentCredential](https://cloud.tencent.com/document/api/1627/129671)

### 第 1 次发布

发布时间：2026-03-23 17:43:24

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateAgentApp](https://cloud.tencent.com/document/api/1627/129628)
* [CreateAgentAppMcpServers](https://cloud.tencent.com/document/api/1627/129627)
* [CreateAgentAppModelServices](https://cloud.tencent.com/document/api/1627/129626)
* [CreateMcpServer](https://cloud.tencent.com/document/api/1627/129634)
* [CreateModel](https://cloud.tencent.com/document/api/1627/129609)
* [CreateModelService](https://cloud.tencent.com/document/api/1627/129615)
* [DeleteAgentApp](https://cloud.tencent.com/document/api/1627/129625)
* [DeleteAgentAppMcpServers](https://cloud.tencent.com/document/api/1627/129624)
* [DeleteAgentAppModelServices](https://cloud.tencent.com/document/api/1627/129623)
* [DeleteMcpServer](https://cloud.tencent.com/document/api/1627/129633)
* [DeleteModel](https://cloud.tencent.com/document/api/1627/129608)
* [DeleteModelService](https://cloud.tencent.com/document/api/1627/129614)
* [DescribeAgentApp](https://cloud.tencent.com/document/api/1627/129622)
* [DescribeAgentAppMcpServers](https://cloud.tencent.com/document/api/1627/129621)
* [DescribeAgentAppModelServices](https://cloud.tencent.com/document/api/1627/129620)
* [DescribeAgentApps](https://cloud.tencent.com/document/api/1627/129619)
* [DescribeMcpServer](https://cloud.tencent.com/document/api/1627/129632)
* [DescribeMcpServers](https://cloud.tencent.com/document/api/1627/129631)
* [DescribeModel](https://cloud.tencent.com/document/api/1627/129607)
* [DescribeModelService](https://cloud.tencent.com/document/api/1627/129613)
* [DescribeModelServices](https://cloud.tencent.com/document/api/1627/129612)
* [DescribeModels](https://cloud.tencent.com/document/api/1627/129606)
* [ModifyAgentApp](https://cloud.tencent.com/document/api/1627/129618)
* [ModifyAgentAppModelServices](https://cloud.tencent.com/document/api/1627/129617)
* [ModifyMcpServer](https://cloud.tencent.com/document/api/1627/129630)
* [ModifyModel](https://cloud.tencent.com/document/api/1627/129605)
* [ModifyModelService](https://cloud.tencent.com/document/api/1627/129611)

新增数据结构：

* [AgentAppMcpServerDTO](https://cloud.tencent.com/document/api/1627/129635#AgentAppMcpServerDTO)
* [AgentAppMcpServerVO](https://cloud.tencent.com/document/api/1627/129635#AgentAppMcpServerVO)
* [AgentAppModelServiceDTO](https://cloud.tencent.com/document/api/1627/129635#AgentAppModelServiceDTO)
* [AgentAppSecretKeyVO](https://cloud.tencent.com/document/api/1627/129635#AgentAppSecretKeyVO)
* [AgentCredentialContentDTO](https://cloud.tencent.com/document/api/1627/129635#AgentCredentialContentDTO)
* [AgentCredentialContentHeaderDTO](https://cloud.tencent.com/document/api/1627/129635#AgentCredentialContentHeaderDTO)
* [BindMcpSecurityRuleDTO](https://cloud.tencent.com/document/api/1627/129635#BindMcpSecurityRuleDTO)
* [BindMcpSecurityRuleVO](https://cloud.tencent.com/document/api/1627/129635#BindMcpSecurityRuleVO)
* [CreateAgentAppResp](https://cloud.tencent.com/document/api/1627/129635#CreateAgentAppResp)
* [DescribeAgentAppMcpServersResp](https://cloud.tencent.com/document/api/1627/129635#DescribeAgentAppMcpServersResp)
* [DescribeAgentAppResp](https://cloud.tencent.com/document/api/1627/129635#DescribeAgentAppResp)
* [DescribeAgentAppsResp](https://cloud.tencent.com/document/api/1627/129635#DescribeAgentAppsResp)
* [DescribeAgentAppsSortDTO](https://cloud.tencent.com/document/api/1627/129635#DescribeAgentAppsSortDTO)
* [DescribeAgentCredentialResp](https://cloud.tencent.com/document/api/1627/129635#DescribeAgentCredentialResp)
* [DescribeMcpServerResponseVO](https://cloud.tencent.com/document/api/1627/129635#DescribeMcpServerResponseVO)
* [DescribeMcpServersResponseVO](https://cloud.tencent.com/document/api/1627/129635#DescribeMcpServersResponseVO)
* [DescribeModelResponseVO](https://cloud.tencent.com/document/api/1627/129635#DescribeModelResponseVO)
* [DescribeModelServiceResponseVO](https://cloud.tencent.com/document/api/1627/129635#DescribeModelServiceResponseVO)
* [DescribeModelServicesResponseVO](https://cloud.tencent.com/document/api/1627/129635#DescribeModelServicesResponseVO)
* [DescribeModelServicesSort](https://cloud.tencent.com/document/api/1627/129635#DescribeModelServicesSort)
* [DescribeModelsResponseVO](https://cloud.tencent.com/document/api/1627/129635#DescribeModelsResponseVO)
* [DescribeModelsSort](https://cloud.tencent.com/document/api/1627/129635#DescribeModelsSort)
* [IDNameVO](https://cloud.tencent.com/document/api/1627/129635#IDNameVO)
* [InvokeLimitConfigDTO](https://cloud.tencent.com/document/api/1627/129635#InvokeLimitConfigDTO)
* [IpConfig](https://cloud.tencent.com/document/api/1627/129635#IpConfig)
* [LimitWindowsDTO](https://cloud.tencent.com/document/api/1627/129635#LimitWindowsDTO)
* [McpInputOutSchema](https://cloud.tencent.com/document/api/1627/129635#McpInputOutSchema)
* [McpSecurityRule](https://cloud.tencent.com/document/api/1627/129635#McpSecurityRule)
* [McpSecurityRulesVO](https://cloud.tencent.com/document/api/1627/129635#McpSecurityRulesVO)
* [McpTool](https://cloud.tencent.com/document/api/1627/129635#McpTool)
* [McpToolAnnotation](https://cloud.tencent.com/document/api/1627/129635#McpToolAnnotation)
* [McpUrlObj](https://cloud.tencent.com/document/api/1627/129635#McpUrlObj)
* [PluginConfigDTO](https://cloud.tencent.com/document/api/1627/129635#PluginConfigDTO)
* [PluginFormValueDTO](https://cloud.tencent.com/document/api/1627/129635#PluginFormValueDTO)
* [ResultIDVO](https://cloud.tencent.com/document/api/1627/129635#ResultIDVO)
* [ResultIDsVO](https://cloud.tencent.com/document/api/1627/129635#ResultIDsVO)
* [StartEndTime](https://cloud.tencent.com/document/api/1627/129635#StartEndTime)
* [TargetHostDTO](https://cloud.tencent.com/document/api/1627/129635#TargetHostDTO)
* [TargetModelDTO](https://cloud.tencent.com/document/api/1627/129635#TargetModelDTO)
* [TmsConfigDTO](https://cloud.tencent.com/document/api/1627/129635#TmsConfigDTO)
* [TokenLimitConfigDTO](https://cloud.tencent.com/document/api/1627/129635#TokenLimitConfigDTO)
* [ToolConfigDTO](https://cloud.tencent.com/document/api/1627/129635#ToolConfigDTO)
* [ToolConfigVO](https://cloud.tencent.com/document/api/1627/129635#ToolConfigVO)



## 云安全一体化平台(csip) 版本：2022-11-21

### 第 69 次发布

发布时间：2026-03-24 01:31:42

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeAIAgentAssetList](https://cloud.tencent.com/document/api/664/129677)

新增数据结构：

* [AIAgentAsset](https://cloud.tencent.com/document/api/664/90825#AIAgentAsset)



## 主机安全(cwp) 版本：2018-02-28

### 第 160 次发布

发布时间：2026-03-24 01:34:33

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeSkillInfo](https://cloud.tencent.com/document/api/296/129678)

新增数据结构：

* [SkillInfo](https://cloud.tencent.com/document/api/296/19867#SkillInfo)



## 智能视图计算平台(iss) 版本：2023-05-17

### 第 30 次发布

发布时间：2026-03-24 02:23:34

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [DescribeDeviceData](https://cloud.tencent.com/document/api/1344/95952#DescribeDeviceData)

	* 新增成员：PushStreamSecureUrl




## 媒体处理(mps) 版本：2019-06-12

### 第 187 次发布

发布时间：2026-03-24 02:46:23

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [SyncDubbing](https://cloud.tencent.com/document/api/862/128038)

	* 新增入参：Output

	* 新增出参：AudioUrl, ExtInfo


新增数据结构：

* [SyncDubbingOutputOption](https://cloud.tencent.com/document/api/862/37615#SyncDubbingOutputOption)



## 文字识别(ocr) 版本：2018-11-19

### 第 240 次发布

发布时间：2026-03-24 02:53:37

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [TextWaybill](https://cloud.tencent.com/document/api/866/33527#TextWaybill)

	* 新增成员：MainWaybillNum




## 云数据库 PostgreSQL(postgres) 版本：2017-03-12

### 第 63 次发布

发布时间：2026-03-24 02:59:51

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateInstances](https://cloud.tencent.com/document/api/409/56107)

	* 新增入参：StorageType

* [DescribeClasses](https://cloud.tencent.com/document/api/409/89019)

	* 新增入参：StorageType

* [DescribeDBVersions](https://cloud.tencent.com/document/api/409/89018)

	* 新增入参：StorageType

* [DescribeProductConfig](https://cloud.tencent.com/document/api/409/16776)

	* 新增入参：StorageType

* [InquiryPriceCreateDBInstances](https://cloud.tencent.com/document/api/409/16777)

	* 新增入参：StorageType


修改数据结构：

* [DBInstance](https://cloud.tencent.com/document/api/409/16778#DBInstance)

	* 新增成员：DBInstanceStorageType




## 云点播(vod) 版本：2024-07-18



## 云点播(vod) 版本：2018-07-17

### 第 239 次发布

发布时间：2026-03-24 04:01:55

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [AigcVideoOutputConfig](https://cloud.tencent.com/document/api/266/31773#AigcVideoOutputConfig)

	* 新增成员：OffPeak




