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

package v20180301

const (
	// 此产品的特有错误码

	// 认证失败。
	AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 该操作不支持跨算法模型版本。
	FAILEDOPERATION_ACROSSVERSIONSERROR = "FailedOperation.AcrossVersionsError"

	// 操作冲突，请勿同时操作相同的Person。
	FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"

	// 增加人脸不支持并发操作。
	FAILEDOPERATION_CREATEFACECONCURRENT = "FailedOperation.CreateFaceConcurrent"

	// 同一人员库中自定义描述字段不可重复。
	FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"

	// 人脸图片质量不符要求。
	FAILEDOPERATION_FACEQUALITYNOTQUALIFIED = "FailedOperation.FaceQualityNotQualified"

	// 人脸框大小小于MinFaceSize设置，人脸被过滤。
	FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"

	// 指定人员库的升级操作无法执行。
	FAILEDOPERATION_GROUPCANNOTUPGRADE = "FailedOperation.GroupCannotUpgrade"

	// 当前组正处于删除状态，请等待。
	FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"

	// 组中已包含对应的人员Id。
	FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"

	// 组中不包含对应的人员Id。
	FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"

	// 图片解码失败。
	FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"

	// 图片下载错误。
	FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"

	// 人脸检测失败。
	FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"

	// 图片分辨率过大。
	FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"

	// 图片短边分辨率小于64。
	FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"

	// base64编码后的图片数据大小不超过5M。
	FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"

	// 任务无法回滚。
	FAILEDOPERATION_JOBCANNNOTROLLBACK = "FailedOperation.JobCannnotRollback"

	// 请求频率超过限制。
	FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"

	// 后端服务超时。
	FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"

	// Rpc调用失败。
	FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"

	// 检索人脸个数超过限制。
	FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"

	// 算法服务异常，请重试。
	FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"

	// 内部错误。
	FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"

	// 人员库升级任务不存在。
	FAILEDOPERATION_UPGRADEJOBIDNOTEXIST = "FailedOperation.UpgradeJobIdNotExist"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数不合法。
	INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"

	// 账号脸数量超出限制。
	INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"

	// 删除人脸数量超出限制。每个人员至少需要包含一张人脸。
	INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"

	// FaceMatchThreshold参数不合法。
	INVALIDPARAMETERVALUE_FACEMATCHTHRESHOLDILLEGAL = "InvalidParameterValue.FaceMatchThresholdIllegal"

	// 算法模型版本不合法。
	INVALIDPARAMETERVALUE_FACEMODELVERSIONILLEGAL = "InvalidParameterValue.FaceModelVersionIllegal"

	// 人员库自定义描述字段数组长度超过限制。最多可以创建5个。
	INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"

	// 人员库自定义描述字段名称不可重复。
	INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"

	// 人员库自定义描述字段名称包含非法字符。人员库自定义描述字段名称只支持中英文、-、_、数字。
	INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"

	// 人员库自定义描述字段名称长度超出限制。
	INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"

	// 人员库人脸数量超出限制。
	INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"

	// 人员库ID已经存在。人员库ID不可重复。
	INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"

	// 人员库ID包含非法字符。人员库ID只支持英文、数字、-%@#&_。
	INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"

	// 人员库ID不存在。
	INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"

	// 对应的人员库ID在库中不存在。
	INVALIDPARAMETERVALUE_GROUPIDNOTEXISTS = "InvalidParameterValue.GroupIdNotExists"

	// 人员库ID超出长度限制。
	INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"

	// 传入的人员库列表超过限制。
	INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"

	// 人员库名称已经存在。人员库名称不可重复。
	INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"

	// 人员库名称包含非法字符。人员库名称只支持中英文、-、_、数字。
	INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"

	// 人员库名称超出长度限制。
	INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"

	// 人员库数量超出限制。如需增加，请联系我们。
	INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"

	// 人员库数量超出限制。单个人员最多可被添加至100个人员库。
	INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"

	// 人员库备注包含非法字符。人员库备注只支持中英文、-、_、数字。
	INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"

	// 人员库备注超出长度限制。
	INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"

	// 图片为空。
	INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"

	// 图片为空。
	INVALIDPARAMETERVALUE_IMAGEEMPTYERROR = "InvalidParameterValue.ImageEmptyError"

	// 返回数量超出限制。
	INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"

	// 指定分组中没有人脸。
	INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"

	// 图片中没有人脸。
	INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"

	// 起始序号过大。请检查需要请求的数组长度。
	INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"

	// 人员自定义描述字段数组长度超过限制。最多5个。
	INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"

	// 人员自定义描述字段名称不可重复。
	INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"

	// 人员自定义描述字段名称包含非法字符。人员自定义描述字段名称只支持中英文、-、_、数字。
	INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"

	// 人员自定义描述字段名称长度超出限制。
	INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"

	// 组中已包含对应的人员Id。
	INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"

	// 人员人脸数量超出限制。单个人员最多可以包含五张人脸。
	INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"

	// 人员性别设置出错。0代表未填写，1代表男性，2代表女性。
	INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"

	// 人员ID已经存在。人员ID不可重复。
	INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"

	// 人员ID包含非法字符。人员ID只支持英文、数字、-%@#&_。
	INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"

	// 人员ID不存在。
	INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"

	// 人员ID超出长度限制。
	INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"

	// 人员名称包含非法字符。人员名称只支持中英文、-、_、数字。
	INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"

	// 人员名称超出长度限制。
	INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"

	// QualityControl参数不合法。
	INVALIDPARAMETERVALUE_QUALITYCONTROLILLEGAL = "InvalidParameterValue.QualityControlIllegal"

	// 搜索的人员数目超过限制。
	INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"

	// UniquePersonControl参数不合法。
	INVALIDPARAMETERVALUE_UNIQUEPERSONCONTROLILLEGAL = "InvalidParameterValue.UniquePersonControlIllegal"

	// 该操作不支持算法模型版本2.0及以下版本。
	INVALIDPARAMETERVALUE_UNSUPPORTEDGROUPFACEMODELVERSION = "InvalidParameterValue.UnsupportedGroupFaceModelVersion"

	// 一次最多上传四张人脸。
	INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"

	// URL格式不合法。
	INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 人脸个数超过限制。
	LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"

	// 必选参数为空。
	MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"

	// 请求的次数超过了频率限制。
	REQUESTLIMITEXCEEDED = "RequestLimitExceeded"

	// 帐号已欠费。
	RESOURCEUNAVAILABLE_CHARGESTATUSEXCEPTION = "ResourceUnavailable.ChargeStatusException"

	// 资源正在发货中。
	RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"

	// 帐号已被冻结。
	RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"

	// 获取认证信息失败。
	RESOURCEUNAVAILABLE_GETAUTHINFOERROR = "ResourceUnavailable.GetAuthInfoError"

	// 帐号已欠费。
	RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"

	// 余额不足。
	RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"

	// 计费状态未知，请确认是否已在控制台开通服务。
	RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"

	// 服务未开通。
	RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"

	// 资源已被回收。
	RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"

	// 帐号已停服。
	RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"

	// 计费状态未知。
	RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"

	// 计费状态异常。
	RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"

	// 未知方法名。
	UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
)
