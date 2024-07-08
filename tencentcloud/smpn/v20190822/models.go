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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type CHPRequest struct {
	// 电话号码
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`
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
	TagType *int64 `json:"TagType,omitnil,omitempty" name:"TagType"`

	// 标记次数
	TagCount *int64 `json:"TagCount,omitnil,omitempty" name:"TagCount"`
}

// Predefined struct for user
type DescribeSmpnChpRequestParams struct {
	// 客户用于计费的资源Id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 终端骚扰保护请求
	RequestData *CHPRequest `json:"RequestData,omitnil,omitempty" name:"RequestData"`
}

type DescribeSmpnChpRequest struct {
	*tchttp.BaseRequest
	
	// 客户用于计费的资源Id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 终端骚扰保护请求
	RequestData *CHPRequest `json:"RequestData,omitnil,omitempty" name:"RequestData"`
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
	ResponseData *CHPResponse `json:"ResponseData,omitnil,omitempty" name:"ResponseData"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	RequestData *FNRRequest `json:"RequestData,omitnil,omitempty" name:"RequestData"`

	// 用于计费的资源ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

type DescribeSmpnFnrRequest struct {
	*tchttp.BaseRequest
	
	// 虚假号码识别请求内容
	RequestData *FNRRequest `json:"RequestData,omitnil,omitempty" name:"RequestData"`

	// 用于计费的资源ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
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
	ResponseData *FNRResponse `json:"ResponseData,omitnil,omitempty" name:"ResponseData"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type FNRRequest struct {
	// 电话号码
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`
}

type FNRResponse struct {
	// 虚假号码描述
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}