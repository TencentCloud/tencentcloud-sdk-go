# Release v1.3.143

## 腾讯云智能体开发平台(adp) 版本：2026-05-20

### 第 12 次发布

发布时间：2026-07-24 13:13:21

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CreateAppTrigger](https://cloud.tencent.com/document/api/1759/135012)
* [CreateTimerTask](https://cloud.tencent.com/document/api/1759/135011)
* [DeleteAppTrigger](https://cloud.tencent.com/document/api/1759/135010)
* [DeleteTimerTask](https://cloud.tencent.com/document/api/1759/135009)
* [DescribeAppTrigger](https://cloud.tencent.com/document/api/1759/135008)
* [DescribeAppTriggerInstance](https://cloud.tencent.com/document/api/1759/135007)
* [DescribeAppTriggerRunLogList](https://cloud.tencent.com/document/api/1759/135006)
* [DescribeAppTriggerSummaryList](https://cloud.tencent.com/document/api/1759/135005)
* [DescribeTimerTask](https://cloud.tencent.com/document/api/1759/135004)
* [DescribeTimerTaskRunLogList](https://cloud.tencent.com/document/api/1759/135003)
* [DescribeTimerTaskSummaryList](https://cloud.tencent.com/document/api/1759/135002)
* [MarkAppTriggerRunLogRead](https://cloud.tencent.com/document/api/1759/135001)
* [MarkTimerTaskRunLogRead](https://cloud.tencent.com/document/api/1759/135000)
* [ModifyAppTrigger](https://cloud.tencent.com/document/api/1759/134999)
* [ModifyTimerTask](https://cloud.tencent.com/document/api/1759/134998)
* [PauseAppTrigger](https://cloud.tencent.com/document/api/1759/134997)
* [PauseTimerTask](https://cloud.tencent.com/document/api/1759/134996)
* [ResumeAppTrigger](https://cloud.tencent.com/document/api/1759/134995)
* [ResumeTimerTask](https://cloud.tencent.com/document/api/1759/134994)
* [RunAppTriggerNow](https://cloud.tencent.com/document/api/1759/134993)
* [RunTimerTaskNow](https://cloud.tencent.com/document/api/1759/134992)

新增数据结构：

* [AppTrigger](https://cloud.tencent.com/document/api/1759/132545#AppTrigger)
* [AppTriggerInstance](https://cloud.tencent.com/document/api/1759/132545#AppTriggerInstance)
* [AppTriggerParamBinding](https://cloud.tencent.com/document/api/1759/132545#AppTriggerParamBinding)
* [AppTriggerParamBindingConfig](https://cloud.tencent.com/document/api/1759/132545#AppTriggerParamBindingConfig)
* [AppTriggerParamBindingValue](https://cloud.tencent.com/document/api/1759/132545#AppTriggerParamBindingValue)
* [AppTriggerParamSchema](https://cloud.tencent.com/document/api/1759/132545#AppTriggerParamSchema)
* [AppTriggerPromptExecuteConfig](https://cloud.tencent.com/document/api/1759/132545#AppTriggerPromptExecuteConfig)
* [AppTriggerRunLog](https://cloud.tencent.com/document/api/1759/132545#AppTriggerRunLog)
* [AppTriggerScheduleConfig](https://cloud.tencent.com/document/api/1759/132545#AppTriggerScheduleConfig)
* [AppTriggerScheduleStatus](https://cloud.tencent.com/document/api/1759/132545#AppTriggerScheduleStatus)
* [AppTriggerSummary](https://cloud.tencent.com/document/api/1759/132545#AppTriggerSummary)
* [AppTriggerWebhookConfig](https://cloud.tencent.com/document/api/1759/132545#AppTriggerWebhookConfig)
* [AppTriggerWebhookParamSchemaConfig](https://cloud.tencent.com/document/api/1759/132545#AppTriggerWebhookParamSchemaConfig)
* [AppTriggerWebhookStatus](https://cloud.tencent.com/document/api/1759/132545#AppTriggerWebhookStatus)
* [AppTriggerWorkflowExecuteConfig](https://cloud.tencent.com/document/api/1759/132545#AppTriggerWorkflowExecuteConfig)
* [CronSchedule](https://cloud.tencent.com/document/api/1759/132545#CronSchedule)
* [DailySchedule](https://cloud.tencent.com/document/api/1759/132545#DailySchedule)
* [ExecuteConfig](https://cloud.tencent.com/document/api/1759/132545#ExecuteConfig)
* [IntervalSchedule](https://cloud.tencent.com/document/api/1759/132545#IntervalSchedule)
* [ManualOnlySchedule](https://cloud.tencent.com/document/api/1759/132545#ManualOnlySchedule)
* [OnceSchedule](https://cloud.tencent.com/document/api/1759/132545#OnceSchedule)
* [TimerConfig](https://cloud.tencent.com/document/api/1759/132545#TimerConfig)
* [TimerProfile](https://cloud.tencent.com/document/api/1759/132545#TimerProfile)
* [TimerPushConfig](https://cloud.tencent.com/document/api/1759/132545#TimerPushConfig)
* [TimerScheduleConfig](https://cloud.tencent.com/document/api/1759/132545#TimerScheduleConfig)
* [TimerStatus](https://cloud.tencent.com/document/api/1759/132545#TimerStatus)
* [TimerTask](https://cloud.tencent.com/document/api/1759/132545#TimerTask)
* [TimerTaskSummary](https://cloud.tencent.com/document/api/1759/132545#TimerTaskSummary)
* [TriggerConfig](https://cloud.tencent.com/document/api/1759/132545#TriggerConfig)
* [TriggerStatus](https://cloud.tencent.com/document/api/1759/132545#TriggerStatus)
* [WeeklySchedule](https://cloud.tencent.com/document/api/1759/132545#WeeklySchedule)
* [WeeklyTime](https://cloud.tencent.com/document/api/1759/132545#WeeklyTime)



## 云防火墙(cfw) 版本：2019-09-04

### 第 111 次发布

发布时间：2026-07-27 01:25:36

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DeleteBlockIgnoreRuleNew](https://cloud.tencent.com/document/api/1132/104037)

	* <font color="#dd0000">**修改入参**：</font>ShowType




## 腾讯电子签企业版(ess) 版本：2020-11-11

### 第 310 次发布

发布时间：2026-07-27 02:00:17

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyPartnerAutoSignAuthUrl](https://cloud.tencent.com/document/api/1323/120051)

	* 新增入参：SealTypes




## 腾讯电子签（基础版）(essbasic) 版本：2021-05-26

### 第 267 次发布

发布时间：2026-07-27 02:02:02

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [ModifyPartnerAutoSignAuthUrl](https://cloud.tencent.com/document/api/1420/120052)

	* 新增入参：SealTypes




## 腾讯电子签（基础版）(essbasic) 版本：2020-12-22



## 媒体处理(mps) 版本：2019-06-12

### 第 224 次发布

发布时间：2026-07-24 18:05:05

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CloneViral](https://cloud.tencent.com/document/api/862/135033)
* [DescribeCloneViralTask](https://cloud.tencent.com/document/api/862/135032)

新增数据结构：

* [CloneViralAIGC](https://cloud.tencent.com/document/api/862/37615#CloneViralAIGC)
* [CloneViralContent](https://cloud.tencent.com/document/api/862/37615#CloneViralContent)
* [CloneViralPersona](https://cloud.tencent.com/document/api/862/37615#CloneViralPersona)
* [CloneViralProduct](https://cloud.tencent.com/document/api/862/37615#CloneViralProduct)



