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

package v20210701

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type DescribeNewUserAcquisitionRequestParams struct {
	// 用户信息
	ServiceParams *UserInfos `json:"ServiceParams,omitempty" name:"ServiceParams"`
}

type DescribeNewUserAcquisitionRequest struct {
	*tchttp.BaseRequest
	
	// 用户信息
	ServiceParams *UserInfos `json:"ServiceParams,omitempty" name:"ServiceParams"`
}

func (r *DescribeNewUserAcquisitionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNewUserAcquisitionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceParams")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNewUserAcquisitionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNewUserAcquisitionResponseParams struct {
	// 用户信誉分，1-5从低到高
	ServiceRsp *Score `json:"ServiceRsp,omitempty" name:"ServiceRsp"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeNewUserAcquisitionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNewUserAcquisitionResponseParams `json:"Response"`
}

func (r *DescribeNewUserAcquisitionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNewUserAcquisitionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStockEstimationRequestParams struct {
	// 用户信息
	ServiceParams *UserInfos `json:"ServiceParams,omitempty" name:"ServiceParams"`
}

type DescribeStockEstimationRequest struct {
	*tchttp.BaseRequest
	
	// 用户信息
	ServiceParams *UserInfos `json:"ServiceParams,omitempty" name:"ServiceParams"`
}

func (r *DescribeStockEstimationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStockEstimationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceParams")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStockEstimationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStockEstimationResponseParams struct {
	// 用户信誉分，1-5从低到高
	ServiceRsp *Score `json:"ServiceRsp,omitempty" name:"ServiceRsp"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeStockEstimationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStockEstimationResponseParams `json:"Response"`
}

func (r *DescribeStockEstimationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStockEstimationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Score struct {
	// 信誉分，1-5从低到高
	Star *int64 `json:"Star,omitempty" name:"Star"`
}

type UserInfos struct {
	// 用户的手机号
	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`

	// 用户的微信OpenID
	Openid *string `json:"Openid,omitempty" name:"Openid"`

	// 用户移动设备的客户端IP
	IP *string `json:"IP,omitempty" name:"IP"`

	// 用户WiFi的BSSID
	WiFiBssid *string `json:"WiFiBssid,omitempty" name:"WiFiBssid"`

	// 用户Android设备的IMEI
	IMEI *string `json:"IMEI,omitempty" name:"IMEI"`

	// 用户Android设备的OAID
	OAID *string `json:"OAID,omitempty" name:"OAID"`

	// 用户iOS设备的IDFA
	IDFA *string `json:"IDFA,omitempty" name:"IDFA"`
}