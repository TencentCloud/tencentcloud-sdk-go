# Release v1.1.35

## 数据湖计算 DLC(dlc) 版本：2021-01-25

### 第 138 次发布

发布时间：2025-09-27 16:52:20

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [DescribeTaskResourceUsage](https://cloud.tencent.com/document/api/1342/123808)

修改接口：

* [CreateDataMaskStrategy](https://cloud.tencent.com/document/api/1342/122619)

	* 新增出参：StrategyId

* [CreateStandardEngineResourceGroup](https://cloud.tencent.com/document/api/1342/122138)

	* 新增入参：DriverGPUSpec, ExecutorGPUSpec, GPULimitSize, GPUSize, PythonGPUSpec

* [UpdateStandardEngineResourceGroupResourceInfo](https://cloud.tencent.com/document/api/1342/122128)

	* 新增入参：DriverGPUSpec, ExecutorGPUSpec, GPULimitSize, GPUSize, PythonGPUSpec


新增数据结构：

* [CoreInfo](https://cloud.tencent.com/document/api/1342/53778#CoreInfo)

修改数据结构：

* [TaskResultInfo](https://cloud.tencent.com/document/api/1342/53778#TaskResultInfo)

	* 新增成员：ResultSetEncode




