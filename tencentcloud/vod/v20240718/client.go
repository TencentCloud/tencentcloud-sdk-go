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

package v20240718

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2024-07-18"

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


func NewCreateIncrementalMigrationStrategyRequest() (request *CreateIncrementalMigrationStrategyRequest) {
    request = &CreateIncrementalMigrationStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateIncrementalMigrationStrategy")
    
    
    return
}

func NewCreateIncrementalMigrationStrategyResponse() (response *CreateIncrementalMigrationStrategyResponse) {
    response = &CreateIncrementalMigrationStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateIncrementalMigrationStrategy
// 创建增量迁移策略。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BUCKETIDNOTMATCH = "FailedOperation.BucketIdNotMatch"
//  FAILEDOPERATION_BUCKETINCREMENTALMIGRATIONSTRATEGYDEPLOYING = "FailedOperation.BucketIncrementalMigrationStrategyDeploying"
//  FAILEDOPERATION_INCREMENTALMIGRATIONSTRATEGYOVERLIMIT = "FailedOperation.IncrementalMigrationStrategyOverLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDSTORAGEREGION = "InvalidParameterValue.UnsupportedStorageRegion"
//  UNAUTHORIZEDOPERATION_SUBAPP = "UnauthorizedOperation.SubApp"
func (c *Client) CreateIncrementalMigrationStrategy(request *CreateIncrementalMigrationStrategyRequest) (response *CreateIncrementalMigrationStrategyResponse, err error) {
    return c.CreateIncrementalMigrationStrategyWithContext(context.Background(), request)
}

// CreateIncrementalMigrationStrategy
// 创建增量迁移策略。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BUCKETIDNOTMATCH = "FailedOperation.BucketIdNotMatch"
//  FAILEDOPERATION_BUCKETINCREMENTALMIGRATIONSTRATEGYDEPLOYING = "FailedOperation.BucketIncrementalMigrationStrategyDeploying"
//  FAILEDOPERATION_INCREMENTALMIGRATIONSTRATEGYOVERLIMIT = "FailedOperation.IncrementalMigrationStrategyOverLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDSTORAGEREGION = "InvalidParameterValue.UnsupportedStorageRegion"
//  UNAUTHORIZEDOPERATION_SUBAPP = "UnauthorizedOperation.SubApp"
func (c *Client) CreateIncrementalMigrationStrategyWithContext(ctx context.Context, request *CreateIncrementalMigrationStrategyRequest) (response *CreateIncrementalMigrationStrategyResponse, err error) {
    if request == nil {
        request = NewCreateIncrementalMigrationStrategyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateIncrementalMigrationStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateIncrementalMigrationStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateStorageRequest() (request *CreateStorageRequest) {
    request = &CreateStorageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateStorage")
    
    
    return
}

func NewCreateStorageResponse() (response *CreateStorageResponse) {
    response = &CreateStorageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateStorage
// 该接口用于为专业版应用创建存储桶。
//
// 
//
// 注：
//
// - 本接口仅用于专业版应用；
//
// - 客户创建点播专业版应用时，系统默认为客户开通了部分地域的存储，用户如果需要开通其它地域的存储，可以通过该接口进行开通；
//
// - 通过 [DescribeStorageRegions](https://cloud.tencent.com/document/product/266/72480) 接口可以查询到所有存储地域及已经开通存储桶的地域。
//
// 可能返回的错误码:
//  FAILEDOPERATION_STORAGEREGIONBUCKETOVERLIMIT = "FailedOperation.StorageRegionBucketOverLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDSTORAGEREGION = "InvalidParameterValue.UnsupportedStorageRegion"
func (c *Client) CreateStorage(request *CreateStorageRequest) (response *CreateStorageResponse, err error) {
    return c.CreateStorageWithContext(context.Background(), request)
}

// CreateStorage
// 该接口用于为专业版应用创建存储桶。
//
// 
//
// 注：
//
// - 本接口仅用于专业版应用；
//
// - 客户创建点播专业版应用时，系统默认为客户开通了部分地域的存储，用户如果需要开通其它地域的存储，可以通过该接口进行开通；
//
// - 通过 [DescribeStorageRegions](https://cloud.tencent.com/document/product/266/72480) 接口可以查询到所有存储地域及已经开通存储桶的地域。
//
// 可能返回的错误码:
//  FAILEDOPERATION_STORAGEREGIONBUCKETOVERLIMIT = "FailedOperation.StorageRegionBucketOverLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDSTORAGEREGION = "InvalidParameterValue.UnsupportedStorageRegion"
func (c *Client) CreateStorageWithContext(ctx context.Context, request *CreateStorageRequest) (response *CreateStorageResponse, err error) {
    if request == nil {
        request = NewCreateStorageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateStorage require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateStorageResponse()
    err = c.Send(request, response)
    return
}

func NewCreateStorageCredentialsRequest() (request *CreateStorageCredentialsRequest) {
    request = &CreateStorageCredentialsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateStorageCredentials")
    
    
    return
}

func NewCreateStorageCredentialsResponse() (response *CreateStorageCredentialsResponse) {
    response = &CreateStorageCredentialsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateStorageCredentials
// 用于按指定策略，生成专业版应用的临时访问凭证，比如生成用于客户端上传的临时凭证。
//
// 可能返回的错误码:
//  FAILEDOPERATION_STORAGEREGIONBUCKETOVERLIMIT = "FailedOperation.StorageRegionBucketOverLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDSTORAGEREGION = "InvalidParameterValue.UnsupportedStorageRegion"
func (c *Client) CreateStorageCredentials(request *CreateStorageCredentialsRequest) (response *CreateStorageCredentialsResponse, err error) {
    return c.CreateStorageCredentialsWithContext(context.Background(), request)
}

// CreateStorageCredentials
// 用于按指定策略，生成专业版应用的临时访问凭证，比如生成用于客户端上传的临时凭证。
//
// 可能返回的错误码:
//  FAILEDOPERATION_STORAGEREGIONBUCKETOVERLIMIT = "FailedOperation.StorageRegionBucketOverLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDSTORAGEREGION = "InvalidParameterValue.UnsupportedStorageRegion"
func (c *Client) CreateStorageCredentialsWithContext(ctx context.Context, request *CreateStorageCredentialsRequest) (response *CreateStorageCredentialsResponse, err error) {
    if request == nil {
        request = NewCreateStorageCredentialsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateStorageCredentials require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateStorageCredentialsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteIncrementalMigrationStrategyRequest() (request *DeleteIncrementalMigrationStrategyRequest) {
    request = &DeleteIncrementalMigrationStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DeleteIncrementalMigrationStrategy")
    
    
    return
}

func NewDeleteIncrementalMigrationStrategyResponse() (response *DeleteIncrementalMigrationStrategyResponse) {
    response = &DeleteIncrementalMigrationStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteIncrementalMigrationStrategy
// 删除增量迁移策略。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BUCKETIDNOTMATCH = "FailedOperation.BucketIdNotMatch"
//  FAILEDOPERATION_BUCKETINCREMENTALMIGRATIONSTRATEGYDEPLOYING = "FailedOperation.BucketIncrementalMigrationStrategyDeploying"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_INCREMENTALMIGRATIONSTRATEGYNOTFOUND = "ResourceNotFound.IncrementalMigrationStrategyNotFound"
//  UNAUTHORIZEDOPERATION_SUBAPP = "UnauthorizedOperation.SubApp"
func (c *Client) DeleteIncrementalMigrationStrategy(request *DeleteIncrementalMigrationStrategyRequest) (response *DeleteIncrementalMigrationStrategyResponse, err error) {
    return c.DeleteIncrementalMigrationStrategyWithContext(context.Background(), request)
}

// DeleteIncrementalMigrationStrategy
// 删除增量迁移策略。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BUCKETIDNOTMATCH = "FailedOperation.BucketIdNotMatch"
//  FAILEDOPERATION_BUCKETINCREMENTALMIGRATIONSTRATEGYDEPLOYING = "FailedOperation.BucketIncrementalMigrationStrategyDeploying"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_INCREMENTALMIGRATIONSTRATEGYNOTFOUND = "ResourceNotFound.IncrementalMigrationStrategyNotFound"
//  UNAUTHORIZEDOPERATION_SUBAPP = "UnauthorizedOperation.SubApp"
func (c *Client) DeleteIncrementalMigrationStrategyWithContext(ctx context.Context, request *DeleteIncrementalMigrationStrategyRequest) (response *DeleteIncrementalMigrationStrategyResponse, err error) {
    if request == nil {
        request = NewDeleteIncrementalMigrationStrategyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteIncrementalMigrationStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteIncrementalMigrationStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIncrementalMigrationStrategyInfosRequest() (request *DescribeIncrementalMigrationStrategyInfosRequest) {
    request = &DescribeIncrementalMigrationStrategyInfosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeIncrementalMigrationStrategyInfos")
    
    
    return
}

func NewDescribeIncrementalMigrationStrategyInfosResponse() (response *DescribeIncrementalMigrationStrategyInfosResponse) {
    response = &DescribeIncrementalMigrationStrategyInfosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIncrementalMigrationStrategyInfos
// 查询增量迁移策略信息。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeIncrementalMigrationStrategyInfos(request *DescribeIncrementalMigrationStrategyInfosRequest) (response *DescribeIncrementalMigrationStrategyInfosResponse, err error) {
    return c.DescribeIncrementalMigrationStrategyInfosWithContext(context.Background(), request)
}

// DescribeIncrementalMigrationStrategyInfos
// 查询增量迁移策略信息。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeIncrementalMigrationStrategyInfosWithContext(ctx context.Context, request *DescribeIncrementalMigrationStrategyInfosRequest) (response *DescribeIncrementalMigrationStrategyInfosResponse, err error) {
    if request == nil {
        request = NewDescribeIncrementalMigrationStrategyInfosRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIncrementalMigrationStrategyInfos require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIncrementalMigrationStrategyInfosResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStorageRequest() (request *DescribeStorageRequest) {
    request = &DescribeStorageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "DescribeStorage")
    
    
    return
}

func NewDescribeStorageResponse() (response *DescribeStorageResponse) {
    response = &DescribeStorageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStorage
// 该接口用于查询专业版应用中的存储桶信息，同时支持分页查询。
//
// 
//
// 注：
//
// - 本接口仅用于专业版应用。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeStorage(request *DescribeStorageRequest) (response *DescribeStorageResponse, err error) {
    return c.DescribeStorageWithContext(context.Background(), request)
}

// DescribeStorage
// 该接口用于查询专业版应用中的存储桶信息，同时支持分页查询。
//
// 
//
// 注：
//
// - 本接口仅用于专业版应用。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeStorageWithContext(ctx context.Context, request *DescribeStorageRequest) (response *DescribeStorageResponse, err error) {
    if request == nil {
        request = NewDescribeStorageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStorage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStorageResponse()
    err = c.Send(request, response)
    return
}

func NewModifyIncrementalMigrationStrategyRequest() (request *ModifyIncrementalMigrationStrategyRequest) {
    request = &ModifyIncrementalMigrationStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "ModifyIncrementalMigrationStrategy")
    
    
    return
}

func NewModifyIncrementalMigrationStrategyResponse() (response *ModifyIncrementalMigrationStrategyResponse) {
    response = &ModifyIncrementalMigrationStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyIncrementalMigrationStrategy
// 创建增量迁移策略。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BUCKETIDNOTMATCH = "FailedOperation.BucketIdNotMatch"
//  FAILEDOPERATION_BUCKETINCREMENTALMIGRATIONSTRATEGYDEPLOYING = "FailedOperation.BucketIncrementalMigrationStrategyDeploying"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_INCREMENTALMIGRATIONSTRATEGYNOTFOUND = "ResourceNotFound.IncrementalMigrationStrategyNotFound"
//  UNAUTHORIZEDOPERATION_SUBAPP = "UnauthorizedOperation.SubApp"
func (c *Client) ModifyIncrementalMigrationStrategy(request *ModifyIncrementalMigrationStrategyRequest) (response *ModifyIncrementalMigrationStrategyResponse, err error) {
    return c.ModifyIncrementalMigrationStrategyWithContext(context.Background(), request)
}

// ModifyIncrementalMigrationStrategy
// 创建增量迁移策略。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BUCKETIDNOTMATCH = "FailedOperation.BucketIdNotMatch"
//  FAILEDOPERATION_BUCKETINCREMENTALMIGRATIONSTRATEGYDEPLOYING = "FailedOperation.BucketIncrementalMigrationStrategyDeploying"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_INCREMENTALMIGRATIONSTRATEGYNOTFOUND = "ResourceNotFound.IncrementalMigrationStrategyNotFound"
//  UNAUTHORIZEDOPERATION_SUBAPP = "UnauthorizedOperation.SubApp"
func (c *Client) ModifyIncrementalMigrationStrategyWithContext(ctx context.Context, request *ModifyIncrementalMigrationStrategyRequest) (response *ModifyIncrementalMigrationStrategyResponse, err error) {
    if request == nil {
        request = NewModifyIncrementalMigrationStrategyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyIncrementalMigrationStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyIncrementalMigrationStrategyResponse()
    err = c.Send(request, response)
    return
}
