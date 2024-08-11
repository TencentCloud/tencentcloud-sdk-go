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

package v20210712

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-07-12"

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


func NewCreateLibraryRequest() (request *CreateLibraryRequest) {
    request = &CreateLibraryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("smh", APIVersion, "CreateLibrary")
    
    
    return
}

func NewCreateLibraryResponse() (response *CreateLibraryResponse) {
    response = &CreateLibraryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLibrary
// 创建 PaaS 服务媒体库
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_POSTPAIDFAIL = "InternalError.PostpaidFail"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BUCKETNAMEINVALID = "InvalidParameterValue.BucketNameInvalid"
//  INVALIDPARAMETERVALUE_BUCKETNAMENOTBELONGYOU = "InvalidParameterValue.BucketNameNotBelongYou"
//  INVALIDPARAMETERVALUE_BUCKETNOTFOUND = "InvalidParameterValue.BucketNotFound"
//  INVALIDPARAMETERVALUE_BUCKETNOTSUPPORT = "InvalidParameterValue.BucketNotSupport"
//  INVALIDPARAMETERVALUE_BUCKETREGIONINVALID = "InvalidParameterValue.BucketRegionInvalid"
//  INVALIDPARAMETERVALUE_COSSTORAGECLASS = "InvalidParameterValue.CosStorageClass"
//  INVALIDPARAMETERVALUE_COSSTORAGECLASSINTELLIGENTTIERING = "InvalidParameterValue.CosStorageClassIntelligentTiering"
//  UNAUTHORIZEDOPERATION_PASSROLE = "UnauthorizedOperation.PassRole"
//  UNAUTHORIZEDOPERATION_SERVICELINKEDROLE = "UnauthorizedOperation.ServiceLinkedRole"
//  UNSUPPORTEDOPERATION_BALANCELESS = "UnsupportedOperation.BalanceLess"
func (c *Client) CreateLibrary(request *CreateLibraryRequest) (response *CreateLibraryResponse, err error) {
    return c.CreateLibraryWithContext(context.Background(), request)
}

// CreateLibrary
// 创建 PaaS 服务媒体库
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_POSTPAIDFAIL = "InternalError.PostpaidFail"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BUCKETNAMEINVALID = "InvalidParameterValue.BucketNameInvalid"
//  INVALIDPARAMETERVALUE_BUCKETNAMENOTBELONGYOU = "InvalidParameterValue.BucketNameNotBelongYou"
//  INVALIDPARAMETERVALUE_BUCKETNOTFOUND = "InvalidParameterValue.BucketNotFound"
//  INVALIDPARAMETERVALUE_BUCKETNOTSUPPORT = "InvalidParameterValue.BucketNotSupport"
//  INVALIDPARAMETERVALUE_BUCKETREGIONINVALID = "InvalidParameterValue.BucketRegionInvalid"
//  INVALIDPARAMETERVALUE_COSSTORAGECLASS = "InvalidParameterValue.CosStorageClass"
//  INVALIDPARAMETERVALUE_COSSTORAGECLASSINTELLIGENTTIERING = "InvalidParameterValue.CosStorageClassIntelligentTiering"
//  UNAUTHORIZEDOPERATION_PASSROLE = "UnauthorizedOperation.PassRole"
//  UNAUTHORIZEDOPERATION_SERVICELINKEDROLE = "UnauthorizedOperation.ServiceLinkedRole"
//  UNSUPPORTEDOPERATION_BALANCELESS = "UnsupportedOperation.BalanceLess"
func (c *Client) CreateLibraryWithContext(ctx context.Context, request *CreateLibraryRequest) (response *CreateLibraryResponse, err error) {
    if request == nil {
        request = NewCreateLibraryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLibrary require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLibraryResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUserRequest() (request *CreateUserRequest) {
    request = &CreateUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("smh", APIVersion, "CreateUser")
    
    
    return
}

func NewCreateUserResponse() (response *CreateUserResponse) {
    response = &CreateUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateUser
// 新建用户。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ONEOFACCOUNTNAMEANDACCOUNTPASSWORDISEMPTY = "InvalidParameter.OneOfAccountNameAndAccountPasswordIsEmpty"
//  INVALIDPARAMETER_ONEOFCOUNTRYCODEANDPHONENUMBERISEMPTY = "InvalidParameter.OneOfCountryCodeAndPhoneNumberIsEmpty"
//  INVALIDPARAMETERVALUE_ACCOUNTNAMETOOLONG = "InvalidParameterValue.AccountNameTooLong"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORD = "InvalidParameterValue.AccountPassword"
//  INVALIDPARAMETERVALUE_ACCOUNTUSERIDTOOLONG = "InvalidParameterValue.AccountUserIdTooLong"
//  INVALIDPARAMETERVALUE_COUNTRYCODETOOLONG = "InvalidParameterValue.CountryCodeTooLong"
//  INVALIDPARAMETERVALUE_DUPLICATEACCOUNTNAME = "InvalidParameterValue.DuplicateAccountName"
//  INVALIDPARAMETERVALUE_DUPLICATEACCOUNTUSERID = "InvalidParameterValue.DuplicateAccountUserId"
//  INVALIDPARAMETERVALUE_DUPLICATEEMAIL = "InvalidParameterValue.DuplicateEmail"
//  INVALIDPARAMETERVALUE_DUPLICATEUSERPHONENUMBER = "InvalidParameterValue.DuplicateUserPhoneNumber"
//  INVALIDPARAMETERVALUE_EMAILTOOLONG = "InvalidParameterValue.EmailTooLong"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAME = "InvalidParameterValue.InvalidAccountName"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTUSERID = "InvalidParameterValue.InvalidAccountUserId"
//  INVALIDPARAMETERVALUE_INVALIDCOUNTRYCODE = "InvalidParameterValue.InvalidCountryCode"
//  INVALIDPARAMETERVALUE_INVALIDEMAIL = "InvalidParameterValue.InvalidEmail"
//  INVALIDPARAMETERVALUE_INVALIDPHONENUMBER = "InvalidParameterValue.InvalidPhoneNumber"
//  INVALIDPARAMETERVALUE_PHONENUMBERTOOLONG = "InvalidParameterValue.PhoneNumberTooLong"
//  INVALIDPARAMETERVALUE_ROLENOTSUPPORT = "InvalidParameterValue.RoleNotSupport"
//  RESOURCENOTFOUND_LIBRARY = "ResourceNotFound.Library"
func (c *Client) CreateUser(request *CreateUserRequest) (response *CreateUserResponse, err error) {
    return c.CreateUserWithContext(context.Background(), request)
}

// CreateUser
// 新建用户。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ONEOFACCOUNTNAMEANDACCOUNTPASSWORDISEMPTY = "InvalidParameter.OneOfAccountNameAndAccountPasswordIsEmpty"
//  INVALIDPARAMETER_ONEOFCOUNTRYCODEANDPHONENUMBERISEMPTY = "InvalidParameter.OneOfCountryCodeAndPhoneNumberIsEmpty"
//  INVALIDPARAMETERVALUE_ACCOUNTNAMETOOLONG = "InvalidParameterValue.AccountNameTooLong"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORD = "InvalidParameterValue.AccountPassword"
//  INVALIDPARAMETERVALUE_ACCOUNTUSERIDTOOLONG = "InvalidParameterValue.AccountUserIdTooLong"
//  INVALIDPARAMETERVALUE_COUNTRYCODETOOLONG = "InvalidParameterValue.CountryCodeTooLong"
//  INVALIDPARAMETERVALUE_DUPLICATEACCOUNTNAME = "InvalidParameterValue.DuplicateAccountName"
//  INVALIDPARAMETERVALUE_DUPLICATEACCOUNTUSERID = "InvalidParameterValue.DuplicateAccountUserId"
//  INVALIDPARAMETERVALUE_DUPLICATEEMAIL = "InvalidParameterValue.DuplicateEmail"
//  INVALIDPARAMETERVALUE_DUPLICATEUSERPHONENUMBER = "InvalidParameterValue.DuplicateUserPhoneNumber"
//  INVALIDPARAMETERVALUE_EMAILTOOLONG = "InvalidParameterValue.EmailTooLong"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAME = "InvalidParameterValue.InvalidAccountName"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTUSERID = "InvalidParameterValue.InvalidAccountUserId"
//  INVALIDPARAMETERVALUE_INVALIDCOUNTRYCODE = "InvalidParameterValue.InvalidCountryCode"
//  INVALIDPARAMETERVALUE_INVALIDEMAIL = "InvalidParameterValue.InvalidEmail"
//  INVALIDPARAMETERVALUE_INVALIDPHONENUMBER = "InvalidParameterValue.InvalidPhoneNumber"
//  INVALIDPARAMETERVALUE_PHONENUMBERTOOLONG = "InvalidParameterValue.PhoneNumberTooLong"
//  INVALIDPARAMETERVALUE_ROLENOTSUPPORT = "InvalidParameterValue.RoleNotSupport"
//  RESOURCENOTFOUND_LIBRARY = "ResourceNotFound.Library"
func (c *Client) CreateUserWithContext(ctx context.Context, request *CreateUserRequest) (response *CreateUserResponse, err error) {
    if request == nil {
        request = NewCreateUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUserResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUserLifecycleRequest() (request *CreateUserLifecycleRequest) {
    request = &CreateUserLifecycleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("smh", APIVersion, "CreateUserLifecycle")
    
    
    return
}

func NewCreateUserLifecycleResponse() (response *CreateUserLifecycleResponse) {
    response = &CreateUserLifecycleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateUserLifecycle
// 设置用户生命周期。如果指定的用户已经设置了生命周期，重复调用此接口将覆盖已有的设置。也可用于清除指定用户的生命周期。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DESTROYTIME = "InvalidParameterValue.DestroyTime"
//  INVALIDPARAMETERVALUE_FILTER = "InvalidParameterValue.Filter"
//  INVALIDPARAMETERVALUE_ISOLATETIME = "InvalidParameterValue.IsolateTime"
//  RESOURCENOTFOUND_LIBRARY = "ResourceNotFound.Library"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) CreateUserLifecycle(request *CreateUserLifecycleRequest) (response *CreateUserLifecycleResponse, err error) {
    return c.CreateUserLifecycleWithContext(context.Background(), request)
}

// CreateUserLifecycle
// 设置用户生命周期。如果指定的用户已经设置了生命周期，重复调用此接口将覆盖已有的设置。也可用于清除指定用户的生命周期。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DESTROYTIME = "InvalidParameterValue.DestroyTime"
//  INVALIDPARAMETERVALUE_FILTER = "InvalidParameterValue.Filter"
//  INVALIDPARAMETERVALUE_ISOLATETIME = "InvalidParameterValue.IsolateTime"
//  RESOURCENOTFOUND_LIBRARY = "ResourceNotFound.Library"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) CreateUserLifecycleWithContext(ctx context.Context, request *CreateUserLifecycleRequest) (response *CreateUserLifecycleResponse, err error) {
    if request == nil {
        request = NewCreateUserLifecycleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUserLifecycle require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUserLifecycleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLibraryRequest() (request *DeleteLibraryRequest) {
    request = &DeleteLibraryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("smh", APIVersion, "DeleteLibrary")
    
    
    return
}

func NewDeleteLibraryResponse() (response *DeleteLibraryResponse) {
    response = &DeleteLibraryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLibrary
// 删除 PaaS 服务媒体库
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MODIFYRESOURCESTATUSFAIL = "InternalError.ModifyResourceStatusFail"
//  RESOURCEINUSE_MULTISPACE = "ResourceInUse.MultiSpace"
//  RESOURCENOTFOUND_LIBRARY = "ResourceNotFound.Library"
func (c *Client) DeleteLibrary(request *DeleteLibraryRequest) (response *DeleteLibraryResponse, err error) {
    return c.DeleteLibraryWithContext(context.Background(), request)
}

// DeleteLibrary
// 删除 PaaS 服务媒体库
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MODIFYRESOURCESTATUSFAIL = "InternalError.ModifyResourceStatusFail"
//  RESOURCEINUSE_MULTISPACE = "ResourceInUse.MultiSpace"
//  RESOURCENOTFOUND_LIBRARY = "ResourceNotFound.Library"
func (c *Client) DeleteLibraryWithContext(ctx context.Context, request *DeleteLibraryRequest) (response *DeleteLibraryResponse, err error) {
    if request == nil {
        request = NewDeleteLibraryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLibrary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLibraryResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUserRequest() (request *DeleteUserRequest) {
    request = &DeleteUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("smh", APIVersion, "DeleteUser")
    
    
    return
}

func NewDeleteUserResponse() (response *DeleteUserResponse) {
    response = &DeleteUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteUser
// 一次删除多个用户。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_FILTER = "InvalidParameterValue.Filter"
//  INVALIDPARAMETERVALUE_TOOMANYFILTERS = "InvalidParameterValue.TooManyFilters"
//  RESOURCENOTFOUND_LIBRARY = "ResourceNotFound.Library"
func (c *Client) DeleteUser(request *DeleteUserRequest) (response *DeleteUserResponse, err error) {
    return c.DeleteUserWithContext(context.Background(), request)
}

// DeleteUser
// 一次删除多个用户。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_FILTER = "InvalidParameterValue.Filter"
//  INVALIDPARAMETERVALUE_TOOMANYFILTERS = "InvalidParameterValue.TooManyFilters"
//  RESOURCENOTFOUND_LIBRARY = "ResourceNotFound.Library"
func (c *Client) DeleteUserWithContext(ctx context.Context, request *DeleteUserRequest) (response *DeleteUserResponse, err error) {
    if request == nil {
        request = NewDeleteUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUserResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLibrariesRequest() (request *DescribeLibrariesRequest) {
    request = &DescribeLibrariesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("smh", APIVersion, "DescribeLibraries")
    
    
    return
}

func NewDescribeLibrariesResponse() (response *DescribeLibrariesResponse) {
    response = &DescribeLibrariesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLibraries
// 查询 PaaS 服务媒体库列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
func (c *Client) DescribeLibraries(request *DescribeLibrariesRequest) (response *DescribeLibrariesResponse, err error) {
    return c.DescribeLibrariesWithContext(context.Background(), request)
}

// DescribeLibraries
// 查询 PaaS 服务媒体库列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
func (c *Client) DescribeLibrariesWithContext(ctx context.Context, request *DescribeLibrariesRequest) (response *DescribeLibrariesResponse, err error) {
    if request == nil {
        request = NewDescribeLibrariesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLibraries require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLibrariesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLibrarySecretRequest() (request *DescribeLibrarySecretRequest) {
    request = &DescribeLibrarySecretRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("smh", APIVersion, "DescribeLibrarySecret")
    
    
    return
}

func NewDescribeLibrarySecretResponse() (response *DescribeLibrarySecretResponse) {
    response = &DescribeLibrarySecretResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLibrarySecret
// 查询 PaaS 服务媒体库密钥
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_LIBRARY = "ResourceNotFound.Library"
func (c *Client) DescribeLibrarySecret(request *DescribeLibrarySecretRequest) (response *DescribeLibrarySecretResponse, err error) {
    return c.DescribeLibrarySecretWithContext(context.Background(), request)
}

// DescribeLibrarySecret
// 查询 PaaS 服务媒体库密钥
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_LIBRARY = "ResourceNotFound.Library"
func (c *Client) DescribeLibrarySecretWithContext(ctx context.Context, request *DescribeLibrarySecretRequest) (response *DescribeLibrarySecretResponse, err error) {
    if request == nil {
        request = NewDescribeLibrarySecretRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLibrarySecret require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLibrarySecretResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOfficialInstancesRequest() (request *DescribeOfficialInstancesRequest) {
    request = &DescribeOfficialInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("smh", APIVersion, "DescribeOfficialInstances")
    
    
    return
}

func NewDescribeOfficialInstancesResponse() (response *DescribeOfficialInstancesResponse) {
    response = &DescribeOfficialInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOfficialInstances
// 查询官方云盘实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
func (c *Client) DescribeOfficialInstances(request *DescribeOfficialInstancesRequest) (response *DescribeOfficialInstancesResponse, err error) {
    return c.DescribeOfficialInstancesWithContext(context.Background(), request)
}

// DescribeOfficialInstances
// 查询官方云盘实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
func (c *Client) DescribeOfficialInstancesWithContext(ctx context.Context, request *DescribeOfficialInstancesRequest) (response *DescribeOfficialInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeOfficialInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOfficialInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOfficialInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOfficialOverviewRequest() (request *DescribeOfficialOverviewRequest) {
    request = &DescribeOfficialOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("smh", APIVersion, "DescribeOfficialOverview")
    
    
    return
}

func NewDescribeOfficialOverviewResponse() (response *DescribeOfficialOverviewResponse) {
    response = &DescribeOfficialOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOfficialOverview
// 查询官方云盘实例概览数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeOfficialOverview(request *DescribeOfficialOverviewRequest) (response *DescribeOfficialOverviewResponse, err error) {
    return c.DescribeOfficialOverviewWithContext(context.Background(), request)
}

// DescribeOfficialOverview
// 查询官方云盘实例概览数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeOfficialOverviewWithContext(ctx context.Context, request *DescribeOfficialOverviewRequest) (response *DescribeOfficialOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeOfficialOverviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOfficialOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOfficialOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTrafficPackagesRequest() (request *DescribeTrafficPackagesRequest) {
    request = &DescribeTrafficPackagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("smh", APIVersion, "DescribeTrafficPackages")
    
    
    return
}

func NewDescribeTrafficPackagesResponse() (response *DescribeTrafficPackagesResponse) {
    response = &DescribeTrafficPackagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTrafficPackages
// 查询流量资源包
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
func (c *Client) DescribeTrafficPackages(request *DescribeTrafficPackagesRequest) (response *DescribeTrafficPackagesResponse, err error) {
    return c.DescribeTrafficPackagesWithContext(context.Background(), request)
}

// DescribeTrafficPackages
// 查询流量资源包
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
func (c *Client) DescribeTrafficPackagesWithContext(ctx context.Context, request *DescribeTrafficPackagesRequest) (response *DescribeTrafficPackagesResponse, err error) {
    if request == nil {
        request = NewDescribeTrafficPackagesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTrafficPackages require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTrafficPackagesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserLifecycleRequest() (request *DescribeUserLifecycleRequest) {
    request = &DescribeUserLifecycleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("smh", APIVersion, "DescribeUserLifecycle")
    
    
    return
}

func NewDescribeUserLifecycleResponse() (response *DescribeUserLifecycleResponse) {
    response = &DescribeUserLifecycleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserLifecycle
// 查询用户生命周期。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_FILTER = "InvalidParameterValue.Filter"
//  RESOURCENOTFOUND_LIBRARY = "ResourceNotFound.Library"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) DescribeUserLifecycle(request *DescribeUserLifecycleRequest) (response *DescribeUserLifecycleResponse, err error) {
    return c.DescribeUserLifecycleWithContext(context.Background(), request)
}

// DescribeUserLifecycle
// 查询用户生命周期。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_FILTER = "InvalidParameterValue.Filter"
//  RESOURCENOTFOUND_LIBRARY = "ResourceNotFound.Library"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) DescribeUserLifecycleWithContext(ctx context.Context, request *DescribeUserLifecycleRequest) (response *DescribeUserLifecycleResponse, err error) {
    if request == nil {
        request = NewDescribeUserLifecycleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserLifecycle require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserLifecycleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLibraryRequest() (request *ModifyLibraryRequest) {
    request = &ModifyLibraryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("smh", APIVersion, "ModifyLibrary")
    
    
    return
}

func NewModifyLibraryResponse() (response *ModifyLibraryResponse) {
    response = &ModifyLibraryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyLibrary
// 修改 PaaS 服务媒体库配置项
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  RESOURCENOTFOUND_LIBRARY = "ResourceNotFound.Library"
func (c *Client) ModifyLibrary(request *ModifyLibraryRequest) (response *ModifyLibraryResponse, err error) {
    return c.ModifyLibraryWithContext(context.Background(), request)
}

// ModifyLibrary
// 修改 PaaS 服务媒体库配置项
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  RESOURCENOTFOUND_LIBRARY = "ResourceNotFound.Library"
func (c *Client) ModifyLibraryWithContext(ctx context.Context, request *ModifyLibraryRequest) (response *ModifyLibraryResponse, err error) {
    if request == nil {
        request = NewModifyLibraryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLibrary require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLibraryResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUserRequest() (request *ModifyUserRequest) {
    request = &ModifyUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("smh", APIVersion, "ModifyUser")
    
    
    return
}

func NewModifyUserResponse() (response *ModifyUserResponse) {
    response = &ModifyUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyUser
// 更新用户信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ONEOFACCOUNTNAMEANDACCOUNTPASSWORDISEMPTY = "InvalidParameter.OneOfAccountNameAndAccountPasswordIsEmpty"
//  INVALIDPARAMETER_ONEOFCOUNTRYCODEANDPHONENUMBERISEMPTY = "InvalidParameter.OneOfCountryCodeAndPhoneNumberIsEmpty"
//  INVALIDPARAMETERVALUE_ACCOUNTNAMETOOLONG = "InvalidParameterValue.AccountNameTooLong"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORD = "InvalidParameterValue.AccountPassword"
//  INVALIDPARAMETERVALUE_ACCOUNTUSERIDTOOLONG = "InvalidParameterValue.AccountUserIdTooLong"
//  INVALIDPARAMETERVALUE_COUNTRYCODETOOLONG = "InvalidParameterValue.CountryCodeTooLong"
//  INVALIDPARAMETERVALUE_DUPLICATEACCOUNTNAME = "InvalidParameterValue.DuplicateAccountName"
//  INVALIDPARAMETERVALUE_DUPLICATEACCOUNTUSERID = "InvalidParameterValue.DuplicateAccountUserId"
//  INVALIDPARAMETERVALUE_DUPLICATEEMAIL = "InvalidParameterValue.DuplicateEmail"
//  INVALIDPARAMETERVALUE_DUPLICATEUSERPHONENUMBER = "InvalidParameterValue.DuplicateUserPhoneNumber"
//  INVALIDPARAMETERVALUE_EMAILTOOLONG = "InvalidParameterValue.EmailTooLong"
//  INVALIDPARAMETERVALUE_FILTER = "InvalidParameterValue.Filter"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAME = "InvalidParameterValue.InvalidAccountName"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTUSERID = "InvalidParameterValue.InvalidAccountUserId"
//  INVALIDPARAMETERVALUE_INVALIDCOUNTRYCODE = "InvalidParameterValue.InvalidCountryCode"
//  INVALIDPARAMETERVALUE_INVALIDEMAIL = "InvalidParameterValue.InvalidEmail"
//  INVALIDPARAMETERVALUE_INVALIDPHONENUMBER = "InvalidParameterValue.InvalidPhoneNumber"
//  INVALIDPARAMETERVALUE_PHONENUMBERTOOLONG = "InvalidParameterValue.PhoneNumberTooLong"
//  INVALIDPARAMETERVALUE_ROLENOTSUPPORT = "InvalidParameterValue.RoleNotSupport"
//  RESOURCENOTFOUND_LIBRARY = "ResourceNotFound.Library"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) ModifyUser(request *ModifyUserRequest) (response *ModifyUserResponse, err error) {
    return c.ModifyUserWithContext(context.Background(), request)
}

// ModifyUser
// 更新用户信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ONEOFACCOUNTNAMEANDACCOUNTPASSWORDISEMPTY = "InvalidParameter.OneOfAccountNameAndAccountPasswordIsEmpty"
//  INVALIDPARAMETER_ONEOFCOUNTRYCODEANDPHONENUMBERISEMPTY = "InvalidParameter.OneOfCountryCodeAndPhoneNumberIsEmpty"
//  INVALIDPARAMETERVALUE_ACCOUNTNAMETOOLONG = "InvalidParameterValue.AccountNameTooLong"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORD = "InvalidParameterValue.AccountPassword"
//  INVALIDPARAMETERVALUE_ACCOUNTUSERIDTOOLONG = "InvalidParameterValue.AccountUserIdTooLong"
//  INVALIDPARAMETERVALUE_COUNTRYCODETOOLONG = "InvalidParameterValue.CountryCodeTooLong"
//  INVALIDPARAMETERVALUE_DUPLICATEACCOUNTNAME = "InvalidParameterValue.DuplicateAccountName"
//  INVALIDPARAMETERVALUE_DUPLICATEACCOUNTUSERID = "InvalidParameterValue.DuplicateAccountUserId"
//  INVALIDPARAMETERVALUE_DUPLICATEEMAIL = "InvalidParameterValue.DuplicateEmail"
//  INVALIDPARAMETERVALUE_DUPLICATEUSERPHONENUMBER = "InvalidParameterValue.DuplicateUserPhoneNumber"
//  INVALIDPARAMETERVALUE_EMAILTOOLONG = "InvalidParameterValue.EmailTooLong"
//  INVALIDPARAMETERVALUE_FILTER = "InvalidParameterValue.Filter"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAME = "InvalidParameterValue.InvalidAccountName"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTUSERID = "InvalidParameterValue.InvalidAccountUserId"
//  INVALIDPARAMETERVALUE_INVALIDCOUNTRYCODE = "InvalidParameterValue.InvalidCountryCode"
//  INVALIDPARAMETERVALUE_INVALIDEMAIL = "InvalidParameterValue.InvalidEmail"
//  INVALIDPARAMETERVALUE_INVALIDPHONENUMBER = "InvalidParameterValue.InvalidPhoneNumber"
//  INVALIDPARAMETERVALUE_PHONENUMBERTOOLONG = "InvalidParameterValue.PhoneNumberTooLong"
//  INVALIDPARAMETERVALUE_ROLENOTSUPPORT = "InvalidParameterValue.RoleNotSupport"
//  RESOURCENOTFOUND_LIBRARY = "ResourceNotFound.Library"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) ModifyUserWithContext(ctx context.Context, request *ModifyUserRequest) (response *ModifyUserResponse, err error) {
    if request == nil {
        request = NewModifyUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUserResponse()
    err = c.Send(request, response)
    return
}

func NewSendSmsCodeRequest() (request *SendSmsCodeRequest) {
    request = &SendSmsCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("smh", APIVersion, "SendSmsCode")
    
    
    return
}

func NewSendSmsCodeResponse() (response *SendSmsCodeResponse) {
    response = &SendSmsCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SendSmsCode
// 发送用于换绑官方云盘实例的超级管理员账号的短信验证码
//
// 可能返回的错误码:
//  INTERNALERROR_SENDSMS = "InternalError.SendSms"
//  INVALIDPARAMETERVALUE_COUNTRYCODE = "InvalidParameterValue.CountryCode"
//  INVALIDPARAMETERVALUE_NOTMODIFIED = "InvalidParameterValue.NotModified"
//  INVALIDPARAMETERVALUE_NOTSUPERADMIN = "InvalidParameterValue.NotSuperAdmin"
//  INVALIDPARAMETERVALUE_PHONENUMBER = "InvalidParameterValue.PhoneNumber"
//  REQUESTLIMITEXCEEDED_SENDSMS = "RequestLimitExceeded.SendSms"
//  RESOURCENOTFOUND_OFFICIALINSTANCE = "ResourceNotFound.OfficialInstance"
//  UNSUPPORTEDOPERATION_PURPOSE = "UnsupportedOperation.Purpose"
func (c *Client) SendSmsCode(request *SendSmsCodeRequest) (response *SendSmsCodeResponse, err error) {
    return c.SendSmsCodeWithContext(context.Background(), request)
}

// SendSmsCode
// 发送用于换绑官方云盘实例的超级管理员账号的短信验证码
//
// 可能返回的错误码:
//  INTERNALERROR_SENDSMS = "InternalError.SendSms"
//  INVALIDPARAMETERVALUE_COUNTRYCODE = "InvalidParameterValue.CountryCode"
//  INVALIDPARAMETERVALUE_NOTMODIFIED = "InvalidParameterValue.NotModified"
//  INVALIDPARAMETERVALUE_NOTSUPERADMIN = "InvalidParameterValue.NotSuperAdmin"
//  INVALIDPARAMETERVALUE_PHONENUMBER = "InvalidParameterValue.PhoneNumber"
//  REQUESTLIMITEXCEEDED_SENDSMS = "RequestLimitExceeded.SendSms"
//  RESOURCENOTFOUND_OFFICIALINSTANCE = "ResourceNotFound.OfficialInstance"
//  UNSUPPORTEDOPERATION_PURPOSE = "UnsupportedOperation.Purpose"
func (c *Client) SendSmsCodeWithContext(ctx context.Context, request *SendSmsCodeRequest) (response *SendSmsCodeResponse, err error) {
    if request == nil {
        request = NewSendSmsCodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SendSmsCode require credential")
    }

    request.SetContext(ctx)
    
    response = NewSendSmsCodeResponse()
    err = c.Send(request, response)
    return
}

func NewVerifySmsCodeRequest() (request *VerifySmsCodeRequest) {
    request = &VerifySmsCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("smh", APIVersion, "VerifySmsCode")
    
    
    return
}

func NewVerifySmsCodeResponse() (response *VerifySmsCodeResponse) {
    response = &VerifySmsCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// VerifySmsCode
// 验证短信验证码以换绑官方云盘实例的超级管理员账号
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_COUNTRYCODE = "InvalidParameterValue.CountryCode"
//  INVALIDPARAMETERVALUE_PHONENUMBER = "InvalidParameterValue.PhoneNumber"
//  LIMITEXCEEDED_USERLIMIT = "LimitExceeded.UserLimit"
//  RESOURCENOTFOUND_OFFICIALINSTANCE = "ResourceNotFound.OfficialInstance"
//  UNAUTHORIZEDOPERATION_SMSCODE = "UnauthorizedOperation.SmsCode"
//  UNAUTHORIZEDOPERATION_SMSCODEEXCEEDED = "UnauthorizedOperation.SmsCodeExceeded"
//  UNSUPPORTEDOPERATION_PURPOSE = "UnsupportedOperation.Purpose"
func (c *Client) VerifySmsCode(request *VerifySmsCodeRequest) (response *VerifySmsCodeResponse, err error) {
    return c.VerifySmsCodeWithContext(context.Background(), request)
}

// VerifySmsCode
// 验证短信验证码以换绑官方云盘实例的超级管理员账号
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_COUNTRYCODE = "InvalidParameterValue.CountryCode"
//  INVALIDPARAMETERVALUE_PHONENUMBER = "InvalidParameterValue.PhoneNumber"
//  LIMITEXCEEDED_USERLIMIT = "LimitExceeded.UserLimit"
//  RESOURCENOTFOUND_OFFICIALINSTANCE = "ResourceNotFound.OfficialInstance"
//  UNAUTHORIZEDOPERATION_SMSCODE = "UnauthorizedOperation.SmsCode"
//  UNAUTHORIZEDOPERATION_SMSCODEEXCEEDED = "UnauthorizedOperation.SmsCodeExceeded"
//  UNSUPPORTEDOPERATION_PURPOSE = "UnsupportedOperation.Purpose"
func (c *Client) VerifySmsCodeWithContext(ctx context.Context, request *VerifySmsCodeRequest) (response *VerifySmsCodeResponse, err error) {
    if request == nil {
        request = NewVerifySmsCodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("VerifySmsCode require credential")
    }

    request.SetContext(ctx)
    
    response = NewVerifySmsCodeResponse()
    err = c.Send(request, response)
    return
}
