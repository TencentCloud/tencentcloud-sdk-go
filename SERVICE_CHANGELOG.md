# Release v1.3.198

## Agent 沙箱服务(ags) 版本：2025-09-20

### 第 30 次发布

发布时间：2026-10-23 01:10:08

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeSandboxInstances](https://cloud.tencent.com/document/api/1814/139848)



## 负载均衡(clb) 版本：2018-03-17

### 第 172 次发布

发布时间：2026-10-23 01:31:52

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyModelAttributes](https://cloud.tencent.com/document/api/214/133670)

	* 新增入参：ModelProvider, Protocol, CMRPrivateNetworkTunnelId, HostHeader, VerifySSL

* [ModifyServiceProviderModelAttributes](https://cloud.tencent.com/document/api/214/133669)

	* 新增入参：NewModelName, NewModelAlias


修改数据结构：

* [StickyConfig](https://cloud.tencent.com/document/api/214/30694#StickyConfig)

	* 新增成员：TimeoutSeconds




## 云安全一体化平台(csip) 版本：2022-11-21

### 第 111 次发布

发布时间：2026-10-23 01:38:39

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [ClusterListItem](https://cloud.tencent.com/document/api/664/90825#ClusterListItem)

	* 新增成员：TagItems, TagModifyInfo




## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 325 次发布

发布时间：2026-10-23 02:10:44

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateReleaseFlow](https://cloud.tencent.com/document/api/1323/88045)

	* 新增入参：FlowName


修改数据结构：

* [ApproverInfo](https://cloud.tencent.com/document/api/1323/70369#ApproverInfo)

	* 新增成员：PinCode

* [FillApproverInfo](https://cloud.tencent.com/document/api/1323/70369#FillApproverInfo)

	* 新增成员：PinCode

* [FlowCreateApprover](https://cloud.tencent.com/document/api/1323/70369#FlowCreateApprover)

	* 新增成员：PinCode

* [ReleasedApprover](https://cloud.tencent.com/document/api/1323/70369#ReleasedApprover)

	* 新增成员：PinCode




## 腾讯电子签（基础版）(essbasic) 版本：2021-05-26

### 第 278 次发布

发布时间：2026-10-23 02:12:33

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ChannelCreateReleaseFlow](https://cloud.tencent.com/document/api/1420/83461)

	* 新增入参：FlowName


修改数据结构：

* [FillApproverInfo](https://cloud.tencent.com/document/api/1420/61525#FillApproverInfo)

	* 新增成员：PinCode

* [FlowApproverInfo](https://cloud.tencent.com/document/api/1420/61525#FlowApproverInfo)

	* 新增成员：PinCode

* [ReleasedApprover](https://cloud.tencent.com/document/api/1420/61525#ReleasedApprover)

	* 新增成员：PinCode




## 腾讯电子签（基础版）(essbasic) 版本：2020-12-22



## 文字识别(ocr) 版本：2018-11-19

### 第 274 次发布

发布时间：2026-10-23 02:44:40

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [VerifyGeneralCardWarn](https://cloud.tencent.com/document/api/866/136016)

	* 新增入参：ReasoningPrompt, ReasoningConfig, IgnoreWatermarkCategories

	* 新增出参：ReasoningResult, IsKeyFieldTamper




## 集团账号管理(organization) 版本：2021-03-31

### 第 63 次发布

发布时间：2026-10-23 02:46:27

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [UpdateRoleConfiguration](https://cloud.tencent.com/document/api/850/109930)

	* 新增入参：IsApply


新增数据结构：

* [FailedAccount](https://cloud.tencent.com/document/api/850/67060#FailedAccount)

修改数据结构：

* [RoleConfiguration](https://cloud.tencent.com/document/api/850/67060#RoleConfiguration)

	* 新增成员：ApplyStatus, AppliedSessionDuration, SessionDurationUpdateTime, FailedAccounts




## 集团账号管理(organization) 版本：2018-12-25



## 云开发 CloudBase(tcb) 版本：2018-06-08

### 第 173 次发布

发布时间：2026-10-23 02:58:01

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateUserKey](https://cloud.tencent.com/document/api/876/139855)
* [DeleteUserKey](https://cloud.tencent.com/document/api/876/139854)
* [GetOneUserKey](https://cloud.tencent.com/document/api/876/139853)
* [GetUserKeyList](https://cloud.tencent.com/document/api/876/139852)
* [UpdateUserKey](https://cloud.tencent.com/document/api/876/139851)

新增数据结构：

* [APIsUserKeyInfo](https://cloud.tencent.com/document/api/876/34822#APIsUserKeyInfo)
* [AgentRuntimeCode](https://cloud.tencent.com/document/api/876/34822#AgentRuntimeCode)
* [AgentRuntimeConfig](https://cloud.tencent.com/document/api/876/34822#AgentRuntimeConfig)
* [AgentRuntimeSessionConfig](https://cloud.tencent.com/document/api/876/34822#AgentRuntimeSessionConfig)
* [BindService](https://cloud.tencent.com/document/api/876/34822#BindService)
* [BindServiceList](https://cloud.tencent.com/document/api/876/34822#BindServiceList)



## 边缘安全加速平台(teo) 版本：2022-09-01

### 第 169 次发布

发布时间：2026-10-23 03:07:21

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeLogAnalysisStatistics](https://cloud.tencent.com/document/api/1552/139856)

新增数据结构：

* [StatisticsRecord](https://cloud.tencent.com/document/api/1552/80721#StatisticsRecord)



## 边缘安全加速平台(teo) 版本：2022-01-06



## 容器服务(tke) 版本：2022-05-01

### 第 33 次发布

发布时间：2026-10-23 03:14:53

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [NativeNodeInfo](https://cloud.tencent.com/document/api/457/103206#NativeNodeInfo)

	* 新增成员：IsMigrated, ResourceId




## 容器服务(tke) 版本：2018-05-25



## 私有网络(vpc) 版本：2017-03-12

### 第 314 次发布

发布时间：2026-10-23 03:25:54

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyAddressAttribute](https://cloud.tencent.com/document/api/215/16704)

	* 新增入参：EipVisibleOnEni

* [ModifyIPv6AddressesAttributes](https://cloud.tencent.com/document/api/215/113675)

	* 新增入参：EipVisibleOnEni


修改数据结构：

* [Address](https://cloud.tencent.com/document/api/215/15824#Address)

	* 新增成员：EipVisibleOnEni




