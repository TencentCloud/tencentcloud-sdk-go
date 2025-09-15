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


func NewDescribeLicenseRequest() (request *DescribeLicenseRequest) {
    request = &DescribeLicenseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudapp", APIVersion, "DescribeLicense")
    
    
    return
}

func NewDescribeLicenseResponse() (response *DescribeLicenseResponse) {
    response = &DescribeLicenseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLicense
// # DescribeLicense
//
// 
//
// 客户根据请求的参数获取当前名下的许可证信息
//
// 
//
// ```json
//
// {
//
//   "Filters": [
//
//     {
//
//       "Name": "QueryType",
//
//       "Values": ["IncludeAddition"]
//
//     }
//
//   ]
//
// }
//
// ```
//
// 
//
// 返回的内容结构如下：
//
// 
//
// - Response.RequestId 为当前请求的唯一 id
//
// - Response.Token 为 license 信息 jwt 加密后的 token 串
//
// 
//
// ```json
//
// {
//
//   "Response": {
//
//     "RequestId": "cd15813b-adff-460e-b9fc-64579e96412d",
//
//     "Token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjkzMjQ4MTc5ODAsImlhdCI6MTc1NjE3Nzk4MCwiaXNzIjoibGljZW5zZS1zZXJ2aWNlIiwicGF5bG9hZCI6eyJNYWluTGljZW5zZSI6eyJMaWNlbnNlTW9kZSI6IlN1YnNjcmlwdGlvbiIsIkJpbGxpbmdNb2RlIjoxLCJDcmVhdGVTb3VyY2UiOiJTTjE3MTk0MDc1NDc0SEJETSIsIkF1dGhvcml6ZWRDbG91ZGFwcFJvbGVJZCI6IjcwMDAwMTgzMzgwNiIsIkF1dGhvcml6ZWRDbG91ZGFwcElkIjoiY2xvdWRhcHAtc2V3ZWM2cHMiLCJBdXRob3JpemVkVXNlclVpbiI6IjcwMDAwMDkxODE1NiIsIkxpZmVTcGFuVW5pdCI6IlkiLCJMaWZlU3BhbiI6MzY1LCJTb2Z0d2FyZVBhY2thZ2VJZCI6InBrZy0xZ2xlaG9tNyIsIlNvZnR3YXJlUGFja2FnZVZlcnNpb24iOiIwLjAuMSIsIkF1dGhvcml6ZWRTcGVjaWZpY2F0aW9uIjpbeyJQYXJhbUtleSI6InZlcnNpb24iLCJQYXJhbUtleU5hbWUiOiLniYjmnKwiLCJQYXJhbVZhbHVlIjoiYmFzaWMiLCJQYXJhbVZhbHVlTmFtZSI6IuWfuuehgOeJiCJ9LHsiUGFyYW1LZXkiOiJzaXplIiwiUGFyYW1LZXlOYW1lIjoi6KeE5qC8IiwiUGFyYW1WYWx1ZSI6IjEwMCIsIlBhcmFtVmFsdWVOYW1lIjoiMTAw5Lq66KeE5qihIn1dLCJQcm92aWRlcklkIjoxMDAwMDAwNzEsIlByb3ZpZGVyVWluIjoiNzAwMDAwOTE4MTU2IiwiSXNzdWVEYXRlIjoiMjAyNC0wNi0yNlQyMToxMjozMiswODowMCIsIkFjdGl2YXRpb25EYXRlIjoiMjAyNC0wNi0yNlQyMToxMjozNSswODowMCIsIkV4cGlyYXRpb25EYXRlIjoiMjM4OS0wNi0yNlQyMToxMjozNSswODowMCIsIkxpY2Vuc2VTdGF0dXMiOiJBY3RpdmUiLCJMaWNlbnNlSWQiOiI3MDAwMDA5MTgxNTY6cGtnLTFnbGVob203OmNsb3VkYXBwLXNld2VjNnBzOjgwMDciLCJMaWNlbnNlVHlwZSI6IlN0YW5kYXJkIiwiTGljZW5zZUxldmVsIjoiTWFzdGVyIn0sIkFkZGl0aW9uTGljZW5zZXMiOltdLCJUaW1lc3RhbXAiOiIyMDI1LTA4LTI2VDExOjEzOjAwKzA4OjAwIn19.G8Lx49xZBW0Rh3lRA15XzZ-PzLJj0bAxwnklx0pTjrHWxqxQdETAdGfU_QaGI_WZfYh2IVbFcwHnRLiRj6pQb4guCMpCbcsgL28BRS4g1wnaFhjcyEQLLtpDdz4_lPnOR2VHHvnfwhLZtccAgsRpeedPMBK1hwO9D3WKisQg2LcIr0V-QB8gmgIqqyqrLW6z37QpjgB4ZyJ5bIC1J-0-VmghskA04xnQRPdGJtlyBhjzVjeDxBq5JOqm3Am0Nqu1jyTd3MuYgSRwJqkDyjVBOGFGGy6mZCIYnxU_ET6-0ZEendqYwXDkpYG4rZZv5YmRCXiSESYz0zx4czwmFWkw-TjRSvUQBxBfsoDcAgyzpY7zBOTnbrr7DyoMvVnnHo7vb0if8_vkub6o0MuRnvdDYxNJtnTtlIScCadWAIvWUQ1DlUw2kzS-h9Ju2h7JhKw9cUeutu0X_6V4arZu9JlgWT9Ns7BtS9Y5JxgQOd36Aan39Rwohy_BrVwjOkbvDuTFLc_yNUlNdq5T2GNbDjABCmi73CGhCuWyPgtRs4ftpPugDRrTe4E95F224jdhf7I0He-nY4i1MoVjz8Zzm4v0vH67cMfcu0XVhs7ywvmu5tBSwm0uuhAXFFIbSrgEzuadxNhSi6qVCFNLnjiPYplK1M9mxG8Hc-fU-0A0TPepx8Q"
//
//   }
//
// }
//
// ```
//
// 
//
// 验签过程：
//
// 对 Response.Token 内容使用公钥进行解码转换得到许可结构体信息，返回的内容结构如下，其中 paylod 中的信息为许可证信息结构：
//
// 
//
// ```json
//
// {
//
//   "exp": 9324758169,
//
//   "iat": 1756118169,
//
//   "iss": "license-service",
//
//   "payload": {
//
//     "MainLicense": {
//
//       "LicenseMode": "Subscription",
//
//       "BillingMode": 1,
//
//       "CreateSource": "SN1719406931EJJ1E",
//
//       "AuthorizedCloudappRoleId": "700001833621",
//
//       "AuthorizedCloudappId": "cloudapp-992nqg9u",
//
//       "AuthorizedUserUin": "700001833621",
//
//       "LifeSpanUnit": "Y",
//
//       "LifeSpan": 365,
//
//       "SoftwarePackageId": "pkg-1glehom7",
//
//       "SoftwarePackageVersion": "0.0.1",
//
//       "AuthorizedSpecification": [
//
//         {
//
//           "ParamKey": "version",
//
//           "ParamKeyName": "版本",
//
//           "ParamValue": "basic",
//
//           "ParamValueName": "基础版"
//
//         },
//
//         {
//
//           "ParamKey": "size",
//
//           "ParamKeyName": "规格",
//
//           "ParamValue": "100",
//
//           "ParamValueName": "100 人规模"
//
//         }
//
//       ],
//
//       "ProviderId": 100000071,
//
//       "ProviderUin": "700000918156",
//
//       "IssueDate": "2024-06-26T21:02:16+08:00",
//
//       "ActivationDate": "2024-06-26T21:02:19+08:00",
//
//       "ExpirationDate": "2389-06-26T21:02:19+08:00",
//
//       "LicenseStatus": "Active",
//
//       "LicenseId": "700000918156:pkg-1glehom7:cloudapp-992nqg9u:3988",
//
//       "LicenseType": "Standard",
//
//       "LicenseLevel": "Master"
//
//     },
//
//     "AdditionLicenses": [
//
//       {
//
//         "LicenseMode": "Subscription",
//
//         "BillingMode": 1,
//
//         "CreateSource": "SN1719406931EJJ1E",
//
//         "AuthorizedCloudappRoleId": "700001833621",
//
//         "AuthorizedCloudappId": "cloudapp-992nqg9u",
//
//         "AuthorizedUserUin": "700001833621",
//
//         "LifeSpanUnit": "Y",
//
//         "LifeSpan": 365,
//
//         "SoftwarePackageId": "pkg-1glehom7",
//
//         "SoftwarePackageVersion": "0.0.1",
//
//         "AuthorizedSpecification": [
//
//           {
//
//             "ParamKey": "version",
//
//             "ParamKeyName": "版本",
//
//             "ParamValue": "basic",
//
//             "ParamValueName": "基础版"
//
//           },
//
//           {
//
//             "ParamKey": "size",
//
//             "ParamKeyName": "规格",
//
//             "ParamValue": "100",
//
//             "ParamValueName": "100 人规模"
//
//           }
//
//         ],
//
//         "ProviderId": 100000071,
//
//         "ProviderUin": "700000918156",
//
//         "IssueDate": "2024-06-26T21:02:16+08:00",
//
//         "ActivationDate": "2024-06-26T21:02:19+08:00",
//
//         "ExpirationDate": "2389-06-26T21:02:19+08:00",
//
//         "LicenseStatus": "Active",
//
//         "LicenseId": "700000918156:pkg-1glehom7:cloudapp-992nqg9u:3988",
//
//         "LicenseType": "Standard",
//
//         "LicenseLevel": "Master"
//
//       }
//
//     ],
//
//     "Timestamp": "2025-08-25T18:36:09+08:00"
//
//   }
//
// }
//
// ```
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_FAILTOSIGN = "InternalError.FailToSign"
//  RESOURCENOTFOUND_LICENSENOTFOUNDERR = "ResourceNotFound.LicenseNotFoundErr"
func (c *Client) DescribeLicense(request *DescribeLicenseRequest) (response *DescribeLicenseResponse, err error) {
    return c.DescribeLicenseWithContext(context.Background(), request)
}

// DescribeLicense
// # DescribeLicense
//
// 
//
// 客户根据请求的参数获取当前名下的许可证信息
//
// 
//
// ```json
//
// {
//
//   "Filters": [
//
//     {
//
//       "Name": "QueryType",
//
//       "Values": ["IncludeAddition"]
//
//     }
//
//   ]
//
// }
//
// ```
//
// 
//
// 返回的内容结构如下：
//
// 
//
// - Response.RequestId 为当前请求的唯一 id
//
// - Response.Token 为 license 信息 jwt 加密后的 token 串
//
// 
//
// ```json
//
// {
//
//   "Response": {
//
//     "RequestId": "cd15813b-adff-460e-b9fc-64579e96412d",
//
//     "Token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjkzMjQ4MTc5ODAsImlhdCI6MTc1NjE3Nzk4MCwiaXNzIjoibGljZW5zZS1zZXJ2aWNlIiwicGF5bG9hZCI6eyJNYWluTGljZW5zZSI6eyJMaWNlbnNlTW9kZSI6IlN1YnNjcmlwdGlvbiIsIkJpbGxpbmdNb2RlIjoxLCJDcmVhdGVTb3VyY2UiOiJTTjE3MTk0MDc1NDc0SEJETSIsIkF1dGhvcml6ZWRDbG91ZGFwcFJvbGVJZCI6IjcwMDAwMTgzMzgwNiIsIkF1dGhvcml6ZWRDbG91ZGFwcElkIjoiY2xvdWRhcHAtc2V3ZWM2cHMiLCJBdXRob3JpemVkVXNlclVpbiI6IjcwMDAwMDkxODE1NiIsIkxpZmVTcGFuVW5pdCI6IlkiLCJMaWZlU3BhbiI6MzY1LCJTb2Z0d2FyZVBhY2thZ2VJZCI6InBrZy0xZ2xlaG9tNyIsIlNvZnR3YXJlUGFja2FnZVZlcnNpb24iOiIwLjAuMSIsIkF1dGhvcml6ZWRTcGVjaWZpY2F0aW9uIjpbeyJQYXJhbUtleSI6InZlcnNpb24iLCJQYXJhbUtleU5hbWUiOiLniYjmnKwiLCJQYXJhbVZhbHVlIjoiYmFzaWMiLCJQYXJhbVZhbHVlTmFtZSI6IuWfuuehgOeJiCJ9LHsiUGFyYW1LZXkiOiJzaXplIiwiUGFyYW1LZXlOYW1lIjoi6KeE5qC8IiwiUGFyYW1WYWx1ZSI6IjEwMCIsIlBhcmFtVmFsdWVOYW1lIjoiMTAw5Lq66KeE5qihIn1dLCJQcm92aWRlcklkIjoxMDAwMDAwNzEsIlByb3ZpZGVyVWluIjoiNzAwMDAwOTE4MTU2IiwiSXNzdWVEYXRlIjoiMjAyNC0wNi0yNlQyMToxMjozMiswODowMCIsIkFjdGl2YXRpb25EYXRlIjoiMjAyNC0wNi0yNlQyMToxMjozNSswODowMCIsIkV4cGlyYXRpb25EYXRlIjoiMjM4OS0wNi0yNlQyMToxMjozNSswODowMCIsIkxpY2Vuc2VTdGF0dXMiOiJBY3RpdmUiLCJMaWNlbnNlSWQiOiI3MDAwMDA5MTgxNTY6cGtnLTFnbGVob203OmNsb3VkYXBwLXNld2VjNnBzOjgwMDciLCJMaWNlbnNlVHlwZSI6IlN0YW5kYXJkIiwiTGljZW5zZUxldmVsIjoiTWFzdGVyIn0sIkFkZGl0aW9uTGljZW5zZXMiOltdLCJUaW1lc3RhbXAiOiIyMDI1LTA4LTI2VDExOjEzOjAwKzA4OjAwIn19.G8Lx49xZBW0Rh3lRA15XzZ-PzLJj0bAxwnklx0pTjrHWxqxQdETAdGfU_QaGI_WZfYh2IVbFcwHnRLiRj6pQb4guCMpCbcsgL28BRS4g1wnaFhjcyEQLLtpDdz4_lPnOR2VHHvnfwhLZtccAgsRpeedPMBK1hwO9D3WKisQg2LcIr0V-QB8gmgIqqyqrLW6z37QpjgB4ZyJ5bIC1J-0-VmghskA04xnQRPdGJtlyBhjzVjeDxBq5JOqm3Am0Nqu1jyTd3MuYgSRwJqkDyjVBOGFGGy6mZCIYnxU_ET6-0ZEendqYwXDkpYG4rZZv5YmRCXiSESYz0zx4czwmFWkw-TjRSvUQBxBfsoDcAgyzpY7zBOTnbrr7DyoMvVnnHo7vb0if8_vkub6o0MuRnvdDYxNJtnTtlIScCadWAIvWUQ1DlUw2kzS-h9Ju2h7JhKw9cUeutu0X_6V4arZu9JlgWT9Ns7BtS9Y5JxgQOd36Aan39Rwohy_BrVwjOkbvDuTFLc_yNUlNdq5T2GNbDjABCmi73CGhCuWyPgtRs4ftpPugDRrTe4E95F224jdhf7I0He-nY4i1MoVjz8Zzm4v0vH67cMfcu0XVhs7ywvmu5tBSwm0uuhAXFFIbSrgEzuadxNhSi6qVCFNLnjiPYplK1M9mxG8Hc-fU-0A0TPepx8Q"
//
//   }
//
// }
//
// ```
//
// 
//
// 验签过程：
//
// 对 Response.Token 内容使用公钥进行解码转换得到许可结构体信息，返回的内容结构如下，其中 paylod 中的信息为许可证信息结构：
//
// 
//
// ```json
//
// {
//
//   "exp": 9324758169,
//
//   "iat": 1756118169,
//
//   "iss": "license-service",
//
//   "payload": {
//
//     "MainLicense": {
//
//       "LicenseMode": "Subscription",
//
//       "BillingMode": 1,
//
//       "CreateSource": "SN1719406931EJJ1E",
//
//       "AuthorizedCloudappRoleId": "700001833621",
//
//       "AuthorizedCloudappId": "cloudapp-992nqg9u",
//
//       "AuthorizedUserUin": "700001833621",
//
//       "LifeSpanUnit": "Y",
//
//       "LifeSpan": 365,
//
//       "SoftwarePackageId": "pkg-1glehom7",
//
//       "SoftwarePackageVersion": "0.0.1",
//
//       "AuthorizedSpecification": [
//
//         {
//
//           "ParamKey": "version",
//
//           "ParamKeyName": "版本",
//
//           "ParamValue": "basic",
//
//           "ParamValueName": "基础版"
//
//         },
//
//         {
//
//           "ParamKey": "size",
//
//           "ParamKeyName": "规格",
//
//           "ParamValue": "100",
//
//           "ParamValueName": "100 人规模"
//
//         }
//
//       ],
//
//       "ProviderId": 100000071,
//
//       "ProviderUin": "700000918156",
//
//       "IssueDate": "2024-06-26T21:02:16+08:00",
//
//       "ActivationDate": "2024-06-26T21:02:19+08:00",
//
//       "ExpirationDate": "2389-06-26T21:02:19+08:00",
//
//       "LicenseStatus": "Active",
//
//       "LicenseId": "700000918156:pkg-1glehom7:cloudapp-992nqg9u:3988",
//
//       "LicenseType": "Standard",
//
//       "LicenseLevel": "Master"
//
//     },
//
//     "AdditionLicenses": [
//
//       {
//
//         "LicenseMode": "Subscription",
//
//         "BillingMode": 1,
//
//         "CreateSource": "SN1719406931EJJ1E",
//
//         "AuthorizedCloudappRoleId": "700001833621",
//
//         "AuthorizedCloudappId": "cloudapp-992nqg9u",
//
//         "AuthorizedUserUin": "700001833621",
//
//         "LifeSpanUnit": "Y",
//
//         "LifeSpan": 365,
//
//         "SoftwarePackageId": "pkg-1glehom7",
//
//         "SoftwarePackageVersion": "0.0.1",
//
//         "AuthorizedSpecification": [
//
//           {
//
//             "ParamKey": "version",
//
//             "ParamKeyName": "版本",
//
//             "ParamValue": "basic",
//
//             "ParamValueName": "基础版"
//
//           },
//
//           {
//
//             "ParamKey": "size",
//
//             "ParamKeyName": "规格",
//
//             "ParamValue": "100",
//
//             "ParamValueName": "100 人规模"
//
//           }
//
//         ],
//
//         "ProviderId": 100000071,
//
//         "ProviderUin": "700000918156",
//
//         "IssueDate": "2024-06-26T21:02:16+08:00",
//
//         "ActivationDate": "2024-06-26T21:02:19+08:00",
//
//         "ExpirationDate": "2389-06-26T21:02:19+08:00",
//
//         "LicenseStatus": "Active",
//
//         "LicenseId": "700000918156:pkg-1glehom7:cloudapp-992nqg9u:3988",
//
//         "LicenseType": "Standard",
//
//         "LicenseLevel": "Master"
//
//       }
//
//     ],
//
//     "Timestamp": "2025-08-25T18:36:09+08:00"
//
//   }
//
// }
//
// ```
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_FAILTOSIGN = "InternalError.FailToSign"
//  RESOURCENOTFOUND_LICENSENOTFOUNDERR = "ResourceNotFound.LicenseNotFoundErr"
func (c *Client) DescribeLicenseWithContext(ctx context.Context, request *DescribeLicenseRequest) (response *DescribeLicenseResponse, err error) {
    if request == nil {
        request = NewDescribeLicenseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cloudapp", APIVersion, "DescribeLicense")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLicense require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLicenseResponse()
    err = c.Send(request, response)
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
    c.InitBaseRequest(&request.BaseRequest, "cloudapp", APIVersion, "VerifyLicense")
    
    if c.GetCredential() == nil {
        return nil, errors.New("VerifyLicense require credential")
    }

    request.SetContext(ctx)
    
    response = NewVerifyLicenseResponse()
    err = c.Send(request, response)
    return
}
