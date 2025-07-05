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

package v20230812

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2023-08-12"

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


func NewCreateApplicationRequest() (request *CreateApplicationRequest) {
    request = &CreateApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hai", APIVersion, "CreateApplication")
    
    
    return
}

func NewCreateApplicationResponse() (response *CreateApplicationResponse) {
    response = &CreateApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateApplication
// 本接口（CreateApplication）用于对HAI实例制作自定义应用。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INSTANCEIDNOTFOUND = "InvalidParameterValue.InstanceIdNotFound"
//  INVALIDPARAMETERVALUE_INVALIDAPPLICATIONNAMEDUPLICATE = "InvalidParameterValue.InvalidApplicationNameDuplicate"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEIDMALFORMED = "InvalidParameterValue.InvalidInstanceIdMalformed"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  UNSUPPORTEDOPERATION_INSTANCECREATEAPPLICATIONNOTSUPPORT = "UnsupportedOperation.InstanceCreateApplicationNotSupport"
func (c *Client) CreateApplication(request *CreateApplicationRequest) (response *CreateApplicationResponse, err error) {
    return c.CreateApplicationWithContext(context.Background(), request)
}

// CreateApplication
// 本接口（CreateApplication）用于对HAI实例制作自定义应用。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INSTANCEIDNOTFOUND = "InvalidParameterValue.InstanceIdNotFound"
//  INVALIDPARAMETERVALUE_INVALIDAPPLICATIONNAMEDUPLICATE = "InvalidParameterValue.InvalidApplicationNameDuplicate"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEIDMALFORMED = "InvalidParameterValue.InvalidInstanceIdMalformed"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  UNSUPPORTEDOPERATION_INSTANCECREATEAPPLICATIONNOTSUPPORT = "UnsupportedOperation.InstanceCreateApplicationNotSupport"
func (c *Client) CreateApplicationWithContext(ctx context.Context, request *CreateApplicationRequest) (response *CreateApplicationResponse, err error) {
    if request == nil {
        request = NewCreateApplicationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMuskPromptRequest() (request *CreateMuskPromptRequest) {
    request = &CreateMuskPromptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hai", APIVersion, "CreateMuskPrompt")
    
    
    return
}

func NewCreateMuskPromptResponse() (response *CreateMuskPromptResponse) {
    response = &CreateMuskPromptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateMuskPrompt
// 创建musk prompt 任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateMuskPrompt(request *CreateMuskPromptRequest) (response *CreateMuskPromptResponse, err error) {
    return c.CreateMuskPromptWithContext(context.Background(), request)
}

// CreateMuskPrompt
// 创建musk prompt 任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateMuskPromptWithContext(ctx context.Context, request *CreateMuskPromptRequest) (response *CreateMuskPromptResponse, err error) {
    if request == nil {
        request = NewCreateMuskPromptRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMuskPrompt require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMuskPromptResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationsRequest() (request *DescribeApplicationsRequest) {
    request = &DescribeApplicationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hai", APIVersion, "DescribeApplications")
    
    
    return
}

func NewDescribeApplicationsResponse() (response *DescribeApplicationsResponse) {
    response = &DescribeApplicationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApplications
// 本接口（DescribeApplications）用于查询应用
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ATMOSTONE = "InvalidParameter.AtMostOne"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INVALIDAPPLICATIONIDMALFORMED = "InvalidParameterValue.InvalidApplicationIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDAPPLICATIONTYPE = "InvalidParameterValue.InvalidApplicationType"
//  INVALIDPARAMETERVALUE_INVALIDORDER = "InvalidParameterValue.InvalidOrder"
//  INVALIDPARAMETERVALUE_INVALIDORDERFIELD = "InvalidParameterValue.InvalidOrderField"
//  INVALIDPARAMETERVALUE_INVALIDSCENEIDMALFORMED = "InvalidParameterValue.InvalidSceneIdMalformed"
func (c *Client) DescribeApplications(request *DescribeApplicationsRequest) (response *DescribeApplicationsResponse, err error) {
    return c.DescribeApplicationsWithContext(context.Background(), request)
}

// DescribeApplications
// 本接口（DescribeApplications）用于查询应用
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ATMOSTONE = "InvalidParameter.AtMostOne"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INVALIDAPPLICATIONIDMALFORMED = "InvalidParameterValue.InvalidApplicationIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDAPPLICATIONTYPE = "InvalidParameterValue.InvalidApplicationType"
//  INVALIDPARAMETERVALUE_INVALIDORDER = "InvalidParameterValue.InvalidOrder"
//  INVALIDPARAMETERVALUE_INVALIDORDERFIELD = "InvalidParameterValue.InvalidOrderField"
//  INVALIDPARAMETERVALUE_INVALIDSCENEIDMALFORMED = "InvalidParameterValue.InvalidSceneIdMalformed"
func (c *Client) DescribeApplicationsWithContext(ctx context.Context, request *DescribeApplicationsRequest) (response *DescribeApplicationsResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplications require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceNetworkStatusRequest() (request *DescribeInstanceNetworkStatusRequest) {
    request = &DescribeInstanceNetworkStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hai", APIVersion, "DescribeInstanceNetworkStatus")
    
    
    return
}

func NewDescribeInstanceNetworkStatusResponse() (response *DescribeInstanceNetworkStatusResponse) {
    response = &DescribeInstanceNetworkStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceNetworkStatus
// 本接口（DescribeInstanceNetworkStatus）用于查询实例的网络配置及消耗情况
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INSTANCEIDNOTFOUND = "InvalidParameterValue.InstanceIdNotFound"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEIDMALFORMED = "InvalidParameterValue.InvalidInstanceIdMalformed"
func (c *Client) DescribeInstanceNetworkStatus(request *DescribeInstanceNetworkStatusRequest) (response *DescribeInstanceNetworkStatusResponse, err error) {
    return c.DescribeInstanceNetworkStatusWithContext(context.Background(), request)
}

// DescribeInstanceNetworkStatus
// 本接口（DescribeInstanceNetworkStatus）用于查询实例的网络配置及消耗情况
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INSTANCEIDNOTFOUND = "InvalidParameterValue.InstanceIdNotFound"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEIDMALFORMED = "InvalidParameterValue.InvalidInstanceIdMalformed"
func (c *Client) DescribeInstanceNetworkStatusWithContext(ctx context.Context, request *DescribeInstanceNetworkStatusRequest) (response *DescribeInstanceNetworkStatusResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceNetworkStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceNetworkStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceNetworkStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hai", APIVersion, "DescribeInstances")
    
    
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstances
// 本接口（DescribeInstances）用户查询实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ATMOSTONE = "InvalidParameter.AtMostOne"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEIDMALFORMED = "InvalidParameterValue.InvalidInstanceIdMalformed"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDUSER = "UnauthorizedOperation.UnauthorizedUser"
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    return c.DescribeInstancesWithContext(context.Background(), request)
}

// DescribeInstances
// 本接口（DescribeInstances）用户查询实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ATMOSTONE = "InvalidParameter.AtMostOne"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEIDMALFORMED = "InvalidParameterValue.InvalidInstanceIdMalformed"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDUSER = "UnauthorizedOperation.UnauthorizedUser"
func (c *Client) DescribeInstancesWithContext(ctx context.Context, request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMuskPromptsRequest() (request *DescribeMuskPromptsRequest) {
    request = &DescribeMuskPromptsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hai", APIVersion, "DescribeMuskPrompts")
    
    
    return
}

func NewDescribeMuskPromptsResponse() (response *DescribeMuskPromptsResponse) {
    response = &DescribeMuskPromptsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMuskPrompts
// 获取prompt任务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeMuskPrompts(request *DescribeMuskPromptsRequest) (response *DescribeMuskPromptsResponse, err error) {
    return c.DescribeMuskPromptsWithContext(context.Background(), request)
}

// DescribeMuskPrompts
// 获取prompt任务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeMuskPromptsWithContext(ctx context.Context, request *DescribeMuskPromptsRequest) (response *DescribeMuskPromptsResponse, err error) {
    if request == nil {
        request = NewDescribeMuskPromptsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMuskPrompts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMuskPromptsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRegionsRequest() (request *DescribeRegionsRequest) {
    request = &DescribeRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hai", APIVersion, "DescribeRegions")
    
    
    return
}

func NewDescribeRegionsResponse() (response *DescribeRegionsResponse) {
    response = &DescribeRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRegions
// 本接口（DescribeRegions）用于查询地域列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_REGIONINVALID = "InvalidParameterValue.RegionInvalid"
func (c *Client) DescribeRegions(request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
    return c.DescribeRegionsWithContext(context.Background(), request)
}

// DescribeRegions
// 本接口（DescribeRegions）用于查询地域列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_REGIONINVALID = "InvalidParameterValue.RegionInvalid"
func (c *Client) DescribeRegionsWithContext(ctx context.Context, request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeRegionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRegions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScenesRequest() (request *DescribeScenesRequest) {
    request = &DescribeScenesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hai", APIVersion, "DescribeScenes")
    
    
    return
}

func NewDescribeScenesResponse() (response *DescribeScenesResponse) {
    response = &DescribeScenesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeScenes
// 本接口（DescribeScenes）用于查询场景
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INVALIDSCENEIDMALFORMED = "InvalidParameterValue.InvalidSceneIdMalformed"
//  INVALIDPARAMETERVALUE_REGIONINVALID = "InvalidParameterValue.RegionInvalid"
func (c *Client) DescribeScenes(request *DescribeScenesRequest) (response *DescribeScenesResponse, err error) {
    return c.DescribeScenesWithContext(context.Background(), request)
}

// DescribeScenes
// 本接口（DescribeScenes）用于查询场景
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INVALIDSCENEIDMALFORMED = "InvalidParameterValue.InvalidSceneIdMalformed"
//  INVALIDPARAMETERVALUE_REGIONINVALID = "InvalidParameterValue.RegionInvalid"
func (c *Client) DescribeScenesWithContext(ctx context.Context, request *DescribeScenesRequest) (response *DescribeScenesResponse, err error) {
    if request == nil {
        request = NewDescribeScenesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScenes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScenesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceLoginSettingsRequest() (request *DescribeServiceLoginSettingsRequest) {
    request = &DescribeServiceLoginSettingsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hai", APIVersion, "DescribeServiceLoginSettings")
    
    
    return
}

func NewDescribeServiceLoginSettingsResponse() (response *DescribeServiceLoginSettingsResponse) {
    response = &DescribeServiceLoginSettingsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeServiceLoginSettings
// 本接口（DescribeServiceLoginSettings）用于查询服务登录配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEIDMALFORMED = "InvalidParameterValue.InvalidInstanceIdMalformed"
func (c *Client) DescribeServiceLoginSettings(request *DescribeServiceLoginSettingsRequest) (response *DescribeServiceLoginSettingsResponse, err error) {
    return c.DescribeServiceLoginSettingsWithContext(context.Background(), request)
}

// DescribeServiceLoginSettings
// 本接口（DescribeServiceLoginSettings）用于查询服务登录配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEIDMALFORMED = "InvalidParameterValue.InvalidInstanceIdMalformed"
func (c *Client) DescribeServiceLoginSettingsWithContext(ctx context.Context, request *DescribeServiceLoginSettingsRequest) (response *DescribeServiceLoginSettingsResponse, err error) {
    if request == nil {
        request = NewDescribeServiceLoginSettingsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServiceLoginSettings require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServiceLoginSettingsResponse()
    err = c.Send(request, response)
    return
}

func NewInquirePriceRunInstancesRequest() (request *InquirePriceRunInstancesRequest) {
    request = &InquirePriceRunInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hai", APIVersion, "InquirePriceRunInstances")
    
    
    return
}

func NewInquirePriceRunInstancesResponse() (response *InquirePriceRunInstancesResponse) {
    response = &InquirePriceRunInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquirePriceRunInstances
// 本接口 (InquirePriceRunInstances) 用于实例询价。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_APPLICATIONIDNOTFOUND = "InvalidParameterValue.ApplicationIdNotFound"
//  INVALIDPARAMETERVALUE_BUNDLETYPENOTFOUND = "InvalidParameterValue.BundleTypeNotFound"
//  INVALIDPARAMETERVALUE_INSTANCENAMETOOLONG = "InvalidParameterValue.InstanceNameTooLong"
//  INVALIDPARAMETERVALUE_INVALIDAPPLICATIONIDMALFORMED = "InvalidParameterValue.InvalidApplicationIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDCHARGETYPE = "InvalidParameterValue.InvalidChargeType"
//  INVALIDPARAMETERVALUE_INVALIDCHARGETYPEANDTIMEUNIT = "InvalidParameterValue.InvalidChargeTypeAndTimeUnit"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCECOUNT = "InvalidParameterValue.InvalidInstanceCount"
//  INVALIDPARAMETERVALUE_SECURITYGROUPNOTFOUND = "InvalidParameterValue.SecurityGroupNotFound"
//  INVALIDPARAMETERVALUE_SUBNETMUSTCOEXISTWITHVPC = "InvalidParameterValue.SubnetMustCoexistWithVpc"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  RESOURCEINSUFFICIENT_BUNDLEINVENTORYSHORTAGE = "ResourceInsufficient.BundleInventoryShortage"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDSTSOPERATION = "UnauthorizedOperation.UnauthorizedSTSOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDUSER = "UnauthorizedOperation.UnauthorizedUser"
func (c *Client) InquirePriceRunInstances(request *InquirePriceRunInstancesRequest) (response *InquirePriceRunInstancesResponse, err error) {
    return c.InquirePriceRunInstancesWithContext(context.Background(), request)
}

// InquirePriceRunInstances
// 本接口 (InquirePriceRunInstances) 用于实例询价。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_APPLICATIONIDNOTFOUND = "InvalidParameterValue.ApplicationIdNotFound"
//  INVALIDPARAMETERVALUE_BUNDLETYPENOTFOUND = "InvalidParameterValue.BundleTypeNotFound"
//  INVALIDPARAMETERVALUE_INSTANCENAMETOOLONG = "InvalidParameterValue.InstanceNameTooLong"
//  INVALIDPARAMETERVALUE_INVALIDAPPLICATIONIDMALFORMED = "InvalidParameterValue.InvalidApplicationIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDCHARGETYPE = "InvalidParameterValue.InvalidChargeType"
//  INVALIDPARAMETERVALUE_INVALIDCHARGETYPEANDTIMEUNIT = "InvalidParameterValue.InvalidChargeTypeAndTimeUnit"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCECOUNT = "InvalidParameterValue.InvalidInstanceCount"
//  INVALIDPARAMETERVALUE_SECURITYGROUPNOTFOUND = "InvalidParameterValue.SecurityGroupNotFound"
//  INVALIDPARAMETERVALUE_SUBNETMUSTCOEXISTWITHVPC = "InvalidParameterValue.SubnetMustCoexistWithVpc"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  RESOURCEINSUFFICIENT_BUNDLEINVENTORYSHORTAGE = "ResourceInsufficient.BundleInventoryShortage"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDSTSOPERATION = "UnauthorizedOperation.UnauthorizedSTSOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDUSER = "UnauthorizedOperation.UnauthorizedUser"
func (c *Client) InquirePriceRunInstancesWithContext(ctx context.Context, request *InquirePriceRunInstancesRequest) (response *InquirePriceRunInstancesResponse, err error) {
    if request == nil {
        request = NewInquirePriceRunInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquirePriceRunInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquirePriceRunInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewResetInstancesPasswordRequest() (request *ResetInstancesPasswordRequest) {
    request = &ResetInstancesPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hai", APIVersion, "ResetInstancesPassword")
    
    
    return
}

func NewResetInstancesPasswordResponse() (response *ResetInstancesPasswordResponse) {
    response = &ResetInstancesPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResetInstancesPassword
// 本接口 (ResetInstancesPassword) 用于重置实例的用户密码。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDNOTFOUND = "InvalidParameterValue.InstanceIdNotFound"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEIDMALFORMED = "InvalidParameterValue.InvalidInstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORD = "InvalidParameterValue.InvalidPassword"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
func (c *Client) ResetInstancesPassword(request *ResetInstancesPasswordRequest) (response *ResetInstancesPasswordResponse, err error) {
    return c.ResetInstancesPasswordWithContext(context.Background(), request)
}

// ResetInstancesPassword
// 本接口 (ResetInstancesPassword) 用于重置实例的用户密码。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDNOTFOUND = "InvalidParameterValue.InstanceIdNotFound"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEIDMALFORMED = "InvalidParameterValue.InvalidInstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORD = "InvalidParameterValue.InvalidPassword"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
func (c *Client) ResetInstancesPasswordWithContext(ctx context.Context, request *ResetInstancesPasswordRequest) (response *ResetInstancesPasswordResponse, err error) {
    if request == nil {
        request = NewResetInstancesPasswordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetInstancesPassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetInstancesPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewResizeInstanceDiskRequest() (request *ResizeInstanceDiskRequest) {
    request = &ResizeInstanceDiskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hai", APIVersion, "ResizeInstanceDisk")
    
    
    return
}

func NewResizeInstanceDiskResponse() (response *ResizeInstanceDiskResponse) {
    response = &ResizeInstanceDiskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResizeInstanceDisk
// 本接口（ResizeInstanceDisk）用于对指定HAI实例进行扩容云硬盘操作。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DISKSIZEMUSTLARGETHANCURRENTDISKSIZE = "InvalidParameterValue.DiskSizeMustLargeThanCurrentDiskSize"
//  INVALIDPARAMETERVALUE_DISKSIZEREACHLIMIT = "InvalidParameterValue.DiskSizeReachLimit"
//  INVALIDPARAMETERVALUE_INSTANCEIDNOTFOUND = "InvalidParameterValue.InstanceIdNotFound"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEIDMALFORMED = "InvalidParameterValue.InvalidInstanceIdMalformed"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
func (c *Client) ResizeInstanceDisk(request *ResizeInstanceDiskRequest) (response *ResizeInstanceDiskResponse, err error) {
    return c.ResizeInstanceDiskWithContext(context.Background(), request)
}

// ResizeInstanceDisk
// 本接口（ResizeInstanceDisk）用于对指定HAI实例进行扩容云硬盘操作。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DISKSIZEMUSTLARGETHANCURRENTDISKSIZE = "InvalidParameterValue.DiskSizeMustLargeThanCurrentDiskSize"
//  INVALIDPARAMETERVALUE_DISKSIZEREACHLIMIT = "InvalidParameterValue.DiskSizeReachLimit"
//  INVALIDPARAMETERVALUE_INSTANCEIDNOTFOUND = "InvalidParameterValue.InstanceIdNotFound"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEIDMALFORMED = "InvalidParameterValue.InvalidInstanceIdMalformed"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
func (c *Client) ResizeInstanceDiskWithContext(ctx context.Context, request *ResizeInstanceDiskRequest) (response *ResizeInstanceDiskResponse, err error) {
    if request == nil {
        request = NewResizeInstanceDiskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResizeInstanceDisk require credential")
    }

    request.SetContext(ctx)
    
    response = NewResizeInstanceDiskResponse()
    err = c.Send(request, response)
    return
}

func NewRunInstancesRequest() (request *RunInstancesRequest) {
    request = &RunInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hai", APIVersion, "RunInstances")
    
    
    return
}

func NewRunInstancesResponse() (response *RunInstancesResponse) {
    response = &RunInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RunInstances
// 本接口 (RunInstances) 用于创建一个或多个指定配置的实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ARREARSACCOUNTCANNOTRUNINSTANCES = "FailedOperation.ArrearsAccountCannotRunInstances"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INVALIDPARAMETERVALUE_APPLICATIONIDNOTFOUND = "InvalidParameterValue.ApplicationIdNotFound"
//  INVALIDPARAMETERVALUE_BUNDLETYPENOTFOUND = "InvalidParameterValue.BundleTypeNotFound"
//  INVALIDPARAMETERVALUE_DISKSIZEMUSTLARGERTHANAPPLICATIONMINREQUIREDSIZE = "InvalidParameterValue.DiskSizeMustLargerThanApplicationMinRequiredSize"
//  INVALIDPARAMETERVALUE_INSTANCENAMETOOLONG = "InvalidParameterValue.InstanceNameTooLong"
//  INVALIDPARAMETERVALUE_INVALIDAPPLICATIONIDMALFORMED = "InvalidParameterValue.InvalidApplicationIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDCHARGETYPE = "InvalidParameterValue.InvalidChargeType"
//  INVALIDPARAMETERVALUE_INVALIDCHARGETYPEANDTIMEUNIT = "InvalidParameterValue.InvalidChargeTypeAndTimeUnit"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCECOUNT = "InvalidParameterValue.InvalidInstanceCount"
//  INVALIDPARAMETERVALUE_SECURITYGROUPNOTFOUND = "InvalidParameterValue.SecurityGroupNotFound"
//  INVALIDPARAMETERVALUE_SUBNETMUSTCOEXISTWITHVPC = "InvalidParameterValue.SubnetMustCoexistWithVpc"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_SECURITYGROUPLIMITEXCEEDED = "LimitExceeded.SecurityGroupLimitExceeded"
//  LIMITEXCEEDED_VPCLIMITEXCEEDED = "LimitExceeded.VpcLimitExceeded"
//  RESOURCEINSUFFICIENT_BUNDLEINVENTORYSHORTAGE = "ResourceInsufficient.BundleInventoryShortage"
//  UNAUTHORIZEDOPERATION_GETROLEERROR = "UnauthorizedOperation.GetRoleError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDNETWORKUSER = "UnauthorizedOperation.UnauthorizedNetworkUser"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDSTSOPERATION = "UnauthorizedOperation.UnauthorizedSTSOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDUSER = "UnauthorizedOperation.UnauthorizedUser"
func (c *Client) RunInstances(request *RunInstancesRequest) (response *RunInstancesResponse, err error) {
    return c.RunInstancesWithContext(context.Background(), request)
}

// RunInstances
// 本接口 (RunInstances) 用于创建一个或多个指定配置的实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ARREARSACCOUNTCANNOTRUNINSTANCES = "FailedOperation.ArrearsAccountCannotRunInstances"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INVALIDPARAMETERVALUE_APPLICATIONIDNOTFOUND = "InvalidParameterValue.ApplicationIdNotFound"
//  INVALIDPARAMETERVALUE_BUNDLETYPENOTFOUND = "InvalidParameterValue.BundleTypeNotFound"
//  INVALIDPARAMETERVALUE_DISKSIZEMUSTLARGERTHANAPPLICATIONMINREQUIREDSIZE = "InvalidParameterValue.DiskSizeMustLargerThanApplicationMinRequiredSize"
//  INVALIDPARAMETERVALUE_INSTANCENAMETOOLONG = "InvalidParameterValue.InstanceNameTooLong"
//  INVALIDPARAMETERVALUE_INVALIDAPPLICATIONIDMALFORMED = "InvalidParameterValue.InvalidApplicationIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDCHARGETYPE = "InvalidParameterValue.InvalidChargeType"
//  INVALIDPARAMETERVALUE_INVALIDCHARGETYPEANDTIMEUNIT = "InvalidParameterValue.InvalidChargeTypeAndTimeUnit"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCECOUNT = "InvalidParameterValue.InvalidInstanceCount"
//  INVALIDPARAMETERVALUE_SECURITYGROUPNOTFOUND = "InvalidParameterValue.SecurityGroupNotFound"
//  INVALIDPARAMETERVALUE_SUBNETMUSTCOEXISTWITHVPC = "InvalidParameterValue.SubnetMustCoexistWithVpc"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_SECURITYGROUPLIMITEXCEEDED = "LimitExceeded.SecurityGroupLimitExceeded"
//  LIMITEXCEEDED_VPCLIMITEXCEEDED = "LimitExceeded.VpcLimitExceeded"
//  RESOURCEINSUFFICIENT_BUNDLEINVENTORYSHORTAGE = "ResourceInsufficient.BundleInventoryShortage"
//  UNAUTHORIZEDOPERATION_GETROLEERROR = "UnauthorizedOperation.GetRoleError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDNETWORKUSER = "UnauthorizedOperation.UnauthorizedNetworkUser"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDSTSOPERATION = "UnauthorizedOperation.UnauthorizedSTSOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDUSER = "UnauthorizedOperation.UnauthorizedUser"
func (c *Client) RunInstancesWithContext(ctx context.Context, request *RunInstancesRequest) (response *RunInstancesResponse, err error) {
    if request == nil {
        request = NewRunInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewRunInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewStartInstanceRequest() (request *StartInstanceRequest) {
    request = &StartInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hai", APIVersion, "StartInstance")
    
    
    return
}

func NewStartInstanceResponse() (response *StartInstanceResponse) {
    response = &StartInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartInstance
// 本接口 (StartInstance) 用于主动启动实例。
//
// ‘运行中’、‘预付费’的实例不支持启动实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATEINSTANCEREPEATEDLY = "FailedOperation.OperateInstanceRepeatedly"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDNOTFOUND = "InvalidParameterValue.InstanceIdNotFound"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEIDMALFORMED = "InvalidParameterValue.InvalidInstanceIdMalformed"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCEINSUFFICIENT_BUNDLEINVENTORYSHORTAGE = "ResourceInsufficient.BundleInventoryShortage"
//  UNSUPPORTEDOPERATION_INSTANCESTATEARREARS = "UnsupportedOperation.InstanceStateArrears"
//  UNSUPPORTEDOPERATION_INSTANCESTATELAUNCHFAILED = "UnsupportedOperation.InstanceStateLaunchFailed"
//  UNSUPPORTEDOPERATION_INSTANCESTATEPENDING = "UnsupportedOperation.InstanceStatePending"
//  UNSUPPORTEDOPERATION_INSTANCESTATERUNNING = "UnsupportedOperation.InstanceStateRunning"
//  UNSUPPORTEDOPERATION_INSTANCESTATETERMINATED = "UnsupportedOperation.InstanceStateTerminated"
func (c *Client) StartInstance(request *StartInstanceRequest) (response *StartInstanceResponse, err error) {
    return c.StartInstanceWithContext(context.Background(), request)
}

// StartInstance
// 本接口 (StartInstance) 用于主动启动实例。
//
// ‘运行中’、‘预付费’的实例不支持启动实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATEINSTANCEREPEATEDLY = "FailedOperation.OperateInstanceRepeatedly"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDNOTFOUND = "InvalidParameterValue.InstanceIdNotFound"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEIDMALFORMED = "InvalidParameterValue.InvalidInstanceIdMalformed"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCEINSUFFICIENT_BUNDLEINVENTORYSHORTAGE = "ResourceInsufficient.BundleInventoryShortage"
//  UNSUPPORTEDOPERATION_INSTANCESTATEARREARS = "UnsupportedOperation.InstanceStateArrears"
//  UNSUPPORTEDOPERATION_INSTANCESTATELAUNCHFAILED = "UnsupportedOperation.InstanceStateLaunchFailed"
//  UNSUPPORTEDOPERATION_INSTANCESTATEPENDING = "UnsupportedOperation.InstanceStatePending"
//  UNSUPPORTEDOPERATION_INSTANCESTATERUNNING = "UnsupportedOperation.InstanceStateRunning"
//  UNSUPPORTEDOPERATION_INSTANCESTATETERMINATED = "UnsupportedOperation.InstanceStateTerminated"
func (c *Client) StartInstanceWithContext(ctx context.Context, request *StartInstanceRequest) (response *StartInstanceResponse, err error) {
    if request == nil {
        request = NewStartInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewStopInstanceRequest() (request *StopInstanceRequest) {
    request = &StopInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hai", APIVersion, "StopInstance")
    
    
    return
}

func NewStopInstanceResponse() (response *StopInstanceResponse) {
    response = &StopInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopInstance
// 本接口 (StopInstance) 用于主动关闭实例。
//
// ‘已关机’、‘预付费’的实例不支持关机
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATEINSTANCEREPEATEDLY = "FailedOperation.OperateInstanceRepeatedly"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDNOTFOUND = "InvalidParameterValue.InstanceIdNotFound"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEIDMALFORMED = "InvalidParameterValue.InvalidInstanceIdMalformed"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  UNSUPPORTEDOPERATION_INSTANCESTATEARREARS = "UnsupportedOperation.InstanceStateArrears"
//  UNSUPPORTEDOPERATION_INSTANCESTATELAUNCHFAILED = "UnsupportedOperation.InstanceStateLaunchFailed"
//  UNSUPPORTEDOPERATION_INSTANCESTATEPENDING = "UnsupportedOperation.InstanceStatePending"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTOPPEDNOCHARGE = "UnsupportedOperation.InstanceStateStoppedNoCharge"
//  UNSUPPORTEDOPERATION_INSTANCESTATETERMINATED = "UnsupportedOperation.InstanceStateTerminated"
func (c *Client) StopInstance(request *StopInstanceRequest) (response *StopInstanceResponse, err error) {
    return c.StopInstanceWithContext(context.Background(), request)
}

// StopInstance
// 本接口 (StopInstance) 用于主动关闭实例。
//
// ‘已关机’、‘预付费’的实例不支持关机
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATEINSTANCEREPEATEDLY = "FailedOperation.OperateInstanceRepeatedly"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDNOTFOUND = "InvalidParameterValue.InstanceIdNotFound"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEIDMALFORMED = "InvalidParameterValue.InvalidInstanceIdMalformed"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  UNSUPPORTEDOPERATION_INSTANCESTATEARREARS = "UnsupportedOperation.InstanceStateArrears"
//  UNSUPPORTEDOPERATION_INSTANCESTATELAUNCHFAILED = "UnsupportedOperation.InstanceStateLaunchFailed"
//  UNSUPPORTEDOPERATION_INSTANCESTATEPENDING = "UnsupportedOperation.InstanceStatePending"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTOPPEDNOCHARGE = "UnsupportedOperation.InstanceStateStoppedNoCharge"
//  UNSUPPORTEDOPERATION_INSTANCESTATETERMINATED = "UnsupportedOperation.InstanceStateTerminated"
func (c *Client) StopInstanceWithContext(ctx context.Context, request *StopInstanceRequest) (response *StopInstanceResponse, err error) {
    if request == nil {
        request = NewStopInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateInstancesRequest() (request *TerminateInstancesRequest) {
    request = &TerminateInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hai", APIVersion, "TerminateInstances")
    
    
    return
}

func NewTerminateInstancesResponse() (response *TerminateInstancesResponse) {
    response = &TerminateInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TerminateInstances
// 本接口 (TerminateInstances) 用于主动退还实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CANNOTRETURN = "FailedOperation.CannotReturn"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDNOTFOUND = "InvalidParameterValue.InstanceIdNotFound"
//  INVALIDPARAMETERVALUE_INVALIDCHARGETYPE = "InvalidParameterValue.InvalidChargeType"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEIDMALFORMED = "InvalidParameterValue.InvalidInstanceIdMalformed"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  UNSUPPORTEDOPERATION_CANNOTTERMINATEEXPIREDANDNOTEXPIREDINSTANCES = "UnsupportedOperation.CannotTerminateExpiredAndNotExpiredInstances"
//  UNSUPPORTEDOPERATION_NOTSUPPORTTERMINATEPREPAIDANDPOSTPAID = "UnsupportedOperation.NotSupportTerminatePrePaidAndPostPaid"
func (c *Client) TerminateInstances(request *TerminateInstancesRequest) (response *TerminateInstancesResponse, err error) {
    return c.TerminateInstancesWithContext(context.Background(), request)
}

// TerminateInstances
// 本接口 (TerminateInstances) 用于主动退还实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CANNOTRETURN = "FailedOperation.CannotReturn"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDNOTFOUND = "InvalidParameterValue.InstanceIdNotFound"
//  INVALIDPARAMETERVALUE_INVALIDCHARGETYPE = "InvalidParameterValue.InvalidChargeType"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEIDMALFORMED = "InvalidParameterValue.InvalidInstanceIdMalformed"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  UNSUPPORTEDOPERATION_CANNOTTERMINATEEXPIREDANDNOTEXPIREDINSTANCES = "UnsupportedOperation.CannotTerminateExpiredAndNotExpiredInstances"
//  UNSUPPORTEDOPERATION_NOTSUPPORTTERMINATEPREPAIDANDPOSTPAID = "UnsupportedOperation.NotSupportTerminatePrePaidAndPostPaid"
func (c *Client) TerminateInstancesWithContext(ctx context.Context, request *TerminateInstancesRequest) (response *TerminateInstancesResponse, err error) {
    if request == nil {
        request = NewTerminateInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminateInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewTerminateInstancesResponse()
    err = c.Send(request, response)
    return
}
