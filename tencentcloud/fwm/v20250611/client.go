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

package v20250611

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2025-06-11"

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


func NewCancelIgnorePolicyRiskRequest() (request *CancelIgnorePolicyRiskRequest) {
    request = &CancelIgnorePolicyRiskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fwm", APIVersion, "CancelIgnorePolicyRisk")
    
    
    return
}

func NewCancelIgnorePolicyRiskResponse() (response *CancelIgnorePolicyRiskResponse) {
    response = &CancelIgnorePolicyRiskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CancelIgnorePolicyRisk
// 取消忽略策略风险
//
// 可能返回的错误码:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CancelIgnorePolicyRisk(request *CancelIgnorePolicyRiskRequest) (response *CancelIgnorePolicyRiskResponse, err error) {
    return c.CancelIgnorePolicyRiskWithContext(context.Background(), request)
}

// CancelIgnorePolicyRisk
// 取消忽略策略风险
//
// 可能返回的错误码:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CancelIgnorePolicyRiskWithContext(ctx context.Context, request *CancelIgnorePolicyRiskRequest) (response *CancelIgnorePolicyRiskResponse, err error) {
    if request == nil {
        request = NewCancelIgnorePolicyRiskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "fwm", APIVersion, "CancelIgnorePolicyRisk")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelIgnorePolicyRisk require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelIgnorePolicyRiskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAnalyzePolicyTaskRequest() (request *CreateAnalyzePolicyTaskRequest) {
    request = &CreateAnalyzePolicyTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fwm", APIVersion, "CreateAnalyzePolicyTask")
    
    
    return
}

func NewCreateAnalyzePolicyTaskResponse() (response *CreateAnalyzePolicyTaskResponse) {
    response = &CreateAnalyzePolicyTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAnalyzePolicyTask
// 创建策略风险分析任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAnalyzePolicyTask(request *CreateAnalyzePolicyTaskRequest) (response *CreateAnalyzePolicyTaskResponse, err error) {
    return c.CreateAnalyzePolicyTaskWithContext(context.Background(), request)
}

// CreateAnalyzePolicyTask
// 创建策略风险分析任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAnalyzePolicyTaskWithContext(ctx context.Context, request *CreateAnalyzePolicyTaskRequest) (response *CreateAnalyzePolicyTaskResponse, err error) {
    if request == nil {
        request = NewCreateAnalyzePolicyTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "fwm", APIVersion, "CreateAnalyzePolicyTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAnalyzePolicyTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAnalyzePolicyTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEdgeAclRuleRequest() (request *CreateEdgeAclRuleRequest) {
    request = &CreateEdgeAclRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fwm", APIVersion, "CreateEdgeAclRule")
    
    
    return
}

func NewCreateEdgeAclRuleResponse() (response *CreateEdgeAclRuleResponse) {
    response = &CreateEdgeAclRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateEdgeAclRule
// 向已有的互联网边界ACL规则组中添加规则。需要先创建规则组，然后通过此接口添加规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateEdgeAclRule(request *CreateEdgeAclRuleRequest) (response *CreateEdgeAclRuleResponse, err error) {
    return c.CreateEdgeAclRuleWithContext(context.Background(), request)
}

// CreateEdgeAclRule
// 向已有的互联网边界ACL规则组中添加规则。需要先创建规则组，然后通过此接口添加规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateEdgeAclRuleWithContext(ctx context.Context, request *CreateEdgeAclRuleRequest) (response *CreateEdgeAclRuleResponse, err error) {
    if request == nil {
        request = NewCreateEdgeAclRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "fwm", APIVersion, "CreateEdgeAclRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEdgeAclRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEdgeAclRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEdgeAclRuleGroupRequest() (request *CreateEdgeAclRuleGroupRequest) {
    request = &CreateEdgeAclRuleGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fwm", APIVersion, "CreateEdgeAclRuleGroup")
    
    
    return
}

func NewCreateEdgeAclRuleGroupResponse() (response *CreateEdgeAclRuleGroupResponse) {
    response = &CreateEdgeAclRuleGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateEdgeAclRuleGroup
// 创建互联网边界ACL规则组，支持同时创建多条规则。Product 必须为 cfw_edge_acl。规则支持 IP、域名、参数模板、实例、标签等多种源/目标类型。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateEdgeAclRuleGroup(request *CreateEdgeAclRuleGroupRequest) (response *CreateEdgeAclRuleGroupResponse, err error) {
    return c.CreateEdgeAclRuleGroupWithContext(context.Background(), request)
}

// CreateEdgeAclRuleGroup
// 创建互联网边界ACL规则组，支持同时创建多条规则。Product 必须为 cfw_edge_acl。规则支持 IP、域名、参数模板、实例、标签等多种源/目标类型。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateEdgeAclRuleGroupWithContext(ctx context.Context, request *CreateEdgeAclRuleGroupRequest) (response *CreateEdgeAclRuleGroupResponse, err error) {
    if request == nil {
        request = NewCreateEdgeAclRuleGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "fwm", APIVersion, "CreateEdgeAclRuleGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEdgeAclRuleGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEdgeAclRuleGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNatAclRuleRequest() (request *CreateNatAclRuleRequest) {
    request = &CreateNatAclRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fwm", APIVersion, "CreateNatAclRule")
    
    
    return
}

func NewCreateNatAclRuleResponse() (response *CreateNatAclRuleResponse) {
    response = &CreateNatAclRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateNatAclRule
// 在已有规则组中添加NAT ACL规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateNatAclRule(request *CreateNatAclRuleRequest) (response *CreateNatAclRuleResponse, err error) {
    return c.CreateNatAclRuleWithContext(context.Background(), request)
}

// CreateNatAclRule
// 在已有规则组中添加NAT ACL规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateNatAclRuleWithContext(ctx context.Context, request *CreateNatAclRuleRequest) (response *CreateNatAclRuleResponse, err error) {
    if request == nil {
        request = NewCreateNatAclRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "fwm", APIVersion, "CreateNatAclRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNatAclRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateNatAclRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNatAclRuleGroupRequest() (request *CreateNatAclRuleGroupRequest) {
    request = &CreateNatAclRuleGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fwm", APIVersion, "CreateNatAclRuleGroup")
    
    
    return
}

func NewCreateNatAclRuleGroupResponse() (response *CreateNatAclRuleGroupResponse) {
    response = &CreateNatAclRuleGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateNatAclRuleGroup
// 创建NAT ACL规则组（NAT边界防火墙规则组管理）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateNatAclRuleGroup(request *CreateNatAclRuleGroupRequest) (response *CreateNatAclRuleGroupResponse, err error) {
    return c.CreateNatAclRuleGroupWithContext(context.Background(), request)
}

// CreateNatAclRuleGroup
// 创建NAT ACL规则组（NAT边界防火墙规则组管理）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateNatAclRuleGroupWithContext(ctx context.Context, request *CreateNatAclRuleGroupRequest) (response *CreateNatAclRuleGroupResponse, err error) {
    if request == nil {
        request = NewCreateNatAclRuleGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "fwm", APIVersion, "CreateNatAclRuleGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNatAclRuleGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateNatAclRuleGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSecurityGroupRuleRequest() (request *CreateSecurityGroupRuleRequest) {
    request = &CreateSecurityGroupRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fwm", APIVersion, "CreateSecurityGroupRule")
    
    
    return
}

func NewCreateSecurityGroupRuleResponse() (response *CreateSecurityGroupRuleResponse) {
    response = &CreateSecurityGroupRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSecurityGroupRule
// 规则组编辑时添加规则（规则组管理）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateSecurityGroupRule(request *CreateSecurityGroupRuleRequest) (response *CreateSecurityGroupRuleResponse, err error) {
    return c.CreateSecurityGroupRuleWithContext(context.Background(), request)
}

// CreateSecurityGroupRule
// 规则组编辑时添加规则（规则组管理）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateSecurityGroupRuleWithContext(ctx context.Context, request *CreateSecurityGroupRuleRequest) (response *CreateSecurityGroupRuleResponse, err error) {
    if request == nil {
        request = NewCreateSecurityGroupRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "fwm", APIVersion, "CreateSecurityGroupRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSecurityGroupRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSecurityGroupRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSecurityGroupRuleGroupRequest() (request *CreateSecurityGroupRuleGroupRequest) {
    request = &CreateSecurityGroupRuleGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fwm", APIVersion, "CreateSecurityGroupRuleGroup")
    
    
    return
}

func NewCreateSecurityGroupRuleGroupResponse() (response *CreateSecurityGroupRuleGroupResponse) {
    response = &CreateSecurityGroupRuleGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSecurityGroupRuleGroup
// 创建规则组（规则组管理）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateSecurityGroupRuleGroup(request *CreateSecurityGroupRuleGroupRequest) (response *CreateSecurityGroupRuleGroupResponse, err error) {
    return c.CreateSecurityGroupRuleGroupWithContext(context.Background(), request)
}

// CreateSecurityGroupRuleGroup
// 创建规则组（规则组管理）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateSecurityGroupRuleGroupWithContext(ctx context.Context, request *CreateSecurityGroupRuleGroupRequest) (response *CreateSecurityGroupRuleGroupResponse, err error) {
    if request == nil {
        request = NewCreateSecurityGroupRuleGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "fwm", APIVersion, "CreateSecurityGroupRuleGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSecurityGroupRuleGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSecurityGroupRuleGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateStrategyRequest() (request *CreateStrategyRequest) {
    request = &CreateStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fwm", APIVersion, "CreateStrategy")
    
    
    return
}

func NewCreateStrategyResponse() (response *CreateStrategyResponse) {
    response = &CreateStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateStrategy
// 创建策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateStrategy(request *CreateStrategyRequest) (response *CreateStrategyResponse, err error) {
    return c.CreateStrategyWithContext(context.Background(), request)
}

// CreateStrategy
// 创建策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateStrategyWithContext(ctx context.Context, request *CreateStrategyRequest) (response *CreateStrategyResponse, err error) {
    if request == nil {
        request = NewCreateStrategyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "fwm", APIVersion, "CreateStrategy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVpcAclRuleRequest() (request *CreateVpcAclRuleRequest) {
    request = &CreateVpcAclRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fwm", APIVersion, "CreateVpcAclRule")
    
    
    return
}

func NewCreateVpcAclRuleResponse() (response *CreateVpcAclRuleResponse) {
    response = &CreateVpcAclRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateVpcAclRule
// 在已有规则组中添加VPC ACL规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateVpcAclRule(request *CreateVpcAclRuleRequest) (response *CreateVpcAclRuleResponse, err error) {
    return c.CreateVpcAclRuleWithContext(context.Background(), request)
}

// CreateVpcAclRule
// 在已有规则组中添加VPC ACL规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateVpcAclRuleWithContext(ctx context.Context, request *CreateVpcAclRuleRequest) (response *CreateVpcAclRuleResponse, err error) {
    if request == nil {
        request = NewCreateVpcAclRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "fwm", APIVersion, "CreateVpcAclRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVpcAclRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateVpcAclRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVpcAclRuleGroupRequest() (request *CreateVpcAclRuleGroupRequest) {
    request = &CreateVpcAclRuleGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fwm", APIVersion, "CreateVpcAclRuleGroup")
    
    
    return
}

func NewCreateVpcAclRuleGroupResponse() (response *CreateVpcAclRuleGroupResponse) {
    response = &CreateVpcAclRuleGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateVpcAclRuleGroup
// 创建VPC ACL规则组（VPC间防火墙规则组管理）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateVpcAclRuleGroup(request *CreateVpcAclRuleGroupRequest) (response *CreateVpcAclRuleGroupResponse, err error) {
    return c.CreateVpcAclRuleGroupWithContext(context.Background(), request)
}

// CreateVpcAclRuleGroup
// 创建VPC ACL规则组（VPC间防火墙规则组管理）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateVpcAclRuleGroupWithContext(ctx context.Context, request *CreateVpcAclRuleGroupRequest) (response *CreateVpcAclRuleGroupResponse, err error) {
    if request == nil {
        request = NewCreateVpcAclRuleGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "fwm", APIVersion, "CreateVpcAclRuleGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVpcAclRuleGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateVpcAclRuleGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEdgeAclRuleRequest() (request *DeleteEdgeAclRuleRequest) {
    request = &DeleteEdgeAclRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fwm", APIVersion, "DeleteEdgeAclRule")
    
    
    return
}

func NewDeleteEdgeAclRuleResponse() (response *DeleteEdgeAclRuleResponse) {
    response = &DeleteEdgeAclRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteEdgeAclRule
// 批量删除互联网边界ACL规则。支持一次删除多条规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteEdgeAclRule(request *DeleteEdgeAclRuleRequest) (response *DeleteEdgeAclRuleResponse, err error) {
    return c.DeleteEdgeAclRuleWithContext(context.Background(), request)
}

// DeleteEdgeAclRule
// 批量删除互联网边界ACL规则。支持一次删除多条规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteEdgeAclRuleWithContext(ctx context.Context, request *DeleteEdgeAclRuleRequest) (response *DeleteEdgeAclRuleResponse, err error) {
    if request == nil {
        request = NewDeleteEdgeAclRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "fwm", APIVersion, "DeleteEdgeAclRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteEdgeAclRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteEdgeAclRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNatAclRuleRequest() (request *DeleteNatAclRuleRequest) {
    request = &DeleteNatAclRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fwm", APIVersion, "DeleteNatAclRule")
    
    
    return
}

func NewDeleteNatAclRuleResponse() (response *DeleteNatAclRuleResponse) {
    response = &DeleteNatAclRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteNatAclRule
// 删除NAT ACL规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteNatAclRule(request *DeleteNatAclRuleRequest) (response *DeleteNatAclRuleResponse, err error) {
    return c.DeleteNatAclRuleWithContext(context.Background(), request)
}

// DeleteNatAclRule
// 删除NAT ACL规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteNatAclRuleWithContext(ctx context.Context, request *DeleteNatAclRuleRequest) (response *DeleteNatAclRuleResponse, err error) {
    if request == nil {
        request = NewDeleteNatAclRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "fwm", APIVersion, "DeleteNatAclRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteNatAclRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteNatAclRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRuleGroupRequest() (request *DeleteRuleGroupRequest) {
    request = &DeleteRuleGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fwm", APIVersion, "DeleteRuleGroup")
    
    
    return
}

func NewDeleteRuleGroupResponse() (response *DeleteRuleGroupResponse) {
    response = &DeleteRuleGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRuleGroup
// 删除规则组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DeleteRuleGroup(request *DeleteRuleGroupRequest) (response *DeleteRuleGroupResponse, err error) {
    return c.DeleteRuleGroupWithContext(context.Background(), request)
}

// DeleteRuleGroup
// 删除规则组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DeleteRuleGroupWithContext(ctx context.Context, request *DeleteRuleGroupRequest) (response *DeleteRuleGroupResponse, err error) {
    if request == nil {
        request = NewDeleteRuleGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "fwm", APIVersion, "DeleteRuleGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRuleGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRuleGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSecurityGroupRuleRequest() (request *DeleteSecurityGroupRuleRequest) {
    request = &DeleteSecurityGroupRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fwm", APIVersion, "DeleteSecurityGroupRule")
    
    
    return
}

func NewDeleteSecurityGroupRuleResponse() (response *DeleteSecurityGroupRuleResponse) {
    response = &DeleteSecurityGroupRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSecurityGroupRule
// 删除规则（规则组管理）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteSecurityGroupRule(request *DeleteSecurityGroupRuleRequest) (response *DeleteSecurityGroupRuleResponse, err error) {
    return c.DeleteSecurityGroupRuleWithContext(context.Background(), request)
}

// DeleteSecurityGroupRule
// 删除规则（规则组管理）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteSecurityGroupRuleWithContext(ctx context.Context, request *DeleteSecurityGroupRuleRequest) (response *DeleteSecurityGroupRuleResponse, err error) {
    if request == nil {
        request = NewDeleteSecurityGroupRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "fwm", APIVersion, "DeleteSecurityGroupRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSecurityGroupRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSecurityGroupRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteStrategyRequest() (request *DeleteStrategyRequest) {
    request = &DeleteStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fwm", APIVersion, "DeleteStrategy")
    
    
    return
}

func NewDeleteStrategyResponse() (response *DeleteStrategyResponse) {
    response = &DeleteStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteStrategy
// 删除策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteStrategy(request *DeleteStrategyRequest) (response *DeleteStrategyResponse, err error) {
    return c.DeleteStrategyWithContext(context.Background(), request)
}

// DeleteStrategy
// 删除策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteStrategyWithContext(ctx context.Context, request *DeleteStrategyRequest) (response *DeleteStrategyResponse, err error) {
    if request == nil {
        request = NewDeleteStrategyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "fwm", APIVersion, "DeleteStrategy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteVpcAclRuleRequest() (request *DeleteVpcAclRuleRequest) {
    request = &DeleteVpcAclRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fwm", APIVersion, "DeleteVpcAclRule")
    
    
    return
}

func NewDeleteVpcAclRuleResponse() (response *DeleteVpcAclRuleResponse) {
    response = &DeleteVpcAclRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteVpcAclRule
// 删除VPC ACL规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteVpcAclRule(request *DeleteVpcAclRuleRequest) (response *DeleteVpcAclRuleResponse, err error) {
    return c.DeleteVpcAclRuleWithContext(context.Background(), request)
}

// DeleteVpcAclRule
// 删除VPC ACL规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteVpcAclRuleWithContext(ctx context.Context, request *DeleteVpcAclRuleRequest) (response *DeleteVpcAclRuleResponse, err error) {
    if request == nil {
        request = NewDeleteVpcAclRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "fwm", APIVersion, "DeleteVpcAclRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteVpcAclRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteVpcAclRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgeAclRulesRequest() (request *DescribeEdgeAclRulesRequest) {
    request = &DescribeEdgeAclRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fwm", APIVersion, "DescribeEdgeAclRules")
    
    
    return
}

func NewDescribeEdgeAclRulesResponse() (response *DescribeEdgeAclRulesResponse) {
    response = &DescribeEdgeAclRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEdgeAclRules
// 查询指定规则组下的互联网边界ACL规则列表。支持分页和多种过滤条件。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeAclRules(request *DescribeEdgeAclRulesRequest) (response *DescribeEdgeAclRulesResponse, err error) {
    return c.DescribeEdgeAclRulesWithContext(context.Background(), request)
}

// DescribeEdgeAclRules
// 查询指定规则组下的互联网边界ACL规则列表。支持分页和多种过滤条件。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeAclRulesWithContext(ctx context.Context, request *DescribeEdgeAclRulesRequest) (response *DescribeEdgeAclRulesResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeAclRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "fwm", APIVersion, "DescribeEdgeAclRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgeAclRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgeAclRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNatAclRulesRequest() (request *DescribeNatAclRulesRequest) {
    request = &DescribeNatAclRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fwm", APIVersion, "DescribeNatAclRules")
    
    
    return
}

func NewDescribeNatAclRulesResponse() (response *DescribeNatAclRulesResponse) {
    response = &DescribeNatAclRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNatAclRules
// 查询NAT ACL规则列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeNatAclRules(request *DescribeNatAclRulesRequest) (response *DescribeNatAclRulesResponse, err error) {
    return c.DescribeNatAclRulesWithContext(context.Background(), request)
}

// DescribeNatAclRules
// 查询NAT ACL规则列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeNatAclRulesWithContext(ctx context.Context, request *DescribeNatAclRulesRequest) (response *DescribeNatAclRulesResponse, err error) {
    if request == nil {
        request = NewDescribeNatAclRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "fwm", APIVersion, "DescribeNatAclRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNatAclRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNatAclRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrganMembersRequest() (request *DescribeOrganMembersRequest) {
    request = &DescribeOrganMembersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fwm", APIVersion, "DescribeOrganMembers")
    
    
    return
}

func NewDescribeOrganMembersResponse() (response *DescribeOrganMembersResponse) {
    response = &DescribeOrganMembersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOrganMembers
// 查询集团下所有纳管成员账号列表，支持分页、排序和多条件筛选，仅管理员可调用
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeOrganMembers(request *DescribeOrganMembersRequest) (response *DescribeOrganMembersResponse, err error) {
    return c.DescribeOrganMembersWithContext(context.Background(), request)
}

// DescribeOrganMembers
// 查询集团下所有纳管成员账号列表，支持分页、排序和多条件筛选，仅管理员可调用
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeOrganMembersWithContext(ctx context.Context, request *DescribeOrganMembersRequest) (response *DescribeOrganMembersResponse, err error) {
    if request == nil {
        request = NewDescribeOrganMembersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "fwm", APIVersion, "DescribeOrganMembers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrganMembers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrganMembersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrganSummaryRequest() (request *DescribeOrganSummaryRequest) {
    request = &DescribeOrganSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fwm", APIVersion, "DescribeOrganSummary")
    
    
    return
}

func NewDescribeOrganSummaryResponse() (response *DescribeOrganSummaryResponse) {
    response = &DescribeOrganSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOrganSummary
// 获取集团概览信息，包括集团名称、管理员信息、成员数量等
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeOrganSummary(request *DescribeOrganSummaryRequest) (response *DescribeOrganSummaryResponse, err error) {
    return c.DescribeOrganSummaryWithContext(context.Background(), request)
}

// DescribeOrganSummary
// 获取集团概览信息，包括集团名称、管理员信息、成员数量等
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeOrganSummaryWithContext(ctx context.Context, request *DescribeOrganSummaryRequest) (response *DescribeOrganSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeOrganSummaryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "fwm", APIVersion, "DescribeOrganSummary")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrganSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrganSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePolicyRiskAccountProductStatsRequest() (request *DescribePolicyRiskAccountProductStatsRequest) {
    request = &DescribePolicyRiskAccountProductStatsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fwm", APIVersion, "DescribePolicyRiskAccountProductStats")
    
    
    return
}

func NewDescribePolicyRiskAccountProductStatsResponse() (response *DescribePolicyRiskAccountProductStatsResponse) {
    response = &DescribePolicyRiskAccountProductStatsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePolicyRiskAccountProductStats
// 查询账号+产品维度风险统计，按账号分组返回各产品的体检策略数、待整改风险数、整改率、最近体检时间等信息，支持按账号名称/ID搜索以及仅看待整改、仅超时未体检筛选
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribePolicyRiskAccountProductStats(request *DescribePolicyRiskAccountProductStatsRequest) (response *DescribePolicyRiskAccountProductStatsResponse, err error) {
    return c.DescribePolicyRiskAccountProductStatsWithContext(context.Background(), request)
}

// DescribePolicyRiskAccountProductStats
// 查询账号+产品维度风险统计，按账号分组返回各产品的体检策略数、待整改风险数、整改率、最近体检时间等信息，支持按账号名称/ID搜索以及仅看待整改、仅超时未体检筛选
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribePolicyRiskAccountProductStatsWithContext(ctx context.Context, request *DescribePolicyRiskAccountProductStatsRequest) (response *DescribePolicyRiskAccountProductStatsResponse, err error) {
    if request == nil {
        request = NewDescribePolicyRiskAccountProductStatsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "fwm", APIVersion, "DescribePolicyRiskAccountProductStats")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePolicyRiskAccountProductStats require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePolicyRiskAccountProductStatsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskAnalysisDetailsRequest() (request *DescribeRiskAnalysisDetailsRequest) {
    request = &DescribeRiskAnalysisDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fwm", APIVersion, "DescribeRiskAnalysisDetails")
    
    
    return
}

func NewDescribeRiskAnalysisDetailsResponse() (response *DescribeRiskAnalysisDetailsResponse) {
    response = &DescribeRiskAnalysisDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRiskAnalysisDetails
// 获取实时分析风险详情
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeRiskAnalysisDetails(request *DescribeRiskAnalysisDetailsRequest) (response *DescribeRiskAnalysisDetailsResponse, err error) {
    return c.DescribeRiskAnalysisDetailsWithContext(context.Background(), request)
}

// DescribeRiskAnalysisDetails
// 获取实时分析风险详情
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeRiskAnalysisDetailsWithContext(ctx context.Context, request *DescribeRiskAnalysisDetailsRequest) (response *DescribeRiskAnalysisDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeRiskAnalysisDetailsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "fwm", APIVersion, "DescribeRiskAnalysisDetails")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskAnalysisDetails require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskAnalysisDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskCategoryStatsRequest() (request *DescribeRiskCategoryStatsRequest) {
    request = &DescribeRiskCategoryStatsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fwm", APIVersion, "DescribeRiskCategoryStats")
    
    
    return
}

func NewDescribeRiskCategoryStatsResponse() (response *DescribeRiskCategoryStatsResponse) {
    response = &DescribeRiskCategoryStatsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRiskCategoryStats
// 查询策略体检风险分类统计数据,包含各类风险的规则数量、处置状态、整改率等信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeRiskCategoryStats(request *DescribeRiskCategoryStatsRequest) (response *DescribeRiskCategoryStatsResponse, err error) {
    return c.DescribeRiskCategoryStatsWithContext(context.Background(), request)
}

// DescribeRiskCategoryStats
// 查询策略体检风险分类统计数据,包含各类风险的规则数量、处置状态、整改率等信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeRiskCategoryStatsWithContext(ctx context.Context, request *DescribeRiskCategoryStatsRequest) (response *DescribeRiskCategoryStatsResponse, err error) {
    if request == nil {
        request = NewDescribeRiskCategoryStatsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "fwm", APIVersion, "DescribeRiskCategoryStats")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskCategoryStats require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskCategoryStatsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskListRequest() (request *DescribeRiskListRequest) {
    request = &DescribeRiskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fwm", APIVersion, "DescribeRiskList")
    
    
    return
}

func NewDescribeRiskListResponse() (response *DescribeRiskListResponse) {
    response = &DescribeRiskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRiskList
// 查询用户所有规则的策略问题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRiskList(request *DescribeRiskListRequest) (response *DescribeRiskListResponse, err error) {
    return c.DescribeRiskListWithContext(context.Background(), request)
}

// DescribeRiskList
// 查询用户所有规则的策略问题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRiskListWithContext(ctx context.Context, request *DescribeRiskListRequest) (response *DescribeRiskListResponse, err error) {
    if request == nil {
        request = NewDescribeRiskListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "fwm", APIVersion, "DescribeRiskList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityGroupRuleRequest() (request *DescribeSecurityGroupRuleRequest) {
    request = &DescribeSecurityGroupRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fwm", APIVersion, "DescribeSecurityGroupRule")
    
    
    return
}

func NewDescribeSecurityGroupRuleResponse() (response *DescribeSecurityGroupRuleResponse) {
    response = &DescribeSecurityGroupRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSecurityGroupRule
// 查询规则详情（规则组管理）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
func (c *Client) DescribeSecurityGroupRule(request *DescribeSecurityGroupRuleRequest) (response *DescribeSecurityGroupRuleResponse, err error) {
    return c.DescribeSecurityGroupRuleWithContext(context.Background(), request)
}

// DescribeSecurityGroupRule
// 查询规则详情（规则组管理）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
func (c *Client) DescribeSecurityGroupRuleWithContext(ctx context.Context, request *DescribeSecurityGroupRuleRequest) (response *DescribeSecurityGroupRuleResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityGroupRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "fwm", APIVersion, "DescribeSecurityGroupRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityGroupRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityGroupRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityGroupRulesRequest() (request *DescribeSecurityGroupRulesRequest) {
    request = &DescribeSecurityGroupRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fwm", APIVersion, "DescribeSecurityGroupRules")
    
    
    return
}

func NewDescribeSecurityGroupRulesResponse() (response *DescribeSecurityGroupRulesResponse) {
    response = &DescribeSecurityGroupRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSecurityGroupRules
// 查询规则组中规则列表接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeSecurityGroupRules(request *DescribeSecurityGroupRulesRequest) (response *DescribeSecurityGroupRulesResponse, err error) {
    return c.DescribeSecurityGroupRulesWithContext(context.Background(), request)
}

// DescribeSecurityGroupRules
// 查询规则组中规则列表接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeSecurityGroupRulesWithContext(ctx context.Context, request *DescribeSecurityGroupRulesRequest) (response *DescribeSecurityGroupRulesResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityGroupRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "fwm", APIVersion, "DescribeSecurityGroupRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityGroupRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityGroupRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStrategiesRequest() (request *DescribeStrategiesRequest) {
    request = &DescribeStrategiesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fwm", APIVersion, "DescribeStrategies")
    
    
    return
}

func NewDescribeStrategiesResponse() (response *DescribeStrategiesResponse) {
    response = &DescribeStrategiesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStrategies
// 查询策略列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_COMMONPARAM = "MissingParameter.CommonParam"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeStrategies(request *DescribeStrategiesRequest) (response *DescribeStrategiesResponse, err error) {
    return c.DescribeStrategiesWithContext(context.Background(), request)
}

// DescribeStrategies
// 查询策略列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_COMMONPARAM = "MissingParameter.CommonParam"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeStrategiesWithContext(ctx context.Context, request *DescribeStrategiesRequest) (response *DescribeStrategiesResponse, err error) {
    if request == nil {
        request = NewDescribeStrategiesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "fwm", APIVersion, "DescribeStrategies")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStrategies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStrategiesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStrategyRequest() (request *DescribeStrategyRequest) {
    request = &DescribeStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fwm", APIVersion, "DescribeStrategy")
    
    
    return
}

func NewDescribeStrategyResponse() (response *DescribeStrategyResponse) {
    response = &DescribeStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStrategy
// 查询策略详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
func (c *Client) DescribeStrategy(request *DescribeStrategyRequest) (response *DescribeStrategyResponse, err error) {
    return c.DescribeStrategyWithContext(context.Background(), request)
}

// DescribeStrategy
// 查询策略详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
func (c *Client) DescribeStrategyWithContext(ctx context.Context, request *DescribeStrategyRequest) (response *DescribeStrategyResponse, err error) {
    if request == nil {
        request = NewDescribeStrategyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "fwm", APIVersion, "DescribeStrategy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStrategyAccountsRequest() (request *DescribeStrategyAccountsRequest) {
    request = &DescribeStrategyAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fwm", APIVersion, "DescribeStrategyAccounts")
    
    
    return
}

func NewDescribeStrategyAccountsResponse() (response *DescribeStrategyAccountsResponse) {
    response = &DescribeStrategyAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStrategyAccounts
// 查看防火墙管理规则下发账号列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeStrategyAccounts(request *DescribeStrategyAccountsRequest) (response *DescribeStrategyAccountsResponse, err error) {
    return c.DescribeStrategyAccountsWithContext(context.Background(), request)
}

// DescribeStrategyAccounts
// 查看防火墙管理规则下发账号列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeStrategyAccountsWithContext(ctx context.Context, request *DescribeStrategyAccountsRequest) (response *DescribeStrategyAccountsResponse, err error) {
    if request == nil {
        request = NewDescribeStrategyAccountsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "fwm", APIVersion, "DescribeStrategyAccounts")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStrategyAccounts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStrategyAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStrategyDispatchStatusRequest() (request *DescribeStrategyDispatchStatusRequest) {
    request = &DescribeStrategyDispatchStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fwm", APIVersion, "DescribeStrategyDispatchStatus")
    
    
    return
}

func NewDescribeStrategyDispatchStatusResponse() (response *DescribeStrategyDispatchStatusResponse) {
    response = &DescribeStrategyDispatchStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStrategyDispatchStatus
// 查询策略下发状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeStrategyDispatchStatus(request *DescribeStrategyDispatchStatusRequest) (response *DescribeStrategyDispatchStatusResponse, err error) {
    return c.DescribeStrategyDispatchStatusWithContext(context.Background(), request)
}

// DescribeStrategyDispatchStatus
// 查询策略下发状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeStrategyDispatchStatusWithContext(ctx context.Context, request *DescribeStrategyDispatchStatusRequest) (response *DescribeStrategyDispatchStatusResponse, err error) {
    if request == nil {
        request = NewDescribeStrategyDispatchStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "fwm", APIVersion, "DescribeStrategyDispatchStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStrategyDispatchStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStrategyDispatchStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVpcAclRulesRequest() (request *DescribeVpcAclRulesRequest) {
    request = &DescribeVpcAclRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fwm", APIVersion, "DescribeVpcAclRules")
    
    
    return
}

func NewDescribeVpcAclRulesResponse() (response *DescribeVpcAclRulesResponse) {
    response = &DescribeVpcAclRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVpcAclRules
// 查询VPC ACL规则列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeVpcAclRules(request *DescribeVpcAclRulesRequest) (response *DescribeVpcAclRulesResponse, err error) {
    return c.DescribeVpcAclRulesWithContext(context.Background(), request)
}

// DescribeVpcAclRules
// 查询VPC ACL规则列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeVpcAclRulesWithContext(ctx context.Context, request *DescribeVpcAclRulesRequest) (response *DescribeVpcAclRulesResponse, err error) {
    if request == nil {
        request = NewDescribeVpcAclRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "fwm", APIVersion, "DescribeVpcAclRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVpcAclRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVpcAclRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDispatchStrategyRequest() (request *DispatchStrategyRequest) {
    request = &DispatchStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fwm", APIVersion, "DispatchStrategy")
    
    
    return
}

func NewDispatchStrategyResponse() (response *DispatchStrategyResponse) {
    response = &DispatchStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DispatchStrategy
// 下发策略
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DispatchStrategy(request *DispatchStrategyRequest) (response *DispatchStrategyResponse, err error) {
    return c.DispatchStrategyWithContext(context.Background(), request)
}

// DispatchStrategy
// 下发策略
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DispatchStrategyWithContext(ctx context.Context, request *DispatchStrategyRequest) (response *DispatchStrategyResponse, err error) {
    if request == nil {
        request = NewDispatchStrategyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "fwm", APIVersion, "DispatchStrategy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DispatchStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDispatchStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewIgnorePolicyRiskRequest() (request *IgnorePolicyRiskRequest) {
    request = &IgnorePolicyRiskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fwm", APIVersion, "IgnorePolicyRisk")
    
    
    return
}

func NewIgnorePolicyRiskResponse() (response *IgnorePolicyRiskResponse) {
    response = &IgnorePolicyRiskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// IgnorePolicyRisk
// 忽略策略问题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) IgnorePolicyRisk(request *IgnorePolicyRiskRequest) (response *IgnorePolicyRiskResponse, err error) {
    return c.IgnorePolicyRiskWithContext(context.Background(), request)
}

// IgnorePolicyRisk
// 忽略策略问题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) IgnorePolicyRiskWithContext(ctx context.Context, request *IgnorePolicyRiskRequest) (response *IgnorePolicyRiskResponse, err error) {
    if request == nil {
        request = NewIgnorePolicyRiskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "fwm", APIVersion, "IgnorePolicyRisk")
    
    if c.GetCredential() == nil {
        return nil, errors.New("IgnorePolicyRisk require credential")
    }

    request.SetContext(ctx)
    
    response = NewIgnorePolicyRiskResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEdgeAclRuleRequest() (request *ModifyEdgeAclRuleRequest) {
    request = &ModifyEdgeAclRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fwm", APIVersion, "ModifyEdgeAclRule")
    
    
    return
}

func NewModifyEdgeAclRuleResponse() (response *ModifyEdgeAclRuleResponse) {
    response = &ModifyEdgeAclRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyEdgeAclRule
// 修改互联网边界ACL规则。Rule 参数中必须包含 RuleId 用于指定要修改的规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyEdgeAclRule(request *ModifyEdgeAclRuleRequest) (response *ModifyEdgeAclRuleResponse, err error) {
    return c.ModifyEdgeAclRuleWithContext(context.Background(), request)
}

// ModifyEdgeAclRule
// 修改互联网边界ACL规则。Rule 参数中必须包含 RuleId 用于指定要修改的规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyEdgeAclRuleWithContext(ctx context.Context, request *ModifyEdgeAclRuleRequest) (response *ModifyEdgeAclRuleResponse, err error) {
    if request == nil {
        request = NewModifyEdgeAclRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "fwm", APIVersion, "ModifyEdgeAclRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEdgeAclRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyEdgeAclRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEdgeAclRuleSequenceRequest() (request *ModifyEdgeAclRuleSequenceRequest) {
    request = &ModifyEdgeAclRuleSequenceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fwm", APIVersion, "ModifyEdgeAclRuleSequence")
    
    
    return
}

func NewModifyEdgeAclRuleSequenceResponse() (response *ModifyEdgeAclRuleSequenceResponse) {
    response = &ModifyEdgeAclRuleSequenceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyEdgeAclRuleSequence
// 批量调整互联网边界ACL规则的执行顺序。Sequences 参数必须包含所有受影响的规则序号映射关系。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyEdgeAclRuleSequence(request *ModifyEdgeAclRuleSequenceRequest) (response *ModifyEdgeAclRuleSequenceResponse, err error) {
    return c.ModifyEdgeAclRuleSequenceWithContext(context.Background(), request)
}

// ModifyEdgeAclRuleSequence
// 批量调整互联网边界ACL规则的执行顺序。Sequences 参数必须包含所有受影响的规则序号映射关系。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyEdgeAclRuleSequenceWithContext(ctx context.Context, request *ModifyEdgeAclRuleSequenceRequest) (response *ModifyEdgeAclRuleSequenceResponse, err error) {
    if request == nil {
        request = NewModifyEdgeAclRuleSequenceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "fwm", APIVersion, "ModifyEdgeAclRuleSequence")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEdgeAclRuleSequence require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyEdgeAclRuleSequenceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNatAclRuleRequest() (request *ModifyNatAclRuleRequest) {
    request = &ModifyNatAclRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fwm", APIVersion, "ModifyNatAclRule")
    
    
    return
}

func NewModifyNatAclRuleResponse() (response *ModifyNatAclRuleResponse) {
    response = &ModifyNatAclRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyNatAclRule
// 修改NAT ACL规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyNatAclRule(request *ModifyNatAclRuleRequest) (response *ModifyNatAclRuleResponse, err error) {
    return c.ModifyNatAclRuleWithContext(context.Background(), request)
}

// ModifyNatAclRule
// 修改NAT ACL规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyNatAclRuleWithContext(ctx context.Context, request *ModifyNatAclRuleRequest) (response *ModifyNatAclRuleResponse, err error) {
    if request == nil {
        request = NewModifyNatAclRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "fwm", APIVersion, "ModifyNatAclRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNatAclRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNatAclRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNatAclRuleSequenceRequest() (request *ModifyNatAclRuleSequenceRequest) {
    request = &ModifyNatAclRuleSequenceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fwm", APIVersion, "ModifyNatAclRuleSequence")
    
    
    return
}

func NewModifyNatAclRuleSequenceResponse() (response *ModifyNatAclRuleSequenceResponse) {
    response = &ModifyNatAclRuleSequenceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyNatAclRuleSequence
// 调整NAT ACL规则优先级顺序
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyNatAclRuleSequence(request *ModifyNatAclRuleSequenceRequest) (response *ModifyNatAclRuleSequenceResponse, err error) {
    return c.ModifyNatAclRuleSequenceWithContext(context.Background(), request)
}

// ModifyNatAclRuleSequence
// 调整NAT ACL规则优先级顺序
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyNatAclRuleSequenceWithContext(ctx context.Context, request *ModifyNatAclRuleSequenceRequest) (response *ModifyNatAclRuleSequenceResponse, err error) {
    if request == nil {
        request = NewModifyNatAclRuleSequenceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "fwm", APIVersion, "ModifyNatAclRuleSequence")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNatAclRuleSequence require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNatAclRuleSequenceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRuleGroupRequest() (request *ModifyRuleGroupRequest) {
    request = &ModifyRuleGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fwm", APIVersion, "ModifyRuleGroup")
    
    
    return
}

func NewModifyRuleGroupResponse() (response *ModifyRuleGroupResponse) {
    response = &ModifyRuleGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRuleGroup
// 修改规则组信息（规则组管理）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyRuleGroup(request *ModifyRuleGroupRequest) (response *ModifyRuleGroupResponse, err error) {
    return c.ModifyRuleGroupWithContext(context.Background(), request)
}

// ModifyRuleGroup
// 修改规则组信息（规则组管理）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyRuleGroupWithContext(ctx context.Context, request *ModifyRuleGroupRequest) (response *ModifyRuleGroupResponse, err error) {
    if request == nil {
        request = NewModifyRuleGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "fwm", APIVersion, "ModifyRuleGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRuleGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRuleGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifySecurityGroupRuleRequest() (request *ModifySecurityGroupRuleRequest) {
    request = &ModifySecurityGroupRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fwm", APIVersion, "ModifySecurityGroupRule")
    
    
    return
}

func NewModifySecurityGroupRuleResponse() (response *ModifySecurityGroupRuleResponse) {
    response = &ModifySecurityGroupRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySecurityGroupRule
// 修改规则（规则组管理）
//
// 可能返回的错误码:
//  MISSINGPARAMETER_COMMONPARAM = "MissingParameter.CommonParam"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
func (c *Client) ModifySecurityGroupRule(request *ModifySecurityGroupRuleRequest) (response *ModifySecurityGroupRuleResponse, err error) {
    return c.ModifySecurityGroupRuleWithContext(context.Background(), request)
}

// ModifySecurityGroupRule
// 修改规则（规则组管理）
//
// 可能返回的错误码:
//  MISSINGPARAMETER_COMMONPARAM = "MissingParameter.CommonParam"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
func (c *Client) ModifySecurityGroupRuleWithContext(ctx context.Context, request *ModifySecurityGroupRuleRequest) (response *ModifySecurityGroupRuleResponse, err error) {
    if request == nil {
        request = NewModifySecurityGroupRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "fwm", APIVersion, "ModifySecurityGroupRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySecurityGroupRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySecurityGroupRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyStrategyRequest() (request *ModifyStrategyRequest) {
    request = &ModifyStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fwm", APIVersion, "ModifyStrategy")
    
    
    return
}

func NewModifyStrategyResponse() (response *ModifyStrategyResponse) {
    response = &ModifyStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyStrategy
// 修改策略信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
func (c *Client) ModifyStrategy(request *ModifyStrategyRequest) (response *ModifyStrategyResponse, err error) {
    return c.ModifyStrategyWithContext(context.Background(), request)
}

// ModifyStrategy
// 修改策略信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
func (c *Client) ModifyStrategyWithContext(ctx context.Context, request *ModifyStrategyRequest) (response *ModifyStrategyResponse, err error) {
    if request == nil {
        request = NewModifyStrategyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "fwm", APIVersion, "ModifyStrategy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyStrategySequenceRequest() (request *ModifyStrategySequenceRequest) {
    request = &ModifyStrategySequenceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fwm", APIVersion, "ModifyStrategySequence")
    
    
    return
}

func NewModifyStrategySequenceResponse() (response *ModifyStrategySequenceResponse) {
    response = &ModifyStrategySequenceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyStrategySequence
// 快速排序修改策略优先级
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyStrategySequence(request *ModifyStrategySequenceRequest) (response *ModifyStrategySequenceResponse, err error) {
    return c.ModifyStrategySequenceWithContext(context.Background(), request)
}

// ModifyStrategySequence
// 快速排序修改策略优先级
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyStrategySequenceWithContext(ctx context.Context, request *ModifyStrategySequenceRequest) (response *ModifyStrategySequenceResponse, err error) {
    if request == nil {
        request = NewModifyStrategySequenceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "fwm", APIVersion, "ModifyStrategySequence")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyStrategySequence require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyStrategySequenceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVpcAclRuleRequest() (request *ModifyVpcAclRuleRequest) {
    request = &ModifyVpcAclRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fwm", APIVersion, "ModifyVpcAclRule")
    
    
    return
}

func NewModifyVpcAclRuleResponse() (response *ModifyVpcAclRuleResponse) {
    response = &ModifyVpcAclRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyVpcAclRule
// 修改VPC ACL规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyVpcAclRule(request *ModifyVpcAclRuleRequest) (response *ModifyVpcAclRuleResponse, err error) {
    return c.ModifyVpcAclRuleWithContext(context.Background(), request)
}

// ModifyVpcAclRule
// 修改VPC ACL规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyVpcAclRuleWithContext(ctx context.Context, request *ModifyVpcAclRuleRequest) (response *ModifyVpcAclRuleResponse, err error) {
    if request == nil {
        request = NewModifyVpcAclRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "fwm", APIVersion, "ModifyVpcAclRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVpcAclRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyVpcAclRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVpcAclRuleSequenceRequest() (request *ModifyVpcAclRuleSequenceRequest) {
    request = &ModifyVpcAclRuleSequenceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fwm", APIVersion, "ModifyVpcAclRuleSequence")
    
    
    return
}

func NewModifyVpcAclRuleSequenceResponse() (response *ModifyVpcAclRuleSequenceResponse) {
    response = &ModifyVpcAclRuleSequenceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyVpcAclRuleSequence
// 调整VPC ACL规则优先级顺序
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyVpcAclRuleSequence(request *ModifyVpcAclRuleSequenceRequest) (response *ModifyVpcAclRuleSequenceResponse, err error) {
    return c.ModifyVpcAclRuleSequenceWithContext(context.Background(), request)
}

// ModifyVpcAclRuleSequence
// 调整VPC ACL规则优先级顺序
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyVpcAclRuleSequenceWithContext(ctx context.Context, request *ModifyVpcAclRuleSequenceRequest) (response *ModifyVpcAclRuleSequenceResponse, err error) {
    if request == nil {
        request = NewModifyVpcAclRuleSequenceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "fwm", APIVersion, "ModifyVpcAclRuleSequence")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVpcAclRuleSequence require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyVpcAclRuleSequenceResponse()
    err = c.Send(request, response)
    return
}
