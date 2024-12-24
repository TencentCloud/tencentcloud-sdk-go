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

package v20210323

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-03-23"

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


func NewGetDrugIndicationsRequest() (request *GetDrugIndicationsRequest) {
    request = &GetDrugIndicationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("aca", APIVersion, "GetDrugIndications")
    
    
    return
}

func NewGetDrugIndicationsResponse() (response *GetDrugIndicationsResponse) {
    response = &GetDrugIndicationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetDrugIndications
// 药品适应症接口
func (c *Client) GetDrugIndications(request *GetDrugIndicationsRequest) (response *GetDrugIndicationsResponse, err error) {
    return c.GetDrugIndicationsWithContext(context.Background(), request)
}

// GetDrugIndications
// 药品适应症接口
func (c *Client) GetDrugIndicationsWithContext(ctx context.Context, request *GetDrugIndicationsRequest) (response *GetDrugIndicationsResponse, err error) {
    if request == nil {
        request = NewGetDrugIndicationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDrugIndications require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDrugIndicationsResponse()
    err = c.Send(request, response)
    return
}

func NewLoginHisToolRequest() (request *LoginHisToolRequest) {
    request = &LoginHisToolRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("aca", APIVersion, "LoginHisTool")
    
    
    return
}

func NewLoginHisToolResponse() (response *LoginHisToolResponse) {
    response = &LoginHisToolResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// LoginHisTool
// 登录获取token
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) LoginHisTool(request *LoginHisToolRequest) (response *LoginHisToolResponse, err error) {
    return c.LoginHisToolWithContext(context.Background(), request)
}

// LoginHisTool
// 登录获取token
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) LoginHisToolWithContext(ctx context.Context, request *LoginHisToolRequest) (response *LoginHisToolResponse, err error) {
    if request == nil {
        request = NewLoginHisToolRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("LoginHisTool require credential")
    }

    request.SetContext(ctx)
    
    response = NewLoginHisToolResponse()
    err = c.Send(request, response)
    return
}

func NewLoginOutHisToolRequest() (request *LoginOutHisToolRequest) {
    request = &LoginOutHisToolRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("aca", APIVersion, "LoginOutHisTool")
    
    
    return
}

func NewLoginOutHisToolResponse() (response *LoginOutHisToolResponse) {
    response = &LoginOutHisToolResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// LoginOutHisTool
// 登出
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) LoginOutHisTool(request *LoginOutHisToolRequest) (response *LoginOutHisToolResponse, err error) {
    return c.LoginOutHisToolWithContext(context.Background(), request)
}

// LoginOutHisTool
// 登出
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) LoginOutHisToolWithContext(ctx context.Context, request *LoginOutHisToolRequest) (response *LoginOutHisToolResponse, err error) {
    if request == nil {
        request = NewLoginOutHisToolRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("LoginOutHisTool require credential")
    }

    request.SetContext(ctx)
    
    response = NewLoginOutHisToolResponse()
    err = c.Send(request, response)
    return
}

func NewSmartDrugInfoRequest() (request *SmartDrugInfoRequest) {
    request = &SmartDrugInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("aca", APIVersion, "SmartDrugInfo")
    
    
    return
}

func NewSmartDrugInfoResponse() (response *SmartDrugInfoResponse) {
    response = &SmartDrugInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SmartDrugInfo
// 智能用药接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) SmartDrugInfo(request *SmartDrugInfoRequest) (response *SmartDrugInfoResponse, err error) {
    return c.SmartDrugInfoWithContext(context.Background(), request)
}

// SmartDrugInfo
// 智能用药接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) SmartDrugInfoWithContext(ctx context.Context, request *SmartDrugInfoRequest) (response *SmartDrugInfoResponse, err error) {
    if request == nil {
        request = NewSmartDrugInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SmartDrugInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewSmartDrugInfoResponse()
    err = c.Send(request, response)
    return
}

func NewSmartPredictRequest() (request *SmartPredictRequest) {
    request = &SmartPredictRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("aca", APIVersion, "SmartPredict")
    
    
    return
}

func NewSmartPredictResponse() (response *SmartPredictResponse) {
    response = &SmartPredictResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SmartPredict
// 辅诊智能预测接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) SmartPredict(request *SmartPredictRequest) (response *SmartPredictResponse, err error) {
    return c.SmartPredictWithContext(context.Background(), request)
}

// SmartPredict
// 辅诊智能预测接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) SmartPredictWithContext(ctx context.Context, request *SmartPredictRequest) (response *SmartPredictResponse, err error) {
    if request == nil {
        request = NewSmartPredictRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SmartPredict require credential")
    }

    request.SetContext(ctx)
    
    response = NewSmartPredictResponse()
    err = c.Send(request, response)
    return
}

func NewSyncDepartmentRequest() (request *SyncDepartmentRequest) {
    request = &SyncDepartmentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("aca", APIVersion, "SyncDepartment")
    
    
    return
}

func NewSyncDepartmentResponse() (response *SyncDepartmentResponse) {
    response = &SyncDepartmentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SyncDepartment
// 用于院方科室管理，获取科室列表和状态、新增或修改科室信息、删除科室。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) SyncDepartment(request *SyncDepartmentRequest) (response *SyncDepartmentResponse, err error) {
    return c.SyncDepartmentWithContext(context.Background(), request)
}

// SyncDepartment
// 用于院方科室管理，获取科室列表和状态、新增或修改科室信息、删除科室。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) SyncDepartmentWithContext(ctx context.Context, request *SyncDepartmentRequest) (response *SyncDepartmentResponse, err error) {
    if request == nil {
        request = NewSyncDepartmentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SyncDepartment require credential")
    }

    request.SetContext(ctx)
    
    response = NewSyncDepartmentResponse()
    err = c.Send(request, response)
    return
}

func NewSyncStandardDictRequest() (request *SyncStandardDictRequest) {
    request = &SyncStandardDictRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("aca", APIVersion, "SyncStandardDict")
    
    
    return
}

func NewSyncStandardDictResponse() (response *SyncStandardDictResponse) {
    response = &SyncStandardDictResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SyncStandardDict
// 同步标准字典，如给药频次、给药途径、科室、诊断等
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) SyncStandardDict(request *SyncStandardDictRequest) (response *SyncStandardDictResponse, err error) {
    return c.SyncStandardDictWithContext(context.Background(), request)
}

// SyncStandardDict
// 同步标准字典，如给药频次、给药途径、科室、诊断等
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) SyncStandardDictWithContext(ctx context.Context, request *SyncStandardDictRequest) (response *SyncStandardDictResponse, err error) {
    if request == nil {
        request = NewSyncStandardDictRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SyncStandardDict require credential")
    }

    request.SetContext(ctx)
    
    response = NewSyncStandardDictResponse()
    err = c.Send(request, response)
    return
}

func NewUploadDrugsRequest() (request *UploadDrugsRequest) {
    request = &UploadDrugsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("aca", APIVersion, "UploadDrugs")
    
    
    return
}

func NewUploadDrugsResponse() (response *UploadDrugsResponse) {
    response = &UploadDrugsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UploadDrugs
// 药品同步，一次同步数据不要超过500个
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UploadDrugs(request *UploadDrugsRequest) (response *UploadDrugsResponse, err error) {
    return c.UploadDrugsWithContext(context.Background(), request)
}

// UploadDrugs
// 药品同步，一次同步数据不要超过500个
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UploadDrugsWithContext(ctx context.Context, request *UploadDrugsRequest) (response *UploadDrugsResponse, err error) {
    if request == nil {
        request = NewUploadDrugsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadDrugs require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadDrugsResponse()
    err = c.Send(request, response)
    return
}
