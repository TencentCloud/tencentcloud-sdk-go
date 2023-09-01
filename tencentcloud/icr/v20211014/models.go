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

package v20211014

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type GetIndustryV1HomeMembersReqPayload struct {
	// 用户ID
	ID *string `json:"ID,omitnil" name:"ID"`
}

// Predefined struct for user
type GetIndustryV1HomeMembersRequestParams struct {
	// 无
	Payload *GetIndustryV1HomeMembersReqPayload `json:"Payload,omitnil" name:"Payload"`

	// 无
	Metadata *ReqMetadata `json:"Metadata,omitnil" name:"Metadata"`
}

type GetIndustryV1HomeMembersRequest struct {
	*tchttp.BaseRequest
	
	// 无
	Payload *GetIndustryV1HomeMembersReqPayload `json:"Payload,omitnil" name:"Payload"`

	// 无
	Metadata *ReqMetadata `json:"Metadata,omitnil" name:"Metadata"`
}

func (r *GetIndustryV1HomeMembersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetIndustryV1HomeMembersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Payload")
	delete(f, "Metadata")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetIndustryV1HomeMembersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetIndustryV1HomeMembersRespData struct {
	// 修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EditTime *int64 `json:"EditTime,omitnil" name:"EditTime"`

	// 功能列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	FeatureList *GetIndustryV1HomeMembersRespFeature `json:"FeatureList,omitnil" name:"FeatureList"`

	// 用户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ID *string `json:"ID,omitnil" name:"ID"`

	// 用户行业分类
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndustryType *string `json:"IndustryType,omitnil" name:"IndustryType"`

	// 子用户数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemberNum *int64 `json:"MemberNum,omitnil" name:"MemberNum"`

	// 机器人列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductList *GetIndustryV1HomeMembersRespProduct `json:"ProductList,omitnil" name:"ProductList"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 是否有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 功能列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TypeList *GetIndustryV1HomeMembersRespType `json:"TypeList,omitnil" name:"TypeList"`

	// 用户账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserAccount *string `json:"UserAccount,omitnil" name:"UserAccount"`
}

type GetIndustryV1HomeMembersRespFeature struct {
	// 功能名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FeatureName *string `json:"FeatureName,omitnil" name:"FeatureName"`

	// 功能ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ID *string `json:"ID,omitnil" name:"ID"`
}

type GetIndustryV1HomeMembersRespIndustry struct {
	// 行业ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ID *string `json:"ID,omitnil" name:"ID"`

	// 行业名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndustryName *string `json:"IndustryName,omitnil" name:"IndustryName"`
}

type GetIndustryV1HomeMembersRespPayload struct {
	// 用户级别
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountLevel *string `json:"AccountLevel,omitnil" name:"AccountLevel"`

	// 用户列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataList []*GetIndustryV1HomeMembersRespData `json:"DataList,omitnil" name:"DataList"`

	// 每页数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 分页偏移量，从0开始
	// 注意：此字段可能返回 null，表示取不到有效值。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 用户总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil" name:"Total"`
}

type GetIndustryV1HomeMembersRespProduct struct {
	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 编辑时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EditTime *string `json:"EditTime,omitnil" name:"EditTime"`

	// 机器人ID（AppKey信息）
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppKey *string `json:"AppKey,omitnil" name:"AppKey"`

	// 机器人图标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Image *string `json:"Image,omitnil" name:"Image"`

	// 行业信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Industry []*GetIndustryV1HomeMembersRespIndustry `json:"Industry,omitnil" name:"Industry"`

	// 操作员列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperatorList *string `json:"OperatorList,omitnil" name:"OperatorList"`

	// 机器人名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductName *string `json:"ProductName,omitnil" name:"ProductName"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 模板列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateList *string `json:"TemplateList,omitnil" name:"TemplateList"`
}

type GetIndustryV1HomeMembersRespType struct {
	// 类型ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 类型名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TypeName *string `json:"TypeName,omitnil" name:"TypeName"`
}

// Predefined struct for user
type GetIndustryV1HomeMembersResponseParams struct {
	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Metadata *RspMetadata `json:"Metadata,omitnil" name:"Metadata"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Payload *GetIndustryV1HomeMembersRespPayload `json:"Payload,omitnil" name:"Payload"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetIndustryV1HomeMembersResponse struct {
	*tchttp.BaseResponse
	Response *GetIndustryV1HomeMembersResponseParams `json:"Response"`
}

func (r *GetIndustryV1HomeMembersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetIndustryV1HomeMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReqMetadata struct {
	// 渠道
	ChannelID *string `json:"ChannelID,omitnil" name:"ChannelID"`

	// 无
	BusinessName *string `json:"BusinessName,omitnil" name:"BusinessName"`

	// 无
	GUID *string `json:"GUID,omitnil" name:"GUID"`

	// 无
	AppKey *string `json:"AppKey,omitnil" name:"AppKey"`

	// 位置定位服务
	LBS *ReqMetadataLBS `json:"LBS,omitnil" name:"LBS"`

	// 透传字段
	Vagrants []*ReqMetadataVagrant `json:"Vagrants,omitnil" name:"Vagrants"`
}

type ReqMetadataLBS struct {
	// 纬度
	Latitude *float64 `json:"Latitude,omitnil" name:"Latitude"`

	// 经度
	Longitude *float64 `json:"Longitude,omitnil" name:"Longitude"`
}

type ReqMetadataVagrant struct {
	// 无
	Key *string `json:"Key,omitnil" name:"Key"`

	// 无
	Value *string `json:"Value,omitnil" name:"Value"`
}

type RspMetadata struct {
	// 无
	Code *int64 `json:"Code,omitnil" name:"Code"`

	// 无
	Message *string `json:"Message,omitnil" name:"Message"`

	// 无
	SessionID *string `json:"SessionID,omitnil" name:"SessionID"`

	// 无
	SessionDelta *string `json:"SessionDelta,omitnil" name:"SessionDelta"`
}