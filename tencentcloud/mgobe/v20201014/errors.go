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

package v20201014

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// DryRun 操作，代表请求将会是成功的，只是多传了 DryRun 参数。
	DRYRUNOPERATION = "DryRunOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 连接信息为空。
	FAILEDOPERATION_ACCESSACCESSINFOEMPTY = "FailedOperation.AccessAccessInfoEmpty"

	// 添加COMM连接信息失败。
	FAILEDOPERATION_ACCESSADDCOMMCONNERR = "FailedOperation.AccessAddCommConnErr"

	// 添加心跳连接信息失败。
	FAILEDOPERATION_ACCESSADDHEARTCONNERR = "FailedOperation.AccessAddHeartConnErr"

	// 添加Relay连接信息失败。
	FAILEDOPERATION_ACCESSADDRELAYCONNERR = "FailedOperation.AccessAddRelayConnErr"

	// 获取Token失败。
	FAILEDOPERATION_ACCESSCMDGETTOKENERR = "FailedOperation.AccessCmdGetTokenErr"

	// 命令字无效错误。
	FAILEDOPERATION_ACCESSCMDINVALIDERR = "FailedOperation.AccessCmdInvalidErr"

	// Token无效或过期。
	FAILEDOPERATION_ACCESSCMDINVALIDTOKEN = "FailedOperation.AccessCmdInvalidToken"

	// Token即将过期。
	FAILEDOPERATION_ACCESSCMDTOKENPREEXPIRE = "FailedOperation.AccessCmdTokenPreExpire"

	// 查找连接信息出错。
	FAILEDOPERATION_ACCESSCONNERR = "FailedOperation.AccessConnErr"

	// 获取COMM连接信息失效。
	FAILEDOPERATION_ACCESSGETCOMMCONNECTERR = "FailedOperation.AccessGetCommConnectErr"

	// 获取RELAY连接信息失效。
	FAILEDOPERATION_ACCESSGETRELAYCONNECTERR = "FailedOperation.AccessGetRelayConnectErr"

	// 获取Relay的RS_IP或RS_PORT出错。
	FAILEDOPERATION_ACCESSGETRSIPERR = "FailedOperation.AccessGetRsIpErr"

	// 心跳包解析出错。
	FAILEDOPERATION_ACCESSHEARTBODYPARSEERR = "FailedOperation.AccessHeartBodyParseErr"

	// 登录用户中心回包解析出错。
	FAILEDOPERATION_ACCESSLOGINBODYPARSEERR = "FailedOperation.AccessLoginBodyParseErr"

	// 转发SVR名字错误，不是relay_svr或state_svr。
	FAILEDOPERATION_ACCESSNOERELAYORSTATESVR = "FailedOperation.AccessNoeRelayOrStateSvr"

	// 用户已经登录，不能重复登录。
	FAILEDOPERATION_ACCESSPLAYERDUPLICATELOGIN = "FailedOperation.AccessPlayerDuplicateLogin"

	// PUSH序列化包失败。
	FAILEDOPERATION_ACCESSPUSHSERIALIZEERR = "FailedOperation.AccessPushSerializeErr"

	// 计费类型相关错误。
	FAILEDOPERATION_BILLINGERROR = "FailedOperation.BillingError"

	// 非法命令字。
	FAILEDOPERATION_CMDINVALID = "FailedOperation.CmdInvalid"

	// 解散房间失败。
	FAILEDOPERATION_DISMISSROOMFAILED = "FailedOperation.DismissRoomFailed"

	// 发送消息频率达到限制。
	FAILEDOPERATION_GROUPCHATFREQUENCYLIMIT = "FailedOperation.GroupChatFrequencyLimit"

	// 无权限修改队组队长。
	FAILEDOPERATION_GROUPMODIFYOWNERNOPERMISSION = "FailedOperation.GroupModifyOwnerNoPermission"

	// 队组不存在。
	FAILEDOPERATION_GROUPNOTEXIST = "FailedOperation.GroupNotExist"

	// 队组操作失败。
	FAILEDOPERATION_GROUPOPERATIONFAILED = "FailedOperation.GroupOperationFailed"

	// 对组中人数超过限制。
	FAILEDOPERATION_GROUPPLAYERNUMLIMITEXCEED = "FailedOperation.GroupPlayerNumLimitExceed"

	// 没有权限移除玩家。
	FAILEDOPERATION_GROUPREMOVEPLAYERNOPERMISSION = "FailedOperation.GroupRemovePlayerNoPermission"

	// 服务器内部错误。
	FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"

	// 无效的修改选项。
	FAILEDOPERATION_INVALIDCHANGEOPTION = "FailedOperation.InvalidChangeOption"

	// 参数错误change_room_option_list。
	FAILEDOPERATION_INVALIDCHANGEROOMOPTION = "FailedOperation.InvalidChangeRoomOption"

	// 业务参数错误。
	FAILEDOPERATION_INVALIDPARAMS = "FailedOperation.InvalidParams"

	// 参数错误create_room_type。
	FAILEDOPERATION_INVALIDPARAMSCREATEROOMTYPE = "FailedOperation.InvalidParamsCreateRoomType"

	// 参数错误device_id。
	FAILEDOPERATION_INVALIDPARAMSDEVICEID = "FailedOperation.InvalidParamsDeviceId"

	// 参数错误game_id。
	FAILEDOPERATION_INVALIDPARAMSGAMEID = "FailedOperation.InvalidParamsGameId"

	// 队组自定义属性参数错误。
	FAILEDOPERATION_INVALIDPARAMSGROUPCUSTOMPROPERTIES = "FailedOperation.InvalidParamsGroupCustomProperties"

	// 队组id参数错误。
	FAILEDOPERATION_INVALIDPARAMSGROUPID = "FailedOperation.InvalidParamsGroupId"

	// 队组名称参数错误。
	FAILEDOPERATION_INVALIDPARAMSGROUPNAME = "FailedOperation.InvalidParamsGroupName"

	// 队组owner参数错误。
	FAILEDOPERATION_INVALIDPARAMSGROUPOWNER = "FailedOperation.InvalidParamsGroupOwner"

	// 队组玩家自定义属性参数错误。
	FAILEDOPERATION_INVALIDPARAMSGROUPPLAYERCUSTOMPROPERTIES = "FailedOperation.InvalidParamsGroupPlayerCustomProperties"

	// 队组玩家自定义状态参数错误。
	FAILEDOPERATION_INVALIDPARAMSGROUPPLAYERCUSTOMSTATUS = "FailedOperation.InvalidParamsGroupPlayerCustomStatus"

	// 队组玩家名称参数错误。
	FAILEDOPERATION_INVALIDPARAMSGROUPPLAYERNAME = "FailedOperation.InvalidParamsGroupPlayerName"

	// 队组接收消息类型参数错误。
	FAILEDOPERATION_INVALIDPARAMSGROUPRECVTYPE = "FailedOperation.InvalidParamsGroupRecvType"

	// 队组类型参数错误。
	FAILEDOPERATION_INVALIDPARAMSGROUPTYPE = "FailedOperation.InvalidParamsGroupType"

	// 参数错误match_code。
	FAILEDOPERATION_INVALIDPARAMSMATCHCODE = "FailedOperation.InvalidParamsMatchCode"

	// 参数错误match_type。
	FAILEDOPERATION_INVALIDPARAMSMATCHTYPE = "FailedOperation.InvalidParamsMatchType"

	// 最大玩家数量参数错误。
	FAILEDOPERATION_INVALIDPARAMSMAXPLAYER = "FailedOperation.InvalidParamsMaxPlayer"

	// 参数错误max_players。
	FAILEDOPERATION_INVALIDPARAMSMAXPLAYERS = "FailedOperation.InvalidParamsMaxPlayers"

	// 参数错误message。
	FAILEDOPERATION_INVALIDPARAMSMESSAGE = "FailedOperation.InvalidParamsMessage"

	// 消息长度超过限制。
	FAILEDOPERATION_INVALIDPARAMSMESSAGELENGTH = "FailedOperation.InvalidParamsMessageLength"

	// 消息队列消息decode参数错误。
	FAILEDOPERATION_INVALIDPARAMSMSGQDECODE = "FailedOperation.InvalidParamsMsgqDecode"

	// 消息队列消息encode参数错误。
	FAILEDOPERATION_INVALIDPARAMSMSGQENCODE = "FailedOperation.InvalidParamsMsgqEncode"

	// 参数错误network_state。
	FAILEDOPERATION_INVALIDPARAMSNETWORKSTATE = "FailedOperation.InvalidParamsNetworkState"

	// 参数错误nonce。
	FAILEDOPERATION_INVALIDPARAMSNONCE = "FailedOperation.InvalidParamsNonce"

	// 参数错误open_id。
	FAILEDOPERATION_INVALIDPARAMSOPENID = "FailedOperation.InvalidParamsOpenId"

	// 参数错误owner。
	FAILEDOPERATION_INVALIDPARAMSOWNER = "FailedOperation.InvalidParamsOwner"

	// 参数错误owner_open_id。
	FAILEDOPERATION_INVALIDPARAMSOWNEROPENID = "FailedOperation.InvalidParamsOwnerOpenId"

	// 参数错误page_no。
	FAILEDOPERATION_INVALIDPARAMSPAGENO = "FailedOperation.InvalidParamsPageNo"

	// 参数错误page_size。
	FAILEDOPERATION_INVALIDPARAMSPAGESIZE = "FailedOperation.InvalidParamsPageSize"

	// 参数错误platform。
	FAILEDOPERATION_INVALIDPARAMSPLATFORM = "FailedOperation.InvalidParamsPlatform"

	// 玩法协议规则表达式错误。
	FAILEDOPERATION_INVALIDPARAMSPLAYMODEEXPRESSION = "FailedOperation.InvalidParamsPlayModeExpression"

	// 玩法协议规则类型错误。
	FAILEDOPERATION_INVALIDPARAMSPLAYMODERULETYPE = "FailedOperation.InvalidParamsPlayModeRuletype"

	// 玩法协议规则团队表达式错误。
	FAILEDOPERATION_INVALIDPARAMSPLAYMODETEAM = "FailedOperation.InvalidParamsPlayModeTeam"

	// 玩法协议版本号错误。
	FAILEDOPERATION_INVALIDPARAMSPLAYMODEVERSION = "FailedOperation.InvalidParamsPlayModeVersion"

	// 参数错误player_id。
	FAILEDOPERATION_INVALIDPARAMSPLAYERID = "FailedOperation.InvalidParamsPlayerId"

	// 参数错误player_info。
	FAILEDOPERATION_INVALIDPARAMSPLAYERINFO = "FailedOperation.InvalidParamsPlayerInfo"

	// 参数错误playerlist。
	FAILEDOPERATION_INVALIDPARAMSPLAYERLIST = "FailedOperation.InvalidParamsPlayerList"

	// 玩家不在队组中不允许操作。
	FAILEDOPERATION_INVALIDPARAMSPLAYERNOTINGROUP = "FailedOperation.InvalidParamsPlayerNotInGroup"

	// 队组接收消息的玩家中存在不在队组中的玩家。
	FAILEDOPERATION_INVALIDPARAMSRECVPLAYERID = "FailedOperation.InvalidParamsRecvPlayerId"

	// 参数错误region。
	FAILEDOPERATION_INVALIDPARAMSREGION = "FailedOperation.InvalidParamsRegion"

	// 参数错误create_type。
	FAILEDOPERATION_INVALIDPARAMSROOMCREATETYPE = "FailedOperation.InvalidParamsRoomCreateType"

	// 参数错误room_name。
	FAILEDOPERATION_INVALIDPARAMSROOMNAME = "FailedOperation.InvalidParamsRoomName"

	// 参数错误room_type。
	FAILEDOPERATION_INVALIDPARAMSROOMTYPE = "FailedOperation.InvalidParamsRoomType"

	// 参数错误sign。
	FAILEDOPERATION_INVALIDPARAMSSIGN = "FailedOperation.InvalidParamsSign"

	// 参数错误timestamp。
	FAILEDOPERATION_INVALIDPARAMSTIMESTAMP = "FailedOperation.InvalidParamsTimestamp"

	// 参数错误token。
	FAILEDOPERATION_INVALIDPARAMSTOKEN = "FailedOperation.InvalidParamsToken"

	// [rm]当前大区找不到合适的匹配,内部接口用。
	FAILEDOPERATION_MATCHCANNOTFOUND = "FailedOperation.MatchCanNotFound"

	// 取消匹配失败。
	FAILEDOPERATION_MATCHCANCELFAILED = "FailedOperation.MatchCancelFailed"

	// 匹配创建房间失败。
	FAILEDOPERATION_MATCHCREATEROOMERR = "FailedOperation.MatchCreateRoomErr"

	// 匹配创房有玩家已经在房间中。
	FAILEDOPERATION_MATCHCREATEROOMPLAYERALREADYINROOM = "FailedOperation.MatchCreateRoomPlayerAlreadyInRoom"

	// 匹配失败。
	FAILEDOPERATION_MATCHERR = "FailedOperation.MatchErr"

	// 游戏信息不存在。
	FAILEDOPERATION_MATCHGAMEINFONOTEXIST = "FailedOperation.MatchGameInfoNotExist"

	// 获取匹配信息失败。
	FAILEDOPERATION_MATCHGETMATCHINFOERR = "FailedOperation.MatchGetMatchInfoErr"

	// 匹配获取玩家属性失败。
	FAILEDOPERATION_MATCHGETPLAYERATTRFAIL = "FailedOperation.MatchGetPlayerAttrFail"

	// 查询匹配队列信息失败。
	FAILEDOPERATION_MATCHGETPLAYERLISTINFOERR = "FailedOperation.MatchGetPlayerListInfoErr"

	// 匹配获取队伍属性失败。
	FAILEDOPERATION_MATCHGETTEAMATTRFAIL = "FailedOperation.MatchGetTeamAttrFail"

	// 匹配小组人数超过队伍上限。
	FAILEDOPERATION_MATCHGROUPNUMEXCEEDLIMIT = "FailedOperation.MatchGroupNumExceedLimit"

	// 匹配无效参数。
	FAILEDOPERATION_MATCHINVALIDPARAMS = "FailedOperation.MatchInvalidParams"

	// 匹配加入房间失败。
	FAILEDOPERATION_MATCHJOINROOMERR = "FailedOperation.MatchJoinRoomErr"

	// 匹配逻辑错误。
	FAILEDOPERATION_MATCHLOGICERR = "FailedOperation.MatchLogicErr"

	// 匹配失败，无任何房间。
	FAILEDOPERATION_MATCHNOROOM = "FailedOperation.MatchNoRoom"

	// 玩家属性无法决定队伍类别。
	FAILEDOPERATION_MATCHNONETEAMTYPEFIT = "FailedOperation.MatchNoneTeamTypeFit"

	// 匹配参数不完整。
	FAILEDOPERATION_MATCHPLAYATTRNOTFOUND = "FailedOperation.MatchPlayAttrNotFound"

	// 匹配规则获取属性匹配区间失败。
	FAILEDOPERATION_MATCHPLAYRULEATTRSEGMENTNOTFOUND = "FailedOperation.MatchPlayRuleAttrSegmentNotFound"

	// 匹配规则算法错误。
	FAILEDOPERATION_MATCHPLAYRULEFUNCERR = "FailedOperation.MatchPlayRuleFuncErr"

	// 匹配规则不存在。
	FAILEDOPERATION_MATCHPLAYRULENOTFOUND = "FailedOperation.MatchPlayRuleNotFound"

	// 匹配规则不可用。
	FAILEDOPERATION_MATCHPLAYRULENOTRUNNING = "FailedOperation.MatchPlayRuleNotRunning"

	// 玩家属性不存在。
	FAILEDOPERATION_MATCHPLAYERATTRNOTFOUND = "FailedOperation.MatchPlayerAttrNotFound"

	// 匹配小组中玩家ID重复。
	FAILEDOPERATION_MATCHPLAYERIDISREPEATED = "FailedOperation.MatchPlayerIdIsRepeated"

	// 用户信息不存在。
	FAILEDOPERATION_MATCHPLAYERINFONOTEXIST = "FailedOperation.MatchPlayerInfoNotExist"

	// 用户已经在匹配中。
	FAILEDOPERATION_MATCHPLAYERISINMATCH = "FailedOperation.MatchPlayerIsInMatch"

	// 用户不在匹配状态。
	FAILEDOPERATION_MATCHPLAYERNOTINMATCH = "FailedOperation.MatchPlayerNotInMatch"

	// 查询游戏信息失败。
	FAILEDOPERATION_MATCHQUERYGAMEERR = "FailedOperation.MatchQueryGameErr"

	// 查询用户信息失败。
	FAILEDOPERATION_MATCHQUERYPLAYERERR = "FailedOperation.MatchQueryPlayerErr"

	// 查询大区信息失败。
	FAILEDOPERATION_MATCHQUERYREGIONERR = "FailedOperation.MatchQueryRegionErr"

	// 无大区信息。
	FAILEDOPERATION_MATCHREGIONINFONOTEXIST = "FailedOperation.MatchRegionInfoNotExist"

	// 匹配已经取消。
	FAILEDOPERATION_MATCHREQUESTCANCELED = "FailedOperation.MatchRequestCanceled"

	// 匹配请求ID已经存在。
	FAILEDOPERATION_MATCHREQUESTIDISEXIST = "FailedOperation.MatchRequestIdIsExist"

	// 匹配请求ID不存在。
	FAILEDOPERATION_MATCHREQUESTIDNOTEXIST = "FailedOperation.MatchRequestIdNotExist"

	// 匹配机器人Group不正确。
	FAILEDOPERATION_MATCHROBOTGROUPNOTRIGHT = "FailedOperation.MatchRobotGroupNotRight"

	// 匹配机器人Team不正确。
	FAILEDOPERATION_MATCHROBOTTEAMNOTRIGHT = "FailedOperation.MatchRobotTeamNotRight"

	// 团队匹配失败。
	FAILEDOPERATION_MATCHTEAMFAIL = "FailedOperation.MatchTeamFail"

	// 队伍匹配失败。
	FAILEDOPERATION_MATCHTEAMMATCHFAIL = "FailedOperation.MatchTeamMatchFail"

	// 玩家伍类别非法。
	FAILEDOPERATION_MATCHTEAMTYPEINVALID = "FailedOperation.MatchTeamTypeInvalid"

	// 匹配超时。
	FAILEDOPERATION_MATCHTIMEOUT = "FailedOperation.MatchTimeout"

	// 更新匹配信息失败。
	FAILEDOPERATION_MATCHUPDATEMATCHINFOERR = "FailedOperation.MatchUpdateMatchInfoErr"

	// 没有队组操作权限。
	FAILEDOPERATION_NOGROUPOPERATIONPERMISSION = "FailedOperation.NoGroupOperationPermission"

	// 没有权限请求。
	FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"

	// 队组禁止玩家加入。
	FAILEDOPERATION_OPERATIONFAILEDGROUPFORBIDJOIN = "FailedOperation.OperationFailedGroupForbidJoin"

	// 参数错误。
	FAILEDOPERATION_PARAMSINVALID = "FailedOperation.ParamsInvalid"

	// 持久化队组数量超过限制。
	FAILEDOPERATION_PERSISTENCEGROUPNUMEXCEEDTHELIMIT = "FailedOperation.PersistenceGroupNumExceedTheLimit"

	// 新增用户信息失败。
	FAILEDOPERATION_PLAYERADDPLAYERFAIL = "FailedOperation.PlayerAddPlayerFail"

	// 清除token缓存失败。
	FAILEDOPERATION_PLAYERCLEARTOKENFAIL = "FailedOperation.PlayerClearTokenFail"

	// 重复请求。
	FAILEDOPERATION_PLAYERDUPLICATEREQ = "FailedOperation.PlayerDuplicateReq"

	// game不存在。
	FAILEDOPERATION_PLAYERGAMENOTEXIST = "FailedOperation.PlayerGameNotExist"

	// 游戏已停止服务。
	FAILEDOPERATION_PLAYERGAMEOUTOFSERVICE = "FailedOperation.PlayerGameOutOfService"

	// 查询token失败。
	FAILEDOPERATION_PLAYERGETTOKENFAIL = "FailedOperation.PlayerGetTokenFail"

	// 玩家加入的对组个数超过限制。
	FAILEDOPERATION_PLAYERGROUPNUMLIMITEXCEED = "FailedOperation.PlayerGroupNumLimitExceed"

	// 玩家已经在队组中。
	FAILEDOPERATION_PLAYERISEXISTGROUP = "FailedOperation.PlayerIsExistGroup"

	// 玩家不在该队组中。
	FAILEDOPERATION_PLAYERISNOTEXISTGROUP = "FailedOperation.PlayerIsNotExistGroup"

	// 获取分布式锁失败。
	FAILEDOPERATION_PLAYERLOCKFAIL = "FailedOperation.PlayerLockFail"

	// 查询game信息失败。
	FAILEDOPERATION_PLAYERQUERYGAMEFAIL = "FailedOperation.PlayerQueryGameFail"

	// 查询用户信息失败。
	FAILEDOPERATION_PLAYERQUERYPLAYERFAIL = "FailedOperation.PlayerQueryPlayerFail"

	// 用户记录数不正确。
	FAILEDOPERATION_PLAYERRECORDNUMERR = "FailedOperation.PlayerRecordNumErr"

	// 保存token缓存失败。
	FAILEDOPERATION_PLAYERSAVETOKENFAIL = "FailedOperation.PlayerSaveTokenFail"

	// 查询secret_key失败。
	FAILEDOPERATION_PLAYERSECRETKEYFAIL = "FailedOperation.PlayerSecretKeyFail"

	// sign校验失败。
	FAILEDOPERATION_PLAYERSIGNERR = "FailedOperation.PlayerSignErr"

	// timestamp非法。
	FAILEDOPERATION_PLAYERTIMESTAMPINVALID = "FailedOperation.PlayerTimestampInvalid"

	// token非法。
	FAILEDOPERATION_PLAYERTOKENINVALID = "FailedOperation.PlayerTokenInvalid"

	// token不存在。
	FAILEDOPERATION_PLAYERTOKENNOTEXIST = "FailedOperation.PlayerTokenNotExist"

	// 释放分布式锁失败。
	FAILEDOPERATION_PLAYERUNLOCKFAIL = "FailedOperation.PlayerUnlockFail"

	// 重复创建。
	FAILEDOPERATION_RELAYALREADYEXISTS = "FailedOperation.RelayAlreadyExists"

	// 清理房间对局数据失败。
	FAILEDOPERATION_RELAYCLEANRELAYROOMFAIL = "FailedOperation.RelayCleanRelayRoomFail"

	// data长度超限制。
	FAILEDOPERATION_RELAYDATAEXCEEDLIMITED = "FailedOperation.RelayDataExceedLimited"

	// 转发到client-sdk失败。
	FAILEDOPERATION_RELAYFORWARDTOCLIENTFAIL = "FailedOperation.RelayForwardToClientFail"

	// 转发到自定义逻辑svr失败。
	FAILEDOPERATION_RELAYFORWARDTOGAMESVRFAIL = "FailedOperation.RelayForwardToGamesvrFail"

	// gamesvr查不到房间信息报错。
	FAILEDOPERATION_RELAYGAMESVRNOTFOUNDROOMFAIL = "FailedOperation.RelayGamesvrNotFoundRoomFail"

	// 自定义扩展服务（gamesvr）未开通。
	FAILEDOPERATION_RELAYGAMESVRSERVICENOTOPEN = "FailedOperation.RelayGamesvrServiceNotOpen"

	// 查询帧缓存失败。
	FAILEDOPERATION_RELAYGETFRAMECACHEFAIL = "FailedOperation.RelayGetFrameCacheFail"

	// 共享内存缓存错误。
	FAILEDOPERATION_RELAYHKVCACHEERROR = "FailedOperation.RelayHkvCacheError"

	// 帧率非法。
	FAILEDOPERATION_RELAYINVALIDFRAMERATE = "FailedOperation.RelayInvalidFrameRate"

	// 成员已存在。
	FAILEDOPERATION_RELAYMEMBERALREADYEXISTS = "FailedOperation.RelayMemberAlreadyExists"

	// 成员不存在。
	FAILEDOPERATION_RELAYMEMBERNOTEXISTS = "FailedOperation.RelayMemberNotExists"

	// 无可用的pod。
	FAILEDOPERATION_RELAYNOAVAILABLEPOD = "FailedOperation.RelayNoAvailablePod"

	// 没任何成员。
	FAILEDOPERATION_RELAYNOMEMBERS = "FailedOperation.RelayNoMembers"

	// 没权限，401开头是权限相关错误。
	FAILEDOPERATION_RELAYNOPERMISSION = "FailedOperation.RelayNoPermission"

	// 服务不存在。
	FAILEDOPERATION_RELAYNOTEXISTS = "FailedOperation.RelayNotExists"

	// 通知自定义服务gamesvr失败，402开头，是自定义gamesvr相关的错误。
	FAILEDOPERATION_RELAYNOTIFYGAMESVRFAIL = "FailedOperation.RelayNotifyGamesvrFail"

	// 通知relayworker失败。
	FAILEDOPERATION_RELAYNOTIFYRELAYWORKERFAIL = "FailedOperation.RelayNotifyRelayworkerFail"

	// redis缓存错误。
	FAILEDOPERATION_RELAYREDISCACHEERROR = "FailedOperation.RelayRedisCacheError"

	// 补帧的时候游戏没有开始。
	FAILEDOPERATION_RELAYREQFRAMEGAMENOTSTARTED = "FailedOperation.RelayReqFrameGameNotStarted"

	// 请求分配pod失败。
	FAILEDOPERATION_RELAYREQPODFAIL = "FailedOperation.RelayReqPodFail"

	// 重置房间对局失败。
	FAILEDOPERATION_RELAYRESETRELAYROOMFAIL = "FailedOperation.RelayResetRelayRoomFail"

	// 开局状态下，G不允许修改帧率。
	FAILEDOPERATION_RELAYSETFRAMERATEFORBIDDEN = "FailedOperation.RelaySetFrameRateForbidden"

	// 状态异常。
	FAILEDOPERATION_RELAYSTATEINVALID = "FailedOperation.RelayStateInvalid"

	// 被移除的玩家Id为空。
	FAILEDOPERATION_REMOVEPLAYERIDISEMPTY = "FailedOperation.RemovePlayerIdIsEmpty"

	// 请求包格式错误。
	FAILEDOPERATION_REQBADPKG = "FailedOperation.ReqBadPkg"

	// ctrlsvr分配relaysvr失败。
	FAILEDOPERATION_ROOMALLOCATERELAYSVRIPPORTERR = "FailedOperation.RoomAllocateRelaysvrIpPortErr"

	// 检查登录失败。
	FAILEDOPERATION_ROOMCHECKLOGINSESSIONERR = "FailedOperation.RoomCheckLoginSessionErr"

	// 创建房间失败。
	FAILEDOPERATION_ROOMCREATEFAIL = "FailedOperation.RoomCreateFail"

	// 创建房间无权限。
	FAILEDOPERATION_ROOMCREATENOPERMISSION = "FailedOperation.RoomCreateNoPermission"

	// 销毁房间无权限。
	FAILEDOPERATION_ROOMDESTORYNOPERMISSION = "FailedOperation.RoomDestoryNoPermission"

	// 无解散房间权限。
	FAILEDOPERATION_ROOMDISSMISSNOPERMISSION = "FailedOperation.RoomDissmissNoPermission"

	// 游戏信息不存在。
	FAILEDOPERATION_ROOMGAMEINFONOTEXIST = "FailedOperation.RoomGameInfoNotExist"

	// 查询用户信息失败。
	FAILEDOPERATION_ROOMGETPLAYERINFOERR = "FailedOperation.RoomGetPlayerInfoErr"

	// 获取房间信息失败。
	FAILEDOPERATION_ROOMGETROOMINFOERR = "FailedOperation.RoomGetRoomInfoErr"

	// 房间信息不存在。
	FAILEDOPERATION_ROOMINFOUNEXIST = "FailedOperation.RoomInfoUnexist"

	// 房间teamId无效。
	FAILEDOPERATION_ROOMINVALIDPARAMSTEAMID = "FailedOperation.RoomInvalidParamsTeamId"

	// 无权限加入房间。
	FAILEDOPERATION_ROOMJOINNOPERMISSION = "FailedOperation.RoomJoinNoPermission"

	// 房间不允许加入用户。
	FAILEDOPERATION_ROOMJOINNOTALLOW = "FailedOperation.RoomJoinNotAllow"

	// 最大用户数值设置非法。
	FAILEDOPERATION_ROOMMAXPLAYERSINVALID = "FailedOperation.RoomMaxPlayersInvalid"

	// 房间数量超过限制。
	FAILEDOPERATION_ROOMMAXROOMNUMBEREXCEEDLIMIT = "FailedOperation.RoomMaxRoomNumberExceedLimit"

	// 修改房主失败。
	FAILEDOPERATION_ROOMMODIFYOWNERERR = "FailedOperation.RoomModifyOwnerErr"

	// 玩家信息操作繁忙，请重试。
	FAILEDOPERATION_ROOMMODIFYPLAYERBUSY = "FailedOperation.RoomModifyPlayerBusy"

	// 无修改房间属性权限。
	FAILEDOPERATION_ROOMMODIFYPROPERTIESNOPEMISSION = "FailedOperation.RoomModifyPropertiesNoPemission"

	// 页号、页数大小参数不合法，可能实际大小没这么大。
	FAILEDOPERATION_ROOMPARAMPAGEINVALID = "FailedOperation.RoomParamPageInvalid"

	// 用户已经在房间内，不能操作创建房间、加房等操作。
	FAILEDOPERATION_ROOMPLAYERALREADYINROOM = "FailedOperation.RoomPlayerAlreadyInRoom"

	// 用户信息不存在。
	FAILEDOPERATION_ROOMPLAYERINFONOTEXIST = "FailedOperation.RoomPlayerInfoNotExist"

	// 用户目前不在房间内，不能操作更改房间属性、踢人等操作。
	FAILEDOPERATION_ROOMPLAYERNOTINROOM = "FailedOperation.RoomPlayerNotInRoom"

	// 用户在房间中掉线，不能开始游戏等操作。
	FAILEDOPERATION_ROOMPLAYEROFFLINE = "FailedOperation.RoomPlayerOffline"

	// 房间内用户数已经达到最大人数不能再加入了。
	FAILEDOPERATION_ROOMPLAYERSEXCEEDLIMIT = "FailedOperation.RoomPlayersExceedLimit"

	// 游戏信息失败。
	FAILEDOPERATION_ROOMQUERYGAMEERR = "FailedOperation.RoomQueryGameErr"

	// 查询用户信息失败。
	FAILEDOPERATION_ROOMQUERYPLAYERERR = "FailedOperation.RoomQueryPlayerErr"

	// 查询地域信息失败。
	FAILEDOPERATION_ROOMQUERYREGIONERR = "FailedOperation.RoomQueryRegionErr"

	// 查询不到accessRegion信息。
	FAILEDOPERATION_ROOMREGIONINFONOTEXIST = "FailedOperation.RoomRegionInfoNotExist"

	// 无踢人权限。
	FAILEDOPERATION_ROOMREMOVEPLAYERNOPERMISSION = "FailedOperation.RoomRemovePlayerNoPermission"

	// 被踢玩家不在房间中。
	FAILEDOPERATION_ROOMREMOVEPLAYERNOTINROOM = "FailedOperation.RoomRemovePlayerNotInRoom"

	// 无踢出自己权限。
	FAILEDOPERATION_ROOMREMOVESELFNOPERMISSION = "FailedOperation.RoomRemoveSelfNoPermission"

	// 房间团队人员已满。
	FAILEDOPERATION_ROOMTEAMMEMBERLIMITEXCEED = "FailedOperation.RoomTeamMemberLimitExceed"

	// 编码失败。
	FAILEDOPERATION_SDKENCODEPARAMFAIL = "FailedOperation.SdkEncodeParamFail"

	// 参数错误。
	FAILEDOPERATION_SDKINVALIDPARAMS = "FailedOperation.SdkInvalidParams"

	// 帧同步鉴权错误。
	FAILEDOPERATION_SDKNOCHECKLOGIN = "FailedOperation.SdkNoCheckLogin"

	// 登录态错误。
	FAILEDOPERATION_SDKNOLOGIN = "FailedOperation.SdkNoLogin"

	// 无房间。
	FAILEDOPERATION_SDKNOROOM = "FailedOperation.SdkNoRoom"

	// 消息响应超时。
	FAILEDOPERATION_SDKRESTIMEOUT = "FailedOperation.SdkResTimeout"

	// 消息发送失败。
	FAILEDOPERATION_SDKSENDFAIL = "FailedOperation.SdkSendFail"

	// Socket断开。
	FAILEDOPERATION_SDKSOCKETCLOSE = "FailedOperation.SdkSocketClose"

	// 网络错误。
	FAILEDOPERATION_SDKSOCKETERROR = "FailedOperation.SdkSocketError"

	// SDK未初始化。
	FAILEDOPERATION_SDKUNINIT = "FailedOperation.SdkUninit"

	// 服务器繁忙。
	FAILEDOPERATION_SERVERBUSY = "FailedOperation.ServerBusy"

	// 标签添加失败。
	FAILEDOPERATION_TAGADDFAILED = "FailedOperation.TagAddFailed"

	// 标签接口调用失败，请稍后再试。若无法解决，请在线咨询。
	FAILEDOPERATION_TAGCALLERFAILED = "FailedOperation.TagCallerFailed"

	// 后端超时错误。
	FAILEDOPERATION_TIMEOUT = "FailedOperation.TimeOut"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 配置房间id管理模块错误。
	INTERNALERROR_CONFROOMIDBUCKETERR = "InternalError.ConfRoomIdBucketErr"

	// 数据格式转化失败。
	INTERNALERROR_DATAFORMATERR = "InternalError.DataFormatErr"

	// hashcode解码失败。
	INTERNALERROR_HASHIDDECODEERR = "InternalError.HashidDecodeErr"

	// hashcode编码失败。
	INTERNALERROR_HASHIDENCODEERR = "InternalError.HashidEncodeErr"

	// hashcode生成失败。
	INTERNALERROR_HASHIDERR = "InternalError.HashidErr"

	// 参数错误recordId。
	INTERNALERROR_INVALIDPARAMSRECOREID = "InternalError.InvalidParamsRecoreId"

	// JSON数据格式转化失败。
	INTERNALERROR_JSONFORMATERR = "InternalError.JsonFormatErr"

	// 玩法数据格式转化失败。
	INTERNALERROR_JSONPLAYMODEFORMATERR = "InternalError.JsonPlayModeFormatErr"

	// 玩法数据格式转化失败。
	INTERNALERROR_JSONPLAYMODEPARISEERR = "InternalError.JsonPlayModePariseErr"

	// 匹配内部逻辑错误。
	INTERNALERROR_MATCHINNERLOGICERR = "InternalError.MatchInnerLogicErr"

	// 匹配内部参数错误。
	INTERNALERROR_MATCHINNERPARAMSERR = "InternalError.MatchInnerParamsErr"

	// 匹配不是GSE类型查询匹配结果失败。
	INTERNALERROR_MATCHRESULTTYPENOTGSE = "InternalError.MatchResultTypeNotGse"

	// 匹配房间添加节点失败。
	INTERNALERROR_MATCHROOMINNERADDNODEERR = "InternalError.MatchRoomInnerAddNodeErr"

	// 匹配房间删除节点失败。
	INTERNALERROR_MATCHROOMINNERDELNODEERR = "InternalError.MatchRoomInnerDelNodeErr"

	// myspp框架返回-1000。
	INTERNALERROR_MYSPPSYSTEMERR = "InternalError.MysppSystemErr"

	// 删除失败。
	INTERNALERROR_MYSQLDELETEFAIL = "InternalError.MysqlDeleteFail"

	// 插入失败。
	INTERNALERROR_MYSQLINSERTFAIL = "InternalError.MysqlInsertFail"

	// 查询为空。
	INTERNALERROR_MYSQLMULTIROWFOUND = "InternalError.MysqlMultiRowFound"

	// 查询为空。
	INTERNALERROR_MYSQLNOROWFOUND = "InternalError.MysqlNoRowFound"

	// 查询失败。
	INTERNALERROR_MYSQLQUERYSFAIL = "InternalError.MysqlQuerysFail"

	// 更新失败。
	INTERNALERROR_MYSQLUPDATEFAIL = "InternalError.MysqlUpdateFail"

	// 反序列化失败。
	INTERNALERROR_PBPARSEFROMSTRERR = "InternalError.PbParseFromStrErr"

	// 序列化失败。
	INTERNALERROR_PBSERIALIZETOSTRERR = "InternalError.PbSerializeToStrErr"

	// redisdel类操作失败。
	INTERNALERROR_REDISDELOPERR = "InternalError.RedisDelOpErr"

	// redis操作异常。
	INTERNALERROR_REDISEXPIREOPERR = "InternalError.RedisExpireOpErr"

	// redisget类操作失败。
	INTERNALERROR_REDISGETOPERR = "InternalError.RedisGetOpErr"

	// redisKEY不存在。
	INTERNALERROR_REDISKEYNOTEXIST = "InternalError.RedisKeyNotExist"

	// redislist操作失败。
	INTERNALERROR_REDISLISTOPERR = "InternalError.RedisListOpErr"

	// redislistpop空结果。
	INTERNALERROR_REDISLISTPOPEMPTY = "InternalError.RedisListPopEmpty"

	// redis加锁冲突类操作失败。
	INTERNALERROR_REDISLOCKALREADYEXIST = "InternalError.RedisLockAlreadyExist"

	// redis加锁类操作失败。
	INTERNALERROR_REDISLOCKOPERR = "InternalError.RedisLockOpErr"

	// redis操作参数不合法。
	INTERNALERROR_REDISOPINVALIDPARAMS = "InternalError.RedisOpInvalidParams"

	// redis实例池获取实例失败。
	INTERNALERROR_REDISPOOLGETINSTANCEFAIL = "InternalError.RedisPoolGetInstanceFail"

	// redisset内为空。
	INTERNALERROR_REDISSETISEMPTY = "InternalError.RedisSetIsEmpty"

	// redisset类操作失败。
	INTERNALERROR_REDISSETOPERR = "InternalError.RedisSetOpErr"

	// 申请service失败。
	INTERNALERROR_ROOMALLOCATESERVICEFAIL = "InternalError.RoomAllocateServiceFail"

	// mysql数据库插入历史房间信息失败。
	INTERNALERROR_ROOMHISTORYINFOINSERTERR = "InternalError.RoomHistoryInfoInsertErr"

	// 检查锁失败，一般是过期。
	INTERNALERROR_ROOMREDISCHECKLOCKERR = "InternalError.RoomRedisCheckLockErr"

	// 删除锁失败。
	INTERNALERROR_ROOMREDISDELLOCKERR = "InternalError.RoomRedisDelLockErr"

	// 获取锁失败。
	INTERNALERROR_ROOMREDISGETLOCKERR = "InternalError.RoomRedisGetLockErr"

	// 数据库更新失败。
	INTERNALERROR_ROOMREDISUPDATEERR = "InternalError.RoomRedisUpdateErr"

	// 删除用户房间映射表信息失败。
	INTERNALERROR_ROOMREMOVEREDISPLAYERROOMMATCHERR = "InternalError.RoomRemoveRedisPlayerRoomMatchErr"

	// 删除房间信息表信息失败。
	INTERNALERROR_ROOMREMOVEREDISROOMINFOERR = "InternalError.RoomRemoveRedisRoomInfoErr"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 玩家ID在玩家列表中重复。
	INVALIDPARAMETER_DUPLICATEPLAYERIDINPLAYERS = "InvalidParameter.DuplicatePlayerIdInPlayers"

	// 无效的游戏描述长度。
	INVALIDPARAMETER_GAMEDESCLENGTH = "InvalidParameter.GameDescLength"

	// 无效的游戏名称长度。
	INVALIDPARAMETER_GAMENAMELENGTH = "InvalidParameter.GameNameLength"

	// 无效的游戏平台。
	INVALIDPARAMETER_GAMEPLATFORM = "InvalidParameter.GamePlatform"

	// 无效的游戏类型。
	INVALIDPARAMETER_GAMETYPE = "InvalidParameter.GameType"

	// 无效的自定义房间属性。
	INVALIDPARAMETER_INVALIDCUSTOMPROPERTIES = "InvalidParameter.InvalidCustomProperties"

	// 无效的最大玩家数量。
	INVALIDPARAMETER_INVALIDMAXPLAYERS = "InvalidParameter.InvalidMaxPlayers"

	// 无效的最小玩家数量。
	INVALIDPARAMETER_INVALIDMINPLAYERS = "InvalidParameter.InvalidMinPlayers"

	// 无效的OpenId长度。
	INVALIDPARAMETER_INVALIDOPENIDLENGTH = "InvalidParameter.InvalidOpenIdLength"

	// 无效的自定义玩家属性长度。
	INVALIDPARAMETER_INVALIDPLAYERCUSTOMPROFILELENGTH = "InvalidParameter.InvalidPlayerCustomProfileLength"

	// 无效的自定义玩家状态。
	INVALIDPARAMETER_INVALIDPLAYERCUSTOMPROFILESTATUS = "InvalidParameter.InvalidPlayerCustomProfileStatus"

	// 无效的玩家昵称长度。
	INVALIDPARAMETER_INVALIDPLAYERNAMELENGTH = "InvalidParameter.InvalidPlayerNameLength"

	// 无效的玩家数量。
	INVALIDPARAMETER_INVALIDPLAYERSSIZE = "InvalidParameter.InvalidPlayersSize"

	// 错误的机器人匹配模式参数。
	INVALIDPARAMETER_INVALIDROBOTMATCHMODELPARAM = "InvalidParameter.InvalidRobotMatchModelParam"

	// 无效的房间名称。
	INVALIDPARAMETER_INVALIDROOMNAME = "InvalidParameter.InvalidRoomName"

	// 无效的房间类型长度。
	INVALIDPARAMETER_INVALIDROOMTYPELENGTH = "InvalidParameter.InvalidRoomTypeLength"

	// 无效的队伍Id长度。
	INVALIDPARAMETER_INVALIDTEAMIDLENGTH = "InvalidParameter.InvalidTeamIdLength"

	// 无效的队伍昵称长度。
	INVALIDPARAMETER_INVALIDTEAMNAMELENGTH = "InvalidParameter.InvalidTeamNameLength"

	// 无效的队伍大小。
	INVALIDPARAMETER_INVALIDTEAMSSIZE = "InvalidParameter.InvalidTeamsSize"

	// 无效的开通联网对战服务选项。
	INVALIDPARAMETER_OPENONLINESERVICE = "InvalidParameter.OpenOnlineService"

	// 房主信息不在玩家列表中。
	INVALIDPARAMETER_OWNERNOTINPLAYERS = "InvalidParameter.OwnerNotInPlayers"

	// 玩家数量不在队伍可容纳范围。
	INVALIDPARAMETER_PLAYERNUMNOTINTEAMRANGE = "InvalidParameter.PlayerNumNotInTeamRange"

	// 玩家ID在玩家列表中重复。
	INVALIDPARAMETER_PLAYEROPENIDINPLAYERSDUPLICATE = "InvalidParameter.PlayerOpenIdInPlayersDuplicate"

	// 创建满员房间但是玩家数量不足。
	INVALIDPARAMETER_PLAYERSIZENOTENOUGH = "InvalidParameter.PlayerSizeNotEnough"

	// 玩家队伍ID不在队伍列表中。
	INVALIDPARAMETER_PLAYERTEAMIDNOTINTEAMS = "InvalidParameter.PlayerTeamIdNotInTeams"

	// 无效的标签列表。
	INVALIDPARAMETER_TAGS = "InvalidParameter.Tags"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 日志集数量超出限制。
	LIMITEXCEEDED_CLSLOGSETEXCEED = "LimitExceeded.CLSLogsetExceed"

	// 日志主题数量超出限制。
	LIMITEXCEEDED_CLSTOPICEXCEED = "LimitExceeded.CLSTopicExceed"

	// 游戏数超过限额。
	LIMITEXCEEDED_GAMERESOURCELIMIT = "LimitExceeded.GameResourceLimit"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 请求的次数超过了频率限制。
	REQUESTLIMITEXCEEDED = "RequestLimitExceeded"

	// 操作过于频繁，请稍等几秒后重试。
	REQUESTLIMITEXCEEDED_FREQUENCYLIMIT = "RequestLimitExceeded.FrequencyLimit"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 房间不存在。
	RESOURCENOTFOUND_ROOMNOTEXIST = "ResourceNotFound.RoomNotExist"

	// 标签资源未找到。
	RESOURCENOTFOUND_TAGNOTFOUND = "ResourceNotFound.TagNotFound"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 日志服务(CLS)不可用，请确保您已在日志服务控制台开通服务。若无法解决，请在线咨询。
	RESOURCEUNAVAILABLE_CLSNOTALLOWED = "ResourceUnavailable.CLSNotAllowed"

	// 游戏已被冻结，操作失败。
	RESOURCEUNAVAILABLE_GAMEFROZEN = "ResourceUnavailable.GameFrozen"

	// 资源售罄。
	RESOURCESSOLDOUT = "ResourcesSoldOut"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 需要授权CAM权限操作。
	UNAUTHORIZEDOPERATION_CAMUNAUTHORIZEDOPERATION = "UnauthorizedOperation.CAMUnauthorizedOperation"

	// 需要授权GSE的CAM权限操作。
	UNAUTHORIZEDOPERATION_GSECAMUNAUTHORIZEDOPERATION = "UnauthorizedOperation.GseCAMUnauthorizedOperation"

	// 无操作权限。
	UNAUTHORIZEDOPERATION_NOTACTIONPERMISSION = "UnauthorizedOperation.NotActionPermission"

	// 无项目权限。
	UNAUTHORIZEDOPERATION_NOTPROJECTPERMISSION = "UnauthorizedOperation.NotProjectPermission"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 标签键不允许重复。
	UNSUPPORTEDOPERATION_TAGKEYDUPLICATE = "UnsupportedOperation.TagKeyDuplicate"
)
