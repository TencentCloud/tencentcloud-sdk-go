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

package v20250425

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type ManageIPPortraitRiskInput struct {
	// 用户公网ip（仅支持IPv4）
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`

	// 渠道号
	Channel *int64 `json:"Channel,omitnil,omitempty" name:"Channel"`
}

type ManageIPPortraitRiskOutput struct {
	// 返回码
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 返回消息
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 结果
	Value *ManageIPPortraitRiskValueOutput `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type ManageIPPortraitRiskRequestParams struct {
	// 请求秒级时间戳
	PostTime *int64 `json:"PostTime,omitnil,omitempty" name:"PostTime"`

	// 业务入参
	BusinessSecurityData *ManageIPPortraitRiskInput `json:"BusinessSecurityData,omitnil,omitempty" name:"BusinessSecurityData"`
}

type ManageIPPortraitRiskRequest struct {
	*tchttp.BaseRequest
	
	// 请求秒级时间戳
	PostTime *int64 `json:"PostTime,omitnil,omitempty" name:"PostTime"`

	// 业务入参
	BusinessSecurityData *ManageIPPortraitRiskInput `json:"BusinessSecurityData,omitnil,omitempty" name:"BusinessSecurityData"`
}

func (r *ManageIPPortraitRiskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManageIPPortraitRiskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PostTime")
	delete(f, "BusinessSecurityData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ManageIPPortraitRiskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ManageIPPortraitRiskResponseParams struct {
	// 出参
	Data *ManageIPPortraitRiskOutput `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ManageIPPortraitRiskResponse struct {
	*tchttp.BaseResponse
	Response *ManageIPPortraitRiskResponseParams `json:"Response"`
}

func (r *ManageIPPortraitRiskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManageIPPortraitRiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ManageIPPortraitRiskValueOutput struct {
	// 对应的IP
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`

	// 返回风险等级, 0 - 4，0代表无风险，数值越大，风险越高
	RiskScore *int64 `json:"RiskScore,omitnil,omitempty" name:"RiskScore"`

	// 风险类型
	RiskType []*int64 `json:"RiskType,omitnil,omitempty" name:"RiskType"`
}