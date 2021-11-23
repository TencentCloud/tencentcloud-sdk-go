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

package v20210408

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-04-08"

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


func NewCreateStructureTaskRequest() (request *CreateStructureTaskRequest) {
    request = &CreateStructureTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cii", APIVersion, "CreateStructureTask")
    
    return
}

func NewCreateStructureTaskResponse() (response *CreateStructureTaskResponse) {
    response = &CreateStructureTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateStructureTask
// 本接口(CreateStructureTask)基于提供的客户及保单信息，创建并启动结构化识别任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateStructureTask(request *CreateStructureTaskRequest) (response *CreateStructureTaskResponse, err error) {
    if request == nil {
        request = NewCreateStructureTaskRequest()
    }
    response = NewCreateStructureTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUnderwriteTaskByIdRequest() (request *CreateUnderwriteTaskByIdRequest) {
    request = &CreateUnderwriteTaskByIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cii", APIVersion, "CreateUnderwriteTaskById")
    
    return
}

func NewCreateUnderwriteTaskByIdResponse() (response *CreateUnderwriteTaskByIdResponse) {
    response = &CreateUnderwriteTaskByIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateUnderwriteTaskById
// 本接口(CreateUnderwriteTaskById)用于根据结构化任务ID创建核保任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateUnderwriteTaskById(request *CreateUnderwriteTaskByIdRequest) (response *CreateUnderwriteTaskByIdResponse, err error) {
    if request == nil {
        request = NewCreateUnderwriteTaskByIdRequest()
    }
    response = NewCreateUnderwriteTaskByIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMachineUnderwriteRequest() (request *DescribeMachineUnderwriteRequest) {
    request = &DescribeMachineUnderwriteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cii", APIVersion, "DescribeMachineUnderwrite")
    
    return
}

func NewDescribeMachineUnderwriteResponse() (response *DescribeMachineUnderwriteResponse) {
    response = &DescribeMachineUnderwriteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMachineUnderwrite
// 本接口(DescribeMachineUnderwrite)用于查询机器核保任务数据
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DescribeMachineUnderwrite(request *DescribeMachineUnderwriteRequest) (response *DescribeMachineUnderwriteResponse, err error) {
    if request == nil {
        request = NewDescribeMachineUnderwriteRequest()
    }
    response = NewDescribeMachineUnderwriteResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStructCompareDataRequest() (request *DescribeStructCompareDataRequest) {
    request = &DescribeStructCompareDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cii", APIVersion, "DescribeStructCompareData")
    
    return
}

func NewDescribeStructCompareDataResponse() (response *DescribeStructCompareDataResponse) {
    response = &DescribeStructCompareDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStructCompareData
// 结构化对比查询接口，对比结构化复核前后数据差异，查询识别正确率，召回率。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeStructCompareData(request *DescribeStructCompareDataRequest) (response *DescribeStructCompareDataResponse, err error) {
    if request == nil {
        request = NewDescribeStructCompareDataRequest()
    }
    response = NewDescribeStructCompareDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStructureDifferenceRequest() (request *DescribeStructureDifferenceRequest) {
    request = &DescribeStructureDifferenceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cii", APIVersion, "DescribeStructureDifference")
    
    return
}

func NewDescribeStructureDifferenceResponse() (response *DescribeStructureDifferenceResponse) {
    response = &DescribeStructureDifferenceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStructureDifference
// 结构化复核差异查询接口，对比结构化复核前后数据差异，返回差异的部分。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeStructureDifference(request *DescribeStructureDifferenceRequest) (response *DescribeStructureDifferenceResponse, err error) {
    if request == nil {
        request = NewDescribeStructureDifferenceRequest()
    }
    response = NewDescribeStructureDifferenceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStructureResultRequest() (request *DescribeStructureResultRequest) {
    request = &DescribeStructureResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cii", APIVersion, "DescribeStructureResult")
    
    return
}

func NewDescribeStructureResultResponse() (response *DescribeStructureResultResponse) {
    response = &DescribeStructureResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStructureResult
// 本接口(DescribeStructureResult)用于查询结构化结果接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeStructureResult(request *DescribeStructureResultRequest) (response *DescribeStructureResultResponse, err error) {
    if request == nil {
        request = NewDescribeStructureResultRequest()
    }
    response = NewDescribeStructureResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStructureTaskResultRequest() (request *DescribeStructureTaskResultRequest) {
    request = &DescribeStructureTaskResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cii", APIVersion, "DescribeStructureTaskResult")
    
    return
}

func NewDescribeStructureTaskResultResponse() (response *DescribeStructureTaskResultResponse) {
    response = &DescribeStructureTaskResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStructureTaskResult
// 依据任务ID获取结构化结果接口。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeStructureTaskResult(request *DescribeStructureTaskResultRequest) (response *DescribeStructureTaskResultResponse, err error) {
    if request == nil {
        request = NewDescribeStructureTaskResultRequest()
    }
    response = NewDescribeStructureTaskResultResponse()
    err = c.Send(request, response)
    return
}

func NewUploadMedicalFileRequest() (request *UploadMedicalFileRequest) {
    request = &UploadMedicalFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cii", APIVersion, "UploadMedicalFile")
    
    return
}

func NewUploadMedicalFileResponse() (response *UploadMedicalFileResponse) {
    response = &UploadMedicalFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UploadMedicalFile
// 上传医疗影像文件，可以用来做结构化。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_ENTERPRISE = "AuthFailure.Enterprise"
//  AUTHFAILURE_PERSONAL = "AuthFailure.Personal"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) UploadMedicalFile(request *UploadMedicalFileRequest) (response *UploadMedicalFileResponse, err error) {
    if request == nil {
        request = NewUploadMedicalFileRequest()
    }
    response = NewUploadMedicalFileResponse()
    err = c.Send(request, response)
    return
}
