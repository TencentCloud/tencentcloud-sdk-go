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

package v20180312

const (
	// 此产品的特有错误码

	// 内部错误。
	INTERNALERROR = "InternalError"

	// url已经添加。
	INVALIDPARAMETER_DUPLICATE = "InvalidParameter.Duplicate"

	// 错误的数据格式。
	INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"

	// 请求的数据不存在。
	INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"

	// 站点已添加至其他监控，不能重复添加。
	LIMITEXCEEDED_MONITORQUOTA = "LimitExceeded.MonitorQuota"

	// 超出购买的扫描次数。
	LIMITEXCEEDED_SCANQUOTA = "LimitExceeded.ScanQuota"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
)
