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

package v20190822

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type CHPRequest struct {
	// 电话号码
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`
}

type CHPResponse struct {
	// 标记类型
	//  0: 无标记
	//  50:骚扰电话
	//  51:房产中介
	//  52:保险理财
	//  53:广告推销
	//  54:诈骗电话
	//  55:快递电话
	//  56:出租车专车
	TagType *int64 `json:"TagType,omitempty" name:"TagType"`

	// 标记次数
	TagCount *int64 `json:"TagCount,omitempty" name:"TagCount"`
}

// Predefined struct for user
type CreateSmpnEpaRequestParams struct {
	// 企业号码认证请求内容
	RequestData *EPARequest `json:"RequestData,omitempty" name:"RequestData"`

	// 用于计费的资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

type CreateSmpnEpaRequest struct {
	*tchttp.BaseRequest
	
	// 企业号码认证请求内容
	RequestData *EPARequest `json:"RequestData,omitempty" name:"RequestData"`

	// 用于计费的资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

func (r *CreateSmpnEpaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSmpnEpaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RequestData")
	delete(f, "ResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSmpnEpaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSmpnEpaResponseParams struct {
	// 业号码认证回应内容
	ResponseData *EPAResponse `json:"ResponseData,omitempty" name:"ResponseData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSmpnEpaResponse struct {
	*tchttp.BaseResponse
	Response *CreateSmpnEpaResponseParams `json:"Response"`
}

func (r *CreateSmpnEpaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSmpnEpaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSmpnChpRequestParams struct {
	// 客户用于计费的资源Id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 终端骚扰保护请求
	RequestData *CHPRequest `json:"RequestData,omitempty" name:"RequestData"`
}

type DescribeSmpnChpRequest struct {
	*tchttp.BaseRequest
	
	// 客户用于计费的资源Id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 终端骚扰保护请求
	RequestData *CHPRequest `json:"RequestData,omitempty" name:"RequestData"`
}

func (r *DescribeSmpnChpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSmpnChpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceId")
	delete(f, "RequestData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSmpnChpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSmpnChpResponseParams struct {
	// 终端骚扰保护回应
	ResponseData *CHPResponse `json:"ResponseData,omitempty" name:"ResponseData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSmpnChpResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSmpnChpResponseParams `json:"Response"`
}

func (r *DescribeSmpnChpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSmpnChpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSmpnFnrRequestParams struct {
	// 虚假号码识别请求内容
	RequestData *FNRRequest `json:"RequestData,omitempty" name:"RequestData"`

	// 用于计费的资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

type DescribeSmpnFnrRequest struct {
	*tchttp.BaseRequest
	
	// 虚假号码识别请求内容
	RequestData *FNRRequest `json:"RequestData,omitempty" name:"RequestData"`

	// 用于计费的资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

func (r *DescribeSmpnFnrRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSmpnFnrRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RequestData")
	delete(f, "ResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSmpnFnrRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSmpnFnrResponseParams struct {
	// 虚假号码识别回应内容
	ResponseData *FNRResponse `json:"ResponseData,omitempty" name:"ResponseData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSmpnFnrResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSmpnFnrResponseParams `json:"Response"`
}

func (r *DescribeSmpnFnrResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSmpnFnrResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSmpnMhmRequestParams struct {
	// 号码营销监控请求内容
	RequestData *MHMRequest `json:"RequestData,omitempty" name:"RequestData"`

	// 用于计费的资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

type DescribeSmpnMhmRequest struct {
	*tchttp.BaseRequest
	
	// 号码营销监控请求内容
	RequestData *MHMRequest `json:"RequestData,omitempty" name:"RequestData"`

	// 用于计费的资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

func (r *DescribeSmpnMhmRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSmpnMhmRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RequestData")
	delete(f, "ResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSmpnMhmRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSmpnMhmResponseParams struct {
	// 号码营销监控回应内容
	ResponseData *MHMResponse `json:"ResponseData,omitempty" name:"ResponseData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSmpnMhmResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSmpnMhmResponseParams `json:"Response"`
}

func (r *DescribeSmpnMhmResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSmpnMhmResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSmpnMrlRequestParams struct {
	// 恶意标记等级请求内容
	RequestData *MRLRequest `json:"RequestData,omitempty" name:"RequestData"`

	// 用于计费的资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

type DescribeSmpnMrlRequest struct {
	*tchttp.BaseRequest
	
	// 恶意标记等级请求内容
	RequestData *MRLRequest `json:"RequestData,omitempty" name:"RequestData"`

	// 用于计费的资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

func (r *DescribeSmpnMrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSmpnMrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RequestData")
	delete(f, "ResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSmpnMrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSmpnMrlResponseParams struct {
	// 恶意标记等级回应内容
	ResponseData *MRLResponse `json:"ResponseData,omitempty" name:"ResponseData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSmpnMrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSmpnMrlResponseParams `json:"Response"`
}

func (r *DescribeSmpnMrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSmpnMrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EPARequest struct {
	// 电话号码
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// 黄页名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

type EPAResponse struct {
	// 0成功 其他失败
	RetCode *int64 `json:"RetCode,omitempty" name:"RetCode"`
}

type FNRRequest struct {
	// 电话号码
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`
}

type FNRResponse struct {
	// 虚假号码描述
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type MHMRequest struct {
	// 电话号码
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`
}

type MHMResponse struct {
	// 标记类型
	//  0: 无标记
	//  50:骚扰电话
	//  51:房产中介
	//  52:保险理财
	//  53:广告推销
	//  54:诈骗电话
	//  55:快递电话
	//  56:出租车专车
	TagType *int64 `json:"TagType,omitempty" name:"TagType"`

	// 标记次数
	TagCount *int64 `json:"TagCount,omitempty" name:"TagCount"`
}

type MRLRequest struct {
	// 电话号码
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`
}

type MRLResponse struct {
	// 骚扰电话恶意标记等级
	DisturbLevel *int64 `json:"DisturbLevel,omitempty" name:"DisturbLevel"`

	// 房产中介恶意标记等级
	HouseAgentLevel *int64 `json:"HouseAgentLevel,omitempty" name:"HouseAgentLevel"`

	// 保险理财恶意标记等级
	InsuranceLevel *int64 `json:"InsuranceLevel,omitempty" name:"InsuranceLevel"`

	// 广告推销恶意标记等级
	SalesLevel *int64 `json:"SalesLevel,omitempty" name:"SalesLevel"`

	// 诈骗电话恶意标记等级
	CheatLevel *int64 `json:"CheatLevel,omitempty" name:"CheatLevel"`
}