// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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

package v20230427

const (
	// 此产品的特有错误码

	// 未找到此应用该api授权信息
	AUTHFAILURE_APIAUTHORIZATIONNOTFOUND = "AuthFailure.ApiAuthorizationNotFound"

	// token已过期,请重新申请
	AUTHFAILURE_TOKENEXPIRED = "AuthFailure.TokenExpired"

	// 未找到该token信息
	AUTHFAILURE_TOKENNOTFOUND = "AuthFailure.TokenNotFound"

	// Api规则配置信息错误
	FAILEDOPERATION_APIRULECONFIGERROR = "FailedOperation.ApiRuleConfigError"

	// 操作DB失败
	FAILEDOPERATION_DBERROR = "FailedOperation.DBError"

	// 获取标签错误
	FAILEDOPERATION_GETTAGSFAILED = "FailedOperation.GetTagsFailed"

	// 锁定Redis缓存错误
	FAILEDOPERATION_LOCKREDISCACHEFAILED = "FailedOperation.LockRedisCacheFailed"

	// 修改设备信息失败
	FAILEDOPERATION_MODIFYDEVICEERROR = "FailedOperation.ModifyDeviceError"

	// redis操作错误
	FAILEDOPERATION_REDISOPERATIONFAILED = "FailedOperation.RedisOperationFailed"

	// 消息发送错误
	FAILEDOPERATION_SENDMSGERROR = "FailedOperation.SendMsgError"

	// 向IOT服务发送消息失败
	FAILEDOPERATION_SENDMSGTOIOTFAILED = "FailedOperation.SendMsgToIOTFailed"

	// 对象存储 - 初始化失败
	FAILEDOPERATION_STORAGEINITFAILED = "FailedOperation.StorageInitFailed"

	// 对象存储 - url生成失败
	FAILEDOPERATION_URLGENERATEFAILED = "FailedOperation.URLGenerateFailed"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// API网关内部错误
	INTERNALERROR_APIGATEWAYINTERNALERROR = "InternalError.ApiGatewayInternalError"

	// 该api访问路径匹配参数异常，请检查请求路径
	INTERNALERROR_APIREQUESTPATHPARAMETERERROR = "InternalError.ApiRequestPathParameterError"

	// 该应用未关联该项目空间数据，无法获取该项目空间数据
	INTERNALERROR_APPAPINOSPACEPERMISSION = "InternalError.AppApiNoSpacePermission"

	// BIM服务内部错误
	INTERNALERROR_BIMAPPINTERNALERROR = "InternalError.BIMAppInternalError"

	// 业务逻辑错误
	INTERNALERROR_BUSINESSLOGICERROR = "InternalError.BusinessLogicError"

	// manager服务操作失败
	INTERNALERROR_MANAGERSRVFAILED = "InternalError.ManagerSrvFailed"

	// 超时
	INTERNALERROR_TIMEOUT = "InternalError.Timeout"

	// 未知错误
	INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"

	// 更新工作空间下园区信息错误
	INTERNALERROR_UPDATEPARKINFOFAILED = "InternalError.UpdateParkInfoFailed"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 园区编码重复
	INVALIDPARAMETER_DUPLICATEPARKCODE = "InvalidParameter.DuplicateParkCode"

	// 园区名字或园区编码长度超过50字符
	INVALIDPARAMETER_EXCEEDPARKLENGTHLIMIT = "InvalidParameter.ExceedParkLengthLimit"

	// 告警content无效
	INVALIDPARAMETER_INVALIDALARMCONTENT = "InvalidParameter.InvalidAlarmContent"

	// 错误的媒体数据
	INVALIDPARAMETER_INVALIDMEDIADATA = "InvalidParameter.InvalidMediaData"

	// 错误的Meta数据，无法构建握手参数
	INVALIDPARAMETER_INVALIDMETADATA = "InvalidParameter.InvalidMetaData"

	// 请求时间非法
	INVALIDPARAMETER_INVALIDREQUESTTIME = "InvalidParameter.InvalidRequestTime"

	// 参数不匹配
	INVALIDPARAMETER_PARAMNOTMATCH = "InvalidParameter.ParamNotMatch"

	// 状态与处理类型不符
	INVALIDPARAMETER_STATUSNOTMATCHPROCESSTYPE = "InvalidParameter.StatusNotMatchProcessType"

	// 未找到token字段
	INVALIDPARAMETER_TOKENFIELDNOTFOUND = "InvalidParameter.TokenFieldNotFound"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 行政区编码长度为0
	INVALIDPARAMETERVALUE_EMPTYADMINISTRATIVECODE = "InvalidParameterValue.EmptyAdministrativeCode"

	// 告警状态无效
	INVALIDPARAMETERVALUE_INVALIDALARMSTATUS = "InvalidParameterValue.InvalidAlarmStatus"

	// 应用id非法
	INVALIDPARAMETERVALUE_INVALIDAPPLICATIONID = "InvalidParameterValue.InvalidApplicationId"

	// 项目id或构件id取值错误
	INVALIDPARAMETERVALUE_INVALIDELEMENTIDORPROJECTID = "InvalidParameterValue.InvalidElementIdOrProjectId"

	// 网络视频录像机GB28181协议错误
	INVALIDPARAMETERVALUE_INVALIDGB28181CONFIG = "InvalidParameterValue.InvalidGB28181Config"

	// 申请应用token时，传递的nonce参数非法
	INVALIDPARAMETERVALUE_INVALIDNONCE = "InvalidParameterValue.InvalidNonce"

	// 非法protocol协议参数
	INVALIDPARAMETERVALUE_INVALIDPROTOCOL = "InvalidParameterValue.InvalidProtocol"

	// 请求时间非法
	INVALIDPARAMETERVALUE_INVALIDREQUESTTIME = "InvalidParameterValue.InvalidRequestTime"

	// 保存类型参数错误
	INVALIDPARAMETERVALUE_INVALIDSAVETYPE = "InvalidParameterValue.InvalidSaveType"

	// 签名非法
	INVALIDPARAMETERVALUE_INVALIDSIGNATURE = "InvalidParameterValue.InvalidSignature"

	// 非法StreamId，StreamId仅支持0或1
	INVALIDPARAMETERVALUE_INVALIDSTREAMID = "InvalidParameterValue.InvalidStreamId"

	// metadata租户信息无效
	INVALIDPARAMETERVALUE_INVALIDTENANTID = "InvalidParameterValue.InvalidTenantId"

	// 时间范围参数错误
	INVALIDPARAMETERVALUE_INVALIDTIMERANGE = "InvalidParameterValue.InvalidTimeRange"

	// 播放速率错误，需要等于0.5、1、1.5、2、4、8或16
	INVALIDPARAMETERVALUE_INVALIDVIDEOPLAYRATE = "InvalidParameterValue.InvalidVideoPlayRate"

	// 错误的工作空间Id
	INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"

	// api接口批量最大个数超限
	LIMITEXCEEDED_APILIMITEXCEEDED = "LimitExceeded.ApiLimitExceeded"

	// 设备数量超过限制
	LIMITEXCEEDED_DEVICELIMITEXCEEDED = "LimitExceeded.DeviceLimitExceeded"

	// 视频流超过阈值
	LIMITEXCEEDED_VIDEOSTREAMTHRESHOLDEXCEEDED = "LimitExceeded.VideoStreamThresholdExceeded"

	// 视频转码超出阈值
	LIMITEXCEEDED_VIDEOTRANSCODE = "LimitExceeded.VideoTranscode"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 告警id为空
	MISSINGPARAMETER_EMPTYALARMID = "MissingParameter.EmptyAlarmId"

	// stream参数缺失
	MISSINGPARAMETER_EMPTYSTREAM = "MissingParameter.EmptyStream"

	// WID参数缺失
	MISSINGPARAMETER_EMPTYWID = "MissingParameter.EmptyWID"

	// Api没有操作权限
	OPERATIONDENIED_APIPERMISSIONDENIED = "OperationDenied.ApiPermissionDenied"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 告警id找不到
	RESOURCENOTFOUND_ALARMIDNOTFOUND = "ResourceNotFound.AlarmIDNotFound"

	// 告警id不存在
	RESOURCENOTFOUND_ALARMIDNOTEXIST = "ResourceNotFound.AlarmIdNotExist"

	// 摄像头基础信息不存在
	RESOURCENOTFOUND_CAMERABASEINFONOTEXIST = "ResourceNotFound.CameraBaseInfoNotExist"

	// 摄像头信息不存在
	RESOURCENOTFOUND_CAMERAINFONOTEXIST = "ResourceNotFound.CameraInfoNotExist"

	// 摄像头状态不存在
	RESOURCENOTFOUND_CAMERASTATUSNOTEXIST = "ResourceNotFound.CameraStatusNotExist"

	// 未查询到构件信息
	RESOURCENOTFOUND_ELEMENTNOTFOUND = "ResourceNotFound.ElementNotFound"

	// 空设备列表
	RESOURCENOTFOUND_EMPTYDEVICELIST = "ResourceNotFound.EmptyDeviceList"

	// 网络视频录像机或中心级视频网络存储设备相关配置信息缺失
	RESOURCENOTFOUND_NVRORCVRCONFIGNOTEXIST = "ResourceNotFound.NVROrCVRConfigNotExist"

	// srs_hook服务节点缺失
	RESOURCENOTFOUND_SRSHOOKSERVICENODE = "ResourceNotFound.SRSHookServiceNode"

	// 该deviceId的srs信息缺失
	RESOURCENOTFOUND_SRSNOTEXIST = "ResourceNotFound.SRSNotExist"

	// 视频流缺失
	RESOURCENOTFOUND_STREAMNOTEXIST = "ResourceNotFound.StreamNotExist"

	// VideoPush服务节点缺失
	RESOURCENOTFOUND_VIDEOPUSHSERVICENODE = "ResourceNotFound.VideoPushServiceNode"

	// 设备WId不存在
	RESOURCENOTFOUND_WIDNOTEXIST = "ResourceNotFound.WIDNotExist"

	// 设备已离线
	RESOURCEUNAVAILABLE_DEVICEOFFLINE = "ResourceUnavailable.DeviceOffline"

	// wId 未找到或被应用api未被授权
	UNAUTHORIZEDOPERATION_APIAUTHFAILED = "UnauthorizedOperation.APIAuthFailed"

	// 应用api的数据授权未配置，请先配置权限
	UNAUTHORIZEDOPERATION_UNAUTHORIZEDAPI = "UnauthorizedOperation.UnauthorizedApi"

	// 未推流
	UNSUPPORTEDOPERATION_NOTSTREAMING = "UnsupportedOperation.NotStreaming"

	// 不支持的控制指令
	UNSUPPORTEDOPERATION_UNSUPPORTEDCMD = "UnsupportedOperation.UnsupportedCMD"
)
