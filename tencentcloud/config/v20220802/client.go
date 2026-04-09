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

package v20220802

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2022-08-02"

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


func NewAddAggregateCompliancePackRequest() (request *AddAggregateCompliancePackRequest) {
    request = &AddAggregateCompliancePackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "AddAggregateCompliancePack")
    
    
    return
}

func NewAddAggregateCompliancePackResponse() (response *AddAggregateCompliancePackResponse) {
    response = &AddAggregateCompliancePackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddAggregateCompliancePack
// 账号组新建合规包
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPLIANCEPACKNAMEISNOTEXIST = "FailedOperation.CompliancePackNameIsNotExist"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_COMPLIANCEPACKMAXNUM = "LimitExceeded.CompliancePackMaxNum"
//  RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"
func (c *Client) AddAggregateCompliancePack(request *AddAggregateCompliancePackRequest) (response *AddAggregateCompliancePackResponse, err error) {
    return c.AddAggregateCompliancePackWithContext(context.Background(), request)
}

// AddAggregateCompliancePack
// 账号组新建合规包
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPLIANCEPACKNAMEISNOTEXIST = "FailedOperation.CompliancePackNameIsNotExist"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_COMPLIANCEPACKMAXNUM = "LimitExceeded.CompliancePackMaxNum"
//  RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"
func (c *Client) AddAggregateCompliancePackWithContext(ctx context.Context, request *AddAggregateCompliancePackRequest) (response *AddAggregateCompliancePackResponse, err error) {
    if request == nil {
        request = NewAddAggregateCompliancePackRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "AddAggregateCompliancePack")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddAggregateCompliancePack require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddAggregateCompliancePackResponse()
    err = c.Send(request, response)
    return
}

func NewAddAggregateConfigRuleRequest() (request *AddAggregateConfigRuleRequest) {
    request = &AddAggregateConfigRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "AddAggregateConfigRule")
    
    
    return
}

func NewAddAggregateConfigRuleResponse() (response *AddAggregateConfigRuleResponse) {
    response = &AddAggregateConfigRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddAggregateConfigRule
// 账号组新建规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_RULENAMEISEXIST = "FailedOperation.RuleNameIsExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ACCOUNTGROUPISNOTEXIST = "ResourceNotFound.AccountGroupIsNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AddAggregateConfigRule(request *AddAggregateConfigRuleRequest) (response *AddAggregateConfigRuleResponse, err error) {
    return c.AddAggregateConfigRuleWithContext(context.Background(), request)
}

// AddAggregateConfigRule
// 账号组新建规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_RULENAMEISEXIST = "FailedOperation.RuleNameIsExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ACCOUNTGROUPISNOTEXIST = "ResourceNotFound.AccountGroupIsNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AddAggregateConfigRuleWithContext(ctx context.Context, request *AddAggregateConfigRuleRequest) (response *AddAggregateConfigRuleResponse, err error) {
    if request == nil {
        request = NewAddAggregateConfigRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "AddAggregateConfigRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddAggregateConfigRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddAggregateConfigRuleResponse()
    err = c.Send(request, response)
    return
}

func NewAddCompliancePackRequest() (request *AddCompliancePackRequest) {
    request = &AddCompliancePackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "AddCompliancePack")
    
    
    return
}

func NewAddCompliancePackResponse() (response *AddCompliancePackResponse) {
    response = &AddCompliancePackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddCompliancePack
// 新建合规包
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPLIANCEPACKNAMEISNOTEXIST = "FailedOperation.CompliancePackNameIsNotExist"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_COMPLIANCEPACKMAXNUM = "LimitExceeded.CompliancePackMaxNum"
//  RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"
func (c *Client) AddCompliancePack(request *AddCompliancePackRequest) (response *AddCompliancePackResponse, err error) {
    return c.AddCompliancePackWithContext(context.Background(), request)
}

// AddCompliancePack
// 新建合规包
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPLIANCEPACKNAMEISNOTEXIST = "FailedOperation.CompliancePackNameIsNotExist"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_COMPLIANCEPACKMAXNUM = "LimitExceeded.CompliancePackMaxNum"
//  RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"
func (c *Client) AddCompliancePackWithContext(ctx context.Context, request *AddCompliancePackRequest) (response *AddCompliancePackResponse, err error) {
    if request == nil {
        request = NewAddCompliancePackRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "AddCompliancePack")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddCompliancePack require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddCompliancePackResponse()
    err = c.Send(request, response)
    return
}

func NewAddConfigRuleRequest() (request *AddConfigRuleRequest) {
    request = &AddConfigRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "AddConfigRule")
    
    
    return
}

func NewAddConfigRuleResponse() (response *AddConfigRuleResponse) {
    response = &AddConfigRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddConfigRule
// 新建 规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_RULENAMEISEXIST = "FailedOperation.RuleNameIsExist"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) AddConfigRule(request *AddConfigRuleRequest) (response *AddConfigRuleResponse, err error) {
    return c.AddConfigRuleWithContext(context.Background(), request)
}

// AddConfigRule
// 新建 规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_RULENAMEISEXIST = "FailedOperation.RuleNameIsExist"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) AddConfigRuleWithContext(ctx context.Context, request *AddConfigRuleRequest) (response *AddConfigRuleResponse, err error) {
    if request == nil {
        request = NewAddConfigRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "AddConfigRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddConfigRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddConfigRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCloseAggregateConfigRuleRequest() (request *CloseAggregateConfigRuleRequest) {
    request = &CloseAggregateConfigRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "CloseAggregateConfigRule")
    
    
    return
}

func NewCloseAggregateConfigRuleResponse() (response *CloseAggregateConfigRuleResponse) {
    response = &CloseAggregateConfigRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CloseAggregateConfigRule
// 账号组关闭规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTALLOWOPERATECONTROLCENTERGOVERNANCEPACKAGERULE = "FailedOperation.NotAllowOperateControlCenterGovernancePackageRule"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ACCOUNTGROUPISNOTEXIST = "ResourceNotFound.AccountGroupIsNotExist"
//  RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CloseAggregateConfigRule(request *CloseAggregateConfigRuleRequest) (response *CloseAggregateConfigRuleResponse, err error) {
    return c.CloseAggregateConfigRuleWithContext(context.Background(), request)
}

// CloseAggregateConfigRule
// 账号组关闭规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTALLOWOPERATECONTROLCENTERGOVERNANCEPACKAGERULE = "FailedOperation.NotAllowOperateControlCenterGovernancePackageRule"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ACCOUNTGROUPISNOTEXIST = "ResourceNotFound.AccountGroupIsNotExist"
//  RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CloseAggregateConfigRuleWithContext(ctx context.Context, request *CloseAggregateConfigRuleRequest) (response *CloseAggregateConfigRuleResponse, err error) {
    if request == nil {
        request = NewCloseAggregateConfigRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "CloseAggregateConfigRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseAggregateConfigRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloseAggregateConfigRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCloseConfigRecorderRequest() (request *CloseConfigRecorderRequest) {
    request = &CloseConfigRecorderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "CloseConfigRecorder")
    
    
    return
}

func NewCloseConfigRecorderResponse() (response *CloseConfigRecorderResponse) {
    response = &CloseConfigRecorderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CloseConfigRecorder
// 资源监控管理-关闭监控
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_RECORDERISNOTEXIST = "ResourceNotFound.RecorderIsNotExist"
func (c *Client) CloseConfigRecorder(request *CloseConfigRecorderRequest) (response *CloseConfigRecorderResponse, err error) {
    return c.CloseConfigRecorderWithContext(context.Background(), request)
}

// CloseConfigRecorder
// 资源监控管理-关闭监控
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_RECORDERISNOTEXIST = "ResourceNotFound.RecorderIsNotExist"
func (c *Client) CloseConfigRecorderWithContext(ctx context.Context, request *CloseConfigRecorderRequest) (response *CloseConfigRecorderResponse, err error) {
    if request == nil {
        request = NewCloseConfigRecorderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "CloseConfigRecorder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseConfigRecorder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloseConfigRecorderResponse()
    err = c.Send(request, response)
    return
}

func NewCloseConfigRuleRequest() (request *CloseConfigRuleRequest) {
    request = &CloseConfigRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "CloseConfigRule")
    
    
    return
}

func NewCloseConfigRuleResponse() (response *CloseConfigRuleResponse) {
    response = &CloseConfigRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CloseConfigRule
// 关闭规则
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"
func (c *Client) CloseConfigRule(request *CloseConfigRuleRequest) (response *CloseConfigRuleResponse, err error) {
    return c.CloseConfigRuleWithContext(context.Background(), request)
}

// CloseConfigRule
// 关闭规则
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"
func (c *Client) CloseConfigRuleWithContext(ctx context.Context, request *CloseConfigRuleRequest) (response *CloseConfigRuleResponse, err error) {
    if request == nil {
        request = NewCloseConfigRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "CloseConfigRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseConfigRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloseConfigRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAggregatorRequest() (request *CreateAggregatorRequest) {
    request = &CreateAggregatorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "CreateAggregator")
    
    
    return
}

func NewCreateAggregatorResponse() (response *CreateAggregatorResponse) {
    response = &CreateAggregatorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAggregator
// 新建账号组
//
// 可能返回的错误码:
//  FAILEDOPERATION_SAMEMEMBERACCOUNTGROUPALREADYEXIST = "FailedOperation.SameMemberAccountGroupAlreadyExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_ACCOUNTGROUPLIMITOVER = "LimitExceeded.AccountGroupLimitOver"
func (c *Client) CreateAggregator(request *CreateAggregatorRequest) (response *CreateAggregatorResponse, err error) {
    return c.CreateAggregatorWithContext(context.Background(), request)
}

// CreateAggregator
// 新建账号组
//
// 可能返回的错误码:
//  FAILEDOPERATION_SAMEMEMBERACCOUNTGROUPALREADYEXIST = "FailedOperation.SameMemberAccountGroupAlreadyExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_ACCOUNTGROUPLIMITOVER = "LimitExceeded.AccountGroupLimitOver"
func (c *Client) CreateAggregatorWithContext(ctx context.Context, request *CreateAggregatorRequest) (response *CreateAggregatorResponse, err error) {
    if request == nil {
        request = NewCreateAggregatorRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "CreateAggregator")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAggregator require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAggregatorResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRemediationRequest() (request *CreateRemediationRequest) {
    request = &CreateRemediationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "CreateRemediation")
    
    
    return
}

func NewCreateRemediationResponse() (response *CreateRemediationResponse) {
    response = &CreateRemediationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRemediation
// 新增规则修正设置
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_REMEDIATIONMAXNUM = "LimitExceeded.RemediationMaxNum"
func (c *Client) CreateRemediation(request *CreateRemediationRequest) (response *CreateRemediationResponse, err error) {
    return c.CreateRemediationWithContext(context.Background(), request)
}

// CreateRemediation
// 新增规则修正设置
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_REMEDIATIONMAXNUM = "LimitExceeded.RemediationMaxNum"
func (c *Client) CreateRemediationWithContext(ctx context.Context, request *CreateRemediationRequest) (response *CreateRemediationResponse, err error) {
    if request == nil {
        request = NewCreateRemediationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "CreateRemediation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRemediation require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRemediationResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAggregateCompliancePackRequest() (request *DeleteAggregateCompliancePackRequest) {
    request = &DeleteAggregateCompliancePackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "DeleteAggregateCompliancePack")
    
    
    return
}

func NewDeleteAggregateCompliancePackResponse() (response *DeleteAggregateCompliancePackResponse) {
    response = &DeleteAggregateCompliancePackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAggregateCompliancePack
// 账号组删除合规包
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPLIANCEPACKISNOTEXIST = "FailedOperation.CompliancePackIsNotExist"
//  FAILEDOPERATION_COMPLIANCEPACKISNOTSTOP = "FailedOperation.CompliancePackIsNotStop"
//  FAILEDOPERATION_NOTALLOWOPERATECONTROLCENTERGOVERNANCEPACKAGE = "FailedOperation.NotAllowOperateControlCenterGovernancePackage"
//  FAILEDOPERATION_RULEISNOTSTOP = "FailedOperation.RuleIsNotStop"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_COMPLIANCEPACKISNOTEXIST = "ResourceNotFound.CompliancePackIsNotExist"
func (c *Client) DeleteAggregateCompliancePack(request *DeleteAggregateCompliancePackRequest) (response *DeleteAggregateCompliancePackResponse, err error) {
    return c.DeleteAggregateCompliancePackWithContext(context.Background(), request)
}

// DeleteAggregateCompliancePack
// 账号组删除合规包
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPLIANCEPACKISNOTEXIST = "FailedOperation.CompliancePackIsNotExist"
//  FAILEDOPERATION_COMPLIANCEPACKISNOTSTOP = "FailedOperation.CompliancePackIsNotStop"
//  FAILEDOPERATION_NOTALLOWOPERATECONTROLCENTERGOVERNANCEPACKAGE = "FailedOperation.NotAllowOperateControlCenterGovernancePackage"
//  FAILEDOPERATION_RULEISNOTSTOP = "FailedOperation.RuleIsNotStop"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_COMPLIANCEPACKISNOTEXIST = "ResourceNotFound.CompliancePackIsNotExist"
func (c *Client) DeleteAggregateCompliancePackWithContext(ctx context.Context, request *DeleteAggregateCompliancePackRequest) (response *DeleteAggregateCompliancePackResponse, err error) {
    if request == nil {
        request = NewDeleteAggregateCompliancePackRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "DeleteAggregateCompliancePack")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAggregateCompliancePack require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAggregateCompliancePackResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAggregateConfigRuleRequest() (request *DeleteAggregateConfigRuleRequest) {
    request = &DeleteAggregateConfigRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "DeleteAggregateConfigRule")
    
    
    return
}

func NewDeleteAggregateConfigRuleResponse() (response *DeleteAggregateConfigRuleResponse) {
    response = &DeleteAggregateConfigRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAggregateConfigRule
// 账号组删除规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTALLOWOPERATECONTROLCENTERGOVERNANCEPACKAGERULE = "FailedOperation.NotAllowOperateControlCenterGovernancePackageRule"
//  FAILEDOPERATION_RULEISNOTSTOP = "FailedOperation.RuleIsNotStop"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ACCOUNTGROUPISNOTEXIST = "ResourceNotFound.AccountGroupIsNotExist"
//  RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteAggregateConfigRule(request *DeleteAggregateConfigRuleRequest) (response *DeleteAggregateConfigRuleResponse, err error) {
    return c.DeleteAggregateConfigRuleWithContext(context.Background(), request)
}

// DeleteAggregateConfigRule
// 账号组删除规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTALLOWOPERATECONTROLCENTERGOVERNANCEPACKAGERULE = "FailedOperation.NotAllowOperateControlCenterGovernancePackageRule"
//  FAILEDOPERATION_RULEISNOTSTOP = "FailedOperation.RuleIsNotStop"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ACCOUNTGROUPISNOTEXIST = "ResourceNotFound.AccountGroupIsNotExist"
//  RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteAggregateConfigRuleWithContext(ctx context.Context, request *DeleteAggregateConfigRuleRequest) (response *DeleteAggregateConfigRuleResponse, err error) {
    if request == nil {
        request = NewDeleteAggregateConfigRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "DeleteAggregateConfigRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAggregateConfigRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAggregateConfigRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCompliancePackRequest() (request *DeleteCompliancePackRequest) {
    request = &DeleteCompliancePackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "DeleteCompliancePack")
    
    
    return
}

func NewDeleteCompliancePackResponse() (response *DeleteCompliancePackResponse) {
    response = &DeleteCompliancePackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCompliancePack
// 删除合规包
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPLIANCEPACKISNOTEXIST = "FailedOperation.CompliancePackIsNotExist"
//  FAILEDOPERATION_COMPLIANCEPACKISNOTSTOP = "FailedOperation.CompliancePackIsNotStop"
//  FAILEDOPERATION_RULEISNOTSTOP = "FailedOperation.RuleIsNotStop"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_COMPLIANCEPACKISNOTEXIST = "ResourceNotFound.CompliancePackIsNotExist"
func (c *Client) DeleteCompliancePack(request *DeleteCompliancePackRequest) (response *DeleteCompliancePackResponse, err error) {
    return c.DeleteCompliancePackWithContext(context.Background(), request)
}

// DeleteCompliancePack
// 删除合规包
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPLIANCEPACKISNOTEXIST = "FailedOperation.CompliancePackIsNotExist"
//  FAILEDOPERATION_COMPLIANCEPACKISNOTSTOP = "FailedOperation.CompliancePackIsNotStop"
//  FAILEDOPERATION_RULEISNOTSTOP = "FailedOperation.RuleIsNotStop"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_COMPLIANCEPACKISNOTEXIST = "ResourceNotFound.CompliancePackIsNotExist"
func (c *Client) DeleteCompliancePackWithContext(ctx context.Context, request *DeleteCompliancePackRequest) (response *DeleteCompliancePackResponse, err error) {
    if request == nil {
        request = NewDeleteCompliancePackRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "DeleteCompliancePack")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCompliancePack require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCompliancePackResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteConfigRuleRequest() (request *DeleteConfigRuleRequest) {
    request = &DeleteConfigRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "DeleteConfigRule")
    
    
    return
}

func NewDeleteConfigRuleResponse() (response *DeleteConfigRuleResponse) {
    response = &DeleteConfigRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteConfigRule
// 删除规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_RULEISNOTSTOP = "FailedOperation.RuleIsNotStop"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"
func (c *Client) DeleteConfigRule(request *DeleteConfigRuleRequest) (response *DeleteConfigRuleResponse, err error) {
    return c.DeleteConfigRuleWithContext(context.Background(), request)
}

// DeleteConfigRule
// 删除规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_RULEISNOTSTOP = "FailedOperation.RuleIsNotStop"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"
func (c *Client) DeleteConfigRuleWithContext(ctx context.Context, request *DeleteConfigRuleRequest) (response *DeleteConfigRuleResponse, err error) {
    if request == nil {
        request = NewDeleteConfigRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "DeleteConfigRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteConfigRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteConfigRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRemediationsRequest() (request *DeleteRemediationsRequest) {
    request = &DeleteRemediationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "DeleteRemediations")
    
    
    return
}

func NewDeleteRemediationsResponse() (response *DeleteRemediationsResponse) {
    response = &DeleteRemediationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRemediations
// 删除规则修正设置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_REMEDIATIONISNOTEXIST = "ResourceNotFound.RemediationIsNotExist"
func (c *Client) DeleteRemediations(request *DeleteRemediationsRequest) (response *DeleteRemediationsResponse, err error) {
    return c.DeleteRemediationsWithContext(context.Background(), request)
}

// DeleteRemediations
// 删除规则修正设置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_REMEDIATIONISNOTEXIST = "ResourceNotFound.RemediationIsNotExist"
func (c *Client) DeleteRemediationsWithContext(ctx context.Context, request *DeleteRemediationsRequest) (response *DeleteRemediationsResponse, err error) {
    if request == nil {
        request = NewDeleteRemediationsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "DeleteRemediations")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRemediations require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRemediationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAggregateCompliancePackRequest() (request *DescribeAggregateCompliancePackRequest) {
    request = &DescribeAggregateCompliancePackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "DescribeAggregateCompliancePack")
    
    
    return
}

func NewDescribeAggregateCompliancePackResponse() (response *DescribeAggregateCompliancePackResponse) {
    response = &DescribeAggregateCompliancePackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAggregateCompliancePack
// 账号组合规包详情
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAggregateCompliancePack(request *DescribeAggregateCompliancePackRequest) (response *DescribeAggregateCompliancePackResponse, err error) {
    return c.DescribeAggregateCompliancePackWithContext(context.Background(), request)
}

// DescribeAggregateCompliancePack
// 账号组合规包详情
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAggregateCompliancePackWithContext(ctx context.Context, request *DescribeAggregateCompliancePackRequest) (response *DescribeAggregateCompliancePackResponse, err error) {
    if request == nil {
        request = NewDescribeAggregateCompliancePackRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "DescribeAggregateCompliancePack")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAggregateCompliancePack require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAggregateCompliancePackResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAggregateConfigDeliverRequest() (request *DescribeAggregateConfigDeliverRequest) {
    request = &DescribeAggregateConfigDeliverRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "DescribeAggregateConfigDeliver")
    
    
    return
}

func NewDescribeAggregateConfigDeliverResponse() (response *DescribeAggregateConfigDeliverResponse) {
    response = &DescribeAggregateConfigDeliverResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAggregateConfigDeliver
// 账号组获取投递设置详情
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_DELIVEREVENTNOTEXIST = "ResourceNotFound.DeliverEventNotExist"
//  RESOURCENOTFOUND_DELIVERISNOTEXIST = "ResourceNotFound.DeliverIsNotExist"
func (c *Client) DescribeAggregateConfigDeliver(request *DescribeAggregateConfigDeliverRequest) (response *DescribeAggregateConfigDeliverResponse, err error) {
    return c.DescribeAggregateConfigDeliverWithContext(context.Background(), request)
}

// DescribeAggregateConfigDeliver
// 账号组获取投递设置详情
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_DELIVEREVENTNOTEXIST = "ResourceNotFound.DeliverEventNotExist"
//  RESOURCENOTFOUND_DELIVERISNOTEXIST = "ResourceNotFound.DeliverIsNotExist"
func (c *Client) DescribeAggregateConfigDeliverWithContext(ctx context.Context, request *DescribeAggregateConfigDeliverRequest) (response *DescribeAggregateConfigDeliverResponse, err error) {
    if request == nil {
        request = NewDescribeAggregateConfigDeliverRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "DescribeAggregateConfigDeliver")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAggregateConfigDeliver require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAggregateConfigDeliverResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAggregateConfigRuleRequest() (request *DescribeAggregateConfigRuleRequest) {
    request = &DescribeAggregateConfigRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "DescribeAggregateConfigRule")
    
    
    return
}

func NewDescribeAggregateConfigRuleResponse() (response *DescribeAggregateConfigRuleResponse) {
    response = &DescribeAggregateConfigRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAggregateConfigRule
// 账号组获取规则详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ACCOUNTGROUPISNOTEXIST = "ResourceNotFound.AccountGroupIsNotExist"
//  RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"
func (c *Client) DescribeAggregateConfigRule(request *DescribeAggregateConfigRuleRequest) (response *DescribeAggregateConfigRuleResponse, err error) {
    return c.DescribeAggregateConfigRuleWithContext(context.Background(), request)
}

// DescribeAggregateConfigRule
// 账号组获取规则详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ACCOUNTGROUPISNOTEXIST = "ResourceNotFound.AccountGroupIsNotExist"
//  RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"
func (c *Client) DescribeAggregateConfigRuleWithContext(ctx context.Context, request *DescribeAggregateConfigRuleRequest) (response *DescribeAggregateConfigRuleResponse, err error) {
    if request == nil {
        request = NewDescribeAggregateConfigRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "DescribeAggregateConfigRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAggregateConfigRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAggregateConfigRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAggregateDiscoveredResourceRequest() (request *DescribeAggregateDiscoveredResourceRequest) {
    request = &DescribeAggregateDiscoveredResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "DescribeAggregateDiscoveredResource")
    
    
    return
}

func NewDescribeAggregateDiscoveredResourceResponse() (response *DescribeAggregateDiscoveredResourceResponse) {
    response = &DescribeAggregateDiscoveredResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAggregateDiscoveredResource
// 账号组资源详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ACCOUNTGROUPISNOTEXIST = "ResourceNotFound.AccountGroupIsNotExist"
//  RESOURCENOTFOUND_RESOURCENOTEXIST = "ResourceNotFound.ResourceNotExist"
func (c *Client) DescribeAggregateDiscoveredResource(request *DescribeAggregateDiscoveredResourceRequest) (response *DescribeAggregateDiscoveredResourceResponse, err error) {
    return c.DescribeAggregateDiscoveredResourceWithContext(context.Background(), request)
}

// DescribeAggregateDiscoveredResource
// 账号组资源详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ACCOUNTGROUPISNOTEXIST = "ResourceNotFound.AccountGroupIsNotExist"
//  RESOURCENOTFOUND_RESOURCENOTEXIST = "ResourceNotFound.ResourceNotExist"
func (c *Client) DescribeAggregateDiscoveredResourceWithContext(ctx context.Context, request *DescribeAggregateDiscoveredResourceRequest) (response *DescribeAggregateDiscoveredResourceResponse, err error) {
    if request == nil {
        request = NewDescribeAggregateDiscoveredResourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "DescribeAggregateDiscoveredResource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAggregateDiscoveredResource require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAggregateDiscoveredResourceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAggregatorRequest() (request *DescribeAggregatorRequest) {
    request = &DescribeAggregatorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "DescribeAggregator")
    
    
    return
}

func NewDescribeAggregatorResponse() (response *DescribeAggregatorResponse) {
    response = &DescribeAggregatorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAggregator
// 账号组详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ACCOUNTGROUPISNOTEXIST = "ResourceNotFound.AccountGroupIsNotExist"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
func (c *Client) DescribeAggregator(request *DescribeAggregatorRequest) (response *DescribeAggregatorResponse, err error) {
    return c.DescribeAggregatorWithContext(context.Background(), request)
}

// DescribeAggregator
// 账号组详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ACCOUNTGROUPISNOTEXIST = "ResourceNotFound.AccountGroupIsNotExist"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
func (c *Client) DescribeAggregatorWithContext(ctx context.Context, request *DescribeAggregatorRequest) (response *DescribeAggregatorResponse, err error) {
    if request == nil {
        request = NewDescribeAggregatorRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "DescribeAggregator")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAggregator require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAggregatorResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCompliancePackRequest() (request *DescribeCompliancePackRequest) {
    request = &DescribeCompliancePackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "DescribeCompliancePack")
    
    
    return
}

func NewDescribeCompliancePackResponse() (response *DescribeCompliancePackResponse) {
    response = &DescribeCompliancePackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCompliancePack
// 合规包详情
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeCompliancePack(request *DescribeCompliancePackRequest) (response *DescribeCompliancePackResponse, err error) {
    return c.DescribeCompliancePackWithContext(context.Background(), request)
}

// DescribeCompliancePack
// 合规包详情
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeCompliancePackWithContext(ctx context.Context, request *DescribeCompliancePackRequest) (response *DescribeCompliancePackResponse, err error) {
    if request == nil {
        request = NewDescribeCompliancePackRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "DescribeCompliancePack")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCompliancePack require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCompliancePackResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigDeliverRequest() (request *DescribeConfigDeliverRequest) {
    request = &DescribeConfigDeliverRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "DescribeConfigDeliver")
    
    
    return
}

func NewDescribeConfigDeliverResponse() (response *DescribeConfigDeliverResponse) {
    response = &DescribeConfigDeliverResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConfigDeliver
// 获取投递设置详情
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_DELIVEREVENTNOTEXIST = "ResourceNotFound.DeliverEventNotExist"
//  RESOURCENOTFOUND_DELIVERISNOTEXIST = "ResourceNotFound.DeliverIsNotExist"
func (c *Client) DescribeConfigDeliver(request *DescribeConfigDeliverRequest) (response *DescribeConfigDeliverResponse, err error) {
    return c.DescribeConfigDeliverWithContext(context.Background(), request)
}

// DescribeConfigDeliver
// 获取投递设置详情
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_DELIVEREVENTNOTEXIST = "ResourceNotFound.DeliverEventNotExist"
//  RESOURCENOTFOUND_DELIVERISNOTEXIST = "ResourceNotFound.DeliverIsNotExist"
func (c *Client) DescribeConfigDeliverWithContext(ctx context.Context, request *DescribeConfigDeliverRequest) (response *DescribeConfigDeliverResponse, err error) {
    if request == nil {
        request = NewDescribeConfigDeliverRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "DescribeConfigDeliver")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConfigDeliver require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConfigDeliverResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigRecorderRequest() (request *DescribeConfigRecorderRequest) {
    request = &DescribeConfigRecorderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "DescribeConfigRecorder")
    
    
    return
}

func NewDescribeConfigRecorderResponse() (response *DescribeConfigRecorderResponse) {
    response = &DescribeConfigRecorderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConfigRecorder
// 获取监控详情
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_RECORDERISNOTEXIST = "ResourceNotFound.RecorderIsNotExist"
func (c *Client) DescribeConfigRecorder(request *DescribeConfigRecorderRequest) (response *DescribeConfigRecorderResponse, err error) {
    return c.DescribeConfigRecorderWithContext(context.Background(), request)
}

// DescribeConfigRecorder
// 获取监控详情
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_RECORDERISNOTEXIST = "ResourceNotFound.RecorderIsNotExist"
func (c *Client) DescribeConfigRecorderWithContext(ctx context.Context, request *DescribeConfigRecorderRequest) (response *DescribeConfigRecorderResponse, err error) {
    if request == nil {
        request = NewDescribeConfigRecorderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "DescribeConfigRecorder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConfigRecorder require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConfigRecorderResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigRuleRequest() (request *DescribeConfigRuleRequest) {
    request = &DescribeConfigRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "DescribeConfigRule")
    
    
    return
}

func NewDescribeConfigRuleResponse() (response *DescribeConfigRuleResponse) {
    response = &DescribeConfigRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConfigRule
// 获取规则详情
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"
func (c *Client) DescribeConfigRule(request *DescribeConfigRuleRequest) (response *DescribeConfigRuleResponse, err error) {
    return c.DescribeConfigRuleWithContext(context.Background(), request)
}

// DescribeConfigRule
// 获取规则详情
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"
func (c *Client) DescribeConfigRuleWithContext(ctx context.Context, request *DescribeConfigRuleRequest) (response *DescribeConfigRuleResponse, err error) {
    if request == nil {
        request = NewDescribeConfigRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "DescribeConfigRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConfigRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConfigRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDiscoveredResourceRequest() (request *DescribeDiscoveredResourceRequest) {
    request = &DescribeDiscoveredResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "DescribeDiscoveredResource")
    
    
    return
}

func NewDescribeDiscoveredResourceResponse() (response *DescribeDiscoveredResourceResponse) {
    response = &DescribeDiscoveredResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDiscoveredResource
// 资源详情
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_RESOURCENOTEXIST = "ResourceNotFound.ResourceNotExist"
func (c *Client) DescribeDiscoveredResource(request *DescribeDiscoveredResourceRequest) (response *DescribeDiscoveredResourceResponse, err error) {
    return c.DescribeDiscoveredResourceWithContext(context.Background(), request)
}

// DescribeDiscoveredResource
// 资源详情
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_RESOURCENOTEXIST = "ResourceNotFound.ResourceNotExist"
func (c *Client) DescribeDiscoveredResourceWithContext(ctx context.Context, request *DescribeDiscoveredResourceRequest) (response *DescribeDiscoveredResourceResponse, err error) {
    if request == nil {
        request = NewDescribeDiscoveredResourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "DescribeDiscoveredResource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDiscoveredResource require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDiscoveredResourceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSystemCompliancePackRequest() (request *DescribeSystemCompliancePackRequest) {
    request = &DescribeSystemCompliancePackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "DescribeSystemCompliancePack")
    
    
    return
}

func NewDescribeSystemCompliancePackResponse() (response *DescribeSystemCompliancePackResponse) {
    response = &DescribeSystemCompliancePackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSystemCompliancePack
// 获取系统合规包详情
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_COMPLIANCEPACKNOTEXIST = "ResourceNotFound.CompliancePackNotExist"
func (c *Client) DescribeSystemCompliancePack(request *DescribeSystemCompliancePackRequest) (response *DescribeSystemCompliancePackResponse, err error) {
    return c.DescribeSystemCompliancePackWithContext(context.Background(), request)
}

// DescribeSystemCompliancePack
// 获取系统合规包详情
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_COMPLIANCEPACKNOTEXIST = "ResourceNotFound.CompliancePackNotExist"
func (c *Client) DescribeSystemCompliancePackWithContext(ctx context.Context, request *DescribeSystemCompliancePackRequest) (response *DescribeSystemCompliancePackResponse, err error) {
    if request == nil {
        request = NewDescribeSystemCompliancePackRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "DescribeSystemCompliancePack")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSystemCompliancePack require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSystemCompliancePackResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSystemRuleRequest() (request *DescribeSystemRuleRequest) {
    request = &DescribeSystemRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "DescribeSystemRule")
    
    
    return
}

func NewDescribeSystemRuleResponse() (response *DescribeSystemRuleResponse) {
    response = &DescribeSystemRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSystemRule
// 预设规则详情
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSystemRule(request *DescribeSystemRuleRequest) (response *DescribeSystemRuleResponse, err error) {
    return c.DescribeSystemRuleWithContext(context.Background(), request)
}

// DescribeSystemRule
// 预设规则详情
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSystemRuleWithContext(ctx context.Context, request *DescribeSystemRuleRequest) (response *DescribeSystemRuleResponse, err error) {
    if request == nil {
        request = NewDescribeSystemRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "DescribeSystemRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSystemRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSystemRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDetachAggregateConfigRuleToCompliancePackRequest() (request *DetachAggregateConfigRuleToCompliancePackRequest) {
    request = &DetachAggregateConfigRuleToCompliancePackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "DetachAggregateConfigRuleToCompliancePack")
    
    
    return
}

func NewDetachAggregateConfigRuleToCompliancePackResponse() (response *DetachAggregateConfigRuleToCompliancePackResponse) {
    response = &DetachAggregateConfigRuleToCompliancePackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DetachAggregateConfigRuleToCompliancePack
// 账号组合规包移除规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPLIANCEPACKISNOTEXIST = "FailedOperation.CompliancePackIsNotExist"
//  FAILEDOPERATION_NOTALLOWOPERATECONTROLCENTERGOVERNANCEPACKAGE = "FailedOperation.NotAllowOperateControlCenterGovernancePackage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_COMPLIANCEPACKISNOTEXIST = "ResourceNotFound.CompliancePackIsNotExist"
//  RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"
func (c *Client) DetachAggregateConfigRuleToCompliancePack(request *DetachAggregateConfigRuleToCompliancePackRequest) (response *DetachAggregateConfigRuleToCompliancePackResponse, err error) {
    return c.DetachAggregateConfigRuleToCompliancePackWithContext(context.Background(), request)
}

// DetachAggregateConfigRuleToCompliancePack
// 账号组合规包移除规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPLIANCEPACKISNOTEXIST = "FailedOperation.CompliancePackIsNotExist"
//  FAILEDOPERATION_NOTALLOWOPERATECONTROLCENTERGOVERNANCEPACKAGE = "FailedOperation.NotAllowOperateControlCenterGovernancePackage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_COMPLIANCEPACKISNOTEXIST = "ResourceNotFound.CompliancePackIsNotExist"
//  RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"
func (c *Client) DetachAggregateConfigRuleToCompliancePackWithContext(ctx context.Context, request *DetachAggregateConfigRuleToCompliancePackRequest) (response *DetachAggregateConfigRuleToCompliancePackResponse, err error) {
    if request == nil {
        request = NewDetachAggregateConfigRuleToCompliancePackRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "DetachAggregateConfigRuleToCompliancePack")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetachAggregateConfigRuleToCompliancePack require credential")
    }

    request.SetContext(ctx)
    
    response = NewDetachAggregateConfigRuleToCompliancePackResponse()
    err = c.Send(request, response)
    return
}

func NewDetachConfigRuleToCompliancePackRequest() (request *DetachConfigRuleToCompliancePackRequest) {
    request = &DetachConfigRuleToCompliancePackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "DetachConfigRuleToCompliancePack")
    
    
    return
}

func NewDetachConfigRuleToCompliancePackResponse() (response *DetachConfigRuleToCompliancePackResponse) {
    response = &DetachConfigRuleToCompliancePackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DetachConfigRuleToCompliancePack
// 合规包移除规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPLIANCEPACKISNOTEXIST = "FailedOperation.CompliancePackIsNotExist"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_COMPLIANCEPACKISNOTEXIST = "ResourceNotFound.CompliancePackIsNotExist"
//  RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"
func (c *Client) DetachConfigRuleToCompliancePack(request *DetachConfigRuleToCompliancePackRequest) (response *DetachConfigRuleToCompliancePackResponse, err error) {
    return c.DetachConfigRuleToCompliancePackWithContext(context.Background(), request)
}

// DetachConfigRuleToCompliancePack
// 合规包移除规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPLIANCEPACKISNOTEXIST = "FailedOperation.CompliancePackIsNotExist"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_COMPLIANCEPACKISNOTEXIST = "ResourceNotFound.CompliancePackIsNotExist"
//  RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"
func (c *Client) DetachConfigRuleToCompliancePackWithContext(ctx context.Context, request *DetachConfigRuleToCompliancePackRequest) (response *DetachConfigRuleToCompliancePackResponse, err error) {
    if request == nil {
        request = NewDetachConfigRuleToCompliancePackRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "DetachConfigRuleToCompliancePack")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetachConfigRuleToCompliancePack require credential")
    }

    request.SetContext(ctx)
    
    response = NewDetachConfigRuleToCompliancePackResponse()
    err = c.Send(request, response)
    return
}

func NewListAggregateCompliancePacksRequest() (request *ListAggregateCompliancePacksRequest) {
    request = &ListAggregateCompliancePacksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "ListAggregateCompliancePacks")
    
    
    return
}

func NewListAggregateCompliancePacksResponse() (response *ListAggregateCompliancePacksResponse) {
    response = &ListAggregateCompliancePacksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListAggregateCompliancePacks
// 账号组获取合规包列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ACCOUNTGROUPISNOTEXIST = "ResourceNotFound.AccountGroupIsNotExist"
func (c *Client) ListAggregateCompliancePacks(request *ListAggregateCompliancePacksRequest) (response *ListAggregateCompliancePacksResponse, err error) {
    return c.ListAggregateCompliancePacksWithContext(context.Background(), request)
}

// ListAggregateCompliancePacks
// 账号组获取合规包列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ACCOUNTGROUPISNOTEXIST = "ResourceNotFound.AccountGroupIsNotExist"
func (c *Client) ListAggregateCompliancePacksWithContext(ctx context.Context, request *ListAggregateCompliancePacksRequest) (response *ListAggregateCompliancePacksResponse, err error) {
    if request == nil {
        request = NewListAggregateCompliancePacksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "ListAggregateCompliancePacks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAggregateCompliancePacks require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAggregateCompliancePacksResponse()
    err = c.Send(request, response)
    return
}

func NewListAggregateConfigRuleEvaluationResultsRequest() (request *ListAggregateConfigRuleEvaluationResultsRequest) {
    request = &ListAggregateConfigRuleEvaluationResultsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "ListAggregateConfigRuleEvaluationResults")
    
    
    return
}

func NewListAggregateConfigRuleEvaluationResultsResponse() (response *ListAggregateConfigRuleEvaluationResultsResponse) {
    response = &ListAggregateConfigRuleEvaluationResultsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListAggregateConfigRuleEvaluationResults
// 账号组获取评估结果--规则维度（某个规则下资源的评估结果列表）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ACCOUNTGROUPISNOTEXIST = "ResourceNotFound.AccountGroupIsNotExist"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
func (c *Client) ListAggregateConfigRuleEvaluationResults(request *ListAggregateConfigRuleEvaluationResultsRequest) (response *ListAggregateConfigRuleEvaluationResultsResponse, err error) {
    return c.ListAggregateConfigRuleEvaluationResultsWithContext(context.Background(), request)
}

// ListAggregateConfigRuleEvaluationResults
// 账号组获取评估结果--规则维度（某个规则下资源的评估结果列表）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ACCOUNTGROUPISNOTEXIST = "ResourceNotFound.AccountGroupIsNotExist"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
func (c *Client) ListAggregateConfigRuleEvaluationResultsWithContext(ctx context.Context, request *ListAggregateConfigRuleEvaluationResultsRequest) (response *ListAggregateConfigRuleEvaluationResultsResponse, err error) {
    if request == nil {
        request = NewListAggregateConfigRuleEvaluationResultsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "ListAggregateConfigRuleEvaluationResults")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAggregateConfigRuleEvaluationResults require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAggregateConfigRuleEvaluationResultsResponse()
    err = c.Send(request, response)
    return
}

func NewListAggregateConfigRulesRequest() (request *ListAggregateConfigRulesRequest) {
    request = &ListAggregateConfigRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "ListAggregateConfigRules")
    
    
    return
}

func NewListAggregateConfigRulesResponse() (response *ListAggregateConfigRulesResponse) {
    response = &ListAggregateConfigRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListAggregateConfigRules
// 账号组获取规则列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ACCOUNTGROUPISNOTEXIST = "ResourceNotFound.AccountGroupIsNotExist"
func (c *Client) ListAggregateConfigRules(request *ListAggregateConfigRulesRequest) (response *ListAggregateConfigRulesResponse, err error) {
    return c.ListAggregateConfigRulesWithContext(context.Background(), request)
}

// ListAggregateConfigRules
// 账号组获取规则列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ACCOUNTGROUPISNOTEXIST = "ResourceNotFound.AccountGroupIsNotExist"
func (c *Client) ListAggregateConfigRulesWithContext(ctx context.Context, request *ListAggregateConfigRulesRequest) (response *ListAggregateConfigRulesResponse, err error) {
    if request == nil {
        request = NewListAggregateConfigRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "ListAggregateConfigRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAggregateConfigRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAggregateConfigRulesResponse()
    err = c.Send(request, response)
    return
}

func NewListAggregateDiscoveredResourcesRequest() (request *ListAggregateDiscoveredResourcesRequest) {
    request = &ListAggregateDiscoveredResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "ListAggregateDiscoveredResources")
    
    
    return
}

func NewListAggregateDiscoveredResourcesResponse() (response *ListAggregateDiscoveredResourcesResponse) {
    response = &ListAggregateDiscoveredResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListAggregateDiscoveredResources
// 账号组获取资源列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ACCOUNTGROUPISNOTEXIST = "ResourceNotFound.AccountGroupIsNotExist"
func (c *Client) ListAggregateDiscoveredResources(request *ListAggregateDiscoveredResourcesRequest) (response *ListAggregateDiscoveredResourcesResponse, err error) {
    return c.ListAggregateDiscoveredResourcesWithContext(context.Background(), request)
}

// ListAggregateDiscoveredResources
// 账号组获取资源列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ACCOUNTGROUPISNOTEXIST = "ResourceNotFound.AccountGroupIsNotExist"
func (c *Client) ListAggregateDiscoveredResourcesWithContext(ctx context.Context, request *ListAggregateDiscoveredResourcesRequest) (response *ListAggregateDiscoveredResourcesResponse, err error) {
    if request == nil {
        request = NewListAggregateDiscoveredResourcesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "ListAggregateDiscoveredResources")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAggregateDiscoveredResources require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAggregateDiscoveredResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewListAggregatorsRequest() (request *ListAggregatorsRequest) {
    request = &ListAggregatorsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "ListAggregators")
    
    
    return
}

func NewListAggregatorsResponse() (response *ListAggregatorsResponse) {
    response = &ListAggregatorsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListAggregators
// 账号组列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListAggregators(request *ListAggregatorsRequest) (response *ListAggregatorsResponse, err error) {
    return c.ListAggregatorsWithContext(context.Background(), request)
}

// ListAggregators
// 账号组列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListAggregatorsWithContext(ctx context.Context, request *ListAggregatorsRequest) (response *ListAggregatorsResponse, err error) {
    if request == nil {
        request = NewListAggregatorsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "ListAggregators")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAggregators require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAggregatorsResponse()
    err = c.Send(request, response)
    return
}

func NewListCompliancePacksRequest() (request *ListCompliancePacksRequest) {
    request = &ListCompliancePacksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "ListCompliancePacks")
    
    
    return
}

func NewListCompliancePacksResponse() (response *ListCompliancePacksResponse) {
    response = &ListCompliancePacksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListCompliancePacks
// 获取合规包列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListCompliancePacks(request *ListCompliancePacksRequest) (response *ListCompliancePacksResponse, err error) {
    return c.ListCompliancePacksWithContext(context.Background(), request)
}

// ListCompliancePacks
// 获取合规包列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListCompliancePacksWithContext(ctx context.Context, request *ListCompliancePacksRequest) (response *ListCompliancePacksResponse, err error) {
    if request == nil {
        request = NewListCompliancePacksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "ListCompliancePacks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListCompliancePacks require credential")
    }

    request.SetContext(ctx)
    
    response = NewListCompliancePacksResponse()
    err = c.Send(request, response)
    return
}

func NewListConfigRuleEvaluationResultsRequest() (request *ListConfigRuleEvaluationResultsRequest) {
    request = &ListConfigRuleEvaluationResultsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "ListConfigRuleEvaluationResults")
    
    
    return
}

func NewListConfigRuleEvaluationResultsResponse() (response *ListConfigRuleEvaluationResultsResponse) {
    response = &ListConfigRuleEvaluationResultsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListConfigRuleEvaluationResults
//  获取评估结果--规则维度（某个规则下资源的评估结果列表）
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"
func (c *Client) ListConfigRuleEvaluationResults(request *ListConfigRuleEvaluationResultsRequest) (response *ListConfigRuleEvaluationResultsResponse, err error) {
    return c.ListConfigRuleEvaluationResultsWithContext(context.Background(), request)
}

// ListConfigRuleEvaluationResults
//  获取评估结果--规则维度（某个规则下资源的评估结果列表）
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"
func (c *Client) ListConfigRuleEvaluationResultsWithContext(ctx context.Context, request *ListConfigRuleEvaluationResultsRequest) (response *ListConfigRuleEvaluationResultsResponse, err error) {
    if request == nil {
        request = NewListConfigRuleEvaluationResultsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "ListConfigRuleEvaluationResults")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListConfigRuleEvaluationResults require credential")
    }

    request.SetContext(ctx)
    
    response = NewListConfigRuleEvaluationResultsResponse()
    err = c.Send(request, response)
    return
}

func NewListConfigRulesRequest() (request *ListConfigRulesRequest) {
    request = &ListConfigRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "ListConfigRules")
    
    
    return
}

func NewListConfigRulesResponse() (response *ListConfigRulesResponse) {
    response = &ListConfigRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListConfigRules
// 获取规则列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListConfigRules(request *ListConfigRulesRequest) (response *ListConfigRulesResponse, err error) {
    return c.ListConfigRulesWithContext(context.Background(), request)
}

// ListConfigRules
// 获取规则列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListConfigRulesWithContext(ctx context.Context, request *ListConfigRulesRequest) (response *ListConfigRulesResponse, err error) {
    if request == nil {
        request = NewListConfigRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "ListConfigRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListConfigRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewListConfigRulesResponse()
    err = c.Send(request, response)
    return
}

func NewListDiscoveredResourcesRequest() (request *ListDiscoveredResourcesRequest) {
    request = &ListDiscoveredResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "ListDiscoveredResources")
    
    
    return
}

func NewListDiscoveredResourcesResponse() (response *ListDiscoveredResourcesResponse) {
    response = &ListDiscoveredResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListDiscoveredResources
// 获取资源列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListDiscoveredResources(request *ListDiscoveredResourcesRequest) (response *ListDiscoveredResourcesResponse, err error) {
    return c.ListDiscoveredResourcesWithContext(context.Background(), request)
}

// ListDiscoveredResources
// 获取资源列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListDiscoveredResourcesWithContext(ctx context.Context, request *ListDiscoveredResourcesRequest) (response *ListDiscoveredResourcesResponse, err error) {
    if request == nil {
        request = NewListDiscoveredResourcesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "ListDiscoveredResources")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListDiscoveredResources require credential")
    }

    request.SetContext(ctx)
    
    response = NewListDiscoveredResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewListRemediationExecutionsRequest() (request *ListRemediationExecutionsRequest) {
    request = &ListRemediationExecutionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "ListRemediationExecutions")
    
    
    return
}

func NewListRemediationExecutionsResponse() (response *ListRemediationExecutionsResponse) {
    response = &ListRemediationExecutionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListRemediationExecutions
// 修正记录
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_REMEDIATIONISNOTEXIST = "ResourceNotFound.RemediationIsNotExist"
func (c *Client) ListRemediationExecutions(request *ListRemediationExecutionsRequest) (response *ListRemediationExecutionsResponse, err error) {
    return c.ListRemediationExecutionsWithContext(context.Background(), request)
}

// ListRemediationExecutions
// 修正记录
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_REMEDIATIONISNOTEXIST = "ResourceNotFound.RemediationIsNotExist"
func (c *Client) ListRemediationExecutionsWithContext(ctx context.Context, request *ListRemediationExecutionsRequest) (response *ListRemediationExecutionsResponse, err error) {
    if request == nil {
        request = NewListRemediationExecutionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "ListRemediationExecutions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListRemediationExecutions require credential")
    }

    request.SetContext(ctx)
    
    response = NewListRemediationExecutionsResponse()
    err = c.Send(request, response)
    return
}

func NewListRemediationsRequest() (request *ListRemediationsRequest) {
    request = &ListRemediationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "ListRemediations")
    
    
    return
}

func NewListRemediationsResponse() (response *ListRemediationsResponse) {
    response = &ListRemediationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListRemediations
// 修正设置列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListRemediations(request *ListRemediationsRequest) (response *ListRemediationsResponse, err error) {
    return c.ListRemediationsWithContext(context.Background(), request)
}

// ListRemediations
// 修正设置列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListRemediationsWithContext(ctx context.Context, request *ListRemediationsRequest) (response *ListRemediationsResponse, err error) {
    if request == nil {
        request = NewListRemediationsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "ListRemediations")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListRemediations require credential")
    }

    request.SetContext(ctx)
    
    response = NewListRemediationsResponse()
    err = c.Send(request, response)
    return
}

func NewListResourceTypesRequest() (request *ListResourceTypesRequest) {
    request = &ListResourceTypesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "ListResourceTypes")
    
    
    return
}

func NewListResourceTypesResponse() (response *ListResourceTypesResponse) {
    response = &ListResourceTypesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListResourceTypes
// 获取资源类型列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListResourceTypes(request *ListResourceTypesRequest) (response *ListResourceTypesResponse, err error) {
    return c.ListResourceTypesWithContext(context.Background(), request)
}

// ListResourceTypes
// 获取资源类型列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListResourceTypesWithContext(ctx context.Context, request *ListResourceTypesRequest) (response *ListResourceTypesResponse, err error) {
    if request == nil {
        request = NewListResourceTypesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "ListResourceTypes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListResourceTypes require credential")
    }

    request.SetContext(ctx)
    
    response = NewListResourceTypesResponse()
    err = c.Send(request, response)
    return
}

func NewListSystemCompliancePacksRequest() (request *ListSystemCompliancePacksRequest) {
    request = &ListSystemCompliancePacksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "ListSystemCompliancePacks")
    
    
    return
}

func NewListSystemCompliancePacksResponse() (response *ListSystemCompliancePacksResponse) {
    response = &ListSystemCompliancePacksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListSystemCompliancePacks
// 获取系统合规包列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListSystemCompliancePacks(request *ListSystemCompliancePacksRequest) (response *ListSystemCompliancePacksResponse, err error) {
    return c.ListSystemCompliancePacksWithContext(context.Background(), request)
}

// ListSystemCompliancePacks
// 获取系统合规包列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListSystemCompliancePacksWithContext(ctx context.Context, request *ListSystemCompliancePacksRequest) (response *ListSystemCompliancePacksResponse, err error) {
    if request == nil {
        request = NewListSystemCompliancePacksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "ListSystemCompliancePacks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListSystemCompliancePacks require credential")
    }

    request.SetContext(ctx)
    
    response = NewListSystemCompliancePacksResponse()
    err = c.Send(request, response)
    return
}

func NewListSystemRulesRequest() (request *ListSystemRulesRequest) {
    request = &ListSystemRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "ListSystemRules")
    
    
    return
}

func NewListSystemRulesResponse() (response *ListSystemRulesResponse) {
    response = &ListSystemRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListSystemRules
// 获取预设规则列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListSystemRules(request *ListSystemRulesRequest) (response *ListSystemRulesResponse, err error) {
    return c.ListSystemRulesWithContext(context.Background(), request)
}

// ListSystemRules
// 获取预设规则列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListSystemRulesWithContext(ctx context.Context, request *ListSystemRulesRequest) (response *ListSystemRulesResponse, err error) {
    if request == nil {
        request = NewListSystemRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "ListSystemRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListSystemRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewListSystemRulesResponse()
    err = c.Send(request, response)
    return
}

func NewOpenAggregateConfigRuleRequest() (request *OpenAggregateConfigRuleRequest) {
    request = &OpenAggregateConfigRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "OpenAggregateConfigRule")
    
    
    return
}

func NewOpenAggregateConfigRuleResponse() (response *OpenAggregateConfigRuleResponse) {
    response = &OpenAggregateConfigRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OpenAggregateConfigRule
// 账号组开启规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTALLOWOPERATECONTROLCENTERGOVERNANCEPACKAGERULE = "FailedOperation.NotAllowOperateControlCenterGovernancePackageRule"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) OpenAggregateConfigRule(request *OpenAggregateConfigRuleRequest) (response *OpenAggregateConfigRuleResponse, err error) {
    return c.OpenAggregateConfigRuleWithContext(context.Background(), request)
}

// OpenAggregateConfigRule
// 账号组开启规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTALLOWOPERATECONTROLCENTERGOVERNANCEPACKAGERULE = "FailedOperation.NotAllowOperateControlCenterGovernancePackageRule"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) OpenAggregateConfigRuleWithContext(ctx context.Context, request *OpenAggregateConfigRuleRequest) (response *OpenAggregateConfigRuleResponse, err error) {
    if request == nil {
        request = NewOpenAggregateConfigRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "OpenAggregateConfigRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenAggregateConfigRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenAggregateConfigRuleResponse()
    err = c.Send(request, response)
    return
}

func NewOpenConfigRecorderRequest() (request *OpenConfigRecorderRequest) {
    request = &OpenConfigRecorderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "OpenConfigRecorder")
    
    
    return
}

func NewOpenConfigRecorderResponse() (response *OpenConfigRecorderResponse) {
    response = &OpenConfigRecorderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OpenConfigRecorder
// 资源监控管理-开启监控
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_OPENRECORDERLIMITOVER = "LimitExceeded.OpenRecorderLimitOver"
func (c *Client) OpenConfigRecorder(request *OpenConfigRecorderRequest) (response *OpenConfigRecorderResponse, err error) {
    return c.OpenConfigRecorderWithContext(context.Background(), request)
}

// OpenConfigRecorder
// 资源监控管理-开启监控
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_OPENRECORDERLIMITOVER = "LimitExceeded.OpenRecorderLimitOver"
func (c *Client) OpenConfigRecorderWithContext(ctx context.Context, request *OpenConfigRecorderRequest) (response *OpenConfigRecorderResponse, err error) {
    if request == nil {
        request = NewOpenConfigRecorderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "OpenConfigRecorder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenConfigRecorder require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenConfigRecorderResponse()
    err = c.Send(request, response)
    return
}

func NewOpenConfigRuleRequest() (request *OpenConfigRuleRequest) {
    request = &OpenConfigRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "OpenConfigRule")
    
    
    return
}

func NewOpenConfigRuleResponse() (response *OpenConfigRuleResponse) {
    response = &OpenConfigRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OpenConfigRule
// 开启规则
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"
func (c *Client) OpenConfigRule(request *OpenConfigRuleRequest) (response *OpenConfigRuleResponse, err error) {
    return c.OpenConfigRuleWithContext(context.Background(), request)
}

// OpenConfigRule
// 开启规则
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"
func (c *Client) OpenConfigRuleWithContext(ctx context.Context, request *OpenConfigRuleRequest) (response *OpenConfigRuleResponse, err error) {
    if request == nil {
        request = NewOpenConfigRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "OpenConfigRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenConfigRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenConfigRuleResponse()
    err = c.Send(request, response)
    return
}

func NewPutEvaluationsRequest() (request *PutEvaluationsRequest) {
    request = &PutEvaluationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "PutEvaluations")
    
    
    return
}

func NewPutEvaluationsResponse() (response *PutEvaluationsResponse) {
    response = &PutEvaluationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PutEvaluations
// 上报自定义规则评估结果
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ACCOUNTGROUPISNOTEXIST = "ResourceNotFound.AccountGroupIsNotExist"
//  RESOURCENOTFOUND_RESOURCENOTEXIST = "ResourceNotFound.ResourceNotExist"
//  RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"
func (c *Client) PutEvaluations(request *PutEvaluationsRequest) (response *PutEvaluationsResponse, err error) {
    return c.PutEvaluationsWithContext(context.Background(), request)
}

// PutEvaluations
// 上报自定义规则评估结果
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ACCOUNTGROUPISNOTEXIST = "ResourceNotFound.AccountGroupIsNotExist"
//  RESOURCENOTFOUND_RESOURCENOTEXIST = "ResourceNotFound.ResourceNotExist"
//  RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"
func (c *Client) PutEvaluationsWithContext(ctx context.Context, request *PutEvaluationsRequest) (response *PutEvaluationsResponse, err error) {
    if request == nil {
        request = NewPutEvaluationsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "PutEvaluations")
    
    if c.GetCredential() == nil {
        return nil, errors.New("PutEvaluations require credential")
    }

    request.SetContext(ctx)
    
    response = NewPutEvaluationsResponse()
    err = c.Send(request, response)
    return
}

func NewStartAggregateConfigRuleEvaluationRequest() (request *StartAggregateConfigRuleEvaluationRequest) {
    request = &StartAggregateConfigRuleEvaluationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "StartAggregateConfigRuleEvaluation")
    
    
    return
}

func NewStartAggregateConfigRuleEvaluationResponse() (response *StartAggregateConfigRuleEvaluationResponse) {
    response = &StartAggregateConfigRuleEvaluationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartAggregateConfigRuleEvaluation
// 账号组触发评估
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ACCOUNTGROUPISNOTEXIST = "ResourceNotFound.AccountGroupIsNotExist"
//  RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) StartAggregateConfigRuleEvaluation(request *StartAggregateConfigRuleEvaluationRequest) (response *StartAggregateConfigRuleEvaluationResponse, err error) {
    return c.StartAggregateConfigRuleEvaluationWithContext(context.Background(), request)
}

// StartAggregateConfigRuleEvaluation
// 账号组触发评估
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ACCOUNTGROUPISNOTEXIST = "ResourceNotFound.AccountGroupIsNotExist"
//  RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) StartAggregateConfigRuleEvaluationWithContext(ctx context.Context, request *StartAggregateConfigRuleEvaluationRequest) (response *StartAggregateConfigRuleEvaluationResponse, err error) {
    if request == nil {
        request = NewStartAggregateConfigRuleEvaluationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "StartAggregateConfigRuleEvaluation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartAggregateConfigRuleEvaluation require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartAggregateConfigRuleEvaluationResponse()
    err = c.Send(request, response)
    return
}

func NewStartConfigRuleEvaluationRequest() (request *StartConfigRuleEvaluationRequest) {
    request = &StartConfigRuleEvaluationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "StartConfigRuleEvaluation")
    
    
    return
}

func NewStartConfigRuleEvaluationResponse() (response *StartConfigRuleEvaluationResponse) {
    response = &StartConfigRuleEvaluationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartConfigRuleEvaluation
// 触发评估
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"
func (c *Client) StartConfigRuleEvaluation(request *StartConfigRuleEvaluationRequest) (response *StartConfigRuleEvaluationResponse, err error) {
    return c.StartConfigRuleEvaluationWithContext(context.Background(), request)
}

// StartConfigRuleEvaluation
// 触发评估
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"
func (c *Client) StartConfigRuleEvaluationWithContext(ctx context.Context, request *StartConfigRuleEvaluationRequest) (response *StartConfigRuleEvaluationResponse, err error) {
    if request == nil {
        request = NewStartConfigRuleEvaluationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "StartConfigRuleEvaluation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartConfigRuleEvaluation require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartConfigRuleEvaluationResponse()
    err = c.Send(request, response)
    return
}

func NewStartRemediationRequest() (request *StartRemediationRequest) {
    request = &StartRemediationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "StartRemediation")
    
    
    return
}

func NewStartRemediationResponse() (response *StartRemediationResponse) {
    response = &StartRemediationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartRemediation
// 手动执行规则修复
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) StartRemediation(request *StartRemediationRequest) (response *StartRemediationResponse, err error) {
    return c.StartRemediationWithContext(context.Background(), request)
}

// StartRemediation
// 手动执行规则修复
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) StartRemediationWithContext(ctx context.Context, request *StartRemediationRequest) (response *StartRemediationResponse, err error) {
    if request == nil {
        request = NewStartRemediationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "StartRemediation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartRemediation require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartRemediationResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAggregateCompliancePackRequest() (request *UpdateAggregateCompliancePackRequest) {
    request = &UpdateAggregateCompliancePackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "UpdateAggregateCompliancePack")
    
    
    return
}

func NewUpdateAggregateCompliancePackResponse() (response *UpdateAggregateCompliancePackResponse) {
    response = &UpdateAggregateCompliancePackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateAggregateCompliancePack
// 账号组编辑合规包
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPLIANCEPACKISNOTEXIST = "FailedOperation.CompliancePackIsNotExist"
//  FAILEDOPERATION_NOTALLOWOPERATECONTROLCENTERGOVERNANCEPACKAGE = "FailedOperation.NotAllowOperateControlCenterGovernancePackage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_COMPLIANCEPACKISNOTEXIST = "ResourceNotFound.CompliancePackIsNotExist"
//  RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"
func (c *Client) UpdateAggregateCompliancePack(request *UpdateAggregateCompliancePackRequest) (response *UpdateAggregateCompliancePackResponse, err error) {
    return c.UpdateAggregateCompliancePackWithContext(context.Background(), request)
}

// UpdateAggregateCompliancePack
// 账号组编辑合规包
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPLIANCEPACKISNOTEXIST = "FailedOperation.CompliancePackIsNotExist"
//  FAILEDOPERATION_NOTALLOWOPERATECONTROLCENTERGOVERNANCEPACKAGE = "FailedOperation.NotAllowOperateControlCenterGovernancePackage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_COMPLIANCEPACKISNOTEXIST = "ResourceNotFound.CompliancePackIsNotExist"
//  RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"
func (c *Client) UpdateAggregateCompliancePackWithContext(ctx context.Context, request *UpdateAggregateCompliancePackRequest) (response *UpdateAggregateCompliancePackResponse, err error) {
    if request == nil {
        request = NewUpdateAggregateCompliancePackRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "UpdateAggregateCompliancePack")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateAggregateCompliancePack require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateAggregateCompliancePackResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAggregateCompliancePackStatusRequest() (request *UpdateAggregateCompliancePackStatusRequest) {
    request = &UpdateAggregateCompliancePackStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "UpdateAggregateCompliancePackStatus")
    
    
    return
}

func NewUpdateAggregateCompliancePackStatusResponse() (response *UpdateAggregateCompliancePackStatusResponse) {
    response = &UpdateAggregateCompliancePackStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateAggregateCompliancePackStatus
// 账号组开启、关闭合规包
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPLIANCEPACKISNOTEXIST = "FailedOperation.CompliancePackIsNotExist"
//  FAILEDOPERATION_NOTALLOWOPERATECONTROLCENTERGOVERNANCEPACKAGE = "FailedOperation.NotAllowOperateControlCenterGovernancePackage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_COMPLIANCEPACKISNOTEXIST = "ResourceNotFound.CompliancePackIsNotExist"
func (c *Client) UpdateAggregateCompliancePackStatus(request *UpdateAggregateCompliancePackStatusRequest) (response *UpdateAggregateCompliancePackStatusResponse, err error) {
    return c.UpdateAggregateCompliancePackStatusWithContext(context.Background(), request)
}

// UpdateAggregateCompliancePackStatus
// 账号组开启、关闭合规包
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPLIANCEPACKISNOTEXIST = "FailedOperation.CompliancePackIsNotExist"
//  FAILEDOPERATION_NOTALLOWOPERATECONTROLCENTERGOVERNANCEPACKAGE = "FailedOperation.NotAllowOperateControlCenterGovernancePackage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_COMPLIANCEPACKISNOTEXIST = "ResourceNotFound.CompliancePackIsNotExist"
func (c *Client) UpdateAggregateCompliancePackStatusWithContext(ctx context.Context, request *UpdateAggregateCompliancePackStatusRequest) (response *UpdateAggregateCompliancePackStatusResponse, err error) {
    if request == nil {
        request = NewUpdateAggregateCompliancePackStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "UpdateAggregateCompliancePackStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateAggregateCompliancePackStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateAggregateCompliancePackStatusResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAggregateConfigDeliverRequest() (request *UpdateAggregateConfigDeliverRequest) {
    request = &UpdateAggregateConfigDeliverRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "UpdateAggregateConfigDeliver")
    
    
    return
}

func NewUpdateAggregateConfigDeliverResponse() (response *UpdateAggregateConfigDeliverResponse) {
    response = &UpdateAggregateConfigDeliverResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateAggregateConfigDeliver
// 账号组编辑投递设置
//
// 可能返回的错误码:
//  FAILEDOPERATION_MEMBERNOTASSIGNMANAGER = "FailedOperation.MemberNotAssignManager"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_DELIVERBUCKETNOTEXIST = "ResourceNotFound.DeliverBucketNotExist"
//  RESOURCENOTFOUND_DELIVEREVENTNOTEXIST = "ResourceNotFound.DeliverEventNotExist"
func (c *Client) UpdateAggregateConfigDeliver(request *UpdateAggregateConfigDeliverRequest) (response *UpdateAggregateConfigDeliverResponse, err error) {
    return c.UpdateAggregateConfigDeliverWithContext(context.Background(), request)
}

// UpdateAggregateConfigDeliver
// 账号组编辑投递设置
//
// 可能返回的错误码:
//  FAILEDOPERATION_MEMBERNOTASSIGNMANAGER = "FailedOperation.MemberNotAssignManager"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_DELIVERBUCKETNOTEXIST = "ResourceNotFound.DeliverBucketNotExist"
//  RESOURCENOTFOUND_DELIVEREVENTNOTEXIST = "ResourceNotFound.DeliverEventNotExist"
func (c *Client) UpdateAggregateConfigDeliverWithContext(ctx context.Context, request *UpdateAggregateConfigDeliverRequest) (response *UpdateAggregateConfigDeliverResponse, err error) {
    if request == nil {
        request = NewUpdateAggregateConfigDeliverRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "UpdateAggregateConfigDeliver")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateAggregateConfigDeliver require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateAggregateConfigDeliverResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAggregateConfigRuleRequest() (request *UpdateAggregateConfigRuleRequest) {
    request = &UpdateAggregateConfigRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "UpdateAggregateConfigRule")
    
    
    return
}

func NewUpdateAggregateConfigRuleResponse() (response *UpdateAggregateConfigRuleResponse) {
    response = &UpdateAggregateConfigRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateAggregateConfigRule
// 账号组编辑规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTALLOWOPERATECONTROLCENTERGOVERNANCEPACKAGERULE = "FailedOperation.NotAllowOperateControlCenterGovernancePackageRule"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ACCOUNTGROUPISNOTEXIST = "ResourceNotFound.AccountGroupIsNotExist"
//  RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) UpdateAggregateConfigRule(request *UpdateAggregateConfigRuleRequest) (response *UpdateAggregateConfigRuleResponse, err error) {
    return c.UpdateAggregateConfigRuleWithContext(context.Background(), request)
}

// UpdateAggregateConfigRule
// 账号组编辑规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTALLOWOPERATECONTROLCENTERGOVERNANCEPACKAGERULE = "FailedOperation.NotAllowOperateControlCenterGovernancePackageRule"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ACCOUNTGROUPISNOTEXIST = "ResourceNotFound.AccountGroupIsNotExist"
//  RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) UpdateAggregateConfigRuleWithContext(ctx context.Context, request *UpdateAggregateConfigRuleRequest) (response *UpdateAggregateConfigRuleResponse, err error) {
    if request == nil {
        request = NewUpdateAggregateConfigRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "UpdateAggregateConfigRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateAggregateConfigRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateAggregateConfigRuleResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateCompliancePackRequest() (request *UpdateCompliancePackRequest) {
    request = &UpdateCompliancePackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "UpdateCompliancePack")
    
    
    return
}

func NewUpdateCompliancePackResponse() (response *UpdateCompliancePackResponse) {
    response = &UpdateCompliancePackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateCompliancePack
// 编辑合规包
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPLIANCEPACKISNOTEXIST = "FailedOperation.CompliancePackIsNotExist"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_COMPLIANCEPACKISNOTEXIST = "ResourceNotFound.CompliancePackIsNotExist"
//  RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"
func (c *Client) UpdateCompliancePack(request *UpdateCompliancePackRequest) (response *UpdateCompliancePackResponse, err error) {
    return c.UpdateCompliancePackWithContext(context.Background(), request)
}

// UpdateCompliancePack
// 编辑合规包
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPLIANCEPACKISNOTEXIST = "FailedOperation.CompliancePackIsNotExist"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_COMPLIANCEPACKISNOTEXIST = "ResourceNotFound.CompliancePackIsNotExist"
//  RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"
func (c *Client) UpdateCompliancePackWithContext(ctx context.Context, request *UpdateCompliancePackRequest) (response *UpdateCompliancePackResponse, err error) {
    if request == nil {
        request = NewUpdateCompliancePackRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "UpdateCompliancePack")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateCompliancePack require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateCompliancePackResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateCompliancePackStatusRequest() (request *UpdateCompliancePackStatusRequest) {
    request = &UpdateCompliancePackStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "UpdateCompliancePackStatus")
    
    
    return
}

func NewUpdateCompliancePackStatusResponse() (response *UpdateCompliancePackStatusResponse) {
    response = &UpdateCompliancePackStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateCompliancePackStatus
// 开启、关闭合规包
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPLIANCEPACKISNOTEXIST = "FailedOperation.CompliancePackIsNotExist"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_COMPLIANCEPACKISNOTEXIST = "ResourceNotFound.CompliancePackIsNotExist"
func (c *Client) UpdateCompliancePackStatus(request *UpdateCompliancePackStatusRequest) (response *UpdateCompliancePackStatusResponse, err error) {
    return c.UpdateCompliancePackStatusWithContext(context.Background(), request)
}

// UpdateCompliancePackStatus
// 开启、关闭合规包
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPLIANCEPACKISNOTEXIST = "FailedOperation.CompliancePackIsNotExist"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_COMPLIANCEPACKISNOTEXIST = "ResourceNotFound.CompliancePackIsNotExist"
func (c *Client) UpdateCompliancePackStatusWithContext(ctx context.Context, request *UpdateCompliancePackStatusRequest) (response *UpdateCompliancePackStatusResponse, err error) {
    if request == nil {
        request = NewUpdateCompliancePackStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "UpdateCompliancePackStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateCompliancePackStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateCompliancePackStatusResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateConfigDeliverRequest() (request *UpdateConfigDeliverRequest) {
    request = &UpdateConfigDeliverRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "UpdateConfigDeliver")
    
    
    return
}

func NewUpdateConfigDeliverResponse() (response *UpdateConfigDeliverResponse) {
    response = &UpdateConfigDeliverResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateConfigDeliver
// 编辑投递设置
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_DELIVEREVENTNOTEXIST = "ResourceNotFound.DeliverEventNotExist"
func (c *Client) UpdateConfigDeliver(request *UpdateConfigDeliverRequest) (response *UpdateConfigDeliverResponse, err error) {
    return c.UpdateConfigDeliverWithContext(context.Background(), request)
}

// UpdateConfigDeliver
// 编辑投递设置
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_DELIVEREVENTNOTEXIST = "ResourceNotFound.DeliverEventNotExist"
func (c *Client) UpdateConfigDeliverWithContext(ctx context.Context, request *UpdateConfigDeliverRequest) (response *UpdateConfigDeliverResponse, err error) {
    if request == nil {
        request = NewUpdateConfigDeliverRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "UpdateConfigDeliver")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateConfigDeliver require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateConfigDeliverResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateConfigRecorderRequest() (request *UpdateConfigRecorderRequest) {
    request = &UpdateConfigRecorderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "UpdateConfigRecorder")
    
    
    return
}

func NewUpdateConfigRecorderResponse() (response *UpdateConfigRecorderResponse) {
    response = &UpdateConfigRecorderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateConfigRecorder
// 编辑监控范围
//
// 可能返回的错误码:
//  FAILEDOPERATION_RECORDERISCLOSE = "FailedOperation.RecorderIsClose"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_UPDATERECORDERLIMITOVER = "LimitExceeded.UpdateRecorderLimitOver"
//  RESOURCENOTFOUND_RECORDERISNOTEXIST = "ResourceNotFound.RecorderIsNotExist"
func (c *Client) UpdateConfigRecorder(request *UpdateConfigRecorderRequest) (response *UpdateConfigRecorderResponse, err error) {
    return c.UpdateConfigRecorderWithContext(context.Background(), request)
}

// UpdateConfigRecorder
// 编辑监控范围
//
// 可能返回的错误码:
//  FAILEDOPERATION_RECORDERISCLOSE = "FailedOperation.RecorderIsClose"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_UPDATERECORDERLIMITOVER = "LimitExceeded.UpdateRecorderLimitOver"
//  RESOURCENOTFOUND_RECORDERISNOTEXIST = "ResourceNotFound.RecorderIsNotExist"
func (c *Client) UpdateConfigRecorderWithContext(ctx context.Context, request *UpdateConfigRecorderRequest) (response *UpdateConfigRecorderResponse, err error) {
    if request == nil {
        request = NewUpdateConfigRecorderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "UpdateConfigRecorder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateConfigRecorder require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateConfigRecorderResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateConfigRuleRequest() (request *UpdateConfigRuleRequest) {
    request = &UpdateConfigRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "UpdateConfigRule")
    
    
    return
}

func NewUpdateConfigRuleResponse() (response *UpdateConfigRuleResponse) {
    response = &UpdateConfigRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateConfigRule
// 编辑规则
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"
func (c *Client) UpdateConfigRule(request *UpdateConfigRuleRequest) (response *UpdateConfigRuleResponse, err error) {
    return c.UpdateConfigRuleWithContext(context.Background(), request)
}

// UpdateConfigRule
// 编辑规则
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"
func (c *Client) UpdateConfigRuleWithContext(ctx context.Context, request *UpdateConfigRuleRequest) (response *UpdateConfigRuleResponse, err error) {
    if request == nil {
        request = NewUpdateConfigRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "UpdateConfigRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateConfigRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateConfigRuleResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateRemediationRequest() (request *UpdateRemediationRequest) {
    request = &UpdateRemediationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "UpdateRemediation")
    
    
    return
}

func NewUpdateRemediationResponse() (response *UpdateRemediationResponse) {
    response = &UpdateRemediationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateRemediation
// 新增告警监控规则
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UpdateRemediation(request *UpdateRemediationRequest) (response *UpdateRemediationResponse, err error) {
    return c.UpdateRemediationWithContext(context.Background(), request)
}

// UpdateRemediation
// 新增告警监控规则
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UpdateRemediationWithContext(ctx context.Context, request *UpdateRemediationRequest) (response *UpdateRemediationResponse, err error) {
    if request == nil {
        request = NewUpdateRemediationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "UpdateRemediation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateRemediation require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateRemediationResponse()
    err = c.Send(request, response)
    return
}
