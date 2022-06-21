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

package v20181011

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type ActivityInfo struct {
	// 活动使用模板id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 活动标题
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActivityTitle *string `json:"ActivityTitle,omitempty" name:"ActivityTitle"`

	// 活动描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActivityDesc *string `json:"ActivityDesc,omitempty" name:"ActivityDesc"`

	// 活动封面地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActivityCover *string `json:"ActivityCover,omitempty" name:"ActivityCover"`

	// 活动类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActivityType *string `json:"ActivityType,omitempty" name:"ActivityType"`

	// 活动id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActivityId *string `json:"ActivityId,omitempty" name:"ActivityId"`

	// 活动模板自定义配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	PersonalConfig *string `json:"PersonalConfig,omitempty" name:"PersonalConfig"`
}

// Predefined struct for user
type CheckStaffChUserRequestParams struct {
	// 员工ID
	UserId []*string `json:"UserId,omitempty" name:"UserId"`

	// 渠道状态：checkpass审核通过, checkreject审核拒绝, enableoperate启用, stopoperate停用
	OperateType *string `json:"OperateType,omitempty" name:"OperateType"`
}

type CheckStaffChUserRequest struct {
	*tchttp.BaseRequest
	
	// 员工ID
	UserId []*string `json:"UserId,omitempty" name:"UserId"`

	// 渠道状态：checkpass审核通过, checkreject审核拒绝, enableoperate启用, stopoperate停用
	OperateType *string `json:"OperateType,omitempty" name:"OperateType"`
}

func (r *CheckStaffChUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckStaffChUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "OperateType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckStaffChUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckStaffChUserResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CheckStaffChUserResponse struct {
	*tchttp.BaseResponse
	Response *CheckStaffChUserResponseParams `json:"Response"`
}

func (r *CheckStaffChUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckStaffChUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CopyActivityChannelRequestParams struct {
	// 活动ID
	ActivityId *string `json:"ActivityId,omitempty" name:"ActivityId"`

	// 来源渠道ID
	ChannelFrom *string `json:"ChannelFrom,omitempty" name:"ChannelFrom"`

	// 目的渠道id
	ChannelTo []*string `json:"ChannelTo,omitempty" name:"ChannelTo"`
}

type CopyActivityChannelRequest struct {
	*tchttp.BaseRequest
	
	// 活动ID
	ActivityId *string `json:"ActivityId,omitempty" name:"ActivityId"`

	// 来源渠道ID
	ChannelFrom *string `json:"ChannelFrom,omitempty" name:"ChannelFrom"`

	// 目的渠道id
	ChannelTo []*string `json:"ChannelTo,omitempty" name:"ChannelTo"`
}

func (r *CopyActivityChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyActivityChannelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ActivityId")
	delete(f, "ChannelFrom")
	delete(f, "ChannelTo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CopyActivityChannelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CopyActivityChannelResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CopyActivityChannelResponse struct {
	*tchttp.BaseResponse
	Response *CopyActivityChannelResponseParams `json:"Response"`
}

func (r *CopyActivityChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyActivityChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProjectRequestParams struct {
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 项目机构
	ProjectOrg *string `json:"ProjectOrg,omitempty" name:"ProjectOrg"`

	// 项目预算
	ProjectBudget *string `json:"ProjectBudget,omitempty" name:"ProjectBudget"`

	// 项目简介
	ProjectIntroduction *string `json:"ProjectIntroduction,omitempty" name:"ProjectIntroduction"`

	// 所属部门ID
	ProjectOrgId *string `json:"ProjectOrgId,omitempty" name:"ProjectOrgId"`
}

type CreateProjectRequest struct {
	*tchttp.BaseRequest
	
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 项目机构
	ProjectOrg *string `json:"ProjectOrg,omitempty" name:"ProjectOrg"`

	// 项目预算
	ProjectBudget *string `json:"ProjectBudget,omitempty" name:"ProjectBudget"`

	// 项目简介
	ProjectIntroduction *string `json:"ProjectIntroduction,omitempty" name:"ProjectIntroduction"`

	// 所属部门ID
	ProjectOrgId *string `json:"ProjectOrgId,omitempty" name:"ProjectOrgId"`
}

func (r *CreateProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectName")
	delete(f, "ProjectOrg")
	delete(f, "ProjectBudget")
	delete(f, "ProjectIntroduction")
	delete(f, "ProjectOrgId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProjectResponseParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateProjectResponse struct {
	*tchttp.BaseResponse
	Response *CreateProjectResponseParams `json:"Response"`
}

func (r *CreateProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSubProjectRequestParams struct {
	// 所属项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 子项目名称
	SubProjectName *string `json:"SubProjectName,omitempty" name:"SubProjectName"`
}

type CreateSubProjectRequest struct {
	*tchttp.BaseRequest
	
	// 所属项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 子项目名称
	SubProjectName *string `json:"SubProjectName,omitempty" name:"SubProjectName"`
}

func (r *CreateSubProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSubProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "SubProjectName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSubProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSubProjectResponseParams struct {
	// 子项目id
	SubProjectId *string `json:"SubProjectId,omitempty" name:"SubProjectId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSubProjectResponse struct {
	*tchttp.BaseResponse
	Response *CreateSubProjectResponseParams `json:"Response"`
}

func (r *CreateSubProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSubProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomerInfo struct {
	// 总活跃度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Activity *int64 `json:"Activity,omitempty" name:"Activity"`

	// 客户ID
	AudienceUserId *string `json:"AudienceUserId,omitempty" name:"AudienceUserId"`

	// 头像
	// 注意：此字段可能返回 null，表示取不到有效值。
	Avatar *string `json:"Avatar,omitempty" name:"Avatar"`

	// 最近记录城市
	// 注意：此字段可能返回 null，表示取不到有效值。
	City *string `json:"City,omitempty" name:"City"`

	// 最活跃时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastActiveTime *string `json:"LastActiveTime,omitempty" name:"LastActiveTime"`

	// 是否星标客户
	// 注意：此字段可能返回 null，表示取不到有效值。
	MarkFlag *string `json:"MarkFlag,omitempty" name:"MarkFlag"`

	// 30天活跃度
	// 注意：此字段可能返回 null，表示取不到有效值。
	MonthActive *int64 `json:"MonthActive,omitempty" name:"MonthActive"`

	// 30天推荐度
	// 注意：此字段可能返回 null，表示取不到有效值。
	MonthRecommend *int64 `json:"MonthRecommend,omitempty" name:"MonthRecommend"`

	// 手机号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 最近记录省份
	// 注意：此字段可能返回 null，表示取不到有效值。
	Province *string `json:"Province,omitempty" name:"Province"`

	// 姓名
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealName *string `json:"RealName,omitempty" name:"RealName"`

	// 员工标识 0 未关联 1 已关联
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelChannelFlag *int64 `json:"RelChannelFlag,omitempty" name:"RelChannelFlag"`

	// 性别 1男 2女
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sex *int64 `json:"Sex,omitempty" name:"Sex"`

	// 传播力（好友数）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Spread *int64 `json:"Spread,omitempty" name:"Spread"`

	// 7天活跃度
	// 注意：此字段可能返回 null，表示取不到有效值。
	WeekActive *int64 `json:"WeekActive,omitempty" name:"WeekActive"`

	// 7天推荐度
	// 注意：此字段可能返回 null，表示取不到有效值。
	WeekRecommend *int64 `json:"WeekRecommend,omitempty" name:"WeekRecommend"`

	// 微信城市
	// 注意：此字段可能返回 null，表示取不到有效值。
	WxCity *string `json:"WxCity,omitempty" name:"WxCity"`

	// 微信国家或地区
	// 注意：此字段可能返回 null，表示取不到有效值。
	WxCountry *string `json:"WxCountry,omitempty" name:"WxCountry"`

	// 微信呢称
	// 注意：此字段可能返回 null，表示取不到有效值。
	WxNickname *string `json:"WxNickname,omitempty" name:"WxNickname"`

	// 微信省份
	// 注意：此字段可能返回 null，表示取不到有效值。
	WxProvince *string `json:"WxProvince,omitempty" name:"WxProvince"`
}

// Predefined struct for user
type DeleteProjectRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type DeleteProjectRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DeleteProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProjectResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteProjectResponse struct {
	*tchttp.BaseResponse
	Response *DeleteProjectResponseParams `json:"Response"`
}

func (r *DeleteProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomerRequestParams struct {
	// 用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

type DescribeCustomerRequest struct {
	*tchttp.BaseRequest
	
	// 用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

func (r *DescribeCustomerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomerResponseParams struct {
	// 地址列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddressList []*string `json:"AddressList,omitempty" name:"AddressList"`

	// 用户id
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 头像
	// 注意：此字段可能返回 null，表示取不到有效值。
	Avatar *string `json:"Avatar,omitempty" name:"Avatar"`

	// 生日
	// 注意：此字段可能返回 null，表示取不到有效值。
	Birthday *string `json:"Birthday,omitempty" name:"Birthday"`

	// 城市
	// 注意：此字段可能返回 null，表示取不到有效值。
	City *string `json:"City,omitempty" name:"City"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 设备
	// 注意：此字段可能返回 null，表示取不到有效值。
	Device *string `json:"Device,omitempty" name:"Device"`

	// 行业
	// 注意：此字段可能返回 null，表示取不到有效值。
	Industrys []*string `json:"Industrys,omitempty" name:"Industrys"`

	// 上次登录时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastActiveTime *string `json:"LastActiveTime,omitempty" name:"LastActiveTime"`

	// 是否星标 1是 0否
	// 注意：此字段可能返回 null，表示取不到有效值。
	MarkFlag *string `json:"MarkFlag,omitempty" name:"MarkFlag"`

	// 手机型号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Model *string `json:"Model,omitempty" name:"Model"`

	// 微信openid
	// 注意：此字段可能返回 null，表示取不到有效值。
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 消费特点
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayFeature *string `json:"PayFeature,omitempty" name:"PayFeature"`

	// 手机号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 手机号码列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneList *string `json:"PhoneList,omitempty" name:"PhoneList"`

	// 最近记录省份
	// 注意：此字段可能返回 null，表示取不到有效值。
	Province *string `json:"Province,omitempty" name:"Province"`

	// 姓名
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealName *string `json:"RealName,omitempty" name:"RealName"`

	// 员工标识 0：非员工 1：员工
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelChannelFlag *string `json:"RelChannelFlag,omitempty" name:"RelChannelFlag"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 性别 1男 2女
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sex *string `json:"Sex,omitempty" name:"Sex"`

	// 最初来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceAudienceVo *string `json:"SourceAudienceVo,omitempty" name:"SourceAudienceVo"`

	// 关注公众号列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubWechats []*string `json:"SubWechats,omitempty" name:"SubWechats"`

	// 微信unionid
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnionId *string `json:"UnionId,omitempty" name:"UnionId"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 用户类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserTypes []*string `json:"UserTypes,omitempty" name:"UserTypes"`

	// 城市
	// 注意：此字段可能返回 null，表示取不到有效值。
	WxCity *string `json:"WxCity,omitempty" name:"WxCity"`

	// 国家
	// 注意：此字段可能返回 null，表示取不到有效值。
	WxCountry *string `json:"WxCountry,omitempty" name:"WxCountry"`

	// 昵称
	// 注意：此字段可能返回 null，表示取不到有效值。
	WxNickname *string `json:"WxNickname,omitempty" name:"WxNickname"`

	// 省份
	// 注意：此字段可能返回 null，表示取不到有效值。
	WxProvince *string `json:"WxProvince,omitempty" name:"WxProvince"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCustomerResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomerResponseParams `json:"Response"`
}

func (r *DescribeCustomerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomersRequestParams struct {
	// 查询类型，0.个人，1负责部门，2.指定部门
	QueryType *string `json:"QueryType,omitempty" name:"QueryType"`

	// 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 是否星级标记 1是 0否
	MarkFlag *int64 `json:"MarkFlag,omitempty" name:"MarkFlag"`

	// 客户标签，多个标签用逗号隔开
	TagIds *string `json:"TagIds,omitempty" name:"TagIds"`

	// 员工标识筛选，0：非员工，1：员工
	RelChannelFlag *string `json:"RelChannelFlag,omitempty" name:"RelChannelFlag"`

	// 必须存在手机 1是 0否
	NeedPhoneFlag *int64 `json:"NeedPhoneFlag,omitempty" name:"NeedPhoneFlag"`

	// 省份
	Province *string `json:"Province,omitempty" name:"Province"`

	// 城市
	City *string `json:"City,omitempty" name:"City"`

	// 性别 1男 2女
	Sex *string `json:"Sex,omitempty" name:"Sex"`

	// 城市
	KeyWord *string `json:"KeyWord,omitempty" name:"KeyWord"`

	// 查询开始位置
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页记录条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 子项目ID
	SubProjectId *string `json:"SubProjectId,omitempty" name:"SubProjectId"`
}

type DescribeCustomersRequest struct {
	*tchttp.BaseRequest
	
	// 查询类型，0.个人，1负责部门，2.指定部门
	QueryType *string `json:"QueryType,omitempty" name:"QueryType"`

	// 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 是否星级标记 1是 0否
	MarkFlag *int64 `json:"MarkFlag,omitempty" name:"MarkFlag"`

	// 客户标签，多个标签用逗号隔开
	TagIds *string `json:"TagIds,omitempty" name:"TagIds"`

	// 员工标识筛选，0：非员工，1：员工
	RelChannelFlag *string `json:"RelChannelFlag,omitempty" name:"RelChannelFlag"`

	// 必须存在手机 1是 0否
	NeedPhoneFlag *int64 `json:"NeedPhoneFlag,omitempty" name:"NeedPhoneFlag"`

	// 省份
	Province *string `json:"Province,omitempty" name:"Province"`

	// 城市
	City *string `json:"City,omitempty" name:"City"`

	// 性别 1男 2女
	Sex *string `json:"Sex,omitempty" name:"Sex"`

	// 城市
	KeyWord *string `json:"KeyWord,omitempty" name:"KeyWord"`

	// 查询开始位置
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页记录条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 子项目ID
	SubProjectId *string `json:"SubProjectId,omitempty" name:"SubProjectId"`
}

func (r *DescribeCustomersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "QueryType")
	delete(f, "GroupId")
	delete(f, "MarkFlag")
	delete(f, "TagIds")
	delete(f, "RelChannelFlag")
	delete(f, "NeedPhoneFlag")
	delete(f, "Province")
	delete(f, "City")
	delete(f, "Sex")
	delete(f, "KeyWord")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SubProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomersResponseParams struct {
	// 总记录条数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 数据列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserList []*CustomerInfo `json:"UserList,omitempty" name:"UserList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCustomersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomersResponseParams `json:"Response"`
}

func (r *DescribeCustomersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type DescribeProjectRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DescribeProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectResponseParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 项目预算
	ProjectBudget *float64 `json:"ProjectBudget,omitempty" name:"ProjectBudget"`

	// 项目机构
	ProjectOrg *string `json:"ProjectOrg,omitempty" name:"ProjectOrg"`

	// 项目简介
	ProjectIntroduction *string `json:"ProjectIntroduction,omitempty" name:"ProjectIntroduction"`

	// 子项目列表
	SubProjectList []*SubProjectInfo `json:"SubProjectList,omitempty" name:"SubProjectList"`

	// 项目状态
	ProjectStatus *string `json:"ProjectStatus,omitempty" name:"ProjectStatus"`

	// 项目机构Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectOrgId *string `json:"ProjectOrgId,omitempty" name:"ProjectOrgId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeProjectResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProjectResponseParams `json:"Response"`
}

func (r *DescribeProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectStockRequestParams struct {
	// 子项目id
	SubProjectId *string `json:"SubProjectId,omitempty" name:"SubProjectId"`
}

type DescribeProjectStockRequest struct {
	*tchttp.BaseRequest
	
	// 子项目id
	SubProjectId *string `json:"SubProjectId,omitempty" name:"SubProjectId"`
}

func (r *DescribeProjectStockRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectStockRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectStockRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectStockResponseParams struct {
	// 项目库存列表
	ProjectStocks []*ProjectStock `json:"ProjectStocks,omitempty" name:"ProjectStocks"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeProjectStockResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProjectStockResponseParams `json:"Response"`
}

func (r *DescribeProjectStockResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectStockResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectsRequestParams struct {
	// 页码
	PageNo *uint64 `json:"PageNo,omitempty" name:"PageNo"`

	// 页面大小
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 过滤规则
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 部门范围过滤
	Filters *Filters `json:"Filters,omitempty" name:"Filters"`

	// 项目状态, 0:编辑中 1:运营中 2:已下线 3:已删除 4:审批中
	ProjectStatus *int64 `json:"ProjectStatus,omitempty" name:"ProjectStatus"`
}

type DescribeProjectsRequest struct {
	*tchttp.BaseRequest
	
	// 页码
	PageNo *uint64 `json:"PageNo,omitempty" name:"PageNo"`

	// 页面大小
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 过滤规则
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 部门范围过滤
	Filters *Filters `json:"Filters,omitempty" name:"Filters"`

	// 项目状态, 0:编辑中 1:运营中 2:已下线 3:已删除 4:审批中
	ProjectStatus *int64 `json:"ProjectStatus,omitempty" name:"ProjectStatus"`
}

func (r *DescribeProjectsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNo")
	delete(f, "PageSize")
	delete(f, "SearchWord")
	delete(f, "Filters")
	delete(f, "ProjectStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectsResponseParams struct {
	// 项目列表
	ProjectList []*ProjectInfo `json:"ProjectList,omitempty" name:"ProjectList"`

	// 项目数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeProjectsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProjectsResponseParams `json:"Response"`
}

func (r *DescribeProjectsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceTemplateHeadersRequestParams struct {
	// 微信公众号appId
	WxAppId *string `json:"WxAppId,omitempty" name:"WxAppId"`
}

type DescribeResourceTemplateHeadersRequest struct {
	*tchttp.BaseRequest
	
	// 微信公众号appId
	WxAppId *string `json:"WxAppId,omitempty" name:"WxAppId"`
}

func (r *DescribeResourceTemplateHeadersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceTemplateHeadersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WxAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourceTemplateHeadersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceTemplateHeadersResponseParams struct {
	// 记录条数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 模板列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TmplList []*ResourceTemplateHeader `json:"TmplList,omitempty" name:"TmplList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeResourceTemplateHeadersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourceTemplateHeadersResponseParams `json:"Response"`
}

func (r *DescribeResourceTemplateHeadersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceTemplateHeadersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubProjectRequestParams struct {
	// 子项目id
	SubProjectId *string `json:"SubProjectId,omitempty" name:"SubProjectId"`
}

type DescribeSubProjectRequest struct {
	*tchttp.BaseRequest
	
	// 子项目id
	SubProjectId *string `json:"SubProjectId,omitempty" name:"SubProjectId"`
}

func (r *DescribeSubProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSubProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubProjectResponseParams struct {
	// 作品信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductInfo *ProductInfo `json:"ProductInfo,omitempty" name:"ProductInfo"`

	// 活动信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActivityInfo *ActivityInfo `json:"ActivityInfo,omitempty" name:"ActivityInfo"`

	// 分享标题
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShareTitle *string `json:"ShareTitle,omitempty" name:"ShareTitle"`

	// 分享描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShareDesc *string `json:"ShareDesc,omitempty" name:"ShareDesc"`

	// 分享图标
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShareImg *string `json:"ShareImg,omitempty" name:"ShareImg"`

	// 是否已创建策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	HasStrategy *uint64 `json:"HasStrategy,omitempty" name:"HasStrategy"`

	// 子项目状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubProjectStatus *string `json:"SubProjectStatus,omitempty" name:"SubProjectStatus"`

	// 分享公众号的appId
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShareAppId *string `json:"ShareAppId,omitempty" name:"ShareAppId"`

	// 分享公众号的wsId
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShareWsId *string `json:"ShareWsId,omitempty" name:"ShareWsId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSubProjectResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSubProjectResponseParams `json:"Response"`
}

func (r *DescribeSubProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExpireFlowRequestParams struct {
	// 工单ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`
}

type ExpireFlowRequest struct {
	*tchttp.BaseRequest
	
	// 工单ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`
}

func (r *ExpireFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExpireFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExpireFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExpireFlowResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ExpireFlowResponse struct {
	*tchttp.BaseResponse
	Response *ExpireFlowResponseParams `json:"Response"`
}

func (r *ExpireFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExpireFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filters struct {
	// 过滤类型, 0: 默认(可见部门+自创) 1: 自创 2: 指定部门(部门在可见范围内)
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 指定部门Id, 类型2使用
	DeptIds []*string `json:"DeptIds,omitempty" name:"DeptIds"`

	// 用户Id列表
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`
}

// Predefined struct for user
type ModifyProjectRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 项目预算
	ProjectBudget *string `json:"ProjectBudget,omitempty" name:"ProjectBudget"`

	// 项目机构
	ProjectOrg *string `json:"ProjectOrg,omitempty" name:"ProjectOrg"`

	// 项目简介
	ProjectIntroduction *string `json:"ProjectIntroduction,omitempty" name:"ProjectIntroduction"`

	// 项目机构Id
	ProjectOrgId *string `json:"ProjectOrgId,omitempty" name:"ProjectOrgId"`
}

type ModifyProjectRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 项目预算
	ProjectBudget *string `json:"ProjectBudget,omitempty" name:"ProjectBudget"`

	// 项目机构
	ProjectOrg *string `json:"ProjectOrg,omitempty" name:"ProjectOrg"`

	// 项目简介
	ProjectIntroduction *string `json:"ProjectIntroduction,omitempty" name:"ProjectIntroduction"`

	// 项目机构Id
	ProjectOrgId *string `json:"ProjectOrgId,omitempty" name:"ProjectOrgId"`
}

func (r *ModifyProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ProjectName")
	delete(f, "ProjectBudget")
	delete(f, "ProjectOrg")
	delete(f, "ProjectIntroduction")
	delete(f, "ProjectOrgId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProjectResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyProjectResponse struct {
	*tchttp.BaseResponse
	Response *ModifyProjectResponseParams `json:"Response"`
}

func (r *ModifyProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OffLineProjectRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type OffLineProjectRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *OffLineProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OffLineProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OffLineProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OffLineProjectResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type OffLineProjectResponse struct {
	*tchttp.BaseResponse
	Response *OffLineProjectResponseParams `json:"Response"`
}

func (r *OffLineProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OffLineProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductInfo struct {
	// 模板id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 模板主题
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductTitle *string `json:"ProductTitle,omitempty" name:"ProductTitle"`

	// 模板描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductDesc *string `json:"ProductDesc,omitempty" name:"ProductDesc"`

	// 模板封面地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductCover *string `json:"ProductCover,omitempty" name:"ProductCover"`

	// 内容作品id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 作品预览链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductUrl *string `json:"ProductUrl,omitempty" name:"ProductUrl"`

	// 作品名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
}

type ProjectInfo struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 项目所属机构
	ProjectOrg *string `json:"ProjectOrg,omitempty" name:"ProjectOrg"`

	// 项目预算
	ProjectBudget *float64 `json:"ProjectBudget,omitempty" name:"ProjectBudget"`

	// 项目状态
	ProjectStatus *string `json:"ProjectStatus,omitempty" name:"ProjectStatus"`

	// 项目创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 项目简介
	ProjectIntroduction *string `json:"ProjectIntroduction,omitempty" name:"ProjectIntroduction"`

	// 项目所属机构Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectOrgId *string `json:"ProjectOrgId,omitempty" name:"ProjectOrgId"`
}

type ProjectStock struct {
	// 奖品id
	PrizeId *string `json:"PrizeId,omitempty" name:"PrizeId"`

	// 奖品批次
	PrizeBat *uint64 `json:"PrizeBat,omitempty" name:"PrizeBat"`

	// 奖品名称
	PrizeName *string `json:"PrizeName,omitempty" name:"PrizeName"`

	// 已分配奖品数量
	UsedStock *uint64 `json:"UsedStock,omitempty" name:"UsedStock"`

	// 该奖品剩余库存数量
	RemainStock *uint64 `json:"RemainStock,omitempty" name:"RemainStock"`

	// 奖品所在奖池index
	PoolIdx *uint64 `json:"PoolIdx,omitempty" name:"PoolIdx"`

	// 奖品所在奖池名称
	PoolName *string `json:"PoolName,omitempty" name:"PoolName"`
}

// Predefined struct for user
type ReplenishProjectStockRequestParams struct {
	// 项目id
	SubProjectId *string `json:"SubProjectId,omitempty" name:"SubProjectId"`

	// 奖品id
	PrizeId *string `json:"PrizeId,omitempty" name:"PrizeId"`

	// 奖品数量
	PrizeNum *uint64 `json:"PrizeNum,omitempty" name:"PrizeNum"`

	// 奖池索引
	PoolIndex *uint64 `json:"PoolIndex,omitempty" name:"PoolIndex"`

	// 奖池名称
	PoolName *string `json:"PoolName,omitempty" name:"PoolName"`
}

type ReplenishProjectStockRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	SubProjectId *string `json:"SubProjectId,omitempty" name:"SubProjectId"`

	// 奖品id
	PrizeId *string `json:"PrizeId,omitempty" name:"PrizeId"`

	// 奖品数量
	PrizeNum *uint64 `json:"PrizeNum,omitempty" name:"PrizeNum"`

	// 奖池索引
	PoolIndex *uint64 `json:"PoolIndex,omitempty" name:"PoolIndex"`

	// 奖池名称
	PoolName *string `json:"PoolName,omitempty" name:"PoolName"`
}

func (r *ReplenishProjectStockRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplenishProjectStockRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubProjectId")
	delete(f, "PrizeId")
	delete(f, "PrizeNum")
	delete(f, "PoolIndex")
	delete(f, "PoolName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReplenishProjectStockRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplenishProjectStockResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ReplenishProjectStockResponse struct {
	*tchttp.BaseResponse
	Response *ReplenishProjectStockResponseParams `json:"Response"`
}

func (r *ReplenishProjectStockResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplenishProjectStockResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceTemplateHeader struct {
	// 模板预览区内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitempty" name:"Content"`

	// 模板预览示例
	// 注意：此字段可能返回 null，表示取不到有效值。
	Example *string `json:"Example,omitempty" name:"Example"`

	// 模板预览区域键数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyArray *string `json:"KeyArray,omitempty" name:"KeyArray"`

	// 模板id
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 模板标题
	// 注意：此字段可能返回 null，表示取不到有效值。
	Title *string `json:"Title,omitempty" name:"Title"`
}

// Predefined struct for user
type SendWxTouchTaskRequestParams struct {
	// 客户分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 去除今日已发送的客户
	DistinctFlag *bool `json:"DistinctFlag,omitempty" name:"DistinctFlag"`

	// 是否立马发送
	IsSendNow *bool `json:"IsSendNow,omitempty" name:"IsSendNow"`

	// 发送时间，一般为0
	SendDate *int64 `json:"SendDate,omitempty" name:"SendDate"`

	// 任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 微信触达类型，text, news, smallapp, tmplmsg
	WxTouchType *string `json:"WxTouchType,omitempty" name:"WxTouchType"`

	// 标题
	Title *string `json:"Title,omitempty" name:"Title"`

	// 文本内容
	Content *string `json:"Content,omitempty" name:"Content"`

	// 图文素材ID
	NewsId *string `json:"NewsId,omitempty" name:"NewsId"`

	// 小程序卡片ID
	SmallProgramId *string `json:"SmallProgramId,omitempty" name:"SmallProgramId"`

	// 模板消息ID
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 微信公众号appId
	WxAppId *string `json:"WxAppId,omitempty" name:"WxAppId"`
}

type SendWxTouchTaskRequest struct {
	*tchttp.BaseRequest
	
	// 客户分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 去除今日已发送的客户
	DistinctFlag *bool `json:"DistinctFlag,omitempty" name:"DistinctFlag"`

	// 是否立马发送
	IsSendNow *bool `json:"IsSendNow,omitempty" name:"IsSendNow"`

	// 发送时间，一般为0
	SendDate *int64 `json:"SendDate,omitempty" name:"SendDate"`

	// 任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 微信触达类型，text, news, smallapp, tmplmsg
	WxTouchType *string `json:"WxTouchType,omitempty" name:"WxTouchType"`

	// 标题
	Title *string `json:"Title,omitempty" name:"Title"`

	// 文本内容
	Content *string `json:"Content,omitempty" name:"Content"`

	// 图文素材ID
	NewsId *string `json:"NewsId,omitempty" name:"NewsId"`

	// 小程序卡片ID
	SmallProgramId *string `json:"SmallProgramId,omitempty" name:"SmallProgramId"`

	// 模板消息ID
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 微信公众号appId
	WxAppId *string `json:"WxAppId,omitempty" name:"WxAppId"`
}

func (r *SendWxTouchTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendWxTouchTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "DistinctFlag")
	delete(f, "IsSendNow")
	delete(f, "SendDate")
	delete(f, "TaskName")
	delete(f, "WxTouchType")
	delete(f, "Title")
	delete(f, "Content")
	delete(f, "NewsId")
	delete(f, "SmallProgramId")
	delete(f, "TemplateId")
	delete(f, "WxAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendWxTouchTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendWxTouchTaskResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SendWxTouchTaskResponse struct {
	*tchttp.BaseResponse
	Response *SendWxTouchTaskResponseParams `json:"Response"`
}

func (r *SendWxTouchTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendWxTouchTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubProjectInfo struct {
	// 子项目id
	SubProjectId *string `json:"SubProjectId,omitempty" name:"SubProjectId"`

	// 子项目名称
	SubProjectName *string `json:"SubProjectName,omitempty" name:"SubProjectName"`

	// 子项目状态
	SubProjectStatus *string `json:"SubProjectStatus,omitempty" name:"SubProjectStatus"`
}