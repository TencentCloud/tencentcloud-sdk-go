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

package v20210515

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 渠道商服务时间已到期。
	AUTHFAILURE_AGENTEXPIRED = "AuthFailure.AgentExpired"

	// 当前没有创建任何企业。
	AUTHFAILURE_CORPEMPTY = "AuthFailure.CorpEmpty"

	// 企业服务时间已到期。
	AUTHFAILURE_CORPEXPIRED = "AuthFailure.CorpExpired"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"
)
