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

package v20190926

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-09-26"

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


func NewMarketingValueJudgementRequest() (request *MarketingValueJudgementRequest) {
    request = &MarketingValueJudgementRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mvj", APIVersion, "MarketingValueJudgement")
    
    
    return
}

func NewMarketingValueJudgementResponse() (response *MarketingValueJudgementResponse) {
    response = &MarketingValueJudgementResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// MarketingValueJudgement
// 欢迎使用营销价值判断（Marketing Value Judgement，简称 MVJ）。
//
// 
//
// 营销价值判断（MVJ）是针对零售场景的风控服务，通过识别高价值顾客，以帮助零售商保障营销资金
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRTIMEOUT = "InternalError.ErrTimeOut"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTFOUND = "InvalidParameterValue.AccountNotFound"
//  INVALIDPARAMETERVALUE_INVALIDIP = "InvalidParameterValue.InvalidIp"
//  INVALIDPARAMETERVALUE_INVALIDMOBILENUMBER = "InvalidParameterValue.InvalidMobileNumber"
func (c *Client) MarketingValueJudgement(request *MarketingValueJudgementRequest) (response *MarketingValueJudgementResponse, err error) {
    return c.MarketingValueJudgementWithContext(context.Background(), request)
}

// MarketingValueJudgement
// 欢迎使用营销价值判断（Marketing Value Judgement，简称 MVJ）。
//
// 
//
// 营销价值判断（MVJ）是针对零售场景的风控服务，通过识别高价值顾客，以帮助零售商保障营销资金
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRTIMEOUT = "InternalError.ErrTimeOut"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTFOUND = "InvalidParameterValue.AccountNotFound"
//  INVALIDPARAMETERVALUE_INVALIDIP = "InvalidParameterValue.InvalidIp"
//  INVALIDPARAMETERVALUE_INVALIDMOBILENUMBER = "InvalidParameterValue.InvalidMobileNumber"
func (c *Client) MarketingValueJudgementWithContext(ctx context.Context, request *MarketingValueJudgementRequest) (response *MarketingValueJudgementResponse, err error) {
    if request == nil {
        request = NewMarketingValueJudgementRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("MarketingValueJudgement require credential")
    }

    request.SetContext(ctx)
    
    response = NewMarketingValueJudgementResponse()
    err = c.Send(request, response)
    return
}
