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

package v20190904

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-09-04"

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


func NewCreateAcRulesRequest() (request *CreateAcRulesRequest) {
    request = &CreateAcRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "CreateAcRules")
    return
}

func NewCreateAcRulesResponse() (response *CreateAcRulesResponse) {
    response = &CreateAcRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAcRules
// 创建规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) CreateAcRules(request *CreateAcRulesRequest) (response *CreateAcRulesResponse, err error) {
    if request == nil {
        request = NewCreateAcRulesRequest()
    }
    response = NewCreateAcRulesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSecurityGroupApiRulesRequest() (request *CreateSecurityGroupApiRulesRequest) {
    request = &CreateSecurityGroupApiRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "CreateSecurityGroupApiRules")
    return
}

func NewCreateSecurityGroupApiRulesResponse() (response *CreateSecurityGroupApiRulesResponse) {
    response = &CreateSecurityGroupApiRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSecurityGroupApiRules
// 创建安全组API规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateSecurityGroupApiRules(request *CreateSecurityGroupApiRulesRequest) (response *CreateSecurityGroupApiRulesResponse, err error) {
    if request == nil {
        request = NewCreateSecurityGroupApiRulesRequest()
    }
    response = NewCreateSecurityGroupApiRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAcRuleRequest() (request *DeleteAcRuleRequest) {
    request = &DeleteAcRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "DeleteAcRule")
    return
}

func NewDeleteAcRuleResponse() (response *DeleteAcRuleResponse) {
    response = &DeleteAcRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAcRule
// 删除规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteAcRule(request *DeleteAcRuleRequest) (response *DeleteAcRuleResponse, err error) {
    if request == nil {
        request = NewDeleteAcRuleRequest()
    }
    response = NewDeleteAcRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAllAccessControlRuleRequest() (request *DeleteAllAccessControlRuleRequest) {
    request = &DeleteAllAccessControlRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "DeleteAllAccessControlRule")
    return
}

func NewDeleteAllAccessControlRuleResponse() (response *DeleteAllAccessControlRuleResponse) {
    response = &DeleteAllAccessControlRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAllAccessControlRule
// 全部删除规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteAllAccessControlRule(request *DeleteAllAccessControlRuleRequest) (response *DeleteAllAccessControlRuleResponse, err error) {
    if request == nil {
        request = NewDeleteAllAccessControlRuleRequest()
    }
    response = NewDeleteAllAccessControlRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSecurityGroupAllRuleRequest() (request *DeleteSecurityGroupAllRuleRequest) {
    request = &DeleteSecurityGroupAllRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "DeleteSecurityGroupAllRule")
    return
}

func NewDeleteSecurityGroupAllRuleResponse() (response *DeleteSecurityGroupAllRuleResponse) {
    response = &DeleteSecurityGroupAllRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteSecurityGroupAllRule
// 删除全部规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteSecurityGroupAllRule(request *DeleteSecurityGroupAllRuleRequest) (response *DeleteSecurityGroupAllRuleResponse, err error) {
    if request == nil {
        request = NewDeleteSecurityGroupAllRuleRequest()
    }
    response = NewDeleteSecurityGroupAllRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSecurityGroupRuleRequest() (request *DeleteSecurityGroupRuleRequest) {
    request = &DeleteSecurityGroupRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "DeleteSecurityGroupRule")
    return
}

func NewDeleteSecurityGroupRuleResponse() (response *DeleteSecurityGroupRuleResponse) {
    response = &DeleteSecurityGroupRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteSecurityGroupRule
// 删除规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteSecurityGroupRule(request *DeleteSecurityGroupRuleRequest) (response *DeleteSecurityGroupRuleResponse, err error) {
    if request == nil {
        request = NewDeleteSecurityGroupRuleRequest()
    }
    response = NewDeleteSecurityGroupRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAcListsRequest() (request *DescribeAcListsRequest) {
    request = &DescribeAcListsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeAcLists")
    return
}

func NewDescribeAcListsResponse() (response *DescribeAcListsResponse) {
    response = &DescribeAcListsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAcLists
// 访问控制列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAcLists(request *DescribeAcListsRequest) (response *DescribeAcListsResponse, err error) {
    if request == nil {
        request = NewDescribeAcListsRequest()
    }
    response = NewDescribeAcListsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssociatedInstanceListRequest() (request *DescribeAssociatedInstanceListRequest) {
    request = &DescribeAssociatedInstanceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeAssociatedInstanceList")
    return
}

func NewDescribeAssociatedInstanceListResponse() (response *DescribeAssociatedInstanceListResponse) {
    response = &DescribeAssociatedInstanceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssociatedInstanceList
// 获取安全组关联实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAssociatedInstanceList(request *DescribeAssociatedInstanceListRequest) (response *DescribeAssociatedInstanceListResponse, err error) {
    if request == nil {
        request = NewDescribeAssociatedInstanceListRequest()
    }
    response = NewDescribeAssociatedInstanceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBlockByIpTimesListRequest() (request *DescribeBlockByIpTimesListRequest) {
    request = &DescribeBlockByIpTimesListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeBlockByIpTimesList")
    return
}

func NewDescribeBlockByIpTimesListResponse() (response *DescribeBlockByIpTimesListResponse) {
    response = &DescribeBlockByIpTimesListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBlockByIpTimesList
// DescribeBlockByIpTimesList 告警中心阻断IP折线图
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBlockByIpTimesList(request *DescribeBlockByIpTimesListRequest) (response *DescribeBlockByIpTimesListResponse, err error) {
    if request == nil {
        request = NewDescribeBlockByIpTimesListRequest()
    }
    response = NewDescribeBlockByIpTimesListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBlockStaticListRequest() (request *DescribeBlockStaticListRequest) {
    request = &DescribeBlockStaticListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeBlockStaticList")
    return
}

func NewDescribeBlockStaticListResponse() (response *DescribeBlockStaticListResponse) {
    response = &DescribeBlockStaticListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBlockStaticList
// DescribeBlockStaticList 告警中心柱形图
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBlockStaticList(request *DescribeBlockStaticListRequest) (response *DescribeBlockStaticListResponse, err error) {
    if request == nil {
        request = NewDescribeBlockStaticListRequest()
    }
    response = NewDescribeBlockStaticListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCfwEipsRequest() (request *DescribeCfwEipsRequest) {
    request = &DescribeCfwEipsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeCfwEips")
    return
}

func NewDescribeCfwEipsResponse() (response *DescribeCfwEipsResponse) {
    response = &DescribeCfwEipsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCfwEips
// 查询防火墙弹性公网IP
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCfwEips(request *DescribeCfwEipsRequest) (response *DescribeCfwEipsResponse, err error) {
    if request == nil {
        request = NewDescribeCfwEipsRequest()
    }
    response = NewDescribeCfwEipsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGuideScanInfoRequest() (request *DescribeGuideScanInfoRequest) {
    request = &DescribeGuideScanInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeGuideScanInfo")
    return
}

func NewDescribeGuideScanInfoResponse() (response *DescribeGuideScanInfoResponse) {
    response = &DescribeGuideScanInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGuideScanInfo
// DescribeGuideScanInfo新手引导扫描接口信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeGuideScanInfo(request *DescribeGuideScanInfoRequest) (response *DescribeGuideScanInfoResponse, err error) {
    if request == nil {
        request = NewDescribeGuideScanInfoRequest()
    }
    response = NewDescribeGuideScanInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNatRuleOverviewRequest() (request *DescribeNatRuleOverviewRequest) {
    request = &DescribeNatRuleOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeNatRuleOverview")
    return
}

func NewDescribeNatRuleOverviewResponse() (response *DescribeNatRuleOverviewResponse) {
    response = &DescribeNatRuleOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNatRuleOverview
// nat规则列表概况
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeNatRuleOverview(request *DescribeNatRuleOverviewRequest) (response *DescribeNatRuleOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeNatRuleOverviewRequest()
    }
    response = NewDescribeNatRuleOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRuleOverviewRequest() (request *DescribeRuleOverviewRequest) {
    request = &DescribeRuleOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeRuleOverview")
    return
}

func NewDescribeRuleOverviewResponse() (response *DescribeRuleOverviewResponse) {
    response = &DescribeRuleOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRuleOverview
// 查询规则列表概况
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeRuleOverview(request *DescribeRuleOverviewRequest) (response *DescribeRuleOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeRuleOverviewRequest()
    }
    response = NewDescribeRuleOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityGroupListRequest() (request *DescribeSecurityGroupListRequest) {
    request = &DescribeSecurityGroupListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeSecurityGroupList")
    return
}

func NewDescribeSecurityGroupListResponse() (response *DescribeSecurityGroupListResponse) {
    response = &DescribeSecurityGroupListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSecurityGroupList
// 查询安全组规则列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeSecurityGroupList(request *DescribeSecurityGroupListRequest) (response *DescribeSecurityGroupListResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityGroupListRequest()
    }
    response = NewDescribeSecurityGroupListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSwitchListsRequest() (request *DescribeSwitchListsRequest) {
    request = &DescribeSwitchListsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeSwitchLists")
    return
}

func NewDescribeSwitchListsResponse() (response *DescribeSwitchListsResponse) {
    response = &DescribeSwitchListsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSwitchLists
// 防火墙开关列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSwitchLists(request *DescribeSwitchListsRequest) (response *DescribeSwitchListsResponse, err error) {
    if request == nil {
        request = NewDescribeSwitchListsRequest()
    }
    response = NewDescribeSwitchListsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSyncAssetStatusRequest() (request *DescribeSyncAssetStatusRequest) {
    request = &DescribeSyncAssetStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeSyncAssetStatus")
    return
}

func NewDescribeSyncAssetStatusResponse() (response *DescribeSyncAssetStatusResponse) {
    response = &DescribeSyncAssetStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSyncAssetStatus
// 同步资产状态查询-互联网&VPC
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeSyncAssetStatus(request *DescribeSyncAssetStatusRequest) (response *DescribeSyncAssetStatusResponse, err error) {
    if request == nil {
        request = NewDescribeSyncAssetStatusRequest()
    }
    response = NewDescribeSyncAssetStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTLogInfoRequest() (request *DescribeTLogInfoRequest) {
    request = &DescribeTLogInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeTLogInfo")
    return
}

func NewDescribeTLogInfoResponse() (response *DescribeTLogInfoResponse) {
    response = &DescribeTLogInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTLogInfo
// DescribeTLogInfo告警中心概况
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTLogInfo(request *DescribeTLogInfoRequest) (response *DescribeTLogInfoResponse, err error) {
    if request == nil {
        request = NewDescribeTLogInfoRequest()
    }
    response = NewDescribeTLogInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTLogIpListRequest() (request *DescribeTLogIpListRequest) {
    request = &DescribeTLogIpListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeTLogIpList")
    return
}

func NewDescribeTLogIpListResponse() (response *DescribeTLogIpListResponse) {
    response = &DescribeTLogIpListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTLogIpList
// DescribeTLogIpList告警中心IP柱形图
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTLogIpList(request *DescribeTLogIpListRequest) (response *DescribeTLogIpListResponse, err error) {
    if request == nil {
        request = NewDescribeTLogIpListRequest()
    }
    response = NewDescribeTLogIpListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTableStatusRequest() (request *DescribeTableStatusRequest) {
    request = &DescribeTableStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeTableStatus")
    return
}

func NewDescribeTableStatusResponse() (response *DescribeTableStatusResponse) {
    response = &DescribeTableStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTableStatus
// 查询规则表状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTableStatus(request *DescribeTableStatusRequest) (response *DescribeTableStatusResponse, err error) {
    if request == nil {
        request = NewDescribeTableStatusRequest()
    }
    response = NewDescribeTableStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUnHandleEventTabListRequest() (request *DescribeUnHandleEventTabListRequest) {
    request = &DescribeUnHandleEventTabListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeUnHandleEventTabList")
    return
}

func NewDescribeUnHandleEventTabListResponse() (response *DescribeUnHandleEventTabListResponse) {
    response = &DescribeUnHandleEventTabListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUnHandleEventTabList
// DescribeUnHandleEventTabList 告警中心伪攻击链事件未处置接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUnHandleEventTabList(request *DescribeUnHandleEventTabListRequest) (response *DescribeUnHandleEventTabListResponse, err error) {
    if request == nil {
        request = NewDescribeUnHandleEventTabListRequest()
    }
    response = NewDescribeUnHandleEventTabListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVpcRuleOverviewRequest() (request *DescribeVpcRuleOverviewRequest) {
    request = &DescribeVpcRuleOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeVpcRuleOverview")
    return
}

func NewDescribeVpcRuleOverviewResponse() (response *DescribeVpcRuleOverviewResponse) {
    response = &DescribeVpcRuleOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVpcRuleOverview
// vpc规则列表概况
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeVpcRuleOverview(request *DescribeVpcRuleOverviewRequest) (response *DescribeVpcRuleOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeVpcRuleOverviewRequest()
    }
    response = NewDescribeVpcRuleOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewExpandCfwVerticalRequest() (request *ExpandCfwVerticalRequest) {
    request = &ExpandCfwVerticalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "ExpandCfwVertical")
    return
}

func NewExpandCfwVerticalResponse() (response *ExpandCfwVerticalResponse) {
    response = &ExpandCfwVerticalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExpandCfwVertical
// 防火墙垂直扩容
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ExpandCfwVertical(request *ExpandCfwVerticalRequest) (response *ExpandCfwVerticalResponse, err error) {
    if request == nil {
        request = NewExpandCfwVerticalRequest()
    }
    response = NewExpandCfwVerticalResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAcRuleRequest() (request *ModifyAcRuleRequest) {
    request = &ModifyAcRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyAcRule")
    return
}

func NewModifyAcRuleResponse() (response *ModifyAcRuleResponse) {
    response = &ModifyAcRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAcRule
// 修改规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyAcRule(request *ModifyAcRuleRequest) (response *ModifyAcRuleResponse, err error) {
    if request == nil {
        request = NewModifyAcRuleRequest()
    }
    response = NewModifyAcRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAllRuleStatusRequest() (request *ModifyAllRuleStatusRequest) {
    request = &ModifyAllRuleStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyAllRuleStatus")
    return
}

func NewModifyAllRuleStatusResponse() (response *ModifyAllRuleStatusResponse) {
    response = &ModifyAllRuleStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAllRuleStatus
// 启用停用全部规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyAllRuleStatus(request *ModifyAllRuleStatusRequest) (response *ModifyAllRuleStatusResponse, err error) {
    if request == nil {
        request = NewModifyAllRuleStatusRequest()
    }
    response = NewModifyAllRuleStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAllSwitchStatusRequest() (request *ModifyAllSwitchStatusRequest) {
    request = &ModifyAllSwitchStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyAllSwitchStatus")
    return
}

func NewModifyAllSwitchStatusResponse() (response *ModifyAllSwitchStatusResponse) {
    response = &ModifyAllSwitchStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAllSwitchStatus
// 一键开启和关闭
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) ModifyAllSwitchStatus(request *ModifyAllSwitchStatusRequest) (response *ModifyAllSwitchStatusResponse, err error) {
    if request == nil {
        request = NewModifyAllSwitchStatusRequest()
    }
    response = NewModifyAllSwitchStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBlockIgnoreListRequest() (request *ModifyBlockIgnoreListRequest) {
    request = &ModifyBlockIgnoreListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyBlockIgnoreList")
    return
}

func NewModifyBlockIgnoreListResponse() (response *ModifyBlockIgnoreListResponse) {
    response = &ModifyBlockIgnoreListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyBlockIgnoreList
// 支持对拦截列表、忽略列表如下操作：
//
// 批量增加拦截IP、忽略IP/域名
//
// 批量删除拦截IP、忽略IP/域名
//
// 批量修改拦截IP、忽略IP/域名生效事件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyBlockIgnoreList(request *ModifyBlockIgnoreListRequest) (response *ModifyBlockIgnoreListResponse, err error) {
    if request == nil {
        request = NewModifyBlockIgnoreListRequest()
    }
    response = NewModifyBlockIgnoreListResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBlockTopRequest() (request *ModifyBlockTopRequest) {
    request = &ModifyBlockTopRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyBlockTop")
    return
}

func NewModifyBlockTopResponse() (response *ModifyBlockTopResponse) {
    response = &ModifyBlockTopResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyBlockTop
// ModifyBlockTop取消置顶接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyBlockTop(request *ModifyBlockTopRequest) (response *ModifyBlockTopResponse, err error) {
    if request == nil {
        request = NewModifyBlockTopRequest()
    }
    response = NewModifyBlockTopResponse()
    err = c.Send(request, response)
    return
}

func NewModifyItemSwitchStatusRequest() (request *ModifyItemSwitchStatusRequest) {
    request = &ModifyItemSwitchStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyItemSwitchStatus")
    return
}

func NewModifyItemSwitchStatusResponse() (response *ModifyItemSwitchStatusResponse) {
    response = &ModifyItemSwitchStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyItemSwitchStatus
// 修改单个防火墙开关
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyItemSwitchStatus(request *ModifyItemSwitchStatusRequest) (response *ModifyItemSwitchStatusResponse, err error) {
    if request == nil {
        request = NewModifyItemSwitchStatusRequest()
    }
    response = NewModifyItemSwitchStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifySecurityGroupAllRuleStatusRequest() (request *ModifySecurityGroupAllRuleStatusRequest) {
    request = &ModifySecurityGroupAllRuleStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "ModifySecurityGroupAllRuleStatus")
    return
}

func NewModifySecurityGroupAllRuleStatusResponse() (response *ModifySecurityGroupAllRuleStatusResponse) {
    response = &ModifySecurityGroupAllRuleStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifySecurityGroupAllRuleStatus
// 启用停用全部规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifySecurityGroupAllRuleStatus(request *ModifySecurityGroupAllRuleStatusRequest) (response *ModifySecurityGroupAllRuleStatusResponse, err error) {
    if request == nil {
        request = NewModifySecurityGroupAllRuleStatusRequest()
    }
    response = NewModifySecurityGroupAllRuleStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifySequenceRulesRequest() (request *ModifySequenceRulesRequest) {
    request = &ModifySequenceRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "ModifySequenceRules")
    return
}

func NewModifySequenceRulesResponse() (response *ModifySequenceRulesResponse) {
    response = &ModifySequenceRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifySequenceRules
// 修改规则执行顺序
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifySequenceRules(request *ModifySequenceRulesRequest) (response *ModifySequenceRulesResponse, err error) {
    if request == nil {
        request = NewModifySequenceRulesRequest()
    }
    response = NewModifySequenceRulesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTableStatusRequest() (request *ModifyTableStatusRequest) {
    request = &ModifyTableStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyTableStatus")
    return
}

func NewModifyTableStatusResponse() (response *ModifyTableStatusResponse) {
    response = &ModifyTableStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyTableStatus
// 修改规则表状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyTableStatus(request *ModifyTableStatusRequest) (response *ModifyTableStatusResponse, err error) {
    if request == nil {
        request = NewModifyTableStatusRequest()
    }
    response = NewModifyTableStatusResponse()
    err = c.Send(request, response)
    return
}

func NewRunSyncAssetRequest() (request *RunSyncAssetRequest) {
    request = &RunSyncAssetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "RunSyncAsset")
    return
}

func NewRunSyncAssetResponse() (response *RunSyncAssetResponse) {
    response = &RunSyncAssetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RunSyncAsset
// 同步资产-互联网&VPC
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) RunSyncAsset(request *RunSyncAssetRequest) (response *RunSyncAssetResponse, err error) {
    if request == nil {
        request = NewRunSyncAssetRequest()
    }
    response = NewRunSyncAssetResponse()
    err = c.Send(request, response)
    return
}

func NewSetNatFwDnatRuleRequest() (request *SetNatFwDnatRuleRequest) {
    request = &SetNatFwDnatRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "SetNatFwDnatRule")
    return
}

func NewSetNatFwDnatRuleResponse() (response *SetNatFwDnatRuleResponse) {
    response = &SetNatFwDnatRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetNatFwDnatRule
// 配置防火墙Dnat规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SetNatFwDnatRule(request *SetNatFwDnatRuleRequest) (response *SetNatFwDnatRuleResponse, err error) {
    if request == nil {
        request = NewSetNatFwDnatRuleRequest()
    }
    response = NewSetNatFwDnatRuleResponse()
    err = c.Send(request, response)
    return
}
