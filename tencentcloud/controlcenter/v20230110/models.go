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

package v20230110

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AccountFactoryItem struct {
	// 账号工厂基线项唯一标识，只能包含英文字母、数字和@、,._[]-:()（）【】+=，。，长度2-128个字符。
	Identifier *string `json:"Identifier,omitnil,omitempty" name:"Identifier"`

	// 基线项名称，功能项名字唯一，仅支持英文字母、数宇、汉字、符号@、＆_[]-的组合，1-25个中文或英文字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 基线项英文名称，基线项名字唯一，仅支持英文字母、数字、空格、符号@、＆_[]-的组合，1-64个英文字符。
	NameEn *string `json:"NameEn,omitnil,omitempty" name:"NameEn"`

	// 基线项权重，数值小权重越高，取值范围大于等于0。
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// 基线项是否必填，1必填，0非必填
	Required *int64 `json:"Required,omitnil,omitempty" name:"Required"`

	// 基线项依赖项，N的取值范围与该基线项依赖的其它基线项个数有关。
	DependsOn []*DependsOnItem `json:"DependsOn,omitnil,omitempty" name:"DependsOn"`

	// 基线描述，长度为2~256个英文或中文字符，默认值为空。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 基线项英文描述，长度为2~1024个英文字符，默认值为空。
	DescriptionEn *string `json:"DescriptionEn,omitnil,omitempty" name:"DescriptionEn"`

	// 基线分类，长度为2~32个英文或中文字符，不能为空。
	Classify *string `json:"Classify,omitnil,omitempty" name:"Classify"`

	// 基线英文分类，长度为2~64个英文字符，不能为空。
	ClassifyEn *string `json:"ClassifyEn,omitnil,omitempty" name:"ClassifyEn"`
}

type BaselineConfigItem struct {
	// 账号工厂基线项唯一标识,只能包含英文字母、数字和@、,._[]-:()（）【】+=，。，长度2-128个字符。
	Identifier *string `json:"Identifier,omitnil,omitempty" name:"Identifier"`

	// 账号工厂基线项配置，不同基线项配置参数不同。
	Configuration *string `json:"Configuration,omitnil,omitempty" name:"Configuration"`
}

type BaselineInfoItem struct {
	// 账号工厂基线项唯一标识，只能包含英文字母、数字和@、,._[]-:()（）【】+=，。，长度2-128个字符。
	Identifier *string `json:"Identifier,omitnil,omitempty" name:"Identifier"`

	// 账号工厂基线项配置，不同的基线项配置参数不同。
	Configuration *string `json:"Configuration,omitnil,omitempty" name:"Configuration"`

	// 基线应用的账号数量。
	ApplyCount *int64 `json:"ApplyCount,omitnil,omitempty" name:"ApplyCount"`
}

type BaselineStepTaskInfo struct {
	// 任务唯一Id，只能包含英文字母、数字，是16个字符的随机字符串。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 基线功能项唯一标识，只能包含英文字母、数字和@、,._[]-:()（）【】+=，。，长度2-128个字符。
	Identifier *string `json:"Identifier,omitnil,omitempty" name:"Identifier"`

	// 被应用基线项的成员账号uin
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 基线项应用的状态,Running表示基线项应用中,Success表示基线项应用成功,Failed表示基线项应用失败,Pending表示基线项待应用,Skipped表示基线项被跳过
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码
	ErrCode *string `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息
	ErrMessage *string `json:"ErrMessage,omitnil,omitempty" name:"ErrMessage"`

	// 基线项部署输出
	Output *string `json:"Output,omitnil,omitempty" name:"Output"`

	// 创建时间，按照ISO8601标准表示，格式为yyyy-MM-dd hh:mm:ss。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间，按照ISO8601标准表示，格式为yyyy-MM-dd hh:mm:ss。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

// Predefined struct for user
type BatchApplyAccountBaselinesRequestParams struct {
	// 成员账号uin，也是被应用基线的账号uin。
	MemberUinList []*int64 `json:"MemberUinList,omitnil,omitempty" name:"MemberUinList"`

	// 基线项配置信息列表。
	BaselineConfigItems []*BaselineConfigItem `json:"BaselineConfigItems,omitnil,omitempty" name:"BaselineConfigItems"`
}

type BatchApplyAccountBaselinesRequest struct {
	*tchttp.BaseRequest
	
	// 成员账号uin，也是被应用基线的账号uin。
	MemberUinList []*int64 `json:"MemberUinList,omitnil,omitempty" name:"MemberUinList"`

	// 基线项配置信息列表。
	BaselineConfigItems []*BaselineConfigItem `json:"BaselineConfigItems,omitnil,omitempty" name:"BaselineConfigItems"`
}

func (r *BatchApplyAccountBaselinesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchApplyAccountBaselinesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUinList")
	delete(f, "BaselineConfigItems")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchApplyAccountBaselinesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchApplyAccountBaselinesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BatchApplyAccountBaselinesResponse struct {
	*tchttp.BaseResponse
	Response *BatchApplyAccountBaselinesResponseParams `json:"Response"`
}

func (r *BatchApplyAccountBaselinesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchApplyAccountBaselinesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DependsOnItem struct {
	// 依赖项类型，只有LandingZoneSetUp或AccountFactorySetUp。LandingZoneSetUp表示landingZone的依赖项；AccountFactorySetUp表示账号工厂的依赖项
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 功能项唯一标识，只能包含英文字母、数字和@、,._[]-:()（）【】+=，。，长度2-128个字符。
	Identifier *string `json:"Identifier,omitnil,omitempty" name:"Identifier"`
}

// Predefined struct for user
type GetAccountFactoryBaselineRequestParams struct {

}

type GetAccountFactoryBaselineRequest struct {
	*tchttp.BaseRequest
	
}

func (r *GetAccountFactoryBaselineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAccountFactoryBaselineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetAccountFactoryBaselineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAccountFactoryBaselineResponseParams struct {
	// 资源所属主账号uin。
	OwnerUin *int64 `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 基线项名称，基线项名字唯一，仅支持英文字母、数宇、汉字、符号@、＆_[]-的组合，1-25个中文或英文字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 基线项配置列表。
	BaselineConfigItems []*BaselineInfoItem `json:"BaselineConfigItems,omitnil,omitempty" name:"BaselineConfigItems"`

	// 创建时间，按照ISO8601标准表示，格式为yyyy-MM-dd hh:mm:ss。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间，按照ISO8601标准表示，格式为yyyy-MM-dd hh:mm:ss。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetAccountFactoryBaselineResponse struct {
	*tchttp.BaseResponse
	Response *GetAccountFactoryBaselineResponseParams `json:"Response"`
}

func (r *GetAccountFactoryBaselineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAccountFactoryBaselineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAccountFactoryBaselineItemsRequestParams struct {
	// 返回记录最大数目,取值范围0~200。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，取值范围大于等于0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type ListAccountFactoryBaselineItemsRequest struct {
	*tchttp.BaseRequest
	
	// 返回记录最大数目,取值范围0~200。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，取值范围大于等于0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *ListAccountFactoryBaselineItemsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAccountFactoryBaselineItemsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAccountFactoryBaselineItemsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAccountFactoryBaselineItemsResponseParams struct {
	// 账号工厂基线列表。
	BaselineItems []*AccountFactoryItem `json:"BaselineItems,omitnil,omitempty" name:"BaselineItems"`

	// 总数。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAccountFactoryBaselineItemsResponse struct {
	*tchttp.BaseResponse
	Response *ListAccountFactoryBaselineItemsResponseParams `json:"Response"`
}

func (r *ListAccountFactoryBaselineItemsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAccountFactoryBaselineItemsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDeployStepTasksRequestParams struct {
	// 功能项唯一标识，只能包含英文字母、数字和@、,._[]-:()（）【】+=，。，长度2-128个字符。
	Identifier *string `json:"Identifier,omitnil,omitempty" name:"Identifier"`

	// 返回记录最大数目,取值范围0~200。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，取值范围大于等于0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type ListDeployStepTasksRequest struct {
	*tchttp.BaseRequest
	
	// 功能项唯一标识，只能包含英文字母、数字和@、,._[]-:()（）【】+=，。，长度2-128个字符。
	Identifier *string `json:"Identifier,omitnil,omitempty" name:"Identifier"`

	// 返回记录最大数目,取值范围0~200。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，取值范围大于等于0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *ListDeployStepTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDeployStepTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Identifier")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListDeployStepTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDeployStepTasksResponseParams struct {
	// 账号工厂基线功能项应用信息列表。
	BaselineDeployStepTaskList []*BaselineStepTaskInfo `json:"BaselineDeployStepTaskList,omitnil,omitempty" name:"BaselineDeployStepTaskList"`

	// 总数。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListDeployStepTasksResponse struct {
	*tchttp.BaseResponse
	Response *ListDeployStepTasksResponseParams `json:"Response"`
}

func (r *ListDeployStepTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDeployStepTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAccountFactoryBaselineRequestParams struct {
	// 基线名称，基线名字唯一，仅支持英文字母、数宇、汉字、符号@、＆_[]-的组合，1-25个中文或英文字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 基线配置，覆盖更新。可以通过controlcenter:GetAccountFactoryBaseline查询现有基线配置。可以通过controlcenter:ListAccountFactoryBaselineItems查询支持的基线列表。
	BaselineConfigItems []*BaselineConfigItem `json:"BaselineConfigItems,omitnil,omitempty" name:"BaselineConfigItems"`
}

type UpdateAccountFactoryBaselineRequest struct {
	*tchttp.BaseRequest
	
	// 基线名称，基线名字唯一，仅支持英文字母、数宇、汉字、符号@、＆_[]-的组合，1-25个中文或英文字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 基线配置，覆盖更新。可以通过controlcenter:GetAccountFactoryBaseline查询现有基线配置。可以通过controlcenter:ListAccountFactoryBaselineItems查询支持的基线列表。
	BaselineConfigItems []*BaselineConfigItem `json:"BaselineConfigItems,omitnil,omitempty" name:"BaselineConfigItems"`
}

func (r *UpdateAccountFactoryBaselineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAccountFactoryBaselineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "BaselineConfigItems")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAccountFactoryBaselineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAccountFactoryBaselineResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateAccountFactoryBaselineResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAccountFactoryBaselineResponseParams `json:"Response"`
}

func (r *UpdateAccountFactoryBaselineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAccountFactoryBaselineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}