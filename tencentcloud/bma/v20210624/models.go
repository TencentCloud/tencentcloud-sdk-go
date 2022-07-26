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

package v20210624

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type CreateCRBlockRequestParams struct {
	// 已存证的作品ID
	WorkId *int64 `json:"WorkId,omitempty" name:"WorkId"`

	// 侵权链接
	TortUrl *string `json:"TortUrl,omitempty" name:"TortUrl"`

	// 侵权标题
	TortTitle *string `json:"TortTitle,omitempty" name:"TortTitle"`

	// 侵权平台
	TortPlat *string `json:"TortPlat,omitempty" name:"TortPlat"`

	// 拦截结果回调地址
	BlockUrl *string `json:"BlockUrl,omitempty" name:"BlockUrl"`

	// x
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// x
	ValidStartDate *string `json:"ValidStartDate,omitempty" name:"ValidStartDate"`

	// x
	ValidEndDate *string `json:"ValidEndDate,omitempty" name:"ValidEndDate"`

	// xx
	TortPic *string `json:"TortPic,omitempty" name:"TortPic"`

	// x
	CommFileUrl *string `json:"CommFileUrl,omitempty" name:"CommFileUrl"`

	// x
	CommValidStartDate *string `json:"CommValidStartDate,omitempty" name:"CommValidStartDate"`

	// x
	CommValidEndDate *string `json:"CommValidEndDate,omitempty" name:"CommValidEndDate"`

	// x
	IsProducer *string `json:"IsProducer,omitempty" name:"IsProducer"`
}

type CreateCRBlockRequest struct {
	*tchttp.BaseRequest
	
	// 已存证的作品ID
	WorkId *int64 `json:"WorkId,omitempty" name:"WorkId"`

	// 侵权链接
	TortUrl *string `json:"TortUrl,omitempty" name:"TortUrl"`

	// 侵权标题
	TortTitle *string `json:"TortTitle,omitempty" name:"TortTitle"`

	// 侵权平台
	TortPlat *string `json:"TortPlat,omitempty" name:"TortPlat"`

	// 拦截结果回调地址
	BlockUrl *string `json:"BlockUrl,omitempty" name:"BlockUrl"`

	// x
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// x
	ValidStartDate *string `json:"ValidStartDate,omitempty" name:"ValidStartDate"`

	// x
	ValidEndDate *string `json:"ValidEndDate,omitempty" name:"ValidEndDate"`

	// xx
	TortPic *string `json:"TortPic,omitempty" name:"TortPic"`

	// x
	CommFileUrl *string `json:"CommFileUrl,omitempty" name:"CommFileUrl"`

	// x
	CommValidStartDate *string `json:"CommValidStartDate,omitempty" name:"CommValidStartDate"`

	// x
	CommValidEndDate *string `json:"CommValidEndDate,omitempty" name:"CommValidEndDate"`

	// x
	IsProducer *string `json:"IsProducer,omitempty" name:"IsProducer"`
}

func (r *CreateCRBlockRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCRBlockRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkId")
	delete(f, "TortUrl")
	delete(f, "TortTitle")
	delete(f, "TortPlat")
	delete(f, "BlockUrl")
	delete(f, "FileUrl")
	delete(f, "ValidStartDate")
	delete(f, "ValidEndDate")
	delete(f, "TortPic")
	delete(f, "CommFileUrl")
	delete(f, "CommValidStartDate")
	delete(f, "CommValidEndDate")
	delete(f, "IsProducer")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCRBlockRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCRBlockResponseParams struct {
	// 侵权ID
	TortId *int64 `json:"TortId,omitempty" name:"TortId"`

	// xxx
	TortNum *string `json:"TortNum,omitempty" name:"TortNum"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCRBlockResponse struct {
	*tchttp.BaseResponse
	Response *CreateCRBlockResponseParams `json:"Response"`
}

func (r *CreateCRBlockResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCRBlockResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCRCompanyVerifyRequestParams struct {
	// 企业名称
	CompanyName *string `json:"CompanyName,omitempty" name:"CompanyName"`

	// 企业证件号码
	CompanyID *string `json:"CompanyID,omitempty" name:"CompanyID"`

	// 企业法人姓名
	CompanyLegalName *string `json:"CompanyLegalName,omitempty" name:"CompanyLegalName"`

	// 管理员名称
	ManagerName *string `json:"ManagerName,omitempty" name:"ManagerName"`

	// 管理员手机号
	ManagerPhone *string `json:"ManagerPhone,omitempty" name:"ManagerPhone"`

	// 手机验证码
	VerificationCode *string `json:"VerificationCode,omitempty" name:"VerificationCode"`

	// 企业认证号码类型 1：社会信用代码 2：组织机构代码 3：企业工商注册码 4：其他 默认为1
	CompanyIDType *string `json:"CompanyIDType,omitempty" name:"CompanyIDType"`

	// xxx
	Type *string `json:"Type,omitempty" name:"Type"`
}

type CreateCRCompanyVerifyRequest struct {
	*tchttp.BaseRequest
	
	// 企业名称
	CompanyName *string `json:"CompanyName,omitempty" name:"CompanyName"`

	// 企业证件号码
	CompanyID *string `json:"CompanyID,omitempty" name:"CompanyID"`

	// 企业法人姓名
	CompanyLegalName *string `json:"CompanyLegalName,omitempty" name:"CompanyLegalName"`

	// 管理员名称
	ManagerName *string `json:"ManagerName,omitempty" name:"ManagerName"`

	// 管理员手机号
	ManagerPhone *string `json:"ManagerPhone,omitempty" name:"ManagerPhone"`

	// 手机验证码
	VerificationCode *string `json:"VerificationCode,omitempty" name:"VerificationCode"`

	// 企业认证号码类型 1：社会信用代码 2：组织机构代码 3：企业工商注册码 4：其他 默认为1
	CompanyIDType *string `json:"CompanyIDType,omitempty" name:"CompanyIDType"`

	// xxx
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *CreateCRCompanyVerifyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCRCompanyVerifyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CompanyName")
	delete(f, "CompanyID")
	delete(f, "CompanyLegalName")
	delete(f, "ManagerName")
	delete(f, "ManagerPhone")
	delete(f, "VerificationCode")
	delete(f, "CompanyIDType")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCRCompanyVerifyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCRCompanyVerifyResponseParams struct {
	// 认证状态 0-认证成功 1-认证失败
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 认证结果返回
	Note *string `json:"Note,omitempty" name:"Note"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCRCompanyVerifyResponse struct {
	*tchttp.BaseResponse
	Response *CreateCRCompanyVerifyResponseParams `json:"Response"`
}

func (r *CreateCRCompanyVerifyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCRCompanyVerifyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCRRightRequestParams struct {
	// 已存证的作品ID
	WorkId *int64 `json:"WorkId,omitempty" name:"WorkId"`

	// 侵权链接
	TortUrl *string `json:"TortUrl,omitempty" name:"TortUrl"`

	// 侵权标题
	TortTitle *string `json:"TortTitle,omitempty" name:"TortTitle"`

	// 侵权平台
	TortPlat *string `json:"TortPlat,omitempty" name:"TortPlat"`

	// 发函结果回调地址
	RightUrl *string `json:"RightUrl,omitempty" name:"RightUrl"`

	// x
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// x
	ValidStartDate *string `json:"ValidStartDate,omitempty" name:"ValidStartDate"`

	// x
	ValidEndDate *string `json:"ValidEndDate,omitempty" name:"ValidEndDate"`

	// x
	CommFileUrl *string `json:"CommFileUrl,omitempty" name:"CommFileUrl"`

	// x
	CommValidStartDate *string `json:"CommValidStartDate,omitempty" name:"CommValidStartDate"`

	// x
	CommValidEndDate *string `json:"CommValidEndDate,omitempty" name:"CommValidEndDate"`

	// x
	HomeFileUrl *string `json:"HomeFileUrl,omitempty" name:"HomeFileUrl"`

	// x
	HomeValidStartDate *string `json:"HomeValidStartDate,omitempty" name:"HomeValidStartDate"`

	// x
	HomeValidEndDate *string `json:"HomeValidEndDate,omitempty" name:"HomeValidEndDate"`

	// x
	IsProducer *string `json:"IsProducer,omitempty" name:"IsProducer"`
}

type CreateCRRightRequest struct {
	*tchttp.BaseRequest
	
	// 已存证的作品ID
	WorkId *int64 `json:"WorkId,omitempty" name:"WorkId"`

	// 侵权链接
	TortUrl *string `json:"TortUrl,omitempty" name:"TortUrl"`

	// 侵权标题
	TortTitle *string `json:"TortTitle,omitempty" name:"TortTitle"`

	// 侵权平台
	TortPlat *string `json:"TortPlat,omitempty" name:"TortPlat"`

	// 发函结果回调地址
	RightUrl *string `json:"RightUrl,omitempty" name:"RightUrl"`

	// x
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// x
	ValidStartDate *string `json:"ValidStartDate,omitempty" name:"ValidStartDate"`

	// x
	ValidEndDate *string `json:"ValidEndDate,omitempty" name:"ValidEndDate"`

	// x
	CommFileUrl *string `json:"CommFileUrl,omitempty" name:"CommFileUrl"`

	// x
	CommValidStartDate *string `json:"CommValidStartDate,omitempty" name:"CommValidStartDate"`

	// x
	CommValidEndDate *string `json:"CommValidEndDate,omitempty" name:"CommValidEndDate"`

	// x
	HomeFileUrl *string `json:"HomeFileUrl,omitempty" name:"HomeFileUrl"`

	// x
	HomeValidStartDate *string `json:"HomeValidStartDate,omitempty" name:"HomeValidStartDate"`

	// x
	HomeValidEndDate *string `json:"HomeValidEndDate,omitempty" name:"HomeValidEndDate"`

	// x
	IsProducer *string `json:"IsProducer,omitempty" name:"IsProducer"`
}

func (r *CreateCRRightRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCRRightRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkId")
	delete(f, "TortUrl")
	delete(f, "TortTitle")
	delete(f, "TortPlat")
	delete(f, "RightUrl")
	delete(f, "FileUrl")
	delete(f, "ValidStartDate")
	delete(f, "ValidEndDate")
	delete(f, "CommFileUrl")
	delete(f, "CommValidStartDate")
	delete(f, "CommValidEndDate")
	delete(f, "HomeFileUrl")
	delete(f, "HomeValidStartDate")
	delete(f, "HomeValidEndDate")
	delete(f, "IsProducer")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCRRightRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCRRightResponseParams struct {
	// 侵权ID
	TortId *int64 `json:"TortId,omitempty" name:"TortId"`

	// xxx
	TortNum *string `json:"TortNum,omitempty" name:"TortNum"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCRRightResponse struct {
	*tchttp.BaseResponse
	Response *CreateCRRightResponseParams `json:"Response"`
}

func (r *CreateCRRightResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCRRightResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCRWorkInfoRequestParams struct {
	// xxx
	WorkId *int64 `json:"WorkId,omitempty" name:"WorkId"`
}

type DescribeCRWorkInfoRequest struct {
	*tchttp.BaseRequest
	
	// xxx
	WorkId *int64 `json:"WorkId,omitempty" name:"WorkId"`
}

func (r *DescribeCRWorkInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCRWorkInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCRWorkInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCRWorkInfoResponseParams struct {
	// x
	WorkName *string `json:"WorkName,omitempty" name:"WorkName"`

	// x
	MonitorStatus *int64 `json:"MonitorStatus,omitempty" name:"MonitorStatus"`

	// x
	AuthStatus *int64 `json:"AuthStatus,omitempty" name:"AuthStatus"`

	// x
	CommStatus *int64 `json:"CommStatus,omitempty" name:"CommStatus"`

	// x
	IsProducer *int64 `json:"IsProducer,omitempty" name:"IsProducer"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCRWorkInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCRWorkInfoResponseParams `json:"Response"`
}

func (r *DescribeCRWorkInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCRWorkInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}