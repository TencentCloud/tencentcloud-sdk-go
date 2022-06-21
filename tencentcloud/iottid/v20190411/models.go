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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type AuthTestTidRequestParams struct {
	// 设备端SDK填入测试TID参数后生成的加密数据串
	Data *string `json:"Data,omitempty" name:"Data"`
}

type AuthTestTidRequest struct {
	*tchttp.BaseRequest
	
	// 设备端SDK填入测试TID参数后生成的加密数据串
	Data *string `json:"Data,omitempty" name:"Data"`
}

func (r *AuthTestTidRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AuthTestTidRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Data")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AuthTestTidRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AuthTestTidResponseParams struct {
	// 认证结果
	Pass *bool `json:"Pass,omitempty" name:"Pass"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AuthTestTidResponse struct {
	*tchttp.BaseResponse
	Response *AuthTestTidResponseParams `json:"Response"`
}

func (r *AuthTestTidResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AuthTestTidResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BurnTidNotifyRequestParams struct {
	// 订单编号
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// TID编号
	Tid *string `json:"Tid,omitempty" name:"Tid"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BurnTidNotifyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrderId")
	delete(f, "Tid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BurnTidNotifyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BurnTidNotifyResponseParams struct {
	// 接收回执成功的TID
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BurnTidNotifyResponse struct {
	*tchttp.BaseResponse
	Response *BurnTidNotifyResponseParams `json:"Response"`
}

func (r *BurnTidNotifyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BurnTidNotifyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeliverTidNotifyRequestParams struct {
	// 订单编号
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// TID编号
	Tid *string `json:"Tid,omitempty" name:"Tid"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeliverTidNotifyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrderId")
	delete(f, "Tid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeliverTidNotifyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeliverTidNotifyResponseParams struct {
	// 剩余空发数量
	RemaindCount *uint64 `json:"RemaindCount,omitempty" name:"RemaindCount"`

	// 已回执的TID编码
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 产品公钥
	ProductKey *string `json:"ProductKey,omitempty" name:"ProductKey"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeliverTidNotifyResponse struct {
	*tchttp.BaseResponse
	Response *DeliverTidNotifyResponseParams `json:"Response"`
}

func (r *DeliverTidNotifyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeliverTidNotifyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeliverTidsRequestParams struct {
	// 订单ID
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 数量，1~100
	Quantity *uint64 `json:"Quantity,omitempty" name:"Quantity"`
}

type DeliverTidsRequest struct {
	*tchttp.BaseRequest
	
	// 订单ID
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 数量，1~100
	Quantity *uint64 `json:"Quantity,omitempty" name:"Quantity"`
}

func (r *DeliverTidsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeliverTidsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrderId")
	delete(f, "Quantity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeliverTidsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeliverTidsResponseParams struct {
	// 空发的TID信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TidSet []*TidKeysInfo `json:"TidSet,omitempty" name:"TidSet"`

	// 产品公钥
	ProductKey *string `json:"ProductKey,omitempty" name:"ProductKey"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeliverTidsResponse struct {
	*tchttp.BaseResponse
	Response *DeliverTidsResponseParams `json:"Response"`
}

func (r *DeliverTidsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeliverTidsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAvailableLibCountRequestParams struct {
	// 订单编号
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`
}

type DescribeAvailableLibCountRequest struct {
	*tchttp.BaseRequest
	
	// 订单编号
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`
}

func (r *DescribeAvailableLibCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAvailableLibCountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrderId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAvailableLibCountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAvailableLibCountResponseParams struct {
	// 可空发的白盒密钥数量
	Quantity *uint64 `json:"Quantity,omitempty" name:"Quantity"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAvailableLibCountResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAvailableLibCountResponseParams `json:"Response"`
}

func (r *DescribeAvailableLibCountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAvailableLibCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePermissionRequestParams struct {

}

type DescribePermissionRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribePermissionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePermissionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePermissionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePermissionResponseParams struct {
	// 企业用户
	EnterpriseUser *bool `json:"EnterpriseUser,omitempty" name:"EnterpriseUser"`

	// 下载控制台权限
	DownloadPermission *string `json:"DownloadPermission,omitempty" name:"DownloadPermission"`

	// 使用控制台权限
	UsePermission *string `json:"UsePermission,omitempty" name:"UsePermission"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePermissionResponse struct {
	*tchttp.BaseResponse
	Response *DescribePermissionResponseParams `json:"Response"`
}

func (r *DescribePermissionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePermissionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadTidsRequestParams struct {
	// 订单编号
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 下载数量：1~10
	Quantity *uint64 `json:"Quantity,omitempty" name:"Quantity"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadTidsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrderId")
	delete(f, "Quantity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DownloadTidsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadTidsResponseParams struct {
	// 下载的TID信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TidSet []*TidKeysInfo `json:"TidSet,omitempty" name:"TidSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DownloadTidsResponse struct {
	*tchttp.BaseResponse
	Response *DownloadTidsResponseParams `json:"Response"`
}

func (r *DownloadTidsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

	// 软加固白盒密钥下载地址
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 软加固设备标识码
	DeviceCode *string `json:"DeviceCode,omitempty" name:"DeviceCode"`
}

// Predefined struct for user
type UploadDeviceUniqueCodeRequestParams struct {
	// 硬件唯一标识码
	CodeSet []*string `json:"CodeSet,omitempty" name:"CodeSet"`

	// 硬件标识码绑定的申请编号
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`
}

type UploadDeviceUniqueCodeRequest struct {
	*tchttp.BaseRequest
	
	// 硬件唯一标识码
	CodeSet []*string `json:"CodeSet,omitempty" name:"CodeSet"`

	// 硬件标识码绑定的申请编号
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`
}

func (r *UploadDeviceUniqueCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadDeviceUniqueCodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CodeSet")
	delete(f, "OrderId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadDeviceUniqueCodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadDeviceUniqueCodeResponseParams struct {
	// 本次已上传数量
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 重复的硬件唯一标识码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExistedCodeSet []*string `json:"ExistedCodeSet,omitempty" name:"ExistedCodeSet"`

	// 剩余可上传数量
	LeftQuantity *uint64 `json:"LeftQuantity,omitempty" name:"LeftQuantity"`

	// 错误的硬件唯一标识码
	// 注意：此字段可能返回 null，表示取不到有效值。
	IllegalCodeSet []*string `json:"IllegalCodeSet,omitempty" name:"IllegalCodeSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UploadDeviceUniqueCodeResponse struct {
	*tchttp.BaseResponse
	Response *UploadDeviceUniqueCodeResponseParams `json:"Response"`
}

func (r *UploadDeviceUniqueCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadDeviceUniqueCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyChipBurnInfoRequestParams struct {
	// 验证数据
	Data *string `json:"Data,omitempty" name:"Data"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyChipBurnInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Data")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifyChipBurnInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyChipBurnInfoResponseParams struct {
	// 验证结果
	Pass *bool `json:"Pass,omitempty" name:"Pass"`

	// 已验证次数
	VerifiedTimes *uint64 `json:"VerifiedTimes,omitempty" name:"VerifiedTimes"`

	// 剩余验证次数
	LeftTimes *uint64 `json:"LeftTimes,omitempty" name:"LeftTimes"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type VerifyChipBurnInfoResponse struct {
	*tchttp.BaseResponse
	Response *VerifyChipBurnInfoResponseParams `json:"Response"`
}

func (r *VerifyChipBurnInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyChipBurnInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}