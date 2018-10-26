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

package v20180423

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type BindPsaTagRequest struct {
	*tchttp.BaseRequest
	// 预授权规则ID
	PsaId *string `json:"PsaId" name:"PsaId"`
	// 需要绑定的标签key
	TagKey *string `json:"TagKey" name:"TagKey"`
	// 需要绑定的标签value
	TagValue *string `json:"TagValue" name:"TagValue"`
}

func (r *BindPsaTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindPsaTagRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindPsaTagResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindPsaTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindPsaTagResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePsaRegulationRequest struct {
	*tchttp.BaseRequest
	// 规则别名
	PsaName *string `json:"PsaName" name:"PsaName"`
	// 关联的故障类型ID列表
	TaskTypeIds []*uint64 `json:"TaskTypeIds" name:"TaskTypeIds" list`
	// 维修实例上限，默认为5
	RepairLimit *uint64 `json:"RepairLimit" name:"RepairLimit"`
	// 规则备注
	PsaDescription *string `json:"PsaDescription" name:"PsaDescription"`
}

func (r *CreatePsaRegulationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePsaRegulationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePsaRegulationResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 创建的预授权规则ID
		PsaId *string `json:"PsaId" name:"PsaId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePsaRegulationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePsaRegulationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePsaRegulationRequest struct {
	*tchttp.BaseRequest
	// 预授权规则ID
	PsaId *string `json:"PsaId" name:"PsaId"`
}

func (r *DeletePsaRegulationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePsaRegulationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePsaRegulationResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePsaRegulationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePsaRegulationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePsaRegulationsRequest struct {
	*tchttp.BaseRequest
	// 数量限制
	Limit *uint64 `json:"Limit" name:"Limit"`
	// 偏移量
	Offset *uint64 `json:"Offset" name:"Offset"`
	// 规则ID过滤，支持模糊查询
	PsaIds []*string `json:"PsaIds" name:"PsaIds" list`
	// 规则别名过滤，支持模糊查询
	PsaNames []*string `json:"PsaNames" name:"PsaNames" list`
	// 标签过滤
	Tags []*Tag `json:"Tags" name:"Tags" list`
	// 排序字段，取值支持：CreateTime
	OrderField *string `json:"OrderField" name:"OrderField"`
	// 排序方式 0:递增(默认) 1:递减
	Order *uint64 `json:"Order" name:"Order"`
}

func (r *DescribePsaRegulationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePsaRegulationsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePsaRegulationsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 返回规则数量
		TotalCount *uint64 `json:"TotalCount" name:"TotalCount"`
		// 返回规则列表
		PsaRegulations []*PsaRegulation `json:"PsaRegulations" name:"PsaRegulations" list`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePsaRegulationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePsaRegulationsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRepairTaskConstantRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeRepairTaskConstantRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRepairTaskConstantRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRepairTaskConstantResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 故障类型ID与对应中文名列表
		TaskTypeSet []*TaskType `json:"TaskTypeSet" name:"TaskTypeSet" list`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRepairTaskConstantResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRepairTaskConstantResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskInfoRequest struct {
	*tchttp.BaseRequest
	// 开始位置
	Offset *uint64 `json:"Offset" name:"Offset"`
	// 数据条数
	Limit *uint64 `json:"Limit" name:"Limit"`
	// 时间过滤下限
	StartDate *string `json:"StartDate" name:"StartDate"`
	// 时间过滤上限
	EndDate *string `json:"EndDate" name:"EndDate"`
	// 任务状态ID过滤
	TaskStatus []*uint64 `json:"TaskStatus" name:"TaskStatus" list`
	// 排序字段，目前支持：CreateTime，AuthTime，EndTime
	OrderField *string `json:"OrderField" name:"OrderField"`
	// 排序方式 0:递增(默认) 1:递减
	Order *uint64 `json:"Order" name:"Order"`
	// 任务ID过滤
	TaskIds []*string `json:"TaskIds" name:"TaskIds" list`
	// 实例ID过滤
	InstanceIds []*string `json:"InstanceIds" name:"InstanceIds" list`
	// 实例别名过滤
	Aliases []*string `json:"Aliases" name:"Aliases" list`
	// 故障类型ID过滤
	TaskTypeIds []*uint64 `json:"TaskTypeIds" name:"TaskTypeIds" list`
}

func (r *DescribeTaskInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 返回任务总数量
		TotalCount *int64 `json:"TotalCount" name:"TotalCount"`
		// 任务信息列表
		TaskInfoSet []*TaskInfo `json:"TaskInfoSet" name:"TaskInfoSet" list`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskOperationLogRequest struct {
	*tchttp.BaseRequest
	// 维修任务ID
	TaskId *string `json:"TaskId" name:"TaskId"`
	// 排序字段，目前支持：OperationTime
	OrderField *string `json:"OrderField" name:"OrderField"`
	// 排序方式 0:递增(默认) 1:递减
	Order *uint64 `json:"Order" name:"Order"`
}

func (r *DescribeTaskOperationLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskOperationLogRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskOperationLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 操作日志
		TaskOperationLogSet []*TaskOperationLog `json:"TaskOperationLogSet" name:"TaskOperationLogSet" list`
		// 日志条数
		TotalCount *uint64 `json:"TotalCount" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskOperationLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskOperationLogResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyPsaRegulationRequest struct {
	*tchttp.BaseRequest
	// 预授权规则ID
	PsaId *string `json:"PsaId" name:"PsaId"`
	// 预授权规则别名
	PsaName *string `json:"PsaName" name:"PsaName"`
	// 维修中的实例上限
	RepairLimit *uint64 `json:"RepairLimit" name:"RepairLimit"`
	// 预授权规则备注
	PsaDescription *string `json:"PsaDescription" name:"PsaDescription"`
	// 预授权规则关联故障类型ID列表
	TaskTypeIds []*uint64 `json:"TaskTypeIds" name:"TaskTypeIds" list`
}

func (r *ModifyPsaRegulationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyPsaRegulationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyPsaRegulationResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPsaRegulationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyPsaRegulationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PsaRegulation struct {
	// 规则ID
	PsaId *string `json:"PsaId" name:"PsaId"`
	// 规则别名
	PsaName *string `json:"PsaName" name:"PsaName"`
	// 关联标签数量
	TagCount *uint64 `json:"TagCount" name:"TagCount"`
	// 关联实例数量
	InstanceCount *uint64 `json:"InstanceCount" name:"InstanceCount"`
	// 故障实例数量
	RepairCount *uint64 `json:"RepairCount" name:"RepairCount"`
	// 故障实例上限
	RepairLimit *uint64 `json:"RepairLimit" name:"RepairLimit"`
	// 创建时间
	CreateTime *string `json:"CreateTime" name:"CreateTime"`
	// 规则备注
	PsaDescription *string `json:"PsaDescription" name:"PsaDescription"`
	// 关联标签
	Tags []*Tag `json:"Tags" name:"Tags" list`
	// 关联故障类型id
	TaskTypeIds []*uint64 `json:"TaskTypeIds" name:"TaskTypeIds" list`
}

type RepairTaskControlRequest struct {
	*tchttp.BaseRequest
	// 维修任务ID
	TaskId *string `json:"TaskId" name:"TaskId"`
	// 操作
	Operate *string `json:"Operate" name:"Operate"`
}

func (r *RepairTaskControlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RepairTaskControlRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RepairTaskControlResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 出参TaskId是黑石异步任务ID，不同于入参TaskId字段。
	// 此字段可作为DescriptionOperationResult查询异步任务状态接口的入参，查询异步任务执行结果。
		TaskId *uint64 `json:"TaskId" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *RepairTaskControlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RepairTaskControlResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Tag struct {
	// 标签键
	TagKey *string `json:"TagKey" name:"TagKey"`
	// 标签键对应的值
	TagValues []*string `json:"TagValues" name:"TagValues" list`
}

type TaskInfo struct {
	// 任务id
	TaskId *string `json:"TaskId" name:"TaskId"`
	// 主机id
	InstanceId *string `json:"InstanceId" name:"InstanceId"`
	// 主机别名
	Alias *string `json:"Alias" name:"Alias"`
	// 故障类型id
	TaskTypeId *uint64 `json:"TaskTypeId" name:"TaskTypeId"`
	// 任务状态id
	TaskStatus *uint64 `json:"TaskStatus" name:"TaskStatus"`
	// 创建时间
	CreateTime *string `json:"CreateTime" name:"CreateTime"`
	// 授权时间
	AuthTime *string `json:"AuthTime" name:"AuthTime"`
	// 结束时间
	EndTime *string `json:"EndTime" name:"EndTime"`
	// 任务详情
	TaskDetail *string `json:"TaskDetail" name:"TaskDetail"`
	// 设备状态
	DeviceStatus *uint64 `json:"DeviceStatus" name:"DeviceStatus"`
	// 设备操作状态
	OperateStatus *uint64 `json:"OperateStatus" name:"OperateStatus"`
	// 可用区
	Zone *string `json:"Zone" name:"Zone"`
	// 地域
	Region *string `json:"Region" name:"Region"`
	// 所属网络
	VpcId *string `json:"VpcId" name:"VpcId"`
	// 所在子网
	SubnetId *string `json:"SubnetId" name:"SubnetId"`
	// 子网名
	SubnetName *string `json:"SubnetName" name:"SubnetName"`
	// VPC名
	VpcName *string `json:"VpcName" name:"VpcName"`
	// VpcCidrBlock
	VpcCidrBlock *string `json:"VpcCidrBlock" name:"VpcCidrBlock"`
	// SubnetCidrBlock
	SubnetCidrBlock *string `json:"SubnetCidrBlock" name:"SubnetCidrBlock"`
	// 公网ip
	WanIp *string `json:"WanIp" name:"WanIp"`
	// 内网IP
	LanIp *string `json:"LanIp" name:"LanIp"`
	// 管理IP
	MgtIp *string `json:"MgtIp" name:"MgtIp"`
}

type TaskOperationLog struct {
	// 操作步骤
	TaskStep *string `json:"TaskStep" name:"TaskStep"`
	// 操作人
	Operator *string `json:"Operator" name:"Operator"`
	// 操作描述
	OperationDetail *string `json:"OperationDetail" name:"OperationDetail"`
	// 操作时间
	OperationTime *string `json:"OperationTime" name:"OperationTime"`
}

type TaskType struct {
	// 故障类ID
	TypeId *uint64 `json:"TypeId" name:"TypeId"`
	// 故障类中文名
	TypeName *string `json:"TypeName" name:"TypeName"`
	// 故障类型父类
	TaskSubType *string `json:"TaskSubType" name:"TaskSubType"`
}

type UnbindPsaTagRequest struct {
	*tchttp.BaseRequest
	// 预授权规则ID
	PsaId *string `json:"PsaId" name:"PsaId"`
	// 需要解绑的标签key
	TagKey *string `json:"TagKey" name:"TagKey"`
	// 需要解绑的标签value
	TagValue *string `json:"TagValue" name:"TagValue"`
}

func (r *UnbindPsaTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnbindPsaTagRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnbindPsaTagResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnbindPsaTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnbindPsaTagResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
