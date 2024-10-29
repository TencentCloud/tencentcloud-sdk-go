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


func NewCheckMigrateIndexMetaDataRequest() (request *CheckMigrateIndexMetaDataRequest) {
    request = &CheckMigrateIndexMetaDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "CheckMigrateIndexMetaData")
    
    
    return
}

func NewCheckMigrateIndexMetaDataResponse() (response *CheckMigrateIndexMetaDataResponse) {
    response = &CheckMigrateIndexMetaDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckMigrateIndexMetaData
// 检查cos迁移索引元数据
func (c *Client) CheckMigrateIndexMetaData(request *CheckMigrateIndexMetaDataRequest) (response *CheckMigrateIndexMetaDataResponse, err error) {
    return c.CheckMigrateIndexMetaDataWithContext(context.Background(), request)
}

// CheckMigrateIndexMetaData
// 检查cos迁移索引元数据
func (c *Client) CheckMigrateIndexMetaDataWithContext(ctx context.Context, request *CheckMigrateIndexMetaDataRequest) (response *CheckMigrateIndexMetaDataResponse, err error) {
    if request == nil {
        request = NewCheckMigrateIndexMetaDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckMigrateIndexMetaData require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckMigrateIndexMetaDataResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClusterSnapshotRequest() (request *CreateClusterSnapshotRequest) {
    request = &CreateClusterSnapshotRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "CreateClusterSnapshot")
    
    
    return
}

func NewCreateClusterSnapshotResponse() (response *CreateClusterSnapshotResponse) {
    response = &CreateClusterSnapshotResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateClusterSnapshot
// 集群快照手动创建
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDESVERSION = "InvalidParameter.InvalidEsVersion"
//  INVALIDPARAMETER_INVALIDUIN = "InvalidParameter.InvalidUin"
//  RESOURCENOTFOUND_CLUSTERINFONOTFOUND = "ResourceNotFound.ClusterInfoNotFound"
func (c *Client) CreateClusterSnapshot(request *CreateClusterSnapshotRequest) (response *CreateClusterSnapshotResponse, err error) {
    return c.CreateClusterSnapshotWithContext(context.Background(), request)
}

// CreateClusterSnapshot
// 集群快照手动创建
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDESVERSION = "InvalidParameter.InvalidEsVersion"
//  INVALIDPARAMETER_INVALIDUIN = "InvalidParameter.InvalidUin"
//  RESOURCENOTFOUND_CLUSTERINFONOTFOUND = "ResourceNotFound.ClusterInfoNotFound"
func (c *Client) CreateClusterSnapshotWithContext(ctx context.Context, request *CreateClusterSnapshotRequest) (response *CreateClusterSnapshotResponse, err error) {
    if request == nil {
        request = NewCreateClusterSnapshotRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateClusterSnapshot require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateClusterSnapshotResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCosMigrateToServerlessInstanceRequest() (request *CreateCosMigrateToServerlessInstanceRequest) {
    request = &CreateCosMigrateToServerlessInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "CreateCosMigrateToServerlessInstance")
    
    
    return
}

func NewCreateCosMigrateToServerlessInstanceResponse() (response *CreateCosMigrateToServerlessInstanceResponse) {
    response = &CreateCosMigrateToServerlessInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCosMigrateToServerlessInstance
// cos迁移流程
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDESVERSION = "InvalidParameter.InvalidEsVersion"
//  INVALIDPARAMETER_INVALIDUIN = "InvalidParameter.InvalidUin"
//  RESOURCENOTFOUND_CLUSTERINFONOTFOUND = "ResourceNotFound.ClusterInfoNotFound"
func (c *Client) CreateCosMigrateToServerlessInstance(request *CreateCosMigrateToServerlessInstanceRequest) (response *CreateCosMigrateToServerlessInstanceResponse, err error) {
    return c.CreateCosMigrateToServerlessInstanceWithContext(context.Background(), request)
}

// CreateCosMigrateToServerlessInstance
// cos迁移流程
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDESVERSION = "InvalidParameter.InvalidEsVersion"
//  INVALIDPARAMETER_INVALIDUIN = "InvalidParameter.InvalidUin"
//  RESOURCENOTFOUND_CLUSTERINFONOTFOUND = "ResourceNotFound.ClusterInfoNotFound"
func (c *Client) CreateCosMigrateToServerlessInstanceWithContext(ctx context.Context, request *CreateCosMigrateToServerlessInstanceRequest) (response *CreateCosMigrateToServerlessInstanceResponse, err error) {
    if request == nil {
        request = NewCreateCosMigrateToServerlessInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCosMigrateToServerlessInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCosMigrateToServerlessInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateIndexRequest() (request *CreateIndexRequest) {
    request = &CreateIndexRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "CreateIndex")
    
    
    return
}

func NewCreateIndexResponse() (response *CreateIndexResponse) {
    response = &CreateIndexResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateIndex
// 创建索引
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDESVERSION = "InvalidParameter.InvalidEsVersion"
//  INVALIDPARAMETER_INVALIDINDEXMETAJSON = "InvalidParameter.InvalidIndexMetaJson"
//  INVALIDPARAMETER_INVALIDINDEXNAME = "InvalidParameter.InvalidIndexName"
//  INVALIDPARAMETER_INVALIDINDEXTYPE = "InvalidParameter.InvalidIndexType"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  RESOURCENOTFOUND_CLUSTERINFONOTFOUND = "ResourceNotFound.ClusterInfoNotFound"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
func (c *Client) CreateIndex(request *CreateIndexRequest) (response *CreateIndexResponse, err error) {
    return c.CreateIndexWithContext(context.Background(), request)
}

// CreateIndex
// 创建索引
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDESVERSION = "InvalidParameter.InvalidEsVersion"
//  INVALIDPARAMETER_INVALIDINDEXMETAJSON = "InvalidParameter.InvalidIndexMetaJson"
//  INVALIDPARAMETER_INVALIDINDEXNAME = "InvalidParameter.InvalidIndexName"
//  INVALIDPARAMETER_INVALIDINDEXTYPE = "InvalidParameter.InvalidIndexType"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  RESOURCENOTFOUND_CLUSTERINFONOTFOUND = "ResourceNotFound.ClusterInfoNotFound"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
func (c *Client) CreateIndexWithContext(ctx context.Context, request *CreateIndexRequest) (response *CreateIndexResponse, err error) {
    if request == nil {
        request = NewCreateIndexRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateIndex require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateIndexResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstanceRequest() (request *CreateInstanceRequest) {
    request = &CreateInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "CreateInstance")
    
    
    return
}

func NewCreateInstanceResponse() (response *CreateInstanceResponse) {
    response = &CreateInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateInstance
// 创建指定规格的ES集群实例
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHCREATEINSTANCE = "AuthFailure.UnAuthCreateInstance"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION_CLUSTERRESOURCELIMITERROR = "FailedOperation.ClusterResourceLimitError"
//  FAILEDOPERATION_DISKCOUNTPARAMERROR = "FailedOperation.DiskCountParamError"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  FAILEDOPERATION_NODENUMANDZONEERROR = "FailedOperation.NodeNumAndZoneError"
//  FAILEDOPERATION_NOTAUTHENTICATED = "FailedOperation.NotAuthenticated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDAUTOVOUCHER = "InvalidParameter.InvalidAutoVoucher"
//  INVALIDPARAMETER_INVALIDDEPLOYMODE = "InvalidParameter.InvalidDeployMode"
//  INVALIDPARAMETER_INVALIDDISKCOUNT = "InvalidParameter.InvalidDiskCount"
//  INVALIDPARAMETER_INVALIDDISKENCRYPT = "InvalidParameter.InvalidDiskEncrypt"
//  INVALIDPARAMETER_INVALIDDISKENHANCE = "InvalidParameter.InvalidDiskEnhance"
//  INVALIDPARAMETER_INVALIDESVERSION = "InvalidParameter.InvalidEsVersion"
//  INVALIDPARAMETER_INVALIDMULTIZONEINFO = "InvalidParameter.InvalidMultiZoneInfo"
//  INVALIDPARAMETER_INVALIDNODENUM = "InvalidParameter.InvalidNodeNum"
//  INVALIDPARAMETER_INVALIDNODETYPE = "InvalidParameter.InvalidNodeType"
//  INVALIDPARAMETER_INVALIDOPTYPE = "InvalidParameter.InvalidOpType"
//  INVALIDPARAMETER_INVALIDOPERATIONDURATION = "InvalidParameter.InvalidOperationDuration"
//  INVALIDPARAMETER_INVALIDREGION = "InvalidParameter.InvalidRegion"
//  INVALIDPARAMETER_INVALIDTAGLIST = "InvalidParameter.InvalidTagList"
//  INVALIDPARAMETER_INVALIDTYPE = "InvalidParameter.InvalidType"
//  INVALIDPARAMETER_INVALIDVOUCHERIDS = "InvalidParameter.InvalidVoucherIds"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  INVALIDPARAMETERVALUE_CHARGETYPE = "InvalidParameterValue.ChargeType"
//  INVALIDPARAMETERVALUE_RENEWFLAG = "InvalidParameterValue.RenewFlag"
//  LIMITEXCEEDED_CLUSTERNUM = "LimitExceeded.ClusterNum"
//  LIMITEXCEEDED_RESOURCELIMIT = "LimitExceeded.ResourceLimit"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_ORDER = "ResourceInUse.Order"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_HIDDENZONE = "ResourceInsufficient.HiddenZone"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
//  RESOURCEINSUFFICIENT_ZONE = "ResourceInsufficient.Zone"
//  RESOURCENOTFOUND_DISKINFONOTFOUND = "ResourceNotFound.DiskInfoNotFound"
//  RESOURCENOTFOUND_TRADECGWNOTFOUND = "ResourceNotFound.TradeCgwNotFound"
//  RESOURCENOTFOUND_VPCINFONOTFOUND = "ResourceNotFound.VPCInfoNotFound"
//  RESOURCENOTFOUND_WHITELISTNOTFOUND = "ResourceNotFound.WhiteListNotFound"
//  UNAUTHORIZEDOPERATION_UINNOTINWHITELIST = "UnauthorizedOperation.UinNotInWhiteList"
//  UNSUPPORTEDOPERATION_BASICSECURITYTYPE = "UnsupportedOperation.BasicSecurityType"
//  UNSUPPORTEDOPERATION_LICENSEERROR = "UnsupportedOperation.LicenseError"
//  UNSUPPORTEDOPERATION_VPCINFONOTFOUND = "UnsupportedOperation.VPCInfoNotFound"
func (c *Client) CreateInstance(request *CreateInstanceRequest) (response *CreateInstanceResponse, err error) {
    return c.CreateInstanceWithContext(context.Background(), request)
}

// CreateInstance
// 创建指定规格的ES集群实例
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHCREATEINSTANCE = "AuthFailure.UnAuthCreateInstance"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION_CLUSTERRESOURCELIMITERROR = "FailedOperation.ClusterResourceLimitError"
//  FAILEDOPERATION_DISKCOUNTPARAMERROR = "FailedOperation.DiskCountParamError"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  FAILEDOPERATION_NODENUMANDZONEERROR = "FailedOperation.NodeNumAndZoneError"
//  FAILEDOPERATION_NOTAUTHENTICATED = "FailedOperation.NotAuthenticated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDAUTOVOUCHER = "InvalidParameter.InvalidAutoVoucher"
//  INVALIDPARAMETER_INVALIDDEPLOYMODE = "InvalidParameter.InvalidDeployMode"
//  INVALIDPARAMETER_INVALIDDISKCOUNT = "InvalidParameter.InvalidDiskCount"
//  INVALIDPARAMETER_INVALIDDISKENCRYPT = "InvalidParameter.InvalidDiskEncrypt"
//  INVALIDPARAMETER_INVALIDDISKENHANCE = "InvalidParameter.InvalidDiskEnhance"
//  INVALIDPARAMETER_INVALIDESVERSION = "InvalidParameter.InvalidEsVersion"
//  INVALIDPARAMETER_INVALIDMULTIZONEINFO = "InvalidParameter.InvalidMultiZoneInfo"
//  INVALIDPARAMETER_INVALIDNODENUM = "InvalidParameter.InvalidNodeNum"
//  INVALIDPARAMETER_INVALIDNODETYPE = "InvalidParameter.InvalidNodeType"
//  INVALIDPARAMETER_INVALIDOPTYPE = "InvalidParameter.InvalidOpType"
//  INVALIDPARAMETER_INVALIDOPERATIONDURATION = "InvalidParameter.InvalidOperationDuration"
//  INVALIDPARAMETER_INVALIDREGION = "InvalidParameter.InvalidRegion"
//  INVALIDPARAMETER_INVALIDTAGLIST = "InvalidParameter.InvalidTagList"
//  INVALIDPARAMETER_INVALIDTYPE = "InvalidParameter.InvalidType"
//  INVALIDPARAMETER_INVALIDVOUCHERIDS = "InvalidParameter.InvalidVoucherIds"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  INVALIDPARAMETERVALUE_CHARGETYPE = "InvalidParameterValue.ChargeType"
//  INVALIDPARAMETERVALUE_RENEWFLAG = "InvalidParameterValue.RenewFlag"
//  LIMITEXCEEDED_CLUSTERNUM = "LimitExceeded.ClusterNum"
//  LIMITEXCEEDED_RESOURCELIMIT = "LimitExceeded.ResourceLimit"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_ORDER = "ResourceInUse.Order"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_HIDDENZONE = "ResourceInsufficient.HiddenZone"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
//  RESOURCEINSUFFICIENT_ZONE = "ResourceInsufficient.Zone"
//  RESOURCENOTFOUND_DISKINFONOTFOUND = "ResourceNotFound.DiskInfoNotFound"
//  RESOURCENOTFOUND_TRADECGWNOTFOUND = "ResourceNotFound.TradeCgwNotFound"
//  RESOURCENOTFOUND_VPCINFONOTFOUND = "ResourceNotFound.VPCInfoNotFound"
//  RESOURCENOTFOUND_WHITELISTNOTFOUND = "ResourceNotFound.WhiteListNotFound"
//  UNAUTHORIZEDOPERATION_UINNOTINWHITELIST = "UnauthorizedOperation.UinNotInWhiteList"
//  UNSUPPORTEDOPERATION_BASICSECURITYTYPE = "UnsupportedOperation.BasicSecurityType"
//  UNSUPPORTEDOPERATION_LICENSEERROR = "UnsupportedOperation.LicenseError"
//  UNSUPPORTEDOPERATION_VPCINFONOTFOUND = "UnsupportedOperation.VPCInfoNotFound"
func (c *Client) CreateInstanceWithContext(ctx context.Context, request *CreateInstanceRequest) (response *CreateInstanceResponse, err error) {
    if request == nil {
        request = NewCreateInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLogstashInstanceRequest() (request *CreateLogstashInstanceRequest) {
    request = &CreateLogstashInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "CreateLogstashInstance")
    
    
    return
}

func NewCreateLogstashInstanceResponse() (response *CreateLogstashInstanceResponse) {
    response = &CreateLogstashInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLogstashInstance
// 用于创建Logstash实例
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHCREATEINSTANCE = "AuthFailure.UnAuthCreateInstance"
//  AUTHFAILURE_UNAUTHDESCRIBEINSTANCES = "AuthFailure.UnAuthDescribeInstances"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION_CLOSEDATATIER = "FailedOperation.CloseDataTier"
//  FAILEDOPERATION_CLUSTERRESOURCELIMITERROR = "FailedOperation.ClusterResourceLimitError"
//  FAILEDOPERATION_ESDICTIONARYINFOERROR = "FailedOperation.EsDictionaryInfoError"
//  FAILEDOPERATION_FILENAMEERROR = "FailedOperation.FileNameError"
//  FAILEDOPERATION_FILESIZEERROR = "FailedOperation.FileSizeError"
//  FAILEDOPERATION_GETTAGINFOERROR = "FailedOperation.GetTagInfoError"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  FAILEDOPERATION_NODENUMANDZONEERROR = "FailedOperation.NodeNumAndZoneError"
//  FAILEDOPERATION_PASSWORD = "FailedOperation.Password"
//  FAILEDOPERATION_REFUNDERROR = "FailedOperation.RefundError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeOut"
//  FAILEDOPERATION_TRADESIGNERROR = "FailedOperation.TradeSignError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_GC = "InvalidParameter.GC"
//  INVALIDPARAMETER_INVALIDAUTOVOUCHER = "InvalidParameter.InvalidAutoVoucher"
//  INVALIDPARAMETER_INVALIDBACKENDS = "InvalidParameter.InvalidBackends"
//  INVALIDPARAMETER_INVALIDDEPLOYMODE = "InvalidParameter.InvalidDeployMode"
//  INVALIDPARAMETER_INVALIDDISKENHANCE = "InvalidParameter.InvalidDiskEnhance"
//  INVALIDPARAMETER_INVALIDDNS = "InvalidParameter.InvalidDns"
//  INVALIDPARAMETER_INVALIDESVERSION = "InvalidParameter.InvalidEsVersion"
//  INVALIDPARAMETER_INVALIDINDEXMETAJSON = "InvalidParameter.InvalidIndexMetaJson"
//  INVALIDPARAMETER_INVALIDINDEXNAME = "InvalidParameter.InvalidIndexName"
//  INVALIDPARAMETER_INVALIDINDEXOPTIONSFIELD = "InvalidParameter.InvalidIndexOptionsField"
//  INVALIDPARAMETER_INVALIDINDEXPOLICYFIELD = "InvalidParameter.InvalidIndexPolicyField"
//  INVALIDPARAMETER_INVALIDINDEXSETTINGSFIELD = "InvalidParameter.InvalidIndexSettingsField"
//  INVALIDPARAMETER_INVALIDINDEXTYPE = "InvalidParameter.InvalidIndexType"
//  INVALIDPARAMETER_INVALIDINSTANCEOPERATIONDURATIONS = "InvalidParameter.InvalidInstanceOperationDurations"
//  INVALIDPARAMETER_INVALIDLOGTYPE = "InvalidParameter.InvalidLogType"
//  INVALIDPARAMETER_INVALIDLOGSTASHVERSION = "InvalidParameter.InvalidLogstashVersion"
//  INVALIDPARAMETER_INVALIDNODENAMES = "InvalidParameter.InvalidNodeNames"
//  INVALIDPARAMETER_INVALIDNODETYPESTORAGEINFO = "InvalidParameter.InvalidNodeTypeStorageInfo"
//  INVALIDPARAMETER_INVALIDOPERATIONDURATION = "InvalidParameter.InvalidOperationDuration"
//  INVALIDPARAMETER_INVALIDPARA = "InvalidParameter.InvalidPara"
//  INVALIDPARAMETER_INVALIDRESTARTMODE = "InvalidParameter.InvalidRestartMode"
//  INVALIDPARAMETER_INVALIDRESTARTTYPE = "InvalidParameter.InvalidRestartType"
//  INVALIDPARAMETER_INVALIDSAMPLEJSON = "InvalidParameter.InvalidSampleJson"
//  INVALIDPARAMETER_INVALIDSECURITYGROUPIDS = "InvalidParameter.InvalidSecurityGroupIds"
//  INVALIDPARAMETER_INVALIDSUBNETUIDLIST = "InvalidParameter.InvalidSubnetUidList"
//  INVALIDPARAMETER_INVALIDTARGETINDEXNAME = "InvalidParameter.InvalidTargetIndexName"
//  INVALIDPARAMETER_INVALIDTARGETNODETYPES = "InvalidParameter.InvalidTargetNodeTypes"
//  INVALIDPARAMETER_INVALIDTIMEPARAM = "InvalidParameter.InvalidTimeParam"
//  INVALIDPARAMETER_INVALIDTRACEUUID = "InvalidParameter.InvalidTraceUuid"
//  INVALIDPARAMETER_INVALIDTYPE = "InvalidParameter.InvalidType"
//  INVALIDPARAMETER_INVALIDUPDATEMETAJSON = "InvalidParameter.InvalidUpdateMetaJson"
//  INVALIDPARAMETER_INVALIDVOUCHERIDS = "InvalidParameter.InvalidVoucherIds"
//  INVALIDPARAMETERVALUE_CHARGEPERIOD = "InvalidParameterValue.ChargePeriod"
//  INVALIDPARAMETERVALUE_CHARGETYPE = "InvalidParameterValue.ChargeType"
//  INVALIDPARAMETERVALUE_CHECKONLY = "InvalidParameterValue.CheckOnly"
//  INVALIDPARAMETERVALUE_CONFIGINFO = "InvalidParameterValue.ConfigInfo"
//  INVALIDPARAMETERVALUE_ENABLEDOUBLEENI = "InvalidParameterValue.EnableDoubleEni"
//  INVALIDPARAMETERVALUE_ESPORT = "InvalidParameterValue.EsPort"
//  INVALIDPARAMETERVALUE_ESVIP = "InvalidParameterValue.EsVip"
//  INVALIDPARAMETERVALUE_INSTALLBUNDLELIST = "InvalidParameterValue.InstallBundleList"
//  INVALIDPARAMETERVALUE_INSTALLPLUGINLIST = "InvalidParameterValue.InstallPluginList"
//  INVALIDPARAMETERVALUE_INVALIDJDK = "InvalidParameterValue.InvalidJDK"
//  INVALIDPARAMETERVALUE_PLUGINTYPE = "InvalidParameterValue.PluginType"
//  INVALIDPARAMETERVALUE_REMOVEBUNDLELIST = "InvalidParameterValue.RemoveBundleList"
//  INVALIDPARAMETERVALUE_REMOVEPLUGINLIST = "InvalidParameterValue.RemovePluginList"
//  INVALIDPARAMETERVALUE_RENEWFLAG = "InvalidParameterValue.RenewFlag"
//  INVALIDPARAMETERVALUE_TARGETPATH = "InvalidParameterValue.TargetPath"
//  INVALIDPARAMETERVALUE_TIMEUNIT = "InvalidParameterValue.TimeUnit"
//  INVALIDPARAMETERVALUE_UPGRADEMODE = "InvalidParameterValue.UpgradeMode"
//  LIMITEXCEEDED_CLUSTERNUM = "LimitExceeded.ClusterNum"
//  LIMITEXCEEDED_DIAGNOSECOUNT = "LimitExceeded.DiagnoseCount"
//  LIMITEXCEEDED_DISKCOUNT = "LimitExceeded.DiskCount"
//  LIMITEXCEEDED_NODENUMORINDICES = "LimitExceeded.NodeNumOrIndices"
//  LIMITEXCEEDED_PLUGININSTALL = "LimitExceeded.PluginInstall"
//  REGIONERROR = "RegionError"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_DIAGNOSE = "ResourceInUse.Diagnose"
//  RESOURCEINUSE_ORDER = "ResourceInUse.Order"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_CVM = "ResourceInsufficient.CVM"
//  RESOURCEINSUFFICIENT_CVMQUOTA = "ResourceInsufficient.CVMQuota"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
//  RESOURCENOTFOUND_ACCOUNTINFONOTFOUND = "ResourceNotFound.AccountInfoNotFound"
//  RESOURCENOTFOUND_BARADMETRICNOTFOUND = "ResourceNotFound.BaradMetricNotFound"
//  RESOURCENOTFOUND_CAMINFONOTFOUND = "ResourceNotFound.CAMInfoNotFound"
//  RESOURCENOTFOUND_CLBNOTFOUND = "ResourceNotFound.CLBNotFound"
//  RESOURCENOTFOUND_COSINFONOTFOUND = "ResourceNotFound.CosInfoNotFound"
//  RESOURCENOTFOUND_DIAGNOSENOTFOUND = "ResourceNotFound.DiagnoseNotFound"
//  RESOURCENOTFOUND_KMSNOTFOUND = "ResourceNotFound.KMSNotFound"
//  RESOURCENOTFOUND_STSNOTFOUND = "ResourceNotFound.STSNotFound"
//  RESOURCENOTFOUND_SECURITYGROUPNOTFOUND = "ResourceNotFound.SecurityGroupNotFound"
//  RESOURCENOTFOUND_TRADECGWNOTFOUND = "ResourceNotFound.TradeCgwNotFound"
//  RESOURCENOTFOUND_VPCINFONOTFOUND = "ResourceNotFound.VPCInfoNotFound"
//  UNSUPPORTEDOPERATION_AUDITSTATUSCONFLICT = "UnsupportedOperation.AuditStatusConflict"
//  UNSUPPORTEDOPERATION_COSBACKUP = "UnsupportedOperation.CosBackUp"
//  UNSUPPORTEDOPERATION_DIAGNOSEJOB = "UnsupportedOperation.DiagnoseJob"
//  UNSUPPORTEDOPERATION_DIAGNOSENOTOPEN = "UnsupportedOperation.DiagnoseNotOpen"
//  UNSUPPORTEDOPERATION_DISKUSE = "UnsupportedOperation.DiskUse"
//  UNSUPPORTEDOPERATION_EDITLISTLENGTH = "UnsupportedOperation.EditListLength"
//  UNSUPPORTEDOPERATION_MIGRATE = "UnsupportedOperation.Migrate"
//  UNSUPPORTEDOPERATION_PASSLOGSTASHID = "UnsupportedOperation.PassLogstashId"
//  UNSUPPORTEDOPERATION_PLUGIN = "UnsupportedOperation.Plugin"
//  UNSUPPORTEDOPERATION_PROTOCOL = "UnsupportedOperation.Protocol"
//  UNSUPPORTEDOPERATION_READRATE = "UnsupportedOperation.ReadRate"
//  UNSUPPORTEDOPERATION_RESTARTMODE = "UnsupportedOperation.RestartMode"
//  UNSUPPORTEDOPERATION_STATUSNOTNORMAL = "UnsupportedOperation.StatusNotNormal"
//  UNSUPPORTEDOPERATION_UPDATEDISKENCRYPT = "UnsupportedOperation.UpdateDiskEncrypt"
//  UNSUPPORTEDOPERATION_UPDATEDISKTYPE = "UnsupportedOperation.UpdateDiskType"
//  UNSUPPORTEDOPERATION_UPDATENODENUMANDNODETYPE = "UnsupportedOperation.UpdateNodeNumAndNodeType"
//  UNSUPPORTEDOPERATION_WEBSERVICETYPE = "UnsupportedOperation.WebServiceType"
//  UNSUPPORTEDOPERATION_WRITERATE = "UnsupportedOperation.WriteRate"
func (c *Client) CreateLogstashInstance(request *CreateLogstashInstanceRequest) (response *CreateLogstashInstanceResponse, err error) {
    return c.CreateLogstashInstanceWithContext(context.Background(), request)
}

// CreateLogstashInstance
// 用于创建Logstash实例
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHCREATEINSTANCE = "AuthFailure.UnAuthCreateInstance"
//  AUTHFAILURE_UNAUTHDESCRIBEINSTANCES = "AuthFailure.UnAuthDescribeInstances"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION_CLOSEDATATIER = "FailedOperation.CloseDataTier"
//  FAILEDOPERATION_CLUSTERRESOURCELIMITERROR = "FailedOperation.ClusterResourceLimitError"
//  FAILEDOPERATION_ESDICTIONARYINFOERROR = "FailedOperation.EsDictionaryInfoError"
//  FAILEDOPERATION_FILENAMEERROR = "FailedOperation.FileNameError"
//  FAILEDOPERATION_FILESIZEERROR = "FailedOperation.FileSizeError"
//  FAILEDOPERATION_GETTAGINFOERROR = "FailedOperation.GetTagInfoError"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  FAILEDOPERATION_NODENUMANDZONEERROR = "FailedOperation.NodeNumAndZoneError"
//  FAILEDOPERATION_PASSWORD = "FailedOperation.Password"
//  FAILEDOPERATION_REFUNDERROR = "FailedOperation.RefundError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeOut"
//  FAILEDOPERATION_TRADESIGNERROR = "FailedOperation.TradeSignError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_GC = "InvalidParameter.GC"
//  INVALIDPARAMETER_INVALIDAUTOVOUCHER = "InvalidParameter.InvalidAutoVoucher"
//  INVALIDPARAMETER_INVALIDBACKENDS = "InvalidParameter.InvalidBackends"
//  INVALIDPARAMETER_INVALIDDEPLOYMODE = "InvalidParameter.InvalidDeployMode"
//  INVALIDPARAMETER_INVALIDDISKENHANCE = "InvalidParameter.InvalidDiskEnhance"
//  INVALIDPARAMETER_INVALIDDNS = "InvalidParameter.InvalidDns"
//  INVALIDPARAMETER_INVALIDESVERSION = "InvalidParameter.InvalidEsVersion"
//  INVALIDPARAMETER_INVALIDINDEXMETAJSON = "InvalidParameter.InvalidIndexMetaJson"
//  INVALIDPARAMETER_INVALIDINDEXNAME = "InvalidParameter.InvalidIndexName"
//  INVALIDPARAMETER_INVALIDINDEXOPTIONSFIELD = "InvalidParameter.InvalidIndexOptionsField"
//  INVALIDPARAMETER_INVALIDINDEXPOLICYFIELD = "InvalidParameter.InvalidIndexPolicyField"
//  INVALIDPARAMETER_INVALIDINDEXSETTINGSFIELD = "InvalidParameter.InvalidIndexSettingsField"
//  INVALIDPARAMETER_INVALIDINDEXTYPE = "InvalidParameter.InvalidIndexType"
//  INVALIDPARAMETER_INVALIDINSTANCEOPERATIONDURATIONS = "InvalidParameter.InvalidInstanceOperationDurations"
//  INVALIDPARAMETER_INVALIDLOGTYPE = "InvalidParameter.InvalidLogType"
//  INVALIDPARAMETER_INVALIDLOGSTASHVERSION = "InvalidParameter.InvalidLogstashVersion"
//  INVALIDPARAMETER_INVALIDNODENAMES = "InvalidParameter.InvalidNodeNames"
//  INVALIDPARAMETER_INVALIDNODETYPESTORAGEINFO = "InvalidParameter.InvalidNodeTypeStorageInfo"
//  INVALIDPARAMETER_INVALIDOPERATIONDURATION = "InvalidParameter.InvalidOperationDuration"
//  INVALIDPARAMETER_INVALIDPARA = "InvalidParameter.InvalidPara"
//  INVALIDPARAMETER_INVALIDRESTARTMODE = "InvalidParameter.InvalidRestartMode"
//  INVALIDPARAMETER_INVALIDRESTARTTYPE = "InvalidParameter.InvalidRestartType"
//  INVALIDPARAMETER_INVALIDSAMPLEJSON = "InvalidParameter.InvalidSampleJson"
//  INVALIDPARAMETER_INVALIDSECURITYGROUPIDS = "InvalidParameter.InvalidSecurityGroupIds"
//  INVALIDPARAMETER_INVALIDSUBNETUIDLIST = "InvalidParameter.InvalidSubnetUidList"
//  INVALIDPARAMETER_INVALIDTARGETINDEXNAME = "InvalidParameter.InvalidTargetIndexName"
//  INVALIDPARAMETER_INVALIDTARGETNODETYPES = "InvalidParameter.InvalidTargetNodeTypes"
//  INVALIDPARAMETER_INVALIDTIMEPARAM = "InvalidParameter.InvalidTimeParam"
//  INVALIDPARAMETER_INVALIDTRACEUUID = "InvalidParameter.InvalidTraceUuid"
//  INVALIDPARAMETER_INVALIDTYPE = "InvalidParameter.InvalidType"
//  INVALIDPARAMETER_INVALIDUPDATEMETAJSON = "InvalidParameter.InvalidUpdateMetaJson"
//  INVALIDPARAMETER_INVALIDVOUCHERIDS = "InvalidParameter.InvalidVoucherIds"
//  INVALIDPARAMETERVALUE_CHARGEPERIOD = "InvalidParameterValue.ChargePeriod"
//  INVALIDPARAMETERVALUE_CHARGETYPE = "InvalidParameterValue.ChargeType"
//  INVALIDPARAMETERVALUE_CHECKONLY = "InvalidParameterValue.CheckOnly"
//  INVALIDPARAMETERVALUE_CONFIGINFO = "InvalidParameterValue.ConfigInfo"
//  INVALIDPARAMETERVALUE_ENABLEDOUBLEENI = "InvalidParameterValue.EnableDoubleEni"
//  INVALIDPARAMETERVALUE_ESPORT = "InvalidParameterValue.EsPort"
//  INVALIDPARAMETERVALUE_ESVIP = "InvalidParameterValue.EsVip"
//  INVALIDPARAMETERVALUE_INSTALLBUNDLELIST = "InvalidParameterValue.InstallBundleList"
//  INVALIDPARAMETERVALUE_INSTALLPLUGINLIST = "InvalidParameterValue.InstallPluginList"
//  INVALIDPARAMETERVALUE_INVALIDJDK = "InvalidParameterValue.InvalidJDK"
//  INVALIDPARAMETERVALUE_PLUGINTYPE = "InvalidParameterValue.PluginType"
//  INVALIDPARAMETERVALUE_REMOVEBUNDLELIST = "InvalidParameterValue.RemoveBundleList"
//  INVALIDPARAMETERVALUE_REMOVEPLUGINLIST = "InvalidParameterValue.RemovePluginList"
//  INVALIDPARAMETERVALUE_RENEWFLAG = "InvalidParameterValue.RenewFlag"
//  INVALIDPARAMETERVALUE_TARGETPATH = "InvalidParameterValue.TargetPath"
//  INVALIDPARAMETERVALUE_TIMEUNIT = "InvalidParameterValue.TimeUnit"
//  INVALIDPARAMETERVALUE_UPGRADEMODE = "InvalidParameterValue.UpgradeMode"
//  LIMITEXCEEDED_CLUSTERNUM = "LimitExceeded.ClusterNum"
//  LIMITEXCEEDED_DIAGNOSECOUNT = "LimitExceeded.DiagnoseCount"
//  LIMITEXCEEDED_DISKCOUNT = "LimitExceeded.DiskCount"
//  LIMITEXCEEDED_NODENUMORINDICES = "LimitExceeded.NodeNumOrIndices"
//  LIMITEXCEEDED_PLUGININSTALL = "LimitExceeded.PluginInstall"
//  REGIONERROR = "RegionError"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_DIAGNOSE = "ResourceInUse.Diagnose"
//  RESOURCEINUSE_ORDER = "ResourceInUse.Order"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_CVM = "ResourceInsufficient.CVM"
//  RESOURCEINSUFFICIENT_CVMQUOTA = "ResourceInsufficient.CVMQuota"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
//  RESOURCENOTFOUND_ACCOUNTINFONOTFOUND = "ResourceNotFound.AccountInfoNotFound"
//  RESOURCENOTFOUND_BARADMETRICNOTFOUND = "ResourceNotFound.BaradMetricNotFound"
//  RESOURCENOTFOUND_CAMINFONOTFOUND = "ResourceNotFound.CAMInfoNotFound"
//  RESOURCENOTFOUND_CLBNOTFOUND = "ResourceNotFound.CLBNotFound"
//  RESOURCENOTFOUND_COSINFONOTFOUND = "ResourceNotFound.CosInfoNotFound"
//  RESOURCENOTFOUND_DIAGNOSENOTFOUND = "ResourceNotFound.DiagnoseNotFound"
//  RESOURCENOTFOUND_KMSNOTFOUND = "ResourceNotFound.KMSNotFound"
//  RESOURCENOTFOUND_STSNOTFOUND = "ResourceNotFound.STSNotFound"
//  RESOURCENOTFOUND_SECURITYGROUPNOTFOUND = "ResourceNotFound.SecurityGroupNotFound"
//  RESOURCENOTFOUND_TRADECGWNOTFOUND = "ResourceNotFound.TradeCgwNotFound"
//  RESOURCENOTFOUND_VPCINFONOTFOUND = "ResourceNotFound.VPCInfoNotFound"
//  UNSUPPORTEDOPERATION_AUDITSTATUSCONFLICT = "UnsupportedOperation.AuditStatusConflict"
//  UNSUPPORTEDOPERATION_COSBACKUP = "UnsupportedOperation.CosBackUp"
//  UNSUPPORTEDOPERATION_DIAGNOSEJOB = "UnsupportedOperation.DiagnoseJob"
//  UNSUPPORTEDOPERATION_DIAGNOSENOTOPEN = "UnsupportedOperation.DiagnoseNotOpen"
//  UNSUPPORTEDOPERATION_DISKUSE = "UnsupportedOperation.DiskUse"
//  UNSUPPORTEDOPERATION_EDITLISTLENGTH = "UnsupportedOperation.EditListLength"
//  UNSUPPORTEDOPERATION_MIGRATE = "UnsupportedOperation.Migrate"
//  UNSUPPORTEDOPERATION_PASSLOGSTASHID = "UnsupportedOperation.PassLogstashId"
//  UNSUPPORTEDOPERATION_PLUGIN = "UnsupportedOperation.Plugin"
//  UNSUPPORTEDOPERATION_PROTOCOL = "UnsupportedOperation.Protocol"
//  UNSUPPORTEDOPERATION_READRATE = "UnsupportedOperation.ReadRate"
//  UNSUPPORTEDOPERATION_RESTARTMODE = "UnsupportedOperation.RestartMode"
//  UNSUPPORTEDOPERATION_STATUSNOTNORMAL = "UnsupportedOperation.StatusNotNormal"
//  UNSUPPORTEDOPERATION_UPDATEDISKENCRYPT = "UnsupportedOperation.UpdateDiskEncrypt"
//  UNSUPPORTEDOPERATION_UPDATEDISKTYPE = "UnsupportedOperation.UpdateDiskType"
//  UNSUPPORTEDOPERATION_UPDATENODENUMANDNODETYPE = "UnsupportedOperation.UpdateNodeNumAndNodeType"
//  UNSUPPORTEDOPERATION_WEBSERVICETYPE = "UnsupportedOperation.WebServiceType"
//  UNSUPPORTEDOPERATION_WRITERATE = "UnsupportedOperation.WriteRate"
func (c *Client) CreateLogstashInstanceWithContext(ctx context.Context, request *CreateLogstashInstanceRequest) (response *CreateLogstashInstanceResponse, err error) {
    if request == nil {
        request = NewCreateLogstashInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLogstashInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLogstashInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateServerlessInstanceRequest() (request *CreateServerlessInstanceRequest) {
    request = &CreateServerlessInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "CreateServerlessInstance")
    
    
    return
}

func NewCreateServerlessInstanceResponse() (response *CreateServerlessInstanceResponse) {
    response = &CreateServerlessInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateServerlessInstance
// 创建Serverless索引
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION_CLUSTERRESOURCELIMITERROR = "FailedOperation.ClusterResourceLimitError"
//  FAILEDOPERATION_DISKCOUNTPARAMERROR = "FailedOperation.DiskCountParamError"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  FAILEDOPERATION_NOTAUTHENTICATED = "FailedOperation.NotAuthenticated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INDEXNAMEEXIST = "InvalidParameter.IndexNameExist"
//  INVALIDPARAMETER_INDEXNAMEINVALID = "InvalidParameter.IndexNameInvalid"
//  INVALIDPARAMETER_USERNAMEEXIST = "InvalidParameter.UsernameExist"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_INDEXCOUNT = "LimitExceeded.IndexCount"
//  LIMITEXCEEDED_INDEXCREATE = "LimitExceeded.IndexCreate"
//  LIMITEXCEEDED_SPACECOUNT = "LimitExceeded.SpaceCount"
//  LIMITEXCEEDED_SPACECREATE = "LimitExceeded.SpaceCreate"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_HIDDENZONE = "ResourceInsufficient.HiddenZone"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
func (c *Client) CreateServerlessInstance(request *CreateServerlessInstanceRequest) (response *CreateServerlessInstanceResponse, err error) {
    return c.CreateServerlessInstanceWithContext(context.Background(), request)
}

// CreateServerlessInstance
// 创建Serverless索引
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION_CLUSTERRESOURCELIMITERROR = "FailedOperation.ClusterResourceLimitError"
//  FAILEDOPERATION_DISKCOUNTPARAMERROR = "FailedOperation.DiskCountParamError"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  FAILEDOPERATION_NOTAUTHENTICATED = "FailedOperation.NotAuthenticated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INDEXNAMEEXIST = "InvalidParameter.IndexNameExist"
//  INVALIDPARAMETER_INDEXNAMEINVALID = "InvalidParameter.IndexNameInvalid"
//  INVALIDPARAMETER_USERNAMEEXIST = "InvalidParameter.UsernameExist"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_INDEXCOUNT = "LimitExceeded.IndexCount"
//  LIMITEXCEEDED_INDEXCREATE = "LimitExceeded.IndexCreate"
//  LIMITEXCEEDED_SPACECOUNT = "LimitExceeded.SpaceCount"
//  LIMITEXCEEDED_SPACECREATE = "LimitExceeded.SpaceCreate"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_HIDDENZONE = "ResourceInsufficient.HiddenZone"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
func (c *Client) CreateServerlessInstanceWithContext(ctx context.Context, request *CreateServerlessInstanceRequest) (response *CreateServerlessInstanceResponse, err error) {
    if request == nil {
        request = NewCreateServerlessInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateServerlessInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateServerlessInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateServerlessSpaceV2Request() (request *CreateServerlessSpaceV2Request) {
    request = &CreateServerlessSpaceV2Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "CreateServerlessSpaceV2")
    
    
    return
}

func NewCreateServerlessSpaceV2Response() (response *CreateServerlessSpaceV2Response) {
    response = &CreateServerlessSpaceV2Response{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateServerlessSpaceV2
// 创建Serverless索引空间
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION_CLUSTERRESOURCELIMITERROR = "FailedOperation.ClusterResourceLimitError"
//  FAILEDOPERATION_DISKCOUNTPARAMERROR = "FailedOperation.DiskCountParamError"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  FAILEDOPERATION_NOTAUTHENTICATED = "FailedOperation.NotAuthenticated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SPACENAMEEXIST = "InvalidParameter.SpaceNameExist"
//  INVALIDPARAMETER_SPACENAMEINVALID = "InvalidParameter.SpaceNameInvalid"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_INDEXCOUNT = "LimitExceeded.IndexCount"
//  LIMITEXCEEDED_INDEXCREATE = "LimitExceeded.IndexCreate"
//  LIMITEXCEEDED_SPACECOUNT = "LimitExceeded.SpaceCount"
//  LIMITEXCEEDED_SPACECREATE = "LimitExceeded.SpaceCreate"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_HIDDENZONE = "ResourceInsufficient.HiddenZone"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
func (c *Client) CreateServerlessSpaceV2(request *CreateServerlessSpaceV2Request) (response *CreateServerlessSpaceV2Response, err error) {
    return c.CreateServerlessSpaceV2WithContext(context.Background(), request)
}

// CreateServerlessSpaceV2
// 创建Serverless索引空间
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION_CLUSTERRESOURCELIMITERROR = "FailedOperation.ClusterResourceLimitError"
//  FAILEDOPERATION_DISKCOUNTPARAMERROR = "FailedOperation.DiskCountParamError"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  FAILEDOPERATION_NOTAUTHENTICATED = "FailedOperation.NotAuthenticated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SPACENAMEEXIST = "InvalidParameter.SpaceNameExist"
//  INVALIDPARAMETER_SPACENAMEINVALID = "InvalidParameter.SpaceNameInvalid"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_INDEXCOUNT = "LimitExceeded.IndexCount"
//  LIMITEXCEEDED_INDEXCREATE = "LimitExceeded.IndexCreate"
//  LIMITEXCEEDED_SPACECOUNT = "LimitExceeded.SpaceCount"
//  LIMITEXCEEDED_SPACECREATE = "LimitExceeded.SpaceCreate"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_HIDDENZONE = "ResourceInsufficient.HiddenZone"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
func (c *Client) CreateServerlessSpaceV2WithContext(ctx context.Context, request *CreateServerlessSpaceV2Request) (response *CreateServerlessSpaceV2Response, err error) {
    if request == nil {
        request = NewCreateServerlessSpaceV2Request()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateServerlessSpaceV2 require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateServerlessSpaceV2Response()
    err = c.Send(request, response)
    return
}

func NewDeleteClusterSnapshotRequest() (request *DeleteClusterSnapshotRequest) {
    request = &DeleteClusterSnapshotRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "DeleteClusterSnapshot")
    
    
    return
}

func NewDeleteClusterSnapshotResponse() (response *DeleteClusterSnapshotResponse) {
    response = &DeleteClusterSnapshotResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteClusterSnapshot
// 删除快照仓库里备份的快照
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION_CLUSTERRESOURCELIMITERROR = "FailedOperation.ClusterResourceLimitError"
//  FAILEDOPERATION_DISKCOUNTPARAMERROR = "FailedOperation.DiskCountParamError"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  FAILEDOPERATION_NOTAUTHENTICATED = "FailedOperation.NotAuthenticated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SPACENAMEEXIST = "InvalidParameter.SpaceNameExist"
//  INVALIDPARAMETER_SPACENAMEINVALID = "InvalidParameter.SpaceNameInvalid"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_INDEXCOUNT = "LimitExceeded.IndexCount"
//  LIMITEXCEEDED_INDEXCREATE = "LimitExceeded.IndexCreate"
//  LIMITEXCEEDED_SPACECOUNT = "LimitExceeded.SpaceCount"
//  LIMITEXCEEDED_SPACECREATE = "LimitExceeded.SpaceCreate"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_HIDDENZONE = "ResourceInsufficient.HiddenZone"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
func (c *Client) DeleteClusterSnapshot(request *DeleteClusterSnapshotRequest) (response *DeleteClusterSnapshotResponse, err error) {
    return c.DeleteClusterSnapshotWithContext(context.Background(), request)
}

// DeleteClusterSnapshot
// 删除快照仓库里备份的快照
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION_CLUSTERRESOURCELIMITERROR = "FailedOperation.ClusterResourceLimitError"
//  FAILEDOPERATION_DISKCOUNTPARAMERROR = "FailedOperation.DiskCountParamError"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  FAILEDOPERATION_NOTAUTHENTICATED = "FailedOperation.NotAuthenticated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SPACENAMEEXIST = "InvalidParameter.SpaceNameExist"
//  INVALIDPARAMETER_SPACENAMEINVALID = "InvalidParameter.SpaceNameInvalid"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_INDEXCOUNT = "LimitExceeded.IndexCount"
//  LIMITEXCEEDED_INDEXCREATE = "LimitExceeded.IndexCreate"
//  LIMITEXCEEDED_SPACECOUNT = "LimitExceeded.SpaceCount"
//  LIMITEXCEEDED_SPACECREATE = "LimitExceeded.SpaceCreate"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_HIDDENZONE = "ResourceInsufficient.HiddenZone"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
func (c *Client) DeleteClusterSnapshotWithContext(ctx context.Context, request *DeleteClusterSnapshotRequest) (response *DeleteClusterSnapshotResponse, err error) {
    if request == nil {
        request = NewDeleteClusterSnapshotRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteClusterSnapshot require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteClusterSnapshotResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteIndexRequest() (request *DeleteIndexRequest) {
    request = &DeleteIndexRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "DeleteIndex")
    
    
    return
}

func NewDeleteIndexResponse() (response *DeleteIndexResponse) {
    response = &DeleteIndexResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteIndex
// 删除索引
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDESVERSION = "InvalidParameter.InvalidEsVersion"
//  INVALIDPARAMETER_INVALIDINDEXTYPE = "InvalidParameter.InvalidIndexType"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  RESOURCENOTFOUND_CLUSTERINFONOTFOUND = "ResourceNotFound.ClusterInfoNotFound"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
func (c *Client) DeleteIndex(request *DeleteIndexRequest) (response *DeleteIndexResponse, err error) {
    return c.DeleteIndexWithContext(context.Background(), request)
}

// DeleteIndex
// 删除索引
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDESVERSION = "InvalidParameter.InvalidEsVersion"
//  INVALIDPARAMETER_INVALIDINDEXTYPE = "InvalidParameter.InvalidIndexType"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  RESOURCENOTFOUND_CLUSTERINFONOTFOUND = "ResourceNotFound.ClusterInfoNotFound"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
func (c *Client) DeleteIndexWithContext(ctx context.Context, request *DeleteIndexRequest) (response *DeleteIndexResponse, err error) {
    if request == nil {
        request = NewDeleteIndexRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteIndex require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteIndexResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteInstanceRequest() (request *DeleteInstanceRequest) {
    request = &DeleteInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "DeleteInstance")
    
    
    return
}

func NewDeleteInstanceResponse() (response *DeleteInstanceResponse) {
    response = &DeleteInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteInstance
// 销毁集群实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_REFUNDERROR = "FailedOperation.RefundError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  INVALIDPARAMETERVALUE_CHARGETYPE = "InvalidParameterValue.ChargeType"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
//  RESOURCENOTFOUND_TRADECGWNOTFOUND = "ResourceNotFound.TradeCgwNotFound"
//  UNSUPPORTEDOPERATION_STATUSNOTSUPPORT = "UnsupportedOperation.StatusNotSupport"
func (c *Client) DeleteInstance(request *DeleteInstanceRequest) (response *DeleteInstanceResponse, err error) {
    return c.DeleteInstanceWithContext(context.Background(), request)
}

// DeleteInstance
// 销毁集群实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_REFUNDERROR = "FailedOperation.RefundError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  INVALIDPARAMETERVALUE_CHARGETYPE = "InvalidParameterValue.ChargeType"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
//  RESOURCENOTFOUND_TRADECGWNOTFOUND = "ResourceNotFound.TradeCgwNotFound"
//  UNSUPPORTEDOPERATION_STATUSNOTSUPPORT = "UnsupportedOperation.StatusNotSupport"
func (c *Client) DeleteInstanceWithContext(ctx context.Context, request *DeleteInstanceRequest) (response *DeleteInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLogstashInstanceRequest() (request *DeleteLogstashInstanceRequest) {
    request = &DeleteLogstashInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "DeleteLogstashInstance")
    
    
    return
}

func NewDeleteLogstashInstanceResponse() (response *DeleteLogstashInstanceResponse) {
    response = &DeleteLogstashInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLogstashInstance
// 用于删除Logstash实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteLogstashInstance(request *DeleteLogstashInstanceRequest) (response *DeleteLogstashInstanceResponse, err error) {
    return c.DeleteLogstashInstanceWithContext(context.Background(), request)
}

// DeleteLogstashInstance
// 用于删除Logstash实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteLogstashInstanceWithContext(ctx context.Context, request *DeleteLogstashInstanceRequest) (response *DeleteLogstashInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteLogstashInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLogstashInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLogstashInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLogstashPipelinesRequest() (request *DeleteLogstashPipelinesRequest) {
    request = &DeleteLogstashPipelinesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "DeleteLogstashPipelines")
    
    
    return
}

func NewDeleteLogstashPipelinesResponse() (response *DeleteLogstashPipelinesResponse) {
    response = &DeleteLogstashPipelinesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLogstashPipelines
// 用于批量删除Logstash管道
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DeleteLogstashPipelines(request *DeleteLogstashPipelinesRequest) (response *DeleteLogstashPipelinesResponse, err error) {
    return c.DeleteLogstashPipelinesWithContext(context.Background(), request)
}

// DeleteLogstashPipelines
// 用于批量删除Logstash管道
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DeleteLogstashPipelinesWithContext(ctx context.Context, request *DeleteLogstashPipelinesRequest) (response *DeleteLogstashPipelinesResponse, err error) {
    if request == nil {
        request = NewDeleteLogstashPipelinesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLogstashPipelines require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLogstashPipelinesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteServerlessInstanceRequest() (request *DeleteServerlessInstanceRequest) {
    request = &DeleteServerlessInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "DeleteServerlessInstance")
    
    
    return
}

func NewDeleteServerlessInstanceResponse() (response *DeleteServerlessInstanceResponse) {
    response = &DeleteServerlessInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteServerlessInstance
// 删除Serverless索引
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION_CLUSTERRESOURCELIMITERROR = "FailedOperation.ClusterResourceLimitError"
//  FAILEDOPERATION_DISKCOUNTPARAMERROR = "FailedOperation.DiskCountParamError"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  FAILEDOPERATION_NOTAUTHENTICATED = "FailedOperation.NotAuthenticated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_HIDDENZONE = "ResourceInsufficient.HiddenZone"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
func (c *Client) DeleteServerlessInstance(request *DeleteServerlessInstanceRequest) (response *DeleteServerlessInstanceResponse, err error) {
    return c.DeleteServerlessInstanceWithContext(context.Background(), request)
}

// DeleteServerlessInstance
// 删除Serverless索引
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION_CLUSTERRESOURCELIMITERROR = "FailedOperation.ClusterResourceLimitError"
//  FAILEDOPERATION_DISKCOUNTPARAMERROR = "FailedOperation.DiskCountParamError"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  FAILEDOPERATION_NOTAUTHENTICATED = "FailedOperation.NotAuthenticated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_HIDDENZONE = "ResourceInsufficient.HiddenZone"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
func (c *Client) DeleteServerlessInstanceWithContext(ctx context.Context, request *DeleteServerlessInstanceRequest) (response *DeleteServerlessInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteServerlessInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteServerlessInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteServerlessInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteServerlessSpaceUserRequest() (request *DeleteServerlessSpaceUserRequest) {
    request = &DeleteServerlessSpaceUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "DeleteServerlessSpaceUser")
    
    
    return
}

func NewDeleteServerlessSpaceUserResponse() (response *DeleteServerlessSpaceUserResponse) {
    response = &DeleteServerlessSpaceUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteServerlessSpaceUser
// 删除Serverless空间子用户
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION_CLUSTERRESOURCELIMITERROR = "FailedOperation.ClusterResourceLimitError"
//  FAILEDOPERATION_DISKCOUNTPARAMERROR = "FailedOperation.DiskCountParamError"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  FAILEDOPERATION_NOTAUTHENTICATED = "FailedOperation.NotAuthenticated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_USERNAMEEXIST = "InvalidParameter.UsernameExist"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_INDEXCOUNT = "LimitExceeded.IndexCount"
//  LIMITEXCEEDED_INDEXCREATE = "LimitExceeded.IndexCreate"
//  LIMITEXCEEDED_SPACECOUNT = "LimitExceeded.SpaceCount"
//  LIMITEXCEEDED_SPACECREATE = "LimitExceeded.SpaceCreate"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_HIDDENZONE = "ResourceInsufficient.HiddenZone"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
func (c *Client) DeleteServerlessSpaceUser(request *DeleteServerlessSpaceUserRequest) (response *DeleteServerlessSpaceUserResponse, err error) {
    return c.DeleteServerlessSpaceUserWithContext(context.Background(), request)
}

// DeleteServerlessSpaceUser
// 删除Serverless空间子用户
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION_CLUSTERRESOURCELIMITERROR = "FailedOperation.ClusterResourceLimitError"
//  FAILEDOPERATION_DISKCOUNTPARAMERROR = "FailedOperation.DiskCountParamError"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  FAILEDOPERATION_NOTAUTHENTICATED = "FailedOperation.NotAuthenticated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_USERNAMEEXIST = "InvalidParameter.UsernameExist"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_INDEXCOUNT = "LimitExceeded.IndexCount"
//  LIMITEXCEEDED_INDEXCREATE = "LimitExceeded.IndexCreate"
//  LIMITEXCEEDED_SPACECOUNT = "LimitExceeded.SpaceCount"
//  LIMITEXCEEDED_SPACECREATE = "LimitExceeded.SpaceCreate"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_HIDDENZONE = "ResourceInsufficient.HiddenZone"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
func (c *Client) DeleteServerlessSpaceUserWithContext(ctx context.Context, request *DeleteServerlessSpaceUserRequest) (response *DeleteServerlessSpaceUserResponse, err error) {
    if request == nil {
        request = NewDeleteServerlessSpaceUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteServerlessSpaceUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteServerlessSpaceUserResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterSnapshotRequest() (request *DescribeClusterSnapshotRequest) {
    request = &DescribeClusterSnapshotRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "DescribeClusterSnapshot")
    
    
    return
}

func NewDescribeClusterSnapshotResponse() (response *DescribeClusterSnapshotResponse) {
    response = &DescribeClusterSnapshotResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterSnapshot
// 获取快照备份列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_COSINFONOTFOUND = "ResourceNotFound.CosInfoNotFound"
func (c *Client) DescribeClusterSnapshot(request *DescribeClusterSnapshotRequest) (response *DescribeClusterSnapshotResponse, err error) {
    return c.DescribeClusterSnapshotWithContext(context.Background(), request)
}

// DescribeClusterSnapshot
// 获取快照备份列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_COSINFONOTFOUND = "ResourceNotFound.CosInfoNotFound"
func (c *Client) DescribeClusterSnapshotWithContext(ctx context.Context, request *DescribeClusterSnapshotRequest) (response *DescribeClusterSnapshotResponse, err error) {
    if request == nil {
        request = NewDescribeClusterSnapshotRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterSnapshot require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterSnapshotResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDiagnoseRequest() (request *DescribeDiagnoseRequest) {
    request = &DescribeDiagnoseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "DescribeDiagnose")
    
    
    return
}

func NewDescribeDiagnoseResponse() (response *DescribeDiagnoseResponse) {
    response = &DescribeDiagnoseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDiagnose
// 查询智能运维诊断结果报告
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  INVALIDPARAMETER_INVALIDREGION = "InvalidParameter.InvalidRegion"
//  INVALIDPARAMETER_INVALIDTIMEPARAM = "InvalidParameter.InvalidTimeParam"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
//  RESOURCENOTFOUND_DIAGNOSENOTFOUND = "ResourceNotFound.DiagnoseNotFound"
//  UNSUPPORTEDOPERATION_DIAGNOSENOTOPEN = "UnsupportedOperation.DiagnoseNotOpen"
func (c *Client) DescribeDiagnose(request *DescribeDiagnoseRequest) (response *DescribeDiagnoseResponse, err error) {
    return c.DescribeDiagnoseWithContext(context.Background(), request)
}

// DescribeDiagnose
// 查询智能运维诊断结果报告
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  INVALIDPARAMETER_INVALIDREGION = "InvalidParameter.InvalidRegion"
//  INVALIDPARAMETER_INVALIDTIMEPARAM = "InvalidParameter.InvalidTimeParam"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
//  RESOURCENOTFOUND_DIAGNOSENOTFOUND = "ResourceNotFound.DiagnoseNotFound"
//  UNSUPPORTEDOPERATION_DIAGNOSENOTOPEN = "UnsupportedOperation.DiagnoseNotOpen"
func (c *Client) DescribeDiagnoseWithContext(ctx context.Context, request *DescribeDiagnoseRequest) (response *DescribeDiagnoseResponse, err error) {
    if request == nil {
        request = NewDescribeDiagnoseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDiagnose require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDiagnoseResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIndexListRequest() (request *DescribeIndexListRequest) {
    request = &DescribeIndexListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "DescribeIndexList")
    
    
    return
}

func NewDescribeIndexListResponse() (response *DescribeIndexListResponse) {
    response = &DescribeIndexListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIndexList
// 获取索引列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDESVERSION = "InvalidParameter.InvalidEsVersion"
//  INVALIDPARAMETER_INVALIDINDEXTYPE = "InvalidParameter.InvalidIndexType"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  RESOURCENOTFOUND_CLUSTERINFONOTFOUND = "ResourceNotFound.ClusterInfoNotFound"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
func (c *Client) DescribeIndexList(request *DescribeIndexListRequest) (response *DescribeIndexListResponse, err error) {
    return c.DescribeIndexListWithContext(context.Background(), request)
}

// DescribeIndexList
// 获取索引列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDESVERSION = "InvalidParameter.InvalidEsVersion"
//  INVALIDPARAMETER_INVALIDINDEXTYPE = "InvalidParameter.InvalidIndexType"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  RESOURCENOTFOUND_CLUSTERINFONOTFOUND = "ResourceNotFound.ClusterInfoNotFound"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
func (c *Client) DescribeIndexListWithContext(ctx context.Context, request *DescribeIndexListRequest) (response *DescribeIndexListResponse, err error) {
    if request == nil {
        request = NewDescribeIndexListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIndexList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIndexListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIndexMetaRequest() (request *DescribeIndexMetaRequest) {
    request = &DescribeIndexMetaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "DescribeIndexMeta")
    
    
    return
}

func NewDescribeIndexMetaResponse() (response *DescribeIndexMetaResponse) {
    response = &DescribeIndexMetaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIndexMeta
// 获取索引元数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDESVERSION = "InvalidParameter.InvalidEsVersion"
//  INVALIDPARAMETER_INVALIDINDEXNAME = "InvalidParameter.InvalidIndexName"
//  INVALIDPARAMETER_INVALIDINDEXTYPE = "InvalidParameter.InvalidIndexType"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  RESOURCENOTFOUND_CLUSTERINFONOTFOUND = "ResourceNotFound.ClusterInfoNotFound"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
func (c *Client) DescribeIndexMeta(request *DescribeIndexMetaRequest) (response *DescribeIndexMetaResponse, err error) {
    return c.DescribeIndexMetaWithContext(context.Background(), request)
}

// DescribeIndexMeta
// 获取索引元数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDESVERSION = "InvalidParameter.InvalidEsVersion"
//  INVALIDPARAMETER_INVALIDINDEXNAME = "InvalidParameter.InvalidIndexName"
//  INVALIDPARAMETER_INVALIDINDEXTYPE = "InvalidParameter.InvalidIndexType"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  RESOURCENOTFOUND_CLUSTERINFONOTFOUND = "ResourceNotFound.ClusterInfoNotFound"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
func (c *Client) DescribeIndexMetaWithContext(ctx context.Context, request *DescribeIndexMetaRequest) (response *DescribeIndexMetaResponse, err error) {
    if request == nil {
        request = NewDescribeIndexMetaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIndexMeta require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIndexMetaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceLogsRequest() (request *DescribeInstanceLogsRequest) {
    request = &DescribeInstanceLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "DescribeInstanceLogs")
    
    
    return
}

func NewDescribeInstanceLogsResponse() (response *DescribeInstanceLogsResponse) {
    response = &DescribeInstanceLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceLogs
// 查询用户该地域下符合条件的ES集群的日志
//
// 可能返回的错误码:
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  INVALIDPARAMETER_INVALIDLIMIT = "InvalidParameter.InvalidLimit"
//  INVALIDPARAMETER_INVALIDLOGTYPE = "InvalidParameter.InvalidLogType"
//  INVALIDPARAMETER_INVALIDOFFSET = "InvalidParameter.InvalidOffset"
//  INVALIDPARAMETER_INVALIDORDERBYKEY = "InvalidParameter.InvalidOrderByKey"
//  INVALIDPARAMETER_INVALIDQUERYSTRING = "InvalidParameter.InvalidQueryString"
//  INVALIDPARAMETER_INVALIDREGION = "InvalidParameter.InvalidRegion"
//  INVALIDPARAMETER_INVALIDTIMEPARAM = "InvalidParameter.InvalidTimeParam"
//  RESOURCENOTFOUND_CLUSTERINFONOTFOUND = "ResourceNotFound.ClusterInfoNotFound"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
func (c *Client) DescribeInstanceLogs(request *DescribeInstanceLogsRequest) (response *DescribeInstanceLogsResponse, err error) {
    return c.DescribeInstanceLogsWithContext(context.Background(), request)
}

// DescribeInstanceLogs
// 查询用户该地域下符合条件的ES集群的日志
//
// 可能返回的错误码:
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  INVALIDPARAMETER_INVALIDLIMIT = "InvalidParameter.InvalidLimit"
//  INVALIDPARAMETER_INVALIDLOGTYPE = "InvalidParameter.InvalidLogType"
//  INVALIDPARAMETER_INVALIDOFFSET = "InvalidParameter.InvalidOffset"
//  INVALIDPARAMETER_INVALIDORDERBYKEY = "InvalidParameter.InvalidOrderByKey"
//  INVALIDPARAMETER_INVALIDQUERYSTRING = "InvalidParameter.InvalidQueryString"
//  INVALIDPARAMETER_INVALIDREGION = "InvalidParameter.InvalidRegion"
//  INVALIDPARAMETER_INVALIDTIMEPARAM = "InvalidParameter.InvalidTimeParam"
//  RESOURCENOTFOUND_CLUSTERINFONOTFOUND = "ResourceNotFound.ClusterInfoNotFound"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
func (c *Client) DescribeInstanceLogsWithContext(ctx context.Context, request *DescribeInstanceLogsRequest) (response *DescribeInstanceLogsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceOperationsRequest() (request *DescribeInstanceOperationsRequest) {
    request = &DescribeInstanceOperationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "DescribeInstanceOperations")
    
    
    return
}

func NewDescribeInstanceOperationsResponse() (response *DescribeInstanceOperationsResponse) {
    response = &DescribeInstanceOperationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceOperations
// 查询实例指定条件下的操作记录
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  INVALIDPARAMETER_INVALIDLIMIT = "InvalidParameter.InvalidLimit"
//  INVALIDPARAMETER_INVALIDTIMEPARAM = "InvalidParameter.InvalidTimeParam"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
//  UNSUPPORTEDOPERATION_STATUSNOTSUPPORT = "UnsupportedOperation.StatusNotSupport"
func (c *Client) DescribeInstanceOperations(request *DescribeInstanceOperationsRequest) (response *DescribeInstanceOperationsResponse, err error) {
    return c.DescribeInstanceOperationsWithContext(context.Background(), request)
}

// DescribeInstanceOperations
// 查询实例指定条件下的操作记录
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  INVALIDPARAMETER_INVALIDLIMIT = "InvalidParameter.InvalidLimit"
//  INVALIDPARAMETER_INVALIDTIMEPARAM = "InvalidParameter.InvalidTimeParam"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
//  UNSUPPORTEDOPERATION_STATUSNOTSUPPORT = "UnsupportedOperation.StatusNotSupport"
func (c *Client) DescribeInstanceOperationsWithContext(ctx context.Context, request *DescribeInstanceOperationsRequest) (response *DescribeInstanceOperationsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceOperationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceOperations require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceOperationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancePluginListRequest() (request *DescribeInstancePluginListRequest) {
    request = &DescribeInstancePluginListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "DescribeInstancePluginList")
    
    
    return
}

func NewDescribeInstancePluginListResponse() (response *DescribeInstancePluginListResponse) {
    response = &DescribeInstancePluginListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstancePluginList
// 查询实例插件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  INVALIDPARAMETERVALUE_CONFIGINFO = "InvalidParameterValue.ConfigInfo"
//  RESOURCENOTFOUND_CLUSTERINFONOTFOUND = "ResourceNotFound.ClusterInfoNotFound"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
func (c *Client) DescribeInstancePluginList(request *DescribeInstancePluginListRequest) (response *DescribeInstancePluginListResponse, err error) {
    return c.DescribeInstancePluginListWithContext(context.Background(), request)
}

// DescribeInstancePluginList
// 查询实例插件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  INVALIDPARAMETERVALUE_CONFIGINFO = "InvalidParameterValue.ConfigInfo"
//  RESOURCENOTFOUND_CLUSTERINFONOTFOUND = "ResourceNotFound.ClusterInfoNotFound"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
func (c *Client) DescribeInstancePluginListWithContext(ctx context.Context, request *DescribeInstancePluginListRequest) (response *DescribeInstancePluginListResponse, err error) {
    if request == nil {
        request = NewDescribeInstancePluginListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstancePluginList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancePluginListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "DescribeInstances")
    
    
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstances
// 查询用户该地域下符合条件的所有实例
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHDESCRIBEINSTANCES = "AuthFailure.UnAuthDescribeInstances"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION_GETTAGINFOERROR = "FailedOperation.GetTagInfoError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDIP = "InvalidParameter.InvalidIp"
//  INVALIDPARAMETER_INVALIDIPLIST = "InvalidParameter.InvalidIpList"
//  INVALIDPARAMETER_INVALIDLIMIT = "InvalidParameter.InvalidLimit"
//  INVALIDPARAMETER_INVALIDNODETYPE = "InvalidParameter.InvalidNodeType"
//  INVALIDPARAMETER_INVALIDOFFSET = "InvalidParameter.InvalidOffset"
//  INVALIDPARAMETER_INVALIDORDERBYKEY = "InvalidParameter.InvalidOrderByKey"
//  INVALIDPARAMETER_INVALIDORDERBYTYPE = "InvalidParameter.InvalidOrderByType"
//  INVALIDPARAMETER_INVALIDREGION = "InvalidParameter.InvalidRegion"
//  INVALIDPARAMETER_INVALIDTAGINFO = "InvalidParameter.InvalidTagInfo"
//  INVALIDPARAMETER_INVALIDTAGLIST = "InvalidParameter.InvalidTagList"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  RESOURCENOTFOUND_CAMINFONOTFOUND = "ResourceNotFound.CAMInfoNotFound"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    return c.DescribeInstancesWithContext(context.Background(), request)
}

// DescribeInstances
// 查询用户该地域下符合条件的所有实例
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHDESCRIBEINSTANCES = "AuthFailure.UnAuthDescribeInstances"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION_GETTAGINFOERROR = "FailedOperation.GetTagInfoError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDIP = "InvalidParameter.InvalidIp"
//  INVALIDPARAMETER_INVALIDIPLIST = "InvalidParameter.InvalidIpList"
//  INVALIDPARAMETER_INVALIDLIMIT = "InvalidParameter.InvalidLimit"
//  INVALIDPARAMETER_INVALIDNODETYPE = "InvalidParameter.InvalidNodeType"
//  INVALIDPARAMETER_INVALIDOFFSET = "InvalidParameter.InvalidOffset"
//  INVALIDPARAMETER_INVALIDORDERBYKEY = "InvalidParameter.InvalidOrderByKey"
//  INVALIDPARAMETER_INVALIDORDERBYTYPE = "InvalidParameter.InvalidOrderByType"
//  INVALIDPARAMETER_INVALIDREGION = "InvalidParameter.InvalidRegion"
//  INVALIDPARAMETER_INVALIDTAGINFO = "InvalidParameter.InvalidTagInfo"
//  INVALIDPARAMETER_INVALIDTAGLIST = "InvalidParameter.InvalidTagList"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  RESOURCENOTFOUND_CAMINFONOTFOUND = "ResourceNotFound.CAMInfoNotFound"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
func (c *Client) DescribeInstancesWithContext(ctx context.Context, request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogstashInstanceLogsRequest() (request *DescribeLogstashInstanceLogsRequest) {
    request = &DescribeLogstashInstanceLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "DescribeLogstashInstanceLogs")
    
    
    return
}

func NewDescribeLogstashInstanceLogsResponse() (response *DescribeLogstashInstanceLogsResponse) {
    response = &DescribeLogstashInstanceLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLogstashInstanceLogs
// 查询用户该地域下符合条件的Logstash实例的日志
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDQUERYSTRING = "InvalidParameter.InvalidQueryString"
func (c *Client) DescribeLogstashInstanceLogs(request *DescribeLogstashInstanceLogsRequest) (response *DescribeLogstashInstanceLogsResponse, err error) {
    return c.DescribeLogstashInstanceLogsWithContext(context.Background(), request)
}

// DescribeLogstashInstanceLogs
// 查询用户该地域下符合条件的Logstash实例的日志
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDQUERYSTRING = "InvalidParameter.InvalidQueryString"
func (c *Client) DescribeLogstashInstanceLogsWithContext(ctx context.Context, request *DescribeLogstashInstanceLogsRequest) (response *DescribeLogstashInstanceLogsResponse, err error) {
    if request == nil {
        request = NewDescribeLogstashInstanceLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogstashInstanceLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogstashInstanceLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogstashInstanceOperationsRequest() (request *DescribeLogstashInstanceOperationsRequest) {
    request = &DescribeLogstashInstanceOperationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "DescribeLogstashInstanceOperations")
    
    
    return
}

func NewDescribeLogstashInstanceOperationsResponse() (response *DescribeLogstashInstanceOperationsResponse) {
    response = &DescribeLogstashInstanceOperationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLogstashInstanceOperations
// 查询实例指定条件下的操作记录
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeLogstashInstanceOperations(request *DescribeLogstashInstanceOperationsRequest) (response *DescribeLogstashInstanceOperationsResponse, err error) {
    return c.DescribeLogstashInstanceOperationsWithContext(context.Background(), request)
}

// DescribeLogstashInstanceOperations
// 查询实例指定条件下的操作记录
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeLogstashInstanceOperationsWithContext(ctx context.Context, request *DescribeLogstashInstanceOperationsRequest) (response *DescribeLogstashInstanceOperationsResponse, err error) {
    if request == nil {
        request = NewDescribeLogstashInstanceOperationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogstashInstanceOperations require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogstashInstanceOperationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogstashInstancesRequest() (request *DescribeLogstashInstancesRequest) {
    request = &DescribeLogstashInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "DescribeLogstashInstances")
    
    
    return
}

func NewDescribeLogstashInstancesResponse() (response *DescribeLogstashInstancesResponse) {
    response = &DescribeLogstashInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLogstashInstances
// 查询用户该地域下符合条件的所有Logstash实例
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeLogstashInstances(request *DescribeLogstashInstancesRequest) (response *DescribeLogstashInstancesResponse, err error) {
    return c.DescribeLogstashInstancesWithContext(context.Background(), request)
}

// DescribeLogstashInstances
// 查询用户该地域下符合条件的所有Logstash实例
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeLogstashInstancesWithContext(ctx context.Context, request *DescribeLogstashInstancesRequest) (response *DescribeLogstashInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeLogstashInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogstashInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogstashInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogstashPipelinesRequest() (request *DescribeLogstashPipelinesRequest) {
    request = &DescribeLogstashPipelinesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "DescribeLogstashPipelines")
    
    
    return
}

func NewDescribeLogstashPipelinesResponse() (response *DescribeLogstashPipelinesResponse) {
    response = &DescribeLogstashPipelinesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLogstashPipelines
// 用于获取Logstash实例管道列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeLogstashPipelines(request *DescribeLogstashPipelinesRequest) (response *DescribeLogstashPipelinesResponse, err error) {
    return c.DescribeLogstashPipelinesWithContext(context.Background(), request)
}

// DescribeLogstashPipelines
// 用于获取Logstash实例管道列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeLogstashPipelinesWithContext(ctx context.Context, request *DescribeLogstashPipelinesRequest) (response *DescribeLogstashPipelinesResponse, err error) {
    if request == nil {
        request = NewDescribeLogstashPipelinesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogstashPipelines require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogstashPipelinesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServerlessInstancesRequest() (request *DescribeServerlessInstancesRequest) {
    request = &DescribeServerlessInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "DescribeServerlessInstances")
    
    
    return
}

func NewDescribeServerlessInstancesResponse() (response *DescribeServerlessInstancesResponse) {
    response = &DescribeServerlessInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeServerlessInstances
// Serverless获取索引列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeServerlessInstances(request *DescribeServerlessInstancesRequest) (response *DescribeServerlessInstancesResponse, err error) {
    return c.DescribeServerlessInstancesWithContext(context.Background(), request)
}

// DescribeServerlessInstances
// Serverless获取索引列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeServerlessInstancesWithContext(ctx context.Context, request *DescribeServerlessInstancesRequest) (response *DescribeServerlessInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeServerlessInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServerlessInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServerlessInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServerlessMetricsRequest() (request *DescribeServerlessMetricsRequest) {
    request = &DescribeServerlessMetricsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "DescribeServerlessMetrics")
    
    
    return
}

func NewDescribeServerlessMetricsResponse() (response *DescribeServerlessMetricsResponse) {
    response = &DescribeServerlessMetricsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeServerlessMetrics
// 获取serverless实例对应指标，获取space维度时不需要传入indexid，获取index时不需要传入spaceid
//
// 获取一段时间时间范围内的指标数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeServerlessMetrics(request *DescribeServerlessMetricsRequest) (response *DescribeServerlessMetricsResponse, err error) {
    return c.DescribeServerlessMetricsWithContext(context.Background(), request)
}

// DescribeServerlessMetrics
// 获取serverless实例对应指标，获取space维度时不需要传入indexid，获取index时不需要传入spaceid
//
// 获取一段时间时间范围内的指标数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeServerlessMetricsWithContext(ctx context.Context, request *DescribeServerlessMetricsRequest) (response *DescribeServerlessMetricsResponse, err error) {
    if request == nil {
        request = NewDescribeServerlessMetricsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServerlessMetrics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServerlessMetricsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServerlessSpaceUserRequest() (request *DescribeServerlessSpaceUserRequest) {
    request = &DescribeServerlessSpaceUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "DescribeServerlessSpaceUser")
    
    
    return
}

func NewDescribeServerlessSpaceUserResponse() (response *DescribeServerlessSpaceUserResponse) {
    response = &DescribeServerlessSpaceUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeServerlessSpaceUser
// 查看Serverless空间子用户
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION_CLUSTERRESOURCELIMITERROR = "FailedOperation.ClusterResourceLimitError"
//  FAILEDOPERATION_DISKCOUNTPARAMERROR = "FailedOperation.DiskCountParamError"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  FAILEDOPERATION_NOTAUTHENTICATED = "FailedOperation.NotAuthenticated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_USERNAMEEXIST = "InvalidParameter.UsernameExist"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_INDEXCOUNT = "LimitExceeded.IndexCount"
//  LIMITEXCEEDED_INDEXCREATE = "LimitExceeded.IndexCreate"
//  LIMITEXCEEDED_SPACECOUNT = "LimitExceeded.SpaceCount"
//  LIMITEXCEEDED_SPACECREATE = "LimitExceeded.SpaceCreate"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_HIDDENZONE = "ResourceInsufficient.HiddenZone"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
func (c *Client) DescribeServerlessSpaceUser(request *DescribeServerlessSpaceUserRequest) (response *DescribeServerlessSpaceUserResponse, err error) {
    return c.DescribeServerlessSpaceUserWithContext(context.Background(), request)
}

// DescribeServerlessSpaceUser
// 查看Serverless空间子用户
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION_CLUSTERRESOURCELIMITERROR = "FailedOperation.ClusterResourceLimitError"
//  FAILEDOPERATION_DISKCOUNTPARAMERROR = "FailedOperation.DiskCountParamError"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  FAILEDOPERATION_NOTAUTHENTICATED = "FailedOperation.NotAuthenticated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_USERNAMEEXIST = "InvalidParameter.UsernameExist"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_INDEXCOUNT = "LimitExceeded.IndexCount"
//  LIMITEXCEEDED_INDEXCREATE = "LimitExceeded.IndexCreate"
//  LIMITEXCEEDED_SPACECOUNT = "LimitExceeded.SpaceCount"
//  LIMITEXCEEDED_SPACECREATE = "LimitExceeded.SpaceCreate"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_HIDDENZONE = "ResourceInsufficient.HiddenZone"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
func (c *Client) DescribeServerlessSpaceUserWithContext(ctx context.Context, request *DescribeServerlessSpaceUserRequest) (response *DescribeServerlessSpaceUserResponse, err error) {
    if request == nil {
        request = NewDescribeServerlessSpaceUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServerlessSpaceUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServerlessSpaceUserResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServerlessSpacesRequest() (request *DescribeServerlessSpacesRequest) {
    request = &DescribeServerlessSpacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "DescribeServerlessSpaces")
    
    
    return
}

func NewDescribeServerlessSpacesResponse() (response *DescribeServerlessSpacesResponse) {
    response = &DescribeServerlessSpacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeServerlessSpaces
// 获取Serverless索引空间列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeServerlessSpaces(request *DescribeServerlessSpacesRequest) (response *DescribeServerlessSpacesResponse, err error) {
    return c.DescribeServerlessSpacesWithContext(context.Background(), request)
}

// DescribeServerlessSpaces
// 获取Serverless索引空间列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeServerlessSpacesWithContext(ctx context.Context, request *DescribeServerlessSpacesRequest) (response *DescribeServerlessSpacesResponse, err error) {
    if request == nil {
        request = NewDescribeServerlessSpacesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServerlessSpaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServerlessSpacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSpaceKibanaToolsRequest() (request *DescribeSpaceKibanaToolsRequest) {
    request = &DescribeSpaceKibanaToolsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "DescribeSpaceKibanaTools")
    
    
    return
}

func NewDescribeSpaceKibanaToolsResponse() (response *DescribeSpaceKibanaToolsResponse) {
    response = &DescribeSpaceKibanaToolsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSpaceKibanaTools
// space维度的kibana获取登录token
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSpaceKibanaTools(request *DescribeSpaceKibanaToolsRequest) (response *DescribeSpaceKibanaToolsResponse, err error) {
    return c.DescribeSpaceKibanaToolsWithContext(context.Background(), request)
}

// DescribeSpaceKibanaTools
// space维度的kibana获取登录token
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSpaceKibanaToolsWithContext(ctx context.Context, request *DescribeSpaceKibanaToolsRequest) (response *DescribeSpaceKibanaToolsResponse, err error) {
    if request == nil {
        request = NewDescribeSpaceKibanaToolsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSpaceKibanaTools require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSpaceKibanaToolsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserCosSnapshotListRequest() (request *DescribeUserCosSnapshotListRequest) {
    request = &DescribeUserCosSnapshotListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "DescribeUserCosSnapshotList")
    
    
    return
}

func NewDescribeUserCosSnapshotListResponse() (response *DescribeUserCosSnapshotListResponse) {
    response = &DescribeUserCosSnapshotListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserCosSnapshotList
// 查询快照信息接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUserCosSnapshotList(request *DescribeUserCosSnapshotListRequest) (response *DescribeUserCosSnapshotListResponse, err error) {
    return c.DescribeUserCosSnapshotListWithContext(context.Background(), request)
}

// DescribeUserCosSnapshotList
// 查询快照信息接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUserCosSnapshotListWithContext(ctx context.Context, request *DescribeUserCosSnapshotListRequest) (response *DescribeUserCosSnapshotListResponse, err error) {
    if request == nil {
        request = NewDescribeUserCosSnapshotListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserCosSnapshotList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserCosSnapshotListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeViewsRequest() (request *DescribeViewsRequest) {
    request = &DescribeViewsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "DescribeViews")
    
    
    return
}

func NewDescribeViewsResponse() (response *DescribeViewsResponse) {
    response = &DescribeViewsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeViews
// 查询集群各视图数据，包括集群维度、节点维度、Kibana维度
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHCREATEINSTANCE = "AuthFailure.UnAuthCreateInstance"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_CAMINFONOTFOUND = "ResourceNotFound.CAMInfoNotFound"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
//  UNSUPPORTEDOPERATION_INSTANCETYPEERROR = "UnsupportedOperation.InstanceTypeError"
//  UNSUPPORTEDOPERATION_STATUSNOTSUPPORT = "UnsupportedOperation.StatusNotSupport"
func (c *Client) DescribeViews(request *DescribeViewsRequest) (response *DescribeViewsResponse, err error) {
    return c.DescribeViewsWithContext(context.Background(), request)
}

// DescribeViews
// 查询集群各视图数据，包括集群维度、节点维度、Kibana维度
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHCREATEINSTANCE = "AuthFailure.UnAuthCreateInstance"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_CAMINFONOTFOUND = "ResourceNotFound.CAMInfoNotFound"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
//  UNSUPPORTEDOPERATION_INSTANCETYPEERROR = "UnsupportedOperation.InstanceTypeError"
//  UNSUPPORTEDOPERATION_STATUSNOTSUPPORT = "UnsupportedOperation.StatusNotSupport"
func (c *Client) DescribeViewsWithContext(ctx context.Context, request *DescribeViewsRequest) (response *DescribeViewsResponse, err error) {
    if request == nil {
        request = NewDescribeViewsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeViews require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeViewsResponse()
    err = c.Send(request, response)
    return
}

func NewDiagnoseInstanceRequest() (request *DiagnoseInstanceRequest) {
    request = &DiagnoseInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "DiagnoseInstance")
    
    
    return
}

func NewDiagnoseInstanceResponse() (response *DiagnoseInstanceResponse) {
    response = &DiagnoseInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DiagnoseInstance
// 智能运维诊断集群
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  INVALIDPARAMETER_INVALIDREGION = "InvalidParameter.InvalidRegion"
//  LIMITEXCEEDED_DIAGNOSECOUNT = "LimitExceeded.DiagnoseCount"
//  RESOURCEINUSE_DIAGNOSE = "ResourceInUse.Diagnose"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
//  RESOURCENOTFOUND_DIAGNOSENOTFOUND = "ResourceNotFound.DiagnoseNotFound"
//  UNSUPPORTEDOPERATION_DIAGNOSEJOB = "UnsupportedOperation.DiagnoseJob"
//  UNSUPPORTEDOPERATION_DIAGNOSENOTOPEN = "UnsupportedOperation.DiagnoseNotOpen"
func (c *Client) DiagnoseInstance(request *DiagnoseInstanceRequest) (response *DiagnoseInstanceResponse, err error) {
    return c.DiagnoseInstanceWithContext(context.Background(), request)
}

// DiagnoseInstance
// 智能运维诊断集群
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  INVALIDPARAMETER_INVALIDREGION = "InvalidParameter.InvalidRegion"
//  LIMITEXCEEDED_DIAGNOSECOUNT = "LimitExceeded.DiagnoseCount"
//  RESOURCEINUSE_DIAGNOSE = "ResourceInUse.Diagnose"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
//  RESOURCENOTFOUND_DIAGNOSENOTFOUND = "ResourceNotFound.DiagnoseNotFound"
//  UNSUPPORTEDOPERATION_DIAGNOSEJOB = "UnsupportedOperation.DiagnoseJob"
//  UNSUPPORTEDOPERATION_DIAGNOSENOTOPEN = "UnsupportedOperation.DiagnoseNotOpen"
func (c *Client) DiagnoseInstanceWithContext(ctx context.Context, request *DiagnoseInstanceRequest) (response *DiagnoseInstanceResponse, err error) {
    if request == nil {
        request = NewDiagnoseInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DiagnoseInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDiagnoseInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewGetDiagnoseSettingsRequest() (request *GetDiagnoseSettingsRequest) {
    request = &GetDiagnoseSettingsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "GetDiagnoseSettings")
    
    
    return
}

func NewGetDiagnoseSettingsResponse() (response *GetDiagnoseSettingsResponse) {
    response = &GetDiagnoseSettingsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetDiagnoseSettings
// 查看智能运维配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  INVALIDPARAMETER_INVALIDREGION = "InvalidParameter.InvalidRegion"
//  INVALIDPARAMETERVALUE_CONFIGINFO = "InvalidParameterValue.ConfigInfo"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
func (c *Client) GetDiagnoseSettings(request *GetDiagnoseSettingsRequest) (response *GetDiagnoseSettingsResponse, err error) {
    return c.GetDiagnoseSettingsWithContext(context.Background(), request)
}

// GetDiagnoseSettings
// 查看智能运维配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  INVALIDPARAMETER_INVALIDREGION = "InvalidParameter.InvalidRegion"
//  INVALIDPARAMETERVALUE_CONFIGINFO = "InvalidParameterValue.ConfigInfo"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
func (c *Client) GetDiagnoseSettingsWithContext(ctx context.Context, request *GetDiagnoseSettingsRequest) (response *GetDiagnoseSettingsResponse, err error) {
    if request == nil {
        request = NewGetDiagnoseSettingsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDiagnoseSettings require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDiagnoseSettingsResponse()
    err = c.Send(request, response)
    return
}

func NewGetRequestTargetNodeTypesRequest() (request *GetRequestTargetNodeTypesRequest) {
    request = &GetRequestTargetNodeTypesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "GetRequestTargetNodeTypes")
    
    
    return
}

func NewGetRequestTargetNodeTypesResponse() (response *GetRequestTargetNodeTypesResponse) {
    response = &GetRequestTargetNodeTypesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetRequestTargetNodeTypes
// 获取接收客户端请求的节点类型
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INSTANCETYPEERROR = "UnsupportedOperation.InstanceTypeError"
func (c *Client) GetRequestTargetNodeTypes(request *GetRequestTargetNodeTypesRequest) (response *GetRequestTargetNodeTypesResponse, err error) {
    return c.GetRequestTargetNodeTypesWithContext(context.Background(), request)
}

// GetRequestTargetNodeTypes
// 获取接收客户端请求的节点类型
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INSTANCETYPEERROR = "UnsupportedOperation.InstanceTypeError"
func (c *Client) GetRequestTargetNodeTypesWithContext(ctx context.Context, request *GetRequestTargetNodeTypesRequest) (response *GetRequestTargetNodeTypesResponse, err error) {
    if request == nil {
        request = NewGetRequestTargetNodeTypesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetRequestTargetNodeTypes require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetRequestTargetNodeTypesResponse()
    err = c.Send(request, response)
    return
}

func NewInquirePriceRenewInstanceRequest() (request *InquirePriceRenewInstanceRequest) {
    request = &InquirePriceRenewInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "InquirePriceRenewInstance")
    
    
    return
}

func NewInquirePriceRenewInstanceResponse() (response *InquirePriceRenewInstanceResponse) {
    response = &InquirePriceRenewInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquirePriceRenewInstance
// 集群续费询价接口，续费前通过调用该接口，可获取集群续费的价格。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
//  RESOURCENOTFOUND_TRADECGWNOTFOUND = "ResourceNotFound.TradeCgwNotFound"
func (c *Client) InquirePriceRenewInstance(request *InquirePriceRenewInstanceRequest) (response *InquirePriceRenewInstanceResponse, err error) {
    return c.InquirePriceRenewInstanceWithContext(context.Background(), request)
}

// InquirePriceRenewInstance
// 集群续费询价接口，续费前通过调用该接口，可获取集群续费的价格。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
//  RESOURCENOTFOUND_TRADECGWNOTFOUND = "ResourceNotFound.TradeCgwNotFound"
func (c *Client) InquirePriceRenewInstanceWithContext(ctx context.Context, request *InquirePriceRenewInstanceRequest) (response *InquirePriceRenewInstanceResponse, err error) {
    if request == nil {
        request = NewInquirePriceRenewInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquirePriceRenewInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquirePriceRenewInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewInstallInstanceModelRequest() (request *InstallInstanceModelRequest) {
    request = &InstallInstanceModelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "InstallInstanceModel")
    
    
    return
}

func NewInstallInstanceModelResponse() (response *InstallInstanceModelResponse) {
    response = &InstallInstanceModelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InstallInstanceModel
// ES集群安装模型接口
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDESVERSION = "InvalidParameter.InvalidEsVersion"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
//  UNSUPPORTEDOPERATION_STATUSNOTNORMAL = "UnsupportedOperation.StatusNotNormal"
func (c *Client) InstallInstanceModel(request *InstallInstanceModelRequest) (response *InstallInstanceModelResponse, err error) {
    return c.InstallInstanceModelWithContext(context.Background(), request)
}

// InstallInstanceModel
// ES集群安装模型接口
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDESVERSION = "InvalidParameter.InvalidEsVersion"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
//  UNSUPPORTEDOPERATION_STATUSNOTNORMAL = "UnsupportedOperation.StatusNotNormal"
func (c *Client) InstallInstanceModelWithContext(ctx context.Context, request *InstallInstanceModelRequest) (response *InstallInstanceModelResponse, err error) {
    if request == nil {
        request = NewInstallInstanceModelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InstallInstanceModel require credential")
    }

    request.SetContext(ctx)
    
    response = NewInstallInstanceModelResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEsVipSecurityGroupRequest() (request *ModifyEsVipSecurityGroupRequest) {
    request = &ModifyEsVipSecurityGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "ModifyEsVipSecurityGroup")
    
    
    return
}

func NewModifyEsVipSecurityGroupResponse() (response *ModifyEsVipSecurityGroupResponse) {
    response = &ModifyEsVipSecurityGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyEsVipSecurityGroup
// 修改绑定VIP的安全组，传安全组id列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  INVALIDPARAMETER_INVALIDSECURITYGROUPIDS = "InvalidParameter.InvalidSecurityGroupIds"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
//  RESOURCENOTFOUND_SECURITYGROUPNOTFOUND = "ResourceNotFound.SecurityGroupNotFound"
//  UNSUPPORTEDOPERATION_STATUSNOTSUPPORT = "UnsupportedOperation.StatusNotSupport"
func (c *Client) ModifyEsVipSecurityGroup(request *ModifyEsVipSecurityGroupRequest) (response *ModifyEsVipSecurityGroupResponse, err error) {
    return c.ModifyEsVipSecurityGroupWithContext(context.Background(), request)
}

// ModifyEsVipSecurityGroup
// 修改绑定VIP的安全组，传安全组id列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  INVALIDPARAMETER_INVALIDSECURITYGROUPIDS = "InvalidParameter.InvalidSecurityGroupIds"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
//  RESOURCENOTFOUND_SECURITYGROUPNOTFOUND = "ResourceNotFound.SecurityGroupNotFound"
//  UNSUPPORTEDOPERATION_STATUSNOTSUPPORT = "UnsupportedOperation.StatusNotSupport"
func (c *Client) ModifyEsVipSecurityGroupWithContext(ctx context.Context, request *ModifyEsVipSecurityGroupRequest) (response *ModifyEsVipSecurityGroupResponse, err error) {
    if request == nil {
        request = NewModifyEsVipSecurityGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEsVipSecurityGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyEsVipSecurityGroupResponse()
    err = c.Send(request, response)
    return
}

func NewRestartInstanceRequest() (request *RestartInstanceRequest) {
    request = &RestartInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "RestartInstance")
    
    
    return
}

func NewRestartInstanceResponse() (response *RestartInstanceResponse) {
    response = &RestartInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RestartInstance
// 重启ES集群实例(用于系统版本更新等操作)
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDESVERSION = "InvalidParameter.InvalidEsVersion"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  INVALIDPARAMETERVALUE_CONFIGINFO = "InvalidParameterValue.ConfigInfo"
//  RESOURCENOTFOUND_CLUSTERINFONOTFOUND = "ResourceNotFound.ClusterInfoNotFound"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
//  UNSUPPORTEDOPERATION_CLUSTERSTATENOREPLICATION = "UnsupportedOperation.ClusterStateNoReplication"
//  UNSUPPORTEDOPERATION_CLUSTERSTATEUNHEALTH = "UnsupportedOperation.ClusterStateUnHealth"
//  UNSUPPORTEDOPERATION_INSTANCETYPEERROR = "UnsupportedOperation.InstanceTypeError"
//  UNSUPPORTEDOPERATION_STATUSNOTNORMAL = "UnsupportedOperation.StatusNotNormal"
//  UNSUPPORTEDOPERATION_STATUSNOTSUPPORT = "UnsupportedOperation.StatusNotSupport"
func (c *Client) RestartInstance(request *RestartInstanceRequest) (response *RestartInstanceResponse, err error) {
    return c.RestartInstanceWithContext(context.Background(), request)
}

// RestartInstance
// 重启ES集群实例(用于系统版本更新等操作)
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDESVERSION = "InvalidParameter.InvalidEsVersion"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  INVALIDPARAMETERVALUE_CONFIGINFO = "InvalidParameterValue.ConfigInfo"
//  RESOURCENOTFOUND_CLUSTERINFONOTFOUND = "ResourceNotFound.ClusterInfoNotFound"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
//  UNSUPPORTEDOPERATION_CLUSTERSTATENOREPLICATION = "UnsupportedOperation.ClusterStateNoReplication"
//  UNSUPPORTEDOPERATION_CLUSTERSTATEUNHEALTH = "UnsupportedOperation.ClusterStateUnHealth"
//  UNSUPPORTEDOPERATION_INSTANCETYPEERROR = "UnsupportedOperation.InstanceTypeError"
//  UNSUPPORTEDOPERATION_STATUSNOTNORMAL = "UnsupportedOperation.StatusNotNormal"
//  UNSUPPORTEDOPERATION_STATUSNOTSUPPORT = "UnsupportedOperation.StatusNotSupport"
func (c *Client) RestartInstanceWithContext(ctx context.Context, request *RestartInstanceRequest) (response *RestartInstanceResponse, err error) {
    if request == nil {
        request = NewRestartInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestartInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewRestartInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewRestartKibanaRequest() (request *RestartKibanaRequest) {
    request = &RestartKibanaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "RestartKibana")
    
    
    return
}

func NewRestartKibanaResponse() (response *RestartKibanaResponse) {
    response = &RestartKibanaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RestartKibana
// 重启Kibana
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
//  UNSUPPORTEDOPERATION_INSTANCETYPEERROR = "UnsupportedOperation.InstanceTypeError"
//  UNSUPPORTEDOPERATION_STATUSNOTNORMAL = "UnsupportedOperation.StatusNotNormal"
//  UNSUPPORTEDOPERATION_STATUSNOTSUPPORT = "UnsupportedOperation.StatusNotSupport"
func (c *Client) RestartKibana(request *RestartKibanaRequest) (response *RestartKibanaResponse, err error) {
    return c.RestartKibanaWithContext(context.Background(), request)
}

// RestartKibana
// 重启Kibana
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
//  UNSUPPORTEDOPERATION_INSTANCETYPEERROR = "UnsupportedOperation.InstanceTypeError"
//  UNSUPPORTEDOPERATION_STATUSNOTNORMAL = "UnsupportedOperation.StatusNotNormal"
//  UNSUPPORTEDOPERATION_STATUSNOTSUPPORT = "UnsupportedOperation.StatusNotSupport"
func (c *Client) RestartKibanaWithContext(ctx context.Context, request *RestartKibanaRequest) (response *RestartKibanaResponse, err error) {
    if request == nil {
        request = NewRestartKibanaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestartKibana require credential")
    }

    request.SetContext(ctx)
    
    response = NewRestartKibanaResponse()
    err = c.Send(request, response)
    return
}

func NewRestartLogstashInstanceRequest() (request *RestartLogstashInstanceRequest) {
    request = &RestartLogstashInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "RestartLogstashInstance")
    
    
    return
}

func NewRestartLogstashInstanceResponse() (response *RestartLogstashInstanceResponse) {
    response = &RestartLogstashInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RestartLogstashInstance
// 用于重启Logstash实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RestartLogstashInstance(request *RestartLogstashInstanceRequest) (response *RestartLogstashInstanceResponse, err error) {
    return c.RestartLogstashInstanceWithContext(context.Background(), request)
}

// RestartLogstashInstance
// 用于重启Logstash实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RestartLogstashInstanceWithContext(ctx context.Context, request *RestartLogstashInstanceRequest) (response *RestartLogstashInstanceResponse, err error) {
    if request == nil {
        request = NewRestartLogstashInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestartLogstashInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewRestartLogstashInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewRestartNodesRequest() (request *RestartNodesRequest) {
    request = &RestartNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "RestartNodes")
    
    
    return
}

func NewRestartNodesResponse() (response *RestartNodesResponse) {
    response = &RestartNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RestartNodes
// 用于重启集群节点
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  INVALIDPARAMETER_INVALIDIPLIST = "InvalidParameter.InvalidIpList"
//  INVALIDPARAMETER_INVALIDNODENAMES = "InvalidParameter.InvalidNodeNames"
//  INVALIDPARAMETER_INVALIDRESTARTMODE = "InvalidParameter.InvalidRestartMode"
//  INVALIDPARAMETER_INVALIDSUBNETUIDLIST = "InvalidParameter.InvalidSubnetUidList"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT_SUBNETIP = "ResourceInsufficient.SubnetIp"
//  RESOURCENOTFOUND_CLUSTERINFONOTFOUND = "ResourceNotFound.ClusterInfoNotFound"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
//  UNSUPPORTEDOPERATION_CLUSTERSTATENOREPLICATION = "UnsupportedOperation.ClusterStateNoReplication"
//  UNSUPPORTEDOPERATION_CLUSTERSTATEUNHEALTH = "UnsupportedOperation.ClusterStateUnHealth"
//  UNSUPPORTEDOPERATION_INSTANCETYPEERROR = "UnsupportedOperation.InstanceTypeError"
//  UNSUPPORTEDOPERATION_RESTARTMODE = "UnsupportedOperation.RestartMode"
//  UNSUPPORTEDOPERATION_STATUSNOTNORMAL = "UnsupportedOperation.StatusNotNormal"
//  UNSUPPORTEDOPERATION_STATUSNOTSUPPORT = "UnsupportedOperation.StatusNotSupport"
func (c *Client) RestartNodes(request *RestartNodesRequest) (response *RestartNodesResponse, err error) {
    return c.RestartNodesWithContext(context.Background(), request)
}

// RestartNodes
// 用于重启集群节点
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  INVALIDPARAMETER_INVALIDIPLIST = "InvalidParameter.InvalidIpList"
//  INVALIDPARAMETER_INVALIDNODENAMES = "InvalidParameter.InvalidNodeNames"
//  INVALIDPARAMETER_INVALIDRESTARTMODE = "InvalidParameter.InvalidRestartMode"
//  INVALIDPARAMETER_INVALIDSUBNETUIDLIST = "InvalidParameter.InvalidSubnetUidList"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT_SUBNETIP = "ResourceInsufficient.SubnetIp"
//  RESOURCENOTFOUND_CLUSTERINFONOTFOUND = "ResourceNotFound.ClusterInfoNotFound"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
//  UNSUPPORTEDOPERATION_CLUSTERSTATENOREPLICATION = "UnsupportedOperation.ClusterStateNoReplication"
//  UNSUPPORTEDOPERATION_CLUSTERSTATEUNHEALTH = "UnsupportedOperation.ClusterStateUnHealth"
//  UNSUPPORTEDOPERATION_INSTANCETYPEERROR = "UnsupportedOperation.InstanceTypeError"
//  UNSUPPORTEDOPERATION_RESTARTMODE = "UnsupportedOperation.RestartMode"
//  UNSUPPORTEDOPERATION_STATUSNOTNORMAL = "UnsupportedOperation.StatusNotNormal"
//  UNSUPPORTEDOPERATION_STATUSNOTSUPPORT = "UnsupportedOperation.StatusNotSupport"
func (c *Client) RestartNodesWithContext(ctx context.Context, request *RestartNodesRequest) (response *RestartNodesResponse, err error) {
    if request == nil {
        request = NewRestartNodesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestartNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewRestartNodesResponse()
    err = c.Send(request, response)
    return
}

func NewRestoreClusterSnapshotRequest() (request *RestoreClusterSnapshotRequest) {
    request = &RestoreClusterSnapshotRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "RestoreClusterSnapshot")
    
    
    return
}

func NewRestoreClusterSnapshotResponse() (response *RestoreClusterSnapshotResponse) {
    response = &RestoreClusterSnapshotResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RestoreClusterSnapshot
// 快照备份恢复，从仓库中恢复快照到集群中
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDESVERSION = "InvalidParameter.InvalidEsVersion"
//  INVALIDPARAMETERVALUE_PASSWORD = "InvalidParameterValue.Password"
//  RESOURCENOTFOUND_COSINFONOTFOUND = "ResourceNotFound.CosInfoNotFound"
func (c *Client) RestoreClusterSnapshot(request *RestoreClusterSnapshotRequest) (response *RestoreClusterSnapshotResponse, err error) {
    return c.RestoreClusterSnapshotWithContext(context.Background(), request)
}

// RestoreClusterSnapshot
// 快照备份恢复，从仓库中恢复快照到集群中
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDESVERSION = "InvalidParameter.InvalidEsVersion"
//  INVALIDPARAMETERVALUE_PASSWORD = "InvalidParameterValue.Password"
//  RESOURCENOTFOUND_COSINFONOTFOUND = "ResourceNotFound.CosInfoNotFound"
func (c *Client) RestoreClusterSnapshotWithContext(ctx context.Context, request *RestoreClusterSnapshotRequest) (response *RestoreClusterSnapshotResponse, err error) {
    if request == nil {
        request = NewRestoreClusterSnapshotRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestoreClusterSnapshot require credential")
    }

    request.SetContext(ctx)
    
    response = NewRestoreClusterSnapshotResponse()
    err = c.Send(request, response)
    return
}

func NewSaveAndDeployLogstashPipelineRequest() (request *SaveAndDeployLogstashPipelineRequest) {
    request = &SaveAndDeployLogstashPipelineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "SaveAndDeployLogstashPipeline")
    
    
    return
}

func NewSaveAndDeployLogstashPipelineResponse() (response *SaveAndDeployLogstashPipelineResponse) {
    response = &SaveAndDeployLogstashPipelineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SaveAndDeployLogstashPipeline
// 用于下发并且部署管道
//
// 可能返回的错误码:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) SaveAndDeployLogstashPipeline(request *SaveAndDeployLogstashPipelineRequest) (response *SaveAndDeployLogstashPipelineResponse, err error) {
    return c.SaveAndDeployLogstashPipelineWithContext(context.Background(), request)
}

// SaveAndDeployLogstashPipeline
// 用于下发并且部署管道
//
// 可能返回的错误码:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) SaveAndDeployLogstashPipelineWithContext(ctx context.Context, request *SaveAndDeployLogstashPipelineRequest) (response *SaveAndDeployLogstashPipelineResponse, err error) {
    if request == nil {
        request = NewSaveAndDeployLogstashPipelineRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SaveAndDeployLogstashPipeline require credential")
    }

    request.SetContext(ctx)
    
    response = NewSaveAndDeployLogstashPipelineResponse()
    err = c.Send(request, response)
    return
}

func NewStartLogstashPipelinesRequest() (request *StartLogstashPipelinesRequest) {
    request = &StartLogstashPipelinesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "StartLogstashPipelines")
    
    
    return
}

func NewStartLogstashPipelinesResponse() (response *StartLogstashPipelinesResponse) {
    response = &StartLogstashPipelinesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartLogstashPipelines
// 用于启动Logstash管道
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) StartLogstashPipelines(request *StartLogstashPipelinesRequest) (response *StartLogstashPipelinesResponse, err error) {
    return c.StartLogstashPipelinesWithContext(context.Background(), request)
}

// StartLogstashPipelines
// 用于启动Logstash管道
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) StartLogstashPipelinesWithContext(ctx context.Context, request *StartLogstashPipelinesRequest) (response *StartLogstashPipelinesResponse, err error) {
    if request == nil {
        request = NewStartLogstashPipelinesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartLogstashPipelines require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartLogstashPipelinesResponse()
    err = c.Send(request, response)
    return
}

func NewStopLogstashPipelinesRequest() (request *StopLogstashPipelinesRequest) {
    request = &StopLogstashPipelinesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "StopLogstashPipelines")
    
    
    return
}

func NewStopLogstashPipelinesResponse() (response *StopLogstashPipelinesResponse) {
    response = &StopLogstashPipelinesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopLogstashPipelines
// 用于批量停止Logstash管道
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) StopLogstashPipelines(request *StopLogstashPipelinesRequest) (response *StopLogstashPipelinesResponse, err error) {
    return c.StopLogstashPipelinesWithContext(context.Background(), request)
}

// StopLogstashPipelines
// 用于批量停止Logstash管道
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) StopLogstashPipelinesWithContext(ctx context.Context, request *StopLogstashPipelinesRequest) (response *StopLogstashPipelinesResponse, err error) {
    if request == nil {
        request = NewStopLogstashPipelinesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopLogstashPipelines require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopLogstashPipelinesResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateDiagnoseSettingsRequest() (request *UpdateDiagnoseSettingsRequest) {
    request = &UpdateDiagnoseSettingsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "UpdateDiagnoseSettings")
    
    
    return
}

func NewUpdateDiagnoseSettingsResponse() (response *UpdateDiagnoseSettingsResponse) {
    response = &UpdateDiagnoseSettingsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateDiagnoseSettings
// 更新智能运维配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  INVALIDPARAMETER_INVALIDREGION = "InvalidParameter.InvalidRegion"
//  INVALIDPARAMETER_INVALIDTIMEPARAM = "InvalidParameter.InvalidTimeParam"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
//  UNSUPPORTEDOPERATION_STATUSNOTSUPPORT = "UnsupportedOperation.StatusNotSupport"
func (c *Client) UpdateDiagnoseSettings(request *UpdateDiagnoseSettingsRequest) (response *UpdateDiagnoseSettingsResponse, err error) {
    return c.UpdateDiagnoseSettingsWithContext(context.Background(), request)
}

// UpdateDiagnoseSettings
// 更新智能运维配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  INVALIDPARAMETER_INVALIDREGION = "InvalidParameter.InvalidRegion"
//  INVALIDPARAMETER_INVALIDTIMEPARAM = "InvalidParameter.InvalidTimeParam"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
//  UNSUPPORTEDOPERATION_STATUSNOTSUPPORT = "UnsupportedOperation.StatusNotSupport"
func (c *Client) UpdateDiagnoseSettingsWithContext(ctx context.Context, request *UpdateDiagnoseSettingsRequest) (response *UpdateDiagnoseSettingsResponse, err error) {
    if request == nil {
        request = NewUpdateDiagnoseSettingsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateDiagnoseSettings require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateDiagnoseSettingsResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateDictionariesRequest() (request *UpdateDictionariesRequest) {
    request = &UpdateDictionariesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "UpdateDictionaries")
    
    
    return
}

func NewUpdateDictionariesResponse() (response *UpdateDictionariesResponse) {
    response = &UpdateDictionariesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateDictionaries
// 更新ES集群词典
//
// 可能返回的错误码:
//  FAILEDOPERATION_ESDICTIONARYINFOERROR = "FailedOperation.EsDictionaryInfoError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  INVALIDPARAMETER_INVALIDUPDATETYPE = "InvalidParameter.InvalidUpdateType"
//  LIMITEXCEEDED_UPDATEITEMLIMIT = "LimitExceeded.UpdateItemLimit"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_CLUSTERINFONOTFOUND = "ResourceNotFound.ClusterInfoNotFound"
//  RESOURCENOTFOUND_COSINFONOTFOUND = "ResourceNotFound.CosInfoNotFound"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
//  RESOURCENOTFOUND_WHITELISTNOTFOUND = "ResourceNotFound.WhiteListNotFound"
//  UNSUPPORTEDOPERATION_CLUSTERSTATENOREPLICATION = "UnsupportedOperation.ClusterStateNoReplication"
//  UNSUPPORTEDOPERATION_CLUSTERSTATEUNHEALTH = "UnsupportedOperation.ClusterStateUnHealth"
//  UNSUPPORTEDOPERATION_PLUGIN = "UnsupportedOperation.Plugin"
//  UNSUPPORTEDOPERATION_STATUSNOTNORMAL = "UnsupportedOperation.StatusNotNormal"
//  UNSUPPORTEDOPERATION_STATUSNOTSUPPORT = "UnsupportedOperation.StatusNotSupport"
func (c *Client) UpdateDictionaries(request *UpdateDictionariesRequest) (response *UpdateDictionariesResponse, err error) {
    return c.UpdateDictionariesWithContext(context.Background(), request)
}

// UpdateDictionaries
// 更新ES集群词典
//
// 可能返回的错误码:
//  FAILEDOPERATION_ESDICTIONARYINFOERROR = "FailedOperation.EsDictionaryInfoError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  INVALIDPARAMETER_INVALIDUPDATETYPE = "InvalidParameter.InvalidUpdateType"
//  LIMITEXCEEDED_UPDATEITEMLIMIT = "LimitExceeded.UpdateItemLimit"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_CLUSTERINFONOTFOUND = "ResourceNotFound.ClusterInfoNotFound"
//  RESOURCENOTFOUND_COSINFONOTFOUND = "ResourceNotFound.CosInfoNotFound"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
//  RESOURCENOTFOUND_WHITELISTNOTFOUND = "ResourceNotFound.WhiteListNotFound"
//  UNSUPPORTEDOPERATION_CLUSTERSTATENOREPLICATION = "UnsupportedOperation.ClusterStateNoReplication"
//  UNSUPPORTEDOPERATION_CLUSTERSTATEUNHEALTH = "UnsupportedOperation.ClusterStateUnHealth"
//  UNSUPPORTEDOPERATION_PLUGIN = "UnsupportedOperation.Plugin"
//  UNSUPPORTEDOPERATION_STATUSNOTNORMAL = "UnsupportedOperation.StatusNotNormal"
//  UNSUPPORTEDOPERATION_STATUSNOTSUPPORT = "UnsupportedOperation.StatusNotSupport"
func (c *Client) UpdateDictionariesWithContext(ctx context.Context, request *UpdateDictionariesRequest) (response *UpdateDictionariesResponse, err error) {
    if request == nil {
        request = NewUpdateDictionariesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateDictionaries require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateDictionariesResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateIndexRequest() (request *UpdateIndexRequest) {
    request = &UpdateIndexRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "UpdateIndex")
    
    
    return
}

func NewUpdateIndexResponse() (response *UpdateIndexResponse) {
    response = &UpdateIndexResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateIndex
// 更新索引
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDESVERSION = "InvalidParameter.InvalidEsVersion"
//  INVALIDPARAMETER_INVALIDINDEXMETAJSON = "InvalidParameter.InvalidIndexMetaJson"
//  INVALIDPARAMETER_INVALIDINDEXTYPE = "InvalidParameter.InvalidIndexType"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  INVALIDPARAMETER_INVALIDUPDATEMETAJSON = "InvalidParameter.InvalidUpdateMetaJson"
//  LIMITEXCEEDED_UPDATEITEMLIMIT = "LimitExceeded.UpdateItemLimit"
//  RESOURCENOTFOUND_CLUSTERINFONOTFOUND = "ResourceNotFound.ClusterInfoNotFound"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
func (c *Client) UpdateIndex(request *UpdateIndexRequest) (response *UpdateIndexResponse, err error) {
    return c.UpdateIndexWithContext(context.Background(), request)
}

// UpdateIndex
// 更新索引
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDESVERSION = "InvalidParameter.InvalidEsVersion"
//  INVALIDPARAMETER_INVALIDINDEXMETAJSON = "InvalidParameter.InvalidIndexMetaJson"
//  INVALIDPARAMETER_INVALIDINDEXTYPE = "InvalidParameter.InvalidIndexType"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  INVALIDPARAMETER_INVALIDUPDATEMETAJSON = "InvalidParameter.InvalidUpdateMetaJson"
//  LIMITEXCEEDED_UPDATEITEMLIMIT = "LimitExceeded.UpdateItemLimit"
//  RESOURCENOTFOUND_CLUSTERINFONOTFOUND = "ResourceNotFound.ClusterInfoNotFound"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
func (c *Client) UpdateIndexWithContext(ctx context.Context, request *UpdateIndexRequest) (response *UpdateIndexResponse, err error) {
    if request == nil {
        request = NewUpdateIndexRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateIndex require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateIndexResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateInstanceRequest() (request *UpdateInstanceRequest) {
    request = &UpdateInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "UpdateInstance")
    
    
    return
}

func NewUpdateInstanceResponse() (response *UpdateInstanceResponse) {
    response = &UpdateInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateInstance
// 对集群进行节点规格变更，修改实例名称，修改配置，重置密码， 添加Kibana黑白名单等操作。参数中InstanceId为必传参数，ForceRestart为选填参数，剩余参数传递组合及含义如下：
//
// - InstanceName：修改实例名称(仅用于标识实例)
//
// - NodeInfoList: 修改节点配置（节点横向扩缩容，纵向扩缩容，增加主节点，增加冷节点等）
//
// - EsConfig：修改集群配置
//
// - Password：修改默认用户elastic的密码
//
// - EsAcl：修改访问控制列表
//
// - CosBackUp: 设置集群COS自动备份信息
//
// 以上参数组合只能传递一种，多传或少传均会导致请求失败
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ESDICTIONARYINFOERROR = "FailedOperation.EsDictionaryInfoError"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  FAILEDOPERATION_NOTAUTHENTICATED = "FailedOperation.NotAuthenticated"
//  FAILEDOPERATION_UNSUPPORTEDLOCALDISKROLLUPSCALEUPORDOWN = "FailedOperation.UnsupportedLocalDiskRollUpScaleUpOrDown"
//  FAILEDOPERATION_UNSUPPORTEDRESETNODETYPEANDSCALEOUTDISK = "FailedOperation.UnsupportedResetNodeTypeAndScaleOutDisk"
//  FAILEDOPERATION_UNSUPPORTEDRESTSCALEDOWNANDMODIFYDISK = "FailedOperation.UnsupportedRestScaleDownAndModifyDisk"
//  FAILEDOPERATION_UNSUPPORTEDREVERSEREGULATIONNODETYPEANDDISK = "FailedOperation.UnsupportedReverseRegulationNodeTypeAndDisk"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDAUTOVOUCHER = "InvalidParameter.InvalidAutoVoucher"
//  INVALIDPARAMETER_INVALIDCLUSTERNAME = "InvalidParameter.InvalidClusterName"
//  INVALIDPARAMETER_INVALIDCOSBACKUPINFO = "InvalidParameter.InvalidCosBackupInfo"
//  INVALIDPARAMETER_INVALIDDISKCOUNT = "InvalidParameter.InvalidDiskCount"
//  INVALIDPARAMETER_INVALIDDISKENCRYPT = "InvalidParameter.InvalidDiskEncrypt"
//  INVALIDPARAMETER_INVALIDDISKSIZE = "InvalidParameter.InvalidDiskSize"
//  INVALIDPARAMETER_INVALIDDISKTYPE = "InvalidParameter.InvalidDiskType"
//  INVALIDPARAMETER_INVALIDESACL = "InvalidParameter.InvalidEsACL"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  INVALIDPARAMETER_INVALIDKIBANAPRIVATEPORT = "InvalidParameter.InvalidKibanaPrivatePort"
//  INVALIDPARAMETER_INVALIDMULTIZONEINFO = "InvalidParameter.InvalidMultiZoneInfo"
//  INVALIDPARAMETER_INVALIDNODENUM = "InvalidParameter.InvalidNodeNum"
//  INVALIDPARAMETER_INVALIDNODETYPE = "InvalidParameter.InvalidNodeType"
//  INVALIDPARAMETER_INVALIDOPERATIONDURATION = "InvalidParameter.InvalidOperationDuration"
//  INVALIDPARAMETER_INVALIDPRIVATEACCESS = "InvalidParameter.InvalidPrivateAccess"
//  INVALIDPARAMETER_INVALIDPUBLICACCESS = "InvalidParameter.InvalidPublicAccess"
//  INVALIDPARAMETER_INVALIDREGION = "InvalidParameter.InvalidRegion"
//  INVALIDPARAMETER_INVALIDRESTARTTYPE = "InvalidParameter.InvalidRestartType"
//  INVALIDPARAMETER_INVALIDSUBNETID = "InvalidParameter.InvalidSubnetId"
//  INVALIDPARAMETER_INVALIDTYPE = "InvalidParameter.InvalidType"
//  INVALIDPARAMETER_INVALIDUPDATETYPE = "InvalidParameter.InvalidUpdateType"
//  INVALIDPARAMETER_INVALIDVOUCHERIDS = "InvalidParameter.InvalidVoucherIds"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  INVALIDPARAMETERVALUE_CHARGETYPE = "InvalidParameterValue.ChargeType"
//  INVALIDPARAMETERVALUE_CONFIGINFO = "InvalidParameterValue.ConfigInfo"
//  INVALIDPARAMETERVALUE_ESCONFIGTYPE = "InvalidParameterValue.EsConfigType"
//  INVALIDPARAMETERVALUE_INSTANCENAME = "InvalidParameterValue.InstanceName"
//  INVALIDPARAMETERVALUE_INVALIDDEADLINE = "InvalidParameterValue.InvalidDeadline"
//  INVALIDPARAMETERVALUE_PASSWORD = "InvalidParameterValue.Password"
//  LIMITEXCEEDED_UPDATEITEMLIMIT = "LimitExceeded.UpdateItemLimit"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_ORDER = "ResourceInUse.Order"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_CVM = "ResourceInsufficient.CVM"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
//  RESOURCEINSUFFICIENT_SUBNETIP = "ResourceInsufficient.SubnetIp"
//  RESOURCENOTFOUND_CLUSTERINFONOTFOUND = "ResourceNotFound.ClusterInfoNotFound"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
//  RESOURCENOTFOUND_DISKINFONOTFOUND = "ResourceNotFound.DiskInfoNotFound"
//  RESOURCENOTFOUND_TRADECGWNOTFOUND = "ResourceNotFound.TradeCgwNotFound"
//  RESOURCENOTFOUND_VPCINFONOTFOUND = "ResourceNotFound.VPCInfoNotFound"
//  RESOURCENOTFOUND_WHITELISTNOTFOUND = "ResourceNotFound.WhiteListNotFound"
//  UNAUTHORIZEDOPERATION_UINNOTINWHITELIST = "UnauthorizedOperation.UinNotInWhiteList"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BASICSECURITYTYPE = "UnsupportedOperation.BasicSecurityType"
//  UNSUPPORTEDOPERATION_CHANGENODETYPE = "UnsupportedOperation.ChangeNodeType"
//  UNSUPPORTEDOPERATION_CLUSTERSHARDNUM = "UnsupportedOperation.ClusterShardNum"
//  UNSUPPORTEDOPERATION_CLUSTERSTATECLOSE = "UnsupportedOperation.ClusterStateClose"
//  UNSUPPORTEDOPERATION_CLUSTERSTATENOREPLICATION = "UnsupportedOperation.ClusterStateNoReplication"
//  UNSUPPORTEDOPERATION_CLUSTERSTATEUNHEALTH = "UnsupportedOperation.ClusterStateUnHealth"
//  UNSUPPORTEDOPERATION_INDEXSETTINGSREQUIRESET = "UnsupportedOperation.IndexSettingsRequireSet"
//  UNSUPPORTEDOPERATION_LICENSEERROR = "UnsupportedOperation.LicenseError"
//  UNSUPPORTEDOPERATION_MULTIZONESUPGRADE = "UnsupportedOperation.MultiZonesUpgrade"
//  UNSUPPORTEDOPERATION_PLUGIN = "UnsupportedOperation.Plugin"
//  UNSUPPORTEDOPERATION_SCALEDOWNTOOMUCH = "UnsupportedOperation.ScaleDownTooMuch"
//  UNSUPPORTEDOPERATION_STATUSNOTNORMAL = "UnsupportedOperation.StatusNotNormal"
//  UNSUPPORTEDOPERATION_STATUSNOTSUPPORT = "UnsupportedOperation.StatusNotSupport"
//  UNSUPPORTEDOPERATION_UPDATEDISKTYPE = "UnsupportedOperation.UpdateDiskType"
//  UNSUPPORTEDOPERATION_WEBSERVICETYPE = "UnsupportedOperation.WebServiceType"
func (c *Client) UpdateInstance(request *UpdateInstanceRequest) (response *UpdateInstanceResponse, err error) {
    return c.UpdateInstanceWithContext(context.Background(), request)
}

// UpdateInstance
// 对集群进行节点规格变更，修改实例名称，修改配置，重置密码， 添加Kibana黑白名单等操作。参数中InstanceId为必传参数，ForceRestart为选填参数，剩余参数传递组合及含义如下：
//
// - InstanceName：修改实例名称(仅用于标识实例)
//
// - NodeInfoList: 修改节点配置（节点横向扩缩容，纵向扩缩容，增加主节点，增加冷节点等）
//
// - EsConfig：修改集群配置
//
// - Password：修改默认用户elastic的密码
//
// - EsAcl：修改访问控制列表
//
// - CosBackUp: 设置集群COS自动备份信息
//
// 以上参数组合只能传递一种，多传或少传均会导致请求失败
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ESDICTIONARYINFOERROR = "FailedOperation.EsDictionaryInfoError"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  FAILEDOPERATION_NOTAUTHENTICATED = "FailedOperation.NotAuthenticated"
//  FAILEDOPERATION_UNSUPPORTEDLOCALDISKROLLUPSCALEUPORDOWN = "FailedOperation.UnsupportedLocalDiskRollUpScaleUpOrDown"
//  FAILEDOPERATION_UNSUPPORTEDRESETNODETYPEANDSCALEOUTDISK = "FailedOperation.UnsupportedResetNodeTypeAndScaleOutDisk"
//  FAILEDOPERATION_UNSUPPORTEDRESTSCALEDOWNANDMODIFYDISK = "FailedOperation.UnsupportedRestScaleDownAndModifyDisk"
//  FAILEDOPERATION_UNSUPPORTEDREVERSEREGULATIONNODETYPEANDDISK = "FailedOperation.UnsupportedReverseRegulationNodeTypeAndDisk"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDAUTOVOUCHER = "InvalidParameter.InvalidAutoVoucher"
//  INVALIDPARAMETER_INVALIDCLUSTERNAME = "InvalidParameter.InvalidClusterName"
//  INVALIDPARAMETER_INVALIDCOSBACKUPINFO = "InvalidParameter.InvalidCosBackupInfo"
//  INVALIDPARAMETER_INVALIDDISKCOUNT = "InvalidParameter.InvalidDiskCount"
//  INVALIDPARAMETER_INVALIDDISKENCRYPT = "InvalidParameter.InvalidDiskEncrypt"
//  INVALIDPARAMETER_INVALIDDISKSIZE = "InvalidParameter.InvalidDiskSize"
//  INVALIDPARAMETER_INVALIDDISKTYPE = "InvalidParameter.InvalidDiskType"
//  INVALIDPARAMETER_INVALIDESACL = "InvalidParameter.InvalidEsACL"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  INVALIDPARAMETER_INVALIDKIBANAPRIVATEPORT = "InvalidParameter.InvalidKibanaPrivatePort"
//  INVALIDPARAMETER_INVALIDMULTIZONEINFO = "InvalidParameter.InvalidMultiZoneInfo"
//  INVALIDPARAMETER_INVALIDNODENUM = "InvalidParameter.InvalidNodeNum"
//  INVALIDPARAMETER_INVALIDNODETYPE = "InvalidParameter.InvalidNodeType"
//  INVALIDPARAMETER_INVALIDOPERATIONDURATION = "InvalidParameter.InvalidOperationDuration"
//  INVALIDPARAMETER_INVALIDPRIVATEACCESS = "InvalidParameter.InvalidPrivateAccess"
//  INVALIDPARAMETER_INVALIDPUBLICACCESS = "InvalidParameter.InvalidPublicAccess"
//  INVALIDPARAMETER_INVALIDREGION = "InvalidParameter.InvalidRegion"
//  INVALIDPARAMETER_INVALIDRESTARTTYPE = "InvalidParameter.InvalidRestartType"
//  INVALIDPARAMETER_INVALIDSUBNETID = "InvalidParameter.InvalidSubnetId"
//  INVALIDPARAMETER_INVALIDTYPE = "InvalidParameter.InvalidType"
//  INVALIDPARAMETER_INVALIDUPDATETYPE = "InvalidParameter.InvalidUpdateType"
//  INVALIDPARAMETER_INVALIDVOUCHERIDS = "InvalidParameter.InvalidVoucherIds"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  INVALIDPARAMETERVALUE_CHARGETYPE = "InvalidParameterValue.ChargeType"
//  INVALIDPARAMETERVALUE_CONFIGINFO = "InvalidParameterValue.ConfigInfo"
//  INVALIDPARAMETERVALUE_ESCONFIGTYPE = "InvalidParameterValue.EsConfigType"
//  INVALIDPARAMETERVALUE_INSTANCENAME = "InvalidParameterValue.InstanceName"
//  INVALIDPARAMETERVALUE_INVALIDDEADLINE = "InvalidParameterValue.InvalidDeadline"
//  INVALIDPARAMETERVALUE_PASSWORD = "InvalidParameterValue.Password"
//  LIMITEXCEEDED_UPDATEITEMLIMIT = "LimitExceeded.UpdateItemLimit"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_ORDER = "ResourceInUse.Order"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_CVM = "ResourceInsufficient.CVM"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
//  RESOURCEINSUFFICIENT_SUBNETIP = "ResourceInsufficient.SubnetIp"
//  RESOURCENOTFOUND_CLUSTERINFONOTFOUND = "ResourceNotFound.ClusterInfoNotFound"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
//  RESOURCENOTFOUND_DISKINFONOTFOUND = "ResourceNotFound.DiskInfoNotFound"
//  RESOURCENOTFOUND_TRADECGWNOTFOUND = "ResourceNotFound.TradeCgwNotFound"
//  RESOURCENOTFOUND_VPCINFONOTFOUND = "ResourceNotFound.VPCInfoNotFound"
//  RESOURCENOTFOUND_WHITELISTNOTFOUND = "ResourceNotFound.WhiteListNotFound"
//  UNAUTHORIZEDOPERATION_UINNOTINWHITELIST = "UnauthorizedOperation.UinNotInWhiteList"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BASICSECURITYTYPE = "UnsupportedOperation.BasicSecurityType"
//  UNSUPPORTEDOPERATION_CHANGENODETYPE = "UnsupportedOperation.ChangeNodeType"
//  UNSUPPORTEDOPERATION_CLUSTERSHARDNUM = "UnsupportedOperation.ClusterShardNum"
//  UNSUPPORTEDOPERATION_CLUSTERSTATECLOSE = "UnsupportedOperation.ClusterStateClose"
//  UNSUPPORTEDOPERATION_CLUSTERSTATENOREPLICATION = "UnsupportedOperation.ClusterStateNoReplication"
//  UNSUPPORTEDOPERATION_CLUSTERSTATEUNHEALTH = "UnsupportedOperation.ClusterStateUnHealth"
//  UNSUPPORTEDOPERATION_INDEXSETTINGSREQUIRESET = "UnsupportedOperation.IndexSettingsRequireSet"
//  UNSUPPORTEDOPERATION_LICENSEERROR = "UnsupportedOperation.LicenseError"
//  UNSUPPORTEDOPERATION_MULTIZONESUPGRADE = "UnsupportedOperation.MultiZonesUpgrade"
//  UNSUPPORTEDOPERATION_PLUGIN = "UnsupportedOperation.Plugin"
//  UNSUPPORTEDOPERATION_SCALEDOWNTOOMUCH = "UnsupportedOperation.ScaleDownTooMuch"
//  UNSUPPORTEDOPERATION_STATUSNOTNORMAL = "UnsupportedOperation.StatusNotNormal"
//  UNSUPPORTEDOPERATION_STATUSNOTSUPPORT = "UnsupportedOperation.StatusNotSupport"
//  UNSUPPORTEDOPERATION_UPDATEDISKTYPE = "UnsupportedOperation.UpdateDiskType"
//  UNSUPPORTEDOPERATION_WEBSERVICETYPE = "UnsupportedOperation.WebServiceType"
func (c *Client) UpdateInstanceWithContext(ctx context.Context, request *UpdateInstanceRequest) (response *UpdateInstanceResponse, err error) {
    if request == nil {
        request = NewUpdateInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateJdkRequest() (request *UpdateJdkRequest) {
    request = &UpdateJdkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "UpdateJdk")
    
    
    return
}

func NewUpdateJdkResponse() (response *UpdateJdkResponse) {
    response = &UpdateJdkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateJdk
// 更新实例Jdk配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_GC = "InvalidParameter.GC"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  INVALIDPARAMETERVALUE_INVALIDJDK = "InvalidParameterValue.InvalidJDK"
//  RESOURCENOTFOUND_CLUSTERINFONOTFOUND = "ResourceNotFound.ClusterInfoNotFound"
//  RESOURCENOTFOUND_OSSINFONOTFOUND = "ResourceNotFound.OssInfoNotFound"
//  UNSUPPORTEDOPERATION_CLUSTERSTATENOREPLICATION = "UnsupportedOperation.ClusterStateNoReplication"
//  UNSUPPORTEDOPERATION_CLUSTERSTATEUNHEALTH = "UnsupportedOperation.ClusterStateUnHealth"
//  UNSUPPORTEDOPERATION_STATUSNOTSUPPORT = "UnsupportedOperation.StatusNotSupport"
func (c *Client) UpdateJdk(request *UpdateJdkRequest) (response *UpdateJdkResponse, err error) {
    return c.UpdateJdkWithContext(context.Background(), request)
}

// UpdateJdk
// 更新实例Jdk配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_GC = "InvalidParameter.GC"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  INVALIDPARAMETERVALUE_INVALIDJDK = "InvalidParameterValue.InvalidJDK"
//  RESOURCENOTFOUND_CLUSTERINFONOTFOUND = "ResourceNotFound.ClusterInfoNotFound"
//  RESOURCENOTFOUND_OSSINFONOTFOUND = "ResourceNotFound.OssInfoNotFound"
//  UNSUPPORTEDOPERATION_CLUSTERSTATENOREPLICATION = "UnsupportedOperation.ClusterStateNoReplication"
//  UNSUPPORTEDOPERATION_CLUSTERSTATEUNHEALTH = "UnsupportedOperation.ClusterStateUnHealth"
//  UNSUPPORTEDOPERATION_STATUSNOTSUPPORT = "UnsupportedOperation.StatusNotSupport"
func (c *Client) UpdateJdkWithContext(ctx context.Context, request *UpdateJdkRequest) (response *UpdateJdkResponse, err error) {
    if request == nil {
        request = NewUpdateJdkRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateJdk require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateJdkResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateLogstashInstanceRequest() (request *UpdateLogstashInstanceRequest) {
    request = &UpdateLogstashInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "UpdateLogstashInstance")
    
    
    return
}

func NewUpdateLogstashInstanceResponse() (response *UpdateLogstashInstanceResponse) {
    response = &UpdateLogstashInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateLogstashInstance
// 对集群进行节点规格变更，修改实例名称，修改配置，等操作。参数中InstanceId为必传参数，参数传递组合及含义如下：
//
// - InstanceName：修改实例名称(仅用于标识实例)
//
// - NodeNum: 修改实例节点数量（节点横向扩缩容，纵向扩缩容等）
//
// - YMLConfig: 修改实例YML配置
//
// - BindedES：修改绑定的ES集群配置
//
// 以上参数组合只能传递一种，多传或少传均会导致请求失败
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLUSTERRESOURCELIMITERROR = "FailedOperation.ClusterResourceLimitError"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_ORDER = "ResourceInUse.Order"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateLogstashInstance(request *UpdateLogstashInstanceRequest) (response *UpdateLogstashInstanceResponse, err error) {
    return c.UpdateLogstashInstanceWithContext(context.Background(), request)
}

// UpdateLogstashInstance
// 对集群进行节点规格变更，修改实例名称，修改配置，等操作。参数中InstanceId为必传参数，参数传递组合及含义如下：
//
// - InstanceName：修改实例名称(仅用于标识实例)
//
// - NodeNum: 修改实例节点数量（节点横向扩缩容，纵向扩缩容等）
//
// - YMLConfig: 修改实例YML配置
//
// - BindedES：修改绑定的ES集群配置
//
// 以上参数组合只能传递一种，多传或少传均会导致请求失败
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLUSTERRESOURCELIMITERROR = "FailedOperation.ClusterResourceLimitError"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_ORDER = "ResourceInUse.Order"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateLogstashInstanceWithContext(ctx context.Context, request *UpdateLogstashInstanceRequest) (response *UpdateLogstashInstanceResponse, err error) {
    if request == nil {
        request = NewUpdateLogstashInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateLogstashInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateLogstashInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateLogstashPipelineDescRequest() (request *UpdateLogstashPipelineDescRequest) {
    request = &UpdateLogstashPipelineDescRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "UpdateLogstashPipelineDesc")
    
    
    return
}

func NewUpdateLogstashPipelineDescResponse() (response *UpdateLogstashPipelineDescResponse) {
    response = &UpdateLogstashPipelineDescResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateLogstashPipelineDesc
// 用于更新管道描述信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) UpdateLogstashPipelineDesc(request *UpdateLogstashPipelineDescRequest) (response *UpdateLogstashPipelineDescResponse, err error) {
    return c.UpdateLogstashPipelineDescWithContext(context.Background(), request)
}

// UpdateLogstashPipelineDesc
// 用于更新管道描述信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) UpdateLogstashPipelineDescWithContext(ctx context.Context, request *UpdateLogstashPipelineDescRequest) (response *UpdateLogstashPipelineDescResponse, err error) {
    if request == nil {
        request = NewUpdateLogstashPipelineDescRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateLogstashPipelineDesc require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateLogstashPipelineDescResponse()
    err = c.Send(request, response)
    return
}

func NewUpdatePluginsRequest() (request *UpdatePluginsRequest) {
    request = &UpdatePluginsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "UpdatePlugins")
    
    
    return
}

func NewUpdatePluginsResponse() (response *UpdatePluginsResponse) {
    response = &UpdatePluginsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdatePlugins
// 变更插件列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_FILENAMEERROR = "FailedOperation.FileNameError"
//  FAILEDOPERATION_FILESIZEERROR = "FailedOperation.FileSizeError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  INVALIDPARAMETER_INVALIDREGION = "InvalidParameter.InvalidRegion"
//  INVALIDPARAMETERVALUE_CONFIGINFO = "InvalidParameterValue.ConfigInfo"
//  INVALIDPARAMETERVALUE_INSTALLPLUGINLIST = "InvalidParameterValue.InstallPluginList"
//  INVALIDPARAMETERVALUE_PLUGINTYPE = "InvalidParameterValue.PluginType"
//  LIMITEXCEEDED_PLUGININSTALL = "LimitExceeded.PluginInstall"
//  LIMITEXCEEDED_UPDATEITEMLIMIT = "LimitExceeded.UpdateItemLimit"
//  RESOURCENOTFOUND_CLUSTERINFONOTFOUND = "ResourceNotFound.ClusterInfoNotFound"
//  RESOURCENOTFOUND_COSINFONOTFOUND = "ResourceNotFound.CosInfoNotFound"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
//  RESOURCENOTFOUND_WHITELISTNOTFOUND = "ResourceNotFound.WhiteListNotFound"
//  UNSUPPORTEDOPERATION_PLUGIN = "UnsupportedOperation.Plugin"
//  UNSUPPORTEDOPERATION_STATUSNOTNORMAL = "UnsupportedOperation.StatusNotNormal"
//  UNSUPPORTEDOPERATION_STATUSNOTSUPPORT = "UnsupportedOperation.StatusNotSupport"
func (c *Client) UpdatePlugins(request *UpdatePluginsRequest) (response *UpdatePluginsResponse, err error) {
    return c.UpdatePluginsWithContext(context.Background(), request)
}

// UpdatePlugins
// 变更插件列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_FILENAMEERROR = "FailedOperation.FileNameError"
//  FAILEDOPERATION_FILESIZEERROR = "FailedOperation.FileSizeError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  INVALIDPARAMETER_INVALIDREGION = "InvalidParameter.InvalidRegion"
//  INVALIDPARAMETERVALUE_CONFIGINFO = "InvalidParameterValue.ConfigInfo"
//  INVALIDPARAMETERVALUE_INSTALLPLUGINLIST = "InvalidParameterValue.InstallPluginList"
//  INVALIDPARAMETERVALUE_PLUGINTYPE = "InvalidParameterValue.PluginType"
//  LIMITEXCEEDED_PLUGININSTALL = "LimitExceeded.PluginInstall"
//  LIMITEXCEEDED_UPDATEITEMLIMIT = "LimitExceeded.UpdateItemLimit"
//  RESOURCENOTFOUND_CLUSTERINFONOTFOUND = "ResourceNotFound.ClusterInfoNotFound"
//  RESOURCENOTFOUND_COSINFONOTFOUND = "ResourceNotFound.CosInfoNotFound"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
//  RESOURCENOTFOUND_WHITELISTNOTFOUND = "ResourceNotFound.WhiteListNotFound"
//  UNSUPPORTEDOPERATION_PLUGIN = "UnsupportedOperation.Plugin"
//  UNSUPPORTEDOPERATION_STATUSNOTNORMAL = "UnsupportedOperation.StatusNotNormal"
//  UNSUPPORTEDOPERATION_STATUSNOTSUPPORT = "UnsupportedOperation.StatusNotSupport"
func (c *Client) UpdatePluginsWithContext(ctx context.Context, request *UpdatePluginsRequest) (response *UpdatePluginsResponse, err error) {
    if request == nil {
        request = NewUpdatePluginsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdatePlugins require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdatePluginsResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateRequestTargetNodeTypesRequest() (request *UpdateRequestTargetNodeTypesRequest) {
    request = &UpdateRequestTargetNodeTypesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "UpdateRequestTargetNodeTypes")
    
    
    return
}

func NewUpdateRequestTargetNodeTypesResponse() (response *UpdateRequestTargetNodeTypesResponse) {
    response = &UpdateRequestTargetNodeTypesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateRequestTargetNodeTypes
// 更新接收客户端请求的节点类型
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  INVALIDPARAMETER_INVALIDNODETYPE = "InvalidParameter.InvalidNodeType"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INSTANCETYPEERROR = "UnsupportedOperation.InstanceTypeError"
//  UNSUPPORTEDOPERATION_STATUSNOTSUPPORT = "UnsupportedOperation.StatusNotSupport"
func (c *Client) UpdateRequestTargetNodeTypes(request *UpdateRequestTargetNodeTypesRequest) (response *UpdateRequestTargetNodeTypesResponse, err error) {
    return c.UpdateRequestTargetNodeTypesWithContext(context.Background(), request)
}

// UpdateRequestTargetNodeTypes
// 更新接收客户端请求的节点类型
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  INVALIDPARAMETER_INVALIDNODETYPE = "InvalidParameter.InvalidNodeType"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INSTANCETYPEERROR = "UnsupportedOperation.InstanceTypeError"
//  UNSUPPORTEDOPERATION_STATUSNOTSUPPORT = "UnsupportedOperation.StatusNotSupport"
func (c *Client) UpdateRequestTargetNodeTypesWithContext(ctx context.Context, request *UpdateRequestTargetNodeTypesRequest) (response *UpdateRequestTargetNodeTypesResponse, err error) {
    if request == nil {
        request = NewUpdateRequestTargetNodeTypesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateRequestTargetNodeTypes require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateRequestTargetNodeTypesResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateServerlessInstanceRequest() (request *UpdateServerlessInstanceRequest) {
    request = &UpdateServerlessInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "UpdateServerlessInstance")
    
    
    return
}

func NewUpdateServerlessInstanceResponse() (response *UpdateServerlessInstanceResponse) {
    response = &UpdateServerlessInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateServerlessInstance
// 更新Serverless索引
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_ISOLATEDSTATUS = "ResourceInUse.IsolatedStatus"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateServerlessInstance(request *UpdateServerlessInstanceRequest) (response *UpdateServerlessInstanceResponse, err error) {
    return c.UpdateServerlessInstanceWithContext(context.Background(), request)
}

// UpdateServerlessInstance
// 更新Serverless索引
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_ISOLATEDSTATUS = "ResourceInUse.IsolatedStatus"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateServerlessInstanceWithContext(ctx context.Context, request *UpdateServerlessInstanceRequest) (response *UpdateServerlessInstanceResponse, err error) {
    if request == nil {
        request = NewUpdateServerlessInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateServerlessInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateServerlessInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateServerlessSpaceRequest() (request *UpdateServerlessSpaceRequest) {
    request = &UpdateServerlessSpaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "UpdateServerlessSpace")
    
    
    return
}

func NewUpdateServerlessSpaceResponse() (response *UpdateServerlessSpaceResponse) {
    response = &UpdateServerlessSpaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateServerlessSpace
// 更新Serverless索引空间
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateServerlessSpace(request *UpdateServerlessSpaceRequest) (response *UpdateServerlessSpaceResponse, err error) {
    return c.UpdateServerlessSpaceWithContext(context.Background(), request)
}

// UpdateServerlessSpace
// 更新Serverless索引空间
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateServerlessSpaceWithContext(ctx context.Context, request *UpdateServerlessSpaceRequest) (response *UpdateServerlessSpaceResponse, err error) {
    if request == nil {
        request = NewUpdateServerlessSpaceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateServerlessSpace require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateServerlessSpaceResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeInstanceRequest() (request *UpgradeInstanceRequest) {
    request = &UpgradeInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "UpgradeInstance")
    
    
    return
}

func NewUpgradeInstanceResponse() (response *UpgradeInstanceResponse) {
    response = &UpgradeInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpgradeInstance
// 升级ES集群版本
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDESVERSION = "InvalidParameter.InvalidEsVersion"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  INVALIDPARAMETER_INVALIDNODETYPE = "InvalidParameter.InvalidNodeType"
//  INVALIDPARAMETERVALUE_CONFIGINFO = "InvalidParameterValue.ConfigInfo"
//  INVALIDPARAMETERVALUE_UPGRADEMODE = "InvalidParameterValue.UpgradeMode"
//  LIMITEXCEEDED_NODENUMORINDICES = "LimitExceeded.NodeNumOrIndices"
//  LIMITEXCEEDED_RESOURCELIMIT = "LimitExceeded.ResourceLimit"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_SUBNETIP = "ResourceInsufficient.SubnetIp"
//  RESOURCENOTFOUND_CLUSTERINFONOTFOUND = "ResourceNotFound.ClusterInfoNotFound"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
//  RESOURCENOTFOUND_WHITELISTNOTFOUND = "ResourceNotFound.WhiteListNotFound"
//  UNAUTHORIZEDOPERATION_UINNOTINWHITELIST = "UnauthorizedOperation.UinNotInWhiteList"
//  UNSUPPORTEDOPERATION_CLUSTERSTATECLOSE = "UnsupportedOperation.ClusterStateClose"
//  UNSUPPORTEDOPERATION_CLUSTERSTATEUNHEALTH = "UnsupportedOperation.ClusterStateUnHealth"
//  UNSUPPORTEDOPERATION_COSBACKUP = "UnsupportedOperation.CosBackUp"
//  UNSUPPORTEDOPERATION_INDEXSETTINGSREQUIRESET = "UnsupportedOperation.IndexSettingsRequireSet"
//  UNSUPPORTEDOPERATION_LICENSEERROR = "UnsupportedOperation.LicenseError"
//  UNSUPPORTEDOPERATION_PLUGIN = "UnsupportedOperation.Plugin"
//  UNSUPPORTEDOPERATION_STATUSNOTSUPPORT = "UnsupportedOperation.StatusNotSupport"
func (c *Client) UpgradeInstance(request *UpgradeInstanceRequest) (response *UpgradeInstanceResponse, err error) {
    return c.UpgradeInstanceWithContext(context.Background(), request)
}

// UpgradeInstance
// 升级ES集群版本
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDESVERSION = "InvalidParameter.InvalidEsVersion"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  INVALIDPARAMETER_INVALIDNODETYPE = "InvalidParameter.InvalidNodeType"
//  INVALIDPARAMETERVALUE_CONFIGINFO = "InvalidParameterValue.ConfigInfo"
//  INVALIDPARAMETERVALUE_UPGRADEMODE = "InvalidParameterValue.UpgradeMode"
//  LIMITEXCEEDED_NODENUMORINDICES = "LimitExceeded.NodeNumOrIndices"
//  LIMITEXCEEDED_RESOURCELIMIT = "LimitExceeded.ResourceLimit"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_SUBNETIP = "ResourceInsufficient.SubnetIp"
//  RESOURCENOTFOUND_CLUSTERINFONOTFOUND = "ResourceNotFound.ClusterInfoNotFound"
//  RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"
//  RESOURCENOTFOUND_WHITELISTNOTFOUND = "ResourceNotFound.WhiteListNotFound"
//  UNAUTHORIZEDOPERATION_UINNOTINWHITELIST = "UnauthorizedOperation.UinNotInWhiteList"
//  UNSUPPORTEDOPERATION_CLUSTERSTATECLOSE = "UnsupportedOperation.ClusterStateClose"
//  UNSUPPORTEDOPERATION_CLUSTERSTATEUNHEALTH = "UnsupportedOperation.ClusterStateUnHealth"
//  UNSUPPORTEDOPERATION_COSBACKUP = "UnsupportedOperation.CosBackUp"
//  UNSUPPORTEDOPERATION_INDEXSETTINGSREQUIRESET = "UnsupportedOperation.IndexSettingsRequireSet"
//  UNSUPPORTEDOPERATION_LICENSEERROR = "UnsupportedOperation.LicenseError"
//  UNSUPPORTEDOPERATION_PLUGIN = "UnsupportedOperation.Plugin"
//  UNSUPPORTEDOPERATION_STATUSNOTSUPPORT = "UnsupportedOperation.StatusNotSupport"
func (c *Client) UpgradeInstanceWithContext(ctx context.Context, request *UpgradeInstanceRequest) (response *UpgradeInstanceResponse, err error) {
    if request == nil {
        request = NewUpgradeInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeLicenseRequest() (request *UpgradeLicenseRequest) {
    request = &UpgradeLicenseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "UpgradeLicense")
    
    
    return
}

func NewUpgradeLicenseResponse() (response *UpgradeLicenseResponse) {
    response = &UpgradeLicenseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpgradeLicense
// 升级ES商业特性
//
// 可能返回的错误码:
//  FAILEDOPERATION_ESDICTIONARYINFOERROR = "FailedOperation.EsDictionaryInfoError"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  FAILEDOPERATION_NOTAUTHENTICATED = "FailedOperation.NotAuthenticated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDAUTOVOUCHER = "InvalidParameter.InvalidAutoVoucher"
//  INVALIDPARAMETER_INVALIDCOSBACKUPINFO = "InvalidParameter.InvalidCosBackupInfo"
//  INVALIDPARAMETER_INVALIDDISKCOUNT = "InvalidParameter.InvalidDiskCount"
//  INVALIDPARAMETER_INVALIDDISKSIZE = "InvalidParameter.InvalidDiskSize"
//  INVALIDPARAMETER_INVALIDDISKTYPE = "InvalidParameter.InvalidDiskType"
//  INVALIDPARAMETER_INVALIDESVERSION = "InvalidParameter.InvalidEsVersion"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  INVALIDPARAMETER_INVALIDKIBANAPRIVATEPORT = "InvalidParameter.InvalidKibanaPrivatePort"
//  INVALIDPARAMETER_INVALIDNODENUM = "InvalidParameter.InvalidNodeNum"
//  INVALIDPARAMETER_INVALIDNODETYPE = "InvalidParameter.InvalidNodeType"
//  INVALIDPARAMETER_INVALIDOPTYPE = "InvalidParameter.InvalidOpType"
//  INVALIDPARAMETER_INVALIDPRIVATEACCESS = "InvalidParameter.InvalidPrivateAccess"
//  INVALIDPARAMETER_INVALIDRESTARTTYPE = "InvalidParameter.InvalidRestartType"
//  INVALIDPARAMETER_INVALIDTYPE = "InvalidParameter.InvalidType"
//  INVALIDPARAMETER_INVALIDVOUCHERIDS = "InvalidParameter.InvalidVoucherIds"
//  INVALIDPARAMETERVALUE_CHARGETYPE = "InvalidParameterValue.ChargeType"
//  INVALIDPARAMETERVALUE_CONFIGINFO = "InvalidParameterValue.ConfigInfo"
//  INVALIDPARAMETERVALUE_ESCONFIGTYPE = "InvalidParameterValue.EsConfigType"
//  INVALIDPARAMETERVALUE_INSTANCENAME = "InvalidParameterValue.InstanceName"
//  INVALIDPARAMETERVALUE_INVALIDDEADLINE = "InvalidParameterValue.InvalidDeadline"
//  LIMITEXCEEDED_UPDATEITEMLIMIT = "LimitExceeded.UpdateItemLimit"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_ORDER = "ResourceInUse.Order"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCENOTFOUND_CLUSTERINFONOTFOUND = "ResourceNotFound.ClusterInfoNotFound"
//  RESOURCENOTFOUND_DATANODENOTFOUND = "ResourceNotFound.DataNodeNotFound"
//  RESOURCENOTFOUND_DISKINFONOTFOUND = "ResourceNotFound.DiskInfoNotFound"
//  RESOURCENOTFOUND_TRADECGWNOTFOUND = "ResourceNotFound.TradeCgwNotFound"
//  UNAUTHORIZEDOPERATION_UINNOTINWHITELIST = "UnauthorizedOperation.UinNotInWhiteList"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BASICSECURITYTYPE = "UnsupportedOperation.BasicSecurityType"
//  UNSUPPORTEDOPERATION_CLUSTERSTATENOREPLICATION = "UnsupportedOperation.ClusterStateNoReplication"
//  UNSUPPORTEDOPERATION_CLUSTERSTATEUNHEALTH = "UnsupportedOperation.ClusterStateUnHealth"
//  UNSUPPORTEDOPERATION_INSTANCETYPEERROR = "UnsupportedOperation.InstanceTypeError"
//  UNSUPPORTEDOPERATION_LICENSEERROR = "UnsupportedOperation.LicenseError"
//  UNSUPPORTEDOPERATION_LOCALDISK = "UnsupportedOperation.LocalDisk"
//  UNSUPPORTEDOPERATION_MULTIZONESUPGRADE = "UnsupportedOperation.MultiZonesUpgrade"
//  UNSUPPORTEDOPERATION_PLUGIN = "UnsupportedOperation.Plugin"
func (c *Client) UpgradeLicense(request *UpgradeLicenseRequest) (response *UpgradeLicenseResponse, err error) {
    return c.UpgradeLicenseWithContext(context.Background(), request)
}

// UpgradeLicense
// 升级ES商业特性
//
// 可能返回的错误码:
//  FAILEDOPERATION_ESDICTIONARYINFOERROR = "FailedOperation.EsDictionaryInfoError"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  FAILEDOPERATION_NOTAUTHENTICATED = "FailedOperation.NotAuthenticated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDAUTOVOUCHER = "InvalidParameter.InvalidAutoVoucher"
//  INVALIDPARAMETER_INVALIDCOSBACKUPINFO = "InvalidParameter.InvalidCosBackupInfo"
//  INVALIDPARAMETER_INVALIDDISKCOUNT = "InvalidParameter.InvalidDiskCount"
//  INVALIDPARAMETER_INVALIDDISKSIZE = "InvalidParameter.InvalidDiskSize"
//  INVALIDPARAMETER_INVALIDDISKTYPE = "InvalidParameter.InvalidDiskType"
//  INVALIDPARAMETER_INVALIDESVERSION = "InvalidParameter.InvalidEsVersion"
//  INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"
//  INVALIDPARAMETER_INVALIDKIBANAPRIVATEPORT = "InvalidParameter.InvalidKibanaPrivatePort"
//  INVALIDPARAMETER_INVALIDNODENUM = "InvalidParameter.InvalidNodeNum"
//  INVALIDPARAMETER_INVALIDNODETYPE = "InvalidParameter.InvalidNodeType"
//  INVALIDPARAMETER_INVALIDOPTYPE = "InvalidParameter.InvalidOpType"
//  INVALIDPARAMETER_INVALIDPRIVATEACCESS = "InvalidParameter.InvalidPrivateAccess"
//  INVALIDPARAMETER_INVALIDRESTARTTYPE = "InvalidParameter.InvalidRestartType"
//  INVALIDPARAMETER_INVALIDTYPE = "InvalidParameter.InvalidType"
//  INVALIDPARAMETER_INVALIDVOUCHERIDS = "InvalidParameter.InvalidVoucherIds"
//  INVALIDPARAMETERVALUE_CHARGETYPE = "InvalidParameterValue.ChargeType"
//  INVALIDPARAMETERVALUE_CONFIGINFO = "InvalidParameterValue.ConfigInfo"
//  INVALIDPARAMETERVALUE_ESCONFIGTYPE = "InvalidParameterValue.EsConfigType"
//  INVALIDPARAMETERVALUE_INSTANCENAME = "InvalidParameterValue.InstanceName"
//  INVALIDPARAMETERVALUE_INVALIDDEADLINE = "InvalidParameterValue.InvalidDeadline"
//  LIMITEXCEEDED_UPDATEITEMLIMIT = "LimitExceeded.UpdateItemLimit"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_ORDER = "ResourceInUse.Order"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCENOTFOUND_CLUSTERINFONOTFOUND = "ResourceNotFound.ClusterInfoNotFound"
//  RESOURCENOTFOUND_DATANODENOTFOUND = "ResourceNotFound.DataNodeNotFound"
//  RESOURCENOTFOUND_DISKINFONOTFOUND = "ResourceNotFound.DiskInfoNotFound"
//  RESOURCENOTFOUND_TRADECGWNOTFOUND = "ResourceNotFound.TradeCgwNotFound"
//  UNAUTHORIZEDOPERATION_UINNOTINWHITELIST = "UnauthorizedOperation.UinNotInWhiteList"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BASICSECURITYTYPE = "UnsupportedOperation.BasicSecurityType"
//  UNSUPPORTEDOPERATION_CLUSTERSTATENOREPLICATION = "UnsupportedOperation.ClusterStateNoReplication"
//  UNSUPPORTEDOPERATION_CLUSTERSTATEUNHEALTH = "UnsupportedOperation.ClusterStateUnHealth"
//  UNSUPPORTEDOPERATION_INSTANCETYPEERROR = "UnsupportedOperation.InstanceTypeError"
//  UNSUPPORTEDOPERATION_LICENSEERROR = "UnsupportedOperation.LicenseError"
//  UNSUPPORTEDOPERATION_LOCALDISK = "UnsupportedOperation.LocalDisk"
//  UNSUPPORTEDOPERATION_MULTIZONESUPGRADE = "UnsupportedOperation.MultiZonesUpgrade"
//  UNSUPPORTEDOPERATION_PLUGIN = "UnsupportedOperation.Plugin"
func (c *Client) UpgradeLicenseWithContext(ctx context.Context, request *UpgradeLicenseRequest) (response *UpgradeLicenseResponse, err error) {
    if request == nil {
        request = NewUpgradeLicenseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeLicense require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeLicenseResponse()
    err = c.Send(request, response)
    return
}
