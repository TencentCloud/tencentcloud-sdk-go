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

package v20180709

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-07-09"

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


func NewCreateAllocationRuleRequest() (request *CreateAllocationRuleRequest) {
    request = &CreateAllocationRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "CreateAllocationRule")
    
    
    return
}

func NewCreateAllocationRuleResponse() (response *CreateAllocationRuleResponse) {
    response = &CreateAllocationRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAllocationRule
// 创建公摊规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_NUMLIMITERROR = "FailedOperation.NumLimitError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATERROR = "InternalError.DbOperatError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateAllocationRule(request *CreateAllocationRuleRequest) (response *CreateAllocationRuleResponse, err error) {
    return c.CreateAllocationRuleWithContext(context.Background(), request)
}

// CreateAllocationRule
// 创建公摊规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_NUMLIMITERROR = "FailedOperation.NumLimitError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATERROR = "InternalError.DbOperatError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateAllocationRuleWithContext(ctx context.Context, request *CreateAllocationRuleRequest) (response *CreateAllocationRuleResponse, err error) {
    if request == nil {
        request = NewCreateAllocationRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "CreateAllocationRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAllocationRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAllocationRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAllocationTagRequest() (request *CreateAllocationTagRequest) {
    request = &CreateAllocationTagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "CreateAllocationTag")
    
    
    return
}

func NewCreateAllocationTagResponse() (response *CreateAllocationTagResponse) {
    response = &CreateAllocationTagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAllocationTag
// 批量设置分账标签
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateAllocationTag(request *CreateAllocationTagRequest) (response *CreateAllocationTagResponse, err error) {
    return c.CreateAllocationTagWithContext(context.Background(), request)
}

// CreateAllocationTag
// 批量设置分账标签
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateAllocationTagWithContext(ctx context.Context, request *CreateAllocationTagRequest) (response *CreateAllocationTagResponse, err error) {
    if request == nil {
        request = NewCreateAllocationTagRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "CreateAllocationTag")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAllocationTag require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAllocationTagResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAllocationUnitRequest() (request *CreateAllocationUnitRequest) {
    request = &CreateAllocationUnitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "CreateAllocationUnit")
    
    
    return
}

func NewCreateAllocationUnitResponse() (response *CreateAllocationUnitResponse) {
    response = &CreateAllocationUnitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAllocationUnit
// 创建分账单元
//
// 可能返回的错误码:
//  FAILEDOPERATION_NUMLIMITERROR = "FailedOperation.NumLimitError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATERROR = "InternalError.DbOperatError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateAllocationUnit(request *CreateAllocationUnitRequest) (response *CreateAllocationUnitResponse, err error) {
    return c.CreateAllocationUnitWithContext(context.Background(), request)
}

// CreateAllocationUnit
// 创建分账单元
//
// 可能返回的错误码:
//  FAILEDOPERATION_NUMLIMITERROR = "FailedOperation.NumLimitError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATERROR = "InternalError.DbOperatError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateAllocationUnitWithContext(ctx context.Context, request *CreateAllocationUnitRequest) (response *CreateAllocationUnitResponse, err error) {
    if request == nil {
        request = NewCreateAllocationUnitRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "CreateAllocationUnit")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAllocationUnit require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAllocationUnitResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBudgetRequest() (request *CreateBudgetRequest) {
    request = &CreateBudgetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "CreateBudget")
    
    
    return
}

func NewCreateBudgetResponse() (response *CreateBudgetResponse) {
    response = &CreateBudgetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateBudget
// 创建预算信息
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) CreateBudget(request *CreateBudgetRequest) (response *CreateBudgetResponse, err error) {
    return c.CreateBudgetWithContext(context.Background(), request)
}

// CreateBudget
// 创建预算信息
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) CreateBudgetWithContext(ctx context.Context, request *CreateBudgetRequest) (response *CreateBudgetResponse, err error) {
    if request == nil {
        request = NewCreateBudgetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "CreateBudget")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBudget require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBudgetResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGatherRuleRequest() (request *CreateGatherRuleRequest) {
    request = &CreateGatherRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "CreateGatherRule")
    
    
    return
}

func NewCreateGatherRuleResponse() (response *CreateGatherRuleResponse) {
    response = &CreateGatherRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateGatherRule
// 创建归集规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_NUMLIMITERROR = "FailedOperation.NumLimitError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATERROR = "InternalError.DbOperatError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateGatherRule(request *CreateGatherRuleRequest) (response *CreateGatherRuleResponse, err error) {
    return c.CreateGatherRuleWithContext(context.Background(), request)
}

// CreateGatherRule
// 创建归集规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_NUMLIMITERROR = "FailedOperation.NumLimitError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATERROR = "InternalError.DbOperatError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateGatherRuleWithContext(ctx context.Context, request *CreateGatherRuleRequest) (response *CreateGatherRuleResponse, err error) {
    if request == nil {
        request = NewCreateGatherRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "CreateGatherRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateGatherRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateGatherRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAllocationRuleRequest() (request *DeleteAllocationRuleRequest) {
    request = &DeleteAllocationRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DeleteAllocationRule")
    
    
    return
}

func NewDeleteAllocationRuleResponse() (response *DeleteAllocationRuleResponse) {
    response = &DeleteAllocationRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAllocationRule
// 公摊规则删除接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATERROR = "InternalError.DbOperatError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DeleteAllocationRule(request *DeleteAllocationRuleRequest) (response *DeleteAllocationRuleResponse, err error) {
    return c.DeleteAllocationRuleWithContext(context.Background(), request)
}

// DeleteAllocationRule
// 公摊规则删除接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATERROR = "InternalError.DbOperatError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DeleteAllocationRuleWithContext(ctx context.Context, request *DeleteAllocationRuleRequest) (response *DeleteAllocationRuleResponse, err error) {
    if request == nil {
        request = NewDeleteAllocationRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DeleteAllocationRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAllocationRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAllocationRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAllocationTagRequest() (request *DeleteAllocationTagRequest) {
    request = &DeleteAllocationTagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DeleteAllocationTag")
    
    
    return
}

func NewDeleteAllocationTagResponse() (response *DeleteAllocationTagResponse) {
    response = &DeleteAllocationTagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAllocationTag
// 批量取消设置分账标签
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteAllocationTag(request *DeleteAllocationTagRequest) (response *DeleteAllocationTagResponse, err error) {
    return c.DeleteAllocationTagWithContext(context.Background(), request)
}

// DeleteAllocationTag
// 批量取消设置分账标签
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteAllocationTagWithContext(ctx context.Context, request *DeleteAllocationTagRequest) (response *DeleteAllocationTagResponse, err error) {
    if request == nil {
        request = NewDeleteAllocationTagRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DeleteAllocationTag")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAllocationTag require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAllocationTagResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAllocationUnitRequest() (request *DeleteAllocationUnitRequest) {
    request = &DeleteAllocationUnitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DeleteAllocationUnit")
    
    
    return
}

func NewDeleteAllocationUnitResponse() (response *DeleteAllocationUnitResponse) {
    response = &DeleteAllocationUnitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAllocationUnit
// 删除分账单元
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DeleteAllocationUnit(request *DeleteAllocationUnitRequest) (response *DeleteAllocationUnitResponse, err error) {
    return c.DeleteAllocationUnitWithContext(context.Background(), request)
}

// DeleteAllocationUnit
// 删除分账单元
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DeleteAllocationUnitWithContext(ctx context.Context, request *DeleteAllocationUnitRequest) (response *DeleteAllocationUnitResponse, err error) {
    if request == nil {
        request = NewDeleteAllocationUnitRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DeleteAllocationUnit")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAllocationUnit require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAllocationUnitResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBudgetRequest() (request *DeleteBudgetRequest) {
    request = &DeleteBudgetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DeleteBudget")
    
    
    return
}

func NewDeleteBudgetResponse() (response *DeleteBudgetResponse) {
    response = &DeleteBudgetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteBudget
// 依据预算ID删除对应预算项目
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DeleteBudget(request *DeleteBudgetRequest) (response *DeleteBudgetResponse, err error) {
    return c.DeleteBudgetWithContext(context.Background(), request)
}

// DeleteBudget
// 依据预算ID删除对应预算项目
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DeleteBudgetWithContext(ctx context.Context, request *DeleteBudgetRequest) (response *DeleteBudgetResponse, err error) {
    if request == nil {
        request = NewDeleteBudgetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DeleteBudget")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteBudget require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteBudgetResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGatherRuleRequest() (request *DeleteGatherRuleRequest) {
    request = &DeleteGatherRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DeleteGatherRule")
    
    
    return
}

func NewDeleteGatherRuleResponse() (response *DeleteGatherRuleResponse) {
    response = &DeleteGatherRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteGatherRule
// 删除归集规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATERROR = "InternalError.DbOperatError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DeleteGatherRule(request *DeleteGatherRuleRequest) (response *DeleteGatherRuleResponse, err error) {
    return c.DeleteGatherRuleWithContext(context.Background(), request)
}

// DeleteGatherRule
// 删除归集规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATERROR = "InternalError.DbOperatError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DeleteGatherRuleWithContext(ctx context.Context, request *DeleteGatherRuleRequest) (response *DeleteGatherRuleResponse, err error) {
    if request == nil {
        request = NewDeleteGatherRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DeleteGatherRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteGatherRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteGatherRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountBalanceRequest() (request *DescribeAccountBalanceRequest) {
    request = &DescribeAccountBalanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeAccountBalance")
    
    
    return
}

func NewDescribeAccountBalanceResponse() (response *DescribeAccountBalanceResponse) {
    response = &DescribeAccountBalanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccountBalance
// 获取云账户余额信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PAYPRICEERROR = "FailedOperation.PayPriceError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAccountBalance(request *DescribeAccountBalanceRequest) (response *DescribeAccountBalanceResponse, err error) {
    return c.DescribeAccountBalanceWithContext(context.Background(), request)
}

// DescribeAccountBalance
// 获取云账户余额信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PAYPRICEERROR = "FailedOperation.PayPriceError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAccountBalanceWithContext(ctx context.Context, request *DescribeAccountBalanceRequest) (response *DescribeAccountBalanceResponse, err error) {
    if request == nil {
        request = NewDescribeAccountBalanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeAccountBalance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccountBalance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccountBalanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAllocateConditionsRequest() (request *DescribeAllocateConditionsRequest) {
    request = &DescribeAllocateConditionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeAllocateConditions")
    
    
    return
}

func NewDescribeAllocateConditionsResponse() (response *DescribeAllocateConditionsResponse) {
    response = &DescribeAllocateConditionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAllocateConditions
// 查询资源目录筛选条件
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAllocateConditions(request *DescribeAllocateConditionsRequest) (response *DescribeAllocateConditionsResponse, err error) {
    return c.DescribeAllocateConditionsWithContext(context.Background(), request)
}

// DescribeAllocateConditions
// 查询资源目录筛选条件
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAllocateConditionsWithContext(ctx context.Context, request *DescribeAllocateConditionsRequest) (response *DescribeAllocateConditionsResponse, err error) {
    if request == nil {
        request = NewDescribeAllocateConditionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeAllocateConditions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAllocateConditions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAllocateConditionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAllocationBillConditionsRequest() (request *DescribeAllocationBillConditionsRequest) {
    request = &DescribeAllocationBillConditionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeAllocationBillConditions")
    
    
    return
}

func NewDescribeAllocationBillConditionsResponse() (response *DescribeAllocationBillConditionsResponse) {
    response = &DescribeAllocationBillConditionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAllocationBillConditions
// 查询分账账单筛选条件
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAllocationBillConditions(request *DescribeAllocationBillConditionsRequest) (response *DescribeAllocationBillConditionsResponse, err error) {
    return c.DescribeAllocationBillConditionsWithContext(context.Background(), request)
}

// DescribeAllocationBillConditions
// 查询分账账单筛选条件
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAllocationBillConditionsWithContext(ctx context.Context, request *DescribeAllocationBillConditionsRequest) (response *DescribeAllocationBillConditionsResponse, err error) {
    if request == nil {
        request = NewDescribeAllocationBillConditionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeAllocationBillConditions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAllocationBillConditions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAllocationBillConditionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAllocationBillDetailRequest() (request *DescribeAllocationBillDetailRequest) {
    request = &DescribeAllocationBillDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeAllocationBillDetail")
    
    
    return
}

func NewDescribeAllocationBillDetailResponse() (response *DescribeAllocationBillDetailResponse) {
    response = &DescribeAllocationBillDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAllocationBillDetail
// 查询分账账单明细
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAllocationBillDetail(request *DescribeAllocationBillDetailRequest) (response *DescribeAllocationBillDetailResponse, err error) {
    return c.DescribeAllocationBillDetailWithContext(context.Background(), request)
}

// DescribeAllocationBillDetail
// 查询分账账单明细
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAllocationBillDetailWithContext(ctx context.Context, request *DescribeAllocationBillDetailRequest) (response *DescribeAllocationBillDetailResponse, err error) {
    if request == nil {
        request = NewDescribeAllocationBillDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeAllocationBillDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAllocationBillDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAllocationBillDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAllocationMonthOverviewRequest() (request *DescribeAllocationMonthOverviewRequest) {
    request = &DescribeAllocationMonthOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeAllocationMonthOverview")
    
    
    return
}

func NewDescribeAllocationMonthOverviewResponse() (response *DescribeAllocationMonthOverviewResponse) {
    response = &DescribeAllocationMonthOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAllocationMonthOverview
// 查询分账账单月概览
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAllocationMonthOverview(request *DescribeAllocationMonthOverviewRequest) (response *DescribeAllocationMonthOverviewResponse, err error) {
    return c.DescribeAllocationMonthOverviewWithContext(context.Background(), request)
}

// DescribeAllocationMonthOverview
// 查询分账账单月概览
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAllocationMonthOverviewWithContext(ctx context.Context, request *DescribeAllocationMonthOverviewRequest) (response *DescribeAllocationMonthOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeAllocationMonthOverviewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeAllocationMonthOverview")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAllocationMonthOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAllocationMonthOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAllocationOverviewRequest() (request *DescribeAllocationOverviewRequest) {
    request = &DescribeAllocationOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeAllocationOverview")
    
    
    return
}

func NewDescribeAllocationOverviewResponse() (response *DescribeAllocationOverviewResponse) {
    response = &DescribeAllocationOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAllocationOverview
// 查询分账账单日概览
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAllocationOverview(request *DescribeAllocationOverviewRequest) (response *DescribeAllocationOverviewResponse, err error) {
    return c.DescribeAllocationOverviewWithContext(context.Background(), request)
}

// DescribeAllocationOverview
// 查询分账账单日概览
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAllocationOverviewWithContext(ctx context.Context, request *DescribeAllocationOverviewRequest) (response *DescribeAllocationOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeAllocationOverviewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeAllocationOverview")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAllocationOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAllocationOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAllocationRuleDetailRequest() (request *DescribeAllocationRuleDetailRequest) {
    request = &DescribeAllocationRuleDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeAllocationRuleDetail")
    
    
    return
}

func NewDescribeAllocationRuleDetailResponse() (response *DescribeAllocationRuleDetailResponse) {
    response = &DescribeAllocationRuleDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAllocationRuleDetail
// 查询公摊规则详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATERROR = "InternalError.DbOperatError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAllocationRuleDetail(request *DescribeAllocationRuleDetailRequest) (response *DescribeAllocationRuleDetailResponse, err error) {
    return c.DescribeAllocationRuleDetailWithContext(context.Background(), request)
}

// DescribeAllocationRuleDetail
// 查询公摊规则详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATERROR = "InternalError.DbOperatError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAllocationRuleDetailWithContext(ctx context.Context, request *DescribeAllocationRuleDetailRequest) (response *DescribeAllocationRuleDetailResponse, err error) {
    if request == nil {
        request = NewDescribeAllocationRuleDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeAllocationRuleDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAllocationRuleDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAllocationRuleDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAllocationRuleSummaryRequest() (request *DescribeAllocationRuleSummaryRequest) {
    request = &DescribeAllocationRuleSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeAllocationRuleSummary")
    
    
    return
}

func NewDescribeAllocationRuleSummaryResponse() (response *DescribeAllocationRuleSummaryResponse) {
    response = &DescribeAllocationRuleSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAllocationRuleSummary
// 查询所有公摊规则概览
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAllocationRuleSummary(request *DescribeAllocationRuleSummaryRequest) (response *DescribeAllocationRuleSummaryResponse, err error) {
    return c.DescribeAllocationRuleSummaryWithContext(context.Background(), request)
}

// DescribeAllocationRuleSummary
// 查询所有公摊规则概览
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAllocationRuleSummaryWithContext(ctx context.Context, request *DescribeAllocationRuleSummaryRequest) (response *DescribeAllocationRuleSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeAllocationRuleSummaryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeAllocationRuleSummary")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAllocationRuleSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAllocationRuleSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAllocationSummaryByBusinessRequest() (request *DescribeAllocationSummaryByBusinessRequest) {
    request = &DescribeAllocationSummaryByBusinessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeAllocationSummaryByBusiness")
    
    
    return
}

func NewDescribeAllocationSummaryByBusinessResponse() (response *DescribeAllocationSummaryByBusinessResponse) {
    response = &DescribeAllocationSummaryByBusinessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAllocationSummaryByBusiness
// 查询分账账单按产品汇总
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAllocationSummaryByBusiness(request *DescribeAllocationSummaryByBusinessRequest) (response *DescribeAllocationSummaryByBusinessResponse, err error) {
    return c.DescribeAllocationSummaryByBusinessWithContext(context.Background(), request)
}

// DescribeAllocationSummaryByBusiness
// 查询分账账单按产品汇总
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAllocationSummaryByBusinessWithContext(ctx context.Context, request *DescribeAllocationSummaryByBusinessRequest) (response *DescribeAllocationSummaryByBusinessResponse, err error) {
    if request == nil {
        request = NewDescribeAllocationSummaryByBusinessRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeAllocationSummaryByBusiness")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAllocationSummaryByBusiness require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAllocationSummaryByBusinessResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAllocationSummaryByItemRequest() (request *DescribeAllocationSummaryByItemRequest) {
    request = &DescribeAllocationSummaryByItemRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeAllocationSummaryByItem")
    
    
    return
}

func NewDescribeAllocationSummaryByItemResponse() (response *DescribeAllocationSummaryByItemResponse) {
    response = &DescribeAllocationSummaryByItemResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAllocationSummaryByItem
// 查询分账账单按组件汇总
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAllocationSummaryByItem(request *DescribeAllocationSummaryByItemRequest) (response *DescribeAllocationSummaryByItemResponse, err error) {
    return c.DescribeAllocationSummaryByItemWithContext(context.Background(), request)
}

// DescribeAllocationSummaryByItem
// 查询分账账单按组件汇总
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAllocationSummaryByItemWithContext(ctx context.Context, request *DescribeAllocationSummaryByItemRequest) (response *DescribeAllocationSummaryByItemResponse, err error) {
    if request == nil {
        request = NewDescribeAllocationSummaryByItemRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeAllocationSummaryByItem")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAllocationSummaryByItem require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAllocationSummaryByItemResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAllocationSummaryByResourceRequest() (request *DescribeAllocationSummaryByResourceRequest) {
    request = &DescribeAllocationSummaryByResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeAllocationSummaryByResource")
    
    
    return
}

func NewDescribeAllocationSummaryByResourceResponse() (response *DescribeAllocationSummaryByResourceResponse) {
    response = &DescribeAllocationSummaryByResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAllocationSummaryByResource
// 查询分账账单按资源汇总
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAllocationSummaryByResource(request *DescribeAllocationSummaryByResourceRequest) (response *DescribeAllocationSummaryByResourceResponse, err error) {
    return c.DescribeAllocationSummaryByResourceWithContext(context.Background(), request)
}

// DescribeAllocationSummaryByResource
// 查询分账账单按资源汇总
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAllocationSummaryByResourceWithContext(ctx context.Context, request *DescribeAllocationSummaryByResourceRequest) (response *DescribeAllocationSummaryByResourceResponse, err error) {
    if request == nil {
        request = NewDescribeAllocationSummaryByResourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeAllocationSummaryByResource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAllocationSummaryByResource require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAllocationSummaryByResourceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAllocationTreeRequest() (request *DescribeAllocationTreeRequest) {
    request = &DescribeAllocationTreeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeAllocationTree")
    
    
    return
}

func NewDescribeAllocationTreeResponse() (response *DescribeAllocationTreeResponse) {
    response = &DescribeAllocationTreeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAllocationTree
// 查询分账目录树
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAllocationTree(request *DescribeAllocationTreeRequest) (response *DescribeAllocationTreeResponse, err error) {
    return c.DescribeAllocationTreeWithContext(context.Background(), request)
}

// DescribeAllocationTree
// 查询分账目录树
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAllocationTreeWithContext(ctx context.Context, request *DescribeAllocationTreeRequest) (response *DescribeAllocationTreeResponse, err error) {
    if request == nil {
        request = NewDescribeAllocationTreeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeAllocationTree")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAllocationTree require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAllocationTreeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAllocationTrendByMonthRequest() (request *DescribeAllocationTrendByMonthRequest) {
    request = &DescribeAllocationTrendByMonthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeAllocationTrendByMonth")
    
    
    return
}

func NewDescribeAllocationTrendByMonthResponse() (response *DescribeAllocationTrendByMonthResponse) {
    response = &DescribeAllocationTrendByMonthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAllocationTrendByMonth
// 查询分账账单费用趋势
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAllocationTrendByMonth(request *DescribeAllocationTrendByMonthRequest) (response *DescribeAllocationTrendByMonthResponse, err error) {
    return c.DescribeAllocationTrendByMonthWithContext(context.Background(), request)
}

// DescribeAllocationTrendByMonth
// 查询分账账单费用趋势
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAllocationTrendByMonthWithContext(ctx context.Context, request *DescribeAllocationTrendByMonthRequest) (response *DescribeAllocationTrendByMonthResponse, err error) {
    if request == nil {
        request = NewDescribeAllocationTrendByMonthRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeAllocationTrendByMonth")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAllocationTrendByMonth require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAllocationTrendByMonthResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAllocationUnitDetailRequest() (request *DescribeAllocationUnitDetailRequest) {
    request = &DescribeAllocationUnitDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeAllocationUnitDetail")
    
    
    return
}

func NewDescribeAllocationUnitDetailResponse() (response *DescribeAllocationUnitDetailResponse) {
    response = &DescribeAllocationUnitDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAllocationUnitDetail
// 查询分账单元详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAllocationUnitDetail(request *DescribeAllocationUnitDetailRequest) (response *DescribeAllocationUnitDetailResponse, err error) {
    return c.DescribeAllocationUnitDetailWithContext(context.Background(), request)
}

// DescribeAllocationUnitDetail
// 查询分账单元详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAllocationUnitDetailWithContext(ctx context.Context, request *DescribeAllocationUnitDetailRequest) (response *DescribeAllocationUnitDetailResponse, err error) {
    if request == nil {
        request = NewDescribeAllocationUnitDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeAllocationUnitDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAllocationUnitDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAllocationUnitDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillAdjustInfoRequest() (request *DescribeBillAdjustInfoRequest) {
    request = &DescribeBillAdjustInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBillAdjustInfo")
    
    
    return
}

func NewDescribeBillAdjustInfoResponse() (response *DescribeBillAdjustInfoResponse) {
    response = &DescribeBillAdjustInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillAdjustInfo
// 可以通过API获取当前UIN是否有调账，客户可以更快地主动地获取调账情况。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) DescribeBillAdjustInfo(request *DescribeBillAdjustInfoRequest) (response *DescribeBillAdjustInfoResponse, err error) {
    return c.DescribeBillAdjustInfoWithContext(context.Background(), request)
}

// DescribeBillAdjustInfo
// 可以通过API获取当前UIN是否有调账，客户可以更快地主动地获取调账情况。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) DescribeBillAdjustInfoWithContext(ctx context.Context, request *DescribeBillAdjustInfoRequest) (response *DescribeBillAdjustInfoResponse, err error) {
    if request == nil {
        request = NewDescribeBillAdjustInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeBillAdjustInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillAdjustInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillAdjustInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillDetailRequest() (request *DescribeBillDetailRequest) {
    request = &DescribeBillDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBillDetail")
    
    
    return
}

func NewDescribeBillDetailResponse() (response *DescribeBillDetailResponse) {
    response = &DescribeBillDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillDetail
// 获取账单明细数据。
//
// 注意事项：
//
// 1.在请求接口时，由于网络不稳定或其它异常，可能会导致请求失败。如果您遇到这种情况，我们建议您在接口请求失败时，手动发起重试操作，这样可以更好地确保您的接口请求能够成功执行。
//
// 2.对于账单明细数据量级很大（例如每月账单明细量级超过20w）的客户，通过 API 调用账单数据效率较低，建议您开通账单数据存储功能，通过存储桶中获取账单文件进行分析。[账单存储至COS桶](https://cloud.tencent.com/document/product/555/61275)
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYCOUNTFAILED = "FailedOperation.QueryCountFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBillDetail(request *DescribeBillDetailRequest) (response *DescribeBillDetailResponse, err error) {
    return c.DescribeBillDetailWithContext(context.Background(), request)
}

// DescribeBillDetail
// 获取账单明细数据。
//
// 注意事项：
//
// 1.在请求接口时，由于网络不稳定或其它异常，可能会导致请求失败。如果您遇到这种情况，我们建议您在接口请求失败时，手动发起重试操作，这样可以更好地确保您的接口请求能够成功执行。
//
// 2.对于账单明细数据量级很大（例如每月账单明细量级超过20w）的客户，通过 API 调用账单数据效率较低，建议您开通账单数据存储功能，通过存储桶中获取账单文件进行分析。[账单存储至COS桶](https://cloud.tencent.com/document/product/555/61275)
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYCOUNTFAILED = "FailedOperation.QueryCountFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBillDetailWithContext(ctx context.Context, request *DescribeBillDetailRequest) (response *DescribeBillDetailResponse, err error) {
    if request == nil {
        request = NewDescribeBillDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeBillDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillDetailForOrganizationRequest() (request *DescribeBillDetailForOrganizationRequest) {
    request = &DescribeBillDetailForOrganizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBillDetailForOrganization")
    
    
    return
}

func NewDescribeBillDetailForOrganizationResponse() (response *DescribeBillDetailForOrganizationResponse) {
    response = &DescribeBillDetailForOrganizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillDetailForOrganization
// 成员账号获取管理账号代付账单（费用明细）。
//
// 注意事项：在请求接口时，由于网络不稳定或其它异常，可能会导致请求失败。如果您遇到这种情况，我们建议您在接口请求失败时，手动发起重试操作，这样可以更好地确保您的接口请求能够成功执行。
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYCOUNTFAILED = "FailedOperation.QueryCountFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBillDetailForOrganization(request *DescribeBillDetailForOrganizationRequest) (response *DescribeBillDetailForOrganizationResponse, err error) {
    return c.DescribeBillDetailForOrganizationWithContext(context.Background(), request)
}

// DescribeBillDetailForOrganization
// 成员账号获取管理账号代付账单（费用明细）。
//
// 注意事项：在请求接口时，由于网络不稳定或其它异常，可能会导致请求失败。如果您遇到这种情况，我们建议您在接口请求失败时，手动发起重试操作，这样可以更好地确保您的接口请求能够成功执行。
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYCOUNTFAILED = "FailedOperation.QueryCountFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBillDetailForOrganizationWithContext(ctx context.Context, request *DescribeBillDetailForOrganizationRequest) (response *DescribeBillDetailForOrganizationResponse, err error) {
    if request == nil {
        request = NewDescribeBillDetailForOrganizationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeBillDetailForOrganization")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillDetailForOrganization require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillDetailForOrganizationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillDownloadUrlRequest() (request *DescribeBillDownloadUrlRequest) {
    request = &DescribeBillDownloadUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBillDownloadUrl")
    
    
    return
}

func NewDescribeBillDownloadUrlResponse() (response *DescribeBillDownloadUrlResponse) {
    response = &DescribeBillDownloadUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillDownloadUrl
// 该接口支持通过传参，获取L0-PDF、L1-汇总、L2-资源、L3-明细、账单包、五类账单文件下载链接
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBillDownloadUrl(request *DescribeBillDownloadUrlRequest) (response *DescribeBillDownloadUrlResponse, err error) {
    return c.DescribeBillDownloadUrlWithContext(context.Background(), request)
}

// DescribeBillDownloadUrl
// 该接口支持通过传参，获取L0-PDF、L1-汇总、L2-资源、L3-明细、账单包、五类账单文件下载链接
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBillDownloadUrlWithContext(ctx context.Context, request *DescribeBillDownloadUrlRequest) (response *DescribeBillDownloadUrlResponse, err error) {
    if request == nil {
        request = NewDescribeBillDownloadUrlRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeBillDownloadUrl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillDownloadUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillDownloadUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillListRequest() (request *DescribeBillListRequest) {
    request = &DescribeBillListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBillList")
    
    
    return
}

func NewDescribeBillListResponse() (response *DescribeBillListResponse) {
    response = &DescribeBillListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillList
// 获取收支明细列表，支持翻页和参数过滤
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeBillList(request *DescribeBillListRequest) (response *DescribeBillListResponse, err error) {
    return c.DescribeBillListWithContext(context.Background(), request)
}

// DescribeBillList
// 获取收支明细列表，支持翻页和参数过滤
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeBillListWithContext(ctx context.Context, request *DescribeBillListRequest) (response *DescribeBillListResponse, err error) {
    if request == nil {
        request = NewDescribeBillListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeBillList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillResourceSummaryRequest() (request *DescribeBillResourceSummaryRequest) {
    request = &DescribeBillResourceSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBillResourceSummary")
    
    
    return
}

func NewDescribeBillResourceSummaryResponse() (response *DescribeBillResourceSummaryResponse) {
    response = &DescribeBillResourceSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillResourceSummary
// 获取账单资源汇总数据
//
// 可能返回的错误码:
//  FAILEDOPERATION_SUMMARYDATANOTREADY = "FailedOperation.SummaryDataNotReady"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DescribeBillResourceSummary(request *DescribeBillResourceSummaryRequest) (response *DescribeBillResourceSummaryResponse, err error) {
    return c.DescribeBillResourceSummaryWithContext(context.Background(), request)
}

// DescribeBillResourceSummary
// 获取账单资源汇总数据
//
// 可能返回的错误码:
//  FAILEDOPERATION_SUMMARYDATANOTREADY = "FailedOperation.SummaryDataNotReady"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DescribeBillResourceSummaryWithContext(ctx context.Context, request *DescribeBillResourceSummaryRequest) (response *DescribeBillResourceSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeBillResourceSummaryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeBillResourceSummary")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillResourceSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillResourceSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillResourceSummaryForOrganizationRequest() (request *DescribeBillResourceSummaryForOrganizationRequest) {
    request = &DescribeBillResourceSummaryForOrganizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBillResourceSummaryForOrganization")
    
    
    return
}

func NewDescribeBillResourceSummaryForOrganizationResponse() (response *DescribeBillResourceSummaryForOrganizationResponse) {
    response = &DescribeBillResourceSummaryForOrganizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillResourceSummaryForOrganization
// 成员账号获取管理账号代付账单（按资源汇总）
//
// 可能返回的错误码:
//  FAILEDOPERATION_SUMMARYDATANOTREADY = "FailedOperation.SummaryDataNotReady"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DescribeBillResourceSummaryForOrganization(request *DescribeBillResourceSummaryForOrganizationRequest) (response *DescribeBillResourceSummaryForOrganizationResponse, err error) {
    return c.DescribeBillResourceSummaryForOrganizationWithContext(context.Background(), request)
}

// DescribeBillResourceSummaryForOrganization
// 成员账号获取管理账号代付账单（按资源汇总）
//
// 可能返回的错误码:
//  FAILEDOPERATION_SUMMARYDATANOTREADY = "FailedOperation.SummaryDataNotReady"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DescribeBillResourceSummaryForOrganizationWithContext(ctx context.Context, request *DescribeBillResourceSummaryForOrganizationRequest) (response *DescribeBillResourceSummaryForOrganizationResponse, err error) {
    if request == nil {
        request = NewDescribeBillResourceSummaryForOrganizationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeBillResourceSummaryForOrganization")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillResourceSummaryForOrganization require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillResourceSummaryForOrganizationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillSummaryRequest() (request *DescribeBillSummaryRequest) {
    request = &DescribeBillSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBillSummary")
    
    
    return
}

func NewDescribeBillSummaryResponse() (response *DescribeBillSummaryResponse) {
    response = &DescribeBillSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillSummary
// 该接口支持通过传参，按照产品、项目、地域、计费模式和标签五个维度获取账单费用明细。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TAGKEYNOTEXIST = "FailedOperation.TagKeyNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummary(request *DescribeBillSummaryRequest) (response *DescribeBillSummaryResponse, err error) {
    return c.DescribeBillSummaryWithContext(context.Background(), request)
}

// DescribeBillSummary
// 该接口支持通过传参，按照产品、项目、地域、计费模式和标签五个维度获取账单费用明细。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TAGKEYNOTEXIST = "FailedOperation.TagKeyNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryWithContext(ctx context.Context, request *DescribeBillSummaryRequest) (response *DescribeBillSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeBillSummaryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeBillSummary")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillSummaryByPayModeRequest() (request *DescribeBillSummaryByPayModeRequest) {
    request = &DescribeBillSummaryByPayModeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBillSummaryByPayMode")
    
    
    return
}

func NewDescribeBillSummaryByPayModeResponse() (response *DescribeBillSummaryByPayModeResponse) {
    response = &DescribeBillSummaryByPayModeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillSummaryByPayMode
// 获取按计费模式汇总费用分布
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryByPayMode(request *DescribeBillSummaryByPayModeRequest) (response *DescribeBillSummaryByPayModeResponse, err error) {
    return c.DescribeBillSummaryByPayModeWithContext(context.Background(), request)
}

// DescribeBillSummaryByPayMode
// 获取按计费模式汇总费用分布
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryByPayModeWithContext(ctx context.Context, request *DescribeBillSummaryByPayModeRequest) (response *DescribeBillSummaryByPayModeResponse, err error) {
    if request == nil {
        request = NewDescribeBillSummaryByPayModeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeBillSummaryByPayMode")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillSummaryByPayMode require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillSummaryByPayModeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillSummaryByProductRequest() (request *DescribeBillSummaryByProductRequest) {
    request = &DescribeBillSummaryByProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBillSummaryByProduct")
    
    
    return
}

func NewDescribeBillSummaryByProductResponse() (response *DescribeBillSummaryByProductResponse) {
    response = &DescribeBillSummaryByProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillSummaryByProduct
// 获取产品汇总费用分布
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryByProduct(request *DescribeBillSummaryByProductRequest) (response *DescribeBillSummaryByProductResponse, err error) {
    return c.DescribeBillSummaryByProductWithContext(context.Background(), request)
}

// DescribeBillSummaryByProduct
// 获取产品汇总费用分布
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryByProductWithContext(ctx context.Context, request *DescribeBillSummaryByProductRequest) (response *DescribeBillSummaryByProductResponse, err error) {
    if request == nil {
        request = NewDescribeBillSummaryByProductRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeBillSummaryByProduct")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillSummaryByProduct require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillSummaryByProductResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillSummaryByProjectRequest() (request *DescribeBillSummaryByProjectRequest) {
    request = &DescribeBillSummaryByProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBillSummaryByProject")
    
    
    return
}

func NewDescribeBillSummaryByProjectResponse() (response *DescribeBillSummaryByProjectResponse) {
    response = &DescribeBillSummaryByProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillSummaryByProject
// 获取按项目汇总费用分布
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryByProject(request *DescribeBillSummaryByProjectRequest) (response *DescribeBillSummaryByProjectResponse, err error) {
    return c.DescribeBillSummaryByProjectWithContext(context.Background(), request)
}

// DescribeBillSummaryByProject
// 获取按项目汇总费用分布
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryByProjectWithContext(ctx context.Context, request *DescribeBillSummaryByProjectRequest) (response *DescribeBillSummaryByProjectResponse, err error) {
    if request == nil {
        request = NewDescribeBillSummaryByProjectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeBillSummaryByProject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillSummaryByProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillSummaryByProjectResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillSummaryByRegionRequest() (request *DescribeBillSummaryByRegionRequest) {
    request = &DescribeBillSummaryByRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBillSummaryByRegion")
    
    
    return
}

func NewDescribeBillSummaryByRegionResponse() (response *DescribeBillSummaryByRegionResponse) {
    response = &DescribeBillSummaryByRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillSummaryByRegion
// 获取按地域汇总费用分布
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryByRegion(request *DescribeBillSummaryByRegionRequest) (response *DescribeBillSummaryByRegionResponse, err error) {
    return c.DescribeBillSummaryByRegionWithContext(context.Background(), request)
}

// DescribeBillSummaryByRegion
// 获取按地域汇总费用分布
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryByRegionWithContext(ctx context.Context, request *DescribeBillSummaryByRegionRequest) (response *DescribeBillSummaryByRegionResponse, err error) {
    if request == nil {
        request = NewDescribeBillSummaryByRegionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeBillSummaryByRegion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillSummaryByRegion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillSummaryByRegionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillSummaryByTagRequest() (request *DescribeBillSummaryByTagRequest) {
    request = &DescribeBillSummaryByTagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBillSummaryByTag")
    
    
    return
}

func NewDescribeBillSummaryByTagResponse() (response *DescribeBillSummaryByTagResponse) {
    response = &DescribeBillSummaryByTagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillSummaryByTag
// 获取按标签汇总费用分布
//
// 可能返回的错误码:
//  FAILEDOPERATION_TAGKEYNOTEXIST = "FailedOperation.TagKeyNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryByTag(request *DescribeBillSummaryByTagRequest) (response *DescribeBillSummaryByTagResponse, err error) {
    return c.DescribeBillSummaryByTagWithContext(context.Background(), request)
}

// DescribeBillSummaryByTag
// 获取按标签汇总费用分布
//
// 可能返回的错误码:
//  FAILEDOPERATION_TAGKEYNOTEXIST = "FailedOperation.TagKeyNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryByTagWithContext(ctx context.Context, request *DescribeBillSummaryByTagRequest) (response *DescribeBillSummaryByTagResponse, err error) {
    if request == nil {
        request = NewDescribeBillSummaryByTagRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeBillSummaryByTag")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillSummaryByTag require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillSummaryByTagResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillSummaryForOrganizationRequest() (request *DescribeBillSummaryForOrganizationRequest) {
    request = &DescribeBillSummaryForOrganizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBillSummaryForOrganization")
    
    
    return
}

func NewDescribeBillSummaryForOrganizationResponse() (response *DescribeBillSummaryForOrganizationResponse) {
    response = &DescribeBillSummaryForOrganizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillSummaryForOrganization
// 该接口支持通过传参，按照产品、项目、地域、计费模式和标签五个维度获取账单费用明细。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TAGKEYNOTEXIST = "FailedOperation.TagKeyNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryForOrganization(request *DescribeBillSummaryForOrganizationRequest) (response *DescribeBillSummaryForOrganizationResponse, err error) {
    return c.DescribeBillSummaryForOrganizationWithContext(context.Background(), request)
}

// DescribeBillSummaryForOrganization
// 该接口支持通过传参，按照产品、项目、地域、计费模式和标签五个维度获取账单费用明细。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TAGKEYNOTEXIST = "FailedOperation.TagKeyNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryForOrganizationWithContext(ctx context.Context, request *DescribeBillSummaryForOrganizationRequest) (response *DescribeBillSummaryForOrganizationResponse, err error) {
    if request == nil {
        request = NewDescribeBillSummaryForOrganizationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeBillSummaryForOrganization")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillSummaryForOrganization require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillSummaryForOrganizationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBudgetRequest() (request *DescribeBudgetRequest) {
    request = &DescribeBudgetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBudget")
    
    
    return
}

func NewDescribeBudgetResponse() (response *DescribeBudgetResponse) {
    response = &DescribeBudgetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBudget
// 获取预算详细信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_TAGKEYNOTEXIST = "FailedOperation.TagKeyNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBudget(request *DescribeBudgetRequest) (response *DescribeBudgetResponse, err error) {
    return c.DescribeBudgetWithContext(context.Background(), request)
}

// DescribeBudget
// 获取预算详细信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_TAGKEYNOTEXIST = "FailedOperation.TagKeyNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBudgetWithContext(ctx context.Context, request *DescribeBudgetRequest) (response *DescribeBudgetResponse, err error) {
    if request == nil {
        request = NewDescribeBudgetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeBudget")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBudget require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBudgetResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBudgetOperationLogRequest() (request *DescribeBudgetOperationLogRequest) {
    request = &DescribeBudgetOperationLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBudgetOperationLog")
    
    
    return
}

func NewDescribeBudgetOperationLogResponse() (response *DescribeBudgetOperationLogResponse) {
    response = &DescribeBudgetOperationLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBudgetOperationLog
// 查询预算修改记录
//
// 可能返回的错误码:
//  FAILEDOPERATION_TAGKEYNOTEXIST = "FailedOperation.TagKeyNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBudgetOperationLog(request *DescribeBudgetOperationLogRequest) (response *DescribeBudgetOperationLogResponse, err error) {
    return c.DescribeBudgetOperationLogWithContext(context.Background(), request)
}

// DescribeBudgetOperationLog
// 查询预算修改记录
//
// 可能返回的错误码:
//  FAILEDOPERATION_TAGKEYNOTEXIST = "FailedOperation.TagKeyNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBudgetOperationLogWithContext(ctx context.Context, request *DescribeBudgetOperationLogRequest) (response *DescribeBudgetOperationLogResponse, err error) {
    if request == nil {
        request = NewDescribeBudgetOperationLogRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeBudgetOperationLog")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBudgetOperationLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBudgetOperationLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBudgetRemindRecordListRequest() (request *DescribeBudgetRemindRecordListRequest) {
    request = &DescribeBudgetRemindRecordListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBudgetRemindRecordList")
    
    
    return
}

func NewDescribeBudgetRemindRecordListResponse() (response *DescribeBudgetRemindRecordListResponse) {
    response = &DescribeBudgetRemindRecordListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBudgetRemindRecordList
// 返回预算提醒记录，包括预算周期、检测时间、提醒时间、提醒类型、提醒内容
//
// 可能返回的错误码:
//  FAILEDOPERATION_TAGKEYNOTEXIST = "FailedOperation.TagKeyNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBudgetRemindRecordList(request *DescribeBudgetRemindRecordListRequest) (response *DescribeBudgetRemindRecordListResponse, err error) {
    return c.DescribeBudgetRemindRecordListWithContext(context.Background(), request)
}

// DescribeBudgetRemindRecordList
// 返回预算提醒记录，包括预算周期、检测时间、提醒时间、提醒类型、提醒内容
//
// 可能返回的错误码:
//  FAILEDOPERATION_TAGKEYNOTEXIST = "FailedOperation.TagKeyNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBudgetRemindRecordListWithContext(ctx context.Context, request *DescribeBudgetRemindRecordListRequest) (response *DescribeBudgetRemindRecordListResponse, err error) {
    if request == nil {
        request = NewDescribeBudgetRemindRecordListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeBudgetRemindRecordList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBudgetRemindRecordList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBudgetRemindRecordListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCostDetailRequest() (request *DescribeCostDetailRequest) {
    request = &DescribeCostDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeCostDetail")
    
    
    return
}

func NewDescribeCostDetailResponse() (response *DescribeCostDetailResponse) {
    response = &DescribeCostDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCostDetail
// 查询消耗明细
//
// 
//
// 注意事项：
//
// 1. 对于消耗明细数据量级很大（例如每月消耗明细量级超过100w）的客户，通过 API 调用明细数据会有超时风险，建议您开通消耗账单数据存储功能，通过存储桶中获取账单文件进行分析。[账单存储至COS桶](https://cloud.tencent.com/document/product/555/61275)
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DescribeCostDetail(request *DescribeCostDetailRequest) (response *DescribeCostDetailResponse, err error) {
    return c.DescribeCostDetailWithContext(context.Background(), request)
}

// DescribeCostDetail
// 查询消耗明细
//
// 
//
// 注意事项：
//
// 1. 对于消耗明细数据量级很大（例如每月消耗明细量级超过100w）的客户，通过 API 调用明细数据会有超时风险，建议您开通消耗账单数据存储功能，通过存储桶中获取账单文件进行分析。[账单存储至COS桶](https://cloud.tencent.com/document/product/555/61275)
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DescribeCostDetailWithContext(ctx context.Context, request *DescribeCostDetailRequest) (response *DescribeCostDetailResponse, err error) {
    if request == nil {
        request = NewDescribeCostDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeCostDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCostDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCostDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCostExplorerSummaryRequest() (request *DescribeCostExplorerSummaryRequest) {
    request = &DescribeCostExplorerSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeCostExplorerSummary")
    
    
    return
}

func NewDescribeCostExplorerSummaryResponse() (response *DescribeCostExplorerSummaryResponse) {
    response = &DescribeCostExplorerSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCostExplorerSummary
// 查看成本分析明细
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DescribeCostExplorerSummary(request *DescribeCostExplorerSummaryRequest) (response *DescribeCostExplorerSummaryResponse, err error) {
    return c.DescribeCostExplorerSummaryWithContext(context.Background(), request)
}

// DescribeCostExplorerSummary
// 查看成本分析明细
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DescribeCostExplorerSummaryWithContext(ctx context.Context, request *DescribeCostExplorerSummaryRequest) (response *DescribeCostExplorerSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeCostExplorerSummaryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeCostExplorerSummary")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCostExplorerSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCostExplorerSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCostSummaryByProductRequest() (request *DescribeCostSummaryByProductRequest) {
    request = &DescribeCostSummaryByProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeCostSummaryByProduct")
    
    
    return
}

func NewDescribeCostSummaryByProductResponse() (response *DescribeCostSummaryByProductResponse) {
    response = &DescribeCostSummaryByProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCostSummaryByProduct
// 获取按产品汇总消耗详情
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DescribeCostSummaryByProduct(request *DescribeCostSummaryByProductRequest) (response *DescribeCostSummaryByProductResponse, err error) {
    return c.DescribeCostSummaryByProductWithContext(context.Background(), request)
}

// DescribeCostSummaryByProduct
// 获取按产品汇总消耗详情
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DescribeCostSummaryByProductWithContext(ctx context.Context, request *DescribeCostSummaryByProductRequest) (response *DescribeCostSummaryByProductResponse, err error) {
    if request == nil {
        request = NewDescribeCostSummaryByProductRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeCostSummaryByProduct")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCostSummaryByProduct require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCostSummaryByProductResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCostSummaryByProjectRequest() (request *DescribeCostSummaryByProjectRequest) {
    request = &DescribeCostSummaryByProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeCostSummaryByProject")
    
    
    return
}

func NewDescribeCostSummaryByProjectResponse() (response *DescribeCostSummaryByProjectResponse) {
    response = &DescribeCostSummaryByProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCostSummaryByProject
// 获取按项目汇总消耗详情
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DescribeCostSummaryByProject(request *DescribeCostSummaryByProjectRequest) (response *DescribeCostSummaryByProjectResponse, err error) {
    return c.DescribeCostSummaryByProjectWithContext(context.Background(), request)
}

// DescribeCostSummaryByProject
// 获取按项目汇总消耗详情
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DescribeCostSummaryByProjectWithContext(ctx context.Context, request *DescribeCostSummaryByProjectRequest) (response *DescribeCostSummaryByProjectResponse, err error) {
    if request == nil {
        request = NewDescribeCostSummaryByProjectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeCostSummaryByProject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCostSummaryByProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCostSummaryByProjectResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCostSummaryByRegionRequest() (request *DescribeCostSummaryByRegionRequest) {
    request = &DescribeCostSummaryByRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeCostSummaryByRegion")
    
    
    return
}

func NewDescribeCostSummaryByRegionResponse() (response *DescribeCostSummaryByRegionResponse) {
    response = &DescribeCostSummaryByRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCostSummaryByRegion
// 获取按地域汇总消耗详情
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DescribeCostSummaryByRegion(request *DescribeCostSummaryByRegionRequest) (response *DescribeCostSummaryByRegionResponse, err error) {
    return c.DescribeCostSummaryByRegionWithContext(context.Background(), request)
}

// DescribeCostSummaryByRegion
// 获取按地域汇总消耗详情
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DescribeCostSummaryByRegionWithContext(ctx context.Context, request *DescribeCostSummaryByRegionRequest) (response *DescribeCostSummaryByRegionResponse, err error) {
    if request == nil {
        request = NewDescribeCostSummaryByRegionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeCostSummaryByRegion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCostSummaryByRegion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCostSummaryByRegionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCostSummaryByResourceRequest() (request *DescribeCostSummaryByResourceRequest) {
    request = &DescribeCostSummaryByResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeCostSummaryByResource")
    
    
    return
}

func NewDescribeCostSummaryByResourceResponse() (response *DescribeCostSummaryByResourceResponse) {
    response = &DescribeCostSummaryByResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCostSummaryByResource
// 获取按资源汇总消耗详情
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DescribeCostSummaryByResource(request *DescribeCostSummaryByResourceRequest) (response *DescribeCostSummaryByResourceResponse, err error) {
    return c.DescribeCostSummaryByResourceWithContext(context.Background(), request)
}

// DescribeCostSummaryByResource
// 获取按资源汇总消耗详情
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DescribeCostSummaryByResourceWithContext(ctx context.Context, request *DescribeCostSummaryByResourceRequest) (response *DescribeCostSummaryByResourceResponse, err error) {
    if request == nil {
        request = NewDescribeCostSummaryByResourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeCostSummaryByResource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCostSummaryByResource require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCostSummaryByResourceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDealsByCondRequest() (request *DescribeDealsByCondRequest) {
    request = &DescribeDealsByCondRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeDealsByCond")
    
    
    return
}

func NewDescribeDealsByCondResponse() (response *DescribeDealsByCondResponse) {
    response = &DescribeDealsByCondResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDealsByCond
// 查询订单
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDealsByCond(request *DescribeDealsByCondRequest) (response *DescribeDealsByCondResponse, err error) {
    return c.DescribeDealsByCondWithContext(context.Background(), request)
}

// DescribeDealsByCond
// 查询订单
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDealsByCondWithContext(ctx context.Context, request *DescribeDealsByCondRequest) (response *DescribeDealsByCondResponse, err error) {
    if request == nil {
        request = NewDescribeDealsByCondRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeDealsByCond")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDealsByCond require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDealsByCondResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDosageCosDetailByDateRequest() (request *DescribeDosageCosDetailByDateRequest) {
    request = &DescribeDosageCosDetailByDateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeDosageCosDetailByDate")
    
    
    return
}

func NewDescribeDosageCosDetailByDateResponse() (response *DescribeDosageCosDetailByDateResponse) {
    response = &DescribeDosageCosDetailByDateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDosageCosDetailByDate
// 获取COS产品用量明细
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DescribeDosageCosDetailByDate(request *DescribeDosageCosDetailByDateRequest) (response *DescribeDosageCosDetailByDateResponse, err error) {
    return c.DescribeDosageCosDetailByDateWithContext(context.Background(), request)
}

// DescribeDosageCosDetailByDate
// 获取COS产品用量明细
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DescribeDosageCosDetailByDateWithContext(ctx context.Context, request *DescribeDosageCosDetailByDateRequest) (response *DescribeDosageCosDetailByDateResponse, err error) {
    if request == nil {
        request = NewDescribeDosageCosDetailByDateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeDosageCosDetailByDate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDosageCosDetailByDate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDosageCosDetailByDateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDosageDetailByDateRequest() (request *DescribeDosageDetailByDateRequest) {
    request = &DescribeDosageDetailByDateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeDosageDetailByDate")
    
    
    return
}

func NewDescribeDosageDetailByDateResponse() (response *DescribeDosageDetailByDateResponse) {
    response = &DescribeDosageDetailByDateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDosageDetailByDate
// 按日期获取产品用量明细
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DescribeDosageDetailByDate(request *DescribeDosageDetailByDateRequest) (response *DescribeDosageDetailByDateResponse, err error) {
    return c.DescribeDosageDetailByDateWithContext(context.Background(), request)
}

// DescribeDosageDetailByDate
// 按日期获取产品用量明细
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DescribeDosageDetailByDateWithContext(ctx context.Context, request *DescribeDosageDetailByDateRequest) (response *DescribeDosageDetailByDateResponse, err error) {
    if request == nil {
        request = NewDescribeDosageDetailByDateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeDosageDetailByDate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDosageDetailByDate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDosageDetailByDateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDosageDetailListRequest() (request *DescribeDosageDetailListRequest) {
    request = &DescribeDosageDetailListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeDosageDetailList")
    
    
    return
}

func NewDescribeDosageDetailListResponse() (response *DescribeDosageDetailListResponse) {
    response = &DescribeDosageDetailListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDosageDetailList
// 获取已接入标准用量明细模板产品的用量明细数据，目前已接入并支持查询的产品包括：云联络中心、实时音视频、实时音视频、智能媒资托管、CODING DevOps、全球IP应用加速
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYDBFAILED = "FailedOperation.QueryDBFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_PARAMCHECKFAILED = "InvalidParameter.ParamCheckFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDosageDetailList(request *DescribeDosageDetailListRequest) (response *DescribeDosageDetailListResponse, err error) {
    return c.DescribeDosageDetailListWithContext(context.Background(), request)
}

// DescribeDosageDetailList
// 获取已接入标准用量明细模板产品的用量明细数据，目前已接入并支持查询的产品包括：云联络中心、实时音视频、实时音视频、智能媒资托管、CODING DevOps、全球IP应用加速
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYDBFAILED = "FailedOperation.QueryDBFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_PARAMCHECKFAILED = "InvalidParameter.ParamCheckFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDosageDetailListWithContext(ctx context.Context, request *DescribeDosageDetailListRequest) (response *DescribeDosageDetailListResponse, err error) {
    if request == nil {
        request = NewDescribeDosageDetailListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeDosageDetailList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDosageDetailList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDosageDetailListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGatherResourceRequest() (request *DescribeGatherResourceRequest) {
    request = &DescribeGatherResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeGatherResource")
    
    
    return
}

func NewDescribeGatherResourceResponse() (response *DescribeGatherResourceResponse) {
    response = &DescribeGatherResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGatherResource
// 查询分账账单资源归集汇总
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYDBFAILED = "FailedOperation.QueryDBFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_PARAMCHECKFAILED = "InvalidParameter.ParamCheckFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeGatherResource(request *DescribeGatherResourceRequest) (response *DescribeGatherResourceResponse, err error) {
    return c.DescribeGatherResourceWithContext(context.Background(), request)
}

// DescribeGatherResource
// 查询分账账单资源归集汇总
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYDBFAILED = "FailedOperation.QueryDBFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_PARAMCHECKFAILED = "InvalidParameter.ParamCheckFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeGatherResourceWithContext(ctx context.Context, request *DescribeGatherResourceRequest) (response *DescribeGatherResourceResponse, err error) {
    if request == nil {
        request = NewDescribeGatherResourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeGatherResource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGatherResource require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGatherResourceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGatherRuleDetailRequest() (request *DescribeGatherRuleDetailRequest) {
    request = &DescribeGatherRuleDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeGatherRuleDetail")
    
    
    return
}

func NewDescribeGatherRuleDetailResponse() (response *DescribeGatherRuleDetailResponse) {
    response = &DescribeGatherRuleDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGatherRuleDetail
// 查询归集规则详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATERROR = "InternalError.DbOperatError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeGatherRuleDetail(request *DescribeGatherRuleDetailRequest) (response *DescribeGatherRuleDetailResponse, err error) {
    return c.DescribeGatherRuleDetailWithContext(context.Background(), request)
}

// DescribeGatherRuleDetail
// 查询归集规则详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATERROR = "InternalError.DbOperatError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeGatherRuleDetailWithContext(ctx context.Context, request *DescribeGatherRuleDetailRequest) (response *DescribeGatherRuleDetailResponse, err error) {
    if request == nil {
        request = NewDescribeGatherRuleDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeGatherRuleDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGatherRuleDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGatherRuleDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSavingPlanResourceInfoRequest() (request *DescribeSavingPlanResourceInfoRequest) {
    request = &DescribeSavingPlanResourceInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeSavingPlanResourceInfo")
    
    
    return
}

func NewDescribeSavingPlanResourceInfoResponse() (response *DescribeSavingPlanResourceInfoResponse) {
    response = &DescribeSavingPlanResourceInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSavingPlanResourceInfo
// 查询节省计划详情
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeSavingPlanResourceInfo(request *DescribeSavingPlanResourceInfoRequest) (response *DescribeSavingPlanResourceInfoResponse, err error) {
    return c.DescribeSavingPlanResourceInfoWithContext(context.Background(), request)
}

// DescribeSavingPlanResourceInfo
// 查询节省计划详情
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeSavingPlanResourceInfoWithContext(ctx context.Context, request *DescribeSavingPlanResourceInfoRequest) (response *DescribeSavingPlanResourceInfoResponse, err error) {
    if request == nil {
        request = NewDescribeSavingPlanResourceInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeSavingPlanResourceInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSavingPlanResourceInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSavingPlanResourceInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTagListRequest() (request *DescribeTagListRequest) {
    request = &DescribeTagListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeTagList")
    
    
    return
}

func NewDescribeTagListResponse() (response *DescribeTagListResponse) {
    response = &DescribeTagListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTagList
// 获取分账标签
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeTagList(request *DescribeTagListRequest) (response *DescribeTagListResponse, err error) {
    return c.DescribeTagListWithContext(context.Background(), request)
}

// DescribeTagList
// 获取分账标签
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeTagListWithContext(ctx context.Context, request *DescribeTagListRequest) (response *DescribeTagListResponse, err error) {
    if request == nil {
        request = NewDescribeTagListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeTagList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTagList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTagListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVoucherInfoRequest() (request *DescribeVoucherInfoRequest) {
    request = &DescribeVoucherInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeVoucherInfo")
    
    
    return
}

func NewDescribeVoucherInfoResponse() (response *DescribeVoucherInfoResponse) {
    response = &DescribeVoucherInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVoucherInfo
// 获取代金券相关信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDAPPID = "FailedOperation.InvalidAppId"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeVoucherInfo(request *DescribeVoucherInfoRequest) (response *DescribeVoucherInfoResponse, err error) {
    return c.DescribeVoucherInfoWithContext(context.Background(), request)
}

// DescribeVoucherInfo
// 获取代金券相关信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDAPPID = "FailedOperation.InvalidAppId"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeVoucherInfoWithContext(ctx context.Context, request *DescribeVoucherInfoRequest) (response *DescribeVoucherInfoResponse, err error) {
    if request == nil {
        request = NewDescribeVoucherInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeVoucherInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVoucherInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVoucherInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVoucherUsageDetailsRequest() (request *DescribeVoucherUsageDetailsRequest) {
    request = &DescribeVoucherUsageDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeVoucherUsageDetails")
    
    
    return
}

func NewDescribeVoucherUsageDetailsResponse() (response *DescribeVoucherUsageDetailsResponse) {
    response = &DescribeVoucherUsageDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVoucherUsageDetails
// 获取代金券使用记录
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDAPPID = "FailedOperation.InvalidAppId"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeVoucherUsageDetails(request *DescribeVoucherUsageDetailsRequest) (response *DescribeVoucherUsageDetailsResponse, err error) {
    return c.DescribeVoucherUsageDetailsWithContext(context.Background(), request)
}

// DescribeVoucherUsageDetails
// 获取代金券使用记录
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDAPPID = "FailedOperation.InvalidAppId"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeVoucherUsageDetailsWithContext(ctx context.Context, request *DescribeVoucherUsageDetailsRequest) (response *DescribeVoucherUsageDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeVoucherUsageDetailsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeVoucherUsageDetails")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVoucherUsageDetails require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVoucherUsageDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAllocationRuleRequest() (request *ModifyAllocationRuleRequest) {
    request = &ModifyAllocationRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "ModifyAllocationRule")
    
    
    return
}

func NewModifyAllocationRuleResponse() (response *ModifyAllocationRuleResponse) {
    response = &ModifyAllocationRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAllocationRule
// 编辑公摊规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_NUMLIMITERROR = "FailedOperation.NumLimitError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATERROR = "InternalError.DbOperatError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyAllocationRule(request *ModifyAllocationRuleRequest) (response *ModifyAllocationRuleResponse, err error) {
    return c.ModifyAllocationRuleWithContext(context.Background(), request)
}

// ModifyAllocationRule
// 编辑公摊规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_NUMLIMITERROR = "FailedOperation.NumLimitError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATERROR = "InternalError.DbOperatError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyAllocationRuleWithContext(ctx context.Context, request *ModifyAllocationRuleRequest) (response *ModifyAllocationRuleResponse, err error) {
    if request == nil {
        request = NewModifyAllocationRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "ModifyAllocationRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAllocationRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAllocationRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAllocationUnitRequest() (request *ModifyAllocationUnitRequest) {
    request = &ModifyAllocationUnitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "ModifyAllocationUnit")
    
    
    return
}

func NewModifyAllocationUnitResponse() (response *ModifyAllocationUnitResponse) {
    response = &ModifyAllocationUnitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAllocationUnit
// 修改分账单元信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATERROR = "InternalError.DbOperatError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyAllocationUnit(request *ModifyAllocationUnitRequest) (response *ModifyAllocationUnitResponse, err error) {
    return c.ModifyAllocationUnitWithContext(context.Background(), request)
}

// ModifyAllocationUnit
// 修改分账单元信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATERROR = "InternalError.DbOperatError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyAllocationUnitWithContext(ctx context.Context, request *ModifyAllocationUnitRequest) (response *ModifyAllocationUnitResponse, err error) {
    if request == nil {
        request = NewModifyAllocationUnitRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "ModifyAllocationUnit")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAllocationUnit require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAllocationUnitResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBudgetRequest() (request *ModifyBudgetRequest) {
    request = &ModifyBudgetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "ModifyBudget")
    
    
    return
}

func NewModifyBudgetResponse() (response *ModifyBudgetResponse) {
    response = &ModifyBudgetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyBudget
// 更新预算信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATERROR = "InternalError.DbOperatError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyBudget(request *ModifyBudgetRequest) (response *ModifyBudgetResponse, err error) {
    return c.ModifyBudgetWithContext(context.Background(), request)
}

// ModifyBudget
// 更新预算信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATERROR = "InternalError.DbOperatError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyBudgetWithContext(ctx context.Context, request *ModifyBudgetRequest) (response *ModifyBudgetResponse, err error) {
    if request == nil {
        request = NewModifyBudgetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "ModifyBudget")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBudget require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBudgetResponse()
    err = c.Send(request, response)
    return
}

func NewModifyGatherRuleRequest() (request *ModifyGatherRuleRequest) {
    request = &ModifyGatherRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "ModifyGatherRule")
    
    
    return
}

func NewModifyGatherRuleResponse() (response *ModifyGatherRuleResponse) {
    response = &ModifyGatherRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyGatherRule
// 编辑归集规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_NUMLIMITERROR = "FailedOperation.NumLimitError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATERROR = "InternalError.DbOperatError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyGatherRule(request *ModifyGatherRuleRequest) (response *ModifyGatherRuleResponse, err error) {
    return c.ModifyGatherRuleWithContext(context.Background(), request)
}

// ModifyGatherRule
// 编辑归集规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_NUMLIMITERROR = "FailedOperation.NumLimitError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATERROR = "InternalError.DbOperatError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyGatherRuleWithContext(ctx context.Context, request *ModifyGatherRuleRequest) (response *ModifyGatherRuleResponse, err error) {
    if request == nil {
        request = NewModifyGatherRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "ModifyGatherRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyGatherRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyGatherRuleResponse()
    err = c.Send(request, response)
    return
}

func NewPayDealsRequest() (request *PayDealsRequest) {
    request = &PayDealsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "PayDeals")
    
    
    return
}

func NewPayDealsResponse() (response *PayDealsResponse) {
    response = &PayDealsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PayDeals
// 支付订单
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENTPAYDEALCANNOTDOWN = "FailedOperation.AgentPayDealCannotDown"
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  FAILEDOPERATION_INVALIDDEAL = "FailedOperation.InvalidDeal"
//  FAILEDOPERATION_INVALIDVOUCHER = "FailedOperation.InvalidVoucher"
//  FAILEDOPERATION_NEEDPAYTOGETER = "FailedOperation.NeedPayTogeter"
//  FAILEDOPERATION_NEEDPAYTOGETHER = "FailedOperation.NeedPayTogether"
//  FAILEDOPERATION_PAYPRICEERROR = "FailedOperation.PayPriceError"
//  FAILEDOPERATION_PAYSUCCDELIVERFAILED = "FailedOperation.PaySuccDeliverFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CERTIFICATIONNEEDUPGRADE = "UnauthorizedOperation.CertificationNeedUpgrade"
//  UNAUTHORIZEDOPERATION_NOTCERTIFICATION = "UnauthorizedOperation.NotCertification"
func (c *Client) PayDeals(request *PayDealsRequest) (response *PayDealsResponse, err error) {
    return c.PayDealsWithContext(context.Background(), request)
}

// PayDeals
// 支付订单
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENTPAYDEALCANNOTDOWN = "FailedOperation.AgentPayDealCannotDown"
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  FAILEDOPERATION_INVALIDDEAL = "FailedOperation.InvalidDeal"
//  FAILEDOPERATION_INVALIDVOUCHER = "FailedOperation.InvalidVoucher"
//  FAILEDOPERATION_NEEDPAYTOGETER = "FailedOperation.NeedPayTogeter"
//  FAILEDOPERATION_NEEDPAYTOGETHER = "FailedOperation.NeedPayTogether"
//  FAILEDOPERATION_PAYPRICEERROR = "FailedOperation.PayPriceError"
//  FAILEDOPERATION_PAYSUCCDELIVERFAILED = "FailedOperation.PaySuccDeliverFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CERTIFICATIONNEEDUPGRADE = "UnauthorizedOperation.CertificationNeedUpgrade"
//  UNAUTHORIZEDOPERATION_NOTCERTIFICATION = "UnauthorizedOperation.NotCertification"
func (c *Client) PayDealsWithContext(ctx context.Context, request *PayDealsRequest) (response *PayDealsResponse, err error) {
    if request == nil {
        request = NewPayDealsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "PayDeals")
    
    if c.GetCredential() == nil {
        return nil, errors.New("PayDeals require credential")
    }

    request.SetContext(ctx)
    
    response = NewPayDealsResponse()
    err = c.Send(request, response)
    return
}
