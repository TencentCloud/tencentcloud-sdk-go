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

package v20200210

const (
	// 此产品的特有错误码

	// 重复账号。
	FAILEDOPERATION_DUPLICATEDACCOUNT = "FailedOperation.DuplicatedAccount"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 内部数据库访问失败。
	INTERNALERROR_DBERROR = "InternalError.DBError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 实例不存在。
	INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 账号不存在。
	INVALIDPARAMETERVALUE_ACCOUNTNOTEXIST = "InvalidParameterValue.AccountNotExist"

	// 实例不存在。
	INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"

	// 查询记录不存在。
	INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 超出数量限制。
	LIMITEXCEEDED_OUTOFCOUNTLIMIT = "LimitExceeded.OutOfCountLimit"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"
)
