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

package v20241024

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type AcceptTccVpcEndPointConnectRequestParams struct {
	// 终端节点服务Id
	EndPointServiceId *string `json:"EndPointServiceId,omitnil,omitempty" name:"EndPointServiceId"`

	// 终端节点id
	EndPointId []*string `json:"EndPointId,omitnil,omitempty" name:"EndPointId"`

	// 是否接受连接
	AcceptFlag *bool `json:"AcceptFlag,omitnil,omitempty" name:"AcceptFlag"`
}

type AcceptTccVpcEndPointConnectRequest struct {
	*tchttp.BaseRequest
	
	// 终端节点服务Id
	EndPointServiceId *string `json:"EndPointServiceId,omitnil,omitempty" name:"EndPointServiceId"`

	// 终端节点id
	EndPointId []*string `json:"EndPointId,omitnil,omitempty" name:"EndPointId"`

	// 是否接受连接
	AcceptFlag *bool `json:"AcceptFlag,omitnil,omitempty" name:"AcceptFlag"`
}

func (r *AcceptTccVpcEndPointConnectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AcceptTccVpcEndPointConnectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EndPointServiceId")
	delete(f, "EndPointId")
	delete(f, "AcceptFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AcceptTccVpcEndPointConnectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AcceptTccVpcEndPointConnectResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AcceptTccVpcEndPointConnectResponse struct {
	*tchttp.BaseResponse
	Response *AcceptTccVpcEndPointConnectResponseParams `json:"Response"`
}

func (r *AcceptTccVpcEndPointConnectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AcceptTccVpcEndPointConnectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindTccVpcEndPointServiceWhiteListRequestParams struct {
	// 终端节点服务Id
	EndPointServiceId *string `json:"EndPointServiceId,omitnil,omitempty" name:"EndPointServiceId"`

	// 需要开白的用户Uin
	UserUin *string `json:"UserUin,omitnil,omitempty" name:"UserUin"`

	// 用户描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type BindTccVpcEndPointServiceWhiteListRequest struct {
	*tchttp.BaseRequest
	
	// 终端节点服务Id
	EndPointServiceId *string `json:"EndPointServiceId,omitnil,omitempty" name:"EndPointServiceId"`

	// 需要开白的用户Uin
	UserUin *string `json:"UserUin,omitnil,omitempty" name:"UserUin"`

	// 用户描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *BindTccVpcEndPointServiceWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindTccVpcEndPointServiceWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EndPointServiceId")
	delete(f, "UserUin")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindTccVpcEndPointServiceWhiteListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindTccVpcEndPointServiceWhiteListResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BindTccVpcEndPointServiceWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *BindTccVpcEndPointServiceWhiteListResponseParams `json:"Response"`
}

func (r *BindTccVpcEndPointServiceWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindTccVpcEndPointServiceWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTccCatalogRequestParams struct {
	// 数据目录Id
	CatalogId *string `json:"CatalogId,omitnil,omitempty" name:"CatalogId"`
}

type DescribeTccCatalogRequest struct {
	*tchttp.BaseRequest
	
	// 数据目录Id
	CatalogId *string `json:"CatalogId,omitnil,omitempty" name:"CatalogId"`
}

func (r *DescribeTccCatalogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTccCatalogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CatalogId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTccCatalogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTccCatalogResponseParams struct {
	// Tcc数据目录信息
	TccCatalog *TccCatalogConfig `json:"TccCatalog,omitnil,omitempty" name:"TccCatalog"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTccCatalogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTccCatalogResponseParams `json:"Response"`
}

func (r *DescribeTccCatalogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTccCatalogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTccCatalogsRequestParams struct {
	// 数据目录Id
	CatalogId *string `json:"CatalogId,omitnil,omitempty" name:"CatalogId"`

	// 数据目录名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数据目录类型,当前支持：TCC-HIVE
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 状态信息：注册中0，待测试1，连接成功2，连接失败3，删除中4，已删除5
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 操作人uin
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`
}

type DescribeTccCatalogsRequest struct {
	*tchttp.BaseRequest
	
	// 数据目录Id
	CatalogId *string `json:"CatalogId,omitnil,omitempty" name:"CatalogId"`

	// 数据目录名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数据目录类型,当前支持：TCC-HIVE
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 状态信息：注册中0，待测试1，连接成功2，连接失败3，删除中4，已删除5
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 操作人uin
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`
}

func (r *DescribeTccCatalogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTccCatalogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CatalogId")
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "Status")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTccCatalogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTccCatalogsResponseParams struct {
	// 数据目录列表
	TccCatalogSet []*TccCatalogSet `json:"TccCatalogSet,omitnil,omitempty" name:"TccCatalogSet"`

	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTccCatalogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTccCatalogsResponseParams `json:"Response"`
}

func (r *DescribeTccCatalogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTccCatalogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetWork struct {
	// vpc实例id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// vpc网段
	VpcCidrBlock *string `json:"VpcCidrBlock,omitnil,omitempty" name:"VpcCidrBlock"`

	// 子网实例id
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 子网网段
	SubnetCidrBlock *string `json:"SubnetCidrBlock,omitnil,omitempty" name:"SubnetCidrBlock"`
}

type TccCatalogConfig struct {
	// 数据目录唯一id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 数据目录名字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数据目录类型,当前支持：TCC-HIVE
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 描述信息
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 状态信息：注册中0，待测试1，连接成功2，连接失败3，删除中4，已删除5
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Tcc数据目录连接信息
	Connection *TccConnectionConfig `json:"Connection,omitnil,omitempty" name:"Connection"`

	// 操作人uin
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type TccCatalogSet struct {
	// 数据目录唯一id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 数据目录名字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数据目录类型,当前支持：TCC-HIVE
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 状态信息：注册中0，待测试1，连接成功2，连接失败3，删除中4，已删除5
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 操作人uin
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type TccConnection struct {
	// 引擎终端节点服务Id
	EndpointServiceId *string `json:"EndpointServiceId,omitnil,omitempty" name:"EndpointServiceId"`

	// 元数据连接串
	MetaStoreUrl *string `json:"MetaStoreUrl,omitnil,omitempty" name:"MetaStoreUrl"`

	// 网络信息
	NetWork *NetWork `json:"NetWork,omitnil,omitempty" name:"NetWork"`

	// hive版本
	HiveVersion *string `json:"HiveVersion,omitnil,omitempty" name:"HiveVersion"`

	// hive location
	Location *string `json:"Location,omitnil,omitempty" name:"Location"`

	// HMS终端节点服务
	HmsEndpointServiceId *string `json:"HmsEndpointServiceId,omitnil,omitempty" name:"HmsEndpointServiceId"`
}

type TccConnectionConfig struct {
	// Tcc数据目录连接配置
	TccHive *TccConnection `json:"TccHive,omitnil,omitempty" name:"TccHive"`
}