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

package v20180129

const (
	// 此产品的特有错误码

	// 数据处理中。
	INTERNALERROR_DATAINPROCESSING = "InternalError.DataInProcessing"

	// 内部接口错误。
	INTERNALERROR_INNERSERVERFAILED = "InternalError.InnerServerFailed"

	// 元数据操作失败。
	INTERNALERROR_METADATAOPFAILED = "InternalError.MetaDataOpFailed"

	// 参数错误
	INVALIDPARAMETER = "InvalidParameter"

	// 参数取值错误
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 缺少参数错误
	MISSINGPARAMETER = "MissingParameter"

	// 资源不存在
	RESOURCENOTFOUND = "ResourceNotFound"

	// 未授权操作
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
)
