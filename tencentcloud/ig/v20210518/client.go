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

package v20210518

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-05-18"

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


func NewDescribeIgOrderListRequest() (request *DescribeIgOrderListRequest) {
    request = &DescribeIgOrderListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ig", APIVersion, "DescribeIgOrderList")
    
    
    return
}

func NewDescribeIgOrderListResponse() (response *DescribeIgOrderListResponse) {
    response = &DescribeIgOrderListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIgOrderList
// 查询智能导诊订单列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeIgOrderList(request *DescribeIgOrderListRequest) (response *DescribeIgOrderListResponse, err error) {
    return c.DescribeIgOrderListWithContext(context.Background(), request)
}

// DescribeIgOrderList
// 查询智能导诊订单列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeIgOrderListWithContext(ctx context.Context, request *DescribeIgOrderListRequest) (response *DescribeIgOrderListResponse, err error) {
    if request == nil {
        request = NewDescribeIgOrderListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ig", APIVersion, "DescribeIgOrderList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIgOrderList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIgOrderListResponse()
    err = c.Send(request, response)
    return
}

func NewGetLLMDiagnosisDrugRequest() (request *GetLLMDiagnosisDrugRequest) {
    request = &GetLLMDiagnosisDrugRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ig", APIVersion, "GetLLMDiagnosisDrug")
    
    
    return
}

func NewGetLLMDiagnosisDrugResponse() (response *GetLLMDiagnosisDrugResponse) {
    response = &GetLLMDiagnosisDrugResponse{} 
    return

}

// GetLLMDiagnosisDrug
// 大模型问药拍药盒
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) GetLLMDiagnosisDrug(request *GetLLMDiagnosisDrugRequest) (response *GetLLMDiagnosisDrugResponse, err error) {
    return c.GetLLMDiagnosisDrugWithContext(context.Background(), request)
}

// GetLLMDiagnosisDrug
// 大模型问药拍药盒
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) GetLLMDiagnosisDrugWithContext(ctx context.Context, request *GetLLMDiagnosisDrugRequest) (response *GetLLMDiagnosisDrugResponse, err error) {
    if request == nil {
        request = NewGetLLMDiagnosisDrugRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ig", APIVersion, "GetLLMDiagnosisDrug")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetLLMDiagnosisDrug require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetLLMDiagnosisDrugResponse()
    err = c.Send(request, response)
    return
}

func NewGetLLMDiagnosisDrugChatRequest() (request *GetLLMDiagnosisDrugChatRequest) {
    request = &GetLLMDiagnosisDrugChatRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ig", APIVersion, "GetLLMDiagnosisDrugChat")
    
    
    return
}

func NewGetLLMDiagnosisDrugChatResponse() (response *GetLLMDiagnosisDrugChatResponse) {
    response = &GetLLMDiagnosisDrugChatResponse{} 
    return

}

// GetLLMDiagnosisDrugChat
// 大模型问药问答
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) GetLLMDiagnosisDrugChat(request *GetLLMDiagnosisDrugChatRequest) (response *GetLLMDiagnosisDrugChatResponse, err error) {
    return c.GetLLMDiagnosisDrugChatWithContext(context.Background(), request)
}

// GetLLMDiagnosisDrugChat
// 大模型问药问答
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) GetLLMDiagnosisDrugChatWithContext(ctx context.Context, request *GetLLMDiagnosisDrugChatRequest) (response *GetLLMDiagnosisDrugChatResponse, err error) {
    if request == nil {
        request = NewGetLLMDiagnosisDrugChatRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ig", APIVersion, "GetLLMDiagnosisDrugChat")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetLLMDiagnosisDrugChat require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetLLMDiagnosisDrugChatResponse()
    err = c.Send(request, response)
    return
}

func NewGetLLMDiagnosisHealthRequest() (request *GetLLMDiagnosisHealthRequest) {
    request = &GetLLMDiagnosisHealthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ig", APIVersion, "GetLLMDiagnosisHealth")
    
    
    return
}

func NewGetLLMDiagnosisHealthResponse() (response *GetLLMDiagnosisHealthResponse) {
    response = &GetLLMDiagnosisHealthResponse{} 
    return

}

// GetLLMDiagnosisHealth
// 大模型健康自诊
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) GetLLMDiagnosisHealth(request *GetLLMDiagnosisHealthRequest) (response *GetLLMDiagnosisHealthResponse, err error) {
    return c.GetLLMDiagnosisHealthWithContext(context.Background(), request)
}

// GetLLMDiagnosisHealth
// 大模型健康自诊
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) GetLLMDiagnosisHealthWithContext(ctx context.Context, request *GetLLMDiagnosisHealthRequest) (response *GetLLMDiagnosisHealthResponse, err error) {
    if request == nil {
        request = NewGetLLMDiagnosisHealthRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ig", APIVersion, "GetLLMDiagnosisHealth")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetLLMDiagnosisHealth require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetLLMDiagnosisHealthResponse()
    err = c.Send(request, response)
    return
}

func NewGetLLMReportInterpretationRequest() (request *GetLLMReportInterpretationRequest) {
    request = &GetLLMReportInterpretationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ig", APIVersion, "GetLLMReportInterpretation")
    
    
    return
}

func NewGetLLMReportInterpretationResponse() (response *GetLLMReportInterpretationResponse) {
    response = &GetLLMReportInterpretationResponse{} 
    return

}

// GetLLMReportInterpretation
// 大模型报告解读
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) GetLLMReportInterpretation(request *GetLLMReportInterpretationRequest) (response *GetLLMReportInterpretationResponse, err error) {
    return c.GetLLMReportInterpretationWithContext(context.Background(), request)
}

// GetLLMReportInterpretation
// 大模型报告解读
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) GetLLMReportInterpretationWithContext(ctx context.Context, request *GetLLMReportInterpretationRequest) (response *GetLLMReportInterpretationResponse, err error) {
    if request == nil {
        request = NewGetLLMReportInterpretationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ig", APIVersion, "GetLLMReportInterpretation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetLLMReportInterpretation require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetLLMReportInterpretationResponse()
    err = c.Send(request, response)
    return
}

func NewQueryDrugInstructionsRequest() (request *QueryDrugInstructionsRequest) {
    request = &QueryDrugInstructionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ig", APIVersion, "QueryDrugInstructions")
    
    
    return
}

func NewQueryDrugInstructionsResponse() (response *QueryDrugInstructionsResponse) {
    response = &QueryDrugInstructionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryDrugInstructions
// 查询药品说明书
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) QueryDrugInstructions(request *QueryDrugInstructionsRequest) (response *QueryDrugInstructionsResponse, err error) {
    return c.QueryDrugInstructionsWithContext(context.Background(), request)
}

// QueryDrugInstructions
// 查询药品说明书
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) QueryDrugInstructionsWithContext(ctx context.Context, request *QueryDrugInstructionsRequest) (response *QueryDrugInstructionsResponse, err error) {
    if request == nil {
        request = NewQueryDrugInstructionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ig", APIVersion, "QueryDrugInstructions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryDrugInstructions require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryDrugInstructionsResponse()
    err = c.Send(request, response)
    return
}
