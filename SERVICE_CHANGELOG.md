# Release v1.1.5

## 应用性能监控(apm) 版本：2021-06-22

### 第 43 次发布

发布时间：2025-08-05 01:05:36

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [APMKV](https://cloud.tencent.com/document/api/1463/64927#APMKV)

	* <font color="#dd0000">**修改成员**：</font>Key, Value

* [ApmField](https://cloud.tencent.com/document/api/1463/64927#ApmField)

	* 新增成员：NameCN




## 云防火墙(cfw) 版本：2019-09-04

### 第 84 次发布

发布时间：2025-08-04 15:32:51

本次发布包含了以下内容：

改善已有的文档。

<font color="#dd0000">**删除接口**：</font>

* CreateIdsWhiteRule
* DeleteIdsWhiteRule
* DescribeIdsWhiteRule

修改接口：

* [ModifyAddressTemplate](https://cloud.tencent.com/document/api/1132/97919)

	* 新增出参：RuleLimitNum

* [ModifyAllRuleStatus](https://cloud.tencent.com/document/api/1132/49063)

	* 新增出参：RuleLimitNum


<font color="#dd0000">**删除数据结构**：</font>

* IdsWhiteInfo



## 数据湖计算 DLC(dlc) 版本：2021-01-25

### 第 132 次发布

发布时间：2025-08-05 01:33:09

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeEngineUsageInfo](https://cloud.tencent.com/document/api/1342/87876)

	* 新增出参：AvailPercent

* [ListTaskJobLogDetail](https://cloud.tencent.com/document/api/1342/75645)

	* 新增入参：DataEngineId, ResourceGroupId

	* <font color="#dd0000">**修改入参**：</font>TaskId


修改数据结构：

* [StandardEngineResourceGroupInfo](https://cloud.tencent.com/document/api/1342/53778#StandardEngineResourceGroupInfo)

	* 新增成员：PublicDomain, RegistryId, RegionName, LaunchTime




## 物联网开发平台(iotexplorer) 版本：2019-04-23

### 第 110 次发布

发布时间：2025-08-04 15:33:56

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateTWeSeeRecognitionTask](https://cloud.tencent.com/document/api/1081/118005)

	* 新增入参：SummaryConfig

* [InvokeTWeSeeRecognitionTask](https://cloud.tencent.com/document/api/1081/118004)

	* 新增入参：SummaryConfig

* [ModifyTWeSeeConfig](https://cloud.tencent.com/document/api/1081/118021)

	* 新增入参：SummaryConfig


新增数据结构：

* [VisionSummaryConfig](https://cloud.tencent.com/document/api/1081/34988#VisionSummaryConfig)

修改数据结构：

* [VisionRecognitionResult](https://cloud.tencent.com/document/api/1081/34988#VisionRecognitionResult)

	* 新增成员：ErrorCode




## TSF-应用管理&Consul(tsf) 版本：2018-03-26

### 第 128 次发布

发布时间：2025-08-05 02:37:59

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeLicenses](https://cloud.tencent.com/document/api/649/122280)
* [DescribeLogCapacity](https://cloud.tencent.com/document/api/649/122279)
* [DescribeResourceConfig](https://cloud.tencent.com/document/api/649/122278)

新增数据结构：

* [ContainerAdditionalResourceRequirement](https://cloud.tencent.com/document/api/649/36099#ContainerAdditionalResourceRequirement)
* [ContainerAdditionalResourceRequirementMap](https://cloud.tencent.com/document/api/649/36099#ContainerAdditionalResourceRequirementMap)
* [ContainerGroupResourceConfig](https://cloud.tencent.com/document/api/649/36099#ContainerGroupResourceConfig)
* [ContainerInstanceResourceConfig](https://cloud.tencent.com/document/api/649/36099#ContainerInstanceResourceConfig)
* [DescribeResourceConfigCluster](https://cloud.tencent.com/document/api/649/36099#DescribeResourceConfigCluster)
* [DescribeResourceConfigClusterContainer](https://cloud.tencent.com/document/api/649/36099#DescribeResourceConfigClusterContainer)
* [DescribeResourceConfigLicense](https://cloud.tencent.com/document/api/649/36099#DescribeResourceConfigLicense)
* [DescribeResourceConfigLicenseFunction](https://cloud.tencent.com/document/api/649/36099#DescribeResourceConfigLicenseFunction)
* [DescribeResourceConfigLicenseResource](https://cloud.tencent.com/document/api/649/36099#DescribeResourceConfigLicenseResource)
* [DescribeResourceConfigResultV2](https://cloud.tencent.com/document/api/649/36099#DescribeResourceConfigResultV2)
* [DescribeResourceConfigSts](https://cloud.tencent.com/document/api/649/36099#DescribeResourceConfigSts)
* [GroupResourceConfig](https://cloud.tencent.com/document/api/649/36099#GroupResourceConfig)
* [InstanceResourceConfig](https://cloud.tencent.com/document/api/649/36099#InstanceResourceConfig)
* [LicenseTag](https://cloud.tencent.com/document/api/649/36099#LicenseTag)
* [PackageConfig](https://cloud.tencent.com/document/api/649/36099#PackageConfig)
* [TsfPageLicenseTag](https://cloud.tencent.com/document/api/649/36099#TsfPageLicenseTag)
* [VmInstanceResourceConfig](https://cloud.tencent.com/document/api/649/36099#VmInstanceResourceConfig)



## 数据开发治理平台 WeData(wedata) 版本：2021-08-20

### 第 166 次发布

发布时间：2025-08-05 02:48:40

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [CreateBaseProject](https://cloud.tencent.com/document/api/1267/121425)

	* 新增入参：Project

* [ModifyProject](https://cloud.tencent.com/document/api/1267/121422)

	* 新增入参：DisplayName, Description, ModifyType


新增数据结构：

* [BaseProject](https://cloud.tencent.com/document/api/1267/76336#BaseProject)

修改数据结构：

* [DataSourceInfo](https://cloud.tencent.com/document/api/1267/76336#DataSourceInfo)

	* 新增成员：ForbidProbe




