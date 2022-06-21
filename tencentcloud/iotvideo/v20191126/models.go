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

package v20191126

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type BindDevInfo struct {
	// 设备TID
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 设备型号
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceModel *string `json:"DeviceModel,omitempty" name:"DeviceModel"`

	// 用户角色，owner：主人，guest：访客
	Role *string `json:"Role,omitempty" name:"Role"`
}

type BindUsrInfo struct {
	// IotVideo平台分配给终端用户的用户id
	AccessId *string `json:"AccessId,omitempty" name:"AccessId"`

	// 用户角色，owner：主人，guest：访客
	Role *string `json:"Role,omitempty" name:"Role"`
}

// Predefined struct for user
type ClearDeviceActiveCodeRequestParams struct {
	// 设备TID列表，0<元素数量<=100
	Tids []*string `json:"Tids,omitempty" name:"Tids"`
}

type ClearDeviceActiveCodeRequest struct {
	*tchttp.BaseRequest
	
	// 设备TID列表，0<元素数量<=100
	Tids []*string `json:"Tids,omitempty" name:"Tids"`
}

func (r *ClearDeviceActiveCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ClearDeviceActiveCodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Tids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ClearDeviceActiveCodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ClearDeviceActiveCodeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ClearDeviceActiveCodeResponse struct {
	*tchttp.BaseResponse
	Response *ClearDeviceActiveCodeResponseParams `json:"Response"`
}

func (r *ClearDeviceActiveCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ClearDeviceActiveCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Contents struct {
	// 英文，长度不超过300个字符
	// 注意：此字段可能返回 null，表示取不到有效值。
	En *string `json:"En,omitempty" name:"En"`

	// 中文简体，长度不超过300个字符
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cn *string `json:"Cn,omitempty" name:"Cn"`

	// 中文繁体(Traditional Chinese)，长度不超过300个字符
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tc *string `json:"Tc,omitempty" name:"Tc"`

	// 默认语言，最多不超过300个字符
	// 注意：此字段可能返回 null，表示取不到有效值。
	Default *string `json:"Default,omitempty" name:"Default"`
}

// Predefined struct for user
type CreateAnonymousAccessTokenRequestParams struct {
	// Token的TTL(time to alive)分钟数,最大值1440(即24小时)
	TtlMinutes *int64 `json:"TtlMinutes,omitempty" name:"TtlMinutes"`

	// 设备ID。创建Token时, 此参数为必须项
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 旧的AccessToken。续期Token时，此参数为必须
	OldAccessToken *string `json:"OldAccessToken,omitempty" name:"OldAccessToken"`
}

type CreateAnonymousAccessTokenRequest struct {
	*tchttp.BaseRequest
	
	// Token的TTL(time to alive)分钟数,最大值1440(即24小时)
	TtlMinutes *int64 `json:"TtlMinutes,omitempty" name:"TtlMinutes"`

	// 设备ID。创建Token时, 此参数为必须项
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 旧的AccessToken。续期Token时，此参数为必须
	OldAccessToken *string `json:"OldAccessToken,omitempty" name:"OldAccessToken"`
}

func (r *CreateAnonymousAccessTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAnonymousAccessTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TtlMinutes")
	delete(f, "Tid")
	delete(f, "OldAccessToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAnonymousAccessTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAnonymousAccessTokenResponseParams struct {
	// 终端用户在IoT Video上的唯一标识ID
	AccessId *string `json:"AccessId,omitempty" name:"AccessId"`

	// IoT Video平台的AccessToken
	AccessToken *string `json:"AccessToken,omitempty" name:"AccessToken"`

	// Token的过期时间，单位秒(UTC时间)
	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAnonymousAccessTokenResponse struct {
	*tchttp.BaseResponse
	Response *CreateAnonymousAccessTokenResponseParams `json:"Response"`
}

func (r *CreateAnonymousAccessTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAnonymousAccessTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAppUsrRequestParams struct {
	// 标识用户的唯一ID，防止同一个用户多次注册
	CunionId *string `json:"CunionId,omitempty" name:"CunionId"`

	// 用于小程序关联手机号
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`
}

type CreateAppUsrRequest struct {
	*tchttp.BaseRequest
	
	// 标识用户的唯一ID，防止同一个用户多次注册
	CunionId *string `json:"CunionId,omitempty" name:"CunionId"`

	// 用于小程序关联手机号
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`
}

func (r *CreateAppUsrRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAppUsrRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CunionId")
	delete(f, "Mobile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAppUsrRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAppUsrResponseParams struct {
	// 厂商云标识用户的唯一ID
	CunionId *string `json:"CunionId,omitempty" name:"CunionId"`

	// 客户的终端用户在IoT Video上的唯一标识ID
	AccessId *string `json:"AccessId,omitempty" name:"AccessId"`

	// 用户是否为新创建
	NewRegist *bool `json:"NewRegist,omitempty" name:"NewRegist"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAppUsrResponse struct {
	*tchttp.BaseResponse
	Response *CreateAppUsrResponseParams `json:"Response"`
}

func (r *CreateAppUsrResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAppUsrResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBindingRequestParams struct {
	// 终端用户在IoT Video上的唯一标识ID
	AccessId *string `json:"AccessId,omitempty" name:"AccessId"`

	// 设备TID
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 用户角色，owner：主人，guest：访客
	Role *string `json:"Role,omitempty" name:"Role"`

	// 是否踢掉之前的主人，true：踢掉；false：不踢掉。当role为guest时，可以不填
	ForceBind *bool `json:"ForceBind,omitempty" name:"ForceBind"`

	// 设备昵称，最多不超过64个字符
	Nick *string `json:"Nick,omitempty" name:"Nick"`

	// 绑定过程中的会话token，由设备通过SDK接口确认是否允许绑定的token，用于增加设备被绑定的安全性
	BindToken *string `json:"BindToken,omitempty" name:"BindToken"`
}

type CreateBindingRequest struct {
	*tchttp.BaseRequest
	
	// 终端用户在IoT Video上的唯一标识ID
	AccessId *string `json:"AccessId,omitempty" name:"AccessId"`

	// 设备TID
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 用户角色，owner：主人，guest：访客
	Role *string `json:"Role,omitempty" name:"Role"`

	// 是否踢掉之前的主人，true：踢掉；false：不踢掉。当role为guest时，可以不填
	ForceBind *bool `json:"ForceBind,omitempty" name:"ForceBind"`

	// 设备昵称，最多不超过64个字符
	Nick *string `json:"Nick,omitempty" name:"Nick"`

	// 绑定过程中的会话token，由设备通过SDK接口确认是否允许绑定的token，用于增加设备被绑定的安全性
	BindToken *string `json:"BindToken,omitempty" name:"BindToken"`
}

func (r *CreateBindingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBindingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessId")
	delete(f, "Tid")
	delete(f, "Role")
	delete(f, "ForceBind")
	delete(f, "Nick")
	delete(f, "BindToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBindingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBindingResponseParams struct {
	// 访问设备的AccessToken
	AccessToken *string `json:"AccessToken,omitempty" name:"AccessToken"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateBindingResponse struct {
	*tchttp.BaseResponse
	Response *CreateBindingResponseParams `json:"Response"`
}

func (r *CreateBindingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBindingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDevTokenRequestParams struct {
	// 客户的终端用户在IoT Video上的唯一标识ID
	AccessId *string `json:"AccessId,omitempty" name:"AccessId"`

	// 设备TID列表,0<元素数量<=100
	Tids []*string `json:"Tids,omitempty" name:"Tids"`

	// Token的TTL(time to alive)分钟数
	TtlMinutes *uint64 `json:"TtlMinutes,omitempty" name:"TtlMinutes"`
}

type CreateDevTokenRequest struct {
	*tchttp.BaseRequest
	
	// 客户的终端用户在IoT Video上的唯一标识ID
	AccessId *string `json:"AccessId,omitempty" name:"AccessId"`

	// 设备TID列表,0<元素数量<=100
	Tids []*string `json:"Tids,omitempty" name:"Tids"`

	// Token的TTL(time to alive)分钟数
	TtlMinutes *uint64 `json:"TtlMinutes,omitempty" name:"TtlMinutes"`
}

func (r *CreateDevTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDevTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessId")
	delete(f, "Tids")
	delete(f, "TtlMinutes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDevTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDevTokenResponseParams struct {
	// 返回的用户token列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*DevTokenInfo `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDevTokenResponse struct {
	*tchttp.BaseResponse
	Response *CreateDevTokenResponseParams `json:"Response"`
}

func (r *CreateDevTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDevTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDevicesRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 创建设备的数量，数量范围1-100
	Number *uint64 `json:"Number,omitempty" name:"Number"`

	// 设备名称前缀，支持英文、数字，不超过10字符
	NamePrefix *string `json:"NamePrefix,omitempty" name:"NamePrefix"`

	// 操作人
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type CreateDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 创建设备的数量，数量范围1-100
	Number *uint64 `json:"Number,omitempty" name:"Number"`

	// 设备名称前缀，支持英文、数字，不超过10字符
	NamePrefix *string `json:"NamePrefix,omitempty" name:"NamePrefix"`

	// 操作人
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *CreateDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "Number")
	delete(f, "NamePrefix")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDevicesResponseParams struct {
	// 新创建设备的认证信息
	Data []*DeviceCertificate `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDevicesResponse struct {
	*tchttp.BaseResponse
	Response *CreateDevicesResponseParams `json:"Response"`
}

func (r *CreateDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGencodeRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 物模型发布版本号,-1代表未发布的，保存的是草稿箱的版本。1代表已发布的物模型。
	Revision *int64 `json:"Revision,omitempty" name:"Revision"`
}

type CreateGencodeRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 物模型发布版本号,-1代表未发布的，保存的是草稿箱的版本。1代表已发布的物模型。
	Revision *int64 `json:"Revision,omitempty" name:"Revision"`
}

func (r *CreateGencodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGencodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "Revision")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateGencodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGencodeResponseParams struct {
	// 生成的源代码(zip压缩后的base64编码)
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZipCode *string `json:"ZipCode,omitempty" name:"ZipCode"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateGencodeResponse struct {
	*tchttp.BaseResponse
	Response *CreateGencodeResponseParams `json:"Response"`
}

func (r *CreateGencodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGencodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIotDataTypeRequestParams struct {
	// 用户自定义数据类型，json格式的字符串
	IotDataType *string `json:"IotDataType,omitempty" name:"IotDataType"`
}

type CreateIotDataTypeRequest struct {
	*tchttp.BaseRequest
	
	// 用户自定义数据类型，json格式的字符串
	IotDataType *string `json:"IotDataType,omitempty" name:"IotDataType"`
}

func (r *CreateIotDataTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIotDataTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IotDataType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIotDataTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIotDataTypeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateIotDataTypeResponse struct {
	*tchttp.BaseResponse
	Response *CreateIotDataTypeResponseParams `json:"Response"`
}

func (r *CreateIotDataTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIotDataTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIotModelRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 物模型json串
	IotModel *string `json:"IotModel,omitempty" name:"IotModel"`
}

type CreateIotModelRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 物模型json串
	IotModel *string `json:"IotModel,omitempty" name:"IotModel"`
}

func (r *CreateIotModelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIotModelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "IotModel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIotModelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIotModelResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateIotModelResponse struct {
	*tchttp.BaseResponse
	Response *CreateIotModelResponseParams `json:"Response"`
}

func (r *CreateIotModelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIotModelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProductRequestParams struct {
	// 产器型号(APP产品,为APP包名)
	ProductModel *string `json:"ProductModel,omitempty" name:"ProductModel"`

	// 产品名称
	// 仅支持中文、英文、数字、下划线，不超过32个字符
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 产品描述信息
	// 不支持单引号、双引号、退格符、回车符、换行符、制表符、反斜杠、下划线、“%”、“#”、“$”，不超过128字符
	ProductDescription *string `json:"ProductDescription,omitempty" name:"ProductDescription"`

	// 设备功能码（ypsxth:音频双向通话 ，spdxth:视频单向通话）
	Features []*string `json:"Features,omitempty" name:"Features"`

	// 主芯片产商ID
	ChipManufactureId *string `json:"ChipManufactureId,omitempty" name:"ChipManufactureId"`

	// 主芯片ID
	ChipId *string `json:"ChipId,omitempty" name:"ChipId"`

	// 地域：
	// China-Mainland（中国大陆）
	// China-Hong Kong, Macao and Taiwan（港澳台地区）
	// America（美国）
	// Europe（欧洲）
	// India（印度）
	// Other-Overseas（其他境外地区）
	ProductRegion *string `json:"ProductRegion,omitempty" name:"ProductRegion"`

	// 设备类型, 0-普通视频设备，1-NVR设备
	ProductCate *uint64 `json:"ProductCate,omitempty" name:"ProductCate"`

	// 接入模型，bit0是0：公版小程序未接入，bit0是1：公版小程序已接入
	AccessMode *int64 `json:"AccessMode,omitempty" name:"AccessMode"`

	// Linux,Android,Liteos等系统
	Os *string `json:"Os,omitempty" name:"Os"`

	// 芯片架构，只是针对操作系统为android的
	ChipArch *string `json:"ChipArch,omitempty" name:"ChipArch"`
}

type CreateProductRequest struct {
	*tchttp.BaseRequest
	
	// 产器型号(APP产品,为APP包名)
	ProductModel *string `json:"ProductModel,omitempty" name:"ProductModel"`

	// 产品名称
	// 仅支持中文、英文、数字、下划线，不超过32个字符
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 产品描述信息
	// 不支持单引号、双引号、退格符、回车符、换行符、制表符、反斜杠、下划线、“%”、“#”、“$”，不超过128字符
	ProductDescription *string `json:"ProductDescription,omitempty" name:"ProductDescription"`

	// 设备功能码（ypsxth:音频双向通话 ，spdxth:视频单向通话）
	Features []*string `json:"Features,omitempty" name:"Features"`

	// 主芯片产商ID
	ChipManufactureId *string `json:"ChipManufactureId,omitempty" name:"ChipManufactureId"`

	// 主芯片ID
	ChipId *string `json:"ChipId,omitempty" name:"ChipId"`

	// 地域：
	// China-Mainland（中国大陆）
	// China-Hong Kong, Macao and Taiwan（港澳台地区）
	// America（美国）
	// Europe（欧洲）
	// India（印度）
	// Other-Overseas（其他境外地区）
	ProductRegion *string `json:"ProductRegion,omitempty" name:"ProductRegion"`

	// 设备类型, 0-普通视频设备，1-NVR设备
	ProductCate *uint64 `json:"ProductCate,omitempty" name:"ProductCate"`

	// 接入模型，bit0是0：公版小程序未接入，bit0是1：公版小程序已接入
	AccessMode *int64 `json:"AccessMode,omitempty" name:"AccessMode"`

	// Linux,Android,Liteos等系统
	Os *string `json:"Os,omitempty" name:"Os"`

	// 芯片架构，只是针对操作系统为android的
	ChipArch *string `json:"ChipArch,omitempty" name:"ChipArch"`
}

func (r *CreateProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductModel")
	delete(f, "ProductName")
	delete(f, "ProductDescription")
	delete(f, "Features")
	delete(f, "ChipManufactureId")
	delete(f, "ChipId")
	delete(f, "ProductRegion")
	delete(f, "ProductCate")
	delete(f, "AccessMode")
	delete(f, "Os")
	delete(f, "ChipArch")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProductResponseParams struct {
	// 产品详细信息
	Data *ProductBase `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateProductResponse struct {
	*tchttp.BaseResponse
	Response *CreateProductResponseParams `json:"Response"`
}

func (r *CreateProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStorageRequestParams struct {
	// 云存套餐ID
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`

	// 设备TID
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 用户唯一标识，由厂商保证内部唯一性
	UserTag *string `json:"UserTag,omitempty" name:"UserTag"`
}

type CreateStorageRequest struct {
	*tchttp.BaseRequest
	
	// 云存套餐ID
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`

	// 设备TID
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 用户唯一标识，由厂商保证内部唯一性
	UserTag *string `json:"UserTag,omitempty" name:"UserTag"`
}

func (r *CreateStorageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStorageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PkgId")
	delete(f, "Tid")
	delete(f, "UserTag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStorageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStorageResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateStorageResponse struct {
	*tchttp.BaseResponse
	Response *CreateStorageResponseParams `json:"Response"`
}

func (r *CreateStorageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStorageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStorageServiceRequestParams struct {
	// 云存套餐ID：
	// yc1m3d ： 全时3天存储月套餐。
	// yc1m7d ： 全时7天存储月套餐。
	// yc1m30d ：全时30天存储月套餐。
	// yc1y3d ：全时3天存储年套餐。
	// yc1y7d ：全时7天存储年套餐。
	// yc1y30d ：全时30天存储年套餐。
	// ye1m3d ：事件3天存储月套餐。
	// ye1m7d ：事件7天存储月套餐。
	// ye1m30d ：事件30天存储月套餐 。
	// ye1y3d ：事件3天存储年套餐。
	// ye1y7d ：事件7天存储年套餐。
	// ye1y30d ：事件30天存储年套餐。
	// yc1w7d : 全时7天存储周套餐。
	// ye1w7d : 事件7天存储周套餐。
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`

	// 设备TID
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 订单数量,可一次性创建多个订单
	OrderCount *int64 `json:"OrderCount,omitempty" name:"OrderCount"`

	// 云存服务所在的区域,如ap-guangzhou,ap-singapore, na-siliconvalley, eu-frankfurt
	StorageRegion *string `json:"StorageRegion,omitempty" name:"StorageRegion"`

	// 视频流通道号。(对于存在多路视频流的设备，如NVR设备，与设备实际视频流通道号对应)
	ChnNum *int64 `json:"ChnNum,omitempty" name:"ChnNum"`

	// 设备主人用户在IoT Video平台的注册ID。该参数用于验证Paas/Saas平台的设备/用户关系链是否一致
	AccessId *string `json:"AccessId,omitempty" name:"AccessId"`

	// 服务生效时间,若不指定此参数，服务立即生效
	EnableTime *int64 `json:"EnableTime,omitempty" name:"EnableTime"`
}

type CreateStorageServiceRequest struct {
	*tchttp.BaseRequest
	
	// 云存套餐ID：
	// yc1m3d ： 全时3天存储月套餐。
	// yc1m7d ： 全时7天存储月套餐。
	// yc1m30d ：全时30天存储月套餐。
	// yc1y3d ：全时3天存储年套餐。
	// yc1y7d ：全时7天存储年套餐。
	// yc1y30d ：全时30天存储年套餐。
	// ye1m3d ：事件3天存储月套餐。
	// ye1m7d ：事件7天存储月套餐。
	// ye1m30d ：事件30天存储月套餐 。
	// ye1y3d ：事件3天存储年套餐。
	// ye1y7d ：事件7天存储年套餐。
	// ye1y30d ：事件30天存储年套餐。
	// yc1w7d : 全时7天存储周套餐。
	// ye1w7d : 事件7天存储周套餐。
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`

	// 设备TID
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 订单数量,可一次性创建多个订单
	OrderCount *int64 `json:"OrderCount,omitempty" name:"OrderCount"`

	// 云存服务所在的区域,如ap-guangzhou,ap-singapore, na-siliconvalley, eu-frankfurt
	StorageRegion *string `json:"StorageRegion,omitempty" name:"StorageRegion"`

	// 视频流通道号。(对于存在多路视频流的设备，如NVR设备，与设备实际视频流通道号对应)
	ChnNum *int64 `json:"ChnNum,omitempty" name:"ChnNum"`

	// 设备主人用户在IoT Video平台的注册ID。该参数用于验证Paas/Saas平台的设备/用户关系链是否一致
	AccessId *string `json:"AccessId,omitempty" name:"AccessId"`

	// 服务生效时间,若不指定此参数，服务立即生效
	EnableTime *int64 `json:"EnableTime,omitempty" name:"EnableTime"`
}

func (r *CreateStorageServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStorageServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PkgId")
	delete(f, "Tid")
	delete(f, "OrderCount")
	delete(f, "StorageRegion")
	delete(f, "ChnNum")
	delete(f, "AccessId")
	delete(f, "EnableTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStorageServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStorageServiceResponseParams struct {
	// 标志是否为续订
	IsRenew *bool `json:"IsRenew,omitempty" name:"IsRenew"`

	// 云存服务ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 云存服务所在的区域
	StorageRegion *string `json:"StorageRegion,omitempty" name:"StorageRegion"`

	// 设备TID
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 视频流通道号。(对于存在多路视频流的设备，如NVR设备，与设备实际视频流通道号对应)
	ChnNum *int64 `json:"ChnNum,omitempty" name:"ChnNum"`

	// 终端用户在IoT Video平台的注册ID
	AccessId *string `json:"AccessId,omitempty" name:"AccessId"`

	// 服务开始时间
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 服务失效时间
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 服务状态
	// 1：正常使用中
	// 2：待续费。设备云存服务已到期，但是历史云存数据未过期。续费后仍可查看这些历史数据。
	// 3：已过期。查询不到设备保存在云端的数据。
	// 4：等待服务生效。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 新增的云存定单列表
	Data []*StorageOrder `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateStorageServiceResponse struct {
	*tchttp.BaseResponse
	Response *CreateStorageServiceResponseParams `json:"Response"`
}

func (r *CreateStorageServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStorageServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTraceIdsRequestParams struct {
	// 设备TID列表
	Tids []*string `json:"Tids,omitempty" name:"Tids"`
}

type CreateTraceIdsRequest struct {
	*tchttp.BaseRequest
	
	// 设备TID列表
	Tids []*string `json:"Tids,omitempty" name:"Tids"`
}

func (r *CreateTraceIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTraceIdsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Tids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTraceIdsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTraceIdsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateTraceIdsResponse struct {
	*tchttp.BaseResponse
	Response *CreateTraceIdsResponseParams `json:"Response"`
}

func (r *CreateTraceIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTraceIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUploadPathRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 固件文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`
}

type CreateUploadPathRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 固件文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`
}

func (r *CreateUploadPathRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUploadPathRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "FileName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUploadPathRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUploadPathResponseParams struct {
	// 固件上传地址URL，用户可将本地的固件文件通过该URL以PUT的请求方式上传。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateUploadPathResponse struct {
	*tchttp.BaseResponse
	Response *CreateUploadPathResponseParams `json:"Response"`
}

func (r *CreateUploadPathResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUploadPathResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUsrTokenRequestParams struct {
	// 终端用户在IoT Video上的唯一标识ID
	AccessId *string `json:"AccessId,omitempty" name:"AccessId"`

	// 终端唯一ID，用于区分同一个用户的多个终端
	UniqueId *string `json:"UniqueId,omitempty" name:"UniqueId"`

	// Token的TTL(time to alive)分钟数
	TtlMinutes *uint64 `json:"TtlMinutes,omitempty" name:"TtlMinutes"`

	// 旧的AccessToken。续期Token时，此参数为必须。
	OldAccessToken *string `json:"OldAccessToken,omitempty" name:"OldAccessToken"`
}

type CreateUsrTokenRequest struct {
	*tchttp.BaseRequest
	
	// 终端用户在IoT Video上的唯一标识ID
	AccessId *string `json:"AccessId,omitempty" name:"AccessId"`

	// 终端唯一ID，用于区分同一个用户的多个终端
	UniqueId *string `json:"UniqueId,omitempty" name:"UniqueId"`

	// Token的TTL(time to alive)分钟数
	TtlMinutes *uint64 `json:"TtlMinutes,omitempty" name:"TtlMinutes"`

	// 旧的AccessToken。续期Token时，此参数为必须。
	OldAccessToken *string `json:"OldAccessToken,omitempty" name:"OldAccessToken"`
}

func (r *CreateUsrTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUsrTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessId")
	delete(f, "UniqueId")
	delete(f, "TtlMinutes")
	delete(f, "OldAccessToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUsrTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUsrTokenResponseParams struct {
	// 终端用户在IoT Video上的唯一标识ID
	AccessId *string `json:"AccessId,omitempty" name:"AccessId"`

	// IoT Video平台的AccessToken
	AccessToken *string `json:"AccessToken,omitempty" name:"AccessToken"`

	// Token的过期时间，单位秒(UTC时间)
	ExpireTime *uint64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 终端ID
	TerminalId *string `json:"TerminalId,omitempty" name:"TerminalId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateUsrTokenResponse struct {
	*tchttp.BaseResponse
	Response *CreateUsrTokenResponseParams `json:"Response"`
}

func (r *CreateUsrTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUsrTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Data struct {
	// 直播协议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 流媒体播放地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	URI *string `json:"URI,omitempty" name:"URI"`

	// 流媒体地址过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 视频编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoCodec *string `json:"VideoCodec,omitempty" name:"VideoCodec"`

	// 音频编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioCodec *string `json:"AudioCodec,omitempty" name:"AudioCodec"`
}

// Predefined struct for user
type DeleteAppUsrRequestParams struct {
	// 客户的终端用户在IoT Video上的唯一标识ID
	AccessId *string `json:"AccessId,omitempty" name:"AccessId"`
}

type DeleteAppUsrRequest struct {
	*tchttp.BaseRequest
	
	// 客户的终端用户在IoT Video上的唯一标识ID
	AccessId *string `json:"AccessId,omitempty" name:"AccessId"`
}

func (r *DeleteAppUsrRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAppUsrRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAppUsrRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAppUsrResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteAppUsrResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAppUsrResponseParams `json:"Response"`
}

func (r *DeleteAppUsrResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAppUsrResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBindingRequestParams struct {
	// 终端用户在IoT Video上的唯一标识ID
	AccessId *string `json:"AccessId,omitempty" name:"AccessId"`

	// 设备TID
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 用户角色，owner：主人，guest：访客
	Role *string `json:"Role,omitempty" name:"Role"`
}

type DeleteBindingRequest struct {
	*tchttp.BaseRequest
	
	// 终端用户在IoT Video上的唯一标识ID
	AccessId *string `json:"AccessId,omitempty" name:"AccessId"`

	// 设备TID
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 用户角色，owner：主人，guest：访客
	Role *string `json:"Role,omitempty" name:"Role"`
}

func (r *DeleteBindingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBindingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessId")
	delete(f, "Tid")
	delete(f, "Role")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBindingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBindingResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteBindingResponse struct {
	*tchttp.BaseResponse
	Response *DeleteBindingResponseParams `json:"Response"`
}

func (r *DeleteBindingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBindingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDeviceRequestParams struct {
	// 设备TID列表
	Tids []*string `json:"Tids,omitempty" name:"Tids"`
}

type DeleteDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 设备TID列表
	Tids []*string `json:"Tids,omitempty" name:"Tids"`
}

func (r *DeleteDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Tids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDeviceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteDeviceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDeviceResponseParams `json:"Response"`
}

func (r *DeleteDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIotDataTypeRequestParams struct {
	// 自定义数据类型的标识符
	TypeId *string `json:"TypeId,omitempty" name:"TypeId"`
}

type DeleteIotDataTypeRequest struct {
	*tchttp.BaseRequest
	
	// 自定义数据类型的标识符
	TypeId *string `json:"TypeId,omitempty" name:"TypeId"`
}

func (r *DeleteIotDataTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIotDataTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TypeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteIotDataTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIotDataTypeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteIotDataTypeResponse struct {
	*tchttp.BaseResponse
	Response *DeleteIotDataTypeResponseParams `json:"Response"`
}

func (r *DeleteIotDataTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIotDataTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMessageQueueRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

type DeleteMessageQueueRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

func (r *DeleteMessageQueueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMessageQueueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMessageQueueRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMessageQueueResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteMessageQueueResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMessageQueueResponseParams `json:"Response"`
}

func (r *DeleteMessageQueueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMessageQueueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOtaVersionRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 固件版本号，格式为x.y.z， x，y 范围0-63，z范围1~524288
	OtaVersion *string `json:"OtaVersion,omitempty" name:"OtaVersion"`

	// 操作人
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type DeleteOtaVersionRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 固件版本号，格式为x.y.z， x，y 范围0-63，z范围1~524288
	OtaVersion *string `json:"OtaVersion,omitempty" name:"OtaVersion"`

	// 操作人
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *DeleteOtaVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOtaVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "OtaVersion")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteOtaVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOtaVersionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteOtaVersionResponse struct {
	*tchttp.BaseResponse
	Response *DeleteOtaVersionResponseParams `json:"Response"`
}

func (r *DeleteOtaVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOtaVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProductRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

type DeleteProductRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

func (r *DeleteProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProductResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteProductResponse struct {
	*tchttp.BaseResponse
	Response *DeleteProductResponseParams `json:"Response"`
}

func (r *DeleteProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTraceIdsRequestParams struct {
	// 设备TID列表
	Tids []*string `json:"Tids,omitempty" name:"Tids"`
}

type DeleteTraceIdsRequest struct {
	*tchttp.BaseRequest
	
	// 设备TID列表
	Tids []*string `json:"Tids,omitempty" name:"Tids"`
}

func (r *DeleteTraceIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTraceIdsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Tids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTraceIdsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTraceIdsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteTraceIdsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTraceIdsResponseParams `json:"Response"`
}

func (r *DeleteTraceIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTraceIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeliverStorageServiceRequestParams struct {
	// 待转移的源云存服务ID
	SrcServiceId *string `json:"SrcServiceId,omitempty" name:"SrcServiceId"`

	// 设备TID
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 视频流通道号。(对于存在多路视频流的设备，如NVR设备，与设备实际视频流通道号对应)
	ChnNum *int64 `json:"ChnNum,omitempty" name:"ChnNum"`

	// 设备主人用户在IoT Video平台的注册ID。该参数用于验证Paas/Saas平台的设备/用户关系链是否一致
	AccessId *string `json:"AccessId,omitempty" name:"AccessId"`
}

type DeliverStorageServiceRequest struct {
	*tchttp.BaseRequest
	
	// 待转移的源云存服务ID
	SrcServiceId *string `json:"SrcServiceId,omitempty" name:"SrcServiceId"`

	// 设备TID
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 视频流通道号。(对于存在多路视频流的设备，如NVR设备，与设备实际视频流通道号对应)
	ChnNum *int64 `json:"ChnNum,omitempty" name:"ChnNum"`

	// 设备主人用户在IoT Video平台的注册ID。该参数用于验证Paas/Saas平台的设备/用户关系链是否一致
	AccessId *string `json:"AccessId,omitempty" name:"AccessId"`
}

func (r *DeliverStorageServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeliverStorageServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SrcServiceId")
	delete(f, "Tid")
	delete(f, "ChnNum")
	delete(f, "AccessId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeliverStorageServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeliverStorageServiceResponseParams struct {
	// 被转出的云存服务ID
	SrcServiceId *string `json:"SrcServiceId,omitempty" name:"SrcServiceId"`

	// 被转入的云存服务ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 云存服务所在的区域
	StorageRegion *string `json:"StorageRegion,omitempty" name:"StorageRegion"`

	// 设备TID
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 视频流通道号。(对于存在多路视频流的设备，如NVR设备，与设备实际视频流通道号对应)
	ChnNum *int64 `json:"ChnNum,omitempty" name:"ChnNum"`

	// 终端用户在IoT Video平台的注册ID
	AccessId *string `json:"AccessId,omitempty" name:"AccessId"`

	// 服务开始时间
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 服务失效时间
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 服务状态
	// 1：正常使用中
	// 2：待续费。设备云存服务已到期，但是历史云存数据未过期。续费后仍可查看这些历史数据。
	// 3：已过期。查询不到设备保存在云端的数据。
	// 4：等待服务生效。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 新增的云存定单列表
	Data []*StorageOrder `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeliverStorageServiceResponse struct {
	*tchttp.BaseResponse
	Response *DeliverStorageServiceResponseParams `json:"Response"`
}

func (r *DeliverStorageServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeliverStorageServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountBalanceRequestParams struct {
	// 账户类型 1:设备接入 2:云存
	AccountType *int64 `json:"AccountType,omitempty" name:"AccountType"`
}

type DescribeAccountBalanceRequest struct {
	*tchttp.BaseRequest
	
	// 账户类型 1:设备接入 2:云存
	AccountType *int64 `json:"AccountType,omitempty" name:"AccountType"`
}

func (r *DescribeAccountBalanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountBalanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccountType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountBalanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountBalanceResponseParams struct {
	// 账户类型 1=设备接入;2=云存。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountType *int64 `json:"AccountType,omitempty" name:"AccountType"`

	// 余额, 单位 : 分(人民币)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Balance *int64 `json:"Balance,omitempty" name:"Balance"`

	// 账户状态，1=正常；8=冻结；9=销户。
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *int64 `json:"State,omitempty" name:"State"`

	// 最后修改时间，UTC值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastUpdateTime *int64 `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAccountBalanceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccountBalanceResponseParams `json:"Response"`
}

func (r *DescribeAccountBalanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountBalanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBindDevRequestParams struct {
	// 终端用户在IoT Video上的唯一标识ID
	AccessId *string `json:"AccessId,omitempty" name:"AccessId"`
}

type DescribeBindDevRequest struct {
	*tchttp.BaseRequest
	
	// 终端用户在IoT Video上的唯一标识ID
	AccessId *string `json:"AccessId,omitempty" name:"AccessId"`
}

func (r *DescribeBindDevRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBindDevRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBindDevRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBindDevResponseParams struct {
	// 绑定的设备列表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*BindDevInfo `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBindDevResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBindDevResponseParams `json:"Response"`
}

func (r *DescribeBindDevResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBindDevResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBindUsrRequestParams struct {
	// 设备TID
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 设备主人的AccessId
	AccessId *string `json:"AccessId,omitempty" name:"AccessId"`
}

type DescribeBindUsrRequest struct {
	*tchttp.BaseRequest
	
	// 设备TID
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 设备主人的AccessId
	AccessId *string `json:"AccessId,omitempty" name:"AccessId"`
}

func (r *DescribeBindUsrRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBindUsrRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Tid")
	delete(f, "AccessId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBindUsrRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBindUsrResponseParams struct {
	// 具有绑定关系的终端用户信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*BindUsrInfo `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBindUsrResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBindUsrResponseParams `json:"Response"`
}

func (r *DescribeBindUsrResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBindUsrResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceModelRequestParams struct {
	// 设备TID
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 物模型的分支路径
	Branch *string `json:"Branch,omitempty" name:"Branch"`
}

type DescribeDeviceModelRequest struct {
	*tchttp.BaseRequest
	
	// 设备TID
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 物模型的分支路径
	Branch *string `json:"Branch,omitempty" name:"Branch"`
}

func (r *DescribeDeviceModelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceModelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Tid")
	delete(f, "Branch")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceModelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceModelResponseParams struct {
	// 设备物模型信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *DeviceModelData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDeviceModelResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceModelResponseParams `json:"Response"`
}

func (r *DescribeDeviceModelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceModelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceRequestParams struct {
	// 设备TID
	Tid *string `json:"Tid,omitempty" name:"Tid"`
}

type DescribeDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 设备TID
	Tid *string `json:"Tid,omitempty" name:"Tid"`
}

func (r *DescribeDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Tid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceResponseParams struct {
	// 设备信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *DeviceData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDeviceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceResponseParams `json:"Response"`
}

func (r *DescribeDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDevicesRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 是否返回全量数据
	// 当该值为false时，返回值中的设备物模型、固件版本、在线状态、最后在线时间字段等字段，都将返回数据类型的零值。
	ReturnModel *bool `json:"ReturnModel,omitempty" name:"ReturnModel"`

	// 分页数量,0<取值范围<=100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移，取值＞0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 指定固件版本号，为空查询此产品下所有设备
	OtaVersion *string `json:"OtaVersion,omitempty" name:"OtaVersion"`

	// 设备名称，支持左前缀模糊匹配
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

type DescribeDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 是否返回全量数据
	// 当该值为false时，返回值中的设备物模型、固件版本、在线状态、最后在线时间字段等字段，都将返回数据类型的零值。
	ReturnModel *bool `json:"ReturnModel,omitempty" name:"ReturnModel"`

	// 分页数量,0<取值范围<=100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移，取值＞0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 指定固件版本号，为空查询此产品下所有设备
	OtaVersion *string `json:"OtaVersion,omitempty" name:"OtaVersion"`

	// 设备名称，支持左前缀模糊匹配
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

func (r *DescribeDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "ReturnModel")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OtaVersion")
	delete(f, "DeviceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDevicesResponseParams struct {
	// 设备信息 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*DevicesData `json:"Data,omitempty" name:"Data"`

	// 设备总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDevicesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDevicesResponseParams `json:"Response"`
}

func (r *DescribeDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIotDataTypeRequestParams struct {
	// 自定义数据类型的标识符，为空则返回全量自定义类型的列表
	TypeId *string `json:"TypeId,omitempty" name:"TypeId"`
}

type DescribeIotDataTypeRequest struct {
	*tchttp.BaseRequest
	
	// 自定义数据类型的标识符，为空则返回全量自定义类型的列表
	TypeId *string `json:"TypeId,omitempty" name:"TypeId"`
}

func (r *DescribeIotDataTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIotDataTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TypeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIotDataTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIotDataTypeResponseParams struct {
	// 自定义数据类型，json格式的字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*string `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIotDataTypeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIotDataTypeResponseParams `json:"Response"`
}

func (r *DescribeIotDataTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIotDataTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIotModelRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 物模型版本号， -1表示最新编辑的（未发布）
	Revision *int64 `json:"Revision,omitempty" name:"Revision"`
}

type DescribeIotModelRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 物模型版本号， -1表示最新编辑的（未发布）
	Revision *int64 `json:"Revision,omitempty" name:"Revision"`
}

func (r *DescribeIotModelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIotModelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "Revision")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIotModelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIotModelResponseParams struct {
	// 物模型定义，json格式的字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIotModelResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIotModelResponseParams `json:"Response"`
}

func (r *DescribeIotModelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIotModelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIotModelsRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

type DescribeIotModelsRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

func (r *DescribeIotModelsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIotModelsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIotModelsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIotModelsResponseParams struct {
	// 历史版本列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*IotModelData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIotModelsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIotModelsResponseParams `json:"Response"`
}

func (r *DescribeIotModelsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIotModelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogsRequestParams struct {
	// 设备TID
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 当前分页的最大条数,0<取值范围<=100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量,取值范围>0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 日志类型 1.在线状态变更 2.ProConst变更 3.ProWritable变更 4.Action控制 5.ProReadonly变更 6.Event事件
	LogType *uint64 `json:"LogType,omitempty" name:"LogType"`

	// 查询的起始时间 UNIX时间戳，单位秒
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 物模型对象索引，用于模糊查询，字符长度<=255，每层节点的字符长度<=16
	DataObject *string `json:"DataObject,omitempty" name:"DataObject"`

	// 查询的结束时间 UNIX时间戳，单位秒
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeLogsRequest struct {
	*tchttp.BaseRequest
	
	// 设备TID
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 当前分页的最大条数,0<取值范围<=100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量,取值范围>0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 日志类型 1.在线状态变更 2.ProConst变更 3.ProWritable变更 4.Action控制 5.ProReadonly变更 6.Event事件
	LogType *uint64 `json:"LogType,omitempty" name:"LogType"`

	// 查询的起始时间 UNIX时间戳，单位秒
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 物模型对象索引，用于模糊查询，字符长度<=255，每层节点的字符长度<=16
	DataObject *string `json:"DataObject,omitempty" name:"DataObject"`

	// 查询的结束时间 UNIX时间戳，单位秒
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Tid")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "LogType")
	delete(f, "StartTime")
	delete(f, "DataObject")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogsResponseParams struct {
	// 设备日志信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*LogData `json:"Data,omitempty" name:"Data"`

	// Data数组所包含的信息条数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLogsResponseParams `json:"Response"`
}

func (r *DescribeLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMessageQueueRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

type DescribeMessageQueueRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

func (r *DescribeMessageQueueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMessageQueueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMessageQueueRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMessageQueueResponseParams struct {
	// 消息队列配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *MsgQueueData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeMessageQueueResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMessageQueueResponseParams `json:"Response"`
}

func (r *DescribeMessageQueueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMessageQueueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelDataRetRequestParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeModelDataRetRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeModelDataRetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelDataRetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelDataRetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelDataRetResponseParams struct {
	// 设备响应结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeModelDataRetResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelDataRetResponseParams `json:"Response"`
}

func (r *DescribeModelDataRetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelDataRetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOsListRequestParams struct {

}

type DescribeOsListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeOsListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOsListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOsListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOsListResponseParams struct {
	// 系统类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *SystemType `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeOsListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOsListResponseParams `json:"Response"`
}

func (r *DescribeOsListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOsListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOtaVersionsRequestParams struct {
	// 分页偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页数量，0<取值范围<=100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 产品ID，为空时查询客户所有产品的版本信息
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 版本号，支持模糊匹配
	OtaVersion *string `json:"OtaVersion,omitempty" name:"OtaVersion"`

	// 版本类型 1未发布 2测试发布 3正式发布 4禁用
	PubStatus *uint64 `json:"PubStatus,omitempty" name:"PubStatus"`
}

type DescribeOtaVersionsRequest struct {
	*tchttp.BaseRequest
	
	// 分页偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页数量，0<取值范围<=100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 产品ID，为空时查询客户所有产品的版本信息
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 版本号，支持模糊匹配
	OtaVersion *string `json:"OtaVersion,omitempty" name:"OtaVersion"`

	// 版本类型 1未发布 2测试发布 3正式发布 4禁用
	PubStatus *uint64 `json:"PubStatus,omitempty" name:"PubStatus"`
}

func (r *DescribeOtaVersionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOtaVersionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ProductId")
	delete(f, "OtaVersion")
	delete(f, "PubStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOtaVersionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOtaVersionsResponseParams struct {
	// 版本数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 版本详细信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*VersionData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeOtaVersionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOtaVersionsResponseParams `json:"Response"`
}

func (r *DescribeOtaVersionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOtaVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

type DescribeProductRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

func (r *DescribeProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductResponseParams struct {
	// 产品详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *ProductData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeProductResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProductResponseParams `json:"Response"`
}

func (r *DescribeProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductsRequestParams struct {
	// 分页大小，当前页面中显示的最大数量，值范围 1-100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移，Offset从0开始
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 产器型号(APP产品,为APP包名)
	ProductModel *string `json:"ProductModel,omitempty" name:"ProductModel"`

	// 开始时间 ，UNIX 时间戳，单位秒
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间 ，UNIX 时间戳，单位秒
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeProductsRequest struct {
	*tchttp.BaseRequest
	
	// 分页大小，当前页面中显示的最大数量，值范围 1-100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移，Offset从0开始
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 产器型号(APP产品,为APP包名)
	ProductModel *string `json:"ProductModel,omitempty" name:"ProductModel"`

	// 开始时间 ，UNIX 时间戳，单位秒
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间 ，UNIX 时间戳，单位秒
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeProductsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "ProductModel")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProductsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductsResponseParams struct {
	// 产品详细信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*ProductData `json:"Data,omitempty" name:"Data"`

	// 产品总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeProductsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProductsResponseParams `json:"Response"`
}

func (r *DescribeProductsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePubVersionsRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

type DescribePubVersionsRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

func (r *DescribePubVersionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePubVersionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePubVersionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePubVersionsResponseParams struct {
	// 历史发布的版本列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*OtaPubHistory `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePubVersionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePubVersionsResponseParams `json:"Response"`
}

func (r *DescribePubVersionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePubVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRechargeRecordsRequestParams struct {
	// 账户类型 1:设备接入 2:云存。
	AccountType *int64 `json:"AccountType,omitempty" name:"AccountType"`

	// 从第几条记录开始显示, 默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 总共查询多少条记录，默认为值50。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeRechargeRecordsRequest struct {
	*tchttp.BaseRequest
	
	// 账户类型 1:设备接入 2:云存。
	AccountType *int64 `json:"AccountType,omitempty" name:"AccountType"`

	// 从第几条记录开始显示, 默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 总共查询多少条记录，默认为值50。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeRechargeRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRechargeRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccountType")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRechargeRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRechargeRecordsResponseParams struct {
	// 账户类型 1:设备接入 2:云存
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountType *int64 `json:"AccountType,omitempty" name:"AccountType"`

	// 充值记录列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Records []*RechargeRecord `json:"Records,omitempty" name:"Records"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRechargeRecordsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRechargeRecordsResponseParams `json:"Response"`
}

func (r *DescribeRechargeRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRechargeRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegistrationStatusRequestParams struct {
	// 终端用户的唯一ID列表，0<元素数量<=100
	CunionIds []*string `json:"CunionIds,omitempty" name:"CunionIds"`
}

type DescribeRegistrationStatusRequest struct {
	*tchttp.BaseRequest
	
	// 终端用户的唯一ID列表，0<元素数量<=100
	CunionIds []*string `json:"CunionIds,omitempty" name:"CunionIds"`
}

func (r *DescribeRegistrationStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegistrationStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CunionIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRegistrationStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegistrationStatusResponseParams struct {
	// 终端用户注册状态列表
	Data []*RegisteredStatus `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRegistrationStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRegistrationStatusResponseParams `json:"Response"`
}

func (r *DescribeRegistrationStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegistrationStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRunLogRequestParams struct {
	// 设备TID
	Tid *string `json:"Tid,omitempty" name:"Tid"`
}

type DescribeRunLogRequest struct {
	*tchttp.BaseRequest
	
	// 设备TID
	Tid *string `json:"Tid,omitempty" name:"Tid"`
}

func (r *DescribeRunLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRunLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Tid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRunLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRunLogResponseParams struct {
	// 设备运行日志文本信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRunLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRunLogResponseParams `json:"Response"`
}

func (r *DescribeRunLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRunLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStorageServiceRequestParams struct {
	// 云存服务ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 是否返回已结束的订单信息(已过期/已退订/已转移)
	GetFinishedOrder *bool `json:"GetFinishedOrder,omitempty" name:"GetFinishedOrder"`
}

type DescribeStorageServiceRequest struct {
	*tchttp.BaseRequest
	
	// 云存服务ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 是否返回已结束的订单信息(已过期/已退订/已转移)
	GetFinishedOrder *bool `json:"GetFinishedOrder,omitempty" name:"GetFinishedOrder"`
}

func (r *DescribeStorageServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStorageServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "GetFinishedOrder")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStorageServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStorageServiceResponseParams struct {
	// 云存服务ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 云存服务所在的区域
	StorageRegion *string `json:"StorageRegion,omitempty" name:"StorageRegion"`

	// 设备TID
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 视频流通道号。(对于存在多路视频流的设备，如NVR设备，与设备实际视频流通道号对应)
	ChnNum *int64 `json:"ChnNum,omitempty" name:"ChnNum"`

	// 终端用户在IoT Video平台的注册ID
	AccessId *string `json:"AccessId,omitempty" name:"AccessId"`

	// 服务开始时间
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 服务失效时间
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 服务状态
	// 1：正常使用中
	// 2：待续费。设备云存服务已到期，但是历史云存数据未过期。续费后仍可查看这些历史数据。
	// 3：已过期。查询不到设备保存在云端的数据。
	// 4：等待服务生效。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 云存定单列表
	Data []*StorageOrder `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeStorageServiceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStorageServiceResponseParams `json:"Response"`
}

func (r *DescribeStorageServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStorageServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamRequestParams struct {
	// 设备TID
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 终端用户ID
	AccessId *string `json:"AccessId,omitempty" name:"AccessId"`

	// 直播协议, 可选值：RTSP、RTMP、HLS、HLS-fmp4
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 音视频流地址
	Address *string `json:"Address,omitempty" name:"Address"`

	// 设备访问token，访问用户未绑定的设备时，需提供该参数
	AccessToken *string `json:"AccessToken,omitempty" name:"AccessToken"`
}

type DescribeStreamRequest struct {
	*tchttp.BaseRequest
	
	// 设备TID
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 终端用户ID
	AccessId *string `json:"AccessId,omitempty" name:"AccessId"`

	// 直播协议, 可选值：RTSP、RTMP、HLS、HLS-fmp4
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 音视频流地址
	Address *string `json:"Address,omitempty" name:"Address"`

	// 设备访问token，访问用户未绑定的设备时，需提供该参数
	AccessToken *string `json:"AccessToken,omitempty" name:"AccessToken"`
}

func (r *DescribeStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Tid")
	delete(f, "AccessId")
	delete(f, "Protocol")
	delete(f, "Address")
	delete(f, "AccessToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamResponseParams struct {
	// 返回参数结构
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *Data `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeStreamResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamResponseParams `json:"Response"`
}

func (r *DescribeStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTraceIdsRequestParams struct {

}

type DescribeTraceIdsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeTraceIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTraceIdsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTraceIdsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTraceIdsResponseParams struct {
	// 设备TID列表，列表元素之间以“,”分隔
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTraceIdsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTraceIdsResponseParams `json:"Response"`
}

func (r *DescribeTraceIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTraceIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTraceStatusRequestParams struct {
	// 设备TID列表
	Tids []*string `json:"Tids,omitempty" name:"Tids"`
}

type DescribeTraceStatusRequest struct {
	*tchttp.BaseRequest
	
	// 设备TID列表
	Tids []*string `json:"Tids,omitempty" name:"Tids"`
}

func (r *DescribeTraceStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTraceStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Tids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTraceStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTraceStatusResponseParams struct {
	// 设备追踪状态列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*TraceStatus `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTraceStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTraceStatusResponseParams `json:"Response"`
}

func (r *DescribeTraceStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTraceStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DevTokenInfo struct {
	// 客户的终端用户在IotVideo上的唯一标识id
	AccessId *string `json:"AccessId,omitempty" name:"AccessId"`

	// 设备TID
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// IotVideo平台的accessToken
	AccessToken *string `json:"AccessToken,omitempty" name:"AccessToken"`

	// Token的过期时间，单位秒(UTC时间)
	ExpireTime *uint64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
}

type DeviceCertificate struct {
	// 设备TID
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 设备初始证书信息，base64编码
	Certificate *string `json:"Certificate,omitempty" name:"Certificate"`

	// 设备私钥下载地址
	WhiteBoxSoUrl *string `json:"WhiteBoxSoUrl,omitempty" name:"WhiteBoxSoUrl"`
}

type DeviceData struct {
	// 设备TID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 激活时间 0代表未激活
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActiveTime *uint64 `json:"ActiveTime,omitempty" name:"ActiveTime"`

	// 设备是否被禁用
	// 注意：此字段可能返回 null，表示取不到有效值。
	Disabled *bool `json:"Disabled,omitempty" name:"Disabled"`

	// 固件版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	OtaVersion *string `json:"OtaVersion,omitempty" name:"OtaVersion"`

	// 设备在线状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Online *uint64 `json:"Online,omitempty" name:"Online"`

	// 设备最后上线时间（mqtt连接成功时间），UNIX时间戳，单位秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastOnlineTime *uint64 `json:"LastOnlineTime,omitempty" name:"LastOnlineTime"`

	// 物模型json数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	IotModel *string `json:"IotModel,omitempty" name:"IotModel"`

	// 设备名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 产品ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备初始证书信息，base64编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Certificate *string `json:"Certificate,omitempty" name:"Certificate"`

	// 设备私钥下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	WhiteBoxSoUrl *string `json:"WhiteBoxSoUrl,omitempty" name:"WhiteBoxSoUrl"`

	// 设备推流状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	StreamStatus *bool `json:"StreamStatus,omitempty" name:"StreamStatus"`
}

type DeviceModelData struct {
	// 设备TID
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 物模型分支路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Branch *string `json:"Branch,omitempty" name:"Branch"`

	// 物模型数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	IotModel *string `json:"IotModel,omitempty" name:"IotModel"`
}

type DevicesData struct {
	// 设备TID
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 激活时间 0代表未激活
	ActiveTime *uint64 `json:"ActiveTime,omitempty" name:"ActiveTime"`

	// 设备是否被禁用
	Disabled *bool `json:"Disabled,omitempty" name:"Disabled"`

	// 设备推流状态
	StreamStatus *bool `json:"StreamStatus,omitempty" name:"StreamStatus"`

	// 固件版本
	OtaVersion *string `json:"OtaVersion,omitempty" name:"OtaVersion"`

	// 设备在线状态
	Online *uint64 `json:"Online,omitempty" name:"Online"`

	// 设备最后上线时间（mqtt连接成功时间），UNIX时间戳，单位秒
	LastOnlineTime *uint64 `json:"LastOnlineTime,omitempty" name:"LastOnlineTime"`

	// 物模型json数据
	IotModel *string `json:"IotModel,omitempty" name:"IotModel"`

	// 设备固件最新更新时间，UNIX时间戳，单位秒
	LastUpdateTime *uint64 `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
}

// Predefined struct for user
type DisableDeviceRequestParams struct {
	// 设备TID ≤100
	Tids []*string `json:"Tids,omitempty" name:"Tids"`
}

type DisableDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 设备TID ≤100
	Tids []*string `json:"Tids,omitempty" name:"Tids"`
}

func (r *DisableDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Tids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableDeviceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DisableDeviceResponse struct {
	*tchttp.BaseResponse
	Response *DisableDeviceResponseParams `json:"Response"`
}

func (r *DisableDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableDeviceStreamRequestParams struct {
	// 设备TID列表
	Tids []*string `json:"Tids,omitempty" name:"Tids"`
}

type DisableDeviceStreamRequest struct {
	*tchttp.BaseRequest
	
	// 设备TID列表
	Tids []*string `json:"Tids,omitempty" name:"Tids"`
}

func (r *DisableDeviceStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableDeviceStreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Tids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableDeviceStreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableDeviceStreamResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DisableDeviceStreamResponse struct {
	*tchttp.BaseResponse
	Response *DisableDeviceStreamResponseParams `json:"Response"`
}

func (r *DisableDeviceStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableDeviceStreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableOtaVersionRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 固件版本号，格式为x.y.z， x，y 范围0-63，z范围1~524288
	OtaVersion *string `json:"OtaVersion,omitempty" name:"OtaVersion"`

	// 操作人
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type DisableOtaVersionRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 固件版本号，格式为x.y.z， x，y 范围0-63，z范围1~524288
	OtaVersion *string `json:"OtaVersion,omitempty" name:"OtaVersion"`

	// 操作人
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *DisableOtaVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableOtaVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "OtaVersion")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableOtaVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableOtaVersionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DisableOtaVersionResponse struct {
	*tchttp.BaseResponse
	Response *DisableOtaVersionResponseParams `json:"Response"`
}

func (r *DisableOtaVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableOtaVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IotModelData struct {
	// 版本号
	Revision *int64 `json:"Revision,omitempty" name:"Revision"`

	// 发布时间
	ReleaseTime *int64 `json:"ReleaseTime,omitempty" name:"ReleaseTime"`
}

type LogData struct {
	// 发生时间 UNIX时间戳，单位秒
	Occurtime *uint64 `json:"Occurtime,omitempty" name:"Occurtime"`

	// 日志类型 1在线状态变更 2FP变更 3SP变更 4CO控制 5ST变更 6EV事件
	LogType *uint64 `json:"LogType,omitempty" name:"LogType"`

	// 物模型对象索引
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataObject *string `json:"DataObject,omitempty" name:"DataObject"`

	// 物模型旧值  json串
	// 注意：此字段可能返回 null，表示取不到有效值。
	OldValue *string `json:"OldValue,omitempty" name:"OldValue"`

	// 物模型新值  json串
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewValue *string `json:"NewValue,omitempty" name:"NewValue"`
}

// Predefined struct for user
type ModifyDeviceActionRequestParams struct {
	// 设备TID
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 如果设备处于休眠状态，是否唤醒设备
	Wakeup *bool `json:"Wakeup,omitempty" name:"Wakeup"`

	// 物模型的分支路径
	Branch *string `json:"Branch,omitempty" name:"Branch"`

	// 写入的物模型数据，如果是json需要转义成字符串
	Value *string `json:"Value,omitempty" name:"Value"`

	// Value字段的类型是否为数值（float、int）
	IsNum *bool `json:"IsNum,omitempty" name:"IsNum"`
}

type ModifyDeviceActionRequest struct {
	*tchttp.BaseRequest
	
	// 设备TID
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 如果设备处于休眠状态，是否唤醒设备
	Wakeup *bool `json:"Wakeup,omitempty" name:"Wakeup"`

	// 物模型的分支路径
	Branch *string `json:"Branch,omitempty" name:"Branch"`

	// 写入的物模型数据，如果是json需要转义成字符串
	Value *string `json:"Value,omitempty" name:"Value"`

	// Value字段的类型是否为数值（float、int）
	IsNum *bool `json:"IsNum,omitempty" name:"IsNum"`
}

func (r *ModifyDeviceActionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDeviceActionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Tid")
	delete(f, "Wakeup")
	delete(f, "Branch")
	delete(f, "Value")
	delete(f, "IsNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDeviceActionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDeviceActionResponseParams struct {
	// 设备端的响应结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitempty" name:"Data"`

	// 任务ID
	// 若设备端未能及时响应时，会返回此字段，用户可以通过DescribeModelDataRet获取设备的最终响应结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDeviceActionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDeviceActionResponseParams `json:"Response"`
}

func (r *ModifyDeviceActionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDeviceActionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDevicePropertyRequestParams struct {
	// 设备TID
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 如果设备处于休眠状态，是否唤醒设备
	Wakeup *bool `json:"Wakeup,omitempty" name:"Wakeup"`

	// 物模型的分支路径
	Branch *string `json:"Branch,omitempty" name:"Branch"`

	// 写入的物模型数据，如果是json需要转义成字符串
	Value *string `json:"Value,omitempty" name:"Value"`

	// Value字段是否为数值（float、int）
	IsNum *bool `json:"IsNum,omitempty" name:"IsNum"`
}

type ModifyDevicePropertyRequest struct {
	*tchttp.BaseRequest
	
	// 设备TID
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 如果设备处于休眠状态，是否唤醒设备
	Wakeup *bool `json:"Wakeup,omitempty" name:"Wakeup"`

	// 物模型的分支路径
	Branch *string `json:"Branch,omitempty" name:"Branch"`

	// 写入的物模型数据，如果是json需要转义成字符串
	Value *string `json:"Value,omitempty" name:"Value"`

	// Value字段是否为数值（float、int）
	IsNum *bool `json:"IsNum,omitempty" name:"IsNum"`
}

func (r *ModifyDevicePropertyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDevicePropertyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Tid")
	delete(f, "Wakeup")
	delete(f, "Branch")
	delete(f, "Value")
	delete(f, "IsNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDevicePropertyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDevicePropertyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDevicePropertyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDevicePropertyResponseParams `json:"Response"`
}

func (r *ModifyDevicePropertyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDevicePropertyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDeviceRequestParams struct {
	// 设备ID
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 用户ID
	AccessId *string `json:"AccessId,omitempty" name:"AccessId"`

	// 设备昵称，最多不超过64个字符
	Nick *string `json:"Nick,omitempty" name:"Nick"`
}

type ModifyDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 设备ID
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 用户ID
	AccessId *string `json:"AccessId,omitempty" name:"AccessId"`

	// 设备昵称，最多不超过64个字符
	Nick *string `json:"Nick,omitempty" name:"Nick"`
}

func (r *ModifyDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Tid")
	delete(f, "AccessId")
	delete(f, "Nick")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDeviceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDeviceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDeviceResponseParams `json:"Response"`
}

func (r *ModifyDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProductRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 产品名称
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 产品描述
	ProductDescription *string `json:"ProductDescription,omitempty" name:"ProductDescription"`

	// 主芯片产商ID
	ChipManufactureId *string `json:"ChipManufactureId,omitempty" name:"ChipManufactureId"`

	// 主芯片ID
	ChipId *string `json:"ChipId,omitempty" name:"ChipId"`
}

type ModifyProductRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 产品名称
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 产品描述
	ProductDescription *string `json:"ProductDescription,omitempty" name:"ProductDescription"`

	// 主芯片产商ID
	ChipManufactureId *string `json:"ChipManufactureId,omitempty" name:"ChipManufactureId"`

	// 主芯片ID
	ChipId *string `json:"ChipId,omitempty" name:"ChipId"`
}

func (r *ModifyProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "ProductName")
	delete(f, "ProductDescription")
	delete(f, "ChipManufactureId")
	delete(f, "ChipId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProductResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyProductResponse struct {
	*tchttp.BaseResponse
	Response *ModifyProductResponseParams `json:"Response"`
}

func (r *ModifyProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVerContentRequestParams struct {
	// 产品id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 需要修改的版本号
	OtaVersion *string `json:"OtaVersion,omitempty" name:"OtaVersion"`

	// 操作人,字符长度<=64
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 备注信息
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 版本发布的描述信息，需要国际化，可以为空
	Contents *Contents `json:"Contents,omitempty" name:"Contents"`
}

type ModifyVerContentRequest struct {
	*tchttp.BaseRequest
	
	// 产品id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 需要修改的版本号
	OtaVersion *string `json:"OtaVersion,omitempty" name:"OtaVersion"`

	// 操作人,字符长度<=64
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 备注信息
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 版本发布的描述信息，需要国际化，可以为空
	Contents *Contents `json:"Contents,omitempty" name:"Contents"`
}

func (r *ModifyVerContentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVerContentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "OtaVersion")
	delete(f, "Operator")
	delete(f, "Remark")
	delete(f, "Contents")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVerContentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVerContentResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyVerContentResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVerContentResponseParams `json:"Response"`
}

func (r *ModifyVerContentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVerContentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MsgQueueData struct {
	// 消息队列类型 1：CMQ 2：kafka
	MsgQueueType *uint64 `json:"MsgQueueType,omitempty" name:"MsgQueueType"`

	// 消息类型列表，整型值（0-31）之间以“,”分隔
	MsgType *string `json:"MsgType,omitempty" name:"MsgType"`

	// 主题名称
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// 实例名称
	Instance *string `json:"Instance,omitempty" name:"Instance"`

	// 消息地域
	MsgRegion *string `json:"MsgRegion,omitempty" name:"MsgRegion"`
}

type OsData struct {
	// 芯片型号
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChipId *string `json:"ChipId,omitempty" name:"ChipId"`

	// 芯片厂商
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChipManufacture *string `json:"ChipManufacture,omitempty" name:"ChipManufacture"`
}

type OtaPubHistory struct {
	// 版本名称
	OtaVersion *string `json:"OtaVersion,omitempty" name:"OtaVersion"`

	// 发布时间，unix时间戳，单位：秒
	PublishTime *uint64 `json:"PublishTime,omitempty" name:"PublishTime"`
}

type ProductBase struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 产器型号(APP产品,为APP包名)
	ProductModel *string `json:"ProductModel,omitempty" name:"ProductModel"`

	// 产品名称
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 产品描述信息
	ProductDescription *string `json:"ProductDescription,omitempty" name:"ProductDescription"`

	// 创建时间，UNIX 时间戳，单位秒
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 物模型发布版本号,0代表物模型尚未发布
	IotModelRevision *uint64 `json:"IotModelRevision,omitempty" name:"IotModelRevision"`

	// 产品密钥
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`

	// 设备功能码
	// ypsxth : 音频双向通话;	
	// spdxth : 视频单向通话(监控);
	// NVR0824 : NVR设备,大于8路，小于等于24路;
	// WifiKeepalive : Wifi保活(低功耗产品);
	// Alexa : Alexa接入;
	// Google : Google接入;
	// 注意：此字段可能返回 null，表示取不到有效值。
	FuncCode []*string `json:"FuncCode,omitempty" name:"FuncCode"`

	// 产品类别，0 : 普通视频设备；1 : NVR设备
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductCate *int64 `json:"ProductCate,omitempty" name:"ProductCate"`

	// 产品地域
	// China-Mainland（中国大陆）
	// China-Hong Kong, Macao and Taiwan（港澳台地区）
	// America（美国）
	// Europe（欧洲）
	// India（印度）
	// Other-Overseas（其他境外地区）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductRegion *string `json:"ProductRegion,omitempty" name:"ProductRegion"`
}

type ProductData struct {
	// 产品ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 产品名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 产品描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductDescription *string `json:"ProductDescription,omitempty" name:"ProductDescription"`

	// 创建时间，UNIX 时间戳，单位秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 物模型发布版本号,0代表物模型尚未发布
	// 注意：此字段可能返回 null，表示取不到有效值。
	IotModelRevision *int64 `json:"IotModelRevision,omitempty" name:"IotModelRevision"`

	// 产品密钥
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`

	// 设备功能码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Features []*string `json:"Features,omitempty" name:"Features"`

	// 产器型号(APP产品,为APP包名)
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductModel *string `json:"ProductModel,omitempty" name:"ProductModel"`

	// 主芯片厂商id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChipManufactureId *string `json:"ChipManufactureId,omitempty" name:"ChipManufactureId"`

	// 主芯片型号
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChipId *string `json:"ChipId,omitempty" name:"ChipId"`

	// 产品类别，0：普通视频设备；1：NVR设备
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductCate *int64 `json:"ProductCate,omitempty" name:"ProductCate"`

	// 产品地区
	// China-Mainland（中国大陆）
	// China-Hong Kong, Macao and Taiwan（港澳台地区）
	// America（美国）
	// Europe（欧洲）
	// India（印度）
	// Other-Overseas（其他境外地区）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductRegion *string `json:"ProductRegion,omitempty" name:"ProductRegion"`

	// 接入模型，bit0是0：公版小程序未接入，bit0是1：公版小程序已接入
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessMode *int64 `json:"AccessMode,omitempty" name:"AccessMode"`

	// linux,android,liteos
	// 注意：此字段可能返回 null，表示取不到有效值。
	Os *string `json:"Os,omitempty" name:"Os"`
}

type RechargeRecord struct {
	// 流水记录号。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WaterId *int64 `json:"WaterId,omitempty" name:"WaterId"`

	// 充值前的余额，单位0.01元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BalanceBeforeRecharge *int64 `json:"BalanceBeforeRecharge,omitempty" name:"BalanceBeforeRecharge"`

	// 充值金额，单位0.01元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Money *int64 `json:"Money,omitempty" name:"Money"`

	// 充值时间, UTC值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperateTime *int64 `json:"OperateTime,omitempty" name:"OperateTime"`
}

// Predefined struct for user
type RefundStorageServiceRequestParams struct {
	// 云存服务ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 云存子订单ID。如果指定子订单ID,则仅退订该子订单，如果未指定子定单ID，则退订所有子订单
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`
}

type RefundStorageServiceRequest struct {
	*tchttp.BaseRequest
	
	// 云存服务ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 云存子订单ID。如果指定子订单ID,则仅退订该子订单，如果未指定子定单ID，则退订所有子订单
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`
}

func (r *RefundStorageServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefundStorageServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "OrderId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RefundStorageServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RefundStorageServiceResponseParams struct {
	// 云存服务ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 云存服务所在的区域
	StorageRegion *string `json:"StorageRegion,omitempty" name:"StorageRegion"`

	// 设备TID
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 视频流通道号。(对于存在多路视频流的设备，如NVR设备，与设备实际视频流通道号对应)
	ChnNum *int64 `json:"ChnNum,omitempty" name:"ChnNum"`

	// 终端用户在IoT Video平台的注册ID
	AccessId *string `json:"AccessId,omitempty" name:"AccessId"`

	// 服务开始时间
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 服务失效时间
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 服务状态
	// 1：正常使用中
	// 2：待续费。设备云存服务已到期，但是历史云存数据未过期。续费后仍可查看这些历史数据。
	// 3：已过期。查询不到设备保存在云端的数据。
	// 4：等待服务生效。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 有效云存定单列表
	Data []*StorageOrder `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RefundStorageServiceResponse struct {
	*tchttp.BaseResponse
	Response *RefundStorageServiceResponseParams `json:"Response"`
}

func (r *RefundStorageServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefundStorageServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegisteredStatus struct {
	// 终端用户的唯一ID
	CunionId *string `json:"CunionId,omitempty" name:"CunionId"`

	// 注册状态
	IsRegisted *bool `json:"IsRegisted,omitempty" name:"IsRegisted"`
}

// Predefined struct for user
type RunDeviceRequestParams struct {
	// TID列表 ≤100
	Tids []*string `json:"Tids,omitempty" name:"Tids"`
}

type RunDeviceRequest struct {
	*tchttp.BaseRequest
	
	// TID列表 ≤100
	Tids []*string `json:"Tids,omitempty" name:"Tids"`
}

func (r *RunDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Tids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunDeviceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RunDeviceResponse struct {
	*tchttp.BaseResponse
	Response *RunDeviceResponseParams `json:"Response"`
}

func (r *RunDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunDeviceStreamRequestParams struct {
	// 设备TID 列表
	Tids []*string `json:"Tids,omitempty" name:"Tids"`
}

type RunDeviceStreamRequest struct {
	*tchttp.BaseRequest
	
	// 设备TID 列表
	Tids []*string `json:"Tids,omitempty" name:"Tids"`
}

func (r *RunDeviceStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunDeviceStreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Tids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunDeviceStreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunDeviceStreamResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RunDeviceStreamResponse struct {
	*tchttp.BaseResponse
	Response *RunDeviceStreamResponseParams `json:"Response"`
}

func (r *RunDeviceStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunDeviceStreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunIotModelRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 物模型定义，json格式的字符串
	IotModel *string `json:"IotModel,omitempty" name:"IotModel"`
}

type RunIotModelRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 物模型定义，json格式的字符串
	IotModel *string `json:"IotModel,omitempty" name:"IotModel"`
}

func (r *RunIotModelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunIotModelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "IotModel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunIotModelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunIotModelResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RunIotModelResponse struct {
	*tchttp.BaseResponse
	Response *RunIotModelResponseParams `json:"Response"`
}

func (r *RunIotModelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunIotModelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunOtaVersionRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 固件版本号，格式为x.y.z， x，y 范围0-63，z范围1~524288
	OtaVersion *string `json:"OtaVersion,omitempty" name:"OtaVersion"`

	// 灰度值,取值范围0-100，为0时相当于暂停发布
	GrayValue *uint64 `json:"GrayValue,omitempty" name:"GrayValue"`

	// 指定的旧版本
	OldVersions []*string `json:"OldVersions,omitempty" name:"OldVersions"`

	// 操作人
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 备注信息
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 版本发布的描述信息，需要国际化，可以为空
	Contents *Contents `json:"Contents,omitempty" name:"Contents"`
}

type RunOtaVersionRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 固件版本号，格式为x.y.z， x，y 范围0-63，z范围1~524288
	OtaVersion *string `json:"OtaVersion,omitempty" name:"OtaVersion"`

	// 灰度值,取值范围0-100，为0时相当于暂停发布
	GrayValue *uint64 `json:"GrayValue,omitempty" name:"GrayValue"`

	// 指定的旧版本
	OldVersions []*string `json:"OldVersions,omitempty" name:"OldVersions"`

	// 操作人
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 备注信息
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 版本发布的描述信息，需要国际化，可以为空
	Contents *Contents `json:"Contents,omitempty" name:"Contents"`
}

func (r *RunOtaVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunOtaVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "OtaVersion")
	delete(f, "GrayValue")
	delete(f, "OldVersions")
	delete(f, "Operator")
	delete(f, "Remark")
	delete(f, "Contents")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunOtaVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunOtaVersionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RunOtaVersionResponse struct {
	*tchttp.BaseResponse
	Response *RunOtaVersionResponseParams `json:"Response"`
}

func (r *RunOtaVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunOtaVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunTestOtaVersionRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 固件版本号，格式为x.y.z， x，y 范围0-63，z范围1~524288
	OtaVersion *string `json:"OtaVersion,omitempty" name:"OtaVersion"`

	// 指定可升级的设备TID
	Tids []*string `json:"Tids,omitempty" name:"Tids"`

	// 操作人
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 备注信息
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type RunTestOtaVersionRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 固件版本号，格式为x.y.z， x，y 范围0-63，z范围1~524288
	OtaVersion *string `json:"OtaVersion,omitempty" name:"OtaVersion"`

	// 指定可升级的设备TID
	Tids []*string `json:"Tids,omitempty" name:"Tids"`

	// 操作人
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 备注信息
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *RunTestOtaVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunTestOtaVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "OtaVersion")
	delete(f, "Tids")
	delete(f, "Operator")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunTestOtaVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunTestOtaVersionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RunTestOtaVersionResponse struct {
	*tchttp.BaseResponse
	Response *RunTestOtaVersionResponseParams `json:"Response"`
}

func (r *RunTestOtaVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunTestOtaVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendOnlineMsgRequestParams struct {
	// 设备TID
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 如果设备处于休眠状态，是否唤醒设备
	Wakeup *bool `json:"Wakeup,omitempty" name:"Wakeup"`

	// 等待回应类型
	// 0：不等待设备回应直接响应请求;
	// 1：要求设备确认消息已接收,或等待超时后返回;
	// 2：要求设备进行响应处理,收到设备的响应数据后,将设备响应数据回应给请求方;
	WaitResp *uint64 `json:"WaitResp,omitempty" name:"WaitResp"`

	// 消息主题
	MsgTopic *string `json:"MsgTopic,omitempty" name:"MsgTopic"`

	// 消息内容，最大长度不超过8k字节
	MsgContent *string `json:"MsgContent,omitempty" name:"MsgContent"`
}

type SendOnlineMsgRequest struct {
	*tchttp.BaseRequest
	
	// 设备TID
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 如果设备处于休眠状态，是否唤醒设备
	Wakeup *bool `json:"Wakeup,omitempty" name:"Wakeup"`

	// 等待回应类型
	// 0：不等待设备回应直接响应请求;
	// 1：要求设备确认消息已接收,或等待超时后返回;
	// 2：要求设备进行响应处理,收到设备的响应数据后,将设备响应数据回应给请求方;
	WaitResp *uint64 `json:"WaitResp,omitempty" name:"WaitResp"`

	// 消息主题
	MsgTopic *string `json:"MsgTopic,omitempty" name:"MsgTopic"`

	// 消息内容，最大长度不超过8k字节
	MsgContent *string `json:"MsgContent,omitempty" name:"MsgContent"`
}

func (r *SendOnlineMsgRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendOnlineMsgRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Tid")
	delete(f, "Wakeup")
	delete(f, "WaitResp")
	delete(f, "MsgTopic")
	delete(f, "MsgContent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendOnlineMsgRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendOnlineMsgResponseParams struct {
	// 若返回此项则表明需要用户用此taskID进行查询请求是否成功(只有waitresp不等于0的情况下才可能会返回该taskID项)
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 设备响应信息
	Data *string `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SendOnlineMsgResponse struct {
	*tchttp.BaseResponse
	Response *SendOnlineMsgResponseParams `json:"Response"`
}

func (r *SendOnlineMsgResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendOnlineMsgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetMessageQueueRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 消息队列类型 1-CMQ; 2-Ckafka
	MsgQueueType *uint64 `json:"MsgQueueType,omitempty" name:"MsgQueueType"`

	// 消息类型,整型值（0-31）之间以“,”分隔
	// 0.设备在线状态变更
	// 1.常亮属性(ProConst)变更
	// 2.可写属性(ProWritable)变更
	// 3.只读属性(ProReadonly)变更
	// 4.设备控制(Action)
	// 5.设备事件(Event)
	// 6.系统事件(System)
	MsgType *string `json:"MsgType,omitempty" name:"MsgType"`

	// 消息队列主题，不超过32字符
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// kafka消息队列的实例名，不超过64字符
	Instance *string `json:"Instance,omitempty" name:"Instance"`

	// 消息地域，不超过32字符
	MsgRegion *string `json:"MsgRegion,omitempty" name:"MsgRegion"`
}

type SetMessageQueueRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 消息队列类型 1-CMQ; 2-Ckafka
	MsgQueueType *uint64 `json:"MsgQueueType,omitempty" name:"MsgQueueType"`

	// 消息类型,整型值（0-31）之间以“,”分隔
	// 0.设备在线状态变更
	// 1.常亮属性(ProConst)变更
	// 2.可写属性(ProWritable)变更
	// 3.只读属性(ProReadonly)变更
	// 4.设备控制(Action)
	// 5.设备事件(Event)
	// 6.系统事件(System)
	MsgType *string `json:"MsgType,omitempty" name:"MsgType"`

	// 消息队列主题，不超过32字符
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// kafka消息队列的实例名，不超过64字符
	Instance *string `json:"Instance,omitempty" name:"Instance"`

	// 消息地域，不超过32字符
	MsgRegion *string `json:"MsgRegion,omitempty" name:"MsgRegion"`
}

func (r *SetMessageQueueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetMessageQueueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "MsgQueueType")
	delete(f, "MsgType")
	delete(f, "Topic")
	delete(f, "Instance")
	delete(f, "MsgRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetMessageQueueRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetMessageQueueResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetMessageQueueResponse struct {
	*tchttp.BaseResponse
	Response *SetMessageQueueResponseParams `json:"Response"`
}

func (r *SetMessageQueueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetMessageQueueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StorageOrder struct {
	// 定单唯一性ID
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 云存套餐ID
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`

	// 定单服务状态
	// 1;订单正在使用。
	// 2:订单未开始。
	// 3:订单已经使用过，现在暂时未开始使用(该订单从其他服务转移而来)。
	// 4:订单已过期。
	// 5:订单已被退订。
	// 6:定单已被转移到其他云存服务。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 定单服务生效时间
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 定单服务失效时间
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

type SystemType struct {
	// 安卓系统
	// 注意：此字段可能返回 null，表示取不到有效值。
	Android []*OsData `json:"Android,omitempty" name:"Android"`

	// linux系统
	// 注意：此字段可能返回 null，表示取不到有效值。
	Linux []*OsData `json:"Linux,omitempty" name:"Linux"`

	// LiteOs系统
	// 注意：此字段可能返回 null，表示取不到有效值。
	LiteOs []*OsData `json:"LiteOs,omitempty" name:"LiteOs"`
}

type TraceStatus struct {
	// 设备TID
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 设备追踪状态
	IsExist *bool `json:"IsExist,omitempty" name:"IsExist"`
}

// Predefined struct for user
type UpgradeDeviceRequestParams struct {
	// 设备TID
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 固件版本号
	OtaVersion *string `json:"OtaVersion,omitempty" name:"OtaVersion"`

	// 是否立即升级
	UpgradeNow *bool `json:"UpgradeNow,omitempty" name:"UpgradeNow"`
}

type UpgradeDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 设备TID
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 固件版本号
	OtaVersion *string `json:"OtaVersion,omitempty" name:"OtaVersion"`

	// 是否立即升级
	UpgradeNow *bool `json:"UpgradeNow,omitempty" name:"UpgradeNow"`
}

func (r *UpgradeDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Tid")
	delete(f, "OtaVersion")
	delete(f, "UpgradeNow")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeDeviceResponseParams struct {
	// 设备端返回的数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpgradeDeviceResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeDeviceResponseParams `json:"Response"`
}

func (r *UpgradeDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadOtaVersionRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 固件版本号，格式为x.y.z， x，y 范围0-63，z范围1~524288
	OtaVersion *string `json:"OtaVersion,omitempty" name:"OtaVersion"`

	// 固件版本URL
	VersionUrl *string `json:"VersionUrl,omitempty" name:"VersionUrl"`

	// 文件大小，单位：byte
	FileSize *uint64 `json:"FileSize,omitempty" name:"FileSize"`

	// 文件md5校验码（32字符）
	Md5 *string `json:"Md5,omitempty" name:"Md5"`

	// 操作人
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 备注信息
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 版本发布的描述信息，需要国际化，可以为空
	Contents *Contents `json:"Contents,omitempty" name:"Contents"`
}

type UploadOtaVersionRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 固件版本号，格式为x.y.z， x，y 范围0-63，z范围1~524288
	OtaVersion *string `json:"OtaVersion,omitempty" name:"OtaVersion"`

	// 固件版本URL
	VersionUrl *string `json:"VersionUrl,omitempty" name:"VersionUrl"`

	// 文件大小，单位：byte
	FileSize *uint64 `json:"FileSize,omitempty" name:"FileSize"`

	// 文件md5校验码（32字符）
	Md5 *string `json:"Md5,omitempty" name:"Md5"`

	// 操作人
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 备注信息
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 版本发布的描述信息，需要国际化，可以为空
	Contents *Contents `json:"Contents,omitempty" name:"Contents"`
}

func (r *UploadOtaVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadOtaVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "OtaVersion")
	delete(f, "VersionUrl")
	delete(f, "FileSize")
	delete(f, "Md5")
	delete(f, "Operator")
	delete(f, "Remark")
	delete(f, "Contents")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadOtaVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadOtaVersionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UploadOtaVersionResponse struct {
	*tchttp.BaseResponse
	Response *UploadOtaVersionResponseParams `json:"Response"`
}

func (r *UploadOtaVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadOtaVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VersionData struct {
	// 产品ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 固件版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OtaVersion *string `json:"OtaVersion,omitempty" name:"OtaVersion"`

	// 版本类型 1未发布 2测试发布 3正式发布 4禁用
	// 注意：此字段可能返回 null，表示取不到有效值。
	PubStatus *uint64 `json:"PubStatus,omitempty" name:"PubStatus"`

	// 固件版本存储路径URL
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionUrl *string `json:"VersionUrl,omitempty" name:"VersionUrl"`

	// 文件大小，byte
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileSize *uint64 `json:"FileSize,omitempty" name:"FileSize"`

	// 文件校验码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Md5 *string `json:"Md5,omitempty" name:"Md5"`

	// 指定的允许升级的旧版本，PubStatus=3时有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	OldVersions *string `json:"OldVersions,omitempty" name:"OldVersions"`

	// 指定的允许升级的旧设备id，PubStatus=2时有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tids *string `json:"Tids,omitempty" name:"Tids"`

	// 灰度值（0-100）,PubStatus=3时有效，表示n%的升级总量
	// 注意：此字段可能返回 null，表示取不到有效值。
	GrayValue *uint64 `json:"GrayValue,omitempty" name:"GrayValue"`

	// 最近一次发布时间，UNIX时间戳，单位秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublishTime *uint64 `json:"PublishTime,omitempty" name:"PublishTime"`

	// 此版本激活的设备总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActiveCount *int64 `json:"ActiveCount,omitempty" name:"ActiveCount"`

	// 此版本在线的设备总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	OnlineCount *int64 `json:"OnlineCount,omitempty" name:"OnlineCount"`

	// 上传固件文件的时间，UNIX时间戳，单位秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 发布记录的最后变更时间，UNIX时间戳，单位秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	UploadTime *uint64 `json:"UploadTime,omitempty" name:"UploadTime"`

	// 该固件版本发布的变更次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifyTimes *uint64 `json:"ModifyTimes,omitempty" name:"ModifyTimes"`

	// 备注信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 版本发布的描述信息，需要国际化，可以为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	Contents *Contents `json:"Contents,omitempty" name:"Contents"`

	// 月活设备数，当月第一天开始有上线的设备数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AliveInMonthCnt *uint64 `json:"AliveInMonthCnt,omitempty" name:"AliveInMonthCnt"`
}