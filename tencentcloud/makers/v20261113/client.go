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

package v20261113

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2026-11-13"

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


func NewCreateDeploymentRequest() (request *CreateDeploymentRequest) {
    request = &CreateDeploymentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("makers", APIVersion, "CreateDeployment")
    
    
    return
}

func NewCreateDeploymentResponse() (response *CreateDeploymentResponse) {
    response = &CreateDeploymentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDeployment
// 创建部署。上传制品时 ViaMeta 和 Provider 传 Upload，TempBucketPath 传已上传对象的路径。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GITAUTHEXPIRED = "FailedOperation.GitAuthExpired"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateDeployment(request *CreateDeploymentRequest) (response *CreateDeploymentResponse, err error) {
    return c.CreateDeploymentWithContext(context.Background(), request)
}

// CreateDeployment
// 创建部署。上传制品时 ViaMeta 和 Provider 传 Upload，TempBucketPath 传已上传对象的路径。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GITAUTHEXPIRED = "FailedOperation.GitAuthExpired"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateDeploymentWithContext(ctx context.Context, request *CreateDeploymentRequest) (response *CreateDeploymentResponse, err error) {
    if request == nil {
        request = NewCreateDeploymentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "makers", APIVersion, "CreateDeployment")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDeployment require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDeploymentResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProjectRequest() (request *CreateProjectRequest) {
    request = &CreateProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("makers", APIVersion, "CreateProject")
    
    
    return
}

func NewCreateProjectResponse() (response *CreateProjectResponse) {
    response = &CreateProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateProject
// 创建项目。上传制品时 Provider 传 Upload，Channel 传 Custom。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateProject(request *CreateProjectRequest) (response *CreateProjectResponse, err error) {
    return c.CreateProjectWithContext(context.Background(), request)
}

// CreateProject
// 创建项目。上传制品时 Provider 传 Upload，Channel 传 Custom。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateProjectWithContext(ctx context.Context, request *CreateProjectRequest) (response *CreateProjectResponse, err error) {
    if request == nil {
        request = NewCreateProjectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "makers", APIVersion, "CreateProject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateProjectResponse()
    err = c.Send(request, response)
    return
}

func NewCreateZoneCustomDomainRequest() (request *CreateZoneCustomDomainRequest) {
    request = &CreateZoneCustomDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("makers", APIVersion, "CreateZoneCustomDomain")
    
    
    return
}

func NewCreateZoneCustomDomainResponse() (response *CreateZoneCustomDomainResponse) {
    response = &CreateZoneCustomDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateZoneCustomDomain
// 把一个自定义域名绑到项目的指定环境。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateZoneCustomDomain(request *CreateZoneCustomDomainRequest) (response *CreateZoneCustomDomainResponse, err error) {
    return c.CreateZoneCustomDomainWithContext(context.Background(), request)
}

// CreateZoneCustomDomain
// 把一个自定义域名绑到项目的指定环境。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateZoneCustomDomainWithContext(ctx context.Context, request *CreateZoneCustomDomainRequest) (response *CreateZoneCustomDomainResponse, err error) {
    if request == nil {
        request = NewCreateZoneCustomDomainRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "makers", APIVersion, "CreateZoneCustomDomain")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateZoneCustomDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateZoneCustomDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteProjectRequest() (request *DeleteProjectRequest) {
    request = &DeleteProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("makers", APIVersion, "DeleteProject")
    
    
    return
}

func NewDeleteProjectResponse() (response *DeleteProjectResponse) {
    response = &DeleteProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteProject
// 删除项目。进行中的构建会阻止删除。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteProject(request *DeleteProjectRequest) (response *DeleteProjectResponse, err error) {
    return c.DeleteProjectWithContext(context.Background(), request)
}

// DeleteProject
// 删除项目。进行中的构建会阻止删除。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteProjectWithContext(ctx context.Context, request *DeleteProjectRequest) (response *DeleteProjectResponse, err error) {
    if request == nil {
        request = NewDeleteProjectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "makers", APIVersion, "DeleteProject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteProjectResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteProjectCustomDomainsRequest() (request *DeleteProjectCustomDomainsRequest) {
    request = &DeleteProjectCustomDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("makers", APIVersion, "DeleteProjectCustomDomains")
    
    
    return
}

func NewDeleteProjectCustomDomainsResponse() (response *DeleteProjectCustomDomainsResponse) {
    response = &DeleteProjectCustomDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteProjectCustomDomains
// 按域名解绑项目的自定义域名。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteProjectCustomDomains(request *DeleteProjectCustomDomainsRequest) (response *DeleteProjectCustomDomainsResponse, err error) {
    return c.DeleteProjectCustomDomainsWithContext(context.Background(), request)
}

// DeleteProjectCustomDomains
// 按域名解绑项目的自定义域名。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteProjectCustomDomainsWithContext(ctx context.Context, request *DeleteProjectCustomDomainsRequest) (response *DeleteProjectCustomDomainsResponse, err error) {
    if request == nil {
        request = NewDeleteProjectCustomDomainsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "makers", APIVersion, "DeleteProjectCustomDomains")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteProjectCustomDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteProjectCustomDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteProjectEnvsRequest() (request *DeleteProjectEnvsRequest) {
    request = &DeleteProjectEnvsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("makers", APIVersion, "DeleteProjectEnvs")
    
    
    return
}

func NewDeleteProjectEnvsResponse() (response *DeleteProjectEnvsResponse) {
    response = &DeleteProjectEnvsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteProjectEnvs
// 按 EnvVars 删除环境变量。按 Id 或 Key 匹配。Value 必填但不参与匹配，可传任意值。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteProjectEnvs(request *DeleteProjectEnvsRequest) (response *DeleteProjectEnvsResponse, err error) {
    return c.DeleteProjectEnvsWithContext(context.Background(), request)
}

// DeleteProjectEnvs
// 按 EnvVars 删除环境变量。按 Id 或 Key 匹配。Value 必填但不参与匹配，可传任意值。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteProjectEnvsWithContext(ctx context.Context, request *DeleteProjectEnvsRequest) (response *DeleteProjectEnvsResponse, err error) {
    if request == nil {
        request = NewDeleteProjectEnvsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "makers", APIVersion, "DeleteProjectEnvs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteProjectEnvs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteProjectEnvsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCosTempTokenRequest() (request *DescribeCosTempTokenRequest) {
    request = &DescribeCosTempTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("makers", APIVersion, "DescribeCosTempToken")
    
    
    return
}

func NewDescribeCosTempTokenResponse() (response *DescribeCosTempTokenResponse) {
    response = &DescribeCosTempTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCosTempToken
// 获取上传制品用的 COS 桶、地域、目标路径和临时密钥。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeCosTempToken(request *DescribeCosTempTokenRequest) (response *DescribeCosTempTokenResponse, err error) {
    return c.DescribeCosTempTokenWithContext(context.Background(), request)
}

// DescribeCosTempToken
// 获取上传制品用的 COS 桶、地域、目标路径和临时密钥。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeCosTempTokenWithContext(ctx context.Context, request *DescribeCosTempTokenRequest) (response *DescribeCosTempTokenResponse, err error) {
    if request == nil {
        request = NewDescribeCosTempTokenRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "makers", APIVersion, "DescribeCosTempToken")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCosTempToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCosTempTokenResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeploymentLogRequest() (request *DescribeDeploymentLogRequest) {
    request = &DescribeDeploymentLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("makers", APIVersion, "DescribeDeploymentLog")
    
    
    return
}

func NewDescribeDeploymentLogResponse() (response *DescribeDeploymentLogResponse) {
    response = &DescribeDeploymentLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeploymentLog
// 返回构建日志的访问地址。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDeploymentLog(request *DescribeDeploymentLogRequest) (response *DescribeDeploymentLogResponse, err error) {
    return c.DescribeDeploymentLogWithContext(context.Background(), request)
}

// DescribeDeploymentLog
// 返回构建日志的访问地址。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDeploymentLogWithContext(ctx context.Context, request *DescribeDeploymentLogRequest) (response *DescribeDeploymentLogResponse, err error) {
    if request == nil {
        request = NewDescribeDeploymentLogRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "makers", APIVersion, "DescribeDeploymentLog")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeploymentLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeploymentLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeploymentsRequest() (request *DescribeDeploymentsRequest) {
    request = &DescribeDeploymentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("makers", APIVersion, "DescribeDeployments")
    
    
    return
}

func NewDescribeDeploymentsResponse() (response *DescribeDeploymentsResponse) {
    response = &DescribeDeploymentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeployments
// 查询项目的部署记录。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDeployments(request *DescribeDeploymentsRequest) (response *DescribeDeploymentsResponse, err error) {
    return c.DescribeDeploymentsWithContext(context.Background(), request)
}

// DescribeDeployments
// 查询项目的部署记录。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDeploymentsWithContext(ctx context.Context, request *DescribeDeploymentsRequest) (response *DescribeDeploymentsResponse, err error) {
    if request == nil {
        request = NewDescribeDeploymentsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "makers", APIVersion, "DescribeDeployments")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeployments require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeploymentsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEnvironmentsRequest() (request *DescribeEnvironmentsRequest) {
    request = &DescribeEnvironmentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("makers", APIVersion, "DescribeEnvironments")
    
    
    return
}

func NewDescribeEnvironmentsResponse() (response *DescribeEnvironmentsResponse) {
    response = &DescribeEnvironmentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEnvironments
// 查询项目下的环境，含内置 Production、Preview 和自定义环境。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeEnvironments(request *DescribeEnvironmentsRequest) (response *DescribeEnvironmentsResponse, err error) {
    return c.DescribeEnvironmentsWithContext(context.Background(), request)
}

// DescribeEnvironments
// 查询项目下的环境，含内置 Production、Preview 和自定义环境。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeEnvironmentsWithContext(ctx context.Context, request *DescribeEnvironmentsRequest) (response *DescribeEnvironmentsResponse, err error) {
    if request == nil {
        request = NewDescribeEnvironmentsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "makers", APIVersion, "DescribeEnvironments")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEnvironments require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEnvironmentsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProjectEnvsRequest() (request *DescribeProjectEnvsRequest) {
    request = &DescribeProjectEnvsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("makers", APIVersion, "DescribeProjectEnvs")
    
    
    return
}

func NewDescribeProjectEnvsResponse() (response *DescribeProjectEnvsResponse) {
    response = &DescribeProjectEnvsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProjectEnvs
// 查询项目环境变量。不传 Env 时返回全部作用域。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeProjectEnvs(request *DescribeProjectEnvsRequest) (response *DescribeProjectEnvsResponse, err error) {
    return c.DescribeProjectEnvsWithContext(context.Background(), request)
}

// DescribeProjectEnvs
// 查询项目环境变量。不传 Env 时返回全部作用域。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeProjectEnvsWithContext(ctx context.Context, request *DescribeProjectEnvsRequest) (response *DescribeProjectEnvsResponse, err error) {
    if request == nil {
        request = NewDescribeProjectEnvsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "makers", APIVersion, "DescribeProjectEnvs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProjectEnvs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProjectEnvsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProjectsRequest() (request *DescribeProjectsRequest) {
    request = &DescribeProjectsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("makers", APIVersion, "DescribeProjects")
    
    
    return
}

func NewDescribeProjectsResponse() (response *DescribeProjectsResponse) {
    response = &DescribeProjectsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProjects
// 分页查询站点下的项目。Offset 是页码偏移，后端按 Offset * Limit 跳过。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeProjects(request *DescribeProjectsRequest) (response *DescribeProjectsResponse, err error) {
    return c.DescribeProjectsWithContext(context.Background(), request)
}

// DescribeProjects
// 分页查询站点下的项目。Offset 是页码偏移，后端按 Offset * Limit 跳过。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeProjectsWithContext(ctx context.Context, request *DescribeProjectsRequest) (response *DescribeProjectsResponse, err error) {
    if request == nil {
        request = NewDescribeProjectsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "makers", APIVersion, "DescribeProjects")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProjects require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProjectsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZoneCustomDomainsRequest() (request *DescribeZoneCustomDomainsRequest) {
    request = &DescribeZoneCustomDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("makers", APIVersion, "DescribeZoneCustomDomains")
    
    
    return
}

func NewDescribeZoneCustomDomainsResponse() (response *DescribeZoneCustomDomainsResponse) {
    response = &DescribeZoneCustomDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeZoneCustomDomains
// 查询项目在站点下的自定义域名和预设域名。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeZoneCustomDomains(request *DescribeZoneCustomDomainsRequest) (response *DescribeZoneCustomDomainsResponse, err error) {
    return c.DescribeZoneCustomDomainsWithContext(context.Background(), request)
}

// DescribeZoneCustomDomains
// 查询项目在站点下的自定义域名和预设域名。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeZoneCustomDomainsWithContext(ctx context.Context, request *DescribeZoneCustomDomainsRequest) (response *DescribeZoneCustomDomainsResponse, err error) {
    if request == nil {
        request = NewDescribeZoneCustomDomainsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "makers", APIVersion, "DescribeZoneCustomDomains")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeZoneCustomDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeZoneCustomDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyProjectRequest() (request *ModifyProjectRequest) {
    request = &ModifyProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("makers", APIVersion, "ModifyProject")
    
    
    return
}

func NewModifyProjectResponse() (response *ModifyProjectResponse) {
    response = &ModifyProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyProject
// 修改项目名称和构建配置。未传的字段保持原值。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyProject(request *ModifyProjectRequest) (response *ModifyProjectResponse, err error) {
    return c.ModifyProjectWithContext(context.Background(), request)
}

// ModifyProject
// 修改项目名称和构建配置。未传的字段保持原值。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyProjectWithContext(ctx context.Context, request *ModifyProjectRequest) (response *ModifyProjectResponse, err error) {
    if request == nil {
        request = NewModifyProjectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "makers", APIVersion, "ModifyProject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyProjectResponse()
    err = c.Send(request, response)
    return
}

func NewModifyProjectEnvsRequest() (request *ModifyProjectEnvsRequest) {
    request = &ModifyProjectEnvsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("makers", APIVersion, "ModifyProjectEnvs")
    
    
    return
}

func NewModifyProjectEnvsResponse() (response *ModifyProjectEnvsResponse) {
    response = &ModifyProjectEnvsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyProjectEnvs
// 按传入的 EnvVars 写入环境变量。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyProjectEnvs(request *ModifyProjectEnvsRequest) (response *ModifyProjectEnvsResponse, err error) {
    return c.ModifyProjectEnvsWithContext(context.Background(), request)
}

// ModifyProjectEnvs
// 按传入的 EnvVars 写入环境变量。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyProjectEnvsWithContext(ctx context.Context, request *ModifyProjectEnvsRequest) (response *ModifyProjectEnvsResponse, err error) {
    if request == nil {
        request = NewModifyProjectEnvsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "makers", APIVersion, "ModifyProjectEnvs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyProjectEnvs require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyProjectEnvsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyZoneCustomDomainEnvRequest() (request *ModifyZoneCustomDomainEnvRequest) {
    request = &ModifyZoneCustomDomainEnvRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("makers", APIVersion, "ModifyZoneCustomDomainEnv")
    
    
    return
}

func NewModifyZoneCustomDomainEnvResponse() (response *ModifyZoneCustomDomainEnvResponse) {
    response = &ModifyZoneCustomDomainEnvResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyZoneCustomDomainEnv
// 把已绑定的自定义域名切到另一个环境。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyZoneCustomDomainEnv(request *ModifyZoneCustomDomainEnvRequest) (response *ModifyZoneCustomDomainEnvResponse, err error) {
    return c.ModifyZoneCustomDomainEnvWithContext(context.Background(), request)
}

// ModifyZoneCustomDomainEnv
// 把已绑定的自定义域名切到另一个环境。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyZoneCustomDomainEnvWithContext(ctx context.Context, request *ModifyZoneCustomDomainEnvRequest) (response *ModifyZoneCustomDomainEnvResponse, err error) {
    if request == nil {
        request = NewModifyZoneCustomDomainEnvRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "makers", APIVersion, "ModifyZoneCustomDomainEnv")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyZoneCustomDomainEnv require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyZoneCustomDomainEnvResponse()
    err = c.Send(request, response)
    return
}
