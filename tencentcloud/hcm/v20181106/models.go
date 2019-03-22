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

package v20181106

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type EvaluationRequest struct {
	*tchttp.BaseRequest

	// 图片唯一标识，一张图片一个SessionId；
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 图片数据，需要使用base64对图片的二进制数据进行编码，与url参数二者填一即可；
	Image *string `json:"Image,omitempty" name:"Image"`

	// 业务应用ID，与账号应用APPID无关，是用来方便客户管理服务的参数，新的 HcmAppId 可以在[控制台](https://console.cloud.tencent.com/hcm)【应用管理】下新建。
	HcmAppid *string `json:"HcmAppid,omitempty" name:"HcmAppid"`

	// 图片url，与Image参数二者填一即可；
	Url *string `json:"Url,omitempty" name:"Url"`
}

func (r *EvaluationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EvaluationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EvaluationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 图片唯一标识，一张图片一个SessionId；
		SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

		// 识别出的算式信息；
		Items []*Item `json:"Items,omitempty" name:"Items" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EvaluationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EvaluationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Item struct {

	// 识别的算式是否正确
	Item *string `json:"Item,omitempty" name:"Item"`

	// 识别的算式
	ItemString *string `json:"ItemString,omitempty" name:"ItemString"`

	// 识别的算式在图片上的位置信息
	ItemCoord *ItemCoord `json:"ItemCoord,omitempty" name:"ItemCoord"`

	// 推荐的答案
	Answer *string `json:"Answer,omitempty" name:"Answer"`
}

type ItemCoord struct {

	// 算式高度
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 算式宽度
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 算式图的左上角横坐标
	X *int64 `json:"X,omitempty" name:"X"`

	// 算式图的左上角纵坐标
	Y *int64 `json:"Y,omitempty" name:"Y"`
}
