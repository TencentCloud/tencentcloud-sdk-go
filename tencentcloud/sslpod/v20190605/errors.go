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

package v20190605

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 重复添加。
	FAILEDOPERATION_REPETITIONADD = "FailedOperation.RepetitionAdd"

	// 解析域名失败。
	FAILEDOPERATION_RESOLVEDOMAINFAILED = "FailedOperation.ResolveDomainFailed"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 无效的域名。
	INVALIDPARAMETER_INVALIDDOMAIN = "InvalidParameter.InvalidDomain"

	// 无效的IP。
	INVALIDPARAMETER_INVALIDIP = "InvalidParameter.InvalidIP"

	// 无效的端口。
	INVALIDPARAMETER_INVALIDPORT = "InvalidParameter.InvalidPort"

	// 无效的监控类型。
	INVALIDPARAMETER_INVALIDSERVERTYPE = "InvalidParameter.InvalidServerType"

	// 标签不合规。
	INVALIDPARAMETER_INVALIDTAGNAME = "InvalidParameter.InvalidTagName"

	// 添加标签过多。
	INVALIDPARAMETER_TOOMANYTAG = "InvalidParameter.TooManyTag"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 无效的通知类型。
	INVALIDPARAMETERVALUE_INVALIDNOTICETYPE = "InvalidParameterValue.InvalidNoticeType"

	// 无效的搜索类型。
	INVALIDPARAMETERVALUE_INVALIDSEARCHTYPE = "InvalidParameterValue.InvalidSearchType"

	// 套餐管理额度已满。
	LIMITEXCEEDED_ADDEXCEEDED = "LimitExceeded.AddExceeded"

	// 套餐监控额度已满。
	LIMITEXCEEDED_MONITOREXCEEDED = "LimitExceeded.MonitorExceeded"

	// 产品不存在。
	RESOURCENOTFOUND_PRODUCT = "ResourceNotFound.Product"
)
