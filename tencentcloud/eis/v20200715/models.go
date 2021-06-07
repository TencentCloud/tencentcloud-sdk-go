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
    "errors"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEisConnectorConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConnectorName")
	delete(f, "ConnectorVersion")
	if len(f) > 0 {
		return errors.New("DescribeEisConnectorConfigRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEisConnectorConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 连接器配置参数描述（json结构），示例如下：
	// {
	//     "attributes":{
	//         "description":"测试",
	//         "displayName":"test",
	//         "name":"test",
	//         "version":"1.0.0"
	//     },
	//     "properties":[
	//         {
	//             "displayName":"日期",
	//             "name":"prop1",
	//             "required":"true",
	//             "type":"date"
	//         }
	//     ],
	//     "operations":{
	//         "get-info":[
	//             {
	//                 "displayName":"para1",
	//                 "name":"para1",
	//                 "required":"true",
	//                 "type":"int"
	//             },
	//             {
	//                 "displayName":"para2",
	//                 "name":"para2",
	//                 "required":"true",
	//                 "type":"float"
	//             },
	//             {
	//                 "displayName":"para3",
	//                 "name":"para3",
	//                 "required":"true",
	//                 "type":"string"
	//             },
	//             {
	//                 "displayName":"para4",
	//                 "name":"para4",
	//                 "required":"true",
	//                 "type":"decimal"
	//             },
	//             {
	//                 "displayName":"para5",
	//                 "name":"para5",
	//                 "required":"true",
	//                 "type":"bool"
	//             },
	//             {
	//                 "displayName":"para6",
	//                 "name":"para6",
	//                 "required":"true",
	//                 "type":"date"
	//             },
	//             {
	//                 "displayName":"para7",
	//                 "name":"para7",
	//                 "required":"true",
	//                 "type":"time"
	//             },
	//             {
	//                 "displayName":"para8",
	//                 "name":"para8",
	//                 "required":"true",
	//                 "type":"datetime"
	//             },
	//             {
	//                 "displayName":"para9",
	//                 "name":"para9",
	//                 "required":"true",
	//                 "type":"struct",
	//                 "children":[
	//                     {
	//                         "displayName":"date",
	//                         "name":"date",
	//                         "required":"true",
	//                         "type":"date"
	//                     },
	//                     {
	//                         "displayName":"time",
	//                         "name":"time",
	//                         "required":"true",
	//                         "type":"time"
	//                     },
	//                     {
	//                         "displayName":"datetime",
	//                         "name":"datetime",
	//                         "required":"true",
	//                         "type":"datetime"
	//                     }
	//                 ]
	//             },
	//             {
	//                 "displayName":"para10",
	//                 "name":"para10",
	//                 "required":"true",
	//                 "type":"list",
	//                 "children":[
	//                     {
	//                         "displayName":"value",
	//                         "name":"value",
	//                         "required":"true",
	//                         "type":"string"
	//                     }
	//                 ]
	//             },
	//             {
	//                 "displayName":"para11",
	//                 "name":"para11",
	//                 "required":"true",
	//                 "type":"dict",
	//                 "children":[
	//                     {
	//                         "displayName":"key",
	//                         "name":"key",
	//                         "required":"true",
	//                         "type":"string"
	//                     },
	//                     {
	//                         "displayName":"value",
	//                         "name":"value",
	//                         "required":"true",
	//                         "type":"string"
	//                     }
	//                 ]
	//             },
	//             {
	//                 "displayName":"para12",
	//                 "name":"para12",
	//                 "required":"true",
	//                 "type":"enum",
	//                 "children":[
	//                     {
	//                         "displayName":"PC",
	//                         "name":"1"
	//                     },
	//                     {
	//                         "displayName":"Mac",
	//                         "name":"2"
	//                     }
	//                 ]
	//             }
	//         ]
	//     }
	// }
		ConnectorParameter *string `json:"ConnectorParameter,omitempty" name:"ConnectorParameter"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEisConnectorConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListEisConnectorOperationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConnectorName")
	delete(f, "ConnectorVersion")
	if len(f) > 0 {
		return errors.New("ListEisConnectorOperationsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListEisConnectorOperationsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 连接器列表
		Operations []*EisConnectionOperation `json:"Operations,omitempty" name:"Operations"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListEisConnectorOperationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListEisConnectorOperationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

// It is highly **NOT** recommended to use this function
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
		return errors.New("ListEisConnectorsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListEisConnectorsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 连接器总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 连接器列表
		Connectors []*EisConnectorSummary `json:"Connectors,omitempty" name:"Connectors"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListEisConnectorsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListEisConnectorsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
