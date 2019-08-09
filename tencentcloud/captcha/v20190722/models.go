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

package v20190722

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type DescribeCaptchaResultRequest struct {
	*tchttp.BaseRequest

	// 验证码类型，9：滑块验证码
	CaptchaType *uint64 `json:"CaptchaType,omitempty" name:"CaptchaType"`

	// 验证码返回给用户的票据
	Ticket *string `json:"Ticket,omitempty" name:"Ticket"`

	// 用户操作来源的外网 IP
	UserIp *string `json:"UserIp,omitempty" name:"UserIp"`

	// 验证票据需要的随机字符串
	Randstr *string `json:"Randstr,omitempty" name:"Randstr"`

	// 验证码应用ID
	CaptchaAppId *uint64 `json:"CaptchaAppId,omitempty" name:"CaptchaAppId"`

	// 用于服务器端校验验证码票据的验证密钥，请妥善保密，请勿泄露给第三方
	AppSecretKey *string `json:"AppSecretKey,omitempty" name:"AppSecretKey"`

	// 业务 ID，网站或应用在多个业务中使用此服务，通过此 ID 区分统计数据
	BusinessId *uint64 `json:"BusinessId,omitempty" name:"BusinessId"`

	// 场景 ID，网站或应用的业务下有多个场景使用此服务，通过此 ID 区分统计数据
	SceneId *uint64 `json:"SceneId,omitempty" name:"SceneId"`

	// mac 地址或设备唯一标识
	MacAddress *string `json:"MacAddress,omitempty" name:"MacAddress"`

	// 手机设备号
	Imei *string `json:"Imei,omitempty" name:"Imei"`
}

func (r *DescribeCaptchaResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCaptchaResultRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCaptchaResultResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 1:验证成功，0:验证失败，100:AppSecretKey参数校验错误
		CaptchaCode *int64 `json:"CaptchaCode,omitempty" name:"CaptchaCode"`

		// 状态描述及验证错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		CaptchaMsg *string `json:"CaptchaMsg,omitempty" name:"CaptchaMsg"`

		// [0,100]，恶意等级
	// 注意：此字段可能返回 null，表示取不到有效值。
		EvilLevel *int64 `json:"EvilLevel,omitempty" name:"EvilLevel"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCaptchaResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCaptchaResultResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
