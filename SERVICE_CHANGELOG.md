# Release v1.1.39

## 云数据库 MySQL(cdb) 版本：2017-03-20

### 第 203 次发布

发布时间：2025-10-13 01:11:08

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeAuditConfig](https://cloud.tencent.com/document/api/236/45455)

	* 新增出参：IsOpening


修改数据结构：

* [InstanceDbAuditStatus](https://cloud.tencent.com/document/api/236/15878#InstanceDbAuditStatus)

	* 新增成员：TrialStatus, TrialStartTime, TrialDuration, TrialCloseTime, TrialDescribeLogHours




## 腾讯云数据仓库 TCHouse-D(cdwdoris) 版本：2021-12-28

### 第 57 次发布

发布时间：2025-10-13 01:12:29

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyUserPrivilegesV3](https://cloud.tencent.com/document/api/1387/109432)

	* 新增入参：ComputeGroupType




## 消息队列 CKafka 版(ckafka) 版本：2019-08-19

### 第 133 次发布

发布时间：2025-10-13 01:14:04

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [MqttConnectParam](https://cloud.tencent.com/document/api/597/40861#MqttConnectParam)

	* 新增成员：Port, ServiceVip, Ip




## 云应用(cloudapp) 版本：2022-05-30

### 第 6 次发布

发布时间：2025-10-13 01:14:30

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [SaleParam](https://cloud.tencent.com/document/api/1689/108875#SaleParam)

	* 新增成员：ParamId, ParamValueId, ModuleId, ModuleKey, ModuleName




## 数据库智能管家 DBbrain(dbbrain) 版本：2021-05-27



## 数据库智能管家 DBbrain(dbbrain) 版本：2019-10-16

### 第 19 次发布

发布时间：2025-10-13 01:16:37

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [MySqlProcess](https://cloud.tencent.com/document/api/1130/39561#MySqlProcess)

	* 新增成员：SqlType




## 专线接入(dc) 版本：2018-04-10

### 第 37 次发布

发布时间：2025-10-13 01:17:12

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyDirectConnectTunnelExtra](https://cloud.tencent.com/document/api/216/49100)

	* 新增入参：ImportDirectRoute


修改数据结构：

* [AccessPoint](https://cloud.tencent.com/document/api/216/18418#AccessPoint)

	* 新增成员：Version, AccessPointServiceType

* [BFDInfo](https://cloud.tencent.com/document/api/216/18418#BFDInfo)

	* 新增成员：EnableBfdMultiHop

* [CloudAttachInfo](https://cloud.tencent.com/document/api/216/18418#CloudAttachInfo)

	* 新增成员：IapCode, IdcPointType, BIapLinkProtected

* [CreateCasInput](https://cloud.tencent.com/document/api/216/18418#CreateCasInput)

	* 新增成员：IdcPointType, BIapLinkProtected

* [DirectConnect](https://cloud.tencent.com/document/api/216/18418#DirectConnect)

	* 新增成员：FaultReportContactEmail




## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 250 次发布

发布时间：2025-10-13 01:21:30

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeContractComparisonTask](https://cloud.tencent.com/document/api/1323/123953)



## 私有网络(vpc) 版本：2017-03-12

### 第 267 次发布

发布时间：2025-10-13 02:33:45

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeAddressTemplateGroupInstances](https://cloud.tencent.com/document/api/215/123956)
* [DescribeAddressTemplateInstances](https://cloud.tencent.com/document/api/215/42837)
* [DescribeServiceTemplateGroupInstances](https://cloud.tencent.com/document/api/215/123955)
* [DescribeServiceTemplateInstances](https://cloud.tencent.com/document/api/215/123954)

修改接口：

* [ModifyPrivateNatGatewayAttribute](https://cloud.tencent.com/document/api/215/107890)

	* 新增入参：DeletionProtectionEnabled

	* <font color="#dd0000">**修改入参**：</font>NatGatewayName




## 数据开发治理平台 WeData(wedata) 版本：2025-08-06

### 第 2 次发布

发布时间：2025-10-13 02:42:08

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [AddCalcEnginesToProject](https://cloud.tencent.com/document/api/1267/123977)
* [AssociateResourceGroupToProject](https://cloud.tencent.com/document/api/1267/123985)
* [CreateDataSource](https://cloud.tencent.com/document/api/1267/123976)
* [CreateProject](https://cloud.tencent.com/document/api/1267/123975)
* [CreateProjectMember](https://cloud.tencent.com/document/api/1267/123974)
* [CreateResourceGroup](https://cloud.tencent.com/document/api/1267/123984)
* [DeleteDataSource](https://cloud.tencent.com/document/api/1267/123973)
* [DeleteLineage](https://cloud.tencent.com/document/api/1267/123997)
* [DeleteProjectMember](https://cloud.tencent.com/document/api/1267/123972)
* [DeleteResourceGroup](https://cloud.tencent.com/document/api/1267/123983)
* [DisableProject](https://cloud.tencent.com/document/api/1267/123971)
* [DissociateResourceGroupFromProject](https://cloud.tencent.com/document/api/1267/123982)
* [EnableProject](https://cloud.tencent.com/document/api/1267/123970)
* [GetDataSource](https://cloud.tencent.com/document/api/1267/123969)
* [GetDataSourceRelatedTasks](https://cloud.tencent.com/document/api/1267/123968)
* [GetProject](https://cloud.tencent.com/document/api/1267/123967)
* [GetResourceGroupMetrics](https://cloud.tencent.com/document/api/1267/123981)
* [GetTable](https://cloud.tencent.com/document/api/1267/123996)
* [GetTableColumns](https://cloud.tencent.com/document/api/1267/123995)
* [GrantMemberProjectRole](https://cloud.tencent.com/document/api/1267/123966)
* [ListCatalog](https://cloud.tencent.com/document/api/1267/123994)
* [ListColumnLineage](https://cloud.tencent.com/document/api/1267/123993)
* [ListDataSources](https://cloud.tencent.com/document/api/1267/123965)
* [ListDatabase](https://cloud.tencent.com/document/api/1267/123992)
* [ListLineage](https://cloud.tencent.com/document/api/1267/123991)
* [ListProcessLineage](https://cloud.tencent.com/document/api/1267/123990)
* [ListProjectMembers](https://cloud.tencent.com/document/api/1267/123964)
* [ListProjectRoles](https://cloud.tencent.com/document/api/1267/123963)
* [ListProjects](https://cloud.tencent.com/document/api/1267/123962)
* [ListResourceGroups](https://cloud.tencent.com/document/api/1267/123980)
* [ListSchema](https://cloud.tencent.com/document/api/1267/123989)
* [ListTable](https://cloud.tencent.com/document/api/1267/123988)
* [ListTenantRoles](https://cloud.tencent.com/document/api/1267/123961)
* [RegisterLineage](https://cloud.tencent.com/document/api/1267/123987)
* [RemoveMemberProjectRole](https://cloud.tencent.com/document/api/1267/123960)
* [UpdateDataSource](https://cloud.tencent.com/document/api/1267/123959)
* [UpdateProject](https://cloud.tencent.com/document/api/1267/123958)
* [UpdateResourceGroup](https://cloud.tencent.com/document/api/1267/123979)

新增数据结构：

* [BindProject](https://cloud.tencent.com/document/api/1267/123643#BindProject)
* [BriefTask](https://cloud.tencent.com/document/api/1267/123643#BriefTask)
* [BusinessMetadata](https://cloud.tencent.com/document/api/1267/123643#BusinessMetadata)
* [CatalogInfo](https://cloud.tencent.com/document/api/1267/123643#CatalogInfo)
* [ColumnInfo](https://cloud.tencent.com/document/api/1267/123643#ColumnInfo)
* [CreateProjectResult](https://cloud.tencent.com/document/api/1267/123643#CreateProjectResult)
* [DLCClusterInfo](https://cloud.tencent.com/document/api/1267/123643#DLCClusterInfo)
* [DataSource](https://cloud.tencent.com/document/api/1267/123643#DataSource)
* [DataSourceFileUpload](https://cloud.tencent.com/document/api/1267/123643#DataSourceFileUpload)
* [DataSourceInfo](https://cloud.tencent.com/document/api/1267/123643#DataSourceInfo)
* [DataSourceResult](https://cloud.tencent.com/document/api/1267/123643#DataSourceResult)
* [DataSourceStatus](https://cloud.tencent.com/document/api/1267/123643#DataSourceStatus)
* [DatabaseInfo](https://cloud.tencent.com/document/api/1267/123643#DatabaseInfo)
* [DatasourceRelationTaskInfo](https://cloud.tencent.com/document/api/1267/123643#DatasourceRelationTaskInfo)
* [ExecutorResourceGroupData](https://cloud.tencent.com/document/api/1267/123643#ExecutorResourceGroupData)
* [ExecutorResourceGroupInfo](https://cloud.tencent.com/document/api/1267/123643#ExecutorResourceGroupInfo)
* [IntegrationResource](https://cloud.tencent.com/document/api/1267/123643#IntegrationResource)
* [LineageNodeInfo](https://cloud.tencent.com/document/api/1267/123643#LineageNodeInfo)
* [LineagePair](https://cloud.tencent.com/document/api/1267/123643#LineagePair)
* [LineageProcess](https://cloud.tencent.com/document/api/1267/123643#LineageProcess)
* [LineageProperty](https://cloud.tencent.com/document/api/1267/123643#LineageProperty)
* [LineageRelation](https://cloud.tencent.com/document/api/1267/123643#LineageRelation)
* [LineageResouce](https://cloud.tencent.com/document/api/1267/123643#LineageResouce)
* [LineageResource](https://cloud.tencent.com/document/api/1267/123643#LineageResource)
* [ListCatalogPage](https://cloud.tencent.com/document/api/1267/123643#ListCatalogPage)
* [ListDatabasePage](https://cloud.tencent.com/document/api/1267/123643#ListDatabasePage)
* [ListLineagePage](https://cloud.tencent.com/document/api/1267/123643#ListLineagePage)
* [ListProcessLineagePage](https://cloud.tencent.com/document/api/1267/123643#ListProcessLineagePage)
* [ListSchemaPage](https://cloud.tencent.com/document/api/1267/123643#ListSchemaPage)
* [ListTablePage](https://cloud.tencent.com/document/api/1267/123643#ListTablePage)
* [MetricData](https://cloud.tencent.com/document/api/1267/123643#MetricData)
* [OperateResult](https://cloud.tencent.com/document/api/1267/123643#OperateResult)
* [PageRoles](https://cloud.tencent.com/document/api/1267/123643#PageRoles)
* [Project](https://cloud.tencent.com/document/api/1267/123643#Project)
* [ProjectBrief](https://cloud.tencent.com/document/api/1267/123643#ProjectBrief)
* [ProjectRequest](https://cloud.tencent.com/document/api/1267/123643#ProjectRequest)
* [ProjectResult](https://cloud.tencent.com/document/api/1267/123643#ProjectResult)
* [ProjectUserRole](https://cloud.tencent.com/document/api/1267/123643#ProjectUserRole)
* [ProjectUsersBrief](https://cloud.tencent.com/document/api/1267/123643#ProjectUsersBrief)
* [RelateTask](https://cloud.tencent.com/document/api/1267/123643#RelateTask)
* [ResourceGroupMetrics](https://cloud.tencent.com/document/api/1267/123643#ResourceGroupMetrics)
* [ResourceGroupSpecification](https://cloud.tencent.com/document/api/1267/123643#ResourceGroupSpecification)
* [ResourceNumber](https://cloud.tencent.com/document/api/1267/123643#ResourceNumber)
* [ResourceResult](https://cloud.tencent.com/document/api/1267/123643#ResourceResult)
* [ResourceStatus](https://cloud.tencent.com/document/api/1267/123643#ResourceStatus)
* [ResourceType](https://cloud.tencent.com/document/api/1267/123643#ResourceType)
* [SchemaInfo](https://cloud.tencent.com/document/api/1267/123643#SchemaInfo)
* [SystemRole](https://cloud.tencent.com/document/api/1267/123643#SystemRole)
* [TableInfo](https://cloud.tencent.com/document/api/1267/123643#TableInfo)
* [TechnicalMetadata](https://cloud.tencent.com/document/api/1267/123643#TechnicalMetadata)
* [TrendData](https://cloud.tencent.com/document/api/1267/123643#TrendData)



## 数据开发治理平台 WeData(wedata) 版本：2021-08-20



