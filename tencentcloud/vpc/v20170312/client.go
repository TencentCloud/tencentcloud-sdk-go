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

package v20170312

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2017-03-12"

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


func NewDescribeNatsEipsInternalRequest() (request *DescribeNatsEipsInternalRequest) {
    request = &DescribeNatsEipsInternalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeNatsEipsInternal")
    
    
    return
}

func NewDescribeNatsEipsInternalResponse() (response *DescribeNatsEipsInternalResponse) {
    response = &DescribeNatsEipsInternalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNatsEipsInternal
// 内部用户查询所有NAT网关的EIP IP列表信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLVPCOSSFAILED = "FailedOperation.CallVpcOssFailed"
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETER_FILTERVALUESNOTLIST = "InvalidParameter.FilterValuesNotList"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeNatsEipsInternal(request *DescribeNatsEipsInternalRequest) (response *DescribeNatsEipsInternalResponse, err error) {
    return c.DescribeNatsEipsInternalWithContext(context.Background(), request)
}

// DescribeNatsEipsInternal
// 内部用户查询所有NAT网关的EIP IP列表信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLVPCOSSFAILED = "FailedOperation.CallVpcOssFailed"
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETER_FILTERVALUESNOTLIST = "InvalidParameter.FilterValuesNotList"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeNatsEipsInternalWithContext(ctx context.Context, request *DescribeNatsEipsInternalRequest) (response *DescribeNatsEipsInternalResponse, err error) {
    if request == nil {
        request = NewDescribeNatsEipsInternalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNatsEipsInternal require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNatsEipsInternalResponse()
    err = c.Send(request, response)
    return
}
