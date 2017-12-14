// Copyright 1999-2018 Tencent Ltd.
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
package dcdb

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-04-11"

type Client struct {
    common.Client
}

func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey)
    return
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithSecretId(credential.SecretId, credential.SecretKey).
        WithProfile(clientProfile)
    return
}


func NewCreateDCDBInstanceRequest() (request *CreateDCDBInstanceRequest) {
    request = &CreateDCDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "CreateDCDBInstance")
    return
}

func NewCreateDCDBInstanceResponse() (response *CreateDCDBInstanceResponse) {
    response = &CreateDCDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateDCDBInstance）用于创建包年包月的云数据库实例，可通过传入实例规格、数据库版本号、购买时长等信息创建云数据库实例。
func (c *Client) CreateDCDBInstance(request *CreateDCDBInstanceRequest) (response *CreateDCDBInstanceResponse, err error) {
    if request == nil {
        request = NewCreateDCDBInstanceRequest()
    }
    response = NewCreateDCDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBLogFilesRequest() (request *DescribeDBLogFilesRequest) {
    request = &DescribeDBLogFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDBLogFiles")
    return
}

func NewDescribeDBLogFilesResponse() (response *DescribeDBLogFilesResponse) {
    response = &DescribeDBLogFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeDBLogFiles)用于获取数据库的各种日志列表，包括冷备、binlog、errlog和slowlog。
func (c *Client) DescribeDBLogFiles(request *DescribeDBLogFilesRequest) (response *DescribeDBLogFilesResponse, err error) {
    if request == nil {
        request = NewDescribeDBLogFilesRequest()
    }
    response = NewDescribeDBLogFilesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDCDBInstancesRequest() (request *DescribeDCDBInstancesRequest) {
    request = &DescribeDCDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDCDBInstances")
    return
}

func NewDescribeDCDBInstancesResponse() (response *DescribeDCDBInstancesResponse) {
    response = &DescribeDCDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询云数据库实例列表，支持通过项目ID、实例ID、内网地址、实例名称等来筛选实例。
// 如果不指定任何筛选条件，则默认返回10条实例记录，单次请求最多支持返回100条实例记录。
func (c *Client) DescribeDCDBInstances(request *DescribeDCDBInstancesRequest) (response *DescribeDCDBInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeDCDBInstancesRequest()
    }
    response = NewDescribeDCDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDCDBPriceRequest() (request *DescribeDCDBPriceRequest) {
    request = &DescribeDCDBPriceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDCDBPrice")
    return
}

func NewDescribeDCDBPriceResponse() (response *DescribeDCDBPriceResponse) {
    response = &DescribeDCDBPriceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeDCDBPrice）用于在购买实例前，查询实例的价格。
func (c *Client) DescribeDCDBPrice(request *DescribeDCDBPriceRequest) (response *DescribeDCDBPriceResponse, err error) {
    if request == nil {
        request = NewDescribeDCDBPriceRequest()
    }
    response = NewDescribeDCDBPriceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDCDBRenewalPriceRequest() (request *DescribeDCDBRenewalPriceRequest) {
    request = &DescribeDCDBRenewalPriceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDCDBRenewalPrice")
    return
}

func NewDescribeDCDBRenewalPriceResponse() (response *DescribeDCDBRenewalPriceResponse) {
    response = &DescribeDCDBRenewalPriceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeDCDBRenewalPrice）用于在续费分布式数据库实例时，查询续费的价格。
func (c *Client) DescribeDCDBRenewalPrice(request *DescribeDCDBRenewalPriceRequest) (response *DescribeDCDBRenewalPriceResponse, err error) {
    if request == nil {
        request = NewDescribeDCDBRenewalPriceRequest()
    }
    response = NewDescribeDCDBRenewalPriceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDCDBSaleInfoRequest() (request *DescribeDCDBSaleInfoRequest) {
    request = &DescribeDCDBSaleInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDCDBSaleInfo")
    return
}

func NewDescribeDCDBSaleInfoResponse() (response *DescribeDCDBSaleInfoResponse) {
    response = &DescribeDCDBSaleInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeDCDBSaleInfo)用于查询分布式数据库可售卖的地域和可用区信息。
func (c *Client) DescribeDCDBSaleInfo(request *DescribeDCDBSaleInfoRequest) (response *DescribeDCDBSaleInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDCDBSaleInfoRequest()
    }
    response = NewDescribeDCDBSaleInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDCDBUpgradePriceRequest() (request *DescribeDCDBUpgradePriceRequest) {
    request = &DescribeDCDBUpgradePriceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDCDBUpgradePrice")
    return
}

func NewDescribeDCDBUpgradePriceResponse() (response *DescribeDCDBUpgradePriceResponse) {
    response = &DescribeDCDBUpgradePriceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeDCDBUpgradePrice）用于查询升级分布式数据库实例价格。
func (c *Client) DescribeDCDBUpgradePrice(request *DescribeDCDBUpgradePriceRequest) (response *DescribeDCDBUpgradePriceResponse, err error) {
    if request == nil {
        request = NewDescribeDCDBUpgradePriceRequest()
    }
    response = NewDescribeDCDBUpgradePriceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrdersRequest() (request *DescribeOrdersRequest) {
    request = &DescribeOrdersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeOrders")
    return
}

func NewDescribeOrdersResponse() (response *DescribeOrdersResponse) {
    response = &DescribeOrdersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeOrders）用于查询分布式数据库订单信息。传入订单Id来查询订单关联的分布式数据库实例，和对应的任务流程ID。
func (c *Client) DescribeOrders(request *DescribeOrdersRequest) (response *DescribeOrdersResponse, err error) {
    if request == nil {
        request = NewDescribeOrdersRequest()
    }
    response = NewDescribeOrdersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeShardSpecRequest() (request *DescribeShardSpecRequest) {
    request = &DescribeShardSpecRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeShardSpec")
    return
}

func NewDescribeShardSpecResponse() (response *DescribeShardSpecResponse) {
    response = &DescribeShardSpecResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询可创建的分布式数据库可售卖的分片规格配置。
func (c *Client) DescribeShardSpec(request *DescribeShardSpecRequest) (response *DescribeShardSpecResponse, err error) {
    if request == nil {
        request = NewDescribeShardSpecRequest()
    }
    response = NewDescribeShardSpecResponse()
    err = c.Send(request, response)
    return
}

func NewRenewDCDBInstanceRequest() (request *RenewDCDBInstanceRequest) {
    request = &RenewDCDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "RenewDCDBInstance")
    return
}

func NewRenewDCDBInstanceResponse() (response *RenewDCDBInstanceResponse) {
    response = &RenewDCDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（RenewDCDBInstance）用于续费分布式数据库实例。
func (c *Client) RenewDCDBInstance(request *RenewDCDBInstanceRequest) (response *RenewDCDBInstanceResponse, err error) {
    if request == nil {
        request = NewRenewDCDBInstanceRequest()
    }
    response = NewRenewDCDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeDCDBInstanceRequest() (request *UpgradeDCDBInstanceRequest) {
    request = &UpgradeDCDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "UpgradeDCDBInstance")
    return
}

func NewUpgradeDCDBInstanceResponse() (response *UpgradeDCDBInstanceResponse) {
    response = &UpgradeDCDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（UpgradeDCDBInstance）用于升级分布式数据库实例。本接口完成下单和支付两个动作，如果发生支付失败的错误，调用用户账户相关接口中的支付订单接口（PayDeals）重新支付即可。
func (c *Client) UpgradeDCDBInstance(request *UpgradeDCDBInstanceRequest) (response *UpgradeDCDBInstanceResponse, err error) {
    if request == nil {
        request = NewUpgradeDCDBInstanceRequest()
    }
    response = NewUpgradeDCDBInstanceResponse()
    err = c.Send(request, response)
    return
}
