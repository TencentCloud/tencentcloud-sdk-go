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

package v20201104

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-11-04"

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


func NewAssociateSecurityGroupRequest() (request *AssociateSecurityGroupRequest) {
    request = &AssociateSecurityGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "AssociateSecurityGroup")
    
    
    return
}

func NewAssociateSecurityGroupResponse() (response *AssociateSecurityGroupResponse) {
    response = &AssociateSecurityGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AssociateSecurityGroup
// 安全组批量关联云桌面
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) AssociateSecurityGroup(request *AssociateSecurityGroupRequest) (response *AssociateSecurityGroupResponse, err error) {
    return c.AssociateSecurityGroupWithContext(context.Background(), request)
}

// AssociateSecurityGroup
// 安全组批量关联云桌面
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) AssociateSecurityGroupWithContext(ctx context.Context, request *AssociateSecurityGroupRequest) (response *AssociateSecurityGroupResponse, err error) {
    if request == nil {
        request = NewAssociateSecurityGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "AssociateSecurityGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AssociateSecurityGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewAssociateSecurityGroupResponse()
    err = c.Send(request, response)
    return
}

func NewBindInstancesRequest() (request *BindInstancesRequest) {
    request = &BindInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "BindInstances")
    
    
    return
}

func NewBindInstancesResponse() (response *BindInstancesResponse) {
    response = &BindInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindInstances
// 绑定桌面
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) BindInstances(request *BindInstancesRequest) (response *BindInstancesResponse, err error) {
    return c.BindInstancesWithContext(context.Background(), request)
}

// BindInstances
// 绑定桌面
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) BindInstancesWithContext(ctx context.Context, request *BindInstancesRequest) (response *BindInstancesResponse, err error) {
    if request == nil {
        request = NewBindInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "BindInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCommandTaskRequest() (request *CreateCommandTaskRequest) {
    request = &CreateCommandTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "CreateCommandTask")
    
    
    return
}

func NewCreateCommandTaskResponse() (response *CreateCommandTaskResponse) {
    response = &CreateCommandTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCommandTask
// 创建批量脚本任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateCommandTask(request *CreateCommandTaskRequest) (response *CreateCommandTaskResponse, err error) {
    return c.CreateCommandTaskWithContext(context.Background(), request)
}

// CreateCommandTask
// 创建批量脚本任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateCommandTaskWithContext(ctx context.Context, request *CreateCommandTaskRequest) (response *CreateCommandTaskResponse, err error) {
    if request == nil {
        request = NewCreateCommandTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "CreateCommandTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCommandTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCommandTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDataDiskRequest() (request *CreateDataDiskRequest) {
    request = &CreateDataDiskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "CreateDataDisk")
    
    
    return
}

func NewCreateDataDiskResponse() (response *CreateDataDiskResponse) {
    response = &CreateDataDiskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDataDisk
// 实例新增数据盘
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateDataDisk(request *CreateDataDiskRequest) (response *CreateDataDiskResponse, err error) {
    return c.CreateDataDiskWithContext(context.Background(), request)
}

// CreateDataDisk
// 实例新增数据盘
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateDataDiskWithContext(ctx context.Context, request *CreateDataDiskRequest) (response *CreateDataDiskResponse, err error) {
    if request == nil {
        request = NewCreateDataDiskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "CreateDataDisk")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDataDisk require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDataDiskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDomainControlPolicyRequest() (request *CreateDomainControlPolicyRequest) {
    request = &CreateDomainControlPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "CreateDomainControlPolicy")
    
    
    return
}

func NewCreateDomainControlPolicyResponse() (response *CreateDomainControlPolicyResponse) {
    response = &CreateDomainControlPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDomainControlPolicy
// 新建域名访问控制策略
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_TOOMANYPOLICIES = "LimitExceeded.TooManyPolicies"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateDomainControlPolicy(request *CreateDomainControlPolicyRequest) (response *CreateDomainControlPolicyResponse, err error) {
    return c.CreateDomainControlPolicyWithContext(context.Background(), request)
}

// CreateDomainControlPolicy
// 新建域名访问控制策略
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_TOOMANYPOLICIES = "LimitExceeded.TooManyPolicies"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateDomainControlPolicyWithContext(ctx context.Context, request *CreateDomainControlPolicyRequest) (response *CreateDomainControlPolicyResponse, err error) {
    if request == nil {
        request = NewCreateDomainControlPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "CreateDomainControlPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDomainControlPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDomainControlPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateImageRequest() (request *CreateImageRequest) {
    request = &CreateImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "CreateImage")
    
    
    return
}

func NewCreateImageResponse() (response *CreateImageResponse) {
    response = &CreateImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateImage
// 创建自定义镜像。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DISK = "InvalidParameterValue.Disk"
//  INVALIDPARAMETERVALUE_INSTANCEID = "InvalidParameterValue.InstanceId"
//  LIMITEXCEEDED_IMAGEQUOTA = "LimitExceeded.ImageQuota"
//  UNSUPPORTEDOPERATION_ASSISTANTSTATE = "UnsupportedOperation.AssistantState"
//  UNSUPPORTEDOPERATION_DISKSIZE = "UnsupportedOperation.DiskSize"
//  UNSUPPORTEDOPERATION_INSTANCESTATE = "UnsupportedOperation.InstanceState"
//  UNSUPPORTEDOPERATION_STOPPEDMODESTOPCHARGING = "UnsupportedOperation.StoppedModeStopCharging"
func (c *Client) CreateImage(request *CreateImageRequest) (response *CreateImageResponse, err error) {
    return c.CreateImageWithContext(context.Background(), request)
}

// CreateImage
// 创建自定义镜像。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DISK = "InvalidParameterValue.Disk"
//  INVALIDPARAMETERVALUE_INSTANCEID = "InvalidParameterValue.InstanceId"
//  LIMITEXCEEDED_IMAGEQUOTA = "LimitExceeded.ImageQuota"
//  UNSUPPORTEDOPERATION_ASSISTANTSTATE = "UnsupportedOperation.AssistantState"
//  UNSUPPORTEDOPERATION_DISKSIZE = "UnsupportedOperation.DiskSize"
//  UNSUPPORTEDOPERATION_INSTANCESTATE = "UnsupportedOperation.InstanceState"
//  UNSUPPORTEDOPERATION_STOPPEDMODESTOPCHARGING = "UnsupportedOperation.StoppedModeStopCharging"
func (c *Client) CreateImageWithContext(ctx context.Context, request *CreateImageRequest) (response *CreateImageResponse, err error) {
    if request == nil {
        request = NewCreateImageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "CreateImage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateImage require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateImageResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstancesRequest() (request *CreateInstancesRequest) {
    request = &CreateInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "CreateInstances")
    
    
    return
}

func NewCreateInstancesResponse() (response *CreateInstancesResponse) {
    response = &CreateInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateInstances
// 创建云桌面实例
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION_DEALCHECKFAILED = "FailedOperation.DealCheckFailed"
//  FAILEDOPERATION_INSUFFICIENTBALANCE = "FailedOperation.InsufficientBalance"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DISK = "InvalidParameterValue.Disk"
//  INVALIDPARAMETERVALUE_IMAGE = "InvalidParameterValue.Image"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  LIMITEXCEEDED_INSTANCEQUOTA = "LimitExceeded.InstanceQuota"
//  LIMITEXCEEDED_PARALLELOPERATION = "LimitExceeded.ParallelOperation"
//  RESOURCEINSUFFICIENT_SPECIFIEDINSTANCETYPE = "ResourceInsufficient.SpecifiedInstanceType"
//  UNAUTHORIZEDOPERATION_SERVICENOTACTIVATED = "UnauthorizedOperation.ServiceNotActivated"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INSTANCESTATE = "UnsupportedOperation.InstanceState"
func (c *Client) CreateInstances(request *CreateInstancesRequest) (response *CreateInstancesResponse, err error) {
    return c.CreateInstancesWithContext(context.Background(), request)
}

// CreateInstances
// 创建云桌面实例
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION_DEALCHECKFAILED = "FailedOperation.DealCheckFailed"
//  FAILEDOPERATION_INSUFFICIENTBALANCE = "FailedOperation.InsufficientBalance"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DISK = "InvalidParameterValue.Disk"
//  INVALIDPARAMETERVALUE_IMAGE = "InvalidParameterValue.Image"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  LIMITEXCEEDED_INSTANCEQUOTA = "LimitExceeded.InstanceQuota"
//  LIMITEXCEEDED_PARALLELOPERATION = "LimitExceeded.ParallelOperation"
//  RESOURCEINSUFFICIENT_SPECIFIEDINSTANCETYPE = "ResourceInsufficient.SpecifiedInstanceType"
//  UNAUTHORIZEDOPERATION_SERVICENOTACTIVATED = "UnauthorizedOperation.ServiceNotActivated"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INSTANCESTATE = "UnsupportedOperation.InstanceState"
func (c *Client) CreateInstancesWithContext(ctx context.Context, request *CreateInstancesRequest) (response *CreateInstancesResponse, err error) {
    if request == nil {
        request = NewCreateInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "CreateInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePolicyRequest() (request *CreatePolicyRequest) {
    request = &CreatePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "CreatePolicy")
    
    
    return
}

func NewCreatePolicyResponse() (response *CreatePolicyResponse) {
    response = &CreatePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePolicy
// 创建策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TOOMANYPOLICIES = "LimitExceeded.TooManyPolicies"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_POLICYEXISTS = "ResourceInUse.PolicyExists"
func (c *Client) CreatePolicy(request *CreatePolicyRequest) (response *CreatePolicyResponse, err error) {
    return c.CreatePolicyWithContext(context.Background(), request)
}

// CreatePolicy
// 创建策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TOOMANYPOLICIES = "LimitExceeded.TooManyPolicies"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_POLICYEXISTS = "ResourceInUse.PolicyExists"
func (c *Client) CreatePolicyWithContext(ctx context.Context, request *CreatePolicyRequest) (response *CreatePolicyResponse, err error) {
    if request == nil {
        request = NewCreatePolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "CreatePolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCommandTasksRequest() (request *DeleteCommandTasksRequest) {
    request = &DeleteCommandTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "DeleteCommandTasks")
    
    
    return
}

func NewDeleteCommandTasksResponse() (response *DeleteCommandTasksResponse) {
    response = &DeleteCommandTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCommandTasks
// 删除命令任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteCommandTasks(request *DeleteCommandTasksRequest) (response *DeleteCommandTasksResponse, err error) {
    return c.DeleteCommandTasksWithContext(context.Background(), request)
}

// DeleteCommandTasks
// 删除命令任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteCommandTasksWithContext(ctx context.Context, request *DeleteCommandTasksRequest) (response *DeleteCommandTasksResponse, err error) {
    if request == nil {
        request = NewDeleteCommandTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "DeleteCommandTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCommandTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCommandTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDomainControlPolicyRequest() (request *DeleteDomainControlPolicyRequest) {
    request = &DeleteDomainControlPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "DeleteDomainControlPolicy")
    
    
    return
}

func NewDeleteDomainControlPolicyResponse() (response *DeleteDomainControlPolicyResponse) {
    response = &DeleteDomainControlPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDomainControlPolicy
// 删除域名访问控制策略
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteDomainControlPolicy(request *DeleteDomainControlPolicyRequest) (response *DeleteDomainControlPolicyResponse, err error) {
    return c.DeleteDomainControlPolicyWithContext(context.Background(), request)
}

// DeleteDomainControlPolicy
// 删除域名访问控制策略
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteDomainControlPolicyWithContext(ctx context.Context, request *DeleteDomainControlPolicyRequest) (response *DeleteDomainControlPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteDomainControlPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "DeleteDomainControlPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDomainControlPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDomainControlPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteImagesRequest() (request *DeleteImagesRequest) {
    request = &DeleteImagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "DeleteImages")
    
    
    return
}

func NewDeleteImagesResponse() (response *DeleteImagesResponse) {
    response = &DeleteImagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteImages
// 用于删除一个或多个自定义镜像。单次最大删除5个。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_IMAGESTATE = "UnsupportedOperation.ImageState"
//  UNSUPPORTEDOPERATION_IMAGETYPE = "UnsupportedOperation.ImageType"
func (c *Client) DeleteImages(request *DeleteImagesRequest) (response *DeleteImagesResponse, err error) {
    return c.DeleteImagesWithContext(context.Background(), request)
}

// DeleteImages
// 用于删除一个或多个自定义镜像。单次最大删除5个。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_IMAGESTATE = "UnsupportedOperation.ImageState"
//  UNSUPPORTEDOPERATION_IMAGETYPE = "UnsupportedOperation.ImageType"
func (c *Client) DeleteImagesWithContext(ctx context.Context, request *DeleteImagesRequest) (response *DeleteImagesResponse, err error) {
    if request == nil {
        request = NewDeleteImagesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "DeleteImages")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteImages require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteImagesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCommandTasksRequest() (request *DescribeCommandTasksRequest) {
    request = &DescribeCommandTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "DescribeCommandTasks")
    
    
    return
}

func NewDescribeCommandTasksResponse() (response *DescribeCommandTasksResponse) {
    response = &DescribeCommandTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCommandTasks
// 拉取命令任务详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeCommandTasks(request *DescribeCommandTasksRequest) (response *DescribeCommandTasksResponse, err error) {
    return c.DescribeCommandTasksWithContext(context.Background(), request)
}

// DescribeCommandTasks
// 拉取命令任务详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeCommandTasksWithContext(ctx context.Context, request *DescribeCommandTasksRequest) (response *DescribeCommandTasksResponse, err error) {
    if request == nil {
        request = NewDescribeCommandTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "DescribeCommandTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCommandTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCommandTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImagesRequest() (request *DescribeImagesRequest) {
    request = &DescribeImagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "DescribeImages")
    
    
    return
}

func NewDescribeImagesResponse() (response *DescribeImagesResponse) {
    response = &DescribeImagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeImages
// 获取云桌面镜像详细信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DescribeImages(request *DescribeImagesRequest) (response *DescribeImagesResponse, err error) {
    return c.DescribeImagesWithContext(context.Background(), request)
}

// DescribeImages
// 获取云桌面镜像详细信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DescribeImagesWithContext(ctx context.Context, request *DescribeImagesRequest) (response *DescribeImagesResponse, err error) {
    if request == nil {
        request = NewDescribeImagesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "DescribeImages")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImages require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImagesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceFetchMsgRequest() (request *DescribeInstanceFetchMsgRequest) {
    request = &DescribeInstanceFetchMsgRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "DescribeInstanceFetchMsg")
    
    
    return
}

func NewDescribeInstanceFetchMsgResponse() (response *DescribeInstanceFetchMsgResponse) {
    response = &DescribeInstanceFetchMsgResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceFetchMsg
// 获取指定实例的登录Fetch信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCEINUSE = "FailedOperation.InstanceInUse"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INTERNALMODULEERROR = "FailedOperation.InternalModuleError"
//  FAILEDOPERATION_INVALIDINSTANCEBOUNDUSER = "FailedOperation.InvalidInstanceBoundUser"
//  FAILEDOPERATION_INVALIDINSTANCEREGION = "FailedOperation.InvalidInstanceRegion"
//  FAILEDOPERATION_INVALIDINSTANCESTATUSOPERATION = "FailedOperation.InvalidInstanceStatusOperation"
//  FAILEDOPERATION_NOSUPPORTINSTANCEREGION = "FailedOperation.NoSupportInstanceRegion"
//  FAILEDOPERATION_UNKNOWNERROR = "FailedOperation.UnknownError"
//  FAILEDOPERATION_USERRESOURCENOTFOUND = "FailedOperation.UserResourceNotFound"
func (c *Client) DescribeInstanceFetchMsg(request *DescribeInstanceFetchMsgRequest) (response *DescribeInstanceFetchMsgResponse, err error) {
    return c.DescribeInstanceFetchMsgWithContext(context.Background(), request)
}

// DescribeInstanceFetchMsg
// 获取指定实例的登录Fetch信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCEINUSE = "FailedOperation.InstanceInUse"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INTERNALMODULEERROR = "FailedOperation.InternalModuleError"
//  FAILEDOPERATION_INVALIDINSTANCEBOUNDUSER = "FailedOperation.InvalidInstanceBoundUser"
//  FAILEDOPERATION_INVALIDINSTANCEREGION = "FailedOperation.InvalidInstanceRegion"
//  FAILEDOPERATION_INVALIDINSTANCESTATUSOPERATION = "FailedOperation.InvalidInstanceStatusOperation"
//  FAILEDOPERATION_NOSUPPORTINSTANCEREGION = "FailedOperation.NoSupportInstanceRegion"
//  FAILEDOPERATION_UNKNOWNERROR = "FailedOperation.UnknownError"
//  FAILEDOPERATION_USERRESOURCENOTFOUND = "FailedOperation.UserResourceNotFound"
func (c *Client) DescribeInstanceFetchMsgWithContext(ctx context.Context, request *DescribeInstanceFetchMsgRequest) (response *DescribeInstanceFetchMsgResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceFetchMsgRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "DescribeInstanceFetchMsg")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceFetchMsg require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceFetchMsgResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceGraphDataRequest() (request *DescribeInstanceGraphDataRequest) {
    request = &DescribeInstanceGraphDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "DescribeInstanceGraphData")
    
    
    return
}

func NewDescribeInstanceGraphDataResponse() (response *DescribeInstanceGraphDataResponse) {
    response = &DescribeInstanceGraphDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceGraphData
// 获取实例磁盘分区监控数据
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeInstanceGraphData(request *DescribeInstanceGraphDataRequest) (response *DescribeInstanceGraphDataResponse, err error) {
    return c.DescribeInstanceGraphDataWithContext(context.Background(), request)
}

// DescribeInstanceGraphData
// 获取实例磁盘分区监控数据
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeInstanceGraphDataWithContext(ctx context.Context, request *DescribeInstanceGraphDataRequest) (response *DescribeInstanceGraphDataResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceGraphDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "DescribeInstanceGraphData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceGraphData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceGraphDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceLoginMsgRequest() (request *DescribeInstanceLoginMsgRequest) {
    request = &DescribeInstanceLoginMsgRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "DescribeInstanceLoginMsg")
    
    
    return
}

func NewDescribeInstanceLoginMsgResponse() (response *DescribeInstanceLoginMsgResponse) {
    response = &DescribeInstanceLoginMsgResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceLoginMsg
// 获取指定实例的登录login信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCEINUSE = "FailedOperation.InstanceInUse"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INTERNALMODULEERROR = "FailedOperation.InternalModuleError"
//  FAILEDOPERATION_INVALIDINSTANCEBOUNDUSER = "FailedOperation.InvalidInstanceBoundUser"
//  FAILEDOPERATION_INVALIDINSTANCEREGION = "FailedOperation.InvalidInstanceRegion"
//  FAILEDOPERATION_INVALIDINSTANCESTATUSOPERATION = "FailedOperation.InvalidInstanceStatusOperation"
//  FAILEDOPERATION_NOSUPPORTINSTANCEREGION = "FailedOperation.NoSupportInstanceRegion"
//  FAILEDOPERATION_UNKNOWNERROR = "FailedOperation.UnknownError"
//  FAILEDOPERATION_USERRESOURCENOTFOUND = "FailedOperation.UserResourceNotFound"
func (c *Client) DescribeInstanceLoginMsg(request *DescribeInstanceLoginMsgRequest) (response *DescribeInstanceLoginMsgResponse, err error) {
    return c.DescribeInstanceLoginMsgWithContext(context.Background(), request)
}

// DescribeInstanceLoginMsg
// 获取指定实例的登录login信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCEINUSE = "FailedOperation.InstanceInUse"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INTERNALMODULEERROR = "FailedOperation.InternalModuleError"
//  FAILEDOPERATION_INVALIDINSTANCEBOUNDUSER = "FailedOperation.InvalidInstanceBoundUser"
//  FAILEDOPERATION_INVALIDINSTANCEREGION = "FailedOperation.InvalidInstanceRegion"
//  FAILEDOPERATION_INVALIDINSTANCESTATUSOPERATION = "FailedOperation.InvalidInstanceStatusOperation"
//  FAILEDOPERATION_NOSUPPORTINSTANCEREGION = "FailedOperation.NoSupportInstanceRegion"
//  FAILEDOPERATION_UNKNOWNERROR = "FailedOperation.UnknownError"
//  FAILEDOPERATION_USERRESOURCENOTFOUND = "FailedOperation.UserResourceNotFound"
func (c *Client) DescribeInstanceLoginMsgWithContext(ctx context.Context, request *DescribeInstanceLoginMsgRequest) (response *DescribeInstanceLoginMsgResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceLoginMsgRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "DescribeInstanceLoginMsg")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceLoginMsg require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceLoginMsgResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceMonitorDataRequest() (request *DescribeInstanceMonitorDataRequest) {
    request = &DescribeInstanceMonitorDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "DescribeInstanceMonitorData")
    
    
    return
}

func NewDescribeInstanceMonitorDataResponse() (response *DescribeInstanceMonitorDataResponse) {
    response = &DescribeInstanceMonitorDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceMonitorData
// 查询实例监控数据
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeInstanceMonitorData(request *DescribeInstanceMonitorDataRequest) (response *DescribeInstanceMonitorDataResponse, err error) {
    return c.DescribeInstanceMonitorDataWithContext(context.Background(), request)
}

// DescribeInstanceMonitorData
// 查询实例监控数据
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeInstanceMonitorDataWithContext(ctx context.Context, request *DescribeInstanceMonitorDataRequest) (response *DescribeInstanceMonitorDataResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceMonitorDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "DescribeInstanceMonitorData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceMonitorData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceMonitorDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceRegionDistributionRequest() (request *DescribeInstanceRegionDistributionRequest) {
    request = &DescribeInstanceRegionDistributionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "DescribeInstanceRegionDistribution")
    
    
    return
}

func NewDescribeInstanceRegionDistributionResponse() (response *DescribeInstanceRegionDistributionResponse) {
    response = &DescribeInstanceRegionDistributionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceRegionDistribution
// 获取用户在各个区域云桌面实例个数
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeInstanceRegionDistribution(request *DescribeInstanceRegionDistributionRequest) (response *DescribeInstanceRegionDistributionResponse, err error) {
    return c.DescribeInstanceRegionDistributionWithContext(context.Background(), request)
}

// DescribeInstanceRegionDistribution
// 获取用户在各个区域云桌面实例个数
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeInstanceRegionDistributionWithContext(ctx context.Context, request *DescribeInstanceRegionDistributionRequest) (response *DescribeInstanceRegionDistributionResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceRegionDistributionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "DescribeInstanceRegionDistribution")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceRegionDistribution require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceRegionDistributionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "DescribeInstances")
    
    
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstances
// 获取云桌面实例详细信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    return c.DescribeInstancesWithContext(context.Background(), request)
}

// DescribeInstances
// 获取云桌面实例详细信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DescribeInstancesWithContext(ctx context.Context, request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "DescribeInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesBindStatusRequest() (request *DescribeInstancesBindStatusRequest) {
    request = &DescribeInstancesBindStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "DescribeInstancesBindStatus")
    
    
    return
}

func NewDescribeInstancesBindStatusResponse() (response *DescribeInstancesBindStatusResponse) {
    response = &DescribeInstancesBindStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstancesBindStatus
// 获取桌面实例绑定状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeInstancesBindStatus(request *DescribeInstancesBindStatusRequest) (response *DescribeInstancesBindStatusResponse, err error) {
    return c.DescribeInstancesBindStatusWithContext(context.Background(), request)
}

// DescribeInstancesBindStatus
// 获取桌面实例绑定状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeInstancesBindStatusWithContext(ctx context.Context, request *DescribeInstancesBindStatusRequest) (response *DescribeInstancesBindStatusResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesBindStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "DescribeInstancesBindStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstancesBindStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesBindStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesStatusRequest() (request *DescribeInstancesStatusRequest) {
    request = &DescribeInstancesStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "DescribeInstancesStatus")
    
    
    return
}

func NewDescribeInstancesStatusResponse() (response *DescribeInstancesStatusResponse) {
    response = &DescribeInstancesStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstancesStatus
// 获取桌面实例状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInstancesStatus(request *DescribeInstancesStatusRequest) (response *DescribeInstancesStatusResponse, err error) {
    return c.DescribeInstancesStatusWithContext(context.Background(), request)
}

// DescribeInstancesStatus
// 获取桌面实例状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInstancesStatusWithContext(ctx context.Context, request *DescribeInstancesStatusRequest) (response *DescribeInstancesStatusResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "DescribeInstancesStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstancesStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInvocationDetailRequest() (request *DescribeInvocationDetailRequest) {
    request = &DescribeInvocationDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "DescribeInvocationDetail")
    
    
    return
}

func NewDescribeInvocationDetailResponse() (response *DescribeInvocationDetailResponse) {
    response = &DescribeInvocationDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInvocationDetail
// 查询脚本任务执行的详细情况。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeInvocationDetail(request *DescribeInvocationDetailRequest) (response *DescribeInvocationDetailResponse, err error) {
    return c.DescribeInvocationDetailWithContext(context.Background(), request)
}

// DescribeInvocationDetail
// 查询脚本任务执行的详细情况。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeInvocationDetailWithContext(ctx context.Context, request *DescribeInvocationDetailRequest) (response *DescribeInvocationDetailResponse, err error) {
    if request == nil {
        request = NewDescribeInvocationDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "DescribeInvocationDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInvocationDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInvocationDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInvocationsRequest() (request *DescribeInvocationsRequest) {
    request = &DescribeInvocationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "DescribeInvocations")
    
    
    return
}

func NewDescribeInvocationsResponse() (response *DescribeInvocationsResponse) {
    response = &DescribeInvocationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInvocations
// 查询指定云桌面上的任务列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeInvocations(request *DescribeInvocationsRequest) (response *DescribeInvocationsResponse, err error) {
    return c.DescribeInvocationsWithContext(context.Background(), request)
}

// DescribeInvocations
// 查询指定云桌面上的任务列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeInvocationsWithContext(ctx context.Context, request *DescribeInvocationsRequest) (response *DescribeInvocationsResponse, err error) {
    if request == nil {
        request = NewDescribeInvocationsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "DescribeInvocations")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInvocations require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInvocationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePolicyDetailRequest() (request *DescribePolicyDetailRequest) {
    request = &DescribePolicyDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "DescribePolicyDetail")
    
    
    return
}

func NewDescribePolicyDetailResponse() (response *DescribePolicyDetailResponse) {
    response = &DescribePolicyDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePolicyDetail
// 拉取租户策略列表 
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribePolicyDetail(request *DescribePolicyDetailRequest) (response *DescribePolicyDetailResponse, err error) {
    return c.DescribePolicyDetailWithContext(context.Background(), request)
}

// DescribePolicyDetail
// 拉取租户策略列表 
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribePolicyDetailWithContext(ctx context.Context, request *DescribePolicyDetailRequest) (response *DescribePolicyDetailResponse, err error) {
    if request == nil {
        request = NewDescribePolicyDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "DescribePolicyDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePolicyDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePolicyDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityGroupsRequest() (request *DescribeSecurityGroupsRequest) {
    request = &DescribeSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "DescribeSecurityGroups")
    
    
    return
}

func NewDescribeSecurityGroupsResponse() (response *DescribeSecurityGroupsResponse) {
    response = &DescribeSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSecurityGroups
// 拉取云桌面关联的安全组信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeSecurityGroups(request *DescribeSecurityGroupsRequest) (response *DescribeSecurityGroupsResponse, err error) {
    return c.DescribeSecurityGroupsWithContext(context.Background(), request)
}

// DescribeSecurityGroups
// 拉取云桌面关联的安全组信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeSecurityGroupsWithContext(ctx context.Context, request *DescribeSecurityGroupsRequest) (response *DescribeSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "DescribeSecurityGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSnapshotsRequest() (request *DescribeSnapshotsRequest) {
    request = &DescribeSnapshotsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "DescribeSnapshots")
    
    
    return
}

func NewDescribeSnapshotsResponse() (response *DescribeSnapshotsResponse) {
    response = &DescribeSnapshotsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSnapshots
// 拉取快照列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DescribeSnapshots(request *DescribeSnapshotsRequest) (response *DescribeSnapshotsResponse, err error) {
    return c.DescribeSnapshotsWithContext(context.Background(), request)
}

// DescribeSnapshots
// 拉取快照列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DescribeSnapshotsWithContext(ctx context.Context, request *DescribeSnapshotsRequest) (response *DescribeSnapshotsResponse, err error) {
    if request == nil {
        request = NewDescribeSnapshotsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "DescribeSnapshots")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSnapshots require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSnapshotsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsageRequest() (request *DescribeUsageRequest) {
    request = &DescribeUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "DescribeUsage")
    
    
    return
}

func NewDescribeUsageResponse() (response *DescribeUsageResponse) {
    response = &DescribeUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUsage
// 拉取实例使用时长详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_USERKEYWORD = "InvalidParameter.UserKeyword"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeUsage(request *DescribeUsageRequest) (response *DescribeUsageResponse, err error) {
    return c.DescribeUsageWithContext(context.Background(), request)
}

// DescribeUsage
// 拉取实例使用时长详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_USERKEYWORD = "InvalidParameter.UserKeyword"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeUsageWithContext(ctx context.Context, request *DescribeUsageRequest) (response *DescribeUsageResponse, err error) {
    if request == nil {
        request = NewDescribeUsageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "DescribeUsage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUsage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUsageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsersRequest() (request *DescribeUsersRequest) {
    request = &DescribeUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "DescribeUsers")
    
    
    return
}

func NewDescribeUsersResponse() (response *DescribeUsersResponse) {
    response = &DescribeUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUsers
// 拉取用户列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  UNAUTHORIZEDOPERATION_SERVICENOTACTIVATED = "UnauthorizedOperation.ServiceNotActivated"
func (c *Client) DescribeUsers(request *DescribeUsersRequest) (response *DescribeUsersResponse, err error) {
    return c.DescribeUsersWithContext(context.Background(), request)
}

// DescribeUsers
// 拉取用户列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  UNAUTHORIZEDOPERATION_SERVICENOTACTIVATED = "UnauthorizedOperation.ServiceNotActivated"
func (c *Client) DescribeUsersWithContext(ctx context.Context, request *DescribeUsersRequest) (response *DescribeUsersResponse, err error) {
    if request == nil {
        request = NewDescribeUsersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "DescribeUsers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUsers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUsersResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateSecurityGroupRequest() (request *DisassociateSecurityGroupRequest) {
    request = &DisassociateSecurityGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "DisassociateSecurityGroup")
    
    
    return
}

func NewDisassociateSecurityGroupResponse() (response *DisassociateSecurityGroupResponse) {
    response = &DisassociateSecurityGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisassociateSecurityGroup
// 安全组批量解关联云桌面
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DisassociateSecurityGroup(request *DisassociateSecurityGroupRequest) (response *DisassociateSecurityGroupResponse, err error) {
    return c.DisassociateSecurityGroupWithContext(context.Background(), request)
}

// DisassociateSecurityGroup
// 安全组批量解关联云桌面
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DisassociateSecurityGroupWithContext(ctx context.Context, request *DisassociateSecurityGroupRequest) (response *DisassociateSecurityGroupResponse, err error) {
    if request == nil {
        request = NewDisassociateSecurityGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "DisassociateSecurityGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisassociateSecurityGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisassociateSecurityGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDiskConfigurationRequest() (request *ModifyDiskConfigurationRequest) {
    request = &ModifyDiskConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "ModifyDiskConfiguration")
    
    
    return
}

func NewModifyDiskConfigurationResponse() (response *ModifyDiskConfigurationResponse) {
    response = &ModifyDiskConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDiskConfiguration
// 实例硬盘扩容
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyDiskConfiguration(request *ModifyDiskConfigurationRequest) (response *ModifyDiskConfigurationResponse, err error) {
    return c.ModifyDiskConfigurationWithContext(context.Background(), request)
}

// ModifyDiskConfiguration
// 实例硬盘扩容
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyDiskConfigurationWithContext(ctx context.Context, request *ModifyDiskConfigurationRequest) (response *ModifyDiskConfigurationResponse, err error) {
    if request == nil {
        request = NewModifyDiskConfigurationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "ModifyDiskConfiguration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDiskConfiguration require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDiskConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyImageAttrsRequest() (request *ModifyImageAttrsRequest) {
    request = &ModifyImageAttrsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "ModifyImageAttrs")
    
    
    return
}

func NewModifyImageAttrsResponse() (response *ModifyImageAttrsResponse) {
    response = &ModifyImageAttrsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyImageAttrs
// 修改镜像属性
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyImageAttrs(request *ModifyImageAttrsRequest) (response *ModifyImageAttrsResponse, err error) {
    return c.ModifyImageAttrsWithContext(context.Background(), request)
}

// ModifyImageAttrs
// 修改镜像属性
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyImageAttrsWithContext(ctx context.Context, request *ModifyImageAttrsRequest) (response *ModifyImageAttrsResponse, err error) {
    if request == nil {
        request = NewModifyImageAttrsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "ModifyImageAttrs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyImageAttrs require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyImageAttrsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceBasicAttributeRequest() (request *ModifyInstanceBasicAttributeRequest) {
    request = &ModifyInstanceBasicAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "ModifyInstanceBasicAttribute")
    
    
    return
}

func NewModifyInstanceBasicAttributeResponse() (response *ModifyInstanceBasicAttributeResponse) {
    response = &ModifyInstanceBasicAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstanceBasicAttribute
// 修改桌面实例属性
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyInstanceBasicAttribute(request *ModifyInstanceBasicAttributeRequest) (response *ModifyInstanceBasicAttributeResponse, err error) {
    return c.ModifyInstanceBasicAttributeWithContext(context.Background(), request)
}

// ModifyInstanceBasicAttribute
// 修改桌面实例属性
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyInstanceBasicAttributeWithContext(ctx context.Context, request *ModifyInstanceBasicAttributeRequest) (response *ModifyInstanceBasicAttributeResponse, err error) {
    if request == nil {
        request = NewModifyInstanceBasicAttributeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "ModifyInstanceBasicAttribute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceBasicAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceBasicAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstancesTypeRequest() (request *ModifyInstancesTypeRequest) {
    request = &ModifyInstancesTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "ModifyInstancesType")
    
    
    return
}

func NewModifyInstancesTypeResponse() (response *ModifyInstancesTypeResponse) {
    response = &ModifyInstancesTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstancesType
// 调整实例机型配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DEALCHECKFAILED = "FailedOperation.DealCheckFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_RESOURCENOTENOUGH = "InvalidParameterValue.ResourceNotEnough"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) ModifyInstancesType(request *ModifyInstancesTypeRequest) (response *ModifyInstancesTypeResponse, err error) {
    return c.ModifyInstancesTypeWithContext(context.Background(), request)
}

// ModifyInstancesType
// 调整实例机型配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DEALCHECKFAILED = "FailedOperation.DealCheckFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_RESOURCENOTENOUGH = "InvalidParameterValue.ResourceNotEnough"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) ModifyInstancesTypeWithContext(ctx context.Context, request *ModifyInstancesTypeRequest) (response *ModifyInstancesTypeResponse, err error) {
    if request == nil {
        request = NewModifyInstancesTypeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "ModifyInstancesType")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstancesType require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstancesTypeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstancesVpcRequest() (request *ModifyInstancesVpcRequest) {
    request = &ModifyInstancesVpcRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "ModifyInstancesVpc")
    
    
    return
}

func NewModifyInstancesVpcResponse() (response *ModifyInstancesVpcResponse) {
    response = &ModifyInstancesVpcResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstancesVpc
// 调整实例私有网络
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyInstancesVpc(request *ModifyInstancesVpcRequest) (response *ModifyInstancesVpcResponse, err error) {
    return c.ModifyInstancesVpcWithContext(context.Background(), request)
}

// ModifyInstancesVpc
// 调整实例私有网络
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyInstancesVpcWithContext(ctx context.Context, request *ModifyInstancesVpcRequest) (response *ModifyInstancesVpcResponse, err error) {
    if request == nil {
        request = NewModifyInstancesVpcRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "ModifyInstancesVpc")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstancesVpc require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstancesVpcResponse()
    err = c.Send(request, response)
    return
}

func NewRebootInstancesRequest() (request *RebootInstancesRequest) {
    request = &RebootInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "RebootInstances")
    
    
    return
}

func NewRebootInstancesResponse() (response *RebootInstancesResponse) {
    response = &RebootInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RebootInstances
// 重启实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) RebootInstances(request *RebootInstancesRequest) (response *RebootInstancesResponse, err error) {
    return c.RebootInstancesWithContext(context.Background(), request)
}

// RebootInstances
// 重启实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) RebootInstancesWithContext(ctx context.Context, request *RebootInstancesRequest) (response *RebootInstancesResponse, err error) {
    if request == nil {
        request = NewRebootInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "RebootInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RebootInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewRebootInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveDataDisksRequest() (request *RemoveDataDisksRequest) {
    request = &RemoveDataDisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "RemoveDataDisks")
    
    
    return
}

func NewRemoveDataDisksResponse() (response *RemoveDataDisksResponse) {
    response = &RemoveDataDisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemoveDataDisks
// 卸载实例数据盘
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) RemoveDataDisks(request *RemoveDataDisksRequest) (response *RemoveDataDisksResponse, err error) {
    return c.RemoveDataDisksWithContext(context.Background(), request)
}

// RemoveDataDisks
// 卸载实例数据盘
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) RemoveDataDisksWithContext(ctx context.Context, request *RemoveDataDisksRequest) (response *RemoveDataDisksResponse, err error) {
    if request == nil {
        request = NewRemoveDataDisksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "RemoveDataDisks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveDataDisks require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveDataDisksResponse()
    err = c.Send(request, response)
    return
}

func NewRemovePolicyRequest() (request *RemovePolicyRequest) {
    request = &RemovePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "RemovePolicy")
    
    
    return
}

func NewRemovePolicyResponse() (response *RemovePolicyResponse) {
    response = &RemovePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemovePolicy
// 删除策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) RemovePolicy(request *RemovePolicyRequest) (response *RemovePolicyResponse, err error) {
    return c.RemovePolicyWithContext(context.Background(), request)
}

// RemovePolicy
// 删除策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) RemovePolicyWithContext(ctx context.Context, request *RemovePolicyRequest) (response *RemovePolicyResponse, err error) {
    if request == nil {
        request = NewRemovePolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "RemovePolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemovePolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemovePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewRenewInstancesRequest() (request *RenewInstancesRequest) {
    request = &RenewInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "RenewInstances")
    
    
    return
}

func NewRenewInstancesResponse() (response *RenewInstancesResponse) {
    response = &RenewInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RenewInstances
// 实例续费
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEALCHECKFAILED = "FailedOperation.DealCheckFailed"
//  FAILEDOPERATION_INSUFFICIENTBALANCE = "FailedOperation.InsufficientBalance"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  UNAUTHORIZEDOPERATION_SERVICENOTACTIVATED = "UnauthorizedOperation.ServiceNotActivated"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RenewInstances(request *RenewInstancesRequest) (response *RenewInstancesResponse, err error) {
    return c.RenewInstancesWithContext(context.Background(), request)
}

// RenewInstances
// 实例续费
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEALCHECKFAILED = "FailedOperation.DealCheckFailed"
//  FAILEDOPERATION_INSUFFICIENTBALANCE = "FailedOperation.InsufficientBalance"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  UNAUTHORIZEDOPERATION_SERVICENOTACTIVATED = "UnauthorizedOperation.ServiceNotActivated"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RenewInstancesWithContext(ctx context.Context, request *RenewInstancesRequest) (response *RenewInstancesResponse, err error) {
    if request == nil {
        request = NewRenewInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "RenewInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RenewInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewRenewInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewResetInstancesRequest() (request *ResetInstancesRequest) {
    request = &ResetInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "ResetInstances")
    
    
    return
}

func NewResetInstancesResponse() (response *ResetInstancesResponse) {
    response = &ResetInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResetInstances
// 实例重置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDIMAGESIZE = "FailedOperation.InvalidImageSize"
//  FAILEDOPERATION_INVALIDSTATUS = "FailedOperation.InvalidStatus"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ResetInstances(request *ResetInstancesRequest) (response *ResetInstancesResponse, err error) {
    return c.ResetInstancesWithContext(context.Background(), request)
}

// ResetInstances
// 实例重置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDIMAGESIZE = "FailedOperation.InvalidImageSize"
//  FAILEDOPERATION_INVALIDSTATUS = "FailedOperation.InvalidStatus"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ResetInstancesWithContext(ctx context.Context, request *ResetInstancesRequest) (response *ResetInstancesResponse, err error) {
    if request == nil {
        request = NewResetInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "ResetInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewRunCommandRequest() (request *RunCommandRequest) {
    request = &RunCommandRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "RunCommand")
    
    
    return
}

func NewRunCommandResponse() (response *RunCommandResponse) {
    response = &RunCommandResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RunCommand
// 在指定的云桌面实例上执行脚本任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) RunCommand(request *RunCommandRequest) (response *RunCommandResponse, err error) {
    return c.RunCommandWithContext(context.Background(), request)
}

// RunCommand
// 在指定的云桌面实例上执行脚本任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) RunCommandWithContext(ctx context.Context, request *RunCommandRequest) (response *RunCommandResponse, err error) {
    if request == nil {
        request = NewRunCommandRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "RunCommand")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunCommand require credential")
    }

    request.SetContext(ctx)
    
    response = NewRunCommandResponse()
    err = c.Send(request, response)
    return
}

func NewStartInstancesRequest() (request *StartInstancesRequest) {
    request = &StartInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "StartInstances")
    
    
    return
}

func NewStartInstancesResponse() (response *StartInstancesResponse) {
    response = &StartInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartInstances
// 开启实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCEINSUFFICIENT = "FailedOperation.ResourceInsufficient"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) StartInstances(request *StartInstancesRequest) (response *StartInstancesResponse, err error) {
    return c.StartInstancesWithContext(context.Background(), request)
}

// StartInstances
// 开启实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCEINSUFFICIENT = "FailedOperation.ResourceInsufficient"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) StartInstancesWithContext(ctx context.Context, request *StartInstancesRequest) (response *StartInstancesResponse, err error) {
    if request == nil {
        request = NewStartInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "StartInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewStopCommandTasksRequest() (request *StopCommandTasksRequest) {
    request = &StopCommandTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "StopCommandTasks")
    
    
    return
}

func NewStopCommandTasksResponse() (response *StopCommandTasksResponse) {
    response = &StopCommandTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopCommandTasks
// 停止命令任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER_COMMANDNOTRUNNING = "InvalidParameter.CommandNotRunning"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) StopCommandTasks(request *StopCommandTasksRequest) (response *StopCommandTasksResponse, err error) {
    return c.StopCommandTasksWithContext(context.Background(), request)
}

// StopCommandTasks
// 停止命令任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER_COMMANDNOTRUNNING = "InvalidParameter.CommandNotRunning"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) StopCommandTasksWithContext(ctx context.Context, request *StopCommandTasksRequest) (response *StopCommandTasksResponse, err error) {
    if request == nil {
        request = NewStopCommandTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "StopCommandTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopCommandTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopCommandTasksResponse()
    err = c.Send(request, response)
    return
}

func NewStopInstancesRequest() (request *StopInstancesRequest) {
    request = &StopInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "StopInstances")
    
    
    return
}

func NewStopInstancesResponse() (response *StopInstancesResponse) {
    response = &StopInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopInstances
// 关闭实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) StopInstances(request *StopInstancesRequest) (response *StopInstancesResponse, err error) {
    return c.StopInstancesWithContext(context.Background(), request)
}

// StopInstances
// 关闭实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) StopInstancesWithContext(ctx context.Context, request *StopInstancesRequest) (response *StopInstancesResponse, err error) {
    if request == nil {
        request = NewStopInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "StopInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewStopInvocationRequest() (request *StopInvocationRequest) {
    request = &StopInvocationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "StopInvocation")
    
    
    return
}

func NewStopInvocationResponse() (response *StopInvocationResponse) {
    response = &StopInvocationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopInvocation
// 停止脚本任务的执行。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) StopInvocation(request *StopInvocationRequest) (response *StopInvocationResponse, err error) {
    return c.StopInvocationWithContext(context.Background(), request)
}

// StopInvocation
// 停止脚本任务的执行。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) StopInvocationWithContext(ctx context.Context, request *StopInvocationRequest) (response *StopInvocationResponse, err error) {
    if request == nil {
        request = NewStopInvocationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "StopInvocation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopInvocation require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopInvocationResponse()
    err = c.Send(request, response)
    return
}

func NewStopSessionRequest() (request *StopSessionRequest) {
    request = &StopSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "StopSession")
    
    
    return
}

func NewStopSessionResponse() (response *StopSessionResponse) {
    response = &StopSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopSession
// 注销或者断开机器当前会话
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) StopSession(request *StopSessionRequest) (response *StopSessionResponse, err error) {
    return c.StopSessionWithContext(context.Background(), request)
}

// StopSession
// 注销或者断开机器当前会话
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) StopSessionWithContext(ctx context.Context, request *StopSessionRequest) (response *StopSessionResponse, err error) {
    if request == nil {
        request = NewStopSessionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "StopSession")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopSession require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopSessionResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateInstancesRequest() (request *TerminateInstancesRequest) {
    request = &TerminateInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "TerminateInstances")
    
    
    return
}

func NewTerminateInstancesResponse() (response *TerminateInstancesResponse) {
    response = &TerminateInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TerminateInstances
// 主动退还、销毁退还实例。 
//
// <li>包年包月实例首次调用本接口，且参数DestroyImmediately为false时，实例进入隔离期，再次调用本接口，实例将被销毁，且不可恢复。</li>
//
// <li>包年包月实例，如果参数DestroyImmediately为true来调用此接口，则实例将被直接销毁，且不可恢复。</li>
//
// <li>按量计费实例调用本接口将被直接销毁，且不可恢复。</li>
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTACTIVATED = "UnauthorizedOperation.ServiceNotActivated"
//  UNSUPPORTEDOPERATION_INSTANCESTATE = "UnsupportedOperation.InstanceState"
func (c *Client) TerminateInstances(request *TerminateInstancesRequest) (response *TerminateInstancesResponse, err error) {
    return c.TerminateInstancesWithContext(context.Background(), request)
}

// TerminateInstances
// 主动退还、销毁退还实例。 
//
// <li>包年包月实例首次调用本接口，且参数DestroyImmediately为false时，实例进入隔离期，再次调用本接口，实例将被销毁，且不可恢复。</li>
//
// <li>包年包月实例，如果参数DestroyImmediately为true来调用此接口，则实例将被直接销毁，且不可恢复。</li>
//
// <li>按量计费实例调用本接口将被直接销毁，且不可恢复。</li>
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTACTIVATED = "UnauthorizedOperation.ServiceNotActivated"
//  UNSUPPORTEDOPERATION_INSTANCESTATE = "UnsupportedOperation.InstanceState"
func (c *Client) TerminateInstancesWithContext(ctx context.Context, request *TerminateInstancesRequest) (response *TerminateInstancesResponse, err error) {
    if request == nil {
        request = NewTerminateInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "TerminateInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminateInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewTerminateInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindInstancesRequest() (request *UnbindInstancesRequest) {
    request = &UnbindInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "UnbindInstances")
    
    
    return
}

func NewUnbindInstancesResponse() (response *UnbindInstancesResponse) {
    response = &UnbindInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UnbindInstances
// 解除用户和机器绑定关系
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) UnbindInstances(request *UnbindInstancesRequest) (response *UnbindInstancesResponse, err error) {
    return c.UnbindInstancesWithContext(context.Background(), request)
}

// UnbindInstances
// 解除用户和机器绑定关系
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) UnbindInstancesWithContext(ctx context.Context, request *UnbindInstancesRequest) (response *UnbindInstancesResponse, err error) {
    if request == nil {
        request = NewUnbindInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "UnbindInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewUpdatePolicyRequest() (request *UpdatePolicyRequest) {
    request = &UpdatePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cvd", APIVersion, "UpdatePolicy")
    
    
    return
}

func NewUpdatePolicyResponse() (response *UpdatePolicyResponse) {
    response = &UpdatePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdatePolicy
// 更新策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_POLICYEXISTS = "ResourceInUse.PolicyExists"
func (c *Client) UpdatePolicy(request *UpdatePolicyRequest) (response *UpdatePolicyResponse, err error) {
    return c.UpdatePolicyWithContext(context.Background(), request)
}

// UpdatePolicy
// 更新策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_POLICYEXISTS = "ResourceInUse.PolicyExists"
func (c *Client) UpdatePolicyWithContext(ctx context.Context, request *UpdatePolicyRequest) (response *UpdatePolicyResponse, err error) {
    if request == nil {
        request = NewUpdatePolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cvd", APIVersion, "UpdatePolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdatePolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdatePolicyResponse()
    err = c.Send(request, response)
    return
}
