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

package v20230418

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2023-04-18"

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


func NewConfirmCommonServiceWorkOrderRequest() (request *ConfirmCommonServiceWorkOrderRequest) {
    request = &ConfirmCommonServiceWorkOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chc", APIVersion, "ConfirmCommonServiceWorkOrder")
    
    
    return
}

func NewConfirmCommonServiceWorkOrderResponse() (response *ConfirmCommonServiceWorkOrderResponse) {
    response = &ConfirmCommonServiceWorkOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ConfirmCommonServiceWorkOrder
// 确认通用服务工单
func (c *Client) ConfirmCommonServiceWorkOrder(request *ConfirmCommonServiceWorkOrderRequest) (response *ConfirmCommonServiceWorkOrderResponse, err error) {
    return c.ConfirmCommonServiceWorkOrderWithContext(context.Background(), request)
}

// ConfirmCommonServiceWorkOrder
// 确认通用服务工单
func (c *Client) ConfirmCommonServiceWorkOrderWithContext(ctx context.Context, request *ConfirmCommonServiceWorkOrderRequest) (response *ConfirmCommonServiceWorkOrderResponse, err error) {
    if request == nil {
        request = NewConfirmCommonServiceWorkOrderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "chc", APIVersion, "ConfirmCommonServiceWorkOrder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ConfirmCommonServiceWorkOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewConfirmCommonServiceWorkOrderResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCommonServiceWorkOrderRequest() (request *CreateCommonServiceWorkOrderRequest) {
    request = &CreateCommonServiceWorkOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chc", APIVersion, "CreateCommonServiceWorkOrder")
    
    
    return
}

func NewCreateCommonServiceWorkOrderResponse() (response *CreateCommonServiceWorkOrderResponse) {
    response = &CreateCommonServiceWorkOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCommonServiceWorkOrder
// 创建通用工单
func (c *Client) CreateCommonServiceWorkOrder(request *CreateCommonServiceWorkOrderRequest) (response *CreateCommonServiceWorkOrderResponse, err error) {
    return c.CreateCommonServiceWorkOrderWithContext(context.Background(), request)
}

// CreateCommonServiceWorkOrder
// 创建通用工单
func (c *Client) CreateCommonServiceWorkOrderWithContext(ctx context.Context, request *CreateCommonServiceWorkOrderRequest) (response *CreateCommonServiceWorkOrderResponse, err error) {
    if request == nil {
        request = NewCreateCommonServiceWorkOrderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "chc", APIVersion, "CreateCommonServiceWorkOrder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCommonServiceWorkOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCommonServiceWorkOrderResponse()
    err = c.Send(request, response)
    return
}

func NewCreateModelEvaluationWorkOrderRequest() (request *CreateModelEvaluationWorkOrderRequest) {
    request = &CreateModelEvaluationWorkOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chc", APIVersion, "CreateModelEvaluationWorkOrder")
    
    
    return
}

func NewCreateModelEvaluationWorkOrderResponse() (response *CreateModelEvaluationWorkOrderResponse) {
    response = &CreateModelEvaluationWorkOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateModelEvaluationWorkOrder
// 创建设备型号评估工单
func (c *Client) CreateModelEvaluationWorkOrder(request *CreateModelEvaluationWorkOrderRequest) (response *CreateModelEvaluationWorkOrderResponse, err error) {
    return c.CreateModelEvaluationWorkOrderWithContext(context.Background(), request)
}

// CreateModelEvaluationWorkOrder
// 创建设备型号评估工单
func (c *Client) CreateModelEvaluationWorkOrderWithContext(ctx context.Context, request *CreateModelEvaluationWorkOrderRequest) (response *CreateModelEvaluationWorkOrderResponse, err error) {
    if request == nil {
        request = NewCreateModelEvaluationWorkOrderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "chc", APIVersion, "CreateModelEvaluationWorkOrder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateModelEvaluationWorkOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateModelEvaluationWorkOrderResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMovingWorkOrderRequest() (request *CreateMovingWorkOrderRequest) {
    request = &CreateMovingWorkOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chc", APIVersion, "CreateMovingWorkOrder")
    
    
    return
}

func NewCreateMovingWorkOrderResponse() (response *CreateMovingWorkOrderResponse) {
    response = &CreateMovingWorkOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateMovingWorkOrder
// 创建设备搬迁工单
func (c *Client) CreateMovingWorkOrder(request *CreateMovingWorkOrderRequest) (response *CreateMovingWorkOrderResponse, err error) {
    return c.CreateMovingWorkOrderWithContext(context.Background(), request)
}

// CreateMovingWorkOrder
// 创建设备搬迁工单
func (c *Client) CreateMovingWorkOrderWithContext(ctx context.Context, request *CreateMovingWorkOrderRequest) (response *CreateMovingWorkOrderResponse, err error) {
    if request == nil {
        request = NewCreateMovingWorkOrderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "chc", APIVersion, "CreateMovingWorkOrder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMovingWorkOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMovingWorkOrderResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNetDeviceModelRequest() (request *CreateNetDeviceModelRequest) {
    request = &CreateNetDeviceModelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chc", APIVersion, "CreateNetDeviceModel")
    
    
    return
}

func NewCreateNetDeviceModelResponse() (response *CreateNetDeviceModelResponse) {
    response = &CreateNetDeviceModelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateNetDeviceModel
// 创建新的网络设备型号
func (c *Client) CreateNetDeviceModel(request *CreateNetDeviceModelRequest) (response *CreateNetDeviceModelResponse, err error) {
    return c.CreateNetDeviceModelWithContext(context.Background(), request)
}

// CreateNetDeviceModel
// 创建新的网络设备型号
func (c *Client) CreateNetDeviceModelWithContext(ctx context.Context, request *CreateNetDeviceModelRequest) (response *CreateNetDeviceModelResponse, err error) {
    if request == nil {
        request = NewCreateNetDeviceModelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "chc", APIVersion, "CreateNetDeviceModel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNetDeviceModel require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateNetDeviceModelResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePersonnelVisitWorkOrderRequest() (request *CreatePersonnelVisitWorkOrderRequest) {
    request = &CreatePersonnelVisitWorkOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chc", APIVersion, "CreatePersonnelVisitWorkOrder")
    
    
    return
}

func NewCreatePersonnelVisitWorkOrderResponse() (response *CreatePersonnelVisitWorkOrderResponse) {
    response = &CreatePersonnelVisitWorkOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePersonnelVisitWorkOrder
// 创建人员到访工单
func (c *Client) CreatePersonnelVisitWorkOrder(request *CreatePersonnelVisitWorkOrderRequest) (response *CreatePersonnelVisitWorkOrderResponse, err error) {
    return c.CreatePersonnelVisitWorkOrderWithContext(context.Background(), request)
}

// CreatePersonnelVisitWorkOrder
// 创建人员到访工单
func (c *Client) CreatePersonnelVisitWorkOrderWithContext(ctx context.Context, request *CreatePersonnelVisitWorkOrderRequest) (response *CreatePersonnelVisitWorkOrderResponse, err error) {
    if request == nil {
        request = NewCreatePersonnelVisitWorkOrderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "chc", APIVersion, "CreatePersonnelVisitWorkOrder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePersonnelVisitWorkOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePersonnelVisitWorkOrderResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePowerOffWorkOrderRequest() (request *CreatePowerOffWorkOrderRequest) {
    request = &CreatePowerOffWorkOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chc", APIVersion, "CreatePowerOffWorkOrder")
    
    
    return
}

func NewCreatePowerOffWorkOrderResponse() (response *CreatePowerOffWorkOrderResponse) {
    response = &CreatePowerOffWorkOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePowerOffWorkOrder
// 创建设备关电工单
func (c *Client) CreatePowerOffWorkOrder(request *CreatePowerOffWorkOrderRequest) (response *CreatePowerOffWorkOrderResponse, err error) {
    return c.CreatePowerOffWorkOrderWithContext(context.Background(), request)
}

// CreatePowerOffWorkOrder
// 创建设备关电工单
func (c *Client) CreatePowerOffWorkOrderWithContext(ctx context.Context, request *CreatePowerOffWorkOrderRequest) (response *CreatePowerOffWorkOrderResponse, err error) {
    if request == nil {
        request = NewCreatePowerOffWorkOrderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "chc", APIVersion, "CreatePowerOffWorkOrder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePowerOffWorkOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePowerOffWorkOrderResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePowerOnWorkOrderRequest() (request *CreatePowerOnWorkOrderRequest) {
    request = &CreatePowerOnWorkOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chc", APIVersion, "CreatePowerOnWorkOrder")
    
    
    return
}

func NewCreatePowerOnWorkOrderResponse() (response *CreatePowerOnWorkOrderResponse) {
    response = &CreatePowerOnWorkOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePowerOnWorkOrder
// 创建设备开电工单
func (c *Client) CreatePowerOnWorkOrder(request *CreatePowerOnWorkOrderRequest) (response *CreatePowerOnWorkOrderResponse, err error) {
    return c.CreatePowerOnWorkOrderWithContext(context.Background(), request)
}

// CreatePowerOnWorkOrder
// 创建设备开电工单
func (c *Client) CreatePowerOnWorkOrderWithContext(ctx context.Context, request *CreatePowerOnWorkOrderRequest) (response *CreatePowerOnWorkOrderResponse, err error) {
    if request == nil {
        request = NewCreatePowerOnWorkOrderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "chc", APIVersion, "CreatePowerOnWorkOrder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePowerOnWorkOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePowerOnWorkOrderResponse()
    err = c.Send(request, response)
    return
}

func NewCreateQuitWorkOrderRequest() (request *CreateQuitWorkOrderRequest) {
    request = &CreateQuitWorkOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chc", APIVersion, "CreateQuitWorkOrder")
    
    
    return
}

func NewCreateQuitWorkOrderResponse() (response *CreateQuitWorkOrderResponse) {
    response = &CreateQuitWorkOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateQuitWorkOrder
// 创建设备退出工单
func (c *Client) CreateQuitWorkOrder(request *CreateQuitWorkOrderRequest) (response *CreateQuitWorkOrderResponse, err error) {
    return c.CreateQuitWorkOrderWithContext(context.Background(), request)
}

// CreateQuitWorkOrder
// 创建设备退出工单
func (c *Client) CreateQuitWorkOrderWithContext(ctx context.Context, request *CreateQuitWorkOrderRequest) (response *CreateQuitWorkOrderResponse, err error) {
    if request == nil {
        request = NewCreateQuitWorkOrderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "chc", APIVersion, "CreateQuitWorkOrder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateQuitWorkOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateQuitWorkOrderResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRackOffWorkOrderRequest() (request *CreateRackOffWorkOrderRequest) {
    request = &CreateRackOffWorkOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chc", APIVersion, "CreateRackOffWorkOrder")
    
    
    return
}

func NewCreateRackOffWorkOrderResponse() (response *CreateRackOffWorkOrderResponse) {
    response = &CreateRackOffWorkOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRackOffWorkOrder
// 创建设备下架工单
func (c *Client) CreateRackOffWorkOrder(request *CreateRackOffWorkOrderRequest) (response *CreateRackOffWorkOrderResponse, err error) {
    return c.CreateRackOffWorkOrderWithContext(context.Background(), request)
}

// CreateRackOffWorkOrder
// 创建设备下架工单
func (c *Client) CreateRackOffWorkOrderWithContext(ctx context.Context, request *CreateRackOffWorkOrderRequest) (response *CreateRackOffWorkOrderResponse, err error) {
    if request == nil {
        request = NewCreateRackOffWorkOrderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "chc", APIVersion, "CreateRackOffWorkOrder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRackOffWorkOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRackOffWorkOrderResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRackOnWorkOrderRequest() (request *CreateRackOnWorkOrderRequest) {
    request = &CreateRackOnWorkOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chc", APIVersion, "CreateRackOnWorkOrder")
    
    
    return
}

func NewCreateRackOnWorkOrderResponse() (response *CreateRackOnWorkOrderResponse) {
    response = &CreateRackOnWorkOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRackOnWorkOrder
// 创建设备上架工单
func (c *Client) CreateRackOnWorkOrder(request *CreateRackOnWorkOrderRequest) (response *CreateRackOnWorkOrderResponse, err error) {
    return c.CreateRackOnWorkOrderWithContext(context.Background(), request)
}

// CreateRackOnWorkOrder
// 创建设备上架工单
func (c *Client) CreateRackOnWorkOrderWithContext(ctx context.Context, request *CreateRackOnWorkOrderRequest) (response *CreateRackOnWorkOrderResponse, err error) {
    if request == nil {
        request = NewCreateRackOnWorkOrderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "chc", APIVersion, "CreateRackOnWorkOrder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRackOnWorkOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRackOnWorkOrderResponse()
    err = c.Send(request, response)
    return
}

func NewCreateReceivingWorkOrderRequest() (request *CreateReceivingWorkOrderRequest) {
    request = &CreateReceivingWorkOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chc", APIVersion, "CreateReceivingWorkOrder")
    
    
    return
}

func NewCreateReceivingWorkOrderResponse() (response *CreateReceivingWorkOrderResponse) {
    response = &CreateReceivingWorkOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateReceivingWorkOrder
// 创建设备收货工单
func (c *Client) CreateReceivingWorkOrder(request *CreateReceivingWorkOrderRequest) (response *CreateReceivingWorkOrderResponse, err error) {
    return c.CreateReceivingWorkOrderWithContext(context.Background(), request)
}

// CreateReceivingWorkOrder
// 创建设备收货工单
func (c *Client) CreateReceivingWorkOrderWithContext(ctx context.Context, request *CreateReceivingWorkOrderRequest) (response *CreateReceivingWorkOrderResponse, err error) {
    if request == nil {
        request = NewCreateReceivingWorkOrderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "chc", APIVersion, "CreateReceivingWorkOrder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateReceivingWorkOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateReceivingWorkOrderResponse()
    err = c.Send(request, response)
    return
}

func NewCreateServerModelRequest() (request *CreateServerModelRequest) {
    request = &CreateServerModelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chc", APIVersion, "CreateServerModel")
    
    
    return
}

func NewCreateServerModelResponse() (response *CreateServerModelResponse) {
    response = &CreateServerModelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateServerModel
// 新增服务器设备型号
func (c *Client) CreateServerModel(request *CreateServerModelRequest) (response *CreateServerModelResponse, err error) {
    return c.CreateServerModelWithContext(context.Background(), request)
}

// CreateServerModel
// 新增服务器设备型号
func (c *Client) CreateServerModelWithContext(ctx context.Context, request *CreateServerModelRequest) (response *CreateServerModelResponse, err error) {
    if request == nil {
        request = NewCreateServerModelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "chc", APIVersion, "CreateServerModel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateServerModel require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateServerModelResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSpeciallyQuitWorkOrderRequest() (request *CreateSpeciallyQuitWorkOrderRequest) {
    request = &CreateSpeciallyQuitWorkOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chc", APIVersion, "CreateSpeciallyQuitWorkOrder")
    
    
    return
}

func NewCreateSpeciallyQuitWorkOrderResponse() (response *CreateSpeciallyQuitWorkOrderResponse) {
    response = &CreateSpeciallyQuitWorkOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSpeciallyQuitWorkOrder
// 创建临时设备退出工单
func (c *Client) CreateSpeciallyQuitWorkOrder(request *CreateSpeciallyQuitWorkOrderRequest) (response *CreateSpeciallyQuitWorkOrderResponse, err error) {
    return c.CreateSpeciallyQuitWorkOrderWithContext(context.Background(), request)
}

// CreateSpeciallyQuitWorkOrder
// 创建临时设备退出工单
func (c *Client) CreateSpeciallyQuitWorkOrderWithContext(ctx context.Context, request *CreateSpeciallyQuitWorkOrderRequest) (response *CreateSpeciallyQuitWorkOrderResponse, err error) {
    if request == nil {
        request = NewCreateSpeciallyQuitWorkOrderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "chc", APIVersion, "CreateSpeciallyQuitWorkOrder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSpeciallyQuitWorkOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSpeciallyQuitWorkOrderResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAvailableModelListRequest() (request *DescribeAvailableModelListRequest) {
    request = &DescribeAvailableModelListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chc", APIVersion, "DescribeAvailableModelList")
    
    
    return
}

func NewDescribeAvailableModelListResponse() (response *DescribeAvailableModelListResponse) {
    response = &DescribeAvailableModelListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAvailableModelList
// 获取机房内可用的型号列表
func (c *Client) DescribeAvailableModelList(request *DescribeAvailableModelListRequest) (response *DescribeAvailableModelListResponse, err error) {
    return c.DescribeAvailableModelListWithContext(context.Background(), request)
}

// DescribeAvailableModelList
// 获取机房内可用的型号列表
func (c *Client) DescribeAvailableModelListWithContext(ctx context.Context, request *DescribeAvailableModelListRequest) (response *DescribeAvailableModelListResponse, err error) {
    if request == nil {
        request = NewDescribeAvailableModelListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "chc", APIVersion, "DescribeAvailableModelList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAvailableModelList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAvailableModelListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCampusListRequest() (request *DescribeCampusListRequest) {
    request = &DescribeCampusListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chc", APIVersion, "DescribeCampusList")
    
    
    return
}

func NewDescribeCampusListResponse() (response *DescribeCampusListResponse) {
    response = &DescribeCampusListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCampusList
// 获取用户可操作的园区列表
func (c *Client) DescribeCampusList(request *DescribeCampusListRequest) (response *DescribeCampusListResponse, err error) {
    return c.DescribeCampusListWithContext(context.Background(), request)
}

// DescribeCampusList
// 获取用户可操作的园区列表
func (c *Client) DescribeCampusListWithContext(ctx context.Context, request *DescribeCampusListRequest) (response *DescribeCampusListResponse, err error) {
    if request == nil {
        request = NewDescribeCampusListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "chc", APIVersion, "DescribeCampusList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCampusList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCampusListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCommonServiceWorkOrderDetailRequest() (request *DescribeCommonServiceWorkOrderDetailRequest) {
    request = &DescribeCommonServiceWorkOrderDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chc", APIVersion, "DescribeCommonServiceWorkOrderDetail")
    
    
    return
}

func NewDescribeCommonServiceWorkOrderDetailResponse() (response *DescribeCommonServiceWorkOrderDetailResponse) {
    response = &DescribeCommonServiceWorkOrderDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCommonServiceWorkOrderDetail
// 查询通用服务工单详情
func (c *Client) DescribeCommonServiceWorkOrderDetail(request *DescribeCommonServiceWorkOrderDetailRequest) (response *DescribeCommonServiceWorkOrderDetailResponse, err error) {
    return c.DescribeCommonServiceWorkOrderDetailWithContext(context.Background(), request)
}

// DescribeCommonServiceWorkOrderDetail
// 查询通用服务工单详情
func (c *Client) DescribeCommonServiceWorkOrderDetailWithContext(ctx context.Context, request *DescribeCommonServiceWorkOrderDetailRequest) (response *DescribeCommonServiceWorkOrderDetailResponse, err error) {
    if request == nil {
        request = NewDescribeCommonServiceWorkOrderDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "chc", APIVersion, "DescribeCommonServiceWorkOrderDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCommonServiceWorkOrderDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCommonServiceWorkOrderDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomerInfoRequest() (request *DescribeCustomerInfoRequest) {
    request = &DescribeCustomerInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chc", APIVersion, "DescribeCustomerInfo")
    
    
    return
}

func NewDescribeCustomerInfoResponse() (response *DescribeCustomerInfoResponse) {
    response = &DescribeCustomerInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCustomerInfo
// 查询客户信息
func (c *Client) DescribeCustomerInfo(request *DescribeCustomerInfoRequest) (response *DescribeCustomerInfoResponse, err error) {
    return c.DescribeCustomerInfoWithContext(context.Background(), request)
}

// DescribeCustomerInfo
// 查询客户信息
func (c *Client) DescribeCustomerInfoWithContext(ctx context.Context, request *DescribeCustomerInfoRequest) (response *DescribeCustomerInfoResponse, err error) {
    if request == nil {
        request = NewDescribeCustomerInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "chc", APIVersion, "DescribeCustomerInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCustomerInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCustomerInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceListRequest() (request *DescribeDeviceListRequest) {
    request = &DescribeDeviceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chc", APIVersion, "DescribeDeviceList")
    
    
    return
}

func NewDescribeDeviceListResponse() (response *DescribeDeviceListResponse) {
    response = &DescribeDeviceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeviceList
// 获取设备列表
func (c *Client) DescribeDeviceList(request *DescribeDeviceListRequest) (response *DescribeDeviceListResponse, err error) {
    return c.DescribeDeviceListWithContext(context.Background(), request)
}

// DescribeDeviceList
// 获取设备列表
func (c *Client) DescribeDeviceListWithContext(ctx context.Context, request *DescribeDeviceListRequest) (response *DescribeDeviceListResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "chc", APIVersion, "DescribeDeviceList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeviceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceWorkOrderDetailRequest() (request *DescribeDeviceWorkOrderDetailRequest) {
    request = &DescribeDeviceWorkOrderDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chc", APIVersion, "DescribeDeviceWorkOrderDetail")
    
    
    return
}

func NewDescribeDeviceWorkOrderDetailResponse() (response *DescribeDeviceWorkOrderDetailResponse) {
    response = &DescribeDeviceWorkOrderDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeviceWorkOrderDetail
// 用于查询设备类工单的工单详情，如：'receiving', 'rackOn', 'powerOn', 'powerOff', 'rackOff', 'quit'
func (c *Client) DescribeDeviceWorkOrderDetail(request *DescribeDeviceWorkOrderDetailRequest) (response *DescribeDeviceWorkOrderDetailResponse, err error) {
    return c.DescribeDeviceWorkOrderDetailWithContext(context.Background(), request)
}

// DescribeDeviceWorkOrderDetail
// 用于查询设备类工单的工单详情，如：'receiving', 'rackOn', 'powerOn', 'powerOff', 'rackOff', 'quit'
func (c *Client) DescribeDeviceWorkOrderDetailWithContext(ctx context.Context, request *DescribeDeviceWorkOrderDetailRequest) (response *DescribeDeviceWorkOrderDetailResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceWorkOrderDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "chc", APIVersion, "DescribeDeviceWorkOrderDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceWorkOrderDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeviceWorkOrderDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIdcUnitAssetDetailRequest() (request *DescribeIdcUnitAssetDetailRequest) {
    request = &DescribeIdcUnitAssetDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chc", APIVersion, "DescribeIdcUnitAssetDetail")
    
    
    return
}

func NewDescribeIdcUnitAssetDetailResponse() (response *DescribeIdcUnitAssetDetailResponse) {
    response = &DescribeIdcUnitAssetDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIdcUnitAssetDetail
// 查询机房管理单元资产详情
func (c *Client) DescribeIdcUnitAssetDetail(request *DescribeIdcUnitAssetDetailRequest) (response *DescribeIdcUnitAssetDetailResponse, err error) {
    return c.DescribeIdcUnitAssetDetailWithContext(context.Background(), request)
}

// DescribeIdcUnitAssetDetail
// 查询机房管理单元资产详情
func (c *Client) DescribeIdcUnitAssetDetailWithContext(ctx context.Context, request *DescribeIdcUnitAssetDetailRequest) (response *DescribeIdcUnitAssetDetailResponse, err error) {
    if request == nil {
        request = NewDescribeIdcUnitAssetDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "chc", APIVersion, "DescribeIdcUnitAssetDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIdcUnitAssetDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIdcUnitAssetDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIdcUnitDetailRequest() (request *DescribeIdcUnitDetailRequest) {
    request = &DescribeIdcUnitDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chc", APIVersion, "DescribeIdcUnitDetail")
    
    
    return
}

func NewDescribeIdcUnitDetailResponse() (response *DescribeIdcUnitDetailResponse) {
    response = &DescribeIdcUnitDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIdcUnitDetail
// 查询机房管理单元详情
func (c *Client) DescribeIdcUnitDetail(request *DescribeIdcUnitDetailRequest) (response *DescribeIdcUnitDetailResponse, err error) {
    return c.DescribeIdcUnitDetailWithContext(context.Background(), request)
}

// DescribeIdcUnitDetail
// 查询机房管理单元详情
func (c *Client) DescribeIdcUnitDetailWithContext(ctx context.Context, request *DescribeIdcUnitDetailRequest) (response *DescribeIdcUnitDetailResponse, err error) {
    if request == nil {
        request = NewDescribeIdcUnitDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "chc", APIVersion, "DescribeIdcUnitDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIdcUnitDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIdcUnitDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIdcsRequest() (request *DescribeIdcsRequest) {
    request = &DescribeIdcsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chc", APIVersion, "DescribeIdcs")
    
    
    return
}

func NewDescribeIdcsResponse() (response *DescribeIdcsResponse) {
    response = &DescribeIdcsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIdcs
// 获取机房和机房管理单元信息
func (c *Client) DescribeIdcs(request *DescribeIdcsRequest) (response *DescribeIdcsResponse, err error) {
    return c.DescribeIdcsWithContext(context.Background(), request)
}

// DescribeIdcs
// 获取机房和机房管理单元信息
func (c *Client) DescribeIdcsWithContext(ctx context.Context, request *DescribeIdcsRequest) (response *DescribeIdcsResponse, err error) {
    if request == nil {
        request = NewDescribeIdcsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "chc", APIVersion, "DescribeIdcs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIdcs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIdcsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeModelRequest() (request *DescribeModelRequest) {
    request = &DescribeModelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chc", APIVersion, "DescribeModel")
    
    
    return
}

func NewDescribeModelResponse() (response *DescribeModelResponse) {
    response = &DescribeModelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeModel
// 查询设备型号详情
func (c *Client) DescribeModel(request *DescribeModelRequest) (response *DescribeModelResponse, err error) {
    return c.DescribeModelWithContext(context.Background(), request)
}

// DescribeModel
// 查询设备型号详情
func (c *Client) DescribeModelWithContext(ctx context.Context, request *DescribeModelRequest) (response *DescribeModelResponse, err error) {
    if request == nil {
        request = NewDescribeModelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "chc", APIVersion, "DescribeModel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeModel require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeModelResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeModelEvaluationWorkOrderDetailRequest() (request *DescribeModelEvaluationWorkOrderDetailRequest) {
    request = &DescribeModelEvaluationWorkOrderDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chc", APIVersion, "DescribeModelEvaluationWorkOrderDetail")
    
    
    return
}

func NewDescribeModelEvaluationWorkOrderDetailResponse() (response *DescribeModelEvaluationWorkOrderDetailResponse) {
    response = &DescribeModelEvaluationWorkOrderDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeModelEvaluationWorkOrderDetail
// 查询设备型号评估工单详情
func (c *Client) DescribeModelEvaluationWorkOrderDetail(request *DescribeModelEvaluationWorkOrderDetailRequest) (response *DescribeModelEvaluationWorkOrderDetailResponse, err error) {
    return c.DescribeModelEvaluationWorkOrderDetailWithContext(context.Background(), request)
}

// DescribeModelEvaluationWorkOrderDetail
// 查询设备型号评估工单详情
func (c *Client) DescribeModelEvaluationWorkOrderDetailWithContext(ctx context.Context, request *DescribeModelEvaluationWorkOrderDetailRequest) (response *DescribeModelEvaluationWorkOrderDetailResponse, err error) {
    if request == nil {
        request = NewDescribeModelEvaluationWorkOrderDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "chc", APIVersion, "DescribeModelEvaluationWorkOrderDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeModelEvaluationWorkOrderDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeModelEvaluationWorkOrderDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeModelTemplateRequest() (request *DescribeModelTemplateRequest) {
    request = &DescribeModelTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chc", APIVersion, "DescribeModelTemplate")
    
    
    return
}

func NewDescribeModelTemplateResponse() (response *DescribeModelTemplateResponse) {
    response = &DescribeModelTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeModelTemplate
// 获取型号的填写模板
func (c *Client) DescribeModelTemplate(request *DescribeModelTemplateRequest) (response *DescribeModelTemplateResponse, err error) {
    return c.DescribeModelTemplateWithContext(context.Background(), request)
}

// DescribeModelTemplate
// 获取型号的填写模板
func (c *Client) DescribeModelTemplateWithContext(ctx context.Context, request *DescribeModelTemplateRequest) (response *DescribeModelTemplateResponse, err error) {
    if request == nil {
        request = NewDescribeModelTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "chc", APIVersion, "DescribeModelTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeModelTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeModelTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeModelVersionListRequest() (request *DescribeModelVersionListRequest) {
    request = &DescribeModelVersionListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chc", APIVersion, "DescribeModelVersionList")
    
    
    return
}

func NewDescribeModelVersionListResponse() (response *DescribeModelVersionListResponse) {
    response = &DescribeModelVersionListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeModelVersionList
// 获取用户的型号和对应的版本数量
func (c *Client) DescribeModelVersionList(request *DescribeModelVersionListRequest) (response *DescribeModelVersionListResponse, err error) {
    return c.DescribeModelVersionListWithContext(context.Background(), request)
}

// DescribeModelVersionList
// 获取用户的型号和对应的版本数量
func (c *Client) DescribeModelVersionListWithContext(ctx context.Context, request *DescribeModelVersionListRequest) (response *DescribeModelVersionListResponse, err error) {
    if request == nil {
        request = NewDescribeModelVersionListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "chc", APIVersion, "DescribeModelVersionList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeModelVersionList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeModelVersionListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePersonnelVisitWorkOrderDetailRequest() (request *DescribePersonnelVisitWorkOrderDetailRequest) {
    request = &DescribePersonnelVisitWorkOrderDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chc", APIVersion, "DescribePersonnelVisitWorkOrderDetail")
    
    
    return
}

func NewDescribePersonnelVisitWorkOrderDetailResponse() (response *DescribePersonnelVisitWorkOrderDetailResponse) {
    response = &DescribePersonnelVisitWorkOrderDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePersonnelVisitWorkOrderDetail
// 查询人员到访工单详情
func (c *Client) DescribePersonnelVisitWorkOrderDetail(request *DescribePersonnelVisitWorkOrderDetailRequest) (response *DescribePersonnelVisitWorkOrderDetailResponse, err error) {
    return c.DescribePersonnelVisitWorkOrderDetailWithContext(context.Background(), request)
}

// DescribePersonnelVisitWorkOrderDetail
// 查询人员到访工单详情
func (c *Client) DescribePersonnelVisitWorkOrderDetailWithContext(ctx context.Context, request *DescribePersonnelVisitWorkOrderDetailRequest) (response *DescribePersonnelVisitWorkOrderDetailResponse, err error) {
    if request == nil {
        request = NewDescribePersonnelVisitWorkOrderDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "chc", APIVersion, "DescribePersonnelVisitWorkOrderDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePersonnelVisitWorkOrderDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePersonnelVisitWorkOrderDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePositionStatusSummaryRequest() (request *DescribePositionStatusSummaryRequest) {
    request = &DescribePositionStatusSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chc", APIVersion, "DescribePositionStatusSummary")
    
    
    return
}

func NewDescribePositionStatusSummaryResponse() (response *DescribePositionStatusSummaryResponse) {
    response = &DescribePositionStatusSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePositionStatusSummary
// 获取机架总数及各状态对应的数量汇总
func (c *Client) DescribePositionStatusSummary(request *DescribePositionStatusSummaryRequest) (response *DescribePositionStatusSummaryResponse, err error) {
    return c.DescribePositionStatusSummaryWithContext(context.Background(), request)
}

// DescribePositionStatusSummary
// 获取机架总数及各状态对应的数量汇总
func (c *Client) DescribePositionStatusSummaryWithContext(ctx context.Context, request *DescribePositionStatusSummaryRequest) (response *DescribePositionStatusSummaryResponse, err error) {
    if request == nil {
        request = NewDescribePositionStatusSummaryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "chc", APIVersion, "DescribePositionStatusSummary")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePositionStatusSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePositionStatusSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePositionsRequest() (request *DescribePositionsRequest) {
    request = &DescribePositionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chc", APIVersion, "DescribePositions")
    
    
    return
}

func NewDescribePositionsResponse() (response *DescribePositionsResponse) {
    response = &DescribePositionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePositions
// 获取机位列表
func (c *Client) DescribePositions(request *DescribePositionsRequest) (response *DescribePositionsResponse, err error) {
    return c.DescribePositionsWithContext(context.Background(), request)
}

// DescribePositions
// 获取机位列表
func (c *Client) DescribePositionsWithContext(ctx context.Context, request *DescribePositionsRequest) (response *DescribePositionsResponse, err error) {
    if request == nil {
        request = NewDescribePositionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "chc", APIVersion, "DescribePositions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePositions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePositionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRacksRequest() (request *DescribeRacksRequest) {
    request = &DescribeRacksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chc", APIVersion, "DescribeRacks")
    
    
    return
}

func NewDescribeRacksResponse() (response *DescribeRacksResponse) {
    response = &DescribeRacksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRacks
// 获取机架列表
func (c *Client) DescribeRacks(request *DescribeRacksRequest) (response *DescribeRacksResponse, err error) {
    return c.DescribeRacksWithContext(context.Background(), request)
}

// DescribeRacks
// 获取机架列表
func (c *Client) DescribeRacksWithContext(ctx context.Context, request *DescribeRacksRequest) (response *DescribeRacksResponse, err error) {
    if request == nil {
        request = NewDescribeRacksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "chc", APIVersion, "DescribeRacks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRacks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRacksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRacksDistributionRequest() (request *DescribeRacksDistributionRequest) {
    request = &DescribeRacksDistributionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chc", APIVersion, "DescribeRacksDistribution")
    
    
    return
}

func NewDescribeRacksDistributionResponse() (response *DescribeRacksDistributionResponse) {
    response = &DescribeRacksDistributionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRacksDistribution
// 获取机房管理单元的机位分布
func (c *Client) DescribeRacksDistribution(request *DescribeRacksDistributionRequest) (response *DescribeRacksDistributionResponse, err error) {
    return c.DescribeRacksDistributionWithContext(context.Background(), request)
}

// DescribeRacksDistribution
// 获取机房管理单元的机位分布
func (c *Client) DescribeRacksDistributionWithContext(ctx context.Context, request *DescribeRacksDistributionRequest) (response *DescribeRacksDistributionResponse, err error) {
    if request == nil {
        request = NewDescribeRacksDistributionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "chc", APIVersion, "DescribeRacksDistribution")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRacksDistribution require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRacksDistributionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourceUsageRequest() (request *DescribeResourceUsageRequest) {
    request = &DescribeResourceUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chc", APIVersion, "DescribeResourceUsage")
    
    
    return
}

func NewDescribeResourceUsageResponse() (response *DescribeResourceUsageResponse) {
    response = &DescribeResourceUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeResourceUsage
// 查询资源汇总
func (c *Client) DescribeResourceUsage(request *DescribeResourceUsageRequest) (response *DescribeResourceUsageResponse, err error) {
    return c.DescribeResourceUsageWithContext(context.Background(), request)
}

// DescribeResourceUsage
// 查询资源汇总
func (c *Client) DescribeResourceUsageWithContext(ctx context.Context, request *DescribeResourceUsageRequest) (response *DescribeResourceUsageResponse, err error) {
    if request == nil {
        request = NewDescribeResourceUsageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "chc", APIVersion, "DescribeResourceUsage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResourceUsage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResourceUsageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWorkOrderListRequest() (request *DescribeWorkOrderListRequest) {
    request = &DescribeWorkOrderListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chc", APIVersion, "DescribeWorkOrderList")
    
    
    return
}

func NewDescribeWorkOrderListResponse() (response *DescribeWorkOrderListResponse) {
    response = &DescribeWorkOrderListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWorkOrderList
// 查询工单列表
func (c *Client) DescribeWorkOrderList(request *DescribeWorkOrderListRequest) (response *DescribeWorkOrderListResponse, err error) {
    return c.DescribeWorkOrderListWithContext(context.Background(), request)
}

// DescribeWorkOrderList
// 查询工单列表
func (c *Client) DescribeWorkOrderListWithContext(ctx context.Context, request *DescribeWorkOrderListRequest) (response *DescribeWorkOrderListResponse, err error) {
    if request == nil {
        request = NewDescribeWorkOrderListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "chc", APIVersion, "DescribeWorkOrderList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWorkOrderList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWorkOrderListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWorkOrderStatisticsRequest() (request *DescribeWorkOrderStatisticsRequest) {
    request = &DescribeWorkOrderStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chc", APIVersion, "DescribeWorkOrderStatistics")
    
    
    return
}

func NewDescribeWorkOrderStatisticsResponse() (response *DescribeWorkOrderStatisticsResponse) {
    response = &DescribeWorkOrderStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWorkOrderStatistics
// 工单统计数据查询
func (c *Client) DescribeWorkOrderStatistics(request *DescribeWorkOrderStatisticsRequest) (response *DescribeWorkOrderStatisticsResponse, err error) {
    return c.DescribeWorkOrderStatisticsWithContext(context.Background(), request)
}

// DescribeWorkOrderStatistics
// 工单统计数据查询
func (c *Client) DescribeWorkOrderStatisticsWithContext(ctx context.Context, request *DescribeWorkOrderStatisticsRequest) (response *DescribeWorkOrderStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeWorkOrderStatisticsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "chc", APIVersion, "DescribeWorkOrderStatistics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWorkOrderStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWorkOrderStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWorkOrderTypesRequest() (request *DescribeWorkOrderTypesRequest) {
    request = &DescribeWorkOrderTypesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chc", APIVersion, "DescribeWorkOrderTypes")
    
    
    return
}

func NewDescribeWorkOrderTypesResponse() (response *DescribeWorkOrderTypesResponse) {
    response = &DescribeWorkOrderTypesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWorkOrderTypes
// 获取用户可用的工单类型
func (c *Client) DescribeWorkOrderTypes(request *DescribeWorkOrderTypesRequest) (response *DescribeWorkOrderTypesResponse, err error) {
    return c.DescribeWorkOrderTypesWithContext(context.Background(), request)
}

// DescribeWorkOrderTypes
// 获取用户可用的工单类型
func (c *Client) DescribeWorkOrderTypesWithContext(ctx context.Context, request *DescribeWorkOrderTypesRequest) (response *DescribeWorkOrderTypesResponse, err error) {
    if request == nil {
        request = NewDescribeWorkOrderTypesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "chc", APIVersion, "DescribeWorkOrderTypes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWorkOrderTypes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWorkOrderTypesResponse()
    err = c.Send(request, response)
    return
}

func NewExportCustomerWorkOrderDetailRequest() (request *ExportCustomerWorkOrderDetailRequest) {
    request = &ExportCustomerWorkOrderDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chc", APIVersion, "ExportCustomerWorkOrderDetail")
    
    
    return
}

func NewExportCustomerWorkOrderDetailResponse() (response *ExportCustomerWorkOrderDetailResponse) {
    response = &ExportCustomerWorkOrderDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ExportCustomerWorkOrderDetail
// 导出工单详情
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_ENDTIMELOWERTHANSTARTTIME = "InvalidParameterValue.EndTimeLowerThanStartTime"
//  INVALIDPARAMETERVALUE_INVALIDTIMEFORMAT = "InvalidParameterValue.InvalidTimeFormat"
//  INVALIDPARAMETERVALUE_INVALIDWORKORDERTYPE = "InvalidParameterValue.InvalidWorkOrderType"
func (c *Client) ExportCustomerWorkOrderDetail(request *ExportCustomerWorkOrderDetailRequest) (response *ExportCustomerWorkOrderDetailResponse, err error) {
    return c.ExportCustomerWorkOrderDetailWithContext(context.Background(), request)
}

// ExportCustomerWorkOrderDetail
// 导出工单详情
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_ENDTIMELOWERTHANSTARTTIME = "InvalidParameterValue.EndTimeLowerThanStartTime"
//  INVALIDPARAMETERVALUE_INVALIDTIMEFORMAT = "InvalidParameterValue.InvalidTimeFormat"
//  INVALIDPARAMETERVALUE_INVALIDWORKORDERTYPE = "InvalidParameterValue.InvalidWorkOrderType"
func (c *Client) ExportCustomerWorkOrderDetailWithContext(ctx context.Context, request *ExportCustomerWorkOrderDetailRequest) (response *ExportCustomerWorkOrderDetailResponse, err error) {
    if request == nil {
        request = NewExportCustomerWorkOrderDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "chc", APIVersion, "ExportCustomerWorkOrderDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportCustomerWorkOrderDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportCustomerWorkOrderDetailResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWorkOrderTypeCollectFlagRequest() (request *ModifyWorkOrderTypeCollectFlagRequest) {
    request = &ModifyWorkOrderTypeCollectFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chc", APIVersion, "ModifyWorkOrderTypeCollectFlag")
    
    
    return
}

func NewModifyWorkOrderTypeCollectFlagResponse() (response *ModifyWorkOrderTypeCollectFlagResponse) {
    response = &ModifyWorkOrderTypeCollectFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyWorkOrderTypeCollectFlag
// 如果当前该工单类型是收藏状态，调用接口后变成未收藏状态，如果是未收藏状态，调用该接口变为收藏状态
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_ENDTIMELOWERTHANSTARTTIME = "InvalidParameterValue.EndTimeLowerThanStartTime"
//  INVALIDPARAMETERVALUE_INVALIDTIMEFORMAT = "InvalidParameterValue.InvalidTimeFormat"
//  INVALIDPARAMETERVALUE_INVALIDWORKORDERTYPE = "InvalidParameterValue.InvalidWorkOrderType"
func (c *Client) ModifyWorkOrderTypeCollectFlag(request *ModifyWorkOrderTypeCollectFlagRequest) (response *ModifyWorkOrderTypeCollectFlagResponse, err error) {
    return c.ModifyWorkOrderTypeCollectFlagWithContext(context.Background(), request)
}

// ModifyWorkOrderTypeCollectFlag
// 如果当前该工单类型是收藏状态，调用接口后变成未收藏状态，如果是未收藏状态，调用该接口变为收藏状态
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_ENDTIMELOWERTHANSTARTTIME = "InvalidParameterValue.EndTimeLowerThanStartTime"
//  INVALIDPARAMETERVALUE_INVALIDTIMEFORMAT = "InvalidParameterValue.InvalidTimeFormat"
//  INVALIDPARAMETERVALUE_INVALIDWORKORDERTYPE = "InvalidParameterValue.InvalidWorkOrderType"
func (c *Client) ModifyWorkOrderTypeCollectFlagWithContext(ctx context.Context, request *ModifyWorkOrderTypeCollectFlagRequest) (response *ModifyWorkOrderTypeCollectFlagResponse, err error) {
    if request == nil {
        request = NewModifyWorkOrderTypeCollectFlagRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "chc", APIVersion, "ModifyWorkOrderTypeCollectFlag")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyWorkOrderTypeCollectFlag require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyWorkOrderTypeCollectFlagResponse()
    err = c.Send(request, response)
    return
}
