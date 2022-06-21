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

package v20200720

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type CreateWeappQRUrlRequestParams struct {
	// 代理角色临时密钥的Token
	SessionKey *string `json:"SessionKey,omitempty" name:"SessionKey"`
}

type CreateWeappQRUrlRequest struct {
	*tchttp.BaseRequest
	
	// 代理角色临时密钥的Token
	SessionKey *string `json:"SessionKey,omitempty" name:"SessionKey"`
}

func (r *CreateWeappQRUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWeappQRUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWeappQRUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWeappQRUrlResponseParams struct {
	// 渠道备案小程序二维码
	Url *string `json:"Url,omitempty" name:"Url"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateWeappQRUrlResponse struct {
	*tchttp.BaseResponse
	Response *CreateWeappQRUrlResponseParams `json:"Response"`
}

func (r *CreateWeappQRUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWeappQRUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGetAuthInfoRequestParams struct {

}

type DescribeGetAuthInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeGetAuthInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGetAuthInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGetAuthInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGetAuthInfoResponseParams struct {
	// 实名认证状态：0未实名，1已实名
	IsTenPayMasked *string `json:"IsTenPayMasked,omitempty" name:"IsTenPayMasked"`

	// 实名认证类型：0个人，1企业
	IsAuthenticated *string `json:"IsAuthenticated,omitempty" name:"IsAuthenticated"`

	// 认证类型，个人0，企业1
	Type *string `json:"Type,omitempty" name:"Type"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeGetAuthInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGetAuthInfoResponseParams `json:"Response"`
}

func (r *DescribeGetAuthInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGetAuthInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SyncIcpOrderWebInfoRequestParams struct {
	// 备案ICP订单号
	IcpOrderId *string `json:"IcpOrderId,omitempty" name:"IcpOrderId"`

	// 订单里的webId
	SourceWebId *string `json:"SourceWebId,omitempty" name:"SourceWebId"`

	// 订单里的webId 数组(如果传入的webIds含有 订单中不包含的webId，会自动跳过)
	TargetWebIds []*string `json:"TargetWebIds,omitempty" name:"TargetWebIds"`

	// 网站信息字段名 数组
	SyncFields []*string `json:"SyncFields,omitempty" name:"SyncFields"`

	// 是否先判断同步的网站负责人是否一致 (这里会判断 sitePersonName, sitePersonCerType,sitePersonCerNum三个字段完全一致)  默认:true. 非必要 不建议关闭修改该参数默认值
	CheckSamePerson *bool `json:"CheckSamePerson,omitempty" name:"CheckSamePerson"`
}

type SyncIcpOrderWebInfoRequest struct {
	*tchttp.BaseRequest
	
	// 备案ICP订单号
	IcpOrderId *string `json:"IcpOrderId,omitempty" name:"IcpOrderId"`

	// 订单里的webId
	SourceWebId *string `json:"SourceWebId,omitempty" name:"SourceWebId"`

	// 订单里的webId 数组(如果传入的webIds含有 订单中不包含的webId，会自动跳过)
	TargetWebIds []*string `json:"TargetWebIds,omitempty" name:"TargetWebIds"`

	// 网站信息字段名 数组
	SyncFields []*string `json:"SyncFields,omitempty" name:"SyncFields"`

	// 是否先判断同步的网站负责人是否一致 (这里会判断 sitePersonName, sitePersonCerType,sitePersonCerNum三个字段完全一致)  默认:true. 非必要 不建议关闭修改该参数默认值
	CheckSamePerson *bool `json:"CheckSamePerson,omitempty" name:"CheckSamePerson"`
}

func (r *SyncIcpOrderWebInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncIcpOrderWebInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IcpOrderId")
	delete(f, "SourceWebId")
	delete(f, "TargetWebIds")
	delete(f, "SyncFields")
	delete(f, "CheckSamePerson")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SyncIcpOrderWebInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SyncIcpOrderWebInfoResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SyncIcpOrderWebInfoResponse struct {
	*tchttp.BaseResponse
	Response *SyncIcpOrderWebInfoResponseParams `json:"Response"`
}

func (r *SyncIcpOrderWebInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncIcpOrderWebInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}