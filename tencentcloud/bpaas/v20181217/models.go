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

package v20181217

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type ApplyParam struct {
	// 审批流中表单唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 表单value
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value []*string `json:"Value,omitnil,omitempty" name:"Value"`

	// 表单参数描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type ApproveOpinion struct {
	// 方式 1:输入文字反馈  2:预设选项
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 审批意见
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*string `json:"Content,omitnil,omitempty" name:"Content"`
}

type ApproveUser struct {
	// 用户uin
	Uin *uint64 `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 用户类型 (1:用户  2:用户组)
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 用户描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 用户昵称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// 动态获取Scf
	// 注意：此字段可能返回 null，表示取不到有效值。
	Scf *Scf `json:"Scf,omitnil,omitempty" name:"Scf"`

	// 审批状态 （取值范围 0:待审批  1:审批通过  2:拒绝  6:其他人已审批）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproveStatus *uint64 `json:"ApproveStatus,omitnil,omitempty" name:"ApproveStatus"`

	// 审批意见
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproveMsg *string `json:"ApproveMsg,omitnil,omitempty" name:"ApproveMsg"`

	// 审批时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproveTime *string `json:"ApproveTime,omitnil,omitempty" name:"ApproveTime"`

	// 审批组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproveGroup *string `json:"ApproveGroup,omitnil,omitempty" name:"ApproveGroup"`
}

// Predefined struct for user
type GetBpaasApproveDetailRequestParams struct {
	// 审批id
	ApproveId *uint64 `json:"ApproveId,omitnil,omitempty" name:"ApproveId"`
}

type GetBpaasApproveDetailRequest struct {
	*tchttp.BaseRequest
	
	// 审批id
	ApproveId *uint64 `json:"ApproveId,omitnil,omitempty" name:"ApproveId"`
}

func (r *GetBpaasApproveDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetBpaasApproveDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApproveId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetBpaasApproveDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetBpaasApproveDetailResponseParams struct {
	// 申请人uin
	ApplyUin *uint64 `json:"ApplyUin,omitnil,omitempty" name:"ApplyUin"`

	// 申请人主账号
	ApplyOwnUin *uint64 `json:"ApplyOwnUin,omitnil,omitempty" name:"ApplyOwnUin"`

	// 申请人昵称
	ApplyUinNick *string `json:"ApplyUinNick,omitnil,omitempty" name:"ApplyUinNick"`

	// 审批流id
	BpaasId *uint64 `json:"BpaasId,omitnil,omitempty" name:"BpaasId"`

	// 审批流名称
	BpaasName *string `json:"BpaasName,omitnil,omitempty" name:"BpaasName"`

	// 申请参数
	ApplicationParams []*ApplyParam `json:"ApplicationParams,omitnil,omitempty" name:"ApplicationParams"`

	// 申请原因
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 申请时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 申请单状态
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 节点信息
	Nodes []*StatusNode `json:"Nodes,omitnil,omitempty" name:"Nodes"`

	// 正在审批的节点id
	ApprovingNodeId *string `json:"ApprovingNodeId,omitnil,omitempty" name:"ApprovingNodeId"`

	// 更新时间，时间格式：2021-12-12 10:12:10	
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetBpaasApproveDetailResponse struct {
	*tchttp.BaseResponse
	Response *GetBpaasApproveDetailResponseParams `json:"Response"`
}

func (r *GetBpaasApproveDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetBpaasApproveDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OutApproveBpaasApplicationRequestParams struct {
	// 状态  1:通过  2:拒绝
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 审批单id
	ApproveId *uint64 `json:"ApproveId,omitnil,omitempty" name:"ApproveId"`

	// 审批意见
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`
}

type OutApproveBpaasApplicationRequest struct {
	*tchttp.BaseRequest
	
	// 状态  1:通过  2:拒绝
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 审批单id
	ApproveId *uint64 `json:"ApproveId,omitnil,omitempty" name:"ApproveId"`

	// 审批意见
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`
}

func (r *OutApproveBpaasApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OutApproveBpaasApplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	delete(f, "ApproveId")
	delete(f, "Msg")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OutApproveBpaasApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OutApproveBpaasApplicationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OutApproveBpaasApplicationResponse struct {
	*tchttp.BaseResponse
	Response *OutApproveBpaasApplicationResponseParams `json:"Response"`
}

func (r *OutApproveBpaasApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OutApproveBpaasApplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Scf struct {
	// Scf函数地域id
	ScfRegion *string `json:"ScfRegion,omitnil,omitempty" name:"ScfRegion"`

	// Scf函数地域
	ScfRegionName *string `json:"ScfRegionName,omitnil,omitempty" name:"ScfRegionName"`

	// Scf函数名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScfName *string `json:"ScfName,omitnil,omitempty" name:"ScfName"`

	// Scf函数入参
	// 注意：此字段可能返回 null，表示取不到有效值。
	Params []*ScfParam `json:"Params,omitnil,omitempty" name:"Params"`
}

type ScfParam struct {
	// 参数Key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 参数类型 1用户输入 2预设参数 3表单参数
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 参数值
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// 参数描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type StatusNode struct {
	// 节点id
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 节点名称
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// 节点类型 1:审批节点 2:执行节点 3:条件节点
	NodeType *uint64 `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// 下一个节点
	NextNode *string `json:"NextNode,omitnil,omitempty" name:"NextNode"`

	// 审批意见模型
	Opinion *ApproveOpinion `json:"Opinion,omitnil,omitempty" name:"Opinion"`

	// scf函数名称
	ScfName *string `json:"ScfName,omitnil,omitempty" name:"ScfName"`

	// 状态（0：待审批，1：审批通过，2：拒绝，3：scf执行失败，4：scf执行成功）18: 外部审批中
	SubStatus *uint64 `json:"SubStatus,omitnil,omitempty" name:"SubStatus"`

	// 审批节点审批人
	ApprovedUin []*uint64 `json:"ApprovedUin,omitnil,omitempty" name:"ApprovedUin"`

	// 审批时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 审批意见信息 审批节点:审批人意见  执行节点:scf函数执行日志
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 有权限审批该节点的uin
	Users *ApproveUser `json:"Users,omitnil,omitempty" name:"Users"`

	// 是否有权限审批该节点
	IsApprove *bool `json:"IsApprove,omitnil,omitempty" name:"IsApprove"`

	// 审批id
	ApproveId *string `json:"ApproveId,omitnil,omitempty" name:"ApproveId"`

	// 审批方式 0或签 1会签
	ApproveMethod *uint64 `json:"ApproveMethod,omitnil,omitempty" name:"ApproveMethod"`

	// 审批节点审批类型，1人工审批 2自动通过 3自动决绝 4外部审批scf
	ApproveType *uint64 `json:"ApproveType,omitnil,omitempty" name:"ApproveType"`

	// 外部审批类型 scf:0或null ; CKafka:1
	CallMethod *uint64 `json:"CallMethod,omitnil,omitempty" name:"CallMethod"`

	// CKafka - 接入资源ID
	DataHubId *string `json:"DataHubId,omitnil,omitempty" name:"DataHubId"`

	// CKafka - 任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// CKafka - 地域
	CKafkaRegion *string `json:"CKafkaRegion,omitnil,omitempty" name:"CKafkaRegion"`

	// 外部审批Url
	ExternalUrl *string `json:"ExternalUrl,omitnil,omitempty" name:"ExternalUrl"`

	// 并行节点 3-4
	ParallelNodes *string `json:"ParallelNodes,omitnil,omitempty" name:"ParallelNodes"`

	// scf拒绝时返回信息
	RejectedCloudFunctionMsg *string `json:"RejectedCloudFunctionMsg,omitnil,omitempty" name:"RejectedCloudFunctionMsg"`

	// 上一个节点
	PrevNode *string `json:"PrevNode,omitnil,omitempty" name:"PrevNode"`
}