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

package v20190411

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AuthTestTidRequest struct {
	*tchttp.BaseRequest

	// 设备端SDK填入测试TID参数后生成的加密数据串
	Data *string `json:"Data,omitempty" name:"Data"`
}

func (r *AuthTestTidRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AuthTestTidRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AuthTestTidResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 认证结果
		Pass *bool `json:"Pass,omitempty" name:"Pass"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AuthTestTidResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AuthTestTidResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BurnTidNotifyRequest struct {
	*tchttp.BaseRequest

	// 订单编号
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// TID编号
	Tid *string `json:"Tid,omitempty" name:"Tid"`
}

func (r *BurnTidNotifyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BurnTidNotifyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BurnTidNotifyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 接收回执成功的TID
		Tid *string `json:"Tid,omitempty" name:"Tid"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BurnTidNotifyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BurnTidNotifyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeliverTidNotifyRequest struct {
	*tchttp.BaseRequest

	// 订单编号
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// TID编号
	Tid *string `json:"Tid,omitempty" name:"Tid"`
}

func (r *DeliverTidNotifyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeliverTidNotifyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeliverTidNotifyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 剩余空发数量
		RemaindCount *uint64 `json:"RemaindCount,omitempty" name:"RemaindCount"`

		// 已回执的TID编码
		Tid *string `json:"Tid,omitempty" name:"Tid"`

		// 产品公钥
		ProductKey *string `json:"ProductKey,omitempty" name:"ProductKey"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeliverTidNotifyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeliverTidNotifyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeliverTidsRequest struct {
	*tchttp.BaseRequest

	// 订单ID
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 数量，1~10
	Quantity *uint64 `json:"Quantity,omitempty" name:"Quantity"`
}

func (r *DeliverTidsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeliverTidsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeliverTidsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 空发的TID信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		TidSet []*TidKeysInfo `json:"TidSet,omitempty" name:"TidSet" list`

		// 产品公钥
		ProductKey *string `json:"ProductKey,omitempty" name:"ProductKey"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeliverTidsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeliverTidsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePermissionRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribePermissionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePermissionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePermissionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 企业用户
		EnterpriseUser *bool `json:"EnterpriseUser,omitempty" name:"EnterpriseUser"`

		// 下载控制台权限
		DownloadPermission *string `json:"DownloadPermission,omitempty" name:"DownloadPermission"`

		// 使用控制台权限
		UsePermission *string `json:"UsePermission,omitempty" name:"UsePermission"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePermissionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePermissionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DownloadTidsRequest struct {
	*tchttp.BaseRequest

	// 订单编号
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 下载数量：1~10
	Quantity *uint64 `json:"Quantity,omitempty" name:"Quantity"`
}

func (r *DownloadTidsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DownloadTidsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DownloadTidsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 下载的TID信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		TidSet []*TidKeysInfo `json:"TidSet,omitempty" name:"TidSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DownloadTidsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DownloadTidsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TidKeysInfo struct {

	// TID号码
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 公钥
	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`

	// 私钥
	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`

	// 共享密钥
	Psk *string `json:"Psk,omitempty" name:"Psk"`
}

type VerifyChipBurnInfoRequest struct {
	*tchttp.BaseRequest

	// 验证数据
	Data *string `json:"Data,omitempty" name:"Data"`
}

func (r *VerifyChipBurnInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *VerifyChipBurnInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type VerifyChipBurnInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 验证结果
		Pass *bool `json:"Pass,omitempty" name:"Pass"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VerifyChipBurnInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *VerifyChipBurnInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
