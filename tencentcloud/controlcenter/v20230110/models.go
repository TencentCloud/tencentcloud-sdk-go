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

package v20230110

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type BaselineConfigItem struct {
	// 账号工厂基线项唯一标识,只能包含英文字母、数字和@、,._[]-:()（）【】+=，。，长度2-128个字符。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Identifier *string `json:"Identifier,omitnil,omitempty" name:"Identifier"`

	// 账号工厂基线项配置，不同基线项配置参数不同。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Configuration *string `json:"Configuration,omitnil,omitempty" name:"Configuration"`
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