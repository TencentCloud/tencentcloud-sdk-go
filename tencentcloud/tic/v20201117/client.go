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

package v20201117

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-11-17"

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


func NewApplyStackRequest() (request *ApplyStackRequest) {
    request = &ApplyStackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tic", APIVersion, "ApplyStack")
    
    
    return
}

func NewApplyStackResponse() (response *ApplyStackResponse) {
    response = &ApplyStackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ApplyStack
// 本接口（ApplyStack）用于触发资源栈下某个版本的Apply事件。
//
// 
//
// - 当版本处于PLAN_IN_PROGRESS或APPLY_IN_PROGRESS状态时，将无法再执行本操作
//
// - 当版本处于APPLY_COMPLETED状态时，本操作无法执行
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTEXIST = "FailedOperation.NotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  UNSUPPORTEDOPERATION_FORBIDOP = "UnsupportedOperation.ForbidOp"
func (c *Client) ApplyStack(request *ApplyStackRequest) (response *ApplyStackResponse, err error) {
    return c.ApplyStackWithContext(context.Background(), request)
}

// ApplyStack
// 本接口（ApplyStack）用于触发资源栈下某个版本的Apply事件。
//
// 
//
// - 当版本处于PLAN_IN_PROGRESS或APPLY_IN_PROGRESS状态时，将无法再执行本操作
//
// - 当版本处于APPLY_COMPLETED状态时，本操作无法执行
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTEXIST = "FailedOperation.NotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  UNSUPPORTEDOPERATION_FORBIDOP = "UnsupportedOperation.ForbidOp"
func (c *Client) ApplyStackWithContext(ctx context.Context, request *ApplyStackRequest) (response *ApplyStackResponse, err error) {
    if request == nil {
        request = NewApplyStackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyStack require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyStackResponse()
    err = c.Send(request, response)
    return
}

func NewCreateStackRequest() (request *CreateStackRequest) {
    request = &CreateStackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tic", APIVersion, "CreateStack")
    
    
    return
}

func NewCreateStackResponse() (response *CreateStackResponse) {
    response = &CreateStackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateStack
// 本接口（CreateStack）用于通过传递一个COS的terraform zip模版URL来创建一个资源栈。创建资源栈后仍需要用户调用对应的plan, apply, destory执行对应的事件。
//
// 可能返回的错误码:
//  FAILEDOPERATION_HTTPREQUESTERROR = "FailedOperation.HttpRequestError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) CreateStack(request *CreateStackRequest) (response *CreateStackResponse, err error) {
    return c.CreateStackWithContext(context.Background(), request)
}

// CreateStack
// 本接口（CreateStack）用于通过传递一个COS的terraform zip模版URL来创建一个资源栈。创建资源栈后仍需要用户调用对应的plan, apply, destory执行对应的事件。
//
// 可能返回的错误码:
//  FAILEDOPERATION_HTTPREQUESTERROR = "FailedOperation.HttpRequestError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) CreateStackWithContext(ctx context.Context, request *CreateStackRequest) (response *CreateStackResponse, err error) {
    if request == nil {
        request = NewCreateStackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateStack require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateStackResponse()
    err = c.Send(request, response)
    return
}

func NewCreateStackVersionRequest() (request *CreateStackVersionRequest) {
    request = &CreateStackVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tic", APIVersion, "CreateStackVersion")
    
    
    return
}

func NewCreateStackVersionResponse() (response *CreateStackVersionResponse) {
    response = &CreateStackVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateStackVersion
// 本接口（CreateStackVersion）用于给资源栈新增一个HCL模版版本，仅限COS链接，且为zip格式。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTEXIST = "FailedOperation.NotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_MQERROR = "InternalError.MqError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  UNSUPPORTEDOPERATION_FORBIDOP = "UnsupportedOperation.ForbidOp"
func (c *Client) CreateStackVersion(request *CreateStackVersionRequest) (response *CreateStackVersionResponse, err error) {
    return c.CreateStackVersionWithContext(context.Background(), request)
}

// CreateStackVersion
// 本接口（CreateStackVersion）用于给资源栈新增一个HCL模版版本，仅限COS链接，且为zip格式。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTEXIST = "FailedOperation.NotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_MQERROR = "InternalError.MqError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  UNSUPPORTEDOPERATION_FORBIDOP = "UnsupportedOperation.ForbidOp"
func (c *Client) CreateStackVersionWithContext(ctx context.Context, request *CreateStackVersionRequest) (response *CreateStackVersionResponse, err error) {
    if request == nil {
        request = NewCreateStackVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateStackVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateStackVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteStackRequest() (request *DeleteStackRequest) {
    request = &DeleteStackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tic", APIVersion, "DeleteStack")
    
    
    return
}

func NewDeleteStackResponse() (response *DeleteStackResponse) {
    response = &DeleteStackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteStack
// 本接口（DeleteStack）用于删除一个资源栈（配置、版本、事件信息）。但不会销毁资源管理的云资源。如果需要销毁资源栈管理的云资源，请调用 DestoryStack 接口销毁云资源。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTEXIST = "FailedOperation.NotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteStack(request *DeleteStackRequest) (response *DeleteStackResponse, err error) {
    return c.DeleteStackWithContext(context.Background(), request)
}

// DeleteStack
// 本接口（DeleteStack）用于删除一个资源栈（配置、版本、事件信息）。但不会销毁资源管理的云资源。如果需要销毁资源栈管理的云资源，请调用 DestoryStack 接口销毁云资源。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTEXIST = "FailedOperation.NotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteStackWithContext(ctx context.Context, request *DeleteStackRequest) (response *DeleteStackResponse, err error) {
    if request == nil {
        request = NewDeleteStackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteStack require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteStackResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteStackVersionRequest() (request *DeleteStackVersionRequest) {
    request = &DeleteStackVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tic", APIVersion, "DeleteStackVersion")
    
    
    return
}

func NewDeleteStackVersionResponse() (response *DeleteStackVersionResponse) {
    response = &DeleteStackVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteStackVersion
// 本接口（DeleteStackVersion）用于删除一个版本，处于PLAN_IN_PROGRESS和APPLY_IN_PROGRESS状态中的版本无法删除。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) DeleteStackVersion(request *DeleteStackVersionRequest) (response *DeleteStackVersionResponse, err error) {
    return c.DeleteStackVersionWithContext(context.Background(), request)
}

// DeleteStackVersion
// 本接口（DeleteStackVersion）用于删除一个版本，处于PLAN_IN_PROGRESS和APPLY_IN_PROGRESS状态中的版本无法删除。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) DeleteStackVersionWithContext(ctx context.Context, request *DeleteStackVersionRequest) (response *DeleteStackVersionResponse, err error) {
    if request == nil {
        request = NewDeleteStackVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteStackVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteStackVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStackEventRequest() (request *DescribeStackEventRequest) {
    request = &DescribeStackEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tic", APIVersion, "DescribeStackEvent")
    
    
    return
}

func NewDescribeStackEventResponse() (response *DescribeStackEventResponse) {
    response = &DescribeStackEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStackEvent
// 本接口（DescribeStackEvent）用于获取单个事件详情，尤其是可以得到事件的详细控制台输出文本。
//
// 可能返回的错误码:
//  FAILEDOPERATION_HTTPREQUESTERROR = "FailedOperation.HttpRequestError"
//  FAILEDOPERATION_NOTEXIST = "FailedOperation.NotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) DescribeStackEvent(request *DescribeStackEventRequest) (response *DescribeStackEventResponse, err error) {
    return c.DescribeStackEventWithContext(context.Background(), request)
}

// DescribeStackEvent
// 本接口（DescribeStackEvent）用于获取单个事件详情，尤其是可以得到事件的详细控制台输出文本。
//
// 可能返回的错误码:
//  FAILEDOPERATION_HTTPREQUESTERROR = "FailedOperation.HttpRequestError"
//  FAILEDOPERATION_NOTEXIST = "FailedOperation.NotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) DescribeStackEventWithContext(ctx context.Context, request *DescribeStackEventRequest) (response *DescribeStackEventResponse, err error) {
    if request == nil {
        request = NewDescribeStackEventRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStackEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStackEventResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStackEventsRequest() (request *DescribeStackEventsRequest) {
    request = &DescribeStackEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tic", APIVersion, "DescribeStackEvents")
    
    
    return
}

func NewDescribeStackEventsResponse() (response *DescribeStackEventsResponse) {
    response = &DescribeStackEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStackEvents
// 本接口（DescribeStackEvents）用于查看一个或多个事件详细信息。
//
// 
//
// - 可以根据事件ID过滤感兴趣的事件
//
// - 也可以根据版本ID，资源栈ID，事件类型，事件状态过滤事件，过滤信息详细请见过滤器Filter
//
// - 如果参数为空，返回当前用户一定数量（Limit所指定的数量，默认为20）的事件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeStackEvents(request *DescribeStackEventsRequest) (response *DescribeStackEventsResponse, err error) {
    return c.DescribeStackEventsWithContext(context.Background(), request)
}

// DescribeStackEvents
// 本接口（DescribeStackEvents）用于查看一个或多个事件详细信息。
//
// 
//
// - 可以根据事件ID过滤感兴趣的事件
//
// - 也可以根据版本ID，资源栈ID，事件类型，事件状态过滤事件，过滤信息详细请见过滤器Filter
//
// - 如果参数为空，返回当前用户一定数量（Limit所指定的数量，默认为20）的事件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeStackEventsWithContext(ctx context.Context, request *DescribeStackEventsRequest) (response *DescribeStackEventsResponse, err error) {
    if request == nil {
        request = NewDescribeStackEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStackEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStackEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStackVersionsRequest() (request *DescribeStackVersionsRequest) {
    request = &DescribeStackVersionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tic", APIVersion, "DescribeStackVersions")
    
    
    return
}

func NewDescribeStackVersionsResponse() (response *DescribeStackVersionsResponse) {
    response = &DescribeStackVersionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStackVersions
// 本接口（DescribeStackVersions）用于查询一个或多个版本的详细信息。
//
// 
//
// - 可以根据版本ID查询感兴趣的版本
//
// - 可以根据版本名字和状态来过滤版本，详见过滤器Filter
//
// - 如果参数为空，返回当前用户一定数量（Limit所指定的数量，默认为20）的版本
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTEXIST = "FailedOperation.NotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeStackVersions(request *DescribeStackVersionsRequest) (response *DescribeStackVersionsResponse, err error) {
    return c.DescribeStackVersionsWithContext(context.Background(), request)
}

// DescribeStackVersions
// 本接口（DescribeStackVersions）用于查询一个或多个版本的详细信息。
//
// 
//
// - 可以根据版本ID查询感兴趣的版本
//
// - 可以根据版本名字和状态来过滤版本，详见过滤器Filter
//
// - 如果参数为空，返回当前用户一定数量（Limit所指定的数量，默认为20）的版本
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTEXIST = "FailedOperation.NotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeStackVersionsWithContext(ctx context.Context, request *DescribeStackVersionsRequest) (response *DescribeStackVersionsResponse, err error) {
    if request == nil {
        request = NewDescribeStackVersionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStackVersions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStackVersionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStacksRequest() (request *DescribeStacksRequest) {
    request = &DescribeStacksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tic", APIVersion, "DescribeStacks")
    
    
    return
}

func NewDescribeStacksResponse() (response *DescribeStacksResponse) {
    response = &DescribeStacksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStacks
// 本接口（DescribeStacks）用于查询一个或多个资源栈的详细信息。
//
// 
//
// - 可以根据资源栈ID来查询感兴趣的资源栈信息
//
// - 若参数为空，返回当前用户一定数量（Limit所指定的数量，默认为20）的资源栈
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeStacks(request *DescribeStacksRequest) (response *DescribeStacksResponse, err error) {
    return c.DescribeStacksWithContext(context.Background(), request)
}

// DescribeStacks
// 本接口（DescribeStacks）用于查询一个或多个资源栈的详细信息。
//
// 
//
// - 可以根据资源栈ID来查询感兴趣的资源栈信息
//
// - 若参数为空，返回当前用户一定数量（Limit所指定的数量，默认为20）的资源栈
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeStacksWithContext(ctx context.Context, request *DescribeStacksRequest) (response *DescribeStacksResponse, err error) {
    if request == nil {
        request = NewDescribeStacksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStacks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStacksResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyStackRequest() (request *DestroyStackRequest) {
    request = &DestroyStackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tic", APIVersion, "DestroyStack")
    
    
    return
}

func NewDestroyStackResponse() (response *DestroyStackResponse) {
    response = &DestroyStackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DestroyStack
// 本接口（DestroyStack）用于删除资源栈下的某个版本所创建的资源。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTEXIST = "FailedOperation.NotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) DestroyStack(request *DestroyStackRequest) (response *DestroyStackResponse, err error) {
    return c.DestroyStackWithContext(context.Background(), request)
}

// DestroyStack
// 本接口（DestroyStack）用于删除资源栈下的某个版本所创建的资源。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTEXIST = "FailedOperation.NotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) DestroyStackWithContext(ctx context.Context, request *DestroyStackRequest) (response *DestroyStackResponse, err error) {
    if request == nil {
        request = NewDestroyStackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DestroyStack require credential")
    }

    request.SetContext(ctx)
    
    response = NewDestroyStackResponse()
    err = c.Send(request, response)
    return
}

func NewPlanStackRequest() (request *PlanStackRequest) {
    request = &PlanStackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tic", APIVersion, "PlanStack")
    
    
    return
}

func NewPlanStackResponse() (response *PlanStackResponse) {
    response = &PlanStackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PlanStack
// 本接口（PlanStack）用于触发资源栈下某个版本的PLAN事件。
//
// 
//
// - 当版本处于PLAN_IN_PROGRESS或APPLY_IN_PROGRESS状态时，将无法再执行本操作
//
// - 当版本处于APPLY_COMPLETED状态时，本操作无法执行
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTEXIST = "FailedOperation.NotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  UNSUPPORTEDOPERATION_FORBIDOP = "UnsupportedOperation.ForbidOp"
func (c *Client) PlanStack(request *PlanStackRequest) (response *PlanStackResponse, err error) {
    return c.PlanStackWithContext(context.Background(), request)
}

// PlanStack
// 本接口（PlanStack）用于触发资源栈下某个版本的PLAN事件。
//
// 
//
// - 当版本处于PLAN_IN_PROGRESS或APPLY_IN_PROGRESS状态时，将无法再执行本操作
//
// - 当版本处于APPLY_COMPLETED状态时，本操作无法执行
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTEXIST = "FailedOperation.NotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  UNSUPPORTEDOPERATION_FORBIDOP = "UnsupportedOperation.ForbidOp"
func (c *Client) PlanStackWithContext(ctx context.Context, request *PlanStackRequest) (response *PlanStackResponse, err error) {
    if request == nil {
        request = NewPlanStackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PlanStack require credential")
    }

    request.SetContext(ctx)
    
    response = NewPlanStackResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateStackRequest() (request *UpdateStackRequest) {
    request = &UpdateStackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tic", APIVersion, "UpdateStack")
    
    
    return
}

func NewUpdateStackResponse() (response *UpdateStackResponse) {
    response = &UpdateStackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateStack
// 本接口（UpdateStack）用于更新资源栈的名称和描述。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTEXIST = "FailedOperation.NotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UpdateStack(request *UpdateStackRequest) (response *UpdateStackResponse, err error) {
    return c.UpdateStackWithContext(context.Background(), request)
}

// UpdateStack
// 本接口（UpdateStack）用于更新资源栈的名称和描述。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTEXIST = "FailedOperation.NotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UpdateStackWithContext(ctx context.Context, request *UpdateStackRequest) (response *UpdateStackResponse, err error) {
    if request == nil {
        request = NewUpdateStackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateStack require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateStackResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateStackVersionRequest() (request *UpdateStackVersionRequest) {
    request = &UpdateStackVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tic", APIVersion, "UpdateStackVersion")
    
    
    return
}

func NewUpdateStackVersionResponse() (response *UpdateStackVersionResponse) {
    response = &UpdateStackVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateStackVersion
// 本接口（UpdateStackVersion）用于更新一个版本的模版内容，名称或描述，模版仅限COS URL，且为zip格式。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) UpdateStackVersion(request *UpdateStackVersionRequest) (response *UpdateStackVersionResponse, err error) {
    return c.UpdateStackVersionWithContext(context.Background(), request)
}

// UpdateStackVersion
// 本接口（UpdateStackVersion）用于更新一个版本的模版内容，名称或描述，模版仅限COS URL，且为zip格式。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) UpdateStackVersionWithContext(ctx context.Context, request *UpdateStackVersionRequest) (response *UpdateStackVersionResponse, err error) {
    if request == nil {
        request = NewUpdateStackVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateStackVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateStackVersionResponse()
    err = c.Send(request, response)
    return
}
