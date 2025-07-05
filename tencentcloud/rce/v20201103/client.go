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

package v20201103

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-11-03"

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


func NewCreateNameListRequest() (request *CreateNameListRequest) {
    request = &CreateNameListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rce", APIVersion, "CreateNameList")
    
    
    return
}

func NewCreateNameListResponse() (response *CreateNameListResponse) {
    response = &CreateNameListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateNameList
// 创建黑白名单，黑白名单数量上限为100
func (c *Client) CreateNameList(request *CreateNameListRequest) (response *CreateNameListResponse, err error) {
    return c.CreateNameListWithContext(context.Background(), request)
}

// CreateNameList
// 创建黑白名单，黑白名单数量上限为100
func (c *Client) CreateNameListWithContext(ctx context.Context, request *CreateNameListRequest) (response *CreateNameListResponse, err error) {
    if request == nil {
        request = NewCreateNameListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNameList require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateNameListResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNameListRequest() (request *DeleteNameListRequest) {
    request = &DeleteNameListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rce", APIVersion, "DeleteNameList")
    
    
    return
}

func NewDeleteNameListResponse() (response *DeleteNameListResponse) {
    response = &DeleteNameListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteNameList
// 修改黑白名单状态 关闭 开启 删除
func (c *Client) DeleteNameList(request *DeleteNameListRequest) (response *DeleteNameListResponse, err error) {
    return c.DeleteNameListWithContext(context.Background(), request)
}

// DeleteNameList
// 修改黑白名单状态 关闭 开启 删除
func (c *Client) DeleteNameListWithContext(ctx context.Context, request *DeleteNameListRequest) (response *DeleteNameListResponse, err error) {
    if request == nil {
        request = NewDeleteNameListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteNameList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteNameListResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNameListDataRequest() (request *DeleteNameListDataRequest) {
    request = &DeleteNameListDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rce", APIVersion, "DeleteNameListData")
    
    
    return
}

func NewDeleteNameListDataResponse() (response *DeleteNameListDataResponse) {
    response = &DeleteNameListDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteNameListData
// 删除黑白名单数据
func (c *Client) DeleteNameListData(request *DeleteNameListDataRequest) (response *DeleteNameListDataResponse, err error) {
    return c.DeleteNameListDataWithContext(context.Background(), request)
}

// DeleteNameListData
// 删除黑白名单数据
func (c *Client) DeleteNameListDataWithContext(ctx context.Context, request *DeleteNameListDataRequest) (response *DeleteNameListDataResponse, err error) {
    if request == nil {
        request = NewDeleteNameListDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteNameListData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteNameListDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNameListRequest() (request *DescribeNameListRequest) {
    request = &DescribeNameListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rce", APIVersion, "DescribeNameList")
    
    
    return
}

func NewDescribeNameListResponse() (response *DescribeNameListResponse) {
    response = &DescribeNameListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNameList
// 列表展示黑白名单列表数据, 包含列表名称, 名单类型, 数据类型, 数据来源, 描述, 状态等
func (c *Client) DescribeNameList(request *DescribeNameListRequest) (response *DescribeNameListResponse, err error) {
    return c.DescribeNameListWithContext(context.Background(), request)
}

// DescribeNameList
// 列表展示黑白名单列表数据, 包含列表名称, 名单类型, 数据类型, 数据来源, 描述, 状态等
func (c *Client) DescribeNameListWithContext(ctx context.Context, request *DescribeNameListRequest) (response *DescribeNameListResponse, err error) {
    if request == nil {
        request = NewDescribeNameListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNameList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNameListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNameListDataListRequest() (request *DescribeNameListDataListRequest) {
    request = &DescribeNameListDataListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rce", APIVersion, "DescribeNameListDataList")
    
    
    return
}

func NewDescribeNameListDataListResponse() (response *DescribeNameListDataListResponse) {
    response = &DescribeNameListDataListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNameListDataList
// 黑白名单详情数据展示 名单id 客户appid uin 数据内容 开始时间和结束时间 状态 描述
func (c *Client) DescribeNameListDataList(request *DescribeNameListDataListRequest) (response *DescribeNameListDataListResponse, err error) {
    return c.DescribeNameListDataListWithContext(context.Background(), request)
}

// DescribeNameListDataList
// 黑白名单详情数据展示 名单id 客户appid uin 数据内容 开始时间和结束时间 状态 描述
func (c *Client) DescribeNameListDataListWithContext(ctx context.Context, request *DescribeNameListDataListRequest) (response *DescribeNameListDataListResponse, err error) {
    if request == nil {
        request = NewDescribeNameListDataListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNameListDataList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNameListDataListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNameListDetailRequest() (request *DescribeNameListDetailRequest) {
    request = &DescribeNameListDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rce", APIVersion, "DescribeNameListDetail")
    
    
    return
}

func NewDescribeNameListDetailResponse() (response *DescribeNameListDetailResponse) {
    response = &DescribeNameListDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNameListDetail
// 查询黑白名单列表详情
func (c *Client) DescribeNameListDetail(request *DescribeNameListDetailRequest) (response *DescribeNameListDetailResponse, err error) {
    return c.DescribeNameListDetailWithContext(context.Background(), request)
}

// DescribeNameListDetail
// 查询黑白名单列表详情
func (c *Client) DescribeNameListDetailWithContext(ctx context.Context, request *DescribeNameListDetailRequest) (response *DescribeNameListDetailResponse, err error) {
    if request == nil {
        request = NewDescribeNameListDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNameListDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNameListDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserUsageCntRequest() (request *DescribeUserUsageCntRequest) {
    request = &DescribeUserUsageCntRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rce", APIVersion, "DescribeUserUsageCnt")
    
    
    return
}

func NewDescribeUserUsageCntResponse() (response *DescribeUserUsageCntResponse) {
    response = &DescribeUserUsageCntResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserUsageCnt
// RCE控制台预付费和后付费次数展示
func (c *Client) DescribeUserUsageCnt(request *DescribeUserUsageCntRequest) (response *DescribeUserUsageCntResponse, err error) {
    return c.DescribeUserUsageCntWithContext(context.Background(), request)
}

// DescribeUserUsageCnt
// RCE控制台预付费和后付费次数展示
func (c *Client) DescribeUserUsageCntWithContext(ctx context.Context, request *DescribeUserUsageCntRequest) (response *DescribeUserUsageCntResponse, err error) {
    if request == nil {
        request = NewDescribeUserUsageCntRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserUsageCnt require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserUsageCntResponse()
    err = c.Send(request, response)
    return
}

func NewImportNameListDataRequest() (request *ImportNameListDataRequest) {
    request = &ImportNameListDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rce", APIVersion, "ImportNameListData")
    
    
    return
}

func NewImportNameListDataResponse() (response *ImportNameListDataResponse) {
    response = &ImportNameListDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ImportNameListData
// 新增黑白名单数据，所有黑白名单数据总量上限为10000
func (c *Client) ImportNameListData(request *ImportNameListDataRequest) (response *ImportNameListDataResponse, err error) {
    return c.ImportNameListDataWithContext(context.Background(), request)
}

// ImportNameListData
// 新增黑白名单数据，所有黑白名单数据总量上限为10000
func (c *Client) ImportNameListDataWithContext(ctx context.Context, request *ImportNameListDataRequest) (response *ImportNameListDataResponse, err error) {
    if request == nil {
        request = NewImportNameListDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ImportNameListData require credential")
    }

    request.SetContext(ctx)
    
    response = NewImportNameListDataResponse()
    err = c.Send(request, response)
    return
}

func NewManageMarketingRiskRequest() (request *ManageMarketingRiskRequest) {
    request = &ManageMarketingRiskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rce", APIVersion, "ManageMarketingRisk")
    
    
    return
}

func NewManageMarketingRiskResponse() (response *ManageMarketingRiskResponse) {
    response = &ManageMarketingRiskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ManageMarketingRisk
// 全栈式风控引擎（RiskControlEngine，RCE）是基于人工智能技术和腾讯20年风控实战沉淀，依托腾讯海量业务构建的风控引擎，以轻量级的 SaaS 服务方式接入，帮助您快速解决注册、登录、营销活动等关键场景遇到的欺诈问题，实时防御黑灰产作恶。
func (c *Client) ManageMarketingRisk(request *ManageMarketingRiskRequest) (response *ManageMarketingRiskResponse, err error) {
    return c.ManageMarketingRiskWithContext(context.Background(), request)
}

// ManageMarketingRisk
// 全栈式风控引擎（RiskControlEngine，RCE）是基于人工智能技术和腾讯20年风控实战沉淀，依托腾讯海量业务构建的风控引擎，以轻量级的 SaaS 服务方式接入，帮助您快速解决注册、登录、营销活动等关键场景遇到的欺诈问题，实时防御黑灰产作恶。
func (c *Client) ManageMarketingRiskWithContext(ctx context.Context, request *ManageMarketingRiskRequest) (response *ManageMarketingRiskResponse, err error) {
    if request == nil {
        request = NewManageMarketingRiskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ManageMarketingRisk require credential")
    }

    request.SetContext(ctx)
    
    response = NewManageMarketingRiskResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNameListRequest() (request *ModifyNameListRequest) {
    request = &ModifyNameListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rce", APIVersion, "ModifyNameList")
    
    
    return
}

func NewModifyNameListResponse() (response *ModifyNameListResponse) {
    response = &ModifyNameListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyNameList
// 修改列表数据 列表名称 列表类型 数据类型 状态 备注
func (c *Client) ModifyNameList(request *ModifyNameListRequest) (response *ModifyNameListResponse, err error) {
    return c.ModifyNameListWithContext(context.Background(), request)
}

// ModifyNameList
// 修改列表数据 列表名称 列表类型 数据类型 状态 备注
func (c *Client) ModifyNameListWithContext(ctx context.Context, request *ModifyNameListRequest) (response *ModifyNameListResponse, err error) {
    if request == nil {
        request = NewModifyNameListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNameList require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNameListResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNameListDataRequest() (request *ModifyNameListDataRequest) {
    request = &ModifyNameListDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rce", APIVersion, "ModifyNameListData")
    
    
    return
}

func NewModifyNameListDataResponse() (response *ModifyNameListDataResponse) {
    response = &ModifyNameListDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyNameListData
// 修改黑白名单列表详情 详情内容 开始和结束时间 状态 备注等
func (c *Client) ModifyNameListData(request *ModifyNameListDataRequest) (response *ModifyNameListDataResponse, err error) {
    return c.ModifyNameListDataWithContext(context.Background(), request)
}

// ModifyNameListData
// 修改黑白名单列表详情 详情内容 开始和结束时间 状态 备注等
func (c *Client) ModifyNameListDataWithContext(ctx context.Context, request *ModifyNameListDataRequest) (response *ModifyNameListDataResponse, err error) {
    if request == nil {
        request = NewModifyNameListDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNameListData require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNameListDataResponse()
    err = c.Send(request, response)
    return
}
