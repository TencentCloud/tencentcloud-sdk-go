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

package v20210111

import (
    "encoding/json"
    "errors"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type DescribeSmsTemplateListRequest struct {
	*tchttp.BaseRequest

	// 模板 ID 数组。
	TemplateIdSet []*uint64 `json:"TemplateIdSet,omitempty" name:"TemplateIdSet" list`

	// 是否国际/港澳台短信：
	// 0：表示国内短信。
	// 1：表示国际/港澳台短信。
	International *uint64 `json:"International,omitempty" name:"International"`
}

func (r *DescribeSmsTemplateListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSmsTemplateListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateIdSet")
	delete(f, "International")
	if len(f) > 0 {
		return errors.New("DescribeSmsTemplateListRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSmsTemplateListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 获取短信模板信息响应
		DescribeTemplateStatusSet []*DescribeTemplateListStatus `json:"DescribeTemplateStatusSet,omitempty" name:"DescribeTemplateStatusSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSmsTemplateListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSmsTemplateListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTemplateListStatus struct {

	// 模板Id
	TemplateId *uint64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 是否国际/港澳台短信：
	// 0：表示国内短信。
	// 1：表示国际/港澳台短信。
	International *uint64 `json:"International,omitempty" name:"International"`

	// 申请签名状态。其中：
	// 0：表示审核通过。
	// 1：表示审核中。
	// -1：表示审核未通过或审核失败。
	StatusCode *int64 `json:"StatusCode,omitempty" name:"StatusCode"`

	// 审核回复，审核人员审核后给出的回复，通常是审核未通过的原因。
	ReviewReply *string `json:"ReviewReply,omitempty" name:"ReviewReply"`

	// 模板名称。
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 提交审核时间，UNIX 时间戳（单位：秒）。
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
}
