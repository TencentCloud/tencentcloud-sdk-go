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

package v20190926

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Data struct {

	// 操作时间戳，单位秒
	PostTime *uint64 `json:"PostTime,omitempty" name:"PostTime"`

	// 用户ID 
	// accountType不同对应不同的用户ID。如果是QQ或微信用户则填入对应的openId
	Uid *string `json:"Uid,omitempty" name:"Uid"`

	// 操作来源的外网IP
	UserIp *string `json:"UserIp,omitempty" name:"UserIp"`

	// 0~100：营销价值评分，分值越高，价值越大
	// [0,50]低价值
	// [50,70]价值一般
	// [70,100]高价值
	ValueScore *uint64 `json:"ValueScore,omitempty" name:"ValueScore"`
}

type MarketingValueJudgementRequest struct {
	*tchttp.BaseRequest

	// 手机账号类型填写4
	AccountType *uint64 `json:"AccountType,omitempty" name:"AccountType"`

	// 填写手机号码，如15317537488
	Uid *string `json:"Uid,omitempty" name:"Uid"`

	// 用户请求时的客户端外网IP
	UserIp *string `json:"UserIp,omitempty" name:"UserIp"`

	// 用户操作时间戳，单位秒（格林威治时间精确到秒，如1501590972）
	PostTime *uint64 `json:"PostTime,omitempty" name:"PostTime"`

	// 用户设备号imei/idfa(建议填写)
	Imei *string `json:"Imei,omitempty" name:"Imei"`

	// 活动链接(建议填写)
	Referer *string `json:"Referer,omitempty" name:"Referer"`
}

func (r *MarketingValueJudgementRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MarketingValueJudgementRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MarketingValueJudgementResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回数据
		Data *Data `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MarketingValueJudgementResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MarketingValueJudgementResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
