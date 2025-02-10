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

package v20240516

const (
	// 此产品的特有错误码

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// FailedOperation.InstanceNotReady
	FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"

	// FailedOperation.PublicKeyVerifyFailed
	FAILEDOPERATION_PUBLICKEYVERIFYFAILED = "FailedOperation.PublicKeyVerifyFailed"

	// LimitExceeded.TopicNum
	LIMITEXCEEDED_TOPICNUM = "LimitExceeded.TopicNum"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// ResourceNotFound.Instance
	RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"

	// ResourceNotFound.Role
	RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"

	// ResourceNotFound.Topic
	RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// UnsupportedOperation.ResourceAlreadyExists
	UNSUPPORTEDOPERATION_RESOURCEALREADYEXISTS = "UnsupportedOperation.ResourceAlreadyExists"
)
