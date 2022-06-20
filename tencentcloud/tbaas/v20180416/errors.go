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

package v20180416

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 用户无权限访问。
	AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"

	// 为带来更好的服务与体验，长安链体验网络维护升级中，请稍后再试。
	FAILEDOPERATION_BAASSTOPSERVING = "FailedOperation.BaaSStopServing"

	// Bcos网络异常。
	FAILEDOPERATION_BCOSSERVICE = "FailedOperation.BcosService"

	// 检查CSR文件非法。
	FAILEDOPERATION_CACHECKCSR = "FailedOperation.CaCheckCsr"

	// 数据库操作错误。
	FAILEDOPERATION_CADBOPTION = "FailedOperation.CaDbOption"

	// 证书已经存在。
	FAILEDOPERATION_CAEXSIT = "FailedOperation.CaExsit"

	// 产生密钥对失败。
	FAILEDOPERATION_CAGENKEY = "FailedOperation.CaGenkey"

	// 初始化CA失败。
	FAILEDOPERATION_CAINIT = "FailedOperation.CaInit"

	// 输入参数非法。
	FAILEDOPERATION_CAINPUTPARAM = "FailedOperation.CaInputParam"

	// 证书不存在。
	FAILEDOPERATION_CANOEXIST = "FailedOperation.CaNoExist"

	// 注销证书失败。
	FAILEDOPERATION_CAREVOKE = "FailedOperation.CaRevoke"

	// 根证书不存在。
	FAILEDOPERATION_CAROOTNONEXIST = "FailedOperation.CaRootNonExist"

	// CA内部错误。
	FAILEDOPERATION_CASERVICE = "FailedOperation.CaService"

	// 签发证书失败。
	FAILEDOPERATION_CASIGNCERT = "FailedOperation.CaSignCert"

	// CFCA云API申请证书失败。
	FAILEDOPERATION_CAYUNAPIAPPLYCERT = "FailedOperation.CaYunApiApplyCert"

	// CFCA云API错误。
	FAILEDOPERATION_CAYUNAPICOMMON = "FailedOperation.CaYunApiCommon"

	// 合约已在通道实例化。
	FAILEDOPERATION_CHAINCODECHANNEL = "FailedOperation.ChainCodeChannel"

	// 合约已存在。
	FAILEDOPERATION_CHAINCODEEXIST = "FailedOperation.ChainCodeExist"

	// 合约实例化错误。
	FAILEDOPERATION_CHAINCODEINIT = "FailedOperation.ChainCodeInit"

	// 合约安装错误。
	FAILEDOPERATION_CHAINCODEINSTALL = "FailedOperation.ChainCodeInstall"

	// Bcos不能编译已经部署的合约。
	FAILEDOPERATION_COMPILEDEPLOYEDCONTRACT = "FailedOperation.CompileDeployedContract"

	// Bcos不能编译正在部署的合约。
	FAILEDOPERATION_COMPILEDEPLOYINGCONTRACT = "FailedOperation.CompileDeployingContract"

	// Bcos只能由该合约所属机构执行操作。
	FAILEDOPERATION_CONTRACTEDITEDBYOTHERAGENCY = "FailedOperation.ContractEditedByOtherAgency"

	// 数据库操作异常。
	FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DatabaseException"

	// Bcos合约已部署，无法删除。
	FAILEDOPERATION_DELETEDEPLOYEDCONTRACT = "FailedOperation.DeleteDeployedContract"

	// Bcos合约正在部署，无法删除。
	FAILEDOPERATION_DELETEDEPLOYINGCONTRACT = "FailedOperation.DeleteDeployingContract"

	// Bcos不能部署未编译的合约。
	FAILEDOPERATION_DEPLOYCONTRACTNOTCOMPILE = "FailedOperation.DeployContractNotCompile"

	// 区块详情查询失败。
	FAILEDOPERATION_FABRICBLOCKDETAIL = "FailedOperation.FabricBlockDetail"

	// 区块列表查询失败。
	FAILEDOPERATION_FABRICBLOCKQUERY = "FailedOperation.FabricBlockQuery"

	// 提交失败。
	FAILEDOPERATION_FABRICCOMMIT = "FailedOperation.FabricCommit"

	// event hub失败。
	FAILEDOPERATION_FABRICEVENTHUB = "FailedOperation.FabricEventHub"

	// 区块链网络请求异常。
	FAILEDOPERATION_FABRICMANAGE = "FailedOperation.FabricManage"

	// 背书失败。
	FAILEDOPERATION_FABRICPROPOSAL = "FailedOperation.FabricProposal"

	// fabric请求参数错误。
	FAILEDOPERATION_FABRICREQUESTPARAMS = "FailedOperation.FabricRequestParams"

	// 参数验证失败。
	FAILEDOPERATION_FABRICREQUSTPARAMS = "FailedOperation.FabricRequstParams"

	// 交易详情查询失败。
	FAILEDOPERATION_FABRICTRANSACTIONDETAIL = "FailedOperation.FabricTransactionDetail"

	// 交易列表查询失败。
	FAILEDOPERATION_FABRICTRANSACTIONQUERY = "FailedOperation.FabricTransactionQuery"

	// Txid 查询失败。
	FAILEDOPERATION_FABRICTXIDQUERY = "FailedOperation.FabricTxIdQuery"

	// fabric invoke交易错误。
	FAILEDOPERATION_FABRICTXINVOKE = "FailedOperation.FabricTxInvoke"

	// fabric query交易错误。
	FAILEDOPERATION_FABRICTXQUERY = "FailedOperation.FabricTxQuery"

	// fabric query交易错误, 无正确。
	FAILEDOPERATION_FABRICTXQUERYNONE = "FailedOperation.FabricTxQueryNone"

	// Bcos前置服务调用失败。
	FAILEDOPERATION_FRONTREQUESTFAIL = "FailedOperation.FrontRequestFail"

	// 用户非法操作。
	FAILEDOPERATION_GROUPILLEGAL = "FailedOperation.GroupIllegal"

	// 无效合约。
	FAILEDOPERATION_INVALIDCHAINCODE = "FailedOperation.InvalidChaincode"

	// 无效通道。
	FAILEDOPERATION_INVALIDCHANNEL = "FailedOperation.InvalidChannel"

	// 无效网络。
	FAILEDOPERATION_INVALIDCLUSTER = "FailedOperation.InvalidCluster"

	// 无效组织。
	FAILEDOPERATION_INVALIDGROUP = "FailedOperation.InvalidGroup"

	// Bcos无效的群组编号。
	FAILEDOPERATION_INVALIDGROUPPK = "FailedOperation.InvalidGroupPk"

	// Bcos无效的私钥用户信息。
	FAILEDOPERATION_INVALIDKEYUSER = "FailedOperation.InvalidKeyUser"

	// 无效操作。
	FAILEDOPERATION_INVALIDOPERATION = "FailedOperation.InvalidOperation"

	// 无效节点。
	FAILEDOPERATION_INVALIDPEER = "FailedOperation.InvalidPeer"

	// 无效资源。
	FAILEDOPERATION_INVALIDRESOURCE = "FailedOperation.InvalidResource"

	// 交易请求异常。
	FAILEDOPERATION_MANAGESERVICE = "FailedOperation.ManageService"

	// Bcos新建合约失败。
	FAILEDOPERATION_NEWCONTRACT = "FailedOperation.NewContract"

	// 合约没有在通道初始化。
	FAILEDOPERATION_NOCHAINCODECHANNEL = "FailedOperation.NoChainCodeChannel"

	// 组织没有加入合约。
	FAILEDOPERATION_NOCHAINCODEGROUP = "FailedOperation.NoChainCodeGroup"

	// 合约没有在节点安装。
	FAILEDOPERATION_NOCHAINCODEPEER = "FailedOperation.NoChainCodePeer"

	// 组织没有加入通道。
	FAILEDOPERATION_NOCHANNELGROUP = "FailedOperation.NoChannelGroup"

	// 节点没有加入通道。
	FAILEDOPERATION_NOCHANNELPEER = "FailedOperation.NoChannelPeer"

	// 对象不存在。
	FAILEDOPERATION_NOOBJECT = "FailedOperation.NoObject"

	// 没有可用节点。
	FAILEDOPERATION_NOPEER = "FailedOperation.NoPeer"

	// Bcos合约未部署。
	FAILEDOPERATION_NOTDEPLOYEDCONTRACT = "FailedOperation.NotDeployedContract"

	// 服务调用失败，请检查参数。
	FAILEDOPERATION_SERVICEFAILED = "FailedOperation.ServiceFailed"

	// 操作状态不匹配。
	FAILEDOPERATION_STATUSNOMATCH = "FailedOperation.StatusNoMatch"

	// 后台服务请求超时。
	FAILEDOPERATION_TIMEOUTURL = "FailedOperation.TimeOutUrl"

	// 交易执行超时，请稍后再试。
	FAILEDOPERATION_TRANSACTIONTIMEOUT = "FailedOperation.TransactionTimeout"

	// Bcos更新已部署合约。
	FAILEDOPERATION_UPDATEDEPLOYEDCONTRACT = "FailedOperation.UpdateDeployedContract"

	// 获取用户认证类型出错。
	FAILEDOPERATION_USERAUTHTYPE = "FailedOperation.UserAuthType"

	// 您因违反用户协议，目前无法使用长安链体验网络。
	FAILEDOPERATION_USERINBLACKLIST = "FailedOperation.UserInBlackList"

	// 用户未加入体验网络。
	FAILEDOPERATION_USERNOJOINDEMOCLUSTER = "FailedOperation.UserNoJoinDemoCluster"

	// Bcos数据库操作异常，请重试。
	INTERNALERROR_DBERROR = "InternalError.DBError"

	// 内部错误。
	INTERNALERROR_FAILPREFORM = "InternalError.FailPreform"

	// 服务器异常，请重试。
	INTERNALERROR_FAILURL = "InternalError.FailUrl"

	// Flask内部错误。
	INTERNALERROR_FLASKEXCEPTION = "InternalError.FlaskException"

	// Bcos无效的合约参数。
	INTERNALERROR_INVALIDCONTRACTPARAM = "InternalError.InvalidContractParam"

	// Bcos不支持的请求类型。
	INTERNALERROR_METHODTYPENOTSUPPORT = "InternalError.MethodTypeNotSupport"

	// 错误码未定义。
	INTERNALERROR_NODEFINEERROR = "InternalError.NoDefineError"

	// Bcos服务器异常，请重试。
	INTERNALERROR_SERVERERROR = "InternalError.ServerError"

	// 服务器异常。
	INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"

	// 服务异常，请重试。
	INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"

	// 交易服务异常，请重试。
	INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"

	// 交易服务内部错误。
	INTERNALERROR_TRANSACTIONSERVICE = "InternalError.TransactionService"

	// 交易服务未知错误，请重试。
	INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"

	// Bcos账号参数错误。
	INVALIDPARAMETER_ACCOUNTPARAMERROR = "InvalidParameter.AccountParamError"

	// Bcos无效的机构信息。
	INVALIDPARAMETER_AGENCYINVALID = "InvalidParameter.AgencyInvalid"

	// Bcos新增机构网络关联信息入参错误。
	INVALIDPARAMETER_AGENCYNETPARAMINVALID = "InvalidParameter.AgencyNetParamInvalid"

	// Bcos所属联盟编号不能为空。
	INVALIDPARAMETER_ALLIANCEIDOFNETEMPTY = "InvalidParameter.AllianceIdOfNetEmpty"

	// Bcos无效的合约编号。
	INVALIDPARAMETER_CONTRACTIDINVALID = "InvalidParameter.ContractIdInvalid"

	// Bcos数据已存在，请勿重复添加。
	INVALIDPARAMETER_DATAHADEXIST = "InvalidParameter.DataHadExist"

	// Bcos无效参数。
	INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"

	// Bcos无效的失效时间。
	INVALIDPARAMETER_EXPIRETIMEINVALID = "InvalidParameter.ExpireTimeInvalid"

	// Bcos前置的IP或端口无效。
	INVALIDPARAMETER_FRONTCONNFAIL = "InvalidParameter.FrontConnFail"

	// Bcos IP格式错误。
	INVALIDPARAMETER_FRONTIPINVALID = "InvalidParameter.FrontIpInvalid"

	// Bcos新增前置服务参数错误。
	INVALIDPARAMETER_FRONTPARAMERROR = "InvalidParameter.FrontParamError"

	// Bcos前置服务调用失败。
	INVALIDPARAMETER_FRONTREQUESTFAIL = "InvalidParameter.FrontRequestFail"

	// 输入参数存在违规内容。
	INVALIDPARAMETER_INPUTDATAVIOLATION = "InvalidParameter.InputDataViolation"

	// Bcos无效的合约参数。
	INVALIDPARAMETER_INVALIDCONTRACTARG = "InvalidParameter.InvalidContractArg"

	// Bcos无效的网络编号。
	INVALIDPARAMETER_NETIDINVALID = "InvalidParameter.NetIdInvalid"

	// Bcos网络参数错误。
	INVALIDPARAMETER_NETPARAMERROR = "InvalidParameter.NetParamError"

	// Bcos新增私钥用户参数错误。
	INVALIDPARAMETER_NEWKEYUSERPARAMERROR = "InvalidParameter.NewKeyUserParamError"

	// Bcos无可删信息,请确认后重试。
	INVALIDPARAMETER_NOINFOTODELETE = "InvalidParameter.NoInfoToDelete"

	// Bcos没有有效的前置服务信息。
	INVALIDPARAMETER_NOTFOUNDVALIDFRONT = "InvalidParameter.NotFoundValidFront"

	// Bcos不能重复部署合约。
	INVALIDPARAMETER_REDEPLOYEDCONTRACT = "InvalidParameter.ReDeployedContract"

	// Bcos合约正在部署中，请勿重复操作。
	INVALIDPARAMETER_REDEPLOYINGCONTRACT = "InvalidParameter.ReDeployingContract"

	// Bcos无效的角色。
	INVALIDPARAMETER_ROLEINVALID = "InvalidParameter.RoleInvalid"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 参数格式不正确。
	INVALIDPARAMETERVALUE_ILLEGALFORMAT = "InvalidParameterValue.IllegalFormat"

	// 参数取值不合法。
	INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"

	// 参数值为空。
	INVALIDPARAMETERVALUE_PARAMETEREMPTY = "InvalidParameterValue.ParameterEmpty"

	// Bcos缺少参数。
	MISSINGPARAMETER_EMPTYPARAM = "MissingParameter.EmptyParam"

	// 只有所有者可以操作。
	OPERATIONDENIED_NOTOWNER = "OperationDenied.NotOwner"

	// Bcos查询不到数据。
	RESOURCENOTFOUND_EMPTYDATA = "ResourceNotFound.EmptyData"
)
