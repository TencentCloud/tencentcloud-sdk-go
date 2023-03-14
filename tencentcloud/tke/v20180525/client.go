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

package v20180525

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-05-25"

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


func NewAcquireClusterAdminRoleRequest() (request *AcquireClusterAdminRoleRequest) {
    request = &AcquireClusterAdminRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "AcquireClusterAdminRole")
    
    
    return
}

func NewAcquireClusterAdminRoleResponse() (response *AcquireClusterAdminRoleResponse) {
    response = &AcquireClusterAdminRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AcquireClusterAdminRole
// 通过此接口，可以获取集群的tke:admin的ClusterRole，即管理员角色，可以用于CAM侧高权限的用户，通过CAM策略给予子账户此接口权限，进而可以通过此接口直接获取到kubernetes集群内的管理员角色。
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_KUBERNETESCREATEOPERATIONERROR = "FailedOperation.KubernetesCreateOperationError"
//  FAILEDOPERATION_KUBERNETESGETOPERATIONERROR = "FailedOperation.KubernetesGetOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_KUBERNETESCLIENTBUILDERROR = "InternalError.KubernetesClientBuildError"
//  INTERNALERROR_KUBERNETESCREATEOPERATIONERROR = "InternalError.KubernetesCreateOperationError"
//  INTERNALERROR_KUBERNETESGETOPERATIONERROR = "InternalError.KubernetesGetOperationError"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNSUPPORTEDOPERATION_NOTINWHITELIST = "UnsupportedOperation.NotInWhitelist"
func (c *Client) AcquireClusterAdminRole(request *AcquireClusterAdminRoleRequest) (response *AcquireClusterAdminRoleResponse, err error) {
    return c.AcquireClusterAdminRoleWithContext(context.Background(), request)
}

// AcquireClusterAdminRole
// 通过此接口，可以获取集群的tke:admin的ClusterRole，即管理员角色，可以用于CAM侧高权限的用户，通过CAM策略给予子账户此接口权限，进而可以通过此接口直接获取到kubernetes集群内的管理员角色。
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_KUBERNETESCREATEOPERATIONERROR = "FailedOperation.KubernetesCreateOperationError"
//  FAILEDOPERATION_KUBERNETESGETOPERATIONERROR = "FailedOperation.KubernetesGetOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_KUBERNETESCLIENTBUILDERROR = "InternalError.KubernetesClientBuildError"
//  INTERNALERROR_KUBERNETESCREATEOPERATIONERROR = "InternalError.KubernetesCreateOperationError"
//  INTERNALERROR_KUBERNETESGETOPERATIONERROR = "InternalError.KubernetesGetOperationError"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNSUPPORTEDOPERATION_NOTINWHITELIST = "UnsupportedOperation.NotInWhitelist"
func (c *Client) AcquireClusterAdminRoleWithContext(ctx context.Context, request *AcquireClusterAdminRoleRequest) (response *AcquireClusterAdminRoleResponse, err error) {
    if request == nil {
        request = NewAcquireClusterAdminRoleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AcquireClusterAdminRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewAcquireClusterAdminRoleResponse()
    err = c.Send(request, response)
    return
}

func NewAddClusterCIDRRequest() (request *AddClusterCIDRRequest) {
    request = &AddClusterCIDRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "AddClusterCIDR")
    
    
    return
}

func NewAddClusterCIDRResponse() (response *AddClusterCIDRResponse) {
    response = &AddClusterCIDRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddClusterCIDR
// 给GR集群增加可用的ClusterCIDR
//
// 可能返回的错误码:
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  INTERNALERROR_KUBECLIENTCREATE = "InternalError.KubeClientCreate"
//  INTERNALERROR_KUBECOMMON = "InternalError.KubeCommon"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CIDRMASKSIZEOUTOFRANGE = "InvalidParameter.CIDRMaskSizeOutOfRange"
//  INVALIDPARAMETER_CIDRCONFLICTWITHOTHERCLUSTER = "InvalidParameter.CidrConflictWithOtherCluster"
//  INVALIDPARAMETER_CIDRCONFLICTWITHOTHERROUTE = "InvalidParameter.CidrConflictWithOtherRoute"
//  INVALIDPARAMETER_CIDRCONFLICTWITHVPCCIDR = "InvalidParameter.CidrConflictWithVpcCidr"
//  INVALIDPARAMETER_CIDRCONFLICTWITHVPCGLOBALROUTE = "InvalidParameter.CidrConflictWithVpcGlobalRoute"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  UNSUPPORTEDOPERATION_CLUSTERNOTSUITADDCLUSTERCIDR = "UnsupportedOperation.ClusterNotSuitAddClusterCIDR"
func (c *Client) AddClusterCIDR(request *AddClusterCIDRRequest) (response *AddClusterCIDRResponse, err error) {
    return c.AddClusterCIDRWithContext(context.Background(), request)
}

// AddClusterCIDR
// 给GR集群增加可用的ClusterCIDR
//
// 可能返回的错误码:
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  INTERNALERROR_KUBECLIENTCREATE = "InternalError.KubeClientCreate"
//  INTERNALERROR_KUBECOMMON = "InternalError.KubeCommon"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CIDRMASKSIZEOUTOFRANGE = "InvalidParameter.CIDRMaskSizeOutOfRange"
//  INVALIDPARAMETER_CIDRCONFLICTWITHOTHERCLUSTER = "InvalidParameter.CidrConflictWithOtherCluster"
//  INVALIDPARAMETER_CIDRCONFLICTWITHOTHERROUTE = "InvalidParameter.CidrConflictWithOtherRoute"
//  INVALIDPARAMETER_CIDRCONFLICTWITHVPCCIDR = "InvalidParameter.CidrConflictWithVpcCidr"
//  INVALIDPARAMETER_CIDRCONFLICTWITHVPCGLOBALROUTE = "InvalidParameter.CidrConflictWithVpcGlobalRoute"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  UNSUPPORTEDOPERATION_CLUSTERNOTSUITADDCLUSTERCIDR = "UnsupportedOperation.ClusterNotSuitAddClusterCIDR"
func (c *Client) AddClusterCIDRWithContext(ctx context.Context, request *AddClusterCIDRRequest) (response *AddClusterCIDRResponse, err error) {
    if request == nil {
        request = NewAddClusterCIDRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddClusterCIDR require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddClusterCIDRResponse()
    err = c.Send(request, response)
    return
}

func NewAddExistedInstancesRequest() (request *AddExistedInstancesRequest) {
    request = &AddExistedInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "AddExistedInstances")
    
    
    return
}

func NewAddExistedInstancesResponse() (response *AddExistedInstancesResponse) {
    response = &AddExistedInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddExistedInstances
// 添加已经存在的实例到集群
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLUSTERNOTFOUND = "FailedOperation.ClusterNotFound"
//  FAILEDOPERATION_CLUSTERSTATE = "FailedOperation.ClusterState"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_CVMCOMMON = "FailedOperation.CvmCommon"
//  FAILEDOPERATION_NETWORKSCALEERROR = "FailedOperation.NetworkScaleError"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  FAILEDOPERATION_QUOTAMAXNODLIMIT = "FailedOperation.QuotaMaxNodLimit"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_CVMCOMMON = "InternalError.CvmCommon"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_QUOTAMAXNODLIMIT = "InternalError.QuotaMaxNodLimit"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) AddExistedInstances(request *AddExistedInstancesRequest) (response *AddExistedInstancesResponse, err error) {
    return c.AddExistedInstancesWithContext(context.Background(), request)
}

// AddExistedInstances
// 添加已经存在的实例到集群
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLUSTERNOTFOUND = "FailedOperation.ClusterNotFound"
//  FAILEDOPERATION_CLUSTERSTATE = "FailedOperation.ClusterState"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_CVMCOMMON = "FailedOperation.CvmCommon"
//  FAILEDOPERATION_NETWORKSCALEERROR = "FailedOperation.NetworkScaleError"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  FAILEDOPERATION_QUOTAMAXNODLIMIT = "FailedOperation.QuotaMaxNodLimit"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_CVMCOMMON = "InternalError.CvmCommon"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_QUOTAMAXNODLIMIT = "InternalError.QuotaMaxNodLimit"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) AddExistedInstancesWithContext(ctx context.Context, request *AddExistedInstancesRequest) (response *AddExistedInstancesResponse, err error) {
    if request == nil {
        request = NewAddExistedInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddExistedInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddExistedInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewAddNodeToNodePoolRequest() (request *AddNodeToNodePoolRequest) {
    request = &AddNodeToNodePoolRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "AddNodeToNodePool")
    
    
    return
}

func NewAddNodeToNodePoolResponse() (response *AddNodeToNodePoolResponse) {
    response = &AddNodeToNodePoolResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddNodeToNodePool
// 将集群内节点移入节点池
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_CVMCOMMON = "FailedOperation.CvmCommon"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) AddNodeToNodePool(request *AddNodeToNodePoolRequest) (response *AddNodeToNodePoolResponse, err error) {
    return c.AddNodeToNodePoolWithContext(context.Background(), request)
}

// AddNodeToNodePool
// 将集群内节点移入节点池
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_CVMCOMMON = "FailedOperation.CvmCommon"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) AddNodeToNodePoolWithContext(ctx context.Context, request *AddNodeToNodePoolRequest) (response *AddNodeToNodePoolResponse, err error) {
    if request == nil {
        request = NewAddNodeToNodePoolRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddNodeToNodePool require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddNodeToNodePoolResponse()
    err = c.Send(request, response)
    return
}

func NewAddVpcCniSubnetsRequest() (request *AddVpcCniSubnetsRequest) {
    request = &AddVpcCniSubnetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "AddVpcCniSubnets")
    
    
    return
}

func NewAddVpcCniSubnetsResponse() (response *AddVpcCniSubnetsResponse) {
    response = &AddVpcCniSubnetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddVpcCniSubnets
// 针对VPC-CNI模式的集群，增加集群容器网络可使用的子网
//
// 可能返回的错误码:
//  FAILEDOPERATION_VPCRECODRNOTFOUND = "FailedOperation.VpcRecodrNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_KUBECOMMON = "InternalError.KubeCommon"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INTERNALERROR_VPCRECODRNOTFOUND = "InternalError.VpcRecodrNotFound"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) AddVpcCniSubnets(request *AddVpcCniSubnetsRequest) (response *AddVpcCniSubnetsResponse, err error) {
    return c.AddVpcCniSubnetsWithContext(context.Background(), request)
}

// AddVpcCniSubnets
// 针对VPC-CNI模式的集群，增加集群容器网络可使用的子网
//
// 可能返回的错误码:
//  FAILEDOPERATION_VPCRECODRNOTFOUND = "FailedOperation.VpcRecodrNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_KUBECOMMON = "InternalError.KubeCommon"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INTERNALERROR_VPCRECODRNOTFOUND = "InternalError.VpcRecodrNotFound"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) AddVpcCniSubnetsWithContext(ctx context.Context, request *AddVpcCniSubnetsRequest) (response *AddVpcCniSubnetsResponse, err error) {
    if request == nil {
        request = NewAddVpcCniSubnetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddVpcCniSubnets require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddVpcCniSubnetsResponse()
    err = c.Send(request, response)
    return
}

func NewCancelClusterReleaseRequest() (request *CancelClusterReleaseRequest) {
    request = &CancelClusterReleaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CancelClusterRelease")
    
    
    return
}

func NewCancelClusterReleaseResponse() (response *CancelClusterReleaseResponse) {
    response = &CancelClusterReleaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CancelClusterRelease
// 在应用市场中取消安装失败的应用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MARKETGETAUTHFAILED = "FailedOperation.MarketGetAuthFailed"
//  FAILEDOPERATION_MARKETRELEASEOPERATION = "FailedOperation.MarketReleaseOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MARKETGETAUTHFAILED = "InternalError.MarketGetAuthFailed"
//  INTERNALERROR_MARKETINTERNALSERVERERROR = "InternalError.MarketInternalServerError"
//  INTERNALERROR_MARKETRELEASEOPERATION = "InternalError.MarketReleaseOperation"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CancelClusterRelease(request *CancelClusterReleaseRequest) (response *CancelClusterReleaseResponse, err error) {
    return c.CancelClusterReleaseWithContext(context.Background(), request)
}

// CancelClusterRelease
// 在应用市场中取消安装失败的应用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MARKETGETAUTHFAILED = "FailedOperation.MarketGetAuthFailed"
//  FAILEDOPERATION_MARKETRELEASEOPERATION = "FailedOperation.MarketReleaseOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MARKETGETAUTHFAILED = "InternalError.MarketGetAuthFailed"
//  INTERNALERROR_MARKETINTERNALSERVERERROR = "InternalError.MarketInternalServerError"
//  INTERNALERROR_MARKETRELEASEOPERATION = "InternalError.MarketReleaseOperation"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CancelClusterReleaseWithContext(ctx context.Context, request *CancelClusterReleaseRequest) (response *CancelClusterReleaseResponse, err error) {
    if request == nil {
        request = NewCancelClusterReleaseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelClusterRelease require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelClusterReleaseResponse()
    err = c.Send(request, response)
    return
}

func NewCheckEdgeClusterCIDRRequest() (request *CheckEdgeClusterCIDRRequest) {
    request = &CheckEdgeClusterCIDRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CheckEdgeClusterCIDR")
    
    
    return
}

func NewCheckEdgeClusterCIDRResponse() (response *CheckEdgeClusterCIDRResponse) {
    response = &CheckEdgeClusterCIDRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckEdgeClusterCIDR
// 检查边缘计算集群的CIDR是否冲突
//
// 可能返回的错误码:
//  FAILEDOPERATION_VPCCOMMON = "FailedOperation.VpcCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CIDRCONFLICTWITHOTHERCLUSTER = "InternalError.CidrConflictWithOtherCluster"
//  INTERNALERROR_CIDRCONFLICTWITHOTHERROUTE = "InternalError.CidrConflictWithOtherRoute"
//  INTERNALERROR_CIDRCONFLICTWITHVPCCIDR = "InternalError.CidrConflictWithVpcCidr"
//  INTERNALERROR_CIDRCONFLICTWITHVPCGLOBALROUTE = "InternalError.CidrConflictWithVpcGlobalRoute"
//  INTERNALERROR_CIDRINVALI = "InternalError.CidrInvali"
//  INTERNALERROR_CIDRMASKSIZEOUTOFRANGE = "InternalError.CidrMaskSizeOutOfRange"
//  INTERNALERROR_CMDTIMEOUT = "InternalError.CmdTimeout"
//  INTERNALERROR_INVALIDPRIVATENETWORKCIDR = "InternalError.InvalidPrivateNetworkCidr"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_VPCCOMMON = "InternalError.VpcCommon"
//  INTERNALERROR_VPCPEERNOTFOUND = "InternalError.VpcPeerNotFound"
//  INTERNALERROR_VPCRECODRNOTFOUND = "InternalError.VpcRecodrNotFound"
//  INTERNALERROR_VSTATIONERROR = "InternalError.VstationError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_ROUTETABLENOTEMPTY = "InvalidParameter.RouteTableNotEmpty"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CheckEdgeClusterCIDR(request *CheckEdgeClusterCIDRRequest) (response *CheckEdgeClusterCIDRResponse, err error) {
    return c.CheckEdgeClusterCIDRWithContext(context.Background(), request)
}

// CheckEdgeClusterCIDR
// 检查边缘计算集群的CIDR是否冲突
//
// 可能返回的错误码:
//  FAILEDOPERATION_VPCCOMMON = "FailedOperation.VpcCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CIDRCONFLICTWITHOTHERCLUSTER = "InternalError.CidrConflictWithOtherCluster"
//  INTERNALERROR_CIDRCONFLICTWITHOTHERROUTE = "InternalError.CidrConflictWithOtherRoute"
//  INTERNALERROR_CIDRCONFLICTWITHVPCCIDR = "InternalError.CidrConflictWithVpcCidr"
//  INTERNALERROR_CIDRCONFLICTWITHVPCGLOBALROUTE = "InternalError.CidrConflictWithVpcGlobalRoute"
//  INTERNALERROR_CIDRINVALI = "InternalError.CidrInvali"
//  INTERNALERROR_CIDRMASKSIZEOUTOFRANGE = "InternalError.CidrMaskSizeOutOfRange"
//  INTERNALERROR_CMDTIMEOUT = "InternalError.CmdTimeout"
//  INTERNALERROR_INVALIDPRIVATENETWORKCIDR = "InternalError.InvalidPrivateNetworkCidr"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_VPCCOMMON = "InternalError.VpcCommon"
//  INTERNALERROR_VPCPEERNOTFOUND = "InternalError.VpcPeerNotFound"
//  INTERNALERROR_VPCRECODRNOTFOUND = "InternalError.VpcRecodrNotFound"
//  INTERNALERROR_VSTATIONERROR = "InternalError.VstationError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_ROUTETABLENOTEMPTY = "InvalidParameter.RouteTableNotEmpty"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CheckEdgeClusterCIDRWithContext(ctx context.Context, request *CheckEdgeClusterCIDRRequest) (response *CheckEdgeClusterCIDRResponse, err error) {
    if request == nil {
        request = NewCheckEdgeClusterCIDRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckEdgeClusterCIDR require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckEdgeClusterCIDRResponse()
    err = c.Send(request, response)
    return
}

func NewCheckInstancesUpgradeAbleRequest() (request *CheckInstancesUpgradeAbleRequest) {
    request = &CheckInstancesUpgradeAbleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CheckInstancesUpgradeAble")
    
    
    return
}

func NewCheckInstancesUpgradeAbleResponse() (response *CheckInstancesUpgradeAbleResponse) {
    response = &CheckInstancesUpgradeAbleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckInstancesUpgradeAble
// 检查给定节点列表中哪些是可升级的
//
// 可能返回的错误码:
//  FAILEDOPERATION_KUBECLIENTCONNECTION = "FailedOperation.KubeClientConnection"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  FAILEDOPERATION_POLICYSERVERCOMMONERROR = "FailedOperation.PolicyServerCommonError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_KUBECLIENTCONNECTION = "InternalError.KubeClientConnection"
//  INTERNALERROR_KUBECOMMON = "InternalError.KubeCommon"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
func (c *Client) CheckInstancesUpgradeAble(request *CheckInstancesUpgradeAbleRequest) (response *CheckInstancesUpgradeAbleResponse, err error) {
    return c.CheckInstancesUpgradeAbleWithContext(context.Background(), request)
}

// CheckInstancesUpgradeAble
// 检查给定节点列表中哪些是可升级的
//
// 可能返回的错误码:
//  FAILEDOPERATION_KUBECLIENTCONNECTION = "FailedOperation.KubeClientConnection"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  FAILEDOPERATION_POLICYSERVERCOMMONERROR = "FailedOperation.PolicyServerCommonError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_KUBECLIENTCONNECTION = "InternalError.KubeClientConnection"
//  INTERNALERROR_KUBECOMMON = "InternalError.KubeCommon"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
func (c *Client) CheckInstancesUpgradeAbleWithContext(ctx context.Context, request *CheckInstancesUpgradeAbleRequest) (response *CheckInstancesUpgradeAbleResponse, err error) {
    if request == nil {
        request = NewCheckInstancesUpgradeAbleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckInstancesUpgradeAble require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckInstancesUpgradeAbleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClusterRequest() (request *CreateClusterRequest) {
    request = &CreateClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CreateCluster")
    
    
    return
}

func NewCreateClusterResponse() (response *CreateClusterResponse) {
    response = &CreateClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCluster
// 创建集群
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTCOMMON = "FailedOperation.AccountCommon"
//  FAILEDOPERATION_ACCOUNTUSERNOTAUTHENTICATED = "FailedOperation.AccountUserNotAuthenticated"
//  FAILEDOPERATION_CAMNOAUTH = "FailedOperation.CamNoAuth"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_CVMCOMMON = "FailedOperation.CvmCommon"
//  FAILEDOPERATION_CVMNUMBERNOTMATCH = "FailedOperation.CvmNumberNotMatch"
//  FAILEDOPERATION_CVMVPCIDNOTMATCH = "FailedOperation.CvmVpcidNotMatch"
//  FAILEDOPERATION_DBRECORDNOTFOUND = "FailedOperation.DbRecordNotFound"
//  FAILEDOPERATION_DFWGETUSGQUOTA = "FailedOperation.DfwGetUSGQuota"
//  FAILEDOPERATION_OSNOTSUPPORT = "FailedOperation.OsNotSupport"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  FAILEDOPERATION_QUOTAMAXCLSLIMIT = "FailedOperation.QuotaMaxClsLimit"
//  FAILEDOPERATION_QUOTAMAXNODLIMIT = "FailedOperation.QuotaMaxNodLimit"
//  FAILEDOPERATION_QUOTAUSGLIMIT = "FailedOperation.QuotaUSGLimit"
//  FAILEDOPERATION_TRADECOMMON = "FailedOperation.TradeCommon"
//  FAILEDOPERATION_VPCCOMMON = "FailedOperation.VpcCommon"
//  FAILEDOPERATION_VPCRECODRNOTFOUND = "FailedOperation.VpcRecodrNotFound"
//  FAILEDOPERATION_WHITELISTUNEXPECTEDERROR = "FailedOperation.WhitelistUnexpectedError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCOMMON = "InternalError.AccountCommon"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_ASCOMMON = "InternalError.AsCommon"
//  INTERNALERROR_CIDRCONFLICTWITHOTHERCLUSTER = "InternalError.CidrConflictWithOtherCluster"
//  INTERNALERROR_CIDRCONFLICTWITHOTHERROUTE = "InternalError.CidrConflictWithOtherRoute"
//  INTERNALERROR_CIDRCONFLICTWITHVPCCIDR = "InternalError.CidrConflictWithVpcCidr"
//  INTERNALERROR_CIDRCONFLICTWITHVPCGLOBALROUTE = "InternalError.CidrConflictWithVpcGlobalRoute"
//  INTERNALERROR_CIDRINVALI = "InternalError.CidrInvali"
//  INTERNALERROR_CIDRMASKSIZEOUTOFRANGE = "InternalError.CidrMaskSizeOutOfRange"
//  INTERNALERROR_COMPONENTCLINETHTTP = "InternalError.ComponentClinetHttp"
//  INTERNALERROR_CREATEMASTERFAILED = "InternalError.CreateMasterFailed"
//  INTERNALERROR_CVMCOMMON = "InternalError.CvmCommon"
//  INTERNALERROR_CVMNUMBERNOTMATCH = "InternalError.CvmNumberNotMatch"
//  INTERNALERROR_CVMSTATUS = "InternalError.CvmStatus"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_DFWGETUSGCOUNT = "InternalError.DfwGetUSGCount"
//  INTERNALERROR_DFWGETUSGQUOTA = "InternalError.DfwGetUSGQuota"
//  INTERNALERROR_INITMASTERFAILED = "InternalError.InitMasterFailed"
//  INTERNALERROR_INVALIDPRIVATENETWORKCIDR = "InternalError.InvalidPrivateNetworkCidr"
//  INTERNALERROR_OSNOTSUPPORT = "InternalError.OsNotSupport"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_PUBLICCLUSTEROPNOTSUPPORT = "InternalError.PublicClusterOpNotSupport"
//  INTERNALERROR_QUOTAMAXCLSLIMIT = "InternalError.QuotaMaxClsLimit"
//  INTERNALERROR_QUOTAMAXNODLIMIT = "InternalError.QuotaMaxNodLimit"
//  INTERNALERROR_QUOTAUSGLIMIT = "InternalError.QuotaUSGLimit"
//  INTERNALERROR_TASKCREATEFAILED = "InternalError.TaskCreateFailed"
//  INTERNALERROR_TRADECOMMON = "InternalError.TradeCommon"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INTERNALERROR_VPCCOMMON = "InternalError.VpcCommon"
//  INTERNALERROR_VPCRECODRNOTFOUND = "InternalError.VpcRecodrNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CIDRMASKSIZEOUTOFRANGE = "InvalidParameter.CIDRMaskSizeOutOfRange"
//  INVALIDPARAMETER_CIDRCONFLICTWITHOTHERCLUSTER = "InvalidParameter.CidrConflictWithOtherCluster"
//  INVALIDPARAMETER_CIDRCONFLICTWITHOTHERROUTE = "InvalidParameter.CidrConflictWithOtherRoute"
//  INVALIDPARAMETER_CIDRCONFLICTWITHVPCCIDR = "InvalidParameter.CidrConflictWithVpcCidr"
//  INVALIDPARAMETER_CIDRCONFLICTWITHVPCGLOBALROUTE = "InvalidParameter.CidrConflictWithVpcGlobalRoute"
//  INVALIDPARAMETER_CIDRINVALI = "InvalidParameter.CidrInvali"
//  INVALIDPARAMETER_CIDRINVALID = "InvalidParameter.CidrInvalid"
//  INVALIDPARAMETER_INVALIDPRIVATENETWORKCIDR = "InvalidParameter.InvalidPrivateNetworkCIDR"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT_SPECIFIEDINSTANCETYPE = "ResourceInsufficient.SpecifiedInstanceType"
func (c *Client) CreateCluster(request *CreateClusterRequest) (response *CreateClusterResponse, err error) {
    return c.CreateClusterWithContext(context.Background(), request)
}

// CreateCluster
// 创建集群
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTCOMMON = "FailedOperation.AccountCommon"
//  FAILEDOPERATION_ACCOUNTUSERNOTAUTHENTICATED = "FailedOperation.AccountUserNotAuthenticated"
//  FAILEDOPERATION_CAMNOAUTH = "FailedOperation.CamNoAuth"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_CVMCOMMON = "FailedOperation.CvmCommon"
//  FAILEDOPERATION_CVMNUMBERNOTMATCH = "FailedOperation.CvmNumberNotMatch"
//  FAILEDOPERATION_CVMVPCIDNOTMATCH = "FailedOperation.CvmVpcidNotMatch"
//  FAILEDOPERATION_DBRECORDNOTFOUND = "FailedOperation.DbRecordNotFound"
//  FAILEDOPERATION_DFWGETUSGQUOTA = "FailedOperation.DfwGetUSGQuota"
//  FAILEDOPERATION_OSNOTSUPPORT = "FailedOperation.OsNotSupport"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  FAILEDOPERATION_QUOTAMAXCLSLIMIT = "FailedOperation.QuotaMaxClsLimit"
//  FAILEDOPERATION_QUOTAMAXNODLIMIT = "FailedOperation.QuotaMaxNodLimit"
//  FAILEDOPERATION_QUOTAUSGLIMIT = "FailedOperation.QuotaUSGLimit"
//  FAILEDOPERATION_TRADECOMMON = "FailedOperation.TradeCommon"
//  FAILEDOPERATION_VPCCOMMON = "FailedOperation.VpcCommon"
//  FAILEDOPERATION_VPCRECODRNOTFOUND = "FailedOperation.VpcRecodrNotFound"
//  FAILEDOPERATION_WHITELISTUNEXPECTEDERROR = "FailedOperation.WhitelistUnexpectedError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCOMMON = "InternalError.AccountCommon"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_ASCOMMON = "InternalError.AsCommon"
//  INTERNALERROR_CIDRCONFLICTWITHOTHERCLUSTER = "InternalError.CidrConflictWithOtherCluster"
//  INTERNALERROR_CIDRCONFLICTWITHOTHERROUTE = "InternalError.CidrConflictWithOtherRoute"
//  INTERNALERROR_CIDRCONFLICTWITHVPCCIDR = "InternalError.CidrConflictWithVpcCidr"
//  INTERNALERROR_CIDRCONFLICTWITHVPCGLOBALROUTE = "InternalError.CidrConflictWithVpcGlobalRoute"
//  INTERNALERROR_CIDRINVALI = "InternalError.CidrInvali"
//  INTERNALERROR_CIDRMASKSIZEOUTOFRANGE = "InternalError.CidrMaskSizeOutOfRange"
//  INTERNALERROR_COMPONENTCLINETHTTP = "InternalError.ComponentClinetHttp"
//  INTERNALERROR_CREATEMASTERFAILED = "InternalError.CreateMasterFailed"
//  INTERNALERROR_CVMCOMMON = "InternalError.CvmCommon"
//  INTERNALERROR_CVMNUMBERNOTMATCH = "InternalError.CvmNumberNotMatch"
//  INTERNALERROR_CVMSTATUS = "InternalError.CvmStatus"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_DFWGETUSGCOUNT = "InternalError.DfwGetUSGCount"
//  INTERNALERROR_DFWGETUSGQUOTA = "InternalError.DfwGetUSGQuota"
//  INTERNALERROR_INITMASTERFAILED = "InternalError.InitMasterFailed"
//  INTERNALERROR_INVALIDPRIVATENETWORKCIDR = "InternalError.InvalidPrivateNetworkCidr"
//  INTERNALERROR_OSNOTSUPPORT = "InternalError.OsNotSupport"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_PUBLICCLUSTEROPNOTSUPPORT = "InternalError.PublicClusterOpNotSupport"
//  INTERNALERROR_QUOTAMAXCLSLIMIT = "InternalError.QuotaMaxClsLimit"
//  INTERNALERROR_QUOTAMAXNODLIMIT = "InternalError.QuotaMaxNodLimit"
//  INTERNALERROR_QUOTAUSGLIMIT = "InternalError.QuotaUSGLimit"
//  INTERNALERROR_TASKCREATEFAILED = "InternalError.TaskCreateFailed"
//  INTERNALERROR_TRADECOMMON = "InternalError.TradeCommon"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INTERNALERROR_VPCCOMMON = "InternalError.VpcCommon"
//  INTERNALERROR_VPCRECODRNOTFOUND = "InternalError.VpcRecodrNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CIDRMASKSIZEOUTOFRANGE = "InvalidParameter.CIDRMaskSizeOutOfRange"
//  INVALIDPARAMETER_CIDRCONFLICTWITHOTHERCLUSTER = "InvalidParameter.CidrConflictWithOtherCluster"
//  INVALIDPARAMETER_CIDRCONFLICTWITHOTHERROUTE = "InvalidParameter.CidrConflictWithOtherRoute"
//  INVALIDPARAMETER_CIDRCONFLICTWITHVPCCIDR = "InvalidParameter.CidrConflictWithVpcCidr"
//  INVALIDPARAMETER_CIDRCONFLICTWITHVPCGLOBALROUTE = "InvalidParameter.CidrConflictWithVpcGlobalRoute"
//  INVALIDPARAMETER_CIDRINVALI = "InvalidParameter.CidrInvali"
//  INVALIDPARAMETER_CIDRINVALID = "InvalidParameter.CidrInvalid"
//  INVALIDPARAMETER_INVALIDPRIVATENETWORKCIDR = "InvalidParameter.InvalidPrivateNetworkCIDR"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT_SPECIFIEDINSTANCETYPE = "ResourceInsufficient.SpecifiedInstanceType"
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

func NewCreateClusterEndpointRequest() (request *CreateClusterEndpointRequest) {
    request = &CreateClusterEndpointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CreateClusterEndpoint")
    
    
    return
}

func NewCreateClusterEndpointResponse() (response *CreateClusterEndpointResponse) {
    response = &CreateClusterEndpointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateClusterEndpoint
// 创建集群访问端口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_KUBECLIENTCONNECTION = "FailedOperation.KubeClientConnection"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_EMPTYCLUSTERNOTSUPPORT = "InternalError.EmptyClusterNotSupport"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateClusterEndpoint(request *CreateClusterEndpointRequest) (response *CreateClusterEndpointResponse, err error) {
    return c.CreateClusterEndpointWithContext(context.Background(), request)
}

// CreateClusterEndpoint
// 创建集群访问端口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_KUBECLIENTCONNECTION = "FailedOperation.KubeClientConnection"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_EMPTYCLUSTERNOTSUPPORT = "InternalError.EmptyClusterNotSupport"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateClusterEndpointWithContext(ctx context.Context, request *CreateClusterEndpointRequest) (response *CreateClusterEndpointResponse, err error) {
    if request == nil {
        request = NewCreateClusterEndpointRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateClusterEndpoint require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateClusterEndpointResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClusterEndpointVipRequest() (request *CreateClusterEndpointVipRequest) {
    request = &CreateClusterEndpointVipRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CreateClusterEndpointVip")
    
    
    return
}

func NewCreateClusterEndpointVipResponse() (response *CreateClusterEndpointVipResponse) {
    response = &CreateClusterEndpointVipResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateClusterEndpointVip
// 创建托管集群外网访问端口（老的方式，仅支持托管集群外网端口）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_CVMCOMMON = "InternalError.CvmCommon"
//  INTERNALERROR_CVMNOTFOUND = "InternalError.CvmNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateClusterEndpointVip(request *CreateClusterEndpointVipRequest) (response *CreateClusterEndpointVipResponse, err error) {
    return c.CreateClusterEndpointVipWithContext(context.Background(), request)
}

// CreateClusterEndpointVip
// 创建托管集群外网访问端口（老的方式，仅支持托管集群外网端口）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_CVMCOMMON = "InternalError.CvmCommon"
//  INTERNALERROR_CVMNOTFOUND = "InternalError.CvmNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateClusterEndpointVipWithContext(ctx context.Context, request *CreateClusterEndpointVipRequest) (response *CreateClusterEndpointVipResponse, err error) {
    if request == nil {
        request = NewCreateClusterEndpointVipRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateClusterEndpointVip require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateClusterEndpointVipResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClusterInstancesRequest() (request *CreateClusterInstancesRequest) {
    request = &CreateClusterInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CreateClusterInstances")
    
    
    return
}

func NewCreateClusterInstancesResponse() (response *CreateClusterInstancesResponse) {
    response = &CreateClusterInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateClusterInstances
// 扩展(新建)集群节点
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCOUNTCOMMON = "FailedOperation.AccountCommon"
//  FAILEDOPERATION_CLUSTERNOTFOUND = "FailedOperation.ClusterNotFound"
//  FAILEDOPERATION_CLUSTERSTATE = "FailedOperation.ClusterState"
//  FAILEDOPERATION_CVMCOMMON = "FailedOperation.CvmCommon"
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  FAILEDOPERATION_NETWORKSCALEERROR = "FailedOperation.NetworkScaleError"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  FAILEDOPERATION_QUOTAMAXNODLIMIT = "FailedOperation.QuotaMaxNodLimit"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCOMMON = "InternalError.AccountCommon"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_COMPONENTCLINETHTTP = "InternalError.ComponentClinetHttp"
//  INTERNALERROR_CVMCOMMON = "InternalError.CvmCommon"
//  INTERNALERROR_CVMNOTFOUND = "InternalError.CvmNotFound"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_IMAGEIDNOTFOUND = "InternalError.ImageIdNotFound"
//  INTERNALERROR_OSNOTSUPPORT = "InternalError.OsNotSupport"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_QUOTAMAXCLSLIMIT = "InternalError.QuotaMaxClsLimit"
//  INTERNALERROR_QUOTAMAXNODLIMIT = "InternalError.QuotaMaxNodLimit"
//  INTERNALERROR_QUOTAMAXRTLIMIT = "InternalError.QuotaMaxRtLimit"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INTERNALERROR_VPCCOMMON = "InternalError.VpcCommon"
//  INTERNALERROR_VPCPEERNOTFOUND = "InternalError.VpcPeerNotFound"
//  INTERNALERROR_VPCRECODRNOTFOUND = "InternalError.VpcRecodrNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT_SPECIFIEDINSTANCETYPE = "ResourceInsufficient.SpecifiedInstanceType"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateClusterInstances(request *CreateClusterInstancesRequest) (response *CreateClusterInstancesResponse, err error) {
    return c.CreateClusterInstancesWithContext(context.Background(), request)
}

// CreateClusterInstances
// 扩展(新建)集群节点
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCOUNTCOMMON = "FailedOperation.AccountCommon"
//  FAILEDOPERATION_CLUSTERNOTFOUND = "FailedOperation.ClusterNotFound"
//  FAILEDOPERATION_CLUSTERSTATE = "FailedOperation.ClusterState"
//  FAILEDOPERATION_CVMCOMMON = "FailedOperation.CvmCommon"
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  FAILEDOPERATION_NETWORKSCALEERROR = "FailedOperation.NetworkScaleError"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  FAILEDOPERATION_QUOTAMAXNODLIMIT = "FailedOperation.QuotaMaxNodLimit"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCOMMON = "InternalError.AccountCommon"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_COMPONENTCLINETHTTP = "InternalError.ComponentClinetHttp"
//  INTERNALERROR_CVMCOMMON = "InternalError.CvmCommon"
//  INTERNALERROR_CVMNOTFOUND = "InternalError.CvmNotFound"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_IMAGEIDNOTFOUND = "InternalError.ImageIdNotFound"
//  INTERNALERROR_OSNOTSUPPORT = "InternalError.OsNotSupport"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_QUOTAMAXCLSLIMIT = "InternalError.QuotaMaxClsLimit"
//  INTERNALERROR_QUOTAMAXNODLIMIT = "InternalError.QuotaMaxNodLimit"
//  INTERNALERROR_QUOTAMAXRTLIMIT = "InternalError.QuotaMaxRtLimit"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INTERNALERROR_VPCCOMMON = "InternalError.VpcCommon"
//  INTERNALERROR_VPCPEERNOTFOUND = "InternalError.VpcPeerNotFound"
//  INTERNALERROR_VPCRECODRNOTFOUND = "InternalError.VpcRecodrNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT_SPECIFIEDINSTANCETYPE = "ResourceInsufficient.SpecifiedInstanceType"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateClusterInstancesWithContext(ctx context.Context, request *CreateClusterInstancesRequest) (response *CreateClusterInstancesResponse, err error) {
    if request == nil {
        request = NewCreateClusterInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateClusterInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateClusterInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClusterNodePoolRequest() (request *CreateClusterNodePoolRequest) {
    request = &CreateClusterNodePoolRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CreateClusterNodePool")
    
    
    return
}

func NewCreateClusterNodePoolResponse() (response *CreateClusterNodePoolResponse) {
    response = &CreateClusterNodePoolResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateClusterNodePool
// 创建节点池
//
// 可能返回的错误码:
//  FAILEDOPERATION_ASCOMMON = "FailedOperation.AsCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_POLICYSERVERCOMMONERROR = "FailedOperation.PolicyServerCommonError"
//  FAILEDOPERATION_UNEXPECTEDERROR = "FailedOperation.UnexpectedError"
//  INTERNALERROR_ASCOMMON = "InternalError.AsCommon"
//  INTERNALERROR_COMPONENTCLIENTHTTP = "InternalError.ComponentClientHttp"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_OSNOTSUPPORT = "InvalidParameter.OsNotSupport"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_ASASGNOTEXIST = "ResourceNotFound.AsAsgNotExist"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) CreateClusterNodePool(request *CreateClusterNodePoolRequest) (response *CreateClusterNodePoolResponse, err error) {
    return c.CreateClusterNodePoolWithContext(context.Background(), request)
}

// CreateClusterNodePool
// 创建节点池
//
// 可能返回的错误码:
//  FAILEDOPERATION_ASCOMMON = "FailedOperation.AsCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_POLICYSERVERCOMMONERROR = "FailedOperation.PolicyServerCommonError"
//  FAILEDOPERATION_UNEXPECTEDERROR = "FailedOperation.UnexpectedError"
//  INTERNALERROR_ASCOMMON = "InternalError.AsCommon"
//  INTERNALERROR_COMPONENTCLIENTHTTP = "InternalError.ComponentClientHttp"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_OSNOTSUPPORT = "InvalidParameter.OsNotSupport"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_ASASGNOTEXIST = "ResourceNotFound.AsAsgNotExist"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) CreateClusterNodePoolWithContext(ctx context.Context, request *CreateClusterNodePoolRequest) (response *CreateClusterNodePoolResponse, err error) {
    if request == nil {
        request = NewCreateClusterNodePoolRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateClusterNodePool require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateClusterNodePoolResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClusterReleaseRequest() (request *CreateClusterReleaseRequest) {
    request = &CreateClusterReleaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CreateClusterRelease")
    
    
    return
}

func NewCreateClusterReleaseResponse() (response *CreateClusterReleaseResponse) {
    response = &CreateClusterReleaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateClusterRelease
// 在应用市场中给集群创建应用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MARKETGETAUTHFAILED = "FailedOperation.MarketGetAuthFailed"
//  FAILEDOPERATION_MARKETRELEASEOPERATION = "FailedOperation.MarketReleaseOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MARKETGETAUTHFAILED = "InternalError.MarketGetAuthFailed"
//  INTERNALERROR_MARKETINTERNALSERVERERROR = "InternalError.MarketInternalServerError"
//  INTERNALERROR_MARKETRELEASEOPERATION = "InternalError.MarketReleaseOperation"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateClusterRelease(request *CreateClusterReleaseRequest) (response *CreateClusterReleaseResponse, err error) {
    return c.CreateClusterReleaseWithContext(context.Background(), request)
}

// CreateClusterRelease
// 在应用市场中给集群创建应用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MARKETGETAUTHFAILED = "FailedOperation.MarketGetAuthFailed"
//  FAILEDOPERATION_MARKETRELEASEOPERATION = "FailedOperation.MarketReleaseOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MARKETGETAUTHFAILED = "InternalError.MarketGetAuthFailed"
//  INTERNALERROR_MARKETINTERNALSERVERERROR = "InternalError.MarketInternalServerError"
//  INTERNALERROR_MARKETRELEASEOPERATION = "InternalError.MarketReleaseOperation"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateClusterReleaseWithContext(ctx context.Context, request *CreateClusterReleaseRequest) (response *CreateClusterReleaseResponse, err error) {
    if request == nil {
        request = NewCreateClusterReleaseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateClusterRelease require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateClusterReleaseResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClusterRouteRequest() (request *CreateClusterRouteRequest) {
    request = &CreateClusterRouteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CreateClusterRoute")
    
    
    return
}

func NewCreateClusterRouteResponse() (response *CreateClusterRouteResponse) {
    response = &CreateClusterRouteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateClusterRoute
// 创建集群路由
//
// 可能返回的错误码:
//  FAILEDOPERATION_VPCCOMMON = "FailedOperation.VpcCommon"
//  INTERNALERROR_CIDRCONFLICTWITHOTHERROUTE = "InternalError.CidrConflictWithOtherRoute"
//  INTERNALERROR_CIDROUTOFROUTETABLE = "InternalError.CidrOutOfRouteTable"
//  INTERNALERROR_CVMNOTFOUND = "InternalError.CvmNotFound"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_GATEWAYALREADYASSOCIATEDCIDR = "InternalError.GatewayAlreadyAssociatedCidr"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_ROUTETABLENOTFOUND = "InternalError.RouteTableNotFound"
//  INTERNALERROR_VPCCOMMON = "InternalError.VpcCommon"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CIDRCONFLICTWITHOTHERROUTE = "InvalidParameter.CidrConflictWithOtherRoute"
//  INVALIDPARAMETER_CIDROUTOFROUTETABLE = "InvalidParameter.CidrOutOfRouteTable"
//  INVALIDPARAMETER_GATEWAYALREADYASSOCIATEDCIDR = "InvalidParameter.GatewayAlreadyAssociatedCidr"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND_ROUTETABLENOTFOUND = "ResourceNotFound.RouteTableNotFound"
func (c *Client) CreateClusterRoute(request *CreateClusterRouteRequest) (response *CreateClusterRouteResponse, err error) {
    return c.CreateClusterRouteWithContext(context.Background(), request)
}

// CreateClusterRoute
// 创建集群路由
//
// 可能返回的错误码:
//  FAILEDOPERATION_VPCCOMMON = "FailedOperation.VpcCommon"
//  INTERNALERROR_CIDRCONFLICTWITHOTHERROUTE = "InternalError.CidrConflictWithOtherRoute"
//  INTERNALERROR_CIDROUTOFROUTETABLE = "InternalError.CidrOutOfRouteTable"
//  INTERNALERROR_CVMNOTFOUND = "InternalError.CvmNotFound"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_GATEWAYALREADYASSOCIATEDCIDR = "InternalError.GatewayAlreadyAssociatedCidr"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_ROUTETABLENOTFOUND = "InternalError.RouteTableNotFound"
//  INTERNALERROR_VPCCOMMON = "InternalError.VpcCommon"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CIDRCONFLICTWITHOTHERROUTE = "InvalidParameter.CidrConflictWithOtherRoute"
//  INVALIDPARAMETER_CIDROUTOFROUTETABLE = "InvalidParameter.CidrOutOfRouteTable"
//  INVALIDPARAMETER_GATEWAYALREADYASSOCIATEDCIDR = "InvalidParameter.GatewayAlreadyAssociatedCidr"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND_ROUTETABLENOTFOUND = "ResourceNotFound.RouteTableNotFound"
func (c *Client) CreateClusterRouteWithContext(ctx context.Context, request *CreateClusterRouteRequest) (response *CreateClusterRouteResponse, err error) {
    if request == nil {
        request = NewCreateClusterRouteRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateClusterRoute require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateClusterRouteResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClusterRouteTableRequest() (request *CreateClusterRouteTableRequest) {
    request = &CreateClusterRouteTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CreateClusterRouteTable")
    
    
    return
}

func NewCreateClusterRouteTableResponse() (response *CreateClusterRouteTableResponse) {
    response = &CreateClusterRouteTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateClusterRouteTable
// 创建集群路由表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CIDRCONFLICTWITHOTHERCLUSTER = "InternalError.CidrConflictWithOtherCluster"
//  INTERNALERROR_CIDRCONFLICTWITHOTHERROUTE = "InternalError.CidrConflictWithOtherRoute"
//  INTERNALERROR_CIDRCONFLICTWITHVPCCIDR = "InternalError.CidrConflictWithVpcCidr"
//  INTERNALERROR_CIDRCONFLICTWITHVPCGLOBALROUTE = "InternalError.CidrConflictWithVpcGlobalRoute"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_QUOTAMAXRTLIMIT = "InternalError.QuotaMaxRtLimit"
//  INTERNALERROR_RESOURCEEXISTALREADY = "InternalError.ResourceExistAlready"
//  INTERNALERROR_VPCCOMMON = "InternalError.VpcCommon"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CIDRCONFLICTWITHOTHERROUTE = "InvalidParameter.CidrConflictWithOtherRoute"
//  INVALIDPARAMETER_CIDRCONFLICTWITHVPCCIDR = "InvalidParameter.CidrConflictWithVpcCidr"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCEINUSE_RESOURCEEXISTALREADY = "ResourceInUse.ResourceExistAlready"
func (c *Client) CreateClusterRouteTable(request *CreateClusterRouteTableRequest) (response *CreateClusterRouteTableResponse, err error) {
    return c.CreateClusterRouteTableWithContext(context.Background(), request)
}

// CreateClusterRouteTable
// 创建集群路由表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CIDRCONFLICTWITHOTHERCLUSTER = "InternalError.CidrConflictWithOtherCluster"
//  INTERNALERROR_CIDRCONFLICTWITHOTHERROUTE = "InternalError.CidrConflictWithOtherRoute"
//  INTERNALERROR_CIDRCONFLICTWITHVPCCIDR = "InternalError.CidrConflictWithVpcCidr"
//  INTERNALERROR_CIDRCONFLICTWITHVPCGLOBALROUTE = "InternalError.CidrConflictWithVpcGlobalRoute"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_QUOTAMAXRTLIMIT = "InternalError.QuotaMaxRtLimit"
//  INTERNALERROR_RESOURCEEXISTALREADY = "InternalError.ResourceExistAlready"
//  INTERNALERROR_VPCCOMMON = "InternalError.VpcCommon"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CIDRCONFLICTWITHOTHERROUTE = "InvalidParameter.CidrConflictWithOtherRoute"
//  INVALIDPARAMETER_CIDRCONFLICTWITHVPCCIDR = "InvalidParameter.CidrConflictWithVpcCidr"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCEINUSE_RESOURCEEXISTALREADY = "ResourceInUse.ResourceExistAlready"
func (c *Client) CreateClusterRouteTableWithContext(ctx context.Context, request *CreateClusterRouteTableRequest) (response *CreateClusterRouteTableResponse, err error) {
    if request == nil {
        request = NewCreateClusterRouteTableRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateClusterRouteTable require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateClusterRouteTableResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClusterVirtualNodeRequest() (request *CreateClusterVirtualNodeRequest) {
    request = &CreateClusterVirtualNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CreateClusterVirtualNode")
    
    
    return
}

func NewCreateClusterVirtualNodeResponse() (response *CreateClusterVirtualNodeResponse) {
    response = &CreateClusterVirtualNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateClusterVirtualNode
// 创建虚拟节点
//
// 可能返回的错误码:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCEINUSE_SUBNETALREADYEXIST = "ResourceInUse.SubnetAlreadyExist"
//  RESOURCEUNAVAILABLE_NODEPOOLSTATENOTNORMAL = "ResourceUnavailable.NodePoolStateNotNormal"
//  UNSUPPORTEDOPERATION_NOTINWHITELIST = "UnsupportedOperation.NotInWhitelist"
//  UNSUPPORTEDOPERATION_NOTSUPPORTINSTALLVIRTUALKUBELET = "UnsupportedOperation.NotSupportInstallVirtualKubelet"
func (c *Client) CreateClusterVirtualNode(request *CreateClusterVirtualNodeRequest) (response *CreateClusterVirtualNodeResponse, err error) {
    return c.CreateClusterVirtualNodeWithContext(context.Background(), request)
}

// CreateClusterVirtualNode
// 创建虚拟节点
//
// 可能返回的错误码:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCEINUSE_SUBNETALREADYEXIST = "ResourceInUse.SubnetAlreadyExist"
//  RESOURCEUNAVAILABLE_NODEPOOLSTATENOTNORMAL = "ResourceUnavailable.NodePoolStateNotNormal"
//  UNSUPPORTEDOPERATION_NOTINWHITELIST = "UnsupportedOperation.NotInWhitelist"
//  UNSUPPORTEDOPERATION_NOTSUPPORTINSTALLVIRTUALKUBELET = "UnsupportedOperation.NotSupportInstallVirtualKubelet"
func (c *Client) CreateClusterVirtualNodeWithContext(ctx context.Context, request *CreateClusterVirtualNodeRequest) (response *CreateClusterVirtualNodeResponse, err error) {
    if request == nil {
        request = NewCreateClusterVirtualNodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateClusterVirtualNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateClusterVirtualNodeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClusterVirtualNodePoolRequest() (request *CreateClusterVirtualNodePoolRequest) {
    request = &CreateClusterVirtualNodePoolRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CreateClusterVirtualNodePool")
    
    
    return
}

func NewCreateClusterVirtualNodePoolResponse() (response *CreateClusterVirtualNodePoolResponse) {
    response = &CreateClusterVirtualNodePoolResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateClusterVirtualNodePool
// 创建虚拟节点池
//
// 可能返回的错误码:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_SUBNETINVALIDERROR = "InvalidParameter.SubnetInvalidError"
//  INVALIDPARAMETER_SUBNETNOTEXIST = "InvalidParameter.SubnetNotExist"
//  RESOURCEINUSE_SUBNETALREADYEXIST = "ResourceInUse.SubnetAlreadyExist"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNSUPPORTEDOPERATION_NOTSUPPORTINSTALLVIRTUALKUBELET = "UnsupportedOperation.NotSupportInstallVirtualKubelet"
func (c *Client) CreateClusterVirtualNodePool(request *CreateClusterVirtualNodePoolRequest) (response *CreateClusterVirtualNodePoolResponse, err error) {
    return c.CreateClusterVirtualNodePoolWithContext(context.Background(), request)
}

// CreateClusterVirtualNodePool
// 创建虚拟节点池
//
// 可能返回的错误码:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_SUBNETINVALIDERROR = "InvalidParameter.SubnetInvalidError"
//  INVALIDPARAMETER_SUBNETNOTEXIST = "InvalidParameter.SubnetNotExist"
//  RESOURCEINUSE_SUBNETALREADYEXIST = "ResourceInUse.SubnetAlreadyExist"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNSUPPORTEDOPERATION_NOTSUPPORTINSTALLVIRTUALKUBELET = "UnsupportedOperation.NotSupportInstallVirtualKubelet"
func (c *Client) CreateClusterVirtualNodePoolWithContext(ctx context.Context, request *CreateClusterVirtualNodePoolRequest) (response *CreateClusterVirtualNodePoolResponse, err error) {
    if request == nil {
        request = NewCreateClusterVirtualNodePoolRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateClusterVirtualNodePool require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateClusterVirtualNodePoolResponse()
    err = c.Send(request, response)
    return
}

func NewCreateECMInstancesRequest() (request *CreateECMInstancesRequest) {
    request = &CreateECMInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CreateECMInstances")
    
    
    return
}

func NewCreateECMInstancesResponse() (response *CreateECMInstancesResponse) {
    response = &CreateECMInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateECMInstances
// 创建边缘计算ECM机器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateECMInstances(request *CreateECMInstancesRequest) (response *CreateECMInstancesResponse, err error) {
    return c.CreateECMInstancesWithContext(context.Background(), request)
}

// CreateECMInstances
// 创建边缘计算ECM机器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateECMInstancesWithContext(ctx context.Context, request *CreateECMInstancesRequest) (response *CreateECMInstancesResponse, err error) {
    if request == nil {
        request = NewCreateECMInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateECMInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateECMInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEKSClusterRequest() (request *CreateEKSClusterRequest) {
    request = &CreateEKSClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CreateEKSCluster")
    
    
    return
}

func NewCreateEKSClusterResponse() (response *CreateEKSClusterResponse) {
    response = &CreateEKSClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateEKSCluster
// 创建弹性集群
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateEKSCluster(request *CreateEKSClusterRequest) (response *CreateEKSClusterResponse, err error) {
    return c.CreateEKSClusterWithContext(context.Background(), request)
}

// CreateEKSCluster
// 创建弹性集群
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateEKSClusterWithContext(ctx context.Context, request *CreateEKSClusterRequest) (response *CreateEKSClusterResponse, err error) {
    if request == nil {
        request = NewCreateEKSClusterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEKSCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEKSClusterResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEKSContainerInstancesRequest() (request *CreateEKSContainerInstancesRequest) {
    request = &CreateEKSContainerInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CreateEKSContainerInstances")
    
    
    return
}

func NewCreateEKSContainerInstancesResponse() (response *CreateEKSContainerInstancesResponse) {
    response = &CreateEKSContainerInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateEKSContainerInstances
// 创建容器实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CMDTIMEOUT = "InternalError.CmdTimeout"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateEKSContainerInstances(request *CreateEKSContainerInstancesRequest) (response *CreateEKSContainerInstancesResponse, err error) {
    return c.CreateEKSContainerInstancesWithContext(context.Background(), request)
}

// CreateEKSContainerInstances
// 创建容器实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CMDTIMEOUT = "InternalError.CmdTimeout"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateEKSContainerInstancesWithContext(ctx context.Context, request *CreateEKSContainerInstancesRequest) (response *CreateEKSContainerInstancesResponse, err error) {
    if request == nil {
        request = NewCreateEKSContainerInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEKSContainerInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEKSContainerInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEdgeCVMInstancesRequest() (request *CreateEdgeCVMInstancesRequest) {
    request = &CreateEdgeCVMInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CreateEdgeCVMInstances")
    
    
    return
}

func NewCreateEdgeCVMInstancesResponse() (response *CreateEdgeCVMInstancesResponse) {
    response = &CreateEdgeCVMInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateEdgeCVMInstances
// 创建边缘容器CVM机器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateEdgeCVMInstances(request *CreateEdgeCVMInstancesRequest) (response *CreateEdgeCVMInstancesResponse, err error) {
    return c.CreateEdgeCVMInstancesWithContext(context.Background(), request)
}

// CreateEdgeCVMInstances
// 创建边缘容器CVM机器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateEdgeCVMInstancesWithContext(ctx context.Context, request *CreateEdgeCVMInstancesRequest) (response *CreateEdgeCVMInstancesResponse, err error) {
    if request == nil {
        request = NewCreateEdgeCVMInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEdgeCVMInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEdgeCVMInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEdgeLogConfigRequest() (request *CreateEdgeLogConfigRequest) {
    request = &CreateEdgeLogConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CreateEdgeLogConfig")
    
    
    return
}

func NewCreateEdgeLogConfigResponse() (response *CreateEdgeLogConfigResponse) {
    response = &CreateEdgeLogConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateEdgeLogConfig
// 创建边缘集群日志采集配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateEdgeLogConfig(request *CreateEdgeLogConfigRequest) (response *CreateEdgeLogConfigResponse, err error) {
    return c.CreateEdgeLogConfigWithContext(context.Background(), request)
}

// CreateEdgeLogConfig
// 创建边缘集群日志采集配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateEdgeLogConfigWithContext(ctx context.Context, request *CreateEdgeLogConfigRequest) (response *CreateEdgeLogConfigResponse, err error) {
    if request == nil {
        request = NewCreateEdgeLogConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEdgeLogConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEdgeLogConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateImageCacheRequest() (request *CreateImageCacheRequest) {
    request = &CreateImageCacheRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CreateImageCache")
    
    
    return
}

func NewCreateImageCacheResponse() (response *CreateImageCacheResponse) {
    response = &CreateImageCacheResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateImageCache
// 创建镜像缓存的接口。创建过程中，请勿删除EKSCI实例和云盘，否则镜像缓存将创建失败。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateImageCache(request *CreateImageCacheRequest) (response *CreateImageCacheResponse, err error) {
    return c.CreateImageCacheWithContext(context.Background(), request)
}

// CreateImageCache
// 创建镜像缓存的接口。创建过程中，请勿删除EKSCI实例和云盘，否则镜像缓存将创建失败。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateImageCacheWithContext(ctx context.Context, request *CreateImageCacheRequest) (response *CreateImageCacheResponse, err error) {
    if request == nil {
        request = NewCreateImageCacheRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateImageCache require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateImageCacheResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrometheusAlertPolicyRequest() (request *CreatePrometheusAlertPolicyRequest) {
    request = &CreatePrometheusAlertPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CreatePrometheusAlertPolicy")
    
    
    return
}

func NewCreatePrometheusAlertPolicyResponse() (response *CreatePrometheusAlertPolicyResponse) {
    response = &CreatePrometheusAlertPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePrometheusAlertPolicy
// 创建告警策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_KUBECLIENTCONF = "FailedOperation.KubeClientConf"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreatePrometheusAlertPolicy(request *CreatePrometheusAlertPolicyRequest) (response *CreatePrometheusAlertPolicyResponse, err error) {
    return c.CreatePrometheusAlertPolicyWithContext(context.Background(), request)
}

// CreatePrometheusAlertPolicy
// 创建告警策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_KUBECLIENTCONF = "FailedOperation.KubeClientConf"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreatePrometheusAlertPolicyWithContext(ctx context.Context, request *CreatePrometheusAlertPolicyRequest) (response *CreatePrometheusAlertPolicyResponse, err error) {
    if request == nil {
        request = NewCreatePrometheusAlertPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrometheusAlertPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrometheusAlertPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrometheusAlertRuleRequest() (request *CreatePrometheusAlertRuleRequest) {
    request = &CreatePrometheusAlertRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CreatePrometheusAlertRule")
    
    
    return
}

func NewCreatePrometheusAlertRuleResponse() (response *CreatePrometheusAlertRuleResponse) {
    response = &CreatePrometheusAlertRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePrometheusAlertRule
// 创建告警规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreatePrometheusAlertRule(request *CreatePrometheusAlertRuleRequest) (response *CreatePrometheusAlertRuleResponse, err error) {
    return c.CreatePrometheusAlertRuleWithContext(context.Background(), request)
}

// CreatePrometheusAlertRule
// 创建告警规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreatePrometheusAlertRuleWithContext(ctx context.Context, request *CreatePrometheusAlertRuleRequest) (response *CreatePrometheusAlertRuleResponse, err error) {
    if request == nil {
        request = NewCreatePrometheusAlertRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrometheusAlertRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrometheusAlertRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrometheusClusterAgentRequest() (request *CreatePrometheusClusterAgentRequest) {
    request = &CreatePrometheusClusterAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CreatePrometheusClusterAgent")
    
    
    return
}

func NewCreatePrometheusClusterAgentResponse() (response *CreatePrometheusClusterAgentResponse) {
    response = &CreatePrometheusClusterAgentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePrometheusClusterAgent
// 与云监控融合的2.0实例关联集群
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) CreatePrometheusClusterAgent(request *CreatePrometheusClusterAgentRequest) (response *CreatePrometheusClusterAgentResponse, err error) {
    return c.CreatePrometheusClusterAgentWithContext(context.Background(), request)
}

// CreatePrometheusClusterAgent
// 与云监控融合的2.0实例关联集群
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) CreatePrometheusClusterAgentWithContext(ctx context.Context, request *CreatePrometheusClusterAgentRequest) (response *CreatePrometheusClusterAgentResponse, err error) {
    if request == nil {
        request = NewCreatePrometheusClusterAgentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrometheusClusterAgent require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrometheusClusterAgentResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrometheusConfigRequest() (request *CreatePrometheusConfigRequest) {
    request = &CreatePrometheusConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CreatePrometheusConfig")
    
    
    return
}

func NewCreatePrometheusConfigResponse() (response *CreatePrometheusConfigResponse) {
    response = &CreatePrometheusConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePrometheusConfig
// 创建prometheus配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  RESOURCEINUSE_RESOURCEEXISTALREADY = "ResourceInUse.ResourceExistAlready"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreatePrometheusConfig(request *CreatePrometheusConfigRequest) (response *CreatePrometheusConfigResponse, err error) {
    return c.CreatePrometheusConfigWithContext(context.Background(), request)
}

// CreatePrometheusConfig
// 创建prometheus配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  RESOURCEINUSE_RESOURCEEXISTALREADY = "ResourceInUse.ResourceExistAlready"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreatePrometheusConfigWithContext(ctx context.Context, request *CreatePrometheusConfigRequest) (response *CreatePrometheusConfigResponse, err error) {
    if request == nil {
        request = NewCreatePrometheusConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrometheusConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrometheusConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrometheusDashboardRequest() (request *CreatePrometheusDashboardRequest) {
    request = &CreatePrometheusDashboardRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CreatePrometheusDashboard")
    
    
    return
}

func NewCreatePrometheusDashboardResponse() (response *CreatePrometheusDashboardResponse) {
    response = &CreatePrometheusDashboardResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePrometheusDashboard
// 创建grafana监控面板
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) CreatePrometheusDashboard(request *CreatePrometheusDashboardRequest) (response *CreatePrometheusDashboardResponse, err error) {
    return c.CreatePrometheusDashboardWithContext(context.Background(), request)
}

// CreatePrometheusDashboard
// 创建grafana监控面板
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) CreatePrometheusDashboardWithContext(ctx context.Context, request *CreatePrometheusDashboardRequest) (response *CreatePrometheusDashboardResponse, err error) {
    if request == nil {
        request = NewCreatePrometheusDashboardRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrometheusDashboard require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrometheusDashboardResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrometheusGlobalNotificationRequest() (request *CreatePrometheusGlobalNotificationRequest) {
    request = &CreatePrometheusGlobalNotificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CreatePrometheusGlobalNotification")
    
    
    return
}

func NewCreatePrometheusGlobalNotificationResponse() (response *CreatePrometheusGlobalNotificationResponse) {
    response = &CreatePrometheusGlobalNotificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePrometheusGlobalNotification
// 创建全局告警通知渠道
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) CreatePrometheusGlobalNotification(request *CreatePrometheusGlobalNotificationRequest) (response *CreatePrometheusGlobalNotificationResponse, err error) {
    return c.CreatePrometheusGlobalNotificationWithContext(context.Background(), request)
}

// CreatePrometheusGlobalNotification
// 创建全局告警通知渠道
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) CreatePrometheusGlobalNotificationWithContext(ctx context.Context, request *CreatePrometheusGlobalNotificationRequest) (response *CreatePrometheusGlobalNotificationResponse, err error) {
    if request == nil {
        request = NewCreatePrometheusGlobalNotificationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrometheusGlobalNotification require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrometheusGlobalNotificationResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrometheusRecordRuleYamlRequest() (request *CreatePrometheusRecordRuleYamlRequest) {
    request = &CreatePrometheusRecordRuleYamlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CreatePrometheusRecordRuleYaml")
    
    
    return
}

func NewCreatePrometheusRecordRuleYamlResponse() (response *CreatePrometheusRecordRuleYamlResponse) {
    response = &CreatePrometheusRecordRuleYamlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePrometheusRecordRuleYaml
// 以Yaml的方式创建聚合规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreatePrometheusRecordRuleYaml(request *CreatePrometheusRecordRuleYamlRequest) (response *CreatePrometheusRecordRuleYamlResponse, err error) {
    return c.CreatePrometheusRecordRuleYamlWithContext(context.Background(), request)
}

// CreatePrometheusRecordRuleYaml
// 以Yaml的方式创建聚合规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreatePrometheusRecordRuleYamlWithContext(ctx context.Context, request *CreatePrometheusRecordRuleYamlRequest) (response *CreatePrometheusRecordRuleYamlResponse, err error) {
    if request == nil {
        request = NewCreatePrometheusRecordRuleYamlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrometheusRecordRuleYaml require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrometheusRecordRuleYamlResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrometheusTempRequest() (request *CreatePrometheusTempRequest) {
    request = &CreatePrometheusTempRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CreatePrometheusTemp")
    
    
    return
}

func NewCreatePrometheusTempResponse() (response *CreatePrometheusTempResponse) {
    response = &CreatePrometheusTempResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePrometheusTemp
// 创建一个云原生Prometheus模板
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) CreatePrometheusTemp(request *CreatePrometheusTempRequest) (response *CreatePrometheusTempResponse, err error) {
    return c.CreatePrometheusTempWithContext(context.Background(), request)
}

// CreatePrometheusTemp
// 创建一个云原生Prometheus模板
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) CreatePrometheusTempWithContext(ctx context.Context, request *CreatePrometheusTempRequest) (response *CreatePrometheusTempResponse, err error) {
    if request == nil {
        request = NewCreatePrometheusTempRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrometheusTemp require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrometheusTempResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrometheusTemplateRequest() (request *CreatePrometheusTemplateRequest) {
    request = &CreatePrometheusTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CreatePrometheusTemplate")
    
    
    return
}

func NewCreatePrometheusTemplateResponse() (response *CreatePrometheusTemplateResponse) {
    response = &CreatePrometheusTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePrometheusTemplate
// 创建一个云原生Prometheus模板实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) CreatePrometheusTemplate(request *CreatePrometheusTemplateRequest) (response *CreatePrometheusTemplateResponse, err error) {
    return c.CreatePrometheusTemplateWithContext(context.Background(), request)
}

// CreatePrometheusTemplate
// 创建一个云原生Prometheus模板实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) CreatePrometheusTemplateWithContext(ctx context.Context, request *CreatePrometheusTemplateRequest) (response *CreatePrometheusTemplateResponse, err error) {
    if request == nil {
        request = NewCreatePrometheusTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrometheusTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrometheusTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTKEEdgeClusterRequest() (request *CreateTKEEdgeClusterRequest) {
    request = &CreateTKEEdgeClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CreateTKEEdgeCluster")
    
    
    return
}

func NewCreateTKEEdgeClusterResponse() (response *CreateTKEEdgeClusterResponse) {
    response = &CreateTKEEdgeClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTKEEdgeCluster
// 创建边缘计算集群
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateTKEEdgeCluster(request *CreateTKEEdgeClusterRequest) (response *CreateTKEEdgeClusterResponse, err error) {
    return c.CreateTKEEdgeClusterWithContext(context.Background(), request)
}

// CreateTKEEdgeCluster
// 创建边缘计算集群
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateTKEEdgeClusterWithContext(ctx context.Context, request *CreateTKEEdgeClusterRequest) (response *CreateTKEEdgeClusterResponse, err error) {
    if request == nil {
        request = NewCreateTKEEdgeClusterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTKEEdgeCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTKEEdgeClusterResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClusterRequest() (request *DeleteClusterRequest) {
    request = &DeleteClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeleteCluster")
    
    
    return
}

func NewDeleteClusterResponse() (response *DeleteClusterResponse) {
    response = &DeleteClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCluster
// 删除集群(YUNAPI V3版本)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CAMNOAUTH = "FailedOperation.CamNoAuth"
//  FAILEDOPERATION_CLUSTERFORBIDDENTODELETE = "FailedOperation.ClusterForbiddenToDelete"
//  FAILEDOPERATION_CLUSTERNOTFOUND = "FailedOperation.ClusterNotFound"
//  FAILEDOPERATION_CLUSTERSTATE = "FailedOperation.ClusterState"
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLINETHTTP = "FailedOperation.ComponentClinetHttp"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_COMPONENTCLINETHTTP = "InternalError.ComponentClinetHttp"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_PUBLICCLUSTEROPNOTSUPPORT = "InternalError.PublicClusterOpNotSupport"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_CLUSTERINDELETIONPROTECTION = "OperationDenied.ClusterInDeletionProtection"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteCluster(request *DeleteClusterRequest) (response *DeleteClusterResponse, err error) {
    return c.DeleteClusterWithContext(context.Background(), request)
}

// DeleteCluster
// 删除集群(YUNAPI V3版本)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CAMNOAUTH = "FailedOperation.CamNoAuth"
//  FAILEDOPERATION_CLUSTERFORBIDDENTODELETE = "FailedOperation.ClusterForbiddenToDelete"
//  FAILEDOPERATION_CLUSTERNOTFOUND = "FailedOperation.ClusterNotFound"
//  FAILEDOPERATION_CLUSTERSTATE = "FailedOperation.ClusterState"
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLINETHTTP = "FailedOperation.ComponentClinetHttp"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_COMPONENTCLINETHTTP = "InternalError.ComponentClinetHttp"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_PUBLICCLUSTEROPNOTSUPPORT = "InternalError.PublicClusterOpNotSupport"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_CLUSTERINDELETIONPROTECTION = "OperationDenied.ClusterInDeletionProtection"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteClusterWithContext(ctx context.Context, request *DeleteClusterRequest) (response *DeleteClusterResponse, err error) {
    if request == nil {
        request = NewDeleteClusterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteClusterResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClusterAsGroupsRequest() (request *DeleteClusterAsGroupsRequest) {
    request = &DeleteClusterAsGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeleteClusterAsGroups")
    
    
    return
}

func NewDeleteClusterAsGroupsResponse() (response *DeleteClusterAsGroupsResponse) {
    response = &DeleteClusterAsGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteClusterAsGroups
// 删除集群伸缩组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLUSTERSTATE = "FailedOperation.ClusterState"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_ASCOMMON = "InternalError.AsCommon"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_PUBLICCLUSTEROPNOTSUPPORT = "InternalError.PublicClusterOpNotSupport"
//  INTERNALERROR_QUOTAMAXCLSLIMIT = "InternalError.QuotaMaxClsLimit"
//  INTERNALERROR_QUOTAMAXNODLIMIT = "InternalError.QuotaMaxNodLimit"
//  INTERNALERROR_QUOTAMAXRTLIMIT = "InternalError.QuotaMaxRtLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ASCOMMONERROR = "InvalidParameter.AsCommonError"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteClusterAsGroups(request *DeleteClusterAsGroupsRequest) (response *DeleteClusterAsGroupsResponse, err error) {
    return c.DeleteClusterAsGroupsWithContext(context.Background(), request)
}

// DeleteClusterAsGroups
// 删除集群伸缩组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLUSTERSTATE = "FailedOperation.ClusterState"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_ASCOMMON = "InternalError.AsCommon"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_PUBLICCLUSTEROPNOTSUPPORT = "InternalError.PublicClusterOpNotSupport"
//  INTERNALERROR_QUOTAMAXCLSLIMIT = "InternalError.QuotaMaxClsLimit"
//  INTERNALERROR_QUOTAMAXNODLIMIT = "InternalError.QuotaMaxNodLimit"
//  INTERNALERROR_QUOTAMAXRTLIMIT = "InternalError.QuotaMaxRtLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ASCOMMONERROR = "InvalidParameter.AsCommonError"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteClusterAsGroupsWithContext(ctx context.Context, request *DeleteClusterAsGroupsRequest) (response *DeleteClusterAsGroupsResponse, err error) {
    if request == nil {
        request = NewDeleteClusterAsGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteClusterAsGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteClusterAsGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClusterEndpointRequest() (request *DeleteClusterEndpointRequest) {
    request = &DeleteClusterEndpointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeleteClusterEndpoint")
    
    
    return
}

func NewDeleteClusterEndpointResponse() (response *DeleteClusterEndpointResponse) {
    response = &DeleteClusterEndpointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteClusterEndpoint
// 删除集群访问端口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_CVMCOMMON = "InternalError.CvmCommon"
//  INTERNALERROR_CVMNOTFOUND = "InternalError.CvmNotFound"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_KUBECOMMON = "InternalError.KubeCommon"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteClusterEndpoint(request *DeleteClusterEndpointRequest) (response *DeleteClusterEndpointResponse, err error) {
    return c.DeleteClusterEndpointWithContext(context.Background(), request)
}

// DeleteClusterEndpoint
// 删除集群访问端口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_CVMCOMMON = "InternalError.CvmCommon"
//  INTERNALERROR_CVMNOTFOUND = "InternalError.CvmNotFound"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_KUBECOMMON = "InternalError.KubeCommon"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteClusterEndpointWithContext(ctx context.Context, request *DeleteClusterEndpointRequest) (response *DeleteClusterEndpointResponse, err error) {
    if request == nil {
        request = NewDeleteClusterEndpointRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteClusterEndpoint require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteClusterEndpointResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClusterEndpointVipRequest() (request *DeleteClusterEndpointVipRequest) {
    request = &DeleteClusterEndpointVipRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeleteClusterEndpointVip")
    
    
    return
}

func NewDeleteClusterEndpointVipResponse() (response *DeleteClusterEndpointVipResponse) {
    response = &DeleteClusterEndpointVipResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteClusterEndpointVip
// 删除托管集群外网访问端口（老的方式，仅支持托管集群外网端口）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_DFWGETUSGCOUNT = "InternalError.DfwGetUSGCount"
//  INTERNALERROR_DFWGETUSGQUOTA = "InternalError.DfwGetUSGQuota"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteClusterEndpointVip(request *DeleteClusterEndpointVipRequest) (response *DeleteClusterEndpointVipResponse, err error) {
    return c.DeleteClusterEndpointVipWithContext(context.Background(), request)
}

// DeleteClusterEndpointVip
// 删除托管集群外网访问端口（老的方式，仅支持托管集群外网端口）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_DFWGETUSGCOUNT = "InternalError.DfwGetUSGCount"
//  INTERNALERROR_DFWGETUSGQUOTA = "InternalError.DfwGetUSGQuota"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteClusterEndpointVipWithContext(ctx context.Context, request *DeleteClusterEndpointVipRequest) (response *DeleteClusterEndpointVipResponse, err error) {
    if request == nil {
        request = NewDeleteClusterEndpointVipRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteClusterEndpointVip require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteClusterEndpointVipResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClusterInstancesRequest() (request *DeleteClusterInstancesRequest) {
    request = &DeleteClusterInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeleteClusterInstances")
    
    
    return
}

func NewDeleteClusterInstancesResponse() (response *DeleteClusterInstancesResponse) {
    response = &DeleteClusterInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteClusterInstances
// 删除集群中的实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_ASCOMMON = "FailedOperation.AsCommon"
//  FAILEDOPERATION_CLUSTERNOTFOUND = "FailedOperation.ClusterNotFound"
//  FAILEDOPERATION_CLUSTERSTATE = "FailedOperation.ClusterState"
//  FAILEDOPERATION_DBRECORDNOTFOUND = "FailedOperation.DbRecordNotFound"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ASCOMMON = "InternalError.AsCommon"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_PUBLICCLUSTEROPNOTSUPPORT = "InternalError.PublicClusterOpNotSupport"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DeleteClusterInstances(request *DeleteClusterInstancesRequest) (response *DeleteClusterInstancesResponse, err error) {
    return c.DeleteClusterInstancesWithContext(context.Background(), request)
}

// DeleteClusterInstances
// 删除集群中的实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_ASCOMMON = "FailedOperation.AsCommon"
//  FAILEDOPERATION_CLUSTERNOTFOUND = "FailedOperation.ClusterNotFound"
//  FAILEDOPERATION_CLUSTERSTATE = "FailedOperation.ClusterState"
//  FAILEDOPERATION_DBRECORDNOTFOUND = "FailedOperation.DbRecordNotFound"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ASCOMMON = "InternalError.AsCommon"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_PUBLICCLUSTEROPNOTSUPPORT = "InternalError.PublicClusterOpNotSupport"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DeleteClusterInstancesWithContext(ctx context.Context, request *DeleteClusterInstancesRequest) (response *DeleteClusterInstancesResponse, err error) {
    if request == nil {
        request = NewDeleteClusterInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteClusterInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteClusterInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClusterNodePoolRequest() (request *DeleteClusterNodePoolRequest) {
    request = &DeleteClusterNodePoolRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeleteClusterNodePool")
    
    
    return
}

func NewDeleteClusterNodePoolResponse() (response *DeleteClusterNodePoolResponse) {
    response = &DeleteClusterNodePoolResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteClusterNodePool
// 删除节点池
//
// 可能返回的错误码:
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteClusterNodePool(request *DeleteClusterNodePoolRequest) (response *DeleteClusterNodePoolResponse, err error) {
    return c.DeleteClusterNodePoolWithContext(context.Background(), request)
}

// DeleteClusterNodePool
// 删除节点池
//
// 可能返回的错误码:
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteClusterNodePoolWithContext(ctx context.Context, request *DeleteClusterNodePoolRequest) (response *DeleteClusterNodePoolResponse, err error) {
    if request == nil {
        request = NewDeleteClusterNodePoolRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteClusterNodePool require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteClusterNodePoolResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClusterRouteRequest() (request *DeleteClusterRouteRequest) {
    request = &DeleteClusterRouteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeleteClusterRoute")
    
    
    return
}

func NewDeleteClusterRouteResponse() (response *DeleteClusterRouteResponse) {
    response = &DeleteClusterRouteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteClusterRoute
// 删除集群路由
//
// 可能返回的错误码:
//  FAILEDOPERATION_VPCCOMMON = "FailedOperation.VpcCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_ROUTETABLENOTFOUND = "InternalError.RouteTableNotFound"
//  INTERNALERROR_VPCCOMMON = "InternalError.VpcCommon"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteClusterRoute(request *DeleteClusterRouteRequest) (response *DeleteClusterRouteResponse, err error) {
    return c.DeleteClusterRouteWithContext(context.Background(), request)
}

// DeleteClusterRoute
// 删除集群路由
//
// 可能返回的错误码:
//  FAILEDOPERATION_VPCCOMMON = "FailedOperation.VpcCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_ROUTETABLENOTFOUND = "InternalError.RouteTableNotFound"
//  INTERNALERROR_VPCCOMMON = "InternalError.VpcCommon"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteClusterRouteWithContext(ctx context.Context, request *DeleteClusterRouteRequest) (response *DeleteClusterRouteResponse, err error) {
    if request == nil {
        request = NewDeleteClusterRouteRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteClusterRoute require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteClusterRouteResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClusterRouteTableRequest() (request *DeleteClusterRouteTableRequest) {
    request = &DeleteClusterRouteTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeleteClusterRouteTable")
    
    
    return
}

func NewDeleteClusterRouteTableResponse() (response *DeleteClusterRouteTableResponse) {
    response = &DeleteClusterRouteTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteClusterRouteTable
// 删除集群路由表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_ROUTETABLENOTEMPTY = "InternalError.RouteTableNotEmpty"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ROUTETABLENOTEMPTY = "InvalidParameter.RouteTableNotEmpty"
func (c *Client) DeleteClusterRouteTable(request *DeleteClusterRouteTableRequest) (response *DeleteClusterRouteTableResponse, err error) {
    return c.DeleteClusterRouteTableWithContext(context.Background(), request)
}

// DeleteClusterRouteTable
// 删除集群路由表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_ROUTETABLENOTEMPTY = "InternalError.RouteTableNotEmpty"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ROUTETABLENOTEMPTY = "InvalidParameter.RouteTableNotEmpty"
func (c *Client) DeleteClusterRouteTableWithContext(ctx context.Context, request *DeleteClusterRouteTableRequest) (response *DeleteClusterRouteTableResponse, err error) {
    if request == nil {
        request = NewDeleteClusterRouteTableRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteClusterRouteTable require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteClusterRouteTableResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClusterVirtualNodeRequest() (request *DeleteClusterVirtualNodeRequest) {
    request = &DeleteClusterVirtualNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeleteClusterVirtualNode")
    
    
    return
}

func NewDeleteClusterVirtualNodeResponse() (response *DeleteClusterVirtualNodeResponse) {
    response = &DeleteClusterVirtualNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteClusterVirtualNode
// 删除虚拟节点
//
// 可能返回的错误码:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCEINUSE_EXISTRUNNINGPOD = "ResourceInUse.ExistRunningPod"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) DeleteClusterVirtualNode(request *DeleteClusterVirtualNodeRequest) (response *DeleteClusterVirtualNodeResponse, err error) {
    return c.DeleteClusterVirtualNodeWithContext(context.Background(), request)
}

// DeleteClusterVirtualNode
// 删除虚拟节点
//
// 可能返回的错误码:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCEINUSE_EXISTRUNNINGPOD = "ResourceInUse.ExistRunningPod"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) DeleteClusterVirtualNodeWithContext(ctx context.Context, request *DeleteClusterVirtualNodeRequest) (response *DeleteClusterVirtualNodeResponse, err error) {
    if request == nil {
        request = NewDeleteClusterVirtualNodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteClusterVirtualNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteClusterVirtualNodeResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClusterVirtualNodePoolRequest() (request *DeleteClusterVirtualNodePoolRequest) {
    request = &DeleteClusterVirtualNodePoolRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeleteClusterVirtualNodePool")
    
    
    return
}

func NewDeleteClusterVirtualNodePoolResponse() (response *DeleteClusterVirtualNodePoolResponse) {
    response = &DeleteClusterVirtualNodePoolResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteClusterVirtualNodePool
// 删除虚拟节点池
//
// 可能返回的错误码:
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  RESOURCEINUSE_EXISTRUNNINGPOD = "ResourceInUse.ExistRunningPod"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) DeleteClusterVirtualNodePool(request *DeleteClusterVirtualNodePoolRequest) (response *DeleteClusterVirtualNodePoolResponse, err error) {
    return c.DeleteClusterVirtualNodePoolWithContext(context.Background(), request)
}

// DeleteClusterVirtualNodePool
// 删除虚拟节点池
//
// 可能返回的错误码:
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  RESOURCEINUSE_EXISTRUNNINGPOD = "ResourceInUse.ExistRunningPod"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) DeleteClusterVirtualNodePoolWithContext(ctx context.Context, request *DeleteClusterVirtualNodePoolRequest) (response *DeleteClusterVirtualNodePoolResponse, err error) {
    if request == nil {
        request = NewDeleteClusterVirtualNodePoolRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteClusterVirtualNodePool require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteClusterVirtualNodePoolResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteECMInstancesRequest() (request *DeleteECMInstancesRequest) {
    request = &DeleteECMInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeleteECMInstances")
    
    
    return
}

func NewDeleteECMInstancesResponse() (response *DeleteECMInstancesResponse) {
    response = &DeleteECMInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteECMInstances
// 删除ECM实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteECMInstances(request *DeleteECMInstancesRequest) (response *DeleteECMInstancesResponse, err error) {
    return c.DeleteECMInstancesWithContext(context.Background(), request)
}

// DeleteECMInstances
// 删除ECM实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteECMInstancesWithContext(ctx context.Context, request *DeleteECMInstancesRequest) (response *DeleteECMInstancesResponse, err error) {
    if request == nil {
        request = NewDeleteECMInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteECMInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteECMInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEKSClusterRequest() (request *DeleteEKSClusterRequest) {
    request = &DeleteEKSClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeleteEKSCluster")
    
    
    return
}

func NewDeleteEKSClusterResponse() (response *DeleteEKSClusterResponse) {
    response = &DeleteEKSClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteEKSCluster
// 删除弹性集群(yunapiv3)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteEKSCluster(request *DeleteEKSClusterRequest) (response *DeleteEKSClusterResponse, err error) {
    return c.DeleteEKSClusterWithContext(context.Background(), request)
}

// DeleteEKSCluster
// 删除弹性集群(yunapiv3)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteEKSClusterWithContext(ctx context.Context, request *DeleteEKSClusterRequest) (response *DeleteEKSClusterResponse, err error) {
    if request == nil {
        request = NewDeleteEKSClusterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteEKSCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteEKSClusterResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEKSContainerInstancesRequest() (request *DeleteEKSContainerInstancesRequest) {
    request = &DeleteEKSContainerInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeleteEKSContainerInstances")
    
    
    return
}

func NewDeleteEKSContainerInstancesResponse() (response *DeleteEKSContainerInstancesResponse) {
    response = &DeleteEKSContainerInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteEKSContainerInstances
// 删除容器实例，可批量删除
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CONTAINERNOTFOUND = "InternalError.ContainerNotFound"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteEKSContainerInstances(request *DeleteEKSContainerInstancesRequest) (response *DeleteEKSContainerInstancesResponse, err error) {
    return c.DeleteEKSContainerInstancesWithContext(context.Background(), request)
}

// DeleteEKSContainerInstances
// 删除容器实例，可批量删除
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CONTAINERNOTFOUND = "InternalError.ContainerNotFound"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteEKSContainerInstancesWithContext(ctx context.Context, request *DeleteEKSContainerInstancesRequest) (response *DeleteEKSContainerInstancesResponse, err error) {
    if request == nil {
        request = NewDeleteEKSContainerInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteEKSContainerInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteEKSContainerInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEdgeCVMInstancesRequest() (request *DeleteEdgeCVMInstancesRequest) {
    request = &DeleteEdgeCVMInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeleteEdgeCVMInstances")
    
    
    return
}

func NewDeleteEdgeCVMInstancesResponse() (response *DeleteEdgeCVMInstancesResponse) {
    response = &DeleteEdgeCVMInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteEdgeCVMInstances
// 删除边缘容器CVM实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteEdgeCVMInstances(request *DeleteEdgeCVMInstancesRequest) (response *DeleteEdgeCVMInstancesResponse, err error) {
    return c.DeleteEdgeCVMInstancesWithContext(context.Background(), request)
}

// DeleteEdgeCVMInstances
// 删除边缘容器CVM实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteEdgeCVMInstancesWithContext(ctx context.Context, request *DeleteEdgeCVMInstancesRequest) (response *DeleteEdgeCVMInstancesResponse, err error) {
    if request == nil {
        request = NewDeleteEdgeCVMInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteEdgeCVMInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteEdgeCVMInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEdgeClusterInstancesRequest() (request *DeleteEdgeClusterInstancesRequest) {
    request = &DeleteEdgeClusterInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeleteEdgeClusterInstances")
    
    
    return
}

func NewDeleteEdgeClusterInstancesResponse() (response *DeleteEdgeClusterInstancesResponse) {
    response = &DeleteEdgeClusterInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteEdgeClusterInstances
// 删除边缘计算实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteEdgeClusterInstances(request *DeleteEdgeClusterInstancesRequest) (response *DeleteEdgeClusterInstancesResponse, err error) {
    return c.DeleteEdgeClusterInstancesWithContext(context.Background(), request)
}

// DeleteEdgeClusterInstances
// 删除边缘计算实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteEdgeClusterInstancesWithContext(ctx context.Context, request *DeleteEdgeClusterInstancesRequest) (response *DeleteEdgeClusterInstancesResponse, err error) {
    if request == nil {
        request = NewDeleteEdgeClusterInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteEdgeClusterInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteEdgeClusterInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteImageCachesRequest() (request *DeleteImageCachesRequest) {
    request = &DeleteImageCachesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeleteImageCaches")
    
    
    return
}

func NewDeleteImageCachesResponse() (response *DeleteImageCachesResponse) {
    response = &DeleteImageCachesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteImageCaches
// 批量删除镜像缓存
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteImageCaches(request *DeleteImageCachesRequest) (response *DeleteImageCachesResponse, err error) {
    return c.DeleteImageCachesWithContext(context.Background(), request)
}

// DeleteImageCaches
// 批量删除镜像缓存
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteImageCachesWithContext(ctx context.Context, request *DeleteImageCachesRequest) (response *DeleteImageCachesResponse, err error) {
    if request == nil {
        request = NewDeleteImageCachesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteImageCaches require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteImageCachesResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePrometheusAlertPolicyRequest() (request *DeletePrometheusAlertPolicyRequest) {
    request = &DeletePrometheusAlertPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeletePrometheusAlertPolicy")
    
    
    return
}

func NewDeletePrometheusAlertPolicyResponse() (response *DeletePrometheusAlertPolicyResponse) {
    response = &DeletePrometheusAlertPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeletePrometheusAlertPolicy
// 删除2.0实例告警策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DeletePrometheusAlertPolicy(request *DeletePrometheusAlertPolicyRequest) (response *DeletePrometheusAlertPolicyResponse, err error) {
    return c.DeletePrometheusAlertPolicyWithContext(context.Background(), request)
}

// DeletePrometheusAlertPolicy
// 删除2.0实例告警策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DeletePrometheusAlertPolicyWithContext(ctx context.Context, request *DeletePrometheusAlertPolicyRequest) (response *DeletePrometheusAlertPolicyResponse, err error) {
    if request == nil {
        request = NewDeletePrometheusAlertPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePrometheusAlertPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePrometheusAlertPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePrometheusAlertRuleRequest() (request *DeletePrometheusAlertRuleRequest) {
    request = &DeletePrometheusAlertRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeletePrometheusAlertRule")
    
    
    return
}

func NewDeletePrometheusAlertRuleResponse() (response *DeletePrometheusAlertRuleResponse) {
    response = &DeletePrometheusAlertRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeletePrometheusAlertRule
// 删除告警规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DeletePrometheusAlertRule(request *DeletePrometheusAlertRuleRequest) (response *DeletePrometheusAlertRuleResponse, err error) {
    return c.DeletePrometheusAlertRuleWithContext(context.Background(), request)
}

// DeletePrometheusAlertRule
// 删除告警规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DeletePrometheusAlertRuleWithContext(ctx context.Context, request *DeletePrometheusAlertRuleRequest) (response *DeletePrometheusAlertRuleResponse, err error) {
    if request == nil {
        request = NewDeletePrometheusAlertRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePrometheusAlertRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePrometheusAlertRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePrometheusClusterAgentRequest() (request *DeletePrometheusClusterAgentRequest) {
    request = &DeletePrometheusClusterAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeletePrometheusClusterAgent")
    
    
    return
}

func NewDeletePrometheusClusterAgentResponse() (response *DeletePrometheusClusterAgentResponse) {
    response = &DeletePrometheusClusterAgentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeletePrometheusClusterAgent
// 解除TMP实例的集群关联
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBRECORDNOTFOUND = "FailedOperation.DbRecordNotFound"
//  FAILEDOPERATION_KUBECLIENTCONF = "FailedOperation.KubeClientConf"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DeletePrometheusClusterAgent(request *DeletePrometheusClusterAgentRequest) (response *DeletePrometheusClusterAgentResponse, err error) {
    return c.DeletePrometheusClusterAgentWithContext(context.Background(), request)
}

// DeletePrometheusClusterAgent
// 解除TMP实例的集群关联
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBRECORDNOTFOUND = "FailedOperation.DbRecordNotFound"
//  FAILEDOPERATION_KUBECLIENTCONF = "FailedOperation.KubeClientConf"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DeletePrometheusClusterAgentWithContext(ctx context.Context, request *DeletePrometheusClusterAgentRequest) (response *DeletePrometheusClusterAgentResponse, err error) {
    if request == nil {
        request = NewDeletePrometheusClusterAgentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePrometheusClusterAgent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePrometheusClusterAgentResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePrometheusConfigRequest() (request *DeletePrometheusConfigRequest) {
    request = &DeletePrometheusConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeletePrometheusConfig")
    
    
    return
}

func NewDeletePrometheusConfigResponse() (response *DeletePrometheusConfigResponse) {
    response = &DeletePrometheusConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeletePrometheusConfig
// 删除Prometheus配置，如果目标不存在，将返回成功
//
// 可能返回的错误码:
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
func (c *Client) DeletePrometheusConfig(request *DeletePrometheusConfigRequest) (response *DeletePrometheusConfigResponse, err error) {
    return c.DeletePrometheusConfigWithContext(context.Background(), request)
}

// DeletePrometheusConfig
// 删除Prometheus配置，如果目标不存在，将返回成功
//
// 可能返回的错误码:
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
func (c *Client) DeletePrometheusConfigWithContext(ctx context.Context, request *DeletePrometheusConfigRequest) (response *DeletePrometheusConfigResponse, err error) {
    if request == nil {
        request = NewDeletePrometheusConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePrometheusConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePrometheusConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePrometheusRecordRuleYamlRequest() (request *DeletePrometheusRecordRuleYamlRequest) {
    request = &DeletePrometheusRecordRuleYamlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeletePrometheusRecordRuleYaml")
    
    
    return
}

func NewDeletePrometheusRecordRuleYamlResponse() (response *DeletePrometheusRecordRuleYamlResponse) {
    response = &DeletePrometheusRecordRuleYamlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeletePrometheusRecordRuleYaml
// 删除聚合实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DeletePrometheusRecordRuleYaml(request *DeletePrometheusRecordRuleYamlRequest) (response *DeletePrometheusRecordRuleYamlResponse, err error) {
    return c.DeletePrometheusRecordRuleYamlWithContext(context.Background(), request)
}

// DeletePrometheusRecordRuleYaml
// 删除聚合实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DeletePrometheusRecordRuleYamlWithContext(ctx context.Context, request *DeletePrometheusRecordRuleYamlRequest) (response *DeletePrometheusRecordRuleYamlResponse, err error) {
    if request == nil {
        request = NewDeletePrometheusRecordRuleYamlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePrometheusRecordRuleYaml require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePrometheusRecordRuleYamlResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePrometheusTempRequest() (request *DeletePrometheusTempRequest) {
    request = &DeletePrometheusTempRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeletePrometheusTemp")
    
    
    return
}

func NewDeletePrometheusTempResponse() (response *DeletePrometheusTempResponse) {
    response = &DeletePrometheusTempResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeletePrometheusTemp
// 删除一个云原生Prometheus配置模板
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
func (c *Client) DeletePrometheusTemp(request *DeletePrometheusTempRequest) (response *DeletePrometheusTempResponse, err error) {
    return c.DeletePrometheusTempWithContext(context.Background(), request)
}

// DeletePrometheusTemp
// 删除一个云原生Prometheus配置模板
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
func (c *Client) DeletePrometheusTempWithContext(ctx context.Context, request *DeletePrometheusTempRequest) (response *DeletePrometheusTempResponse, err error) {
    if request == nil {
        request = NewDeletePrometheusTempRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePrometheusTemp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePrometheusTempResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePrometheusTempSyncRequest() (request *DeletePrometheusTempSyncRequest) {
    request = &DeletePrometheusTempSyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeletePrometheusTempSync")
    
    
    return
}

func NewDeletePrometheusTempSyncResponse() (response *DeletePrometheusTempSyncResponse) {
    response = &DeletePrometheusTempSyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeletePrometheusTempSync
// 解除模板同步，这将会删除目标中该模板所生产的配置，针对V2版本实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_DBRECORDNOTFOUND = "FailedOperation.DbRecordNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
func (c *Client) DeletePrometheusTempSync(request *DeletePrometheusTempSyncRequest) (response *DeletePrometheusTempSyncResponse, err error) {
    return c.DeletePrometheusTempSyncWithContext(context.Background(), request)
}

// DeletePrometheusTempSync
// 解除模板同步，这将会删除目标中该模板所生产的配置，针对V2版本实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_DBRECORDNOTFOUND = "FailedOperation.DbRecordNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
func (c *Client) DeletePrometheusTempSyncWithContext(ctx context.Context, request *DeletePrometheusTempSyncRequest) (response *DeletePrometheusTempSyncResponse, err error) {
    if request == nil {
        request = NewDeletePrometheusTempSyncRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePrometheusTempSync require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePrometheusTempSyncResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePrometheusTemplateRequest() (request *DeletePrometheusTemplateRequest) {
    request = &DeletePrometheusTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeletePrometheusTemplate")
    
    
    return
}

func NewDeletePrometheusTemplateResponse() (response *DeletePrometheusTemplateResponse) {
    response = &DeletePrometheusTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeletePrometheusTemplate
// 删除一个云原生Prometheus配置模板
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
func (c *Client) DeletePrometheusTemplate(request *DeletePrometheusTemplateRequest) (response *DeletePrometheusTemplateResponse, err error) {
    return c.DeletePrometheusTemplateWithContext(context.Background(), request)
}

// DeletePrometheusTemplate
// 删除一个云原生Prometheus配置模板
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
func (c *Client) DeletePrometheusTemplateWithContext(ctx context.Context, request *DeletePrometheusTemplateRequest) (response *DeletePrometheusTemplateResponse, err error) {
    if request == nil {
        request = NewDeletePrometheusTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePrometheusTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePrometheusTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePrometheusTemplateSyncRequest() (request *DeletePrometheusTemplateSyncRequest) {
    request = &DeletePrometheusTemplateSyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeletePrometheusTemplateSync")
    
    
    return
}

func NewDeletePrometheusTemplateSyncResponse() (response *DeletePrometheusTemplateSyncResponse) {
    response = &DeletePrometheusTemplateSyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeletePrometheusTemplateSync
// 取消模板同步，这将会删除目标中该模板所生产的配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBRECORDNOTFOUND = "FailedOperation.DbRecordNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeletePrometheusTemplateSync(request *DeletePrometheusTemplateSyncRequest) (response *DeletePrometheusTemplateSyncResponse, err error) {
    return c.DeletePrometheusTemplateSyncWithContext(context.Background(), request)
}

// DeletePrometheusTemplateSync
// 取消模板同步，这将会删除目标中该模板所生产的配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBRECORDNOTFOUND = "FailedOperation.DbRecordNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeletePrometheusTemplateSyncWithContext(ctx context.Context, request *DeletePrometheusTemplateSyncRequest) (response *DeletePrometheusTemplateSyncResponse, err error) {
    if request == nil {
        request = NewDeletePrometheusTemplateSyncRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePrometheusTemplateSync require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePrometheusTemplateSyncResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTKEEdgeClusterRequest() (request *DeleteTKEEdgeClusterRequest) {
    request = &DeleteTKEEdgeClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeleteTKEEdgeCluster")
    
    
    return
}

func NewDeleteTKEEdgeClusterResponse() (response *DeleteTKEEdgeClusterResponse) {
    response = &DeleteTKEEdgeClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteTKEEdgeCluster
// 删除边缘计算集群
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteTKEEdgeCluster(request *DeleteTKEEdgeClusterRequest) (response *DeleteTKEEdgeClusterResponse, err error) {
    return c.DeleteTKEEdgeClusterWithContext(context.Background(), request)
}

// DeleteTKEEdgeCluster
// 删除边缘计算集群
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteTKEEdgeClusterWithContext(ctx context.Context, request *DeleteTKEEdgeClusterRequest) (response *DeleteTKEEdgeClusterResponse, err error) {
    if request == nil {
        request = NewDeleteTKEEdgeClusterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTKEEdgeCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTKEEdgeClusterResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAvailableClusterVersionRequest() (request *DescribeAvailableClusterVersionRequest) {
    request = &DescribeAvailableClusterVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeAvailableClusterVersion")
    
    
    return
}

func NewDescribeAvailableClusterVersionResponse() (response *DescribeAvailableClusterVersionResponse) {
    response = &DescribeAvailableClusterVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAvailableClusterVersion
// 获取集群可以升级的所有版本
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERINABNORMALSTAT = "ResourceUnavailable.ClusterInAbnormalStat"
func (c *Client) DescribeAvailableClusterVersion(request *DescribeAvailableClusterVersionRequest) (response *DescribeAvailableClusterVersionResponse, err error) {
    return c.DescribeAvailableClusterVersionWithContext(context.Background(), request)
}

// DescribeAvailableClusterVersion
// 获取集群可以升级的所有版本
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERINABNORMALSTAT = "ResourceUnavailable.ClusterInAbnormalStat"
func (c *Client) DescribeAvailableClusterVersionWithContext(ctx context.Context, request *DescribeAvailableClusterVersionRequest) (response *DescribeAvailableClusterVersionResponse, err error) {
    if request == nil {
        request = NewDescribeAvailableClusterVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAvailableClusterVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAvailableClusterVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAvailableTKEEdgeVersionRequest() (request *DescribeAvailableTKEEdgeVersionRequest) {
    request = &DescribeAvailableTKEEdgeVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeAvailableTKEEdgeVersion")
    
    
    return
}

func NewDescribeAvailableTKEEdgeVersionResponse() (response *DescribeAvailableTKEEdgeVersionResponse) {
    response = &DescribeAvailableTKEEdgeVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAvailableTKEEdgeVersion
// 边缘计算支持版本和k8s版本
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAvailableTKEEdgeVersion(request *DescribeAvailableTKEEdgeVersionRequest) (response *DescribeAvailableTKEEdgeVersionResponse, err error) {
    return c.DescribeAvailableTKEEdgeVersionWithContext(context.Background(), request)
}

// DescribeAvailableTKEEdgeVersion
// 边缘计算支持版本和k8s版本
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAvailableTKEEdgeVersionWithContext(ctx context.Context, request *DescribeAvailableTKEEdgeVersionRequest) (response *DescribeAvailableTKEEdgeVersionResponse, err error) {
    if request == nil {
        request = NewDescribeAvailableTKEEdgeVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAvailableTKEEdgeVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAvailableTKEEdgeVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterAsGroupOptionRequest() (request *DescribeClusterAsGroupOptionRequest) {
    request = &DescribeClusterAsGroupOptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterAsGroupOption")
    
    
    return
}

func NewDescribeClusterAsGroupOptionResponse() (response *DescribeClusterAsGroupOptionResponse) {
    response = &DescribeClusterAsGroupOptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClusterAsGroupOption
// 集群弹性伸缩配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLUSTERNOTFOUND = "FailedOperation.ClusterNotFound"
//  FAILEDOPERATION_CLUSTERSTATE = "FailedOperation.ClusterState"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_ASCOMMON = "InternalError.AsCommon"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeClusterAsGroupOption(request *DescribeClusterAsGroupOptionRequest) (response *DescribeClusterAsGroupOptionResponse, err error) {
    return c.DescribeClusterAsGroupOptionWithContext(context.Background(), request)
}

// DescribeClusterAsGroupOption
// 集群弹性伸缩配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLUSTERNOTFOUND = "FailedOperation.ClusterNotFound"
//  FAILEDOPERATION_CLUSTERSTATE = "FailedOperation.ClusterState"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_ASCOMMON = "InternalError.AsCommon"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeClusterAsGroupOptionWithContext(ctx context.Context, request *DescribeClusterAsGroupOptionRequest) (response *DescribeClusterAsGroupOptionResponse, err error) {
    if request == nil {
        request = NewDescribeClusterAsGroupOptionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterAsGroupOption require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterAsGroupOptionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterAsGroupsRequest() (request *DescribeClusterAsGroupsRequest) {
    request = &DescribeClusterAsGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterAsGroups")
    
    
    return
}

func NewDescribeClusterAsGroupsResponse() (response *DescribeClusterAsGroupsResponse) {
    response = &DescribeClusterAsGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClusterAsGroups
// 集群关联的伸缩组列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLUSTERNOTFOUND = "FailedOperation.ClusterNotFound"
//  FAILEDOPERATION_CLUSTERSTATE = "FailedOperation.ClusterState"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_ASCOMMON = "InternalError.AsCommon"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_PODNOTFOUND = "InternalError.PodNotFound"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_VPCCOMMON = "InternalError.VpcCommon"
//  INTERNALERROR_VPCPEERNOTFOUND = "InternalError.VpcPeerNotFound"
//  INTERNALERROR_VPCRECODRNOTFOUND = "InternalError.VpcRecodrNotFound"
func (c *Client) DescribeClusterAsGroups(request *DescribeClusterAsGroupsRequest) (response *DescribeClusterAsGroupsResponse, err error) {
    return c.DescribeClusterAsGroupsWithContext(context.Background(), request)
}

// DescribeClusterAsGroups
// 集群关联的伸缩组列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLUSTERNOTFOUND = "FailedOperation.ClusterNotFound"
//  FAILEDOPERATION_CLUSTERSTATE = "FailedOperation.ClusterState"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_ASCOMMON = "InternalError.AsCommon"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_PODNOTFOUND = "InternalError.PodNotFound"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_VPCCOMMON = "InternalError.VpcCommon"
//  INTERNALERROR_VPCPEERNOTFOUND = "InternalError.VpcPeerNotFound"
//  INTERNALERROR_VPCRECODRNOTFOUND = "InternalError.VpcRecodrNotFound"
func (c *Client) DescribeClusterAsGroupsWithContext(ctx context.Context, request *DescribeClusterAsGroupsRequest) (response *DescribeClusterAsGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterAsGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterAsGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterAsGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterAuthenticationOptionsRequest() (request *DescribeClusterAuthenticationOptionsRequest) {
    request = &DescribeClusterAuthenticationOptionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterAuthenticationOptions")
    
    
    return
}

func NewDescribeClusterAuthenticationOptionsResponse() (response *DescribeClusterAuthenticationOptionsResponse) {
    response = &DescribeClusterAuthenticationOptionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClusterAuthenticationOptions
// 查看集群认证配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) DescribeClusterAuthenticationOptions(request *DescribeClusterAuthenticationOptionsRequest) (response *DescribeClusterAuthenticationOptionsResponse, err error) {
    return c.DescribeClusterAuthenticationOptionsWithContext(context.Background(), request)
}

// DescribeClusterAuthenticationOptions
// 查看集群认证配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) DescribeClusterAuthenticationOptionsWithContext(ctx context.Context, request *DescribeClusterAuthenticationOptionsRequest) (response *DescribeClusterAuthenticationOptionsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterAuthenticationOptionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterAuthenticationOptions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterAuthenticationOptionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterCommonNamesRequest() (request *DescribeClusterCommonNamesRequest) {
    request = &DescribeClusterCommonNamesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterCommonNames")
    
    
    return
}

func NewDescribeClusterCommonNamesResponse() (response *DescribeClusterCommonNamesResponse) {
    response = &DescribeClusterCommonNamesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClusterCommonNames
// 获取指定子账户在RBAC授权模式中对应kube-apiserver客户端证书的CommonName字段，如果没有客户端证书，将会签发一个，此接口有最大传入子账户数量上限，当前为50
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INTERNALERROR_WHITELISTUNEXPECTEDERROR = "InternalError.WhitelistUnexpectedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNSUPPORTEDOPERATION_NOTINWHITELIST = "UnsupportedOperation.NotInWhitelist"
func (c *Client) DescribeClusterCommonNames(request *DescribeClusterCommonNamesRequest) (response *DescribeClusterCommonNamesResponse, err error) {
    return c.DescribeClusterCommonNamesWithContext(context.Background(), request)
}

// DescribeClusterCommonNames
// 获取指定子账户在RBAC授权模式中对应kube-apiserver客户端证书的CommonName字段，如果没有客户端证书，将会签发一个，此接口有最大传入子账户数量上限，当前为50
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INTERNALERROR_WHITELISTUNEXPECTEDERROR = "InternalError.WhitelistUnexpectedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNSUPPORTEDOPERATION_NOTINWHITELIST = "UnsupportedOperation.NotInWhitelist"
func (c *Client) DescribeClusterCommonNamesWithContext(ctx context.Context, request *DescribeClusterCommonNamesRequest) (response *DescribeClusterCommonNamesResponse, err error) {
    if request == nil {
        request = NewDescribeClusterCommonNamesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterCommonNames require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterCommonNamesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterControllersRequest() (request *DescribeClusterControllersRequest) {
    request = &DescribeClusterControllersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterControllers")
    
    
    return
}

func NewDescribeClusterControllersResponse() (response *DescribeClusterControllersResponse) {
    response = &DescribeClusterControllersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClusterControllers
// 用于查询Kubernetes的各个原生控制器是否开启
//
// 可能返回的错误码:
//  INTERNALERROR_KUBECLIENTCREATE = "InternalError.KubeClientCreate"
//  INTERNALERROR_KUBECOMMON = "InternalError.KubeCommon"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) DescribeClusterControllers(request *DescribeClusterControllersRequest) (response *DescribeClusterControllersResponse, err error) {
    return c.DescribeClusterControllersWithContext(context.Background(), request)
}

// DescribeClusterControllers
// 用于查询Kubernetes的各个原生控制器是否开启
//
// 可能返回的错误码:
//  INTERNALERROR_KUBECLIENTCREATE = "InternalError.KubeClientCreate"
//  INTERNALERROR_KUBECOMMON = "InternalError.KubeCommon"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) DescribeClusterControllersWithContext(ctx context.Context, request *DescribeClusterControllersRequest) (response *DescribeClusterControllersResponse, err error) {
    if request == nil {
        request = NewDescribeClusterControllersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterControllers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterControllersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterEndpointStatusRequest() (request *DescribeClusterEndpointStatusRequest) {
    request = &DescribeClusterEndpointStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterEndpointStatus")
    
    
    return
}

func NewDescribeClusterEndpointStatusResponse() (response *DescribeClusterEndpointStatusResponse) {
    response = &DescribeClusterEndpointStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClusterEndpointStatus
// 查询集群访问端口状态(独立集群开启内网/外网访问，托管集群支持开启内网访问)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_KUBECLIENTCONNECTION = "FailedOperation.KubeClientConnection"
//  FAILEDOPERATION_KUBERNETESINTERNAL = "FailedOperation.KubernetesInternal"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_KUBECLIENTCONNECTION = "InternalError.KubeClientConnection"
//  INTERNALERROR_KUBECOMMON = "InternalError.KubeCommon"
//  INTERNALERROR_KUBERNETESINTERNAL = "InternalError.KubernetesInternal"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INTERNALERROR_VPCCOMMON = "InternalError.VpcCommon"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeClusterEndpointStatus(request *DescribeClusterEndpointStatusRequest) (response *DescribeClusterEndpointStatusResponse, err error) {
    return c.DescribeClusterEndpointStatusWithContext(context.Background(), request)
}

// DescribeClusterEndpointStatus
// 查询集群访问端口状态(独立集群开启内网/外网访问，托管集群支持开启内网访问)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_KUBECLIENTCONNECTION = "FailedOperation.KubeClientConnection"
//  FAILEDOPERATION_KUBERNETESINTERNAL = "FailedOperation.KubernetesInternal"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_KUBECLIENTCONNECTION = "InternalError.KubeClientConnection"
//  INTERNALERROR_KUBECOMMON = "InternalError.KubeCommon"
//  INTERNALERROR_KUBERNETESINTERNAL = "InternalError.KubernetesInternal"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INTERNALERROR_VPCCOMMON = "InternalError.VpcCommon"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeClusterEndpointStatusWithContext(ctx context.Context, request *DescribeClusterEndpointStatusRequest) (response *DescribeClusterEndpointStatusResponse, err error) {
    if request == nil {
        request = NewDescribeClusterEndpointStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterEndpointStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterEndpointStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterEndpointVipStatusRequest() (request *DescribeClusterEndpointVipStatusRequest) {
    request = &DescribeClusterEndpointVipStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterEndpointVipStatus")
    
    
    return
}

func NewDescribeClusterEndpointVipStatusResponse() (response *DescribeClusterEndpointVipStatusResponse) {
    response = &DescribeClusterEndpointVipStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClusterEndpointVipStatus
// 查询集群开启端口流程状态(仅支持托管集群外网端口)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_ASCOMMON = "InternalError.AsCommon"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_DFWGETUSGCOUNT = "InternalError.DfwGetUSGCount"
//  INTERNALERROR_DFWGETUSGQUOTA = "InternalError.DfwGetUSGQuota"
//  INTERNALERROR_IMAGEIDNOTFOUND = "InternalError.ImageIdNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ASCOMMONERROR = "InvalidParameter.AsCommonError"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeClusterEndpointVipStatus(request *DescribeClusterEndpointVipStatusRequest) (response *DescribeClusterEndpointVipStatusResponse, err error) {
    return c.DescribeClusterEndpointVipStatusWithContext(context.Background(), request)
}

// DescribeClusterEndpointVipStatus
// 查询集群开启端口流程状态(仅支持托管集群外网端口)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_ASCOMMON = "InternalError.AsCommon"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_DFWGETUSGCOUNT = "InternalError.DfwGetUSGCount"
//  INTERNALERROR_DFWGETUSGQUOTA = "InternalError.DfwGetUSGQuota"
//  INTERNALERROR_IMAGEIDNOTFOUND = "InternalError.ImageIdNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ASCOMMONERROR = "InvalidParameter.AsCommonError"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeClusterEndpointVipStatusWithContext(ctx context.Context, request *DescribeClusterEndpointVipStatusRequest) (response *DescribeClusterEndpointVipStatusResponse, err error) {
    if request == nil {
        request = NewDescribeClusterEndpointVipStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterEndpointVipStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterEndpointVipStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterEndpointsRequest() (request *DescribeClusterEndpointsRequest) {
    request = &DescribeClusterEndpointsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterEndpoints")
    
    
    return
}

func NewDescribeClusterEndpointsResponse() (response *DescribeClusterEndpointsResponse) {
    response = &DescribeClusterEndpointsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClusterEndpoints
// 获取集群的访问地址，包括内网地址，外网地址，外网域名，外网访问安全策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  FAILEDOPERATION_KUBERNETESCLIENTBUILDERROR = "FailedOperation.KubernetesClientBuildError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COMPONENTCLIENTHTTP = "InternalError.ComponentClientHttp"
//  INTERNALERROR_KUBECLIENTCONNECTION = "InternalError.KubeClientConnection"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_KUBERESOURCENOTFOUND = "ResourceNotFound.KubeResourceNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERINABNORMALSTAT = "ResourceUnavailable.ClusterInAbnormalStat"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeClusterEndpoints(request *DescribeClusterEndpointsRequest) (response *DescribeClusterEndpointsResponse, err error) {
    return c.DescribeClusterEndpointsWithContext(context.Background(), request)
}

// DescribeClusterEndpoints
// 获取集群的访问地址，包括内网地址，外网地址，外网域名，外网访问安全策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  FAILEDOPERATION_KUBERNETESCLIENTBUILDERROR = "FailedOperation.KubernetesClientBuildError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COMPONENTCLIENTHTTP = "InternalError.ComponentClientHttp"
//  INTERNALERROR_KUBECLIENTCONNECTION = "InternalError.KubeClientConnection"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_KUBERESOURCENOTFOUND = "ResourceNotFound.KubeResourceNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERINABNORMALSTAT = "ResourceUnavailable.ClusterInAbnormalStat"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeClusterEndpointsWithContext(ctx context.Context, request *DescribeClusterEndpointsRequest) (response *DescribeClusterEndpointsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterEndpointsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterEndpoints require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterEndpointsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterInstancesRequest() (request *DescribeClusterInstancesRequest) {
    request = &DescribeClusterInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterInstances")
    
    
    return
}

func NewDescribeClusterInstancesResponse() (response *DescribeClusterInstancesResponse) {
    response = &DescribeClusterInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClusterInstances
// 查询集群下节点实例信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_INITMASTERFAILED = "InternalError.InitMasterFailed"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_PUBLICCLUSTEROPNOTSUPPORT = "InternalError.PublicClusterOpNotSupport"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeClusterInstances(request *DescribeClusterInstancesRequest) (response *DescribeClusterInstancesResponse, err error) {
    return c.DescribeClusterInstancesWithContext(context.Background(), request)
}

// DescribeClusterInstances
// 查询集群下节点实例信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_INITMASTERFAILED = "InternalError.InitMasterFailed"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_PUBLICCLUSTEROPNOTSUPPORT = "InternalError.PublicClusterOpNotSupport"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeClusterInstancesWithContext(ctx context.Context, request *DescribeClusterInstancesRequest) (response *DescribeClusterInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeClusterInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterKubeconfigRequest() (request *DescribeClusterKubeconfigRequest) {
    request = &DescribeClusterKubeconfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterKubeconfig")
    
    
    return
}

func NewDescribeClusterKubeconfigResponse() (response *DescribeClusterKubeconfigResponse) {
    response = &DescribeClusterKubeconfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClusterKubeconfig
// 获取集群的kubeconfig文件，不同子账户获取自己的kubeconfig文件，该文件中有每个子账户自己的kube-apiserver的客户端证书，默认首次调此接口时候创建客户端证书，时效20年，未授予任何权限，如果是集群所有者或者主账户，则默认是cluster-admin权限。
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  FAILEDOPERATION_KUBERNETESCLIENTBUILDERROR = "FailedOperation.KubernetesClientBuildError"
//  FAILEDOPERATION_UNEXPECTEDERROR = "FailedOperation.UnexpectedError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_COMPONENTCLIENTHTTP = "InternalError.ComponentClientHttp"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_KUBECLIENTCONNECTION = "InternalError.KubeClientConnection"
//  INTERNALERROR_KUBERNETESCLIENTBUILDERROR = "InternalError.KubernetesClientBuildError"
//  INTERNALERROR_KUBERNETESCREATEOPERATIONERROR = "InternalError.KubernetesCreateOperationError"
//  INTERNALERROR_KUBERNETESDELETEOPERATIONERROR = "InternalError.KubernetesDeleteOperationError"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INTERNALERROR_WHITELISTUNEXPECTEDERROR = "InternalError.WhitelistUnexpectedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_KUBERNETESRESOURCENOTFOUND = "ResourceNotFound.KubernetesResourceNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERINABNORMALSTAT = "ResourceUnavailable.ClusterInAbnormalStat"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeClusterKubeconfig(request *DescribeClusterKubeconfigRequest) (response *DescribeClusterKubeconfigResponse, err error) {
    return c.DescribeClusterKubeconfigWithContext(context.Background(), request)
}

// DescribeClusterKubeconfig
// 获取集群的kubeconfig文件，不同子账户获取自己的kubeconfig文件，该文件中有每个子账户自己的kube-apiserver的客户端证书，默认首次调此接口时候创建客户端证书，时效20年，未授予任何权限，如果是集群所有者或者主账户，则默认是cluster-admin权限。
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  FAILEDOPERATION_KUBERNETESCLIENTBUILDERROR = "FailedOperation.KubernetesClientBuildError"
//  FAILEDOPERATION_UNEXPECTEDERROR = "FailedOperation.UnexpectedError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_COMPONENTCLIENTHTTP = "InternalError.ComponentClientHttp"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_KUBECLIENTCONNECTION = "InternalError.KubeClientConnection"
//  INTERNALERROR_KUBERNETESCLIENTBUILDERROR = "InternalError.KubernetesClientBuildError"
//  INTERNALERROR_KUBERNETESCREATEOPERATIONERROR = "InternalError.KubernetesCreateOperationError"
//  INTERNALERROR_KUBERNETESDELETEOPERATIONERROR = "InternalError.KubernetesDeleteOperationError"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INTERNALERROR_WHITELISTUNEXPECTEDERROR = "InternalError.WhitelistUnexpectedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_KUBERNETESRESOURCENOTFOUND = "ResourceNotFound.KubernetesResourceNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERINABNORMALSTAT = "ResourceUnavailable.ClusterInAbnormalStat"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeClusterKubeconfigWithContext(ctx context.Context, request *DescribeClusterKubeconfigRequest) (response *DescribeClusterKubeconfigResponse, err error) {
    if request == nil {
        request = NewDescribeClusterKubeconfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterKubeconfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterKubeconfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterLevelAttributeRequest() (request *DescribeClusterLevelAttributeRequest) {
    request = &DescribeClusterLevelAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterLevelAttribute")
    
    
    return
}

func NewDescribeClusterLevelAttributeResponse() (response *DescribeClusterLevelAttributeResponse) {
    response = &DescribeClusterLevelAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClusterLevelAttribute
// 获取集群规模
//
// 可能返回的错误码:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeClusterLevelAttribute(request *DescribeClusterLevelAttributeRequest) (response *DescribeClusterLevelAttributeResponse, err error) {
    return c.DescribeClusterLevelAttributeWithContext(context.Background(), request)
}

// DescribeClusterLevelAttribute
// 获取集群规模
//
// 可能返回的错误码:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeClusterLevelAttributeWithContext(ctx context.Context, request *DescribeClusterLevelAttributeRequest) (response *DescribeClusterLevelAttributeResponse, err error) {
    if request == nil {
        request = NewDescribeClusterLevelAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterLevelAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterLevelAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterLevelChangeRecordsRequest() (request *DescribeClusterLevelChangeRecordsRequest) {
    request = &DescribeClusterLevelChangeRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterLevelChangeRecords")
    
    
    return
}

func NewDescribeClusterLevelChangeRecordsResponse() (response *DescribeClusterLevelChangeRecordsResponse) {
    response = &DescribeClusterLevelChangeRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClusterLevelChangeRecords
// 查询集群变配记录
//
// 可能返回的错误码:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeClusterLevelChangeRecords(request *DescribeClusterLevelChangeRecordsRequest) (response *DescribeClusterLevelChangeRecordsResponse, err error) {
    return c.DescribeClusterLevelChangeRecordsWithContext(context.Background(), request)
}

// DescribeClusterLevelChangeRecords
// 查询集群变配记录
//
// 可能返回的错误码:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeClusterLevelChangeRecordsWithContext(ctx context.Context, request *DescribeClusterLevelChangeRecordsRequest) (response *DescribeClusterLevelChangeRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterLevelChangeRecordsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterLevelChangeRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterLevelChangeRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterNodePoolDetailRequest() (request *DescribeClusterNodePoolDetailRequest) {
    request = &DescribeClusterNodePoolDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterNodePoolDetail")
    
    
    return
}

func NewDescribeClusterNodePoolDetailResponse() (response *DescribeClusterNodePoolDetailResponse) {
    response = &DescribeClusterNodePoolDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClusterNodePoolDetail
// 查询节点池详情
//
// 可能返回的错误码:
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
func (c *Client) DescribeClusterNodePoolDetail(request *DescribeClusterNodePoolDetailRequest) (response *DescribeClusterNodePoolDetailResponse, err error) {
    return c.DescribeClusterNodePoolDetailWithContext(context.Background(), request)
}

// DescribeClusterNodePoolDetail
// 查询节点池详情
//
// 可能返回的错误码:
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
func (c *Client) DescribeClusterNodePoolDetailWithContext(ctx context.Context, request *DescribeClusterNodePoolDetailRequest) (response *DescribeClusterNodePoolDetailResponse, err error) {
    if request == nil {
        request = NewDescribeClusterNodePoolDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterNodePoolDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterNodePoolDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterNodePoolsRequest() (request *DescribeClusterNodePoolsRequest) {
    request = &DescribeClusterNodePoolsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterNodePools")
    
    
    return
}

func NewDescribeClusterNodePoolsResponse() (response *DescribeClusterNodePoolsResponse) {
    response = &DescribeClusterNodePoolsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClusterNodePools
// 查询节点池列表
//
// 可能返回的错误码:
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
func (c *Client) DescribeClusterNodePools(request *DescribeClusterNodePoolsRequest) (response *DescribeClusterNodePoolsResponse, err error) {
    return c.DescribeClusterNodePoolsWithContext(context.Background(), request)
}

// DescribeClusterNodePools
// 查询节点池列表
//
// 可能返回的错误码:
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
func (c *Client) DescribeClusterNodePoolsWithContext(ctx context.Context, request *DescribeClusterNodePoolsRequest) (response *DescribeClusterNodePoolsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterNodePoolsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterNodePools require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterNodePoolsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterPendingReleasesRequest() (request *DescribeClusterPendingReleasesRequest) {
    request = &DescribeClusterPendingReleasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterPendingReleases")
    
    
    return
}

func NewDescribeClusterPendingReleasesResponse() (response *DescribeClusterPendingReleasesResponse) {
    response = &DescribeClusterPendingReleasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClusterPendingReleases
// 在应用市场中查询正在安装中的应用列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MARKETGETAUTHFAILED = "FailedOperation.MarketGetAuthFailed"
//  FAILEDOPERATION_MARKETRELEASEOPERATION = "FailedOperation.MarketReleaseOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MARKETGETAUTHFAILED = "InternalError.MarketGetAuthFailed"
//  INTERNALERROR_MARKETINTERNALSERVERERROR = "InternalError.MarketInternalServerError"
//  INTERNALERROR_MARKETRELEASEOPERATION = "InternalError.MarketReleaseOperation"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeClusterPendingReleases(request *DescribeClusterPendingReleasesRequest) (response *DescribeClusterPendingReleasesResponse, err error) {
    return c.DescribeClusterPendingReleasesWithContext(context.Background(), request)
}

// DescribeClusterPendingReleases
// 在应用市场中查询正在安装中的应用列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MARKETGETAUTHFAILED = "FailedOperation.MarketGetAuthFailed"
//  FAILEDOPERATION_MARKETRELEASEOPERATION = "FailedOperation.MarketReleaseOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MARKETGETAUTHFAILED = "InternalError.MarketGetAuthFailed"
//  INTERNALERROR_MARKETINTERNALSERVERERROR = "InternalError.MarketInternalServerError"
//  INTERNALERROR_MARKETRELEASEOPERATION = "InternalError.MarketReleaseOperation"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeClusterPendingReleasesWithContext(ctx context.Context, request *DescribeClusterPendingReleasesRequest) (response *DescribeClusterPendingReleasesResponse, err error) {
    if request == nil {
        request = NewDescribeClusterPendingReleasesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterPendingReleases require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterPendingReleasesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterReleaseDetailsRequest() (request *DescribeClusterReleaseDetailsRequest) {
    request = &DescribeClusterReleaseDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterReleaseDetails")
    
    
    return
}

func NewDescribeClusterReleaseDetailsResponse() (response *DescribeClusterReleaseDetailsResponse) {
    response = &DescribeClusterReleaseDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClusterReleaseDetails
// 查询通过应用市场安装的某个应用详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MARKETGETAUTHFAILED = "FailedOperation.MarketGetAuthFailed"
//  FAILEDOPERATION_MARKETRELEASEOPERATION = "FailedOperation.MarketReleaseOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MARKETGETAUTHFAILED = "InternalError.MarketGetAuthFailed"
//  INTERNALERROR_MARKETINTERNALSERVERERROR = "InternalError.MarketInternalServerError"
//  INTERNALERROR_MARKETRELEASEOPERATION = "InternalError.MarketReleaseOperation"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeClusterReleaseDetails(request *DescribeClusterReleaseDetailsRequest) (response *DescribeClusterReleaseDetailsResponse, err error) {
    return c.DescribeClusterReleaseDetailsWithContext(context.Background(), request)
}

// DescribeClusterReleaseDetails
// 查询通过应用市场安装的某个应用详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MARKETGETAUTHFAILED = "FailedOperation.MarketGetAuthFailed"
//  FAILEDOPERATION_MARKETRELEASEOPERATION = "FailedOperation.MarketReleaseOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MARKETGETAUTHFAILED = "InternalError.MarketGetAuthFailed"
//  INTERNALERROR_MARKETINTERNALSERVERERROR = "InternalError.MarketInternalServerError"
//  INTERNALERROR_MARKETRELEASEOPERATION = "InternalError.MarketReleaseOperation"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeClusterReleaseDetailsWithContext(ctx context.Context, request *DescribeClusterReleaseDetailsRequest) (response *DescribeClusterReleaseDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterReleaseDetailsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterReleaseDetails require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterReleaseDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterReleaseHistoryRequest() (request *DescribeClusterReleaseHistoryRequest) {
    request = &DescribeClusterReleaseHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterReleaseHistory")
    
    
    return
}

func NewDescribeClusterReleaseHistoryResponse() (response *DescribeClusterReleaseHistoryResponse) {
    response = &DescribeClusterReleaseHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClusterReleaseHistory
// 查询集群在应用市场中某个已安装应用的版本历史
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MARKETGETAUTHFAILED = "FailedOperation.MarketGetAuthFailed"
//  FAILEDOPERATION_MARKETRELEASEOPERATION = "FailedOperation.MarketReleaseOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MARKETGETAUTHFAILED = "InternalError.MarketGetAuthFailed"
//  INTERNALERROR_MARKETINTERNALSERVERERROR = "InternalError.MarketInternalServerError"
//  INTERNALERROR_MARKETRELEASEOPERATION = "InternalError.MarketReleaseOperation"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeClusterReleaseHistory(request *DescribeClusterReleaseHistoryRequest) (response *DescribeClusterReleaseHistoryResponse, err error) {
    return c.DescribeClusterReleaseHistoryWithContext(context.Background(), request)
}

// DescribeClusterReleaseHistory
// 查询集群在应用市场中某个已安装应用的版本历史
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MARKETGETAUTHFAILED = "FailedOperation.MarketGetAuthFailed"
//  FAILEDOPERATION_MARKETRELEASEOPERATION = "FailedOperation.MarketReleaseOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MARKETGETAUTHFAILED = "InternalError.MarketGetAuthFailed"
//  INTERNALERROR_MARKETINTERNALSERVERERROR = "InternalError.MarketInternalServerError"
//  INTERNALERROR_MARKETRELEASEOPERATION = "InternalError.MarketReleaseOperation"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeClusterReleaseHistoryWithContext(ctx context.Context, request *DescribeClusterReleaseHistoryRequest) (response *DescribeClusterReleaseHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeClusterReleaseHistoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterReleaseHistory require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterReleaseHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterReleasesRequest() (request *DescribeClusterReleasesRequest) {
    request = &DescribeClusterReleasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterReleases")
    
    
    return
}

func NewDescribeClusterReleasesResponse() (response *DescribeClusterReleasesResponse) {
    response = &DescribeClusterReleasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClusterReleases
// 查询集群在应用市场中已安装应用列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MARKETGETAUTHFAILED = "FailedOperation.MarketGetAuthFailed"
//  FAILEDOPERATION_MARKETRELEASEOPERATION = "FailedOperation.MarketReleaseOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MARKETGETAUTHFAILED = "InternalError.MarketGetAuthFailed"
//  INTERNALERROR_MARKETINTERNALSERVERERROR = "InternalError.MarketInternalServerError"
//  INTERNALERROR_MARKETRELEASEOPERATION = "InternalError.MarketReleaseOperation"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeClusterReleases(request *DescribeClusterReleasesRequest) (response *DescribeClusterReleasesResponse, err error) {
    return c.DescribeClusterReleasesWithContext(context.Background(), request)
}

// DescribeClusterReleases
// 查询集群在应用市场中已安装应用列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MARKETGETAUTHFAILED = "FailedOperation.MarketGetAuthFailed"
//  FAILEDOPERATION_MARKETRELEASEOPERATION = "FailedOperation.MarketReleaseOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MARKETGETAUTHFAILED = "InternalError.MarketGetAuthFailed"
//  INTERNALERROR_MARKETINTERNALSERVERERROR = "InternalError.MarketInternalServerError"
//  INTERNALERROR_MARKETRELEASEOPERATION = "InternalError.MarketReleaseOperation"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeClusterReleasesWithContext(ctx context.Context, request *DescribeClusterReleasesRequest) (response *DescribeClusterReleasesResponse, err error) {
    if request == nil {
        request = NewDescribeClusterReleasesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterReleases require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterReleasesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterRouteTablesRequest() (request *DescribeClusterRouteTablesRequest) {
    request = &DescribeClusterRouteTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterRouteTables")
    
    
    return
}

func NewDescribeClusterRouteTablesResponse() (response *DescribeClusterRouteTablesResponse) {
    response = &DescribeClusterRouteTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClusterRouteTables
// 查询集群路由表
//
// 可能返回的错误码:
//  INTERNALERROR_DB = "InternalError.Db"
func (c *Client) DescribeClusterRouteTables(request *DescribeClusterRouteTablesRequest) (response *DescribeClusterRouteTablesResponse, err error) {
    return c.DescribeClusterRouteTablesWithContext(context.Background(), request)
}

// DescribeClusterRouteTables
// 查询集群路由表
//
// 可能返回的错误码:
//  INTERNALERROR_DB = "InternalError.Db"
func (c *Client) DescribeClusterRouteTablesWithContext(ctx context.Context, request *DescribeClusterRouteTablesRequest) (response *DescribeClusterRouteTablesResponse, err error) {
    if request == nil {
        request = NewDescribeClusterRouteTablesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterRouteTables require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterRouteTablesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterRoutesRequest() (request *DescribeClusterRoutesRequest) {
    request = &DescribeClusterRoutesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterRoutes")
    
    
    return
}

func NewDescribeClusterRoutesResponse() (response *DescribeClusterRoutesResponse) {
    response = &DescribeClusterRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClusterRoutes
// 查询集群路由
//
// 可能返回的错误码:
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeClusterRoutes(request *DescribeClusterRoutesRequest) (response *DescribeClusterRoutesResponse, err error) {
    return c.DescribeClusterRoutesWithContext(context.Background(), request)
}

// DescribeClusterRoutes
// 查询集群路由
//
// 可能返回的错误码:
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeClusterRoutesWithContext(ctx context.Context, request *DescribeClusterRoutesRequest) (response *DescribeClusterRoutesResponse, err error) {
    if request == nil {
        request = NewDescribeClusterRoutesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterRoutes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterRoutesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterSecurityRequest() (request *DescribeClusterSecurityRequest) {
    request = &DescribeClusterSecurityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterSecurity")
    
    
    return
}

func NewDescribeClusterSecurityResponse() (response *DescribeClusterSecurityResponse) {
    response = &DescribeClusterSecurityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClusterSecurity
// 集群的密钥信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  FAILEDOPERATION_DBRECORDNOTFOUND = "FailedOperation.DbRecordNotFound"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  FAILEDOPERATION_LBCOMMON = "FailedOperation.LbCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_COMPONENTCLIENTHTTP = "InternalError.ComponentClientHttp"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_KUBECOMMON = "InternalError.KubeCommon"
//  INTERNALERROR_LBCOMMON = "InternalError.LbCommon"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_CIDRINVALID = "InvalidParameter.CidrInvalid"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_KUBERESOURCENOTFOUND = "ResourceNotFound.KubeResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CLUSTERINABNORMALSTAT = "ResourceUnavailable.ClusterInAbnormalStat"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeClusterSecurity(request *DescribeClusterSecurityRequest) (response *DescribeClusterSecurityResponse, err error) {
    return c.DescribeClusterSecurityWithContext(context.Background(), request)
}

// DescribeClusterSecurity
// 集群的密钥信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  FAILEDOPERATION_DBRECORDNOTFOUND = "FailedOperation.DbRecordNotFound"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  FAILEDOPERATION_LBCOMMON = "FailedOperation.LbCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_COMPONENTCLIENTHTTP = "InternalError.ComponentClientHttp"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_KUBECOMMON = "InternalError.KubeCommon"
//  INTERNALERROR_LBCOMMON = "InternalError.LbCommon"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_CIDRINVALID = "InvalidParameter.CidrInvalid"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_KUBERESOURCENOTFOUND = "ResourceNotFound.KubeResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CLUSTERINABNORMALSTAT = "ResourceUnavailable.ClusterInAbnormalStat"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeClusterSecurityWithContext(ctx context.Context, request *DescribeClusterSecurityRequest) (response *DescribeClusterSecurityResponse, err error) {
    if request == nil {
        request = NewDescribeClusterSecurityRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterSecurity require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterSecurityResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterStatusRequest() (request *DescribeClusterStatusRequest) {
    request = &DescribeClusterStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterStatus")
    
    
    return
}

func NewDescribeClusterStatusResponse() (response *DescribeClusterStatusResponse) {
    response = &DescribeClusterStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClusterStatus
// 查看集群状态列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CAMNOAUTH = "FailedOperation.CamNoAuth"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_COMPONENTCLINETHTTP = "FailedOperation.ComponentClinetHttp"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_COMPONENTCLINETHTTP = "InternalError.ComponentClinetHttp"
//  INTERNALERROR_CVMCOMMON = "InternalError.CvmCommon"
//  INTERNALERROR_CVMNOTFOUND = "InternalError.CvmNotFound"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeClusterStatus(request *DescribeClusterStatusRequest) (response *DescribeClusterStatusResponse, err error) {
    return c.DescribeClusterStatusWithContext(context.Background(), request)
}

// DescribeClusterStatus
// 查看集群状态列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CAMNOAUTH = "FailedOperation.CamNoAuth"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_COMPONENTCLINETHTTP = "FailedOperation.ComponentClinetHttp"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_COMPONENTCLINETHTTP = "InternalError.ComponentClinetHttp"
//  INTERNALERROR_CVMCOMMON = "InternalError.CvmCommon"
//  INTERNALERROR_CVMNOTFOUND = "InternalError.CvmNotFound"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeClusterStatusWithContext(ctx context.Context, request *DescribeClusterStatusRequest) (response *DescribeClusterStatusResponse, err error) {
    if request == nil {
        request = NewDescribeClusterStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterVirtualNodeRequest() (request *DescribeClusterVirtualNodeRequest) {
    request = &DescribeClusterVirtualNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterVirtualNode")
    
    
    return
}

func NewDescribeClusterVirtualNodeResponse() (response *DescribeClusterVirtualNodeResponse) {
    response = &DescribeClusterVirtualNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClusterVirtualNode
// 查看虚拟节点列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_KUBERNETESLISTOPERATIONERROR = "FailedOperation.KubernetesListOperationError"
//  INTERNALERROR_KUBERNETESLISTOPERATIONERROR = "InternalError.KubernetesListOperationError"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) DescribeClusterVirtualNode(request *DescribeClusterVirtualNodeRequest) (response *DescribeClusterVirtualNodeResponse, err error) {
    return c.DescribeClusterVirtualNodeWithContext(context.Background(), request)
}

// DescribeClusterVirtualNode
// 查看虚拟节点列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_KUBERNETESLISTOPERATIONERROR = "FailedOperation.KubernetesListOperationError"
//  INTERNALERROR_KUBERNETESLISTOPERATIONERROR = "InternalError.KubernetesListOperationError"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) DescribeClusterVirtualNodeWithContext(ctx context.Context, request *DescribeClusterVirtualNodeRequest) (response *DescribeClusterVirtualNodeResponse, err error) {
    if request == nil {
        request = NewDescribeClusterVirtualNodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterVirtualNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterVirtualNodeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterVirtualNodePoolsRequest() (request *DescribeClusterVirtualNodePoolsRequest) {
    request = &DescribeClusterVirtualNodePoolsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterVirtualNodePools")
    
    
    return
}

func NewDescribeClusterVirtualNodePoolsResponse() (response *DescribeClusterVirtualNodePoolsResponse) {
    response = &DescribeClusterVirtualNodePoolsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClusterVirtualNodePools
// 查看虚拟节点池列表
//
// 可能返回的错误码:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) DescribeClusterVirtualNodePools(request *DescribeClusterVirtualNodePoolsRequest) (response *DescribeClusterVirtualNodePoolsResponse, err error) {
    return c.DescribeClusterVirtualNodePoolsWithContext(context.Background(), request)
}

// DescribeClusterVirtualNodePools
// 查看虚拟节点池列表
//
// 可能返回的错误码:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) DescribeClusterVirtualNodePoolsWithContext(ctx context.Context, request *DescribeClusterVirtualNodePoolsRequest) (response *DescribeClusterVirtualNodePoolsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterVirtualNodePoolsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterVirtualNodePools require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterVirtualNodePoolsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClustersRequest() (request *DescribeClustersRequest) {
    request = &DescribeClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusters")
    
    
    return
}

func NewDescribeClustersResponse() (response *DescribeClustersResponse) {
    response = &DescribeClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClusters
// 查询集群列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNEXPECTEDERROR = "FailedOperation.UnexpectedError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_PUBLICCLUSTEROPNOTSUPPORT = "InternalError.PublicClusterOpNotSupport"
//  INTERNALERROR_QUOTAMAXCLSLIMIT = "InternalError.QuotaMaxClsLimit"
//  INTERNALERROR_QUOTAMAXNODLIMIT = "InternalError.QuotaMaxNodLimit"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeClusters(request *DescribeClustersRequest) (response *DescribeClustersResponse, err error) {
    return c.DescribeClustersWithContext(context.Background(), request)
}

// DescribeClusters
// 查询集群列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNEXPECTEDERROR = "FailedOperation.UnexpectedError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_PUBLICCLUSTEROPNOTSUPPORT = "InternalError.PublicClusterOpNotSupport"
//  INTERNALERROR_QUOTAMAXCLSLIMIT = "InternalError.QuotaMaxClsLimit"
//  INTERNALERROR_QUOTAMAXNODLIMIT = "InternalError.QuotaMaxNodLimit"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeClustersWithContext(ctx context.Context, request *DescribeClustersRequest) (response *DescribeClustersResponse, err error) {
    if request == nil {
        request = NewDescribeClustersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusters require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClustersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeECMInstancesRequest() (request *DescribeECMInstancesRequest) {
    request = &DescribeECMInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeECMInstances")
    
    
    return
}

func NewDescribeECMInstancesResponse() (response *DescribeECMInstancesResponse) {
    response = &DescribeECMInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeECMInstances
// 获取ECM实例相关信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeECMInstances(request *DescribeECMInstancesRequest) (response *DescribeECMInstancesResponse, err error) {
    return c.DescribeECMInstancesWithContext(context.Background(), request)
}

// DescribeECMInstances
// 获取ECM实例相关信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeECMInstancesWithContext(ctx context.Context, request *DescribeECMInstancesRequest) (response *DescribeECMInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeECMInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeECMInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeECMInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEKSClusterCredentialRequest() (request *DescribeEKSClusterCredentialRequest) {
    request = &DescribeEKSClusterCredentialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeEKSClusterCredential")
    
    
    return
}

func NewDescribeEKSClusterCredentialResponse() (response *DescribeEKSClusterCredentialResponse) {
    response = &DescribeEKSClusterCredentialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEKSClusterCredential
// 获取弹性容器集群的接入认证信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEKSClusterCredential(request *DescribeEKSClusterCredentialRequest) (response *DescribeEKSClusterCredentialResponse, err error) {
    return c.DescribeEKSClusterCredentialWithContext(context.Background(), request)
}

// DescribeEKSClusterCredential
// 获取弹性容器集群的接入认证信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEKSClusterCredentialWithContext(ctx context.Context, request *DescribeEKSClusterCredentialRequest) (response *DescribeEKSClusterCredentialResponse, err error) {
    if request == nil {
        request = NewDescribeEKSClusterCredentialRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEKSClusterCredential require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEKSClusterCredentialResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEKSClustersRequest() (request *DescribeEKSClustersRequest) {
    request = &DescribeEKSClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeEKSClusters")
    
    
    return
}

func NewDescribeEKSClustersResponse() (response *DescribeEKSClustersResponse) {
    response = &DescribeEKSClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEKSClusters
// 查询弹性集群列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEKSClusters(request *DescribeEKSClustersRequest) (response *DescribeEKSClustersResponse, err error) {
    return c.DescribeEKSClustersWithContext(context.Background(), request)
}

// DescribeEKSClusters
// 查询弹性集群列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEKSClustersWithContext(ctx context.Context, request *DescribeEKSClustersRequest) (response *DescribeEKSClustersResponse, err error) {
    if request == nil {
        request = NewDescribeEKSClustersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEKSClusters require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEKSClustersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEKSContainerInstanceEventRequest() (request *DescribeEKSContainerInstanceEventRequest) {
    request = &DescribeEKSContainerInstanceEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeEKSContainerInstanceEvent")
    
    
    return
}

func NewDescribeEKSContainerInstanceEventResponse() (response *DescribeEKSContainerInstanceEventResponse) {
    response = &DescribeEKSContainerInstanceEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEKSContainerInstanceEvent
// 查询容器实例的事件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEKSContainerInstanceEvent(request *DescribeEKSContainerInstanceEventRequest) (response *DescribeEKSContainerInstanceEventResponse, err error) {
    return c.DescribeEKSContainerInstanceEventWithContext(context.Background(), request)
}

// DescribeEKSContainerInstanceEvent
// 查询容器实例的事件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEKSContainerInstanceEventWithContext(ctx context.Context, request *DescribeEKSContainerInstanceEventRequest) (response *DescribeEKSContainerInstanceEventResponse, err error) {
    if request == nil {
        request = NewDescribeEKSContainerInstanceEventRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEKSContainerInstanceEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEKSContainerInstanceEventResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEKSContainerInstanceRegionsRequest() (request *DescribeEKSContainerInstanceRegionsRequest) {
    request = &DescribeEKSContainerInstanceRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeEKSContainerInstanceRegions")
    
    
    return
}

func NewDescribeEKSContainerInstanceRegionsResponse() (response *DescribeEKSContainerInstanceRegionsResponse) {
    response = &DescribeEKSContainerInstanceRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEKSContainerInstanceRegions
// 查询容器实例支持的地域
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeEKSContainerInstanceRegions(request *DescribeEKSContainerInstanceRegionsRequest) (response *DescribeEKSContainerInstanceRegionsResponse, err error) {
    return c.DescribeEKSContainerInstanceRegionsWithContext(context.Background(), request)
}

// DescribeEKSContainerInstanceRegions
// 查询容器实例支持的地域
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeEKSContainerInstanceRegionsWithContext(ctx context.Context, request *DescribeEKSContainerInstanceRegionsRequest) (response *DescribeEKSContainerInstanceRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeEKSContainerInstanceRegionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEKSContainerInstanceRegions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEKSContainerInstanceRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEKSContainerInstancesRequest() (request *DescribeEKSContainerInstancesRequest) {
    request = &DescribeEKSContainerInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeEKSContainerInstances")
    
    
    return
}

func NewDescribeEKSContainerInstancesResponse() (response *DescribeEKSContainerInstancesResponse) {
    response = &DescribeEKSContainerInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEKSContainerInstances
// 查询容器实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RBACFORBIDDEN = "FailedOperation.RBACForbidden"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeEKSContainerInstances(request *DescribeEKSContainerInstancesRequest) (response *DescribeEKSContainerInstancesResponse, err error) {
    return c.DescribeEKSContainerInstancesWithContext(context.Background(), request)
}

// DescribeEKSContainerInstances
// 查询容器实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RBACFORBIDDEN = "FailedOperation.RBACForbidden"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeEKSContainerInstancesWithContext(ctx context.Context, request *DescribeEKSContainerInstancesRequest) (response *DescribeEKSContainerInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeEKSContainerInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEKSContainerInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEKSContainerInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgeAvailableExtraArgsRequest() (request *DescribeEdgeAvailableExtraArgsRequest) {
    request = &DescribeEdgeAvailableExtraArgsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeEdgeAvailableExtraArgs")
    
    
    return
}

func NewDescribeEdgeAvailableExtraArgsResponse() (response *DescribeEdgeAvailableExtraArgsResponse) {
    response = &DescribeEdgeAvailableExtraArgsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEdgeAvailableExtraArgs
// 查询边缘容器集群可用的自定义参数
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) DescribeEdgeAvailableExtraArgs(request *DescribeEdgeAvailableExtraArgsRequest) (response *DescribeEdgeAvailableExtraArgsResponse, err error) {
    return c.DescribeEdgeAvailableExtraArgsWithContext(context.Background(), request)
}

// DescribeEdgeAvailableExtraArgs
// 查询边缘容器集群可用的自定义参数
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) DescribeEdgeAvailableExtraArgsWithContext(ctx context.Context, request *DescribeEdgeAvailableExtraArgsRequest) (response *DescribeEdgeAvailableExtraArgsResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeAvailableExtraArgsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgeAvailableExtraArgs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgeAvailableExtraArgsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgeCVMInstancesRequest() (request *DescribeEdgeCVMInstancesRequest) {
    request = &DescribeEdgeCVMInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeEdgeCVMInstances")
    
    
    return
}

func NewDescribeEdgeCVMInstancesResponse() (response *DescribeEdgeCVMInstancesResponse) {
    response = &DescribeEdgeCVMInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEdgeCVMInstances
// 获取边缘容器CVM实例相关信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeCVMInstances(request *DescribeEdgeCVMInstancesRequest) (response *DescribeEdgeCVMInstancesResponse, err error) {
    return c.DescribeEdgeCVMInstancesWithContext(context.Background(), request)
}

// DescribeEdgeCVMInstances
// 获取边缘容器CVM实例相关信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeCVMInstancesWithContext(ctx context.Context, request *DescribeEdgeCVMInstancesRequest) (response *DescribeEdgeCVMInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeCVMInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgeCVMInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgeCVMInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgeClusterExtraArgsRequest() (request *DescribeEdgeClusterExtraArgsRequest) {
    request = &DescribeEdgeClusterExtraArgsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeEdgeClusterExtraArgs")
    
    
    return
}

func NewDescribeEdgeClusterExtraArgsResponse() (response *DescribeEdgeClusterExtraArgsResponse) {
    response = &DescribeEdgeClusterExtraArgsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEdgeClusterExtraArgs
// 查询边缘集群自定义参数
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeClusterExtraArgs(request *DescribeEdgeClusterExtraArgsRequest) (response *DescribeEdgeClusterExtraArgsResponse, err error) {
    return c.DescribeEdgeClusterExtraArgsWithContext(context.Background(), request)
}

// DescribeEdgeClusterExtraArgs
// 查询边缘集群自定义参数
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeClusterExtraArgsWithContext(ctx context.Context, request *DescribeEdgeClusterExtraArgsRequest) (response *DescribeEdgeClusterExtraArgsResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeClusterExtraArgsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgeClusterExtraArgs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgeClusterExtraArgsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgeClusterInstancesRequest() (request *DescribeEdgeClusterInstancesRequest) {
    request = &DescribeEdgeClusterInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeEdgeClusterInstances")
    
    
    return
}

func NewDescribeEdgeClusterInstancesResponse() (response *DescribeEdgeClusterInstancesResponse) {
    response = &DescribeEdgeClusterInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEdgeClusterInstances
// 查询边缘计算集群的节点信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeClusterInstances(request *DescribeEdgeClusterInstancesRequest) (response *DescribeEdgeClusterInstancesResponse, err error) {
    return c.DescribeEdgeClusterInstancesWithContext(context.Background(), request)
}

// DescribeEdgeClusterInstances
// 查询边缘计算集群的节点信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeClusterInstancesWithContext(ctx context.Context, request *DescribeEdgeClusterInstancesRequest) (response *DescribeEdgeClusterInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeClusterInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgeClusterInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgeClusterInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgeClusterUpgradeInfoRequest() (request *DescribeEdgeClusterUpgradeInfoRequest) {
    request = &DescribeEdgeClusterUpgradeInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeEdgeClusterUpgradeInfo")
    
    
    return
}

func NewDescribeEdgeClusterUpgradeInfoResponse() (response *DescribeEdgeClusterUpgradeInfoResponse) {
    response = &DescribeEdgeClusterUpgradeInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEdgeClusterUpgradeInfo
// 可以查询边缘集群升级信息，包含可以升级的组件，当前升级状态和升级错误信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeEdgeClusterUpgradeInfo(request *DescribeEdgeClusterUpgradeInfoRequest) (response *DescribeEdgeClusterUpgradeInfoResponse, err error) {
    return c.DescribeEdgeClusterUpgradeInfoWithContext(context.Background(), request)
}

// DescribeEdgeClusterUpgradeInfo
// 可以查询边缘集群升级信息，包含可以升级的组件，当前升级状态和升级错误信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeEdgeClusterUpgradeInfoWithContext(ctx context.Context, request *DescribeEdgeClusterUpgradeInfoRequest) (response *DescribeEdgeClusterUpgradeInfoResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeClusterUpgradeInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgeClusterUpgradeInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgeClusterUpgradeInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgeLogSwitchesRequest() (request *DescribeEdgeLogSwitchesRequest) {
    request = &DescribeEdgeLogSwitchesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeEdgeLogSwitches")
    
    
    return
}

func NewDescribeEdgeLogSwitchesResponse() (response *DescribeEdgeLogSwitchesResponse) {
    response = &DescribeEdgeLogSwitchesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEdgeLogSwitches
// 获取事件、审计和日志的状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RBACFORBIDDEN = "FailedOperation.RBACForbidden"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_KUBERNETESCLIENTBUILDERROR = "InternalError.KubernetesClientBuildError"
//  INTERNALERROR_KUBERNETESGETOPERATIONERROR = "InternalError.KubernetesGetOperationError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeLogSwitches(request *DescribeEdgeLogSwitchesRequest) (response *DescribeEdgeLogSwitchesResponse, err error) {
    return c.DescribeEdgeLogSwitchesWithContext(context.Background(), request)
}

// DescribeEdgeLogSwitches
// 获取事件、审计和日志的状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RBACFORBIDDEN = "FailedOperation.RBACForbidden"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_KUBERNETESCLIENTBUILDERROR = "InternalError.KubernetesClientBuildError"
//  INTERNALERROR_KUBERNETESGETOPERATIONERROR = "InternalError.KubernetesGetOperationError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeLogSwitchesWithContext(ctx context.Context, request *DescribeEdgeLogSwitchesRequest) (response *DescribeEdgeLogSwitchesResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeLogSwitchesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgeLogSwitches require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgeLogSwitchesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEksContainerInstanceLogRequest() (request *DescribeEksContainerInstanceLogRequest) {
    request = &DescribeEksContainerInstanceLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeEksContainerInstanceLog")
    
    
    return
}

func NewDescribeEksContainerInstanceLogResponse() (response *DescribeEksContainerInstanceLogResponse) {
    response = &DescribeEksContainerInstanceLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEksContainerInstanceLog
// 查询容器实例中容器日志
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CONTAINERNOTFOUND = "InternalError.ContainerNotFound"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_EKSCONTAINERSTATUS = "ResourceUnavailable.EksContainerStatus"
func (c *Client) DescribeEksContainerInstanceLog(request *DescribeEksContainerInstanceLogRequest) (response *DescribeEksContainerInstanceLogResponse, err error) {
    return c.DescribeEksContainerInstanceLogWithContext(context.Background(), request)
}

// DescribeEksContainerInstanceLog
// 查询容器实例中容器日志
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CONTAINERNOTFOUND = "InternalError.ContainerNotFound"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_EKSCONTAINERSTATUS = "ResourceUnavailable.EksContainerStatus"
func (c *Client) DescribeEksContainerInstanceLogWithContext(ctx context.Context, request *DescribeEksContainerInstanceLogRequest) (response *DescribeEksContainerInstanceLogResponse, err error) {
    if request == nil {
        request = NewDescribeEksContainerInstanceLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEksContainerInstanceLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEksContainerInstanceLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEnableVpcCniProgressRequest() (request *DescribeEnableVpcCniProgressRequest) {
    request = &DescribeEnableVpcCniProgressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeEnableVpcCniProgress")
    
    
    return
}

func NewDescribeEnableVpcCniProgressResponse() (response *DescribeEnableVpcCniProgressResponse) {
    response = &DescribeEnableVpcCniProgressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEnableVpcCniProgress
// 本接口用于查询开启vpc-cni模式的任务进度
//
// 可能返回的错误码:
//  FAILEDOPERATION_ENABLEVPCCNIFAILED = "FailedOperation.EnableVPCCNIFailed"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
func (c *Client) DescribeEnableVpcCniProgress(request *DescribeEnableVpcCniProgressRequest) (response *DescribeEnableVpcCniProgressResponse, err error) {
    return c.DescribeEnableVpcCniProgressWithContext(context.Background(), request)
}

// DescribeEnableVpcCniProgress
// 本接口用于查询开启vpc-cni模式的任务进度
//
// 可能返回的错误码:
//  FAILEDOPERATION_ENABLEVPCCNIFAILED = "FailedOperation.EnableVPCCNIFailed"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
func (c *Client) DescribeEnableVpcCniProgressWithContext(ctx context.Context, request *DescribeEnableVpcCniProgressRequest) (response *DescribeEnableVpcCniProgressResponse, err error) {
    if request == nil {
        request = NewDescribeEnableVpcCniProgressRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEnableVpcCniProgress require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEnableVpcCniProgressResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExistedInstancesRequest() (request *DescribeExistedInstancesRequest) {
    request = &DescribeExistedInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeExistedInstances")
    
    
    return
}

func NewDescribeExistedInstancesResponse() (response *DescribeExistedInstancesResponse) {
    response = &DescribeExistedInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeExistedInstances
// 查询已经存在的节点，判断是否可以加入集群
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ASCOMMON = "FailedOperation.AsCommon"
//  FAILEDOPERATION_CLUSTERNOTFOUND = "FailedOperation.ClusterNotFound"
//  FAILEDOPERATION_CLUSTERSTATE = "FailedOperation.ClusterState"
//  FAILEDOPERATION_CVMCOMMON = "FailedOperation.CvmCommon"
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_ASCOMMON = "InternalError.AsCommon"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_CREATEMASTERFAILED = "InternalError.CreateMasterFailed"
//  INTERNALERROR_CVMCOMMON = "InternalError.CvmCommon"
//  INTERNALERROR_CVMNOTFOUND = "InternalError.CvmNotFound"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_IMAGEIDNOTFOUND = "InternalError.ImageIdNotFound"
//  INTERNALERROR_INITMASTERFAILED = "InternalError.InitMasterFailed"
//  INTERNALERROR_INVALIDPRIVATENETWORKCIDR = "InternalError.InvalidPrivateNetworkCidr"
//  INTERNALERROR_OSNOTSUPPORT = "InternalError.OsNotSupport"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_VPCCOMMON = "InternalError.VpcCommon"
//  INTERNALERROR_VPCRECODRNOTFOUND = "InternalError.VpcRecodrNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeExistedInstances(request *DescribeExistedInstancesRequest) (response *DescribeExistedInstancesResponse, err error) {
    return c.DescribeExistedInstancesWithContext(context.Background(), request)
}

// DescribeExistedInstances
// 查询已经存在的节点，判断是否可以加入集群
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ASCOMMON = "FailedOperation.AsCommon"
//  FAILEDOPERATION_CLUSTERNOTFOUND = "FailedOperation.ClusterNotFound"
//  FAILEDOPERATION_CLUSTERSTATE = "FailedOperation.ClusterState"
//  FAILEDOPERATION_CVMCOMMON = "FailedOperation.CvmCommon"
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_ASCOMMON = "InternalError.AsCommon"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_CREATEMASTERFAILED = "InternalError.CreateMasterFailed"
//  INTERNALERROR_CVMCOMMON = "InternalError.CvmCommon"
//  INTERNALERROR_CVMNOTFOUND = "InternalError.CvmNotFound"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_IMAGEIDNOTFOUND = "InternalError.ImageIdNotFound"
//  INTERNALERROR_INITMASTERFAILED = "InternalError.InitMasterFailed"
//  INTERNALERROR_INVALIDPRIVATENETWORKCIDR = "InternalError.InvalidPrivateNetworkCidr"
//  INTERNALERROR_OSNOTSUPPORT = "InternalError.OsNotSupport"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_VPCCOMMON = "InternalError.VpcCommon"
//  INTERNALERROR_VPCRECODRNOTFOUND = "InternalError.VpcRecodrNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeExistedInstancesWithContext(ctx context.Context, request *DescribeExistedInstancesRequest) (response *DescribeExistedInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeExistedInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExistedInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExistedInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExternalClusterSpecRequest() (request *DescribeExternalClusterSpecRequest) {
    request = &DescribeExternalClusterSpecRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeExternalClusterSpec")
    
    
    return
}

func NewDescribeExternalClusterSpecResponse() (response *DescribeExternalClusterSpecResponse) {
    response = &DescribeExternalClusterSpecResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeExternalClusterSpec
// 获取导入第三方集群YAML定义
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_KUBERNETESCLIENTBUILDERROR = "InternalError.KubernetesClientBuildError"
//  INTERNALERROR_KUBERNETESCREATEOPERATIONERROR = "InternalError.KubernetesCreateOperationError"
//  INTERNALERROR_KUBERNETESDELETEOPERATIONERROR = "InternalError.KubernetesDeleteOperationError"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INTERNALERROR_WHITELISTUNEXPECTEDERROR = "InternalError.WhitelistUnexpectedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_KUBERNETESRESOURCENOTFOUND = "ResourceNotFound.KubernetesResourceNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeExternalClusterSpec(request *DescribeExternalClusterSpecRequest) (response *DescribeExternalClusterSpecResponse, err error) {
    return c.DescribeExternalClusterSpecWithContext(context.Background(), request)
}

// DescribeExternalClusterSpec
// 获取导入第三方集群YAML定义
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_KUBERNETESCLIENTBUILDERROR = "InternalError.KubernetesClientBuildError"
//  INTERNALERROR_KUBERNETESCREATEOPERATIONERROR = "InternalError.KubernetesCreateOperationError"
//  INTERNALERROR_KUBERNETESDELETEOPERATIONERROR = "InternalError.KubernetesDeleteOperationError"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INTERNALERROR_WHITELISTUNEXPECTEDERROR = "InternalError.WhitelistUnexpectedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_KUBERNETESRESOURCENOTFOUND = "ResourceNotFound.KubernetesResourceNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeExternalClusterSpecWithContext(ctx context.Context, request *DescribeExternalClusterSpecRequest) (response *DescribeExternalClusterSpecResponse, err error) {
    if request == nil {
        request = NewDescribeExternalClusterSpecRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExternalClusterSpec require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExternalClusterSpecResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageCachesRequest() (request *DescribeImageCachesRequest) {
    request = &DescribeImageCachesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeImageCaches")
    
    
    return
}

func NewDescribeImageCachesResponse() (response *DescribeImageCachesResponse) {
    response = &DescribeImageCachesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeImageCaches
// 查询镜像缓存信息接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeImageCaches(request *DescribeImageCachesRequest) (response *DescribeImageCachesResponse, err error) {
    return c.DescribeImageCachesWithContext(context.Background(), request)
}

// DescribeImageCaches
// 查询镜像缓存信息接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeImageCachesWithContext(ctx context.Context, request *DescribeImageCachesRequest) (response *DescribeImageCachesResponse, err error) {
    if request == nil {
        request = NewDescribeImageCachesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImageCaches require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImageCachesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImagesRequest() (request *DescribeImagesRequest) {
    request = &DescribeImagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeImages")
    
    
    return
}

func NewDescribeImagesResponse() (response *DescribeImagesResponse) {
    response = &DescribeImagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeImages
// 获取镜像信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_COMPONENTCLINETHTTP = "InternalError.ComponentClinetHttp"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_IMAGEIDNOTFOUND = "InternalError.ImageIdNotFound"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_ROUTETABLENOTEMPTY = "InvalidParameter.RouteTableNotEmpty"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeImages(request *DescribeImagesRequest) (response *DescribeImagesResponse, err error) {
    return c.DescribeImagesWithContext(context.Background(), request)
}

// DescribeImages
// 获取镜像信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_COMPONENTCLINETHTTP = "InternalError.ComponentClinetHttp"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_IMAGEIDNOTFOUND = "InternalError.ImageIdNotFound"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_ROUTETABLENOTEMPTY = "InvalidParameter.RouteTableNotEmpty"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeImagesWithContext(ctx context.Context, request *DescribeImagesRequest) (response *DescribeImagesResponse, err error) {
    if request == nil {
        request = NewDescribeImagesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImages require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImagesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusAgentInstancesRequest() (request *DescribePrometheusAgentInstancesRequest) {
    request = &DescribePrometheusAgentInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusAgentInstances")
    
    
    return
}

func NewDescribePrometheusAgentInstancesResponse() (response *DescribePrometheusAgentInstancesResponse) {
    response = &DescribePrometheusAgentInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePrometheusAgentInstances
// 获取关联目标集群的实例列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribePrometheusAgentInstances(request *DescribePrometheusAgentInstancesRequest) (response *DescribePrometheusAgentInstancesResponse, err error) {
    return c.DescribePrometheusAgentInstancesWithContext(context.Background(), request)
}

// DescribePrometheusAgentInstances
// 获取关联目标集群的实例列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribePrometheusAgentInstancesWithContext(ctx context.Context, request *DescribePrometheusAgentInstancesRequest) (response *DescribePrometheusAgentInstancesResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusAgentInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusAgentInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusAgentInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusAgentsRequest() (request *DescribePrometheusAgentsRequest) {
    request = &DescribePrometheusAgentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusAgents")
    
    
    return
}

func NewDescribePrometheusAgentsResponse() (response *DescribePrometheusAgentsResponse) {
    response = &DescribePrometheusAgentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePrometheusAgents
// 获取被关联集群列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) DescribePrometheusAgents(request *DescribePrometheusAgentsRequest) (response *DescribePrometheusAgentsResponse, err error) {
    return c.DescribePrometheusAgentsWithContext(context.Background(), request)
}

// DescribePrometheusAgents
// 获取被关联集群列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) DescribePrometheusAgentsWithContext(ctx context.Context, request *DescribePrometheusAgentsRequest) (response *DescribePrometheusAgentsResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusAgentsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusAgents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusAgentsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusAlertHistoryRequest() (request *DescribePrometheusAlertHistoryRequest) {
    request = &DescribePrometheusAlertHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusAlertHistory")
    
    
    return
}

func NewDescribePrometheusAlertHistoryResponse() (response *DescribePrometheusAlertHistoryResponse) {
    response = &DescribePrometheusAlertHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePrometheusAlertHistory
// 获取告警历史
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusAlertHistory(request *DescribePrometheusAlertHistoryRequest) (response *DescribePrometheusAlertHistoryResponse, err error) {
    return c.DescribePrometheusAlertHistoryWithContext(context.Background(), request)
}

// DescribePrometheusAlertHistory
// 获取告警历史
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusAlertHistoryWithContext(ctx context.Context, request *DescribePrometheusAlertHistoryRequest) (response *DescribePrometheusAlertHistoryResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusAlertHistoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusAlertHistory require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusAlertHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusAlertPolicyRequest() (request *DescribePrometheusAlertPolicyRequest) {
    request = &DescribePrometheusAlertPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusAlertPolicy")
    
    
    return
}

func NewDescribePrometheusAlertPolicyResponse() (response *DescribePrometheusAlertPolicyResponse) {
    response = &DescribePrometheusAlertPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePrometheusAlertPolicy
// 获取2.0实例告警策略列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_KUBECLIENTCONF = "FailedOperation.KubeClientConf"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusAlertPolicy(request *DescribePrometheusAlertPolicyRequest) (response *DescribePrometheusAlertPolicyResponse, err error) {
    return c.DescribePrometheusAlertPolicyWithContext(context.Background(), request)
}

// DescribePrometheusAlertPolicy
// 获取2.0实例告警策略列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_KUBECLIENTCONF = "FailedOperation.KubeClientConf"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusAlertPolicyWithContext(ctx context.Context, request *DescribePrometheusAlertPolicyRequest) (response *DescribePrometheusAlertPolicyResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusAlertPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusAlertPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusAlertPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusAlertRuleRequest() (request *DescribePrometheusAlertRuleRequest) {
    request = &DescribePrometheusAlertRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusAlertRule")
    
    
    return
}

func NewDescribePrometheusAlertRuleResponse() (response *DescribePrometheusAlertRuleResponse) {
    response = &DescribePrometheusAlertRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePrometheusAlertRule
// 获取告警规则列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusAlertRule(request *DescribePrometheusAlertRuleRequest) (response *DescribePrometheusAlertRuleResponse, err error) {
    return c.DescribePrometheusAlertRuleWithContext(context.Background(), request)
}

// DescribePrometheusAlertRule
// 获取告警规则列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusAlertRuleWithContext(ctx context.Context, request *DescribePrometheusAlertRuleRequest) (response *DescribePrometheusAlertRuleResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusAlertRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusAlertRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusAlertRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusClusterAgentsRequest() (request *DescribePrometheusClusterAgentsRequest) {
    request = &DescribePrometheusClusterAgentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusClusterAgents")
    
    
    return
}

func NewDescribePrometheusClusterAgentsResponse() (response *DescribePrometheusClusterAgentsResponse) {
    response = &DescribePrometheusClusterAgentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePrometheusClusterAgents
// 获取TMP实例关联集群列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) DescribePrometheusClusterAgents(request *DescribePrometheusClusterAgentsRequest) (response *DescribePrometheusClusterAgentsResponse, err error) {
    return c.DescribePrometheusClusterAgentsWithContext(context.Background(), request)
}

// DescribePrometheusClusterAgents
// 获取TMP实例关联集群列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) DescribePrometheusClusterAgentsWithContext(ctx context.Context, request *DescribePrometheusClusterAgentsRequest) (response *DescribePrometheusClusterAgentsResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusClusterAgentsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusClusterAgents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusClusterAgentsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusConfigRequest() (request *DescribePrometheusConfigRequest) {
    request = &DescribePrometheusConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusConfig")
    
    
    return
}

func NewDescribePrometheusConfigResponse() (response *DescribePrometheusConfigResponse) {
    response = &DescribePrometheusConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePrometheusConfig
// 拉取Prometheus配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  FAILEDOPERATION_KUBECLIENTCONF = "FailedOperation.KubeClientConf"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusConfig(request *DescribePrometheusConfigRequest) (response *DescribePrometheusConfigResponse, err error) {
    return c.DescribePrometheusConfigWithContext(context.Background(), request)
}

// DescribePrometheusConfig
// 拉取Prometheus配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  FAILEDOPERATION_KUBECLIENTCONF = "FailedOperation.KubeClientConf"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusConfigWithContext(ctx context.Context, request *DescribePrometheusConfigRequest) (response *DescribePrometheusConfigResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusGlobalConfigRequest() (request *DescribePrometheusGlobalConfigRequest) {
    request = &DescribePrometheusGlobalConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusGlobalConfig")
    
    
    return
}

func NewDescribePrometheusGlobalConfigResponse() (response *DescribePrometheusGlobalConfigResponse) {
    response = &DescribePrometheusGlobalConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePrometheusGlobalConfig
// 获得实例级别抓取配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusGlobalConfig(request *DescribePrometheusGlobalConfigRequest) (response *DescribePrometheusGlobalConfigResponse, err error) {
    return c.DescribePrometheusGlobalConfigWithContext(context.Background(), request)
}

// DescribePrometheusGlobalConfig
// 获得实例级别抓取配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusGlobalConfigWithContext(ctx context.Context, request *DescribePrometheusGlobalConfigRequest) (response *DescribePrometheusGlobalConfigResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusGlobalConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusGlobalConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusGlobalConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusGlobalNotificationRequest() (request *DescribePrometheusGlobalNotificationRequest) {
    request = &DescribePrometheusGlobalNotificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusGlobalNotification")
    
    
    return
}

func NewDescribePrometheusGlobalNotificationResponse() (response *DescribePrometheusGlobalNotificationResponse) {
    response = &DescribePrometheusGlobalNotificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePrometheusGlobalNotification
// 查询全局告警通知渠道
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_KUBECLIENTCONF = "FailedOperation.KubeClientConf"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusGlobalNotification(request *DescribePrometheusGlobalNotificationRequest) (response *DescribePrometheusGlobalNotificationResponse, err error) {
    return c.DescribePrometheusGlobalNotificationWithContext(context.Background(), request)
}

// DescribePrometheusGlobalNotification
// 查询全局告警通知渠道
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_KUBECLIENTCONF = "FailedOperation.KubeClientConf"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusGlobalNotificationWithContext(ctx context.Context, request *DescribePrometheusGlobalNotificationRequest) (response *DescribePrometheusGlobalNotificationResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusGlobalNotificationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusGlobalNotification require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusGlobalNotificationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusInstanceRequest() (request *DescribePrometheusInstanceRequest) {
    request = &DescribePrometheusInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusInstance")
    
    
    return
}

func NewDescribePrometheusInstanceResponse() (response *DescribePrometheusInstanceResponse) {
    response = &DescribePrometheusInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePrometheusInstance
// 获取实例详细信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusInstance(request *DescribePrometheusInstanceRequest) (response *DescribePrometheusInstanceResponse, err error) {
    return c.DescribePrometheusInstanceWithContext(context.Background(), request)
}

// DescribePrometheusInstance
// 获取实例详细信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusInstanceWithContext(ctx context.Context, request *DescribePrometheusInstanceRequest) (response *DescribePrometheusInstanceResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusInstanceInitStatusRequest() (request *DescribePrometheusInstanceInitStatusRequest) {
    request = &DescribePrometheusInstanceInitStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusInstanceInitStatus")
    
    
    return
}

func NewDescribePrometheusInstanceInitStatusResponse() (response *DescribePrometheusInstanceInitStatusResponse) {
    response = &DescribePrometheusInstanceInitStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePrometheusInstanceInitStatus
// 获取2.0实例初始化任务状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_TASKNOTFOUND = "InternalError.TaskNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) DescribePrometheusInstanceInitStatus(request *DescribePrometheusInstanceInitStatusRequest) (response *DescribePrometheusInstanceInitStatusResponse, err error) {
    return c.DescribePrometheusInstanceInitStatusWithContext(context.Background(), request)
}

// DescribePrometheusInstanceInitStatus
// 获取2.0实例初始化任务状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_TASKNOTFOUND = "InternalError.TaskNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) DescribePrometheusInstanceInitStatusWithContext(ctx context.Context, request *DescribePrometheusInstanceInitStatusRequest) (response *DescribePrometheusInstanceInitStatusResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusInstanceInitStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusInstanceInitStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusInstanceInitStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusInstancesOverviewRequest() (request *DescribePrometheusInstancesOverviewRequest) {
    request = &DescribePrometheusInstancesOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusInstancesOverview")
    
    
    return
}

func NewDescribePrometheusInstancesOverviewResponse() (response *DescribePrometheusInstancesOverviewResponse) {
    response = &DescribePrometheusInstancesOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePrometheusInstancesOverview
// 获取与云监控融合实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) DescribePrometheusInstancesOverview(request *DescribePrometheusInstancesOverviewRequest) (response *DescribePrometheusInstancesOverviewResponse, err error) {
    return c.DescribePrometheusInstancesOverviewWithContext(context.Background(), request)
}

// DescribePrometheusInstancesOverview
// 获取与云监控融合实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) DescribePrometheusInstancesOverviewWithContext(ctx context.Context, request *DescribePrometheusInstancesOverviewRequest) (response *DescribePrometheusInstancesOverviewResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusInstancesOverviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusInstancesOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusInstancesOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusOverviewsRequest() (request *DescribePrometheusOverviewsRequest) {
    request = &DescribePrometheusOverviewsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusOverviews")
    
    
    return
}

func NewDescribePrometheusOverviewsResponse() (response *DescribePrometheusOverviewsResponse) {
    response = &DescribePrometheusOverviewsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePrometheusOverviews
// 获取实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
func (c *Client) DescribePrometheusOverviews(request *DescribePrometheusOverviewsRequest) (response *DescribePrometheusOverviewsResponse, err error) {
    return c.DescribePrometheusOverviewsWithContext(context.Background(), request)
}

// DescribePrometheusOverviews
// 获取实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
func (c *Client) DescribePrometheusOverviewsWithContext(ctx context.Context, request *DescribePrometheusOverviewsRequest) (response *DescribePrometheusOverviewsResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusOverviewsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusOverviews require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusOverviewsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusRecordRulesRequest() (request *DescribePrometheusRecordRulesRequest) {
    request = &DescribePrometheusRecordRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusRecordRules")
    
    
    return
}

func NewDescribePrometheusRecordRulesResponse() (response *DescribePrometheusRecordRulesResponse) {
    response = &DescribePrometheusRecordRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePrometheusRecordRules
// 获取聚合规则列表，包含关联集群内crd资源创建的record rule
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_KUBECLIENTCONF = "FailedOperation.KubeClientConf"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusRecordRules(request *DescribePrometheusRecordRulesRequest) (response *DescribePrometheusRecordRulesResponse, err error) {
    return c.DescribePrometheusRecordRulesWithContext(context.Background(), request)
}

// DescribePrometheusRecordRules
// 获取聚合规则列表，包含关联集群内crd资源创建的record rule
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_KUBECLIENTCONF = "FailedOperation.KubeClientConf"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusRecordRulesWithContext(ctx context.Context, request *DescribePrometheusRecordRulesRequest) (response *DescribePrometheusRecordRulesResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusRecordRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusRecordRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusRecordRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusTargetsRequest() (request *DescribePrometheusTargetsRequest) {
    request = &DescribePrometheusTargetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusTargets")
    
    
    return
}

func NewDescribePrometheusTargetsResponse() (response *DescribePrometheusTargetsResponse) {
    response = &DescribePrometheusTargetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePrometheusTargets
// 获取targets信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusTargets(request *DescribePrometheusTargetsRequest) (response *DescribePrometheusTargetsResponse, err error) {
    return c.DescribePrometheusTargetsWithContext(context.Background(), request)
}

// DescribePrometheusTargets
// 获取targets信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusTargetsWithContext(ctx context.Context, request *DescribePrometheusTargetsRequest) (response *DescribePrometheusTargetsResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusTargetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusTargets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusTargetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusTempRequest() (request *DescribePrometheusTempRequest) {
    request = &DescribePrometheusTempRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusTemp")
    
    
    return
}

func NewDescribePrometheusTempResponse() (response *DescribePrometheusTempResponse) {
    response = &DescribePrometheusTempResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePrometheusTemp
// 拉取模板列表，默认模板将总是在最前面
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) DescribePrometheusTemp(request *DescribePrometheusTempRequest) (response *DescribePrometheusTempResponse, err error) {
    return c.DescribePrometheusTempWithContext(context.Background(), request)
}

// DescribePrometheusTemp
// 拉取模板列表，默认模板将总是在最前面
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) DescribePrometheusTempWithContext(ctx context.Context, request *DescribePrometheusTempRequest) (response *DescribePrometheusTempResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusTempRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusTemp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusTempResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusTempSyncRequest() (request *DescribePrometheusTempSyncRequest) {
    request = &DescribePrometheusTempSyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusTempSync")
    
    
    return
}

func NewDescribePrometheusTempSyncResponse() (response *DescribePrometheusTempSyncResponse) {
    response = &DescribePrometheusTempSyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePrometheusTempSync
// 获取模板关联实例信息，针对V2版本实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribePrometheusTempSync(request *DescribePrometheusTempSyncRequest) (response *DescribePrometheusTempSyncResponse, err error) {
    return c.DescribePrometheusTempSyncWithContext(context.Background(), request)
}

// DescribePrometheusTempSync
// 获取模板关联实例信息，针对V2版本实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribePrometheusTempSyncWithContext(ctx context.Context, request *DescribePrometheusTempSyncRequest) (response *DescribePrometheusTempSyncResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusTempSyncRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusTempSync require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusTempSyncResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusTemplateSyncRequest() (request *DescribePrometheusTemplateSyncRequest) {
    request = &DescribePrometheusTemplateSyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusTemplateSync")
    
    
    return
}

func NewDescribePrometheusTemplateSyncResponse() (response *DescribePrometheusTemplateSyncResponse) {
    response = &DescribePrometheusTemplateSyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePrometheusTemplateSync
// 获取模板同步信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribePrometheusTemplateSync(request *DescribePrometheusTemplateSyncRequest) (response *DescribePrometheusTemplateSyncResponse, err error) {
    return c.DescribePrometheusTemplateSyncWithContext(context.Background(), request)
}

// DescribePrometheusTemplateSync
// 获取模板同步信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribePrometheusTemplateSyncWithContext(ctx context.Context, request *DescribePrometheusTemplateSyncRequest) (response *DescribePrometheusTemplateSyncResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusTemplateSyncRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusTemplateSync require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusTemplateSyncResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusTemplatesRequest() (request *DescribePrometheusTemplatesRequest) {
    request = &DescribePrometheusTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusTemplates")
    
    
    return
}

func NewDescribePrometheusTemplatesResponse() (response *DescribePrometheusTemplatesResponse) {
    response = &DescribePrometheusTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePrometheusTemplates
// 拉取模板列表，默认模板将总是在最前面
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
func (c *Client) DescribePrometheusTemplates(request *DescribePrometheusTemplatesRequest) (response *DescribePrometheusTemplatesResponse, err error) {
    return c.DescribePrometheusTemplatesWithContext(context.Background(), request)
}

// DescribePrometheusTemplates
// 拉取模板列表，默认模板将总是在最前面
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
func (c *Client) DescribePrometheusTemplatesWithContext(ctx context.Context, request *DescribePrometheusTemplatesRequest) (response *DescribePrometheusTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRegionsRequest() (request *DescribeRegionsRequest) {
    request = &DescribeRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeRegions")
    
    
    return
}

func NewDescribeRegionsResponse() (response *DescribeRegionsResponse) {
    response = &DescribeRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRegions
// 获取容器服务支持的所有地域
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRegions(request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
    return c.DescribeRegionsWithContext(context.Background(), request)
}

// DescribeRegions
// 获取容器服务支持的所有地域
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
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

func NewDescribeResourceUsageRequest() (request *DescribeResourceUsageRequest) {
    request = &DescribeResourceUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeResourceUsage")
    
    
    return
}

func NewDescribeResourceUsageResponse() (response *DescribeResourceUsageResponse) {
    response = &DescribeResourceUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeResourceUsage
// 获取集群资源使用量
//
// 可能返回的错误码:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) DescribeResourceUsage(request *DescribeResourceUsageRequest) (response *DescribeResourceUsageResponse, err error) {
    return c.DescribeResourceUsageWithContext(context.Background(), request)
}

// DescribeResourceUsage
// 获取集群资源使用量
//
// 可能返回的错误码:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) DescribeResourceUsageWithContext(ctx context.Context, request *DescribeResourceUsageRequest) (response *DescribeResourceUsageResponse, err error) {
    if request == nil {
        request = NewDescribeResourceUsageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResourceUsage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResourceUsageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRouteTableConflictsRequest() (request *DescribeRouteTableConflictsRequest) {
    request = &DescribeRouteTableConflictsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeRouteTableConflicts")
    
    
    return
}

func NewDescribeRouteTableConflictsResponse() (response *DescribeRouteTableConflictsResponse) {
    response = &DescribeRouteTableConflictsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRouteTableConflicts
// 查询路由表冲突列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CIDRMASKSIZEOUTOFRANGE = "InternalError.CidrMaskSizeOutOfRange"
//  INTERNALERROR_INVALIDPRIVATENETWORKCIDR = "InternalError.InvalidPrivateNetworkCidr"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_VPCRECODRNOTFOUND = "InternalError.VpcRecodrNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRouteTableConflicts(request *DescribeRouteTableConflictsRequest) (response *DescribeRouteTableConflictsResponse, err error) {
    return c.DescribeRouteTableConflictsWithContext(context.Background(), request)
}

// DescribeRouteTableConflicts
// 查询路由表冲突列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CIDRMASKSIZEOUTOFRANGE = "InternalError.CidrMaskSizeOutOfRange"
//  INTERNALERROR_INVALIDPRIVATENETWORKCIDR = "InternalError.InvalidPrivateNetworkCidr"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_VPCRECODRNOTFOUND = "InternalError.VpcRecodrNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRouteTableConflictsWithContext(ctx context.Context, request *DescribeRouteTableConflictsRequest) (response *DescribeRouteTableConflictsResponse, err error) {
    if request == nil {
        request = NewDescribeRouteTableConflictsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRouteTableConflicts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRouteTableConflictsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTKEEdgeClusterCredentialRequest() (request *DescribeTKEEdgeClusterCredentialRequest) {
    request = &DescribeTKEEdgeClusterCredentialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeTKEEdgeClusterCredential")
    
    
    return
}

func NewDescribeTKEEdgeClusterCredentialResponse() (response *DescribeTKEEdgeClusterCredentialResponse) {
    response = &DescribeTKEEdgeClusterCredentialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTKEEdgeClusterCredential
// 获取边缘计算集群的认证信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTKEEdgeClusterCredential(request *DescribeTKEEdgeClusterCredentialRequest) (response *DescribeTKEEdgeClusterCredentialResponse, err error) {
    return c.DescribeTKEEdgeClusterCredentialWithContext(context.Background(), request)
}

// DescribeTKEEdgeClusterCredential
// 获取边缘计算集群的认证信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTKEEdgeClusterCredentialWithContext(ctx context.Context, request *DescribeTKEEdgeClusterCredentialRequest) (response *DescribeTKEEdgeClusterCredentialResponse, err error) {
    if request == nil {
        request = NewDescribeTKEEdgeClusterCredentialRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTKEEdgeClusterCredential require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTKEEdgeClusterCredentialResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTKEEdgeClusterStatusRequest() (request *DescribeTKEEdgeClusterStatusRequest) {
    request = &DescribeTKEEdgeClusterStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeTKEEdgeClusterStatus")
    
    
    return
}

func NewDescribeTKEEdgeClusterStatusResponse() (response *DescribeTKEEdgeClusterStatusResponse) {
    response = &DescribeTKEEdgeClusterStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTKEEdgeClusterStatus
// 获取边缘计算集群的当前状态以及过程信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTKEEdgeClusterStatus(request *DescribeTKEEdgeClusterStatusRequest) (response *DescribeTKEEdgeClusterStatusResponse, err error) {
    return c.DescribeTKEEdgeClusterStatusWithContext(context.Background(), request)
}

// DescribeTKEEdgeClusterStatus
// 获取边缘计算集群的当前状态以及过程信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTKEEdgeClusterStatusWithContext(ctx context.Context, request *DescribeTKEEdgeClusterStatusRequest) (response *DescribeTKEEdgeClusterStatusResponse, err error) {
    if request == nil {
        request = NewDescribeTKEEdgeClusterStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTKEEdgeClusterStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTKEEdgeClusterStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTKEEdgeClustersRequest() (request *DescribeTKEEdgeClustersRequest) {
    request = &DescribeTKEEdgeClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeTKEEdgeClusters")
    
    
    return
}

func NewDescribeTKEEdgeClustersResponse() (response *DescribeTKEEdgeClustersResponse) {
    response = &DescribeTKEEdgeClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTKEEdgeClusters
// 查询边缘集群列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTKEEdgeClusters(request *DescribeTKEEdgeClustersRequest) (response *DescribeTKEEdgeClustersResponse, err error) {
    return c.DescribeTKEEdgeClustersWithContext(context.Background(), request)
}

// DescribeTKEEdgeClusters
// 查询边缘集群列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTKEEdgeClustersWithContext(ctx context.Context, request *DescribeTKEEdgeClustersRequest) (response *DescribeTKEEdgeClustersResponse, err error) {
    if request == nil {
        request = NewDescribeTKEEdgeClustersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTKEEdgeClusters require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTKEEdgeClustersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTKEEdgeExternalKubeconfigRequest() (request *DescribeTKEEdgeExternalKubeconfigRequest) {
    request = &DescribeTKEEdgeExternalKubeconfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeTKEEdgeExternalKubeconfig")
    
    
    return
}

func NewDescribeTKEEdgeExternalKubeconfigResponse() (response *DescribeTKEEdgeExternalKubeconfigResponse) {
    response = &DescribeTKEEdgeExternalKubeconfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTKEEdgeExternalKubeconfig
// 获取边缘计算外部访问的kubeconfig
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTKEEdgeExternalKubeconfig(request *DescribeTKEEdgeExternalKubeconfigRequest) (response *DescribeTKEEdgeExternalKubeconfigResponse, err error) {
    return c.DescribeTKEEdgeExternalKubeconfigWithContext(context.Background(), request)
}

// DescribeTKEEdgeExternalKubeconfig
// 获取边缘计算外部访问的kubeconfig
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTKEEdgeExternalKubeconfigWithContext(ctx context.Context, request *DescribeTKEEdgeExternalKubeconfigRequest) (response *DescribeTKEEdgeExternalKubeconfigResponse, err error) {
    if request == nil {
        request = NewDescribeTKEEdgeExternalKubeconfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTKEEdgeExternalKubeconfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTKEEdgeExternalKubeconfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTKEEdgeScriptRequest() (request *DescribeTKEEdgeScriptRequest) {
    request = &DescribeTKEEdgeScriptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeTKEEdgeScript")
    
    
    return
}

func NewDescribeTKEEdgeScriptResponse() (response *DescribeTKEEdgeScriptResponse) {
    response = &DescribeTKEEdgeScriptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTKEEdgeScript
// 获取边缘脚本链接，此接口用于添加第三方节点，通过下载脚本从而将节点添加到边缘集群。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTKEEdgeScript(request *DescribeTKEEdgeScriptRequest) (response *DescribeTKEEdgeScriptResponse, err error) {
    return c.DescribeTKEEdgeScriptWithContext(context.Background(), request)
}

// DescribeTKEEdgeScript
// 获取边缘脚本链接，此接口用于添加第三方节点，通过下载脚本从而将节点添加到边缘集群。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTKEEdgeScriptWithContext(ctx context.Context, request *DescribeTKEEdgeScriptRequest) (response *DescribeTKEEdgeScriptResponse, err error) {
    if request == nil {
        request = NewDescribeTKEEdgeScriptRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTKEEdgeScript require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTKEEdgeScriptResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVersionsRequest() (request *DescribeVersionsRequest) {
    request = &DescribeVersionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeVersions")
    
    
    return
}

func NewDescribeVersionsResponse() (response *DescribeVersionsResponse) {
    response = &DescribeVersionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVersions
// 获取集群版本信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLINETHTTP = "FailedOperation.ComponentClinetHttp"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_COMPONENTCLINETHTTP = "InternalError.ComponentClinetHttp"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_ROUTETABLENOTEMPTY = "InvalidParameter.RouteTableNotEmpty"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeVersions(request *DescribeVersionsRequest) (response *DescribeVersionsResponse, err error) {
    return c.DescribeVersionsWithContext(context.Background(), request)
}

// DescribeVersions
// 获取集群版本信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLINETHTTP = "FailedOperation.ComponentClinetHttp"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_COMPONENTCLINETHTTP = "InternalError.ComponentClinetHttp"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_ROUTETABLENOTEMPTY = "InvalidParameter.RouteTableNotEmpty"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeVersionsWithContext(ctx context.Context, request *DescribeVersionsRequest) (response *DescribeVersionsResponse, err error) {
    if request == nil {
        request = NewDescribeVersionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVersions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVersionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVpcCniPodLimitsRequest() (request *DescribeVpcCniPodLimitsRequest) {
    request = &DescribeVpcCniPodLimitsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeVpcCniPodLimits")
    
    
    return
}

func NewDescribeVpcCniPodLimitsResponse() (response *DescribeVpcCniPodLimitsResponse) {
    response = &DescribeVpcCniPodLimitsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVpcCniPodLimits
// 本接口查询当前用户和地域在指定可用区下的机型可支持的最大 TKE VPC-CNI 网络模式的 Pod 数量
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeVpcCniPodLimits(request *DescribeVpcCniPodLimitsRequest) (response *DescribeVpcCniPodLimitsResponse, err error) {
    return c.DescribeVpcCniPodLimitsWithContext(context.Background(), request)
}

// DescribeVpcCniPodLimits
// 本接口查询当前用户和地域在指定可用区下的机型可支持的最大 TKE VPC-CNI 网络模式的 Pod 数量
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeVpcCniPodLimitsWithContext(ctx context.Context, request *DescribeVpcCniPodLimitsRequest) (response *DescribeVpcCniPodLimitsResponse, err error) {
    if request == nil {
        request = NewDescribeVpcCniPodLimitsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVpcCniPodLimits require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVpcCniPodLimitsResponse()
    err = c.Send(request, response)
    return
}

func NewDisableClusterAuditRequest() (request *DisableClusterAuditRequest) {
    request = &DisableClusterAuditRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DisableClusterAudit")
    
    
    return
}

func NewDisableClusterAuditResponse() (response *DisableClusterAuditResponse) {
    response = &DisableClusterAuditResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisableClusterAudit
// 关闭集群审计
//
// 可能返回的错误码:
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  FAILEDOPERATION_KUBERNETESPATCHOPERATIONERROR = "FailedOperation.KubernetesPatchOperationError"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
func (c *Client) DisableClusterAudit(request *DisableClusterAuditRequest) (response *DisableClusterAuditResponse, err error) {
    return c.DisableClusterAuditWithContext(context.Background(), request)
}

// DisableClusterAudit
// 关闭集群审计
//
// 可能返回的错误码:
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  FAILEDOPERATION_KUBERNETESPATCHOPERATIONERROR = "FailedOperation.KubernetesPatchOperationError"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
func (c *Client) DisableClusterAuditWithContext(ctx context.Context, request *DisableClusterAuditRequest) (response *DisableClusterAuditResponse, err error) {
    if request == nil {
        request = NewDisableClusterAuditRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableClusterAudit require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableClusterAuditResponse()
    err = c.Send(request, response)
    return
}

func NewDisableClusterDeletionProtectionRequest() (request *DisableClusterDeletionProtectionRequest) {
    request = &DisableClusterDeletionProtectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DisableClusterDeletionProtection")
    
    
    return
}

func NewDisableClusterDeletionProtectionResponse() (response *DisableClusterDeletionProtectionResponse) {
    response = &DisableClusterDeletionProtectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisableClusterDeletionProtection
// 关闭集群删除保护
//
// 可能返回的错误码:
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DisableClusterDeletionProtection(request *DisableClusterDeletionProtectionRequest) (response *DisableClusterDeletionProtectionResponse, err error) {
    return c.DisableClusterDeletionProtectionWithContext(context.Background(), request)
}

// DisableClusterDeletionProtection
// 关闭集群删除保护
//
// 可能返回的错误码:
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DisableClusterDeletionProtectionWithContext(ctx context.Context, request *DisableClusterDeletionProtectionRequest) (response *DisableClusterDeletionProtectionResponse, err error) {
    if request == nil {
        request = NewDisableClusterDeletionProtectionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableClusterDeletionProtection require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableClusterDeletionProtectionResponse()
    err = c.Send(request, response)
    return
}

func NewDisableEventPersistenceRequest() (request *DisableEventPersistenceRequest) {
    request = &DisableEventPersistenceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DisableEventPersistence")
    
    
    return
}

func NewDisableEventPersistenceResponse() (response *DisableEventPersistenceResponse) {
    response = &DisableEventPersistenceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisableEventPersistence
// 关闭事件持久化功能
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLUSTERNOTFOUND = "FailedOperation.ClusterNotFound"
//  FAILEDOPERATION_KUBERNETESDELETEOPERATIONERROR = "FailedOperation.KubernetesDeleteOperationError"
//  FAILEDOPERATION_RBACFORBIDDEN = "FailedOperation.RBACForbidden"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_KUBERNETESCLIENTBUILDERROR = "InternalError.KubernetesClientBuildError"
//  INTERNALERROR_KUBERNETESDELETEOPERATIONERROR = "InternalError.KubernetesDeleteOperationError"
//  INTERNALERROR_KUBERNETESGETOPERATIONERROR = "InternalError.KubernetesGetOperationError"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
func (c *Client) DisableEventPersistence(request *DisableEventPersistenceRequest) (response *DisableEventPersistenceResponse, err error) {
    return c.DisableEventPersistenceWithContext(context.Background(), request)
}

// DisableEventPersistence
// 关闭事件持久化功能
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLUSTERNOTFOUND = "FailedOperation.ClusterNotFound"
//  FAILEDOPERATION_KUBERNETESDELETEOPERATIONERROR = "FailedOperation.KubernetesDeleteOperationError"
//  FAILEDOPERATION_RBACFORBIDDEN = "FailedOperation.RBACForbidden"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_KUBERNETESCLIENTBUILDERROR = "InternalError.KubernetesClientBuildError"
//  INTERNALERROR_KUBERNETESDELETEOPERATIONERROR = "InternalError.KubernetesDeleteOperationError"
//  INTERNALERROR_KUBERNETESGETOPERATIONERROR = "InternalError.KubernetesGetOperationError"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
func (c *Client) DisableEventPersistenceWithContext(ctx context.Context, request *DisableEventPersistenceRequest) (response *DisableEventPersistenceResponse, err error) {
    if request == nil {
        request = NewDisableEventPersistenceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableEventPersistence require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableEventPersistenceResponse()
    err = c.Send(request, response)
    return
}

func NewDisableVpcCniNetworkTypeRequest() (request *DisableVpcCniNetworkTypeRequest) {
    request = &DisableVpcCniNetworkTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DisableVpcCniNetworkType")
    
    
    return
}

func NewDisableVpcCniNetworkTypeResponse() (response *DisableVpcCniNetworkTypeResponse) {
    response = &DisableVpcCniNetworkTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisableVpcCniNetworkType
// 提供给附加了VPC-CNI能力的Global-Route集群关闭VPC-CNI
//
// 可能返回的错误码:
//  FAILEDOPERATION_DISABLEVPCCNIFAILED = "FailedOperation.DisableVPCCNIFailed"
//  INTERNALERROR_KUBECLIENTCREATE = "InternalError.KubeClientCreate"
//  INTERNALERROR_KUBECOMMON = "InternalError.KubeCommon"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) DisableVpcCniNetworkType(request *DisableVpcCniNetworkTypeRequest) (response *DisableVpcCniNetworkTypeResponse, err error) {
    return c.DisableVpcCniNetworkTypeWithContext(context.Background(), request)
}

// DisableVpcCniNetworkType
// 提供给附加了VPC-CNI能力的Global-Route集群关闭VPC-CNI
//
// 可能返回的错误码:
//  FAILEDOPERATION_DISABLEVPCCNIFAILED = "FailedOperation.DisableVPCCNIFailed"
//  INTERNALERROR_KUBECLIENTCREATE = "InternalError.KubeClientCreate"
//  INTERNALERROR_KUBECOMMON = "InternalError.KubeCommon"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) DisableVpcCniNetworkTypeWithContext(ctx context.Context, request *DisableVpcCniNetworkTypeRequest) (response *DisableVpcCniNetworkTypeResponse, err error) {
    if request == nil {
        request = NewDisableVpcCniNetworkTypeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableVpcCniNetworkType require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableVpcCniNetworkTypeResponse()
    err = c.Send(request, response)
    return
}

func NewDrainClusterVirtualNodeRequest() (request *DrainClusterVirtualNodeRequest) {
    request = &DrainClusterVirtualNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DrainClusterVirtualNode")
    
    
    return
}

func NewDrainClusterVirtualNodeResponse() (response *DrainClusterVirtualNodeResponse) {
    response = &DrainClusterVirtualNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DrainClusterVirtualNode
// 驱逐虚拟节点
//
// 可能返回的错误码:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DrainClusterVirtualNode(request *DrainClusterVirtualNodeRequest) (response *DrainClusterVirtualNodeResponse, err error) {
    return c.DrainClusterVirtualNodeWithContext(context.Background(), request)
}

// DrainClusterVirtualNode
// 驱逐虚拟节点
//
// 可能返回的错误码:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DrainClusterVirtualNodeWithContext(ctx context.Context, request *DrainClusterVirtualNodeRequest) (response *DrainClusterVirtualNodeResponse, err error) {
    if request == nil {
        request = NewDrainClusterVirtualNodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DrainClusterVirtualNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewDrainClusterVirtualNodeResponse()
    err = c.Send(request, response)
    return
}

func NewEnableClusterAuditRequest() (request *EnableClusterAuditRequest) {
    request = &EnableClusterAuditRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "EnableClusterAudit")
    
    
    return
}

func NewEnableClusterAuditResponse() (response *EnableClusterAuditResponse) {
    response = &EnableClusterAuditResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EnableClusterAudit
// 开启集群审计
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_CREATECLSINDEX = "FailedOperation.CreateClsIndex"
//  FAILEDOPERATION_CREATECLSMACHINEGROUP = "FailedOperation.CreateClsMachineGroup"
//  FAILEDOPERATION_CREATECLSTOPIC = "FailedOperation.CreateClsTopic"
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  FAILEDOPERATION_GETCLSMACHINEGROUP = "FailedOperation.GetClsMachineGroup"
//  FAILEDOPERATION_KUBERNETESCREATEOPERATIONERROR = "FailedOperation.KubernetesCreateOperationError"
//  FAILEDOPERATION_KUBERNETESDELETEOPERATIONERROR = "FailedOperation.KubernetesDeleteOperationError"
//  FAILEDOPERATION_KUBERNETESPATCHOPERATIONERROR = "FailedOperation.KubernetesPatchOperationError"
//  FAILEDOPERATION_MODIFYCLSINDEX = "FailedOperation.ModifyClsIndex"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
func (c *Client) EnableClusterAudit(request *EnableClusterAuditRequest) (response *EnableClusterAuditResponse, err error) {
    return c.EnableClusterAuditWithContext(context.Background(), request)
}

// EnableClusterAudit
// 开启集群审计
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_CREATECLSINDEX = "FailedOperation.CreateClsIndex"
//  FAILEDOPERATION_CREATECLSMACHINEGROUP = "FailedOperation.CreateClsMachineGroup"
//  FAILEDOPERATION_CREATECLSTOPIC = "FailedOperation.CreateClsTopic"
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  FAILEDOPERATION_GETCLSMACHINEGROUP = "FailedOperation.GetClsMachineGroup"
//  FAILEDOPERATION_KUBERNETESCREATEOPERATIONERROR = "FailedOperation.KubernetesCreateOperationError"
//  FAILEDOPERATION_KUBERNETESDELETEOPERATIONERROR = "FailedOperation.KubernetesDeleteOperationError"
//  FAILEDOPERATION_KUBERNETESPATCHOPERATIONERROR = "FailedOperation.KubernetesPatchOperationError"
//  FAILEDOPERATION_MODIFYCLSINDEX = "FailedOperation.ModifyClsIndex"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
func (c *Client) EnableClusterAuditWithContext(ctx context.Context, request *EnableClusterAuditRequest) (response *EnableClusterAuditResponse, err error) {
    if request == nil {
        request = NewEnableClusterAuditRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableClusterAudit require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableClusterAuditResponse()
    err = c.Send(request, response)
    return
}

func NewEnableClusterDeletionProtectionRequest() (request *EnableClusterDeletionProtectionRequest) {
    request = &EnableClusterDeletionProtectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "EnableClusterDeletionProtection")
    
    
    return
}

func NewEnableClusterDeletionProtectionResponse() (response *EnableClusterDeletionProtectionResponse) {
    response = &EnableClusterDeletionProtectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EnableClusterDeletionProtection
// 启用集群删除保护
//
// 可能返回的错误码:
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) EnableClusterDeletionProtection(request *EnableClusterDeletionProtectionRequest) (response *EnableClusterDeletionProtectionResponse, err error) {
    return c.EnableClusterDeletionProtectionWithContext(context.Background(), request)
}

// EnableClusterDeletionProtection
// 启用集群删除保护
//
// 可能返回的错误码:
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) EnableClusterDeletionProtectionWithContext(ctx context.Context, request *EnableClusterDeletionProtectionRequest) (response *EnableClusterDeletionProtectionResponse, err error) {
    if request == nil {
        request = NewEnableClusterDeletionProtectionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableClusterDeletionProtection require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableClusterDeletionProtectionResponse()
    err = c.Send(request, response)
    return
}

func NewEnableEventPersistenceRequest() (request *EnableEventPersistenceRequest) {
    request = &EnableEventPersistenceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "EnableEventPersistence")
    
    
    return
}

func NewEnableEventPersistenceResponse() (response *EnableEventPersistenceResponse) {
    response = &EnableEventPersistenceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EnableEventPersistence
// 开启事件持久化功能
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLUSTERNOTFOUND = "FailedOperation.ClusterNotFound"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_CREATECLSCLIENT = "FailedOperation.CreateClsClient"
//  FAILEDOPERATION_CREATECLSCONFIG = "FailedOperation.CreateClsConfig"
//  FAILEDOPERATION_CREATECLSINDEX = "FailedOperation.CreateClsIndex"
//  FAILEDOPERATION_CREATECLSLOGSET = "FailedOperation.CreateClsLogSet"
//  FAILEDOPERATION_CREATECLSTOPIC = "FailedOperation.CreateClsTopic"
//  FAILEDOPERATION_GETCLSCONFIG = "FailedOperation.GetClsConfig"
//  FAILEDOPERATION_GETCLSLOGSET = "FailedOperation.GetClsLogSet"
//  FAILEDOPERATION_MODIFYCLSINDEX = "FailedOperation.ModifyClsIndex"
//  FAILEDOPERATION_RBACFORBIDDEN = "FailedOperation.RBACForbidden"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_KUBERNETESCLIENTBUILDERROR = "InternalError.KubernetesClientBuildError"
//  INTERNALERROR_KUBERNETESCREATEOPERATIONERROR = "InternalError.KubernetesCreateOperationError"
//  INTERNALERROR_KUBERNETESDELETEOPERATIONERROR = "InternalError.KubernetesDeleteOperationError"
//  INTERNALERROR_KUBERNETESGETOPERATIONERROR = "InternalError.KubernetesGetOperationError"
//  INTERNALERROR_KUBERNETESPATCHOPERATIONERROR = "InternalError.KubernetesPatchOperationError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
func (c *Client) EnableEventPersistence(request *EnableEventPersistenceRequest) (response *EnableEventPersistenceResponse, err error) {
    return c.EnableEventPersistenceWithContext(context.Background(), request)
}

// EnableEventPersistence
// 开启事件持久化功能
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLUSTERNOTFOUND = "FailedOperation.ClusterNotFound"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_CREATECLSCLIENT = "FailedOperation.CreateClsClient"
//  FAILEDOPERATION_CREATECLSCONFIG = "FailedOperation.CreateClsConfig"
//  FAILEDOPERATION_CREATECLSINDEX = "FailedOperation.CreateClsIndex"
//  FAILEDOPERATION_CREATECLSLOGSET = "FailedOperation.CreateClsLogSet"
//  FAILEDOPERATION_CREATECLSTOPIC = "FailedOperation.CreateClsTopic"
//  FAILEDOPERATION_GETCLSCONFIG = "FailedOperation.GetClsConfig"
//  FAILEDOPERATION_GETCLSLOGSET = "FailedOperation.GetClsLogSet"
//  FAILEDOPERATION_MODIFYCLSINDEX = "FailedOperation.ModifyClsIndex"
//  FAILEDOPERATION_RBACFORBIDDEN = "FailedOperation.RBACForbidden"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_KUBERNETESCLIENTBUILDERROR = "InternalError.KubernetesClientBuildError"
//  INTERNALERROR_KUBERNETESCREATEOPERATIONERROR = "InternalError.KubernetesCreateOperationError"
//  INTERNALERROR_KUBERNETESDELETEOPERATIONERROR = "InternalError.KubernetesDeleteOperationError"
//  INTERNALERROR_KUBERNETESGETOPERATIONERROR = "InternalError.KubernetesGetOperationError"
//  INTERNALERROR_KUBERNETESPATCHOPERATIONERROR = "InternalError.KubernetesPatchOperationError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
func (c *Client) EnableEventPersistenceWithContext(ctx context.Context, request *EnableEventPersistenceRequest) (response *EnableEventPersistenceResponse, err error) {
    if request == nil {
        request = NewEnableEventPersistenceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableEventPersistence require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableEventPersistenceResponse()
    err = c.Send(request, response)
    return
}

func NewEnableVpcCniNetworkTypeRequest() (request *EnableVpcCniNetworkTypeRequest) {
    request = &EnableVpcCniNetworkTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "EnableVpcCniNetworkType")
    
    
    return
}

func NewEnableVpcCniNetworkTypeResponse() (response *EnableVpcCniNetworkTypeResponse) {
    response = &EnableVpcCniNetworkTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EnableVpcCniNetworkType
// GR集群可以通过本接口附加vpc-cni容器网络插件，开启vpc-cni容器网络能力
//
// 可能返回的错误码:
//  FAILEDOPERATION_ENABLEVPCCNIFAILED = "FailedOperation.EnableVPCCNIFailed"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  UNSUPPORTEDOPERATION_CLUSTERNOTSUITENABLEVPCCNI = "UnsupportedOperation.ClusterNotSuitEnableVPCCNI"
func (c *Client) EnableVpcCniNetworkType(request *EnableVpcCniNetworkTypeRequest) (response *EnableVpcCniNetworkTypeResponse, err error) {
    return c.EnableVpcCniNetworkTypeWithContext(context.Background(), request)
}

// EnableVpcCniNetworkType
// GR集群可以通过本接口附加vpc-cni容器网络插件，开启vpc-cni容器网络能力
//
// 可能返回的错误码:
//  FAILEDOPERATION_ENABLEVPCCNIFAILED = "FailedOperation.EnableVPCCNIFailed"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  UNSUPPORTEDOPERATION_CLUSTERNOTSUITENABLEVPCCNI = "UnsupportedOperation.ClusterNotSuitEnableVPCCNI"
func (c *Client) EnableVpcCniNetworkTypeWithContext(ctx context.Context, request *EnableVpcCniNetworkTypeRequest) (response *EnableVpcCniNetworkTypeResponse, err error) {
    if request == nil {
        request = NewEnableVpcCniNetworkTypeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableVpcCniNetworkType require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableVpcCniNetworkTypeResponse()
    err = c.Send(request, response)
    return
}

func NewForwardApplicationRequestV3Request() (request *ForwardApplicationRequestV3Request) {
    request = &ForwardApplicationRequestV3Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "ForwardApplicationRequestV3")
    
    
    return
}

func NewForwardApplicationRequestV3Response() (response *ForwardApplicationRequestV3Response) {
    response = &ForwardApplicationRequestV3Response{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ForwardApplicationRequestV3
// 操作TKE集群的addon
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RBACFORBIDDEN = "FailedOperation.RBACForbidden"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_NOTINWHITELIST = "UnsupportedOperation.NotInWhitelist"
func (c *Client) ForwardApplicationRequestV3(request *ForwardApplicationRequestV3Request) (response *ForwardApplicationRequestV3Response, err error) {
    return c.ForwardApplicationRequestV3WithContext(context.Background(), request)
}

// ForwardApplicationRequestV3
// 操作TKE集群的addon
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RBACFORBIDDEN = "FailedOperation.RBACForbidden"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_NOTINWHITELIST = "UnsupportedOperation.NotInWhitelist"
func (c *Client) ForwardApplicationRequestV3WithContext(ctx context.Context, request *ForwardApplicationRequestV3Request) (response *ForwardApplicationRequestV3Response, err error) {
    if request == nil {
        request = NewForwardApplicationRequestV3Request()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ForwardApplicationRequestV3 require credential")
    }

    request.SetContext(ctx)
    
    response = NewForwardApplicationRequestV3Response()
    err = c.Send(request, response)
    return
}

func NewForwardTKEEdgeApplicationRequestV3Request() (request *ForwardTKEEdgeApplicationRequestV3Request) {
    request = &ForwardTKEEdgeApplicationRequestV3Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "ForwardTKEEdgeApplicationRequestV3")
    
    
    return
}

func NewForwardTKEEdgeApplicationRequestV3Response() (response *ForwardTKEEdgeApplicationRequestV3Response) {
    response = &ForwardTKEEdgeApplicationRequestV3Response{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ForwardTKEEdgeApplicationRequestV3
// 操作TKEEdge集群的addon
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RBACFORBIDDEN = "FailedOperation.RBACForbidden"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_NOTINWHITELIST = "UnsupportedOperation.NotInWhitelist"
func (c *Client) ForwardTKEEdgeApplicationRequestV3(request *ForwardTKEEdgeApplicationRequestV3Request) (response *ForwardTKEEdgeApplicationRequestV3Response, err error) {
    return c.ForwardTKEEdgeApplicationRequestV3WithContext(context.Background(), request)
}

// ForwardTKEEdgeApplicationRequestV3
// 操作TKEEdge集群的addon
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RBACFORBIDDEN = "FailedOperation.RBACForbidden"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_NOTINWHITELIST = "UnsupportedOperation.NotInWhitelist"
func (c *Client) ForwardTKEEdgeApplicationRequestV3WithContext(ctx context.Context, request *ForwardTKEEdgeApplicationRequestV3Request) (response *ForwardTKEEdgeApplicationRequestV3Response, err error) {
    if request == nil {
        request = NewForwardTKEEdgeApplicationRequestV3Request()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ForwardTKEEdgeApplicationRequestV3 require credential")
    }

    request.SetContext(ctx)
    
    response = NewForwardTKEEdgeApplicationRequestV3Response()
    err = c.Send(request, response)
    return
}

func NewGetClusterLevelPriceRequest() (request *GetClusterLevelPriceRequest) {
    request = &GetClusterLevelPriceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "GetClusterLevelPrice")
    
    
    return
}

func NewGetClusterLevelPriceResponse() (response *GetClusterLevelPriceResponse) {
    response = &GetClusterLevelPriceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetClusterLevelPrice
// 获取集群规模价格
//
// 可能返回的错误码:
//  FAILEDOPERATION_TRADECOMMON = "FailedOperation.TradeCommon"
//  INTERNALERROR_TRADECOMMON = "InternalError.TradeCommon"
func (c *Client) GetClusterLevelPrice(request *GetClusterLevelPriceRequest) (response *GetClusterLevelPriceResponse, err error) {
    return c.GetClusterLevelPriceWithContext(context.Background(), request)
}

// GetClusterLevelPrice
// 获取集群规模价格
//
// 可能返回的错误码:
//  FAILEDOPERATION_TRADECOMMON = "FailedOperation.TradeCommon"
//  INTERNALERROR_TRADECOMMON = "InternalError.TradeCommon"
func (c *Client) GetClusterLevelPriceWithContext(ctx context.Context, request *GetClusterLevelPriceRequest) (response *GetClusterLevelPriceResponse, err error) {
    if request == nil {
        request = NewGetClusterLevelPriceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetClusterLevelPrice require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetClusterLevelPriceResponse()
    err = c.Send(request, response)
    return
}

func NewGetMostSuitableImageCacheRequest() (request *GetMostSuitableImageCacheRequest) {
    request = &GetMostSuitableImageCacheRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "GetMostSuitableImageCache")
    
    
    return
}

func NewGetMostSuitableImageCacheResponse() (response *GetMostSuitableImageCacheResponse) {
    response = &GetMostSuitableImageCacheResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetMostSuitableImageCache
// 根据镜像列表，查询匹配的镜像缓存
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) GetMostSuitableImageCache(request *GetMostSuitableImageCacheRequest) (response *GetMostSuitableImageCacheResponse, err error) {
    return c.GetMostSuitableImageCacheWithContext(context.Background(), request)
}

// GetMostSuitableImageCache
// 根据镜像列表，查询匹配的镜像缓存
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) GetMostSuitableImageCacheWithContext(ctx context.Context, request *GetMostSuitableImageCacheRequest) (response *GetMostSuitableImageCacheResponse, err error) {
    if request == nil {
        request = NewGetMostSuitableImageCacheRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetMostSuitableImageCache require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetMostSuitableImageCacheResponse()
    err = c.Send(request, response)
    return
}

func NewGetTkeAppChartListRequest() (request *GetTkeAppChartListRequest) {
    request = &GetTkeAppChartListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "GetTkeAppChartList")
    
    
    return
}

func NewGetTkeAppChartListResponse() (response *GetTkeAppChartListResponse) {
    response = &GetTkeAppChartListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetTkeAppChartList
// 获取TKE支持的App列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) GetTkeAppChartList(request *GetTkeAppChartListRequest) (response *GetTkeAppChartListResponse, err error) {
    return c.GetTkeAppChartListWithContext(context.Background(), request)
}

// GetTkeAppChartList
// 获取TKE支持的App列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) GetTkeAppChartListWithContext(ctx context.Context, request *GetTkeAppChartListRequest) (response *GetTkeAppChartListResponse, err error) {
    if request == nil {
        request = NewGetTkeAppChartListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTkeAppChartList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTkeAppChartListResponse()
    err = c.Send(request, response)
    return
}

func NewGetUpgradeInstanceProgressRequest() (request *GetUpgradeInstanceProgressRequest) {
    request = &GetUpgradeInstanceProgressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "GetUpgradeInstanceProgress")
    
    
    return
}

func NewGetUpgradeInstanceProgressResponse() (response *GetUpgradeInstanceProgressResponse) {
    response = &GetUpgradeInstanceProgressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetUpgradeInstanceProgress
// 获得节点升级当前的进度
//
// 可能返回的错误码:
//  INTERNALERROR_TASKNOTFOUND = "InternalError.TaskNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) GetUpgradeInstanceProgress(request *GetUpgradeInstanceProgressRequest) (response *GetUpgradeInstanceProgressResponse, err error) {
    return c.GetUpgradeInstanceProgressWithContext(context.Background(), request)
}

// GetUpgradeInstanceProgress
// 获得节点升级当前的进度
//
// 可能返回的错误码:
//  INTERNALERROR_TASKNOTFOUND = "InternalError.TaskNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) GetUpgradeInstanceProgressWithContext(ctx context.Context, request *GetUpgradeInstanceProgressRequest) (response *GetUpgradeInstanceProgressResponse, err error) {
    if request == nil {
        request = NewGetUpgradeInstanceProgressRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetUpgradeInstanceProgress require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetUpgradeInstanceProgressResponse()
    err = c.Send(request, response)
    return
}

func NewInstallEdgeLogAgentRequest() (request *InstallEdgeLogAgentRequest) {
    request = &InstallEdgeLogAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "InstallEdgeLogAgent")
    
    
    return
}

func NewInstallEdgeLogAgentResponse() (response *InstallEdgeLogAgentResponse) {
    response = &InstallEdgeLogAgentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InstallEdgeLogAgent
// 在tke@edge集群的边缘节点上安装日志采集组件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) InstallEdgeLogAgent(request *InstallEdgeLogAgentRequest) (response *InstallEdgeLogAgentResponse, err error) {
    return c.InstallEdgeLogAgentWithContext(context.Background(), request)
}

// InstallEdgeLogAgent
// 在tke@edge集群的边缘节点上安装日志采集组件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) InstallEdgeLogAgentWithContext(ctx context.Context, request *InstallEdgeLogAgentRequest) (response *InstallEdgeLogAgentResponse, err error) {
    if request == nil {
        request = NewInstallEdgeLogAgentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InstallEdgeLogAgent require credential")
    }

    request.SetContext(ctx)
    
    response = NewInstallEdgeLogAgentResponse()
    err = c.Send(request, response)
    return
}

func NewInstallLogAgentRequest() (request *InstallLogAgentRequest) {
    request = &InstallLogAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "InstallLogAgent")
    
    
    return
}

func NewInstallLogAgentResponse() (response *InstallLogAgentResponse) {
    response = &InstallLogAgentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InstallLogAgent
// 在TKE集群中安装CLS日志采集组件
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_KUBERNETESCREATEOPERATIONERROR = "FailedOperation.KubernetesCreateOperationError"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) InstallLogAgent(request *InstallLogAgentRequest) (response *InstallLogAgentResponse, err error) {
    return c.InstallLogAgentWithContext(context.Background(), request)
}

// InstallLogAgent
// 在TKE集群中安装CLS日志采集组件
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_KUBERNETESCREATEOPERATIONERROR = "FailedOperation.KubernetesCreateOperationError"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) InstallLogAgentWithContext(ctx context.Context, request *InstallLogAgentRequest) (response *InstallLogAgentResponse, err error) {
    if request == nil {
        request = NewInstallLogAgentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InstallLogAgent require credential")
    }

    request.SetContext(ctx)
    
    response = NewInstallLogAgentResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterAsGroupAttributeRequest() (request *ModifyClusterAsGroupAttributeRequest) {
    request = &ModifyClusterAsGroupAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "ModifyClusterAsGroupAttribute")
    
    
    return
}

func NewModifyClusterAsGroupAttributeResponse() (response *ModifyClusterAsGroupAttributeResponse) {
    response = &ModifyClusterAsGroupAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyClusterAsGroupAttribute
// 修改集群伸缩组属性
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_ASCOMMON = "InternalError.AsCommon"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ASCOMMONERROR = "InvalidParameter.AsCommonError"
//  INVALIDPARAMETER_CIDROUTOFROUTETABLE = "InvalidParameter.CidrOutOfRouteTable"
//  INVALIDPARAMETER_GATEWAYALREADYASSOCIATEDCIDR = "InvalidParameter.GatewayAlreadyAssociatedCidr"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_ROUTETABLENOTEMPTY = "InvalidParameter.RouteTableNotEmpty"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyClusterAsGroupAttribute(request *ModifyClusterAsGroupAttributeRequest) (response *ModifyClusterAsGroupAttributeResponse, err error) {
    return c.ModifyClusterAsGroupAttributeWithContext(context.Background(), request)
}

// ModifyClusterAsGroupAttribute
// 修改集群伸缩组属性
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_ASCOMMON = "InternalError.AsCommon"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ASCOMMONERROR = "InvalidParameter.AsCommonError"
//  INVALIDPARAMETER_CIDROUTOFROUTETABLE = "InvalidParameter.CidrOutOfRouteTable"
//  INVALIDPARAMETER_GATEWAYALREADYASSOCIATEDCIDR = "InvalidParameter.GatewayAlreadyAssociatedCidr"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_ROUTETABLENOTEMPTY = "InvalidParameter.RouteTableNotEmpty"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyClusterAsGroupAttributeWithContext(ctx context.Context, request *ModifyClusterAsGroupAttributeRequest) (response *ModifyClusterAsGroupAttributeResponse, err error) {
    if request == nil {
        request = NewModifyClusterAsGroupAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClusterAsGroupAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClusterAsGroupAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterAsGroupOptionAttributeRequest() (request *ModifyClusterAsGroupOptionAttributeRequest) {
    request = &ModifyClusterAsGroupOptionAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "ModifyClusterAsGroupOptionAttribute")
    
    
    return
}

func NewModifyClusterAsGroupOptionAttributeResponse() (response *ModifyClusterAsGroupOptionAttributeResponse) {
    response = &ModifyClusterAsGroupOptionAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyClusterAsGroupOptionAttribute
// 修改集群弹性伸缩属性
//
// 可能返回的错误码:
//  FAILEDOPERATION_KUBERNETESINTERNAL = "FailedOperation.KubernetesInternal"
//  FAILEDOPERATION_UNEXPECTEDERROR = "FailedOperation.UnexpectedError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ASCOMMON = "InternalError.AsCommon"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_CVMCOMMON = "InternalError.CvmCommon"
//  INTERNALERROR_CVMNOTFOUND = "InternalError.CvmNotFound"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyClusterAsGroupOptionAttribute(request *ModifyClusterAsGroupOptionAttributeRequest) (response *ModifyClusterAsGroupOptionAttributeResponse, err error) {
    return c.ModifyClusterAsGroupOptionAttributeWithContext(context.Background(), request)
}

// ModifyClusterAsGroupOptionAttribute
// 修改集群弹性伸缩属性
//
// 可能返回的错误码:
//  FAILEDOPERATION_KUBERNETESINTERNAL = "FailedOperation.KubernetesInternal"
//  FAILEDOPERATION_UNEXPECTEDERROR = "FailedOperation.UnexpectedError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ASCOMMON = "InternalError.AsCommon"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_CVMCOMMON = "InternalError.CvmCommon"
//  INTERNALERROR_CVMNOTFOUND = "InternalError.CvmNotFound"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyClusterAsGroupOptionAttributeWithContext(ctx context.Context, request *ModifyClusterAsGroupOptionAttributeRequest) (response *ModifyClusterAsGroupOptionAttributeResponse, err error) {
    if request == nil {
        request = NewModifyClusterAsGroupOptionAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClusterAsGroupOptionAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClusterAsGroupOptionAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterAttributeRequest() (request *ModifyClusterAttributeRequest) {
    request = &ModifyClusterAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "ModifyClusterAttribute")
    
    
    return
}

func NewModifyClusterAttributeResponse() (response *ModifyClusterAttributeResponse) {
    response = &ModifyClusterAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyClusterAttribute
// 修改集群属性
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_TRADECOMMON = "FailedOperation.TradeCommon"
//  FAILEDOPERATION_UNEXPECTEDERROR = "FailedOperation.UnexpectedError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_TRADECOMMON = "InternalError.TradeCommon"
//  INTERNALERROR_TRADEINSUFFICIENTBALANCE = "InternalError.TradeInsufficientBalance"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) ModifyClusterAttribute(request *ModifyClusterAttributeRequest) (response *ModifyClusterAttributeResponse, err error) {
    return c.ModifyClusterAttributeWithContext(context.Background(), request)
}

// ModifyClusterAttribute
// 修改集群属性
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_TRADECOMMON = "FailedOperation.TradeCommon"
//  FAILEDOPERATION_UNEXPECTEDERROR = "FailedOperation.UnexpectedError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_TRADECOMMON = "InternalError.TradeCommon"
//  INTERNALERROR_TRADEINSUFFICIENTBALANCE = "InternalError.TradeInsufficientBalance"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) ModifyClusterAttributeWithContext(ctx context.Context, request *ModifyClusterAttributeRequest) (response *ModifyClusterAttributeResponse, err error) {
    if request == nil {
        request = NewModifyClusterAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClusterAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClusterAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterAuthenticationOptionsRequest() (request *ModifyClusterAuthenticationOptionsRequest) {
    request = &ModifyClusterAuthenticationOptionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "ModifyClusterAuthenticationOptions")
    
    
    return
}

func NewModifyClusterAuthenticationOptionsResponse() (response *ModifyClusterAuthenticationOptionsResponse) {
    response = &ModifyClusterAuthenticationOptionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyClusterAuthenticationOptions
// 修改集群认证配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) ModifyClusterAuthenticationOptions(request *ModifyClusterAuthenticationOptionsRequest) (response *ModifyClusterAuthenticationOptionsResponse, err error) {
    return c.ModifyClusterAuthenticationOptionsWithContext(context.Background(), request)
}

// ModifyClusterAuthenticationOptions
// 修改集群认证配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) ModifyClusterAuthenticationOptionsWithContext(ctx context.Context, request *ModifyClusterAuthenticationOptionsRequest) (response *ModifyClusterAuthenticationOptionsResponse, err error) {
    if request == nil {
        request = NewModifyClusterAuthenticationOptionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClusterAuthenticationOptions require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClusterAuthenticationOptionsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterEndpointSPRequest() (request *ModifyClusterEndpointSPRequest) {
    request = &ModifyClusterEndpointSPRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "ModifyClusterEndpointSP")
    
    
    return
}

func NewModifyClusterEndpointSPResponse() (response *ModifyClusterEndpointSPResponse) {
    response = &ModifyClusterEndpointSPResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyClusterEndpointSP
// 修改托管集群外网端口的安全策略（老的方式，仅支持托管集群外网端口）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_VPCUNEXPECTEDERROR = "FailedOperation.VPCUnexpectedError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INTERNALERROR_VPCUNEXPECTEDERROR = "InternalError.VPCUnexpectedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyClusterEndpointSP(request *ModifyClusterEndpointSPRequest) (response *ModifyClusterEndpointSPResponse, err error) {
    return c.ModifyClusterEndpointSPWithContext(context.Background(), request)
}

// ModifyClusterEndpointSP
// 修改托管集群外网端口的安全策略（老的方式，仅支持托管集群外网端口）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_VPCUNEXPECTEDERROR = "FailedOperation.VPCUnexpectedError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INTERNALERROR_VPCUNEXPECTEDERROR = "InternalError.VPCUnexpectedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyClusterEndpointSPWithContext(ctx context.Context, request *ModifyClusterEndpointSPRequest) (response *ModifyClusterEndpointSPResponse, err error) {
    if request == nil {
        request = NewModifyClusterEndpointSPRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClusterEndpointSP require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClusterEndpointSPResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterNodePoolRequest() (request *ModifyClusterNodePoolRequest) {
    request = &ModifyClusterNodePoolRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "ModifyClusterNodePool")
    
    
    return
}

func NewModifyClusterNodePoolResponse() (response *ModifyClusterNodePoolResponse) {
    response = &ModifyClusterNodePoolResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyClusterNodePool
// 编辑节点池
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_ASASGNOTEXIST = "ResourceNotFound.AsAsgNotExist"
//  UNSUPPORTEDOPERATION_CAENABLEFAILED = "UnsupportedOperation.CaEnableFailed"
func (c *Client) ModifyClusterNodePool(request *ModifyClusterNodePoolRequest) (response *ModifyClusterNodePoolResponse, err error) {
    return c.ModifyClusterNodePoolWithContext(context.Background(), request)
}

// ModifyClusterNodePool
// 编辑节点池
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_ASASGNOTEXIST = "ResourceNotFound.AsAsgNotExist"
//  UNSUPPORTEDOPERATION_CAENABLEFAILED = "UnsupportedOperation.CaEnableFailed"
func (c *Client) ModifyClusterNodePoolWithContext(ctx context.Context, request *ModifyClusterNodePoolRequest) (response *ModifyClusterNodePoolResponse, err error) {
    if request == nil {
        request = NewModifyClusterNodePoolRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClusterNodePool require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClusterNodePoolResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterVirtualNodePoolRequest() (request *ModifyClusterVirtualNodePoolRequest) {
    request = &ModifyClusterVirtualNodePoolRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "ModifyClusterVirtualNodePool")
    
    
    return
}

func NewModifyClusterVirtualNodePoolResponse() (response *ModifyClusterVirtualNodePoolResponse) {
    response = &ModifyClusterVirtualNodePoolResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyClusterVirtualNodePool
// 修改虚拟节点池
//
// 可能返回的错误码:
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_NODEPOOLSTATENOTNORMAL = "ResourceUnavailable.NodePoolStateNotNormal"
func (c *Client) ModifyClusterVirtualNodePool(request *ModifyClusterVirtualNodePoolRequest) (response *ModifyClusterVirtualNodePoolResponse, err error) {
    return c.ModifyClusterVirtualNodePoolWithContext(context.Background(), request)
}

// ModifyClusterVirtualNodePool
// 修改虚拟节点池
//
// 可能返回的错误码:
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_NODEPOOLSTATENOTNORMAL = "ResourceUnavailable.NodePoolStateNotNormal"
func (c *Client) ModifyClusterVirtualNodePoolWithContext(ctx context.Context, request *ModifyClusterVirtualNodePoolRequest) (response *ModifyClusterVirtualNodePoolResponse, err error) {
    if request == nil {
        request = NewModifyClusterVirtualNodePoolRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClusterVirtualNodePool require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClusterVirtualNodePoolResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNodePoolDesiredCapacityAboutAsgRequest() (request *ModifyNodePoolDesiredCapacityAboutAsgRequest) {
    request = &ModifyNodePoolDesiredCapacityAboutAsgRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "ModifyNodePoolDesiredCapacityAboutAsg")
    
    
    return
}

func NewModifyNodePoolDesiredCapacityAboutAsgResponse() (response *ModifyNodePoolDesiredCapacityAboutAsgResponse) {
    response = &ModifyNodePoolDesiredCapacityAboutAsgResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyNodePoolDesiredCapacityAboutAsg
// 修改节点池关联伸缩组的期望实例数
//
// 可能返回的错误码:
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_ASASGNOTEXIST = "ResourceNotFound.AsAsgNotExist"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyNodePoolDesiredCapacityAboutAsg(request *ModifyNodePoolDesiredCapacityAboutAsgRequest) (response *ModifyNodePoolDesiredCapacityAboutAsgResponse, err error) {
    return c.ModifyNodePoolDesiredCapacityAboutAsgWithContext(context.Background(), request)
}

// ModifyNodePoolDesiredCapacityAboutAsg
// 修改节点池关联伸缩组的期望实例数
//
// 可能返回的错误码:
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_ASASGNOTEXIST = "ResourceNotFound.AsAsgNotExist"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyNodePoolDesiredCapacityAboutAsgWithContext(ctx context.Context, request *ModifyNodePoolDesiredCapacityAboutAsgRequest) (response *ModifyNodePoolDesiredCapacityAboutAsgResponse, err error) {
    if request == nil {
        request = NewModifyNodePoolDesiredCapacityAboutAsgRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNodePoolDesiredCapacityAboutAsg require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNodePoolDesiredCapacityAboutAsgResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNodePoolInstanceTypesRequest() (request *ModifyNodePoolInstanceTypesRequest) {
    request = &ModifyNodePoolInstanceTypesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "ModifyNodePoolInstanceTypes")
    
    
    return
}

func NewModifyNodePoolInstanceTypesResponse() (response *ModifyNodePoolInstanceTypesResponse) {
    response = &ModifyNodePoolInstanceTypesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyNodePoolInstanceTypes
// 修改节点池的机型配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) ModifyNodePoolInstanceTypes(request *ModifyNodePoolInstanceTypesRequest) (response *ModifyNodePoolInstanceTypesResponse, err error) {
    return c.ModifyNodePoolInstanceTypesWithContext(context.Background(), request)
}

// ModifyNodePoolInstanceTypes
// 修改节点池的机型配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) ModifyNodePoolInstanceTypesWithContext(ctx context.Context, request *ModifyNodePoolInstanceTypesRequest) (response *ModifyNodePoolInstanceTypesResponse, err error) {
    if request == nil {
        request = NewModifyNodePoolInstanceTypesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNodePoolInstanceTypes require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNodePoolInstanceTypesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPrometheusAgentExternalLabelsRequest() (request *ModifyPrometheusAgentExternalLabelsRequest) {
    request = &ModifyPrometheusAgentExternalLabelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "ModifyPrometheusAgentExternalLabels")
    
    
    return
}

func NewModifyPrometheusAgentExternalLabelsResponse() (response *ModifyPrometheusAgentExternalLabelsResponse) {
    response = &ModifyPrometheusAgentExternalLabelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyPrometheusAgentExternalLabels
// 修改被关联集群的external labels
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
func (c *Client) ModifyPrometheusAgentExternalLabels(request *ModifyPrometheusAgentExternalLabelsRequest) (response *ModifyPrometheusAgentExternalLabelsResponse, err error) {
    return c.ModifyPrometheusAgentExternalLabelsWithContext(context.Background(), request)
}

// ModifyPrometheusAgentExternalLabels
// 修改被关联集群的external labels
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
func (c *Client) ModifyPrometheusAgentExternalLabelsWithContext(ctx context.Context, request *ModifyPrometheusAgentExternalLabelsRequest) (response *ModifyPrometheusAgentExternalLabelsResponse, err error) {
    if request == nil {
        request = NewModifyPrometheusAgentExternalLabelsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPrometheusAgentExternalLabels require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPrometheusAgentExternalLabelsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPrometheusAlertPolicyRequest() (request *ModifyPrometheusAlertPolicyRequest) {
    request = &ModifyPrometheusAlertPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "ModifyPrometheusAlertPolicy")
    
    
    return
}

func NewModifyPrometheusAlertPolicyResponse() (response *ModifyPrometheusAlertPolicyResponse) {
    response = &ModifyPrometheusAlertPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyPrometheusAlertPolicy
// 修改2.0实例告警策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) ModifyPrometheusAlertPolicy(request *ModifyPrometheusAlertPolicyRequest) (response *ModifyPrometheusAlertPolicyResponse, err error) {
    return c.ModifyPrometheusAlertPolicyWithContext(context.Background(), request)
}

// ModifyPrometheusAlertPolicy
// 修改2.0实例告警策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) ModifyPrometheusAlertPolicyWithContext(ctx context.Context, request *ModifyPrometheusAlertPolicyRequest) (response *ModifyPrometheusAlertPolicyResponse, err error) {
    if request == nil {
        request = NewModifyPrometheusAlertPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPrometheusAlertPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPrometheusAlertPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPrometheusAlertRuleRequest() (request *ModifyPrometheusAlertRuleRequest) {
    request = &ModifyPrometheusAlertRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "ModifyPrometheusAlertRule")
    
    
    return
}

func NewModifyPrometheusAlertRuleResponse() (response *ModifyPrometheusAlertRuleResponse) {
    response = &ModifyPrometheusAlertRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyPrometheusAlertRule
// 修改告警规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) ModifyPrometheusAlertRule(request *ModifyPrometheusAlertRuleRequest) (response *ModifyPrometheusAlertRuleResponse, err error) {
    return c.ModifyPrometheusAlertRuleWithContext(context.Background(), request)
}

// ModifyPrometheusAlertRule
// 修改告警规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) ModifyPrometheusAlertRuleWithContext(ctx context.Context, request *ModifyPrometheusAlertRuleRequest) (response *ModifyPrometheusAlertRuleResponse, err error) {
    if request == nil {
        request = NewModifyPrometheusAlertRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPrometheusAlertRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPrometheusAlertRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPrometheusConfigRequest() (request *ModifyPrometheusConfigRequest) {
    request = &ModifyPrometheusConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "ModifyPrometheusConfig")
    
    
    return
}

func NewModifyPrometheusConfigResponse() (response *ModifyPrometheusConfigResponse) {
    response = &ModifyPrometheusConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyPrometheusConfig
// 修改prometheus配置，如果配置项不存在，则会新增
//
// 可能返回的错误码:
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
func (c *Client) ModifyPrometheusConfig(request *ModifyPrometheusConfigRequest) (response *ModifyPrometheusConfigResponse, err error) {
    return c.ModifyPrometheusConfigWithContext(context.Background(), request)
}

// ModifyPrometheusConfig
// 修改prometheus配置，如果配置项不存在，则会新增
//
// 可能返回的错误码:
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
func (c *Client) ModifyPrometheusConfigWithContext(ctx context.Context, request *ModifyPrometheusConfigRequest) (response *ModifyPrometheusConfigResponse, err error) {
    if request == nil {
        request = NewModifyPrometheusConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPrometheusConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPrometheusConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPrometheusGlobalNotificationRequest() (request *ModifyPrometheusGlobalNotificationRequest) {
    request = &ModifyPrometheusGlobalNotificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "ModifyPrometheusGlobalNotification")
    
    
    return
}

func NewModifyPrometheusGlobalNotificationResponse() (response *ModifyPrometheusGlobalNotificationResponse) {
    response = &ModifyPrometheusGlobalNotificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyPrometheusGlobalNotification
// 修改全局告警通知渠道
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) ModifyPrometheusGlobalNotification(request *ModifyPrometheusGlobalNotificationRequest) (response *ModifyPrometheusGlobalNotificationResponse, err error) {
    return c.ModifyPrometheusGlobalNotificationWithContext(context.Background(), request)
}

// ModifyPrometheusGlobalNotification
// 修改全局告警通知渠道
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) ModifyPrometheusGlobalNotificationWithContext(ctx context.Context, request *ModifyPrometheusGlobalNotificationRequest) (response *ModifyPrometheusGlobalNotificationResponse, err error) {
    if request == nil {
        request = NewModifyPrometheusGlobalNotificationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPrometheusGlobalNotification require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPrometheusGlobalNotificationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPrometheusRecordRuleYamlRequest() (request *ModifyPrometheusRecordRuleYamlRequest) {
    request = &ModifyPrometheusRecordRuleYamlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "ModifyPrometheusRecordRuleYaml")
    
    
    return
}

func NewModifyPrometheusRecordRuleYamlResponse() (response *ModifyPrometheusRecordRuleYamlResponse) {
    response = &ModifyPrometheusRecordRuleYamlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyPrometheusRecordRuleYaml
// 通过yaml的方式修改Prometheus聚合实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) ModifyPrometheusRecordRuleYaml(request *ModifyPrometheusRecordRuleYamlRequest) (response *ModifyPrometheusRecordRuleYamlResponse, err error) {
    return c.ModifyPrometheusRecordRuleYamlWithContext(context.Background(), request)
}

// ModifyPrometheusRecordRuleYaml
// 通过yaml的方式修改Prometheus聚合实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) ModifyPrometheusRecordRuleYamlWithContext(ctx context.Context, request *ModifyPrometheusRecordRuleYamlRequest) (response *ModifyPrometheusRecordRuleYamlResponse, err error) {
    if request == nil {
        request = NewModifyPrometheusRecordRuleYamlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPrometheusRecordRuleYaml require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPrometheusRecordRuleYamlResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPrometheusTempRequest() (request *ModifyPrometheusTempRequest) {
    request = &ModifyPrometheusTempRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "ModifyPrometheusTemp")
    
    
    return
}

func NewModifyPrometheusTempResponse() (response *ModifyPrometheusTempResponse) {
    response = &ModifyPrometheusTempResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyPrometheusTemp
// 修改模板内容
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyPrometheusTemp(request *ModifyPrometheusTempRequest) (response *ModifyPrometheusTempResponse, err error) {
    return c.ModifyPrometheusTempWithContext(context.Background(), request)
}

// ModifyPrometheusTemp
// 修改模板内容
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyPrometheusTempWithContext(ctx context.Context, request *ModifyPrometheusTempRequest) (response *ModifyPrometheusTempResponse, err error) {
    if request == nil {
        request = NewModifyPrometheusTempRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPrometheusTemp require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPrometheusTempResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPrometheusTemplateRequest() (request *ModifyPrometheusTemplateRequest) {
    request = &ModifyPrometheusTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "ModifyPrometheusTemplate")
    
    
    return
}

func NewModifyPrometheusTemplateResponse() (response *ModifyPrometheusTemplateResponse) {
    response = &ModifyPrometheusTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyPrometheusTemplate
// 修改模板内容
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
func (c *Client) ModifyPrometheusTemplate(request *ModifyPrometheusTemplateRequest) (response *ModifyPrometheusTemplateResponse, err error) {
    return c.ModifyPrometheusTemplateWithContext(context.Background(), request)
}

// ModifyPrometheusTemplate
// 修改模板内容
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
func (c *Client) ModifyPrometheusTemplateWithContext(ctx context.Context, request *ModifyPrometheusTemplateRequest) (response *ModifyPrometheusTemplateResponse, err error) {
    if request == nil {
        request = NewModifyPrometheusTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPrometheusTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPrometheusTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveNodeFromNodePoolRequest() (request *RemoveNodeFromNodePoolRequest) {
    request = &RemoveNodeFromNodePoolRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "RemoveNodeFromNodePool")
    
    
    return
}

func NewRemoveNodeFromNodePoolResponse() (response *RemoveNodeFromNodePoolResponse) {
    response = &RemoveNodeFromNodePoolResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RemoveNodeFromNodePool
// 移出节点池节点，但保留在集群内
//
// 可能返回的错误码:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) RemoveNodeFromNodePool(request *RemoveNodeFromNodePoolRequest) (response *RemoveNodeFromNodePoolResponse, err error) {
    return c.RemoveNodeFromNodePoolWithContext(context.Background(), request)
}

// RemoveNodeFromNodePool
// 移出节点池节点，但保留在集群内
//
// 可能返回的错误码:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) RemoveNodeFromNodePoolWithContext(ctx context.Context, request *RemoveNodeFromNodePoolRequest) (response *RemoveNodeFromNodePoolResponse, err error) {
    if request == nil {
        request = NewRemoveNodeFromNodePoolRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveNodeFromNodePool require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveNodeFromNodePoolResponse()
    err = c.Send(request, response)
    return
}

func NewRestartEKSContainerInstancesRequest() (request *RestartEKSContainerInstancesRequest) {
    request = &RestartEKSContainerInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "RestartEKSContainerInstances")
    
    
    return
}

func NewRestartEKSContainerInstancesResponse() (response *RestartEKSContainerInstancesResponse) {
    response = &RestartEKSContainerInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RestartEKSContainerInstances
// 重启弹性容器实例，支持批量操作
//
// 可能返回的错误码:
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RestartEKSContainerInstances(request *RestartEKSContainerInstancesRequest) (response *RestartEKSContainerInstancesResponse, err error) {
    return c.RestartEKSContainerInstancesWithContext(context.Background(), request)
}

// RestartEKSContainerInstances
// 重启弹性容器实例，支持批量操作
//
// 可能返回的错误码:
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RestartEKSContainerInstancesWithContext(ctx context.Context, request *RestartEKSContainerInstancesRequest) (response *RestartEKSContainerInstancesResponse, err error) {
    if request == nil {
        request = NewRestartEKSContainerInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestartEKSContainerInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewRestartEKSContainerInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewRollbackClusterReleaseRequest() (request *RollbackClusterReleaseRequest) {
    request = &RollbackClusterReleaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "RollbackClusterRelease")
    
    
    return
}

func NewRollbackClusterReleaseResponse() (response *RollbackClusterReleaseResponse) {
    response = &RollbackClusterReleaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RollbackClusterRelease
// 在应用市场中集群回滚应用至某个历史版本
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MARKETRELEASEOPERATION = "FailedOperation.MarketReleaseOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MARKETGETAUTHFAILED = "InternalError.MarketGetAuthFailed"
//  INTERNALERROR_MARKETINTERNALSERVERERROR = "InternalError.MarketInternalServerError"
//  INTERNALERROR_MARKETRELEASEOPERATION = "InternalError.MarketReleaseOperation"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) RollbackClusterRelease(request *RollbackClusterReleaseRequest) (response *RollbackClusterReleaseResponse, err error) {
    return c.RollbackClusterReleaseWithContext(context.Background(), request)
}

// RollbackClusterRelease
// 在应用市场中集群回滚应用至某个历史版本
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MARKETRELEASEOPERATION = "FailedOperation.MarketReleaseOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MARKETGETAUTHFAILED = "InternalError.MarketGetAuthFailed"
//  INTERNALERROR_MARKETINTERNALSERVERERROR = "InternalError.MarketInternalServerError"
//  INTERNALERROR_MARKETRELEASEOPERATION = "InternalError.MarketReleaseOperation"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) RollbackClusterReleaseWithContext(ctx context.Context, request *RollbackClusterReleaseRequest) (response *RollbackClusterReleaseResponse, err error) {
    if request == nil {
        request = NewRollbackClusterReleaseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RollbackClusterRelease require credential")
    }

    request.SetContext(ctx)
    
    response = NewRollbackClusterReleaseResponse()
    err = c.Send(request, response)
    return
}

func NewRunPrometheusInstanceRequest() (request *RunPrometheusInstanceRequest) {
    request = &RunPrometheusInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "RunPrometheusInstance")
    
    
    return
}

func NewRunPrometheusInstanceResponse() (response *RunPrometheusInstanceResponse) {
    response = &RunPrometheusInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RunPrometheusInstance
// 初始化TMP实例，开启集成中心时调用
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) RunPrometheusInstance(request *RunPrometheusInstanceRequest) (response *RunPrometheusInstanceResponse, err error) {
    return c.RunPrometheusInstanceWithContext(context.Background(), request)
}

// RunPrometheusInstance
// 初始化TMP实例，开启集成中心时调用
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) RunPrometheusInstanceWithContext(ctx context.Context, request *RunPrometheusInstanceRequest) (response *RunPrometheusInstanceResponse, err error) {
    if request == nil {
        request = NewRunPrometheusInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunPrometheusInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewRunPrometheusInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewScaleInClusterMasterRequest() (request *ScaleInClusterMasterRequest) {
    request = &ScaleInClusterMasterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "ScaleInClusterMaster")
    
    
    return
}

func NewScaleInClusterMasterResponse() (response *ScaleInClusterMasterResponse) {
    response = &ScaleInClusterMasterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ScaleInClusterMaster
// 缩容独立集群master节点
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ScaleInClusterMaster(request *ScaleInClusterMasterRequest) (response *ScaleInClusterMasterResponse, err error) {
    return c.ScaleInClusterMasterWithContext(context.Background(), request)
}

// ScaleInClusterMaster
// 缩容独立集群master节点
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ScaleInClusterMasterWithContext(ctx context.Context, request *ScaleInClusterMasterRequest) (response *ScaleInClusterMasterResponse, err error) {
    if request == nil {
        request = NewScaleInClusterMasterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ScaleInClusterMaster require credential")
    }

    request.SetContext(ctx)
    
    response = NewScaleInClusterMasterResponse()
    err = c.Send(request, response)
    return
}

func NewScaleOutClusterMasterRequest() (request *ScaleOutClusterMasterRequest) {
    request = &ScaleOutClusterMasterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "ScaleOutClusterMaster")
    
    
    return
}

func NewScaleOutClusterMasterResponse() (response *ScaleOutClusterMasterResponse) {
    response = &ScaleOutClusterMasterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ScaleOutClusterMaster
// 扩容独立集群master节点
//
// 可能返回的错误码:
//  FAILEDOPERATION_CVMUNEXPECTEDERROR = "FailedOperation.CVMUnexpectedError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) ScaleOutClusterMaster(request *ScaleOutClusterMasterRequest) (response *ScaleOutClusterMasterResponse, err error) {
    return c.ScaleOutClusterMasterWithContext(context.Background(), request)
}

// ScaleOutClusterMaster
// 扩容独立集群master节点
//
// 可能返回的错误码:
//  FAILEDOPERATION_CVMUNEXPECTEDERROR = "FailedOperation.CVMUnexpectedError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) ScaleOutClusterMasterWithContext(ctx context.Context, request *ScaleOutClusterMasterRequest) (response *ScaleOutClusterMasterResponse, err error) {
    if request == nil {
        request = NewScaleOutClusterMasterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ScaleOutClusterMaster require credential")
    }

    request.SetContext(ctx)
    
    response = NewScaleOutClusterMasterResponse()
    err = c.Send(request, response)
    return
}

func NewSetNodePoolNodeProtectionRequest() (request *SetNodePoolNodeProtectionRequest) {
    request = &SetNodePoolNodeProtectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "SetNodePoolNodeProtection")
    
    
    return
}

func NewSetNodePoolNodeProtectionResponse() (response *SetNodePoolNodeProtectionResponse) {
    response = &SetNodePoolNodeProtectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetNodePoolNodeProtection
// 仅能设置节点池中处于伸缩组的节点
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) SetNodePoolNodeProtection(request *SetNodePoolNodeProtectionRequest) (response *SetNodePoolNodeProtectionResponse, err error) {
    return c.SetNodePoolNodeProtectionWithContext(context.Background(), request)
}

// SetNodePoolNodeProtection
// 仅能设置节点池中处于伸缩组的节点
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) SetNodePoolNodeProtectionWithContext(ctx context.Context, request *SetNodePoolNodeProtectionRequest) (response *SetNodePoolNodeProtectionResponse, err error) {
    if request == nil {
        request = NewSetNodePoolNodeProtectionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetNodePoolNodeProtection require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetNodePoolNodeProtectionResponse()
    err = c.Send(request, response)
    return
}

func NewSyncPrometheusTempRequest() (request *SyncPrometheusTempRequest) {
    request = &SyncPrometheusTempRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "SyncPrometheusTemp")
    
    
    return
}

func NewSyncPrometheusTempResponse() (response *SyncPrometheusTempResponse) {
    response = &SyncPrometheusTempResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SyncPrometheusTemp
// 同步模板到实例或者集群，针对V2版本实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SyncPrometheusTemp(request *SyncPrometheusTempRequest) (response *SyncPrometheusTempResponse, err error) {
    return c.SyncPrometheusTempWithContext(context.Background(), request)
}

// SyncPrometheusTemp
// 同步模板到实例或者集群，针对V2版本实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SyncPrometheusTempWithContext(ctx context.Context, request *SyncPrometheusTempRequest) (response *SyncPrometheusTempResponse, err error) {
    if request == nil {
        request = NewSyncPrometheusTempRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SyncPrometheusTemp require credential")
    }

    request.SetContext(ctx)
    
    response = NewSyncPrometheusTempResponse()
    err = c.Send(request, response)
    return
}

func NewSyncPrometheusTemplateRequest() (request *SyncPrometheusTemplateRequest) {
    request = &SyncPrometheusTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "SyncPrometheusTemplate")
    
    
    return
}

func NewSyncPrometheusTemplateResponse() (response *SyncPrometheusTemplateResponse) {
    response = &SyncPrometheusTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SyncPrometheusTemplate
// 同步模板到实例或者集群
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SyncPrometheusTemplate(request *SyncPrometheusTemplateRequest) (response *SyncPrometheusTemplateResponse, err error) {
    return c.SyncPrometheusTemplateWithContext(context.Background(), request)
}

// SyncPrometheusTemplate
// 同步模板到实例或者集群
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SyncPrometheusTemplateWithContext(ctx context.Context, request *SyncPrometheusTemplateRequest) (response *SyncPrometheusTemplateResponse, err error) {
    if request == nil {
        request = NewSyncPrometheusTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SyncPrometheusTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewSyncPrometheusTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewUninstallClusterReleaseRequest() (request *UninstallClusterReleaseRequest) {
    request = &UninstallClusterReleaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "UninstallClusterRelease")
    
    
    return
}

func NewUninstallClusterReleaseResponse() (response *UninstallClusterReleaseResponse) {
    response = &UninstallClusterReleaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UninstallClusterRelease
// 在应用市场中集群删除某个应用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MARKETGETAUTHFAILED = "FailedOperation.MarketGetAuthFailed"
//  FAILEDOPERATION_MARKETRELEASEOPERATION = "FailedOperation.MarketReleaseOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MARKETGETAUTHFAILED = "InternalError.MarketGetAuthFailed"
//  INTERNALERROR_MARKETINTERNALSERVERERROR = "InternalError.MarketInternalServerError"
//  INTERNALERROR_MARKETRELEASEOPERATION = "InternalError.MarketReleaseOperation"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) UninstallClusterRelease(request *UninstallClusterReleaseRequest) (response *UninstallClusterReleaseResponse, err error) {
    return c.UninstallClusterReleaseWithContext(context.Background(), request)
}

// UninstallClusterRelease
// 在应用市场中集群删除某个应用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MARKETGETAUTHFAILED = "FailedOperation.MarketGetAuthFailed"
//  FAILEDOPERATION_MARKETRELEASEOPERATION = "FailedOperation.MarketReleaseOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MARKETGETAUTHFAILED = "InternalError.MarketGetAuthFailed"
//  INTERNALERROR_MARKETINTERNALSERVERERROR = "InternalError.MarketInternalServerError"
//  INTERNALERROR_MARKETRELEASEOPERATION = "InternalError.MarketReleaseOperation"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) UninstallClusterReleaseWithContext(ctx context.Context, request *UninstallClusterReleaseRequest) (response *UninstallClusterReleaseResponse, err error) {
    if request == nil {
        request = NewUninstallClusterReleaseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UninstallClusterRelease require credential")
    }

    request.SetContext(ctx)
    
    response = NewUninstallClusterReleaseResponse()
    err = c.Send(request, response)
    return
}

func NewUninstallEdgeLogAgentRequest() (request *UninstallEdgeLogAgentRequest) {
    request = &UninstallEdgeLogAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "UninstallEdgeLogAgent")
    
    
    return
}

func NewUninstallEdgeLogAgentResponse() (response *UninstallEdgeLogAgentResponse) {
    response = &UninstallEdgeLogAgentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UninstallEdgeLogAgent
// 从tke@edge集群边缘节点上卸载日志采集组件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UninstallEdgeLogAgent(request *UninstallEdgeLogAgentRequest) (response *UninstallEdgeLogAgentResponse, err error) {
    return c.UninstallEdgeLogAgentWithContext(context.Background(), request)
}

// UninstallEdgeLogAgent
// 从tke@edge集群边缘节点上卸载日志采集组件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UninstallEdgeLogAgentWithContext(ctx context.Context, request *UninstallEdgeLogAgentRequest) (response *UninstallEdgeLogAgentResponse, err error) {
    if request == nil {
        request = NewUninstallEdgeLogAgentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UninstallEdgeLogAgent require credential")
    }

    request.SetContext(ctx)
    
    response = NewUninstallEdgeLogAgentResponse()
    err = c.Send(request, response)
    return
}

func NewUninstallLogAgentRequest() (request *UninstallLogAgentRequest) {
    request = &UninstallLogAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "UninstallLogAgent")
    
    
    return
}

func NewUninstallLogAgentResponse() (response *UninstallLogAgentResponse) {
    response = &UninstallLogAgentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UninstallLogAgent
// 从TKE集群中卸载CLS日志采集组件
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLUSTERSTATE = "FailedOperation.ClusterState"
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  FAILEDOPERATION_KUBERNETESDELETEOPERATIONERROR = "FailedOperation.KubernetesDeleteOperationError"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) UninstallLogAgent(request *UninstallLogAgentRequest) (response *UninstallLogAgentResponse, err error) {
    return c.UninstallLogAgentWithContext(context.Background(), request)
}

// UninstallLogAgent
// 从TKE集群中卸载CLS日志采集组件
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLUSTERSTATE = "FailedOperation.ClusterState"
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  FAILEDOPERATION_KUBERNETESDELETEOPERATIONERROR = "FailedOperation.KubernetesDeleteOperationError"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) UninstallLogAgentWithContext(ctx context.Context, request *UninstallLogAgentRequest) (response *UninstallLogAgentResponse, err error) {
    if request == nil {
        request = NewUninstallLogAgentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UninstallLogAgent require credential")
    }

    request.SetContext(ctx)
    
    response = NewUninstallLogAgentResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateClusterVersionRequest() (request *UpdateClusterVersionRequest) {
    request = &UpdateClusterVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "UpdateClusterVersion")
    
    
    return
}

func NewUpdateClusterVersionResponse() (response *UpdateClusterVersionResponse) {
    response = &UpdateClusterVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateClusterVersion
// 升级集群 Master 组件到指定版本
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLUSTERUPGRADENODEVERSION = "FailedOperation.ClusterUpgradeNodeVersion"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CLUSTERUPGRADENODEVERSION = "InternalError.ClusterUpgradeNodeVersion"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) UpdateClusterVersion(request *UpdateClusterVersionRequest) (response *UpdateClusterVersionResponse, err error) {
    return c.UpdateClusterVersionWithContext(context.Background(), request)
}

// UpdateClusterVersion
// 升级集群 Master 组件到指定版本
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLUSTERUPGRADENODEVERSION = "FailedOperation.ClusterUpgradeNodeVersion"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CLUSTERUPGRADENODEVERSION = "InternalError.ClusterUpgradeNodeVersion"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) UpdateClusterVersionWithContext(ctx context.Context, request *UpdateClusterVersionRequest) (response *UpdateClusterVersionResponse, err error) {
    if request == nil {
        request = NewUpdateClusterVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateClusterVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateClusterVersionResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateEKSClusterRequest() (request *UpdateEKSClusterRequest) {
    request = &UpdateEKSClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "UpdateEKSCluster")
    
    
    return
}

func NewUpdateEKSClusterResponse() (response *UpdateEKSClusterResponse) {
    response = &UpdateEKSClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateEKSCluster
// 修改弹性集群名称等属性 
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateEKSCluster(request *UpdateEKSClusterRequest) (response *UpdateEKSClusterResponse, err error) {
    return c.UpdateEKSClusterWithContext(context.Background(), request)
}

// UpdateEKSCluster
// 修改弹性集群名称等属性 
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateEKSClusterWithContext(ctx context.Context, request *UpdateEKSClusterRequest) (response *UpdateEKSClusterResponse, err error) {
    if request == nil {
        request = NewUpdateEKSClusterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateEKSCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateEKSClusterResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateEKSContainerInstanceRequest() (request *UpdateEKSContainerInstanceRequest) {
    request = &UpdateEKSContainerInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "UpdateEKSContainerInstance")
    
    
    return
}

func NewUpdateEKSContainerInstanceResponse() (response *UpdateEKSContainerInstanceResponse) {
    response = &UpdateEKSContainerInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateEKSContainerInstance
// 更新容器实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateEKSContainerInstance(request *UpdateEKSContainerInstanceRequest) (response *UpdateEKSContainerInstanceResponse, err error) {
    return c.UpdateEKSContainerInstanceWithContext(context.Background(), request)
}

// UpdateEKSContainerInstance
// 更新容器实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateEKSContainerInstanceWithContext(ctx context.Context, request *UpdateEKSContainerInstanceRequest) (response *UpdateEKSContainerInstanceResponse, err error) {
    if request == nil {
        request = NewUpdateEKSContainerInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateEKSContainerInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateEKSContainerInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateEdgeClusterVersionRequest() (request *UpdateEdgeClusterVersionRequest) {
    request = &UpdateEdgeClusterVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "UpdateEdgeClusterVersion")
    
    
    return
}

func NewUpdateEdgeClusterVersionResponse() (response *UpdateEdgeClusterVersionResponse) {
    response = &UpdateEdgeClusterVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateEdgeClusterVersion
// 升级边缘集群组件到指定版本，此版本为TKEEdge专用版本。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLUSTERUPGRADENODEVERSION = "FailedOperation.ClusterUpgradeNodeVersion"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CLUSTERUPGRADENODEVERSION = "InternalError.ClusterUpgradeNodeVersion"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) UpdateEdgeClusterVersion(request *UpdateEdgeClusterVersionRequest) (response *UpdateEdgeClusterVersionResponse, err error) {
    return c.UpdateEdgeClusterVersionWithContext(context.Background(), request)
}

// UpdateEdgeClusterVersion
// 升级边缘集群组件到指定版本，此版本为TKEEdge专用版本。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLUSTERUPGRADENODEVERSION = "FailedOperation.ClusterUpgradeNodeVersion"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CLUSTERUPGRADENODEVERSION = "InternalError.ClusterUpgradeNodeVersion"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) UpdateEdgeClusterVersionWithContext(ctx context.Context, request *UpdateEdgeClusterVersionRequest) (response *UpdateEdgeClusterVersionResponse, err error) {
    if request == nil {
        request = NewUpdateEdgeClusterVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateEdgeClusterVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateEdgeClusterVersionResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateImageCacheRequest() (request *UpdateImageCacheRequest) {
    request = &UpdateImageCacheRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "UpdateImageCache")
    
    
    return
}

func NewUpdateImageCacheResponse() (response *UpdateImageCacheResponse) {
    response = &UpdateImageCacheResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateImageCache
// 更新镜像缓存接口
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateImageCache(request *UpdateImageCacheRequest) (response *UpdateImageCacheResponse, err error) {
    return c.UpdateImageCacheWithContext(context.Background(), request)
}

// UpdateImageCache
// 更新镜像缓存接口
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateImageCacheWithContext(ctx context.Context, request *UpdateImageCacheRequest) (response *UpdateImageCacheResponse, err error) {
    if request == nil {
        request = NewUpdateImageCacheRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateImageCache require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateImageCacheResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateTKEEdgeClusterRequest() (request *UpdateTKEEdgeClusterRequest) {
    request = &UpdateTKEEdgeClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "UpdateTKEEdgeCluster")
    
    
    return
}

func NewUpdateTKEEdgeClusterResponse() (response *UpdateTKEEdgeClusterResponse) {
    response = &UpdateTKEEdgeClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateTKEEdgeCluster
// 修改边缘计算集群名称等属性 
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateTKEEdgeCluster(request *UpdateTKEEdgeClusterRequest) (response *UpdateTKEEdgeClusterResponse, err error) {
    return c.UpdateTKEEdgeClusterWithContext(context.Background(), request)
}

// UpdateTKEEdgeCluster
// 修改边缘计算集群名称等属性 
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateTKEEdgeClusterWithContext(ctx context.Context, request *UpdateTKEEdgeClusterRequest) (response *UpdateTKEEdgeClusterResponse, err error) {
    if request == nil {
        request = NewUpdateTKEEdgeClusterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateTKEEdgeCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateTKEEdgeClusterResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeClusterInstancesRequest() (request *UpgradeClusterInstancesRequest) {
    request = &UpgradeClusterInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "UpgradeClusterInstances")
    
    
    return
}

func NewUpgradeClusterInstancesResponse() (response *UpgradeClusterInstancesResponse) {
    response = &UpgradeClusterInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpgradeClusterInstances
// 给集群的一批work节点进行升级
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKALREADYRUNNING = "FailedOperation.TaskAlreadyRunning"
//  FAILEDOPERATION_TASKLIFESTATEERROR = "FailedOperation.TaskLifeStateError"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_TASKALREADYRUNNING = "InternalError.TaskAlreadyRunning"
//  INTERNALERROR_TASKLIFESTATEERROR = "InternalError.TaskLifeStateError"
//  INTERNALERROR_TASKNOTFOUND = "InternalError.TaskNotFound"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) UpgradeClusterInstances(request *UpgradeClusterInstancesRequest) (response *UpgradeClusterInstancesResponse, err error) {
    return c.UpgradeClusterInstancesWithContext(context.Background(), request)
}

// UpgradeClusterInstances
// 给集群的一批work节点进行升级
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKALREADYRUNNING = "FailedOperation.TaskAlreadyRunning"
//  FAILEDOPERATION_TASKLIFESTATEERROR = "FailedOperation.TaskLifeStateError"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_TASKALREADYRUNNING = "InternalError.TaskAlreadyRunning"
//  INTERNALERROR_TASKLIFESTATEERROR = "InternalError.TaskLifeStateError"
//  INTERNALERROR_TASKNOTFOUND = "InternalError.TaskNotFound"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) UpgradeClusterInstancesWithContext(ctx context.Context, request *UpgradeClusterInstancesRequest) (response *UpgradeClusterInstancesResponse, err error) {
    if request == nil {
        request = NewUpgradeClusterInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeClusterInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeClusterInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeClusterReleaseRequest() (request *UpgradeClusterReleaseRequest) {
    request = &UpgradeClusterReleaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "UpgradeClusterRelease")
    
    
    return
}

func NewUpgradeClusterReleaseResponse() (response *UpgradeClusterReleaseResponse) {
    response = &UpgradeClusterReleaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpgradeClusterRelease
// 升级集群中已安装的应用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MARKETGETAUTHFAILED = "FailedOperation.MarketGetAuthFailed"
//  FAILEDOPERATION_MARKETRELEASEOPERATION = "FailedOperation.MarketReleaseOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MARKETGETAUTHFAILED = "InternalError.MarketGetAuthFailed"
//  INTERNALERROR_MARKETINTERNALSERVERERROR = "InternalError.MarketInternalServerError"
//  INTERNALERROR_MARKETRELEASEOPERATION = "InternalError.MarketReleaseOperation"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) UpgradeClusterRelease(request *UpgradeClusterReleaseRequest) (response *UpgradeClusterReleaseResponse, err error) {
    return c.UpgradeClusterReleaseWithContext(context.Background(), request)
}

// UpgradeClusterRelease
// 升级集群中已安装的应用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MARKETGETAUTHFAILED = "FailedOperation.MarketGetAuthFailed"
//  FAILEDOPERATION_MARKETRELEASEOPERATION = "FailedOperation.MarketReleaseOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MARKETGETAUTHFAILED = "InternalError.MarketGetAuthFailed"
//  INTERNALERROR_MARKETINTERNALSERVERERROR = "InternalError.MarketInternalServerError"
//  INTERNALERROR_MARKETRELEASEOPERATION = "InternalError.MarketReleaseOperation"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) UpgradeClusterReleaseWithContext(ctx context.Context, request *UpgradeClusterReleaseRequest) (response *UpgradeClusterReleaseResponse, err error) {
    if request == nil {
        request = NewUpgradeClusterReleaseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeClusterRelease require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeClusterReleaseResponse()
    err = c.Send(request, response)
    return
}
