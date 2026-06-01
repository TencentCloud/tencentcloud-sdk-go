# Release v1.3.108

## AI Agent 安全网关(apis) 版本：2024-08-01

### 第 8 次发布

发布时间：2026-06-02 01:11:30

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateAgentApp](https://cloud.tencent.com/document/api/1627/129628)

	* 新增入参：ConnectorIDs

* [ModifyAgentApp](https://cloud.tencent.com/document/api/1627/129618)

	* 新增入参：ConnectorIDs


修改数据结构：

* [DescribeAgentAppResp](https://cloud.tencent.com/document/api/1627/129635#DescribeAgentAppResp)

	* 新增成员：ConnectorIDs, ServicesNum

* [DescribeAgentCredentialResp](https://cloud.tencent.com/document/api/1627/129635#DescribeAgentCredentialResp)

	* 新增成员：RelateServiceNum




## 商业智能分析 BI(bi) 版本：2022-01-05

### 第 35 次发布

发布时间：2026-06-02 01:14:57

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateUserGroup](https://cloud.tencent.com/document/api/590/130845)

	* 新增入参：AdminUserId, Description

	* 新增出参：Extra, Msg, Data

* [CreateUserGroupMember](https://cloud.tencent.com/document/api/590/130844)

	* 新增入参：GroupId, UserIdList

	* 新增出参：Extra, Msg, Data

* [DeleteUserGroup](https://cloud.tencent.com/document/api/590/130843)

	* 新增出参：Extra, Msg, Data

* [DeleteUserGroupMember](https://cloud.tencent.com/document/api/590/130842)

	* 新增出参：Extra, Msg, Data

* [DescribeUserGroupMemberList](https://cloud.tencent.com/document/api/590/130840)

	* 新增出参：Extra, Msg, Data

* [ModifyUserGroup](https://cloud.tencent.com/document/api/590/130839)

	* 新增出参：Extra, Msg, Data

* [QueryUserGroupMember](https://cloud.tencent.com/document/api/590/130838)

	* 新增出参：Extra, Msg, Data


新增数据结构：

* [DescribeUserGroupMemberPageListContainer](https://cloud.tencent.com/document/api/590/73726#DescribeUserGroupMemberPageListContainer)
* [UserGroupMemberVO](https://cloud.tencent.com/document/api/590/73726#UserGroupMemberVO)
* [UserGroupUserInfoVO](https://cloud.tencent.com/document/api/590/73726#UserGroupUserInfoVO)
* [UserGroupVO](https://cloud.tencent.com/document/api/590/73726#UserGroupVO)



## 文件存储(cfs) 版本：2019-07-19

### 第 50 次发布

发布时间：2026-06-02 01:24:36

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateDataRetrieval](https://cloud.tencent.com/document/api/582/132388)
* [DeleteDataRetrieval](https://cloud.tencent.com/document/api/582/132387)
* [DescribeDataRetrieval](https://cloud.tencent.com/document/api/582/132386)
* [DescribeDataRetrievalTask](https://cloud.tencent.com/document/api/582/132385)
* [ModifyDataRetrieval](https://cloud.tencent.com/document/api/582/132384)
* [RunDataRetrievalTask](https://cloud.tencent.com/document/api/582/132383)

新增数据结构：

* [DataRetrievalInfo](https://cloud.tencent.com/document/api/582/38175#DataRetrievalInfo)
* [DataRetrievalTaskInfo](https://cloud.tencent.com/document/api/582/38175#DataRetrievalTaskInfo)



## 日志服务(cls) 版本：2020-10-16

### 第 161 次发布

发布时间：2026-06-02 01:30:06

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [AccessControlRule](https://cloud.tencent.com/document/api/614/56471#AccessControlRule)

	* 新增成员：CidrBlocks, Action




## 暴露面管理服务(ctem) 版本：2023-11-28

### 第 18 次发布

发布时间：2026-06-02 01:35:31

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateCustomer](https://cloud.tencent.com/document/api/1755/120315)

	* 新增入参：PortScanQps, SingleIPTaskLimit, HighRiskAck, ScanRateAckChecklist

* [CreateJobRecord](https://cloud.tencent.com/document/api/1755/120319)

	* 新增入参：PortScanQps, SingleIPTaskLimit, HighRiskAck, ScanRateAckChecklist

* [ModifyCustomer](https://cloud.tencent.com/document/api/1755/120313)

	* 新增入参：PortScanQps, SingleIPTaskLimit, HighRiskAck, ScanRateAckChecklist


修改数据结构：

* [Customer](https://cloud.tencent.com/document/api/1755/120320#Customer)

	* 新增成员：SingleIPTaskLimit, PortScanQps

* [DisplayApiSec](https://cloud.tencent.com/document/api/1755/120320#DisplayApiSec)

	* 新增成员：AggregationCount

* [DisplayConfig](https://cloud.tencent.com/document/api/1755/120320#DisplayConfig)

	* 新增成员：AggregationCount

* [DisplayHttp](https://cloud.tencent.com/document/api/1755/120320#DisplayHttp)

	* 新增成员：AggregationCount

* [DisplayManage](https://cloud.tencent.com/document/api/1755/120320#DisplayManage)

	* 新增成员：AggregationCount

* [DisplayPort](https://cloud.tencent.com/document/api/1755/120320#DisplayPort)

	* 新增成员：AggregationCount

* [DisplaySensitiveInfo](https://cloud.tencent.com/document/api/1755/120320#DisplaySensitiveInfo)

	* 新增成员：AggregationCount

* [DisplaySubDomain](https://cloud.tencent.com/document/api/1755/120320#DisplaySubDomain)

	* 新增成员：AggregationCount

* [DisplaySuspiciousAsset](https://cloud.tencent.com/document/api/1755/120320#DisplaySuspiciousAsset)

	* 新增成员：AggregationCount




## 轻量应用服务器(lighthouse) 版本：2020-03-24

### 第 80 次发布

发布时间：2026-06-02 02:15:44

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeBlueprintsShareAcrossAccountInfos](https://cloud.tencent.com/document/api/1207/132389)

新增数据结构：

* [BlueprintShareAcrossAccountInfo](https://cloud.tencent.com/document/api/1207/47576#BlueprintShareAcrossAccountInfo)



## 云开发 CloudBase(tcb) 版本：2018-06-08

### 第 144 次发布

发布时间：2026-06-02 02:43:07

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeBillingInfo](https://cloud.tencent.com/document/api/876/94390)

	* 新增入参：EnvIds, Limit, Offset

	* 新增出参：Total




## 云托管 CloudBase Run(tcbr) 版本：2022-02-17

### 第 29 次发布

发布时间：2026-06-02 02:44:03

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [PublicNetConf](https://cloud.tencent.com/document/api/1243/75713#PublicNetConf)

修改数据结构：

* [DiffConfigItem](https://cloud.tencent.com/document/api/1243/75713#DiffConfigItem)

	* 新增成员：PublicNetConf

* [ServerBaseConfig](https://cloud.tencent.com/document/api/1243/75713#ServerBaseConfig)

	* 新增成员：PublicNetConf




## 边缘安全加速平台(teo) 版本：2022-09-01

### 第 147 次发布

发布时间：2026-06-02 02:52:35

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateFunctionReplica](https://cloud.tencent.com/document/api/1552/132393)
* [DeleteFunctionReplica](https://cloud.tencent.com/document/api/1552/132392)
* [DescribeFunctionReplicas](https://cloud.tencent.com/document/api/1552/132391)
* [ModifyFunctionReplica](https://cloud.tencent.com/document/api/1552/132390)

新增数据结构：

* [FunctionReplica](https://cloud.tencent.com/document/api/1552/80721#FunctionReplica)



## 边缘安全加速平台(teo) 版本：2022-01-06



## 云点播(vod) 版本：2024-07-18



## 云点播(vod) 版本：2018-07-17

### 第 261 次发布

发布时间：2026-06-02 03:08:04

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [MediaTranscodeItem](https://cloud.tencent.com/document/api/266/31773#MediaTranscodeItem)

	* 新增成员：FileId




