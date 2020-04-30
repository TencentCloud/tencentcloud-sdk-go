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

type CreateAppUsrRequest struct {
	*tchttp.BaseRequest

	// 标识用户的唯一ID，防止同一个用户多次注册
	CunionId *string `json:"CunionId,omitempty" name:"CunionId"`
}

func (r *CreateAppUsrRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAppUsrRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAppUsrResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 厂商云标识用户的唯一ID
		CunionId *string `json:"CunionId,omitempty" name:"CunionId"`

		// 客户的终端用户在IoT Video上的唯一标识ID
		AccessId *string `json:"AccessId,omitempty" name:"AccessId"`

		// 用户是否为新创建
		NewRegist *bool `json:"NewRegist,omitempty" name:"NewRegist"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAppUsrResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAppUsrResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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
}

func (r *CreateBindingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateBindingRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateBindingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 访问设备的AccessToken
		AccessToken *string `json:"AccessToken,omitempty" name:"AccessToken"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateBindingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateBindingResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDevTokenRequest struct {
	*tchttp.BaseRequest

	// 客户的终端用户在IoT Video上的唯一标识ID
	AccessId *string `json:"AccessId,omitempty" name:"AccessId"`

	// 设备TID列表,0<元素数量<=100
	Tids []*string `json:"Tids,omitempty" name:"Tids" list`

	// Token的TTL(time to alive)分钟数
	TtlMinutes *uint64 `json:"TtlMinutes,omitempty" name:"TtlMinutes"`
}

func (r *CreateDevTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDevTokenRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDevTokenResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的用户token列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data []*DevTokenInfo `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDevTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDevTokenResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *CreateDevicesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDevicesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 新创建设备的认证信息
		Data []*DeviceCertificate `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDevicesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateGencodeRequest struct {
	*tchttp.BaseRequest

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 物模型发布版本号，-1代表最新编辑（未发布）的版本
	Revision *int64 `json:"Revision,omitempty" name:"Revision"`
}

func (r *CreateGencodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateGencodeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateGencodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 生成的源代码(zip压缩后的base64编码)
	// 注意：此字段可能返回 null，表示取不到有效值。
		ZipCode *string `json:"ZipCode,omitempty" name:"ZipCode"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateGencodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateGencodeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *CreateIotDataTypeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateIotDataTypeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateIotDataTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateIotDataTypeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *CreateIotModelRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateIotModelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateIotModelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateIotModelResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateProductRequest struct {
	*tchttp.BaseRequest

	// 产器型号(APP产品,为APP包名)
	ProductModel *string `json:"ProductModel,omitempty" name:"ProductModel"`

	// 设备功能码（ypsxth:音频双向通话 ，spdxth:视频单向通话）
	Features []*string `json:"Features,omitempty" name:"Features" list`

	// 产品名称
	// 仅支持中文、英文、数字、下划线，不超过32个字符
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 产品描述信息
	// 不支持单引号、双引号、退格符、回车符、换行符、制表符、反斜杠、下划线、“%”、“#”、“$”，不超过128字符
	ProductDescription *string `json:"ProductDescription,omitempty" name:"ProductDescription"`

	// 主芯片产商ID
	ChipManufactureId *string `json:"ChipManufactureId,omitempty" name:"ChipManufactureId"`

	// 主芯片ID
	ChipId *string `json:"ChipId,omitempty" name:"ChipId"`
}

func (r *CreateProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateProductRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateProductResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 产品详细信息
		Data *ProductBase `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateProductResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *CreateStorageRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateStorageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateStorageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateStorageResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTraceIdsRequest struct {
	*tchttp.BaseRequest

	// 设备TID列表
	Tids []*string `json:"Tids,omitempty" name:"Tids" list`
}

func (r *CreateTraceIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTraceIdsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTraceIdsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTraceIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTraceIdsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *CreateUploadPathRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateUploadPathResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 固件上传地址URL，用户可将本地的固件文件通过该URL以PUT的请求方式上传。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data *string `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateUploadPathResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateUploadPathResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateUsrTokenRequest struct {
	*tchttp.BaseRequest

	// 终端用户在IoT Video上的唯一标识ID
	AccessId *string `json:"AccessId,omitempty" name:"AccessId"`

	// 终端唯一ID，用于区分同一个用户的多个终端
	UniqueId *string `json:"UniqueId,omitempty" name:"UniqueId"`

	// Token的TTL(time to alive)分钟数
	TtlMinutes *uint64 `json:"TtlMinutes,omitempty" name:"TtlMinutes"`
}

func (r *CreateUsrTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateUsrTokenRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateUsrTokenResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
}

func (r *CreateUsrTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateUsrTokenResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DeleteAppUsrRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAppUsrResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAppUsrResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAppUsrResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DeleteBindingRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteBindingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBindingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteBindingResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteDeviceRequest struct {
	*tchttp.BaseRequest

	// 设备TID列表
	Tids []*string `json:"Tids,omitempty" name:"Tids" list`
}

func (r *DeleteDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteDeviceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteDeviceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteDeviceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DeleteIotDataTypeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteIotDataTypeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteIotDataTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteIotDataTypeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DeleteMessageQueueRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMessageQueueResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMessageQueueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMessageQueueResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DeleteOtaVersionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteOtaVersionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteOtaVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteOtaVersionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DeleteProductRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteProductResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteProductResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTraceIdsRequest struct {
	*tchttp.BaseRequest

	// 设备TID列表
	Tids []*string `json:"Tids,omitempty" name:"Tids" list`
}

func (r *DeleteTraceIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTraceIdsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTraceIdsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTraceIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTraceIdsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeBindDevRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBindDevResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 绑定的设备列表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data []*BindDevInfo `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBindDevResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBindDevResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeBindUsrRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBindUsrResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 具有绑定关系的终端用户信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data []*BindUsrInfo `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBindUsrResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBindUsrResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeDeviceModelRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceModelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 设备物模型信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data *DeviceModelData `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeviceModelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDeviceModelResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeDeviceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 设备信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data *DeviceData `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDeviceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeDevicesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDevicesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 设备信息 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data []*DevicesData `json:"Data,omitempty" name:"Data" list`

		// 设备总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDevicesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeIotDataTypeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeIotDataTypeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 自定义数据类型，json格式的字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data []*string `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIotDataTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeIotDataTypeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeIotModelRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeIotModelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 物模型定义，json格式的字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data *string `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIotModelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeIotModelResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeIotModelsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeIotModelsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 历史版本列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data []*IotModelData `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIotModelsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeIotModelsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeLogsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 设备日志信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data []*LogData `json:"Data,omitempty" name:"Data" list`

		// Data数组所包含的信息条数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLogsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeMessageQueueRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMessageQueueResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 消息队列配置
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data *MsgQueueData `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMessageQueueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMessageQueueResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeModelDataRetRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeModelDataRetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 设备响应结果
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data *string `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeModelDataRetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeModelDataRetResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeOtaVersionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeOtaVersionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 版本数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 版本详细信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data []*VersionData `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOtaVersionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeOtaVersionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeProductRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProductResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 产品详情
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data *ProductData `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProductResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeProductsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProductsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 产品详细信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data []*ProductData `json:"Data,omitempty" name:"Data" list`

		// 产品总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProductsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribePubVersionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePubVersionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 历史发布的版本列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data []*OtaPubHistory `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePubVersionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePubVersionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRegistrationStatusRequest struct {
	*tchttp.BaseRequest

	// 终端用户的唯一ID列表，0<元素数量<=100
	CunionIds []*string `json:"CunionIds,omitempty" name:"CunionIds" list`
}

func (r *DescribeRegistrationStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRegistrationStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRegistrationStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 终端用户注册状态列表
		Data []*RegisteredStatus `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRegistrationStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRegistrationStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeRunLogRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRunLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 设备运行日志文本信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data *string `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRunLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRunLogResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTraceIdsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeTraceIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTraceIdsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTraceIdsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 设备TID列表，列表元素之间以“,”分隔
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data *string `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTraceIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTraceIdsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTraceStatusRequest struct {
	*tchttp.BaseRequest

	// 设备TID列表
	Tids []*string `json:"Tids,omitempty" name:"Tids" list`
}

func (r *DescribeTraceStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTraceStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTraceStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 设备追踪状态列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data []*TraceStatus `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTraceStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

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

type DisableDeviceRequest struct {
	*tchttp.BaseRequest

	// 设备TID ≤100
	Tids []*string `json:"Tids,omitempty" name:"Tids" list`
}

func (r *DisableDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableDeviceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableDeviceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableDeviceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableDeviceStreamRequest struct {
	*tchttp.BaseRequest

	// 设备TID列表
	Tids []*string `json:"Tids,omitempty" name:"Tids" list`
}

func (r *DisableDeviceStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableDeviceStreamRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableDeviceStreamResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableDeviceStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableDeviceStreamResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DisableOtaVersionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableOtaVersionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableOtaVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

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

func (r *ModifyDeviceActionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDeviceActionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 设备端的响应结果
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data *string `json:"Data,omitempty" name:"Data"`

		// 任务ID
	// 若设备端未能及时响应时，会返回此字段，用户可以通过DescribeModelDataRet获取设备的最终响应结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDeviceActionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDeviceActionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *ModifyDevicePropertyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDevicePropertyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDevicePropertyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDevicePropertyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *ModifyProductRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyProductResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyProductResponse) FromJsonString(s string) error {
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
	Features []*string `json:"Features,omitempty" name:"Features" list`

	// 产器型号(APP产品,为APP包名)
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductModel *string `json:"ProductModel,omitempty" name:"ProductModel"`

	// 主芯片厂商id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChipManufactureId *string `json:"ChipManufactureId,omitempty" name:"ChipManufactureId"`

	// 主芯片型号
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChipId *string `json:"ChipId,omitempty" name:"ChipId"`
}

type RegisteredStatus struct {

	// 终端用户的唯一ID
	CunionId *string `json:"CunionId,omitempty" name:"CunionId"`

	// 注册状态
	IsRegisted *bool `json:"IsRegisted,omitempty" name:"IsRegisted"`
}

type RunDeviceRequest struct {
	*tchttp.BaseRequest

	// TID列表 ≤100
	Tids []*string `json:"Tids,omitempty" name:"Tids" list`
}

func (r *RunDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RunDeviceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RunDeviceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RunDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RunDeviceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RunDeviceStreamRequest struct {
	*tchttp.BaseRequest

	// 设备TID 列表
	Tids []*string `json:"Tids,omitempty" name:"Tids" list`
}

func (r *RunDeviceStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RunDeviceStreamRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RunDeviceStreamResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RunDeviceStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RunDeviceStreamResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *RunIotModelRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RunIotModelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RunIotModelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RunIotModelResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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
	OldVersions []*string `json:"OldVersions,omitempty" name:"OldVersions" list`

	// 操作人
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *RunOtaVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RunOtaVersionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RunOtaVersionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RunOtaVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RunOtaVersionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RunTestOtaVersionRequest struct {
	*tchttp.BaseRequest

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 固件版本号，格式为x.y.z， x，y 范围0-63，z范围1~524288
	OtaVersion *string `json:"OtaVersion,omitempty" name:"OtaVersion"`

	// 指定可升级的设备TID
	Tids []*string `json:"Tids,omitempty" name:"Tids" list`

	// 操作人
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *RunTestOtaVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RunTestOtaVersionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RunTestOtaVersionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RunTestOtaVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RunTestOtaVersionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *SendOnlineMsgRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SendOnlineMsgResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 若返回此项则表明需要用户用此taskID进行查询请求是否成功(只有waitresp不等于0的情况下才可能会返回该taskID项)
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 设备响应信息
		Data *string `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendOnlineMsgResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SendOnlineMsgResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetMessageQueueRequest struct {
	*tchttp.BaseRequest

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 消息队列类型 1-CMQ; 2-Ckafka
	MsgQueueType *uint64 `json:"MsgQueueType,omitempty" name:"MsgQueueType"`

	// 消息类型,整型值（0-31）之间以“,”分隔
	// 0：在线状态变更
	// 1.固件版本变更
	// 2.设置参数变更
	// 3.控制状态变更
	// 4.状态信息变更
	// 5.事件发布
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

func (r *SetMessageQueueRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetMessageQueueResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetMessageQueueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetMessageQueueResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TraceStatus struct {

	// 设备TID
	Tid *string `json:"Tid,omitempty" name:"Tid"`

	// 设备追踪状态
	IsExist *bool `json:"IsExist,omitempty" name:"IsExist"`
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

func (r *UpgradeDeviceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeDeviceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 设备端返回的数据
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data *string `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpgradeDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeDeviceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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
}

func (r *UploadOtaVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UploadOtaVersionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UploadOtaVersionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UploadOtaVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

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
}
