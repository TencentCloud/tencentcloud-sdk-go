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

package v20200721

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type Conditions struct {
	// 警告条件ID
	ConditionId *uint64 `json:"ConditionId,omitnil,omitempty" name:"ConditionId"`

	// 警告级别，2:中风险，3:高风险
	Level *uint64 `json:"Level,omitnil,omitempty" name:"Level"`

	// 警告级别描述
	LevelDesc *string `json:"LevelDesc,omitnil,omitempty" name:"LevelDesc"`

	// 警告条件描述
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`
}

// Predefined struct for user
type CreateAdvisorAuthorizationRequestParams struct {

}

type CreateAdvisorAuthorizationRequest struct {
	*tchttp.BaseRequest
	
}

func (r *CreateAdvisorAuthorizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAdvisorAuthorizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAdvisorAuthorizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAdvisorAuthorizationResponseParams struct {
	// 返回信息
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAdvisorAuthorizationResponse struct {
	*tchttp.BaseResponse
	Response *CreateAdvisorAuthorizationResponseParams `json:"Response"`
}

func (r *CreateAdvisorAuthorizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAdvisorAuthorizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStrategiesRequestParams struct {

}

type DescribeStrategiesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeStrategiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStrategiesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStrategiesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStrategiesResponseParams struct {
	// 评估项列表
	Strategies []*Strategies `json:"Strategies,omitnil,omitempty" name:"Strategies"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeStrategiesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStrategiesResponseParams `json:"Response"`
}

func (r *DescribeStrategiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStrategiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskStrategyRisksRequestParams struct {
	// 评估项ID
	StrategyId *uint64 `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`

	// 返回数量,默认值为100,最大值为200
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量,默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// 任务类型
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`
}

type DescribeTaskStrategyRisksRequest struct {
	*tchttp.BaseRequest
	
	// 评估项ID
	StrategyId *uint64 `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`

	// 返回数量,默认值为100,最大值为200
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量,默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// 任务类型
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`
}

func (r *DescribeTaskStrategyRisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskStrategyRisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StrategyId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Env")
	delete(f, "TaskType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskStrategyRisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskStrategyRisksResponseParams struct {
	// 根据此配置，匹配风险实例列表（Risks）对应字段，例如:
	// {"Response":{"RequestId":"111","RiskFieldsDesc":[{"Field":"InstanceId","FieldName":"ID","FieldType":"string","FieldDict":{}},{"Field":"InstanceName","FieldName":"名称","FieldType":"string","FieldDict":{}},{"Field":"InstanceState","FieldName":"状态","FieldType":"string","FieldDict":{"LAUNCH_FAILED":"创建失败","PENDING":"创建中","REBOOTING":"重启中","RUNNING":"运行中","SHUTDOWN":"停止待销毁","STARTING":"开机中","STOPPED":"关机","STOPPING":"关机中","TERMINATING":"销毁中"}},{"Field":"Zone","FieldName":"可用区","FieldType":"string","FieldDict":{}},{"Field":"PrivateIPAddresses","FieldName":"IP地址(内)","FieldType":"stringSlice","FieldDict":{}},{"Field":"PublicIPAddresses","FieldName":"IP地址(公)","FieldType":"stringSlice","FieldDict":{}},{"Field":"Region","FieldName":"地域","FieldType":"string","FieldDict":{}},{"Field":"Tags","FieldName":"标签","FieldType":"tags","FieldDict":{}}],"RiskTotalCount":3,"Risks":"[{\"InstanceId\":\"ins-xxx1\",\"InstanceName\":\"xxx1\",\"InstanceState\":\"RUNNING\",\"PrivateIPAddresses\":[\"1.17.64.2\"],\"PublicIPAddresses\":null,\"Region\":\"ap-shanghai\",\"Tags\":null,\"Zone\":\"ap-shanghai-2\"},{\"InstanceId\":\"ins-xxx2\",\"InstanceName\":\"xxx2\",\"InstanceState\":\"RUNNING\",\"PrivateIPAddresses\":[\"1.17.64.11\"],\"PublicIPAddresses\":null,\"Region\":\"ap-shanghai\",\"Tags\":null,\"Zone\":\"ap-shanghai-2\"}]","StrategyId":9}}
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskFieldsDesc []*RiskFieldsDesc `json:"RiskFieldsDesc,omitnil,omitempty" name:"RiskFieldsDesc"`

	// 评估项ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	StrategyId *uint64 `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`

	// 风险实例个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskTotalCount *uint64 `json:"RiskTotalCount,omitnil,omitempty" name:"RiskTotalCount"`

	// 风险实例详情列表，需要json decode
	// 注意：此字段可能返回 null，表示取不到有效值。
	Risks *string `json:"Risks,omitnil,omitempty" name:"Risks"`

	// 巡检资源数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceCount *uint64 `json:"ResourceCount,omitnil,omitempty" name:"ResourceCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTaskStrategyRisksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskStrategyRisksResponseParams `json:"Response"`
}

func (r *DescribeTaskStrategyRisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskStrategyRisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KeyValue struct {
	// 键名
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 键名对应值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type RiskFieldsDesc struct {
	// 字段ID
	Field *string `json:"Field,omitnil,omitempty" name:"Field"`

	// 字段名称
	FieldName *string `json:"FieldName,omitnil,omitempty" name:"FieldName"`

	// 字段类型, 
	// string: 字符串类型，例如"aa"
	// int: 整形，例如 111
	// stringSlice : 字符串数组类型，例如["a", "b"]
	// tags: 标签类型, 例如: [{"Key":"kkk","Value":"vvv"},{"Key":"kkk2","Value":"vvv2"}]
	FieldType *string `json:"FieldType,omitnil,omitempty" name:"FieldType"`

	// 字段值对应字典
	// 注意：此字段可能返回 null，表示取不到有效值。
	FieldDict []*KeyValue `json:"FieldDict,omitnil,omitempty" name:"FieldDict"`
}

type Strategies struct {
	// 评估项ID
	StrategyId *uint64 `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`

	// 评估项名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 评估项描述
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 评估项对应产品ID
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 评估项对应产品名称
	ProductDesc *string `json:"ProductDesc,omitnil,omitempty" name:"ProductDesc"`

	// 评估项优化建议
	Repair *string `json:"Repair,omitnil,omitempty" name:"Repair"`

	// 评估项类别ID
	GroupId *uint64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 评估项类别名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 评估项风险列表
	Conditions []*Conditions `json:"Conditions,omitnil,omitempty" name:"Conditions"`
}