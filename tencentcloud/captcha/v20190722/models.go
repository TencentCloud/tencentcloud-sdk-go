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

		// 1	OK	验证通过
	// 6	user code len error	验证码长度不匹配
	// 7	captcha no match	验证码答案不匹配/Randstr参数不匹配
	// 8	verify timeout	验证码签名超时
	// 9	Sequnce repeat	验证码签名重放
	// 10	Sequnce invalid	验证码签名序列
	// 11	Cookie invalid	验证码cooking信息不合法
	// 12	sig len error	签名长度错误
	// 13	verify ip no match	ip不匹配
	// 15	decrypt fail	验证码签名解密失败
	// 16	appid no match	验证码强校验appid错误
	// 17	cmd no much	验证码系统命令不匹配
	// 18	uin no match	号码不匹配
	// 19	seq redirect	重定向验证
	// 20	opt no vcode	操作使用pt免验证码校验错误
	// 21	diff	差别，验证错误
	// 22	captcha type not match	验证码类型与拉取时不一致
	// 23	verify type error	验证类型错误
	// 24	invalid pkg	非法请求包
	// 25	bad visitor	策略拦截
	// 26	system busy	系统内部错误
	// 100	param err	appsecretkey 参数校验错误
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
