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

package v20220420

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2022-04-20"

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


func NewCreateBlockNodeRecordsRequest() (request *CreateBlockNodeRecordsRequest) {
    request = &CreateBlockNodeRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tan", APIVersion, "CreateBlockNodeRecords")
    
    
    return
}

func NewCreateBlockNodeRecordsResponse() (response *CreateBlockNodeRecordsResponse) {
    response = &CreateBlockNodeRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateBlockNodeRecords
// 推送节点数据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_GROUPNAMEEXISTED = "InvalidParameter.GroupNameExisted"
//  INVALIDPARAMETER_RECORDEXCEEDSLIMIT = "InvalidParameter.RecordExceedsLimit"
//  INVALIDPARAMETER_RECORDPARAMETERCHECKFAIL = "InvalidParameter.RecordParameterCheckFail"
//  INVALIDPARAMETER_RECORDPARAMETERPARSEFAIL = "InvalidParameter.RecordParameterParseFail"
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
func (c *Client) CreateBlockNodeRecords(request *CreateBlockNodeRecordsRequest) (response *CreateBlockNodeRecordsResponse, err error) {
    return c.CreateBlockNodeRecordsWithContext(context.Background(), request)
}

// CreateBlockNodeRecords
// 推送节点数据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_GROUPNAMEEXISTED = "InvalidParameter.GroupNameExisted"
//  INVALIDPARAMETER_RECORDEXCEEDSLIMIT = "InvalidParameter.RecordExceedsLimit"
//  INVALIDPARAMETER_RECORDPARAMETERCHECKFAIL = "InvalidParameter.RecordParameterCheckFail"
//  INVALIDPARAMETER_RECORDPARAMETERPARSEFAIL = "InvalidParameter.RecordParameterParseFail"
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
func (c *Client) CreateBlockNodeRecordsWithContext(ctx context.Context, request *CreateBlockNodeRecordsRequest) (response *CreateBlockNodeRecordsResponse, err error) {
    if request == nil {
        request = NewCreateBlockNodeRecordsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBlockNodeRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBlockNodeRecordsResponse()
    err = c.Send(request, response)
    return
}
