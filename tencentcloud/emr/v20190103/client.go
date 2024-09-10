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

package v20190103

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-01-03"

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


func NewAddMetricScaleStrategyRequest() (request *AddMetricScaleStrategyRequest) {
    request = &AddMetricScaleStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "AddMetricScaleStrategy")
    
    
    return
}

func NewAddMetricScaleStrategyResponse() (response *AddMetricScaleStrategyResponse) {
    response = &AddMetricScaleStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddMetricScaleStrategy
// 添加扩缩容规则，按负载和时间
//
// 可能返回的错误码:
//  FAILEDOPERATION_MORESTRATEGYNOTALLOWED = "FailedOperation.MoreStrategyNotAllowed"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INVALIDPARAMETER_INVALIDCOMPAREMETHOD = "InvalidParameter.InvalidCompareMethod"
//  INVALIDPARAMETER_INVALIDCONDITIONNUM = "InvalidParameter.InvalidConditionNum"
//  INVALIDPARAMETER_INVALIDPARAMTERINVALIDSOFTINFO = "InvalidParameter.InvalidParamterInvalidSoftInfo"
//  INVALIDPARAMETER_INVALIDPROCESSMETHOD = "InvalidParameter.InvalidProcessMethod"
//  INVALIDPARAMETER_INVALIDSCALEACTION = "InvalidParameter.InvalidScaleAction"
//  INVALIDPARAMETER_INVALIDSOFTWARE = "InvalidParameter.InvalidSoftWare"
//  INVALIDPARAMETER_INVALIDSTRATEGY = "InvalidParameter.InvalidStrategy"
//  INVALIDPARAMETER_INVALIDSTRATEGYSPEC = "InvalidParameter.InvalidStrategySpec"
//  INVALIDPARAMETER_INVALIDSTRATEGYTYPE = "InvalidParameter.InvalidStrategyType"
//  INVALIDPARAMETER_INVALIDTIMELAYOUT = "InvalidParameter.InvalidTimeLayout"
//  INVALIDPARAMETER_REPEATEDEXECUTIONTIME = "InvalidParameter.RepeatedExecutionTime"
//  INVALIDPARAMETER_REPEATEDSTRATEGYNAME = "InvalidParameter.RepeatedStrategyName"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) AddMetricScaleStrategy(request *AddMetricScaleStrategyRequest) (response *AddMetricScaleStrategyResponse, err error) {
    return c.AddMetricScaleStrategyWithContext(context.Background(), request)
}

// AddMetricScaleStrategy
// 添加扩缩容规则，按负载和时间
//
// 可能返回的错误码:
//  FAILEDOPERATION_MORESTRATEGYNOTALLOWED = "FailedOperation.MoreStrategyNotAllowed"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INVALIDPARAMETER_INVALIDCOMPAREMETHOD = "InvalidParameter.InvalidCompareMethod"
//  INVALIDPARAMETER_INVALIDCONDITIONNUM = "InvalidParameter.InvalidConditionNum"
//  INVALIDPARAMETER_INVALIDPARAMTERINVALIDSOFTINFO = "InvalidParameter.InvalidParamterInvalidSoftInfo"
//  INVALIDPARAMETER_INVALIDPROCESSMETHOD = "InvalidParameter.InvalidProcessMethod"
//  INVALIDPARAMETER_INVALIDSCALEACTION = "InvalidParameter.InvalidScaleAction"
//  INVALIDPARAMETER_INVALIDSOFTWARE = "InvalidParameter.InvalidSoftWare"
//  INVALIDPARAMETER_INVALIDSTRATEGY = "InvalidParameter.InvalidStrategy"
//  INVALIDPARAMETER_INVALIDSTRATEGYSPEC = "InvalidParameter.InvalidStrategySpec"
//  INVALIDPARAMETER_INVALIDSTRATEGYTYPE = "InvalidParameter.InvalidStrategyType"
//  INVALIDPARAMETER_INVALIDTIMELAYOUT = "InvalidParameter.InvalidTimeLayout"
//  INVALIDPARAMETER_REPEATEDEXECUTIONTIME = "InvalidParameter.RepeatedExecutionTime"
//  INVALIDPARAMETER_REPEATEDSTRATEGYNAME = "InvalidParameter.RepeatedStrategyName"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) AddMetricScaleStrategyWithContext(ctx context.Context, request *AddMetricScaleStrategyRequest) (response *AddMetricScaleStrategyResponse, err error) {
    if request == nil {
        request = NewAddMetricScaleStrategyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddMetricScaleStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddMetricScaleStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewAddUsersForUserManagerRequest() (request *AddUsersForUserManagerRequest) {
    request = &AddUsersForUserManagerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "AddUsersForUserManager")
    
    
    return
}

func NewAddUsersForUserManagerResponse() (response *AddUsersForUserManagerResponse) {
    response = &AddUsersForUserManagerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddUsersForUserManager
// 该接口支持安装了OpenLdap组件的集群。
//
// 新增用户列表（用户管理）。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) AddUsersForUserManager(request *AddUsersForUserManagerRequest) (response *AddUsersForUserManagerResponse, err error) {
    return c.AddUsersForUserManagerWithContext(context.Background(), request)
}

// AddUsersForUserManager
// 该接口支持安装了OpenLdap组件的集群。
//
// 新增用户列表（用户管理）。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) AddUsersForUserManagerWithContext(ctx context.Context, request *AddUsersForUserManagerRequest) (response *AddUsersForUserManagerResponse, err error) {
    if request == nil {
        request = NewAddUsersForUserManagerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddUsersForUserManager require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddUsersForUserManagerResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClusterRequest() (request *CreateClusterRequest) {
    request = &CreateClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "CreateCluster")
    
    
    return
}

func NewCreateClusterResponse() (response *CreateClusterResponse) {
    response = &CreateClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCluster
// 创建EMR集群实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETCVMSERVERFAILED = "FailedOperation.GetCvmServerFailed"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INVALIDPARAMETER_HALESSMASTERCOUNT = "InvalidParameter.HALessMasterCount"
//  INVALIDPARAMETER_INCORRECTMASTERCOUNT = "InvalidParameter.IncorrectMasterCount"
//  INVALIDPARAMETER_INVALIDALLNODERESOURCESPEC = "InvalidParameter.InvalidAllNodeResourceSpec"
//  INVALIDPARAMETER_INVALIDCOMPONENT = "InvalidParameter.InvalidComponent"
//  INVALIDPARAMETER_INVALIDCORECOUNT = "InvalidParameter.InvalidCoreCount"
//  INVALIDPARAMETER_INVALIDCOREDISKTYPE = "InvalidParameter.InvalidCoreDiskType"
//  INVALIDPARAMETER_INVALIDCOSBUCKET = "InvalidParameter.InvalidCosBucket"
//  INVALIDPARAMETER_INVALIDDEPENDSERVICEANDENABLEKERBEROSCONFLICT = "InvalidParameter.InvalidDependServiceAndEnableKerberosConflict"
//  INVALIDPARAMETER_INVALIDDISKNUM = "InvalidParameter.InvalidDiskNum"
//  INVALIDPARAMETER_INVALIDINSTANCECHARGETYPE = "InvalidParameter.InvalidInstanceChargeType"
//  INVALIDPARAMETER_INVALIDINSTANCETYPE = "InvalidParameter.InvalidInstanceType"
//  INVALIDPARAMETER_INVALIDMASTERDISKTYPE = "InvalidParameter.InvalidMasterDiskType"
//  INVALIDPARAMETER_INVALIDMETATYPE = "InvalidParameter.InvalidMetaType"
//  INVALIDPARAMETER_INVALIDPASSWORD = "InvalidParameter.InvalidPassword"
//  INVALIDPARAMETER_INVALIDPRODUCTVERSION = "InvalidParameter.InvalidProductVersion"
//  INVALIDPARAMETER_INVALIDRENEWFLAG = "InvalidParameter.InvalidRenewFlag"
//  INVALIDPARAMETER_INVALIDRESOURCESPEC = "InvalidParameter.InvalidResourceSpec"
//  INVALIDPARAMETER_INVALIDSCRIPTBOOTSTRAPACTIONCONFIG = "InvalidParameter.InvalidScriptBootstrapActionConfig"
//  INVALIDPARAMETER_INVALIDSOFTINFO = "InvalidParameter.InvalidSoftInfo"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  INVALIDPARAMETER_KERBEROSSUPPORT = "InvalidParameter.KerberosSupport"
//  RESOURCEINSUFFICIENT_INSTANCEINSUFFICIENT = "ResourceInsufficient.InstanceInsufficient"
//  RESOURCENOTFOUND_SUBNETNOTFOUND = "ResourceNotFound.SubnetNotFound"
//  RESOURCESSOLDOUT_CVMSOLDOUT = "ResourcesSoldOut.CvmSoldOut"
func (c *Client) CreateCluster(request *CreateClusterRequest) (response *CreateClusterResponse, err error) {
    return c.CreateClusterWithContext(context.Background(), request)
}

// CreateCluster
// 创建EMR集群实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETCVMSERVERFAILED = "FailedOperation.GetCvmServerFailed"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INVALIDPARAMETER_HALESSMASTERCOUNT = "InvalidParameter.HALessMasterCount"
//  INVALIDPARAMETER_INCORRECTMASTERCOUNT = "InvalidParameter.IncorrectMasterCount"
//  INVALIDPARAMETER_INVALIDALLNODERESOURCESPEC = "InvalidParameter.InvalidAllNodeResourceSpec"
//  INVALIDPARAMETER_INVALIDCOMPONENT = "InvalidParameter.InvalidComponent"
//  INVALIDPARAMETER_INVALIDCORECOUNT = "InvalidParameter.InvalidCoreCount"
//  INVALIDPARAMETER_INVALIDCOREDISKTYPE = "InvalidParameter.InvalidCoreDiskType"
//  INVALIDPARAMETER_INVALIDCOSBUCKET = "InvalidParameter.InvalidCosBucket"
//  INVALIDPARAMETER_INVALIDDEPENDSERVICEANDENABLEKERBEROSCONFLICT = "InvalidParameter.InvalidDependServiceAndEnableKerberosConflict"
//  INVALIDPARAMETER_INVALIDDISKNUM = "InvalidParameter.InvalidDiskNum"
//  INVALIDPARAMETER_INVALIDINSTANCECHARGETYPE = "InvalidParameter.InvalidInstanceChargeType"
//  INVALIDPARAMETER_INVALIDINSTANCETYPE = "InvalidParameter.InvalidInstanceType"
//  INVALIDPARAMETER_INVALIDMASTERDISKTYPE = "InvalidParameter.InvalidMasterDiskType"
//  INVALIDPARAMETER_INVALIDMETATYPE = "InvalidParameter.InvalidMetaType"
//  INVALIDPARAMETER_INVALIDPASSWORD = "InvalidParameter.InvalidPassword"
//  INVALIDPARAMETER_INVALIDPRODUCTVERSION = "InvalidParameter.InvalidProductVersion"
//  INVALIDPARAMETER_INVALIDRENEWFLAG = "InvalidParameter.InvalidRenewFlag"
//  INVALIDPARAMETER_INVALIDRESOURCESPEC = "InvalidParameter.InvalidResourceSpec"
//  INVALIDPARAMETER_INVALIDSCRIPTBOOTSTRAPACTIONCONFIG = "InvalidParameter.InvalidScriptBootstrapActionConfig"
//  INVALIDPARAMETER_INVALIDSOFTINFO = "InvalidParameter.InvalidSoftInfo"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  INVALIDPARAMETER_KERBEROSSUPPORT = "InvalidParameter.KerberosSupport"
//  RESOURCEINSUFFICIENT_INSTANCEINSUFFICIENT = "ResourceInsufficient.InstanceInsufficient"
//  RESOURCENOTFOUND_SUBNETNOTFOUND = "ResourceNotFound.SubnetNotFound"
//  RESOURCESSOLDOUT_CVMSOLDOUT = "ResourcesSoldOut.CvmSoldOut"
func (c *Client) CreateClusterWithContext(ctx context.Context, request *CreateClusterRequest) (response *CreateClusterResponse, err error) {
    if request == nil {
        request = NewCreateClusterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateClusterResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstanceRequest() (request *CreateInstanceRequest) {
    request = &CreateInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "CreateInstance")
    
    
    return
}

func NewCreateInstanceResponse() (response *CreateInstanceResponse) {
    response = &CreateInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateInstance
// 创建EMR集群实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DUPLICATEORDERNOTALLOWED = "FailedOperation.DuplicateOrderNotAllowed"
//  FAILEDOPERATION_GETCVMSERVERFAILED = "FailedOperation.GetCvmServerFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CHECKQUOTAERR = "InternalError.CheckQuotaErr"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_HALESSMASTERCOUNT = "InvalidParameter.HALessMasterCount"
//  INVALIDPARAMETER_INCORRECTCOMMONCOUNT = "InvalidParameter.IncorrectCommonCount"
//  INVALIDPARAMETER_INCORRECTMASTERCOUNT = "InvalidParameter.IncorrectMasterCount"
//  INVALIDPARAMETER_INVALIDAUTORENEW = "InvalidParameter.InvalidAutoRenew"
//  INVALIDPARAMETER_INVALIDCLIENTTOKEN = "InvalidParameter.InvalidClientToken"
//  INVALIDPARAMETER_INVALIDCOMPONENT = "InvalidParameter.InvalidComponent"
//  INVALIDPARAMETER_INVALIDCORECOUNT = "InvalidParameter.InvalidCoreCount"
//  INVALIDPARAMETER_INVALIDCOSBUCKET = "InvalidParameter.InvalidCosBucket"
//  INVALIDPARAMETER_INVALIDCOSFILEURI = "InvalidParameter.InvalidCosFileURI"
//  INVALIDPARAMETER_INVALIDDISKSIZE = "InvalidParameter.InvalidDiskSize"
//  INVALIDPARAMETER_INVALIDEXTENDFIELD = "InvalidParameter.InvalidExtendField"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDINSTANCETYPE = "InvalidParameter.InvalidInstanceType"
//  INVALIDPARAMETER_INVALIDLOGINSETTING = "InvalidParameter.InvalidLoginSetting"
//  INVALIDPARAMETER_INVALIDMETADATAJDBCURL = "InvalidParameter.InvalidMetaDataJdbcUrl"
//  INVALIDPARAMETER_INVALIDMETATYPE = "InvalidParameter.InvalidMetaType"
//  INVALIDPARAMETER_INVALIDPASSWORD = "InvalidParameter.InvalidPassword"
//  INVALIDPARAMETER_INVALIDPAYMODE = "InvalidParameter.InvalidPaymode"
//  INVALIDPARAMETER_INVALIDPREEXECUTEDFILE = "InvalidParameter.InvalidPreExecutedFile"
//  INVALIDPARAMETER_INVALIDPRODUCTID = "InvalidParameter.InvalidProductId"
//  INVALIDPARAMETER_INVALIDPROJECTID = "InvalidParameter.InvalidProjectId"
//  INVALIDPARAMETER_INVALIDRESOURCESPEC = "InvalidParameter.InvalidResourceSpec"
//  INVALIDPARAMETER_INVALIDSECURITYSUPPORT = "InvalidParameter.InvalidSecuritySupport"
//  INVALIDPARAMETER_INVALIDSERCURITYGRPUPID = "InvalidParameter.InvalidSercurityGrpupId"
//  INVALIDPARAMETER_INVALIDSERVICENAME = "InvalidParameter.InvalidServiceName"
//  INVALIDPARAMETER_INVALIDSOFTDEPLOYINFO = "InvalidParameter.InvalidSoftDeployInfo"
//  INVALIDPARAMETER_INVALIDSOFTINFO = "InvalidParameter.InvalidSoftInfo"
//  INVALIDPARAMETER_INVALIDSOFTWARE = "InvalidParameter.InvalidSoftWare"
//  INVALIDPARAMETER_INVALIDSOFTWARENAME = "InvalidParameter.InvalidSoftWareName"
//  INVALIDPARAMETER_INVALIDSOFTWAREVERSION = "InvalidParameter.InvalidSoftWareVersion"
//  INVALIDPARAMETER_INVALIDSUBNETID = "InvalidParameter.InvalidSubnetId"
//  INVALIDPARAMETER_INVALIDSUPPORTHA = "InvalidParameter.InvalidSupportHA"
//  INVALIDPARAMETER_INVALIDTIMESPAN = "InvalidParameter.InvalidTimeSpan"
//  INVALIDPARAMETER_INVALIDTIMEUNIT = "InvalidParameter.InvalidTimeUnit"
//  INVALIDPARAMETER_INVALIDUNIFYMETA = "InvalidParameter.InvalidUnifyMeta"
//  INVALIDPARAMETER_INVALIDVPCID = "InvalidParameter.InvalidVpcId"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  INVALIDPARAMETER_NOTCONTAINMUSTSELECTSOFTWARE = "InvalidParameter.NotContainMustSelectSoftware"
//  INVALIDPARAMETER_PAYMODERESOURCENOTMATCH = "InvalidParameter.PayModeResourceNotMatch"
//  INVALIDPARAMETER_SOFTWARENOTINPRODUCT = "InvalidParameter.SoftwareNotInProduct"
//  INVALIDPARAMETER_UNGRANTEDPOLICY = "InvalidParameter.UngrantedPolicy"
//  INVALIDPARAMETER_UNGRANTEDROLE = "InvalidParameter.UngrantedRole"
//  INVALIDPARAMETER_ZONERESOURCENOTMATCH = "InvalidParameter.ZoneResourceNotMatch"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_SECURITYGROUPNUMLIMITEXCEEDED = "LimitExceeded.SecurityGroupNumLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT_DISKINSUFFICIENT = "ResourceInsufficient.DiskInsufficient"
//  RESOURCEINSUFFICIENT_INSTANCEINSUFFICIENT = "ResourceInsufficient.InstanceInsufficient"
//  RESOURCENOTFOUND_SUBNETNOTFOUND = "ResourceNotFound.SubnetNotFound"
//  RESOURCENOTFOUND_TAGSNOTFOUND = "ResourceNotFound.TagsNotFound"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  RESOURCESSOLDOUT_CBSSOLDOUT = "ResourcesSoldOut.CbsSoldOut"
//  RESOURCESSOLDOUT_CVMSOLDOUT = "ResourcesSoldOut.CvmSoldOut"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateInstance(request *CreateInstanceRequest) (response *CreateInstanceResponse, err error) {
    return c.CreateInstanceWithContext(context.Background(), request)
}

// CreateInstance
// 创建EMR集群实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DUPLICATEORDERNOTALLOWED = "FailedOperation.DuplicateOrderNotAllowed"
//  FAILEDOPERATION_GETCVMSERVERFAILED = "FailedOperation.GetCvmServerFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CHECKQUOTAERR = "InternalError.CheckQuotaErr"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_HALESSMASTERCOUNT = "InvalidParameter.HALessMasterCount"
//  INVALIDPARAMETER_INCORRECTCOMMONCOUNT = "InvalidParameter.IncorrectCommonCount"
//  INVALIDPARAMETER_INCORRECTMASTERCOUNT = "InvalidParameter.IncorrectMasterCount"
//  INVALIDPARAMETER_INVALIDAUTORENEW = "InvalidParameter.InvalidAutoRenew"
//  INVALIDPARAMETER_INVALIDCLIENTTOKEN = "InvalidParameter.InvalidClientToken"
//  INVALIDPARAMETER_INVALIDCOMPONENT = "InvalidParameter.InvalidComponent"
//  INVALIDPARAMETER_INVALIDCORECOUNT = "InvalidParameter.InvalidCoreCount"
//  INVALIDPARAMETER_INVALIDCOSBUCKET = "InvalidParameter.InvalidCosBucket"
//  INVALIDPARAMETER_INVALIDCOSFILEURI = "InvalidParameter.InvalidCosFileURI"
//  INVALIDPARAMETER_INVALIDDISKSIZE = "InvalidParameter.InvalidDiskSize"
//  INVALIDPARAMETER_INVALIDEXTENDFIELD = "InvalidParameter.InvalidExtendField"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDINSTANCETYPE = "InvalidParameter.InvalidInstanceType"
//  INVALIDPARAMETER_INVALIDLOGINSETTING = "InvalidParameter.InvalidLoginSetting"
//  INVALIDPARAMETER_INVALIDMETADATAJDBCURL = "InvalidParameter.InvalidMetaDataJdbcUrl"
//  INVALIDPARAMETER_INVALIDMETATYPE = "InvalidParameter.InvalidMetaType"
//  INVALIDPARAMETER_INVALIDPASSWORD = "InvalidParameter.InvalidPassword"
//  INVALIDPARAMETER_INVALIDPAYMODE = "InvalidParameter.InvalidPaymode"
//  INVALIDPARAMETER_INVALIDPREEXECUTEDFILE = "InvalidParameter.InvalidPreExecutedFile"
//  INVALIDPARAMETER_INVALIDPRODUCTID = "InvalidParameter.InvalidProductId"
//  INVALIDPARAMETER_INVALIDPROJECTID = "InvalidParameter.InvalidProjectId"
//  INVALIDPARAMETER_INVALIDRESOURCESPEC = "InvalidParameter.InvalidResourceSpec"
//  INVALIDPARAMETER_INVALIDSECURITYSUPPORT = "InvalidParameter.InvalidSecuritySupport"
//  INVALIDPARAMETER_INVALIDSERCURITYGRPUPID = "InvalidParameter.InvalidSercurityGrpupId"
//  INVALIDPARAMETER_INVALIDSERVICENAME = "InvalidParameter.InvalidServiceName"
//  INVALIDPARAMETER_INVALIDSOFTDEPLOYINFO = "InvalidParameter.InvalidSoftDeployInfo"
//  INVALIDPARAMETER_INVALIDSOFTINFO = "InvalidParameter.InvalidSoftInfo"
//  INVALIDPARAMETER_INVALIDSOFTWARE = "InvalidParameter.InvalidSoftWare"
//  INVALIDPARAMETER_INVALIDSOFTWARENAME = "InvalidParameter.InvalidSoftWareName"
//  INVALIDPARAMETER_INVALIDSOFTWAREVERSION = "InvalidParameter.InvalidSoftWareVersion"
//  INVALIDPARAMETER_INVALIDSUBNETID = "InvalidParameter.InvalidSubnetId"
//  INVALIDPARAMETER_INVALIDSUPPORTHA = "InvalidParameter.InvalidSupportHA"
//  INVALIDPARAMETER_INVALIDTIMESPAN = "InvalidParameter.InvalidTimeSpan"
//  INVALIDPARAMETER_INVALIDTIMEUNIT = "InvalidParameter.InvalidTimeUnit"
//  INVALIDPARAMETER_INVALIDUNIFYMETA = "InvalidParameter.InvalidUnifyMeta"
//  INVALIDPARAMETER_INVALIDVPCID = "InvalidParameter.InvalidVpcId"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  INVALIDPARAMETER_NOTCONTAINMUSTSELECTSOFTWARE = "InvalidParameter.NotContainMustSelectSoftware"
//  INVALIDPARAMETER_PAYMODERESOURCENOTMATCH = "InvalidParameter.PayModeResourceNotMatch"
//  INVALIDPARAMETER_SOFTWARENOTINPRODUCT = "InvalidParameter.SoftwareNotInProduct"
//  INVALIDPARAMETER_UNGRANTEDPOLICY = "InvalidParameter.UngrantedPolicy"
//  INVALIDPARAMETER_UNGRANTEDROLE = "InvalidParameter.UngrantedRole"
//  INVALIDPARAMETER_ZONERESOURCENOTMATCH = "InvalidParameter.ZoneResourceNotMatch"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_SECURITYGROUPNUMLIMITEXCEEDED = "LimitExceeded.SecurityGroupNumLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT_DISKINSUFFICIENT = "ResourceInsufficient.DiskInsufficient"
//  RESOURCEINSUFFICIENT_INSTANCEINSUFFICIENT = "ResourceInsufficient.InstanceInsufficient"
//  RESOURCENOTFOUND_SUBNETNOTFOUND = "ResourceNotFound.SubnetNotFound"
//  RESOURCENOTFOUND_TAGSNOTFOUND = "ResourceNotFound.TagsNotFound"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  RESOURCESSOLDOUT_CBSSOLDOUT = "ResourcesSoldOut.CbsSoldOut"
//  RESOURCESSOLDOUT_CVMSOLDOUT = "ResourcesSoldOut.CvmSoldOut"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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

func NewCreateSLInstanceRequest() (request *CreateSLInstanceRequest) {
    request = &CreateSLInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "CreateSLInstance")
    
    
    return
}

func NewCreateSLInstanceResponse() (response *CreateSLInstanceResponse) {
    response = &CreateSLInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSLInstance
// 本接口（CreateSLInstance）用于创建 Lite HBase 实例
//
// - 接口调用成功，会创建Lite HBase实例，创建实例请求成功会返回创建实例的 InstaceId 和请求的 RequestID。
//
// - 接口为异步接口，接口返回时操作并未立即完成，实例操作结果可以通过调用 DescribeInstancesList 查看当前实例的 StatusDesc 状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REFUNDCVMFAILED = "FailedOperation.RefundCvmFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  INVALIDPARAMETER_INVALIDDISKSIZE = "InvalidParameter.InvalidDiskSize"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDNODECOUNT = "InvalidParameter.InvalidNodeCount"
//  INVALIDPARAMETER_INVALIDNODETYPE = "InvalidParameter.InvalidNodeType"
//  INVALIDPARAMETER_INVALIDPASSWORD = "InvalidParameter.InvalidPassword"
//  INVALIDPARAMETER_INVALIDPAYMODE = "InvalidParameter.InvalidPaymode"
//  INVALIDPARAMETER_INVALIDRESOURCESPEC = "InvalidParameter.InvalidResourceSpec"
//  INVALIDPARAMETER_INVALIDUINNUM = "InvalidParameter.InvalidUinNum"
//  RESOURCEUNAVAILABLE_RESOURCESPECNOTEXIST = "ResourceUnavailable.ResourceSpecNotExist"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
//  UNSUPPORTEDOPERATION_SERVICENOTSUPPORT = "UnsupportedOperation.ServiceNotSupport"
func (c *Client) CreateSLInstance(request *CreateSLInstanceRequest) (response *CreateSLInstanceResponse, err error) {
    return c.CreateSLInstanceWithContext(context.Background(), request)
}

// CreateSLInstance
// 本接口（CreateSLInstance）用于创建 Lite HBase 实例
//
// - 接口调用成功，会创建Lite HBase实例，创建实例请求成功会返回创建实例的 InstaceId 和请求的 RequestID。
//
// - 接口为异步接口，接口返回时操作并未立即完成，实例操作结果可以通过调用 DescribeInstancesList 查看当前实例的 StatusDesc 状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REFUNDCVMFAILED = "FailedOperation.RefundCvmFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  INVALIDPARAMETER_INVALIDDISKSIZE = "InvalidParameter.InvalidDiskSize"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDNODECOUNT = "InvalidParameter.InvalidNodeCount"
//  INVALIDPARAMETER_INVALIDNODETYPE = "InvalidParameter.InvalidNodeType"
//  INVALIDPARAMETER_INVALIDPASSWORD = "InvalidParameter.InvalidPassword"
//  INVALIDPARAMETER_INVALIDPAYMODE = "InvalidParameter.InvalidPaymode"
//  INVALIDPARAMETER_INVALIDRESOURCESPEC = "InvalidParameter.InvalidResourceSpec"
//  INVALIDPARAMETER_INVALIDUINNUM = "InvalidParameter.InvalidUinNum"
//  RESOURCEUNAVAILABLE_RESOURCESPECNOTEXIST = "ResourceUnavailable.ResourceSpecNotExist"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
//  UNSUPPORTEDOPERATION_SERVICENOTSUPPORT = "UnsupportedOperation.ServiceNotSupport"
func (c *Client) CreateSLInstanceWithContext(ctx context.Context, request *CreateSLInstanceRequest) (response *CreateSLInstanceResponse, err error) {
    if request == nil {
        request = NewCreateSLInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSLInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSLInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAutoScaleStrategyRequest() (request *DeleteAutoScaleStrategyRequest) {
    request = &DeleteAutoScaleStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DeleteAutoScaleStrategy")
    
    
    return
}

func NewDeleteAutoScaleStrategyResponse() (response *DeleteAutoScaleStrategyResponse) {
    response = &DeleteAutoScaleStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAutoScaleStrategy
// 删除自动扩缩容规则，后台销毁根据该规则扩缩容出来的节点
//
// 可能返回的错误码:
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INVALIDPARAMETER_INVALIDSTRATEGYTYPE = "InvalidParameter.InvalidStrategyType"
//  RESOURCENOTFOUND_STRATEGYNOTFOUND = "ResourceNotFound.StrategyNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) DeleteAutoScaleStrategy(request *DeleteAutoScaleStrategyRequest) (response *DeleteAutoScaleStrategyResponse, err error) {
    return c.DeleteAutoScaleStrategyWithContext(context.Background(), request)
}

// DeleteAutoScaleStrategy
// 删除自动扩缩容规则，后台销毁根据该规则扩缩容出来的节点
//
// 可能返回的错误码:
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INVALIDPARAMETER_INVALIDSTRATEGYTYPE = "InvalidParameter.InvalidStrategyType"
//  RESOURCENOTFOUND_STRATEGYNOTFOUND = "ResourceNotFound.StrategyNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) DeleteAutoScaleStrategyWithContext(ctx context.Context, request *DeleteAutoScaleStrategyRequest) (response *DeleteAutoScaleStrategyResponse, err error) {
    if request == nil {
        request = NewDeleteAutoScaleStrategyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAutoScaleStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAutoScaleStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUserManagerUserListRequest() (request *DeleteUserManagerUserListRequest) {
    request = &DeleteUserManagerUserListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DeleteUserManagerUserList")
    
    
    return
}

func NewDeleteUserManagerUserListResponse() (response *DeleteUserManagerUserListResponse) {
    response = &DeleteUserManagerUserListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteUserManagerUserList
// 删除用户列表（用户管理）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) DeleteUserManagerUserList(request *DeleteUserManagerUserListRequest) (response *DeleteUserManagerUserListResponse, err error) {
    return c.DeleteUserManagerUserListWithContext(context.Background(), request)
}

// DeleteUserManagerUserList
// 删除用户列表（用户管理）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) DeleteUserManagerUserListWithContext(ctx context.Context, request *DeleteUserManagerUserListRequest) (response *DeleteUserManagerUserListResponse, err error) {
    if request == nil {
        request = NewDeleteUserManagerUserListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUserManagerUserList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUserManagerUserListResponse()
    err = c.Send(request, response)
    return
}

func NewDeployYarnConfRequest() (request *DeployYarnConfRequest) {
    request = &DeployYarnConfRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DeployYarnConf")
    
    
    return
}

func NewDeployYarnConfResponse() (response *DeployYarnConfResponse) {
    response = &DeployYarnConfResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeployYarnConf
// yarn资源调度-部署生效
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) DeployYarnConf(request *DeployYarnConfRequest) (response *DeployYarnConfResponse, err error) {
    return c.DeployYarnConfWithContext(context.Background(), request)
}

// DeployYarnConf
// yarn资源调度-部署生效
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) DeployYarnConfWithContext(ctx context.Context, request *DeployYarnConfRequest) (response *DeployYarnConfResponse, err error) {
    if request == nil {
        request = NewDeployYarnConfRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeployYarnConf require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeployYarnConfResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAutoScaleGroupGlobalConfRequest() (request *DescribeAutoScaleGroupGlobalConfRequest) {
    request = &DescribeAutoScaleGroupGlobalConfRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeAutoScaleGroupGlobalConf")
    
    
    return
}

func NewDescribeAutoScaleGroupGlobalConfResponse() (response *DescribeAutoScaleGroupGlobalConfResponse) {
    response = &DescribeAutoScaleGroupGlobalConfResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAutoScaleGroupGlobalConf
// 获取自动扩缩容全局配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) DescribeAutoScaleGroupGlobalConf(request *DescribeAutoScaleGroupGlobalConfRequest) (response *DescribeAutoScaleGroupGlobalConfResponse, err error) {
    return c.DescribeAutoScaleGroupGlobalConfWithContext(context.Background(), request)
}

// DescribeAutoScaleGroupGlobalConf
// 获取自动扩缩容全局配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) DescribeAutoScaleGroupGlobalConfWithContext(ctx context.Context, request *DescribeAutoScaleGroupGlobalConfRequest) (response *DescribeAutoScaleGroupGlobalConfResponse, err error) {
    if request == nil {
        request = NewDescribeAutoScaleGroupGlobalConfRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAutoScaleGroupGlobalConf require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAutoScaleGroupGlobalConfResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAutoScaleRecordsRequest() (request *DescribeAutoScaleRecordsRequest) {
    request = &DescribeAutoScaleRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeAutoScaleRecords")
    
    
    return
}

func NewDescribeAutoScaleRecordsResponse() (response *DescribeAutoScaleRecordsResponse) {
    response = &DescribeAutoScaleRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAutoScaleRecords
// 获取集群的自动扩缩容的详细记录
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDFILTERKEY = "InvalidParameter.InvalidFilterKey"
//  INVALIDPARAMETER_INVALIDSTARTTIMEORENDTIME = "InvalidParameter.InvalidStartTimeOrEndTime"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) DescribeAutoScaleRecords(request *DescribeAutoScaleRecordsRequest) (response *DescribeAutoScaleRecordsResponse, err error) {
    return c.DescribeAutoScaleRecordsWithContext(context.Background(), request)
}

// DescribeAutoScaleRecords
// 获取集群的自动扩缩容的详细记录
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDFILTERKEY = "InvalidParameter.InvalidFilterKey"
//  INVALIDPARAMETER_INVALIDSTARTTIMEORENDTIME = "InvalidParameter.InvalidStartTimeOrEndTime"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) DescribeAutoScaleRecordsWithContext(ctx context.Context, request *DescribeAutoScaleRecordsRequest) (response *DescribeAutoScaleRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeAutoScaleRecordsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAutoScaleRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAutoScaleRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAutoScaleStrategiesRequest() (request *DescribeAutoScaleStrategiesRequest) {
    request = &DescribeAutoScaleStrategiesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeAutoScaleStrategies")
    
    
    return
}

func NewDescribeAutoScaleStrategiesResponse() (response *DescribeAutoScaleStrategiesResponse) {
    response = &DescribeAutoScaleStrategiesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAutoScaleStrategies
// 获取自动扩缩容规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER_INVALIDSTRATEGYTYPE = "InvalidParameter.InvalidStrategyType"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) DescribeAutoScaleStrategies(request *DescribeAutoScaleStrategiesRequest) (response *DescribeAutoScaleStrategiesResponse, err error) {
    return c.DescribeAutoScaleStrategiesWithContext(context.Background(), request)
}

// DescribeAutoScaleStrategies
// 获取自动扩缩容规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER_INVALIDSTRATEGYTYPE = "InvalidParameter.InvalidStrategyType"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) DescribeAutoScaleStrategiesWithContext(ctx context.Context, request *DescribeAutoScaleStrategiesRequest) (response *DescribeAutoScaleStrategiesResponse, err error) {
    if request == nil {
        request = NewDescribeAutoScaleStrategiesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAutoScaleStrategies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAutoScaleStrategiesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterFlowStatusDetailRequest() (request *DescribeClusterFlowStatusDetailRequest) {
    request = &DescribeClusterFlowStatusDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeClusterFlowStatusDetail")
    
    
    return
}

func NewDescribeClusterFlowStatusDetailResponse() (response *DescribeClusterFlowStatusDetailResponse) {
    response = &DescribeClusterFlowStatusDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterFlowStatusDetail
// 查询EMR任务运行详情状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) DescribeClusterFlowStatusDetail(request *DescribeClusterFlowStatusDetailRequest) (response *DescribeClusterFlowStatusDetailResponse, err error) {
    return c.DescribeClusterFlowStatusDetailWithContext(context.Background(), request)
}

// DescribeClusterFlowStatusDetail
// 查询EMR任务运行详情状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) DescribeClusterFlowStatusDetailWithContext(ctx context.Context, request *DescribeClusterFlowStatusDetailRequest) (response *DescribeClusterFlowStatusDetailResponse, err error) {
    if request == nil {
        request = NewDescribeClusterFlowStatusDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterFlowStatusDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterFlowStatusDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterNodesRequest() (request *DescribeClusterNodesRequest) {
    request = &DescribeClusterNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeClusterNodes")
    
    
    return
}

func NewDescribeClusterNodesResponse() (response *DescribeClusterNodesResponse) {
    response = &DescribeClusterNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterNodes
// 查询集群节点信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETCAMROLEFAILED = "FailedOperation.GetCamRoleFailed"
//  FAILEDOPERATION_GETCAMSERVERFAILED = "FailedOperation.GetCamServerFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDNODETYPE = "InvalidParameter.InvalidNodeType"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeClusterNodes(request *DescribeClusterNodesRequest) (response *DescribeClusterNodesResponse, err error) {
    return c.DescribeClusterNodesWithContext(context.Background(), request)
}

// DescribeClusterNodes
// 查询集群节点信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETCAMROLEFAILED = "FailedOperation.GetCamRoleFailed"
//  FAILEDOPERATION_GETCAMSERVERFAILED = "FailedOperation.GetCamServerFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDNODETYPE = "InvalidParameter.InvalidNodeType"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeClusterNodesWithContext(ctx context.Context, request *DescribeClusterNodesRequest) (response *DescribeClusterNodesResponse, err error) {
    if request == nil {
        request = NewDescribeClusterNodesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterNodesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCvmQuotaRequest() (request *DescribeCvmQuotaRequest) {
    request = &DescribeCvmQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeCvmQuota")
    
    
    return
}

func NewDescribeCvmQuotaResponse() (response *DescribeCvmQuotaResponse) {
    response = &DescribeCvmQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCvmQuota
// 获取账户的CVM配额
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_EKSERROR = "InternalError.EKSError"
func (c *Client) DescribeCvmQuota(request *DescribeCvmQuotaRequest) (response *DescribeCvmQuotaResponse, err error) {
    return c.DescribeCvmQuotaWithContext(context.Background(), request)
}

// DescribeCvmQuota
// 获取账户的CVM配额
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_EKSERROR = "InternalError.EKSError"
func (c *Client) DescribeCvmQuotaWithContext(ctx context.Context, request *DescribeCvmQuotaRequest) (response *DescribeCvmQuotaResponse, err error) {
    if request == nil {
        request = NewDescribeCvmQuotaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCvmQuota require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCvmQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEmrApplicationStaticsRequest() (request *DescribeEmrApplicationStaticsRequest) {
    request = &DescribeEmrApplicationStaticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeEmrApplicationStatics")
    
    
    return
}

func NewDescribeEmrApplicationStaticsResponse() (response *DescribeEmrApplicationStaticsResponse) {
    response = &DescribeEmrApplicationStaticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEmrApplicationStatics
// yarn application 统计接口查询
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) DescribeEmrApplicationStatics(request *DescribeEmrApplicationStaticsRequest) (response *DescribeEmrApplicationStaticsResponse, err error) {
    return c.DescribeEmrApplicationStaticsWithContext(context.Background(), request)
}

// DescribeEmrApplicationStatics
// yarn application 统计接口查询
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) DescribeEmrApplicationStaticsWithContext(ctx context.Context, request *DescribeEmrApplicationStaticsRequest) (response *DescribeEmrApplicationStaticsResponse, err error) {
    if request == nil {
        request = NewDescribeEmrApplicationStaticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEmrApplicationStatics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEmrApplicationStaticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEmrOverviewMetricsRequest() (request *DescribeEmrOverviewMetricsRequest) {
    request = &DescribeEmrOverviewMetricsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeEmrOverviewMetrics")
    
    
    return
}

func NewDescribeEmrOverviewMetricsResponse() (response *DescribeEmrOverviewMetricsResponse) {
    response = &DescribeEmrOverviewMetricsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEmrOverviewMetrics
// 查询监控概览页指标数据
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_APPIDMISMATCHED = "UnauthorizedOperation.AppIdMismatched"
func (c *Client) DescribeEmrOverviewMetrics(request *DescribeEmrOverviewMetricsRequest) (response *DescribeEmrOverviewMetricsResponse, err error) {
    return c.DescribeEmrOverviewMetricsWithContext(context.Background(), request)
}

// DescribeEmrOverviewMetrics
// 查询监控概览页指标数据
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_APPIDMISMATCHED = "UnauthorizedOperation.AppIdMismatched"
func (c *Client) DescribeEmrOverviewMetricsWithContext(ctx context.Context, request *DescribeEmrOverviewMetricsRequest) (response *DescribeEmrOverviewMetricsResponse, err error) {
    if request == nil {
        request = NewDescribeEmrOverviewMetricsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEmrOverviewMetrics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEmrOverviewMetricsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHBaseTableOverviewRequest() (request *DescribeHBaseTableOverviewRequest) {
    request = &DescribeHBaseTableOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeHBaseTableOverview")
    
    
    return
}

func NewDescribeHBaseTableOverviewResponse() (response *DescribeHBaseTableOverviewResponse) {
    response = &DescribeHBaseTableOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHBaseTableOverview
// 获取Hbase表级监控数据概览接口
//
// 可能返回的错误码:
//  INTERNALERROR_DOOPENTSDBREQUESTEXCEPTION = "InternalError.DoOpenTSDBRequestException"
//  INTERNALERROR_OPENTSDBHTTPRETURNCODENOTOK = "InternalError.OpenTSDBHttpReturnCodeNotOK"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
func (c *Client) DescribeHBaseTableOverview(request *DescribeHBaseTableOverviewRequest) (response *DescribeHBaseTableOverviewResponse, err error) {
    return c.DescribeHBaseTableOverviewWithContext(context.Background(), request)
}

// DescribeHBaseTableOverview
// 获取Hbase表级监控数据概览接口
//
// 可能返回的错误码:
//  INTERNALERROR_DOOPENTSDBREQUESTEXCEPTION = "InternalError.DoOpenTSDBRequestException"
//  INTERNALERROR_OPENTSDBHTTPRETURNCODENOTOK = "InternalError.OpenTSDBHttpReturnCodeNotOK"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
func (c *Client) DescribeHBaseTableOverviewWithContext(ctx context.Context, request *DescribeHBaseTableOverviewRequest) (response *DescribeHBaseTableOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeHBaseTableOverviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHBaseTableOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHBaseTableOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHiveQueriesRequest() (request *DescribeHiveQueriesRequest) {
    request = &DescribeHiveQueriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeHiveQueries")
    
    
    return
}

func NewDescribeHiveQueriesResponse() (response *DescribeHiveQueriesResponse) {
    response = &DescribeHiveQueriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHiveQueries
// 获取hive查询信息
//
// 可能返回的错误码:
//  INTERNALERROR_DBQUERYEXCEPTION = "InternalError.DBQueryException"
//  INVALIDPARAMETER_IMPALAQUERYEXCEPTION = "InvalidParameter.ImpalaQueryException"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  UNAUTHORIZEDOPERATION_APPIDMISMATCHED = "UnauthorizedOperation.AppIdMismatched"
func (c *Client) DescribeHiveQueries(request *DescribeHiveQueriesRequest) (response *DescribeHiveQueriesResponse, err error) {
    return c.DescribeHiveQueriesWithContext(context.Background(), request)
}

// DescribeHiveQueries
// 获取hive查询信息
//
// 可能返回的错误码:
//  INTERNALERROR_DBQUERYEXCEPTION = "InternalError.DBQueryException"
//  INVALIDPARAMETER_IMPALAQUERYEXCEPTION = "InvalidParameter.ImpalaQueryException"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  UNAUTHORIZEDOPERATION_APPIDMISMATCHED = "UnauthorizedOperation.AppIdMismatched"
func (c *Client) DescribeHiveQueriesWithContext(ctx context.Context, request *DescribeHiveQueriesRequest) (response *DescribeHiveQueriesResponse, err error) {
    if request == nil {
        request = NewDescribeHiveQueriesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHiveQueries require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHiveQueriesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImpalaQueriesRequest() (request *DescribeImpalaQueriesRequest) {
    request = &DescribeImpalaQueriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeImpalaQueries")
    
    
    return
}

func NewDescribeImpalaQueriesResponse() (response *DescribeImpalaQueriesResponse) {
    response = &DescribeImpalaQueriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeImpalaQueries
// DescribeImpalaQueries
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
func (c *Client) DescribeImpalaQueries(request *DescribeImpalaQueriesRequest) (response *DescribeImpalaQueriesResponse, err error) {
    return c.DescribeImpalaQueriesWithContext(context.Background(), request)
}

// DescribeImpalaQueries
// DescribeImpalaQueries
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
func (c *Client) DescribeImpalaQueriesWithContext(ctx context.Context, request *DescribeImpalaQueriesRequest) (response *DescribeImpalaQueriesResponse, err error) {
    if request == nil {
        request = NewDescribeImpalaQueriesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImpalaQueries require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImpalaQueriesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInsightListRequest() (request *DescribeInsightListRequest) {
    request = &DescribeInsightListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeInsightList")
    
    
    return
}

func NewDescribeInsightListResponse() (response *DescribeInsightListResponse) {
    response = &DescribeInsightListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInsightList
// 获取洞察结果信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  UNAUTHORIZEDOPERATION_APPIDMISMATCHED = "UnauthorizedOperation.AppIdMismatched"
func (c *Client) DescribeInsightList(request *DescribeInsightListRequest) (response *DescribeInsightListResponse, err error) {
    return c.DescribeInsightListWithContext(context.Background(), request)
}

// DescribeInsightList
// 获取洞察结果信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  UNAUTHORIZEDOPERATION_APPIDMISMATCHED = "UnauthorizedOperation.AppIdMismatched"
func (c *Client) DescribeInsightListWithContext(ctx context.Context, request *DescribeInsightListRequest) (response *DescribeInsightListResponse, err error) {
    if request == nil {
        request = NewDescribeInsightListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInsightList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInsightListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceRenewNodesRequest() (request *DescribeInstanceRenewNodesRequest) {
    request = &DescribeInstanceRenewNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeInstanceRenewNodes")
    
    
    return
}

func NewDescribeInstanceRenewNodesResponse() (response *DescribeInstanceRenewNodesResponse) {
    response = &DescribeInstanceRenewNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceRenewNodes
// 查询待续费节点信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_GETCAMSERVERFAILED = "FailedOperation.GetCamServerFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDNODETYPE = "InvalidParameter.InvalidNodeType"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInstanceRenewNodes(request *DescribeInstanceRenewNodesRequest) (response *DescribeInstanceRenewNodesResponse, err error) {
    return c.DescribeInstanceRenewNodesWithContext(context.Background(), request)
}

// DescribeInstanceRenewNodes
// 查询待续费节点信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_GETCAMSERVERFAILED = "FailedOperation.GetCamServerFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDNODETYPE = "InvalidParameter.InvalidNodeType"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInstanceRenewNodesWithContext(ctx context.Context, request *DescribeInstanceRenewNodesRequest) (response *DescribeInstanceRenewNodesResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceRenewNodesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceRenewNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceRenewNodesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeInstances")
    
    
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstances
// 查询集群实例信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DESCRIBERESOURCETAGSFAILED = "FailedOperation.DescribeResourceTagsFailed"
//  FAILEDOPERATION_GETCAMROLEFAILED = "FailedOperation.GetCamRoleFailed"
//  FAILEDOPERATION_GETCAMSERVERFAILED = "FailedOperation.GetCamServerFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DISPLAYSTRATEGYNOTMATCH = "InvalidParameter.DisplayStrategyNotMatch"
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_ORDERFIELDNOTMATCH = "InvalidParameter.OrderFieldNotMatch"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_HARDWAREINFONOTFOUND = "ResourceNotFound.HardwareInfoNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    return c.DescribeInstancesWithContext(context.Background(), request)
}

// DescribeInstances
// 查询集群实例信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DESCRIBERESOURCETAGSFAILED = "FailedOperation.DescribeResourceTagsFailed"
//  FAILEDOPERATION_GETCAMROLEFAILED = "FailedOperation.GetCamRoleFailed"
//  FAILEDOPERATION_GETCAMSERVERFAILED = "FailedOperation.GetCamServerFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DISPLAYSTRATEGYNOTMATCH = "InvalidParameter.DisplayStrategyNotMatch"
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_ORDERFIELDNOTMATCH = "InvalidParameter.OrderFieldNotMatch"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_HARDWAREINFONOTFOUND = "ResourceNotFound.HardwareInfoNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
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

func NewDescribeInstancesListRequest() (request *DescribeInstancesListRequest) {
    request = &DescribeInstancesListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeInstancesList")
    
    
    return
}

func NewDescribeInstancesListResponse() (response *DescribeInstancesListResponse) {
    response = &DescribeInstancesListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstancesList
// 查询集群列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DESCRIBERESOURCETAGSFAILED = "FailedOperation.DescribeResourceTagsFailed"
//  FAILEDOPERATION_GETCAMROLEFAILED = "FailedOperation.GetCamRoleFailed"
//  FAILEDOPERATION_GETCAMSERVERFAILED = "FailedOperation.GetCamServerFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DISPLAYSTRATEGYNOTMATCH = "InvalidParameter.DisplayStrategyNotMatch"
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  INVALIDPARAMETER_ORDERFIELDNOTMATCH = "InvalidParameter.OrderFieldNotMatch"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
//  UNSUPPORTEDOPERATION_NOTINWHITELIST = "UnsupportedOperation.NotInWhiteList"
func (c *Client) DescribeInstancesList(request *DescribeInstancesListRequest) (response *DescribeInstancesListResponse, err error) {
    return c.DescribeInstancesListWithContext(context.Background(), request)
}

// DescribeInstancesList
// 查询集群列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DESCRIBERESOURCETAGSFAILED = "FailedOperation.DescribeResourceTagsFailed"
//  FAILEDOPERATION_GETCAMROLEFAILED = "FailedOperation.GetCamRoleFailed"
//  FAILEDOPERATION_GETCAMSERVERFAILED = "FailedOperation.GetCamServerFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DISPLAYSTRATEGYNOTMATCH = "InvalidParameter.DisplayStrategyNotMatch"
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  INVALIDPARAMETER_ORDERFIELDNOTMATCH = "InvalidParameter.OrderFieldNotMatch"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
//  UNSUPPORTEDOPERATION_NOTINWHITELIST = "UnsupportedOperation.NotInWhiteList"
func (c *Client) DescribeInstancesListWithContext(ctx context.Context, request *DescribeInstancesListRequest) (response *DescribeInstancesListResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstancesList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeJobFlowRequest() (request *DescribeJobFlowRequest) {
    request = &DescribeJobFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeJobFlow")
    
    
    return
}

func NewDescribeJobFlowResponse() (response *DescribeJobFlowResponse) {
    response = &DescribeJobFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeJobFlow
// 查询流程任务
//
// 可能返回的错误码:
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER_INVALIDJOBFLOW = "InvalidParameter.InvalidJobFlow"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
func (c *Client) DescribeJobFlow(request *DescribeJobFlowRequest) (response *DescribeJobFlowResponse, err error) {
    return c.DescribeJobFlowWithContext(context.Background(), request)
}

// DescribeJobFlow
// 查询流程任务
//
// 可能返回的错误码:
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER_INVALIDJOBFLOW = "InvalidParameter.InvalidJobFlow"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
func (c *Client) DescribeJobFlowWithContext(ctx context.Context, request *DescribeJobFlowRequest) (response *DescribeJobFlowResponse, err error) {
    if request == nil {
        request = NewDescribeJobFlowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeJobFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeJobFlowResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourceScheduleRequest() (request *DescribeResourceScheduleRequest) {
    request = &DescribeResourceScheduleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeResourceSchedule")
    
    
    return
}

func NewDescribeResourceScheduleResponse() (response *DescribeResourceScheduleResponse) {
    response = &DescribeResourceScheduleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeResourceSchedule
// 查询YARN资源调度数据信息。已废弃，请使用`DescribeYarnQueue`去查询队列信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) DescribeResourceSchedule(request *DescribeResourceScheduleRequest) (response *DescribeResourceScheduleResponse, err error) {
    return c.DescribeResourceScheduleWithContext(context.Background(), request)
}

// DescribeResourceSchedule
// 查询YARN资源调度数据信息。已废弃，请使用`DescribeYarnQueue`去查询队列信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) DescribeResourceScheduleWithContext(ctx context.Context, request *DescribeResourceScheduleRequest) (response *DescribeResourceScheduleResponse, err error) {
    if request == nil {
        request = NewDescribeResourceScheduleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResourceSchedule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResourceScheduleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourceScheduleDiffDetailRequest() (request *DescribeResourceScheduleDiffDetailRequest) {
    request = &DescribeResourceScheduleDiffDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeResourceScheduleDiffDetail")
    
    
    return
}

func NewDescribeResourceScheduleDiffDetailResponse() (response *DescribeResourceScheduleDiffDetailResponse) {
    response = &DescribeResourceScheduleDiffDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeResourceScheduleDiffDetail
// YARN资源调度-变更详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) DescribeResourceScheduleDiffDetail(request *DescribeResourceScheduleDiffDetailRequest) (response *DescribeResourceScheduleDiffDetailResponse, err error) {
    return c.DescribeResourceScheduleDiffDetailWithContext(context.Background(), request)
}

// DescribeResourceScheduleDiffDetail
// YARN资源调度-变更详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) DescribeResourceScheduleDiffDetailWithContext(ctx context.Context, request *DescribeResourceScheduleDiffDetailRequest) (response *DescribeResourceScheduleDiffDetailResponse, err error) {
    if request == nil {
        request = NewDescribeResourceScheduleDiffDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResourceScheduleDiffDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResourceScheduleDiffDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSLInstanceRequest() (request *DescribeSLInstanceRequest) {
    request = &DescribeSLInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeSLInstance")
    
    
    return
}

func NewDescribeSLInstanceResponse() (response *DescribeSLInstanceResponse) {
    response = &DescribeSLInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSLInstance
// 本接口（DescribeSLInstance）用于查询 Lite HBase 实例基本信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REFUNDCVMFAILED = "FailedOperation.RefundCvmFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_RESOURCESPECNOTEXIST = "ResourceUnavailable.ResourceSpecNotExist"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
//  UNSUPPORTEDOPERATION_SERVICENOTSUPPORT = "UnsupportedOperation.ServiceNotSupport"
func (c *Client) DescribeSLInstance(request *DescribeSLInstanceRequest) (response *DescribeSLInstanceResponse, err error) {
    return c.DescribeSLInstanceWithContext(context.Background(), request)
}

// DescribeSLInstance
// 本接口（DescribeSLInstance）用于查询 Lite HBase 实例基本信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REFUNDCVMFAILED = "FailedOperation.RefundCvmFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_RESOURCESPECNOTEXIST = "ResourceUnavailable.ResourceSpecNotExist"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
//  UNSUPPORTEDOPERATION_SERVICENOTSUPPORT = "UnsupportedOperation.ServiceNotSupport"
func (c *Client) DescribeSLInstanceWithContext(ctx context.Context, request *DescribeSLInstanceRequest) (response *DescribeSLInstanceResponse, err error) {
    if request == nil {
        request = NewDescribeSLInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSLInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSLInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSLInstanceListRequest() (request *DescribeSLInstanceListRequest) {
    request = &DescribeSLInstanceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeSLInstanceList")
    
    
    return
}

func NewDescribeSLInstanceListResponse() (response *DescribeSLInstanceListResponse) {
    response = &DescribeSLInstanceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSLInstanceList
// 本接口（DescribeSLInstanceList）用于查询 Lite HBase 实例列表详细信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DUPLICATEORDERNOTALLOWED = "FailedOperation.DuplicateOrderNotAllowed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CHECKQUOTAERR = "InternalError.CheckQuotaErr"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DISPLAYSTRATEGYNOTMATCH = "InvalidParameter.DisplayStrategyNotMatch"
//  INVALIDPARAMETER_INVALIDAUTORENEW = "InvalidParameter.InvalidAutoRenew"
//  INVALIDPARAMETER_INVALIDCLIENTTOKEN = "InvalidParameter.InvalidClientToken"
//  INVALIDPARAMETER_INVALIDCOMPONENT = "InvalidParameter.InvalidComponent"
//  INVALIDPARAMETER_INVALIDCORECOUNT = "InvalidParameter.InvalidCoreCount"
//  INVALIDPARAMETER_INVALIDDISKSIZE = "InvalidParameter.InvalidDiskSize"
//  INVALIDPARAMETER_INVALIDEXTENDFIELD = "InvalidParameter.InvalidExtendField"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDINSTANCETYPE = "InvalidParameter.InvalidInstanceType"
//  INVALIDPARAMETER_INVALIDLOGINSETTING = "InvalidParameter.InvalidLoginSetting"
//  INVALIDPARAMETER_INVALIDMETATYPE = "InvalidParameter.InvalidMetaType"
//  INVALIDPARAMETER_INVALIDPASSWORD = "InvalidParameter.InvalidPassword"
//  INVALIDPARAMETER_INVALIDPAYMODE = "InvalidParameter.InvalidPaymode"
//  INVALIDPARAMETER_INVALIDPREEXECUTEDFILE = "InvalidParameter.InvalidPreExecutedFile"
//  INVALIDPARAMETER_INVALIDPRODUCTID = "InvalidParameter.InvalidProductId"
//  INVALIDPARAMETER_INVALIDPROJECTID = "InvalidParameter.InvalidProjectId"
//  INVALIDPARAMETER_INVALIDRESOURCESPEC = "InvalidParameter.InvalidResourceSpec"
//  INVALIDPARAMETER_INVALIDSERCURITYGRPUPID = "InvalidParameter.InvalidSercurityGrpupId"
//  INVALIDPARAMETER_INVALIDSERVICENAME = "InvalidParameter.InvalidServiceName"
//  INVALIDPARAMETER_INVALIDSOFTDEPLOYINFO = "InvalidParameter.InvalidSoftDeployInfo"
//  INVALIDPARAMETER_INVALIDSOFTINFO = "InvalidParameter.InvalidSoftInfo"
//  INVALIDPARAMETER_INVALIDSOFTWARE = "InvalidParameter.InvalidSoftWare"
//  INVALIDPARAMETER_INVALIDSOFTWARENAME = "InvalidParameter.InvalidSoftWareName"
//  INVALIDPARAMETER_INVALIDSOFTWAREVERSION = "InvalidParameter.InvalidSoftWareVersion"
//  INVALIDPARAMETER_INVALIDSUBNETID = "InvalidParameter.InvalidSubnetId"
//  INVALIDPARAMETER_INVALIDSUPPORTHA = "InvalidParameter.InvalidSupportHA"
//  INVALIDPARAMETER_INVALIDTIMESPAN = "InvalidParameter.InvalidTimeSpan"
//  INVALIDPARAMETER_INVALIDTIMEUNIT = "InvalidParameter.InvalidTimeUnit"
//  INVALIDPARAMETER_INVALIDVPCID = "InvalidParameter.InvalidVpcId"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  INVALIDPARAMETER_NOTCONTAINMUSTSELECTSOFTWARE = "InvalidParameter.NotContainMustSelectSoftware"
//  INVALIDPARAMETER_ORDERFIELDNOTMATCH = "InvalidParameter.OrderFieldNotMatch"
//  INVALIDPARAMETER_PAYMODERESOURCENOTMATCH = "InvalidParameter.PayModeResourceNotMatch"
//  INVALIDPARAMETER_SOFTWARENOTINPRODUCT = "InvalidParameter.SoftwareNotInProduct"
//  INVALIDPARAMETER_UNGRANTEDPOLICY = "InvalidParameter.UngrantedPolicy"
//  INVALIDPARAMETER_UNGRANTEDROLE = "InvalidParameter.UngrantedRole"
//  INVALIDPARAMETER_ZONERESOURCENOTMATCH = "InvalidParameter.ZoneResourceNotMatch"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT_DISKINSUFFICIENT = "ResourceInsufficient.DiskInsufficient"
//  RESOURCEINSUFFICIENT_INSTANCEINSUFFICIENT = "ResourceInsufficient.InstanceInsufficient"
//  RESOURCENOTFOUND_TAGSNOTFOUND = "ResourceNotFound.TagsNotFound"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  RESOURCESSOLDOUT_CBSSOLDOUT = "ResourcesSoldOut.CbsSoldOut"
//  RESOURCESSOLDOUT_CVMSOLDOUT = "ResourcesSoldOut.CvmSoldOut"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSLInstanceList(request *DescribeSLInstanceListRequest) (response *DescribeSLInstanceListResponse, err error) {
    return c.DescribeSLInstanceListWithContext(context.Background(), request)
}

// DescribeSLInstanceList
// 本接口（DescribeSLInstanceList）用于查询 Lite HBase 实例列表详细信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DUPLICATEORDERNOTALLOWED = "FailedOperation.DuplicateOrderNotAllowed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CHECKQUOTAERR = "InternalError.CheckQuotaErr"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DISPLAYSTRATEGYNOTMATCH = "InvalidParameter.DisplayStrategyNotMatch"
//  INVALIDPARAMETER_INVALIDAUTORENEW = "InvalidParameter.InvalidAutoRenew"
//  INVALIDPARAMETER_INVALIDCLIENTTOKEN = "InvalidParameter.InvalidClientToken"
//  INVALIDPARAMETER_INVALIDCOMPONENT = "InvalidParameter.InvalidComponent"
//  INVALIDPARAMETER_INVALIDCORECOUNT = "InvalidParameter.InvalidCoreCount"
//  INVALIDPARAMETER_INVALIDDISKSIZE = "InvalidParameter.InvalidDiskSize"
//  INVALIDPARAMETER_INVALIDEXTENDFIELD = "InvalidParameter.InvalidExtendField"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDINSTANCETYPE = "InvalidParameter.InvalidInstanceType"
//  INVALIDPARAMETER_INVALIDLOGINSETTING = "InvalidParameter.InvalidLoginSetting"
//  INVALIDPARAMETER_INVALIDMETATYPE = "InvalidParameter.InvalidMetaType"
//  INVALIDPARAMETER_INVALIDPASSWORD = "InvalidParameter.InvalidPassword"
//  INVALIDPARAMETER_INVALIDPAYMODE = "InvalidParameter.InvalidPaymode"
//  INVALIDPARAMETER_INVALIDPREEXECUTEDFILE = "InvalidParameter.InvalidPreExecutedFile"
//  INVALIDPARAMETER_INVALIDPRODUCTID = "InvalidParameter.InvalidProductId"
//  INVALIDPARAMETER_INVALIDPROJECTID = "InvalidParameter.InvalidProjectId"
//  INVALIDPARAMETER_INVALIDRESOURCESPEC = "InvalidParameter.InvalidResourceSpec"
//  INVALIDPARAMETER_INVALIDSERCURITYGRPUPID = "InvalidParameter.InvalidSercurityGrpupId"
//  INVALIDPARAMETER_INVALIDSERVICENAME = "InvalidParameter.InvalidServiceName"
//  INVALIDPARAMETER_INVALIDSOFTDEPLOYINFO = "InvalidParameter.InvalidSoftDeployInfo"
//  INVALIDPARAMETER_INVALIDSOFTINFO = "InvalidParameter.InvalidSoftInfo"
//  INVALIDPARAMETER_INVALIDSOFTWARE = "InvalidParameter.InvalidSoftWare"
//  INVALIDPARAMETER_INVALIDSOFTWARENAME = "InvalidParameter.InvalidSoftWareName"
//  INVALIDPARAMETER_INVALIDSOFTWAREVERSION = "InvalidParameter.InvalidSoftWareVersion"
//  INVALIDPARAMETER_INVALIDSUBNETID = "InvalidParameter.InvalidSubnetId"
//  INVALIDPARAMETER_INVALIDSUPPORTHA = "InvalidParameter.InvalidSupportHA"
//  INVALIDPARAMETER_INVALIDTIMESPAN = "InvalidParameter.InvalidTimeSpan"
//  INVALIDPARAMETER_INVALIDTIMEUNIT = "InvalidParameter.InvalidTimeUnit"
//  INVALIDPARAMETER_INVALIDVPCID = "InvalidParameter.InvalidVpcId"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  INVALIDPARAMETER_NOTCONTAINMUSTSELECTSOFTWARE = "InvalidParameter.NotContainMustSelectSoftware"
//  INVALIDPARAMETER_ORDERFIELDNOTMATCH = "InvalidParameter.OrderFieldNotMatch"
//  INVALIDPARAMETER_PAYMODERESOURCENOTMATCH = "InvalidParameter.PayModeResourceNotMatch"
//  INVALIDPARAMETER_SOFTWARENOTINPRODUCT = "InvalidParameter.SoftwareNotInProduct"
//  INVALIDPARAMETER_UNGRANTEDPOLICY = "InvalidParameter.UngrantedPolicy"
//  INVALIDPARAMETER_UNGRANTEDROLE = "InvalidParameter.UngrantedRole"
//  INVALIDPARAMETER_ZONERESOURCENOTMATCH = "InvalidParameter.ZoneResourceNotMatch"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT_DISKINSUFFICIENT = "ResourceInsufficient.DiskInsufficient"
//  RESOURCEINSUFFICIENT_INSTANCEINSUFFICIENT = "ResourceInsufficient.InstanceInsufficient"
//  RESOURCENOTFOUND_TAGSNOTFOUND = "ResourceNotFound.TagsNotFound"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  RESOURCESSOLDOUT_CBSSOLDOUT = "ResourcesSoldOut.CbsSoldOut"
//  RESOURCESSOLDOUT_CVMSOLDOUT = "ResourcesSoldOut.CvmSoldOut"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSLInstanceListWithContext(ctx context.Context, request *DescribeSLInstanceListRequest) (response *DescribeSLInstanceListResponse, err error) {
    if request == nil {
        request = NewDescribeSLInstanceListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSLInstanceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSLInstanceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceNodeInfosRequest() (request *DescribeServiceNodeInfosRequest) {
    request = &DescribeServiceNodeInfosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeServiceNodeInfos")
    
    
    return
}

func NewDescribeServiceNodeInfosResponse() (response *DescribeServiceNodeInfosResponse) {
    response = &DescribeServiceNodeInfosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeServiceNodeInfos
// 查询服务进程信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
//  UNSUPPORTEDOPERATION_SERVICENOTSUPPORT = "UnsupportedOperation.ServiceNotSupport"
func (c *Client) DescribeServiceNodeInfos(request *DescribeServiceNodeInfosRequest) (response *DescribeServiceNodeInfosResponse, err error) {
    return c.DescribeServiceNodeInfosWithContext(context.Background(), request)
}

// DescribeServiceNodeInfos
// 查询服务进程信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
//  UNSUPPORTEDOPERATION_SERVICENOTSUPPORT = "UnsupportedOperation.ServiceNotSupport"
func (c *Client) DescribeServiceNodeInfosWithContext(ctx context.Context, request *DescribeServiceNodeInfosRequest) (response *DescribeServiceNodeInfosResponse, err error) {
    if request == nil {
        request = NewDescribeServiceNodeInfosRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServiceNodeInfos require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServiceNodeInfosResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTrinoQueryInfoRequest() (request *DescribeTrinoQueryInfoRequest) {
    request = &DescribeTrinoQueryInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeTrinoQueryInfo")
    
    
    return
}

func NewDescribeTrinoQueryInfoResponse() (response *DescribeTrinoQueryInfoResponse) {
    response = &DescribeTrinoQueryInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTrinoQueryInfo
// 获取trino查询结果
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  UNAUTHORIZEDOPERATION_APPIDMISMATCHED = "UnauthorizedOperation.AppIdMismatched"
func (c *Client) DescribeTrinoQueryInfo(request *DescribeTrinoQueryInfoRequest) (response *DescribeTrinoQueryInfoResponse, err error) {
    return c.DescribeTrinoQueryInfoWithContext(context.Background(), request)
}

// DescribeTrinoQueryInfo
// 获取trino查询结果
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  UNAUTHORIZEDOPERATION_APPIDMISMATCHED = "UnauthorizedOperation.AppIdMismatched"
func (c *Client) DescribeTrinoQueryInfoWithContext(ctx context.Context, request *DescribeTrinoQueryInfoRequest) (response *DescribeTrinoQueryInfoResponse, err error) {
    if request == nil {
        request = NewDescribeTrinoQueryInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTrinoQueryInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTrinoQueryInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsersForUserManagerRequest() (request *DescribeUsersForUserManagerRequest) {
    request = &DescribeUsersForUserManagerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeUsersForUserManager")
    
    
    return
}

func NewDescribeUsersForUserManagerResponse() (response *DescribeUsersForUserManagerResponse) {
    response = &DescribeUsersForUserManagerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUsersForUserManager
// 该接口支持安装了OpenLdap组件的集群。
//
// 批量导出用户。对于kerberos集群，如果需要kertab文件下载地址，可以将NeedKeytabInfo设置为true；注意SupportDownLoadKeyTab为true，但是DownLoadKeyTabUrl为空字符串，表示keytab文件在后台没有准备好（正在生成）。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) DescribeUsersForUserManager(request *DescribeUsersForUserManagerRequest) (response *DescribeUsersForUserManagerResponse, err error) {
    return c.DescribeUsersForUserManagerWithContext(context.Background(), request)
}

// DescribeUsersForUserManager
// 该接口支持安装了OpenLdap组件的集群。
//
// 批量导出用户。对于kerberos集群，如果需要kertab文件下载地址，可以将NeedKeytabInfo设置为true；注意SupportDownLoadKeyTab为true，但是DownLoadKeyTabUrl为空字符串，表示keytab文件在后台没有准备好（正在生成）。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) DescribeUsersForUserManagerWithContext(ctx context.Context, request *DescribeUsersForUserManagerRequest) (response *DescribeUsersForUserManagerResponse, err error) {
    if request == nil {
        request = NewDescribeUsersForUserManagerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUsersForUserManager require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUsersForUserManagerResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeYarnApplicationsRequest() (request *DescribeYarnApplicationsRequest) {
    request = &DescribeYarnApplicationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeYarnApplications")
    
    
    return
}

func NewDescribeYarnApplicationsResponse() (response *DescribeYarnApplicationsResponse) {
    response = &DescribeYarnApplicationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeYarnApplications
// DescribeYarnApplications
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IMPALAQUERYEXCEPTION = "InvalidParameter.ImpalaQueryException"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  UNAUTHORIZEDOPERATION_APPIDMISMATCHED = "UnauthorizedOperation.AppIdMismatched"
func (c *Client) DescribeYarnApplications(request *DescribeYarnApplicationsRequest) (response *DescribeYarnApplicationsResponse, err error) {
    return c.DescribeYarnApplicationsWithContext(context.Background(), request)
}

// DescribeYarnApplications
// DescribeYarnApplications
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IMPALAQUERYEXCEPTION = "InvalidParameter.ImpalaQueryException"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  UNAUTHORIZEDOPERATION_APPIDMISMATCHED = "UnauthorizedOperation.AppIdMismatched"
func (c *Client) DescribeYarnApplicationsWithContext(ctx context.Context, request *DescribeYarnApplicationsRequest) (response *DescribeYarnApplicationsResponse, err error) {
    if request == nil {
        request = NewDescribeYarnApplicationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeYarnApplications require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeYarnApplicationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeYarnQueueRequest() (request *DescribeYarnQueueRequest) {
    request = &DescribeYarnQueueRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeYarnQueue")
    
    
    return
}

func NewDescribeYarnQueueResponse() (response *DescribeYarnQueueResponse) {
    response = &DescribeYarnQueueResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeYarnQueue
// 获取资源调度中的队列信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) DescribeYarnQueue(request *DescribeYarnQueueRequest) (response *DescribeYarnQueueResponse, err error) {
    return c.DescribeYarnQueueWithContext(context.Background(), request)
}

// DescribeYarnQueue
// 获取资源调度中的队列信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) DescribeYarnQueueWithContext(ctx context.Context, request *DescribeYarnQueueRequest) (response *DescribeYarnQueueResponse, err error) {
    if request == nil {
        request = NewDescribeYarnQueueRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeYarnQueue require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeYarnQueueResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeYarnScheduleHistoryRequest() (request *DescribeYarnScheduleHistoryRequest) {
    request = &DescribeYarnScheduleHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeYarnScheduleHistory")
    
    
    return
}

func NewDescribeYarnScheduleHistoryResponse() (response *DescribeYarnScheduleHistoryResponse) {
    response = &DescribeYarnScheduleHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeYarnScheduleHistory
// 查看yarn资源调度的调度历史。废弃，请使用流程中心查看历史记录。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
func (c *Client) DescribeYarnScheduleHistory(request *DescribeYarnScheduleHistoryRequest) (response *DescribeYarnScheduleHistoryResponse, err error) {
    return c.DescribeYarnScheduleHistoryWithContext(context.Background(), request)
}

// DescribeYarnScheduleHistory
// 查看yarn资源调度的调度历史。废弃，请使用流程中心查看历史记录。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
func (c *Client) DescribeYarnScheduleHistoryWithContext(ctx context.Context, request *DescribeYarnScheduleHistoryRequest) (response *DescribeYarnScheduleHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeYarnScheduleHistoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeYarnScheduleHistory require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeYarnScheduleHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewInquirePriceRenewEmrRequest() (request *InquirePriceRenewEmrRequest) {
    request = &InquirePriceRenewEmrRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "InquirePriceRenewEmr")
    
    
    return
}

func NewInquirePriceRenewEmrResponse() (response *InquirePriceRenewEmrResponse) {
    response = &InquirePriceRenewEmrResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquirePriceRenewEmr
// 集群续费询价。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CHECKQUOTAERR = "InternalError.CheckQuotaErr"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDRESOURCEIDS = "InvalidParameter.InvalidResourceIds"
//  INVALIDPARAMETER_INVALIDTIMESPAN = "InvalidParameter.InvalidTimeSpan"
//  INVALIDPARAMETER_INVALIDTIMEUNIT = "InvalidParameter.InvalidTimeUnit"
//  INVALIDPARAMETER_PROJECTRESOURCENOTMATCH = "InvalidParameter.ProjectResourceNotMatch"
//  INVALIDPARAMETER_ZONERESOURCENOTMATCH = "InvalidParameter.ZoneResourceNotMatch"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) InquirePriceRenewEmr(request *InquirePriceRenewEmrRequest) (response *InquirePriceRenewEmrResponse, err error) {
    return c.InquirePriceRenewEmrWithContext(context.Background(), request)
}

// InquirePriceRenewEmr
// 集群续费询价。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CHECKQUOTAERR = "InternalError.CheckQuotaErr"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDRESOURCEIDS = "InvalidParameter.InvalidResourceIds"
//  INVALIDPARAMETER_INVALIDTIMESPAN = "InvalidParameter.InvalidTimeSpan"
//  INVALIDPARAMETER_INVALIDTIMEUNIT = "InvalidParameter.InvalidTimeUnit"
//  INVALIDPARAMETER_PROJECTRESOURCENOTMATCH = "InvalidParameter.ProjectResourceNotMatch"
//  INVALIDPARAMETER_ZONERESOURCENOTMATCH = "InvalidParameter.ZoneResourceNotMatch"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) InquirePriceRenewEmrWithContext(ctx context.Context, request *InquirePriceRenewEmrRequest) (response *InquirePriceRenewEmrResponse, err error) {
    if request == nil {
        request = NewInquirePriceRenewEmrRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquirePriceRenewEmr require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquirePriceRenewEmrResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceCreateInstanceRequest() (request *InquiryPriceCreateInstanceRequest) {
    request = &InquiryPriceCreateInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "InquiryPriceCreateInstance")
    
    
    return
}

func NewInquiryPriceCreateInstanceResponse() (response *InquiryPriceCreateInstanceResponse) {
    response = &InquiryPriceCreateInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquiryPriceCreateInstance
// 创建实例询价
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETCAMSERVERFAILED = "FailedOperation.GetCamServerFailed"
//  FAILEDOPERATION_GETTRADESERVERFAILED = "FailedOperation.GetTradeServerFailed"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_HALESSMASTERCOUNT = "InvalidParameter.HALessMasterCount"
//  INVALIDPARAMETER_INCORRECTCOMMONCOUNT = "InvalidParameter.IncorrectCommonCount"
//  INVALIDPARAMETER_INCORRECTMASTERCOUNT = "InvalidParameter.IncorrectMasterCount"
//  INVALIDPARAMETER_INVALIDCOMMONDISKTYPE = "InvalidParameter.InvalidCommonDiskType"
//  INVALIDPARAMETER_INVALIDCORECOUNT = "InvalidParameter.InvalidCoreCount"
//  INVALIDPARAMETER_INVALIDINSTANCETYPE = "InvalidParameter.InvalidInstanceType"
//  INVALIDPARAMETER_INVALIDMASTERDISKTYPE = "InvalidParameter.InvalidMasterDiskType"
//  INVALIDPARAMETER_INVALIDRESOURCESPEC = "InvalidParameter.InvalidResourceSpec"
//  INVALIDPARAMETER_INVALIDSOFTWARENAME = "InvalidParameter.InvalidSoftWareName"
//  INVALIDPARAMETER_INVALIDTIMEUNIT = "InvalidParameter.InvalidTimeUnit"
//  INVALIDPARAMETER_INVALIDUNIFYMETA = "InvalidParameter.InvalidUnifyMeta"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  INVALIDPARAMETER_NOTCONTAINMUSTSELECTSOFTWARE = "InvalidParameter.NotContainMustSelectSoftware"
//  INVALIDPARAMETER_SOFTWARENOTINPRODUCT = "InvalidParameter.SoftwareNotInProduct"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT_DISKINSUFFICIENT = "ResourceInsufficient.DiskInsufficient"
//  RESOURCEINSUFFICIENT_INSTANCEINSUFFICIENT = "ResourceInsufficient.InstanceInsufficient"
//  RESOURCENOTFOUND_SUBNETNOTFOUND = "ResourceNotFound.SubnetNotFound"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  RESOURCESSOLDOUT_CBSSOLDOUT = "ResourcesSoldOut.CbsSoldOut"
//  RESOURCESSOLDOUT_CVMSOLDOUT = "ResourcesSoldOut.CvmSoldOut"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) InquiryPriceCreateInstance(request *InquiryPriceCreateInstanceRequest) (response *InquiryPriceCreateInstanceResponse, err error) {
    return c.InquiryPriceCreateInstanceWithContext(context.Background(), request)
}

// InquiryPriceCreateInstance
// 创建实例询价
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETCAMSERVERFAILED = "FailedOperation.GetCamServerFailed"
//  FAILEDOPERATION_GETTRADESERVERFAILED = "FailedOperation.GetTradeServerFailed"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_HALESSMASTERCOUNT = "InvalidParameter.HALessMasterCount"
//  INVALIDPARAMETER_INCORRECTCOMMONCOUNT = "InvalidParameter.IncorrectCommonCount"
//  INVALIDPARAMETER_INCORRECTMASTERCOUNT = "InvalidParameter.IncorrectMasterCount"
//  INVALIDPARAMETER_INVALIDCOMMONDISKTYPE = "InvalidParameter.InvalidCommonDiskType"
//  INVALIDPARAMETER_INVALIDCORECOUNT = "InvalidParameter.InvalidCoreCount"
//  INVALIDPARAMETER_INVALIDINSTANCETYPE = "InvalidParameter.InvalidInstanceType"
//  INVALIDPARAMETER_INVALIDMASTERDISKTYPE = "InvalidParameter.InvalidMasterDiskType"
//  INVALIDPARAMETER_INVALIDRESOURCESPEC = "InvalidParameter.InvalidResourceSpec"
//  INVALIDPARAMETER_INVALIDSOFTWARENAME = "InvalidParameter.InvalidSoftWareName"
//  INVALIDPARAMETER_INVALIDTIMEUNIT = "InvalidParameter.InvalidTimeUnit"
//  INVALIDPARAMETER_INVALIDUNIFYMETA = "InvalidParameter.InvalidUnifyMeta"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  INVALIDPARAMETER_NOTCONTAINMUSTSELECTSOFTWARE = "InvalidParameter.NotContainMustSelectSoftware"
//  INVALIDPARAMETER_SOFTWARENOTINPRODUCT = "InvalidParameter.SoftwareNotInProduct"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT_DISKINSUFFICIENT = "ResourceInsufficient.DiskInsufficient"
//  RESOURCEINSUFFICIENT_INSTANCEINSUFFICIENT = "ResourceInsufficient.InstanceInsufficient"
//  RESOURCENOTFOUND_SUBNETNOTFOUND = "ResourceNotFound.SubnetNotFound"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  RESOURCESSOLDOUT_CBSSOLDOUT = "ResourcesSoldOut.CbsSoldOut"
//  RESOURCESSOLDOUT_CVMSOLDOUT = "ResourcesSoldOut.CvmSoldOut"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) InquiryPriceCreateInstanceWithContext(ctx context.Context, request *InquiryPriceCreateInstanceRequest) (response *InquiryPriceCreateInstanceResponse, err error) {
    if request == nil {
        request = NewInquiryPriceCreateInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquiryPriceCreateInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquiryPriceCreateInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceRenewInstanceRequest() (request *InquiryPriceRenewInstanceRequest) {
    request = &InquiryPriceRenewInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "InquiryPriceRenewInstance")
    
    
    return
}

func NewInquiryPriceRenewInstanceResponse() (response *InquiryPriceRenewInstanceResponse) {
    response = &InquiryPriceRenewInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquiryPriceRenewInstance
// 续费询价。
//
// 可能返回的错误码:
//  FAILEDOPERATION_GETCVMSERVERFAILED = "FailedOperation.GetCvmServerFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CHECKQUOTAERR = "InternalError.CheckQuotaErr"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER_INVALIDRESOURCEIDS = "InvalidParameter.InvalidResourceIds"
//  INVALIDPARAMETER_INVALIDRESOURCESPEC = "InvalidParameter.InvalidResourceSpec"
//  INVALIDPARAMETER_INVALIDTIMESPAN = "InvalidParameter.InvalidTimeSpan"
//  INVALIDPARAMETER_INVALIDTIMEUNIT = "InvalidParameter.InvalidTimeUnit"
//  INVALIDPARAMETER_PAYMODERESOURCENOTMATCH = "InvalidParameter.PayModeResourceNotMatch"
//  INVALIDPARAMETER_PROJECTRESOURCENOTMATCH = "InvalidParameter.ProjectResourceNotMatch"
//  INVALIDPARAMETER_ZONERESOURCENOTMATCH = "InvalidParameter.ZoneResourceNotMatch"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) InquiryPriceRenewInstance(request *InquiryPriceRenewInstanceRequest) (response *InquiryPriceRenewInstanceResponse, err error) {
    return c.InquiryPriceRenewInstanceWithContext(context.Background(), request)
}

// InquiryPriceRenewInstance
// 续费询价。
//
// 可能返回的错误码:
//  FAILEDOPERATION_GETCVMSERVERFAILED = "FailedOperation.GetCvmServerFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CHECKQUOTAERR = "InternalError.CheckQuotaErr"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER_INVALIDRESOURCEIDS = "InvalidParameter.InvalidResourceIds"
//  INVALIDPARAMETER_INVALIDRESOURCESPEC = "InvalidParameter.InvalidResourceSpec"
//  INVALIDPARAMETER_INVALIDTIMESPAN = "InvalidParameter.InvalidTimeSpan"
//  INVALIDPARAMETER_INVALIDTIMEUNIT = "InvalidParameter.InvalidTimeUnit"
//  INVALIDPARAMETER_PAYMODERESOURCENOTMATCH = "InvalidParameter.PayModeResourceNotMatch"
//  INVALIDPARAMETER_PROJECTRESOURCENOTMATCH = "InvalidParameter.ProjectResourceNotMatch"
//  INVALIDPARAMETER_ZONERESOURCENOTMATCH = "InvalidParameter.ZoneResourceNotMatch"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) InquiryPriceRenewInstanceWithContext(ctx context.Context, request *InquiryPriceRenewInstanceRequest) (response *InquiryPriceRenewInstanceResponse, err error) {
    if request == nil {
        request = NewInquiryPriceRenewInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquiryPriceRenewInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquiryPriceRenewInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceScaleOutInstanceRequest() (request *InquiryPriceScaleOutInstanceRequest) {
    request = &InquiryPriceScaleOutInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "InquiryPriceScaleOutInstance")
    
    
    return
}

func NewInquiryPriceScaleOutInstanceResponse() (response *InquiryPriceScaleOutInstanceResponse) {
    response = &InquiryPriceScaleOutInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquiryPriceScaleOutInstance
// 扩容询价. 当扩容时候，请通过该接口查询价格。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CHECKQUOTAERR = "InternalError.CheckQuotaErr"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDCORECOUNT = "InvalidParameter.InvalidCoreCount"
//  INVALIDPARAMETER_INVALIDCOUNTNUM = "InvalidParameter.InvalidCountNum"
//  INVALIDPARAMETER_INVALIDINSTANCETYPE = "InvalidParameter.InvalidInstanceType"
//  INVALIDPARAMETER_INVALIDMODIFYSPEC = "InvalidParameter.InvalidModifySpec"
//  INVALIDPARAMETER_INVALIDPAYMODE = "InvalidParameter.InvalidPaymode"
//  INVALIDPARAMETER_INVALIDRESOURCESPEC = "InvalidParameter.InvalidResourceSpec"
//  INVALIDPARAMETER_INVALIDTIMESPAN = "InvalidParameter.InvalidTimeSpan"
//  INVALIDPARAMETER_INVALIDTIMEUNIT = "InvalidParameter.InvalidTimeUnit"
//  INVALIDPARAMETER_INVALIDVPCID = "InvalidParameter.InvalidVpcId"
//  RESOURCEINSUFFICIENT_DISKINSUFFICIENT = "ResourceInsufficient.DiskInsufficient"
//  RESOURCEINSUFFICIENT_INSTANCEINSUFFICIENT = "ResourceInsufficient.InstanceInsufficient"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_RESOURCESPECNOTDEFAULTSPEC = "ResourceUnavailable.ResourceSpecNotDefaultSpec"
//  RESOURCESSOLDOUT_CVMSOLDOUT = "ResourcesSoldOut.CvmSoldOut"
func (c *Client) InquiryPriceScaleOutInstance(request *InquiryPriceScaleOutInstanceRequest) (response *InquiryPriceScaleOutInstanceResponse, err error) {
    return c.InquiryPriceScaleOutInstanceWithContext(context.Background(), request)
}

// InquiryPriceScaleOutInstance
// 扩容询价. 当扩容时候，请通过该接口查询价格。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CHECKQUOTAERR = "InternalError.CheckQuotaErr"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDCORECOUNT = "InvalidParameter.InvalidCoreCount"
//  INVALIDPARAMETER_INVALIDCOUNTNUM = "InvalidParameter.InvalidCountNum"
//  INVALIDPARAMETER_INVALIDINSTANCETYPE = "InvalidParameter.InvalidInstanceType"
//  INVALIDPARAMETER_INVALIDMODIFYSPEC = "InvalidParameter.InvalidModifySpec"
//  INVALIDPARAMETER_INVALIDPAYMODE = "InvalidParameter.InvalidPaymode"
//  INVALIDPARAMETER_INVALIDRESOURCESPEC = "InvalidParameter.InvalidResourceSpec"
//  INVALIDPARAMETER_INVALIDTIMESPAN = "InvalidParameter.InvalidTimeSpan"
//  INVALIDPARAMETER_INVALIDTIMEUNIT = "InvalidParameter.InvalidTimeUnit"
//  INVALIDPARAMETER_INVALIDVPCID = "InvalidParameter.InvalidVpcId"
//  RESOURCEINSUFFICIENT_DISKINSUFFICIENT = "ResourceInsufficient.DiskInsufficient"
//  RESOURCEINSUFFICIENT_INSTANCEINSUFFICIENT = "ResourceInsufficient.InstanceInsufficient"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_RESOURCESPECNOTDEFAULTSPEC = "ResourceUnavailable.ResourceSpecNotDefaultSpec"
//  RESOURCESSOLDOUT_CVMSOLDOUT = "ResourcesSoldOut.CvmSoldOut"
func (c *Client) InquiryPriceScaleOutInstanceWithContext(ctx context.Context, request *InquiryPriceScaleOutInstanceRequest) (response *InquiryPriceScaleOutInstanceResponse, err error) {
    if request == nil {
        request = NewInquiryPriceScaleOutInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquiryPriceScaleOutInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquiryPriceScaleOutInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceUpdateInstanceRequest() (request *InquiryPriceUpdateInstanceRequest) {
    request = &InquiryPriceUpdateInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "InquiryPriceUpdateInstance")
    
    
    return
}

func NewInquiryPriceUpdateInstanceResponse() (response *InquiryPriceUpdateInstanceResponse) {
    response = &InquiryPriceUpdateInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquiryPriceUpdateInstance
// 变配询价
//
// 可能返回的错误码:
//  FAILEDOPERATION_GETTRADESERVERFAILED = "FailedOperation.GetTradeServerFailed"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CHECKQUOTAERR = "InternalError.CheckQuotaErr"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  INVALIDPARAMETER_INVALIDMODIFYSPEC = "InvalidParameter.InvalidModifySpec"
//  INVALIDPARAMETER_INVALIDRESOURCESPEC = "InvalidParameter.InvalidResourceSpec"
//  INVALIDPARAMETER_INVALIDTIMESPAN = "InvalidParameter.InvalidTimeSpan"
//  INVALIDPARAMETER_INVALIDTIMEUNIT = "InvalidParameter.InvalidTimeUnit"
//  INVALIDPARAMETER_INVALIDVPCID = "InvalidParameter.InvalidVpcId"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  INVALIDPARAMETER_ZONERESOURCENOTMATCH = "InvalidParameter.ZoneResourceNotMatch"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) InquiryPriceUpdateInstance(request *InquiryPriceUpdateInstanceRequest) (response *InquiryPriceUpdateInstanceResponse, err error) {
    return c.InquiryPriceUpdateInstanceWithContext(context.Background(), request)
}

// InquiryPriceUpdateInstance
// 变配询价
//
// 可能返回的错误码:
//  FAILEDOPERATION_GETTRADESERVERFAILED = "FailedOperation.GetTradeServerFailed"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CHECKQUOTAERR = "InternalError.CheckQuotaErr"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  INVALIDPARAMETER_INVALIDMODIFYSPEC = "InvalidParameter.InvalidModifySpec"
//  INVALIDPARAMETER_INVALIDRESOURCESPEC = "InvalidParameter.InvalidResourceSpec"
//  INVALIDPARAMETER_INVALIDTIMESPAN = "InvalidParameter.InvalidTimeSpan"
//  INVALIDPARAMETER_INVALIDTIMEUNIT = "InvalidParameter.InvalidTimeUnit"
//  INVALIDPARAMETER_INVALIDVPCID = "InvalidParameter.InvalidVpcId"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  INVALIDPARAMETER_ZONERESOURCENOTMATCH = "InvalidParameter.ZoneResourceNotMatch"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) InquiryPriceUpdateInstanceWithContext(ctx context.Context, request *InquiryPriceUpdateInstanceRequest) (response *InquiryPriceUpdateInstanceResponse, err error) {
    if request == nil {
        request = NewInquiryPriceUpdateInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquiryPriceUpdateInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquiryPriceUpdateInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAutoRenewFlagRequest() (request *ModifyAutoRenewFlagRequest) {
    request = &ModifyAutoRenewFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "ModifyAutoRenewFlag")
    
    
    return
}

func NewModifyAutoRenewFlagResponse() (response *ModifyAutoRenewFlagResponse) {
    response = &ModifyAutoRenewFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAutoRenewFlag
// 前提：预付费集群
//
// 资源级别开启或关闭自动续费
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) ModifyAutoRenewFlag(request *ModifyAutoRenewFlagRequest) (response *ModifyAutoRenewFlagResponse, err error) {
    return c.ModifyAutoRenewFlagWithContext(context.Background(), request)
}

// ModifyAutoRenewFlag
// 前提：预付费集群
//
// 资源级别开启或关闭自动续费
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) ModifyAutoRenewFlagWithContext(ctx context.Context, request *ModifyAutoRenewFlagRequest) (response *ModifyAutoRenewFlagResponse, err error) {
    if request == nil {
        request = NewModifyAutoRenewFlagRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAutoRenewFlag require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAutoRenewFlagResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAutoScaleStrategyRequest() (request *ModifyAutoScaleStrategyRequest) {
    request = &ModifyAutoScaleStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "ModifyAutoScaleStrategy")
    
    
    return
}

func NewModifyAutoScaleStrategyResponse() (response *ModifyAutoScaleStrategyResponse) {
    response = &ModifyAutoScaleStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAutoScaleStrategy
// 修改自动扩缩容规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INVALIDPARAMETER_INVALIDSTATISTICPERIODORTRIGGERTHRESHOLD = "InvalidParameter.InvalidStatisticPeriodOrTriggerThreshold"
//  INVALIDPARAMETER_INVALIDSTRATEGY = "InvalidParameter.InvalidStrategy"
//  INVALIDPARAMETER_INVALIDSTRATEGYPRIORITY = "InvalidParameter.InvalidStrategyPriority"
//  INVALIDPARAMETER_INVALIDSTRATEGYSPEC = "InvalidParameter.InvalidStrategySpec"
//  INVALIDPARAMETER_INVALIDSTRATEGYTYPE = "InvalidParameter.InvalidStrategyType"
//  INVALIDPARAMETER_INVALIDTIMELAYOUT = "InvalidParameter.InvalidTimeLayout"
//  INVALIDPARAMETER_REPEATEDEXECUTIONTIME = "InvalidParameter.RepeatedExecutionTime"
//  INVALIDPARAMETER_REPEATEDSTRATEGYNAME = "InvalidParameter.RepeatedStrategyName"
//  RESOURCENOTFOUND_STRATEGYNOTFOUND = "ResourceNotFound.StrategyNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) ModifyAutoScaleStrategy(request *ModifyAutoScaleStrategyRequest) (response *ModifyAutoScaleStrategyResponse, err error) {
    return c.ModifyAutoScaleStrategyWithContext(context.Background(), request)
}

// ModifyAutoScaleStrategy
// 修改自动扩缩容规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INVALIDPARAMETER_INVALIDSTATISTICPERIODORTRIGGERTHRESHOLD = "InvalidParameter.InvalidStatisticPeriodOrTriggerThreshold"
//  INVALIDPARAMETER_INVALIDSTRATEGY = "InvalidParameter.InvalidStrategy"
//  INVALIDPARAMETER_INVALIDSTRATEGYPRIORITY = "InvalidParameter.InvalidStrategyPriority"
//  INVALIDPARAMETER_INVALIDSTRATEGYSPEC = "InvalidParameter.InvalidStrategySpec"
//  INVALIDPARAMETER_INVALIDSTRATEGYTYPE = "InvalidParameter.InvalidStrategyType"
//  INVALIDPARAMETER_INVALIDTIMELAYOUT = "InvalidParameter.InvalidTimeLayout"
//  INVALIDPARAMETER_REPEATEDEXECUTIONTIME = "InvalidParameter.RepeatedExecutionTime"
//  INVALIDPARAMETER_REPEATEDSTRATEGYNAME = "InvalidParameter.RepeatedStrategyName"
//  RESOURCENOTFOUND_STRATEGYNOTFOUND = "ResourceNotFound.StrategyNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) ModifyAutoScaleStrategyWithContext(ctx context.Context, request *ModifyAutoScaleStrategyRequest) (response *ModifyAutoScaleStrategyResponse, err error) {
    if request == nil {
        request = NewModifyAutoScaleStrategyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAutoScaleStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAutoScaleStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyResourcePoolsRequest() (request *ModifyResourcePoolsRequest) {
    request = &ModifyResourcePoolsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "ModifyResourcePools")
    
    
    return
}

func NewModifyResourcePoolsResponse() (response *ModifyResourcePoolsResponse) {
    response = &ModifyResourcePoolsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyResourcePools
// 刷新YARN的动态资源池。已废弃，请使用`DeployYarnConf`
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) ModifyResourcePools(request *ModifyResourcePoolsRequest) (response *ModifyResourcePoolsResponse, err error) {
    return c.ModifyResourcePoolsWithContext(context.Background(), request)
}

// ModifyResourcePools
// 刷新YARN的动态资源池。已废弃，请使用`DeployYarnConf`
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) ModifyResourcePoolsWithContext(ctx context.Context, request *ModifyResourcePoolsRequest) (response *ModifyResourcePoolsResponse, err error) {
    if request == nil {
        request = NewModifyResourcePoolsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyResourcePools require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyResourcePoolsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyResourceScheduleConfigRequest() (request *ModifyResourceScheduleConfigRequest) {
    request = &ModifyResourceScheduleConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "ModifyResourceScheduleConfig")
    
    
    return
}

func NewModifyResourceScheduleConfigResponse() (response *ModifyResourceScheduleConfigResponse) {
    response = &ModifyResourceScheduleConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyResourceScheduleConfig
// 修改YARN资源调度的资源配置。已废弃，请使用`ModifyYarnQueueV2`来修改队列配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) ModifyResourceScheduleConfig(request *ModifyResourceScheduleConfigRequest) (response *ModifyResourceScheduleConfigResponse, err error) {
    return c.ModifyResourceScheduleConfigWithContext(context.Background(), request)
}

// ModifyResourceScheduleConfig
// 修改YARN资源调度的资源配置。已废弃，请使用`ModifyYarnQueueV2`来修改队列配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) ModifyResourceScheduleConfigWithContext(ctx context.Context, request *ModifyResourceScheduleConfigRequest) (response *ModifyResourceScheduleConfigResponse, err error) {
    if request == nil {
        request = NewModifyResourceScheduleConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyResourceScheduleConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyResourceScheduleConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyResourceSchedulerRequest() (request *ModifyResourceSchedulerRequest) {
    request = &ModifyResourceSchedulerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "ModifyResourceScheduler")
    
    
    return
}

func NewModifyResourceSchedulerResponse() (response *ModifyResourceSchedulerResponse) {
    response = &ModifyResourceSchedulerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyResourceScheduler
// 修改了yarn的资源调度器，点击部署生效。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) ModifyResourceScheduler(request *ModifyResourceSchedulerRequest) (response *ModifyResourceSchedulerResponse, err error) {
    return c.ModifyResourceSchedulerWithContext(context.Background(), request)
}

// ModifyResourceScheduler
// 修改了yarn的资源调度器，点击部署生效。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) ModifyResourceSchedulerWithContext(ctx context.Context, request *ModifyResourceSchedulerRequest) (response *ModifyResourceSchedulerResponse, err error) {
    if request == nil {
        request = NewModifyResourceSchedulerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyResourceScheduler require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyResourceSchedulerResponse()
    err = c.Send(request, response)
    return
}

func NewModifyResourcesTagsRequest() (request *ModifyResourcesTagsRequest) {
    request = &ModifyResourcesTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "ModifyResourcesTags")
    
    
    return
}

func NewModifyResourcesTagsResponse() (response *ModifyResourcesTagsResponse) {
    response = &ModifyResourcesTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyResourcesTags
// 强制修改标签
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyResourcesTags(request *ModifyResourcesTagsRequest) (response *ModifyResourcesTagsResponse, err error) {
    return c.ModifyResourcesTagsWithContext(context.Background(), request)
}

// ModifyResourcesTags
// 强制修改标签
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyResourcesTagsWithContext(ctx context.Context, request *ModifyResourcesTagsRequest) (response *ModifyResourcesTagsResponse, err error) {
    if request == nil {
        request = NewModifyResourcesTagsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyResourcesTags require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyResourcesTagsResponse()
    err = c.Send(request, response)
    return
}

func NewModifySLInstanceRequest() (request *ModifySLInstanceRequest) {
    request = &ModifySLInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "ModifySLInstance")
    
    
    return
}

func NewModifySLInstanceResponse() (response *ModifySLInstanceResponse) {
    response = &ModifySLInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySLInstance
// 本接口（ModifySLInstance）用于修改Lite HBase 实例节点数。
//
// - 接口调用成功，会创建Lite HBase实例，创建实例请求成功会返回请求的 RequestID。
//
// - 接口为异步接口，接口返回时操作并未立即完成，实例操作结果可以通过调用 DescribeInstancesList 查看当前实例的 StatusDesc 状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REFUNDCVMFAILED = "FailedOperation.RefundCvmFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDNODECOUNT = "InvalidParameter.InvalidNodeCount"
//  INVALIDPARAMETER_INVALIDPAYMODE = "InvalidParameter.InvalidPaymode"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_RESOURCESPECNOTEXIST = "ResourceUnavailable.ResourceSpecNotExist"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
//  UNSUPPORTEDOPERATION_SERVICENOTSUPPORT = "UnsupportedOperation.ServiceNotSupport"
func (c *Client) ModifySLInstance(request *ModifySLInstanceRequest) (response *ModifySLInstanceResponse, err error) {
    return c.ModifySLInstanceWithContext(context.Background(), request)
}

// ModifySLInstance
// 本接口（ModifySLInstance）用于修改Lite HBase 实例节点数。
//
// - 接口调用成功，会创建Lite HBase实例，创建实例请求成功会返回请求的 RequestID。
//
// - 接口为异步接口，接口返回时操作并未立即完成，实例操作结果可以通过调用 DescribeInstancesList 查看当前实例的 StatusDesc 状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REFUNDCVMFAILED = "FailedOperation.RefundCvmFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDNODECOUNT = "InvalidParameter.InvalidNodeCount"
//  INVALIDPARAMETER_INVALIDPAYMODE = "InvalidParameter.InvalidPaymode"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_RESOURCESPECNOTEXIST = "ResourceUnavailable.ResourceSpecNotExist"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
//  UNSUPPORTEDOPERATION_SERVICENOTSUPPORT = "UnsupportedOperation.ServiceNotSupport"
func (c *Client) ModifySLInstanceWithContext(ctx context.Context, request *ModifySLInstanceRequest) (response *ModifySLInstanceResponse, err error) {
    if request == nil {
        request = NewModifySLInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySLInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySLInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUserManagerPwdRequest() (request *ModifyUserManagerPwdRequest) {
    request = &ModifyUserManagerPwdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "ModifyUserManagerPwd")
    
    
    return
}

func NewModifyUserManagerPwdResponse() (response *ModifyUserManagerPwdResponse) {
    response = &ModifyUserManagerPwdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyUserManagerPwd
// 修改用户密码（用户管理）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) ModifyUserManagerPwd(request *ModifyUserManagerPwdRequest) (response *ModifyUserManagerPwdResponse, err error) {
    return c.ModifyUserManagerPwdWithContext(context.Background(), request)
}

// ModifyUserManagerPwd
// 修改用户密码（用户管理）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) ModifyUserManagerPwdWithContext(ctx context.Context, request *ModifyUserManagerPwdRequest) (response *ModifyUserManagerPwdResponse, err error) {
    if request == nil {
        request = NewModifyUserManagerPwdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUserManagerPwd require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUserManagerPwdResponse()
    err = c.Send(request, response)
    return
}

func NewModifyYarnDeployRequest() (request *ModifyYarnDeployRequest) {
    request = &ModifyYarnDeployRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "ModifyYarnDeploy")
    
    
    return
}

func NewModifyYarnDeployResponse() (response *ModifyYarnDeployResponse) {
    response = &ModifyYarnDeployResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyYarnDeploy
// 部署生效。已废弃，请使用`DeployYarnConf`接口进行部署生效
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyYarnDeploy(request *ModifyYarnDeployRequest) (response *ModifyYarnDeployResponse, err error) {
    return c.ModifyYarnDeployWithContext(context.Background(), request)
}

// ModifyYarnDeploy
// 部署生效。已废弃，请使用`DeployYarnConf`接口进行部署生效
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyYarnDeployWithContext(ctx context.Context, request *ModifyYarnDeployRequest) (response *ModifyYarnDeployResponse, err error) {
    if request == nil {
        request = NewModifyYarnDeployRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyYarnDeploy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyYarnDeployResponse()
    err = c.Send(request, response)
    return
}

func NewModifyYarnQueueV2Request() (request *ModifyYarnQueueV2Request) {
    request = &ModifyYarnQueueV2Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "ModifyYarnQueueV2")
    
    
    return
}

func NewModifyYarnQueueV2Response() (response *ModifyYarnQueueV2Response) {
    response = &ModifyYarnQueueV2Response{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyYarnQueueV2
// 修改资源调度中队列信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) ModifyYarnQueueV2(request *ModifyYarnQueueV2Request) (response *ModifyYarnQueueV2Response, err error) {
    return c.ModifyYarnQueueV2WithContext(context.Background(), request)
}

// ModifyYarnQueueV2
// 修改资源调度中队列信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) ModifyYarnQueueV2WithContext(ctx context.Context, request *ModifyYarnQueueV2Request) (response *ModifyYarnQueueV2Response, err error) {
    if request == nil {
        request = NewModifyYarnQueueV2Request()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyYarnQueueV2 require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyYarnQueueV2Response()
    err = c.Send(request, response)
    return
}

func NewResetYarnConfigRequest() (request *ResetYarnConfigRequest) {
    request = &ResetYarnConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "ResetYarnConfig")
    
    
    return
}

func NewResetYarnConfigResponse() (response *ResetYarnConfigResponse) {
    response = &ResetYarnConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResetYarnConfig
// 修改YARN资源调度的资源配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) ResetYarnConfig(request *ResetYarnConfigRequest) (response *ResetYarnConfigResponse, err error) {
    return c.ResetYarnConfigWithContext(context.Background(), request)
}

// ResetYarnConfig
// 修改YARN资源调度的资源配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) ResetYarnConfigWithContext(ctx context.Context, request *ResetYarnConfigRequest) (response *ResetYarnConfigResponse, err error) {
    if request == nil {
        request = NewResetYarnConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetYarnConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetYarnConfigResponse()
    err = c.Send(request, response)
    return
}

func NewRunJobFlowRequest() (request *RunJobFlowRequest) {
    request = &RunJobFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "RunJobFlow")
    
    
    return
}

func NewRunJobFlowResponse() (response *RunJobFlowResponse) {
    response = &RunJobFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RunJobFlow
// 创建流程作业
//
// 可能返回的错误码:
//  FAILEDOPERATION_DUPLICATEORDERNOTALLOWED = "FailedOperation.DuplicateOrderNotAllowed"
//  FAILEDOPERATION_GETCVMSERVERFAILED = "FailedOperation.GetCvmServerFailed"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CHECKQUOTAERR = "InternalError.CheckQuotaErr"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INVALIDPARAMETER_INCORRECTCOMMONCOUNT = "InvalidParameter.IncorrectCommonCount"
//  INVALIDPARAMETER_INCORRECTMASTERCOUNT = "InvalidParameter.IncorrectMasterCount"
//  INVALIDPARAMETER_INVALIDBOOTSTRAPACTION = "InvalidParameter.InvalidBootstrapAction"
//  INVALIDPARAMETER_INVALIDCORECOUNT = "InvalidParameter.InvalidCoreCount"
//  INVALIDPARAMETER_INVALIDCOSFILEURI = "InvalidParameter.InvalidCosFileURI"
//  INVALIDPARAMETER_INVALIDDISKNUM = "InvalidParameter.InvalidDiskNum"
//  INVALIDPARAMETER_INVALIDFAILUREPOLICY = "InvalidParameter.InvalidFailurePolicy"
//  INVALIDPARAMETER_INVALIDINSTANCE = "InvalidParameter.InvalidInstance"
//  INVALIDPARAMETER_INVALIDINSTANCEPOLICY = "InvalidParameter.InvalidInstancePolicy"
//  INVALIDPARAMETER_INVALIDINSTANCETYPE = "InvalidParameter.InvalidInstanceType"
//  INVALIDPARAMETER_INVALIDJOBTYPE = "InvalidParameter.InvalidJobType"
//  INVALIDPARAMETER_INVALIDMETATYPE = "InvalidParameter.InvalidMetaType"
//  INVALIDPARAMETER_INVALIDPRODUCT = "InvalidParameter.InvalidProduct"
//  INVALIDPARAMETER_INVALIDPRODUCTID = "InvalidParameter.InvalidProductId"
//  INVALIDPARAMETER_INVALIDPRODUCTVERSION = "InvalidParameter.InvalidProductVersion"
//  INVALIDPARAMETER_INVALIDSOFTINFO = "InvalidParameter.InvalidSoftInfo"
//  INVALIDPARAMETER_INVALIDSOFTWARE = "InvalidParameter.InvalidSoftWare"
//  INVALIDPARAMETER_INVALIDTAGSGROUP = "InvalidParameter.InvalidTagsGroup"
//  INVALIDPARAMETER_INVALIDTIMESPAN = "InvalidParameter.InvalidTimeSpan"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  INVALIDPARAMETER_SOFTWARENOTINPRODUCT = "InvalidParameter.SoftwareNotInProduct"
//  LIMITEXCEEDED_BOOTSTRAPACTIONSNUMLIMITEXCEEDED = "LimitExceeded.BootstrapActionsNumLimitExceeded"
//  RESOURCEINSUFFICIENT_DISKINSUFFICIENT = "ResourceInsufficient.DiskInsufficient"
//  RESOURCEINSUFFICIENT_INSTANCEINSUFFICIENT = "ResourceInsufficient.InstanceInsufficient"
//  RESOURCENOTFOUND_CDBINFONOTFOUND = "ResourceNotFound.CDBInfoNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_SUBNETNOTFOUND = "ResourceNotFound.SubnetNotFound"
//  RESOURCENOTFOUND_TAGSNOTFOUND = "ResourceNotFound.TagsNotFound"
//  RESOURCESSOLDOUT_CVMSOLDOUT = "ResourcesSoldOut.CvmSoldOut"
//  UNSUPPORTEDOPERATION_SERVICENOTSUPPORT = "UnsupportedOperation.ServiceNotSupport"
func (c *Client) RunJobFlow(request *RunJobFlowRequest) (response *RunJobFlowResponse, err error) {
    return c.RunJobFlowWithContext(context.Background(), request)
}

// RunJobFlow
// 创建流程作业
//
// 可能返回的错误码:
//  FAILEDOPERATION_DUPLICATEORDERNOTALLOWED = "FailedOperation.DuplicateOrderNotAllowed"
//  FAILEDOPERATION_GETCVMSERVERFAILED = "FailedOperation.GetCvmServerFailed"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CHECKQUOTAERR = "InternalError.CheckQuotaErr"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INVALIDPARAMETER_INCORRECTCOMMONCOUNT = "InvalidParameter.IncorrectCommonCount"
//  INVALIDPARAMETER_INCORRECTMASTERCOUNT = "InvalidParameter.IncorrectMasterCount"
//  INVALIDPARAMETER_INVALIDBOOTSTRAPACTION = "InvalidParameter.InvalidBootstrapAction"
//  INVALIDPARAMETER_INVALIDCORECOUNT = "InvalidParameter.InvalidCoreCount"
//  INVALIDPARAMETER_INVALIDCOSFILEURI = "InvalidParameter.InvalidCosFileURI"
//  INVALIDPARAMETER_INVALIDDISKNUM = "InvalidParameter.InvalidDiskNum"
//  INVALIDPARAMETER_INVALIDFAILUREPOLICY = "InvalidParameter.InvalidFailurePolicy"
//  INVALIDPARAMETER_INVALIDINSTANCE = "InvalidParameter.InvalidInstance"
//  INVALIDPARAMETER_INVALIDINSTANCEPOLICY = "InvalidParameter.InvalidInstancePolicy"
//  INVALIDPARAMETER_INVALIDINSTANCETYPE = "InvalidParameter.InvalidInstanceType"
//  INVALIDPARAMETER_INVALIDJOBTYPE = "InvalidParameter.InvalidJobType"
//  INVALIDPARAMETER_INVALIDMETATYPE = "InvalidParameter.InvalidMetaType"
//  INVALIDPARAMETER_INVALIDPRODUCT = "InvalidParameter.InvalidProduct"
//  INVALIDPARAMETER_INVALIDPRODUCTID = "InvalidParameter.InvalidProductId"
//  INVALIDPARAMETER_INVALIDPRODUCTVERSION = "InvalidParameter.InvalidProductVersion"
//  INVALIDPARAMETER_INVALIDSOFTINFO = "InvalidParameter.InvalidSoftInfo"
//  INVALIDPARAMETER_INVALIDSOFTWARE = "InvalidParameter.InvalidSoftWare"
//  INVALIDPARAMETER_INVALIDTAGSGROUP = "InvalidParameter.InvalidTagsGroup"
//  INVALIDPARAMETER_INVALIDTIMESPAN = "InvalidParameter.InvalidTimeSpan"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  INVALIDPARAMETER_SOFTWARENOTINPRODUCT = "InvalidParameter.SoftwareNotInProduct"
//  LIMITEXCEEDED_BOOTSTRAPACTIONSNUMLIMITEXCEEDED = "LimitExceeded.BootstrapActionsNumLimitExceeded"
//  RESOURCEINSUFFICIENT_DISKINSUFFICIENT = "ResourceInsufficient.DiskInsufficient"
//  RESOURCEINSUFFICIENT_INSTANCEINSUFFICIENT = "ResourceInsufficient.InstanceInsufficient"
//  RESOURCENOTFOUND_CDBINFONOTFOUND = "ResourceNotFound.CDBInfoNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_SUBNETNOTFOUND = "ResourceNotFound.SubnetNotFound"
//  RESOURCENOTFOUND_TAGSNOTFOUND = "ResourceNotFound.TagsNotFound"
//  RESOURCESSOLDOUT_CVMSOLDOUT = "ResourcesSoldOut.CvmSoldOut"
//  UNSUPPORTEDOPERATION_SERVICENOTSUPPORT = "UnsupportedOperation.ServiceNotSupport"
func (c *Client) RunJobFlowWithContext(ctx context.Context, request *RunJobFlowRequest) (response *RunJobFlowResponse, err error) {
    if request == nil {
        request = NewRunJobFlowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunJobFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewRunJobFlowResponse()
    err = c.Send(request, response)
    return
}

func NewScaleOutClusterRequest() (request *ScaleOutClusterRequest) {
    request = &ScaleOutClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "ScaleOutCluster")
    
    
    return
}

func NewScaleOutClusterResponse() (response *ScaleOutClusterResponse) {
    response = &ScaleOutClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ScaleOutCluster
// 扩容集群节点
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETCVMCONFIGQUOTAFAILED = "FailedOperation.GetCvmConfigQuotaFailed"
//  INVALIDPARAMETER_INVALIDINSTANCECHARGETYPE = "InvalidParameter.InvalidInstanceChargeType"
//  INVALIDPARAMETER_INVALIDINSTANCETYPE = "InvalidParameter.InvalidInstanceType"
//  INVALIDPARAMETER_INVALIDNODEFLAG = "InvalidParameter.InvalidNodeFlag"
//  INVALIDPARAMETER_INVALIDSOFTDEPLOYINFO = "InvalidParameter.InvalidSoftDeployInfo"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  RESOURCEINSUFFICIENT_DISKINSUFFICIENT = "ResourceInsufficient.DiskInsufficient"
//  RESOURCEINSUFFICIENT_INSTANCEINSUFFICIENT = "ResourceInsufficient.InstanceInsufficient"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_RESOURCESPECNOTDEFAULTSPEC = "ResourceUnavailable.ResourceSpecNotDefaultSpec"
func (c *Client) ScaleOutCluster(request *ScaleOutClusterRequest) (response *ScaleOutClusterResponse, err error) {
    return c.ScaleOutClusterWithContext(context.Background(), request)
}

// ScaleOutCluster
// 扩容集群节点
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETCVMCONFIGQUOTAFAILED = "FailedOperation.GetCvmConfigQuotaFailed"
//  INVALIDPARAMETER_INVALIDINSTANCECHARGETYPE = "InvalidParameter.InvalidInstanceChargeType"
//  INVALIDPARAMETER_INVALIDINSTANCETYPE = "InvalidParameter.InvalidInstanceType"
//  INVALIDPARAMETER_INVALIDNODEFLAG = "InvalidParameter.InvalidNodeFlag"
//  INVALIDPARAMETER_INVALIDSOFTDEPLOYINFO = "InvalidParameter.InvalidSoftDeployInfo"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  RESOURCEINSUFFICIENT_DISKINSUFFICIENT = "ResourceInsufficient.DiskInsufficient"
//  RESOURCEINSUFFICIENT_INSTANCEINSUFFICIENT = "ResourceInsufficient.InstanceInsufficient"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_RESOURCESPECNOTDEFAULTSPEC = "ResourceUnavailable.ResourceSpecNotDefaultSpec"
func (c *Client) ScaleOutClusterWithContext(ctx context.Context, request *ScaleOutClusterRequest) (response *ScaleOutClusterResponse, err error) {
    if request == nil {
        request = NewScaleOutClusterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ScaleOutCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewScaleOutClusterResponse()
    err = c.Send(request, response)
    return
}

func NewScaleOutInstanceRequest() (request *ScaleOutInstanceRequest) {
    request = &ScaleOutInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "ScaleOutInstance")
    
    
    return
}

func NewScaleOutInstanceResponse() (response *ScaleOutInstanceResponse) {
    response = &ScaleOutInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ScaleOutInstance
// 扩容节点
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHECKIFSUPPORTPODSTRETCH = "FailedOperation.CheckIfSupportPodStretch"
//  FAILEDOPERATION_DUPLICATEORDERNOTALLOWED = "FailedOperation.DuplicateOrderNotAllowed"
//  FAILEDOPERATION_GETCVMCONFIGQUOTAFAILED = "FailedOperation.GetCvmConfigQuotaFailed"
//  FAILEDOPERATION_GETCVMSERVERFAILED = "FailedOperation.GetCvmServerFailed"
//  FAILEDOPERATION_NOTSUPPORTPOD = "FailedOperation.NotSupportPod"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_EKSERROR = "InternalError.EKSError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TKEERROR = "InternalError.TKEError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPIDRESOURCENOTMATCH = "InvalidParameter.AppIdResourceNotMatch"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDCLICKHOUSECLUSTER = "InvalidParameter.InvalidClickHouseCluster"
//  INVALIDPARAMETER_INVALIDCLIENTTOKEN = "InvalidParameter.InvalidClientToken"
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  INVALIDPARAMETER_INVALIDCORECOUNT = "InvalidParameter.InvalidCoreCount"
//  INVALIDPARAMETER_INVALIDCOUNT = "InvalidParameter.InvalidCount"
//  INVALIDPARAMETER_INVALIDCOUNTNUM = "InvalidParameter.InvalidCountNum"
//  INVALIDPARAMETER_INVALIDCUSTOMIZEDPODPARAM = "InvalidParameter.InvalidCustomizedPodParam"
//  INVALIDPARAMETER_INVALIDEKSINSTANCE = "InvalidParameter.InvalidEksInstance"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDJOBFLOW = "InvalidParameter.InvalidJobFlow"
//  INVALIDPARAMETER_INVALIDPAYMODE = "InvalidParameter.InvalidPaymode"
//  INVALIDPARAMETER_INVALIDRESOURCESPEC = "InvalidParameter.InvalidResourceSpec"
//  INVALIDPARAMETER_INVALIDSERCURITYGRPUPID = "InvalidParameter.InvalidSercurityGrpupId"
//  INVALIDPARAMETER_INVALIDSERVICENODEINFO = "InvalidParameter.InvalidServiceNodeInfo"
//  INVALIDPARAMETER_INVALIDSOFTDEPLOYINFO = "InvalidParameter.InvalidSoftDeployInfo"
//  INVALIDPARAMETER_INVALIDTASKCOUNT = "InvalidParameter.InvalidTaskCount"
//  INVALIDPARAMETER_INVALIDTIMESPAN = "InvalidParameter.InvalidTimeSpan"
//  INVALIDPARAMETER_INVALIDTIMEUNIT = "InvalidParameter.InvalidTimeUnit"
//  INVALIDPARAMETER_INVALIDTKEINSTANCE = "InvalidParameter.InvalidTkeInstance"
//  INVALIDPARAMETERVALUE_INVALIDTKEINSTANCE = "InvalidParameterValue.InvalidTkeInstance"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  RESOURCEINSUFFICIENT_DISKINSUFFICIENT = "ResourceInsufficient.DiskInsufficient"
//  RESOURCEINSUFFICIENT_INSTANCEINSUFFICIENT = "ResourceInsufficient.InstanceInsufficient"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_TKEPRECONDITIONNOTFOUND = "ResourceNotFound.TKEPreconditionNotFound"
//  RESOURCENOTFOUND_TAGSNOTFOUND = "ResourceNotFound.TagsNotFound"
//  RESOURCEUNAVAILABLE_RESOURCESPECNOTDEFAULTSPEC = "ResourceUnavailable.ResourceSpecNotDefaultSpec"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  RESOURCESSOLDOUT_CBSSOLDOUT = "ResourcesSoldOut.CbsSoldOut"
//  RESOURCESSOLDOUT_CVMSOLDOUT = "ResourcesSoldOut.CvmSoldOut"
func (c *Client) ScaleOutInstance(request *ScaleOutInstanceRequest) (response *ScaleOutInstanceResponse, err error) {
    return c.ScaleOutInstanceWithContext(context.Background(), request)
}

// ScaleOutInstance
// 扩容节点
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHECKIFSUPPORTPODSTRETCH = "FailedOperation.CheckIfSupportPodStretch"
//  FAILEDOPERATION_DUPLICATEORDERNOTALLOWED = "FailedOperation.DuplicateOrderNotAllowed"
//  FAILEDOPERATION_GETCVMCONFIGQUOTAFAILED = "FailedOperation.GetCvmConfigQuotaFailed"
//  FAILEDOPERATION_GETCVMSERVERFAILED = "FailedOperation.GetCvmServerFailed"
//  FAILEDOPERATION_NOTSUPPORTPOD = "FailedOperation.NotSupportPod"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_EKSERROR = "InternalError.EKSError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TKEERROR = "InternalError.TKEError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPIDRESOURCENOTMATCH = "InvalidParameter.AppIdResourceNotMatch"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDCLICKHOUSECLUSTER = "InvalidParameter.InvalidClickHouseCluster"
//  INVALIDPARAMETER_INVALIDCLIENTTOKEN = "InvalidParameter.InvalidClientToken"
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  INVALIDPARAMETER_INVALIDCORECOUNT = "InvalidParameter.InvalidCoreCount"
//  INVALIDPARAMETER_INVALIDCOUNT = "InvalidParameter.InvalidCount"
//  INVALIDPARAMETER_INVALIDCOUNTNUM = "InvalidParameter.InvalidCountNum"
//  INVALIDPARAMETER_INVALIDCUSTOMIZEDPODPARAM = "InvalidParameter.InvalidCustomizedPodParam"
//  INVALIDPARAMETER_INVALIDEKSINSTANCE = "InvalidParameter.InvalidEksInstance"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDJOBFLOW = "InvalidParameter.InvalidJobFlow"
//  INVALIDPARAMETER_INVALIDPAYMODE = "InvalidParameter.InvalidPaymode"
//  INVALIDPARAMETER_INVALIDRESOURCESPEC = "InvalidParameter.InvalidResourceSpec"
//  INVALIDPARAMETER_INVALIDSERCURITYGRPUPID = "InvalidParameter.InvalidSercurityGrpupId"
//  INVALIDPARAMETER_INVALIDSERVICENODEINFO = "InvalidParameter.InvalidServiceNodeInfo"
//  INVALIDPARAMETER_INVALIDSOFTDEPLOYINFO = "InvalidParameter.InvalidSoftDeployInfo"
//  INVALIDPARAMETER_INVALIDTASKCOUNT = "InvalidParameter.InvalidTaskCount"
//  INVALIDPARAMETER_INVALIDTIMESPAN = "InvalidParameter.InvalidTimeSpan"
//  INVALIDPARAMETER_INVALIDTIMEUNIT = "InvalidParameter.InvalidTimeUnit"
//  INVALIDPARAMETER_INVALIDTKEINSTANCE = "InvalidParameter.InvalidTkeInstance"
//  INVALIDPARAMETERVALUE_INVALIDTKEINSTANCE = "InvalidParameterValue.InvalidTkeInstance"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  RESOURCEINSUFFICIENT_DISKINSUFFICIENT = "ResourceInsufficient.DiskInsufficient"
//  RESOURCEINSUFFICIENT_INSTANCEINSUFFICIENT = "ResourceInsufficient.InstanceInsufficient"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_TKEPRECONDITIONNOTFOUND = "ResourceNotFound.TKEPreconditionNotFound"
//  RESOURCENOTFOUND_TAGSNOTFOUND = "ResourceNotFound.TagsNotFound"
//  RESOURCEUNAVAILABLE_RESOURCESPECNOTDEFAULTSPEC = "ResourceUnavailable.ResourceSpecNotDefaultSpec"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  RESOURCESSOLDOUT_CBSSOLDOUT = "ResourcesSoldOut.CbsSoldOut"
//  RESOURCESSOLDOUT_CVMSOLDOUT = "ResourcesSoldOut.CvmSoldOut"
func (c *Client) ScaleOutInstanceWithContext(ctx context.Context, request *ScaleOutInstanceRequest) (response *ScaleOutInstanceResponse, err error) {
    if request == nil {
        request = NewScaleOutInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ScaleOutInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewScaleOutInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewStartStopServiceOrMonitorRequest() (request *StartStopServiceOrMonitorRequest) {
    request = &StartStopServiceOrMonitorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "StartStopServiceOrMonitor")
    
    
    return
}

func NewStartStopServiceOrMonitorResponse() (response *StartStopServiceOrMonitorResponse) {
    response = &StartStopServiceOrMonitorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartStopServiceOrMonitor
// 用于启停服务 重启服务等功能
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDJOBFLOW = "InvalidParameter.InvalidJobFlow"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
//  UNSUPPORTEDOPERATION_SERVICENOTSUPPORT = "UnsupportedOperation.ServiceNotSupport"
func (c *Client) StartStopServiceOrMonitor(request *StartStopServiceOrMonitorRequest) (response *StartStopServiceOrMonitorResponse, err error) {
    return c.StartStopServiceOrMonitorWithContext(context.Background(), request)
}

// StartStopServiceOrMonitor
// 用于启停服务 重启服务等功能
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDJOBFLOW = "InvalidParameter.InvalidJobFlow"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
//  UNSUPPORTEDOPERATION_SERVICENOTSUPPORT = "UnsupportedOperation.ServiceNotSupport"
func (c *Client) StartStopServiceOrMonitorWithContext(ctx context.Context, request *StartStopServiceOrMonitorRequest) (response *StartStopServiceOrMonitorResponse, err error) {
    if request == nil {
        request = NewStartStopServiceOrMonitorRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartStopServiceOrMonitor require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartStopServiceOrMonitorResponse()
    err = c.Send(request, response)
    return
}

func NewSyncPodStateRequest() (request *SyncPodStateRequest) {
    request = &SyncPodStateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "SyncPodState")
    
    
    return
}

func NewSyncPodStateResponse() (response *SyncPodStateResponse) {
    response = &SyncPodStateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SyncPodState
// EMR同步TKE中POD状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DUPLICATEORDERNOTALLOWED = "FailedOperation.DuplicateOrderNotAllowed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CHECKQUOTAERR = "InternalError.CheckQuotaErr"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INCORRECTCOMMONCOUNT = "InvalidParameter.IncorrectCommonCount"
//  INVALIDPARAMETER_INCORRECTMASTERCOUNT = "InvalidParameter.IncorrectMasterCount"
//  INVALIDPARAMETER_INVALIDAUTORENEW = "InvalidParameter.InvalidAutoRenew"
//  INVALIDPARAMETER_INVALIDCLIENTTOKEN = "InvalidParameter.InvalidClientToken"
//  INVALIDPARAMETER_INVALIDCOMPONENT = "InvalidParameter.InvalidComponent"
//  INVALIDPARAMETER_INVALIDCORECOUNT = "InvalidParameter.InvalidCoreCount"
//  INVALIDPARAMETER_INVALIDDISKSIZE = "InvalidParameter.InvalidDiskSize"
//  INVALIDPARAMETER_INVALIDEXTENDFIELD = "InvalidParameter.InvalidExtendField"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDLOGINSETTING = "InvalidParameter.InvalidLoginSetting"
//  INVALIDPARAMETER_INVALIDMETATYPE = "InvalidParameter.InvalidMetaType"
//  INVALIDPARAMETER_INVALIDPASSWORD = "InvalidParameter.InvalidPassword"
//  INVALIDPARAMETER_INVALIDPAYMODE = "InvalidParameter.InvalidPaymode"
//  INVALIDPARAMETER_INVALIDPREEXECUTEDFILE = "InvalidParameter.InvalidPreExecutedFile"
//  INVALIDPARAMETER_INVALIDPRODUCTID = "InvalidParameter.InvalidProductId"
//  INVALIDPARAMETER_INVALIDPROJECTID = "InvalidParameter.InvalidProjectId"
//  INVALIDPARAMETER_INVALIDRESOURCESPEC = "InvalidParameter.InvalidResourceSpec"
//  INVALIDPARAMETER_INVALIDSERCURITYGRPUPID = "InvalidParameter.InvalidSercurityGrpupId"
//  INVALIDPARAMETER_INVALIDSERVICENAME = "InvalidParameter.InvalidServiceName"
//  INVALIDPARAMETER_INVALIDSOFTDEPLOYINFO = "InvalidParameter.InvalidSoftDeployInfo"
//  INVALIDPARAMETER_INVALIDSOFTINFO = "InvalidParameter.InvalidSoftInfo"
//  INVALIDPARAMETER_INVALIDSOFTWAREVERSION = "InvalidParameter.InvalidSoftWareVersion"
//  INVALIDPARAMETER_INVALIDSUBNETID = "InvalidParameter.InvalidSubnetId"
//  INVALIDPARAMETER_INVALIDSUPPORTHA = "InvalidParameter.InvalidSupportHA"
//  INVALIDPARAMETER_INVALIDTIMESPAN = "InvalidParameter.InvalidTimeSpan"
//  INVALIDPARAMETER_INVALIDTIMEUNIT = "InvalidParameter.InvalidTimeUnit"
//  INVALIDPARAMETER_INVALIDVPCID = "InvalidParameter.InvalidVpcId"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  INVALIDPARAMETER_NOTCONTAINMUSTSELECTSOFTWARE = "InvalidParameter.NotContainMustSelectSoftware"
//  INVALIDPARAMETER_PAYMODERESOURCENOTMATCH = "InvalidParameter.PayModeResourceNotMatch"
//  INVALIDPARAMETER_SOFTWARENOTINPRODUCT = "InvalidParameter.SoftwareNotInProduct"
//  INVALIDPARAMETER_UNGRANTEDPOLICY = "InvalidParameter.UngrantedPolicy"
//  INVALIDPARAMETER_UNGRANTEDROLE = "InvalidParameter.UngrantedRole"
//  INVALIDPARAMETER_ZONERESOURCENOTMATCH = "InvalidParameter.ZoneResourceNotMatch"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT_DISKINSUFFICIENT = "ResourceInsufficient.DiskInsufficient"
//  RESOURCEINSUFFICIENT_INSTANCEINSUFFICIENT = "ResourceInsufficient.InstanceInsufficient"
//  RESOURCENOTFOUND_TAGSNOTFOUND = "ResourceNotFound.TagsNotFound"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  RESOURCESSOLDOUT_CBSSOLDOUT = "ResourcesSoldOut.CbsSoldOut"
//  RESOURCESSOLDOUT_CVMSOLDOUT = "ResourcesSoldOut.CvmSoldOut"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SyncPodState(request *SyncPodStateRequest) (response *SyncPodStateResponse, err error) {
    return c.SyncPodStateWithContext(context.Background(), request)
}

// SyncPodState
// EMR同步TKE中POD状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DUPLICATEORDERNOTALLOWED = "FailedOperation.DuplicateOrderNotAllowed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CHECKQUOTAERR = "InternalError.CheckQuotaErr"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INCORRECTCOMMONCOUNT = "InvalidParameter.IncorrectCommonCount"
//  INVALIDPARAMETER_INCORRECTMASTERCOUNT = "InvalidParameter.IncorrectMasterCount"
//  INVALIDPARAMETER_INVALIDAUTORENEW = "InvalidParameter.InvalidAutoRenew"
//  INVALIDPARAMETER_INVALIDCLIENTTOKEN = "InvalidParameter.InvalidClientToken"
//  INVALIDPARAMETER_INVALIDCOMPONENT = "InvalidParameter.InvalidComponent"
//  INVALIDPARAMETER_INVALIDCORECOUNT = "InvalidParameter.InvalidCoreCount"
//  INVALIDPARAMETER_INVALIDDISKSIZE = "InvalidParameter.InvalidDiskSize"
//  INVALIDPARAMETER_INVALIDEXTENDFIELD = "InvalidParameter.InvalidExtendField"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDLOGINSETTING = "InvalidParameter.InvalidLoginSetting"
//  INVALIDPARAMETER_INVALIDMETATYPE = "InvalidParameter.InvalidMetaType"
//  INVALIDPARAMETER_INVALIDPASSWORD = "InvalidParameter.InvalidPassword"
//  INVALIDPARAMETER_INVALIDPAYMODE = "InvalidParameter.InvalidPaymode"
//  INVALIDPARAMETER_INVALIDPREEXECUTEDFILE = "InvalidParameter.InvalidPreExecutedFile"
//  INVALIDPARAMETER_INVALIDPRODUCTID = "InvalidParameter.InvalidProductId"
//  INVALIDPARAMETER_INVALIDPROJECTID = "InvalidParameter.InvalidProjectId"
//  INVALIDPARAMETER_INVALIDRESOURCESPEC = "InvalidParameter.InvalidResourceSpec"
//  INVALIDPARAMETER_INVALIDSERCURITYGRPUPID = "InvalidParameter.InvalidSercurityGrpupId"
//  INVALIDPARAMETER_INVALIDSERVICENAME = "InvalidParameter.InvalidServiceName"
//  INVALIDPARAMETER_INVALIDSOFTDEPLOYINFO = "InvalidParameter.InvalidSoftDeployInfo"
//  INVALIDPARAMETER_INVALIDSOFTINFO = "InvalidParameter.InvalidSoftInfo"
//  INVALIDPARAMETER_INVALIDSOFTWAREVERSION = "InvalidParameter.InvalidSoftWareVersion"
//  INVALIDPARAMETER_INVALIDSUBNETID = "InvalidParameter.InvalidSubnetId"
//  INVALIDPARAMETER_INVALIDSUPPORTHA = "InvalidParameter.InvalidSupportHA"
//  INVALIDPARAMETER_INVALIDTIMESPAN = "InvalidParameter.InvalidTimeSpan"
//  INVALIDPARAMETER_INVALIDTIMEUNIT = "InvalidParameter.InvalidTimeUnit"
//  INVALIDPARAMETER_INVALIDVPCID = "InvalidParameter.InvalidVpcId"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  INVALIDPARAMETER_NOTCONTAINMUSTSELECTSOFTWARE = "InvalidParameter.NotContainMustSelectSoftware"
//  INVALIDPARAMETER_PAYMODERESOURCENOTMATCH = "InvalidParameter.PayModeResourceNotMatch"
//  INVALIDPARAMETER_SOFTWARENOTINPRODUCT = "InvalidParameter.SoftwareNotInProduct"
//  INVALIDPARAMETER_UNGRANTEDPOLICY = "InvalidParameter.UngrantedPolicy"
//  INVALIDPARAMETER_UNGRANTEDROLE = "InvalidParameter.UngrantedRole"
//  INVALIDPARAMETER_ZONERESOURCENOTMATCH = "InvalidParameter.ZoneResourceNotMatch"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT_DISKINSUFFICIENT = "ResourceInsufficient.DiskInsufficient"
//  RESOURCEINSUFFICIENT_INSTANCEINSUFFICIENT = "ResourceInsufficient.InstanceInsufficient"
//  RESOURCENOTFOUND_TAGSNOTFOUND = "ResourceNotFound.TagsNotFound"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  RESOURCESSOLDOUT_CBSSOLDOUT = "ResourcesSoldOut.CbsSoldOut"
//  RESOURCESSOLDOUT_CVMSOLDOUT = "ResourcesSoldOut.CvmSoldOut"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SyncPodStateWithContext(ctx context.Context, request *SyncPodStateRequest) (response *SyncPodStateResponse, err error) {
    if request == nil {
        request = NewSyncPodStateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SyncPodState require credential")
    }

    request.SetContext(ctx)
    
    response = NewSyncPodStateResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateClusterNodesRequest() (request *TerminateClusterNodesRequest) {
    request = &TerminateClusterNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "TerminateClusterNodes")
    
    
    return
}

func NewTerminateClusterNodesResponse() (response *TerminateClusterNodesResponse) {
    response = &TerminateClusterNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TerminateClusterNodes
// 销毁集群节点
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDNODEFLAG = "InvalidParameter.InvalidNodeFlag"
//  INVALIDPARAMETER_INVALIDRESOURCEIDS = "InvalidParameter.InvalidResourceIds"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  RESOURCENOTFOUND_CVMINSTANCENOTFOUND = "ResourceNotFound.CvmInstanceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) TerminateClusterNodes(request *TerminateClusterNodesRequest) (response *TerminateClusterNodesResponse, err error) {
    return c.TerminateClusterNodesWithContext(context.Background(), request)
}

// TerminateClusterNodes
// 销毁集群节点
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDNODEFLAG = "InvalidParameter.InvalidNodeFlag"
//  INVALIDPARAMETER_INVALIDRESOURCEIDS = "InvalidParameter.InvalidResourceIds"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  RESOURCENOTFOUND_CVMINSTANCENOTFOUND = "ResourceNotFound.CvmInstanceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) TerminateClusterNodesWithContext(ctx context.Context, request *TerminateClusterNodesRequest) (response *TerminateClusterNodesResponse, err error) {
    if request == nil {
        request = NewTerminateClusterNodesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminateClusterNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewTerminateClusterNodesResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateInstanceRequest() (request *TerminateInstanceRequest) {
    request = &TerminateInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "TerminateInstance")
    
    
    return
}

func NewTerminateInstanceResponse() (response *TerminateInstanceResponse) {
    response = &TerminateInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TerminateInstance
// 销毁EMR实例。此接口仅支持弹性MapReduce正式计费版本。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REFUNDCVMFAILED = "FailedOperation.RefundCvmFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
//  UNSUPPORTEDOPERATION_SERVICENOTSUPPORT = "UnsupportedOperation.ServiceNotSupport"
func (c *Client) TerminateInstance(request *TerminateInstanceRequest) (response *TerminateInstanceResponse, err error) {
    return c.TerminateInstanceWithContext(context.Background(), request)
}

// TerminateInstance
// 销毁EMR实例。此接口仅支持弹性MapReduce正式计费版本。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REFUNDCVMFAILED = "FailedOperation.RefundCvmFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
//  UNSUPPORTEDOPERATION_SERVICENOTSUPPORT = "UnsupportedOperation.ServiceNotSupport"
func (c *Client) TerminateInstanceWithContext(ctx context.Context, request *TerminateInstanceRequest) (response *TerminateInstanceResponse, err error) {
    if request == nil {
        request = NewTerminateInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminateInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewTerminateInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateSLInstanceRequest() (request *TerminateSLInstanceRequest) {
    request = &TerminateSLInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "TerminateSLInstance")
    
    
    return
}

func NewTerminateSLInstanceResponse() (response *TerminateSLInstanceResponse) {
    response = &TerminateSLInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TerminateSLInstance
// 本接口（TerminateSLInstance）用于销毁 Lite HBase 实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REFUNDCVMFAILED = "FailedOperation.RefundCvmFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
//  UNSUPPORTEDOPERATION_SERVICENOTSUPPORT = "UnsupportedOperation.ServiceNotSupport"
func (c *Client) TerminateSLInstance(request *TerminateSLInstanceRequest) (response *TerminateSLInstanceResponse, err error) {
    return c.TerminateSLInstanceWithContext(context.Background(), request)
}

// TerminateSLInstance
// 本接口（TerminateSLInstance）用于销毁 Lite HBase 实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REFUNDCVMFAILED = "FailedOperation.RefundCvmFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
//  UNSUPPORTEDOPERATION_SERVICENOTSUPPORT = "UnsupportedOperation.ServiceNotSupport"
func (c *Client) TerminateSLInstanceWithContext(ctx context.Context, request *TerminateSLInstanceRequest) (response *TerminateSLInstanceResponse, err error) {
    if request == nil {
        request = NewTerminateSLInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminateSLInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewTerminateSLInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateTasksRequest() (request *TerminateTasksRequest) {
    request = &TerminateTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "TerminateTasks")
    
    
    return
}

func NewTerminateTasksResponse() (response *TerminateTasksResponse) {
    response = &TerminateTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TerminateTasks
// 缩容Task节点
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDJOBFLOW = "InvalidParameter.InvalidJobFlow"
//  INVALIDPARAMETER_INVALIDRESOURCEIDS = "InvalidParameter.InvalidResourceIds"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) TerminateTasks(request *TerminateTasksRequest) (response *TerminateTasksResponse, err error) {
    return c.TerminateTasksWithContext(context.Background(), request)
}

// TerminateTasks
// 缩容Task节点
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDJOBFLOW = "InvalidParameter.InvalidJobFlow"
//  INVALIDPARAMETER_INVALIDRESOURCEIDS = "InvalidParameter.InvalidResourceIds"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) TerminateTasksWithContext(ctx context.Context, request *TerminateTasksRequest) (response *TerminateTasksResponse, err error) {
    if request == nil {
        request = NewTerminateTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminateTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewTerminateTasksResponse()
    err = c.Send(request, response)
    return
}
