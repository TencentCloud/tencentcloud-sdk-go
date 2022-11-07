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

package v20201214

const (
	// 此产品的特有错误码

	// CDC绑定VPC失败。
	FAILEDOPERATION_CDCBINDVPCFAIL = "FailedOperation.CdcBindVpcFail"

	// 删除site失败。
	FAILEDOPERATION_FAILDELETESITE = "FailedOperation.FailDeleteSite"

	// 该机型暂不支持。
	INVALIDPARAMETER_INSTANCETYPENOTSUPPORT = "InvalidParameter.InstanceTypeNotSupport"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// ccdb返回结果不为0。
	INVALIDPARAMETERVALUE_INVALIDAPPIDFORMAT = "InvalidParameterValue.InvalidAppIdFormat"

	// 本地专用集群COS大小不正确。
	INVALIDPARAMETERVALUE_INVALIDVALUEDEDICATEDCLUSTERCOSSIZE = "InvalidParameterValue.InvalidValueDedicatedClusterCosSize"

	// CBS大小不正确，它必须是40的整数倍。
	INVALIDPARAMETERVALUE_INVALIDVALUEDEDICATEDCLUSTERDATASTEPSIZE = "InvalidParameterValue.InvalidValueDedicatedClusterDataStepSize"

	// region无效。
	INVALIDPARAMETERVALUE_INVALIDVALUEREGION = "InvalidParameterValue.InvalidValueRegion"

	// 超出大小限制。
	INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"

	// 无效Region ID 。
	INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"

	// 参数名过长。
	INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"

	// zone和region不匹配。
	INVALIDPARAMETERVALUE_ZONEMISMATCHREGION = "InvalidParameterValue.ZoneMismatchRegion"

	// 当前可用区暂未支持。
	INVALIDPARAMETERVALUE_ZONENOTSUPPORTED = "InvalidParameterValue.ZoneNotSupported"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 至少输入一个入参。
	MISSINGPARAMETER_ATLEASTONE = "MissingParameter.AtLeastOne"

	// 云硬盘余量不足。
	RESOURCEINSUFFICIENT_CLOUDDISKUNAVAILABLE = "ResourceInsufficient.CloudDiskUnavailable"

	// 该资源没有找到。
	RESOURCENOTFOUND_INVALIDDEDICATEDCLUSTERID = "ResourceNotFound.InvalidDedicatedClusterId"

	// 请确认资源ID 是否存在。
	RESOURCENOTFOUND_INVALIDDEDICATEDCLUSTERORDERID = "ResourceNotFound.InvalidDedicatedClusterOrderId"

	// 该资源没有找到。
	RESOURCENOTFOUND_INVALIDDEDICATEDCLUSTERTYPEID = "ResourceNotFound.InvalidDedicatedClusterTypeId"

	// 站点机房无效。
	RESOURCENOTFOUND_INVALIDSITEID = "ResourceNotFound.InvalidSiteId"

	// 不支持非CUSTOMER类型的app id。
	UNSUPPORTEDOPERATION_NONCUSTOMERAPPIDNOTSUPPORT = "UnsupportedOperation.NonCustomerAppIdNotSupport"
)
