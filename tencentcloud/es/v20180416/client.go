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
func (c *Client) CreateInstance(request *CreateInstanceRequest) (response *CreateInstanceResponse, err error) {
    return c.CreateInstanceWithContext(context.Background(), request)
}

// CreateInstance
// 创建指定规格的ES集群实例
//
// 可能返回的错误码:
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteInstance(request *DeleteInstanceRequest) (response *DeleteInstanceResponse, err error) {
    return c.DeleteInstanceWithContext(context.Background(), request)
}

// DeleteInstance
// 销毁集群实例 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeInstanceLogs(request *DescribeInstanceLogsRequest) (response *DescribeInstanceLogsResponse, err error) {
    return c.DescribeInstanceLogsWithContext(context.Background(), request)
}

// DescribeInstanceLogs
// 查询用户该地域下符合条件的ES集群的日志
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeInstanceOperations(request *DescribeInstanceOperationsRequest) (response *DescribeInstanceOperationsResponse, err error) {
    return c.DescribeInstanceOperationsWithContext(context.Background(), request)
}

// DescribeInstanceOperations
// 查询实例指定条件下的操作记录
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
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
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    return c.DescribeInstancesWithContext(context.Background(), request)
}

// DescribeInstances
// 查询用户该地域下符合条件的所有实例
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeViews(request *DescribeViewsRequest) (response *DescribeViewsResponse, err error) {
    return c.DescribeViewsWithContext(context.Background(), request)
}

// DescribeViews
// 查询集群各视图数据，包括集群维度、节点维度、Kibana维度
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
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
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DiagnoseInstance(request *DiagnoseInstanceRequest) (response *DiagnoseInstanceResponse, err error) {
    return c.DiagnoseInstanceWithContext(context.Background(), request)
}

// DiagnoseInstance
// 智能运维诊断集群
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetRequestTargetNodeTypes(request *GetRequestTargetNodeTypesRequest) (response *GetRequestTargetNodeTypesResponse, err error) {
    return c.GetRequestTargetNodeTypesWithContext(context.Background(), request)
}

// GetRequestTargetNodeTypes
// 获取接收客户端请求的节点类型
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) RestartInstance(request *RestartInstanceRequest) (response *RestartInstanceResponse, err error) {
    return c.RestartInstanceWithContext(context.Background(), request)
}

// RestartInstance
// 重启ES集群实例(用于系统版本更新等操作) 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
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
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) RestartKibana(request *RestartKibanaRequest) (response *RestartKibanaResponse, err error) {
    return c.RestartKibanaWithContext(context.Background(), request)
}

// RestartKibana
// 重启Kibana 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
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
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RestartNodes(request *RestartNodesRequest) (response *RestartNodesResponse, err error) {
    return c.RestartNodesWithContext(context.Background(), request)
}

// RestartNodes
// 用于重启集群节点
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateDiagnoseSettings(request *UpdateDiagnoseSettingsRequest) (response *UpdateDiagnoseSettingsResponse, err error) {
    return c.UpdateDiagnoseSettingsWithContext(context.Background(), request)
}

// UpdateDiagnoseSettings
// 更新智能运维配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) UpdateDictionaries(request *UpdateDictionariesRequest) (response *UpdateDictionariesResponse, err error) {
    return c.UpdateDictionariesWithContext(context.Background(), request)
}

// UpdateDictionaries
// 更新ES集群词典
//
// 可能返回的错误码:
//  RESOURCEINUSE = "ResourceInUse"
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
//  FAILEDOPERATION_CLUSTERRESOURCELIMITERROR = "FailedOperation.ClusterResourceLimitError"
//  FAILEDOPERATION_DISKCOUNTPARAMERROR = "FailedOperation.DiskCountParamError"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  FAILEDOPERATION_UNSUPPORTRESETNODETYPEANDSCALEOUTDISK = "FailedOperation.UnsupportResetNodeTypeAndScaleoutDisk"
//  FAILEDOPERATION_UNSUPPORTRESETSCALEDOWNANDMODIFYDISK = "FailedOperation.UnsupportResetScaledownAndModifyDisk"
//  FAILEDOPERATION_UNSUPPORTREVERSEREGULATIONNODETYPEANDDISK = "FailedOperation.UnsupportReverseRegulationNodeTypeAndDisk"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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
//  FAILEDOPERATION_CLUSTERRESOURCELIMITERROR = "FailedOperation.ClusterResourceLimitError"
//  FAILEDOPERATION_DISKCOUNTPARAMERROR = "FailedOperation.DiskCountParamError"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  FAILEDOPERATION_UNSUPPORTRESETNODETYPEANDSCALEOUTDISK = "FailedOperation.UnsupportResetNodeTypeAndScaleoutDisk"
//  FAILEDOPERATION_UNSUPPORTRESETSCALEDOWNANDMODIFYDISK = "FailedOperation.UnsupportResetScaledownAndModifyDisk"
//  FAILEDOPERATION_UNSUPPORTREVERSEREGULATIONNODETYPEANDDISK = "FailedOperation.UnsupportReverseRegulationNodeTypeAndDisk"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATENOREPLICATION = "FailedOperation.ErrorClusterStateNoReplication"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateJdk(request *UpdateJdkRequest) (response *UpdateJdkResponse, err error) {
    return c.UpdateJdkWithContext(context.Background(), request)
}

// UpdateJdk
// 更新实例Jdk配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATENOREPLICATION = "FailedOperation.ErrorClusterStateNoReplication"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATENOREPLICATION = "FailedOperation.ErrorClusterStateNoReplication"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdatePlugins(request *UpdatePluginsRequest) (response *UpdatePluginsResponse, err error) {
    return c.UpdatePluginsWithContext(context.Background(), request)
}

// UpdatePlugins
// 变更插件列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATENOREPLICATION = "FailedOperation.ErrorClusterStateNoReplication"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateRequestTargetNodeTypes(request *UpdateRequestTargetNodeTypesRequest) (response *UpdateRequestTargetNodeTypesResponse, err error) {
    return c.UpdateRequestTargetNodeTypesWithContext(context.Background(), request)
}

// UpdateRequestTargetNodeTypes
// 更新接收客户端请求的节点类型
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpgradeLicense(request *UpgradeLicenseRequest) (response *UpgradeLicenseResponse, err error) {
    return c.UpgradeLicenseWithContext(context.Background(), request)
}

// UpgradeLicense
// 升级ES商业特性
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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
