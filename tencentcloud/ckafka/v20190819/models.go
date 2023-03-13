// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
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
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Acl struct {
	// Acl资源类型，（0:UNKNOWN，1:ANY，2:TOPIC，3:GROUP，4:CLUSTER，5:TRANSACTIONAL_ID）当前只有TOPIC，
	ResourceType *int64 `json:"ResourceType,omitempty" name:"ResourceType"`

	// 资源名称，和resourceType相关如当resourceType为TOPIC时，则该字段表示topic名称，当resourceType为GROUP时，该字段表示group名称
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// 用户列表，默认为User:*，表示任何user都可以访问，当前用户只能是用户列表中包含的用户
	// 注意：此字段可能返回 null，表示取不到有效值。
	Principal *string `json:"Principal,omitempty" name:"Principal"`

	// 默认为*，表示任何host都可以访问，当前ckafka不支持host为*，但是后面开源kafka的产品化会直接支持
	// 注意：此字段可能返回 null，表示取不到有效值。
	Host *string `json:"Host,omitempty" name:"Host"`

	// Acl操作方式(0:UNKNOWN，1:ANY，2:ALL，3:READ，4:WRITE，5:CREATE，6:DELETE，7:ALTER，8:DESCRIBE，9:CLUSTER_ACTION，10:DESCRIBE_CONFIGS，11:ALTER_CONFIGS，12:IDEMPOTEN_WRITE)
	Operation *int64 `json:"Operation,omitempty" name:"Operation"`

	// 权限类型(0:UNKNOWN，1:ANY，2:DENY，3:ALLOW)
	PermissionType *int64 `json:"PermissionType,omitempty" name:"PermissionType"`
}

type AclResponse struct {
	// 符合条件的总数据条数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// ACL列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AclList []*Acl `json:"AclList,omitempty" name:"AclList"`
}

type AclRule struct {
	// Acl规则名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 匹配类型，目前只支持前缀匹配，枚举值列表：PREFIXED
	// 注意：此字段可能返回 null，表示取不到有效值。
	PatternType *string `json:"PatternType,omitempty" name:"PatternType"`

	// 表示前缀匹配的前缀的值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pattern *string `json:"Pattern,omitempty" name:"Pattern"`

	// Acl资源类型,目前只支持Topic,枚举值列表：Topic
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 该规则所包含的ACL信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AclList *string `json:"AclList,omitempty" name:"AclList"`

	// 规则所创建的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTimeStamp *string `json:"CreateTimeStamp,omitempty" name:"CreateTimeStamp"`

	// 预设ACL规则是否应用到新增的topic中
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsApplied *int64 `json:"IsApplied,omitempty" name:"IsApplied"`

	// 规则更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTimeStamp *string `json:"UpdateTimeStamp,omitempty" name:"UpdateTimeStamp"`

	// 规则的备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 其中一个显示的对应的TopicName
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 应用该ACL规则的Topic数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicCount *int64 `json:"TopicCount,omitempty" name:"TopicCount"`

	// patternType的中文显示
	// 注意：此字段可能返回 null，表示取不到有效值。
	PatternTypeTitle *string `json:"PatternTypeTitle,omitempty" name:"PatternTypeTitle"`
}

type AclRuleInfo struct {
	// Acl操作方式，枚举值(所有操作: All, 读：Read，写：Write)
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 权限类型，(Deny，Allow)
	PermissionType *string `json:"PermissionType,omitempty" name:"PermissionType"`

	// 默认为*，表示任何host都可以访问，当前ckafka不支持host为*和ip网段
	Host *string `json:"Host,omitempty" name:"Host"`

	// 用户列表，默认为User:*，表示任何user都可以访问，当前用户只能是用户列表中包含的用户。传入格式需要带【User:】前缀。例如用户A，传入为User:A。
	Principal *string `json:"Principal,omitempty" name:"Principal"`
}

type AnalyseParam struct {
	// 解析格式，JSON，DELIMITER分隔符，REGULAR正则提取，SOURCE处理上层所有结果
	Format *string `json:"Format,omitempty" name:"Format"`

	// 分隔符、正则表达式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Regex *string `json:"Regex,omitempty" name:"Regex"`

	// 需再次处理的KEY——模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputValueType *string `json:"InputValueType,omitempty" name:"InputValueType"`

	// 需再次处理的KEY——KEY表达式
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputValue *string `json:"InputValue,omitempty" name:"InputValue"`
}

type AppIdResponse struct {
	// 符合要求的所有AppId数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 符合要求的App Id列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppIdList []*int64 `json:"AppIdList,omitempty" name:"AppIdList"`
}

type Assignment struct {
	// assingment版本信息
	Version *int64 `json:"Version,omitempty" name:"Version"`

	// topic信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Topics []*GroupInfoTopics `json:"Topics,omitempty" name:"Topics"`
}

// Predefined struct for user
type AuthorizeTokenRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 用户
	User *string `json:"User,omitempty" name:"User"`

	// token串
	Tokens *string `json:"Tokens,omitempty" name:"Tokens"`
}

type AuthorizeTokenRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 用户
	User *string `json:"User,omitempty" name:"User"`

	// token串
	Tokens *string `json:"Tokens,omitempty" name:"Tokens"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *int64 `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	Format *string `json:"Format,omitempty" name:"Format"`
}

type BatchContent struct {
	// 发送的消息体
	Body *string `json:"Body,omitempty" name:"Body"`

	// 发送消息的键名
	Key *string `json:"Key,omitempty" name:"Key"`
}

// Predefined struct for user
type BatchCreateAclRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Acl资源类型，(2:TOPIC）
	ResourceType *int64 `json:"ResourceType,omitempty" name:"ResourceType"`

	// 资源列表数组
	ResourceNames []*string `json:"ResourceNames,omitempty" name:"ResourceNames"`

	// 设置的ACL规则列表
	RuleList []*AclRuleInfo `json:"RuleList,omitempty" name:"RuleList"`
}

type BatchCreateAclRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Acl资源类型，(2:TOPIC）
	ResourceType *int64 `json:"ResourceType,omitempty" name:"ResourceType"`

	// 资源列表数组
	ResourceNames []*string `json:"ResourceNames,omitempty" name:"ResourceNames"`

	// 设置的ACL规则列表
	RuleList []*AclRuleInfo `json:"RuleList,omitempty" name:"RuleList"`
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
	// 状态码
	Result *int64 `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 实例名称
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// partition信息
	Partitions []*Partitions `json:"Partitions,omitempty" name:"Partitions"`

	// 指定topic，默认所有topic
	TopicName []*string `json:"TopicName,omitempty" name:"TopicName"`
}

type BatchModifyGroupOffsetsRequest struct {
	*tchttp.BaseRequest
	
	// 消费分组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 实例名称
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// partition信息
	Partitions []*Partitions `json:"Partitions,omitempty" name:"Partitions"`

	// 指定topic，默认所有topic
	TopicName []*string `json:"TopicName,omitempty" name:"TopicName"`
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
	Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 主题属性列表
	Topic []*BatchModifyTopicInfo `json:"Topic,omitempty" name:"Topic"`
}

type BatchModifyTopicAttributesRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 主题属性列表
	Topic []*BatchModifyTopicInfo `json:"Topic,omitempty" name:"Topic"`
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
	Result []*BatchModifyTopicResultDTO `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// topic名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 分区数
	PartitionNum *int64 `json:"PartitionNum,omitempty" name:"PartitionNum"`

	// 备注
	Note *string `json:"Note,omitempty" name:"Note"`

	// 副本数
	ReplicaNum *int64 `json:"ReplicaNum,omitempty" name:"ReplicaNum"`

	// 消息删除策略，可以选择delete 或者compact
	CleanUpPolicy *string `json:"CleanUpPolicy,omitempty" name:"CleanUpPolicy"`

	// 当producer设置request.required.acks为-1时，min.insync.replicas指定replicas的最小数目
	MinInsyncReplicas *int64 `json:"MinInsyncReplicas,omitempty" name:"MinInsyncReplicas"`

	// 是否允许非ISR的副本成为Leader
	UncleanLeaderElectionEnable *bool `json:"UncleanLeaderElectionEnable,omitempty" name:"UncleanLeaderElectionEnable"`

	// topic维度的消息保留时间（毫秒）范围1 分钟到90 天
	RetentionMs *int64 `json:"RetentionMs,omitempty" name:"RetentionMs"`

	// topic维度的消息保留大小，范围1 MB到1024 GB
	RetentionBytes *int64 `json:"RetentionBytes,omitempty" name:"RetentionBytes"`

	// Segment分片滚动的时长（毫秒），范围1 到90 天
	SegmentMs *int64 `json:"SegmentMs,omitempty" name:"SegmentMs"`

	// 批次的消息大小，范围1 KB到12 MB
	MaxMessageBytes *int64 `json:"MaxMessageBytes,omitempty" name:"MaxMessageBytes"`
}

type BatchModifyTopicResultDTO struct {
	// 实例id
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// topic名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 状态码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReturnCode *string `json:"ReturnCode,omitempty" name:"ReturnCode"`

	// 状态消息
	Message *string `json:"Message,omitempty" name:"Message"`
}

// Predefined struct for user
type CancelAuthorizationTokenRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 用户
	User *string `json:"User,omitempty" name:"User"`

	// token串
	Tokens *string `json:"Tokens,omitempty" name:"Tokens"`
}

type CancelAuthorizationTokenRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 用户
	User *string `json:"User,omitempty" name:"User"`

	// token串
	Tokens *string `json:"Tokens,omitempty" name:"Tokens"`
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
	// 0 成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *int64 `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

// Predefined struct for user
type CheckCdcClusterRequestParams struct {
	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

type CheckCdcClusterRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// ClickHouse连接源的用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// ClickHouse连接源的密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitempty" name:"Password"`

	// ClickHouse连接源的实例资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// ClickHouse连接源是否为自建集群
	// 注意：此字段可能返回 null，表示取不到有效值。
	SelfBuilt *bool `json:"SelfBuilt,omitempty" name:"SelfBuilt"`

	// ClickHouse连接源的实例vip，当为腾讯云实例时，必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceVip *string `json:"ServiceVip,omitempty" name:"ServiceVip"`

	// ClickHouse连接源的vpcId，当为腾讯云实例时，必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// 是否更新到关联的Datahub任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsUpdate *bool `json:"IsUpdate,omitempty" name:"IsUpdate"`
}

type ClickHouseModifyConnectParam struct {
	// ClickHouse连接源的实例资源【不支持修改】
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// ClickHouse的连接port【不支持修改】
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// ClickHouse连接源的实例vip【不支持修改】
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceVip *string `json:"ServiceVip,omitempty" name:"ServiceVip"`

	// ClickHouse连接源的vpcId【不支持修改】
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// ClickHouse连接源的用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// ClickHouse连接源的密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitempty" name:"Password"`

	// ClickHouse连接源是否为自建集群【不支持修改】
	// 注意：此字段可能返回 null，表示取不到有效值。
	SelfBuilt *bool `json:"SelfBuilt,omitempty" name:"SelfBuilt"`

	// 是否更新到关联的Datahub任务，默认为true
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsUpdate *bool `json:"IsUpdate,omitempty" name:"IsUpdate"`
}

type ClickHouseParam struct {
	// ClickHouse的集群
	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`

	// ClickHouse的数据库名
	Database *string `json:"Database,omitempty" name:"Database"`

	// ClickHouse的数据表名
	Table *string `json:"Table,omitempty" name:"Table"`

	// ClickHouse的schema
	Schema []*ClickHouseSchema `json:"Schema,omitempty" name:"Schema"`

	// 实例资源
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// ClickHouse的连接ip
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// ClickHouse的连接port
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// ClickHouse的用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// ClickHouse的密码
	Password *string `json:"Password,omitempty" name:"Password"`

	// 实例vip
	ServiceVip *string `json:"ServiceVip,omitempty" name:"ServiceVip"`

	// 实例的vpcId
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// 是否为自建集群
	SelfBuilt *bool `json:"SelfBuilt,omitempty" name:"SelfBuilt"`

	// ClickHouse是否抛弃解析失败的消息，默认为true
	DropInvalidMessage *bool `json:"DropInvalidMessage,omitempty" name:"DropInvalidMessage"`

	// ClickHouse 类型，emr-clickhouse : "emr";cdw-clickhouse : "cdwch";自建 : ""
	Type *string `json:"Type,omitempty" name:"Type"`

	// 当设置成员参数DropInvalidMessageToCls设置为true时,DropInvalidMessage参数失效
	DropCls *DropCls `json:"DropCls,omitempty" name:"DropCls"`
}

type ClickHouseSchema struct {
	// 表的列名
	ColumnName *string `json:"ColumnName,omitempty" name:"ColumnName"`

	// 该列对应的jsonKey名
	JsonKey *string `json:"JsonKey,omitempty" name:"JsonKey"`

	// 表列项的类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 列项是否允许为空
	AllowNull *bool `json:"AllowNull,omitempty" name:"AllowNull"`
}

type ClsParam struct {
	// 生产的信息是否为json格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	DecodeJson *bool `json:"DecodeJson,omitempty" name:"DecodeJson"`

	// cls日志主题id
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// cls日志集id
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogSet *string `json:"LogSet,omitempty" name:"LogSet"`

	// 当DecodeJson为false时必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContentKey *string `json:"ContentKey,omitempty" name:"ContentKey"`

	// 指定消息中的某字段内容作为cls日志的时间。
	// 字段内容格式需要是秒级时间戳
	TimeField *string `json:"TimeField,omitempty" name:"TimeField"`
}

type ClusterInfo struct {
	// 集群Id
	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 集群最大磁盘 单位GB
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxDiskSize *int64 `json:"MaxDiskSize,omitempty" name:"MaxDiskSize"`

	// 集群最大带宽 单位MB/s
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxBandWidth *int64 `json:"MaxBandWidth,omitempty" name:"MaxBandWidth"`

	// 集群当前可用磁盘  单位GB
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvailableDiskSize *int64 `json:"AvailableDiskSize,omitempty" name:"AvailableDiskSize"`

	// 集群当前可用带宽 单位MB/s
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvailableBandWidth *int64 `json:"AvailableBandWidth,omitempty" name:"AvailableBandWidth"`

	// 集群所属可用区，表明集群归属的可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 集群节点所在的可用区，若该集群为跨可用区集群，则包含该集群节点所在的多个可用区。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneIds []*int64 `json:"ZoneIds,omitempty" name:"ZoneIds"`
}

type Config struct {
	// 消息保留时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Retention *int64 `json:"Retention,omitempty" name:"Retention"`

	// 最小同步复制数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinInsyncReplicas *int64 `json:"MinInsyncReplicas,omitempty" name:"MinInsyncReplicas"`

	// 日志清理模式，默认 delete。
	// delete：日志按保存时间删除；compact：日志按 key 压缩；compact, delete：日志按 key 压缩且会保存时间删除。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CleanUpPolicy *string `json:"CleanUpPolicy,omitempty" name:"CleanUpPolicy"`

	// Segment 分片滚动的时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	SegmentMs *int64 `json:"SegmentMs,omitempty" name:"SegmentMs"`

	// 0表示 false。 1表示 true。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UncleanLeaderElectionEnable *int64 `json:"UncleanLeaderElectionEnable,omitempty" name:"UncleanLeaderElectionEnable"`

	// Segment 分片滚动的字节数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SegmentBytes *int64 `json:"SegmentBytes,omitempty" name:"SegmentBytes"`

	// 最大消息字节数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxMessageBytes *int64 `json:"MaxMessageBytes,omitempty" name:"MaxMessageBytes"`

	// 消息保留文件大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	RetentionBytes *int64 `json:"RetentionBytes,omitempty" name:"RetentionBytes"`
}

type ConnectResourceResourceIdResp struct {
	// 连接源的Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

type Connection struct {
	// Topic名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 消费组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Topic的Id
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

type ConsumerGroup struct {
	// 用户组名称
	ConsumerGroupName *string `json:"ConsumerGroupName,omitempty" name:"ConsumerGroupName"`

	// 订阅信息实体
	SubscribedInfo []*SubscribedInfo `json:"SubscribedInfo,omitempty" name:"SubscribedInfo"`
}

type ConsumerGroupResponse struct {
	// 符合条件的消费组数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 主题列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicList []*ConsumerGroupTopic `json:"TopicList,omitempty" name:"TopicList"`

	// 消费分组List
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupList []*ConsumerGroup `json:"GroupList,omitempty" name:"GroupList"`

	// 所有分区数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPartition *int64 `json:"TotalPartition,omitempty" name:"TotalPartition"`

	// 监控的分区列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PartitionListForMonitor []*Partition `json:"PartitionListForMonitor,omitempty" name:"PartitionListForMonitor"`

	// 主题总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalTopic *int64 `json:"TotalTopic,omitempty" name:"TotalTopic"`

	// 监控的主题列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicListForMonitor []*ConsumerGroupTopic `json:"TopicListForMonitor,omitempty" name:"TopicListForMonitor"`

	// 监控的组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupListForMonitor []*Group `json:"GroupListForMonitor,omitempty" name:"GroupListForMonitor"`
}

type ConsumerGroupTopic struct {
	// 主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 主题名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

type ConsumerRecord struct {
	// 主题名
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// 分区id
	Partition *int64 `json:"Partition,omitempty" name:"Partition"`

	// 位点
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 消息key
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 消息value
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 消息时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// 消息headers
	// 注意：此字段可能返回 null，表示取不到有效值。
	Headers *string `json:"Headers,omitempty" name:"Headers"`
}

type CosParam struct {
	// cos 存储桶名称
	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`

	// 地域代码
	Region *string `json:"Region,omitempty" name:"Region"`

	// 对象名称
	ObjectKey *string `json:"ObjectKey,omitempty" name:"ObjectKey"`

	// 汇聚消息量的大小（单位：MB)
	AggregateBatchSize *uint64 `json:"AggregateBatchSize,omitempty" name:"AggregateBatchSize"`

	// 汇聚的时间间隔（单位：小时）
	AggregateInterval *uint64 `json:"AggregateInterval,omitempty" name:"AggregateInterval"`

	// 消息汇聚后的文件格式（支持csv, json）
	FormatOutputType *string `json:"FormatOutputType,omitempty" name:"FormatOutputType"`

	// 转储的对象目录前缀
	ObjectKeyPrefix *string `json:"ObjectKeyPrefix,omitempty" name:"ObjectKeyPrefix"`

	// 根据strptime 时间格式化的分区格式
	DirectoryTimeFormat *string `json:"DirectoryTimeFormat,omitempty" name:"DirectoryTimeFormat"`
}

// Predefined struct for user
type CreateAclRequestParams struct {
	// 实例id信息
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Acl资源类型，(2:TOPIC，3:GROUP，4:CLUSTER)
	ResourceType *int64 `json:"ResourceType,omitempty" name:"ResourceType"`

	// Acl操作方式，(2:ALL，3:READ，4:WRITE，5:CREATE，6:DELETE，7:ALTER，8:DESCRIBE，9:CLUSTER_ACTION，10:DESCRIBE_CONFIGS，11:ALTER_CONFIGS，12:IDEMPOTENT_WRITE)
	Operation *int64 `json:"Operation,omitempty" name:"Operation"`

	// 权限类型，(2:DENY，3:ALLOW)，当前ckakfa支持ALLOW(相当于白名单)，其它用于后续兼容开源kafka的acl时使用
	PermissionType *int64 `json:"PermissionType,omitempty" name:"PermissionType"`

	// 资源名称，和resourceType相关，如当resourceType为TOPIC时，则该字段表示topic名称，当resourceType为GROUP时，该字段表示group名称，当resourceType为CLUSTER时，该字段可为空。
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// 默认为\*，表示任何host都可以访问，当前ckafka不支持host为\*，但是后面开源kafka的产品化会直接支持
	Host *string `json:"Host,omitempty" name:"Host"`

	// 用户列表，默认为User:*，表示任何user都可以访问，当前用户只能是用户列表中包含的用户。传入时需要加 User: 前缀,如用户A则传入User:A。
	Principal *string `json:"Principal,omitempty" name:"Principal"`

	// 资源名称列表,Json字符串格式。ResourceName和resourceNameList只能指定其中一个。
	ResourceNameList *string `json:"ResourceNameList,omitempty" name:"ResourceNameList"`
}

type CreateAclRequest struct {
	*tchttp.BaseRequest
	
	// 实例id信息
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Acl资源类型，(2:TOPIC，3:GROUP，4:CLUSTER)
	ResourceType *int64 `json:"ResourceType,omitempty" name:"ResourceType"`

	// Acl操作方式，(2:ALL，3:READ，4:WRITE，5:CREATE，6:DELETE，7:ALTER，8:DESCRIBE，9:CLUSTER_ACTION，10:DESCRIBE_CONFIGS，11:ALTER_CONFIGS，12:IDEMPOTENT_WRITE)
	Operation *int64 `json:"Operation,omitempty" name:"Operation"`

	// 权限类型，(2:DENY，3:ALLOW)，当前ckakfa支持ALLOW(相当于白名单)，其它用于后续兼容开源kafka的acl时使用
	PermissionType *int64 `json:"PermissionType,omitempty" name:"PermissionType"`

	// 资源名称，和resourceType相关，如当resourceType为TOPIC时，则该字段表示topic名称，当resourceType为GROUP时，该字段表示group名称，当resourceType为CLUSTER时，该字段可为空。
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// 默认为\*，表示任何host都可以访问，当前ckafka不支持host为\*，但是后面开源kafka的产品化会直接支持
	Host *string `json:"Host,omitempty" name:"Host"`

	// 用户列表，默认为User:*，表示任何user都可以访问，当前用户只能是用户列表中包含的用户。传入时需要加 User: 前缀,如用户A则传入User:A。
	Principal *string `json:"Principal,omitempty" name:"Principal"`

	// 资源名称列表,Json字符串格式。ResourceName和resourceNameList只能指定其中一个。
	ResourceNameList *string `json:"ResourceNameList,omitempty" name:"ResourceNameList"`
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
	Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type CreateCdcClusterRequestParams struct {
	// cdc的id
	CdcId *string `json:"CdcId,omitempty" name:"CdcId"`

	// vpcId,一个地域只有唯一一个vpcid用于CDC
	CdcVpcId *string `json:"CdcVpcId,omitempty" name:"CdcVpcId"`

	// 每个CDC集群有唯一一个子网ID
	CdcSubnetId *string `json:"CdcSubnetId,omitempty" name:"CdcSubnetId"`

	// 所在可用区ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// cdc集群的总带宽
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// cdc集群的总磁盘
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 数据盘类型
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 系统盘类型
	SystemDiskType *string `json:"SystemDiskType,omitempty" name:"SystemDiskType"`
}

type CreateCdcClusterRequest struct {
	*tchttp.BaseRequest
	
	// cdc的id
	CdcId *string `json:"CdcId,omitempty" name:"CdcId"`

	// vpcId,一个地域只有唯一一个vpcid用于CDC
	CdcVpcId *string `json:"CdcVpcId,omitempty" name:"CdcVpcId"`

	// 每个CDC集群有唯一一个子网ID
	CdcSubnetId *string `json:"CdcSubnetId,omitempty" name:"CdcSubnetId"`

	// 所在可用区ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// cdc集群的总带宽
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// cdc集群的总磁盘
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 数据盘类型
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 系统盘类型
	SystemDiskType *string `json:"SystemDiskType,omitempty" name:"SystemDiskType"`
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
	Result *CdcClusterResponse `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// 连接源类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 连接源描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// Dts配置，Type为DTS时必填
	DtsConnectParam *DtsConnectParam `json:"DtsConnectParam,omitempty" name:"DtsConnectParam"`

	// MongoDB配置，Type为MONGODB时必填
	MongoDBConnectParam *MongoDBConnectParam `json:"MongoDBConnectParam,omitempty" name:"MongoDBConnectParam"`

	// Es配置，Type为ES时必填
	EsConnectParam *EsConnectParam `json:"EsConnectParam,omitempty" name:"EsConnectParam"`

	// ClickHouse配置，Type为CLICKHOUSE时必填
	ClickHouseConnectParam *ClickHouseConnectParam `json:"ClickHouseConnectParam,omitempty" name:"ClickHouseConnectParam"`

	// MySQL配置，Type为MYSQL或TDSQL_C_MYSQL时必填
	MySQLConnectParam *MySQLConnectParam `json:"MySQLConnectParam,omitempty" name:"MySQLConnectParam"`

	// PostgreSQL配置，Type为POSTGRESQL或TDSQL_C_POSTGRESQL时必填
	PostgreSQLConnectParam *PostgreSQLConnectParam `json:"PostgreSQLConnectParam,omitempty" name:"PostgreSQLConnectParam"`

	// MariaDB配置，Type为MARIADB时必填
	MariaDBConnectParam *MariaDBConnectParam `json:"MariaDBConnectParam,omitempty" name:"MariaDBConnectParam"`

	// SQLServer配置，Type为SQLSERVER时必填
	SQLServerConnectParam *SQLServerConnectParam `json:"SQLServerConnectParam,omitempty" name:"SQLServerConnectParam"`

	// Doris 配置，Type为 DORIS 时必填
	DorisConnectParam *DorisConnectParam `json:"DorisConnectParam,omitempty" name:"DorisConnectParam"`
}

type CreateConnectResourceRequest struct {
	*tchttp.BaseRequest
	
	// 连接源名称
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// 连接源类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 连接源描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// Dts配置，Type为DTS时必填
	DtsConnectParam *DtsConnectParam `json:"DtsConnectParam,omitempty" name:"DtsConnectParam"`

	// MongoDB配置，Type为MONGODB时必填
	MongoDBConnectParam *MongoDBConnectParam `json:"MongoDBConnectParam,omitempty" name:"MongoDBConnectParam"`

	// Es配置，Type为ES时必填
	EsConnectParam *EsConnectParam `json:"EsConnectParam,omitempty" name:"EsConnectParam"`

	// ClickHouse配置，Type为CLICKHOUSE时必填
	ClickHouseConnectParam *ClickHouseConnectParam `json:"ClickHouseConnectParam,omitempty" name:"ClickHouseConnectParam"`

	// MySQL配置，Type为MYSQL或TDSQL_C_MYSQL时必填
	MySQLConnectParam *MySQLConnectParam `json:"MySQLConnectParam,omitempty" name:"MySQLConnectParam"`

	// PostgreSQL配置，Type为POSTGRESQL或TDSQL_C_POSTGRESQL时必填
	PostgreSQLConnectParam *PostgreSQLConnectParam `json:"PostgreSQLConnectParam,omitempty" name:"PostgreSQLConnectParam"`

	// MariaDB配置，Type为MARIADB时必填
	MariaDBConnectParam *MariaDBConnectParam `json:"MariaDBConnectParam,omitempty" name:"MariaDBConnectParam"`

	// SQLServer配置，Type为SQLSERVER时必填
	SQLServerConnectParam *SQLServerConnectParam `json:"SQLServerConnectParam,omitempty" name:"SQLServerConnectParam"`

	// Doris 配置，Type为 DORIS 时必填
	DorisConnectParam *DorisConnectParam `json:"DorisConnectParam,omitempty" name:"DorisConnectParam"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateConnectResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConnectResourceResponseParams struct {
	// 连接源的Id
	Result *ConnectResourceResourceIdResp `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// group名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// topic名称，TopicName、TopicNameList 需要显示指定一个存在的topic名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// topic名称数组
	TopicNameList []*string `json:"TopicNameList,omitempty" name:"TopicNameList"`
}

type CreateConsumerRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// group名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// topic名称，TopicName、TopicNameList 需要显示指定一个存在的topic名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// topic名称数组
	TopicNameList []*string `json:"TopicNameList,omitempty" name:"TopicNameList"`
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
	// 创建group描述
	Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 任务类型，SOURCE数据接入，SINK数据流出
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 数据源
	SourceResource *DatahubResource `json:"SourceResource,omitempty" name:"SourceResource"`

	// 数据目标
	TargetResource *DatahubResource `json:"TargetResource,omitempty" name:"TargetResource"`

	// 数据处理规则
	TransformParam *TransformParam `json:"TransformParam,omitempty" name:"TransformParam"`

	// 实例连接参数【已废弃】
	PrivateLinkParam *PrivateLinkParam `json:"PrivateLinkParam,omitempty" name:"PrivateLinkParam"`

	// 选择所要绑定的SchemaId
	SchemaId *string `json:"SchemaId,omitempty" name:"SchemaId"`

	// 数据处理规则
	TransformsParam *TransformsParam `json:"TransformsParam,omitempty" name:"TransformsParam"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type CreateDatahubTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 任务类型，SOURCE数据接入，SINK数据流出
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 数据源
	SourceResource *DatahubResource `json:"SourceResource,omitempty" name:"SourceResource"`

	// 数据目标
	TargetResource *DatahubResource `json:"TargetResource,omitempty" name:"TargetResource"`

	// 数据处理规则
	TransformParam *TransformParam `json:"TransformParam,omitempty" name:"TransformParam"`

	// 实例连接参数【已废弃】
	PrivateLinkParam *PrivateLinkParam `json:"PrivateLinkParam,omitempty" name:"PrivateLinkParam"`

	// 选择所要绑定的SchemaId
	SchemaId *string `json:"SchemaId,omitempty" name:"SchemaId"`

	// 数据处理规则
	TransformsParam *TransformsParam `json:"TransformsParam,omitempty" name:"TransformsParam"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDatahubTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDatahubTaskRes struct {
	// 转储任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 数据转储Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatahubId *string `json:"DatahubId,omitempty" name:"DatahubId"`
}

// Predefined struct for user
type CreateDatahubTaskResponseParams struct {
	// 任务id
	Result *CreateDatahubTaskRes `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type CreateInstancePostRequestParams struct {
	// 实例名称，是一个不超过 64 个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例带宽
	BandWidth *int64 `json:"BandWidth,omitempty" name:"BandWidth"`

	// vpcId，不填默认基础网络
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网id，vpc网络需要传该参数，基础网络可以不传
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 可选。实例日志的最长保留时间，单位分钟，默认为10080（7天），最大30天，不填默认0，代表不开启日志保留时间回收策略
	MsgRetentionTime *int64 `json:"MsgRetentionTime,omitempty" name:"MsgRetentionTime"`

	// 可用区
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 创建实例时可以选择集群Id, 该入参表示集群Id
	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
}

type CreateInstancePostRequest struct {
	*tchttp.BaseRequest
	
	// 实例名称，是一个不超过 64 个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例带宽
	BandWidth *int64 `json:"BandWidth,omitempty" name:"BandWidth"`

	// vpcId，不填默认基础网络
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网id，vpc网络需要传该参数，基础网络可以不传
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 可选。实例日志的最长保留时间，单位分钟，默认为10080（7天），最大30天，不填默认0，代表不开启日志保留时间回收策略
	MsgRetentionTime *int64 `json:"MsgRetentionTime,omitempty" name:"MsgRetentionTime"`

	// 可用区
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 创建实例时可以选择集群Id, 该入参表示集群Id
	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *CreateInstancePostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstancePostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceName")
	delete(f, "BandWidth")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "MsgRetentionTime")
	delete(f, "ZoneId")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstancePostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstancePostResponseParams struct {
	// 返回结果
	Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateInstancePostResponse struct {
	*tchttp.BaseResponse
	Response *CreateInstancePostResponseParams `json:"Response"`
}

func (r *CreateInstancePostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstancePostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstancePreData struct {
	// CreateInstancePre返回固定为0，不能作为CheckTaskStatus的查询条件。只是为了保证和后台数据结构对齐。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 订单号列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`

	// 实例Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

// Predefined struct for user
type CreateInstancePreRequestParams struct {
	// 实例名称，是一个不超过 64 个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 可用区，购买多可用区实例时，填写ZoneIds.N字段中的任意一个值
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 预付费购买时长，例如 "1m",就是一个月
	Period *string `json:"Period,omitempty" name:"Period"`

	// 实例规格说明 专业版实例[所有规格]填写1.
	// 标准版实例 ([入门型(general)]填写1，[标准型(standard)]填写2，[进阶型(advanced)]填写3，[容量型(capacity)]填写4，[高阶型1(specialized-1)]填写5，[高阶性2(specialized-2)]填写6,[高阶型3(specialized-3)]填写7,[高阶型4(specialized-4)]填写8，[独占型(exclusive)]填写9。
	InstanceType *int64 `json:"InstanceType,omitempty" name:"InstanceType"`

	// vpcId，不填默认基础网络
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网id，vpc网络需要传该参数，基础网络可以不传
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 可选。实例日志的最长保留时间，单位分钟，默认为10080（7天），最大30天，不填默认0，代表不开启日志保留时间回收策略
	MsgRetentionTime *int64 `json:"MsgRetentionTime,omitempty" name:"MsgRetentionTime"`

	// 创建实例时可以选择集群Id, 该入参表示集群Id
	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`

	// 预付费自动续费标记，0表示默认状态(用户未设置，即初始状态)， 1表示自动续费，2表示明确不自动续费(用户设置)
	RenewFlag *int64 `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// CKafka版本号[0.10.2、1.1.1、2.4.1], 默认是1.1.1
	KafkaVersion *string `json:"KafkaVersion,omitempty" name:"KafkaVersion"`

	// 实例类型: [标准版实例]填写 standard(默认), [专业版实例]填写 profession
	SpecificationsType *string `json:"SpecificationsType,omitempty" name:"SpecificationsType"`

	// 磁盘大小,专业版不填写默认最小磁盘,填写后根据磁盘带宽分区数弹性计算
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 带宽,专业版不填写默认最小带宽,填写后根据磁盘带宽分区数弹性计算
	BandWidth *int64 `json:"BandWidth,omitempty" name:"BandWidth"`

	// 分区大小,专业版不填写默认最小分区数,填写后根据磁盘带宽分区数弹性计算
	Partition *int64 `json:"Partition,omitempty" name:"Partition"`

	// 标签
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 磁盘类型（ssd填写CLOUD_SSD，sata填写CLOUD_BASIC）
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 跨可用区，zoneIds必填
	MultiZoneFlag *bool `json:"MultiZoneFlag,omitempty" name:"MultiZoneFlag"`

	// 可用区列表，购买多可用区实例时为必填项
	ZoneIds []*int64 `json:"ZoneIds,omitempty" name:"ZoneIds"`
}

type CreateInstancePreRequest struct {
	*tchttp.BaseRequest
	
	// 实例名称，是一个不超过 64 个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 可用区，购买多可用区实例时，填写ZoneIds.N字段中的任意一个值
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 预付费购买时长，例如 "1m",就是一个月
	Period *string `json:"Period,omitempty" name:"Period"`

	// 实例规格说明 专业版实例[所有规格]填写1.
	// 标准版实例 ([入门型(general)]填写1，[标准型(standard)]填写2，[进阶型(advanced)]填写3，[容量型(capacity)]填写4，[高阶型1(specialized-1)]填写5，[高阶性2(specialized-2)]填写6,[高阶型3(specialized-3)]填写7,[高阶型4(specialized-4)]填写8，[独占型(exclusive)]填写9。
	InstanceType *int64 `json:"InstanceType,omitempty" name:"InstanceType"`

	// vpcId，不填默认基础网络
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网id，vpc网络需要传该参数，基础网络可以不传
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 可选。实例日志的最长保留时间，单位分钟，默认为10080（7天），最大30天，不填默认0，代表不开启日志保留时间回收策略
	MsgRetentionTime *int64 `json:"MsgRetentionTime,omitempty" name:"MsgRetentionTime"`

	// 创建实例时可以选择集群Id, 该入参表示集群Id
	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`

	// 预付费自动续费标记，0表示默认状态(用户未设置，即初始状态)， 1表示自动续费，2表示明确不自动续费(用户设置)
	RenewFlag *int64 `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// CKafka版本号[0.10.2、1.1.1、2.4.1], 默认是1.1.1
	KafkaVersion *string `json:"KafkaVersion,omitempty" name:"KafkaVersion"`

	// 实例类型: [标准版实例]填写 standard(默认), [专业版实例]填写 profession
	SpecificationsType *string `json:"SpecificationsType,omitempty" name:"SpecificationsType"`

	// 磁盘大小,专业版不填写默认最小磁盘,填写后根据磁盘带宽分区数弹性计算
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 带宽,专业版不填写默认最小带宽,填写后根据磁盘带宽分区数弹性计算
	BandWidth *int64 `json:"BandWidth,omitempty" name:"BandWidth"`

	// 分区大小,专业版不填写默认最小分区数,填写后根据磁盘带宽分区数弹性计算
	Partition *int64 `json:"Partition,omitempty" name:"Partition"`

	// 标签
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 磁盘类型（ssd填写CLOUD_SSD，sata填写CLOUD_BASIC）
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 跨可用区，zoneIds必填
	MultiZoneFlag *bool `json:"MultiZoneFlag,omitempty" name:"MultiZoneFlag"`

	// 可用区列表，购买多可用区实例时为必填项
	ZoneIds []*int64 `json:"ZoneIds,omitempty" name:"ZoneIds"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstancePreRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstancePreResp struct {
	// 返回的code，0为正常，非0为错误
	ReturnCode *string `json:"ReturnCode,omitempty" name:"ReturnCode"`

	// 成功消息
	ReturnMessage *string `json:"ReturnMessage,omitempty" name:"ReturnMessage"`

	// 操作型返回的Data数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *CreateInstancePreData `json:"Data,omitempty" name:"Data"`

	// 删除是时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteRouteTimestamp *string `json:"DeleteRouteTimestamp,omitempty" name:"DeleteRouteTimestamp"`
}

// Predefined struct for user
type CreateInstancePreResponseParams struct {
	// 返回结果
	Result *CreateInstancePreResp `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 主题名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 主题分区个数
	PartitionNum *int64 `json:"PartitionNum,omitempty" name:"PartitionNum"`
}

type CreatePartitionRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 主题名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 主题分区个数
	PartitionNum *int64 `json:"PartitionNum,omitempty" name:"PartitionNum"`
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
	Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type CreateRouteRequestParams struct {
	// 实例唯一id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 路由网络类型(3:vpc路由;4:标准版支撑路由;7:专业版支撑路由)
	VipType *int64 `json:"VipType,omitempty" name:"VipType"`

	// vpc网络Id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// vpc子网id
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 访问类型
	AccessType *int64 `json:"AccessType,omitempty" name:"AccessType"`

	// 是否需要权限管理
	AuthFlag *int64 `json:"AuthFlag,omitempty" name:"AuthFlag"`

	// 调用方appId
	CallerAppid *int64 `json:"CallerAppid,omitempty" name:"CallerAppid"`

	// 公网带宽
	PublicNetwork *int64 `json:"PublicNetwork,omitempty" name:"PublicNetwork"`

	// vip地址
	Ip *string `json:"Ip,omitempty" name:"Ip"`
}

type CreateRouteRequest struct {
	*tchttp.BaseRequest
	
	// 实例唯一id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 路由网络类型(3:vpc路由;4:标准版支撑路由;7:专业版支撑路由)
	VipType *int64 `json:"VipType,omitempty" name:"VipType"`

	// vpc网络Id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// vpc子网id
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 访问类型
	AccessType *int64 `json:"AccessType,omitempty" name:"AccessType"`

	// 是否需要权限管理
	AuthFlag *int64 `json:"AuthFlag,omitempty" name:"AuthFlag"`

	// 调用方appId
	CallerAppid *int64 `json:"CallerAppid,omitempty" name:"CallerAppid"`

	// 公网带宽
	PublicNetwork *int64 `json:"PublicNetwork,omitempty" name:"PublicNetwork"`

	// vip地址
	Ip *string `json:"Ip,omitempty" name:"Ip"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRouteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRouteResponseParams struct {
	// 返回结果
	Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 用户名
	User *string `json:"User,omitempty" name:"User"`
}

type CreateTokenRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 用户名
	User *string `json:"User,omitempty" name:"User"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 主题名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// ip白名单列表
	IpWhiteList []*string `json:"IpWhiteList,omitempty" name:"IpWhiteList"`
}

type CreateTopicIpWhiteListRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 主题名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// ip白名单列表
	IpWhiteList []*string `json:"IpWhiteList,omitempty" name:"IpWhiteList"`
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
	Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 主题名称，是一个不超过 128 个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Partition个数，大于0
	PartitionNum *int64 `json:"PartitionNum,omitempty" name:"PartitionNum"`

	// 副本个数，不能多于 broker 数，最大为3
	ReplicaNum *int64 `json:"ReplicaNum,omitempty" name:"ReplicaNum"`

	// ip白名单开关, 1:打开  0:关闭，默认不打开
	EnableWhiteList *int64 `json:"EnableWhiteList,omitempty" name:"EnableWhiteList"`

	// Ip白名单列表，配额限制，enableWhileList=1时必选
	IpWhiteList []*string `json:"IpWhiteList,omitempty" name:"IpWhiteList"`

	// 清理日志策略，日志清理模式，默认为"delete"。"delete"：日志按保存时间删除，"compact"：日志按 key 压缩，"compact, delete"：日志按 key 压缩且会按保存时间删除。
	CleanUpPolicy *string `json:"CleanUpPolicy,omitempty" name:"CleanUpPolicy"`

	// 主题备注，是一个不超过 64 个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)
	Note *string `json:"Note,omitempty" name:"Note"`

	// 默认为1
	MinInsyncReplicas *int64 `json:"MinInsyncReplicas,omitempty" name:"MinInsyncReplicas"`

	// 是否允许未同步的副本选为leader，false:不允许，true:允许，默认不允许
	UncleanLeaderElectionEnable *int64 `json:"UncleanLeaderElectionEnable,omitempty" name:"UncleanLeaderElectionEnable"`

	// 可选参数。消息保留时间，单位ms，当前最小值为60000ms
	RetentionMs *int64 `json:"RetentionMs,omitempty" name:"RetentionMs"`

	// Segment分片滚动的时长，单位ms，当前最小为3600000ms
	SegmentMs *int64 `json:"SegmentMs,omitempty" name:"SegmentMs"`

	// 主题消息最大值，单位为 Byte，最小值1024Byte(即1KB)，最大值为8388608Byte（即8MB）。
	MaxMessageBytes *int64 `json:"MaxMessageBytes,omitempty" name:"MaxMessageBytes"`

	// 预设ACL规则, 1:打开  0:关闭，默认不打开
	EnableAclRule *int64 `json:"EnableAclRule,omitempty" name:"EnableAclRule"`

	// 预设ACL规则的名称
	AclRuleName *string `json:"AclRuleName,omitempty" name:"AclRuleName"`

	// 可选, 保留文件大小. 默认为-1,单位bytes, 当前最小值为1048576B
	RetentionBytes *int64 `json:"RetentionBytes,omitempty" name:"RetentionBytes"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type CreateTopicRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 主题名称，是一个不超过 128 个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Partition个数，大于0
	PartitionNum *int64 `json:"PartitionNum,omitempty" name:"PartitionNum"`

	// 副本个数，不能多于 broker 数，最大为3
	ReplicaNum *int64 `json:"ReplicaNum,omitempty" name:"ReplicaNum"`

	// ip白名单开关, 1:打开  0:关闭，默认不打开
	EnableWhiteList *int64 `json:"EnableWhiteList,omitempty" name:"EnableWhiteList"`

	// Ip白名单列表，配额限制，enableWhileList=1时必选
	IpWhiteList []*string `json:"IpWhiteList,omitempty" name:"IpWhiteList"`

	// 清理日志策略，日志清理模式，默认为"delete"。"delete"：日志按保存时间删除，"compact"：日志按 key 压缩，"compact, delete"：日志按 key 压缩且会按保存时间删除。
	CleanUpPolicy *string `json:"CleanUpPolicy,omitempty" name:"CleanUpPolicy"`

	// 主题备注，是一个不超过 64 个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)
	Note *string `json:"Note,omitempty" name:"Note"`

	// 默认为1
	MinInsyncReplicas *int64 `json:"MinInsyncReplicas,omitempty" name:"MinInsyncReplicas"`

	// 是否允许未同步的副本选为leader，false:不允许，true:允许，默认不允许
	UncleanLeaderElectionEnable *int64 `json:"UncleanLeaderElectionEnable,omitempty" name:"UncleanLeaderElectionEnable"`

	// 可选参数。消息保留时间，单位ms，当前最小值为60000ms
	RetentionMs *int64 `json:"RetentionMs,omitempty" name:"RetentionMs"`

	// Segment分片滚动的时长，单位ms，当前最小为3600000ms
	SegmentMs *int64 `json:"SegmentMs,omitempty" name:"SegmentMs"`

	// 主题消息最大值，单位为 Byte，最小值1024Byte(即1KB)，最大值为8388608Byte（即8MB）。
	MaxMessageBytes *int64 `json:"MaxMessageBytes,omitempty" name:"MaxMessageBytes"`

	// 预设ACL规则, 1:打开  0:关闭，默认不打开
	EnableAclRule *int64 `json:"EnableAclRule,omitempty" name:"EnableAclRule"`

	// 预设ACL规则的名称
	AclRuleName *string `json:"AclRuleName,omitempty" name:"AclRuleName"`

	// 可选, 保留文件大小. 默认为-1,单位bytes, 当前最小值为1048576B
	RetentionBytes *int64 `json:"RetentionBytes,omitempty" name:"RetentionBytes"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateTopicResp struct {
	// 主题Id
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

// Predefined struct for user
type CreateTopicResponseParams struct {
	// 返回创建结果
	Result *CreateTopicResp `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 用户名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 用户密码
	Password *string `json:"Password,omitempty" name:"Password"`
}

type CreateUserRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 用户名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 用户密码
	Password *string `json:"Password,omitempty" name:"Password"`
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
	// 返回的结果
	Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Ctsdb连接源的实例vip
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceVip *string `json:"ServiceVip,omitempty" name:"ServiceVip"`

	// Ctsdb连接源的vpcId
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// Ctsdb连接源的用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Ctsdb连接源的密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitempty" name:"Password"`

	// Ctsdb连接源的实例资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resource *string `json:"Resource,omitempty" name:"Resource"`
}

type CtsdbModifyConnectParam struct {
	// Ctsdb的连接port
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Ctsdb连接源的实例vip
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceVip *string `json:"ServiceVip,omitempty" name:"ServiceVip"`

	// Ctsdb连接源的vpcId
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// Ctsdb连接源的用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Ctsdb连接源的密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitempty" name:"Password"`

	// Ctsdb连接源的实例资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resource *string `json:"Resource,omitempty" name:"Resource"`
}

type CtsdbParam struct {
	// 连接管理实例资源
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// Ctsdb的metric
	CtsdbMetric *string `json:"CtsdbMetric,omitempty" name:"CtsdbMetric"`
}

type DatahubResource struct {
	// 资源类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// ckafka配置，Type为KAFKA时必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	KafkaParam *KafkaParam `json:"KafkaParam,omitempty" name:"KafkaParam"`

	// EB配置，Type为EB时必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventBusParam *EventBusParam `json:"EventBusParam,omitempty" name:"EventBusParam"`

	// MongoDB配置，Type为MONGODB时必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	MongoDBParam *MongoDBParam `json:"MongoDBParam,omitempty" name:"MongoDBParam"`

	// Es配置，Type为ES时必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	EsParam *EsParam `json:"EsParam,omitempty" name:"EsParam"`

	// Tdw配置，Type为TDW时必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	TdwParam *TdwParam `json:"TdwParam,omitempty" name:"TdwParam"`

	// Dts配置，Type为DTS时必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	DtsParam *DtsParam `json:"DtsParam,omitempty" name:"DtsParam"`

	// ClickHouse配置，Type为CLICKHOUSE时必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClickHouseParam *ClickHouseParam `json:"ClickHouseParam,omitempty" name:"ClickHouseParam"`

	// Cls配置，Type为CLS时必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClsParam *ClsParam `json:"ClsParam,omitempty" name:"ClsParam"`

	// Cos配置，Type为COS时必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosParam *CosParam `json:"CosParam,omitempty" name:"CosParam"`

	// MySQL配置，Type为MYSQL时必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	MySQLParam *MySQLParam `json:"MySQLParam,omitempty" name:"MySQLParam"`

	// PostgreSQL配置，Type为POSTGRESQL或TDSQL_C_POSTGRESQL时必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	PostgreSQLParam *PostgreSQLParam `json:"PostgreSQLParam,omitempty" name:"PostgreSQLParam"`

	// Topic配置，Type为Topic时必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicParam *TopicParam `json:"TopicParam,omitempty" name:"TopicParam"`

	// MariaDB配置，Type为MARIADB时必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	MariaDBParam *MariaDBParam `json:"MariaDBParam,omitempty" name:"MariaDBParam"`

	// SQLServer配置，Type为SQLSERVER时必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	SQLServerParam *SQLServerParam `json:"SQLServerParam,omitempty" name:"SQLServerParam"`

	// Ctsdb配置，Type为CTSDB时必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	CtsdbParam *CtsdbParam `json:"CtsdbParam,omitempty" name:"CtsdbParam"`

	// Scf配置，Type为SCF时必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScfParam *ScfParam `json:"ScfParam,omitempty" name:"ScfParam"`
}

type DatahubTaskIdRes struct {
	// 任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type DatahubTaskInfo struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 任务类型，SOURCE数据接入，SINK数据流出
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 状态，-1创建失败，0创建中，1运行中，2删除中，3已删除，4删除失败，5暂停中，6已暂停，7暂停失败，8恢复中，9恢复失败
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 数据源
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceResource *DatahubResource `json:"SourceResource,omitempty" name:"SourceResource"`

	// 数据目标
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetResource *DatahubResource `json:"TargetResource,omitempty" name:"TargetResource"`

	// 任务创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 异常信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`

	// 创建进度百分比
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskProgress *float64 `json:"TaskProgress,omitempty" name:"TaskProgress"`

	// 任务当前处于的步骤
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskCurrentStep *string `json:"TaskCurrentStep,omitempty" name:"TaskCurrentStep"`

	// Datahub转储Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatahubId *string `json:"DatahubId,omitempty" name:"DatahubId"`

	// 步骤列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepList []*string `json:"StepList,omitempty" name:"StepList"`
}

type DatahubTopicDTO struct {
	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// Topic名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Topic Id
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 分区数
	PartitionNum *uint64 `json:"PartitionNum,omitempty" name:"PartitionNum"`

	// 过期时间
	RetentionMs *uint64 `json:"RetentionMs,omitempty" name:"RetentionMs"`

	// 备注
	Note *string `json:"Note,omitempty" name:"Note"`

	// 状态，1使用中，2删除中
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type DateParam struct {
	// 时间格式
	Format *string `json:"Format,omitempty" name:"Format"`

	// 输入类型，string，unix时间戳，默认string
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`

	// 时区，默认GMT+8
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeZone *string `json:"TimeZone,omitempty" name:"TimeZone"`
}

// Predefined struct for user
type DeleteAclRequestParams struct {
	// 实例id信息
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Acl资源类型，(2:TOPIC，3:GROUP，4:CLUSTER)
	ResourceType *int64 `json:"ResourceType,omitempty" name:"ResourceType"`

	// 资源名称，和resourceType相关，如当resourceType为TOPIC时，则该字段表示topic名称，当resourceType为GROUP时，该字段表示group名称，当resourceType为CLUSTER时，该字段可为空。
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// Acl操作方式，(2:ALL，3:READ，4:WRITE，5:CREATE，6:DELETE，7:ALTER，8:DESCRIBE，9:CLUSTER_ACTION，10:DESCRIBE_CONFIGS，11:ALTER_CONFIGS，12:IDEMPOTENT_WRITE)
	Operation *int64 `json:"Operation,omitempty" name:"Operation"`

	// 权限类型，(2:DENY，3:ALLOW)，当前ckakfa支持ALLOW(相当于白名单)，其它用于后续兼容开源kafka的acl时使用
	PermissionType *int64 `json:"PermissionType,omitempty" name:"PermissionType"`

	// 默认为\*，表示任何host都可以访问，当前ckafka不支持host为\*，但是后面开源kafka的产品化会直接支持
	Host *string `json:"Host,omitempty" name:"Host"`

	// 用户列表，默认为*，表示任何user都可以访问，当前用户只能是用户列表中包含的用户
	Principal *string `json:"Principal,omitempty" name:"Principal"`
}

type DeleteAclRequest struct {
	*tchttp.BaseRequest
	
	// 实例id信息
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Acl资源类型，(2:TOPIC，3:GROUP，4:CLUSTER)
	ResourceType *int64 `json:"ResourceType,omitempty" name:"ResourceType"`

	// 资源名称，和resourceType相关，如当resourceType为TOPIC时，则该字段表示topic名称，当resourceType为GROUP时，该字段表示group名称，当resourceType为CLUSTER时，该字段可为空。
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// Acl操作方式，(2:ALL，3:READ，4:WRITE，5:CREATE，6:DELETE，7:ALTER，8:DESCRIBE，9:CLUSTER_ACTION，10:DESCRIBE_CONFIGS，11:ALTER_CONFIGS，12:IDEMPOTENT_WRITE)
	Operation *int64 `json:"Operation,omitempty" name:"Operation"`

	// 权限类型，(2:DENY，3:ALLOW)，当前ckakfa支持ALLOW(相当于白名单)，其它用于后续兼容开源kafka的acl时使用
	PermissionType *int64 `json:"PermissionType,omitempty" name:"PermissionType"`

	// 默认为\*，表示任何host都可以访问，当前ckafka不支持host为\*，但是后面开源kafka的产品化会直接支持
	Host *string `json:"Host,omitempty" name:"Host"`

	// 用户列表，默认为*，表示任何user都可以访问，当前用户只能是用户列表中包含的用户
	Principal *string `json:"Principal,omitempty" name:"Principal"`
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
	Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例id信息
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// acl规则名称
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
}

type DeleteAclRuleRequest struct {
	*tchttp.BaseRequest
	
	// 实例id信息
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// acl规则名称
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
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
	Result *int64 `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

type DeleteConnectResourceRequest struct {
	*tchttp.BaseRequest
	
	// 连接源的Id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
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
	Result *ConnectResourceResourceIdResp `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type DeleteDatahubTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
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
	// 任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *DatahubTaskIdRes `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Name *string `json:"Name,omitempty" name:"Name"`
}

type DeleteDatahubTopicRequest struct {
	*tchttp.BaseRequest
	
	// Topic名称
	Name *string `json:"Name,omitempty" name:"Name"`
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
	Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 消费分组
	Group *string `json:"Group,omitempty" name:"Group"`
}

type DeleteGroupRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 消费分组
	Group *string `json:"Group,omitempty" name:"Group"`
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
	Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DeleteInstancePreRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DeleteInstancePreRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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
	Result *CreateInstancePreResp `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例唯一id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 路由id
	RouteId *int64 `json:"RouteId,omitempty" name:"RouteId"`

	// 调用方appId
	CallerAppid *int64 `json:"CallerAppid,omitempty" name:"CallerAppid"`

	// 删除路由时间
	DeleteRouteTime *string `json:"DeleteRouteTime,omitempty" name:"DeleteRouteTime"`
}

type DeleteRouteRequest struct {
	*tchttp.BaseRequest
	
	// 实例唯一id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 路由id
	RouteId *int64 `json:"RouteId,omitempty" name:"RouteId"`

	// 调用方appId
	CallerAppid *int64 `json:"CallerAppid,omitempty" name:"CallerAppid"`

	// 删除路由时间
	DeleteRouteTime *string `json:"DeleteRouteTime,omitempty" name:"DeleteRouteTime"`
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
	Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 修改时间
	DelayTime *string `json:"DelayTime,omitempty" name:"DelayTime"`
}

type DeleteRouteTriggerTimeRequest struct {
	*tchttp.BaseRequest
	
	// 修改时间
	DelayTime *string `json:"DelayTime,omitempty" name:"DelayTime"`
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
	delete(f, "DelayTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRouteTriggerTimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRouteTriggerTimeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 主题名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// ip白名单列表
	IpWhiteList []*string `json:"IpWhiteList,omitempty" name:"IpWhiteList"`
}

type DeleteTopicIpWhiteListRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 主题名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// ip白名单列表
	IpWhiteList []*string `json:"IpWhiteList,omitempty" name:"IpWhiteList"`
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
	Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// ckafka 主题名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

type DeleteTopicRequest struct {
	*tchttp.BaseRequest
	
	// ckafka 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// ckafka 主题名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
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
	Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 用户名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

type DeleteUserRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 用户名称
	Name *string `json:"Name,omitempty" name:"Name"`
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
	Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

// Predefined struct for user
type DescribeACLRequestParams struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Acl资源类型，(2:TOPIC，3:GROUP，4:CLUSTER)
	ResourceType *int64 `json:"ResourceType,omitempty" name:"ResourceType"`

	// 资源名称，和resourceType相关，如当resourceType为TOPIC时，则该字段表示topic名称，当resourceType为GROUP时，该字段表示group名称，当resourceType为CLUSTER时，该字段可为空。
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// 偏移位置
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 个数限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 关键字匹配
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
}

type DescribeACLRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Acl资源类型，(2:TOPIC，3:GROUP，4:CLUSTER)
	ResourceType *int64 `json:"ResourceType,omitempty" name:"ResourceType"`

	// 资源名称，和resourceType相关，如当resourceType为TOPIC时，则该字段表示topic名称，当resourceType为GROUP时，该字段表示group名称，当resourceType为CLUSTER时，该字段可为空。
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// 偏移位置
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 个数限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 关键字匹配
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
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
	Result *AclResponse `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeAppInfoRequestParams struct {
	// 偏移位置
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 本次查询用户数目最大数量限制，最大值为50，默认50
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeAppInfoRequest struct {
	*tchttp.BaseRequest
	
	// 偏移位置
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 本次查询用户数目最大数量限制，最大值为50，默认50
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAppInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAppInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppInfoResponseParams struct {
	// 返回的符合要求的App Id列表
	Result *AppIdResponse `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAppInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAppInfoResponseParams `json:"Response"`
}

func (r *DescribeAppInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCkafkaZoneRequestParams struct {
	// cdc专业集群业务参数
	CdcId *string `json:"CdcId,omitempty" name:"CdcId"`
}

type DescribeCkafkaZoneRequest struct {
	*tchttp.BaseRequest
	
	// cdc专业集群业务参数
	CdcId *string `json:"CdcId,omitempty" name:"CdcId"`
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
	Result *ZoneResponse `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	IpAddr *string `json:"IpAddr,omitempty" name:"IpAddr"`

	// 连结时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Time *string `json:"Time,omitempty" name:"Time"`

	// 是否支持的版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsUnSupportVersion *bool `json:"IsUnSupportVersion,omitempty" name:"IsUnSupportVersion"`
}

type DescribeConnectResource struct {
	// 连接源的Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 连接源名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// 连接源描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 连接源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 连接源的状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 连接源的创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 连接源的异常信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`

	// 连接源的当前所处步骤
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentStep *string `json:"CurrentStep,omitempty" name:"CurrentStep"`

	// 该连接源关联的Datahub任务数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatahubTaskCount *int64 `json:"DatahubTaskCount,omitempty" name:"DatahubTaskCount"`

	// Dts配置，Type为DTS时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	DtsConnectParam *DtsConnectParam `json:"DtsConnectParam,omitempty" name:"DtsConnectParam"`

	// MongoDB配置，Type为MONGODB时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	MongoDBConnectParam *MongoDBConnectParam `json:"MongoDBConnectParam,omitempty" name:"MongoDBConnectParam"`

	// Es配置，Type为ES时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	EsConnectParam *EsConnectParam `json:"EsConnectParam,omitempty" name:"EsConnectParam"`

	// ClickHouse配置，Type为CLICKHOUSE时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClickHouseConnectParam *ClickHouseConnectParam `json:"ClickHouseConnectParam,omitempty" name:"ClickHouseConnectParam"`

	// MySQL配置，Type为MYSQL或TDSQL_C_MYSQL时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	MySQLConnectParam *MySQLConnectParam `json:"MySQLConnectParam,omitempty" name:"MySQLConnectParam"`

	// PostgreSQL配置，Type为POSTGRESQL或TDSQL_C_POSTGRESQL时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	PostgreSQLConnectParam *PostgreSQLConnectParam `json:"PostgreSQLConnectParam,omitempty" name:"PostgreSQLConnectParam"`

	// MariaDB配置，Type为MARIADB时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	MariaDBConnectParam *MariaDBConnectParam `json:"MariaDBConnectParam,omitempty" name:"MariaDBConnectParam"`

	// SQLServer配置，Type为SQLSERVER时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	SQLServerConnectParam *SQLServerConnectParam `json:"SQLServerConnectParam,omitempty" name:"SQLServerConnectParam"`

	// Ctsdb配置，Type为CTSDB时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	CtsdbConnectParam *CtsdbConnectParam `json:"CtsdbConnectParam,omitempty" name:"CtsdbConnectParam"`

	// Doris 配置，Type 为 DORIS 时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	DorisConnectParam *DorisConnectParam `json:"DorisConnectParam,omitempty" name:"DorisConnectParam"`
}

// Predefined struct for user
type DescribeConnectResourceRequestParams struct {
	// 连接源的Id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

type DescribeConnectResourceRequest struct {
	*tchttp.BaseRequest
	
	// 连接源的Id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 连接源名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// 连接源描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 连接源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 连接源的状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 连接源的创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 连接源的异常信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`

	// 连接源的当前所处步骤
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentStep *string `json:"CurrentStep,omitempty" name:"CurrentStep"`

	// 步骤列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepList []*string `json:"StepList,omitempty" name:"StepList"`

	// MySQL配置，Type为MYSQL或TDSQL_C_MYSQL时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	MySQLConnectParam *MySQLConnectParam `json:"MySQLConnectParam,omitempty" name:"MySQLConnectParam"`

	// PostgreSQL配置，Type为POSTGRESQL或TDSQL_C_POSTGRESQL时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	PostgreSQLConnectParam *PostgreSQLConnectParam `json:"PostgreSQLConnectParam,omitempty" name:"PostgreSQLConnectParam"`

	// Dts配置，Type为DTS时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	DtsConnectParam *DtsConnectParam `json:"DtsConnectParam,omitempty" name:"DtsConnectParam"`

	// MongoDB配置，Type为MONGODB时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	MongoDBConnectParam *MongoDBConnectParam `json:"MongoDBConnectParam,omitempty" name:"MongoDBConnectParam"`

	// Es配置，Type为ES时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	EsConnectParam *EsConnectParam `json:"EsConnectParam,omitempty" name:"EsConnectParam"`

	// ClickHouse配置，Type为CLICKHOUSE时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClickHouseConnectParam *ClickHouseConnectParam `json:"ClickHouseConnectParam,omitempty" name:"ClickHouseConnectParam"`

	// MariaDB配置，Type为MARIADB时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	MariaDBConnectParam *MariaDBConnectParam `json:"MariaDBConnectParam,omitempty" name:"MariaDBConnectParam"`

	// SQLServer配置，Type为SQLSERVER时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	SQLServerConnectParam *SQLServerConnectParam `json:"SQLServerConnectParam,omitempty" name:"SQLServerConnectParam"`

	// Ctsdb配置，Type为CTSDB时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	CtsdbConnectParam *CtsdbConnectParam `json:"CtsdbConnectParam,omitempty" name:"CtsdbConnectParam"`

	// Doris 配置，Type 为 DORIS 时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	DorisConnectParam *DorisConnectParam `json:"DorisConnectParam,omitempty" name:"DorisConnectParam"`
}

// Predefined struct for user
type DescribeConnectResourceResponseParams struct {
	// 连接源的Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *DescribeConnectResourceResp `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Type *string `json:"Type,omitempty" name:"Type"`

	// 连接源名称的关键字查询
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 分页偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeConnectResourcesRequest struct {
	*tchttp.BaseRequest
	
	// 连接源类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 连接源名称的关键字查询
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 分页偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConnectResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConnectResourcesResp struct {
	// 连接源个数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 连接源数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConnectResourceList []*DescribeConnectResource `json:"ConnectResourceList,omitempty" name:"ConnectResourceList"`
}

// Predefined struct for user
type DescribeConnectResourcesResponseParams struct {
	// 连接源列表
	Result *DescribeConnectResourcesResp `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// ckafka实例id。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 可选，用户需要查询的group名称。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 可选，用户需要查询的group中的对应的topic名称，如果指定了该参数，而group又未指定则忽略该参数。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 本次返回个数限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移位置
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeConsumerGroupRequest struct {
	*tchttp.BaseRequest
	
	// ckafka实例id。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 可选，用户需要查询的group名称。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 可选，用户需要查询的group中的对应的topic名称，如果指定了该参数，而group又未指定则忽略该参数。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 本次返回个数限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移位置
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
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
	Result *ConsumerGroupResponse `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeDatahubGroupOffsetsRequestParams struct {
	// （过滤条件）按照实例 ID 过滤
	Name *string `json:"Name,omitempty" name:"Name"`

	// Kafka 消费分组
	Group *string `json:"Group,omitempty" name:"Group"`

	// 模糊匹配 topicName
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 本次查询的偏移位置，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 本次返回结果的最大个数，默认为50，最大值为50
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeDatahubGroupOffsetsRequest struct {
	*tchttp.BaseRequest
	
	// （过滤条件）按照实例 ID 过滤
	Name *string `json:"Name,omitempty" name:"Name"`

	// Kafka 消费分组
	Group *string `json:"Group,omitempty" name:"Group"`

	// 模糊匹配 topicName
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 本次查询的偏移位置，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 本次返回结果的最大个数，默认为50，最大值为50
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
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
	Result *GroupOffsetResponse `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeDatahubTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
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
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 任务类型，SOURCE数据接入，SINK数据流出
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 状态，-1创建失败，0创建中，1运行中，2删除中，3已删除，4删除失败，5暂停中，6已暂停，7暂停失败，8恢复中，9恢复失败
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 数据源
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceResource *DatahubResource `json:"SourceResource,omitempty" name:"SourceResource"`

	// 数据目标
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetResource *DatahubResource `json:"TargetResource,omitempty" name:"TargetResource"`

	// Connection列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Connections []*Connection `json:"Connections,omitempty" name:"Connections"`

	// 任务创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 消息处理规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransformParam *TransformParam `json:"TransformParam,omitempty" name:"TransformParam"`

	// 数据接入ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatahubId *string `json:"DatahubId,omitempty" name:"DatahubId"`

	// 绑定的SchemaId
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchemaId *string `json:"SchemaId,omitempty" name:"SchemaId"`

	// 绑定的Schema名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchemaName *string `json:"SchemaName,omitempty" name:"SchemaName"`

	// 数据处理规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransformsParam *TransformsParam `json:"TransformsParam,omitempty" name:"TransformsParam"`

	// 异常信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`

	// 任务标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

// Predefined struct for user
type DescribeDatahubTaskResponseParams struct {
	// 返回结果
	Result *DescribeDatahubTaskRes `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 返回数量，默认为20，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件，按照 TaskName 过滤，支持模糊查询
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 转储的目标类型
	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`

	// 任务类型，SOURCE数据接入，SINK数据流出
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 转储的源类型
	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`

	// 转储的资源
	Resource *string `json:"Resource,omitempty" name:"Resource"`
}

type DescribeDatahubTasksRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为20，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件，按照 TaskName 过滤，支持模糊查询
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 转储的目标类型
	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`

	// 任务类型，SOURCE数据接入，SINK数据流出
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 转储的源类型
	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`

	// 转储的资源
	Resource *string `json:"Resource,omitempty" name:"Resource"`
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
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Datahub任务信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskList []*DatahubTaskInfo `json:"TaskList,omitempty" name:"TaskList"`
}

// Predefined struct for user
type DescribeDatahubTasksResponseParams struct {
	// 返回任务查询结果
	Result *DescribeDatahubTasksRes `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

type DescribeDatahubTopicRequest struct {
	*tchttp.BaseRequest
	
	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`
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
	Name *string `json:"Name,omitempty" name:"Name"`

	// Topic名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Topic Id
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 分区数
	PartitionNum *uint64 `json:"PartitionNum,omitempty" name:"PartitionNum"`

	// 过期时间
	RetentionMs *uint64 `json:"RetentionMs,omitempty" name:"RetentionMs"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Note *string `json:"Note,omitempty" name:"Note"`

	// 用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 密码
	Password *string `json:"Password,omitempty" name:"Password"`

	// 状态，1使用中，2删除中
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 服务路由地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Address *string `json:"Address,omitempty" name:"Address"`
}

// Predefined struct for user
type DescribeDatahubTopicResponseParams struct {
	// 返回的结果对象
	Result *DescribeDatahubTopicResp `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 查询值
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 本次查询的偏移位置，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 本次返回结果的最大个数，默认为50，最大值为50
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeDatahubTopicsRequest struct {
	*tchttp.BaseRequest
	
	// 查询值
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 本次查询的偏移位置，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 本次返回结果的最大个数，默认为50，最大值为50
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatahubTopicsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDatahubTopicsResp struct {
	// 总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Topic列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicList []*DatahubTopicDTO `json:"TopicList,omitempty" name:"TopicList"`
}

// Predefined struct for user
type DescribeDatahubTopicsResponseParams struct {
	// 主题列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *DescribeDatahubTopicsResp `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// groupId
	Group *string `json:"Group,omitempty" name:"Group"`

	// 该 group 使用的协议。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

// Predefined struct for user
type DescribeGroupInfoRequestParams struct {
	// （过滤条件）按照实例 ID 过滤。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Kafka 消费分组，Consumer-group，这里是数组形式，格式：GroupList.0=xxx&GroupList.1=yyy。
	GroupList []*string `json:"GroupList,omitempty" name:"GroupList"`
}

type DescribeGroupInfoRequest struct {
	*tchttp.BaseRequest
	
	// （过滤条件）按照实例 ID 过滤。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Kafka 消费分组，Consumer-group，这里是数组形式，格式：GroupList.0=xxx&GroupList.1=yyy。
	GroupList []*string `json:"GroupList,omitempty" name:"GroupList"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result []*GroupInfoResponse `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// （过滤条件）按照实例 ID 过滤
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Kafka 消费分组
	Group *string `json:"Group,omitempty" name:"Group"`

	// group 订阅的主题名称数组，如果没有该数组，则表示指定的 group 下所有 topic 信息
	Topics []*string `json:"Topics,omitempty" name:"Topics"`

	// 模糊匹配 topicName
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 本次查询的偏移位置，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 本次返回结果的最大个数，默认为50，最大值为50
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeGroupOffsetsRequest struct {
	*tchttp.BaseRequest
	
	// （过滤条件）按照实例 ID 过滤
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Kafka 消费分组
	Group *string `json:"Group,omitempty" name:"Group"`

	// group 订阅的主题名称数组，如果没有该数组，则表示指定的 group 下所有 topic 信息
	Topics []*string `json:"Topics,omitempty" name:"Topics"`

	// 模糊匹配 topicName
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 本次查询的偏移位置，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 本次返回结果的最大个数，默认为50，最大值为50
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
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
	// 返回的结果对象
	Result *GroupOffsetResponse `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 搜索关键字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 最大返回数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeGroupRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 搜索关键字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 最大返回数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGroupResponseParams struct {
	// 返回结果集列表
	Result *GroupResponse `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeInstanceAttributesRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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
	Result *InstanceAttributesResponse `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// （过滤条件）按照实例名,实例Id,可用区,私有网络id,子网id 过滤，支持模糊查询
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// （过滤条件）实例的状态。0：创建中，1：运行中，2：删除中，不填默认返回全部
	Status []*int64 `json:"Status,omitempty" name:"Status"`

	// 偏移量，不填默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，不填则默认10，最大值20。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 匹配标签key值。
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 过滤器。filter.Name 支持('Ip', 'VpcId', 'SubNetId', 'InstanceType','InstanceId') ,filter.Values最多传递10个值.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 已经废弃， 使用InstanceIdList
	InstanceIds *string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 按照实例ID过滤
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList"`

	// 根据标签列表过滤实例（取交集）
	TagList []*Tag `json:"TagList,omitempty" name:"TagList"`
}

type DescribeInstancesDetailRequest struct {
	*tchttp.BaseRequest
	
	// （过滤条件）按照实例ID过滤
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// （过滤条件）按照实例名,实例Id,可用区,私有网络id,子网id 过滤，支持模糊查询
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// （过滤条件）实例的状态。0：创建中，1：运行中，2：删除中，不填默认返回全部
	Status []*int64 `json:"Status,omitempty" name:"Status"`

	// 偏移量，不填默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，不填则默认10，最大值20。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 匹配标签key值。
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 过滤器。filter.Name 支持('Ip', 'VpcId', 'SubNetId', 'InstanceType','InstanceId') ,filter.Values最多传递10个值.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 已经废弃， 使用InstanceIdList
	InstanceIds *string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 按照实例ID过滤
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList"`

	// 根据标签列表过滤实例（取交集）
	TagList []*Tag `json:"TagList,omitempty" name:"TagList"`
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
	Result *InstanceDetailResponse `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// （过滤条件）按照实例ID过滤
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// （过滤条件）按照实例名称过滤，支持模糊查询
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// （过滤条件）实例的状态。0：创建中，1：运行中，2：删除中，不填默认返回全部
	Status []*int64 `json:"Status,omitempty" name:"Status"`

	// 偏移量，不填默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，不填则默认10，最大值100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 已废弃。匹配标签key值。
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 私有网络Id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// （过滤条件）按照实例ID过滤
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// （过滤条件）按照实例名称过滤，支持模糊查询
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// （过滤条件）实例的状态。0：创建中，1：运行中，2：删除中，不填默认返回全部
	Status []*int64 `json:"Status,omitempty" name:"Status"`

	// 偏移量，不填默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，不填则默认10，最大值100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 已废弃。匹配标签key值。
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 私有网络Id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
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
	Result *InstanceResponse `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeRegionRequestParams struct {
	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回最大结果数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 业务字段，可忽略
	Business *string `json:"Business,omitempty" name:"Business"`

	// cdc专有集群业务字段，可忽略
	CdcId *string `json:"CdcId,omitempty" name:"CdcId"`
}

type DescribeRegionRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回最大结果数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 业务字段，可忽略
	Business *string `json:"Business,omitempty" name:"Business"`

	// cdc专有集群业务字段，可忽略
	CdcId *string `json:"CdcId,omitempty" name:"CdcId"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result []*Region `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例唯一id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeRouteRequest struct {
	*tchttp.BaseRequest
	
	// 实例唯一id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRouteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRouteResponseParams struct {
	// 返回的路由信息结果集
	Result *RouteResponse `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeTopicAttributesRequestParams struct {
	// 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 主题名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

type DescribeTopicAttributesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 主题名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
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
	Result *TopicAttributesResponse `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// （过滤条件）按照topicName过滤，支持模糊查询
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 偏移量，不填默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，不填则默认 10，最大值20，取值要大于0
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Acl预设策略名称
	AclRuleName *string `json:"AclRuleName,omitempty" name:"AclRuleName"`
}

type DescribeTopicDetailRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// （过滤条件）按照topicName过滤，支持模糊查询
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 偏移量，不填默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，不填则默认 10，最大值20，取值要大于0
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Acl预设策略名称
	AclRuleName *string `json:"AclRuleName,omitempty" name:"AclRuleName"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopicDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicDetailResponseParams struct {
	// 返回的主题详情实体
	Result *TopicDetailResponse `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeTopicProduceConnectionRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// topic名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

type DescribeTopicProduceConnectionRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// topic名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
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
	Result []*DescribeConnectInfoResultDTO `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 过滤条件，按照 topicName 过滤，支持模糊查询
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 偏移量，不填默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，不填则默认为20，最大值为50
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Acl预设策略名称
	AclRuleName *string `json:"AclRuleName,omitempty" name:"AclRuleName"`
}

type DescribeTopicRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 过滤条件，按照 topicName 过滤，支持模糊查询
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 偏移量，不填默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，不填则默认为20，最大值为50
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Acl预设策略名称
	AclRuleName *string `json:"AclRuleName,omitempty" name:"AclRuleName"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *TopicResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 主题名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 分页时的起始位置
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页时的个数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeTopicSubscribeGroupRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 主题名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 分页时的起始位置
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页时的个数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
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
	Result *TopicSubscribeGroup `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 主题名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 偏移量，不填默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，不填则默认10，最大值20。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 仅筛选未同步副本
	OutOfSyncReplicaOnly *bool `json:"OutOfSyncReplicaOnly,omitempty" name:"OutOfSyncReplicaOnly"`
}

type DescribeTopicSyncReplicaRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 主题名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 偏移量，不填默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，不填则默认10，最大值20。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 仅筛选未同步副本
	OutOfSyncReplicaOnly *bool `json:"OutOfSyncReplicaOnly,omitempty" name:"OutOfSyncReplicaOnly"`
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
	Result *TopicInSyncReplicaResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeUserRequestParams struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 按照名称过滤
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 本次返回个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeUserRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 按照名称过滤
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 本次返回个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
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
	// 返回结果列表
	Result *UserResponse `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Doris 连接源的用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Doris 连接源的密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitempty" name:"Password"`

	// Doris 连接源的实例资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// Doris 连接源的实例vip，当为腾讯云实例时，必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceVip *string `json:"ServiceVip,omitempty" name:"ServiceVip"`

	// Doris 连接源的vpcId，当为腾讯云实例时，必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// 是否更新到关联的Datahub任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsUpdate *bool `json:"IsUpdate,omitempty" name:"IsUpdate"`

	// Doris 连接源是否为自建集群
	// 注意：此字段可能返回 null，表示取不到有效值。
	SelfBuilt *bool `json:"SelfBuilt,omitempty" name:"SelfBuilt"`

	// Doris 的 http 负载均衡连接 port，通常映射到 be 的 8040 端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	BePort *int64 `json:"BePort,omitempty" name:"BePort"`
}

type DorisModifyConnectParam struct {
	// Doris 连接源的实例资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// Doris jdbc 负载均衡连接 port，通常映射到 fe 的 9030 端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Doris 连接源的实例vip，当为腾讯云实例时，必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceVip *string `json:"ServiceVip,omitempty" name:"ServiceVip"`

	// Doris 连接源的vpcId，当为腾讯云实例时，必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// Doris 连接源的用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Doris 连接源的密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 是否更新到关联的Datahub任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsUpdate *bool `json:"IsUpdate,omitempty" name:"IsUpdate"`

	// Doris 连接源是否为自建集群
	// 注意：此字段可能返回 null，表示取不到有效值。
	SelfBuilt *bool `json:"SelfBuilt,omitempty" name:"SelfBuilt"`

	// Doris 的 http 负载均衡连接 port，通常映射到 be 的 8040 端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	BePort *int64 `json:"BePort,omitempty" name:"BePort"`
}

type DropCls struct {
	// 是否投递到cls
	// 注意：此字段可能返回 null，表示取不到有效值。
	DropInvalidMessageToCls *bool `json:"DropInvalidMessageToCls,omitempty" name:"DropInvalidMessageToCls"`

	// 投递cls的地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	DropClsRegion *string `json:"DropClsRegion,omitempty" name:"DropClsRegion"`

	// 投递cls的账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	DropClsOwneruin *string `json:"DropClsOwneruin,omitempty" name:"DropClsOwneruin"`

	// 投递cls的主题
	// 注意：此字段可能返回 null，表示取不到有效值。
	DropClsTopicId *string `json:"DropClsTopicId,omitempty" name:"DropClsTopicId"`

	// 投递cls的日志集id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DropClsLogSet *string `json:"DropClsLogSet,omitempty" name:"DropClsLogSet"`
}

type DtsConnectParam struct {
	// Dts的连接port
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Dts消费分组的Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Dts消费分组的账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Dts消费分组的密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitempty" name:"Password"`

	// Dts实例Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// Dts订阅的topic
	// 注意：此字段可能返回 null，表示取不到有效值。
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// 是否更新到关联的Datahub任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsUpdate *bool `json:"IsUpdate,omitempty" name:"IsUpdate"`
}

type DtsModifyConnectParam struct {
	// Dts实例Id【不支持修改】
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// Dts的连接port【不支持修改】
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Dts消费分组的Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Dts消费分组的账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Dts消费分组的密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 是否更新到关联的Datahub任务，默认为true
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsUpdate *bool `json:"IsUpdate,omitempty" name:"IsUpdate"`

	// Dts订阅的topic【不支持修改】
	// 注意：此字段可能返回 null，表示取不到有效值。
	Topic *string `json:"Topic,omitempty" name:"Topic"`
}

type DtsParam struct {
	// Dts实例Id
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// Dts的连接ip
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Dts的连接port
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Dts订阅的topic
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// Dts消费分组的Id
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Dts消费分组的账号
	GroupUser *string `json:"GroupUser,omitempty" name:"GroupUser"`

	// Dts消费分组的密码
	GroupPassword *string `json:"GroupPassword,omitempty" name:"GroupPassword"`

	// false同步原始数据，true同步解析后的json格式数据,默认true
	TranSql *bool `json:"TranSql,omitempty" name:"TranSql"`
}

type DynamicDiskConfig struct {
	// 动态硬盘扩容配置开关（0: 关闭，1: 开启）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// 每次磁盘动态扩容大小百分比
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepForwardPercentage *int64 `json:"StepForwardPercentage,omitempty" name:"StepForwardPercentage"`

	// 磁盘配额百分比触发条件，即消息达到此值触发硬盘自动扩容事件
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskQuotaPercentage *int64 `json:"DiskQuotaPercentage,omitempty" name:"DiskQuotaPercentage"`

	// 最大扩容硬盘大小，以 GB 为单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxDiskSpace *int64 `json:"MaxDiskSpace,omitempty" name:"MaxDiskSpace"`
}

type DynamicRetentionTime struct {
	// 动态消息保留时间配置开关（0: 关闭，1: 开启）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// 磁盘配额百分比触发条件，即消息达到此值触发消息保留时间变更事件
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskQuotaPercentage *int64 `json:"DiskQuotaPercentage,omitempty" name:"DiskQuotaPercentage"`

	// 每次向前调整消息保留时间百分比
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepForwardPercentage *int64 `json:"StepForwardPercentage,omitempty" name:"StepForwardPercentage"`

	// 保底时长，单位分钟
	// 注意：此字段可能返回 null，表示取不到有效值。
	BottomRetention *int64 `json:"BottomRetention,omitempty" name:"BottomRetention"`
}

type EsConnectParam struct {
	// Es的连接port
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Es连接源的用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Es连接源的密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitempty" name:"Password"`

	// Es连接源的实例资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// Es连接源是否为自建集群
	// 注意：此字段可能返回 null，表示取不到有效值。
	SelfBuilt *bool `json:"SelfBuilt,omitempty" name:"SelfBuilt"`

	// Es连接源的实例vip，当为腾讯云实例时，必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceVip *string `json:"ServiceVip,omitempty" name:"ServiceVip"`

	// Es连接源的vpcId，当为腾讯云实例时，必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// 是否更新到关联的Datahub任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsUpdate *bool `json:"IsUpdate,omitempty" name:"IsUpdate"`
}

type EsModifyConnectParam struct {
	// Es连接源的实例资源【不支持修改】
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// Es的连接port【不支持修改】
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Es连接源的实例vip【不支持修改】
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceVip *string `json:"ServiceVip,omitempty" name:"ServiceVip"`

	// Es连接源的vpcId【不支持修改】
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// Es连接源的用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Es连接源的密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitempty" name:"Password"`

	// Es连接源是否为自建集群【不支持修改】
	// 注意：此字段可能返回 null，表示取不到有效值。
	SelfBuilt *bool `json:"SelfBuilt,omitempty" name:"SelfBuilt"`

	// 是否更新到关联的Datahub任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsUpdate *bool `json:"IsUpdate,omitempty" name:"IsUpdate"`
}

type EsParam struct {
	// 实例资源
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// Es的连接port
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Es用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Es密码
	Password *string `json:"Password,omitempty" name:"Password"`

	// 是否为自建集群
	SelfBuilt *bool `json:"SelfBuilt,omitempty" name:"SelfBuilt"`

	// 实例vip
	ServiceVip *string `json:"ServiceVip,omitempty" name:"ServiceVip"`

	// 实例的vpcId
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// Es是否抛弃解析失败的消息
	DropInvalidMessage *bool `json:"DropInvalidMessage,omitempty" name:"DropInvalidMessage"`

	// Es自定义index名称
	Index *string `json:"Index,omitempty" name:"Index"`

	// Es自定义日期后缀
	DateFormat *string `json:"DateFormat,omitempty" name:"DateFormat"`

	// 非json格式数据的自定义key
	ContentKey *string `json:"ContentKey,omitempty" name:"ContentKey"`

	// Es是否抛弃非json格式的消息
	DropInvalidJsonMessage *bool `json:"DropInvalidJsonMessage,omitempty" name:"DropInvalidJsonMessage"`

	// 转储到Es中的文档ID取值字段名
	DocumentIdField *string `json:"DocumentIdField,omitempty" name:"DocumentIdField"`

	// Es自定义index名称的类型，STRING，JSONPATH，默认为STRING
	IndexType *string `json:"IndexType,omitempty" name:"IndexType"`

	// 当设置成员参数DropInvalidMessageToCls设置为true时,DropInvalidMessage参数失效
	DropCls *DropCls `json:"DropCls,omitempty" name:"DropCls"`

	// 转储到ES的消息为Database的binlog时，如果需要同步数据库操作，即增删改的操作到ES时填写数据库表主键
	DatabasePrimaryKey *string `json:"DatabasePrimaryKey,omitempty" name:"DatabasePrimaryKey"`

	// 死信队列
	DropDlq *FailureParam `json:"DropDlq,omitempty" name:"DropDlq"`
}

type EventBusParam struct {
	// 资源类型。EB_COS/EB_ES/EB_CLS
	Type *string `json:"Type,omitempty" name:"Type"`

	// 是否为自建集群
	SelfBuilt *bool `json:"SelfBuilt,omitempty" name:"SelfBuilt"`

	// 实例资源
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// SCF云函数命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// SCF云函数函数名
	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`

	// SCF云函数版本及别名
	Qualifier *string `json:"Qualifier,omitempty" name:"Qualifier"`
}

type FailureParam struct {
	// 类型，DLQ死信队列，IGNORE_ERROR保留，DROP废弃
	Type *string `json:"Type,omitempty" name:"Type"`

	// Ckafka类型死信队列
	KafkaParam *KafkaParam `json:"KafkaParam,omitempty" name:"KafkaParam"`

	// 重试间隔
	RetryInterval *uint64 `json:"RetryInterval,omitempty" name:"RetryInterval"`

	// 重试次数
	MaxRetryAttempts *uint64 `json:"MaxRetryAttempts,omitempty" name:"MaxRetryAttempts"`

	// DIP Topic类型死信队列
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicParam *TopicParam `json:"TopicParam,omitempty" name:"TopicParam"`

	// 死信队列类型，CKAFKA，TOPIC
	// 注意：此字段可能返回 null，表示取不到有效值。
	DlqType *string `json:"DlqType,omitempty" name:"DlqType"`
}

// Predefined struct for user
type FetchDatahubMessageByOffsetRequestParams struct {
	// 主题名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分区id
	Partition *int64 `json:"Partition,omitempty" name:"Partition"`

	// 位点信息，必填
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type FetchDatahubMessageByOffsetRequest struct {
	*tchttp.BaseRequest
	
	// 主题名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分区id
	Partition *int64 `json:"Partition,omitempty" name:"Partition"`

	// 位点信息，必填
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
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
	Result *ConsumerRecord `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 主题名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分区id
	Partition *int64 `json:"Partition,omitempty" name:"Partition"`

	// 位点信息
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 最大查询条数，最小1，最大100
	MessageCount *int64 `json:"MessageCount,omitempty" name:"MessageCount"`
}

type FetchLatestDatahubMessageListRequest struct {
	*tchttp.BaseRequest
	
	// 主题名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分区id
	Partition *int64 `json:"Partition,omitempty" name:"Partition"`

	// 位点信息
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 最大查询条数，最小1，最大100
	MessageCount *int64 `json:"MessageCount,omitempty" name:"MessageCount"`
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
	Result []*ConsumerRecord `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 主题名
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// 分区id
	Partition *int64 `json:"Partition,omitempty" name:"Partition"`

	// 位点信息，必填
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type FetchMessageByOffsetRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 主题名
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// 分区id
	Partition *int64 `json:"Partition,omitempty" name:"Partition"`

	// 位点信息，必填
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
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
	Result *ConsumerRecord `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 主题名
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// 分区id
	Partition *int64 `json:"Partition,omitempty" name:"Partition"`

	// 位点信息
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 最大查询条数，默认20，最大20
	SinglePartitionRecordNumber *int64 `json:"SinglePartitionRecordNumber,omitempty" name:"SinglePartitionRecordNumber"`
}

type FetchMessageListByOffsetRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 主题名
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// 分区id
	Partition *int64 `json:"Partition,omitempty" name:"Partition"`

	// 位点信息
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 最大查询条数，默认20，最大20
	SinglePartitionRecordNumber *int64 `json:"SinglePartitionRecordNumber,omitempty" name:"SinglePartitionRecordNumber"`
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
	Result []*ConsumerRecord `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type FieldParam struct {
	// 解析
	Analyse *AnalyseParam `json:"Analyse,omitempty" name:"Analyse"`

	// 二次解析
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecondaryAnalyse *SecondaryAnalyseParam `json:"SecondaryAnalyse,omitempty" name:"SecondaryAnalyse"`

	// 数据处理
	// 注意：此字段可能返回 null，表示取不到有效值。
	SMT []*SMTParam `json:"SMT,omitempty" name:"SMT"`

	// 测试结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 解析结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnalyseResult []*SMTParam `json:"AnalyseResult,omitempty" name:"AnalyseResult"`

	// 二次解析结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecondaryAnalyseResult []*SMTParam `json:"SecondaryAnalyseResult,omitempty" name:"SecondaryAnalyseResult"`

	// JSON格式解析结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnalyseJsonResult *string `json:"AnalyseJsonResult,omitempty" name:"AnalyseJsonResult"`

	// JSON格式二次解析结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecondaryAnalyseJsonResult *string `json:"SecondaryAnalyseJsonResult,omitempty" name:"SecondaryAnalyseJsonResult"`
}

type Filter struct {
	// 需要过滤的字段。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 字段的过滤值。
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type FilterMapParam struct {
	// Key值
	Key *string `json:"Key,omitempty" name:"Key"`

	// 匹配模式，前缀匹配PREFIX，后缀匹配SUFFIX，包含匹配CONTAINS，EXCEPT除外匹配，数值匹配NUMBER，IP匹配IP
	MatchMode *string `json:"MatchMode,omitempty" name:"MatchMode"`

	// Value值
	Value *string `json:"Value,omitempty" name:"Value"`

	// 固定REGULAR
	Type *string `json:"Type,omitempty" name:"Type"`
}

type Group struct {
	// 组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

type GroupInfoMember struct {
	// coordinator 为消费分组中的消费者生成的唯一 ID
	MemberId *string `json:"MemberId,omitempty" name:"MemberId"`

	// 客户消费者 SDK 自己设置的 client.id 信息
	ClientId *string `json:"ClientId,omitempty" name:"ClientId"`

	// 一般存储客户的 IP 地址
	ClientHost *string `json:"ClientHost,omitempty" name:"ClientHost"`

	// 存储着分配给该消费者的 partition 信息
	Assignment *Assignment `json:"Assignment,omitempty" name:"Assignment"`
}

type GroupInfoResponse struct {
	// 错误码，正常为0
	ErrorCode *string `json:"ErrorCode,omitempty" name:"ErrorCode"`

	// group 状态描述（常见的为 Empty、Stable、Dead 三种状态）：
	// Dead：消费分组不存在
	// Empty：消费分组，当前没有任何消费者订阅
	// PreparingRebalance：消费分组处于 rebalance 状态
	// CompletingRebalance：消费分组处于 rebalance 状态
	// Stable：消费分组中各个消费者已经加入，处于稳定状态
	State *string `json:"State,omitempty" name:"State"`

	// 消费分组选择的协议类型正常的消费者一般为 consumer 但有些系统采用了自己的协议如 kafka-connect 用的就是 connect。只有标准的 consumer 协议，本接口才知道具体的分配方式的格式，才能解析到具体的 partition 的分配情况
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// 消费者 partition 分配算法常见的有如下几种(Kafka 消费者 SDK 默认的选择项为 range)：range、 roundrobin、 sticky
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 仅当 state 为 Stable 且 protocol_type 为 consumer 时， 该数组才包含信息
	Members []*GroupInfoMember `json:"Members,omitempty" name:"Members"`

	// Kafka 消费分组
	Group *string `json:"Group,omitempty" name:"Group"`
}

type GroupInfoTopics struct {
	// 分配的 topic 名称
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// 分配的 partition 信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Partitions []*int64 `json:"Partitions,omitempty" name:"Partitions"`
}

type GroupOffsetPartition struct {
	// topic 的 partitionId
	Partition *int64 `json:"Partition,omitempty" name:"Partition"`

	// consumer 提交的 offset 位置
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 支持消费者提交消息时，传入 metadata 作为它用，当前一般为空字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	Metadata *string `json:"Metadata,omitempty" name:"Metadata"`

	// 错误码
	ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`

	// 当前 partition 最新的 offset
	LogEndOffset *int64 `json:"LogEndOffset,omitempty" name:"LogEndOffset"`

	// 未消费的消息个数
	Lag *int64 `json:"Lag,omitempty" name:"Lag"`
}

type GroupOffsetResponse struct {
	// 符合调节的总结果数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 该主题分区数组，其中每个元素为一个 json object
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicList []*GroupOffsetTopic `json:"TopicList,omitempty" name:"TopicList"`
}

type GroupOffsetTopic struct {
	// 主题名称
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// 该主题分区数组，其中每个元素为一个 json object
	// 注意：此字段可能返回 null，表示取不到有效值。
	Partitions []*GroupOffsetPartition `json:"Partitions,omitempty" name:"Partitions"`
}

type GroupResponse struct {
	// 计数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// GroupList
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupList []*DescribeGroup `json:"GroupList,omitempty" name:"GroupList"`

	// 消费分组配额
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupCountQuota *uint64 `json:"GroupCountQuota,omitempty" name:"GroupCountQuota"`
}

// Predefined struct for user
type InquireCkafkaPriceRequestParams struct {
	// 国内站标准版填写standards2, 专业版填写profession
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 购买/续费付费类型(购买时不填的话, 默认获取购买包年包月一个月的费用)
	InstanceChargeParam *InstanceChargeParam `json:"InstanceChargeParam,omitempty" name:"InstanceChargeParam"`

	// 购买/续费时购买的实例数量(不填时, 默认为1个)
	InstanceNum *int64 `json:"InstanceNum,omitempty" name:"InstanceNum"`

	// 实例内网带宽大小, 单位MB/s (购买时必填)
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 实例的硬盘购买类型以及大小 (购买时必填)
	InquiryDiskParam *InquiryDiskParam `json:"InquiryDiskParam,omitempty" name:"InquiryDiskParam"`

	// 实例消息保留时间大小, 单位小时 (购买时必填)
	MessageRetention *int64 `json:"MessageRetention,omitempty" name:"MessageRetention"`

	// 购买实例topic数, 单位个 (购买时必填)
	Topic *int64 `json:"Topic,omitempty" name:"Topic"`

	// 购买实例分区数, 单位个 (购买时必填)
	Partition *int64 `json:"Partition,omitempty" name:"Partition"`

	// 购买地域, 可通过查看DescribeCkafkaZone这个接口获取ZoneId
	ZoneIds []*int64 `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 标记操作, 新购填写purchase, 续费填写renew, (不填时, 默认为purchase)
	CategoryAction *string `json:"CategoryAction,omitempty" name:"CategoryAction"`

	// 国内站购买的版本, sv_ckafka_instance_s2_1(入门型), sv_ckafka_instance_s2_2(标准版), sv_ckafka_instance_s2_3(进阶型), 如果instanceType为standards2, 但该参数为空, 则默认值为sv_ckafka_instance_s2_1
	BillType *string `json:"BillType,omitempty" name:"BillType"`

	// 公网带宽计费模式, 目前只有专业版支持公网带宽 (购买公网带宽时必填)
	PublicNetworkParam *InquiryPublicNetworkParam `json:"PublicNetworkParam,omitempty" name:"PublicNetworkParam"`

	// 续费时的实例id, 续费时填写
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type InquireCkafkaPriceRequest struct {
	*tchttp.BaseRequest
	
	// 国内站标准版填写standards2, 专业版填写profession
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 购买/续费付费类型(购买时不填的话, 默认获取购买包年包月一个月的费用)
	InstanceChargeParam *InstanceChargeParam `json:"InstanceChargeParam,omitempty" name:"InstanceChargeParam"`

	// 购买/续费时购买的实例数量(不填时, 默认为1个)
	InstanceNum *int64 `json:"InstanceNum,omitempty" name:"InstanceNum"`

	// 实例内网带宽大小, 单位MB/s (购买时必填)
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 实例的硬盘购买类型以及大小 (购买时必填)
	InquiryDiskParam *InquiryDiskParam `json:"InquiryDiskParam,omitempty" name:"InquiryDiskParam"`

	// 实例消息保留时间大小, 单位小时 (购买时必填)
	MessageRetention *int64 `json:"MessageRetention,omitempty" name:"MessageRetention"`

	// 购买实例topic数, 单位个 (购买时必填)
	Topic *int64 `json:"Topic,omitempty" name:"Topic"`

	// 购买实例分区数, 单位个 (购买时必填)
	Partition *int64 `json:"Partition,omitempty" name:"Partition"`

	// 购买地域, 可通过查看DescribeCkafkaZone这个接口获取ZoneId
	ZoneIds []*int64 `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// 标记操作, 新购填写purchase, 续费填写renew, (不填时, 默认为purchase)
	CategoryAction *string `json:"CategoryAction,omitempty" name:"CategoryAction"`

	// 国内站购买的版本, sv_ckafka_instance_s2_1(入门型), sv_ckafka_instance_s2_2(标准版), sv_ckafka_instance_s2_3(进阶型), 如果instanceType为standards2, 但该参数为空, 则默认值为sv_ckafka_instance_s2_1
	BillType *string `json:"BillType,omitempty" name:"BillType"`

	// 公网带宽计费模式, 目前只有专业版支持公网带宽 (购买公网带宽时必填)
	PublicNetworkParam *InquiryPublicNetworkParam `json:"PublicNetworkParam,omitempty" name:"PublicNetworkParam"`

	// 续费时的实例id, 续费时填写
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstancePrice *InquiryPrice `json:"InstancePrice,omitempty" name:"InstancePrice"`

	// 公网带宽价格
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicNetworkBandwidthPrice *InquiryPrice `json:"PublicNetworkBandwidthPrice,omitempty" name:"PublicNetworkBandwidthPrice"`
}

// Predefined struct for user
type InquireCkafkaPriceResponseParams struct {
	// 出参
	Result *InquireCkafkaPriceResp `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnitPrice *float64 `json:"UnitPrice,omitempty" name:"UnitPrice"`

	// 折扣单位价格
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnitPriceDiscount *float64 `json:"UnitPriceDiscount,omitempty" name:"UnitPriceDiscount"`

	// 合计原价
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

	// 折扣合计价格
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`

	// 折扣(单位是%)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Discount *float64 `json:"Discount,omitempty" name:"Discount"`

	// 商品数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// 付费货币
	// 注意：此字段可能返回 null，表示取不到有效值。
	Currency *string `json:"Currency,omitempty" name:"Currency"`

	// 硬盘专用返回参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 购买时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 购买时长单位("m"按月, "h"按小时)
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 购买数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type InquiryDetailPrice struct {
	// 额外内网带宽价格
	// 注意：此字段可能返回 null，表示取不到有效值。
	BandwidthPrice *InquiryBasePrice `json:"BandwidthPrice,omitempty" name:"BandwidthPrice"`

	// 硬盘价格
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskPrice *InquiryBasePrice `json:"DiskPrice,omitempty" name:"DiskPrice"`

	// 额外分区价格
	// 注意：此字段可能返回 null，表示取不到有效值。
	PartitionPrice *InquiryBasePrice `json:"PartitionPrice,omitempty" name:"PartitionPrice"`

	// 额外Topic价格
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicPrice *InquiryBasePrice `json:"TopicPrice,omitempty" name:"TopicPrice"`

	// 实例套餐价格
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceTypePrice *InquiryBasePrice `json:"InstanceTypePrice,omitempty" name:"InstanceTypePrice"`
}

type InquiryDiskParam struct {
	// 购买硬盘类型: SSD(SSD), CLOUD_SSD(SSD云硬盘), CLOUD_PREMIUM(高性能云硬盘), CLOUD_BASIC(云盘)
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 购买硬盘大小: 单位GB
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
}

type InquiryPrice struct {
	// 单位原价
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnitPrice *float64 `json:"UnitPrice,omitempty" name:"UnitPrice"`

	// 折扣单位价格
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnitPriceDiscount *float64 `json:"UnitPriceDiscount,omitempty" name:"UnitPriceDiscount"`

	// 合计原价
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

	// 折扣合计价格
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`

	// 折扣(单位是%)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Discount *float64 `json:"Discount,omitempty" name:"Discount"`

	// 商品数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// 付费货币
	// 注意：此字段可能返回 null，表示取不到有效值。
	Currency *string `json:"Currency,omitempty" name:"Currency"`

	// 硬盘专用返回参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 购买时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 购买时长单位("m"按月, "h"按小时)
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 购买数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *int64 `json:"Value,omitempty" name:"Value"`

	// 详细类别的价格
	// 注意：此字段可能返回 null，表示取不到有效值。
	DetailPrices *InquiryDetailPrice `json:"DetailPrices,omitempty" name:"DetailPrices"`
}

type InquiryPublicNetworkParam struct {
	// 公网计费模式: BANDWIDTH_PREPAID(包年包月), BANDWIDTH_POSTPAID_BY_HOUR(带宽按小时计费)
	PublicNetworkChargeType *string `json:"PublicNetworkChargeType,omitempty" name:"PublicNetworkChargeType"`

	// 公网带宽, 单位MB
	PublicNetworkMonthly *int64 `json:"PublicNetworkMonthly,omitempty" name:"PublicNetworkMonthly"`
}

type Instance struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例的状态。0：创建中，1：运行中，2：删除中 ， 5 隔离中，-1 创建失败
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 是否开源实例。开源：true，不开源：false
	// 注意：此字段可能返回 null，表示取不到有效值。
	IfCommunity *bool `json:"IfCommunity,omitempty" name:"IfCommunity"`
}

type InstanceAttributesResponse struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 接入点 VIP 列表信息
	VipList []*VipEntity `json:"VipList,omitempty" name:"VipList"`

	// 虚拟IP
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 虚拟端口
	Vport *string `json:"Vport,omitempty" name:"Vport"`

	// 实例的状态。0：创建中，1：运行中，2：删除中
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 实例带宽，单位：Mbps
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 实例的存储大小，单位：GB
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 可用区
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// VPC 的 ID，为空表示是基础网络
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网 ID， 为空表示基础网络
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 实例健康状态， 1：健康，2：告警，3：异常
	Healthy *int64 `json:"Healthy,omitempty" name:"Healthy"`

	// 实例健康信息，当前会展示磁盘利用率，最大长度为256
	HealthyMessage *string `json:"HealthyMessage,omitempty" name:"HealthyMessage"`

	// 创建时间
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 消息保存时间,单位为分钟
	MsgRetentionTime *int64 `json:"MsgRetentionTime,omitempty" name:"MsgRetentionTime"`

	// 自动创建 Topic 配置， 若该字段为空，则表示未开启自动创建
	Config *InstanceConfigDO `json:"Config,omitempty" name:"Config"`

	// 剩余创建分区数
	RemainderPartitions *int64 `json:"RemainderPartitions,omitempty" name:"RemainderPartitions"`

	// 剩余创建主题数
	RemainderTopics *int64 `json:"RemainderTopics,omitempty" name:"RemainderTopics"`

	// 当前创建分区数
	CreatedPartitions *int64 `json:"CreatedPartitions,omitempty" name:"CreatedPartitions"`

	// 当前创建主题数
	CreatedTopics *int64 `json:"CreatedTopics,omitempty" name:"CreatedTopics"`

	// 标签数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *uint64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 跨可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneIds []*int64 `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// kafka版本信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitempty" name:"Version"`

	// 最大分组数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxGroupNum *int64 `json:"MaxGroupNum,omitempty" name:"MaxGroupNum"`

	// 售卖类型,0:标准版,1:专业版
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cvm *int64 `json:"Cvm,omitempty" name:"Cvm"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 表示该实例支持的特性。FEATURE_SUBNET_ACL:表示acl策略支持设置子网。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Features []*string `json:"Features,omitempty" name:"Features"`

	// 动态消息保留策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	RetentionTimeConfig *DynamicRetentionTime `json:"RetentionTimeConfig,omitempty" name:"RetentionTimeConfig"`

	// 最大连接数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxConnection *uint64 `json:"MaxConnection,omitempty" name:"MaxConnection"`

	// 公网带宽
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicNetwork *int64 `json:"PublicNetwork,omitempty" name:"PublicNetwork"`

	// 时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteRouteTimestamp *string `json:"DeleteRouteTimestamp,omitempty" name:"DeleteRouteTimestamp"`

	// 剩余创建分区数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RemainingPartitions *int64 `json:"RemainingPartitions,omitempty" name:"RemainingPartitions"`

	// 剩余创建主题数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RemainingTopics *int64 `json:"RemainingTopics,omitempty" name:"RemainingTopics"`

	// 动态硬盘扩容策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	DynamicDiskConfig *DynamicDiskConfig `json:"DynamicDiskConfig,omitempty" name:"DynamicDiskConfig"`
}

type InstanceChargeParam struct {
	// 实例付费类型: PREPAID(包年包月), POSTPAID_BY_HOUR(按量付费)
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 购买时长: 包年包月时需要填写, 按量计费无需填写
	InstanceChargePeriod *int64 `json:"InstanceChargePeriod,omitempty" name:"InstanceChargePeriod"`
}

type InstanceConfigDO struct {
	// 是否自动创建主题
	AutoCreateTopicsEnable *bool `json:"AutoCreateTopicsEnable,omitempty" name:"AutoCreateTopicsEnable"`

	// 分区数
	DefaultNumPartitions *int64 `json:"DefaultNumPartitions,omitempty" name:"DefaultNumPartitions"`

	// 默认的复制Factor
	DefaultReplicationFactor *int64 `json:"DefaultReplicationFactor,omitempty" name:"DefaultReplicationFactor"`
}

type InstanceDetail struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 访问实例的vip 信息
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 访问实例的端口信息
	Vport *string `json:"Vport,omitempty" name:"Vport"`

	// 虚拟IP列表
	VipList []*VipEntity `json:"VipList,omitempty" name:"VipList"`

	// 实例的状态。0：创建中，1：运行中，2：删除中：5隔离中， -1 创建失败
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 实例带宽，单位Mbps
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 实例的存储大小，单位GB
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 可用区域ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// vpcId，如果为空，说明是基础网络
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网id
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 实例是否续费，int  枚举值：1表示自动续费，2表示明确不自动续费
	RenewFlag *int64 `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// 实例状态 int：1表示健康，2表示告警，3 表示实例状态异常
	Healthy *int64 `json:"Healthy,omitempty" name:"Healthy"`

	// 实例状态信息
	HealthyMessage *string `json:"HealthyMessage,omitempty" name:"HealthyMessage"`

	// 实例创建时间时间
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 实例过期时间
	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 是否为内部客户。值为1 表示内部客户
	IsInternal *int64 `json:"IsInternal,omitempty" name:"IsInternal"`

	// Topic个数
	TopicNum *int64 `json:"TopicNum,omitempty" name:"TopicNum"`

	// 标识tag
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// kafka版本信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitempty" name:"Version"`

	// 跨可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneIds []*int64 `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// ckafka售卖类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cvm *int64 `json:"Cvm,omitempty" name:"Cvm"`

	// ckafka实例类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 磁盘类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 当前规格最大Topic数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxTopicNumber *int64 `json:"MaxTopicNumber,omitempty" name:"MaxTopicNumber"`

	// 当前规格最大Partition数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxPartitionNumber *int64 `json:"MaxPartitionNumber,omitempty" name:"MaxPartitionNumber"`

	// 计划升级配置时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	RebalanceTime *string `json:"RebalanceTime,omitempty" name:"RebalanceTime"`

	// 实例当前partition数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	PartitionNumber *uint64 `json:"PartitionNumber,omitempty" name:"PartitionNumber"`

	// 公网带宽类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicNetworkChargeType *string `json:"PublicNetworkChargeType,omitempty" name:"PublicNetworkChargeType"`

	// 公网带宽值
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicNetwork *int64 `json:"PublicNetwork,omitempty" name:"PublicNetwork"`

	// 实例类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 实例功能列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Features []*string `json:"Features,omitempty" name:"Features"`
}

type InstanceDetailResponse struct {
	// 符合条件的实例总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 符合条件的实例详情列表
	InstanceList []*InstanceDetail `json:"InstanceList,omitempty" name:"InstanceList"`
}

type InstanceQuotaConfigResp struct {
	// 生产限流大小，单位 MB/s
	// 注意：此字段可能返回 null，表示取不到有效值。
	QuotaProducerByteRate *int64 `json:"QuotaProducerByteRate,omitempty" name:"QuotaProducerByteRate"`

	// 消费限流大小，单位 MB/s
	// 注意：此字段可能返回 null，表示取不到有效值。
	QuotaConsumerByteRate *int64 `json:"QuotaConsumerByteRate,omitempty" name:"QuotaConsumerByteRate"`
}

type InstanceResponse struct {
	// 符合条件的实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceList []*Instance `json:"InstanceList,omitempty" name:"InstanceList"`

	// 符合条件的结果总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type JgwOperateResponse struct {
	// 返回的code，0为正常，非0为错误
	ReturnCode *string `json:"ReturnCode,omitempty" name:"ReturnCode"`

	// 成功消息
	ReturnMessage *string `json:"ReturnMessage,omitempty" name:"ReturnMessage"`

	// 操作型返回的Data数据,可能有flowId等
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *OperateResponseData `json:"Data,omitempty" name:"Data"`
}

type JsonPathReplaceParam struct {
	// 被替换值，Jsonpath表达式
	OldValue *string `json:"OldValue,omitempty" name:"OldValue"`

	// 替换值，Jsonpath表达式或字符串
	NewValue *string `json:"NewValue,omitempty" name:"NewValue"`
}

type KVParam struct {
	// 分隔符
	Delimiter *string `json:"Delimiter,omitempty" name:"Delimiter"`

	// key-value二次解析分隔符
	Regex *string `json:"Regex,omitempty" name:"Regex"`

	// 保留源Key，默认为false不保留
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeepOriginalKey *string `json:"KeepOriginalKey,omitempty" name:"KeepOriginalKey"`
}

type KafkaParam struct {
	// 是否为自建集群
	SelfBuilt *bool `json:"SelfBuilt,omitempty" name:"SelfBuilt"`

	// 实例资源
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// Topic名称，多个以“,”分隔
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// Offset类型，最开始位置earliest，最新位置latest，时间点位置timestamp
	// 注意：此字段可能返回 null，表示取不到有效值。
	OffsetType *string `json:"OffsetType,omitempty" name:"OffsetType"`

	// Offset类型为timestamp时必传，传时间戳，精确到秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 实例资源名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// Zone ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// Topic的Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Topic的分区数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PartitionNum *int64 `json:"PartitionNum,omitempty" name:"PartitionNum"`

	// 启用容错实例/开启死信队列
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableToleration *bool `json:"EnableToleration,omitempty" name:"EnableToleration"`

	// Qps 限制
	// 注意：此字段可能返回 null，表示取不到有效值。
	QpsLimit *uint64 `json:"QpsLimit,omitempty" name:"QpsLimit"`

	// Table到Topic的路由，「分发到多个topic」开关打开时必传
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableMappings []*TableMapping `json:"TableMappings,omitempty" name:"TableMappings"`

	// 「分发到多个topic」开关，默认为false
	// 注意：此字段可能返回 null，表示取不到有效值。
	UseTableMapping *bool `json:"UseTableMapping,omitempty" name:"UseTableMapping"`

	// 使用的Topic是否需要自动创建（目前只支持SOURCE流入任务，如果不使用分发到多个topic，需要在Topic字段填写需要自动创建的topic名）
	// 注意：此字段可能返回 null，表示取不到有效值。
	UseAutoCreateTopic *bool `json:"UseAutoCreateTopic,omitempty" name:"UseAutoCreateTopic"`

	// 写入Topic时是否进行压缩，不开启填"none"，开启的话，填写"open"。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompressionType *string `json:"CompressionType,omitempty" name:"CompressionType"`

	// 源topic消息1条扩增成msgMultiple条写入目标topic(该参数目前只有ckafka流入ckafka适用)
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgMultiple *int64 `json:"MsgMultiple,omitempty" name:"MsgMultiple"`
}

type LowercaseParam struct {

}

type MapParam struct {
	// key值
	Key *string `json:"Key,omitempty" name:"Key"`

	// 类型，DEFAULT默认，DATE系统预设-时间戳，CUSTOMIZE自定义，MAPPING映射
	Type *string `json:"Type,omitempty" name:"Type"`

	// 值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type MariaDBConnectParam struct {
	// MariaDB的连接port
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// MariaDB连接源的用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// MariaDB连接源的密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitempty" name:"Password"`

	// MariaDB连接源的实例资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// MariaDB连接源的实例vip，当为腾讯云实例时，必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceVip *string `json:"ServiceVip,omitempty" name:"ServiceVip"`

	// MariaDB连接源的vpcId，当为腾讯云实例时，必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// 是否更新到关联的Datahub任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsUpdate *bool `json:"IsUpdate,omitempty" name:"IsUpdate"`
}

type MariaDBModifyConnectParam struct {
	// MariaDB连接源的实例资源【不支持修改】
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// MariaDB的连接port【不支持修改】
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// MariaDB连接源的实例vip【不支持修改】
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceVip *string `json:"ServiceVip,omitempty" name:"ServiceVip"`

	// MariaDB连接源的vpcId【不支持修改】
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// MariaDB连接源的用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// MariaDB连接源的密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 是否更新到关联的Datahub任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsUpdate *bool `json:"IsUpdate,omitempty" name:"IsUpdate"`
}

type MariaDBParam struct {
	// MariaDB的数据库名称，"*"为全数据库
	Database *string `json:"Database,omitempty" name:"Database"`

	// MariaDB的数据表名称，"*"为所监听的所有数据库中的非系统表，可以","间隔，监听多个数据表，但数据表需要以"数据库名.数据表名"的格式进行填写
	Table *string `json:"Table,omitempty" name:"Table"`

	// 该MariaDB在连接管理内的Id
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// 复制存量信息(schema_only不复制, initial全量)，默认位initial
	SnapshotMode *string `json:"SnapshotMode,omitempty" name:"SnapshotMode"`

	// 格式：库1.表1:字段1,字段2;库2.表2:字段2，表之间;（分号）隔开，字段之间,（逗号）隔开。不指定的表默认取表的主键
	KeyColumns *string `json:"KeyColumns,omitempty" name:"KeyColumns"`

	// 当Table输入的是前缀时，该项值为true，否则为false
	IsTablePrefix *bool `json:"IsTablePrefix,omitempty" name:"IsTablePrefix"`

	// 输出格式，DEFAULT、CANAL_1、CANAL_2
	OutputFormat *string `json:"OutputFormat,omitempty" name:"OutputFormat"`

	// 如果该值为all，则DDL数据以及DML数据也会写入到选中的topic；若该值为dml，则只有DML数据写入到选中的topic
	IncludeContentChanges *string `json:"IncludeContentChanges,omitempty" name:"IncludeContentChanges"`

	// 如果该值为true，且MySQL中"binlog_rows_query_log_events"配置项的值为"ON"，则流入到topic的数据包含原SQL语句；若该值为false，流入到topic的数据不包含原SQL语句
	IncludeQuery *bool `json:"IncludeQuery,omitempty" name:"IncludeQuery"`

	// 如果该值为 true，则消息中会携带消息结构体对应的schema，如果该值为false则不会携带
	RecordWithSchema *bool `json:"RecordWithSchema,omitempty" name:"RecordWithSchema"`
}

// Predefined struct for user
type ModifyConnectResourceRequestParams struct {
	// 连接源的Id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 连接源名称，为空时不修改
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// 连接源描述，为空时不修改
	Description *string `json:"Description,omitempty" name:"Description"`

	// 连接源类型，修改数据源参数时，需要与原Type相同，否则编辑数据源无效
	Type *string `json:"Type,omitempty" name:"Type"`

	// Dts配置，Type为DTS时必填
	DtsConnectParam *DtsModifyConnectParam `json:"DtsConnectParam,omitempty" name:"DtsConnectParam"`

	// MongoDB配置，Type为MONGODB时必填
	MongoDBConnectParam *MongoDBModifyConnectParam `json:"MongoDBConnectParam,omitempty" name:"MongoDBConnectParam"`

	// Es配置，Type为ES时必填
	EsConnectParam *EsModifyConnectParam `json:"EsConnectParam,omitempty" name:"EsConnectParam"`

	// ClickHouse配置，Type为CLICKHOUSE时必填
	ClickHouseConnectParam *ClickHouseModifyConnectParam `json:"ClickHouseConnectParam,omitempty" name:"ClickHouseConnectParam"`

	// MySQL配置，Type为MYSQL或TDSQL_C_MYSQL时必填
	MySQLConnectParam *MySQLModifyConnectParam `json:"MySQLConnectParam,omitempty" name:"MySQLConnectParam"`

	// PostgreSQL配置，Type为POSTGRESQL或TDSQL_C_POSTGRESQL时必填
	PostgreSQLConnectParam *PostgreSQLModifyConnectParam `json:"PostgreSQLConnectParam,omitempty" name:"PostgreSQLConnectParam"`

	// MariaDB配置，Type为MARIADB时必填
	MariaDBConnectParam *MariaDBModifyConnectParam `json:"MariaDBConnectParam,omitempty" name:"MariaDBConnectParam"`

	// SQLServer配置，Type为SQLSERVER时必填
	SQLServerConnectParam *SQLServerModifyConnectParam `json:"SQLServerConnectParam,omitempty" name:"SQLServerConnectParam"`

	// Ctsdb配置，Type为CTSDB
	CtsdbConnectParam *CtsdbModifyConnectParam `json:"CtsdbConnectParam,omitempty" name:"CtsdbConnectParam"`

	// Doris配置，Type为DORIS
	DorisConnectParam *DorisModifyConnectParam `json:"DorisConnectParam,omitempty" name:"DorisConnectParam"`
}

type ModifyConnectResourceRequest struct {
	*tchttp.BaseRequest
	
	// 连接源的Id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 连接源名称，为空时不修改
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// 连接源描述，为空时不修改
	Description *string `json:"Description,omitempty" name:"Description"`

	// 连接源类型，修改数据源参数时，需要与原Type相同，否则编辑数据源无效
	Type *string `json:"Type,omitempty" name:"Type"`

	// Dts配置，Type为DTS时必填
	DtsConnectParam *DtsModifyConnectParam `json:"DtsConnectParam,omitempty" name:"DtsConnectParam"`

	// MongoDB配置，Type为MONGODB时必填
	MongoDBConnectParam *MongoDBModifyConnectParam `json:"MongoDBConnectParam,omitempty" name:"MongoDBConnectParam"`

	// Es配置，Type为ES时必填
	EsConnectParam *EsModifyConnectParam `json:"EsConnectParam,omitempty" name:"EsConnectParam"`

	// ClickHouse配置，Type为CLICKHOUSE时必填
	ClickHouseConnectParam *ClickHouseModifyConnectParam `json:"ClickHouseConnectParam,omitempty" name:"ClickHouseConnectParam"`

	// MySQL配置，Type为MYSQL或TDSQL_C_MYSQL时必填
	MySQLConnectParam *MySQLModifyConnectParam `json:"MySQLConnectParam,omitempty" name:"MySQLConnectParam"`

	// PostgreSQL配置，Type为POSTGRESQL或TDSQL_C_POSTGRESQL时必填
	PostgreSQLConnectParam *PostgreSQLModifyConnectParam `json:"PostgreSQLConnectParam,omitempty" name:"PostgreSQLConnectParam"`

	// MariaDB配置，Type为MARIADB时必填
	MariaDBConnectParam *MariaDBModifyConnectParam `json:"MariaDBConnectParam,omitempty" name:"MariaDBConnectParam"`

	// SQLServer配置，Type为SQLSERVER时必填
	SQLServerConnectParam *SQLServerModifyConnectParam `json:"SQLServerConnectParam,omitempty" name:"SQLServerConnectParam"`

	// Ctsdb配置，Type为CTSDB
	CtsdbConnectParam *CtsdbModifyConnectParam `json:"CtsdbConnectParam,omitempty" name:"CtsdbConnectParam"`

	// Doris配置，Type为DORIS
	DorisConnectParam *DorisModifyConnectParam `json:"DorisConnectParam,omitempty" name:"DorisConnectParam"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyConnectResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyConnectResourceResponseParams struct {
	// 连接源的Id
	Result *ConnectResourceResourceIdResp `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
}

type ModifyDatahubTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDatahubTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDatahubTaskResponseParams struct {
	// 任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *DatahubTaskIdRes `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ModifyGroupOffsetsRequestParams struct {
	// kafka实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// kafka 消费分组
	Group *string `json:"Group,omitempty" name:"Group"`

	// 重置offset的策略，入参含义 0. 对齐shift-by参数，代表把offset向前或向后移动shift条 1. 对齐参考(by-duration,to-datetime,to-earliest,to-latest),代表把offset移动到指定timestamp的位置 2. 对齐参考(to-offset)，代表把offset移动到指定的offset位置
	Strategy *int64 `json:"Strategy,omitempty" name:"Strategy"`

	// 表示需要重置的topics， 不填表示全部
	Topics []*string `json:"Topics,omitempty" name:"Topics"`

	// 当strategy为0时，必须包含该字段，可以大于零代表会把offset向后移动shift条，小于零则将offset向前回溯shift条数。正确重置后新的offset应该是(old_offset + shift)，需要注意的是如果新的offset小于partition的earliest则会设置为earliest，如果大于partition 的latest则会设置为latest
	Shift *int64 `json:"Shift,omitempty" name:"Shift"`

	// 单位ms。当strategy为1时，必须包含该字段，其中-2表示重置offset到最开始的位置，-1表示重置到最新的位置(相当于清空)，其它值则代表指定的时间，会获取topic中指定时间的offset然后进行重置，需要注意的时，如果指定的时间不存在消息，则获取最末尾的offset。
	ShiftTimestamp *int64 `json:"ShiftTimestamp,omitempty" name:"ShiftTimestamp"`

	// 需要重新设置的offset位置。当strategy为2，必须包含该字段。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 需要重新设置的partition的列表，如果没有指定Topics参数。则重置全部topics的对应的Partition列表里的partition。指定Topics时则重置指定的topic列表的对应的Partitions列表的partition。
	Partitions []*int64 `json:"Partitions,omitempty" name:"Partitions"`
}

type ModifyGroupOffsetsRequest struct {
	*tchttp.BaseRequest
	
	// kafka实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// kafka 消费分组
	Group *string `json:"Group,omitempty" name:"Group"`

	// 重置offset的策略，入参含义 0. 对齐shift-by参数，代表把offset向前或向后移动shift条 1. 对齐参考(by-duration,to-datetime,to-earliest,to-latest),代表把offset移动到指定timestamp的位置 2. 对齐参考(to-offset)，代表把offset移动到指定的offset位置
	Strategy *int64 `json:"Strategy,omitempty" name:"Strategy"`

	// 表示需要重置的topics， 不填表示全部
	Topics []*string `json:"Topics,omitempty" name:"Topics"`

	// 当strategy为0时，必须包含该字段，可以大于零代表会把offset向后移动shift条，小于零则将offset向前回溯shift条数。正确重置后新的offset应该是(old_offset + shift)，需要注意的是如果新的offset小于partition的earliest则会设置为earliest，如果大于partition 的latest则会设置为latest
	Shift *int64 `json:"Shift,omitempty" name:"Shift"`

	// 单位ms。当strategy为1时，必须包含该字段，其中-2表示重置offset到最开始的位置，-1表示重置到最新的位置(相当于清空)，其它值则代表指定的时间，会获取topic中指定时间的offset然后进行重置，需要注意的时，如果指定的时间不存在消息，则获取最末尾的offset。
	ShiftTimestamp *int64 `json:"ShiftTimestamp,omitempty" name:"ShiftTimestamp"`

	// 需要重新设置的offset位置。当strategy为2，必须包含该字段。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 需要重新设置的partition的列表，如果没有指定Topics参数。则重置全部topics的对应的Partition列表里的partition。指定Topics时则重置指定的topic列表的对应的Partitions列表的partition。
	Partitions []*int64 `json:"Partitions,omitempty" name:"Partitions"`
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
	Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	AutoCreateTopicEnable *bool `json:"AutoCreateTopicEnable,omitempty" name:"AutoCreateTopicEnable"`

	// 可选，如果auto.create.topic.enable设置为true没有设置该值时，默认设置为3
	DefaultNumPartitions *int64 `json:"DefaultNumPartitions,omitempty" name:"DefaultNumPartitions"`

	// 如歌auto.create.topic.enable设置为true没有指定该值时默认设置为2
	DefaultReplicationFactor *int64 `json:"DefaultReplicationFactor,omitempty" name:"DefaultReplicationFactor"`
}

// Predefined struct for user
type ModifyInstanceAttributesRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例日志的最长保留时间，单位分钟，最大30天，0代表不开启日志保留时间回收策略
	MsgRetentionTime *int64 `json:"MsgRetentionTime,omitempty" name:"MsgRetentionTime"`

	// 实例名称，是一个不超过 64 个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例配置
	Config *ModifyInstanceAttributesConfig `json:"Config,omitempty" name:"Config"`

	// 动态消息保留策略配置
	DynamicRetentionConfig *DynamicRetentionTime `json:"DynamicRetentionConfig,omitempty" name:"DynamicRetentionConfig"`

	// 修改升配置rebalance时间
	RebalanceTime *int64 `json:"RebalanceTime,omitempty" name:"RebalanceTime"`

	// 公网带宽
	PublicNetwork *int64 `json:"PublicNetwork,omitempty" name:"PublicNetwork"`

	// 动态硬盘扩容策略配置
	DynamicDiskConfig *DynamicDiskConfig `json:"DynamicDiskConfig,omitempty" name:"DynamicDiskConfig"`

	// 实例级别单条消息大小（单位byte)
	MaxMessageByte *uint64 `json:"MaxMessageByte,omitempty" name:"MaxMessageByte"`
}

type ModifyInstanceAttributesRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例日志的最长保留时间，单位分钟，最大30天，0代表不开启日志保留时间回收策略
	MsgRetentionTime *int64 `json:"MsgRetentionTime,omitempty" name:"MsgRetentionTime"`

	// 实例名称，是一个不超过 64 个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例配置
	Config *ModifyInstanceAttributesConfig `json:"Config,omitempty" name:"Config"`

	// 动态消息保留策略配置
	DynamicRetentionConfig *DynamicRetentionTime `json:"DynamicRetentionConfig,omitempty" name:"DynamicRetentionConfig"`

	// 修改升配置rebalance时间
	RebalanceTime *int64 `json:"RebalanceTime,omitempty" name:"RebalanceTime"`

	// 公网带宽
	PublicNetwork *int64 `json:"PublicNetwork,omitempty" name:"PublicNetwork"`

	// 动态硬盘扩容策略配置
	DynamicDiskConfig *DynamicDiskConfig `json:"DynamicDiskConfig,omitempty" name:"DynamicDiskConfig"`

	// 实例级别单条消息大小（单位byte)
	MaxMessageByte *uint64 `json:"MaxMessageByte,omitempty" name:"MaxMessageByte"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceAttributesResponseParams struct {
	// 返回结果
	Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例名称
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 预计磁盘，根据磁盘步长，规格向上调整。
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 预计带宽，根据带宽步长，规格向上调整。
	BandWidth *int64 `json:"BandWidth,omitempty" name:"BandWidth"`

	// 预计分区，根据带宽步长，规格向上调整。
	Partition *int64 `json:"Partition,omitempty" name:"Partition"`
}

type ModifyInstancePreRequest struct {
	*tchttp.BaseRequest
	
	// 实例名称
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 预计磁盘，根据磁盘步长，规格向上调整。
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 预计带宽，根据带宽步长，规格向上调整。
	BandWidth *int64 `json:"BandWidth,omitempty" name:"BandWidth"`

	// 预计分区，根据带宽步长，规格向上调整。
	Partition *int64 `json:"Partition,omitempty" name:"Partition"`
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
	Result *CreateInstancePreResp `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 用户名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 用户当前密码
	Password *string `json:"Password,omitempty" name:"Password"`

	// 用户新密码
	PasswordNew *string `json:"PasswordNew,omitempty" name:"PasswordNew"`
}

type ModifyPasswordRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 用户名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 用户当前密码
	Password *string `json:"Password,omitempty" name:"Password"`

	// 用户新密码
	PasswordNew *string `json:"PasswordNew,omitempty" name:"PasswordNew"`
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
	Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ModifyTopicAttributesRequestParams struct {
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 主题名称。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 主题备注，是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线-。
	Note *string `json:"Note,omitempty" name:"Note"`

	// IP 白名单开关，1：打开；0：关闭。
	EnableWhiteList *int64 `json:"EnableWhiteList,omitempty" name:"EnableWhiteList"`

	// 默认为1。
	MinInsyncReplicas *int64 `json:"MinInsyncReplicas,omitempty" name:"MinInsyncReplicas"`

	// 默认为 0，0：false；1：true。
	UncleanLeaderElectionEnable *int64 `json:"UncleanLeaderElectionEnable,omitempty" name:"UncleanLeaderElectionEnable"`

	// 消息保留时间，单位：ms，当前最小值为60000ms。
	RetentionMs *int64 `json:"RetentionMs,omitempty" name:"RetentionMs"`

	// Segment 分片滚动的时长，单位：ms，当前最小为86400000ms。
	SegmentMs *int64 `json:"SegmentMs,omitempty" name:"SegmentMs"`

	// 主题消息最大值，单位为 Byte，最大值为12582912Byte（即12MB）。
	MaxMessageBytes *int64 `json:"MaxMessageBytes,omitempty" name:"MaxMessageBytes"`

	// 消息删除策略，可以选择delete 或者compact
	CleanUpPolicy *string `json:"CleanUpPolicy,omitempty" name:"CleanUpPolicy"`

	// Ip白名单列表，配额限制，enableWhileList=1时必选
	IpWhiteList []*string `json:"IpWhiteList,omitempty" name:"IpWhiteList"`

	// 预设ACL规则, 1:打开  0:关闭，默认不打开
	EnableAclRule *int64 `json:"EnableAclRule,omitempty" name:"EnableAclRule"`

	// 预设ACL规则的名称
	AclRuleName *string `json:"AclRuleName,omitempty" name:"AclRuleName"`

	// 可选, 保留文件大小. 默认为-1,单位bytes, 当前最小值为1048576B
	RetentionBytes *int64 `json:"RetentionBytes,omitempty" name:"RetentionBytes"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 生产限流，单位 MB/s
	QuotaProducerByteRate *int64 `json:"QuotaProducerByteRate,omitempty" name:"QuotaProducerByteRate"`

	// 消费限流，单位 MB/s
	QuotaConsumerByteRate *int64 `json:"QuotaConsumerByteRate,omitempty" name:"QuotaConsumerByteRate"`

	// 调整topic副本数
	ReplicaNum *int64 `json:"ReplicaNum,omitempty" name:"ReplicaNum"`
}

type ModifyTopicAttributesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 主题名称。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 主题备注，是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线-。
	Note *string `json:"Note,omitempty" name:"Note"`

	// IP 白名单开关，1：打开；0：关闭。
	EnableWhiteList *int64 `json:"EnableWhiteList,omitempty" name:"EnableWhiteList"`

	// 默认为1。
	MinInsyncReplicas *int64 `json:"MinInsyncReplicas,omitempty" name:"MinInsyncReplicas"`

	// 默认为 0，0：false；1：true。
	UncleanLeaderElectionEnable *int64 `json:"UncleanLeaderElectionEnable,omitempty" name:"UncleanLeaderElectionEnable"`

	// 消息保留时间，单位：ms，当前最小值为60000ms。
	RetentionMs *int64 `json:"RetentionMs,omitempty" name:"RetentionMs"`

	// Segment 分片滚动的时长，单位：ms，当前最小为86400000ms。
	SegmentMs *int64 `json:"SegmentMs,omitempty" name:"SegmentMs"`

	// 主题消息最大值，单位为 Byte，最大值为12582912Byte（即12MB）。
	MaxMessageBytes *int64 `json:"MaxMessageBytes,omitempty" name:"MaxMessageBytes"`

	// 消息删除策略，可以选择delete 或者compact
	CleanUpPolicy *string `json:"CleanUpPolicy,omitempty" name:"CleanUpPolicy"`

	// Ip白名单列表，配额限制，enableWhileList=1时必选
	IpWhiteList []*string `json:"IpWhiteList,omitempty" name:"IpWhiteList"`

	// 预设ACL规则, 1:打开  0:关闭，默认不打开
	EnableAclRule *int64 `json:"EnableAclRule,omitempty" name:"EnableAclRule"`

	// 预设ACL规则的名称
	AclRuleName *string `json:"AclRuleName,omitempty" name:"AclRuleName"`

	// 可选, 保留文件大小. 默认为-1,单位bytes, 当前最小值为1048576B
	RetentionBytes *int64 `json:"RetentionBytes,omitempty" name:"RetentionBytes"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 生产限流，单位 MB/s
	QuotaProducerByteRate *int64 `json:"QuotaProducerByteRate,omitempty" name:"QuotaProducerByteRate"`

	// 消费限流，单位 MB/s
	QuotaConsumerByteRate *int64 `json:"QuotaConsumerByteRate,omitempty" name:"QuotaConsumerByteRate"`

	// 调整topic副本数
	ReplicaNum *int64 `json:"ReplicaNum,omitempty" name:"ReplicaNum"`
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
	delete(f, "SegmentMs")
	delete(f, "MaxMessageBytes")
	delete(f, "CleanUpPolicy")
	delete(f, "IpWhiteList")
	delete(f, "EnableAclRule")
	delete(f, "AclRuleName")
	delete(f, "RetentionBytes")
	delete(f, "Tags")
	delete(f, "QuotaProducerByteRate")
	delete(f, "QuotaConsumerByteRate")
	delete(f, "ReplicaNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTopicAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTopicAttributesResponseParams struct {
	// 返回结果集
	Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// MongoDB连接源的用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// MongoDB连接源的密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitempty" name:"Password"`

	// MongoDB连接源的实例资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// MongoDB连接源是否为自建集群
	// 注意：此字段可能返回 null，表示取不到有效值。
	SelfBuilt *bool `json:"SelfBuilt,omitempty" name:"SelfBuilt"`

	// MongoDB连接源的实例vip，当为腾讯云实例时，必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceVip *string `json:"ServiceVip,omitempty" name:"ServiceVip"`

	// MongoDB连接源的vpcId，当为腾讯云实例时，必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// 是否更新到关联的Datahub任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsUpdate *bool `json:"IsUpdate,omitempty" name:"IsUpdate"`
}

type MongoDBModifyConnectParam struct {
	// MongoDB连接源的实例资源【不支持修改】
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// MongoDB的连接port【不支持修改】
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// MongoDB连接源的实例vip【不支持修改】
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceVip *string `json:"ServiceVip,omitempty" name:"ServiceVip"`

	// MongoDB连接源的vpcId【不支持修改】
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// MongoDB连接源的用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// MongoDB连接源的密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitempty" name:"Password"`

	// MongoDB连接源是否为自建集群【不支持修改】
	// 注意：此字段可能返回 null，表示取不到有效值。
	SelfBuilt *bool `json:"SelfBuilt,omitempty" name:"SelfBuilt"`

	// 是否更新到关联的Datahub任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsUpdate *bool `json:"IsUpdate,omitempty" name:"IsUpdate"`
}

type MongoDBParam struct {
	// MongoDB的数据库名称
	Database *string `json:"Database,omitempty" name:"Database"`

	// MongoDB的集群
	Collection *string `json:"Collection,omitempty" name:"Collection"`

	// 是否复制存量数据，默认传参true
	CopyExisting *bool `json:"CopyExisting,omitempty" name:"CopyExisting"`

	// 实例资源
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// MongoDB的连接ip
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// MongoDB的连接port
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// MongoDB数据库用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// MongoDB数据库密码
	Password *string `json:"Password,omitempty" name:"Password"`

	// 监听事件类型，为空时表示全选。取值包括insert,update,replace,delete,invalidate,drop,dropdatabase,rename，多个类型间使用,逗号分隔
	ListeningEvent *string `json:"ListeningEvent,omitempty" name:"ListeningEvent"`

	// 主从优先级，默认主节点
	ReadPreference *string `json:"ReadPreference,omitempty" name:"ReadPreference"`

	// 聚合管道
	Pipeline *string `json:"Pipeline,omitempty" name:"Pipeline"`

	// 是否为自建集群
	SelfBuilt *bool `json:"SelfBuilt,omitempty" name:"SelfBuilt"`
}

type MySQLConnectParam struct {
	// MySQL的连接port
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// MySQL连接源的用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// MySQL连接源的密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitempty" name:"Password"`

	// MySQL连接源的实例资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// MySQL连接源的实例vip，当为腾讯云实例时，必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceVip *string `json:"ServiceVip,omitempty" name:"ServiceVip"`

	// MySQL连接源的vpcId，当为腾讯云实例时，必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// 是否更新到关联的Datahub任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsUpdate *bool `json:"IsUpdate,omitempty" name:"IsUpdate"`

	// 当type为TDSQL_C_MYSQL时，必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Mysql 连接源是否为自建集群
	// 注意：此字段可能返回 null，表示取不到有效值。
	SelfBuilt *bool `json:"SelfBuilt,omitempty" name:"SelfBuilt"`
}

type MySQLModifyConnectParam struct {
	// MySQL连接源的实例资源【不支持修改】
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// MySQL的连接port【不支持修改】
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// MySQL连接源的实例vip【不支持修改】
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceVip *string `json:"ServiceVip,omitempty" name:"ServiceVip"`

	// MySQL连接源的vpcId【不支持修改】
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// MySQL连接源的用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// MySQL连接源的密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 是否更新到关联的Datahub任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsUpdate *bool `json:"IsUpdate,omitempty" name:"IsUpdate"`

	// 当type为TDSQL_C_MYSQL时
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 是否是自建的集群
	// 注意：此字段可能返回 null，表示取不到有效值。
	SelfBuilt *bool `json:"SelfBuilt,omitempty" name:"SelfBuilt"`
}

type MySQLParam struct {
	// MySQL的数据库名称，"*"为全数据库
	Database *string `json:"Database,omitempty" name:"Database"`

	// MySQL的数据表名称，"*"为所监听的所有数据库中的非系统表，可以","间隔，监听多个数据表，但数据表需要以"数据库名.数据表名"的格式进行填写，需要填入正则表达式时，格式为"数据库名\\.数据表名"
	Table *string `json:"Table,omitempty" name:"Table"`

	// 该MySQL在连接管理内的Id
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// 复制存量信息(schema_only不复制, initial全量)，默认位initial
	SnapshotMode *string `json:"SnapshotMode,omitempty" name:"SnapshotMode"`

	// 存放MySQL的Ddl信息的Topic，为空则默认不存放
	DdlTopic *string `json:"DdlTopic,omitempty" name:"DdlTopic"`

	// "TABLE" 表示读取项为 table，"QUERY" 表示读取项为 query
	DataSourceMonitorMode *string `json:"DataSourceMonitorMode,omitempty" name:"DataSourceMonitorMode"`

	// 当 "DataMonitorMode"="TABLE" 时，传入需要读取的 Table；当 "DataMonitorMode"="QUERY" 时，传入需要读取的查询 sql 语句
	DataSourceMonitorResource *string `json:"DataSourceMonitorResource,omitempty" name:"DataSourceMonitorResource"`

	// "TIMESTAMP" 表示增量列为时间戳类型，"INCREMENT" 表示增量列为自增 id 类型
	DataSourceIncrementMode *string `json:"DataSourceIncrementMode,omitempty" name:"DataSourceIncrementMode"`

	// 传入需要监听的列名称
	DataSourceIncrementColumn *string `json:"DataSourceIncrementColumn,omitempty" name:"DataSourceIncrementColumn"`

	// "HEAD" 表示复制存量 + 增量数据，"TAIL" 表示只复制增量数据
	DataSourceStartFrom *string `json:"DataSourceStartFrom,omitempty" name:"DataSourceStartFrom"`

	// "INSERT" 表示使用 Insert 模式插入，"UPSERT" 表示使用 Upsert 模式插入
	DataTargetInsertMode *string `json:"DataTargetInsertMode,omitempty" name:"DataTargetInsertMode"`

	// 当 "DataInsertMode"="UPSERT" 时，传入当前 upsert 时依赖的主键
	DataTargetPrimaryKeyField *string `json:"DataTargetPrimaryKeyField,omitempty" name:"DataTargetPrimaryKeyField"`

	// 表与消息间的映射关系
	DataTargetRecordMapping []*RecordMapping `json:"DataTargetRecordMapping,omitempty" name:"DataTargetRecordMapping"`

	// 事件路由到特定主题的正则表达式，默认为(.*)
	TopicRegex *string `json:"TopicRegex,omitempty" name:"TopicRegex"`

	// TopicRegex的引用组，指定$1、$2等
	TopicReplacement *string `json:"TopicReplacement,omitempty" name:"TopicReplacement"`

	// 格式：库1.表1:字段1,字段2;库2.表2:字段2，表之间;（分号）隔开，字段之间,（逗号）隔开。不指定的表默认取表的主键
	KeyColumns *string `json:"KeyColumns,omitempty" name:"KeyColumns"`

	// Mysql 是否抛弃解析失败的消息，默认为true
	DropInvalidMessage *bool `json:"DropInvalidMessage,omitempty" name:"DropInvalidMessage"`

	// 当设置成员参数DropInvalidMessageToCls设置为true时,DropInvalidMessage参数失效
	DropCls *DropCls `json:"DropCls,omitempty" name:"DropCls"`

	// 输出格式，DEFAULT、CANAL_1、CANAL_2
	OutputFormat *string `json:"OutputFormat,omitempty" name:"OutputFormat"`

	// 当Table输入的是前缀时，该项值为true，否则为false
	IsTablePrefix *bool `json:"IsTablePrefix,omitempty" name:"IsTablePrefix"`

	// 如果该值为all，则DDL数据以及DML数据也会写入到选中的topic；若该值为dml，则只有DML数据写入到选中的topic
	IncludeContentChanges *string `json:"IncludeContentChanges,omitempty" name:"IncludeContentChanges"`

	// 如果该值为true，且MySQL中"binlog_rows_query_log_events"配置项的值为"ON"，则流入到topic的数据包含原SQL语句；若该值为false，流入到topic的数据不包含原SQL语句
	IncludeQuery *bool `json:"IncludeQuery,omitempty" name:"IncludeQuery"`

	// 如果该值为 true，则消息中会携带消息结构体对应的schema，如果该值为false则不会携带
	RecordWithSchema *bool `json:"RecordWithSchema,omitempty" name:"RecordWithSchema"`

	// 存放信令表的数据库名称
	SignalDatabase *string `json:"SignalDatabase,omitempty" name:"SignalDatabase"`

	// 输入的table是否为正则表达式，如果该选项以及IsTablePrefix同时为true，该选项的判断优先级高于IsTablePrefix
	IsTableRegular *bool `json:"IsTableRegular,omitempty" name:"IsTableRegular"`
}

type OperateResponseData struct {
	// FlowId11
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
}

type Partition struct {
	// 分区ID
	PartitionId *int64 `json:"PartitionId,omitempty" name:"PartitionId"`
}

type PartitionOffset struct {
	// Partition,例如"0"或"1"
	// 注意：此字段可能返回 null，表示取不到有效值。
	Partition *string `json:"Partition,omitempty" name:"Partition"`

	// Offset,例如100
	// 注意：此字段可能返回 null，表示取不到有效值。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type Partitions struct {
	// 分区
	Partition *int64 `json:"Partition,omitempty" name:"Partition"`

	// partition 消费位移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type PostgreSQLConnectParam struct {
	// PostgreSQL的连接port
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// PostgreSQL连接源的用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// PostgreSQL连接源的密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitempty" name:"Password"`

	// PostgreSQL连接源的实例资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// PostgreSQL连接源的实例vip，当为腾讯云实例时，必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceVip *string `json:"ServiceVip,omitempty" name:"ServiceVip"`

	// PostgreSQL连接源的vpcId，当为腾讯云实例时，必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// 当type为TDSQL_C_POSTGRESQL时，必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 是否更新到关联的Datahub任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsUpdate *bool `json:"IsUpdate,omitempty" name:"IsUpdate"`

	// PostgreSQL连接源是否为自建集群
	// 注意：此字段可能返回 null，表示取不到有效值。
	SelfBuilt *bool `json:"SelfBuilt,omitempty" name:"SelfBuilt"`
}

type PostgreSQLModifyConnectParam struct {
	// PostgreSQL连接源的实例资源【不支持修改】
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// PostgreSQL的连接port【不支持修改】
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// PostgreSQL连接源的实例vip【不支持修改】
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceVip *string `json:"ServiceVip,omitempty" name:"ServiceVip"`

	// PostgreSQL连接源的vpcId【不支持修改】
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// PostgreSQL连接源的用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// PostgreSQL连接源的密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 当type为TDSQL_C_POSTGRESQL时，该参数才有值【不支持修改】
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 是否更新到关联的Datahub任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsUpdate *bool `json:"IsUpdate,omitempty" name:"IsUpdate"`

	// 是否为自建集群
	// 注意：此字段可能返回 null，表示取不到有效值。
	SelfBuilt *bool `json:"SelfBuilt,omitempty" name:"SelfBuilt"`
}

type PostgreSQLParam struct {
	// PostgreSQL的数据库名称
	Database *string `json:"Database,omitempty" name:"Database"`

	// PostgreSQL的数据表名称，"*"为所监听的所有数据库中的非系统表，可以","间隔，监听多个数据表，但数据表需要以"Schema名.数据表名"的格式进行填写，需要填入正则表达式时，格式为"Schema名\\.数据表名"
	Table *string `json:"Table,omitempty" name:"Table"`

	// 该PostgreSQL在连接管理内的Id
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// 插件名(decoderbufs/pgoutput)，默认为decoderbufs
	PluginName *string `json:"PluginName,omitempty" name:"PluginName"`

	// 复制存量信息(never增量, initial全量)，默认为initial
	SnapshotMode *string `json:"SnapshotMode,omitempty" name:"SnapshotMode"`

	// 上游数据格式(JSON/Debezium), 当数据库同步模式为默认字段匹配时,必填
	DataFormat *string `json:"DataFormat,omitempty" name:"DataFormat"`

	// "INSERT" 表示使用 Insert 模式插入，"UPSERT" 表示使用 Upsert 模式插入
	DataTargetInsertMode *string `json:"DataTargetInsertMode,omitempty" name:"DataTargetInsertMode"`

	// 当 "DataInsertMode"="UPSERT" 时，传入当前 upsert 时依赖的主键
	DataTargetPrimaryKeyField *string `json:"DataTargetPrimaryKeyField,omitempty" name:"DataTargetPrimaryKeyField"`

	// 表与消息间的映射关系
	DataTargetRecordMapping []*RecordMapping `json:"DataTargetRecordMapping,omitempty" name:"DataTargetRecordMapping"`

	// 是否抛弃解析失败的消息，默认为true
	DropInvalidMessage *bool `json:"DropInvalidMessage,omitempty" name:"DropInvalidMessage"`

	// 输入的table是否为正则表达式
	IsTableRegular *bool `json:"IsTableRegular,omitempty" name:"IsTableRegular"`

	// 格式：库1.表1:字段1,字段2;库2.表2:字段2，表之间;（分号）隔开，字段之间,（逗号）隔开。不指定的表默认取表的主键
	KeyColumns *string `json:"KeyColumns,omitempty" name:"KeyColumns"`

	// 如果该值为 true，则消息中会携带消息结构体对应的schema，如果该值为false则不会携带
	RecordWithSchema *bool `json:"RecordWithSchema,omitempty" name:"RecordWithSchema"`
}

type Price struct {
	// 折扣价
	RealTotalCost *float64 `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// 原价
	TotalCost *float64 `json:"TotalCost,omitempty" name:"TotalCost"`
}

type PrivateLinkParam struct {
	// 客户实例的vip
	ServiceVip *string `json:"ServiceVip,omitempty" name:"ServiceVip"`

	// 客户实例的vpcId
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
}

type RecordMapping struct {
	// 消息的 key 名称
	JsonKey *string `json:"JsonKey,omitempty" name:"JsonKey"`

	// 消息类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 消息是否允许为空
	AllowNull *bool `json:"AllowNull,omitempty" name:"AllowNull"`

	// 对应映射列名称
	ColumnName *string `json:"ColumnName,omitempty" name:"ColumnName"`

	// 数据库表额外字段
	ExtraInfo *string `json:"ExtraInfo,omitempty" name:"ExtraInfo"`

	// 当前列大小
	ColumnSize *string `json:"ColumnSize,omitempty" name:"ColumnSize"`

	// 当前列精度
	DecimalDigits *string `json:"DecimalDigits,omitempty" name:"DecimalDigits"`

	// 是否为自增列
	AutoIncrement *bool `json:"AutoIncrement,omitempty" name:"AutoIncrement"`

	// 数据库表默认参数
	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`
}

type RegexReplaceParam struct {
	// 正则表达式
	Regex *string `json:"Regex,omitempty" name:"Regex"`

	// 替换新值
	NewValue *string `json:"NewValue,omitempty" name:"NewValue"`
}

type Region struct {
	// 地域ID
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// 地域名称
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 区域名称
	AreaName *string `json:"AreaName,omitempty" name:"AreaName"`

	// 地域代码
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionCode *string `json:"RegionCode,omitempty" name:"RegionCode"`

	// 地域代码（V3版本）
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionCodeV3 *string `json:"RegionCodeV3,omitempty" name:"RegionCodeV3"`

	// NONE:默认值不支持任何特殊机型\nCVM:支持CVM类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Support *string `json:"Support,omitempty" name:"Support"`

	// 是否支持ipv6, 0：表示不支持，1：表示支持
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ipv6 *int64 `json:"Ipv6,omitempty" name:"Ipv6"`

	// 是否支持跨可用区, 0：表示不支持，1：表示支持
	// 注意：此字段可能返回 null，表示取不到有效值。
	MultiZone *int64 `json:"MultiZone,omitempty" name:"MultiZone"`
}

type ReplaceParam struct {
	// 被替换值
	OldValue *string `json:"OldValue,omitempty" name:"OldValue"`

	// 替换值
	NewValue *string `json:"NewValue,omitempty" name:"NewValue"`
}

type Route struct {
	// 实例接入方式
	// 0：PLAINTEXT (明文方式，没有带用户信息老版本及社区版本都支持)
	// 1：SASL_PLAINTEXT（明文方式，不过在数据开始时，会通过SASL方式登录鉴权，仅社区版本支持）
	// 2：SSL（SSL加密通信，没有带用户信息，老版本及社区版本都支持）
	// 3：SASL_SSL（SSL加密通信，在数据开始时，会通过SASL方式登录鉴权，仅社区版本支持）
	AccessType *int64 `json:"AccessType,omitempty" name:"AccessType"`

	// 路由ID
	RouteId *int64 `json:"RouteId,omitempty" name:"RouteId"`

	// vip网络类型（1:外网TGW  2:基础网络 3:VPC网络 4:支撑网络(idc 环境) 5:SSL外网访问方式访问 6:黑石环境vpc 7:支撑网络(cvm 环境）
	VipType *int64 `json:"VipType,omitempty" name:"VipType"`

	// 虚拟IP列表
	VipList []*VipEntity `json:"VipList,omitempty" name:"VipList"`

	// 域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名port
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainPort *int64 `json:"DomainPort,omitempty" name:"DomainPort"`

	// 时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteTimestamp *string `json:"DeleteTimestamp,omitempty" name:"DeleteTimestamp"`
}

type RouteResponse struct {
	// 路由信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Routers []*Route `json:"Routers,omitempty" name:"Routers"`
}

type RowParam struct {
	// 行内容，KEY_VALUE，VALUE
	RowContent *string `json:"RowContent,omitempty" name:"RowContent"`

	// key和value间的分隔符
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyValueDelimiter *string `json:"KeyValueDelimiter,omitempty" name:"KeyValueDelimiter"`

	// 元素建的分隔符
	// 注意：此字段可能返回 null，表示取不到有效值。
	EntryDelimiter *string `json:"EntryDelimiter,omitempty" name:"EntryDelimiter"`
}

type SMTParam struct {
	// 数据处理KEY
	Key *string `json:"Key,omitempty" name:"Key"`

	// 操作，DATE系统预设-时间戳，CUSTOMIZE自定义，MAPPING映射，JSONPATH
	Operate *string `json:"Operate,omitempty" name:"Operate"`

	// 数据类型，ORIGINAL原始，STRING，INT64，FLOAT64，BOOLEAN，MAP，ARRAY
	SchemeType *string `json:"SchemeType,omitempty" name:"SchemeType"`

	// 数据处理VALUE
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// VALUE处理
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValueOperate *ValueParam `json:"ValueOperate,omitempty" name:"ValueOperate"`

	// 原始VALUE
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginalValue *string `json:"OriginalValue,omitempty" name:"OriginalValue"`

	// VALUE处理链
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValueOperates []*ValueParam `json:"ValueOperates,omitempty" name:"ValueOperates"`
}

type SQLServerConnectParam struct {
	// SQLServer的连接port
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// SQLServer连接源的用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// SQLServer连接源的密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitempty" name:"Password"`

	// SQLServer连接源的实例资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// SQLServer连接源的实例vip，当为腾讯云实例时，必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceVip *string `json:"ServiceVip,omitempty" name:"ServiceVip"`

	// SQLServer连接源的vpcId，当为腾讯云实例时，必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// 是否更新到关联的Dip任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsUpdate *bool `json:"IsUpdate,omitempty" name:"IsUpdate"`
}

type SQLServerModifyConnectParam struct {
	// SQLServer连接源的实例资源【不支持修改】
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// SQLServer的连接port【不支持修改】
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// SQLServer连接源的实例vip【不支持修改】
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceVip *string `json:"ServiceVip,omitempty" name:"ServiceVip"`

	// SQLServer连接源的vpcId【不支持修改】
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// SQLServer连接源的用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// SQLServer连接源的密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 是否更新到关联的Dip任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsUpdate *bool `json:"IsUpdate,omitempty" name:"IsUpdate"`
}

type SQLServerParam struct {
	// SQLServer的数据库名称
	Database *string `json:"Database,omitempty" name:"Database"`

	// SQLServer的数据表名称，"*"为所监听的所有数据库中的非系统表，可以","间隔，监听多个数据表，但数据表需要以"数据库名.数据表名"的格式进行填写
	Table *string `json:"Table,omitempty" name:"Table"`

	// 该SQLServer在连接管理内的Id
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// 复制存量信息(schema_only增量, initial全量)，默认为initial
	SnapshotMode *string `json:"SnapshotMode,omitempty" name:"SnapshotMode"`
}

type SaleInfo struct {
	// 手动设置的flag标志
	// 注意：此字段可能返回 null，表示取不到有效值。
	Flag *bool `json:"Flag,omitempty" name:"Flag"`

	// ckakfa版本号(1.1.1/2.4.2/0.10.2)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitempty" name:"Version"`

	// 专业版、标准版标志
	// 注意：此字段可能返回 null，表示取不到有效值。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 售罄标志：true售罄
	// 注意：此字段可能返回 null，表示取不到有效值。
	SoldOut *bool `json:"SoldOut,omitempty" name:"SoldOut"`
}

type ScfParam struct {
	// SCF云函数函数名
	// 注意：此字段可能返回 null，表示取不到有效值。
	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`

	// SCF云函数命名空间, 默认为default
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// SCF云函数版本及别名, 默认为$DEFAULT
	// 注意：此字段可能返回 null，表示取不到有效值。
	Qualifier *string `json:"Qualifier,omitempty" name:"Qualifier"`

	// 每批最大发送消息数, 默认为1000
	BatchSize *int64 `json:"BatchSize,omitempty" name:"BatchSize"`

	// SCF调用失败后重试次数, 默认为5
	MaxRetries *int64 `json:"MaxRetries,omitempty" name:"MaxRetries"`
}

type SecondaryAnalyseParam struct {
	// 分隔符
	Regex *string `json:"Regex,omitempty" name:"Regex"`
}

// Predefined struct for user
type SendMessageRequestParams struct {
	// DataHub接入ID
	DataHubId *string `json:"DataHubId,omitempty" name:"DataHubId"`

	// 发送消息内容(单次请求最多500条)
	Message []*BatchContent `json:"Message,omitempty" name:"Message"`
}

type SendMessageRequest struct {
	*tchttp.BaseRequest
	
	// DataHub接入ID
	DataHubId *string `json:"DataHubId,omitempty" name:"DataHubId"`

	// 发送消息内容(单次请求最多500条)
	Message []*BatchContent `json:"Message,omitempty" name:"Message"`
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
	MessageId []*string `json:"MessageId,omitempty" name:"MessageId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Regex *string `json:"Regex,omitempty" name:"Regex"`
}

type SubscribedInfo struct {
	// 订阅的主题名
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 订阅的分区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Partition []*int64 `json:"Partition,omitempty" name:"Partition"`

	// 分区offset信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PartitionOffset []*PartitionOffset `json:"PartitionOffset,omitempty" name:"PartitionOffset"`

	// 订阅的主题ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

type SubstrParam struct {
	// 截取起始位置
	Start *int64 `json:"Start,omitempty" name:"Start"`

	// 截取截止位置
	End *int64 `json:"End,omitempty" name:"End"`
}

type TableMapping struct {
	// 库名
	Database *string `json:"Database,omitempty" name:"Database"`

	// 表名，多个表,（逗号）隔开
	Table *string `json:"Table,omitempty" name:"Table"`

	// Topic名称
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// Topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

type Tag struct {
	// 标签的key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签的值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type TdwParam struct {
	// Tdw的bid
	Bid *string `json:"Bid,omitempty" name:"Bid"`

	// Tdw的tid
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 默认true
	IsDomestic *bool `json:"IsDomestic,omitempty" name:"IsDomestic"`

	// TDW地址，默认tl-tdbank-tdmanager.tencent-distribute.com
	TdwHost *string `json:"TdwHost,omitempty" name:"TdwHost"`

	// TDW端口，默认8099
	TdwPort *int64 `json:"TdwPort,omitempty" name:"TdwPort"`
}

type Topic struct {
	// 主题的ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 主题的名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Note *string `json:"Note,omitempty" name:"Note"`
}

type TopicAttributesResponse struct {
	// 主题 ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 创建时间
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 主题备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Note *string `json:"Note,omitempty" name:"Note"`

	// 分区个数
	PartitionNum *int64 `json:"PartitionNum,omitempty" name:"PartitionNum"`

	// IP 白名单开关，1：打开； 0：关闭
	EnableWhiteList *int64 `json:"EnableWhiteList,omitempty" name:"EnableWhiteList"`

	// IP 白名单列表
	IpWhiteList []*string `json:"IpWhiteList,omitempty" name:"IpWhiteList"`

	// topic 配置数组
	Config *Config `json:"Config,omitempty" name:"Config"`

	// 分区详情
	Partitions []*TopicPartitionDO `json:"Partitions,omitempty" name:"Partitions"`

	// ACL预设策略开关，1：打开； 0：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableAclRule *int64 `json:"EnableAclRule,omitempty" name:"EnableAclRule"`

	// 预设策略列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AclRuleList []*AclRule `json:"AclRuleList,omitempty" name:"AclRuleList"`

	// topic 限流策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	QuotaConfig *InstanceQuotaConfigResp `json:"QuotaConfig,omitempty" name:"QuotaConfig"`

	// 副本数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReplicaNum *int64 `json:"ReplicaNum,omitempty" name:"ReplicaNum"`
}

type TopicDetail struct {
	// 主题名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 分区数
	PartitionNum *int64 `json:"PartitionNum,omitempty" name:"PartitionNum"`

	// 副本数
	ReplicaNum *int64 `json:"ReplicaNum,omitempty" name:"ReplicaNum"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Note *string `json:"Note,omitempty" name:"Note"`

	// 创建时间
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 是否开启ip鉴权白名单，true表示开启，false表示不开启
	EnableWhiteList *bool `json:"EnableWhiteList,omitempty" name:"EnableWhiteList"`

	// ip白名单中ip个数
	IpWhiteListCount *int64 `json:"IpWhiteListCount,omitempty" name:"IpWhiteListCount"`

	// 数据备份cos bucket: 转存到cos 的bucket地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	ForwardCosBucket *string `json:"ForwardCosBucket,omitempty" name:"ForwardCosBucket"`

	// 数据备份cos 状态： 1 不开启数据备份，0 开启数据备份
	ForwardStatus *int64 `json:"ForwardStatus,omitempty" name:"ForwardStatus"`

	// 数据备份到cos的周期频率
	ForwardInterval *int64 `json:"ForwardInterval,omitempty" name:"ForwardInterval"`

	// 高级配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Config *Config `json:"Config,omitempty" name:"Config"`

	// 消息保留时间配置(用于动态配置变更记录)
	// 注意：此字段可能返回 null，表示取不到有效值。
	RetentionTimeConfig *TopicRetentionTimeConfigRsp `json:"RetentionTimeConfig,omitempty" name:"RetentionTimeConfig"`

	// 0:正常，1：已删除，2：删除中
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type TopicDetailResponse struct {
	// 返回的主题详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicList []*TopicDetail `json:"TopicList,omitempty" name:"TopicList"`

	// 符合条件的所有主题详情数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type TopicInSyncReplicaInfo struct {
	// 分区名称
	Partition *string `json:"Partition,omitempty" name:"Partition"`

	// Leader Id
	Leader *uint64 `json:"Leader,omitempty" name:"Leader"`

	// 副本集
	Replica *string `json:"Replica,omitempty" name:"Replica"`

	// ISR
	InSyncReplica *string `json:"InSyncReplica,omitempty" name:"InSyncReplica"`

	// 起始Offset
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeginOffset *uint64 `json:"BeginOffset,omitempty" name:"BeginOffset"`

	// 末端Offset
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndOffset *uint64 `json:"EndOffset,omitempty" name:"EndOffset"`

	// 消息数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MessageCount *uint64 `json:"MessageCount,omitempty" name:"MessageCount"`

	// 未同步副本集
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutOfSyncReplica *string `json:"OutOfSyncReplica,omitempty" name:"OutOfSyncReplica"`
}

type TopicInSyncReplicaResult struct {
	// Topic详情及副本合集
	TopicInSyncReplicaList []*TopicInSyncReplicaInfo `json:"TopicInSyncReplicaList,omitempty" name:"TopicInSyncReplicaList"`

	// 总计个数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type TopicParam struct {
	// 单独售卖Topic的Topic名称
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// Offset类型，最开始位置earliest，最新位置latest，时间点位置timestamp
	// 注意：此字段可能返回 null，表示取不到有效值。
	OffsetType *string `json:"OffsetType,omitempty" name:"OffsetType"`

	// Offset类型为timestamp时必传，传时间戳，精确到秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// Topic的TopicId【出参】
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 写入Topic时是否进行压缩，不开启填"none"，开启的话，可选择"gzip", "snappy", "lz4"中的一个进行填写。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompressionType *string `json:"CompressionType,omitempty" name:"CompressionType"`

	// 使用的Topic是否需要自动创建（目前只支持SOURCE流入任务）
	// 注意：此字段可能返回 null，表示取不到有效值。
	UseAutoCreateTopic *bool `json:"UseAutoCreateTopic,omitempty" name:"UseAutoCreateTopic"`

	// 源topic消息1条扩增成msgMultiple条写入目标topic(该参数目前只有ckafka流入ckafka适用)
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgMultiple *int64 `json:"MsgMultiple,omitempty" name:"MsgMultiple"`
}

type TopicPartitionDO struct {
	// Partition ID
	Partition *int64 `json:"Partition,omitempty" name:"Partition"`

	// Leader 运行状态
	LeaderStatus *int64 `json:"LeaderStatus,omitempty" name:"LeaderStatus"`

	// ISR 个数
	IsrNum *int64 `json:"IsrNum,omitempty" name:"IsrNum"`

	// 副本个数
	ReplicaNum *int64 `json:"ReplicaNum,omitempty" name:"ReplicaNum"`
}

type TopicResult struct {
	// 返回的主题信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicList []*Topic `json:"TopicList,omitempty" name:"TopicList"`

	// 符合条件的 topic 数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type TopicRetentionTimeConfigRsp struct {
	// 期望值，即用户配置的Topic消息保留时间(单位分钟)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Expect *int64 `json:"Expect,omitempty" name:"Expect"`

	// 当前值，即当前生效值(可能存在动态调整，单位分钟)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Current *int64 `json:"Current,omitempty" name:"Current"`

	// 最近变更时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModTimeStamp *int64 `json:"ModTimeStamp,omitempty" name:"ModTimeStamp"`
}

type TopicSubscribeGroup struct {
	// 总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 消费分组状态数量信息
	StatusCountInfo *string `json:"StatusCountInfo,omitempty" name:"StatusCountInfo"`

	// 消费分组信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupsInfo []*GroupInfoResponse `json:"GroupsInfo,omitempty" name:"GroupsInfo"`

	// 此次请求是否异步的状态。实例里分组较少的会直接返回结果,Status为1。当分组较多时,会异步更新缓存，Status为0时不会返回分组信息，直至Status为1更新完毕返回结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type TransformParam struct {
	// 解析格式，JSON，DELIMITER分隔符，REGULAR正则提取
	AnalysisFormat *string `json:"AnalysisFormat,omitempty" name:"AnalysisFormat"`

	// 输出格式
	OutputFormat *string `json:"OutputFormat,omitempty" name:"OutputFormat"`

	// 是否保留解析失败数据
	FailureParam *FailureParam `json:"FailureParam,omitempty" name:"FailureParam"`

	// 原始数据
	Content *string `json:"Content,omitempty" name:"Content"`

	// 数据来源，TOPIC从源topic拉取，CUSTOMIZE自定义
	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`

	// 分隔符、正则表达式
	Regex *string `json:"Regex,omitempty" name:"Regex"`

	// Map
	MapParam []*MapParam `json:"MapParam,omitempty" name:"MapParam"`

	// 过滤器
	FilterParam []*FilterMapParam `json:"FilterParam,omitempty" name:"FilterParam"`

	// 测试结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 解析结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnalyseResult []*MapParam `json:"AnalyseResult,omitempty" name:"AnalyseResult"`

	// 底层引擎是否使用eb
	// 注意：此字段可能返回 null，表示取不到有效值。
	UseEventBus *bool `json:"UseEventBus,omitempty" name:"UseEventBus"`
}

type TransformsParam struct {
	// 原始数据
	Content *string `json:"Content,omitempty" name:"Content"`

	// 处理链
	FieldChain []*FieldParam `json:"FieldChain,omitempty" name:"FieldChain"`

	// 过滤器
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilterParam []*FilterMapParam `json:"FilterParam,omitempty" name:"FilterParam"`

	// 失败处理
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailureParam *FailureParam `json:"FailureParam,omitempty" name:"FailureParam"`

	// 测试结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 数据来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`

	// 输出格式，JSON，ROW，默认为JSON
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputFormat *string `json:"OutputFormat,omitempty" name:"OutputFormat"`

	// 输出格式为ROW必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	RowParam *RowParam `json:"RowParam,omitempty" name:"RowParam"`

	// 是否保留数据源Topic元数据信息（源Topic、Partition、Offset），默认为false
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeepMetadata *bool `json:"KeepMetadata,omitempty" name:"KeepMetadata"`

	// 数组解析
	// 注意：此字段可能返回 null，表示取不到有效值。
	BatchAnalyse *BatchAnalyseParam `json:"BatchAnalyse,omitempty" name:"BatchAnalyse"`
}

type UrlDecodeParam struct {
	// 编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	CharsetName *string `json:"CharsetName,omitempty" name:"CharsetName"`
}

type User struct {
	// 用户id
	UserId *int64 `json:"UserId,omitempty" name:"UserId"`

	// 用户名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 最后更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type UserResponse struct {
	// 符合条件的用户列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Users []*User `json:"Users,omitempty" name:"Users"`

	// 符合条件的总用户数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type ValueParam struct {
	// 处理模式，REPLACE替换，SUBSTR截取，DATE日期转换，TRIM去除前后空格，REGEX_REPLACE正则替换，URL_DECODE，LOWERCASE转换为小写
	Type *string `json:"Type,omitempty" name:"Type"`

	// 替换，TYPE=REPLACE时必传
	// 注意：此字段可能返回 null，表示取不到有效值。
	Replace *ReplaceParam `json:"Replace,omitempty" name:"Replace"`

	// 截取，TYPE=SUBSTR时必传
	// 注意：此字段可能返回 null，表示取不到有效值。
	Substr *SubstrParam `json:"Substr,omitempty" name:"Substr"`

	// 时间转换，TYPE=DATE时必传
	// 注意：此字段可能返回 null，表示取不到有效值。
	Date *DateParam `json:"Date,omitempty" name:"Date"`

	// 正则替换，TYPE=REGEX_REPLACE时必传
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegexReplace *RegexReplaceParam `json:"RegexReplace,omitempty" name:"RegexReplace"`

	// 值支持一拆多，TYPE=SPLIT时必传
	// 注意：此字段可能返回 null，表示取不到有效值。
	Split *SplitParam `json:"Split,omitempty" name:"Split"`

	// key-value二次解析，TYPE=KV时必传
	// 注意：此字段可能返回 null，表示取不到有效值。
	KV *KVParam `json:"KV,omitempty" name:"KV"`

	// 处理结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *string `json:"Result,omitempty" name:"Result"`

	// JsonPath替换，TYPE=JSON_PATH_REPLACE时必传
	// 注意：此字段可能返回 null，表示取不到有效值。
	JsonPathReplace *JsonPathReplaceParam `json:"JsonPathReplace,omitempty" name:"JsonPathReplace"`

	// Url解析
	// 注意：此字段可能返回 null，表示取不到有效值。
	UrlDecode *UrlDecodeParam `json:"UrlDecode,omitempty" name:"UrlDecode"`

	// 小写字符解析
	// 注意：此字段可能返回 null，表示取不到有效值。
	Lowercase *LowercaseParam `json:"Lowercase,omitempty" name:"Lowercase"`
}

type VipEntity struct {
	// 虚拟IP
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 虚拟端口
	Vport *string `json:"Vport,omitempty" name:"Vport"`
}

type ZoneInfo struct {
	// zone的id
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 是否内部APP
	IsInternalApp *int64 `json:"IsInternalApp,omitempty" name:"IsInternalApp"`

	// app id
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 标识
	Flag *bool `json:"Flag,omitempty" name:"Flag"`

	// zone名称
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// zone状态
	ZoneStatus *int64 `json:"ZoneStatus,omitempty" name:"ZoneStatus"`

	// 额外标识
	Exflag *string `json:"Exflag,omitempty" name:"Exflag"`

	// json对象，key为机型，value true为售罄，false为未售罄
	SoldOut *string `json:"SoldOut,omitempty" name:"SoldOut"`

	// 标准版售罄信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SalesInfo []*SaleInfo `json:"SalesInfo,omitempty" name:"SalesInfo"`
}

type ZoneResponse struct {
	// zone列表
	ZoneList []*ZoneInfo `json:"ZoneList,omitempty" name:"ZoneList"`

	// 最大购买实例个数
	MaxBuyInstanceNum *int64 `json:"MaxBuyInstanceNum,omitempty" name:"MaxBuyInstanceNum"`

	// 最大购买带宽 单位Mb/s
	MaxBandwidth *int64 `json:"MaxBandwidth,omitempty" name:"MaxBandwidth"`

	// 后付费单位价格
	UnitPrice *Price `json:"UnitPrice,omitempty" name:"UnitPrice"`

	// 后付费消息单价
	MessagePrice *Price `json:"MessagePrice,omitempty" name:"MessagePrice"`

	// 用户独占集群信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterInfo []*ClusterInfo `json:"ClusterInfo,omitempty" name:"ClusterInfo"`

	// 购买标准版配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Standard *string `json:"Standard,omitempty" name:"Standard"`

	// 购买标准版S2配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	StandardS2 *string `json:"StandardS2,omitempty" name:"StandardS2"`

	// 购买专业版配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Profession *string `json:"Profession,omitempty" name:"Profession"`

	// 购买物理独占版配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Physical *string `json:"Physical,omitempty" name:"Physical"`

	// 公网带宽
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicNetwork *string `json:"PublicNetwork,omitempty" name:"PublicNetwork"`

	// 公网带宽配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicNetworkLimit *string `json:"PublicNetworkLimit,omitempty" name:"PublicNetworkLimit"`
}