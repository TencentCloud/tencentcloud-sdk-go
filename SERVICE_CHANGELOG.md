# Release v1.3.180

## 腾讯云智能体开发平台(adp) 版本：2026-05-20

### 第 24 次发布

发布时间：2026-09-14 20:48:54

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [CheckLabel](https://cloud.tencent.com/document/api/1759/138154)
* [CreateCategory](https://cloud.tencent.com/document/api/1759/138153)
* [CreateKB](https://cloud.tencent.com/document/api/1759/138152)
* [CreateLabel](https://cloud.tencent.com/document/api/1759/138151)
* [CreateQAGenerationTask](https://cloud.tencent.com/document/api/1759/138150)
* [CreateQAList](https://cloud.tencent.com/document/api/1759/138149)
* [CreateSimilarQuestion](https://cloud.tencent.com/document/api/1759/138148)
* [DeleteCategory](https://cloud.tencent.com/document/api/1759/138147)
* [DeleteDocList](https://cloud.tencent.com/document/api/1759/138146)
* [DeleteKB](https://cloud.tencent.com/document/api/1759/138145)
* [DeleteLabelList](https://cloud.tencent.com/document/api/1759/138144)
* [DeleteQAList](https://cloud.tencent.com/document/api/1759/138143)
* [DescribeCategoryList](https://cloud.tencent.com/document/api/1759/138142)
* [DescribeConflictQA](https://cloud.tencent.com/document/api/1759/138141)
* [DescribeConflictQASummaryList](https://cloud.tencent.com/document/api/1759/138140)
* [DescribeDoc](https://cloud.tencent.com/document/api/1759/138139)
* [DescribeDocSummaryList](https://cloud.tencent.com/document/api/1759/138138)
* [DescribeKB](https://cloud.tencent.com/document/api/1759/138137)
* [DescribeKBSummaryList](https://cloud.tencent.com/document/api/1759/138136)
* [DescribeLabel](https://cloud.tencent.com/document/api/1759/138135)
* [DescribeLabelSummaryList](https://cloud.tencent.com/document/api/1759/138134)
* [DescribeQA](https://cloud.tencent.com/document/api/1759/138133)
* [DescribeQASummaryList](https://cloud.tencent.com/document/api/1759/138132)
* [ExportQA](https://cloud.tencent.com/document/api/1759/138131)
* [ImportDocList](https://cloud.tencent.com/document/api/1759/138130)
* [ModifyCategory](https://cloud.tencent.com/document/api/1759/138129)
* [ModifyConflictQA](https://cloud.tencent.com/document/api/1759/138128)
* [ModifyDoc](https://cloud.tencent.com/document/api/1759/138127)
* [ModifyDocList](https://cloud.tencent.com/document/api/1759/138126)
* [ModifyKB](https://cloud.tencent.com/document/api/1759/138125)
* [ModifyLabel](https://cloud.tencent.com/document/api/1759/138124)
* [ModifyQA](https://cloud.tencent.com/document/api/1759/138123)
* [ModifyQAList](https://cloud.tencent.com/document/api/1759/138122)
* [SearchKnowledge](https://cloud.tencent.com/document/api/1759/138121)

新增数据结构：

* [CategoryModifyFields](https://cloud.tencent.com/document/api/1759/132545#CategoryModifyFields)
* [CategoryPath](https://cloud.tencent.com/document/api/1759/132545#CategoryPath)
* [CheckResult](https://cloud.tencent.com/document/api/1759/132545#CheckResult)
* [ConflictQA](https://cloud.tencent.com/document/api/1759/132545#ConflictQA)
* [ConflictQASummary](https://cloud.tencent.com/document/api/1759/132545#ConflictQASummary)
* [ContentFilter](https://cloud.tencent.com/document/api/1759/132545#ContentFilter)
* [DBRetrievalConfig](https://cloud.tencent.com/document/api/1759/132545#DBRetrievalConfig)
* [DeDuplicateStrategy](https://cloud.tencent.com/document/api/1759/132545#DeDuplicateStrategy)
* [DocExternalLink](https://cloud.tencent.com/document/api/1759/132545#DocExternalLink)
* [DocImportSpec](https://cloud.tencent.com/document/api/1759/132545#DocImportSpec)
* [DocLifecycle](https://cloud.tencent.com/document/api/1759/132545#DocLifecycle)
* [DocLink](https://cloud.tencent.com/document/api/1759/132545#DocLink)
* [DocMetadata](https://cloud.tencent.com/document/api/1759/132545#DocMetadata)
* [DocModifyFields](https://cloud.tencent.com/document/api/1759/132545#DocModifyFields)
* [DocOperator](https://cloud.tencent.com/document/api/1759/132545#DocOperator)
* [DocParseConfig](https://cloud.tencent.com/document/api/1759/132545#DocParseConfig)
* [DocPermission](https://cloud.tencent.com/document/api/1759/132545#DocPermission)
* [DocQuery](https://cloud.tencent.com/document/api/1759/132545#DocQuery)
* [DocRetrievalConfig](https://cloud.tencent.com/document/api/1759/132545#DocRetrievalConfig)
* [DocSummary](https://cloud.tencent.com/document/api/1759/132545#DocSummary)
* [DocSwitch](https://cloud.tencent.com/document/api/1759/132545#DocSwitch)
* [DocTaskStatus](https://cloud.tencent.com/document/api/1759/132545#DocTaskStatus)
* [DocUpdatePeriod](https://cloud.tencent.com/document/api/1759/132545#DocUpdatePeriod)
* [ESConfig](https://cloud.tencent.com/document/api/1759/132545#ESConfig)
* [EffectivePeriod](https://cloud.tencent.com/document/api/1759/132545#EffectivePeriod)
* [ExpirationAwareness](https://cloud.tencent.com/document/api/1759/132545#ExpirationAwareness)
* [ExpirationPolicy](https://cloud.tencent.com/document/api/1759/132545#ExpirationPolicy)
* [FinalRerankConfig](https://cloud.tencent.com/document/api/1759/132545#FinalRerankConfig)
* [GraphRAG](https://cloud.tencent.com/document/api/1759/132545#GraphRAG)
* [KBCapacity](https://cloud.tencent.com/document/api/1759/132545#KBCapacity)
* [KBCategory](https://cloud.tencent.com/document/api/1759/132545#KBCategory)
* [KBModelConfig](https://cloud.tencent.com/document/api/1759/132545#KBModelConfig)
* [KBModifyExtendFields](https://cloud.tencent.com/document/api/1759/132545#KBModifyExtendFields)
* [KBRetrievalConfig](https://cloud.tencent.com/document/api/1759/132545#KBRetrievalConfig)
* [KBSpec](https://cloud.tencent.com/document/api/1759/132545#KBSpec)
* [KBSummary](https://cloud.tencent.com/document/api/1759/132545#KBSummary)
* [KVPair](https://cloud.tencent.com/document/api/1759/132545#KVPair)
* [KnowledgeResult](https://cloud.tencent.com/document/api/1759/132545#KnowledgeResult)
* [KnowledgeScope](https://cloud.tencent.com/document/api/1759/132545#KnowledgeScope)
* [KnowledgeSnippetProfile](https://cloud.tencent.com/document/api/1759/132545#KnowledgeSnippetProfile)
* [KnowledgeSource](https://cloud.tencent.com/document/api/1759/132545#KnowledgeSource)
* [LabelModifyFields](https://cloud.tencent.com/document/api/1759/132545#LabelModifyFields)
* [LabelRef](https://cloud.tencent.com/document/api/1759/132545#LabelRef)
* [LabelRefIdentity](https://cloud.tencent.com/document/api/1759/132545#LabelRefIdentity)
* [LabelRefIdentityList](https://cloud.tencent.com/document/api/1759/132545#LabelRefIdentityList)
* [LabelSummary](https://cloud.tencent.com/document/api/1759/132545#LabelSummary)
* [LabelTerm](https://cloud.tencent.com/document/api/1759/132545#LabelTerm)
* [LabelTermCheckResult](https://cloud.tencent.com/document/api/1759/132545#LabelTermCheckResult)
* [LabelTermModifyItem](https://cloud.tencent.com/document/api/1759/132545#LabelTermModifyItem)
* [MetaValue](https://cloud.tencent.com/document/api/1759/132545#MetaValue)
* [OperationResult](https://cloud.tencent.com/document/api/1759/132545#OperationResult)
* [Operator](https://cloud.tencent.com/document/api/1759/132545#Operator)
* [QACreateSpec](https://cloud.tencent.com/document/api/1759/132545#QACreateSpec)
* [QALifecycle](https://cloud.tencent.com/document/api/1759/132545#QALifecycle)
* [QAMetadata](https://cloud.tencent.com/document/api/1759/132545#QAMetadata)
* [QAModifyFields](https://cloud.tencent.com/document/api/1759/132545#QAModifyFields)
* [QAOperator](https://cloud.tencent.com/document/api/1759/132545#QAOperator)
* [QAPermission](https://cloud.tencent.com/document/api/1759/132545#QAPermission)
* [QAQuery](https://cloud.tencent.com/document/api/1759/132545#QAQuery)
* [QARetrievalConfig](https://cloud.tencent.com/document/api/1759/132545#QARetrievalConfig)
* [QASegmentHighlight](https://cloud.tencent.com/document/api/1759/132545#QASegmentHighlight)
* [QASourceInfo](https://cloud.tencent.com/document/api/1759/132545#QASourceInfo)
* [QASummary](https://cloud.tencent.com/document/api/1759/132545#QASummary)
* [RerankConfig](https://cloud.tencent.com/document/api/1759/132545#RerankConfig)
* [RetrievalOption](https://cloud.tencent.com/document/api/1759/132545#RetrievalOption)
* [SearchAdvancedConfig](https://cloud.tencent.com/document/api/1759/132545#SearchAdvancedConfig)
* [SearchBilling](https://cloud.tencent.com/document/api/1759/132545#SearchBilling)
* [SearchContext](https://cloud.tencent.com/document/api/1759/132545#SearchContext)
* [SearchFilter](https://cloud.tencent.com/document/api/1759/132545#SearchFilter)
* [SearchFilterConfig](https://cloud.tencent.com/document/api/1759/132545#SearchFilterConfig)
* [SearchInput](https://cloud.tencent.com/document/api/1759/132545#SearchInput)
* [SearchResultPayload](https://cloud.tencent.com/document/api/1759/132545#SearchResultPayload)
* [SimilarQuestion](https://cloud.tencent.com/document/api/1759/132545#SimilarQuestion)
* [SimilarQuestionExtra](https://cloud.tencent.com/document/api/1759/132545#SimilarQuestionExtra)
* [SimilarQuestionModifySpec](https://cloud.tencent.com/document/api/1759/132545#SimilarQuestionModifySpec)
* [SimilarQuestionStat](https://cloud.tencent.com/document/api/1759/132545#SimilarQuestionStat)
* [SummaryListSwitch](https://cloud.tencent.com/document/api/1759/132545#SummaryListSwitch)
* [TableEnhancement](https://cloud.tencent.com/document/api/1759/132545#TableEnhancement)
* [TokenUsage](https://cloud.tencent.com/document/api/1759/132545#TokenUsage)
* [UserAccessConfig](https://cloud.tencent.com/document/api/1759/132545#UserAccessConfig)



## 云联络中心(ccc) 版本：2020-02-10

### 第 135 次发布

发布时间：2026-09-15 01:22:42

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [AICallAPICallAttempt](https://cloud.tencent.com/document/api/679/47715#AICallAPICallAttempt)
* [AICallAPICallDetail](https://cloud.tencent.com/document/api/679/47715#AICallAPICallDetail)
* [AICallAPICallRequestDetail](https://cloud.tencent.com/document/api/679/47715#AICallAPICallRequestDetail)
* [AICallAPICallResponseDetail](https://cloud.tencent.com/document/api/679/47715#AICallAPICallResponseDetail)

修改数据结构：

* [AIRoundPath](https://cloud.tencent.com/document/api/679/47715#AIRoundPath)

	* 新增成员：APICall

* [AISpeakEvent](https://cloud.tencent.com/document/api/679/47715#AISpeakEvent)

	* 新增成员：KnowledgeName




## 消息队列 CKafka 版(ckafka) 版本：2019-08-19

### 第 153 次发布

发布时间：2026-09-15 01:31:30

本次发布包含了以下内容：

改善已有的文档。

修改数据结构：

* [EsParam](https://cloud.tencent.com/document/api/597/40861#EsParam)

	* 新增成员：SourceDataFormat




## 全球加速(ga2) 版本：2025-01-15

### 第 14 次发布

发布时间：2026-09-15 02:08:47

本次发布包含了以下内容：

改善已有的文档。

修改接口：

* [DescribeAccessLogParam](https://cloud.tencent.com/document/api/1817/134690)

	* 新增入参：GlobalAcceleratorId

* [DescribeForwardingPolicy](https://cloud.tencent.com/document/api/1817/132227)

	* 新增入参：Filters

* [DescribeForwardingRule](https://cloud.tencent.com/document/api/1817/130162)

	* 新增入参：Filters

* [DescribeGlobalAcceleratorAclRules](https://cloud.tencent.com/document/api/1817/134687)

	* 新增入参：Filters

* [ModifyListener](https://cloud.tencent.com/document/api/1817/130155)

	* 新增入参：HttpVersion


修改数据结构：

* [EndpointGroupConfigurationSet](https://cloud.tencent.com/document/api/1817/130045#EndpointGroupConfigurationSet)

	* 新增成员：OriginPrivateIps, OriginPublicCidrs

* [GlobalAcceleratorAccessLog](https://cloud.tencent.com/document/api/1817/130045#GlobalAcceleratorAccessLog)

	* 新增成员：LogPushTaskId, GlobalAcceleratorId, ListenerId, EndpointGroupId, FlowLogDescription, CloudRegion, CloudLogId, CloudLogSetId, FieldKeys, Status




## 物联网开发平台(iotexplorer) 版本：2019-04-23

### 第 158 次发布

发布时间：2026-09-14 16:20:51

本次发布包含了以下内容：

改善已有的文档。

新增接口：

* [BatchDescribeTWeSeeOrders](https://cloud.tencent.com/document/api/1081/138113)
* [OperateTWeSeeDirectUploadObject](https://cloud.tencent.com/document/api/1081/138112)

新增数据结构：

* [BatchDescribeTWeSeeOrdersEntry](https://cloud.tencent.com/document/api/1081/34988#BatchDescribeTWeSeeOrdersEntry)
* [BatchDescribeTWeSeeOrdersResult](https://cloud.tencent.com/document/api/1081/34988#BatchDescribeTWeSeeOrdersResult)
* [SeeObjectListOptions](https://cloud.tencent.com/document/api/1081/34988#SeeObjectListOptions)
* [SeeObjectListing](https://cloud.tencent.com/document/api/1081/34988#SeeObjectListing)
* [SeeObjectMetadata](https://cloud.tencent.com/document/api/1081/34988#SeeObjectMetadata)
* [SeeObjectMetadataEntry](https://cloud.tencent.com/document/api/1081/34988#SeeObjectMetadataEntry)
* [SeeObjectSummary](https://cloud.tencent.com/document/api/1081/34988#SeeObjectSummary)



## 边缘安全加速平台(teo) 版本：2022-09-01

### 第 162 次发布

发布时间：2026-09-15 03:40:10

本次发布包含了以下内容：

改善已有的文档。

新增数据结构：

* [CustomVariable](https://cloud.tencent.com/document/api/1552/80721#CustomVariable)
* [CustomVariableOperation](https://cloud.tencent.com/document/api/1552/80721#CustomVariableOperation)
* [CustomVariableOperationRuleAction](https://cloud.tencent.com/document/api/1552/80721#CustomVariableOperationRuleAction)
* [CustomVariableOperationRuleBranch](https://cloud.tencent.com/document/api/1552/80721#CustomVariableOperationRuleBranch)
* [CustomVariableOperationSubRule](https://cloud.tencent.com/document/api/1552/80721#CustomVariableOperationSubRule)
* [SetParameters](https://cloud.tencent.com/document/api/1552/80721#SetParameters)

修改数据结构：

* [RuleEngineItem](https://cloud.tencent.com/document/api/1552/80721#RuleEngineItem)

	* 新增成员：CustomVariables, CustomVariableOperations




## 边缘安全加速平台(teo) 版本：2022-01-06



