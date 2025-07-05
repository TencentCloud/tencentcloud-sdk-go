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

package v20220530

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2022-05-30"

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


func NewVerifyLicenseRequest() (request *VerifyLicenseRequest) {
    request = &VerifyLicenseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudapp", APIVersion, "VerifyLicense")
    
    
    return
}

func NewVerifyLicenseResponse() (response *VerifyLicenseResponse) {
    response = &VerifyLicenseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// VerifyLicense
// 从软件进程读取 LICENSE。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_FAILTOSIGN = "InternalError.FailToSign"
//  RESOURCENOTFOUND_LICENSENOTFOUNDERR = "ResourceNotFound.LicenseNotFoundErr"
func (c *Client) VerifyLicense(request *VerifyLicenseRequest) (response *VerifyLicenseResponse, err error) {
    return c.VerifyLicenseWithContext(context.Background(), request)
}

// VerifyLicense
// 从软件进程读取 LICENSE。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_FAILTOSIGN = "InternalError.FailToSign"
//  RESOURCENOTFOUND_LICENSENOTFOUNDERR = "ResourceNotFound.LicenseNotFoundErr"
func (c *Client) VerifyLicenseWithContext(ctx context.Context, request *VerifyLicenseRequest) (response *VerifyLicenseResponse, err error) {
    if request == nil {
        request = NewVerifyLicenseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("VerifyLicense require credential")
    }

    request.SetContext(ctx)
    
    response = NewVerifyLicenseResponse()
    err = c.Send(request, response)
    return
}
