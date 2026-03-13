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

package v20180608

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-06-08"

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


func NewBindCloudBaseAccessDomainRequest() (request *BindCloudBaseAccessDomainRequest) {
    request = &BindCloudBaseAccessDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "BindCloudBaseAccessDomain")
    
    
    return
}

func NewBindCloudBaseAccessDomainResponse() (response *BindCloudBaseAccessDomainResponse) {
    response = &BindCloudBaseAccessDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindCloudBaseAccessDomain
// 绑定云开发自定义域名，用于云接入和静态托管
//
// 可能返回的错误码:
//  INVALIDPARAMETER_CNAMENOTMATCH = "InvalidParameter.CNAMENotMatch"
//  INVALIDPARAMETER_DOMAINEXIST = "InvalidParameter.DomainExist"
//  INVALIDPARAMETER_SERVICEEVIL = "InvalidParameter.ServiceEvil"
//  INVALIDPARAMETER_SERVICEICP = "InvalidParameter.ServiceICP"
//  INVALIDPARAMETER_SERVICETHRESHOLD = "InvalidParameter.ServiceThreshold"
//  RESOURCEUNAVAILABLE_CDNFREEZED = "ResourceUnavailable.CDNFreezed"
func (c *Client) BindCloudBaseAccessDomain(request *BindCloudBaseAccessDomainRequest) (response *BindCloudBaseAccessDomainResponse, err error) {
    return c.BindCloudBaseAccessDomainWithContext(context.Background(), request)
}

// BindCloudBaseAccessDomain
// 绑定云开发自定义域名，用于云接入和静态托管
//
// 可能返回的错误码:
//  INVALIDPARAMETER_CNAMENOTMATCH = "InvalidParameter.CNAMENotMatch"
//  INVALIDPARAMETER_DOMAINEXIST = "InvalidParameter.DomainExist"
//  INVALIDPARAMETER_SERVICEEVIL = "InvalidParameter.ServiceEvil"
//  INVALIDPARAMETER_SERVICEICP = "InvalidParameter.ServiceICP"
//  INVALIDPARAMETER_SERVICETHRESHOLD = "InvalidParameter.ServiceThreshold"
//  RESOURCEUNAVAILABLE_CDNFREEZED = "ResourceUnavailable.CDNFreezed"
func (c *Client) BindCloudBaseAccessDomainWithContext(ctx context.Context, request *BindCloudBaseAccessDomainRequest) (response *BindCloudBaseAccessDomainResponse, err error) {
    if request == nil {
        request = NewBindCloudBaseAccessDomainRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "BindCloudBaseAccessDomain")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindCloudBaseAccessDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindCloudBaseAccessDomainResponse()
    err = c.Send(request, response)
    return
}

func NewBindCloudBaseGWDomainRequest() (request *BindCloudBaseGWDomainRequest) {
    request = &BindCloudBaseGWDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "BindCloudBaseGWDomain")
    
    
    return
}

func NewBindCloudBaseGWDomainResponse() (response *BindCloudBaseGWDomainResponse) {
    response = &BindCloudBaseGWDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindCloudBaseGWDomain
// 绑定自定义域名
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEMFAIL = "InternalError.SystemFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APINOEXIST = "InvalidParameter.APINoExist"
//  INVALIDPARAMETER_CNAMENOTMATCH = "InvalidParameter.CNAMENotMatch"
//  INVALIDPARAMETER_DOMAINEXIST = "InvalidParameter.DomainExist"
//  INVALIDPARAMETER_DOMAINNOTEXIST = "InvalidParameter.DomainNotExist"
//  INVALIDPARAMETER_EXCLUSIVECERT = "InvalidParameter.ExclusiveCert"
//  INVALIDPARAMETER_SERVICEEVIL = "InvalidParameter.ServiceEvil"
//  INVALIDPARAMETER_SERVICEICP = "InvalidParameter.ServiceICP"
//  INVALIDPARAMETER_SERVICETHRESHOLD = "InvalidParameter.ServiceThreshold"
func (c *Client) BindCloudBaseGWDomain(request *BindCloudBaseGWDomainRequest) (response *BindCloudBaseGWDomainResponse, err error) {
    return c.BindCloudBaseGWDomainWithContext(context.Background(), request)
}

// BindCloudBaseGWDomain
// 绑定自定义域名
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEMFAIL = "InternalError.SystemFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APINOEXIST = "InvalidParameter.APINoExist"
//  INVALIDPARAMETER_CNAMENOTMATCH = "InvalidParameter.CNAMENotMatch"
//  INVALIDPARAMETER_DOMAINEXIST = "InvalidParameter.DomainExist"
//  INVALIDPARAMETER_DOMAINNOTEXIST = "InvalidParameter.DomainNotExist"
//  INVALIDPARAMETER_EXCLUSIVECERT = "InvalidParameter.ExclusiveCert"
//  INVALIDPARAMETER_SERVICEEVIL = "InvalidParameter.ServiceEvil"
//  INVALIDPARAMETER_SERVICEICP = "InvalidParameter.ServiceICP"
//  INVALIDPARAMETER_SERVICETHRESHOLD = "InvalidParameter.ServiceThreshold"
func (c *Client) BindCloudBaseGWDomainWithContext(ctx context.Context, request *BindCloudBaseGWDomainRequest) (response *BindCloudBaseGWDomainResponse, err error) {
    if request == nil {
        request = NewBindCloudBaseGWDomainRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "BindCloudBaseGWDomain")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindCloudBaseGWDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindCloudBaseGWDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCheckTcbServiceRequest() (request *CheckTcbServiceRequest) {
    request = &CheckTcbServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "CheckTcbService")
    
    
    return
}

func NewCheckTcbServiceResponse() (response *CheckTcbServiceResponse) {
    response = &CheckTcbServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckTcbService
// 检查是否开通Tcb服务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) CheckTcbService(request *CheckTcbServiceRequest) (response *CheckTcbServiceResponse, err error) {
    return c.CheckTcbServiceWithContext(context.Background(), request)
}

// CheckTcbService
// 检查是否开通Tcb服务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) CheckTcbServiceWithContext(ctx context.Context, request *CheckTcbServiceRequest) (response *CheckTcbServiceResponse, err error) {
    if request == nil {
        request = NewCheckTcbServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "CheckTcbService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckTcbService require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckTcbServiceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAuthDomainRequest() (request *CreateAuthDomainRequest) {
    request = &CreateAuthDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "CreateAuthDomain")
    
    
    return
}

func NewCreateAuthDomainResponse() (response *CreateAuthDomainResponse) {
    response = &CreateAuthDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAuthDomain
// 增加安全域名。
//
// 云开发会校验网页应用请求的来源域名，您需要将来源域名加入到WEB安全域名列表中。
//
// 可以通过接口 [DescribeAuthDomains](https://cloud.tencent.com/document/product/876/42151) 获取当前已绑定生效的安全域名。
//
// 
//
// 注意⚠️
//
//   安全域名绑定成功之后，需要几分钟时间逐步生效。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  OPERATIONDENIED_FREEPACKAGEDENIED = "OperationDenied.FreePackageDenied"
//  OPERATIONDENIED_RESOURCEFROZEN = "OperationDenied.ResourceFrozen"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateAuthDomain(request *CreateAuthDomainRequest) (response *CreateAuthDomainResponse, err error) {
    return c.CreateAuthDomainWithContext(context.Background(), request)
}

// CreateAuthDomain
// 增加安全域名。
//
// 云开发会校验网页应用请求的来源域名，您需要将来源域名加入到WEB安全域名列表中。
//
// 可以通过接口 [DescribeAuthDomains](https://cloud.tencent.com/document/product/876/42151) 获取当前已绑定生效的安全域名。
//
// 
//
// 注意⚠️
//
//   安全域名绑定成功之后，需要几分钟时间逐步生效。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  OPERATIONDENIED_FREEPACKAGEDENIED = "OperationDenied.FreePackageDenied"
//  OPERATIONDENIED_RESOURCEFROZEN = "OperationDenied.ResourceFrozen"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateAuthDomainWithContext(ctx context.Context, request *CreateAuthDomainRequest) (response *CreateAuthDomainResponse, err error) {
    if request == nil {
        request = NewCreateAuthDomainRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "CreateAuthDomain")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAuthDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAuthDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBillDealRequest() (request *CreateBillDealRequest) {
    request = &CreateBillDealRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "CreateBillDeal")
    
    
    return
}

func NewCreateBillDealResponse() (response *CreateBillDealResponse) {
    response = &CreateBillDealResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateBillDeal
// 创建云开发产品计费订单，用于以下几种场景：
//
// 1. 购买云开发环境
//
// 2. 续费云开发环境
//
// 3. 变更云开发环境套餐
//
// 4. 购买云开发资源包
//
// 5. 购买云开发大促包
//
// 
//
// 该接口支持下单并支付(CreateAndPay=true时)，此时会自动在腾讯云账户中扣除余额（余额不足会下单失败）。
//
// 该接口支持自动扣除代金券（AutoVoucher=true时），符合条件的代金券会被自动扣除。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  OPERATIONDENIED_FREEPACKAGEDENIED = "OperationDenied.FreePackageDenied"
//  OPERATIONDENIED_RESOURCEFROZEN = "OperationDenied.ResourceFrozen"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateBillDeal(request *CreateBillDealRequest) (response *CreateBillDealResponse, err error) {
    return c.CreateBillDealWithContext(context.Background(), request)
}

// CreateBillDeal
// 创建云开发产品计费订单，用于以下几种场景：
//
// 1. 购买云开发环境
//
// 2. 续费云开发环境
//
// 3. 变更云开发环境套餐
//
// 4. 购买云开发资源包
//
// 5. 购买云开发大促包
//
// 
//
// 该接口支持下单并支付(CreateAndPay=true时)，此时会自动在腾讯云账户中扣除余额（余额不足会下单失败）。
//
// 该接口支持自动扣除代金券（AutoVoucher=true时），符合条件的代金券会被自动扣除。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  OPERATIONDENIED_FREEPACKAGEDENIED = "OperationDenied.FreePackageDenied"
//  OPERATIONDENIED_RESOURCEFROZEN = "OperationDenied.ResourceFrozen"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateBillDealWithContext(ctx context.Context, request *CreateBillDealRequest) (response *CreateBillDealResponse, err error) {
    if request == nil {
        request = NewCreateBillDealRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "CreateBillDeal")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBillDeal require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBillDealResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCloudBaseGWAPIRequest() (request *CreateCloudBaseGWAPIRequest) {
    request = &CreateCloudBaseGWAPIRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "CreateCloudBaseGWAPI")
    
    
    return
}

func NewCreateCloudBaseGWAPIResponse() (response *CreateCloudBaseGWAPIResponse) {
    response = &CreateCloudBaseGWAPIResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCloudBaseGWAPI
// 创建云开发网关API
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEMFAIL = "InternalError.SystemFail"
//  INVALIDPARAMETER_APICREATED = "InvalidParameter.APICreated"
//  INVALIDPARAMETER_APINOTEXIST = "InvalidParameter.APINotExist"
//  INVALIDPARAMETER_APITHRESHOLD = "InvalidParameter.APIThreshold"
//  INVALIDPARAMETER_APITYPENOTSUPPORT = "InvalidParameter.APITypeNotSupport"
//  INVALIDPARAMETER_PATHEXIST = "InvalidParameter.PathExist"
//  INVALIDPARAMETER_SERVICEEVIL = "InvalidParameter.ServiceEvil"
//  RESOURCEUNAVAILABLE_CDNFREEZED = "ResourceUnavailable.CDNFreezed"
func (c *Client) CreateCloudBaseGWAPI(request *CreateCloudBaseGWAPIRequest) (response *CreateCloudBaseGWAPIResponse, err error) {
    return c.CreateCloudBaseGWAPIWithContext(context.Background(), request)
}

// CreateCloudBaseGWAPI
// 创建云开发网关API
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEMFAIL = "InternalError.SystemFail"
//  INVALIDPARAMETER_APICREATED = "InvalidParameter.APICreated"
//  INVALIDPARAMETER_APINOTEXIST = "InvalidParameter.APINotExist"
//  INVALIDPARAMETER_APITHRESHOLD = "InvalidParameter.APIThreshold"
//  INVALIDPARAMETER_APITYPENOTSUPPORT = "InvalidParameter.APITypeNotSupport"
//  INVALIDPARAMETER_PATHEXIST = "InvalidParameter.PathExist"
//  INVALIDPARAMETER_SERVICEEVIL = "InvalidParameter.ServiceEvil"
//  RESOURCEUNAVAILABLE_CDNFREEZED = "ResourceUnavailable.CDNFreezed"
func (c *Client) CreateCloudBaseGWAPIWithContext(ctx context.Context, request *CreateCloudBaseGWAPIRequest) (response *CreateCloudBaseGWAPIResponse, err error) {
    if request == nil {
        request = NewCreateCloudBaseGWAPIRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "CreateCloudBaseGWAPI")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudBaseGWAPI require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudBaseGWAPIResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEnvRequest() (request *CreateEnvRequest) {
    request = &CreateEnvRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "CreateEnv")
    
    
    return
}

func NewCreateEnvResponse() (response *CreateEnvResponse) {
    response = &CreateEnvResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateEnv
// 本接口用于购买云开发环境。
//
// 该接口会自动下单并支付，会在腾讯云账户中扣除余额（余额不足会下单失败）。
//
// 该接口支持自动扣除代金券（AutoVoucher=true时），符合条件的代金券会被自动扣除。
//
// 环境下单成功之后会返回EnvId。EnvId是全局唯一表示。
//
// 环境发货是异步行为，后续可以通过接口 [DescribeEnvs ](https://cloud.tencent.com/document/product/876/34820) 查询环境状态和各项资源信息；通过 [DescribeBillingInfo](https://cloud.tencent.com/document/product/876/94390) 查询环境套餐信息，包括 到期时间、当前套餐等。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) CreateEnv(request *CreateEnvRequest) (response *CreateEnvResponse, err error) {
    return c.CreateEnvWithContext(context.Background(), request)
}

// CreateEnv
// 本接口用于购买云开发环境。
//
// 该接口会自动下单并支付，会在腾讯云账户中扣除余额（余额不足会下单失败）。
//
// 该接口支持自动扣除代金券（AutoVoucher=true时），符合条件的代金券会被自动扣除。
//
// 环境下单成功之后会返回EnvId。EnvId是全局唯一表示。
//
// 环境发货是异步行为，后续可以通过接口 [DescribeEnvs ](https://cloud.tencent.com/document/product/876/34820) 查询环境状态和各项资源信息；通过 [DescribeBillingInfo](https://cloud.tencent.com/document/product/876/94390) 查询环境套餐信息，包括 到期时间、当前套餐等。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) CreateEnvWithContext(ctx context.Context, request *CreateEnvRequest) (response *CreateEnvResponse, err error) {
    if request == nil {
        request = NewCreateEnvRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "CreateEnv")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEnv require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEnvResponse()
    err = c.Send(request, response)
    return
}

func NewCreateHostingDomainRequest() (request *CreateHostingDomainRequest) {
    request = &CreateHostingDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "CreateHostingDomain")
    
    
    return
}

func NewCreateHostingDomainResponse() (response *CreateHostingDomainResponse) {
    response = &CreateHostingDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateHostingDomain
// 创建托管域名
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  OPERATIONDENIED_RESOURCEFROZEN = "OperationDenied.ResourceFrozen"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateHostingDomain(request *CreateHostingDomainRequest) (response *CreateHostingDomainResponse, err error) {
    return c.CreateHostingDomainWithContext(context.Background(), request)
}

// CreateHostingDomain
// 创建托管域名
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  OPERATIONDENIED_RESOURCEFROZEN = "OperationDenied.ResourceFrozen"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateHostingDomainWithContext(ctx context.Context, request *CreateHostingDomainRequest) (response *CreateHostingDomainResponse, err error) {
    if request == nil {
        request = NewCreateHostingDomainRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "CreateHostingDomain")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateHostingDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateHostingDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMySQLRequest() (request *CreateMySQLRequest) {
    request = &CreateMySQLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "CreateMySQL")
    
    
    return
}

func NewCreateMySQLResponse() (response *CreateMySQLResponse) {
    response = &CreateMySQLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateMySQL
// 开通Mysql型数据库
//
// 
//
// 开通后，可通过 [DescribeCreateMySQLResult ](https://cloud.tencent.com/document/api/876/128185) 查询开通结果，Mysql开通成功后，可通过接口设置数据库账号相关功能包括但不限于【创建账号、删除账号、查询可授权权限列表、查询账号已有权限、修改主机、修改配置、修改账号库表权限】、集群操作相关【查询集群参数、修改集群参数】，连接设置相关【关闭外网、开通外网、查询集群信息】，备份回档相关【创建手动回档、删除手动回档、修改自动备份配置信息、查询备份文件列表、集群回档、查询任务列表、获取table列表、获取集群数据库列表、查询备份下载地址】，相关功能接口文档：[TDSQL-C MySQL API文档](https://cloud.tencent.com/document/product/1003/48106)，可以通过 [RunSql](https://cloud.tencent.com/document/api/876/127880) 接口来执行 sql 命令，比如创建表格、插入数据、删除表格等 sql 命令
//
// 可能返回的错误码:
//  INTERNALERROR_SYS_ERR = "InternalError.SYS_ERR"
//  INVALIDPARAMETER_INVALID_PARAM = "InvalidParameter.INVALID_PARAM"
func (c *Client) CreateMySQL(request *CreateMySQLRequest) (response *CreateMySQLResponse, err error) {
    return c.CreateMySQLWithContext(context.Background(), request)
}

// CreateMySQL
// 开通Mysql型数据库
//
// 
//
// 开通后，可通过 [DescribeCreateMySQLResult ](https://cloud.tencent.com/document/api/876/128185) 查询开通结果，Mysql开通成功后，可通过接口设置数据库账号相关功能包括但不限于【创建账号、删除账号、查询可授权权限列表、查询账号已有权限、修改主机、修改配置、修改账号库表权限】、集群操作相关【查询集群参数、修改集群参数】，连接设置相关【关闭外网、开通外网、查询集群信息】，备份回档相关【创建手动回档、删除手动回档、修改自动备份配置信息、查询备份文件列表、集群回档、查询任务列表、获取table列表、获取集群数据库列表、查询备份下载地址】，相关功能接口文档：[TDSQL-C MySQL API文档](https://cloud.tencent.com/document/product/1003/48106)，可以通过 [RunSql](https://cloud.tencent.com/document/api/876/127880) 接口来执行 sql 命令，比如创建表格、插入数据、删除表格等 sql 命令
//
// 可能返回的错误码:
//  INTERNALERROR_SYS_ERR = "InternalError.SYS_ERR"
//  INVALIDPARAMETER_INVALID_PARAM = "InvalidParameter.INVALID_PARAM"
func (c *Client) CreateMySQLWithContext(ctx context.Context, request *CreateMySQLRequest) (response *CreateMySQLResponse, err error) {
    if request == nil {
        request = NewCreateMySQLRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "CreateMySQL")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMySQL require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMySQLResponse()
    err = c.Send(request, response)
    return
}

func NewCreateStaticStoreRequest() (request *CreateStaticStoreRequest) {
    request = &CreateStaticStoreRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "CreateStaticStore")
    
    
    return
}

func NewCreateStaticStoreResponse() (response *CreateStaticStoreResponse) {
    response = &CreateStaticStoreResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateStaticStore
// 创建静态托管资源，包括COS和CDN，异步任务创建，查看创建结果需要根据DescribeStaticStore接口来查看
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateStaticStore(request *CreateStaticStoreRequest) (response *CreateStaticStoreResponse, err error) {
    return c.CreateStaticStoreWithContext(context.Background(), request)
}

// CreateStaticStore
// 创建静态托管资源，包括COS和CDN，异步任务创建，查看创建结果需要根据DescribeStaticStore接口来查看
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateStaticStoreWithContext(ctx context.Context, request *CreateStaticStoreRequest) (response *CreateStaticStoreResponse, err error) {
    if request == nil {
        request = NewCreateStaticStoreRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "CreateStaticStore")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateStaticStore require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateStaticStoreResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTableRequest() (request *CreateTableRequest) {
    request = &CreateTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "CreateTable")
    
    
    return
}

func NewCreateTableResponse() (response *CreateTableResponse) {
    response = &CreateTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTable
// 本接口(CreateTable)用于创建文档型数据库表，支持创建capped类型集合，暂时不支持分片表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATETABLE = "FailedOperation.CreateTable"
//  FAILEDOPERATION_LISTTABLE = "FailedOperation.ListTable"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_NOVALIDCONNECTION = "LimitExceeded.NoValidConnection"
//  LIMITEXCEEDED_OUTOFTABLEQUOTA = "LimitExceeded.OutOfTableQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_CONNECTOR = "ResourceNotFound.Connector"
//  RESOURCENOTFOUND_TABLE = "ResourceNotFound.Table"
//  RESOURCEUNAVAILABLE_MONGOISOLATED = "ResourceUnavailable.MongoIsolated"
//  RESOURCEUNAVAILABLE_RESOURCEEXIST = "ResourceUnavailable.ResourceExist"
//  RESOURCEUNAVAILABLE_RESOURCEOVERDUE = "ResourceUnavailable.ResourceOverdue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateTable(request *CreateTableRequest) (response *CreateTableResponse, err error) {
    return c.CreateTableWithContext(context.Background(), request)
}

// CreateTable
// 本接口(CreateTable)用于创建文档型数据库表，支持创建capped类型集合，暂时不支持分片表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATETABLE = "FailedOperation.CreateTable"
//  FAILEDOPERATION_LISTTABLE = "FailedOperation.ListTable"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_NOVALIDCONNECTION = "LimitExceeded.NoValidConnection"
//  LIMITEXCEEDED_OUTOFTABLEQUOTA = "LimitExceeded.OutOfTableQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_CONNECTOR = "ResourceNotFound.Connector"
//  RESOURCENOTFOUND_TABLE = "ResourceNotFound.Table"
//  RESOURCEUNAVAILABLE_MONGOISOLATED = "ResourceUnavailable.MongoIsolated"
//  RESOURCEUNAVAILABLE_RESOURCEEXIST = "ResourceUnavailable.ResourceExist"
//  RESOURCEUNAVAILABLE_RESOURCEOVERDUE = "ResourceUnavailable.ResourceOverdue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateTableWithContext(ctx context.Context, request *CreateTableRequest) (response *CreateTableResponse, err error) {
    if request == nil {
        request = NewCreateTableRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "CreateTable")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTable require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTableResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUserRequest() (request *CreateUserRequest) {
    request = &CreateUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "CreateUser")
    
    
    return
}

func NewCreateUserResponse() (response *CreateUserResponse) {
    response = &CreateUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateUser
// 创建tcb用户
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DUPLICATEDDATA = "FailedOperation.DuplicatedData"
//  FAILEDOPERATION_FLEXDBRESOURCEOVERDUE = "FailedOperation.FlexdbResourceOverdue"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateUser(request *CreateUserRequest) (response *CreateUserResponse, err error) {
    return c.CreateUserWithContext(context.Background(), request)
}

// CreateUser
// 创建tcb用户
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DUPLICATEDDATA = "FailedOperation.DuplicatedData"
//  FAILEDOPERATION_FLEXDBRESOURCEOVERDUE = "FailedOperation.FlexdbResourceOverdue"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateUserWithContext(ctx context.Context, request *CreateUserRequest) (response *CreateUserResponse, err error) {
    if request == nil {
        request = NewCreateUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "CreateUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUserResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAuthDomainRequest() (request *DeleteAuthDomainRequest) {
    request = &DeleteAuthDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DeleteAuthDomain")
    
    
    return
}

func NewDeleteAuthDomainResponse() (response *DeleteAuthDomainResponse) {
    response = &DeleteAuthDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAuthDomain
// 删除合法域名。
//
// 云开发会校验网页应用请求的来源域名，您需要将来源域名加入到WEB安全域名列表中。
//
// 可以通过接口 [DescribeAuthDomains](https://cloud.tencent.com/document/product/876/42151) 获取当前已绑定生效的安全域名。
//
// 
//
// 注意⚠️
//
// 安全域名被删除之后，可能会引起跨域问题，请谨慎操作。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  OPERATIONDENIED_RESOURCEFROZEN = "OperationDenied.ResourceFrozen"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteAuthDomain(request *DeleteAuthDomainRequest) (response *DeleteAuthDomainResponse, err error) {
    return c.DeleteAuthDomainWithContext(context.Background(), request)
}

// DeleteAuthDomain
// 删除合法域名。
//
// 云开发会校验网页应用请求的来源域名，您需要将来源域名加入到WEB安全域名列表中。
//
// 可以通过接口 [DescribeAuthDomains](https://cloud.tencent.com/document/product/876/42151) 获取当前已绑定生效的安全域名。
//
// 
//
// 注意⚠️
//
// 安全域名被删除之后，可能会引起跨域问题，请谨慎操作。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  OPERATIONDENIED_RESOURCEFROZEN = "OperationDenied.ResourceFrozen"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteAuthDomainWithContext(ctx context.Context, request *DeleteAuthDomainRequest) (response *DeleteAuthDomainResponse, err error) {
    if request == nil {
        request = NewDeleteAuthDomainRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DeleteAuthDomain")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAuthDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAuthDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCloudBaseGWAPIRequest() (request *DeleteCloudBaseGWAPIRequest) {
    request = &DeleteCloudBaseGWAPIRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DeleteCloudBaseGWAPI")
    
    
    return
}

func NewDeleteCloudBaseGWAPIResponse() (response *DeleteCloudBaseGWAPIResponse) {
    response = &DeleteCloudBaseGWAPIResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCloudBaseGWAPI
// 删除网关API
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEMFAIL = "InternalError.SystemFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APITYPENOTSUPPORT = "InvalidParameter.APITypeNotSupport"
//  INVALIDPARAMETER_SERVICEEVIL = "InvalidParameter.ServiceEvil"
//  RESOURCEUNAVAILABLE_CDNFREEZED = "ResourceUnavailable.CDNFreezed"
func (c *Client) DeleteCloudBaseGWAPI(request *DeleteCloudBaseGWAPIRequest) (response *DeleteCloudBaseGWAPIResponse, err error) {
    return c.DeleteCloudBaseGWAPIWithContext(context.Background(), request)
}

// DeleteCloudBaseGWAPI
// 删除网关API
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEMFAIL = "InternalError.SystemFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APITYPENOTSUPPORT = "InvalidParameter.APITypeNotSupport"
//  INVALIDPARAMETER_SERVICEEVIL = "InvalidParameter.ServiceEvil"
//  RESOURCEUNAVAILABLE_CDNFREEZED = "ResourceUnavailable.CDNFreezed"
func (c *Client) DeleteCloudBaseGWAPIWithContext(ctx context.Context, request *DeleteCloudBaseGWAPIRequest) (response *DeleteCloudBaseGWAPIResponse, err error) {
    if request == nil {
        request = NewDeleteCloudBaseGWAPIRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DeleteCloudBaseGWAPI")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCloudBaseGWAPI require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCloudBaseGWAPIResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCloudBaseGWDomainRequest() (request *DeleteCloudBaseGWDomainRequest) {
    request = &DeleteCloudBaseGWDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DeleteCloudBaseGWDomain")
    
    
    return
}

func NewDeleteCloudBaseGWDomainResponse() (response *DeleteCloudBaseGWDomainResponse) {
    response = &DeleteCloudBaseGWDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCloudBaseGWDomain
// 删除网关域名
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEMFAIL = "InternalError.SystemFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINNOTEXIST = "InvalidParameter.DomainNotExist"
//  INVALIDPARAMETER_SERVICEEVIL = "InvalidParameter.ServiceEvil"
func (c *Client) DeleteCloudBaseGWDomain(request *DeleteCloudBaseGWDomainRequest) (response *DeleteCloudBaseGWDomainResponse, err error) {
    return c.DeleteCloudBaseGWDomainWithContext(context.Background(), request)
}

// DeleteCloudBaseGWDomain
// 删除网关域名
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEMFAIL = "InternalError.SystemFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINNOTEXIST = "InvalidParameter.DomainNotExist"
//  INVALIDPARAMETER_SERVICEEVIL = "InvalidParameter.ServiceEvil"
func (c *Client) DeleteCloudBaseGWDomainWithContext(ctx context.Context, request *DeleteCloudBaseGWDomainRequest) (response *DeleteCloudBaseGWDomainResponse, err error) {
    if request == nil {
        request = NewDeleteCloudBaseGWDomainRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DeleteCloudBaseGWDomain")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCloudBaseGWDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCloudBaseGWDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTableRequest() (request *DeleteTableRequest) {
    request = &DeleteTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DeleteTable")
    
    
    return
}

func NewDeleteTableResponse() (response *DeleteTableResponse) {
    response = &DeleteTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTable
// 本接口(DeleteTable)用于删除文档型数据库表，删除表后表中数据将会被删除且无法恢复，请谨慎操作
//
// 
//
// 接口入参中的 Tag 为文档型数据库的实例 Id，可以通过 [DescribeEnvs](https://cloud.tencent.com/document/api/876/34820) 接口返回的 EnvList[0].Databases[0].InstanceId 获取
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LISTTABLE = "FailedOperation.ListTable"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_NOVALIDCONNECTION = "LimitExceeded.NoValidConnection"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CONNECTOR = "ResourceNotFound.Connector"
//  RESOURCENOTFOUND_TABLE = "ResourceNotFound.Table"
//  RESOURCEUNAVAILABLE_MONGOISOLATED = "ResourceUnavailable.MongoIsolated"
//  RESOURCEUNAVAILABLE_RESOURCEOVERDUE = "ResourceUnavailable.ResourceOverdue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteTable(request *DeleteTableRequest) (response *DeleteTableResponse, err error) {
    return c.DeleteTableWithContext(context.Background(), request)
}

// DeleteTable
// 本接口(DeleteTable)用于删除文档型数据库表，删除表后表中数据将会被删除且无法恢复，请谨慎操作
//
// 
//
// 接口入参中的 Tag 为文档型数据库的实例 Id，可以通过 [DescribeEnvs](https://cloud.tencent.com/document/api/876/34820) 接口返回的 EnvList[0].Databases[0].InstanceId 获取
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LISTTABLE = "FailedOperation.ListTable"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_NOVALIDCONNECTION = "LimitExceeded.NoValidConnection"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CONNECTOR = "ResourceNotFound.Connector"
//  RESOURCENOTFOUND_TABLE = "ResourceNotFound.Table"
//  RESOURCEUNAVAILABLE_MONGOISOLATED = "ResourceUnavailable.MongoIsolated"
//  RESOURCEUNAVAILABLE_RESOURCEOVERDUE = "ResourceUnavailable.ResourceOverdue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteTableWithContext(ctx context.Context, request *DeleteTableRequest) (response *DeleteTableResponse, err error) {
    if request == nil {
        request = NewDeleteTableRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DeleteTable")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTable require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTableResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUsersRequest() (request *DeleteUsersRequest) {
    request = &DeleteUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DeleteUsers")
    
    
    return
}

func NewDeleteUsersResponse() (response *DeleteUsersResponse) {
    response = &DeleteUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteUsers
// 删除tcb用户
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteUsers(request *DeleteUsersRequest) (response *DeleteUsersResponse, err error) {
    return c.DeleteUsersWithContext(context.Background(), request)
}

// DeleteUsers
// 删除tcb用户
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteUsersWithContext(ctx context.Context, request *DeleteUsersRequest) (response *DeleteUsersResponse, err error) {
    if request == nil {
        request = NewDeleteUsersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DeleteUsers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUsers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUsersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuthDomainsRequest() (request *DescribeAuthDomainsRequest) {
    request = &DescribeAuthDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeAuthDomains")
    
    
    return
}

func NewDescribeAuthDomainsResponse() (response *DescribeAuthDomainsResponse) {
    response = &DescribeAuthDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAuthDomains
// 本接口用于获取当前环境的安全域名列表。
//
// 云开发会校验网页应用请求的来源域名，您需要将来源域名加入到WEB安全域名列表中。
//
// 可以通过接口 [CreateAuthDomain](https://cloud.tencent.com/document/product/876/42764) 增加安全域名。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) DescribeAuthDomains(request *DescribeAuthDomainsRequest) (response *DescribeAuthDomainsResponse, err error) {
    return c.DescribeAuthDomainsWithContext(context.Background(), request)
}

// DescribeAuthDomains
// 本接口用于获取当前环境的安全域名列表。
//
// 云开发会校验网页应用请求的来源域名，您需要将来源域名加入到WEB安全域名列表中。
//
// 可以通过接口 [CreateAuthDomain](https://cloud.tencent.com/document/product/876/42764) 增加安全域名。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) DescribeAuthDomainsWithContext(ctx context.Context, request *DescribeAuthDomainsRequest) (response *DescribeAuthDomainsResponse, err error) {
    if request == nil {
        request = NewDescribeAuthDomainsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeAuthDomains")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAuthDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAuthDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaasPackageListRequest() (request *DescribeBaasPackageListRequest) {
    request = &DescribeBaasPackageListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeBaasPackageList")
    
    
    return
}

func NewDescribeBaasPackageListResponse() (response *DescribeBaasPackageListResponse) {
    response = &DescribeBaasPackageListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBaasPackageList
// 获取新套餐列表，含详情，如果传了PackageId，则只获取指定套餐详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) DescribeBaasPackageList(request *DescribeBaasPackageListRequest) (response *DescribeBaasPackageListResponse, err error) {
    return c.DescribeBaasPackageListWithContext(context.Background(), request)
}

// DescribeBaasPackageList
// 获取新套餐列表，含详情，如果传了PackageId，则只获取指定套餐详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) DescribeBaasPackageListWithContext(ctx context.Context, request *DescribeBaasPackageListRequest) (response *DescribeBaasPackageListResponse, err error) {
    if request == nil {
        request = NewDescribeBaasPackageListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeBaasPackageList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBaasPackageList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBaasPackageListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudBaseBuildServiceRequest() (request *DescribeCloudBaseBuildServiceRequest) {
    request = &DescribeCloudBaseBuildServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeCloudBaseBuildService")
    
    
    return
}

func NewDescribeCloudBaseBuildServiceResponse() (response *DescribeCloudBaseBuildServiceResponse) {
    response = &DescribeCloudBaseBuildServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudBaseBuildService
// 获取云托管代码上传url
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SERVICENOTEXIST = "InvalidParameter.ServiceNotExist"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_VERSIONNOTFOUND = "ResourceNotFound.VersionNotFound"
//  RESOURCEUNAVAILABLE_RESOURCEBANNED = "ResourceUnavailable.ResourceBanned"
//  RESOURCEUNAVAILABLE_RESOURCEFROZEN = "ResourceUnavailable.ResourceFrozen"
//  RESOURCEUNAVAILABLE_RESOURCEISOLATED = "ResourceUnavailable.ResourceIsolated"
func (c *Client) DescribeCloudBaseBuildService(request *DescribeCloudBaseBuildServiceRequest) (response *DescribeCloudBaseBuildServiceResponse, err error) {
    return c.DescribeCloudBaseBuildServiceWithContext(context.Background(), request)
}

// DescribeCloudBaseBuildService
// 获取云托管代码上传url
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SERVICENOTEXIST = "InvalidParameter.ServiceNotExist"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_VERSIONNOTFOUND = "ResourceNotFound.VersionNotFound"
//  RESOURCEUNAVAILABLE_RESOURCEBANNED = "ResourceUnavailable.ResourceBanned"
//  RESOURCEUNAVAILABLE_RESOURCEFROZEN = "ResourceUnavailable.ResourceFrozen"
//  RESOURCEUNAVAILABLE_RESOURCEISOLATED = "ResourceUnavailable.ResourceIsolated"
func (c *Client) DescribeCloudBaseBuildServiceWithContext(ctx context.Context, request *DescribeCloudBaseBuildServiceRequest) (response *DescribeCloudBaseBuildServiceResponse, err error) {
    if request == nil {
        request = NewDescribeCloudBaseBuildServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeCloudBaseBuildService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudBaseBuildService require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudBaseBuildServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudBaseGWAPIRequest() (request *DescribeCloudBaseGWAPIRequest) {
    request = &DescribeCloudBaseGWAPIRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeCloudBaseGWAPI")
    
    
    return
}

func NewDescribeCloudBaseGWAPIResponse() (response *DescribeCloudBaseGWAPIResponse) {
    response = &DescribeCloudBaseGWAPIResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudBaseGWAPI
// 获取网关API列表
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEMFAIL = "InternalError.SystemFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APINOEXIST = "InvalidParameter.APINoExist"
//  INVALIDPARAMETER_APITYPENOTSUPPORT = "InvalidParameter.APITypeNotSupport"
//  INVALIDPARAMETER_EXCLUSIVECERT = "InvalidParameter.ExclusiveCert"
//  INVALIDPARAMETER_SERVICEEVIL = "InvalidParameter.ServiceEvil"
func (c *Client) DescribeCloudBaseGWAPI(request *DescribeCloudBaseGWAPIRequest) (response *DescribeCloudBaseGWAPIResponse, err error) {
    return c.DescribeCloudBaseGWAPIWithContext(context.Background(), request)
}

// DescribeCloudBaseGWAPI
// 获取网关API列表
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEMFAIL = "InternalError.SystemFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APINOEXIST = "InvalidParameter.APINoExist"
//  INVALIDPARAMETER_APITYPENOTSUPPORT = "InvalidParameter.APITypeNotSupport"
//  INVALIDPARAMETER_EXCLUSIVECERT = "InvalidParameter.ExclusiveCert"
//  INVALIDPARAMETER_SERVICEEVIL = "InvalidParameter.ServiceEvil"
func (c *Client) DescribeCloudBaseGWAPIWithContext(ctx context.Context, request *DescribeCloudBaseGWAPIRequest) (response *DescribeCloudBaseGWAPIResponse, err error) {
    if request == nil {
        request = NewDescribeCloudBaseGWAPIRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeCloudBaseGWAPI")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudBaseGWAPI require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudBaseGWAPIResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudBaseGWServiceRequest() (request *DescribeCloudBaseGWServiceRequest) {
    request = &DescribeCloudBaseGWServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeCloudBaseGWService")
    
    
    return
}

func NewDescribeCloudBaseGWServiceResponse() (response *DescribeCloudBaseGWServiceResponse) {
    response = &DescribeCloudBaseGWServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudBaseGWService
// 获取网关服务
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEMFAIL = "InternalError.SystemFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINNOTEXIST = "InvalidParameter.DomainNotExist"
//  INVALIDPARAMETER_SERVICEEVIL = "InvalidParameter.ServiceEvil"
//  INVALIDPARAMETER_SERVICEICP = "InvalidParameter.ServiceICP"
func (c *Client) DescribeCloudBaseGWService(request *DescribeCloudBaseGWServiceRequest) (response *DescribeCloudBaseGWServiceResponse, err error) {
    return c.DescribeCloudBaseGWServiceWithContext(context.Background(), request)
}

// DescribeCloudBaseGWService
// 获取网关服务
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEMFAIL = "InternalError.SystemFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINNOTEXIST = "InvalidParameter.DomainNotExist"
//  INVALIDPARAMETER_SERVICEEVIL = "InvalidParameter.ServiceEvil"
//  INVALIDPARAMETER_SERVICEICP = "InvalidParameter.ServiceICP"
func (c *Client) DescribeCloudBaseGWServiceWithContext(ctx context.Context, request *DescribeCloudBaseGWServiceRequest) (response *DescribeCloudBaseGWServiceResponse, err error) {
    if request == nil {
        request = NewDescribeCloudBaseGWServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeCloudBaseGWService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudBaseGWService require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudBaseGWServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCreateMySQLResultRequest() (request *DescribeCreateMySQLResultRequest) {
    request = &DescribeCreateMySQLResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeCreateMySQLResult")
    
    
    return
}

func NewDescribeCreateMySQLResultResponse() (response *DescribeCreateMySQLResultResponse) {
    response = &DescribeCreateMySQLResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCreateMySQLResult
// 查询开通Mysql结果，`Response.Data.Status = "notexist"` 表示未开通，如果未开通，可以调用 [CreateMySQL](https://cloud.tencent.com/document/api/876/128186) 来开通
//
//  `Response.Data. Status = "success"` 表示开通成功，Mysql开通成功后，可通过接口设置数据库账号相关功能包括但不限于【创建账号、删除账号、查询可授权权限列表、查询账号已有权限、修改主机、修改配置、修改账号库表权限】、集群操作相关【查询集群参数、修改集群参数】，连接设置相关【关闭外网、开通外网、查询集群信息】，备份回档相关【创建手动回档、删除手动回档、修改自动备份配置信息、查询备份文件列表、集群回档、查询任务列表、获取table列表、获取集群数据库列表、查询备份下载地址】，相关功能接口文档：[TDSQL-C MySQL API文档](https://cloud.tencent.com/document/product/1003/48106)，可以通过 [RunSql](https://cloud.tencent.com/document/api/876/127880) 接口来执行 sql 命令，比如创建表格、插入数据、删除表格等 sql 命令
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEMFAIL = "InternalError.SystemFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINNOTEXIST = "InvalidParameter.DomainNotExist"
//  INVALIDPARAMETER_SERVICEEVIL = "InvalidParameter.ServiceEvil"
//  INVALIDPARAMETER_SERVICEICP = "InvalidParameter.ServiceICP"
func (c *Client) DescribeCreateMySQLResult(request *DescribeCreateMySQLResultRequest) (response *DescribeCreateMySQLResultResponse, err error) {
    return c.DescribeCreateMySQLResultWithContext(context.Background(), request)
}

// DescribeCreateMySQLResult
// 查询开通Mysql结果，`Response.Data.Status = "notexist"` 表示未开通，如果未开通，可以调用 [CreateMySQL](https://cloud.tencent.com/document/api/876/128186) 来开通
//
//  `Response.Data. Status = "success"` 表示开通成功，Mysql开通成功后，可通过接口设置数据库账号相关功能包括但不限于【创建账号、删除账号、查询可授权权限列表、查询账号已有权限、修改主机、修改配置、修改账号库表权限】、集群操作相关【查询集群参数、修改集群参数】，连接设置相关【关闭外网、开通外网、查询集群信息】，备份回档相关【创建手动回档、删除手动回档、修改自动备份配置信息、查询备份文件列表、集群回档、查询任务列表、获取table列表、获取集群数据库列表、查询备份下载地址】，相关功能接口文档：[TDSQL-C MySQL API文档](https://cloud.tencent.com/document/product/1003/48106)，可以通过 [RunSql](https://cloud.tencent.com/document/api/876/127880) 接口来执行 sql 命令，比如创建表格、插入数据、删除表格等 sql 命令
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEMFAIL = "InternalError.SystemFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINNOTEXIST = "InvalidParameter.DomainNotExist"
//  INVALIDPARAMETER_SERVICEEVIL = "InvalidParameter.ServiceEvil"
//  INVALIDPARAMETER_SERVICEICP = "InvalidParameter.ServiceICP"
func (c *Client) DescribeCreateMySQLResultWithContext(ctx context.Context, request *DescribeCreateMySQLResultRequest) (response *DescribeCreateMySQLResultResponse, err error) {
    if request == nil {
        request = NewDescribeCreateMySQLResultRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeCreateMySQLResult")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCreateMySQLResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCreateMySQLResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatabaseACLRequest() (request *DescribeDatabaseACLRequest) {
    request = &DescribeDatabaseACLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeDatabaseACL")
    
    
    return
}

func NewDescribeDatabaseACLResponse() (response *DescribeDatabaseACLResponse) {
    response = &DescribeDatabaseACLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDatabaseACL
// 获取文档型数据库权限
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) DescribeDatabaseACL(request *DescribeDatabaseACLRequest) (response *DescribeDatabaseACLResponse, err error) {
    return c.DescribeDatabaseACLWithContext(context.Background(), request)
}

// DescribeDatabaseACL
// 获取文档型数据库权限
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) DescribeDatabaseACLWithContext(ctx context.Context, request *DescribeDatabaseACLRequest) (response *DescribeDatabaseACLResponse, err error) {
    if request == nil {
        request = NewDescribeDatabaseACLRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeDatabaseACL")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDatabaseACL require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDatabaseACLResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEnvAccountCircleRequest() (request *DescribeEnvAccountCircleRequest) {
    request = &DescribeEnvAccountCircleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeEnvAccountCircle")
    
    
    return
}

func NewDescribeEnvAccountCircleResponse() (response *DescribeEnvAccountCircleResponse) {
    response = &DescribeEnvAccountCircleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEnvAccountCircle
// 查询环境计费周期。
//
// 云开发环境的资源点都是按月结算的，每个月都有一定的抵扣额度。
//
// 
//
// 例如：
//
//   某个环境在 2026-01-05 购买了3个月个人版(到期时间: 2026-04-05)，则他可以在以下3个周期内，分别享有40000资源点的额度：
//
//   1. 2026-01-05 ~ 2026-02-05 23:59:59
//
//   2. 2026-02-06 ~ 2026-03-05 23:59:59
//
//   3. 2026-03-06 ~ 2026-04-05 23:59:59
//
// 
//
// 本接口，用于获取环境当前属于哪个计费周期内。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCEUNAVAILABLE_RESOURCEEXPIRED = "ResourceUnavailable.ResourceExpired"
func (c *Client) DescribeEnvAccountCircle(request *DescribeEnvAccountCircleRequest) (response *DescribeEnvAccountCircleResponse, err error) {
    return c.DescribeEnvAccountCircleWithContext(context.Background(), request)
}

// DescribeEnvAccountCircle
// 查询环境计费周期。
//
// 云开发环境的资源点都是按月结算的，每个月都有一定的抵扣额度。
//
// 
//
// 例如：
//
//   某个环境在 2026-01-05 购买了3个月个人版(到期时间: 2026-04-05)，则他可以在以下3个周期内，分别享有40000资源点的额度：
//
//   1. 2026-01-05 ~ 2026-02-05 23:59:59
//
//   2. 2026-02-06 ~ 2026-03-05 23:59:59
//
//   3. 2026-03-06 ~ 2026-04-05 23:59:59
//
// 
//
// 本接口，用于获取环境当前属于哪个计费周期内。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCEUNAVAILABLE_RESOURCEEXPIRED = "ResourceUnavailable.ResourceExpired"
func (c *Client) DescribeEnvAccountCircleWithContext(ctx context.Context, request *DescribeEnvAccountCircleRequest) (response *DescribeEnvAccountCircleResponse, err error) {
    if request == nil {
        request = NewDescribeEnvAccountCircleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeEnvAccountCircle")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEnvAccountCircle require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEnvAccountCircleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEnvLimitRequest() (request *DescribeEnvLimitRequest) {
    request = &DescribeEnvLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeEnvLimit")
    
    
    return
}

func NewDescribeEnvLimitResponse() (response *DescribeEnvLimitResponse) {
    response = &DescribeEnvLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEnvLimit
// 查询环境个数上限
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) DescribeEnvLimit(request *DescribeEnvLimitRequest) (response *DescribeEnvLimitResponse, err error) {
    return c.DescribeEnvLimitWithContext(context.Background(), request)
}

// DescribeEnvLimit
// 查询环境个数上限
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) DescribeEnvLimitWithContext(ctx context.Context, request *DescribeEnvLimitRequest) (response *DescribeEnvLimitResponse, err error) {
    if request == nil {
        request = NewDescribeEnvLimitRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeEnvLimit")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEnvLimit require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEnvLimitResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEnvsRequest() (request *DescribeEnvsRequest) {
    request = &DescribeEnvsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeEnvs")
    
    
    return
}

func NewDescribeEnvsResponse() (response *DescribeEnvsResponse) {
    response = &DescribeEnvsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEnvs
// 获取环境列表，含环境下的各个资源信息。尤其是各资源的唯一标识，是请求各资源的关键参数
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACTION = "InvalidParameter.Action"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCENOTFOUND_USERNOTEXISTS = "ResourceNotFound.UserNotExists"
func (c *Client) DescribeEnvs(request *DescribeEnvsRequest) (response *DescribeEnvsResponse, err error) {
    return c.DescribeEnvsWithContext(context.Background(), request)
}

// DescribeEnvs
// 获取环境列表，含环境下的各个资源信息。尤其是各资源的唯一标识，是请求各资源的关键参数
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACTION = "InvalidParameter.Action"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCENOTFOUND_USERNOTEXISTS = "ResourceNotFound.UserNotExists"
func (c *Client) DescribeEnvsWithContext(ctx context.Context, request *DescribeEnvsRequest) (response *DescribeEnvsResponse, err error) {
    if request == nil {
        request = NewDescribeEnvsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeEnvs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEnvs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEnvsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostingDomainTaskRequest() (request *DescribeHostingDomainTaskRequest) {
    request = &DescribeHostingDomainTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeHostingDomainTask")
    
    
    return
}

func NewDescribeHostingDomainTaskResponse() (response *DescribeHostingDomainTaskResponse) {
    response = &DescribeHostingDomainTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHostingDomainTask
// 查询静态托管域名任务状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_CONCURRENT = "LimitExceeded.Concurrent"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeHostingDomainTask(request *DescribeHostingDomainTaskRequest) (response *DescribeHostingDomainTaskResponse, err error) {
    return c.DescribeHostingDomainTaskWithContext(context.Background(), request)
}

// DescribeHostingDomainTask
// 查询静态托管域名任务状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_CONCURRENT = "LimitExceeded.Concurrent"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeHostingDomainTaskWithContext(ctx context.Context, request *DescribeHostingDomainTaskRequest) (response *DescribeHostingDomainTaskResponse, err error) {
    if request == nil {
        request = NewDescribeHostingDomainTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeHostingDomainTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHostingDomainTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHostingDomainTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMySQLClusterDetailRequest() (request *DescribeMySQLClusterDetailRequest) {
    request = &DescribeMySQLClusterDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeMySQLClusterDetail")
    
    
    return
}

func NewDescribeMySQLClusterDetailResponse() (response *DescribeMySQLClusterDetailResponse) {
    response = &DescribeMySQLClusterDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMySQLClusterDetail
// 查询Mysql集群信息
//
// 
//
// 调用该接口前需要先查询Mysql是否开通，可通过 [DescribeCreateMySQLResult ](https://cloud.tencent.com/document/api/876/128185) 查询，只有已开通的才能查到集群信息，Mysql开通成功后，可通过接口设置数据库账号相关功能包括但不限于【创建账号、删除账号、查询可授权权限列表、查询账号已有权限、修改主机、修改配置、修改账号库表权限】、集群操作相关【查询集群参数、修改集群参数】，连接设置相关【关闭外网、开通外网、查询集群信息】，备份回档相关【创建手动回档、删除手动回档、修改自动备份配置信息、查询备份文件列表、集群回档、查询任务列表、获取table列表、获取集群数据库列表、查询备份下载地址】，相关功能接口文档：[TDSQL-C MySQL API文档](https://cloud.tencent.com/document/product/1003/48106)，可以通过 [RunSql](https://cloud.tencent.com/document/api/876/127880) 接口来执行 sql 命令，比如创建表格、插入数据、删除表格等 sql 命令
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATASOURCENOTEXIST = "FailedOperation.DataSourceNotExist"
func (c *Client) DescribeMySQLClusterDetail(request *DescribeMySQLClusterDetailRequest) (response *DescribeMySQLClusterDetailResponse, err error) {
    return c.DescribeMySQLClusterDetailWithContext(context.Background(), request)
}

// DescribeMySQLClusterDetail
// 查询Mysql集群信息
//
// 
//
// 调用该接口前需要先查询Mysql是否开通，可通过 [DescribeCreateMySQLResult ](https://cloud.tencent.com/document/api/876/128185) 查询，只有已开通的才能查到集群信息，Mysql开通成功后，可通过接口设置数据库账号相关功能包括但不限于【创建账号、删除账号、查询可授权权限列表、查询账号已有权限、修改主机、修改配置、修改账号库表权限】、集群操作相关【查询集群参数、修改集群参数】，连接设置相关【关闭外网、开通外网、查询集群信息】，备份回档相关【创建手动回档、删除手动回档、修改自动备份配置信息、查询备份文件列表、集群回档、查询任务列表、获取table列表、获取集群数据库列表、查询备份下载地址】，相关功能接口文档：[TDSQL-C MySQL API文档](https://cloud.tencent.com/document/product/1003/48106)，可以通过 [RunSql](https://cloud.tencent.com/document/api/876/127880) 接口来执行 sql 命令，比如创建表格、插入数据、删除表格等 sql 命令
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATASOURCENOTEXIST = "FailedOperation.DataSourceNotExist"
func (c *Client) DescribeMySQLClusterDetailWithContext(ctx context.Context, request *DescribeMySQLClusterDetailRequest) (response *DescribeMySQLClusterDetailResponse, err error) {
    if request == nil {
        request = NewDescribeMySQLClusterDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeMySQLClusterDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMySQLClusterDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMySQLClusterDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMySQLTaskStatusRequest() (request *DescribeMySQLTaskStatusRequest) {
    request = &DescribeMySQLTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeMySQLTaskStatus")
    
    
    return
}

func NewDescribeMySQLTaskStatusResponse() (response *DescribeMySQLTaskStatusResponse) {
    response = &DescribeMySQLTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMySQLTaskStatus
// 查询Mysql任务状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATASOURCENOTEXIST = "FailedOperation.DataSourceNotExist"
func (c *Client) DescribeMySQLTaskStatus(request *DescribeMySQLTaskStatusRequest) (response *DescribeMySQLTaskStatusResponse, err error) {
    return c.DescribeMySQLTaskStatusWithContext(context.Background(), request)
}

// DescribeMySQLTaskStatus
// 查询Mysql任务状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATASOURCENOTEXIST = "FailedOperation.DataSourceNotExist"
func (c *Client) DescribeMySQLTaskStatusWithContext(ctx context.Context, request *DescribeMySQLTaskStatusRequest) (response *DescribeMySQLTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeMySQLTaskStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeMySQLTaskStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMySQLTaskStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMySQLTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeQuotaDataRequest() (request *DescribeQuotaDataRequest) {
    request = &DescribeQuotaDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeQuotaData")
    
    
    return
}

func NewDescribeQuotaDataResponse() (response *DescribeQuotaDataResponse) {
    response = &DescribeQuotaDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeQuotaData
// 查询指定指标的配额使用量
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeQuotaData(request *DescribeQuotaDataRequest) (response *DescribeQuotaDataResponse, err error) {
    return c.DescribeQuotaDataWithContext(context.Background(), request)
}

// DescribeQuotaData
// 查询指定指标的配额使用量
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeQuotaDataWithContext(ctx context.Context, request *DescribeQuotaDataRequest) (response *DescribeQuotaDataResponse, err error) {
    if request == nil {
        request = NewDescribeQuotaDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeQuotaData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeQuotaData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeQuotaDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSafeRuleRequest() (request *DescribeSafeRuleRequest) {
    request = &DescribeSafeRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeSafeRule")
    
    
    return
}

func NewDescribeSafeRuleResponse() (response *DescribeSafeRuleResponse) {
    response = &DescribeSafeRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSafeRule
// 查询数据库安全规则。
//
// 安全规则，用于控制C端用户的访问权限。详见 [安全规则介绍](https://cloud.tencent.com/document/product/876/123478) 。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) DescribeSafeRule(request *DescribeSafeRuleRequest) (response *DescribeSafeRuleResponse, err error) {
    return c.DescribeSafeRuleWithContext(context.Background(), request)
}

// DescribeSafeRule
// 查询数据库安全规则。
//
// 安全规则，用于控制C端用户的访问权限。详见 [安全规则介绍](https://cloud.tencent.com/document/product/876/123478) 。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) DescribeSafeRuleWithContext(ctx context.Context, request *DescribeSafeRuleRequest) (response *DescribeSafeRuleResponse, err error) {
    if request == nil {
        request = NewDescribeSafeRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeSafeRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSafeRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSafeRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStaticStoreRequest() (request *DescribeStaticStoreRequest) {
    request = &DescribeStaticStoreRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeStaticStore")
    
    
    return
}

func NewDescribeStaticStoreResponse() (response *DescribeStaticStoreResponse) {
    response = &DescribeStaticStoreResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStaticStore
// 查看当前环境下静态托管资源信息，根据返回结果判断静态资源的状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeStaticStore(request *DescribeStaticStoreRequest) (response *DescribeStaticStoreResponse, err error) {
    return c.DescribeStaticStoreWithContext(context.Background(), request)
}

// DescribeStaticStore
// 查看当前环境下静态托管资源信息，根据返回结果判断静态资源的状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeStaticStoreWithContext(ctx context.Context, request *DescribeStaticStoreRequest) (response *DescribeStaticStoreResponse, err error) {
    if request == nil {
        request = NewDescribeStaticStoreRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeStaticStore")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStaticStore require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStaticStoreResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTableRequest() (request *DescribeTableRequest) {
    request = &DescribeTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeTable")
    
    
    return
}

func NewDescribeTableResponse() (response *DescribeTableResponse) {
    response = &DescribeTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTable
// 查询文档型数据库表的相关信息，包括索引等信息
//
// 
//
// 接口入参中的 Tag 为文档型数据库的实例 Id，可以通过 [DescribeEnvs](https://cloud.tencent.com/document/api/876/34820) 接口返回的 EnvList[0].Databases[0].InstanceId 获取
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LISTTABLE = "FailedOperation.ListTable"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_NOVALIDCONNECTION = "LimitExceeded.NoValidConnection"
//  LIMITEXCEEDED_OUTOFREADREQUESTQUOTA = "LimitExceeded.OutOfReadRequestQuota"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CONNECTOR = "ResourceNotFound.Connector"
//  RESOURCENOTFOUND_TABLE = "ResourceNotFound.Table"
//  RESOURCEUNAVAILABLE_MONGOISOLATED = "ResourceUnavailable.MongoIsolated"
//  RESOURCEUNAVAILABLE_RESOURCEOVERDUE = "ResourceUnavailable.ResourceOverdue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTable(request *DescribeTableRequest) (response *DescribeTableResponse, err error) {
    return c.DescribeTableWithContext(context.Background(), request)
}

// DescribeTable
// 查询文档型数据库表的相关信息，包括索引等信息
//
// 
//
// 接口入参中的 Tag 为文档型数据库的实例 Id，可以通过 [DescribeEnvs](https://cloud.tencent.com/document/api/876/34820) 接口返回的 EnvList[0].Databases[0].InstanceId 获取
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LISTTABLE = "FailedOperation.ListTable"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_NOVALIDCONNECTION = "LimitExceeded.NoValidConnection"
//  LIMITEXCEEDED_OUTOFREADREQUESTQUOTA = "LimitExceeded.OutOfReadRequestQuota"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CONNECTOR = "ResourceNotFound.Connector"
//  RESOURCENOTFOUND_TABLE = "ResourceNotFound.Table"
//  RESOURCEUNAVAILABLE_MONGOISOLATED = "ResourceUnavailable.MongoIsolated"
//  RESOURCEUNAVAILABLE_RESOURCEOVERDUE = "ResourceUnavailable.ResourceOverdue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTableWithContext(ctx context.Context, request *DescribeTableRequest) (response *DescribeTableResponse, err error) {
    if request == nil {
        request = NewDescribeTableRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeTable")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTable require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTableResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTablesRequest() (request *DescribeTablesRequest) {
    request = &DescribeTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeTables")
    
    
    return
}

func NewDescribeTablesResponse() (response *DescribeTablesResponse) {
    response = &DescribeTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTables
// 本接口(DescribeTables)用于查询文档型数据库所有表信息，包括表名、表中数据条数、表中数据量、索引个数及索引的大小等
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LISTTABLE = "FailedOperation.ListTable"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_NOVALIDCONNECTION = "LimitExceeded.NoValidConnection"
//  LIMITEXCEEDED_OUTOFREADREQUESTQUOTA = "LimitExceeded.OutOfReadRequestQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_CONNECTOR = "ResourceNotFound.Connector"
//  RESOURCENOTFOUND_TABLE = "ResourceNotFound.Table"
//  RESOURCEUNAVAILABLE_MONGOISOLATED = "ResourceUnavailable.MongoIsolated"
//  RESOURCEUNAVAILABLE_RESOURCEOVERDUE = "ResourceUnavailable.ResourceOverdue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTables(request *DescribeTablesRequest) (response *DescribeTablesResponse, err error) {
    return c.DescribeTablesWithContext(context.Background(), request)
}

// DescribeTables
// 本接口(DescribeTables)用于查询文档型数据库所有表信息，包括表名、表中数据条数、表中数据量、索引个数及索引的大小等
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LISTTABLE = "FailedOperation.ListTable"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_NOVALIDCONNECTION = "LimitExceeded.NoValidConnection"
//  LIMITEXCEEDED_OUTOFREADREQUESTQUOTA = "LimitExceeded.OutOfReadRequestQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_CONNECTOR = "ResourceNotFound.Connector"
//  RESOURCENOTFOUND_TABLE = "ResourceNotFound.Table"
//  RESOURCEUNAVAILABLE_MONGOISOLATED = "ResourceUnavailable.MongoIsolated"
//  RESOURCEUNAVAILABLE_RESOURCEOVERDUE = "ResourceUnavailable.ResourceOverdue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTablesWithContext(ctx context.Context, request *DescribeTablesRequest) (response *DescribeTablesResponse, err error) {
    if request == nil {
        request = NewDescribeTablesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeTables")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTables require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTablesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserListRequest() (request *DescribeUserListRequest) {
    request = &DescribeUserListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeUserList")
    
    
    return
}

func NewDescribeUserListResponse() (response *DescribeUserListResponse) {
    response = &DescribeUserListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserList
// 查询tcb用户列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FLEXDBRESOURCEOVERDUE = "FailedOperation.FlexdbResourceOverdue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeUserList(request *DescribeUserListRequest) (response *DescribeUserListResponse, err error) {
    return c.DescribeUserListWithContext(context.Background(), request)
}

// DescribeUserList
// 查询tcb用户列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FLEXDBRESOURCEOVERDUE = "FailedOperation.FlexdbResourceOverdue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeUserListWithContext(ctx context.Context, request *DescribeUserListRequest) (response *DescribeUserListResponse, err error) {
    if request == nil {
        request = NewDescribeUserListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeUserList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserListResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyEnvRequest() (request *DestroyEnvRequest) {
    request = &DestroyEnvRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DestroyEnv")
    
    
    return
}

func NewDestroyEnvResponse() (response *DestroyEnvResponse) {
    response = &DestroyEnvResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DestroyEnv
// 本接口用于销毁云开发环境。
//
// 云开发环境遵循腾讯云包年包月预付费产品生命周期，因此环境销毁需要分两步：
//
// 1. 资源退费。此时会根据当前环境剩余有效期，自动退还相关费用(代金券不退)。退款后，环境进入隔离期。
//
// 2. 环境删除。环境在进入隔离期后15天会自动删除。也可以通过本接口，指定 IsForce=true 来强制删除隔离期环境。
//
// 
//
// **注意**⚠️
//
//   1. 环境退费后进入隔离期，则所有资源均无法访问，控制台无法操作和管理。
//
//   2. 环境被彻底删除后，所有数据均无法找回。请谨慎操作。
//
// 
//
// 可以通过接口 [tcb:DescribeBillingInfo](https://cloud.tencent.com/document/product/876/94390) 查询环境计费状态。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_INVOICEAMOUNTLACK = "ResourceUnavailable.InvoiceAmountLack"
//  RESOURCEUNAVAILABLE_RESOURCEOVERDUE = "ResourceUnavailable.ResourceOverdue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DestroyEnv(request *DestroyEnvRequest) (response *DestroyEnvResponse, err error) {
    return c.DestroyEnvWithContext(context.Background(), request)
}

// DestroyEnv
// 本接口用于销毁云开发环境。
//
// 云开发环境遵循腾讯云包年包月预付费产品生命周期，因此环境销毁需要分两步：
//
// 1. 资源退费。此时会根据当前环境剩余有效期，自动退还相关费用(代金券不退)。退款后，环境进入隔离期。
//
// 2. 环境删除。环境在进入隔离期后15天会自动删除。也可以通过本接口，指定 IsForce=true 来强制删除隔离期环境。
//
// 
//
// **注意**⚠️
//
//   1. 环境退费后进入隔离期，则所有资源均无法访问，控制台无法操作和管理。
//
//   2. 环境被彻底删除后，所有数据均无法找回。请谨慎操作。
//
// 
//
// 可以通过接口 [tcb:DescribeBillingInfo](https://cloud.tencent.com/document/product/876/94390) 查询环境计费状态。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_INVOICEAMOUNTLACK = "ResourceUnavailable.InvoiceAmountLack"
//  RESOURCEUNAVAILABLE_RESOURCEOVERDUE = "ResourceUnavailable.ResourceOverdue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DestroyEnvWithContext(ctx context.Context, request *DestroyEnvRequest) (response *DestroyEnvResponse, err error) {
    if request == nil {
        request = NewDestroyEnvRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DestroyEnv")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DestroyEnv require credential")
    }

    request.SetContext(ctx)
    
    response = NewDestroyEnvResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyMySQLRequest() (request *DestroyMySQLRequest) {
    request = &DestroyMySQLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DestroyMySQL")
    
    
    return
}

func NewDestroyMySQLResponse() (response *DestroyMySQLResponse) {
    response = &DestroyMySQLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DestroyMySQL
// 销毁Mysql
//
// 
//
// 销毁后可以通过 [DescribeMySQLTaskStatus](https://cloud.tencent.com/document/api/876/128183) 接口查询销毁结果，如果 `Response.Data. Status = FAILED ` 表示销毁失败，可以重新调用销毁接口重试
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATASOURCENOTEXIST = "FailedOperation.DataSourceNotExist"
func (c *Client) DestroyMySQL(request *DestroyMySQLRequest) (response *DestroyMySQLResponse, err error) {
    return c.DestroyMySQLWithContext(context.Background(), request)
}

// DestroyMySQL
// 销毁Mysql
//
// 
//
// 销毁后可以通过 [DescribeMySQLTaskStatus](https://cloud.tencent.com/document/api/876/128183) 接口查询销毁结果，如果 `Response.Data. Status = FAILED ` 表示销毁失败，可以重新调用销毁接口重试
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATASOURCENOTEXIST = "FailedOperation.DataSourceNotExist"
func (c *Client) DestroyMySQLWithContext(ctx context.Context, request *DestroyMySQLRequest) (response *DestroyMySQLResponse, err error) {
    if request == nil {
        request = NewDestroyMySQLRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DestroyMySQL")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DestroyMySQL require credential")
    }

    request.SetContext(ctx)
    
    response = NewDestroyMySQLResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyStaticStoreRequest() (request *DestroyStaticStoreRequest) {
    request = &DestroyStaticStoreRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DestroyStaticStore")
    
    
    return
}

func NewDestroyStaticStoreResponse() (response *DestroyStaticStoreResponse) {
    response = &DestroyStaticStoreResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DestroyStaticStore
// 销毁静态托管资源，该接口创建异步销毁任务，资源最终状态可从DestroyStaticStore接口查看
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DestroyStaticStore(request *DestroyStaticStoreRequest) (response *DestroyStaticStoreResponse, err error) {
    return c.DestroyStaticStoreWithContext(context.Background(), request)
}

// DestroyStaticStore
// 销毁静态托管资源，该接口创建异步销毁任务，资源最终状态可从DestroyStaticStore接口查看
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DestroyStaticStoreWithContext(ctx context.Context, request *DestroyStaticStoreRequest) (response *DestroyStaticStoreResponse, err error) {
    if request == nil {
        request = NewDestroyStaticStoreRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DestroyStaticStore")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DestroyStaticStore require credential")
    }

    request.SetContext(ctx)
    
    response = NewDestroyStaticStoreResponse()
    err = c.Send(request, response)
    return
}

func NewEditAuthConfigRequest() (request *EditAuthConfigRequest) {
    request = &EditAuthConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "EditAuthConfig")
    
    
    return
}

func NewEditAuthConfigResponse() (response *EditAuthConfigResponse) {
    response = &EditAuthConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EditAuthConfig
// 修改登录配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) EditAuthConfig(request *EditAuthConfigRequest) (response *EditAuthConfigResponse, err error) {
    return c.EditAuthConfigWithContext(context.Background(), request)
}

// EditAuthConfig
// 修改登录配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) EditAuthConfigWithContext(ctx context.Context, request *EditAuthConfigRequest) (response *EditAuthConfigResponse, err error) {
    if request == nil {
        request = NewEditAuthConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "EditAuthConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EditAuthConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewEditAuthConfigResponse()
    err = c.Send(request, response)
    return
}

func NewListTablesRequest() (request *ListTablesRequest) {
    request = &ListTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "ListTables")
    
    
    return
}

func NewListTablesResponse() (response *ListTablesResponse) {
    response = &ListTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListTables
// 本接口(ListTables)用于查询文档型数据库所有表信息，包括表名、表中数据条数、表中数据量、索引个数及索引的大小等
//
// 
//
// 该接口跟 [DescribeTables](https://cloud.tencent.com/document/api/876/127962) 接口功能一致，后续该接口可能会下线，请使用 [DescribeTable](https://cloud.tencent.com/document/api/876/127962)接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LISTTABLE = "FailedOperation.ListTable"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_NOVALIDCONNECTION = "LimitExceeded.NoValidConnection"
//  LIMITEXCEEDED_OUTOFREADREQUESTQUOTA = "LimitExceeded.OutOfReadRequestQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE_RESOURCEOVERDUE = "ResourceUnavailable.ResourceOverdue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ListTables(request *ListTablesRequest) (response *ListTablesResponse, err error) {
    return c.ListTablesWithContext(context.Background(), request)
}

// ListTables
// 本接口(ListTables)用于查询文档型数据库所有表信息，包括表名、表中数据条数、表中数据量、索引个数及索引的大小等
//
// 
//
// 该接口跟 [DescribeTables](https://cloud.tencent.com/document/api/876/127962) 接口功能一致，后续该接口可能会下线，请使用 [DescribeTable](https://cloud.tencent.com/document/api/876/127962)接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LISTTABLE = "FailedOperation.ListTable"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_NOVALIDCONNECTION = "LimitExceeded.NoValidConnection"
//  LIMITEXCEEDED_OUTOFREADREQUESTQUOTA = "LimitExceeded.OutOfReadRequestQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE_RESOURCEOVERDUE = "ResourceUnavailable.ResourceOverdue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ListTablesWithContext(ctx context.Context, request *ListTablesRequest) (response *ListTablesResponse, err error) {
    if request == nil {
        request = NewListTablesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "ListTables")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListTables require credential")
    }

    request.SetContext(ctx)
    
    response = NewListTablesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCloudBaseGWAPIRequest() (request *ModifyCloudBaseGWAPIRequest) {
    request = &ModifyCloudBaseGWAPIRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "ModifyCloudBaseGWAPI")
    
    
    return
}

func NewModifyCloudBaseGWAPIResponse() (response *ModifyCloudBaseGWAPIResponse) {
    response = &ModifyCloudBaseGWAPIResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCloudBaseGWAPI
// 修改云开发网关API
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SERVICEEVIL = "InvalidParameter.ServiceEvil"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_CDNFREEZED = "ResourceUnavailable.CDNFreezed"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyCloudBaseGWAPI(request *ModifyCloudBaseGWAPIRequest) (response *ModifyCloudBaseGWAPIResponse, err error) {
    return c.ModifyCloudBaseGWAPIWithContext(context.Background(), request)
}

// ModifyCloudBaseGWAPI
// 修改云开发网关API
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SERVICEEVIL = "InvalidParameter.ServiceEvil"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_CDNFREEZED = "ResourceUnavailable.CDNFreezed"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyCloudBaseGWAPIWithContext(ctx context.Context, request *ModifyCloudBaseGWAPIRequest) (response *ModifyCloudBaseGWAPIResponse, err error) {
    if request == nil {
        request = NewModifyCloudBaseGWAPIRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "ModifyCloudBaseGWAPI")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCloudBaseGWAPI require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCloudBaseGWAPIResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClsTopicRequest() (request *ModifyClsTopicRequest) {
    request = &ModifyClsTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "ModifyClsTopic")
    
    
    return
}

func NewModifyClsTopicResponse() (response *ModifyClsTopicResponse) {
    response = &ModifyClsTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyClsTopic
// 修改日志主题
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_CONCURRENT = "LimitExceeded.Concurrent"
//  LIMITEXCEEDED_REQUEST = "LimitExceeded.Request"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyClsTopic(request *ModifyClsTopicRequest) (response *ModifyClsTopicResponse, err error) {
    return c.ModifyClsTopicWithContext(context.Background(), request)
}

// ModifyClsTopic
// 修改日志主题
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_CONCURRENT = "LimitExceeded.Concurrent"
//  LIMITEXCEEDED_REQUEST = "LimitExceeded.Request"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyClsTopicWithContext(ctx context.Context, request *ModifyClsTopicRequest) (response *ModifyClsTopicResponse, err error) {
    if request == nil {
        request = NewModifyClsTopicRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "ModifyClsTopic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClsTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClsTopicResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDatabaseACLRequest() (request *ModifyDatabaseACLRequest) {
    request = &ModifyDatabaseACLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "ModifyDatabaseACL")
    
    
    return
}

func NewModifyDatabaseACLResponse() (response *ModifyDatabaseACLResponse) {
    response = &ModifyDatabaseACLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDatabaseACL
// 修改文档型数据库权限
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCEINUSE_FSACLJOBUNDONE = "ResourceInUse.FsACLJobUnDone"
func (c *Client) ModifyDatabaseACL(request *ModifyDatabaseACLRequest) (response *ModifyDatabaseACLResponse, err error) {
    return c.ModifyDatabaseACLWithContext(context.Background(), request)
}

// ModifyDatabaseACL
// 修改文档型数据库权限
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCEINUSE_FSACLJOBUNDONE = "ResourceInUse.FsACLJobUnDone"
func (c *Client) ModifyDatabaseACLWithContext(ctx context.Context, request *ModifyDatabaseACLRequest) (response *ModifyDatabaseACLResponse, err error) {
    if request == nil {
        request = NewModifyDatabaseACLRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "ModifyDatabaseACL")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDatabaseACL require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDatabaseACLResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEnvRequest() (request *ModifyEnvRequest) {
    request = &ModifyEnvRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "ModifyEnv")
    
    
    return
}

func NewModifyEnvResponse() (response *ModifyEnvResponse) {
    response = &ModifyEnvResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyEnv
// 更新环境信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) ModifyEnv(request *ModifyEnvRequest) (response *ModifyEnvResponse, err error) {
    return c.ModifyEnvWithContext(context.Background(), request)
}

// ModifyEnv
// 更新环境信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) ModifyEnvWithContext(ctx context.Context, request *ModifyEnvRequest) (response *ModifyEnvResponse, err error) {
    if request == nil {
        request = NewModifyEnvRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "ModifyEnv")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEnv require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyEnvResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEnvPlanRequest() (request *ModifyEnvPlanRequest) {
    request = &ModifyEnvPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "ModifyEnvPlan")
    
    
    return
}

func NewModifyEnvPlanResponse() (response *ModifyEnvPlanResponse) {
    response = &ModifyEnvPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyEnvPlan
// 本接口用于变更云开发环境套餐。
//
// 该接口会自动下单并支付，会在腾讯云账户中扣除余额（余额不足会下单失败）。
//
// 该接口支持自动扣除代金券（AutoVoucher=true时），符合条件的代金券会被自动扣除。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
func (c *Client) ModifyEnvPlan(request *ModifyEnvPlanRequest) (response *ModifyEnvPlanResponse, err error) {
    return c.ModifyEnvPlanWithContext(context.Background(), request)
}

// ModifyEnvPlan
// 本接口用于变更云开发环境套餐。
//
// 该接口会自动下单并支付，会在腾讯云账户中扣除余额（余额不足会下单失败）。
//
// 该接口支持自动扣除代金券（AutoVoucher=true时），符合条件的代金券会被自动扣除。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
func (c *Client) ModifyEnvPlanWithContext(ctx context.Context, request *ModifyEnvPlanRequest) (response *ModifyEnvPlanResponse, err error) {
    if request == nil {
        request = NewModifyEnvPlanRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "ModifyEnvPlan")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEnvPlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyEnvPlanResponse()
    err = c.Send(request, response)
    return
}

func NewModifySafeRuleRequest() (request *ModifySafeRuleRequest) {
    request = &ModifySafeRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "ModifySafeRule")
    
    
    return
}

func NewModifySafeRuleResponse() (response *ModifySafeRuleResponse) {
    response = &ModifySafeRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySafeRule
// 设置数据库安全规则。
//
// 安全规则，用于控制C端用户的访问权限。详见 [安全规则介绍 ](https://cloud.tencent.com/document/product/876/123478)。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) ModifySafeRule(request *ModifySafeRuleRequest) (response *ModifySafeRuleResponse, err error) {
    return c.ModifySafeRuleWithContext(context.Background(), request)
}

// ModifySafeRule
// 设置数据库安全规则。
//
// 安全规则，用于控制C端用户的访问权限。详见 [安全规则介绍 ](https://cloud.tencent.com/document/product/876/123478)。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) ModifySafeRuleWithContext(ctx context.Context, request *ModifySafeRuleRequest) (response *ModifySafeRuleResponse, err error) {
    if request == nil {
        request = NewModifySafeRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "ModifySafeRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySafeRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySafeRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUserRequest() (request *ModifyUserRequest) {
    request = &ModifyUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "ModifyUser")
    
    
    return
}

func NewModifyUserResponse() (response *ModifyUserResponse) {
    response = &ModifyUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyUser
// 修改tcb用户
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DUPLICATEDDATA = "FailedOperation.DuplicatedData"
//  FAILEDOPERATION_FLEXDBRESOURCEOVERDUE = "FailedOperation.FlexdbResourceOverdue"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALID_PARAM = "InvalidParameter.INVALID_PARAM"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  RESOURCENOTFOUND_USERNOTEXISTS = "ResourceNotFound.UserNotExists"
func (c *Client) ModifyUser(request *ModifyUserRequest) (response *ModifyUserResponse, err error) {
    return c.ModifyUserWithContext(context.Background(), request)
}

// ModifyUser
// 修改tcb用户
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DUPLICATEDDATA = "FailedOperation.DuplicatedData"
//  FAILEDOPERATION_FLEXDBRESOURCEOVERDUE = "FailedOperation.FlexdbResourceOverdue"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALID_PARAM = "InvalidParameter.INVALID_PARAM"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  RESOURCENOTFOUND_USERNOTEXISTS = "ResourceNotFound.UserNotExists"
func (c *Client) ModifyUserWithContext(ctx context.Context, request *ModifyUserRequest) (response *ModifyUserResponse, err error) {
    if request == nil {
        request = NewModifyUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "ModifyUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUserResponse()
    err = c.Send(request, response)
    return
}

func NewReinstateEnvRequest() (request *ReinstateEnvRequest) {
    request = &ReinstateEnvRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "ReinstateEnv")
    
    
    return
}

func NewReinstateEnvResponse() (response *ReinstateEnvResponse) {
    response = &ReinstateEnvResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReinstateEnv
// 针对已隔离的免费环境，可以通过本接口将其恢复访问。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ReinstateEnv(request *ReinstateEnvRequest) (response *ReinstateEnvResponse, err error) {
    return c.ReinstateEnvWithContext(context.Background(), request)
}

// ReinstateEnv
// 针对已隔离的免费环境，可以通过本接口将其恢复访问。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ReinstateEnvWithContext(ctx context.Context, request *ReinstateEnvRequest) (response *ReinstateEnvResponse, err error) {
    if request == nil {
        request = NewReinstateEnvRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "ReinstateEnv")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReinstateEnv require credential")
    }

    request.SetContext(ctx)
    
    response = NewReinstateEnvResponse()
    err = c.Send(request, response)
    return
}

func NewRenewEnvRequest() (request *RenewEnvRequest) {
    request = &RenewEnvRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "RenewEnv")
    
    
    return
}

func NewRenewEnvResponse() (response *RenewEnvResponse) {
    response = &RenewEnvResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RenewEnv
// 本接口用于云开发环境套餐续费。
//
// 该接口会自动下单并支付，会在腾讯云账户中扣除余额（余额不足会下单失败）。
//
// 该接口支持自动扣除代金券（AutoVoucher=true时），符合条件的代金券会被自动扣除。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
func (c *Client) RenewEnv(request *RenewEnvRequest) (response *RenewEnvResponse, err error) {
    return c.RenewEnvWithContext(context.Background(), request)
}

// RenewEnv
// 本接口用于云开发环境套餐续费。
//
// 该接口会自动下单并支付，会在腾讯云账户中扣除余额（余额不足会下单失败）。
//
// 该接口支持自动扣除代金券（AutoVoucher=true时），符合条件的代金券会被自动扣除。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
func (c *Client) RenewEnvWithContext(ctx context.Context, request *RenewEnvRequest) (response *RenewEnvResponse, err error) {
    if request == nil {
        request = NewRenewEnvRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "RenewEnv")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RenewEnv require credential")
    }

    request.SetContext(ctx)
    
    response = NewRenewEnvResponse()
    err = c.Send(request, response)
    return
}

func NewRunCommandsRequest() (request *RunCommandsRequest) {
    request = &RunCommandsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "RunCommands")
    
    
    return
}

func NewRunCommandsResponse() (response *RunCommandsResponse) {
    response = &RunCommandsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RunCommands
// 本接口用于执行文档型数据库命令
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDDOC = "InvalidParameterValue.InvalidDoc"
//  LIMITEXCEEDED_NOVALIDCONNECTION = "LimitExceeded.NoValidConnection"
//  LIMITEXCEEDED_OUTOFRESULTSIZELIMIT = "LimitExceeded.OutOfResultSizeLimit"
//  RESOURCENOTFOUND_CONNECTOR = "ResourceNotFound.Connector"
//  RESOURCENOTFOUND_TABLE = "ResourceNotFound.Table"
//  RESOURCEUNAVAILABLE_MONGOISOLATED = "ResourceUnavailable.MongoIsolated"
//  RESOURCEUNAVAILABLE_RESOURCEOVERDUE = "ResourceUnavailable.ResourceOverdue"
func (c *Client) RunCommands(request *RunCommandsRequest) (response *RunCommandsResponse, err error) {
    return c.RunCommandsWithContext(context.Background(), request)
}

// RunCommands
// 本接口用于执行文档型数据库命令
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDDOC = "InvalidParameterValue.InvalidDoc"
//  LIMITEXCEEDED_NOVALIDCONNECTION = "LimitExceeded.NoValidConnection"
//  LIMITEXCEEDED_OUTOFRESULTSIZELIMIT = "LimitExceeded.OutOfResultSizeLimit"
//  RESOURCENOTFOUND_CONNECTOR = "ResourceNotFound.Connector"
//  RESOURCENOTFOUND_TABLE = "ResourceNotFound.Table"
//  RESOURCEUNAVAILABLE_MONGOISOLATED = "ResourceUnavailable.MongoIsolated"
//  RESOURCEUNAVAILABLE_RESOURCEOVERDUE = "ResourceUnavailable.ResourceOverdue"
func (c *Client) RunCommandsWithContext(ctx context.Context, request *RunCommandsRequest) (response *RunCommandsResponse, err error) {
    if request == nil {
        request = NewRunCommandsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "RunCommands")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunCommands require credential")
    }

    request.SetContext(ctx)
    
    response = NewRunCommandsResponse()
    err = c.Send(request, response)
    return
}

func NewRunSqlRequest() (request *RunSqlRequest) {
    request = &RunSqlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "RunSql")
    
    
    return
}

func NewRunSqlResponse() (response *RunSqlResponse) {
    response = &RunSqlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RunSql
// 执行MySQL语句
//
// 
//
// 该接口用来执行 MySql 语句，比如创建表格、插入数据、修改数据、删除字段、添加索引等可以通过sql 语句实现的都可以使用该接口
//
// 
//
// 调用该接口前需要先查询Mysql是否开通，可通过 [DescribeCreateMySQLResult ](https://cloud.tencent.com/document/api/876/128185) 查询，只有开通成功才能操作
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASECONNECTERROR = "FailedOperation.DatabaseConnectError"
//  FAILEDOPERATION_DATABASEEXECSQLERROR = "FailedOperation.DatabaseExecSqlError"
//  FAILEDOPERATION_DATABASESCHEMAERROR = "FailedOperation.DatabaseSchemaError"
//  FAILEDOPERATION_EMPTYDATABASEENDPOINT = "FailedOperation.EmptyDatabaseEndpoint"
//  FAILEDOPERATION_TDSQLPAUSED = "FailedOperation.TdsqlPaused"
//  INTERNALERROR_SYS_ERR = "InternalError.SYS_ERR"
//  INVALIDPARAMETER_INVALID_PARAM = "InvalidParameter.INVALID_PARAM"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_TABLENOTFOUND = "ResourceNotFound.TableNotFound"
//  UNSUPPORTEDOPERATION_TOOMANYTABLES = "UnsupportedOperation.TooManyTables"
func (c *Client) RunSql(request *RunSqlRequest) (response *RunSqlResponse, err error) {
    return c.RunSqlWithContext(context.Background(), request)
}

// RunSql
// 执行MySQL语句
//
// 
//
// 该接口用来执行 MySql 语句，比如创建表格、插入数据、修改数据、删除字段、添加索引等可以通过sql 语句实现的都可以使用该接口
//
// 
//
// 调用该接口前需要先查询Mysql是否开通，可通过 [DescribeCreateMySQLResult ](https://cloud.tencent.com/document/api/876/128185) 查询，只有开通成功才能操作
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASECONNECTERROR = "FailedOperation.DatabaseConnectError"
//  FAILEDOPERATION_DATABASEEXECSQLERROR = "FailedOperation.DatabaseExecSqlError"
//  FAILEDOPERATION_DATABASESCHEMAERROR = "FailedOperation.DatabaseSchemaError"
//  FAILEDOPERATION_EMPTYDATABASEENDPOINT = "FailedOperation.EmptyDatabaseEndpoint"
//  FAILEDOPERATION_TDSQLPAUSED = "FailedOperation.TdsqlPaused"
//  INTERNALERROR_SYS_ERR = "InternalError.SYS_ERR"
//  INVALIDPARAMETER_INVALID_PARAM = "InvalidParameter.INVALID_PARAM"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_TABLENOTFOUND = "ResourceNotFound.TableNotFound"
//  UNSUPPORTEDOPERATION_TOOMANYTABLES = "UnsupportedOperation.TooManyTables"
func (c *Client) RunSqlWithContext(ctx context.Context, request *RunSqlRequest) (response *RunSqlResponse, err error) {
    if request == nil {
        request = NewRunSqlRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "RunSql")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunSql require credential")
    }

    request.SetContext(ctx)
    
    response = NewRunSqlResponse()
    err = c.Send(request, response)
    return
}

func NewSearchClsLogRequest() (request *SearchClsLogRequest) {
    request = &SearchClsLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "SearchClsLog")
    
    
    return
}

func NewSearchClsLogResponse() (response *SearchClsLogResponse) {
    response = &SearchClsLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SearchClsLog
// 搜索用户调用日志
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_INVALIDCONTEXT = "FailedOperation.InvalidContext"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_QUERYERROR = "FailedOperation.QueryError"
//  FAILEDOPERATION_SYNTAXERROR = "FailedOperation.SyntaxError"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) SearchClsLog(request *SearchClsLogRequest) (response *SearchClsLogResponse, err error) {
    return c.SearchClsLogWithContext(context.Background(), request)
}

// SearchClsLog
// 搜索用户调用日志
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_INVALIDCONTEXT = "FailedOperation.InvalidContext"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_QUERYERROR = "FailedOperation.QueryError"
//  FAILEDOPERATION_SYNTAXERROR = "FailedOperation.SyntaxError"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) SearchClsLogWithContext(ctx context.Context, request *SearchClsLogRequest) (response *SearchClsLogResponse, err error) {
    if request == nil {
        request = NewSearchClsLogRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "SearchClsLog")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchClsLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchClsLogResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateTableRequest() (request *UpdateTableRequest) {
    request = &UpdateTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "UpdateTable")
    
    
    return
}

func NewUpdateTableResponse() (response *UpdateTableResponse) {
    response = &UpdateTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateTable
// 本接口(UpdateTable)用于修改文档型数据库表信息，当前可以支持创建和删除索引
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LISTTABLE = "FailedOperation.ListTable"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_NOVALIDCONNECTION = "LimitExceeded.NoValidConnection"
//  LIMITEXCEEDED_OUTOFINDEXQUOTA = "LimitExceeded.OutOfIndexQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE_INDEXCREATING = "ResourceInUse.IndexCreating"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CONNECTOR = "ResourceNotFound.Connector"
//  RESOURCENOTFOUND_TABLE = "ResourceNotFound.Table"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_INDEXOPTIONSCONFLICT = "ResourceUnavailable.IndexOptionsConflict"
//  RESOURCEUNAVAILABLE_MONGOISOLATED = "ResourceUnavailable.MongoIsolated"
//  RESOURCEUNAVAILABLE_RESOURCEOVERDUE = "ResourceUnavailable.ResourceOverdue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) UpdateTable(request *UpdateTableRequest) (response *UpdateTableResponse, err error) {
    return c.UpdateTableWithContext(context.Background(), request)
}

// UpdateTable
// 本接口(UpdateTable)用于修改文档型数据库表信息，当前可以支持创建和删除索引
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LISTTABLE = "FailedOperation.ListTable"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_NOVALIDCONNECTION = "LimitExceeded.NoValidConnection"
//  LIMITEXCEEDED_OUTOFINDEXQUOTA = "LimitExceeded.OutOfIndexQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE_INDEXCREATING = "ResourceInUse.IndexCreating"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CONNECTOR = "ResourceNotFound.Connector"
//  RESOURCENOTFOUND_TABLE = "ResourceNotFound.Table"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_INDEXOPTIONSCONFLICT = "ResourceUnavailable.IndexOptionsConflict"
//  RESOURCEUNAVAILABLE_MONGOISOLATED = "ResourceUnavailable.MongoIsolated"
//  RESOURCEUNAVAILABLE_RESOURCEOVERDUE = "ResourceUnavailable.ResourceOverdue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) UpdateTableWithContext(ctx context.Context, request *UpdateTableRequest) (response *UpdateTableResponse, err error) {
    if request == nil {
        request = NewUpdateTableRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "UpdateTable")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateTable require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateTableResponse()
    err = c.Send(request, response)
    return
}
