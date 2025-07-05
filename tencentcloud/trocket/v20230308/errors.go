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

package v20230308

const (
	// 此产品的特有错误码

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 实例未就绪，请稍后重试。
	FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"

	// 已达到主题数上限。
	LIMITEXCEEDED_TOPICNUM = "LimitExceeded.TopicNum"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// 客户端不存在。
	RESOURCENOTFOUND_CLIENT = "ResourceNotFound.Client"

	// 接入点不存在。
	RESOURCENOTFOUND_ENDPOINT = "ResourceNotFound.Endpoint"

	// 消费组不存在，请检查后重试。
	RESOURCENOTFOUND_GROUP = "ResourceNotFound.Group"

	// 实例不存在。
	RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"

	// 消息不存在。
	RESOURCENOTFOUND_MESSAGE = "ResourceNotFound.Message"

	// 迁移任务不存在，请检查后重试。
	RESOURCENOTFOUND_MIGRATIONTASK = "ResourceNotFound.MigrationTask"

	// 角色不存在，请检查后重试。
	RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"

	// 主题不存在，请检查后重试。
	RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 实例Topic数量不能调整到已使用额度以下。
	UNSUPPORTEDOPERATION_INSTANCETOPICNUMDOWNGRADE = "UnsupportedOperation.InstanceTopicNumDowngrade"

	// 资源已存在，请检查后重试。
	UNSUPPORTEDOPERATION_RESOURCEALREADYEXISTS = "UnsupportedOperation.ResourceAlreadyExists"
)
