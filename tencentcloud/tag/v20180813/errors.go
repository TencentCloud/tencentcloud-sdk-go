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

package v20180813

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 未通过CAM鉴权。
	AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 项目已禁用。
	FAILEDOPERATION_PROJECTDISABLED = "FailedOperation.ProjectDisabled"

	// 项目数超过限制。
	FAILEDOPERATION_PROJECTNUMEXCEED = "FailedOperation.ProjectNumExceed"

	// 单次请求的资源appId必须相同。
	FAILEDOPERATION_RESOURCEAPPIDNOTSAME = "FailedOperation.ResourceAppIdNotSame"

	// 资源标签正在处理中。
	FAILEDOPERATION_RESOURCETAGPROCESSING = "FailedOperation.ResourceTagProcessing"

	// 标签已经关联配额。
	FAILEDOPERATION_TAGATTACHEDQUOTA = "FailedOperation.TagAttachedQuota"

	// 已关联资源的标签无法删除。
	FAILEDOPERATION_TAGATTACHEDRESOURCE = "FailedOperation.TagAttachedResource"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// PaginationToken参数非法。
	INVALIDPARAMETER_PAGINATIONTOKENINVALID = "InvalidParameter.PaginationTokenInvalid"

	// 项目名称已存在。
	INVALIDPARAMETER_PROJECTNAMEEXISTED = "InvalidParameter.ProjectNameExisted"

	// 系统预留标签键 qcloud、tencent和project 禁止创建。
	INVALIDPARAMETER_RESERVEDTAGKEY = "InvalidParameter.ReservedTagKey"

	// Tag参数错误。
	INVALIDPARAMETER_TAG = "InvalidParameter.Tag"

	// 当前业务不支持标签操作。
	INVALIDPARAMETER_UNSUPPORTEDSERVICE = "InvalidParameter.UnsupportedService"

	// DeleteTags中不能包含ReplaceTags或AddTags中元素。
	INVALIDPARAMETERVALUE_DELETETAGSPARAMERROR = "InvalidParameterValue.DeleteTagsParamError"

	// offset error。
	INVALIDPARAMETERVALUE_OFFSETINVALID = "InvalidParameterValue.OffsetInvalid"

	// 地域错误。
	INVALIDPARAMETERVALUE_REGIONINVALID = "InvalidParameterValue.RegionInvalid"

	// 系统预留标签键 qcloud、tencent和project 禁止创建。
	INVALIDPARAMETERVALUE_RESERVEDTAGKEY = "InvalidParameterValue.ReservedTagKey"

	// 资源描述错误。
	INVALIDPARAMETERVALUE_RESOURCEDESCRIPTIONERROR = "InvalidParameterValue.ResourceDescriptionError"

	// 资源Id错误。
	INVALIDPARAMETERVALUE_RESOURCEIDINVALID = "InvalidParameterValue.ResourceIdInvalid"

	// 资源前缀错误。
	INVALIDPARAMETERVALUE_RESOURCEPREFIXINVALID = "InvalidParameterValue.ResourcePrefixInvalid"

	// 业务类型错误。
	INVALIDPARAMETERVALUE_SERVICETYPEINVALID = "InvalidParameterValue.ServiceTypeInvalid"

	// TagFilters参数错误。
	INVALIDPARAMETERVALUE_TAGFILTERS = "InvalidParameterValue.TagFilters"

	// 过滤标签键对应标签值达到上限数 6。
	INVALIDPARAMETERVALUE_TAGFILTERSLENGTHEXCEEDED = "InvalidParameterValue.TagFiltersLengthExceeded"

	// 标签键包含非法字符。
	INVALIDPARAMETERVALUE_TAGKEYCHARACTERILLEGAL = "InvalidParameterValue.TagKeyCharacterIllegal"

	// TagList中存在重复的TagKey。
	INVALIDPARAMETERVALUE_TAGKEYDUPLICATE = "InvalidParameterValue.TagKeyDuplicate"

	// 标签键不能为空值。
	INVALIDPARAMETERVALUE_TAGKEYEMPTY = "InvalidParameterValue.TagKeyEmpty"

	// 标签键长度超过限制。
	INVALIDPARAMETERVALUE_TAGKEYLENGTHEXCEEDED = "InvalidParameterValue.TagKeyLengthExceeded"

	// 标签值包含非法字符。
	INVALIDPARAMETERVALUE_TAGVALUECHARACTERILLEGAL = "InvalidParameterValue.TagValueCharacterIllegal"

	// 标签值不能为空值。
	INVALIDPARAMETERVALUE_TAGVALUEEMPTY = "InvalidParameterValue.TagValueEmpty"

	// 标签值长度超过限制。
	INVALIDPARAMETERVALUE_TAGVALUELENGTHEXCEEDED = "InvalidParameterValue.TagValueLengthExceeded"

	// Uin参数不合法。
	INVALIDPARAMETERVALUE_UININVALID = "InvalidParameterValue.UinInvalid"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 超过配额限制。
	LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"

	// 资源关联的标签数超过限制。
	LIMITEXCEEDED_RESOURCEATTACHEDTAGS = "LimitExceeded.ResourceAttachedTags"

	// 单次请求的资源数达到上限。
	LIMITEXCEEDED_RESOURCENUMPERREQUEST = "LimitExceeded.ResourceNumPerRequest"

	// 用户创建标签键达到上限数 1000。
	LIMITEXCEEDED_TAGKEY = "LimitExceeded.TagKey"

	// 单次请求的标签数超过上限。
	LIMITEXCEEDED_TAGNUMPERREQUEST = "LimitExceeded.TagNumPerRequest"

	// 单个标签键对应标签值达到上限数 1000。
	LIMITEXCEEDED_TAGVALUE = "LimitExceeded.TagValue"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 标签已存在。
	RESOURCEINUSE_TAGDUPLICATE = "ResourceInUse.TagDuplicate"

	// 对应的标签键和资源已关联。
	RESOURCEINUSE_TAGKEYATTACHED = "ResourceInUse.TagKeyAttached"

	// 资源关联的标签键不存在。
	RESOURCENOTFOUND_ATTACHEDTAGKEYNOTFOUND = "ResourceNotFound.AttachedTagKeyNotFound"

	// 记录不存在。
	RESOURCENOTFOUND_RECORDNOTFOUND = "ResourceNotFound.RecordNotFound"

	// 标签不存在。
	RESOURCENOTFOUND_TAGNONEXIST = "ResourceNotFound.TagNonExist"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"
)
