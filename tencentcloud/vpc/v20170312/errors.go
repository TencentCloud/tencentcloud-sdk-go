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

package v20170312

const (
	// 此产品的特有错误码

	// 调用VPCOSS失败
	FAILEDOPERATION_CALLVPCOSSFAILED = "FailedOperation.CallVpcOssFailed"

	// 指定过滤条件不存在。
	INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"

	// 指定过滤选项值不是列表。
	INVALIDPARAMETER_FILTERVALUESNOTLIST = "InvalidParameter.FilterValuesNotList"

	// 参数值超出限制。
	INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"

	// 入参格式不合法。
	INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"

	// 参数值不在指定范围。
	INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"
)
