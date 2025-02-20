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

package v20200210

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type ManagePortraitRiskInput struct {
	// 请求时间戳秒
	PostTime *int64 `json:"PostTime,omitnil,omitempty" name:"PostTime"`

	// 用户公网ip（仅支持IPv4）
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`

	// 渠道号
	Channel *int64 `json:"Channel,omitnil,omitempty" name:"Channel"`
}

type ManagePortraitRiskOutput struct {
	// 返回码（0，成功，其他失败）
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 返回码对应的信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *ManagePortraitRiskValueOutput `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type ManagePortraitRiskRequestParams struct {
	// 业务入参
	BusinessSecurityData *ManagePortraitRiskInput `json:"BusinessSecurityData,omitnil,omitempty" name:"BusinessSecurityData"`
}

type ManagePortraitRiskRequest struct {
	*tchttp.BaseRequest
	
	// 业务入参
	BusinessSecurityData *ManagePortraitRiskInput `json:"BusinessSecurityData,omitnil,omitempty" name:"BusinessSecurityData"`
}

func (r *ManagePortraitRiskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManagePortraitRiskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BusinessSecurityData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ManagePortraitRiskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ManagePortraitRiskResponseParams struct {
	// 业务出参
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *ManagePortraitRiskOutput `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ManagePortraitRiskResponse struct {
	*tchttp.BaseResponse
	Response *ManagePortraitRiskResponseParams `json:"Response"`
}

func (r *ManagePortraitRiskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManagePortraitRiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ManagePortraitRiskValueOutput struct {
	// 对应的IP
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`

	// 返回风险等级, 0 - 4，0代表无风险，数值越大，风险越高
	Level *int64 `json:"Level,omitnil,omitempty" name:"Level"`
}