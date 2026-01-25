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

package v20230517

const (
	// 此产品的特有错误码

	// 操作失败
	FAILEDOPERATION = "FailedOperation"

	// AI任务状态已经是Off
	FAILEDOPERATION_AITASKSTATUSISOFF = "FailedOperation.AITaskStatusIsOff"

	// 不允许删除或更改状态为On的AI任务
	FAILEDOPERATION_AITASKSTATUSISON = "FailedOperation.AITaskStatusIsOn"

	// 数据库错误
	FAILEDOPERATION_DATABASEERROR = "FailedOperation.DatabaseError"

	// 设备响应超时
	FAILEDOPERATION_DEVICERESPONSETIMEOUT = "FailedOperation.DeviceResponseTimeOut"

	// 设备端结果响应超时
	FAILEDOPERATION_DEVICERESULTTIMEOUT = "FailedOperation.DeviceResultTimeOut"

	// 没有该子用户
	FAILEDOPERATION_NOTHAVESUBUSER = "FailedOperation.NotHaveSubUser"

	// 请求超时
	FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"

	// 服务器内部错误
	INTERNALERROR_SERVERINTERNALERROR = "InternalError.ServerInternalError"

	// 无效的输入参数
	INVALIDPARAMETER = "InvalidParameter"

	// 非法的下载URL
	INVALIDPARAMETER_DOWNLOADURLERROR = "InvalidParameter.DownloadUrlError"

	// 下载URL已过期
	INVALIDPARAMETER_DOWNLOADURLHASEXPIRED = "InvalidParameter.DownloadUrlHasExpired"

	// 操作类型不合法
	INVALIDPARAMETER_INVALIDACTION = "InvalidParameter.InvalidAction"

	// 无效的消息体格式
	INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"

	// 无效的通道参数
	INVALIDPARAMETER_INVALIDCHANNELS = "InvalidParameter.InvalidChannels"

	// 无效的生命周期参数
	INVALIDPARAMETER_INVALIDLIFERULEPARAM = "InvalidParameter.InvalidLifeRuleParam"

	// 无效的组织参数
	INVALIDPARAMETER_INVALIDORGANIZATIONPARAM = "InvalidParameter.InvalidOrganizationParam"

	// 无效的输入参数
	INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"

	// 参数格式不对或缺少参数
	INVALIDPARAMETER_INVALIDPARAMETERFORMAT = "InvalidParameter.InvalidParameterFormat"

	// 无效的时间片段
	INVALIDPARAMETER_INVALIDTIMESECTION = "InvalidParameter.InvalidTimeSection"

	// 必要请求头参数不能为空
	INVALIDPARAMETER_REQUIREDHEADERPARAMETEREMPTY = "InvalidParameter.RequiredHeaderParameterEmpty"

	// TaskId不存在
	INVALIDPARAMETER_TASKIDNOTEXIST = "InvalidParameter.TaskIdNotExist"

	// 无效的参数值
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 存在不满足（上云时间段=前一天的录像时间段/上云倍速）的设定，请检查并确认
	INVALIDPARAMETERVALUE_BAKTIMENOTENOUGH = "InvalidParameterValue.BakTimeNotEnough"

	// CallbackURL包含特殊字符
	INVALIDPARAMETERVALUE_CALLBACKURLCONTAINILLEGALCHARACTER = "InvalidParameterValue.CallbackURLContainIllegalCharacter"

	// 包含已存在其他AI任务的ChannelId
	INVALIDPARAMETERVALUE_CHANNELIDALREADYEXISTSINOTHERAITASKS = "InvalidParameterValue.ChannelIdAlreadyExistsInOtherAITasks"

	// ChannelId不能为空
	INVALIDPARAMETERVALUE_CHANNELIDMUSTBENOTEMPTY = "InvalidParameterValue.ChannelIdMustBeNotEmpty"

	// ChannelList包含特殊字符
	INVALIDPARAMETERVALUE_CHANNELLISTCONTAINILLEGALCHARACTER = "InvalidParameterValue.ChannelListContainIllegalCharacter"

	// ChannelList不能为空
	INVALIDPARAMETERVALUE_CHANNELLISTMUSTBENOTEMPTY = "InvalidParameterValue.ChannelListMustBeNotEmpty"

	// ChannelList的数量必须小于1000
	INVALIDPARAMETERVALUE_CHANNELNUMBERMUSTBELESSTHANONETHOUSAND = "InvalidParameterValue.ChannelNumberMustBeLessThanOneThousand"

	// 通道数量超过限制范围
	INVALIDPARAMETERVALUE_CHANNELSEXCEEDSRANGE = "InvalidParameterValue.ChannelsExceedsRange"

	// 包含无效的ChannelId
	INVALIDPARAMETERVALUE_CONTAININVALIDCHANNELID = "InvalidParameterValue.ContainInvalidChannelId"

	// 包含无效的设备ID
	INVALIDPARAMETERVALUE_CONTAININVALIDDEVICEID = "InvalidParameterValue.ContainInvalidDeviceId"

	// 周期超出限制（1-7）
	INVALIDPARAMETERVALUE_DATAOUTOFRANGE = "InvalidParameterValue.DataOutOfRange"

	// 设备ID和通道ID不匹配
	INVALIDPARAMETERVALUE_DEVICEMISMATCHCHANNEL = "InvalidParameterValue.DeviceMismatchChannel"

	// 通道ID不能为空
	INVALIDPARAMETERVALUE_EMPTYCHANNELID = "InvalidParameterValue.EmptyChannelId"

	// 设备ID不能为空
	INVALIDPARAMETERVALUE_EMPTYDEVICEID = "InvalidParameterValue.EmptyDeviceId"

	// 名称不能为空
	INVALIDPARAMETERVALUE_EMPTYNAME = "InvalidParameterValue.EmptyName"

	// 组织ID不能为空
	INVALIDPARAMETERVALUE_EMPTYORGANIZATIONID = "InvalidParameterValue.EmptyOrganizationId"

	// 关联模板不能为空
	INVALIDPARAMETERVALUE_EMPTYTEMPLATEID = "InvalidParameterValue.EmptyTemplateId"

	// 结束时间不能为空
	INVALIDPARAMETERVALUE_ENDTIMEZERO = "InvalidParameterValue.EndTimeZero"

	// OperTimeSlot值有重复
	INVALIDPARAMETERVALUE_HASDUPLICATEOPERTIMESLOT = "InvalidParameterValue.HasDuplicateOperTimeSlot"

	// 通道ID包含了非法字符
	INVALIDPARAMETERVALUE_ILLEGALCHANNELID = "InvalidParameterValue.IllegalChannelId"

	// 描述中包含不符合要求的字符(仅支持中文、英文、数字、_、-)
	INVALIDPARAMETERVALUE_ILLEGALDESCRIBE = "InvalidParameterValue.IllegalDescribe"

	// 设备ID包含了非法字符
	INVALIDPARAMETERVALUE_ILLEGALDEVICEID = "InvalidParameterValue.IllegalDeviceId"

	// 名称中包含不符合要求的字符(仅支持中文、英文、数字、_、-)
	INVALIDPARAMETERVALUE_ILLEGALNAME = "InvalidParameterValue.IllegalName"

	// 组织ID包含了非法字符
	INVALIDPARAMETERVALUE_ILLEGALORGANIZATIONID = "InvalidParameterValue.IllegalOrganizationId"

	// 码流类型包含非法字符(仅支持英文、数字、_、-)
	INVALIDPARAMETERVALUE_ILLEGALSTREAMTYPE = "InvalidParameterValue.IllegalStreamType"

	// 无效的AI任务描述信息
	INVALIDPARAMETERVALUE_INVALIDAITASKDESC = "InvalidParameterValue.InvalidAITaskDesc"

	// 无效的AI任务ID
	INVALIDPARAMETERVALUE_INVALIDAITASKID = "InvalidParameterValue.InvalidAITaskID"

	// 无效的AI任务名称
	INVALIDPARAMETERVALUE_INVALIDAITASKNAME = "InvalidParameterValue.InvalidAITaskName"

	// 无效的AI任务状态
	INVALIDPARAMETERVALUE_INVALIDAITASKSTATUS = "InvalidParameterValue.InvalidAITaskStatus"

	// 无效的设备接入协议
	INVALIDPARAMETERVALUE_INVALIDACCESSPROTOCOL = "InvalidParameterValue.InvalidAccessProtocol"

	// 无效的开始或结束时间
	INVALIDPARAMETERVALUE_INVALIDBEGINANDENDTIME = "InvalidParameterValue.InvalidBeginAndEndTime"

	// 无效的开始时间
	INVALIDPARAMETERVALUE_INVALIDBEGINTIME = "InvalidParameterValue.InvalidBeginTime"

	// 无效的通道ID
	INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"

	// 包含无效的通道ID或者设备ID
	INVALIDPARAMETERVALUE_INVALIDCHANNELIDORDEVICEID = "InvalidParameterValue.InvalidChannelIdOrDeviceId"

	// 通道名称为空或填写有误
	INVALIDPARAMETERVALUE_INVALIDCHANNELNAME = "InvalidParameterValue.InvalidChannelName"

	// 无效的集群ID，长度或内容不符合规则
	INVALIDPARAMETERVALUE_INVALIDCLUSTERID = "InvalidParameterValue.InvalidClusterId"

	// 描述填写有误
	INVALIDPARAMETERVALUE_INVALIDDESCRIPTION = "InvalidParameterValue.InvalidDescription"

	// 无效的AI分析类型
	INVALIDPARAMETERVALUE_INVALIDDETECTTYPE = "InvalidParameterValue.InvalidDetectType"

	// 无效的设备ID
	INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"

	// 设备名称为空或填写有误
	INVALIDPARAMETERVALUE_INVALIDDEVICENAME = "InvalidParameterValue.InvalidDeviceName"

	// 设备密码长度应当大于1位并且小于64位
	INVALIDPARAMETERVALUE_INVALIDDEVICEPASSWORDLENGTH = "InvalidParameterValue.InvalidDevicePasswordLength"

	// 无效的设备状态
	INVALIDPARAMETERVALUE_INVALIDDEVICESTATUS = "InvalidParameterValue.InvalidDeviceStatus"

	// 无效的设备类型
	INVALIDPARAMETERVALUE_INVALIDDEVICETYPE = "InvalidParameterValue.InvalidDeviceType"

	// 无效的域名类型参数
	INVALIDPARAMETERVALUE_INVALIDDOMAINTYPE = "InvalidParameterValue.InvalidDomainType"

	// 无效的使能参数
	INVALIDPARAMETERVALUE_INVALIDENALBEVALUE = "InvalidParameterValue.InvalidEnalbeValue"

	// 无效的结束时间
	INVALIDPARAMETERVALUE_INVALIDENDTIME = "InvalidParameterValue.InvalidEndTime"

	// 无效的副本有效期,范围是1-365天
	INVALIDPARAMETERVALUE_INVALIDEXPIRATION = "InvalidParameterValue.InvalidExpiration"

	// 冷存储至少60天, 并且热存储和冷存储总天数不能超过3650天
	INVALIDPARAMETERVALUE_INVALIDEXPIRATIONRANGE = "InvalidParameterValue.InvalidExpirationRange"

	// 无效的文件格式
	INVALIDPARAMETERVALUE_INVALIDFILETYPE = "InvalidParameterValue.InvalidFileType"

	// 无效的网关ID，长度或内容不符合规则
	INVALIDPARAMETERVALUE_INVALIDGATEWAYID = "InvalidParameterValue.InvalidGatewayId"

	// 无效的网关接入协议
	INVALIDPARAMETERVALUE_INVALIDGATEWAYPROTOCOLTYPE = "InvalidParameterValue.InvalidGatewayProtocolType"

	// 无效的Ipv4地址
	INVALIDPARAMETERVALUE_INVALIDIPV4 = "InvalidParameterValue.InvalidIpv4"

	// 无效的关键字
	INVALIDPARAMETERVALUE_INVALIDKEYWORD = "InvalidParameterValue.InvalidKeyword"

	// 名称为空或填写有误
	INVALIDPARAMETERVALUE_INVALIDNAME = "InvalidParameterValue.InvalidName"

	// 无效的OperTimeSlot格式
	INVALIDPARAMETERVALUE_INVALIDOPERTIMESLOTFORMAT = "InvalidParameterValue.InvalidOperTimeSlotFormat"

	// 组织名称为空或填写有误
	INVALIDPARAMETERVALUE_INVALIDORGNAME = "InvalidParameterValue.InvalidOrgName"

	// 无效的组织ID
	INVALIDPARAMETERVALUE_INVALIDORGANIZATIONID = "InvalidParameterValue.InvalidOrganizationId"

	// 无效的PageNumber
	INVALIDPARAMETERVALUE_INVALIDPAGENUMBER = "InvalidParameterValue.InvalidPageNumber"

	// 页参数取值错误
	INVALIDPARAMETERVALUE_INVALIDPAGEPARAMETER = "InvalidParameterValue.InvalidPageParameter"

	// 无效的PageSize
	INVALIDPARAMETERVALUE_INVALIDPAGESIZE = "InvalidParameterValue.InvalidPageSize"

	// 无效的计划ID
	INVALIDPARAMETERVALUE_INVALIDPLANID = "InvalidParameterValue.InvalidPlanId"

	// 无效的拉流鉴权开关值
	INVALIDPARAMETERVALUE_INVALIDPULLSTATE = "InvalidParameterValue.InvalidPullState"

	// 无效的推流鉴权开关值
	INVALIDPARAMETERVALUE_INVALIDPUSHSTATE = "InvalidParameterValue.InvalidPushState"

	// 无效的RTMP推流AppName
	INVALIDPARAMETERVALUE_INVALIDRTMPAPPNAME = "InvalidParameterValue.InvalidRTMPAppName"

	// 无效的RTMP推流StreamName
	INVALIDPARAMETERVALUE_INVALIDRTMPSTREAMNAME = "InvalidParameterValue.InvalidRTMPStreamName"

	// 无效的录像补录参数
	INVALIDPARAMETERVALUE_INVALIDREPAIRMODE = "InvalidParameterValue.InvalidRepairMode"

	// 无效的取回模式
	INVALIDPARAMETERVALUE_INVALIDRETRIEVALMODE = "InvalidParameterValue.InvalidRetrievalMode"

	// 无效的取回任务ID
	INVALIDPARAMETERVALUE_INVALIDRETRIEVETASKID = "InvalidParameterValue.InvalidRetrieveTaskId"

	// 无效的密钥值，长度或内容不符合规则
	INVALIDPARAMETERVALUE_INVALIDSECRET = "InvalidParameterValue.InvalidSecret"

	// 无效的开始或结束时间
	INVALIDPARAMETERVALUE_INVALIDSTARTTIMEORENDTIME = "InvalidParameterValue.InvalidStartTimeOrEndTime"

	// 状态取值错误
	INVALIDPARAMETERVALUE_INVALIDSTATUS = "InvalidParameterValue.InvalidStatus"

	// 无效的流类型
	INVALIDPARAMETERVALUE_INVALIDSTREAMTYPE = "InvalidParameterValue.InvalidStreamType"

	// 无效的模板ID
	INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"

	// 无效的TemplateTag
	INVALIDPARAMETERVALUE_INVALIDTEMPLATETAG = "InvalidParameterValue.InvalidTemplateTag"

	// 时间格式错误
	INVALIDPARAMETERVALUE_INVALIDTIMEFORMAT = "InvalidParameterValue.InvalidTimeFormat"

	// 无效的截图频率
	INVALIDPARAMETERVALUE_INVALIDTIMEINTERVAL = "InvalidParameterValue.InvalidTimeInterval"

	// 无效时间片段取值
	INVALIDPARAMETERVALUE_INVALIDTIMESECTIONVALUE = "InvalidParameterValue.InvalidTimeSectionValue"

	// 热存储至少1天, 最多3650天
	INVALIDPARAMETERVALUE_INVALIDTRANSITIONRANGE = "InvalidParameterValue.InvalidTransitionRange"

	// 无效的用户参数
	INVALIDPARAMETERVALUE_INVALIDUSERPARAM = "InvalidParameterValue.InvalidUserParam"

	// 无效的用户名，长度或内容不符合规则
	INVALIDPARAMETERVALUE_INVALIDUSERNAME = "InvalidParameterValue.InvalidUsername"

	// 开启录像补录热存不能小于3天
	INVALIDPARAMETERVALUE_LIMITREPAIRTRANSITION = "InvalidParameterValue.LimitRepairTransition"

	// 时间不能跨天
	INVALIDPARAMETERVALUE_NONSAMEDAY = "InvalidParameterValue.NonSameDay"

	// Object包含非法字符
	INVALIDPARAMETERVALUE_OBJECTCONTAINILLEGALCHARACTER = "InvalidParameterValue.ObjectContainIllegalCharacter"

	// OperTimeSlot包含特殊字符
	INVALIDPARAMETERVALUE_OPERTIMESLOTCONTAINILLEGALCHARACTER = "InvalidParameterValue.OperTimeSlotContainIllegalCharacter"

	// OperTimeSlot容量必须小于5
	INVALIDPARAMETERVALUE_OPERTIMESLOTNUMBERMUSTBELESSTHANFIVE = "InvalidParameterValue.OperTimeSlotNumberMustBeLessThanFive"

	// OperTimeSlot开始时间必须小于结束时间
	INVALIDPARAMETERVALUE_OPERTIMESLOTSTARTMUSTBELESSTHANEND = "InvalidParameterValue.OperTimeSlotStartMustBeLessThanEnd"

	// 组织名称不能重复
	INVALIDPARAMETERVALUE_ORGNAMEREPEAT = "InvalidParameterValue.OrgNameRepeat"

	// 组织数量不能大于1000个
	INVALIDPARAMETERVALUE_ORGANIZATIONCOUNTEXCEEDSRANGE = "InvalidParameterValue.OrganizationCountExceedsRange"

	// 时间范围超限
	INVALIDPARAMETERVALUE_OUTOFTIMERANGE = "InvalidParameterValue.OutOfTimeRange"

	// 页面大小超出限制
	INVALIDPARAMETERVALUE_PAGESIZEEXCEEDLIMIT = "InvalidParameterValue.PageSizeExceedLimit"

	// 通道数量一次最多添加5000路
	INVALIDPARAMETERVALUE_PLANCHANNELSEXCEEDSRANGE = "InvalidParameterValue.PlanChannelsExceedsRange"

	// 计划名称不能重复
	INVALIDPARAMETERVALUE_PLANNAMEREPEAT = "InvalidParameterValue.PlanNameRepeat"

	// RTMP推流自定义AppName及StreamName不能重复配置
	INVALIDPARAMETERVALUE_RTMPPUSHSTREAMPARAMREPEAT = "InvalidParameterValue.RTMPPushStreamParamRepeat"

	// 一个取回任务最多添加32个设备通道
	INVALIDPARAMETERVALUE_RETRIEVETASKCHANNELSEXCEEDSRANGE = "InvalidParameterValue.RetrieveTaskChannelsExceedsRange"

	// 开始时间不能大于结束时间
	INVALIDPARAMETERVALUE_STARTOVERENDTIME = "InvalidParameterValue.StartOverEndTime"

	// 开始时间不能大于当前时间
	INVALIDPARAMETERVALUE_STARTOVERNOWTIME = "InvalidParameterValue.StartOverNowTime"

	// 开始时间不能大于等于结束时间
	INVALIDPARAMETERVALUE_STARTTIMEGREATERTHANOREQUALENDTIME = "InvalidParameterValue.StartTimeGreaterThanOrEqualEndTime"

	// 开始时间不能为空
	INVALIDPARAMETERVALUE_STARTTIMEZERO = "InvalidParameterValue.StartTimeZero"

	// AI任务状态不能为空
	INVALIDPARAMETERVALUE_STATUSMUSTBENOTEMPTY = "InvalidParameterValue.StatusMustBeNotEmpty"

	// 不支持任务类型
	INVALIDPARAMETERVALUE_TASKTYPENOTSUPPORTED = "InvalidParameterValue.TaskTypeNotSupported"

	// TemplateTag必须一致
	INVALIDPARAMETERVALUE_TEMPLATETAGMUSTBECONSISTENT = "InvalidParameterValue.TemplateTagMustBeConsistent"

	// 时间间隔不能小于10分钟
	INVALIDPARAMETERVALUE_TIMELESSTHANTENMINUTES = "InvalidParameterValue.TimeLessThanTenMinutes"

	// 描述字段过长
	INVALIDPARAMETERVALUE_TOOLONGDESCRIBE = "InvalidParameterValue.TooLongDescribe"

	// 名称过长
	INVALIDPARAMETERVALUE_TOOLONGNAME = "InvalidParameterValue.TooLongName"

	// 码流类型过长
	INVALIDPARAMETERVALUE_TOOLONGSTREAMTYPE = "InvalidParameterValue.TooLongStreamType"

	// 不支持的操作命令
	INVALIDPARAMETERVALUE_UNSUPPORTOPERATECMD = "InvalidParameterValue.UnSupportOperateCMD"

	// 不支持该倍速
	INVALIDPARAMETERVALUE_UNSUPPORTSCALE = "InvalidParameterValue.UnSupportScale"

	// 不支持的接入类型
	INVALIDPARAMETERVALUE_UNSUPPORTEDACCESSTYPE = "InvalidParameterValue.UnSupportedAccessType"

	// 不支持的云台控制指令
	INVALIDPARAMETERVALUE_UNSUPPORTEDPTZCMD = "InvalidParameterValue.UnSupportedPTZCMD"

	// 不支持的预置位控制指令
	INVALIDPARAMETERVALUE_UNSUPPORTEDPRESETCMD = "InvalidParameterValue.UnSupportedPresetCMD"

	// 不支持的倍速播放参数
	INVALIDPARAMETERVALUE_UNSUPPORTEDSCALEPARAM = "InvalidParameterValue.UnSupportedScaleParam"

	// 不支持的流协议
	INVALIDPARAMETERVALUE_UNSUPPORTEDSTREAMPROTOCOL = "InvalidParameterValue.UnsupportedStreamProtocol"

	// 通道数量超出限制
	LIMITEXCEEDED_CHANNELNUMEXCEEDED = "LimitExceeded.ChannelNumExceeded"

	// 查询的设备数量不能超过200个
	LIMITEXCEEDED_DEICEIDQUANTITYEXCEEDEDLIMIT = "LimitExceeded.DeiceIdQuantityExceededLimit"

	// 查询的国标设备数量不能超过500个
	LIMITEXCEEDED_GBDEVICENUMEXCEEDED = "LimitExceeded.GBDeviceNumExceeded"

	// 时间片段不存在
	MISSINGPARAMETER_EMPTYTIMESECTION = "MissingParameter.EmptyTimeSection"

	// 缺少生命周期参数
	MISSINGPARAMETER_MISSINGLIFERULEPARAM = "MissingParameter.MissingLifeRuleParam"

	// 缺少必须参数
	MISSINGPARAMETER_MISSINGMUSTPARAMETER = "MissingParameter.MissingMustParameter"

	// 缺少取回任务参数
	MISSINGPARAMETER_MISSINGRETRIEVETASKPARAM = "MissingParameter.MissingRetrieveTaskParam"

	// 没有更新信息
	MISSINGPARAMETER_MISSINGUPDATEDINFO = "MissingParameter.MissingUpdatedInfo"

	// 缺少用户信息
	MISSINGPARAMETER_MISSINGUSERINFO = "MissingParameter.MissingUserInfo"

	// 带宽受限
	OPERATIONDENIED_BANDWIDTHLIMITZERO = "OperationDenied.BandwidthLimitZero"

	// 并发下载数超限
	OPERATIONDENIED_CONCURRENTDOWNLOADSOVERLIMIT = "OperationDenied.ConcurrentDownloadsOverLimit"

	// 链接数受限
	OPERATIONDENIED_CONNECTSLIMITZERO = "OperationDenied.ConnectsLimitZero"

	// 新增组织树失败，超出组织树的最大深度
	OPERATIONDENIED_EXCEEDEDMAXIMUMDEPTH = "OperationDenied.ExceededMaximumDepth"

	// 该功能需要申请白名单配置
	OPERATIONDENIED_NOWHITELIST = "OperationDenied.NoWhitelist"

	// 播放鉴权未开启
	OPERATIONDENIED_PLAYAUTHNOTENABLED = "OperationDenied.PlayAuthNotEnabled"

	// 资源不可达，该资源不属于该地域
	REGIONERROR_RESOURCEUNREACHABLE = "RegionError.ResourceUnreachable"

	// 存在重复添加的通道
	RESOURCEINUSE_CHANNELREPEATADD = "ResourceInUse.ChannelRepeatAdd"

	// 计划删除中无法修改
	RESOURCEINUSE_PLANDELETING = "ResourceInUse.PlanDeleting"

	// 该模板被计划关联，不允许删除
	RESOURCEINUSE_PLANLINKTEMPLATE = "ResourceInUse.PlanLinkTemplate"

	// 计划名称不能重复
	RESOURCEINUSE_PLANNAMEREPEAT = "ResourceInUse.PlanNameRepeat"

	// 取回任务正在执行中，无法删除
	RESOURCEINUSE_RETRIEVETASKEXECUTING = "ResourceInUse.RetrieveTaskExecuting"

	// 取回任务名称不能重复
	RESOURCEINUSE_RETRIEVETASKNAMEREPEAT = "ResourceInUse.RetrieveTaskNameRepeat"

	// 模板名称不能重复
	RESOURCEINUSE_TEMPLATENAMEREPEAT = "ResourceInUse.TemplateNameRepeat"

	// 资源不存在
	RESOURCENOTFOUND = "ResourceNotFound"

	// AI任务不存在
	RESOURCENOTFOUND_AITASKNOTEXISTED = "ResourceNotFound.AITaskNotExisted"

	// 通道不存在
	RESOURCENOTFOUND_CHANNELNOTEXIST = "ResourceNotFound.ChannelNotExist"

	// 设备不存在
	RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"

	// 没有找到可用的服务节点
	RESOURCENOTFOUND_NOTFOUNDCLUSTER = "ResourceNotFound.NotFoundCluster"

	// 组织ID不存在
	RESOURCENOTFOUND_ORGANIZATIONIDNOTEXIST = "ResourceNotFound.OrganizationIdNotExist"

	// 计划不存在
	RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"

	// 取回任务不存在
	RESOURCENOTFOUND_RETRIEVETASKNOTEXIST = "ResourceNotFound.RetrieveTaskNotExist"

	// 视频流不存在或已关闭
	RESOURCENOTFOUND_STREAMNOTEXISTORCLOSE = "ResourceNotFound.StreamNotExistOrClose"

	// 模板不存在
	RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"

	// 录像不存在
	RESOURCENOTFOUND_VIDEONOTFOUND = "ResourceNotFound.VideoNotFound"

	// 通道离线
	RESOURCEUNAVAILABLE_CHANNELOFFLINE = "ResourceUnavailable.ChannelOffline"

	// 设备被禁用
	RESOURCEUNAVAILABLE_DEVDISABLE = "ResourceUnavailable.DevDisable"

	// 设备未注册
	RESOURCEUNAVAILABLE_DEVNOREGISTER = "ResourceUnavailable.DevNoRegister"

	// 设备离线
	RESOURCEUNAVAILABLE_DEVOFFLINE = "ResourceUnavailable.DevOffline"

	// 录像已归档
	RESOURCEUNAVAILABLE_VIDEOARCHIVED = "ResourceUnavailable.VideoArchived"

	// 组织下仍有主账号或其他子账号添加的设备，您无法查看或删除组织
	UNSUPPORTEDOPERATION_ORGLINKDEV = "UnsupportedOperation.OrgLinkDev"

	// 组织下有挂靠的组织，不允许删除
	UNSUPPORTEDOPERATION_ORGLINKORG = "UnsupportedOperation.OrgLinkOrg"

	// RTMP推流域名不存在
	UNSUPPORTEDOPERATION_PUSHDOMAINNOTEXIST = "UnsupportedOperation.PushDomainNotExist"

	// Scale和Pos参数不支持同时下发
	UNSUPPORTEDOPERATION_SCALEANDPOSBOTHEXIST = "UnsupportedOperation.ScaleAndPosBothExist"

	// StreamType和Resolution参数不支持同时下发
	UNSUPPORTEDOPERATION_STREAMTYPEORRESOLUTION = "UnsupportedOperation.StreamTypeOrResolution"

	// 该地域未开通内网服务
	UNSUPPORTEDOPERATION_UNOPENEDINTRANETSERVICESINREGION = "UnsupportedOperation.UnopenedIntranetServicesInRegion"
)
