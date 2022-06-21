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

package v20200715

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type DescribeEisConnectorConfigRequestParams struct {
	// 连接器名称
	ConnectorName *string `json:"ConnectorName,omitempty" name:"ConnectorName"`

	// 连接器版本
	ConnectorVersion *string `json:"ConnectorVersion,omitempty" name:"ConnectorVersion"`
}

type DescribeEisConnectorConfigRequest struct {
	*tchttp.BaseRequest
	
	// 连接器名称
	ConnectorName *string `json:"ConnectorName,omitempty" name:"ConnectorName"`

	// 连接器版本
	ConnectorVersion *string `json:"ConnectorVersion,omitempty" name:"ConnectorVersion"`
}

func (r *DescribeEisConnectorConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEisConnectorConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConnectorName")
	delete(f, "ConnectorVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEisConnectorConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEisConnectorConfigResponseParams struct {
	// 连接器配置参数描述（json结构），示例如下：
	// {
	//     "attributes":{
	//         "description":"测试", // 连接器的描述
	//         "displayName":"测试", // 连接器的展示名
	//         "name":"test", // 连接器的名称
	//         "version":"1.0.0" // 连接器的版本号
	//     },
	//     "properties":[
	//         {
	//             "attributes":{
	//                 "displayName":"企业ID", // 参数的展示名
	//                 "name":"para1", // 参数名
	//                 "required":"true", // 是否必填
	//                 "type":"int" // 参数的类型
	//             }
	//         },
	//         {
	//             "attributes":{
	//                 "displayName":"成员管理密钥",
	//                 "name":"para2",
	//                 "required":"true",
	//                 "type":"float"
	//             }
	//         },
	//         {
	//             "attributes":{
	//                 "displayName":"应用管理密钥",
	//                 "name":"para3",
	//                 "required":"true",
	//                 "type":"string"
	//             }
	//         },
	//         {
	//             "attributes":{
	//                 "displayName":"企业ID",
	//                 "name":"para4",
	//                 "required":"true",
	//                 "type":"decimal"
	//             }
	//         },
	//         {
	//             "attributes":{
	//                 "displayName":"成员管理密钥",
	//                 "name":"para5",
	//                 "required":"true",
	//                 "type":"bool"
	//             }
	//         },
	//         {
	//             "attributes":{
	//                 "displayName":"应用管理密钥",
	//                 "name":"para6",
	//                 "required":"true",
	//                 "type":"date"
	//             }
	//         },
	//         {
	//             "attributes":{
	//                 "displayName":"企业ID",
	//                 "name":"para7",
	//                 "required":"true",
	//                 "type":"time"
	//             }
	//         },
	//         {
	//             "attributes":{
	//                 "displayName":"成员管理密钥",
	//                 "name":"para8",
	//                 "required":"true",
	//                 "type":"datetime"
	//             }
	//         },
	//         {
	//             "attributes":{
	//                 "displayName":"应用管理密钥",
	//                 "name":"para9",
	//                 "required":"true",
	//                 "type":"map"
	//             },
	//             "children":[
	//                 {
	//                     "attributes":{
	//                         "displayName":"key",
	//                         "name":"key",
	//                         "required":"true",
	//                         "type":"string"
	//                     }
	//                 },
	//                 {
	//                     "attributes":{
	//                         "displayName":"value",
	//                         "name":"value",
	//                         "required":"true",
	//                         "type":"any"
	//                     }
	//                 }
	//             ]
	//         },
	//         {
	//             "attributes":{
	//                 "displayName":"企业ID",
	//                 "name":"para10",
	//                 "required":"true",
	//                 "type":"list" // list，list里元素的类型是结构体，children里是结构体的描述
	//             },
	//             "children":[
	//                 {
	//                     "attributes":{
	//                         "displayName":"field1",
	//                         "name":"field1",
	//                         "required":"true",
	//                         "type":"string"
	//                     }
	//                 },
	//                 {
	//                     "attributes":{
	//                         "displayName":"field2",
	//                         "name":"field2",
	//                         "required":"true",
	//                         "type":"any"
	//                     }
	//                 }
	//             ]
	//         },
	//         {
	//             "attributes":{
	//                 "displayName":"成员管理密钥",
	//                 "name":"para11",
	//                 "required":"true",
	//                 "type":"struct"
	//             },
	//             "children":[
	//                 {
	//                     "attributes":{
	//                         "displayName":"field1", // 结构体属性的展示名
	//                         "name":"field1", // 结构体属性的名称
	//                         "required":"true", // 是否必填
	//                         "type":"string" // 属性的类型
	//                     }
	//                 },
	//                 {
	//                     "attributes":{
	//                         "displayName":"field2",
	//                         "name":"field2",
	//                         "required":"true",
	//                         "type":"any"
	//                     }
	//                 }
	//             ]
	//         },
	//         {
	//             "attributes":{
	//                 "displayName":"应用管理密钥",
	//                 "name":"para12",
	//                 "required":"true",
	//                 "type":"enum"
	//             },
	//             "children":[
	//                 {
	//                     "attributes":{
	//                         "displayName":"PC", // 枚举值的展示名
	//                         "name":"PC" // 枚举值的名称
	//                     }
	//                 },
	//                 {
	//                     "attributes":{
	//                         "displayName":"MAC",
	//                         "name":"MAC"
	//                     }
	//                 }
	//             ]
	//         }
	//     ]
	// }
	ConnectorParameter *string `json:"ConnectorParameter,omitempty" name:"ConnectorParameter"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEisConnectorConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEisConnectorConfigResponseParams `json:"Response"`
}

func (r *DescribeEisConnectorConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEisConnectorConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EisConnectionOperation struct {
	// 连接器操作名称
	OperationName *string `json:"OperationName,omitempty" name:"OperationName"`

	// 连接器展示名称
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// 操作是否为触发器
	IsTrigger *bool `json:"IsTrigger,omitempty" name:"IsTrigger"`
}

type EisConnectorSummary struct {
	// 连接器名称
	ConnectorName *string `json:"ConnectorName,omitempty" name:"ConnectorName"`

	// 连接器展示名称
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// 连接器对应企业
	Company *string `json:"Company,omitempty" name:"Company"`

	// 连接器对应产品
	Product *string `json:"Product,omitempty" name:"Product"`

	// 连接器版本
	ConnectorVersion *string `json:"ConnectorVersion,omitempty" name:"ConnectorVersion"`

	// 连接器创建时间
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`
}

// Predefined struct for user
type ListEisConnectorOperationsRequestParams struct {
	// 连接器名称
	ConnectorName *string `json:"ConnectorName,omitempty" name:"ConnectorName"`

	// 连接器版本
	ConnectorVersion *string `json:"ConnectorVersion,omitempty" name:"ConnectorVersion"`
}

type ListEisConnectorOperationsRequest struct {
	*tchttp.BaseRequest
	
	// 连接器名称
	ConnectorName *string `json:"ConnectorName,omitempty" name:"ConnectorName"`

	// 连接器版本
	ConnectorVersion *string `json:"ConnectorVersion,omitempty" name:"ConnectorVersion"`
}

func (r *ListEisConnectorOperationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListEisConnectorOperationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConnectorName")
	delete(f, "ConnectorVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListEisConnectorOperationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListEisConnectorOperationsResponseParams struct {
	// 连接器列表
	Operations []*EisConnectionOperation `json:"Operations,omitempty" name:"Operations"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListEisConnectorOperationsResponse struct {
	*tchttp.BaseResponse
	Response *ListEisConnectorOperationsResponseParams `json:"Response"`
}

func (r *ListEisConnectorOperationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListEisConnectorOperationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListEisConnectorsRequestParams struct {
	// 连接器名称,非必输，如输入则按照输入值模糊匹配
	ConnectorName *string `json:"ConnectorName,omitempty" name:"ConnectorName"`

	// 分页参数,数据偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页参数,每页显示的条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type ListEisConnectorsRequest struct {
	*tchttp.BaseRequest
	
	// 连接器名称,非必输，如输入则按照输入值模糊匹配
	ConnectorName *string `json:"ConnectorName,omitempty" name:"ConnectorName"`

	// 分页参数,数据偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页参数,每页显示的条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListEisConnectorsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListEisConnectorsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConnectorName")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListEisConnectorsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListEisConnectorsResponseParams struct {
	// 连接器总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 连接器列表
	Connectors []*EisConnectorSummary `json:"Connectors,omitempty" name:"Connectors"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListEisConnectorsResponse struct {
	*tchttp.BaseResponse
	Response *ListEisConnectorsResponseParams `json:"Response"`
}

func (r *ListEisConnectorsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListEisConnectorsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}