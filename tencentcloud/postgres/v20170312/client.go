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

package v20170312

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2017-03-12"

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


func NewAddDBInstanceToReadOnlyGroupRequest() (request *AddDBInstanceToReadOnlyGroupRequest) {
    request = &AddDBInstanceToReadOnlyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "AddDBInstanceToReadOnlyGroup")
    
    
    return
}

func NewAddDBInstanceToReadOnlyGroupResponse() (response *AddDBInstanceToReadOnlyGroupResponse) {
    response = &AddDBInstanceToReadOnlyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddDBInstanceToReadOnlyGroup
// 本接口（AddDBInstanceToReadOnlyGroup）用于添加只读实例到只读组
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_ROGROUPMASTERINSTANCENOTRIGHT = "FailedOperation.ROGroupMasterInstanceNotRight"
//  FAILEDOPERATION_ROGROUPNOTFOUNDERROR = "FailedOperation.ROGroupNotFoundError"
//  FAILEDOPERATION_ROINSTANCEHASINROGROUPERROR = "FailedOperation.ROInstanceHasInROGroupError"
//  INSTANCENOTEXIST = "InstanceNotExist"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDREADONLYGROUPSTATUS = "InvalidParameterValue.InvalidReadOnlyGroupStatus"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_READONLYGROUPNOTEXIST = "InvalidParameterValue.ReadOnlyGroupNotExist"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  OPERATIONDENIED_ROGROUPSTATUSERROR = "OperationDenied.ROGroupStatusError"
//  OPERATIONDENIED_ROINSTANCESTATUSLIMITOPERROR = "OperationDenied.ROInstanceStatusLimitOpError"
//  OPERATIONDENIED_ROINSTANCECOUNTEXEEDERROR = "OperationDenied.RoInstanceCountExeedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
func (c *Client) AddDBInstanceToReadOnlyGroup(request *AddDBInstanceToReadOnlyGroupRequest) (response *AddDBInstanceToReadOnlyGroupResponse, err error) {
    return c.AddDBInstanceToReadOnlyGroupWithContext(context.Background(), request)
}

// AddDBInstanceToReadOnlyGroup
// 本接口（AddDBInstanceToReadOnlyGroup）用于添加只读实例到只读组
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_ROGROUPMASTERINSTANCENOTRIGHT = "FailedOperation.ROGroupMasterInstanceNotRight"
//  FAILEDOPERATION_ROGROUPNOTFOUNDERROR = "FailedOperation.ROGroupNotFoundError"
//  FAILEDOPERATION_ROINSTANCEHASINROGROUPERROR = "FailedOperation.ROInstanceHasInROGroupError"
//  INSTANCENOTEXIST = "InstanceNotExist"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDREADONLYGROUPSTATUS = "InvalidParameterValue.InvalidReadOnlyGroupStatus"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_READONLYGROUPNOTEXIST = "InvalidParameterValue.ReadOnlyGroupNotExist"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  OPERATIONDENIED_ROGROUPSTATUSERROR = "OperationDenied.ROGroupStatusError"
//  OPERATIONDENIED_ROINSTANCESTATUSLIMITOPERROR = "OperationDenied.ROInstanceStatusLimitOpError"
//  OPERATIONDENIED_ROINSTANCECOUNTEXEEDERROR = "OperationDenied.RoInstanceCountExeedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
func (c *Client) AddDBInstanceToReadOnlyGroupWithContext(ctx context.Context, request *AddDBInstanceToReadOnlyGroupRequest) (response *AddDBInstanceToReadOnlyGroupResponse, err error) {
    if request == nil {
        request = NewAddDBInstanceToReadOnlyGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddDBInstanceToReadOnlyGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddDBInstanceToReadOnlyGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCloneDBInstanceRequest() (request *CloneDBInstanceRequest) {
    request = &CloneDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "CloneDBInstance")
    
    
    return
}

func NewCloneDBInstanceResponse() (response *CloneDBInstanceResponse) {
    response = &CloneDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CloneDBInstance
// 用于克隆实例，支持指定备份集、指定时间点进行克隆。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_ALLOCATEQUOTASERROR = "FailedOperation.AllocateQuotasError"
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_CDBCGWCONNECTERROR = "FailedOperation.CdbCgwConnectError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_GETSUBNETERROR = "FailedOperation.GetSubnetError"
//  FAILEDOPERATION_GETVPCINFOERROR = "FailedOperation.GetVpcInfoError"
//  FAILEDOPERATION_QUERYSPECERROR = "FailedOperation.QuerySpecError"
//  FAILEDOPERATION_QUERYVPCFAILED = "FailedOperation.QueryVpcFailed"
//  FAILEDOPERATION_QUERYVPCFALIED = "FailedOperation.QueryVpcFalied"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_ILLEGALPROJECTID = "InvalidParameterValue.IllegalProjectId"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  INVALIDPARAMETERVALUE_SPECNOTRECOGNIZEDERROR = "InvalidParameterValue.SpecNotRecognizedError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
func (c *Client) CloneDBInstance(request *CloneDBInstanceRequest) (response *CloneDBInstanceResponse, err error) {
    return c.CloneDBInstanceWithContext(context.Background(), request)
}

// CloneDBInstance
// 用于克隆实例，支持指定备份集、指定时间点进行克隆。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_ALLOCATEQUOTASERROR = "FailedOperation.AllocateQuotasError"
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_CDBCGWCONNECTERROR = "FailedOperation.CdbCgwConnectError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_GETSUBNETERROR = "FailedOperation.GetSubnetError"
//  FAILEDOPERATION_GETVPCINFOERROR = "FailedOperation.GetVpcInfoError"
//  FAILEDOPERATION_QUERYSPECERROR = "FailedOperation.QuerySpecError"
//  FAILEDOPERATION_QUERYVPCFAILED = "FailedOperation.QueryVpcFailed"
//  FAILEDOPERATION_QUERYVPCFALIED = "FailedOperation.QueryVpcFalied"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_ILLEGALPROJECTID = "InvalidParameterValue.IllegalProjectId"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  INVALIDPARAMETERVALUE_SPECNOTRECOGNIZEDERROR = "InvalidParameterValue.SpecNotRecognizedError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
func (c *Client) CloneDBInstanceWithContext(ctx context.Context, request *CloneDBInstanceRequest) (response *CloneDBInstanceResponse, err error) {
    if request == nil {
        request = NewCloneDBInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloneDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloneDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCloseDBExtranetAccessRequest() (request *CloseDBExtranetAccessRequest) {
    request = &CloseDBExtranetAccessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "CloseDBExtranetAccess")
    
    
    return
}

func NewCloseDBExtranetAccessResponse() (response *CloseDBExtranetAccessResponse) {
    response = &CloseDBExtranetAccessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CloseDBExtranetAccess
// 本接口（CloseDBExtranetAccess）用于关闭实例公网地址。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_FLOWERROR = "InternalError.FlowError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
func (c *Client) CloseDBExtranetAccess(request *CloseDBExtranetAccessRequest) (response *CloseDBExtranetAccessResponse, err error) {
    return c.CloseDBExtranetAccessWithContext(context.Background(), request)
}

// CloseDBExtranetAccess
// 本接口（CloseDBExtranetAccess）用于关闭实例公网地址。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_FLOWERROR = "InternalError.FlowError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
func (c *Client) CloseDBExtranetAccessWithContext(ctx context.Context, request *CloseDBExtranetAccessRequest) (response *CloseDBExtranetAccessResponse, err error) {
    if request == nil {
        request = NewCloseDBExtranetAccessRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseDBExtranetAccess require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloseDBExtranetAccessResponse()
    err = c.Send(request, response)
    return
}

func NewCloseServerlessDBExtranetAccessRequest() (request *CloseServerlessDBExtranetAccessRequest) {
    request = &CloseServerlessDBExtranetAccessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "CloseServerlessDBExtranetAccess")
    
    
    return
}

func NewCloseServerlessDBExtranetAccessResponse() (response *CloseServerlessDBExtranetAccessResponse) {
    response = &CloseServerlessDBExtranetAccessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CloseServerlessDBExtranetAccess
// 本接口（CloseServerlessDBExtranetAccess）用于关闭serverlessDB实例公网地址
//
// 可能返回的错误码:
//  ACCOUNTNOTEXIST = "AccountNotExist"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DBERROR = "DBError"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_INVALIDTRADEOPERATE = "FailedOperation.InvalidTradeOperate"
//  FAILEDOPERATION_LIMITOPERATION = "FailedOperation.LimitOperation"
//  FAILEDOPERATION_QUERYVPCFALIED = "FailedOperation.QueryVpcFalied"
//  FLOWERROR = "FlowError"
//  INSTANCENOTEXIST = "InstanceNotExist"
//  INTERFACENAMENOTFOUND = "InterfaceNameNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_FLOWERROR = "InternalError.FlowError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDACCOUNTSTATUS = "InvalidAccountStatus"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INTERFACENAMENOTFOUND = "InvalidParameterValue.InterfaceNameNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_STRUCTPARSEFAILED = "InvalidParameterValue.StructParseFailed"
//  LIMITOPERATION = "LimitOperation"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
//  RESOURCEUNAVAILABLE_RESOURCENOPERMISSION = "ResourceUnavailable.ResourceNoPermission"
//  RESOURCEUNAVAILABLE_VPCRESOURCENOTFOUND = "ResourceUnavailable.VpcResourceNotFound"
//  STRUCTPARSEFAILED = "StructParseFailed"
//  SYSTEMERROR = "SystemError"
//  UNKNOWNERROR = "UnknownError"
//  VPCERROR = "VpcError"
func (c *Client) CloseServerlessDBExtranetAccess(request *CloseServerlessDBExtranetAccessRequest) (response *CloseServerlessDBExtranetAccessResponse, err error) {
    return c.CloseServerlessDBExtranetAccessWithContext(context.Background(), request)
}

// CloseServerlessDBExtranetAccess
// 本接口（CloseServerlessDBExtranetAccess）用于关闭serverlessDB实例公网地址
//
// 可能返回的错误码:
//  ACCOUNTNOTEXIST = "AccountNotExist"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DBERROR = "DBError"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_INVALIDTRADEOPERATE = "FailedOperation.InvalidTradeOperate"
//  FAILEDOPERATION_LIMITOPERATION = "FailedOperation.LimitOperation"
//  FAILEDOPERATION_QUERYVPCFALIED = "FailedOperation.QueryVpcFalied"
//  FLOWERROR = "FlowError"
//  INSTANCENOTEXIST = "InstanceNotExist"
//  INTERFACENAMENOTFOUND = "InterfaceNameNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_FLOWERROR = "InternalError.FlowError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDACCOUNTSTATUS = "InvalidAccountStatus"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INTERFACENAMENOTFOUND = "InvalidParameterValue.InterfaceNameNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_STRUCTPARSEFAILED = "InvalidParameterValue.StructParseFailed"
//  LIMITOPERATION = "LimitOperation"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
//  RESOURCEUNAVAILABLE_RESOURCENOPERMISSION = "ResourceUnavailable.ResourceNoPermission"
//  RESOURCEUNAVAILABLE_VPCRESOURCENOTFOUND = "ResourceUnavailable.VpcResourceNotFound"
//  STRUCTPARSEFAILED = "StructParseFailed"
//  SYSTEMERROR = "SystemError"
//  UNKNOWNERROR = "UnknownError"
//  VPCERROR = "VpcError"
func (c *Client) CloseServerlessDBExtranetAccessWithContext(ctx context.Context, request *CloseServerlessDBExtranetAccessRequest) (response *CloseServerlessDBExtranetAccessResponse, err error) {
    if request == nil {
        request = NewCloseServerlessDBExtranetAccessRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseServerlessDBExtranetAccess require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloseServerlessDBExtranetAccessResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAccountRequest() (request *CreateAccountRequest) {
    request = &CreateAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "CreateAccount")
    
    
    return
}

func NewCreateAccountResponse() (response *CreateAccountResponse) {
    response = &CreateAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAccount
// 此接口用于创建数据账号，返回的Oid为账号唯一标识。与数据库系统表pg_roles中记录的oid一致。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_PRECHECKERROR = "FailedOperation.PreCheckError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_DECRYPTPASSWORDFAILED = "InvalidParameterValue.DecryptPasswordFailed"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAMEERROR = "InvalidParameterValue.InvalidAccountNameError"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAMEFORMATERROR = "InvalidParameterValue.InvalidAccountNameFormatError"
//  INVALIDPARAMETERVALUE_INVALIDDATABASENAMEFORMATERROR = "InvalidParameterValue.InvalidDatabaseNameFormatError"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORDLENGTHERROR = "InvalidParameterValue.InvalidPasswordLengthError"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORDVALUEERROR = "InvalidParameterValue.InvalidPasswordValueError"
//  INVALIDPARAMETERVALUE_PARAMETERCHARACTERERROR = "InvalidParameterValue.ParameterCharacterError"
//  INVALIDPARAMETERVALUE_PARAMETERLENGTHLIMITERROR = "InvalidParameterValue.ParameterLengthLimitError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  OPERATIONDENIED_KERNELVERSIONERROR = "OperationDenied.KernelVersionError"
func (c *Client) CreateAccount(request *CreateAccountRequest) (response *CreateAccountResponse, err error) {
    return c.CreateAccountWithContext(context.Background(), request)
}

// CreateAccount
// 此接口用于创建数据账号，返回的Oid为账号唯一标识。与数据库系统表pg_roles中记录的oid一致。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_PRECHECKERROR = "FailedOperation.PreCheckError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_DECRYPTPASSWORDFAILED = "InvalidParameterValue.DecryptPasswordFailed"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAMEERROR = "InvalidParameterValue.InvalidAccountNameError"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAMEFORMATERROR = "InvalidParameterValue.InvalidAccountNameFormatError"
//  INVALIDPARAMETERVALUE_INVALIDDATABASENAMEFORMATERROR = "InvalidParameterValue.InvalidDatabaseNameFormatError"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORDLENGTHERROR = "InvalidParameterValue.InvalidPasswordLengthError"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORDVALUEERROR = "InvalidParameterValue.InvalidPasswordValueError"
//  INVALIDPARAMETERVALUE_PARAMETERCHARACTERERROR = "InvalidParameterValue.ParameterCharacterError"
//  INVALIDPARAMETERVALUE_PARAMETERLENGTHLIMITERROR = "InvalidParameterValue.ParameterLengthLimitError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  OPERATIONDENIED_KERNELVERSIONERROR = "OperationDenied.KernelVersionError"
func (c *Client) CreateAccountWithContext(ctx context.Context, request *CreateAccountRequest) (response *CreateAccountResponse, err error) {
    if request == nil {
        request = NewCreateAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAccountResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBaseBackupRequest() (request *CreateBaseBackupRequest) {
    request = &CreateBaseBackupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "CreateBaseBackup")
    
    
    return
}

func NewCreateBaseBackupResponse() (response *CreateBaseBackupResponse) {
    response = &CreateBaseBackupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateBaseBackup
// 本接口（CreateBaseBackup）用于创建实例的数据备份。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) CreateBaseBackup(request *CreateBaseBackupRequest) (response *CreateBaseBackupResponse, err error) {
    return c.CreateBaseBackupWithContext(context.Background(), request)
}

// CreateBaseBackup
// 本接口（CreateBaseBackup）用于创建实例的数据备份。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) CreateBaseBackupWithContext(ctx context.Context, request *CreateBaseBackupRequest) (response *CreateBaseBackupResponse, err error) {
    if request == nil {
        request = NewCreateBaseBackupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBaseBackup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBaseBackupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDBInstanceNetworkAccessRequest() (request *CreateDBInstanceNetworkAccessRequest) {
    request = &CreateDBInstanceNetworkAccessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "CreateDBInstanceNetworkAccess")
    
    
    return
}

func NewCreateDBInstanceNetworkAccessResponse() (response *CreateDBInstanceNetworkAccessResponse) {
    response = &CreateDBInstanceNetworkAccessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDBInstanceNetworkAccess
// 本接口（CreateDBInstanceNetworkAccess）用于创建实例网络。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_CREATEBASICNETWORKDENIEDERROR = "FailedOperation.CreateBasicNetworkDeniedError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_GETSUBNETERROR = "FailedOperation.GetSubnetError"
//  FAILEDOPERATION_GETVPCINFOERROR = "FailedOperation.GetVpcInfoError"
//  FAILEDOPERATION_NETWORKNUMLIMITERROR = "FailedOperation.NetworkNumLimitError"
//  FAILEDOPERATION_QUERYSPECERROR = "FailedOperation.QuerySpecError"
//  FAILEDOPERATION_SERVERLESSNOTSUPPORTEDERROR = "FailedOperation.ServerlessNotSupportedError"
//  FLOWERROR = "FlowError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_MARSHALERROR = "InternalError.MarshalError"
//  INVALIDPARAMETER_VPCNOTFOUNDERROR = "InvalidParameter.VpcNotFoundError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCEIPV6NOTSUPPORTEDERROR = "OperationDenied.InstanceIpv6NotSupportedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) CreateDBInstanceNetworkAccess(request *CreateDBInstanceNetworkAccessRequest) (response *CreateDBInstanceNetworkAccessResponse, err error) {
    return c.CreateDBInstanceNetworkAccessWithContext(context.Background(), request)
}

// CreateDBInstanceNetworkAccess
// 本接口（CreateDBInstanceNetworkAccess）用于创建实例网络。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_CREATEBASICNETWORKDENIEDERROR = "FailedOperation.CreateBasicNetworkDeniedError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_GETSUBNETERROR = "FailedOperation.GetSubnetError"
//  FAILEDOPERATION_GETVPCINFOERROR = "FailedOperation.GetVpcInfoError"
//  FAILEDOPERATION_NETWORKNUMLIMITERROR = "FailedOperation.NetworkNumLimitError"
//  FAILEDOPERATION_QUERYSPECERROR = "FailedOperation.QuerySpecError"
//  FAILEDOPERATION_SERVERLESSNOTSUPPORTEDERROR = "FailedOperation.ServerlessNotSupportedError"
//  FLOWERROR = "FlowError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_MARSHALERROR = "InternalError.MarshalError"
//  INVALIDPARAMETER_VPCNOTFOUNDERROR = "InvalidParameter.VpcNotFoundError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCEIPV6NOTSUPPORTEDERROR = "OperationDenied.InstanceIpv6NotSupportedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) CreateDBInstanceNetworkAccessWithContext(ctx context.Context, request *CreateDBInstanceNetworkAccessRequest) (response *CreateDBInstanceNetworkAccessResponse, err error) {
    if request == nil {
        request = NewCreateDBInstanceNetworkAccessRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDBInstanceNetworkAccess require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDBInstanceNetworkAccessResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDBInstancesRequest() (request *CreateDBInstancesRequest) {
    request = &CreateDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "CreateDBInstances")
    
    
    return
}

func NewCreateDBInstancesResponse() (response *CreateDBInstancesResponse) {
    response = &CreateDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDBInstances
// 本接口（CreateDBInstances）用于创建一个或者多个PostgreSQL实例,仅发货实例不会进行初始化。本接口已废弃，推荐使用接口[CreateInstances](https://cloud.tencent.com/document/api/409/56107)替代。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_ALLOCATEQUOTASERROR = "FailedOperation.AllocateQuotasError"
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_CDBCGWCONNECTERROR = "FailedOperation.CdbCgwConnectError"
//  FAILEDOPERATION_CREATEBASICNETWORKDENIEDERROR = "FailedOperation.CreateBasicNetworkDeniedError"
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_GETSUBNETERROR = "FailedOperation.GetSubnetError"
//  FAILEDOPERATION_GETVPCINFOERROR = "FailedOperation.GetVpcInfoError"
//  FAILEDOPERATION_INVALIDTRADEOPERATE = "FailedOperation.InvalidTradeOperate"
//  FAILEDOPERATION_PAYORDERFAILED = "FailedOperation.PayOrderFailed"
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  FAILEDOPERATION_QUERYSPECERROR = "FailedOperation.QuerySpecError"
//  FAILEDOPERATION_QUERYVPCFAILED = "FailedOperation.QueryVpcFailed"
//  FAILEDOPERATION_QUERYVPCFALIED = "FailedOperation.QueryVpcFalied"
//  FAILEDOPERATION_SERVICEACCESSERROR = "FailedOperation.ServiceAccessError"
//  FAILEDOPERATION_STORAGEMEMORYCHECKERROR = "FailedOperation.StorageMemoryCheckError"
//  FAILEDOPERATION_TRADECREATEERROR = "FailedOperation.TradeCreateError"
//  FAILEDOPERATION_TRADEPAYORDERSERROR = "FailedOperation.TradePayOrdersError"
//  INTERNALERROR_CGWERROR = "InternalError.CgwError"
//  INTERNALERROR_CNSERROR = "InternalError.CnsError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DFWERROR = "InternalError.DfwError"
//  INTERNALERROR_FLOWERROR = "InternalError.FlowError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETER_TRADEACCESSDENIEDERROR = "InvalidParameter.TradeAccessDeniedError"
//  INVALIDPARAMETER_VPCNOTFOUNDERROR = "InvalidParameter.VpcNotFoundError"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_BADSPEC = "InvalidParameterValue.BadSpec"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCECHARGETYPE = "InvalidParameterValue.IllegalInstanceChargeType"
//  INVALIDPARAMETERVALUE_ILLEGALPROJECTID = "InvalidParameterValue.IllegalProjectId"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_INTERFACENAMENOTFOUND = "InvalidParameterValue.InterfaceNameNotFound"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTFORMAT = "InvalidParameterValue.InvalidAccountFormat"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAME = "InvalidParameterValue.InvalidAccountName"
//  INVALIDPARAMETERVALUE_INVALIDCHARSET = "InvalidParameterValue.InvalidCharset"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCENUM = "InvalidParameterValue.InvalidInstanceNum"
//  INVALIDPARAMETERVALUE_INVALIDORDERNUM = "InvalidParameterValue.InvalidOrderNum"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORDFORMAT = "InvalidParameterValue.InvalidPasswordFormat"
//  INVALIDPARAMETERVALUE_INVALIDPID = "InvalidParameterValue.InvalidPid"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  INVALIDPARAMETERVALUE_PARAMETERCHARACTERERROR = "InvalidParameterValue.ParameterCharacterError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_PARAMETERLENGTHLIMITERROR = "InvalidParameterValue.ParameterLengthLimitError"
//  INVALIDPARAMETERVALUE_PARAMETEROUTRANGEERROR = "InvalidParameterValue.ParameterOutRangeError"
//  INVALIDPARAMETERVALUE_REGIONNOTSUPPORTED = "InvalidParameterValue.RegionNotSupported"
//  INVALIDPARAMETERVALUE_SPECNOTRECOGNIZEDERROR = "InvalidParameterValue.SpecNotRecognizedError"
//  INVALIDPARAMETERVALUE_STRUCTPARSEFAILED = "InvalidParameterValue.StructParseFailed"
//  INVALIDPID = "InvalidPid"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  OPERATIONDENIED_USERNOTAUTHENTICATEDERROR = "OperationDenied.UserNotAuthenticatedError"
//  OPERATIONDENIED_VPCDENIEDERROR = "OperationDenied.VpcDeniedError"
//  RESOURCEINSUFFICIENT_RESOURCENOTENOUGH = "ResourceInsufficient.ResourceNotEnough"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
//  RESOURCEUNAVAILABLE_RESOURCENOPERMISSION = "ResourceUnavailable.ResourceNoPermission"
//  RESOURCEUNAVAILABLE_VPCRESOURCENOTFOUND = "ResourceUnavailable.VpcResourceNotFound"
//  UNAUTHORIZEDOPERATION_USERHASNOREALNAMEAUTHENTICATION = "UnauthorizedOperation.UserHasNoRealnameAuthentication"
//  UNKNOWNERROR = "UnknownError"
func (c *Client) CreateDBInstances(request *CreateDBInstancesRequest) (response *CreateDBInstancesResponse, err error) {
    return c.CreateDBInstancesWithContext(context.Background(), request)
}

// CreateDBInstances
// 本接口（CreateDBInstances）用于创建一个或者多个PostgreSQL实例,仅发货实例不会进行初始化。本接口已废弃，推荐使用接口[CreateInstances](https://cloud.tencent.com/document/api/409/56107)替代。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_ALLOCATEQUOTASERROR = "FailedOperation.AllocateQuotasError"
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_CDBCGWCONNECTERROR = "FailedOperation.CdbCgwConnectError"
//  FAILEDOPERATION_CREATEBASICNETWORKDENIEDERROR = "FailedOperation.CreateBasicNetworkDeniedError"
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_GETSUBNETERROR = "FailedOperation.GetSubnetError"
//  FAILEDOPERATION_GETVPCINFOERROR = "FailedOperation.GetVpcInfoError"
//  FAILEDOPERATION_INVALIDTRADEOPERATE = "FailedOperation.InvalidTradeOperate"
//  FAILEDOPERATION_PAYORDERFAILED = "FailedOperation.PayOrderFailed"
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  FAILEDOPERATION_QUERYSPECERROR = "FailedOperation.QuerySpecError"
//  FAILEDOPERATION_QUERYVPCFAILED = "FailedOperation.QueryVpcFailed"
//  FAILEDOPERATION_QUERYVPCFALIED = "FailedOperation.QueryVpcFalied"
//  FAILEDOPERATION_SERVICEACCESSERROR = "FailedOperation.ServiceAccessError"
//  FAILEDOPERATION_STORAGEMEMORYCHECKERROR = "FailedOperation.StorageMemoryCheckError"
//  FAILEDOPERATION_TRADECREATEERROR = "FailedOperation.TradeCreateError"
//  FAILEDOPERATION_TRADEPAYORDERSERROR = "FailedOperation.TradePayOrdersError"
//  INTERNALERROR_CGWERROR = "InternalError.CgwError"
//  INTERNALERROR_CNSERROR = "InternalError.CnsError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DFWERROR = "InternalError.DfwError"
//  INTERNALERROR_FLOWERROR = "InternalError.FlowError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETER_TRADEACCESSDENIEDERROR = "InvalidParameter.TradeAccessDeniedError"
//  INVALIDPARAMETER_VPCNOTFOUNDERROR = "InvalidParameter.VpcNotFoundError"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_BADSPEC = "InvalidParameterValue.BadSpec"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCECHARGETYPE = "InvalidParameterValue.IllegalInstanceChargeType"
//  INVALIDPARAMETERVALUE_ILLEGALPROJECTID = "InvalidParameterValue.IllegalProjectId"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_INTERFACENAMENOTFOUND = "InvalidParameterValue.InterfaceNameNotFound"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTFORMAT = "InvalidParameterValue.InvalidAccountFormat"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAME = "InvalidParameterValue.InvalidAccountName"
//  INVALIDPARAMETERVALUE_INVALIDCHARSET = "InvalidParameterValue.InvalidCharset"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCENUM = "InvalidParameterValue.InvalidInstanceNum"
//  INVALIDPARAMETERVALUE_INVALIDORDERNUM = "InvalidParameterValue.InvalidOrderNum"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORDFORMAT = "InvalidParameterValue.InvalidPasswordFormat"
//  INVALIDPARAMETERVALUE_INVALIDPID = "InvalidParameterValue.InvalidPid"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  INVALIDPARAMETERVALUE_PARAMETERCHARACTERERROR = "InvalidParameterValue.ParameterCharacterError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_PARAMETERLENGTHLIMITERROR = "InvalidParameterValue.ParameterLengthLimitError"
//  INVALIDPARAMETERVALUE_PARAMETEROUTRANGEERROR = "InvalidParameterValue.ParameterOutRangeError"
//  INVALIDPARAMETERVALUE_REGIONNOTSUPPORTED = "InvalidParameterValue.RegionNotSupported"
//  INVALIDPARAMETERVALUE_SPECNOTRECOGNIZEDERROR = "InvalidParameterValue.SpecNotRecognizedError"
//  INVALIDPARAMETERVALUE_STRUCTPARSEFAILED = "InvalidParameterValue.StructParseFailed"
//  INVALIDPID = "InvalidPid"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  OPERATIONDENIED_USERNOTAUTHENTICATEDERROR = "OperationDenied.UserNotAuthenticatedError"
//  OPERATIONDENIED_VPCDENIEDERROR = "OperationDenied.VpcDeniedError"
//  RESOURCEINSUFFICIENT_RESOURCENOTENOUGH = "ResourceInsufficient.ResourceNotEnough"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
//  RESOURCEUNAVAILABLE_RESOURCENOPERMISSION = "ResourceUnavailable.ResourceNoPermission"
//  RESOURCEUNAVAILABLE_VPCRESOURCENOTFOUND = "ResourceUnavailable.VpcResourceNotFound"
//  UNAUTHORIZEDOPERATION_USERHASNOREALNAMEAUTHENTICATION = "UnauthorizedOperation.UserHasNoRealnameAuthentication"
//  UNKNOWNERROR = "UnknownError"
func (c *Client) CreateDBInstancesWithContext(ctx context.Context, request *CreateDBInstancesRequest) (response *CreateDBInstancesResponse, err error) {
    if request == nil {
        request = NewCreateDBInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstancesRequest() (request *CreateInstancesRequest) {
    request = &CreateInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "CreateInstances")
    
    
    return
}

func NewCreateInstancesResponse() (response *CreateInstancesResponse) {
    response = &CreateInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateInstances
// 本接口 (CreateInstances) 用于创建一个或者多个PostgreSQL实例，通过此接口创建的实例无需进行初始化，可直接使用。
//
// <li>实例创建成功后将自动开机启动，实例状态变为“运行中”。</li>
//
// <li>预付费实例的购买会预先扣除本次实例购买所需金额，按小时后付费实例购买会预先冻结本次实例购买一小时内所需金额，在调用本接口前请确保账户余额充足。</li>
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_ALLOCATEQUOTASERROR = "FailedOperation.AllocateQuotasError"
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_CDBCGWCONNECTERROR = "FailedOperation.CdbCgwConnectError"
//  FAILEDOPERATION_CREATEBASICNETWORKDENIEDERROR = "FailedOperation.CreateBasicNetworkDeniedError"
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_GETSUBNETERROR = "FailedOperation.GetSubnetError"
//  FAILEDOPERATION_GETVPCINFOERROR = "FailedOperation.GetVpcInfoError"
//  FAILEDOPERATION_INVALIDTRADEOPERATE = "FailedOperation.InvalidTradeOperate"
//  FAILEDOPERATION_PAYORDERFAILED = "FailedOperation.PayOrderFailed"
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  FAILEDOPERATION_QUERYSPECERROR = "FailedOperation.QuerySpecError"
//  FAILEDOPERATION_QUERYVPCFAILED = "FailedOperation.QueryVpcFailed"
//  FAILEDOPERATION_QUERYVPCFALIED = "FailedOperation.QueryVpcFalied"
//  FAILEDOPERATION_SERVICEACCESSERROR = "FailedOperation.ServiceAccessError"
//  FAILEDOPERATION_STORAGEMEMORYCHECKERROR = "FailedOperation.StorageMemoryCheckError"
//  FAILEDOPERATION_TRADECREATEERROR = "FailedOperation.TradeCreateError"
//  FAILEDOPERATION_TRADEPAYORDERSERROR = "FailedOperation.TradePayOrdersError"
//  FAILEDOPERATION_WHITELISTCONNECTERROR = "FailedOperation.WhitelistConnectError"
//  INTERNALERROR_CGWERROR = "InternalError.CgwError"
//  INTERNALERROR_CNSERROR = "InternalError.CnsError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DFWERROR = "InternalError.DfwError"
//  INTERNALERROR_FLOWERROR = "InternalError.FlowError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_JSONPARSEERROR = "InternalError.JsonParseError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETER_TRADEACCESSDENIEDERROR = "InvalidParameter.TradeAccessDeniedError"
//  INVALIDPARAMETER_VPCNOTFOUNDERROR = "InvalidParameter.VpcNotFoundError"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_BADSPEC = "InvalidParameterValue.BadSpec"
//  INVALIDPARAMETERVALUE_CHARSETNOTFOUNDERROR = "InvalidParameterValue.CharsetNotFoundError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_DECRYPTPASSWORDFAILED = "InvalidParameterValue.DecryptPasswordFailed"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCECHARGETYPE = "InvalidParameterValue.IllegalInstanceChargeType"
//  INVALIDPARAMETERVALUE_ILLEGALPROJECTID = "InvalidParameterValue.IllegalProjectId"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_INTERFACENAMENOTFOUND = "InvalidParameterValue.InterfaceNameNotFound"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTERROR = "InvalidParameterValue.InvalidAccountError"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTFORMAT = "InvalidParameterValue.InvalidAccountFormat"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAME = "InvalidParameterValue.InvalidAccountName"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAMEERROR = "InvalidParameterValue.InvalidAccountNameError"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAMEFORMATERROR = "InvalidParameterValue.InvalidAccountNameFormatError"
//  INVALIDPARAMETERVALUE_INVALIDCHARSET = "InvalidParameterValue.InvalidCharset"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCENUM = "InvalidParameterValue.InvalidInstanceNum"
//  INVALIDPARAMETERVALUE_INVALIDORDERNUM = "InvalidParameterValue.InvalidOrderNum"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORDFORMAT = "InvalidParameterValue.InvalidPasswordFormat"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORDLENGTHERROR = "InvalidParameterValue.InvalidPasswordLengthError"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORDVALUEERROR = "InvalidParameterValue.InvalidPasswordValueError"
//  INVALIDPARAMETERVALUE_INVALIDPID = "InvalidParameterValue.InvalidPid"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  INVALIDPARAMETERVALUE_PARAMETERCHARACTERERROR = "InvalidParameterValue.ParameterCharacterError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_PARAMETERLENGTHLIMITERROR = "InvalidParameterValue.ParameterLengthLimitError"
//  INVALIDPARAMETERVALUE_PARAMETEROUTRANGEERROR = "InvalidParameterValue.ParameterOutRangeError"
//  INVALIDPARAMETERVALUE_REGIONNOTSUPPORTED = "InvalidParameterValue.RegionNotSupported"
//  INVALIDPARAMETERVALUE_SPECNOTRECOGNIZEDERROR = "InvalidParameterValue.SpecNotRecognizedError"
//  INVALIDPARAMETERVALUE_STRUCTPARSEFAILED = "InvalidParameterValue.StructParseFailed"
//  INVALIDPID = "InvalidPid"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  OPERATIONDENIED_USERNOTAUTHENTICATEDERROR = "OperationDenied.UserNotAuthenticatedError"
//  OPERATIONDENIED_VPCDENIEDERROR = "OperationDenied.VpcDeniedError"
//  RESOURCEINSUFFICIENT_RESOURCENOTENOUGH = "ResourceInsufficient.ResourceNotEnough"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
//  RESOURCEUNAVAILABLE_RESOURCENOPERMISSION = "ResourceUnavailable.ResourceNoPermission"
//  RESOURCEUNAVAILABLE_VPCRESOURCENOTFOUND = "ResourceUnavailable.VpcResourceNotFound"
//  UNAUTHORIZEDOPERATION_USERHASNOREALNAMEAUTHENTICATION = "UnauthorizedOperation.UserHasNoRealnameAuthentication"
//  UNKNOWNERROR = "UnknownError"
func (c *Client) CreateInstances(request *CreateInstancesRequest) (response *CreateInstancesResponse, err error) {
    return c.CreateInstancesWithContext(context.Background(), request)
}

// CreateInstances
// 本接口 (CreateInstances) 用于创建一个或者多个PostgreSQL实例，通过此接口创建的实例无需进行初始化，可直接使用。
//
// <li>实例创建成功后将自动开机启动，实例状态变为“运行中”。</li>
//
// <li>预付费实例的购买会预先扣除本次实例购买所需金额，按小时后付费实例购买会预先冻结本次实例购买一小时内所需金额，在调用本接口前请确保账户余额充足。</li>
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_ALLOCATEQUOTASERROR = "FailedOperation.AllocateQuotasError"
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_CDBCGWCONNECTERROR = "FailedOperation.CdbCgwConnectError"
//  FAILEDOPERATION_CREATEBASICNETWORKDENIEDERROR = "FailedOperation.CreateBasicNetworkDeniedError"
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_GETSUBNETERROR = "FailedOperation.GetSubnetError"
//  FAILEDOPERATION_GETVPCINFOERROR = "FailedOperation.GetVpcInfoError"
//  FAILEDOPERATION_INVALIDTRADEOPERATE = "FailedOperation.InvalidTradeOperate"
//  FAILEDOPERATION_PAYORDERFAILED = "FailedOperation.PayOrderFailed"
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  FAILEDOPERATION_QUERYSPECERROR = "FailedOperation.QuerySpecError"
//  FAILEDOPERATION_QUERYVPCFAILED = "FailedOperation.QueryVpcFailed"
//  FAILEDOPERATION_QUERYVPCFALIED = "FailedOperation.QueryVpcFalied"
//  FAILEDOPERATION_SERVICEACCESSERROR = "FailedOperation.ServiceAccessError"
//  FAILEDOPERATION_STORAGEMEMORYCHECKERROR = "FailedOperation.StorageMemoryCheckError"
//  FAILEDOPERATION_TRADECREATEERROR = "FailedOperation.TradeCreateError"
//  FAILEDOPERATION_TRADEPAYORDERSERROR = "FailedOperation.TradePayOrdersError"
//  FAILEDOPERATION_WHITELISTCONNECTERROR = "FailedOperation.WhitelistConnectError"
//  INTERNALERROR_CGWERROR = "InternalError.CgwError"
//  INTERNALERROR_CNSERROR = "InternalError.CnsError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DFWERROR = "InternalError.DfwError"
//  INTERNALERROR_FLOWERROR = "InternalError.FlowError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_JSONPARSEERROR = "InternalError.JsonParseError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETER_TRADEACCESSDENIEDERROR = "InvalidParameter.TradeAccessDeniedError"
//  INVALIDPARAMETER_VPCNOTFOUNDERROR = "InvalidParameter.VpcNotFoundError"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_BADSPEC = "InvalidParameterValue.BadSpec"
//  INVALIDPARAMETERVALUE_CHARSETNOTFOUNDERROR = "InvalidParameterValue.CharsetNotFoundError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_DECRYPTPASSWORDFAILED = "InvalidParameterValue.DecryptPasswordFailed"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCECHARGETYPE = "InvalidParameterValue.IllegalInstanceChargeType"
//  INVALIDPARAMETERVALUE_ILLEGALPROJECTID = "InvalidParameterValue.IllegalProjectId"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_INTERFACENAMENOTFOUND = "InvalidParameterValue.InterfaceNameNotFound"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTERROR = "InvalidParameterValue.InvalidAccountError"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTFORMAT = "InvalidParameterValue.InvalidAccountFormat"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAME = "InvalidParameterValue.InvalidAccountName"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAMEERROR = "InvalidParameterValue.InvalidAccountNameError"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAMEFORMATERROR = "InvalidParameterValue.InvalidAccountNameFormatError"
//  INVALIDPARAMETERVALUE_INVALIDCHARSET = "InvalidParameterValue.InvalidCharset"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCENUM = "InvalidParameterValue.InvalidInstanceNum"
//  INVALIDPARAMETERVALUE_INVALIDORDERNUM = "InvalidParameterValue.InvalidOrderNum"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORDFORMAT = "InvalidParameterValue.InvalidPasswordFormat"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORDLENGTHERROR = "InvalidParameterValue.InvalidPasswordLengthError"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORDVALUEERROR = "InvalidParameterValue.InvalidPasswordValueError"
//  INVALIDPARAMETERVALUE_INVALIDPID = "InvalidParameterValue.InvalidPid"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  INVALIDPARAMETERVALUE_PARAMETERCHARACTERERROR = "InvalidParameterValue.ParameterCharacterError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_PARAMETERLENGTHLIMITERROR = "InvalidParameterValue.ParameterLengthLimitError"
//  INVALIDPARAMETERVALUE_PARAMETEROUTRANGEERROR = "InvalidParameterValue.ParameterOutRangeError"
//  INVALIDPARAMETERVALUE_REGIONNOTSUPPORTED = "InvalidParameterValue.RegionNotSupported"
//  INVALIDPARAMETERVALUE_SPECNOTRECOGNIZEDERROR = "InvalidParameterValue.SpecNotRecognizedError"
//  INVALIDPARAMETERVALUE_STRUCTPARSEFAILED = "InvalidParameterValue.StructParseFailed"
//  INVALIDPID = "InvalidPid"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  OPERATIONDENIED_USERNOTAUTHENTICATEDERROR = "OperationDenied.UserNotAuthenticatedError"
//  OPERATIONDENIED_VPCDENIEDERROR = "OperationDenied.VpcDeniedError"
//  RESOURCEINSUFFICIENT_RESOURCENOTENOUGH = "ResourceInsufficient.ResourceNotEnough"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
//  RESOURCEUNAVAILABLE_RESOURCENOPERMISSION = "ResourceUnavailable.ResourceNoPermission"
//  RESOURCEUNAVAILABLE_VPCRESOURCENOTFOUND = "ResourceUnavailable.VpcResourceNotFound"
//  UNAUTHORIZEDOPERATION_USERHASNOREALNAMEAUTHENTICATION = "UnauthorizedOperation.UserHasNoRealnameAuthentication"
//  UNKNOWNERROR = "UnknownError"
func (c *Client) CreateInstancesWithContext(ctx context.Context, request *CreateInstancesRequest) (response *CreateInstancesResponse, err error) {
    if request == nil {
        request = NewCreateInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateParameterTemplateRequest() (request *CreateParameterTemplateRequest) {
    request = &CreateParameterTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "CreateParameterTemplate")
    
    
    return
}

func NewCreateParameterTemplateResponse() (response *CreateParameterTemplateResponse) {
    response = &CreateParameterTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateParameterTemplate
// 本接口 (CreateParameterTemplate) 用于创建参数模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
func (c *Client) CreateParameterTemplate(request *CreateParameterTemplateRequest) (response *CreateParameterTemplateResponse, err error) {
    return c.CreateParameterTemplateWithContext(context.Background(), request)
}

// CreateParameterTemplate
// 本接口 (CreateParameterTemplate) 用于创建参数模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
func (c *Client) CreateParameterTemplateWithContext(ctx context.Context, request *CreateParameterTemplateRequest) (response *CreateParameterTemplateResponse, err error) {
    if request == nil {
        request = NewCreateParameterTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateParameterTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateParameterTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateReadOnlyDBInstanceRequest() (request *CreateReadOnlyDBInstanceRequest) {
    request = &CreateReadOnlyDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "CreateReadOnlyDBInstance")
    
    
    return
}

func NewCreateReadOnlyDBInstanceResponse() (response *CreateReadOnlyDBInstanceResponse) {
    response = &CreateReadOnlyDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateReadOnlyDBInstance
// 本接口(CreateReadOnlyDBInstance)用于创建只读实例
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_ALLOCATEQUOTASERROR = "FailedOperation.AllocateQuotasError"
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_CDBCGWCONNECTERROR = "FailedOperation.CdbCgwConnectError"
//  FAILEDOPERATION_CREATEBASICNETWORKDENIEDERROR = "FailedOperation.CreateBasicNetworkDeniedError"
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_GETSUBNETERROR = "FailedOperation.GetSubnetError"
//  FAILEDOPERATION_GETVPCINFOERROR = "FailedOperation.GetVpcInfoError"
//  FAILEDOPERATION_INVALIDTRADEOPERATE = "FailedOperation.InvalidTradeOperate"
//  FAILEDOPERATION_PAYORDERFAILED = "FailedOperation.PayOrderFailed"
//  FAILEDOPERATION_PRECHECKERROR = "FailedOperation.PreCheckError"
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  FAILEDOPERATION_QUERYSPECERROR = "FailedOperation.QuerySpecError"
//  FAILEDOPERATION_QUERYVPCFAILED = "FailedOperation.QueryVpcFailed"
//  FAILEDOPERATION_QUERYVPCFALIED = "FailedOperation.QueryVpcFalied"
//  FAILEDOPERATION_ROGROUPNOTFOUNDERROR = "FailedOperation.ROGroupNotFoundError"
//  FAILEDOPERATION_ROGROUPNUMEXCEEDERROR = "FailedOperation.ROGroupNumExceedError"
//  FAILEDOPERATION_SERVICEACCESSERROR = "FailedOperation.ServiceAccessError"
//  FAILEDOPERATION_STORAGEMEMORYCHECKERROR = "FailedOperation.StorageMemoryCheckError"
//  FAILEDOPERATION_TRADECREATEERROR = "FailedOperation.TradeCreateError"
//  FAILEDOPERATION_TRADEPAYORDERSERROR = "FailedOperation.TradePayOrdersError"
//  FAILEDOPERATION_WHITELISTCONNECTERROR = "FailedOperation.WhitelistConnectError"
//  INTERNALERROR_CGWERROR = "InternalError.CgwError"
//  INTERNALERROR_CNSERROR = "InternalError.CnsError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DFWERROR = "InternalError.DfwError"
//  INTERNALERROR_FLOWERROR = "InternalError.FlowError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETER_TRADEACCESSDENIEDERROR = "InvalidParameter.TradeAccessDeniedError"
//  INVALIDPARAMETER_VPCNOTFOUNDERROR = "InvalidParameter.VpcNotFoundError"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_BADSPEC = "InvalidParameterValue.BadSpec"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCECHARGETYPE = "InvalidParameterValue.IllegalInstanceChargeType"
//  INVALIDPARAMETERVALUE_ILLEGALPROJECTID = "InvalidParameterValue.IllegalProjectId"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_INTERFACENAMENOTFOUND = "InvalidParameterValue.InterfaceNameNotFound"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTFORMAT = "InvalidParameterValue.InvalidAccountFormat"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAME = "InvalidParameterValue.InvalidAccountName"
//  INVALIDPARAMETERVALUE_INVALIDCHARSET = "InvalidParameterValue.InvalidCharset"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCENUM = "InvalidParameterValue.InvalidInstanceNum"
//  INVALIDPARAMETERVALUE_INVALIDORDERNUM = "InvalidParameterValue.InvalidOrderNum"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORDFORMAT = "InvalidParameterValue.InvalidPasswordFormat"
//  INVALIDPARAMETERVALUE_INVALIDPID = "InvalidParameterValue.InvalidPid"
//  INVALIDPARAMETERVALUE_PARAMETERCHARACTERERROR = "InvalidParameterValue.ParameterCharacterError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_PARAMETERLENGTHLIMITERROR = "InvalidParameterValue.ParameterLengthLimitError"
//  INVALIDPARAMETERVALUE_REGIONNOTSUPPORTED = "InvalidParameterValue.RegionNotSupported"
//  INVALIDPARAMETERVALUE_STRUCTPARSEFAILED = "InvalidParameterValue.StructParseFailed"
//  INVALIDPID = "InvalidPid"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  OPERATIONDENIED_PAYMODEERROR = "OperationDenied.PayModeError"
//  OPERATIONDENIED_ROGROUPSTATUSERROR = "OperationDenied.ROGroupStatusError"
//  OPERATIONDENIED_ROINSTANCECOUNTEXEEDERROR = "OperationDenied.RoInstanceCountExeedError"
//  OPERATIONDENIED_USERNOTAUTHENTICATEDERROR = "OperationDenied.UserNotAuthenticatedError"
//  OPERATIONDENIED_VPCDENIEDERROR = "OperationDenied.VpcDeniedError"
//  RESOURCEINSUFFICIENT_RESOURCENOTENOUGH = "ResourceInsufficient.ResourceNotEnough"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
//  RESOURCEUNAVAILABLE_RESOURCENOPERMISSION = "ResourceUnavailable.ResourceNoPermission"
//  RESOURCEUNAVAILABLE_VPCRESOURCENOTFOUND = "ResourceUnavailable.VpcResourceNotFound"
//  UNAUTHORIZEDOPERATION_USERHASNOREALNAMEAUTHENTICATION = "UnauthorizedOperation.UserHasNoRealnameAuthentication"
//  UNKNOWNERROR = "UnknownError"
func (c *Client) CreateReadOnlyDBInstance(request *CreateReadOnlyDBInstanceRequest) (response *CreateReadOnlyDBInstanceResponse, err error) {
    return c.CreateReadOnlyDBInstanceWithContext(context.Background(), request)
}

// CreateReadOnlyDBInstance
// 本接口(CreateReadOnlyDBInstance)用于创建只读实例
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_ALLOCATEQUOTASERROR = "FailedOperation.AllocateQuotasError"
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_CDBCGWCONNECTERROR = "FailedOperation.CdbCgwConnectError"
//  FAILEDOPERATION_CREATEBASICNETWORKDENIEDERROR = "FailedOperation.CreateBasicNetworkDeniedError"
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_GETSUBNETERROR = "FailedOperation.GetSubnetError"
//  FAILEDOPERATION_GETVPCINFOERROR = "FailedOperation.GetVpcInfoError"
//  FAILEDOPERATION_INVALIDTRADEOPERATE = "FailedOperation.InvalidTradeOperate"
//  FAILEDOPERATION_PAYORDERFAILED = "FailedOperation.PayOrderFailed"
//  FAILEDOPERATION_PRECHECKERROR = "FailedOperation.PreCheckError"
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  FAILEDOPERATION_QUERYSPECERROR = "FailedOperation.QuerySpecError"
//  FAILEDOPERATION_QUERYVPCFAILED = "FailedOperation.QueryVpcFailed"
//  FAILEDOPERATION_QUERYVPCFALIED = "FailedOperation.QueryVpcFalied"
//  FAILEDOPERATION_ROGROUPNOTFOUNDERROR = "FailedOperation.ROGroupNotFoundError"
//  FAILEDOPERATION_ROGROUPNUMEXCEEDERROR = "FailedOperation.ROGroupNumExceedError"
//  FAILEDOPERATION_SERVICEACCESSERROR = "FailedOperation.ServiceAccessError"
//  FAILEDOPERATION_STORAGEMEMORYCHECKERROR = "FailedOperation.StorageMemoryCheckError"
//  FAILEDOPERATION_TRADECREATEERROR = "FailedOperation.TradeCreateError"
//  FAILEDOPERATION_TRADEPAYORDERSERROR = "FailedOperation.TradePayOrdersError"
//  FAILEDOPERATION_WHITELISTCONNECTERROR = "FailedOperation.WhitelistConnectError"
//  INTERNALERROR_CGWERROR = "InternalError.CgwError"
//  INTERNALERROR_CNSERROR = "InternalError.CnsError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DFWERROR = "InternalError.DfwError"
//  INTERNALERROR_FLOWERROR = "InternalError.FlowError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETER_TRADEACCESSDENIEDERROR = "InvalidParameter.TradeAccessDeniedError"
//  INVALIDPARAMETER_VPCNOTFOUNDERROR = "InvalidParameter.VpcNotFoundError"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_BADSPEC = "InvalidParameterValue.BadSpec"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCECHARGETYPE = "InvalidParameterValue.IllegalInstanceChargeType"
//  INVALIDPARAMETERVALUE_ILLEGALPROJECTID = "InvalidParameterValue.IllegalProjectId"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_INTERFACENAMENOTFOUND = "InvalidParameterValue.InterfaceNameNotFound"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTFORMAT = "InvalidParameterValue.InvalidAccountFormat"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAME = "InvalidParameterValue.InvalidAccountName"
//  INVALIDPARAMETERVALUE_INVALIDCHARSET = "InvalidParameterValue.InvalidCharset"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCENUM = "InvalidParameterValue.InvalidInstanceNum"
//  INVALIDPARAMETERVALUE_INVALIDORDERNUM = "InvalidParameterValue.InvalidOrderNum"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORDFORMAT = "InvalidParameterValue.InvalidPasswordFormat"
//  INVALIDPARAMETERVALUE_INVALIDPID = "InvalidParameterValue.InvalidPid"
//  INVALIDPARAMETERVALUE_PARAMETERCHARACTERERROR = "InvalidParameterValue.ParameterCharacterError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_PARAMETERLENGTHLIMITERROR = "InvalidParameterValue.ParameterLengthLimitError"
//  INVALIDPARAMETERVALUE_REGIONNOTSUPPORTED = "InvalidParameterValue.RegionNotSupported"
//  INVALIDPARAMETERVALUE_STRUCTPARSEFAILED = "InvalidParameterValue.StructParseFailed"
//  INVALIDPID = "InvalidPid"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  OPERATIONDENIED_PAYMODEERROR = "OperationDenied.PayModeError"
//  OPERATIONDENIED_ROGROUPSTATUSERROR = "OperationDenied.ROGroupStatusError"
//  OPERATIONDENIED_ROINSTANCECOUNTEXEEDERROR = "OperationDenied.RoInstanceCountExeedError"
//  OPERATIONDENIED_USERNOTAUTHENTICATEDERROR = "OperationDenied.UserNotAuthenticatedError"
//  OPERATIONDENIED_VPCDENIEDERROR = "OperationDenied.VpcDeniedError"
//  RESOURCEINSUFFICIENT_RESOURCENOTENOUGH = "ResourceInsufficient.ResourceNotEnough"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
//  RESOURCEUNAVAILABLE_RESOURCENOPERMISSION = "ResourceUnavailable.ResourceNoPermission"
//  RESOURCEUNAVAILABLE_VPCRESOURCENOTFOUND = "ResourceUnavailable.VpcResourceNotFound"
//  UNAUTHORIZEDOPERATION_USERHASNOREALNAMEAUTHENTICATION = "UnauthorizedOperation.UserHasNoRealnameAuthentication"
//  UNKNOWNERROR = "UnknownError"
func (c *Client) CreateReadOnlyDBInstanceWithContext(ctx context.Context, request *CreateReadOnlyDBInstanceRequest) (response *CreateReadOnlyDBInstanceResponse, err error) {
    if request == nil {
        request = NewCreateReadOnlyDBInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateReadOnlyDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateReadOnlyDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateReadOnlyGroupRequest() (request *CreateReadOnlyGroupRequest) {
    request = &CreateReadOnlyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "CreateReadOnlyGroup")
    
    
    return
}

func NewCreateReadOnlyGroupResponse() (response *CreateReadOnlyGroupResponse) {
    response = &CreateReadOnlyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateReadOnlyGroup
// 本接口（CreateReadOnlyGroup）用于创建只读组
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_ALLOCATEQUOTASERROR = "FailedOperation.AllocateQuotasError"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_CREATEBASICNETWORKDENIEDERROR = "FailedOperation.CreateBasicNetworkDeniedError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_GETVPCINFOERROR = "FailedOperation.GetVpcInfoError"
//  FAILEDOPERATION_ILLEGALROINSTANCENUM = "FailedOperation.IllegalROInstanceNum"
//  FAILEDOPERATION_MASTERINSTANCEQUERYERROR = "FailedOperation.MasterInstanceQueryError"
//  FAILEDOPERATION_ROGROUPNUMEXCEEDERROR = "FailedOperation.ROGroupNumExceedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETER_VPCNOTFOUNDERROR = "InvalidParameter.VpcNotFoundError"
//  INVALIDPARAMETERVALUE_INTERFACENAMENOTFOUND = "InvalidParameterValue.InterfaceNameNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERCHARACTERERROR = "InvalidParameterValue.ParameterCharacterError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_PARAMETERLENGTHLIMITERROR = "InvalidParameterValue.ParameterLengthLimitError"
//  INVALIDPARAMETERVALUE_REGIONNOTSUPPORTED = "InvalidParameterValue.RegionNotSupported"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  STRUCTPARSEFAILED = "StructParseFailed"
//  UNKNOWNERROR = "UnknownError"
func (c *Client) CreateReadOnlyGroup(request *CreateReadOnlyGroupRequest) (response *CreateReadOnlyGroupResponse, err error) {
    return c.CreateReadOnlyGroupWithContext(context.Background(), request)
}

// CreateReadOnlyGroup
// 本接口（CreateReadOnlyGroup）用于创建只读组
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_ALLOCATEQUOTASERROR = "FailedOperation.AllocateQuotasError"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_CREATEBASICNETWORKDENIEDERROR = "FailedOperation.CreateBasicNetworkDeniedError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_GETVPCINFOERROR = "FailedOperation.GetVpcInfoError"
//  FAILEDOPERATION_ILLEGALROINSTANCENUM = "FailedOperation.IllegalROInstanceNum"
//  FAILEDOPERATION_MASTERINSTANCEQUERYERROR = "FailedOperation.MasterInstanceQueryError"
//  FAILEDOPERATION_ROGROUPNUMEXCEEDERROR = "FailedOperation.ROGroupNumExceedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETER_VPCNOTFOUNDERROR = "InvalidParameter.VpcNotFoundError"
//  INVALIDPARAMETERVALUE_INTERFACENAMENOTFOUND = "InvalidParameterValue.InterfaceNameNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERCHARACTERERROR = "InvalidParameterValue.ParameterCharacterError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_PARAMETERLENGTHLIMITERROR = "InvalidParameterValue.ParameterLengthLimitError"
//  INVALIDPARAMETERVALUE_REGIONNOTSUPPORTED = "InvalidParameterValue.RegionNotSupported"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  STRUCTPARSEFAILED = "StructParseFailed"
//  UNKNOWNERROR = "UnknownError"
func (c *Client) CreateReadOnlyGroupWithContext(ctx context.Context, request *CreateReadOnlyGroupRequest) (response *CreateReadOnlyGroupResponse, err error) {
    if request == nil {
        request = NewCreateReadOnlyGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateReadOnlyGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateReadOnlyGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateReadOnlyGroupNetworkAccessRequest() (request *CreateReadOnlyGroupNetworkAccessRequest) {
    request = &CreateReadOnlyGroupNetworkAccessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "CreateReadOnlyGroupNetworkAccess")
    
    
    return
}

func NewCreateReadOnlyGroupNetworkAccessResponse() (response *CreateReadOnlyGroupNetworkAccessResponse) {
    response = &CreateReadOnlyGroupNetworkAccessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateReadOnlyGroupNetworkAccess
// 本接口（CreateReadOnlyGroupNetworkAccess）用于创建RO组的网络。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_CREATEBASICNETWORKDENIEDERROR = "FailedOperation.CreateBasicNetworkDeniedError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_GETSUBNETERROR = "FailedOperation.GetSubnetError"
//  FAILEDOPERATION_GETVPCINFOERROR = "FailedOperation.GetVpcInfoError"
//  FAILEDOPERATION_NETWORKNUMLIMITERROR = "FailedOperation.NetworkNumLimitError"
//  FAILEDOPERATION_ROGROUPNOTFOUNDERROR = "FailedOperation.ROGroupNotFoundError"
//  FLOWERROR = "FlowError"
//  INTERNALERROR_FLOWERROR = "InternalError.FlowError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_MARSHALERROR = "InternalError.MarshalError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_ROGROUPSTATUSERROR = "OperationDenied.ROGroupStatusError"
func (c *Client) CreateReadOnlyGroupNetworkAccess(request *CreateReadOnlyGroupNetworkAccessRequest) (response *CreateReadOnlyGroupNetworkAccessResponse, err error) {
    return c.CreateReadOnlyGroupNetworkAccessWithContext(context.Background(), request)
}

// CreateReadOnlyGroupNetworkAccess
// 本接口（CreateReadOnlyGroupNetworkAccess）用于创建RO组的网络。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_CREATEBASICNETWORKDENIEDERROR = "FailedOperation.CreateBasicNetworkDeniedError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_GETSUBNETERROR = "FailedOperation.GetSubnetError"
//  FAILEDOPERATION_GETVPCINFOERROR = "FailedOperation.GetVpcInfoError"
//  FAILEDOPERATION_NETWORKNUMLIMITERROR = "FailedOperation.NetworkNumLimitError"
//  FAILEDOPERATION_ROGROUPNOTFOUNDERROR = "FailedOperation.ROGroupNotFoundError"
//  FLOWERROR = "FlowError"
//  INTERNALERROR_FLOWERROR = "InternalError.FlowError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_MARSHALERROR = "InternalError.MarshalError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_ROGROUPSTATUSERROR = "OperationDenied.ROGroupStatusError"
func (c *Client) CreateReadOnlyGroupNetworkAccessWithContext(ctx context.Context, request *CreateReadOnlyGroupNetworkAccessRequest) (response *CreateReadOnlyGroupNetworkAccessResponse, err error) {
    if request == nil {
        request = NewCreateReadOnlyGroupNetworkAccessRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateReadOnlyGroupNetworkAccess require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateReadOnlyGroupNetworkAccessResponse()
    err = c.Send(request, response)
    return
}

func NewCreateServerlessDBInstanceRequest() (request *CreateServerlessDBInstanceRequest) {
    request = &CreateServerlessDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "CreateServerlessDBInstance")
    
    
    return
}

func NewCreateServerlessDBInstanceResponse() (response *CreateServerlessDBInstanceResponse) {
    response = &CreateServerlessDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateServerlessDBInstance
// 本接口 (CreateServerlessDBInstance) 用于创建一个ServerlessDB实例，创建成功返回实例ID。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_ALLOCATEQUOTASERROR = "FailedOperation.AllocateQuotasError"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_CREATEBASICNETWORKDENIEDERROR = "FailedOperation.CreateBasicNetworkDeniedError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_GETVPCINFOERROR = "FailedOperation.GetVpcInfoError"
//  FAILEDOPERATION_QUERYVPCFAILED = "FailedOperation.QueryVpcFailed"
//  FLOWERROR = "FlowError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_FLOWERROR = "InternalError.FlowError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENAMEEXIST = "InvalidParameter.InstanceNameExist"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_INSTANCENAMEEXIST = "InvalidParameterValue.InstanceNameExist"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTFORMAT = "InvalidParameterValue.InvalidAccountFormat"
//  INVALIDPARAMETERVALUE_INVALIDCHARSET = "InvalidParameterValue.InvalidCharset"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORDFORMAT = "InvalidParameterValue.InvalidPasswordFormat"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  INVALIDPARAMETERVALUE_PARAMETERCHARACTERERROR = "InvalidParameterValue.ParameterCharacterError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_PARAMETERLENGTHLIMITERROR = "InvalidParameterValue.ParameterLengthLimitError"
//  INVALIDPARAMETERVALUE_REGIONNOTSUPPORTED = "InvalidParameterValue.RegionNotSupported"
//  INVALIDPARAMETERVALUE_STRUCTPARSEFAILED = "InvalidParameterValue.StructParseFailed"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_NOTSUPPORTZONEERROR = "OperationDenied.NotSupportZoneError"
//  OPERATIONDENIED_VERSIONNOTSUPPORTERROR = "OperationDenied.VersionNotSupportError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
//  UNAUTHORIZEDOPERATION_USERHASNOREALNAMEAUTHENTICATION = "UnauthorizedOperation.UserHasNoRealnameAuthentication"
func (c *Client) CreateServerlessDBInstance(request *CreateServerlessDBInstanceRequest) (response *CreateServerlessDBInstanceResponse, err error) {
    return c.CreateServerlessDBInstanceWithContext(context.Background(), request)
}

// CreateServerlessDBInstance
// 本接口 (CreateServerlessDBInstance) 用于创建一个ServerlessDB实例，创建成功返回实例ID。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_ALLOCATEQUOTASERROR = "FailedOperation.AllocateQuotasError"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_CREATEBASICNETWORKDENIEDERROR = "FailedOperation.CreateBasicNetworkDeniedError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_GETVPCINFOERROR = "FailedOperation.GetVpcInfoError"
//  FAILEDOPERATION_QUERYVPCFAILED = "FailedOperation.QueryVpcFailed"
//  FLOWERROR = "FlowError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_FLOWERROR = "InternalError.FlowError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENAMEEXIST = "InvalidParameter.InstanceNameExist"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_INSTANCENAMEEXIST = "InvalidParameterValue.InstanceNameExist"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTFORMAT = "InvalidParameterValue.InvalidAccountFormat"
//  INVALIDPARAMETERVALUE_INVALIDCHARSET = "InvalidParameterValue.InvalidCharset"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORDFORMAT = "InvalidParameterValue.InvalidPasswordFormat"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  INVALIDPARAMETERVALUE_PARAMETERCHARACTERERROR = "InvalidParameterValue.ParameterCharacterError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_PARAMETERLENGTHLIMITERROR = "InvalidParameterValue.ParameterLengthLimitError"
//  INVALIDPARAMETERVALUE_REGIONNOTSUPPORTED = "InvalidParameterValue.RegionNotSupported"
//  INVALIDPARAMETERVALUE_STRUCTPARSEFAILED = "InvalidParameterValue.StructParseFailed"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_NOTSUPPORTZONEERROR = "OperationDenied.NotSupportZoneError"
//  OPERATIONDENIED_VERSIONNOTSUPPORTERROR = "OperationDenied.VersionNotSupportError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
//  UNAUTHORIZEDOPERATION_USERHASNOREALNAMEAUTHENTICATION = "UnauthorizedOperation.UserHasNoRealnameAuthentication"
func (c *Client) CreateServerlessDBInstanceWithContext(ctx context.Context, request *CreateServerlessDBInstanceRequest) (response *CreateServerlessDBInstanceResponse, err error) {
    if request == nil {
        request = NewCreateServerlessDBInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateServerlessDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateServerlessDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAccountRequest() (request *DeleteAccountRequest) {
    request = &DeleteAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DeleteAccount")
    
    
    return
}

func NewDeleteAccountResponse() (response *DeleteAccountResponse) {
    response = &DeleteAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAccount
// 此接口用于删除数据库账号，需要同时输入Oid与UserName，避免误删。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_PRECHECKERROR = "FailedOperation.PreCheckError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAMEERROR = "InvalidParameterValue.InvalidAccountNameError"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAMEFORMATERROR = "InvalidParameterValue.InvalidAccountNameFormatError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DeleteAccount(request *DeleteAccountRequest) (response *DeleteAccountResponse, err error) {
    return c.DeleteAccountWithContext(context.Background(), request)
}

// DeleteAccount
// 此接口用于删除数据库账号，需要同时输入Oid与UserName，避免误删。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_PRECHECKERROR = "FailedOperation.PreCheckError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAMEERROR = "InvalidParameterValue.InvalidAccountNameError"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAMEFORMATERROR = "InvalidParameterValue.InvalidAccountNameFormatError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DeleteAccountWithContext(ctx context.Context, request *DeleteAccountRequest) (response *DeleteAccountResponse, err error) {
    if request == nil {
        request = NewDeleteAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAccountResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBaseBackupRequest() (request *DeleteBaseBackupRequest) {
    request = &DeleteBaseBackupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DeleteBaseBackup")
    
    
    return
}

func NewDeleteBaseBackupResponse() (response *DeleteBaseBackupResponse) {
    response = &DeleteBaseBackupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteBaseBackup
// 本接口（DeleteBaseBackup）用于删除实例指定数据备份。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DeleteBaseBackup(request *DeleteBaseBackupRequest) (response *DeleteBaseBackupResponse, err error) {
    return c.DeleteBaseBackupWithContext(context.Background(), request)
}

// DeleteBaseBackup
// 本接口（DeleteBaseBackup）用于删除实例指定数据备份。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DeleteBaseBackupWithContext(ctx context.Context, request *DeleteBaseBackupRequest) (response *DeleteBaseBackupResponse, err error) {
    if request == nil {
        request = NewDeleteBaseBackupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteBaseBackup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteBaseBackupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDBInstanceNetworkAccessRequest() (request *DeleteDBInstanceNetworkAccessRequest) {
    request = &DeleteDBInstanceNetworkAccessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DeleteDBInstanceNetworkAccess")
    
    
    return
}

func NewDeleteDBInstanceNetworkAccessResponse() (response *DeleteDBInstanceNetworkAccessResponse) {
    response = &DeleteDBInstanceNetworkAccessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDBInstanceNetworkAccess
// 可对实例进行网络的删除操作。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_GETVPCINFOERROR = "FailedOperation.GetVpcInfoError"
//  FAILEDOPERATION_NETWORKNUMLIMITERROR = "FailedOperation.NetworkNumLimitError"
//  FAILEDOPERATION_QUERYSPECERROR = "FailedOperation.QuerySpecError"
//  FAILEDOPERATION_SERVERLESSNOTSUPPORTEDERROR = "FailedOperation.ServerlessNotSupportedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETER_VPCNOTFOUNDERROR = "InvalidParameter.VpcNotFoundError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  OPERATIONDENIED_ROINSTANCEIPV6NOTSUPPORTEDERROR = "OperationDenied.ROInstanceIpv6NotSupportedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DeleteDBInstanceNetworkAccess(request *DeleteDBInstanceNetworkAccessRequest) (response *DeleteDBInstanceNetworkAccessResponse, err error) {
    return c.DeleteDBInstanceNetworkAccessWithContext(context.Background(), request)
}

// DeleteDBInstanceNetworkAccess
// 可对实例进行网络的删除操作。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_GETVPCINFOERROR = "FailedOperation.GetVpcInfoError"
//  FAILEDOPERATION_NETWORKNUMLIMITERROR = "FailedOperation.NetworkNumLimitError"
//  FAILEDOPERATION_QUERYSPECERROR = "FailedOperation.QuerySpecError"
//  FAILEDOPERATION_SERVERLESSNOTSUPPORTEDERROR = "FailedOperation.ServerlessNotSupportedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETER_VPCNOTFOUNDERROR = "InvalidParameter.VpcNotFoundError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  OPERATIONDENIED_ROINSTANCEIPV6NOTSUPPORTEDERROR = "OperationDenied.ROInstanceIpv6NotSupportedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DeleteDBInstanceNetworkAccessWithContext(ctx context.Context, request *DeleteDBInstanceNetworkAccessRequest) (response *DeleteDBInstanceNetworkAccessResponse, err error) {
    if request == nil {
        request = NewDeleteDBInstanceNetworkAccessRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDBInstanceNetworkAccess require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDBInstanceNetworkAccessResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLogBackupRequest() (request *DeleteLogBackupRequest) {
    request = &DeleteLogBackupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DeleteLogBackup")
    
    
    return
}

func NewDeleteLogBackupResponse() (response *DeleteLogBackupResponse) {
    response = &DeleteLogBackupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLogBackup
// 本接口（DeleteLogBackup）用于删除实例指定日志备份。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DeleteLogBackup(request *DeleteLogBackupRequest) (response *DeleteLogBackupResponse, err error) {
    return c.DeleteLogBackupWithContext(context.Background(), request)
}

// DeleteLogBackup
// 本接口（DeleteLogBackup）用于删除实例指定日志备份。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DeleteLogBackupWithContext(ctx context.Context, request *DeleteLogBackupRequest) (response *DeleteLogBackupResponse, err error) {
    if request == nil {
        request = NewDeleteLogBackupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLogBackup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLogBackupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteParameterTemplateRequest() (request *DeleteParameterTemplateRequest) {
    request = &DeleteParameterTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DeleteParameterTemplate")
    
    
    return
}

func NewDeleteParameterTemplateResponse() (response *DeleteParameterTemplateResponse) {
    response = &DeleteParameterTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteParameterTemplate
// 本接口（DeleteParameterTemplate）主要用于删除某个参数模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
func (c *Client) DeleteParameterTemplate(request *DeleteParameterTemplateRequest) (response *DeleteParameterTemplateResponse, err error) {
    return c.DeleteParameterTemplateWithContext(context.Background(), request)
}

// DeleteParameterTemplate
// 本接口（DeleteParameterTemplate）主要用于删除某个参数模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
func (c *Client) DeleteParameterTemplateWithContext(ctx context.Context, request *DeleteParameterTemplateRequest) (response *DeleteParameterTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteParameterTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteParameterTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteParameterTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteReadOnlyGroupRequest() (request *DeleteReadOnlyGroupRequest) {
    request = &DeleteReadOnlyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DeleteReadOnlyGroup")
    
    
    return
}

func NewDeleteReadOnlyGroupResponse() (response *DeleteReadOnlyGroupResponse) {
    response = &DeleteReadOnlyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteReadOnlyGroup
// 本接口(DeleteReadOnlyGroup)用于删除指定的只读组
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_ROGROUPCANNOTBEDELETEDERROR = "FailedOperation.ROGroupCannotBeDeletedError"
//  FAILEDOPERATION_ROGROUPNOTFOUNDERROR = "FailedOperation.ROGroupNotFoundError"
//  INTERNALERROR_FLOWERROR = "InternalError.FlowError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INTERFACENAMENOTFOUND = "InvalidParameterValue.InterfaceNameNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  SYSTEMERROR = "SystemError"
func (c *Client) DeleteReadOnlyGroup(request *DeleteReadOnlyGroupRequest) (response *DeleteReadOnlyGroupResponse, err error) {
    return c.DeleteReadOnlyGroupWithContext(context.Background(), request)
}

// DeleteReadOnlyGroup
// 本接口(DeleteReadOnlyGroup)用于删除指定的只读组
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_ROGROUPCANNOTBEDELETEDERROR = "FailedOperation.ROGroupCannotBeDeletedError"
//  FAILEDOPERATION_ROGROUPNOTFOUNDERROR = "FailedOperation.ROGroupNotFoundError"
//  INTERNALERROR_FLOWERROR = "InternalError.FlowError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INTERFACENAMENOTFOUND = "InvalidParameterValue.InterfaceNameNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  SYSTEMERROR = "SystemError"
func (c *Client) DeleteReadOnlyGroupWithContext(ctx context.Context, request *DeleteReadOnlyGroupRequest) (response *DeleteReadOnlyGroupResponse, err error) {
    if request == nil {
        request = NewDeleteReadOnlyGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteReadOnlyGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteReadOnlyGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteReadOnlyGroupNetworkAccessRequest() (request *DeleteReadOnlyGroupNetworkAccessRequest) {
    request = &DeleteReadOnlyGroupNetworkAccessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DeleteReadOnlyGroupNetworkAccess")
    
    
    return
}

func NewDeleteReadOnlyGroupNetworkAccessResponse() (response *DeleteReadOnlyGroupNetworkAccessResponse) {
    response = &DeleteReadOnlyGroupNetworkAccessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteReadOnlyGroupNetworkAccess
// 可对RO组进行网络的删除操作。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_GETVPCINFOERROR = "FailedOperation.GetVpcInfoError"
//  FAILEDOPERATION_NETWORKNUMLIMITERROR = "FailedOperation.NetworkNumLimitError"
//  FAILEDOPERATION_ROGROUPNOTFOUNDERROR = "FailedOperation.ROGroupNotFoundError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETER_VPCNOTFOUNDERROR = "InvalidParameter.VpcNotFoundError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_ROGROUPSTATUSERROR = "OperationDenied.ROGroupStatusError"
func (c *Client) DeleteReadOnlyGroupNetworkAccess(request *DeleteReadOnlyGroupNetworkAccessRequest) (response *DeleteReadOnlyGroupNetworkAccessResponse, err error) {
    return c.DeleteReadOnlyGroupNetworkAccessWithContext(context.Background(), request)
}

// DeleteReadOnlyGroupNetworkAccess
// 可对RO组进行网络的删除操作。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_GETVPCINFOERROR = "FailedOperation.GetVpcInfoError"
//  FAILEDOPERATION_NETWORKNUMLIMITERROR = "FailedOperation.NetworkNumLimitError"
//  FAILEDOPERATION_ROGROUPNOTFOUNDERROR = "FailedOperation.ROGroupNotFoundError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETER_VPCNOTFOUNDERROR = "InvalidParameter.VpcNotFoundError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_ROGROUPSTATUSERROR = "OperationDenied.ROGroupStatusError"
func (c *Client) DeleteReadOnlyGroupNetworkAccessWithContext(ctx context.Context, request *DeleteReadOnlyGroupNetworkAccessRequest) (response *DeleteReadOnlyGroupNetworkAccessResponse, err error) {
    if request == nil {
        request = NewDeleteReadOnlyGroupNetworkAccessRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteReadOnlyGroupNetworkAccess require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteReadOnlyGroupNetworkAccessResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteServerlessDBInstanceRequest() (request *DeleteServerlessDBInstanceRequest) {
    request = &DeleteServerlessDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DeleteServerlessDBInstance")
    
    
    return
}

func NewDeleteServerlessDBInstanceResponse() (response *DeleteServerlessDBInstanceResponse) {
    response = &DeleteServerlessDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteServerlessDBInstance
// 本接口 (DeleteServerlessDBInstance) 用于删除一个ServerlessDB实例。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_FLOWERROR = "InternalError.FlowError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDINSTANCESTATUS = "InvalidInstanceStatus"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_STRUCTPARSEFAILED = "InvalidParameterValue.StructParseFailed"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  REGIONNOTSUPPORTED = "RegionNotSupported"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
func (c *Client) DeleteServerlessDBInstance(request *DeleteServerlessDBInstanceRequest) (response *DeleteServerlessDBInstanceResponse, err error) {
    return c.DeleteServerlessDBInstanceWithContext(context.Background(), request)
}

// DeleteServerlessDBInstance
// 本接口 (DeleteServerlessDBInstance) 用于删除一个ServerlessDB实例。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_FLOWERROR = "InternalError.FlowError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDINSTANCESTATUS = "InvalidInstanceStatus"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_STRUCTPARSEFAILED = "InvalidParameterValue.StructParseFailed"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  REGIONNOTSUPPORTED = "RegionNotSupported"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
func (c *Client) DeleteServerlessDBInstanceWithContext(ctx context.Context, request *DeleteServerlessDBInstanceRequest) (response *DeleteServerlessDBInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteServerlessDBInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteServerlessDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteServerlessDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountPrivilegesRequest() (request *DescribeAccountPrivilegesRequest) {
    request = &DescribeAccountPrivilegesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeAccountPrivileges")
    
    
    return
}

func NewDescribeAccountPrivilegesResponse() (response *DescribeAccountPrivilegesResponse) {
    response = &DescribeAccountPrivilegesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccountPrivileges
// 查询数据库账号对某数据库对象拥有的权限列表。
//
// 可能返回的错误码:
//  DBERROR = "DBError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_PRECHECKERROR = "FailedOperation.PreCheckError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAMEERROR = "InvalidParameterValue.InvalidAccountNameError"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAMEFORMATERROR = "InvalidParameterValue.InvalidAccountNameFormatError"
//  INVALIDPARAMETERVALUE_INVALIDDATABASEOBJECTNAME = "InvalidParameterValue.InvalidDatabaseObjectName"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeAccountPrivileges(request *DescribeAccountPrivilegesRequest) (response *DescribeAccountPrivilegesResponse, err error) {
    return c.DescribeAccountPrivilegesWithContext(context.Background(), request)
}

// DescribeAccountPrivileges
// 查询数据库账号对某数据库对象拥有的权限列表。
//
// 可能返回的错误码:
//  DBERROR = "DBError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_PRECHECKERROR = "FailedOperation.PreCheckError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAMEERROR = "InvalidParameterValue.InvalidAccountNameError"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAMEFORMATERROR = "InvalidParameterValue.InvalidAccountNameFormatError"
//  INVALIDPARAMETERVALUE_INVALIDDATABASEOBJECTNAME = "InvalidParameterValue.InvalidDatabaseObjectName"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeAccountPrivilegesWithContext(ctx context.Context, request *DescribeAccountPrivilegesRequest) (response *DescribeAccountPrivilegesResponse, err error) {
    if request == nil {
        request = NewDescribeAccountPrivilegesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccountPrivileges require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccountPrivilegesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountsRequest() (request *DescribeAccountsRequest) {
    request = &DescribeAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeAccounts")
    
    
    return
}

func NewDescribeAccountsResponse() (response *DescribeAccountsResponse) {
    response = &DescribeAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccounts
// 本接口（DescribeAccounts）用于查询实例的数据库账号列表。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DBERROR = "DBError"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_PRECHECKERROR = "FailedOperation.PreCheckError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXIST = "InvalidParameterValue.AccountNotExist"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeAccounts(request *DescribeAccountsRequest) (response *DescribeAccountsResponse, err error) {
    return c.DescribeAccountsWithContext(context.Background(), request)
}

// DescribeAccounts
// 本接口（DescribeAccounts）用于查询实例的数据库账号列表。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DBERROR = "DBError"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_PRECHECKERROR = "FailedOperation.PreCheckError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXIST = "InvalidParameterValue.AccountNotExist"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeAccountsWithContext(ctx context.Context, request *DescribeAccountsRequest) (response *DescribeAccountsResponse, err error) {
    if request == nil {
        request = NewDescribeAccountsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccounts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAvailableRecoveryTimeRequest() (request *DescribeAvailableRecoveryTimeRequest) {
    request = &DescribeAvailableRecoveryTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeAvailableRecoveryTime")
    
    
    return
}

func NewDescribeAvailableRecoveryTimeResponse() (response *DescribeAvailableRecoveryTimeResponse) {
    response = &DescribeAvailableRecoveryTimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAvailableRecoveryTime
// 本接口（DescribeAvailableRecoveryTime）用于查询实例可恢复的时间范围。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeAvailableRecoveryTime(request *DescribeAvailableRecoveryTimeRequest) (response *DescribeAvailableRecoveryTimeResponse, err error) {
    return c.DescribeAvailableRecoveryTimeWithContext(context.Background(), request)
}

// DescribeAvailableRecoveryTime
// 本接口（DescribeAvailableRecoveryTime）用于查询实例可恢复的时间范围。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeAvailableRecoveryTimeWithContext(ctx context.Context, request *DescribeAvailableRecoveryTimeRequest) (response *DescribeAvailableRecoveryTimeResponse, err error) {
    if request == nil {
        request = NewDescribeAvailableRecoveryTimeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAvailableRecoveryTime require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAvailableRecoveryTimeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupDownloadRestrictionRequest() (request *DescribeBackupDownloadRestrictionRequest) {
    request = &DescribeBackupDownloadRestrictionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeBackupDownloadRestriction")
    
    
    return
}

func NewDescribeBackupDownloadRestrictionResponse() (response *DescribeBackupDownloadRestrictionResponse) {
    response = &DescribeBackupDownloadRestrictionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackupDownloadRestriction
// 本接口（DescribeBackupDownloadRestriction）用于查询备份文件下载限制。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
func (c *Client) DescribeBackupDownloadRestriction(request *DescribeBackupDownloadRestrictionRequest) (response *DescribeBackupDownloadRestrictionResponse, err error) {
    return c.DescribeBackupDownloadRestrictionWithContext(context.Background(), request)
}

// DescribeBackupDownloadRestriction
// 本接口（DescribeBackupDownloadRestriction）用于查询备份文件下载限制。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
func (c *Client) DescribeBackupDownloadRestrictionWithContext(ctx context.Context, request *DescribeBackupDownloadRestrictionRequest) (response *DescribeBackupDownloadRestrictionResponse, err error) {
    if request == nil {
        request = NewDescribeBackupDownloadRestrictionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupDownloadRestriction require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupDownloadRestrictionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupDownloadURLRequest() (request *DescribeBackupDownloadURLRequest) {
    request = &DescribeBackupDownloadURLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeBackupDownloadURL")
    
    
    return
}

func NewDescribeBackupDownloadURLResponse() (response *DescribeBackupDownloadURLResponse) {
    response = &DescribeBackupDownloadURLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackupDownloadURL
// 本接口 (DescribeBackupDownloadURL) 用于查询指定备份集的下载地址，可包括全量备份集、增量日志备份集。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INSTANCEDATAERROR = "InternalError.InstanceDataError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEEXCEEDERROR = "InvalidParameterValue.ParameterValueExceedError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeBackupDownloadURL(request *DescribeBackupDownloadURLRequest) (response *DescribeBackupDownloadURLResponse, err error) {
    return c.DescribeBackupDownloadURLWithContext(context.Background(), request)
}

// DescribeBackupDownloadURL
// 本接口 (DescribeBackupDownloadURL) 用于查询指定备份集的下载地址，可包括全量备份集、增量日志备份集。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INSTANCEDATAERROR = "InternalError.InstanceDataError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEEXCEEDERROR = "InvalidParameterValue.ParameterValueExceedError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeBackupDownloadURLWithContext(ctx context.Context, request *DescribeBackupDownloadURLRequest) (response *DescribeBackupDownloadURLResponse, err error) {
    if request == nil {
        request = NewDescribeBackupDownloadURLRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupDownloadURL require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupDownloadURLResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupOverviewRequest() (request *DescribeBackupOverviewRequest) {
    request = &DescribeBackupOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeBackupOverview")
    
    
    return
}

func NewDescribeBackupOverviewResponse() (response *DescribeBackupOverviewResponse) {
    response = &DescribeBackupOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackupOverview
// 本接口（DescribeBackupOverview）用于查询用户的备份概览信息。返回用户当前备份个数、备份占用容量、免费容量、收费容量等信息（容量单位为字节）。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
func (c *Client) DescribeBackupOverview(request *DescribeBackupOverviewRequest) (response *DescribeBackupOverviewResponse, err error) {
    return c.DescribeBackupOverviewWithContext(context.Background(), request)
}

// DescribeBackupOverview
// 本接口（DescribeBackupOverview）用于查询用户的备份概览信息。返回用户当前备份个数、备份占用容量、免费容量、收费容量等信息（容量单位为字节）。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
func (c *Client) DescribeBackupOverviewWithContext(ctx context.Context, request *DescribeBackupOverviewRequest) (response *DescribeBackupOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeBackupOverviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupPlansRequest() (request *DescribeBackupPlansRequest) {
    request = &DescribeBackupPlansRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeBackupPlans")
    
    
    return
}

func NewDescribeBackupPlansResponse() (response *DescribeBackupPlansResponse) {
    response = &DescribeBackupPlansResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackupPlans
// 本接口 (DescribeBackupPlans) 用于实例所有的备份计划查询
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeBackupPlans(request *DescribeBackupPlansRequest) (response *DescribeBackupPlansResponse, err error) {
    return c.DescribeBackupPlansWithContext(context.Background(), request)
}

// DescribeBackupPlans
// 本接口 (DescribeBackupPlans) 用于实例所有的备份计划查询
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeBackupPlansWithContext(ctx context.Context, request *DescribeBackupPlansRequest) (response *DescribeBackupPlansResponse, err error) {
    if request == nil {
        request = NewDescribeBackupPlansRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupPlans require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupPlansResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupSummariesRequest() (request *DescribeBackupSummariesRequest) {
    request = &DescribeBackupSummariesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeBackupSummaries")
    
    
    return
}

func NewDescribeBackupSummariesResponse() (response *DescribeBackupSummariesResponse) {
    response = &DescribeBackupSummariesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackupSummaries
// 本接口(DescribeBackupSummaries)用于查询实例备份的统计信息，返回以实例为维度的备份个数、占用容量等信息（容量单位为字节）。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeBackupSummaries(request *DescribeBackupSummariesRequest) (response *DescribeBackupSummariesResponse, err error) {
    return c.DescribeBackupSummariesWithContext(context.Background(), request)
}

// DescribeBackupSummaries
// 本接口(DescribeBackupSummaries)用于查询实例备份的统计信息，返回以实例为维度的备份个数、占用容量等信息（容量单位为字节）。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeBackupSummariesWithContext(ctx context.Context, request *DescribeBackupSummariesRequest) (response *DescribeBackupSummariesResponse, err error) {
    if request == nil {
        request = NewDescribeBackupSummariesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupSummaries require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupSummariesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaseBackupsRequest() (request *DescribeBaseBackupsRequest) {
    request = &DescribeBaseBackupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeBaseBackups")
    
    
    return
}

func NewDescribeBaseBackupsResponse() (response *DescribeBaseBackupsResponse) {
    response = &DescribeBaseBackupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBaseBackups
// 本接口（DescribeBaseBackups）用于查询数据备份列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
func (c *Client) DescribeBaseBackups(request *DescribeBaseBackupsRequest) (response *DescribeBaseBackupsResponse, err error) {
    return c.DescribeBaseBackupsWithContext(context.Background(), request)
}

// DescribeBaseBackups
// 本接口（DescribeBaseBackups）用于查询数据备份列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
func (c *Client) DescribeBaseBackupsWithContext(ctx context.Context, request *DescribeBaseBackupsRequest) (response *DescribeBaseBackupsResponse, err error) {
    if request == nil {
        request = NewDescribeBaseBackupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBaseBackups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBaseBackupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClassesRequest() (request *DescribeClassesRequest) {
    request = &DescribeClassesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeClasses")
    
    
    return
}

func NewDescribeClassesResponse() (response *DescribeClassesResponse) {
    response = &DescribeClassesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClasses
// 本接口（DescribeClasses）用于查询实例售卖规格。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_WHITELISTCONNECTERROR = "FailedOperation.WhitelistConnectError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
func (c *Client) DescribeClasses(request *DescribeClassesRequest) (response *DescribeClassesResponse, err error) {
    return c.DescribeClassesWithContext(context.Background(), request)
}

// DescribeClasses
// 本接口（DescribeClasses）用于查询实例售卖规格。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_WHITELISTCONNECTERROR = "FailedOperation.WhitelistConnectError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
func (c *Client) DescribeClassesWithContext(ctx context.Context, request *DescribeClassesRequest) (response *DescribeClassesResponse, err error) {
    if request == nil {
        request = NewDescribeClassesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClasses require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClassesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloneDBInstanceSpecRequest() (request *DescribeCloneDBInstanceSpecRequest) {
    request = &DescribeCloneDBInstanceSpecRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeCloneDBInstanceSpec")
    
    
    return
}

func NewDescribeCloneDBInstanceSpecResponse() (response *DescribeCloneDBInstanceSpecResponse) {
    response = &DescribeCloneDBInstanceSpecResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloneDBInstanceSpec
// 本接口（DescribeCloneDBInstanceSpec）用于查询克隆实例可选择的最小规格，包括SpecCode和磁盘。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_QUERYSPECERROR = "FailedOperation.QuerySpecError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeCloneDBInstanceSpec(request *DescribeCloneDBInstanceSpecRequest) (response *DescribeCloneDBInstanceSpecResponse, err error) {
    return c.DescribeCloneDBInstanceSpecWithContext(context.Background(), request)
}

// DescribeCloneDBInstanceSpec
// 本接口（DescribeCloneDBInstanceSpec）用于查询克隆实例可选择的最小规格，包括SpecCode和磁盘。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_QUERYSPECERROR = "FailedOperation.QuerySpecError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeCloneDBInstanceSpecWithContext(ctx context.Context, request *DescribeCloneDBInstanceSpecRequest) (response *DescribeCloneDBInstanceSpecResponse, err error) {
    if request == nil {
        request = NewDescribeCloneDBInstanceSpecRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloneDBInstanceSpec require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloneDBInstanceSpecResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBBackupsRequest() (request *DescribeDBBackupsRequest) {
    request = &DescribeDBBackupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeDBBackups")
    
    
    return
}

func NewDescribeDBBackupsResponse() (response *DescribeDBBackupsResponse) {
    response = &DescribeDBBackupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBBackups
// 本接口（DescribeDBBackups）用于查询实例备份列表。本接口已废弃，推荐使用接口[DescribeBaseBackups](https://cloud.tencent.com/document/api/409/89022)替代。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_OSSACCESSERROR = "FailedOperation.OssAccessError"
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  FAILEDOPERATION_PRESIGNEDURLGETERROR = "FailedOperation.PresignedURLGetError"
//  FAILEDOPERATION_STOREPATHSPLITERROR = "FailedOperation.StorePathSplitError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INSTANCEDATAERROR = "InternalError.InstanceDataError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNKNOWNERROR = "UnknownError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDBBackups(request *DescribeDBBackupsRequest) (response *DescribeDBBackupsResponse, err error) {
    return c.DescribeDBBackupsWithContext(context.Background(), request)
}

// DescribeDBBackups
// 本接口（DescribeDBBackups）用于查询实例备份列表。本接口已废弃，推荐使用接口[DescribeBaseBackups](https://cloud.tencent.com/document/api/409/89022)替代。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_OSSACCESSERROR = "FailedOperation.OssAccessError"
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  FAILEDOPERATION_PRESIGNEDURLGETERROR = "FailedOperation.PresignedURLGetError"
//  FAILEDOPERATION_STOREPATHSPLITERROR = "FailedOperation.StorePathSplitError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INSTANCEDATAERROR = "InternalError.InstanceDataError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNKNOWNERROR = "UnknownError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDBBackupsWithContext(ctx context.Context, request *DescribeDBBackupsRequest) (response *DescribeDBBackupsResponse, err error) {
    if request == nil {
        request = NewDescribeDBBackupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBBackups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBBackupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBErrlogsRequest() (request *DescribeDBErrlogsRequest) {
    request = &DescribeDBErrlogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeDBErrlogs")
    
    
    return
}

func NewDescribeDBErrlogsResponse() (response *DescribeDBErrlogsResponse) {
    response = &DescribeDBErrlogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBErrlogs
// 本接口（DescribeDBErrlogs）用于查询错误日志。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_ESCONNECTERROR = "FailedOperation.ESConnectError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INTERFACENAMENOTFOUND = "InvalidParameterValue.InterfaceNameNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEEXCEEDERROR = "InvalidParameterValue.ParameterValueExceedError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeDBErrlogs(request *DescribeDBErrlogsRequest) (response *DescribeDBErrlogsResponse, err error) {
    return c.DescribeDBErrlogsWithContext(context.Background(), request)
}

// DescribeDBErrlogs
// 本接口（DescribeDBErrlogs）用于查询错误日志。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_ESCONNECTERROR = "FailedOperation.ESConnectError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INTERFACENAMENOTFOUND = "InvalidParameterValue.InterfaceNameNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEEXCEEDERROR = "InvalidParameterValue.ParameterValueExceedError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeDBErrlogsWithContext(ctx context.Context, request *DescribeDBErrlogsRequest) (response *DescribeDBErrlogsResponse, err error) {
    if request == nil {
        request = NewDescribeDBErrlogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBErrlogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBErrlogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstanceAttributeRequest() (request *DescribeDBInstanceAttributeRequest) {
    request = &DescribeDBInstanceAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeDBInstanceAttribute")
    
    
    return
}

func NewDescribeDBInstanceAttributeResponse() (response *DescribeDBInstanceAttributeResponse) {
    response = &DescribeDBInstanceAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBInstanceAttribute
// 本接口 (DescribeDBInstanceAttribute) 用于查询某个实例的详情信息。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_QUERYINSTANCEBATCHERROR = "FailedOperation.QueryInstanceBatchError"
//  FAILEDOPERATION_QUERYSPECBYSPECCODEERROR = "FailedOperation.QuerySpecBySpecCodeError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_TRANSACTIOBEGINERROR = "InternalError.TransactioBeginError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeDBInstanceAttribute(request *DescribeDBInstanceAttributeRequest) (response *DescribeDBInstanceAttributeResponse, err error) {
    return c.DescribeDBInstanceAttributeWithContext(context.Background(), request)
}

// DescribeDBInstanceAttribute
// 本接口 (DescribeDBInstanceAttribute) 用于查询某个实例的详情信息。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_QUERYINSTANCEBATCHERROR = "FailedOperation.QueryInstanceBatchError"
//  FAILEDOPERATION_QUERYSPECBYSPECCODEERROR = "FailedOperation.QuerySpecBySpecCodeError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_TRANSACTIOBEGINERROR = "InternalError.TransactioBeginError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeDBInstanceAttributeWithContext(ctx context.Context, request *DescribeDBInstanceAttributeRequest) (response *DescribeDBInstanceAttributeResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBInstanceAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBInstanceAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstanceHAConfigRequest() (request *DescribeDBInstanceHAConfigRequest) {
    request = &DescribeDBInstanceHAConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeDBInstanceHAConfig")
    
    
    return
}

func NewDescribeDBInstanceHAConfigResponse() (response *DescribeDBInstanceHAConfigResponse) {
    response = &DescribeDBInstanceHAConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBInstanceHAConfig
// 本接口（DescribeDBInstanceHAConfig）用于查询实例HA配置信息。其中HA配置信息包括：
//
// <li>允许备节点切换为主节点的条件配置
//
// <li>半同步实例使用同步复制或异步复制的条件配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeDBInstanceHAConfig(request *DescribeDBInstanceHAConfigRequest) (response *DescribeDBInstanceHAConfigResponse, err error) {
    return c.DescribeDBInstanceHAConfigWithContext(context.Background(), request)
}

// DescribeDBInstanceHAConfig
// 本接口（DescribeDBInstanceHAConfig）用于查询实例HA配置信息。其中HA配置信息包括：
//
// <li>允许备节点切换为主节点的条件配置
//
// <li>半同步实例使用同步复制或异步复制的条件配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeDBInstanceHAConfigWithContext(ctx context.Context, request *DescribeDBInstanceHAConfigRequest) (response *DescribeDBInstanceHAConfigResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceHAConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBInstanceHAConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBInstanceHAConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstanceParametersRequest() (request *DescribeDBInstanceParametersRequest) {
    request = &DescribeDBInstanceParametersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeDBInstanceParameters")
    
    
    return
}

func NewDescribeDBInstanceParametersResponse() (response *DescribeDBInstanceParametersResponse) {
    response = &DescribeDBInstanceParametersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBInstanceParameters
// 本接口（DescribeDBInstanceAttribute）用于查询实例的参数信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNKNOWNERROR = "UnknownError"
func (c *Client) DescribeDBInstanceParameters(request *DescribeDBInstanceParametersRequest) (response *DescribeDBInstanceParametersResponse, err error) {
    return c.DescribeDBInstanceParametersWithContext(context.Background(), request)
}

// DescribeDBInstanceParameters
// 本接口（DescribeDBInstanceAttribute）用于查询实例的参数信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNKNOWNERROR = "UnknownError"
func (c *Client) DescribeDBInstanceParametersWithContext(ctx context.Context, request *DescribeDBInstanceParametersRequest) (response *DescribeDBInstanceParametersResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceParametersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBInstanceParameters require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBInstanceParametersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstanceSecurityGroupsRequest() (request *DescribeDBInstanceSecurityGroupsRequest) {
    request = &DescribeDBInstanceSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeDBInstanceSecurityGroups")
    
    
    return
}

func NewDescribeDBInstanceSecurityGroupsResponse() (response *DescribeDBInstanceSecurityGroupsResponse) {
    response = &DescribeDBInstanceSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBInstanceSecurityGroups
// 本接口（DescribeDBInstanceSecurityGroups）用于查询实例安全组。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_ROGROUPNOTFOUNDERROR = "FailedOperation.ROGroupNotFoundError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  OPERATIONDENIED_ROGROUPSTATUSERROR = "OperationDenied.ROGroupStatusError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeDBInstanceSecurityGroups(request *DescribeDBInstanceSecurityGroupsRequest) (response *DescribeDBInstanceSecurityGroupsResponse, err error) {
    return c.DescribeDBInstanceSecurityGroupsWithContext(context.Background(), request)
}

// DescribeDBInstanceSecurityGroups
// 本接口（DescribeDBInstanceSecurityGroups）用于查询实例安全组。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_ROGROUPNOTFOUNDERROR = "FailedOperation.ROGroupNotFoundError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  OPERATIONDENIED_ROGROUPSTATUSERROR = "OperationDenied.ROGroupStatusError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeDBInstanceSecurityGroupsWithContext(ctx context.Context, request *DescribeDBInstanceSecurityGroupsRequest) (response *DescribeDBInstanceSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceSecurityGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBInstanceSecurityGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBInstanceSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstancesRequest() (request *DescribeDBInstancesRequest) {
    request = &DescribeDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeDBInstances")
    
    
    return
}

func NewDescribeDBInstancesResponse() (response *DescribeDBInstancesResponse) {
    response = &DescribeDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBInstances
// 本接口 (DescribeDBInstances) 用于查询一个或多个实例的详细信息。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_CAMCHECKRESOURCEFAILED = "FailedOperation.CamCheckResourceFailed"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_LIMITOPERATION = "FailedOperation.LimitOperation"
//  FAILEDOPERATION_OSSACCESSERROR = "FailedOperation.OssAccessError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_TRANSACTIOBEGINERROR = "InternalError.TransactioBeginError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDREGIONIDERROR = "InvalidParameterValue.InvalidRegionIdError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_REGIONNOTSUPPORTED = "InvalidParameterValue.RegionNotSupported"
//  INVALIDPARAMETERVALUE_STRUCTPARSEFAILED = "InvalidParameterValue.StructParseFailed"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDBInstances(request *DescribeDBInstancesRequest) (response *DescribeDBInstancesResponse, err error) {
    return c.DescribeDBInstancesWithContext(context.Background(), request)
}

// DescribeDBInstances
// 本接口 (DescribeDBInstances) 用于查询一个或多个实例的详细信息。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_CAMCHECKRESOURCEFAILED = "FailedOperation.CamCheckResourceFailed"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_LIMITOPERATION = "FailedOperation.LimitOperation"
//  FAILEDOPERATION_OSSACCESSERROR = "FailedOperation.OssAccessError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_TRANSACTIOBEGINERROR = "InternalError.TransactioBeginError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDREGIONIDERROR = "InvalidParameterValue.InvalidRegionIdError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_REGIONNOTSUPPORTED = "InvalidParameterValue.RegionNotSupported"
//  INVALIDPARAMETERVALUE_STRUCTPARSEFAILED = "InvalidParameterValue.StructParseFailed"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDBInstancesWithContext(ctx context.Context, request *DescribeDBInstancesRequest) (response *DescribeDBInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBSlowlogsRequest() (request *DescribeDBSlowlogsRequest) {
    request = &DescribeDBSlowlogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeDBSlowlogs")
    
    
    return
}

func NewDescribeDBSlowlogsResponse() (response *DescribeDBSlowlogsResponse) {
    response = &DescribeDBSlowlogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBSlowlogs
// 本接口（DescribeDBSlowlogs）用于获取慢查询日志。本接口已于2021.09.01日废弃，后续此接口将不再返回任何数据，推荐使用接口[DescribeSlowQueryList](https://cloud.tencent.com/document/api/409/60540)替代。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_ESCONNECTERROR = "FailedOperation.ESConnectError"
//  FAILEDOPERATION_ESQUERYERROR = "FailedOperation.ESQueryError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeDBSlowlogs(request *DescribeDBSlowlogsRequest) (response *DescribeDBSlowlogsResponse, err error) {
    return c.DescribeDBSlowlogsWithContext(context.Background(), request)
}

// DescribeDBSlowlogs
// 本接口（DescribeDBSlowlogs）用于获取慢查询日志。本接口已于2021.09.01日废弃，后续此接口将不再返回任何数据，推荐使用接口[DescribeSlowQueryList](https://cloud.tencent.com/document/api/409/60540)替代。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_ESCONNECTERROR = "FailedOperation.ESConnectError"
//  FAILEDOPERATION_ESQUERYERROR = "FailedOperation.ESQueryError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeDBSlowlogsWithContext(ctx context.Context, request *DescribeDBSlowlogsRequest) (response *DescribeDBSlowlogsResponse, err error) {
    if request == nil {
        request = NewDescribeDBSlowlogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBSlowlogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBSlowlogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBVersionsRequest() (request *DescribeDBVersionsRequest) {
    request = &DescribeDBVersionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeDBVersions")
    
    
    return
}

func NewDescribeDBVersionsResponse() (response *DescribeDBVersionsResponse) {
    response = &DescribeDBVersionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBVersions
// 本接口（DescribeDBVersions）用于查询支持的数据库版本。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
func (c *Client) DescribeDBVersions(request *DescribeDBVersionsRequest) (response *DescribeDBVersionsResponse, err error) {
    return c.DescribeDBVersionsWithContext(context.Background(), request)
}

// DescribeDBVersions
// 本接口（DescribeDBVersions）用于查询支持的数据库版本。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
func (c *Client) DescribeDBVersionsWithContext(ctx context.Context, request *DescribeDBVersionsRequest) (response *DescribeDBVersionsResponse, err error) {
    if request == nil {
        request = NewDescribeDBVersionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBVersions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBVersionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBXlogsRequest() (request *DescribeDBXlogsRequest) {
    request = &DescribeDBXlogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeDBXlogs")
    
    
    return
}

func NewDescribeDBXlogsResponse() (response *DescribeDBXlogsResponse) {
    response = &DescribeDBXlogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBXlogs
// 本接口（DescribeDBXlogs）用于获取实例Xlog列表。本接口已废弃，推荐使用接口[DescribeLogBackups](https://cloud.tencent.com/document/api/409/89021)替代。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_OSSACCESSERROR = "FailedOperation.OssAccessError"
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  FAILEDOPERATION_PRESIGNEDURLGETERROR = "FailedOperation.PresignedURLGetError"
//  FAILEDOPERATION_STOREPATHSPLITERROR = "FailedOperation.StorePathSplitError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INSTANCEDATAERROR = "InternalError.InstanceDataError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNKNOWNERROR = "UnknownError"
func (c *Client) DescribeDBXlogs(request *DescribeDBXlogsRequest) (response *DescribeDBXlogsResponse, err error) {
    return c.DescribeDBXlogsWithContext(context.Background(), request)
}

// DescribeDBXlogs
// 本接口（DescribeDBXlogs）用于获取实例Xlog列表。本接口已废弃，推荐使用接口[DescribeLogBackups](https://cloud.tencent.com/document/api/409/89021)替代。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_OSSACCESSERROR = "FailedOperation.OssAccessError"
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  FAILEDOPERATION_PRESIGNEDURLGETERROR = "FailedOperation.PresignedURLGetError"
//  FAILEDOPERATION_STOREPATHSPLITERROR = "FailedOperation.StorePathSplitError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INSTANCEDATAERROR = "InternalError.InstanceDataError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNKNOWNERROR = "UnknownError"
func (c *Client) DescribeDBXlogsWithContext(ctx context.Context, request *DescribeDBXlogsRequest) (response *DescribeDBXlogsResponse, err error) {
    if request == nil {
        request = NewDescribeDBXlogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBXlogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBXlogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatabaseObjectsRequest() (request *DescribeDatabaseObjectsRequest) {
    request = &DescribeDatabaseObjectsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeDatabaseObjects")
    
    
    return
}

func NewDescribeDatabaseObjectsResponse() (response *DescribeDatabaseObjectsResponse) {
    response = &DescribeDatabaseObjectsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDatabaseObjects
// 本接口用于查询数据库对象列表。例如查询test数据库下的模式列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_PRECHECKERROR = "FailedOperation.PreCheckError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeDatabaseObjects(request *DescribeDatabaseObjectsRequest) (response *DescribeDatabaseObjectsResponse, err error) {
    return c.DescribeDatabaseObjectsWithContext(context.Background(), request)
}

// DescribeDatabaseObjects
// 本接口用于查询数据库对象列表。例如查询test数据库下的模式列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_PRECHECKERROR = "FailedOperation.PreCheckError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeDatabaseObjectsWithContext(ctx context.Context, request *DescribeDatabaseObjectsRequest) (response *DescribeDatabaseObjectsResponse, err error) {
    if request == nil {
        request = NewDescribeDatabaseObjectsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDatabaseObjects require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDatabaseObjectsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatabasesRequest() (request *DescribeDatabasesRequest) {
    request = &DescribeDatabasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeDatabases")
    
    
    return
}

func NewDescribeDatabasesResponse() (response *DescribeDatabasesResponse) {
    response = &DescribeDatabasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDatabases
// 接口（DescribeDatabases）用来查询实例的数据库列表。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DBERROR = "DBError"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_LIMITOPERATION = "FailedOperation.LimitOperation"
//  FAILEDOPERATION_OSSACCESSERROR = "FailedOperation.OssAccessError"
//  FAILEDOPERATION_PRECHECKERROR = "FailedOperation.PreCheckError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INSTANCEDATAERROR = "InternalError.InstanceDataError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDINSTANCESTATUS = "InvalidInstanceStatus"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
func (c *Client) DescribeDatabases(request *DescribeDatabasesRequest) (response *DescribeDatabasesResponse, err error) {
    return c.DescribeDatabasesWithContext(context.Background(), request)
}

// DescribeDatabases
// 接口（DescribeDatabases）用来查询实例的数据库列表。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DBERROR = "DBError"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_LIMITOPERATION = "FailedOperation.LimitOperation"
//  FAILEDOPERATION_OSSACCESSERROR = "FailedOperation.OssAccessError"
//  FAILEDOPERATION_PRECHECKERROR = "FailedOperation.PreCheckError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INSTANCEDATAERROR = "InternalError.InstanceDataError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDINSTANCESTATUS = "InvalidInstanceStatus"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
func (c *Client) DescribeDatabasesWithContext(ctx context.Context, request *DescribeDatabasesRequest) (response *DescribeDatabasesResponse, err error) {
    if request == nil {
        request = NewDescribeDatabasesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDatabases require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDatabasesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDefaultParametersRequest() (request *DescribeDefaultParametersRequest) {
    request = &DescribeDefaultParametersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeDefaultParameters")
    
    
    return
}

func NewDescribeDefaultParametersResponse() (response *DescribeDefaultParametersResponse) {
    response = &DescribeDefaultParametersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDefaultParameters
// 本接口（DescribeDefaultParameters）主要用于查询某个数据库版本和引擎支持的所有参数。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) DescribeDefaultParameters(request *DescribeDefaultParametersRequest) (response *DescribeDefaultParametersResponse, err error) {
    return c.DescribeDefaultParametersWithContext(context.Background(), request)
}

// DescribeDefaultParameters
// 本接口（DescribeDefaultParameters）主要用于查询某个数据库版本和引擎支持的所有参数。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) DescribeDefaultParametersWithContext(ctx context.Context, request *DescribeDefaultParametersRequest) (response *DescribeDefaultParametersResponse, err error) {
    if request == nil {
        request = NewDescribeDefaultParametersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDefaultParameters require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDefaultParametersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEncryptionKeysRequest() (request *DescribeEncryptionKeysRequest) {
    request = &DescribeEncryptionKeysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeEncryptionKeys")
    
    
    return
}

func NewDescribeEncryptionKeysResponse() (response *DescribeEncryptionKeysResponse) {
    response = &DescribeEncryptionKeysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEncryptionKeys
// 本接口 （DescribeEncryptionKeys） 用于查询实例的密钥信息列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
func (c *Client) DescribeEncryptionKeys(request *DescribeEncryptionKeysRequest) (response *DescribeEncryptionKeysResponse, err error) {
    return c.DescribeEncryptionKeysWithContext(context.Background(), request)
}

// DescribeEncryptionKeys
// 本接口 （DescribeEncryptionKeys） 用于查询实例的密钥信息列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
func (c *Client) DescribeEncryptionKeysWithContext(ctx context.Context, request *DescribeEncryptionKeysRequest) (response *DescribeEncryptionKeysResponse, err error) {
    if request == nil {
        request = NewDescribeEncryptionKeysRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEncryptionKeys require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEncryptionKeysResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogBackupsRequest() (request *DescribeLogBackupsRequest) {
    request = &DescribeLogBackupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeLogBackups")
    
    
    return
}

func NewDescribeLogBackupsResponse() (response *DescribeLogBackupsResponse) {
    response = &DescribeLogBackupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLogBackups
// 本接口（DescribeLogBackups）用于查询日志备份列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
func (c *Client) DescribeLogBackups(request *DescribeLogBackupsRequest) (response *DescribeLogBackupsResponse, err error) {
    return c.DescribeLogBackupsWithContext(context.Background(), request)
}

// DescribeLogBackups
// 本接口（DescribeLogBackups）用于查询日志备份列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
func (c *Client) DescribeLogBackupsWithContext(ctx context.Context, request *DescribeLogBackupsRequest) (response *DescribeLogBackupsResponse, err error) {
    if request == nil {
        request = NewDescribeLogBackupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogBackups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogBackupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrdersRequest() (request *DescribeOrdersRequest) {
    request = &DescribeOrdersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeOrders")
    
    
    return
}

func NewDescribeOrdersResponse() (response *DescribeOrdersResponse) {
    response = &DescribeOrdersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOrders
// 本接口（DescribeOrders）用于查询订单信息。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_QUERYDEALFAILED = "FailedOperation.QueryDealFailed"
//  FAILEDOPERATION_TRADEQUERYORDERERROR = "FailedOperation.TradeQueryOrderError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_NULLDEALNAMES = "InvalidParameterValue.NullDealNames"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeOrders(request *DescribeOrdersRequest) (response *DescribeOrdersResponse, err error) {
    return c.DescribeOrdersWithContext(context.Background(), request)
}

// DescribeOrders
// 本接口（DescribeOrders）用于查询订单信息。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_QUERYDEALFAILED = "FailedOperation.QueryDealFailed"
//  FAILEDOPERATION_TRADEQUERYORDERERROR = "FailedOperation.TradeQueryOrderError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_NULLDEALNAMES = "InvalidParameterValue.NullDealNames"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeOrdersWithContext(ctx context.Context, request *DescribeOrdersRequest) (response *DescribeOrdersResponse, err error) {
    if request == nil {
        request = NewDescribeOrdersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrders require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrdersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeParameterTemplateAttributesRequest() (request *DescribeParameterTemplateAttributesRequest) {
    request = &DescribeParameterTemplateAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeParameterTemplateAttributes")
    
    
    return
}

func NewDescribeParameterTemplateAttributesResponse() (response *DescribeParameterTemplateAttributesResponse) {
    response = &DescribeParameterTemplateAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeParameterTemplateAttributes
// 本接口（DescribeParameterTemplateAttributes）用于查询某个参数模板的具体内容，包括基本信息和参数信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
func (c *Client) DescribeParameterTemplateAttributes(request *DescribeParameterTemplateAttributesRequest) (response *DescribeParameterTemplateAttributesResponse, err error) {
    return c.DescribeParameterTemplateAttributesWithContext(context.Background(), request)
}

// DescribeParameterTemplateAttributes
// 本接口（DescribeParameterTemplateAttributes）用于查询某个参数模板的具体内容，包括基本信息和参数信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
func (c *Client) DescribeParameterTemplateAttributesWithContext(ctx context.Context, request *DescribeParameterTemplateAttributesRequest) (response *DescribeParameterTemplateAttributesResponse, err error) {
    if request == nil {
        request = NewDescribeParameterTemplateAttributesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeParameterTemplateAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeParameterTemplateAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeParameterTemplatesRequest() (request *DescribeParameterTemplatesRequest) {
    request = &DescribeParameterTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeParameterTemplates")
    
    
    return
}

func NewDescribeParameterTemplatesResponse() (response *DescribeParameterTemplatesResponse) {
    response = &DescribeParameterTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeParameterTemplates
// 本接口 (DescribeParameterTemplates) 用于查询参数模板列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
func (c *Client) DescribeParameterTemplates(request *DescribeParameterTemplatesRequest) (response *DescribeParameterTemplatesResponse, err error) {
    return c.DescribeParameterTemplatesWithContext(context.Background(), request)
}

// DescribeParameterTemplates
// 本接口 (DescribeParameterTemplates) 用于查询参数模板列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
func (c *Client) DescribeParameterTemplatesWithContext(ctx context.Context, request *DescribeParameterTemplatesRequest) (response *DescribeParameterTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeParameterTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeParameterTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeParameterTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeParamsEventRequest() (request *DescribeParamsEventRequest) {
    request = &DescribeParamsEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeParamsEvent")
    
    
    return
}

func NewDescribeParamsEventResponse() (response *DescribeParamsEventResponse) {
    response = &DescribeParamsEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeParamsEvent
// 本接口（DescribeParamsEvent）用于查询参数修改事件。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeParamsEvent(request *DescribeParamsEventRequest) (response *DescribeParamsEventResponse, err error) {
    return c.DescribeParamsEventWithContext(context.Background(), request)
}

// DescribeParamsEvent
// 本接口（DescribeParamsEvent）用于查询参数修改事件。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeParamsEventWithContext(ctx context.Context, request *DescribeParamsEventRequest) (response *DescribeParamsEventResponse, err error) {
    if request == nil {
        request = NewDescribeParamsEventRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeParamsEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeParamsEventResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProductConfigRequest() (request *DescribeProductConfigRequest) {
    request = &DescribeProductConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeProductConfig")
    
    
    return
}

func NewDescribeProductConfigResponse() (response *DescribeProductConfigResponse) {
    response = &DescribeProductConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProductConfig
// 本接口（DescribeProductConfig）用于查询售卖规格配置。本接口已废弃，推荐使用新接口[DescribeClasses](https://cloud.tencent.com/document/api/409/89019)。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_WHITELISTCONNECTERROR = "FailedOperation.WhitelistConnectError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_BADSPEC = "InvalidParameterValue.BadSpec"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeProductConfig(request *DescribeProductConfigRequest) (response *DescribeProductConfigResponse, err error) {
    return c.DescribeProductConfigWithContext(context.Background(), request)
}

// DescribeProductConfig
// 本接口（DescribeProductConfig）用于查询售卖规格配置。本接口已废弃，推荐使用新接口[DescribeClasses](https://cloud.tencent.com/document/api/409/89019)。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_WHITELISTCONNECTERROR = "FailedOperation.WhitelistConnectError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_BADSPEC = "InvalidParameterValue.BadSpec"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeProductConfigWithContext(ctx context.Context, request *DescribeProductConfigRequest) (response *DescribeProductConfigResponse, err error) {
    if request == nil {
        request = NewDescribeProductConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProductConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProductConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReadOnlyGroupsRequest() (request *DescribeReadOnlyGroupsRequest) {
    request = &DescribeReadOnlyGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeReadOnlyGroups")
    
    
    return
}

func NewDescribeReadOnlyGroupsResponse() (response *DescribeReadOnlyGroupsResponse) {
    response = &DescribeReadOnlyGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeReadOnlyGroups
// 本接口（DescribeReadOnlyGroups）用于查询只读组列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_INTERFACENAMENOTFOUND = "InvalidParameterValue.InterfaceNameNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeReadOnlyGroups(request *DescribeReadOnlyGroupsRequest) (response *DescribeReadOnlyGroupsResponse, err error) {
    return c.DescribeReadOnlyGroupsWithContext(context.Background(), request)
}

// DescribeReadOnlyGroups
// 本接口（DescribeReadOnlyGroups）用于查询只读组列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_INTERFACENAMENOTFOUND = "InvalidParameterValue.InterfaceNameNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeReadOnlyGroupsWithContext(ctx context.Context, request *DescribeReadOnlyGroupsRequest) (response *DescribeReadOnlyGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeReadOnlyGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReadOnlyGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReadOnlyGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRegionsRequest() (request *DescribeRegionsRequest) {
    request = &DescribeRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeRegions")
    
    
    return
}

func NewDescribeRegionsResponse() (response *DescribeRegionsResponse) {
    response = &DescribeRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRegions
// 本接口 (DescribeRegions) 用于查询售卖地域信息。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_WHITELISTCONNECTERROR = "FailedOperation.WhitelistConnectError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INTERFACENAMENOTFOUND = "InvalidParameterValue.InterfaceNameNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_REGIONNOTSUPPORTED = "InvalidParameterValue.RegionNotSupported"
//  INVALIDPARAMETERVALUE_STRUCTPARSEFAILED = "InvalidParameterValue.StructParseFailed"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRegions(request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
    return c.DescribeRegionsWithContext(context.Background(), request)
}

// DescribeRegions
// 本接口 (DescribeRegions) 用于查询售卖地域信息。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_WHITELISTCONNECTERROR = "FailedOperation.WhitelistConnectError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INTERFACENAMENOTFOUND = "InvalidParameterValue.InterfaceNameNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_REGIONNOTSUPPORTED = "InvalidParameterValue.RegionNotSupported"
//  INVALIDPARAMETERVALUE_STRUCTPARSEFAILED = "InvalidParameterValue.StructParseFailed"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRegionsWithContext(ctx context.Context, request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeRegionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRegions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServerlessDBInstancesRequest() (request *DescribeServerlessDBInstancesRequest) {
    request = &DescribeServerlessDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeServerlessDBInstances")
    
    
    return
}

func NewDescribeServerlessDBInstancesResponse() (response *DescribeServerlessDBInstancesResponse) {
    response = &DescribeServerlessDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeServerlessDBInstances
// 用于查询一个或多个serverlessDB实例的详细信息
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DBERROR = "DBError"
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERFACENAMENOTFOUND = "InterfaceNameNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_VPCNOTFOUNDERROR = "InvalidParameter.VpcNotFoundError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_ILLEGALPROJECTID = "InvalidParameterValue.IllegalProjectId"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_VPCRESOURCENOTFOUND = "ResourceUnavailable.VpcResourceNotFound"
//  STRUCTPARSEFAILED = "StructParseFailed"
//  SYSTEMERROR = "SystemError"
//  UNKNOWNERROR = "UnknownError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeServerlessDBInstances(request *DescribeServerlessDBInstancesRequest) (response *DescribeServerlessDBInstancesResponse, err error) {
    return c.DescribeServerlessDBInstancesWithContext(context.Background(), request)
}

// DescribeServerlessDBInstances
// 用于查询一个或多个serverlessDB实例的详细信息
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DBERROR = "DBError"
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERFACENAMENOTFOUND = "InterfaceNameNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_VPCNOTFOUNDERROR = "InvalidParameter.VpcNotFoundError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_ILLEGALPROJECTID = "InvalidParameterValue.IllegalProjectId"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_VPCRESOURCENOTFOUND = "ResourceUnavailable.VpcResourceNotFound"
//  STRUCTPARSEFAILED = "StructParseFailed"
//  SYSTEMERROR = "SystemError"
//  UNKNOWNERROR = "UnknownError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeServerlessDBInstancesWithContext(ctx context.Context, request *DescribeServerlessDBInstancesRequest) (response *DescribeServerlessDBInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeServerlessDBInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServerlessDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServerlessDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSlowQueryAnalysisRequest() (request *DescribeSlowQueryAnalysisRequest) {
    request = &DescribeSlowQueryAnalysisRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeSlowQueryAnalysis")
    
    
    return
}

func NewDescribeSlowQueryAnalysisResponse() (response *DescribeSlowQueryAnalysisResponse) {
    response = &DescribeSlowQueryAnalysisResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSlowQueryAnalysis
// 此接口（DescribeSlowQueryAnalysis）用于统计指定时间范围内的所有慢查询，根据SQL语句抽象参数后，进行聚合分析，并返回同类SQL列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ESQUERYERROR = "FailedOperation.ESQueryError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEEXCEEDERROR = "InvalidParameterValue.ParameterValueExceedError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeSlowQueryAnalysis(request *DescribeSlowQueryAnalysisRequest) (response *DescribeSlowQueryAnalysisResponse, err error) {
    return c.DescribeSlowQueryAnalysisWithContext(context.Background(), request)
}

// DescribeSlowQueryAnalysis
// 此接口（DescribeSlowQueryAnalysis）用于统计指定时间范围内的所有慢查询，根据SQL语句抽象参数后，进行聚合分析，并返回同类SQL列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ESQUERYERROR = "FailedOperation.ESQueryError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEEXCEEDERROR = "InvalidParameterValue.ParameterValueExceedError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeSlowQueryAnalysisWithContext(ctx context.Context, request *DescribeSlowQueryAnalysisRequest) (response *DescribeSlowQueryAnalysisResponse, err error) {
    if request == nil {
        request = NewDescribeSlowQueryAnalysisRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSlowQueryAnalysis require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSlowQueryAnalysisResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSlowQueryListRequest() (request *DescribeSlowQueryListRequest) {
    request = &DescribeSlowQueryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeSlowQueryList")
    
    
    return
}

func NewDescribeSlowQueryListResponse() (response *DescribeSlowQueryListResponse) {
    response = &DescribeSlowQueryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSlowQueryList
// 此接口（DescribeSlowQueryList）用于查询指定时间范围内的所有慢查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEEXCEEDERROR = "InvalidParameterValue.ParameterValueExceedError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeSlowQueryList(request *DescribeSlowQueryListRequest) (response *DescribeSlowQueryListResponse, err error) {
    return c.DescribeSlowQueryListWithContext(context.Background(), request)
}

// DescribeSlowQueryList
// 此接口（DescribeSlowQueryList）用于查询指定时间范围内的所有慢查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEEXCEEDERROR = "InvalidParameterValue.ParameterValueExceedError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeSlowQueryListWithContext(ctx context.Context, request *DescribeSlowQueryListRequest) (response *DescribeSlowQueryListResponse, err error) {
    if request == nil {
        request = NewDescribeSlowQueryListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSlowQueryList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSlowQueryListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZonesRequest() (request *DescribeZonesRequest) {
    request = &DescribeZonesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeZones")
    
    
    return
}

func NewDescribeZonesResponse() (response *DescribeZonesResponse) {
    response = &DescribeZonesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeZones
// 本接口 (DescribeZones) 用于查询支持的可用区信息。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_WHITELISTCONNECTERROR = "FailedOperation.WhitelistConnectError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_REGIONNOTSUPPORTED = "InvalidParameterValue.RegionNotSupported"
//  INVALIDPARAMETERVALUE_STRUCTPARSEFAILED = "InvalidParameterValue.StructParseFailed"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeZones(request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
    return c.DescribeZonesWithContext(context.Background(), request)
}

// DescribeZones
// 本接口 (DescribeZones) 用于查询支持的可用区信息。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_WHITELISTCONNECTERROR = "FailedOperation.WhitelistConnectError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_REGIONNOTSUPPORTED = "InvalidParameterValue.RegionNotSupported"
//  INVALIDPARAMETERVALUE_STRUCTPARSEFAILED = "InvalidParameterValue.StructParseFailed"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeZonesWithContext(ctx context.Context, request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
    if request == nil {
        request = NewDescribeZonesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeZones require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeZonesResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyDBInstanceRequest() (request *DestroyDBInstanceRequest) {
    request = &DestroyDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DestroyDBInstance")
    
    
    return
}

func NewDestroyDBInstanceResponse() (response *DestroyDBInstanceResponse) {
    response = &DestroyDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DestroyDBInstance
// 本接口 (DestroyDBInstance) 用于彻底销毁指定DBInstanceId对应的实例，销毁后实例数据将彻底删除，无法找回，调用前请仔细确认要操作的实例。只能销毁隔离中的实例。
//
// 可能返回的错误码:
//  ACCOUNTNOTEXIST = "AccountNotExist"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CMQRESPONSEERROR = "FailedOperation.CMQResponseError"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DELETERESOURCEPROJECTTAGERROR = "FailedOperation.DeleteResourceProjectTagError"
//  FAILEDOPERATION_DELETERESOURCESTOTAGERROR = "FailedOperation.DeleteResourcesToTagError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_LIMITOPERATION = "FailedOperation.LimitOperation"
//  FAILEDOPERATION_OPERATEFREQUENCYLIMITEDERROR = "FailedOperation.OperateFrequencyLimitedError"
//  FAILEDOPERATION_POSTPAIDUNBLOCKERROR = "FailedOperation.PostPaidUnblockError"
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  FAILEDOPERATION_TRADEACCESSERROR = "FailedOperation.TradeAccessError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_FLOWERROR = "InternalError.FlowError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDACCOUNTSTATUS = "InvalidAccountStatus"
//  INVALIDINSTANCESTATUS = "InvalidInstanceStatus"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCECHARGETYPE = "InvalidParameterValue.IllegalInstanceChargeType"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_STRUCTPARSEFAILED = "InvalidParameterValue.StructParseFailed"
//  LIMITOPERATION = "LimitOperation"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
//  STRUCTPARSEFAILED = "StructParseFailed"
//  SYSTEMERROR = "SystemError"
//  UNKNOWNERROR = "UnknownError"
func (c *Client) DestroyDBInstance(request *DestroyDBInstanceRequest) (response *DestroyDBInstanceResponse, err error) {
    return c.DestroyDBInstanceWithContext(context.Background(), request)
}

// DestroyDBInstance
// 本接口 (DestroyDBInstance) 用于彻底销毁指定DBInstanceId对应的实例，销毁后实例数据将彻底删除，无法找回，调用前请仔细确认要操作的实例。只能销毁隔离中的实例。
//
// 可能返回的错误码:
//  ACCOUNTNOTEXIST = "AccountNotExist"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CMQRESPONSEERROR = "FailedOperation.CMQResponseError"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DELETERESOURCEPROJECTTAGERROR = "FailedOperation.DeleteResourceProjectTagError"
//  FAILEDOPERATION_DELETERESOURCESTOTAGERROR = "FailedOperation.DeleteResourcesToTagError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_LIMITOPERATION = "FailedOperation.LimitOperation"
//  FAILEDOPERATION_OPERATEFREQUENCYLIMITEDERROR = "FailedOperation.OperateFrequencyLimitedError"
//  FAILEDOPERATION_POSTPAIDUNBLOCKERROR = "FailedOperation.PostPaidUnblockError"
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  FAILEDOPERATION_TRADEACCESSERROR = "FailedOperation.TradeAccessError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_FLOWERROR = "InternalError.FlowError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDACCOUNTSTATUS = "InvalidAccountStatus"
//  INVALIDINSTANCESTATUS = "InvalidInstanceStatus"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCECHARGETYPE = "InvalidParameterValue.IllegalInstanceChargeType"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_STRUCTPARSEFAILED = "InvalidParameterValue.StructParseFailed"
//  LIMITOPERATION = "LimitOperation"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
//  STRUCTPARSEFAILED = "StructParseFailed"
//  SYSTEMERROR = "SystemError"
//  UNKNOWNERROR = "UnknownError"
func (c *Client) DestroyDBInstanceWithContext(ctx context.Context, request *DestroyDBInstanceRequest) (response *DestroyDBInstanceResponse, err error) {
    if request == nil {
        request = NewDestroyDBInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DestroyDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDestroyDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDisIsolateDBInstancesRequest() (request *DisIsolateDBInstancesRequest) {
    request = &DisIsolateDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "DisIsolateDBInstances")
    
    
    return
}

func NewDisIsolateDBInstancesResponse() (response *DisIsolateDBInstancesResponse) {
    response = &DisIsolateDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisIsolateDBInstances
// 本接口（DisIsolateDBInstances）用于解隔离实例
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_CAMCHECKRESOURCEFAILED = "FailedOperation.CamCheckResourceFailed"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_INVALIDTRADEOPERATE = "FailedOperation.InvalidTradeOperate"
//  FAILEDOPERATION_QUERYTRADESTATUSERROR = "FailedOperation.QueryTradeStatusError"
//  FAILEDOPERATION_TRADEACCESSERROR = "FailedOperation.TradeAccessError"
//  FAILEDOPERATION_TRADECREATEERROR = "FailedOperation.TradeCreateError"
//  FAILEDOPERATION_TRADEPAYORDERSERROR = "FailedOperation.TradePayOrdersError"
//  FAILEDOPERATION_TRADEQUERYORDERERROR = "FailedOperation.TradeQueryOrderError"
//  FAILEDOPERATION_TRADEQUERYPRICEERROR = "FailedOperation.TradeQueryPriceError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDINSTANCESTATUS = "InvalidInstanceStatus"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETER_TRADEACCESSDENIEDERROR = "InvalidParameter.TradeAccessDeniedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERCHARACTERLIMITERROR = "InvalidParameterValue.ParameterCharacterLimitError"
//  INVALIDPARAMETERVALUE_PARAMETERLENGTHLIMITERROR = "InvalidParameterValue.ParameterLengthLimitError"
//  INVALIDPARAMETERVALUE_STRUCTPARSEFAILED = "InvalidParameterValue.StructParseFailed"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITOPERATION = "LimitOperation"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DisIsolateDBInstances(request *DisIsolateDBInstancesRequest) (response *DisIsolateDBInstancesResponse, err error) {
    return c.DisIsolateDBInstancesWithContext(context.Background(), request)
}

// DisIsolateDBInstances
// 本接口（DisIsolateDBInstances）用于解隔离实例
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_CAMCHECKRESOURCEFAILED = "FailedOperation.CamCheckResourceFailed"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_INVALIDTRADEOPERATE = "FailedOperation.InvalidTradeOperate"
//  FAILEDOPERATION_QUERYTRADESTATUSERROR = "FailedOperation.QueryTradeStatusError"
//  FAILEDOPERATION_TRADEACCESSERROR = "FailedOperation.TradeAccessError"
//  FAILEDOPERATION_TRADECREATEERROR = "FailedOperation.TradeCreateError"
//  FAILEDOPERATION_TRADEPAYORDERSERROR = "FailedOperation.TradePayOrdersError"
//  FAILEDOPERATION_TRADEQUERYORDERERROR = "FailedOperation.TradeQueryOrderError"
//  FAILEDOPERATION_TRADEQUERYPRICEERROR = "FailedOperation.TradeQueryPriceError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDINSTANCESTATUS = "InvalidInstanceStatus"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETER_TRADEACCESSDENIEDERROR = "InvalidParameter.TradeAccessDeniedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERCHARACTERLIMITERROR = "InvalidParameterValue.ParameterCharacterLimitError"
//  INVALIDPARAMETERVALUE_PARAMETERLENGTHLIMITERROR = "InvalidParameterValue.ParameterLengthLimitError"
//  INVALIDPARAMETERVALUE_STRUCTPARSEFAILED = "InvalidParameterValue.StructParseFailed"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITOPERATION = "LimitOperation"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DisIsolateDBInstancesWithContext(ctx context.Context, request *DisIsolateDBInstancesRequest) (response *DisIsolateDBInstancesResponse, err error) {
    if request == nil {
        request = NewDisIsolateDBInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisIsolateDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisIsolateDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewInitDBInstancesRequest() (request *InitDBInstancesRequest) {
    request = &InitDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "InitDBInstances")
    
    
    return
}

func NewInitDBInstancesResponse() (response *InitDBInstancesResponse) {
    response = &InitDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InitDBInstances
// 本接口（InitDBInstances）用于初始化云数据库PostgreSQL实例。本接口已废弃，推荐使用接口[CreateInstances](https://cloud.tencent.com/document/api/409/56107)替代。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_CHARSETNOTFOUNDERROR = "InvalidParameterValue.CharsetNotFoundError"
//  INVALIDPARAMETERVALUE_DECRYPTPASSWORDFAILED = "InvalidParameterValue.DecryptPasswordFailed"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTERROR = "InvalidParameterValue.InvalidAccountError"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTFORMAT = "InvalidParameterValue.InvalidAccountFormat"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAMEERROR = "InvalidParameterValue.InvalidAccountNameError"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAMEFORMATERROR = "InvalidParameterValue.InvalidAccountNameFormatError"
//  INVALIDPARAMETERVALUE_INVALIDCHARSET = "InvalidParameterValue.InvalidCharset"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORDFORMAT = "InvalidParameterValue.InvalidPasswordFormat"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORDLENGTHERROR = "InvalidParameterValue.InvalidPasswordLengthError"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORDVALUEERROR = "InvalidParameterValue.InvalidPasswordValueError"
//  INVALIDPARAMETERVALUE_PARAMETERCHARACTERLIMITERROR = "InvalidParameterValue.ParameterCharacterLimitError"
//  INVALIDPARAMETERVALUE_PARAMETERCHARACTERPRELIMITERROR = "InvalidParameterValue.ParameterCharacterPreLimitError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
func (c *Client) InitDBInstances(request *InitDBInstancesRequest) (response *InitDBInstancesResponse, err error) {
    return c.InitDBInstancesWithContext(context.Background(), request)
}

// InitDBInstances
// 本接口（InitDBInstances）用于初始化云数据库PostgreSQL实例。本接口已废弃，推荐使用接口[CreateInstances](https://cloud.tencent.com/document/api/409/56107)替代。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_CHARSETNOTFOUNDERROR = "InvalidParameterValue.CharsetNotFoundError"
//  INVALIDPARAMETERVALUE_DECRYPTPASSWORDFAILED = "InvalidParameterValue.DecryptPasswordFailed"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTERROR = "InvalidParameterValue.InvalidAccountError"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTFORMAT = "InvalidParameterValue.InvalidAccountFormat"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAMEERROR = "InvalidParameterValue.InvalidAccountNameError"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAMEFORMATERROR = "InvalidParameterValue.InvalidAccountNameFormatError"
//  INVALIDPARAMETERVALUE_INVALIDCHARSET = "InvalidParameterValue.InvalidCharset"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORDFORMAT = "InvalidParameterValue.InvalidPasswordFormat"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORDLENGTHERROR = "InvalidParameterValue.InvalidPasswordLengthError"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORDVALUEERROR = "InvalidParameterValue.InvalidPasswordValueError"
//  INVALIDPARAMETERVALUE_PARAMETERCHARACTERLIMITERROR = "InvalidParameterValue.ParameterCharacterLimitError"
//  INVALIDPARAMETERVALUE_PARAMETERCHARACTERPRELIMITERROR = "InvalidParameterValue.ParameterCharacterPreLimitError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
func (c *Client) InitDBInstancesWithContext(ctx context.Context, request *InitDBInstancesRequest) (response *InitDBInstancesResponse, err error) {
    if request == nil {
        request = NewInitDBInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InitDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewInitDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceCreateDBInstancesRequest() (request *InquiryPriceCreateDBInstancesRequest) {
    request = &InquiryPriceCreateDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "InquiryPriceCreateDBInstances")
    
    
    return
}

func NewInquiryPriceCreateDBInstancesResponse() (response *InquiryPriceCreateDBInstancesResponse) {
    response = &InquiryPriceCreateDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquiryPriceCreateDBInstances
// 本接口 (InquiryPriceCreateDBInstances) 用于查询购买实例的价格信息。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  FAILEDOPERATION_QUERYSPECBYSPECCODEERROR = "FailedOperation.QuerySpecBySpecCodeError"
//  FAILEDOPERATION_TRADEQUERYPRICEERROR = "FailedOperation.TradeQueryPriceError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_BADSPEC = "InvalidParameterValue.BadSpec"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCECHARGETYPE = "InvalidParameterValue.IllegalInstanceChargeType"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_SPECNOTRECOGNIZEDERROR = "InvalidParameterValue.SpecNotRecognizedError"
//  INVALIDPID = "InvalidPid"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_PAYMODEERROR = "OperationDenied.PayModeError"
func (c *Client) InquiryPriceCreateDBInstances(request *InquiryPriceCreateDBInstancesRequest) (response *InquiryPriceCreateDBInstancesResponse, err error) {
    return c.InquiryPriceCreateDBInstancesWithContext(context.Background(), request)
}

// InquiryPriceCreateDBInstances
// 本接口 (InquiryPriceCreateDBInstances) 用于查询购买实例的价格信息。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  FAILEDOPERATION_QUERYSPECBYSPECCODEERROR = "FailedOperation.QuerySpecBySpecCodeError"
//  FAILEDOPERATION_TRADEQUERYPRICEERROR = "FailedOperation.TradeQueryPriceError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_BADSPEC = "InvalidParameterValue.BadSpec"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCECHARGETYPE = "InvalidParameterValue.IllegalInstanceChargeType"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_SPECNOTRECOGNIZEDERROR = "InvalidParameterValue.SpecNotRecognizedError"
//  INVALIDPID = "InvalidPid"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_PAYMODEERROR = "OperationDenied.PayModeError"
func (c *Client) InquiryPriceCreateDBInstancesWithContext(ctx context.Context, request *InquiryPriceCreateDBInstancesRequest) (response *InquiryPriceCreateDBInstancesResponse, err error) {
    if request == nil {
        request = NewInquiryPriceCreateDBInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquiryPriceCreateDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquiryPriceCreateDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceRenewDBInstanceRequest() (request *InquiryPriceRenewDBInstanceRequest) {
    request = &InquiryPriceRenewDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "InquiryPriceRenewDBInstance")
    
    
    return
}

func NewInquiryPriceRenewDBInstanceResponse() (response *InquiryPriceRenewDBInstanceResponse) {
    response = &InquiryPriceRenewDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquiryPriceRenewDBInstance
// 本接口（InquiryPriceRenewDBInstance）用于查询续费实例的价格。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  FAILEDOPERATION_TRADEQUERYPRICEERROR = "FailedOperation.TradeQueryPriceError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_BADSPEC = "InvalidParameterValue.BadSpec"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCECHARGETYPE = "InvalidParameterValue.IllegalInstanceChargeType"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  OPERATIONDENIED_POSTPAIDPAYMODEERROR = "OperationDenied.PostPaidPayModeError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) InquiryPriceRenewDBInstance(request *InquiryPriceRenewDBInstanceRequest) (response *InquiryPriceRenewDBInstanceResponse, err error) {
    return c.InquiryPriceRenewDBInstanceWithContext(context.Background(), request)
}

// InquiryPriceRenewDBInstance
// 本接口（InquiryPriceRenewDBInstance）用于查询续费实例的价格。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  FAILEDOPERATION_TRADEQUERYPRICEERROR = "FailedOperation.TradeQueryPriceError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_BADSPEC = "InvalidParameterValue.BadSpec"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCECHARGETYPE = "InvalidParameterValue.IllegalInstanceChargeType"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  OPERATIONDENIED_POSTPAIDPAYMODEERROR = "OperationDenied.PostPaidPayModeError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) InquiryPriceRenewDBInstanceWithContext(ctx context.Context, request *InquiryPriceRenewDBInstanceRequest) (response *InquiryPriceRenewDBInstanceResponse, err error) {
    if request == nil {
        request = NewInquiryPriceRenewDBInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquiryPriceRenewDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquiryPriceRenewDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceUpgradeDBInstanceRequest() (request *InquiryPriceUpgradeDBInstanceRequest) {
    request = &InquiryPriceUpgradeDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "InquiryPriceUpgradeDBInstance")
    
    
    return
}

func NewInquiryPriceUpgradeDBInstanceResponse() (response *InquiryPriceUpgradeDBInstanceResponse) {
    response = &InquiryPriceUpgradeDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquiryPriceUpgradeDBInstance
// 本接口（InquiryPriceUpgradeDBInstance）用于查询升级实例的价格。只支持按量计费实例。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_CDBCGWCONNECTERROR = "FailedOperation.CdbCgwConnectError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_GETPURCHASERECORDFAILED = "FailedOperation.GetPurchaseRecordFailed"
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  FAILEDOPERATION_TRADEQUERYPRICEERROR = "FailedOperation.TradeQueryPriceError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BADSPEC = "InvalidParameterValue.BadSpec"
//  INVALIDPARAMETERVALUE_ILLEGALCHARGETYPE = "InvalidParameterValue.IllegalChargeType"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCECHARGETYPE = "InvalidParameterValue.IllegalInstanceChargeType"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEVOLUME = "InvalidParameterValue.InvalidInstanceVolume"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
func (c *Client) InquiryPriceUpgradeDBInstance(request *InquiryPriceUpgradeDBInstanceRequest) (response *InquiryPriceUpgradeDBInstanceResponse, err error) {
    return c.InquiryPriceUpgradeDBInstanceWithContext(context.Background(), request)
}

// InquiryPriceUpgradeDBInstance
// 本接口（InquiryPriceUpgradeDBInstance）用于查询升级实例的价格。只支持按量计费实例。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_CDBCGWCONNECTERROR = "FailedOperation.CdbCgwConnectError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_GETPURCHASERECORDFAILED = "FailedOperation.GetPurchaseRecordFailed"
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  FAILEDOPERATION_TRADEQUERYPRICEERROR = "FailedOperation.TradeQueryPriceError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BADSPEC = "InvalidParameterValue.BadSpec"
//  INVALIDPARAMETERVALUE_ILLEGALCHARGETYPE = "InvalidParameterValue.IllegalChargeType"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCECHARGETYPE = "InvalidParameterValue.IllegalInstanceChargeType"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEVOLUME = "InvalidParameterValue.InvalidInstanceVolume"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
func (c *Client) InquiryPriceUpgradeDBInstanceWithContext(ctx context.Context, request *InquiryPriceUpgradeDBInstanceRequest) (response *InquiryPriceUpgradeDBInstanceResponse, err error) {
    if request == nil {
        request = NewInquiryPriceUpgradeDBInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquiryPriceUpgradeDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquiryPriceUpgradeDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewIsolateDBInstancesRequest() (request *IsolateDBInstancesRequest) {
    request = &IsolateDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "IsolateDBInstances")
    
    
    return
}

func NewIsolateDBInstancesResponse() (response *IsolateDBInstancesResponse) {
    response = &IsolateDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// IsolateDBInstances
// 本接口（IsolateDBInstances）用于隔离实例。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DBERROR = "DBError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_GETINSTANCEBYRESOURCEIDERROR = "FailedOperation.GetInstanceByResourceIdError"
//  FAILEDOPERATION_INVALIDACCOUNTSTATUS = "FailedOperation.InvalidAccountStatus"
//  FAILEDOPERATION_OPERATEFREQUENCYLIMITEDERROR = "FailedOperation.OperateFrequencyLimitedError"
//  FAILEDOPERATION_QUERYTRADESTATUSERROR = "FailedOperation.QueryTradeStatusError"
//  FAILEDOPERATION_TRADEACCESSERROR = "FailedOperation.TradeAccessError"
//  FAILEDOPERATION_TRADECREATEERROR = "FailedOperation.TradeCreateError"
//  FAILEDOPERATION_TRADEPAYORDERSERROR = "FailedOperation.TradePayOrdersError"
//  FAILEDOPERATION_TRADEQUERYORDERERROR = "FailedOperation.TradeQueryOrderError"
//  FAILEDOPERATION_TRADEQUERYPRICEERROR = "FailedOperation.TradeQueryPriceError"
//  INSTANCENOTEXIST = "InstanceNotExist"
//  INTERNALERROR_INSTANCEDATAERROR = "InternalError.InstanceDataError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDINSTANCESTATUS = "InvalidInstanceStatus"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TRADEACCESSDENIEDERROR = "InvalidParameter.TradeAccessDeniedError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITOPERATION = "LimitOperation"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_DTSINSTANCESTATUSERROR = "OperationDenied.DTSInstanceStatusError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  REGIONNOTSUPPORTED = "RegionNotSupported"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) IsolateDBInstances(request *IsolateDBInstancesRequest) (response *IsolateDBInstancesResponse, err error) {
    return c.IsolateDBInstancesWithContext(context.Background(), request)
}

// IsolateDBInstances
// 本接口（IsolateDBInstances）用于隔离实例。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DBERROR = "DBError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_GETINSTANCEBYRESOURCEIDERROR = "FailedOperation.GetInstanceByResourceIdError"
//  FAILEDOPERATION_INVALIDACCOUNTSTATUS = "FailedOperation.InvalidAccountStatus"
//  FAILEDOPERATION_OPERATEFREQUENCYLIMITEDERROR = "FailedOperation.OperateFrequencyLimitedError"
//  FAILEDOPERATION_QUERYTRADESTATUSERROR = "FailedOperation.QueryTradeStatusError"
//  FAILEDOPERATION_TRADEACCESSERROR = "FailedOperation.TradeAccessError"
//  FAILEDOPERATION_TRADECREATEERROR = "FailedOperation.TradeCreateError"
//  FAILEDOPERATION_TRADEPAYORDERSERROR = "FailedOperation.TradePayOrdersError"
//  FAILEDOPERATION_TRADEQUERYORDERERROR = "FailedOperation.TradeQueryOrderError"
//  FAILEDOPERATION_TRADEQUERYPRICEERROR = "FailedOperation.TradeQueryPriceError"
//  INSTANCENOTEXIST = "InstanceNotExist"
//  INTERNALERROR_INSTANCEDATAERROR = "InternalError.InstanceDataError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDINSTANCESTATUS = "InvalidInstanceStatus"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TRADEACCESSDENIEDERROR = "InvalidParameter.TradeAccessDeniedError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITOPERATION = "LimitOperation"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_DTSINSTANCESTATUSERROR = "OperationDenied.DTSInstanceStatusError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  REGIONNOTSUPPORTED = "RegionNotSupported"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) IsolateDBInstancesWithContext(ctx context.Context, request *IsolateDBInstancesRequest) (response *IsolateDBInstancesResponse, err error) {
    if request == nil {
        request = NewIsolateDBInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("IsolateDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewIsolateDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewLockAccountRequest() (request *LockAccountRequest) {
    request = &LockAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "LockAccount")
    
    
    return
}

func NewLockAccountResponse() (response *LockAccountResponse) {
    response = &LockAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// LockAccount
// 此接口用于锁定数据库账号，锁定后账号当前连接会断开，并且无法建立新连接。
//
// 可能返回的错误码:
//  DBERROR = "DBError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_PRECHECKERROR = "FailedOperation.PreCheckError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAMEERROR = "InvalidParameterValue.InvalidAccountNameError"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAMEFORMATERROR = "InvalidParameterValue.InvalidAccountNameFormatError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) LockAccount(request *LockAccountRequest) (response *LockAccountResponse, err error) {
    return c.LockAccountWithContext(context.Background(), request)
}

// LockAccount
// 此接口用于锁定数据库账号，锁定后账号当前连接会断开，并且无法建立新连接。
//
// 可能返回的错误码:
//  DBERROR = "DBError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_PRECHECKERROR = "FailedOperation.PreCheckError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAMEERROR = "InvalidParameterValue.InvalidAccountNameError"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAMEFORMATERROR = "InvalidParameterValue.InvalidAccountNameFormatError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) LockAccountWithContext(ctx context.Context, request *LockAccountRequest) (response *LockAccountResponse, err error) {
    if request == nil {
        request = NewLockAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("LockAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewLockAccountResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccountPrivilegesRequest() (request *ModifyAccountPrivilegesRequest) {
    request = &ModifyAccountPrivilegesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "ModifyAccountPrivileges")
    
    
    return
}

func NewModifyAccountPrivilegesResponse() (response *ModifyAccountPrivilegesResponse) {
    response = &ModifyAccountPrivilegesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAccountPrivileges
// 修改某账号对某数据库对象的权限、修改账号类型。
//
// 可能返回的错误码:
//  DBERROR = "DBError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_PRECHECKERROR = "FailedOperation.PreCheckError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAMEERROR = "InvalidParameterValue.InvalidAccountNameError"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAMEFORMATERROR = "InvalidParameterValue.InvalidAccountNameFormatError"
//  INVALIDPARAMETERVALUE_INVALIDDATABASEOBJECTNAME = "InvalidParameterValue.InvalidDatabaseObjectName"
//  INVALIDPARAMETERVALUE_PARAMETEROUTOFRANGE = "InvalidParameterValue.ParameterOutOfRange"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  OPERATIONDENIED_KERNELVERSIONERROR = "OperationDenied.KernelVersionError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) ModifyAccountPrivileges(request *ModifyAccountPrivilegesRequest) (response *ModifyAccountPrivilegesResponse, err error) {
    return c.ModifyAccountPrivilegesWithContext(context.Background(), request)
}

// ModifyAccountPrivileges
// 修改某账号对某数据库对象的权限、修改账号类型。
//
// 可能返回的错误码:
//  DBERROR = "DBError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_PRECHECKERROR = "FailedOperation.PreCheckError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAMEERROR = "InvalidParameterValue.InvalidAccountNameError"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAMEFORMATERROR = "InvalidParameterValue.InvalidAccountNameFormatError"
//  INVALIDPARAMETERVALUE_INVALIDDATABASEOBJECTNAME = "InvalidParameterValue.InvalidDatabaseObjectName"
//  INVALIDPARAMETERVALUE_PARAMETEROUTOFRANGE = "InvalidParameterValue.ParameterOutOfRange"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  OPERATIONDENIED_KERNELVERSIONERROR = "OperationDenied.KernelVersionError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) ModifyAccountPrivilegesWithContext(ctx context.Context, request *ModifyAccountPrivilegesRequest) (response *ModifyAccountPrivilegesResponse, err error) {
    if request == nil {
        request = NewModifyAccountPrivilegesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccountPrivileges require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAccountPrivilegesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccountRemarkRequest() (request *ModifyAccountRemarkRequest) {
    request = &ModifyAccountRemarkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "ModifyAccountRemark")
    
    
    return
}

func NewModifyAccountRemarkResponse() (response *ModifyAccountRemarkResponse) {
    response = &ModifyAccountRemarkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAccountRemark
// 本接口（ModifyAccountRemark）用于修改账号备注。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTFORMAT = "InvalidParameterValue.InvalidAccountFormat"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAMEERROR = "InvalidParameterValue.InvalidAccountNameError"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAMEFORMATERROR = "InvalidParameterValue.InvalidAccountNameFormatError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERCHARACTERERROR = "InvalidParameterValue.ParameterCharacterError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_PARAMETERLENGTHLIMITERROR = "InvalidParameterValue.ParameterLengthLimitError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) ModifyAccountRemark(request *ModifyAccountRemarkRequest) (response *ModifyAccountRemarkResponse, err error) {
    return c.ModifyAccountRemarkWithContext(context.Background(), request)
}

// ModifyAccountRemark
// 本接口（ModifyAccountRemark）用于修改账号备注。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTFORMAT = "InvalidParameterValue.InvalidAccountFormat"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAMEERROR = "InvalidParameterValue.InvalidAccountNameError"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAMEFORMATERROR = "InvalidParameterValue.InvalidAccountNameFormatError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERCHARACTERERROR = "InvalidParameterValue.ParameterCharacterError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_PARAMETERLENGTHLIMITERROR = "InvalidParameterValue.ParameterLengthLimitError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) ModifyAccountRemarkWithContext(ctx context.Context, request *ModifyAccountRemarkRequest) (response *ModifyAccountRemarkResponse, err error) {
    if request == nil {
        request = NewModifyAccountRemarkRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccountRemark require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAccountRemarkResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBackupDownloadRestrictionRequest() (request *ModifyBackupDownloadRestrictionRequest) {
    request = &ModifyBackupDownloadRestrictionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "ModifyBackupDownloadRestriction")
    
    
    return
}

func NewModifyBackupDownloadRestrictionResponse() (response *ModifyBackupDownloadRestrictionResponse) {
    response = &ModifyBackupDownloadRestrictionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyBackupDownloadRestriction
// 本接口（ModifyBackupDownloadRestriction）用于修改备份文件下载限制。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETER_VPCNOTFOUNDERROR = "InvalidParameter.VpcNotFoundError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
func (c *Client) ModifyBackupDownloadRestriction(request *ModifyBackupDownloadRestrictionRequest) (response *ModifyBackupDownloadRestrictionResponse, err error) {
    return c.ModifyBackupDownloadRestrictionWithContext(context.Background(), request)
}

// ModifyBackupDownloadRestriction
// 本接口（ModifyBackupDownloadRestriction）用于修改备份文件下载限制。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETER_VPCNOTFOUNDERROR = "InvalidParameter.VpcNotFoundError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
func (c *Client) ModifyBackupDownloadRestrictionWithContext(ctx context.Context, request *ModifyBackupDownloadRestrictionRequest) (response *ModifyBackupDownloadRestrictionResponse, err error) {
    if request == nil {
        request = NewModifyBackupDownloadRestrictionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBackupDownloadRestriction require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBackupDownloadRestrictionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBackupPlanRequest() (request *ModifyBackupPlanRequest) {
    request = &ModifyBackupPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "ModifyBackupPlan")
    
    
    return
}

func NewModifyBackupPlanResponse() (response *ModifyBackupPlanResponse) {
    response = &ModifyBackupPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyBackupPlan
// 本接口 (ModifyBackupPlan) 用于实例备份计划的修改，默认是在每天的凌晨开始全量备份，备份保留时长是7天。可以根据此接口指定时间进行实例的备份。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) ModifyBackupPlan(request *ModifyBackupPlanRequest) (response *ModifyBackupPlanResponse, err error) {
    return c.ModifyBackupPlanWithContext(context.Background(), request)
}

// ModifyBackupPlan
// 本接口 (ModifyBackupPlan) 用于实例备份计划的修改，默认是在每天的凌晨开始全量备份，备份保留时长是7天。可以根据此接口指定时间进行实例的备份。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) ModifyBackupPlanWithContext(ctx context.Context, request *ModifyBackupPlanRequest) (response *ModifyBackupPlanResponse, err error) {
    if request == nil {
        request = NewModifyBackupPlanRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBackupPlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBackupPlanResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBaseBackupExpireTimeRequest() (request *ModifyBaseBackupExpireTimeRequest) {
    request = &ModifyBaseBackupExpireTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "ModifyBaseBackupExpireTime")
    
    
    return
}

func NewModifyBaseBackupExpireTimeResponse() (response *ModifyBaseBackupExpireTimeResponse) {
    response = &ModifyBaseBackupExpireTimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyBaseBackupExpireTime
// 本接口（ModifyBaseBackupExpireTime）用于修改实例指定数据备份的过期时间。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) ModifyBaseBackupExpireTime(request *ModifyBaseBackupExpireTimeRequest) (response *ModifyBaseBackupExpireTimeResponse, err error) {
    return c.ModifyBaseBackupExpireTimeWithContext(context.Background(), request)
}

// ModifyBaseBackupExpireTime
// 本接口（ModifyBaseBackupExpireTime）用于修改实例指定数据备份的过期时间。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) ModifyBaseBackupExpireTimeWithContext(ctx context.Context, request *ModifyBaseBackupExpireTimeRequest) (response *ModifyBaseBackupExpireTimeResponse, err error) {
    if request == nil {
        request = NewModifyBaseBackupExpireTimeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBaseBackupExpireTime require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBaseBackupExpireTimeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceChargeTypeRequest() (request *ModifyDBInstanceChargeTypeRequest) {
    request = &ModifyDBInstanceChargeTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "ModifyDBInstanceChargeType")
    
    
    return
}

func NewModifyDBInstanceChargeTypeResponse() (response *ModifyDBInstanceChargeTypeResponse) {
    response = &ModifyDBInstanceChargeTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDBInstanceChargeType
// 支持实例的计费类型转换（目前仅支持按量计费转包年包月）
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
func (c *Client) ModifyDBInstanceChargeType(request *ModifyDBInstanceChargeTypeRequest) (response *ModifyDBInstanceChargeTypeResponse, err error) {
    return c.ModifyDBInstanceChargeTypeWithContext(context.Background(), request)
}

// ModifyDBInstanceChargeType
// 支持实例的计费类型转换（目前仅支持按量计费转包年包月）
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
func (c *Client) ModifyDBInstanceChargeTypeWithContext(ctx context.Context, request *ModifyDBInstanceChargeTypeRequest) (response *ModifyDBInstanceChargeTypeResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceChargeTypeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBInstanceChargeType require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBInstanceChargeTypeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceDeploymentRequest() (request *ModifyDBInstanceDeploymentRequest) {
    request = &ModifyDBInstanceDeploymentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "ModifyDBInstanceDeployment")
    
    
    return
}

func NewModifyDBInstanceDeploymentResponse() (response *ModifyDBInstanceDeploymentResponse) {
    response = &ModifyDBInstanceDeploymentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDBInstanceDeployment
// 本接口（ModifyDBInstanceDeployment）用于修改节点可用区部署方式，仅支持主实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
func (c *Client) ModifyDBInstanceDeployment(request *ModifyDBInstanceDeploymentRequest) (response *ModifyDBInstanceDeploymentResponse, err error) {
    return c.ModifyDBInstanceDeploymentWithContext(context.Background(), request)
}

// ModifyDBInstanceDeployment
// 本接口（ModifyDBInstanceDeployment）用于修改节点可用区部署方式，仅支持主实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
func (c *Client) ModifyDBInstanceDeploymentWithContext(ctx context.Context, request *ModifyDBInstanceDeploymentRequest) (response *ModifyDBInstanceDeploymentResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceDeploymentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBInstanceDeployment require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBInstanceDeploymentResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceHAConfigRequest() (request *ModifyDBInstanceHAConfigRequest) {
    request = &ModifyDBInstanceHAConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "ModifyDBInstanceHAConfig")
    
    
    return
}

func NewModifyDBInstanceHAConfigResponse() (response *ModifyDBInstanceHAConfigResponse) {
    response = &ModifyDBInstanceHAConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDBInstanceHAConfig
// 本接口（ModifyDBInstanceHAConfig）用于修改实例HA配置信息。其中HA配置信息包括：
//
// <li>允许备节点切换为主节点的条件配置
//
// <li>半同步实例使用同步复制或异步复制的条件配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETER_PARAMETERSNOTSET = "InvalidParameter.ParametersNotSet"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETEROUTOFRANGE = "InvalidParameterValue.ParameterOutOfRange"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) ModifyDBInstanceHAConfig(request *ModifyDBInstanceHAConfigRequest) (response *ModifyDBInstanceHAConfigResponse, err error) {
    return c.ModifyDBInstanceHAConfigWithContext(context.Background(), request)
}

// ModifyDBInstanceHAConfig
// 本接口（ModifyDBInstanceHAConfig）用于修改实例HA配置信息。其中HA配置信息包括：
//
// <li>允许备节点切换为主节点的条件配置
//
// <li>半同步实例使用同步复制或异步复制的条件配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETER_PARAMETERSNOTSET = "InvalidParameter.ParametersNotSet"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETEROUTOFRANGE = "InvalidParameterValue.ParameterOutOfRange"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) ModifyDBInstanceHAConfigWithContext(ctx context.Context, request *ModifyDBInstanceHAConfigRequest) (response *ModifyDBInstanceHAConfigResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceHAConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBInstanceHAConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBInstanceHAConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceNameRequest() (request *ModifyDBInstanceNameRequest) {
    request = &ModifyDBInstanceNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "ModifyDBInstanceName")
    
    
    return
}

func NewModifyDBInstanceNameResponse() (response *ModifyDBInstanceNameResponse) {
    response = &ModifyDBInstanceNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDBInstanceName
// 本接口（ModifyDBInstanceName）用于修改postgresql实例名字。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_OSSACCESSERROR = "FailedOperation.OssAccessError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTFORMAT = "InvalidParameterValue.InvalidAccountFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERCHARACTERERROR = "InvalidParameterValue.ParameterCharacterError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_PARAMETERLENGTHLIMITERROR = "InvalidParameterValue.ParameterLengthLimitError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) ModifyDBInstanceName(request *ModifyDBInstanceNameRequest) (response *ModifyDBInstanceNameResponse, err error) {
    return c.ModifyDBInstanceNameWithContext(context.Background(), request)
}

// ModifyDBInstanceName
// 本接口（ModifyDBInstanceName）用于修改postgresql实例名字。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_OSSACCESSERROR = "FailedOperation.OssAccessError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTFORMAT = "InvalidParameterValue.InvalidAccountFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERCHARACTERERROR = "InvalidParameterValue.ParameterCharacterError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_PARAMETERLENGTHLIMITERROR = "InvalidParameterValue.ParameterLengthLimitError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) ModifyDBInstanceNameWithContext(ctx context.Context, request *ModifyDBInstanceNameRequest) (response *ModifyDBInstanceNameResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceNameRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBInstanceName require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBInstanceNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceParametersRequest() (request *ModifyDBInstanceParametersRequest) {
    request = &ModifyDBInstanceParametersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "ModifyDBInstanceParameters")
    
    
    return
}

func NewModifyDBInstanceParametersResponse() (response *ModifyDBInstanceParametersResponse) {
    response = &ModifyDBInstanceParametersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDBInstanceParameters
// 本接口 (ModifyDBInstanceParameters) 用于修改实例参数。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) ModifyDBInstanceParameters(request *ModifyDBInstanceParametersRequest) (response *ModifyDBInstanceParametersResponse, err error) {
    return c.ModifyDBInstanceParametersWithContext(context.Background(), request)
}

// ModifyDBInstanceParameters
// 本接口 (ModifyDBInstanceParameters) 用于修改实例参数。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) ModifyDBInstanceParametersWithContext(ctx context.Context, request *ModifyDBInstanceParametersRequest) (response *ModifyDBInstanceParametersResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceParametersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBInstanceParameters require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBInstanceParametersResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceReadOnlyGroupRequest() (request *ModifyDBInstanceReadOnlyGroupRequest) {
    request = &ModifyDBInstanceReadOnlyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "ModifyDBInstanceReadOnlyGroup")
    
    
    return
}

func NewModifyDBInstanceReadOnlyGroupResponse() (response *ModifyDBInstanceReadOnlyGroupResponse) {
    response = &ModifyDBInstanceReadOnlyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDBInstanceReadOnlyGroup
// 本接口（ModifyDBInstanceReadOnlyGroup）用于修改实例所属的只读组
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_ILLEGALROINSTANCENUM = "FailedOperation.IllegalROInstanceNum"
//  FAILEDOPERATION_ROGROUPNOTFOUNDERROR = "FailedOperation.ROGroupNotFoundError"
//  INSTANCENOTEXIST = "InstanceNotExist"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDINSTANCESTATUS = "InvalidInstanceStatus"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDREADONLYGROUPSTATUS = "InvalidParameterValue.InvalidReadOnlyGroupStatus"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_READONLYGROUPNOTEXIST = "InvalidParameterValue.ReadOnlyGroupNotExist"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_ROGROUPSTATUSERROR = "OperationDenied.ROGroupStatusError"
//  OPERATIONDENIED_ROINSTANCECOUNTEXEEDERROR = "OperationDenied.RoInstanceCountExeedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
func (c *Client) ModifyDBInstanceReadOnlyGroup(request *ModifyDBInstanceReadOnlyGroupRequest) (response *ModifyDBInstanceReadOnlyGroupResponse, err error) {
    return c.ModifyDBInstanceReadOnlyGroupWithContext(context.Background(), request)
}

// ModifyDBInstanceReadOnlyGroup
// 本接口（ModifyDBInstanceReadOnlyGroup）用于修改实例所属的只读组
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_ILLEGALROINSTANCENUM = "FailedOperation.IllegalROInstanceNum"
//  FAILEDOPERATION_ROGROUPNOTFOUNDERROR = "FailedOperation.ROGroupNotFoundError"
//  INSTANCENOTEXIST = "InstanceNotExist"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDINSTANCESTATUS = "InvalidInstanceStatus"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDREADONLYGROUPSTATUS = "InvalidParameterValue.InvalidReadOnlyGroupStatus"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_READONLYGROUPNOTEXIST = "InvalidParameterValue.ReadOnlyGroupNotExist"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_ROGROUPSTATUSERROR = "OperationDenied.ROGroupStatusError"
//  OPERATIONDENIED_ROINSTANCECOUNTEXEEDERROR = "OperationDenied.RoInstanceCountExeedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
func (c *Client) ModifyDBInstanceReadOnlyGroupWithContext(ctx context.Context, request *ModifyDBInstanceReadOnlyGroupRequest) (response *ModifyDBInstanceReadOnlyGroupResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceReadOnlyGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBInstanceReadOnlyGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBInstanceReadOnlyGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceSecurityGroupsRequest() (request *ModifyDBInstanceSecurityGroupsRequest) {
    request = &ModifyDBInstanceSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "ModifyDBInstanceSecurityGroups")
    
    
    return
}

func NewModifyDBInstanceSecurityGroupsResponse() (response *ModifyDBInstanceSecurityGroupsResponse) {
    response = &ModifyDBInstanceSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDBInstanceSecurityGroups
// 本接口（ModifyDBInstanceSecurityGroups）用于修改实例安全组。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_ROGROUPNOTFOUNDERROR = "FailedOperation.ROGroupNotFoundError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  OPERATIONDENIED_ROGROUPSTATUSERROR = "OperationDenied.ROGroupStatusError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) ModifyDBInstanceSecurityGroups(request *ModifyDBInstanceSecurityGroupsRequest) (response *ModifyDBInstanceSecurityGroupsResponse, err error) {
    return c.ModifyDBInstanceSecurityGroupsWithContext(context.Background(), request)
}

// ModifyDBInstanceSecurityGroups
// 本接口（ModifyDBInstanceSecurityGroups）用于修改实例安全组。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_ROGROUPNOTFOUNDERROR = "FailedOperation.ROGroupNotFoundError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  OPERATIONDENIED_ROGROUPSTATUSERROR = "OperationDenied.ROGroupStatusError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) ModifyDBInstanceSecurityGroupsWithContext(ctx context.Context, request *ModifyDBInstanceSecurityGroupsRequest) (response *ModifyDBInstanceSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceSecurityGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBInstanceSecurityGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBInstanceSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceSpecRequest() (request *ModifyDBInstanceSpecRequest) {
    request = &ModifyDBInstanceSpecRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "ModifyDBInstanceSpec")
    
    
    return
}

func NewModifyDBInstanceSpecResponse() (response *ModifyDBInstanceSpecResponse) {
    response = &ModifyDBInstanceSpecResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDBInstanceSpec
// 本接口（ModifyDBInstanceSpec）用于修改实例规格，包括内存、磁盘。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_PAYORDERFAILED = "FailedOperation.PayOrderFailed"
//  FAILEDOPERATION_PRECHECKERROR = "FailedOperation.PreCheckError"
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  FAILEDOPERATION_QUERYSPECERROR = "FailedOperation.QuerySpecError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) ModifyDBInstanceSpec(request *ModifyDBInstanceSpecRequest) (response *ModifyDBInstanceSpecResponse, err error) {
    return c.ModifyDBInstanceSpecWithContext(context.Background(), request)
}

// ModifyDBInstanceSpec
// 本接口（ModifyDBInstanceSpec）用于修改实例规格，包括内存、磁盘。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_PAYORDERFAILED = "FailedOperation.PayOrderFailed"
//  FAILEDOPERATION_PRECHECKERROR = "FailedOperation.PreCheckError"
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  FAILEDOPERATION_QUERYSPECERROR = "FailedOperation.QuerySpecError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) ModifyDBInstanceSpecWithContext(ctx context.Context, request *ModifyDBInstanceSpecRequest) (response *ModifyDBInstanceSpecResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceSpecRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBInstanceSpec require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBInstanceSpecResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstancesProjectRequest() (request *ModifyDBInstancesProjectRequest) {
    request = &ModifyDBInstancesProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "ModifyDBInstancesProject")
    
    
    return
}

func NewModifyDBInstancesProjectResponse() (response *ModifyDBInstancesProjectResponse) {
    response = &ModifyDBInstancesProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDBInstancesProject
// 本接口（ModifyDBInstancesProject）用于修改实例所属项目。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_CDBCGWCONNECTERROR = "FailedOperation.CdbCgwConnectError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_UPDATERESOURCEPROJECTTAGVALUEERROR = "FailedOperation.UpdateResourceProjectTagValueError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ILLEGALPROJECTID = "InvalidParameterValue.IllegalProjectId"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) ModifyDBInstancesProject(request *ModifyDBInstancesProjectRequest) (response *ModifyDBInstancesProjectResponse, err error) {
    return c.ModifyDBInstancesProjectWithContext(context.Background(), request)
}

// ModifyDBInstancesProject
// 本接口（ModifyDBInstancesProject）用于修改实例所属项目。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_CDBCGWCONNECTERROR = "FailedOperation.CdbCgwConnectError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_UPDATERESOURCEPROJECTTAGVALUEERROR = "FailedOperation.UpdateResourceProjectTagValueError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ILLEGALPROJECTID = "InvalidParameterValue.IllegalProjectId"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) ModifyDBInstancesProjectWithContext(ctx context.Context, request *ModifyDBInstancesProjectRequest) (response *ModifyDBInstancesProjectResponse, err error) {
    if request == nil {
        request = NewModifyDBInstancesProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBInstancesProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBInstancesProjectResponse()
    err = c.Send(request, response)
    return
}

func NewModifyParameterTemplateRequest() (request *ModifyParameterTemplateRequest) {
    request = &ModifyParameterTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "ModifyParameterTemplate")
    
    
    return
}

func NewModifyParameterTemplateResponse() (response *ModifyParameterTemplateResponse) {
    response = &ModifyParameterTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyParameterTemplate
// 本接口（ModifyParameterTemplate）主要用于修改参数模板名称，描述等配置，也可用于管理参数模板中的参数列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
func (c *Client) ModifyParameterTemplate(request *ModifyParameterTemplateRequest) (response *ModifyParameterTemplateResponse, err error) {
    return c.ModifyParameterTemplateWithContext(context.Background(), request)
}

// ModifyParameterTemplate
// 本接口（ModifyParameterTemplate）主要用于修改参数模板名称，描述等配置，也可用于管理参数模板中的参数列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
func (c *Client) ModifyParameterTemplateWithContext(ctx context.Context, request *ModifyParameterTemplateRequest) (response *ModifyParameterTemplateResponse, err error) {
    if request == nil {
        request = NewModifyParameterTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyParameterTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyParameterTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyReadOnlyGroupConfigRequest() (request *ModifyReadOnlyGroupConfigRequest) {
    request = &ModifyReadOnlyGroupConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "ModifyReadOnlyGroupConfig")
    
    
    return
}

func NewModifyReadOnlyGroupConfigResponse() (response *ModifyReadOnlyGroupConfigResponse) {
    response = &ModifyReadOnlyGroupConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyReadOnlyGroupConfig
// 本接口(ModifyReadOnlyGroupConfig)用于更新只读组配置信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_BASENETWORKACCESSERROR = "FailedOperation.BaseNetworkAccessError"
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DELETEALLROUTE = "FailedOperation.DeleteAllRoute"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_MODIFYROGROUPERROR = "FailedOperation.ModifyROGroupError"
//  FAILEDOPERATION_ROGROUPNOTFOUNDERROR = "FailedOperation.ROGroupNotFoundError"
//  FAILEDOPERATION_SERVICEACCESSERROR = "FailedOperation.ServiceAccessError"
//  FAILEDOPERATION_VPCRESETSERVICEERROR = "FailedOperation.VPCResetServiceError"
//  FAILEDOPERATION_VPCUPDATEROUTEERROR = "FailedOperation.VPCUpdateRouteError"
//  INTERFACENAMENOTFOUND = "InterfaceNameNotFound"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_READONLYGROUPNOTEXIST = "InvalidParameterValue.ReadOnlyGroupNotExist"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) ModifyReadOnlyGroupConfig(request *ModifyReadOnlyGroupConfigRequest) (response *ModifyReadOnlyGroupConfigResponse, err error) {
    return c.ModifyReadOnlyGroupConfigWithContext(context.Background(), request)
}

// ModifyReadOnlyGroupConfig
// 本接口(ModifyReadOnlyGroupConfig)用于更新只读组配置信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_BASENETWORKACCESSERROR = "FailedOperation.BaseNetworkAccessError"
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DELETEALLROUTE = "FailedOperation.DeleteAllRoute"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_MODIFYROGROUPERROR = "FailedOperation.ModifyROGroupError"
//  FAILEDOPERATION_ROGROUPNOTFOUNDERROR = "FailedOperation.ROGroupNotFoundError"
//  FAILEDOPERATION_SERVICEACCESSERROR = "FailedOperation.ServiceAccessError"
//  FAILEDOPERATION_VPCRESETSERVICEERROR = "FailedOperation.VPCResetServiceError"
//  FAILEDOPERATION_VPCUPDATEROUTEERROR = "FailedOperation.VPCUpdateRouteError"
//  INTERFACENAMENOTFOUND = "InterfaceNameNotFound"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_READONLYGROUPNOTEXIST = "InvalidParameterValue.ReadOnlyGroupNotExist"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) ModifyReadOnlyGroupConfigWithContext(ctx context.Context, request *ModifyReadOnlyGroupConfigRequest) (response *ModifyReadOnlyGroupConfigResponse, err error) {
    if request == nil {
        request = NewModifyReadOnlyGroupConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyReadOnlyGroupConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyReadOnlyGroupConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifySwitchTimePeriodRequest() (request *ModifySwitchTimePeriodRequest) {
    request = &ModifySwitchTimePeriodRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "ModifySwitchTimePeriod")
    
    
    return
}

func NewModifySwitchTimePeriodResponse() (response *ModifySwitchTimePeriodResponse) {
    response = &ModifySwitchTimePeriodResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySwitchTimePeriod
// 当升级完成后，对处于等待切换状态下的实例，强制实例立即切换。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_CAMCHECKRESOURCEFAILED = "FailedOperation.CamCheckResourceFailed"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_QUERYSPECERROR = "FailedOperation.QuerySpecError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNKNOWNERROR = "UnknownError"
func (c *Client) ModifySwitchTimePeriod(request *ModifySwitchTimePeriodRequest) (response *ModifySwitchTimePeriodResponse, err error) {
    return c.ModifySwitchTimePeriodWithContext(context.Background(), request)
}

// ModifySwitchTimePeriod
// 当升级完成后，对处于等待切换状态下的实例，强制实例立即切换。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_CAMCHECKRESOURCEFAILED = "FailedOperation.CamCheckResourceFailed"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_QUERYSPECERROR = "FailedOperation.QuerySpecError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNKNOWNERROR = "UnknownError"
func (c *Client) ModifySwitchTimePeriodWithContext(ctx context.Context, request *ModifySwitchTimePeriodRequest) (response *ModifySwitchTimePeriodResponse, err error) {
    if request == nil {
        request = NewModifySwitchTimePeriodRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySwitchTimePeriod require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySwitchTimePeriodResponse()
    err = c.Send(request, response)
    return
}

func NewOpenDBExtranetAccessRequest() (request *OpenDBExtranetAccessRequest) {
    request = &OpenDBExtranetAccessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "OpenDBExtranetAccess")
    
    
    return
}

func NewOpenDBExtranetAccessResponse() (response *OpenDBExtranetAccessResponse) {
    response = &OpenDBExtranetAccessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OpenDBExtranetAccess
// 本接口（OpenDBExtranetAccess）用于开通实例公网地址。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_FLOWERROR = "InternalError.FlowError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCEIPV6NOTSUPPORTEDERROR = "OperationDenied.InstanceIpv6NotSupportedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
func (c *Client) OpenDBExtranetAccess(request *OpenDBExtranetAccessRequest) (response *OpenDBExtranetAccessResponse, err error) {
    return c.OpenDBExtranetAccessWithContext(context.Background(), request)
}

// OpenDBExtranetAccess
// 本接口（OpenDBExtranetAccess）用于开通实例公网地址。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_FLOWERROR = "InternalError.FlowError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCEIPV6NOTSUPPORTEDERROR = "OperationDenied.InstanceIpv6NotSupportedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
func (c *Client) OpenDBExtranetAccessWithContext(ctx context.Context, request *OpenDBExtranetAccessRequest) (response *OpenDBExtranetAccessResponse, err error) {
    if request == nil {
        request = NewOpenDBExtranetAccessRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenDBExtranetAccess require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenDBExtranetAccessResponse()
    err = c.Send(request, response)
    return
}

func NewOpenServerlessDBExtranetAccessRequest() (request *OpenServerlessDBExtranetAccessRequest) {
    request = &OpenServerlessDBExtranetAccessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "OpenServerlessDBExtranetAccess")
    
    
    return
}

func NewOpenServerlessDBExtranetAccessResponse() (response *OpenServerlessDBExtranetAccessResponse) {
    response = &OpenServerlessDBExtranetAccessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OpenServerlessDBExtranetAccess
// 本接口（OpenServerlessDBExtranetAccess）用于开通serverlessDB实例公网地址。
//
// 可能返回的错误码:
//  ACCOUNTNOTEXIST = "AccountNotExist"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DBERROR = "DBError"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_QUERYVPCFALIED = "FailedOperation.QueryVpcFalied"
//  FLOWERROR = "FlowError"
//  INSTANCENOTEXIST = "InstanceNotExist"
//  INTERFACENAMENOTFOUND = "InterfaceNameNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_FLOWERROR = "InternalError.FlowError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDACCOUNTSTATUS = "InvalidAccountStatus"
//  INVALIDINSTANCESTATUS = "InvalidInstanceStatus"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INTERFACENAMENOTFOUND = "InvalidParameterValue.InterfaceNameNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_STRUCTPARSEFAILED = "InvalidParameterValue.StructParseFailed"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
//  RESOURCEUNAVAILABLE_RESOURCENOPERMISSION = "ResourceUnavailable.ResourceNoPermission"
//  RESOURCEUNAVAILABLE_VPCRESOURCENOTFOUND = "ResourceUnavailable.VpcResourceNotFound"
//  STRUCTPARSEFAILED = "StructParseFailed"
//  SYSTEMERROR = "SystemError"
//  UNKNOWNERROR = "UnknownError"
//  VPCERROR = "VpcError"
func (c *Client) OpenServerlessDBExtranetAccess(request *OpenServerlessDBExtranetAccessRequest) (response *OpenServerlessDBExtranetAccessResponse, err error) {
    return c.OpenServerlessDBExtranetAccessWithContext(context.Background(), request)
}

// OpenServerlessDBExtranetAccess
// 本接口（OpenServerlessDBExtranetAccess）用于开通serverlessDB实例公网地址。
//
// 可能返回的错误码:
//  ACCOUNTNOTEXIST = "AccountNotExist"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DBERROR = "DBError"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_QUERYVPCFALIED = "FailedOperation.QueryVpcFalied"
//  FLOWERROR = "FlowError"
//  INSTANCENOTEXIST = "InstanceNotExist"
//  INTERFACENAMENOTFOUND = "InterfaceNameNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_FLOWERROR = "InternalError.FlowError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDACCOUNTSTATUS = "InvalidAccountStatus"
//  INVALIDINSTANCESTATUS = "InvalidInstanceStatus"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INTERFACENAMENOTFOUND = "InvalidParameterValue.InterfaceNameNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_STRUCTPARSEFAILED = "InvalidParameterValue.StructParseFailed"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
//  RESOURCEUNAVAILABLE_RESOURCENOPERMISSION = "ResourceUnavailable.ResourceNoPermission"
//  RESOURCEUNAVAILABLE_VPCRESOURCENOTFOUND = "ResourceUnavailable.VpcResourceNotFound"
//  STRUCTPARSEFAILED = "StructParseFailed"
//  SYSTEMERROR = "SystemError"
//  UNKNOWNERROR = "UnknownError"
//  VPCERROR = "VpcError"
func (c *Client) OpenServerlessDBExtranetAccessWithContext(ctx context.Context, request *OpenServerlessDBExtranetAccessRequest) (response *OpenServerlessDBExtranetAccessResponse, err error) {
    if request == nil {
        request = NewOpenServerlessDBExtranetAccessRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenServerlessDBExtranetAccess require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenServerlessDBExtranetAccessResponse()
    err = c.Send(request, response)
    return
}

func NewRebalanceReadOnlyGroupRequest() (request *RebalanceReadOnlyGroupRequest) {
    request = &RebalanceReadOnlyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "RebalanceReadOnlyGroup")
    
    
    return
}

func NewRebalanceReadOnlyGroupResponse() (response *RebalanceReadOnlyGroupResponse) {
    response = &RebalanceReadOnlyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RebalanceReadOnlyGroup
// 本接口(RebalanceReadOnlyGroup)用于重新均衡 RO 组内实例的负载。注意，RO 组内 RO 实例会有一次数据库连接瞬断，请确保应用程序能重连数据库，谨慎操作。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BASENETWORKACCESSERROR = "FailedOperation.BaseNetworkAccessError"
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DELETEALLROUTE = "FailedOperation.DeleteAllRoute"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_ROGROUPNOTFOUNDERROR = "FailedOperation.ROGroupNotFoundError"
//  FAILEDOPERATION_VPCRESETSERVICEERROR = "FailedOperation.VPCResetServiceError"
//  FAILEDOPERATION_VPCUPDATEROUTEERROR = "FailedOperation.VPCUpdateRouteError"
//  INTERFACENAMENOTFOUND = "InterfaceNameNotFound"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  VPCERROR = "VpcError"
func (c *Client) RebalanceReadOnlyGroup(request *RebalanceReadOnlyGroupRequest) (response *RebalanceReadOnlyGroupResponse, err error) {
    return c.RebalanceReadOnlyGroupWithContext(context.Background(), request)
}

// RebalanceReadOnlyGroup
// 本接口(RebalanceReadOnlyGroup)用于重新均衡 RO 组内实例的负载。注意，RO 组内 RO 实例会有一次数据库连接瞬断，请确保应用程序能重连数据库，谨慎操作。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BASENETWORKACCESSERROR = "FailedOperation.BaseNetworkAccessError"
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DELETEALLROUTE = "FailedOperation.DeleteAllRoute"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_ROGROUPNOTFOUNDERROR = "FailedOperation.ROGroupNotFoundError"
//  FAILEDOPERATION_VPCRESETSERVICEERROR = "FailedOperation.VPCResetServiceError"
//  FAILEDOPERATION_VPCUPDATEROUTEERROR = "FailedOperation.VPCUpdateRouteError"
//  INTERFACENAMENOTFOUND = "InterfaceNameNotFound"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  VPCERROR = "VpcError"
func (c *Client) RebalanceReadOnlyGroupWithContext(ctx context.Context, request *RebalanceReadOnlyGroupRequest) (response *RebalanceReadOnlyGroupResponse, err error) {
    if request == nil {
        request = NewRebalanceReadOnlyGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RebalanceReadOnlyGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewRebalanceReadOnlyGroupResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveDBInstanceFromReadOnlyGroupRequest() (request *RemoveDBInstanceFromReadOnlyGroupRequest) {
    request = &RemoveDBInstanceFromReadOnlyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "RemoveDBInstanceFromReadOnlyGroup")
    
    
    return
}

func NewRemoveDBInstanceFromReadOnlyGroupResponse() (response *RemoveDBInstanceFromReadOnlyGroupResponse) {
    response = &RemoveDBInstanceFromReadOnlyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemoveDBInstanceFromReadOnlyGroup
// 本接口（RemoveDBInstanceFromReadOnlyGroup）用户将只读实例从只读组中移除
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_ILLEGALROINSTANCENUM = "FailedOperation.IllegalROInstanceNum"
//  FAILEDOPERATION_ROGROUPMASTERINSTANCENOTRIGHT = "FailedOperation.ROGroupMasterInstanceNotRight"
//  FAILEDOPERATION_ROGROUPNOTFOUNDERROR = "FailedOperation.ROGroupNotFoundError"
//  INSTANCENOTEXIST = "InstanceNotExist"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDINSTANCESTATUS = "InvalidInstanceStatus"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDREADONLYGROUPSTATUS = "InvalidParameterValue.InvalidReadOnlyGroupStatus"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_READONLYGROUPNOTEXIST = "InvalidParameterValue.ReadOnlyGroupNotExist"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_ROGROUPSTATUSERROR = "OperationDenied.ROGroupStatusError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
func (c *Client) RemoveDBInstanceFromReadOnlyGroup(request *RemoveDBInstanceFromReadOnlyGroupRequest) (response *RemoveDBInstanceFromReadOnlyGroupResponse, err error) {
    return c.RemoveDBInstanceFromReadOnlyGroupWithContext(context.Background(), request)
}

// RemoveDBInstanceFromReadOnlyGroup
// 本接口（RemoveDBInstanceFromReadOnlyGroup）用户将只读实例从只读组中移除
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_ILLEGALROINSTANCENUM = "FailedOperation.IllegalROInstanceNum"
//  FAILEDOPERATION_ROGROUPMASTERINSTANCENOTRIGHT = "FailedOperation.ROGroupMasterInstanceNotRight"
//  FAILEDOPERATION_ROGROUPNOTFOUNDERROR = "FailedOperation.ROGroupNotFoundError"
//  INSTANCENOTEXIST = "InstanceNotExist"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDINSTANCESTATUS = "InvalidInstanceStatus"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDREADONLYGROUPSTATUS = "InvalidParameterValue.InvalidReadOnlyGroupStatus"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_READONLYGROUPNOTEXIST = "InvalidParameterValue.ReadOnlyGroupNotExist"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_ROGROUPSTATUSERROR = "OperationDenied.ROGroupStatusError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
func (c *Client) RemoveDBInstanceFromReadOnlyGroupWithContext(ctx context.Context, request *RemoveDBInstanceFromReadOnlyGroupRequest) (response *RemoveDBInstanceFromReadOnlyGroupResponse, err error) {
    if request == nil {
        request = NewRemoveDBInstanceFromReadOnlyGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveDBInstanceFromReadOnlyGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveDBInstanceFromReadOnlyGroupResponse()
    err = c.Send(request, response)
    return
}

func NewRenewInstanceRequest() (request *RenewInstanceRequest) {
    request = &RenewInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "RenewInstance")
    
    
    return
}

func NewRenewInstanceResponse() (response *RenewInstanceResponse) {
    response = &RenewInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RenewInstance
// 本接口（RenewInstance）用于续费实例。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_GETVPCINFOERROR = "FailedOperation.GetVpcInfoError"
//  FAILEDOPERATION_QUERYVPCFAILED = "FailedOperation.QueryVpcFailed"
//  FAILEDOPERATION_QUERYVPCFALIED = "FailedOperation.QueryVpcFalied"
//  FAILEDOPERATION_TRADECREATEERROR = "FailedOperation.TradeCreateError"
//  FAILEDOPERATION_TRADEPAYORDERSERROR = "FailedOperation.TradePayOrdersError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_VPCNOTFOUNDERROR = "InvalidParameter.VpcNotFoundError"
//  INVALIDPARAMETERVALUE_BADSPEC = "InvalidParameterValue.BadSpec"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCECHARGETYPE = "InvalidParameterValue.IllegalInstanceChargeType"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  OPERATIONDENIED_POSTPAIDPAYMODEERROR = "OperationDenied.PostPaidPayModeError"
//  OPERATIONDENIED_VPCDENIEDERROR = "OperationDenied.VpcDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCENOPERMISSION = "ResourceUnavailable.ResourceNoPermission"
//  RESOURCEUNAVAILABLE_VPCRESOURCENOTFOUND = "ResourceUnavailable.VpcResourceNotFound"
func (c *Client) RenewInstance(request *RenewInstanceRequest) (response *RenewInstanceResponse, err error) {
    return c.RenewInstanceWithContext(context.Background(), request)
}

// RenewInstance
// 本接口（RenewInstance）用于续费实例。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_GETVPCINFOERROR = "FailedOperation.GetVpcInfoError"
//  FAILEDOPERATION_QUERYVPCFAILED = "FailedOperation.QueryVpcFailed"
//  FAILEDOPERATION_QUERYVPCFALIED = "FailedOperation.QueryVpcFalied"
//  FAILEDOPERATION_TRADECREATEERROR = "FailedOperation.TradeCreateError"
//  FAILEDOPERATION_TRADEPAYORDERSERROR = "FailedOperation.TradePayOrdersError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_VPCNOTFOUNDERROR = "InvalidParameter.VpcNotFoundError"
//  INVALIDPARAMETERVALUE_BADSPEC = "InvalidParameterValue.BadSpec"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCECHARGETYPE = "InvalidParameterValue.IllegalInstanceChargeType"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  OPERATIONDENIED_POSTPAIDPAYMODEERROR = "OperationDenied.PostPaidPayModeError"
//  OPERATIONDENIED_VPCDENIEDERROR = "OperationDenied.VpcDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCENOPERMISSION = "ResourceUnavailable.ResourceNoPermission"
//  RESOURCEUNAVAILABLE_VPCRESOURCENOTFOUND = "ResourceUnavailable.VpcResourceNotFound"
func (c *Client) RenewInstanceWithContext(ctx context.Context, request *RenewInstanceRequest) (response *RenewInstanceResponse, err error) {
    if request == nil {
        request = NewRenewInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RenewInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewRenewInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewResetAccountPasswordRequest() (request *ResetAccountPasswordRequest) {
    request = &ResetAccountPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "ResetAccountPassword")
    
    
    return
}

func NewResetAccountPasswordResponse() (response *ResetAccountPasswordResponse) {
    response = &ResetAccountPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResetAccountPassword
// 本接口（ResetAccountPassword）用于重置实例的账户密码。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_OSSACCESSERROR = "FailedOperation.OssAccessError"
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  FAILEDOPERATION_PRECHECKERROR = "FailedOperation.PreCheckError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDACCOUNTPASSWORD = "InvalidAccountPassword"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_DECRYPTPASSWORDFAILED = "InvalidParameterValue.DecryptPasswordFailed"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAMEERROR = "InvalidParameterValue.InvalidAccountNameError"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAMEFORMATERROR = "InvalidParameterValue.InvalidAccountNameFormatError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORDFORMAT = "InvalidParameterValue.InvalidPasswordFormat"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORDLENGTHERROR = "InvalidParameterValue.InvalidPasswordLengthError"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORDVALUEERROR = "InvalidParameterValue.InvalidPasswordValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) ResetAccountPassword(request *ResetAccountPasswordRequest) (response *ResetAccountPasswordResponse, err error) {
    return c.ResetAccountPasswordWithContext(context.Background(), request)
}

// ResetAccountPassword
// 本接口（ResetAccountPassword）用于重置实例的账户密码。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_OSSACCESSERROR = "FailedOperation.OssAccessError"
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  FAILEDOPERATION_PRECHECKERROR = "FailedOperation.PreCheckError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDACCOUNTPASSWORD = "InvalidAccountPassword"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_DECRYPTPASSWORDFAILED = "InvalidParameterValue.DecryptPasswordFailed"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAMEERROR = "InvalidParameterValue.InvalidAccountNameError"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAMEFORMATERROR = "InvalidParameterValue.InvalidAccountNameFormatError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORDFORMAT = "InvalidParameterValue.InvalidPasswordFormat"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORDLENGTHERROR = "InvalidParameterValue.InvalidPasswordLengthError"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORDVALUEERROR = "InvalidParameterValue.InvalidPasswordValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) ResetAccountPasswordWithContext(ctx context.Context, request *ResetAccountPasswordRequest) (response *ResetAccountPasswordResponse, err error) {
    if request == nil {
        request = NewResetAccountPasswordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetAccountPassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetAccountPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewRestartDBInstanceRequest() (request *RestartDBInstanceRequest) {
    request = &RestartDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "RestartDBInstance")
    
    
    return
}

func NewRestartDBInstanceResponse() (response *RestartDBInstanceResponse) {
    response = &RestartDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RestartDBInstance
// 本接口（RestartDBInstance）用于重启实例。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_FLOWERROR = "InternalError.FlowError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
func (c *Client) RestartDBInstance(request *RestartDBInstanceRequest) (response *RestartDBInstanceResponse, err error) {
    return c.RestartDBInstanceWithContext(context.Background(), request)
}

// RestartDBInstance
// 本接口（RestartDBInstance）用于重启实例。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_FLOWERROR = "InternalError.FlowError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
func (c *Client) RestartDBInstanceWithContext(ctx context.Context, request *RestartDBInstanceRequest) (response *RestartDBInstanceResponse, err error) {
    if request == nil {
        request = NewRestartDBInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestartDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewRestartDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewRestoreDBInstanceObjectsRequest() (request *RestoreDBInstanceObjectsRequest) {
    request = &RestoreDBInstanceObjectsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "RestoreDBInstanceObjects")
    
    
    return
}

func NewRestoreDBInstanceObjectsResponse() (response *RestoreDBInstanceObjectsResponse) {
    response = &RestoreDBInstanceObjectsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RestoreDBInstanceObjects
// 根据备份集或恢复目标时间，在原实例上恢复数据库相关对象，例如数据库、表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) RestoreDBInstanceObjects(request *RestoreDBInstanceObjectsRequest) (response *RestoreDBInstanceObjectsResponse, err error) {
    return c.RestoreDBInstanceObjectsWithContext(context.Background(), request)
}

// RestoreDBInstanceObjects
// 根据备份集或恢复目标时间，在原实例上恢复数据库相关对象，例如数据库、表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) RestoreDBInstanceObjectsWithContext(ctx context.Context, request *RestoreDBInstanceObjectsRequest) (response *RestoreDBInstanceObjectsResponse, err error) {
    if request == nil {
        request = NewRestoreDBInstanceObjectsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestoreDBInstanceObjects require credential")
    }

    request.SetContext(ctx)
    
    response = NewRestoreDBInstanceObjectsResponse()
    err = c.Send(request, response)
    return
}

func NewSetAutoRenewFlagRequest() (request *SetAutoRenewFlagRequest) {
    request = &SetAutoRenewFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "SetAutoRenewFlag")
    
    
    return
}

func NewSetAutoRenewFlagResponse() (response *SetAutoRenewFlagResponse) {
    response = &SetAutoRenewFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetAutoRenewFlag
// 本接口（SetAutoRenewFlag）用于设置自动续费。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_SETAUTORENEWFLAGERROR = "FailedOperation.SetAutoRenewFlagError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCECHARGETYPE = "InvalidParameterValue.IllegalInstanceChargeType"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  OPERATIONDENIED_POSTPAIDPAYMODEERROR = "OperationDenied.PostPaidPayModeError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
func (c *Client) SetAutoRenewFlag(request *SetAutoRenewFlagRequest) (response *SetAutoRenewFlagResponse, err error) {
    return c.SetAutoRenewFlagWithContext(context.Background(), request)
}

// SetAutoRenewFlag
// 本接口（SetAutoRenewFlag）用于设置自动续费。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_SETAUTORENEWFLAGERROR = "FailedOperation.SetAutoRenewFlagError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCECHARGETYPE = "InvalidParameterValue.IllegalInstanceChargeType"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  OPERATIONDENIED_POSTPAIDPAYMODEERROR = "OperationDenied.PostPaidPayModeError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
func (c *Client) SetAutoRenewFlagWithContext(ctx context.Context, request *SetAutoRenewFlagRequest) (response *SetAutoRenewFlagResponse, err error) {
    if request == nil {
        request = NewSetAutoRenewFlagRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetAutoRenewFlag require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetAutoRenewFlagResponse()
    err = c.Send(request, response)
    return
}

func NewSwitchDBInstancePrimaryRequest() (request *SwitchDBInstancePrimaryRequest) {
    request = &SwitchDBInstancePrimaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "SwitchDBInstancePrimary")
    
    
    return
}

func NewSwitchDBInstancePrimaryResponse() (response *SwitchDBInstancePrimaryResponse) {
    response = &SwitchDBInstancePrimaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SwitchDBInstancePrimary
// 本接口（SwitchDBInstancePrimary）用于切换实例主备关系。
//
// <li>通过主动发起切换，可以验证业务能否正确处理实例主备切换的场景
//
// <li>通过使用强制切换，可以在备节点延迟不满足切换条件时，强制发起主从切换
//
// <li>只有主实例可以执行该操作
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_NOAVAILABLESTANDBY = "FailedOperation.NoAvailableStandby"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
//  OPERATIONDENIED_INSTANCETYPEDENIED = "OperationDenied.InstanceTypeDenied"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) SwitchDBInstancePrimary(request *SwitchDBInstancePrimaryRequest) (response *SwitchDBInstancePrimaryResponse, err error) {
    return c.SwitchDBInstancePrimaryWithContext(context.Background(), request)
}

// SwitchDBInstancePrimary
// 本接口（SwitchDBInstancePrimary）用于切换实例主备关系。
//
// <li>通过主动发起切换，可以验证业务能否正确处理实例主备切换的场景
//
// <li>通过使用强制切换，可以在备节点延迟不满足切换条件时，强制发起主从切换
//
// <li>只有主实例可以执行该操作
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_NOAVAILABLESTANDBY = "FailedOperation.NoAvailableStandby"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
//  OPERATIONDENIED_INSTANCETYPEDENIED = "OperationDenied.InstanceTypeDenied"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) SwitchDBInstancePrimaryWithContext(ctx context.Context, request *SwitchDBInstancePrimaryRequest) (response *SwitchDBInstancePrimaryResponse, err error) {
    if request == nil {
        request = NewSwitchDBInstancePrimaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SwitchDBInstancePrimary require credential")
    }

    request.SetContext(ctx)
    
    response = NewSwitchDBInstancePrimaryResponse()
    err = c.Send(request, response)
    return
}

func NewUnlockAccountRequest() (request *UnlockAccountRequest) {
    request = &UnlockAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "UnlockAccount")
    
    
    return
}

func NewUnlockAccountResponse() (response *UnlockAccountResponse) {
    response = &UnlockAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UnlockAccount
// 解除数据库账号的锁定，解锁后账号可以登陆数据库。
//
// 可能返回的错误码:
//  DBERROR = "DBError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_PRECHECKERROR = "FailedOperation.PreCheckError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAMEERROR = "InvalidParameterValue.InvalidAccountNameError"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAMEFORMATERROR = "InvalidParameterValue.InvalidAccountNameFormatError"
//  INVALIDPARAMETERVALUE_INVALIDDATABASEOBJECTNAME = "InvalidParameterValue.InvalidDatabaseObjectName"
//  INVALIDPARAMETERVALUE_PARAMETERLENGTHLIMITERROR = "InvalidParameterValue.ParameterLengthLimitError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) UnlockAccount(request *UnlockAccountRequest) (response *UnlockAccountResponse, err error) {
    return c.UnlockAccountWithContext(context.Background(), request)
}

// UnlockAccount
// 解除数据库账号的锁定，解锁后账号可以登陆数据库。
//
// 可能返回的错误码:
//  DBERROR = "DBError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_PRECHECKERROR = "FailedOperation.PreCheckError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAMEERROR = "InvalidParameterValue.InvalidAccountNameError"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAMEFORMATERROR = "InvalidParameterValue.InvalidAccountNameFormatError"
//  INVALIDPARAMETERVALUE_INVALIDDATABASEOBJECTNAME = "InvalidParameterValue.InvalidDatabaseObjectName"
//  INVALIDPARAMETERVALUE_PARAMETERLENGTHLIMITERROR = "InvalidParameterValue.ParameterLengthLimitError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) UnlockAccountWithContext(ctx context.Context, request *UnlockAccountRequest) (response *UnlockAccountResponse, err error) {
    if request == nil {
        request = NewUnlockAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnlockAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnlockAccountResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeDBInstanceRequest() (request *UpgradeDBInstanceRequest) {
    request = &UpgradeDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "UpgradeDBInstance")
    
    
    return
}

func NewUpgradeDBInstanceResponse() (response *UpgradeDBInstanceResponse) {
    response = &UpgradeDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpgradeDBInstance
// 本接口（UpgradeDBInstance）用于升级实例配置。本接口已废弃，推荐使用接口[ModifyDBInstanceSpec](https://cloud.tencent.com/document/api/409/63689)替代。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_GETPURCHASERECORDFAILED = "FailedOperation.GetPurchaseRecordFailed"
//  FAILEDOPERATION_PAYORDERFAILED = "FailedOperation.PayOrderFailed"
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  FAILEDOPERATION_QUERYSPECERROR = "FailedOperation.QuerySpecError"
//  FAILEDOPERATION_STORAGEMEMORYCHECKERROR = "FailedOperation.StorageMemoryCheckError"
//  FAILEDOPERATION_TRADEACCESSERROR = "FailedOperation.TradeAccessError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BADSPEC = "InvalidParameterValue.BadSpec"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEVOLUME = "InvalidParameterValue.InvalidInstanceVolume"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPID = "InvalidPid"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
func (c *Client) UpgradeDBInstance(request *UpgradeDBInstanceRequest) (response *UpgradeDBInstanceResponse, err error) {
    return c.UpgradeDBInstanceWithContext(context.Background(), request)
}

// UpgradeDBInstance
// 本接口（UpgradeDBInstance）用于升级实例配置。本接口已废弃，推荐使用接口[ModifyDBInstanceSpec](https://cloud.tencent.com/document/api/409/63689)替代。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_GETPURCHASERECORDFAILED = "FailedOperation.GetPurchaseRecordFailed"
//  FAILEDOPERATION_PAYORDERFAILED = "FailedOperation.PayOrderFailed"
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  FAILEDOPERATION_QUERYSPECERROR = "FailedOperation.QuerySpecError"
//  FAILEDOPERATION_STORAGEMEMORYCHECKERROR = "FailedOperation.StorageMemoryCheckError"
//  FAILEDOPERATION_TRADEACCESSERROR = "FailedOperation.TradeAccessError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BADSPEC = "InvalidParameterValue.BadSpec"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEVOLUME = "InvalidParameterValue.InvalidInstanceVolume"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPID = "InvalidPid"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
func (c *Client) UpgradeDBInstanceWithContext(ctx context.Context, request *UpgradeDBInstanceRequest) (response *UpgradeDBInstanceResponse, err error) {
    if request == nil {
        request = NewUpgradeDBInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeDBInstanceKernelVersionRequest() (request *UpgradeDBInstanceKernelVersionRequest) {
    request = &UpgradeDBInstanceKernelVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "UpgradeDBInstanceKernelVersion")
    
    
    return
}

func NewUpgradeDBInstanceKernelVersionResponse() (response *UpgradeDBInstanceKernelVersionResponse) {
    response = &UpgradeDBInstanceKernelVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpgradeDBInstanceKernelVersion
// 本接口（UpgradeDBInstanceKernelVersion）用于升级实例的内核版本号。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_PRECHECKERROR = "FailedOperation.PreCheckError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) UpgradeDBInstanceKernelVersion(request *UpgradeDBInstanceKernelVersionRequest) (response *UpgradeDBInstanceKernelVersionResponse, err error) {
    return c.UpgradeDBInstanceKernelVersionWithContext(context.Background(), request)
}

// UpgradeDBInstanceKernelVersion
// 本接口（UpgradeDBInstanceKernelVersion）用于升级实例的内核版本号。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_PRECHECKERROR = "FailedOperation.PreCheckError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) UpgradeDBInstanceKernelVersionWithContext(ctx context.Context, request *UpgradeDBInstanceKernelVersionRequest) (response *UpgradeDBInstanceKernelVersionResponse, err error) {
    if request == nil {
        request = NewUpgradeDBInstanceKernelVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeDBInstanceKernelVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeDBInstanceKernelVersionResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeDBInstanceMajorVersionRequest() (request *UpgradeDBInstanceMajorVersionRequest) {
    request = &UpgradeDBInstanceMajorVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("postgres", APIVersion, "UpgradeDBInstanceMajorVersion")
    
    
    return
}

func NewUpgradeDBInstanceMajorVersionResponse() (response *UpgradeDBInstanceMajorVersionResponse) {
    response = &UpgradeDBInstanceMajorVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpgradeDBInstanceMajorVersion
// 本接口（UpgradeDBInstanceMajorVersion）用于升级实例内核大版本，例如从PostgreSQL 12升级到PostgreSQL 15。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) UpgradeDBInstanceMajorVersion(request *UpgradeDBInstanceMajorVersionRequest) (response *UpgradeDBInstanceMajorVersionResponse, err error) {
    return c.UpgradeDBInstanceMajorVersionWithContext(context.Background(), request)
}

// UpgradeDBInstanceMajorVersion
// 本接口（UpgradeDBInstanceMajorVersion）用于升级实例内核大版本，例如从PostgreSQL 12升级到PostgreSQL 15。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) UpgradeDBInstanceMajorVersionWithContext(ctx context.Context, request *UpgradeDBInstanceMajorVersionRequest) (response *UpgradeDBInstanceMajorVersionResponse, err error) {
    if request == nil {
        request = NewUpgradeDBInstanceMajorVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeDBInstanceMajorVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeDBInstanceMajorVersionResponse()
    err = c.Send(request, response)
    return
}
