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

package v20210129

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type LiveCodeDetail struct {

	// 活码id
	LiveCodeId *uint64 `json:"LiveCodeId,omitempty" name:"LiveCodeId"`

	// 活码名称
	LiveCodeName *string `json:"LiveCodeName,omitempty" name:"LiveCodeName"`

	// 短链url
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShortChainAddress *string `json:"ShortChainAddress,omitempty" name:"ShortChainAddress"`

	// 活码二维码
	// 注意：此字段可能返回 null，表示取不到有效值。
	LiveCodePreview *string `json:"LiveCodePreview,omitempty" name:"LiveCodePreview"`

	// 活动id
	ActivityId *int64 `json:"ActivityId,omitempty" name:"ActivityId"`

	// 活动名称
	ActivityName *string `json:"ActivityName,omitempty" name:"ActivityName"`

	// 活码状态，-1：删除，0：启用，1禁用，默认为0
	LiveCodeState *int64 `json:"LiveCodeState,omitempty" name:"LiveCodeState"`

	// 活码参数，每个活码参数都是不一样的， 这个的值对应的是字符串json类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	LiveCodeData *string `json:"LiveCodeData,omitempty" name:"LiveCodeData"`

	// 创建时间戳，单位为秒
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间戳，单位为秒
	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type QueryActivityLiveCodeListRequest struct {
	*tchttp.BaseRequest

	// 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor *string `json:"Cursor,omitempty" name:"Cursor"`

	// 返回的最大记录数，整型，最大值100，默认值50，超过最大值时取最大值
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *QueryActivityLiveCodeListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryActivityLiveCodeListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Cursor")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryActivityLiveCodeListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryActivityLiveCodeListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分页游标，再下次请求时填写以获取之后分页的记录，如果已经没有更多的数据则返回空
	// 注意：此字段可能返回 null，表示取不到有效值。
		NextCursor *string `json:"NextCursor,omitempty" name:"NextCursor"`

		// 活码列表响应参数
	// 注意：此字段可能返回 null，表示取不到有效值。
		PageData []*LiveCodeDetail `json:"PageData,omitempty" name:"PageData"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryActivityLiveCodeListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryActivityLiveCodeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
