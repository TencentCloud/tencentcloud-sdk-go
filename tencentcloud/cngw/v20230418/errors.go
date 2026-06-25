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

package v20230418

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// AuthFailure.SecretIdNotFound
	AUTHFAILURE_SECRETIDNOTFOUND = "AuthFailure.SecretIdNotFound"

	// AuthFailure.UnauthorizedOperation
	AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"

	// DryRun 操作，代表请求将会是成功的，只是多传了 DryRun 参数。
	DRYRUNOPERATION = "DryRunOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// FailedOperation.Cls
	FAILEDOPERATION_CLS = "FailedOperation.Cls"

	// FailedOperation.ClsConfig
	FAILEDOPERATION_CLSCONFIG = "FailedOperation.ClsConfig"

	// FailedOperation.Cos
	FAILEDOPERATION_COS = "FailedOperation.Cos"

	// FailedOperation.FailedOperation
	FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"

	// FailedOperation.InternalError
	FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"

	// FailedOperation.Plugin
	FAILEDOPERATION_PLUGIN = "FailedOperation.Plugin"

	// FailedOperation.Resource
	FAILEDOPERATION_RESOURCE = "FailedOperation.Resource"

	// FailedOperation.Role
	FAILEDOPERATION_ROLE = "FailedOperation.Role"

	// FailedOperation.Vpc
	FAILEDOPERATION_VPC = "FailedOperation.Vpc"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// InternalError.CLBFailure
	INTERNALERROR_CLBFAILURE = "InternalError.CLBFailure"

	// InternalError.CamNoAuth
	INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"

	// InternalError.CreateError
	INTERNALERROR_CREATEERROR = "InternalError.CreateError"

	// InternalError.DecodeError
	INTERNALERROR_DECODEERROR = "InternalError.DecodeError"

	// InternalError.EKSFailure
	INTERNALERROR_EKSFAILURE = "InternalError.EKSFailure"

	// InternalError.GetCredential
	INTERNALERROR_GETCREDENTIAL = "InternalError.GetCredential"

	// InternalError.GetRoleError
	INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"

	// InternalError.GovernanceFailure
	INTERNALERROR_GOVERNANCEFAILURE = "InternalError.GovernanceFailure"

	// InternalError.HttpStatusCodeError
	INTERNALERROR_HTTPSTATUSCODEERROR = "InternalError.HttpStatusCodeError"

	// InternalError.IOError
	INTERNALERROR_IOERROR = "InternalError.IOError"

	// InternalError.InternalError
	INTERNALERROR_INTERNALERROR = "InternalError.InternalError"

	// InternalError.OperationFailed
	INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"

	// InternalError.QueryError
	INTERNALERROR_QUERYERROR = "InternalError.QueryError"

	// InternalError.TKEClusterNotReady
	INTERNALERROR_TKECLUSTERNOTREADY = "InternalError.TKEClusterNotReady"

	// InternalError.TKEFailure
	INTERNALERROR_TKEFAILURE = "InternalError.TKEFailure"

	// InternalError.TagFailure
	INTERNALERROR_TAGFAILURE = "InternalError.TagFailure"

	// InternalError.UnknownError
	INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"

	// InternalError.UpdateError
	INTERNALERROR_UPDATEERROR = "InternalError.UpdateError"

	// InternalError.VPCFailure
	INTERNALERROR_VPCFAILURE = "InternalError.VPCFailure"

	// 无效的过滤器
	INVALIDFILTER = "InvalidFilter"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// InvalidParameter.ParseJsonError
	INVALIDPARAMETER_PARSEJSONERROR = "InvalidParameter.ParseJsonError"

	// 无效参数组合
	INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// InvalidParameterValue.Action
	INVALIDPARAMETERVALUE_ACTION = "InvalidParameterValue.Action"

	// InvalidParameterValue.BadRequestFormat
	INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"

	// InvalidParameterValue.BucketName
	INVALIDPARAMETERVALUE_BUCKETNAME = "InvalidParameterValue.BucketName"

	// InvalidParameterValue.Cos
	INVALIDPARAMETERVALUE_COS = "InvalidParameterValue.Cos"

	// InvalidParameterValue.CreateError
	INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"

	// InvalidParameterValue.Description
	INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"

	// InvalidParameterValue.GatewayId
	INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"

	// InvalidParameterValue.GatewayVersion
	INVALIDPARAMETERVALUE_GATEWAYVERSION = "InvalidParameterValue.GatewayVersion"

	// InvalidParameterValue.InvalidParameterValue
	INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"

	// InvalidParameterValue.KeyNotExist
	INVALIDPARAMETERVALUE_KEYNOTEXIST = "InvalidParameterValue.KeyNotExist"

	// InvalidParameterValue.Name
	INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"

	// InvalidParameterValue.ObjectName
	INVALIDPARAMETERVALUE_OBJECTNAME = "InvalidParameterValue.ObjectName"

	// InvalidParameterValue.OperationFailed
	INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"

	// InvalidParameterValue.Plugin
	INVALIDPARAMETERVALUE_PLUGIN = "InvalidParameterValue.Plugin"

	// InvalidParameterValue.QueryError
	INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"

	// InvalidParameterValue.Region
	INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"

	// InvalidParameterValue.ResourceAlreadyExist
	INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"

	// InvalidParameterValue.Specification
	INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"

	// InvalidParameterValue.SystemParameter
	INVALIDPARAMETERVALUE_SYSTEMPARAMETER = "InvalidParameterValue.SystemParameter"

	// InvalidParameterValue.Type
	INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"

	// InvalidParameterValue.UpdateError
	INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"

	// InvalidParameterValue.ZipFile
	INVALIDPARAMETERVALUE_ZIPFILE = "InvalidParameterValue.ZipFile"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// LimitExceeded.LBDomains
	LIMITEXCEEDED_LBDOMAINS = "LimitExceeded.LBDomains"

	// LimitExceeded.LimitExceeded
	LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// MissingParameter.CreateError
	MISSINGPARAMETER_CREATEERROR = "MissingParameter.CreateError"

	// MissingParameter.MissParameter
	MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"

	// MissingParameter.UpdateError
	MISSINGPARAMETER_UPDATEERROR = "MissingParameter.UpdateError"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// OperationDenied.LimitExceeded
	OPERATIONDENIED_LIMITEXCEEDED = "OperationDenied.LimitExceeded"

	// OperationDenied.OperationDenied
	OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"

	// 地域错误
	REGIONERROR = "RegionError"

	// 请求的次数超过了频率限制。
	REQUESTLIMITEXCEEDED = "RequestLimitExceeded"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// ResourceInUse.GatewayServiceExistGroups
	RESOURCEINUSE_GATEWAYSERVICEEXISTGROUPS = "ResourceInUse.GatewayServiceExistGroups"

	// ResourceInUse.GatewayServiceExistRouter
	RESOURCEINUSE_GATEWAYSERVICEEXISTROUTER = "ResourceInUse.GatewayServiceExistRouter"

	// ResourceInUse.GatewayServiceGroupExistUpstream
	RESOURCEINUSE_GATEWAYSERVICEGROUPEXISTUPSTREAM = "ResourceInUse.GatewayServiceGroupExistUpstream"

	// ResourceInUse.GatewayServiceSourceExistService
	RESOURCEINUSE_GATEWAYSERVICESOURCEEXISTSERVICE = "ResourceInUse.GatewayServiceSourceExistService"

	// ResourceInUse.InstancesExistedInService
	RESOURCEINUSE_INSTANCESEXISTEDINSERVICE = "ResourceInUse.InstancesExistedInService"

	// ResourceInUse.ServicesExistedInNamespace
	RESOURCEINUSE_SERVICESEXISTEDINNAMESPACE = "ResourceInUse.ServicesExistedInNamespace"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// ResourceNotFound.ConfigNotExist
	RESOURCENOTFOUND_CONFIGNOTEXIST = "ResourceNotFound.ConfigNotExist"

	// ResourceNotFound.Forbidden
	RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"

	// ResourceNotFound.InstanceNotFound
	RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"

	// ResourceNotFound.ResourceNotFound
	RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 资源售罄。
	RESOURCESSOLDOUT = "ResourcesSoldOut"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// UnauthorizedOperation.CamNoAuth
	UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"

	// UnauthorizedOperation.CamPassRoleNotExist
	UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"

	// UnauthorizedOperation.ClsNotActivated
	UNAUTHORIZEDOPERATION_CLSNOTACTIVATED = "UnauthorizedOperation.ClsNotActivated"

	// UnauthorizedOperation.Uin
	UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"

	// UnauthorizedOperation.UnauthorizedOperation
	UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// UnsupportedOperation.GovernanceNotSupportVisible
	UNSUPPORTEDOPERATION_GOVERNANCENOTSUPPORTVISIBLE = "UnsupportedOperation.GovernanceNotSupportVisible"
)
