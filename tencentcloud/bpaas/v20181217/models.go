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

package v20181217

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type ApplyParam struct {
	// 审批流中表单唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 表单value
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value []*string `json:"Value,omitempty" name:"Value"`

	// 表单参数描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`
}

type ApproveOpinion struct {
	// 方式 1:输入文字反馈  2:预设选项
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 审批意见
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*string `json:"Content,omitempty" name:"Content"`
}

type ApproveUser struct {
	// 用户uin
	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`

	// 用户类型 (1:用户  2:用户组)
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 用户描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// 用户昵称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Nick *string `json:"Nick,omitempty" name:"Nick"`

	// 动态获取Scf
	// 注意：此字段可能返回 null，表示取不到有效值。
	Scf *Scf `json:"Scf,omitempty" name:"Scf"`
}

// Predefined struct for user
type GetBpaasApproveDetailRequestParams struct {
	// 审批id
	ApproveId *uint64 `json:"ApproveId,omitempty" name:"ApproveId"`
}

type GetBpaasApproveDetailRequest struct {
	*tchttp.BaseRequest
	
	// 审批id
	ApproveId *uint64 `json:"ApproveId,omitempty" name:"ApproveId"`
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
	ApplyUin *uint64 `json:"ApplyUin,omitempty" name:"ApplyUin"`

	// 申请人主账号
	ApplyOwnUin *uint64 `json:"ApplyOwnUin,omitempty" name:"ApplyOwnUin"`

	// 申请人昵称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplyUinNick *string `json:"ApplyUinNick,omitempty" name:"ApplyUinNick"`

	// 审批流id
	BpaasId *uint64 `json:"BpaasId,omitempty" name:"BpaasId"`

	// 审批流名称
	BpaasName *string `json:"BpaasName,omitempty" name:"BpaasName"`

	// 申请参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationParams []*ApplyParam `json:"ApplicationParams,omitempty" name:"ApplicationParams"`

	// 申请原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// 申请时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 申请单状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 节点信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Nodes []*StatusNode `json:"Nodes,omitempty" name:"Nodes"`

	// 正在审批的节点id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApprovingNodeId *string `json:"ApprovingNodeId,omitempty" name:"ApprovingNodeId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 审批单id
	ApproveId *uint64 `json:"ApproveId,omitempty" name:"ApproveId"`

	// 审批意见
	Msg *string `json:"Msg,omitempty" name:"Msg"`
}

type OutApproveBpaasApplicationRequest struct {
	*tchttp.BaseRequest
	
	// 状态  1:通过  2:拒绝
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 审批单id
	ApproveId *uint64 `json:"ApproveId,omitempty" name:"ApproveId"`

	// 审批意见
	Msg *string `json:"Msg,omitempty" name:"Msg"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	ScfRegion *string `json:"ScfRegion,omitempty" name:"ScfRegion"`

	// Scf函数地域
	ScfRegionName *string `json:"ScfRegionName,omitempty" name:"ScfRegionName"`

	// Scf函数名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScfName *string `json:"ScfName,omitempty" name:"ScfName"`

	// Scf函数入参
	// 注意：此字段可能返回 null，表示取不到有效值。
	Params []*ScfParam `json:"Params,omitempty" name:"Params"`
}

type ScfParam struct {
	// 参数Key
	Key *string `json:"Key,omitempty" name:"Key"`

	// 参数类型 1用户输入 2预设参数 3表单参数
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 参数值
	Values []*string `json:"Values,omitempty" name:"Values"`

	// 参数描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`
}

type StatusNode struct {
	// 节点id
	NodeId *string `json:"NodeId,omitempty" name:"NodeId"`

	// 节点名称
	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`

	// 节点类型 1:审批节点 2:执行节点 3:条件节点
	NodeType *uint64 `json:"NodeType,omitempty" name:"NodeType"`

	// 下一个节点
	NextNode *string `json:"NextNode,omitempty" name:"NextNode"`

	// 审批意见模型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Opinion *ApproveOpinion `json:"Opinion,omitempty" name:"Opinion"`

	// scf函数名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScfName *string `json:"ScfName,omitempty" name:"ScfName"`

	// 状态（0：待审批，1：审批通过，2：拒绝，3：scf执行失败，4：scf执行成功）18: 外部审批中
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubStatus *uint64 `json:"SubStatus,omitempty" name:"SubStatus"`

	// 审批节点审批人
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApprovedUin []*uint64 `json:"ApprovedUin,omitempty" name:"ApprovedUin"`

	// 审批时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 审批意见信息 审批节点:审批人意见  执行节点:scf函数执行日志
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 有权限审批该节点的uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Users *ApproveUser `json:"Users,omitempty" name:"Users"`

	// 是否有权限审批该节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsApprove *bool `json:"IsApprove,omitempty" name:"IsApprove"`

	// 审批id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproveId *string `json:"ApproveId,omitempty" name:"ApproveId"`

	// 审批方式 0或签 1会签
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproveMethod *uint64 `json:"ApproveMethod,omitempty" name:"ApproveMethod"`

	// 审批节点审批类型，1人工审批 2自动通过 3自动决绝 4外部审批scf
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproveType *uint64 `json:"ApproveType,omitempty" name:"ApproveType"`

	// 外部审批类型 scf:0或null ; CKafka:1
	// 注意：此字段可能返回 null，表示取不到有效值。
	CallMethod *uint64 `json:"CallMethod,omitempty" name:"CallMethod"`

	// CKafka - 接入资源ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataHubId *string `json:"DataHubId,omitempty" name:"DataHubId"`

	// CKafka - 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// CKafka - 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	CKafkaRegion *string `json:"CKafkaRegion,omitempty" name:"CKafkaRegion"`

	// 外部审批Url
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalUrl *string `json:"ExternalUrl,omitempty" name:"ExternalUrl"`

	// 并行节点 3-4
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParallelNodes *string `json:"ParallelNodes,omitempty" name:"ParallelNodes"`
}