// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20190819

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type Acl struct {
	// Acl资源类型，（0:UNKNOWN，1:ANY，2:TOPIC，3:GROUP，4:CLUSTER，5:TRANSACTIONAL_ID）当前只有TOPIC，
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 资源名称，和resourceType相关如当resourceType为TOPIC时，则该字段表示topic名称，当resourceType为GROUP时，该字段表示group名称
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 用户列表，默认为User:*，表示任何user都可以访问，当前用户只能是用户列表中包含的用户
	Principal *string `json:"Principal,omitnil,omitempty" name:"Principal"`

	// 默认\*,表示任何host都可以访问，当前ckafka不支持host为\*，但是后面开源kafka的产品化会直接支持
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// Acl操作方式(0:UNKNOWN，1:ANY，2:ALL，3:READ，4:WRITE，5:CREATE，6:DELETE，7:ALTER，8:DESCRIBE，9:CLUSTER_ACTION，10:DESCRIBE_CONFIGS，11:ALTER_CONFIGS，12:IDEMPOTEN_WRITE)
	Operation *int64 `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 权限类型(0:UNKNOWN，1:ANY，2:DENY，3:ALLOW)
	PermissionType *int64 `json:"PermissionType,omitnil,omitempty" name:"PermissionType"`
}

type AclResponse struct {
	// 符合条件的总数据条数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// ACL列表
	AclList []*Acl `json:"AclList,omitnil,omitempty" name:"AclList"`
}

type AclRule struct {
	// ACL规则名
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// ckafka集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// ACL规则匹配类型，目前只支持前缀匹配，枚举值列表：PREFIXED
	PatternType *string `json:"PatternType,omitnil,omitempty" name:"PatternType"`

	// 表示前缀匹配的前缀的值
	Pattern *string `json:"Pattern,omitnil,omitempty" name:"Pattern"`

	// Acl资源类型,目前只支持Topic,枚举值列表：Topic
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 该规则所包含的ACL信息
	AclList *string `json:"AclList,omitnil,omitempty" name:"AclList"`

	// 规则所创建的时间
	CreateTimeStamp *string `json:"CreateTimeStamp,omitnil,omitempty" name:"CreateTimeStamp"`

	// 预设ACL规则是否应用到新增的topic中
	IsApplied *int64 `json:"IsApplied,omitnil,omitempty" name:"IsApplied"`

	// 规则更新时间
	UpdateTimeStamp *string `json:"UpdateTimeStamp,omitnil,omitempty" name:"UpdateTimeStamp"`

	// 规则的备注
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 其中一个显示的对应的TopicName
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 应用该ACL规则的Topic数
	TopicCount *int64 `json:"TopicCount,omitnil,omitempty" name:"TopicCount"`

	// patternType的中文显示
	PatternTypeTitle *string `json:"PatternTypeTitle,omitnil,omitempty" name:"PatternTypeTitle"`
}

type AclRuleInfo struct {
	// Acl操作方式，枚举值(所有操作: All, 读：Read，写：Write)
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 权限类型，Deny：拒绝，Allow：允许。
	PermissionType *string `json:"PermissionType,omitnil,omitempty" name:"PermissionType"`

	// 表示任何host都可以访问
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 用户，User:*表示任何user都可以访问，当前用户只能是用户列表中包含的用户。传入格式需要带【User:】前缀。例如用户A，传入为User:A。
	Principal *string `json:"Principal,omitnil,omitempty" name:"Principal"`
}

type AclRuleResp struct {
	// 总数据条数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// AclRule列表
	AclRuleList []*AclRule `json:"AclRuleList,omitnil,omitempty" name:"AclRuleList"`
}

type AnalyseParam struct {
	// 解析格式，JSON，DELIMITER分隔符，REGULAR正则提取，SOURCE处理上层所有结果
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 分隔符、正则表达式
	Regex *string `json:"Regex,omitnil,omitempty" name:"Regex"`

	// 需再次处理的KEY——模式
	InputValueType *string `json:"InputValueType,omitnil,omitempty" name:"InputValueType"`

	// 需再次处理的KEY——KEY表达式
	InputValue *string `json:"InputValue,omitnil,omitempty" name:"InputValue"`
}

type Assignment struct {
	// assingment版本信息
	Version *int64 `json:"Version,omitnil,omitempty" name:"Version"`

	// topic信息列表
	Topics []*GroupInfoTopics `json:"Topics,omitnil,omitempty" name:"Topics"`
}

// Predefined struct for user
type AuthorizeTokenRequestParams struct {
	// ckafka集群实例Id, 可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// token串
	Tokens *string `json:"Tokens,omitnil,omitempty" name:"Tokens"`
}

type AuthorizeTokenRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id, 可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// token串
	Tokens *string `json:"Tokens,omitnil,omitempty" name:"Tokens"`
}

func (r *AuthorizeTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AuthorizeTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "User")
	delete(f, "Tokens")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AuthorizeTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AuthorizeTokenResponseParams struct {
	// 0 成功
	Result *int64 `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AuthorizeTokenResponse struct {
	*tchttp.BaseResponse
	Response *AuthorizeTokenResponseParams `json:"Response"`
}

func (r *AuthorizeTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AuthorizeTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchAnalyseParam struct {
	// ONE_BY_ONE单条输出，MERGE合并输出
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`
}

type BatchContent struct {
	// 发送的消息体
	Body *string `json:"Body,omitnil,omitempty" name:"Body"`

	// 发送消息的键名
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`
}

// Predefined struct for user
type BatchCreateAclRequestParams struct {
	// ckafka集群实例Id，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Acl资源类型，(2:TOPIC）
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 资源列表数组，可通过[DescribeTopic](https://cloud.tencent.com/document/product/597/40847)接口获取。
	ResourceNames []*string `json:"ResourceNames,omitnil,omitempty" name:"ResourceNames"`

	// 设置的ACL规则列表，可通过[DescribeAclRule](https://cloud.tencent.com/document/product/597/89217)接口获取。
	RuleList []*AclRuleInfo `json:"RuleList,omitnil,omitempty" name:"RuleList"`
}

type BatchCreateAclRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Acl资源类型，(2:TOPIC）
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 资源列表数组，可通过[DescribeTopic](https://cloud.tencent.com/document/product/597/40847)接口获取。
	ResourceNames []*string `json:"ResourceNames,omitnil,omitempty" name:"ResourceNames"`

	// 设置的ACL规则列表，可通过[DescribeAclRule](https://cloud.tencent.com/document/product/597/89217)接口获取。
	RuleList []*AclRuleInfo `json:"RuleList,omitnil,omitempty" name:"RuleList"`
}

func (r *BatchCreateAclRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchCreateAclRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ResourceType")
	delete(f, "ResourceNames")
	delete(f, "RuleList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchCreateAclRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchCreateAclResponseParams struct {
	// 状态码：0-修改成功，否则修改失败
	Result *int64 `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BatchCreateAclResponse struct {
	*tchttp.BaseResponse
	Response *BatchCreateAclResponseParams `json:"Response"`
}

func (r *BatchCreateAclResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchCreateAclResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchModifyGroupOffsetsRequestParams struct {
	// 消费分组名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// ckafka集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// partition信息
	Partitions []*Partitions `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// 指定topic，默认所有topic
	TopicName []*string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

type BatchModifyGroupOffsetsRequest struct {
	*tchttp.BaseRequest
	
	// 消费分组名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// ckafka集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// partition信息
	Partitions []*Partitions `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// 指定topic，默认所有topic
	TopicName []*string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

func (r *BatchModifyGroupOffsetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchModifyGroupOffsetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupName")
	delete(f, "InstanceId")
	delete(f, "Partitions")
	delete(f, "TopicName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchModifyGroupOffsetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchModifyGroupOffsetsResponseParams struct {
	// 返回结果
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BatchModifyGroupOffsetsResponse struct {
	*tchttp.BaseResponse
	Response *BatchModifyGroupOffsetsResponseParams `json:"Response"`
}

func (r *BatchModifyGroupOffsetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchModifyGroupOffsetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchModifyTopicAttributesRequestParams struct {
	// ckafka集群实例Id，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题属性列表 (同一个批次最多支持10个)，可通过[DescribeTopic](https://cloud.tencent.com/document/product/597/40847)接口获取
	Topic []*BatchModifyTopicInfo `json:"Topic,omitnil,omitempty" name:"Topic"`
}

type BatchModifyTopicAttributesRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题属性列表 (同一个批次最多支持10个)，可通过[DescribeTopic](https://cloud.tencent.com/document/product/597/40847)接口获取
	Topic []*BatchModifyTopicInfo `json:"Topic,omitnil,omitempty" name:"Topic"`
}

func (r *BatchModifyTopicAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchModifyTopicAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Topic")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchModifyTopicAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchModifyTopicAttributesResponseParams struct {
	// 返回结果
	Result []*BatchModifyTopicResultDTO `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BatchModifyTopicAttributesResponse struct {
	*tchttp.BaseResponse
	Response *BatchModifyTopicAttributesResponseParams `json:"Response"`
}

func (r *BatchModifyTopicAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchModifyTopicAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchModifyTopicInfo struct {
	// 主题名
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 分区数
	PartitionNum *int64 `json:"PartitionNum,omitnil,omitempty" name:"PartitionNum"`

	// 备注
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// 副本数
	ReplicaNum *int64 `json:"ReplicaNum,omitnil,omitempty" name:"ReplicaNum"`

	// 消息删除策略，可以选择delete 或者compact
	CleanUpPolicy *string `json:"CleanUpPolicy,omitnil,omitempty" name:"CleanUpPolicy"`

	// 当producer设置request.required.acks为-1时，min.insync.replicas指定replicas的最小数目
	MinInsyncReplicas *int64 `json:"MinInsyncReplicas,omitnil,omitempty" name:"MinInsyncReplicas"`

	// 是否允许非ISR的副本成为Leader
	UncleanLeaderElectionEnable *bool `json:"UncleanLeaderElectionEnable,omitnil,omitempty" name:"UncleanLeaderElectionEnable"`

	// topic维度的消息保留时间（毫秒）范围1 分钟到90 天
	RetentionMs *int64 `json:"RetentionMs,omitnil,omitempty" name:"RetentionMs"`

	// topic维度的消息保留大小，单位为Byte，范围1 GB到1024 GB。
	RetentionBytes *int64 `json:"RetentionBytes,omitnil,omitempty" name:"RetentionBytes"`

	// Segment分片滚动的时长（毫秒），范围1 天到90 天
	SegmentMs *int64 `json:"SegmentMs,omitnil,omitempty" name:"SegmentMs"`

	// 批次的消息大小，范围1 KB到12 MB
	MaxMessageBytes *int64 `json:"MaxMessageBytes,omitnil,omitempty" name:"MaxMessageBytes"`

	// 消息保存的时间类型：CreateTime/LogAppendTime
	LogMsgTimestampType *string `json:"LogMsgTimestampType,omitnil,omitempty" name:"LogMsgTimestampType"`
}

type BatchModifyTopicResultDTO struct {
	// ckafka集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 操作返回码
	ReturnCode *string `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// 操作返回信息
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

type BrokerTopicData struct {
	// 主题名称
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 主题Id
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 主题占用Broker 容量大小，单位为Bytes。
	DataSize *uint64 `json:"DataSize,omitnil,omitempty" name:"DataSize"`
}

type BrokerTopicFlowData struct {
	// 主题名
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 主题Id
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Topic 流量(MB)
	TopicTraffic *string `json:"TopicTraffic,omitnil,omitempty" name:"TopicTraffic"`
}

// Predefined struct for user
type CancelAuthorizationTokenRequestParams struct {
	// ckafka集群实例Id,可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// token串
	Tokens *string `json:"Tokens,omitnil,omitempty" name:"Tokens"`
}

type CancelAuthorizationTokenRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id,可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// token串
	Tokens *string `json:"Tokens,omitnil,omitempty" name:"Tokens"`
}

func (r *CancelAuthorizationTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelAuthorizationTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "User")
	delete(f, "Tokens")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelAuthorizationTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelAuthorizationTokenResponseParams struct {
	// 0 成功 非0 失败
	Result *int64 `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CancelAuthorizationTokenResponse struct {
	*tchttp.BaseResponse
	Response *CancelAuthorizationTokenResponseParams `json:"Response"`
}

func (r *CancelAuthorizationTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelAuthorizationTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CdcClusterResponse struct {
	// 任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

// Predefined struct for user
type CheckCdcClusterRequestParams struct {
	// 任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type CheckCdcClusterRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *CheckCdcClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckCdcClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckCdcClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckCdcClusterResponseParams struct {
	// 返回结果状态Success
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckCdcClusterResponse struct {
	*tchttp.BaseResponse
	Response *CheckCdcClusterResponseParams `json:"Response"`
}

func (r *CheckCdcClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckCdcClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClickHouseConnectParam struct {
	// ClickHouse的连接port
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// ClickHouse连接源的用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// ClickHouse连接源的密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// ClickHouse连接源的实例资源
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// ClickHouse连接源是否为自建集群
	SelfBuilt *bool `json:"SelfBuilt,omitnil,omitempty" name:"SelfBuilt"`

	// ClickHouse连接源的实例vip，当为腾讯云实例时，必填
	ServiceVip *string `json:"ServiceVip,omitnil,omitempty" name:"ServiceVip"`

	// ClickHouse连接源的vpcId，当为腾讯云实例时，必填
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 是否更新到关联的Datahub任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsUpdate *bool `json:"IsUpdate,omitnil,omitempty" name:"IsUpdate"`
}

type ClickHouseModifyConnectParam struct {
	// ClickHouse连接源的实例资源【不支持修改】
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// ClickHouse的连接port【不支持修改】
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// ClickHouse连接源的实例vip【不支持修改】
	ServiceVip *string `json:"ServiceVip,omitnil,omitempty" name:"ServiceVip"`

	// ClickHouse连接源的vpcId【不支持修改】
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// ClickHouse连接源的用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// ClickHouse连接源的密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// ClickHouse连接源是否为自建集群【不支持修改】
	SelfBuilt *bool `json:"SelfBuilt,omitnil,omitempty" name:"SelfBuilt"`

	// 是否更新到关联的Datahub任务，默认为true
	IsUpdate *bool `json:"IsUpdate,omitnil,omitempty" name:"IsUpdate"`
}

type ClickHouseParam struct {
	// ClickHouse的集群
	Cluster *string `json:"Cluster,omitnil,omitempty" name:"Cluster"`

	// ClickHouse的数据库名
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// ClickHouse的数据表名
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// ClickHouse的schema
	Schema []*ClickHouseSchema `json:"Schema,omitnil,omitempty" name:"Schema"`

	// 实例资源
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// ClickHouse的连接ip
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// ClickHouse的连接port
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// ClickHouse的用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// ClickHouse的密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 实例vip
	ServiceVip *string `json:"ServiceVip,omitnil,omitempty" name:"ServiceVip"`

	// 实例的vpcId
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 是否为自建集群
	SelfBuilt *bool `json:"SelfBuilt,omitnil,omitempty" name:"SelfBuilt"`

	// ClickHouse是否抛弃解析失败的消息，默认为true
	DropInvalidMessage *bool `json:"DropInvalidMessage,omitnil,omitempty" name:"DropInvalidMessage"`

	// ClickHouse 类型，emr-clickhouse : "emr";cdw-clickhouse : "cdwch";自建 : ""
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 当设置成员参数DropInvalidMessageToCls设置为true时,DropInvalidMessage参数失效
	DropCls *DropCls `json:"DropCls,omitnil,omitempty" name:"DropCls"`

	// 每批次投递到 ClickHouse 表消息数量，默认为 1000 条。
	// 提高该参数值，有利于减少往  ClickHouse 投递的次数，但在错误消息过多及网络不稳定等极端情况下时，可能导致频繁重试影响效率。
	BatchSize *int64 `json:"BatchSize,omitnil,omitempty" name:"BatchSize"`

	// 每次从 topic 中拉取消息大小，默认为 1MB，即至少要从 topic 中批量拉取 1MB 消息，才进行数据投递到 ClickHouse 操作。
	// 提高该参数值，有利于减少往  ClickHouse 投递的次数，但在错误消息过多及网络不稳定等极端情况下时，可能导致频繁重试影响效率。
	ConsumerFetchMinBytes *int64 `json:"ConsumerFetchMinBytes,omitnil,omitempty" name:"ConsumerFetchMinBytes"`

	// 每次从 topic 拉取消息最大等待时间，当超过当前最大等待时间时，即使没有拉取到 ConsumerFetchMinBytes 大小，也将进行 ClickHouse 投递操作。
	// 提高该参数值，有利于减少往  ClickHouse 投递的次数，但在错误消息过多及网络不稳定等极端情况下时，可能导致频繁重试影响效率。
	ConsumerFetchMaxWaitMs *int64 `json:"ConsumerFetchMaxWaitMs,omitnil,omitempty" name:"ConsumerFetchMaxWaitMs"`
}

type ClickHouseSchema struct {
	// 表的列名
	ColumnName *string `json:"ColumnName,omitnil,omitempty" name:"ColumnName"`

	// 该列对应的jsonKey名
	JsonKey *string `json:"JsonKey,omitnil,omitempty" name:"JsonKey"`

	// 表列项的类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 列项是否允许为空
	AllowNull *bool `json:"AllowNull,omitnil,omitempty" name:"AllowNull"`
}

type ClsParam struct {
	// 生产的信息是否为json格式
	DecodeJson *bool `json:"DecodeJson,omitnil,omitempty" name:"DecodeJson"`

	// cls日志主题id
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// cls日志集id
	LogSet *string `json:"LogSet,omitnil,omitempty" name:"LogSet"`

	// 当DecodeJson为false时必填
	ContentKey *string `json:"ContentKey,omitnil,omitempty" name:"ContentKey"`

	// 指定消息中的某字段内容作为cls日志的时间。
	// 字段内容格式需要是秒级时间戳
	TimeField *string `json:"TimeField,omitnil,omitempty" name:"TimeField"`
}

type ClusterInfo struct {
	// 集群Id
	ClusterId *int64 `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 集群最大磁盘 单位GB
	MaxDiskSize *int64 `json:"MaxDiskSize,omitnil,omitempty" name:"MaxDiskSize"`

	// 集群最大带宽 单位MB/s
	MaxBandWidth *int64 `json:"MaxBandWidth,omitnil,omitempty" name:"MaxBandWidth"`

	// 集群当前可用磁盘  单位GB
	AvailableDiskSize *int64 `json:"AvailableDiskSize,omitnil,omitempty" name:"AvailableDiskSize"`

	// 集群当前可用带宽 单位MB/s
	AvailableBandWidth *int64 `json:"AvailableBandWidth,omitnil,omitempty" name:"AvailableBandWidth"`

	// 集群所属可用区，表明集群归属的可用区
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 集群节点所在的可用区，若该集群为跨可用区集群，则包含该集群节点所在的多个可用区。
	ZoneIds []*int64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`
}

type Config struct {
	// 消息保留时间，单位ms。
	Retention *int64 `json:"Retention,omitnil,omitempty" name:"Retention"`

	// 最小同步复制数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinInsyncReplicas *int64 `json:"MinInsyncReplicas,omitnil,omitempty" name:"MinInsyncReplicas"`

	// 日志清理模式，默认 delete。
	// delete：日志按保存时间删除；compact：日志按 key 压缩；compact, delete：日志按 key 压缩且会保存时间删除。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CleanUpPolicy *string `json:"CleanUpPolicy,omitnil,omitempty" name:"CleanUpPolicy"`

	// Segment 分片滚动的时长，单位ms。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SegmentMs *int64 `json:"SegmentMs,omitnil,omitempty" name:"SegmentMs"`

	// 0表示 false。 1表示 true。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UncleanLeaderElectionEnable *int64 `json:"UncleanLeaderElectionEnable,omitnil,omitempty" name:"UncleanLeaderElectionEnable"`

	// Segment 分片滚动的字节数，单位bytes
	// 注意：此字段可能返回 null，表示取不到有效值。
	SegmentBytes *int64 `json:"SegmentBytes,omitnil,omitempty" name:"SegmentBytes"`

	// 最大消息字节数，单位bytes
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxMessageBytes *int64 `json:"MaxMessageBytes,omitnil,omitempty" name:"MaxMessageBytes"`

	// 消息保留文件大小，单位Bytes
	// 注意：此字段可能返回 null，表示取不到有效值。
	RetentionBytes *int64 `json:"RetentionBytes,omitnil,omitempty" name:"RetentionBytes"`

	// 消息保存的时间类型，CreateTime表示生产者创建这条消息的时间;LogAppendTime表示broker接收到消息的时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogMsgTimestampType *string `json:"LogMsgTimestampType,omitnil,omitempty" name:"LogMsgTimestampType"`
}

type ConnectResourceResourceIdResp struct {
	// 连接源的Id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

type Connection struct {
	// 主题名
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 消费组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 主题Id
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

type ConsumerGroup struct {
	// 用户组名称
	ConsumerGroupName *string `json:"ConsumerGroupName,omitnil,omitempty" name:"ConsumerGroupName"`

	// 订阅信息实体
	SubscribedInfo []*SubscribedInfo `json:"SubscribedInfo,omitnil,omitempty" name:"SubscribedInfo"`
}

type ConsumerGroupResponse struct {
	// 符合条件的消费组数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 主题列表
	TopicList []*ConsumerGroupTopic `json:"TopicList,omitnil,omitempty" name:"TopicList"`

	// 消费分组List
	GroupList []*ConsumerGroup `json:"GroupList,omitnil,omitempty" name:"GroupList"`

	// 所有分区数量
	TotalPartition *int64 `json:"TotalPartition,omitnil,omitempty" name:"TotalPartition"`

	// 监控的分区列表
	PartitionListForMonitor []*Partition `json:"PartitionListForMonitor,omitnil,omitempty" name:"PartitionListForMonitor"`

	// 主题总数
	TotalTopic *int64 `json:"TotalTopic,omitnil,omitempty" name:"TotalTopic"`

	// 监控的主题列表
	TopicListForMonitor []*ConsumerGroupTopic `json:"TopicListForMonitor,omitnil,omitempty" name:"TopicListForMonitor"`

	// 监控的组列表
	GroupListForMonitor []*Group `json:"GroupListForMonitor,omitnil,omitempty" name:"GroupListForMonitor"`
}

type ConsumerGroupSpeed struct {
	// 消费者组名称
	ConsumerGroupName *string `json:"ConsumerGroupName,omitnil,omitempty" name:"ConsumerGroupName"`

	// 消费速度 Count/Minute
	Speed *uint64 `json:"Speed,omitnil,omitempty" name:"Speed"`
}

type ConsumerGroupTopic struct {
	// 主题ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 主题名称
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

type ConsumerRecord struct {
	// 主题名
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 分区id
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// 位点
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 消息key
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 消息value
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 消息时间戳
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 消息headers
	// 注意：此字段可能返回 null，表示取不到有效值。
	Headers *string `json:"Headers,omitnil,omitempty" name:"Headers"`
}

type CosParam struct {
	// cos 存储桶名称
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// 地域代码
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 对象名称
	ObjectKey *string `json:"ObjectKey,omitnil,omitempty" name:"ObjectKey"`

	// 汇聚消息量的大小（单位：MB)
	AggregateBatchSize *uint64 `json:"AggregateBatchSize,omitnil,omitempty" name:"AggregateBatchSize"`

	// 汇聚的时间间隔（单位：小时）
	AggregateInterval *uint64 `json:"AggregateInterval,omitnil,omitempty" name:"AggregateInterval"`

	// 消息汇聚后的文件格式（支持csv, json）
	FormatOutputType *string `json:"FormatOutputType,omitnil,omitempty" name:"FormatOutputType"`

	// 转储的对象目录前缀
	ObjectKeyPrefix *string `json:"ObjectKeyPrefix,omitnil,omitempty" name:"ObjectKeyPrefix"`

	// 根据strptime 时间格式化的分区格式
	DirectoryTimeFormat *string `json:"DirectoryTimeFormat,omitnil,omitempty" name:"DirectoryTimeFormat"`
}

// Predefined struct for user
type CreateAclRequestParams struct {
	// ckafka集群实例Id，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Acl资源类型，(2:TOPIC，3:GROUP，4:CLUSTER)
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// Acl操作方式，(2:ALL，3:READ，4:WRITE，5:CREATE，6:DELETE，7:ALTER，8:DESCRIBE，9:CLUSTER_ACTION，10:DESCRIBE_CONFIGS，11:ALTER_CONFIGS，12:IDEMPOTENT_WRITE)
	Operation *int64 `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 权限类型，(2:DENY，3:ALLOW)，当前ckafka支持ALLOW(相当于白名单)，其它用于后续兼容开源kafka的acl时使用
	PermissionType *int64 `json:"PermissionType,omitnil,omitempty" name:"PermissionType"`

	// 资源名称，和resourceType相关，如当resourceType为TOPIC时，则该字段表示topic名称，当resourceType为GROUP时，该字段表示group名称，当resourceType为CLUSTER时，该字段可为空。
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 默认为*，表示任何host都可以访问。支持填写IP或网段，支持“;”分隔。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 用户列表，默认为User:*，表示任何user都可以访问，当前用户只能是用户列表中包含的用户。传入时需要加 User: 前缀,如用户A则传入User:A。
	Principal *string `json:"Principal,omitnil,omitempty" name:"Principal"`

	// 资源名称列表,Json字符串格式。ResourceName和resourceNameList只能指定其中一个。
	ResourceNameList *string `json:"ResourceNameList,omitnil,omitempty" name:"ResourceNameList"`
}

type CreateAclRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Acl资源类型，(2:TOPIC，3:GROUP，4:CLUSTER)
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// Acl操作方式，(2:ALL，3:READ，4:WRITE，5:CREATE，6:DELETE，7:ALTER，8:DESCRIBE，9:CLUSTER_ACTION，10:DESCRIBE_CONFIGS，11:ALTER_CONFIGS，12:IDEMPOTENT_WRITE)
	Operation *int64 `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 权限类型，(2:DENY，3:ALLOW)，当前ckafka支持ALLOW(相当于白名单)，其它用于后续兼容开源kafka的acl时使用
	PermissionType *int64 `json:"PermissionType,omitnil,omitempty" name:"PermissionType"`

	// 资源名称，和resourceType相关，如当resourceType为TOPIC时，则该字段表示topic名称，当resourceType为GROUP时，该字段表示group名称，当resourceType为CLUSTER时，该字段可为空。
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 默认为*，表示任何host都可以访问。支持填写IP或网段，支持“;”分隔。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 用户列表，默认为User:*，表示任何user都可以访问，当前用户只能是用户列表中包含的用户。传入时需要加 User: 前缀,如用户A则传入User:A。
	Principal *string `json:"Principal,omitnil,omitempty" name:"Principal"`

	// 资源名称列表,Json字符串格式。ResourceName和resourceNameList只能指定其中一个。
	ResourceNameList *string `json:"ResourceNameList,omitnil,omitempty" name:"ResourceNameList"`
}

func (r *CreateAclRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAclRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ResourceType")
	delete(f, "Operation")
	delete(f, "PermissionType")
	delete(f, "ResourceName")
	delete(f, "Host")
	delete(f, "Principal")
	delete(f, "ResourceNameList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAclRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAclResponseParams struct {
	// 返回结果
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAclResponse struct {
	*tchttp.BaseResponse
	Response *CreateAclResponseParams `json:"Response"`
}

func (r *CreateAclResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAclResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAclRuleRequestParams struct {
	// ckafka集群实例Id，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Acl资源类型,目前只支持Topic,枚举值列表：Topic
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// ACL规则匹配类型，目前支持前缀匹配与预设策略，枚举值列表：PREFIXED/PRESET
	PatternType *string `json:"PatternType,omitnil,omitempty" name:"PatternType"`

	// 规则名称
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 设置的ACL规则列表
	RuleList []*AclRuleInfo `json:"RuleList,omitnil,omitempty" name:"RuleList"`

	// 表示前缀匹配的前缀的值 (当PatternType取值为PREFIXED时，此参数必填)
	Pattern *string `json:"Pattern,omitnil,omitempty" name:"Pattern"`

	// 预设ACL规则是否应用到新增的topic中。默认为0，表示否。取值为1时表示是。
	IsApplied *int64 `json:"IsApplied,omitnil,omitempty" name:"IsApplied"`

	// ACL规则的备注
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

type CreateAclRuleRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Acl资源类型,目前只支持Topic,枚举值列表：Topic
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// ACL规则匹配类型，目前支持前缀匹配与预设策略，枚举值列表：PREFIXED/PRESET
	PatternType *string `json:"PatternType,omitnil,omitempty" name:"PatternType"`

	// 规则名称
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 设置的ACL规则列表
	RuleList []*AclRuleInfo `json:"RuleList,omitnil,omitempty" name:"RuleList"`

	// 表示前缀匹配的前缀的值 (当PatternType取值为PREFIXED时，此参数必填)
	Pattern *string `json:"Pattern,omitnil,omitempty" name:"Pattern"`

	// 预设ACL规则是否应用到新增的topic中。默认为0，表示否。取值为1时表示是。
	IsApplied *int64 `json:"IsApplied,omitnil,omitempty" name:"IsApplied"`

	// ACL规则的备注
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

func (r *CreateAclRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAclRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ResourceType")
	delete(f, "PatternType")
	delete(f, "RuleName")
	delete(f, "RuleList")
	delete(f, "Pattern")
	delete(f, "IsApplied")
	delete(f, "Comment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAclRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAclRuleResponseParams struct {
	// 规则的唯一表示Key
	Result *int64 `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAclRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateAclRuleResponseParams `json:"Response"`
}

func (r *CreateAclRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAclRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCdcClusterRequestParams struct {
	// cdc的id
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`

	// vpcId,一个地域只有唯一一个vpcid用于CDC
	CdcVpcId *string `json:"CdcVpcId,omitnil,omitempty" name:"CdcVpcId"`

	// 每个CDC集群有唯一一个子网ID
	CdcSubnetId *string `json:"CdcSubnetId,omitnil,omitempty" name:"CdcSubnetId"`

	// 所在可用区ID
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 实例带宽,单位MB/s; 最小值:20MB/s, 高级版最大值:360MB/s,专业版最大值:100000MB/s  标准版固定带宽规格: 40MB/s, 100MB/s, 150MB/s
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// cdc集群的总磁盘
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// ckafka集群实例磁盘类型
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 系统盘类型
	SystemDiskType *string `json:"SystemDiskType,omitnil,omitempty" name:"SystemDiskType"`
}

type CreateCdcClusterRequest struct {
	*tchttp.BaseRequest
	
	// cdc的id
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`

	// vpcId,一个地域只有唯一一个vpcid用于CDC
	CdcVpcId *string `json:"CdcVpcId,omitnil,omitempty" name:"CdcVpcId"`

	// 每个CDC集群有唯一一个子网ID
	CdcSubnetId *string `json:"CdcSubnetId,omitnil,omitempty" name:"CdcSubnetId"`

	// 所在可用区ID
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 实例带宽,单位MB/s; 最小值:20MB/s, 高级版最大值:360MB/s,专业版最大值:100000MB/s  标准版固定带宽规格: 40MB/s, 100MB/s, 150MB/s
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// cdc集群的总磁盘
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// ckafka集群实例磁盘类型
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 系统盘类型
	SystemDiskType *string `json:"SystemDiskType,omitnil,omitempty" name:"SystemDiskType"`
}

func (r *CreateCdcClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCdcClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CdcId")
	delete(f, "CdcVpcId")
	delete(f, "CdcSubnetId")
	delete(f, "ZoneId")
	delete(f, "Bandwidth")
	delete(f, "DiskSize")
	delete(f, "DiskType")
	delete(f, "SystemDiskType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCdcClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCdcClusterResponseParams struct {
	// 无
	Result *CdcClusterResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCdcClusterResponse struct {
	*tchttp.BaseResponse
	Response *CreateCdcClusterResponseParams `json:"Response"`
}

func (r *CreateCdcClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCdcClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConnectResourceRequestParams struct {
	// 连接源名称
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 连接源类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 连接源描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Dts配置，Type为DTS时必填
	DtsConnectParam *DtsConnectParam `json:"DtsConnectParam,omitnil,omitempty" name:"DtsConnectParam"`

	// MongoDB配置，Type为MONGODB时必填
	MongoDBConnectParam *MongoDBConnectParam `json:"MongoDBConnectParam,omitnil,omitempty" name:"MongoDBConnectParam"`

	// Es配置，Type为ES时必填
	EsConnectParam *EsConnectParam `json:"EsConnectParam,omitnil,omitempty" name:"EsConnectParam"`

	// ClickHouse配置，Type为CLICKHOUSE时必填
	ClickHouseConnectParam *ClickHouseConnectParam `json:"ClickHouseConnectParam,omitnil,omitempty" name:"ClickHouseConnectParam"`

	// MySQL配置，Type为MYSQL或TDSQL_C_MYSQL时必填
	MySQLConnectParam *MySQLConnectParam `json:"MySQLConnectParam,omitnil,omitempty" name:"MySQLConnectParam"`

	// PostgreSQL配置，Type为POSTGRESQL或TDSQL_C_POSTGRESQL时必填
	PostgreSQLConnectParam *PostgreSQLConnectParam `json:"PostgreSQLConnectParam,omitnil,omitempty" name:"PostgreSQLConnectParam"`

	// MariaDB配置，Type为MARIADB时必填
	MariaDBConnectParam *MariaDBConnectParam `json:"MariaDBConnectParam,omitnil,omitempty" name:"MariaDBConnectParam"`

	// SQLServer配置，Type为SQLSERVER时必填
	SQLServerConnectParam *SQLServerConnectParam `json:"SQLServerConnectParam,omitnil,omitempty" name:"SQLServerConnectParam"`

	// Doris 配置，Type为 DORIS 时必填
	DorisConnectParam *DorisConnectParam `json:"DorisConnectParam,omitnil,omitempty" name:"DorisConnectParam"`

	// Kafka配置，Type为 KAFKA 时必填
	KafkaConnectParam *KafkaConnectParam `json:"KafkaConnectParam,omitnil,omitempty" name:"KafkaConnectParam"`

	// MQTT配置，Type为 MQTT 时必填
	MqttConnectParam *MqttConnectParam `json:"MqttConnectParam,omitnil,omitempty" name:"MqttConnectParam"`
}

type CreateConnectResourceRequest struct {
	*tchttp.BaseRequest
	
	// 连接源名称
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 连接源类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 连接源描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Dts配置，Type为DTS时必填
	DtsConnectParam *DtsConnectParam `json:"DtsConnectParam,omitnil,omitempty" name:"DtsConnectParam"`

	// MongoDB配置，Type为MONGODB时必填
	MongoDBConnectParam *MongoDBConnectParam `json:"MongoDBConnectParam,omitnil,omitempty" name:"MongoDBConnectParam"`

	// Es配置，Type为ES时必填
	EsConnectParam *EsConnectParam `json:"EsConnectParam,omitnil,omitempty" name:"EsConnectParam"`

	// ClickHouse配置，Type为CLICKHOUSE时必填
	ClickHouseConnectParam *ClickHouseConnectParam `json:"ClickHouseConnectParam,omitnil,omitempty" name:"ClickHouseConnectParam"`

	// MySQL配置，Type为MYSQL或TDSQL_C_MYSQL时必填
	MySQLConnectParam *MySQLConnectParam `json:"MySQLConnectParam,omitnil,omitempty" name:"MySQLConnectParam"`

	// PostgreSQL配置，Type为POSTGRESQL或TDSQL_C_POSTGRESQL时必填
	PostgreSQLConnectParam *PostgreSQLConnectParam `json:"PostgreSQLConnectParam,omitnil,omitempty" name:"PostgreSQLConnectParam"`

	// MariaDB配置，Type为MARIADB时必填
	MariaDBConnectParam *MariaDBConnectParam `json:"MariaDBConnectParam,omitnil,omitempty" name:"MariaDBConnectParam"`

	// SQLServer配置，Type为SQLSERVER时必填
	SQLServerConnectParam *SQLServerConnectParam `json:"SQLServerConnectParam,omitnil,omitempty" name:"SQLServerConnectParam"`

	// Doris 配置，Type为 DORIS 时必填
	DorisConnectParam *DorisConnectParam `json:"DorisConnectParam,omitnil,omitempty" name:"DorisConnectParam"`

	// Kafka配置，Type为 KAFKA 时必填
	KafkaConnectParam *KafkaConnectParam `json:"KafkaConnectParam,omitnil,omitempty" name:"KafkaConnectParam"`

	// MQTT配置，Type为 MQTT 时必填
	MqttConnectParam *MqttConnectParam `json:"MqttConnectParam,omitnil,omitempty" name:"MqttConnectParam"`
}

func (r *CreateConnectResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConnectResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceName")
	delete(f, "Type")
	delete(f, "Description")
	delete(f, "DtsConnectParam")
	delete(f, "MongoDBConnectParam")
	delete(f, "EsConnectParam")
	delete(f, "ClickHouseConnectParam")
	delete(f, "MySQLConnectParam")
	delete(f, "PostgreSQLConnectParam")
	delete(f, "MariaDBConnectParam")
	delete(f, "SQLServerConnectParam")
	delete(f, "DorisConnectParam")
	delete(f, "KafkaConnectParam")
	delete(f, "MqttConnectParam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateConnectResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConnectResourceResponseParams struct {
	// 连接源的Id
	Result *ConnectResourceResourceIdResp `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateConnectResourceResponse struct {
	*tchttp.BaseResponse
	Response *CreateConnectResourceResponseParams `json:"Response"`
}

func (r *CreateConnectResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConnectResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConsumerRequestParams struct {
	// ckafka集群实例Id，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 消费分组名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 主题名，TopicName、TopicNameList 需要显示指定一个存在的主题名
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 主题名列表
	TopicNameList []*string `json:"TopicNameList,omitnil,omitempty" name:"TopicNameList"`
}

type CreateConsumerRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 消费分组名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 主题名，TopicName、TopicNameList 需要显示指定一个存在的主题名
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 主题名列表
	TopicNameList []*string `json:"TopicNameList,omitnil,omitempty" name:"TopicNameList"`
}

func (r *CreateConsumerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConsumerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "GroupName")
	delete(f, "TopicName")
	delete(f, "TopicNameList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateConsumerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConsumerResponseParams struct {
	// 创建消费者组返回结果
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateConsumerResponse struct {
	*tchttp.BaseResponse
	Response *CreateConsumerResponseParams `json:"Response"`
}

func (r *CreateConsumerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConsumerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDatahubTaskRequestParams struct {
	// 任务名称,只能以字母起始,允许包含字母、数字、- 、.  、 下划线且长度不超过64 (、为分割符号规则不包含)
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 任务类型，SOURCE数据接入，SINK数据流出
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 数据源
	SourceResource *DatahubResource `json:"SourceResource,omitnil,omitempty" name:"SourceResource"`

	// 数据目标
	TargetResource *DatahubResource `json:"TargetResource,omitnil,omitempty" name:"TargetResource"`

	// 数据处理规则
	TransformParam *TransformParam `json:"TransformParam,omitnil,omitempty" name:"TransformParam"`

	// 实例连接参数【已废弃】
	//
	// Deprecated: PrivateLinkParam is deprecated.
	PrivateLinkParam *PrivateLinkParam `json:"PrivateLinkParam,omitnil,omitempty" name:"PrivateLinkParam"`

	// 选择所要绑定的SchemaId
	SchemaId *string `json:"SchemaId,omitnil,omitempty" name:"SchemaId"`

	// 数据处理规则
	TransformsParam *TransformsParam `json:"TransformsParam,omitnil,omitempty" name:"TransformsParam"`

	// 任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 任务描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateDatahubTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务名称,只能以字母起始,允许包含字母、数字、- 、.  、 下划线且长度不超过64 (、为分割符号规则不包含)
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 任务类型，SOURCE数据接入，SINK数据流出
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 数据源
	SourceResource *DatahubResource `json:"SourceResource,omitnil,omitempty" name:"SourceResource"`

	// 数据目标
	TargetResource *DatahubResource `json:"TargetResource,omitnil,omitempty" name:"TargetResource"`

	// 数据处理规则
	TransformParam *TransformParam `json:"TransformParam,omitnil,omitempty" name:"TransformParam"`

	// 实例连接参数【已废弃】
	PrivateLinkParam *PrivateLinkParam `json:"PrivateLinkParam,omitnil,omitempty" name:"PrivateLinkParam"`

	// 选择所要绑定的SchemaId
	SchemaId *string `json:"SchemaId,omitnil,omitempty" name:"SchemaId"`

	// 数据处理规则
	TransformsParam *TransformsParam `json:"TransformsParam,omitnil,omitempty" name:"TransformsParam"`

	// 任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 任务描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateDatahubTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDatahubTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskName")
	delete(f, "TaskType")
	delete(f, "SourceResource")
	delete(f, "TargetResource")
	delete(f, "TransformParam")
	delete(f, "PrivateLinkParam")
	delete(f, "SchemaId")
	delete(f, "TransformsParam")
	delete(f, "TaskId")
	delete(f, "Tags")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDatahubTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDatahubTaskRes struct {
	// 转储任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 数据转储Id
	DatahubId *string `json:"DatahubId,omitnil,omitempty" name:"DatahubId"`
}

// Predefined struct for user
type CreateDatahubTaskResponseParams struct {
	// 返回结果
	Result *CreateDatahubTaskRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDatahubTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateDatahubTaskResponseParams `json:"Response"`
}

func (r *CreateDatahubTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDatahubTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDatahubTopicRequestParams struct {
	// 名称，是一个不超过 128 个字符的字符串，必须以“AppId-”为首字符，剩余部分可以包含字母、数字和横划线(-)
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Partition个数，大于0
	PartitionNum *int64 `json:"PartitionNum,omitnil,omitempty" name:"PartitionNum"`

	// 消息保留时间，单位ms，当前最小值为60000ms
	RetentionMs *int64 `json:"RetentionMs,omitnil,omitempty" name:"RetentionMs"`

	// 主题备注，是一个不超过 64 个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateDatahubTopicRequest struct {
	*tchttp.BaseRequest
	
	// 名称，是一个不超过 128 个字符的字符串，必须以“AppId-”为首字符，剩余部分可以包含字母、数字和横划线(-)
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Partition个数，大于0
	PartitionNum *int64 `json:"PartitionNum,omitnil,omitempty" name:"PartitionNum"`

	// 消息保留时间，单位ms，当前最小值为60000ms
	RetentionMs *int64 `json:"RetentionMs,omitnil,omitempty" name:"RetentionMs"`

	// 主题备注，是一个不超过 64 个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *CreateDatahubTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDatahubTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "PartitionNum")
	delete(f, "RetentionMs")
	delete(f, "Note")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDatahubTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDatahubTopicResponseParams struct {
	// 返回创建结果
	Result *DatahubTopicResp `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDatahubTopicResponse struct {
	*tchttp.BaseResponse
	Response *CreateDatahubTopicResponseParams `json:"Response"`
}

func (r *CreateDatahubTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDatahubTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstancePostData struct {
	// CreateInstancePre返回固定为0，不能作为CheckTaskStatus的查询条件。只是为了保证和后台数据结构对齐。
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 订单号列表
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// ckafka集群实例Id，当购买多个实例时，默认返回购买的第一个实例 id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 订单和购买实例对应映射列表
	DealNameInstanceIdMapping []*DealInstanceDTO `json:"DealNameInstanceIdMapping,omitnil,omitempty" name:"DealNameInstanceIdMapping"`
}

type CreateInstancePostResp struct {
	// 返回的code，0为正常，非0为错误
	ReturnCode *string `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// 接口返回消息，当接口报错时提示错误信息
	ReturnMessage *string `json:"ReturnMessage,omitnil,omitempty" name:"ReturnMessage"`

	// 返回的Data数据
	Data *CreateInstancePostData `json:"Data,omitnil,omitempty" name:"Data"`
}

type CreateInstancePreData struct {
	// CreateInstancePre返回固定为0，不能作为CheckTaskStatus的查询条件。只是为了保证和后台数据结构对齐。
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 订单号列表
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// ckafka集群实例Id，当购买多个实例时，默认返回购买的第一个实例 id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 订单和购买实例对应映射列表
	DealNameInstanceIdMapping []*DealInstanceDTO `json:"DealNameInstanceIdMapping,omitnil,omitempty" name:"DealNameInstanceIdMapping"`
}

// Predefined struct for user
type CreateInstancePreRequestParams struct {
	// <p>ckafka集群实例Name，是一个不超过 128 个字符的任意字符串。</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>可用区。当购买多可用区实例时，当前参数为主可用区。  <a href="https://cloud.tencent.com/document/product/597/55246">查看可用区</a></p>
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>预付费购买时长，例如 "1m",就是一个月,取值范围 1m~36m</p>
	Period *string `json:"Period,omitnil,omitempty" name:"Period"`

	// <p>国际站标准版实例规格。目前只有国际站标准版使用当前字段区分规格，国内站标准版使用峰值带宽区分规格。除了国际站标准版外的所有实例填写 1 即可。国际站标准版实例：入门型(general)]填写1；[标准型(standard)]填写2；[进阶型(advanced)]填写3；[容量型(capacity)]填写4；[高阶型1(specialized-1)]填写5；[高阶型2(specialized-2)]填写6；[高阶型3(specialized-3)]填写7；[高阶型4(specialized-4)]填写8。</p>
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// <p>私有网络Id</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>子网id</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>可选。实例日志的最长保留时间，单位分钟，不填默认为1440（1天），可设置范围为1分钟到90天。</p>
	MsgRetentionTime *int64 `json:"MsgRetentionTime,omitnil,omitempty" name:"MsgRetentionTime"`

	// <p>创建实例时可以选择集群Id, 该入参表示集群Id</p>
	ClusterId *int64 `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>预付费自动续费标记，0表示默认状态(用户未设置，即初始状态)， 1表示自动续费，2表示明确不自动续费(用户设置)</p>
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// <p>CKafka版本号[2.4.1, 2.4.2, 2.8.1, 3.2.3], 默认取值是2.4.1。2.4.1 与 2.4.2 属于同一个版本，传任意一个均可。</p>
	KafkaVersion *string `json:"KafkaVersion,omitnil,omitempty" name:"KafkaVersion"`

	// <p>实例类型: [标准版实例]填写 "standard" (默认), [专业版实例]填写 "profession",[高级版实例]填写"premium"</p>
	SpecificationsType *string `json:"SpecificationsType,omitnil,omitempty" name:"SpecificationsType"`

	// <p>磁盘大小，如果跟控制台规格配比不相符，则无法创建成功。默认取值为500，步长设置为100。可以通过以下链接查看计费规格：https://cloud.tencent.com/document/product/597/122562</p>
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// <p>实例带宽,默认值为40，单位MB/s; 最小值:20MB/s, 高级版最大值:360MB/s,专业版最大值:100000MB/s  标准版固定带宽规格: 40MB/s, 100MB/s, 150MB/s。可以通过以下链接查看计费规格：https://cloud.tencent.com/document/product/597/11745</p>
	BandWidth *int64 `json:"BandWidth,omitnil,omitempty" name:"BandWidth"`

	// <p>分区大小，如果跟控制台规格配比不相符，则无法创建成功。默认值为800，步长为100。可以通过以下链接查看计费规格：https://cloud.tencent.com/document/product/597/122563</p>
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// <p>标签</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>专业版/高级版实例磁盘类型，标准版实例不需要填写。"CLOUD_SSD"：SSD云硬盘；"CLOUD_BASIC"：高性能云硬盘。不传默认为 "CLOUD_BASIC"</p>
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// <p>是否创建跨可用区实例，当前参数为 true 时，zoneIds必填</p>
	MultiZoneFlag *bool `json:"MultiZoneFlag,omitnil,omitempty" name:"MultiZoneFlag"`

	// <p>可用区列表，购买多可用区实例时为必填项</p>
	ZoneIds []*int64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// <p>公网带宽大小，单位 Mbps。默认是没有加上免费 3Mbps 带宽。例如总共需要 3Mbps 公网带宽，此处传 0；总共需要 6Mbps 公网带宽，此处传 3。默认值为 0。需要保证传入参数为 3 的整数倍</p>
	PublicNetworkMonthly *int64 `json:"PublicNetworkMonthly,omitnil,omitempty" name:"PublicNetworkMonthly"`

	// <p>购买实例数量。非必填，默认值为 1。当传入该参数时，会创建多个 instanceName 加后缀区分的实例</p>
	InstanceNum *int64 `json:"InstanceNum,omitnil,omitempty" name:"InstanceNum"`

	// <p>是否自动选择代金券:1-是;0否。默认为0</p>
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// <p>弹性带宽开关 0不开启  1开启（0默认）</p>
	ElasticBandwidthSwitch *int64 `json:"ElasticBandwidthSwitch,omitnil,omitempty" name:"ElasticBandwidthSwitch"`

	// <p>自定义证书Id,仅当SpecificationsType为profession时生效,支持自定义证书能力</p><p>可通过<a href="https://cloud.tencent.com/document/product/400/41673">DescribeCertificateDetail</a>接口获取</p>
	CustomSSLCertId *string `json:"CustomSSLCertId,omitnil,omitempty" name:"CustomSSLCertId"`
}

type CreateInstancePreRequest struct {
	*tchttp.BaseRequest
	
	// <p>ckafka集群实例Name，是一个不超过 128 个字符的任意字符串。</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>可用区。当购买多可用区实例时，当前参数为主可用区。  <a href="https://cloud.tencent.com/document/product/597/55246">查看可用区</a></p>
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>预付费购买时长，例如 "1m",就是一个月,取值范围 1m~36m</p>
	Period *string `json:"Period,omitnil,omitempty" name:"Period"`

	// <p>国际站标准版实例规格。目前只有国际站标准版使用当前字段区分规格，国内站标准版使用峰值带宽区分规格。除了国际站标准版外的所有实例填写 1 即可。国际站标准版实例：入门型(general)]填写1；[标准型(standard)]填写2；[进阶型(advanced)]填写3；[容量型(capacity)]填写4；[高阶型1(specialized-1)]填写5；[高阶型2(specialized-2)]填写6；[高阶型3(specialized-3)]填写7；[高阶型4(specialized-4)]填写8。</p>
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// <p>私有网络Id</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>子网id</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>可选。实例日志的最长保留时间，单位分钟，不填默认为1440（1天），可设置范围为1分钟到90天。</p>
	MsgRetentionTime *int64 `json:"MsgRetentionTime,omitnil,omitempty" name:"MsgRetentionTime"`

	// <p>创建实例时可以选择集群Id, 该入参表示集群Id</p>
	ClusterId *int64 `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>预付费自动续费标记，0表示默认状态(用户未设置，即初始状态)， 1表示自动续费，2表示明确不自动续费(用户设置)</p>
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// <p>CKafka版本号[2.4.1, 2.4.2, 2.8.1, 3.2.3], 默认取值是2.4.1。2.4.1 与 2.4.2 属于同一个版本，传任意一个均可。</p>
	KafkaVersion *string `json:"KafkaVersion,omitnil,omitempty" name:"KafkaVersion"`

	// <p>实例类型: [标准版实例]填写 "standard" (默认), [专业版实例]填写 "profession",[高级版实例]填写"premium"</p>
	SpecificationsType *string `json:"SpecificationsType,omitnil,omitempty" name:"SpecificationsType"`

	// <p>磁盘大小，如果跟控制台规格配比不相符，则无法创建成功。默认取值为500，步长设置为100。可以通过以下链接查看计费规格：https://cloud.tencent.com/document/product/597/122562</p>
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// <p>实例带宽,默认值为40，单位MB/s; 最小值:20MB/s, 高级版最大值:360MB/s,专业版最大值:100000MB/s  标准版固定带宽规格: 40MB/s, 100MB/s, 150MB/s。可以通过以下链接查看计费规格：https://cloud.tencent.com/document/product/597/11745</p>
	BandWidth *int64 `json:"BandWidth,omitnil,omitempty" name:"BandWidth"`

	// <p>分区大小，如果跟控制台规格配比不相符，则无法创建成功。默认值为800，步长为100。可以通过以下链接查看计费规格：https://cloud.tencent.com/document/product/597/122563</p>
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// <p>标签</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>专业版/高级版实例磁盘类型，标准版实例不需要填写。"CLOUD_SSD"：SSD云硬盘；"CLOUD_BASIC"：高性能云硬盘。不传默认为 "CLOUD_BASIC"</p>
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// <p>是否创建跨可用区实例，当前参数为 true 时，zoneIds必填</p>
	MultiZoneFlag *bool `json:"MultiZoneFlag,omitnil,omitempty" name:"MultiZoneFlag"`

	// <p>可用区列表，购买多可用区实例时为必填项</p>
	ZoneIds []*int64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// <p>公网带宽大小，单位 Mbps。默认是没有加上免费 3Mbps 带宽。例如总共需要 3Mbps 公网带宽，此处传 0；总共需要 6Mbps 公网带宽，此处传 3。默认值为 0。需要保证传入参数为 3 的整数倍</p>
	PublicNetworkMonthly *int64 `json:"PublicNetworkMonthly,omitnil,omitempty" name:"PublicNetworkMonthly"`

	// <p>购买实例数量。非必填，默认值为 1。当传入该参数时，会创建多个 instanceName 加后缀区分的实例</p>
	InstanceNum *int64 `json:"InstanceNum,omitnil,omitempty" name:"InstanceNum"`

	// <p>是否自动选择代金券:1-是;0否。默认为0</p>
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// <p>弹性带宽开关 0不开启  1开启（0默认）</p>
	ElasticBandwidthSwitch *int64 `json:"ElasticBandwidthSwitch,omitnil,omitempty" name:"ElasticBandwidthSwitch"`

	// <p>自定义证书Id,仅当SpecificationsType为profession时生效,支持自定义证书能力</p><p>可通过<a href="https://cloud.tencent.com/document/product/400/41673">DescribeCertificateDetail</a>接口获取</p>
	CustomSSLCertId *string `json:"CustomSSLCertId,omitnil,omitempty" name:"CustomSSLCertId"`
}

func (r *CreateInstancePreRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstancePreRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceName")
	delete(f, "ZoneId")
	delete(f, "Period")
	delete(f, "InstanceType")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "MsgRetentionTime")
	delete(f, "ClusterId")
	delete(f, "RenewFlag")
	delete(f, "KafkaVersion")
	delete(f, "SpecificationsType")
	delete(f, "DiskSize")
	delete(f, "BandWidth")
	delete(f, "Partition")
	delete(f, "Tags")
	delete(f, "DiskType")
	delete(f, "MultiZoneFlag")
	delete(f, "ZoneIds")
	delete(f, "PublicNetworkMonthly")
	delete(f, "InstanceNum")
	delete(f, "AutoVoucher")
	delete(f, "ElasticBandwidthSwitch")
	delete(f, "CustomSSLCertId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstancePreRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstancePreResp struct {
	// 返回的code，0为正常，非0为错误
	ReturnCode *string `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// 成功消息
	ReturnMessage *string `json:"ReturnMessage,omitnil,omitempty" name:"ReturnMessage"`

	// 操作型返回的Data数据
	Data *CreateInstancePreData `json:"Data,omitnil,omitempty" name:"Data"`

	// 删除时间。目前该参数字段已废弃，将会在未来被删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: DeleteRouteTimestamp is deprecated.
	DeleteRouteTimestamp *string `json:"DeleteRouteTimestamp,omitnil,omitempty" name:"DeleteRouteTimestamp"`
}

// Predefined struct for user
type CreateInstancePreResponseParams struct {
	// <p>返回结果</p>
	Result *CreateInstancePreResp `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateInstancePreResponse struct {
	*tchttp.BaseResponse
	Response *CreateInstancePreResponseParams `json:"Response"`
}

func (r *CreateInstancePreResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstancePreResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePartitionRequestParams struct {
	// ckafka集群实例Id，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名称，可通过[DescribeTopic](https://cloud.tencent.com/document/product/597/40847)接口获取
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 主题分区个数，传入参数为修改后的分区数，而不是增加的分区数，因此传入参数需要大于当前主题分区个数。
	PartitionNum *int64 `json:"PartitionNum,omitnil,omitempty" name:"PartitionNum"`
}

type CreatePartitionRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名称，可通过[DescribeTopic](https://cloud.tencent.com/document/product/597/40847)接口获取
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 主题分区个数，传入参数为修改后的分区数，而不是增加的分区数，因此传入参数需要大于当前主题分区个数。
	PartitionNum *int64 `json:"PartitionNum,omitnil,omitempty" name:"PartitionNum"`
}

func (r *CreatePartitionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePartitionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "TopicName")
	delete(f, "PartitionNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePartitionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePartitionResponseParams struct {
	// 返回的结果集
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePartitionResponse struct {
	*tchttp.BaseResponse
	Response *CreatePartitionResponseParams `json:"Response"`
}

func (r *CreatePartitionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePartitionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePostPaidInstanceRequestParams struct {
	// <p>私有网络Id,可通过<a href="https://cloud.tencent.com/document/product/215/15778">DescribeVpcs</a>接口获取</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>子网Id,可通过<a href="https://cloud.tencent.com/document/product/215/15784">DescribeSubnets</a>接口获取</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>ckafka集群实例名称，是一个长度不超过128的任意字符。</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>国际站标准版实例规格。目前只有国际站标准版使用当前字段区分规格，国内站标准版使用峰值带宽区分规格。除了国际站标准版外的所有实例填写 1 即可。国际站标准版实例：入门型(general)]填写1；[标准型(standard)]填写2；[进阶型(advanced)]填写3；[容量型(capacity)]填写4；[高阶型1(specialized-1)]填写5；[高阶型2(specialized-2)]填写6；[高阶型3(specialized-3)]填写7；[高阶型4(specialized-4)]填写8。</p>
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// <p>实例日志的默认最长保留时间，单位分钟。不传入该参数时默认为 1440 分钟（1天），最大30天。当 topic 显式设置消息保留时间时，以 topic 保留时间为准</p>
	MsgRetentionTime *int64 `json:"MsgRetentionTime,omitnil,omitempty" name:"MsgRetentionTime"`

	// <p>创建实例时可以选择集群Id, 该入参表示集群Id。不指定实例所在集群则不传入该参数</p>
	ClusterId *int64 `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>实例版本。目前支持当前支持"2.4.1", "2.4.2","2.8.1", "3.2.3"，默认取值"2.4.1"。"2.4.1" 与 "2.4.2" 属于同一个版本，传任意一个均可。</p>
	KafkaVersion *string `json:"KafkaVersion,omitnil,omitempty" name:"KafkaVersion"`

	// <p>实例类型。"standard"：标准版，"profession"：专业版。  (标准版仅国际站支持，国内站目前支持专业版)</p>
	SpecificationsType *string `json:"SpecificationsType,omitnil,omitempty" name:"SpecificationsType"`

	// <p>专业版实例磁盘类型，标准版实例不需要填写。"CLOUD_SSD"：SSD云硬盘；"CLOUD_BASIC"：高性能云硬盘。不传默认值为 "CLOUD_BASIC"</p>
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// <p>实例内网峰值带宽，默认值为40。单位 MB/s。标准版需传入当前实例规格所对应的峰值带宽。注意如果创建的实例为专业版实例，峰值带宽，分区数等参数配置需要满足专业版的计费规格，可以通过以下链接查看计费规格：https://cloud.tencent.com/document/product/597/11745</p>
	BandWidth *int64 `json:"BandWidth,omitnil,omitempty" name:"BandWidth"`

	// <p>实例硬盘大小，默认取值为500，步长设置为100。需要满足当前实例的计费规格，可以通过以下链接查看计费规格：https://cloud.tencent.com/document/product/597/122562</p>
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// <p>实例最大分区数量，需要满足当前实例的计费规格。默认值为800，步长为100。可以通过以下链接查看计费规格：https://cloud.tencent.com/document/product/597/122563</p>
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// <p>实例最大 topic 数量，需要满足当前实例的计费规格。默认值为800，步长设置为100。</p>
	TopicNum *int64 `json:"TopicNum,omitnil,omitempty" name:"TopicNum"`

	// <p>实例所在的可用区。当创建多可用区实例时，该参数为创建的默认接入点所在子网的可用区 id。ZoneId、ZoneIds不能同时为空，可通过<a href="https://cloud.tencent.com/document/product/597/55246">DescribeCkafkaZone</a>接口获取。</p>
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>当前实例是否为多可用区实例。</p>
	MultiZoneFlag *bool `json:"MultiZoneFlag,omitnil,omitempty" name:"MultiZoneFlag"`

	// <p>当实例为多可用区实例时，多可用区 id 列表。注意参数 ZoneId 对应的多可用区需要包含在该参数数组中。ZoneId、ZoneIds不能同时为空，可通过<a href="https://cloud.tencent.com/document/product/597/55246">DescribeCkafkaZone</a>接口获取。</p>
	ZoneIds []*int64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// <p>购买实例数量。非必填，默认值为 1。当传入该参数时，会创建多个 instanceName 加后缀区分的实例</p>
	InstanceNum *int64 `json:"InstanceNum,omitnil,omitempty" name:"InstanceNum"`

	// <p>公网带宽大小，单位 Mbps。默认是没有加上免费 3Mbps 带宽。例如总共需要 3Mbps 公网带宽，此处传 0；总共需要 6Mbps 公网带宽，此处传 3。需要保证传入参数为 3 的整数倍</p>
	PublicNetworkMonthly *int64 `json:"PublicNetworkMonthly,omitnil,omitempty" name:"PublicNetworkMonthly"`

	// <p>标签</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>弹性带宽开关 0不开启  1开启（0默认)</p>
	ElasticBandwidthSwitch *int64 `json:"ElasticBandwidthSwitch,omitnil,omitempty" name:"ElasticBandwidthSwitch"`

	// <p>自定义证书Id,仅当SpecificationsType为profession时生效,支持自定义证书能力</p><p>可通过<a href="https://cloud.tencent.com/document/product/400/41673">DescribeCertificateDetail</a>接口获取</p>
	CustomSSLCertId *string `json:"CustomSSLCertId,omitnil,omitempty" name:"CustomSSLCertId"`
}

type CreatePostPaidInstanceRequest struct {
	*tchttp.BaseRequest
	
	// <p>私有网络Id,可通过<a href="https://cloud.tencent.com/document/product/215/15778">DescribeVpcs</a>接口获取</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>子网Id,可通过<a href="https://cloud.tencent.com/document/product/215/15784">DescribeSubnets</a>接口获取</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>ckafka集群实例名称，是一个长度不超过128的任意字符。</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>国际站标准版实例规格。目前只有国际站标准版使用当前字段区分规格，国内站标准版使用峰值带宽区分规格。除了国际站标准版外的所有实例填写 1 即可。国际站标准版实例：入门型(general)]填写1；[标准型(standard)]填写2；[进阶型(advanced)]填写3；[容量型(capacity)]填写4；[高阶型1(specialized-1)]填写5；[高阶型2(specialized-2)]填写6；[高阶型3(specialized-3)]填写7；[高阶型4(specialized-4)]填写8。</p>
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// <p>实例日志的默认最长保留时间，单位分钟。不传入该参数时默认为 1440 分钟（1天），最大30天。当 topic 显式设置消息保留时间时，以 topic 保留时间为准</p>
	MsgRetentionTime *int64 `json:"MsgRetentionTime,omitnil,omitempty" name:"MsgRetentionTime"`

	// <p>创建实例时可以选择集群Id, 该入参表示集群Id。不指定实例所在集群则不传入该参数</p>
	ClusterId *int64 `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>实例版本。目前支持当前支持"2.4.1", "2.4.2","2.8.1", "3.2.3"，默认取值"2.4.1"。"2.4.1" 与 "2.4.2" 属于同一个版本，传任意一个均可。</p>
	KafkaVersion *string `json:"KafkaVersion,omitnil,omitempty" name:"KafkaVersion"`

	// <p>实例类型。"standard"：标准版，"profession"：专业版。  (标准版仅国际站支持，国内站目前支持专业版)</p>
	SpecificationsType *string `json:"SpecificationsType,omitnil,omitempty" name:"SpecificationsType"`

	// <p>专业版实例磁盘类型，标准版实例不需要填写。"CLOUD_SSD"：SSD云硬盘；"CLOUD_BASIC"：高性能云硬盘。不传默认值为 "CLOUD_BASIC"</p>
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// <p>实例内网峰值带宽，默认值为40。单位 MB/s。标准版需传入当前实例规格所对应的峰值带宽。注意如果创建的实例为专业版实例，峰值带宽，分区数等参数配置需要满足专业版的计费规格，可以通过以下链接查看计费规格：https://cloud.tencent.com/document/product/597/11745</p>
	BandWidth *int64 `json:"BandWidth,omitnil,omitempty" name:"BandWidth"`

	// <p>实例硬盘大小，默认取值为500，步长设置为100。需要满足当前实例的计费规格，可以通过以下链接查看计费规格：https://cloud.tencent.com/document/product/597/122562</p>
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// <p>实例最大分区数量，需要满足当前实例的计费规格。默认值为800，步长为100。可以通过以下链接查看计费规格：https://cloud.tencent.com/document/product/597/122563</p>
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// <p>实例最大 topic 数量，需要满足当前实例的计费规格。默认值为800，步长设置为100。</p>
	TopicNum *int64 `json:"TopicNum,omitnil,omitempty" name:"TopicNum"`

	// <p>实例所在的可用区。当创建多可用区实例时，该参数为创建的默认接入点所在子网的可用区 id。ZoneId、ZoneIds不能同时为空，可通过<a href="https://cloud.tencent.com/document/product/597/55246">DescribeCkafkaZone</a>接口获取。</p>
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>当前实例是否为多可用区实例。</p>
	MultiZoneFlag *bool `json:"MultiZoneFlag,omitnil,omitempty" name:"MultiZoneFlag"`

	// <p>当实例为多可用区实例时，多可用区 id 列表。注意参数 ZoneId 对应的多可用区需要包含在该参数数组中。ZoneId、ZoneIds不能同时为空，可通过<a href="https://cloud.tencent.com/document/product/597/55246">DescribeCkafkaZone</a>接口获取。</p>
	ZoneIds []*int64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// <p>购买实例数量。非必填，默认值为 1。当传入该参数时，会创建多个 instanceName 加后缀区分的实例</p>
	InstanceNum *int64 `json:"InstanceNum,omitnil,omitempty" name:"InstanceNum"`

	// <p>公网带宽大小，单位 Mbps。默认是没有加上免费 3Mbps 带宽。例如总共需要 3Mbps 公网带宽，此处传 0；总共需要 6Mbps 公网带宽，此处传 3。需要保证传入参数为 3 的整数倍</p>
	PublicNetworkMonthly *int64 `json:"PublicNetworkMonthly,omitnil,omitempty" name:"PublicNetworkMonthly"`

	// <p>标签</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>弹性带宽开关 0不开启  1开启（0默认)</p>
	ElasticBandwidthSwitch *int64 `json:"ElasticBandwidthSwitch,omitnil,omitempty" name:"ElasticBandwidthSwitch"`

	// <p>自定义证书Id,仅当SpecificationsType为profession时生效,支持自定义证书能力</p><p>可通过<a href="https://cloud.tencent.com/document/product/400/41673">DescribeCertificateDetail</a>接口获取</p>
	CustomSSLCertId *string `json:"CustomSSLCertId,omitnil,omitempty" name:"CustomSSLCertId"`
}

func (r *CreatePostPaidInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePostPaidInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "InstanceName")
	delete(f, "InstanceType")
	delete(f, "MsgRetentionTime")
	delete(f, "ClusterId")
	delete(f, "KafkaVersion")
	delete(f, "SpecificationsType")
	delete(f, "DiskType")
	delete(f, "BandWidth")
	delete(f, "DiskSize")
	delete(f, "Partition")
	delete(f, "TopicNum")
	delete(f, "ZoneId")
	delete(f, "MultiZoneFlag")
	delete(f, "ZoneIds")
	delete(f, "InstanceNum")
	delete(f, "PublicNetworkMonthly")
	delete(f, "Tags")
	delete(f, "ElasticBandwidthSwitch")
	delete(f, "CustomSSLCertId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePostPaidInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePostPaidInstanceResponseParams struct {
	// <p>返回结果</p>
	Result *CreateInstancePostResp `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePostPaidInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreatePostPaidInstanceResponseParams `json:"Response"`
}

func (r *CreatePostPaidInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePostPaidInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusRequestParams struct {
	// ckafka集群实例Id,可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 私有网络Id,可通过[DescribeVpcs](https://cloud.tencent.com/document/product/215/15778)接口获取
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网Id,可通过[DescribeSubnets](https://cloud.tencent.com/document/product/215/15784)接口获取
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`
}

type CreatePrometheusRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id,可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 私有网络Id,可通过[DescribeVpcs](https://cloud.tencent.com/document/product/215/15778)接口获取
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网Id,可通过[DescribeSubnets](https://cloud.tencent.com/document/product/215/15784)接口获取
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`
}

func (r *CreatePrometheusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrometheusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusResponseParams struct {
	// 打通普罗米修斯
	Result *PrometheusResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePrometheusResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrometheusResponseParams `json:"Response"`
}

func (r *CreatePrometheusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRouteRequestParams struct {
	// <p>ckafka集群实例id,可通过<a href="https://cloud.tencent.com/document/product/597/40835">DescribeInstances</a>接口获取</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>路由网络类型(3:vpc路由;7:内部支撑路由;1:公网路由)</p>
	VipType *int64 `json:"VipType,omitnil,omitempty" name:"VipType"`

	// <p>vpc网络Id,当vipType为3时必填</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>vpc子网id,当vipType为3时必填</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>访问类型：0-plaintext；1-sasl_plaintext；3-sasl_ssl; 4-sasl_scram_sha_256; 5-sasl_scram_sha_512  默认为0vipType=3,支持 0,1,3,4,5vipType=7,支持0,1,3vipType=1,支持1,3</p>
	AccessType *int64 `json:"AccessType,omitnil,omitempty" name:"AccessType"`

	// <p>是否需要权限管理,该字段已废弃</p>
	AuthFlag *int64 `json:"AuthFlag,omitnil,omitempty" name:"AuthFlag"`

	// <p>调用方appId</p>
	CallerAppid *int64 `json:"CallerAppid,omitnil,omitempty" name:"CallerAppid"`

	// <p>公网带宽,公网路由必传,且是3的倍数,无默认值</p>
	PublicNetwork *int64 `json:"PublicNetwork,omitnil,omitempty" name:"PublicNetwork"`

	// <p>vip地址</p>
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// <p>备注信息</p>
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// <p>关联安全组有序列表</p>
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

type CreateRouteRequest struct {
	*tchttp.BaseRequest
	
	// <p>ckafka集群实例id,可通过<a href="https://cloud.tencent.com/document/product/597/40835">DescribeInstances</a>接口获取</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>路由网络类型(3:vpc路由;7:内部支撑路由;1:公网路由)</p>
	VipType *int64 `json:"VipType,omitnil,omitempty" name:"VipType"`

	// <p>vpc网络Id,当vipType为3时必填</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>vpc子网id,当vipType为3时必填</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>访问类型：0-plaintext；1-sasl_plaintext；3-sasl_ssl; 4-sasl_scram_sha_256; 5-sasl_scram_sha_512  默认为0vipType=3,支持 0,1,3,4,5vipType=7,支持0,1,3vipType=1,支持1,3</p>
	AccessType *int64 `json:"AccessType,omitnil,omitempty" name:"AccessType"`

	// <p>是否需要权限管理,该字段已废弃</p>
	AuthFlag *int64 `json:"AuthFlag,omitnil,omitempty" name:"AuthFlag"`

	// <p>调用方appId</p>
	CallerAppid *int64 `json:"CallerAppid,omitnil,omitempty" name:"CallerAppid"`

	// <p>公网带宽,公网路由必传,且是3的倍数,无默认值</p>
	PublicNetwork *int64 `json:"PublicNetwork,omitnil,omitempty" name:"PublicNetwork"`

	// <p>vip地址</p>
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// <p>备注信息</p>
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// <p>关联安全组有序列表</p>
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

func (r *CreateRouteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRouteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "VipType")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "AccessType")
	delete(f, "AuthFlag")
	delete(f, "CallerAppid")
	delete(f, "PublicNetwork")
	delete(f, "Ip")
	delete(f, "Note")
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRouteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRouteResponseParams struct {
	// <p>返回结果</p>
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRouteResponse struct {
	*tchttp.BaseResponse
	Response *CreateRouteResponseParams `json:"Response"`
}

func (r *CreateRouteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTokenRequestParams struct {
	// ckafka集群实例Id,可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名
	User *string `json:"User,omitnil,omitempty" name:"User"`
}

type CreateTokenRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id,可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名
	User *string `json:"User,omitnil,omitempty" name:"User"`
}

func (r *CreateTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "User")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTokenResponseParams struct {
	// token串
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTokenResponse struct {
	*tchttp.BaseResponse
	Response *CreateTokenResponseParams `json:"Response"`
}

func (r *CreateTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTopicIpWhiteListRequestParams struct {
	// ckafka集群实例Id，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名称，可通过[DescribeTopic](https://cloud.tencent.com/document/product/597/40847)接口获取
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// ip白名单列表，最大值为512，即最大允许传入512个ip。
	IpWhiteList []*string `json:"IpWhiteList,omitnil,omitempty" name:"IpWhiteList"`
}

type CreateTopicIpWhiteListRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名称，可通过[DescribeTopic](https://cloud.tencent.com/document/product/597/40847)接口获取
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// ip白名单列表，最大值为512，即最大允许传入512个ip。
	IpWhiteList []*string `json:"IpWhiteList,omitnil,omitempty" name:"IpWhiteList"`
}

func (r *CreateTopicIpWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTopicIpWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "TopicName")
	delete(f, "IpWhiteList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTopicIpWhiteListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTopicIpWhiteListResponseParams struct {
	// 删除主题IP白名单结果
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTopicIpWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *CreateTopicIpWhiteListResponseParams `json:"Response"`
}

func (r *CreateTopicIpWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTopicIpWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTopicRequestParams struct {
	// 实例Id，可通过DescribeInstances接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 只能包含字母、数字、下划线、“-”、“.”
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Partition个数，大于0
	PartitionNum *int64 `json:"PartitionNum,omitnil,omitempty" name:"PartitionNum"`

	// 副本个数，不能多于 broker 数，最大为3
	ReplicaNum *int64 `json:"ReplicaNum,omitnil,omitempty" name:"ReplicaNum"`

	// ip白名单开关, 1:打开  0:关闭，默认不打开
	EnableWhiteList *int64 `json:"EnableWhiteList,omitnil,omitempty" name:"EnableWhiteList"`

	// Ip白名单列表，配额限制，enableWhileList=1时必选
	IpWhiteList []*string `json:"IpWhiteList,omitnil,omitempty" name:"IpWhiteList"`

	// 清理日志策略，日志清理模式，默认为"delete"。"delete"：日志按保存时间删除，"compact"：日志按 key 压缩，"compact, delete"：日志按 key 压缩且会按保存时间删除。
	CleanUpPolicy *string `json:"CleanUpPolicy,omitnil,omitempty" name:"CleanUpPolicy"`

	// 主题备注，是一个不超过 64 个字符的字符串，可以用字母和数字为首字符，剩余部分可以包含字母、数字和横划线(-)
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// 最小同步副本数，默认为1
	MinInsyncReplicas *int64 `json:"MinInsyncReplicas,omitnil,omitempty" name:"MinInsyncReplicas"`

	// 是否允许未同步的副本选为leader，0:不允许，1:允许，默认不允许
	UncleanLeaderElectionEnable *int64 `json:"UncleanLeaderElectionEnable,omitnil,omitempty" name:"UncleanLeaderElectionEnable"`

	// 可选参数。消息保留时间，单位ms，当前最小值为60000。默认值为7200000ms（2小时），最大值为7776000000 ms（90天）。
	RetentionMs *int64 `json:"RetentionMs,omitnil,omitempty" name:"RetentionMs"`

	// Segment分片滚动的时长，单位ms，最小值为86400000ms（1天）。
	SegmentMs *int64 `json:"SegmentMs,omitnil,omitempty" name:"SegmentMs"`

	// 主题消息最大值，单位为 Byte，最小值1024Bytes(即1KB)，最大值为12582912Bytes（即12MB）
	MaxMessageBytes *int64 `json:"MaxMessageBytes,omitnil,omitempty" name:"MaxMessageBytes"`

	// 预设ACL规则, 1:打开  0:关闭，默认不打开
	EnableAclRule *int64 `json:"EnableAclRule,omitnil,omitempty" name:"EnableAclRule"`

	// 预设ACL规则的名称
	AclRuleName *string `json:"AclRuleName,omitnil,omitempty" name:"AclRuleName"`

	// 可选, 保留文件大小. 默认为-1,单位Byte, 当前最小值为1073741824。
	RetentionBytes *int64 `json:"RetentionBytes,omitnil,omitempty" name:"RetentionBytes"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 消息保存的时间类型:CreateTime/LogAppendTime
	LogMsgTimestampType *string `json:"LogMsgTimestampType,omitnil,omitempty" name:"LogMsgTimestampType"`
}

type CreateTopicRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id，可通过DescribeInstances接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 只能包含字母、数字、下划线、“-”、“.”
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Partition个数，大于0
	PartitionNum *int64 `json:"PartitionNum,omitnil,omitempty" name:"PartitionNum"`

	// 副本个数，不能多于 broker 数，最大为3
	ReplicaNum *int64 `json:"ReplicaNum,omitnil,omitempty" name:"ReplicaNum"`

	// ip白名单开关, 1:打开  0:关闭，默认不打开
	EnableWhiteList *int64 `json:"EnableWhiteList,omitnil,omitempty" name:"EnableWhiteList"`

	// Ip白名单列表，配额限制，enableWhileList=1时必选
	IpWhiteList []*string `json:"IpWhiteList,omitnil,omitempty" name:"IpWhiteList"`

	// 清理日志策略，日志清理模式，默认为"delete"。"delete"：日志按保存时间删除，"compact"：日志按 key 压缩，"compact, delete"：日志按 key 压缩且会按保存时间删除。
	CleanUpPolicy *string `json:"CleanUpPolicy,omitnil,omitempty" name:"CleanUpPolicy"`

	// 主题备注，是一个不超过 64 个字符的字符串，可以用字母和数字为首字符，剩余部分可以包含字母、数字和横划线(-)
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// 最小同步副本数，默认为1
	MinInsyncReplicas *int64 `json:"MinInsyncReplicas,omitnil,omitempty" name:"MinInsyncReplicas"`

	// 是否允许未同步的副本选为leader，0:不允许，1:允许，默认不允许
	UncleanLeaderElectionEnable *int64 `json:"UncleanLeaderElectionEnable,omitnil,omitempty" name:"UncleanLeaderElectionEnable"`

	// 可选参数。消息保留时间，单位ms，当前最小值为60000。默认值为7200000ms（2小时），最大值为7776000000 ms（90天）。
	RetentionMs *int64 `json:"RetentionMs,omitnil,omitempty" name:"RetentionMs"`

	// Segment分片滚动的时长，单位ms，最小值为86400000ms（1天）。
	SegmentMs *int64 `json:"SegmentMs,omitnil,omitempty" name:"SegmentMs"`

	// 主题消息最大值，单位为 Byte，最小值1024Bytes(即1KB)，最大值为12582912Bytes（即12MB）
	MaxMessageBytes *int64 `json:"MaxMessageBytes,omitnil,omitempty" name:"MaxMessageBytes"`

	// 预设ACL规则, 1:打开  0:关闭，默认不打开
	EnableAclRule *int64 `json:"EnableAclRule,omitnil,omitempty" name:"EnableAclRule"`

	// 预设ACL规则的名称
	AclRuleName *string `json:"AclRuleName,omitnil,omitempty" name:"AclRuleName"`

	// 可选, 保留文件大小. 默认为-1,单位Byte, 当前最小值为1073741824。
	RetentionBytes *int64 `json:"RetentionBytes,omitnil,omitempty" name:"RetentionBytes"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 消息保存的时间类型:CreateTime/LogAppendTime
	LogMsgTimestampType *string `json:"LogMsgTimestampType,omitnil,omitempty" name:"LogMsgTimestampType"`
}

func (r *CreateTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "TopicName")
	delete(f, "PartitionNum")
	delete(f, "ReplicaNum")
	delete(f, "EnableWhiteList")
	delete(f, "IpWhiteList")
	delete(f, "CleanUpPolicy")
	delete(f, "Note")
	delete(f, "MinInsyncReplicas")
	delete(f, "UncleanLeaderElectionEnable")
	delete(f, "RetentionMs")
	delete(f, "SegmentMs")
	delete(f, "MaxMessageBytes")
	delete(f, "EnableAclRule")
	delete(f, "AclRuleName")
	delete(f, "RetentionBytes")
	delete(f, "Tags")
	delete(f, "LogMsgTimestampType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateTopicResp struct {
	// 主题Id
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

// Predefined struct for user
type CreateTopicResponseParams struct {
	// 返回创建结果
	Result *CreateTopicResp `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTopicResponse struct {
	*tchttp.BaseResponse
	Response *CreateTopicResponseParams `json:"Response"`
}

func (r *CreateTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserRequestParams struct {
	// ckafka集群实例Id，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 用户密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

type CreateUserRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 用户密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

func (r *CreateUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Name")
	delete(f, "Password")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserResponseParams struct {
	// 返回结果
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateUserResponse struct {
	*tchttp.BaseResponse
	Response *CreateUserResponseParams `json:"Response"`
}

func (r *CreateUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CtsdbConnectParam struct {
	// Ctsdb的连接port
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Ctsdb连接源的实例vip
	ServiceVip *string `json:"ServiceVip,omitnil,omitempty" name:"ServiceVip"`

	// Ctsdb连接源的vpcId
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// Ctsdb连接源的用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Ctsdb连接源的密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Ctsdb连接源的实例资源
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`
}

type CtsdbModifyConnectParam struct {
	// Ctsdb的连接port
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Ctsdb连接源的实例vip
	ServiceVip *string `json:"ServiceVip,omitnil,omitempty" name:"ServiceVip"`

	// Ctsdb连接源的vpcId
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// Ctsdb连接源的用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Ctsdb连接源的密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Ctsdb连接源的实例资源
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`
}

type CtsdbParam struct {
	// 连接管理实例资源
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// Ctsdb的metric
	CtsdbMetric *string `json:"CtsdbMetric,omitnil,omitempty" name:"CtsdbMetric"`
}

type CvmAndIpInfo struct {
	// ckafka集群实例Id
	CkafkaInstanceId *string `json:"CkafkaInstanceId,omitnil,omitempty" name:"CkafkaInstanceId"`

	// CVM实例ID(ins-test )或POD IP(10.0.0.30)  
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// IP地址
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`
}

type DatahubResource struct {
	// 资源类型  type类型如下: 
	// KAFKA,
	// EB_ES,
	// EB_COS,
	// EB_CLS,
	// EB_,
	// MONGODB,
	// HTTP,
	// TDW,
	// ES,
	// CLICKHOUSE,
	// DTS,
	// CLS,
	// COS,
	// TOPIC,
	// MYSQL,
	// MQTT,
	// MYSQL_DATA,
	// DORIS,
	// POSTGRESQL,
	// TDSQL_C_POSTGRESQL,
	// TDSQL_POSTGRESQL,
	// WAREHOUSE_POSTGRESQL,
	// TDSQL_C_MYSQL,
	// MARIADB,
	// SQLSERVER,
	// CTSDB,
	// SCF
	// 
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// ckafka配置，Type为KAFKA时必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	KafkaParam *KafkaParam `json:"KafkaParam,omitnil,omitempty" name:"KafkaParam"`

	// EB配置，Type为EB时必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventBusParam *EventBusParam `json:"EventBusParam,omitnil,omitempty" name:"EventBusParam"`

	// MongoDB配置，Type为MONGODB时必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	MongoDBParam *MongoDBParam `json:"MongoDBParam,omitnil,omitempty" name:"MongoDBParam"`

	// Es配置，Type为ES时必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	EsParam *EsParam `json:"EsParam,omitnil,omitempty" name:"EsParam"`

	// Tdw配置，Type为TDW时必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	TdwParam *TdwParam `json:"TdwParam,omitnil,omitempty" name:"TdwParam"`

	// Dts配置，Type为DTS时必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	DtsParam *DtsParam `json:"DtsParam,omitnil,omitempty" name:"DtsParam"`

	// ClickHouse配置，Type为CLICKHOUSE时必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClickHouseParam *ClickHouseParam `json:"ClickHouseParam,omitnil,omitempty" name:"ClickHouseParam"`

	// Cls配置，Type为CLS时必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClsParam *ClsParam `json:"ClsParam,omitnil,omitempty" name:"ClsParam"`

	// Cos配置，Type为COS时必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosParam *CosParam `json:"CosParam,omitnil,omitempty" name:"CosParam"`

	// MySQL配置，Type为MYSQL时必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	MySQLParam *MySQLParam `json:"MySQLParam,omitnil,omitempty" name:"MySQLParam"`

	// PostgreSQL配置，Type为POSTGRESQL或TDSQL_C_POSTGRESQL时必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	PostgreSQLParam *PostgreSQLParam `json:"PostgreSQLParam,omitnil,omitempty" name:"PostgreSQLParam"`

	// Topic配置，Type为Topic时必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicParam *TopicParam `json:"TopicParam,omitnil,omitempty" name:"TopicParam"`

	// MariaDB配置，Type为MARIADB时必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	MariaDBParam *MariaDBParam `json:"MariaDBParam,omitnil,omitempty" name:"MariaDBParam"`

	// SQLServer配置，Type为SQLSERVER时必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	SQLServerParam *SQLServerParam `json:"SQLServerParam,omitnil,omitempty" name:"SQLServerParam"`

	// Ctsdb配置，Type为CTSDB时必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	CtsdbParam *CtsdbParam `json:"CtsdbParam,omitnil,omitempty" name:"CtsdbParam"`

	// Scf配置，Type为SCF时必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScfParam *ScfParam `json:"ScfParam,omitnil,omitempty" name:"ScfParam"`

	// MQTT配置，Type为 MQTT 时必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	MqttParam *MqttParam `json:"MqttParam,omitnil,omitempty" name:"MqttParam"`
}

type DatahubTaskIdRes struct {
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DatahubTaskInfo struct {
	// 任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 任务类型，SOURCE数据接入，SINK数据流出
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 状态，-1创建失败，0创建中，1运行中，2删除中，3已删除，4删除失败，5暂停中，6已暂停，7暂停失败，8恢复中，9恢复失败
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 数据源
	SourceResource *DatahubResource `json:"SourceResource,omitnil,omitempty" name:"SourceResource"`

	// 数据目标
	TargetResource *DatahubResource `json:"TargetResource,omitnil,omitempty" name:"TargetResource"`

	// 任务创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 异常信息
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// 创建进度百分比
	TaskProgress *float64 `json:"TaskProgress,omitnil,omitempty" name:"TaskProgress"`

	// 任务当前处于的步骤
	TaskCurrentStep *string `json:"TaskCurrentStep,omitnil,omitempty" name:"TaskCurrentStep"`

	// Datahub转储Id
	DatahubId *string `json:"DatahubId,omitnil,omitempty" name:"DatahubId"`

	// 步骤列表
	StepList []*string `json:"StepList,omitnil,omitempty" name:"StepList"`

	// 任务描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type DatahubTopicDTO struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Topic名称
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Topic Id
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 分区数
	PartitionNum *uint64 `json:"PartitionNum,omitnil,omitempty" name:"PartitionNum"`

	// 过期时间，单位ms
	RetentionMs *uint64 `json:"RetentionMs,omitnil,omitempty" name:"RetentionMs"`

	// 备注
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// 状态，1使用中，2删除中
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type DatahubTopicResp struct {
	// 主题名称
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 主题Id
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

type DateParam struct {
	// 时间格式
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 输入类型，string，unix时间戳，默认string
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 时区，默认GMT+8
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`
}

type DealInstanceDTO struct {
	// 订单流水
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// 订单流水对应购买的 CKafka 实例 id 列表
	InstanceIdList []*string `json:"InstanceIdList,omitnil,omitempty" name:"InstanceIdList"`
}

// Predefined struct for user
type DeleteAclRequestParams struct {
	// ckafka集群实例Id，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Acl资源类型，(2:TOPIC，3:GROUP，4:CLUSTER)
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 资源名称，和resourceType相关，如当resourceType为TOPIC时，则该字段表示topic名称，当resourceType为GROUP时，该字段表示group名称，当resourceType为CLUSTER时，该字段可为空。
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// Acl操作方式，(2:ALL，3:READ，4:WRITE，5:CREATE，6:DELETE，7:ALTER，8:DESCRIBE，9:CLUSTER_ACTION，10:DESCRIBE_CONFIGS，11:ALTER_CONFIGS，12:IDEMPOTENT_WRITE)
	Operation *int64 `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 权限类型，(2:DENY，3:ALLOW)，当前ckafka支持ALLOW(相当于白名单)，其它用于后续兼容开源kafka的acl时使用
	PermissionType *int64 `json:"PermissionType,omitnil,omitempty" name:"PermissionType"`

	// 默认为\*，表示任何host都可以访问，当前ckafka不支持host为\*，但是后面开源kafka的产品化会直接支持
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 用户列表，默认为User:*，表示任何user都可以访问，当前用户只能是用户列表中包含的用户
	Principal *string `json:"Principal,omitnil,omitempty" name:"Principal"`
}

type DeleteAclRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Acl资源类型，(2:TOPIC，3:GROUP，4:CLUSTER)
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 资源名称，和resourceType相关，如当resourceType为TOPIC时，则该字段表示topic名称，当resourceType为GROUP时，该字段表示group名称，当resourceType为CLUSTER时，该字段可为空。
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// Acl操作方式，(2:ALL，3:READ，4:WRITE，5:CREATE，6:DELETE，7:ALTER，8:DESCRIBE，9:CLUSTER_ACTION，10:DESCRIBE_CONFIGS，11:ALTER_CONFIGS，12:IDEMPOTENT_WRITE)
	Operation *int64 `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 权限类型，(2:DENY，3:ALLOW)，当前ckafka支持ALLOW(相当于白名单)，其它用于后续兼容开源kafka的acl时使用
	PermissionType *int64 `json:"PermissionType,omitnil,omitempty" name:"PermissionType"`

	// 默认为\*，表示任何host都可以访问，当前ckafka不支持host为\*，但是后面开源kafka的产品化会直接支持
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 用户列表，默认为User:*，表示任何user都可以访问，当前用户只能是用户列表中包含的用户
	Principal *string `json:"Principal,omitnil,omitempty" name:"Principal"`
}

func (r *DeleteAclRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAclRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ResourceType")
	delete(f, "ResourceName")
	delete(f, "Operation")
	delete(f, "PermissionType")
	delete(f, "Host")
	delete(f, "Principal")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAclRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAclResponseParams struct {
	// 返回结果
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAclResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAclResponseParams `json:"Response"`
}

func (r *DeleteAclResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAclResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAclRuleRequestParams struct {
	// 实例id信息，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// acl规则名称，可通过DescribeAclRule接口获取。
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`
}

type DeleteAclRuleRequest struct {
	*tchttp.BaseRequest
	
	// 实例id信息，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// acl规则名称，可通过DescribeAclRule接口获取。
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`
}

func (r *DeleteAclRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAclRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "RuleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAclRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAclRuleResponseParams struct {
	// 返回被删除的规则的ID
	Result *int64 `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAclRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAclRuleResponseParams `json:"Response"`
}

func (r *DeleteAclRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAclRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteConnectResourceRequestParams struct {
	// 连接源的Id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

type DeleteConnectResourceRequest struct {
	*tchttp.BaseRequest
	
	// 连接源的Id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

func (r *DeleteConnectResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConnectResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteConnectResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteConnectResourceResponseParams struct {
	// 连接源的Id
	Result *ConnectResourceResourceIdResp `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteConnectResourceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteConnectResourceResponseParams `json:"Response"`
}

func (r *DeleteConnectResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConnectResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDatahubTaskRequestParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DeleteDatahubTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DeleteDatahubTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDatahubTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDatahubTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDatahubTaskResponseParams struct {
	// 操作结果
	Result *DatahubTaskIdRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDatahubTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDatahubTaskResponseParams `json:"Response"`
}

func (r *DeleteDatahubTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDatahubTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDatahubTopicRequestParams struct {
	// Topic名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DeleteDatahubTopicRequest struct {
	*tchttp.BaseRequest
	
	// Topic名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *DeleteDatahubTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDatahubTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDatahubTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDatahubTopicResponseParams struct {
	// 返回的结果集
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDatahubTopicResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDatahubTopicResponseParams `json:"Response"`
}

func (r *DeleteDatahubTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDatahubTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGroupRequestParams struct {
	// ckafka集群实例Id，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 消费组名称，可通过[DescribeConsumerGroup](https://cloud.tencent.com/document/product/597/40841)接口获取。
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`
}

type DeleteGroupRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 消费组名称，可通过[DescribeConsumerGroup](https://cloud.tencent.com/document/product/597/40841)接口获取。
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`
}

func (r *DeleteGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Group")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGroupResponseParams struct {
	// 返回结果
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteGroupResponseParams `json:"Response"`
}

func (r *DeleteGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGroupSubscribeTopicRequestParams struct {
	// ckafka集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 消费分组名称
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// 主题名
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`
}

type DeleteGroupSubscribeTopicRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 消费分组名称
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// 主题名
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`
}

func (r *DeleteGroupSubscribeTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGroupSubscribeTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Group")
	delete(f, "Topic")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteGroupSubscribeTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGroupSubscribeTopicResponseParams struct {
	// 返回结果
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteGroupSubscribeTopicResponse struct {
	*tchttp.BaseResponse
	Response *DeleteGroupSubscribeTopicResponseParams `json:"Response"`
}

func (r *DeleteGroupSubscribeTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGroupSubscribeTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstancePostRequestParams struct {
	// ckafka集群实例Id，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DeleteInstancePostRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DeleteInstancePostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstancePostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteInstancePostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstancePostResponseParams struct {
	// 返回结果
	Result *InstanceDeleteResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteInstancePostResponse struct {
	*tchttp.BaseResponse
	Response *DeleteInstancePostResponseParams `json:"Response"`
}

func (r *DeleteInstancePostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstancePostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstancePreRequestParams struct {
	// ckafka集群实例Id，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DeleteInstancePreRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DeleteInstancePreRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstancePreRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteInstancePreRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstancePreResponseParams struct {
	// 返回结果
	Result *CreateInstancePreResp `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteInstancePreResponse struct {
	*tchttp.BaseResponse
	Response *DeleteInstancePreResponseParams `json:"Response"`
}

func (r *DeleteInstancePreResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstancePreResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRouteRequestParams struct {
	// ckafka集群实例Id,可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 路由id,可通过[DescribeRoute](https://cloud.tencent.com/document/product/597/45484)接口获取
	RouteId *int64 `json:"RouteId,omitnil,omitempty" name:"RouteId"`

	// 调用方appId
	CallerAppid *int64 `json:"CallerAppid,omitnil,omitempty" name:"CallerAppid"`

	// 设置定时删除路由时间,仅类型为公网路由支持定时删除,可选择未来的24小时的任意时间
	DeleteRouteTime *string `json:"DeleteRouteTime,omitnil,omitempty" name:"DeleteRouteTime"`
}

type DeleteRouteRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id,可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 路由id,可通过[DescribeRoute](https://cloud.tencent.com/document/product/597/45484)接口获取
	RouteId *int64 `json:"RouteId,omitnil,omitempty" name:"RouteId"`

	// 调用方appId
	CallerAppid *int64 `json:"CallerAppid,omitnil,omitempty" name:"CallerAppid"`

	// 设置定时删除路由时间,仅类型为公网路由支持定时删除,可选择未来的24小时的任意时间
	DeleteRouteTime *string `json:"DeleteRouteTime,omitnil,omitempty" name:"DeleteRouteTime"`
}

func (r *DeleteRouteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRouteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "RouteId")
	delete(f, "CallerAppid")
	delete(f, "DeleteRouteTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRouteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRouteResponseParams struct {
	// 返回结果
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRouteResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRouteResponseParams `json:"Response"`
}

func (r *DeleteRouteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRouteTriggerTimeRequestParams struct {
	// ckafka集群实例Id,可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 修改删除路由的定时时间
	DelayTime *string `json:"DelayTime,omitnil,omitempty" name:"DelayTime"`
}

type DeleteRouteTriggerTimeRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id,可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 修改删除路由的定时时间
	DelayTime *string `json:"DelayTime,omitnil,omitempty" name:"DelayTime"`
}

func (r *DeleteRouteTriggerTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRouteTriggerTimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DelayTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRouteTriggerTimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRouteTriggerTimeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRouteTriggerTimeResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRouteTriggerTimeResponseParams `json:"Response"`
}

func (r *DeleteRouteTriggerTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRouteTriggerTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTopicIpWhiteListRequestParams struct {
	// ckafka集群实例Id，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名，可通过[DescribeTopic](https://cloud.tencent.com/document/product/597/40847)接口获取。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// ip白名单列表
	IpWhiteList []*string `json:"IpWhiteList,omitnil,omitempty" name:"IpWhiteList"`
}

type DeleteTopicIpWhiteListRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名，可通过[DescribeTopic](https://cloud.tencent.com/document/product/597/40847)接口获取。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// ip白名单列表
	IpWhiteList []*string `json:"IpWhiteList,omitnil,omitempty" name:"IpWhiteList"`
}

func (r *DeleteTopicIpWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTopicIpWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "TopicName")
	delete(f, "IpWhiteList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTopicIpWhiteListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTopicIpWhiteListResponseParams struct {
	// 删除主题IP白名单结果
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTopicIpWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTopicIpWhiteListResponseParams `json:"Response"`
}

func (r *DeleteTopicIpWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTopicIpWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTopicRequestParams struct {
	// ckafka 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// ckafka 主题名称
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

type DeleteTopicRequest struct {
	*tchttp.BaseRequest
	
	// ckafka 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// ckafka 主题名称
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

func (r *DeleteTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "TopicName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTopicResponseParams struct {
	// 返回的结果集
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTopicResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTopicResponseParams `json:"Response"`
}

func (r *DeleteTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserRequestParams struct {
	// ckafka集群实例Id，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名称，可通过[DescribeUser](https://cloud.tencent.com/document/product/597/40855)接口获取。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DeleteUserRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名称，可通过[DescribeUser](https://cloud.tencent.com/document/product/597/40855)接口获取。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *DeleteUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserResponseParams struct {
	// 返回结果
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteUserResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUserResponseParams `json:"Response"`
}

func (r *DeleteUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescModifyType struct {
	// 变配类型
	ModifyType *int64 `json:"ModifyType,omitnil,omitempty" name:"ModifyType"`

	// 是否迁移标志
	MigrateFlag *bool `json:"MigrateFlag,omitnil,omitempty" name:"MigrateFlag"`

	// 迁移预计耗时(稳定模式)秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	MigrateCostTime *int64 `json:"MigrateCostTime,omitnil,omitempty" name:"MigrateCostTime"`

	// 升配模式(1:稳定模式，2:高速模式)
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpgradeStrategy *int64 `json:"UpgradeStrategy,omitnil,omitempty" name:"UpgradeStrategy"`

	// 迁移预计耗时(高速模式)秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	MigrateCostTimeHighSpeed *int64 `json:"MigrateCostTimeHighSpeed,omitnil,omitempty" name:"MigrateCostTimeHighSpeed"`
}

// Predefined struct for user
type DescribeACLRequestParams struct {
	// ckafka集群实例Id，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Acl资源类型，(2:TOPIC，3:GROUP，4:CLUSTER)
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 资源名称，和resourceType相关，如当resourceType为TOPIC时，则该字段表示topic名称，当resourceType为GROUP时，该字段表示group名称，当resourceType为CLUSTER时，该字段可为空。
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 偏移位置
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 个数限制，默认值为50，最大值为50。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 关键字匹配
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`
}

type DescribeACLRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Acl资源类型，(2:TOPIC，3:GROUP，4:CLUSTER)
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 资源名称，和resourceType相关，如当resourceType为TOPIC时，则该字段表示topic名称，当resourceType为GROUP时，该字段表示group名称，当resourceType为CLUSTER时，该字段可为空。
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 偏移位置
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 个数限制，默认值为50，最大值为50。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 关键字匹配
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`
}

func (r *DescribeACLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeACLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ResourceType")
	delete(f, "ResourceName")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchWord")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeACLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeACLResponseParams struct {
	// 返回的ACL结果集对象
	Result *AclResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeACLResponse struct {
	*tchttp.BaseResponse
	Response *DescribeACLResponseParams `json:"Response"`
}

func (r *DescribeACLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeACLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAclRuleRequestParams struct {
	// ckafka集群实例Id，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// ACL规则名
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// ACL规则匹配类型 （PREFIXED：前缀匹配，PRESET：预设策略）
	PatternType *string `json:"PatternType,omitnil,omitempty" name:"PatternType"`

	// 是否读取简略的ACL规则，默认值为false，表示不读取简略的ACL规则。
	IsSimplified *bool `json:"IsSimplified,omitnil,omitempty" name:"IsSimplified"`
}

type DescribeAclRuleRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// ACL规则名
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// ACL规则匹配类型 （PREFIXED：前缀匹配，PRESET：预设策略）
	PatternType *string `json:"PatternType,omitnil,omitempty" name:"PatternType"`

	// 是否读取简略的ACL规则，默认值为false，表示不读取简略的ACL规则。
	IsSimplified *bool `json:"IsSimplified,omitnil,omitempty" name:"IsSimplified"`
}

func (r *DescribeAclRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAclRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "RuleName")
	delete(f, "PatternType")
	delete(f, "IsSimplified")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAclRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAclRuleResponseParams struct {
	// 返回的AclRule结果集对象
	Result *AclRuleResp `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAclRuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAclRuleResponseParams `json:"Response"`
}

func (r *DescribeAclRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAclRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCkafkaVersionRequestParams struct {
	// ckafka集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeCkafkaVersionRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeCkafkaVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCkafkaVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCkafkaVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCkafkaVersionResponseParams struct {
	// 实例版本信息
	Result *InstanceVersion `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCkafkaVersionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCkafkaVersionResponseParams `json:"Response"`
}

func (r *DescribeCkafkaVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCkafkaVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCkafkaZoneRequestParams struct {
	// cdc集群Id
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`
}

type DescribeCkafkaZoneRequest struct {
	*tchttp.BaseRequest
	
	// cdc集群Id
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`
}

func (r *DescribeCkafkaZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCkafkaZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CdcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCkafkaZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCkafkaZoneResponseParams struct {
	// 查询结果复杂对象实体
	Result *ZoneResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCkafkaZoneResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCkafkaZoneResponseParams `json:"Response"`
}

func (r *DescribeCkafkaZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCkafkaZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConnectInfoResultDTO struct {
	// ip地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpAddr *string `json:"IpAddr,omitnil,omitempty" name:"IpAddr"`

	// 连结时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 是否支持的版本
	IsUnSupportVersion *bool `json:"IsUnSupportVersion,omitnil,omitempty" name:"IsUnSupportVersion"`
}

type DescribeConnectResource struct {
	// 连接源的Id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 连接源名称
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 连接源描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 连接源类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 连接源的状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 连接源的创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 连接源的异常信息
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// 该连接源关联的Datahub任务数
	DatahubTaskCount *int64 `json:"DatahubTaskCount,omitnil,omitempty" name:"DatahubTaskCount"`

	// 连接源的当前所处步骤
	CurrentStep *string `json:"CurrentStep,omitnil,omitempty" name:"CurrentStep"`

	// 创建进度百分比
	TaskProgress *float64 `json:"TaskProgress,omitnil,omitempty" name:"TaskProgress"`

	// 步骤列表
	StepList []*string `json:"StepList,omitnil,omitempty" name:"StepList"`

	// Dts配置，Type为DTS时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	DtsConnectParam *DtsConnectParam `json:"DtsConnectParam,omitnil,omitempty" name:"DtsConnectParam"`

	// MongoDB配置，Type为MONGODB时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	MongoDBConnectParam *MongoDBConnectParam `json:"MongoDBConnectParam,omitnil,omitempty" name:"MongoDBConnectParam"`

	// Es配置，Type为ES时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	EsConnectParam *EsConnectParam `json:"EsConnectParam,omitnil,omitempty" name:"EsConnectParam"`

	// ClickHouse配置，Type为CLICKHOUSE时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClickHouseConnectParam *ClickHouseConnectParam `json:"ClickHouseConnectParam,omitnil,omitempty" name:"ClickHouseConnectParam"`

	// MySQL配置，Type为MYSQL或TDSQL_C_MYSQL时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	MySQLConnectParam *MySQLConnectParam `json:"MySQLConnectParam,omitnil,omitempty" name:"MySQLConnectParam"`

	// PostgreSQL配置，Type为POSTGRESQL或TDSQL_C_POSTGRESQL时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	PostgreSQLConnectParam *PostgreSQLConnectParam `json:"PostgreSQLConnectParam,omitnil,omitempty" name:"PostgreSQLConnectParam"`

	// MariaDB配置，Type为MARIADB时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	MariaDBConnectParam *MariaDBConnectParam `json:"MariaDBConnectParam,omitnil,omitempty" name:"MariaDBConnectParam"`

	// SQLServer配置，Type为SQLSERVER时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	SQLServerConnectParam *SQLServerConnectParam `json:"SQLServerConnectParam,omitnil,omitempty" name:"SQLServerConnectParam"`

	// Ctsdb配置，Type为CTSDB时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	CtsdbConnectParam *CtsdbConnectParam `json:"CtsdbConnectParam,omitnil,omitempty" name:"CtsdbConnectParam"`

	// Doris 配置，Type 为 DORIS 时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	DorisConnectParam *DorisConnectParam `json:"DorisConnectParam,omitnil,omitempty" name:"DorisConnectParam"`

	// Kafka配置，Type 为 KAFKA 时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	KafkaConnectParam *KafkaConnectParam `json:"KafkaConnectParam,omitnil,omitempty" name:"KafkaConnectParam"`

	// MQTT配置，Type 为 MQTT 时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	MqttConnectParam *MqttConnectParam `json:"MqttConnectParam,omitnil,omitempty" name:"MqttConnectParam"`
}

// Predefined struct for user
type DescribeConnectResourceRequestParams struct {
	// 连接源的Id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

type DescribeConnectResourceRequest struct {
	*tchttp.BaseRequest
	
	// 连接源的Id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

func (r *DescribeConnectResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConnectResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConnectResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConnectResourceResp struct {
	// 连接源的Id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 连接源名称
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 连接源描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 连接源类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 连接源的状态  枚举值: -1 (创建失败) 、0 (创建中) 、 1 (运行中)、 2 (删除中) 、 4 (删除失败) 、 5 (配置更改中) 、 6 (配置更改失败) 、 7 (异常)
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 连接源的创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 连接源的异常信息
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// 连接源的当前所处步骤
	CurrentStep *string `json:"CurrentStep,omitnil,omitempty" name:"CurrentStep"`

	// 步骤列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepList []*string `json:"StepList,omitnil,omitempty" name:"StepList"`

	// MySQL配置，Type为MYSQL或TDSQL_C_MYSQL时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	MySQLConnectParam *MySQLConnectParam `json:"MySQLConnectParam,omitnil,omitempty" name:"MySQLConnectParam"`

	// PostgreSQL配置，Type为POSTGRESQL或TDSQL_C_POSTGRESQL时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	PostgreSQLConnectParam *PostgreSQLConnectParam `json:"PostgreSQLConnectParam,omitnil,omitempty" name:"PostgreSQLConnectParam"`

	// Dts配置，Type为DTS时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	DtsConnectParam *DtsConnectParam `json:"DtsConnectParam,omitnil,omitempty" name:"DtsConnectParam"`

	// MongoDB配置，Type为MONGODB时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	MongoDBConnectParam *MongoDBConnectParam `json:"MongoDBConnectParam,omitnil,omitempty" name:"MongoDBConnectParam"`

	// Es配置，Type为ES时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	EsConnectParam *EsConnectParam `json:"EsConnectParam,omitnil,omitempty" name:"EsConnectParam"`

	// ClickHouse配置，Type为CLICKHOUSE时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClickHouseConnectParam *ClickHouseConnectParam `json:"ClickHouseConnectParam,omitnil,omitempty" name:"ClickHouseConnectParam"`

	// MariaDB配置，Type为MARIADB时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	MariaDBConnectParam *MariaDBConnectParam `json:"MariaDBConnectParam,omitnil,omitempty" name:"MariaDBConnectParam"`

	// SQLServer配置，Type为SQLSERVER时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	SQLServerConnectParam *SQLServerConnectParam `json:"SQLServerConnectParam,omitnil,omitempty" name:"SQLServerConnectParam"`

	// Ctsdb配置，Type为CTSDB时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	CtsdbConnectParam *CtsdbConnectParam `json:"CtsdbConnectParam,omitnil,omitempty" name:"CtsdbConnectParam"`

	// Doris 配置，Type 为 DORIS 时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	DorisConnectParam *DorisConnectParam `json:"DorisConnectParam,omitnil,omitempty" name:"DorisConnectParam"`

	// Kafka配置，Type 为 KAFKA 时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	KafkaConnectParam *KafkaConnectParam `json:"KafkaConnectParam,omitnil,omitempty" name:"KafkaConnectParam"`

	// MQTT配置，Type 为 MQTT 时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	MqttConnectParam *MqttConnectParam `json:"MqttConnectParam,omitnil,omitempty" name:"MqttConnectParam"`
}

// Predefined struct for user
type DescribeConnectResourceResponseParams struct {
	// 连接源数据信息
	Result *DescribeConnectResourceResp `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConnectResourceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConnectResourceResponseParams `json:"Response"`
}

func (r *DescribeConnectResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConnectResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConnectResourcesRequestParams struct {
	// 连接源类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 连接源名称的关键字查询,支持模糊匹配
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// 分页偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为1000 (超过1000,则限制为1000)
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 连接源的关键字查询, 根据地域查询本地域内连接管理列表中的连接(仅支持包含region输入的连接源)
	ResourceRegion *string `json:"ResourceRegion,omitnil,omitempty" name:"ResourceRegion"`
}

type DescribeConnectResourcesRequest struct {
	*tchttp.BaseRequest
	
	// 连接源类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 连接源名称的关键字查询,支持模糊匹配
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// 分页偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为1000 (超过1000,则限制为1000)
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 连接源的关键字查询, 根据地域查询本地域内连接管理列表中的连接(仅支持包含region输入的连接源)
	ResourceRegion *string `json:"ResourceRegion,omitnil,omitempty" name:"ResourceRegion"`
}

func (r *DescribeConnectResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConnectResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "SearchWord")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ResourceRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConnectResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConnectResourcesResp struct {
	// 连接源个数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 连接源数据
	ConnectResourceList []*DescribeConnectResource `json:"ConnectResourceList,omitnil,omitempty" name:"ConnectResourceList"`
}

// Predefined struct for user
type DescribeConnectResourcesResponseParams struct {
	// 连接源列表
	Result *DescribeConnectResourcesResp `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConnectResourcesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConnectResourcesResponseParams `json:"Response"`
}

func (r *DescribeConnectResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConnectResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConsumerGroupRequestParams struct {
	// ckafka集群实例Id,可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户需要查询的group名称。
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 用户需要查询的group中的对应的topic名称，如果指定了该参数，而group又未指定则忽略该参数。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 返回消费组的限制数量，最大支持50
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 消费组列表的起始偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeConsumerGroupRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id,可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户需要查询的group名称。
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 用户需要查询的group中的对应的topic名称，如果指定了该参数，而group又未指定则忽略该参数。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 返回消费组的限制数量，最大支持50
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 消费组列表的起始偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeConsumerGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConsumerGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "GroupName")
	delete(f, "TopicName")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConsumerGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConsumerGroupResponseParams struct {
	// 返回的消费分组信息
	Result *ConsumerGroupResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConsumerGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConsumerGroupResponseParams `json:"Response"`
}

func (r *DescribeConsumerGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConsumerGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCvmInfoRequestParams struct {
	// ckafka集群实例Id,可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeCvmInfoRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id,可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeCvmInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCvmInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCvmInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCvmInfoResponseParams struct {
	// 返回结果
	Result *ListCvmAndIpInfoRsp `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCvmInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCvmInfoResponseParams `json:"Response"`
}

func (r *DescribeCvmInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCvmInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatahubGroupOffsetsRequestParams struct {
	// （过滤条件）按照实例 ID 过滤
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Kafka 消费分组
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// 模糊匹配 topicName
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// 本次查询的偏移位置，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 本次返回结果的最大个数，默认为50，最大值为50
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeDatahubGroupOffsetsRequest struct {
	*tchttp.BaseRequest
	
	// （过滤条件）按照实例 ID 过滤
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Kafka 消费分组
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// 模糊匹配 topicName
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// 本次查询的偏移位置，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 本次返回结果的最大个数，默认为50，最大值为50
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeDatahubGroupOffsetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatahubGroupOffsetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Group")
	delete(f, "SearchWord")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatahubGroupOffsetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatahubGroupOffsetsResponseParams struct {
	// 返回的结果对象
	Result *GroupOffsetResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDatahubGroupOffsetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatahubGroupOffsetsResponseParams `json:"Response"`
}

func (r *DescribeDatahubGroupOffsetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatahubGroupOffsetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatahubTaskRequestParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeDatahubTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DescribeDatahubTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatahubTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatahubTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDatahubTaskRes struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 任务类型，SOURCE数据接入，SINK数据流出
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 状态，-1创建失败，0创建中，1运行中，2删除中，3已删除，4删除失败，5暂停中，6已暂停，7暂停失败，8恢复中，9恢复失败
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 数据源
	SourceResource *DatahubResource `json:"SourceResource,omitnil,omitempty" name:"SourceResource"`

	// 数据目标
	TargetResource *DatahubResource `json:"TargetResource,omitnil,omitempty" name:"TargetResource"`

	// Connection列表
	Connections []*Connection `json:"Connections,omitnil,omitempty" name:"Connections"`

	// 任务创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 消息处理规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransformParam *TransformParam `json:"TransformParam,omitnil,omitempty" name:"TransformParam"`

	// 数据接入ID
	DatahubId *string `json:"DatahubId,omitnil,omitempty" name:"DatahubId"`

	// 绑定的SchemaId
	SchemaId *string `json:"SchemaId,omitnil,omitempty" name:"SchemaId"`

	// 绑定的Schema名称
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// 数据处理规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransformsParam *TransformsParam `json:"TransformsParam,omitnil,omitempty" name:"TransformsParam"`

	// 异常信息
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// 任务标签列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 任务描述信息	
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 1:正常 2:隔离中
	IsolateStatus *int64 `json:"IsolateStatus,omitnil,omitempty" name:"IsolateStatus"`
}

// Predefined struct for user
type DescribeDatahubTaskResponseParams struct {
	// 返回结果
	Result *DescribeDatahubTaskRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDatahubTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatahubTaskResponseParams `json:"Response"`
}

func (r *DescribeDatahubTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatahubTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatahubTasksRequestParams struct {
	// 返回数量，默认为20，最大值为100 (超过100限制为100)
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤条件，按照 TaskName 过滤，支持模糊查询
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// 转储的目标类型
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 任务类型，SOURCE数据接入，SINK数据流出
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 转储的源类型
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 转储的资源
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`
}

type DescribeDatahubTasksRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为20，最大值为100 (超过100限制为100)
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤条件，按照 TaskName 过滤，支持模糊查询
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// 转储的目标类型
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 任务类型，SOURCE数据接入，SINK数据流出
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 转储的源类型
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 转储的资源
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`
}

func (r *DescribeDatahubTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatahubTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "SearchWord")
	delete(f, "TargetType")
	delete(f, "TaskType")
	delete(f, "SourceType")
	delete(f, "Resource")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatahubTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDatahubTasksRes struct {
	// 任务总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Datahub任务信息列表
	TaskList []*DatahubTaskInfo `json:"TaskList,omitnil,omitempty" name:"TaskList"`
}

// Predefined struct for user
type DescribeDatahubTasksResponseParams struct {
	// 返回任务查询结果
	Result *DescribeDatahubTasksRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDatahubTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatahubTasksResponseParams `json:"Response"`
}

func (r *DescribeDatahubTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatahubTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatahubTopicRequestParams struct {
	// 弹性topic名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DescribeDatahubTopicRequest struct {
	*tchttp.BaseRequest
	
	// 弹性topic名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *DescribeDatahubTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatahubTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatahubTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDatahubTopicResp struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Topic名称
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Topic Id
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 分区数
	PartitionNum *uint64 `json:"PartitionNum,omitnil,omitempty" name:"PartitionNum"`

	// 过期时间，单位ms
	RetentionMs *uint64 `json:"RetentionMs,omitnil,omitempty" name:"RetentionMs"`

	// 备注
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 状态，1使用中，2删除中
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 服务路由地址
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`
}

// Predefined struct for user
type DescribeDatahubTopicResponseParams struct {
	// 返回的结果对象
	Result *DescribeDatahubTopicResp `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDatahubTopicResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatahubTopicResponseParams `json:"Response"`
}

func (r *DescribeDatahubTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatahubTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatahubTopicsRequestParams struct {
	// 搜索词
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// 本次查询的偏移位置，默认为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 本次返回结果的最大个数，默认为50，最大值为50
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 是否从连接查询topic列表
	QueryFromConnectResource *bool `json:"QueryFromConnectResource,omitnil,omitempty" name:"QueryFromConnectResource"`

	// 连接的ID
	ConnectResourceId *string `json:"ConnectResourceId,omitnil,omitempty" name:"ConnectResourceId"`

	// topic资源表达式
	TopicRegularExpression *string `json:"TopicRegularExpression,omitnil,omitempty" name:"TopicRegularExpression"`
}

type DescribeDatahubTopicsRequest struct {
	*tchttp.BaseRequest
	
	// 搜索词
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// 本次查询的偏移位置，默认为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 本次返回结果的最大个数，默认为50，最大值为50
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 是否从连接查询topic列表
	QueryFromConnectResource *bool `json:"QueryFromConnectResource,omitnil,omitempty" name:"QueryFromConnectResource"`

	// 连接的ID
	ConnectResourceId *string `json:"ConnectResourceId,omitnil,omitempty" name:"ConnectResourceId"`

	// topic资源表达式
	TopicRegularExpression *string `json:"TopicRegularExpression,omitnil,omitempty" name:"TopicRegularExpression"`
}

func (r *DescribeDatahubTopicsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatahubTopicsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SearchWord")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "QueryFromConnectResource")
	delete(f, "ConnectResourceId")
	delete(f, "TopicRegularExpression")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatahubTopicsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDatahubTopicsResp struct {
	// 总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Topic列表
	TopicList []*DatahubTopicDTO `json:"TopicList,omitnil,omitempty" name:"TopicList"`
}

// Predefined struct for user
type DescribeDatahubTopicsResponseParams struct {
	// 主题列表
	Result *DescribeDatahubTopicsResp `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDatahubTopicsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatahubTopicsResponseParams `json:"Response"`
}

func (r *DescribeDatahubTopicsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatahubTopicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGroup struct {
	// 消费分组名称
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// 该 group 使用的协议。
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`
}

// Predefined struct for user
type DescribeGroupInfoRequestParams struct {
	// ckafka集群实例Id,可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Kafka 消费分组列表,可通过[DescribeConsumerGroup](https://cloud.tencent.com/document/product/597/40841)接口获取
	GroupList []*string `json:"GroupList,omitnil,omitempty" name:"GroupList"`
}

type DescribeGroupInfoRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id,可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Kafka 消费分组列表,可通过[DescribeConsumerGroup](https://cloud.tencent.com/document/product/597/40841)接口获取
	GroupList []*string `json:"GroupList,omitnil,omitempty" name:"GroupList"`
}

func (r *DescribeGroupInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "GroupList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGroupInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGroupInfoResponseParams struct {
	// 返回的结果
	Result []*GroupInfoResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGroupInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGroupInfoResponseParams `json:"Response"`
}

func (r *DescribeGroupInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGroupOffsetsRequestParams struct {
	// ckafka集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Kafka 消费分组
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// group 订阅的主题名称数组，如果没有该数组，则表示指定的 group 下所有 topic 信息
	Topics []*string `json:"Topics,omitnil,omitempty" name:"Topics"`

	// 模糊匹配 topicName
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// 本次查询的偏移位置，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 本次返回结果的最大个数，默认为50，最大值为50
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeGroupOffsetsRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Kafka 消费分组
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// group 订阅的主题名称数组，如果没有该数组，则表示指定的 group 下所有 topic 信息
	Topics []*string `json:"Topics,omitnil,omitempty" name:"Topics"`

	// 模糊匹配 topicName
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// 本次查询的偏移位置，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 本次返回结果的最大个数，默认为50，最大值为50
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeGroupOffsetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupOffsetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Group")
	delete(f, "Topics")
	delete(f, "SearchWord")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGroupOffsetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGroupOffsetsResponseParams struct {
	// 返回结果
	Result *GroupOffsetResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGroupOffsetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGroupOffsetsResponseParams `json:"Response"`
}

func (r *DescribeGroupOffsetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupOffsetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGroupRequestParams struct {
	// ckafka集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 搜索关键字
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大返回数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 仅支持 GroupState 筛选,   支持的筛选状态有 Empty/Stable  注意：该参数只能在2.8/3.2 版本生效
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeGroupRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 搜索关键字
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大返回数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 仅支持 GroupState 筛选,   支持的筛选状态有 Empty/Stable  注意：该参数只能在2.8/3.2 版本生效
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SearchWord")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGroupResponseParams struct {
	// 返回结果
	Result *GroupResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGroupResponseParams `json:"Response"`
}

func (r *DescribeGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceAttributesRequestParams struct {
	// ckafka集群实例Id,可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceAttributesRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id,可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceAttributesResponseParams struct {
	// 实例属性返回结果对象。
	Result *InstanceAttributesResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceAttributesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceAttributesResponseParams `json:"Response"`
}

func (r *DescribeInstanceAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesDetailRequestParams struct {
	// （过滤条件）按照实例ID过滤
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// （过滤条件）按照实例名,实例Id,可用区,私有网络id,子网id 过滤，支持模糊查询
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// （过滤条件）实例的状态。0：创建中，1：运行中，2：删除中，不填默认返回全部
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 偏移量，不填默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，不填则默认10，最大值20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 匹配标签key值。
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 过滤器。filter.Name 支持('Ip', 'VpcId', 'SubNetId', 'InstanceType','InstanceId') ,filter.Values最多传递10个值.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 已经废弃， 使用InstanceIdList
	//
	// Deprecated: InstanceIds is deprecated.
	InstanceIds *string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 按照实例ID过滤
	InstanceIdList []*string `json:"InstanceIdList,omitnil,omitempty" name:"InstanceIdList"`

	// 根据标签列表过滤实例（取交集）
	TagList []*Tag `json:"TagList,omitnil,omitempty" name:"TagList"`
}

type DescribeInstancesDetailRequest struct {
	*tchttp.BaseRequest
	
	// （过滤条件）按照实例ID过滤
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// （过滤条件）按照实例名,实例Id,可用区,私有网络id,子网id 过滤，支持模糊查询
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// （过滤条件）实例的状态。0：创建中，1：运行中，2：删除中，不填默认返回全部
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 偏移量，不填默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，不填则默认10，最大值20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 匹配标签key值。
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 过滤器。filter.Name 支持('Ip', 'VpcId', 'SubNetId', 'InstanceType','InstanceId') ,filter.Values最多传递10个值.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 已经废弃， 使用InstanceIdList
	InstanceIds *string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 按照实例ID过滤
	InstanceIdList []*string `json:"InstanceIdList,omitnil,omitempty" name:"InstanceIdList"`

	// 根据标签列表过滤实例（取交集）
	TagList []*Tag `json:"TagList,omitnil,omitempty" name:"TagList"`
}

func (r *DescribeInstancesDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SearchWord")
	delete(f, "Status")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "TagKey")
	delete(f, "Filters")
	delete(f, "InstanceIds")
	delete(f, "InstanceIdList")
	delete(f, "TagList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesDetailResponseParams struct {
	// 返回的实例详情结果对象
	Result *InstanceDetailResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstancesDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesDetailResponseParams `json:"Response"`
}

func (r *DescribeInstancesDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesRequestParams struct {
	// （查询条件）按照ckafka集群实例Id过滤
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 搜索词   ex:（查询条件）按照实例名称过滤，支持模糊查询
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// （查询条件）实例的状态  0：创建中，1：运行中，2：删除中，5: 隔离中,  7:升级中 不填默认返回全部
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 偏移量，不填默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，不填则默认10，最大值100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 已废弃。匹配标签key值。
	//
	// Deprecated: TagKey is deprecated.
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// （查询条件）私有网络Id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// （查询条件）按照ckafka集群实例Id过滤
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 搜索词   ex:（查询条件）按照实例名称过滤，支持模糊查询
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// （查询条件）实例的状态  0：创建中，1：运行中，2：删除中，5: 隔离中,  7:升级中 不填默认返回全部
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 偏移量，不填默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，不填则默认10，最大值100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 已废弃。匹配标签key值。
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// （查询条件）私有网络Id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SearchWord")
	delete(f, "Status")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "TagKey")
	delete(f, "VpcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesResponseParams struct {
	// 返回的结果
	Result *InstanceResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesResponseParams `json:"Response"`
}

func (r *DescribeInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModifyTypeRequestParams struct {
	// ckafka集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 升配后的带宽，单位mb
	BandWidth *int64 `json:"BandWidth,omitnil,omitempty" name:"BandWidth"`

	// 升配后的磁盘，单位G
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 磁盘类型，例如 CLOUD_PREMIUM
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 分区数量
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// topic数量
	Topic *int64 `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 实例类型例如 sp_ckafka_profession
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 变配入口
	ModifyEntry *string `json:"ModifyEntry,omitnil,omitempty" name:"ModifyEntry"`
}

type DescribeModifyTypeRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 升配后的带宽，单位mb
	BandWidth *int64 `json:"BandWidth,omitnil,omitempty" name:"BandWidth"`

	// 升配后的磁盘，单位G
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 磁盘类型，例如 CLOUD_PREMIUM
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 分区数量
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// topic数量
	Topic *int64 `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 实例类型例如 sp_ckafka_profession
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 变配入口
	ModifyEntry *string `json:"ModifyEntry,omitnil,omitempty" name:"ModifyEntry"`
}

func (r *DescribeModifyTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModifyTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BandWidth")
	delete(f, "DiskSize")
	delete(f, "DiskType")
	delete(f, "Partition")
	delete(f, "Topic")
	delete(f, "Type")
	delete(f, "ModifyEntry")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModifyTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModifyTypeResponseParams struct {
	// 返回的变配类型结构
	Result *DescModifyType `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModifyTypeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModifyTypeResponseParams `json:"Response"`
}

func (r *DescribeModifyTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModifyTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusRequestParams struct {
	// ckafka集群实例Id,可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribePrometheusRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id,可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribePrometheusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusResponseParams struct {
	// Prometheus监控映射列表
	Result []*PrometheusDTO `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrometheusResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusResponseParams `json:"Response"`
}

func (r *DescribePrometheusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegionRequestParams struct {
	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回最大结果数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 业务字段，可忽略
	Business *string `json:"Business,omitnil,omitempty" name:"Business"`

	// cdc专有集群业务字段，可忽略
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`
}

type DescribeRegionRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回最大结果数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 业务字段，可忽略
	Business *string `json:"Business,omitnil,omitempty" name:"Business"`

	// cdc专有集群业务字段，可忽略
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`
}

func (r *DescribeRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Business")
	delete(f, "CdcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRegionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegionResponseParams struct {
	// 返回地域枚举结果列表
	Result []*Region `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRegionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRegionResponseParams `json:"Response"`
}

func (r *DescribeRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRouteRequestParams struct {
	// ckafka集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 路由Id
	RouteId *int64 `json:"RouteId,omitnil,omitempty" name:"RouteId"`

	// 是否显示主路由，true时会在返回原路由列表的基础上,再额外展示实例创建时的主路由信息(且不被InternalFlag/UsedFor等参数过滤影响)	
	MainRouteFlag *bool `json:"MainRouteFlag,omitnil,omitempty" name:"MainRouteFlag"`
}

type DescribeRouteRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 路由Id
	RouteId *int64 `json:"RouteId,omitnil,omitempty" name:"RouteId"`

	// 是否显示主路由，true时会在返回原路由列表的基础上,再额外展示实例创建时的主路由信息(且不被InternalFlag/UsedFor等参数过滤影响)	
	MainRouteFlag *bool `json:"MainRouteFlag,omitnil,omitempty" name:"MainRouteFlag"`
}

func (r *DescribeRouteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRouteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "RouteId")
	delete(f, "MainRouteFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRouteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRouteResponseParams struct {
	// 返回的路由信息结果集
	Result *RouteResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRouteResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRouteResponseParams `json:"Response"`
}

func (r *DescribeRouteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupRoutesRequestParams struct {
	// 路由信息
	InstanceRoute *InstanceRoute `json:"InstanceRoute,omitnil,omitempty" name:"InstanceRoute"`

	// 过滤器
	Filters []*RouteFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页Offset,默认0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页Limit,默认20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 关键词,可根据实例id/实例名称/vip模糊搜索
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`
}

type DescribeSecurityGroupRoutesRequest struct {
	*tchttp.BaseRequest
	
	// 路由信息
	InstanceRoute *InstanceRoute `json:"InstanceRoute,omitnil,omitempty" name:"InstanceRoute"`

	// 过滤器
	Filters []*RouteFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页Offset,默认0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页Limit,默认20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 关键词,可根据实例id/实例名称/vip模糊搜索
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`
}

func (r *DescribeSecurityGroupRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceRoute")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchWord")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityGroupRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupRoutesResponseParams struct {
	// 返回的安全组路由信息结果对象
	Result *SecurityGroupRouteResp `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSecurityGroupRoutesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityGroupRoutesResponseParams `json:"Response"`
}

func (r *DescribeSecurityGroupRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskStatusRequestParams struct {
	// 流程Id
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`
}

type DescribeTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// 流程Id
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`
}

func (r *DescribeTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskStatusResponseParams struct {
	// 返回结果
	Result *TaskStatusResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskStatusResponseParams `json:"Response"`
}

func (r *DescribeTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicAttributesRequestParams struct {
	// ckafka集群实例Id，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名称，可通过[DescribeTopic](https://cloud.tencent.com/document/product/597/40847)接口获取
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

type DescribeTopicAttributesRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名称，可通过[DescribeTopic](https://cloud.tencent.com/document/product/597/40847)接口获取
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

func (r *DescribeTopicAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "TopicName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopicAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicAttributesResponseParams struct {
	// 返回的结果对象
	Result *TopicAttributesResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTopicAttributesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopicAttributesResponseParams `json:"Response"`
}

func (r *DescribeTopicAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicDetailRequestParams struct {
	// ckafka集群实例Id，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// （过滤条件）按照topicName过滤，支持模糊查询
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// 偏移量，不填默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，不填则默认 20，取值要大于0
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Acl预设策略名称
	AclRuleName *string `json:"AclRuleName,omitnil,omitempty" name:"AclRuleName"`

	// 根据特定的属性排序(目前支持PartitionNum/CreateTime)，默认值为CreateTime。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 0-顺序、1-倒序，默认值为0。
	OrderType *int64 `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// 目前支持 ReplicaNum （副本数）筛选
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeTopicDetailRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// （过滤条件）按照topicName过滤，支持模糊查询
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// 偏移量，不填默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，不填则默认 20，取值要大于0
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Acl预设策略名称
	AclRuleName *string `json:"AclRuleName,omitnil,omitempty" name:"AclRuleName"`

	// 根据特定的属性排序(目前支持PartitionNum/CreateTime)，默认值为CreateTime。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 0-顺序、1-倒序，默认值为0。
	OrderType *int64 `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// 目前支持 ReplicaNum （副本数）筛选
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeTopicDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SearchWord")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "AclRuleName")
	delete(f, "OrderBy")
	delete(f, "OrderType")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopicDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicDetailResponseParams struct {
	// 返回的主题详情实体
	Result *TopicDetailResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTopicDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopicDetailResponseParams `json:"Response"`
}

func (r *DescribeTopicDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicFlowRankingRequestParams struct {
	// ckafka集群实例Id，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 排行类别，PRO：Topic生产流量；CON：Topic消费流量
	RankingType *string `json:"RankingType,omitnil,omitempty" name:"RankingType"`

	// 排行起始日期
	BeginDate *string `json:"BeginDate,omitnil,omitempty" name:"BeginDate"`

	// 排行结束日期
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// Broker IP 地址
	BrokerIp *string `json:"BrokerIp,omitnil,omitempty" name:"BrokerIp"`
}

type DescribeTopicFlowRankingRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 排行类别，PRO：Topic生产流量；CON：Topic消费流量
	RankingType *string `json:"RankingType,omitnil,omitempty" name:"RankingType"`

	// 排行起始日期
	BeginDate *string `json:"BeginDate,omitnil,omitempty" name:"BeginDate"`

	// 排行结束日期
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// Broker IP 地址
	BrokerIp *string `json:"BrokerIp,omitnil,omitempty" name:"BrokerIp"`
}

func (r *DescribeTopicFlowRankingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicFlowRankingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "RankingType")
	delete(f, "BeginDate")
	delete(f, "EndDate")
	delete(f, "BrokerIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopicFlowRankingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicFlowRankingResponseParams struct {
	// 流量排行返回结果
	Result *TopicFlowRankingResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTopicFlowRankingResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopicFlowRankingResponseParams `json:"Response"`
}

func (r *DescribeTopicFlowRankingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicFlowRankingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicProduceConnectionRequestParams struct {
	// ckafka集群实例Id，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名，可通过[DescribeTopic](https://cloud.tencent.com/document/product/597/40847)接口获取。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

type DescribeTopicProduceConnectionRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名，可通过[DescribeTopic](https://cloud.tencent.com/document/product/597/40847)接口获取。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

func (r *DescribeTopicProduceConnectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicProduceConnectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "TopicName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopicProduceConnectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicProduceConnectionResponseParams struct {
	// 链接信息返回结果集
	Result []*DescribeConnectInfoResultDTO `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTopicProduceConnectionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopicProduceConnectionResponseParams `json:"Response"`
}

func (r *DescribeTopicProduceConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicProduceConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicRequestParams struct {
	// ckafka集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 过滤条件，按照 topicName 过滤，支持模糊查询
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// 偏移量，不填默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，不填则默认为20，最大值为50
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Acl预设策略名称
	AclRuleName *string `json:"AclRuleName,omitnil,omitempty" name:"AclRuleName"`
}

type DescribeTopicRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 过滤条件，按照 topicName 过滤，支持模糊查询
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// 偏移量，不填默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，不填则默认为20，最大值为50
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Acl预设策略名称
	AclRuleName *string `json:"AclRuleName,omitnil,omitempty" name:"AclRuleName"`
}

func (r *DescribeTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SearchWord")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "AclRuleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicResponseParams struct {
	// 返回的结果
	Result *TopicResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTopicResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopicResponseParams `json:"Response"`
}

func (r *DescribeTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicSubscribeGroupRequestParams struct {
	// ckafka集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 分页时的起始位置
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页时的个数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeTopicSubscribeGroupRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 分页时的起始位置
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页时的个数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeTopicSubscribeGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicSubscribeGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "TopicName")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopicSubscribeGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicSubscribeGroupResponseParams struct {
	// 返回结果
	Result *TopicSubscribeGroup `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTopicSubscribeGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopicSubscribeGroupResponseParams `json:"Response"`
}

func (r *DescribeTopicSubscribeGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicSubscribeGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicSyncReplicaRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名称
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 偏移量，不填默认为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认值为20，必须大于0。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 仅筛选未同步副本
	OutOfSyncReplicaOnly *bool `json:"OutOfSyncReplicaOnly,omitnil,omitempty" name:"OutOfSyncReplicaOnly"`
}

type DescribeTopicSyncReplicaRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名称
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 偏移量，不填默认为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认值为20，必须大于0。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 仅筛选未同步副本
	OutOfSyncReplicaOnly *bool `json:"OutOfSyncReplicaOnly,omitnil,omitempty" name:"OutOfSyncReplicaOnly"`
}

func (r *DescribeTopicSyncReplicaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicSyncReplicaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "TopicName")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OutOfSyncReplicaOnly")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopicSyncReplicaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicSyncReplicaResponseParams struct {
	// 返回topic 副本详情
	Result *TopicInSyncReplicaResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTopicSyncReplicaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopicSyncReplicaResponseParams `json:"Response"`
}

func (r *DescribeTopicSyncReplicaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicSyncReplicaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTypeInstancesRequestParams struct {
	// （过滤条件）按照实例ID过滤
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// （过滤条件）按照实例名称过滤，支持模糊查询
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// （过滤条件）实例的状态。0：创建中，1：运行中，2：删除中，不填默认返回全部
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 偏移量，不填默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，不填则默认10，最大值100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 匹配标签key值。
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`
}

type DescribeTypeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// （过滤条件）按照实例ID过滤
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// （过滤条件）按照实例名称过滤，支持模糊查询
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// （过滤条件）实例的状态。0：创建中，1：运行中，2：删除中，不填默认返回全部
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 偏移量，不填默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，不填则默认10，最大值100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 匹配标签key值。
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`
}

func (r *DescribeTypeInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTypeInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SearchWord")
	delete(f, "Status")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "TagKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTypeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTypeInstancesResponseParams struct {
	// 返回的结果
	Result *InstanceResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTypeInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTypeInstancesResponseParams `json:"Response"`
}

func (r *DescribeTypeInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTypeInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserRequestParams struct {
	// ckafka集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 按照名称过滤
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeUserRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 按照名称过滤
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SearchWord")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserResponseParams struct {
	// 返回结果
	Result *UserResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserResponseParams `json:"Response"`
}

func (r *DescribeUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DorisConnectParam struct {
	// Doris jdbc 负载均衡连接 port，通常映射到 fe 的 9030 端口
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Doris 连接源的用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Doris 连接源的密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Doris 连接源的实例资源
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// Doris 连接源的实例vip，当为腾讯云实例时，必填
	ServiceVip *string `json:"ServiceVip,omitnil,omitempty" name:"ServiceVip"`

	// Doris 连接源的vpcId，当为腾讯云实例时，必填
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 是否更新到关联的Datahub任务
	IsUpdate *bool `json:"IsUpdate,omitnil,omitempty" name:"IsUpdate"`

	// Doris 连接源是否为自建集群
	SelfBuilt *bool `json:"SelfBuilt,omitnil,omitempty" name:"SelfBuilt"`

	// Doris 的 http 负载均衡连接 port，通常映射到 be 的 8040 端口
	BePort *int64 `json:"BePort,omitnil,omitempty" name:"BePort"`
}

type DorisModifyConnectParam struct {
	// Doris 连接源的实例资源
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// Doris jdbc 负载均衡连接 port，通常映射到 fe 的 9030 端口
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Doris 连接源的实例vip，当为腾讯云实例时，必填
	ServiceVip *string `json:"ServiceVip,omitnil,omitempty" name:"ServiceVip"`

	// Doris 连接源的vpcId，当为腾讯云实例时，必填
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// Doris 连接源的用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Doris 连接源的密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 是否更新到关联的Datahub任务
	IsUpdate *bool `json:"IsUpdate,omitnil,omitempty" name:"IsUpdate"`

	// Doris 连接源是否为自建集群
	SelfBuilt *bool `json:"SelfBuilt,omitnil,omitempty" name:"SelfBuilt"`

	// Doris 的 http 负载均衡连接 port，通常映射到 be 的 8040 端口
	BePort *int64 `json:"BePort,omitnil,omitempty" name:"BePort"`
}

type DropCls struct {
	// 是否投递到cls
	DropInvalidMessageToCls *bool `json:"DropInvalidMessageToCls,omitnil,omitempty" name:"DropInvalidMessageToCls"`

	// 投递cls的地域
	DropClsRegion *string `json:"DropClsRegion,omitnil,omitempty" name:"DropClsRegion"`

	// 投递cls的账号
	DropClsOwneruin *string `json:"DropClsOwneruin,omitnil,omitempty" name:"DropClsOwneruin"`

	// 投递cls的主题
	DropClsTopicId *string `json:"DropClsTopicId,omitnil,omitempty" name:"DropClsTopicId"`

	// 投递cls的日志集id
	DropClsLogSet *string `json:"DropClsLogSet,omitnil,omitempty" name:"DropClsLogSet"`
}

type DtsConnectParam struct {
	// Dts的连接port
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Dts消费分组的Id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Dts消费分组的账号
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Dts消费分组的密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Dts实例Id
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// Dts订阅的topic
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 是否更新到关联的Datahub任务
	IsUpdate *bool `json:"IsUpdate,omitnil,omitempty" name:"IsUpdate"`
}

type DtsModifyConnectParam struct {
	// Dts实例Id【不支持修改】
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// Dts的连接port【不支持修改】
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Dts消费分组的Id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Dts消费分组的账号
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Dts消费分组的密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 是否更新到关联的Datahub任务，默认为true
	IsUpdate *bool `json:"IsUpdate,omitnil,omitempty" name:"IsUpdate"`

	// Dts订阅的topic【不支持修改】
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`
}

type DtsParam struct {
	// Dts实例Id
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// Dts的连接ip
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// Dts的连接port
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Dts订阅的topic
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// Dts消费分组的Id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Dts消费分组的账号
	GroupUser *string `json:"GroupUser,omitnil,omitempty" name:"GroupUser"`

	// Dts消费分组的密码
	GroupPassword *string `json:"GroupPassword,omitnil,omitempty" name:"GroupPassword"`

	// false同步原始数据，true同步解析后的json格式数据,默认true
	TranSql *bool `json:"TranSql,omitnil,omitempty" name:"TranSql"`
}

type DynamicDiskConfig struct {
	// 动态硬盘扩容配置开关（0: 关闭，1: 开启）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 每次磁盘动态扩容大小百分比
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepForwardPercentage *int64 `json:"StepForwardPercentage,omitnil,omitempty" name:"StepForwardPercentage"`

	// 磁盘配额百分比触发条件，即消息达到此值触发硬盘自动扩容事件
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskQuotaPercentage *int64 `json:"DiskQuotaPercentage,omitnil,omitempty" name:"DiskQuotaPercentage"`

	// 最大扩容硬盘大小，以 GB 为单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxDiskSpace *int64 `json:"MaxDiskSpace,omitnil,omitempty" name:"MaxDiskSpace"`
}

type DynamicRetentionTime struct {
	// 动态消息保留时间配置开关（0: 关闭，1: 开启）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 磁盘配额百分比触发条件，即消息达到此值触发消息保留时间变更事件
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskQuotaPercentage *int64 `json:"DiskQuotaPercentage,omitnil,omitempty" name:"DiskQuotaPercentage"`

	// 每次向前调整消息保留时间百分比
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepForwardPercentage *int64 `json:"StepForwardPercentage,omitnil,omitempty" name:"StepForwardPercentage"`

	// 保底时长，单位分钟
	// 注意：此字段可能返回 null，表示取不到有效值。
	BottomRetention *int64 `json:"BottomRetention,omitnil,omitempty" name:"BottomRetention"`
}

type EsConnectParam struct {
	// Es的连接port
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Es连接源的用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Es连接源的密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Es连接源的实例资源
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// Es连接源是否为自建集群
	SelfBuilt *bool `json:"SelfBuilt,omitnil,omitempty" name:"SelfBuilt"`

	// Es连接源的实例vip，当为腾讯云实例时，必填
	ServiceVip *string `json:"ServiceVip,omitnil,omitempty" name:"ServiceVip"`

	// Es连接源的vpcId，当为腾讯云实例时，必填
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 是否更新到关联的Datahub任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsUpdate *bool `json:"IsUpdate,omitnil,omitempty" name:"IsUpdate"`
}

type EsModifyConnectParam struct {
	// Es连接源的实例资源【不支持修改】
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// Es的连接port【不支持修改】
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Es连接源的实例vip【不支持修改】
	ServiceVip *string `json:"ServiceVip,omitnil,omitempty" name:"ServiceVip"`

	// Es连接源的vpcId【不支持修改】
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// Es连接源的用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Es连接源的密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Es连接源是否为自建集群【不支持修改】
	SelfBuilt *bool `json:"SelfBuilt,omitnil,omitempty" name:"SelfBuilt"`

	// 是否更新到关联的Datahub任务
	IsUpdate *bool `json:"IsUpdate,omitnil,omitempty" name:"IsUpdate"`
}

type EsParam struct {
	// Es实例资源Id
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// Es的连接port
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Es用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Es密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 是否为自建集群
	SelfBuilt *bool `json:"SelfBuilt,omitnil,omitempty" name:"SelfBuilt"`

	// 实例vip
	ServiceVip *string `json:"ServiceVip,omitnil,omitempty" name:"ServiceVip"`

	// 实例的vpcId
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// Es是否抛弃解析失败的消息
	DropInvalidMessage *bool `json:"DropInvalidMessage,omitnil,omitempty" name:"DropInvalidMessage"`

	// Es自定义index名称
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// Es自定义日期后缀
	DateFormat *string `json:"DateFormat,omitnil,omitempty" name:"DateFormat"`

	// 非json格式数据的自定义key
	ContentKey *string `json:"ContentKey,omitnil,omitempty" name:"ContentKey"`

	// Es是否抛弃非json格式的消息
	DropInvalidJsonMessage *bool `json:"DropInvalidJsonMessage,omitnil,omitempty" name:"DropInvalidJsonMessage"`

	// 转储到Es中的文档ID取值字段名
	DocumentIdField *string `json:"DocumentIdField,omitnil,omitempty" name:"DocumentIdField"`

	// Es自定义index名称的类型，STRING，JSONPATH，默认为STRING
	IndexType *string `json:"IndexType,omitnil,omitempty" name:"IndexType"`

	// 当设置成员参数DropInvalidMessageToCls设置为true时,DropInvalidMessage参数失效
	DropCls *DropCls `json:"DropCls,omitnil,omitempty" name:"DropCls"`

	// 转储到ES的消息为Database的binlog时，如果需要同步数据库操作，即增删改的操作到ES时填写数据库表主键
	DatabasePrimaryKey *string `json:"DatabasePrimaryKey,omitnil,omitempty" name:"DatabasePrimaryKey"`

	// 死信队列
	DropDlq *FailureParam `json:"DropDlq,omitnil,omitempty" name:"DropDlq"`

	// 使用数据订阅格式导入 es 时，消息与 es 索引字段映射关系。不填默认为默认字段匹配
	RecordMappingList []*EsRecordMapping `json:"RecordMappingList,omitnil,omitempty" name:"RecordMappingList"`

	// 消息要映射为 es 索引中 @timestamp 的字段，如果当前配置为空，则使用消息的时间戳进行映射
	DateField *string `json:"DateField,omitnil,omitempty" name:"DateField"`

	// 用来区分当前索引映射，属于新建索引还是存量索引。"EXIST_MAPPING"：从存量索引中选择；"NEW_MAPPING"：新建索引
	RecordMappingMode *string `json:"RecordMappingMode,omitnil,omitempty" name:"RecordMappingMode"`
}

type EsRecordMapping struct {
	// es 索引成员名称
	ColumnName *string `json:"ColumnName,omitnil,omitempty" name:"ColumnName"`

	// 消息字段名称
	JsonKey *string `json:"JsonKey,omitnil,omitempty" name:"JsonKey"`
}

type EventBusParam struct {
	// 资源类型。COS/ES/CLS
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 是否为自建集群
	SelfBuilt *bool `json:"SelfBuilt,omitnil,omitempty" name:"SelfBuilt"`

	// 实例资源
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// SCF云函数命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// SCF云函数函数名
	FunctionName *string `json:"FunctionName,omitnil,omitempty" name:"FunctionName"`

	// SCF云函数版本及别名
	Qualifier *string `json:"Qualifier,omitnil,omitempty" name:"Qualifier"`
}

type FailureParam struct {
	// 类型，DLQ死信队列，IGNORE_ERROR保留，DROP丢弃
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Ckafka类型死信队列
	KafkaParam *KafkaParam `json:"KafkaParam,omitnil,omitempty" name:"KafkaParam"`

	// 重试间隔
	RetryInterval *uint64 `json:"RetryInterval,omitnil,omitempty" name:"RetryInterval"`

	// 重试次数
	MaxRetryAttempts *uint64 `json:"MaxRetryAttempts,omitnil,omitempty" name:"MaxRetryAttempts"`

	// DIP Topic类型死信队列
	TopicParam *TopicParam `json:"TopicParam,omitnil,omitempty" name:"TopicParam"`

	// 死信队列类型，CKAFKA，TOPIC
	DlqType *string `json:"DlqType,omitnil,omitempty" name:"DlqType"`
}

// Predefined struct for user
type FetchDatahubMessageByOffsetRequestParams struct {
	// 弹性topic名称，可通过[DescribeDatahubTopics](https://cloud.tencent.com/document/product/597/86863)接口获取
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分区id
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// 位点信息
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type FetchDatahubMessageByOffsetRequest struct {
	*tchttp.BaseRequest
	
	// 弹性topic名称，可通过[DescribeDatahubTopics](https://cloud.tencent.com/document/product/597/86863)接口获取
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分区id
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// 位点信息
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *FetchDatahubMessageByOffsetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FetchDatahubMessageByOffsetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Partition")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FetchDatahubMessageByOffsetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FetchDatahubMessageByOffsetResponseParams struct {
	// 返回结果
	Result *ConsumerRecord `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type FetchDatahubMessageByOffsetResponse struct {
	*tchttp.BaseResponse
	Response *FetchDatahubMessageByOffsetResponseParams `json:"Response"`
}

func (r *FetchDatahubMessageByOffsetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FetchDatahubMessageByOffsetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FetchLatestDatahubMessageListRequestParams struct {
	// 弹性topic名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分区id
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// 位点信息
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大查询条数，最小1，最大100
	MessageCount *int64 `json:"MessageCount,omitnil,omitempty" name:"MessageCount"`
}

type FetchLatestDatahubMessageListRequest struct {
	*tchttp.BaseRequest
	
	// 弹性topic名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分区id
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// 位点信息
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大查询条数，最小1，最大100
	MessageCount *int64 `json:"MessageCount,omitnil,omitempty" name:"MessageCount"`
}

func (r *FetchLatestDatahubMessageListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FetchLatestDatahubMessageListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Partition")
	delete(f, "Offset")
	delete(f, "MessageCount")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FetchLatestDatahubMessageListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FetchLatestDatahubMessageListResponseParams struct {
	// 返回结果。
	Result []*ConsumerRecord `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type FetchLatestDatahubMessageListResponse struct {
	*tchttp.BaseResponse
	Response *FetchLatestDatahubMessageListResponseParams `json:"Response"`
}

func (r *FetchLatestDatahubMessageListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FetchLatestDatahubMessageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FetchMessageByOffsetRequestParams struct {
	// ckafka集群实例Id，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名，可通过[DescribeTopic](https://cloud.tencent.com/document/product/597/40847)接口获取
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 分区id
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// 位点信息
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type FetchMessageByOffsetRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名，可通过[DescribeTopic](https://cloud.tencent.com/document/product/597/40847)接口获取
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 分区id
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// 位点信息
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *FetchMessageByOffsetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FetchMessageByOffsetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Topic")
	delete(f, "Partition")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FetchMessageByOffsetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FetchMessageByOffsetResponseParams struct {
	// 返回结果
	Result *ConsumerRecord `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type FetchMessageByOffsetResponse struct {
	*tchttp.BaseResponse
	Response *FetchMessageByOffsetResponseParams `json:"Response"`
}

func (r *FetchMessageByOffsetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FetchMessageByOffsetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FetchMessageListByOffsetRequestParams struct {
	// ckafka集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 分区id
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// 位点信息
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大查询条数，默认20，最大20
	SinglePartitionRecordNumber *int64 `json:"SinglePartitionRecordNumber,omitnil,omitempty" name:"SinglePartitionRecordNumber"`
}

type FetchMessageListByOffsetRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 分区id
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// 位点信息
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大查询条数，默认20，最大20
	SinglePartitionRecordNumber *int64 `json:"SinglePartitionRecordNumber,omitnil,omitempty" name:"SinglePartitionRecordNumber"`
}

func (r *FetchMessageListByOffsetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FetchMessageListByOffsetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Topic")
	delete(f, "Partition")
	delete(f, "Offset")
	delete(f, "SinglePartitionRecordNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FetchMessageListByOffsetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FetchMessageListByOffsetResponseParams struct {
	// 返回结果。注意，列表中不返回具体的消息内容（key、value），如果需要查询具体消息内容，请使用FetchMessageByOffset接口
	Result []*ConsumerRecord `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type FetchMessageListByOffsetResponse struct {
	*tchttp.BaseResponse
	Response *FetchMessageListByOffsetResponseParams `json:"Response"`
}

func (r *FetchMessageListByOffsetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FetchMessageListByOffsetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FetchMessageListByTimestampRequestParams struct {
	// ckafka集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 分区id
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// 查询开始时间，13位时间戳
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 最大查询条数，默认20，最大20, 最小1
	SinglePartitionRecordNumber *int64 `json:"SinglePartitionRecordNumber,omitnil,omitempty" name:"SinglePartitionRecordNumber"`
}

type FetchMessageListByTimestampRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 分区id
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// 查询开始时间，13位时间戳
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 最大查询条数，默认20，最大20, 最小1
	SinglePartitionRecordNumber *int64 `json:"SinglePartitionRecordNumber,omitnil,omitempty" name:"SinglePartitionRecordNumber"`
}

func (r *FetchMessageListByTimestampRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FetchMessageListByTimestampRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Topic")
	delete(f, "Partition")
	delete(f, "StartTime")
	delete(f, "SinglePartitionRecordNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FetchMessageListByTimestampRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FetchMessageListByTimestampResponseParams struct {
	// 返回结果。注意，列表中不返回具体的消息内容（key、value），如果需要查询具体消息内容，请使用FetchMessageByOffset接口
	Result []*ConsumerRecord `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type FetchMessageListByTimestampResponse struct {
	*tchttp.BaseResponse
	Response *FetchMessageListByTimestampResponseParams `json:"Response"`
}

func (r *FetchMessageListByTimestampResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FetchMessageListByTimestampResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FieldParam struct {
	// 解析
	Analyse *AnalyseParam `json:"Analyse,omitnil,omitempty" name:"Analyse"`

	// 二次解析
	SecondaryAnalyse *SecondaryAnalyseParam `json:"SecondaryAnalyse,omitnil,omitempty" name:"SecondaryAnalyse"`

	// 数据处理
	SMT []*SMTParam `json:"SMT,omitnil,omitempty" name:"SMT"`

	// 测试结果
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 解析结果
	AnalyseResult []*SMTParam `json:"AnalyseResult,omitnil,omitempty" name:"AnalyseResult"`

	// 二次解析结果
	SecondaryAnalyseResult []*SMTParam `json:"SecondaryAnalyseResult,omitnil,omitempty" name:"SecondaryAnalyseResult"`

	// JSON格式解析结果
	AnalyseJsonResult *string `json:"AnalyseJsonResult,omitnil,omitempty" name:"AnalyseJsonResult"`

	// JSON格式二次解析结果
	SecondaryAnalyseJsonResult *string `json:"SecondaryAnalyseJsonResult,omitnil,omitempty" name:"SecondaryAnalyseJsonResult"`
}

type Filter struct {
	// 需要过滤的字段。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 字段的过滤值。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type FilterMapParam struct {
	// Key值
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 匹配模式，前缀匹配PREFIX，后缀匹配SUFFIX，包含匹配CONTAINS，EXCEPT除外匹配，数值匹配NUMBER，IP匹配IP
	MatchMode *string `json:"MatchMode,omitnil,omitempty" name:"MatchMode"`

	// Value值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 固定REGULAR
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type Group struct {
	// 消费分组名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`
}

type GroupInfoMember struct {
	// coordinator 为消费分组中的消费者生成的唯一 ID
	MemberId *string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 客户消费者 SDK 自己设置的 client.id 信息
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// 一般存储客户的 IP 地址
	ClientHost *string `json:"ClientHost,omitnil,omitempty" name:"ClientHost"`

	// 存储着分配给该消费者的 partition 信息
	Assignment *Assignment `json:"Assignment,omitnil,omitempty" name:"Assignment"`
}

type GroupInfoResponse struct {
	// 错误码，正常为0
	ErrorCode *string `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// group 状态描述（常见的为 Empty、Stable、Dead 三种状态）：
	// Dead：消费分组不存在
	// Empty：消费分组，当前没有任何消费者订阅
	// PreparingRebalance：消费分组处于 rebalance 状态
	// CompletingRebalance：消费分组处于 rebalance 状态
	// Stable：消费分组中各个消费者已经加入，处于稳定状态
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 消费分组选择的协议类型正常的消费者一般为 consumer 但有些系统采用了自己的协议如 kafka-connect 用的就是 connect。只有标准的 consumer 协议，本接口才知道具体的分配方式的格式，才能解析到具体的 partition 的分配情况
	ProtocolType *string `json:"ProtocolType,omitnil,omitempty" name:"ProtocolType"`

	// 消费者 partition 分配算法常见的有如下几种(Kafka 消费者 SDK 默认的选择项为 range)：range、 roundrobin、 sticky
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 仅当 state 为 Stable 且 protocol_type 为 consumer 时， 该数组才包含信息
	Members []*GroupInfoMember `json:"Members,omitnil,omitempty" name:"Members"`

	// 消费分组名称
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`
}

type GroupInfoTopics struct {
	// 分配的 topic 名称
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 分配的 partition 信息
	Partitions []*int64 `json:"Partitions,omitnil,omitempty" name:"Partitions"`
}

type GroupOffsetPartition struct {
	// topic 的 partitionId
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// consumer 提交的 offset 位置
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 支持消费者提交消息时，传入 metadata 作为它用，当前一般为空字符串
	Metadata *string `json:"Metadata,omitnil,omitempty" name:"Metadata"`

	// 错误码
	ErrorCode *int64 `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// 当前 partition 最新的 offset
	LogEndOffset *int64 `json:"LogEndOffset,omitnil,omitempty" name:"LogEndOffset"`

	// 未消费的消息个数
	Lag *int64 `json:"Lag,omitnil,omitempty" name:"Lag"`
}

type GroupOffsetResponse struct {
	// 符合调节的总结果数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 该主题分区数组，其中每个元素为一个 json object
	TopicList []*GroupOffsetTopic `json:"TopicList,omitnil,omitempty" name:"TopicList"`
}

type GroupOffsetTopic struct {
	// 主题名称
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 该主题分区数组，其中每个元素为一个 json object
	Partitions []*GroupOffsetPartition `json:"Partitions,omitnil,omitempty" name:"Partitions"`
}

type GroupResponse struct {
	// 计数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// GroupList
	GroupList []*DescribeGroup `json:"GroupList,omitnil,omitempty" name:"GroupList"`

	// 消费分组配额
	GroupCountQuota *uint64 `json:"GroupCountQuota,omitnil,omitempty" name:"GroupCountQuota"`
}

// Predefined struct for user
type InquireCkafkaPriceRequestParams struct {
	// 国内站标准版填写standards2, 国际站标准版填写standard,专业版填写profession,高级版填写premium
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 购买/续费付费类型(购买时不填的话, 默认获取购买包年包月一个月的费用)
	InstanceChargeParam *InstanceChargeParam `json:"InstanceChargeParam,omitnil,omitempty" name:"InstanceChargeParam"`

	// 购买/续费时购买的实例数量(不填时, 默认为1个)
	InstanceNum *int64 `json:"InstanceNum,omitnil,omitempty" name:"InstanceNum"`

	// 实例内网带宽大小, 单位MB/s (购买时必填，专业版/高级版询价时带宽信息必填)
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 实例的硬盘购买类型以及大小 (购买时必填，专业版/高级版询价时磁盘信息必填)
	InquiryDiskParam *InquiryDiskParam `json:"InquiryDiskParam,omitnil,omitempty" name:"InquiryDiskParam"`

	// 实例消息保留时间大小, 单位小时 (购买时必填)
	MessageRetention *int64 `json:"MessageRetention,omitnil,omitempty" name:"MessageRetention"`

	// 购买实例topic数, 单位个 (购买时必填)
	Topic *int64 `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 购买实例分区数, 单位个 (购买时必填，专业版/高级版询价时带宽信息必填)
	// 分区上限 最大值: 40000,步长: 100
	// 可以通过以下链接查看规格限制: https://cloud.tencent.com/document/product/597/122563
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// 购买地域, 可通过查看DescribeCkafkaZone这个接口获取ZoneId
	ZoneIds []*int64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 标记操作, 新购填写purchase, 续费填写renew, (不填时, 默认为purchase)
	CategoryAction *string `json:"CategoryAction,omitnil,omitempty" name:"CategoryAction"`

	// 国内站购买的版本, sv_ckafka_instance_s2_1(入门型), sv_ckafka_instance_s2_2(标准版), sv_ckafka_instance_s2_3(进阶型), 如果instanceType为standards2, 但该参数为空, 则默认值为sv_ckafka_instance_s2_1
	BillType *string `json:"BillType,omitnil,omitempty" name:"BillType"`

	// 公网带宽计费模式, 目前只有专业版支持公网带宽 (购买公网带宽时必填),取值为3的倍数
	PublicNetworkParam *InquiryPublicNetworkParam `json:"PublicNetworkParam,omitnil,omitempty" name:"PublicNetworkParam"`

	// 续费时的实例id, 续费时填写
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type InquireCkafkaPriceRequest struct {
	*tchttp.BaseRequest
	
	// 国内站标准版填写standards2, 国际站标准版填写standard,专业版填写profession,高级版填写premium
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 购买/续费付费类型(购买时不填的话, 默认获取购买包年包月一个月的费用)
	InstanceChargeParam *InstanceChargeParam `json:"InstanceChargeParam,omitnil,omitempty" name:"InstanceChargeParam"`

	// 购买/续费时购买的实例数量(不填时, 默认为1个)
	InstanceNum *int64 `json:"InstanceNum,omitnil,omitempty" name:"InstanceNum"`

	// 实例内网带宽大小, 单位MB/s (购买时必填，专业版/高级版询价时带宽信息必填)
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 实例的硬盘购买类型以及大小 (购买时必填，专业版/高级版询价时磁盘信息必填)
	InquiryDiskParam *InquiryDiskParam `json:"InquiryDiskParam,omitnil,omitempty" name:"InquiryDiskParam"`

	// 实例消息保留时间大小, 单位小时 (购买时必填)
	MessageRetention *int64 `json:"MessageRetention,omitnil,omitempty" name:"MessageRetention"`

	// 购买实例topic数, 单位个 (购买时必填)
	Topic *int64 `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 购买实例分区数, 单位个 (购买时必填，专业版/高级版询价时带宽信息必填)
	// 分区上限 最大值: 40000,步长: 100
	// 可以通过以下链接查看规格限制: https://cloud.tencent.com/document/product/597/122563
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// 购买地域, 可通过查看DescribeCkafkaZone这个接口获取ZoneId
	ZoneIds []*int64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 标记操作, 新购填写purchase, 续费填写renew, (不填时, 默认为purchase)
	CategoryAction *string `json:"CategoryAction,omitnil,omitempty" name:"CategoryAction"`

	// 国内站购买的版本, sv_ckafka_instance_s2_1(入门型), sv_ckafka_instance_s2_2(标准版), sv_ckafka_instance_s2_3(进阶型), 如果instanceType为standards2, 但该参数为空, 则默认值为sv_ckafka_instance_s2_1
	BillType *string `json:"BillType,omitnil,omitempty" name:"BillType"`

	// 公网带宽计费模式, 目前只有专业版支持公网带宽 (购买公网带宽时必填),取值为3的倍数
	PublicNetworkParam *InquiryPublicNetworkParam `json:"PublicNetworkParam,omitnil,omitempty" name:"PublicNetworkParam"`

	// 续费时的实例id, 续费时填写
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *InquireCkafkaPriceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquireCkafkaPriceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceType")
	delete(f, "InstanceChargeParam")
	delete(f, "InstanceNum")
	delete(f, "Bandwidth")
	delete(f, "InquiryDiskParam")
	delete(f, "MessageRetention")
	delete(f, "Topic")
	delete(f, "Partition")
	delete(f, "ZoneIds")
	delete(f, "CategoryAction")
	delete(f, "BillType")
	delete(f, "PublicNetworkParam")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquireCkafkaPriceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type InquireCkafkaPriceResp struct {
	// 实例价格
	InstancePrice *InquiryPrice `json:"InstancePrice,omitnil,omitempty" name:"InstancePrice"`

	// 公网带宽价格
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicNetworkBandwidthPrice *InquiryPrice `json:"PublicNetworkBandwidthPrice,omitnil,omitempty" name:"PublicNetworkBandwidthPrice"`
}

// Predefined struct for user
type InquireCkafkaPriceResponseParams struct {
	// 返回结果
	Result *InquireCkafkaPriceResp `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquireCkafkaPriceResponse struct {
	*tchttp.BaseResponse
	Response *InquireCkafkaPriceResponseParams `json:"Response"`
}

func (r *InquireCkafkaPriceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquireCkafkaPriceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryBasePrice struct {
	// 单位原价
	UnitPrice *float64 `json:"UnitPrice,omitnil,omitempty" name:"UnitPrice"`

	// 折扣单位价格
	UnitPriceDiscount *float64 `json:"UnitPriceDiscount,omitnil,omitempty" name:"UnitPriceDiscount"`

	// 合计原价
	OriginalPrice *float64 `json:"OriginalPrice,omitnil,omitempty" name:"OriginalPrice"`

	// 折扣合计价格
	DiscountPrice *float64 `json:"DiscountPrice,omitnil,omitempty" name:"DiscountPrice"`

	// 折扣(单位是%)
	Discount *float64 `json:"Discount,omitnil,omitempty" name:"Discount"`

	// 商品数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// 付费货币
	// 注意：此字段可能返回 null，表示取不到有效值。
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`

	// 硬盘专用返回参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 购买时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 购买时长单位("m"按月, "h"按小时)
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 购买数量
	Value *int64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type InquiryDetailPrice struct {
	// 额外内网带宽价格
	// 注意：此字段可能返回 null，表示取不到有效值。
	BandwidthPrice *InquiryBasePrice `json:"BandwidthPrice,omitnil,omitempty" name:"BandwidthPrice"`

	// 硬盘价格
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskPrice *InquiryBasePrice `json:"DiskPrice,omitnil,omitempty" name:"DiskPrice"`

	// 额外分区价格
	// 注意：此字段可能返回 null，表示取不到有效值。
	PartitionPrice *InquiryBasePrice `json:"PartitionPrice,omitnil,omitempty" name:"PartitionPrice"`

	// 额外Topic价格
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicPrice *InquiryBasePrice `json:"TopicPrice,omitnil,omitempty" name:"TopicPrice"`

	// 实例套餐价格
	InstanceTypePrice *InquiryBasePrice `json:"InstanceTypePrice,omitnil,omitempty" name:"InstanceTypePrice"`
}

type InquiryDiskParam struct {
	// 购买硬盘类型: SSD(SSD), CLOUD_SSD(SSD云硬盘), CLOUD_PREMIUM(高性能云硬盘), CLOUD_BASIC(云盘)
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 购买硬盘大小: 单位GB
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`
}

type InquiryPrice struct {
	// 单位原价
	UnitPrice *float64 `json:"UnitPrice,omitnil,omitempty" name:"UnitPrice"`

	// 折扣单位价格
	UnitPriceDiscount *float64 `json:"UnitPriceDiscount,omitnil,omitempty" name:"UnitPriceDiscount"`

	// 合计原价
	OriginalPrice *float64 `json:"OriginalPrice,omitnil,omitempty" name:"OriginalPrice"`

	// 折扣合计价格
	DiscountPrice *float64 `json:"DiscountPrice,omitnil,omitempty" name:"DiscountPrice"`

	// 折扣(单位是%)
	Discount *float64 `json:"Discount,omitnil,omitempty" name:"Discount"`

	// 商品数量
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// 付费货币
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`

	// 硬盘专用返回参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 购买时长
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 购买时长单位("m"按月, "h"按小时)
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 购买数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *int64 `json:"Value,omitnil,omitempty" name:"Value"`

	// 详细类别的价格
	// 注意：此字段可能返回 null，表示取不到有效值。
	DetailPrices *InquiryDetailPrice `json:"DetailPrices,omitnil,omitempty" name:"DetailPrices"`
}

type InquiryPublicNetworkParam struct {
	// 公网计费模式: BANDWIDTH_PREPAID(包年包月), BANDWIDTH_POSTPAID_BY_HOUR(带宽按小时计费)
	PublicNetworkChargeType *string `json:"PublicNetworkChargeType,omitnil,omitempty" name:"PublicNetworkChargeType"`

	// 公网带宽, 单位MB 取值需是0，或是3的倍数
	PublicNetworkMonthly *int64 `json:"PublicNetworkMonthly,omitnil,omitempty" name:"PublicNetworkMonthly"`
}

type Instance struct {
	// ckafka集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// ckafka集群实例Name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例的状态。0: 创建中，1: 运行中，2: 删除中,  3: 已删除,  5: 隔离中,  7: 升级中,  -1: 创建失败 
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 是否开源实例。开源：true，不开源：false
	IfCommunity *bool `json:"IfCommunity,omitnil,omitempty" name:"IfCommunity"`
}

type InstanceAttributesResponse struct {
	// <p>ckafka集群实例Id</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>ckafka集群实例Name</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>接入点 VIP 列表信息</p>
	VipList []*VipEntity `json:"VipList,omitnil,omitempty" name:"VipList"`

	// <p>虚拟IP</p>
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// <p>虚拟端口</p>
	Vport *string `json:"Vport,omitnil,omitempty" name:"Vport"`

	// <p>实例的状态。0: 创建中，1: 运行中，2: 删除中,  3: 已删除,  5: 隔离中,  7: 升级中,  -1: 创建失败 </p>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>实例带宽，单位：Mbps</p>
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// <p>实例的存储大小，单位：GB</p>
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// <p>可用区</p>
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>VPC 的 ID，为空表示是基础网络</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>子网 ID， 为空表示基础网络</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>实例健康状态， 1：健康，2：告警，3：异常</p>
	Healthy *int64 `json:"Healthy,omitnil,omitempty" name:"Healthy"`

	// <p>实例健康信息，当前会展示磁盘利用率，最大长度为256</p>
	HealthyMessage *string `json:"HealthyMessage,omitnil,omitempty" name:"HealthyMessage"`

	// <p>创建时间</p>
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>消息保存时间,单位为分钟</p>
	MsgRetentionTime *int64 `json:"MsgRetentionTime,omitnil,omitempty" name:"MsgRetentionTime"`

	// <p>自动创建 Topic 配置， 若该字段为空，则表示未开启自动创建</p>
	Config *InstanceConfigDO `json:"Config,omitnil,omitempty" name:"Config"`

	// <p>剩余创建分区数</p>
	RemainderPartitions *int64 `json:"RemainderPartitions,omitnil,omitempty" name:"RemainderPartitions"`

	// <p>剩余创建主题数</p>
	RemainderTopics *int64 `json:"RemainderTopics,omitnil,omitempty" name:"RemainderTopics"`

	// <p>当前创建分区数</p>
	CreatedPartitions *int64 `json:"CreatedPartitions,omitnil,omitempty" name:"CreatedPartitions"`

	// <p>当前创建主题数</p>
	CreatedTopics *int64 `json:"CreatedTopics,omitnil,omitempty" name:"CreatedTopics"`

	// <p>标签数组</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>过期时间</p>
	ExpireTime *uint64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// <p>可用区列表</p>
	ZoneIds []*int64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// <p>ckafka集群实例版本</p>
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// <p>最大分组数</p>
	MaxGroupNum *int64 `json:"MaxGroupNum,omitnil,omitempty" name:"MaxGroupNum"`

	// <p>售卖类型,0:标准版,1:专业版</p>
	Cvm *int64 `json:"Cvm,omitnil,omitempty" name:"Cvm"`

	// <p>实例类型  枚举列表: profession  :专业版  <br />standards2  :标准版premium   :高级版serverless  :serverless版</p>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// <p>表示该实例支持的特性。FEATURE_SUBNET_ACL:表示acl策略支持设置子网。</p>
	Features []*string `json:"Features,omitnil,omitempty" name:"Features"`

	// <p>动态消息保留策略</p>
	RetentionTimeConfig *DynamicRetentionTime `json:"RetentionTimeConfig,omitnil,omitempty" name:"RetentionTimeConfig"`

	// <p>最大连接数</p>
	MaxConnection *uint64 `json:"MaxConnection,omitnil,omitempty" name:"MaxConnection"`

	// <p>公网带宽</p>
	PublicNetwork *int64 `json:"PublicNetwork,omitnil,omitempty" name:"PublicNetwork"`

	// <p>该字段已废弃,无实际含义</p>
	DeleteRouteTimestamp *string `json:"DeleteRouteTimestamp,omitnil,omitempty" name:"DeleteRouteTimestamp"`

	// <p>剩余创建分区数</p>
	RemainingPartitions *int64 `json:"RemainingPartitions,omitnil,omitempty" name:"RemainingPartitions"`

	// <p>剩余创建主题数</p>
	RemainingTopics *int64 `json:"RemainingTopics,omitnil,omitempty" name:"RemainingTopics"`

	// <p>动态硬盘扩容策略</p>
	DynamicDiskConfig *DynamicDiskConfig `json:"DynamicDiskConfig,omitnil,omitempty" name:"DynamicDiskConfig"`

	// <p>系统维护时间</p>
	SystemMaintenanceTime *string `json:"SystemMaintenanceTime,omitnil,omitempty" name:"SystemMaintenanceTime"`

	// <p>实例级别消息最大大小</p>
	MaxMessageByte *uint64 `json:"MaxMessageByte,omitnil,omitempty" name:"MaxMessageByte"`

	// <p>实例计费类型  POSTPAID_BY_HOUR 按小时付费; PREPAID 包年包月</p>
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// <p>是否开启弹性带宽白名单 <br />1:已开启弹性带宽白名单;0:未开启弹性带宽白名单;</p>
	ElasticBandwidthSwitch *int64 `json:"ElasticBandwidthSwitch,omitnil,omitempty" name:"ElasticBandwidthSwitch"`

	// <p>弹性带宽开通状态1:未开启弹性带宽;16: 开启弹性带宽中;32:开启弹性带宽成功;33:关闭弹性带宽中;34:关闭弹性带宽成功;64:开启弹性带宽失败;65:关闭弹性带宽失败;</p>
	ElasticBandwidthOpenStatus *int64 `json:"ElasticBandwidthOpenStatus,omitnil,omitempty" name:"ElasticBandwidthOpenStatus"`

	// <p>集群类型<br />CLOUD_IDC IDC集群CLOUD_CVM_SHARE CVM共享集群CLOUD_CVM_YUNTI 云梯CVM集群CLOUD_CVM    CVM集群CLOUD_CDC CDC集群CLOUD_EKS_TSE EKS集群</p>
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// <p>免费分区数量</p>
	FreePartitionNumber *int64 `json:"FreePartitionNumber,omitnil,omitempty" name:"FreePartitionNumber"`

	// <p>弹性带宽上浮值</p>
	ElasticFloatBandwidth *int64 `json:"ElasticFloatBandwidth,omitnil,omitempty" name:"ElasticFloatBandwidth"`

	// <p>ssl自定义证书id  仅自定义证书实例集群返回</p>
	CustomCertId *string `json:"CustomCertId,omitnil,omitempty" name:"CustomCertId"`

	// <p>集群topic默认 unclean.leader.election.enable配置: 1 开启 0 关闭</p>
	UncleanLeaderElectionEnable *int64 `json:"UncleanLeaderElectionEnable,omitnil,omitempty" name:"UncleanLeaderElectionEnable"`

	// <p>实例删除保护开关: 1 开启 0 关闭</p>
	DeleteProtectionEnable *int64 `json:"DeleteProtectionEnable,omitnil,omitempty" name:"DeleteProtectionEnable"`
}

type InstanceChargeParam struct {
	// 实例付费类型: PREPAID(包年包月), POSTPAID_BY_HOUR(按量付费)
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 购买时长: 包年包月时需要填写, 按量计费无需填写
	InstanceChargePeriod *int64 `json:"InstanceChargePeriod,omitnil,omitempty" name:"InstanceChargePeriod"`
}

type InstanceConfigDO struct {
	// 是否自动创建主题
	AutoCreateTopicsEnable *bool `json:"AutoCreateTopicsEnable,omitnil,omitempty" name:"AutoCreateTopicsEnable"`

	// 分区数
	DefaultNumPartitions *int64 `json:"DefaultNumPartitions,omitnil,omitempty" name:"DefaultNumPartitions"`

	// 默认的复制Factor
	DefaultReplicationFactor *int64 `json:"DefaultReplicationFactor,omitnil,omitempty" name:"DefaultReplicationFactor"`
}

type InstanceDeleteResponse struct {
	// 删除实例返回的任务Id
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`
}

type InstanceDetail struct {
	// <p>ckafka集群实例Id</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>ckafka集群实例名称</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>访问实例的vip 信息</p>
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// <p>访问实例的端口信息</p>
	Vport *string `json:"Vport,omitnil,omitempty" name:"Vport"`

	// <p>虚拟IP列表</p>
	VipList []*VipEntity `json:"VipList,omitnil,omitempty" name:"VipList"`

	// <p>实例的状态。0: 创建中，1: 运行中，2: 删除中,  3: 已删除,  5: 隔离中,  7: 升级中,  -1: 创建失败 </p>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>实例带宽，单位Mbps</p>
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// <p>ckafka集群实例磁盘大小，单位G</p>
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// <p>可用区域ID</p>
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>vpcId，如果为空，说明是基础网络</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>子网id</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>实例是否续费，int  枚举值：1表示自动续费，2表示明确不自动续费</p>
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// <p>实例状态 int：1表示健康，2表示告警，3 表示实例状态异常</p>
	Healthy *int64 `json:"Healthy,omitnil,omitempty" name:"Healthy"`

	// <p>实例状态信息</p>
	HealthyMessage *string `json:"HealthyMessage,omitnil,omitempty" name:"HealthyMessage"`

	// <p>实例创建时间</p>
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>实例过期时间</p>
	ExpireTime *int64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// <p>是否为内部客户。值为1 表示内部客户</p>
	IsInternal *int64 `json:"IsInternal,omitnil,omitempty" name:"IsInternal"`

	// <p>Topic个数</p>
	TopicNum *int64 `json:"TopicNum,omitnil,omitempty" name:"TopicNum"`

	// <p>标识tag</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>kafka版本信息</p>
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// <p>跨可用区</p>
	ZoneIds []*int64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// <p>ckafka售卖类型</p>
	Cvm *int64 `json:"Cvm,omitnil,omitempty" name:"Cvm"`

	// <p>ckafka集群实例类型</p>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// <p>ckafka集群实例磁盘类型</p>
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// <p>当前规格最大Topic数</p>
	MaxTopicNumber *int64 `json:"MaxTopicNumber,omitnil,omitempty" name:"MaxTopicNumber"`

	// <p>当前规格最大Partition数</p>
	MaxPartitionNumber *int64 `json:"MaxPartitionNumber,omitnil,omitempty" name:"MaxPartitionNumber"`

	// <p>计划升级配置时间</p>
	RebalanceTime *string `json:"RebalanceTime,omitnil,omitempty" name:"RebalanceTime"`

	// <p>实例当前partition数量</p>
	PartitionNumber *uint64 `json:"PartitionNumber,omitnil,omitempty" name:"PartitionNumber"`

	// <p>ckafka集群实例公网带宽类型</p>
	PublicNetworkChargeType *string `json:"PublicNetworkChargeType,omitnil,omitempty" name:"PublicNetworkChargeType"`

	// <p>公网带宽 最小3Mbps  最大999Mbps 仅专业版支持填写</p>
	PublicNetwork *int64 `json:"PublicNetwork,omitnil,omitempty" name:"PublicNetwork"`

	// <p>ckafka集群实例底层集群类型</p>
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// <p>实例功能列表</p>
	Features []*string `json:"Features,omitnil,omitempty" name:"Features"`
}

type InstanceDetailResponse struct {
	// 符合条件的实例总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 符合条件的实例详情列表
	InstanceList []*InstanceDetail `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`
}

type InstanceQuotaConfigResp struct {
	// 生产限流大小，单位 MB/s
	// 注意：此字段可能返回 null，表示取不到有效值。
	QuotaProducerByteRate *int64 `json:"QuotaProducerByteRate,omitnil,omitempty" name:"QuotaProducerByteRate"`

	// 消费限流大小，单位 MB/s
	// 注意：此字段可能返回 null，表示取不到有效值。
	QuotaConsumerByteRate *int64 `json:"QuotaConsumerByteRate,omitnil,omitempty" name:"QuotaConsumerByteRate"`
}

type InstanceResponse struct {
	// 符合条件的实例列表
	InstanceList []*Instance `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// 符合条件的结果总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type InstanceRoute struct {
	// ckafka集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 路由Id
	RouteId *int64 `json:"RouteId,omitnil,omitempty" name:"RouteId"`
}

// Predefined struct for user
type InstanceScalingDownRequestParams struct {
	// ckafka集群实例Id,可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 缩容模式  1:稳定变配 
	// 2.高速变配
	UpgradeStrategy *int64 `json:"UpgradeStrategy,omitnil,omitempty" name:"UpgradeStrategy"`

	// 磁盘大小 单位 GB     最大值为500000,步长100
	// 可以通过以下链接查看规格限制：https://cloud.tencent.com/document/product/597/122562
	// 
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 峰值带宽 单位 MB/s
	// 可以通过以下链接查看规格限制及对应步长: https://cloud.tencent.com/document/product/597/11745
	BandWidth *int64 `json:"BandWidth,omitnil,omitempty" name:"BandWidth"`

	// 分区上限 最大值: 40000, 步长: 100
	// 可以通过以下链接查看规格限制: https://cloud.tencent.com/document/product/597/122563
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`
}

type InstanceScalingDownRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id,可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 缩容模式  1:稳定变配 
	// 2.高速变配
	UpgradeStrategy *int64 `json:"UpgradeStrategy,omitnil,omitempty" name:"UpgradeStrategy"`

	// 磁盘大小 单位 GB     最大值为500000,步长100
	// 可以通过以下链接查看规格限制：https://cloud.tencent.com/document/product/597/122562
	// 
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 峰值带宽 单位 MB/s
	// 可以通过以下链接查看规格限制及对应步长: https://cloud.tencent.com/document/product/597/11745
	BandWidth *int64 `json:"BandWidth,omitnil,omitempty" name:"BandWidth"`

	// 分区上限 最大值: 40000, 步长: 100
	// 可以通过以下链接查看规格限制: https://cloud.tencent.com/document/product/597/122563
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`
}

func (r *InstanceScalingDownRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InstanceScalingDownRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UpgradeStrategy")
	delete(f, "DiskSize")
	delete(f, "BandWidth")
	delete(f, "Partition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InstanceScalingDownRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InstanceScalingDownResponseParams struct {
	// 返回结果
	Result *ScalingDownResp `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InstanceScalingDownResponse struct {
	*tchttp.BaseResponse
	Response *InstanceScalingDownResponseParams `json:"Response"`
}

func (r *InstanceScalingDownResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InstanceScalingDownResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceVersion struct {
	// ckafka集群实例版本
	KafkaVersion *string `json:"KafkaVersion,omitnil,omitempty" name:"KafkaVersion"`

	// broker版本信息
	CurBrokerVersion *string `json:"CurBrokerVersion,omitnil,omitempty" name:"CurBrokerVersion"`

	// 最新版本信息
	LatestBrokerVersion []*LatestBrokerVersion `json:"LatestBrokerVersion,omitnil,omitempty" name:"LatestBrokerVersion"`

	// 允许跨大版本内核升级
	AllowUpgradeHighVersion *bool `json:"AllowUpgradeHighVersion,omitnil,omitempty" name:"AllowUpgradeHighVersion"`

	// 允许升级的大版本
	HighVersionSet []*string `json:"HighVersionSet,omitnil,omitempty" name:"HighVersionSet"`

	// 允许小版本号配置自动删除消费者组
	AllowAutoDeleteTimestamp *bool `json:"AllowAutoDeleteTimestamp,omitnil,omitempty" name:"AllowAutoDeleteTimestamp"`
}

type JgwOperateResponse struct {
	// 返回的code，0为正常，非0为错误
	ReturnCode *string `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// 成功消息
	ReturnMessage *string `json:"ReturnMessage,omitnil,omitempty" name:"ReturnMessage"`

	// 操作型返回的Data数据,可能有flowId等
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *OperateResponseData `json:"Data,omitnil,omitempty" name:"Data"`
}

type JsonPathReplaceParam struct {
	// 被替换值，Jsonpath表达式
	OldValue *string `json:"OldValue,omitnil,omitempty" name:"OldValue"`

	// 替换值，Jsonpath表达式或字符串
	NewValue *string `json:"NewValue,omitnil,omitempty" name:"NewValue"`
}

type KVParam struct {
	// 分隔符
	Delimiter *string `json:"Delimiter,omitnil,omitempty" name:"Delimiter"`

	// key-value二次解析分隔符
	Regex *string `json:"Regex,omitnil,omitempty" name:"Regex"`

	// 保留源Key，默认为false不保留
	KeepOriginalKey *string `json:"KeepOriginalKey,omitnil,omitempty" name:"KeepOriginalKey"`
}

type KafkaConnectParam struct {
	// Kafka连接源的实例资源, 非自建时必填，NetworkType=VPC时传clb实例id
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// 是否为自建集群
	SelfBuilt *bool `json:"SelfBuilt,omitnil,omitempty" name:"SelfBuilt"`

	// 是否更新到关联的Dip任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsUpdate *bool `json:"IsUpdate,omitnil,omitempty" name:"IsUpdate"`

	// Kafka连接的broker地址, NetworkType=PUBLIC公网时必填
	BrokerAddress *string `json:"BrokerAddress,omitnil,omitempty" name:"BrokerAddress"`

	// CKafka连接源的实例资源地域, 跨地域时必填
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 网络类型：PUBLIC公网；VPC
	NetworkType *string `json:"NetworkType,omitnil,omitempty" name:"NetworkType"`

	// vpcId，NetworkType=VPC时必传
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// vip，NetworkType=VPC时必传
	ServiceVip *string `json:"ServiceVip,omitnil,omitempty" name:"ServiceVip"`

	// 端口，NetworkType=VPC时必传
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 跨云同步下需要客户传递唯一Id标志一组资源
	CrossNetResourceUniqueId *string `json:"CrossNetResourceUniqueId,omitnil,omitempty" name:"CrossNetResourceUniqueId"`

	// 跨云子网ID
	CrossNetVpcSubNetId *string `json:"CrossNetVpcSubNetId,omitnil,omitempty" name:"CrossNetVpcSubNetId"`
}

type KafkaParam struct {
	// 是否为自建集群
	SelfBuilt *bool `json:"SelfBuilt,omitnil,omitempty" name:"SelfBuilt"`

	// ckafka集群实例Id
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// 主题名，多个以“,”分隔
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// Offset类型，最开始位置earliest，最新位置latest，时间点位置timestamp
	// 注意：此字段可能返回 null，表示取不到有效值。
	OffsetType *string `json:"OffsetType,omitnil,omitempty" name:"OffsetType"`

	// Offset类型为timestamp时必传，传时间戳，精确到秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 实例资源名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 主题Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Topic的分区数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PartitionNum *int64 `json:"PartitionNum,omitnil,omitempty" name:"PartitionNum"`

	// 启用容错实例/开启死信队列
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableToleration *bool `json:"EnableToleration,omitnil,omitempty" name:"EnableToleration"`

	// Qps 限制
	QpsLimit *uint64 `json:"QpsLimit,omitnil,omitempty" name:"QpsLimit"`

	// Table到Topic的路由，「分发到多个topic」开关打开时必传
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableMappings []*TableMapping `json:"TableMappings,omitnil,omitempty" name:"TableMappings"`

	// 「分发到多个topic」开关，默认为false
	UseTableMapping *bool `json:"UseTableMapping,omitnil,omitempty" name:"UseTableMapping"`

	// 使用的Topic是否需要自动创建（目前只支持SOURCE流入任务，如果不使用分发到多个topic，需要在Topic字段填写需要自动创建的topic名）
	UseAutoCreateTopic *bool `json:"UseAutoCreateTopic,omitnil,omitempty" name:"UseAutoCreateTopic"`

	// 写入Topic时是否进行压缩，不开启填"none"，开启的话，填写"open"。
	CompressionType *string `json:"CompressionType,omitnil,omitempty" name:"CompressionType"`

	// 源topic消息1条扩增成msgMultiple条写入目标topic(该参数目前只有ckafka流入ckafka适用)
	MsgMultiple *int64 `json:"MsgMultiple,omitnil,omitempty" name:"MsgMultiple"`

	// 数据同步专用参数, 正常数据处理可为空, 实例级别同步: 仅同步元数据填写"META_SYNC_INSTANCE_TYPE", 同步元数据及全部topic内消息的填写"META_AND_DATA_SYNC_INSTANCE_TYPE"; topic级别同步: 选中的源和目标topic中的消息(需要目标实例也包含该topic)填写"DATA_SYNC_TYPE"
	ConnectorSyncType *string `json:"ConnectorSyncType,omitnil,omitempty" name:"ConnectorSyncType"`

	// 数据同步专用参数, 当通过时,希望下游的消息写入分区与上游的一致,则填true,但下游分区小于上游时,会报错; 不需要一致则为false, 默认为false
	KeepPartition *bool `json:"KeepPartition,omitnil,omitempty" name:"KeepPartition"`

	// 正则匹配Topic列表
	TopicRegularExpression *string `json:"TopicRegularExpression,omitnil,omitempty" name:"TopicRegularExpression"`

	// Topic 前缀
	Prefix *string `json:"Prefix,omitnil,omitempty" name:"Prefix"`

	// Topic前缀分隔符
	Separator *string `json:"Separator,omitnil,omitempty" name:"Separator"`
}

type LatestBrokerVersion struct {
	// ckafka集群实例版本
	KafkaVersion *string `json:"KafkaVersion,omitnil,omitempty" name:"KafkaVersion"`

	// broker版本号
	BrokerVersion *string `json:"BrokerVersion,omitnil,omitempty" name:"BrokerVersion"`
}

type ListCvmAndIpInfoRsp struct {
	// cvm和IP 列表
	CvmList []*CvmAndIpInfo `json:"CvmList,omitnil,omitempty" name:"CvmList"`

	// 实例数据量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type LowercaseParam struct {

}

type MapParam struct {
	// key值
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 类型，DEFAULT默认，DATE系统预设-时间戳，CUSTOMIZE自定义，MAPPING映射
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type MariaDBConnectParam struct {
	// MariaDB的连接port
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// MariaDB连接源的用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// MariaDB连接源的密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// MariaDB连接源的实例资源
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// MariaDB连接源的实例vip，当为腾讯云实例时，必填
	ServiceVip *string `json:"ServiceVip,omitnil,omitempty" name:"ServiceVip"`

	// MariaDB连接源的vpcId，当为腾讯云实例时，必填
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 是否更新到关联的Datahub任务
	IsUpdate *bool `json:"IsUpdate,omitnil,omitempty" name:"IsUpdate"`
}

type MariaDBModifyConnectParam struct {
	// MariaDB连接源的实例资源【不支持修改】
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// MariaDB的连接port【不支持修改】
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// MariaDB连接源的实例vip【不支持修改】
	ServiceVip *string `json:"ServiceVip,omitnil,omitempty" name:"ServiceVip"`

	// MariaDB连接源的vpcId【不支持修改】
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// MariaDB连接源的用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// MariaDB连接源的密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 是否更新到关联的Datahub任务
	IsUpdate *bool `json:"IsUpdate,omitnil,omitempty" name:"IsUpdate"`
}

type MariaDBParam struct {
	// MariaDB的数据库名称，"*"为全数据库
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// MariaDB的数据表名称，"*"为所监听的所有数据库中的非系统表，可以","间隔，监听多个数据表，但数据表需要以"数据库名.数据表名"的格式进行填写
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// 该MariaDB在连接管理内的Id
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// 复制存量信息(schema_only不复制, initial全量)，默认值initial
	SnapshotMode *string `json:"SnapshotMode,omitnil,omitempty" name:"SnapshotMode"`

	// 格式：库1.表1:字段1,字段2;库2.表2:字段2，表之间;（分号）隔开，字段之间,（逗号）隔开。不指定的表默认取表的主键
	KeyColumns *string `json:"KeyColumns,omitnil,omitempty" name:"KeyColumns"`

	// 当Table输入的是前缀时，该项值为true，否则为false
	IsTablePrefix *bool `json:"IsTablePrefix,omitnil,omitempty" name:"IsTablePrefix"`

	// 输出格式，DEFAULT、CANAL_1、CANAL_2
	OutputFormat *string `json:"OutputFormat,omitnil,omitempty" name:"OutputFormat"`

	// 如果该值为all，则DDL数据以及DML数据也会写入到选中的topic；若该值为dml，则只有DML数据写入到选中的topic
	IncludeContentChanges *string `json:"IncludeContentChanges,omitnil,omitempty" name:"IncludeContentChanges"`

	// 如果该值为true，且MySQL中"binlog_rows_query_log_events"配置项的值为"ON"，则流入到topic的数据包含原SQL语句；若该值为false，流入到topic的数据不包含原SQL语句
	IncludeQuery *bool `json:"IncludeQuery,omitnil,omitempty" name:"IncludeQuery"`

	// 如果该值为 true，则消息中会携带消息结构体对应的schema，如果该值为false则不会携带
	RecordWithSchema *bool `json:"RecordWithSchema,omitnil,omitempty" name:"RecordWithSchema"`
}

// Predefined struct for user
type ModifyAclRuleRequestParams struct {
	// ckafka集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// ACL规则名
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 修改预设规则时传入,是否应用到新增的Topic
	IsApplied *int64 `json:"IsApplied,omitnil,omitempty" name:"IsApplied"`
}

type ModifyAclRuleRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// ACL规则名
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 修改预设规则时传入,是否应用到新增的Topic
	IsApplied *int64 `json:"IsApplied,omitnil,omitempty" name:"IsApplied"`
}

func (r *ModifyAclRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAclRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "RuleName")
	delete(f, "IsApplied")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAclRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAclRuleResponseParams struct {
	// 规则的唯一表示Key
	Result *int64 `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAclRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAclRuleResponseParams `json:"Response"`
}

func (r *ModifyAclRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAclRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyConnectResourceRequestParams struct {
	// 连接源的Id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 连接源名称，为空时不修改
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 连接源描述，为空时不修改
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 连接源类型，修改数据源参数时，需要与原Type相同，否则编辑数据源无效
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Dts配置，Type为DTS时必填
	DtsConnectParam *DtsModifyConnectParam `json:"DtsConnectParam,omitnil,omitempty" name:"DtsConnectParam"`

	// MongoDB配置，Type为MONGODB时必填
	MongoDBConnectParam *MongoDBModifyConnectParam `json:"MongoDBConnectParam,omitnil,omitempty" name:"MongoDBConnectParam"`

	// Es配置，Type为ES时必填
	EsConnectParam *EsModifyConnectParam `json:"EsConnectParam,omitnil,omitempty" name:"EsConnectParam"`

	// ClickHouse配置，Type为CLICKHOUSE时必填
	ClickHouseConnectParam *ClickHouseModifyConnectParam `json:"ClickHouseConnectParam,omitnil,omitempty" name:"ClickHouseConnectParam"`

	// MySQL配置，Type为MYSQL或TDSQL_C_MYSQL时必填
	MySQLConnectParam *MySQLModifyConnectParam `json:"MySQLConnectParam,omitnil,omitempty" name:"MySQLConnectParam"`

	// PostgreSQL配置，Type为POSTGRESQL或TDSQL_C_POSTGRESQL时必填
	PostgreSQLConnectParam *PostgreSQLModifyConnectParam `json:"PostgreSQLConnectParam,omitnil,omitempty" name:"PostgreSQLConnectParam"`

	// MariaDB配置，Type为MARIADB时必填
	MariaDBConnectParam *MariaDBModifyConnectParam `json:"MariaDBConnectParam,omitnil,omitempty" name:"MariaDBConnectParam"`

	// SQLServer配置，Type为SQLSERVER时必填
	SQLServerConnectParam *SQLServerModifyConnectParam `json:"SQLServerConnectParam,omitnil,omitempty" name:"SQLServerConnectParam"`

	// Ctsdb配置，Type为CTSDB
	CtsdbConnectParam *CtsdbModifyConnectParam `json:"CtsdbConnectParam,omitnil,omitempty" name:"CtsdbConnectParam"`

	// Doris配置，Type为DORIS
	DorisConnectParam *DorisModifyConnectParam `json:"DorisConnectParam,omitnil,omitempty" name:"DorisConnectParam"`

	// Kafka配置，Type为 KAFKA 时必填
	KafkaConnectParam *KafkaConnectParam `json:"KafkaConnectParam,omitnil,omitempty" name:"KafkaConnectParam"`

	// MQTT配置，Type为 MQTT 时必填
	MqttConnectParam *MqttConnectParam `json:"MqttConnectParam,omitnil,omitempty" name:"MqttConnectParam"`
}

type ModifyConnectResourceRequest struct {
	*tchttp.BaseRequest
	
	// 连接源的Id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 连接源名称，为空时不修改
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 连接源描述，为空时不修改
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 连接源类型，修改数据源参数时，需要与原Type相同，否则编辑数据源无效
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Dts配置，Type为DTS时必填
	DtsConnectParam *DtsModifyConnectParam `json:"DtsConnectParam,omitnil,omitempty" name:"DtsConnectParam"`

	// MongoDB配置，Type为MONGODB时必填
	MongoDBConnectParam *MongoDBModifyConnectParam `json:"MongoDBConnectParam,omitnil,omitempty" name:"MongoDBConnectParam"`

	// Es配置，Type为ES时必填
	EsConnectParam *EsModifyConnectParam `json:"EsConnectParam,omitnil,omitempty" name:"EsConnectParam"`

	// ClickHouse配置，Type为CLICKHOUSE时必填
	ClickHouseConnectParam *ClickHouseModifyConnectParam `json:"ClickHouseConnectParam,omitnil,omitempty" name:"ClickHouseConnectParam"`

	// MySQL配置，Type为MYSQL或TDSQL_C_MYSQL时必填
	MySQLConnectParam *MySQLModifyConnectParam `json:"MySQLConnectParam,omitnil,omitempty" name:"MySQLConnectParam"`

	// PostgreSQL配置，Type为POSTGRESQL或TDSQL_C_POSTGRESQL时必填
	PostgreSQLConnectParam *PostgreSQLModifyConnectParam `json:"PostgreSQLConnectParam,omitnil,omitempty" name:"PostgreSQLConnectParam"`

	// MariaDB配置，Type为MARIADB时必填
	MariaDBConnectParam *MariaDBModifyConnectParam `json:"MariaDBConnectParam,omitnil,omitempty" name:"MariaDBConnectParam"`

	// SQLServer配置，Type为SQLSERVER时必填
	SQLServerConnectParam *SQLServerModifyConnectParam `json:"SQLServerConnectParam,omitnil,omitempty" name:"SQLServerConnectParam"`

	// Ctsdb配置，Type为CTSDB
	CtsdbConnectParam *CtsdbModifyConnectParam `json:"CtsdbConnectParam,omitnil,omitempty" name:"CtsdbConnectParam"`

	// Doris配置，Type为DORIS
	DorisConnectParam *DorisModifyConnectParam `json:"DorisConnectParam,omitnil,omitempty" name:"DorisConnectParam"`

	// Kafka配置，Type为 KAFKA 时必填
	KafkaConnectParam *KafkaConnectParam `json:"KafkaConnectParam,omitnil,omitempty" name:"KafkaConnectParam"`

	// MQTT配置，Type为 MQTT 时必填
	MqttConnectParam *MqttConnectParam `json:"MqttConnectParam,omitnil,omitempty" name:"MqttConnectParam"`
}

func (r *ModifyConnectResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConnectResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceId")
	delete(f, "ResourceName")
	delete(f, "Description")
	delete(f, "Type")
	delete(f, "DtsConnectParam")
	delete(f, "MongoDBConnectParam")
	delete(f, "EsConnectParam")
	delete(f, "ClickHouseConnectParam")
	delete(f, "MySQLConnectParam")
	delete(f, "PostgreSQLConnectParam")
	delete(f, "MariaDBConnectParam")
	delete(f, "SQLServerConnectParam")
	delete(f, "CtsdbConnectParam")
	delete(f, "DorisConnectParam")
	delete(f, "KafkaConnectParam")
	delete(f, "MqttConnectParam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyConnectResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyConnectResourceResponseParams struct {
	// 连接源的Id
	Result *ConnectResourceResourceIdResp `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyConnectResourceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyConnectResourceResponseParams `json:"Response"`
}

func (r *ModifyConnectResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConnectResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDatahubTaskRequestParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 任务描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ModifyDatahubTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 任务描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *ModifyDatahubTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDatahubTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "TaskName")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDatahubTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDatahubTaskResponseParams struct {
	// 任务id
	Result *DatahubTaskIdRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDatahubTaskResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDatahubTaskResponseParams `json:"Response"`
}

func (r *ModifyDatahubTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDatahubTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDatahubTopicRequestParams struct {
	// 弹性topic名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 消息保留时间，单位：ms，当前最小值为60000ms。
	RetentionMs *int64 `json:"RetentionMs,omitnil,omitempty" name:"RetentionMs"`

	// 主题备注，是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线-。
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type ModifyDatahubTopicRequest struct {
	*tchttp.BaseRequest
	
	// 弹性topic名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 消息保留时间，单位：ms，当前最小值为60000ms。
	RetentionMs *int64 `json:"RetentionMs,omitnil,omitempty" name:"RetentionMs"`

	// 主题备注，是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线-。
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *ModifyDatahubTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDatahubTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "RetentionMs")
	delete(f, "Note")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDatahubTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDatahubTopicResponseParams struct {
	// 返回结果集
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDatahubTopicResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDatahubTopicResponseParams `json:"Response"`
}

func (r *ModifyDatahubTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDatahubTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGroupOffsetsRequestParams struct {
	// ckafka集群实例Id,可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 消费分组名称,可通过[DescribeConsumerGroup](https://cloud.tencent.com/document/product/597/40841)接口获取
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// 重置offset的策略,入参含义 0. 对齐shift-by参数,代表把offset向前或向后移动shift条 1. 对齐参考(by-duration,to-datetime,to-earliest,to-latest),代表把offset移动到指定timestamp的位置 2. 对齐参考(to-offset),代表把offset移动到指定的offset位置
	Strategy *int64 `json:"Strategy,omitnil,omitempty" name:"Strategy"`

	// 需要重置的主题名列表
	Topics []*string `json:"Topics,omitnil,omitempty" name:"Topics"`

	// 当strategy为0时，必须包含该字段，可以大于零代表会把offset向后移动shift条，小于零则将offset向前回溯shift条数。正确重置后新的offset应该是(old_offset + shift)，需要注意的是如果新的offset小于partition的earliest则会设置为earliest，如果大于partition 的latest则会设置为latest
	Shift *int64 `json:"Shift,omitnil,omitempty" name:"Shift"`

	// 单位ms。当strategy为1时，必须包含该字段，其中-2表示重置offset到最开始的位置，-1表示重置到最新的位置(相当于清空)，其它值则代表指定的时间，会获取topic中指定时间的offset然后进行重置，需要注意的是，如果指定的时间不存在消息，则获取最末尾的offset。
	ShiftTimestamp *int64 `json:"ShiftTimestamp,omitnil,omitempty" name:"ShiftTimestamp"`

	// 需要重新设置的offset位置。当strategy为2，必须包含该字段。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 需要重新设置的partition的列表，如果没有指定Topics参数。则重置全部topics的对应的Partition列表里的partition。指定Topics时则重置指定的topic列表的对应的Partitions列表的partition。
	Partitions []*int64 `json:"Partitions,omitnil,omitempty" name:"Partitions"`
}

type ModifyGroupOffsetsRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id,可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 消费分组名称,可通过[DescribeConsumerGroup](https://cloud.tencent.com/document/product/597/40841)接口获取
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// 重置offset的策略,入参含义 0. 对齐shift-by参数,代表把offset向前或向后移动shift条 1. 对齐参考(by-duration,to-datetime,to-earliest,to-latest),代表把offset移动到指定timestamp的位置 2. 对齐参考(to-offset),代表把offset移动到指定的offset位置
	Strategy *int64 `json:"Strategy,omitnil,omitempty" name:"Strategy"`

	// 需要重置的主题名列表
	Topics []*string `json:"Topics,omitnil,omitempty" name:"Topics"`

	// 当strategy为0时，必须包含该字段，可以大于零代表会把offset向后移动shift条，小于零则将offset向前回溯shift条数。正确重置后新的offset应该是(old_offset + shift)，需要注意的是如果新的offset小于partition的earliest则会设置为earliest，如果大于partition 的latest则会设置为latest
	Shift *int64 `json:"Shift,omitnil,omitempty" name:"Shift"`

	// 单位ms。当strategy为1时，必须包含该字段，其中-2表示重置offset到最开始的位置，-1表示重置到最新的位置(相当于清空)，其它值则代表指定的时间，会获取topic中指定时间的offset然后进行重置，需要注意的是，如果指定的时间不存在消息，则获取最末尾的offset。
	ShiftTimestamp *int64 `json:"ShiftTimestamp,omitnil,omitempty" name:"ShiftTimestamp"`

	// 需要重新设置的offset位置。当strategy为2，必须包含该字段。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 需要重新设置的partition的列表，如果没有指定Topics参数。则重置全部topics的对应的Partition列表里的partition。指定Topics时则重置指定的topic列表的对应的Partitions列表的partition。
	Partitions []*int64 `json:"Partitions,omitnil,omitempty" name:"Partitions"`
}

func (r *ModifyGroupOffsetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGroupOffsetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Group")
	delete(f, "Strategy")
	delete(f, "Topics")
	delete(f, "Shift")
	delete(f, "ShiftTimestamp")
	delete(f, "Offset")
	delete(f, "Partitions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyGroupOffsetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGroupOffsetsResponseParams struct {
	// 返回结果
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyGroupOffsetsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyGroupOffsetsResponseParams `json:"Response"`
}

func (r *ModifyGroupOffsetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGroupOffsetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceAttributesConfig struct {
	// 自动创建 true 表示开启，false 表示不开启
	AutoCreateTopicEnable *bool `json:"AutoCreateTopicEnable,omitnil,omitempty" name:"AutoCreateTopicEnable"`

	// 新创建主题的默认分区数,如果AutoCreateTopicEnable设置为true没有设置该值时，默认设置为3
	DefaultNumPartitions *int64 `json:"DefaultNumPartitions,omitnil,omitempty" name:"DefaultNumPartitions"`

	// 新创建主题的默认副本数,如果AutoCreateTopicEnable设置为true没有指定该值时默认设置为2
	DefaultReplicationFactor *int64 `json:"DefaultReplicationFactor,omitnil,omitempty" name:"DefaultReplicationFactor"`
}

// Predefined struct for user
type ModifyInstanceAttributesRequestParams struct {
	// ckafka集群实例Id,可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例日志的最长保留时间，单位分钟，最大90天，最小为1min
	MsgRetentionTime *int64 `json:"MsgRetentionTime,omitnil,omitempty" name:"MsgRetentionTime"`

	// ckafka集群实例Name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例配置
	Config *ModifyInstanceAttributesConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// 动态消息保留策略配置
	DynamicRetentionConfig *DynamicRetentionTime `json:"DynamicRetentionConfig,omitnil,omitempty" name:"DynamicRetentionConfig"`

	// 用于修改升级版本或升配定时任务的执行时间，Unix时间戳，精确到秒
	RebalanceTime *int64 `json:"RebalanceTime,omitnil,omitempty" name:"RebalanceTime"`

	// 公网带宽 最小3Mbps  最大999Mbps 仅专业版支持填写
	PublicNetwork *int64 `json:"PublicNetwork,omitnil,omitempty" name:"PublicNetwork"`

	// 动态硬盘扩容策略配置
	//
	// Deprecated: DynamicDiskConfig is deprecated.
	DynamicDiskConfig *DynamicDiskConfig `json:"DynamicDiskConfig,omitnil,omitempty" name:"DynamicDiskConfig"`

	// 实例级别单条消息大小（单位byte)  最大 12582912(不包含)  最小1024(不包含)
	MaxMessageByte *uint64 `json:"MaxMessageByte,omitnil,omitempty" name:"MaxMessageByte"`

	// 是否允许未同步的副本选为 leader: 1 开启  0 关闭
	UncleanLeaderElectionEnable *int64 `json:"UncleanLeaderElectionEnable,omitnil,omitempty" name:"UncleanLeaderElectionEnable"`

	// 实例删除保护开关: 1 开启  0 关闭
	DeleteProtectionEnable *int64 `json:"DeleteProtectionEnable,omitnil,omitempty" name:"DeleteProtectionEnable"`
}

type ModifyInstanceAttributesRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id,可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例日志的最长保留时间，单位分钟，最大90天，最小为1min
	MsgRetentionTime *int64 `json:"MsgRetentionTime,omitnil,omitempty" name:"MsgRetentionTime"`

	// ckafka集群实例Name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例配置
	Config *ModifyInstanceAttributesConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// 动态消息保留策略配置
	DynamicRetentionConfig *DynamicRetentionTime `json:"DynamicRetentionConfig,omitnil,omitempty" name:"DynamicRetentionConfig"`

	// 用于修改升级版本或升配定时任务的执行时间，Unix时间戳，精确到秒
	RebalanceTime *int64 `json:"RebalanceTime,omitnil,omitempty" name:"RebalanceTime"`

	// 公网带宽 最小3Mbps  最大999Mbps 仅专业版支持填写
	PublicNetwork *int64 `json:"PublicNetwork,omitnil,omitempty" name:"PublicNetwork"`

	// 动态硬盘扩容策略配置
	DynamicDiskConfig *DynamicDiskConfig `json:"DynamicDiskConfig,omitnil,omitempty" name:"DynamicDiskConfig"`

	// 实例级别单条消息大小（单位byte)  最大 12582912(不包含)  最小1024(不包含)
	MaxMessageByte *uint64 `json:"MaxMessageByte,omitnil,omitempty" name:"MaxMessageByte"`

	// 是否允许未同步的副本选为 leader: 1 开启  0 关闭
	UncleanLeaderElectionEnable *int64 `json:"UncleanLeaderElectionEnable,omitnil,omitempty" name:"UncleanLeaderElectionEnable"`

	// 实例删除保护开关: 1 开启  0 关闭
	DeleteProtectionEnable *int64 `json:"DeleteProtectionEnable,omitnil,omitempty" name:"DeleteProtectionEnable"`
}

func (r *ModifyInstanceAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "MsgRetentionTime")
	delete(f, "InstanceName")
	delete(f, "Config")
	delete(f, "DynamicRetentionConfig")
	delete(f, "RebalanceTime")
	delete(f, "PublicNetwork")
	delete(f, "DynamicDiskConfig")
	delete(f, "MaxMessageByte")
	delete(f, "UncleanLeaderElectionEnable")
	delete(f, "DeleteProtectionEnable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceAttributesResponseParams struct {
	// 返回结果
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceAttributesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceAttributesResponseParams `json:"Response"`
}

func (r *ModifyInstanceAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstancePreRequestParams struct {
	// ckafka集群实例Id,可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 磁盘大小 单位 GB     最大值为500000,步长100
	// 可以通过以下链接查看规格限制：https://cloud.tencent.com/document/product/597/122562
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 峰值带宽 单位 MB/s
	// 可以通过以下链接查看规格限制及对应步长: https://cloud.tencent.com/document/product/597/11745
	BandWidth *int64 `json:"BandWidth,omitnil,omitempty" name:"BandWidth"`

	// 分区上限 最大值: 40000, 步长: 100
	// 可以通过以下链接查看规格限制: https://cloud.tencent.com/document/product/597/122563
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`
}

type ModifyInstancePreRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id,可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 磁盘大小 单位 GB     最大值为500000,步长100
	// 可以通过以下链接查看规格限制：https://cloud.tencent.com/document/product/597/122562
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 峰值带宽 单位 MB/s
	// 可以通过以下链接查看规格限制及对应步长: https://cloud.tencent.com/document/product/597/11745
	BandWidth *int64 `json:"BandWidth,omitnil,omitempty" name:"BandWidth"`

	// 分区上限 最大值: 40000, 步长: 100
	// 可以通过以下链接查看规格限制: https://cloud.tencent.com/document/product/597/122563
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`
}

func (r *ModifyInstancePreRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancePreRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DiskSize")
	delete(f, "BandWidth")
	delete(f, "Partition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstancePreRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstancePreResponseParams struct {
	// 变更预付费实例配置返回结构
	Result *CreateInstancePreResp `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstancePreResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstancePreResponseParams `json:"Response"`
}

func (r *ModifyInstancePreResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancePreResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPasswordRequestParams struct {
	// 实例Id，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名称，可通过[DescribeUser](https://cloud.tencent.com/document/product/597/40855)接口获取。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 用户当前密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 用户新密码
	PasswordNew *string `json:"PasswordNew,omitnil,omitempty" name:"PasswordNew"`
}

type ModifyPasswordRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id，可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名称，可通过[DescribeUser](https://cloud.tencent.com/document/product/597/40855)接口获取。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 用户当前密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 用户新密码
	PasswordNew *string `json:"PasswordNew,omitnil,omitempty" name:"PasswordNew"`
}

func (r *ModifyPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Name")
	delete(f, "Password")
	delete(f, "PasswordNew")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPasswordResponseParams struct {
	// 返回结果
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyPasswordResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPasswordResponseParams `json:"Response"`
}

func (r *ModifyPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRoutineMaintenanceTaskRequestParams struct {
	// ckafka集群实例id,可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 自动化运维类别, 类别如下: QUOTA、ANALYSIS、RE_BALANCE、ELASTIC_BANDWIDTH
	MaintenanceType *string `json:"MaintenanceType,omitnil,omitempty" name:"MaintenanceType"`

	// INSTANCE_STORAGE_CAPACITY(磁盘自动扩容)/MESSAGE_RETENTION_PERIOD(磁盘动态消息保留策略)
	MaintenanceSubtype *string `json:"MaintenanceSubtype,omitnil,omitempty" name:"MaintenanceSubtype"`

	// 主题名
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 任务触发阈值
	ConfigureThreshold *int64 `json:"ConfigureThreshold,omitnil,omitempty" name:"ConfigureThreshold"`

	// 任务调整步长
	ConfigureStepSize *int64 `json:"ConfigureStepSize,omitnil,omitempty" name:"ConfigureStepSize"`

	// 任务调整上限
	ConfigureLimit *int64 `json:"ConfigureLimit,omitnil,omitempty" name:"ConfigureLimit"`

	// 任务预期触发时间，存储从当日 0AM 开始偏移的秒数
	PlannedTime *int64 `json:"PlannedTime,omitnil,omitempty" name:"PlannedTime"`

	// 任务额外信息
	ExtraConfig *string `json:"ExtraConfig,omitnil,omitempty" name:"ExtraConfig"`

	// 任务状态,0 开启,1 关闭
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 执行week day
	Week *string `json:"Week,omitnil,omitempty" name:"Week"`
}

type ModifyRoutineMaintenanceTaskRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例id,可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 自动化运维类别, 类别如下: QUOTA、ANALYSIS、RE_BALANCE、ELASTIC_BANDWIDTH
	MaintenanceType *string `json:"MaintenanceType,omitnil,omitempty" name:"MaintenanceType"`

	// INSTANCE_STORAGE_CAPACITY(磁盘自动扩容)/MESSAGE_RETENTION_PERIOD(磁盘动态消息保留策略)
	MaintenanceSubtype *string `json:"MaintenanceSubtype,omitnil,omitempty" name:"MaintenanceSubtype"`

	// 主题名
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 任务触发阈值
	ConfigureThreshold *int64 `json:"ConfigureThreshold,omitnil,omitempty" name:"ConfigureThreshold"`

	// 任务调整步长
	ConfigureStepSize *int64 `json:"ConfigureStepSize,omitnil,omitempty" name:"ConfigureStepSize"`

	// 任务调整上限
	ConfigureLimit *int64 `json:"ConfigureLimit,omitnil,omitempty" name:"ConfigureLimit"`

	// 任务预期触发时间，存储从当日 0AM 开始偏移的秒数
	PlannedTime *int64 `json:"PlannedTime,omitnil,omitempty" name:"PlannedTime"`

	// 任务额外信息
	ExtraConfig *string `json:"ExtraConfig,omitnil,omitempty" name:"ExtraConfig"`

	// 任务状态,0 开启,1 关闭
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 执行week day
	Week *string `json:"Week,omitnil,omitempty" name:"Week"`
}

func (r *ModifyRoutineMaintenanceTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRoutineMaintenanceTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "MaintenanceType")
	delete(f, "MaintenanceSubtype")
	delete(f, "TopicName")
	delete(f, "ConfigureThreshold")
	delete(f, "ConfigureStepSize")
	delete(f, "ConfigureLimit")
	delete(f, "PlannedTime")
	delete(f, "ExtraConfig")
	delete(f, "Status")
	delete(f, "Week")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRoutineMaintenanceTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRoutineMaintenanceTaskResponseParams struct {
	// 返回结果
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRoutineMaintenanceTaskResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRoutineMaintenanceTaskResponseParams `json:"Response"`
}

func (r *ModifyRoutineMaintenanceTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRoutineMaintenanceTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTopicAttributesRequestParams struct {
	// ckafka集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 主题备注，是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线-。
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// IP 白名单开关，1：打开；0：关闭。
	EnableWhiteList *int64 `json:"EnableWhiteList,omitnil,omitempty" name:"EnableWhiteList"`

	// 默认为1。
	MinInsyncReplicas *int64 `json:"MinInsyncReplicas,omitnil,omitempty" name:"MinInsyncReplicas"`

	// 默认为 0，0：false；1：true。
	UncleanLeaderElectionEnable *int64 `json:"UncleanLeaderElectionEnable,omitnil,omitempty" name:"UncleanLeaderElectionEnable"`

	// 消息保留时间，单位：ms，当前最小值为60000ms。
	RetentionMs *int64 `json:"RetentionMs,omitnil,omitempty" name:"RetentionMs"`

	// 主题消息最大值，单位为 Byte，最大值为12582912Byte（即12MB）。
	MaxMessageBytes *int64 `json:"MaxMessageBytes,omitnil,omitempty" name:"MaxMessageBytes"`

	// Segment 分片滚动的时长，单位：ms，当前最小值86400000ms。
	SegmentMs *int64 `json:"SegmentMs,omitnil,omitempty" name:"SegmentMs"`

	// 消息删除策略，可以选择delete 或者compact
	CleanUpPolicy *string `json:"CleanUpPolicy,omitnil,omitempty" name:"CleanUpPolicy"`

	// Ip白名单列表，配额限制，enableWhileList=1时必选
	IpWhiteList []*string `json:"IpWhiteList,omitnil,omitempty" name:"IpWhiteList"`

	// 预设ACL规则, 1:打开  0:关闭，默认不打开
	EnableAclRule *int64 `json:"EnableAclRule,omitnil,omitempty" name:"EnableAclRule"`

	// ACL规则名
	AclRuleName *string `json:"AclRuleName,omitnil,omitempty" name:"AclRuleName"`

	// 可选, 保留文件大小. 默认为-1,单位bytes, 当前最小值为1048576B
	RetentionBytes *int64 `json:"RetentionBytes,omitnil,omitempty" name:"RetentionBytes"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 生产限流，单位 MB/s；设置为-1，则生产不限流
	QuotaProducerByteRate *int64 `json:"QuotaProducerByteRate,omitnil,omitempty" name:"QuotaProducerByteRate"`

	// 消费限流，单位 MB/s；设置为-1，则消费不限流
	QuotaConsumerByteRate *int64 `json:"QuotaConsumerByteRate,omitnil,omitempty" name:"QuotaConsumerByteRate"`

	// topic副本数  最小值 1,最大值 3
	ReplicaNum *int64 `json:"ReplicaNum,omitnil,omitempty" name:"ReplicaNum"`

	// 消息保存的时间类型：CreateTime/LogAppendTime
	LogMsgTimestampType *string `json:"LogMsgTimestampType,omitnil,omitempty" name:"LogMsgTimestampType"`
}

type ModifyTopicAttributesRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 主题备注，是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线-。
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// IP 白名单开关，1：打开；0：关闭。
	EnableWhiteList *int64 `json:"EnableWhiteList,omitnil,omitempty" name:"EnableWhiteList"`

	// 默认为1。
	MinInsyncReplicas *int64 `json:"MinInsyncReplicas,omitnil,omitempty" name:"MinInsyncReplicas"`

	// 默认为 0，0：false；1：true。
	UncleanLeaderElectionEnable *int64 `json:"UncleanLeaderElectionEnable,omitnil,omitempty" name:"UncleanLeaderElectionEnable"`

	// 消息保留时间，单位：ms，当前最小值为60000ms。
	RetentionMs *int64 `json:"RetentionMs,omitnil,omitempty" name:"RetentionMs"`

	// 主题消息最大值，单位为 Byte，最大值为12582912Byte（即12MB）。
	MaxMessageBytes *int64 `json:"MaxMessageBytes,omitnil,omitempty" name:"MaxMessageBytes"`

	// Segment 分片滚动的时长，单位：ms，当前最小值86400000ms。
	SegmentMs *int64 `json:"SegmentMs,omitnil,omitempty" name:"SegmentMs"`

	// 消息删除策略，可以选择delete 或者compact
	CleanUpPolicy *string `json:"CleanUpPolicy,omitnil,omitempty" name:"CleanUpPolicy"`

	// Ip白名单列表，配额限制，enableWhileList=1时必选
	IpWhiteList []*string `json:"IpWhiteList,omitnil,omitempty" name:"IpWhiteList"`

	// 预设ACL规则, 1:打开  0:关闭，默认不打开
	EnableAclRule *int64 `json:"EnableAclRule,omitnil,omitempty" name:"EnableAclRule"`

	// ACL规则名
	AclRuleName *string `json:"AclRuleName,omitnil,omitempty" name:"AclRuleName"`

	// 可选, 保留文件大小. 默认为-1,单位bytes, 当前最小值为1048576B
	RetentionBytes *int64 `json:"RetentionBytes,omitnil,omitempty" name:"RetentionBytes"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 生产限流，单位 MB/s；设置为-1，则生产不限流
	QuotaProducerByteRate *int64 `json:"QuotaProducerByteRate,omitnil,omitempty" name:"QuotaProducerByteRate"`

	// 消费限流，单位 MB/s；设置为-1，则消费不限流
	QuotaConsumerByteRate *int64 `json:"QuotaConsumerByteRate,omitnil,omitempty" name:"QuotaConsumerByteRate"`

	// topic副本数  最小值 1,最大值 3
	ReplicaNum *int64 `json:"ReplicaNum,omitnil,omitempty" name:"ReplicaNum"`

	// 消息保存的时间类型：CreateTime/LogAppendTime
	LogMsgTimestampType *string `json:"LogMsgTimestampType,omitnil,omitempty" name:"LogMsgTimestampType"`
}

func (r *ModifyTopicAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTopicAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "TopicName")
	delete(f, "Note")
	delete(f, "EnableWhiteList")
	delete(f, "MinInsyncReplicas")
	delete(f, "UncleanLeaderElectionEnable")
	delete(f, "RetentionMs")
	delete(f, "MaxMessageBytes")
	delete(f, "SegmentMs")
	delete(f, "CleanUpPolicy")
	delete(f, "IpWhiteList")
	delete(f, "EnableAclRule")
	delete(f, "AclRuleName")
	delete(f, "RetentionBytes")
	delete(f, "Tags")
	delete(f, "QuotaProducerByteRate")
	delete(f, "QuotaConsumerByteRate")
	delete(f, "ReplicaNum")
	delete(f, "LogMsgTimestampType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTopicAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTopicAttributesResponseParams struct {
	// 返回结果
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTopicAttributesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTopicAttributesResponseParams `json:"Response"`
}

func (r *ModifyTopicAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTopicAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MongoDBConnectParam struct {
	// MongoDB的连接port
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// MongoDB连接源的用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// MongoDB连接源的密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// MongoDB连接源的实例资源
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// MongoDB连接源是否为自建集群
	SelfBuilt *bool `json:"SelfBuilt,omitnil,omitempty" name:"SelfBuilt"`

	// MongoDB连接源的实例vip，当为腾讯云实例时，必填
	ServiceVip *string `json:"ServiceVip,omitnil,omitempty" name:"ServiceVip"`

	// MongoDB连接源的vpcId，当为腾讯云实例时，必填
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 是否更新到关联的Datahub任务
	IsUpdate *bool `json:"IsUpdate,omitnil,omitempty" name:"IsUpdate"`
}

type MongoDBModifyConnectParam struct {
	// MongoDB连接源的实例资源【不支持修改】
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// MongoDB的连接port【不支持修改】
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// MongoDB连接源的实例vip【不支持修改】
	ServiceVip *string `json:"ServiceVip,omitnil,omitempty" name:"ServiceVip"`

	// MongoDB连接源的vpcId【不支持修改】
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// MongoDB连接源的用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// MongoDB连接源的密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// MongoDB连接源是否为自建集群【不支持修改】
	SelfBuilt *bool `json:"SelfBuilt,omitnil,omitempty" name:"SelfBuilt"`

	// 是否更新到关联的Datahub任务
	IsUpdate *bool `json:"IsUpdate,omitnil,omitempty" name:"IsUpdate"`
}

type MongoDBParam struct {
	// MongoDB的数据库名称
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// MongoDB的集群
	Collection *string `json:"Collection,omitnil,omitempty" name:"Collection"`

	// 是否复制存量数据，默认传参true
	CopyExisting *bool `json:"CopyExisting,omitnil,omitempty" name:"CopyExisting"`

	// 实例资源
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// MongoDB的连接ip
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// MongoDB的连接port
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// MongoDB数据库用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// MongoDB数据库密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 监听事件类型，为空时表示全选。取值包括insert,update,replace,delete,invalidate,drop,dropdatabase,rename，多个类型间使用,逗号分隔
	ListeningEvent *string `json:"ListeningEvent,omitnil,omitempty" name:"ListeningEvent"`

	// 主从优先级，默认主节点
	ReadPreference *string `json:"ReadPreference,omitnil,omitempty" name:"ReadPreference"`

	// 聚合管道
	Pipeline *string `json:"Pipeline,omitnil,omitempty" name:"Pipeline"`

	// 是否为自建集群
	SelfBuilt *bool `json:"SelfBuilt,omitnil,omitempty" name:"SelfBuilt"`
}

type MqttConnectParam struct {
	// MQTT的连接port
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// MQTT连接源的用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// MQTT连接源的密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// MQTT连接源的实例资源
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// MQTT的连接ip
	ServiceVip *string `json:"ServiceVip,omitnil,omitempty" name:"ServiceVip"`

	// MQTT Instance vpc-id
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 是否为自建集群
	SelfBuilt *bool `json:"SelfBuilt,omitnil,omitempty" name:"SelfBuilt"`

	// 是否更新到关联的Dip任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsUpdate *bool `json:"IsUpdate,omitnil,omitempty" name:"IsUpdate"`

	// MQTT连接源的实例资源地域, 跨地域时必填
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// IP
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`
}

type MqttParam struct {
	// 需要同步的MQTT Topic列表, CSV格式
	Topics *string `json:"Topics,omitnil,omitempty" name:"Topics"`

	// 用于控制会话的持久性。cleanSession 为true时，连接时会创建一个全新的会话。 cleanSession = false时，连接时会恢复之前的会话。
	CleanSession *bool `json:"CleanSession,omitnil,omitempty" name:"CleanSession"`

	// MQTT instance-id
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// MQTT实例VIP
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// MQTT VIP 端口
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// MQTT实例用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// MQTT实例内账户密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// QoS
	Qos *int64 `json:"Qos,omitnil,omitempty" name:"Qos"`

	// tasks.max 订阅Topic的并发Task个数, 默认为1; 当设置大于1时, 使用Shared Subscription
	MaxTasks *int64 `json:"MaxTasks,omitnil,omitempty" name:"MaxTasks"`

	// MQTT 实例的Service VIP
	ServiceVip *string `json:"ServiceVip,omitnil,omitempty" name:"ServiceVip"`

	// MQTT实例的VPC ID
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 是否为自建集群, MQTT只支持非自建集群
	SelfBuilt *bool `json:"SelfBuilt,omitnil,omitempty" name:"SelfBuilt"`
}

type MySQLConnectParam struct {
	// MySQL的连接port
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// MySQL连接源的用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// MySQL连接源的密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// MySQL连接源的实例资源
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// MySQL连接源的实例vip，当为腾讯云实例时，必填
	ServiceVip *string `json:"ServiceVip,omitnil,omitempty" name:"ServiceVip"`

	// MySQL连接源的vpcId，当为腾讯云实例时，必填
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 是否更新到关联的Datahub任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsUpdate *bool `json:"IsUpdate,omitnil,omitempty" name:"IsUpdate"`

	// 当type为TDSQL_C_MYSQL时，必填
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Mysql 连接源是否为自建集群
	SelfBuilt *bool `json:"SelfBuilt,omitnil,omitempty" name:"SelfBuilt"`
}

type MySQLModifyConnectParam struct {
	// MySQL连接源的实例资源【不支持修改】
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// MySQL的连接port【不支持修改】
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// MySQL连接源的实例vip【不支持修改】
	ServiceVip *string `json:"ServiceVip,omitnil,omitempty" name:"ServiceVip"`

	// MySQL连接源的vpcId【不支持修改】
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// MySQL连接源的用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// MySQL连接源的密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 是否更新到关联的Datahub任务
	IsUpdate *bool `json:"IsUpdate,omitnil,omitempty" name:"IsUpdate"`

	// 当type为TDSQL_C_MYSQL时
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 是否是自建的集群
	SelfBuilt *bool `json:"SelfBuilt,omitnil,omitempty" name:"SelfBuilt"`
}

type MySQLParam struct {
	// MySQL的数据库名称，"*"为全数据库
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// MySQL的数据表名称，"*"为所监听的所有数据库中的非系统表，可以","间隔，监听多个数据表，但数据表需要以"数据库名.数据表名"的格式进行填写，需要填入正则表达式时，格式为"数据库名\\.数据表名"
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// 该MySQL在连接管理内的Id
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// 复制存量信息(schema_only不复制, initial全量)，默认值initial
	SnapshotMode *string `json:"SnapshotMode,omitnil,omitempty" name:"SnapshotMode"`

	// 存放MySQL的Ddl信息的Topic，为空则默认不存放
	DdlTopic *string `json:"DdlTopic,omitnil,omitempty" name:"DdlTopic"`

	// "TABLE" 表示读取项为 table，"QUERY" 表示读取项为 query
	DataSourceMonitorMode *string `json:"DataSourceMonitorMode,omitnil,omitempty" name:"DataSourceMonitorMode"`

	// 当 "DataMonitorMode"="TABLE" 时，传入需要读取的 Table；当 "DataMonitorMode"="QUERY" 时，传入需要读取的查询 sql 语句
	DataSourceMonitorResource *string `json:"DataSourceMonitorResource,omitnil,omitempty" name:"DataSourceMonitorResource"`

	// "TIMESTAMP" 表示增量列为时间戳类型，"INCREMENT" 表示增量列为自增 id 类型
	DataSourceIncrementMode *string `json:"DataSourceIncrementMode,omitnil,omitempty" name:"DataSourceIncrementMode"`

	// 传入需要监听的列名称
	DataSourceIncrementColumn *string `json:"DataSourceIncrementColumn,omitnil,omitempty" name:"DataSourceIncrementColumn"`

	// "HEAD" 表示复制存量 + 增量数据，"TAIL" 表示只复制增量数据
	DataSourceStartFrom *string `json:"DataSourceStartFrom,omitnil,omitempty" name:"DataSourceStartFrom"`

	// "INSERT" 表示使用 Insert 模式插入，"UPSERT" 表示使用 Upsert 模式插入
	DataTargetInsertMode *string `json:"DataTargetInsertMode,omitnil,omitempty" name:"DataTargetInsertMode"`

	// 当 "DataInsertMode"="UPSERT" 时，传入当前 upsert 时依赖的主键
	DataTargetPrimaryKeyField *string `json:"DataTargetPrimaryKeyField,omitnil,omitempty" name:"DataTargetPrimaryKeyField"`

	// 表与消息间的映射关系
	DataTargetRecordMapping []*RecordMapping `json:"DataTargetRecordMapping,omitnil,omitempty" name:"DataTargetRecordMapping"`

	// 事件路由到特定主题的正则表达式，默认为(.*)
	TopicRegex *string `json:"TopicRegex,omitnil,omitempty" name:"TopicRegex"`

	// TopicRegex的引用组，指定$1、$2等
	TopicReplacement *string `json:"TopicReplacement,omitnil,omitempty" name:"TopicReplacement"`

	// 格式：库1.表1:字段1,字段2;库2.表2:字段2，表之间;（分号）隔开，字段之间,（逗号）隔开。不指定的表默认取表的主键
	KeyColumns *string `json:"KeyColumns,omitnil,omitempty" name:"KeyColumns"`

	// Mysql 是否抛弃解析失败的消息，默认为true
	DropInvalidMessage *bool `json:"DropInvalidMessage,omitnil,omitempty" name:"DropInvalidMessage"`

	// 当设置成员参数DropInvalidMessageToCls设置为true时,DropInvalidMessage参数失效
	DropCls *DropCls `json:"DropCls,omitnil,omitempty" name:"DropCls"`

	// 输出格式，DEFAULT、CANAL_1、CANAL_2
	OutputFormat *string `json:"OutputFormat,omitnil,omitempty" name:"OutputFormat"`

	// 当Table输入的是前缀时，该项值为true，否则为false
	IsTablePrefix *bool `json:"IsTablePrefix,omitnil,omitempty" name:"IsTablePrefix"`

	// 如果该值为all，则DDL数据以及DML数据也会写入到选中的topic；若该值为dml，则只有DML数据写入到选中的topic
	IncludeContentChanges *string `json:"IncludeContentChanges,omitnil,omitempty" name:"IncludeContentChanges"`

	// 如果该值为true，且MySQL中"binlog_rows_query_log_events"配置项的值为"ON"，则流入到topic的数据包含原SQL语句；若该值为false，流入到topic的数据不包含原SQL语句
	IncludeQuery *bool `json:"IncludeQuery,omitnil,omitempty" name:"IncludeQuery"`

	// 如果该值为 true，则消息中会携带消息结构体对应的schema，如果该值为false则不会携带
	RecordWithSchema *bool `json:"RecordWithSchema,omitnil,omitempty" name:"RecordWithSchema"`

	// 存放信令表的数据库名称
	SignalDatabase *string `json:"SignalDatabase,omitnil,omitempty" name:"SignalDatabase"`

	// 输入的table是否为正则表达式，如果该选项以及IsTablePrefix同时为true，该选项的判断优先级高于IsTablePrefix
	IsTableRegular *bool `json:"IsTableRegular,omitnil,omitempty" name:"IsTableRegular"`

	// 信号表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SignalTable *string `json:"SignalTable,omitnil,omitempty" name:"SignalTable"`

	// datetime 类型字段转换为时间戳的时区
	DateTimeZone *string `json:"DateTimeZone,omitnil,omitempty" name:"DateTimeZone"`

	// 是否为自建集群
	SelfBuilt *bool `json:"SelfBuilt,omitnil,omitempty" name:"SelfBuilt"`
}

type OperateResponseData struct {
	// 流程Id
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// RouteIdDto
	RouteDTO *RouteDTO `json:"RouteDTO,omitnil,omitempty" name:"RouteDTO"`
}

type Partition struct {
	// 分区ID
	PartitionId *int64 `json:"PartitionId,omitnil,omitempty" name:"PartitionId"`
}

type PartitionOffset struct {
	// 分区
	Partition *string `json:"Partition,omitnil,omitempty" name:"Partition"`

	// 位点偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type Partitions struct {
	// 分区
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// partition 消费位移
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

// Predefined struct for user
type PauseDatahubTaskRequestParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type PauseDatahubTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *PauseDatahubTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PauseDatahubTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PauseDatahubTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PauseDatahubTaskResponseParams struct {
	// 任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *DatahubTaskIdRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PauseDatahubTaskResponse struct {
	*tchttp.BaseResponse
	Response *PauseDatahubTaskResponseParams `json:"Response"`
}

func (r *PauseDatahubTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PauseDatahubTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PostgreSQLConnectParam struct {
	// PostgreSQL的连接port
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// PostgreSQL连接源的用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// PostgreSQL连接源的密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// PostgreSQL连接源的实例资源
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// PostgreSQL连接源的实例vip，当为腾讯云实例时，必填
	ServiceVip *string `json:"ServiceVip,omitnil,omitempty" name:"ServiceVip"`

	// PostgreSQL连接源的vpcId，当为腾讯云实例时，必填
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 当type为TDSQL_C_POSTGRESQL时，必填
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 是否更新到关联的Datahub任务
	IsUpdate *bool `json:"IsUpdate,omitnil,omitempty" name:"IsUpdate"`

	// PostgreSQL连接源是否为自建集群
	SelfBuilt *bool `json:"SelfBuilt,omitnil,omitempty" name:"SelfBuilt"`
}

type PostgreSQLModifyConnectParam struct {
	// PostgreSQL连接源的实例资源【不支持修改】
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// PostgreSQL的连接port【不支持修改】
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// PostgreSQL连接源的实例vip【不支持修改】
	ServiceVip *string `json:"ServiceVip,omitnil,omitempty" name:"ServiceVip"`

	// PostgreSQL连接源的vpcId【不支持修改】
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// PostgreSQL连接源的用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// PostgreSQL连接源的密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 当type为TDSQL_C_POSTGRESQL时，该参数才有值【不支持修改】
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 是否更新到关联的Datahub任务
	IsUpdate *bool `json:"IsUpdate,omitnil,omitempty" name:"IsUpdate"`

	// 是否为自建集群
	SelfBuilt *bool `json:"SelfBuilt,omitnil,omitempty" name:"SelfBuilt"`
}

type PostgreSQLParam struct {
	// PostgreSQL的数据库名称
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// PostgreSQL的数据表名称，"*"为所监听的所有数据库中的非系统表，可以","间隔，监听多个数据表，但数据表需要以"Schema名.数据表名"的格式进行填写，需要填入正则表达式时，格式为"Schema名\\.数据表名"
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// 该PostgreSQL在连接管理内的Id
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// 插件名(decoderbufs/pgoutput)，默认为decoderbufs
	PluginName *string `json:"PluginName,omitnil,omitempty" name:"PluginName"`

	// 复制存量信息(never增量, initial全量)，默认为initial
	SnapshotMode *string `json:"SnapshotMode,omitnil,omitempty" name:"SnapshotMode"`

	// 上游数据格式(JSON/Debezium), 当数据库同步模式为默认字段匹配时,必填
	DataFormat *string `json:"DataFormat,omitnil,omitempty" name:"DataFormat"`

	// "INSERT" 表示使用 Insert 模式插入，"UPSERT" 表示使用 Upsert 模式插入
	DataTargetInsertMode *string `json:"DataTargetInsertMode,omitnil,omitempty" name:"DataTargetInsertMode"`

	// 当 "DataInsertMode"="UPSERT" 时，传入当前 upsert 时依赖的主键
	DataTargetPrimaryKeyField *string `json:"DataTargetPrimaryKeyField,omitnil,omitempty" name:"DataTargetPrimaryKeyField"`

	// 表与消息间的映射关系
	DataTargetRecordMapping []*RecordMapping `json:"DataTargetRecordMapping,omitnil,omitempty" name:"DataTargetRecordMapping"`

	// 是否抛弃解析失败的消息，默认为true
	DropInvalidMessage *bool `json:"DropInvalidMessage,omitnil,omitempty" name:"DropInvalidMessage"`

	// 输入的table是否为正则表达式
	IsTableRegular *bool `json:"IsTableRegular,omitnil,omitempty" name:"IsTableRegular"`

	// 格式：库1.表1:字段1,字段2;库2.表2:字段2，表之间;（分号）隔开，字段之间,（逗号）隔开。不指定的表默认取表的主键
	KeyColumns *string `json:"KeyColumns,omitnil,omitempty" name:"KeyColumns"`

	// 如果该值为 true，则消息中会携带消息结构体对应的schema，如果该值为false则不会携带
	RecordWithSchema *bool `json:"RecordWithSchema,omitnil,omitempty" name:"RecordWithSchema"`
}

type Price struct {
	// 折扣价
	RealTotalCost *float64 `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// 原价
	TotalCost *float64 `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`
}

type PrivateLinkParam struct {
	// 客户实例的vip
	ServiceVip *string `json:"ServiceVip,omitnil,omitempty" name:"ServiceVip"`

	// 客户实例的vpcId
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`
}

type PrometheusDTO struct {
	// export类型（JmxExport\NodeExport）
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// vip
	SourceIp *string `json:"SourceIp,omitnil,omitempty" name:"SourceIp"`

	// vport
	SourcePort *int64 `json:"SourcePort,omitnil,omitempty" name:"SourcePort"`

	// broker地址
	BrokerIp *string `json:"BrokerIp,omitnil,omitempty" name:"BrokerIp"`

	// VPC ID信息
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID信息
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`
}

type PrometheusResult struct {
	// 返回的code，0为正常，非0为错误
	ReturnCode *string `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// 成功消息
	ReturnMessage *string `json:"ReturnMessage,omitnil,omitempty" name:"ReturnMessage"`

	// 操作型返回的Data数据,可能有flowId等
	Data *OperateResponseData `json:"Data,omitnil,omitempty" name:"Data"`
}

type RecordMapping struct {
	// 消息的 key 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	JsonKey *string `json:"JsonKey,omitnil,omitempty" name:"JsonKey"`

	// 消息类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 消息是否允许为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllowNull *bool `json:"AllowNull,omitnil,omitempty" name:"AllowNull"`

	// 对应映射列名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColumnName *string `json:"ColumnName,omitnil,omitempty" name:"ColumnName"`

	// 数据库表额外字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraInfo *string `json:"ExtraInfo,omitnil,omitempty" name:"ExtraInfo"`

	// 当前列大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColumnSize *string `json:"ColumnSize,omitnil,omitempty" name:"ColumnSize"`

	// 当前列精度
	// 注意：此字段可能返回 null，表示取不到有效值。
	DecimalDigits *string `json:"DecimalDigits,omitnil,omitempty" name:"DecimalDigits"`

	// 是否为自增列
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoIncrement *bool `json:"AutoIncrement,omitnil,omitempty" name:"AutoIncrement"`

	// 数据库表默认参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultValue *string `json:"DefaultValue,omitnil,omitempty" name:"DefaultValue"`
}

type RegexReplaceParam struct {
	// 正则表达式
	Regex *string `json:"Regex,omitnil,omitempty" name:"Regex"`

	// 替换新值
	NewValue *string `json:"NewValue,omitnil,omitempty" name:"NewValue"`
}

type Region struct {
	// 地域ID
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 地域名称
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// 区域名称
	AreaName *string `json:"AreaName,omitnil,omitempty" name:"AreaName"`

	// 地域代码
	RegionCode *string `json:"RegionCode,omitnil,omitempty" name:"RegionCode"`

	// 地域代码（V3版本）
	RegionCodeV3 *string `json:"RegionCodeV3,omitnil,omitempty" name:"RegionCodeV3"`

	// NONE:默认值不支持任何特殊类型 实例类型
	Support *string `json:"Support,omitnil,omitempty" name:"Support"`

	// 是否支持ipv6, 0：表示不支持，1：表示支持
	Ipv6 *int64 `json:"Ipv6,omitnil,omitempty" name:"Ipv6"`

	// 是否支持跨可用区, 0：表示不支持，1：表示支持
	MultiZone *int64 `json:"MultiZone,omitnil,omitempty" name:"MultiZone"`
}

// Predefined struct for user
type RenewCkafkaInstanceRequestParams struct {
	// ckafka集群实例Id,可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 续费时长, 默认为1, 单位是月
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`
}

type RenewCkafkaInstanceRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id,可通过[DescribeInstances](https://cloud.tencent.com/document/product/597/40835)接口获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 续费时长, 默认为1, 单位是月
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`
}

func (r *RenewCkafkaInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewCkafkaInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "TimeSpan")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewCkafkaInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RenewCkafkaInstanceResp struct {
	// 订单号
	BigDealId *string `json:"BigDealId,omitnil,omitempty" name:"BigDealId"`

	// 子订单号
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`
}

// Predefined struct for user
type RenewCkafkaInstanceResponseParams struct {
	// 返回值
	Result *RenewCkafkaInstanceResp `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RenewCkafkaInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RenewCkafkaInstanceResponseParams `json:"Response"`
}

func (r *RenewCkafkaInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewCkafkaInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReplaceParam struct {
	// 被替换值
	OldValue *string `json:"OldValue,omitnil,omitempty" name:"OldValue"`

	// 替换值
	NewValue *string `json:"NewValue,omitnil,omitempty" name:"NewValue"`
}

// Predefined struct for user
type RestartDatahubTaskRequestParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type RestartDatahubTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *RestartDatahubTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartDatahubTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartDatahubTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartDatahubTaskResponseParams struct {
	// 任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *DatahubTaskIdRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RestartDatahubTaskResponse struct {
	*tchttp.BaseResponse
	Response *RestartDatahubTaskResponseParams `json:"Response"`
}

func (r *RestartDatahubTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartDatahubTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeDatahubTaskRequestParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type ResumeDatahubTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *ResumeDatahubTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeDatahubTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResumeDatahubTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeDatahubTaskResponseParams struct {
	// 任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *DatahubTaskIdRes `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResumeDatahubTaskResponse struct {
	*tchttp.BaseResponse
	Response *ResumeDatahubTaskResponseParams `json:"Response"`
}

func (r *ResumeDatahubTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeDatahubTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Route struct {
	// 实例接入方式
	// 0：PLAINTEXT (明文方式，没有带用户信息老版本及社区版本都支持)
	// 1：SASL_PLAINTEXT（明文方式，不过在数据开始时，会通过SASL方式登录鉴权，仅社区版本支持）
	// 2：SSL（SSL加密通信，没有带用户信息，老版本及社区版本都支持）
	// 3：SASL_SSL（SSL加密通信，在数据开始时，会通过SASL方式登录鉴权，仅社区版本支持）
	AccessType *int64 `json:"AccessType,omitnil,omitempty" name:"AccessType"`

	// 路由Id
	RouteId *int64 `json:"RouteId,omitnil,omitempty" name:"RouteId"`

	// 路由网络类型(3:vpc路由;7:内部支撑路由;1:公网路由)
	VipType *int64 `json:"VipType,omitnil,omitempty" name:"VipType"`

	// 虚拟IP列表
	VipList []*VipEntity `json:"VipList,omitnil,omitempty" name:"VipList"`

	// 域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名port
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainPort *int64 `json:"DomainPort,omitnil,omitempty" name:"DomainPort"`

	// 时间戳
	DeleteTimestamp *string `json:"DeleteTimestamp,omitnil,omitempty" name:"DeleteTimestamp"`

	// 子网Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Subnet *string `json:"Subnet,omitnil,omitempty" name:"Subnet"`

	// 虚拟IP列表(1对1 broker节点)
	BrokerVipList []*VipEntity `json:"BrokerVipList,omitnil,omitempty" name:"BrokerVipList"`

	// 私有网络Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 备注信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// 路由的状态。1: 创建中，2: 创建成功，3: 创建失败，4: 删除中，6: 删除失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type RouteDTO struct {
	// 路由Id
	RouteId *int64 `json:"RouteId,omitnil,omitempty" name:"RouteId"`
}

type RouteFilter struct {
	// 过滤名称,目前支持security-group-id,按安全组关联过滤
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 过滤值,当过滤名称为security-group-id时仅支持传单个value
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// 过滤关系,支持IN和NOT_IN,默认为IN
	Relation *string `json:"Relation,omitnil,omitempty" name:"Relation"`
}

type RouteResponse struct {
	// 路由信息列表
	Routers []*Route `json:"Routers,omitnil,omitempty" name:"Routers"`
}

type RowParam struct {
	// 行内容，KEY_VALUE，VALUE
	RowContent *string `json:"RowContent,omitnil,omitempty" name:"RowContent"`

	// key和value间的分隔符
	KeyValueDelimiter *string `json:"KeyValueDelimiter,omitnil,omitempty" name:"KeyValueDelimiter"`

	// 元素建的分隔符
	EntryDelimiter *string `json:"EntryDelimiter,omitnil,omitempty" name:"EntryDelimiter"`
}

type SMTParam struct {
	// 数据处理KEY
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 操作，DATE系统预设-时间戳，CUSTOMIZE自定义，MAPPING映射，JSONPATH
	Operate *string `json:"Operate,omitnil,omitempty" name:"Operate"`

	// 数据类型，ORIGINAL原始，STRING，INT64，FLOAT64，BOOLEAN，MAP，ARRAY
	SchemeType *string `json:"SchemeType,omitnil,omitempty" name:"SchemeType"`

	// 数据处理VALUE
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// VALUE处理
	ValueOperate *ValueParam `json:"ValueOperate,omitnil,omitempty" name:"ValueOperate"`

	// 原始VALUE
	OriginalValue *string `json:"OriginalValue,omitnil,omitempty" name:"OriginalValue"`

	// VALUE处理链
	ValueOperates []*ValueParam `json:"ValueOperates,omitnil,omitempty" name:"ValueOperates"`
}

type SQLServerConnectParam struct {
	// SQLServer的连接port
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// SQLServer连接源的用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// SQLServer连接源的密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// SQLServer连接源的实例资源
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// SQLServer连接源的实例vip，当为腾讯云实例时，必填
	ServiceVip *string `json:"ServiceVip,omitnil,omitempty" name:"ServiceVip"`

	// SQLServer连接源的vpcId，当为腾讯云实例时，必填
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 是否更新到关联的Dip任务
	IsUpdate *bool `json:"IsUpdate,omitnil,omitempty" name:"IsUpdate"`
}

type SQLServerModifyConnectParam struct {
	// SQLServer连接源的实例资源【不支持修改】
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// SQLServer的连接port【不支持修改】
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// SQLServer连接源的实例vip【不支持修改】
	ServiceVip *string `json:"ServiceVip,omitnil,omitempty" name:"ServiceVip"`

	// SQLServer连接源的vpcId【不支持修改】
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// SQLServer连接源的用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// SQLServer连接源的密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 是否更新到关联的Dip任务
	IsUpdate *bool `json:"IsUpdate,omitnil,omitempty" name:"IsUpdate"`
}

type SQLServerParam struct {
	// SQLServer的数据库名称
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// SQLServer的数据表名称，"*"为所监听的所有数据库中的非系统表，可以","间隔，监听多个数据表，但数据表需要以"数据库名.数据表名"的格式进行填写
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// 该SQLServer在连接管理内的Id
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// 复制存量信息(schema_only增量, initial全量)，默认为initial
	SnapshotMode *string `json:"SnapshotMode,omitnil,omitempty" name:"SnapshotMode"`
}

type SaleInfo struct {
	// 手动设置的flag标志，true表示售罄，false表示可售。
	Flag *bool `json:"Flag,omitnil,omitempty" name:"Flag"`

	// ckafka版本号(1.1.1/2.4.2/0.10.2)
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 专业版、标准版标志
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 售罄标志：true售罄
	SoldOut *bool `json:"SoldOut,omitnil,omitempty" name:"SoldOut"`
}

type ScalingDownResp struct {
	// 订单号列表
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`
}

type ScfParam struct {
	// SCF云函数函数名
	FunctionName *string `json:"FunctionName,omitnil,omitempty" name:"FunctionName"`

	// SCF云函数命名空间, 默认为default
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// SCF云函数版本及别名, 默认为$DEFAULT
	Qualifier *string `json:"Qualifier,omitnil,omitempty" name:"Qualifier"`

	// 每批最大发送消息数, 默认为1000
	BatchSize *int64 `json:"BatchSize,omitnil,omitempty" name:"BatchSize"`

	// SCF调用失败后重试次数, 默认为5
	MaxRetries *int64 `json:"MaxRetries,omitnil,omitempty" name:"MaxRetries"`
}

type SecondaryAnalyseParam struct {
	// 分隔符
	Regex *string `json:"Regex,omitnil,omitempty" name:"Regex"`
}

type SecurityGroupRoute struct {
	// 路由信息
	InstanceRoute *InstanceRoute `json:"InstanceRoute,omitnil,omitempty" name:"InstanceRoute"`

	// 关联的安全组列表
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// ckafka集群实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 路由vpcId
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 路由vip
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`
}

type SecurityGroupRouteResp struct {
	// 符合条件的安全组路由信息总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 符合条件的安全组路由信息列表
	SecurityGroupRoutes []*SecurityGroupRoute `json:"SecurityGroupRoutes,omitnil,omitempty" name:"SecurityGroupRoutes"`
}

// Predefined struct for user
type SendMessageRequestParams struct {
	// DataHub接入ID
	DataHubId *string `json:"DataHubId,omitnil,omitempty" name:"DataHubId"`

	// 发送消息内容(单次请求最多500条)
	Message []*BatchContent `json:"Message,omitnil,omitempty" name:"Message"`
}

type SendMessageRequest struct {
	*tchttp.BaseRequest
	
	// DataHub接入ID
	DataHubId *string `json:"DataHubId,omitnil,omitempty" name:"DataHubId"`

	// 发送消息内容(单次请求最多500条)
	Message []*BatchContent `json:"Message,omitnil,omitempty" name:"Message"`
}

func (r *SendMessageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendMessageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataHubId")
	delete(f, "Message")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendMessageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendMessageResponseParams struct {
	// 消息ID列表
	MessageId []*string `json:"MessageId,omitnil,omitempty" name:"MessageId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SendMessageResponse struct {
	*tchttp.BaseResponse
	Response *SendMessageResponseParams `json:"Response"`
}

func (r *SendMessageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SplitParam struct {
	// 分隔符
	Regex *string `json:"Regex,omitnil,omitempty" name:"Regex"`
}

type SubscribedInfo struct {
	// 订阅的主题名
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 订阅的分区
	Partition []*int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// 分区offset信息
	PartitionOffset []*PartitionOffset `json:"PartitionOffset,omitnil,omitempty" name:"PartitionOffset"`

	// 订阅的主题ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

type SubstrParam struct {
	// 截取起始位置
	Start *int64 `json:"Start,omitnil,omitempty" name:"Start"`

	// 截取截止位置
	End *int64 `json:"End,omitnil,omitempty" name:"End"`
}

type TableMapping struct {
	// 库名
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// 表名，多个表,（逗号）隔开
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// Topic名称
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// Topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

type Tag struct {
	// 标签的key
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签的值
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TaskStatusResponse struct {
	// 任务状态:
	// 0 成功
	// 1 失败
	// 2 进行中
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 输出信息
	Output *string `json:"Output,omitnil,omitempty" name:"Output"`
}

type TdwParam struct {
	// Tdw的bid
	Bid *string `json:"Bid,omitnil,omitempty" name:"Bid"`

	// Tdw的tid
	Tid *string `json:"Tid,omitnil,omitempty" name:"Tid"`

	// 默认true
	IsDomestic *bool `json:"IsDomestic,omitnil,omitempty" name:"IsDomestic"`

	// TDW地址，默认tl-tdbank-tdmanager.tencent-distribute.com
	TdwHost *string `json:"TdwHost,omitnil,omitempty" name:"TdwHost"`

	// TDW端口，默认8099
	TdwPort *int64 `json:"TdwPort,omitnil,omitempty" name:"TdwPort"`
}

type Topic struct {
	// 主题的ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 主题的名称
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`
}

type TopicAttributesResponse struct {
	// 主题 ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 创建时间的秒级时间戳
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 主题备注
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// 分区个数
	PartitionNum *int64 `json:"PartitionNum,omitnil,omitempty" name:"PartitionNum"`

	// IP 白名单开关，1：打开； 0：关闭
	EnableWhiteList *int64 `json:"EnableWhiteList,omitnil,omitempty" name:"EnableWhiteList"`

	// IP 白名单列表
	IpWhiteList []*string `json:"IpWhiteList,omitnil,omitempty" name:"IpWhiteList"`

	// topic 配置数组
	Config *Config `json:"Config,omitnil,omitempty" name:"Config"`

	// 分区详情
	Partitions []*TopicPartitionDO `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// ACL预设策略开关，1：打开； 0：关闭
	EnableAclRule *int64 `json:"EnableAclRule,omitnil,omitempty" name:"EnableAclRule"`

	// 预设策略列表
	AclRuleList []*AclRule `json:"AclRuleList,omitnil,omitempty" name:"AclRuleList"`

	// topic 限流策略
	QuotaConfig *InstanceQuotaConfigResp `json:"QuotaConfig,omitnil,omitempty" name:"QuotaConfig"`

	// 副本数
	ReplicaNum *int64 `json:"ReplicaNum,omitnil,omitempty" name:"ReplicaNum"`
}

type TopicDetail struct {
	// 主题名
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 主题Id
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 分区数
	PartitionNum *int64 `json:"PartitionNum,omitnil,omitempty" name:"PartitionNum"`

	// topic副本数  最小值 1,最大值 3
	ReplicaNum *int64 `json:"ReplicaNum,omitnil,omitempty" name:"ReplicaNum"`

	// 备注
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// 创建时间
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 是否开启ip鉴权白名单，true表示开启，false表示不开启
	EnableWhiteList *bool `json:"EnableWhiteList,omitnil,omitempty" name:"EnableWhiteList"`

	// ip白名单中ip个数
	IpWhiteListCount *int64 `json:"IpWhiteListCount,omitnil,omitempty" name:"IpWhiteListCount"`

	// 数据备份cos bucket: 转存到cos 的bucket地址
	ForwardCosBucket *string `json:"ForwardCosBucket,omitnil,omitempty" name:"ForwardCosBucket"`

	// 数据备份cos 状态： 1 不开启数据备份，0 开启数据备份
	ForwardStatus *int64 `json:"ForwardStatus,omitnil,omitempty" name:"ForwardStatus"`

	// 数据备份到cos的周期频率
	ForwardInterval *int64 `json:"ForwardInterval,omitnil,omitempty" name:"ForwardInterval"`

	// 高级配置
	Config *Config `json:"Config,omitnil,omitempty" name:"Config"`

	// 消息保留时间配置(用于动态配置变更记录)
	RetentionTimeConfig *TopicRetentionTimeConfigRsp `json:"RetentionTimeConfig,omitnil,omitempty" name:"RetentionTimeConfig"`

	// 0:正常，1：已删除，2：删除中
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type TopicDetailResponse struct {
	// 返回的主题详情列表
	TopicList []*TopicDetail `json:"TopicList,omitnil,omitempty" name:"TopicList"`

	// 符合条件的所有主题详情数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type TopicFlowRanking struct {
	// 主题Id
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 主题名称
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 分区数
	PartitionNum *uint64 `json:"PartitionNum,omitnil,omitempty" name:"PartitionNum"`

	// 副本数
	ReplicaNum *uint64 `json:"ReplicaNum,omitnil,omitempty" name:"ReplicaNum"`

	// Topic 流量,单位MB(设置date时以sum方式聚合)
	TopicTraffic *string `json:"TopicTraffic,omitnil,omitempty" name:"TopicTraffic"`

	// Topic 消息堆积
	MessageHeap *uint64 `json:"MessageHeap,omitnil,omitempty" name:"MessageHeap"`
}

type TopicFlowRankingResult struct {
	// Topic 流量数组
	TopicFlow []*TopicFlowRanking `json:"TopicFlow,omitnil,omitempty" name:"TopicFlow"`

	// 消费者组消费速度排行速度
	ConsumeSpeed []*ConsumerGroupSpeed `json:"ConsumeSpeed,omitnil,omitempty" name:"ConsumeSpeed"`

	// Topic 消息堆积/占用磁盘排行
	TopicMessageHeap []*TopicMessageHeapRanking `json:"TopicMessageHeap,omitnil,omitempty" name:"TopicMessageHeap"`

	// Broker Ip 列表
	BrokerIp []*string `json:"BrokerIp,omitnil,omitempty" name:"BrokerIp"`

	// 单个broker 节点 Topic占用的数据大小
	BrokerTopicData []*BrokerTopicData `json:"BrokerTopicData,omitnil,omitempty" name:"BrokerTopicData"`

	// 单个Broker 节点Topic 流量的大小(单位MB)
	BrokerTopicFlowData []*BrokerTopicFlowData `json:"BrokerTopicFlowData,omitnil,omitempty" name:"BrokerTopicFlowData"`
}

type TopicInSyncReplicaInfo struct {
	// 分区名称
	Partition *string `json:"Partition,omitnil,omitempty" name:"Partition"`

	// Leader Id
	Leader *uint64 `json:"Leader,omitnil,omitempty" name:"Leader"`

	// 副本集
	Replica *string `json:"Replica,omitnil,omitempty" name:"Replica"`

	// ISR
	InSyncReplica *string `json:"InSyncReplica,omitnil,omitempty" name:"InSyncReplica"`

	// 起始Offset
	BeginOffset *uint64 `json:"BeginOffset,omitnil,omitempty" name:"BeginOffset"`

	// 末端Offset
	EndOffset *uint64 `json:"EndOffset,omitnil,omitempty" name:"EndOffset"`

	// 消息数
	MessageCount *uint64 `json:"MessageCount,omitnil,omitempty" name:"MessageCount"`

	// 未同步副本集
	OutOfSyncReplica *string `json:"OutOfSyncReplica,omitnil,omitempty" name:"OutOfSyncReplica"`
}

type TopicInSyncReplicaResult struct {
	// Topic详情及副本合集
	TopicInSyncReplicaList []*TopicInSyncReplicaInfo `json:"TopicInSyncReplicaList,omitnil,omitempty" name:"TopicInSyncReplicaList"`

	// 总计个数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type TopicMessageHeapRanking struct {
	// 主题ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 主题名称
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 分区数
	PartitionNum *uint64 `json:"PartitionNum,omitnil,omitempty" name:"PartitionNum"`

	// 副本数
	ReplicaNum *uint64 `json:"ReplicaNum,omitnil,omitempty" name:"ReplicaNum"`

	// Topic 流量，单位为MB。
	TopicTraffic *string `json:"TopicTraffic,omitnil,omitempty" name:"TopicTraffic"`

	// topic消息堆积/占用磁盘，单位为Bytes。
	MessageHeap *uint64 `json:"MessageHeap,omitnil,omitempty" name:"MessageHeap"`
}

type TopicParam struct {
	// 单独售卖Topic的Topic名称
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// Offset类型，最开始位置earliest，最新位置latest，时间点位置timestamp
	// 注意：此字段可能返回 null，表示取不到有效值。
	OffsetType *string `json:"OffsetType,omitnil,omitempty" name:"OffsetType"`

	// Offset类型为timestamp时必传，传时间戳，精确到秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Topic的TopicId【出参】
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 写入Topic时是否进行压缩，不开启填"none"，开启的话，可选择"gzip", "snappy", "lz4"中的一个进行填写。
	CompressionType *string `json:"CompressionType,omitnil,omitempty" name:"CompressionType"`

	// 使用的Topic是否需要自动创建（目前只支持SOURCE流入任务）
	UseAutoCreateTopic *bool `json:"UseAutoCreateTopic,omitnil,omitempty" name:"UseAutoCreateTopic"`

	// 源topic消息1条扩增成msgMultiple条写入目标topic(该参数目前只有ckafka流入ckafka适用)
	MsgMultiple *int64 `json:"MsgMultiple,omitnil,omitempty" name:"MsgMultiple"`
}

type TopicPartitionDO struct {
	// Partition 分区ID
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// Leader 运行状态，0表示正常运行
	LeaderStatus *int64 `json:"LeaderStatus,omitnil,omitempty" name:"LeaderStatus"`

	// ISR 个数
	IsrNum *int64 `json:"IsrNum,omitnil,omitempty" name:"IsrNum"`

	// 副本个数
	ReplicaNum *int64 `json:"ReplicaNum,omitnil,omitempty" name:"ReplicaNum"`
}

type TopicResult struct {
	// 返回的主题信息列表
	TopicList []*Topic `json:"TopicList,omitnil,omitempty" name:"TopicList"`

	// 符合条件的 topic 数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type TopicRetentionTimeConfigRsp struct {
	// 期望值，即用户配置的Topic消息保留时间(单位分钟)
	Expect *int64 `json:"Expect,omitnil,omitempty" name:"Expect"`

	// 当前值，即当前生效值(可能存在动态调整，单位分钟)
	Current *int64 `json:"Current,omitnil,omitempty" name:"Current"`

	// 最近变更时间
	ModTimeStamp *int64 `json:"ModTimeStamp,omitnil,omitempty" name:"ModTimeStamp"`
}

type TopicSubscribeGroup struct {
	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 消费分组状态数量信息
	StatusCountInfo *string `json:"StatusCountInfo,omitnil,omitempty" name:"StatusCountInfo"`

	// 消费分组信息
	GroupsInfo []*GroupInfoResponse `json:"GroupsInfo,omitnil,omitempty" name:"GroupsInfo"`

	// 此次请求是否异步的状态。实例里分组较少的会直接返回结果,Status为1。当分组较多时,会异步更新缓存，Status为0时不会返回分组信息，直至Status为1更新完毕返回结果。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type TransformParam struct {
	// 解析格式，JSON，DELIMITER分隔符，REGULAR正则提取
	AnalysisFormat *string `json:"AnalysisFormat,omitnil,omitempty" name:"AnalysisFormat"`

	// 输出格式
	OutputFormat *string `json:"OutputFormat,omitnil,omitempty" name:"OutputFormat"`

	// 是否保留解析失败数据
	FailureParam *FailureParam `json:"FailureParam,omitnil,omitempty" name:"FailureParam"`

	// 原始数据
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 数据来源，TOPIC从源topic拉取，CUSTOMIZE自定义
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 分隔符、正则表达式
	Regex *string `json:"Regex,omitnil,omitempty" name:"Regex"`

	// Map
	MapParam []*MapParam `json:"MapParam,omitnil,omitempty" name:"MapParam"`

	// 过滤器
	FilterParam []*FilterMapParam `json:"FilterParam,omitnil,omitempty" name:"FilterParam"`

	// 测试结果
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 解析结果
	AnalyseResult []*MapParam `json:"AnalyseResult,omitnil,omitempty" name:"AnalyseResult"`

	// 底层引擎是否使用eb
	UseEventBus *bool `json:"UseEventBus,omitnil,omitempty" name:"UseEventBus"`
}

type TransformsParam struct {
	// 原始数据
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 处理链
	FieldChain []*FieldParam `json:"FieldChain,omitnil,omitempty" name:"FieldChain"`

	// 过滤器
	FilterParam []*FilterMapParam `json:"FilterParam,omitnil,omitempty" name:"FilterParam"`

	// 失败处理
	FailureParam *FailureParam `json:"FailureParam,omitnil,omitempty" name:"FailureParam"`

	// 测试结果
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 数据来源
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 输出格式，JSON，ROW，默认为JSON
	OutputFormat *string `json:"OutputFormat,omitnil,omitempty" name:"OutputFormat"`

	// 输出格式为ROW必填
	RowParam *RowParam `json:"RowParam,omitnil,omitempty" name:"RowParam"`

	// 是否保留数据源Topic元数据信息（源Topic、Partition、Offset），默认为false
	KeepMetadata *bool `json:"KeepMetadata,omitnil,omitempty" name:"KeepMetadata"`

	// 数组解析
	BatchAnalyse *BatchAnalyseParam `json:"BatchAnalyse,omitnil,omitempty" name:"BatchAnalyse"`
}

// Predefined struct for user
type UpgradeBrokerVersionRequestParams struct {
	// ckafka集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 1.平滑升配.2.垂直升配
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 版本号
	SourceVersion *string `json:"SourceVersion,omitnil,omitempty" name:"SourceVersion"`

	// 版本号
	TargetVersion *string `json:"TargetVersion,omitnil,omitempty" name:"TargetVersion"`

	// 延迟时间
	DelayTimeStamp *string `json:"DelayTimeStamp,omitnil,omitempty" name:"DelayTimeStamp"`
}

type UpgradeBrokerVersionRequest struct {
	*tchttp.BaseRequest
	
	// ckafka集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 1.平滑升配.2.垂直升配
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 版本号
	SourceVersion *string `json:"SourceVersion,omitnil,omitempty" name:"SourceVersion"`

	// 版本号
	TargetVersion *string `json:"TargetVersion,omitnil,omitempty" name:"TargetVersion"`

	// 延迟时间
	DelayTimeStamp *string `json:"DelayTimeStamp,omitnil,omitempty" name:"DelayTimeStamp"`
}

func (r *UpgradeBrokerVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeBrokerVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Type")
	delete(f, "SourceVersion")
	delete(f, "TargetVersion")
	delete(f, "DelayTimeStamp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeBrokerVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeBrokerVersionResponseParams struct {
	// 升配结果
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeBrokerVersionResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeBrokerVersionResponseParams `json:"Response"`
}

func (r *UpgradeBrokerVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeBrokerVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UrlDecodeParam struct {
	// 编码
	CharsetName *string `json:"CharsetName,omitnil,omitempty" name:"CharsetName"`
}

type User struct {
	// 用户id
	UserId *int64 `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 最后更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type UserResponse struct {
	// 符合条件的用户列表
	Users []*User `json:"Users,omitnil,omitempty" name:"Users"`

	// 符合条件的总用户数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type ValueParam struct {
	// 处理模式，REPLACE替换，SUBSTR截取，DATE日期转换，TRIM去除前后空格，REGEX_REPLACE正则替换，URL_DECODE，LOWERCASE转换为小写
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 替换，TYPE=REPLACE时必传
	// 注意：此字段可能返回 null，表示取不到有效值。
	Replace *ReplaceParam `json:"Replace,omitnil,omitempty" name:"Replace"`

	// 截取，TYPE=SUBSTR时必传
	// 注意：此字段可能返回 null，表示取不到有效值。
	Substr *SubstrParam `json:"Substr,omitnil,omitempty" name:"Substr"`

	// 时间转换，TYPE=DATE时必传
	// 注意：此字段可能返回 null，表示取不到有效值。
	Date *DateParam `json:"Date,omitnil,omitempty" name:"Date"`

	// 正则替换，TYPE=REGEX_REPLACE时必传
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegexReplace *RegexReplaceParam `json:"RegexReplace,omitnil,omitempty" name:"RegexReplace"`

	// 值支持一拆多，TYPE=SPLIT时必传
	// 注意：此字段可能返回 null，表示取不到有效值。
	Split *SplitParam `json:"Split,omitnil,omitempty" name:"Split"`

	// key-value二次解析，TYPE=KV时必传
	// 注意：此字段可能返回 null，表示取不到有效值。
	KV *KVParam `json:"KV,omitnil,omitempty" name:"KV"`

	// 处理结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// JsonPath替换，TYPE=JSON_PATH_REPLACE时必传
	// 注意：此字段可能返回 null，表示取不到有效值。
	JsonPathReplace *JsonPathReplaceParam `json:"JsonPathReplace,omitnil,omitempty" name:"JsonPathReplace"`

	// Url解析
	// 注意：此字段可能返回 null，表示取不到有效值。
	UrlDecode *UrlDecodeParam `json:"UrlDecode,omitnil,omitempty" name:"UrlDecode"`

	// 小写字符解析
	// 注意：此字段可能返回 null，表示取不到有效值。
	Lowercase *LowercaseParam `json:"Lowercase,omitnil,omitempty" name:"Lowercase"`
}

type VipEntity struct {
	// 虚拟IP
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 虚拟端口
	Vport *string `json:"Vport,omitnil,omitempty" name:"Vport"`
}

type ZoneInfo struct {
	// 可用区
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 是否内部APP
	IsInternalApp *int64 `json:"IsInternalApp,omitnil,omitempty" name:"IsInternalApp"`

	// 应用标识
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 可用区是否售罄标识，true表示已售罄，false表示未售罄。
	Flag *bool `json:"Flag,omitnil,omitempty" name:"Flag"`

	// 可用区名称
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// 可用区状态  枚举示例:  3: 开启，4: 关闭;  可用区状态以SoldOut为准
	ZoneStatus *int64 `json:"ZoneStatus,omitnil,omitempty" name:"ZoneStatus"`

	// 额外标识
	//
	// Deprecated: Exflag is deprecated.
	Exflag *string `json:"Exflag,omitnil,omitempty" name:"Exflag"`

	// true为售罄，false为未售罄
	SoldOut *string `json:"SoldOut,omitnil,omitempty" name:"SoldOut"`

	// 标准版售罄信息
	SalesInfo []*SaleInfo `json:"SalesInfo,omitnil,omitempty" name:"SalesInfo"`

	// 额外标识
	ExtraFlag *string `json:"ExtraFlag,omitnil,omitempty" name:"ExtraFlag"`
}

type ZoneResponse struct {
	// <p>zone列表</p>
	ZoneList []*ZoneInfo `json:"ZoneList,omitnil,omitempty" name:"ZoneList"`

	// <p>最大购买实例个数</p>
	MaxBuyInstanceNum *int64 `json:"MaxBuyInstanceNum,omitnil,omitempty" name:"MaxBuyInstanceNum"`

	// <p>最大购买带宽 单位Mb/s</p>
	MaxBandwidth *int64 `json:"MaxBandwidth,omitnil,omitempty" name:"MaxBandwidth"`

	// <p>后付费单位价格</p>
	UnitPrice *Price `json:"UnitPrice,omitnil,omitempty" name:"UnitPrice"`

	// <p>后付费消息单价</p>
	MessagePrice *Price `json:"MessagePrice,omitnil,omitempty" name:"MessagePrice"`

	// <p>用户独占集群信息</p>
	ClusterInfo []*ClusterInfo `json:"ClusterInfo,omitnil,omitempty" name:"ClusterInfo"`

	// <p>购买标准版配置</p>
	Standard *string `json:"Standard,omitnil,omitempty" name:"Standard"`

	// <p>购买标准版S2配置</p>
	StandardS2 *string `json:"StandardS2,omitnil,omitempty" name:"StandardS2"`

	// <p>购买专业版配置</p>
	Profession *string `json:"Profession,omitnil,omitempty" name:"Profession"`

	// <p>购买物理独占版配置</p>
	Physical *string `json:"Physical,omitnil,omitempty" name:"Physical"`

	// <p>公网带宽 最小3Mbps  最大999Mbps 仅专业版支持填写   已废弃,无实际意义</p>
	PublicNetwork *string `json:"PublicNetwork,omitnil,omitempty" name:"PublicNetwork"`

	// <p>公网带宽配置</p>
	PublicNetworkLimit *string `json:"PublicNetworkLimit,omitnil,omitempty" name:"PublicNetworkLimit"`

	// <p>请求Id</p>
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`

	// <p>分页offset</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>分页limit</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>是否必须录入tag</p>
	ForceCheckTag *bool `json:"ForceCheckTag,omitnil,omitempty" name:"ForceCheckTag"`
}