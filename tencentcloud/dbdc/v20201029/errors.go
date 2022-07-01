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

package v20201029

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// DryRun 操作，代表请求将会是成功的，只是多传了 DryRun 参数。
	DRYRUNOPERATION = "DryRunOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 修改资源信息失败。
	FAILEDOPERATION_MODIFYRESOURCEINFOERROR = "FailedOperation.ModifyResourceInfoError"

	// 访问管控服务失败，请稍后重试。如果持续不成功，请联系客服进行处理。
	FAILEDOPERATION_OSSACCESSERROR = "FailedOperation.OssAccessError"

	// 查询资源信息失败。
	FAILEDOPERATION_QUERYRESOURCEERROR = "FailedOperation.QueryResourceError"

	// 获取独享集群规格信息失败。
	FAILEDOPERATION_QUERYSPECINFOERROR = "FailedOperation.QuerySpecInfoError"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 参数值传入错误。
	INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"

	// 集群的参数非法。
	INVALIDPARAMETERVALUE_RESOURCEPARAMETERERROR = "InvalidParameterValue.ResourceParameterError"

	// 获取资源失败，AppId：{{1}}，ResourceId：{{2}}。
	RESOURCENOTFOUND_FETCHRESOURCEERROR = "ResourceNotFound.FetchResourceError"

	// 获取资源列表失败，ErrMsg：{{1}}。
	RESOURCENOTFOUND_FETCHRESOURCELISTERROR = "ResourceNotFound.FetchResourceListError"

	// 资源状态异常。
	RESOURCEUNAVAILABLE_RESOURCESTATUSABNORMALERROR = "ResourceUnavailable.ResourceStatusAbnormalError"
)
