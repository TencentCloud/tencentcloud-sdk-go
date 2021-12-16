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

package v20210303

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 帐号已欠费。
	FAILEDOPERATION_ARREARSERROR = "FailedOperation.ArrearsError"

	// 今日次数达到限制。
	FAILEDOPERATION_COUNTLIMITERROR = "FailedOperation.CountLimitError"

	// 数据明文内容过长。
	FAILEDOPERATION_DATAINFOTOOLONG = "FailedOperation.DataInfoTooLong"

	// 文件下载失败。
	FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"

	// 文件编码格式错误。
	FAILEDOPERATION_FILEENCODINDFORMATERROR = "FailedOperation.FileEncodindFormatError"

	// 文件读取失败。
	FAILEDOPERATION_FILEREADFAILED = "FailedOperation.FileReadFailed"

	// 哈希不匹配。
	FAILEDOPERATION_HASHNOMATCH = "FailedOperation.HashNoMatch"

	// 上链失败。
	FAILEDOPERATION_ONCHAINFAILURE = "FailedOperation.OnChainFailure"

	// 查询无记录。
	FAILEDOPERATION_QUERYNORECORD = "FailedOperation.QueryNoRecord"

	// 敏感数据。
	FAILEDOPERATION_SENSITIVEDATA = "FailedOperation.SensitiveData"

	// 未知错误。
	FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 用户数据无效。
	INVALIDPARAMETER_ACCOUNTINFOINVALID = "InvalidParameter.AccountInfoInvalid"

	// 文件后缀不正确。
	INVALIDPARAMETER_INVALIDFILESUFFIX = "InvalidParameter.InvalidFileSuffix"

	// 参数值错误。
	INVALIDPARAMETER_INVALIDPARAMETERVALUE = "InvalidParameter.InvalidParameterValue"

	// URL无效。
	INVALIDPARAMETER_INVALIDURL = "InvalidParameter.InvalidURL"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 数据明文内容过长。
	INVALIDPARAMETERVALUE_DATAINFOTOOLONG = "InvalidParameterValue.DataInfoTooLong"

	// 哈希不匹配。
	INVALIDPARAMETERVALUE_HASHNOMATCH = "InvalidParameterValue.HashNoMatch"

	// 文件后缀不正确。
	INVALIDPARAMETERVALUE_INVALIDFILESUFFIX = "InvalidParameterValue.InvalidFileSuffix"

	// URL无效。
	INVALIDPARAMETERVALUE_INVALIDURL = "InvalidParameterValue.InvalidURL"

	// 文件内容太大。
	INVALIDPARAMETERVALUE_TOOLARGEFILEERROR = "InvalidParameterValue.TooLargeFileError"

	// 文件内容太大。
	LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"

	// 余额不足。
	RESOURCEINSUFFICIENT_LOWBALANCE = "ResourceInsufficient.LowBalance"

	// 文件下载失败。
	RESOURCENOTFOUND_DOWNLOADERROR = "ResourceNotFound.DownLoadError"

	// 资源未完成开通。
	RESOURCEUNAVAILABLE_RESOURCENOTOPENED = "ResourceUnavailable.ResourceNotOpened"
)
