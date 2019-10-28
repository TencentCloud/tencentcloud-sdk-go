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

package v20181225

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AcceptOrganizationInvitationRequest struct {
	*tchttp.BaseRequest

	// 邀请ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *AcceptOrganizationInvitationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AcceptOrganizationInvitationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AcceptOrganizationInvitationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AcceptOrganizationInvitationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AcceptOrganizationInvitationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddOrganizationNodeRequest struct {
	*tchttp.BaseRequest

	// 父组织单元ID
	ParentNodeId *uint64 `json:"ParentNodeId,omitempty" name:"ParentNodeId"`

	// 组织单元名字
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *AddOrganizationNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddOrganizationNodeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddOrganizationNodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 组织单元ID
		NodeId *uint64 `json:"NodeId,omitempty" name:"NodeId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddOrganizationNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddOrganizationNodeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CancelOrganizationInvitationRequest struct {
	*tchttp.BaseRequest

	// 邀请ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *CancelOrganizationInvitationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CancelOrganizationInvitationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CancelOrganizationInvitationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelOrganizationInvitationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CancelOrganizationInvitationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateOrganizationRequest struct {
	*tchttp.BaseRequest

	// 组织类型（目前固定为1）
	OrgType *uint64 `json:"OrgType,omitempty" name:"OrgType"`
}

func (r *CreateOrganizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateOrganizationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateOrganizationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 企业组织ID
		OrgId *uint64 `json:"OrgId,omitempty" name:"OrgId"`

		// 创建者昵称
		Nickname *string `json:"Nickname,omitempty" name:"Nickname"`

		// 创建者邮箱
		Mail *string `json:"Mail,omitempty" name:"Mail"`

		// 组织类型
		OrgType *uint64 `json:"OrgType,omitempty" name:"OrgType"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOrganizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateOrganizationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteOrganizationMemberFromNodeRequest struct {
	*tchttp.BaseRequest

	// 被删除成员UIN
	MemberUin *uint64 `json:"MemberUin,omitempty" name:"MemberUin"`

	// 组织单元ID
	NodeId *uint64 `json:"NodeId,omitempty" name:"NodeId"`
}

func (r *DeleteOrganizationMemberFromNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteOrganizationMemberFromNodeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteOrganizationMemberFromNodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteOrganizationMemberFromNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteOrganizationMemberFromNodeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteOrganizationMembersRequest struct {
	*tchttp.BaseRequest

	// 被删除成员的UIN列表
	Uins []*uint64 `json:"Uins,omitempty" name:"Uins" list`
}

func (r *DeleteOrganizationMembersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteOrganizationMembersRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteOrganizationMembersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteOrganizationMembersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteOrganizationMembersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteOrganizationNodesRequest struct {
	*tchttp.BaseRequest

	// 组织单元ID列表
	NodeIds []*uint64 `json:"NodeIds,omitempty" name:"NodeIds" list`
}

func (r *DeleteOrganizationNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteOrganizationNodesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteOrganizationNodesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteOrganizationNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteOrganizationNodesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteOrganizationRequest struct {
	*tchttp.BaseRequest
}

func (r *DeleteOrganizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteOrganizationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteOrganizationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteOrganizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteOrganizationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DenyOrganizationInvitationRequest struct {
	*tchttp.BaseRequest

	// 邀请ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DenyOrganizationInvitationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DenyOrganizationInvitationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DenyOrganizationInvitationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DenyOrganizationInvitationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DenyOrganizationInvitationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetOrganizationMemberRequest struct {
	*tchttp.BaseRequest

	// 组织成员UIN
	MemberUin *uint64 `json:"MemberUin,omitempty" name:"MemberUin"`
}

func (r *GetOrganizationMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetOrganizationMemberRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetOrganizationMemberResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 组织成员UIN
		Uin *uint64 `json:"Uin,omitempty" name:"Uin"`

		// 组织成员名称
		Name *string `json:"Name,omitempty" name:"Name"`

		// 备注
		Remark *string `json:"Remark,omitempty" name:"Remark"`

		// 加入时间
		JoinTime *string `json:"JoinTime,omitempty" name:"JoinTime"`

		// 组织单元ID
		NodeId *uint64 `json:"NodeId,omitempty" name:"NodeId"`

		// 组织单元名称
		NodeName *string `json:"NodeName,omitempty" name:"NodeName"`

		// 父组织单元ID
		ParentNodeId *uint64 `json:"ParentNodeId,omitempty" name:"ParentNodeId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetOrganizationMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetOrganizationMemberResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetOrganizationRequest struct {
	*tchttp.BaseRequest
}

func (r *GetOrganizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetOrganizationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetOrganizationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 企业组织ID
		OrgId *uint64 `json:"OrgId,omitempty" name:"OrgId"`

		// 创建者UIN
		HostUin *uint64 `json:"HostUin,omitempty" name:"HostUin"`

		// 创建者昵称
		Nickname *string `json:"Nickname,omitempty" name:"Nickname"`

		// 创建者邮箱
		Mail *string `json:"Mail,omitempty" name:"Mail"`

		// 企业组织类型
		OrgType *uint64 `json:"OrgType,omitempty" name:"OrgType"`

		// 是否为空
		IsEmpty *uint64 `json:"IsEmpty,omitempty" name:"IsEmpty"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetOrganizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetOrganizationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListOrganizationInvitationsRequest struct {
	*tchttp.BaseRequest

	// 是否被邀请。1：被邀请，0：发出的邀请
	Invited *uint64 `json:"Invited,omitempty" name:"Invited"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListOrganizationInvitationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListOrganizationInvitationsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListOrganizationInvitationsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 邀请信息列表
		Invitations []*OrgInvitation `json:"Invitations,omitempty" name:"Invitations" list`

		// 总数目
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListOrganizationInvitationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListOrganizationInvitationsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListOrganizationMembersRequest struct {
	*tchttp.BaseRequest

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListOrganizationMembersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListOrganizationMembersRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListOrganizationMembersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 成员列表
		Members []*OrgMember `json:"Members,omitempty" name:"Members" list`

		// 总数目
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListOrganizationMembersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListOrganizationMembersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListOrganizationNodeMembersRequest struct {
	*tchttp.BaseRequest

	// 企业组织单元ID
	NodeId *uint64 `json:"NodeId,omitempty" name:"NodeId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListOrganizationNodeMembersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListOrganizationNodeMembersRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListOrganizationNodeMembersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数目
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 成员列表
		Members []*OrgMember `json:"Members,omitempty" name:"Members" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListOrganizationNodeMembersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListOrganizationNodeMembersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListOrganizationNodesRequest struct {
	*tchttp.BaseRequest
}

func (r *ListOrganizationNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListOrganizationNodesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListOrganizationNodesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 企业组织单元列表
		Nodes []*OrgNode `json:"Nodes,omitempty" name:"Nodes" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListOrganizationNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListOrganizationNodesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MoveOrganizationMembersToNodeRequest struct {
	*tchttp.BaseRequest

	// 组织单元ID
	NodeId *uint64 `json:"NodeId,omitempty" name:"NodeId"`

	// 成员UIN列表
	Uins []*uint64 `json:"Uins,omitempty" name:"Uins" list`
}

func (r *MoveOrganizationMembersToNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MoveOrganizationMembersToNodeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MoveOrganizationMembersToNodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MoveOrganizationMembersToNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MoveOrganizationMembersToNodeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OrgInvitation struct {

	// 邀请ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 被邀请UIN
	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`

	// 创建者UIN
	HostUin *uint64 `json:"HostUin,omitempty" name:"HostUin"`

	// 创建者名称
	HostName *string `json:"HostName,omitempty" name:"HostName"`

	// 创建者邮箱
	HostMail *string `json:"HostMail,omitempty" name:"HostMail"`

	// 邀请状态。-1：已过期，0：正常，1：已接受，2：已失效，3：已取消
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 企业组织类型
	OrgType *uint64 `json:"OrgType,omitempty" name:"OrgType"`

	// 邀请时间
	InviteTime *string `json:"InviteTime,omitempty" name:"InviteTime"`

	// 过期时间
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
}

type OrgMember struct {

	// UIN
	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 加入时间
	JoinTime *string `json:"JoinTime,omitempty" name:"JoinTime"`
}

type OrgNode struct {

	// 组织单元ID
	NodeId *uint64 `json:"NodeId,omitempty" name:"NodeId"`

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 父单元ID
	ParentNodeId *uint64 `json:"ParentNodeId,omitempty" name:"ParentNodeId"`

	// 成员数量
	MemberCount *uint64 `json:"MemberCount,omitempty" name:"MemberCount"`
}

type QuitOrganizationRequest struct {
	*tchttp.BaseRequest

	// 企业组织ID
	OrgId *uint64 `json:"OrgId,omitempty" name:"OrgId"`
}

func (r *QuitOrganizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QuitOrganizationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QuitOrganizationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QuitOrganizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QuitOrganizationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SendOrganizationInvitationRequest struct {
	*tchttp.BaseRequest

	// 被邀请账户UIN
	InviteUin *uint64 `json:"InviteUin,omitempty" name:"InviteUin"`

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *SendOrganizationInvitationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SendOrganizationInvitationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SendOrganizationInvitationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendOrganizationInvitationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SendOrganizationInvitationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateOrganizationMemberRequest struct {
	*tchttp.BaseRequest

	// 成员UIN
	MemberUin *uint64 `json:"MemberUin,omitempty" name:"MemberUin"`

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *UpdateOrganizationMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateOrganizationMemberRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateOrganizationMemberResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateOrganizationMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateOrganizationMemberResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateOrganizationNodeRequest struct {
	*tchttp.BaseRequest

	// 企业组织单元ID
	NodeId *uint64 `json:"NodeId,omitempty" name:"NodeId"`

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 父单元ID
	ParentNodeId *uint64 `json:"ParentNodeId,omitempty" name:"ParentNodeId"`
}

func (r *UpdateOrganizationNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateOrganizationNodeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateOrganizationNodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateOrganizationNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateOrganizationNodeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
