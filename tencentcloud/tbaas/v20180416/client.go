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

package v20180416

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-04-16"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewApplyChainMakerBatchUserCertRequest() (request *ApplyChainMakerBatchUserCertRequest) {
    request = &ApplyChainMakerBatchUserCertRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tbaas", APIVersion, "ApplyChainMakerBatchUserCert")
    
    
    return
}

func NewApplyChainMakerBatchUserCertResponse() (response *ApplyChainMakerBatchUserCertResponse) {
    response = &ApplyChainMakerBatchUserCertResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ApplyChainMakerBatchUserCert
// 批量申请长安链用户签名证书
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_BAASSTOPSERVING = "FailedOperation.BaaSStopServing"
//  FAILEDOPERATION_INVALIDOPERATION = "FailedOperation.InvalidOperation"
//  FAILEDOPERATION_USERINBLACKLIST = "FailedOperation.UserInBlackList"
//  FAILEDOPERATION_USERNOJOINDEMOCLUSTER = "FailedOperation.UserNoJoinDemoCluster"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
func (c *Client) ApplyChainMakerBatchUserCert(request *ApplyChainMakerBatchUserCertRequest) (response *ApplyChainMakerBatchUserCertResponse, err error) {
    return c.ApplyChainMakerBatchUserCertWithContext(context.Background(), request)
}

// ApplyChainMakerBatchUserCert
// 批量申请长安链用户签名证书
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_BAASSTOPSERVING = "FailedOperation.BaaSStopServing"
//  FAILEDOPERATION_INVALIDOPERATION = "FailedOperation.InvalidOperation"
//  FAILEDOPERATION_USERINBLACKLIST = "FailedOperation.UserInBlackList"
//  FAILEDOPERATION_USERNOJOINDEMOCLUSTER = "FailedOperation.UserNoJoinDemoCluster"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
func (c *Client) ApplyChainMakerBatchUserCertWithContext(ctx context.Context, request *ApplyChainMakerBatchUserCertRequest) (response *ApplyChainMakerBatchUserCertResponse, err error) {
    if request == nil {
        request = NewApplyChainMakerBatchUserCertRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tbaas", APIVersion, "ApplyChainMakerBatchUserCert")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyChainMakerBatchUserCert require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyChainMakerBatchUserCertResponse()
    err = c.Send(request, response)
    return
}

func NewApplyUserCertRequest() (request *ApplyUserCertRequest) {
    request = &ApplyUserCertRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tbaas", APIVersion, "ApplyUserCert")
    
    
    return
}

func NewApplyUserCertResponse() (response *ApplyUserCertResponse) {
    response = &ApplyUserCertResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ApplyUserCert
// 申请用户证书
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CACHECKCSR = "FailedOperation.CaCheckCsr"
//  FAILEDOPERATION_CADBOPTION = "FailedOperation.CaDbOption"
//  FAILEDOPERATION_CAEXSIT = "FailedOperation.CaExsit"
//  FAILEDOPERATION_CAGENKEY = "FailedOperation.CaGenkey"
//  FAILEDOPERATION_CAINIT = "FailedOperation.CaInit"
//  FAILEDOPERATION_CAINPUTPARAM = "FailedOperation.CaInputParam"
//  FAILEDOPERATION_CANOEXIST = "FailedOperation.CaNoExist"
//  FAILEDOPERATION_CAREVOKE = "FailedOperation.CaRevoke"
//  FAILEDOPERATION_CAROOTNONEXIST = "FailedOperation.CaRootNonExist"
//  FAILEDOPERATION_CASERVICE = "FailedOperation.CaService"
//  FAILEDOPERATION_CASIGNCERT = "FailedOperation.CaSignCert"
//  FAILEDOPERATION_CAYUNAPIAPPLYCERT = "FailedOperation.CaYunApiApplyCert"
//  FAILEDOPERATION_CAYUNAPICOMMON = "FailedOperation.CaYunApiCommon"
//  FAILEDOPERATION_GROUPILLEGAL = "FailedOperation.GroupIllegal"
//  FAILEDOPERATION_USERAUTHTYPE = "FailedOperation.UserAuthType"
//  INTERNALERROR_FLASKEXCEPTION = "InternalError.FlaskException"
//  INTERNALERROR_NODEFINEERROR = "InternalError.NoDefineError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INVALIDPARAMETERVALUE_ILLEGALFORMAT = "InvalidParameterValue.IllegalFormat"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  INVALIDPARAMETERVALUE_PARAMETEREMPTY = "InvalidParameterValue.ParameterEmpty"
func (c *Client) ApplyUserCert(request *ApplyUserCertRequest) (response *ApplyUserCertResponse, err error) {
    return c.ApplyUserCertWithContext(context.Background(), request)
}

// ApplyUserCert
// 申请用户证书
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CACHECKCSR = "FailedOperation.CaCheckCsr"
//  FAILEDOPERATION_CADBOPTION = "FailedOperation.CaDbOption"
//  FAILEDOPERATION_CAEXSIT = "FailedOperation.CaExsit"
//  FAILEDOPERATION_CAGENKEY = "FailedOperation.CaGenkey"
//  FAILEDOPERATION_CAINIT = "FailedOperation.CaInit"
//  FAILEDOPERATION_CAINPUTPARAM = "FailedOperation.CaInputParam"
//  FAILEDOPERATION_CANOEXIST = "FailedOperation.CaNoExist"
//  FAILEDOPERATION_CAREVOKE = "FailedOperation.CaRevoke"
//  FAILEDOPERATION_CAROOTNONEXIST = "FailedOperation.CaRootNonExist"
//  FAILEDOPERATION_CASERVICE = "FailedOperation.CaService"
//  FAILEDOPERATION_CASIGNCERT = "FailedOperation.CaSignCert"
//  FAILEDOPERATION_CAYUNAPIAPPLYCERT = "FailedOperation.CaYunApiApplyCert"
//  FAILEDOPERATION_CAYUNAPICOMMON = "FailedOperation.CaYunApiCommon"
//  FAILEDOPERATION_GROUPILLEGAL = "FailedOperation.GroupIllegal"
//  FAILEDOPERATION_USERAUTHTYPE = "FailedOperation.UserAuthType"
//  INTERNALERROR_FLASKEXCEPTION = "InternalError.FlaskException"
//  INTERNALERROR_NODEFINEERROR = "InternalError.NoDefineError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INVALIDPARAMETERVALUE_ILLEGALFORMAT = "InvalidParameterValue.IllegalFormat"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  INVALIDPARAMETERVALUE_PARAMETEREMPTY = "InvalidParameterValue.ParameterEmpty"
func (c *Client) ApplyUserCertWithContext(ctx context.Context, request *ApplyUserCertRequest) (response *ApplyUserCertResponse, err error) {
    if request == nil {
        request = NewApplyUserCertRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tbaas", APIVersion, "ApplyUserCert")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyUserCert require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyUserCertResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFabricBlockRequest() (request *DescribeFabricBlockRequest) {
    request = &DescribeFabricBlockRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tbaas", APIVersion, "DescribeFabricBlock")
    
    
    return
}

func NewDescribeFabricBlockResponse() (response *DescribeFabricBlockResponse) {
    response = &DescribeFabricBlockResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFabricBlock
// 获取Fabric某区块的详细信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_FABRICBLOCKNOEXIST = "FailedOperation.FabricBlockNoExist"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
func (c *Client) DescribeFabricBlock(request *DescribeFabricBlockRequest) (response *DescribeFabricBlockResponse, err error) {
    return c.DescribeFabricBlockWithContext(context.Background(), request)
}

// DescribeFabricBlock
// 获取Fabric某区块的详细信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_FABRICBLOCKNOEXIST = "FailedOperation.FabricBlockNoExist"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
func (c *Client) DescribeFabricBlockWithContext(ctx context.Context, request *DescribeFabricBlockRequest) (response *DescribeFabricBlockResponse, err error) {
    if request == nil {
        request = NewDescribeFabricBlockRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tbaas", APIVersion, "DescribeFabricBlock")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFabricBlock require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFabricBlockResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFabricTransactionRequest() (request *DescribeFabricTransactionRequest) {
    request = &DescribeFabricTransactionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tbaas", APIVersion, "DescribeFabricTransaction")
    
    
    return
}

func NewDescribeFabricTransactionResponse() (response *DescribeFabricTransactionResponse) {
    response = &DescribeFabricTransactionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFabricTransaction
// 获取Fabric交易的详细信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_FABRICTRANSACTIONNOEXIST = "FailedOperation.FabricTransactionNoExist"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
func (c *Client) DescribeFabricTransaction(request *DescribeFabricTransactionRequest) (response *DescribeFabricTransactionResponse, err error) {
    return c.DescribeFabricTransactionWithContext(context.Background(), request)
}

// DescribeFabricTransaction
// 获取Fabric交易的详细信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_FABRICTRANSACTIONNOEXIST = "FailedOperation.FabricTransactionNoExist"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
func (c *Client) DescribeFabricTransactionWithContext(ctx context.Context, request *DescribeFabricTransactionRequest) (response *DescribeFabricTransactionResponse, err error) {
    if request == nil {
        request = NewDescribeFabricTransactionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tbaas", APIVersion, "DescribeFabricTransaction")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFabricTransaction require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFabricTransactionResponse()
    err = c.Send(request, response)
    return
}

func NewDownloadUserCertRequest() (request *DownloadUserCertRequest) {
    request = &DownloadUserCertRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tbaas", APIVersion, "DownloadUserCert")
    
    
    return
}

func NewDownloadUserCertResponse() (response *DownloadUserCertResponse) {
    response = &DownloadUserCertResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DownloadUserCert
// 下载用户证书
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CACHECKCSR = "FailedOperation.CaCheckCsr"
//  FAILEDOPERATION_CADBOPTION = "FailedOperation.CaDbOption"
//  FAILEDOPERATION_CAEXSIT = "FailedOperation.CaExsit"
//  FAILEDOPERATION_CAGENKEY = "FailedOperation.CaGenkey"
//  FAILEDOPERATION_CAINIT = "FailedOperation.CaInit"
//  FAILEDOPERATION_CAINPUTPARAM = "FailedOperation.CaInputParam"
//  FAILEDOPERATION_CANOEXIST = "FailedOperation.CaNoExist"
//  FAILEDOPERATION_CAREVOKE = "FailedOperation.CaRevoke"
//  FAILEDOPERATION_CAROOTNONEXIST = "FailedOperation.CaRootNonExist"
//  FAILEDOPERATION_CASERVICE = "FailedOperation.CaService"
//  FAILEDOPERATION_CASIGNCERT = "FailedOperation.CaSignCert"
//  FAILEDOPERATION_CAYUNAPIAPPLYCERT = "FailedOperation.CaYunApiApplyCert"
//  FAILEDOPERATION_CAYUNAPICOMMON = "FailedOperation.CaYunApiCommon"
//  FAILEDOPERATION_GROUPILLEGAL = "FailedOperation.GroupIllegal"
//  INTERNALERROR_FLASKEXCEPTION = "InternalError.FlaskException"
//  INTERNALERROR_NODEFINEERROR = "InternalError.NoDefineError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INVALIDPARAMETERVALUE_ILLEGALFORMAT = "InvalidParameterValue.IllegalFormat"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  INVALIDPARAMETERVALUE_PARAMETEREMPTY = "InvalidParameterValue.ParameterEmpty"
func (c *Client) DownloadUserCert(request *DownloadUserCertRequest) (response *DownloadUserCertResponse, err error) {
    return c.DownloadUserCertWithContext(context.Background(), request)
}

// DownloadUserCert
// 下载用户证书
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CACHECKCSR = "FailedOperation.CaCheckCsr"
//  FAILEDOPERATION_CADBOPTION = "FailedOperation.CaDbOption"
//  FAILEDOPERATION_CAEXSIT = "FailedOperation.CaExsit"
//  FAILEDOPERATION_CAGENKEY = "FailedOperation.CaGenkey"
//  FAILEDOPERATION_CAINIT = "FailedOperation.CaInit"
//  FAILEDOPERATION_CAINPUTPARAM = "FailedOperation.CaInputParam"
//  FAILEDOPERATION_CANOEXIST = "FailedOperation.CaNoExist"
//  FAILEDOPERATION_CAREVOKE = "FailedOperation.CaRevoke"
//  FAILEDOPERATION_CAROOTNONEXIST = "FailedOperation.CaRootNonExist"
//  FAILEDOPERATION_CASERVICE = "FailedOperation.CaService"
//  FAILEDOPERATION_CASIGNCERT = "FailedOperation.CaSignCert"
//  FAILEDOPERATION_CAYUNAPIAPPLYCERT = "FailedOperation.CaYunApiApplyCert"
//  FAILEDOPERATION_CAYUNAPICOMMON = "FailedOperation.CaYunApiCommon"
//  FAILEDOPERATION_GROUPILLEGAL = "FailedOperation.GroupIllegal"
//  INTERNALERROR_FLASKEXCEPTION = "InternalError.FlaskException"
//  INTERNALERROR_NODEFINEERROR = "InternalError.NoDefineError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INVALIDPARAMETERVALUE_ILLEGALFORMAT = "InvalidParameterValue.IllegalFormat"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  INVALIDPARAMETERVALUE_PARAMETEREMPTY = "InvalidParameterValue.ParameterEmpty"
func (c *Client) DownloadUserCertWithContext(ctx context.Context, request *DownloadUserCertRequest) (response *DownloadUserCertResponse, err error) {
    if request == nil {
        request = NewDownloadUserCertRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tbaas", APIVersion, "DownloadUserCert")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DownloadUserCert require credential")
    }

    request.SetContext(ctx)
    
    response = NewDownloadUserCertResponse()
    err = c.Send(request, response)
    return
}

func NewGetBlockListRequest() (request *GetBlockListRequest) {
    request = &GetBlockListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tbaas", APIVersion, "GetBlockList")
    
    
    return
}

func NewGetBlockListResponse() (response *GetBlockListResponse) {
    response = &GetBlockListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetBlockList
// 查看当前网络下的所有区块列表，分页展示
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_FABRICBLOCKQUERY = "FailedOperation.FabricBlockQuery"
//  FAILEDOPERATION_FABRICREQUSTPARAMS = "FailedOperation.FabricRequstParams"
//  FAILEDOPERATION_GROUPILLEGAL = "FailedOperation.GroupIllegal"
//  FAILEDOPERATION_NOCHANNELPEER = "FailedOperation.NoChannelPeer"
//  FAILEDOPERATION_NOOBJECT = "FailedOperation.NoObject"
//  FAILEDOPERATION_TIMEOUTURL = "FailedOperation.TimeOutUrl"
//  INTERNALERROR_FLASKEXCEPTION = "InternalError.FlaskException"
//  INTERNALERROR_NODEFINEERROR = "InternalError.NoDefineError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INVALIDPARAMETERVALUE_ILLEGALFORMAT = "InvalidParameterValue.IllegalFormat"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  INVALIDPARAMETERVALUE_PARAMETEREMPTY = "InvalidParameterValue.ParameterEmpty"
func (c *Client) GetBlockList(request *GetBlockListRequest) (response *GetBlockListResponse, err error) {
    return c.GetBlockListWithContext(context.Background(), request)
}

// GetBlockList
// 查看当前网络下的所有区块列表，分页展示
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_FABRICBLOCKQUERY = "FailedOperation.FabricBlockQuery"
//  FAILEDOPERATION_FABRICREQUSTPARAMS = "FailedOperation.FabricRequstParams"
//  FAILEDOPERATION_GROUPILLEGAL = "FailedOperation.GroupIllegal"
//  FAILEDOPERATION_NOCHANNELPEER = "FailedOperation.NoChannelPeer"
//  FAILEDOPERATION_NOOBJECT = "FailedOperation.NoObject"
//  FAILEDOPERATION_TIMEOUTURL = "FailedOperation.TimeOutUrl"
//  INTERNALERROR_FLASKEXCEPTION = "InternalError.FlaskException"
//  INTERNALERROR_NODEFINEERROR = "InternalError.NoDefineError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INVALIDPARAMETERVALUE_ILLEGALFORMAT = "InvalidParameterValue.IllegalFormat"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  INVALIDPARAMETERVALUE_PARAMETEREMPTY = "InvalidParameterValue.ParameterEmpty"
func (c *Client) GetBlockListWithContext(ctx context.Context, request *GetBlockListRequest) (response *GetBlockListResponse, err error) {
    if request == nil {
        request = NewGetBlockListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tbaas", APIVersion, "GetBlockList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetBlockList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetBlockListResponse()
    err = c.Send(request, response)
    return
}

func NewGetBlockTransactionListForUserRequest() (request *GetBlockTransactionListForUserRequest) {
    request = &GetBlockTransactionListForUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tbaas", APIVersion, "GetBlockTransactionListForUser")
    
    
    return
}

func NewGetBlockTransactionListForUserResponse() (response *GetBlockTransactionListForUserResponse) {
    response = &GetBlockTransactionListForUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetBlockTransactionListForUser
// 获取区块内交易列表
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DatabaseException"
//  FAILEDOPERATION_FABRICBLOCKDETAIL = "FailedOperation.FabricBlockDetail"
//  FAILEDOPERATION_FABRICTRANSACTIONDETAIL = "FailedOperation.FabricTransactionDetail"
//  FAILEDOPERATION_FABRICTRANSACTIONQUERY = "FailedOperation.FabricTransactionQuery"
//  FAILEDOPERATION_GROUPILLEGAL = "FailedOperation.GroupIllegal"
//  FAILEDOPERATION_NOCHANNELPEER = "FailedOperation.NoChannelPeer"
//  FAILEDOPERATION_NOOBJECT = "FailedOperation.NoObject"
//  FAILEDOPERATION_NOPEER = "FailedOperation.NoPeer"
//  FAILEDOPERATION_STATUSNOMATCH = "FailedOperation.StatusNoMatch"
//  FAILEDOPERATION_TIMEOUTURL = "FailedOperation.TimeOutUrl"
//  INTERNALERROR_FLASKEXCEPTION = "InternalError.FlaskException"
//  INTERNALERROR_NODEFINEERROR = "InternalError.NoDefineError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALFORMAT = "InvalidParameterValue.IllegalFormat"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  INVALIDPARAMETERVALUE_PARAMETEREMPTY = "InvalidParameterValue.ParameterEmpty"
func (c *Client) GetBlockTransactionListForUser(request *GetBlockTransactionListForUserRequest) (response *GetBlockTransactionListForUserResponse, err error) {
    return c.GetBlockTransactionListForUserWithContext(context.Background(), request)
}

// GetBlockTransactionListForUser
// 获取区块内交易列表
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DatabaseException"
//  FAILEDOPERATION_FABRICBLOCKDETAIL = "FailedOperation.FabricBlockDetail"
//  FAILEDOPERATION_FABRICTRANSACTIONDETAIL = "FailedOperation.FabricTransactionDetail"
//  FAILEDOPERATION_FABRICTRANSACTIONQUERY = "FailedOperation.FabricTransactionQuery"
//  FAILEDOPERATION_GROUPILLEGAL = "FailedOperation.GroupIllegal"
//  FAILEDOPERATION_NOCHANNELPEER = "FailedOperation.NoChannelPeer"
//  FAILEDOPERATION_NOOBJECT = "FailedOperation.NoObject"
//  FAILEDOPERATION_NOPEER = "FailedOperation.NoPeer"
//  FAILEDOPERATION_STATUSNOMATCH = "FailedOperation.StatusNoMatch"
//  FAILEDOPERATION_TIMEOUTURL = "FailedOperation.TimeOutUrl"
//  INTERNALERROR_FLASKEXCEPTION = "InternalError.FlaskException"
//  INTERNALERROR_NODEFINEERROR = "InternalError.NoDefineError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALFORMAT = "InvalidParameterValue.IllegalFormat"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  INVALIDPARAMETERVALUE_PARAMETEREMPTY = "InvalidParameterValue.ParameterEmpty"
func (c *Client) GetBlockTransactionListForUserWithContext(ctx context.Context, request *GetBlockTransactionListForUserRequest) (response *GetBlockTransactionListForUserResponse, err error) {
    if request == nil {
        request = NewGetBlockTransactionListForUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tbaas", APIVersion, "GetBlockTransactionListForUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetBlockTransactionListForUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetBlockTransactionListForUserResponse()
    err = c.Send(request, response)
    return
}

func NewGetClusterSummaryRequest() (request *GetClusterSummaryRequest) {
    request = &GetClusterSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tbaas", APIVersion, "GetClusterSummary")
    
    
    return
}

func NewGetClusterSummaryResponse() (response *GetClusterSummaryResponse) {
    response = &GetClusterSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetClusterSummary
// 获取区块链网络概要
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_GROUPILLEGAL = "FailedOperation.GroupIllegal"
//  FAILEDOPERATION_NOOBJECT = "FailedOperation.NoObject"
//  FAILEDOPERATION_TIMEOUTURL = "FailedOperation.TimeOutUrl"
//  INTERNALERROR_FLASKEXCEPTION = "InternalError.FlaskException"
//  INTERNALERROR_NODEFINEERROR = "InternalError.NoDefineError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INVALIDPARAMETERVALUE_ILLEGALFORMAT = "InvalidParameterValue.IllegalFormat"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  INVALIDPARAMETERVALUE_PARAMETEREMPTY = "InvalidParameterValue.ParameterEmpty"
func (c *Client) GetClusterSummary(request *GetClusterSummaryRequest) (response *GetClusterSummaryResponse, err error) {
    return c.GetClusterSummaryWithContext(context.Background(), request)
}

// GetClusterSummary
// 获取区块链网络概要
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_GROUPILLEGAL = "FailedOperation.GroupIllegal"
//  FAILEDOPERATION_NOOBJECT = "FailedOperation.NoObject"
//  FAILEDOPERATION_TIMEOUTURL = "FailedOperation.TimeOutUrl"
//  INTERNALERROR_FLASKEXCEPTION = "InternalError.FlaskException"
//  INTERNALERROR_NODEFINEERROR = "InternalError.NoDefineError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INVALIDPARAMETERVALUE_ILLEGALFORMAT = "InvalidParameterValue.IllegalFormat"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  INVALIDPARAMETERVALUE_PARAMETEREMPTY = "InvalidParameterValue.ParameterEmpty"
func (c *Client) GetClusterSummaryWithContext(ctx context.Context, request *GetClusterSummaryRequest) (response *GetClusterSummaryResponse, err error) {
    if request == nil {
        request = NewGetClusterSummaryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tbaas", APIVersion, "GetClusterSummary")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetClusterSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetClusterSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewGetInvokeTxRequest() (request *GetInvokeTxRequest) {
    request = &GetInvokeTxRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tbaas", APIVersion, "GetInvokeTx")
    
    
    return
}

func NewGetInvokeTxResponse() (response *GetInvokeTxResponse) {
    response = &GetInvokeTxResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetInvokeTx
// Invoke异步调用结果查询
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DatabaseException"
//  FAILEDOPERATION_FABRICREQUESTPARAMS = "FailedOperation.FabricRequestParams"
//  FAILEDOPERATION_FABRICREQUSTPARAMS = "FailedOperation.FabricRequstParams"
//  FAILEDOPERATION_FABRICTXIDQUERY = "FailedOperation.FabricTxIdQuery"
//  FAILEDOPERATION_GROUPILLEGAL = "FailedOperation.GroupIllegal"
//  FAILEDOPERATION_MANAGESERVICE = "FailedOperation.ManageService"
//  FAILEDOPERATION_NOCHANNELPEER = "FailedOperation.NoChannelPeer"
//  FAILEDOPERATION_NOOBJECT = "FailedOperation.NoObject"
//  FAILEDOPERATION_TRANSACTIONTIMEOUT = "FailedOperation.TransactionTimeout"
//  INTERNALERROR_FLASKEXCEPTION = "InternalError.FlaskException"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_TRANSACTIONSERVICE = "InternalError.TransactionService"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  INVALIDPARAMETERVALUE_PARAMETEREMPTY = "InvalidParameterValue.ParameterEmpty"
func (c *Client) GetInvokeTx(request *GetInvokeTxRequest) (response *GetInvokeTxResponse, err error) {
    return c.GetInvokeTxWithContext(context.Background(), request)
}

// GetInvokeTx
// Invoke异步调用结果查询
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DatabaseException"
//  FAILEDOPERATION_FABRICREQUESTPARAMS = "FailedOperation.FabricRequestParams"
//  FAILEDOPERATION_FABRICREQUSTPARAMS = "FailedOperation.FabricRequstParams"
//  FAILEDOPERATION_FABRICTXIDQUERY = "FailedOperation.FabricTxIdQuery"
//  FAILEDOPERATION_GROUPILLEGAL = "FailedOperation.GroupIllegal"
//  FAILEDOPERATION_MANAGESERVICE = "FailedOperation.ManageService"
//  FAILEDOPERATION_NOCHANNELPEER = "FailedOperation.NoChannelPeer"
//  FAILEDOPERATION_NOOBJECT = "FailedOperation.NoObject"
//  FAILEDOPERATION_TRANSACTIONTIMEOUT = "FailedOperation.TransactionTimeout"
//  INTERNALERROR_FLASKEXCEPTION = "InternalError.FlaskException"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_TRANSACTIONSERVICE = "InternalError.TransactionService"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  INVALIDPARAMETERVALUE_PARAMETEREMPTY = "InvalidParameterValue.ParameterEmpty"
func (c *Client) GetInvokeTxWithContext(ctx context.Context, request *GetInvokeTxRequest) (response *GetInvokeTxResponse, err error) {
    if request == nil {
        request = NewGetInvokeTxRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tbaas", APIVersion, "GetInvokeTx")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetInvokeTx require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetInvokeTxResponse()
    err = c.Send(request, response)
    return
}

func NewGetLatestTransactionListRequest() (request *GetLatestTransactionListRequest) {
    request = &GetLatestTransactionListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tbaas", APIVersion, "GetLatestTransactionList")
    
    
    return
}

func NewGetLatestTransactionListResponse() (response *GetLatestTransactionListResponse) {
    response = &GetLatestTransactionListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetLatestTransactionList
// 获取fabric最新交易列表
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_FABRICBLOCKDETAIL = "FailedOperation.FabricBlockDetail"
//  FAILEDOPERATION_FABRICTRANSACTIONQUERY = "FailedOperation.FabricTransactionQuery"
//  FAILEDOPERATION_GROUPILLEGAL = "FailedOperation.GroupIllegal"
//  FAILEDOPERATION_NOCHANNELGROUP = "FailedOperation.NoChannelGroup"
//  FAILEDOPERATION_NOCHANNELPEER = "FailedOperation.NoChannelPeer"
//  FAILEDOPERATION_NOOBJECT = "FailedOperation.NoObject"
//  INTERNALERROR_FLASKEXCEPTION = "InternalError.FlaskException"
//  INTERNALERROR_NODEFINEERROR = "InternalError.NoDefineError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INVALIDPARAMETERVALUE_ILLEGALFORMAT = "InvalidParameterValue.IllegalFormat"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  INVALIDPARAMETERVALUE_PARAMETEREMPTY = "InvalidParameterValue.ParameterEmpty"
func (c *Client) GetLatestTransactionList(request *GetLatestTransactionListRequest) (response *GetLatestTransactionListResponse, err error) {
    return c.GetLatestTransactionListWithContext(context.Background(), request)
}

// GetLatestTransactionList
// 获取fabric最新交易列表
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_FABRICBLOCKDETAIL = "FailedOperation.FabricBlockDetail"
//  FAILEDOPERATION_FABRICTRANSACTIONQUERY = "FailedOperation.FabricTransactionQuery"
//  FAILEDOPERATION_GROUPILLEGAL = "FailedOperation.GroupIllegal"
//  FAILEDOPERATION_NOCHANNELGROUP = "FailedOperation.NoChannelGroup"
//  FAILEDOPERATION_NOCHANNELPEER = "FailedOperation.NoChannelPeer"
//  FAILEDOPERATION_NOOBJECT = "FailedOperation.NoObject"
//  INTERNALERROR_FLASKEXCEPTION = "InternalError.FlaskException"
//  INTERNALERROR_NODEFINEERROR = "InternalError.NoDefineError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INVALIDPARAMETERVALUE_ILLEGALFORMAT = "InvalidParameterValue.IllegalFormat"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  INVALIDPARAMETERVALUE_PARAMETEREMPTY = "InvalidParameterValue.ParameterEmpty"
func (c *Client) GetLatestTransactionListWithContext(ctx context.Context, request *GetLatestTransactionListRequest) (response *GetLatestTransactionListResponse, err error) {
    if request == nil {
        request = NewGetLatestTransactionListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tbaas", APIVersion, "GetLatestTransactionList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetLatestTransactionList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetLatestTransactionListResponse()
    err = c.Send(request, response)
    return
}

func NewGetTransactionDetailForUserRequest() (request *GetTransactionDetailForUserRequest) {
    request = &GetTransactionDetailForUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tbaas", APIVersion, "GetTransactionDetailForUser")
    
    
    return
}

func NewGetTransactionDetailForUserResponse() (response *GetTransactionDetailForUserResponse) {
    response = &GetTransactionDetailForUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetTransactionDetailForUser
// 获取交易的详情
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DatabaseException"
//  FAILEDOPERATION_FABRICBLOCKDETAIL = "FailedOperation.FabricBlockDetail"
//  FAILEDOPERATION_FABRICTRANSACTIONDETAIL = "FailedOperation.FabricTransactionDetail"
//  FAILEDOPERATION_GROUPILLEGAL = "FailedOperation.GroupIllegal"
//  FAILEDOPERATION_NOOBJECT = "FailedOperation.NoObject"
//  FAILEDOPERATION_TIMEOUTURL = "FailedOperation.TimeOutUrl"
//  INTERNALERROR_FAILURL = "InternalError.FailUrl"
//  INTERNALERROR_FLASKEXCEPTION = "InternalError.FlaskException"
//  INTERNALERROR_NODEFINEERROR = "InternalError.NoDefineError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_TRANSACTIONSERVICE = "InternalError.TransactionService"
//  INVALIDPARAMETERVALUE_ILLEGALFORMAT = "InvalidParameterValue.IllegalFormat"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  INVALIDPARAMETERVALUE_PARAMETEREMPTY = "InvalidParameterValue.ParameterEmpty"
func (c *Client) GetTransactionDetailForUser(request *GetTransactionDetailForUserRequest) (response *GetTransactionDetailForUserResponse, err error) {
    return c.GetTransactionDetailForUserWithContext(context.Background(), request)
}

// GetTransactionDetailForUser
// 获取交易的详情
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DatabaseException"
//  FAILEDOPERATION_FABRICBLOCKDETAIL = "FailedOperation.FabricBlockDetail"
//  FAILEDOPERATION_FABRICTRANSACTIONDETAIL = "FailedOperation.FabricTransactionDetail"
//  FAILEDOPERATION_GROUPILLEGAL = "FailedOperation.GroupIllegal"
//  FAILEDOPERATION_NOOBJECT = "FailedOperation.NoObject"
//  FAILEDOPERATION_TIMEOUTURL = "FailedOperation.TimeOutUrl"
//  INTERNALERROR_FAILURL = "InternalError.FailUrl"
//  INTERNALERROR_FLASKEXCEPTION = "InternalError.FlaskException"
//  INTERNALERROR_NODEFINEERROR = "InternalError.NoDefineError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_TRANSACTIONSERVICE = "InternalError.TransactionService"
//  INVALIDPARAMETERVALUE_ILLEGALFORMAT = "InvalidParameterValue.IllegalFormat"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  INVALIDPARAMETERVALUE_PARAMETEREMPTY = "InvalidParameterValue.ParameterEmpty"
func (c *Client) GetTransactionDetailForUserWithContext(ctx context.Context, request *GetTransactionDetailForUserRequest) (response *GetTransactionDetailForUserResponse, err error) {
    if request == nil {
        request = NewGetTransactionDetailForUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tbaas", APIVersion, "GetTransactionDetailForUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTransactionDetailForUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTransactionDetailForUserResponse()
    err = c.Send(request, response)
    return
}

func NewInvokeRequest() (request *InvokeRequest) {
    request = &InvokeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tbaas", APIVersion, "Invoke")
    
    
    return
}

func NewInvokeResponse() (response *InvokeResponse) {
    response = &InvokeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// Invoke
// 新增交易
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DatabaseException"
//  FAILEDOPERATION_FABRICCOMMIT = "FailedOperation.FabricCommit"
//  FAILEDOPERATION_FABRICEVENTHUB = "FailedOperation.FabricEventHub"
//  FAILEDOPERATION_FABRICPROPOSAL = "FailedOperation.FabricProposal"
//  FAILEDOPERATION_FABRICREQUESTPARAMS = "FailedOperation.FabricRequestParams"
//  FAILEDOPERATION_FABRICREQUSTPARAMS = "FailedOperation.FabricRequstParams"
//  FAILEDOPERATION_FABRICTXINVOKE = "FailedOperation.FabricTxInvoke"
//  FAILEDOPERATION_FABRICTXQUERY = "FailedOperation.FabricTxQuery"
//  FAILEDOPERATION_FABRICTXQUERYNONE = "FailedOperation.FabricTxQueryNone"
//  FAILEDOPERATION_GROUPILLEGAL = "FailedOperation.GroupIllegal"
//  FAILEDOPERATION_MANAGESERVICE = "FailedOperation.ManageService"
//  FAILEDOPERATION_NOCHAINCODECHANNEL = "FailedOperation.NoChainCodeChannel"
//  FAILEDOPERATION_NOCHAINCODEGROUP = "FailedOperation.NoChainCodeGroup"
//  FAILEDOPERATION_NOCHAINCODEPEER = "FailedOperation.NoChainCodePeer"
//  FAILEDOPERATION_NOCHANNELGROUP = "FailedOperation.NoChannelGroup"
//  FAILEDOPERATION_NOCHANNELPEER = "FailedOperation.NoChannelPeer"
//  FAILEDOPERATION_NOOBJECT = "FailedOperation.NoObject"
//  FAILEDOPERATION_NOPEER = "FailedOperation.NoPeer"
//  FAILEDOPERATION_TRANSACTIONTIMEOUT = "FailedOperation.TransactionTimeout"
//  INTERNALERROR_FLASKEXCEPTION = "InternalError.FlaskException"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_TRANSACTIONSERVICE = "InternalError.TransactionService"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  INVALIDPARAMETERVALUE_PARAMETEREMPTY = "InvalidParameterValue.ParameterEmpty"
func (c *Client) Invoke(request *InvokeRequest) (response *InvokeResponse, err error) {
    return c.InvokeWithContext(context.Background(), request)
}

// Invoke
// 新增交易
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DatabaseException"
//  FAILEDOPERATION_FABRICCOMMIT = "FailedOperation.FabricCommit"
//  FAILEDOPERATION_FABRICEVENTHUB = "FailedOperation.FabricEventHub"
//  FAILEDOPERATION_FABRICPROPOSAL = "FailedOperation.FabricProposal"
//  FAILEDOPERATION_FABRICREQUESTPARAMS = "FailedOperation.FabricRequestParams"
//  FAILEDOPERATION_FABRICREQUSTPARAMS = "FailedOperation.FabricRequstParams"
//  FAILEDOPERATION_FABRICTXINVOKE = "FailedOperation.FabricTxInvoke"
//  FAILEDOPERATION_FABRICTXQUERY = "FailedOperation.FabricTxQuery"
//  FAILEDOPERATION_FABRICTXQUERYNONE = "FailedOperation.FabricTxQueryNone"
//  FAILEDOPERATION_GROUPILLEGAL = "FailedOperation.GroupIllegal"
//  FAILEDOPERATION_MANAGESERVICE = "FailedOperation.ManageService"
//  FAILEDOPERATION_NOCHAINCODECHANNEL = "FailedOperation.NoChainCodeChannel"
//  FAILEDOPERATION_NOCHAINCODEGROUP = "FailedOperation.NoChainCodeGroup"
//  FAILEDOPERATION_NOCHAINCODEPEER = "FailedOperation.NoChainCodePeer"
//  FAILEDOPERATION_NOCHANNELGROUP = "FailedOperation.NoChannelGroup"
//  FAILEDOPERATION_NOCHANNELPEER = "FailedOperation.NoChannelPeer"
//  FAILEDOPERATION_NOOBJECT = "FailedOperation.NoObject"
//  FAILEDOPERATION_NOPEER = "FailedOperation.NoPeer"
//  FAILEDOPERATION_TRANSACTIONTIMEOUT = "FailedOperation.TransactionTimeout"
//  INTERNALERROR_FLASKEXCEPTION = "InternalError.FlaskException"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_TRANSACTIONSERVICE = "InternalError.TransactionService"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  INVALIDPARAMETERVALUE_PARAMETEREMPTY = "InvalidParameterValue.ParameterEmpty"
func (c *Client) InvokeWithContext(ctx context.Context, request *InvokeRequest) (response *InvokeResponse, err error) {
    if request == nil {
        request = NewInvokeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tbaas", APIVersion, "Invoke")
    
    if c.GetCredential() == nil {
        return nil, errors.New("Invoke require credential")
    }

    request.SetContext(ctx)
    
    response = NewInvokeResponse()
    err = c.Send(request, response)
    return
}

func NewInvokeChainMakerContractRequest() (request *InvokeChainMakerContractRequest) {
    request = &InvokeChainMakerContractRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tbaas", APIVersion, "InvokeChainMakerContract")
    
    
    return
}

func NewInvokeChainMakerContractResponse() (response *InvokeChainMakerContractResponse) {
    response = &InvokeChainMakerContractResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InvokeChainMakerContract
// 调用长安链合约执行交易
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_BCOSSERVICE = "FailedOperation.BcosService"
//  FAILEDOPERATION_CHAINMAKERCHAINCODEINVOKEFAILED = "FailedOperation.ChainMakerChaincodeInvokeFailed"
//  FAILEDOPERATION_COMPILEDEPLOYEDCONTRACT = "FailedOperation.CompileDeployedContract"
//  FAILEDOPERATION_COMPILEDEPLOYINGCONTRACT = "FailedOperation.CompileDeployingContract"
//  FAILEDOPERATION_CONTRACTEDITEDBYOTHERAGENCY = "FailedOperation.ContractEditedByOtherAgency"
//  FAILEDOPERATION_DELETEDEPLOYEDCONTRACT = "FailedOperation.DeleteDeployedContract"
//  FAILEDOPERATION_DELETEDEPLOYINGCONTRACT = "FailedOperation.DeleteDeployingContract"
//  FAILEDOPERATION_DEPLOYCONTRACTNOTCOMPILE = "FailedOperation.DeployContractNotCompile"
//  FAILEDOPERATION_FRONTREQUESTFAIL = "FailedOperation.FrontRequestFail"
//  FAILEDOPERATION_INVALIDGROUPPK = "FailedOperation.InvalidGroupPk"
//  FAILEDOPERATION_INVALIDKEYUSER = "FailedOperation.InvalidKeyUser"
//  FAILEDOPERATION_NEWCONTRACT = "FailedOperation.NewContract"
//  FAILEDOPERATION_NOTDEPLOYEDCONTRACT = "FailedOperation.NotDeployedContract"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INVALIDCONTRACTPARAM = "InternalError.InvalidContractParam"
//  INTERNALERROR_METHODTYPENOTSUPPORT = "InternalError.MethodTypeNotSupport"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INVALIDPARAMETER_ACCOUNTPARAMERROR = "InvalidParameter.AccountParamError"
//  INVALIDPARAMETER_AGENCYINVALID = "InvalidParameter.AgencyInvalid"
//  INVALIDPARAMETER_AGENCYNETPARAMINVALID = "InvalidParameter.AgencyNetParamInvalid"
//  INVALIDPARAMETER_ALLIANCEIDOFNETEMPTY = "InvalidParameter.AllianceIdOfNetEmpty"
//  INVALIDPARAMETER_CONTRACTIDINVALID = "InvalidParameter.ContractIdInvalid"
//  INVALIDPARAMETER_DATAHADEXIST = "InvalidParameter.DataHadExist"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_EXPIRETIMEINVALID = "InvalidParameter.ExpireTimeInvalid"
//  INVALIDPARAMETER_FRONTCONNFAIL = "InvalidParameter.FrontConnFail"
//  INVALIDPARAMETER_FRONTIPINVALID = "InvalidParameter.FrontIpInvalid"
//  INVALIDPARAMETER_FRONTPARAMERROR = "InvalidParameter.FrontParamError"
//  INVALIDPARAMETER_FRONTREQUESTFAIL = "InvalidParameter.FrontRequestFail"
//  INVALIDPARAMETER_INVALIDCONTRACTARG = "InvalidParameter.InvalidContractArg"
//  INVALIDPARAMETER_NETIDINVALID = "InvalidParameter.NetIdInvalid"
//  INVALIDPARAMETER_NETPARAMERROR = "InvalidParameter.NetParamError"
//  INVALIDPARAMETER_NEWKEYUSERPARAMERROR = "InvalidParameter.NewKeyUserParamError"
//  INVALIDPARAMETER_NOINFOTODELETE = "InvalidParameter.NoInfoToDelete"
//  INVALIDPARAMETER_NOTFOUNDVALIDFRONT = "InvalidParameter.NotFoundValidFront"
//  INVALIDPARAMETER_REDEPLOYEDCONTRACT = "InvalidParameter.ReDeployedContract"
//  INVALIDPARAMETER_REDEPLOYINGCONTRACT = "InvalidParameter.ReDeployingContract"
//  INVALIDPARAMETER_ROLEINVALID = "InvalidParameter.RoleInvalid"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  MISSINGPARAMETER_EMPTYPARAM = "MissingParameter.EmptyParam"
//  RESOURCENOTFOUND_EMPTYDATA = "ResourceNotFound.EmptyData"
func (c *Client) InvokeChainMakerContract(request *InvokeChainMakerContractRequest) (response *InvokeChainMakerContractResponse, err error) {
    return c.InvokeChainMakerContractWithContext(context.Background(), request)
}

// InvokeChainMakerContract
// 调用长安链合约执行交易
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_BCOSSERVICE = "FailedOperation.BcosService"
//  FAILEDOPERATION_CHAINMAKERCHAINCODEINVOKEFAILED = "FailedOperation.ChainMakerChaincodeInvokeFailed"
//  FAILEDOPERATION_COMPILEDEPLOYEDCONTRACT = "FailedOperation.CompileDeployedContract"
//  FAILEDOPERATION_COMPILEDEPLOYINGCONTRACT = "FailedOperation.CompileDeployingContract"
//  FAILEDOPERATION_CONTRACTEDITEDBYOTHERAGENCY = "FailedOperation.ContractEditedByOtherAgency"
//  FAILEDOPERATION_DELETEDEPLOYEDCONTRACT = "FailedOperation.DeleteDeployedContract"
//  FAILEDOPERATION_DELETEDEPLOYINGCONTRACT = "FailedOperation.DeleteDeployingContract"
//  FAILEDOPERATION_DEPLOYCONTRACTNOTCOMPILE = "FailedOperation.DeployContractNotCompile"
//  FAILEDOPERATION_FRONTREQUESTFAIL = "FailedOperation.FrontRequestFail"
//  FAILEDOPERATION_INVALIDGROUPPK = "FailedOperation.InvalidGroupPk"
//  FAILEDOPERATION_INVALIDKEYUSER = "FailedOperation.InvalidKeyUser"
//  FAILEDOPERATION_NEWCONTRACT = "FailedOperation.NewContract"
//  FAILEDOPERATION_NOTDEPLOYEDCONTRACT = "FailedOperation.NotDeployedContract"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INVALIDCONTRACTPARAM = "InternalError.InvalidContractParam"
//  INTERNALERROR_METHODTYPENOTSUPPORT = "InternalError.MethodTypeNotSupport"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INVALIDPARAMETER_ACCOUNTPARAMERROR = "InvalidParameter.AccountParamError"
//  INVALIDPARAMETER_AGENCYINVALID = "InvalidParameter.AgencyInvalid"
//  INVALIDPARAMETER_AGENCYNETPARAMINVALID = "InvalidParameter.AgencyNetParamInvalid"
//  INVALIDPARAMETER_ALLIANCEIDOFNETEMPTY = "InvalidParameter.AllianceIdOfNetEmpty"
//  INVALIDPARAMETER_CONTRACTIDINVALID = "InvalidParameter.ContractIdInvalid"
//  INVALIDPARAMETER_DATAHADEXIST = "InvalidParameter.DataHadExist"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_EXPIRETIMEINVALID = "InvalidParameter.ExpireTimeInvalid"
//  INVALIDPARAMETER_FRONTCONNFAIL = "InvalidParameter.FrontConnFail"
//  INVALIDPARAMETER_FRONTIPINVALID = "InvalidParameter.FrontIpInvalid"
//  INVALIDPARAMETER_FRONTPARAMERROR = "InvalidParameter.FrontParamError"
//  INVALIDPARAMETER_FRONTREQUESTFAIL = "InvalidParameter.FrontRequestFail"
//  INVALIDPARAMETER_INVALIDCONTRACTARG = "InvalidParameter.InvalidContractArg"
//  INVALIDPARAMETER_NETIDINVALID = "InvalidParameter.NetIdInvalid"
//  INVALIDPARAMETER_NETPARAMERROR = "InvalidParameter.NetParamError"
//  INVALIDPARAMETER_NEWKEYUSERPARAMERROR = "InvalidParameter.NewKeyUserParamError"
//  INVALIDPARAMETER_NOINFOTODELETE = "InvalidParameter.NoInfoToDelete"
//  INVALIDPARAMETER_NOTFOUNDVALIDFRONT = "InvalidParameter.NotFoundValidFront"
//  INVALIDPARAMETER_REDEPLOYEDCONTRACT = "InvalidParameter.ReDeployedContract"
//  INVALIDPARAMETER_REDEPLOYINGCONTRACT = "InvalidParameter.ReDeployingContract"
//  INVALIDPARAMETER_ROLEINVALID = "InvalidParameter.RoleInvalid"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  MISSINGPARAMETER_EMPTYPARAM = "MissingParameter.EmptyParam"
//  RESOURCENOTFOUND_EMPTYDATA = "ResourceNotFound.EmptyData"
func (c *Client) InvokeChainMakerContractWithContext(ctx context.Context, request *InvokeChainMakerContractRequest) (response *InvokeChainMakerContractResponse, err error) {
    if request == nil {
        request = NewInvokeChainMakerContractRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tbaas", APIVersion, "InvokeChainMakerContract")
    
    if c.GetCredential() == nil {
        return nil, errors.New("InvokeChainMakerContract require credential")
    }

    request.SetContext(ctx)
    
    response = NewInvokeChainMakerContractResponse()
    err = c.Send(request, response)
    return
}

func NewInvokeChainMakerDemoContractRequest() (request *InvokeChainMakerDemoContractRequest) {
    request = &InvokeChainMakerDemoContractRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tbaas", APIVersion, "InvokeChainMakerDemoContract")
    
    
    return
}

func NewInvokeChainMakerDemoContractResponse() (response *InvokeChainMakerDemoContractResponse) {
    response = &InvokeChainMakerDemoContractResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InvokeChainMakerDemoContract
// 调用长安链体验网络合约执行交易
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_BAASSTOPSERVING = "FailedOperation.BaaSStopServing"
//  FAILEDOPERATION_INVALIDOPERATION = "FailedOperation.InvalidOperation"
//  FAILEDOPERATION_USERINBLACKLIST = "FailedOperation.UserInBlackList"
//  FAILEDOPERATION_USERNOJOINDEMOCLUSTER = "FailedOperation.UserNoJoinDemoCluster"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INPUTDATAVIOLATION = "InvalidParameter.InputDataViolation"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
func (c *Client) InvokeChainMakerDemoContract(request *InvokeChainMakerDemoContractRequest) (response *InvokeChainMakerDemoContractResponse, err error) {
    return c.InvokeChainMakerDemoContractWithContext(context.Background(), request)
}

// InvokeChainMakerDemoContract
// 调用长安链体验网络合约执行交易
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_BAASSTOPSERVING = "FailedOperation.BaaSStopServing"
//  FAILEDOPERATION_INVALIDOPERATION = "FailedOperation.InvalidOperation"
//  FAILEDOPERATION_USERINBLACKLIST = "FailedOperation.UserInBlackList"
//  FAILEDOPERATION_USERNOJOINDEMOCLUSTER = "FailedOperation.UserNoJoinDemoCluster"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INPUTDATAVIOLATION = "InvalidParameter.InputDataViolation"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
func (c *Client) InvokeChainMakerDemoContractWithContext(ctx context.Context, request *InvokeChainMakerDemoContractRequest) (response *InvokeChainMakerDemoContractResponse, err error) {
    if request == nil {
        request = NewInvokeChainMakerDemoContractRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tbaas", APIVersion, "InvokeChainMakerDemoContract")
    
    if c.GetCredential() == nil {
        return nil, errors.New("InvokeChainMakerDemoContract require credential")
    }

    request.SetContext(ctx)
    
    response = NewInvokeChainMakerDemoContractResponse()
    err = c.Send(request, response)
    return
}

func NewInvokeFabricChaincodeRequest() (request *InvokeFabricChaincodeRequest) {
    request = &InvokeFabricChaincodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tbaas", APIVersion, "InvokeFabricChaincode")
    
    
    return
}

func NewInvokeFabricChaincodeResponse() (response *InvokeFabricChaincodeResponse) {
    response = &InvokeFabricChaincodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InvokeFabricChaincode
// 调用Fabric用户合约执行交易
//
// 可能返回的错误码:
//  FAILEDOPERATION_FABRICCHAINCODEINVOKEFAILED = "FailedOperation.FabricChaincodeInvokeFailed"
//  FAILEDOPERATION_FABRICCHAINCODENOEXIST = "FailedOperation.FabricChaincodeNoExist"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
func (c *Client) InvokeFabricChaincode(request *InvokeFabricChaincodeRequest) (response *InvokeFabricChaincodeResponse, err error) {
    return c.InvokeFabricChaincodeWithContext(context.Background(), request)
}

// InvokeFabricChaincode
// 调用Fabric用户合约执行交易
//
// 可能返回的错误码:
//  FAILEDOPERATION_FABRICCHAINCODEINVOKEFAILED = "FailedOperation.FabricChaincodeInvokeFailed"
//  FAILEDOPERATION_FABRICCHAINCODENOEXIST = "FailedOperation.FabricChaincodeNoExist"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
func (c *Client) InvokeFabricChaincodeWithContext(ctx context.Context, request *InvokeFabricChaincodeRequest) (response *InvokeFabricChaincodeResponse, err error) {
    if request == nil {
        request = NewInvokeFabricChaincodeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tbaas", APIVersion, "InvokeFabricChaincode")
    
    if c.GetCredential() == nil {
        return nil, errors.New("InvokeFabricChaincode require credential")
    }

    request.SetContext(ctx)
    
    response = NewInvokeFabricChaincodeResponse()
    err = c.Send(request, response)
    return
}

func NewQueryRequest() (request *QueryRequest) {
    request = &QueryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tbaas", APIVersion, "Query")
    
    
    return
}

func NewQueryResponse() (response *QueryResponse) {
    response = &QueryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// Query
// 查询交易
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DatabaseException"
//  FAILEDOPERATION_FABRICCOMMIT = "FailedOperation.FabricCommit"
//  FAILEDOPERATION_FABRICPROPOSAL = "FailedOperation.FabricProposal"
//  FAILEDOPERATION_FABRICREQUESTPARAMS = "FailedOperation.FabricRequestParams"
//  FAILEDOPERATION_FABRICREQUSTPARAMS = "FailedOperation.FabricRequstParams"
//  FAILEDOPERATION_FABRICTXINVOKE = "FailedOperation.FabricTxInvoke"
//  FAILEDOPERATION_FABRICTXQUERY = "FailedOperation.FabricTxQuery"
//  FAILEDOPERATION_FABRICTXQUERYNONE = "FailedOperation.FabricTxQueryNone"
//  FAILEDOPERATION_GROUPILLEGAL = "FailedOperation.GroupIllegal"
//  FAILEDOPERATION_MANAGESERVICE = "FailedOperation.ManageService"
//  FAILEDOPERATION_NOCHAINCODECHANNEL = "FailedOperation.NoChainCodeChannel"
//  FAILEDOPERATION_NOCHAINCODEGROUP = "FailedOperation.NoChainCodeGroup"
//  FAILEDOPERATION_NOCHAINCODEPEER = "FailedOperation.NoChainCodePeer"
//  FAILEDOPERATION_NOCHANNELGROUP = "FailedOperation.NoChannelGroup"
//  FAILEDOPERATION_NOCHANNELPEER = "FailedOperation.NoChannelPeer"
//  FAILEDOPERATION_NOOBJECT = "FailedOperation.NoObject"
//  FAILEDOPERATION_TRANSACTIONTIMEOUT = "FailedOperation.TransactionTimeout"
//  INTERNALERROR_FLASKEXCEPTION = "InternalError.FlaskException"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_TRANSACTIONSERVICE = "InternalError.TransactionService"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  INVALIDPARAMETERVALUE_PARAMETEREMPTY = "InvalidParameterValue.ParameterEmpty"
func (c *Client) Query(request *QueryRequest) (response *QueryResponse, err error) {
    return c.QueryWithContext(context.Background(), request)
}

// Query
// 查询交易
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DATABASEEXCEPTION = "FailedOperation.DatabaseException"
//  FAILEDOPERATION_FABRICCOMMIT = "FailedOperation.FabricCommit"
//  FAILEDOPERATION_FABRICPROPOSAL = "FailedOperation.FabricProposal"
//  FAILEDOPERATION_FABRICREQUESTPARAMS = "FailedOperation.FabricRequestParams"
//  FAILEDOPERATION_FABRICREQUSTPARAMS = "FailedOperation.FabricRequstParams"
//  FAILEDOPERATION_FABRICTXINVOKE = "FailedOperation.FabricTxInvoke"
//  FAILEDOPERATION_FABRICTXQUERY = "FailedOperation.FabricTxQuery"
//  FAILEDOPERATION_FABRICTXQUERYNONE = "FailedOperation.FabricTxQueryNone"
//  FAILEDOPERATION_GROUPILLEGAL = "FailedOperation.GroupIllegal"
//  FAILEDOPERATION_MANAGESERVICE = "FailedOperation.ManageService"
//  FAILEDOPERATION_NOCHAINCODECHANNEL = "FailedOperation.NoChainCodeChannel"
//  FAILEDOPERATION_NOCHAINCODEGROUP = "FailedOperation.NoChainCodeGroup"
//  FAILEDOPERATION_NOCHAINCODEPEER = "FailedOperation.NoChainCodePeer"
//  FAILEDOPERATION_NOCHANNELGROUP = "FailedOperation.NoChannelGroup"
//  FAILEDOPERATION_NOCHANNELPEER = "FailedOperation.NoChannelPeer"
//  FAILEDOPERATION_NOOBJECT = "FailedOperation.NoObject"
//  FAILEDOPERATION_TRANSACTIONTIMEOUT = "FailedOperation.TransactionTimeout"
//  INTERNALERROR_FLASKEXCEPTION = "InternalError.FlaskException"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_TRANSACTIONSERVICE = "InternalError.TransactionService"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  INVALIDPARAMETERVALUE_PARAMETEREMPTY = "InvalidParameterValue.ParameterEmpty"
func (c *Client) QueryWithContext(ctx context.Context, request *QueryRequest) (response *QueryResponse, err error) {
    if request == nil {
        request = NewQueryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tbaas", APIVersion, "Query")
    
    if c.GetCredential() == nil {
        return nil, errors.New("Query require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryResponse()
    err = c.Send(request, response)
    return
}

func NewQueryChainMakerBlockTransactionRequest() (request *QueryChainMakerBlockTransactionRequest) {
    request = &QueryChainMakerBlockTransactionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tbaas", APIVersion, "QueryChainMakerBlockTransaction")
    
    
    return
}

func NewQueryChainMakerBlockTransactionResponse() (response *QueryChainMakerBlockTransactionResponse) {
    response = &QueryChainMakerBlockTransactionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryChainMakerBlockTransaction
// 查询长安链指定高度区块的交易
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_BCOSSERVICE = "FailedOperation.BcosService"
//  FAILEDOPERATION_COMPILEDEPLOYEDCONTRACT = "FailedOperation.CompileDeployedContract"
//  FAILEDOPERATION_COMPILEDEPLOYINGCONTRACT = "FailedOperation.CompileDeployingContract"
//  FAILEDOPERATION_CONTRACTEDITEDBYOTHERAGENCY = "FailedOperation.ContractEditedByOtherAgency"
//  FAILEDOPERATION_DELETEDEPLOYEDCONTRACT = "FailedOperation.DeleteDeployedContract"
//  FAILEDOPERATION_DELETEDEPLOYINGCONTRACT = "FailedOperation.DeleteDeployingContract"
//  FAILEDOPERATION_DEPLOYCONTRACTNOTCOMPILE = "FailedOperation.DeployContractNotCompile"
//  FAILEDOPERATION_FRONTREQUESTFAIL = "FailedOperation.FrontRequestFail"
//  FAILEDOPERATION_INVALIDGROUPPK = "FailedOperation.InvalidGroupPk"
//  FAILEDOPERATION_INVALIDKEYUSER = "FailedOperation.InvalidKeyUser"
//  FAILEDOPERATION_NEWCONTRACT = "FailedOperation.NewContract"
//  FAILEDOPERATION_NOTDEPLOYEDCONTRACT = "FailedOperation.NotDeployedContract"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INVALIDCONTRACTPARAM = "InternalError.InvalidContractParam"
//  INTERNALERROR_METHODTYPENOTSUPPORT = "InternalError.MethodTypeNotSupport"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_ACCOUNTPARAMERROR = "InvalidParameter.AccountParamError"
//  INVALIDPARAMETER_AGENCYINVALID = "InvalidParameter.AgencyInvalid"
//  INVALIDPARAMETER_AGENCYNETPARAMINVALID = "InvalidParameter.AgencyNetParamInvalid"
//  INVALIDPARAMETER_ALLIANCEIDOFNETEMPTY = "InvalidParameter.AllianceIdOfNetEmpty"
//  INVALIDPARAMETER_CONTRACTIDINVALID = "InvalidParameter.ContractIdInvalid"
//  INVALIDPARAMETER_DATAHADEXIST = "InvalidParameter.DataHadExist"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_EXPIRETIMEINVALID = "InvalidParameter.ExpireTimeInvalid"
//  INVALIDPARAMETER_FRONTCONNFAIL = "InvalidParameter.FrontConnFail"
//  INVALIDPARAMETER_FRONTIPINVALID = "InvalidParameter.FrontIpInvalid"
//  INVALIDPARAMETER_FRONTPARAMERROR = "InvalidParameter.FrontParamError"
//  INVALIDPARAMETER_FRONTREQUESTFAIL = "InvalidParameter.FrontRequestFail"
//  INVALIDPARAMETER_INVALIDCONTRACTARG = "InvalidParameter.InvalidContractArg"
//  INVALIDPARAMETER_NETIDINVALID = "InvalidParameter.NetIdInvalid"
//  INVALIDPARAMETER_NETPARAMERROR = "InvalidParameter.NetParamError"
//  INVALIDPARAMETER_NEWKEYUSERPARAMERROR = "InvalidParameter.NewKeyUserParamError"
//  INVALIDPARAMETER_NOINFOTODELETE = "InvalidParameter.NoInfoToDelete"
//  INVALIDPARAMETER_NOTFOUNDVALIDFRONT = "InvalidParameter.NotFoundValidFront"
//  INVALIDPARAMETER_REDEPLOYEDCONTRACT = "InvalidParameter.ReDeployedContract"
//  INVALIDPARAMETER_REDEPLOYINGCONTRACT = "InvalidParameter.ReDeployingContract"
//  INVALIDPARAMETER_ROLEINVALID = "InvalidParameter.RoleInvalid"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  MISSINGPARAMETER_EMPTYPARAM = "MissingParameter.EmptyParam"
//  RESOURCENOTFOUND_EMPTYDATA = "ResourceNotFound.EmptyData"
func (c *Client) QueryChainMakerBlockTransaction(request *QueryChainMakerBlockTransactionRequest) (response *QueryChainMakerBlockTransactionResponse, err error) {
    return c.QueryChainMakerBlockTransactionWithContext(context.Background(), request)
}

// QueryChainMakerBlockTransaction
// 查询长安链指定高度区块的交易
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_BCOSSERVICE = "FailedOperation.BcosService"
//  FAILEDOPERATION_COMPILEDEPLOYEDCONTRACT = "FailedOperation.CompileDeployedContract"
//  FAILEDOPERATION_COMPILEDEPLOYINGCONTRACT = "FailedOperation.CompileDeployingContract"
//  FAILEDOPERATION_CONTRACTEDITEDBYOTHERAGENCY = "FailedOperation.ContractEditedByOtherAgency"
//  FAILEDOPERATION_DELETEDEPLOYEDCONTRACT = "FailedOperation.DeleteDeployedContract"
//  FAILEDOPERATION_DELETEDEPLOYINGCONTRACT = "FailedOperation.DeleteDeployingContract"
//  FAILEDOPERATION_DEPLOYCONTRACTNOTCOMPILE = "FailedOperation.DeployContractNotCompile"
//  FAILEDOPERATION_FRONTREQUESTFAIL = "FailedOperation.FrontRequestFail"
//  FAILEDOPERATION_INVALIDGROUPPK = "FailedOperation.InvalidGroupPk"
//  FAILEDOPERATION_INVALIDKEYUSER = "FailedOperation.InvalidKeyUser"
//  FAILEDOPERATION_NEWCONTRACT = "FailedOperation.NewContract"
//  FAILEDOPERATION_NOTDEPLOYEDCONTRACT = "FailedOperation.NotDeployedContract"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INVALIDCONTRACTPARAM = "InternalError.InvalidContractParam"
//  INTERNALERROR_METHODTYPENOTSUPPORT = "InternalError.MethodTypeNotSupport"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_ACCOUNTPARAMERROR = "InvalidParameter.AccountParamError"
//  INVALIDPARAMETER_AGENCYINVALID = "InvalidParameter.AgencyInvalid"
//  INVALIDPARAMETER_AGENCYNETPARAMINVALID = "InvalidParameter.AgencyNetParamInvalid"
//  INVALIDPARAMETER_ALLIANCEIDOFNETEMPTY = "InvalidParameter.AllianceIdOfNetEmpty"
//  INVALIDPARAMETER_CONTRACTIDINVALID = "InvalidParameter.ContractIdInvalid"
//  INVALIDPARAMETER_DATAHADEXIST = "InvalidParameter.DataHadExist"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_EXPIRETIMEINVALID = "InvalidParameter.ExpireTimeInvalid"
//  INVALIDPARAMETER_FRONTCONNFAIL = "InvalidParameter.FrontConnFail"
//  INVALIDPARAMETER_FRONTIPINVALID = "InvalidParameter.FrontIpInvalid"
//  INVALIDPARAMETER_FRONTPARAMERROR = "InvalidParameter.FrontParamError"
//  INVALIDPARAMETER_FRONTREQUESTFAIL = "InvalidParameter.FrontRequestFail"
//  INVALIDPARAMETER_INVALIDCONTRACTARG = "InvalidParameter.InvalidContractArg"
//  INVALIDPARAMETER_NETIDINVALID = "InvalidParameter.NetIdInvalid"
//  INVALIDPARAMETER_NETPARAMERROR = "InvalidParameter.NetParamError"
//  INVALIDPARAMETER_NEWKEYUSERPARAMERROR = "InvalidParameter.NewKeyUserParamError"
//  INVALIDPARAMETER_NOINFOTODELETE = "InvalidParameter.NoInfoToDelete"
//  INVALIDPARAMETER_NOTFOUNDVALIDFRONT = "InvalidParameter.NotFoundValidFront"
//  INVALIDPARAMETER_REDEPLOYEDCONTRACT = "InvalidParameter.ReDeployedContract"
//  INVALIDPARAMETER_REDEPLOYINGCONTRACT = "InvalidParameter.ReDeployingContract"
//  INVALIDPARAMETER_ROLEINVALID = "InvalidParameter.RoleInvalid"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  MISSINGPARAMETER_EMPTYPARAM = "MissingParameter.EmptyParam"
//  RESOURCENOTFOUND_EMPTYDATA = "ResourceNotFound.EmptyData"
func (c *Client) QueryChainMakerBlockTransactionWithContext(ctx context.Context, request *QueryChainMakerBlockTransactionRequest) (response *QueryChainMakerBlockTransactionResponse, err error) {
    if request == nil {
        request = NewQueryChainMakerBlockTransactionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tbaas", APIVersion, "QueryChainMakerBlockTransaction")
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryChainMakerBlockTransaction require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryChainMakerBlockTransactionResponse()
    err = c.Send(request, response)
    return
}

func NewQueryChainMakerContractRequest() (request *QueryChainMakerContractRequest) {
    request = &QueryChainMakerContractRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tbaas", APIVersion, "QueryChainMakerContract")
    
    
    return
}

func NewQueryChainMakerContractResponse() (response *QueryChainMakerContractResponse) {
    response = &QueryChainMakerContractResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryChainMakerContract
// 调用长安链合约查询
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_BCOSSERVICE = "FailedOperation.BcosService"
//  FAILEDOPERATION_CHAINMAKERCHAINCODEQUERYFAILED = "FailedOperation.ChainMakerChaincodeQueryFailed"
//  FAILEDOPERATION_COMPILEDEPLOYEDCONTRACT = "FailedOperation.CompileDeployedContract"
//  FAILEDOPERATION_COMPILEDEPLOYINGCONTRACT = "FailedOperation.CompileDeployingContract"
//  FAILEDOPERATION_CONTRACTEDITEDBYOTHERAGENCY = "FailedOperation.ContractEditedByOtherAgency"
//  FAILEDOPERATION_DELETEDEPLOYEDCONTRACT = "FailedOperation.DeleteDeployedContract"
//  FAILEDOPERATION_DELETEDEPLOYINGCONTRACT = "FailedOperation.DeleteDeployingContract"
//  FAILEDOPERATION_DEPLOYCONTRACTNOTCOMPILE = "FailedOperation.DeployContractNotCompile"
//  FAILEDOPERATION_FRONTREQUESTFAIL = "FailedOperation.FrontRequestFail"
//  FAILEDOPERATION_INVALIDGROUPPK = "FailedOperation.InvalidGroupPk"
//  FAILEDOPERATION_INVALIDKEYUSER = "FailedOperation.InvalidKeyUser"
//  FAILEDOPERATION_NEWCONTRACT = "FailedOperation.NewContract"
//  FAILEDOPERATION_NOTDEPLOYEDCONTRACT = "FailedOperation.NotDeployedContract"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INVALIDCONTRACTPARAM = "InternalError.InvalidContractParam"
//  INTERNALERROR_METHODTYPENOTSUPPORT = "InternalError.MethodTypeNotSupport"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INVALIDPARAMETER_ACCOUNTPARAMERROR = "InvalidParameter.AccountParamError"
//  INVALIDPARAMETER_AGENCYINVALID = "InvalidParameter.AgencyInvalid"
//  INVALIDPARAMETER_AGENCYNETPARAMINVALID = "InvalidParameter.AgencyNetParamInvalid"
//  INVALIDPARAMETER_ALLIANCEIDOFNETEMPTY = "InvalidParameter.AllianceIdOfNetEmpty"
//  INVALIDPARAMETER_CONTRACTIDINVALID = "InvalidParameter.ContractIdInvalid"
//  INVALIDPARAMETER_DATAHADEXIST = "InvalidParameter.DataHadExist"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_EXPIRETIMEINVALID = "InvalidParameter.ExpireTimeInvalid"
//  INVALIDPARAMETER_FRONTCONNFAIL = "InvalidParameter.FrontConnFail"
//  INVALIDPARAMETER_FRONTIPINVALID = "InvalidParameter.FrontIpInvalid"
//  INVALIDPARAMETER_FRONTPARAMERROR = "InvalidParameter.FrontParamError"
//  INVALIDPARAMETER_FRONTREQUESTFAIL = "InvalidParameter.FrontRequestFail"
//  INVALIDPARAMETER_INVALIDCONTRACTARG = "InvalidParameter.InvalidContractArg"
//  INVALIDPARAMETER_NETIDINVALID = "InvalidParameter.NetIdInvalid"
//  INVALIDPARAMETER_NETPARAMERROR = "InvalidParameter.NetParamError"
//  INVALIDPARAMETER_NEWKEYUSERPARAMERROR = "InvalidParameter.NewKeyUserParamError"
//  INVALIDPARAMETER_NOINFOTODELETE = "InvalidParameter.NoInfoToDelete"
//  INVALIDPARAMETER_NOTFOUNDVALIDFRONT = "InvalidParameter.NotFoundValidFront"
//  INVALIDPARAMETER_REDEPLOYEDCONTRACT = "InvalidParameter.ReDeployedContract"
//  INVALIDPARAMETER_REDEPLOYINGCONTRACT = "InvalidParameter.ReDeployingContract"
//  INVALIDPARAMETER_ROLEINVALID = "InvalidParameter.RoleInvalid"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  MISSINGPARAMETER_EMPTYPARAM = "MissingParameter.EmptyParam"
//  RESOURCENOTFOUND_EMPTYDATA = "ResourceNotFound.EmptyData"
func (c *Client) QueryChainMakerContract(request *QueryChainMakerContractRequest) (response *QueryChainMakerContractResponse, err error) {
    return c.QueryChainMakerContractWithContext(context.Background(), request)
}

// QueryChainMakerContract
// 调用长安链合约查询
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_BCOSSERVICE = "FailedOperation.BcosService"
//  FAILEDOPERATION_CHAINMAKERCHAINCODEQUERYFAILED = "FailedOperation.ChainMakerChaincodeQueryFailed"
//  FAILEDOPERATION_COMPILEDEPLOYEDCONTRACT = "FailedOperation.CompileDeployedContract"
//  FAILEDOPERATION_COMPILEDEPLOYINGCONTRACT = "FailedOperation.CompileDeployingContract"
//  FAILEDOPERATION_CONTRACTEDITEDBYOTHERAGENCY = "FailedOperation.ContractEditedByOtherAgency"
//  FAILEDOPERATION_DELETEDEPLOYEDCONTRACT = "FailedOperation.DeleteDeployedContract"
//  FAILEDOPERATION_DELETEDEPLOYINGCONTRACT = "FailedOperation.DeleteDeployingContract"
//  FAILEDOPERATION_DEPLOYCONTRACTNOTCOMPILE = "FailedOperation.DeployContractNotCompile"
//  FAILEDOPERATION_FRONTREQUESTFAIL = "FailedOperation.FrontRequestFail"
//  FAILEDOPERATION_INVALIDGROUPPK = "FailedOperation.InvalidGroupPk"
//  FAILEDOPERATION_INVALIDKEYUSER = "FailedOperation.InvalidKeyUser"
//  FAILEDOPERATION_NEWCONTRACT = "FailedOperation.NewContract"
//  FAILEDOPERATION_NOTDEPLOYEDCONTRACT = "FailedOperation.NotDeployedContract"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INVALIDCONTRACTPARAM = "InternalError.InvalidContractParam"
//  INTERNALERROR_METHODTYPENOTSUPPORT = "InternalError.MethodTypeNotSupport"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INVALIDPARAMETER_ACCOUNTPARAMERROR = "InvalidParameter.AccountParamError"
//  INVALIDPARAMETER_AGENCYINVALID = "InvalidParameter.AgencyInvalid"
//  INVALIDPARAMETER_AGENCYNETPARAMINVALID = "InvalidParameter.AgencyNetParamInvalid"
//  INVALIDPARAMETER_ALLIANCEIDOFNETEMPTY = "InvalidParameter.AllianceIdOfNetEmpty"
//  INVALIDPARAMETER_CONTRACTIDINVALID = "InvalidParameter.ContractIdInvalid"
//  INVALIDPARAMETER_DATAHADEXIST = "InvalidParameter.DataHadExist"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_EXPIRETIMEINVALID = "InvalidParameter.ExpireTimeInvalid"
//  INVALIDPARAMETER_FRONTCONNFAIL = "InvalidParameter.FrontConnFail"
//  INVALIDPARAMETER_FRONTIPINVALID = "InvalidParameter.FrontIpInvalid"
//  INVALIDPARAMETER_FRONTPARAMERROR = "InvalidParameter.FrontParamError"
//  INVALIDPARAMETER_FRONTREQUESTFAIL = "InvalidParameter.FrontRequestFail"
//  INVALIDPARAMETER_INVALIDCONTRACTARG = "InvalidParameter.InvalidContractArg"
//  INVALIDPARAMETER_NETIDINVALID = "InvalidParameter.NetIdInvalid"
//  INVALIDPARAMETER_NETPARAMERROR = "InvalidParameter.NetParamError"
//  INVALIDPARAMETER_NEWKEYUSERPARAMERROR = "InvalidParameter.NewKeyUserParamError"
//  INVALIDPARAMETER_NOINFOTODELETE = "InvalidParameter.NoInfoToDelete"
//  INVALIDPARAMETER_NOTFOUNDVALIDFRONT = "InvalidParameter.NotFoundValidFront"
//  INVALIDPARAMETER_REDEPLOYEDCONTRACT = "InvalidParameter.ReDeployedContract"
//  INVALIDPARAMETER_REDEPLOYINGCONTRACT = "InvalidParameter.ReDeployingContract"
//  INVALIDPARAMETER_ROLEINVALID = "InvalidParameter.RoleInvalid"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  MISSINGPARAMETER_EMPTYPARAM = "MissingParameter.EmptyParam"
//  RESOURCENOTFOUND_EMPTYDATA = "ResourceNotFound.EmptyData"
func (c *Client) QueryChainMakerContractWithContext(ctx context.Context, request *QueryChainMakerContractRequest) (response *QueryChainMakerContractResponse, err error) {
    if request == nil {
        request = NewQueryChainMakerContractRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tbaas", APIVersion, "QueryChainMakerContract")
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryChainMakerContract require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryChainMakerContractResponse()
    err = c.Send(request, response)
    return
}

func NewQueryChainMakerDemoBlockTransactionRequest() (request *QueryChainMakerDemoBlockTransactionRequest) {
    request = &QueryChainMakerDemoBlockTransactionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tbaas", APIVersion, "QueryChainMakerDemoBlockTransaction")
    
    
    return
}

func NewQueryChainMakerDemoBlockTransactionResponse() (response *QueryChainMakerDemoBlockTransactionResponse) {
    response = &QueryChainMakerDemoBlockTransactionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryChainMakerDemoBlockTransaction
// 查询长安链体验网络指定高度区块的交易
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_BAASSTOPSERVING = "FailedOperation.BaaSStopServing"
//  FAILEDOPERATION_INVALIDOPERATION = "FailedOperation.InvalidOperation"
//  FAILEDOPERATION_USERINBLACKLIST = "FailedOperation.UserInBlackList"
//  FAILEDOPERATION_USERNOJOINDEMOCLUSTER = "FailedOperation.UserNoJoinDemoCluster"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
func (c *Client) QueryChainMakerDemoBlockTransaction(request *QueryChainMakerDemoBlockTransactionRequest) (response *QueryChainMakerDemoBlockTransactionResponse, err error) {
    return c.QueryChainMakerDemoBlockTransactionWithContext(context.Background(), request)
}

// QueryChainMakerDemoBlockTransaction
// 查询长安链体验网络指定高度区块的交易
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_BAASSTOPSERVING = "FailedOperation.BaaSStopServing"
//  FAILEDOPERATION_INVALIDOPERATION = "FailedOperation.InvalidOperation"
//  FAILEDOPERATION_USERINBLACKLIST = "FailedOperation.UserInBlackList"
//  FAILEDOPERATION_USERNOJOINDEMOCLUSTER = "FailedOperation.UserNoJoinDemoCluster"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
func (c *Client) QueryChainMakerDemoBlockTransactionWithContext(ctx context.Context, request *QueryChainMakerDemoBlockTransactionRequest) (response *QueryChainMakerDemoBlockTransactionResponse, err error) {
    if request == nil {
        request = NewQueryChainMakerDemoBlockTransactionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tbaas", APIVersion, "QueryChainMakerDemoBlockTransaction")
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryChainMakerDemoBlockTransaction require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryChainMakerDemoBlockTransactionResponse()
    err = c.Send(request, response)
    return
}

func NewQueryChainMakerDemoContractRequest() (request *QueryChainMakerDemoContractRequest) {
    request = &QueryChainMakerDemoContractRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tbaas", APIVersion, "QueryChainMakerDemoContract")
    
    
    return
}

func NewQueryChainMakerDemoContractResponse() (response *QueryChainMakerDemoContractResponse) {
    response = &QueryChainMakerDemoContractResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryChainMakerDemoContract
// 调用长安链体验网络合约查询
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_BAASSTOPSERVING = "FailedOperation.BaaSStopServing"
//  FAILEDOPERATION_INVALIDOPERATION = "FailedOperation.InvalidOperation"
//  FAILEDOPERATION_USERINBLACKLIST = "FailedOperation.UserInBlackList"
//  FAILEDOPERATION_USERNOJOINDEMOCLUSTER = "FailedOperation.UserNoJoinDemoCluster"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
func (c *Client) QueryChainMakerDemoContract(request *QueryChainMakerDemoContractRequest) (response *QueryChainMakerDemoContractResponse, err error) {
    return c.QueryChainMakerDemoContractWithContext(context.Background(), request)
}

// QueryChainMakerDemoContract
// 调用长安链体验网络合约查询
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_BAASSTOPSERVING = "FailedOperation.BaaSStopServing"
//  FAILEDOPERATION_INVALIDOPERATION = "FailedOperation.InvalidOperation"
//  FAILEDOPERATION_USERINBLACKLIST = "FailedOperation.UserInBlackList"
//  FAILEDOPERATION_USERNOJOINDEMOCLUSTER = "FailedOperation.UserNoJoinDemoCluster"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
func (c *Client) QueryChainMakerDemoContractWithContext(ctx context.Context, request *QueryChainMakerDemoContractRequest) (response *QueryChainMakerDemoContractResponse, err error) {
    if request == nil {
        request = NewQueryChainMakerDemoContractRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tbaas", APIVersion, "QueryChainMakerDemoContract")
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryChainMakerDemoContract require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryChainMakerDemoContractResponse()
    err = c.Send(request, response)
    return
}

func NewQueryChainMakerDemoTransactionRequest() (request *QueryChainMakerDemoTransactionRequest) {
    request = &QueryChainMakerDemoTransactionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tbaas", APIVersion, "QueryChainMakerDemoTransaction")
    
    
    return
}

func NewQueryChainMakerDemoTransactionResponse() (response *QueryChainMakerDemoTransactionResponse) {
    response = &QueryChainMakerDemoTransactionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryChainMakerDemoTransaction
// 通过交易ID查询长安链体验网络交易
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_BAASSTOPSERVING = "FailedOperation.BaaSStopServing"
//  FAILEDOPERATION_INVALIDOPERATION = "FailedOperation.InvalidOperation"
//  FAILEDOPERATION_USERINBLACKLIST = "FailedOperation.UserInBlackList"
//  FAILEDOPERATION_USERNOJOINDEMOCLUSTER = "FailedOperation.UserNoJoinDemoCluster"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
func (c *Client) QueryChainMakerDemoTransaction(request *QueryChainMakerDemoTransactionRequest) (response *QueryChainMakerDemoTransactionResponse, err error) {
    return c.QueryChainMakerDemoTransactionWithContext(context.Background(), request)
}

// QueryChainMakerDemoTransaction
// 通过交易ID查询长安链体验网络交易
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_BAASSTOPSERVING = "FailedOperation.BaaSStopServing"
//  FAILEDOPERATION_INVALIDOPERATION = "FailedOperation.InvalidOperation"
//  FAILEDOPERATION_USERINBLACKLIST = "FailedOperation.UserInBlackList"
//  FAILEDOPERATION_USERNOJOINDEMOCLUSTER = "FailedOperation.UserNoJoinDemoCluster"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
func (c *Client) QueryChainMakerDemoTransactionWithContext(ctx context.Context, request *QueryChainMakerDemoTransactionRequest) (response *QueryChainMakerDemoTransactionResponse, err error) {
    if request == nil {
        request = NewQueryChainMakerDemoTransactionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tbaas", APIVersion, "QueryChainMakerDemoTransaction")
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryChainMakerDemoTransaction require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryChainMakerDemoTransactionResponse()
    err = c.Send(request, response)
    return
}

func NewQueryChainMakerTransactionRequest() (request *QueryChainMakerTransactionRequest) {
    request = &QueryChainMakerTransactionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tbaas", APIVersion, "QueryChainMakerTransaction")
    
    
    return
}

func NewQueryChainMakerTransactionResponse() (response *QueryChainMakerTransactionResponse) {
    response = &QueryChainMakerTransactionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryChainMakerTransaction
// 通过交易ID查询长安链交易
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_BCOSSERVICE = "FailedOperation.BcosService"
//  FAILEDOPERATION_COMPILEDEPLOYEDCONTRACT = "FailedOperation.CompileDeployedContract"
//  FAILEDOPERATION_COMPILEDEPLOYINGCONTRACT = "FailedOperation.CompileDeployingContract"
//  FAILEDOPERATION_CONTRACTEDITEDBYOTHERAGENCY = "FailedOperation.ContractEditedByOtherAgency"
//  FAILEDOPERATION_DELETEDEPLOYEDCONTRACT = "FailedOperation.DeleteDeployedContract"
//  FAILEDOPERATION_DELETEDEPLOYINGCONTRACT = "FailedOperation.DeleteDeployingContract"
//  FAILEDOPERATION_DEPLOYCONTRACTNOTCOMPILE = "FailedOperation.DeployContractNotCompile"
//  FAILEDOPERATION_FRONTREQUESTFAIL = "FailedOperation.FrontRequestFail"
//  FAILEDOPERATION_INVALIDGROUPPK = "FailedOperation.InvalidGroupPk"
//  FAILEDOPERATION_INVALIDKEYUSER = "FailedOperation.InvalidKeyUser"
//  FAILEDOPERATION_NEWCONTRACT = "FailedOperation.NewContract"
//  FAILEDOPERATION_NOTDEPLOYEDCONTRACT = "FailedOperation.NotDeployedContract"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INVALIDCONTRACTPARAM = "InternalError.InvalidContractParam"
//  INTERNALERROR_METHODTYPENOTSUPPORT = "InternalError.MethodTypeNotSupport"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INVALIDPARAMETER_ACCOUNTPARAMERROR = "InvalidParameter.AccountParamError"
//  INVALIDPARAMETER_AGENCYINVALID = "InvalidParameter.AgencyInvalid"
//  INVALIDPARAMETER_AGENCYNETPARAMINVALID = "InvalidParameter.AgencyNetParamInvalid"
//  INVALIDPARAMETER_ALLIANCEIDOFNETEMPTY = "InvalidParameter.AllianceIdOfNetEmpty"
//  INVALIDPARAMETER_CONTRACTIDINVALID = "InvalidParameter.ContractIdInvalid"
//  INVALIDPARAMETER_DATAHADEXIST = "InvalidParameter.DataHadExist"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_EXPIRETIMEINVALID = "InvalidParameter.ExpireTimeInvalid"
//  INVALIDPARAMETER_FRONTCONNFAIL = "InvalidParameter.FrontConnFail"
//  INVALIDPARAMETER_FRONTIPINVALID = "InvalidParameter.FrontIpInvalid"
//  INVALIDPARAMETER_FRONTPARAMERROR = "InvalidParameter.FrontParamError"
//  INVALIDPARAMETER_FRONTREQUESTFAIL = "InvalidParameter.FrontRequestFail"
//  INVALIDPARAMETER_INVALIDCONTRACTARG = "InvalidParameter.InvalidContractArg"
//  INVALIDPARAMETER_NETIDINVALID = "InvalidParameter.NetIdInvalid"
//  INVALIDPARAMETER_NETPARAMERROR = "InvalidParameter.NetParamError"
//  INVALIDPARAMETER_NEWKEYUSERPARAMERROR = "InvalidParameter.NewKeyUserParamError"
//  INVALIDPARAMETER_NOINFOTODELETE = "InvalidParameter.NoInfoToDelete"
//  INVALIDPARAMETER_NOTFOUNDVALIDFRONT = "InvalidParameter.NotFoundValidFront"
//  INVALIDPARAMETER_REDEPLOYEDCONTRACT = "InvalidParameter.ReDeployedContract"
//  INVALIDPARAMETER_REDEPLOYINGCONTRACT = "InvalidParameter.ReDeployingContract"
//  INVALIDPARAMETER_ROLEINVALID = "InvalidParameter.RoleInvalid"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  MISSINGPARAMETER_EMPTYPARAM = "MissingParameter.EmptyParam"
//  RESOURCENOTFOUND_EMPTYDATA = "ResourceNotFound.EmptyData"
func (c *Client) QueryChainMakerTransaction(request *QueryChainMakerTransactionRequest) (response *QueryChainMakerTransactionResponse, err error) {
    return c.QueryChainMakerTransactionWithContext(context.Background(), request)
}

// QueryChainMakerTransaction
// 通过交易ID查询长安链交易
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_BCOSSERVICE = "FailedOperation.BcosService"
//  FAILEDOPERATION_COMPILEDEPLOYEDCONTRACT = "FailedOperation.CompileDeployedContract"
//  FAILEDOPERATION_COMPILEDEPLOYINGCONTRACT = "FailedOperation.CompileDeployingContract"
//  FAILEDOPERATION_CONTRACTEDITEDBYOTHERAGENCY = "FailedOperation.ContractEditedByOtherAgency"
//  FAILEDOPERATION_DELETEDEPLOYEDCONTRACT = "FailedOperation.DeleteDeployedContract"
//  FAILEDOPERATION_DELETEDEPLOYINGCONTRACT = "FailedOperation.DeleteDeployingContract"
//  FAILEDOPERATION_DEPLOYCONTRACTNOTCOMPILE = "FailedOperation.DeployContractNotCompile"
//  FAILEDOPERATION_FRONTREQUESTFAIL = "FailedOperation.FrontRequestFail"
//  FAILEDOPERATION_INVALIDGROUPPK = "FailedOperation.InvalidGroupPk"
//  FAILEDOPERATION_INVALIDKEYUSER = "FailedOperation.InvalidKeyUser"
//  FAILEDOPERATION_NEWCONTRACT = "FailedOperation.NewContract"
//  FAILEDOPERATION_NOTDEPLOYEDCONTRACT = "FailedOperation.NotDeployedContract"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INVALIDCONTRACTPARAM = "InternalError.InvalidContractParam"
//  INTERNALERROR_METHODTYPENOTSUPPORT = "InternalError.MethodTypeNotSupport"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INVALIDPARAMETER_ACCOUNTPARAMERROR = "InvalidParameter.AccountParamError"
//  INVALIDPARAMETER_AGENCYINVALID = "InvalidParameter.AgencyInvalid"
//  INVALIDPARAMETER_AGENCYNETPARAMINVALID = "InvalidParameter.AgencyNetParamInvalid"
//  INVALIDPARAMETER_ALLIANCEIDOFNETEMPTY = "InvalidParameter.AllianceIdOfNetEmpty"
//  INVALIDPARAMETER_CONTRACTIDINVALID = "InvalidParameter.ContractIdInvalid"
//  INVALIDPARAMETER_DATAHADEXIST = "InvalidParameter.DataHadExist"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_EXPIRETIMEINVALID = "InvalidParameter.ExpireTimeInvalid"
//  INVALIDPARAMETER_FRONTCONNFAIL = "InvalidParameter.FrontConnFail"
//  INVALIDPARAMETER_FRONTIPINVALID = "InvalidParameter.FrontIpInvalid"
//  INVALIDPARAMETER_FRONTPARAMERROR = "InvalidParameter.FrontParamError"
//  INVALIDPARAMETER_FRONTREQUESTFAIL = "InvalidParameter.FrontRequestFail"
//  INVALIDPARAMETER_INVALIDCONTRACTARG = "InvalidParameter.InvalidContractArg"
//  INVALIDPARAMETER_NETIDINVALID = "InvalidParameter.NetIdInvalid"
//  INVALIDPARAMETER_NETPARAMERROR = "InvalidParameter.NetParamError"
//  INVALIDPARAMETER_NEWKEYUSERPARAMERROR = "InvalidParameter.NewKeyUserParamError"
//  INVALIDPARAMETER_NOINFOTODELETE = "InvalidParameter.NoInfoToDelete"
//  INVALIDPARAMETER_NOTFOUNDVALIDFRONT = "InvalidParameter.NotFoundValidFront"
//  INVALIDPARAMETER_REDEPLOYEDCONTRACT = "InvalidParameter.ReDeployedContract"
//  INVALIDPARAMETER_REDEPLOYINGCONTRACT = "InvalidParameter.ReDeployingContract"
//  INVALIDPARAMETER_ROLEINVALID = "InvalidParameter.RoleInvalid"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  MISSINGPARAMETER_EMPTYPARAM = "MissingParameter.EmptyParam"
//  RESOURCENOTFOUND_EMPTYDATA = "ResourceNotFound.EmptyData"
func (c *Client) QueryChainMakerTransactionWithContext(ctx context.Context, request *QueryChainMakerTransactionRequest) (response *QueryChainMakerTransactionResponse, err error) {
    if request == nil {
        request = NewQueryChainMakerTransactionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tbaas", APIVersion, "QueryChainMakerTransaction")
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryChainMakerTransaction require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryChainMakerTransactionResponse()
    err = c.Send(request, response)
    return
}

func NewQueryFabricChaincodeRequest() (request *QueryFabricChaincodeRequest) {
    request = &QueryFabricChaincodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tbaas", APIVersion, "QueryFabricChaincode")
    
    
    return
}

func NewQueryFabricChaincodeResponse() (response *QueryFabricChaincodeResponse) {
    response = &QueryFabricChaincodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryFabricChaincode
// 调用Fabric用户合约查询
//
// 可能返回的错误码:
//  FAILEDOPERATION_FABRICCHAINCODENOEXIST = "FailedOperation.FabricChaincodeNoExist"
//  FAILEDOPERATION_FABRICCHAINCODEQUERYFAILED = "FailedOperation.FabricChaincodeQueryFailed"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
func (c *Client) QueryFabricChaincode(request *QueryFabricChaincodeRequest) (response *QueryFabricChaincodeResponse, err error) {
    return c.QueryFabricChaincodeWithContext(context.Background(), request)
}

// QueryFabricChaincode
// 调用Fabric用户合约查询
//
// 可能返回的错误码:
//  FAILEDOPERATION_FABRICCHAINCODENOEXIST = "FailedOperation.FabricChaincodeNoExist"
//  FAILEDOPERATION_FABRICCHAINCODEQUERYFAILED = "FailedOperation.FabricChaincodeQueryFailed"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
func (c *Client) QueryFabricChaincodeWithContext(ctx context.Context, request *QueryFabricChaincodeRequest) (response *QueryFabricChaincodeResponse, err error) {
    if request == nil {
        request = NewQueryFabricChaincodeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tbaas", APIVersion, "QueryFabricChaincode")
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryFabricChaincode require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryFabricChaincodeResponse()
    err = c.Send(request, response)
    return
}

func NewSrvInvokeRequest() (request *SrvInvokeRequest) {
    request = &SrvInvokeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tbaas", APIVersion, "SrvInvoke")
    
    
    return
}

func NewSrvInvokeResponse() (response *SrvInvokeResponse) {
    response = &SrvInvokeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SrvInvoke
// trustsql服务统一接口
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_SERVICEFAILED = "FailedOperation.ServiceFailed"
//  OPERATIONDENIED_NOTOWNER = "OperationDenied.NotOwner"
func (c *Client) SrvInvoke(request *SrvInvokeRequest) (response *SrvInvokeResponse, err error) {
    return c.SrvInvokeWithContext(context.Background(), request)
}

// SrvInvoke
// trustsql服务统一接口
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_SERVICEFAILED = "FailedOperation.ServiceFailed"
//  OPERATIONDENIED_NOTOWNER = "OperationDenied.NotOwner"
func (c *Client) SrvInvokeWithContext(ctx context.Context, request *SrvInvokeRequest) (response *SrvInvokeResponse, err error) {
    if request == nil {
        request = NewSrvInvokeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tbaas", APIVersion, "SrvInvoke")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SrvInvoke require credential")
    }

    request.SetContext(ctx)
    
    response = NewSrvInvokeResponse()
    err = c.Send(request, response)
    return
}
