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

package v20200721

const (
	// 此产品的特有错误码

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 依赖的db出错。
	INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"

	// 系统错误。
	INTERNALERROR_SYSTEM = "InternalError.System"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数错误。
	INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"
)
