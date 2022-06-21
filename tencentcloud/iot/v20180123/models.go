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

package v20180123

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Action struct {
	// 转发至topic
	// 注意：此字段可能返回 null，表示取不到有效值。
	Topic *TopicAction `json:"Topic,omitempty" name:"Topic"`

	// 转发至第三发
	// 注意：此字段可能返回 null，表示取不到有效值。
	Service *ServiceAction `json:"Service,omitempty" name:"Service"`

	// 转发至第三发Ckafka
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ckafka *CkafkaAction `json:"Ckafka,omitempty" name:"Ckafka"`
}

// Predefined struct for user
type ActivateRuleRequestParams struct {
	// 规则Id
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

type ActivateRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则Id
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *ActivateRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ActivateRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ActivateRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ActivateRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ActivateRuleResponse struct {
	*tchttp.BaseResponse
	Response *ActivateRuleResponseParams `json:"Response"`
}

func (r *ActivateRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ActivateRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddDeviceRequestParams struct {
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称，唯一标识某产品下的一个设备
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

type AddDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称，唯一标识某产品下的一个设备
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

func (r *AddDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddDeviceResponseParams struct {
	// 设备信息
	Device *Device `json:"Device,omitempty" name:"Device"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddDeviceResponse struct {
	*tchttp.BaseResponse
	Response *AddDeviceResponseParams `json:"Response"`
}

func (r *AddDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddProductRequestParams struct {
	// 产品名称，同一区域产品名称需唯一，支持中文、英文字母、中划线和下划线，长度不超过31个字符，中文占两个字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 产品描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 数据模版
	DataTemplate []*DataTemplate `json:"DataTemplate,omitempty" name:"DataTemplate"`

	// 产品版本（native表示基础版，template表示高级版，默认值为template）
	DataProtocol *string `json:"DataProtocol,omitempty" name:"DataProtocol"`

	// 设备认证方式（1：动态令牌，2：签名直连鉴权）
	AuthType *uint64 `json:"AuthType,omitempty" name:"AuthType"`

	// 通信方式（other/wifi/cellular/nb-iot）
	CommProtocol *string `json:"CommProtocol,omitempty" name:"CommProtocol"`

	// 产品的设备类型（device: 直连设备；sub_device：子设备；gateway：网关设备）
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`
}

type AddProductRequest struct {
	*tchttp.BaseRequest
	
	// 产品名称，同一区域产品名称需唯一，支持中文、英文字母、中划线和下划线，长度不超过31个字符，中文占两个字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 产品描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 数据模版
	DataTemplate []*DataTemplate `json:"DataTemplate,omitempty" name:"DataTemplate"`

	// 产品版本（native表示基础版，template表示高级版，默认值为template）
	DataProtocol *string `json:"DataProtocol,omitempty" name:"DataProtocol"`

	// 设备认证方式（1：动态令牌，2：签名直连鉴权）
	AuthType *uint64 `json:"AuthType,omitempty" name:"AuthType"`

	// 通信方式（other/wifi/cellular/nb-iot）
	CommProtocol *string `json:"CommProtocol,omitempty" name:"CommProtocol"`

	// 产品的设备类型（device: 直连设备；sub_device：子设备；gateway：网关设备）
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`
}

func (r *AddProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddProductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "DataTemplate")
	delete(f, "DataProtocol")
	delete(f, "AuthType")
	delete(f, "CommProtocol")
	delete(f, "DeviceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddProductResponseParams struct {
	// 产品信息
	Product *Product `json:"Product,omitempty" name:"Product"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddProductResponse struct {
	*tchttp.BaseResponse
	Response *AddProductResponseParams `json:"Response"`
}

func (r *AddProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddRuleRequestParams struct {
	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 查询
	Query *RuleQuery `json:"Query,omitempty" name:"Query"`

	// 转发动作列表
	Actions []*Action `json:"Actions,omitempty" name:"Actions"`

	// 数据类型（0：文本，1：二进制）
	DataType *uint64 `json:"DataType,omitempty" name:"DataType"`
}

type AddRuleRequest struct {
	*tchttp.BaseRequest
	
	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 查询
	Query *RuleQuery `json:"Query,omitempty" name:"Query"`

	// 转发动作列表
	Actions []*Action `json:"Actions,omitempty" name:"Actions"`

	// 数据类型（0：文本，1：二进制）
	DataType *uint64 `json:"DataType,omitempty" name:"DataType"`
}

func (r *AddRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "Query")
	delete(f, "Actions")
	delete(f, "DataType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddRuleResponseParams struct {
	// 规则
	Rule *Rule `json:"Rule,omitempty" name:"Rule"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddRuleResponse struct {
	*tchttp.BaseResponse
	Response *AddRuleResponseParams `json:"Response"`
}

func (r *AddRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddTopicRequestParams struct {
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// Topic名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

type AddTopicRequest struct {
	*tchttp.BaseRequest
	
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// Topic名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

func (r *AddTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "TopicName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddTopicResponseParams struct {
	// Topic信息
	Topic *Topic `json:"Topic,omitempty" name:"Topic"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddTopicResponse struct {
	*tchttp.BaseResponse
	Response *AddTopicResponseParams `json:"Response"`
}

func (r *AddTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AppAddUserRequestParams struct {
	// 用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 密码
	Password *string `json:"Password,omitempty" name:"Password"`
}

type AppAddUserRequest struct {
	*tchttp.BaseRequest
	
	// 用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 密码
	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *AppAddUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AppAddUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserName")
	delete(f, "Password")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AppAddUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AppAddUserResponseParams struct {
	// 应用用户
	AppUser *AppUser `json:"AppUser,omitempty" name:"AppUser"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AppAddUserResponse struct {
	*tchttp.BaseResponse
	Response *AppAddUserResponseParams `json:"Response"`
}

func (r *AppAddUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AppAddUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AppDeleteDeviceRequestParams struct {
	// 访问Token
	AccessToken *string `json:"AccessToken,omitempty" name:"AccessToken"`

	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

type AppDeleteDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 访问Token
	AccessToken *string `json:"AccessToken,omitempty" name:"AccessToken"`

	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

func (r *AppDeleteDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AppDeleteDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessToken")
	delete(f, "ProductId")
	delete(f, "DeviceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AppDeleteDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AppDeleteDeviceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AppDeleteDeviceResponse struct {
	*tchttp.BaseResponse
	Response *AppDeleteDeviceResponseParams `json:"Response"`
}

func (r *AppDeleteDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AppDeleteDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AppDevice struct {
	// 设备Id
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 所属产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 别名
	AliasName *string `json:"AliasName,omitempty" name:"AliasName"`

	// 地区
	Region *string `json:"Region,omitempty" name:"Region"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type AppDeviceDetail struct {
	// 设备Id
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 所属产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 别名
	AliasName *string `json:"AliasName,omitempty" name:"AliasName"`

	// 地区
	Region *string `json:"Region,omitempty" name:"Region"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 设备信息（json）
	DeviceInfo *string `json:"DeviceInfo,omitempty" name:"DeviceInfo"`

	// 数据模板
	DataTemplate []*DataTemplate `json:"DataTemplate,omitempty" name:"DataTemplate"`
}

// Predefined struct for user
type AppGetDeviceDataRequestParams struct {
	// 访问Token
	AccessToken *string `json:"AccessToken,omitempty" name:"AccessToken"`

	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

type AppGetDeviceDataRequest struct {
	*tchttp.BaseRequest
	
	// 访问Token
	AccessToken *string `json:"AccessToken,omitempty" name:"AccessToken"`

	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

func (r *AppGetDeviceDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AppGetDeviceDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessToken")
	delete(f, "ProductId")
	delete(f, "DeviceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AppGetDeviceDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AppGetDeviceDataResponseParams struct {
	// 设备数据。
	DeviceData *string `json:"DeviceData,omitempty" name:"DeviceData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AppGetDeviceDataResponse struct {
	*tchttp.BaseResponse
	Response *AppGetDeviceDataResponseParams `json:"Response"`
}

func (r *AppGetDeviceDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AppGetDeviceDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AppGetDeviceRequestParams struct {
	// 访问Token
	AccessToken *string `json:"AccessToken,omitempty" name:"AccessToken"`

	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

type AppGetDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 访问Token
	AccessToken *string `json:"AccessToken,omitempty" name:"AccessToken"`

	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

func (r *AppGetDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AppGetDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessToken")
	delete(f, "ProductId")
	delete(f, "DeviceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AppGetDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AppGetDeviceResponseParams struct {
	// 绑定设备详情
	AppDevice *AppDeviceDetail `json:"AppDevice,omitempty" name:"AppDevice"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AppGetDeviceResponse struct {
	*tchttp.BaseResponse
	Response *AppGetDeviceResponseParams `json:"Response"`
}

func (r *AppGetDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AppGetDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AppGetDeviceStatusesRequestParams struct {
	// 访问Token
	AccessToken *string `json:"AccessToken,omitempty" name:"AccessToken"`

	// 设备Id列表（单次限制1000个设备）
	DeviceIds []*string `json:"DeviceIds,omitempty" name:"DeviceIds"`
}

type AppGetDeviceStatusesRequest struct {
	*tchttp.BaseRequest
	
	// 访问Token
	AccessToken *string `json:"AccessToken,omitempty" name:"AccessToken"`

	// 设备Id列表（单次限制1000个设备）
	DeviceIds []*string `json:"DeviceIds,omitempty" name:"DeviceIds"`
}

func (r *AppGetDeviceStatusesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AppGetDeviceStatusesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessToken")
	delete(f, "DeviceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AppGetDeviceStatusesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AppGetDeviceStatusesResponseParams struct {
	// 设备状态
	DeviceStatuses []*DeviceStatus `json:"DeviceStatuses,omitempty" name:"DeviceStatuses"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AppGetDeviceStatusesResponse struct {
	*tchttp.BaseResponse
	Response *AppGetDeviceStatusesResponseParams `json:"Response"`
}

func (r *AppGetDeviceStatusesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AppGetDeviceStatusesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AppGetDevicesRequestParams struct {
	// 访问Token
	AccessToken *string `json:"AccessToken,omitempty" name:"AccessToken"`
}

type AppGetDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 访问Token
	AccessToken *string `json:"AccessToken,omitempty" name:"AccessToken"`
}

func (r *AppGetDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AppGetDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AppGetDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AppGetDevicesResponseParams struct {
	// 绑定设备列表
	Devices []*AppDevice `json:"Devices,omitempty" name:"Devices"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AppGetDevicesResponse struct {
	*tchttp.BaseResponse
	Response *AppGetDevicesResponseParams `json:"Response"`
}

func (r *AppGetDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AppGetDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AppGetTokenRequestParams struct {
	// 用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 密码
	Password *string `json:"Password,omitempty" name:"Password"`

	// TTL
	Expire *uint64 `json:"Expire,omitempty" name:"Expire"`
}

type AppGetTokenRequest struct {
	*tchttp.BaseRequest
	
	// 用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 密码
	Password *string `json:"Password,omitempty" name:"Password"`

	// TTL
	Expire *uint64 `json:"Expire,omitempty" name:"Expire"`
}

func (r *AppGetTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AppGetTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserName")
	delete(f, "Password")
	delete(f, "Expire")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AppGetTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AppGetTokenResponseParams struct {
	// 访问Token
	AccessToken *string `json:"AccessToken,omitempty" name:"AccessToken"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AppGetTokenResponse struct {
	*tchttp.BaseResponse
	Response *AppGetTokenResponseParams `json:"Response"`
}

func (r *AppGetTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AppGetTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AppGetUserRequestParams struct {
	// 访问Token
	AccessToken *string `json:"AccessToken,omitempty" name:"AccessToken"`
}

type AppGetUserRequest struct {
	*tchttp.BaseRequest
	
	// 访问Token
	AccessToken *string `json:"AccessToken,omitempty" name:"AccessToken"`
}

func (r *AppGetUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AppGetUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AppGetUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AppGetUserResponseParams struct {
	// 用户信息
	AppUser *AppUser `json:"AppUser,omitempty" name:"AppUser"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AppGetUserResponse struct {
	*tchttp.BaseResponse
	Response *AppGetUserResponseParams `json:"Response"`
}

func (r *AppGetUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AppGetUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AppIssueDeviceControlRequestParams struct {
	// 访问Token
	AccessToken *string `json:"AccessToken,omitempty" name:"AccessToken"`

	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 控制数据（json）
	ControlData *string `json:"ControlData,omitempty" name:"ControlData"`

	// 是否发送metadata字段
	Metadata *bool `json:"Metadata,omitempty" name:"Metadata"`
}

type AppIssueDeviceControlRequest struct {
	*tchttp.BaseRequest
	
	// 访问Token
	AccessToken *string `json:"AccessToken,omitempty" name:"AccessToken"`

	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 控制数据（json）
	ControlData *string `json:"ControlData,omitempty" name:"ControlData"`

	// 是否发送metadata字段
	Metadata *bool `json:"Metadata,omitempty" name:"Metadata"`
}

func (r *AppIssueDeviceControlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AppIssueDeviceControlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessToken")
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "ControlData")
	delete(f, "Metadata")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AppIssueDeviceControlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AppIssueDeviceControlResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AppIssueDeviceControlResponse struct {
	*tchttp.BaseResponse
	Response *AppIssueDeviceControlResponseParams `json:"Response"`
}

func (r *AppIssueDeviceControlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AppIssueDeviceControlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AppResetPasswordRequestParams struct {
	// 访问Token
	AccessToken *string `json:"AccessToken,omitempty" name:"AccessToken"`

	// 旧密码
	OldPassword *string `json:"OldPassword,omitempty" name:"OldPassword"`

	// 新密码
	NewPassword *string `json:"NewPassword,omitempty" name:"NewPassword"`
}

type AppResetPasswordRequest struct {
	*tchttp.BaseRequest
	
	// 访问Token
	AccessToken *string `json:"AccessToken,omitempty" name:"AccessToken"`

	// 旧密码
	OldPassword *string `json:"OldPassword,omitempty" name:"OldPassword"`

	// 新密码
	NewPassword *string `json:"NewPassword,omitempty" name:"NewPassword"`
}

func (r *AppResetPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AppResetPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessToken")
	delete(f, "OldPassword")
	delete(f, "NewPassword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AppResetPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AppResetPasswordResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AppResetPasswordResponse struct {
	*tchttp.BaseResponse
	Response *AppResetPasswordResponseParams `json:"Response"`
}

func (r *AppResetPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AppResetPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AppSecureAddDeviceRequestParams struct {
	// 访问Token
	AccessToken *string `json:"AccessToken,omitempty" name:"AccessToken"`

	// 设备签名
	DeviceSignature *string `json:"DeviceSignature,omitempty" name:"DeviceSignature"`
}

type AppSecureAddDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 访问Token
	AccessToken *string `json:"AccessToken,omitempty" name:"AccessToken"`

	// 设备签名
	DeviceSignature *string `json:"DeviceSignature,omitempty" name:"DeviceSignature"`
}

func (r *AppSecureAddDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AppSecureAddDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessToken")
	delete(f, "DeviceSignature")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AppSecureAddDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AppSecureAddDeviceResponseParams struct {
	// 绑定设备信息
	AppDevice *AppDevice `json:"AppDevice,omitempty" name:"AppDevice"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AppSecureAddDeviceResponse struct {
	*tchttp.BaseResponse
	Response *AppSecureAddDeviceResponseParams `json:"Response"`
}

func (r *AppSecureAddDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AppSecureAddDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AppUpdateDeviceRequestParams struct {
	// 访问Token
	AccessToken *string `json:"AccessToken,omitempty" name:"AccessToken"`

	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 设备别名
	AliasName *string `json:"AliasName,omitempty" name:"AliasName"`
}

type AppUpdateDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 访问Token
	AccessToken *string `json:"AccessToken,omitempty" name:"AccessToken"`

	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 设备别名
	AliasName *string `json:"AliasName,omitempty" name:"AliasName"`
}

func (r *AppUpdateDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AppUpdateDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessToken")
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "AliasName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AppUpdateDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AppUpdateDeviceResponseParams struct {
	// 设备信息
	AppDevice *AppDevice `json:"AppDevice,omitempty" name:"AppDevice"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AppUpdateDeviceResponse struct {
	*tchttp.BaseResponse
	Response *AppUpdateDeviceResponseParams `json:"Response"`
}

func (r *AppUpdateDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AppUpdateDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AppUpdateUserRequestParams struct {
	// 访问Token
	AccessToken *string `json:"AccessToken,omitempty" name:"AccessToken"`

	// 昵称
	NickName *string `json:"NickName,omitempty" name:"NickName"`
}

type AppUpdateUserRequest struct {
	*tchttp.BaseRequest
	
	// 访问Token
	AccessToken *string `json:"AccessToken,omitempty" name:"AccessToken"`

	// 昵称
	NickName *string `json:"NickName,omitempty" name:"NickName"`
}

func (r *AppUpdateUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AppUpdateUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessToken")
	delete(f, "NickName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AppUpdateUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AppUpdateUserResponseParams struct {
	// 应用用户
	AppUser *AppUser `json:"AppUser,omitempty" name:"AppUser"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AppUpdateUserResponse struct {
	*tchttp.BaseResponse
	Response *AppUpdateUserResponseParams `json:"Response"`
}

func (r *AppUpdateUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AppUpdateUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AppUser struct {
	// 应用Id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 昵称
	NickName *string `json:"NickName,omitempty" name:"NickName"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 修改时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

// Predefined struct for user
type AssociateSubDeviceToGatewayProductRequestParams struct {
	// 子设备产品Id
	SubDeviceProductId *string `json:"SubDeviceProductId,omitempty" name:"SubDeviceProductId"`

	// 网关产品Id
	GatewayProductId *string `json:"GatewayProductId,omitempty" name:"GatewayProductId"`
}

type AssociateSubDeviceToGatewayProductRequest struct {
	*tchttp.BaseRequest
	
	// 子设备产品Id
	SubDeviceProductId *string `json:"SubDeviceProductId,omitempty" name:"SubDeviceProductId"`

	// 网关产品Id
	GatewayProductId *string `json:"GatewayProductId,omitempty" name:"GatewayProductId"`
}

func (r *AssociateSubDeviceToGatewayProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateSubDeviceToGatewayProductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubDeviceProductId")
	delete(f, "GatewayProductId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateSubDeviceToGatewayProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateSubDeviceToGatewayProductResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AssociateSubDeviceToGatewayProductResponse struct {
	*tchttp.BaseResponse
	Response *AssociateSubDeviceToGatewayProductResponseParams `json:"Response"`
}

func (r *AssociateSubDeviceToGatewayProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateSubDeviceToGatewayProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BoolData struct {
	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 描述
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// 读写模式
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// 取值列表
	Range []*bool `json:"Range,omitempty" name:"Range"`
}

type CkafkaAction struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// topic名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`
}

type DataHistoryEntry struct {
	// 日志id
	Id *string `json:"Id,omitempty" name:"Id"`

	// 时间戳
	Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 数据
	Data *string `json:"Data,omitempty" name:"Data"`
}

type DataTemplate struct {
	// 数字类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Number *NumberData `json:"Number,omitempty" name:"Number"`

	// 字符串类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	String *StringData `json:"String,omitempty" name:"String"`

	// 枚举类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enum *EnumData `json:"Enum,omitempty" name:"Enum"`

	// 布尔类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bool *BoolData `json:"Bool,omitempty" name:"Bool"`
}

// Predefined struct for user
type DeactivateRuleRequestParams struct {
	// 规则Id
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

type DeactivateRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则Id
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *DeactivateRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeactivateRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeactivateRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeactivateRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeactivateRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeactivateRuleResponseParams `json:"Response"`
}

func (r *DeactivateRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeactivateRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DebugLogEntry struct {
	// 日志id
	Id *string `json:"Id,omitempty" name:"Id"`

	// 行为（事件）
	Event *string `json:"Event,omitempty" name:"Event"`

	// shadow/action/mqtt, 分别表示：影子/规则引擎/上下线日志
	LogType *string `json:"LogType,omitempty" name:"LogType"`

	// 时间戳
	Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// success/fail
	Result *string `json:"Result,omitempty" name:"Result"`

	// 日志详细内容
	Data *string `json:"Data,omitempty" name:"Data"`

	// 数据来源topic
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

// Predefined struct for user
type DeleteDeviceRequestParams struct {
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

type DeleteDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
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
	delete(f, "ProductId")
	delete(f, "DeviceName")
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
type DeleteProductRequestParams struct {
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

type DeleteProductRequest struct {
	*tchttp.BaseRequest
	
	// 产品Id
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
type DeleteRuleRequestParams struct {
	// 规则Id
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

type DeleteRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则Id
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *DeleteRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRuleResponseParams `json:"Response"`
}

func (r *DeleteRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTopicRequestParams struct {
	// TopicId
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

type DeleteTopicRequest struct {
	*tchttp.BaseRequest
	
	// TopicId
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

func (r *DeleteTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "ProductId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTopicResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteTopicResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTopicResponseParams `json:"Response"`
}

func (r *DeleteTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Device struct {
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 设备密钥
	DeviceSecret *string `json:"DeviceSecret,omitempty" name:"DeviceSecret"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 设备信息（json）
	DeviceInfo *string `json:"DeviceInfo,omitempty" name:"DeviceInfo"`
}

type DeviceEntry struct {
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 设备密钥
	DeviceSecret *string `json:"DeviceSecret,omitempty" name:"DeviceSecret"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type DeviceLogEntry struct {
	// 日志id
	Id *string `json:"Id,omitempty" name:"Id"`

	// 日志内容
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 状态码
	Code *string `json:"Code,omitempty" name:"Code"`

	// 时间戳
	Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 设备动作
	Method *string `json:"Method,omitempty" name:"Method"`
}

type DeviceSignature struct {
	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 设备签名
	DeviceSignature *string `json:"DeviceSignature,omitempty" name:"DeviceSignature"`
}

type DeviceStatData struct {
	// 时间点
	Datetime *string `json:"Datetime,omitempty" name:"Datetime"`

	// 在线设备数
	DeviceOnline *uint64 `json:"DeviceOnline,omitempty" name:"DeviceOnline"`

	// 激活设备数
	DeviceActive *uint64 `json:"DeviceActive,omitempty" name:"DeviceActive"`

	// 设备总数
	DeviceTotal *uint64 `json:"DeviceTotal,omitempty" name:"DeviceTotal"`
}

type DeviceStatus struct {
	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 设备状态（inactive, online, offline）
	Status *string `json:"Status,omitempty" name:"Status"`

	// 首次上线时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstOnline *string `json:"FirstOnline,omitempty" name:"FirstOnline"`

	// 最后上线时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastOnline *string `json:"LastOnline,omitempty" name:"LastOnline"`

	// 上线次数
	OnlineTimes *uint64 `json:"OnlineTimes,omitempty" name:"OnlineTimes"`
}

type EnumData struct {
	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 描述
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// 读写模式
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// 取值列表
	Range []*string `json:"Range,omitempty" name:"Range"`
}

// Predefined struct for user
type GetDataHistoryRequestParams struct {
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称列表，允许最多一次100台
	DeviceNames []*string `json:"DeviceNames,omitempty" name:"DeviceNames"`

	// 查询开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询数据量
	Size *uint64 `json:"Size,omitempty" name:"Size"`

	// 时间排序（desc/asc）
	Order *string `json:"Order,omitempty" name:"Order"`

	// 查询游标
	ScrollId *string `json:"ScrollId,omitempty" name:"ScrollId"`
}

type GetDataHistoryRequest struct {
	*tchttp.BaseRequest
	
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称列表，允许最多一次100台
	DeviceNames []*string `json:"DeviceNames,omitempty" name:"DeviceNames"`

	// 查询开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询数据量
	Size *uint64 `json:"Size,omitempty" name:"Size"`

	// 时间排序（desc/asc）
	Order *string `json:"Order,omitempty" name:"Order"`

	// 查询游标
	ScrollId *string `json:"ScrollId,omitempty" name:"ScrollId"`
}

func (r *GetDataHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDataHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceNames")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Size")
	delete(f, "Order")
	delete(f, "ScrollId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDataHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDataHistoryResponseParams struct {
	// 数据历史
	DataHistory []*DataHistoryEntry `json:"DataHistory,omitempty" name:"DataHistory"`

	// 查询游标
	ScrollId *string `json:"ScrollId,omitempty" name:"ScrollId"`

	// 查询游标超时
	ScrollTimeout *uint64 `json:"ScrollTimeout,omitempty" name:"ScrollTimeout"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetDataHistoryResponse struct {
	*tchttp.BaseResponse
	Response *GetDataHistoryResponseParams `json:"Response"`
}

func (r *GetDataHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDataHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDebugLogRequestParams struct {
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称列表，最大支持100台
	DeviceNames []*string `json:"DeviceNames,omitempty" name:"DeviceNames"`

	// 查询开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询数据量
	Size *uint64 `json:"Size,omitempty" name:"Size"`

	// 时间排序（desc/asc）
	Order *string `json:"Order,omitempty" name:"Order"`

	// 查询游标
	ScrollId *string `json:"ScrollId,omitempty" name:"ScrollId"`

	// 日志类型（shadow/action/mqtt）
	Type *string `json:"Type,omitempty" name:"Type"`
}

type GetDebugLogRequest struct {
	*tchttp.BaseRequest
	
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称列表，最大支持100台
	DeviceNames []*string `json:"DeviceNames,omitempty" name:"DeviceNames"`

	// 查询开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询数据量
	Size *uint64 `json:"Size,omitempty" name:"Size"`

	// 时间排序（desc/asc）
	Order *string `json:"Order,omitempty" name:"Order"`

	// 查询游标
	ScrollId *string `json:"ScrollId,omitempty" name:"ScrollId"`

	// 日志类型（shadow/action/mqtt）
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *GetDebugLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDebugLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceNames")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Size")
	delete(f, "Order")
	delete(f, "ScrollId")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDebugLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDebugLogResponseParams struct {
	// 调试日志
	DebugLog []*DebugLogEntry `json:"DebugLog,omitempty" name:"DebugLog"`

	// 查询游标
	ScrollId *string `json:"ScrollId,omitempty" name:"ScrollId"`

	// 游标超时
	ScrollTimeout *uint64 `json:"ScrollTimeout,omitempty" name:"ScrollTimeout"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetDebugLogResponse struct {
	*tchttp.BaseResponse
	Response *GetDebugLogResponseParams `json:"Response"`
}

func (r *GetDebugLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDebugLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDeviceDataRequestParams struct {
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

type GetDeviceDataRequest struct {
	*tchttp.BaseRequest
	
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

func (r *GetDeviceDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDeviceDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDeviceDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDeviceDataResponseParams struct {
	// 设备数据
	DeviceData *string `json:"DeviceData,omitempty" name:"DeviceData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetDeviceDataResponse struct {
	*tchttp.BaseResponse
	Response *GetDeviceDataResponseParams `json:"Response"`
}

func (r *GetDeviceDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDeviceDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDeviceLogRequestParams struct {
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称列表，最大支持100台
	DeviceNames []*string `json:"DeviceNames,omitempty" name:"DeviceNames"`

	// 查询开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询数据量
	Size *uint64 `json:"Size,omitempty" name:"Size"`

	// 时间排序（desc/asc）
	Order *string `json:"Order,omitempty" name:"Order"`

	// 查询游标
	ScrollId *string `json:"ScrollId,omitempty" name:"ScrollId"`

	// 日志类型（comm/status）
	Type *string `json:"Type,omitempty" name:"Type"`
}

type GetDeviceLogRequest struct {
	*tchttp.BaseRequest
	
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称列表，最大支持100台
	DeviceNames []*string `json:"DeviceNames,omitempty" name:"DeviceNames"`

	// 查询开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询数据量
	Size *uint64 `json:"Size,omitempty" name:"Size"`

	// 时间排序（desc/asc）
	Order *string `json:"Order,omitempty" name:"Order"`

	// 查询游标
	ScrollId *string `json:"ScrollId,omitempty" name:"ScrollId"`

	// 日志类型（comm/status）
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *GetDeviceLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDeviceLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceNames")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Size")
	delete(f, "Order")
	delete(f, "ScrollId")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDeviceLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDeviceLogResponseParams struct {
	// 设备日志
	DeviceLog []*DeviceLogEntry `json:"DeviceLog,omitempty" name:"DeviceLog"`

	// 查询游标
	ScrollId *string `json:"ScrollId,omitempty" name:"ScrollId"`

	// 游标超时
	ScrollTimeout *uint64 `json:"ScrollTimeout,omitempty" name:"ScrollTimeout"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetDeviceLogResponse struct {
	*tchttp.BaseResponse
	Response *GetDeviceLogResponseParams `json:"Response"`
}

func (r *GetDeviceLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDeviceLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDeviceRequestParams struct {
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

type GetDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

func (r *GetDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDeviceResponseParams struct {
	// 设备信息
	Device *Device `json:"Device,omitempty" name:"Device"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetDeviceResponse struct {
	*tchttp.BaseResponse
	Response *GetDeviceResponseParams `json:"Response"`
}

func (r *GetDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDeviceSignaturesRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称列表（单次限制1000个设备）
	DeviceNames []*string `json:"DeviceNames,omitempty" name:"DeviceNames"`

	// 过期时间
	Expire *uint64 `json:"Expire,omitempty" name:"Expire"`
}

type GetDeviceSignaturesRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称列表（单次限制1000个设备）
	DeviceNames []*string `json:"DeviceNames,omitempty" name:"DeviceNames"`

	// 过期时间
	Expire *uint64 `json:"Expire,omitempty" name:"Expire"`
}

func (r *GetDeviceSignaturesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDeviceSignaturesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceNames")
	delete(f, "Expire")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDeviceSignaturesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDeviceSignaturesResponseParams struct {
	// 设备绑定签名列表
	DeviceSignatures []*DeviceSignature `json:"DeviceSignatures,omitempty" name:"DeviceSignatures"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetDeviceSignaturesResponse struct {
	*tchttp.BaseResponse
	Response *GetDeviceSignaturesResponseParams `json:"Response"`
}

func (r *GetDeviceSignaturesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDeviceSignaturesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDeviceStatisticsRequestParams struct {
	// 产品Id列表
	Products []*string `json:"Products,omitempty" name:"Products"`

	// 开始日期
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 结束日期
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

type GetDeviceStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 产品Id列表
	Products []*string `json:"Products,omitempty" name:"Products"`

	// 开始日期
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 结束日期
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

func (r *GetDeviceStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDeviceStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Products")
	delete(f, "StartDate")
	delete(f, "EndDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDeviceStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDeviceStatisticsResponseParams struct {
	// 统计数据
	DeviceStatistics []*DeviceStatData `json:"DeviceStatistics,omitempty" name:"DeviceStatistics"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetDeviceStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *GetDeviceStatisticsResponseParams `json:"Response"`
}

func (r *GetDeviceStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDeviceStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDeviceStatusesRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称列表（单次限制1000个设备）
	DeviceNames []*string `json:"DeviceNames,omitempty" name:"DeviceNames"`
}

type GetDeviceStatusesRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称列表（单次限制1000个设备）
	DeviceNames []*string `json:"DeviceNames,omitempty" name:"DeviceNames"`
}

func (r *GetDeviceStatusesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDeviceStatusesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDeviceStatusesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDeviceStatusesResponseParams struct {
	// 设备状态列表
	DeviceStatuses []*DeviceStatus `json:"DeviceStatuses,omitempty" name:"DeviceStatuses"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetDeviceStatusesResponse struct {
	*tchttp.BaseResponse
	Response *GetDeviceStatusesResponseParams `json:"Response"`
}

func (r *GetDeviceStatusesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDeviceStatusesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDevicesRequestParams struct {
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 长度
	Length *uint64 `json:"Length,omitempty" name:"Length"`

	// 关键字查询
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

type GetDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 长度
	Length *uint64 `json:"Length,omitempty" name:"Length"`

	// 关键字查询
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

func (r *GetDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "Offset")
	delete(f, "Length")
	delete(f, "Keyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDevicesResponseParams struct {
	// 设备列表
	Devices []*DeviceEntry `json:"Devices,omitempty" name:"Devices"`

	// 设备总数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetDevicesResponse struct {
	*tchttp.BaseResponse
	Response *GetDevicesResponseParams `json:"Response"`
}

func (r *GetDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetProductRequestParams struct {
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

type GetProductRequest struct {
	*tchttp.BaseRequest
	
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

func (r *GetProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetProductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetProductResponseParams struct {
	// 产品信息
	Product *Product `json:"Product,omitempty" name:"Product"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetProductResponse struct {
	*tchttp.BaseResponse
	Response *GetProductResponseParams `json:"Response"`
}

func (r *GetProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetProductsRequestParams struct {
	// 偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 长度
	Length *uint64 `json:"Length,omitempty" name:"Length"`
}

type GetProductsRequest struct {
	*tchttp.BaseRequest
	
	// 偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 长度
	Length *uint64 `json:"Length,omitempty" name:"Length"`
}

func (r *GetProductsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetProductsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Length")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetProductsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetProductsResponseParams struct {
	// Product列表
	Products []*ProductEntry `json:"Products,omitempty" name:"Products"`

	// Product总数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetProductsResponse struct {
	*tchttp.BaseResponse
	Response *GetProductsResponseParams `json:"Response"`
}

func (r *GetProductsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetProductsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRuleRequestParams struct {
	// 规则Id
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

type GetRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则Id
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *GetRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRuleResponseParams struct {
	// 规则
	Rule *Rule `json:"Rule,omitempty" name:"Rule"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetRuleResponse struct {
	*tchttp.BaseResponse
	Response *GetRuleResponseParams `json:"Response"`
}

func (r *GetRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRulesRequestParams struct {
	// 偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 长度
	Length *uint64 `json:"Length,omitempty" name:"Length"`
}

type GetRulesRequest struct {
	*tchttp.BaseRequest
	
	// 偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 长度
	Length *uint64 `json:"Length,omitempty" name:"Length"`
}

func (r *GetRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Length")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRulesResponseParams struct {
	// 规则列表
	Rules []*Rule `json:"Rules,omitempty" name:"Rules"`

	// 规则总数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetRulesResponse struct {
	*tchttp.BaseResponse
	Response *GetRulesResponseParams `json:"Response"`
}

func (r *GetRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTopicRequestParams struct {
	// TopicId
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

type GetTopicRequest struct {
	*tchttp.BaseRequest
	
	// TopicId
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

func (r *GetTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "ProductId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTopicResponseParams struct {
	// Topic信息
	Topic *Topic `json:"Topic,omitempty" name:"Topic"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetTopicResponse struct {
	*tchttp.BaseResponse
	Response *GetTopicResponseParams `json:"Response"`
}

func (r *GetTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTopicsRequestParams struct {
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 长度
	Length *uint64 `json:"Length,omitempty" name:"Length"`
}

type GetTopicsRequest struct {
	*tchttp.BaseRequest
	
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 长度
	Length *uint64 `json:"Length,omitempty" name:"Length"`
}

func (r *GetTopicsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTopicsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "Offset")
	delete(f, "Length")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTopicsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTopicsResponseParams struct {
	// Topic列表
	Topics []*Topic `json:"Topics,omitempty" name:"Topics"`

	// Topic总数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetTopicsResponse struct {
	*tchttp.BaseResponse
	Response *GetTopicsResponseParams `json:"Response"`
}

func (r *GetTopicsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTopicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IssueDeviceControlRequestParams struct {
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 控制数据（json）
	ControlData *string `json:"ControlData,omitempty" name:"ControlData"`

	// 是否发送metadata字段
	Metadata *bool `json:"Metadata,omitempty" name:"Metadata"`
}

type IssueDeviceControlRequest struct {
	*tchttp.BaseRequest
	
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 控制数据（json）
	ControlData *string `json:"ControlData,omitempty" name:"ControlData"`

	// 是否发送metadata字段
	Metadata *bool `json:"Metadata,omitempty" name:"Metadata"`
}

func (r *IssueDeviceControlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IssueDeviceControlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "ControlData")
	delete(f, "Metadata")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IssueDeviceControlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IssueDeviceControlResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type IssueDeviceControlResponse struct {
	*tchttp.BaseResponse
	Response *IssueDeviceControlResponseParams `json:"Response"`
}

func (r *IssueDeviceControlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IssueDeviceControlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NumberData struct {
	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 描述
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// 读写模式
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// 取值范围
	Range []*float64 `json:"Range,omitempty" name:"Range"`
}

type Product struct {
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 产品Key
	ProductKey *string `json:"ProductKey,omitempty" name:"ProductKey"`

	// AppId
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// 产品名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 产品描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 连接域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 产品规格
	Standard *uint64 `json:"Standard,omitempty" name:"Standard"`

	// 鉴权类型（0：直连，1：Token）
	AuthType *uint64 `json:"AuthType,omitempty" name:"AuthType"`

	// 删除（0未删除）
	Deleted *uint64 `json:"Deleted,omitempty" name:"Deleted"`

	// 备注
	Message *string `json:"Message,omitempty" name:"Message"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 数据模版
	DataTemplate []*DataTemplate `json:"DataTemplate,omitempty" name:"DataTemplate"`

	// 数据协议（native/template）
	DataProtocol *string `json:"DataProtocol,omitempty" name:"DataProtocol"`

	// 直连用户名
	Username *string `json:"Username,omitempty" name:"Username"`

	// 直连密码
	Password *string `json:"Password,omitempty" name:"Password"`

	// 通信方式
	CommProtocol *string `json:"CommProtocol,omitempty" name:"CommProtocol"`

	// qps
	Qps *uint64 `json:"Qps,omitempty" name:"Qps"`

	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 产品的设备类型
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// 关联的产品列表
	AssociatedProducts []*string `json:"AssociatedProducts,omitempty" name:"AssociatedProducts"`
}

type ProductEntry struct {
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 产品Key
	ProductKey *string `json:"ProductKey,omitempty" name:"ProductKey"`

	// AppId
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// 产品名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 产品描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 连接域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 鉴权类型（0：直连，1：Token）
	AuthType *uint64 `json:"AuthType,omitempty" name:"AuthType"`

	// 数据协议（native/template）
	DataProtocol *string `json:"DataProtocol,omitempty" name:"DataProtocol"`

	// 删除（0未删除）
	Deleted *uint64 `json:"Deleted,omitempty" name:"Deleted"`

	// 备注
	Message *string `json:"Message,omitempty" name:"Message"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 通信方式
	CommProtocol *string `json:"CommProtocol,omitempty" name:"CommProtocol"`

	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 设备类型
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`
}

// Predefined struct for user
type PublishMsgRequestParams struct {
	// Topic
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// 消息内容
	Message *string `json:"Message,omitempty" name:"Message"`

	// Qos(目前QoS支持0与1)
	Qos *int64 `json:"Qos,omitempty" name:"Qos"`
}

type PublishMsgRequest struct {
	*tchttp.BaseRequest
	
	// Topic
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// 消息内容
	Message *string `json:"Message,omitempty" name:"Message"`

	// Qos(目前QoS支持0与1)
	Qos *int64 `json:"Qos,omitempty" name:"Qos"`
}

func (r *PublishMsgRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PublishMsgRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Topic")
	delete(f, "Message")
	delete(f, "Qos")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PublishMsgRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PublishMsgResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type PublishMsgResponse struct {
	*tchttp.BaseResponse
	Response *PublishMsgResponseParams `json:"Response"`
}

func (r *PublishMsgResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PublishMsgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetDeviceRequestParams struct {
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

type ResetDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

func (r *ResetDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetDeviceResponseParams struct {
	// 设备信息
	Device *Device `json:"Device,omitempty" name:"Device"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ResetDeviceResponse struct {
	*tchttp.BaseResponse
	Response *ResetDeviceResponseParams `json:"Response"`
}

func (r *ResetDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Rule struct {
	// 规则Id
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// AppId
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 查询
	Query *RuleQuery `json:"Query,omitempty" name:"Query"`

	// 转发
	Actions []*Action `json:"Actions,omitempty" name:"Actions"`

	// 已启动
	Active *uint64 `json:"Active,omitempty" name:"Active"`

	// 已删除
	Deleted *uint64 `json:"Deleted,omitempty" name:"Deleted"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 消息顺序
	MsgOrder *uint64 `json:"MsgOrder,omitempty" name:"MsgOrder"`

	// 数据类型（0：文本，1：二进制）
	DataType *uint64 `json:"DataType,omitempty" name:"DataType"`
}

type RuleQuery struct {
	// 字段
	Field *string `json:"Field,omitempty" name:"Field"`

	// 过滤规则
	Condition *string `json:"Condition,omitempty" name:"Condition"`

	// Topic
	// 注意：此字段可能返回 null，表示取不到有效值。
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// 产品Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

type ServiceAction struct {
	// 服务url地址
	Url *string `json:"Url,omitempty" name:"Url"`
}

type StringData struct {
	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 描述
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// 读写模式
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// 长度范围
	Range []*uint64 `json:"Range,omitempty" name:"Range"`
}

type Topic struct {
	// TopicId
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Topic名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 消息最大生命周期
	MsgLife *uint64 `json:"MsgLife,omitempty" name:"MsgLife"`

	// 消息最大大小
	MsgSize *uint64 `json:"MsgSize,omitempty" name:"MsgSize"`

	// 消息最大数量
	MsgCount *uint64 `json:"MsgCount,omitempty" name:"MsgCount"`

	// 已删除
	Deleted *uint64 `json:"Deleted,omitempty" name:"Deleted"`

	// Topic完整路径
	Path *string `json:"Path,omitempty" name:"Path"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type TopicAction struct {
	// 目标topic
	Topic *string `json:"Topic,omitempty" name:"Topic"`
}

// Predefined struct for user
type UnassociateSubDeviceFromGatewayProductRequestParams struct {
	// 子设备产品Id
	SubDeviceProductId *string `json:"SubDeviceProductId,omitempty" name:"SubDeviceProductId"`

	// 网关设备产品Id
	GatewayProductId *string `json:"GatewayProductId,omitempty" name:"GatewayProductId"`
}

type UnassociateSubDeviceFromGatewayProductRequest struct {
	*tchttp.BaseRequest
	
	// 子设备产品Id
	SubDeviceProductId *string `json:"SubDeviceProductId,omitempty" name:"SubDeviceProductId"`

	// 网关设备产品Id
	GatewayProductId *string `json:"GatewayProductId,omitempty" name:"GatewayProductId"`
}

func (r *UnassociateSubDeviceFromGatewayProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnassociateSubDeviceFromGatewayProductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubDeviceProductId")
	delete(f, "GatewayProductId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnassociateSubDeviceFromGatewayProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnassociateSubDeviceFromGatewayProductResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnassociateSubDeviceFromGatewayProductResponse struct {
	*tchttp.BaseResponse
	Response *UnassociateSubDeviceFromGatewayProductResponseParams `json:"Response"`
}

func (r *UnassociateSubDeviceFromGatewayProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnassociateSubDeviceFromGatewayProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateProductRequestParams struct {
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 产品名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 产品描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 数据模版
	DataTemplate []*DataTemplate `json:"DataTemplate,omitempty" name:"DataTemplate"`
}

type UpdateProductRequest struct {
	*tchttp.BaseRequest
	
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 产品名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 产品描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 数据模版
	DataTemplate []*DataTemplate `json:"DataTemplate,omitempty" name:"DataTemplate"`
}

func (r *UpdateProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateProductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "DataTemplate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateProductResponseParams struct {
	// 更新后的产品信息
	Product *Product `json:"Product,omitempty" name:"Product"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateProductResponse struct {
	*tchttp.BaseResponse
	Response *UpdateProductResponseParams `json:"Response"`
}

func (r *UpdateProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRuleRequestParams struct {
	// 规则Id
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 查询
	Query *RuleQuery `json:"Query,omitempty" name:"Query"`

	// 转发动作列表
	Actions []*Action `json:"Actions,omitempty" name:"Actions"`

	// 数据类型（0：文本，1：二进制）
	DataType *uint64 `json:"DataType,omitempty" name:"DataType"`
}

type UpdateRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则Id
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 查询
	Query *RuleQuery `json:"Query,omitempty" name:"Query"`

	// 转发动作列表
	Actions []*Action `json:"Actions,omitempty" name:"Actions"`

	// 数据类型（0：文本，1：二进制）
	DataType *uint64 `json:"DataType,omitempty" name:"DataType"`
}

func (r *UpdateRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "Query")
	delete(f, "Actions")
	delete(f, "DataType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRuleResponseParams struct {
	// 规则
	Rule *Rule `json:"Rule,omitempty" name:"Rule"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateRuleResponse struct {
	*tchttp.BaseResponse
	Response *UpdateRuleResponseParams `json:"Response"`
}

func (r *UpdateRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}